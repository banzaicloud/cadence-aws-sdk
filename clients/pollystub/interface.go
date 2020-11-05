// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7

package pollystub

import (
	"github.com/aws/aws-sdk-go/service/polly"
	"go.uber.org/cadence/workflow"

	"github.com/banzaicloud/cadence-aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	DeleteLexicon(ctx workflow.Context, input *polly.DeleteLexiconInput) (*polly.DeleteLexiconOutput, error)
	DeleteLexiconAsync(ctx workflow.Context, input *polly.DeleteLexiconInput) *PollyDeleteLexiconFuture

	DescribeVoices(ctx workflow.Context, input *polly.DescribeVoicesInput) (*polly.DescribeVoicesOutput, error)
	DescribeVoicesAsync(ctx workflow.Context, input *polly.DescribeVoicesInput) *PollyDescribeVoicesFuture

	GetLexicon(ctx workflow.Context, input *polly.GetLexiconInput) (*polly.GetLexiconOutput, error)
	GetLexiconAsync(ctx workflow.Context, input *polly.GetLexiconInput) *PollyGetLexiconFuture

	GetSpeechSynthesisTask(ctx workflow.Context, input *polly.GetSpeechSynthesisTaskInput) (*polly.GetSpeechSynthesisTaskOutput, error)
	GetSpeechSynthesisTaskAsync(ctx workflow.Context, input *polly.GetSpeechSynthesisTaskInput) *PollyGetSpeechSynthesisTaskFuture

	ListLexicons(ctx workflow.Context, input *polly.ListLexiconsInput) (*polly.ListLexiconsOutput, error)
	ListLexiconsAsync(ctx workflow.Context, input *polly.ListLexiconsInput) *PollyListLexiconsFuture

	ListSpeechSynthesisTasks(ctx workflow.Context, input *polly.ListSpeechSynthesisTasksInput) (*polly.ListSpeechSynthesisTasksOutput, error)
	ListSpeechSynthesisTasksAsync(ctx workflow.Context, input *polly.ListSpeechSynthesisTasksInput) *PollyListSpeechSynthesisTasksFuture

	PutLexicon(ctx workflow.Context, input *polly.PutLexiconInput) (*polly.PutLexiconOutput, error)
	PutLexiconAsync(ctx workflow.Context, input *polly.PutLexiconInput) *PollyPutLexiconFuture

	StartSpeechSynthesisTask(ctx workflow.Context, input *polly.StartSpeechSynthesisTaskInput) (*polly.StartSpeechSynthesisTaskOutput, error)
	StartSpeechSynthesisTaskAsync(ctx workflow.Context, input *polly.StartSpeechSynthesisTaskInput) *PollyStartSpeechSynthesisTaskFuture

	SynthesizeSpeech(ctx workflow.Context, input *polly.SynthesizeSpeechInput) (*polly.SynthesizeSpeechOutput, error)
	SynthesizeSpeechAsync(ctx workflow.Context, input *polly.SynthesizeSpeechInput) *PollySynthesizeSpeechFuture
}

func NewClient() Client {
	return &stub{}
}
