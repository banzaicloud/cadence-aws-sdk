// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7

package cloudhsmstub

import (
	"github.com/aws/aws-sdk-go/service/cloudhsm"
	"go.uber.org/cadence/workflow"

	"github.com/banzaicloud/cadence-aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type stub struct{}

type AddTagsToResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AddTagsToResourceFuture) Get(ctx workflow.Context) (*cloudhsm.AddTagsToResourceOutput, error) {
	var output cloudhsm.AddTagsToResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CreateHapgFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CreateHapgFuture) Get(ctx workflow.Context) (*cloudhsm.CreateHapgOutput, error) {
	var output cloudhsm.CreateHapgOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CreateHsmFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CreateHsmFuture) Get(ctx workflow.Context) (*cloudhsm.CreateHsmOutput, error) {
	var output cloudhsm.CreateHsmOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CreateLunaClientFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CreateLunaClientFuture) Get(ctx workflow.Context) (*cloudhsm.CreateLunaClientOutput, error) {
	var output cloudhsm.CreateLunaClientOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DeleteHapgFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DeleteHapgFuture) Get(ctx workflow.Context) (*cloudhsm.DeleteHapgOutput, error) {
	var output cloudhsm.DeleteHapgOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DeleteHsmFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DeleteHsmFuture) Get(ctx workflow.Context) (*cloudhsm.DeleteHsmOutput, error) {
	var output cloudhsm.DeleteHsmOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DeleteLunaClientFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DeleteLunaClientFuture) Get(ctx workflow.Context) (*cloudhsm.DeleteLunaClientOutput, error) {
	var output cloudhsm.DeleteLunaClientOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DescribeHapgFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DescribeHapgFuture) Get(ctx workflow.Context) (*cloudhsm.DescribeHapgOutput, error) {
	var output cloudhsm.DescribeHapgOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DescribeHsmFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DescribeHsmFuture) Get(ctx workflow.Context) (*cloudhsm.DescribeHsmOutput, error) {
	var output cloudhsm.DescribeHsmOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DescribeLunaClientFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DescribeLunaClientFuture) Get(ctx workflow.Context) (*cloudhsm.DescribeLunaClientOutput, error) {
	var output cloudhsm.DescribeLunaClientOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GetConfigFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GetConfigFuture) Get(ctx workflow.Context) (*cloudhsm.GetConfigOutput, error) {
	var output cloudhsm.GetConfigOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListAvailableZonesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListAvailableZonesFuture) Get(ctx workflow.Context) (*cloudhsm.ListAvailableZonesOutput, error) {
	var output cloudhsm.ListAvailableZonesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListHapgsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListHapgsFuture) Get(ctx workflow.Context) (*cloudhsm.ListHapgsOutput, error) {
	var output cloudhsm.ListHapgsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListHsmsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListHsmsFuture) Get(ctx workflow.Context) (*cloudhsm.ListHsmsOutput, error) {
	var output cloudhsm.ListHsmsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListLunaClientsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListLunaClientsFuture) Get(ctx workflow.Context) (*cloudhsm.ListLunaClientsOutput, error) {
	var output cloudhsm.ListLunaClientsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListTagsForResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListTagsForResourceFuture) Get(ctx workflow.Context) (*cloudhsm.ListTagsForResourceOutput, error) {
	var output cloudhsm.ListTagsForResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ModifyHapgFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ModifyHapgFuture) Get(ctx workflow.Context) (*cloudhsm.ModifyHapgOutput, error) {
	var output cloudhsm.ModifyHapgOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ModifyHsmFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ModifyHsmFuture) Get(ctx workflow.Context) (*cloudhsm.ModifyHsmOutput, error) {
	var output cloudhsm.ModifyHsmOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ModifyLunaClientFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ModifyLunaClientFuture) Get(ctx workflow.Context) (*cloudhsm.ModifyLunaClientOutput, error) {
	var output cloudhsm.ModifyLunaClientOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type RemoveTagsFromResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *RemoveTagsFromResourceFuture) Get(ctx workflow.Context) (*cloudhsm.RemoveTagsFromResourceOutput, error) {
	var output cloudhsm.RemoveTagsFromResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *stub) AddTagsToResource(ctx workflow.Context, input *cloudhsm.AddTagsToResourceInput) (*cloudhsm.AddTagsToResourceOutput, error) {
	var output cloudhsm.AddTagsToResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws-cloudhsm-AddTagsToResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) AddTagsToResourceAsync(ctx workflow.Context, input *cloudhsm.AddTagsToResourceInput) *AddTagsToResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws-cloudhsm-AddTagsToResource", input)
	return &AddTagsToResourceFuture{Future: future}
}

func (a *stub) CreateHapg(ctx workflow.Context, input *cloudhsm.CreateHapgInput) (*cloudhsm.CreateHapgOutput, error) {
	var output cloudhsm.CreateHapgOutput
	err := workflow.ExecuteActivity(ctx, "aws-cloudhsm-CreateHapg", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateHapgAsync(ctx workflow.Context, input *cloudhsm.CreateHapgInput) *CreateHapgFuture {
	future := workflow.ExecuteActivity(ctx, "aws-cloudhsm-CreateHapg", input)
	return &CreateHapgFuture{Future: future}
}

func (a *stub) CreateHsm(ctx workflow.Context, input *cloudhsm.CreateHsmInput) (*cloudhsm.CreateHsmOutput, error) {
	var output cloudhsm.CreateHsmOutput
	err := workflow.ExecuteActivity(ctx, "aws-cloudhsm-CreateHsm", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateHsmAsync(ctx workflow.Context, input *cloudhsm.CreateHsmInput) *CreateHsmFuture {
	future := workflow.ExecuteActivity(ctx, "aws-cloudhsm-CreateHsm", input)
	return &CreateHsmFuture{Future: future}
}

func (a *stub) CreateLunaClient(ctx workflow.Context, input *cloudhsm.CreateLunaClientInput) (*cloudhsm.CreateLunaClientOutput, error) {
	var output cloudhsm.CreateLunaClientOutput
	err := workflow.ExecuteActivity(ctx, "aws-cloudhsm-CreateLunaClient", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateLunaClientAsync(ctx workflow.Context, input *cloudhsm.CreateLunaClientInput) *CreateLunaClientFuture {
	future := workflow.ExecuteActivity(ctx, "aws-cloudhsm-CreateLunaClient", input)
	return &CreateLunaClientFuture{Future: future}
}

func (a *stub) DeleteHapg(ctx workflow.Context, input *cloudhsm.DeleteHapgInput) (*cloudhsm.DeleteHapgOutput, error) {
	var output cloudhsm.DeleteHapgOutput
	err := workflow.ExecuteActivity(ctx, "aws-cloudhsm-DeleteHapg", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteHapgAsync(ctx workflow.Context, input *cloudhsm.DeleteHapgInput) *DeleteHapgFuture {
	future := workflow.ExecuteActivity(ctx, "aws-cloudhsm-DeleteHapg", input)
	return &DeleteHapgFuture{Future: future}
}

func (a *stub) DeleteHsm(ctx workflow.Context, input *cloudhsm.DeleteHsmInput) (*cloudhsm.DeleteHsmOutput, error) {
	var output cloudhsm.DeleteHsmOutput
	err := workflow.ExecuteActivity(ctx, "aws-cloudhsm-DeleteHsm", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteHsmAsync(ctx workflow.Context, input *cloudhsm.DeleteHsmInput) *DeleteHsmFuture {
	future := workflow.ExecuteActivity(ctx, "aws-cloudhsm-DeleteHsm", input)
	return &DeleteHsmFuture{Future: future}
}

func (a *stub) DeleteLunaClient(ctx workflow.Context, input *cloudhsm.DeleteLunaClientInput) (*cloudhsm.DeleteLunaClientOutput, error) {
	var output cloudhsm.DeleteLunaClientOutput
	err := workflow.ExecuteActivity(ctx, "aws-cloudhsm-DeleteLunaClient", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteLunaClientAsync(ctx workflow.Context, input *cloudhsm.DeleteLunaClientInput) *DeleteLunaClientFuture {
	future := workflow.ExecuteActivity(ctx, "aws-cloudhsm-DeleteLunaClient", input)
	return &DeleteLunaClientFuture{Future: future}
}

func (a *stub) DescribeHapg(ctx workflow.Context, input *cloudhsm.DescribeHapgInput) (*cloudhsm.DescribeHapgOutput, error) {
	var output cloudhsm.DescribeHapgOutput
	err := workflow.ExecuteActivity(ctx, "aws-cloudhsm-DescribeHapg", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeHapgAsync(ctx workflow.Context, input *cloudhsm.DescribeHapgInput) *DescribeHapgFuture {
	future := workflow.ExecuteActivity(ctx, "aws-cloudhsm-DescribeHapg", input)
	return &DescribeHapgFuture{Future: future}
}

func (a *stub) DescribeHsm(ctx workflow.Context, input *cloudhsm.DescribeHsmInput) (*cloudhsm.DescribeHsmOutput, error) {
	var output cloudhsm.DescribeHsmOutput
	err := workflow.ExecuteActivity(ctx, "aws-cloudhsm-DescribeHsm", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeHsmAsync(ctx workflow.Context, input *cloudhsm.DescribeHsmInput) *DescribeHsmFuture {
	future := workflow.ExecuteActivity(ctx, "aws-cloudhsm-DescribeHsm", input)
	return &DescribeHsmFuture{Future: future}
}

func (a *stub) DescribeLunaClient(ctx workflow.Context, input *cloudhsm.DescribeLunaClientInput) (*cloudhsm.DescribeLunaClientOutput, error) {
	var output cloudhsm.DescribeLunaClientOutput
	err := workflow.ExecuteActivity(ctx, "aws-cloudhsm-DescribeLunaClient", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeLunaClientAsync(ctx workflow.Context, input *cloudhsm.DescribeLunaClientInput) *DescribeLunaClientFuture {
	future := workflow.ExecuteActivity(ctx, "aws-cloudhsm-DescribeLunaClient", input)
	return &DescribeLunaClientFuture{Future: future}
}

func (a *stub) GetConfig(ctx workflow.Context, input *cloudhsm.GetConfigInput) (*cloudhsm.GetConfigOutput, error) {
	var output cloudhsm.GetConfigOutput
	err := workflow.ExecuteActivity(ctx, "aws-cloudhsm-GetConfig", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetConfigAsync(ctx workflow.Context, input *cloudhsm.GetConfigInput) *GetConfigFuture {
	future := workflow.ExecuteActivity(ctx, "aws-cloudhsm-GetConfig", input)
	return &GetConfigFuture{Future: future}
}

func (a *stub) ListAvailableZones(ctx workflow.Context, input *cloudhsm.ListAvailableZonesInput) (*cloudhsm.ListAvailableZonesOutput, error) {
	var output cloudhsm.ListAvailableZonesOutput
	err := workflow.ExecuteActivity(ctx, "aws-cloudhsm-ListAvailableZones", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListAvailableZonesAsync(ctx workflow.Context, input *cloudhsm.ListAvailableZonesInput) *ListAvailableZonesFuture {
	future := workflow.ExecuteActivity(ctx, "aws-cloudhsm-ListAvailableZones", input)
	return &ListAvailableZonesFuture{Future: future}
}

func (a *stub) ListHapgs(ctx workflow.Context, input *cloudhsm.ListHapgsInput) (*cloudhsm.ListHapgsOutput, error) {
	var output cloudhsm.ListHapgsOutput
	err := workflow.ExecuteActivity(ctx, "aws-cloudhsm-ListHapgs", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListHapgsAsync(ctx workflow.Context, input *cloudhsm.ListHapgsInput) *ListHapgsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-cloudhsm-ListHapgs", input)
	return &ListHapgsFuture{Future: future}
}

func (a *stub) ListHsms(ctx workflow.Context, input *cloudhsm.ListHsmsInput) (*cloudhsm.ListHsmsOutput, error) {
	var output cloudhsm.ListHsmsOutput
	err := workflow.ExecuteActivity(ctx, "aws-cloudhsm-ListHsms", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListHsmsAsync(ctx workflow.Context, input *cloudhsm.ListHsmsInput) *ListHsmsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-cloudhsm-ListHsms", input)
	return &ListHsmsFuture{Future: future}
}

func (a *stub) ListLunaClients(ctx workflow.Context, input *cloudhsm.ListLunaClientsInput) (*cloudhsm.ListLunaClientsOutput, error) {
	var output cloudhsm.ListLunaClientsOutput
	err := workflow.ExecuteActivity(ctx, "aws-cloudhsm-ListLunaClients", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListLunaClientsAsync(ctx workflow.Context, input *cloudhsm.ListLunaClientsInput) *ListLunaClientsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-cloudhsm-ListLunaClients", input)
	return &ListLunaClientsFuture{Future: future}
}

func (a *stub) ListTagsForResource(ctx workflow.Context, input *cloudhsm.ListTagsForResourceInput) (*cloudhsm.ListTagsForResourceOutput, error) {
	var output cloudhsm.ListTagsForResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws-cloudhsm-ListTagsForResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListTagsForResourceAsync(ctx workflow.Context, input *cloudhsm.ListTagsForResourceInput) *ListTagsForResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws-cloudhsm-ListTagsForResource", input)
	return &ListTagsForResourceFuture{Future: future}
}

func (a *stub) ModifyHapg(ctx workflow.Context, input *cloudhsm.ModifyHapgInput) (*cloudhsm.ModifyHapgOutput, error) {
	var output cloudhsm.ModifyHapgOutput
	err := workflow.ExecuteActivity(ctx, "aws-cloudhsm-ModifyHapg", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ModifyHapgAsync(ctx workflow.Context, input *cloudhsm.ModifyHapgInput) *ModifyHapgFuture {
	future := workflow.ExecuteActivity(ctx, "aws-cloudhsm-ModifyHapg", input)
	return &ModifyHapgFuture{Future: future}
}

func (a *stub) ModifyHsm(ctx workflow.Context, input *cloudhsm.ModifyHsmInput) (*cloudhsm.ModifyHsmOutput, error) {
	var output cloudhsm.ModifyHsmOutput
	err := workflow.ExecuteActivity(ctx, "aws-cloudhsm-ModifyHsm", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ModifyHsmAsync(ctx workflow.Context, input *cloudhsm.ModifyHsmInput) *ModifyHsmFuture {
	future := workflow.ExecuteActivity(ctx, "aws-cloudhsm-ModifyHsm", input)
	return &ModifyHsmFuture{Future: future}
}

func (a *stub) ModifyLunaClient(ctx workflow.Context, input *cloudhsm.ModifyLunaClientInput) (*cloudhsm.ModifyLunaClientOutput, error) {
	var output cloudhsm.ModifyLunaClientOutput
	err := workflow.ExecuteActivity(ctx, "aws-cloudhsm-ModifyLunaClient", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ModifyLunaClientAsync(ctx workflow.Context, input *cloudhsm.ModifyLunaClientInput) *ModifyLunaClientFuture {
	future := workflow.ExecuteActivity(ctx, "aws-cloudhsm-ModifyLunaClient", input)
	return &ModifyLunaClientFuture{Future: future}
}

func (a *stub) RemoveTagsFromResource(ctx workflow.Context, input *cloudhsm.RemoveTagsFromResourceInput) (*cloudhsm.RemoveTagsFromResourceOutput, error) {
	var output cloudhsm.RemoveTagsFromResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws-cloudhsm-RemoveTagsFromResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) RemoveTagsFromResourceAsync(ctx workflow.Context, input *cloudhsm.RemoveTagsFromResourceInput) *RemoveTagsFromResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws-cloudhsm-RemoveTagsFromResource", input)
	return &RemoveTagsFromResourceFuture{Future: future}
}
