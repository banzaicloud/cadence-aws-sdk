// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.
// Copyright (c) 2020 Banzai Cloud All rights reserved.

package ioteventsstub

import (
	"github.com/aws/aws-sdk-go/service/iotevents"
	"go.uber.org/cadence/workflow"

	"github.com/banzaicloud/cadence-aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type stub struct{}

type IoTEventsCreateDetectorModelFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *IoTEventsCreateDetectorModelFuture) Get(ctx workflow.Context) (*iotevents.CreateDetectorModelOutput, error) {
	var output iotevents.CreateDetectorModelOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type IoTEventsCreateInputFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *IoTEventsCreateInputFuture) Get(ctx workflow.Context) (*iotevents.CreateInputOutput, error) {
	var output iotevents.CreateInputOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type IoTEventsDeleteDetectorModelFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *IoTEventsDeleteDetectorModelFuture) Get(ctx workflow.Context) (*iotevents.DeleteDetectorModelOutput, error) {
	var output iotevents.DeleteDetectorModelOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type IoTEventsDeleteInputFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *IoTEventsDeleteInputFuture) Get(ctx workflow.Context) (*iotevents.DeleteInputOutput, error) {
	var output iotevents.DeleteInputOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type IoTEventsDescribeDetectorModelFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *IoTEventsDescribeDetectorModelFuture) Get(ctx workflow.Context) (*iotevents.DescribeDetectorModelOutput, error) {
	var output iotevents.DescribeDetectorModelOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type IoTEventsDescribeInputFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *IoTEventsDescribeInputFuture) Get(ctx workflow.Context) (*iotevents.DescribeInputOutput, error) {
	var output iotevents.DescribeInputOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type IoTEventsDescribeLoggingOptionsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *IoTEventsDescribeLoggingOptionsFuture) Get(ctx workflow.Context) (*iotevents.DescribeLoggingOptionsOutput, error) {
	var output iotevents.DescribeLoggingOptionsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type IoTEventsListDetectorModelVersionsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *IoTEventsListDetectorModelVersionsFuture) Get(ctx workflow.Context) (*iotevents.ListDetectorModelVersionsOutput, error) {
	var output iotevents.ListDetectorModelVersionsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type IoTEventsListDetectorModelsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *IoTEventsListDetectorModelsFuture) Get(ctx workflow.Context) (*iotevents.ListDetectorModelsOutput, error) {
	var output iotevents.ListDetectorModelsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type IoTEventsListInputsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *IoTEventsListInputsFuture) Get(ctx workflow.Context) (*iotevents.ListInputsOutput, error) {
	var output iotevents.ListInputsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type IoTEventsListTagsForResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *IoTEventsListTagsForResourceFuture) Get(ctx workflow.Context) (*iotevents.ListTagsForResourceOutput, error) {
	var output iotevents.ListTagsForResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type IoTEventsPutLoggingOptionsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *IoTEventsPutLoggingOptionsFuture) Get(ctx workflow.Context) (*iotevents.PutLoggingOptionsOutput, error) {
	var output iotevents.PutLoggingOptionsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type IoTEventsTagResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *IoTEventsTagResourceFuture) Get(ctx workflow.Context) (*iotevents.TagResourceOutput, error) {
	var output iotevents.TagResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type IoTEventsUntagResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *IoTEventsUntagResourceFuture) Get(ctx workflow.Context) (*iotevents.UntagResourceOutput, error) {
	var output iotevents.UntagResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type IoTEventsUpdateDetectorModelFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *IoTEventsUpdateDetectorModelFuture) Get(ctx workflow.Context) (*iotevents.UpdateDetectorModelOutput, error) {
	var output iotevents.UpdateDetectorModelOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type IoTEventsUpdateInputFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *IoTEventsUpdateInputFuture) Get(ctx workflow.Context) (*iotevents.UpdateInputOutput, error) {
	var output iotevents.UpdateInputOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateDetectorModel(ctx workflow.Context, input *iotevents.CreateDetectorModelInput) (*iotevents.CreateDetectorModelOutput, error) {
	var output iotevents.CreateDetectorModelOutput
	err := workflow.ExecuteActivity(ctx, "aws-iotevents-CreateDetectorModel", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateDetectorModelAsync(ctx workflow.Context, input *iotevents.CreateDetectorModelInput) *IoTEventsCreateDetectorModelFuture {
	future := workflow.ExecuteActivity(ctx, "aws-iotevents-CreateDetectorModel", input)
	return &IoTEventsCreateDetectorModelFuture{Future: future}
}

func (a *stub) CreateInput(ctx workflow.Context, input *iotevents.CreateInputInput) (*iotevents.CreateInputOutput, error) {
	var output iotevents.CreateInputOutput
	err := workflow.ExecuteActivity(ctx, "aws-iotevents-CreateInput", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateInputAsync(ctx workflow.Context, input *iotevents.CreateInputInput) *IoTEventsCreateInputFuture {
	future := workflow.ExecuteActivity(ctx, "aws-iotevents-CreateInput", input)
	return &IoTEventsCreateInputFuture{Future: future}
}

func (a *stub) DeleteDetectorModel(ctx workflow.Context, input *iotevents.DeleteDetectorModelInput) (*iotevents.DeleteDetectorModelOutput, error) {
	var output iotevents.DeleteDetectorModelOutput
	err := workflow.ExecuteActivity(ctx, "aws-iotevents-DeleteDetectorModel", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteDetectorModelAsync(ctx workflow.Context, input *iotevents.DeleteDetectorModelInput) *IoTEventsDeleteDetectorModelFuture {
	future := workflow.ExecuteActivity(ctx, "aws-iotevents-DeleteDetectorModel", input)
	return &IoTEventsDeleteDetectorModelFuture{Future: future}
}

func (a *stub) DeleteInput(ctx workflow.Context, input *iotevents.DeleteInputInput) (*iotevents.DeleteInputOutput, error) {
	var output iotevents.DeleteInputOutput
	err := workflow.ExecuteActivity(ctx, "aws-iotevents-DeleteInput", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteInputAsync(ctx workflow.Context, input *iotevents.DeleteInputInput) *IoTEventsDeleteInputFuture {
	future := workflow.ExecuteActivity(ctx, "aws-iotevents-DeleteInput", input)
	return &IoTEventsDeleteInputFuture{Future: future}
}

func (a *stub) DescribeDetectorModel(ctx workflow.Context, input *iotevents.DescribeDetectorModelInput) (*iotevents.DescribeDetectorModelOutput, error) {
	var output iotevents.DescribeDetectorModelOutput
	err := workflow.ExecuteActivity(ctx, "aws-iotevents-DescribeDetectorModel", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeDetectorModelAsync(ctx workflow.Context, input *iotevents.DescribeDetectorModelInput) *IoTEventsDescribeDetectorModelFuture {
	future := workflow.ExecuteActivity(ctx, "aws-iotevents-DescribeDetectorModel", input)
	return &IoTEventsDescribeDetectorModelFuture{Future: future}
}

func (a *stub) DescribeInput(ctx workflow.Context, input *iotevents.DescribeInputInput) (*iotevents.DescribeInputOutput, error) {
	var output iotevents.DescribeInputOutput
	err := workflow.ExecuteActivity(ctx, "aws-iotevents-DescribeInput", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeInputAsync(ctx workflow.Context, input *iotevents.DescribeInputInput) *IoTEventsDescribeInputFuture {
	future := workflow.ExecuteActivity(ctx, "aws-iotevents-DescribeInput", input)
	return &IoTEventsDescribeInputFuture{Future: future}
}

func (a *stub) DescribeLoggingOptions(ctx workflow.Context, input *iotevents.DescribeLoggingOptionsInput) (*iotevents.DescribeLoggingOptionsOutput, error) {
	var output iotevents.DescribeLoggingOptionsOutput
	err := workflow.ExecuteActivity(ctx, "aws-iotevents-DescribeLoggingOptions", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeLoggingOptionsAsync(ctx workflow.Context, input *iotevents.DescribeLoggingOptionsInput) *IoTEventsDescribeLoggingOptionsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-iotevents-DescribeLoggingOptions", input)
	return &IoTEventsDescribeLoggingOptionsFuture{Future: future}
}

func (a *stub) ListDetectorModelVersions(ctx workflow.Context, input *iotevents.ListDetectorModelVersionsInput) (*iotevents.ListDetectorModelVersionsOutput, error) {
	var output iotevents.ListDetectorModelVersionsOutput
	err := workflow.ExecuteActivity(ctx, "aws-iotevents-ListDetectorModelVersions", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListDetectorModelVersionsAsync(ctx workflow.Context, input *iotevents.ListDetectorModelVersionsInput) *IoTEventsListDetectorModelVersionsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-iotevents-ListDetectorModelVersions", input)
	return &IoTEventsListDetectorModelVersionsFuture{Future: future}
}

func (a *stub) ListDetectorModels(ctx workflow.Context, input *iotevents.ListDetectorModelsInput) (*iotevents.ListDetectorModelsOutput, error) {
	var output iotevents.ListDetectorModelsOutput
	err := workflow.ExecuteActivity(ctx, "aws-iotevents-ListDetectorModels", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListDetectorModelsAsync(ctx workflow.Context, input *iotevents.ListDetectorModelsInput) *IoTEventsListDetectorModelsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-iotevents-ListDetectorModels", input)
	return &IoTEventsListDetectorModelsFuture{Future: future}
}

func (a *stub) ListInputs(ctx workflow.Context, input *iotevents.ListInputsInput) (*iotevents.ListInputsOutput, error) {
	var output iotevents.ListInputsOutput
	err := workflow.ExecuteActivity(ctx, "aws-iotevents-ListInputs", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListInputsAsync(ctx workflow.Context, input *iotevents.ListInputsInput) *IoTEventsListInputsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-iotevents-ListInputs", input)
	return &IoTEventsListInputsFuture{Future: future}
}

func (a *stub) ListTagsForResource(ctx workflow.Context, input *iotevents.ListTagsForResourceInput) (*iotevents.ListTagsForResourceOutput, error) {
	var output iotevents.ListTagsForResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws-iotevents-ListTagsForResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListTagsForResourceAsync(ctx workflow.Context, input *iotevents.ListTagsForResourceInput) *IoTEventsListTagsForResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws-iotevents-ListTagsForResource", input)
	return &IoTEventsListTagsForResourceFuture{Future: future}
}

func (a *stub) PutLoggingOptions(ctx workflow.Context, input *iotevents.PutLoggingOptionsInput) (*iotevents.PutLoggingOptionsOutput, error) {
	var output iotevents.PutLoggingOptionsOutput
	err := workflow.ExecuteActivity(ctx, "aws-iotevents-PutLoggingOptions", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) PutLoggingOptionsAsync(ctx workflow.Context, input *iotevents.PutLoggingOptionsInput) *IoTEventsPutLoggingOptionsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-iotevents-PutLoggingOptions", input)
	return &IoTEventsPutLoggingOptionsFuture{Future: future}
}

func (a *stub) TagResource(ctx workflow.Context, input *iotevents.TagResourceInput) (*iotevents.TagResourceOutput, error) {
	var output iotevents.TagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws-iotevents-TagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) TagResourceAsync(ctx workflow.Context, input *iotevents.TagResourceInput) *IoTEventsTagResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws-iotevents-TagResource", input)
	return &IoTEventsTagResourceFuture{Future: future}
}

func (a *stub) UntagResource(ctx workflow.Context, input *iotevents.UntagResourceInput) (*iotevents.UntagResourceOutput, error) {
	var output iotevents.UntagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws-iotevents-UntagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UntagResourceAsync(ctx workflow.Context, input *iotevents.UntagResourceInput) *IoTEventsUntagResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws-iotevents-UntagResource", input)
	return &IoTEventsUntagResourceFuture{Future: future}
}

func (a *stub) UpdateDetectorModel(ctx workflow.Context, input *iotevents.UpdateDetectorModelInput) (*iotevents.UpdateDetectorModelOutput, error) {
	var output iotevents.UpdateDetectorModelOutput
	err := workflow.ExecuteActivity(ctx, "aws-iotevents-UpdateDetectorModel", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateDetectorModelAsync(ctx workflow.Context, input *iotevents.UpdateDetectorModelInput) *IoTEventsUpdateDetectorModelFuture {
	future := workflow.ExecuteActivity(ctx, "aws-iotevents-UpdateDetectorModel", input)
	return &IoTEventsUpdateDetectorModelFuture{Future: future}
}

func (a *stub) UpdateInput(ctx workflow.Context, input *iotevents.UpdateInputInput) (*iotevents.UpdateInputOutput, error) {
	var output iotevents.UpdateInputOutput
	err := workflow.ExecuteActivity(ctx, "aws-iotevents-UpdateInput", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateInputAsync(ctx workflow.Context, input *iotevents.UpdateInputInput) *IoTEventsUpdateInputFuture {
	future := workflow.ExecuteActivity(ctx, "aws-iotevents-UpdateInput", input)
	return &IoTEventsUpdateInputFuture{Future: future}
}
