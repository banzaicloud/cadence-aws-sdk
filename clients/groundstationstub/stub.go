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

type stub struct{}

type GroundStationCancelContactFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GroundStationCancelContactFuture) Get(ctx workflow.Context) (*groundstation.CancelContactOutput, error) {
	var output groundstation.CancelContactOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GroundStationCreateConfigFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GroundStationCreateConfigFuture) Get(ctx workflow.Context) (*groundstation.CreateConfigOutput, error) {
	var output groundstation.CreateConfigOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GroundStationCreateDataflowEndpointGroupFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GroundStationCreateDataflowEndpointGroupFuture) Get(ctx workflow.Context) (*groundstation.CreateDataflowEndpointGroupOutput, error) {
	var output groundstation.CreateDataflowEndpointGroupOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GroundStationCreateMissionProfileFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GroundStationCreateMissionProfileFuture) Get(ctx workflow.Context) (*groundstation.CreateMissionProfileOutput, error) {
	var output groundstation.CreateMissionProfileOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GroundStationDeleteConfigFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GroundStationDeleteConfigFuture) Get(ctx workflow.Context) (*groundstation.DeleteConfigOutput, error) {
	var output groundstation.DeleteConfigOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GroundStationDeleteDataflowEndpointGroupFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GroundStationDeleteDataflowEndpointGroupFuture) Get(ctx workflow.Context) (*groundstation.DeleteDataflowEndpointGroupOutput, error) {
	var output groundstation.DeleteDataflowEndpointGroupOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GroundStationDeleteMissionProfileFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GroundStationDeleteMissionProfileFuture) Get(ctx workflow.Context) (*groundstation.DeleteMissionProfileOutput, error) {
	var output groundstation.DeleteMissionProfileOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GroundStationDescribeContactFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GroundStationDescribeContactFuture) Get(ctx workflow.Context) (*groundstation.DescribeContactOutput, error) {
	var output groundstation.DescribeContactOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GroundStationGetConfigFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GroundStationGetConfigFuture) Get(ctx workflow.Context) (*groundstation.GetConfigOutput, error) {
	var output groundstation.GetConfigOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GroundStationGetDataflowEndpointGroupFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GroundStationGetDataflowEndpointGroupFuture) Get(ctx workflow.Context) (*groundstation.GetDataflowEndpointGroupOutput, error) {
	var output groundstation.GetDataflowEndpointGroupOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GroundStationGetMinuteUsageFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GroundStationGetMinuteUsageFuture) Get(ctx workflow.Context) (*groundstation.GetMinuteUsageOutput, error) {
	var output groundstation.GetMinuteUsageOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GroundStationGetMissionProfileFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GroundStationGetMissionProfileFuture) Get(ctx workflow.Context) (*groundstation.GetMissionProfileOutput, error) {
	var output groundstation.GetMissionProfileOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GroundStationGetSatelliteFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GroundStationGetSatelliteFuture) Get(ctx workflow.Context) (*groundstation.GetSatelliteOutput, error) {
	var output groundstation.GetSatelliteOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GroundStationListConfigsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GroundStationListConfigsFuture) Get(ctx workflow.Context) (*groundstation.ListConfigsOutput, error) {
	var output groundstation.ListConfigsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GroundStationListContactsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GroundStationListContactsFuture) Get(ctx workflow.Context) (*groundstation.ListContactsOutput, error) {
	var output groundstation.ListContactsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GroundStationListDataflowEndpointGroupsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GroundStationListDataflowEndpointGroupsFuture) Get(ctx workflow.Context) (*groundstation.ListDataflowEndpointGroupsOutput, error) {
	var output groundstation.ListDataflowEndpointGroupsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GroundStationListGroundStationsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GroundStationListGroundStationsFuture) Get(ctx workflow.Context) (*groundstation.ListGroundStationsOutput, error) {
	var output groundstation.ListGroundStationsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GroundStationListMissionProfilesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GroundStationListMissionProfilesFuture) Get(ctx workflow.Context) (*groundstation.ListMissionProfilesOutput, error) {
	var output groundstation.ListMissionProfilesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GroundStationListSatellitesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GroundStationListSatellitesFuture) Get(ctx workflow.Context) (*groundstation.ListSatellitesOutput, error) {
	var output groundstation.ListSatellitesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GroundStationListTagsForResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GroundStationListTagsForResourceFuture) Get(ctx workflow.Context) (*groundstation.ListTagsForResourceOutput, error) {
	var output groundstation.ListTagsForResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GroundStationReserveContactFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GroundStationReserveContactFuture) Get(ctx workflow.Context) (*groundstation.ReserveContactOutput, error) {
	var output groundstation.ReserveContactOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GroundStationTagResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GroundStationTagResourceFuture) Get(ctx workflow.Context) (*groundstation.TagResourceOutput, error) {
	var output groundstation.TagResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GroundStationUntagResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GroundStationUntagResourceFuture) Get(ctx workflow.Context) (*groundstation.UntagResourceOutput, error) {
	var output groundstation.UntagResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GroundStationUpdateConfigFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GroundStationUpdateConfigFuture) Get(ctx workflow.Context) (*groundstation.UpdateConfigOutput, error) {
	var output groundstation.UpdateConfigOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GroundStationUpdateMissionProfileFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GroundStationUpdateMissionProfileFuture) Get(ctx workflow.Context) (*groundstation.UpdateMissionProfileOutput, error) {
	var output groundstation.UpdateMissionProfileOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *stub) CancelContact(ctx workflow.Context, input *groundstation.CancelContactInput) (*groundstation.CancelContactOutput, error) {
	var output groundstation.CancelContactOutput
	err := workflow.ExecuteActivity(ctx, "aws-groundstation-CancelContact", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CancelContactAsync(ctx workflow.Context, input *groundstation.CancelContactInput) *GroundStationCancelContactFuture {
	future := workflow.ExecuteActivity(ctx, "aws-groundstation-CancelContact", input)
	return &GroundStationCancelContactFuture{Future: future}
}

func (a *stub) CreateConfig(ctx workflow.Context, input *groundstation.CreateConfigInput) (*groundstation.CreateConfigOutput, error) {
	var output groundstation.CreateConfigOutput
	err := workflow.ExecuteActivity(ctx, "aws-groundstation-CreateConfig", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateConfigAsync(ctx workflow.Context, input *groundstation.CreateConfigInput) *GroundStationCreateConfigFuture {
	future := workflow.ExecuteActivity(ctx, "aws-groundstation-CreateConfig", input)
	return &GroundStationCreateConfigFuture{Future: future}
}

func (a *stub) CreateDataflowEndpointGroup(ctx workflow.Context, input *groundstation.CreateDataflowEndpointGroupInput) (*groundstation.CreateDataflowEndpointGroupOutput, error) {
	var output groundstation.CreateDataflowEndpointGroupOutput
	err := workflow.ExecuteActivity(ctx, "aws-groundstation-CreateDataflowEndpointGroup", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateDataflowEndpointGroupAsync(ctx workflow.Context, input *groundstation.CreateDataflowEndpointGroupInput) *GroundStationCreateDataflowEndpointGroupFuture {
	future := workflow.ExecuteActivity(ctx, "aws-groundstation-CreateDataflowEndpointGroup", input)
	return &GroundStationCreateDataflowEndpointGroupFuture{Future: future}
}

func (a *stub) CreateMissionProfile(ctx workflow.Context, input *groundstation.CreateMissionProfileInput) (*groundstation.CreateMissionProfileOutput, error) {
	var output groundstation.CreateMissionProfileOutput
	err := workflow.ExecuteActivity(ctx, "aws-groundstation-CreateMissionProfile", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateMissionProfileAsync(ctx workflow.Context, input *groundstation.CreateMissionProfileInput) *GroundStationCreateMissionProfileFuture {
	future := workflow.ExecuteActivity(ctx, "aws-groundstation-CreateMissionProfile", input)
	return &GroundStationCreateMissionProfileFuture{Future: future}
}

func (a *stub) DeleteConfig(ctx workflow.Context, input *groundstation.DeleteConfigInput) (*groundstation.DeleteConfigOutput, error) {
	var output groundstation.DeleteConfigOutput
	err := workflow.ExecuteActivity(ctx, "aws-groundstation-DeleteConfig", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteConfigAsync(ctx workflow.Context, input *groundstation.DeleteConfigInput) *GroundStationDeleteConfigFuture {
	future := workflow.ExecuteActivity(ctx, "aws-groundstation-DeleteConfig", input)
	return &GroundStationDeleteConfigFuture{Future: future}
}

func (a *stub) DeleteDataflowEndpointGroup(ctx workflow.Context, input *groundstation.DeleteDataflowEndpointGroupInput) (*groundstation.DeleteDataflowEndpointGroupOutput, error) {
	var output groundstation.DeleteDataflowEndpointGroupOutput
	err := workflow.ExecuteActivity(ctx, "aws-groundstation-DeleteDataflowEndpointGroup", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteDataflowEndpointGroupAsync(ctx workflow.Context, input *groundstation.DeleteDataflowEndpointGroupInput) *GroundStationDeleteDataflowEndpointGroupFuture {
	future := workflow.ExecuteActivity(ctx, "aws-groundstation-DeleteDataflowEndpointGroup", input)
	return &GroundStationDeleteDataflowEndpointGroupFuture{Future: future}
}

func (a *stub) DeleteMissionProfile(ctx workflow.Context, input *groundstation.DeleteMissionProfileInput) (*groundstation.DeleteMissionProfileOutput, error) {
	var output groundstation.DeleteMissionProfileOutput
	err := workflow.ExecuteActivity(ctx, "aws-groundstation-DeleteMissionProfile", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteMissionProfileAsync(ctx workflow.Context, input *groundstation.DeleteMissionProfileInput) *GroundStationDeleteMissionProfileFuture {
	future := workflow.ExecuteActivity(ctx, "aws-groundstation-DeleteMissionProfile", input)
	return &GroundStationDeleteMissionProfileFuture{Future: future}
}

func (a *stub) DescribeContact(ctx workflow.Context, input *groundstation.DescribeContactInput) (*groundstation.DescribeContactOutput, error) {
	var output groundstation.DescribeContactOutput
	err := workflow.ExecuteActivity(ctx, "aws-groundstation-DescribeContact", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeContactAsync(ctx workflow.Context, input *groundstation.DescribeContactInput) *GroundStationDescribeContactFuture {
	future := workflow.ExecuteActivity(ctx, "aws-groundstation-DescribeContact", input)
	return &GroundStationDescribeContactFuture{Future: future}
}

func (a *stub) GetConfig(ctx workflow.Context, input *groundstation.GetConfigInput) (*groundstation.GetConfigOutput, error) {
	var output groundstation.GetConfigOutput
	err := workflow.ExecuteActivity(ctx, "aws-groundstation-GetConfig", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetConfigAsync(ctx workflow.Context, input *groundstation.GetConfigInput) *GroundStationGetConfigFuture {
	future := workflow.ExecuteActivity(ctx, "aws-groundstation-GetConfig", input)
	return &GroundStationGetConfigFuture{Future: future}
}

func (a *stub) GetDataflowEndpointGroup(ctx workflow.Context, input *groundstation.GetDataflowEndpointGroupInput) (*groundstation.GetDataflowEndpointGroupOutput, error) {
	var output groundstation.GetDataflowEndpointGroupOutput
	err := workflow.ExecuteActivity(ctx, "aws-groundstation-GetDataflowEndpointGroup", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetDataflowEndpointGroupAsync(ctx workflow.Context, input *groundstation.GetDataflowEndpointGroupInput) *GroundStationGetDataflowEndpointGroupFuture {
	future := workflow.ExecuteActivity(ctx, "aws-groundstation-GetDataflowEndpointGroup", input)
	return &GroundStationGetDataflowEndpointGroupFuture{Future: future}
}

func (a *stub) GetMinuteUsage(ctx workflow.Context, input *groundstation.GetMinuteUsageInput) (*groundstation.GetMinuteUsageOutput, error) {
	var output groundstation.GetMinuteUsageOutput
	err := workflow.ExecuteActivity(ctx, "aws-groundstation-GetMinuteUsage", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetMinuteUsageAsync(ctx workflow.Context, input *groundstation.GetMinuteUsageInput) *GroundStationGetMinuteUsageFuture {
	future := workflow.ExecuteActivity(ctx, "aws-groundstation-GetMinuteUsage", input)
	return &GroundStationGetMinuteUsageFuture{Future: future}
}

func (a *stub) GetMissionProfile(ctx workflow.Context, input *groundstation.GetMissionProfileInput) (*groundstation.GetMissionProfileOutput, error) {
	var output groundstation.GetMissionProfileOutput
	err := workflow.ExecuteActivity(ctx, "aws-groundstation-GetMissionProfile", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetMissionProfileAsync(ctx workflow.Context, input *groundstation.GetMissionProfileInput) *GroundStationGetMissionProfileFuture {
	future := workflow.ExecuteActivity(ctx, "aws-groundstation-GetMissionProfile", input)
	return &GroundStationGetMissionProfileFuture{Future: future}
}

func (a *stub) GetSatellite(ctx workflow.Context, input *groundstation.GetSatelliteInput) (*groundstation.GetSatelliteOutput, error) {
	var output groundstation.GetSatelliteOutput
	err := workflow.ExecuteActivity(ctx, "aws-groundstation-GetSatellite", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetSatelliteAsync(ctx workflow.Context, input *groundstation.GetSatelliteInput) *GroundStationGetSatelliteFuture {
	future := workflow.ExecuteActivity(ctx, "aws-groundstation-GetSatellite", input)
	return &GroundStationGetSatelliteFuture{Future: future}
}

func (a *stub) ListConfigs(ctx workflow.Context, input *groundstation.ListConfigsInput) (*groundstation.ListConfigsOutput, error) {
	var output groundstation.ListConfigsOutput
	err := workflow.ExecuteActivity(ctx, "aws-groundstation-ListConfigs", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListConfigsAsync(ctx workflow.Context, input *groundstation.ListConfigsInput) *GroundStationListConfigsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-groundstation-ListConfigs", input)
	return &GroundStationListConfigsFuture{Future: future}
}

func (a *stub) ListContacts(ctx workflow.Context, input *groundstation.ListContactsInput) (*groundstation.ListContactsOutput, error) {
	var output groundstation.ListContactsOutput
	err := workflow.ExecuteActivity(ctx, "aws-groundstation-ListContacts", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListContactsAsync(ctx workflow.Context, input *groundstation.ListContactsInput) *GroundStationListContactsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-groundstation-ListContacts", input)
	return &GroundStationListContactsFuture{Future: future}
}

func (a *stub) ListDataflowEndpointGroups(ctx workflow.Context, input *groundstation.ListDataflowEndpointGroupsInput) (*groundstation.ListDataflowEndpointGroupsOutput, error) {
	var output groundstation.ListDataflowEndpointGroupsOutput
	err := workflow.ExecuteActivity(ctx, "aws-groundstation-ListDataflowEndpointGroups", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListDataflowEndpointGroupsAsync(ctx workflow.Context, input *groundstation.ListDataflowEndpointGroupsInput) *GroundStationListDataflowEndpointGroupsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-groundstation-ListDataflowEndpointGroups", input)
	return &GroundStationListDataflowEndpointGroupsFuture{Future: future}
}

func (a *stub) ListGroundStations(ctx workflow.Context, input *groundstation.ListGroundStationsInput) (*groundstation.ListGroundStationsOutput, error) {
	var output groundstation.ListGroundStationsOutput
	err := workflow.ExecuteActivity(ctx, "aws-groundstation-ListGroundStations", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListGroundStationsAsync(ctx workflow.Context, input *groundstation.ListGroundStationsInput) *GroundStationListGroundStationsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-groundstation-ListGroundStations", input)
	return &GroundStationListGroundStationsFuture{Future: future}
}

func (a *stub) ListMissionProfiles(ctx workflow.Context, input *groundstation.ListMissionProfilesInput) (*groundstation.ListMissionProfilesOutput, error) {
	var output groundstation.ListMissionProfilesOutput
	err := workflow.ExecuteActivity(ctx, "aws-groundstation-ListMissionProfiles", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListMissionProfilesAsync(ctx workflow.Context, input *groundstation.ListMissionProfilesInput) *GroundStationListMissionProfilesFuture {
	future := workflow.ExecuteActivity(ctx, "aws-groundstation-ListMissionProfiles", input)
	return &GroundStationListMissionProfilesFuture{Future: future}
}

func (a *stub) ListSatellites(ctx workflow.Context, input *groundstation.ListSatellitesInput) (*groundstation.ListSatellitesOutput, error) {
	var output groundstation.ListSatellitesOutput
	err := workflow.ExecuteActivity(ctx, "aws-groundstation-ListSatellites", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListSatellitesAsync(ctx workflow.Context, input *groundstation.ListSatellitesInput) *GroundStationListSatellitesFuture {
	future := workflow.ExecuteActivity(ctx, "aws-groundstation-ListSatellites", input)
	return &GroundStationListSatellitesFuture{Future: future}
}

func (a *stub) ListTagsForResource(ctx workflow.Context, input *groundstation.ListTagsForResourceInput) (*groundstation.ListTagsForResourceOutput, error) {
	var output groundstation.ListTagsForResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws-groundstation-ListTagsForResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListTagsForResourceAsync(ctx workflow.Context, input *groundstation.ListTagsForResourceInput) *GroundStationListTagsForResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws-groundstation-ListTagsForResource", input)
	return &GroundStationListTagsForResourceFuture{Future: future}
}

func (a *stub) ReserveContact(ctx workflow.Context, input *groundstation.ReserveContactInput) (*groundstation.ReserveContactOutput, error) {
	var output groundstation.ReserveContactOutput
	err := workflow.ExecuteActivity(ctx, "aws-groundstation-ReserveContact", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ReserveContactAsync(ctx workflow.Context, input *groundstation.ReserveContactInput) *GroundStationReserveContactFuture {
	future := workflow.ExecuteActivity(ctx, "aws-groundstation-ReserveContact", input)
	return &GroundStationReserveContactFuture{Future: future}
}

func (a *stub) TagResource(ctx workflow.Context, input *groundstation.TagResourceInput) (*groundstation.TagResourceOutput, error) {
	var output groundstation.TagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws-groundstation-TagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) TagResourceAsync(ctx workflow.Context, input *groundstation.TagResourceInput) *GroundStationTagResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws-groundstation-TagResource", input)
	return &GroundStationTagResourceFuture{Future: future}
}

func (a *stub) UntagResource(ctx workflow.Context, input *groundstation.UntagResourceInput) (*groundstation.UntagResourceOutput, error) {
	var output groundstation.UntagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws-groundstation-UntagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UntagResourceAsync(ctx workflow.Context, input *groundstation.UntagResourceInput) *GroundStationUntagResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws-groundstation-UntagResource", input)
	return &GroundStationUntagResourceFuture{Future: future}
}

func (a *stub) UpdateConfig(ctx workflow.Context, input *groundstation.UpdateConfigInput) (*groundstation.UpdateConfigOutput, error) {
	var output groundstation.UpdateConfigOutput
	err := workflow.ExecuteActivity(ctx, "aws-groundstation-UpdateConfig", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateConfigAsync(ctx workflow.Context, input *groundstation.UpdateConfigInput) *GroundStationUpdateConfigFuture {
	future := workflow.ExecuteActivity(ctx, "aws-groundstation-UpdateConfig", input)
	return &GroundStationUpdateConfigFuture{Future: future}
}

func (a *stub) UpdateMissionProfile(ctx workflow.Context, input *groundstation.UpdateMissionProfileInput) (*groundstation.UpdateMissionProfileOutput, error) {
	var output groundstation.UpdateMissionProfileOutput
	err := workflow.ExecuteActivity(ctx, "aws-groundstation-UpdateMissionProfile", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateMissionProfileAsync(ctx workflow.Context, input *groundstation.UpdateMissionProfileInput) *GroundStationUpdateMissionProfileFuture {
	future := workflow.ExecuteActivity(ctx, "aws-groundstation-UpdateMissionProfile", input)
	return &GroundStationUpdateMissionProfileFuture{Future: future}
}
