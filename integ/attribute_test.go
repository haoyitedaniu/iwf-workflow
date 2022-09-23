package integ

import (
	"context"
	"fmt"
	"github.com/cadence-oss/iwf-server/gen/iwfidl"
	"github.com/cadence-oss/iwf-server/integ/attribute"
	"github.com/cadence-oss/iwf-server/service"
	"github.com/cadence-oss/iwf-server/service/api"
	"github.com/cadence-oss/iwf-server/service/interpreter/temporal"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"strconv"
	"testing"
	"time"
)

func TestAttributeWorkflow(t *testing.T) {
	// start test workflow server
	wfHandler, basicWorkflow := attribute.NewAttributeWorkflow()
	testWorkflowServerPort := "9714"
	wfServer := &http.Server{
		Addr:    ":" + testWorkflowServerPort,
		Handler: basicWorkflow,
	}
	defer wfServer.Close()
	go func() {
		if err := wfServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// start iwf api server
	temporalClient := createTemporalClient()
	iwfService := api.NewService(temporalClient)
	testIwfServerPort := "9715"
	iwfServer := &http.Server{
		Addr:    ":" + testIwfServerPort,
		Handler: iwfService,
	}
	defer iwfServer.Close()
	go func() {
		if err := iwfServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// start iwf interpreter worker
	interpreter := temporal.NewInterpreterWorker(temporalClient)
	interpreter.Start()
	defer interpreter.Close()

	// start a workflow
	apiClient := iwfidl.NewAPIClient(&iwfidl.Configuration{
		Servers: []iwfidl.ServerConfiguration{
			{
				URL: "http://localhost:" + testIwfServerPort,
			},
		},
	})
	wfId := attribute.WorkflowType + strconv.Itoa(int(time.Now().Unix()))
	req := apiClient.DefaultApi.ApiV1WorkflowStartPost(context.Background())
	resp, httpResp, err := req.WorkflowStartRequest(iwfidl.WorkflowStartRequest{
		WorkflowId:             iwfidl.PtrString(wfId),
		IwfWorkflowType:        iwfidl.PtrString(attribute.WorkflowType),
		WorkflowTimeoutSeconds: iwfidl.PtrInt32(10),
		IwfWorkerUrl:           iwfidl.PtrString("http://localhost:" + testWorkflowServerPort),
		StartStateId:           iwfidl.PtrString(attribute.State1),
	}).Execute()
	if err != nil {
		log.Fatalf("Fail to invoke start api %v", err)
	}
	if httpResp.StatusCode != http.StatusOK {
		log.Fatalf("Status not success" + httpResp.Status)
	}
	fmt.Println(*resp)
	defer temporalClient.TerminateWorkflow(context.Background(), wfId, "", "terminate incase not completed")

	// wait for the workflow
	run := temporalClient.GetWorkflow(context.Background(), wfId, "")
	_ = run.Get(context.Background(), nil)

	qres, err := temporalClient.QueryWorkflow(context.Background(), wfId, "", service.AttributeQueryType, nil)
	if err != nil {
		log.Fatalf("Fail to invoke query workflow for all attrs%v", err)
	}

	var queryResult1 service.QueryAttributeResponse
	err = qres.Get(&queryResult1)
	if err != nil {
		log.Fatalf("Fail to invoke query workflow for all attrs%v", err)
	}

	//qres, err = temporalClient.QueryWorkflow(context.Background(), wfId, "", service.AttributeQueryType, service.QueryAttributeRequest{
	//	Keys: []string{
	//		attribute.TestQueryAttributeKey,
	//	},
	//})

	req2 := apiClient.DefaultApi.ApiV1WorkflowQueryPost(context.Background())
	_, httpResp2, err := req2.WorkflowQueryRequest(iwfidl.WorkflowQueryRequest{
		WorkflowId: iwfidl.PtrString(wfId),
		AttributeKeys: []string{
			attribute.TestQueryAttributeKey,
		},
	}).Execute()

	if err != nil || httpResp2.StatusCode != 200 {
		log.Fatalf("Fail to invoke query workflow for sigle attr %v %v", err, httpResp2)
	}

	var queryResult2 service.QueryAttributeResponse
	err = qres.Get(&queryResult2)
	if err != nil {
		log.Fatalf("Fail to invoke query workflow for single attr %v", err)
	}

	// assertion
	history, data := wfHandler.GetTestResult()
	assertions := assert.New(t)
	assertions.Equalf(map[string]int64{
		"S1_start":  1,
		"S1_decide": 1,
		"S2_start":  1,
		"S2_decide": 1,
	}, history, "attribute test fail, %v", history)

	if attribute.EnableTestingSearchAttribute {
		assertions.Equal(map[string]interface{}{
			"S1_decide_intSaFounds": 1,
			"S1_decide_kwSaFounds":  1,
			"S2_decide_intSaFounds": 1,
			"S2_decide_kwSaFounds":  1,
			"S2_start_intSaFounds":  1,
			"S2_start_kwSaFounds":   1,

			"S1_decide_localAttFound": true,
			"S1_decide_queryAttFound": true,
			"S2_decide_queryAttFound": true,
			"S2_start_queryAttFound":  true,
		}, data)
	} else {
		// map[S1_decide_intSaFounds:0 S1_decide_kwSaFounds:0 S1_decide_localAttFound:false
		//S1_decide_queryAttFound:true S2_decide_intSaFounds:0 S2_decide_kwSaFounds:0
		//S2_decide_queryAttFound:false S2_start_intSaFounds:0 S2_start_kwSaFounds:0 S2_start_queryAttFound:false]
		assertions.Equal(map[string]interface{}{
			"S1_decide_intSaFounds": 0,
			"S1_decide_kwSaFounds":  0,
			"S2_decide_intSaFounds": 0,
			"S2_decide_kwSaFounds":  0,
			"S2_start_intSaFounds":  0,
			"S2_start_kwSaFounds":   0,

			"S1_decide_localAttFound": true,
			"S1_decide_queryAttFound": true,
			"S2_decide_queryAttFound": true,
			"S2_start_queryAttFound":  true,
		}, data)
	}

	expected := []iwfidl.KeyValue{
		{
			Key:   iwfidl.PtrString(attribute.TestQueryAttributeKey),
			Value: &attribute.TestQueryVal2,
		},
	}
	assertions.Equal(expected, queryResult2.AttributeValues)
	assertions.Equal(expected, queryResult1.AttributeValues)
}