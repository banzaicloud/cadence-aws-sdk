// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7

package iotdataplanestub

import (
	"github.com/aws/aws-sdk-go/service/iotdataplane"
	"go.uber.org/cadence/workflow"

	"github.com/banzaicloud/cadence-aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	DeleteThingShadow(ctx workflow.Context, input *iotdataplane.DeleteThingShadowInput) (*iotdataplane.DeleteThingShadowOutput, error)
	DeleteThingShadowAsync(ctx workflow.Context, input *iotdataplane.DeleteThingShadowInput) *IoTDataPlaneDeleteThingShadowFuture

	GetThingShadow(ctx workflow.Context, input *iotdataplane.GetThingShadowInput) (*iotdataplane.GetThingShadowOutput, error)
	GetThingShadowAsync(ctx workflow.Context, input *iotdataplane.GetThingShadowInput) *IoTDataPlaneGetThingShadowFuture

	ListNamedShadowsForThing(ctx workflow.Context, input *iotdataplane.ListNamedShadowsForThingInput) (*iotdataplane.ListNamedShadowsForThingOutput, error)
	ListNamedShadowsForThingAsync(ctx workflow.Context, input *iotdataplane.ListNamedShadowsForThingInput) *IoTDataPlaneListNamedShadowsForThingFuture

	Publish(ctx workflow.Context, input *iotdataplane.PublishInput) (*iotdataplane.PublishOutput, error)
	PublishAsync(ctx workflow.Context, input *iotdataplane.PublishInput) *IoTDataPlanePublishFuture

	UpdateThingShadow(ctx workflow.Context, input *iotdataplane.UpdateThingShadowInput) (*iotdataplane.UpdateThingShadowOutput, error)
	UpdateThingShadowAsync(ctx workflow.Context, input *iotdataplane.UpdateThingShadowInput) *IoTDataPlaneUpdateThingShadowFuture
}

func NewClient() Client {
	return &stub{}
}
