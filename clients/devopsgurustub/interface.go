// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7

package devopsgurustub

import (
	"github.com/aws/aws-sdk-go/service/devopsguru"
	"go.uber.org/cadence/workflow"

	"github.com/banzaicloud/cadence-aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	AddNotificationChannel(ctx workflow.Context, input *devopsguru.AddNotificationChannelInput) (*devopsguru.AddNotificationChannelOutput, error)
	AddNotificationChannelAsync(ctx workflow.Context, input *devopsguru.AddNotificationChannelInput) *AddNotificationChannelFuture

	DescribeAccountHealth(ctx workflow.Context, input *devopsguru.DescribeAccountHealthInput) (*devopsguru.DescribeAccountHealthOutput, error)
	DescribeAccountHealthAsync(ctx workflow.Context, input *devopsguru.DescribeAccountHealthInput) *DescribeAccountHealthFuture

	DescribeAccountOverview(ctx workflow.Context, input *devopsguru.DescribeAccountOverviewInput) (*devopsguru.DescribeAccountOverviewOutput, error)
	DescribeAccountOverviewAsync(ctx workflow.Context, input *devopsguru.DescribeAccountOverviewInput) *DescribeAccountOverviewFuture

	DescribeAnomaly(ctx workflow.Context, input *devopsguru.DescribeAnomalyInput) (*devopsguru.DescribeAnomalyOutput, error)
	DescribeAnomalyAsync(ctx workflow.Context, input *devopsguru.DescribeAnomalyInput) *DescribeAnomalyFuture

	DescribeInsight(ctx workflow.Context, input *devopsguru.DescribeInsightInput) (*devopsguru.DescribeInsightOutput, error)
	DescribeInsightAsync(ctx workflow.Context, input *devopsguru.DescribeInsightInput) *DescribeInsightFuture

	DescribeResourceCollectionHealth(ctx workflow.Context, input *devopsguru.DescribeResourceCollectionHealthInput) (*devopsguru.DescribeResourceCollectionHealthOutput, error)
	DescribeResourceCollectionHealthAsync(ctx workflow.Context, input *devopsguru.DescribeResourceCollectionHealthInput) *DescribeResourceCollectionHealthFuture

	DescribeServiceIntegration(ctx workflow.Context, input *devopsguru.DescribeServiceIntegrationInput) (*devopsguru.DescribeServiceIntegrationOutput, error)
	DescribeServiceIntegrationAsync(ctx workflow.Context, input *devopsguru.DescribeServiceIntegrationInput) *DescribeServiceIntegrationFuture

	GetResourceCollection(ctx workflow.Context, input *devopsguru.GetResourceCollectionInput) (*devopsguru.GetResourceCollectionOutput, error)
	GetResourceCollectionAsync(ctx workflow.Context, input *devopsguru.GetResourceCollectionInput) *GetResourceCollectionFuture

	ListAnomaliesForInsight(ctx workflow.Context, input *devopsguru.ListAnomaliesForInsightInput) (*devopsguru.ListAnomaliesForInsightOutput, error)
	ListAnomaliesForInsightAsync(ctx workflow.Context, input *devopsguru.ListAnomaliesForInsightInput) *ListAnomaliesForInsightFuture

	ListEvents(ctx workflow.Context, input *devopsguru.ListEventsInput) (*devopsguru.ListEventsOutput, error)
	ListEventsAsync(ctx workflow.Context, input *devopsguru.ListEventsInput) *ListEventsFuture

	ListInsights(ctx workflow.Context, input *devopsguru.ListInsightsInput) (*devopsguru.ListInsightsOutput, error)
	ListInsightsAsync(ctx workflow.Context, input *devopsguru.ListInsightsInput) *ListInsightsFuture

	ListNotificationChannels(ctx workflow.Context, input *devopsguru.ListNotificationChannelsInput) (*devopsguru.ListNotificationChannelsOutput, error)
	ListNotificationChannelsAsync(ctx workflow.Context, input *devopsguru.ListNotificationChannelsInput) *ListNotificationChannelsFuture

	ListRecommendations(ctx workflow.Context, input *devopsguru.ListRecommendationsInput) (*devopsguru.ListRecommendationsOutput, error)
	ListRecommendationsAsync(ctx workflow.Context, input *devopsguru.ListRecommendationsInput) *ListRecommendationsFuture

	PutFeedback(ctx workflow.Context, input *devopsguru.PutFeedbackInput) (*devopsguru.PutFeedbackOutput, error)
	PutFeedbackAsync(ctx workflow.Context, input *devopsguru.PutFeedbackInput) *PutFeedbackFuture

	RemoveNotificationChannel(ctx workflow.Context, input *devopsguru.RemoveNotificationChannelInput) (*devopsguru.RemoveNotificationChannelOutput, error)
	RemoveNotificationChannelAsync(ctx workflow.Context, input *devopsguru.RemoveNotificationChannelInput) *RemoveNotificationChannelFuture

	SearchInsights(ctx workflow.Context, input *devopsguru.SearchInsightsInput) (*devopsguru.SearchInsightsOutput, error)
	SearchInsightsAsync(ctx workflow.Context, input *devopsguru.SearchInsightsInput) *SearchInsightsFuture

	UpdateResourceCollection(ctx workflow.Context, input *devopsguru.UpdateResourceCollectionInput) (*devopsguru.UpdateResourceCollectionOutput, error)
	UpdateResourceCollectionAsync(ctx workflow.Context, input *devopsguru.UpdateResourceCollectionInput) *UpdateResourceCollectionFuture

	UpdateServiceIntegration(ctx workflow.Context, input *devopsguru.UpdateServiceIntegrationInput) (*devopsguru.UpdateServiceIntegrationOutput, error)
	UpdateServiceIntegrationAsync(ctx workflow.Context, input *devopsguru.UpdateServiceIntegrationInput) *UpdateServiceIntegrationFuture
}

func NewClient() Client {
	return &stub{}
}
