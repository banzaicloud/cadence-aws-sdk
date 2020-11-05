// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7

package mediaconvertstub

import (
	"github.com/aws/aws-sdk-go/service/mediaconvert"
	"go.uber.org/cadence/workflow"

	"github.com/banzaicloud/cadence-aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	AssociateCertificate(ctx workflow.Context, input *mediaconvert.AssociateCertificateInput) (*mediaconvert.AssociateCertificateOutput, error)
	AssociateCertificateAsync(ctx workflow.Context, input *mediaconvert.AssociateCertificateInput) *MediaConvertAssociateCertificateFuture

	CancelJob(ctx workflow.Context, input *mediaconvert.CancelJobInput) (*mediaconvert.CancelJobOutput, error)
	CancelJobAsync(ctx workflow.Context, input *mediaconvert.CancelJobInput) *MediaConvertCancelJobFuture

	CreateJob(ctx workflow.Context, input *mediaconvert.CreateJobInput) (*mediaconvert.CreateJobOutput, error)
	CreateJobAsync(ctx workflow.Context, input *mediaconvert.CreateJobInput) *MediaConvertCreateJobFuture

	CreateJobTemplate(ctx workflow.Context, input *mediaconvert.CreateJobTemplateInput) (*mediaconvert.CreateJobTemplateOutput, error)
	CreateJobTemplateAsync(ctx workflow.Context, input *mediaconvert.CreateJobTemplateInput) *MediaConvertCreateJobTemplateFuture

	CreatePreset(ctx workflow.Context, input *mediaconvert.CreatePresetInput) (*mediaconvert.CreatePresetOutput, error)
	CreatePresetAsync(ctx workflow.Context, input *mediaconvert.CreatePresetInput) *MediaConvertCreatePresetFuture

	CreateQueue(ctx workflow.Context, input *mediaconvert.CreateQueueInput) (*mediaconvert.CreateQueueOutput, error)
	CreateQueueAsync(ctx workflow.Context, input *mediaconvert.CreateQueueInput) *MediaConvertCreateQueueFuture

	DeleteJobTemplate(ctx workflow.Context, input *mediaconvert.DeleteJobTemplateInput) (*mediaconvert.DeleteJobTemplateOutput, error)
	DeleteJobTemplateAsync(ctx workflow.Context, input *mediaconvert.DeleteJobTemplateInput) *MediaConvertDeleteJobTemplateFuture

	DeletePreset(ctx workflow.Context, input *mediaconvert.DeletePresetInput) (*mediaconvert.DeletePresetOutput, error)
	DeletePresetAsync(ctx workflow.Context, input *mediaconvert.DeletePresetInput) *MediaConvertDeletePresetFuture

	DeleteQueue(ctx workflow.Context, input *mediaconvert.DeleteQueueInput) (*mediaconvert.DeleteQueueOutput, error)
	DeleteQueueAsync(ctx workflow.Context, input *mediaconvert.DeleteQueueInput) *MediaConvertDeleteQueueFuture

	DescribeEndpoints(ctx workflow.Context, input *mediaconvert.DescribeEndpointsInput) (*mediaconvert.DescribeEndpointsOutput, error)
	DescribeEndpointsAsync(ctx workflow.Context, input *mediaconvert.DescribeEndpointsInput) *MediaConvertDescribeEndpointsFuture

	DisassociateCertificate(ctx workflow.Context, input *mediaconvert.DisassociateCertificateInput) (*mediaconvert.DisassociateCertificateOutput, error)
	DisassociateCertificateAsync(ctx workflow.Context, input *mediaconvert.DisassociateCertificateInput) *MediaConvertDisassociateCertificateFuture

	GetJob(ctx workflow.Context, input *mediaconvert.GetJobInput) (*mediaconvert.GetJobOutput, error)
	GetJobAsync(ctx workflow.Context, input *mediaconvert.GetJobInput) *MediaConvertGetJobFuture

	GetJobTemplate(ctx workflow.Context, input *mediaconvert.GetJobTemplateInput) (*mediaconvert.GetJobTemplateOutput, error)
	GetJobTemplateAsync(ctx workflow.Context, input *mediaconvert.GetJobTemplateInput) *MediaConvertGetJobTemplateFuture

	GetPreset(ctx workflow.Context, input *mediaconvert.GetPresetInput) (*mediaconvert.GetPresetOutput, error)
	GetPresetAsync(ctx workflow.Context, input *mediaconvert.GetPresetInput) *MediaConvertGetPresetFuture

	GetQueue(ctx workflow.Context, input *mediaconvert.GetQueueInput) (*mediaconvert.GetQueueOutput, error)
	GetQueueAsync(ctx workflow.Context, input *mediaconvert.GetQueueInput) *MediaConvertGetQueueFuture

	ListJobTemplates(ctx workflow.Context, input *mediaconvert.ListJobTemplatesInput) (*mediaconvert.ListJobTemplatesOutput, error)
	ListJobTemplatesAsync(ctx workflow.Context, input *mediaconvert.ListJobTemplatesInput) *MediaConvertListJobTemplatesFuture

	ListJobs(ctx workflow.Context, input *mediaconvert.ListJobsInput) (*mediaconvert.ListJobsOutput, error)
	ListJobsAsync(ctx workflow.Context, input *mediaconvert.ListJobsInput) *MediaConvertListJobsFuture

	ListPresets(ctx workflow.Context, input *mediaconvert.ListPresetsInput) (*mediaconvert.ListPresetsOutput, error)
	ListPresetsAsync(ctx workflow.Context, input *mediaconvert.ListPresetsInput) *MediaConvertListPresetsFuture

	ListQueues(ctx workflow.Context, input *mediaconvert.ListQueuesInput) (*mediaconvert.ListQueuesOutput, error)
	ListQueuesAsync(ctx workflow.Context, input *mediaconvert.ListQueuesInput) *MediaConvertListQueuesFuture

	ListTagsForResource(ctx workflow.Context, input *mediaconvert.ListTagsForResourceInput) (*mediaconvert.ListTagsForResourceOutput, error)
	ListTagsForResourceAsync(ctx workflow.Context, input *mediaconvert.ListTagsForResourceInput) *MediaConvertListTagsForResourceFuture

	TagResource(ctx workflow.Context, input *mediaconvert.TagResourceInput) (*mediaconvert.TagResourceOutput, error)
	TagResourceAsync(ctx workflow.Context, input *mediaconvert.TagResourceInput) *MediaConvertTagResourceFuture

	UntagResource(ctx workflow.Context, input *mediaconvert.UntagResourceInput) (*mediaconvert.UntagResourceOutput, error)
	UntagResourceAsync(ctx workflow.Context, input *mediaconvert.UntagResourceInput) *MediaConvertUntagResourceFuture

	UpdateJobTemplate(ctx workflow.Context, input *mediaconvert.UpdateJobTemplateInput) (*mediaconvert.UpdateJobTemplateOutput, error)
	UpdateJobTemplateAsync(ctx workflow.Context, input *mediaconvert.UpdateJobTemplateInput) *MediaConvertUpdateJobTemplateFuture

	UpdatePreset(ctx workflow.Context, input *mediaconvert.UpdatePresetInput) (*mediaconvert.UpdatePresetOutput, error)
	UpdatePresetAsync(ctx workflow.Context, input *mediaconvert.UpdatePresetInput) *MediaConvertUpdatePresetFuture

	UpdateQueue(ctx workflow.Context, input *mediaconvert.UpdateQueueInput) (*mediaconvert.UpdateQueueOutput, error)
	UpdateQueueAsync(ctx workflow.Context, input *mediaconvert.UpdateQueueInput) *MediaConvertUpdateQueueFuture
}

func NewClient() Client {
	return &stub{}
}
