// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7

package appstreamstub

import (
	"github.com/aws/aws-sdk-go/service/appstream"
	"go.uber.org/cadence/workflow"

	"github.com/banzaicloud/cadence-aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	AssociateFleet(ctx workflow.Context, input *appstream.AssociateFleetInput) (*appstream.AssociateFleetOutput, error)
	AssociateFleetAsync(ctx workflow.Context, input *appstream.AssociateFleetInput) *AppStreamAssociateFleetFuture

	BatchAssociateUserStack(ctx workflow.Context, input *appstream.BatchAssociateUserStackInput) (*appstream.BatchAssociateUserStackOutput, error)
	BatchAssociateUserStackAsync(ctx workflow.Context, input *appstream.BatchAssociateUserStackInput) *AppStreamBatchAssociateUserStackFuture

	BatchDisassociateUserStack(ctx workflow.Context, input *appstream.BatchDisassociateUserStackInput) (*appstream.BatchDisassociateUserStackOutput, error)
	BatchDisassociateUserStackAsync(ctx workflow.Context, input *appstream.BatchDisassociateUserStackInput) *AppStreamBatchDisassociateUserStackFuture

	CopyImage(ctx workflow.Context, input *appstream.CopyImageInput) (*appstream.CopyImageOutput, error)
	CopyImageAsync(ctx workflow.Context, input *appstream.CopyImageInput) *AppStreamCopyImageFuture

	CreateDirectoryConfig(ctx workflow.Context, input *appstream.CreateDirectoryConfigInput) (*appstream.CreateDirectoryConfigOutput, error)
	CreateDirectoryConfigAsync(ctx workflow.Context, input *appstream.CreateDirectoryConfigInput) *AppStreamCreateDirectoryConfigFuture

	CreateFleet(ctx workflow.Context, input *appstream.CreateFleetInput) (*appstream.CreateFleetOutput, error)
	CreateFleetAsync(ctx workflow.Context, input *appstream.CreateFleetInput) *AppStreamCreateFleetFuture

	CreateImageBuilder(ctx workflow.Context, input *appstream.CreateImageBuilderInput) (*appstream.CreateImageBuilderOutput, error)
	CreateImageBuilderAsync(ctx workflow.Context, input *appstream.CreateImageBuilderInput) *AppStreamCreateImageBuilderFuture

	CreateImageBuilderStreamingURL(ctx workflow.Context, input *appstream.CreateImageBuilderStreamingURLInput) (*appstream.CreateImageBuilderStreamingURLOutput, error)
	CreateImageBuilderStreamingURLAsync(ctx workflow.Context, input *appstream.CreateImageBuilderStreamingURLInput) *AppStreamCreateImageBuilderStreamingURLFuture

	CreateStack(ctx workflow.Context, input *appstream.CreateStackInput) (*appstream.CreateStackOutput, error)
	CreateStackAsync(ctx workflow.Context, input *appstream.CreateStackInput) *AppStreamCreateStackFuture

	CreateStreamingURL(ctx workflow.Context, input *appstream.CreateStreamingURLInput) (*appstream.CreateStreamingURLOutput, error)
	CreateStreamingURLAsync(ctx workflow.Context, input *appstream.CreateStreamingURLInput) *AppStreamCreateStreamingURLFuture

	CreateUsageReportSubscription(ctx workflow.Context, input *appstream.CreateUsageReportSubscriptionInput) (*appstream.CreateUsageReportSubscriptionOutput, error)
	CreateUsageReportSubscriptionAsync(ctx workflow.Context, input *appstream.CreateUsageReportSubscriptionInput) *AppStreamCreateUsageReportSubscriptionFuture

	CreateUser(ctx workflow.Context, input *appstream.CreateUserInput) (*appstream.CreateUserOutput, error)
	CreateUserAsync(ctx workflow.Context, input *appstream.CreateUserInput) *AppStreamCreateUserFuture

	DeleteDirectoryConfig(ctx workflow.Context, input *appstream.DeleteDirectoryConfigInput) (*appstream.DeleteDirectoryConfigOutput, error)
	DeleteDirectoryConfigAsync(ctx workflow.Context, input *appstream.DeleteDirectoryConfigInput) *AppStreamDeleteDirectoryConfigFuture

	DeleteFleet(ctx workflow.Context, input *appstream.DeleteFleetInput) (*appstream.DeleteFleetOutput, error)
	DeleteFleetAsync(ctx workflow.Context, input *appstream.DeleteFleetInput) *AppStreamDeleteFleetFuture

	DeleteImage(ctx workflow.Context, input *appstream.DeleteImageInput) (*appstream.DeleteImageOutput, error)
	DeleteImageAsync(ctx workflow.Context, input *appstream.DeleteImageInput) *AppStreamDeleteImageFuture

	DeleteImageBuilder(ctx workflow.Context, input *appstream.DeleteImageBuilderInput) (*appstream.DeleteImageBuilderOutput, error)
	DeleteImageBuilderAsync(ctx workflow.Context, input *appstream.DeleteImageBuilderInput) *AppStreamDeleteImageBuilderFuture

	DeleteImagePermissions(ctx workflow.Context, input *appstream.DeleteImagePermissionsInput) (*appstream.DeleteImagePermissionsOutput, error)
	DeleteImagePermissionsAsync(ctx workflow.Context, input *appstream.DeleteImagePermissionsInput) *AppStreamDeleteImagePermissionsFuture

	DeleteStack(ctx workflow.Context, input *appstream.DeleteStackInput) (*appstream.DeleteStackOutput, error)
	DeleteStackAsync(ctx workflow.Context, input *appstream.DeleteStackInput) *AppStreamDeleteStackFuture

	DeleteUsageReportSubscription(ctx workflow.Context, input *appstream.DeleteUsageReportSubscriptionInput) (*appstream.DeleteUsageReportSubscriptionOutput, error)
	DeleteUsageReportSubscriptionAsync(ctx workflow.Context, input *appstream.DeleteUsageReportSubscriptionInput) *AppStreamDeleteUsageReportSubscriptionFuture

	DeleteUser(ctx workflow.Context, input *appstream.DeleteUserInput) (*appstream.DeleteUserOutput, error)
	DeleteUserAsync(ctx workflow.Context, input *appstream.DeleteUserInput) *AppStreamDeleteUserFuture

	DescribeDirectoryConfigs(ctx workflow.Context, input *appstream.DescribeDirectoryConfigsInput) (*appstream.DescribeDirectoryConfigsOutput, error)
	DescribeDirectoryConfigsAsync(ctx workflow.Context, input *appstream.DescribeDirectoryConfigsInput) *AppStreamDescribeDirectoryConfigsFuture

	DescribeFleets(ctx workflow.Context, input *appstream.DescribeFleetsInput) (*appstream.DescribeFleetsOutput, error)
	DescribeFleetsAsync(ctx workflow.Context, input *appstream.DescribeFleetsInput) *AppStreamDescribeFleetsFuture

	DescribeImageBuilders(ctx workflow.Context, input *appstream.DescribeImageBuildersInput) (*appstream.DescribeImageBuildersOutput, error)
	DescribeImageBuildersAsync(ctx workflow.Context, input *appstream.DescribeImageBuildersInput) *AppStreamDescribeImageBuildersFuture

	DescribeImagePermissions(ctx workflow.Context, input *appstream.DescribeImagePermissionsInput) (*appstream.DescribeImagePermissionsOutput, error)
	DescribeImagePermissionsAsync(ctx workflow.Context, input *appstream.DescribeImagePermissionsInput) *AppStreamDescribeImagePermissionsFuture

	DescribeImages(ctx workflow.Context, input *appstream.DescribeImagesInput) (*appstream.DescribeImagesOutput, error)
	DescribeImagesAsync(ctx workflow.Context, input *appstream.DescribeImagesInput) *AppStreamDescribeImagesFuture

	DescribeSessions(ctx workflow.Context, input *appstream.DescribeSessionsInput) (*appstream.DescribeSessionsOutput, error)
	DescribeSessionsAsync(ctx workflow.Context, input *appstream.DescribeSessionsInput) *AppStreamDescribeSessionsFuture

	DescribeStacks(ctx workflow.Context, input *appstream.DescribeStacksInput) (*appstream.DescribeStacksOutput, error)
	DescribeStacksAsync(ctx workflow.Context, input *appstream.DescribeStacksInput) *AppStreamDescribeStacksFuture

	DescribeUsageReportSubscriptions(ctx workflow.Context, input *appstream.DescribeUsageReportSubscriptionsInput) (*appstream.DescribeUsageReportSubscriptionsOutput, error)
	DescribeUsageReportSubscriptionsAsync(ctx workflow.Context, input *appstream.DescribeUsageReportSubscriptionsInput) *AppStreamDescribeUsageReportSubscriptionsFuture

	DescribeUserStackAssociations(ctx workflow.Context, input *appstream.DescribeUserStackAssociationsInput) (*appstream.DescribeUserStackAssociationsOutput, error)
	DescribeUserStackAssociationsAsync(ctx workflow.Context, input *appstream.DescribeUserStackAssociationsInput) *AppStreamDescribeUserStackAssociationsFuture

	DescribeUsers(ctx workflow.Context, input *appstream.DescribeUsersInput) (*appstream.DescribeUsersOutput, error)
	DescribeUsersAsync(ctx workflow.Context, input *appstream.DescribeUsersInput) *AppStreamDescribeUsersFuture

	DisableUser(ctx workflow.Context, input *appstream.DisableUserInput) (*appstream.DisableUserOutput, error)
	DisableUserAsync(ctx workflow.Context, input *appstream.DisableUserInput) *AppStreamDisableUserFuture

	DisassociateFleet(ctx workflow.Context, input *appstream.DisassociateFleetInput) (*appstream.DisassociateFleetOutput, error)
	DisassociateFleetAsync(ctx workflow.Context, input *appstream.DisassociateFleetInput) *AppStreamDisassociateFleetFuture

	EnableUser(ctx workflow.Context, input *appstream.EnableUserInput) (*appstream.EnableUserOutput, error)
	EnableUserAsync(ctx workflow.Context, input *appstream.EnableUserInput) *AppStreamEnableUserFuture

	ExpireSession(ctx workflow.Context, input *appstream.ExpireSessionInput) (*appstream.ExpireSessionOutput, error)
	ExpireSessionAsync(ctx workflow.Context, input *appstream.ExpireSessionInput) *AppStreamExpireSessionFuture

	ListAssociatedFleets(ctx workflow.Context, input *appstream.ListAssociatedFleetsInput) (*appstream.ListAssociatedFleetsOutput, error)
	ListAssociatedFleetsAsync(ctx workflow.Context, input *appstream.ListAssociatedFleetsInput) *AppStreamListAssociatedFleetsFuture

	ListAssociatedStacks(ctx workflow.Context, input *appstream.ListAssociatedStacksInput) (*appstream.ListAssociatedStacksOutput, error)
	ListAssociatedStacksAsync(ctx workflow.Context, input *appstream.ListAssociatedStacksInput) *AppStreamListAssociatedStacksFuture

	ListTagsForResource(ctx workflow.Context, input *appstream.ListTagsForResourceInput) (*appstream.ListTagsForResourceOutput, error)
	ListTagsForResourceAsync(ctx workflow.Context, input *appstream.ListTagsForResourceInput) *AppStreamListTagsForResourceFuture

	StartFleet(ctx workflow.Context, input *appstream.StartFleetInput) (*appstream.StartFleetOutput, error)
	StartFleetAsync(ctx workflow.Context, input *appstream.StartFleetInput) *AppStreamStartFleetFuture

	StartImageBuilder(ctx workflow.Context, input *appstream.StartImageBuilderInput) (*appstream.StartImageBuilderOutput, error)
	StartImageBuilderAsync(ctx workflow.Context, input *appstream.StartImageBuilderInput) *AppStreamStartImageBuilderFuture

	StopFleet(ctx workflow.Context, input *appstream.StopFleetInput) (*appstream.StopFleetOutput, error)
	StopFleetAsync(ctx workflow.Context, input *appstream.StopFleetInput) *AppStreamStopFleetFuture

	StopImageBuilder(ctx workflow.Context, input *appstream.StopImageBuilderInput) (*appstream.StopImageBuilderOutput, error)
	StopImageBuilderAsync(ctx workflow.Context, input *appstream.StopImageBuilderInput) *AppStreamStopImageBuilderFuture

	TagResource(ctx workflow.Context, input *appstream.TagResourceInput) (*appstream.TagResourceOutput, error)
	TagResourceAsync(ctx workflow.Context, input *appstream.TagResourceInput) *AppStreamTagResourceFuture

	UntagResource(ctx workflow.Context, input *appstream.UntagResourceInput) (*appstream.UntagResourceOutput, error)
	UntagResourceAsync(ctx workflow.Context, input *appstream.UntagResourceInput) *AppStreamUntagResourceFuture

	UpdateDirectoryConfig(ctx workflow.Context, input *appstream.UpdateDirectoryConfigInput) (*appstream.UpdateDirectoryConfigOutput, error)
	UpdateDirectoryConfigAsync(ctx workflow.Context, input *appstream.UpdateDirectoryConfigInput) *AppStreamUpdateDirectoryConfigFuture

	UpdateFleet(ctx workflow.Context, input *appstream.UpdateFleetInput) (*appstream.UpdateFleetOutput, error)
	UpdateFleetAsync(ctx workflow.Context, input *appstream.UpdateFleetInput) *AppStreamUpdateFleetFuture

	UpdateImagePermissions(ctx workflow.Context, input *appstream.UpdateImagePermissionsInput) (*appstream.UpdateImagePermissionsOutput, error)
	UpdateImagePermissionsAsync(ctx workflow.Context, input *appstream.UpdateImagePermissionsInput) *AppStreamUpdateImagePermissionsFuture

	UpdateStack(ctx workflow.Context, input *appstream.UpdateStackInput) (*appstream.UpdateStackOutput, error)
	UpdateStackAsync(ctx workflow.Context, input *appstream.UpdateStackInput) *AppStreamUpdateStackFuture

	WaitUntilFleetStarted(ctx workflow.Context, input *appstream.DescribeFleetsInput) error
	WaitUntilFleetStartedAsync(ctx workflow.Context, input *appstream.DescribeFleetsInput) *clients.VoidFuture

	WaitUntilFleetStopped(ctx workflow.Context, input *appstream.DescribeFleetsInput) error
	WaitUntilFleetStoppedAsync(ctx workflow.Context, input *appstream.DescribeFleetsInput) *clients.VoidFuture
}

func NewClient() Client {
	return &stub{}
}
