package cloudformation

import (
	"context"

	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/aws/aws-sdk-go/service/cloudformation/cloudformationiface"
	"github.com/banzaicloud/cadence-aws-sdk/workflow/internal/errors"
)

// Activities implements Cadence activities for the API operation methods of this AWS service.
type Activities struct {
	Client cloudformationiface.CloudFormationAPI
}

// DescribeStackResourcesInput is the input for the DescribeStackResources API operation.
type DescribeStackResourcesInput struct {
	Input *cloudformation.DescribeStackResourcesInput
}

// DescribeStackResourcesOutput is the output for the DescribeStackResources API operation.
type DescribeStackResourcesOutput struct {
	Output *cloudformation.DescribeStackResourcesOutput
}

// DescribeStackResources implements the DescribeStackResources API operation.
func (a Activities) DescribeStackResources(ctx context.Context, input DescribeStackResourcesInput) (*DescribeStackResourcesOutput, error) {
	output, err := a.Client.DescribeStackResourcesWithContext(ctx, input.Input)
	if err != nil {
		return nil, errors.EncodeError(err)
	}

	return &DescribeStackResourcesOutput{Output: output}, nil
}
