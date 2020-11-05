// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7

package batchstub

import (
	"github.com/aws/aws-sdk-go/service/batch"
	"go.uber.org/cadence/workflow"

	"github.com/banzaicloud/cadence-aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type stub struct{}

type BatchCancelJobFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *BatchCancelJobFuture) Get(ctx workflow.Context) (*batch.CancelJobOutput, error) {
	var output batch.CancelJobOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type BatchCreateComputeEnvironmentFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *BatchCreateComputeEnvironmentFuture) Get(ctx workflow.Context) (*batch.CreateComputeEnvironmentOutput, error) {
	var output batch.CreateComputeEnvironmentOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type BatchCreateJobQueueFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *BatchCreateJobQueueFuture) Get(ctx workflow.Context) (*batch.CreateJobQueueOutput, error) {
	var output batch.CreateJobQueueOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type BatchDeleteComputeEnvironmentFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *BatchDeleteComputeEnvironmentFuture) Get(ctx workflow.Context) (*batch.DeleteComputeEnvironmentOutput, error) {
	var output batch.DeleteComputeEnvironmentOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type BatchDeleteJobQueueFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *BatchDeleteJobQueueFuture) Get(ctx workflow.Context) (*batch.DeleteJobQueueOutput, error) {
	var output batch.DeleteJobQueueOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type BatchDeregisterJobDefinitionFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *BatchDeregisterJobDefinitionFuture) Get(ctx workflow.Context) (*batch.DeregisterJobDefinitionOutput, error) {
	var output batch.DeregisterJobDefinitionOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type BatchDescribeComputeEnvironmentsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *BatchDescribeComputeEnvironmentsFuture) Get(ctx workflow.Context) (*batch.DescribeComputeEnvironmentsOutput, error) {
	var output batch.DescribeComputeEnvironmentsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type BatchDescribeJobDefinitionsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *BatchDescribeJobDefinitionsFuture) Get(ctx workflow.Context) (*batch.DescribeJobDefinitionsOutput, error) {
	var output batch.DescribeJobDefinitionsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type BatchDescribeJobQueuesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *BatchDescribeJobQueuesFuture) Get(ctx workflow.Context) (*batch.DescribeJobQueuesOutput, error) {
	var output batch.DescribeJobQueuesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type BatchDescribeJobsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *BatchDescribeJobsFuture) Get(ctx workflow.Context) (*batch.DescribeJobsOutput, error) {
	var output batch.DescribeJobsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type BatchListJobsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *BatchListJobsFuture) Get(ctx workflow.Context) (*batch.ListJobsOutput, error) {
	var output batch.ListJobsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type BatchListTagsForResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *BatchListTagsForResourceFuture) Get(ctx workflow.Context) (*batch.ListTagsForResourceOutput, error) {
	var output batch.ListTagsForResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type BatchRegisterJobDefinitionFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *BatchRegisterJobDefinitionFuture) Get(ctx workflow.Context) (*batch.RegisterJobDefinitionOutput, error) {
	var output batch.RegisterJobDefinitionOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type BatchSubmitJobFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *BatchSubmitJobFuture) Get(ctx workflow.Context) (*batch.SubmitJobOutput, error) {
	var output batch.SubmitJobOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type BatchTagResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *BatchTagResourceFuture) Get(ctx workflow.Context) (*batch.TagResourceOutput, error) {
	var output batch.TagResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type BatchTerminateJobFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *BatchTerminateJobFuture) Get(ctx workflow.Context) (*batch.TerminateJobOutput, error) {
	var output batch.TerminateJobOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type BatchUntagResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *BatchUntagResourceFuture) Get(ctx workflow.Context) (*batch.UntagResourceOutput, error) {
	var output batch.UntagResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type BatchUpdateComputeEnvironmentFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *BatchUpdateComputeEnvironmentFuture) Get(ctx workflow.Context) (*batch.UpdateComputeEnvironmentOutput, error) {
	var output batch.UpdateComputeEnvironmentOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type BatchUpdateJobQueueFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *BatchUpdateJobQueueFuture) Get(ctx workflow.Context) (*batch.UpdateJobQueueOutput, error) {
	var output batch.UpdateJobQueueOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *stub) CancelJob(ctx workflow.Context, input *batch.CancelJobInput) (*batch.CancelJobOutput, error) {
	var output batch.CancelJobOutput
	err := workflow.ExecuteActivity(ctx, "aws-batch-CancelJob", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CancelJobAsync(ctx workflow.Context, input *batch.CancelJobInput) *BatchCancelJobFuture {
	future := workflow.ExecuteActivity(ctx, "aws-batch-CancelJob", input)
	return &BatchCancelJobFuture{Future: future}
}

func (a *stub) CreateComputeEnvironment(ctx workflow.Context, input *batch.CreateComputeEnvironmentInput) (*batch.CreateComputeEnvironmentOutput, error) {
	var output batch.CreateComputeEnvironmentOutput
	err := workflow.ExecuteActivity(ctx, "aws-batch-CreateComputeEnvironment", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateComputeEnvironmentAsync(ctx workflow.Context, input *batch.CreateComputeEnvironmentInput) *BatchCreateComputeEnvironmentFuture {
	future := workflow.ExecuteActivity(ctx, "aws-batch-CreateComputeEnvironment", input)
	return &BatchCreateComputeEnvironmentFuture{Future: future}
}

func (a *stub) CreateJobQueue(ctx workflow.Context, input *batch.CreateJobQueueInput) (*batch.CreateJobQueueOutput, error) {
	var output batch.CreateJobQueueOutput
	err := workflow.ExecuteActivity(ctx, "aws-batch-CreateJobQueue", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateJobQueueAsync(ctx workflow.Context, input *batch.CreateJobQueueInput) *BatchCreateJobQueueFuture {
	future := workflow.ExecuteActivity(ctx, "aws-batch-CreateJobQueue", input)
	return &BatchCreateJobQueueFuture{Future: future}
}

func (a *stub) DeleteComputeEnvironment(ctx workflow.Context, input *batch.DeleteComputeEnvironmentInput) (*batch.DeleteComputeEnvironmentOutput, error) {
	var output batch.DeleteComputeEnvironmentOutput
	err := workflow.ExecuteActivity(ctx, "aws-batch-DeleteComputeEnvironment", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteComputeEnvironmentAsync(ctx workflow.Context, input *batch.DeleteComputeEnvironmentInput) *BatchDeleteComputeEnvironmentFuture {
	future := workflow.ExecuteActivity(ctx, "aws-batch-DeleteComputeEnvironment", input)
	return &BatchDeleteComputeEnvironmentFuture{Future: future}
}

func (a *stub) DeleteJobQueue(ctx workflow.Context, input *batch.DeleteJobQueueInput) (*batch.DeleteJobQueueOutput, error) {
	var output batch.DeleteJobQueueOutput
	err := workflow.ExecuteActivity(ctx, "aws-batch-DeleteJobQueue", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteJobQueueAsync(ctx workflow.Context, input *batch.DeleteJobQueueInput) *BatchDeleteJobQueueFuture {
	future := workflow.ExecuteActivity(ctx, "aws-batch-DeleteJobQueue", input)
	return &BatchDeleteJobQueueFuture{Future: future}
}

func (a *stub) DeregisterJobDefinition(ctx workflow.Context, input *batch.DeregisterJobDefinitionInput) (*batch.DeregisterJobDefinitionOutput, error) {
	var output batch.DeregisterJobDefinitionOutput
	err := workflow.ExecuteActivity(ctx, "aws-batch-DeregisterJobDefinition", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeregisterJobDefinitionAsync(ctx workflow.Context, input *batch.DeregisterJobDefinitionInput) *BatchDeregisterJobDefinitionFuture {
	future := workflow.ExecuteActivity(ctx, "aws-batch-DeregisterJobDefinition", input)
	return &BatchDeregisterJobDefinitionFuture{Future: future}
}

func (a *stub) DescribeComputeEnvironments(ctx workflow.Context, input *batch.DescribeComputeEnvironmentsInput) (*batch.DescribeComputeEnvironmentsOutput, error) {
	var output batch.DescribeComputeEnvironmentsOutput
	err := workflow.ExecuteActivity(ctx, "aws-batch-DescribeComputeEnvironments", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeComputeEnvironmentsAsync(ctx workflow.Context, input *batch.DescribeComputeEnvironmentsInput) *BatchDescribeComputeEnvironmentsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-batch-DescribeComputeEnvironments", input)
	return &BatchDescribeComputeEnvironmentsFuture{Future: future}
}

func (a *stub) DescribeJobDefinitions(ctx workflow.Context, input *batch.DescribeJobDefinitionsInput) (*batch.DescribeJobDefinitionsOutput, error) {
	var output batch.DescribeJobDefinitionsOutput
	err := workflow.ExecuteActivity(ctx, "aws-batch-DescribeJobDefinitions", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeJobDefinitionsAsync(ctx workflow.Context, input *batch.DescribeJobDefinitionsInput) *BatchDescribeJobDefinitionsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-batch-DescribeJobDefinitions", input)
	return &BatchDescribeJobDefinitionsFuture{Future: future}
}

func (a *stub) DescribeJobQueues(ctx workflow.Context, input *batch.DescribeJobQueuesInput) (*batch.DescribeJobQueuesOutput, error) {
	var output batch.DescribeJobQueuesOutput
	err := workflow.ExecuteActivity(ctx, "aws-batch-DescribeJobQueues", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeJobQueuesAsync(ctx workflow.Context, input *batch.DescribeJobQueuesInput) *BatchDescribeJobQueuesFuture {
	future := workflow.ExecuteActivity(ctx, "aws-batch-DescribeJobQueues", input)
	return &BatchDescribeJobQueuesFuture{Future: future}
}

func (a *stub) DescribeJobs(ctx workflow.Context, input *batch.DescribeJobsInput) (*batch.DescribeJobsOutput, error) {
	var output batch.DescribeJobsOutput
	err := workflow.ExecuteActivity(ctx, "aws-batch-DescribeJobs", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeJobsAsync(ctx workflow.Context, input *batch.DescribeJobsInput) *BatchDescribeJobsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-batch-DescribeJobs", input)
	return &BatchDescribeJobsFuture{Future: future}
}

func (a *stub) ListJobs(ctx workflow.Context, input *batch.ListJobsInput) (*batch.ListJobsOutput, error) {
	var output batch.ListJobsOutput
	err := workflow.ExecuteActivity(ctx, "aws-batch-ListJobs", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListJobsAsync(ctx workflow.Context, input *batch.ListJobsInput) *BatchListJobsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-batch-ListJobs", input)
	return &BatchListJobsFuture{Future: future}
}

func (a *stub) ListTagsForResource(ctx workflow.Context, input *batch.ListTagsForResourceInput) (*batch.ListTagsForResourceOutput, error) {
	var output batch.ListTagsForResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws-batch-ListTagsForResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListTagsForResourceAsync(ctx workflow.Context, input *batch.ListTagsForResourceInput) *BatchListTagsForResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws-batch-ListTagsForResource", input)
	return &BatchListTagsForResourceFuture{Future: future}
}

func (a *stub) RegisterJobDefinition(ctx workflow.Context, input *batch.RegisterJobDefinitionInput) (*batch.RegisterJobDefinitionOutput, error) {
	var output batch.RegisterJobDefinitionOutput
	err := workflow.ExecuteActivity(ctx, "aws-batch-RegisterJobDefinition", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) RegisterJobDefinitionAsync(ctx workflow.Context, input *batch.RegisterJobDefinitionInput) *BatchRegisterJobDefinitionFuture {
	future := workflow.ExecuteActivity(ctx, "aws-batch-RegisterJobDefinition", input)
	return &BatchRegisterJobDefinitionFuture{Future: future}
}

func (a *stub) SubmitJob(ctx workflow.Context, input *batch.SubmitJobInput) (*batch.SubmitJobOutput, error) {
	var output batch.SubmitJobOutput
	err := workflow.ExecuteActivity(ctx, "aws-batch-SubmitJob", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) SubmitJobAsync(ctx workflow.Context, input *batch.SubmitJobInput) *BatchSubmitJobFuture {
	future := workflow.ExecuteActivity(ctx, "aws-batch-SubmitJob", input)
	return &BatchSubmitJobFuture{Future: future}
}

func (a *stub) TagResource(ctx workflow.Context, input *batch.TagResourceInput) (*batch.TagResourceOutput, error) {
	var output batch.TagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws-batch-TagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) TagResourceAsync(ctx workflow.Context, input *batch.TagResourceInput) *BatchTagResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws-batch-TagResource", input)
	return &BatchTagResourceFuture{Future: future}
}

func (a *stub) TerminateJob(ctx workflow.Context, input *batch.TerminateJobInput) (*batch.TerminateJobOutput, error) {
	var output batch.TerminateJobOutput
	err := workflow.ExecuteActivity(ctx, "aws-batch-TerminateJob", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) TerminateJobAsync(ctx workflow.Context, input *batch.TerminateJobInput) *BatchTerminateJobFuture {
	future := workflow.ExecuteActivity(ctx, "aws-batch-TerminateJob", input)
	return &BatchTerminateJobFuture{Future: future}
}

func (a *stub) UntagResource(ctx workflow.Context, input *batch.UntagResourceInput) (*batch.UntagResourceOutput, error) {
	var output batch.UntagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws-batch-UntagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UntagResourceAsync(ctx workflow.Context, input *batch.UntagResourceInput) *BatchUntagResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws-batch-UntagResource", input)
	return &BatchUntagResourceFuture{Future: future}
}

func (a *stub) UpdateComputeEnvironment(ctx workflow.Context, input *batch.UpdateComputeEnvironmentInput) (*batch.UpdateComputeEnvironmentOutput, error) {
	var output batch.UpdateComputeEnvironmentOutput
	err := workflow.ExecuteActivity(ctx, "aws-batch-UpdateComputeEnvironment", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateComputeEnvironmentAsync(ctx workflow.Context, input *batch.UpdateComputeEnvironmentInput) *BatchUpdateComputeEnvironmentFuture {
	future := workflow.ExecuteActivity(ctx, "aws-batch-UpdateComputeEnvironment", input)
	return &BatchUpdateComputeEnvironmentFuture{Future: future}
}

func (a *stub) UpdateJobQueue(ctx workflow.Context, input *batch.UpdateJobQueueInput) (*batch.UpdateJobQueueOutput, error) {
	var output batch.UpdateJobQueueOutput
	err := workflow.ExecuteActivity(ctx, "aws-batch-UpdateJobQueue", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateJobQueueAsync(ctx workflow.Context, input *batch.UpdateJobQueueInput) *BatchUpdateJobQueueFuture {
	future := workflow.ExecuteActivity(ctx, "aws-batch-UpdateJobQueue", input)
	return &BatchUpdateJobQueueFuture{Future: future}
}
