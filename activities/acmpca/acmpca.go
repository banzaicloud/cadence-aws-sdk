// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7

package acmpca

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/acmpca"
	"github.com/aws/aws-sdk-go/service/acmpca/acmpcaiface"

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
	client acmpcaiface.ACMPCAAPI

	sessionFactory SessionFactory
}

func NewActivities(sess *session.Session, config ...*aws.Config) *Activities {
	client := acmpca.New(sess, config...)
	return &Activities{client: client}
}

func NewActivitiesWithSessionFactory(sessionFactory SessionFactory) *Activities {
	return &Activities{sessionFactory: sessionFactory}
}

func (a *Activities) getClient(ctx context.Context) (acmpcaiface.ACMPCAAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}

	return acmpca.New(sess), nil
}

func (a *Activities) CreateCertificateAuthority(ctx context.Context, input *acmpca.CreateCertificateAuthorityInput) (*acmpca.CreateCertificateAuthorityOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateCertificateAuthorityWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateCertificateAuthorityAuditReport(ctx context.Context, input *acmpca.CreateCertificateAuthorityAuditReportInput) (*acmpca.CreateCertificateAuthorityAuditReportOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateCertificateAuthorityAuditReportWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreatePermission(ctx context.Context, input *acmpca.CreatePermissionInput) (*acmpca.CreatePermissionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreatePermissionWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteCertificateAuthority(ctx context.Context, input *acmpca.DeleteCertificateAuthorityInput) (*acmpca.DeleteCertificateAuthorityOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteCertificateAuthorityWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeletePermission(ctx context.Context, input *acmpca.DeletePermissionInput) (*acmpca.DeletePermissionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeletePermissionWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeletePolicy(ctx context.Context, input *acmpca.DeletePolicyInput) (*acmpca.DeletePolicyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeletePolicyWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeCertificateAuthority(ctx context.Context, input *acmpca.DescribeCertificateAuthorityInput) (*acmpca.DescribeCertificateAuthorityOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeCertificateAuthorityWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeCertificateAuthorityAuditReport(ctx context.Context, input *acmpca.DescribeCertificateAuthorityAuditReportInput) (*acmpca.DescribeCertificateAuthorityAuditReportOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeCertificateAuthorityAuditReportWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetCertificate(ctx context.Context, input *acmpca.GetCertificateInput) (*acmpca.GetCertificateOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetCertificateWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetCertificateAuthorityCertificate(ctx context.Context, input *acmpca.GetCertificateAuthorityCertificateInput) (*acmpca.GetCertificateAuthorityCertificateOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetCertificateAuthorityCertificateWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetCertificateAuthorityCsr(ctx context.Context, input *acmpca.GetCertificateAuthorityCsrInput) (*acmpca.GetCertificateAuthorityCsrOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetCertificateAuthorityCsrWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetPolicy(ctx context.Context, input *acmpca.GetPolicyInput) (*acmpca.GetPolicyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetPolicyWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ImportCertificateAuthorityCertificate(ctx context.Context, input *acmpca.ImportCertificateAuthorityCertificateInput) (*acmpca.ImportCertificateAuthorityCertificateOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ImportCertificateAuthorityCertificateWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) IssueCertificate(ctx context.Context, input *acmpca.IssueCertificateInput) (*acmpca.IssueCertificateOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.IssueCertificateWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListCertificateAuthorities(ctx context.Context, input *acmpca.ListCertificateAuthoritiesInput) (*acmpca.ListCertificateAuthoritiesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListCertificateAuthoritiesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListPermissions(ctx context.Context, input *acmpca.ListPermissionsInput) (*acmpca.ListPermissionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListPermissionsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListTags(ctx context.Context, input *acmpca.ListTagsInput) (*acmpca.ListTagsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListTagsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) PutPolicy(ctx context.Context, input *acmpca.PutPolicyInput) (*acmpca.PutPolicyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.PutPolicyWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) RestoreCertificateAuthority(ctx context.Context, input *acmpca.RestoreCertificateAuthorityInput) (*acmpca.RestoreCertificateAuthorityOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.RestoreCertificateAuthorityWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) RevokeCertificate(ctx context.Context, input *acmpca.RevokeCertificateInput) (*acmpca.RevokeCertificateOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.RevokeCertificateWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) TagCertificateAuthority(ctx context.Context, input *acmpca.TagCertificateAuthorityInput) (*acmpca.TagCertificateAuthorityOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.TagCertificateAuthorityWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UntagCertificateAuthority(ctx context.Context, input *acmpca.UntagCertificateAuthorityInput) (*acmpca.UntagCertificateAuthorityOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UntagCertificateAuthorityWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateCertificateAuthority(ctx context.Context, input *acmpca.UpdateCertificateAuthorityInput) (*acmpca.UpdateCertificateAuthorityOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateCertificateAuthorityWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) WaitUntilAuditReportCreated(ctx context.Context, input *acmpca.DescribeCertificateAuthorityAuditReportInput) error {
	client, err := a.getClient(ctx)
	if err != nil {
		return internal.EncodeError(err)
	}
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return internal.EncodeError(client.WaitUntilAuditReportCreatedWithContext(ctx, input, options...))
	})
}

func (a *Activities) WaitUntilCertificateAuthorityCSRCreated(ctx context.Context, input *acmpca.GetCertificateAuthorityCsrInput) error {
	client, err := a.getClient(ctx)
	if err != nil {
		return internal.EncodeError(err)
	}
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return internal.EncodeError(client.WaitUntilCertificateAuthorityCSRCreatedWithContext(ctx, input, options...))
	})
}

func (a *Activities) WaitUntilCertificateIssued(ctx context.Context, input *acmpca.GetCertificateInput) error {
	client, err := a.getClient(ctx)
	if err != nil {
		return internal.EncodeError(err)
	}
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return internal.EncodeError(client.WaitUntilCertificateIssuedWithContext(ctx, input, options...))
	})
}
