// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7

package connect

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/connect"
	"github.com/aws/aws-sdk-go/service/connect/connectiface"

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
	client connectiface.ConnectAPI

	sessionFactory SessionFactory
}

func NewActivities(sess *session.Session, config ...*aws.Config) *Activities {
	client := connect.New(sess, config...)
	return &Activities{client: client}
}

func NewActivitiesWithSessionFactory(sessionFactory SessionFactory) *Activities {
	return &Activities{sessionFactory: sessionFactory}
}

func (a *Activities) getClient(ctx context.Context) (connectiface.ConnectAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}

	return connect.New(sess), nil
}

func (a *Activities) AssociateRoutingProfileQueues(ctx context.Context, input *connect.AssociateRoutingProfileQueuesInput) (*connect.AssociateRoutingProfileQueuesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.AssociateRoutingProfileQueuesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateContactFlow(ctx context.Context, input *connect.CreateContactFlowInput) (*connect.CreateContactFlowOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateContactFlowWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateRoutingProfile(ctx context.Context, input *connect.CreateRoutingProfileInput) (*connect.CreateRoutingProfileOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateRoutingProfileWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateUser(ctx context.Context, input *connect.CreateUserInput) (*connect.CreateUserOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateUserWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteUser(ctx context.Context, input *connect.DeleteUserInput) (*connect.DeleteUserOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteUserWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeContactFlow(ctx context.Context, input *connect.DescribeContactFlowInput) (*connect.DescribeContactFlowOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeContactFlowWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeRoutingProfile(ctx context.Context, input *connect.DescribeRoutingProfileInput) (*connect.DescribeRoutingProfileOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeRoutingProfileWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeUser(ctx context.Context, input *connect.DescribeUserInput) (*connect.DescribeUserOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeUserWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeUserHierarchyGroup(ctx context.Context, input *connect.DescribeUserHierarchyGroupInput) (*connect.DescribeUserHierarchyGroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeUserHierarchyGroupWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeUserHierarchyStructure(ctx context.Context, input *connect.DescribeUserHierarchyStructureInput) (*connect.DescribeUserHierarchyStructureOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeUserHierarchyStructureWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DisassociateRoutingProfileQueues(ctx context.Context, input *connect.DisassociateRoutingProfileQueuesInput) (*connect.DisassociateRoutingProfileQueuesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DisassociateRoutingProfileQueuesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetContactAttributes(ctx context.Context, input *connect.GetContactAttributesInput) (*connect.GetContactAttributesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetContactAttributesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetCurrentMetricData(ctx context.Context, input *connect.GetCurrentMetricDataInput) (*connect.GetCurrentMetricDataOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetCurrentMetricDataWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetFederationToken(ctx context.Context, input *connect.GetFederationTokenInput) (*connect.GetFederationTokenOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetFederationTokenWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetMetricData(ctx context.Context, input *connect.GetMetricDataInput) (*connect.GetMetricDataOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetMetricDataWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListContactFlows(ctx context.Context, input *connect.ListContactFlowsInput) (*connect.ListContactFlowsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListContactFlowsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListHoursOfOperations(ctx context.Context, input *connect.ListHoursOfOperationsInput) (*connect.ListHoursOfOperationsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListHoursOfOperationsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListPhoneNumbers(ctx context.Context, input *connect.ListPhoneNumbersInput) (*connect.ListPhoneNumbersOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListPhoneNumbersWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListPrompts(ctx context.Context, input *connect.ListPromptsInput) (*connect.ListPromptsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListPromptsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListQueues(ctx context.Context, input *connect.ListQueuesInput) (*connect.ListQueuesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListQueuesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListRoutingProfileQueues(ctx context.Context, input *connect.ListRoutingProfileQueuesInput) (*connect.ListRoutingProfileQueuesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListRoutingProfileQueuesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListRoutingProfiles(ctx context.Context, input *connect.ListRoutingProfilesInput) (*connect.ListRoutingProfilesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListRoutingProfilesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListSecurityProfiles(ctx context.Context, input *connect.ListSecurityProfilesInput) (*connect.ListSecurityProfilesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListSecurityProfilesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListTagsForResource(ctx context.Context, input *connect.ListTagsForResourceInput) (*connect.ListTagsForResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListTagsForResourceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListUserHierarchyGroups(ctx context.Context, input *connect.ListUserHierarchyGroupsInput) (*connect.ListUserHierarchyGroupsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListUserHierarchyGroupsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListUsers(ctx context.Context, input *connect.ListUsersInput) (*connect.ListUsersOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListUsersWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ResumeContactRecording(ctx context.Context, input *connect.ResumeContactRecordingInput) (*connect.ResumeContactRecordingOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ResumeContactRecordingWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) StartChatContact(ctx context.Context, input *connect.StartChatContactInput) (*connect.StartChatContactOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	internal.SetClientToken(ctx, &input.ClientToken)
	output, err := client.StartChatContactWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) StartContactRecording(ctx context.Context, input *connect.StartContactRecordingInput) (*connect.StartContactRecordingOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.StartContactRecordingWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) StartOutboundVoiceContact(ctx context.Context, input *connect.StartOutboundVoiceContactInput) (*connect.StartOutboundVoiceContactOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	internal.SetClientToken(ctx, &input.ClientToken)
	output, err := client.StartOutboundVoiceContactWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) StopContact(ctx context.Context, input *connect.StopContactInput) (*connect.StopContactOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.StopContactWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) StopContactRecording(ctx context.Context, input *connect.StopContactRecordingInput) (*connect.StopContactRecordingOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.StopContactRecordingWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) SuspendContactRecording(ctx context.Context, input *connect.SuspendContactRecordingInput) (*connect.SuspendContactRecordingOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.SuspendContactRecordingWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) TagResource(ctx context.Context, input *connect.TagResourceInput) (*connect.TagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.TagResourceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UntagResource(ctx context.Context, input *connect.UntagResourceInput) (*connect.UntagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UntagResourceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateContactAttributes(ctx context.Context, input *connect.UpdateContactAttributesInput) (*connect.UpdateContactAttributesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateContactAttributesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateContactFlowContent(ctx context.Context, input *connect.UpdateContactFlowContentInput) (*connect.UpdateContactFlowContentOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateContactFlowContentWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateContactFlowName(ctx context.Context, input *connect.UpdateContactFlowNameInput) (*connect.UpdateContactFlowNameOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateContactFlowNameWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateRoutingProfileConcurrency(ctx context.Context, input *connect.UpdateRoutingProfileConcurrencyInput) (*connect.UpdateRoutingProfileConcurrencyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateRoutingProfileConcurrencyWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateRoutingProfileDefaultOutboundQueue(ctx context.Context, input *connect.UpdateRoutingProfileDefaultOutboundQueueInput) (*connect.UpdateRoutingProfileDefaultOutboundQueueOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateRoutingProfileDefaultOutboundQueueWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateRoutingProfileName(ctx context.Context, input *connect.UpdateRoutingProfileNameInput) (*connect.UpdateRoutingProfileNameOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateRoutingProfileNameWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateRoutingProfileQueues(ctx context.Context, input *connect.UpdateRoutingProfileQueuesInput) (*connect.UpdateRoutingProfileQueuesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateRoutingProfileQueuesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateUserHierarchy(ctx context.Context, input *connect.UpdateUserHierarchyInput) (*connect.UpdateUserHierarchyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateUserHierarchyWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateUserIdentityInfo(ctx context.Context, input *connect.UpdateUserIdentityInfoInput) (*connect.UpdateUserIdentityInfoOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateUserIdentityInfoWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateUserPhoneConfig(ctx context.Context, input *connect.UpdateUserPhoneConfigInput) (*connect.UpdateUserPhoneConfigOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateUserPhoneConfigWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateUserRoutingProfile(ctx context.Context, input *connect.UpdateUserRoutingProfileInput) (*connect.UpdateUserRoutingProfileOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateUserRoutingProfileWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateUserSecurityProfiles(ctx context.Context, input *connect.UpdateUserSecurityProfilesInput) (*connect.UpdateUserSecurityProfilesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateUserSecurityProfilesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}
