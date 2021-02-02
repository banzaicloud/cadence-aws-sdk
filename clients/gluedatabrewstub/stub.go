// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7

package gluedatabrewstub

import (
	"github.com/aws/aws-sdk-go/service/gluedatabrew"
	"go.uber.org/cadence/workflow"

	"github.com/banzaicloud/cadence-aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type stub struct{}

type BatchDeleteRecipeVersionFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *BatchDeleteRecipeVersionFuture) Get(ctx workflow.Context) (*gluedatabrew.BatchDeleteRecipeVersionOutput, error) {
	var output gluedatabrew.BatchDeleteRecipeVersionOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CreateDatasetFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CreateDatasetFuture) Get(ctx workflow.Context) (*gluedatabrew.CreateDatasetOutput, error) {
	var output gluedatabrew.CreateDatasetOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CreateProfileJobFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CreateProfileJobFuture) Get(ctx workflow.Context) (*gluedatabrew.CreateProfileJobOutput, error) {
	var output gluedatabrew.CreateProfileJobOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CreateProjectFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CreateProjectFuture) Get(ctx workflow.Context) (*gluedatabrew.CreateProjectOutput, error) {
	var output gluedatabrew.CreateProjectOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CreateRecipeFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CreateRecipeFuture) Get(ctx workflow.Context) (*gluedatabrew.CreateRecipeOutput, error) {
	var output gluedatabrew.CreateRecipeOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CreateRecipeJobFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CreateRecipeJobFuture) Get(ctx workflow.Context) (*gluedatabrew.CreateRecipeJobOutput, error) {
	var output gluedatabrew.CreateRecipeJobOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CreateScheduleFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CreateScheduleFuture) Get(ctx workflow.Context) (*gluedatabrew.CreateScheduleOutput, error) {
	var output gluedatabrew.CreateScheduleOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DeleteDatasetFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DeleteDatasetFuture) Get(ctx workflow.Context) (*gluedatabrew.DeleteDatasetOutput, error) {
	var output gluedatabrew.DeleteDatasetOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DeleteJobFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DeleteJobFuture) Get(ctx workflow.Context) (*gluedatabrew.DeleteJobOutput, error) {
	var output gluedatabrew.DeleteJobOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DeleteProjectFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DeleteProjectFuture) Get(ctx workflow.Context) (*gluedatabrew.DeleteProjectOutput, error) {
	var output gluedatabrew.DeleteProjectOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DeleteRecipeVersionFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DeleteRecipeVersionFuture) Get(ctx workflow.Context) (*gluedatabrew.DeleteRecipeVersionOutput, error) {
	var output gluedatabrew.DeleteRecipeVersionOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DeleteScheduleFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DeleteScheduleFuture) Get(ctx workflow.Context) (*gluedatabrew.DeleteScheduleOutput, error) {
	var output gluedatabrew.DeleteScheduleOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DescribeDatasetFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DescribeDatasetFuture) Get(ctx workflow.Context) (*gluedatabrew.DescribeDatasetOutput, error) {
	var output gluedatabrew.DescribeDatasetOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DescribeJobFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DescribeJobFuture) Get(ctx workflow.Context) (*gluedatabrew.DescribeJobOutput, error) {
	var output gluedatabrew.DescribeJobOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DescribeProjectFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DescribeProjectFuture) Get(ctx workflow.Context) (*gluedatabrew.DescribeProjectOutput, error) {
	var output gluedatabrew.DescribeProjectOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DescribeRecipeFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DescribeRecipeFuture) Get(ctx workflow.Context) (*gluedatabrew.DescribeRecipeOutput, error) {
	var output gluedatabrew.DescribeRecipeOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DescribeScheduleFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DescribeScheduleFuture) Get(ctx workflow.Context) (*gluedatabrew.DescribeScheduleOutput, error) {
	var output gluedatabrew.DescribeScheduleOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListDatasetsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListDatasetsFuture) Get(ctx workflow.Context) (*gluedatabrew.ListDatasetsOutput, error) {
	var output gluedatabrew.ListDatasetsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListJobRunsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListJobRunsFuture) Get(ctx workflow.Context) (*gluedatabrew.ListJobRunsOutput, error) {
	var output gluedatabrew.ListJobRunsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListJobsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListJobsFuture) Get(ctx workflow.Context) (*gluedatabrew.ListJobsOutput, error) {
	var output gluedatabrew.ListJobsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListProjectsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListProjectsFuture) Get(ctx workflow.Context) (*gluedatabrew.ListProjectsOutput, error) {
	var output gluedatabrew.ListProjectsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListRecipeVersionsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListRecipeVersionsFuture) Get(ctx workflow.Context) (*gluedatabrew.ListRecipeVersionsOutput, error) {
	var output gluedatabrew.ListRecipeVersionsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListRecipesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListRecipesFuture) Get(ctx workflow.Context) (*gluedatabrew.ListRecipesOutput, error) {
	var output gluedatabrew.ListRecipesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListSchedulesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListSchedulesFuture) Get(ctx workflow.Context) (*gluedatabrew.ListSchedulesOutput, error) {
	var output gluedatabrew.ListSchedulesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListTagsForResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListTagsForResourceFuture) Get(ctx workflow.Context) (*gluedatabrew.ListTagsForResourceOutput, error) {
	var output gluedatabrew.ListTagsForResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type PublishRecipeFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *PublishRecipeFuture) Get(ctx workflow.Context) (*gluedatabrew.PublishRecipeOutput, error) {
	var output gluedatabrew.PublishRecipeOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SendProjectSessionActionFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SendProjectSessionActionFuture) Get(ctx workflow.Context) (*gluedatabrew.SendProjectSessionActionOutput, error) {
	var output gluedatabrew.SendProjectSessionActionOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type StartJobRunFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *StartJobRunFuture) Get(ctx workflow.Context) (*gluedatabrew.StartJobRunOutput, error) {
	var output gluedatabrew.StartJobRunOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type StartProjectSessionFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *StartProjectSessionFuture) Get(ctx workflow.Context) (*gluedatabrew.StartProjectSessionOutput, error) {
	var output gluedatabrew.StartProjectSessionOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type StopJobRunFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *StopJobRunFuture) Get(ctx workflow.Context) (*gluedatabrew.StopJobRunOutput, error) {
	var output gluedatabrew.StopJobRunOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type TagResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *TagResourceFuture) Get(ctx workflow.Context) (*gluedatabrew.TagResourceOutput, error) {
	var output gluedatabrew.TagResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type UntagResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *UntagResourceFuture) Get(ctx workflow.Context) (*gluedatabrew.UntagResourceOutput, error) {
	var output gluedatabrew.UntagResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type UpdateDatasetFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *UpdateDatasetFuture) Get(ctx workflow.Context) (*gluedatabrew.UpdateDatasetOutput, error) {
	var output gluedatabrew.UpdateDatasetOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type UpdateProfileJobFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *UpdateProfileJobFuture) Get(ctx workflow.Context) (*gluedatabrew.UpdateProfileJobOutput, error) {
	var output gluedatabrew.UpdateProfileJobOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type UpdateProjectFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *UpdateProjectFuture) Get(ctx workflow.Context) (*gluedatabrew.UpdateProjectOutput, error) {
	var output gluedatabrew.UpdateProjectOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type UpdateRecipeFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *UpdateRecipeFuture) Get(ctx workflow.Context) (*gluedatabrew.UpdateRecipeOutput, error) {
	var output gluedatabrew.UpdateRecipeOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type UpdateRecipeJobFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *UpdateRecipeJobFuture) Get(ctx workflow.Context) (*gluedatabrew.UpdateRecipeJobOutput, error) {
	var output gluedatabrew.UpdateRecipeJobOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type UpdateScheduleFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *UpdateScheduleFuture) Get(ctx workflow.Context) (*gluedatabrew.UpdateScheduleOutput, error) {
	var output gluedatabrew.UpdateScheduleOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *stub) BatchDeleteRecipeVersion(ctx workflow.Context, input *gluedatabrew.BatchDeleteRecipeVersionInput) (*gluedatabrew.BatchDeleteRecipeVersionOutput, error) {
	var output gluedatabrew.BatchDeleteRecipeVersionOutput
	err := workflow.ExecuteActivity(ctx, "aws-gluedatabrew-BatchDeleteRecipeVersion", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) BatchDeleteRecipeVersionAsync(ctx workflow.Context, input *gluedatabrew.BatchDeleteRecipeVersionInput) *BatchDeleteRecipeVersionFuture {
	future := workflow.ExecuteActivity(ctx, "aws-gluedatabrew-BatchDeleteRecipeVersion", input)
	return &BatchDeleteRecipeVersionFuture{Future: future}
}

func (a *stub) CreateDataset(ctx workflow.Context, input *gluedatabrew.CreateDatasetInput) (*gluedatabrew.CreateDatasetOutput, error) {
	var output gluedatabrew.CreateDatasetOutput
	err := workflow.ExecuteActivity(ctx, "aws-gluedatabrew-CreateDataset", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateDatasetAsync(ctx workflow.Context, input *gluedatabrew.CreateDatasetInput) *CreateDatasetFuture {
	future := workflow.ExecuteActivity(ctx, "aws-gluedatabrew-CreateDataset", input)
	return &CreateDatasetFuture{Future: future}
}

func (a *stub) CreateProfileJob(ctx workflow.Context, input *gluedatabrew.CreateProfileJobInput) (*gluedatabrew.CreateProfileJobOutput, error) {
	var output gluedatabrew.CreateProfileJobOutput
	err := workflow.ExecuteActivity(ctx, "aws-gluedatabrew-CreateProfileJob", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateProfileJobAsync(ctx workflow.Context, input *gluedatabrew.CreateProfileJobInput) *CreateProfileJobFuture {
	future := workflow.ExecuteActivity(ctx, "aws-gluedatabrew-CreateProfileJob", input)
	return &CreateProfileJobFuture{Future: future}
}

func (a *stub) CreateProject(ctx workflow.Context, input *gluedatabrew.CreateProjectInput) (*gluedatabrew.CreateProjectOutput, error) {
	var output gluedatabrew.CreateProjectOutput
	err := workflow.ExecuteActivity(ctx, "aws-gluedatabrew-CreateProject", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateProjectAsync(ctx workflow.Context, input *gluedatabrew.CreateProjectInput) *CreateProjectFuture {
	future := workflow.ExecuteActivity(ctx, "aws-gluedatabrew-CreateProject", input)
	return &CreateProjectFuture{Future: future}
}

func (a *stub) CreateRecipe(ctx workflow.Context, input *gluedatabrew.CreateRecipeInput) (*gluedatabrew.CreateRecipeOutput, error) {
	var output gluedatabrew.CreateRecipeOutput
	err := workflow.ExecuteActivity(ctx, "aws-gluedatabrew-CreateRecipe", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateRecipeAsync(ctx workflow.Context, input *gluedatabrew.CreateRecipeInput) *CreateRecipeFuture {
	future := workflow.ExecuteActivity(ctx, "aws-gluedatabrew-CreateRecipe", input)
	return &CreateRecipeFuture{Future: future}
}

func (a *stub) CreateRecipeJob(ctx workflow.Context, input *gluedatabrew.CreateRecipeJobInput) (*gluedatabrew.CreateRecipeJobOutput, error) {
	var output gluedatabrew.CreateRecipeJobOutput
	err := workflow.ExecuteActivity(ctx, "aws-gluedatabrew-CreateRecipeJob", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateRecipeJobAsync(ctx workflow.Context, input *gluedatabrew.CreateRecipeJobInput) *CreateRecipeJobFuture {
	future := workflow.ExecuteActivity(ctx, "aws-gluedatabrew-CreateRecipeJob", input)
	return &CreateRecipeJobFuture{Future: future}
}

func (a *stub) CreateSchedule(ctx workflow.Context, input *gluedatabrew.CreateScheduleInput) (*gluedatabrew.CreateScheduleOutput, error) {
	var output gluedatabrew.CreateScheduleOutput
	err := workflow.ExecuteActivity(ctx, "aws-gluedatabrew-CreateSchedule", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateScheduleAsync(ctx workflow.Context, input *gluedatabrew.CreateScheduleInput) *CreateScheduleFuture {
	future := workflow.ExecuteActivity(ctx, "aws-gluedatabrew-CreateSchedule", input)
	return &CreateScheduleFuture{Future: future}
}

func (a *stub) DeleteDataset(ctx workflow.Context, input *gluedatabrew.DeleteDatasetInput) (*gluedatabrew.DeleteDatasetOutput, error) {
	var output gluedatabrew.DeleteDatasetOutput
	err := workflow.ExecuteActivity(ctx, "aws-gluedatabrew-DeleteDataset", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteDatasetAsync(ctx workflow.Context, input *gluedatabrew.DeleteDatasetInput) *DeleteDatasetFuture {
	future := workflow.ExecuteActivity(ctx, "aws-gluedatabrew-DeleteDataset", input)
	return &DeleteDatasetFuture{Future: future}
}

func (a *stub) DeleteJob(ctx workflow.Context, input *gluedatabrew.DeleteJobInput) (*gluedatabrew.DeleteJobOutput, error) {
	var output gluedatabrew.DeleteJobOutput
	err := workflow.ExecuteActivity(ctx, "aws-gluedatabrew-DeleteJob", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteJobAsync(ctx workflow.Context, input *gluedatabrew.DeleteJobInput) *DeleteJobFuture {
	future := workflow.ExecuteActivity(ctx, "aws-gluedatabrew-DeleteJob", input)
	return &DeleteJobFuture{Future: future}
}

func (a *stub) DeleteProject(ctx workflow.Context, input *gluedatabrew.DeleteProjectInput) (*gluedatabrew.DeleteProjectOutput, error) {
	var output gluedatabrew.DeleteProjectOutput
	err := workflow.ExecuteActivity(ctx, "aws-gluedatabrew-DeleteProject", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteProjectAsync(ctx workflow.Context, input *gluedatabrew.DeleteProjectInput) *DeleteProjectFuture {
	future := workflow.ExecuteActivity(ctx, "aws-gluedatabrew-DeleteProject", input)
	return &DeleteProjectFuture{Future: future}
}

func (a *stub) DeleteRecipeVersion(ctx workflow.Context, input *gluedatabrew.DeleteRecipeVersionInput) (*gluedatabrew.DeleteRecipeVersionOutput, error) {
	var output gluedatabrew.DeleteRecipeVersionOutput
	err := workflow.ExecuteActivity(ctx, "aws-gluedatabrew-DeleteRecipeVersion", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteRecipeVersionAsync(ctx workflow.Context, input *gluedatabrew.DeleteRecipeVersionInput) *DeleteRecipeVersionFuture {
	future := workflow.ExecuteActivity(ctx, "aws-gluedatabrew-DeleteRecipeVersion", input)
	return &DeleteRecipeVersionFuture{Future: future}
}

func (a *stub) DeleteSchedule(ctx workflow.Context, input *gluedatabrew.DeleteScheduleInput) (*gluedatabrew.DeleteScheduleOutput, error) {
	var output gluedatabrew.DeleteScheduleOutput
	err := workflow.ExecuteActivity(ctx, "aws-gluedatabrew-DeleteSchedule", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteScheduleAsync(ctx workflow.Context, input *gluedatabrew.DeleteScheduleInput) *DeleteScheduleFuture {
	future := workflow.ExecuteActivity(ctx, "aws-gluedatabrew-DeleteSchedule", input)
	return &DeleteScheduleFuture{Future: future}
}

func (a *stub) DescribeDataset(ctx workflow.Context, input *gluedatabrew.DescribeDatasetInput) (*gluedatabrew.DescribeDatasetOutput, error) {
	var output gluedatabrew.DescribeDatasetOutput
	err := workflow.ExecuteActivity(ctx, "aws-gluedatabrew-DescribeDataset", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeDatasetAsync(ctx workflow.Context, input *gluedatabrew.DescribeDatasetInput) *DescribeDatasetFuture {
	future := workflow.ExecuteActivity(ctx, "aws-gluedatabrew-DescribeDataset", input)
	return &DescribeDatasetFuture{Future: future}
}

func (a *stub) DescribeJob(ctx workflow.Context, input *gluedatabrew.DescribeJobInput) (*gluedatabrew.DescribeJobOutput, error) {
	var output gluedatabrew.DescribeJobOutput
	err := workflow.ExecuteActivity(ctx, "aws-gluedatabrew-DescribeJob", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeJobAsync(ctx workflow.Context, input *gluedatabrew.DescribeJobInput) *DescribeJobFuture {
	future := workflow.ExecuteActivity(ctx, "aws-gluedatabrew-DescribeJob", input)
	return &DescribeJobFuture{Future: future}
}

func (a *stub) DescribeProject(ctx workflow.Context, input *gluedatabrew.DescribeProjectInput) (*gluedatabrew.DescribeProjectOutput, error) {
	var output gluedatabrew.DescribeProjectOutput
	err := workflow.ExecuteActivity(ctx, "aws-gluedatabrew-DescribeProject", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeProjectAsync(ctx workflow.Context, input *gluedatabrew.DescribeProjectInput) *DescribeProjectFuture {
	future := workflow.ExecuteActivity(ctx, "aws-gluedatabrew-DescribeProject", input)
	return &DescribeProjectFuture{Future: future}
}

func (a *stub) DescribeRecipe(ctx workflow.Context, input *gluedatabrew.DescribeRecipeInput) (*gluedatabrew.DescribeRecipeOutput, error) {
	var output gluedatabrew.DescribeRecipeOutput
	err := workflow.ExecuteActivity(ctx, "aws-gluedatabrew-DescribeRecipe", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeRecipeAsync(ctx workflow.Context, input *gluedatabrew.DescribeRecipeInput) *DescribeRecipeFuture {
	future := workflow.ExecuteActivity(ctx, "aws-gluedatabrew-DescribeRecipe", input)
	return &DescribeRecipeFuture{Future: future}
}

func (a *stub) DescribeSchedule(ctx workflow.Context, input *gluedatabrew.DescribeScheduleInput) (*gluedatabrew.DescribeScheduleOutput, error) {
	var output gluedatabrew.DescribeScheduleOutput
	err := workflow.ExecuteActivity(ctx, "aws-gluedatabrew-DescribeSchedule", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeScheduleAsync(ctx workflow.Context, input *gluedatabrew.DescribeScheduleInput) *DescribeScheduleFuture {
	future := workflow.ExecuteActivity(ctx, "aws-gluedatabrew-DescribeSchedule", input)
	return &DescribeScheduleFuture{Future: future}
}

func (a *stub) ListDatasets(ctx workflow.Context, input *gluedatabrew.ListDatasetsInput) (*gluedatabrew.ListDatasetsOutput, error) {
	var output gluedatabrew.ListDatasetsOutput
	err := workflow.ExecuteActivity(ctx, "aws-gluedatabrew-ListDatasets", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListDatasetsAsync(ctx workflow.Context, input *gluedatabrew.ListDatasetsInput) *ListDatasetsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-gluedatabrew-ListDatasets", input)
	return &ListDatasetsFuture{Future: future}
}

func (a *stub) ListJobRuns(ctx workflow.Context, input *gluedatabrew.ListJobRunsInput) (*gluedatabrew.ListJobRunsOutput, error) {
	var output gluedatabrew.ListJobRunsOutput
	err := workflow.ExecuteActivity(ctx, "aws-gluedatabrew-ListJobRuns", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListJobRunsAsync(ctx workflow.Context, input *gluedatabrew.ListJobRunsInput) *ListJobRunsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-gluedatabrew-ListJobRuns", input)
	return &ListJobRunsFuture{Future: future}
}

func (a *stub) ListJobs(ctx workflow.Context, input *gluedatabrew.ListJobsInput) (*gluedatabrew.ListJobsOutput, error) {
	var output gluedatabrew.ListJobsOutput
	err := workflow.ExecuteActivity(ctx, "aws-gluedatabrew-ListJobs", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListJobsAsync(ctx workflow.Context, input *gluedatabrew.ListJobsInput) *ListJobsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-gluedatabrew-ListJobs", input)
	return &ListJobsFuture{Future: future}
}

func (a *stub) ListProjects(ctx workflow.Context, input *gluedatabrew.ListProjectsInput) (*gluedatabrew.ListProjectsOutput, error) {
	var output gluedatabrew.ListProjectsOutput
	err := workflow.ExecuteActivity(ctx, "aws-gluedatabrew-ListProjects", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListProjectsAsync(ctx workflow.Context, input *gluedatabrew.ListProjectsInput) *ListProjectsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-gluedatabrew-ListProjects", input)
	return &ListProjectsFuture{Future: future}
}

func (a *stub) ListRecipeVersions(ctx workflow.Context, input *gluedatabrew.ListRecipeVersionsInput) (*gluedatabrew.ListRecipeVersionsOutput, error) {
	var output gluedatabrew.ListRecipeVersionsOutput
	err := workflow.ExecuteActivity(ctx, "aws-gluedatabrew-ListRecipeVersions", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListRecipeVersionsAsync(ctx workflow.Context, input *gluedatabrew.ListRecipeVersionsInput) *ListRecipeVersionsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-gluedatabrew-ListRecipeVersions", input)
	return &ListRecipeVersionsFuture{Future: future}
}

func (a *stub) ListRecipes(ctx workflow.Context, input *gluedatabrew.ListRecipesInput) (*gluedatabrew.ListRecipesOutput, error) {
	var output gluedatabrew.ListRecipesOutput
	err := workflow.ExecuteActivity(ctx, "aws-gluedatabrew-ListRecipes", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListRecipesAsync(ctx workflow.Context, input *gluedatabrew.ListRecipesInput) *ListRecipesFuture {
	future := workflow.ExecuteActivity(ctx, "aws-gluedatabrew-ListRecipes", input)
	return &ListRecipesFuture{Future: future}
}

func (a *stub) ListSchedules(ctx workflow.Context, input *gluedatabrew.ListSchedulesInput) (*gluedatabrew.ListSchedulesOutput, error) {
	var output gluedatabrew.ListSchedulesOutput
	err := workflow.ExecuteActivity(ctx, "aws-gluedatabrew-ListSchedules", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListSchedulesAsync(ctx workflow.Context, input *gluedatabrew.ListSchedulesInput) *ListSchedulesFuture {
	future := workflow.ExecuteActivity(ctx, "aws-gluedatabrew-ListSchedules", input)
	return &ListSchedulesFuture{Future: future}
}

func (a *stub) ListTagsForResource(ctx workflow.Context, input *gluedatabrew.ListTagsForResourceInput) (*gluedatabrew.ListTagsForResourceOutput, error) {
	var output gluedatabrew.ListTagsForResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws-gluedatabrew-ListTagsForResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListTagsForResourceAsync(ctx workflow.Context, input *gluedatabrew.ListTagsForResourceInput) *ListTagsForResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws-gluedatabrew-ListTagsForResource", input)
	return &ListTagsForResourceFuture{Future: future}
}

func (a *stub) PublishRecipe(ctx workflow.Context, input *gluedatabrew.PublishRecipeInput) (*gluedatabrew.PublishRecipeOutput, error) {
	var output gluedatabrew.PublishRecipeOutput
	err := workflow.ExecuteActivity(ctx, "aws-gluedatabrew-PublishRecipe", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) PublishRecipeAsync(ctx workflow.Context, input *gluedatabrew.PublishRecipeInput) *PublishRecipeFuture {
	future := workflow.ExecuteActivity(ctx, "aws-gluedatabrew-PublishRecipe", input)
	return &PublishRecipeFuture{Future: future}
}

func (a *stub) SendProjectSessionAction(ctx workflow.Context, input *gluedatabrew.SendProjectSessionActionInput) (*gluedatabrew.SendProjectSessionActionOutput, error) {
	var output gluedatabrew.SendProjectSessionActionOutput
	err := workflow.ExecuteActivity(ctx, "aws-gluedatabrew-SendProjectSessionAction", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) SendProjectSessionActionAsync(ctx workflow.Context, input *gluedatabrew.SendProjectSessionActionInput) *SendProjectSessionActionFuture {
	future := workflow.ExecuteActivity(ctx, "aws-gluedatabrew-SendProjectSessionAction", input)
	return &SendProjectSessionActionFuture{Future: future}
}

func (a *stub) StartJobRun(ctx workflow.Context, input *gluedatabrew.StartJobRunInput) (*gluedatabrew.StartJobRunOutput, error) {
	var output gluedatabrew.StartJobRunOutput
	err := workflow.ExecuteActivity(ctx, "aws-gluedatabrew-StartJobRun", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) StartJobRunAsync(ctx workflow.Context, input *gluedatabrew.StartJobRunInput) *StartJobRunFuture {
	future := workflow.ExecuteActivity(ctx, "aws-gluedatabrew-StartJobRun", input)
	return &StartJobRunFuture{Future: future}
}

func (a *stub) StartProjectSession(ctx workflow.Context, input *gluedatabrew.StartProjectSessionInput) (*gluedatabrew.StartProjectSessionOutput, error) {
	var output gluedatabrew.StartProjectSessionOutput
	err := workflow.ExecuteActivity(ctx, "aws-gluedatabrew-StartProjectSession", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) StartProjectSessionAsync(ctx workflow.Context, input *gluedatabrew.StartProjectSessionInput) *StartProjectSessionFuture {
	future := workflow.ExecuteActivity(ctx, "aws-gluedatabrew-StartProjectSession", input)
	return &StartProjectSessionFuture{Future: future}
}

func (a *stub) StopJobRun(ctx workflow.Context, input *gluedatabrew.StopJobRunInput) (*gluedatabrew.StopJobRunOutput, error) {
	var output gluedatabrew.StopJobRunOutput
	err := workflow.ExecuteActivity(ctx, "aws-gluedatabrew-StopJobRun", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) StopJobRunAsync(ctx workflow.Context, input *gluedatabrew.StopJobRunInput) *StopJobRunFuture {
	future := workflow.ExecuteActivity(ctx, "aws-gluedatabrew-StopJobRun", input)
	return &StopJobRunFuture{Future: future}
}

func (a *stub) TagResource(ctx workflow.Context, input *gluedatabrew.TagResourceInput) (*gluedatabrew.TagResourceOutput, error) {
	var output gluedatabrew.TagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws-gluedatabrew-TagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) TagResourceAsync(ctx workflow.Context, input *gluedatabrew.TagResourceInput) *TagResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws-gluedatabrew-TagResource", input)
	return &TagResourceFuture{Future: future}
}

func (a *stub) UntagResource(ctx workflow.Context, input *gluedatabrew.UntagResourceInput) (*gluedatabrew.UntagResourceOutput, error) {
	var output gluedatabrew.UntagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws-gluedatabrew-UntagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UntagResourceAsync(ctx workflow.Context, input *gluedatabrew.UntagResourceInput) *UntagResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws-gluedatabrew-UntagResource", input)
	return &UntagResourceFuture{Future: future}
}

func (a *stub) UpdateDataset(ctx workflow.Context, input *gluedatabrew.UpdateDatasetInput) (*gluedatabrew.UpdateDatasetOutput, error) {
	var output gluedatabrew.UpdateDatasetOutput
	err := workflow.ExecuteActivity(ctx, "aws-gluedatabrew-UpdateDataset", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateDatasetAsync(ctx workflow.Context, input *gluedatabrew.UpdateDatasetInput) *UpdateDatasetFuture {
	future := workflow.ExecuteActivity(ctx, "aws-gluedatabrew-UpdateDataset", input)
	return &UpdateDatasetFuture{Future: future}
}

func (a *stub) UpdateProfileJob(ctx workflow.Context, input *gluedatabrew.UpdateProfileJobInput) (*gluedatabrew.UpdateProfileJobOutput, error) {
	var output gluedatabrew.UpdateProfileJobOutput
	err := workflow.ExecuteActivity(ctx, "aws-gluedatabrew-UpdateProfileJob", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateProfileJobAsync(ctx workflow.Context, input *gluedatabrew.UpdateProfileJobInput) *UpdateProfileJobFuture {
	future := workflow.ExecuteActivity(ctx, "aws-gluedatabrew-UpdateProfileJob", input)
	return &UpdateProfileJobFuture{Future: future}
}

func (a *stub) UpdateProject(ctx workflow.Context, input *gluedatabrew.UpdateProjectInput) (*gluedatabrew.UpdateProjectOutput, error) {
	var output gluedatabrew.UpdateProjectOutput
	err := workflow.ExecuteActivity(ctx, "aws-gluedatabrew-UpdateProject", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateProjectAsync(ctx workflow.Context, input *gluedatabrew.UpdateProjectInput) *UpdateProjectFuture {
	future := workflow.ExecuteActivity(ctx, "aws-gluedatabrew-UpdateProject", input)
	return &UpdateProjectFuture{Future: future}
}

func (a *stub) UpdateRecipe(ctx workflow.Context, input *gluedatabrew.UpdateRecipeInput) (*gluedatabrew.UpdateRecipeOutput, error) {
	var output gluedatabrew.UpdateRecipeOutput
	err := workflow.ExecuteActivity(ctx, "aws-gluedatabrew-UpdateRecipe", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateRecipeAsync(ctx workflow.Context, input *gluedatabrew.UpdateRecipeInput) *UpdateRecipeFuture {
	future := workflow.ExecuteActivity(ctx, "aws-gluedatabrew-UpdateRecipe", input)
	return &UpdateRecipeFuture{Future: future}
}

func (a *stub) UpdateRecipeJob(ctx workflow.Context, input *gluedatabrew.UpdateRecipeJobInput) (*gluedatabrew.UpdateRecipeJobOutput, error) {
	var output gluedatabrew.UpdateRecipeJobOutput
	err := workflow.ExecuteActivity(ctx, "aws-gluedatabrew-UpdateRecipeJob", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateRecipeJobAsync(ctx workflow.Context, input *gluedatabrew.UpdateRecipeJobInput) *UpdateRecipeJobFuture {
	future := workflow.ExecuteActivity(ctx, "aws-gluedatabrew-UpdateRecipeJob", input)
	return &UpdateRecipeJobFuture{Future: future}
}

func (a *stub) UpdateSchedule(ctx workflow.Context, input *gluedatabrew.UpdateScheduleInput) (*gluedatabrew.UpdateScheduleOutput, error) {
	var output gluedatabrew.UpdateScheduleOutput
	err := workflow.ExecuteActivity(ctx, "aws-gluedatabrew-UpdateSchedule", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateScheduleAsync(ctx workflow.Context, input *gluedatabrew.UpdateScheduleInput) *UpdateScheduleFuture {
	future := workflow.ExecuteActivity(ctx, "aws-gluedatabrew-UpdateSchedule", input)
	return &UpdateScheduleFuture{Future: future}
}