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

// WaitForQueryAttributeChangeCommand struct for WaitForQueryAttributeChangeCommand
type WaitForQueryAttributeChangeCommand struct {
	CommandId string `json:"commandId"`
	AttributeKey string `json:"attributeKey"`
}

// NewWaitForQueryAttributeChangeCommand instantiates a new WaitForQueryAttributeChangeCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWaitForQueryAttributeChangeCommand(commandId string, attributeKey string) *WaitForQueryAttributeChangeCommand {
	this := WaitForQueryAttributeChangeCommand{}
	this.CommandId = commandId
	this.AttributeKey = attributeKey
	return &this
}

// NewWaitForQueryAttributeChangeCommandWithDefaults instantiates a new WaitForQueryAttributeChangeCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWaitForQueryAttributeChangeCommandWithDefaults() *WaitForQueryAttributeChangeCommand {
	this := WaitForQueryAttributeChangeCommand{}
	return &this
}

// GetCommandId returns the CommandId field value
func (o *WaitForQueryAttributeChangeCommand) GetCommandId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CommandId
}

// GetCommandIdOk returns a tuple with the CommandId field value
// and a boolean to check if the value has been set.
func (o *WaitForQueryAttributeChangeCommand) GetCommandIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CommandId, true
}

// SetCommandId sets field value
func (o *WaitForQueryAttributeChangeCommand) SetCommandId(v string) {
	o.CommandId = v
}

// GetAttributeKey returns the AttributeKey field value
func (o *WaitForQueryAttributeChangeCommand) GetAttributeKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AttributeKey
}

// GetAttributeKeyOk returns a tuple with the AttributeKey field value
// and a boolean to check if the value has been set.
func (o *WaitForQueryAttributeChangeCommand) GetAttributeKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AttributeKey, true
}

// SetAttributeKey sets field value
func (o *WaitForQueryAttributeChangeCommand) SetAttributeKey(v string) {
	o.AttributeKey = v
}

func (o WaitForQueryAttributeChangeCommand) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["commandId"] = o.CommandId
	}
	if true {
		toSerialize["attributeKey"] = o.AttributeKey
	}
	return json.Marshal(toSerialize)
}

type NullableWaitForQueryAttributeChangeCommand struct {
	value *WaitForQueryAttributeChangeCommand
	isSet bool
}

func (v NullableWaitForQueryAttributeChangeCommand) Get() *WaitForQueryAttributeChangeCommand {
	return v.value
}

func (v *NullableWaitForQueryAttributeChangeCommand) Set(val *WaitForQueryAttributeChangeCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableWaitForQueryAttributeChangeCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableWaitForQueryAttributeChangeCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWaitForQueryAttributeChangeCommand(val *WaitForQueryAttributeChangeCommand) *NullableWaitForQueryAttributeChangeCommand {
	return &NullableWaitForQueryAttributeChangeCommand{value: val, isSet: true}
}

func (v NullableWaitForQueryAttributeChangeCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWaitForQueryAttributeChangeCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


