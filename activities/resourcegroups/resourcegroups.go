// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7

package resourcegroups

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/resourcegroups"
	"github.com/aws/aws-sdk-go/service/resourcegroups/resourcegroupsiface"

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
	client resourcegroupsiface.ResourceGroupsAPI

	sessionFactory SessionFactory
}

func NewActivities(sess *session.Session, config ...*aws.Config) *Activities {
	client := resourcegroups.New(sess, config...)
	return &Activities{client: client}
}

func NewActivitiesWithSessionFactory(sessionFactory SessionFactory) *Activities {
	return &Activities{sessionFactory: sessionFactory}
}

func (a *Activities) getClient(ctx context.Context) (resourcegroupsiface.ResourceGroupsAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}

	return resourcegroups.New(sess), nil
}

func (a *Activities) CreateGroup(ctx context.Context, input *resourcegroups.CreateGroupInput) (*resourcegroups.CreateGroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateGroupWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteGroup(ctx context.Context, input *resourcegroups.DeleteGroupInput) (*resourcegroups.DeleteGroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteGroupWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetGroup(ctx context.Context, input *resourcegroups.GetGroupInput) (*resourcegroups.GetGroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetGroupWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetGroupConfiguration(ctx context.Context, input *resourcegroups.GetGroupConfigurationInput) (*resourcegroups.GetGroupConfigurationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetGroupConfigurationWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetGroupQuery(ctx context.Context, input *resourcegroups.GetGroupQueryInput) (*resourcegroups.GetGroupQueryOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetGroupQueryWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetTags(ctx context.Context, input *resourcegroups.GetTagsInput) (*resourcegroups.GetTagsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetTagsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GroupResources(ctx context.Context, input *resourcegroups.GroupResourcesInput) (*resourcegroups.GroupResourcesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GroupResourcesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListGroupResources(ctx context.Context, input *resourcegroups.ListGroupResourcesInput) (*resourcegroups.ListGroupResourcesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListGroupResourcesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListGroups(ctx context.Context, input *resourcegroups.ListGroupsInput) (*resourcegroups.ListGroupsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListGroupsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) PutGroupConfiguration(ctx context.Context, input *resourcegroups.PutGroupConfigurationInput) (*resourcegroups.PutGroupConfigurationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.PutGroupConfigurationWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) SearchResources(ctx context.Context, input *resourcegroups.SearchResourcesInput) (*resourcegroups.SearchResourcesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.SearchResourcesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) Tag(ctx context.Context, input *resourcegroups.TagInput) (*resourcegroups.TagOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.TagWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UngroupResources(ctx context.Context, input *resourcegroups.UngroupResourcesInput) (*resourcegroups.UngroupResourcesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UngroupResourcesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) Untag(ctx context.Context, input *resourcegroups.UntagInput) (*resourcegroups.UntagOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UntagWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateGroup(ctx context.Context, input *resourcegroups.UpdateGroupInput) (*resourcegroups.UpdateGroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateGroupWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateGroupQuery(ctx context.Context, input *resourcegroups.UpdateGroupQueryInput) (*resourcegroups.UpdateGroupQueryOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateGroupQueryWithContext(ctx, input)

	return output, internal.EncodeError(err)
}
