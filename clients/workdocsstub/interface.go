// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7

package workdocsstub

import (
	"github.com/aws/aws-sdk-go/service/workdocs"
	"go.uber.org/cadence/workflow"

	"github.com/banzaicloud/cadence-aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	AbortDocumentVersionUpload(ctx workflow.Context, input *workdocs.AbortDocumentVersionUploadInput) (*workdocs.AbortDocumentVersionUploadOutput, error)
	AbortDocumentVersionUploadAsync(ctx workflow.Context, input *workdocs.AbortDocumentVersionUploadInput) *AbortDocumentVersionUploadFuture

	ActivateUser(ctx workflow.Context, input *workdocs.ActivateUserInput) (*workdocs.ActivateUserOutput, error)
	ActivateUserAsync(ctx workflow.Context, input *workdocs.ActivateUserInput) *ActivateUserFuture

	AddResourcePermissions(ctx workflow.Context, input *workdocs.AddResourcePermissionsInput) (*workdocs.AddResourcePermissionsOutput, error)
	AddResourcePermissionsAsync(ctx workflow.Context, input *workdocs.AddResourcePermissionsInput) *AddResourcePermissionsFuture

	CreateComment(ctx workflow.Context, input *workdocs.CreateCommentInput) (*workdocs.CreateCommentOutput, error)
	CreateCommentAsync(ctx workflow.Context, input *workdocs.CreateCommentInput) *CreateCommentFuture

	CreateCustomMetadata(ctx workflow.Context, input *workdocs.CreateCustomMetadataInput) (*workdocs.CreateCustomMetadataOutput, error)
	CreateCustomMetadataAsync(ctx workflow.Context, input *workdocs.CreateCustomMetadataInput) *CreateCustomMetadataFuture

	CreateFolder(ctx workflow.Context, input *workdocs.CreateFolderInput) (*workdocs.CreateFolderOutput, error)
	CreateFolderAsync(ctx workflow.Context, input *workdocs.CreateFolderInput) *CreateFolderFuture

	CreateLabels(ctx workflow.Context, input *workdocs.CreateLabelsInput) (*workdocs.CreateLabelsOutput, error)
	CreateLabelsAsync(ctx workflow.Context, input *workdocs.CreateLabelsInput) *CreateLabelsFuture

	CreateNotificationSubscription(ctx workflow.Context, input *workdocs.CreateNotificationSubscriptionInput) (*workdocs.CreateNotificationSubscriptionOutput, error)
	CreateNotificationSubscriptionAsync(ctx workflow.Context, input *workdocs.CreateNotificationSubscriptionInput) *CreateNotificationSubscriptionFuture

	CreateUser(ctx workflow.Context, input *workdocs.CreateUserInput) (*workdocs.CreateUserOutput, error)
	CreateUserAsync(ctx workflow.Context, input *workdocs.CreateUserInput) *CreateUserFuture

	DeactivateUser(ctx workflow.Context, input *workdocs.DeactivateUserInput) (*workdocs.DeactivateUserOutput, error)
	DeactivateUserAsync(ctx workflow.Context, input *workdocs.DeactivateUserInput) *DeactivateUserFuture

	DeleteComment(ctx workflow.Context, input *workdocs.DeleteCommentInput) (*workdocs.DeleteCommentOutput, error)
	DeleteCommentAsync(ctx workflow.Context, input *workdocs.DeleteCommentInput) *DeleteCommentFuture

	DeleteCustomMetadata(ctx workflow.Context, input *workdocs.DeleteCustomMetadataInput) (*workdocs.DeleteCustomMetadataOutput, error)
	DeleteCustomMetadataAsync(ctx workflow.Context, input *workdocs.DeleteCustomMetadataInput) *DeleteCustomMetadataFuture

	DeleteDocument(ctx workflow.Context, input *workdocs.DeleteDocumentInput) (*workdocs.DeleteDocumentOutput, error)
	DeleteDocumentAsync(ctx workflow.Context, input *workdocs.DeleteDocumentInput) *DeleteDocumentFuture

	DeleteFolder(ctx workflow.Context, input *workdocs.DeleteFolderInput) (*workdocs.DeleteFolderOutput, error)
	DeleteFolderAsync(ctx workflow.Context, input *workdocs.DeleteFolderInput) *DeleteFolderFuture

	DeleteFolderContents(ctx workflow.Context, input *workdocs.DeleteFolderContentsInput) (*workdocs.DeleteFolderContentsOutput, error)
	DeleteFolderContentsAsync(ctx workflow.Context, input *workdocs.DeleteFolderContentsInput) *DeleteFolderContentsFuture

	DeleteLabels(ctx workflow.Context, input *workdocs.DeleteLabelsInput) (*workdocs.DeleteLabelsOutput, error)
	DeleteLabelsAsync(ctx workflow.Context, input *workdocs.DeleteLabelsInput) *DeleteLabelsFuture

	DeleteNotificationSubscription(ctx workflow.Context, input *workdocs.DeleteNotificationSubscriptionInput) (*workdocs.DeleteNotificationSubscriptionOutput, error)
	DeleteNotificationSubscriptionAsync(ctx workflow.Context, input *workdocs.DeleteNotificationSubscriptionInput) *DeleteNotificationSubscriptionFuture

	DeleteUser(ctx workflow.Context, input *workdocs.DeleteUserInput) (*workdocs.DeleteUserOutput, error)
	DeleteUserAsync(ctx workflow.Context, input *workdocs.DeleteUserInput) *DeleteUserFuture

	DescribeActivities(ctx workflow.Context, input *workdocs.DescribeActivitiesInput) (*workdocs.DescribeActivitiesOutput, error)
	DescribeActivitiesAsync(ctx workflow.Context, input *workdocs.DescribeActivitiesInput) *DescribeActivitiesFuture

	DescribeComments(ctx workflow.Context, input *workdocs.DescribeCommentsInput) (*workdocs.DescribeCommentsOutput, error)
	DescribeCommentsAsync(ctx workflow.Context, input *workdocs.DescribeCommentsInput) *DescribeCommentsFuture

	DescribeDocumentVersions(ctx workflow.Context, input *workdocs.DescribeDocumentVersionsInput) (*workdocs.DescribeDocumentVersionsOutput, error)
	DescribeDocumentVersionsAsync(ctx workflow.Context, input *workdocs.DescribeDocumentVersionsInput) *DescribeDocumentVersionsFuture

	DescribeFolderContents(ctx workflow.Context, input *workdocs.DescribeFolderContentsInput) (*workdocs.DescribeFolderContentsOutput, error)
	DescribeFolderContentsAsync(ctx workflow.Context, input *workdocs.DescribeFolderContentsInput) *DescribeFolderContentsFuture

	DescribeGroups(ctx workflow.Context, input *workdocs.DescribeGroupsInput) (*workdocs.DescribeGroupsOutput, error)
	DescribeGroupsAsync(ctx workflow.Context, input *workdocs.DescribeGroupsInput) *DescribeGroupsFuture

	DescribeNotificationSubscriptions(ctx workflow.Context, input *workdocs.DescribeNotificationSubscriptionsInput) (*workdocs.DescribeNotificationSubscriptionsOutput, error)
	DescribeNotificationSubscriptionsAsync(ctx workflow.Context, input *workdocs.DescribeNotificationSubscriptionsInput) *DescribeNotificationSubscriptionsFuture

	DescribeResourcePermissions(ctx workflow.Context, input *workdocs.DescribeResourcePermissionsInput) (*workdocs.DescribeResourcePermissionsOutput, error)
	DescribeResourcePermissionsAsync(ctx workflow.Context, input *workdocs.DescribeResourcePermissionsInput) *DescribeResourcePermissionsFuture

	DescribeRootFolders(ctx workflow.Context, input *workdocs.DescribeRootFoldersInput) (*workdocs.DescribeRootFoldersOutput, error)
	DescribeRootFoldersAsync(ctx workflow.Context, input *workdocs.DescribeRootFoldersInput) *DescribeRootFoldersFuture

	DescribeUsers(ctx workflow.Context, input *workdocs.DescribeUsersInput) (*workdocs.DescribeUsersOutput, error)
	DescribeUsersAsync(ctx workflow.Context, input *workdocs.DescribeUsersInput) *DescribeUsersFuture

	GetCurrentUser(ctx workflow.Context, input *workdocs.GetCurrentUserInput) (*workdocs.GetCurrentUserOutput, error)
	GetCurrentUserAsync(ctx workflow.Context, input *workdocs.GetCurrentUserInput) *GetCurrentUserFuture

	GetDocument(ctx workflow.Context, input *workdocs.GetDocumentInput) (*workdocs.GetDocumentOutput, error)
	GetDocumentAsync(ctx workflow.Context, input *workdocs.GetDocumentInput) *GetDocumentFuture

	GetDocumentPath(ctx workflow.Context, input *workdocs.GetDocumentPathInput) (*workdocs.GetDocumentPathOutput, error)
	GetDocumentPathAsync(ctx workflow.Context, input *workdocs.GetDocumentPathInput) *GetDocumentPathFuture

	GetDocumentVersion(ctx workflow.Context, input *workdocs.GetDocumentVersionInput) (*workdocs.GetDocumentVersionOutput, error)
	GetDocumentVersionAsync(ctx workflow.Context, input *workdocs.GetDocumentVersionInput) *GetDocumentVersionFuture

	GetFolder(ctx workflow.Context, input *workdocs.GetFolderInput) (*workdocs.GetFolderOutput, error)
	GetFolderAsync(ctx workflow.Context, input *workdocs.GetFolderInput) *GetFolderFuture

	GetFolderPath(ctx workflow.Context, input *workdocs.GetFolderPathInput) (*workdocs.GetFolderPathOutput, error)
	GetFolderPathAsync(ctx workflow.Context, input *workdocs.GetFolderPathInput) *GetFolderPathFuture

	GetResources(ctx workflow.Context, input *workdocs.GetResourcesInput) (*workdocs.GetResourcesOutput, error)
	GetResourcesAsync(ctx workflow.Context, input *workdocs.GetResourcesInput) *GetResourcesFuture

	InitiateDocumentVersionUpload(ctx workflow.Context, input *workdocs.InitiateDocumentVersionUploadInput) (*workdocs.InitiateDocumentVersionUploadOutput, error)
	InitiateDocumentVersionUploadAsync(ctx workflow.Context, input *workdocs.InitiateDocumentVersionUploadInput) *InitiateDocumentVersionUploadFuture

	RemoveAllResourcePermissions(ctx workflow.Context, input *workdocs.RemoveAllResourcePermissionsInput) (*workdocs.RemoveAllResourcePermissionsOutput, error)
	RemoveAllResourcePermissionsAsync(ctx workflow.Context, input *workdocs.RemoveAllResourcePermissionsInput) *RemoveAllResourcePermissionsFuture

	RemoveResourcePermission(ctx workflow.Context, input *workdocs.RemoveResourcePermissionInput) (*workdocs.RemoveResourcePermissionOutput, error)
	RemoveResourcePermissionAsync(ctx workflow.Context, input *workdocs.RemoveResourcePermissionInput) *RemoveResourcePermissionFuture

	UpdateDocument(ctx workflow.Context, input *workdocs.UpdateDocumentInput) (*workdocs.UpdateDocumentOutput, error)
	UpdateDocumentAsync(ctx workflow.Context, input *workdocs.UpdateDocumentInput) *UpdateDocumentFuture

	UpdateDocumentVersion(ctx workflow.Context, input *workdocs.UpdateDocumentVersionInput) (*workdocs.UpdateDocumentVersionOutput, error)
	UpdateDocumentVersionAsync(ctx workflow.Context, input *workdocs.UpdateDocumentVersionInput) *UpdateDocumentVersionFuture

	UpdateFolder(ctx workflow.Context, input *workdocs.UpdateFolderInput) (*workdocs.UpdateFolderOutput, error)
	UpdateFolderAsync(ctx workflow.Context, input *workdocs.UpdateFolderInput) *UpdateFolderFuture

	UpdateUser(ctx workflow.Context, input *workdocs.UpdateUserInput) (*workdocs.UpdateUserOutput, error)
	UpdateUserAsync(ctx workflow.Context, input *workdocs.UpdateUserInput) *UpdateUserFuture
}

func NewClient() Client {
	return &stub{}
}
