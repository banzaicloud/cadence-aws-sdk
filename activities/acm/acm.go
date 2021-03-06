// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7

package acm

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/acm"
	"github.com/aws/aws-sdk-go/service/acm/acmiface"

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
	client acmiface.ACMAPI

	sessionFactory SessionFactory
}

func NewActivities(sess *session.Session, config ...*aws.Config) *Activities {
	client := acm.New(sess, config...)
	return &Activities{client: client}
}

func NewActivitiesWithSessionFactory(sessionFactory SessionFactory) *Activities {
	return &Activities{sessionFactory: sessionFactory}
}

func (a *Activities) getClient(ctx context.Context) (acmiface.ACMAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}

	return acm.New(sess), nil
}

func (a *Activities) AddTagsToCertificate(ctx context.Context, input *acm.AddTagsToCertificateInput) (*acm.AddTagsToCertificateOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.AddTagsToCertificateWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteCertificate(ctx context.Context, input *acm.DeleteCertificateInput) (*acm.DeleteCertificateOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteCertificateWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeCertificate(ctx context.Context, input *acm.DescribeCertificateInput) (*acm.DescribeCertificateOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeCertificateWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ExportCertificate(ctx context.Context, input *acm.ExportCertificateInput) (*acm.ExportCertificateOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ExportCertificateWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetCertificate(ctx context.Context, input *acm.GetCertificateInput) (*acm.GetCertificateOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetCertificateWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ImportCertificate(ctx context.Context, input *acm.ImportCertificateInput) (*acm.ImportCertificateOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ImportCertificateWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListCertificates(ctx context.Context, input *acm.ListCertificatesInput) (*acm.ListCertificatesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListCertificatesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListTagsForCertificate(ctx context.Context, input *acm.ListTagsForCertificateInput) (*acm.ListTagsForCertificateOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListTagsForCertificateWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) RemoveTagsFromCertificate(ctx context.Context, input *acm.RemoveTagsFromCertificateInput) (*acm.RemoveTagsFromCertificateOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.RemoveTagsFromCertificateWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) RenewCertificate(ctx context.Context, input *acm.RenewCertificateInput) (*acm.RenewCertificateOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.RenewCertificateWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) RequestCertificate(ctx context.Context, input *acm.RequestCertificateInput) (*acm.RequestCertificateOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.RequestCertificateWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ResendValidationEmail(ctx context.Context, input *acm.ResendValidationEmailInput) (*acm.ResendValidationEmailOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ResendValidationEmailWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateCertificateOptions(ctx context.Context, input *acm.UpdateCertificateOptionsInput) (*acm.UpdateCertificateOptionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateCertificateOptionsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) WaitUntilCertificateValidated(ctx context.Context, input *acm.DescribeCertificateInput) error {
	client, err := a.getClient(ctx)
	if err != nil {
		return internal.EncodeError(err)
	}
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return internal.EncodeError(client.WaitUntilCertificateValidatedWithContext(ctx, input, options...))
	})
}
