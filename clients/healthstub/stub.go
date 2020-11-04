// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.
// Copyright (c) 2020 Banzai Cloud All rights reserved.

package healthstub

import (
	"github.com/aws/aws-sdk-go/service/health"
	"go.uber.org/cadence/workflow"

	"github.com/banzaicloud/cadence-aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type stub struct{}

type HealthDescribeAffectedAccountsForOrganizationFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *HealthDescribeAffectedAccountsForOrganizationFuture) Get(ctx workflow.Context) (*health.DescribeAffectedAccountsForOrganizationOutput, error) {
	var output health.DescribeAffectedAccountsForOrganizationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type HealthDescribeAffectedEntitiesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *HealthDescribeAffectedEntitiesFuture) Get(ctx workflow.Context) (*health.DescribeAffectedEntitiesOutput, error) {
	var output health.DescribeAffectedEntitiesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type HealthDescribeAffectedEntitiesForOrganizationFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *HealthDescribeAffectedEntitiesForOrganizationFuture) Get(ctx workflow.Context) (*health.DescribeAffectedEntitiesForOrganizationOutput, error) {
	var output health.DescribeAffectedEntitiesForOrganizationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type HealthDescribeEntityAggregatesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *HealthDescribeEntityAggregatesFuture) Get(ctx workflow.Context) (*health.DescribeEntityAggregatesOutput, error) {
	var output health.DescribeEntityAggregatesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type HealthDescribeEventAggregatesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *HealthDescribeEventAggregatesFuture) Get(ctx workflow.Context) (*health.DescribeEventAggregatesOutput, error) {
	var output health.DescribeEventAggregatesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type HealthDescribeEventDetailsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *HealthDescribeEventDetailsFuture) Get(ctx workflow.Context) (*health.DescribeEventDetailsOutput, error) {
	var output health.DescribeEventDetailsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type HealthDescribeEventDetailsForOrganizationFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *HealthDescribeEventDetailsForOrganizationFuture) Get(ctx workflow.Context) (*health.DescribeEventDetailsForOrganizationOutput, error) {
	var output health.DescribeEventDetailsForOrganizationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type HealthDescribeEventTypesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *HealthDescribeEventTypesFuture) Get(ctx workflow.Context) (*health.DescribeEventTypesOutput, error) {
	var output health.DescribeEventTypesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type HealthDescribeEventsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *HealthDescribeEventsFuture) Get(ctx workflow.Context) (*health.DescribeEventsOutput, error) {
	var output health.DescribeEventsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type HealthDescribeEventsForOrganizationFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *HealthDescribeEventsForOrganizationFuture) Get(ctx workflow.Context) (*health.DescribeEventsForOrganizationOutput, error) {
	var output health.DescribeEventsForOrganizationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type HealthDescribeHealthServiceStatusForOrganizationFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *HealthDescribeHealthServiceStatusForOrganizationFuture) Get(ctx workflow.Context) (*health.DescribeHealthServiceStatusForOrganizationOutput, error) {
	var output health.DescribeHealthServiceStatusForOrganizationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type HealthDisableHealthServiceAccessForOrganizationFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *HealthDisableHealthServiceAccessForOrganizationFuture) Get(ctx workflow.Context) (*health.DisableHealthServiceAccessForOrganizationOutput, error) {
	var output health.DisableHealthServiceAccessForOrganizationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type HealthEnableHealthServiceAccessForOrganizationFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *HealthEnableHealthServiceAccessForOrganizationFuture) Get(ctx workflow.Context) (*health.EnableHealthServiceAccessForOrganizationOutput, error) {
	var output health.EnableHealthServiceAccessForOrganizationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeAffectedAccountsForOrganization(ctx workflow.Context, input *health.DescribeAffectedAccountsForOrganizationInput) (*health.DescribeAffectedAccountsForOrganizationOutput, error) {
	var output health.DescribeAffectedAccountsForOrganizationOutput
	err := workflow.ExecuteActivity(ctx, "aws-health-DescribeAffectedAccountsForOrganization", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeAffectedAccountsForOrganizationAsync(ctx workflow.Context, input *health.DescribeAffectedAccountsForOrganizationInput) *HealthDescribeAffectedAccountsForOrganizationFuture {
	future := workflow.ExecuteActivity(ctx, "aws-health-DescribeAffectedAccountsForOrganization", input)
	return &HealthDescribeAffectedAccountsForOrganizationFuture{Future: future}
}

func (a *stub) DescribeAffectedEntities(ctx workflow.Context, input *health.DescribeAffectedEntitiesInput) (*health.DescribeAffectedEntitiesOutput, error) {
	var output health.DescribeAffectedEntitiesOutput
	err := workflow.ExecuteActivity(ctx, "aws-health-DescribeAffectedEntities", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeAffectedEntitiesAsync(ctx workflow.Context, input *health.DescribeAffectedEntitiesInput) *HealthDescribeAffectedEntitiesFuture {
	future := workflow.ExecuteActivity(ctx, "aws-health-DescribeAffectedEntities", input)
	return &HealthDescribeAffectedEntitiesFuture{Future: future}
}

func (a *stub) DescribeAffectedEntitiesForOrganization(ctx workflow.Context, input *health.DescribeAffectedEntitiesForOrganizationInput) (*health.DescribeAffectedEntitiesForOrganizationOutput, error) {
	var output health.DescribeAffectedEntitiesForOrganizationOutput
	err := workflow.ExecuteActivity(ctx, "aws-health-DescribeAffectedEntitiesForOrganization", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeAffectedEntitiesForOrganizationAsync(ctx workflow.Context, input *health.DescribeAffectedEntitiesForOrganizationInput) *HealthDescribeAffectedEntitiesForOrganizationFuture {
	future := workflow.ExecuteActivity(ctx, "aws-health-DescribeAffectedEntitiesForOrganization", input)
	return &HealthDescribeAffectedEntitiesForOrganizationFuture{Future: future}
}

func (a *stub) DescribeEntityAggregates(ctx workflow.Context, input *health.DescribeEntityAggregatesInput) (*health.DescribeEntityAggregatesOutput, error) {
	var output health.DescribeEntityAggregatesOutput
	err := workflow.ExecuteActivity(ctx, "aws-health-DescribeEntityAggregates", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeEntityAggregatesAsync(ctx workflow.Context, input *health.DescribeEntityAggregatesInput) *HealthDescribeEntityAggregatesFuture {
	future := workflow.ExecuteActivity(ctx, "aws-health-DescribeEntityAggregates", input)
	return &HealthDescribeEntityAggregatesFuture{Future: future}
}

func (a *stub) DescribeEventAggregates(ctx workflow.Context, input *health.DescribeEventAggregatesInput) (*health.DescribeEventAggregatesOutput, error) {
	var output health.DescribeEventAggregatesOutput
	err := workflow.ExecuteActivity(ctx, "aws-health-DescribeEventAggregates", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeEventAggregatesAsync(ctx workflow.Context, input *health.DescribeEventAggregatesInput) *HealthDescribeEventAggregatesFuture {
	future := workflow.ExecuteActivity(ctx, "aws-health-DescribeEventAggregates", input)
	return &HealthDescribeEventAggregatesFuture{Future: future}
}

func (a *stub) DescribeEventDetails(ctx workflow.Context, input *health.DescribeEventDetailsInput) (*health.DescribeEventDetailsOutput, error) {
	var output health.DescribeEventDetailsOutput
	err := workflow.ExecuteActivity(ctx, "aws-health-DescribeEventDetails", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeEventDetailsAsync(ctx workflow.Context, input *health.DescribeEventDetailsInput) *HealthDescribeEventDetailsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-health-DescribeEventDetails", input)
	return &HealthDescribeEventDetailsFuture{Future: future}
}

func (a *stub) DescribeEventDetailsForOrganization(ctx workflow.Context, input *health.DescribeEventDetailsForOrganizationInput) (*health.DescribeEventDetailsForOrganizationOutput, error) {
	var output health.DescribeEventDetailsForOrganizationOutput
	err := workflow.ExecuteActivity(ctx, "aws-health-DescribeEventDetailsForOrganization", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeEventDetailsForOrganizationAsync(ctx workflow.Context, input *health.DescribeEventDetailsForOrganizationInput) *HealthDescribeEventDetailsForOrganizationFuture {
	future := workflow.ExecuteActivity(ctx, "aws-health-DescribeEventDetailsForOrganization", input)
	return &HealthDescribeEventDetailsForOrganizationFuture{Future: future}
}

func (a *stub) DescribeEventTypes(ctx workflow.Context, input *health.DescribeEventTypesInput) (*health.DescribeEventTypesOutput, error) {
	var output health.DescribeEventTypesOutput
	err := workflow.ExecuteActivity(ctx, "aws-health-DescribeEventTypes", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeEventTypesAsync(ctx workflow.Context, input *health.DescribeEventTypesInput) *HealthDescribeEventTypesFuture {
	future := workflow.ExecuteActivity(ctx, "aws-health-DescribeEventTypes", input)
	return &HealthDescribeEventTypesFuture{Future: future}
}

func (a *stub) DescribeEvents(ctx workflow.Context, input *health.DescribeEventsInput) (*health.DescribeEventsOutput, error) {
	var output health.DescribeEventsOutput
	err := workflow.ExecuteActivity(ctx, "aws-health-DescribeEvents", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeEventsAsync(ctx workflow.Context, input *health.DescribeEventsInput) *HealthDescribeEventsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-health-DescribeEvents", input)
	return &HealthDescribeEventsFuture{Future: future}
}

func (a *stub) DescribeEventsForOrganization(ctx workflow.Context, input *health.DescribeEventsForOrganizationInput) (*health.DescribeEventsForOrganizationOutput, error) {
	var output health.DescribeEventsForOrganizationOutput
	err := workflow.ExecuteActivity(ctx, "aws-health-DescribeEventsForOrganization", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeEventsForOrganizationAsync(ctx workflow.Context, input *health.DescribeEventsForOrganizationInput) *HealthDescribeEventsForOrganizationFuture {
	future := workflow.ExecuteActivity(ctx, "aws-health-DescribeEventsForOrganization", input)
	return &HealthDescribeEventsForOrganizationFuture{Future: future}
}

func (a *stub) DescribeHealthServiceStatusForOrganization(ctx workflow.Context, input *health.DescribeHealthServiceStatusForOrganizationInput) (*health.DescribeHealthServiceStatusForOrganizationOutput, error) {
	var output health.DescribeHealthServiceStatusForOrganizationOutput
	err := workflow.ExecuteActivity(ctx, "aws-health-DescribeHealthServiceStatusForOrganization", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeHealthServiceStatusForOrganizationAsync(ctx workflow.Context, input *health.DescribeHealthServiceStatusForOrganizationInput) *HealthDescribeHealthServiceStatusForOrganizationFuture {
	future := workflow.ExecuteActivity(ctx, "aws-health-DescribeHealthServiceStatusForOrganization", input)
	return &HealthDescribeHealthServiceStatusForOrganizationFuture{Future: future}
}

func (a *stub) DisableHealthServiceAccessForOrganization(ctx workflow.Context, input *health.DisableHealthServiceAccessForOrganizationInput) (*health.DisableHealthServiceAccessForOrganizationOutput, error) {
	var output health.DisableHealthServiceAccessForOrganizationOutput
	err := workflow.ExecuteActivity(ctx, "aws-health-DisableHealthServiceAccessForOrganization", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DisableHealthServiceAccessForOrganizationAsync(ctx workflow.Context, input *health.DisableHealthServiceAccessForOrganizationInput) *HealthDisableHealthServiceAccessForOrganizationFuture {
	future := workflow.ExecuteActivity(ctx, "aws-health-DisableHealthServiceAccessForOrganization", input)
	return &HealthDisableHealthServiceAccessForOrganizationFuture{Future: future}
}

func (a *stub) EnableHealthServiceAccessForOrganization(ctx workflow.Context, input *health.EnableHealthServiceAccessForOrganizationInput) (*health.EnableHealthServiceAccessForOrganizationOutput, error) {
	var output health.EnableHealthServiceAccessForOrganizationOutput
	err := workflow.ExecuteActivity(ctx, "aws-health-EnableHealthServiceAccessForOrganization", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) EnableHealthServiceAccessForOrganizationAsync(ctx workflow.Context, input *health.EnableHealthServiceAccessForOrganizationInput) *HealthEnableHealthServiceAccessForOrganizationFuture {
	future := workflow.ExecuteActivity(ctx, "aws-health-EnableHealthServiceAccessForOrganization", input)
	return &HealthEnableHealthServiceAccessForOrganizationFuture{Future: future}
}
