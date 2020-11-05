// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7

package marketplacecatalogstub

import (
	"github.com/aws/aws-sdk-go/service/marketplacecatalog"
	"go.uber.org/cadence/workflow"

	"github.com/banzaicloud/cadence-aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	CancelChangeSet(ctx workflow.Context, input *marketplacecatalog.CancelChangeSetInput) (*marketplacecatalog.CancelChangeSetOutput, error)
	CancelChangeSetAsync(ctx workflow.Context, input *marketplacecatalog.CancelChangeSetInput) *MarketplaceCatalogCancelChangeSetFuture

	DescribeChangeSet(ctx workflow.Context, input *marketplacecatalog.DescribeChangeSetInput) (*marketplacecatalog.DescribeChangeSetOutput, error)
	DescribeChangeSetAsync(ctx workflow.Context, input *marketplacecatalog.DescribeChangeSetInput) *MarketplaceCatalogDescribeChangeSetFuture

	DescribeEntity(ctx workflow.Context, input *marketplacecatalog.DescribeEntityInput) (*marketplacecatalog.DescribeEntityOutput, error)
	DescribeEntityAsync(ctx workflow.Context, input *marketplacecatalog.DescribeEntityInput) *MarketplaceCatalogDescribeEntityFuture

	ListChangeSets(ctx workflow.Context, input *marketplacecatalog.ListChangeSetsInput) (*marketplacecatalog.ListChangeSetsOutput, error)
	ListChangeSetsAsync(ctx workflow.Context, input *marketplacecatalog.ListChangeSetsInput) *MarketplaceCatalogListChangeSetsFuture

	ListEntities(ctx workflow.Context, input *marketplacecatalog.ListEntitiesInput) (*marketplacecatalog.ListEntitiesOutput, error)
	ListEntitiesAsync(ctx workflow.Context, input *marketplacecatalog.ListEntitiesInput) *MarketplaceCatalogListEntitiesFuture

	StartChangeSet(ctx workflow.Context, input *marketplacecatalog.StartChangeSetInput) (*marketplacecatalog.StartChangeSetOutput, error)
	StartChangeSetAsync(ctx workflow.Context, input *marketplacecatalog.StartChangeSetInput) *MarketplaceCatalogStartChangeSetFuture
}

func NewClient() Client {
	return &stub{}
}
