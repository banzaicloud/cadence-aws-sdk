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

type Client interface {
	CreateWorkspace(ctx workflow.Context, input *prometheusservice.CreateWorkspaceInput) (*prometheusservice.CreateWorkspaceOutput, error)
	CreateWorkspaceAsync(ctx workflow.Context, input *prometheusservice.CreateWorkspaceInput) *CreateWorkspaceFuture

	DeleteWorkspace(ctx workflow.Context, input *prometheusservice.DeleteWorkspaceInput) (*prometheusservice.DeleteWorkspaceOutput, error)
	DeleteWorkspaceAsync(ctx workflow.Context, input *prometheusservice.DeleteWorkspaceInput) *DeleteWorkspaceFuture

	DescribeWorkspace(ctx workflow.Context, input *prometheusservice.DescribeWorkspaceInput) (*prometheusservice.DescribeWorkspaceOutput, error)
	DescribeWorkspaceAsync(ctx workflow.Context, input *prometheusservice.DescribeWorkspaceInput) *DescribeWorkspaceFuture

	ListWorkspaces(ctx workflow.Context, input *prometheusservice.ListWorkspacesInput) (*prometheusservice.ListWorkspacesOutput, error)
	ListWorkspacesAsync(ctx workflow.Context, input *prometheusservice.ListWorkspacesInput) *ListWorkspacesFuture

	UpdateWorkspaceAlias(ctx workflow.Context, input *prometheusservice.UpdateWorkspaceAliasInput) (*prometheusservice.UpdateWorkspaceAliasOutput, error)
	UpdateWorkspaceAliasAsync(ctx workflow.Context, input *prometheusservice.UpdateWorkspaceAliasInput) *UpdateWorkspaceAliasFuture
}

func NewClient() Client {
	return &stub{}
}
