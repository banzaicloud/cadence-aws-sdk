// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.
// Copyright (c) 2020 Banzai Cloud All rights reserved.

package inspectorstub

import (
	"github.com/aws/aws-sdk-go/service/inspector"
	"go.uber.org/cadence/workflow"

	"github.com/banzaicloud/cadence-aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type stub struct{}

type InspectorAddAttributesToFindingsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *InspectorAddAttributesToFindingsFuture) Get(ctx workflow.Context) (*inspector.AddAttributesToFindingsOutput, error) {
	var output inspector.AddAttributesToFindingsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type InspectorCreateAssessmentTargetFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *InspectorCreateAssessmentTargetFuture) Get(ctx workflow.Context) (*inspector.CreateAssessmentTargetOutput, error) {
	var output inspector.CreateAssessmentTargetOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type InspectorCreateAssessmentTemplateFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *InspectorCreateAssessmentTemplateFuture) Get(ctx workflow.Context) (*inspector.CreateAssessmentTemplateOutput, error) {
	var output inspector.CreateAssessmentTemplateOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type InspectorCreateExclusionsPreviewFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *InspectorCreateExclusionsPreviewFuture) Get(ctx workflow.Context) (*inspector.CreateExclusionsPreviewOutput, error) {
	var output inspector.CreateExclusionsPreviewOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type InspectorCreateResourceGroupFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *InspectorCreateResourceGroupFuture) Get(ctx workflow.Context) (*inspector.CreateResourceGroupOutput, error) {
	var output inspector.CreateResourceGroupOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type InspectorDeleteAssessmentRunFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *InspectorDeleteAssessmentRunFuture) Get(ctx workflow.Context) (*inspector.DeleteAssessmentRunOutput, error) {
	var output inspector.DeleteAssessmentRunOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type InspectorDeleteAssessmentTargetFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *InspectorDeleteAssessmentTargetFuture) Get(ctx workflow.Context) (*inspector.DeleteAssessmentTargetOutput, error) {
	var output inspector.DeleteAssessmentTargetOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type InspectorDeleteAssessmentTemplateFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *InspectorDeleteAssessmentTemplateFuture) Get(ctx workflow.Context) (*inspector.DeleteAssessmentTemplateOutput, error) {
	var output inspector.DeleteAssessmentTemplateOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type InspectorDescribeAssessmentRunsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *InspectorDescribeAssessmentRunsFuture) Get(ctx workflow.Context) (*inspector.DescribeAssessmentRunsOutput, error) {
	var output inspector.DescribeAssessmentRunsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type InspectorDescribeAssessmentTargetsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *InspectorDescribeAssessmentTargetsFuture) Get(ctx workflow.Context) (*inspector.DescribeAssessmentTargetsOutput, error) {
	var output inspector.DescribeAssessmentTargetsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type InspectorDescribeAssessmentTemplatesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *InspectorDescribeAssessmentTemplatesFuture) Get(ctx workflow.Context) (*inspector.DescribeAssessmentTemplatesOutput, error) {
	var output inspector.DescribeAssessmentTemplatesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type InspectorDescribeCrossAccountAccessRoleFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *InspectorDescribeCrossAccountAccessRoleFuture) Get(ctx workflow.Context) (*inspector.DescribeCrossAccountAccessRoleOutput, error) {
	var output inspector.DescribeCrossAccountAccessRoleOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type InspectorDescribeExclusionsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *InspectorDescribeExclusionsFuture) Get(ctx workflow.Context) (*inspector.DescribeExclusionsOutput, error) {
	var output inspector.DescribeExclusionsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type InspectorDescribeFindingsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *InspectorDescribeFindingsFuture) Get(ctx workflow.Context) (*inspector.DescribeFindingsOutput, error) {
	var output inspector.DescribeFindingsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type InspectorDescribeResourceGroupsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *InspectorDescribeResourceGroupsFuture) Get(ctx workflow.Context) (*inspector.DescribeResourceGroupsOutput, error) {
	var output inspector.DescribeResourceGroupsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type InspectorDescribeRulesPackagesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *InspectorDescribeRulesPackagesFuture) Get(ctx workflow.Context) (*inspector.DescribeRulesPackagesOutput, error) {
	var output inspector.DescribeRulesPackagesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type InspectorGetAssessmentReportFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *InspectorGetAssessmentReportFuture) Get(ctx workflow.Context) (*inspector.GetAssessmentReportOutput, error) {
	var output inspector.GetAssessmentReportOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type InspectorGetExclusionsPreviewFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *InspectorGetExclusionsPreviewFuture) Get(ctx workflow.Context) (*inspector.GetExclusionsPreviewOutput, error) {
	var output inspector.GetExclusionsPreviewOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type InspectorGetTelemetryMetadataFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *InspectorGetTelemetryMetadataFuture) Get(ctx workflow.Context) (*inspector.GetTelemetryMetadataOutput, error) {
	var output inspector.GetTelemetryMetadataOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type InspectorListAssessmentRunAgentsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *InspectorListAssessmentRunAgentsFuture) Get(ctx workflow.Context) (*inspector.ListAssessmentRunAgentsOutput, error) {
	var output inspector.ListAssessmentRunAgentsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type InspectorListAssessmentRunsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *InspectorListAssessmentRunsFuture) Get(ctx workflow.Context) (*inspector.ListAssessmentRunsOutput, error) {
	var output inspector.ListAssessmentRunsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type InspectorListAssessmentTargetsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *InspectorListAssessmentTargetsFuture) Get(ctx workflow.Context) (*inspector.ListAssessmentTargetsOutput, error) {
	var output inspector.ListAssessmentTargetsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type InspectorListAssessmentTemplatesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *InspectorListAssessmentTemplatesFuture) Get(ctx workflow.Context) (*inspector.ListAssessmentTemplatesOutput, error) {
	var output inspector.ListAssessmentTemplatesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type InspectorListEventSubscriptionsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *InspectorListEventSubscriptionsFuture) Get(ctx workflow.Context) (*inspector.ListEventSubscriptionsOutput, error) {
	var output inspector.ListEventSubscriptionsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type InspectorListExclusionsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *InspectorListExclusionsFuture) Get(ctx workflow.Context) (*inspector.ListExclusionsOutput, error) {
	var output inspector.ListExclusionsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type InspectorListFindingsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *InspectorListFindingsFuture) Get(ctx workflow.Context) (*inspector.ListFindingsOutput, error) {
	var output inspector.ListFindingsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type InspectorListRulesPackagesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *InspectorListRulesPackagesFuture) Get(ctx workflow.Context) (*inspector.ListRulesPackagesOutput, error) {
	var output inspector.ListRulesPackagesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type InspectorListTagsForResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *InspectorListTagsForResourceFuture) Get(ctx workflow.Context) (*inspector.ListTagsForResourceOutput, error) {
	var output inspector.ListTagsForResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type InspectorPreviewAgentsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *InspectorPreviewAgentsFuture) Get(ctx workflow.Context) (*inspector.PreviewAgentsOutput, error) {
	var output inspector.PreviewAgentsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type InspectorRegisterCrossAccountAccessRoleFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *InspectorRegisterCrossAccountAccessRoleFuture) Get(ctx workflow.Context) (*inspector.RegisterCrossAccountAccessRoleOutput, error) {
	var output inspector.RegisterCrossAccountAccessRoleOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type InspectorRemoveAttributesFromFindingsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *InspectorRemoveAttributesFromFindingsFuture) Get(ctx workflow.Context) (*inspector.RemoveAttributesFromFindingsOutput, error) {
	var output inspector.RemoveAttributesFromFindingsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type InspectorSetTagsForResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *InspectorSetTagsForResourceFuture) Get(ctx workflow.Context) (*inspector.SetTagsForResourceOutput, error) {
	var output inspector.SetTagsForResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type InspectorStartAssessmentRunFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *InspectorStartAssessmentRunFuture) Get(ctx workflow.Context) (*inspector.StartAssessmentRunOutput, error) {
	var output inspector.StartAssessmentRunOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type InspectorStopAssessmentRunFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *InspectorStopAssessmentRunFuture) Get(ctx workflow.Context) (*inspector.StopAssessmentRunOutput, error) {
	var output inspector.StopAssessmentRunOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type InspectorSubscribeToEventFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *InspectorSubscribeToEventFuture) Get(ctx workflow.Context) (*inspector.SubscribeToEventOutput, error) {
	var output inspector.SubscribeToEventOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type InspectorUnsubscribeFromEventFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *InspectorUnsubscribeFromEventFuture) Get(ctx workflow.Context) (*inspector.UnsubscribeFromEventOutput, error) {
	var output inspector.UnsubscribeFromEventOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type InspectorUpdateAssessmentTargetFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *InspectorUpdateAssessmentTargetFuture) Get(ctx workflow.Context) (*inspector.UpdateAssessmentTargetOutput, error) {
	var output inspector.UpdateAssessmentTargetOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *stub) AddAttributesToFindings(ctx workflow.Context, input *inspector.AddAttributesToFindingsInput) (*inspector.AddAttributesToFindingsOutput, error) {
	var output inspector.AddAttributesToFindingsOutput
	err := workflow.ExecuteActivity(ctx, "aws-inspector-AddAttributesToFindings", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) AddAttributesToFindingsAsync(ctx workflow.Context, input *inspector.AddAttributesToFindingsInput) *InspectorAddAttributesToFindingsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-inspector-AddAttributesToFindings", input)
	return &InspectorAddAttributesToFindingsFuture{Future: future}
}

func (a *stub) CreateAssessmentTarget(ctx workflow.Context, input *inspector.CreateAssessmentTargetInput) (*inspector.CreateAssessmentTargetOutput, error) {
	var output inspector.CreateAssessmentTargetOutput
	err := workflow.ExecuteActivity(ctx, "aws-inspector-CreateAssessmentTarget", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateAssessmentTargetAsync(ctx workflow.Context, input *inspector.CreateAssessmentTargetInput) *InspectorCreateAssessmentTargetFuture {
	future := workflow.ExecuteActivity(ctx, "aws-inspector-CreateAssessmentTarget", input)
	return &InspectorCreateAssessmentTargetFuture{Future: future}
}

func (a *stub) CreateAssessmentTemplate(ctx workflow.Context, input *inspector.CreateAssessmentTemplateInput) (*inspector.CreateAssessmentTemplateOutput, error) {
	var output inspector.CreateAssessmentTemplateOutput
	err := workflow.ExecuteActivity(ctx, "aws-inspector-CreateAssessmentTemplate", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateAssessmentTemplateAsync(ctx workflow.Context, input *inspector.CreateAssessmentTemplateInput) *InspectorCreateAssessmentTemplateFuture {
	future := workflow.ExecuteActivity(ctx, "aws-inspector-CreateAssessmentTemplate", input)
	return &InspectorCreateAssessmentTemplateFuture{Future: future}
}

func (a *stub) CreateExclusionsPreview(ctx workflow.Context, input *inspector.CreateExclusionsPreviewInput) (*inspector.CreateExclusionsPreviewOutput, error) {
	var output inspector.CreateExclusionsPreviewOutput
	err := workflow.ExecuteActivity(ctx, "aws-inspector-CreateExclusionsPreview", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateExclusionsPreviewAsync(ctx workflow.Context, input *inspector.CreateExclusionsPreviewInput) *InspectorCreateExclusionsPreviewFuture {
	future := workflow.ExecuteActivity(ctx, "aws-inspector-CreateExclusionsPreview", input)
	return &InspectorCreateExclusionsPreviewFuture{Future: future}
}

func (a *stub) CreateResourceGroup(ctx workflow.Context, input *inspector.CreateResourceGroupInput) (*inspector.CreateResourceGroupOutput, error) {
	var output inspector.CreateResourceGroupOutput
	err := workflow.ExecuteActivity(ctx, "aws-inspector-CreateResourceGroup", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateResourceGroupAsync(ctx workflow.Context, input *inspector.CreateResourceGroupInput) *InspectorCreateResourceGroupFuture {
	future := workflow.ExecuteActivity(ctx, "aws-inspector-CreateResourceGroup", input)
	return &InspectorCreateResourceGroupFuture{Future: future}
}

func (a *stub) DeleteAssessmentRun(ctx workflow.Context, input *inspector.DeleteAssessmentRunInput) (*inspector.DeleteAssessmentRunOutput, error) {
	var output inspector.DeleteAssessmentRunOutput
	err := workflow.ExecuteActivity(ctx, "aws-inspector-DeleteAssessmentRun", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteAssessmentRunAsync(ctx workflow.Context, input *inspector.DeleteAssessmentRunInput) *InspectorDeleteAssessmentRunFuture {
	future := workflow.ExecuteActivity(ctx, "aws-inspector-DeleteAssessmentRun", input)
	return &InspectorDeleteAssessmentRunFuture{Future: future}
}

func (a *stub) DeleteAssessmentTarget(ctx workflow.Context, input *inspector.DeleteAssessmentTargetInput) (*inspector.DeleteAssessmentTargetOutput, error) {
	var output inspector.DeleteAssessmentTargetOutput
	err := workflow.ExecuteActivity(ctx, "aws-inspector-DeleteAssessmentTarget", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteAssessmentTargetAsync(ctx workflow.Context, input *inspector.DeleteAssessmentTargetInput) *InspectorDeleteAssessmentTargetFuture {
	future := workflow.ExecuteActivity(ctx, "aws-inspector-DeleteAssessmentTarget", input)
	return &InspectorDeleteAssessmentTargetFuture{Future: future}
}

func (a *stub) DeleteAssessmentTemplate(ctx workflow.Context, input *inspector.DeleteAssessmentTemplateInput) (*inspector.DeleteAssessmentTemplateOutput, error) {
	var output inspector.DeleteAssessmentTemplateOutput
	err := workflow.ExecuteActivity(ctx, "aws-inspector-DeleteAssessmentTemplate", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteAssessmentTemplateAsync(ctx workflow.Context, input *inspector.DeleteAssessmentTemplateInput) *InspectorDeleteAssessmentTemplateFuture {
	future := workflow.ExecuteActivity(ctx, "aws-inspector-DeleteAssessmentTemplate", input)
	return &InspectorDeleteAssessmentTemplateFuture{Future: future}
}

func (a *stub) DescribeAssessmentRuns(ctx workflow.Context, input *inspector.DescribeAssessmentRunsInput) (*inspector.DescribeAssessmentRunsOutput, error) {
	var output inspector.DescribeAssessmentRunsOutput
	err := workflow.ExecuteActivity(ctx, "aws-inspector-DescribeAssessmentRuns", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeAssessmentRunsAsync(ctx workflow.Context, input *inspector.DescribeAssessmentRunsInput) *InspectorDescribeAssessmentRunsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-inspector-DescribeAssessmentRuns", input)
	return &InspectorDescribeAssessmentRunsFuture{Future: future}
}

func (a *stub) DescribeAssessmentTargets(ctx workflow.Context, input *inspector.DescribeAssessmentTargetsInput) (*inspector.DescribeAssessmentTargetsOutput, error) {
	var output inspector.DescribeAssessmentTargetsOutput
	err := workflow.ExecuteActivity(ctx, "aws-inspector-DescribeAssessmentTargets", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeAssessmentTargetsAsync(ctx workflow.Context, input *inspector.DescribeAssessmentTargetsInput) *InspectorDescribeAssessmentTargetsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-inspector-DescribeAssessmentTargets", input)
	return &InspectorDescribeAssessmentTargetsFuture{Future: future}
}

func (a *stub) DescribeAssessmentTemplates(ctx workflow.Context, input *inspector.DescribeAssessmentTemplatesInput) (*inspector.DescribeAssessmentTemplatesOutput, error) {
	var output inspector.DescribeAssessmentTemplatesOutput
	err := workflow.ExecuteActivity(ctx, "aws-inspector-DescribeAssessmentTemplates", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeAssessmentTemplatesAsync(ctx workflow.Context, input *inspector.DescribeAssessmentTemplatesInput) *InspectorDescribeAssessmentTemplatesFuture {
	future := workflow.ExecuteActivity(ctx, "aws-inspector-DescribeAssessmentTemplates", input)
	return &InspectorDescribeAssessmentTemplatesFuture{Future: future}
}

func (a *stub) DescribeCrossAccountAccessRole(ctx workflow.Context, input *inspector.DescribeCrossAccountAccessRoleInput) (*inspector.DescribeCrossAccountAccessRoleOutput, error) {
	var output inspector.DescribeCrossAccountAccessRoleOutput
	err := workflow.ExecuteActivity(ctx, "aws-inspector-DescribeCrossAccountAccessRole", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeCrossAccountAccessRoleAsync(ctx workflow.Context, input *inspector.DescribeCrossAccountAccessRoleInput) *InspectorDescribeCrossAccountAccessRoleFuture {
	future := workflow.ExecuteActivity(ctx, "aws-inspector-DescribeCrossAccountAccessRole", input)
	return &InspectorDescribeCrossAccountAccessRoleFuture{Future: future}
}

func (a *stub) DescribeExclusions(ctx workflow.Context, input *inspector.DescribeExclusionsInput) (*inspector.DescribeExclusionsOutput, error) {
	var output inspector.DescribeExclusionsOutput
	err := workflow.ExecuteActivity(ctx, "aws-inspector-DescribeExclusions", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeExclusionsAsync(ctx workflow.Context, input *inspector.DescribeExclusionsInput) *InspectorDescribeExclusionsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-inspector-DescribeExclusions", input)
	return &InspectorDescribeExclusionsFuture{Future: future}
}

func (a *stub) DescribeFindings(ctx workflow.Context, input *inspector.DescribeFindingsInput) (*inspector.DescribeFindingsOutput, error) {
	var output inspector.DescribeFindingsOutput
	err := workflow.ExecuteActivity(ctx, "aws-inspector-DescribeFindings", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeFindingsAsync(ctx workflow.Context, input *inspector.DescribeFindingsInput) *InspectorDescribeFindingsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-inspector-DescribeFindings", input)
	return &InspectorDescribeFindingsFuture{Future: future}
}

func (a *stub) DescribeResourceGroups(ctx workflow.Context, input *inspector.DescribeResourceGroupsInput) (*inspector.DescribeResourceGroupsOutput, error) {
	var output inspector.DescribeResourceGroupsOutput
	err := workflow.ExecuteActivity(ctx, "aws-inspector-DescribeResourceGroups", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeResourceGroupsAsync(ctx workflow.Context, input *inspector.DescribeResourceGroupsInput) *InspectorDescribeResourceGroupsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-inspector-DescribeResourceGroups", input)
	return &InspectorDescribeResourceGroupsFuture{Future: future}
}

func (a *stub) DescribeRulesPackages(ctx workflow.Context, input *inspector.DescribeRulesPackagesInput) (*inspector.DescribeRulesPackagesOutput, error) {
	var output inspector.DescribeRulesPackagesOutput
	err := workflow.ExecuteActivity(ctx, "aws-inspector-DescribeRulesPackages", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeRulesPackagesAsync(ctx workflow.Context, input *inspector.DescribeRulesPackagesInput) *InspectorDescribeRulesPackagesFuture {
	future := workflow.ExecuteActivity(ctx, "aws-inspector-DescribeRulesPackages", input)
	return &InspectorDescribeRulesPackagesFuture{Future: future}
}

func (a *stub) GetAssessmentReport(ctx workflow.Context, input *inspector.GetAssessmentReportInput) (*inspector.GetAssessmentReportOutput, error) {
	var output inspector.GetAssessmentReportOutput
	err := workflow.ExecuteActivity(ctx, "aws-inspector-GetAssessmentReport", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetAssessmentReportAsync(ctx workflow.Context, input *inspector.GetAssessmentReportInput) *InspectorGetAssessmentReportFuture {
	future := workflow.ExecuteActivity(ctx, "aws-inspector-GetAssessmentReport", input)
	return &InspectorGetAssessmentReportFuture{Future: future}
}

func (a *stub) GetExclusionsPreview(ctx workflow.Context, input *inspector.GetExclusionsPreviewInput) (*inspector.GetExclusionsPreviewOutput, error) {
	var output inspector.GetExclusionsPreviewOutput
	err := workflow.ExecuteActivity(ctx, "aws-inspector-GetExclusionsPreview", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetExclusionsPreviewAsync(ctx workflow.Context, input *inspector.GetExclusionsPreviewInput) *InspectorGetExclusionsPreviewFuture {
	future := workflow.ExecuteActivity(ctx, "aws-inspector-GetExclusionsPreview", input)
	return &InspectorGetExclusionsPreviewFuture{Future: future}
}

func (a *stub) GetTelemetryMetadata(ctx workflow.Context, input *inspector.GetTelemetryMetadataInput) (*inspector.GetTelemetryMetadataOutput, error) {
	var output inspector.GetTelemetryMetadataOutput
	err := workflow.ExecuteActivity(ctx, "aws-inspector-GetTelemetryMetadata", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetTelemetryMetadataAsync(ctx workflow.Context, input *inspector.GetTelemetryMetadataInput) *InspectorGetTelemetryMetadataFuture {
	future := workflow.ExecuteActivity(ctx, "aws-inspector-GetTelemetryMetadata", input)
	return &InspectorGetTelemetryMetadataFuture{Future: future}
}

func (a *stub) ListAssessmentRunAgents(ctx workflow.Context, input *inspector.ListAssessmentRunAgentsInput) (*inspector.ListAssessmentRunAgentsOutput, error) {
	var output inspector.ListAssessmentRunAgentsOutput
	err := workflow.ExecuteActivity(ctx, "aws-inspector-ListAssessmentRunAgents", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListAssessmentRunAgentsAsync(ctx workflow.Context, input *inspector.ListAssessmentRunAgentsInput) *InspectorListAssessmentRunAgentsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-inspector-ListAssessmentRunAgents", input)
	return &InspectorListAssessmentRunAgentsFuture{Future: future}
}

func (a *stub) ListAssessmentRuns(ctx workflow.Context, input *inspector.ListAssessmentRunsInput) (*inspector.ListAssessmentRunsOutput, error) {
	var output inspector.ListAssessmentRunsOutput
	err := workflow.ExecuteActivity(ctx, "aws-inspector-ListAssessmentRuns", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListAssessmentRunsAsync(ctx workflow.Context, input *inspector.ListAssessmentRunsInput) *InspectorListAssessmentRunsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-inspector-ListAssessmentRuns", input)
	return &InspectorListAssessmentRunsFuture{Future: future}
}

func (a *stub) ListAssessmentTargets(ctx workflow.Context, input *inspector.ListAssessmentTargetsInput) (*inspector.ListAssessmentTargetsOutput, error) {
	var output inspector.ListAssessmentTargetsOutput
	err := workflow.ExecuteActivity(ctx, "aws-inspector-ListAssessmentTargets", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListAssessmentTargetsAsync(ctx workflow.Context, input *inspector.ListAssessmentTargetsInput) *InspectorListAssessmentTargetsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-inspector-ListAssessmentTargets", input)
	return &InspectorListAssessmentTargetsFuture{Future: future}
}

func (a *stub) ListAssessmentTemplates(ctx workflow.Context, input *inspector.ListAssessmentTemplatesInput) (*inspector.ListAssessmentTemplatesOutput, error) {
	var output inspector.ListAssessmentTemplatesOutput
	err := workflow.ExecuteActivity(ctx, "aws-inspector-ListAssessmentTemplates", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListAssessmentTemplatesAsync(ctx workflow.Context, input *inspector.ListAssessmentTemplatesInput) *InspectorListAssessmentTemplatesFuture {
	future := workflow.ExecuteActivity(ctx, "aws-inspector-ListAssessmentTemplates", input)
	return &InspectorListAssessmentTemplatesFuture{Future: future}
}

func (a *stub) ListEventSubscriptions(ctx workflow.Context, input *inspector.ListEventSubscriptionsInput) (*inspector.ListEventSubscriptionsOutput, error) {
	var output inspector.ListEventSubscriptionsOutput
	err := workflow.ExecuteActivity(ctx, "aws-inspector-ListEventSubscriptions", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListEventSubscriptionsAsync(ctx workflow.Context, input *inspector.ListEventSubscriptionsInput) *InspectorListEventSubscriptionsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-inspector-ListEventSubscriptions", input)
	return &InspectorListEventSubscriptionsFuture{Future: future}
}

func (a *stub) ListExclusions(ctx workflow.Context, input *inspector.ListExclusionsInput) (*inspector.ListExclusionsOutput, error) {
	var output inspector.ListExclusionsOutput
	err := workflow.ExecuteActivity(ctx, "aws-inspector-ListExclusions", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListExclusionsAsync(ctx workflow.Context, input *inspector.ListExclusionsInput) *InspectorListExclusionsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-inspector-ListExclusions", input)
	return &InspectorListExclusionsFuture{Future: future}
}

func (a *stub) ListFindings(ctx workflow.Context, input *inspector.ListFindingsInput) (*inspector.ListFindingsOutput, error) {
	var output inspector.ListFindingsOutput
	err := workflow.ExecuteActivity(ctx, "aws-inspector-ListFindings", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListFindingsAsync(ctx workflow.Context, input *inspector.ListFindingsInput) *InspectorListFindingsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-inspector-ListFindings", input)
	return &InspectorListFindingsFuture{Future: future}
}

func (a *stub) ListRulesPackages(ctx workflow.Context, input *inspector.ListRulesPackagesInput) (*inspector.ListRulesPackagesOutput, error) {
	var output inspector.ListRulesPackagesOutput
	err := workflow.ExecuteActivity(ctx, "aws-inspector-ListRulesPackages", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListRulesPackagesAsync(ctx workflow.Context, input *inspector.ListRulesPackagesInput) *InspectorListRulesPackagesFuture {
	future := workflow.ExecuteActivity(ctx, "aws-inspector-ListRulesPackages", input)
	return &InspectorListRulesPackagesFuture{Future: future}
}

func (a *stub) ListTagsForResource(ctx workflow.Context, input *inspector.ListTagsForResourceInput) (*inspector.ListTagsForResourceOutput, error) {
	var output inspector.ListTagsForResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws-inspector-ListTagsForResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListTagsForResourceAsync(ctx workflow.Context, input *inspector.ListTagsForResourceInput) *InspectorListTagsForResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws-inspector-ListTagsForResource", input)
	return &InspectorListTagsForResourceFuture{Future: future}
}

func (a *stub) PreviewAgents(ctx workflow.Context, input *inspector.PreviewAgentsInput) (*inspector.PreviewAgentsOutput, error) {
	var output inspector.PreviewAgentsOutput
	err := workflow.ExecuteActivity(ctx, "aws-inspector-PreviewAgents", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) PreviewAgentsAsync(ctx workflow.Context, input *inspector.PreviewAgentsInput) *InspectorPreviewAgentsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-inspector-PreviewAgents", input)
	return &InspectorPreviewAgentsFuture{Future: future}
}

func (a *stub) RegisterCrossAccountAccessRole(ctx workflow.Context, input *inspector.RegisterCrossAccountAccessRoleInput) (*inspector.RegisterCrossAccountAccessRoleOutput, error) {
	var output inspector.RegisterCrossAccountAccessRoleOutput
	err := workflow.ExecuteActivity(ctx, "aws-inspector-RegisterCrossAccountAccessRole", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) RegisterCrossAccountAccessRoleAsync(ctx workflow.Context, input *inspector.RegisterCrossAccountAccessRoleInput) *InspectorRegisterCrossAccountAccessRoleFuture {
	future := workflow.ExecuteActivity(ctx, "aws-inspector-RegisterCrossAccountAccessRole", input)
	return &InspectorRegisterCrossAccountAccessRoleFuture{Future: future}
}

func (a *stub) RemoveAttributesFromFindings(ctx workflow.Context, input *inspector.RemoveAttributesFromFindingsInput) (*inspector.RemoveAttributesFromFindingsOutput, error) {
	var output inspector.RemoveAttributesFromFindingsOutput
	err := workflow.ExecuteActivity(ctx, "aws-inspector-RemoveAttributesFromFindings", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) RemoveAttributesFromFindingsAsync(ctx workflow.Context, input *inspector.RemoveAttributesFromFindingsInput) *InspectorRemoveAttributesFromFindingsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-inspector-RemoveAttributesFromFindings", input)
	return &InspectorRemoveAttributesFromFindingsFuture{Future: future}
}

func (a *stub) SetTagsForResource(ctx workflow.Context, input *inspector.SetTagsForResourceInput) (*inspector.SetTagsForResourceOutput, error) {
	var output inspector.SetTagsForResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws-inspector-SetTagsForResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) SetTagsForResourceAsync(ctx workflow.Context, input *inspector.SetTagsForResourceInput) *InspectorSetTagsForResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws-inspector-SetTagsForResource", input)
	return &InspectorSetTagsForResourceFuture{Future: future}
}

func (a *stub) StartAssessmentRun(ctx workflow.Context, input *inspector.StartAssessmentRunInput) (*inspector.StartAssessmentRunOutput, error) {
	var output inspector.StartAssessmentRunOutput
	err := workflow.ExecuteActivity(ctx, "aws-inspector-StartAssessmentRun", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) StartAssessmentRunAsync(ctx workflow.Context, input *inspector.StartAssessmentRunInput) *InspectorStartAssessmentRunFuture {
	future := workflow.ExecuteActivity(ctx, "aws-inspector-StartAssessmentRun", input)
	return &InspectorStartAssessmentRunFuture{Future: future}
}

func (a *stub) StopAssessmentRun(ctx workflow.Context, input *inspector.StopAssessmentRunInput) (*inspector.StopAssessmentRunOutput, error) {
	var output inspector.StopAssessmentRunOutput
	err := workflow.ExecuteActivity(ctx, "aws-inspector-StopAssessmentRun", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) StopAssessmentRunAsync(ctx workflow.Context, input *inspector.StopAssessmentRunInput) *InspectorStopAssessmentRunFuture {
	future := workflow.ExecuteActivity(ctx, "aws-inspector-StopAssessmentRun", input)
	return &InspectorStopAssessmentRunFuture{Future: future}
}

func (a *stub) SubscribeToEvent(ctx workflow.Context, input *inspector.SubscribeToEventInput) (*inspector.SubscribeToEventOutput, error) {
	var output inspector.SubscribeToEventOutput
	err := workflow.ExecuteActivity(ctx, "aws-inspector-SubscribeToEvent", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) SubscribeToEventAsync(ctx workflow.Context, input *inspector.SubscribeToEventInput) *InspectorSubscribeToEventFuture {
	future := workflow.ExecuteActivity(ctx, "aws-inspector-SubscribeToEvent", input)
	return &InspectorSubscribeToEventFuture{Future: future}
}

func (a *stub) UnsubscribeFromEvent(ctx workflow.Context, input *inspector.UnsubscribeFromEventInput) (*inspector.UnsubscribeFromEventOutput, error) {
	var output inspector.UnsubscribeFromEventOutput
	err := workflow.ExecuteActivity(ctx, "aws-inspector-UnsubscribeFromEvent", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UnsubscribeFromEventAsync(ctx workflow.Context, input *inspector.UnsubscribeFromEventInput) *InspectorUnsubscribeFromEventFuture {
	future := workflow.ExecuteActivity(ctx, "aws-inspector-UnsubscribeFromEvent", input)
	return &InspectorUnsubscribeFromEventFuture{Future: future}
}

func (a *stub) UpdateAssessmentTarget(ctx workflow.Context, input *inspector.UpdateAssessmentTargetInput) (*inspector.UpdateAssessmentTargetOutput, error) {
	var output inspector.UpdateAssessmentTargetOutput
	err := workflow.ExecuteActivity(ctx, "aws-inspector-UpdateAssessmentTarget", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateAssessmentTargetAsync(ctx workflow.Context, input *inspector.UpdateAssessmentTargetInput) *InspectorUpdateAssessmentTargetFuture {
	future := workflow.ExecuteActivity(ctx, "aws-inspector-UpdateAssessmentTarget", input)
	return &InspectorUpdateAssessmentTargetFuture{Future: future}
}
