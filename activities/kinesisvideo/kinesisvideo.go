// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.
// Copyright (c) 2020 Banzai Cloud All rights reserved.

package kinesisvideo

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kinesisvideo"
	"github.com/aws/aws-sdk-go/service/kinesisvideo/kinesisvideoiface"

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
	client kinesisvideoiface.KinesisVideoAPI

	sessionFactory SessionFactory
}

func NewActivities(sess *session.Session, config ...*aws.Config) *Activities {
	client := kinesisvideo.New(sess, config...)
	return &Activities{client: client}
}

func NewActivitiesWithSessionFactory(sessionFactory SessionFactory) *Activities {
	return &Activities{sessionFactory: sessionFactory}
}

func (a *Activities) getClient(ctx context.Context) (kinesisvideoiface.KinesisVideoAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}

	return kinesisvideo.New(sess), nil
}

func (a *Activities) CreateSignalingChannel(ctx context.Context, input *kinesisvideo.CreateSignalingChannelInput) (*kinesisvideo.CreateSignalingChannelOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateSignalingChannelWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateStream(ctx context.Context, input *kinesisvideo.CreateStreamInput) (*kinesisvideo.CreateStreamOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateStreamWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteSignalingChannel(ctx context.Context, input *kinesisvideo.DeleteSignalingChannelInput) (*kinesisvideo.DeleteSignalingChannelOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteSignalingChannelWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteStream(ctx context.Context, input *kinesisvideo.DeleteStreamInput) (*kinesisvideo.DeleteStreamOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteStreamWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeSignalingChannel(ctx context.Context, input *kinesisvideo.DescribeSignalingChannelInput) (*kinesisvideo.DescribeSignalingChannelOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeSignalingChannelWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeStream(ctx context.Context, input *kinesisvideo.DescribeStreamInput) (*kinesisvideo.DescribeStreamOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeStreamWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetDataEndpoint(ctx context.Context, input *kinesisvideo.GetDataEndpointInput) (*kinesisvideo.GetDataEndpointOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetDataEndpointWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetSignalingChannelEndpoint(ctx context.Context, input *kinesisvideo.GetSignalingChannelEndpointInput) (*kinesisvideo.GetSignalingChannelEndpointOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetSignalingChannelEndpointWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListSignalingChannels(ctx context.Context, input *kinesisvideo.ListSignalingChannelsInput) (*kinesisvideo.ListSignalingChannelsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListSignalingChannelsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListStreams(ctx context.Context, input *kinesisvideo.ListStreamsInput) (*kinesisvideo.ListStreamsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListStreamsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListTagsForResource(ctx context.Context, input *kinesisvideo.ListTagsForResourceInput) (*kinesisvideo.ListTagsForResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListTagsForResourceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListTagsForStream(ctx context.Context, input *kinesisvideo.ListTagsForStreamInput) (*kinesisvideo.ListTagsForStreamOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListTagsForStreamWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) TagResource(ctx context.Context, input *kinesisvideo.TagResourceInput) (*kinesisvideo.TagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.TagResourceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) TagStream(ctx context.Context, input *kinesisvideo.TagStreamInput) (*kinesisvideo.TagStreamOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.TagStreamWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UntagResource(ctx context.Context, input *kinesisvideo.UntagResourceInput) (*kinesisvideo.UntagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UntagResourceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UntagStream(ctx context.Context, input *kinesisvideo.UntagStreamInput) (*kinesisvideo.UntagStreamOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UntagStreamWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateDataRetention(ctx context.Context, input *kinesisvideo.UpdateDataRetentionInput) (*kinesisvideo.UpdateDataRetentionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateDataRetentionWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateSignalingChannel(ctx context.Context, input *kinesisvideo.UpdateSignalingChannelInput) (*kinesisvideo.UpdateSignalingChannelOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateSignalingChannelWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateStream(ctx context.Context, input *kinesisvideo.UpdateStreamInput) (*kinesisvideo.UpdateStreamOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateStreamWithContext(ctx, input)

	return output, internal.EncodeError(err)
}
