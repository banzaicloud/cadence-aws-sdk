// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.
// Copyright (c) 2020 Banzai Cloud All rights reserved.

package personalize

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/personalize"
	"github.com/aws/aws-sdk-go/service/personalize/personalizeiface"

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
	client personalizeiface.PersonalizeAPI

	sessionFactory SessionFactory
}

func NewActivities(sess *session.Session, config ...*aws.Config) *Activities {
	client := personalize.New(sess, config...)
	return &Activities{client: client}
}

func NewActivitiesWithSessionFactory(sessionFactory SessionFactory) *Activities {
	return &Activities{sessionFactory: sessionFactory}
}

func (a *Activities) getClient(ctx context.Context) (personalizeiface.PersonalizeAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}

	return personalize.New(sess), nil
}

func (a *Activities) CreateBatchInferenceJob(ctx context.Context, input *personalize.CreateBatchInferenceJobInput) (*personalize.CreateBatchInferenceJobOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateBatchInferenceJobWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateCampaign(ctx context.Context, input *personalize.CreateCampaignInput) (*personalize.CreateCampaignOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateCampaignWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateDataset(ctx context.Context, input *personalize.CreateDatasetInput) (*personalize.CreateDatasetOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateDatasetWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateDatasetGroup(ctx context.Context, input *personalize.CreateDatasetGroupInput) (*personalize.CreateDatasetGroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateDatasetGroupWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateDatasetImportJob(ctx context.Context, input *personalize.CreateDatasetImportJobInput) (*personalize.CreateDatasetImportJobOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateDatasetImportJobWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateEventTracker(ctx context.Context, input *personalize.CreateEventTrackerInput) (*personalize.CreateEventTrackerOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateEventTrackerWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateFilter(ctx context.Context, input *personalize.CreateFilterInput) (*personalize.CreateFilterOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateFilterWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateSchema(ctx context.Context, input *personalize.CreateSchemaInput) (*personalize.CreateSchemaOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateSchemaWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateSolution(ctx context.Context, input *personalize.CreateSolutionInput) (*personalize.CreateSolutionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateSolutionWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateSolutionVersion(ctx context.Context, input *personalize.CreateSolutionVersionInput) (*personalize.CreateSolutionVersionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateSolutionVersionWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteCampaign(ctx context.Context, input *personalize.DeleteCampaignInput) (*personalize.DeleteCampaignOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteCampaignWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteDataset(ctx context.Context, input *personalize.DeleteDatasetInput) (*personalize.DeleteDatasetOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteDatasetWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteDatasetGroup(ctx context.Context, input *personalize.DeleteDatasetGroupInput) (*personalize.DeleteDatasetGroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteDatasetGroupWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteEventTracker(ctx context.Context, input *personalize.DeleteEventTrackerInput) (*personalize.DeleteEventTrackerOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteEventTrackerWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteFilter(ctx context.Context, input *personalize.DeleteFilterInput) (*personalize.DeleteFilterOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteFilterWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteSchema(ctx context.Context, input *personalize.DeleteSchemaInput) (*personalize.DeleteSchemaOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteSchemaWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteSolution(ctx context.Context, input *personalize.DeleteSolutionInput) (*personalize.DeleteSolutionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteSolutionWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeAlgorithm(ctx context.Context, input *personalize.DescribeAlgorithmInput) (*personalize.DescribeAlgorithmOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeAlgorithmWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeBatchInferenceJob(ctx context.Context, input *personalize.DescribeBatchInferenceJobInput) (*personalize.DescribeBatchInferenceJobOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeBatchInferenceJobWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeCampaign(ctx context.Context, input *personalize.DescribeCampaignInput) (*personalize.DescribeCampaignOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeCampaignWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeDataset(ctx context.Context, input *personalize.DescribeDatasetInput) (*personalize.DescribeDatasetOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeDatasetWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeDatasetGroup(ctx context.Context, input *personalize.DescribeDatasetGroupInput) (*personalize.DescribeDatasetGroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeDatasetGroupWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeDatasetImportJob(ctx context.Context, input *personalize.DescribeDatasetImportJobInput) (*personalize.DescribeDatasetImportJobOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeDatasetImportJobWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeEventTracker(ctx context.Context, input *personalize.DescribeEventTrackerInput) (*personalize.DescribeEventTrackerOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeEventTrackerWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeFeatureTransformation(ctx context.Context, input *personalize.DescribeFeatureTransformationInput) (*personalize.DescribeFeatureTransformationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeFeatureTransformationWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeFilter(ctx context.Context, input *personalize.DescribeFilterInput) (*personalize.DescribeFilterOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeFilterWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeRecipe(ctx context.Context, input *personalize.DescribeRecipeInput) (*personalize.DescribeRecipeOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeRecipeWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeSchema(ctx context.Context, input *personalize.DescribeSchemaInput) (*personalize.DescribeSchemaOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeSchemaWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeSolution(ctx context.Context, input *personalize.DescribeSolutionInput) (*personalize.DescribeSolutionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeSolutionWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeSolutionVersion(ctx context.Context, input *personalize.DescribeSolutionVersionInput) (*personalize.DescribeSolutionVersionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeSolutionVersionWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetSolutionMetrics(ctx context.Context, input *personalize.GetSolutionMetricsInput) (*personalize.GetSolutionMetricsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetSolutionMetricsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListBatchInferenceJobs(ctx context.Context, input *personalize.ListBatchInferenceJobsInput) (*personalize.ListBatchInferenceJobsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListBatchInferenceJobsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListCampaigns(ctx context.Context, input *personalize.ListCampaignsInput) (*personalize.ListCampaignsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListCampaignsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListDatasetGroups(ctx context.Context, input *personalize.ListDatasetGroupsInput) (*personalize.ListDatasetGroupsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListDatasetGroupsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListDatasetImportJobs(ctx context.Context, input *personalize.ListDatasetImportJobsInput) (*personalize.ListDatasetImportJobsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListDatasetImportJobsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListDatasets(ctx context.Context, input *personalize.ListDatasetsInput) (*personalize.ListDatasetsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListDatasetsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListEventTrackers(ctx context.Context, input *personalize.ListEventTrackersInput) (*personalize.ListEventTrackersOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListEventTrackersWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListFilters(ctx context.Context, input *personalize.ListFiltersInput) (*personalize.ListFiltersOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListFiltersWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListRecipes(ctx context.Context, input *personalize.ListRecipesInput) (*personalize.ListRecipesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListRecipesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListSchemas(ctx context.Context, input *personalize.ListSchemasInput) (*personalize.ListSchemasOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListSchemasWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListSolutionVersions(ctx context.Context, input *personalize.ListSolutionVersionsInput) (*personalize.ListSolutionVersionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListSolutionVersionsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListSolutions(ctx context.Context, input *personalize.ListSolutionsInput) (*personalize.ListSolutionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListSolutionsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateCampaign(ctx context.Context, input *personalize.UpdateCampaignInput) (*personalize.UpdateCampaignOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateCampaignWithContext(ctx, input)

	return output, internal.EncodeError(err)
}
