// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.
// Copyright (c) 2020 Banzai Cloud All rights reserved.

package eks

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/eks"
	"github.com/aws/aws-sdk-go/service/eks/eksiface"

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
	client eksiface.EKSAPI

	sessionFactory SessionFactory
}

func NewActivities(sess *session.Session, config ...*aws.Config) *Activities {
	client := eks.New(sess, config...)
	return &Activities{client: client}
}

func NewActivitiesWithSessionFactory(sessionFactory SessionFactory) *Activities {
	return &Activities{sessionFactory: sessionFactory}
}

func (a *Activities) getClient(ctx context.Context) (eksiface.EKSAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}

	return eks.New(sess), nil
}

func (a *Activities) CreateCluster(ctx context.Context, input *eks.CreateClusterInput) (*eks.CreateClusterOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateClusterWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateFargateProfile(ctx context.Context, input *eks.CreateFargateProfileInput) (*eks.CreateFargateProfileOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateFargateProfileWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateNodegroup(ctx context.Context, input *eks.CreateNodegroupInput) (*eks.CreateNodegroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateNodegroupWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteCluster(ctx context.Context, input *eks.DeleteClusterInput) (*eks.DeleteClusterOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteClusterWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteFargateProfile(ctx context.Context, input *eks.DeleteFargateProfileInput) (*eks.DeleteFargateProfileOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteFargateProfileWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteNodegroup(ctx context.Context, input *eks.DeleteNodegroupInput) (*eks.DeleteNodegroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteNodegroupWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeCluster(ctx context.Context, input *eks.DescribeClusterInput) (*eks.DescribeClusterOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeClusterWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeFargateProfile(ctx context.Context, input *eks.DescribeFargateProfileInput) (*eks.DescribeFargateProfileOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeFargateProfileWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeNodegroup(ctx context.Context, input *eks.DescribeNodegroupInput) (*eks.DescribeNodegroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeNodegroupWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeUpdate(ctx context.Context, input *eks.DescribeUpdateInput) (*eks.DescribeUpdateOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeUpdateWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListClusters(ctx context.Context, input *eks.ListClustersInput) (*eks.ListClustersOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListClustersWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListFargateProfiles(ctx context.Context, input *eks.ListFargateProfilesInput) (*eks.ListFargateProfilesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListFargateProfilesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListNodegroups(ctx context.Context, input *eks.ListNodegroupsInput) (*eks.ListNodegroupsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListNodegroupsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListTagsForResource(ctx context.Context, input *eks.ListTagsForResourceInput) (*eks.ListTagsForResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListTagsForResourceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListUpdates(ctx context.Context, input *eks.ListUpdatesInput) (*eks.ListUpdatesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListUpdatesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) TagResource(ctx context.Context, input *eks.TagResourceInput) (*eks.TagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.TagResourceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UntagResource(ctx context.Context, input *eks.UntagResourceInput) (*eks.UntagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UntagResourceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateClusterConfig(ctx context.Context, input *eks.UpdateClusterConfigInput) (*eks.UpdateClusterConfigOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateClusterConfigWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateClusterVersion(ctx context.Context, input *eks.UpdateClusterVersionInput) (*eks.UpdateClusterVersionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateClusterVersionWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateNodegroupConfig(ctx context.Context, input *eks.UpdateNodegroupConfigInput) (*eks.UpdateNodegroupConfigOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateNodegroupConfigWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateNodegroupVersion(ctx context.Context, input *eks.UpdateNodegroupVersionInput) (*eks.UpdateNodegroupVersionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateNodegroupVersionWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) WaitUntilClusterActive(ctx context.Context, input *eks.DescribeClusterInput) error {
	client, err := a.getClient(ctx)
	if err != nil {
		return internal.EncodeError(err)
	}
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return internal.EncodeError(client.WaitUntilClusterActiveWithContext(ctx, input, options...))
	})
}

func (a *Activities) WaitUntilClusterDeleted(ctx context.Context, input *eks.DescribeClusterInput) error {
	client, err := a.getClient(ctx)
	if err != nil {
		return internal.EncodeError(err)
	}
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return internal.EncodeError(client.WaitUntilClusterDeletedWithContext(ctx, input, options...))
	})
}

func (a *Activities) WaitUntilNodegroupActive(ctx context.Context, input *eks.DescribeNodegroupInput) error {
	client, err := a.getClient(ctx)
	if err != nil {
		return internal.EncodeError(err)
	}
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return internal.EncodeError(client.WaitUntilNodegroupActiveWithContext(ctx, input, options...))
	})
}

func (a *Activities) WaitUntilNodegroupDeleted(ctx context.Context, input *eks.DescribeNodegroupInput) error {
	client, err := a.getClient(ctx)
	if err != nil {
		return internal.EncodeError(err)
	}
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return internal.EncodeError(client.WaitUntilNodegroupDeletedWithContext(ctx, input, options...))
	})
}
