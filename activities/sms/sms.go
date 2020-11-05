// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7

package sms

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sms"
	"github.com/aws/aws-sdk-go/service/sms/smsiface"

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
	client smsiface.SMSAPI

	sessionFactory SessionFactory
}

func NewActivities(sess *session.Session, config ...*aws.Config) *Activities {
	client := sms.New(sess, config...)
	return &Activities{client: client}
}

func NewActivitiesWithSessionFactory(sessionFactory SessionFactory) *Activities {
	return &Activities{sessionFactory: sessionFactory}
}

func (a *Activities) getClient(ctx context.Context) (smsiface.SMSAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}

	return sms.New(sess), nil
}

func (a *Activities) CreateApp(ctx context.Context, input *sms.CreateAppInput) (*sms.CreateAppOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	internal.SetClientToken(ctx, &input.ClientToken)
	output, err := client.CreateAppWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateReplicationJob(ctx context.Context, input *sms.CreateReplicationJobInput) (*sms.CreateReplicationJobOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateReplicationJobWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteApp(ctx context.Context, input *sms.DeleteAppInput) (*sms.DeleteAppOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteAppWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteAppLaunchConfiguration(ctx context.Context, input *sms.DeleteAppLaunchConfigurationInput) (*sms.DeleteAppLaunchConfigurationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteAppLaunchConfigurationWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteAppReplicationConfiguration(ctx context.Context, input *sms.DeleteAppReplicationConfigurationInput) (*sms.DeleteAppReplicationConfigurationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteAppReplicationConfigurationWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteAppValidationConfiguration(ctx context.Context, input *sms.DeleteAppValidationConfigurationInput) (*sms.DeleteAppValidationConfigurationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteAppValidationConfigurationWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteReplicationJob(ctx context.Context, input *sms.DeleteReplicationJobInput) (*sms.DeleteReplicationJobOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteReplicationJobWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteServerCatalog(ctx context.Context, input *sms.DeleteServerCatalogInput) (*sms.DeleteServerCatalogOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteServerCatalogWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DisassociateConnector(ctx context.Context, input *sms.DisassociateConnectorInput) (*sms.DisassociateConnectorOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DisassociateConnectorWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GenerateChangeSet(ctx context.Context, input *sms.GenerateChangeSetInput) (*sms.GenerateChangeSetOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GenerateChangeSetWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GenerateTemplate(ctx context.Context, input *sms.GenerateTemplateInput) (*sms.GenerateTemplateOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GenerateTemplateWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetApp(ctx context.Context, input *sms.GetAppInput) (*sms.GetAppOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetAppWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetAppLaunchConfiguration(ctx context.Context, input *sms.GetAppLaunchConfigurationInput) (*sms.GetAppLaunchConfigurationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetAppLaunchConfigurationWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetAppReplicationConfiguration(ctx context.Context, input *sms.GetAppReplicationConfigurationInput) (*sms.GetAppReplicationConfigurationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetAppReplicationConfigurationWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetAppValidationConfiguration(ctx context.Context, input *sms.GetAppValidationConfigurationInput) (*sms.GetAppValidationConfigurationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetAppValidationConfigurationWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetAppValidationOutput(ctx context.Context, input *sms.GetAppValidationOutputInput) (*sms.GetAppValidationOutputOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetAppValidationOutputWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetConnectors(ctx context.Context, input *sms.GetConnectorsInput) (*sms.GetConnectorsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetConnectorsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetReplicationJobs(ctx context.Context, input *sms.GetReplicationJobsInput) (*sms.GetReplicationJobsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetReplicationJobsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetReplicationRuns(ctx context.Context, input *sms.GetReplicationRunsInput) (*sms.GetReplicationRunsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetReplicationRunsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetServers(ctx context.Context, input *sms.GetServersInput) (*sms.GetServersOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetServersWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ImportAppCatalog(ctx context.Context, input *sms.ImportAppCatalogInput) (*sms.ImportAppCatalogOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ImportAppCatalogWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ImportServerCatalog(ctx context.Context, input *sms.ImportServerCatalogInput) (*sms.ImportServerCatalogOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ImportServerCatalogWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) LaunchApp(ctx context.Context, input *sms.LaunchAppInput) (*sms.LaunchAppOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.LaunchAppWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListApps(ctx context.Context, input *sms.ListAppsInput) (*sms.ListAppsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListAppsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) NotifyAppValidationOutput(ctx context.Context, input *sms.NotifyAppValidationOutputInput) (*sms.NotifyAppValidationOutputOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.NotifyAppValidationOutputWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) PutAppLaunchConfiguration(ctx context.Context, input *sms.PutAppLaunchConfigurationInput) (*sms.PutAppLaunchConfigurationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.PutAppLaunchConfigurationWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) PutAppReplicationConfiguration(ctx context.Context, input *sms.PutAppReplicationConfigurationInput) (*sms.PutAppReplicationConfigurationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.PutAppReplicationConfigurationWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) PutAppValidationConfiguration(ctx context.Context, input *sms.PutAppValidationConfigurationInput) (*sms.PutAppValidationConfigurationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.PutAppValidationConfigurationWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) StartAppReplication(ctx context.Context, input *sms.StartAppReplicationInput) (*sms.StartAppReplicationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.StartAppReplicationWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) StartOnDemandAppReplication(ctx context.Context, input *sms.StartOnDemandAppReplicationInput) (*sms.StartOnDemandAppReplicationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.StartOnDemandAppReplicationWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) StartOnDemandReplicationRun(ctx context.Context, input *sms.StartOnDemandReplicationRunInput) (*sms.StartOnDemandReplicationRunOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.StartOnDemandReplicationRunWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) StopAppReplication(ctx context.Context, input *sms.StopAppReplicationInput) (*sms.StopAppReplicationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.StopAppReplicationWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) TerminateApp(ctx context.Context, input *sms.TerminateAppInput) (*sms.TerminateAppOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.TerminateAppWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateApp(ctx context.Context, input *sms.UpdateAppInput) (*sms.UpdateAppOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateAppWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateReplicationJob(ctx context.Context, input *sms.UpdateReplicationJobInput) (*sms.UpdateReplicationJobOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateReplicationJobWithContext(ctx, input)

	return output, internal.EncodeError(err)
}
