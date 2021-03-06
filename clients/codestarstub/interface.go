// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7

package codestarstub

import (
	"github.com/aws/aws-sdk-go/service/codestar"
	"go.uber.org/cadence/workflow"

	"github.com/banzaicloud/cadence-aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	AssociateTeamMember(ctx workflow.Context, input *codestar.AssociateTeamMemberInput) (*codestar.AssociateTeamMemberOutput, error)
	AssociateTeamMemberAsync(ctx workflow.Context, input *codestar.AssociateTeamMemberInput) *AssociateTeamMemberFuture

	CreateProject(ctx workflow.Context, input *codestar.CreateProjectInput) (*codestar.CreateProjectOutput, error)
	CreateProjectAsync(ctx workflow.Context, input *codestar.CreateProjectInput) *CreateProjectFuture

	CreateUserProfile(ctx workflow.Context, input *codestar.CreateUserProfileInput) (*codestar.CreateUserProfileOutput, error)
	CreateUserProfileAsync(ctx workflow.Context, input *codestar.CreateUserProfileInput) *CreateUserProfileFuture

	DeleteProject(ctx workflow.Context, input *codestar.DeleteProjectInput) (*codestar.DeleteProjectOutput, error)
	DeleteProjectAsync(ctx workflow.Context, input *codestar.DeleteProjectInput) *DeleteProjectFuture

	DeleteUserProfile(ctx workflow.Context, input *codestar.DeleteUserProfileInput) (*codestar.DeleteUserProfileOutput, error)
	DeleteUserProfileAsync(ctx workflow.Context, input *codestar.DeleteUserProfileInput) *DeleteUserProfileFuture

	DescribeProject(ctx workflow.Context, input *codestar.DescribeProjectInput) (*codestar.DescribeProjectOutput, error)
	DescribeProjectAsync(ctx workflow.Context, input *codestar.DescribeProjectInput) *DescribeProjectFuture

	DescribeUserProfile(ctx workflow.Context, input *codestar.DescribeUserProfileInput) (*codestar.DescribeUserProfileOutput, error)
	DescribeUserProfileAsync(ctx workflow.Context, input *codestar.DescribeUserProfileInput) *DescribeUserProfileFuture

	DisassociateTeamMember(ctx workflow.Context, input *codestar.DisassociateTeamMemberInput) (*codestar.DisassociateTeamMemberOutput, error)
	DisassociateTeamMemberAsync(ctx workflow.Context, input *codestar.DisassociateTeamMemberInput) *DisassociateTeamMemberFuture

	ListProjects(ctx workflow.Context, input *codestar.ListProjectsInput) (*codestar.ListProjectsOutput, error)
	ListProjectsAsync(ctx workflow.Context, input *codestar.ListProjectsInput) *ListProjectsFuture

	ListResources(ctx workflow.Context, input *codestar.ListResourcesInput) (*codestar.ListResourcesOutput, error)
	ListResourcesAsync(ctx workflow.Context, input *codestar.ListResourcesInput) *ListResourcesFuture

	ListTagsForProject(ctx workflow.Context, input *codestar.ListTagsForProjectInput) (*codestar.ListTagsForProjectOutput, error)
	ListTagsForProjectAsync(ctx workflow.Context, input *codestar.ListTagsForProjectInput) *ListTagsForProjectFuture

	ListTeamMembers(ctx workflow.Context, input *codestar.ListTeamMembersInput) (*codestar.ListTeamMembersOutput, error)
	ListTeamMembersAsync(ctx workflow.Context, input *codestar.ListTeamMembersInput) *ListTeamMembersFuture

	ListUserProfiles(ctx workflow.Context, input *codestar.ListUserProfilesInput) (*codestar.ListUserProfilesOutput, error)
	ListUserProfilesAsync(ctx workflow.Context, input *codestar.ListUserProfilesInput) *ListUserProfilesFuture

	TagProject(ctx workflow.Context, input *codestar.TagProjectInput) (*codestar.TagProjectOutput, error)
	TagProjectAsync(ctx workflow.Context, input *codestar.TagProjectInput) *TagProjectFuture

	UntagProject(ctx workflow.Context, input *codestar.UntagProjectInput) (*codestar.UntagProjectOutput, error)
	UntagProjectAsync(ctx workflow.Context, input *codestar.UntagProjectInput) *UntagProjectFuture

	UpdateProject(ctx workflow.Context, input *codestar.UpdateProjectInput) (*codestar.UpdateProjectOutput, error)
	UpdateProjectAsync(ctx workflow.Context, input *codestar.UpdateProjectInput) *UpdateProjectFuture

	UpdateTeamMember(ctx workflow.Context, input *codestar.UpdateTeamMemberInput) (*codestar.UpdateTeamMemberOutput, error)
	UpdateTeamMemberAsync(ctx workflow.Context, input *codestar.UpdateTeamMemberInput) *UpdateTeamMemberFuture

	UpdateUserProfile(ctx workflow.Context, input *codestar.UpdateUserProfileInput) (*codestar.UpdateUserProfileOutput, error)
	UpdateUserProfileAsync(ctx workflow.Context, input *codestar.UpdateUserProfileInput) *UpdateUserProfileFuture
}

func NewClient() Client {
	return &stub{}
}
