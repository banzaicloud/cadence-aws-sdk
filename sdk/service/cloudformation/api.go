package cloudformation

import (
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/banzaicloud/cadence-aws-sdk/sdk/internal/errors"
	"go.uber.org/cadence/workflow"
)

// DescribeStackResourcesInput is the input for the DescribeStackResources operation.
type DescribeStackResourcesInput = cloudformation.DescribeStackResourcesInput

// DescribeStackResourcesOutput  is the output for the DescribeStackResources operation.
type DescribeStackResourcesOutput = cloudformation.DescribeStackResourcesOutput

// CloudFormation implements AWS CF operations and wraps them in a Cadence specific context.
type CloudFormation interface {
	// DescribeStackResources implements the DescribeStackResources API operation.
	DescribeStackResources(ctx workflow.Context, input *DescribeStackResourcesInput) (*DescribeStackResourcesOutput, error)
}

type cloudFormation struct{}

const describeStackResourcesActivityName = "aws-cloudformation-DescribeStackResources"

type describeStackResourcesActivityInput struct {
	Input *DescribeStackResourcesInput
}

type describeStackResourcesActivityOutput struct {
	Output *DescribeStackResourcesOutput
}

func (c cloudFormation) DescribeStackResources(ctx workflow.Context, input *DescribeStackResourcesInput) (*DescribeStackResourcesOutput, error) {
	in := describeStackResourcesActivityInput{
		Input: input,
	}

	var output describeStackResourcesActivityOutput

	err := workflow.ExecuteActivity(ctx, describeStackResourcesActivityName, in).Get(ctx, &output) // nolint: lll
	if err != nil {
		return nil, errors.DecodeError(err)
	}

	return output.Output, nil
}
