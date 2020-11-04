// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.
// Copyright (c) 2020 Banzai Cloud All rights reserved.

package codeguruprofilerstub

import (
	"github.com/aws/aws-sdk-go/service/codeguruprofiler"
	"go.uber.org/cadence/workflow"

	"github.com/banzaicloud/cadence-aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	AddNotificationChannels(ctx workflow.Context, input *codeguruprofiler.AddNotificationChannelsInput) (*codeguruprofiler.AddNotificationChannelsOutput, error)
	AddNotificationChannelsAsync(ctx workflow.Context, input *codeguruprofiler.AddNotificationChannelsInput) *CodeGuruProfilerAddNotificationChannelsFuture

	BatchGetFrameMetricData(ctx workflow.Context, input *codeguruprofiler.BatchGetFrameMetricDataInput) (*codeguruprofiler.BatchGetFrameMetricDataOutput, error)
	BatchGetFrameMetricDataAsync(ctx workflow.Context, input *codeguruprofiler.BatchGetFrameMetricDataInput) *CodeGuruProfilerBatchGetFrameMetricDataFuture

	ConfigureAgent(ctx workflow.Context, input *codeguruprofiler.ConfigureAgentInput) (*codeguruprofiler.ConfigureAgentOutput, error)
	ConfigureAgentAsync(ctx workflow.Context, input *codeguruprofiler.ConfigureAgentInput) *CodeGuruProfilerConfigureAgentFuture

	CreateProfilingGroup(ctx workflow.Context, input *codeguruprofiler.CreateProfilingGroupInput) (*codeguruprofiler.CreateProfilingGroupOutput, error)
	CreateProfilingGroupAsync(ctx workflow.Context, input *codeguruprofiler.CreateProfilingGroupInput) *CodeGuruProfilerCreateProfilingGroupFuture

	DeleteProfilingGroup(ctx workflow.Context, input *codeguruprofiler.DeleteProfilingGroupInput) (*codeguruprofiler.DeleteProfilingGroupOutput, error)
	DeleteProfilingGroupAsync(ctx workflow.Context, input *codeguruprofiler.DeleteProfilingGroupInput) *CodeGuruProfilerDeleteProfilingGroupFuture

	DescribeProfilingGroup(ctx workflow.Context, input *codeguruprofiler.DescribeProfilingGroupInput) (*codeguruprofiler.DescribeProfilingGroupOutput, error)
	DescribeProfilingGroupAsync(ctx workflow.Context, input *codeguruprofiler.DescribeProfilingGroupInput) *CodeGuruProfilerDescribeProfilingGroupFuture

	GetFindingsReportAccountSummary(ctx workflow.Context, input *codeguruprofiler.GetFindingsReportAccountSummaryInput) (*codeguruprofiler.GetFindingsReportAccountSummaryOutput, error)
	GetFindingsReportAccountSummaryAsync(ctx workflow.Context, input *codeguruprofiler.GetFindingsReportAccountSummaryInput) *CodeGuruProfilerGetFindingsReportAccountSummaryFuture

	GetNotificationConfiguration(ctx workflow.Context, input *codeguruprofiler.GetNotificationConfigurationInput) (*codeguruprofiler.GetNotificationConfigurationOutput, error)
	GetNotificationConfigurationAsync(ctx workflow.Context, input *codeguruprofiler.GetNotificationConfigurationInput) *CodeGuruProfilerGetNotificationConfigurationFuture

	GetPolicy(ctx workflow.Context, input *codeguruprofiler.GetPolicyInput) (*codeguruprofiler.GetPolicyOutput, error)
	GetPolicyAsync(ctx workflow.Context, input *codeguruprofiler.GetPolicyInput) *CodeGuruProfilerGetPolicyFuture

	GetProfile(ctx workflow.Context, input *codeguruprofiler.GetProfileInput) (*codeguruprofiler.GetProfileOutput, error)
	GetProfileAsync(ctx workflow.Context, input *codeguruprofiler.GetProfileInput) *CodeGuruProfilerGetProfileFuture

	GetRecommendations(ctx workflow.Context, input *codeguruprofiler.GetRecommendationsInput) (*codeguruprofiler.GetRecommendationsOutput, error)
	GetRecommendationsAsync(ctx workflow.Context, input *codeguruprofiler.GetRecommendationsInput) *CodeGuruProfilerGetRecommendationsFuture

	ListFindingsReports(ctx workflow.Context, input *codeguruprofiler.ListFindingsReportsInput) (*codeguruprofiler.ListFindingsReportsOutput, error)
	ListFindingsReportsAsync(ctx workflow.Context, input *codeguruprofiler.ListFindingsReportsInput) *CodeGuruProfilerListFindingsReportsFuture

	ListProfileTimes(ctx workflow.Context, input *codeguruprofiler.ListProfileTimesInput) (*codeguruprofiler.ListProfileTimesOutput, error)
	ListProfileTimesAsync(ctx workflow.Context, input *codeguruprofiler.ListProfileTimesInput) *CodeGuruProfilerListProfileTimesFuture

	ListProfilingGroups(ctx workflow.Context, input *codeguruprofiler.ListProfilingGroupsInput) (*codeguruprofiler.ListProfilingGroupsOutput, error)
	ListProfilingGroupsAsync(ctx workflow.Context, input *codeguruprofiler.ListProfilingGroupsInput) *CodeGuruProfilerListProfilingGroupsFuture

	ListTagsForResource(ctx workflow.Context, input *codeguruprofiler.ListTagsForResourceInput) (*codeguruprofiler.ListTagsForResourceOutput, error)
	ListTagsForResourceAsync(ctx workflow.Context, input *codeguruprofiler.ListTagsForResourceInput) *CodeGuruProfilerListTagsForResourceFuture

	PostAgentProfile(ctx workflow.Context, input *codeguruprofiler.PostAgentProfileInput) (*codeguruprofiler.PostAgentProfileOutput, error)
	PostAgentProfileAsync(ctx workflow.Context, input *codeguruprofiler.PostAgentProfileInput) *CodeGuruProfilerPostAgentProfileFuture

	PutPermission(ctx workflow.Context, input *codeguruprofiler.PutPermissionInput) (*codeguruprofiler.PutPermissionOutput, error)
	PutPermissionAsync(ctx workflow.Context, input *codeguruprofiler.PutPermissionInput) *CodeGuruProfilerPutPermissionFuture

	RemoveNotificationChannel(ctx workflow.Context, input *codeguruprofiler.RemoveNotificationChannelInput) (*codeguruprofiler.RemoveNotificationChannelOutput, error)
	RemoveNotificationChannelAsync(ctx workflow.Context, input *codeguruprofiler.RemoveNotificationChannelInput) *CodeGuruProfilerRemoveNotificationChannelFuture

	RemovePermission(ctx workflow.Context, input *codeguruprofiler.RemovePermissionInput) (*codeguruprofiler.RemovePermissionOutput, error)
	RemovePermissionAsync(ctx workflow.Context, input *codeguruprofiler.RemovePermissionInput) *CodeGuruProfilerRemovePermissionFuture

	SubmitFeedback(ctx workflow.Context, input *codeguruprofiler.SubmitFeedbackInput) (*codeguruprofiler.SubmitFeedbackOutput, error)
	SubmitFeedbackAsync(ctx workflow.Context, input *codeguruprofiler.SubmitFeedbackInput) *CodeGuruProfilerSubmitFeedbackFuture

	TagResource(ctx workflow.Context, input *codeguruprofiler.TagResourceInput) (*codeguruprofiler.TagResourceOutput, error)
	TagResourceAsync(ctx workflow.Context, input *codeguruprofiler.TagResourceInput) *CodeGuruProfilerTagResourceFuture

	UntagResource(ctx workflow.Context, input *codeguruprofiler.UntagResourceInput) (*codeguruprofiler.UntagResourceOutput, error)
	UntagResourceAsync(ctx workflow.Context, input *codeguruprofiler.UntagResourceInput) *CodeGuruProfilerUntagResourceFuture

	UpdateProfilingGroup(ctx workflow.Context, input *codeguruprofiler.UpdateProfilingGroupInput) (*codeguruprofiler.UpdateProfilingGroupOutput, error)
	UpdateProfilingGroupAsync(ctx workflow.Context, input *codeguruprofiler.UpdateProfilingGroupInput) *CodeGuruProfilerUpdateProfilingGroupFuture
}

func NewClient() Client {
	return &stub{}
}
