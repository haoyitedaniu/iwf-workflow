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

// WorkflowStateStartRequest struct for WorkflowStateStartRequest
type WorkflowStateStartRequest struct {
	Context *Context `json:"context,omitempty"`
	WorkflowType *string `json:"workflowType,omitempty"`
	WorkflowStateId *string `json:"workflowStateId,omitempty"`
	StateInput *EncodedObject `json:"stateInput,omitempty"`
	SearchAttributes []SearchAttribute `json:"searchAttributes,omitempty"`
	QueryAttributes []KeyValue `json:"queryAttributes,omitempty"`
}

// NewWorkflowStateStartRequest instantiates a new WorkflowStateStartRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowStateStartRequest() *WorkflowStateStartRequest {
	this := WorkflowStateStartRequest{}
	return &this
}

// NewWorkflowStateStartRequestWithDefaults instantiates a new WorkflowStateStartRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowStateStartRequestWithDefaults() *WorkflowStateStartRequest {
	this := WorkflowStateStartRequest{}
	return &this
}

// GetContext returns the Context field value if set, zero value otherwise.
func (o *WorkflowStateStartRequest) GetContext() Context {
	if o == nil || o.Context == nil {
		var ret Context
		return ret
	}
	return *o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowStateStartRequest) GetContextOk() (*Context, bool) {
	if o == nil || o.Context == nil {
		return nil, false
	}
	return o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *WorkflowStateStartRequest) HasContext() bool {
	if o != nil && o.Context != nil {
		return true
	}

	return false
}

// SetContext gets a reference to the given Context and assigns it to the Context field.
func (o *WorkflowStateStartRequest) SetContext(v Context) {
	o.Context = &v
}

// GetWorkflowType returns the WorkflowType field value if set, zero value otherwise.
func (o *WorkflowStateStartRequest) GetWorkflowType() string {
	if o == nil || o.WorkflowType == nil {
		var ret string
		return ret
	}
	return *o.WorkflowType
}

// GetWorkflowTypeOk returns a tuple with the WorkflowType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowStateStartRequest) GetWorkflowTypeOk() (*string, bool) {
	if o == nil || o.WorkflowType == nil {
		return nil, false
	}
	return o.WorkflowType, true
}

// HasWorkflowType returns a boolean if a field has been set.
func (o *WorkflowStateStartRequest) HasWorkflowType() bool {
	if o != nil && o.WorkflowType != nil {
		return true
	}

	return false
}

// SetWorkflowType gets a reference to the given string and assigns it to the WorkflowType field.
func (o *WorkflowStateStartRequest) SetWorkflowType(v string) {
	o.WorkflowType = &v
}

// GetWorkflowStateId returns the WorkflowStateId field value if set, zero value otherwise.
func (o *WorkflowStateStartRequest) GetWorkflowStateId() string {
	if o == nil || o.WorkflowStateId == nil {
		var ret string
		return ret
	}
	return *o.WorkflowStateId
}

// GetWorkflowStateIdOk returns a tuple with the WorkflowStateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowStateStartRequest) GetWorkflowStateIdOk() (*string, bool) {
	if o == nil || o.WorkflowStateId == nil {
		return nil, false
	}
	return o.WorkflowStateId, true
}

// HasWorkflowStateId returns a boolean if a field has been set.
func (o *WorkflowStateStartRequest) HasWorkflowStateId() bool {
	if o != nil && o.WorkflowStateId != nil {
		return true
	}

	return false
}

// SetWorkflowStateId gets a reference to the given string and assigns it to the WorkflowStateId field.
func (o *WorkflowStateStartRequest) SetWorkflowStateId(v string) {
	o.WorkflowStateId = &v
}

// GetStateInput returns the StateInput field value if set, zero value otherwise.
func (o *WorkflowStateStartRequest) GetStateInput() EncodedObject {
	if o == nil || o.StateInput == nil {
		var ret EncodedObject
		return ret
	}
	return *o.StateInput
}

// GetStateInputOk returns a tuple with the StateInput field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowStateStartRequest) GetStateInputOk() (*EncodedObject, bool) {
	if o == nil || o.StateInput == nil {
		return nil, false
	}
	return o.StateInput, true
}

// HasStateInput returns a boolean if a field has been set.
func (o *WorkflowStateStartRequest) HasStateInput() bool {
	if o != nil && o.StateInput != nil {
		return true
	}

	return false
}

// SetStateInput gets a reference to the given EncodedObject and assigns it to the StateInput field.
func (o *WorkflowStateStartRequest) SetStateInput(v EncodedObject) {
	o.StateInput = &v
}

// GetSearchAttributes returns the SearchAttributes field value if set, zero value otherwise.
func (o *WorkflowStateStartRequest) GetSearchAttributes() []SearchAttribute {
	if o == nil || o.SearchAttributes == nil {
		var ret []SearchAttribute
		return ret
	}
	return o.SearchAttributes
}

// GetSearchAttributesOk returns a tuple with the SearchAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowStateStartRequest) GetSearchAttributesOk() ([]SearchAttribute, bool) {
	if o == nil || o.SearchAttributes == nil {
		return nil, false
	}
	return o.SearchAttributes, true
}

// HasSearchAttributes returns a boolean if a field has been set.
func (o *WorkflowStateStartRequest) HasSearchAttributes() bool {
	if o != nil && o.SearchAttributes != nil {
		return true
	}

	return false
}

// SetSearchAttributes gets a reference to the given []SearchAttribute and assigns it to the SearchAttributes field.
func (o *WorkflowStateStartRequest) SetSearchAttributes(v []SearchAttribute) {
	o.SearchAttributes = v
}

// GetQueryAttributes returns the QueryAttributes field value if set, zero value otherwise.
func (o *WorkflowStateStartRequest) GetQueryAttributes() []KeyValue {
	if o == nil || o.QueryAttributes == nil {
		var ret []KeyValue
		return ret
	}
	return o.QueryAttributes
}

// GetQueryAttributesOk returns a tuple with the QueryAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowStateStartRequest) GetQueryAttributesOk() ([]KeyValue, bool) {
	if o == nil || o.QueryAttributes == nil {
		return nil, false
	}
	return o.QueryAttributes, true
}

// HasQueryAttributes returns a boolean if a field has been set.
func (o *WorkflowStateStartRequest) HasQueryAttributes() bool {
	if o != nil && o.QueryAttributes != nil {
		return true
	}

	return false
}

// SetQueryAttributes gets a reference to the given []KeyValue and assigns it to the QueryAttributes field.
func (o *WorkflowStateStartRequest) SetQueryAttributes(v []KeyValue) {
	o.QueryAttributes = v
}

func (o WorkflowStateStartRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Context != nil {
		toSerialize["context"] = o.Context
	}
	if o.WorkflowType != nil {
		toSerialize["workflowType"] = o.WorkflowType
	}
	if o.WorkflowStateId != nil {
		toSerialize["workflowStateId"] = o.WorkflowStateId
	}
	if o.StateInput != nil {
		toSerialize["stateInput"] = o.StateInput
	}
	if o.SearchAttributes != nil {
		toSerialize["searchAttributes"] = o.SearchAttributes
	}
	if o.QueryAttributes != nil {
		toSerialize["queryAttributes"] = o.QueryAttributes
	}
	return json.Marshal(toSerialize)
}

type NullableWorkflowStateStartRequest struct {
	value *WorkflowStateStartRequest
	isSet bool
}

func (v NullableWorkflowStateStartRequest) Get() *WorkflowStateStartRequest {
	return v.value
}

func (v *NullableWorkflowStateStartRequest) Set(val *WorkflowStateStartRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowStateStartRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowStateStartRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowStateStartRequest(val *WorkflowStateStartRequest) *NullableWorkflowStateStartRequest {
	return &NullableWorkflowStateStartRequest{value: val, isSet: true}
}

func (v NullableWorkflowStateStartRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowStateStartRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

