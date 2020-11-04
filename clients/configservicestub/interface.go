// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.
// Copyright (c) 2020 Banzai Cloud All rights reserved.

package configservicestub

import (
	"github.com/aws/aws-sdk-go/service/configservice"
	"go.uber.org/cadence/workflow"

	"github.com/banzaicloud/cadence-aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	BatchGetAggregateResourceConfig(ctx workflow.Context, input *configservice.BatchGetAggregateResourceConfigInput) (*configservice.BatchGetAggregateResourceConfigOutput, error)
	BatchGetAggregateResourceConfigAsync(ctx workflow.Context, input *configservice.BatchGetAggregateResourceConfigInput) *ConfigServiceBatchGetAggregateResourceConfigFuture

	BatchGetResourceConfig(ctx workflow.Context, input *configservice.BatchGetResourceConfigInput) (*configservice.BatchGetResourceConfigOutput, error)
	BatchGetResourceConfigAsync(ctx workflow.Context, input *configservice.BatchGetResourceConfigInput) *ConfigServiceBatchGetResourceConfigFuture

	DeleteAggregationAuthorization(ctx workflow.Context, input *configservice.DeleteAggregationAuthorizationInput) (*configservice.DeleteAggregationAuthorizationOutput, error)
	DeleteAggregationAuthorizationAsync(ctx workflow.Context, input *configservice.DeleteAggregationAuthorizationInput) *ConfigServiceDeleteAggregationAuthorizationFuture

	DeleteConfigRule(ctx workflow.Context, input *configservice.DeleteConfigRuleInput) (*configservice.DeleteConfigRuleOutput, error)
	DeleteConfigRuleAsync(ctx workflow.Context, input *configservice.DeleteConfigRuleInput) *ConfigServiceDeleteConfigRuleFuture

	DeleteConfigurationAggregator(ctx workflow.Context, input *configservice.DeleteConfigurationAggregatorInput) (*configservice.DeleteConfigurationAggregatorOutput, error)
	DeleteConfigurationAggregatorAsync(ctx workflow.Context, input *configservice.DeleteConfigurationAggregatorInput) *ConfigServiceDeleteConfigurationAggregatorFuture

	DeleteConfigurationRecorder(ctx workflow.Context, input *configservice.DeleteConfigurationRecorderInput) (*configservice.DeleteConfigurationRecorderOutput, error)
	DeleteConfigurationRecorderAsync(ctx workflow.Context, input *configservice.DeleteConfigurationRecorderInput) *ConfigServiceDeleteConfigurationRecorderFuture

	DeleteConformancePack(ctx workflow.Context, input *configservice.DeleteConformancePackInput) (*configservice.DeleteConformancePackOutput, error)
	DeleteConformancePackAsync(ctx workflow.Context, input *configservice.DeleteConformancePackInput) *ConfigServiceDeleteConformancePackFuture

	DeleteDeliveryChannel(ctx workflow.Context, input *configservice.DeleteDeliveryChannelInput) (*configservice.DeleteDeliveryChannelOutput, error)
	DeleteDeliveryChannelAsync(ctx workflow.Context, input *configservice.DeleteDeliveryChannelInput) *ConfigServiceDeleteDeliveryChannelFuture

	DeleteEvaluationResults(ctx workflow.Context, input *configservice.DeleteEvaluationResultsInput) (*configservice.DeleteEvaluationResultsOutput, error)
	DeleteEvaluationResultsAsync(ctx workflow.Context, input *configservice.DeleteEvaluationResultsInput) *ConfigServiceDeleteEvaluationResultsFuture

	DeleteOrganizationConfigRule(ctx workflow.Context, input *configservice.DeleteOrganizationConfigRuleInput) (*configservice.DeleteOrganizationConfigRuleOutput, error)
	DeleteOrganizationConfigRuleAsync(ctx workflow.Context, input *configservice.DeleteOrganizationConfigRuleInput) *ConfigServiceDeleteOrganizationConfigRuleFuture

	DeleteOrganizationConformancePack(ctx workflow.Context, input *configservice.DeleteOrganizationConformancePackInput) (*configservice.DeleteOrganizationConformancePackOutput, error)
	DeleteOrganizationConformancePackAsync(ctx workflow.Context, input *configservice.DeleteOrganizationConformancePackInput) *ConfigServiceDeleteOrganizationConformancePackFuture

	DeleteRemediationConfiguration(ctx workflow.Context, input *configservice.DeleteRemediationConfigurationInput) (*configservice.DeleteRemediationConfigurationOutput, error)
	DeleteRemediationConfigurationAsync(ctx workflow.Context, input *configservice.DeleteRemediationConfigurationInput) *ConfigServiceDeleteRemediationConfigurationFuture

	DeleteRemediationExceptions(ctx workflow.Context, input *configservice.DeleteRemediationExceptionsInput) (*configservice.DeleteRemediationExceptionsOutput, error)
	DeleteRemediationExceptionsAsync(ctx workflow.Context, input *configservice.DeleteRemediationExceptionsInput) *ConfigServiceDeleteRemediationExceptionsFuture

	DeleteResourceConfig(ctx workflow.Context, input *configservice.DeleteResourceConfigInput) (*configservice.DeleteResourceConfigOutput, error)
	DeleteResourceConfigAsync(ctx workflow.Context, input *configservice.DeleteResourceConfigInput) *ConfigServiceDeleteResourceConfigFuture

	DeleteRetentionConfiguration(ctx workflow.Context, input *configservice.DeleteRetentionConfigurationInput) (*configservice.DeleteRetentionConfigurationOutput, error)
	DeleteRetentionConfigurationAsync(ctx workflow.Context, input *configservice.DeleteRetentionConfigurationInput) *ConfigServiceDeleteRetentionConfigurationFuture

	DeliverConfigSnapshot(ctx workflow.Context, input *configservice.DeliverConfigSnapshotInput) (*configservice.DeliverConfigSnapshotOutput, error)
	DeliverConfigSnapshotAsync(ctx workflow.Context, input *configservice.DeliverConfigSnapshotInput) *ConfigServiceDeliverConfigSnapshotFuture

	DescribeAggregateComplianceByConfigRules(ctx workflow.Context, input *configservice.DescribeAggregateComplianceByConfigRulesInput) (*configservice.DescribeAggregateComplianceByConfigRulesOutput, error)
	DescribeAggregateComplianceByConfigRulesAsync(ctx workflow.Context, input *configservice.DescribeAggregateComplianceByConfigRulesInput) *ConfigServiceDescribeAggregateComplianceByConfigRulesFuture

	DescribeAggregationAuthorizations(ctx workflow.Context, input *configservice.DescribeAggregationAuthorizationsInput) (*configservice.DescribeAggregationAuthorizationsOutput, error)
	DescribeAggregationAuthorizationsAsync(ctx workflow.Context, input *configservice.DescribeAggregationAuthorizationsInput) *ConfigServiceDescribeAggregationAuthorizationsFuture

	DescribeComplianceByConfigRule(ctx workflow.Context, input *configservice.DescribeComplianceByConfigRuleInput) (*configservice.DescribeComplianceByConfigRuleOutput, error)
	DescribeComplianceByConfigRuleAsync(ctx workflow.Context, input *configservice.DescribeComplianceByConfigRuleInput) *ConfigServiceDescribeComplianceByConfigRuleFuture

	DescribeComplianceByResource(ctx workflow.Context, input *configservice.DescribeComplianceByResourceInput) (*configservice.DescribeComplianceByResourceOutput, error)
	DescribeComplianceByResourceAsync(ctx workflow.Context, input *configservice.DescribeComplianceByResourceInput) *ConfigServiceDescribeComplianceByResourceFuture

	DescribeConfigRuleEvaluationStatus(ctx workflow.Context, input *configservice.DescribeConfigRuleEvaluationStatusInput) (*configservice.DescribeConfigRuleEvaluationStatusOutput, error)
	DescribeConfigRuleEvaluationStatusAsync(ctx workflow.Context, input *configservice.DescribeConfigRuleEvaluationStatusInput) *ConfigServiceDescribeConfigRuleEvaluationStatusFuture

	DescribeConfigRules(ctx workflow.Context, input *configservice.DescribeConfigRulesInput) (*configservice.DescribeConfigRulesOutput, error)
	DescribeConfigRulesAsync(ctx workflow.Context, input *configservice.DescribeConfigRulesInput) *ConfigServiceDescribeConfigRulesFuture

	DescribeConfigurationAggregatorSourcesStatus(ctx workflow.Context, input *configservice.DescribeConfigurationAggregatorSourcesStatusInput) (*configservice.DescribeConfigurationAggregatorSourcesStatusOutput, error)
	DescribeConfigurationAggregatorSourcesStatusAsync(ctx workflow.Context, input *configservice.DescribeConfigurationAggregatorSourcesStatusInput) *ConfigServiceDescribeConfigurationAggregatorSourcesStatusFuture

	DescribeConfigurationAggregators(ctx workflow.Context, input *configservice.DescribeConfigurationAggregatorsInput) (*configservice.DescribeConfigurationAggregatorsOutput, error)
	DescribeConfigurationAggregatorsAsync(ctx workflow.Context, input *configservice.DescribeConfigurationAggregatorsInput) *ConfigServiceDescribeConfigurationAggregatorsFuture

	DescribeConfigurationRecorderStatus(ctx workflow.Context, input *configservice.DescribeConfigurationRecorderStatusInput) (*configservice.DescribeConfigurationRecorderStatusOutput, error)
	DescribeConfigurationRecorderStatusAsync(ctx workflow.Context, input *configservice.DescribeConfigurationRecorderStatusInput) *ConfigServiceDescribeConfigurationRecorderStatusFuture

	DescribeConfigurationRecorders(ctx workflow.Context, input *configservice.DescribeConfigurationRecordersInput) (*configservice.DescribeConfigurationRecordersOutput, error)
	DescribeConfigurationRecordersAsync(ctx workflow.Context, input *configservice.DescribeConfigurationRecordersInput) *ConfigServiceDescribeConfigurationRecordersFuture

	DescribeConformancePackCompliance(ctx workflow.Context, input *configservice.DescribeConformancePackComplianceInput) (*configservice.DescribeConformancePackComplianceOutput, error)
	DescribeConformancePackComplianceAsync(ctx workflow.Context, input *configservice.DescribeConformancePackComplianceInput) *ConfigServiceDescribeConformancePackComplianceFuture

	DescribeConformancePackStatus(ctx workflow.Context, input *configservice.DescribeConformancePackStatusInput) (*configservice.DescribeConformancePackStatusOutput, error)
	DescribeConformancePackStatusAsync(ctx workflow.Context, input *configservice.DescribeConformancePackStatusInput) *ConfigServiceDescribeConformancePackStatusFuture

	DescribeConformancePacks(ctx workflow.Context, input *configservice.DescribeConformancePacksInput) (*configservice.DescribeConformancePacksOutput, error)
	DescribeConformancePacksAsync(ctx workflow.Context, input *configservice.DescribeConformancePacksInput) *ConfigServiceDescribeConformancePacksFuture

	DescribeDeliveryChannelStatus(ctx workflow.Context, input *configservice.DescribeDeliveryChannelStatusInput) (*configservice.DescribeDeliveryChannelStatusOutput, error)
	DescribeDeliveryChannelStatusAsync(ctx workflow.Context, input *configservice.DescribeDeliveryChannelStatusInput) *ConfigServiceDescribeDeliveryChannelStatusFuture

	DescribeDeliveryChannels(ctx workflow.Context, input *configservice.DescribeDeliveryChannelsInput) (*configservice.DescribeDeliveryChannelsOutput, error)
	DescribeDeliveryChannelsAsync(ctx workflow.Context, input *configservice.DescribeDeliveryChannelsInput) *ConfigServiceDescribeDeliveryChannelsFuture

	DescribeOrganizationConfigRuleStatuses(ctx workflow.Context, input *configservice.DescribeOrganizationConfigRuleStatusesInput) (*configservice.DescribeOrganizationConfigRuleStatusesOutput, error)
	DescribeOrganizationConfigRuleStatusesAsync(ctx workflow.Context, input *configservice.DescribeOrganizationConfigRuleStatusesInput) *ConfigServiceDescribeOrganizationConfigRuleStatusesFuture

	DescribeOrganizationConfigRules(ctx workflow.Context, input *configservice.DescribeOrganizationConfigRulesInput) (*configservice.DescribeOrganizationConfigRulesOutput, error)
	DescribeOrganizationConfigRulesAsync(ctx workflow.Context, input *configservice.DescribeOrganizationConfigRulesInput) *ConfigServiceDescribeOrganizationConfigRulesFuture

	DescribeOrganizationConformancePackStatuses(ctx workflow.Context, input *configservice.DescribeOrganizationConformancePackStatusesInput) (*configservice.DescribeOrganizationConformancePackStatusesOutput, error)
	DescribeOrganizationConformancePackStatusesAsync(ctx workflow.Context, input *configservice.DescribeOrganizationConformancePackStatusesInput) *ConfigServiceDescribeOrganizationConformancePackStatusesFuture

	DescribeOrganizationConformancePacks(ctx workflow.Context, input *configservice.DescribeOrganizationConformancePacksInput) (*configservice.DescribeOrganizationConformancePacksOutput, error)
	DescribeOrganizationConformancePacksAsync(ctx workflow.Context, input *configservice.DescribeOrganizationConformancePacksInput) *ConfigServiceDescribeOrganizationConformancePacksFuture

	DescribePendingAggregationRequests(ctx workflow.Context, input *configservice.DescribePendingAggregationRequestsInput) (*configservice.DescribePendingAggregationRequestsOutput, error)
	DescribePendingAggregationRequestsAsync(ctx workflow.Context, input *configservice.DescribePendingAggregationRequestsInput) *ConfigServiceDescribePendingAggregationRequestsFuture

	DescribeRemediationConfigurations(ctx workflow.Context, input *configservice.DescribeRemediationConfigurationsInput) (*configservice.DescribeRemediationConfigurationsOutput, error)
	DescribeRemediationConfigurationsAsync(ctx workflow.Context, input *configservice.DescribeRemediationConfigurationsInput) *ConfigServiceDescribeRemediationConfigurationsFuture

	DescribeRemediationExceptions(ctx workflow.Context, input *configservice.DescribeRemediationExceptionsInput) (*configservice.DescribeRemediationExceptionsOutput, error)
	DescribeRemediationExceptionsAsync(ctx workflow.Context, input *configservice.DescribeRemediationExceptionsInput) *ConfigServiceDescribeRemediationExceptionsFuture

	DescribeRemediationExecutionStatus(ctx workflow.Context, input *configservice.DescribeRemediationExecutionStatusInput) (*configservice.DescribeRemediationExecutionStatusOutput, error)
	DescribeRemediationExecutionStatusAsync(ctx workflow.Context, input *configservice.DescribeRemediationExecutionStatusInput) *ConfigServiceDescribeRemediationExecutionStatusFuture

	DescribeRetentionConfigurations(ctx workflow.Context, input *configservice.DescribeRetentionConfigurationsInput) (*configservice.DescribeRetentionConfigurationsOutput, error)
	DescribeRetentionConfigurationsAsync(ctx workflow.Context, input *configservice.DescribeRetentionConfigurationsInput) *ConfigServiceDescribeRetentionConfigurationsFuture

	GetAggregateComplianceDetailsByConfigRule(ctx workflow.Context, input *configservice.GetAggregateComplianceDetailsByConfigRuleInput) (*configservice.GetAggregateComplianceDetailsByConfigRuleOutput, error)
	GetAggregateComplianceDetailsByConfigRuleAsync(ctx workflow.Context, input *configservice.GetAggregateComplianceDetailsByConfigRuleInput) *ConfigServiceGetAggregateComplianceDetailsByConfigRuleFuture

	GetAggregateConfigRuleComplianceSummary(ctx workflow.Context, input *configservice.GetAggregateConfigRuleComplianceSummaryInput) (*configservice.GetAggregateConfigRuleComplianceSummaryOutput, error)
	GetAggregateConfigRuleComplianceSummaryAsync(ctx workflow.Context, input *configservice.GetAggregateConfigRuleComplianceSummaryInput) *ConfigServiceGetAggregateConfigRuleComplianceSummaryFuture

	GetAggregateDiscoveredResourceCounts(ctx workflow.Context, input *configservice.GetAggregateDiscoveredResourceCountsInput) (*configservice.GetAggregateDiscoveredResourceCountsOutput, error)
	GetAggregateDiscoveredResourceCountsAsync(ctx workflow.Context, input *configservice.GetAggregateDiscoveredResourceCountsInput) *ConfigServiceGetAggregateDiscoveredResourceCountsFuture

	GetAggregateResourceConfig(ctx workflow.Context, input *configservice.GetAggregateResourceConfigInput) (*configservice.GetAggregateResourceConfigOutput, error)
	GetAggregateResourceConfigAsync(ctx workflow.Context, input *configservice.GetAggregateResourceConfigInput) *ConfigServiceGetAggregateResourceConfigFuture

	GetComplianceDetailsByConfigRule(ctx workflow.Context, input *configservice.GetComplianceDetailsByConfigRuleInput) (*configservice.GetComplianceDetailsByConfigRuleOutput, error)
	GetComplianceDetailsByConfigRuleAsync(ctx workflow.Context, input *configservice.GetComplianceDetailsByConfigRuleInput) *ConfigServiceGetComplianceDetailsByConfigRuleFuture

	GetComplianceDetailsByResource(ctx workflow.Context, input *configservice.GetComplianceDetailsByResourceInput) (*configservice.GetComplianceDetailsByResourceOutput, error)
	GetComplianceDetailsByResourceAsync(ctx workflow.Context, input *configservice.GetComplianceDetailsByResourceInput) *ConfigServiceGetComplianceDetailsByResourceFuture

	GetComplianceSummaryByConfigRule(ctx workflow.Context, input *configservice.GetComplianceSummaryByConfigRuleInput) (*configservice.GetComplianceSummaryByConfigRuleOutput, error)
	GetComplianceSummaryByConfigRuleAsync(ctx workflow.Context, input *configservice.GetComplianceSummaryByConfigRuleInput) *ConfigServiceGetComplianceSummaryByConfigRuleFuture

	GetComplianceSummaryByResourceType(ctx workflow.Context, input *configservice.GetComplianceSummaryByResourceTypeInput) (*configservice.GetComplianceSummaryByResourceTypeOutput, error)
	GetComplianceSummaryByResourceTypeAsync(ctx workflow.Context, input *configservice.GetComplianceSummaryByResourceTypeInput) *ConfigServiceGetComplianceSummaryByResourceTypeFuture

	GetConformancePackComplianceDetails(ctx workflow.Context, input *configservice.GetConformancePackComplianceDetailsInput) (*configservice.GetConformancePackComplianceDetailsOutput, error)
	GetConformancePackComplianceDetailsAsync(ctx workflow.Context, input *configservice.GetConformancePackComplianceDetailsInput) *ConfigServiceGetConformancePackComplianceDetailsFuture

	GetConformancePackComplianceSummary(ctx workflow.Context, input *configservice.GetConformancePackComplianceSummaryInput) (*configservice.GetConformancePackComplianceSummaryOutput, error)
	GetConformancePackComplianceSummaryAsync(ctx workflow.Context, input *configservice.GetConformancePackComplianceSummaryInput) *ConfigServiceGetConformancePackComplianceSummaryFuture

	GetDiscoveredResourceCounts(ctx workflow.Context, input *configservice.GetDiscoveredResourceCountsInput) (*configservice.GetDiscoveredResourceCountsOutput, error)
	GetDiscoveredResourceCountsAsync(ctx workflow.Context, input *configservice.GetDiscoveredResourceCountsInput) *ConfigServiceGetDiscoveredResourceCountsFuture

	GetOrganizationConfigRuleDetailedStatus(ctx workflow.Context, input *configservice.GetOrganizationConfigRuleDetailedStatusInput) (*configservice.GetOrganizationConfigRuleDetailedStatusOutput, error)
	GetOrganizationConfigRuleDetailedStatusAsync(ctx workflow.Context, input *configservice.GetOrganizationConfigRuleDetailedStatusInput) *ConfigServiceGetOrganizationConfigRuleDetailedStatusFuture

	GetOrganizationConformancePackDetailedStatus(ctx workflow.Context, input *configservice.GetOrganizationConformancePackDetailedStatusInput) (*configservice.GetOrganizationConformancePackDetailedStatusOutput, error)
	GetOrganizationConformancePackDetailedStatusAsync(ctx workflow.Context, input *configservice.GetOrganizationConformancePackDetailedStatusInput) *ConfigServiceGetOrganizationConformancePackDetailedStatusFuture

	GetResourceConfigHistory(ctx workflow.Context, input *configservice.GetResourceConfigHistoryInput) (*configservice.GetResourceConfigHistoryOutput, error)
	GetResourceConfigHistoryAsync(ctx workflow.Context, input *configservice.GetResourceConfigHistoryInput) *ConfigServiceGetResourceConfigHistoryFuture

	ListAggregateDiscoveredResources(ctx workflow.Context, input *configservice.ListAggregateDiscoveredResourcesInput) (*configservice.ListAggregateDiscoveredResourcesOutput, error)
	ListAggregateDiscoveredResourcesAsync(ctx workflow.Context, input *configservice.ListAggregateDiscoveredResourcesInput) *ConfigServiceListAggregateDiscoveredResourcesFuture

	ListDiscoveredResources(ctx workflow.Context, input *configservice.ListDiscoveredResourcesInput) (*configservice.ListDiscoveredResourcesOutput, error)
	ListDiscoveredResourcesAsync(ctx workflow.Context, input *configservice.ListDiscoveredResourcesInput) *ConfigServiceListDiscoveredResourcesFuture

	ListTagsForResource(ctx workflow.Context, input *configservice.ListTagsForResourceInput) (*configservice.ListTagsForResourceOutput, error)
	ListTagsForResourceAsync(ctx workflow.Context, input *configservice.ListTagsForResourceInput) *ConfigServiceListTagsForResourceFuture

	PutAggregationAuthorization(ctx workflow.Context, input *configservice.PutAggregationAuthorizationInput) (*configservice.PutAggregationAuthorizationOutput, error)
	PutAggregationAuthorizationAsync(ctx workflow.Context, input *configservice.PutAggregationAuthorizationInput) *ConfigServicePutAggregationAuthorizationFuture

	PutConfigRule(ctx workflow.Context, input *configservice.PutConfigRuleInput) (*configservice.PutConfigRuleOutput, error)
	PutConfigRuleAsync(ctx workflow.Context, input *configservice.PutConfigRuleInput) *ConfigServicePutConfigRuleFuture

	PutConfigurationAggregator(ctx workflow.Context, input *configservice.PutConfigurationAggregatorInput) (*configservice.PutConfigurationAggregatorOutput, error)
	PutConfigurationAggregatorAsync(ctx workflow.Context, input *configservice.PutConfigurationAggregatorInput) *ConfigServicePutConfigurationAggregatorFuture

	PutConfigurationRecorder(ctx workflow.Context, input *configservice.PutConfigurationRecorderInput) (*configservice.PutConfigurationRecorderOutput, error)
	PutConfigurationRecorderAsync(ctx workflow.Context, input *configservice.PutConfigurationRecorderInput) *ConfigServicePutConfigurationRecorderFuture

	PutConformancePack(ctx workflow.Context, input *configservice.PutConformancePackInput) (*configservice.PutConformancePackOutput, error)
	PutConformancePackAsync(ctx workflow.Context, input *configservice.PutConformancePackInput) *ConfigServicePutConformancePackFuture

	PutDeliveryChannel(ctx workflow.Context, input *configservice.PutDeliveryChannelInput) (*configservice.PutDeliveryChannelOutput, error)
	PutDeliveryChannelAsync(ctx workflow.Context, input *configservice.PutDeliveryChannelInput) *ConfigServicePutDeliveryChannelFuture

	PutEvaluations(ctx workflow.Context, input *configservice.PutEvaluationsInput) (*configservice.PutEvaluationsOutput, error)
	PutEvaluationsAsync(ctx workflow.Context, input *configservice.PutEvaluationsInput) *ConfigServicePutEvaluationsFuture

	PutOrganizationConfigRule(ctx workflow.Context, input *configservice.PutOrganizationConfigRuleInput) (*configservice.PutOrganizationConfigRuleOutput, error)
	PutOrganizationConfigRuleAsync(ctx workflow.Context, input *configservice.PutOrganizationConfigRuleInput) *ConfigServicePutOrganizationConfigRuleFuture

	PutOrganizationConformancePack(ctx workflow.Context, input *configservice.PutOrganizationConformancePackInput) (*configservice.PutOrganizationConformancePackOutput, error)
	PutOrganizationConformancePackAsync(ctx workflow.Context, input *configservice.PutOrganizationConformancePackInput) *ConfigServicePutOrganizationConformancePackFuture

	PutRemediationConfigurations(ctx workflow.Context, input *configservice.PutRemediationConfigurationsInput) (*configservice.PutRemediationConfigurationsOutput, error)
	PutRemediationConfigurationsAsync(ctx workflow.Context, input *configservice.PutRemediationConfigurationsInput) *ConfigServicePutRemediationConfigurationsFuture

	PutRemediationExceptions(ctx workflow.Context, input *configservice.PutRemediationExceptionsInput) (*configservice.PutRemediationExceptionsOutput, error)
	PutRemediationExceptionsAsync(ctx workflow.Context, input *configservice.PutRemediationExceptionsInput) *ConfigServicePutRemediationExceptionsFuture

	PutResourceConfig(ctx workflow.Context, input *configservice.PutResourceConfigInput) (*configservice.PutResourceConfigOutput, error)
	PutResourceConfigAsync(ctx workflow.Context, input *configservice.PutResourceConfigInput) *ConfigServicePutResourceConfigFuture

	PutRetentionConfiguration(ctx workflow.Context, input *configservice.PutRetentionConfigurationInput) (*configservice.PutRetentionConfigurationOutput, error)
	PutRetentionConfigurationAsync(ctx workflow.Context, input *configservice.PutRetentionConfigurationInput) *ConfigServicePutRetentionConfigurationFuture

	SelectAggregateResourceConfig(ctx workflow.Context, input *configservice.SelectAggregateResourceConfigInput) (*configservice.SelectAggregateResourceConfigOutput, error)
	SelectAggregateResourceConfigAsync(ctx workflow.Context, input *configservice.SelectAggregateResourceConfigInput) *ConfigServiceSelectAggregateResourceConfigFuture

	SelectResourceConfig(ctx workflow.Context, input *configservice.SelectResourceConfigInput) (*configservice.SelectResourceConfigOutput, error)
	SelectResourceConfigAsync(ctx workflow.Context, input *configservice.SelectResourceConfigInput) *ConfigServiceSelectResourceConfigFuture

	StartConfigRulesEvaluation(ctx workflow.Context, input *configservice.StartConfigRulesEvaluationInput) (*configservice.StartConfigRulesEvaluationOutput, error)
	StartConfigRulesEvaluationAsync(ctx workflow.Context, input *configservice.StartConfigRulesEvaluationInput) *ConfigServiceStartConfigRulesEvaluationFuture

	StartConfigurationRecorder(ctx workflow.Context, input *configservice.StartConfigurationRecorderInput) (*configservice.StartConfigurationRecorderOutput, error)
	StartConfigurationRecorderAsync(ctx workflow.Context, input *configservice.StartConfigurationRecorderInput) *ConfigServiceStartConfigurationRecorderFuture

	StartRemediationExecution(ctx workflow.Context, input *configservice.StartRemediationExecutionInput) (*configservice.StartRemediationExecutionOutput, error)
	StartRemediationExecutionAsync(ctx workflow.Context, input *configservice.StartRemediationExecutionInput) *ConfigServiceStartRemediationExecutionFuture

	StopConfigurationRecorder(ctx workflow.Context, input *configservice.StopConfigurationRecorderInput) (*configservice.StopConfigurationRecorderOutput, error)
	StopConfigurationRecorderAsync(ctx workflow.Context, input *configservice.StopConfigurationRecorderInput) *ConfigServiceStopConfigurationRecorderFuture

	TagResource(ctx workflow.Context, input *configservice.TagResourceInput) (*configservice.TagResourceOutput, error)
	TagResourceAsync(ctx workflow.Context, input *configservice.TagResourceInput) *ConfigServiceTagResourceFuture

	UntagResource(ctx workflow.Context, input *configservice.UntagResourceInput) (*configservice.UntagResourceOutput, error)
	UntagResourceAsync(ctx workflow.Context, input *configservice.UntagResourceInput) *ConfigServiceUntagResourceFuture
}

func NewClient() Client {
	return &stub{}
}