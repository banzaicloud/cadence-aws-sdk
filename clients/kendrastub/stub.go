// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7

package kendrastub

import (
	"github.com/aws/aws-sdk-go/service/kendra"
	"go.uber.org/cadence/workflow"

	"github.com/banzaicloud/cadence-aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type stub struct{}

type BatchDeleteDocumentFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *BatchDeleteDocumentFuture) Get(ctx workflow.Context) (*kendra.BatchDeleteDocumentOutput, error) {
	var output kendra.BatchDeleteDocumentOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type BatchPutDocumentFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *BatchPutDocumentFuture) Get(ctx workflow.Context) (*kendra.BatchPutDocumentOutput, error) {
	var output kendra.BatchPutDocumentOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CreateDataSourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CreateDataSourceFuture) Get(ctx workflow.Context) (*kendra.CreateDataSourceOutput, error) {
	var output kendra.CreateDataSourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CreateFaqFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CreateFaqFuture) Get(ctx workflow.Context) (*kendra.CreateFaqOutput, error) {
	var output kendra.CreateFaqOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CreateIndexFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CreateIndexFuture) Get(ctx workflow.Context) (*kendra.CreateIndexOutput, error) {
	var output kendra.CreateIndexOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CreateThesaurusFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CreateThesaurusFuture) Get(ctx workflow.Context) (*kendra.CreateThesaurusOutput, error) {
	var output kendra.CreateThesaurusOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DeleteDataSourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DeleteDataSourceFuture) Get(ctx workflow.Context) (*kendra.DeleteDataSourceOutput, error) {
	var output kendra.DeleteDataSourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DeleteFaqFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DeleteFaqFuture) Get(ctx workflow.Context) (*kendra.DeleteFaqOutput, error) {
	var output kendra.DeleteFaqOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DeleteIndexFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DeleteIndexFuture) Get(ctx workflow.Context) (*kendra.DeleteIndexOutput, error) {
	var output kendra.DeleteIndexOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DeleteThesaurusFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DeleteThesaurusFuture) Get(ctx workflow.Context) (*kendra.DeleteThesaurusOutput, error) {
	var output kendra.DeleteThesaurusOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DescribeDataSourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DescribeDataSourceFuture) Get(ctx workflow.Context) (*kendra.DescribeDataSourceOutput, error) {
	var output kendra.DescribeDataSourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DescribeFaqFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DescribeFaqFuture) Get(ctx workflow.Context) (*kendra.DescribeFaqOutput, error) {
	var output kendra.DescribeFaqOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DescribeIndexFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DescribeIndexFuture) Get(ctx workflow.Context) (*kendra.DescribeIndexOutput, error) {
	var output kendra.DescribeIndexOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DescribeThesaurusFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DescribeThesaurusFuture) Get(ctx workflow.Context) (*kendra.DescribeThesaurusOutput, error) {
	var output kendra.DescribeThesaurusOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListDataSourceSyncJobsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListDataSourceSyncJobsFuture) Get(ctx workflow.Context) (*kendra.ListDataSourceSyncJobsOutput, error) {
	var output kendra.ListDataSourceSyncJobsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListDataSourcesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListDataSourcesFuture) Get(ctx workflow.Context) (*kendra.ListDataSourcesOutput, error) {
	var output kendra.ListDataSourcesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListFaqsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListFaqsFuture) Get(ctx workflow.Context) (*kendra.ListFaqsOutput, error) {
	var output kendra.ListFaqsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListIndicesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListIndicesFuture) Get(ctx workflow.Context) (*kendra.ListIndicesOutput, error) {
	var output kendra.ListIndicesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListTagsForResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListTagsForResourceFuture) Get(ctx workflow.Context) (*kendra.ListTagsForResourceOutput, error) {
	var output kendra.ListTagsForResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListThesauriFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListThesauriFuture) Get(ctx workflow.Context) (*kendra.ListThesauriOutput, error) {
	var output kendra.ListThesauriOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type QueryFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *QueryFuture) Get(ctx workflow.Context) (*kendra.QueryOutput, error) {
	var output kendra.QueryOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type StartDataSourceSyncJobFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *StartDataSourceSyncJobFuture) Get(ctx workflow.Context) (*kendra.StartDataSourceSyncJobOutput, error) {
	var output kendra.StartDataSourceSyncJobOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type StopDataSourceSyncJobFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *StopDataSourceSyncJobFuture) Get(ctx workflow.Context) (*kendra.StopDataSourceSyncJobOutput, error) {
	var output kendra.StopDataSourceSyncJobOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SubmitFeedbackFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SubmitFeedbackFuture) Get(ctx workflow.Context) (*kendra.SubmitFeedbackOutput, error) {
	var output kendra.SubmitFeedbackOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type TagResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *TagResourceFuture) Get(ctx workflow.Context) (*kendra.TagResourceOutput, error) {
	var output kendra.TagResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type UntagResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *UntagResourceFuture) Get(ctx workflow.Context) (*kendra.UntagResourceOutput, error) {
	var output kendra.UntagResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type UpdateDataSourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *UpdateDataSourceFuture) Get(ctx workflow.Context) (*kendra.UpdateDataSourceOutput, error) {
	var output kendra.UpdateDataSourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type UpdateIndexFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *UpdateIndexFuture) Get(ctx workflow.Context) (*kendra.UpdateIndexOutput, error) {
	var output kendra.UpdateIndexOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type UpdateThesaurusFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *UpdateThesaurusFuture) Get(ctx workflow.Context) (*kendra.UpdateThesaurusOutput, error) {
	var output kendra.UpdateThesaurusOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *stub) BatchDeleteDocument(ctx workflow.Context, input *kendra.BatchDeleteDocumentInput) (*kendra.BatchDeleteDocumentOutput, error) {
	var output kendra.BatchDeleteDocumentOutput
	err := workflow.ExecuteActivity(ctx, "aws-kendra-BatchDeleteDocument", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) BatchDeleteDocumentAsync(ctx workflow.Context, input *kendra.BatchDeleteDocumentInput) *BatchDeleteDocumentFuture {
	future := workflow.ExecuteActivity(ctx, "aws-kendra-BatchDeleteDocument", input)
	return &BatchDeleteDocumentFuture{Future: future}
}

func (a *stub) BatchPutDocument(ctx workflow.Context, input *kendra.BatchPutDocumentInput) (*kendra.BatchPutDocumentOutput, error) {
	var output kendra.BatchPutDocumentOutput
	err := workflow.ExecuteActivity(ctx, "aws-kendra-BatchPutDocument", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) BatchPutDocumentAsync(ctx workflow.Context, input *kendra.BatchPutDocumentInput) *BatchPutDocumentFuture {
	future := workflow.ExecuteActivity(ctx, "aws-kendra-BatchPutDocument", input)
	return &BatchPutDocumentFuture{Future: future}
}

func (a *stub) CreateDataSource(ctx workflow.Context, input *kendra.CreateDataSourceInput) (*kendra.CreateDataSourceOutput, error) {
	var output kendra.CreateDataSourceOutput
	err := workflow.ExecuteActivity(ctx, "aws-kendra-CreateDataSource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateDataSourceAsync(ctx workflow.Context, input *kendra.CreateDataSourceInput) *CreateDataSourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws-kendra-CreateDataSource", input)
	return &CreateDataSourceFuture{Future: future}
}

func (a *stub) CreateFaq(ctx workflow.Context, input *kendra.CreateFaqInput) (*kendra.CreateFaqOutput, error) {
	var output kendra.CreateFaqOutput
	err := workflow.ExecuteActivity(ctx, "aws-kendra-CreateFaq", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateFaqAsync(ctx workflow.Context, input *kendra.CreateFaqInput) *CreateFaqFuture {
	future := workflow.ExecuteActivity(ctx, "aws-kendra-CreateFaq", input)
	return &CreateFaqFuture{Future: future}
}

func (a *stub) CreateIndex(ctx workflow.Context, input *kendra.CreateIndexInput) (*kendra.CreateIndexOutput, error) {
	var output kendra.CreateIndexOutput
	err := workflow.ExecuteActivity(ctx, "aws-kendra-CreateIndex", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateIndexAsync(ctx workflow.Context, input *kendra.CreateIndexInput) *CreateIndexFuture {
	future := workflow.ExecuteActivity(ctx, "aws-kendra-CreateIndex", input)
	return &CreateIndexFuture{Future: future}
}

func (a *stub) CreateThesaurus(ctx workflow.Context, input *kendra.CreateThesaurusInput) (*kendra.CreateThesaurusOutput, error) {
	var output kendra.CreateThesaurusOutput
	err := workflow.ExecuteActivity(ctx, "aws-kendra-CreateThesaurus", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateThesaurusAsync(ctx workflow.Context, input *kendra.CreateThesaurusInput) *CreateThesaurusFuture {
	future := workflow.ExecuteActivity(ctx, "aws-kendra-CreateThesaurus", input)
	return &CreateThesaurusFuture{Future: future}
}

func (a *stub) DeleteDataSource(ctx workflow.Context, input *kendra.DeleteDataSourceInput) (*kendra.DeleteDataSourceOutput, error) {
	var output kendra.DeleteDataSourceOutput
	err := workflow.ExecuteActivity(ctx, "aws-kendra-DeleteDataSource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteDataSourceAsync(ctx workflow.Context, input *kendra.DeleteDataSourceInput) *DeleteDataSourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws-kendra-DeleteDataSource", input)
	return &DeleteDataSourceFuture{Future: future}
}

func (a *stub) DeleteFaq(ctx workflow.Context, input *kendra.DeleteFaqInput) (*kendra.DeleteFaqOutput, error) {
	var output kendra.DeleteFaqOutput
	err := workflow.ExecuteActivity(ctx, "aws-kendra-DeleteFaq", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteFaqAsync(ctx workflow.Context, input *kendra.DeleteFaqInput) *DeleteFaqFuture {
	future := workflow.ExecuteActivity(ctx, "aws-kendra-DeleteFaq", input)
	return &DeleteFaqFuture{Future: future}
}

func (a *stub) DeleteIndex(ctx workflow.Context, input *kendra.DeleteIndexInput) (*kendra.DeleteIndexOutput, error) {
	var output kendra.DeleteIndexOutput
	err := workflow.ExecuteActivity(ctx, "aws-kendra-DeleteIndex", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteIndexAsync(ctx workflow.Context, input *kendra.DeleteIndexInput) *DeleteIndexFuture {
	future := workflow.ExecuteActivity(ctx, "aws-kendra-DeleteIndex", input)
	return &DeleteIndexFuture{Future: future}
}

func (a *stub) DeleteThesaurus(ctx workflow.Context, input *kendra.DeleteThesaurusInput) (*kendra.DeleteThesaurusOutput, error) {
	var output kendra.DeleteThesaurusOutput
	err := workflow.ExecuteActivity(ctx, "aws-kendra-DeleteThesaurus", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteThesaurusAsync(ctx workflow.Context, input *kendra.DeleteThesaurusInput) *DeleteThesaurusFuture {
	future := workflow.ExecuteActivity(ctx, "aws-kendra-DeleteThesaurus", input)
	return &DeleteThesaurusFuture{Future: future}
}

func (a *stub) DescribeDataSource(ctx workflow.Context, input *kendra.DescribeDataSourceInput) (*kendra.DescribeDataSourceOutput, error) {
	var output kendra.DescribeDataSourceOutput
	err := workflow.ExecuteActivity(ctx, "aws-kendra-DescribeDataSource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeDataSourceAsync(ctx workflow.Context, input *kendra.DescribeDataSourceInput) *DescribeDataSourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws-kendra-DescribeDataSource", input)
	return &DescribeDataSourceFuture{Future: future}
}

func (a *stub) DescribeFaq(ctx workflow.Context, input *kendra.DescribeFaqInput) (*kendra.DescribeFaqOutput, error) {
	var output kendra.DescribeFaqOutput
	err := workflow.ExecuteActivity(ctx, "aws-kendra-DescribeFaq", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeFaqAsync(ctx workflow.Context, input *kendra.DescribeFaqInput) *DescribeFaqFuture {
	future := workflow.ExecuteActivity(ctx, "aws-kendra-DescribeFaq", input)
	return &DescribeFaqFuture{Future: future}
}

func (a *stub) DescribeIndex(ctx workflow.Context, input *kendra.DescribeIndexInput) (*kendra.DescribeIndexOutput, error) {
	var output kendra.DescribeIndexOutput
	err := workflow.ExecuteActivity(ctx, "aws-kendra-DescribeIndex", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeIndexAsync(ctx workflow.Context, input *kendra.DescribeIndexInput) *DescribeIndexFuture {
	future := workflow.ExecuteActivity(ctx, "aws-kendra-DescribeIndex", input)
	return &DescribeIndexFuture{Future: future}
}

func (a *stub) DescribeThesaurus(ctx workflow.Context, input *kendra.DescribeThesaurusInput) (*kendra.DescribeThesaurusOutput, error) {
	var output kendra.DescribeThesaurusOutput
	err := workflow.ExecuteActivity(ctx, "aws-kendra-DescribeThesaurus", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeThesaurusAsync(ctx workflow.Context, input *kendra.DescribeThesaurusInput) *DescribeThesaurusFuture {
	future := workflow.ExecuteActivity(ctx, "aws-kendra-DescribeThesaurus", input)
	return &DescribeThesaurusFuture{Future: future}
}

func (a *stub) ListDataSourceSyncJobs(ctx workflow.Context, input *kendra.ListDataSourceSyncJobsInput) (*kendra.ListDataSourceSyncJobsOutput, error) {
	var output kendra.ListDataSourceSyncJobsOutput
	err := workflow.ExecuteActivity(ctx, "aws-kendra-ListDataSourceSyncJobs", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListDataSourceSyncJobsAsync(ctx workflow.Context, input *kendra.ListDataSourceSyncJobsInput) *ListDataSourceSyncJobsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-kendra-ListDataSourceSyncJobs", input)
	return &ListDataSourceSyncJobsFuture{Future: future}
}

func (a *stub) ListDataSources(ctx workflow.Context, input *kendra.ListDataSourcesInput) (*kendra.ListDataSourcesOutput, error) {
	var output kendra.ListDataSourcesOutput
	err := workflow.ExecuteActivity(ctx, "aws-kendra-ListDataSources", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListDataSourcesAsync(ctx workflow.Context, input *kendra.ListDataSourcesInput) *ListDataSourcesFuture {
	future := workflow.ExecuteActivity(ctx, "aws-kendra-ListDataSources", input)
	return &ListDataSourcesFuture{Future: future}
}

func (a *stub) ListFaqs(ctx workflow.Context, input *kendra.ListFaqsInput) (*kendra.ListFaqsOutput, error) {
	var output kendra.ListFaqsOutput
	err := workflow.ExecuteActivity(ctx, "aws-kendra-ListFaqs", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListFaqsAsync(ctx workflow.Context, input *kendra.ListFaqsInput) *ListFaqsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-kendra-ListFaqs", input)
	return &ListFaqsFuture{Future: future}
}

func (a *stub) ListIndices(ctx workflow.Context, input *kendra.ListIndicesInput) (*kendra.ListIndicesOutput, error) {
	var output kendra.ListIndicesOutput
	err := workflow.ExecuteActivity(ctx, "aws-kendra-ListIndices", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListIndicesAsync(ctx workflow.Context, input *kendra.ListIndicesInput) *ListIndicesFuture {
	future := workflow.ExecuteActivity(ctx, "aws-kendra-ListIndices", input)
	return &ListIndicesFuture{Future: future}
}

func (a *stub) ListTagsForResource(ctx workflow.Context, input *kendra.ListTagsForResourceInput) (*kendra.ListTagsForResourceOutput, error) {
	var output kendra.ListTagsForResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws-kendra-ListTagsForResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListTagsForResourceAsync(ctx workflow.Context, input *kendra.ListTagsForResourceInput) *ListTagsForResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws-kendra-ListTagsForResource", input)
	return &ListTagsForResourceFuture{Future: future}
}

func (a *stub) ListThesauri(ctx workflow.Context, input *kendra.ListThesauriInput) (*kendra.ListThesauriOutput, error) {
	var output kendra.ListThesauriOutput
	err := workflow.ExecuteActivity(ctx, "aws-kendra-ListThesauri", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListThesauriAsync(ctx workflow.Context, input *kendra.ListThesauriInput) *ListThesauriFuture {
	future := workflow.ExecuteActivity(ctx, "aws-kendra-ListThesauri", input)
	return &ListThesauriFuture{Future: future}
}

func (a *stub) Query(ctx workflow.Context, input *kendra.QueryInput) (*kendra.QueryOutput, error) {
	var output kendra.QueryOutput
	err := workflow.ExecuteActivity(ctx, "aws-kendra-Query", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) QueryAsync(ctx workflow.Context, input *kendra.QueryInput) *QueryFuture {
	future := workflow.ExecuteActivity(ctx, "aws-kendra-Query", input)
	return &QueryFuture{Future: future}
}

func (a *stub) StartDataSourceSyncJob(ctx workflow.Context, input *kendra.StartDataSourceSyncJobInput) (*kendra.StartDataSourceSyncJobOutput, error) {
	var output kendra.StartDataSourceSyncJobOutput
	err := workflow.ExecuteActivity(ctx, "aws-kendra-StartDataSourceSyncJob", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) StartDataSourceSyncJobAsync(ctx workflow.Context, input *kendra.StartDataSourceSyncJobInput) *StartDataSourceSyncJobFuture {
	future := workflow.ExecuteActivity(ctx, "aws-kendra-StartDataSourceSyncJob", input)
	return &StartDataSourceSyncJobFuture{Future: future}
}

func (a *stub) StopDataSourceSyncJob(ctx workflow.Context, input *kendra.StopDataSourceSyncJobInput) (*kendra.StopDataSourceSyncJobOutput, error) {
	var output kendra.StopDataSourceSyncJobOutput
	err := workflow.ExecuteActivity(ctx, "aws-kendra-StopDataSourceSyncJob", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) StopDataSourceSyncJobAsync(ctx workflow.Context, input *kendra.StopDataSourceSyncJobInput) *StopDataSourceSyncJobFuture {
	future := workflow.ExecuteActivity(ctx, "aws-kendra-StopDataSourceSyncJob", input)
	return &StopDataSourceSyncJobFuture{Future: future}
}

func (a *stub) SubmitFeedback(ctx workflow.Context, input *kendra.SubmitFeedbackInput) (*kendra.SubmitFeedbackOutput, error) {
	var output kendra.SubmitFeedbackOutput
	err := workflow.ExecuteActivity(ctx, "aws-kendra-SubmitFeedback", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) SubmitFeedbackAsync(ctx workflow.Context, input *kendra.SubmitFeedbackInput) *SubmitFeedbackFuture {
	future := workflow.ExecuteActivity(ctx, "aws-kendra-SubmitFeedback", input)
	return &SubmitFeedbackFuture{Future: future}
}

func (a *stub) TagResource(ctx workflow.Context, input *kendra.TagResourceInput) (*kendra.TagResourceOutput, error) {
	var output kendra.TagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws-kendra-TagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) TagResourceAsync(ctx workflow.Context, input *kendra.TagResourceInput) *TagResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws-kendra-TagResource", input)
	return &TagResourceFuture{Future: future}
}

func (a *stub) UntagResource(ctx workflow.Context, input *kendra.UntagResourceInput) (*kendra.UntagResourceOutput, error) {
	var output kendra.UntagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws-kendra-UntagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UntagResourceAsync(ctx workflow.Context, input *kendra.UntagResourceInput) *UntagResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws-kendra-UntagResource", input)
	return &UntagResourceFuture{Future: future}
}

func (a *stub) UpdateDataSource(ctx workflow.Context, input *kendra.UpdateDataSourceInput) (*kendra.UpdateDataSourceOutput, error) {
	var output kendra.UpdateDataSourceOutput
	err := workflow.ExecuteActivity(ctx, "aws-kendra-UpdateDataSource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateDataSourceAsync(ctx workflow.Context, input *kendra.UpdateDataSourceInput) *UpdateDataSourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws-kendra-UpdateDataSource", input)
	return &UpdateDataSourceFuture{Future: future}
}

func (a *stub) UpdateIndex(ctx workflow.Context, input *kendra.UpdateIndexInput) (*kendra.UpdateIndexOutput, error) {
	var output kendra.UpdateIndexOutput
	err := workflow.ExecuteActivity(ctx, "aws-kendra-UpdateIndex", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateIndexAsync(ctx workflow.Context, input *kendra.UpdateIndexInput) *UpdateIndexFuture {
	future := workflow.ExecuteActivity(ctx, "aws-kendra-UpdateIndex", input)
	return &UpdateIndexFuture{Future: future}
}

func (a *stub) UpdateThesaurus(ctx workflow.Context, input *kendra.UpdateThesaurusInput) (*kendra.UpdateThesaurusOutput, error) {
	var output kendra.UpdateThesaurusOutput
	err := workflow.ExecuteActivity(ctx, "aws-kendra-UpdateThesaurus", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateThesaurusAsync(ctx workflow.Context, input *kendra.UpdateThesaurusInput) *UpdateThesaurusFuture {
	future := workflow.ExecuteActivity(ctx, "aws-kendra-UpdateThesaurus", input)
	return &UpdateThesaurusFuture{Future: future}
}
