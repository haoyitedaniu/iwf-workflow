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

// StateCompletionOutput struct for StateCompletionOutput
type StateCompletionOutput struct {
	CompletedStateId          string         `json:"completedStateId"`
	CompletedStateExecutionId string         `json:"completedStateExecutionId"`
	CompletedStateOutput      *EncodedObject `json:"completedStateOutput,omitempty"`
}

// NewStateCompletionOutput instantiates a new StateCompletionOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStateCompletionOutput(completedStateId string, completedStateExecutionId string) *StateCompletionOutput {
	this := StateCompletionOutput{}
	this.CompletedStateId = completedStateId
	this.CompletedStateExecutionId = completedStateExecutionId
	return &this
}

// NewStateCompletionOutputWithDefaults instantiates a new StateCompletionOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStateCompletionOutputWithDefaults() *StateCompletionOutput {
	this := StateCompletionOutput{}
	return &this
}

// GetCompletedStateId returns the CompletedStateId field value
func (o *StateCompletionOutput) GetCompletedStateId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CompletedStateId
}

// GetCompletedStateIdOk returns a tuple with the CompletedStateId field value
// and a boolean to check if the value has been set.
func (o *StateCompletionOutput) GetCompletedStateIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CompletedStateId, true
}

// SetCompletedStateId sets field value
func (o *StateCompletionOutput) SetCompletedStateId(v string) {
	o.CompletedStateId = v
}

// GetCompletedStateExecutionId returns the CompletedStateExecutionId field value
func (o *StateCompletionOutput) GetCompletedStateExecutionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CompletedStateExecutionId
}

// GetCompletedStateExecutionIdOk returns a tuple with the CompletedStateExecutionId field value
// and a boolean to check if the value has been set.
func (o *StateCompletionOutput) GetCompletedStateExecutionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CompletedStateExecutionId, true
}

// SetCompletedStateExecutionId sets field value
func (o *StateCompletionOutput) SetCompletedStateExecutionId(v string) {
	o.CompletedStateExecutionId = v
}

// GetCompletedStateOutput returns the CompletedStateOutput field value if set, zero value otherwise.
func (o *StateCompletionOutput) GetCompletedStateOutput() EncodedObject {
	if o == nil || isNil(o.CompletedStateOutput) {
		var ret EncodedObject
		return ret
	}
	return *o.CompletedStateOutput
}

// GetCompletedStateOutputOk returns a tuple with the CompletedStateOutput field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StateCompletionOutput) GetCompletedStateOutputOk() (*EncodedObject, bool) {
	if o == nil || isNil(o.CompletedStateOutput) {
		return nil, false
	}
	return o.CompletedStateOutput, true
}

// HasCompletedStateOutput returns a boolean if a field has been set.
func (o *StateCompletionOutput) HasCompletedStateOutput() bool {
	if o != nil && !isNil(o.CompletedStateOutput) {
		return true
	}

	return false
}

// SetCompletedStateOutput gets a reference to the given EncodedObject and assigns it to the CompletedStateOutput field.
func (o *StateCompletionOutput) SetCompletedStateOutput(v EncodedObject) {
	o.CompletedStateOutput = &v
}

func (o StateCompletionOutput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["completedStateId"] = o.CompletedStateId
	}
	if true {
		toSerialize["completedStateExecutionId"] = o.CompletedStateExecutionId
	}
	if !isNil(o.CompletedStateOutput) {
		toSerialize["completedStateOutput"] = o.CompletedStateOutput
	}
	return json.Marshal(toSerialize)
}

type NullableStateCompletionOutput struct {
	value *StateCompletionOutput
	isSet bool
}

func (v NullableStateCompletionOutput) Get() *StateCompletionOutput {
	return v.value
}

func (v *NullableStateCompletionOutput) Set(val *StateCompletionOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableStateCompletionOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableStateCompletionOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStateCompletionOutput(val *StateCompletionOutput) *NullableStateCompletionOutput {
	return &NullableStateCompletionOutput{value: val, isSet: true}
}

func (v NullableStateCompletionOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStateCompletionOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
