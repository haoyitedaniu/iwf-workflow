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

// WorkflowSearchResponse struct for WorkflowSearchResponse
type WorkflowSearchResponse struct {
	WorkflowExecutions []WorkflowSearchResponseEntry `json:"workflowExecutions,omitempty"`
	NextPageToken      *string                       `json:"nextPageToken,omitempty"`
}

// NewWorkflowSearchResponse instantiates a new WorkflowSearchResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowSearchResponse() *WorkflowSearchResponse {
	this := WorkflowSearchResponse{}
	return &this
}

// NewWorkflowSearchResponseWithDefaults instantiates a new WorkflowSearchResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowSearchResponseWithDefaults() *WorkflowSearchResponse {
	this := WorkflowSearchResponse{}
	return &this
}

// GetWorkflowExecutions returns the WorkflowExecutions field value if set, zero value otherwise.
func (o *WorkflowSearchResponse) GetWorkflowExecutions() []WorkflowSearchResponseEntry {
	if o == nil || isNil(o.WorkflowExecutions) {
		var ret []WorkflowSearchResponseEntry
		return ret
	}
	return o.WorkflowExecutions
}

// GetWorkflowExecutionsOk returns a tuple with the WorkflowExecutions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowSearchResponse) GetWorkflowExecutionsOk() ([]WorkflowSearchResponseEntry, bool) {
	if o == nil || isNil(o.WorkflowExecutions) {
		return nil, false
	}
	return o.WorkflowExecutions, true
}

// HasWorkflowExecutions returns a boolean if a field has been set.
func (o *WorkflowSearchResponse) HasWorkflowExecutions() bool {
	if o != nil && !isNil(o.WorkflowExecutions) {
		return true
	}

	return false
}

// SetWorkflowExecutions gets a reference to the given []WorkflowSearchResponseEntry and assigns it to the WorkflowExecutions field.
func (o *WorkflowSearchResponse) SetWorkflowExecutions(v []WorkflowSearchResponseEntry) {
	o.WorkflowExecutions = v
}

// GetNextPageToken returns the NextPageToken field value if set, zero value otherwise.
func (o *WorkflowSearchResponse) GetNextPageToken() string {
	if o == nil || isNil(o.NextPageToken) {
		var ret string
		return ret
	}
	return *o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowSearchResponse) GetNextPageTokenOk() (*string, bool) {
	if o == nil || isNil(o.NextPageToken) {
		return nil, false
	}
	return o.NextPageToken, true
}

// HasNextPageToken returns a boolean if a field has been set.
func (o *WorkflowSearchResponse) HasNextPageToken() bool {
	if o != nil && !isNil(o.NextPageToken) {
		return true
	}

	return false
}

// SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.
func (o *WorkflowSearchResponse) SetNextPageToken(v string) {
	o.NextPageToken = &v
}

func (o WorkflowSearchResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.WorkflowExecutions) {
		toSerialize["workflowExecutions"] = o.WorkflowExecutions
	}
	if !isNil(o.NextPageToken) {
		toSerialize["nextPageToken"] = o.NextPageToken
	}
	return json.Marshal(toSerialize)
}

type NullableWorkflowSearchResponse struct {
	value *WorkflowSearchResponse
	isSet bool
}

func (v NullableWorkflowSearchResponse) Get() *WorkflowSearchResponse {
	return v.value
}

func (v *NullableWorkflowSearchResponse) Set(val *WorkflowSearchResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowSearchResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowSearchResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowSearchResponse(val *WorkflowSearchResponse) *NullableWorkflowSearchResponse {
	return &NullableWorkflowSearchResponse{value: val, isSet: true}
}

func (v NullableWorkflowSearchResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowSearchResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
