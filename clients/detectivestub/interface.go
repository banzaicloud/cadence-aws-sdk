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

type Client interface {
	AcceptInvitation(ctx workflow.Context, input *detective.AcceptInvitationInput) (*detective.AcceptInvitationOutput, error)
	AcceptInvitationAsync(ctx workflow.Context, input *detective.AcceptInvitationInput) *AcceptInvitationFuture

	CreateGraph(ctx workflow.Context, input *detective.CreateGraphInput) (*detective.CreateGraphOutput, error)
	CreateGraphAsync(ctx workflow.Context, input *detective.CreateGraphInput) *CreateGraphFuture

	CreateMembers(ctx workflow.Context, input *detective.CreateMembersInput) (*detective.CreateMembersOutput, error)
	CreateMembersAsync(ctx workflow.Context, input *detective.CreateMembersInput) *CreateMembersFuture

	DeleteGraph(ctx workflow.Context, input *detective.DeleteGraphInput) (*detective.DeleteGraphOutput, error)
	DeleteGraphAsync(ctx workflow.Context, input *detective.DeleteGraphInput) *DeleteGraphFuture

	DeleteMembers(ctx workflow.Context, input *detective.DeleteMembersInput) (*detective.DeleteMembersOutput, error)
	DeleteMembersAsync(ctx workflow.Context, input *detective.DeleteMembersInput) *DeleteMembersFuture

	DisassociateMembership(ctx workflow.Context, input *detective.DisassociateMembershipInput) (*detective.DisassociateMembershipOutput, error)
	DisassociateMembershipAsync(ctx workflow.Context, input *detective.DisassociateMembershipInput) *DisassociateMembershipFuture

	GetMembers(ctx workflow.Context, input *detective.GetMembersInput) (*detective.GetMembersOutput, error)
	GetMembersAsync(ctx workflow.Context, input *detective.GetMembersInput) *GetMembersFuture

	ListGraphs(ctx workflow.Context, input *detective.ListGraphsInput) (*detective.ListGraphsOutput, error)
	ListGraphsAsync(ctx workflow.Context, input *detective.ListGraphsInput) *ListGraphsFuture

	ListInvitations(ctx workflow.Context, input *detective.ListInvitationsInput) (*detective.ListInvitationsOutput, error)
	ListInvitationsAsync(ctx workflow.Context, input *detective.ListInvitationsInput) *ListInvitationsFuture

	ListMembers(ctx workflow.Context, input *detective.ListMembersInput) (*detective.ListMembersOutput, error)
	ListMembersAsync(ctx workflow.Context, input *detective.ListMembersInput) *ListMembersFuture

	RejectInvitation(ctx workflow.Context, input *detective.RejectInvitationInput) (*detective.RejectInvitationOutput, error)
	RejectInvitationAsync(ctx workflow.Context, input *detective.RejectInvitationInput) *RejectInvitationFuture

	StartMonitoringMember(ctx workflow.Context, input *detective.StartMonitoringMemberInput) (*detective.StartMonitoringMemberOutput, error)
	StartMonitoringMemberAsync(ctx workflow.Context, input *detective.StartMonitoringMemberInput) *StartMonitoringMemberFuture
}

func NewClient() Client {
	return &stub{}
}
