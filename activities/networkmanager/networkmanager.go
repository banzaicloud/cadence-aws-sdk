// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7

package networkmanager

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/networkmanager"
	"github.com/aws/aws-sdk-go/service/networkmanager/networkmanageriface"

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
	client networkmanageriface.NetworkManagerAPI

	sessionFactory SessionFactory
}

func NewActivities(sess *session.Session, config ...*aws.Config) *Activities {
	client := networkmanager.New(sess, config...)
	return &Activities{client: client}
}

func NewActivitiesWithSessionFactory(sessionFactory SessionFactory) *Activities {
	return &Activities{sessionFactory: sessionFactory}
}

func (a *Activities) getClient(ctx context.Context) (networkmanageriface.NetworkManagerAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}

	return networkmanager.New(sess), nil
}

func (a *Activities) AssociateCustomerGateway(ctx context.Context, input *networkmanager.AssociateCustomerGatewayInput) (*networkmanager.AssociateCustomerGatewayOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.AssociateCustomerGatewayWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) AssociateLink(ctx context.Context, input *networkmanager.AssociateLinkInput) (*networkmanager.AssociateLinkOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.AssociateLinkWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) AssociateTransitGatewayConnectPeer(ctx context.Context, input *networkmanager.AssociateTransitGatewayConnectPeerInput) (*networkmanager.AssociateTransitGatewayConnectPeerOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.AssociateTransitGatewayConnectPeerWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateConnection(ctx context.Context, input *networkmanager.CreateConnectionInput) (*networkmanager.CreateConnectionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateConnectionWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateDevice(ctx context.Context, input *networkmanager.CreateDeviceInput) (*networkmanager.CreateDeviceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateDeviceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateGlobalNetwork(ctx context.Context, input *networkmanager.CreateGlobalNetworkInput) (*networkmanager.CreateGlobalNetworkOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateGlobalNetworkWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateLink(ctx context.Context, input *networkmanager.CreateLinkInput) (*networkmanager.CreateLinkOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateLinkWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateSite(ctx context.Context, input *networkmanager.CreateSiteInput) (*networkmanager.CreateSiteOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateSiteWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteConnection(ctx context.Context, input *networkmanager.DeleteConnectionInput) (*networkmanager.DeleteConnectionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteConnectionWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteDevice(ctx context.Context, input *networkmanager.DeleteDeviceInput) (*networkmanager.DeleteDeviceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteDeviceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteGlobalNetwork(ctx context.Context, input *networkmanager.DeleteGlobalNetworkInput) (*networkmanager.DeleteGlobalNetworkOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteGlobalNetworkWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteLink(ctx context.Context, input *networkmanager.DeleteLinkInput) (*networkmanager.DeleteLinkOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteLinkWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteSite(ctx context.Context, input *networkmanager.DeleteSiteInput) (*networkmanager.DeleteSiteOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteSiteWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeregisterTransitGateway(ctx context.Context, input *networkmanager.DeregisterTransitGatewayInput) (*networkmanager.DeregisterTransitGatewayOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeregisterTransitGatewayWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeGlobalNetworks(ctx context.Context, input *networkmanager.DescribeGlobalNetworksInput) (*networkmanager.DescribeGlobalNetworksOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeGlobalNetworksWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DisassociateCustomerGateway(ctx context.Context, input *networkmanager.DisassociateCustomerGatewayInput) (*networkmanager.DisassociateCustomerGatewayOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DisassociateCustomerGatewayWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DisassociateLink(ctx context.Context, input *networkmanager.DisassociateLinkInput) (*networkmanager.DisassociateLinkOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DisassociateLinkWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DisassociateTransitGatewayConnectPeer(ctx context.Context, input *networkmanager.DisassociateTransitGatewayConnectPeerInput) (*networkmanager.DisassociateTransitGatewayConnectPeerOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DisassociateTransitGatewayConnectPeerWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetConnections(ctx context.Context, input *networkmanager.GetConnectionsInput) (*networkmanager.GetConnectionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetConnectionsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetCustomerGatewayAssociations(ctx context.Context, input *networkmanager.GetCustomerGatewayAssociationsInput) (*networkmanager.GetCustomerGatewayAssociationsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetCustomerGatewayAssociationsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetDevices(ctx context.Context, input *networkmanager.GetDevicesInput) (*networkmanager.GetDevicesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetDevicesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetLinkAssociations(ctx context.Context, input *networkmanager.GetLinkAssociationsInput) (*networkmanager.GetLinkAssociationsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetLinkAssociationsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetLinks(ctx context.Context, input *networkmanager.GetLinksInput) (*networkmanager.GetLinksOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetLinksWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetSites(ctx context.Context, input *networkmanager.GetSitesInput) (*networkmanager.GetSitesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetSitesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetTransitGatewayConnectPeerAssociations(ctx context.Context, input *networkmanager.GetTransitGatewayConnectPeerAssociationsInput) (*networkmanager.GetTransitGatewayConnectPeerAssociationsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetTransitGatewayConnectPeerAssociationsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetTransitGatewayRegistrations(ctx context.Context, input *networkmanager.GetTransitGatewayRegistrationsInput) (*networkmanager.GetTransitGatewayRegistrationsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetTransitGatewayRegistrationsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListTagsForResource(ctx context.Context, input *networkmanager.ListTagsForResourceInput) (*networkmanager.ListTagsForResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListTagsForResourceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) RegisterTransitGateway(ctx context.Context, input *networkmanager.RegisterTransitGatewayInput) (*networkmanager.RegisterTransitGatewayOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.RegisterTransitGatewayWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) TagResource(ctx context.Context, input *networkmanager.TagResourceInput) (*networkmanager.TagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.TagResourceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UntagResource(ctx context.Context, input *networkmanager.UntagResourceInput) (*networkmanager.UntagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UntagResourceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateConnection(ctx context.Context, input *networkmanager.UpdateConnectionInput) (*networkmanager.UpdateConnectionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateConnectionWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateDevice(ctx context.Context, input *networkmanager.UpdateDeviceInput) (*networkmanager.UpdateDeviceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateDeviceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateGlobalNetwork(ctx context.Context, input *networkmanager.UpdateGlobalNetworkInput) (*networkmanager.UpdateGlobalNetworkOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateGlobalNetworkWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateLink(ctx context.Context, input *networkmanager.UpdateLinkInput) (*networkmanager.UpdateLinkOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateLinkWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateSite(ctx context.Context, input *networkmanager.UpdateSiteInput) (*networkmanager.UpdateSiteOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateSiteWithContext(ctx, input)

	return output, internal.EncodeError(err)
}
