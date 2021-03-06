// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7

package detectivestub

import (
	"github.com/aws/aws-sdk-go/service/detective"
	"go.uber.org/cadence/workflow"

	"github.com/banzaicloud/cadence-aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type stub struct{}

type AcceptInvitationFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AcceptInvitationFuture) Get(ctx workflow.Context) (*detective.AcceptInvitationOutput, error) {
	var output detective.AcceptInvitationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CreateGraphFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CreateGraphFuture) Get(ctx workflow.Context) (*detective.CreateGraphOutput, error) {
	var output detective.CreateGraphOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CreateMembersFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CreateMembersFuture) Get(ctx workflow.Context) (*detective.CreateMembersOutput, error) {
	var output detective.CreateMembersOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DeleteGraphFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DeleteGraphFuture) Get(ctx workflow.Context) (*detective.DeleteGraphOutput, error) {
	var output detective.DeleteGraphOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DeleteMembersFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DeleteMembersFuture) Get(ctx workflow.Context) (*detective.DeleteMembersOutput, error) {
	var output detective.DeleteMembersOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DisassociateMembershipFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DisassociateMembershipFuture) Get(ctx workflow.Context) (*detective.DisassociateMembershipOutput, error) {
	var output detective.DisassociateMembershipOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GetMembersFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GetMembersFuture) Get(ctx workflow.Context) (*detective.GetMembersOutput, error) {
	var output detective.GetMembersOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListGraphsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListGraphsFuture) Get(ctx workflow.Context) (*detective.ListGraphsOutput, error) {
	var output detective.ListGraphsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListInvitationsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListInvitationsFuture) Get(ctx workflow.Context) (*detective.ListInvitationsOutput, error) {
	var output detective.ListInvitationsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListMembersFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListMembersFuture) Get(ctx workflow.Context) (*detective.ListMembersOutput, error) {
	var output detective.ListMembersOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type RejectInvitationFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *RejectInvitationFuture) Get(ctx workflow.Context) (*detective.RejectInvitationOutput, error) {
	var output detective.RejectInvitationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type StartMonitoringMemberFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *StartMonitoringMemberFuture) Get(ctx workflow.Context) (*detective.StartMonitoringMemberOutput, error) {
	var output detective.StartMonitoringMemberOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *stub) AcceptInvitation(ctx workflow.Context, input *detective.AcceptInvitationInput) (*detective.AcceptInvitationOutput, error) {
	var output detective.AcceptInvitationOutput
	err := workflow.ExecuteActivity(ctx, "aws-detective-AcceptInvitation", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) AcceptInvitationAsync(ctx workflow.Context, input *detective.AcceptInvitationInput) *AcceptInvitationFuture {
	future := workflow.ExecuteActivity(ctx, "aws-detective-AcceptInvitation", input)
	return &AcceptInvitationFuture{Future: future}
}

func (a *stub) CreateGraph(ctx workflow.Context, input *detective.CreateGraphInput) (*detective.CreateGraphOutput, error) {
	var output detective.CreateGraphOutput
	err := workflow.ExecuteActivity(ctx, "aws-detective-CreateGraph", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateGraphAsync(ctx workflow.Context, input *detective.CreateGraphInput) *CreateGraphFuture {
	future := workflow.ExecuteActivity(ctx, "aws-detective-CreateGraph", input)
	return &CreateGraphFuture{Future: future}
}

func (a *stub) CreateMembers(ctx workflow.Context, input *detective.CreateMembersInput) (*detective.CreateMembersOutput, error) {
	var output detective.CreateMembersOutput
	err := workflow.ExecuteActivity(ctx, "aws-detective-CreateMembers", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateMembersAsync(ctx workflow.Context, input *detective.CreateMembersInput) *CreateMembersFuture {
	future := workflow.ExecuteActivity(ctx, "aws-detective-CreateMembers", input)
	return &CreateMembersFuture{Future: future}
}

func (a *stub) DeleteGraph(ctx workflow.Context, input *detective.DeleteGraphInput) (*detective.DeleteGraphOutput, error) {
	var output detective.DeleteGraphOutput
	err := workflow.ExecuteActivity(ctx, "aws-detective-DeleteGraph", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteGraphAsync(ctx workflow.Context, input *detective.DeleteGraphInput) *DeleteGraphFuture {
	future := workflow.ExecuteActivity(ctx, "aws-detective-DeleteGraph", input)
	return &DeleteGraphFuture{Future: future}
}

func (a *stub) DeleteMembers(ctx workflow.Context, input *detective.DeleteMembersInput) (*detective.DeleteMembersOutput, error) {
	var output detective.DeleteMembersOutput
	err := workflow.ExecuteActivity(ctx, "aws-detective-DeleteMembers", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteMembersAsync(ctx workflow.Context, input *detective.DeleteMembersInput) *DeleteMembersFuture {
	future := workflow.ExecuteActivity(ctx, "aws-detective-DeleteMembers", input)
	return &DeleteMembersFuture{Future: future}
}

func (a *stub) DisassociateMembership(ctx workflow.Context, input *detective.DisassociateMembershipInput) (*detective.DisassociateMembershipOutput, error) {
	var output detective.DisassociateMembershipOutput
	err := workflow.ExecuteActivity(ctx, "aws-detective-DisassociateMembership", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DisassociateMembershipAsync(ctx workflow.Context, input *detective.DisassociateMembershipInput) *DisassociateMembershipFuture {
	future := workflow.ExecuteActivity(ctx, "aws-detective-DisassociateMembership", input)
	return &DisassociateMembershipFuture{Future: future}
}

func (a *stub) GetMembers(ctx workflow.Context, input *detective.GetMembersInput) (*detective.GetMembersOutput, error) {
	var output detective.GetMembersOutput
	err := workflow.ExecuteActivity(ctx, "aws-detective-GetMembers", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetMembersAsync(ctx workflow.Context, input *detective.GetMembersInput) *GetMembersFuture {
	future := workflow.ExecuteActivity(ctx, "aws-detective-GetMembers", input)
	return &GetMembersFuture{Future: future}
}

func (a *stub) ListGraphs(ctx workflow.Context, input *detective.ListGraphsInput) (*detective.ListGraphsOutput, error) {
	var output detective.ListGraphsOutput
	err := workflow.ExecuteActivity(ctx, "aws-detective-ListGraphs", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListGraphsAsync(ctx workflow.Context, input *detective.ListGraphsInput) *ListGraphsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-detective-ListGraphs", input)
	return &ListGraphsFuture{Future: future}
}

func (a *stub) ListInvitations(ctx workflow.Context, input *detective.ListInvitationsInput) (*detective.ListInvitationsOutput, error) {
	var output detective.ListInvitationsOutput
	err := workflow.ExecuteActivity(ctx, "aws-detective-ListInvitations", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListInvitationsAsync(ctx workflow.Context, input *detective.ListInvitationsInput) *ListInvitationsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-detective-ListInvitations", input)
	return &ListInvitationsFuture{Future: future}
}

func (a *stub) ListMembers(ctx workflow.Context, input *detective.ListMembersInput) (*detective.ListMembersOutput, error) {
	var output detective.ListMembersOutput
	err := workflow.ExecuteActivity(ctx, "aws-detective-ListMembers", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListMembersAsync(ctx workflow.Context, input *detective.ListMembersInput) *ListMembersFuture {
	future := workflow.ExecuteActivity(ctx, "aws-detective-ListMembers", input)
	return &ListMembersFuture{Future: future}
}

func (a *stub) RejectInvitation(ctx workflow.Context, input *detective.RejectInvitationInput) (*detective.RejectInvitationOutput, error) {
	var output detective.RejectInvitationOutput
	err := workflow.ExecuteActivity(ctx, "aws-detective-RejectInvitation", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) RejectInvitationAsync(ctx workflow.Context, input *detective.RejectInvitationInput) *RejectInvitationFuture {
	future := workflow.ExecuteActivity(ctx, "aws-detective-RejectInvitation", input)
	return &RejectInvitationFuture{Future: future}
}

func (a *stub) StartMonitoringMember(ctx workflow.Context, input *detective.StartMonitoringMemberInput) (*detective.StartMonitoringMemberOutput, error) {
	var output detective.StartMonitoringMemberOutput
	err := workflow.ExecuteActivity(ctx, "aws-detective-StartMonitoringMember", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) StartMonitoringMemberAsync(ctx workflow.Context, input *detective.StartMonitoringMemberInput) *StartMonitoringMemberFuture {
	future := workflow.ExecuteActivity(ctx, "aws-detective-StartMonitoringMember", input)
	return &StartMonitoringMemberFuture{Future: future}
}
