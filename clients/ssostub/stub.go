// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.
// Copyright (c) 2020 Banzai Cloud All rights reserved.

package ssostub

import (
	"github.com/aws/aws-sdk-go/service/sso"
	"go.uber.org/cadence/workflow"

	"github.com/banzaicloud/cadence-aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type stub struct{}

type SSOGetRoleCredentialsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SSOGetRoleCredentialsFuture) Get(ctx workflow.Context) (*sso.GetRoleCredentialsOutput, error) {
	var output sso.GetRoleCredentialsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SSOListAccountRolesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SSOListAccountRolesFuture) Get(ctx workflow.Context) (*sso.ListAccountRolesOutput, error) {
	var output sso.ListAccountRolesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SSOListAccountsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SSOListAccountsFuture) Get(ctx workflow.Context) (*sso.ListAccountsOutput, error) {
	var output sso.ListAccountsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SSOLogoutFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SSOLogoutFuture) Get(ctx workflow.Context) (*sso.LogoutOutput, error) {
	var output sso.LogoutOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *stub) GetRoleCredentials(ctx workflow.Context, input *sso.GetRoleCredentialsInput) (*sso.GetRoleCredentialsOutput, error) {
	var output sso.GetRoleCredentialsOutput
	err := workflow.ExecuteActivity(ctx, "aws-sso-GetRoleCredentials", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetRoleCredentialsAsync(ctx workflow.Context, input *sso.GetRoleCredentialsInput) *SSOGetRoleCredentialsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-sso-GetRoleCredentials", input)
	return &SSOGetRoleCredentialsFuture{Future: future}
}

func (a *stub) ListAccountRoles(ctx workflow.Context, input *sso.ListAccountRolesInput) (*sso.ListAccountRolesOutput, error) {
	var output sso.ListAccountRolesOutput
	err := workflow.ExecuteActivity(ctx, "aws-sso-ListAccountRoles", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListAccountRolesAsync(ctx workflow.Context, input *sso.ListAccountRolesInput) *SSOListAccountRolesFuture {
	future := workflow.ExecuteActivity(ctx, "aws-sso-ListAccountRoles", input)
	return &SSOListAccountRolesFuture{Future: future}
}

func (a *stub) ListAccounts(ctx workflow.Context, input *sso.ListAccountsInput) (*sso.ListAccountsOutput, error) {
	var output sso.ListAccountsOutput
	err := workflow.ExecuteActivity(ctx, "aws-sso-ListAccounts", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListAccountsAsync(ctx workflow.Context, input *sso.ListAccountsInput) *SSOListAccountsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-sso-ListAccounts", input)
	return &SSOListAccountsFuture{Future: future}
}

func (a *stub) Logout(ctx workflow.Context, input *sso.LogoutInput) (*sso.LogoutOutput, error) {
	var output sso.LogoutOutput
	err := workflow.ExecuteActivity(ctx, "aws-sso-Logout", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) LogoutAsync(ctx workflow.Context, input *sso.LogoutInput) *SSOLogoutFuture {
	future := workflow.ExecuteActivity(ctx, "aws-sso-Logout", input)
	return &SSOLogoutFuture{Future: future}
}
