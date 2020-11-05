// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7

package costandusagereportservicestub

import (
	"github.com/aws/aws-sdk-go/service/costandusagereportservice"
	"go.uber.org/cadence/workflow"

	"github.com/banzaicloud/cadence-aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type stub struct{}

type CostandUsageReportServiceDeleteReportDefinitionFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CostandUsageReportServiceDeleteReportDefinitionFuture) Get(ctx workflow.Context) (*costandusagereportservice.DeleteReportDefinitionOutput, error) {
	var output costandusagereportservice.DeleteReportDefinitionOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CostandUsageReportServiceDescribeReportDefinitionsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CostandUsageReportServiceDescribeReportDefinitionsFuture) Get(ctx workflow.Context) (*costandusagereportservice.DescribeReportDefinitionsOutput, error) {
	var output costandusagereportservice.DescribeReportDefinitionsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CostandUsageReportServiceModifyReportDefinitionFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CostandUsageReportServiceModifyReportDefinitionFuture) Get(ctx workflow.Context) (*costandusagereportservice.ModifyReportDefinitionOutput, error) {
	var output costandusagereportservice.ModifyReportDefinitionOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CostandUsageReportServicePutReportDefinitionFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CostandUsageReportServicePutReportDefinitionFuture) Get(ctx workflow.Context) (*costandusagereportservice.PutReportDefinitionOutput, error) {
	var output costandusagereportservice.PutReportDefinitionOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteReportDefinition(ctx workflow.Context, input *costandusagereportservice.DeleteReportDefinitionInput) (*costandusagereportservice.DeleteReportDefinitionOutput, error) {
	var output costandusagereportservice.DeleteReportDefinitionOutput
	err := workflow.ExecuteActivity(ctx, "aws-costandusagereportservice-DeleteReportDefinition", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteReportDefinitionAsync(ctx workflow.Context, input *costandusagereportservice.DeleteReportDefinitionInput) *CostandUsageReportServiceDeleteReportDefinitionFuture {
	future := workflow.ExecuteActivity(ctx, "aws-costandusagereportservice-DeleteReportDefinition", input)
	return &CostandUsageReportServiceDeleteReportDefinitionFuture{Future: future}
}

func (a *stub) DescribeReportDefinitions(ctx workflow.Context, input *costandusagereportservice.DescribeReportDefinitionsInput) (*costandusagereportservice.DescribeReportDefinitionsOutput, error) {
	var output costandusagereportservice.DescribeReportDefinitionsOutput
	err := workflow.ExecuteActivity(ctx, "aws-costandusagereportservice-DescribeReportDefinitions", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeReportDefinitionsAsync(ctx workflow.Context, input *costandusagereportservice.DescribeReportDefinitionsInput) *CostandUsageReportServiceDescribeReportDefinitionsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-costandusagereportservice-DescribeReportDefinitions", input)
	return &CostandUsageReportServiceDescribeReportDefinitionsFuture{Future: future}
}

func (a *stub) ModifyReportDefinition(ctx workflow.Context, input *costandusagereportservice.ModifyReportDefinitionInput) (*costandusagereportservice.ModifyReportDefinitionOutput, error) {
	var output costandusagereportservice.ModifyReportDefinitionOutput
	err := workflow.ExecuteActivity(ctx, "aws-costandusagereportservice-ModifyReportDefinition", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ModifyReportDefinitionAsync(ctx workflow.Context, input *costandusagereportservice.ModifyReportDefinitionInput) *CostandUsageReportServiceModifyReportDefinitionFuture {
	future := workflow.ExecuteActivity(ctx, "aws-costandusagereportservice-ModifyReportDefinition", input)
	return &CostandUsageReportServiceModifyReportDefinitionFuture{Future: future}
}

func (a *stub) PutReportDefinition(ctx workflow.Context, input *costandusagereportservice.PutReportDefinitionInput) (*costandusagereportservice.PutReportDefinitionOutput, error) {
	var output costandusagereportservice.PutReportDefinitionOutput
	err := workflow.ExecuteActivity(ctx, "aws-costandusagereportservice-PutReportDefinition", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) PutReportDefinitionAsync(ctx workflow.Context, input *costandusagereportservice.PutReportDefinitionInput) *CostandUsageReportServicePutReportDefinitionFuture {
	future := workflow.ExecuteActivity(ctx, "aws-costandusagereportservice-PutReportDefinition", input)
	return &CostandUsageReportServicePutReportDefinitionFuture{Future: future}
}
