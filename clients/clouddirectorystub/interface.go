// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7

package clouddirectorystub

import (
	"github.com/aws/aws-sdk-go/service/clouddirectory"
	"go.uber.org/cadence/workflow"

	"github.com/banzaicloud/cadence-aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	AddFacetToObject(ctx workflow.Context, input *clouddirectory.AddFacetToObjectInput) (*clouddirectory.AddFacetToObjectOutput, error)
	AddFacetToObjectAsync(ctx workflow.Context, input *clouddirectory.AddFacetToObjectInput) *CloudDirectoryAddFacetToObjectFuture

	ApplySchema(ctx workflow.Context, input *clouddirectory.ApplySchemaInput) (*clouddirectory.ApplySchemaOutput, error)
	ApplySchemaAsync(ctx workflow.Context, input *clouddirectory.ApplySchemaInput) *CloudDirectoryApplySchemaFuture

	AttachObject(ctx workflow.Context, input *clouddirectory.AttachObjectInput) (*clouddirectory.AttachObjectOutput, error)
	AttachObjectAsync(ctx workflow.Context, input *clouddirectory.AttachObjectInput) *CloudDirectoryAttachObjectFuture

	AttachPolicy(ctx workflow.Context, input *clouddirectory.AttachPolicyInput) (*clouddirectory.AttachPolicyOutput, error)
	AttachPolicyAsync(ctx workflow.Context, input *clouddirectory.AttachPolicyInput) *CloudDirectoryAttachPolicyFuture

	AttachToIndex(ctx workflow.Context, input *clouddirectory.AttachToIndexInput) (*clouddirectory.AttachToIndexOutput, error)
	AttachToIndexAsync(ctx workflow.Context, input *clouddirectory.AttachToIndexInput) *CloudDirectoryAttachToIndexFuture

	AttachTypedLink(ctx workflow.Context, input *clouddirectory.AttachTypedLinkInput) (*clouddirectory.AttachTypedLinkOutput, error)
	AttachTypedLinkAsync(ctx workflow.Context, input *clouddirectory.AttachTypedLinkInput) *CloudDirectoryAttachTypedLinkFuture

	BatchRead(ctx workflow.Context, input *clouddirectory.BatchReadInput) (*clouddirectory.BatchReadOutput, error)
	BatchReadAsync(ctx workflow.Context, input *clouddirectory.BatchReadInput) *CloudDirectoryBatchReadFuture

	BatchWrite(ctx workflow.Context, input *clouddirectory.BatchWriteInput) (*clouddirectory.BatchWriteOutput, error)
	BatchWriteAsync(ctx workflow.Context, input *clouddirectory.BatchWriteInput) *CloudDirectoryBatchWriteFuture

	CreateDirectory(ctx workflow.Context, input *clouddirectory.CreateDirectoryInput) (*clouddirectory.CreateDirectoryOutput, error)
	CreateDirectoryAsync(ctx workflow.Context, input *clouddirectory.CreateDirectoryInput) *CloudDirectoryCreateDirectoryFuture

	CreateFacet(ctx workflow.Context, input *clouddirectory.CreateFacetInput) (*clouddirectory.CreateFacetOutput, error)
	CreateFacetAsync(ctx workflow.Context, input *clouddirectory.CreateFacetInput) *CloudDirectoryCreateFacetFuture

	CreateIndex(ctx workflow.Context, input *clouddirectory.CreateIndexInput) (*clouddirectory.CreateIndexOutput, error)
	CreateIndexAsync(ctx workflow.Context, input *clouddirectory.CreateIndexInput) *CloudDirectoryCreateIndexFuture

	CreateObject(ctx workflow.Context, input *clouddirectory.CreateObjectInput) (*clouddirectory.CreateObjectOutput, error)
	CreateObjectAsync(ctx workflow.Context, input *clouddirectory.CreateObjectInput) *CloudDirectoryCreateObjectFuture

	CreateSchema(ctx workflow.Context, input *clouddirectory.CreateSchemaInput) (*clouddirectory.CreateSchemaOutput, error)
	CreateSchemaAsync(ctx workflow.Context, input *clouddirectory.CreateSchemaInput) *CloudDirectoryCreateSchemaFuture

	CreateTypedLinkFacet(ctx workflow.Context, input *clouddirectory.CreateTypedLinkFacetInput) (*clouddirectory.CreateTypedLinkFacetOutput, error)
	CreateTypedLinkFacetAsync(ctx workflow.Context, input *clouddirectory.CreateTypedLinkFacetInput) *CloudDirectoryCreateTypedLinkFacetFuture

	DeleteDirectory(ctx workflow.Context, input *clouddirectory.DeleteDirectoryInput) (*clouddirectory.DeleteDirectoryOutput, error)
	DeleteDirectoryAsync(ctx workflow.Context, input *clouddirectory.DeleteDirectoryInput) *CloudDirectoryDeleteDirectoryFuture

	DeleteFacet(ctx workflow.Context, input *clouddirectory.DeleteFacetInput) (*clouddirectory.DeleteFacetOutput, error)
	DeleteFacetAsync(ctx workflow.Context, input *clouddirectory.DeleteFacetInput) *CloudDirectoryDeleteFacetFuture

	DeleteObject(ctx workflow.Context, input *clouddirectory.DeleteObjectInput) (*clouddirectory.DeleteObjectOutput, error)
	DeleteObjectAsync(ctx workflow.Context, input *clouddirectory.DeleteObjectInput) *CloudDirectoryDeleteObjectFuture

	DeleteSchema(ctx workflow.Context, input *clouddirectory.DeleteSchemaInput) (*clouddirectory.DeleteSchemaOutput, error)
	DeleteSchemaAsync(ctx workflow.Context, input *clouddirectory.DeleteSchemaInput) *CloudDirectoryDeleteSchemaFuture

	DeleteTypedLinkFacet(ctx workflow.Context, input *clouddirectory.DeleteTypedLinkFacetInput) (*clouddirectory.DeleteTypedLinkFacetOutput, error)
	DeleteTypedLinkFacetAsync(ctx workflow.Context, input *clouddirectory.DeleteTypedLinkFacetInput) *CloudDirectoryDeleteTypedLinkFacetFuture

	DetachFromIndex(ctx workflow.Context, input *clouddirectory.DetachFromIndexInput) (*clouddirectory.DetachFromIndexOutput, error)
	DetachFromIndexAsync(ctx workflow.Context, input *clouddirectory.DetachFromIndexInput) *CloudDirectoryDetachFromIndexFuture

	DetachObject(ctx workflow.Context, input *clouddirectory.DetachObjectInput) (*clouddirectory.DetachObjectOutput, error)
	DetachObjectAsync(ctx workflow.Context, input *clouddirectory.DetachObjectInput) *CloudDirectoryDetachObjectFuture

	DetachPolicy(ctx workflow.Context, input *clouddirectory.DetachPolicyInput) (*clouddirectory.DetachPolicyOutput, error)
	DetachPolicyAsync(ctx workflow.Context, input *clouddirectory.DetachPolicyInput) *CloudDirectoryDetachPolicyFuture

	DetachTypedLink(ctx workflow.Context, input *clouddirectory.DetachTypedLinkInput) (*clouddirectory.DetachTypedLinkOutput, error)
	DetachTypedLinkAsync(ctx workflow.Context, input *clouddirectory.DetachTypedLinkInput) *CloudDirectoryDetachTypedLinkFuture

	DisableDirectory(ctx workflow.Context, input *clouddirectory.DisableDirectoryInput) (*clouddirectory.DisableDirectoryOutput, error)
	DisableDirectoryAsync(ctx workflow.Context, input *clouddirectory.DisableDirectoryInput) *CloudDirectoryDisableDirectoryFuture

	EnableDirectory(ctx workflow.Context, input *clouddirectory.EnableDirectoryInput) (*clouddirectory.EnableDirectoryOutput, error)
	EnableDirectoryAsync(ctx workflow.Context, input *clouddirectory.EnableDirectoryInput) *CloudDirectoryEnableDirectoryFuture

	GetAppliedSchemaVersion(ctx workflow.Context, input *clouddirectory.GetAppliedSchemaVersionInput) (*clouddirectory.GetAppliedSchemaVersionOutput, error)
	GetAppliedSchemaVersionAsync(ctx workflow.Context, input *clouddirectory.GetAppliedSchemaVersionInput) *CloudDirectoryGetAppliedSchemaVersionFuture

	GetDirectory(ctx workflow.Context, input *clouddirectory.GetDirectoryInput) (*clouddirectory.GetDirectoryOutput, error)
	GetDirectoryAsync(ctx workflow.Context, input *clouddirectory.GetDirectoryInput) *CloudDirectoryGetDirectoryFuture

	GetFacet(ctx workflow.Context, input *clouddirectory.GetFacetInput) (*clouddirectory.GetFacetOutput, error)
	GetFacetAsync(ctx workflow.Context, input *clouddirectory.GetFacetInput) *CloudDirectoryGetFacetFuture

	GetLinkAttributes(ctx workflow.Context, input *clouddirectory.GetLinkAttributesInput) (*clouddirectory.GetLinkAttributesOutput, error)
	GetLinkAttributesAsync(ctx workflow.Context, input *clouddirectory.GetLinkAttributesInput) *CloudDirectoryGetLinkAttributesFuture

	GetObjectAttributes(ctx workflow.Context, input *clouddirectory.GetObjectAttributesInput) (*clouddirectory.GetObjectAttributesOutput, error)
	GetObjectAttributesAsync(ctx workflow.Context, input *clouddirectory.GetObjectAttributesInput) *CloudDirectoryGetObjectAttributesFuture

	GetObjectInformation(ctx workflow.Context, input *clouddirectory.GetObjectInformationInput) (*clouddirectory.GetObjectInformationOutput, error)
	GetObjectInformationAsync(ctx workflow.Context, input *clouddirectory.GetObjectInformationInput) *CloudDirectoryGetObjectInformationFuture

	GetSchemaAsJson(ctx workflow.Context, input *clouddirectory.GetSchemaAsJsonInput) (*clouddirectory.GetSchemaAsJsonOutput, error)
	GetSchemaAsJsonAsync(ctx workflow.Context, input *clouddirectory.GetSchemaAsJsonInput) *CloudDirectoryGetSchemaAsJsonFuture

	GetTypedLinkFacetInformation(ctx workflow.Context, input *clouddirectory.GetTypedLinkFacetInformationInput) (*clouddirectory.GetTypedLinkFacetInformationOutput, error)
	GetTypedLinkFacetInformationAsync(ctx workflow.Context, input *clouddirectory.GetTypedLinkFacetInformationInput) *CloudDirectoryGetTypedLinkFacetInformationFuture

	ListAppliedSchemaArns(ctx workflow.Context, input *clouddirectory.ListAppliedSchemaArnsInput) (*clouddirectory.ListAppliedSchemaArnsOutput, error)
	ListAppliedSchemaArnsAsync(ctx workflow.Context, input *clouddirectory.ListAppliedSchemaArnsInput) *CloudDirectoryListAppliedSchemaArnsFuture

	ListAttachedIndices(ctx workflow.Context, input *clouddirectory.ListAttachedIndicesInput) (*clouddirectory.ListAttachedIndicesOutput, error)
	ListAttachedIndicesAsync(ctx workflow.Context, input *clouddirectory.ListAttachedIndicesInput) *CloudDirectoryListAttachedIndicesFuture

	ListDevelopmentSchemaArns(ctx workflow.Context, input *clouddirectory.ListDevelopmentSchemaArnsInput) (*clouddirectory.ListDevelopmentSchemaArnsOutput, error)
	ListDevelopmentSchemaArnsAsync(ctx workflow.Context, input *clouddirectory.ListDevelopmentSchemaArnsInput) *CloudDirectoryListDevelopmentSchemaArnsFuture

	ListDirectories(ctx workflow.Context, input *clouddirectory.ListDirectoriesInput) (*clouddirectory.ListDirectoriesOutput, error)
	ListDirectoriesAsync(ctx workflow.Context, input *clouddirectory.ListDirectoriesInput) *CloudDirectoryListDirectoriesFuture

	ListFacetAttributes(ctx workflow.Context, input *clouddirectory.ListFacetAttributesInput) (*clouddirectory.ListFacetAttributesOutput, error)
	ListFacetAttributesAsync(ctx workflow.Context, input *clouddirectory.ListFacetAttributesInput) *CloudDirectoryListFacetAttributesFuture

	ListFacetNames(ctx workflow.Context, input *clouddirectory.ListFacetNamesInput) (*clouddirectory.ListFacetNamesOutput, error)
	ListFacetNamesAsync(ctx workflow.Context, input *clouddirectory.ListFacetNamesInput) *CloudDirectoryListFacetNamesFuture

	ListIncomingTypedLinks(ctx workflow.Context, input *clouddirectory.ListIncomingTypedLinksInput) (*clouddirectory.ListIncomingTypedLinksOutput, error)
	ListIncomingTypedLinksAsync(ctx workflow.Context, input *clouddirectory.ListIncomingTypedLinksInput) *CloudDirectoryListIncomingTypedLinksFuture

	ListIndex(ctx workflow.Context, input *clouddirectory.ListIndexInput) (*clouddirectory.ListIndexOutput, error)
	ListIndexAsync(ctx workflow.Context, input *clouddirectory.ListIndexInput) *CloudDirectoryListIndexFuture

	ListManagedSchemaArns(ctx workflow.Context, input *clouddirectory.ListManagedSchemaArnsInput) (*clouddirectory.ListManagedSchemaArnsOutput, error)
	ListManagedSchemaArnsAsync(ctx workflow.Context, input *clouddirectory.ListManagedSchemaArnsInput) *CloudDirectoryListManagedSchemaArnsFuture

	ListObjectAttributes(ctx workflow.Context, input *clouddirectory.ListObjectAttributesInput) (*clouddirectory.ListObjectAttributesOutput, error)
	ListObjectAttributesAsync(ctx workflow.Context, input *clouddirectory.ListObjectAttributesInput) *CloudDirectoryListObjectAttributesFuture

	ListObjectChildren(ctx workflow.Context, input *clouddirectory.ListObjectChildrenInput) (*clouddirectory.ListObjectChildrenOutput, error)
	ListObjectChildrenAsync(ctx workflow.Context, input *clouddirectory.ListObjectChildrenInput) *CloudDirectoryListObjectChildrenFuture

	ListObjectParentPaths(ctx workflow.Context, input *clouddirectory.ListObjectParentPathsInput) (*clouddirectory.ListObjectParentPathsOutput, error)
	ListObjectParentPathsAsync(ctx workflow.Context, input *clouddirectory.ListObjectParentPathsInput) *CloudDirectoryListObjectParentPathsFuture

	ListObjectParents(ctx workflow.Context, input *clouddirectory.ListObjectParentsInput) (*clouddirectory.ListObjectParentsOutput, error)
	ListObjectParentsAsync(ctx workflow.Context, input *clouddirectory.ListObjectParentsInput) *CloudDirectoryListObjectParentsFuture

	ListObjectPolicies(ctx workflow.Context, input *clouddirectory.ListObjectPoliciesInput) (*clouddirectory.ListObjectPoliciesOutput, error)
	ListObjectPoliciesAsync(ctx workflow.Context, input *clouddirectory.ListObjectPoliciesInput) *CloudDirectoryListObjectPoliciesFuture

	ListOutgoingTypedLinks(ctx workflow.Context, input *clouddirectory.ListOutgoingTypedLinksInput) (*clouddirectory.ListOutgoingTypedLinksOutput, error)
	ListOutgoingTypedLinksAsync(ctx workflow.Context, input *clouddirectory.ListOutgoingTypedLinksInput) *CloudDirectoryListOutgoingTypedLinksFuture

	ListPolicyAttachments(ctx workflow.Context, input *clouddirectory.ListPolicyAttachmentsInput) (*clouddirectory.ListPolicyAttachmentsOutput, error)
	ListPolicyAttachmentsAsync(ctx workflow.Context, input *clouddirectory.ListPolicyAttachmentsInput) *CloudDirectoryListPolicyAttachmentsFuture

	ListPublishedSchemaArns(ctx workflow.Context, input *clouddirectory.ListPublishedSchemaArnsInput) (*clouddirectory.ListPublishedSchemaArnsOutput, error)
	ListPublishedSchemaArnsAsync(ctx workflow.Context, input *clouddirectory.ListPublishedSchemaArnsInput) *CloudDirectoryListPublishedSchemaArnsFuture

	ListTagsForResource(ctx workflow.Context, input *clouddirectory.ListTagsForResourceInput) (*clouddirectory.ListTagsForResourceOutput, error)
	ListTagsForResourceAsync(ctx workflow.Context, input *clouddirectory.ListTagsForResourceInput) *CloudDirectoryListTagsForResourceFuture

	ListTypedLinkFacetAttributes(ctx workflow.Context, input *clouddirectory.ListTypedLinkFacetAttributesInput) (*clouddirectory.ListTypedLinkFacetAttributesOutput, error)
	ListTypedLinkFacetAttributesAsync(ctx workflow.Context, input *clouddirectory.ListTypedLinkFacetAttributesInput) *CloudDirectoryListTypedLinkFacetAttributesFuture

	ListTypedLinkFacetNames(ctx workflow.Context, input *clouddirectory.ListTypedLinkFacetNamesInput) (*clouddirectory.ListTypedLinkFacetNamesOutput, error)
	ListTypedLinkFacetNamesAsync(ctx workflow.Context, input *clouddirectory.ListTypedLinkFacetNamesInput) *CloudDirectoryListTypedLinkFacetNamesFuture

	LookupPolicy(ctx workflow.Context, input *clouddirectory.LookupPolicyInput) (*clouddirectory.LookupPolicyOutput, error)
	LookupPolicyAsync(ctx workflow.Context, input *clouddirectory.LookupPolicyInput) *CloudDirectoryLookupPolicyFuture

	PublishSchema(ctx workflow.Context, input *clouddirectory.PublishSchemaInput) (*clouddirectory.PublishSchemaOutput, error)
	PublishSchemaAsync(ctx workflow.Context, input *clouddirectory.PublishSchemaInput) *CloudDirectoryPublishSchemaFuture

	PutSchemaFromJson(ctx workflow.Context, input *clouddirectory.PutSchemaFromJsonInput) (*clouddirectory.PutSchemaFromJsonOutput, error)
	PutSchemaFromJsonAsync(ctx workflow.Context, input *clouddirectory.PutSchemaFromJsonInput) *CloudDirectoryPutSchemaFromJsonFuture

	RemoveFacetFromObject(ctx workflow.Context, input *clouddirectory.RemoveFacetFromObjectInput) (*clouddirectory.RemoveFacetFromObjectOutput, error)
	RemoveFacetFromObjectAsync(ctx workflow.Context, input *clouddirectory.RemoveFacetFromObjectInput) *CloudDirectoryRemoveFacetFromObjectFuture

	TagResource(ctx workflow.Context, input *clouddirectory.TagResourceInput) (*clouddirectory.TagResourceOutput, error)
	TagResourceAsync(ctx workflow.Context, input *clouddirectory.TagResourceInput) *CloudDirectoryTagResourceFuture

	UntagResource(ctx workflow.Context, input *clouddirectory.UntagResourceInput) (*clouddirectory.UntagResourceOutput, error)
	UntagResourceAsync(ctx workflow.Context, input *clouddirectory.UntagResourceInput) *CloudDirectoryUntagResourceFuture

	UpdateFacet(ctx workflow.Context, input *clouddirectory.UpdateFacetInput) (*clouddirectory.UpdateFacetOutput, error)
	UpdateFacetAsync(ctx workflow.Context, input *clouddirectory.UpdateFacetInput) *CloudDirectoryUpdateFacetFuture

	UpdateLinkAttributes(ctx workflow.Context, input *clouddirectory.UpdateLinkAttributesInput) (*clouddirectory.UpdateLinkAttributesOutput, error)
	UpdateLinkAttributesAsync(ctx workflow.Context, input *clouddirectory.UpdateLinkAttributesInput) *CloudDirectoryUpdateLinkAttributesFuture

	UpdateObjectAttributes(ctx workflow.Context, input *clouddirectory.UpdateObjectAttributesInput) (*clouddirectory.UpdateObjectAttributesOutput, error)
	UpdateObjectAttributesAsync(ctx workflow.Context, input *clouddirectory.UpdateObjectAttributesInput) *CloudDirectoryUpdateObjectAttributesFuture

	UpdateSchema(ctx workflow.Context, input *clouddirectory.UpdateSchemaInput) (*clouddirectory.UpdateSchemaOutput, error)
	UpdateSchemaAsync(ctx workflow.Context, input *clouddirectory.UpdateSchemaInput) *CloudDirectoryUpdateSchemaFuture

	UpdateTypedLinkFacet(ctx workflow.Context, input *clouddirectory.UpdateTypedLinkFacetInput) (*clouddirectory.UpdateTypedLinkFacetOutput, error)
	UpdateTypedLinkFacetAsync(ctx workflow.Context, input *clouddirectory.UpdateTypedLinkFacetInput) *CloudDirectoryUpdateTypedLinkFacetFuture

	UpgradeAppliedSchema(ctx workflow.Context, input *clouddirectory.UpgradeAppliedSchemaInput) (*clouddirectory.UpgradeAppliedSchemaOutput, error)
	UpgradeAppliedSchemaAsync(ctx workflow.Context, input *clouddirectory.UpgradeAppliedSchemaInput) *CloudDirectoryUpgradeAppliedSchemaFuture

	UpgradePublishedSchema(ctx workflow.Context, input *clouddirectory.UpgradePublishedSchemaInput) (*clouddirectory.UpgradePublishedSchemaOutput, error)
	UpgradePublishedSchemaAsync(ctx workflow.Context, input *clouddirectory.UpgradePublishedSchemaInput) *CloudDirectoryUpgradePublishedSchemaFuture
}

func NewClient() Client {
	return &stub{}
}
