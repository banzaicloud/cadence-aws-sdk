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

type Client interface {
	AddListenerCertificates(ctx workflow.Context, input *elbv2.AddListenerCertificatesInput) (*elbv2.AddListenerCertificatesOutput, error)
	AddListenerCertificatesAsync(ctx workflow.Context, input *elbv2.AddListenerCertificatesInput) *ELBV2AddListenerCertificatesFuture

	AddTags(ctx workflow.Context, input *elbv2.AddTagsInput) (*elbv2.AddTagsOutput, error)
	AddTagsAsync(ctx workflow.Context, input *elbv2.AddTagsInput) *ELBV2AddTagsFuture

	CreateListener(ctx workflow.Context, input *elbv2.CreateListenerInput) (*elbv2.CreateListenerOutput, error)
	CreateListenerAsync(ctx workflow.Context, input *elbv2.CreateListenerInput) *ELBV2CreateListenerFuture

	CreateLoadBalancer(ctx workflow.Context, input *elbv2.CreateLoadBalancerInput) (*elbv2.CreateLoadBalancerOutput, error)
	CreateLoadBalancerAsync(ctx workflow.Context, input *elbv2.CreateLoadBalancerInput) *ELBV2CreateLoadBalancerFuture

	CreateRule(ctx workflow.Context, input *elbv2.CreateRuleInput) (*elbv2.CreateRuleOutput, error)
	CreateRuleAsync(ctx workflow.Context, input *elbv2.CreateRuleInput) *ELBV2CreateRuleFuture

	CreateTargetGroup(ctx workflow.Context, input *elbv2.CreateTargetGroupInput) (*elbv2.CreateTargetGroupOutput, error)
	CreateTargetGroupAsync(ctx workflow.Context, input *elbv2.CreateTargetGroupInput) *ELBV2CreateTargetGroupFuture

	DeleteListener(ctx workflow.Context, input *elbv2.DeleteListenerInput) (*elbv2.DeleteListenerOutput, error)
	DeleteListenerAsync(ctx workflow.Context, input *elbv2.DeleteListenerInput) *ELBV2DeleteListenerFuture

	DeleteLoadBalancer(ctx workflow.Context, input *elbv2.DeleteLoadBalancerInput) (*elbv2.DeleteLoadBalancerOutput, error)
	DeleteLoadBalancerAsync(ctx workflow.Context, input *elbv2.DeleteLoadBalancerInput) *ELBV2DeleteLoadBalancerFuture

	DeleteRule(ctx workflow.Context, input *elbv2.DeleteRuleInput) (*elbv2.DeleteRuleOutput, error)
	DeleteRuleAsync(ctx workflow.Context, input *elbv2.DeleteRuleInput) *ELBV2DeleteRuleFuture

	DeleteTargetGroup(ctx workflow.Context, input *elbv2.DeleteTargetGroupInput) (*elbv2.DeleteTargetGroupOutput, error)
	DeleteTargetGroupAsync(ctx workflow.Context, input *elbv2.DeleteTargetGroupInput) *ELBV2DeleteTargetGroupFuture

	DeregisterTargets(ctx workflow.Context, input *elbv2.DeregisterTargetsInput) (*elbv2.DeregisterTargetsOutput, error)
	DeregisterTargetsAsync(ctx workflow.Context, input *elbv2.DeregisterTargetsInput) *ELBV2DeregisterTargetsFuture

	DescribeAccountLimits(ctx workflow.Context, input *elbv2.DescribeAccountLimitsInput) (*elbv2.DescribeAccountLimitsOutput, error)
	DescribeAccountLimitsAsync(ctx workflow.Context, input *elbv2.DescribeAccountLimitsInput) *ELBV2DescribeAccountLimitsFuture

	DescribeListenerCertificates(ctx workflow.Context, input *elbv2.DescribeListenerCertificatesInput) (*elbv2.DescribeListenerCertificatesOutput, error)
	DescribeListenerCertificatesAsync(ctx workflow.Context, input *elbv2.DescribeListenerCertificatesInput) *ELBV2DescribeListenerCertificatesFuture

	DescribeListeners(ctx workflow.Context, input *elbv2.DescribeListenersInput) (*elbv2.DescribeListenersOutput, error)
	DescribeListenersAsync(ctx workflow.Context, input *elbv2.DescribeListenersInput) *ELBV2DescribeListenersFuture

	DescribeLoadBalancerAttributes(ctx workflow.Context, input *elbv2.DescribeLoadBalancerAttributesInput) (*elbv2.DescribeLoadBalancerAttributesOutput, error)
	DescribeLoadBalancerAttributesAsync(ctx workflow.Context, input *elbv2.DescribeLoadBalancerAttributesInput) *ELBV2DescribeLoadBalancerAttributesFuture

	DescribeLoadBalancers(ctx workflow.Context, input *elbv2.DescribeLoadBalancersInput) (*elbv2.DescribeLoadBalancersOutput, error)
	DescribeLoadBalancersAsync(ctx workflow.Context, input *elbv2.DescribeLoadBalancersInput) *ELBV2DescribeLoadBalancersFuture

	DescribeRules(ctx workflow.Context, input *elbv2.DescribeRulesInput) (*elbv2.DescribeRulesOutput, error)
	DescribeRulesAsync(ctx workflow.Context, input *elbv2.DescribeRulesInput) *ELBV2DescribeRulesFuture

	DescribeSSLPolicies(ctx workflow.Context, input *elbv2.DescribeSSLPoliciesInput) (*elbv2.DescribeSSLPoliciesOutput, error)
	DescribeSSLPoliciesAsync(ctx workflow.Context, input *elbv2.DescribeSSLPoliciesInput) *ELBV2DescribeSSLPoliciesFuture

	DescribeTags(ctx workflow.Context, input *elbv2.DescribeTagsInput) (*elbv2.DescribeTagsOutput, error)
	DescribeTagsAsync(ctx workflow.Context, input *elbv2.DescribeTagsInput) *ELBV2DescribeTagsFuture

	DescribeTargetGroupAttributes(ctx workflow.Context, input *elbv2.DescribeTargetGroupAttributesInput) (*elbv2.DescribeTargetGroupAttributesOutput, error)
	DescribeTargetGroupAttributesAsync(ctx workflow.Context, input *elbv2.DescribeTargetGroupAttributesInput) *ELBV2DescribeTargetGroupAttributesFuture

	DescribeTargetGroups(ctx workflow.Context, input *elbv2.DescribeTargetGroupsInput) (*elbv2.DescribeTargetGroupsOutput, error)
	DescribeTargetGroupsAsync(ctx workflow.Context, input *elbv2.DescribeTargetGroupsInput) *ELBV2DescribeTargetGroupsFuture

	DescribeTargetHealth(ctx workflow.Context, input *elbv2.DescribeTargetHealthInput) (*elbv2.DescribeTargetHealthOutput, error)
	DescribeTargetHealthAsync(ctx workflow.Context, input *elbv2.DescribeTargetHealthInput) *ELBV2DescribeTargetHealthFuture

	ModifyListener(ctx workflow.Context, input *elbv2.ModifyListenerInput) (*elbv2.ModifyListenerOutput, error)
	ModifyListenerAsync(ctx workflow.Context, input *elbv2.ModifyListenerInput) *ELBV2ModifyListenerFuture

	ModifyLoadBalancerAttributes(ctx workflow.Context, input *elbv2.ModifyLoadBalancerAttributesInput) (*elbv2.ModifyLoadBalancerAttributesOutput, error)
	ModifyLoadBalancerAttributesAsync(ctx workflow.Context, input *elbv2.ModifyLoadBalancerAttributesInput) *ELBV2ModifyLoadBalancerAttributesFuture

	ModifyRule(ctx workflow.Context, input *elbv2.ModifyRuleInput) (*elbv2.ModifyRuleOutput, error)
	ModifyRuleAsync(ctx workflow.Context, input *elbv2.ModifyRuleInput) *ELBV2ModifyRuleFuture

	ModifyTargetGroup(ctx workflow.Context, input *elbv2.ModifyTargetGroupInput) (*elbv2.ModifyTargetGroupOutput, error)
	ModifyTargetGroupAsync(ctx workflow.Context, input *elbv2.ModifyTargetGroupInput) *ELBV2ModifyTargetGroupFuture

	ModifyTargetGroupAttributes(ctx workflow.Context, input *elbv2.ModifyTargetGroupAttributesInput) (*elbv2.ModifyTargetGroupAttributesOutput, error)
	ModifyTargetGroupAttributesAsync(ctx workflow.Context, input *elbv2.ModifyTargetGroupAttributesInput) *ELBV2ModifyTargetGroupAttributesFuture

	RegisterTargets(ctx workflow.Context, input *elbv2.RegisterTargetsInput) (*elbv2.RegisterTargetsOutput, error)
	RegisterTargetsAsync(ctx workflow.Context, input *elbv2.RegisterTargetsInput) *ELBV2RegisterTargetsFuture

	RemoveListenerCertificates(ctx workflow.Context, input *elbv2.RemoveListenerCertificatesInput) (*elbv2.RemoveListenerCertificatesOutput, error)
	RemoveListenerCertificatesAsync(ctx workflow.Context, input *elbv2.RemoveListenerCertificatesInput) *ELBV2RemoveListenerCertificatesFuture

	RemoveTags(ctx workflow.Context, input *elbv2.RemoveTagsInput) (*elbv2.RemoveTagsOutput, error)
	RemoveTagsAsync(ctx workflow.Context, input *elbv2.RemoveTagsInput) *ELBV2RemoveTagsFuture

	SetIpAddressType(ctx workflow.Context, input *elbv2.SetIpAddressTypeInput) (*elbv2.SetIpAddressTypeOutput, error)
	SetIpAddressTypeAsync(ctx workflow.Context, input *elbv2.SetIpAddressTypeInput) *ELBV2SetIpAddressTypeFuture

	SetRulePriorities(ctx workflow.Context, input *elbv2.SetRulePrioritiesInput) (*elbv2.SetRulePrioritiesOutput, error)
	SetRulePrioritiesAsync(ctx workflow.Context, input *elbv2.SetRulePrioritiesInput) *ELBV2SetRulePrioritiesFuture

	SetSecurityGroups(ctx workflow.Context, input *elbv2.SetSecurityGroupsInput) (*elbv2.SetSecurityGroupsOutput, error)
	SetSecurityGroupsAsync(ctx workflow.Context, input *elbv2.SetSecurityGroupsInput) *ELBV2SetSecurityGroupsFuture

	SetSubnets(ctx workflow.Context, input *elbv2.SetSubnetsInput) (*elbv2.SetSubnetsOutput, error)
	SetSubnetsAsync(ctx workflow.Context, input *elbv2.SetSubnetsInput) *ELBV2SetSubnetsFuture

	WaitUntilLoadBalancerAvailable(ctx workflow.Context, input *elbv2.DescribeLoadBalancersInput) error
	WaitUntilLoadBalancerAvailableAsync(ctx workflow.Context, input *elbv2.DescribeLoadBalancersInput) *clients.VoidFuture

	WaitUntilLoadBalancerExists(ctx workflow.Context, input *elbv2.DescribeLoadBalancersInput) error
	WaitUntilLoadBalancerExistsAsync(ctx workflow.Context, input *elbv2.DescribeLoadBalancersInput) *clients.VoidFuture

	WaitUntilLoadBalancersDeleted(ctx workflow.Context, input *elbv2.DescribeLoadBalancersInput) error
	WaitUntilLoadBalancersDeletedAsync(ctx workflow.Context, input *elbv2.DescribeLoadBalancersInput) *clients.VoidFuture

	WaitUntilTargetDeregistered(ctx workflow.Context, input *elbv2.DescribeTargetHealthInput) error
	WaitUntilTargetDeregisteredAsync(ctx workflow.Context, input *elbv2.DescribeTargetHealthInput) *clients.VoidFuture

	WaitUntilTargetInService(ctx workflow.Context, input *elbv2.DescribeTargetHealthInput) error
	WaitUntilTargetInServiceAsync(ctx workflow.Context, input *elbv2.DescribeTargetHealthInput) *clients.VoidFuture
}

func NewClient() Client {
	return &stub{}
}
