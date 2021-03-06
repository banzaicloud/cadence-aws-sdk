// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7

package mobilestub

import (
	"github.com/aws/aws-sdk-go/service/mobile"
	"go.uber.org/cadence/workflow"

	"github.com/banzaicloud/cadence-aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type stub struct{}

type CreateProjectFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CreateProjectFuture) Get(ctx workflow.Context) (*mobile.CreateProjectOutput, error) {
	var output mobile.CreateProjectOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DeleteProjectFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DeleteProjectFuture) Get(ctx workflow.Context) (*mobile.DeleteProjectOutput, error) {
	var output mobile.DeleteProjectOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DescribeBundleFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DescribeBundleFuture) Get(ctx workflow.Context) (*mobile.DescribeBundleOutput, error) {
	var output mobile.DescribeBundleOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DescribeProjectFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DescribeProjectFuture) Get(ctx workflow.Context) (*mobile.DescribeProjectOutput, error) {
	var output mobile.DescribeProjectOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ExportBundleFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ExportBundleFuture) Get(ctx workflow.Context) (*mobile.ExportBundleOutput, error) {
	var output mobile.ExportBundleOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ExportProjectFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ExportProjectFuture) Get(ctx workflow.Context) (*mobile.ExportProjectOutput, error) {
	var output mobile.ExportProjectOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListBundlesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListBundlesFuture) Get(ctx workflow.Context) (*mobile.ListBundlesOutput, error) {
	var output mobile.ListBundlesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListProjectsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListProjectsFuture) Get(ctx workflow.Context) (*mobile.ListProjectsOutput, error) {
	var output mobile.ListProjectsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type UpdateProjectFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *UpdateProjectFuture) Get(ctx workflow.Context) (*mobile.UpdateProjectOutput, error) {
	var output mobile.UpdateProjectOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateProject(ctx workflow.Context, input *mobile.CreateProjectInput) (*mobile.CreateProjectOutput, error) {
	var output mobile.CreateProjectOutput
	err := workflow.ExecuteActivity(ctx, "aws-mobile-CreateProject", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateProjectAsync(ctx workflow.Context, input *mobile.CreateProjectInput) *CreateProjectFuture {
	future := workflow.ExecuteActivity(ctx, "aws-mobile-CreateProject", input)
	return &CreateProjectFuture{Future: future}
}

func (a *stub) DeleteProject(ctx workflow.Context, input *mobile.DeleteProjectInput) (*mobile.DeleteProjectOutput, error) {
	var output mobile.DeleteProjectOutput
	err := workflow.ExecuteActivity(ctx, "aws-mobile-DeleteProject", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteProjectAsync(ctx workflow.Context, input *mobile.DeleteProjectInput) *DeleteProjectFuture {
	future := workflow.ExecuteActivity(ctx, "aws-mobile-DeleteProject", input)
	return &DeleteProjectFuture{Future: future}
}

func (a *stub) DescribeBundle(ctx workflow.Context, input *mobile.DescribeBundleInput) (*mobile.DescribeBundleOutput, error) {
	var output mobile.DescribeBundleOutput
	err := workflow.ExecuteActivity(ctx, "aws-mobile-DescribeBundle", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeBundleAsync(ctx workflow.Context, input *mobile.DescribeBundleInput) *DescribeBundleFuture {
	future := workflow.ExecuteActivity(ctx, "aws-mobile-DescribeBundle", input)
	return &DescribeBundleFuture{Future: future}
}

func (a *stub) DescribeProject(ctx workflow.Context, input *mobile.DescribeProjectInput) (*mobile.DescribeProjectOutput, error) {
	var output mobile.DescribeProjectOutput
	err := workflow.ExecuteActivity(ctx, "aws-mobile-DescribeProject", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeProjectAsync(ctx workflow.Context, input *mobile.DescribeProjectInput) *DescribeProjectFuture {
	future := workflow.ExecuteActivity(ctx, "aws-mobile-DescribeProject", input)
	return &DescribeProjectFuture{Future: future}
}

func (a *stub) ExportBundle(ctx workflow.Context, input *mobile.ExportBundleInput) (*mobile.ExportBundleOutput, error) {
	var output mobile.ExportBundleOutput
	err := workflow.ExecuteActivity(ctx, "aws-mobile-ExportBundle", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ExportBundleAsync(ctx workflow.Context, input *mobile.ExportBundleInput) *ExportBundleFuture {
	future := workflow.ExecuteActivity(ctx, "aws-mobile-ExportBundle", input)
	return &ExportBundleFuture{Future: future}
}

func (a *stub) ExportProject(ctx workflow.Context, input *mobile.ExportProjectInput) (*mobile.ExportProjectOutput, error) {
	var output mobile.ExportProjectOutput
	err := workflow.ExecuteActivity(ctx, "aws-mobile-ExportProject", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ExportProjectAsync(ctx workflow.Context, input *mobile.ExportProjectInput) *ExportProjectFuture {
	future := workflow.ExecuteActivity(ctx, "aws-mobile-ExportProject", input)
	return &ExportProjectFuture{Future: future}
}

func (a *stub) ListBundles(ctx workflow.Context, input *mobile.ListBundlesInput) (*mobile.ListBundlesOutput, error) {
	var output mobile.ListBundlesOutput
	err := workflow.ExecuteActivity(ctx, "aws-mobile-ListBundles", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListBundlesAsync(ctx workflow.Context, input *mobile.ListBundlesInput) *ListBundlesFuture {
	future := workflow.ExecuteActivity(ctx, "aws-mobile-ListBundles", input)
	return &ListBundlesFuture{Future: future}
}

func (a *stub) ListProjects(ctx workflow.Context, input *mobile.ListProjectsInput) (*mobile.ListProjectsOutput, error) {
	var output mobile.ListProjectsOutput
	err := workflow.ExecuteActivity(ctx, "aws-mobile-ListProjects", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListProjectsAsync(ctx workflow.Context, input *mobile.ListProjectsInput) *ListProjectsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-mobile-ListProjects", input)
	return &ListProjectsFuture{Future: future}
}

func (a *stub) UpdateProject(ctx workflow.Context, input *mobile.UpdateProjectInput) (*mobile.UpdateProjectOutput, error) {
	var output mobile.UpdateProjectOutput
	err := workflow.ExecuteActivity(ctx, "aws-mobile-UpdateProject", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateProjectAsync(ctx workflow.Context, input *mobile.UpdateProjectInput) *UpdateProjectFuture {
	future := workflow.ExecuteActivity(ctx, "aws-mobile-UpdateProject", input)
	return &UpdateProjectFuture{Future: future}
}
