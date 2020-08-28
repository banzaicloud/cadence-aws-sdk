package cloudformation

import (
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/banzaicloud/cadence-aws-sdk/sdk/internal/errors"
	cloudformationwf "github.com/banzaicloud/cadence-aws-sdk/workflow/service/cloudformation"
	"go.uber.org/cadence/workflow"
)

// DescribeStackResourcesInput is the input for the DescribeStackResources operation.
type DescribeStackResourcesInput = cloudformation.DescribeStackResourcesInput

// DescribeStackResourcesOutput is the output for the DescribeStackResources operation.
type DescribeStackResourcesOutput = cloudformation.DescribeStackResourcesOutput

// CloudFormation provides the API operation methods for making requests to AWS in Cadence workflows.
type CloudFormation interface {
	DescribeStackResources(ctx workflow.Context, input *DescribeStackResourcesInput) (*DescribeStackResourcesOutput, error)
}

type cloudFormation struct{}

// New creates a new instance of the CloudFormation client.
func New() CloudFormation {
	return cloudFormation{}
}

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
