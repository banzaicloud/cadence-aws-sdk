// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7

package efsstub

import (
	"github.com/aws/aws-sdk-go/service/efs"
	"go.uber.org/cadence/workflow"

	"github.com/banzaicloud/cadence-aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type stub struct{}

type EFSCreateAccessPointFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *EFSCreateAccessPointFuture) Get(ctx workflow.Context) (*efs.CreateAccessPointOutput, error) {
	var output efs.CreateAccessPointOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type EFSCreateFileSystemFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *EFSCreateFileSystemFuture) Get(ctx workflow.Context) (*efs.FileSystemDescription, error) {
	var output efs.FileSystemDescription
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type EFSCreateMountTargetFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *EFSCreateMountTargetFuture) Get(ctx workflow.Context) (*efs.MountTargetDescription, error) {
	var output efs.MountTargetDescription
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type EFSCreateTagsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *EFSCreateTagsFuture) Get(ctx workflow.Context) (*efs.CreateTagsOutput, error) {
	var output efs.CreateTagsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type EFSDeleteAccessPointFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *EFSDeleteAccessPointFuture) Get(ctx workflow.Context) (*efs.DeleteAccessPointOutput, error) {
	var output efs.DeleteAccessPointOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type EFSDeleteFileSystemFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *EFSDeleteFileSystemFuture) Get(ctx workflow.Context) (*efs.DeleteFileSystemOutput, error) {
	var output efs.DeleteFileSystemOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type EFSDeleteFileSystemPolicyFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *EFSDeleteFileSystemPolicyFuture) Get(ctx workflow.Context) (*efs.DeleteFileSystemPolicyOutput, error) {
	var output efs.DeleteFileSystemPolicyOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type EFSDeleteMountTargetFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *EFSDeleteMountTargetFuture) Get(ctx workflow.Context) (*efs.DeleteMountTargetOutput, error) {
	var output efs.DeleteMountTargetOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type EFSDeleteTagsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *EFSDeleteTagsFuture) Get(ctx workflow.Context) (*efs.DeleteTagsOutput, error) {
	var output efs.DeleteTagsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type EFSDescribeAccessPointsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *EFSDescribeAccessPointsFuture) Get(ctx workflow.Context) (*efs.DescribeAccessPointsOutput, error) {
	var output efs.DescribeAccessPointsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type EFSDescribeBackupPolicyFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *EFSDescribeBackupPolicyFuture) Get(ctx workflow.Context) (*efs.DescribeBackupPolicyOutput, error) {
	var output efs.DescribeBackupPolicyOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type EFSDescribeFileSystemPolicyFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *EFSDescribeFileSystemPolicyFuture) Get(ctx workflow.Context) (*efs.DescribeFileSystemPolicyOutput, error) {
	var output efs.DescribeFileSystemPolicyOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type EFSDescribeFileSystemsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *EFSDescribeFileSystemsFuture) Get(ctx workflow.Context) (*efs.DescribeFileSystemsOutput, error) {
	var output efs.DescribeFileSystemsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type EFSDescribeLifecycleConfigurationFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *EFSDescribeLifecycleConfigurationFuture) Get(ctx workflow.Context) (*efs.DescribeLifecycleConfigurationOutput, error) {
	var output efs.DescribeLifecycleConfigurationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type EFSDescribeMountTargetSecurityGroupsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *EFSDescribeMountTargetSecurityGroupsFuture) Get(ctx workflow.Context) (*efs.DescribeMountTargetSecurityGroupsOutput, error) {
	var output efs.DescribeMountTargetSecurityGroupsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type EFSDescribeMountTargetsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *EFSDescribeMountTargetsFuture) Get(ctx workflow.Context) (*efs.DescribeMountTargetsOutput, error) {
	var output efs.DescribeMountTargetsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type EFSDescribeTagsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *EFSDescribeTagsFuture) Get(ctx workflow.Context) (*efs.DescribeTagsOutput, error) {
	var output efs.DescribeTagsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type EFSListTagsForResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *EFSListTagsForResourceFuture) Get(ctx workflow.Context) (*efs.ListTagsForResourceOutput, error) {
	var output efs.ListTagsForResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type EFSModifyMountTargetSecurityGroupsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *EFSModifyMountTargetSecurityGroupsFuture) Get(ctx workflow.Context) (*efs.ModifyMountTargetSecurityGroupsOutput, error) {
	var output efs.ModifyMountTargetSecurityGroupsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type EFSPutBackupPolicyFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *EFSPutBackupPolicyFuture) Get(ctx workflow.Context) (*efs.PutBackupPolicyOutput, error) {
	var output efs.PutBackupPolicyOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type EFSPutFileSystemPolicyFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *EFSPutFileSystemPolicyFuture) Get(ctx workflow.Context) (*efs.PutFileSystemPolicyOutput, error) {
	var output efs.PutFileSystemPolicyOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type EFSPutLifecycleConfigurationFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *EFSPutLifecycleConfigurationFuture) Get(ctx workflow.Context) (*efs.PutLifecycleConfigurationOutput, error) {
	var output efs.PutLifecycleConfigurationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type EFSTagResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *EFSTagResourceFuture) Get(ctx workflow.Context) (*efs.TagResourceOutput, error) {
	var output efs.TagResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type EFSUntagResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *EFSUntagResourceFuture) Get(ctx workflow.Context) (*efs.UntagResourceOutput, error) {
	var output efs.UntagResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type EFSUpdateFileSystemFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *EFSUpdateFileSystemFuture) Get(ctx workflow.Context) (*efs.UpdateFileSystemOutput, error) {
	var output efs.UpdateFileSystemOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateAccessPoint(ctx workflow.Context, input *efs.CreateAccessPointInput) (*efs.CreateAccessPointOutput, error) {
	var output efs.CreateAccessPointOutput
	err := workflow.ExecuteActivity(ctx, "aws-efs-CreateAccessPoint", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateAccessPointAsync(ctx workflow.Context, input *efs.CreateAccessPointInput) *EFSCreateAccessPointFuture {
	future := workflow.ExecuteActivity(ctx, "aws-efs-CreateAccessPoint", input)
	return &EFSCreateAccessPointFuture{Future: future}
}

func (a *stub) CreateFileSystem(ctx workflow.Context, input *efs.CreateFileSystemInput) (*efs.FileSystemDescription, error) {
	var output efs.FileSystemDescription
	err := workflow.ExecuteActivity(ctx, "aws-efs-CreateFileSystem", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateFileSystemAsync(ctx workflow.Context, input *efs.CreateFileSystemInput) *EFSCreateFileSystemFuture {
	future := workflow.ExecuteActivity(ctx, "aws-efs-CreateFileSystem", input)
	return &EFSCreateFileSystemFuture{Future: future}
}

func (a *stub) CreateMountTarget(ctx workflow.Context, input *efs.CreateMountTargetInput) (*efs.MountTargetDescription, error) {
	var output efs.MountTargetDescription
	err := workflow.ExecuteActivity(ctx, "aws-efs-CreateMountTarget", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateMountTargetAsync(ctx workflow.Context, input *efs.CreateMountTargetInput) *EFSCreateMountTargetFuture {
	future := workflow.ExecuteActivity(ctx, "aws-efs-CreateMountTarget", input)
	return &EFSCreateMountTargetFuture{Future: future}
}

func (a *stub) CreateTags(ctx workflow.Context, input *efs.CreateTagsInput) (*efs.CreateTagsOutput, error) {
	var output efs.CreateTagsOutput
	err := workflow.ExecuteActivity(ctx, "aws-efs-CreateTags", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateTagsAsync(ctx workflow.Context, input *efs.CreateTagsInput) *EFSCreateTagsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-efs-CreateTags", input)
	return &EFSCreateTagsFuture{Future: future}
}

func (a *stub) DeleteAccessPoint(ctx workflow.Context, input *efs.DeleteAccessPointInput) (*efs.DeleteAccessPointOutput, error) {
	var output efs.DeleteAccessPointOutput
	err := workflow.ExecuteActivity(ctx, "aws-efs-DeleteAccessPoint", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteAccessPointAsync(ctx workflow.Context, input *efs.DeleteAccessPointInput) *EFSDeleteAccessPointFuture {
	future := workflow.ExecuteActivity(ctx, "aws-efs-DeleteAccessPoint", input)
	return &EFSDeleteAccessPointFuture{Future: future}
}

func (a *stub) DeleteFileSystem(ctx workflow.Context, input *efs.DeleteFileSystemInput) (*efs.DeleteFileSystemOutput, error) {
	var output efs.DeleteFileSystemOutput
	err := workflow.ExecuteActivity(ctx, "aws-efs-DeleteFileSystem", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteFileSystemAsync(ctx workflow.Context, input *efs.DeleteFileSystemInput) *EFSDeleteFileSystemFuture {
	future := workflow.ExecuteActivity(ctx, "aws-efs-DeleteFileSystem", input)
	return &EFSDeleteFileSystemFuture{Future: future}
}

func (a *stub) DeleteFileSystemPolicy(ctx workflow.Context, input *efs.DeleteFileSystemPolicyInput) (*efs.DeleteFileSystemPolicyOutput, error) {
	var output efs.DeleteFileSystemPolicyOutput
	err := workflow.ExecuteActivity(ctx, "aws-efs-DeleteFileSystemPolicy", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteFileSystemPolicyAsync(ctx workflow.Context, input *efs.DeleteFileSystemPolicyInput) *EFSDeleteFileSystemPolicyFuture {
	future := workflow.ExecuteActivity(ctx, "aws-efs-DeleteFileSystemPolicy", input)
	return &EFSDeleteFileSystemPolicyFuture{Future: future}
}

func (a *stub) DeleteMountTarget(ctx workflow.Context, input *efs.DeleteMountTargetInput) (*efs.DeleteMountTargetOutput, error) {
	var output efs.DeleteMountTargetOutput
	err := workflow.ExecuteActivity(ctx, "aws-efs-DeleteMountTarget", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteMountTargetAsync(ctx workflow.Context, input *efs.DeleteMountTargetInput) *EFSDeleteMountTargetFuture {
	future := workflow.ExecuteActivity(ctx, "aws-efs-DeleteMountTarget", input)
	return &EFSDeleteMountTargetFuture{Future: future}
}

func (a *stub) DeleteTags(ctx workflow.Context, input *efs.DeleteTagsInput) (*efs.DeleteTagsOutput, error) {
	var output efs.DeleteTagsOutput
	err := workflow.ExecuteActivity(ctx, "aws-efs-DeleteTags", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteTagsAsync(ctx workflow.Context, input *efs.DeleteTagsInput) *EFSDeleteTagsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-efs-DeleteTags", input)
	return &EFSDeleteTagsFuture{Future: future}
}

func (a *stub) DescribeAccessPoints(ctx workflow.Context, input *efs.DescribeAccessPointsInput) (*efs.DescribeAccessPointsOutput, error) {
	var output efs.DescribeAccessPointsOutput
	err := workflow.ExecuteActivity(ctx, "aws-efs-DescribeAccessPoints", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeAccessPointsAsync(ctx workflow.Context, input *efs.DescribeAccessPointsInput) *EFSDescribeAccessPointsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-efs-DescribeAccessPoints", input)
	return &EFSDescribeAccessPointsFuture{Future: future}
}

func (a *stub) DescribeBackupPolicy(ctx workflow.Context, input *efs.DescribeBackupPolicyInput) (*efs.DescribeBackupPolicyOutput, error) {
	var output efs.DescribeBackupPolicyOutput
	err := workflow.ExecuteActivity(ctx, "aws-efs-DescribeBackupPolicy", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeBackupPolicyAsync(ctx workflow.Context, input *efs.DescribeBackupPolicyInput) *EFSDescribeBackupPolicyFuture {
	future := workflow.ExecuteActivity(ctx, "aws-efs-DescribeBackupPolicy", input)
	return &EFSDescribeBackupPolicyFuture{Future: future}
}

func (a *stub) DescribeFileSystemPolicy(ctx workflow.Context, input *efs.DescribeFileSystemPolicyInput) (*efs.DescribeFileSystemPolicyOutput, error) {
	var output efs.DescribeFileSystemPolicyOutput
	err := workflow.ExecuteActivity(ctx, "aws-efs-DescribeFileSystemPolicy", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeFileSystemPolicyAsync(ctx workflow.Context, input *efs.DescribeFileSystemPolicyInput) *EFSDescribeFileSystemPolicyFuture {
	future := workflow.ExecuteActivity(ctx, "aws-efs-DescribeFileSystemPolicy", input)
	return &EFSDescribeFileSystemPolicyFuture{Future: future}
}

func (a *stub) DescribeFileSystems(ctx workflow.Context, input *efs.DescribeFileSystemsInput) (*efs.DescribeFileSystemsOutput, error) {
	var output efs.DescribeFileSystemsOutput
	err := workflow.ExecuteActivity(ctx, "aws-efs-DescribeFileSystems", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeFileSystemsAsync(ctx workflow.Context, input *efs.DescribeFileSystemsInput) *EFSDescribeFileSystemsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-efs-DescribeFileSystems", input)
	return &EFSDescribeFileSystemsFuture{Future: future}
}

func (a *stub) DescribeLifecycleConfiguration(ctx workflow.Context, input *efs.DescribeLifecycleConfigurationInput) (*efs.DescribeLifecycleConfigurationOutput, error) {
	var output efs.DescribeLifecycleConfigurationOutput
	err := workflow.ExecuteActivity(ctx, "aws-efs-DescribeLifecycleConfiguration", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeLifecycleConfigurationAsync(ctx workflow.Context, input *efs.DescribeLifecycleConfigurationInput) *EFSDescribeLifecycleConfigurationFuture {
	future := workflow.ExecuteActivity(ctx, "aws-efs-DescribeLifecycleConfiguration", input)
	return &EFSDescribeLifecycleConfigurationFuture{Future: future}
}

func (a *stub) DescribeMountTargetSecurityGroups(ctx workflow.Context, input *efs.DescribeMountTargetSecurityGroupsInput) (*efs.DescribeMountTargetSecurityGroupsOutput, error) {
	var output efs.DescribeMountTargetSecurityGroupsOutput
	err := workflow.ExecuteActivity(ctx, "aws-efs-DescribeMountTargetSecurityGroups", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeMountTargetSecurityGroupsAsync(ctx workflow.Context, input *efs.DescribeMountTargetSecurityGroupsInput) *EFSDescribeMountTargetSecurityGroupsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-efs-DescribeMountTargetSecurityGroups", input)
	return &EFSDescribeMountTargetSecurityGroupsFuture{Future: future}
}

func (a *stub) DescribeMountTargets(ctx workflow.Context, input *efs.DescribeMountTargetsInput) (*efs.DescribeMountTargetsOutput, error) {
	var output efs.DescribeMountTargetsOutput
	err := workflow.ExecuteActivity(ctx, "aws-efs-DescribeMountTargets", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeMountTargetsAsync(ctx workflow.Context, input *efs.DescribeMountTargetsInput) *EFSDescribeMountTargetsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-efs-DescribeMountTargets", input)
	return &EFSDescribeMountTargetsFuture{Future: future}
}

func (a *stub) DescribeTags(ctx workflow.Context, input *efs.DescribeTagsInput) (*efs.DescribeTagsOutput, error) {
	var output efs.DescribeTagsOutput
	err := workflow.ExecuteActivity(ctx, "aws-efs-DescribeTags", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeTagsAsync(ctx workflow.Context, input *efs.DescribeTagsInput) *EFSDescribeTagsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-efs-DescribeTags", input)
	return &EFSDescribeTagsFuture{Future: future}
}

func (a *stub) ListTagsForResource(ctx workflow.Context, input *efs.ListTagsForResourceInput) (*efs.ListTagsForResourceOutput, error) {
	var output efs.ListTagsForResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws-efs-ListTagsForResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListTagsForResourceAsync(ctx workflow.Context, input *efs.ListTagsForResourceInput) *EFSListTagsForResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws-efs-ListTagsForResource", input)
	return &EFSListTagsForResourceFuture{Future: future}
}

func (a *stub) ModifyMountTargetSecurityGroups(ctx workflow.Context, input *efs.ModifyMountTargetSecurityGroupsInput) (*efs.ModifyMountTargetSecurityGroupsOutput, error) {
	var output efs.ModifyMountTargetSecurityGroupsOutput
	err := workflow.ExecuteActivity(ctx, "aws-efs-ModifyMountTargetSecurityGroups", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ModifyMountTargetSecurityGroupsAsync(ctx workflow.Context, input *efs.ModifyMountTargetSecurityGroupsInput) *EFSModifyMountTargetSecurityGroupsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-efs-ModifyMountTargetSecurityGroups", input)
	return &EFSModifyMountTargetSecurityGroupsFuture{Future: future}
}

func (a *stub) PutBackupPolicy(ctx workflow.Context, input *efs.PutBackupPolicyInput) (*efs.PutBackupPolicyOutput, error) {
	var output efs.PutBackupPolicyOutput
	err := workflow.ExecuteActivity(ctx, "aws-efs-PutBackupPolicy", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) PutBackupPolicyAsync(ctx workflow.Context, input *efs.PutBackupPolicyInput) *EFSPutBackupPolicyFuture {
	future := workflow.ExecuteActivity(ctx, "aws-efs-PutBackupPolicy", input)
	return &EFSPutBackupPolicyFuture{Future: future}
}

func (a *stub) PutFileSystemPolicy(ctx workflow.Context, input *efs.PutFileSystemPolicyInput) (*efs.PutFileSystemPolicyOutput, error) {
	var output efs.PutFileSystemPolicyOutput
	err := workflow.ExecuteActivity(ctx, "aws-efs-PutFileSystemPolicy", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) PutFileSystemPolicyAsync(ctx workflow.Context, input *efs.PutFileSystemPolicyInput) *EFSPutFileSystemPolicyFuture {
	future := workflow.ExecuteActivity(ctx, "aws-efs-PutFileSystemPolicy", input)
	return &EFSPutFileSystemPolicyFuture{Future: future}
}

func (a *stub) PutLifecycleConfiguration(ctx workflow.Context, input *efs.PutLifecycleConfigurationInput) (*efs.PutLifecycleConfigurationOutput, error) {
	var output efs.PutLifecycleConfigurationOutput
	err := workflow.ExecuteActivity(ctx, "aws-efs-PutLifecycleConfiguration", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) PutLifecycleConfigurationAsync(ctx workflow.Context, input *efs.PutLifecycleConfigurationInput) *EFSPutLifecycleConfigurationFuture {
	future := workflow.ExecuteActivity(ctx, "aws-efs-PutLifecycleConfiguration", input)
	return &EFSPutLifecycleConfigurationFuture{Future: future}
}

func (a *stub) TagResource(ctx workflow.Context, input *efs.TagResourceInput) (*efs.TagResourceOutput, error) {
	var output efs.TagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws-efs-TagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) TagResourceAsync(ctx workflow.Context, input *efs.TagResourceInput) *EFSTagResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws-efs-TagResource", input)
	return &EFSTagResourceFuture{Future: future}
}

func (a *stub) UntagResource(ctx workflow.Context, input *efs.UntagResourceInput) (*efs.UntagResourceOutput, error) {
	var output efs.UntagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws-efs-UntagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UntagResourceAsync(ctx workflow.Context, input *efs.UntagResourceInput) *EFSUntagResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws-efs-UntagResource", input)
	return &EFSUntagResourceFuture{Future: future}
}

func (a *stub) UpdateFileSystem(ctx workflow.Context, input *efs.UpdateFileSystemInput) (*efs.UpdateFileSystemOutput, error) {
	var output efs.UpdateFileSystemOutput
	err := workflow.ExecuteActivity(ctx, "aws-efs-UpdateFileSystem", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateFileSystemAsync(ctx workflow.Context, input *efs.UpdateFileSystemInput) *EFSUpdateFileSystemFuture {
	future := workflow.ExecuteActivity(ctx, "aws-efs-UpdateFileSystem", input)
	return &EFSUpdateFileSystemFuture{Future: future}
}
