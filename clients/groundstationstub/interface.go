// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.
// Copyright (c) 2020 Banzai Cloud All rights reserved.

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
	CancelContactAsync(ctx workflow.Context, input *groundstation.CancelContactInput) *GroundStationCancelContactFuture

	CreateConfig(ctx workflow.Context, input *groundstation.CreateConfigInput) (*groundstation.CreateConfigOutput, error)
	CreateConfigAsync(ctx workflow.Context, input *groundstation.CreateConfigInput) *GroundStationCreateConfigFuture

	CreateDataflowEndpointGroup(ctx workflow.Context, input *groundstation.CreateDataflowEndpointGroupInput) (*groundstation.CreateDataflowEndpointGroupOutput, error)
	CreateDataflowEndpointGroupAsync(ctx workflow.Context, input *groundstation.CreateDataflowEndpointGroupInput) *GroundStationCreateDataflowEndpointGroupFuture

	CreateMissionProfile(ctx workflow.Context, input *groundstation.CreateMissionProfileInput) (*groundstation.CreateMissionProfileOutput, error)
	CreateMissionProfileAsync(ctx workflow.Context, input *groundstation.CreateMissionProfileInput) *GroundStationCreateMissionProfileFuture

	DeleteConfig(ctx workflow.Context, input *groundstation.DeleteConfigInput) (*groundstation.DeleteConfigOutput, error)
	DeleteConfigAsync(ctx workflow.Context, input *groundstation.DeleteConfigInput) *GroundStationDeleteConfigFuture

	DeleteDataflowEndpointGroup(ctx workflow.Context, input *groundstation.DeleteDataflowEndpointGroupInput) (*groundstation.DeleteDataflowEndpointGroupOutput, error)
	DeleteDataflowEndpointGroupAsync(ctx workflow.Context, input *groundstation.DeleteDataflowEndpointGroupInput) *GroundStationDeleteDataflowEndpointGroupFuture

	DeleteMissionProfile(ctx workflow.Context, input *groundstation.DeleteMissionProfileInput) (*groundstation.DeleteMissionProfileOutput, error)
	DeleteMissionProfileAsync(ctx workflow.Context, input *groundstation.DeleteMissionProfileInput) *GroundStationDeleteMissionProfileFuture

	DescribeContact(ctx workflow.Context, input *groundstation.DescribeContactInput) (*groundstation.DescribeContactOutput, error)
	DescribeContactAsync(ctx workflow.Context, input *groundstation.DescribeContactInput) *GroundStationDescribeContactFuture

	GetConfig(ctx workflow.Context, input *groundstation.GetConfigInput) (*groundstation.GetConfigOutput, error)
	GetConfigAsync(ctx workflow.Context, input *groundstation.GetConfigInput) *GroundStationGetConfigFuture

	GetDataflowEndpointGroup(ctx workflow.Context, input *groundstation.GetDataflowEndpointGroupInput) (*groundstation.GetDataflowEndpointGroupOutput, error)
	GetDataflowEndpointGroupAsync(ctx workflow.Context, input *groundstation.GetDataflowEndpointGroupInput) *GroundStationGetDataflowEndpointGroupFuture

	GetMinuteUsage(ctx workflow.Context, input *groundstation.GetMinuteUsageInput) (*groundstation.GetMinuteUsageOutput, error)
	GetMinuteUsageAsync(ctx workflow.Context, input *groundstation.GetMinuteUsageInput) *GroundStationGetMinuteUsageFuture

	GetMissionProfile(ctx workflow.Context, input *groundstation.GetMissionProfileInput) (*groundstation.GetMissionProfileOutput, error)
	GetMissionProfileAsync(ctx workflow.Context, input *groundstation.GetMissionProfileInput) *GroundStationGetMissionProfileFuture

	GetSatellite(ctx workflow.Context, input *groundstation.GetSatelliteInput) (*groundstation.GetSatelliteOutput, error)
	GetSatelliteAsync(ctx workflow.Context, input *groundstation.GetSatelliteInput) *GroundStationGetSatelliteFuture

	ListConfigs(ctx workflow.Context, input *groundstation.ListConfigsInput) (*groundstation.ListConfigsOutput, error)
	ListConfigsAsync(ctx workflow.Context, input *groundstation.ListConfigsInput) *GroundStationListConfigsFuture

	ListContacts(ctx workflow.Context, input *groundstation.ListContactsInput) (*groundstation.ListContactsOutput, error)
	ListContactsAsync(ctx workflow.Context, input *groundstation.ListContactsInput) *GroundStationListContactsFuture

	ListDataflowEndpointGroups(ctx workflow.Context, input *groundstation.ListDataflowEndpointGroupsInput) (*groundstation.ListDataflowEndpointGroupsOutput, error)
	ListDataflowEndpointGroupsAsync(ctx workflow.Context, input *groundstation.ListDataflowEndpointGroupsInput) *GroundStationListDataflowEndpointGroupsFuture

	ListGroundStations(ctx workflow.Context, input *groundstation.ListGroundStationsInput) (*groundstation.ListGroundStationsOutput, error)
	ListGroundStationsAsync(ctx workflow.Context, input *groundstation.ListGroundStationsInput) *GroundStationListGroundStationsFuture

	ListMissionProfiles(ctx workflow.Context, input *groundstation.ListMissionProfilesInput) (*groundstation.ListMissionProfilesOutput, error)
	ListMissionProfilesAsync(ctx workflow.Context, input *groundstation.ListMissionProfilesInput) *GroundStationListMissionProfilesFuture

	ListSatellites(ctx workflow.Context, input *groundstation.ListSatellitesInput) (*groundstation.ListSatellitesOutput, error)
	ListSatellitesAsync(ctx workflow.Context, input *groundstation.ListSatellitesInput) *GroundStationListSatellitesFuture

	ListTagsForResource(ctx workflow.Context, input *groundstation.ListTagsForResourceInput) (*groundstation.ListTagsForResourceOutput, error)
	ListTagsForResourceAsync(ctx workflow.Context, input *groundstation.ListTagsForResourceInput) *GroundStationListTagsForResourceFuture

	ReserveContact(ctx workflow.Context, input *groundstation.ReserveContactInput) (*groundstation.ReserveContactOutput, error)
	ReserveContactAsync(ctx workflow.Context, input *groundstation.ReserveContactInput) *GroundStationReserveContactFuture

	TagResource(ctx workflow.Context, input *groundstation.TagResourceInput) (*groundstation.TagResourceOutput, error)
	TagResourceAsync(ctx workflow.Context, input *groundstation.TagResourceInput) *GroundStationTagResourceFuture

	UntagResource(ctx workflow.Context, input *groundstation.UntagResourceInput) (*groundstation.UntagResourceOutput, error)
	UntagResourceAsync(ctx workflow.Context, input *groundstation.UntagResourceInput) *GroundStationUntagResourceFuture

	UpdateConfig(ctx workflow.Context, input *groundstation.UpdateConfigInput) (*groundstation.UpdateConfigOutput, error)
	UpdateConfigAsync(ctx workflow.Context, input *groundstation.UpdateConfigInput) *GroundStationUpdateConfigFuture

	UpdateMissionProfile(ctx workflow.Context, input *groundstation.UpdateMissionProfileInput) (*groundstation.UpdateMissionProfileOutput, error)
	UpdateMissionProfileAsync(ctx workflow.Context, input *groundstation.UpdateMissionProfileInput) *GroundStationUpdateMissionProfileFuture
}

func NewClient() Client {
	return &stub{}
}