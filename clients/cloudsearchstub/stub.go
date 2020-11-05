// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7

package cloudsearchstub

import (
	"github.com/aws/aws-sdk-go/service/cloudsearch"
	"go.uber.org/cadence/workflow"

	"github.com/banzaicloud/cadence-aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type stub struct{}

type CloudSearchBuildSuggestersFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CloudSearchBuildSuggestersFuture) Get(ctx workflow.Context) (*cloudsearch.BuildSuggestersOutput, error) {
	var output cloudsearch.BuildSuggestersOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CloudSearchCreateDomainFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CloudSearchCreateDomainFuture) Get(ctx workflow.Context) (*cloudsearch.CreateDomainOutput, error) {
	var output cloudsearch.CreateDomainOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CloudSearchDefineAnalysisSchemeFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CloudSearchDefineAnalysisSchemeFuture) Get(ctx workflow.Context) (*cloudsearch.DefineAnalysisSchemeOutput, error) {
	var output cloudsearch.DefineAnalysisSchemeOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CloudSearchDefineExpressionFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CloudSearchDefineExpressionFuture) Get(ctx workflow.Context) (*cloudsearch.DefineExpressionOutput, error) {
	var output cloudsearch.DefineExpressionOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CloudSearchDefineIndexFieldFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CloudSearchDefineIndexFieldFuture) Get(ctx workflow.Context) (*cloudsearch.DefineIndexFieldOutput, error) {
	var output cloudsearch.DefineIndexFieldOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CloudSearchDefineSuggesterFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CloudSearchDefineSuggesterFuture) Get(ctx workflow.Context) (*cloudsearch.DefineSuggesterOutput, error) {
	var output cloudsearch.DefineSuggesterOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CloudSearchDeleteAnalysisSchemeFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CloudSearchDeleteAnalysisSchemeFuture) Get(ctx workflow.Context) (*cloudsearch.DeleteAnalysisSchemeOutput, error) {
	var output cloudsearch.DeleteAnalysisSchemeOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CloudSearchDeleteDomainFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CloudSearchDeleteDomainFuture) Get(ctx workflow.Context) (*cloudsearch.DeleteDomainOutput, error) {
	var output cloudsearch.DeleteDomainOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CloudSearchDeleteExpressionFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CloudSearchDeleteExpressionFuture) Get(ctx workflow.Context) (*cloudsearch.DeleteExpressionOutput, error) {
	var output cloudsearch.DeleteExpressionOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CloudSearchDeleteIndexFieldFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CloudSearchDeleteIndexFieldFuture) Get(ctx workflow.Context) (*cloudsearch.DeleteIndexFieldOutput, error) {
	var output cloudsearch.DeleteIndexFieldOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CloudSearchDeleteSuggesterFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CloudSearchDeleteSuggesterFuture) Get(ctx workflow.Context) (*cloudsearch.DeleteSuggesterOutput, error) {
	var output cloudsearch.DeleteSuggesterOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CloudSearchDescribeAnalysisSchemesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CloudSearchDescribeAnalysisSchemesFuture) Get(ctx workflow.Context) (*cloudsearch.DescribeAnalysisSchemesOutput, error) {
	var output cloudsearch.DescribeAnalysisSchemesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CloudSearchDescribeAvailabilityOptionsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CloudSearchDescribeAvailabilityOptionsFuture) Get(ctx workflow.Context) (*cloudsearch.DescribeAvailabilityOptionsOutput, error) {
	var output cloudsearch.DescribeAvailabilityOptionsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CloudSearchDescribeDomainEndpointOptionsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CloudSearchDescribeDomainEndpointOptionsFuture) Get(ctx workflow.Context) (*cloudsearch.DescribeDomainEndpointOptionsOutput, error) {
	var output cloudsearch.DescribeDomainEndpointOptionsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CloudSearchDescribeDomainsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CloudSearchDescribeDomainsFuture) Get(ctx workflow.Context) (*cloudsearch.DescribeDomainsOutput, error) {
	var output cloudsearch.DescribeDomainsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CloudSearchDescribeExpressionsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CloudSearchDescribeExpressionsFuture) Get(ctx workflow.Context) (*cloudsearch.DescribeExpressionsOutput, error) {
	var output cloudsearch.DescribeExpressionsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CloudSearchDescribeIndexFieldsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CloudSearchDescribeIndexFieldsFuture) Get(ctx workflow.Context) (*cloudsearch.DescribeIndexFieldsOutput, error) {
	var output cloudsearch.DescribeIndexFieldsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CloudSearchDescribeScalingParametersFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CloudSearchDescribeScalingParametersFuture) Get(ctx workflow.Context) (*cloudsearch.DescribeScalingParametersOutput, error) {
	var output cloudsearch.DescribeScalingParametersOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CloudSearchDescribeServiceAccessPoliciesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CloudSearchDescribeServiceAccessPoliciesFuture) Get(ctx workflow.Context) (*cloudsearch.DescribeServiceAccessPoliciesOutput, error) {
	var output cloudsearch.DescribeServiceAccessPoliciesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CloudSearchDescribeSuggestersFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CloudSearchDescribeSuggestersFuture) Get(ctx workflow.Context) (*cloudsearch.DescribeSuggestersOutput, error) {
	var output cloudsearch.DescribeSuggestersOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CloudSearchIndexDocumentsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CloudSearchIndexDocumentsFuture) Get(ctx workflow.Context) (*cloudsearch.IndexDocumentsOutput, error) {
	var output cloudsearch.IndexDocumentsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CloudSearchListDomainNamesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CloudSearchListDomainNamesFuture) Get(ctx workflow.Context) (*cloudsearch.ListDomainNamesOutput, error) {
	var output cloudsearch.ListDomainNamesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CloudSearchUpdateAvailabilityOptionsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CloudSearchUpdateAvailabilityOptionsFuture) Get(ctx workflow.Context) (*cloudsearch.UpdateAvailabilityOptionsOutput, error) {
	var output cloudsearch.UpdateAvailabilityOptionsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CloudSearchUpdateDomainEndpointOptionsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CloudSearchUpdateDomainEndpointOptionsFuture) Get(ctx workflow.Context) (*cloudsearch.UpdateDomainEndpointOptionsOutput, error) {
	var output cloudsearch.UpdateDomainEndpointOptionsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CloudSearchUpdateScalingParametersFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CloudSearchUpdateScalingParametersFuture) Get(ctx workflow.Context) (*cloudsearch.UpdateScalingParametersOutput, error) {
	var output cloudsearch.UpdateScalingParametersOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CloudSearchUpdateServiceAccessPoliciesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CloudSearchUpdateServiceAccessPoliciesFuture) Get(ctx workflow.Context) (*cloudsearch.UpdateServiceAccessPoliciesOutput, error) {
	var output cloudsearch.UpdateServiceAccessPoliciesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *stub) BuildSuggesters(ctx workflow.Context, input *cloudsearch.BuildSuggestersInput) (*cloudsearch.BuildSuggestersOutput, error) {
	var output cloudsearch.BuildSuggestersOutput
	err := workflow.ExecuteActivity(ctx, "aws-cloudsearch-BuildSuggesters", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) BuildSuggestersAsync(ctx workflow.Context, input *cloudsearch.BuildSuggestersInput) *CloudSearchBuildSuggestersFuture {
	future := workflow.ExecuteActivity(ctx, "aws-cloudsearch-BuildSuggesters", input)
	return &CloudSearchBuildSuggestersFuture{Future: future}
}

func (a *stub) CreateDomain(ctx workflow.Context, input *cloudsearch.CreateDomainInput) (*cloudsearch.CreateDomainOutput, error) {
	var output cloudsearch.CreateDomainOutput
	err := workflow.ExecuteActivity(ctx, "aws-cloudsearch-CreateDomain", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateDomainAsync(ctx workflow.Context, input *cloudsearch.CreateDomainInput) *CloudSearchCreateDomainFuture {
	future := workflow.ExecuteActivity(ctx, "aws-cloudsearch-CreateDomain", input)
	return &CloudSearchCreateDomainFuture{Future: future}
}

func (a *stub) DefineAnalysisScheme(ctx workflow.Context, input *cloudsearch.DefineAnalysisSchemeInput) (*cloudsearch.DefineAnalysisSchemeOutput, error) {
	var output cloudsearch.DefineAnalysisSchemeOutput
	err := workflow.ExecuteActivity(ctx, "aws-cloudsearch-DefineAnalysisScheme", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DefineAnalysisSchemeAsync(ctx workflow.Context, input *cloudsearch.DefineAnalysisSchemeInput) *CloudSearchDefineAnalysisSchemeFuture {
	future := workflow.ExecuteActivity(ctx, "aws-cloudsearch-DefineAnalysisScheme", input)
	return &CloudSearchDefineAnalysisSchemeFuture{Future: future}
}

func (a *stub) DefineExpression(ctx workflow.Context, input *cloudsearch.DefineExpressionInput) (*cloudsearch.DefineExpressionOutput, error) {
	var output cloudsearch.DefineExpressionOutput
	err := workflow.ExecuteActivity(ctx, "aws-cloudsearch-DefineExpression", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DefineExpressionAsync(ctx workflow.Context, input *cloudsearch.DefineExpressionInput) *CloudSearchDefineExpressionFuture {
	future := workflow.ExecuteActivity(ctx, "aws-cloudsearch-DefineExpression", input)
	return &CloudSearchDefineExpressionFuture{Future: future}
}

func (a *stub) DefineIndexField(ctx workflow.Context, input *cloudsearch.DefineIndexFieldInput) (*cloudsearch.DefineIndexFieldOutput, error) {
	var output cloudsearch.DefineIndexFieldOutput
	err := workflow.ExecuteActivity(ctx, "aws-cloudsearch-DefineIndexField", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DefineIndexFieldAsync(ctx workflow.Context, input *cloudsearch.DefineIndexFieldInput) *CloudSearchDefineIndexFieldFuture {
	future := workflow.ExecuteActivity(ctx, "aws-cloudsearch-DefineIndexField", input)
	return &CloudSearchDefineIndexFieldFuture{Future: future}
}

func (a *stub) DefineSuggester(ctx workflow.Context, input *cloudsearch.DefineSuggesterInput) (*cloudsearch.DefineSuggesterOutput, error) {
	var output cloudsearch.DefineSuggesterOutput
	err := workflow.ExecuteActivity(ctx, "aws-cloudsearch-DefineSuggester", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DefineSuggesterAsync(ctx workflow.Context, input *cloudsearch.DefineSuggesterInput) *CloudSearchDefineSuggesterFuture {
	future := workflow.ExecuteActivity(ctx, "aws-cloudsearch-DefineSuggester", input)
	return &CloudSearchDefineSuggesterFuture{Future: future}
}

func (a *stub) DeleteAnalysisScheme(ctx workflow.Context, input *cloudsearch.DeleteAnalysisSchemeInput) (*cloudsearch.DeleteAnalysisSchemeOutput, error) {
	var output cloudsearch.DeleteAnalysisSchemeOutput
	err := workflow.ExecuteActivity(ctx, "aws-cloudsearch-DeleteAnalysisScheme", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteAnalysisSchemeAsync(ctx workflow.Context, input *cloudsearch.DeleteAnalysisSchemeInput) *CloudSearchDeleteAnalysisSchemeFuture {
	future := workflow.ExecuteActivity(ctx, "aws-cloudsearch-DeleteAnalysisScheme", input)
	return &CloudSearchDeleteAnalysisSchemeFuture{Future: future}
}

func (a *stub) DeleteDomain(ctx workflow.Context, input *cloudsearch.DeleteDomainInput) (*cloudsearch.DeleteDomainOutput, error) {
	var output cloudsearch.DeleteDomainOutput
	err := workflow.ExecuteActivity(ctx, "aws-cloudsearch-DeleteDomain", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteDomainAsync(ctx workflow.Context, input *cloudsearch.DeleteDomainInput) *CloudSearchDeleteDomainFuture {
	future := workflow.ExecuteActivity(ctx, "aws-cloudsearch-DeleteDomain", input)
	return &CloudSearchDeleteDomainFuture{Future: future}
}

func (a *stub) DeleteExpression(ctx workflow.Context, input *cloudsearch.DeleteExpressionInput) (*cloudsearch.DeleteExpressionOutput, error) {
	var output cloudsearch.DeleteExpressionOutput
	err := workflow.ExecuteActivity(ctx, "aws-cloudsearch-DeleteExpression", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteExpressionAsync(ctx workflow.Context, input *cloudsearch.DeleteExpressionInput) *CloudSearchDeleteExpressionFuture {
	future := workflow.ExecuteActivity(ctx, "aws-cloudsearch-DeleteExpression", input)
	return &CloudSearchDeleteExpressionFuture{Future: future}
}

func (a *stub) DeleteIndexField(ctx workflow.Context, input *cloudsearch.DeleteIndexFieldInput) (*cloudsearch.DeleteIndexFieldOutput, error) {
	var output cloudsearch.DeleteIndexFieldOutput
	err := workflow.ExecuteActivity(ctx, "aws-cloudsearch-DeleteIndexField", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteIndexFieldAsync(ctx workflow.Context, input *cloudsearch.DeleteIndexFieldInput) *CloudSearchDeleteIndexFieldFuture {
	future := workflow.ExecuteActivity(ctx, "aws-cloudsearch-DeleteIndexField", input)
	return &CloudSearchDeleteIndexFieldFuture{Future: future}
}

func (a *stub) DeleteSuggester(ctx workflow.Context, input *cloudsearch.DeleteSuggesterInput) (*cloudsearch.DeleteSuggesterOutput, error) {
	var output cloudsearch.DeleteSuggesterOutput
	err := workflow.ExecuteActivity(ctx, "aws-cloudsearch-DeleteSuggester", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteSuggesterAsync(ctx workflow.Context, input *cloudsearch.DeleteSuggesterInput) *CloudSearchDeleteSuggesterFuture {
	future := workflow.ExecuteActivity(ctx, "aws-cloudsearch-DeleteSuggester", input)
	return &CloudSearchDeleteSuggesterFuture{Future: future}
}

func (a *stub) DescribeAnalysisSchemes(ctx workflow.Context, input *cloudsearch.DescribeAnalysisSchemesInput) (*cloudsearch.DescribeAnalysisSchemesOutput, error) {
	var output cloudsearch.DescribeAnalysisSchemesOutput
	err := workflow.ExecuteActivity(ctx, "aws-cloudsearch-DescribeAnalysisSchemes", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeAnalysisSchemesAsync(ctx workflow.Context, input *cloudsearch.DescribeAnalysisSchemesInput) *CloudSearchDescribeAnalysisSchemesFuture {
	future := workflow.ExecuteActivity(ctx, "aws-cloudsearch-DescribeAnalysisSchemes", input)
	return &CloudSearchDescribeAnalysisSchemesFuture{Future: future}
}

func (a *stub) DescribeAvailabilityOptions(ctx workflow.Context, input *cloudsearch.DescribeAvailabilityOptionsInput) (*cloudsearch.DescribeAvailabilityOptionsOutput, error) {
	var output cloudsearch.DescribeAvailabilityOptionsOutput
	err := workflow.ExecuteActivity(ctx, "aws-cloudsearch-DescribeAvailabilityOptions", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeAvailabilityOptionsAsync(ctx workflow.Context, input *cloudsearch.DescribeAvailabilityOptionsInput) *CloudSearchDescribeAvailabilityOptionsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-cloudsearch-DescribeAvailabilityOptions", input)
	return &CloudSearchDescribeAvailabilityOptionsFuture{Future: future}
}

func (a *stub) DescribeDomainEndpointOptions(ctx workflow.Context, input *cloudsearch.DescribeDomainEndpointOptionsInput) (*cloudsearch.DescribeDomainEndpointOptionsOutput, error) {
	var output cloudsearch.DescribeDomainEndpointOptionsOutput
	err := workflow.ExecuteActivity(ctx, "aws-cloudsearch-DescribeDomainEndpointOptions", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeDomainEndpointOptionsAsync(ctx workflow.Context, input *cloudsearch.DescribeDomainEndpointOptionsInput) *CloudSearchDescribeDomainEndpointOptionsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-cloudsearch-DescribeDomainEndpointOptions", input)
	return &CloudSearchDescribeDomainEndpointOptionsFuture{Future: future}
}

func (a *stub) DescribeDomains(ctx workflow.Context, input *cloudsearch.DescribeDomainsInput) (*cloudsearch.DescribeDomainsOutput, error) {
	var output cloudsearch.DescribeDomainsOutput
	err := workflow.ExecuteActivity(ctx, "aws-cloudsearch-DescribeDomains", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeDomainsAsync(ctx workflow.Context, input *cloudsearch.DescribeDomainsInput) *CloudSearchDescribeDomainsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-cloudsearch-DescribeDomains", input)
	return &CloudSearchDescribeDomainsFuture{Future: future}
}

func (a *stub) DescribeExpressions(ctx workflow.Context, input *cloudsearch.DescribeExpressionsInput) (*cloudsearch.DescribeExpressionsOutput, error) {
	var output cloudsearch.DescribeExpressionsOutput
	err := workflow.ExecuteActivity(ctx, "aws-cloudsearch-DescribeExpressions", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeExpressionsAsync(ctx workflow.Context, input *cloudsearch.DescribeExpressionsInput) *CloudSearchDescribeExpressionsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-cloudsearch-DescribeExpressions", input)
	return &CloudSearchDescribeExpressionsFuture{Future: future}
}

func (a *stub) DescribeIndexFields(ctx workflow.Context, input *cloudsearch.DescribeIndexFieldsInput) (*cloudsearch.DescribeIndexFieldsOutput, error) {
	var output cloudsearch.DescribeIndexFieldsOutput
	err := workflow.ExecuteActivity(ctx, "aws-cloudsearch-DescribeIndexFields", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeIndexFieldsAsync(ctx workflow.Context, input *cloudsearch.DescribeIndexFieldsInput) *CloudSearchDescribeIndexFieldsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-cloudsearch-DescribeIndexFields", input)
	return &CloudSearchDescribeIndexFieldsFuture{Future: future}
}

func (a *stub) DescribeScalingParameters(ctx workflow.Context, input *cloudsearch.DescribeScalingParametersInput) (*cloudsearch.DescribeScalingParametersOutput, error) {
	var output cloudsearch.DescribeScalingParametersOutput
	err := workflow.ExecuteActivity(ctx, "aws-cloudsearch-DescribeScalingParameters", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeScalingParametersAsync(ctx workflow.Context, input *cloudsearch.DescribeScalingParametersInput) *CloudSearchDescribeScalingParametersFuture {
	future := workflow.ExecuteActivity(ctx, "aws-cloudsearch-DescribeScalingParameters", input)
	return &CloudSearchDescribeScalingParametersFuture{Future: future}
}

func (a *stub) DescribeServiceAccessPolicies(ctx workflow.Context, input *cloudsearch.DescribeServiceAccessPoliciesInput) (*cloudsearch.DescribeServiceAccessPoliciesOutput, error) {
	var output cloudsearch.DescribeServiceAccessPoliciesOutput
	err := workflow.ExecuteActivity(ctx, "aws-cloudsearch-DescribeServiceAccessPolicies", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeServiceAccessPoliciesAsync(ctx workflow.Context, input *cloudsearch.DescribeServiceAccessPoliciesInput) *CloudSearchDescribeServiceAccessPoliciesFuture {
	future := workflow.ExecuteActivity(ctx, "aws-cloudsearch-DescribeServiceAccessPolicies", input)
	return &CloudSearchDescribeServiceAccessPoliciesFuture{Future: future}
}

func (a *stub) DescribeSuggesters(ctx workflow.Context, input *cloudsearch.DescribeSuggestersInput) (*cloudsearch.DescribeSuggestersOutput, error) {
	var output cloudsearch.DescribeSuggestersOutput
	err := workflow.ExecuteActivity(ctx, "aws-cloudsearch-DescribeSuggesters", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeSuggestersAsync(ctx workflow.Context, input *cloudsearch.DescribeSuggestersInput) *CloudSearchDescribeSuggestersFuture {
	future := workflow.ExecuteActivity(ctx, "aws-cloudsearch-DescribeSuggesters", input)
	return &CloudSearchDescribeSuggestersFuture{Future: future}
}

func (a *stub) IndexDocuments(ctx workflow.Context, input *cloudsearch.IndexDocumentsInput) (*cloudsearch.IndexDocumentsOutput, error) {
	var output cloudsearch.IndexDocumentsOutput
	err := workflow.ExecuteActivity(ctx, "aws-cloudsearch-IndexDocuments", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) IndexDocumentsAsync(ctx workflow.Context, input *cloudsearch.IndexDocumentsInput) *CloudSearchIndexDocumentsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-cloudsearch-IndexDocuments", input)
	return &CloudSearchIndexDocumentsFuture{Future: future}
}

func (a *stub) ListDomainNames(ctx workflow.Context, input *cloudsearch.ListDomainNamesInput) (*cloudsearch.ListDomainNamesOutput, error) {
	var output cloudsearch.ListDomainNamesOutput
	err := workflow.ExecuteActivity(ctx, "aws-cloudsearch-ListDomainNames", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListDomainNamesAsync(ctx workflow.Context, input *cloudsearch.ListDomainNamesInput) *CloudSearchListDomainNamesFuture {
	future := workflow.ExecuteActivity(ctx, "aws-cloudsearch-ListDomainNames", input)
	return &CloudSearchListDomainNamesFuture{Future: future}
}

func (a *stub) UpdateAvailabilityOptions(ctx workflow.Context, input *cloudsearch.UpdateAvailabilityOptionsInput) (*cloudsearch.UpdateAvailabilityOptionsOutput, error) {
	var output cloudsearch.UpdateAvailabilityOptionsOutput
	err := workflow.ExecuteActivity(ctx, "aws-cloudsearch-UpdateAvailabilityOptions", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateAvailabilityOptionsAsync(ctx workflow.Context, input *cloudsearch.UpdateAvailabilityOptionsInput) *CloudSearchUpdateAvailabilityOptionsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-cloudsearch-UpdateAvailabilityOptions", input)
	return &CloudSearchUpdateAvailabilityOptionsFuture{Future: future}
}

func (a *stub) UpdateDomainEndpointOptions(ctx workflow.Context, input *cloudsearch.UpdateDomainEndpointOptionsInput) (*cloudsearch.UpdateDomainEndpointOptionsOutput, error) {
	var output cloudsearch.UpdateDomainEndpointOptionsOutput
	err := workflow.ExecuteActivity(ctx, "aws-cloudsearch-UpdateDomainEndpointOptions", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateDomainEndpointOptionsAsync(ctx workflow.Context, input *cloudsearch.UpdateDomainEndpointOptionsInput) *CloudSearchUpdateDomainEndpointOptionsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-cloudsearch-UpdateDomainEndpointOptions", input)
	return &CloudSearchUpdateDomainEndpointOptionsFuture{Future: future}
}

func (a *stub) UpdateScalingParameters(ctx workflow.Context, input *cloudsearch.UpdateScalingParametersInput) (*cloudsearch.UpdateScalingParametersOutput, error) {
	var output cloudsearch.UpdateScalingParametersOutput
	err := workflow.ExecuteActivity(ctx, "aws-cloudsearch-UpdateScalingParameters", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateScalingParametersAsync(ctx workflow.Context, input *cloudsearch.UpdateScalingParametersInput) *CloudSearchUpdateScalingParametersFuture {
	future := workflow.ExecuteActivity(ctx, "aws-cloudsearch-UpdateScalingParameters", input)
	return &CloudSearchUpdateScalingParametersFuture{Future: future}
}

func (a *stub) UpdateServiceAccessPolicies(ctx workflow.Context, input *cloudsearch.UpdateServiceAccessPoliciesInput) (*cloudsearch.UpdateServiceAccessPoliciesOutput, error) {
	var output cloudsearch.UpdateServiceAccessPoliciesOutput
	err := workflow.ExecuteActivity(ctx, "aws-cloudsearch-UpdateServiceAccessPolicies", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateServiceAccessPoliciesAsync(ctx workflow.Context, input *cloudsearch.UpdateServiceAccessPoliciesInput) *CloudSearchUpdateServiceAccessPoliciesFuture {
	future := workflow.ExecuteActivity(ctx, "aws-cloudsearch-UpdateServiceAccessPolicies", input)
	return &CloudSearchUpdateServiceAccessPoliciesFuture{Future: future}
}
