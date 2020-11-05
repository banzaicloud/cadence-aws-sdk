// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7

package lexmodelbuildingservicestub

import (
	"github.com/aws/aws-sdk-go/service/lexmodelbuildingservice"
	"go.uber.org/cadence/workflow"

	"github.com/banzaicloud/cadence-aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	CreateBotVersion(ctx workflow.Context, input *lexmodelbuildingservice.CreateBotVersionInput) (*lexmodelbuildingservice.CreateBotVersionOutput, error)
	CreateBotVersionAsync(ctx workflow.Context, input *lexmodelbuildingservice.CreateBotVersionInput) *LexModelBuildingServiceCreateBotVersionFuture

	CreateIntentVersion(ctx workflow.Context, input *lexmodelbuildingservice.CreateIntentVersionInput) (*lexmodelbuildingservice.CreateIntentVersionOutput, error)
	CreateIntentVersionAsync(ctx workflow.Context, input *lexmodelbuildingservice.CreateIntentVersionInput) *LexModelBuildingServiceCreateIntentVersionFuture

	CreateSlotTypeVersion(ctx workflow.Context, input *lexmodelbuildingservice.CreateSlotTypeVersionInput) (*lexmodelbuildingservice.CreateSlotTypeVersionOutput, error)
	CreateSlotTypeVersionAsync(ctx workflow.Context, input *lexmodelbuildingservice.CreateSlotTypeVersionInput) *LexModelBuildingServiceCreateSlotTypeVersionFuture

	DeleteBot(ctx workflow.Context, input *lexmodelbuildingservice.DeleteBotInput) (*lexmodelbuildingservice.DeleteBotOutput, error)
	DeleteBotAsync(ctx workflow.Context, input *lexmodelbuildingservice.DeleteBotInput) *LexModelBuildingServiceDeleteBotFuture

	DeleteBotAlias(ctx workflow.Context, input *lexmodelbuildingservice.DeleteBotAliasInput) (*lexmodelbuildingservice.DeleteBotAliasOutput, error)
	DeleteBotAliasAsync(ctx workflow.Context, input *lexmodelbuildingservice.DeleteBotAliasInput) *LexModelBuildingServiceDeleteBotAliasFuture

	DeleteBotChannelAssociation(ctx workflow.Context, input *lexmodelbuildingservice.DeleteBotChannelAssociationInput) (*lexmodelbuildingservice.DeleteBotChannelAssociationOutput, error)
	DeleteBotChannelAssociationAsync(ctx workflow.Context, input *lexmodelbuildingservice.DeleteBotChannelAssociationInput) *LexModelBuildingServiceDeleteBotChannelAssociationFuture

	DeleteBotVersion(ctx workflow.Context, input *lexmodelbuildingservice.DeleteBotVersionInput) (*lexmodelbuildingservice.DeleteBotVersionOutput, error)
	DeleteBotVersionAsync(ctx workflow.Context, input *lexmodelbuildingservice.DeleteBotVersionInput) *LexModelBuildingServiceDeleteBotVersionFuture

	DeleteIntent(ctx workflow.Context, input *lexmodelbuildingservice.DeleteIntentInput) (*lexmodelbuildingservice.DeleteIntentOutput, error)
	DeleteIntentAsync(ctx workflow.Context, input *lexmodelbuildingservice.DeleteIntentInput) *LexModelBuildingServiceDeleteIntentFuture

	DeleteIntentVersion(ctx workflow.Context, input *lexmodelbuildingservice.DeleteIntentVersionInput) (*lexmodelbuildingservice.DeleteIntentVersionOutput, error)
	DeleteIntentVersionAsync(ctx workflow.Context, input *lexmodelbuildingservice.DeleteIntentVersionInput) *LexModelBuildingServiceDeleteIntentVersionFuture

	DeleteSlotType(ctx workflow.Context, input *lexmodelbuildingservice.DeleteSlotTypeInput) (*lexmodelbuildingservice.DeleteSlotTypeOutput, error)
	DeleteSlotTypeAsync(ctx workflow.Context, input *lexmodelbuildingservice.DeleteSlotTypeInput) *LexModelBuildingServiceDeleteSlotTypeFuture

	DeleteSlotTypeVersion(ctx workflow.Context, input *lexmodelbuildingservice.DeleteSlotTypeVersionInput) (*lexmodelbuildingservice.DeleteSlotTypeVersionOutput, error)
	DeleteSlotTypeVersionAsync(ctx workflow.Context, input *lexmodelbuildingservice.DeleteSlotTypeVersionInput) *LexModelBuildingServiceDeleteSlotTypeVersionFuture

	DeleteUtterances(ctx workflow.Context, input *lexmodelbuildingservice.DeleteUtterancesInput) (*lexmodelbuildingservice.DeleteUtterancesOutput, error)
	DeleteUtterancesAsync(ctx workflow.Context, input *lexmodelbuildingservice.DeleteUtterancesInput) *LexModelBuildingServiceDeleteUtterancesFuture

	GetBot(ctx workflow.Context, input *lexmodelbuildingservice.GetBotInput) (*lexmodelbuildingservice.GetBotOutput, error)
	GetBotAsync(ctx workflow.Context, input *lexmodelbuildingservice.GetBotInput) *LexModelBuildingServiceGetBotFuture

	GetBotAlias(ctx workflow.Context, input *lexmodelbuildingservice.GetBotAliasInput) (*lexmodelbuildingservice.GetBotAliasOutput, error)
	GetBotAliasAsync(ctx workflow.Context, input *lexmodelbuildingservice.GetBotAliasInput) *LexModelBuildingServiceGetBotAliasFuture

	GetBotAliases(ctx workflow.Context, input *lexmodelbuildingservice.GetBotAliasesInput) (*lexmodelbuildingservice.GetBotAliasesOutput, error)
	GetBotAliasesAsync(ctx workflow.Context, input *lexmodelbuildingservice.GetBotAliasesInput) *LexModelBuildingServiceGetBotAliasesFuture

	GetBotChannelAssociation(ctx workflow.Context, input *lexmodelbuildingservice.GetBotChannelAssociationInput) (*lexmodelbuildingservice.GetBotChannelAssociationOutput, error)
	GetBotChannelAssociationAsync(ctx workflow.Context, input *lexmodelbuildingservice.GetBotChannelAssociationInput) *LexModelBuildingServiceGetBotChannelAssociationFuture

	GetBotChannelAssociations(ctx workflow.Context, input *lexmodelbuildingservice.GetBotChannelAssociationsInput) (*lexmodelbuildingservice.GetBotChannelAssociationsOutput, error)
	GetBotChannelAssociationsAsync(ctx workflow.Context, input *lexmodelbuildingservice.GetBotChannelAssociationsInput) *LexModelBuildingServiceGetBotChannelAssociationsFuture

	GetBotVersions(ctx workflow.Context, input *lexmodelbuildingservice.GetBotVersionsInput) (*lexmodelbuildingservice.GetBotVersionsOutput, error)
	GetBotVersionsAsync(ctx workflow.Context, input *lexmodelbuildingservice.GetBotVersionsInput) *LexModelBuildingServiceGetBotVersionsFuture

	GetBots(ctx workflow.Context, input *lexmodelbuildingservice.GetBotsInput) (*lexmodelbuildingservice.GetBotsOutput, error)
	GetBotsAsync(ctx workflow.Context, input *lexmodelbuildingservice.GetBotsInput) *LexModelBuildingServiceGetBotsFuture

	GetBuiltinIntent(ctx workflow.Context, input *lexmodelbuildingservice.GetBuiltinIntentInput) (*lexmodelbuildingservice.GetBuiltinIntentOutput, error)
	GetBuiltinIntentAsync(ctx workflow.Context, input *lexmodelbuildingservice.GetBuiltinIntentInput) *LexModelBuildingServiceGetBuiltinIntentFuture

	GetBuiltinIntents(ctx workflow.Context, input *lexmodelbuildingservice.GetBuiltinIntentsInput) (*lexmodelbuildingservice.GetBuiltinIntentsOutput, error)
	GetBuiltinIntentsAsync(ctx workflow.Context, input *lexmodelbuildingservice.GetBuiltinIntentsInput) *LexModelBuildingServiceGetBuiltinIntentsFuture

	GetBuiltinSlotTypes(ctx workflow.Context, input *lexmodelbuildingservice.GetBuiltinSlotTypesInput) (*lexmodelbuildingservice.GetBuiltinSlotTypesOutput, error)
	GetBuiltinSlotTypesAsync(ctx workflow.Context, input *lexmodelbuildingservice.GetBuiltinSlotTypesInput) *LexModelBuildingServiceGetBuiltinSlotTypesFuture

	GetExport(ctx workflow.Context, input *lexmodelbuildingservice.GetExportInput) (*lexmodelbuildingservice.GetExportOutput, error)
	GetExportAsync(ctx workflow.Context, input *lexmodelbuildingservice.GetExportInput) *LexModelBuildingServiceGetExportFuture

	GetImport(ctx workflow.Context, input *lexmodelbuildingservice.GetImportInput) (*lexmodelbuildingservice.GetImportOutput, error)
	GetImportAsync(ctx workflow.Context, input *lexmodelbuildingservice.GetImportInput) *LexModelBuildingServiceGetImportFuture

	GetIntent(ctx workflow.Context, input *lexmodelbuildingservice.GetIntentInput) (*lexmodelbuildingservice.GetIntentOutput, error)
	GetIntentAsync(ctx workflow.Context, input *lexmodelbuildingservice.GetIntentInput) *LexModelBuildingServiceGetIntentFuture

	GetIntentVersions(ctx workflow.Context, input *lexmodelbuildingservice.GetIntentVersionsInput) (*lexmodelbuildingservice.GetIntentVersionsOutput, error)
	GetIntentVersionsAsync(ctx workflow.Context, input *lexmodelbuildingservice.GetIntentVersionsInput) *LexModelBuildingServiceGetIntentVersionsFuture

	GetIntents(ctx workflow.Context, input *lexmodelbuildingservice.GetIntentsInput) (*lexmodelbuildingservice.GetIntentsOutput, error)
	GetIntentsAsync(ctx workflow.Context, input *lexmodelbuildingservice.GetIntentsInput) *LexModelBuildingServiceGetIntentsFuture

	GetSlotType(ctx workflow.Context, input *lexmodelbuildingservice.GetSlotTypeInput) (*lexmodelbuildingservice.GetSlotTypeOutput, error)
	GetSlotTypeAsync(ctx workflow.Context, input *lexmodelbuildingservice.GetSlotTypeInput) *LexModelBuildingServiceGetSlotTypeFuture

	GetSlotTypeVersions(ctx workflow.Context, input *lexmodelbuildingservice.GetSlotTypeVersionsInput) (*lexmodelbuildingservice.GetSlotTypeVersionsOutput, error)
	GetSlotTypeVersionsAsync(ctx workflow.Context, input *lexmodelbuildingservice.GetSlotTypeVersionsInput) *LexModelBuildingServiceGetSlotTypeVersionsFuture

	GetSlotTypes(ctx workflow.Context, input *lexmodelbuildingservice.GetSlotTypesInput) (*lexmodelbuildingservice.GetSlotTypesOutput, error)
	GetSlotTypesAsync(ctx workflow.Context, input *lexmodelbuildingservice.GetSlotTypesInput) *LexModelBuildingServiceGetSlotTypesFuture

	GetUtterancesView(ctx workflow.Context, input *lexmodelbuildingservice.GetUtterancesViewInput) (*lexmodelbuildingservice.GetUtterancesViewOutput, error)
	GetUtterancesViewAsync(ctx workflow.Context, input *lexmodelbuildingservice.GetUtterancesViewInput) *LexModelBuildingServiceGetUtterancesViewFuture

	ListTagsForResource(ctx workflow.Context, input *lexmodelbuildingservice.ListTagsForResourceInput) (*lexmodelbuildingservice.ListTagsForResourceOutput, error)
	ListTagsForResourceAsync(ctx workflow.Context, input *lexmodelbuildingservice.ListTagsForResourceInput) *LexModelBuildingServiceListTagsForResourceFuture

	PutBot(ctx workflow.Context, input *lexmodelbuildingservice.PutBotInput) (*lexmodelbuildingservice.PutBotOutput, error)
	PutBotAsync(ctx workflow.Context, input *lexmodelbuildingservice.PutBotInput) *LexModelBuildingServicePutBotFuture

	PutBotAlias(ctx workflow.Context, input *lexmodelbuildingservice.PutBotAliasInput) (*lexmodelbuildingservice.PutBotAliasOutput, error)
	PutBotAliasAsync(ctx workflow.Context, input *lexmodelbuildingservice.PutBotAliasInput) *LexModelBuildingServicePutBotAliasFuture

	PutIntent(ctx workflow.Context, input *lexmodelbuildingservice.PutIntentInput) (*lexmodelbuildingservice.PutIntentOutput, error)
	PutIntentAsync(ctx workflow.Context, input *lexmodelbuildingservice.PutIntentInput) *LexModelBuildingServicePutIntentFuture

	PutSlotType(ctx workflow.Context, input *lexmodelbuildingservice.PutSlotTypeInput) (*lexmodelbuildingservice.PutSlotTypeOutput, error)
	PutSlotTypeAsync(ctx workflow.Context, input *lexmodelbuildingservice.PutSlotTypeInput) *LexModelBuildingServicePutSlotTypeFuture

	StartImport(ctx workflow.Context, input *lexmodelbuildingservice.StartImportInput) (*lexmodelbuildingservice.StartImportOutput, error)
	StartImportAsync(ctx workflow.Context, input *lexmodelbuildingservice.StartImportInput) *LexModelBuildingServiceStartImportFuture

	TagResource(ctx workflow.Context, input *lexmodelbuildingservice.TagResourceInput) (*lexmodelbuildingservice.TagResourceOutput, error)
	TagResourceAsync(ctx workflow.Context, input *lexmodelbuildingservice.TagResourceInput) *LexModelBuildingServiceTagResourceFuture

	UntagResource(ctx workflow.Context, input *lexmodelbuildingservice.UntagResourceInput) (*lexmodelbuildingservice.UntagResourceOutput, error)
	UntagResourceAsync(ctx workflow.Context, input *lexmodelbuildingservice.UntagResourceInput) *LexModelBuildingServiceUntagResourceFuture
}

func NewClient() Client {
	return &stub{}
}
