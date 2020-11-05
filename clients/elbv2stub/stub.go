// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7

package elbv2stub

import (
	"github.com/aws/aws-sdk-go/service/elbv2"
	"go.uber.org/cadence/workflow"

	"github.com/banzaicloud/cadence-aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type stub struct{}

type ELBV2AddListenerCertificatesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ELBV2AddListenerCertificatesFuture) Get(ctx workflow.Context) (*elbv2.AddListenerCertificatesOutput, error) {
	var output elbv2.AddListenerCertificatesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ELBV2AddTagsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ELBV2AddTagsFuture) Get(ctx workflow.Context) (*elbv2.AddTagsOutput, error) {
	var output elbv2.AddTagsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ELBV2CreateListenerFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ELBV2CreateListenerFuture) Get(ctx workflow.Context) (*elbv2.CreateListenerOutput, error) {
	var output elbv2.CreateListenerOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ELBV2CreateLoadBalancerFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ELBV2CreateLoadBalancerFuture) Get(ctx workflow.Context) (*elbv2.CreateLoadBalancerOutput, error) {
	var output elbv2.CreateLoadBalancerOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ELBV2CreateRuleFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ELBV2CreateRuleFuture) Get(ctx workflow.Context) (*elbv2.CreateRuleOutput, error) {
	var output elbv2.CreateRuleOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ELBV2CreateTargetGroupFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ELBV2CreateTargetGroupFuture) Get(ctx workflow.Context) (*elbv2.CreateTargetGroupOutput, error) {
	var output elbv2.CreateTargetGroupOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ELBV2DeleteListenerFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ELBV2DeleteListenerFuture) Get(ctx workflow.Context) (*elbv2.DeleteListenerOutput, error) {
	var output elbv2.DeleteListenerOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ELBV2DeleteLoadBalancerFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ELBV2DeleteLoadBalancerFuture) Get(ctx workflow.Context) (*elbv2.DeleteLoadBalancerOutput, error) {
	var output elbv2.DeleteLoadBalancerOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ELBV2DeleteRuleFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ELBV2DeleteRuleFuture) Get(ctx workflow.Context) (*elbv2.DeleteRuleOutput, error) {
	var output elbv2.DeleteRuleOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ELBV2DeleteTargetGroupFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ELBV2DeleteTargetGroupFuture) Get(ctx workflow.Context) (*elbv2.DeleteTargetGroupOutput, error) {
	var output elbv2.DeleteTargetGroupOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ELBV2DeregisterTargetsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ELBV2DeregisterTargetsFuture) Get(ctx workflow.Context) (*elbv2.DeregisterTargetsOutput, error) {
	var output elbv2.DeregisterTargetsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ELBV2DescribeAccountLimitsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ELBV2DescribeAccountLimitsFuture) Get(ctx workflow.Context) (*elbv2.DescribeAccountLimitsOutput, error) {
	var output elbv2.DescribeAccountLimitsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ELBV2DescribeListenerCertificatesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ELBV2DescribeListenerCertificatesFuture) Get(ctx workflow.Context) (*elbv2.DescribeListenerCertificatesOutput, error) {
	var output elbv2.DescribeListenerCertificatesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ELBV2DescribeListenersFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ELBV2DescribeListenersFuture) Get(ctx workflow.Context) (*elbv2.DescribeListenersOutput, error) {
	var output elbv2.DescribeListenersOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ELBV2DescribeLoadBalancerAttributesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ELBV2DescribeLoadBalancerAttributesFuture) Get(ctx workflow.Context) (*elbv2.DescribeLoadBalancerAttributesOutput, error) {
	var output elbv2.DescribeLoadBalancerAttributesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ELBV2DescribeLoadBalancersFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ELBV2DescribeLoadBalancersFuture) Get(ctx workflow.Context) (*elbv2.DescribeLoadBalancersOutput, error) {
	var output elbv2.DescribeLoadBalancersOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ELBV2DescribeRulesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ELBV2DescribeRulesFuture) Get(ctx workflow.Context) (*elbv2.DescribeRulesOutput, error) {
	var output elbv2.DescribeRulesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ELBV2DescribeSSLPoliciesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ELBV2DescribeSSLPoliciesFuture) Get(ctx workflow.Context) (*elbv2.DescribeSSLPoliciesOutput, error) {
	var output elbv2.DescribeSSLPoliciesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ELBV2DescribeTagsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ELBV2DescribeTagsFuture) Get(ctx workflow.Context) (*elbv2.DescribeTagsOutput, error) {
	var output elbv2.DescribeTagsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ELBV2DescribeTargetGroupAttributesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ELBV2DescribeTargetGroupAttributesFuture) Get(ctx workflow.Context) (*elbv2.DescribeTargetGroupAttributesOutput, error) {
	var output elbv2.DescribeTargetGroupAttributesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ELBV2DescribeTargetGroupsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ELBV2DescribeTargetGroupsFuture) Get(ctx workflow.Context) (*elbv2.DescribeTargetGroupsOutput, error) {
	var output elbv2.DescribeTargetGroupsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ELBV2DescribeTargetHealthFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ELBV2DescribeTargetHealthFuture) Get(ctx workflow.Context) (*elbv2.DescribeTargetHealthOutput, error) {
	var output elbv2.DescribeTargetHealthOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ELBV2ModifyListenerFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ELBV2ModifyListenerFuture) Get(ctx workflow.Context) (*elbv2.ModifyListenerOutput, error) {
	var output elbv2.ModifyListenerOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ELBV2ModifyLoadBalancerAttributesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ELBV2ModifyLoadBalancerAttributesFuture) Get(ctx workflow.Context) (*elbv2.ModifyLoadBalancerAttributesOutput, error) {
	var output elbv2.ModifyLoadBalancerAttributesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ELBV2ModifyRuleFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ELBV2ModifyRuleFuture) Get(ctx workflow.Context) (*elbv2.ModifyRuleOutput, error) {
	var output elbv2.ModifyRuleOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ELBV2ModifyTargetGroupFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ELBV2ModifyTargetGroupFuture) Get(ctx workflow.Context) (*elbv2.ModifyTargetGroupOutput, error) {
	var output elbv2.ModifyTargetGroupOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ELBV2ModifyTargetGroupAttributesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ELBV2ModifyTargetGroupAttributesFuture) Get(ctx workflow.Context) (*elbv2.ModifyTargetGroupAttributesOutput, error) {
	var output elbv2.ModifyTargetGroupAttributesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ELBV2RegisterTargetsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ELBV2RegisterTargetsFuture) Get(ctx workflow.Context) (*elbv2.RegisterTargetsOutput, error) {
	var output elbv2.RegisterTargetsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ELBV2RemoveListenerCertificatesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ELBV2RemoveListenerCertificatesFuture) Get(ctx workflow.Context) (*elbv2.RemoveListenerCertificatesOutput, error) {
	var output elbv2.RemoveListenerCertificatesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ELBV2RemoveTagsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ELBV2RemoveTagsFuture) Get(ctx workflow.Context) (*elbv2.RemoveTagsOutput, error) {
	var output elbv2.RemoveTagsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ELBV2SetIpAddressTypeFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ELBV2SetIpAddressTypeFuture) Get(ctx workflow.Context) (*elbv2.SetIpAddressTypeOutput, error) {
	var output elbv2.SetIpAddressTypeOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ELBV2SetRulePrioritiesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ELBV2SetRulePrioritiesFuture) Get(ctx workflow.Context) (*elbv2.SetRulePrioritiesOutput, error) {
	var output elbv2.SetRulePrioritiesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ELBV2SetSecurityGroupsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ELBV2SetSecurityGroupsFuture) Get(ctx workflow.Context) (*elbv2.SetSecurityGroupsOutput, error) {
	var output elbv2.SetSecurityGroupsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ELBV2SetSubnetsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ELBV2SetSubnetsFuture) Get(ctx workflow.Context) (*elbv2.SetSubnetsOutput, error) {
	var output elbv2.SetSubnetsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *stub) AddListenerCertificates(ctx workflow.Context, input *elbv2.AddListenerCertificatesInput) (*elbv2.AddListenerCertificatesOutput, error) {
	var output elbv2.AddListenerCertificatesOutput
	err := workflow.ExecuteActivity(ctx, "aws-elbv2-AddListenerCertificates", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) AddListenerCertificatesAsync(ctx workflow.Context, input *elbv2.AddListenerCertificatesInput) *ELBV2AddListenerCertificatesFuture {
	future := workflow.ExecuteActivity(ctx, "aws-elbv2-AddListenerCertificates", input)
	return &ELBV2AddListenerCertificatesFuture{Future: future}
}

func (a *stub) AddTags(ctx workflow.Context, input *elbv2.AddTagsInput) (*elbv2.AddTagsOutput, error) {
	var output elbv2.AddTagsOutput
	err := workflow.ExecuteActivity(ctx, "aws-elbv2-AddTags", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) AddTagsAsync(ctx workflow.Context, input *elbv2.AddTagsInput) *ELBV2AddTagsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-elbv2-AddTags", input)
	return &ELBV2AddTagsFuture{Future: future}
}

func (a *stub) CreateListener(ctx workflow.Context, input *elbv2.CreateListenerInput) (*elbv2.CreateListenerOutput, error) {
	var output elbv2.CreateListenerOutput
	err := workflow.ExecuteActivity(ctx, "aws-elbv2-CreateListener", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateListenerAsync(ctx workflow.Context, input *elbv2.CreateListenerInput) *ELBV2CreateListenerFuture {
	future := workflow.ExecuteActivity(ctx, "aws-elbv2-CreateListener", input)
	return &ELBV2CreateListenerFuture{Future: future}
}

func (a *stub) CreateLoadBalancer(ctx workflow.Context, input *elbv2.CreateLoadBalancerInput) (*elbv2.CreateLoadBalancerOutput, error) {
	var output elbv2.CreateLoadBalancerOutput
	err := workflow.ExecuteActivity(ctx, "aws-elbv2-CreateLoadBalancer", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateLoadBalancerAsync(ctx workflow.Context, input *elbv2.CreateLoadBalancerInput) *ELBV2CreateLoadBalancerFuture {
	future := workflow.ExecuteActivity(ctx, "aws-elbv2-CreateLoadBalancer", input)
	return &ELBV2CreateLoadBalancerFuture{Future: future}
}

func (a *stub) CreateRule(ctx workflow.Context, input *elbv2.CreateRuleInput) (*elbv2.CreateRuleOutput, error) {
	var output elbv2.CreateRuleOutput
	err := workflow.ExecuteActivity(ctx, "aws-elbv2-CreateRule", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateRuleAsync(ctx workflow.Context, input *elbv2.CreateRuleInput) *ELBV2CreateRuleFuture {
	future := workflow.ExecuteActivity(ctx, "aws-elbv2-CreateRule", input)
	return &ELBV2CreateRuleFuture{Future: future}
}

func (a *stub) CreateTargetGroup(ctx workflow.Context, input *elbv2.CreateTargetGroupInput) (*elbv2.CreateTargetGroupOutput, error) {
	var output elbv2.CreateTargetGroupOutput
	err := workflow.ExecuteActivity(ctx, "aws-elbv2-CreateTargetGroup", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateTargetGroupAsync(ctx workflow.Context, input *elbv2.CreateTargetGroupInput) *ELBV2CreateTargetGroupFuture {
	future := workflow.ExecuteActivity(ctx, "aws-elbv2-CreateTargetGroup", input)
	return &ELBV2CreateTargetGroupFuture{Future: future}
}

func (a *stub) DeleteListener(ctx workflow.Context, input *elbv2.DeleteListenerInput) (*elbv2.DeleteListenerOutput, error) {
	var output elbv2.DeleteListenerOutput
	err := workflow.ExecuteActivity(ctx, "aws-elbv2-DeleteListener", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteListenerAsync(ctx workflow.Context, input *elbv2.DeleteListenerInput) *ELBV2DeleteListenerFuture {
	future := workflow.ExecuteActivity(ctx, "aws-elbv2-DeleteListener", input)
	return &ELBV2DeleteListenerFuture{Future: future}
}

func (a *stub) DeleteLoadBalancer(ctx workflow.Context, input *elbv2.DeleteLoadBalancerInput) (*elbv2.DeleteLoadBalancerOutput, error) {
	var output elbv2.DeleteLoadBalancerOutput
	err := workflow.ExecuteActivity(ctx, "aws-elbv2-DeleteLoadBalancer", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteLoadBalancerAsync(ctx workflow.Context, input *elbv2.DeleteLoadBalancerInput) *ELBV2DeleteLoadBalancerFuture {
	future := workflow.ExecuteActivity(ctx, "aws-elbv2-DeleteLoadBalancer", input)
	return &ELBV2DeleteLoadBalancerFuture{Future: future}
}

func (a *stub) DeleteRule(ctx workflow.Context, input *elbv2.DeleteRuleInput) (*elbv2.DeleteRuleOutput, error) {
	var output elbv2.DeleteRuleOutput
	err := workflow.ExecuteActivity(ctx, "aws-elbv2-DeleteRule", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteRuleAsync(ctx workflow.Context, input *elbv2.DeleteRuleInput) *ELBV2DeleteRuleFuture {
	future := workflow.ExecuteActivity(ctx, "aws-elbv2-DeleteRule", input)
	return &ELBV2DeleteRuleFuture{Future: future}
}

func (a *stub) DeleteTargetGroup(ctx workflow.Context, input *elbv2.DeleteTargetGroupInput) (*elbv2.DeleteTargetGroupOutput, error) {
	var output elbv2.DeleteTargetGroupOutput
	err := workflow.ExecuteActivity(ctx, "aws-elbv2-DeleteTargetGroup", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteTargetGroupAsync(ctx workflow.Context, input *elbv2.DeleteTargetGroupInput) *ELBV2DeleteTargetGroupFuture {
	future := workflow.ExecuteActivity(ctx, "aws-elbv2-DeleteTargetGroup", input)
	return &ELBV2DeleteTargetGroupFuture{Future: future}
}

func (a *stub) DeregisterTargets(ctx workflow.Context, input *elbv2.DeregisterTargetsInput) (*elbv2.DeregisterTargetsOutput, error) {
	var output elbv2.DeregisterTargetsOutput
	err := workflow.ExecuteActivity(ctx, "aws-elbv2-DeregisterTargets", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeregisterTargetsAsync(ctx workflow.Context, input *elbv2.DeregisterTargetsInput) *ELBV2DeregisterTargetsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-elbv2-DeregisterTargets", input)
	return &ELBV2DeregisterTargetsFuture{Future: future}
}

func (a *stub) DescribeAccountLimits(ctx workflow.Context, input *elbv2.DescribeAccountLimitsInput) (*elbv2.DescribeAccountLimitsOutput, error) {
	var output elbv2.DescribeAccountLimitsOutput
	err := workflow.ExecuteActivity(ctx, "aws-elbv2-DescribeAccountLimits", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeAccountLimitsAsync(ctx workflow.Context, input *elbv2.DescribeAccountLimitsInput) *ELBV2DescribeAccountLimitsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-elbv2-DescribeAccountLimits", input)
	return &ELBV2DescribeAccountLimitsFuture{Future: future}
}

func (a *stub) DescribeListenerCertificates(ctx workflow.Context, input *elbv2.DescribeListenerCertificatesInput) (*elbv2.DescribeListenerCertificatesOutput, error) {
	var output elbv2.DescribeListenerCertificatesOutput
	err := workflow.ExecuteActivity(ctx, "aws-elbv2-DescribeListenerCertificates", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeListenerCertificatesAsync(ctx workflow.Context, input *elbv2.DescribeListenerCertificatesInput) *ELBV2DescribeListenerCertificatesFuture {
	future := workflow.ExecuteActivity(ctx, "aws-elbv2-DescribeListenerCertificates", input)
	return &ELBV2DescribeListenerCertificatesFuture{Future: future}
}

func (a *stub) DescribeListeners(ctx workflow.Context, input *elbv2.DescribeListenersInput) (*elbv2.DescribeListenersOutput, error) {
	var output elbv2.DescribeListenersOutput
	err := workflow.ExecuteActivity(ctx, "aws-elbv2-DescribeListeners", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeListenersAsync(ctx workflow.Context, input *elbv2.DescribeListenersInput) *ELBV2DescribeListenersFuture {
	future := workflow.ExecuteActivity(ctx, "aws-elbv2-DescribeListeners", input)
	return &ELBV2DescribeListenersFuture{Future: future}
}

func (a *stub) DescribeLoadBalancerAttributes(ctx workflow.Context, input *elbv2.DescribeLoadBalancerAttributesInput) (*elbv2.DescribeLoadBalancerAttributesOutput, error) {
	var output elbv2.DescribeLoadBalancerAttributesOutput
	err := workflow.ExecuteActivity(ctx, "aws-elbv2-DescribeLoadBalancerAttributes", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeLoadBalancerAttributesAsync(ctx workflow.Context, input *elbv2.DescribeLoadBalancerAttributesInput) *ELBV2DescribeLoadBalancerAttributesFuture {
	future := workflow.ExecuteActivity(ctx, "aws-elbv2-DescribeLoadBalancerAttributes", input)
	return &ELBV2DescribeLoadBalancerAttributesFuture{Future: future}
}

func (a *stub) DescribeLoadBalancers(ctx workflow.Context, input *elbv2.DescribeLoadBalancersInput) (*elbv2.DescribeLoadBalancersOutput, error) {
	var output elbv2.DescribeLoadBalancersOutput
	err := workflow.ExecuteActivity(ctx, "aws-elbv2-DescribeLoadBalancers", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeLoadBalancersAsync(ctx workflow.Context, input *elbv2.DescribeLoadBalancersInput) *ELBV2DescribeLoadBalancersFuture {
	future := workflow.ExecuteActivity(ctx, "aws-elbv2-DescribeLoadBalancers", input)
	return &ELBV2DescribeLoadBalancersFuture{Future: future}
}

func (a *stub) DescribeRules(ctx workflow.Context, input *elbv2.DescribeRulesInput) (*elbv2.DescribeRulesOutput, error) {
	var output elbv2.DescribeRulesOutput
	err := workflow.ExecuteActivity(ctx, "aws-elbv2-DescribeRules", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeRulesAsync(ctx workflow.Context, input *elbv2.DescribeRulesInput) *ELBV2DescribeRulesFuture {
	future := workflow.ExecuteActivity(ctx, "aws-elbv2-DescribeRules", input)
	return &ELBV2DescribeRulesFuture{Future: future}
}

func (a *stub) DescribeSSLPolicies(ctx workflow.Context, input *elbv2.DescribeSSLPoliciesInput) (*elbv2.DescribeSSLPoliciesOutput, error) {
	var output elbv2.DescribeSSLPoliciesOutput
	err := workflow.ExecuteActivity(ctx, "aws-elbv2-DescribeSSLPolicies", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeSSLPoliciesAsync(ctx workflow.Context, input *elbv2.DescribeSSLPoliciesInput) *ELBV2DescribeSSLPoliciesFuture {
	future := workflow.ExecuteActivity(ctx, "aws-elbv2-DescribeSSLPolicies", input)
	return &ELBV2DescribeSSLPoliciesFuture{Future: future}
}

func (a *stub) DescribeTags(ctx workflow.Context, input *elbv2.DescribeTagsInput) (*elbv2.DescribeTagsOutput, error) {
	var output elbv2.DescribeTagsOutput
	err := workflow.ExecuteActivity(ctx, "aws-elbv2-DescribeTags", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeTagsAsync(ctx workflow.Context, input *elbv2.DescribeTagsInput) *ELBV2DescribeTagsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-elbv2-DescribeTags", input)
	return &ELBV2DescribeTagsFuture{Future: future}
}

func (a *stub) DescribeTargetGroupAttributes(ctx workflow.Context, input *elbv2.DescribeTargetGroupAttributesInput) (*elbv2.DescribeTargetGroupAttributesOutput, error) {
	var output elbv2.DescribeTargetGroupAttributesOutput
	err := workflow.ExecuteActivity(ctx, "aws-elbv2-DescribeTargetGroupAttributes", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeTargetGroupAttributesAsync(ctx workflow.Context, input *elbv2.DescribeTargetGroupAttributesInput) *ELBV2DescribeTargetGroupAttributesFuture {
	future := workflow.ExecuteActivity(ctx, "aws-elbv2-DescribeTargetGroupAttributes", input)
	return &ELBV2DescribeTargetGroupAttributesFuture{Future: future}
}

func (a *stub) DescribeTargetGroups(ctx workflow.Context, input *elbv2.DescribeTargetGroupsInput) (*elbv2.DescribeTargetGroupsOutput, error) {
	var output elbv2.DescribeTargetGroupsOutput
	err := workflow.ExecuteActivity(ctx, "aws-elbv2-DescribeTargetGroups", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeTargetGroupsAsync(ctx workflow.Context, input *elbv2.DescribeTargetGroupsInput) *ELBV2DescribeTargetGroupsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-elbv2-DescribeTargetGroups", input)
	return &ELBV2DescribeTargetGroupsFuture{Future: future}
}

func (a *stub) DescribeTargetHealth(ctx workflow.Context, input *elbv2.DescribeTargetHealthInput) (*elbv2.DescribeTargetHealthOutput, error) {
	var output elbv2.DescribeTargetHealthOutput
	err := workflow.ExecuteActivity(ctx, "aws-elbv2-DescribeTargetHealth", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeTargetHealthAsync(ctx workflow.Context, input *elbv2.DescribeTargetHealthInput) *ELBV2DescribeTargetHealthFuture {
	future := workflow.ExecuteActivity(ctx, "aws-elbv2-DescribeTargetHealth", input)
	return &ELBV2DescribeTargetHealthFuture{Future: future}
}

func (a *stub) ModifyListener(ctx workflow.Context, input *elbv2.ModifyListenerInput) (*elbv2.ModifyListenerOutput, error) {
	var output elbv2.ModifyListenerOutput
	err := workflow.ExecuteActivity(ctx, "aws-elbv2-ModifyListener", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ModifyListenerAsync(ctx workflow.Context, input *elbv2.ModifyListenerInput) *ELBV2ModifyListenerFuture {
	future := workflow.ExecuteActivity(ctx, "aws-elbv2-ModifyListener", input)
	return &ELBV2ModifyListenerFuture{Future: future}
}

func (a *stub) ModifyLoadBalancerAttributes(ctx workflow.Context, input *elbv2.ModifyLoadBalancerAttributesInput) (*elbv2.ModifyLoadBalancerAttributesOutput, error) {
	var output elbv2.ModifyLoadBalancerAttributesOutput
	err := workflow.ExecuteActivity(ctx, "aws-elbv2-ModifyLoadBalancerAttributes", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ModifyLoadBalancerAttributesAsync(ctx workflow.Context, input *elbv2.ModifyLoadBalancerAttributesInput) *ELBV2ModifyLoadBalancerAttributesFuture {
	future := workflow.ExecuteActivity(ctx, "aws-elbv2-ModifyLoadBalancerAttributes", input)
	return &ELBV2ModifyLoadBalancerAttributesFuture{Future: future}
}

func (a *stub) ModifyRule(ctx workflow.Context, input *elbv2.ModifyRuleInput) (*elbv2.ModifyRuleOutput, error) {
	var output elbv2.ModifyRuleOutput
	err := workflow.ExecuteActivity(ctx, "aws-elbv2-ModifyRule", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ModifyRuleAsync(ctx workflow.Context, input *elbv2.ModifyRuleInput) *ELBV2ModifyRuleFuture {
	future := workflow.ExecuteActivity(ctx, "aws-elbv2-ModifyRule", input)
	return &ELBV2ModifyRuleFuture{Future: future}
}

func (a *stub) ModifyTargetGroup(ctx workflow.Context, input *elbv2.ModifyTargetGroupInput) (*elbv2.ModifyTargetGroupOutput, error) {
	var output elbv2.ModifyTargetGroupOutput
	err := workflow.ExecuteActivity(ctx, "aws-elbv2-ModifyTargetGroup", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ModifyTargetGroupAsync(ctx workflow.Context, input *elbv2.ModifyTargetGroupInput) *ELBV2ModifyTargetGroupFuture {
	future := workflow.ExecuteActivity(ctx, "aws-elbv2-ModifyTargetGroup", input)
	return &ELBV2ModifyTargetGroupFuture{Future: future}
}

func (a *stub) ModifyTargetGroupAttributes(ctx workflow.Context, input *elbv2.ModifyTargetGroupAttributesInput) (*elbv2.ModifyTargetGroupAttributesOutput, error) {
	var output elbv2.ModifyTargetGroupAttributesOutput
	err := workflow.ExecuteActivity(ctx, "aws-elbv2-ModifyTargetGroupAttributes", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ModifyTargetGroupAttributesAsync(ctx workflow.Context, input *elbv2.ModifyTargetGroupAttributesInput) *ELBV2ModifyTargetGroupAttributesFuture {
	future := workflow.ExecuteActivity(ctx, "aws-elbv2-ModifyTargetGroupAttributes", input)
	return &ELBV2ModifyTargetGroupAttributesFuture{Future: future}
}

func (a *stub) RegisterTargets(ctx workflow.Context, input *elbv2.RegisterTargetsInput) (*elbv2.RegisterTargetsOutput, error) {
	var output elbv2.RegisterTargetsOutput
	err := workflow.ExecuteActivity(ctx, "aws-elbv2-RegisterTargets", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) RegisterTargetsAsync(ctx workflow.Context, input *elbv2.RegisterTargetsInput) *ELBV2RegisterTargetsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-elbv2-RegisterTargets", input)
	return &ELBV2RegisterTargetsFuture{Future: future}
}

func (a *stub) RemoveListenerCertificates(ctx workflow.Context, input *elbv2.RemoveListenerCertificatesInput) (*elbv2.RemoveListenerCertificatesOutput, error) {
	var output elbv2.RemoveListenerCertificatesOutput
	err := workflow.ExecuteActivity(ctx, "aws-elbv2-RemoveListenerCertificates", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) RemoveListenerCertificatesAsync(ctx workflow.Context, input *elbv2.RemoveListenerCertificatesInput) *ELBV2RemoveListenerCertificatesFuture {
	future := workflow.ExecuteActivity(ctx, "aws-elbv2-RemoveListenerCertificates", input)
	return &ELBV2RemoveListenerCertificatesFuture{Future: future}
}

func (a *stub) RemoveTags(ctx workflow.Context, input *elbv2.RemoveTagsInput) (*elbv2.RemoveTagsOutput, error) {
	var output elbv2.RemoveTagsOutput
	err := workflow.ExecuteActivity(ctx, "aws-elbv2-RemoveTags", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) RemoveTagsAsync(ctx workflow.Context, input *elbv2.RemoveTagsInput) *ELBV2RemoveTagsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-elbv2-RemoveTags", input)
	return &ELBV2RemoveTagsFuture{Future: future}
}

func (a *stub) SetIpAddressType(ctx workflow.Context, input *elbv2.SetIpAddressTypeInput) (*elbv2.SetIpAddressTypeOutput, error) {
	var output elbv2.SetIpAddressTypeOutput
	err := workflow.ExecuteActivity(ctx, "aws-elbv2-SetIpAddressType", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) SetIpAddressTypeAsync(ctx workflow.Context, input *elbv2.SetIpAddressTypeInput) *ELBV2SetIpAddressTypeFuture {
	future := workflow.ExecuteActivity(ctx, "aws-elbv2-SetIpAddressType", input)
	return &ELBV2SetIpAddressTypeFuture{Future: future}
}

func (a *stub) SetRulePriorities(ctx workflow.Context, input *elbv2.SetRulePrioritiesInput) (*elbv2.SetRulePrioritiesOutput, error) {
	var output elbv2.SetRulePrioritiesOutput
	err := workflow.ExecuteActivity(ctx, "aws-elbv2-SetRulePriorities", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) SetRulePrioritiesAsync(ctx workflow.Context, input *elbv2.SetRulePrioritiesInput) *ELBV2SetRulePrioritiesFuture {
	future := workflow.ExecuteActivity(ctx, "aws-elbv2-SetRulePriorities", input)
	return &ELBV2SetRulePrioritiesFuture{Future: future}
}

func (a *stub) SetSecurityGroups(ctx workflow.Context, input *elbv2.SetSecurityGroupsInput) (*elbv2.SetSecurityGroupsOutput, error) {
	var output elbv2.SetSecurityGroupsOutput
	err := workflow.ExecuteActivity(ctx, "aws-elbv2-SetSecurityGroups", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) SetSecurityGroupsAsync(ctx workflow.Context, input *elbv2.SetSecurityGroupsInput) *ELBV2SetSecurityGroupsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-elbv2-SetSecurityGroups", input)
	return &ELBV2SetSecurityGroupsFuture{Future: future}
}

func (a *stub) SetSubnets(ctx workflow.Context, input *elbv2.SetSubnetsInput) (*elbv2.SetSubnetsOutput, error) {
	var output elbv2.SetSubnetsOutput
	err := workflow.ExecuteActivity(ctx, "aws-elbv2-SetSubnets", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) SetSubnetsAsync(ctx workflow.Context, input *elbv2.SetSubnetsInput) *ELBV2SetSubnetsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-elbv2-SetSubnets", input)
	return &ELBV2SetSubnetsFuture{Future: future}
}

func (a *stub) WaitUntilLoadBalancerAvailable(ctx workflow.Context, input *elbv2.DescribeLoadBalancersInput) error {
	return workflow.ExecuteActivity(ctx, "aws-elbv2-WaitUntilLoadBalancerAvailable", input).Get(ctx, nil)
}

func (a *stub) WaitUntilLoadBalancerAvailableAsync(ctx workflow.Context, input *elbv2.DescribeLoadBalancersInput) *clients.VoidFuture {
	future := workflow.ExecuteActivity(ctx, "aws-elbv2-WaitUntilLoadBalancerAvailable", input)
	return clients.NewVoidFuture(future)
}

func (a *stub) WaitUntilLoadBalancerExists(ctx workflow.Context, input *elbv2.DescribeLoadBalancersInput) error {
	return workflow.ExecuteActivity(ctx, "aws-elbv2-WaitUntilLoadBalancerExists", input).Get(ctx, nil)
}

func (a *stub) WaitUntilLoadBalancerExistsAsync(ctx workflow.Context, input *elbv2.DescribeLoadBalancersInput) *clients.VoidFuture {
	future := workflow.ExecuteActivity(ctx, "aws-elbv2-WaitUntilLoadBalancerExists", input)
	return clients.NewVoidFuture(future)
}

func (a *stub) WaitUntilLoadBalancersDeleted(ctx workflow.Context, input *elbv2.DescribeLoadBalancersInput) error {
	return workflow.ExecuteActivity(ctx, "aws-elbv2-WaitUntilLoadBalancersDeleted", input).Get(ctx, nil)
}

func (a *stub) WaitUntilLoadBalancersDeletedAsync(ctx workflow.Context, input *elbv2.DescribeLoadBalancersInput) *clients.VoidFuture {
	future := workflow.ExecuteActivity(ctx, "aws-elbv2-WaitUntilLoadBalancersDeleted", input)
	return clients.NewVoidFuture(future)
}

func (a *stub) WaitUntilTargetDeregistered(ctx workflow.Context, input *elbv2.DescribeTargetHealthInput) error {
	return workflow.ExecuteActivity(ctx, "aws-elbv2-WaitUntilTargetDeregistered", input).Get(ctx, nil)
}

func (a *stub) WaitUntilTargetDeregisteredAsync(ctx workflow.Context, input *elbv2.DescribeTargetHealthInput) *clients.VoidFuture {
	future := workflow.ExecuteActivity(ctx, "aws-elbv2-WaitUntilTargetDeregistered", input)
	return clients.NewVoidFuture(future)
}

func (a *stub) WaitUntilTargetInService(ctx workflow.Context, input *elbv2.DescribeTargetHealthInput) error {
	return workflow.ExecuteActivity(ctx, "aws-elbv2-WaitUntilTargetInService", input).Get(ctx, nil)
}

func (a *stub) WaitUntilTargetInServiceAsync(ctx workflow.Context, input *elbv2.DescribeTargetHealthInput) *clients.VoidFuture {
	future := workflow.ExecuteActivity(ctx, "aws-elbv2-WaitUntilTargetInService", input)
	return clients.NewVoidFuture(future)
}
