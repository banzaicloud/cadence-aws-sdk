// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.
// Copyright (c) 2020 Banzai Cloud All rights reserved.

package cloud9stub

import (
	"github.com/aws/aws-sdk-go/service/cloud9"
	"go.uber.org/cadence/workflow"

	"github.com/banzaicloud/cadence-aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type stub struct{}

type Cloud9CreateEnvironmentEC2Future struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *Cloud9CreateEnvironmentEC2Future) Get(ctx workflow.Context) (*cloud9.CreateEnvironmentEC2Output, error) {
	var output cloud9.CreateEnvironmentEC2Output
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type Cloud9CreateEnvironmentMembershipFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *Cloud9CreateEnvironmentMembershipFuture) Get(ctx workflow.Context) (*cloud9.CreateEnvironmentMembershipOutput, error) {
	var output cloud9.CreateEnvironmentMembershipOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type Cloud9DeleteEnvironmentFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *Cloud9DeleteEnvironmentFuture) Get(ctx workflow.Context) (*cloud9.DeleteEnvironmentOutput, error) {
	var output cloud9.DeleteEnvironmentOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type Cloud9DeleteEnvironmentMembershipFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *Cloud9DeleteEnvironmentMembershipFuture) Get(ctx workflow.Context) (*cloud9.DeleteEnvironmentMembershipOutput, error) {
	var output cloud9.DeleteEnvironmentMembershipOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type Cloud9DescribeEnvironmentMembershipsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *Cloud9DescribeEnvironmentMembershipsFuture) Get(ctx workflow.Context) (*cloud9.DescribeEnvironmentMembershipsOutput, error) {
	var output cloud9.DescribeEnvironmentMembershipsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type Cloud9DescribeEnvironmentStatusFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *Cloud9DescribeEnvironmentStatusFuture) Get(ctx workflow.Context) (*cloud9.DescribeEnvironmentStatusOutput, error) {
	var output cloud9.DescribeEnvironmentStatusOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type Cloud9DescribeEnvironmentsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *Cloud9DescribeEnvironmentsFuture) Get(ctx workflow.Context) (*cloud9.DescribeEnvironmentsOutput, error) {
	var output cloud9.DescribeEnvironmentsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type Cloud9ListEnvironmentsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *Cloud9ListEnvironmentsFuture) Get(ctx workflow.Context) (*cloud9.ListEnvironmentsOutput, error) {
	var output cloud9.ListEnvironmentsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type Cloud9ListTagsForResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *Cloud9ListTagsForResourceFuture) Get(ctx workflow.Context) (*cloud9.ListTagsForResourceOutput, error) {
	var output cloud9.ListTagsForResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type Cloud9TagResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *Cloud9TagResourceFuture) Get(ctx workflow.Context) (*cloud9.TagResourceOutput, error) {
	var output cloud9.TagResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type Cloud9UntagResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *Cloud9UntagResourceFuture) Get(ctx workflow.Context) (*cloud9.UntagResourceOutput, error) {
	var output cloud9.UntagResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type Cloud9UpdateEnvironmentFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *Cloud9UpdateEnvironmentFuture) Get(ctx workflow.Context) (*cloud9.UpdateEnvironmentOutput, error) {
	var output cloud9.UpdateEnvironmentOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type Cloud9UpdateEnvironmentMembershipFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *Cloud9UpdateEnvironmentMembershipFuture) Get(ctx workflow.Context) (*cloud9.UpdateEnvironmentMembershipOutput, error) {
	var output cloud9.UpdateEnvironmentMembershipOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateEnvironmentEC2(ctx workflow.Context, input *cloud9.CreateEnvironmentEC2Input) (*cloud9.CreateEnvironmentEC2Output, error) {
	var output cloud9.CreateEnvironmentEC2Output
	err := workflow.ExecuteActivity(ctx, "aws-cloud9-CreateEnvironmentEC2", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateEnvironmentEC2Async(ctx workflow.Context, input *cloud9.CreateEnvironmentEC2Input) *Cloud9CreateEnvironmentEC2Future {
	future := workflow.ExecuteActivity(ctx, "aws-cloud9-CreateEnvironmentEC2", input)
	return &Cloud9CreateEnvironmentEC2Future{Future: future}
}

func (a *stub) CreateEnvironmentMembership(ctx workflow.Context, input *cloud9.CreateEnvironmentMembershipInput) (*cloud9.CreateEnvironmentMembershipOutput, error) {
	var output cloud9.CreateEnvironmentMembershipOutput
	err := workflow.ExecuteActivity(ctx, "aws-cloud9-CreateEnvironmentMembership", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateEnvironmentMembershipAsync(ctx workflow.Context, input *cloud9.CreateEnvironmentMembershipInput) *Cloud9CreateEnvironmentMembershipFuture {
	future := workflow.ExecuteActivity(ctx, "aws-cloud9-CreateEnvironmentMembership", input)
	return &Cloud9CreateEnvironmentMembershipFuture{Future: future}
}

func (a *stub) DeleteEnvironment(ctx workflow.Context, input *cloud9.DeleteEnvironmentInput) (*cloud9.DeleteEnvironmentOutput, error) {
	var output cloud9.DeleteEnvironmentOutput
	err := workflow.ExecuteActivity(ctx, "aws-cloud9-DeleteEnvironment", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteEnvironmentAsync(ctx workflow.Context, input *cloud9.DeleteEnvironmentInput) *Cloud9DeleteEnvironmentFuture {
	future := workflow.ExecuteActivity(ctx, "aws-cloud9-DeleteEnvironment", input)
	return &Cloud9DeleteEnvironmentFuture{Future: future}
}

func (a *stub) DeleteEnvironmentMembership(ctx workflow.Context, input *cloud9.DeleteEnvironmentMembershipInput) (*cloud9.DeleteEnvironmentMembershipOutput, error) {
	var output cloud9.DeleteEnvironmentMembershipOutput
	err := workflow.ExecuteActivity(ctx, "aws-cloud9-DeleteEnvironmentMembership", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteEnvironmentMembershipAsync(ctx workflow.Context, input *cloud9.DeleteEnvironmentMembershipInput) *Cloud9DeleteEnvironmentMembershipFuture {
	future := workflow.ExecuteActivity(ctx, "aws-cloud9-DeleteEnvironmentMembership", input)
	return &Cloud9DeleteEnvironmentMembershipFuture{Future: future}
}

func (a *stub) DescribeEnvironmentMemberships(ctx workflow.Context, input *cloud9.DescribeEnvironmentMembershipsInput) (*cloud9.DescribeEnvironmentMembershipsOutput, error) {
	var output cloud9.DescribeEnvironmentMembershipsOutput
	err := workflow.ExecuteActivity(ctx, "aws-cloud9-DescribeEnvironmentMemberships", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeEnvironmentMembershipsAsync(ctx workflow.Context, input *cloud9.DescribeEnvironmentMembershipsInput) *Cloud9DescribeEnvironmentMembershipsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-cloud9-DescribeEnvironmentMemberships", input)
	return &Cloud9DescribeEnvironmentMembershipsFuture{Future: future}
}

func (a *stub) DescribeEnvironmentStatus(ctx workflow.Context, input *cloud9.DescribeEnvironmentStatusInput) (*cloud9.DescribeEnvironmentStatusOutput, error) {
	var output cloud9.DescribeEnvironmentStatusOutput
	err := workflow.ExecuteActivity(ctx, "aws-cloud9-DescribeEnvironmentStatus", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeEnvironmentStatusAsync(ctx workflow.Context, input *cloud9.DescribeEnvironmentStatusInput) *Cloud9DescribeEnvironmentStatusFuture {
	future := workflow.ExecuteActivity(ctx, "aws-cloud9-DescribeEnvironmentStatus", input)
	return &Cloud9DescribeEnvironmentStatusFuture{Future: future}
}

func (a *stub) DescribeEnvironments(ctx workflow.Context, input *cloud9.DescribeEnvironmentsInput) (*cloud9.DescribeEnvironmentsOutput, error) {
	var output cloud9.DescribeEnvironmentsOutput
	err := workflow.ExecuteActivity(ctx, "aws-cloud9-DescribeEnvironments", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeEnvironmentsAsync(ctx workflow.Context, input *cloud9.DescribeEnvironmentsInput) *Cloud9DescribeEnvironmentsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-cloud9-DescribeEnvironments", input)
	return &Cloud9DescribeEnvironmentsFuture{Future: future}
}

func (a *stub) ListEnvironments(ctx workflow.Context, input *cloud9.ListEnvironmentsInput) (*cloud9.ListEnvironmentsOutput, error) {
	var output cloud9.ListEnvironmentsOutput
	err := workflow.ExecuteActivity(ctx, "aws-cloud9-ListEnvironments", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListEnvironmentsAsync(ctx workflow.Context, input *cloud9.ListEnvironmentsInput) *Cloud9ListEnvironmentsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-cloud9-ListEnvironments", input)
	return &Cloud9ListEnvironmentsFuture{Future: future}
}

func (a *stub) ListTagsForResource(ctx workflow.Context, input *cloud9.ListTagsForResourceInput) (*cloud9.ListTagsForResourceOutput, error) {
	var output cloud9.ListTagsForResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws-cloud9-ListTagsForResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListTagsForResourceAsync(ctx workflow.Context, input *cloud9.ListTagsForResourceInput) *Cloud9ListTagsForResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws-cloud9-ListTagsForResource", input)
	return &Cloud9ListTagsForResourceFuture{Future: future}
}

func (a *stub) TagResource(ctx workflow.Context, input *cloud9.TagResourceInput) (*cloud9.TagResourceOutput, error) {
	var output cloud9.TagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws-cloud9-TagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) TagResourceAsync(ctx workflow.Context, input *cloud9.TagResourceInput) *Cloud9TagResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws-cloud9-TagResource", input)
	return &Cloud9TagResourceFuture{Future: future}
}

func (a *stub) UntagResource(ctx workflow.Context, input *cloud9.UntagResourceInput) (*cloud9.UntagResourceOutput, error) {
	var output cloud9.UntagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws-cloud9-UntagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UntagResourceAsync(ctx workflow.Context, input *cloud9.UntagResourceInput) *Cloud9UntagResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws-cloud9-UntagResource", input)
	return &Cloud9UntagResourceFuture{Future: future}
}

func (a *stub) UpdateEnvironment(ctx workflow.Context, input *cloud9.UpdateEnvironmentInput) (*cloud9.UpdateEnvironmentOutput, error) {
	var output cloud9.UpdateEnvironmentOutput
	err := workflow.ExecuteActivity(ctx, "aws-cloud9-UpdateEnvironment", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateEnvironmentAsync(ctx workflow.Context, input *cloud9.UpdateEnvironmentInput) *Cloud9UpdateEnvironmentFuture {
	future := workflow.ExecuteActivity(ctx, "aws-cloud9-UpdateEnvironment", input)
	return &Cloud9UpdateEnvironmentFuture{Future: future}
}

func (a *stub) UpdateEnvironmentMembership(ctx workflow.Context, input *cloud9.UpdateEnvironmentMembershipInput) (*cloud9.UpdateEnvironmentMembershipOutput, error) {
	var output cloud9.UpdateEnvironmentMembershipOutput
	err := workflow.ExecuteActivity(ctx, "aws-cloud9-UpdateEnvironmentMembership", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateEnvironmentMembershipAsync(ctx workflow.Context, input *cloud9.UpdateEnvironmentMembershipInput) *Cloud9UpdateEnvironmentMembershipFuture {
	future := workflow.ExecuteActivity(ctx, "aws-cloud9-UpdateEnvironmentMembership", input)
	return &Cloud9UpdateEnvironmentMembershipFuture{Future: future}
}
