// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.
// Copyright (c) 2020 Banzai Cloud All rights reserved.

package lakeformationstub

import (
	"github.com/aws/aws-sdk-go/service/lakeformation"
	"go.uber.org/cadence/workflow"

	"github.com/banzaicloud/cadence-aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type stub struct{}

type LakeFormationBatchGrantPermissionsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *LakeFormationBatchGrantPermissionsFuture) Get(ctx workflow.Context) (*lakeformation.BatchGrantPermissionsOutput, error) {
	var output lakeformation.BatchGrantPermissionsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type LakeFormationBatchRevokePermissionsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *LakeFormationBatchRevokePermissionsFuture) Get(ctx workflow.Context) (*lakeformation.BatchRevokePermissionsOutput, error) {
	var output lakeformation.BatchRevokePermissionsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type LakeFormationDeregisterResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *LakeFormationDeregisterResourceFuture) Get(ctx workflow.Context) (*lakeformation.DeregisterResourceOutput, error) {
	var output lakeformation.DeregisterResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type LakeFormationDescribeResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *LakeFormationDescribeResourceFuture) Get(ctx workflow.Context) (*lakeformation.DescribeResourceOutput, error) {
	var output lakeformation.DescribeResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type LakeFormationGetDataLakeSettingsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *LakeFormationGetDataLakeSettingsFuture) Get(ctx workflow.Context) (*lakeformation.GetDataLakeSettingsOutput, error) {
	var output lakeformation.GetDataLakeSettingsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type LakeFormationGetEffectivePermissionsForPathFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *LakeFormationGetEffectivePermissionsForPathFuture) Get(ctx workflow.Context) (*lakeformation.GetEffectivePermissionsForPathOutput, error) {
	var output lakeformation.GetEffectivePermissionsForPathOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type LakeFormationGrantPermissionsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *LakeFormationGrantPermissionsFuture) Get(ctx workflow.Context) (*lakeformation.GrantPermissionsOutput, error) {
	var output lakeformation.GrantPermissionsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type LakeFormationListPermissionsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *LakeFormationListPermissionsFuture) Get(ctx workflow.Context) (*lakeformation.ListPermissionsOutput, error) {
	var output lakeformation.ListPermissionsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type LakeFormationListResourcesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *LakeFormationListResourcesFuture) Get(ctx workflow.Context) (*lakeformation.ListResourcesOutput, error) {
	var output lakeformation.ListResourcesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type LakeFormationPutDataLakeSettingsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *LakeFormationPutDataLakeSettingsFuture) Get(ctx workflow.Context) (*lakeformation.PutDataLakeSettingsOutput, error) {
	var output lakeformation.PutDataLakeSettingsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type LakeFormationRegisterResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *LakeFormationRegisterResourceFuture) Get(ctx workflow.Context) (*lakeformation.RegisterResourceOutput, error) {
	var output lakeformation.RegisterResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type LakeFormationRevokePermissionsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *LakeFormationRevokePermissionsFuture) Get(ctx workflow.Context) (*lakeformation.RevokePermissionsOutput, error) {
	var output lakeformation.RevokePermissionsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type LakeFormationUpdateResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *LakeFormationUpdateResourceFuture) Get(ctx workflow.Context) (*lakeformation.UpdateResourceOutput, error) {
	var output lakeformation.UpdateResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *stub) BatchGrantPermissions(ctx workflow.Context, input *lakeformation.BatchGrantPermissionsInput) (*lakeformation.BatchGrantPermissionsOutput, error) {
	var output lakeformation.BatchGrantPermissionsOutput
	err := workflow.ExecuteActivity(ctx, "aws-lakeformation-BatchGrantPermissions", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) BatchGrantPermissionsAsync(ctx workflow.Context, input *lakeformation.BatchGrantPermissionsInput) *LakeFormationBatchGrantPermissionsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-lakeformation-BatchGrantPermissions", input)
	return &LakeFormationBatchGrantPermissionsFuture{Future: future}
}

func (a *stub) BatchRevokePermissions(ctx workflow.Context, input *lakeformation.BatchRevokePermissionsInput) (*lakeformation.BatchRevokePermissionsOutput, error) {
	var output lakeformation.BatchRevokePermissionsOutput
	err := workflow.ExecuteActivity(ctx, "aws-lakeformation-BatchRevokePermissions", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) BatchRevokePermissionsAsync(ctx workflow.Context, input *lakeformation.BatchRevokePermissionsInput) *LakeFormationBatchRevokePermissionsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-lakeformation-BatchRevokePermissions", input)
	return &LakeFormationBatchRevokePermissionsFuture{Future: future}
}

func (a *stub) DeregisterResource(ctx workflow.Context, input *lakeformation.DeregisterResourceInput) (*lakeformation.DeregisterResourceOutput, error) {
	var output lakeformation.DeregisterResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws-lakeformation-DeregisterResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeregisterResourceAsync(ctx workflow.Context, input *lakeformation.DeregisterResourceInput) *LakeFormationDeregisterResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws-lakeformation-DeregisterResource", input)
	return &LakeFormationDeregisterResourceFuture{Future: future}
}

func (a *stub) DescribeResource(ctx workflow.Context, input *lakeformation.DescribeResourceInput) (*lakeformation.DescribeResourceOutput, error) {
	var output lakeformation.DescribeResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws-lakeformation-DescribeResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeResourceAsync(ctx workflow.Context, input *lakeformation.DescribeResourceInput) *LakeFormationDescribeResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws-lakeformation-DescribeResource", input)
	return &LakeFormationDescribeResourceFuture{Future: future}
}

func (a *stub) GetDataLakeSettings(ctx workflow.Context, input *lakeformation.GetDataLakeSettingsInput) (*lakeformation.GetDataLakeSettingsOutput, error) {
	var output lakeformation.GetDataLakeSettingsOutput
	err := workflow.ExecuteActivity(ctx, "aws-lakeformation-GetDataLakeSettings", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetDataLakeSettingsAsync(ctx workflow.Context, input *lakeformation.GetDataLakeSettingsInput) *LakeFormationGetDataLakeSettingsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-lakeformation-GetDataLakeSettings", input)
	return &LakeFormationGetDataLakeSettingsFuture{Future: future}
}

func (a *stub) GetEffectivePermissionsForPath(ctx workflow.Context, input *lakeformation.GetEffectivePermissionsForPathInput) (*lakeformation.GetEffectivePermissionsForPathOutput, error) {
	var output lakeformation.GetEffectivePermissionsForPathOutput
	err := workflow.ExecuteActivity(ctx, "aws-lakeformation-GetEffectivePermissionsForPath", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetEffectivePermissionsForPathAsync(ctx workflow.Context, input *lakeformation.GetEffectivePermissionsForPathInput) *LakeFormationGetEffectivePermissionsForPathFuture {
	future := workflow.ExecuteActivity(ctx, "aws-lakeformation-GetEffectivePermissionsForPath", input)
	return &LakeFormationGetEffectivePermissionsForPathFuture{Future: future}
}

func (a *stub) GrantPermissions(ctx workflow.Context, input *lakeformation.GrantPermissionsInput) (*lakeformation.GrantPermissionsOutput, error) {
	var output lakeformation.GrantPermissionsOutput
	err := workflow.ExecuteActivity(ctx, "aws-lakeformation-GrantPermissions", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GrantPermissionsAsync(ctx workflow.Context, input *lakeformation.GrantPermissionsInput) *LakeFormationGrantPermissionsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-lakeformation-GrantPermissions", input)
	return &LakeFormationGrantPermissionsFuture{Future: future}
}

func (a *stub) ListPermissions(ctx workflow.Context, input *lakeformation.ListPermissionsInput) (*lakeformation.ListPermissionsOutput, error) {
	var output lakeformation.ListPermissionsOutput
	err := workflow.ExecuteActivity(ctx, "aws-lakeformation-ListPermissions", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListPermissionsAsync(ctx workflow.Context, input *lakeformation.ListPermissionsInput) *LakeFormationListPermissionsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-lakeformation-ListPermissions", input)
	return &LakeFormationListPermissionsFuture{Future: future}
}

func (a *stub) ListResources(ctx workflow.Context, input *lakeformation.ListResourcesInput) (*lakeformation.ListResourcesOutput, error) {
	var output lakeformation.ListResourcesOutput
	err := workflow.ExecuteActivity(ctx, "aws-lakeformation-ListResources", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListResourcesAsync(ctx workflow.Context, input *lakeformation.ListResourcesInput) *LakeFormationListResourcesFuture {
	future := workflow.ExecuteActivity(ctx, "aws-lakeformation-ListResources", input)
	return &LakeFormationListResourcesFuture{Future: future}
}

func (a *stub) PutDataLakeSettings(ctx workflow.Context, input *lakeformation.PutDataLakeSettingsInput) (*lakeformation.PutDataLakeSettingsOutput, error) {
	var output lakeformation.PutDataLakeSettingsOutput
	err := workflow.ExecuteActivity(ctx, "aws-lakeformation-PutDataLakeSettings", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) PutDataLakeSettingsAsync(ctx workflow.Context, input *lakeformation.PutDataLakeSettingsInput) *LakeFormationPutDataLakeSettingsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-lakeformation-PutDataLakeSettings", input)
	return &LakeFormationPutDataLakeSettingsFuture{Future: future}
}

func (a *stub) RegisterResource(ctx workflow.Context, input *lakeformation.RegisterResourceInput) (*lakeformation.RegisterResourceOutput, error) {
	var output lakeformation.RegisterResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws-lakeformation-RegisterResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) RegisterResourceAsync(ctx workflow.Context, input *lakeformation.RegisterResourceInput) *LakeFormationRegisterResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws-lakeformation-RegisterResource", input)
	return &LakeFormationRegisterResourceFuture{Future: future}
}

func (a *stub) RevokePermissions(ctx workflow.Context, input *lakeformation.RevokePermissionsInput) (*lakeformation.RevokePermissionsOutput, error) {
	var output lakeformation.RevokePermissionsOutput
	err := workflow.ExecuteActivity(ctx, "aws-lakeformation-RevokePermissions", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) RevokePermissionsAsync(ctx workflow.Context, input *lakeformation.RevokePermissionsInput) *LakeFormationRevokePermissionsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-lakeformation-RevokePermissions", input)
	return &LakeFormationRevokePermissionsFuture{Future: future}
}

func (a *stub) UpdateResource(ctx workflow.Context, input *lakeformation.UpdateResourceInput) (*lakeformation.UpdateResourceOutput, error) {
	var output lakeformation.UpdateResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws-lakeformation-UpdateResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateResourceAsync(ctx workflow.Context, input *lakeformation.UpdateResourceInput) *LakeFormationUpdateResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws-lakeformation-UpdateResource", input)
	return &LakeFormationUpdateResourceFuture{Future: future}
}
