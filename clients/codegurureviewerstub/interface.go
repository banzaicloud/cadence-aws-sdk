// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7

package codegurureviewerstub

import (
	"github.com/aws/aws-sdk-go/service/codegurureviewer"
	"go.uber.org/cadence/workflow"

	"github.com/banzaicloud/cadence-aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	AssociateRepository(ctx workflow.Context, input *codegurureviewer.AssociateRepositoryInput) (*codegurureviewer.AssociateRepositoryOutput, error)
	AssociateRepositoryAsync(ctx workflow.Context, input *codegurureviewer.AssociateRepositoryInput) *AssociateRepositoryFuture

	CreateCodeReview(ctx workflow.Context, input *codegurureviewer.CreateCodeReviewInput) (*codegurureviewer.CreateCodeReviewOutput, error)
	CreateCodeReviewAsync(ctx workflow.Context, input *codegurureviewer.CreateCodeReviewInput) *CreateCodeReviewFuture

	DescribeCodeReview(ctx workflow.Context, input *codegurureviewer.DescribeCodeReviewInput) (*codegurureviewer.DescribeCodeReviewOutput, error)
	DescribeCodeReviewAsync(ctx workflow.Context, input *codegurureviewer.DescribeCodeReviewInput) *DescribeCodeReviewFuture

	DescribeRecommendationFeedback(ctx workflow.Context, input *codegurureviewer.DescribeRecommendationFeedbackInput) (*codegurureviewer.DescribeRecommendationFeedbackOutput, error)
	DescribeRecommendationFeedbackAsync(ctx workflow.Context, input *codegurureviewer.DescribeRecommendationFeedbackInput) *DescribeRecommendationFeedbackFuture

	DescribeRepositoryAssociation(ctx workflow.Context, input *codegurureviewer.DescribeRepositoryAssociationInput) (*codegurureviewer.DescribeRepositoryAssociationOutput, error)
	DescribeRepositoryAssociationAsync(ctx workflow.Context, input *codegurureviewer.DescribeRepositoryAssociationInput) *DescribeRepositoryAssociationFuture

	DisassociateRepository(ctx workflow.Context, input *codegurureviewer.DisassociateRepositoryInput) (*codegurureviewer.DisassociateRepositoryOutput, error)
	DisassociateRepositoryAsync(ctx workflow.Context, input *codegurureviewer.DisassociateRepositoryInput) *DisassociateRepositoryFuture

	ListCodeReviews(ctx workflow.Context, input *codegurureviewer.ListCodeReviewsInput) (*codegurureviewer.ListCodeReviewsOutput, error)
	ListCodeReviewsAsync(ctx workflow.Context, input *codegurureviewer.ListCodeReviewsInput) *ListCodeReviewsFuture

	ListRecommendationFeedback(ctx workflow.Context, input *codegurureviewer.ListRecommendationFeedbackInput) (*codegurureviewer.ListRecommendationFeedbackOutput, error)
	ListRecommendationFeedbackAsync(ctx workflow.Context, input *codegurureviewer.ListRecommendationFeedbackInput) *ListRecommendationFeedbackFuture

	ListRecommendations(ctx workflow.Context, input *codegurureviewer.ListRecommendationsInput) (*codegurureviewer.ListRecommendationsOutput, error)
	ListRecommendationsAsync(ctx workflow.Context, input *codegurureviewer.ListRecommendationsInput) *ListRecommendationsFuture

	ListRepositoryAssociations(ctx workflow.Context, input *codegurureviewer.ListRepositoryAssociationsInput) (*codegurureviewer.ListRepositoryAssociationsOutput, error)
	ListRepositoryAssociationsAsync(ctx workflow.Context, input *codegurureviewer.ListRepositoryAssociationsInput) *ListRepositoryAssociationsFuture

	ListTagsForResource(ctx workflow.Context, input *codegurureviewer.ListTagsForResourceInput) (*codegurureviewer.ListTagsForResourceOutput, error)
	ListTagsForResourceAsync(ctx workflow.Context, input *codegurureviewer.ListTagsForResourceInput) *ListTagsForResourceFuture

	PutRecommendationFeedback(ctx workflow.Context, input *codegurureviewer.PutRecommendationFeedbackInput) (*codegurureviewer.PutRecommendationFeedbackOutput, error)
	PutRecommendationFeedbackAsync(ctx workflow.Context, input *codegurureviewer.PutRecommendationFeedbackInput) *PutRecommendationFeedbackFuture

	TagResource(ctx workflow.Context, input *codegurureviewer.TagResourceInput) (*codegurureviewer.TagResourceOutput, error)
	TagResourceAsync(ctx workflow.Context, input *codegurureviewer.TagResourceInput) *TagResourceFuture

	UntagResource(ctx workflow.Context, input *codegurureviewer.UntagResourceInput) (*codegurureviewer.UntagResourceOutput, error)
	UntagResourceAsync(ctx workflow.Context, input *codegurureviewer.UntagResourceInput) *UntagResourceFuture
}

func NewClient() Client {
	return &stub{}
}
