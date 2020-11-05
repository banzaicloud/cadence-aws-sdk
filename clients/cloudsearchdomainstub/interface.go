// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7

package cloudsearchdomainstub

import (
	"github.com/aws/aws-sdk-go/service/cloudsearchdomain"
	"go.uber.org/cadence/workflow"

	"github.com/banzaicloud/cadence-aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	Search(ctx workflow.Context, input *cloudsearchdomain.SearchInput) (*cloudsearchdomain.SearchOutput, error)
	SearchAsync(ctx workflow.Context, input *cloudsearchdomain.SearchInput) *CloudSearchDomainSearchFuture

	Suggest(ctx workflow.Context, input *cloudsearchdomain.SuggestInput) (*cloudsearchdomain.SuggestOutput, error)
	SuggestAsync(ctx workflow.Context, input *cloudsearchdomain.SuggestInput) *CloudSearchDomainSuggestFuture

	UploadDocuments(ctx workflow.Context, input *cloudsearchdomain.UploadDocumentsInput) (*cloudsearchdomain.UploadDocumentsOutput, error)
	UploadDocumentsAsync(ctx workflow.Context, input *cloudsearchdomain.UploadDocumentsInput) *CloudSearchDomainUploadDocumentsFuture
}

func NewClient() Client {
	return &stub{}
}
