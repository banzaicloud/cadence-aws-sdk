// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7

package cloudhsmv2

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudhsmv2"
	"github.com/aws/aws-sdk-go/service/cloudhsmv2/cloudhsmv2iface"

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
	client cloudhsmv2iface.CloudHSMV2API

	sessionFactory SessionFactory
}

func NewActivities(sess *session.Session, config ...*aws.Config) *Activities {
	client := cloudhsmv2.New(sess, config...)
	return &Activities{client: client}
}

func NewActivitiesWithSessionFactory(sessionFactory SessionFactory) *Activities {
	return &Activities{sessionFactory: sessionFactory}
}

func (a *Activities) getClient(ctx context.Context) (cloudhsmv2iface.CloudHSMV2API, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}

	return cloudhsmv2.New(sess), nil
}

func (a *Activities) CopyBackupToRegion(ctx context.Context, input *cloudhsmv2.CopyBackupToRegionInput) (*cloudhsmv2.CopyBackupToRegionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CopyBackupToRegionWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateCluster(ctx context.Context, input *cloudhsmv2.CreateClusterInput) (*cloudhsmv2.CreateClusterOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateClusterWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateHsm(ctx context.Context, input *cloudhsmv2.CreateHsmInput) (*cloudhsmv2.CreateHsmOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateHsmWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteBackup(ctx context.Context, input *cloudhsmv2.DeleteBackupInput) (*cloudhsmv2.DeleteBackupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteBackupWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteCluster(ctx context.Context, input *cloudhsmv2.DeleteClusterInput) (*cloudhsmv2.DeleteClusterOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteClusterWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteHsm(ctx context.Context, input *cloudhsmv2.DeleteHsmInput) (*cloudhsmv2.DeleteHsmOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteHsmWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeBackups(ctx context.Context, input *cloudhsmv2.DescribeBackupsInput) (*cloudhsmv2.DescribeBackupsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeBackupsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeClusters(ctx context.Context, input *cloudhsmv2.DescribeClustersInput) (*cloudhsmv2.DescribeClustersOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeClustersWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) InitializeCluster(ctx context.Context, input *cloudhsmv2.InitializeClusterInput) (*cloudhsmv2.InitializeClusterOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.InitializeClusterWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListTags(ctx context.Context, input *cloudhsmv2.ListTagsInput) (*cloudhsmv2.ListTagsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListTagsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ModifyBackupAttributes(ctx context.Context, input *cloudhsmv2.ModifyBackupAttributesInput) (*cloudhsmv2.ModifyBackupAttributesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ModifyBackupAttributesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ModifyCluster(ctx context.Context, input *cloudhsmv2.ModifyClusterInput) (*cloudhsmv2.ModifyClusterOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ModifyClusterWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) RestoreBackup(ctx context.Context, input *cloudhsmv2.RestoreBackupInput) (*cloudhsmv2.RestoreBackupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.RestoreBackupWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) TagResource(ctx context.Context, input *cloudhsmv2.TagResourceInput) (*cloudhsmv2.TagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.TagResourceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UntagResource(ctx context.Context, input *cloudhsmv2.UntagResourceInput) (*cloudhsmv2.UntagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UntagResourceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}
