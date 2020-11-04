// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.
// Copyright (c) 2020 Banzai Cloud All rights reserved.

package cloudwatcheventsstub

import (
	"github.com/aws/aws-sdk-go/service/cloudwatchevents"
	"go.uber.org/cadence/workflow"

	"github.com/banzaicloud/cadence-aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	ActivateEventSource(ctx workflow.Context, input *cloudwatchevents.ActivateEventSourceInput) (*cloudwatchevents.ActivateEventSourceOutput, error)
	ActivateEventSourceAsync(ctx workflow.Context, input *cloudwatchevents.ActivateEventSourceInput) *CloudWatchEventsActivateEventSourceFuture

	CreateEventBus(ctx workflow.Context, input *cloudwatchevents.CreateEventBusInput) (*cloudwatchevents.CreateEventBusOutput, error)
	CreateEventBusAsync(ctx workflow.Context, input *cloudwatchevents.CreateEventBusInput) *CloudWatchEventsCreateEventBusFuture

	CreatePartnerEventSource(ctx workflow.Context, input *cloudwatchevents.CreatePartnerEventSourceInput) (*cloudwatchevents.CreatePartnerEventSourceOutput, error)
	CreatePartnerEventSourceAsync(ctx workflow.Context, input *cloudwatchevents.CreatePartnerEventSourceInput) *CloudWatchEventsCreatePartnerEventSourceFuture

	DeactivateEventSource(ctx workflow.Context, input *cloudwatchevents.DeactivateEventSourceInput) (*cloudwatchevents.DeactivateEventSourceOutput, error)
	DeactivateEventSourceAsync(ctx workflow.Context, input *cloudwatchevents.DeactivateEventSourceInput) *CloudWatchEventsDeactivateEventSourceFuture

	DeleteEventBus(ctx workflow.Context, input *cloudwatchevents.DeleteEventBusInput) (*cloudwatchevents.DeleteEventBusOutput, error)
	DeleteEventBusAsync(ctx workflow.Context, input *cloudwatchevents.DeleteEventBusInput) *CloudWatchEventsDeleteEventBusFuture

	DeletePartnerEventSource(ctx workflow.Context, input *cloudwatchevents.DeletePartnerEventSourceInput) (*cloudwatchevents.DeletePartnerEventSourceOutput, error)
	DeletePartnerEventSourceAsync(ctx workflow.Context, input *cloudwatchevents.DeletePartnerEventSourceInput) *CloudWatchEventsDeletePartnerEventSourceFuture

	DeleteRule(ctx workflow.Context, input *cloudwatchevents.DeleteRuleInput) (*cloudwatchevents.DeleteRuleOutput, error)
	DeleteRuleAsync(ctx workflow.Context, input *cloudwatchevents.DeleteRuleInput) *CloudWatchEventsDeleteRuleFuture

	DescribeEventBus(ctx workflow.Context, input *cloudwatchevents.DescribeEventBusInput) (*cloudwatchevents.DescribeEventBusOutput, error)
	DescribeEventBusAsync(ctx workflow.Context, input *cloudwatchevents.DescribeEventBusInput) *CloudWatchEventsDescribeEventBusFuture

	DescribeEventSource(ctx workflow.Context, input *cloudwatchevents.DescribeEventSourceInput) (*cloudwatchevents.DescribeEventSourceOutput, error)
	DescribeEventSourceAsync(ctx workflow.Context, input *cloudwatchevents.DescribeEventSourceInput) *CloudWatchEventsDescribeEventSourceFuture

	DescribePartnerEventSource(ctx workflow.Context, input *cloudwatchevents.DescribePartnerEventSourceInput) (*cloudwatchevents.DescribePartnerEventSourceOutput, error)
	DescribePartnerEventSourceAsync(ctx workflow.Context, input *cloudwatchevents.DescribePartnerEventSourceInput) *CloudWatchEventsDescribePartnerEventSourceFuture

	DescribeRule(ctx workflow.Context, input *cloudwatchevents.DescribeRuleInput) (*cloudwatchevents.DescribeRuleOutput, error)
	DescribeRuleAsync(ctx workflow.Context, input *cloudwatchevents.DescribeRuleInput) *CloudWatchEventsDescribeRuleFuture

	DisableRule(ctx workflow.Context, input *cloudwatchevents.DisableRuleInput) (*cloudwatchevents.DisableRuleOutput, error)
	DisableRuleAsync(ctx workflow.Context, input *cloudwatchevents.DisableRuleInput) *CloudWatchEventsDisableRuleFuture

	EnableRule(ctx workflow.Context, input *cloudwatchevents.EnableRuleInput) (*cloudwatchevents.EnableRuleOutput, error)
	EnableRuleAsync(ctx workflow.Context, input *cloudwatchevents.EnableRuleInput) *CloudWatchEventsEnableRuleFuture

	ListEventBuses(ctx workflow.Context, input *cloudwatchevents.ListEventBusesInput) (*cloudwatchevents.ListEventBusesOutput, error)
	ListEventBusesAsync(ctx workflow.Context, input *cloudwatchevents.ListEventBusesInput) *CloudWatchEventsListEventBusesFuture

	ListEventSources(ctx workflow.Context, input *cloudwatchevents.ListEventSourcesInput) (*cloudwatchevents.ListEventSourcesOutput, error)
	ListEventSourcesAsync(ctx workflow.Context, input *cloudwatchevents.ListEventSourcesInput) *CloudWatchEventsListEventSourcesFuture

	ListPartnerEventSourceAccounts(ctx workflow.Context, input *cloudwatchevents.ListPartnerEventSourceAccountsInput) (*cloudwatchevents.ListPartnerEventSourceAccountsOutput, error)
	ListPartnerEventSourceAccountsAsync(ctx workflow.Context, input *cloudwatchevents.ListPartnerEventSourceAccountsInput) *CloudWatchEventsListPartnerEventSourceAccountsFuture

	ListPartnerEventSources(ctx workflow.Context, input *cloudwatchevents.ListPartnerEventSourcesInput) (*cloudwatchevents.ListPartnerEventSourcesOutput, error)
	ListPartnerEventSourcesAsync(ctx workflow.Context, input *cloudwatchevents.ListPartnerEventSourcesInput) *CloudWatchEventsListPartnerEventSourcesFuture

	ListRuleNamesByTarget(ctx workflow.Context, input *cloudwatchevents.ListRuleNamesByTargetInput) (*cloudwatchevents.ListRuleNamesByTargetOutput, error)
	ListRuleNamesByTargetAsync(ctx workflow.Context, input *cloudwatchevents.ListRuleNamesByTargetInput) *CloudWatchEventsListRuleNamesByTargetFuture

	ListRules(ctx workflow.Context, input *cloudwatchevents.ListRulesInput) (*cloudwatchevents.ListRulesOutput, error)
	ListRulesAsync(ctx workflow.Context, input *cloudwatchevents.ListRulesInput) *CloudWatchEventsListRulesFuture

	ListTagsForResource(ctx workflow.Context, input *cloudwatchevents.ListTagsForResourceInput) (*cloudwatchevents.ListTagsForResourceOutput, error)
	ListTagsForResourceAsync(ctx workflow.Context, input *cloudwatchevents.ListTagsForResourceInput) *CloudWatchEventsListTagsForResourceFuture

	ListTargetsByRule(ctx workflow.Context, input *cloudwatchevents.ListTargetsByRuleInput) (*cloudwatchevents.ListTargetsByRuleOutput, error)
	ListTargetsByRuleAsync(ctx workflow.Context, input *cloudwatchevents.ListTargetsByRuleInput) *CloudWatchEventsListTargetsByRuleFuture

	PutEvents(ctx workflow.Context, input *cloudwatchevents.PutEventsInput) (*cloudwatchevents.PutEventsOutput, error)
	PutEventsAsync(ctx workflow.Context, input *cloudwatchevents.PutEventsInput) *CloudWatchEventsPutEventsFuture

	PutPartnerEvents(ctx workflow.Context, input *cloudwatchevents.PutPartnerEventsInput) (*cloudwatchevents.PutPartnerEventsOutput, error)
	PutPartnerEventsAsync(ctx workflow.Context, input *cloudwatchevents.PutPartnerEventsInput) *CloudWatchEventsPutPartnerEventsFuture

	PutPermission(ctx workflow.Context, input *cloudwatchevents.PutPermissionInput) (*cloudwatchevents.PutPermissionOutput, error)
	PutPermissionAsync(ctx workflow.Context, input *cloudwatchevents.PutPermissionInput) *CloudWatchEventsPutPermissionFuture

	PutRule(ctx workflow.Context, input *cloudwatchevents.PutRuleInput) (*cloudwatchevents.PutRuleOutput, error)
	PutRuleAsync(ctx workflow.Context, input *cloudwatchevents.PutRuleInput) *CloudWatchEventsPutRuleFuture

	PutTargets(ctx workflow.Context, input *cloudwatchevents.PutTargetsInput) (*cloudwatchevents.PutTargetsOutput, error)
	PutTargetsAsync(ctx workflow.Context, input *cloudwatchevents.PutTargetsInput) *CloudWatchEventsPutTargetsFuture

	RemovePermission(ctx workflow.Context, input *cloudwatchevents.RemovePermissionInput) (*cloudwatchevents.RemovePermissionOutput, error)
	RemovePermissionAsync(ctx workflow.Context, input *cloudwatchevents.RemovePermissionInput) *CloudWatchEventsRemovePermissionFuture

	RemoveTargets(ctx workflow.Context, input *cloudwatchevents.RemoveTargetsInput) (*cloudwatchevents.RemoveTargetsOutput, error)
	RemoveTargetsAsync(ctx workflow.Context, input *cloudwatchevents.RemoveTargetsInput) *CloudWatchEventsRemoveTargetsFuture

	TagResource(ctx workflow.Context, input *cloudwatchevents.TagResourceInput) (*cloudwatchevents.TagResourceOutput, error)
	TagResourceAsync(ctx workflow.Context, input *cloudwatchevents.TagResourceInput) *CloudWatchEventsTagResourceFuture

	TestEventPattern(ctx workflow.Context, input *cloudwatchevents.TestEventPatternInput) (*cloudwatchevents.TestEventPatternOutput, error)
	TestEventPatternAsync(ctx workflow.Context, input *cloudwatchevents.TestEventPatternInput) *CloudWatchEventsTestEventPatternFuture

	UntagResource(ctx workflow.Context, input *cloudwatchevents.UntagResourceInput) (*cloudwatchevents.UntagResourceOutput, error)
	UntagResourceAsync(ctx workflow.Context, input *cloudwatchevents.UntagResourceInput) *CloudWatchEventsUntagResourceFuture
}

func NewClient() Client {
	return &stub{}
}
