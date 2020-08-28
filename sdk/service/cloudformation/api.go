package cloudformation

import (
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/banzaicloud/cadence-aws-sdk/sdk/internal/errors"
	cloudformationwf "github.com/banzaicloud/cadence-aws-sdk/workflow/service/cloudformation"
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

func (c cloudFormation) DescribeStackResources(ctx workflow.Context, input *DescribeStackResourcesInput) (*DescribeStackResourcesOutput, error) {
	in := cloudformationwf.DescribeStackResourcesInput{
		Input: input,
	}

	var output cloudformationwf.DescribeStackResourcesOutput

	err := workflow.ExecuteActivity(ctx, "aws-cloudformation-DescribeStackResources", in).Get(ctx, &output) // nolint: lll
	if err != nil {
		return nil, errors.DecodeError(err)
	}

	return output.Output, nil
}
