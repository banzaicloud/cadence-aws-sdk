// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7

package networkfirewall

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/networkfirewall"
	"github.com/aws/aws-sdk-go/service/networkfirewall/networkfirewalliface"

	"github.com/banzaicloud/cadence-aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken

type _ request.Option

// SessionFactory returns an aws.Session based on the activity context.
type SessionFactory interface {
	Session(ctx context.Context) (*session.Session, error)
}

type Activities struct {
	client networkfirewalliface.NetworkFirewallAPI

	sessionFactory SessionFactory
}

func NewActivities(sess *session.Session, config ...*aws.Config) *Activities {
	client := networkfirewall.New(sess, config...)
	return &Activities{client: client}
}

func NewActivitiesWithSessionFactory(sessionFactory SessionFactory) *Activities {
	return &Activities{sessionFactory: sessionFactory}
}

func (a *Activities) getClient(ctx context.Context) (networkfirewalliface.NetworkFirewallAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}

	return networkfirewall.New(sess), nil
}

func (a *Activities) AssociateFirewallPolicy(ctx context.Context, input *networkfirewall.AssociateFirewallPolicyInput) (*networkfirewall.AssociateFirewallPolicyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.AssociateFirewallPolicyWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) AssociateSubnets(ctx context.Context, input *networkfirewall.AssociateSubnetsInput) (*networkfirewall.AssociateSubnetsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.AssociateSubnetsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateFirewall(ctx context.Context, input *networkfirewall.CreateFirewallInput) (*networkfirewall.CreateFirewallOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateFirewallWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateFirewallPolicy(ctx context.Context, input *networkfirewall.CreateFirewallPolicyInput) (*networkfirewall.CreateFirewallPolicyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateFirewallPolicyWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateRuleGroup(ctx context.Context, input *networkfirewall.CreateRuleGroupInput) (*networkfirewall.CreateRuleGroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateRuleGroupWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteFirewall(ctx context.Context, input *networkfirewall.DeleteFirewallInput) (*networkfirewall.DeleteFirewallOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteFirewallWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteFirewallPolicy(ctx context.Context, input *networkfirewall.DeleteFirewallPolicyInput) (*networkfirewall.DeleteFirewallPolicyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteFirewallPolicyWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteResourcePolicy(ctx context.Context, input *networkfirewall.DeleteResourcePolicyInput) (*networkfirewall.DeleteResourcePolicyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteResourcePolicyWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteRuleGroup(ctx context.Context, input *networkfirewall.DeleteRuleGroupInput) (*networkfirewall.DeleteRuleGroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteRuleGroupWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeFirewall(ctx context.Context, input *networkfirewall.DescribeFirewallInput) (*networkfirewall.DescribeFirewallOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeFirewallWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeFirewallPolicy(ctx context.Context, input *networkfirewall.DescribeFirewallPolicyInput) (*networkfirewall.DescribeFirewallPolicyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeFirewallPolicyWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeLoggingConfiguration(ctx context.Context, input *networkfirewall.DescribeLoggingConfigurationInput) (*networkfirewall.DescribeLoggingConfigurationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeLoggingConfigurationWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeResourcePolicy(ctx context.Context, input *networkfirewall.DescribeResourcePolicyInput) (*networkfirewall.DescribeResourcePolicyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeResourcePolicyWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeRuleGroup(ctx context.Context, input *networkfirewall.DescribeRuleGroupInput) (*networkfirewall.DescribeRuleGroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeRuleGroupWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DisassociateSubnets(ctx context.Context, input *networkfirewall.DisassociateSubnetsInput) (*networkfirewall.DisassociateSubnetsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DisassociateSubnetsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListFirewallPolicies(ctx context.Context, input *networkfirewall.ListFirewallPoliciesInput) (*networkfirewall.ListFirewallPoliciesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListFirewallPoliciesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListFirewalls(ctx context.Context, input *networkfirewall.ListFirewallsInput) (*networkfirewall.ListFirewallsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListFirewallsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListRuleGroups(ctx context.Context, input *networkfirewall.ListRuleGroupsInput) (*networkfirewall.ListRuleGroupsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListRuleGroupsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListTagsForResource(ctx context.Context, input *networkfirewall.ListTagsForResourceInput) (*networkfirewall.ListTagsForResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListTagsForResourceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) PutResourcePolicy(ctx context.Context, input *networkfirewall.PutResourcePolicyInput) (*networkfirewall.PutResourcePolicyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.PutResourcePolicyWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) TagResource(ctx context.Context, input *networkfirewall.TagResourceInput) (*networkfirewall.TagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.TagResourceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UntagResource(ctx context.Context, input *networkfirewall.UntagResourceInput) (*networkfirewall.UntagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UntagResourceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateFirewallDeleteProtection(ctx context.Context, input *networkfirewall.UpdateFirewallDeleteProtectionInput) (*networkfirewall.UpdateFirewallDeleteProtectionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateFirewallDeleteProtectionWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateFirewallDescription(ctx context.Context, input *networkfirewall.UpdateFirewallDescriptionInput) (*networkfirewall.UpdateFirewallDescriptionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateFirewallDescriptionWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateFirewallPolicy(ctx context.Context, input *networkfirewall.UpdateFirewallPolicyInput) (*networkfirewall.UpdateFirewallPolicyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateFirewallPolicyWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateFirewallPolicyChangeProtection(ctx context.Context, input *networkfirewall.UpdateFirewallPolicyChangeProtectionInput) (*networkfirewall.UpdateFirewallPolicyChangeProtectionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateFirewallPolicyChangeProtectionWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateLoggingConfiguration(ctx context.Context, input *networkfirewall.UpdateLoggingConfigurationInput) (*networkfirewall.UpdateLoggingConfigurationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateLoggingConfigurationWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateRuleGroup(ctx context.Context, input *networkfirewall.UpdateRuleGroupInput) (*networkfirewall.UpdateRuleGroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateRuleGroupWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateSubnetChangeProtection(ctx context.Context, input *networkfirewall.UpdateSubnetChangeProtectionInput) (*networkfirewall.UpdateSubnetChangeProtectionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateSubnetChangeProtectionWithContext(ctx, input)

	return output, internal.EncodeError(err)
}
