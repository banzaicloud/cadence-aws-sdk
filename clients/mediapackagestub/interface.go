// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.
// Copyright (c) 2020 Banzai Cloud All rights reserved.

package mediapackagestub

import (
	"github.com/aws/aws-sdk-go/service/mediapackage"
	"go.uber.org/cadence/workflow"

	"github.com/banzaicloud/cadence-aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	ConfigureLogs(ctx workflow.Context, input *mediapackage.ConfigureLogsInput) (*mediapackage.ConfigureLogsOutput, error)
	ConfigureLogsAsync(ctx workflow.Context, input *mediapackage.ConfigureLogsInput) *MediaPackageConfigureLogsFuture

	CreateChannel(ctx workflow.Context, input *mediapackage.CreateChannelInput) (*mediapackage.CreateChannelOutput, error)
	CreateChannelAsync(ctx workflow.Context, input *mediapackage.CreateChannelInput) *MediaPackageCreateChannelFuture

	CreateHarvestJob(ctx workflow.Context, input *mediapackage.CreateHarvestJobInput) (*mediapackage.CreateHarvestJobOutput, error)
	CreateHarvestJobAsync(ctx workflow.Context, input *mediapackage.CreateHarvestJobInput) *MediaPackageCreateHarvestJobFuture

	CreateOriginEndpoint(ctx workflow.Context, input *mediapackage.CreateOriginEndpointInput) (*mediapackage.CreateOriginEndpointOutput, error)
	CreateOriginEndpointAsync(ctx workflow.Context, input *mediapackage.CreateOriginEndpointInput) *MediaPackageCreateOriginEndpointFuture

	DeleteChannel(ctx workflow.Context, input *mediapackage.DeleteChannelInput) (*mediapackage.DeleteChannelOutput, error)
	DeleteChannelAsync(ctx workflow.Context, input *mediapackage.DeleteChannelInput) *MediaPackageDeleteChannelFuture

	DeleteOriginEndpoint(ctx workflow.Context, input *mediapackage.DeleteOriginEndpointInput) (*mediapackage.DeleteOriginEndpointOutput, error)
	DeleteOriginEndpointAsync(ctx workflow.Context, input *mediapackage.DeleteOriginEndpointInput) *MediaPackageDeleteOriginEndpointFuture

	DescribeChannel(ctx workflow.Context, input *mediapackage.DescribeChannelInput) (*mediapackage.DescribeChannelOutput, error)
	DescribeChannelAsync(ctx workflow.Context, input *mediapackage.DescribeChannelInput) *MediaPackageDescribeChannelFuture

	DescribeHarvestJob(ctx workflow.Context, input *mediapackage.DescribeHarvestJobInput) (*mediapackage.DescribeHarvestJobOutput, error)
	DescribeHarvestJobAsync(ctx workflow.Context, input *mediapackage.DescribeHarvestJobInput) *MediaPackageDescribeHarvestJobFuture

	DescribeOriginEndpoint(ctx workflow.Context, input *mediapackage.DescribeOriginEndpointInput) (*mediapackage.DescribeOriginEndpointOutput, error)
	DescribeOriginEndpointAsync(ctx workflow.Context, input *mediapackage.DescribeOriginEndpointInput) *MediaPackageDescribeOriginEndpointFuture

	ListChannels(ctx workflow.Context, input *mediapackage.ListChannelsInput) (*mediapackage.ListChannelsOutput, error)
	ListChannelsAsync(ctx workflow.Context, input *mediapackage.ListChannelsInput) *MediaPackageListChannelsFuture

	ListHarvestJobs(ctx workflow.Context, input *mediapackage.ListHarvestJobsInput) (*mediapackage.ListHarvestJobsOutput, error)
	ListHarvestJobsAsync(ctx workflow.Context, input *mediapackage.ListHarvestJobsInput) *MediaPackageListHarvestJobsFuture

	ListOriginEndpoints(ctx workflow.Context, input *mediapackage.ListOriginEndpointsInput) (*mediapackage.ListOriginEndpointsOutput, error)
	ListOriginEndpointsAsync(ctx workflow.Context, input *mediapackage.ListOriginEndpointsInput) *MediaPackageListOriginEndpointsFuture

	ListTagsForResource(ctx workflow.Context, input *mediapackage.ListTagsForResourceInput) (*mediapackage.ListTagsForResourceOutput, error)
	ListTagsForResourceAsync(ctx workflow.Context, input *mediapackage.ListTagsForResourceInput) *MediaPackageListTagsForResourceFuture

	RotateChannelCredentials(ctx workflow.Context, input *mediapackage.RotateChannelCredentialsInput) (*mediapackage.RotateChannelCredentialsOutput, error)
	RotateChannelCredentialsAsync(ctx workflow.Context, input *mediapackage.RotateChannelCredentialsInput) *MediaPackageRotateChannelCredentialsFuture

	RotateIngestEndpointCredentials(ctx workflow.Context, input *mediapackage.RotateIngestEndpointCredentialsInput) (*mediapackage.RotateIngestEndpointCredentialsOutput, error)
	RotateIngestEndpointCredentialsAsync(ctx workflow.Context, input *mediapackage.RotateIngestEndpointCredentialsInput) *MediaPackageRotateIngestEndpointCredentialsFuture

	TagResource(ctx workflow.Context, input *mediapackage.TagResourceInput) (*mediapackage.TagResourceOutput, error)
	TagResourceAsync(ctx workflow.Context, input *mediapackage.TagResourceInput) *MediaPackageTagResourceFuture

	UntagResource(ctx workflow.Context, input *mediapackage.UntagResourceInput) (*mediapackage.UntagResourceOutput, error)
	UntagResourceAsync(ctx workflow.Context, input *mediapackage.UntagResourceInput) *MediaPackageUntagResourceFuture

	UpdateChannel(ctx workflow.Context, input *mediapackage.UpdateChannelInput) (*mediapackage.UpdateChannelOutput, error)
	UpdateChannelAsync(ctx workflow.Context, input *mediapackage.UpdateChannelInput) *MediaPackageUpdateChannelFuture

	UpdateOriginEndpoint(ctx workflow.Context, input *mediapackage.UpdateOriginEndpointInput) (*mediapackage.UpdateOriginEndpointOutput, error)
	UpdateOriginEndpointAsync(ctx workflow.Context, input *mediapackage.UpdateOriginEndpointInput) *MediaPackageUpdateOriginEndpointFuture
}

func NewClient() Client {
	return &stub{}
}