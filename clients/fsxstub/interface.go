// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7

package fsxstub

import (
	"github.com/aws/aws-sdk-go/service/fsx"
	"go.uber.org/cadence/workflow"

	"github.com/banzaicloud/cadence-aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	AssociateFileSystemAliases(ctx workflow.Context, input *fsx.AssociateFileSystemAliasesInput) (*fsx.AssociateFileSystemAliasesOutput, error)
	AssociateFileSystemAliasesAsync(ctx workflow.Context, input *fsx.AssociateFileSystemAliasesInput) *AssociateFileSystemAliasesFuture

	CancelDataRepositoryTask(ctx workflow.Context, input *fsx.CancelDataRepositoryTaskInput) (*fsx.CancelDataRepositoryTaskOutput, error)
	CancelDataRepositoryTaskAsync(ctx workflow.Context, input *fsx.CancelDataRepositoryTaskInput) *CancelDataRepositoryTaskFuture

	CreateBackup(ctx workflow.Context, input *fsx.CreateBackupInput) (*fsx.CreateBackupOutput, error)
	CreateBackupAsync(ctx workflow.Context, input *fsx.CreateBackupInput) *CreateBackupFuture

	CreateDataRepositoryTask(ctx workflow.Context, input *fsx.CreateDataRepositoryTaskInput) (*fsx.CreateDataRepositoryTaskOutput, error)
	CreateDataRepositoryTaskAsync(ctx workflow.Context, input *fsx.CreateDataRepositoryTaskInput) *CreateDataRepositoryTaskFuture

	CreateFileSystem(ctx workflow.Context, input *fsx.CreateFileSystemInput) (*fsx.CreateFileSystemOutput, error)
	CreateFileSystemAsync(ctx workflow.Context, input *fsx.CreateFileSystemInput) *CreateFileSystemFuture

	CreateFileSystemFromBackup(ctx workflow.Context, input *fsx.CreateFileSystemFromBackupInput) (*fsx.CreateFileSystemFromBackupOutput, error)
	CreateFileSystemFromBackupAsync(ctx workflow.Context, input *fsx.CreateFileSystemFromBackupInput) *CreateFileSystemFromBackupFuture

	DeleteBackup(ctx workflow.Context, input *fsx.DeleteBackupInput) (*fsx.DeleteBackupOutput, error)
	DeleteBackupAsync(ctx workflow.Context, input *fsx.DeleteBackupInput) *DeleteBackupFuture

	DeleteFileSystem(ctx workflow.Context, input *fsx.DeleteFileSystemInput) (*fsx.DeleteFileSystemOutput, error)
	DeleteFileSystemAsync(ctx workflow.Context, input *fsx.DeleteFileSystemInput) *DeleteFileSystemFuture

	DescribeBackups(ctx workflow.Context, input *fsx.DescribeBackupsInput) (*fsx.DescribeBackupsOutput, error)
	DescribeBackupsAsync(ctx workflow.Context, input *fsx.DescribeBackupsInput) *DescribeBackupsFuture

	DescribeDataRepositoryTasks(ctx workflow.Context, input *fsx.DescribeDataRepositoryTasksInput) (*fsx.DescribeDataRepositoryTasksOutput, error)
	DescribeDataRepositoryTasksAsync(ctx workflow.Context, input *fsx.DescribeDataRepositoryTasksInput) *DescribeDataRepositoryTasksFuture

	DescribeFileSystemAliases(ctx workflow.Context, input *fsx.DescribeFileSystemAliasesInput) (*fsx.DescribeFileSystemAliasesOutput, error)
	DescribeFileSystemAliasesAsync(ctx workflow.Context, input *fsx.DescribeFileSystemAliasesInput) *DescribeFileSystemAliasesFuture

	DescribeFileSystems(ctx workflow.Context, input *fsx.DescribeFileSystemsInput) (*fsx.DescribeFileSystemsOutput, error)
	DescribeFileSystemsAsync(ctx workflow.Context, input *fsx.DescribeFileSystemsInput) *DescribeFileSystemsFuture

	DisassociateFileSystemAliases(ctx workflow.Context, input *fsx.DisassociateFileSystemAliasesInput) (*fsx.DisassociateFileSystemAliasesOutput, error)
	DisassociateFileSystemAliasesAsync(ctx workflow.Context, input *fsx.DisassociateFileSystemAliasesInput) *DisassociateFileSystemAliasesFuture

	ListTagsForResource(ctx workflow.Context, input *fsx.ListTagsForResourceInput) (*fsx.ListTagsForResourceOutput, error)
	ListTagsForResourceAsync(ctx workflow.Context, input *fsx.ListTagsForResourceInput) *ListTagsForResourceFuture

	TagResource(ctx workflow.Context, input *fsx.TagResourceInput) (*fsx.TagResourceOutput, error)
	TagResourceAsync(ctx workflow.Context, input *fsx.TagResourceInput) *TagResourceFuture

	UntagResource(ctx workflow.Context, input *fsx.UntagResourceInput) (*fsx.UntagResourceOutput, error)
	UntagResourceAsync(ctx workflow.Context, input *fsx.UntagResourceInput) *UntagResourceFuture

	UpdateFileSystem(ctx workflow.Context, input *fsx.UpdateFileSystemInput) (*fsx.UpdateFileSystemOutput, error)
	UpdateFileSystemAsync(ctx workflow.Context, input *fsx.UpdateFileSystemInput) *UpdateFileSystemFuture
}

func NewClient() Client {
	return &stub{}
}
