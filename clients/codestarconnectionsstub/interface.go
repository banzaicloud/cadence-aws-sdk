// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.
// Copyright (c) 2020 Banzai Cloud All rights reserved.

package codestarconnectionsstub

import (
	"github.com/aws/aws-sdk-go/service/codestarconnections"
	"go.uber.org/cadence/workflow"

	"github.com/banzaicloud/cadence-aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	CreateConnection(ctx workflow.Context, input *codestarconnections.CreateConnectionInput) (*codestarconnections.CreateConnectionOutput, error)
	CreateConnectionAsync(ctx workflow.Context, input *codestarconnections.CreateConnectionInput) *CodeStarConnectionsCreateConnectionFuture

	CreateHost(ctx workflow.Context, input *codestarconnections.CreateHostInput) (*codestarconnections.CreateHostOutput, error)
	CreateHostAsync(ctx workflow.Context, input *codestarconnections.CreateHostInput) *CodeStarConnectionsCreateHostFuture

	DeleteConnection(ctx workflow.Context, input *codestarconnections.DeleteConnectionInput) (*codestarconnections.DeleteConnectionOutput, error)
	DeleteConnectionAsync(ctx workflow.Context, input *codestarconnections.DeleteConnectionInput) *CodeStarConnectionsDeleteConnectionFuture

	DeleteHost(ctx workflow.Context, input *codestarconnections.DeleteHostInput) (*codestarconnections.DeleteHostOutput, error)
	DeleteHostAsync(ctx workflow.Context, input *codestarconnections.DeleteHostInput) *CodeStarConnectionsDeleteHostFuture

	GetConnection(ctx workflow.Context, input *codestarconnections.GetConnectionInput) (*codestarconnections.GetConnectionOutput, error)
	GetConnectionAsync(ctx workflow.Context, input *codestarconnections.GetConnectionInput) *CodeStarConnectionsGetConnectionFuture

	GetHost(ctx workflow.Context, input *codestarconnections.GetHostInput) (*codestarconnections.GetHostOutput, error)
	GetHostAsync(ctx workflow.Context, input *codestarconnections.GetHostInput) *CodeStarConnectionsGetHostFuture

	ListConnections(ctx workflow.Context, input *codestarconnections.ListConnectionsInput) (*codestarconnections.ListConnectionsOutput, error)
	ListConnectionsAsync(ctx workflow.Context, input *codestarconnections.ListConnectionsInput) *CodeStarConnectionsListConnectionsFuture

	ListHosts(ctx workflow.Context, input *codestarconnections.ListHostsInput) (*codestarconnections.ListHostsOutput, error)
	ListHostsAsync(ctx workflow.Context, input *codestarconnections.ListHostsInput) *CodeStarConnectionsListHostsFuture

	ListTagsForResource(ctx workflow.Context, input *codestarconnections.ListTagsForResourceInput) (*codestarconnections.ListTagsForResourceOutput, error)
	ListTagsForResourceAsync(ctx workflow.Context, input *codestarconnections.ListTagsForResourceInput) *CodeStarConnectionsListTagsForResourceFuture

	TagResource(ctx workflow.Context, input *codestarconnections.TagResourceInput) (*codestarconnections.TagResourceOutput, error)
	TagResourceAsync(ctx workflow.Context, input *codestarconnections.TagResourceInput) *CodeStarConnectionsTagResourceFuture

	UntagResource(ctx workflow.Context, input *codestarconnections.UntagResourceInput) (*codestarconnections.UntagResourceOutput, error)
	UntagResourceAsync(ctx workflow.Context, input *codestarconnections.UntagResourceInput) *CodeStarConnectionsUntagResourceFuture
}

func NewClient() Client {
	return &stub{}
}
