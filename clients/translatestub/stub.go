// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7

package translatestub

import (
	"github.com/aws/aws-sdk-go/service/translate"
	"go.uber.org/cadence/workflow"

	"github.com/banzaicloud/cadence-aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type stub struct{}

type TranslateDeleteTerminologyFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *TranslateDeleteTerminologyFuture) Get(ctx workflow.Context) (*translate.DeleteTerminologyOutput, error) {
	var output translate.DeleteTerminologyOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type TranslateDescribeTextTranslationJobFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *TranslateDescribeTextTranslationJobFuture) Get(ctx workflow.Context) (*translate.DescribeTextTranslationJobOutput, error) {
	var output translate.DescribeTextTranslationJobOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type TranslateGetTerminologyFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *TranslateGetTerminologyFuture) Get(ctx workflow.Context) (*translate.GetTerminologyOutput, error) {
	var output translate.GetTerminologyOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type TranslateImportTerminologyFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *TranslateImportTerminologyFuture) Get(ctx workflow.Context) (*translate.ImportTerminologyOutput, error) {
	var output translate.ImportTerminologyOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type TranslateListTerminologiesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *TranslateListTerminologiesFuture) Get(ctx workflow.Context) (*translate.ListTerminologiesOutput, error) {
	var output translate.ListTerminologiesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type TranslateListTextTranslationJobsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *TranslateListTextTranslationJobsFuture) Get(ctx workflow.Context) (*translate.ListTextTranslationJobsOutput, error) {
	var output translate.ListTextTranslationJobsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type TranslateStartTextTranslationJobFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *TranslateStartTextTranslationJobFuture) Get(ctx workflow.Context) (*translate.StartTextTranslationJobOutput, error) {
	var output translate.StartTextTranslationJobOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type TranslateStopTextTranslationJobFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *TranslateStopTextTranslationJobFuture) Get(ctx workflow.Context) (*translate.StopTextTranslationJobOutput, error) {
	var output translate.StopTextTranslationJobOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type TranslateTextFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *TranslateTextFuture) Get(ctx workflow.Context) (*translate.TextOutput, error) {
	var output translate.TextOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteTerminology(ctx workflow.Context, input *translate.DeleteTerminologyInput) (*translate.DeleteTerminologyOutput, error) {
	var output translate.DeleteTerminologyOutput
	err := workflow.ExecuteActivity(ctx, "aws-translate-DeleteTerminology", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteTerminologyAsync(ctx workflow.Context, input *translate.DeleteTerminologyInput) *TranslateDeleteTerminologyFuture {
	future := workflow.ExecuteActivity(ctx, "aws-translate-DeleteTerminology", input)
	return &TranslateDeleteTerminologyFuture{Future: future}
}

func (a *stub) DescribeTextTranslationJob(ctx workflow.Context, input *translate.DescribeTextTranslationJobInput) (*translate.DescribeTextTranslationJobOutput, error) {
	var output translate.DescribeTextTranslationJobOutput
	err := workflow.ExecuteActivity(ctx, "aws-translate-DescribeTextTranslationJob", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeTextTranslationJobAsync(ctx workflow.Context, input *translate.DescribeTextTranslationJobInput) *TranslateDescribeTextTranslationJobFuture {
	future := workflow.ExecuteActivity(ctx, "aws-translate-DescribeTextTranslationJob", input)
	return &TranslateDescribeTextTranslationJobFuture{Future: future}
}

func (a *stub) GetTerminology(ctx workflow.Context, input *translate.GetTerminologyInput) (*translate.GetTerminologyOutput, error) {
	var output translate.GetTerminologyOutput
	err := workflow.ExecuteActivity(ctx, "aws-translate-GetTerminology", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetTerminologyAsync(ctx workflow.Context, input *translate.GetTerminologyInput) *TranslateGetTerminologyFuture {
	future := workflow.ExecuteActivity(ctx, "aws-translate-GetTerminology", input)
	return &TranslateGetTerminologyFuture{Future: future}
}

func (a *stub) ImportTerminology(ctx workflow.Context, input *translate.ImportTerminologyInput) (*translate.ImportTerminologyOutput, error) {
	var output translate.ImportTerminologyOutput
	err := workflow.ExecuteActivity(ctx, "aws-translate-ImportTerminology", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ImportTerminologyAsync(ctx workflow.Context, input *translate.ImportTerminologyInput) *TranslateImportTerminologyFuture {
	future := workflow.ExecuteActivity(ctx, "aws-translate-ImportTerminology", input)
	return &TranslateImportTerminologyFuture{Future: future}
}

func (a *stub) ListTerminologies(ctx workflow.Context, input *translate.ListTerminologiesInput) (*translate.ListTerminologiesOutput, error) {
	var output translate.ListTerminologiesOutput
	err := workflow.ExecuteActivity(ctx, "aws-translate-ListTerminologies", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListTerminologiesAsync(ctx workflow.Context, input *translate.ListTerminologiesInput) *TranslateListTerminologiesFuture {
	future := workflow.ExecuteActivity(ctx, "aws-translate-ListTerminologies", input)
	return &TranslateListTerminologiesFuture{Future: future}
}

func (a *stub) ListTextTranslationJobs(ctx workflow.Context, input *translate.ListTextTranslationJobsInput) (*translate.ListTextTranslationJobsOutput, error) {
	var output translate.ListTextTranslationJobsOutput
	err := workflow.ExecuteActivity(ctx, "aws-translate-ListTextTranslationJobs", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListTextTranslationJobsAsync(ctx workflow.Context, input *translate.ListTextTranslationJobsInput) *TranslateListTextTranslationJobsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-translate-ListTextTranslationJobs", input)
	return &TranslateListTextTranslationJobsFuture{Future: future}
}

func (a *stub) StartTextTranslationJob(ctx workflow.Context, input *translate.StartTextTranslationJobInput) (*translate.StartTextTranslationJobOutput, error) {
	var output translate.StartTextTranslationJobOutput
	err := workflow.ExecuteActivity(ctx, "aws-translate-StartTextTranslationJob", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) StartTextTranslationJobAsync(ctx workflow.Context, input *translate.StartTextTranslationJobInput) *TranslateStartTextTranslationJobFuture {
	future := workflow.ExecuteActivity(ctx, "aws-translate-StartTextTranslationJob", input)
	return &TranslateStartTextTranslationJobFuture{Future: future}
}

func (a *stub) StopTextTranslationJob(ctx workflow.Context, input *translate.StopTextTranslationJobInput) (*translate.StopTextTranslationJobOutput, error) {
	var output translate.StopTextTranslationJobOutput
	err := workflow.ExecuteActivity(ctx, "aws-translate-StopTextTranslationJob", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) StopTextTranslationJobAsync(ctx workflow.Context, input *translate.StopTextTranslationJobInput) *TranslateStopTextTranslationJobFuture {
	future := workflow.ExecuteActivity(ctx, "aws-translate-StopTextTranslationJob", input)
	return &TranslateStopTextTranslationJobFuture{Future: future}
}

func (a *stub) Text(ctx workflow.Context, input *translate.TextInput) (*translate.TextOutput, error) {
	var output translate.TextOutput
	err := workflow.ExecuteActivity(ctx, "aws-translate-Text", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) TextAsync(ctx workflow.Context, input *translate.TextInput) *TranslateTextFuture {
	future := workflow.ExecuteActivity(ctx, "aws-translate-Text", input)
	return &TranslateTextFuture{Future: future}
}
