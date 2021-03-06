// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7

package prometheusservicestub

import (
	"github.com/aws/aws-sdk-go/service/prometheusservice"
	"go.uber.org/cadence/workflow"

	"github.com/banzaicloud/cadence-aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type stub struct{}

type CreateWorkspaceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CreateWorkspaceFuture) Get(ctx workflow.Context) (*prometheusservice.CreateWorkspaceOutput, error) {
	var output prometheusservice.CreateWorkspaceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DeleteWorkspaceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DeleteWorkspaceFuture) Get(ctx workflow.Context) (*prometheusservice.DeleteWorkspaceOutput, error) {
	var output prometheusservice.DeleteWorkspaceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DescribeWorkspaceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DescribeWorkspaceFuture) Get(ctx workflow.Context) (*prometheusservice.DescribeWorkspaceOutput, error) {
	var output prometheusservice.DescribeWorkspaceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListWorkspacesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListWorkspacesFuture) Get(ctx workflow.Context) (*prometheusservice.ListWorkspacesOutput, error) {
	var output prometheusservice.ListWorkspacesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type UpdateWorkspaceAliasFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *UpdateWorkspaceAliasFuture) Get(ctx workflow.Context) (*prometheusservice.UpdateWorkspaceAliasOutput, error) {
	var output prometheusservice.UpdateWorkspaceAliasOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateWorkspace(ctx workflow.Context, input *prometheusservice.CreateWorkspaceInput) (*prometheusservice.CreateWorkspaceOutput, error) {
	var output prometheusservice.CreateWorkspaceOutput
	err := workflow.ExecuteActivity(ctx, "aws-prometheusservice-CreateWorkspace", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateWorkspaceAsync(ctx workflow.Context, input *prometheusservice.CreateWorkspaceInput) *CreateWorkspaceFuture {
	future := workflow.ExecuteActivity(ctx, "aws-prometheusservice-CreateWorkspace", input)
	return &CreateWorkspaceFuture{Future: future}
}

func (a *stub) DeleteWorkspace(ctx workflow.Context, input *prometheusservice.DeleteWorkspaceInput) (*prometheusservice.DeleteWorkspaceOutput, error) {
	var output prometheusservice.DeleteWorkspaceOutput
	err := workflow.ExecuteActivity(ctx, "aws-prometheusservice-DeleteWorkspace", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteWorkspaceAsync(ctx workflow.Context, input *prometheusservice.DeleteWorkspaceInput) *DeleteWorkspaceFuture {
	future := workflow.ExecuteActivity(ctx, "aws-prometheusservice-DeleteWorkspace", input)
	return &DeleteWorkspaceFuture{Future: future}
}

func (a *stub) DescribeWorkspace(ctx workflow.Context, input *prometheusservice.DescribeWorkspaceInput) (*prometheusservice.DescribeWorkspaceOutput, error) {
	var output prometheusservice.DescribeWorkspaceOutput
	err := workflow.ExecuteActivity(ctx, "aws-prometheusservice-DescribeWorkspace", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeWorkspaceAsync(ctx workflow.Context, input *prometheusservice.DescribeWorkspaceInput) *DescribeWorkspaceFuture {
	future := workflow.ExecuteActivity(ctx, "aws-prometheusservice-DescribeWorkspace", input)
	return &DescribeWorkspaceFuture{Future: future}
}

func (a *stub) ListWorkspaces(ctx workflow.Context, input *prometheusservice.ListWorkspacesInput) (*prometheusservice.ListWorkspacesOutput, error) {
	var output prometheusservice.ListWorkspacesOutput
	err := workflow.ExecuteActivity(ctx, "aws-prometheusservice-ListWorkspaces", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListWorkspacesAsync(ctx workflow.Context, input *prometheusservice.ListWorkspacesInput) *ListWorkspacesFuture {
	future := workflow.ExecuteActivity(ctx, "aws-prometheusservice-ListWorkspaces", input)
	return &ListWorkspacesFuture{Future: future}
}

func (a *stub) UpdateWorkspaceAlias(ctx workflow.Context, input *prometheusservice.UpdateWorkspaceAliasInput) (*prometheusservice.UpdateWorkspaceAliasOutput, error) {
	var output prometheusservice.UpdateWorkspaceAliasOutput
	err := workflow.ExecuteActivity(ctx, "aws-prometheusservice-UpdateWorkspaceAlias", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateWorkspaceAliasAsync(ctx workflow.Context, input *prometheusservice.UpdateWorkspaceAliasInput) *UpdateWorkspaceAliasFuture {
	future := workflow.ExecuteActivity(ctx, "aws-prometheusservice-UpdateWorkspaceAlias", input)
	return &UpdateWorkspaceAliasFuture{Future: future}
}
