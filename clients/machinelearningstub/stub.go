// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7

package machinelearningstub

import (
	"github.com/aws/aws-sdk-go/service/machinelearning"
	"go.uber.org/cadence/workflow"

	"github.com/banzaicloud/cadence-aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type stub struct{}

type AddTagsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AddTagsFuture) Get(ctx workflow.Context) (*machinelearning.AddTagsOutput, error) {
	var output machinelearning.AddTagsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CreateBatchPredictionFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CreateBatchPredictionFuture) Get(ctx workflow.Context) (*machinelearning.CreateBatchPredictionOutput, error) {
	var output machinelearning.CreateBatchPredictionOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CreateDataSourceFromRDSFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CreateDataSourceFromRDSFuture) Get(ctx workflow.Context) (*machinelearning.CreateDataSourceFromRDSOutput, error) {
	var output machinelearning.CreateDataSourceFromRDSOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CreateDataSourceFromRedshiftFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CreateDataSourceFromRedshiftFuture) Get(ctx workflow.Context) (*machinelearning.CreateDataSourceFromRedshiftOutput, error) {
	var output machinelearning.CreateDataSourceFromRedshiftOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CreateDataSourceFromS3Future struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CreateDataSourceFromS3Future) Get(ctx workflow.Context) (*machinelearning.CreateDataSourceFromS3Output, error) {
	var output machinelearning.CreateDataSourceFromS3Output
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CreateEvaluationFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CreateEvaluationFuture) Get(ctx workflow.Context) (*machinelearning.CreateEvaluationOutput, error) {
	var output machinelearning.CreateEvaluationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CreateMLModelFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CreateMLModelFuture) Get(ctx workflow.Context) (*machinelearning.CreateMLModelOutput, error) {
	var output machinelearning.CreateMLModelOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CreateRealtimeEndpointFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CreateRealtimeEndpointFuture) Get(ctx workflow.Context) (*machinelearning.CreateRealtimeEndpointOutput, error) {
	var output machinelearning.CreateRealtimeEndpointOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DeleteBatchPredictionFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DeleteBatchPredictionFuture) Get(ctx workflow.Context) (*machinelearning.DeleteBatchPredictionOutput, error) {
	var output machinelearning.DeleteBatchPredictionOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DeleteDataSourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DeleteDataSourceFuture) Get(ctx workflow.Context) (*machinelearning.DeleteDataSourceOutput, error) {
	var output machinelearning.DeleteDataSourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DeleteEvaluationFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DeleteEvaluationFuture) Get(ctx workflow.Context) (*machinelearning.DeleteEvaluationOutput, error) {
	var output machinelearning.DeleteEvaluationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DeleteMLModelFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DeleteMLModelFuture) Get(ctx workflow.Context) (*machinelearning.DeleteMLModelOutput, error) {
	var output machinelearning.DeleteMLModelOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DeleteRealtimeEndpointFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DeleteRealtimeEndpointFuture) Get(ctx workflow.Context) (*machinelearning.DeleteRealtimeEndpointOutput, error) {
	var output machinelearning.DeleteRealtimeEndpointOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DeleteTagsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DeleteTagsFuture) Get(ctx workflow.Context) (*machinelearning.DeleteTagsOutput, error) {
	var output machinelearning.DeleteTagsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DescribeBatchPredictionsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DescribeBatchPredictionsFuture) Get(ctx workflow.Context) (*machinelearning.DescribeBatchPredictionsOutput, error) {
	var output machinelearning.DescribeBatchPredictionsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DescribeDataSourcesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DescribeDataSourcesFuture) Get(ctx workflow.Context) (*machinelearning.DescribeDataSourcesOutput, error) {
	var output machinelearning.DescribeDataSourcesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DescribeEvaluationsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DescribeEvaluationsFuture) Get(ctx workflow.Context) (*machinelearning.DescribeEvaluationsOutput, error) {
	var output machinelearning.DescribeEvaluationsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DescribeMLModelsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DescribeMLModelsFuture) Get(ctx workflow.Context) (*machinelearning.DescribeMLModelsOutput, error) {
	var output machinelearning.DescribeMLModelsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DescribeTagsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DescribeTagsFuture) Get(ctx workflow.Context) (*machinelearning.DescribeTagsOutput, error) {
	var output machinelearning.DescribeTagsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GetBatchPredictionFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GetBatchPredictionFuture) Get(ctx workflow.Context) (*machinelearning.GetBatchPredictionOutput, error) {
	var output machinelearning.GetBatchPredictionOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GetDataSourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GetDataSourceFuture) Get(ctx workflow.Context) (*machinelearning.GetDataSourceOutput, error) {
	var output machinelearning.GetDataSourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GetEvaluationFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GetEvaluationFuture) Get(ctx workflow.Context) (*machinelearning.GetEvaluationOutput, error) {
	var output machinelearning.GetEvaluationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GetMLModelFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GetMLModelFuture) Get(ctx workflow.Context) (*machinelearning.GetMLModelOutput, error) {
	var output machinelearning.GetMLModelOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type PredictFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *PredictFuture) Get(ctx workflow.Context) (*machinelearning.PredictOutput, error) {
	var output machinelearning.PredictOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type UpdateBatchPredictionFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *UpdateBatchPredictionFuture) Get(ctx workflow.Context) (*machinelearning.UpdateBatchPredictionOutput, error) {
	var output machinelearning.UpdateBatchPredictionOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type UpdateDataSourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *UpdateDataSourceFuture) Get(ctx workflow.Context) (*machinelearning.UpdateDataSourceOutput, error) {
	var output machinelearning.UpdateDataSourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type UpdateEvaluationFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *UpdateEvaluationFuture) Get(ctx workflow.Context) (*machinelearning.UpdateEvaluationOutput, error) {
	var output machinelearning.UpdateEvaluationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type UpdateMLModelFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *UpdateMLModelFuture) Get(ctx workflow.Context) (*machinelearning.UpdateMLModelOutput, error) {
	var output machinelearning.UpdateMLModelOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *stub) AddTags(ctx workflow.Context, input *machinelearning.AddTagsInput) (*machinelearning.AddTagsOutput, error) {
	var output machinelearning.AddTagsOutput
	err := workflow.ExecuteActivity(ctx, "aws-machinelearning-AddTags", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) AddTagsAsync(ctx workflow.Context, input *machinelearning.AddTagsInput) *AddTagsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-machinelearning-AddTags", input)
	return &AddTagsFuture{Future: future}
}

func (a *stub) CreateBatchPrediction(ctx workflow.Context, input *machinelearning.CreateBatchPredictionInput) (*machinelearning.CreateBatchPredictionOutput, error) {
	var output machinelearning.CreateBatchPredictionOutput
	err := workflow.ExecuteActivity(ctx, "aws-machinelearning-CreateBatchPrediction", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateBatchPredictionAsync(ctx workflow.Context, input *machinelearning.CreateBatchPredictionInput) *CreateBatchPredictionFuture {
	future := workflow.ExecuteActivity(ctx, "aws-machinelearning-CreateBatchPrediction", input)
	return &CreateBatchPredictionFuture{Future: future}
}

func (a *stub) CreateDataSourceFromRDS(ctx workflow.Context, input *machinelearning.CreateDataSourceFromRDSInput) (*machinelearning.CreateDataSourceFromRDSOutput, error) {
	var output machinelearning.CreateDataSourceFromRDSOutput
	err := workflow.ExecuteActivity(ctx, "aws-machinelearning-CreateDataSourceFromRDS", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateDataSourceFromRDSAsync(ctx workflow.Context, input *machinelearning.CreateDataSourceFromRDSInput) *CreateDataSourceFromRDSFuture {
	future := workflow.ExecuteActivity(ctx, "aws-machinelearning-CreateDataSourceFromRDS", input)
	return &CreateDataSourceFromRDSFuture{Future: future}
}

func (a *stub) CreateDataSourceFromRedshift(ctx workflow.Context, input *machinelearning.CreateDataSourceFromRedshiftInput) (*machinelearning.CreateDataSourceFromRedshiftOutput, error) {
	var output machinelearning.CreateDataSourceFromRedshiftOutput
	err := workflow.ExecuteActivity(ctx, "aws-machinelearning-CreateDataSourceFromRedshift", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateDataSourceFromRedshiftAsync(ctx workflow.Context, input *machinelearning.CreateDataSourceFromRedshiftInput) *CreateDataSourceFromRedshiftFuture {
	future := workflow.ExecuteActivity(ctx, "aws-machinelearning-CreateDataSourceFromRedshift", input)
	return &CreateDataSourceFromRedshiftFuture{Future: future}
}

func (a *stub) CreateDataSourceFromS3(ctx workflow.Context, input *machinelearning.CreateDataSourceFromS3Input) (*machinelearning.CreateDataSourceFromS3Output, error) {
	var output machinelearning.CreateDataSourceFromS3Output
	err := workflow.ExecuteActivity(ctx, "aws-machinelearning-CreateDataSourceFromS3", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateDataSourceFromS3Async(ctx workflow.Context, input *machinelearning.CreateDataSourceFromS3Input) *CreateDataSourceFromS3Future {
	future := workflow.ExecuteActivity(ctx, "aws-machinelearning-CreateDataSourceFromS3", input)
	return &CreateDataSourceFromS3Future{Future: future}
}

func (a *stub) CreateEvaluation(ctx workflow.Context, input *machinelearning.CreateEvaluationInput) (*machinelearning.CreateEvaluationOutput, error) {
	var output machinelearning.CreateEvaluationOutput
	err := workflow.ExecuteActivity(ctx, "aws-machinelearning-CreateEvaluation", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateEvaluationAsync(ctx workflow.Context, input *machinelearning.CreateEvaluationInput) *CreateEvaluationFuture {
	future := workflow.ExecuteActivity(ctx, "aws-machinelearning-CreateEvaluation", input)
	return &CreateEvaluationFuture{Future: future}
}

func (a *stub) CreateMLModel(ctx workflow.Context, input *machinelearning.CreateMLModelInput) (*machinelearning.CreateMLModelOutput, error) {
	var output machinelearning.CreateMLModelOutput
	err := workflow.ExecuteActivity(ctx, "aws-machinelearning-CreateMLModel", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateMLModelAsync(ctx workflow.Context, input *machinelearning.CreateMLModelInput) *CreateMLModelFuture {
	future := workflow.ExecuteActivity(ctx, "aws-machinelearning-CreateMLModel", input)
	return &CreateMLModelFuture{Future: future}
}

func (a *stub) CreateRealtimeEndpoint(ctx workflow.Context, input *machinelearning.CreateRealtimeEndpointInput) (*machinelearning.CreateRealtimeEndpointOutput, error) {
	var output machinelearning.CreateRealtimeEndpointOutput
	err := workflow.ExecuteActivity(ctx, "aws-machinelearning-CreateRealtimeEndpoint", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateRealtimeEndpointAsync(ctx workflow.Context, input *machinelearning.CreateRealtimeEndpointInput) *CreateRealtimeEndpointFuture {
	future := workflow.ExecuteActivity(ctx, "aws-machinelearning-CreateRealtimeEndpoint", input)
	return &CreateRealtimeEndpointFuture{Future: future}
}

func (a *stub) DeleteBatchPrediction(ctx workflow.Context, input *machinelearning.DeleteBatchPredictionInput) (*machinelearning.DeleteBatchPredictionOutput, error) {
	var output machinelearning.DeleteBatchPredictionOutput
	err := workflow.ExecuteActivity(ctx, "aws-machinelearning-DeleteBatchPrediction", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteBatchPredictionAsync(ctx workflow.Context, input *machinelearning.DeleteBatchPredictionInput) *DeleteBatchPredictionFuture {
	future := workflow.ExecuteActivity(ctx, "aws-machinelearning-DeleteBatchPrediction", input)
	return &DeleteBatchPredictionFuture{Future: future}
}

func (a *stub) DeleteDataSource(ctx workflow.Context, input *machinelearning.DeleteDataSourceInput) (*machinelearning.DeleteDataSourceOutput, error) {
	var output machinelearning.DeleteDataSourceOutput
	err := workflow.ExecuteActivity(ctx, "aws-machinelearning-DeleteDataSource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteDataSourceAsync(ctx workflow.Context, input *machinelearning.DeleteDataSourceInput) *DeleteDataSourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws-machinelearning-DeleteDataSource", input)
	return &DeleteDataSourceFuture{Future: future}
}

func (a *stub) DeleteEvaluation(ctx workflow.Context, input *machinelearning.DeleteEvaluationInput) (*machinelearning.DeleteEvaluationOutput, error) {
	var output machinelearning.DeleteEvaluationOutput
	err := workflow.ExecuteActivity(ctx, "aws-machinelearning-DeleteEvaluation", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteEvaluationAsync(ctx workflow.Context, input *machinelearning.DeleteEvaluationInput) *DeleteEvaluationFuture {
	future := workflow.ExecuteActivity(ctx, "aws-machinelearning-DeleteEvaluation", input)
	return &DeleteEvaluationFuture{Future: future}
}

func (a *stub) DeleteMLModel(ctx workflow.Context, input *machinelearning.DeleteMLModelInput) (*machinelearning.DeleteMLModelOutput, error) {
	var output machinelearning.DeleteMLModelOutput
	err := workflow.ExecuteActivity(ctx, "aws-machinelearning-DeleteMLModel", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteMLModelAsync(ctx workflow.Context, input *machinelearning.DeleteMLModelInput) *DeleteMLModelFuture {
	future := workflow.ExecuteActivity(ctx, "aws-machinelearning-DeleteMLModel", input)
	return &DeleteMLModelFuture{Future: future}
}

func (a *stub) DeleteRealtimeEndpoint(ctx workflow.Context, input *machinelearning.DeleteRealtimeEndpointInput) (*machinelearning.DeleteRealtimeEndpointOutput, error) {
	var output machinelearning.DeleteRealtimeEndpointOutput
	err := workflow.ExecuteActivity(ctx, "aws-machinelearning-DeleteRealtimeEndpoint", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteRealtimeEndpointAsync(ctx workflow.Context, input *machinelearning.DeleteRealtimeEndpointInput) *DeleteRealtimeEndpointFuture {
	future := workflow.ExecuteActivity(ctx, "aws-machinelearning-DeleteRealtimeEndpoint", input)
	return &DeleteRealtimeEndpointFuture{Future: future}
}

func (a *stub) DeleteTags(ctx workflow.Context, input *machinelearning.DeleteTagsInput) (*machinelearning.DeleteTagsOutput, error) {
	var output machinelearning.DeleteTagsOutput
	err := workflow.ExecuteActivity(ctx, "aws-machinelearning-DeleteTags", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteTagsAsync(ctx workflow.Context, input *machinelearning.DeleteTagsInput) *DeleteTagsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-machinelearning-DeleteTags", input)
	return &DeleteTagsFuture{Future: future}
}

func (a *stub) DescribeBatchPredictions(ctx workflow.Context, input *machinelearning.DescribeBatchPredictionsInput) (*machinelearning.DescribeBatchPredictionsOutput, error) {
	var output machinelearning.DescribeBatchPredictionsOutput
	err := workflow.ExecuteActivity(ctx, "aws-machinelearning-DescribeBatchPredictions", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeBatchPredictionsAsync(ctx workflow.Context, input *machinelearning.DescribeBatchPredictionsInput) *DescribeBatchPredictionsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-machinelearning-DescribeBatchPredictions", input)
	return &DescribeBatchPredictionsFuture{Future: future}
}

func (a *stub) DescribeDataSources(ctx workflow.Context, input *machinelearning.DescribeDataSourcesInput) (*machinelearning.DescribeDataSourcesOutput, error) {
	var output machinelearning.DescribeDataSourcesOutput
	err := workflow.ExecuteActivity(ctx, "aws-machinelearning-DescribeDataSources", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeDataSourcesAsync(ctx workflow.Context, input *machinelearning.DescribeDataSourcesInput) *DescribeDataSourcesFuture {
	future := workflow.ExecuteActivity(ctx, "aws-machinelearning-DescribeDataSources", input)
	return &DescribeDataSourcesFuture{Future: future}
}

func (a *stub) DescribeEvaluations(ctx workflow.Context, input *machinelearning.DescribeEvaluationsInput) (*machinelearning.DescribeEvaluationsOutput, error) {
	var output machinelearning.DescribeEvaluationsOutput
	err := workflow.ExecuteActivity(ctx, "aws-machinelearning-DescribeEvaluations", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeEvaluationsAsync(ctx workflow.Context, input *machinelearning.DescribeEvaluationsInput) *DescribeEvaluationsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-machinelearning-DescribeEvaluations", input)
	return &DescribeEvaluationsFuture{Future: future}
}

func (a *stub) DescribeMLModels(ctx workflow.Context, input *machinelearning.DescribeMLModelsInput) (*machinelearning.DescribeMLModelsOutput, error) {
	var output machinelearning.DescribeMLModelsOutput
	err := workflow.ExecuteActivity(ctx, "aws-machinelearning-DescribeMLModels", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeMLModelsAsync(ctx workflow.Context, input *machinelearning.DescribeMLModelsInput) *DescribeMLModelsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-machinelearning-DescribeMLModels", input)
	return &DescribeMLModelsFuture{Future: future}
}

func (a *stub) DescribeTags(ctx workflow.Context, input *machinelearning.DescribeTagsInput) (*machinelearning.DescribeTagsOutput, error) {
	var output machinelearning.DescribeTagsOutput
	err := workflow.ExecuteActivity(ctx, "aws-machinelearning-DescribeTags", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeTagsAsync(ctx workflow.Context, input *machinelearning.DescribeTagsInput) *DescribeTagsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-machinelearning-DescribeTags", input)
	return &DescribeTagsFuture{Future: future}
}

func (a *stub) GetBatchPrediction(ctx workflow.Context, input *machinelearning.GetBatchPredictionInput) (*machinelearning.GetBatchPredictionOutput, error) {
	var output machinelearning.GetBatchPredictionOutput
	err := workflow.ExecuteActivity(ctx, "aws-machinelearning-GetBatchPrediction", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetBatchPredictionAsync(ctx workflow.Context, input *machinelearning.GetBatchPredictionInput) *GetBatchPredictionFuture {
	future := workflow.ExecuteActivity(ctx, "aws-machinelearning-GetBatchPrediction", input)
	return &GetBatchPredictionFuture{Future: future}
}

func (a *stub) GetDataSource(ctx workflow.Context, input *machinelearning.GetDataSourceInput) (*machinelearning.GetDataSourceOutput, error) {
	var output machinelearning.GetDataSourceOutput
	err := workflow.ExecuteActivity(ctx, "aws-machinelearning-GetDataSource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetDataSourceAsync(ctx workflow.Context, input *machinelearning.GetDataSourceInput) *GetDataSourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws-machinelearning-GetDataSource", input)
	return &GetDataSourceFuture{Future: future}
}

func (a *stub) GetEvaluation(ctx workflow.Context, input *machinelearning.GetEvaluationInput) (*machinelearning.GetEvaluationOutput, error) {
	var output machinelearning.GetEvaluationOutput
	err := workflow.ExecuteActivity(ctx, "aws-machinelearning-GetEvaluation", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetEvaluationAsync(ctx workflow.Context, input *machinelearning.GetEvaluationInput) *GetEvaluationFuture {
	future := workflow.ExecuteActivity(ctx, "aws-machinelearning-GetEvaluation", input)
	return &GetEvaluationFuture{Future: future}
}

func (a *stub) GetMLModel(ctx workflow.Context, input *machinelearning.GetMLModelInput) (*machinelearning.GetMLModelOutput, error) {
	var output machinelearning.GetMLModelOutput
	err := workflow.ExecuteActivity(ctx, "aws-machinelearning-GetMLModel", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetMLModelAsync(ctx workflow.Context, input *machinelearning.GetMLModelInput) *GetMLModelFuture {
	future := workflow.ExecuteActivity(ctx, "aws-machinelearning-GetMLModel", input)
	return &GetMLModelFuture{Future: future}
}

func (a *stub) Predict(ctx workflow.Context, input *machinelearning.PredictInput) (*machinelearning.PredictOutput, error) {
	var output machinelearning.PredictOutput
	err := workflow.ExecuteActivity(ctx, "aws-machinelearning-Predict", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) PredictAsync(ctx workflow.Context, input *machinelearning.PredictInput) *PredictFuture {
	future := workflow.ExecuteActivity(ctx, "aws-machinelearning-Predict", input)
	return &PredictFuture{Future: future}
}

func (a *stub) UpdateBatchPrediction(ctx workflow.Context, input *machinelearning.UpdateBatchPredictionInput) (*machinelearning.UpdateBatchPredictionOutput, error) {
	var output machinelearning.UpdateBatchPredictionOutput
	err := workflow.ExecuteActivity(ctx, "aws-machinelearning-UpdateBatchPrediction", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateBatchPredictionAsync(ctx workflow.Context, input *machinelearning.UpdateBatchPredictionInput) *UpdateBatchPredictionFuture {
	future := workflow.ExecuteActivity(ctx, "aws-machinelearning-UpdateBatchPrediction", input)
	return &UpdateBatchPredictionFuture{Future: future}
}

func (a *stub) UpdateDataSource(ctx workflow.Context, input *machinelearning.UpdateDataSourceInput) (*machinelearning.UpdateDataSourceOutput, error) {
	var output machinelearning.UpdateDataSourceOutput
	err := workflow.ExecuteActivity(ctx, "aws-machinelearning-UpdateDataSource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateDataSourceAsync(ctx workflow.Context, input *machinelearning.UpdateDataSourceInput) *UpdateDataSourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws-machinelearning-UpdateDataSource", input)
	return &UpdateDataSourceFuture{Future: future}
}

func (a *stub) UpdateEvaluation(ctx workflow.Context, input *machinelearning.UpdateEvaluationInput) (*machinelearning.UpdateEvaluationOutput, error) {
	var output machinelearning.UpdateEvaluationOutput
	err := workflow.ExecuteActivity(ctx, "aws-machinelearning-UpdateEvaluation", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateEvaluationAsync(ctx workflow.Context, input *machinelearning.UpdateEvaluationInput) *UpdateEvaluationFuture {
	future := workflow.ExecuteActivity(ctx, "aws-machinelearning-UpdateEvaluation", input)
	return &UpdateEvaluationFuture{Future: future}
}

func (a *stub) UpdateMLModel(ctx workflow.Context, input *machinelearning.UpdateMLModelInput) (*machinelearning.UpdateMLModelOutput, error) {
	var output machinelearning.UpdateMLModelOutput
	err := workflow.ExecuteActivity(ctx, "aws-machinelearning-UpdateMLModel", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateMLModelAsync(ctx workflow.Context, input *machinelearning.UpdateMLModelInput) *UpdateMLModelFuture {
	future := workflow.ExecuteActivity(ctx, "aws-machinelearning-UpdateMLModel", input)
	return &UpdateMLModelFuture{Future: future}
}

func (a *stub) WaitUntilBatchPredictionAvailable(ctx workflow.Context, input *machinelearning.DescribeBatchPredictionsInput) error {
	return workflow.ExecuteActivity(ctx, "aws-machinelearning-WaitUntilBatchPredictionAvailable", input).Get(ctx, nil)
}

func (a *stub) WaitUntilBatchPredictionAvailableAsync(ctx workflow.Context, input *machinelearning.DescribeBatchPredictionsInput) *clients.VoidFuture {
	future := workflow.ExecuteActivity(ctx, "aws-machinelearning-WaitUntilBatchPredictionAvailable", input)
	return clients.NewVoidFuture(future)
}

func (a *stub) WaitUntilDataSourceAvailable(ctx workflow.Context, input *machinelearning.DescribeDataSourcesInput) error {
	return workflow.ExecuteActivity(ctx, "aws-machinelearning-WaitUntilDataSourceAvailable", input).Get(ctx, nil)
}

func (a *stub) WaitUntilDataSourceAvailableAsync(ctx workflow.Context, input *machinelearning.DescribeDataSourcesInput) *clients.VoidFuture {
	future := workflow.ExecuteActivity(ctx, "aws-machinelearning-WaitUntilDataSourceAvailable", input)
	return clients.NewVoidFuture(future)
}

func (a *stub) WaitUntilEvaluationAvailable(ctx workflow.Context, input *machinelearning.DescribeEvaluationsInput) error {
	return workflow.ExecuteActivity(ctx, "aws-machinelearning-WaitUntilEvaluationAvailable", input).Get(ctx, nil)
}

func (a *stub) WaitUntilEvaluationAvailableAsync(ctx workflow.Context, input *machinelearning.DescribeEvaluationsInput) *clients.VoidFuture {
	future := workflow.ExecuteActivity(ctx, "aws-machinelearning-WaitUntilEvaluationAvailable", input)
	return clients.NewVoidFuture(future)
}

func (a *stub) WaitUntilMLModelAvailable(ctx workflow.Context, input *machinelearning.DescribeMLModelsInput) error {
	return workflow.ExecuteActivity(ctx, "aws-machinelearning-WaitUntilMLModelAvailable", input).Get(ctx, nil)
}

func (a *stub) WaitUntilMLModelAvailableAsync(ctx workflow.Context, input *machinelearning.DescribeMLModelsInput) *clients.VoidFuture {
	future := workflow.ExecuteActivity(ctx, "aws-machinelearning-WaitUntilMLModelAvailable", input)
	return clients.NewVoidFuture(future)
}
