// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.
// Copyright (c) 2020 Banzai Cloud All rights reserved.

package mobile

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/mobile"
	"github.com/aws/aws-sdk-go/service/mobile/mobileiface"

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
	client mobileiface.MobileAPI

	sessionFactory SessionFactory
}

func NewActivities(sess *session.Session, config ...*aws.Config) *Activities {
	client := mobile.New(sess, config...)
	return &Activities{client: client}
}

func NewActivitiesWithSessionFactory(sessionFactory SessionFactory) *Activities {
	return &Activities{sessionFactory: sessionFactory}
}

func (a *Activities) getClient(ctx context.Context) (mobileiface.MobileAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}

	return mobile.New(sess), nil
}

func (a *Activities) CreateProject(ctx context.Context, input *mobile.CreateProjectInput) (*mobile.CreateProjectOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateProjectWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteProject(ctx context.Context, input *mobile.DeleteProjectInput) (*mobile.DeleteProjectOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteProjectWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeBundle(ctx context.Context, input *mobile.DescribeBundleInput) (*mobile.DescribeBundleOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeBundleWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeProject(ctx context.Context, input *mobile.DescribeProjectInput) (*mobile.DescribeProjectOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeProjectWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ExportBundle(ctx context.Context, input *mobile.ExportBundleInput) (*mobile.ExportBundleOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ExportBundleWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ExportProject(ctx context.Context, input *mobile.ExportProjectInput) (*mobile.ExportProjectOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ExportProjectWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListBundles(ctx context.Context, input *mobile.ListBundlesInput) (*mobile.ListBundlesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListBundlesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListProjects(ctx context.Context, input *mobile.ListProjectsInput) (*mobile.ListProjectsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListProjectsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateProject(ctx context.Context, input *mobile.UpdateProjectInput) (*mobile.UpdateProjectOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateProjectWithContext(ctx, input)

	return output, internal.EncodeError(err)
}