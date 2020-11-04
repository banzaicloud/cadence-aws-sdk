// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.
// Copyright (c) 2020 Banzai Cloud All rights reserved.

package braketstub

import (
	"github.com/aws/aws-sdk-go/service/braket"
	"go.uber.org/cadence/workflow"

	"github.com/banzaicloud/cadence-aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type stub struct{}

type BraketCancelQuantumTaskFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *BraketCancelQuantumTaskFuture) Get(ctx workflow.Context) (*braket.CancelQuantumTaskOutput, error) {
	var output braket.CancelQuantumTaskOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type BraketCreateQuantumTaskFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *BraketCreateQuantumTaskFuture) Get(ctx workflow.Context) (*braket.CreateQuantumTaskOutput, error) {
	var output braket.CreateQuantumTaskOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type BraketGetDeviceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *BraketGetDeviceFuture) Get(ctx workflow.Context) (*braket.GetDeviceOutput, error) {
	var output braket.GetDeviceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type BraketGetQuantumTaskFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *BraketGetQuantumTaskFuture) Get(ctx workflow.Context) (*braket.GetQuantumTaskOutput, error) {
	var output braket.GetQuantumTaskOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type BraketListTagsForResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *BraketListTagsForResourceFuture) Get(ctx workflow.Context) (*braket.ListTagsForResourceOutput, error) {
	var output braket.ListTagsForResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type BraketSearchDevicesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *BraketSearchDevicesFuture) Get(ctx workflow.Context) (*braket.SearchDevicesOutput, error) {
	var output braket.SearchDevicesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type BraketSearchQuantumTasksFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *BraketSearchQuantumTasksFuture) Get(ctx workflow.Context) (*braket.SearchQuantumTasksOutput, error) {
	var output braket.SearchQuantumTasksOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type BraketTagResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *BraketTagResourceFuture) Get(ctx workflow.Context) (*braket.TagResourceOutput, error) {
	var output braket.TagResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type BraketUntagResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *BraketUntagResourceFuture) Get(ctx workflow.Context) (*braket.UntagResourceOutput, error) {
	var output braket.UntagResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *stub) CancelQuantumTask(ctx workflow.Context, input *braket.CancelQuantumTaskInput) (*braket.CancelQuantumTaskOutput, error) {
	var output braket.CancelQuantumTaskOutput
	err := workflow.ExecuteActivity(ctx, "aws-braket-CancelQuantumTask", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CancelQuantumTaskAsync(ctx workflow.Context, input *braket.CancelQuantumTaskInput) *BraketCancelQuantumTaskFuture {
	future := workflow.ExecuteActivity(ctx, "aws-braket-CancelQuantumTask", input)
	return &BraketCancelQuantumTaskFuture{Future: future}
}

func (a *stub) CreateQuantumTask(ctx workflow.Context, input *braket.CreateQuantumTaskInput) (*braket.CreateQuantumTaskOutput, error) {
	var output braket.CreateQuantumTaskOutput
	err := workflow.ExecuteActivity(ctx, "aws-braket-CreateQuantumTask", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateQuantumTaskAsync(ctx workflow.Context, input *braket.CreateQuantumTaskInput) *BraketCreateQuantumTaskFuture {
	future := workflow.ExecuteActivity(ctx, "aws-braket-CreateQuantumTask", input)
	return &BraketCreateQuantumTaskFuture{Future: future}
}

func (a *stub) GetDevice(ctx workflow.Context, input *braket.GetDeviceInput) (*braket.GetDeviceOutput, error) {
	var output braket.GetDeviceOutput
	err := workflow.ExecuteActivity(ctx, "aws-braket-GetDevice", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetDeviceAsync(ctx workflow.Context, input *braket.GetDeviceInput) *BraketGetDeviceFuture {
	future := workflow.ExecuteActivity(ctx, "aws-braket-GetDevice", input)
	return &BraketGetDeviceFuture{Future: future}
}

func (a *stub) GetQuantumTask(ctx workflow.Context, input *braket.GetQuantumTaskInput) (*braket.GetQuantumTaskOutput, error) {
	var output braket.GetQuantumTaskOutput
	err := workflow.ExecuteActivity(ctx, "aws-braket-GetQuantumTask", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetQuantumTaskAsync(ctx workflow.Context, input *braket.GetQuantumTaskInput) *BraketGetQuantumTaskFuture {
	future := workflow.ExecuteActivity(ctx, "aws-braket-GetQuantumTask", input)
	return &BraketGetQuantumTaskFuture{Future: future}
}

func (a *stub) ListTagsForResource(ctx workflow.Context, input *braket.ListTagsForResourceInput) (*braket.ListTagsForResourceOutput, error) {
	var output braket.ListTagsForResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws-braket-ListTagsForResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListTagsForResourceAsync(ctx workflow.Context, input *braket.ListTagsForResourceInput) *BraketListTagsForResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws-braket-ListTagsForResource", input)
	return &BraketListTagsForResourceFuture{Future: future}
}

func (a *stub) SearchDevices(ctx workflow.Context, input *braket.SearchDevicesInput) (*braket.SearchDevicesOutput, error) {
	var output braket.SearchDevicesOutput
	err := workflow.ExecuteActivity(ctx, "aws-braket-SearchDevices", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) SearchDevicesAsync(ctx workflow.Context, input *braket.SearchDevicesInput) *BraketSearchDevicesFuture {
	future := workflow.ExecuteActivity(ctx, "aws-braket-SearchDevices", input)
	return &BraketSearchDevicesFuture{Future: future}
}

func (a *stub) SearchQuantumTasks(ctx workflow.Context, input *braket.SearchQuantumTasksInput) (*braket.SearchQuantumTasksOutput, error) {
	var output braket.SearchQuantumTasksOutput
	err := workflow.ExecuteActivity(ctx, "aws-braket-SearchQuantumTasks", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) SearchQuantumTasksAsync(ctx workflow.Context, input *braket.SearchQuantumTasksInput) *BraketSearchQuantumTasksFuture {
	future := workflow.ExecuteActivity(ctx, "aws-braket-SearchQuantumTasks", input)
	return &BraketSearchQuantumTasksFuture{Future: future}
}

func (a *stub) TagResource(ctx workflow.Context, input *braket.TagResourceInput) (*braket.TagResourceOutput, error) {
	var output braket.TagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws-braket-TagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) TagResourceAsync(ctx workflow.Context, input *braket.TagResourceInput) *BraketTagResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws-braket-TagResource", input)
	return &BraketTagResourceFuture{Future: future}
}

func (a *stub) UntagResource(ctx workflow.Context, input *braket.UntagResourceInput) (*braket.UntagResourceOutput, error) {
	var output braket.UntagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws-braket-UntagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UntagResourceAsync(ctx workflow.Context, input *braket.UntagResourceInput) *BraketUntagResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws-braket-UntagResource", input)
	return &BraketUntagResourceFuture{Future: future}
}
