// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7

package groundstationstub

import (
	"github.com/aws/aws-sdk-go/service/groundstation"
	"go.uber.org/cadence/workflow"

	"github.com/banzaicloud/cadence-aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	CancelContact(ctx workflow.Context, input *groundstation.CancelContactInput) (*groundstation.CancelContactOutput, error)
	CancelContactAsync(ctx workflow.Context, input *groundstation.CancelContactInput) *CancelContactFuture

	CreateConfig(ctx workflow.Context, input *groundstation.CreateConfigInput) (*groundstation.CreateConfigOutput, error)
	CreateConfigAsync(ctx workflow.Context, input *groundstation.CreateConfigInput) *CreateConfigFuture

	CreateDataflowEndpointGroup(ctx workflow.Context, input *groundstation.CreateDataflowEndpointGroupInput) (*groundstation.CreateDataflowEndpointGroupOutput, error)
	CreateDataflowEndpointGroupAsync(ctx workflow.Context, input *groundstation.CreateDataflowEndpointGroupInput) *CreateDataflowEndpointGroupFuture

	CreateMissionProfile(ctx workflow.Context, input *groundstation.CreateMissionProfileInput) (*groundstation.CreateMissionProfileOutput, error)
	CreateMissionProfileAsync(ctx workflow.Context, input *groundstation.CreateMissionProfileInput) *CreateMissionProfileFuture

	DeleteConfig(ctx workflow.Context, input *groundstation.DeleteConfigInput) (*groundstation.DeleteConfigOutput, error)
	DeleteConfigAsync(ctx workflow.Context, input *groundstation.DeleteConfigInput) *DeleteConfigFuture

	DeleteDataflowEndpointGroup(ctx workflow.Context, input *groundstation.DeleteDataflowEndpointGroupInput) (*groundstation.DeleteDataflowEndpointGroupOutput, error)
	DeleteDataflowEndpointGroupAsync(ctx workflow.Context, input *groundstation.DeleteDataflowEndpointGroupInput) *DeleteDataflowEndpointGroupFuture

	DeleteMissionProfile(ctx workflow.Context, input *groundstation.DeleteMissionProfileInput) (*groundstation.DeleteMissionProfileOutput, error)
	DeleteMissionProfileAsync(ctx workflow.Context, input *groundstation.DeleteMissionProfileInput) *DeleteMissionProfileFuture

	DescribeContact(ctx workflow.Context, input *groundstation.DescribeContactInput) (*groundstation.DescribeContactOutput, error)
	DescribeContactAsync(ctx workflow.Context, input *groundstation.DescribeContactInput) *DescribeContactFuture

	GetConfig(ctx workflow.Context, input *groundstation.GetConfigInput) (*groundstation.GetConfigOutput, error)
	GetConfigAsync(ctx workflow.Context, input *groundstation.GetConfigInput) *GetConfigFuture

	GetDataflowEndpointGroup(ctx workflow.Context, input *groundstation.GetDataflowEndpointGroupInput) (*groundstation.GetDataflowEndpointGroupOutput, error)
	GetDataflowEndpointGroupAsync(ctx workflow.Context, input *groundstation.GetDataflowEndpointGroupInput) *GetDataflowEndpointGroupFuture

	GetMinuteUsage(ctx workflow.Context, input *groundstation.GetMinuteUsageInput) (*groundstation.GetMinuteUsageOutput, error)
	GetMinuteUsageAsync(ctx workflow.Context, input *groundstation.GetMinuteUsageInput) *GetMinuteUsageFuture

	GetMissionProfile(ctx workflow.Context, input *groundstation.GetMissionProfileInput) (*groundstation.GetMissionProfileOutput, error)
	GetMissionProfileAsync(ctx workflow.Context, input *groundstation.GetMissionProfileInput) *GetMissionProfileFuture

	GetSatellite(ctx workflow.Context, input *groundstation.GetSatelliteInput) (*groundstation.GetSatelliteOutput, error)
	GetSatelliteAsync(ctx workflow.Context, input *groundstation.GetSatelliteInput) *GetSatelliteFuture

	ListConfigs(ctx workflow.Context, input *groundstation.ListConfigsInput) (*groundstation.ListConfigsOutput, error)
	ListConfigsAsync(ctx workflow.Context, input *groundstation.ListConfigsInput) *ListConfigsFuture

	ListContacts(ctx workflow.Context, input *groundstation.ListContactsInput) (*groundstation.ListContactsOutput, error)
	ListContactsAsync(ctx workflow.Context, input *groundstation.ListContactsInput) *ListContactsFuture

	ListDataflowEndpointGroups(ctx workflow.Context, input *groundstation.ListDataflowEndpointGroupsInput) (*groundstation.ListDataflowEndpointGroupsOutput, error)
	ListDataflowEndpointGroupsAsync(ctx workflow.Context, input *groundstation.ListDataflowEndpointGroupsInput) *ListDataflowEndpointGroupsFuture

	ListGroundStations(ctx workflow.Context, input *groundstation.ListGroundStationsInput) (*groundstation.ListGroundStationsOutput, error)
	ListGroundStationsAsync(ctx workflow.Context, input *groundstation.ListGroundStationsInput) *ListGroundStationsFuture

	ListMissionProfiles(ctx workflow.Context, input *groundstation.ListMissionProfilesInput) (*groundstation.ListMissionProfilesOutput, error)
	ListMissionProfilesAsync(ctx workflow.Context, input *groundstation.ListMissionProfilesInput) *ListMissionProfilesFuture

	ListSatellites(ctx workflow.Context, input *groundstation.ListSatellitesInput) (*groundstation.ListSatellitesOutput, error)
	ListSatellitesAsync(ctx workflow.Context, input *groundstation.ListSatellitesInput) *ListSatellitesFuture

	ListTagsForResource(ctx workflow.Context, input *groundstation.ListTagsForResourceInput) (*groundstation.ListTagsForResourceOutput, error)
	ListTagsForResourceAsync(ctx workflow.Context, input *groundstation.ListTagsForResourceInput) *ListTagsForResourceFuture

	ReserveContact(ctx workflow.Context, input *groundstation.ReserveContactInput) (*groundstation.ReserveContactOutput, error)
	ReserveContactAsync(ctx workflow.Context, input *groundstation.ReserveContactInput) *ReserveContactFuture

	TagResource(ctx workflow.Context, input *groundstation.TagResourceInput) (*groundstation.TagResourceOutput, error)
	TagResourceAsync(ctx workflow.Context, input *groundstation.TagResourceInput) *TagResourceFuture

	UntagResource(ctx workflow.Context, input *groundstation.UntagResourceInput) (*groundstation.UntagResourceOutput, error)
	UntagResourceAsync(ctx workflow.Context, input *groundstation.UntagResourceInput) *UntagResourceFuture

	UpdateConfig(ctx workflow.Context, input *groundstation.UpdateConfigInput) (*groundstation.UpdateConfigOutput, error)
	UpdateConfigAsync(ctx workflow.Context, input *groundstation.UpdateConfigInput) *UpdateConfigFuture

	UpdateMissionProfile(ctx workflow.Context, input *groundstation.UpdateMissionProfileInput) (*groundstation.UpdateMissionProfileOutput, error)
	UpdateMissionProfileAsync(ctx workflow.Context, input *groundstation.UpdateMissionProfileInput) *UpdateMissionProfileFuture
}

func NewClient() Client {
	return &stub{}
}
