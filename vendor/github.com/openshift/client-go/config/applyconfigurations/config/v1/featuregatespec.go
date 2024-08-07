// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	configv1 "github.com/openshift/api/config/v1"
)

// FeatureGateSpecApplyConfiguration represents a declarative configuration of the FeatureGateSpec type for use
// with apply.
type FeatureGateSpecApplyConfiguration struct {
	FeatureGateSelectionApplyConfiguration `json:",inline"`
}

// FeatureGateSpecApplyConfiguration constructs a declarative configuration of the FeatureGateSpec type for use with
// apply.
func FeatureGateSpec() *FeatureGateSpecApplyConfiguration {
	return &FeatureGateSpecApplyConfiguration{}
}

// WithFeatureSet sets the FeatureSet field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the FeatureSet field is set to the value of the last call.
func (b *FeatureGateSpecApplyConfiguration) WithFeatureSet(value configv1.FeatureSet) *FeatureGateSpecApplyConfiguration {
	b.FeatureSet = &value
	return b
}

// WithCustomNoUpgrade sets the CustomNoUpgrade field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CustomNoUpgrade field is set to the value of the last call.
func (b *FeatureGateSpecApplyConfiguration) WithCustomNoUpgrade(value *CustomFeatureGatesApplyConfiguration) *FeatureGateSpecApplyConfiguration {
	b.CustomNoUpgrade = value
	return b
}
