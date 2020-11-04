// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.
// Copyright (c) 2020 Banzai Cloud All rights reserved.

package comprehendmedical

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/comprehendmedical"
	"github.com/aws/aws-sdk-go/service/comprehendmedical/comprehendmedicaliface"

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
	client comprehendmedicaliface.ComprehendMedicalAPI

	sessionFactory SessionFactory
}

func NewActivities(sess *session.Session, config ...*aws.Config) *Activities {
	client := comprehendmedical.New(sess, config...)
	return &Activities{client: client}
}

func NewActivitiesWithSessionFactory(sessionFactory SessionFactory) *Activities {
	return &Activities{sessionFactory: sessionFactory}
}

func (a *Activities) getClient(ctx context.Context) (comprehendmedicaliface.ComprehendMedicalAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}

	return comprehendmedical.New(sess), nil
}

func (a *Activities) DescribeEntitiesDetectionV2Job(ctx context.Context, input *comprehendmedical.DescribeEntitiesDetectionV2JobInput) (*comprehendmedical.DescribeEntitiesDetectionV2JobOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeEntitiesDetectionV2JobWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeICD10CMInferenceJob(ctx context.Context, input *comprehendmedical.DescribeICD10CMInferenceJobInput) (*comprehendmedical.DescribeICD10CMInferenceJobOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeICD10CMInferenceJobWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribePHIDetectionJob(ctx context.Context, input *comprehendmedical.DescribePHIDetectionJobInput) (*comprehendmedical.DescribePHIDetectionJobOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribePHIDetectionJobWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeRxNormInferenceJob(ctx context.Context, input *comprehendmedical.DescribeRxNormInferenceJobInput) (*comprehendmedical.DescribeRxNormInferenceJobOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeRxNormInferenceJobWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DetectEntities(ctx context.Context, input *comprehendmedical.DetectEntitiesInput) (*comprehendmedical.DetectEntitiesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DetectEntitiesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DetectEntitiesV2(ctx context.Context, input *comprehendmedical.DetectEntitiesV2Input) (*comprehendmedical.DetectEntitiesV2Output, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DetectEntitiesV2WithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DetectPHI(ctx context.Context, input *comprehendmedical.DetectPHIInput) (*comprehendmedical.DetectPHIOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DetectPHIWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) InferICD10CM(ctx context.Context, input *comprehendmedical.InferICD10CMInput) (*comprehendmedical.InferICD10CMOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.InferICD10CMWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) InferRxNorm(ctx context.Context, input *comprehendmedical.InferRxNormInput) (*comprehendmedical.InferRxNormOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.InferRxNormWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListEntitiesDetectionV2Jobs(ctx context.Context, input *comprehendmedical.ListEntitiesDetectionV2JobsInput) (*comprehendmedical.ListEntitiesDetectionV2JobsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListEntitiesDetectionV2JobsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListICD10CMInferenceJobs(ctx context.Context, input *comprehendmedical.ListICD10CMInferenceJobsInput) (*comprehendmedical.ListICD10CMInferenceJobsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListICD10CMInferenceJobsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListPHIDetectionJobs(ctx context.Context, input *comprehendmedical.ListPHIDetectionJobsInput) (*comprehendmedical.ListPHIDetectionJobsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListPHIDetectionJobsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListRxNormInferenceJobs(ctx context.Context, input *comprehendmedical.ListRxNormInferenceJobsInput) (*comprehendmedical.ListRxNormInferenceJobsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListRxNormInferenceJobsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) StartEntitiesDetectionV2Job(ctx context.Context, input *comprehendmedical.StartEntitiesDetectionV2JobInput) (*comprehendmedical.StartEntitiesDetectionV2JobOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.StartEntitiesDetectionV2JobWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) StartICD10CMInferenceJob(ctx context.Context, input *comprehendmedical.StartICD10CMInferenceJobInput) (*comprehendmedical.StartICD10CMInferenceJobOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.StartICD10CMInferenceJobWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) StartPHIDetectionJob(ctx context.Context, input *comprehendmedical.StartPHIDetectionJobInput) (*comprehendmedical.StartPHIDetectionJobOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.StartPHIDetectionJobWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) StartRxNormInferenceJob(ctx context.Context, input *comprehendmedical.StartRxNormInferenceJobInput) (*comprehendmedical.StartRxNormInferenceJobOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.StartRxNormInferenceJobWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) StopEntitiesDetectionV2Job(ctx context.Context, input *comprehendmedical.StopEntitiesDetectionV2JobInput) (*comprehendmedical.StopEntitiesDetectionV2JobOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.StopEntitiesDetectionV2JobWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) StopICD10CMInferenceJob(ctx context.Context, input *comprehendmedical.StopICD10CMInferenceJobInput) (*comprehendmedical.StopICD10CMInferenceJobOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.StopICD10CMInferenceJobWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) StopPHIDetectionJob(ctx context.Context, input *comprehendmedical.StopPHIDetectionJobInput) (*comprehendmedical.StopPHIDetectionJobOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.StopPHIDetectionJobWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) StopRxNormInferenceJob(ctx context.Context, input *comprehendmedical.StopRxNormInferenceJobInput) (*comprehendmedical.StopRxNormInferenceJobOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.StopRxNormInferenceJobWithContext(ctx, input)

	return output, internal.EncodeError(err)
}
