// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7

package connectparticipant

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/connectparticipant"
	"github.com/aws/aws-sdk-go/service/connectparticipant/connectparticipantiface"

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
	client connectparticipantiface.ConnectParticipantAPI

	sessionFactory SessionFactory
}

func NewActivities(sess *session.Session, config ...*aws.Config) *Activities {
	client := connectparticipant.New(sess, config...)
	return &Activities{client: client}
}

func NewActivitiesWithSessionFactory(sessionFactory SessionFactory) *Activities {
	return &Activities{sessionFactory: sessionFactory}
}

func (a *Activities) getClient(ctx context.Context) (connectparticipantiface.ConnectParticipantAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}

	return connectparticipant.New(sess), nil
}

func (a *Activities) CompleteAttachmentUpload(ctx context.Context, input *connectparticipant.CompleteAttachmentUploadInput) (*connectparticipant.CompleteAttachmentUploadOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	internal.SetClientToken(ctx, &input.ClientToken)
	output, err := client.CompleteAttachmentUploadWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateParticipantConnection(ctx context.Context, input *connectparticipant.CreateParticipantConnectionInput) (*connectparticipant.CreateParticipantConnectionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateParticipantConnectionWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DisconnectParticipant(ctx context.Context, input *connectparticipant.DisconnectParticipantInput) (*connectparticipant.DisconnectParticipantOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	internal.SetClientToken(ctx, &input.ClientToken)
	output, err := client.DisconnectParticipantWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetAttachment(ctx context.Context, input *connectparticipant.GetAttachmentInput) (*connectparticipant.GetAttachmentOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetAttachmentWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetTranscript(ctx context.Context, input *connectparticipant.GetTranscriptInput) (*connectparticipant.GetTranscriptOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetTranscriptWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) SendEvent(ctx context.Context, input *connectparticipant.SendEventInput) (*connectparticipant.SendEventOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	internal.SetClientToken(ctx, &input.ClientToken)
	output, err := client.SendEventWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) SendMessage(ctx context.Context, input *connectparticipant.SendMessageInput) (*connectparticipant.SendMessageOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	internal.SetClientToken(ctx, &input.ClientToken)
	output, err := client.SendMessageWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) StartAttachmentUpload(ctx context.Context, input *connectparticipant.StartAttachmentUploadInput) (*connectparticipant.StartAttachmentUploadOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	internal.SetClientToken(ctx, &input.ClientToken)
	output, err := client.StartAttachmentUploadWithContext(ctx, input)

	return output, internal.EncodeError(err)
}
