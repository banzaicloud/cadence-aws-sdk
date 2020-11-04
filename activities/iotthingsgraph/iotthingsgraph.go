// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.
// Copyright (c) 2020 Banzai Cloud All rights reserved.

package iotthingsgraph

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/iotthingsgraph"
	"github.com/aws/aws-sdk-go/service/iotthingsgraph/iotthingsgraphiface"

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
	client iotthingsgraphiface.IoTThingsGraphAPI

	sessionFactory SessionFactory
}

func NewActivities(sess *session.Session, config ...*aws.Config) *Activities {
	client := iotthingsgraph.New(sess, config...)
	return &Activities{client: client}
}

func NewActivitiesWithSessionFactory(sessionFactory SessionFactory) *Activities {
	return &Activities{sessionFactory: sessionFactory}
}

func (a *Activities) getClient(ctx context.Context) (iotthingsgraphiface.IoTThingsGraphAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}

	return iotthingsgraph.New(sess), nil
}

func (a *Activities) AssociateEntityToThing(ctx context.Context, input *iotthingsgraph.AssociateEntityToThingInput) (*iotthingsgraph.AssociateEntityToThingOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.AssociateEntityToThingWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateFlowTemplate(ctx context.Context, input *iotthingsgraph.CreateFlowTemplateInput) (*iotthingsgraph.CreateFlowTemplateOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateFlowTemplateWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateSystemInstance(ctx context.Context, input *iotthingsgraph.CreateSystemInstanceInput) (*iotthingsgraph.CreateSystemInstanceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateSystemInstanceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateSystemTemplate(ctx context.Context, input *iotthingsgraph.CreateSystemTemplateInput) (*iotthingsgraph.CreateSystemTemplateOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateSystemTemplateWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteFlowTemplate(ctx context.Context, input *iotthingsgraph.DeleteFlowTemplateInput) (*iotthingsgraph.DeleteFlowTemplateOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteFlowTemplateWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteNamespace(ctx context.Context, input *iotthingsgraph.DeleteNamespaceInput) (*iotthingsgraph.DeleteNamespaceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteNamespaceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteSystemInstance(ctx context.Context, input *iotthingsgraph.DeleteSystemInstanceInput) (*iotthingsgraph.DeleteSystemInstanceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteSystemInstanceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteSystemTemplate(ctx context.Context, input *iotthingsgraph.DeleteSystemTemplateInput) (*iotthingsgraph.DeleteSystemTemplateOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteSystemTemplateWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeploySystemInstance(ctx context.Context, input *iotthingsgraph.DeploySystemInstanceInput) (*iotthingsgraph.DeploySystemInstanceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeploySystemInstanceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeprecateFlowTemplate(ctx context.Context, input *iotthingsgraph.DeprecateFlowTemplateInput) (*iotthingsgraph.DeprecateFlowTemplateOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeprecateFlowTemplateWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeprecateSystemTemplate(ctx context.Context, input *iotthingsgraph.DeprecateSystemTemplateInput) (*iotthingsgraph.DeprecateSystemTemplateOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeprecateSystemTemplateWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeNamespace(ctx context.Context, input *iotthingsgraph.DescribeNamespaceInput) (*iotthingsgraph.DescribeNamespaceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeNamespaceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DissociateEntityFromThing(ctx context.Context, input *iotthingsgraph.DissociateEntityFromThingInput) (*iotthingsgraph.DissociateEntityFromThingOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DissociateEntityFromThingWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetEntities(ctx context.Context, input *iotthingsgraph.GetEntitiesInput) (*iotthingsgraph.GetEntitiesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetEntitiesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetFlowTemplate(ctx context.Context, input *iotthingsgraph.GetFlowTemplateInput) (*iotthingsgraph.GetFlowTemplateOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetFlowTemplateWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetFlowTemplateRevisions(ctx context.Context, input *iotthingsgraph.GetFlowTemplateRevisionsInput) (*iotthingsgraph.GetFlowTemplateRevisionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetFlowTemplateRevisionsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetNamespaceDeletionStatus(ctx context.Context, input *iotthingsgraph.GetNamespaceDeletionStatusInput) (*iotthingsgraph.GetNamespaceDeletionStatusOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetNamespaceDeletionStatusWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetSystemInstance(ctx context.Context, input *iotthingsgraph.GetSystemInstanceInput) (*iotthingsgraph.GetSystemInstanceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetSystemInstanceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetSystemTemplate(ctx context.Context, input *iotthingsgraph.GetSystemTemplateInput) (*iotthingsgraph.GetSystemTemplateOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetSystemTemplateWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetSystemTemplateRevisions(ctx context.Context, input *iotthingsgraph.GetSystemTemplateRevisionsInput) (*iotthingsgraph.GetSystemTemplateRevisionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetSystemTemplateRevisionsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetUploadStatus(ctx context.Context, input *iotthingsgraph.GetUploadStatusInput) (*iotthingsgraph.GetUploadStatusOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetUploadStatusWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListFlowExecutionMessages(ctx context.Context, input *iotthingsgraph.ListFlowExecutionMessagesInput) (*iotthingsgraph.ListFlowExecutionMessagesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListFlowExecutionMessagesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListTagsForResource(ctx context.Context, input *iotthingsgraph.ListTagsForResourceInput) (*iotthingsgraph.ListTagsForResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListTagsForResourceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) SearchEntities(ctx context.Context, input *iotthingsgraph.SearchEntitiesInput) (*iotthingsgraph.SearchEntitiesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.SearchEntitiesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) SearchFlowExecutions(ctx context.Context, input *iotthingsgraph.SearchFlowExecutionsInput) (*iotthingsgraph.SearchFlowExecutionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.SearchFlowExecutionsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) SearchFlowTemplates(ctx context.Context, input *iotthingsgraph.SearchFlowTemplatesInput) (*iotthingsgraph.SearchFlowTemplatesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.SearchFlowTemplatesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) SearchSystemInstances(ctx context.Context, input *iotthingsgraph.SearchSystemInstancesInput) (*iotthingsgraph.SearchSystemInstancesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.SearchSystemInstancesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) SearchSystemTemplates(ctx context.Context, input *iotthingsgraph.SearchSystemTemplatesInput) (*iotthingsgraph.SearchSystemTemplatesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.SearchSystemTemplatesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) SearchThings(ctx context.Context, input *iotthingsgraph.SearchThingsInput) (*iotthingsgraph.SearchThingsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.SearchThingsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) TagResource(ctx context.Context, input *iotthingsgraph.TagResourceInput) (*iotthingsgraph.TagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.TagResourceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UndeploySystemInstance(ctx context.Context, input *iotthingsgraph.UndeploySystemInstanceInput) (*iotthingsgraph.UndeploySystemInstanceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UndeploySystemInstanceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UntagResource(ctx context.Context, input *iotthingsgraph.UntagResourceInput) (*iotthingsgraph.UntagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UntagResourceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateFlowTemplate(ctx context.Context, input *iotthingsgraph.UpdateFlowTemplateInput) (*iotthingsgraph.UpdateFlowTemplateOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateFlowTemplateWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateSystemTemplate(ctx context.Context, input *iotthingsgraph.UpdateSystemTemplateInput) (*iotthingsgraph.UpdateSystemTemplateOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateSystemTemplateWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UploadEntityDefinitions(ctx context.Context, input *iotthingsgraph.UploadEntityDefinitionsInput) (*iotthingsgraph.UploadEntityDefinitionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UploadEntityDefinitionsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}
