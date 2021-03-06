// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7

package identitystorestub

import (
	"github.com/aws/aws-sdk-go/service/identitystore"
	"go.uber.org/cadence/workflow"

	"github.com/banzaicloud/cadence-aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	DescribeGroup(ctx workflow.Context, input *identitystore.DescribeGroupInput) (*identitystore.DescribeGroupOutput, error)
	DescribeGroupAsync(ctx workflow.Context, input *identitystore.DescribeGroupInput) *DescribeGroupFuture

	DescribeUser(ctx workflow.Context, input *identitystore.DescribeUserInput) (*identitystore.DescribeUserOutput, error)
	DescribeUserAsync(ctx workflow.Context, input *identitystore.DescribeUserInput) *DescribeUserFuture

	ListGroups(ctx workflow.Context, input *identitystore.ListGroupsInput) (*identitystore.ListGroupsOutput, error)
	ListGroupsAsync(ctx workflow.Context, input *identitystore.ListGroupsInput) *ListGroupsFuture

	ListUsers(ctx workflow.Context, input *identitystore.ListUsersInput) (*identitystore.ListUsersOutput, error)
	ListUsersAsync(ctx workflow.Context, input *identitystore.ListUsersInput) *ListUsersFuture
}

func NewClient() Client {
	return &stub{}
}
