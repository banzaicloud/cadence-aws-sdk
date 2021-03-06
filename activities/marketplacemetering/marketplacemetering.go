// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7

package marketplacemetering

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/marketplacemetering"
	"github.com/aws/aws-sdk-go/service/marketplacemetering/marketplacemeteringiface"

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
	client marketplacemeteringiface.MarketplaceMeteringAPI

	sessionFactory SessionFactory
}

func NewActivities(sess *session.Session, config ...*aws.Config) *Activities {
	client := marketplacemetering.New(sess, config...)
	return &Activities{client: client}
}

func NewActivitiesWithSessionFactory(sessionFactory SessionFactory) *Activities {
	return &Activities{sessionFactory: sessionFactory}
}

func (a *Activities) getClient(ctx context.Context) (marketplacemeteringiface.MarketplaceMeteringAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}

	return marketplacemetering.New(sess), nil
}

func (a *Activities) BatchMeterUsage(ctx context.Context, input *marketplacemetering.BatchMeterUsageInput) (*marketplacemetering.BatchMeterUsageOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.BatchMeterUsageWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) MeterUsage(ctx context.Context, input *marketplacemetering.MeterUsageInput) (*marketplacemetering.MeterUsageOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.MeterUsageWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) RegisterUsage(ctx context.Context, input *marketplacemetering.RegisterUsageInput) (*marketplacemetering.RegisterUsageOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.RegisterUsageWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ResolveCustomer(ctx context.Context, input *marketplacemetering.ResolveCustomerInput) (*marketplacemetering.ResolveCustomerOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ResolveCustomerWithContext(ctx, input)

	return output, internal.EncodeError(err)
}
