// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7

package gamelift

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/gamelift"
	"github.com/aws/aws-sdk-go/service/gamelift/gameliftiface"

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
	client gameliftiface.GameLiftAPI

	sessionFactory SessionFactory
}

func NewActivities(sess *session.Session, config ...*aws.Config) *Activities {
	client := gamelift.New(sess, config...)
	return &Activities{client: client}
}

func NewActivitiesWithSessionFactory(sessionFactory SessionFactory) *Activities {
	return &Activities{sessionFactory: sessionFactory}
}

func (a *Activities) getClient(ctx context.Context) (gameliftiface.GameLiftAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}

	return gamelift.New(sess), nil
}

func (a *Activities) AcceptMatch(ctx context.Context, input *gamelift.AcceptMatchInput) (*gamelift.AcceptMatchOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.AcceptMatchWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ClaimGameServer(ctx context.Context, input *gamelift.ClaimGameServerInput) (*gamelift.ClaimGameServerOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ClaimGameServerWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateAlias(ctx context.Context, input *gamelift.CreateAliasInput) (*gamelift.CreateAliasOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateAliasWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateBuild(ctx context.Context, input *gamelift.CreateBuildInput) (*gamelift.CreateBuildOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateBuildWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateFleet(ctx context.Context, input *gamelift.CreateFleetInput) (*gamelift.CreateFleetOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateFleetWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateGameServerGroup(ctx context.Context, input *gamelift.CreateGameServerGroupInput) (*gamelift.CreateGameServerGroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateGameServerGroupWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateGameSession(ctx context.Context, input *gamelift.CreateGameSessionInput) (*gamelift.CreateGameSessionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateGameSessionWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateGameSessionQueue(ctx context.Context, input *gamelift.CreateGameSessionQueueInput) (*gamelift.CreateGameSessionQueueOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateGameSessionQueueWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateMatchmakingConfiguration(ctx context.Context, input *gamelift.CreateMatchmakingConfigurationInput) (*gamelift.CreateMatchmakingConfigurationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateMatchmakingConfigurationWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateMatchmakingRuleSet(ctx context.Context, input *gamelift.CreateMatchmakingRuleSetInput) (*gamelift.CreateMatchmakingRuleSetOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateMatchmakingRuleSetWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreatePlayerSession(ctx context.Context, input *gamelift.CreatePlayerSessionInput) (*gamelift.CreatePlayerSessionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreatePlayerSessionWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreatePlayerSessions(ctx context.Context, input *gamelift.CreatePlayerSessionsInput) (*gamelift.CreatePlayerSessionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreatePlayerSessionsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateScript(ctx context.Context, input *gamelift.CreateScriptInput) (*gamelift.CreateScriptOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateScriptWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateVpcPeeringAuthorization(ctx context.Context, input *gamelift.CreateVpcPeeringAuthorizationInput) (*gamelift.CreateVpcPeeringAuthorizationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateVpcPeeringAuthorizationWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateVpcPeeringConnection(ctx context.Context, input *gamelift.CreateVpcPeeringConnectionInput) (*gamelift.CreateVpcPeeringConnectionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateVpcPeeringConnectionWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteAlias(ctx context.Context, input *gamelift.DeleteAliasInput) (*gamelift.DeleteAliasOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteAliasWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteBuild(ctx context.Context, input *gamelift.DeleteBuildInput) (*gamelift.DeleteBuildOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteBuildWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteFleet(ctx context.Context, input *gamelift.DeleteFleetInput) (*gamelift.DeleteFleetOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteFleetWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteGameServerGroup(ctx context.Context, input *gamelift.DeleteGameServerGroupInput) (*gamelift.DeleteGameServerGroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteGameServerGroupWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteGameSessionQueue(ctx context.Context, input *gamelift.DeleteGameSessionQueueInput) (*gamelift.DeleteGameSessionQueueOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteGameSessionQueueWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteMatchmakingConfiguration(ctx context.Context, input *gamelift.DeleteMatchmakingConfigurationInput) (*gamelift.DeleteMatchmakingConfigurationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteMatchmakingConfigurationWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteMatchmakingRuleSet(ctx context.Context, input *gamelift.DeleteMatchmakingRuleSetInput) (*gamelift.DeleteMatchmakingRuleSetOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteMatchmakingRuleSetWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteScalingPolicy(ctx context.Context, input *gamelift.DeleteScalingPolicyInput) (*gamelift.DeleteScalingPolicyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteScalingPolicyWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteScript(ctx context.Context, input *gamelift.DeleteScriptInput) (*gamelift.DeleteScriptOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteScriptWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteVpcPeeringAuthorization(ctx context.Context, input *gamelift.DeleteVpcPeeringAuthorizationInput) (*gamelift.DeleteVpcPeeringAuthorizationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteVpcPeeringAuthorizationWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteVpcPeeringConnection(ctx context.Context, input *gamelift.DeleteVpcPeeringConnectionInput) (*gamelift.DeleteVpcPeeringConnectionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteVpcPeeringConnectionWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeregisterGameServer(ctx context.Context, input *gamelift.DeregisterGameServerInput) (*gamelift.DeregisterGameServerOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeregisterGameServerWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeAlias(ctx context.Context, input *gamelift.DescribeAliasInput) (*gamelift.DescribeAliasOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeAliasWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeBuild(ctx context.Context, input *gamelift.DescribeBuildInput) (*gamelift.DescribeBuildOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeBuildWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeEC2InstanceLimits(ctx context.Context, input *gamelift.DescribeEC2InstanceLimitsInput) (*gamelift.DescribeEC2InstanceLimitsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeEC2InstanceLimitsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeFleetAttributes(ctx context.Context, input *gamelift.DescribeFleetAttributesInput) (*gamelift.DescribeFleetAttributesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeFleetAttributesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeFleetCapacity(ctx context.Context, input *gamelift.DescribeFleetCapacityInput) (*gamelift.DescribeFleetCapacityOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeFleetCapacityWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeFleetEvents(ctx context.Context, input *gamelift.DescribeFleetEventsInput) (*gamelift.DescribeFleetEventsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeFleetEventsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeFleetPortSettings(ctx context.Context, input *gamelift.DescribeFleetPortSettingsInput) (*gamelift.DescribeFleetPortSettingsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeFleetPortSettingsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeFleetUtilization(ctx context.Context, input *gamelift.DescribeFleetUtilizationInput) (*gamelift.DescribeFleetUtilizationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeFleetUtilizationWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeGameServer(ctx context.Context, input *gamelift.DescribeGameServerInput) (*gamelift.DescribeGameServerOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeGameServerWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeGameServerGroup(ctx context.Context, input *gamelift.DescribeGameServerGroupInput) (*gamelift.DescribeGameServerGroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeGameServerGroupWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeGameServerInstances(ctx context.Context, input *gamelift.DescribeGameServerInstancesInput) (*gamelift.DescribeGameServerInstancesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeGameServerInstancesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeGameSessionDetails(ctx context.Context, input *gamelift.DescribeGameSessionDetailsInput) (*gamelift.DescribeGameSessionDetailsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeGameSessionDetailsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeGameSessionPlacement(ctx context.Context, input *gamelift.DescribeGameSessionPlacementInput) (*gamelift.DescribeGameSessionPlacementOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeGameSessionPlacementWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeGameSessionQueues(ctx context.Context, input *gamelift.DescribeGameSessionQueuesInput) (*gamelift.DescribeGameSessionQueuesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeGameSessionQueuesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeGameSessions(ctx context.Context, input *gamelift.DescribeGameSessionsInput) (*gamelift.DescribeGameSessionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeGameSessionsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeInstances(ctx context.Context, input *gamelift.DescribeInstancesInput) (*gamelift.DescribeInstancesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeInstancesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeMatchmaking(ctx context.Context, input *gamelift.DescribeMatchmakingInput) (*gamelift.DescribeMatchmakingOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeMatchmakingWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeMatchmakingConfigurations(ctx context.Context, input *gamelift.DescribeMatchmakingConfigurationsInput) (*gamelift.DescribeMatchmakingConfigurationsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeMatchmakingConfigurationsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeMatchmakingRuleSets(ctx context.Context, input *gamelift.DescribeMatchmakingRuleSetsInput) (*gamelift.DescribeMatchmakingRuleSetsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeMatchmakingRuleSetsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribePlayerSessions(ctx context.Context, input *gamelift.DescribePlayerSessionsInput) (*gamelift.DescribePlayerSessionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribePlayerSessionsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeRuntimeConfiguration(ctx context.Context, input *gamelift.DescribeRuntimeConfigurationInput) (*gamelift.DescribeRuntimeConfigurationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeRuntimeConfigurationWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeScalingPolicies(ctx context.Context, input *gamelift.DescribeScalingPoliciesInput) (*gamelift.DescribeScalingPoliciesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeScalingPoliciesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeScript(ctx context.Context, input *gamelift.DescribeScriptInput) (*gamelift.DescribeScriptOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeScriptWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeVpcPeeringAuthorizations(ctx context.Context, input *gamelift.DescribeVpcPeeringAuthorizationsInput) (*gamelift.DescribeVpcPeeringAuthorizationsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeVpcPeeringAuthorizationsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeVpcPeeringConnections(ctx context.Context, input *gamelift.DescribeVpcPeeringConnectionsInput) (*gamelift.DescribeVpcPeeringConnectionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeVpcPeeringConnectionsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetGameSessionLogUrl(ctx context.Context, input *gamelift.GetGameSessionLogUrlInput) (*gamelift.GetGameSessionLogUrlOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetGameSessionLogUrlWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetInstanceAccess(ctx context.Context, input *gamelift.GetInstanceAccessInput) (*gamelift.GetInstanceAccessOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetInstanceAccessWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListAliases(ctx context.Context, input *gamelift.ListAliasesInput) (*gamelift.ListAliasesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListAliasesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListBuilds(ctx context.Context, input *gamelift.ListBuildsInput) (*gamelift.ListBuildsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListBuildsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListFleets(ctx context.Context, input *gamelift.ListFleetsInput) (*gamelift.ListFleetsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListFleetsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListGameServerGroups(ctx context.Context, input *gamelift.ListGameServerGroupsInput) (*gamelift.ListGameServerGroupsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListGameServerGroupsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListGameServers(ctx context.Context, input *gamelift.ListGameServersInput) (*gamelift.ListGameServersOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListGameServersWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListScripts(ctx context.Context, input *gamelift.ListScriptsInput) (*gamelift.ListScriptsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListScriptsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListTagsForResource(ctx context.Context, input *gamelift.ListTagsForResourceInput) (*gamelift.ListTagsForResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListTagsForResourceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) PutScalingPolicy(ctx context.Context, input *gamelift.PutScalingPolicyInput) (*gamelift.PutScalingPolicyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.PutScalingPolicyWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) RegisterGameServer(ctx context.Context, input *gamelift.RegisterGameServerInput) (*gamelift.RegisterGameServerOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.RegisterGameServerWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) RequestUploadCredentials(ctx context.Context, input *gamelift.RequestUploadCredentialsInput) (*gamelift.RequestUploadCredentialsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.RequestUploadCredentialsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ResolveAlias(ctx context.Context, input *gamelift.ResolveAliasInput) (*gamelift.ResolveAliasOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ResolveAliasWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ResumeGameServerGroup(ctx context.Context, input *gamelift.ResumeGameServerGroupInput) (*gamelift.ResumeGameServerGroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ResumeGameServerGroupWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) SearchGameSessions(ctx context.Context, input *gamelift.SearchGameSessionsInput) (*gamelift.SearchGameSessionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.SearchGameSessionsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) StartFleetActions(ctx context.Context, input *gamelift.StartFleetActionsInput) (*gamelift.StartFleetActionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.StartFleetActionsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) StartGameSessionPlacement(ctx context.Context, input *gamelift.StartGameSessionPlacementInput) (*gamelift.StartGameSessionPlacementOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.StartGameSessionPlacementWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) StartMatchBackfill(ctx context.Context, input *gamelift.StartMatchBackfillInput) (*gamelift.StartMatchBackfillOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.StartMatchBackfillWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) StartMatchmaking(ctx context.Context, input *gamelift.StartMatchmakingInput) (*gamelift.StartMatchmakingOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.StartMatchmakingWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) StopFleetActions(ctx context.Context, input *gamelift.StopFleetActionsInput) (*gamelift.StopFleetActionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.StopFleetActionsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) StopGameSessionPlacement(ctx context.Context, input *gamelift.StopGameSessionPlacementInput) (*gamelift.StopGameSessionPlacementOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.StopGameSessionPlacementWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) StopMatchmaking(ctx context.Context, input *gamelift.StopMatchmakingInput) (*gamelift.StopMatchmakingOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.StopMatchmakingWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) SuspendGameServerGroup(ctx context.Context, input *gamelift.SuspendGameServerGroupInput) (*gamelift.SuspendGameServerGroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.SuspendGameServerGroupWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) TagResource(ctx context.Context, input *gamelift.TagResourceInput) (*gamelift.TagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.TagResourceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UntagResource(ctx context.Context, input *gamelift.UntagResourceInput) (*gamelift.UntagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UntagResourceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateAlias(ctx context.Context, input *gamelift.UpdateAliasInput) (*gamelift.UpdateAliasOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateAliasWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateBuild(ctx context.Context, input *gamelift.UpdateBuildInput) (*gamelift.UpdateBuildOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateBuildWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateFleetAttributes(ctx context.Context, input *gamelift.UpdateFleetAttributesInput) (*gamelift.UpdateFleetAttributesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateFleetAttributesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateFleetCapacity(ctx context.Context, input *gamelift.UpdateFleetCapacityInput) (*gamelift.UpdateFleetCapacityOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateFleetCapacityWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateFleetPortSettings(ctx context.Context, input *gamelift.UpdateFleetPortSettingsInput) (*gamelift.UpdateFleetPortSettingsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateFleetPortSettingsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateGameServer(ctx context.Context, input *gamelift.UpdateGameServerInput) (*gamelift.UpdateGameServerOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateGameServerWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateGameServerGroup(ctx context.Context, input *gamelift.UpdateGameServerGroupInput) (*gamelift.UpdateGameServerGroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateGameServerGroupWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateGameSession(ctx context.Context, input *gamelift.UpdateGameSessionInput) (*gamelift.UpdateGameSessionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateGameSessionWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateGameSessionQueue(ctx context.Context, input *gamelift.UpdateGameSessionQueueInput) (*gamelift.UpdateGameSessionQueueOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateGameSessionQueueWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateMatchmakingConfiguration(ctx context.Context, input *gamelift.UpdateMatchmakingConfigurationInput) (*gamelift.UpdateMatchmakingConfigurationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateMatchmakingConfigurationWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateRuntimeConfiguration(ctx context.Context, input *gamelift.UpdateRuntimeConfigurationInput) (*gamelift.UpdateRuntimeConfigurationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateRuntimeConfigurationWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateScript(ctx context.Context, input *gamelift.UpdateScriptInput) (*gamelift.UpdateScriptOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateScriptWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ValidateMatchmakingRuleSet(ctx context.Context, input *gamelift.ValidateMatchmakingRuleSetInput) (*gamelift.ValidateMatchmakingRuleSetOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ValidateMatchmakingRuleSetWithContext(ctx, input)

	return output, internal.EncodeError(err)
}
