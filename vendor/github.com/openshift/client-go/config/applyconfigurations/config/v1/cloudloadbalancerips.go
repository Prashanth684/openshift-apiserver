// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/openshift/api/config/v1"
)

// CloudLoadBalancerIPsApplyConfiguration represents an declarative configuration of the CloudLoadBalancerIPs type for use
// with apply.
type CloudLoadBalancerIPsApplyConfiguration struct {
	APIIntLoadBalancerIPs  []v1.IP `json:"apiIntLoadBalancerIPs,omitempty"`
	APILoadBalancerIPs     []v1.IP `json:"apiLoadBalancerIPs,omitempty"`
	IngressLoadBalancerIPs []v1.IP `json:"ingressLoadBalancerIPs,omitempty"`
}

// CloudLoadBalancerIPsApplyConfiguration constructs an declarative configuration of the CloudLoadBalancerIPs type for use with
// apply.
func CloudLoadBalancerIPs() *CloudLoadBalancerIPsApplyConfiguration {
	return &CloudLoadBalancerIPsApplyConfiguration{}
}

// WithAPIIntLoadBalancerIPs adds the given value to the APIIntLoadBalancerIPs field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the APIIntLoadBalancerIPs field.
func (b *CloudLoadBalancerIPsApplyConfiguration) WithAPIIntLoadBalancerIPs(values ...v1.IP) *CloudLoadBalancerIPsApplyConfiguration {
	for i := range values {
		b.APIIntLoadBalancerIPs = append(b.APIIntLoadBalancerIPs, values[i])
	}
	return b
}

// WithAPILoadBalancerIPs adds the given value to the APILoadBalancerIPs field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the APILoadBalancerIPs field.
func (b *CloudLoadBalancerIPsApplyConfiguration) WithAPILoadBalancerIPs(values ...v1.IP) *CloudLoadBalancerIPsApplyConfiguration {
	for i := range values {
		b.APILoadBalancerIPs = append(b.APILoadBalancerIPs, values[i])
	}
	return b
}

// WithIngressLoadBalancerIPs adds the given value to the IngressLoadBalancerIPs field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the IngressLoadBalancerIPs field.
func (b *CloudLoadBalancerIPsApplyConfiguration) WithIngressLoadBalancerIPs(values ...v1.IP) *CloudLoadBalancerIPsApplyConfiguration {
	for i := range values {
		b.IngressLoadBalancerIPs = append(b.IngressLoadBalancerIPs, values[i])
	}
	return b
}
