// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.
// Copyright (c) 2020 Banzai Cloud All rights reserved.

package mediaconvert

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/mediaconvert"
	"github.com/aws/aws-sdk-go/service/mediaconvert/mediaconvertiface"

	"github.com/banzaicloud/cadence-aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken

type _ request.Option

// SessionFactory returns an aws.Session based on the activity context.
type SessionFactory interface {
	Session(ctx context.Context) (*session.Session, error)
}

type Activities struct {
	client mediaconvertiface.MediaConvertAPI

	sessionFactory SessionFactory
}

func NewActivities(sess *session.Session, config ...*aws.Config) *Activities {
	client := mediaconvert.New(sess, config...)
	return &Activities{client: client}
}

func NewActivitiesWithSessionFactory(sessionFactory SessionFactory) *Activities {
	return &Activities{sessionFactory: sessionFactory}
}

func (a *Activities) getClient(ctx context.Context) (mediaconvertiface.MediaConvertAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}

	return mediaconvert.New(sess), nil
}

func (a *Activities) AssociateCertificate(ctx context.Context, input *mediaconvert.AssociateCertificateInput) (*mediaconvert.AssociateCertificateOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.AssociateCertificateWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CancelJob(ctx context.Context, input *mediaconvert.CancelJobInput) (*mediaconvert.CancelJobOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CancelJobWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateJob(ctx context.Context, input *mediaconvert.CreateJobInput) (*mediaconvert.CreateJobOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateJobWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateJobTemplate(ctx context.Context, input *mediaconvert.CreateJobTemplateInput) (*mediaconvert.CreateJobTemplateOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateJobTemplateWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreatePreset(ctx context.Context, input *mediaconvert.CreatePresetInput) (*mediaconvert.CreatePresetOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreatePresetWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateQueue(ctx context.Context, input *mediaconvert.CreateQueueInput) (*mediaconvert.CreateQueueOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateQueueWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteJobTemplate(ctx context.Context, input *mediaconvert.DeleteJobTemplateInput) (*mediaconvert.DeleteJobTemplateOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteJobTemplateWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeletePreset(ctx context.Context, input *mediaconvert.DeletePresetInput) (*mediaconvert.DeletePresetOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeletePresetWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteQueue(ctx context.Context, input *mediaconvert.DeleteQueueInput) (*mediaconvert.DeleteQueueOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteQueueWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeEndpoints(ctx context.Context, input *mediaconvert.DescribeEndpointsInput) (*mediaconvert.DescribeEndpointsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeEndpointsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DisassociateCertificate(ctx context.Context, input *mediaconvert.DisassociateCertificateInput) (*mediaconvert.DisassociateCertificateOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DisassociateCertificateWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetJob(ctx context.Context, input *mediaconvert.GetJobInput) (*mediaconvert.GetJobOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetJobWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetJobTemplate(ctx context.Context, input *mediaconvert.GetJobTemplateInput) (*mediaconvert.GetJobTemplateOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetJobTemplateWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetPreset(ctx context.Context, input *mediaconvert.GetPresetInput) (*mediaconvert.GetPresetOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetPresetWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetQueue(ctx context.Context, input *mediaconvert.GetQueueInput) (*mediaconvert.GetQueueOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetQueueWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListJobTemplates(ctx context.Context, input *mediaconvert.ListJobTemplatesInput) (*mediaconvert.ListJobTemplatesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListJobTemplatesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListJobs(ctx context.Context, input *mediaconvert.ListJobsInput) (*mediaconvert.ListJobsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListJobsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListPresets(ctx context.Context, input *mediaconvert.ListPresetsInput) (*mediaconvert.ListPresetsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListPresetsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListQueues(ctx context.Context, input *mediaconvert.ListQueuesInput) (*mediaconvert.ListQueuesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListQueuesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListTagsForResource(ctx context.Context, input *mediaconvert.ListTagsForResourceInput) (*mediaconvert.ListTagsForResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListTagsForResourceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) TagResource(ctx context.Context, input *mediaconvert.TagResourceInput) (*mediaconvert.TagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.TagResourceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UntagResource(ctx context.Context, input *mediaconvert.UntagResourceInput) (*mediaconvert.UntagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UntagResourceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateJobTemplate(ctx context.Context, input *mediaconvert.UpdateJobTemplateInput) (*mediaconvert.UpdateJobTemplateOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateJobTemplateWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdatePreset(ctx context.Context, input *mediaconvert.UpdatePresetInput) (*mediaconvert.UpdatePresetOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdatePresetWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateQueue(ctx context.Context, input *mediaconvert.UpdateQueueInput) (*mediaconvert.UpdateQueueOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateQueueWithContext(ctx, input)

	return output, internal.EncodeError(err)
}
