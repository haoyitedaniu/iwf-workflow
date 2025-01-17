/*
Workflow APIs

This APIs for iwf SDKs to operate workflows

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package iwfidl

import (
	"encoding/json"
)

// WorkflowSearchResponseEntry struct for WorkflowSearchResponseEntry
type WorkflowSearchResponseEntry struct {
	WorkflowId    string `json:"workflowId"`
	WorkflowRunId string `json:"workflowRunId"`
}

// NewWorkflowSearchResponseEntry instantiates a new WorkflowSearchResponseEntry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowSearchResponseEntry(workflowId string, workflowRunId string) *WorkflowSearchResponseEntry {
	this := WorkflowSearchResponseEntry{}
	this.WorkflowId = workflowId
	this.WorkflowRunId = workflowRunId
	return &this
}

// NewWorkflowSearchResponseEntryWithDefaults instantiates a new WorkflowSearchResponseEntry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowSearchResponseEntryWithDefaults() *WorkflowSearchResponseEntry {
	this := WorkflowSearchResponseEntry{}
	return &this
}

// GetWorkflowId returns the WorkflowId field value
func (o *WorkflowSearchResponseEntry) GetWorkflowId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WorkflowId
}

// GetWorkflowIdOk returns a tuple with the WorkflowId field value
// and a boolean to check if the value has been set.
func (o *WorkflowSearchResponseEntry) GetWorkflowIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WorkflowId, true
}

// SetWorkflowId sets field value
func (o *WorkflowSearchResponseEntry) SetWorkflowId(v string) {
	o.WorkflowId = v
}

// GetWorkflowRunId returns the WorkflowRunId field value
func (o *WorkflowSearchResponseEntry) GetWorkflowRunId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WorkflowRunId
}

// GetWorkflowRunIdOk returns a tuple with the WorkflowRunId field value
// and a boolean to check if the value has been set.
func (o *WorkflowSearchResponseEntry) GetWorkflowRunIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WorkflowRunId, true
}

// SetWorkflowRunId sets field value
func (o *WorkflowSearchResponseEntry) SetWorkflowRunId(v string) {
	o.WorkflowRunId = v
}

func (o WorkflowSearchResponseEntry) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["workflowId"] = o.WorkflowId
	}
	if true {
		toSerialize["workflowRunId"] = o.WorkflowRunId
	}
	return json.Marshal(toSerialize)
}

type NullableWorkflowSearchResponseEntry struct {
	value *WorkflowSearchResponseEntry
	isSet bool
}

func (v NullableWorkflowSearchResponseEntry) Get() *WorkflowSearchResponseEntry {
	return v.value
}

func (v *NullableWorkflowSearchResponseEntry) Set(val *WorkflowSearchResponseEntry) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowSearchResponseEntry) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowSearchResponseEntry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowSearchResponseEntry(val *WorkflowSearchResponseEntry) *NullableWorkflowSearchResponseEntry {
	return &NullableWorkflowSearchResponseEntry{value: val, isSet: true}
}

func (v NullableWorkflowSearchResponseEntry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowSearchResponseEntry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
