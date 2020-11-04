// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.
// Copyright (c) 2020 Banzai Cloud All rights reserved.

package elasticbeanstalkstub

import (
	"github.com/aws/aws-sdk-go/service/elasticbeanstalk"
	"go.uber.org/cadence/workflow"

	"github.com/banzaicloud/cadence-aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	AbortEnvironmentUpdate(ctx workflow.Context, input *elasticbeanstalk.AbortEnvironmentUpdateInput) (*elasticbeanstalk.AbortEnvironmentUpdateOutput, error)
	AbortEnvironmentUpdateAsync(ctx workflow.Context, input *elasticbeanstalk.AbortEnvironmentUpdateInput) *ElasticBeanstalkAbortEnvironmentUpdateFuture

	ApplyEnvironmentManagedAction(ctx workflow.Context, input *elasticbeanstalk.ApplyEnvironmentManagedActionInput) (*elasticbeanstalk.ApplyEnvironmentManagedActionOutput, error)
	ApplyEnvironmentManagedActionAsync(ctx workflow.Context, input *elasticbeanstalk.ApplyEnvironmentManagedActionInput) *ElasticBeanstalkApplyEnvironmentManagedActionFuture

	AssociateEnvironmentOperationsRole(ctx workflow.Context, input *elasticbeanstalk.AssociateEnvironmentOperationsRoleInput) (*elasticbeanstalk.AssociateEnvironmentOperationsRoleOutput, error)
	AssociateEnvironmentOperationsRoleAsync(ctx workflow.Context, input *elasticbeanstalk.AssociateEnvironmentOperationsRoleInput) *ElasticBeanstalkAssociateEnvironmentOperationsRoleFuture

	CheckDNSAvailability(ctx workflow.Context, input *elasticbeanstalk.CheckDNSAvailabilityInput) (*elasticbeanstalk.CheckDNSAvailabilityOutput, error)
	CheckDNSAvailabilityAsync(ctx workflow.Context, input *elasticbeanstalk.CheckDNSAvailabilityInput) *ElasticBeanstalkCheckDNSAvailabilityFuture

	ComposeEnvironments(ctx workflow.Context, input *elasticbeanstalk.ComposeEnvironmentsInput) (*elasticbeanstalk.EnvironmentDescriptionsMessage, error)
	ComposeEnvironmentsAsync(ctx workflow.Context, input *elasticbeanstalk.ComposeEnvironmentsInput) *ElasticBeanstalkComposeEnvironmentsFuture

	CreateApplication(ctx workflow.Context, input *elasticbeanstalk.CreateApplicationInput) (*elasticbeanstalk.ApplicationDescriptionMessage, error)
	CreateApplicationAsync(ctx workflow.Context, input *elasticbeanstalk.CreateApplicationInput) *ElasticBeanstalkCreateApplicationFuture

	CreateApplicationVersion(ctx workflow.Context, input *elasticbeanstalk.CreateApplicationVersionInput) (*elasticbeanstalk.ApplicationVersionDescriptionMessage, error)
	CreateApplicationVersionAsync(ctx workflow.Context, input *elasticbeanstalk.CreateApplicationVersionInput) *ElasticBeanstalkCreateApplicationVersionFuture

	CreateConfigurationTemplate(ctx workflow.Context, input *elasticbeanstalk.CreateConfigurationTemplateInput) (*elasticbeanstalk.ConfigurationSettingsDescription, error)
	CreateConfigurationTemplateAsync(ctx workflow.Context, input *elasticbeanstalk.CreateConfigurationTemplateInput) *ElasticBeanstalkCreateConfigurationTemplateFuture

	CreateEnvironment(ctx workflow.Context, input *elasticbeanstalk.CreateEnvironmentInput) (*elasticbeanstalk.EnvironmentDescription, error)
	CreateEnvironmentAsync(ctx workflow.Context, input *elasticbeanstalk.CreateEnvironmentInput) *ElasticBeanstalkCreateEnvironmentFuture

	CreatePlatformVersion(ctx workflow.Context, input *elasticbeanstalk.CreatePlatformVersionInput) (*elasticbeanstalk.CreatePlatformVersionOutput, error)
	CreatePlatformVersionAsync(ctx workflow.Context, input *elasticbeanstalk.CreatePlatformVersionInput) *ElasticBeanstalkCreatePlatformVersionFuture

	CreateStorageLocation(ctx workflow.Context, input *elasticbeanstalk.CreateStorageLocationInput) (*elasticbeanstalk.CreateStorageLocationOutput, error)
	CreateStorageLocationAsync(ctx workflow.Context, input *elasticbeanstalk.CreateStorageLocationInput) *ElasticBeanstalkCreateStorageLocationFuture

	DeleteApplication(ctx workflow.Context, input *elasticbeanstalk.DeleteApplicationInput) (*elasticbeanstalk.DeleteApplicationOutput, error)
	DeleteApplicationAsync(ctx workflow.Context, input *elasticbeanstalk.DeleteApplicationInput) *ElasticBeanstalkDeleteApplicationFuture

	DeleteApplicationVersion(ctx workflow.Context, input *elasticbeanstalk.DeleteApplicationVersionInput) (*elasticbeanstalk.DeleteApplicationVersionOutput, error)
	DeleteApplicationVersionAsync(ctx workflow.Context, input *elasticbeanstalk.DeleteApplicationVersionInput) *ElasticBeanstalkDeleteApplicationVersionFuture

	DeleteConfigurationTemplate(ctx workflow.Context, input *elasticbeanstalk.DeleteConfigurationTemplateInput) (*elasticbeanstalk.DeleteConfigurationTemplateOutput, error)
	DeleteConfigurationTemplateAsync(ctx workflow.Context, input *elasticbeanstalk.DeleteConfigurationTemplateInput) *ElasticBeanstalkDeleteConfigurationTemplateFuture

	DeleteEnvironmentConfiguration(ctx workflow.Context, input *elasticbeanstalk.DeleteEnvironmentConfigurationInput) (*elasticbeanstalk.DeleteEnvironmentConfigurationOutput, error)
	DeleteEnvironmentConfigurationAsync(ctx workflow.Context, input *elasticbeanstalk.DeleteEnvironmentConfigurationInput) *ElasticBeanstalkDeleteEnvironmentConfigurationFuture

	DeletePlatformVersion(ctx workflow.Context, input *elasticbeanstalk.DeletePlatformVersionInput) (*elasticbeanstalk.DeletePlatformVersionOutput, error)
	DeletePlatformVersionAsync(ctx workflow.Context, input *elasticbeanstalk.DeletePlatformVersionInput) *ElasticBeanstalkDeletePlatformVersionFuture

	DescribeAccountAttributes(ctx workflow.Context, input *elasticbeanstalk.DescribeAccountAttributesInput) (*elasticbeanstalk.DescribeAccountAttributesOutput, error)
	DescribeAccountAttributesAsync(ctx workflow.Context, input *elasticbeanstalk.DescribeAccountAttributesInput) *ElasticBeanstalkDescribeAccountAttributesFuture

	DescribeApplicationVersions(ctx workflow.Context, input *elasticbeanstalk.DescribeApplicationVersionsInput) (*elasticbeanstalk.DescribeApplicationVersionsOutput, error)
	DescribeApplicationVersionsAsync(ctx workflow.Context, input *elasticbeanstalk.DescribeApplicationVersionsInput) *ElasticBeanstalkDescribeApplicationVersionsFuture

	DescribeApplications(ctx workflow.Context, input *elasticbeanstalk.DescribeApplicationsInput) (*elasticbeanstalk.DescribeApplicationsOutput, error)
	DescribeApplicationsAsync(ctx workflow.Context, input *elasticbeanstalk.DescribeApplicationsInput) *ElasticBeanstalkDescribeApplicationsFuture

	DescribeConfigurationOptions(ctx workflow.Context, input *elasticbeanstalk.DescribeConfigurationOptionsInput) (*elasticbeanstalk.DescribeConfigurationOptionsOutput, error)
	DescribeConfigurationOptionsAsync(ctx workflow.Context, input *elasticbeanstalk.DescribeConfigurationOptionsInput) *ElasticBeanstalkDescribeConfigurationOptionsFuture

	DescribeConfigurationSettings(ctx workflow.Context, input *elasticbeanstalk.DescribeConfigurationSettingsInput) (*elasticbeanstalk.DescribeConfigurationSettingsOutput, error)
	DescribeConfigurationSettingsAsync(ctx workflow.Context, input *elasticbeanstalk.DescribeConfigurationSettingsInput) *ElasticBeanstalkDescribeConfigurationSettingsFuture

	DescribeEnvironmentHealth(ctx workflow.Context, input *elasticbeanstalk.DescribeEnvironmentHealthInput) (*elasticbeanstalk.DescribeEnvironmentHealthOutput, error)
	DescribeEnvironmentHealthAsync(ctx workflow.Context, input *elasticbeanstalk.DescribeEnvironmentHealthInput) *ElasticBeanstalkDescribeEnvironmentHealthFuture

	DescribeEnvironmentManagedActionHistory(ctx workflow.Context, input *elasticbeanstalk.DescribeEnvironmentManagedActionHistoryInput) (*elasticbeanstalk.DescribeEnvironmentManagedActionHistoryOutput, error)
	DescribeEnvironmentManagedActionHistoryAsync(ctx workflow.Context, input *elasticbeanstalk.DescribeEnvironmentManagedActionHistoryInput) *ElasticBeanstalkDescribeEnvironmentManagedActionHistoryFuture

	DescribeEnvironmentManagedActions(ctx workflow.Context, input *elasticbeanstalk.DescribeEnvironmentManagedActionsInput) (*elasticbeanstalk.DescribeEnvironmentManagedActionsOutput, error)
	DescribeEnvironmentManagedActionsAsync(ctx workflow.Context, input *elasticbeanstalk.DescribeEnvironmentManagedActionsInput) *ElasticBeanstalkDescribeEnvironmentManagedActionsFuture

	DescribeEnvironmentResources(ctx workflow.Context, input *elasticbeanstalk.DescribeEnvironmentResourcesInput) (*elasticbeanstalk.DescribeEnvironmentResourcesOutput, error)
	DescribeEnvironmentResourcesAsync(ctx workflow.Context, input *elasticbeanstalk.DescribeEnvironmentResourcesInput) *ElasticBeanstalkDescribeEnvironmentResourcesFuture

	DescribeEnvironments(ctx workflow.Context, input *elasticbeanstalk.DescribeEnvironmentsInput) (*elasticbeanstalk.EnvironmentDescriptionsMessage, error)
	DescribeEnvironmentsAsync(ctx workflow.Context, input *elasticbeanstalk.DescribeEnvironmentsInput) *ElasticBeanstalkDescribeEnvironmentsFuture

	DescribeEvents(ctx workflow.Context, input *elasticbeanstalk.DescribeEventsInput) (*elasticbeanstalk.DescribeEventsOutput, error)
	DescribeEventsAsync(ctx workflow.Context, input *elasticbeanstalk.DescribeEventsInput) *ElasticBeanstalkDescribeEventsFuture

	DescribeInstancesHealth(ctx workflow.Context, input *elasticbeanstalk.DescribeInstancesHealthInput) (*elasticbeanstalk.DescribeInstancesHealthOutput, error)
	DescribeInstancesHealthAsync(ctx workflow.Context, input *elasticbeanstalk.DescribeInstancesHealthInput) *ElasticBeanstalkDescribeInstancesHealthFuture

	DescribePlatformVersion(ctx workflow.Context, input *elasticbeanstalk.DescribePlatformVersionInput) (*elasticbeanstalk.DescribePlatformVersionOutput, error)
	DescribePlatformVersionAsync(ctx workflow.Context, input *elasticbeanstalk.DescribePlatformVersionInput) *ElasticBeanstalkDescribePlatformVersionFuture

	DisassociateEnvironmentOperationsRole(ctx workflow.Context, input *elasticbeanstalk.DisassociateEnvironmentOperationsRoleInput) (*elasticbeanstalk.DisassociateEnvironmentOperationsRoleOutput, error)
	DisassociateEnvironmentOperationsRoleAsync(ctx workflow.Context, input *elasticbeanstalk.DisassociateEnvironmentOperationsRoleInput) *ElasticBeanstalkDisassociateEnvironmentOperationsRoleFuture

	ListAvailableSolutionStacks(ctx workflow.Context, input *elasticbeanstalk.ListAvailableSolutionStacksInput) (*elasticbeanstalk.ListAvailableSolutionStacksOutput, error)
	ListAvailableSolutionStacksAsync(ctx workflow.Context, input *elasticbeanstalk.ListAvailableSolutionStacksInput) *ElasticBeanstalkListAvailableSolutionStacksFuture

	ListPlatformBranches(ctx workflow.Context, input *elasticbeanstalk.ListPlatformBranchesInput) (*elasticbeanstalk.ListPlatformBranchesOutput, error)
	ListPlatformBranchesAsync(ctx workflow.Context, input *elasticbeanstalk.ListPlatformBranchesInput) *ElasticBeanstalkListPlatformBranchesFuture

	ListPlatformVersions(ctx workflow.Context, input *elasticbeanstalk.ListPlatformVersionsInput) (*elasticbeanstalk.ListPlatformVersionsOutput, error)
	ListPlatformVersionsAsync(ctx workflow.Context, input *elasticbeanstalk.ListPlatformVersionsInput) *ElasticBeanstalkListPlatformVersionsFuture

	ListTagsForResource(ctx workflow.Context, input *elasticbeanstalk.ListTagsForResourceInput) (*elasticbeanstalk.ListTagsForResourceOutput, error)
	ListTagsForResourceAsync(ctx workflow.Context, input *elasticbeanstalk.ListTagsForResourceInput) *ElasticBeanstalkListTagsForResourceFuture

	RebuildEnvironment(ctx workflow.Context, input *elasticbeanstalk.RebuildEnvironmentInput) (*elasticbeanstalk.RebuildEnvironmentOutput, error)
	RebuildEnvironmentAsync(ctx workflow.Context, input *elasticbeanstalk.RebuildEnvironmentInput) *ElasticBeanstalkRebuildEnvironmentFuture

	RequestEnvironmentInfo(ctx workflow.Context, input *elasticbeanstalk.RequestEnvironmentInfoInput) (*elasticbeanstalk.RequestEnvironmentInfoOutput, error)
	RequestEnvironmentInfoAsync(ctx workflow.Context, input *elasticbeanstalk.RequestEnvironmentInfoInput) *ElasticBeanstalkRequestEnvironmentInfoFuture

	RestartAppServer(ctx workflow.Context, input *elasticbeanstalk.RestartAppServerInput) (*elasticbeanstalk.RestartAppServerOutput, error)
	RestartAppServerAsync(ctx workflow.Context, input *elasticbeanstalk.RestartAppServerInput) *ElasticBeanstalkRestartAppServerFuture

	RetrieveEnvironmentInfo(ctx workflow.Context, input *elasticbeanstalk.RetrieveEnvironmentInfoInput) (*elasticbeanstalk.RetrieveEnvironmentInfoOutput, error)
	RetrieveEnvironmentInfoAsync(ctx workflow.Context, input *elasticbeanstalk.RetrieveEnvironmentInfoInput) *ElasticBeanstalkRetrieveEnvironmentInfoFuture

	SwapEnvironmentCNAMEs(ctx workflow.Context, input *elasticbeanstalk.SwapEnvironmentCNAMEsInput) (*elasticbeanstalk.SwapEnvironmentCNAMEsOutput, error)
	SwapEnvironmentCNAMEsAsync(ctx workflow.Context, input *elasticbeanstalk.SwapEnvironmentCNAMEsInput) *ElasticBeanstalkSwapEnvironmentCNAMEsFuture

	TerminateEnvironment(ctx workflow.Context, input *elasticbeanstalk.TerminateEnvironmentInput) (*elasticbeanstalk.EnvironmentDescription, error)
	TerminateEnvironmentAsync(ctx workflow.Context, input *elasticbeanstalk.TerminateEnvironmentInput) *ElasticBeanstalkTerminateEnvironmentFuture

	UpdateApplication(ctx workflow.Context, input *elasticbeanstalk.UpdateApplicationInput) (*elasticbeanstalk.ApplicationDescriptionMessage, error)
	UpdateApplicationAsync(ctx workflow.Context, input *elasticbeanstalk.UpdateApplicationInput) *ElasticBeanstalkUpdateApplicationFuture

	UpdateApplicationResourceLifecycle(ctx workflow.Context, input *elasticbeanstalk.UpdateApplicationResourceLifecycleInput) (*elasticbeanstalk.UpdateApplicationResourceLifecycleOutput, error)
	UpdateApplicationResourceLifecycleAsync(ctx workflow.Context, input *elasticbeanstalk.UpdateApplicationResourceLifecycleInput) *ElasticBeanstalkUpdateApplicationResourceLifecycleFuture

	UpdateApplicationVersion(ctx workflow.Context, input *elasticbeanstalk.UpdateApplicationVersionInput) (*elasticbeanstalk.ApplicationVersionDescriptionMessage, error)
	UpdateApplicationVersionAsync(ctx workflow.Context, input *elasticbeanstalk.UpdateApplicationVersionInput) *ElasticBeanstalkUpdateApplicationVersionFuture

	UpdateConfigurationTemplate(ctx workflow.Context, input *elasticbeanstalk.UpdateConfigurationTemplateInput) (*elasticbeanstalk.ConfigurationSettingsDescription, error)
	UpdateConfigurationTemplateAsync(ctx workflow.Context, input *elasticbeanstalk.UpdateConfigurationTemplateInput) *ElasticBeanstalkUpdateConfigurationTemplateFuture

	UpdateEnvironment(ctx workflow.Context, input *elasticbeanstalk.UpdateEnvironmentInput) (*elasticbeanstalk.EnvironmentDescription, error)
	UpdateEnvironmentAsync(ctx workflow.Context, input *elasticbeanstalk.UpdateEnvironmentInput) *ElasticBeanstalkUpdateEnvironmentFuture

	UpdateTagsForResource(ctx workflow.Context, input *elasticbeanstalk.UpdateTagsForResourceInput) (*elasticbeanstalk.UpdateTagsForResourceOutput, error)
	UpdateTagsForResourceAsync(ctx workflow.Context, input *elasticbeanstalk.UpdateTagsForResourceInput) *ElasticBeanstalkUpdateTagsForResourceFuture

	ValidateConfigurationSettings(ctx workflow.Context, input *elasticbeanstalk.ValidateConfigurationSettingsInput) (*elasticbeanstalk.ValidateConfigurationSettingsOutput, error)
	ValidateConfigurationSettingsAsync(ctx workflow.Context, input *elasticbeanstalk.ValidateConfigurationSettingsInput) *ElasticBeanstalkValidateConfigurationSettingsFuture

	WaitUntilEnvironmentExists(ctx workflow.Context, input *elasticbeanstalk.DescribeEnvironmentsInput) error
	WaitUntilEnvironmentExistsAsync(ctx workflow.Context, input *elasticbeanstalk.DescribeEnvironmentsInput) *clients.VoidFuture

	WaitUntilEnvironmentTerminated(ctx workflow.Context, input *elasticbeanstalk.DescribeEnvironmentsInput) error
	WaitUntilEnvironmentTerminatedAsync(ctx workflow.Context, input *elasticbeanstalk.DescribeEnvironmentsInput) *clients.VoidFuture

	WaitUntilEnvironmentUpdated(ctx workflow.Context, input *elasticbeanstalk.DescribeEnvironmentsInput) error
	WaitUntilEnvironmentUpdatedAsync(ctx workflow.Context, input *elasticbeanstalk.DescribeEnvironmentsInput) *clients.VoidFuture
}

func NewClient() Client {
	return &stub{}
}