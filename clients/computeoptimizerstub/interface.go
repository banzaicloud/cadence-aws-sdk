// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.
// Copyright (c) 2020 Banzai Cloud All rights reserved.

package computeoptimizerstub

import (
	"github.com/aws/aws-sdk-go/service/computeoptimizer"
	"go.uber.org/cadence/workflow"

	"github.com/banzaicloud/cadence-aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	DescribeRecommendationExportJobs(ctx workflow.Context, input *computeoptimizer.DescribeRecommendationExportJobsInput) (*computeoptimizer.DescribeRecommendationExportJobsOutput, error)
	DescribeRecommendationExportJobsAsync(ctx workflow.Context, input *computeoptimizer.DescribeRecommendationExportJobsInput) *ComputeOptimizerDescribeRecommendationExportJobsFuture

	ExportAutoScalingGroupRecommendations(ctx workflow.Context, input *computeoptimizer.ExportAutoScalingGroupRecommendationsInput) (*computeoptimizer.ExportAutoScalingGroupRecommendationsOutput, error)
	ExportAutoScalingGroupRecommendationsAsync(ctx workflow.Context, input *computeoptimizer.ExportAutoScalingGroupRecommendationsInput) *ComputeOptimizerExportAutoScalingGroupRecommendationsFuture

	ExportEC2InstanceRecommendations(ctx workflow.Context, input *computeoptimizer.ExportEC2InstanceRecommendationsInput) (*computeoptimizer.ExportEC2InstanceRecommendationsOutput, error)
	ExportEC2InstanceRecommendationsAsync(ctx workflow.Context, input *computeoptimizer.ExportEC2InstanceRecommendationsInput) *ComputeOptimizerExportEC2InstanceRecommendationsFuture

	GetAutoScalingGroupRecommendations(ctx workflow.Context, input *computeoptimizer.GetAutoScalingGroupRecommendationsInput) (*computeoptimizer.GetAutoScalingGroupRecommendationsOutput, error)
	GetAutoScalingGroupRecommendationsAsync(ctx workflow.Context, input *computeoptimizer.GetAutoScalingGroupRecommendationsInput) *ComputeOptimizerGetAutoScalingGroupRecommendationsFuture

	GetEC2InstanceRecommendations(ctx workflow.Context, input *computeoptimizer.GetEC2InstanceRecommendationsInput) (*computeoptimizer.GetEC2InstanceRecommendationsOutput, error)
	GetEC2InstanceRecommendationsAsync(ctx workflow.Context, input *computeoptimizer.GetEC2InstanceRecommendationsInput) *ComputeOptimizerGetEC2InstanceRecommendationsFuture

	GetEC2RecommendationProjectedMetrics(ctx workflow.Context, input *computeoptimizer.GetEC2RecommendationProjectedMetricsInput) (*computeoptimizer.GetEC2RecommendationProjectedMetricsOutput, error)
	GetEC2RecommendationProjectedMetricsAsync(ctx workflow.Context, input *computeoptimizer.GetEC2RecommendationProjectedMetricsInput) *ComputeOptimizerGetEC2RecommendationProjectedMetricsFuture

	GetEnrollmentStatus(ctx workflow.Context, input *computeoptimizer.GetEnrollmentStatusInput) (*computeoptimizer.GetEnrollmentStatusOutput, error)
	GetEnrollmentStatusAsync(ctx workflow.Context, input *computeoptimizer.GetEnrollmentStatusInput) *ComputeOptimizerGetEnrollmentStatusFuture

	GetRecommendationSummaries(ctx workflow.Context, input *computeoptimizer.GetRecommendationSummariesInput) (*computeoptimizer.GetRecommendationSummariesOutput, error)
	GetRecommendationSummariesAsync(ctx workflow.Context, input *computeoptimizer.GetRecommendationSummariesInput) *ComputeOptimizerGetRecommendationSummariesFuture

	UpdateEnrollmentStatus(ctx workflow.Context, input *computeoptimizer.UpdateEnrollmentStatusInput) (*computeoptimizer.UpdateEnrollmentStatusOutput, error)
	UpdateEnrollmentStatusAsync(ctx workflow.Context, input *computeoptimizer.UpdateEnrollmentStatusInput) *ComputeOptimizerUpdateEnrollmentStatusFuture
}

func NewClient() Client {
	return &stub{}
}
