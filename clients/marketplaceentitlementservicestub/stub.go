// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7

package marketplaceentitlementservicestub

import (
	"github.com/aws/aws-sdk-go/service/marketplaceentitlementservice"
	"go.uber.org/cadence/workflow"

	"github.com/banzaicloud/cadence-aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type stub struct{}

type MarketplaceEntitlementServiceGetEntitlementsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MarketplaceEntitlementServiceGetEntitlementsFuture) Get(ctx workflow.Context) (*marketplaceentitlementservice.GetEntitlementsOutput, error) {
	var output marketplaceentitlementservice.GetEntitlementsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *stub) GetEntitlements(ctx workflow.Context, input *marketplaceentitlementservice.GetEntitlementsInput) (*marketplaceentitlementservice.GetEntitlementsOutput, error) {
	var output marketplaceentitlementservice.GetEntitlementsOutput
	err := workflow.ExecuteActivity(ctx, "aws-marketplaceentitlementservice-GetEntitlements", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetEntitlementsAsync(ctx workflow.Context, input *marketplaceentitlementservice.GetEntitlementsInput) *MarketplaceEntitlementServiceGetEntitlementsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-marketplaceentitlementservice-GetEntitlements", input)
	return &MarketplaceEntitlementServiceGetEntitlementsFuture{Future: future}
}
