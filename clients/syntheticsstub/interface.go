// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7

package syntheticsstub

import (
	"github.com/aws/aws-sdk-go/service/synthetics"
	"go.uber.org/cadence/workflow"

	"github.com/banzaicloud/cadence-aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	CreateCanary(ctx workflow.Context, input *synthetics.CreateCanaryInput) (*synthetics.CreateCanaryOutput, error)
	CreateCanaryAsync(ctx workflow.Context, input *synthetics.CreateCanaryInput) *SyntheticsCreateCanaryFuture

	DeleteCanary(ctx workflow.Context, input *synthetics.DeleteCanaryInput) (*synthetics.DeleteCanaryOutput, error)
	DeleteCanaryAsync(ctx workflow.Context, input *synthetics.DeleteCanaryInput) *SyntheticsDeleteCanaryFuture

	DescribeCanaries(ctx workflow.Context, input *synthetics.DescribeCanariesInput) (*synthetics.DescribeCanariesOutput, error)
	DescribeCanariesAsync(ctx workflow.Context, input *synthetics.DescribeCanariesInput) *SyntheticsDescribeCanariesFuture

	DescribeCanariesLastRun(ctx workflow.Context, input *synthetics.DescribeCanariesLastRunInput) (*synthetics.DescribeCanariesLastRunOutput, error)
	DescribeCanariesLastRunAsync(ctx workflow.Context, input *synthetics.DescribeCanariesLastRunInput) *SyntheticsDescribeCanariesLastRunFuture

	DescribeRuntimeVersions(ctx workflow.Context, input *synthetics.DescribeRuntimeVersionsInput) (*synthetics.DescribeRuntimeVersionsOutput, error)
	DescribeRuntimeVersionsAsync(ctx workflow.Context, input *synthetics.DescribeRuntimeVersionsInput) *SyntheticsDescribeRuntimeVersionsFuture

	GetCanary(ctx workflow.Context, input *synthetics.GetCanaryInput) (*synthetics.GetCanaryOutput, error)
	GetCanaryAsync(ctx workflow.Context, input *synthetics.GetCanaryInput) *SyntheticsGetCanaryFuture

	GetCanaryRuns(ctx workflow.Context, input *synthetics.GetCanaryRunsInput) (*synthetics.GetCanaryRunsOutput, error)
	GetCanaryRunsAsync(ctx workflow.Context, input *synthetics.GetCanaryRunsInput) *SyntheticsGetCanaryRunsFuture

	ListTagsForResource(ctx workflow.Context, input *synthetics.ListTagsForResourceInput) (*synthetics.ListTagsForResourceOutput, error)
	ListTagsForResourceAsync(ctx workflow.Context, input *synthetics.ListTagsForResourceInput) *SyntheticsListTagsForResourceFuture

	StartCanary(ctx workflow.Context, input *synthetics.StartCanaryInput) (*synthetics.StartCanaryOutput, error)
	StartCanaryAsync(ctx workflow.Context, input *synthetics.StartCanaryInput) *SyntheticsStartCanaryFuture

	StopCanary(ctx workflow.Context, input *synthetics.StopCanaryInput) (*synthetics.StopCanaryOutput, error)
	StopCanaryAsync(ctx workflow.Context, input *synthetics.StopCanaryInput) *SyntheticsStopCanaryFuture

	TagResource(ctx workflow.Context, input *synthetics.TagResourceInput) (*synthetics.TagResourceOutput, error)
	TagResourceAsync(ctx workflow.Context, input *synthetics.TagResourceInput) *SyntheticsTagResourceFuture

	UntagResource(ctx workflow.Context, input *synthetics.UntagResourceInput) (*synthetics.UntagResourceOutput, error)
	UntagResourceAsync(ctx workflow.Context, input *synthetics.UntagResourceInput) *SyntheticsUntagResourceFuture

	UpdateCanary(ctx workflow.Context, input *synthetics.UpdateCanaryInput) (*synthetics.UpdateCanaryOutput, error)
	UpdateCanaryAsync(ctx workflow.Context, input *synthetics.UpdateCanaryInput) *SyntheticsUpdateCanaryFuture
}

func NewClient() Client {
	return &stub{}
}
