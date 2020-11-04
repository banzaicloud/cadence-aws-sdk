// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.
// Copyright (c) 2020 Banzai Cloud All rights reserved.

package appconfigstub

import (
	"github.com/aws/aws-sdk-go/service/appconfig"
	"go.uber.org/cadence/workflow"

	"github.com/banzaicloud/cadence-aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type stub struct{}

type AppConfigCreateApplicationFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AppConfigCreateApplicationFuture) Get(ctx workflow.Context) (*appconfig.CreateApplicationOutput, error) {
	var output appconfig.CreateApplicationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AppConfigCreateConfigurationProfileFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AppConfigCreateConfigurationProfileFuture) Get(ctx workflow.Context) (*appconfig.CreateConfigurationProfileOutput, error) {
	var output appconfig.CreateConfigurationProfileOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AppConfigCreateDeploymentStrategyFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AppConfigCreateDeploymentStrategyFuture) Get(ctx workflow.Context) (*appconfig.CreateDeploymentStrategyOutput, error) {
	var output appconfig.CreateDeploymentStrategyOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AppConfigCreateEnvironmentFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AppConfigCreateEnvironmentFuture) Get(ctx workflow.Context) (*appconfig.CreateEnvironmentOutput, error) {
	var output appconfig.CreateEnvironmentOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AppConfigCreateHostedConfigurationVersionFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AppConfigCreateHostedConfigurationVersionFuture) Get(ctx workflow.Context) (*appconfig.CreateHostedConfigurationVersionOutput, error) {
	var output appconfig.CreateHostedConfigurationVersionOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AppConfigDeleteApplicationFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AppConfigDeleteApplicationFuture) Get(ctx workflow.Context) (*appconfig.DeleteApplicationOutput, error) {
	var output appconfig.DeleteApplicationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AppConfigDeleteConfigurationProfileFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AppConfigDeleteConfigurationProfileFuture) Get(ctx workflow.Context) (*appconfig.DeleteConfigurationProfileOutput, error) {
	var output appconfig.DeleteConfigurationProfileOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AppConfigDeleteDeploymentStrategyFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AppConfigDeleteDeploymentStrategyFuture) Get(ctx workflow.Context) (*appconfig.DeleteDeploymentStrategyOutput, error) {
	var output appconfig.DeleteDeploymentStrategyOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AppConfigDeleteEnvironmentFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AppConfigDeleteEnvironmentFuture) Get(ctx workflow.Context) (*appconfig.DeleteEnvironmentOutput, error) {
	var output appconfig.DeleteEnvironmentOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AppConfigDeleteHostedConfigurationVersionFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AppConfigDeleteHostedConfigurationVersionFuture) Get(ctx workflow.Context) (*appconfig.DeleteHostedConfigurationVersionOutput, error) {
	var output appconfig.DeleteHostedConfigurationVersionOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AppConfigGetApplicationFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AppConfigGetApplicationFuture) Get(ctx workflow.Context) (*appconfig.GetApplicationOutput, error) {
	var output appconfig.GetApplicationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AppConfigGetConfigurationFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AppConfigGetConfigurationFuture) Get(ctx workflow.Context) (*appconfig.GetConfigurationOutput, error) {
	var output appconfig.GetConfigurationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AppConfigGetConfigurationProfileFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AppConfigGetConfigurationProfileFuture) Get(ctx workflow.Context) (*appconfig.GetConfigurationProfileOutput, error) {
	var output appconfig.GetConfigurationProfileOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AppConfigGetDeploymentFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AppConfigGetDeploymentFuture) Get(ctx workflow.Context) (*appconfig.GetDeploymentOutput, error) {
	var output appconfig.GetDeploymentOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AppConfigGetDeploymentStrategyFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AppConfigGetDeploymentStrategyFuture) Get(ctx workflow.Context) (*appconfig.GetDeploymentStrategyOutput, error) {
	var output appconfig.GetDeploymentStrategyOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AppConfigGetEnvironmentFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AppConfigGetEnvironmentFuture) Get(ctx workflow.Context) (*appconfig.GetEnvironmentOutput, error) {
	var output appconfig.GetEnvironmentOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AppConfigGetHostedConfigurationVersionFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AppConfigGetHostedConfigurationVersionFuture) Get(ctx workflow.Context) (*appconfig.GetHostedConfigurationVersionOutput, error) {
	var output appconfig.GetHostedConfigurationVersionOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AppConfigListApplicationsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AppConfigListApplicationsFuture) Get(ctx workflow.Context) (*appconfig.ListApplicationsOutput, error) {
	var output appconfig.ListApplicationsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AppConfigListConfigurationProfilesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AppConfigListConfigurationProfilesFuture) Get(ctx workflow.Context) (*appconfig.ListConfigurationProfilesOutput, error) {
	var output appconfig.ListConfigurationProfilesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AppConfigListDeploymentStrategiesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AppConfigListDeploymentStrategiesFuture) Get(ctx workflow.Context) (*appconfig.ListDeploymentStrategiesOutput, error) {
	var output appconfig.ListDeploymentStrategiesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AppConfigListDeploymentsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AppConfigListDeploymentsFuture) Get(ctx workflow.Context) (*appconfig.ListDeploymentsOutput, error) {
	var output appconfig.ListDeploymentsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AppConfigListEnvironmentsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AppConfigListEnvironmentsFuture) Get(ctx workflow.Context) (*appconfig.ListEnvironmentsOutput, error) {
	var output appconfig.ListEnvironmentsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AppConfigListHostedConfigurationVersionsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AppConfigListHostedConfigurationVersionsFuture) Get(ctx workflow.Context) (*appconfig.ListHostedConfigurationVersionsOutput, error) {
	var output appconfig.ListHostedConfigurationVersionsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AppConfigListTagsForResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AppConfigListTagsForResourceFuture) Get(ctx workflow.Context) (*appconfig.ListTagsForResourceOutput, error) {
	var output appconfig.ListTagsForResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AppConfigStartDeploymentFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AppConfigStartDeploymentFuture) Get(ctx workflow.Context) (*appconfig.StartDeploymentOutput, error) {
	var output appconfig.StartDeploymentOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AppConfigStopDeploymentFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AppConfigStopDeploymentFuture) Get(ctx workflow.Context) (*appconfig.StopDeploymentOutput, error) {
	var output appconfig.StopDeploymentOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AppConfigTagResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AppConfigTagResourceFuture) Get(ctx workflow.Context) (*appconfig.TagResourceOutput, error) {
	var output appconfig.TagResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AppConfigUntagResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AppConfigUntagResourceFuture) Get(ctx workflow.Context) (*appconfig.UntagResourceOutput, error) {
	var output appconfig.UntagResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AppConfigUpdateApplicationFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AppConfigUpdateApplicationFuture) Get(ctx workflow.Context) (*appconfig.UpdateApplicationOutput, error) {
	var output appconfig.UpdateApplicationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AppConfigUpdateConfigurationProfileFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AppConfigUpdateConfigurationProfileFuture) Get(ctx workflow.Context) (*appconfig.UpdateConfigurationProfileOutput, error) {
	var output appconfig.UpdateConfigurationProfileOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AppConfigUpdateDeploymentStrategyFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AppConfigUpdateDeploymentStrategyFuture) Get(ctx workflow.Context) (*appconfig.UpdateDeploymentStrategyOutput, error) {
	var output appconfig.UpdateDeploymentStrategyOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AppConfigUpdateEnvironmentFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AppConfigUpdateEnvironmentFuture) Get(ctx workflow.Context) (*appconfig.UpdateEnvironmentOutput, error) {
	var output appconfig.UpdateEnvironmentOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AppConfigValidateConfigurationFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AppConfigValidateConfigurationFuture) Get(ctx workflow.Context) (*appconfig.ValidateConfigurationOutput, error) {
	var output appconfig.ValidateConfigurationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateApplication(ctx workflow.Context, input *appconfig.CreateApplicationInput) (*appconfig.CreateApplicationOutput, error) {
	var output appconfig.CreateApplicationOutput
	err := workflow.ExecuteActivity(ctx, "aws-appconfig-CreateApplication", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateApplicationAsync(ctx workflow.Context, input *appconfig.CreateApplicationInput) *AppConfigCreateApplicationFuture {
	future := workflow.ExecuteActivity(ctx, "aws-appconfig-CreateApplication", input)
	return &AppConfigCreateApplicationFuture{Future: future}
}

func (a *stub) CreateConfigurationProfile(ctx workflow.Context, input *appconfig.CreateConfigurationProfileInput) (*appconfig.CreateConfigurationProfileOutput, error) {
	var output appconfig.CreateConfigurationProfileOutput
	err := workflow.ExecuteActivity(ctx, "aws-appconfig-CreateConfigurationProfile", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateConfigurationProfileAsync(ctx workflow.Context, input *appconfig.CreateConfigurationProfileInput) *AppConfigCreateConfigurationProfileFuture {
	future := workflow.ExecuteActivity(ctx, "aws-appconfig-CreateConfigurationProfile", input)
	return &AppConfigCreateConfigurationProfileFuture{Future: future}
}

func (a *stub) CreateDeploymentStrategy(ctx workflow.Context, input *appconfig.CreateDeploymentStrategyInput) (*appconfig.CreateDeploymentStrategyOutput, error) {
	var output appconfig.CreateDeploymentStrategyOutput
	err := workflow.ExecuteActivity(ctx, "aws-appconfig-CreateDeploymentStrategy", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateDeploymentStrategyAsync(ctx workflow.Context, input *appconfig.CreateDeploymentStrategyInput) *AppConfigCreateDeploymentStrategyFuture {
	future := workflow.ExecuteActivity(ctx, "aws-appconfig-CreateDeploymentStrategy", input)
	return &AppConfigCreateDeploymentStrategyFuture{Future: future}
}

func (a *stub) CreateEnvironment(ctx workflow.Context, input *appconfig.CreateEnvironmentInput) (*appconfig.CreateEnvironmentOutput, error) {
	var output appconfig.CreateEnvironmentOutput
	err := workflow.ExecuteActivity(ctx, "aws-appconfig-CreateEnvironment", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateEnvironmentAsync(ctx workflow.Context, input *appconfig.CreateEnvironmentInput) *AppConfigCreateEnvironmentFuture {
	future := workflow.ExecuteActivity(ctx, "aws-appconfig-CreateEnvironment", input)
	return &AppConfigCreateEnvironmentFuture{Future: future}
}

func (a *stub) CreateHostedConfigurationVersion(ctx workflow.Context, input *appconfig.CreateHostedConfigurationVersionInput) (*appconfig.CreateHostedConfigurationVersionOutput, error) {
	var output appconfig.CreateHostedConfigurationVersionOutput
	err := workflow.ExecuteActivity(ctx, "aws-appconfig-CreateHostedConfigurationVersion", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateHostedConfigurationVersionAsync(ctx workflow.Context, input *appconfig.CreateHostedConfigurationVersionInput) *AppConfigCreateHostedConfigurationVersionFuture {
	future := workflow.ExecuteActivity(ctx, "aws-appconfig-CreateHostedConfigurationVersion", input)
	return &AppConfigCreateHostedConfigurationVersionFuture{Future: future}
}

func (a *stub) DeleteApplication(ctx workflow.Context, input *appconfig.DeleteApplicationInput) (*appconfig.DeleteApplicationOutput, error) {
	var output appconfig.DeleteApplicationOutput
	err := workflow.ExecuteActivity(ctx, "aws-appconfig-DeleteApplication", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteApplicationAsync(ctx workflow.Context, input *appconfig.DeleteApplicationInput) *AppConfigDeleteApplicationFuture {
	future := workflow.ExecuteActivity(ctx, "aws-appconfig-DeleteApplication", input)
	return &AppConfigDeleteApplicationFuture{Future: future}
}

func (a *stub) DeleteConfigurationProfile(ctx workflow.Context, input *appconfig.DeleteConfigurationProfileInput) (*appconfig.DeleteConfigurationProfileOutput, error) {
	var output appconfig.DeleteConfigurationProfileOutput
	err := workflow.ExecuteActivity(ctx, "aws-appconfig-DeleteConfigurationProfile", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteConfigurationProfileAsync(ctx workflow.Context, input *appconfig.DeleteConfigurationProfileInput) *AppConfigDeleteConfigurationProfileFuture {
	future := workflow.ExecuteActivity(ctx, "aws-appconfig-DeleteConfigurationProfile", input)
	return &AppConfigDeleteConfigurationProfileFuture{Future: future}
}

func (a *stub) DeleteDeploymentStrategy(ctx workflow.Context, input *appconfig.DeleteDeploymentStrategyInput) (*appconfig.DeleteDeploymentStrategyOutput, error) {
	var output appconfig.DeleteDeploymentStrategyOutput
	err := workflow.ExecuteActivity(ctx, "aws-appconfig-DeleteDeploymentStrategy", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteDeploymentStrategyAsync(ctx workflow.Context, input *appconfig.DeleteDeploymentStrategyInput) *AppConfigDeleteDeploymentStrategyFuture {
	future := workflow.ExecuteActivity(ctx, "aws-appconfig-DeleteDeploymentStrategy", input)
	return &AppConfigDeleteDeploymentStrategyFuture{Future: future}
}

func (a *stub) DeleteEnvironment(ctx workflow.Context, input *appconfig.DeleteEnvironmentInput) (*appconfig.DeleteEnvironmentOutput, error) {
	var output appconfig.DeleteEnvironmentOutput
	err := workflow.ExecuteActivity(ctx, "aws-appconfig-DeleteEnvironment", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteEnvironmentAsync(ctx workflow.Context, input *appconfig.DeleteEnvironmentInput) *AppConfigDeleteEnvironmentFuture {
	future := workflow.ExecuteActivity(ctx, "aws-appconfig-DeleteEnvironment", input)
	return &AppConfigDeleteEnvironmentFuture{Future: future}
}

func (a *stub) DeleteHostedConfigurationVersion(ctx workflow.Context, input *appconfig.DeleteHostedConfigurationVersionInput) (*appconfig.DeleteHostedConfigurationVersionOutput, error) {
	var output appconfig.DeleteHostedConfigurationVersionOutput
	err := workflow.ExecuteActivity(ctx, "aws-appconfig-DeleteHostedConfigurationVersion", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteHostedConfigurationVersionAsync(ctx workflow.Context, input *appconfig.DeleteHostedConfigurationVersionInput) *AppConfigDeleteHostedConfigurationVersionFuture {
	future := workflow.ExecuteActivity(ctx, "aws-appconfig-DeleteHostedConfigurationVersion", input)
	return &AppConfigDeleteHostedConfigurationVersionFuture{Future: future}
}

func (a *stub) GetApplication(ctx workflow.Context, input *appconfig.GetApplicationInput) (*appconfig.GetApplicationOutput, error) {
	var output appconfig.GetApplicationOutput
	err := workflow.ExecuteActivity(ctx, "aws-appconfig-GetApplication", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetApplicationAsync(ctx workflow.Context, input *appconfig.GetApplicationInput) *AppConfigGetApplicationFuture {
	future := workflow.ExecuteActivity(ctx, "aws-appconfig-GetApplication", input)
	return &AppConfigGetApplicationFuture{Future: future}
}

func (a *stub) GetConfiguration(ctx workflow.Context, input *appconfig.GetConfigurationInput) (*appconfig.GetConfigurationOutput, error) {
	var output appconfig.GetConfigurationOutput
	err := workflow.ExecuteActivity(ctx, "aws-appconfig-GetConfiguration", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetConfigurationAsync(ctx workflow.Context, input *appconfig.GetConfigurationInput) *AppConfigGetConfigurationFuture {
	future := workflow.ExecuteActivity(ctx, "aws-appconfig-GetConfiguration", input)
	return &AppConfigGetConfigurationFuture{Future: future}
}

func (a *stub) GetConfigurationProfile(ctx workflow.Context, input *appconfig.GetConfigurationProfileInput) (*appconfig.GetConfigurationProfileOutput, error) {
	var output appconfig.GetConfigurationProfileOutput
	err := workflow.ExecuteActivity(ctx, "aws-appconfig-GetConfigurationProfile", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetConfigurationProfileAsync(ctx workflow.Context, input *appconfig.GetConfigurationProfileInput) *AppConfigGetConfigurationProfileFuture {
	future := workflow.ExecuteActivity(ctx, "aws-appconfig-GetConfigurationProfile", input)
	return &AppConfigGetConfigurationProfileFuture{Future: future}
}

func (a *stub) GetDeployment(ctx workflow.Context, input *appconfig.GetDeploymentInput) (*appconfig.GetDeploymentOutput, error) {
	var output appconfig.GetDeploymentOutput
	err := workflow.ExecuteActivity(ctx, "aws-appconfig-GetDeployment", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetDeploymentAsync(ctx workflow.Context, input *appconfig.GetDeploymentInput) *AppConfigGetDeploymentFuture {
	future := workflow.ExecuteActivity(ctx, "aws-appconfig-GetDeployment", input)
	return &AppConfigGetDeploymentFuture{Future: future}
}

func (a *stub) GetDeploymentStrategy(ctx workflow.Context, input *appconfig.GetDeploymentStrategyInput) (*appconfig.GetDeploymentStrategyOutput, error) {
	var output appconfig.GetDeploymentStrategyOutput
	err := workflow.ExecuteActivity(ctx, "aws-appconfig-GetDeploymentStrategy", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetDeploymentStrategyAsync(ctx workflow.Context, input *appconfig.GetDeploymentStrategyInput) *AppConfigGetDeploymentStrategyFuture {
	future := workflow.ExecuteActivity(ctx, "aws-appconfig-GetDeploymentStrategy", input)
	return &AppConfigGetDeploymentStrategyFuture{Future: future}
}

func (a *stub) GetEnvironment(ctx workflow.Context, input *appconfig.GetEnvironmentInput) (*appconfig.GetEnvironmentOutput, error) {
	var output appconfig.GetEnvironmentOutput
	err := workflow.ExecuteActivity(ctx, "aws-appconfig-GetEnvironment", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetEnvironmentAsync(ctx workflow.Context, input *appconfig.GetEnvironmentInput) *AppConfigGetEnvironmentFuture {
	future := workflow.ExecuteActivity(ctx, "aws-appconfig-GetEnvironment", input)
	return &AppConfigGetEnvironmentFuture{Future: future}
}

func (a *stub) GetHostedConfigurationVersion(ctx workflow.Context, input *appconfig.GetHostedConfigurationVersionInput) (*appconfig.GetHostedConfigurationVersionOutput, error) {
	var output appconfig.GetHostedConfigurationVersionOutput
	err := workflow.ExecuteActivity(ctx, "aws-appconfig-GetHostedConfigurationVersion", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetHostedConfigurationVersionAsync(ctx workflow.Context, input *appconfig.GetHostedConfigurationVersionInput) *AppConfigGetHostedConfigurationVersionFuture {
	future := workflow.ExecuteActivity(ctx, "aws-appconfig-GetHostedConfigurationVersion", input)
	return &AppConfigGetHostedConfigurationVersionFuture{Future: future}
}

func (a *stub) ListApplications(ctx workflow.Context, input *appconfig.ListApplicationsInput) (*appconfig.ListApplicationsOutput, error) {
	var output appconfig.ListApplicationsOutput
	err := workflow.ExecuteActivity(ctx, "aws-appconfig-ListApplications", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListApplicationsAsync(ctx workflow.Context, input *appconfig.ListApplicationsInput) *AppConfigListApplicationsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-appconfig-ListApplications", input)
	return &AppConfigListApplicationsFuture{Future: future}
}

func (a *stub) ListConfigurationProfiles(ctx workflow.Context, input *appconfig.ListConfigurationProfilesInput) (*appconfig.ListConfigurationProfilesOutput, error) {
	var output appconfig.ListConfigurationProfilesOutput
	err := workflow.ExecuteActivity(ctx, "aws-appconfig-ListConfigurationProfiles", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListConfigurationProfilesAsync(ctx workflow.Context, input *appconfig.ListConfigurationProfilesInput) *AppConfigListConfigurationProfilesFuture {
	future := workflow.ExecuteActivity(ctx, "aws-appconfig-ListConfigurationProfiles", input)
	return &AppConfigListConfigurationProfilesFuture{Future: future}
}

func (a *stub) ListDeploymentStrategies(ctx workflow.Context, input *appconfig.ListDeploymentStrategiesInput) (*appconfig.ListDeploymentStrategiesOutput, error) {
	var output appconfig.ListDeploymentStrategiesOutput
	err := workflow.ExecuteActivity(ctx, "aws-appconfig-ListDeploymentStrategies", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListDeploymentStrategiesAsync(ctx workflow.Context, input *appconfig.ListDeploymentStrategiesInput) *AppConfigListDeploymentStrategiesFuture {
	future := workflow.ExecuteActivity(ctx, "aws-appconfig-ListDeploymentStrategies", input)
	return &AppConfigListDeploymentStrategiesFuture{Future: future}
}

func (a *stub) ListDeployments(ctx workflow.Context, input *appconfig.ListDeploymentsInput) (*appconfig.ListDeploymentsOutput, error) {
	var output appconfig.ListDeploymentsOutput
	err := workflow.ExecuteActivity(ctx, "aws-appconfig-ListDeployments", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListDeploymentsAsync(ctx workflow.Context, input *appconfig.ListDeploymentsInput) *AppConfigListDeploymentsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-appconfig-ListDeployments", input)
	return &AppConfigListDeploymentsFuture{Future: future}
}

func (a *stub) ListEnvironments(ctx workflow.Context, input *appconfig.ListEnvironmentsInput) (*appconfig.ListEnvironmentsOutput, error) {
	var output appconfig.ListEnvironmentsOutput
	err := workflow.ExecuteActivity(ctx, "aws-appconfig-ListEnvironments", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListEnvironmentsAsync(ctx workflow.Context, input *appconfig.ListEnvironmentsInput) *AppConfigListEnvironmentsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-appconfig-ListEnvironments", input)
	return &AppConfigListEnvironmentsFuture{Future: future}
}

func (a *stub) ListHostedConfigurationVersions(ctx workflow.Context, input *appconfig.ListHostedConfigurationVersionsInput) (*appconfig.ListHostedConfigurationVersionsOutput, error) {
	var output appconfig.ListHostedConfigurationVersionsOutput
	err := workflow.ExecuteActivity(ctx, "aws-appconfig-ListHostedConfigurationVersions", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListHostedConfigurationVersionsAsync(ctx workflow.Context, input *appconfig.ListHostedConfigurationVersionsInput) *AppConfigListHostedConfigurationVersionsFuture {
	future := workflow.ExecuteActivity(ctx, "aws-appconfig-ListHostedConfigurationVersions", input)
	return &AppConfigListHostedConfigurationVersionsFuture{Future: future}
}

func (a *stub) ListTagsForResource(ctx workflow.Context, input *appconfig.ListTagsForResourceInput) (*appconfig.ListTagsForResourceOutput, error) {
	var output appconfig.ListTagsForResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws-appconfig-ListTagsForResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListTagsForResourceAsync(ctx workflow.Context, input *appconfig.ListTagsForResourceInput) *AppConfigListTagsForResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws-appconfig-ListTagsForResource", input)
	return &AppConfigListTagsForResourceFuture{Future: future}
}

func (a *stub) StartDeployment(ctx workflow.Context, input *appconfig.StartDeploymentInput) (*appconfig.StartDeploymentOutput, error) {
	var output appconfig.StartDeploymentOutput
	err := workflow.ExecuteActivity(ctx, "aws-appconfig-StartDeployment", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) StartDeploymentAsync(ctx workflow.Context, input *appconfig.StartDeploymentInput) *AppConfigStartDeploymentFuture {
	future := workflow.ExecuteActivity(ctx, "aws-appconfig-StartDeployment", input)
	return &AppConfigStartDeploymentFuture{Future: future}
}

func (a *stub) StopDeployment(ctx workflow.Context, input *appconfig.StopDeploymentInput) (*appconfig.StopDeploymentOutput, error) {
	var output appconfig.StopDeploymentOutput
	err := workflow.ExecuteActivity(ctx, "aws-appconfig-StopDeployment", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) StopDeploymentAsync(ctx workflow.Context, input *appconfig.StopDeploymentInput) *AppConfigStopDeploymentFuture {
	future := workflow.ExecuteActivity(ctx, "aws-appconfig-StopDeployment", input)
	return &AppConfigStopDeploymentFuture{Future: future}
}

func (a *stub) TagResource(ctx workflow.Context, input *appconfig.TagResourceInput) (*appconfig.TagResourceOutput, error) {
	var output appconfig.TagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws-appconfig-TagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) TagResourceAsync(ctx workflow.Context, input *appconfig.TagResourceInput) *AppConfigTagResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws-appconfig-TagResource", input)
	return &AppConfigTagResourceFuture{Future: future}
}

func (a *stub) UntagResource(ctx workflow.Context, input *appconfig.UntagResourceInput) (*appconfig.UntagResourceOutput, error) {
	var output appconfig.UntagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws-appconfig-UntagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UntagResourceAsync(ctx workflow.Context, input *appconfig.UntagResourceInput) *AppConfigUntagResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws-appconfig-UntagResource", input)
	return &AppConfigUntagResourceFuture{Future: future}
}

func (a *stub) UpdateApplication(ctx workflow.Context, input *appconfig.UpdateApplicationInput) (*appconfig.UpdateApplicationOutput, error) {
	var output appconfig.UpdateApplicationOutput
	err := workflow.ExecuteActivity(ctx, "aws-appconfig-UpdateApplication", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateApplicationAsync(ctx workflow.Context, input *appconfig.UpdateApplicationInput) *AppConfigUpdateApplicationFuture {
	future := workflow.ExecuteActivity(ctx, "aws-appconfig-UpdateApplication", input)
	return &AppConfigUpdateApplicationFuture{Future: future}
}

func (a *stub) UpdateConfigurationProfile(ctx workflow.Context, input *appconfig.UpdateConfigurationProfileInput) (*appconfig.UpdateConfigurationProfileOutput, error) {
	var output appconfig.UpdateConfigurationProfileOutput
	err := workflow.ExecuteActivity(ctx, "aws-appconfig-UpdateConfigurationProfile", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateConfigurationProfileAsync(ctx workflow.Context, input *appconfig.UpdateConfigurationProfileInput) *AppConfigUpdateConfigurationProfileFuture {
	future := workflow.ExecuteActivity(ctx, "aws-appconfig-UpdateConfigurationProfile", input)
	return &AppConfigUpdateConfigurationProfileFuture{Future: future}
}

func (a *stub) UpdateDeploymentStrategy(ctx workflow.Context, input *appconfig.UpdateDeploymentStrategyInput) (*appconfig.UpdateDeploymentStrategyOutput, error) {
	var output appconfig.UpdateDeploymentStrategyOutput
	err := workflow.ExecuteActivity(ctx, "aws-appconfig-UpdateDeploymentStrategy", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateDeploymentStrategyAsync(ctx workflow.Context, input *appconfig.UpdateDeploymentStrategyInput) *AppConfigUpdateDeploymentStrategyFuture {
	future := workflow.ExecuteActivity(ctx, "aws-appconfig-UpdateDeploymentStrategy", input)
	return &AppConfigUpdateDeploymentStrategyFuture{Future: future}
}

func (a *stub) UpdateEnvironment(ctx workflow.Context, input *appconfig.UpdateEnvironmentInput) (*appconfig.UpdateEnvironmentOutput, error) {
	var output appconfig.UpdateEnvironmentOutput
	err := workflow.ExecuteActivity(ctx, "aws-appconfig-UpdateEnvironment", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateEnvironmentAsync(ctx workflow.Context, input *appconfig.UpdateEnvironmentInput) *AppConfigUpdateEnvironmentFuture {
	future := workflow.ExecuteActivity(ctx, "aws-appconfig-UpdateEnvironment", input)
	return &AppConfigUpdateEnvironmentFuture{Future: future}
}

func (a *stub) ValidateConfiguration(ctx workflow.Context, input *appconfig.ValidateConfigurationInput) (*appconfig.ValidateConfigurationOutput, error) {
	var output appconfig.ValidateConfigurationOutput
	err := workflow.ExecuteActivity(ctx, "aws-appconfig-ValidateConfiguration", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ValidateConfigurationAsync(ctx workflow.Context, input *appconfig.ValidateConfigurationInput) *AppConfigValidateConfigurationFuture {
	future := workflow.ExecuteActivity(ctx, "aws-appconfig-ValidateConfiguration", input)
	return &AppConfigValidateConfigurationFuture{Future: future}
}