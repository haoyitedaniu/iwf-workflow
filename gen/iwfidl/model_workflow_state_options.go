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

// WorkflowStateOptions struct for WorkflowStateOptions
type WorkflowStateOptions struct {
	SearchAttributesLoadingPolicy *PersistenceLoadingPolicy `json:"searchAttributesLoadingPolicy,omitempty"`
	DataObjectsLoadingPolicy      *PersistenceLoadingPolicy `json:"dataObjectsLoadingPolicy,omitempty"`
	CommandCarryOverPolicy        *CommandCarryOverPolicy   `json:"commandCarryOverPolicy,omitempty"`
	StartApiTimeoutSeconds        *int32                    `json:"startApiTimeoutSeconds,omitempty"`
	DecideApiTimeoutSeconds       *int32                    `json:"decideApiTimeoutSeconds,omitempty"`
	StartApiRetryPolicy           *RetryPolicy              `json:"startApiRetryPolicy,omitempty"`
	DecideApiRetryPolicy          *RetryPolicy              `json:"decideApiRetryPolicy,omitempty"`
}

// NewWorkflowStateOptions instantiates a new WorkflowStateOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowStateOptions() *WorkflowStateOptions {
	this := WorkflowStateOptions{}
	return &this
}

// NewWorkflowStateOptionsWithDefaults instantiates a new WorkflowStateOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowStateOptionsWithDefaults() *WorkflowStateOptions {
	this := WorkflowStateOptions{}
	return &this
}

// GetSearchAttributesLoadingPolicy returns the SearchAttributesLoadingPolicy field value if set, zero value otherwise.
func (o *WorkflowStateOptions) GetSearchAttributesLoadingPolicy() PersistenceLoadingPolicy {
	if o == nil || isNil(o.SearchAttributesLoadingPolicy) {
		var ret PersistenceLoadingPolicy
		return ret
	}
	return *o.SearchAttributesLoadingPolicy
}

// GetSearchAttributesLoadingPolicyOk returns a tuple with the SearchAttributesLoadingPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowStateOptions) GetSearchAttributesLoadingPolicyOk() (*PersistenceLoadingPolicy, bool) {
	if o == nil || isNil(o.SearchAttributesLoadingPolicy) {
		return nil, false
	}
	return o.SearchAttributesLoadingPolicy, true
}

// HasSearchAttributesLoadingPolicy returns a boolean if a field has been set.
func (o *WorkflowStateOptions) HasSearchAttributesLoadingPolicy() bool {
	if o != nil && !isNil(o.SearchAttributesLoadingPolicy) {
		return true
	}

	return false
}

// SetSearchAttributesLoadingPolicy gets a reference to the given PersistenceLoadingPolicy and assigns it to the SearchAttributesLoadingPolicy field.
func (o *WorkflowStateOptions) SetSearchAttributesLoadingPolicy(v PersistenceLoadingPolicy) {
	o.SearchAttributesLoadingPolicy = &v
}

// GetDataObjectsLoadingPolicy returns the DataObjectsLoadingPolicy field value if set, zero value otherwise.
func (o *WorkflowStateOptions) GetDataObjectsLoadingPolicy() PersistenceLoadingPolicy {
	if o == nil || isNil(o.DataObjectsLoadingPolicy) {
		var ret PersistenceLoadingPolicy
		return ret
	}
	return *o.DataObjectsLoadingPolicy
}

// GetDataObjectsLoadingPolicyOk returns a tuple with the DataObjectsLoadingPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowStateOptions) GetDataObjectsLoadingPolicyOk() (*PersistenceLoadingPolicy, bool) {
	if o == nil || isNil(o.DataObjectsLoadingPolicy) {
		return nil, false
	}
	return o.DataObjectsLoadingPolicy, true
}

// HasDataObjectsLoadingPolicy returns a boolean if a field has been set.
func (o *WorkflowStateOptions) HasDataObjectsLoadingPolicy() bool {
	if o != nil && !isNil(o.DataObjectsLoadingPolicy) {
		return true
	}

	return false
}

// SetDataObjectsLoadingPolicy gets a reference to the given PersistenceLoadingPolicy and assigns it to the DataObjectsLoadingPolicy field.
func (o *WorkflowStateOptions) SetDataObjectsLoadingPolicy(v PersistenceLoadingPolicy) {
	o.DataObjectsLoadingPolicy = &v
}

// GetCommandCarryOverPolicy returns the CommandCarryOverPolicy field value if set, zero value otherwise.
func (o *WorkflowStateOptions) GetCommandCarryOverPolicy() CommandCarryOverPolicy {
	if o == nil || isNil(o.CommandCarryOverPolicy) {
		var ret CommandCarryOverPolicy
		return ret
	}
	return *o.CommandCarryOverPolicy
}

// GetCommandCarryOverPolicyOk returns a tuple with the CommandCarryOverPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowStateOptions) GetCommandCarryOverPolicyOk() (*CommandCarryOverPolicy, bool) {
	if o == nil || isNil(o.CommandCarryOverPolicy) {
		return nil, false
	}
	return o.CommandCarryOverPolicy, true
}

// HasCommandCarryOverPolicy returns a boolean if a field has been set.
func (o *WorkflowStateOptions) HasCommandCarryOverPolicy() bool {
	if o != nil && !isNil(o.CommandCarryOverPolicy) {
		return true
	}

	return false
}

// SetCommandCarryOverPolicy gets a reference to the given CommandCarryOverPolicy and assigns it to the CommandCarryOverPolicy field.
func (o *WorkflowStateOptions) SetCommandCarryOverPolicy(v CommandCarryOverPolicy) {
	o.CommandCarryOverPolicy = &v
}

// GetStartApiTimeoutSeconds returns the StartApiTimeoutSeconds field value if set, zero value otherwise.
func (o *WorkflowStateOptions) GetStartApiTimeoutSeconds() int32 {
	if o == nil || isNil(o.StartApiTimeoutSeconds) {
		var ret int32
		return ret
	}
	return *o.StartApiTimeoutSeconds
}

// GetStartApiTimeoutSecondsOk returns a tuple with the StartApiTimeoutSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowStateOptions) GetStartApiTimeoutSecondsOk() (*int32, bool) {
	if o == nil || isNil(o.StartApiTimeoutSeconds) {
		return nil, false
	}
	return o.StartApiTimeoutSeconds, true
}

// HasStartApiTimeoutSeconds returns a boolean if a field has been set.
func (o *WorkflowStateOptions) HasStartApiTimeoutSeconds() bool {
	if o != nil && !isNil(o.StartApiTimeoutSeconds) {
		return true
	}

	return false
}

// SetStartApiTimeoutSeconds gets a reference to the given int32 and assigns it to the StartApiTimeoutSeconds field.
func (o *WorkflowStateOptions) SetStartApiTimeoutSeconds(v int32) {
	o.StartApiTimeoutSeconds = &v
}

// GetDecideApiTimeoutSeconds returns the DecideApiTimeoutSeconds field value if set, zero value otherwise.
func (o *WorkflowStateOptions) GetDecideApiTimeoutSeconds() int32 {
	if o == nil || isNil(o.DecideApiTimeoutSeconds) {
		var ret int32
		return ret
	}
	return *o.DecideApiTimeoutSeconds
}

// GetDecideApiTimeoutSecondsOk returns a tuple with the DecideApiTimeoutSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowStateOptions) GetDecideApiTimeoutSecondsOk() (*int32, bool) {
	if o == nil || isNil(o.DecideApiTimeoutSeconds) {
		return nil, false
	}
	return o.DecideApiTimeoutSeconds, true
}

// HasDecideApiTimeoutSeconds returns a boolean if a field has been set.
func (o *WorkflowStateOptions) HasDecideApiTimeoutSeconds() bool {
	if o != nil && !isNil(o.DecideApiTimeoutSeconds) {
		return true
	}

	return false
}

// SetDecideApiTimeoutSeconds gets a reference to the given int32 and assigns it to the DecideApiTimeoutSeconds field.
func (o *WorkflowStateOptions) SetDecideApiTimeoutSeconds(v int32) {
	o.DecideApiTimeoutSeconds = &v
}

// GetStartApiRetryPolicy returns the StartApiRetryPolicy field value if set, zero value otherwise.
func (o *WorkflowStateOptions) GetStartApiRetryPolicy() RetryPolicy {
	if o == nil || isNil(o.StartApiRetryPolicy) {
		var ret RetryPolicy
		return ret
	}
	return *o.StartApiRetryPolicy
}

// GetStartApiRetryPolicyOk returns a tuple with the StartApiRetryPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowStateOptions) GetStartApiRetryPolicyOk() (*RetryPolicy, bool) {
	if o == nil || isNil(o.StartApiRetryPolicy) {
		return nil, false
	}
	return o.StartApiRetryPolicy, true
}

// HasStartApiRetryPolicy returns a boolean if a field has been set.
func (o *WorkflowStateOptions) HasStartApiRetryPolicy() bool {
	if o != nil && !isNil(o.StartApiRetryPolicy) {
		return true
	}

	return false
}

// SetStartApiRetryPolicy gets a reference to the given RetryPolicy and assigns it to the StartApiRetryPolicy field.
func (o *WorkflowStateOptions) SetStartApiRetryPolicy(v RetryPolicy) {
	o.StartApiRetryPolicy = &v
}

// GetDecideApiRetryPolicy returns the DecideApiRetryPolicy field value if set, zero value otherwise.
func (o *WorkflowStateOptions) GetDecideApiRetryPolicy() RetryPolicy {
	if o == nil || isNil(o.DecideApiRetryPolicy) {
		var ret RetryPolicy
		return ret
	}
	return *o.DecideApiRetryPolicy
}

// GetDecideApiRetryPolicyOk returns a tuple with the DecideApiRetryPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowStateOptions) GetDecideApiRetryPolicyOk() (*RetryPolicy, bool) {
	if o == nil || isNil(o.DecideApiRetryPolicy) {
		return nil, false
	}
	return o.DecideApiRetryPolicy, true
}

// HasDecideApiRetryPolicy returns a boolean if a field has been set.
func (o *WorkflowStateOptions) HasDecideApiRetryPolicy() bool {
	if o != nil && !isNil(o.DecideApiRetryPolicy) {
		return true
	}

	return false
}

// SetDecideApiRetryPolicy gets a reference to the given RetryPolicy and assigns it to the DecideApiRetryPolicy field.
func (o *WorkflowStateOptions) SetDecideApiRetryPolicy(v RetryPolicy) {
	o.DecideApiRetryPolicy = &v
}

func (o WorkflowStateOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.SearchAttributesLoadingPolicy) {
		toSerialize["searchAttributesLoadingPolicy"] = o.SearchAttributesLoadingPolicy
	}
	if !isNil(o.DataObjectsLoadingPolicy) {
		toSerialize["dataObjectsLoadingPolicy"] = o.DataObjectsLoadingPolicy
	}
	if !isNil(o.CommandCarryOverPolicy) {
		toSerialize["commandCarryOverPolicy"] = o.CommandCarryOverPolicy
	}
	if !isNil(o.StartApiTimeoutSeconds) {
		toSerialize["startApiTimeoutSeconds"] = o.StartApiTimeoutSeconds
	}
	if !isNil(o.DecideApiTimeoutSeconds) {
		toSerialize["decideApiTimeoutSeconds"] = o.DecideApiTimeoutSeconds
	}
	if !isNil(o.StartApiRetryPolicy) {
		toSerialize["startApiRetryPolicy"] = o.StartApiRetryPolicy
	}
	if !isNil(o.DecideApiRetryPolicy) {
		toSerialize["decideApiRetryPolicy"] = o.DecideApiRetryPolicy
	}
	return json.Marshal(toSerialize)
}

type NullableWorkflowStateOptions struct {
	value *WorkflowStateOptions
	isSet bool
}

func (v NullableWorkflowStateOptions) Get() *WorkflowStateOptions {
	return v.value
}

func (v *NullableWorkflowStateOptions) Set(val *WorkflowStateOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowStateOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowStateOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowStateOptions(val *WorkflowStateOptions) *NullableWorkflowStateOptions {
	return &NullableWorkflowStateOptions{value: val, isSet: true}
}

func (v NullableWorkflowStateOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowStateOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
