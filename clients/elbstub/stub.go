// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7

package elbstub

import (
	"github.com/aws/aws-sdk-go/service/elb"
	"go.uber.org/cadence/workflow"

	"github.com/banzaicloud/cadence-aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type stub struct{}

type ELBAddTagsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ELBAddTagsFuture) Get(ctx workflow.Context) (*elb.AddTagsOutput, error) {
	var output elb.AddTagsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ELBApplySecurityGroupsToLoadBalancerFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ELBApplySecurityGroupsToLoadBalancerFuture) Get(ctx workflow.Context) (*elb.ApplySecurityGroupsToLoadBalancerOutput, error) {
	var output elb.ApplySecurityGroupsToLoadBalancerOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ELBAttachLoadBalancerToSubnetsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ELBAttachLoadBalancerToSubnetsFuture) Get(ctx workflow.Context) (*elb.AttachLoadBalancerToSubnetsOutput, error) {
	var output elb.AttachLoadBalancerToSubnetsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ELBConfigureHealthCheckFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ELBConfigureHealthCheckFuture) Get(ctx workflow.Context) (*elb.ConfigureHealthCheckOutput, error) {
	var output elb.ConfigureHealthCheckOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ELBCreateAppCookieStickinessPolicyFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ELBCreateAppCookieStickinessPolicyFuture) Get(ctx workflow.Context) (*elb.CreateAppCookieStickinessPolicyOutput, error) {
	var output elb.CreateAppCookieStickinessPolicyOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ELBCreateLBCookieStickinessPolicyFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ELBCreateLBCookieStickinessPolicyFuture) Get(ctx workflow.Context) (*elb.CreateLBCookieStickinessPolicyOutput, error) {
	var output elb.CreateLBCookieStickinessPolicyOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ELBCreateLoadBalancerFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ELBCreateLoadBalancerFuture) Get(ctx workflow.Context) (*elb.CreateLoadBalancerOutput, error) {
	var output elb.CreateLoadBalancerOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ELBCreateLoadBalancerListenersFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ELBCreateLoadBalancerListenersFuture) Get(ctx workflow.Context) (*elb.CreateLoadBalancerListenersOutput, error) {
	var output elb.CreateLoadBalancerListenersOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ELBCreateLoadBalancerPolicyFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ELBCreateLoadBalancerPolicyFuture) Get(ctx workflow.Context) (*elb.CreateLoadBalancerPolicyOutput, error) {
	var output elb.CreateLoadBalancerPolicyOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ELBDeleteLoadBalancerFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ELBDeleteLoadBalancerFuture) Get(ctx workflow.Context) (*elb.DeleteLoadBalancerOutput, error) {
	var output elb.DeleteLoadBalancerOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ELBDeleteLoadBalancerListenersFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ELBDeleteLoadBalancerListenersFuture) Get(ctx workflow.Context) (*elb.DeleteLoadBalancerListenersOutput, error) {
	var output elb.DeleteLoadBalancerListenersOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ELBDeleteLoadBalancerPolicyFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ELBDeleteLoadBalancerPolicyFuture) Get(ctx workflow.Context) (*elb.DeleteLoadBalancerPolicyOutput, error) {
	var output elb.DeleteLoadBalancerPolicyOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ELBDeregisterInstancesFromLoadBalancerFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ELBDeregisterInstancesFromLoadBalancerFuture) Get(ctx workflow.Context) (*elb.DeregisterInstancesFromLoadBalancerOutput, error) {
	var output elb.DeregisterInstancesFromLoadBalancerOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ELBDescribeAccountLimitsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ELBDescribeAccountLimitsFuture) Get(ctx workflow.Context) (*elb.DescribeAccountLimitsOutput, error) {
	var output elb.DescribeAccountLimitsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ELBDescribeInstanceHealthFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ELBDescribeInstanceHealthFuture) Get(ctx workflow.Context) (*elb.DescribeInstanceHealthOutput, error) {
	var output elb.DescribeInstanceHealthOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ELBDescribeLoadBalancerAttributesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ELBDescribeLoadBalancerAttributesFuture) Get(ctx workflow.Context) (*elb.DescribeLoadBalancerAttributesOutput, error) {
	var output elb.DescribeLoadBalancerAttributesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ELBDescribeLoadBalancerPoliciesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ELBDescribeLoadBalancerPoliciesFuture) Get(ctx workflow.Context) (*elb.DescribeLoadBalancerPoliciesOutput, error) {
	var output elb.DescribeLoadBalancerPoliciesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ELBDescribeLoadBalancerPolicyTypesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ELBDescribeLoadBalancerPolicyTypesFuture) Get(ctx workflow.Context) (*elb.DescribeLoadBalancerPolicyTypesOutput, error) {
	var output elb.DescribeLoadBalancerPolicyTypesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ELBDescribeLoadBalancersFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ELBDescribeLoadBalancersFuture) Get(ctx workflow.Context) (*elb.DescribeLoadBalancersOutput, error) {
	var output elb.DescribeLoadBalancersOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ELBDescribeTagsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ELBDescribeTagsFuture) Get(ctx workflow.Context) (*elb.DescribeTagsOutput, error) {
	var output elb.DescribeTagsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ELBDetachLoadBalancerFromSubnetsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ELBDetachLoadBalancerFromSubnetsFuture) Get(ctx workflow.Context) (*elb.DetachLoadBalancerFromSubnetsOutput, error) {
	var output elb.DetachLoadBalancerFromSubnetsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ELBDisableAvailabilityZonesForLoadBalancerFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ELBDisableAvailabilityZonesForLoadBalancerFuture) Get(ctx workflow.Context) (*elb.DisableAvailabilityZonesForLoadBalancerOutput, error) {
	var output elb.DisableAvailabilityZonesForLoadBalancerOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ELBEnableAvailabilityZonesForLoadBalancerFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ELBEnableAvailabilityZonesForLoadBalancerFuture) Get(ctx workflow.Context) (*elb.EnableAvailabilityZonesForLoadBalancerOutput, error) {
	var output elb.EnableAvailabilityZonesForLoadBalancerOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ELBModifyLoadBalancerAttributesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ELBModifyLoadBalancerAttributesFuture) Get(ctx workflow.Context) (*elb.ModifyLoadBalancerAttributesOutput, error) {
	var output elb.ModifyLoadBalancerAttributesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ELBRegisterInstancesWithLoadBalancerFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ELBRegisterInstancesWithLoadBalancerFuture) Get(ctx workflow.Context) (*elb.RegisterInstancesWithLoadBalancerOutput, error) {
	var output elb.RegisterInstancesWithLoadBalancerOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ELBRemoveTagsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ELBRemoveTagsFuture) Get(ctx workflow.Context) (*elb.RemoveTagsOutput, error) {
	var output elb.RemoveTagsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ELBSetLoadBalancerListenerSSLCertificateFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ELBSetLoadBalancerListenerSSLCertificateFuture) Get(ctx workflow.Context) (*elb.SetLoadBalancerListenerSSLCertificateOutput, error) {
	var output elb.SetLoadBalancerListenerSSLCertificateOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ELBSetLoadBalancerPoliciesForBackendServerFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ELBSetLoadBalancerPoliciesForBackendServerFuture) Get(ctx workflow.Context) (*elb.SetLoadBalancerPoliciesForBackendServerOutput, error) {
	var output elb.SetLoadBalancerPoliciesForBackendServerOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ELBSetLoadBalancerPoliciesOfListenerFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ELBSetLoadBalancerPoliciesOfListenerFuture) Get(ctx workflow.Context) (*elb.SetLoadBalancerPoliciesOfListenerOutput, error) {
	var output elb.SetLoadBalancerPoliciesOfListenerOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *stub) AddTags(ctx workflow.Context, input *elb.AddTagsInput) (*elb.AddTagsOutput, error) {
	var output elb.AddTagsOutput
	err := workflow.ExecuteActivity(ctx, "aws-elb-AddTags", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) AddTagsAsync(ctx workflow.Context, input *elb.AddTagsInput) *ELBAddTagsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-elb-AddTags", input)
	return &ELBAddTagsFuture{Future: future}
}

func (a *stub) ApplySecurityGroupsToLoadBalancer(ctx workflow.Context, input *elb.ApplySecurityGroupsToLoadBalancerInput) (*elb.ApplySecurityGroupsToLoadBalancerOutput, error) {
	var output elb.ApplySecurityGroupsToLoadBalancerOutput
	err := workflow.ExecuteActivity(ctx, "aws-elb-ApplySecurityGroupsToLoadBalancer", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ApplySecurityGroupsToLoadBalancerAsync(ctx workflow.Context, input *elb.ApplySecurityGroupsToLoadBalancerInput) *ELBApplySecurityGroupsToLoadBalancerFuture {
	future := workflow.ExecuteActivity(ctx, "aws-elb-ApplySecurityGroupsToLoadBalancer", input)
	return &ELBApplySecurityGroupsToLoadBalancerFuture{Future: future}
}

func (a *stub) AttachLoadBalancerToSubnets(ctx workflow.Context, input *elb.AttachLoadBalancerToSubnetsInput) (*elb.AttachLoadBalancerToSubnetsOutput, error) {
	var output elb.AttachLoadBalancerToSubnetsOutput
	err := workflow.ExecuteActivity(ctx, "aws-elb-AttachLoadBalancerToSubnets", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) AttachLoadBalancerToSubnetsAsync(ctx workflow.Context, input *elb.AttachLoadBalancerToSubnetsInput) *ELBAttachLoadBalancerToSubnetsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-elb-AttachLoadBalancerToSubnets", input)
	return &ELBAttachLoadBalancerToSubnetsFuture{Future: future}
}

func (a *stub) ConfigureHealthCheck(ctx workflow.Context, input *elb.ConfigureHealthCheckInput) (*elb.ConfigureHealthCheckOutput, error) {
	var output elb.ConfigureHealthCheckOutput
	err := workflow.ExecuteActivity(ctx, "aws-elb-ConfigureHealthCheck", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ConfigureHealthCheckAsync(ctx workflow.Context, input *elb.ConfigureHealthCheckInput) *ELBConfigureHealthCheckFuture {
	future := workflow.ExecuteActivity(ctx, "aws-elb-ConfigureHealthCheck", input)
	return &ELBConfigureHealthCheckFuture{Future: future}
}

func (a *stub) CreateAppCookieStickinessPolicy(ctx workflow.Context, input *elb.CreateAppCookieStickinessPolicyInput) (*elb.CreateAppCookieStickinessPolicyOutput, error) {
	var output elb.CreateAppCookieStickinessPolicyOutput
	err := workflow.ExecuteActivity(ctx, "aws-elb-CreateAppCookieStickinessPolicy", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateAppCookieStickinessPolicyAsync(ctx workflow.Context, input *elb.CreateAppCookieStickinessPolicyInput) *ELBCreateAppCookieStickinessPolicyFuture {
	future := workflow.ExecuteActivity(ctx, "aws-elb-CreateAppCookieStickinessPolicy", input)
	return &ELBCreateAppCookieStickinessPolicyFuture{Future: future}
}

func (a *stub) CreateLBCookieStickinessPolicy(ctx workflow.Context, input *elb.CreateLBCookieStickinessPolicyInput) (*elb.CreateLBCookieStickinessPolicyOutput, error) {
	var output elb.CreateLBCookieStickinessPolicyOutput
	err := workflow.ExecuteActivity(ctx, "aws-elb-CreateLBCookieStickinessPolicy", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateLBCookieStickinessPolicyAsync(ctx workflow.Context, input *elb.CreateLBCookieStickinessPolicyInput) *ELBCreateLBCookieStickinessPolicyFuture {
	future := workflow.ExecuteActivity(ctx, "aws-elb-CreateLBCookieStickinessPolicy", input)
	return &ELBCreateLBCookieStickinessPolicyFuture{Future: future}
}

func (a *stub) CreateLoadBalancer(ctx workflow.Context, input *elb.CreateLoadBalancerInput) (*elb.CreateLoadBalancerOutput, error) {
	var output elb.CreateLoadBalancerOutput
	err := workflow.ExecuteActivity(ctx, "aws-elb-CreateLoadBalancer", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateLoadBalancerAsync(ctx workflow.Context, input *elb.CreateLoadBalancerInput) *ELBCreateLoadBalancerFuture {
	future := workflow.ExecuteActivity(ctx, "aws-elb-CreateLoadBalancer", input)
	return &ELBCreateLoadBalancerFuture{Future: future}
}

func (a *stub) CreateLoadBalancerListeners(ctx workflow.Context, input *elb.CreateLoadBalancerListenersInput) (*elb.CreateLoadBalancerListenersOutput, error) {
	var output elb.CreateLoadBalancerListenersOutput
	err := workflow.ExecuteActivity(ctx, "aws-elb-CreateLoadBalancerListeners", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateLoadBalancerListenersAsync(ctx workflow.Context, input *elb.CreateLoadBalancerListenersInput) *ELBCreateLoadBalancerListenersFuture {
	future := workflow.ExecuteActivity(ctx, "aws-elb-CreateLoadBalancerListeners", input)
	return &ELBCreateLoadBalancerListenersFuture{Future: future}
}

func (a *stub) CreateLoadBalancerPolicy(ctx workflow.Context, input *elb.CreateLoadBalancerPolicyInput) (*elb.CreateLoadBalancerPolicyOutput, error) {
	var output elb.CreateLoadBalancerPolicyOutput
	err := workflow.ExecuteActivity(ctx, "aws-elb-CreateLoadBalancerPolicy", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateLoadBalancerPolicyAsync(ctx workflow.Context, input *elb.CreateLoadBalancerPolicyInput) *ELBCreateLoadBalancerPolicyFuture {
	future := workflow.ExecuteActivity(ctx, "aws-elb-CreateLoadBalancerPolicy", input)
	return &ELBCreateLoadBalancerPolicyFuture{Future: future}
}

func (a *stub) DeleteLoadBalancer(ctx workflow.Context, input *elb.DeleteLoadBalancerInput) (*elb.DeleteLoadBalancerOutput, error) {
	var output elb.DeleteLoadBalancerOutput
	err := workflow.ExecuteActivity(ctx, "aws-elb-DeleteLoadBalancer", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteLoadBalancerAsync(ctx workflow.Context, input *elb.DeleteLoadBalancerInput) *ELBDeleteLoadBalancerFuture {
	future := workflow.ExecuteActivity(ctx, "aws-elb-DeleteLoadBalancer", input)
	return &ELBDeleteLoadBalancerFuture{Future: future}
}

func (a *stub) DeleteLoadBalancerListeners(ctx workflow.Context, input *elb.DeleteLoadBalancerListenersInput) (*elb.DeleteLoadBalancerListenersOutput, error) {
	var output elb.DeleteLoadBalancerListenersOutput
	err := workflow.ExecuteActivity(ctx, "aws-elb-DeleteLoadBalancerListeners", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteLoadBalancerListenersAsync(ctx workflow.Context, input *elb.DeleteLoadBalancerListenersInput) *ELBDeleteLoadBalancerListenersFuture {
	future := workflow.ExecuteActivity(ctx, "aws-elb-DeleteLoadBalancerListeners", input)
	return &ELBDeleteLoadBalancerListenersFuture{Future: future}
}

func (a *stub) DeleteLoadBalancerPolicy(ctx workflow.Context, input *elb.DeleteLoadBalancerPolicyInput) (*elb.DeleteLoadBalancerPolicyOutput, error) {
	var output elb.DeleteLoadBalancerPolicyOutput
	err := workflow.ExecuteActivity(ctx, "aws-elb-DeleteLoadBalancerPolicy", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteLoadBalancerPolicyAsync(ctx workflow.Context, input *elb.DeleteLoadBalancerPolicyInput) *ELBDeleteLoadBalancerPolicyFuture {
	future := workflow.ExecuteActivity(ctx, "aws-elb-DeleteLoadBalancerPolicy", input)
	return &ELBDeleteLoadBalancerPolicyFuture{Future: future}
}

func (a *stub) DeregisterInstancesFromLoadBalancer(ctx workflow.Context, input *elb.DeregisterInstancesFromLoadBalancerInput) (*elb.DeregisterInstancesFromLoadBalancerOutput, error) {
	var output elb.DeregisterInstancesFromLoadBalancerOutput
	err := workflow.ExecuteActivity(ctx, "aws-elb-DeregisterInstancesFromLoadBalancer", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeregisterInstancesFromLoadBalancerAsync(ctx workflow.Context, input *elb.DeregisterInstancesFromLoadBalancerInput) *ELBDeregisterInstancesFromLoadBalancerFuture {
	future := workflow.ExecuteActivity(ctx, "aws-elb-DeregisterInstancesFromLoadBalancer", input)
	return &ELBDeregisterInstancesFromLoadBalancerFuture{Future: future}
}

func (a *stub) DescribeAccountLimits(ctx workflow.Context, input *elb.DescribeAccountLimitsInput) (*elb.DescribeAccountLimitsOutput, error) {
	var output elb.DescribeAccountLimitsOutput
	err := workflow.ExecuteActivity(ctx, "aws-elb-DescribeAccountLimits", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeAccountLimitsAsync(ctx workflow.Context, input *elb.DescribeAccountLimitsInput) *ELBDescribeAccountLimitsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-elb-DescribeAccountLimits", input)
	return &ELBDescribeAccountLimitsFuture{Future: future}
}

func (a *stub) DescribeInstanceHealth(ctx workflow.Context, input *elb.DescribeInstanceHealthInput) (*elb.DescribeInstanceHealthOutput, error) {
	var output elb.DescribeInstanceHealthOutput
	err := workflow.ExecuteActivity(ctx, "aws-elb-DescribeInstanceHealth", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeInstanceHealthAsync(ctx workflow.Context, input *elb.DescribeInstanceHealthInput) *ELBDescribeInstanceHealthFuture {
	future := workflow.ExecuteActivity(ctx, "aws-elb-DescribeInstanceHealth", input)
	return &ELBDescribeInstanceHealthFuture{Future: future}
}

func (a *stub) DescribeLoadBalancerAttributes(ctx workflow.Context, input *elb.DescribeLoadBalancerAttributesInput) (*elb.DescribeLoadBalancerAttributesOutput, error) {
	var output elb.DescribeLoadBalancerAttributesOutput
	err := workflow.ExecuteActivity(ctx, "aws-elb-DescribeLoadBalancerAttributes", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeLoadBalancerAttributesAsync(ctx workflow.Context, input *elb.DescribeLoadBalancerAttributesInput) *ELBDescribeLoadBalancerAttributesFuture {
	future := workflow.ExecuteActivity(ctx, "aws-elb-DescribeLoadBalancerAttributes", input)
	return &ELBDescribeLoadBalancerAttributesFuture{Future: future}
}

func (a *stub) DescribeLoadBalancerPolicies(ctx workflow.Context, input *elb.DescribeLoadBalancerPoliciesInput) (*elb.DescribeLoadBalancerPoliciesOutput, error) {
	var output elb.DescribeLoadBalancerPoliciesOutput
	err := workflow.ExecuteActivity(ctx, "aws-elb-DescribeLoadBalancerPolicies", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeLoadBalancerPoliciesAsync(ctx workflow.Context, input *elb.DescribeLoadBalancerPoliciesInput) *ELBDescribeLoadBalancerPoliciesFuture {
	future := workflow.ExecuteActivity(ctx, "aws-elb-DescribeLoadBalancerPolicies", input)
	return &ELBDescribeLoadBalancerPoliciesFuture{Future: future}
}

func (a *stub) DescribeLoadBalancerPolicyTypes(ctx workflow.Context, input *elb.DescribeLoadBalancerPolicyTypesInput) (*elb.DescribeLoadBalancerPolicyTypesOutput, error) {
	var output elb.DescribeLoadBalancerPolicyTypesOutput
	err := workflow.ExecuteActivity(ctx, "aws-elb-DescribeLoadBalancerPolicyTypes", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeLoadBalancerPolicyTypesAsync(ctx workflow.Context, input *elb.DescribeLoadBalancerPolicyTypesInput) *ELBDescribeLoadBalancerPolicyTypesFuture {
	future := workflow.ExecuteActivity(ctx, "aws-elb-DescribeLoadBalancerPolicyTypes", input)
	return &ELBDescribeLoadBalancerPolicyTypesFuture{Future: future}
}

func (a *stub) DescribeLoadBalancers(ctx workflow.Context, input *elb.DescribeLoadBalancersInput) (*elb.DescribeLoadBalancersOutput, error) {
	var output elb.DescribeLoadBalancersOutput
	err := workflow.ExecuteActivity(ctx, "aws-elb-DescribeLoadBalancers", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeLoadBalancersAsync(ctx workflow.Context, input *elb.DescribeLoadBalancersInput) *ELBDescribeLoadBalancersFuture {
	future := workflow.ExecuteActivity(ctx, "aws-elb-DescribeLoadBalancers", input)
	return &ELBDescribeLoadBalancersFuture{Future: future}
}

func (a *stub) DescribeTags(ctx workflow.Context, input *elb.DescribeTagsInput) (*elb.DescribeTagsOutput, error) {
	var output elb.DescribeTagsOutput
	err := workflow.ExecuteActivity(ctx, "aws-elb-DescribeTags", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeTagsAsync(ctx workflow.Context, input *elb.DescribeTagsInput) *ELBDescribeTagsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-elb-DescribeTags", input)
	return &ELBDescribeTagsFuture{Future: future}
}

func (a *stub) DetachLoadBalancerFromSubnets(ctx workflow.Context, input *elb.DetachLoadBalancerFromSubnetsInput) (*elb.DetachLoadBalancerFromSubnetsOutput, error) {
	var output elb.DetachLoadBalancerFromSubnetsOutput
	err := workflow.ExecuteActivity(ctx, "aws-elb-DetachLoadBalancerFromSubnets", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DetachLoadBalancerFromSubnetsAsync(ctx workflow.Context, input *elb.DetachLoadBalancerFromSubnetsInput) *ELBDetachLoadBalancerFromSubnetsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-elb-DetachLoadBalancerFromSubnets", input)
	return &ELBDetachLoadBalancerFromSubnetsFuture{Future: future}
}

func (a *stub) DisableAvailabilityZonesForLoadBalancer(ctx workflow.Context, input *elb.DisableAvailabilityZonesForLoadBalancerInput) (*elb.DisableAvailabilityZonesForLoadBalancerOutput, error) {
	var output elb.DisableAvailabilityZonesForLoadBalancerOutput
	err := workflow.ExecuteActivity(ctx, "aws-elb-DisableAvailabilityZonesForLoadBalancer", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DisableAvailabilityZonesForLoadBalancerAsync(ctx workflow.Context, input *elb.DisableAvailabilityZonesForLoadBalancerInput) *ELBDisableAvailabilityZonesForLoadBalancerFuture {
	future := workflow.ExecuteActivity(ctx, "aws-elb-DisableAvailabilityZonesForLoadBalancer", input)
	return &ELBDisableAvailabilityZonesForLoadBalancerFuture{Future: future}
}

func (a *stub) EnableAvailabilityZonesForLoadBalancer(ctx workflow.Context, input *elb.EnableAvailabilityZonesForLoadBalancerInput) (*elb.EnableAvailabilityZonesForLoadBalancerOutput, error) {
	var output elb.EnableAvailabilityZonesForLoadBalancerOutput
	err := workflow.ExecuteActivity(ctx, "aws-elb-EnableAvailabilityZonesForLoadBalancer", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) EnableAvailabilityZonesForLoadBalancerAsync(ctx workflow.Context, input *elb.EnableAvailabilityZonesForLoadBalancerInput) *ELBEnableAvailabilityZonesForLoadBalancerFuture {
	future := workflow.ExecuteActivity(ctx, "aws-elb-EnableAvailabilityZonesForLoadBalancer", input)
	return &ELBEnableAvailabilityZonesForLoadBalancerFuture{Future: future}
}

func (a *stub) ModifyLoadBalancerAttributes(ctx workflow.Context, input *elb.ModifyLoadBalancerAttributesInput) (*elb.ModifyLoadBalancerAttributesOutput, error) {
	var output elb.ModifyLoadBalancerAttributesOutput
	err := workflow.ExecuteActivity(ctx, "aws-elb-ModifyLoadBalancerAttributes", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ModifyLoadBalancerAttributesAsync(ctx workflow.Context, input *elb.ModifyLoadBalancerAttributesInput) *ELBModifyLoadBalancerAttributesFuture {
	future := workflow.ExecuteActivity(ctx, "aws-elb-ModifyLoadBalancerAttributes", input)
	return &ELBModifyLoadBalancerAttributesFuture{Future: future}
}

func (a *stub) RegisterInstancesWithLoadBalancer(ctx workflow.Context, input *elb.RegisterInstancesWithLoadBalancerInput) (*elb.RegisterInstancesWithLoadBalancerOutput, error) {
	var output elb.RegisterInstancesWithLoadBalancerOutput
	err := workflow.ExecuteActivity(ctx, "aws-elb-RegisterInstancesWithLoadBalancer", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) RegisterInstancesWithLoadBalancerAsync(ctx workflow.Context, input *elb.RegisterInstancesWithLoadBalancerInput) *ELBRegisterInstancesWithLoadBalancerFuture {
	future := workflow.ExecuteActivity(ctx, "aws-elb-RegisterInstancesWithLoadBalancer", input)
	return &ELBRegisterInstancesWithLoadBalancerFuture{Future: future}
}

func (a *stub) RemoveTags(ctx workflow.Context, input *elb.RemoveTagsInput) (*elb.RemoveTagsOutput, error) {
	var output elb.RemoveTagsOutput
	err := workflow.ExecuteActivity(ctx, "aws-elb-RemoveTags", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) RemoveTagsAsync(ctx workflow.Context, input *elb.RemoveTagsInput) *ELBRemoveTagsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-elb-RemoveTags", input)
	return &ELBRemoveTagsFuture{Future: future}
}

func (a *stub) SetLoadBalancerListenerSSLCertificate(ctx workflow.Context, input *elb.SetLoadBalancerListenerSSLCertificateInput) (*elb.SetLoadBalancerListenerSSLCertificateOutput, error) {
	var output elb.SetLoadBalancerListenerSSLCertificateOutput
	err := workflow.ExecuteActivity(ctx, "aws-elb-SetLoadBalancerListenerSSLCertificate", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) SetLoadBalancerListenerSSLCertificateAsync(ctx workflow.Context, input *elb.SetLoadBalancerListenerSSLCertificateInput) *ELBSetLoadBalancerListenerSSLCertificateFuture {
	future := workflow.ExecuteActivity(ctx, "aws-elb-SetLoadBalancerListenerSSLCertificate", input)
	return &ELBSetLoadBalancerListenerSSLCertificateFuture{Future: future}
}

func (a *stub) SetLoadBalancerPoliciesForBackendServer(ctx workflow.Context, input *elb.SetLoadBalancerPoliciesForBackendServerInput) (*elb.SetLoadBalancerPoliciesForBackendServerOutput, error) {
	var output elb.SetLoadBalancerPoliciesForBackendServerOutput
	err := workflow.ExecuteActivity(ctx, "aws-elb-SetLoadBalancerPoliciesForBackendServer", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) SetLoadBalancerPoliciesForBackendServerAsync(ctx workflow.Context, input *elb.SetLoadBalancerPoliciesForBackendServerInput) *ELBSetLoadBalancerPoliciesForBackendServerFuture {
	future := workflow.ExecuteActivity(ctx, "aws-elb-SetLoadBalancerPoliciesForBackendServer", input)
	return &ELBSetLoadBalancerPoliciesForBackendServerFuture{Future: future}
}

func (a *stub) SetLoadBalancerPoliciesOfListener(ctx workflow.Context, input *elb.SetLoadBalancerPoliciesOfListenerInput) (*elb.SetLoadBalancerPoliciesOfListenerOutput, error) {
	var output elb.SetLoadBalancerPoliciesOfListenerOutput
	err := workflow.ExecuteActivity(ctx, "aws-elb-SetLoadBalancerPoliciesOfListener", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) SetLoadBalancerPoliciesOfListenerAsync(ctx workflow.Context, input *elb.SetLoadBalancerPoliciesOfListenerInput) *ELBSetLoadBalancerPoliciesOfListenerFuture {
	future := workflow.ExecuteActivity(ctx, "aws-elb-SetLoadBalancerPoliciesOfListener", input)
	return &ELBSetLoadBalancerPoliciesOfListenerFuture{Future: future}
}

func (a *stub) WaitUntilAnyInstanceInService(ctx workflow.Context, input *elb.DescribeInstanceHealthInput) error {
	return workflow.ExecuteActivity(ctx, "aws-elb-WaitUntilAnyInstanceInService", input).Get(ctx, nil)
}

func (a *stub) WaitUntilAnyInstanceInServiceAsync(ctx workflow.Context, input *elb.DescribeInstanceHealthInput) *clients.VoidFuture {
	future := workflow.ExecuteActivity(ctx, "aws-elb-WaitUntilAnyInstanceInService", input)
	return clients.NewVoidFuture(future)
}

func (a *stub) WaitUntilInstanceDeregistered(ctx workflow.Context, input *elb.DescribeInstanceHealthInput) error {
	return workflow.ExecuteActivity(ctx, "aws-elb-WaitUntilInstanceDeregistered", input).Get(ctx, nil)
}

func (a *stub) WaitUntilInstanceDeregisteredAsync(ctx workflow.Context, input *elb.DescribeInstanceHealthInput) *clients.VoidFuture {
	future := workflow.ExecuteActivity(ctx, "aws-elb-WaitUntilInstanceDeregistered", input)
	return clients.NewVoidFuture(future)
}

func (a *stub) WaitUntilInstanceInService(ctx workflow.Context, input *elb.DescribeInstanceHealthInput) error {
	return workflow.ExecuteActivity(ctx, "aws-elb-WaitUntilInstanceInService", input).Get(ctx, nil)
}

func (a *stub) WaitUntilInstanceInServiceAsync(ctx workflow.Context, input *elb.DescribeInstanceHealthInput) *clients.VoidFuture {
	future := workflow.ExecuteActivity(ctx, "aws-elb-WaitUntilInstanceInService", input)
	return clients.NewVoidFuture(future)
}
