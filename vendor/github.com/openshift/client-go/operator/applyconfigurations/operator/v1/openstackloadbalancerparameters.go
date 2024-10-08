// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// OpenStackLoadBalancerParametersApplyConfiguration represents a declarative configuration of the OpenStackLoadBalancerParameters type for use
// with apply.
type OpenStackLoadBalancerParametersApplyConfiguration struct {
	LoadBalancerIP *string `json:"loadBalancerIP,omitempty"`
}

// OpenStackLoadBalancerParametersApplyConfiguration constructs a declarative configuration of the OpenStackLoadBalancerParameters type for use with
// apply.
func OpenStackLoadBalancerParameters() *OpenStackLoadBalancerParametersApplyConfiguration {
	return &OpenStackLoadBalancerParametersApplyConfiguration{}
}

// WithLoadBalancerIP sets the LoadBalancerIP field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the LoadBalancerIP field is set to the value of the last call.
func (b *OpenStackLoadBalancerParametersApplyConfiguration) WithLoadBalancerIP(value string) *OpenStackLoadBalancerParametersApplyConfiguration {
	b.LoadBalancerIP = &value
	return b
}
