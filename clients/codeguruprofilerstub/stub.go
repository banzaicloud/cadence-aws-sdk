// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7

package codeguruprofilerstub

import (
	"github.com/aws/aws-sdk-go/service/codeguruprofiler"
	"go.uber.org/cadence/workflow"

	"github.com/banzaicloud/cadence-aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type stub struct{}

type CodeGuruProfilerAddNotificationChannelsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CodeGuruProfilerAddNotificationChannelsFuture) Get(ctx workflow.Context) (*codeguruprofiler.AddNotificationChannelsOutput, error) {
	var output codeguruprofiler.AddNotificationChannelsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CodeGuruProfilerBatchGetFrameMetricDataFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CodeGuruProfilerBatchGetFrameMetricDataFuture) Get(ctx workflow.Context) (*codeguruprofiler.BatchGetFrameMetricDataOutput, error) {
	var output codeguruprofiler.BatchGetFrameMetricDataOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CodeGuruProfilerConfigureAgentFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CodeGuruProfilerConfigureAgentFuture) Get(ctx workflow.Context) (*codeguruprofiler.ConfigureAgentOutput, error) {
	var output codeguruprofiler.ConfigureAgentOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CodeGuruProfilerCreateProfilingGroupFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CodeGuruProfilerCreateProfilingGroupFuture) Get(ctx workflow.Context) (*codeguruprofiler.CreateProfilingGroupOutput, error) {
	var output codeguruprofiler.CreateProfilingGroupOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CodeGuruProfilerDeleteProfilingGroupFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CodeGuruProfilerDeleteProfilingGroupFuture) Get(ctx workflow.Context) (*codeguruprofiler.DeleteProfilingGroupOutput, error) {
	var output codeguruprofiler.DeleteProfilingGroupOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CodeGuruProfilerDescribeProfilingGroupFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CodeGuruProfilerDescribeProfilingGroupFuture) Get(ctx workflow.Context) (*codeguruprofiler.DescribeProfilingGroupOutput, error) {
	var output codeguruprofiler.DescribeProfilingGroupOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CodeGuruProfilerGetFindingsReportAccountSummaryFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CodeGuruProfilerGetFindingsReportAccountSummaryFuture) Get(ctx workflow.Context) (*codeguruprofiler.GetFindingsReportAccountSummaryOutput, error) {
	var output codeguruprofiler.GetFindingsReportAccountSummaryOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CodeGuruProfilerGetNotificationConfigurationFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CodeGuruProfilerGetNotificationConfigurationFuture) Get(ctx workflow.Context) (*codeguruprofiler.GetNotificationConfigurationOutput, error) {
	var output codeguruprofiler.GetNotificationConfigurationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CodeGuruProfilerGetPolicyFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CodeGuruProfilerGetPolicyFuture) Get(ctx workflow.Context) (*codeguruprofiler.GetPolicyOutput, error) {
	var output codeguruprofiler.GetPolicyOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CodeGuruProfilerGetProfileFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CodeGuruProfilerGetProfileFuture) Get(ctx workflow.Context) (*codeguruprofiler.GetProfileOutput, error) {
	var output codeguruprofiler.GetProfileOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CodeGuruProfilerGetRecommendationsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CodeGuruProfilerGetRecommendationsFuture) Get(ctx workflow.Context) (*codeguruprofiler.GetRecommendationsOutput, error) {
	var output codeguruprofiler.GetRecommendationsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CodeGuruProfilerListFindingsReportsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CodeGuruProfilerListFindingsReportsFuture) Get(ctx workflow.Context) (*codeguruprofiler.ListFindingsReportsOutput, error) {
	var output codeguruprofiler.ListFindingsReportsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CodeGuruProfilerListProfileTimesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CodeGuruProfilerListProfileTimesFuture) Get(ctx workflow.Context) (*codeguruprofiler.ListProfileTimesOutput, error) {
	var output codeguruprofiler.ListProfileTimesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CodeGuruProfilerListProfilingGroupsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CodeGuruProfilerListProfilingGroupsFuture) Get(ctx workflow.Context) (*codeguruprofiler.ListProfilingGroupsOutput, error) {
	var output codeguruprofiler.ListProfilingGroupsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CodeGuruProfilerListTagsForResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CodeGuruProfilerListTagsForResourceFuture) Get(ctx workflow.Context) (*codeguruprofiler.ListTagsForResourceOutput, error) {
	var output codeguruprofiler.ListTagsForResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CodeGuruProfilerPostAgentProfileFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CodeGuruProfilerPostAgentProfileFuture) Get(ctx workflow.Context) (*codeguruprofiler.PostAgentProfileOutput, error) {
	var output codeguruprofiler.PostAgentProfileOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CodeGuruProfilerPutPermissionFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CodeGuruProfilerPutPermissionFuture) Get(ctx workflow.Context) (*codeguruprofiler.PutPermissionOutput, error) {
	var output codeguruprofiler.PutPermissionOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CodeGuruProfilerRemoveNotificationChannelFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CodeGuruProfilerRemoveNotificationChannelFuture) Get(ctx workflow.Context) (*codeguruprofiler.RemoveNotificationChannelOutput, error) {
	var output codeguruprofiler.RemoveNotificationChannelOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CodeGuruProfilerRemovePermissionFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CodeGuruProfilerRemovePermissionFuture) Get(ctx workflow.Context) (*codeguruprofiler.RemovePermissionOutput, error) {
	var output codeguruprofiler.RemovePermissionOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CodeGuruProfilerSubmitFeedbackFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CodeGuruProfilerSubmitFeedbackFuture) Get(ctx workflow.Context) (*codeguruprofiler.SubmitFeedbackOutput, error) {
	var output codeguruprofiler.SubmitFeedbackOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CodeGuruProfilerTagResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CodeGuruProfilerTagResourceFuture) Get(ctx workflow.Context) (*codeguruprofiler.TagResourceOutput, error) {
	var output codeguruprofiler.TagResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CodeGuruProfilerUntagResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CodeGuruProfilerUntagResourceFuture) Get(ctx workflow.Context) (*codeguruprofiler.UntagResourceOutput, error) {
	var output codeguruprofiler.UntagResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CodeGuruProfilerUpdateProfilingGroupFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CodeGuruProfilerUpdateProfilingGroupFuture) Get(ctx workflow.Context) (*codeguruprofiler.UpdateProfilingGroupOutput, error) {
	var output codeguruprofiler.UpdateProfilingGroupOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *stub) AddNotificationChannels(ctx workflow.Context, input *codeguruprofiler.AddNotificationChannelsInput) (*codeguruprofiler.AddNotificationChannelsOutput, error) {
	var output codeguruprofiler.AddNotificationChannelsOutput
	err := workflow.ExecuteActivity(ctx, "aws-codeguruprofiler-AddNotificationChannels", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) AddNotificationChannelsAsync(ctx workflow.Context, input *codeguruprofiler.AddNotificationChannelsInput) *CodeGuruProfilerAddNotificationChannelsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-codeguruprofiler-AddNotificationChannels", input)
	return &CodeGuruProfilerAddNotificationChannelsFuture{Future: future}
}

func (a *stub) BatchGetFrameMetricData(ctx workflow.Context, input *codeguruprofiler.BatchGetFrameMetricDataInput) (*codeguruprofiler.BatchGetFrameMetricDataOutput, error) {
	var output codeguruprofiler.BatchGetFrameMetricDataOutput
	err := workflow.ExecuteActivity(ctx, "aws-codeguruprofiler-BatchGetFrameMetricData", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) BatchGetFrameMetricDataAsync(ctx workflow.Context, input *codeguruprofiler.BatchGetFrameMetricDataInput) *CodeGuruProfilerBatchGetFrameMetricDataFuture {
	future := workflow.ExecuteActivity(ctx, "aws-codeguruprofiler-BatchGetFrameMetricData", input)
	return &CodeGuruProfilerBatchGetFrameMetricDataFuture{Future: future}
}

func (a *stub) ConfigureAgent(ctx workflow.Context, input *codeguruprofiler.ConfigureAgentInput) (*codeguruprofiler.ConfigureAgentOutput, error) {
	var output codeguruprofiler.ConfigureAgentOutput
	err := workflow.ExecuteActivity(ctx, "aws-codeguruprofiler-ConfigureAgent", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ConfigureAgentAsync(ctx workflow.Context, input *codeguruprofiler.ConfigureAgentInput) *CodeGuruProfilerConfigureAgentFuture {
	future := workflow.ExecuteActivity(ctx, "aws-codeguruprofiler-ConfigureAgent", input)
	return &CodeGuruProfilerConfigureAgentFuture{Future: future}
}

func (a *stub) CreateProfilingGroup(ctx workflow.Context, input *codeguruprofiler.CreateProfilingGroupInput) (*codeguruprofiler.CreateProfilingGroupOutput, error) {
	var output codeguruprofiler.CreateProfilingGroupOutput
	err := workflow.ExecuteActivity(ctx, "aws-codeguruprofiler-CreateProfilingGroup", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateProfilingGroupAsync(ctx workflow.Context, input *codeguruprofiler.CreateProfilingGroupInput) *CodeGuruProfilerCreateProfilingGroupFuture {
	future := workflow.ExecuteActivity(ctx, "aws-codeguruprofiler-CreateProfilingGroup", input)
	return &CodeGuruProfilerCreateProfilingGroupFuture{Future: future}
}

func (a *stub) DeleteProfilingGroup(ctx workflow.Context, input *codeguruprofiler.DeleteProfilingGroupInput) (*codeguruprofiler.DeleteProfilingGroupOutput, error) {
	var output codeguruprofiler.DeleteProfilingGroupOutput
	err := workflow.ExecuteActivity(ctx, "aws-codeguruprofiler-DeleteProfilingGroup", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteProfilingGroupAsync(ctx workflow.Context, input *codeguruprofiler.DeleteProfilingGroupInput) *CodeGuruProfilerDeleteProfilingGroupFuture {
	future := workflow.ExecuteActivity(ctx, "aws-codeguruprofiler-DeleteProfilingGroup", input)
	return &CodeGuruProfilerDeleteProfilingGroupFuture{Future: future}
}

func (a *stub) DescribeProfilingGroup(ctx workflow.Context, input *codeguruprofiler.DescribeProfilingGroupInput) (*codeguruprofiler.DescribeProfilingGroupOutput, error) {
	var output codeguruprofiler.DescribeProfilingGroupOutput
	err := workflow.ExecuteActivity(ctx, "aws-codeguruprofiler-DescribeProfilingGroup", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeProfilingGroupAsync(ctx workflow.Context, input *codeguruprofiler.DescribeProfilingGroupInput) *CodeGuruProfilerDescribeProfilingGroupFuture {
	future := workflow.ExecuteActivity(ctx, "aws-codeguruprofiler-DescribeProfilingGroup", input)
	return &CodeGuruProfilerDescribeProfilingGroupFuture{Future: future}
}

func (a *stub) GetFindingsReportAccountSummary(ctx workflow.Context, input *codeguruprofiler.GetFindingsReportAccountSummaryInput) (*codeguruprofiler.GetFindingsReportAccountSummaryOutput, error) {
	var output codeguruprofiler.GetFindingsReportAccountSummaryOutput
	err := workflow.ExecuteActivity(ctx, "aws-codeguruprofiler-GetFindingsReportAccountSummary", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetFindingsReportAccountSummaryAsync(ctx workflow.Context, input *codeguruprofiler.GetFindingsReportAccountSummaryInput) *CodeGuruProfilerGetFindingsReportAccountSummaryFuture {
	future := workflow.ExecuteActivity(ctx, "aws-codeguruprofiler-GetFindingsReportAccountSummary", input)
	return &CodeGuruProfilerGetFindingsReportAccountSummaryFuture{Future: future}
}

func (a *stub) GetNotificationConfiguration(ctx workflow.Context, input *codeguruprofiler.GetNotificationConfigurationInput) (*codeguruprofiler.GetNotificationConfigurationOutput, error) {
	var output codeguruprofiler.GetNotificationConfigurationOutput
	err := workflow.ExecuteActivity(ctx, "aws-codeguruprofiler-GetNotificationConfiguration", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetNotificationConfigurationAsync(ctx workflow.Context, input *codeguruprofiler.GetNotificationConfigurationInput) *CodeGuruProfilerGetNotificationConfigurationFuture {
	future := workflow.ExecuteActivity(ctx, "aws-codeguruprofiler-GetNotificationConfiguration", input)
	return &CodeGuruProfilerGetNotificationConfigurationFuture{Future: future}
}

func (a *stub) GetPolicy(ctx workflow.Context, input *codeguruprofiler.GetPolicyInput) (*codeguruprofiler.GetPolicyOutput, error) {
	var output codeguruprofiler.GetPolicyOutput
	err := workflow.ExecuteActivity(ctx, "aws-codeguruprofiler-GetPolicy", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetPolicyAsync(ctx workflow.Context, input *codeguruprofiler.GetPolicyInput) *CodeGuruProfilerGetPolicyFuture {
	future := workflow.ExecuteActivity(ctx, "aws-codeguruprofiler-GetPolicy", input)
	return &CodeGuruProfilerGetPolicyFuture{Future: future}
}

func (a *stub) GetProfile(ctx workflow.Context, input *codeguruprofiler.GetProfileInput) (*codeguruprofiler.GetProfileOutput, error) {
	var output codeguruprofiler.GetProfileOutput
	err := workflow.ExecuteActivity(ctx, "aws-codeguruprofiler-GetProfile", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetProfileAsync(ctx workflow.Context, input *codeguruprofiler.GetProfileInput) *CodeGuruProfilerGetProfileFuture {
	future := workflow.ExecuteActivity(ctx, "aws-codeguruprofiler-GetProfile", input)
	return &CodeGuruProfilerGetProfileFuture{Future: future}
}

func (a *stub) GetRecommendations(ctx workflow.Context, input *codeguruprofiler.GetRecommendationsInput) (*codeguruprofiler.GetRecommendationsOutput, error) {
	var output codeguruprofiler.GetRecommendationsOutput
	err := workflow.ExecuteActivity(ctx, "aws-codeguruprofiler-GetRecommendations", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetRecommendationsAsync(ctx workflow.Context, input *codeguruprofiler.GetRecommendationsInput) *CodeGuruProfilerGetRecommendationsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-codeguruprofiler-GetRecommendations", input)
	return &CodeGuruProfilerGetRecommendationsFuture{Future: future}
}

func (a *stub) ListFindingsReports(ctx workflow.Context, input *codeguruprofiler.ListFindingsReportsInput) (*codeguruprofiler.ListFindingsReportsOutput, error) {
	var output codeguruprofiler.ListFindingsReportsOutput
	err := workflow.ExecuteActivity(ctx, "aws-codeguruprofiler-ListFindingsReports", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListFindingsReportsAsync(ctx workflow.Context, input *codeguruprofiler.ListFindingsReportsInput) *CodeGuruProfilerListFindingsReportsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-codeguruprofiler-ListFindingsReports", input)
	return &CodeGuruProfilerListFindingsReportsFuture{Future: future}
}

func (a *stub) ListProfileTimes(ctx workflow.Context, input *codeguruprofiler.ListProfileTimesInput) (*codeguruprofiler.ListProfileTimesOutput, error) {
	var output codeguruprofiler.ListProfileTimesOutput
	err := workflow.ExecuteActivity(ctx, "aws-codeguruprofiler-ListProfileTimes", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListProfileTimesAsync(ctx workflow.Context, input *codeguruprofiler.ListProfileTimesInput) *CodeGuruProfilerListProfileTimesFuture {
	future := workflow.ExecuteActivity(ctx, "aws-codeguruprofiler-ListProfileTimes", input)
	return &CodeGuruProfilerListProfileTimesFuture{Future: future}
}

func (a *stub) ListProfilingGroups(ctx workflow.Context, input *codeguruprofiler.ListProfilingGroupsInput) (*codeguruprofiler.ListProfilingGroupsOutput, error) {
	var output codeguruprofiler.ListProfilingGroupsOutput
	err := workflow.ExecuteActivity(ctx, "aws-codeguruprofiler-ListProfilingGroups", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListProfilingGroupsAsync(ctx workflow.Context, input *codeguruprofiler.ListProfilingGroupsInput) *CodeGuruProfilerListProfilingGroupsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-codeguruprofiler-ListProfilingGroups", input)
	return &CodeGuruProfilerListProfilingGroupsFuture{Future: future}
}

func (a *stub) ListTagsForResource(ctx workflow.Context, input *codeguruprofiler.ListTagsForResourceInput) (*codeguruprofiler.ListTagsForResourceOutput, error) {
	var output codeguruprofiler.ListTagsForResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws-codeguruprofiler-ListTagsForResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListTagsForResourceAsync(ctx workflow.Context, input *codeguruprofiler.ListTagsForResourceInput) *CodeGuruProfilerListTagsForResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws-codeguruprofiler-ListTagsForResource", input)
	return &CodeGuruProfilerListTagsForResourceFuture{Future: future}
}

func (a *stub) PostAgentProfile(ctx workflow.Context, input *codeguruprofiler.PostAgentProfileInput) (*codeguruprofiler.PostAgentProfileOutput, error) {
	var output codeguruprofiler.PostAgentProfileOutput
	err := workflow.ExecuteActivity(ctx, "aws-codeguruprofiler-PostAgentProfile", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) PostAgentProfileAsync(ctx workflow.Context, input *codeguruprofiler.PostAgentProfileInput) *CodeGuruProfilerPostAgentProfileFuture {
	future := workflow.ExecuteActivity(ctx, "aws-codeguruprofiler-PostAgentProfile", input)
	return &CodeGuruProfilerPostAgentProfileFuture{Future: future}
}

func (a *stub) PutPermission(ctx workflow.Context, input *codeguruprofiler.PutPermissionInput) (*codeguruprofiler.PutPermissionOutput, error) {
	var output codeguruprofiler.PutPermissionOutput
	err := workflow.ExecuteActivity(ctx, "aws-codeguruprofiler-PutPermission", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) PutPermissionAsync(ctx workflow.Context, input *codeguruprofiler.PutPermissionInput) *CodeGuruProfilerPutPermissionFuture {
	future := workflow.ExecuteActivity(ctx, "aws-codeguruprofiler-PutPermission", input)
	return &CodeGuruProfilerPutPermissionFuture{Future: future}
}

func (a *stub) RemoveNotificationChannel(ctx workflow.Context, input *codeguruprofiler.RemoveNotificationChannelInput) (*codeguruprofiler.RemoveNotificationChannelOutput, error) {
	var output codeguruprofiler.RemoveNotificationChannelOutput
	err := workflow.ExecuteActivity(ctx, "aws-codeguruprofiler-RemoveNotificationChannel", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) RemoveNotificationChannelAsync(ctx workflow.Context, input *codeguruprofiler.RemoveNotificationChannelInput) *CodeGuruProfilerRemoveNotificationChannelFuture {
	future := workflow.ExecuteActivity(ctx, "aws-codeguruprofiler-RemoveNotificationChannel", input)
	return &CodeGuruProfilerRemoveNotificationChannelFuture{Future: future}
}

func (a *stub) RemovePermission(ctx workflow.Context, input *codeguruprofiler.RemovePermissionInput) (*codeguruprofiler.RemovePermissionOutput, error) {
	var output codeguruprofiler.RemovePermissionOutput
	err := workflow.ExecuteActivity(ctx, "aws-codeguruprofiler-RemovePermission", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) RemovePermissionAsync(ctx workflow.Context, input *codeguruprofiler.RemovePermissionInput) *CodeGuruProfilerRemovePermissionFuture {
	future := workflow.ExecuteActivity(ctx, "aws-codeguruprofiler-RemovePermission", input)
	return &CodeGuruProfilerRemovePermissionFuture{Future: future}
}

func (a *stub) SubmitFeedback(ctx workflow.Context, input *codeguruprofiler.SubmitFeedbackInput) (*codeguruprofiler.SubmitFeedbackOutput, error) {
	var output codeguruprofiler.SubmitFeedbackOutput
	err := workflow.ExecuteActivity(ctx, "aws-codeguruprofiler-SubmitFeedback", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) SubmitFeedbackAsync(ctx workflow.Context, input *codeguruprofiler.SubmitFeedbackInput) *CodeGuruProfilerSubmitFeedbackFuture {
	future := workflow.ExecuteActivity(ctx, "aws-codeguruprofiler-SubmitFeedback", input)
	return &CodeGuruProfilerSubmitFeedbackFuture{Future: future}
}

func (a *stub) TagResource(ctx workflow.Context, input *codeguruprofiler.TagResourceInput) (*codeguruprofiler.TagResourceOutput, error) {
	var output codeguruprofiler.TagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws-codeguruprofiler-TagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) TagResourceAsync(ctx workflow.Context, input *codeguruprofiler.TagResourceInput) *CodeGuruProfilerTagResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws-codeguruprofiler-TagResource", input)
	return &CodeGuruProfilerTagResourceFuture{Future: future}
}

func (a *stub) UntagResource(ctx workflow.Context, input *codeguruprofiler.UntagResourceInput) (*codeguruprofiler.UntagResourceOutput, error) {
	var output codeguruprofiler.UntagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws-codeguruprofiler-UntagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UntagResourceAsync(ctx workflow.Context, input *codeguruprofiler.UntagResourceInput) *CodeGuruProfilerUntagResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws-codeguruprofiler-UntagResource", input)
	return &CodeGuruProfilerUntagResourceFuture{Future: future}
}

func (a *stub) UpdateProfilingGroup(ctx workflow.Context, input *codeguruprofiler.UpdateProfilingGroupInput) (*codeguruprofiler.UpdateProfilingGroupOutput, error) {
	var output codeguruprofiler.UpdateProfilingGroupOutput
	err := workflow.ExecuteActivity(ctx, "aws-codeguruprofiler-UpdateProfilingGroup", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateProfilingGroupAsync(ctx workflow.Context, input *codeguruprofiler.UpdateProfilingGroupInput) *CodeGuruProfilerUpdateProfilingGroupFuture {
	future := workflow.ExecuteActivity(ctx, "aws-codeguruprofiler-UpdateProfilingGroup", input)
	return &CodeGuruProfilerUpdateProfilingGroupFuture{Future: future}
}
