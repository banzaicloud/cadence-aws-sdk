// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7

package snsstub

import (
	"github.com/aws/aws-sdk-go/service/sns"
	"go.uber.org/cadence/workflow"

	"github.com/banzaicloud/cadence-aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	AddPermission(ctx workflow.Context, input *sns.AddPermissionInput) (*sns.AddPermissionOutput, error)
	AddPermissionAsync(ctx workflow.Context, input *sns.AddPermissionInput) *SNSAddPermissionFuture

	CheckIfPhoneNumberIsOptedOut(ctx workflow.Context, input *sns.CheckIfPhoneNumberIsOptedOutInput) (*sns.CheckIfPhoneNumberIsOptedOutOutput, error)
	CheckIfPhoneNumberIsOptedOutAsync(ctx workflow.Context, input *sns.CheckIfPhoneNumberIsOptedOutInput) *SNSCheckIfPhoneNumberIsOptedOutFuture

	ConfirmSubscription(ctx workflow.Context, input *sns.ConfirmSubscriptionInput) (*sns.ConfirmSubscriptionOutput, error)
	ConfirmSubscriptionAsync(ctx workflow.Context, input *sns.ConfirmSubscriptionInput) *SNSConfirmSubscriptionFuture

	CreatePlatformApplication(ctx workflow.Context, input *sns.CreatePlatformApplicationInput) (*sns.CreatePlatformApplicationOutput, error)
	CreatePlatformApplicationAsync(ctx workflow.Context, input *sns.CreatePlatformApplicationInput) *SNSCreatePlatformApplicationFuture

	CreatePlatformEndpoint(ctx workflow.Context, input *sns.CreatePlatformEndpointInput) (*sns.CreatePlatformEndpointOutput, error)
	CreatePlatformEndpointAsync(ctx workflow.Context, input *sns.CreatePlatformEndpointInput) *SNSCreatePlatformEndpointFuture

	CreateTopic(ctx workflow.Context, input *sns.CreateTopicInput) (*sns.CreateTopicOutput, error)
	CreateTopicAsync(ctx workflow.Context, input *sns.CreateTopicInput) *SNSCreateTopicFuture

	DeleteEndpoint(ctx workflow.Context, input *sns.DeleteEndpointInput) (*sns.DeleteEndpointOutput, error)
	DeleteEndpointAsync(ctx workflow.Context, input *sns.DeleteEndpointInput) *SNSDeleteEndpointFuture

	DeletePlatformApplication(ctx workflow.Context, input *sns.DeletePlatformApplicationInput) (*sns.DeletePlatformApplicationOutput, error)
	DeletePlatformApplicationAsync(ctx workflow.Context, input *sns.DeletePlatformApplicationInput) *SNSDeletePlatformApplicationFuture

	DeleteTopic(ctx workflow.Context, input *sns.DeleteTopicInput) (*sns.DeleteTopicOutput, error)
	DeleteTopicAsync(ctx workflow.Context, input *sns.DeleteTopicInput) *SNSDeleteTopicFuture

	GetEndpointAttributes(ctx workflow.Context, input *sns.GetEndpointAttributesInput) (*sns.GetEndpointAttributesOutput, error)
	GetEndpointAttributesAsync(ctx workflow.Context, input *sns.GetEndpointAttributesInput) *SNSGetEndpointAttributesFuture

	GetPlatformApplicationAttributes(ctx workflow.Context, input *sns.GetPlatformApplicationAttributesInput) (*sns.GetPlatformApplicationAttributesOutput, error)
	GetPlatformApplicationAttributesAsync(ctx workflow.Context, input *sns.GetPlatformApplicationAttributesInput) *SNSGetPlatformApplicationAttributesFuture

	GetSMSAttributes(ctx workflow.Context, input *sns.GetSMSAttributesInput) (*sns.GetSMSAttributesOutput, error)
	GetSMSAttributesAsync(ctx workflow.Context, input *sns.GetSMSAttributesInput) *SNSGetSMSAttributesFuture

	GetSubscriptionAttributes(ctx workflow.Context, input *sns.GetSubscriptionAttributesInput) (*sns.GetSubscriptionAttributesOutput, error)
	GetSubscriptionAttributesAsync(ctx workflow.Context, input *sns.GetSubscriptionAttributesInput) *SNSGetSubscriptionAttributesFuture

	GetTopicAttributes(ctx workflow.Context, input *sns.GetTopicAttributesInput) (*sns.GetTopicAttributesOutput, error)
	GetTopicAttributesAsync(ctx workflow.Context, input *sns.GetTopicAttributesInput) *SNSGetTopicAttributesFuture

	ListEndpointsByPlatformApplication(ctx workflow.Context, input *sns.ListEndpointsByPlatformApplicationInput) (*sns.ListEndpointsByPlatformApplicationOutput, error)
	ListEndpointsByPlatformApplicationAsync(ctx workflow.Context, input *sns.ListEndpointsByPlatformApplicationInput) *SNSListEndpointsByPlatformApplicationFuture

	ListPhoneNumbersOptedOut(ctx workflow.Context, input *sns.ListPhoneNumbersOptedOutInput) (*sns.ListPhoneNumbersOptedOutOutput, error)
	ListPhoneNumbersOptedOutAsync(ctx workflow.Context, input *sns.ListPhoneNumbersOptedOutInput) *SNSListPhoneNumbersOptedOutFuture

	ListPlatformApplications(ctx workflow.Context, input *sns.ListPlatformApplicationsInput) (*sns.ListPlatformApplicationsOutput, error)
	ListPlatformApplicationsAsync(ctx workflow.Context, input *sns.ListPlatformApplicationsInput) *SNSListPlatformApplicationsFuture

	ListSubscriptions(ctx workflow.Context, input *sns.ListSubscriptionsInput) (*sns.ListSubscriptionsOutput, error)
	ListSubscriptionsAsync(ctx workflow.Context, input *sns.ListSubscriptionsInput) *SNSListSubscriptionsFuture

	ListSubscriptionsByTopic(ctx workflow.Context, input *sns.ListSubscriptionsByTopicInput) (*sns.ListSubscriptionsByTopicOutput, error)
	ListSubscriptionsByTopicAsync(ctx workflow.Context, input *sns.ListSubscriptionsByTopicInput) *SNSListSubscriptionsByTopicFuture

	ListTagsForResource(ctx workflow.Context, input *sns.ListTagsForResourceInput) (*sns.ListTagsForResourceOutput, error)
	ListTagsForResourceAsync(ctx workflow.Context, input *sns.ListTagsForResourceInput) *SNSListTagsForResourceFuture

	ListTopics(ctx workflow.Context, input *sns.ListTopicsInput) (*sns.ListTopicsOutput, error)
	ListTopicsAsync(ctx workflow.Context, input *sns.ListTopicsInput) *SNSListTopicsFuture

	OptInPhoneNumber(ctx workflow.Context, input *sns.OptInPhoneNumberInput) (*sns.OptInPhoneNumberOutput, error)
	OptInPhoneNumberAsync(ctx workflow.Context, input *sns.OptInPhoneNumberInput) *SNSOptInPhoneNumberFuture

	Publish(ctx workflow.Context, input *sns.PublishInput) (*sns.PublishOutput, error)
	PublishAsync(ctx workflow.Context, input *sns.PublishInput) *SNSPublishFuture

	RemovePermission(ctx workflow.Context, input *sns.RemovePermissionInput) (*sns.RemovePermissionOutput, error)
	RemovePermissionAsync(ctx workflow.Context, input *sns.RemovePermissionInput) *SNSRemovePermissionFuture

	SetEndpointAttributes(ctx workflow.Context, input *sns.SetEndpointAttributesInput) (*sns.SetEndpointAttributesOutput, error)
	SetEndpointAttributesAsync(ctx workflow.Context, input *sns.SetEndpointAttributesInput) *SNSSetEndpointAttributesFuture

	SetPlatformApplicationAttributes(ctx workflow.Context, input *sns.SetPlatformApplicationAttributesInput) (*sns.SetPlatformApplicationAttributesOutput, error)
	SetPlatformApplicationAttributesAsync(ctx workflow.Context, input *sns.SetPlatformApplicationAttributesInput) *SNSSetPlatformApplicationAttributesFuture

	SetSMSAttributes(ctx workflow.Context, input *sns.SetSMSAttributesInput) (*sns.SetSMSAttributesOutput, error)
	SetSMSAttributesAsync(ctx workflow.Context, input *sns.SetSMSAttributesInput) *SNSSetSMSAttributesFuture

	SetSubscriptionAttributes(ctx workflow.Context, input *sns.SetSubscriptionAttributesInput) (*sns.SetSubscriptionAttributesOutput, error)
	SetSubscriptionAttributesAsync(ctx workflow.Context, input *sns.SetSubscriptionAttributesInput) *SNSSetSubscriptionAttributesFuture

	SetTopicAttributes(ctx workflow.Context, input *sns.SetTopicAttributesInput) (*sns.SetTopicAttributesOutput, error)
	SetTopicAttributesAsync(ctx workflow.Context, input *sns.SetTopicAttributesInput) *SNSSetTopicAttributesFuture

	Subscribe(ctx workflow.Context, input *sns.SubscribeInput) (*sns.SubscribeOutput, error)
	SubscribeAsync(ctx workflow.Context, input *sns.SubscribeInput) *SNSSubscribeFuture

	TagResource(ctx workflow.Context, input *sns.TagResourceInput) (*sns.TagResourceOutput, error)
	TagResourceAsync(ctx workflow.Context, input *sns.TagResourceInput) *SNSTagResourceFuture

	Unsubscribe(ctx workflow.Context, input *sns.UnsubscribeInput) (*sns.UnsubscribeOutput, error)
	UnsubscribeAsync(ctx workflow.Context, input *sns.UnsubscribeInput) *SNSUnsubscribeFuture

	UntagResource(ctx workflow.Context, input *sns.UntagResourceInput) (*sns.UntagResourceOutput, error)
	UntagResourceAsync(ctx workflow.Context, input *sns.UntagResourceInput) *SNSUntagResourceFuture
}

func NewClient() Client {
	return &stub{}
}
