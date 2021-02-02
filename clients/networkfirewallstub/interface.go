// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7

package networkfirewallstub

import (
	"github.com/aws/aws-sdk-go/service/networkfirewall"
	"go.uber.org/cadence/workflow"

	"github.com/banzaicloud/cadence-aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	AssociateFirewallPolicy(ctx workflow.Context, input *networkfirewall.AssociateFirewallPolicyInput) (*networkfirewall.AssociateFirewallPolicyOutput, error)
	AssociateFirewallPolicyAsync(ctx workflow.Context, input *networkfirewall.AssociateFirewallPolicyInput) *AssociateFirewallPolicyFuture

	AssociateSubnets(ctx workflow.Context, input *networkfirewall.AssociateSubnetsInput) (*networkfirewall.AssociateSubnetsOutput, error)
	AssociateSubnetsAsync(ctx workflow.Context, input *networkfirewall.AssociateSubnetsInput) *AssociateSubnetsFuture

	CreateFirewall(ctx workflow.Context, input *networkfirewall.CreateFirewallInput) (*networkfirewall.CreateFirewallOutput, error)
	CreateFirewallAsync(ctx workflow.Context, input *networkfirewall.CreateFirewallInput) *CreateFirewallFuture

	CreateFirewallPolicy(ctx workflow.Context, input *networkfirewall.CreateFirewallPolicyInput) (*networkfirewall.CreateFirewallPolicyOutput, error)
	CreateFirewallPolicyAsync(ctx workflow.Context, input *networkfirewall.CreateFirewallPolicyInput) *CreateFirewallPolicyFuture

	CreateRuleGroup(ctx workflow.Context, input *networkfirewall.CreateRuleGroupInput) (*networkfirewall.CreateRuleGroupOutput, error)
	CreateRuleGroupAsync(ctx workflow.Context, input *networkfirewall.CreateRuleGroupInput) *CreateRuleGroupFuture

	DeleteFirewall(ctx workflow.Context, input *networkfirewall.DeleteFirewallInput) (*networkfirewall.DeleteFirewallOutput, error)
	DeleteFirewallAsync(ctx workflow.Context, input *networkfirewall.DeleteFirewallInput) *DeleteFirewallFuture

	DeleteFirewallPolicy(ctx workflow.Context, input *networkfirewall.DeleteFirewallPolicyInput) (*networkfirewall.DeleteFirewallPolicyOutput, error)
	DeleteFirewallPolicyAsync(ctx workflow.Context, input *networkfirewall.DeleteFirewallPolicyInput) *DeleteFirewallPolicyFuture

	DeleteResourcePolicy(ctx workflow.Context, input *networkfirewall.DeleteResourcePolicyInput) (*networkfirewall.DeleteResourcePolicyOutput, error)
	DeleteResourcePolicyAsync(ctx workflow.Context, input *networkfirewall.DeleteResourcePolicyInput) *DeleteResourcePolicyFuture

	DeleteRuleGroup(ctx workflow.Context, input *networkfirewall.DeleteRuleGroupInput) (*networkfirewall.DeleteRuleGroupOutput, error)
	DeleteRuleGroupAsync(ctx workflow.Context, input *networkfirewall.DeleteRuleGroupInput) *DeleteRuleGroupFuture

	DescribeFirewall(ctx workflow.Context, input *networkfirewall.DescribeFirewallInput) (*networkfirewall.DescribeFirewallOutput, error)
	DescribeFirewallAsync(ctx workflow.Context, input *networkfirewall.DescribeFirewallInput) *DescribeFirewallFuture

	DescribeFirewallPolicy(ctx workflow.Context, input *networkfirewall.DescribeFirewallPolicyInput) (*networkfirewall.DescribeFirewallPolicyOutput, error)
	DescribeFirewallPolicyAsync(ctx workflow.Context, input *networkfirewall.DescribeFirewallPolicyInput) *DescribeFirewallPolicyFuture

	DescribeLoggingConfiguration(ctx workflow.Context, input *networkfirewall.DescribeLoggingConfigurationInput) (*networkfirewall.DescribeLoggingConfigurationOutput, error)
	DescribeLoggingConfigurationAsync(ctx workflow.Context, input *networkfirewall.DescribeLoggingConfigurationInput) *DescribeLoggingConfigurationFuture

	DescribeResourcePolicy(ctx workflow.Context, input *networkfirewall.DescribeResourcePolicyInput) (*networkfirewall.DescribeResourcePolicyOutput, error)
	DescribeResourcePolicyAsync(ctx workflow.Context, input *networkfirewall.DescribeResourcePolicyInput) *DescribeResourcePolicyFuture

	DescribeRuleGroup(ctx workflow.Context, input *networkfirewall.DescribeRuleGroupInput) (*networkfirewall.DescribeRuleGroupOutput, error)
	DescribeRuleGroupAsync(ctx workflow.Context, input *networkfirewall.DescribeRuleGroupInput) *DescribeRuleGroupFuture

	DisassociateSubnets(ctx workflow.Context, input *networkfirewall.DisassociateSubnetsInput) (*networkfirewall.DisassociateSubnetsOutput, error)
	DisassociateSubnetsAsync(ctx workflow.Context, input *networkfirewall.DisassociateSubnetsInput) *DisassociateSubnetsFuture

	ListFirewallPolicies(ctx workflow.Context, input *networkfirewall.ListFirewallPoliciesInput) (*networkfirewall.ListFirewallPoliciesOutput, error)
	ListFirewallPoliciesAsync(ctx workflow.Context, input *networkfirewall.ListFirewallPoliciesInput) *ListFirewallPoliciesFuture

	ListFirewalls(ctx workflow.Context, input *networkfirewall.ListFirewallsInput) (*networkfirewall.ListFirewallsOutput, error)
	ListFirewallsAsync(ctx workflow.Context, input *networkfirewall.ListFirewallsInput) *ListFirewallsFuture

	ListRuleGroups(ctx workflow.Context, input *networkfirewall.ListRuleGroupsInput) (*networkfirewall.ListRuleGroupsOutput, error)
	ListRuleGroupsAsync(ctx workflow.Context, input *networkfirewall.ListRuleGroupsInput) *ListRuleGroupsFuture

	ListTagsForResource(ctx workflow.Context, input *networkfirewall.ListTagsForResourceInput) (*networkfirewall.ListTagsForResourceOutput, error)
	ListTagsForResourceAsync(ctx workflow.Context, input *networkfirewall.ListTagsForResourceInput) *ListTagsForResourceFuture

	PutResourcePolicy(ctx workflow.Context, input *networkfirewall.PutResourcePolicyInput) (*networkfirewall.PutResourcePolicyOutput, error)
	PutResourcePolicyAsync(ctx workflow.Context, input *networkfirewall.PutResourcePolicyInput) *PutResourcePolicyFuture

	TagResource(ctx workflow.Context, input *networkfirewall.TagResourceInput) (*networkfirewall.TagResourceOutput, error)
	TagResourceAsync(ctx workflow.Context, input *networkfirewall.TagResourceInput) *TagResourceFuture

	UntagResource(ctx workflow.Context, input *networkfirewall.UntagResourceInput) (*networkfirewall.UntagResourceOutput, error)
	UntagResourceAsync(ctx workflow.Context, input *networkfirewall.UntagResourceInput) *UntagResourceFuture

	UpdateFirewallDeleteProtection(ctx workflow.Context, input *networkfirewall.UpdateFirewallDeleteProtectionInput) (*networkfirewall.UpdateFirewallDeleteProtectionOutput, error)
	UpdateFirewallDeleteProtectionAsync(ctx workflow.Context, input *networkfirewall.UpdateFirewallDeleteProtectionInput) *UpdateFirewallDeleteProtectionFuture

	UpdateFirewallDescription(ctx workflow.Context, input *networkfirewall.UpdateFirewallDescriptionInput) (*networkfirewall.UpdateFirewallDescriptionOutput, error)
	UpdateFirewallDescriptionAsync(ctx workflow.Context, input *networkfirewall.UpdateFirewallDescriptionInput) *UpdateFirewallDescriptionFuture

	UpdateFirewallPolicy(ctx workflow.Context, input *networkfirewall.UpdateFirewallPolicyInput) (*networkfirewall.UpdateFirewallPolicyOutput, error)
	UpdateFirewallPolicyAsync(ctx workflow.Context, input *networkfirewall.UpdateFirewallPolicyInput) *UpdateFirewallPolicyFuture

	UpdateFirewallPolicyChangeProtection(ctx workflow.Context, input *networkfirewall.UpdateFirewallPolicyChangeProtectionInput) (*networkfirewall.UpdateFirewallPolicyChangeProtectionOutput, error)
	UpdateFirewallPolicyChangeProtectionAsync(ctx workflow.Context, input *networkfirewall.UpdateFirewallPolicyChangeProtectionInput) *UpdateFirewallPolicyChangeProtectionFuture

	UpdateLoggingConfiguration(ctx workflow.Context, input *networkfirewall.UpdateLoggingConfigurationInput) (*networkfirewall.UpdateLoggingConfigurationOutput, error)
	UpdateLoggingConfigurationAsync(ctx workflow.Context, input *networkfirewall.UpdateLoggingConfigurationInput) *UpdateLoggingConfigurationFuture

	UpdateRuleGroup(ctx workflow.Context, input *networkfirewall.UpdateRuleGroupInput) (*networkfirewall.UpdateRuleGroupOutput, error)
	UpdateRuleGroupAsync(ctx workflow.Context, input *networkfirewall.UpdateRuleGroupInput) *UpdateRuleGroupFuture

	UpdateSubnetChangeProtection(ctx workflow.Context, input *networkfirewall.UpdateSubnetChangeProtectionInput) (*networkfirewall.UpdateSubnetChangeProtectionOutput, error)
	UpdateSubnetChangeProtectionAsync(ctx workflow.Context, input *networkfirewall.UpdateSubnetChangeProtectionInput) *UpdateSubnetChangeProtectionFuture
}

func NewClient() Client {
	return &stub{}
}
