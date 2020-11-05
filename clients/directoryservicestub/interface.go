// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7

package directoryservicestub

import (
	"github.com/aws/aws-sdk-go/service/directoryservice"
	"go.uber.org/cadence/workflow"

	"github.com/banzaicloud/cadence-aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	AcceptSharedDirectory(ctx workflow.Context, input *directoryservice.AcceptSharedDirectoryInput) (*directoryservice.AcceptSharedDirectoryOutput, error)
	AcceptSharedDirectoryAsync(ctx workflow.Context, input *directoryservice.AcceptSharedDirectoryInput) *DirectoryServiceAcceptSharedDirectoryFuture

	AddIpRoutes(ctx workflow.Context, input *directoryservice.AddIpRoutesInput) (*directoryservice.AddIpRoutesOutput, error)
	AddIpRoutesAsync(ctx workflow.Context, input *directoryservice.AddIpRoutesInput) *DirectoryServiceAddIpRoutesFuture

	AddTagsToResource(ctx workflow.Context, input *directoryservice.AddTagsToResourceInput) (*directoryservice.AddTagsToResourceOutput, error)
	AddTagsToResourceAsync(ctx workflow.Context, input *directoryservice.AddTagsToResourceInput) *DirectoryServiceAddTagsToResourceFuture

	CancelSchemaExtension(ctx workflow.Context, input *directoryservice.CancelSchemaExtensionInput) (*directoryservice.CancelSchemaExtensionOutput, error)
	CancelSchemaExtensionAsync(ctx workflow.Context, input *directoryservice.CancelSchemaExtensionInput) *DirectoryServiceCancelSchemaExtensionFuture

	ConnectDirectory(ctx workflow.Context, input *directoryservice.ConnectDirectoryInput) (*directoryservice.ConnectDirectoryOutput, error)
	ConnectDirectoryAsync(ctx workflow.Context, input *directoryservice.ConnectDirectoryInput) *DirectoryServiceConnectDirectoryFuture

	CreateAlias(ctx workflow.Context, input *directoryservice.CreateAliasInput) (*directoryservice.CreateAliasOutput, error)
	CreateAliasAsync(ctx workflow.Context, input *directoryservice.CreateAliasInput) *DirectoryServiceCreateAliasFuture

	CreateComputer(ctx workflow.Context, input *directoryservice.CreateComputerInput) (*directoryservice.CreateComputerOutput, error)
	CreateComputerAsync(ctx workflow.Context, input *directoryservice.CreateComputerInput) *DirectoryServiceCreateComputerFuture

	CreateConditionalForwarder(ctx workflow.Context, input *directoryservice.CreateConditionalForwarderInput) (*directoryservice.CreateConditionalForwarderOutput, error)
	CreateConditionalForwarderAsync(ctx workflow.Context, input *directoryservice.CreateConditionalForwarderInput) *DirectoryServiceCreateConditionalForwarderFuture

	CreateDirectory(ctx workflow.Context, input *directoryservice.CreateDirectoryInput) (*directoryservice.CreateDirectoryOutput, error)
	CreateDirectoryAsync(ctx workflow.Context, input *directoryservice.CreateDirectoryInput) *DirectoryServiceCreateDirectoryFuture

	CreateLogSubscription(ctx workflow.Context, input *directoryservice.CreateLogSubscriptionInput) (*directoryservice.CreateLogSubscriptionOutput, error)
	CreateLogSubscriptionAsync(ctx workflow.Context, input *directoryservice.CreateLogSubscriptionInput) *DirectoryServiceCreateLogSubscriptionFuture

	CreateMicrosoftAD(ctx workflow.Context, input *directoryservice.CreateMicrosoftADInput) (*directoryservice.CreateMicrosoftADOutput, error)
	CreateMicrosoftADAsync(ctx workflow.Context, input *directoryservice.CreateMicrosoftADInput) *DirectoryServiceCreateMicrosoftADFuture

	CreateSnapshot(ctx workflow.Context, input *directoryservice.CreateSnapshotInput) (*directoryservice.CreateSnapshotOutput, error)
	CreateSnapshotAsync(ctx workflow.Context, input *directoryservice.CreateSnapshotInput) *DirectoryServiceCreateSnapshotFuture

	CreateTrust(ctx workflow.Context, input *directoryservice.CreateTrustInput) (*directoryservice.CreateTrustOutput, error)
	CreateTrustAsync(ctx workflow.Context, input *directoryservice.CreateTrustInput) *DirectoryServiceCreateTrustFuture

	DeleteConditionalForwarder(ctx workflow.Context, input *directoryservice.DeleteConditionalForwarderInput) (*directoryservice.DeleteConditionalForwarderOutput, error)
	DeleteConditionalForwarderAsync(ctx workflow.Context, input *directoryservice.DeleteConditionalForwarderInput) *DirectoryServiceDeleteConditionalForwarderFuture

	DeleteDirectory(ctx workflow.Context, input *directoryservice.DeleteDirectoryInput) (*directoryservice.DeleteDirectoryOutput, error)
	DeleteDirectoryAsync(ctx workflow.Context, input *directoryservice.DeleteDirectoryInput) *DirectoryServiceDeleteDirectoryFuture

	DeleteLogSubscription(ctx workflow.Context, input *directoryservice.DeleteLogSubscriptionInput) (*directoryservice.DeleteLogSubscriptionOutput, error)
	DeleteLogSubscriptionAsync(ctx workflow.Context, input *directoryservice.DeleteLogSubscriptionInput) *DirectoryServiceDeleteLogSubscriptionFuture

	DeleteSnapshot(ctx workflow.Context, input *directoryservice.DeleteSnapshotInput) (*directoryservice.DeleteSnapshotOutput, error)
	DeleteSnapshotAsync(ctx workflow.Context, input *directoryservice.DeleteSnapshotInput) *DirectoryServiceDeleteSnapshotFuture

	DeleteTrust(ctx workflow.Context, input *directoryservice.DeleteTrustInput) (*directoryservice.DeleteTrustOutput, error)
	DeleteTrustAsync(ctx workflow.Context, input *directoryservice.DeleteTrustInput) *DirectoryServiceDeleteTrustFuture

	DeregisterCertificate(ctx workflow.Context, input *directoryservice.DeregisterCertificateInput) (*directoryservice.DeregisterCertificateOutput, error)
	DeregisterCertificateAsync(ctx workflow.Context, input *directoryservice.DeregisterCertificateInput) *DirectoryServiceDeregisterCertificateFuture

	DeregisterEventTopic(ctx workflow.Context, input *directoryservice.DeregisterEventTopicInput) (*directoryservice.DeregisterEventTopicOutput, error)
	DeregisterEventTopicAsync(ctx workflow.Context, input *directoryservice.DeregisterEventTopicInput) *DirectoryServiceDeregisterEventTopicFuture

	DescribeCertificate(ctx workflow.Context, input *directoryservice.DescribeCertificateInput) (*directoryservice.DescribeCertificateOutput, error)
	DescribeCertificateAsync(ctx workflow.Context, input *directoryservice.DescribeCertificateInput) *DirectoryServiceDescribeCertificateFuture

	DescribeConditionalForwarders(ctx workflow.Context, input *directoryservice.DescribeConditionalForwardersInput) (*directoryservice.DescribeConditionalForwardersOutput, error)
	DescribeConditionalForwardersAsync(ctx workflow.Context, input *directoryservice.DescribeConditionalForwardersInput) *DirectoryServiceDescribeConditionalForwardersFuture

	DescribeDirectories(ctx workflow.Context, input *directoryservice.DescribeDirectoriesInput) (*directoryservice.DescribeDirectoriesOutput, error)
	DescribeDirectoriesAsync(ctx workflow.Context, input *directoryservice.DescribeDirectoriesInput) *DirectoryServiceDescribeDirectoriesFuture

	DescribeDomainControllers(ctx workflow.Context, input *directoryservice.DescribeDomainControllersInput) (*directoryservice.DescribeDomainControllersOutput, error)
	DescribeDomainControllersAsync(ctx workflow.Context, input *directoryservice.DescribeDomainControllersInput) *DirectoryServiceDescribeDomainControllersFuture

	DescribeEventTopics(ctx workflow.Context, input *directoryservice.DescribeEventTopicsInput) (*directoryservice.DescribeEventTopicsOutput, error)
	DescribeEventTopicsAsync(ctx workflow.Context, input *directoryservice.DescribeEventTopicsInput) *DirectoryServiceDescribeEventTopicsFuture

	DescribeLDAPSSettings(ctx workflow.Context, input *directoryservice.DescribeLDAPSSettingsInput) (*directoryservice.DescribeLDAPSSettingsOutput, error)
	DescribeLDAPSSettingsAsync(ctx workflow.Context, input *directoryservice.DescribeLDAPSSettingsInput) *DirectoryServiceDescribeLDAPSSettingsFuture

	DescribeSharedDirectories(ctx workflow.Context, input *directoryservice.DescribeSharedDirectoriesInput) (*directoryservice.DescribeSharedDirectoriesOutput, error)
	DescribeSharedDirectoriesAsync(ctx workflow.Context, input *directoryservice.DescribeSharedDirectoriesInput) *DirectoryServiceDescribeSharedDirectoriesFuture

	DescribeSnapshots(ctx workflow.Context, input *directoryservice.DescribeSnapshotsInput) (*directoryservice.DescribeSnapshotsOutput, error)
	DescribeSnapshotsAsync(ctx workflow.Context, input *directoryservice.DescribeSnapshotsInput) *DirectoryServiceDescribeSnapshotsFuture

	DescribeTrusts(ctx workflow.Context, input *directoryservice.DescribeTrustsInput) (*directoryservice.DescribeTrustsOutput, error)
	DescribeTrustsAsync(ctx workflow.Context, input *directoryservice.DescribeTrustsInput) *DirectoryServiceDescribeTrustsFuture

	DisableLDAPS(ctx workflow.Context, input *directoryservice.DisableLDAPSInput) (*directoryservice.DisableLDAPSOutput, error)
	DisableLDAPSAsync(ctx workflow.Context, input *directoryservice.DisableLDAPSInput) *DirectoryServiceDisableLDAPSFuture

	DisableRadius(ctx workflow.Context, input *directoryservice.DisableRadiusInput) (*directoryservice.DisableRadiusOutput, error)
	DisableRadiusAsync(ctx workflow.Context, input *directoryservice.DisableRadiusInput) *DirectoryServiceDisableRadiusFuture

	DisableSso(ctx workflow.Context, input *directoryservice.DisableSsoInput) (*directoryservice.DisableSsoOutput, error)
	DisableSsoAsync(ctx workflow.Context, input *directoryservice.DisableSsoInput) *DirectoryServiceDisableSsoFuture

	EnableLDAPS(ctx workflow.Context, input *directoryservice.EnableLDAPSInput) (*directoryservice.EnableLDAPSOutput, error)
	EnableLDAPSAsync(ctx workflow.Context, input *directoryservice.EnableLDAPSInput) *DirectoryServiceEnableLDAPSFuture

	EnableRadius(ctx workflow.Context, input *directoryservice.EnableRadiusInput) (*directoryservice.EnableRadiusOutput, error)
	EnableRadiusAsync(ctx workflow.Context, input *directoryservice.EnableRadiusInput) *DirectoryServiceEnableRadiusFuture

	EnableSso(ctx workflow.Context, input *directoryservice.EnableSsoInput) (*directoryservice.EnableSsoOutput, error)
	EnableSsoAsync(ctx workflow.Context, input *directoryservice.EnableSsoInput) *DirectoryServiceEnableSsoFuture

	GetDirectoryLimits(ctx workflow.Context, input *directoryservice.GetDirectoryLimitsInput) (*directoryservice.GetDirectoryLimitsOutput, error)
	GetDirectoryLimitsAsync(ctx workflow.Context, input *directoryservice.GetDirectoryLimitsInput) *DirectoryServiceGetDirectoryLimitsFuture

	GetSnapshotLimits(ctx workflow.Context, input *directoryservice.GetSnapshotLimitsInput) (*directoryservice.GetSnapshotLimitsOutput, error)
	GetSnapshotLimitsAsync(ctx workflow.Context, input *directoryservice.GetSnapshotLimitsInput) *DirectoryServiceGetSnapshotLimitsFuture

	ListCertificates(ctx workflow.Context, input *directoryservice.ListCertificatesInput) (*directoryservice.ListCertificatesOutput, error)
	ListCertificatesAsync(ctx workflow.Context, input *directoryservice.ListCertificatesInput) *DirectoryServiceListCertificatesFuture

	ListIpRoutes(ctx workflow.Context, input *directoryservice.ListIpRoutesInput) (*directoryservice.ListIpRoutesOutput, error)
	ListIpRoutesAsync(ctx workflow.Context, input *directoryservice.ListIpRoutesInput) *DirectoryServiceListIpRoutesFuture

	ListLogSubscriptions(ctx workflow.Context, input *directoryservice.ListLogSubscriptionsInput) (*directoryservice.ListLogSubscriptionsOutput, error)
	ListLogSubscriptionsAsync(ctx workflow.Context, input *directoryservice.ListLogSubscriptionsInput) *DirectoryServiceListLogSubscriptionsFuture

	ListSchemaExtensions(ctx workflow.Context, input *directoryservice.ListSchemaExtensionsInput) (*directoryservice.ListSchemaExtensionsOutput, error)
	ListSchemaExtensionsAsync(ctx workflow.Context, input *directoryservice.ListSchemaExtensionsInput) *DirectoryServiceListSchemaExtensionsFuture

	ListTagsForResource(ctx workflow.Context, input *directoryservice.ListTagsForResourceInput) (*directoryservice.ListTagsForResourceOutput, error)
	ListTagsForResourceAsync(ctx workflow.Context, input *directoryservice.ListTagsForResourceInput) *DirectoryServiceListTagsForResourceFuture

	RegisterCertificate(ctx workflow.Context, input *directoryservice.RegisterCertificateInput) (*directoryservice.RegisterCertificateOutput, error)
	RegisterCertificateAsync(ctx workflow.Context, input *directoryservice.RegisterCertificateInput) *DirectoryServiceRegisterCertificateFuture

	RegisterEventTopic(ctx workflow.Context, input *directoryservice.RegisterEventTopicInput) (*directoryservice.RegisterEventTopicOutput, error)
	RegisterEventTopicAsync(ctx workflow.Context, input *directoryservice.RegisterEventTopicInput) *DirectoryServiceRegisterEventTopicFuture

	RejectSharedDirectory(ctx workflow.Context, input *directoryservice.RejectSharedDirectoryInput) (*directoryservice.RejectSharedDirectoryOutput, error)
	RejectSharedDirectoryAsync(ctx workflow.Context, input *directoryservice.RejectSharedDirectoryInput) *DirectoryServiceRejectSharedDirectoryFuture

	RemoveIpRoutes(ctx workflow.Context, input *directoryservice.RemoveIpRoutesInput) (*directoryservice.RemoveIpRoutesOutput, error)
	RemoveIpRoutesAsync(ctx workflow.Context, input *directoryservice.RemoveIpRoutesInput) *DirectoryServiceRemoveIpRoutesFuture

	RemoveTagsFromResource(ctx workflow.Context, input *directoryservice.RemoveTagsFromResourceInput) (*directoryservice.RemoveTagsFromResourceOutput, error)
	RemoveTagsFromResourceAsync(ctx workflow.Context, input *directoryservice.RemoveTagsFromResourceInput) *DirectoryServiceRemoveTagsFromResourceFuture

	ResetUserPassword(ctx workflow.Context, input *directoryservice.ResetUserPasswordInput) (*directoryservice.ResetUserPasswordOutput, error)
	ResetUserPasswordAsync(ctx workflow.Context, input *directoryservice.ResetUserPasswordInput) *DirectoryServiceResetUserPasswordFuture

	RestoreFromSnapshot(ctx workflow.Context, input *directoryservice.RestoreFromSnapshotInput) (*directoryservice.RestoreFromSnapshotOutput, error)
	RestoreFromSnapshotAsync(ctx workflow.Context, input *directoryservice.RestoreFromSnapshotInput) *DirectoryServiceRestoreFromSnapshotFuture

	ShareDirectory(ctx workflow.Context, input *directoryservice.ShareDirectoryInput) (*directoryservice.ShareDirectoryOutput, error)
	ShareDirectoryAsync(ctx workflow.Context, input *directoryservice.ShareDirectoryInput) *DirectoryServiceShareDirectoryFuture

	StartSchemaExtension(ctx workflow.Context, input *directoryservice.StartSchemaExtensionInput) (*directoryservice.StartSchemaExtensionOutput, error)
	StartSchemaExtensionAsync(ctx workflow.Context, input *directoryservice.StartSchemaExtensionInput) *DirectoryServiceStartSchemaExtensionFuture

	UnshareDirectory(ctx workflow.Context, input *directoryservice.UnshareDirectoryInput) (*directoryservice.UnshareDirectoryOutput, error)
	UnshareDirectoryAsync(ctx workflow.Context, input *directoryservice.UnshareDirectoryInput) *DirectoryServiceUnshareDirectoryFuture

	UpdateConditionalForwarder(ctx workflow.Context, input *directoryservice.UpdateConditionalForwarderInput) (*directoryservice.UpdateConditionalForwarderOutput, error)
	UpdateConditionalForwarderAsync(ctx workflow.Context, input *directoryservice.UpdateConditionalForwarderInput) *DirectoryServiceUpdateConditionalForwarderFuture

	UpdateNumberOfDomainControllers(ctx workflow.Context, input *directoryservice.UpdateNumberOfDomainControllersInput) (*directoryservice.UpdateNumberOfDomainControllersOutput, error)
	UpdateNumberOfDomainControllersAsync(ctx workflow.Context, input *directoryservice.UpdateNumberOfDomainControllersInput) *DirectoryServiceUpdateNumberOfDomainControllersFuture

	UpdateRadius(ctx workflow.Context, input *directoryservice.UpdateRadiusInput) (*directoryservice.UpdateRadiusOutput, error)
	UpdateRadiusAsync(ctx workflow.Context, input *directoryservice.UpdateRadiusInput) *DirectoryServiceUpdateRadiusFuture

	UpdateTrust(ctx workflow.Context, input *directoryservice.UpdateTrustInput) (*directoryservice.UpdateTrustOutput, error)
	UpdateTrustAsync(ctx workflow.Context, input *directoryservice.UpdateTrustInput) *DirectoryServiceUpdateTrustFuture

	VerifyTrust(ctx workflow.Context, input *directoryservice.VerifyTrustInput) (*directoryservice.VerifyTrustOutput, error)
	VerifyTrustAsync(ctx workflow.Context, input *directoryservice.VerifyTrustInput) *DirectoryServiceVerifyTrustFuture
}

func NewClient() Client {
	return &stub{}
}
