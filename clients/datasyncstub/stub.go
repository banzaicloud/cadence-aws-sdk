// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7

package datasyncstub

import (
	"github.com/aws/aws-sdk-go/service/datasync"
	"go.uber.org/cadence/workflow"

	"github.com/banzaicloud/cadence-aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type stub struct{}

type DataSyncCancelTaskExecutionFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DataSyncCancelTaskExecutionFuture) Get(ctx workflow.Context) (*datasync.CancelTaskExecutionOutput, error) {
	var output datasync.CancelTaskExecutionOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DataSyncCreateAgentFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DataSyncCreateAgentFuture) Get(ctx workflow.Context) (*datasync.CreateAgentOutput, error) {
	var output datasync.CreateAgentOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DataSyncCreateLocationEfsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DataSyncCreateLocationEfsFuture) Get(ctx workflow.Context) (*datasync.CreateLocationEfsOutput, error) {
	var output datasync.CreateLocationEfsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DataSyncCreateLocationFsxWindowsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DataSyncCreateLocationFsxWindowsFuture) Get(ctx workflow.Context) (*datasync.CreateLocationFsxWindowsOutput, error) {
	var output datasync.CreateLocationFsxWindowsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DataSyncCreateLocationNfsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DataSyncCreateLocationNfsFuture) Get(ctx workflow.Context) (*datasync.CreateLocationNfsOutput, error) {
	var output datasync.CreateLocationNfsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DataSyncCreateLocationObjectStorageFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DataSyncCreateLocationObjectStorageFuture) Get(ctx workflow.Context) (*datasync.CreateLocationObjectStorageOutput, error) {
	var output datasync.CreateLocationObjectStorageOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DataSyncCreateLocationS3Future struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DataSyncCreateLocationS3Future) Get(ctx workflow.Context) (*datasync.CreateLocationS3Output, error) {
	var output datasync.CreateLocationS3Output
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DataSyncCreateLocationSmbFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DataSyncCreateLocationSmbFuture) Get(ctx workflow.Context) (*datasync.CreateLocationSmbOutput, error) {
	var output datasync.CreateLocationSmbOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DataSyncCreateTaskFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DataSyncCreateTaskFuture) Get(ctx workflow.Context) (*datasync.CreateTaskOutput, error) {
	var output datasync.CreateTaskOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DataSyncDeleteAgentFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DataSyncDeleteAgentFuture) Get(ctx workflow.Context) (*datasync.DeleteAgentOutput, error) {
	var output datasync.DeleteAgentOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DataSyncDeleteLocationFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DataSyncDeleteLocationFuture) Get(ctx workflow.Context) (*datasync.DeleteLocationOutput, error) {
	var output datasync.DeleteLocationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DataSyncDeleteTaskFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DataSyncDeleteTaskFuture) Get(ctx workflow.Context) (*datasync.DeleteTaskOutput, error) {
	var output datasync.DeleteTaskOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DataSyncDescribeAgentFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DataSyncDescribeAgentFuture) Get(ctx workflow.Context) (*datasync.DescribeAgentOutput, error) {
	var output datasync.DescribeAgentOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DataSyncDescribeLocationEfsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DataSyncDescribeLocationEfsFuture) Get(ctx workflow.Context) (*datasync.DescribeLocationEfsOutput, error) {
	var output datasync.DescribeLocationEfsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DataSyncDescribeLocationFsxWindowsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DataSyncDescribeLocationFsxWindowsFuture) Get(ctx workflow.Context) (*datasync.DescribeLocationFsxWindowsOutput, error) {
	var output datasync.DescribeLocationFsxWindowsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DataSyncDescribeLocationNfsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DataSyncDescribeLocationNfsFuture) Get(ctx workflow.Context) (*datasync.DescribeLocationNfsOutput, error) {
	var output datasync.DescribeLocationNfsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DataSyncDescribeLocationObjectStorageFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DataSyncDescribeLocationObjectStorageFuture) Get(ctx workflow.Context) (*datasync.DescribeLocationObjectStorageOutput, error) {
	var output datasync.DescribeLocationObjectStorageOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DataSyncDescribeLocationS3Future struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DataSyncDescribeLocationS3Future) Get(ctx workflow.Context) (*datasync.DescribeLocationS3Output, error) {
	var output datasync.DescribeLocationS3Output
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DataSyncDescribeLocationSmbFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DataSyncDescribeLocationSmbFuture) Get(ctx workflow.Context) (*datasync.DescribeLocationSmbOutput, error) {
	var output datasync.DescribeLocationSmbOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DataSyncDescribeTaskFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DataSyncDescribeTaskFuture) Get(ctx workflow.Context) (*datasync.DescribeTaskOutput, error) {
	var output datasync.DescribeTaskOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DataSyncDescribeTaskExecutionFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DataSyncDescribeTaskExecutionFuture) Get(ctx workflow.Context) (*datasync.DescribeTaskExecutionOutput, error) {
	var output datasync.DescribeTaskExecutionOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DataSyncListAgentsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DataSyncListAgentsFuture) Get(ctx workflow.Context) (*datasync.ListAgentsOutput, error) {
	var output datasync.ListAgentsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DataSyncListLocationsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DataSyncListLocationsFuture) Get(ctx workflow.Context) (*datasync.ListLocationsOutput, error) {
	var output datasync.ListLocationsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DataSyncListTagsForResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DataSyncListTagsForResourceFuture) Get(ctx workflow.Context) (*datasync.ListTagsForResourceOutput, error) {
	var output datasync.ListTagsForResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DataSyncListTaskExecutionsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DataSyncListTaskExecutionsFuture) Get(ctx workflow.Context) (*datasync.ListTaskExecutionsOutput, error) {
	var output datasync.ListTaskExecutionsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DataSyncListTasksFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DataSyncListTasksFuture) Get(ctx workflow.Context) (*datasync.ListTasksOutput, error) {
	var output datasync.ListTasksOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DataSyncStartTaskExecutionFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DataSyncStartTaskExecutionFuture) Get(ctx workflow.Context) (*datasync.StartTaskExecutionOutput, error) {
	var output datasync.StartTaskExecutionOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DataSyncTagResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DataSyncTagResourceFuture) Get(ctx workflow.Context) (*datasync.TagResourceOutput, error) {
	var output datasync.TagResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DataSyncUntagResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DataSyncUntagResourceFuture) Get(ctx workflow.Context) (*datasync.UntagResourceOutput, error) {
	var output datasync.UntagResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DataSyncUpdateAgentFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DataSyncUpdateAgentFuture) Get(ctx workflow.Context) (*datasync.UpdateAgentOutput, error) {
	var output datasync.UpdateAgentOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DataSyncUpdateTaskFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DataSyncUpdateTaskFuture) Get(ctx workflow.Context) (*datasync.UpdateTaskOutput, error) {
	var output datasync.UpdateTaskOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *stub) CancelTaskExecution(ctx workflow.Context, input *datasync.CancelTaskExecutionInput) (*datasync.CancelTaskExecutionOutput, error) {
	var output datasync.CancelTaskExecutionOutput
	err := workflow.ExecuteActivity(ctx, "aws-datasync-CancelTaskExecution", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CancelTaskExecutionAsync(ctx workflow.Context, input *datasync.CancelTaskExecutionInput) *DataSyncCancelTaskExecutionFuture {
	future := workflow.ExecuteActivity(ctx, "aws-datasync-CancelTaskExecution", input)
	return &DataSyncCancelTaskExecutionFuture{Future: future}
}

func (a *stub) CreateAgent(ctx workflow.Context, input *datasync.CreateAgentInput) (*datasync.CreateAgentOutput, error) {
	var output datasync.CreateAgentOutput
	err := workflow.ExecuteActivity(ctx, "aws-datasync-CreateAgent", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateAgentAsync(ctx workflow.Context, input *datasync.CreateAgentInput) *DataSyncCreateAgentFuture {
	future := workflow.ExecuteActivity(ctx, "aws-datasync-CreateAgent", input)
	return &DataSyncCreateAgentFuture{Future: future}
}

func (a *stub) CreateLocationEfs(ctx workflow.Context, input *datasync.CreateLocationEfsInput) (*datasync.CreateLocationEfsOutput, error) {
	var output datasync.CreateLocationEfsOutput
	err := workflow.ExecuteActivity(ctx, "aws-datasync-CreateLocationEfs", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateLocationEfsAsync(ctx workflow.Context, input *datasync.CreateLocationEfsInput) *DataSyncCreateLocationEfsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-datasync-CreateLocationEfs", input)
	return &DataSyncCreateLocationEfsFuture{Future: future}
}

func (a *stub) CreateLocationFsxWindows(ctx workflow.Context, input *datasync.CreateLocationFsxWindowsInput) (*datasync.CreateLocationFsxWindowsOutput, error) {
	var output datasync.CreateLocationFsxWindowsOutput
	err := workflow.ExecuteActivity(ctx, "aws-datasync-CreateLocationFsxWindows", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateLocationFsxWindowsAsync(ctx workflow.Context, input *datasync.CreateLocationFsxWindowsInput) *DataSyncCreateLocationFsxWindowsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-datasync-CreateLocationFsxWindows", input)
	return &DataSyncCreateLocationFsxWindowsFuture{Future: future}
}

func (a *stub) CreateLocationNfs(ctx workflow.Context, input *datasync.CreateLocationNfsInput) (*datasync.CreateLocationNfsOutput, error) {
	var output datasync.CreateLocationNfsOutput
	err := workflow.ExecuteActivity(ctx, "aws-datasync-CreateLocationNfs", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateLocationNfsAsync(ctx workflow.Context, input *datasync.CreateLocationNfsInput) *DataSyncCreateLocationNfsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-datasync-CreateLocationNfs", input)
	return &DataSyncCreateLocationNfsFuture{Future: future}
}

func (a *stub) CreateLocationObjectStorage(ctx workflow.Context, input *datasync.CreateLocationObjectStorageInput) (*datasync.CreateLocationObjectStorageOutput, error) {
	var output datasync.CreateLocationObjectStorageOutput
	err := workflow.ExecuteActivity(ctx, "aws-datasync-CreateLocationObjectStorage", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateLocationObjectStorageAsync(ctx workflow.Context, input *datasync.CreateLocationObjectStorageInput) *DataSyncCreateLocationObjectStorageFuture {
	future := workflow.ExecuteActivity(ctx, "aws-datasync-CreateLocationObjectStorage", input)
	return &DataSyncCreateLocationObjectStorageFuture{Future: future}
}

func (a *stub) CreateLocationS3(ctx workflow.Context, input *datasync.CreateLocationS3Input) (*datasync.CreateLocationS3Output, error) {
	var output datasync.CreateLocationS3Output
	err := workflow.ExecuteActivity(ctx, "aws-datasync-CreateLocationS3", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateLocationS3Async(ctx workflow.Context, input *datasync.CreateLocationS3Input) *DataSyncCreateLocationS3Future {
	future := workflow.ExecuteActivity(ctx, "aws-datasync-CreateLocationS3", input)
	return &DataSyncCreateLocationS3Future{Future: future}
}

func (a *stub) CreateLocationSmb(ctx workflow.Context, input *datasync.CreateLocationSmbInput) (*datasync.CreateLocationSmbOutput, error) {
	var output datasync.CreateLocationSmbOutput
	err := workflow.ExecuteActivity(ctx, "aws-datasync-CreateLocationSmb", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateLocationSmbAsync(ctx workflow.Context, input *datasync.CreateLocationSmbInput) *DataSyncCreateLocationSmbFuture {
	future := workflow.ExecuteActivity(ctx, "aws-datasync-CreateLocationSmb", input)
	return &DataSyncCreateLocationSmbFuture{Future: future}
}

func (a *stub) CreateTask(ctx workflow.Context, input *datasync.CreateTaskInput) (*datasync.CreateTaskOutput, error) {
	var output datasync.CreateTaskOutput
	err := workflow.ExecuteActivity(ctx, "aws-datasync-CreateTask", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateTaskAsync(ctx workflow.Context, input *datasync.CreateTaskInput) *DataSyncCreateTaskFuture {
	future := workflow.ExecuteActivity(ctx, "aws-datasync-CreateTask", input)
	return &DataSyncCreateTaskFuture{Future: future}
}

func (a *stub) DeleteAgent(ctx workflow.Context, input *datasync.DeleteAgentInput) (*datasync.DeleteAgentOutput, error) {
	var output datasync.DeleteAgentOutput
	err := workflow.ExecuteActivity(ctx, "aws-datasync-DeleteAgent", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteAgentAsync(ctx workflow.Context, input *datasync.DeleteAgentInput) *DataSyncDeleteAgentFuture {
	future := workflow.ExecuteActivity(ctx, "aws-datasync-DeleteAgent", input)
	return &DataSyncDeleteAgentFuture{Future: future}
}

func (a *stub) DeleteLocation(ctx workflow.Context, input *datasync.DeleteLocationInput) (*datasync.DeleteLocationOutput, error) {
	var output datasync.DeleteLocationOutput
	err := workflow.ExecuteActivity(ctx, "aws-datasync-DeleteLocation", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteLocationAsync(ctx workflow.Context, input *datasync.DeleteLocationInput) *DataSyncDeleteLocationFuture {
	future := workflow.ExecuteActivity(ctx, "aws-datasync-DeleteLocation", input)
	return &DataSyncDeleteLocationFuture{Future: future}
}

func (a *stub) DeleteTask(ctx workflow.Context, input *datasync.DeleteTaskInput) (*datasync.DeleteTaskOutput, error) {
	var output datasync.DeleteTaskOutput
	err := workflow.ExecuteActivity(ctx, "aws-datasync-DeleteTask", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteTaskAsync(ctx workflow.Context, input *datasync.DeleteTaskInput) *DataSyncDeleteTaskFuture {
	future := workflow.ExecuteActivity(ctx, "aws-datasync-DeleteTask", input)
	return &DataSyncDeleteTaskFuture{Future: future}
}

func (a *stub) DescribeAgent(ctx workflow.Context, input *datasync.DescribeAgentInput) (*datasync.DescribeAgentOutput, error) {
	var output datasync.DescribeAgentOutput
	err := workflow.ExecuteActivity(ctx, "aws-datasync-DescribeAgent", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeAgentAsync(ctx workflow.Context, input *datasync.DescribeAgentInput) *DataSyncDescribeAgentFuture {
	future := workflow.ExecuteActivity(ctx, "aws-datasync-DescribeAgent", input)
	return &DataSyncDescribeAgentFuture{Future: future}
}

func (a *stub) DescribeLocationEfs(ctx workflow.Context, input *datasync.DescribeLocationEfsInput) (*datasync.DescribeLocationEfsOutput, error) {
	var output datasync.DescribeLocationEfsOutput
	err := workflow.ExecuteActivity(ctx, "aws-datasync-DescribeLocationEfs", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeLocationEfsAsync(ctx workflow.Context, input *datasync.DescribeLocationEfsInput) *DataSyncDescribeLocationEfsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-datasync-DescribeLocationEfs", input)
	return &DataSyncDescribeLocationEfsFuture{Future: future}
}

func (a *stub) DescribeLocationFsxWindows(ctx workflow.Context, input *datasync.DescribeLocationFsxWindowsInput) (*datasync.DescribeLocationFsxWindowsOutput, error) {
	var output datasync.DescribeLocationFsxWindowsOutput
	err := workflow.ExecuteActivity(ctx, "aws-datasync-DescribeLocationFsxWindows", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeLocationFsxWindowsAsync(ctx workflow.Context, input *datasync.DescribeLocationFsxWindowsInput) *DataSyncDescribeLocationFsxWindowsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-datasync-DescribeLocationFsxWindows", input)
	return &DataSyncDescribeLocationFsxWindowsFuture{Future: future}
}

func (a *stub) DescribeLocationNfs(ctx workflow.Context, input *datasync.DescribeLocationNfsInput) (*datasync.DescribeLocationNfsOutput, error) {
	var output datasync.DescribeLocationNfsOutput
	err := workflow.ExecuteActivity(ctx, "aws-datasync-DescribeLocationNfs", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeLocationNfsAsync(ctx workflow.Context, input *datasync.DescribeLocationNfsInput) *DataSyncDescribeLocationNfsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-datasync-DescribeLocationNfs", input)
	return &DataSyncDescribeLocationNfsFuture{Future: future}
}

func (a *stub) DescribeLocationObjectStorage(ctx workflow.Context, input *datasync.DescribeLocationObjectStorageInput) (*datasync.DescribeLocationObjectStorageOutput, error) {
	var output datasync.DescribeLocationObjectStorageOutput
	err := workflow.ExecuteActivity(ctx, "aws-datasync-DescribeLocationObjectStorage", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeLocationObjectStorageAsync(ctx workflow.Context, input *datasync.DescribeLocationObjectStorageInput) *DataSyncDescribeLocationObjectStorageFuture {
	future := workflow.ExecuteActivity(ctx, "aws-datasync-DescribeLocationObjectStorage", input)
	return &DataSyncDescribeLocationObjectStorageFuture{Future: future}
}

func (a *stub) DescribeLocationS3(ctx workflow.Context, input *datasync.DescribeLocationS3Input) (*datasync.DescribeLocationS3Output, error) {
	var output datasync.DescribeLocationS3Output
	err := workflow.ExecuteActivity(ctx, "aws-datasync-DescribeLocationS3", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeLocationS3Async(ctx workflow.Context, input *datasync.DescribeLocationS3Input) *DataSyncDescribeLocationS3Future {
	future := workflow.ExecuteActivity(ctx, "aws-datasync-DescribeLocationS3", input)
	return &DataSyncDescribeLocationS3Future{Future: future}
}

func (a *stub) DescribeLocationSmb(ctx workflow.Context, input *datasync.DescribeLocationSmbInput) (*datasync.DescribeLocationSmbOutput, error) {
	var output datasync.DescribeLocationSmbOutput
	err := workflow.ExecuteActivity(ctx, "aws-datasync-DescribeLocationSmb", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeLocationSmbAsync(ctx workflow.Context, input *datasync.DescribeLocationSmbInput) *DataSyncDescribeLocationSmbFuture {
	future := workflow.ExecuteActivity(ctx, "aws-datasync-DescribeLocationSmb", input)
	return &DataSyncDescribeLocationSmbFuture{Future: future}
}

func (a *stub) DescribeTask(ctx workflow.Context, input *datasync.DescribeTaskInput) (*datasync.DescribeTaskOutput, error) {
	var output datasync.DescribeTaskOutput
	err := workflow.ExecuteActivity(ctx, "aws-datasync-DescribeTask", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeTaskAsync(ctx workflow.Context, input *datasync.DescribeTaskInput) *DataSyncDescribeTaskFuture {
	future := workflow.ExecuteActivity(ctx, "aws-datasync-DescribeTask", input)
	return &DataSyncDescribeTaskFuture{Future: future}
}

func (a *stub) DescribeTaskExecution(ctx workflow.Context, input *datasync.DescribeTaskExecutionInput) (*datasync.DescribeTaskExecutionOutput, error) {
	var output datasync.DescribeTaskExecutionOutput
	err := workflow.ExecuteActivity(ctx, "aws-datasync-DescribeTaskExecution", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeTaskExecutionAsync(ctx workflow.Context, input *datasync.DescribeTaskExecutionInput) *DataSyncDescribeTaskExecutionFuture {
	future := workflow.ExecuteActivity(ctx, "aws-datasync-DescribeTaskExecution", input)
	return &DataSyncDescribeTaskExecutionFuture{Future: future}
}

func (a *stub) ListAgents(ctx workflow.Context, input *datasync.ListAgentsInput) (*datasync.ListAgentsOutput, error) {
	var output datasync.ListAgentsOutput
	err := workflow.ExecuteActivity(ctx, "aws-datasync-ListAgents", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListAgentsAsync(ctx workflow.Context, input *datasync.ListAgentsInput) *DataSyncListAgentsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-datasync-ListAgents", input)
	return &DataSyncListAgentsFuture{Future: future}
}

func (a *stub) ListLocations(ctx workflow.Context, input *datasync.ListLocationsInput) (*datasync.ListLocationsOutput, error) {
	var output datasync.ListLocationsOutput
	err := workflow.ExecuteActivity(ctx, "aws-datasync-ListLocations", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListLocationsAsync(ctx workflow.Context, input *datasync.ListLocationsInput) *DataSyncListLocationsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-datasync-ListLocations", input)
	return &DataSyncListLocationsFuture{Future: future}
}

func (a *stub) ListTagsForResource(ctx workflow.Context, input *datasync.ListTagsForResourceInput) (*datasync.ListTagsForResourceOutput, error) {
	var output datasync.ListTagsForResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws-datasync-ListTagsForResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListTagsForResourceAsync(ctx workflow.Context, input *datasync.ListTagsForResourceInput) *DataSyncListTagsForResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws-datasync-ListTagsForResource", input)
	return &DataSyncListTagsForResourceFuture{Future: future}
}

func (a *stub) ListTaskExecutions(ctx workflow.Context, input *datasync.ListTaskExecutionsInput) (*datasync.ListTaskExecutionsOutput, error) {
	var output datasync.ListTaskExecutionsOutput
	err := workflow.ExecuteActivity(ctx, "aws-datasync-ListTaskExecutions", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListTaskExecutionsAsync(ctx workflow.Context, input *datasync.ListTaskExecutionsInput) *DataSyncListTaskExecutionsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-datasync-ListTaskExecutions", input)
	return &DataSyncListTaskExecutionsFuture{Future: future}
}

func (a *stub) ListTasks(ctx workflow.Context, input *datasync.ListTasksInput) (*datasync.ListTasksOutput, error) {
	var output datasync.ListTasksOutput
	err := workflow.ExecuteActivity(ctx, "aws-datasync-ListTasks", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListTasksAsync(ctx workflow.Context, input *datasync.ListTasksInput) *DataSyncListTasksFuture {
	future := workflow.ExecuteActivity(ctx, "aws-datasync-ListTasks", input)
	return &DataSyncListTasksFuture{Future: future}
}

func (a *stub) StartTaskExecution(ctx workflow.Context, input *datasync.StartTaskExecutionInput) (*datasync.StartTaskExecutionOutput, error) {
	var output datasync.StartTaskExecutionOutput
	err := workflow.ExecuteActivity(ctx, "aws-datasync-StartTaskExecution", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) StartTaskExecutionAsync(ctx workflow.Context, input *datasync.StartTaskExecutionInput) *DataSyncStartTaskExecutionFuture {
	future := workflow.ExecuteActivity(ctx, "aws-datasync-StartTaskExecution", input)
	return &DataSyncStartTaskExecutionFuture{Future: future}
}

func (a *stub) TagResource(ctx workflow.Context, input *datasync.TagResourceInput) (*datasync.TagResourceOutput, error) {
	var output datasync.TagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws-datasync-TagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) TagResourceAsync(ctx workflow.Context, input *datasync.TagResourceInput) *DataSyncTagResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws-datasync-TagResource", input)
	return &DataSyncTagResourceFuture{Future: future}
}

func (a *stub) UntagResource(ctx workflow.Context, input *datasync.UntagResourceInput) (*datasync.UntagResourceOutput, error) {
	var output datasync.UntagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws-datasync-UntagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UntagResourceAsync(ctx workflow.Context, input *datasync.UntagResourceInput) *DataSyncUntagResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws-datasync-UntagResource", input)
	return &DataSyncUntagResourceFuture{Future: future}
}

func (a *stub) UpdateAgent(ctx workflow.Context, input *datasync.UpdateAgentInput) (*datasync.UpdateAgentOutput, error) {
	var output datasync.UpdateAgentOutput
	err := workflow.ExecuteActivity(ctx, "aws-datasync-UpdateAgent", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateAgentAsync(ctx workflow.Context, input *datasync.UpdateAgentInput) *DataSyncUpdateAgentFuture {
	future := workflow.ExecuteActivity(ctx, "aws-datasync-UpdateAgent", input)
	return &DataSyncUpdateAgentFuture{Future: future}
}

func (a *stub) UpdateTask(ctx workflow.Context, input *datasync.UpdateTaskInput) (*datasync.UpdateTaskOutput, error) {
	var output datasync.UpdateTaskOutput
	err := workflow.ExecuteActivity(ctx, "aws-datasync-UpdateTask", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateTaskAsync(ctx workflow.Context, input *datasync.UpdateTaskInput) *DataSyncUpdateTaskFuture {
	future := workflow.ExecuteActivity(ctx, "aws-datasync-UpdateTask", input)
	return &DataSyncUpdateTaskFuture{Future: future}
}
