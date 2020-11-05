// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7

package eventbridgestub

import (
	"github.com/aws/aws-sdk-go/service/eventbridge"
	"go.uber.org/cadence/workflow"

	"github.com/banzaicloud/cadence-aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type stub struct{}

type EventBridgeActivateEventSourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *EventBridgeActivateEventSourceFuture) Get(ctx workflow.Context) (*eventbridge.ActivateEventSourceOutput, error) {
	var output eventbridge.ActivateEventSourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type EventBridgeCreateEventBusFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *EventBridgeCreateEventBusFuture) Get(ctx workflow.Context) (*eventbridge.CreateEventBusOutput, error) {
	var output eventbridge.CreateEventBusOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type EventBridgeCreatePartnerEventSourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *EventBridgeCreatePartnerEventSourceFuture) Get(ctx workflow.Context) (*eventbridge.CreatePartnerEventSourceOutput, error) {
	var output eventbridge.CreatePartnerEventSourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type EventBridgeDeactivateEventSourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *EventBridgeDeactivateEventSourceFuture) Get(ctx workflow.Context) (*eventbridge.DeactivateEventSourceOutput, error) {
	var output eventbridge.DeactivateEventSourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type EventBridgeDeleteEventBusFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *EventBridgeDeleteEventBusFuture) Get(ctx workflow.Context) (*eventbridge.DeleteEventBusOutput, error) {
	var output eventbridge.DeleteEventBusOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type EventBridgeDeletePartnerEventSourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *EventBridgeDeletePartnerEventSourceFuture) Get(ctx workflow.Context) (*eventbridge.DeletePartnerEventSourceOutput, error) {
	var output eventbridge.DeletePartnerEventSourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type EventBridgeDeleteRuleFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *EventBridgeDeleteRuleFuture) Get(ctx workflow.Context) (*eventbridge.DeleteRuleOutput, error) {
	var output eventbridge.DeleteRuleOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type EventBridgeDescribeEventBusFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *EventBridgeDescribeEventBusFuture) Get(ctx workflow.Context) (*eventbridge.DescribeEventBusOutput, error) {
	var output eventbridge.DescribeEventBusOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type EventBridgeDescribeEventSourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *EventBridgeDescribeEventSourceFuture) Get(ctx workflow.Context) (*eventbridge.DescribeEventSourceOutput, error) {
	var output eventbridge.DescribeEventSourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type EventBridgeDescribePartnerEventSourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *EventBridgeDescribePartnerEventSourceFuture) Get(ctx workflow.Context) (*eventbridge.DescribePartnerEventSourceOutput, error) {
	var output eventbridge.DescribePartnerEventSourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type EventBridgeDescribeRuleFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *EventBridgeDescribeRuleFuture) Get(ctx workflow.Context) (*eventbridge.DescribeRuleOutput, error) {
	var output eventbridge.DescribeRuleOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type EventBridgeDisableRuleFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *EventBridgeDisableRuleFuture) Get(ctx workflow.Context) (*eventbridge.DisableRuleOutput, error) {
	var output eventbridge.DisableRuleOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type EventBridgeEnableRuleFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *EventBridgeEnableRuleFuture) Get(ctx workflow.Context) (*eventbridge.EnableRuleOutput, error) {
	var output eventbridge.EnableRuleOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type EventBridgeListEventBusesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *EventBridgeListEventBusesFuture) Get(ctx workflow.Context) (*eventbridge.ListEventBusesOutput, error) {
	var output eventbridge.ListEventBusesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type EventBridgeListEventSourcesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *EventBridgeListEventSourcesFuture) Get(ctx workflow.Context) (*eventbridge.ListEventSourcesOutput, error) {
	var output eventbridge.ListEventSourcesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type EventBridgeListPartnerEventSourceAccountsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *EventBridgeListPartnerEventSourceAccountsFuture) Get(ctx workflow.Context) (*eventbridge.ListPartnerEventSourceAccountsOutput, error) {
	var output eventbridge.ListPartnerEventSourceAccountsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type EventBridgeListPartnerEventSourcesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *EventBridgeListPartnerEventSourcesFuture) Get(ctx workflow.Context) (*eventbridge.ListPartnerEventSourcesOutput, error) {
	var output eventbridge.ListPartnerEventSourcesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type EventBridgeListRuleNamesByTargetFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *EventBridgeListRuleNamesByTargetFuture) Get(ctx workflow.Context) (*eventbridge.ListRuleNamesByTargetOutput, error) {
	var output eventbridge.ListRuleNamesByTargetOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type EventBridgeListRulesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *EventBridgeListRulesFuture) Get(ctx workflow.Context) (*eventbridge.ListRulesOutput, error) {
	var output eventbridge.ListRulesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type EventBridgeListTagsForResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *EventBridgeListTagsForResourceFuture) Get(ctx workflow.Context) (*eventbridge.ListTagsForResourceOutput, error) {
	var output eventbridge.ListTagsForResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type EventBridgeListTargetsByRuleFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *EventBridgeListTargetsByRuleFuture) Get(ctx workflow.Context) (*eventbridge.ListTargetsByRuleOutput, error) {
	var output eventbridge.ListTargetsByRuleOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type EventBridgePutEventsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *EventBridgePutEventsFuture) Get(ctx workflow.Context) (*eventbridge.PutEventsOutput, error) {
	var output eventbridge.PutEventsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type EventBridgePutPartnerEventsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *EventBridgePutPartnerEventsFuture) Get(ctx workflow.Context) (*eventbridge.PutPartnerEventsOutput, error) {
	var output eventbridge.PutPartnerEventsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type EventBridgePutPermissionFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *EventBridgePutPermissionFuture) Get(ctx workflow.Context) (*eventbridge.PutPermissionOutput, error) {
	var output eventbridge.PutPermissionOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type EventBridgePutRuleFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *EventBridgePutRuleFuture) Get(ctx workflow.Context) (*eventbridge.PutRuleOutput, error) {
	var output eventbridge.PutRuleOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type EventBridgePutTargetsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *EventBridgePutTargetsFuture) Get(ctx workflow.Context) (*eventbridge.PutTargetsOutput, error) {
	var output eventbridge.PutTargetsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type EventBridgeRemovePermissionFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *EventBridgeRemovePermissionFuture) Get(ctx workflow.Context) (*eventbridge.RemovePermissionOutput, error) {
	var output eventbridge.RemovePermissionOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type EventBridgeRemoveTargetsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *EventBridgeRemoveTargetsFuture) Get(ctx workflow.Context) (*eventbridge.RemoveTargetsOutput, error) {
	var output eventbridge.RemoveTargetsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type EventBridgeTagResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *EventBridgeTagResourceFuture) Get(ctx workflow.Context) (*eventbridge.TagResourceOutput, error) {
	var output eventbridge.TagResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type EventBridgeTestEventPatternFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *EventBridgeTestEventPatternFuture) Get(ctx workflow.Context) (*eventbridge.TestEventPatternOutput, error) {
	var output eventbridge.TestEventPatternOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type EventBridgeUntagResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *EventBridgeUntagResourceFuture) Get(ctx workflow.Context) (*eventbridge.UntagResourceOutput, error) {
	var output eventbridge.UntagResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *stub) ActivateEventSource(ctx workflow.Context, input *eventbridge.ActivateEventSourceInput) (*eventbridge.ActivateEventSourceOutput, error) {
	var output eventbridge.ActivateEventSourceOutput
	err := workflow.ExecuteActivity(ctx, "aws-eventbridge-ActivateEventSource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ActivateEventSourceAsync(ctx workflow.Context, input *eventbridge.ActivateEventSourceInput) *EventBridgeActivateEventSourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws-eventbridge-ActivateEventSource", input)
	return &EventBridgeActivateEventSourceFuture{Future: future}
}

func (a *stub) CreateEventBus(ctx workflow.Context, input *eventbridge.CreateEventBusInput) (*eventbridge.CreateEventBusOutput, error) {
	var output eventbridge.CreateEventBusOutput
	err := workflow.ExecuteActivity(ctx, "aws-eventbridge-CreateEventBus", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateEventBusAsync(ctx workflow.Context, input *eventbridge.CreateEventBusInput) *EventBridgeCreateEventBusFuture {
	future := workflow.ExecuteActivity(ctx, "aws-eventbridge-CreateEventBus", input)
	return &EventBridgeCreateEventBusFuture{Future: future}
}

func (a *stub) CreatePartnerEventSource(ctx workflow.Context, input *eventbridge.CreatePartnerEventSourceInput) (*eventbridge.CreatePartnerEventSourceOutput, error) {
	var output eventbridge.CreatePartnerEventSourceOutput
	err := workflow.ExecuteActivity(ctx, "aws-eventbridge-CreatePartnerEventSource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreatePartnerEventSourceAsync(ctx workflow.Context, input *eventbridge.CreatePartnerEventSourceInput) *EventBridgeCreatePartnerEventSourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws-eventbridge-CreatePartnerEventSource", input)
	return &EventBridgeCreatePartnerEventSourceFuture{Future: future}
}

func (a *stub) DeactivateEventSource(ctx workflow.Context, input *eventbridge.DeactivateEventSourceInput) (*eventbridge.DeactivateEventSourceOutput, error) {
	var output eventbridge.DeactivateEventSourceOutput
	err := workflow.ExecuteActivity(ctx, "aws-eventbridge-DeactivateEventSource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeactivateEventSourceAsync(ctx workflow.Context, input *eventbridge.DeactivateEventSourceInput) *EventBridgeDeactivateEventSourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws-eventbridge-DeactivateEventSource", input)
	return &EventBridgeDeactivateEventSourceFuture{Future: future}
}

func (a *stub) DeleteEventBus(ctx workflow.Context, input *eventbridge.DeleteEventBusInput) (*eventbridge.DeleteEventBusOutput, error) {
	var output eventbridge.DeleteEventBusOutput
	err := workflow.ExecuteActivity(ctx, "aws-eventbridge-DeleteEventBus", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteEventBusAsync(ctx workflow.Context, input *eventbridge.DeleteEventBusInput) *EventBridgeDeleteEventBusFuture {
	future := workflow.ExecuteActivity(ctx, "aws-eventbridge-DeleteEventBus", input)
	return &EventBridgeDeleteEventBusFuture{Future: future}
}

func (a *stub) DeletePartnerEventSource(ctx workflow.Context, input *eventbridge.DeletePartnerEventSourceInput) (*eventbridge.DeletePartnerEventSourceOutput, error) {
	var output eventbridge.DeletePartnerEventSourceOutput
	err := workflow.ExecuteActivity(ctx, "aws-eventbridge-DeletePartnerEventSource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeletePartnerEventSourceAsync(ctx workflow.Context, input *eventbridge.DeletePartnerEventSourceInput) *EventBridgeDeletePartnerEventSourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws-eventbridge-DeletePartnerEventSource", input)
	return &EventBridgeDeletePartnerEventSourceFuture{Future: future}
}

func (a *stub) DeleteRule(ctx workflow.Context, input *eventbridge.DeleteRuleInput) (*eventbridge.DeleteRuleOutput, error) {
	var output eventbridge.DeleteRuleOutput
	err := workflow.ExecuteActivity(ctx, "aws-eventbridge-DeleteRule", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteRuleAsync(ctx workflow.Context, input *eventbridge.DeleteRuleInput) *EventBridgeDeleteRuleFuture {
	future := workflow.ExecuteActivity(ctx, "aws-eventbridge-DeleteRule", input)
	return &EventBridgeDeleteRuleFuture{Future: future}
}

func (a *stub) DescribeEventBus(ctx workflow.Context, input *eventbridge.DescribeEventBusInput) (*eventbridge.DescribeEventBusOutput, error) {
	var output eventbridge.DescribeEventBusOutput
	err := workflow.ExecuteActivity(ctx, "aws-eventbridge-DescribeEventBus", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeEventBusAsync(ctx workflow.Context, input *eventbridge.DescribeEventBusInput) *EventBridgeDescribeEventBusFuture {
	future := workflow.ExecuteActivity(ctx, "aws-eventbridge-DescribeEventBus", input)
	return &EventBridgeDescribeEventBusFuture{Future: future}
}

func (a *stub) DescribeEventSource(ctx workflow.Context, input *eventbridge.DescribeEventSourceInput) (*eventbridge.DescribeEventSourceOutput, error) {
	var output eventbridge.DescribeEventSourceOutput
	err := workflow.ExecuteActivity(ctx, "aws-eventbridge-DescribeEventSource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeEventSourceAsync(ctx workflow.Context, input *eventbridge.DescribeEventSourceInput) *EventBridgeDescribeEventSourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws-eventbridge-DescribeEventSource", input)
	return &EventBridgeDescribeEventSourceFuture{Future: future}
}

func (a *stub) DescribePartnerEventSource(ctx workflow.Context, input *eventbridge.DescribePartnerEventSourceInput) (*eventbridge.DescribePartnerEventSourceOutput, error) {
	var output eventbridge.DescribePartnerEventSourceOutput
	err := workflow.ExecuteActivity(ctx, "aws-eventbridge-DescribePartnerEventSource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribePartnerEventSourceAsync(ctx workflow.Context, input *eventbridge.DescribePartnerEventSourceInput) *EventBridgeDescribePartnerEventSourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws-eventbridge-DescribePartnerEventSource", input)
	return &EventBridgeDescribePartnerEventSourceFuture{Future: future}
}

func (a *stub) DescribeRule(ctx workflow.Context, input *eventbridge.DescribeRuleInput) (*eventbridge.DescribeRuleOutput, error) {
	var output eventbridge.DescribeRuleOutput
	err := workflow.ExecuteActivity(ctx, "aws-eventbridge-DescribeRule", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeRuleAsync(ctx workflow.Context, input *eventbridge.DescribeRuleInput) *EventBridgeDescribeRuleFuture {
	future := workflow.ExecuteActivity(ctx, "aws-eventbridge-DescribeRule", input)
	return &EventBridgeDescribeRuleFuture{Future: future}
}

func (a *stub) DisableRule(ctx workflow.Context, input *eventbridge.DisableRuleInput) (*eventbridge.DisableRuleOutput, error) {
	var output eventbridge.DisableRuleOutput
	err := workflow.ExecuteActivity(ctx, "aws-eventbridge-DisableRule", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DisableRuleAsync(ctx workflow.Context, input *eventbridge.DisableRuleInput) *EventBridgeDisableRuleFuture {
	future := workflow.ExecuteActivity(ctx, "aws-eventbridge-DisableRule", input)
	return &EventBridgeDisableRuleFuture{Future: future}
}

func (a *stub) EnableRule(ctx workflow.Context, input *eventbridge.EnableRuleInput) (*eventbridge.EnableRuleOutput, error) {
	var output eventbridge.EnableRuleOutput
	err := workflow.ExecuteActivity(ctx, "aws-eventbridge-EnableRule", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) EnableRuleAsync(ctx workflow.Context, input *eventbridge.EnableRuleInput) *EventBridgeEnableRuleFuture {
	future := workflow.ExecuteActivity(ctx, "aws-eventbridge-EnableRule", input)
	return &EventBridgeEnableRuleFuture{Future: future}
}

func (a *stub) ListEventBuses(ctx workflow.Context, input *eventbridge.ListEventBusesInput) (*eventbridge.ListEventBusesOutput, error) {
	var output eventbridge.ListEventBusesOutput
	err := workflow.ExecuteActivity(ctx, "aws-eventbridge-ListEventBuses", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListEventBusesAsync(ctx workflow.Context, input *eventbridge.ListEventBusesInput) *EventBridgeListEventBusesFuture {
	future := workflow.ExecuteActivity(ctx, "aws-eventbridge-ListEventBuses", input)
	return &EventBridgeListEventBusesFuture{Future: future}
}

func (a *stub) ListEventSources(ctx workflow.Context, input *eventbridge.ListEventSourcesInput) (*eventbridge.ListEventSourcesOutput, error) {
	var output eventbridge.ListEventSourcesOutput
	err := workflow.ExecuteActivity(ctx, "aws-eventbridge-ListEventSources", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListEventSourcesAsync(ctx workflow.Context, input *eventbridge.ListEventSourcesInput) *EventBridgeListEventSourcesFuture {
	future := workflow.ExecuteActivity(ctx, "aws-eventbridge-ListEventSources", input)
	return &EventBridgeListEventSourcesFuture{Future: future}
}

func (a *stub) ListPartnerEventSourceAccounts(ctx workflow.Context, input *eventbridge.ListPartnerEventSourceAccountsInput) (*eventbridge.ListPartnerEventSourceAccountsOutput, error) {
	var output eventbridge.ListPartnerEventSourceAccountsOutput
	err := workflow.ExecuteActivity(ctx, "aws-eventbridge-ListPartnerEventSourceAccounts", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListPartnerEventSourceAccountsAsync(ctx workflow.Context, input *eventbridge.ListPartnerEventSourceAccountsInput) *EventBridgeListPartnerEventSourceAccountsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-eventbridge-ListPartnerEventSourceAccounts", input)
	return &EventBridgeListPartnerEventSourceAccountsFuture{Future: future}
}

func (a *stub) ListPartnerEventSources(ctx workflow.Context, input *eventbridge.ListPartnerEventSourcesInput) (*eventbridge.ListPartnerEventSourcesOutput, error) {
	var output eventbridge.ListPartnerEventSourcesOutput
	err := workflow.ExecuteActivity(ctx, "aws-eventbridge-ListPartnerEventSources", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListPartnerEventSourcesAsync(ctx workflow.Context, input *eventbridge.ListPartnerEventSourcesInput) *EventBridgeListPartnerEventSourcesFuture {
	future := workflow.ExecuteActivity(ctx, "aws-eventbridge-ListPartnerEventSources", input)
	return &EventBridgeListPartnerEventSourcesFuture{Future: future}
}

func (a *stub) ListRuleNamesByTarget(ctx workflow.Context, input *eventbridge.ListRuleNamesByTargetInput) (*eventbridge.ListRuleNamesByTargetOutput, error) {
	var output eventbridge.ListRuleNamesByTargetOutput
	err := workflow.ExecuteActivity(ctx, "aws-eventbridge-ListRuleNamesByTarget", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListRuleNamesByTargetAsync(ctx workflow.Context, input *eventbridge.ListRuleNamesByTargetInput) *EventBridgeListRuleNamesByTargetFuture {
	future := workflow.ExecuteActivity(ctx, "aws-eventbridge-ListRuleNamesByTarget", input)
	return &EventBridgeListRuleNamesByTargetFuture{Future: future}
}

func (a *stub) ListRules(ctx workflow.Context, input *eventbridge.ListRulesInput) (*eventbridge.ListRulesOutput, error) {
	var output eventbridge.ListRulesOutput
	err := workflow.ExecuteActivity(ctx, "aws-eventbridge-ListRules", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListRulesAsync(ctx workflow.Context, input *eventbridge.ListRulesInput) *EventBridgeListRulesFuture {
	future := workflow.ExecuteActivity(ctx, "aws-eventbridge-ListRules", input)
	return &EventBridgeListRulesFuture{Future: future}
}

func (a *stub) ListTagsForResource(ctx workflow.Context, input *eventbridge.ListTagsForResourceInput) (*eventbridge.ListTagsForResourceOutput, error) {
	var output eventbridge.ListTagsForResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws-eventbridge-ListTagsForResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListTagsForResourceAsync(ctx workflow.Context, input *eventbridge.ListTagsForResourceInput) *EventBridgeListTagsForResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws-eventbridge-ListTagsForResource", input)
	return &EventBridgeListTagsForResourceFuture{Future: future}
}

func (a *stub) ListTargetsByRule(ctx workflow.Context, input *eventbridge.ListTargetsByRuleInput) (*eventbridge.ListTargetsByRuleOutput, error) {
	var output eventbridge.ListTargetsByRuleOutput
	err := workflow.ExecuteActivity(ctx, "aws-eventbridge-ListTargetsByRule", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListTargetsByRuleAsync(ctx workflow.Context, input *eventbridge.ListTargetsByRuleInput) *EventBridgeListTargetsByRuleFuture {
	future := workflow.ExecuteActivity(ctx, "aws-eventbridge-ListTargetsByRule", input)
	return &EventBridgeListTargetsByRuleFuture{Future: future}
}

func (a *stub) PutEvents(ctx workflow.Context, input *eventbridge.PutEventsInput) (*eventbridge.PutEventsOutput, error) {
	var output eventbridge.PutEventsOutput
	err := workflow.ExecuteActivity(ctx, "aws-eventbridge-PutEvents", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) PutEventsAsync(ctx workflow.Context, input *eventbridge.PutEventsInput) *EventBridgePutEventsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-eventbridge-PutEvents", input)
	return &EventBridgePutEventsFuture{Future: future}
}

func (a *stub) PutPartnerEvents(ctx workflow.Context, input *eventbridge.PutPartnerEventsInput) (*eventbridge.PutPartnerEventsOutput, error) {
	var output eventbridge.PutPartnerEventsOutput
	err := workflow.ExecuteActivity(ctx, "aws-eventbridge-PutPartnerEvents", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) PutPartnerEventsAsync(ctx workflow.Context, input *eventbridge.PutPartnerEventsInput) *EventBridgePutPartnerEventsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-eventbridge-PutPartnerEvents", input)
	return &EventBridgePutPartnerEventsFuture{Future: future}
}

func (a *stub) PutPermission(ctx workflow.Context, input *eventbridge.PutPermissionInput) (*eventbridge.PutPermissionOutput, error) {
	var output eventbridge.PutPermissionOutput
	err := workflow.ExecuteActivity(ctx, "aws-eventbridge-PutPermission", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) PutPermissionAsync(ctx workflow.Context, input *eventbridge.PutPermissionInput) *EventBridgePutPermissionFuture {
	future := workflow.ExecuteActivity(ctx, "aws-eventbridge-PutPermission", input)
	return &EventBridgePutPermissionFuture{Future: future}
}

func (a *stub) PutRule(ctx workflow.Context, input *eventbridge.PutRuleInput) (*eventbridge.PutRuleOutput, error) {
	var output eventbridge.PutRuleOutput
	err := workflow.ExecuteActivity(ctx, "aws-eventbridge-PutRule", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) PutRuleAsync(ctx workflow.Context, input *eventbridge.PutRuleInput) *EventBridgePutRuleFuture {
	future := workflow.ExecuteActivity(ctx, "aws-eventbridge-PutRule", input)
	return &EventBridgePutRuleFuture{Future: future}
}

func (a *stub) PutTargets(ctx workflow.Context, input *eventbridge.PutTargetsInput) (*eventbridge.PutTargetsOutput, error) {
	var output eventbridge.PutTargetsOutput
	err := workflow.ExecuteActivity(ctx, "aws-eventbridge-PutTargets", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) PutTargetsAsync(ctx workflow.Context, input *eventbridge.PutTargetsInput) *EventBridgePutTargetsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-eventbridge-PutTargets", input)
	return &EventBridgePutTargetsFuture{Future: future}
}

func (a *stub) RemovePermission(ctx workflow.Context, input *eventbridge.RemovePermissionInput) (*eventbridge.RemovePermissionOutput, error) {
	var output eventbridge.RemovePermissionOutput
	err := workflow.ExecuteActivity(ctx, "aws-eventbridge-RemovePermission", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) RemovePermissionAsync(ctx workflow.Context, input *eventbridge.RemovePermissionInput) *EventBridgeRemovePermissionFuture {
	future := workflow.ExecuteActivity(ctx, "aws-eventbridge-RemovePermission", input)
	return &EventBridgeRemovePermissionFuture{Future: future}
}

func (a *stub) RemoveTargets(ctx workflow.Context, input *eventbridge.RemoveTargetsInput) (*eventbridge.RemoveTargetsOutput, error) {
	var output eventbridge.RemoveTargetsOutput
	err := workflow.ExecuteActivity(ctx, "aws-eventbridge-RemoveTargets", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) RemoveTargetsAsync(ctx workflow.Context, input *eventbridge.RemoveTargetsInput) *EventBridgeRemoveTargetsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-eventbridge-RemoveTargets", input)
	return &EventBridgeRemoveTargetsFuture{Future: future}
}

func (a *stub) TagResource(ctx workflow.Context, input *eventbridge.TagResourceInput) (*eventbridge.TagResourceOutput, error) {
	var output eventbridge.TagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws-eventbridge-TagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) TagResourceAsync(ctx workflow.Context, input *eventbridge.TagResourceInput) *EventBridgeTagResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws-eventbridge-TagResource", input)
	return &EventBridgeTagResourceFuture{Future: future}
}

func (a *stub) TestEventPattern(ctx workflow.Context, input *eventbridge.TestEventPatternInput) (*eventbridge.TestEventPatternOutput, error) {
	var output eventbridge.TestEventPatternOutput
	err := workflow.ExecuteActivity(ctx, "aws-eventbridge-TestEventPattern", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) TestEventPatternAsync(ctx workflow.Context, input *eventbridge.TestEventPatternInput) *EventBridgeTestEventPatternFuture {
	future := workflow.ExecuteActivity(ctx, "aws-eventbridge-TestEventPattern", input)
	return &EventBridgeTestEventPatternFuture{Future: future}
}

func (a *stub) UntagResource(ctx workflow.Context, input *eventbridge.UntagResourceInput) (*eventbridge.UntagResourceOutput, error) {
	var output eventbridge.UntagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws-eventbridge-UntagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UntagResourceAsync(ctx workflow.Context, input *eventbridge.UntagResourceInput) *EventBridgeUntagResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws-eventbridge-UntagResource", input)
	return &EventBridgeUntagResourceFuture{Future: future}
}
