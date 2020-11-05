// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7

package kafka

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kafka"
	"github.com/aws/aws-sdk-go/service/kafka/kafkaiface"

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
	client kafkaiface.KafkaAPI

	sessionFactory SessionFactory
}

func NewActivities(sess *session.Session, config ...*aws.Config) *Activities {
	client := kafka.New(sess, config...)
	return &Activities{client: client}
}

func NewActivitiesWithSessionFactory(sessionFactory SessionFactory) *Activities {
	return &Activities{sessionFactory: sessionFactory}
}

func (a *Activities) getClient(ctx context.Context) (kafkaiface.KafkaAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}

	return kafka.New(sess), nil
}

func (a *Activities) BatchAssociateScramSecret(ctx context.Context, input *kafka.BatchAssociateScramSecretInput) (*kafka.BatchAssociateScramSecretOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.BatchAssociateScramSecretWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) BatchDisassociateScramSecret(ctx context.Context, input *kafka.BatchDisassociateScramSecretInput) (*kafka.BatchDisassociateScramSecretOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.BatchDisassociateScramSecretWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateCluster(ctx context.Context, input *kafka.CreateClusterInput) (*kafka.CreateClusterOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateClusterWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateConfiguration(ctx context.Context, input *kafka.CreateConfigurationInput) (*kafka.CreateConfigurationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateConfigurationWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteCluster(ctx context.Context, input *kafka.DeleteClusterInput) (*kafka.DeleteClusterOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteClusterWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteConfiguration(ctx context.Context, input *kafka.DeleteConfigurationInput) (*kafka.DeleteConfigurationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteConfigurationWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeCluster(ctx context.Context, input *kafka.DescribeClusterInput) (*kafka.DescribeClusterOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeClusterWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeClusterOperation(ctx context.Context, input *kafka.DescribeClusterOperationInput) (*kafka.DescribeClusterOperationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeClusterOperationWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeConfiguration(ctx context.Context, input *kafka.DescribeConfigurationInput) (*kafka.DescribeConfigurationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeConfigurationWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeConfigurationRevision(ctx context.Context, input *kafka.DescribeConfigurationRevisionInput) (*kafka.DescribeConfigurationRevisionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeConfigurationRevisionWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetBootstrapBrokers(ctx context.Context, input *kafka.GetBootstrapBrokersInput) (*kafka.GetBootstrapBrokersOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetBootstrapBrokersWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetCompatibleKafkaVersions(ctx context.Context, input *kafka.GetCompatibleKafkaVersionsInput) (*kafka.GetCompatibleKafkaVersionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetCompatibleKafkaVersionsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListClusterOperations(ctx context.Context, input *kafka.ListClusterOperationsInput) (*kafka.ListClusterOperationsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListClusterOperationsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListClusters(ctx context.Context, input *kafka.ListClustersInput) (*kafka.ListClustersOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListClustersWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListConfigurationRevisions(ctx context.Context, input *kafka.ListConfigurationRevisionsInput) (*kafka.ListConfigurationRevisionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListConfigurationRevisionsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListConfigurations(ctx context.Context, input *kafka.ListConfigurationsInput) (*kafka.ListConfigurationsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListConfigurationsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListKafkaVersions(ctx context.Context, input *kafka.ListKafkaVersionsInput) (*kafka.ListKafkaVersionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListKafkaVersionsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListNodes(ctx context.Context, input *kafka.ListNodesInput) (*kafka.ListNodesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListNodesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListScramSecrets(ctx context.Context, input *kafka.ListScramSecretsInput) (*kafka.ListScramSecretsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListScramSecretsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListTagsForResource(ctx context.Context, input *kafka.ListTagsForResourceInput) (*kafka.ListTagsForResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListTagsForResourceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) RebootBroker(ctx context.Context, input *kafka.RebootBrokerInput) (*kafka.RebootBrokerOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.RebootBrokerWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) TagResource(ctx context.Context, input *kafka.TagResourceInput) (*kafka.TagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.TagResourceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UntagResource(ctx context.Context, input *kafka.UntagResourceInput) (*kafka.UntagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UntagResourceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateBrokerCount(ctx context.Context, input *kafka.UpdateBrokerCountInput) (*kafka.UpdateBrokerCountOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateBrokerCountWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateBrokerStorage(ctx context.Context, input *kafka.UpdateBrokerStorageInput) (*kafka.UpdateBrokerStorageOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateBrokerStorageWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateClusterConfiguration(ctx context.Context, input *kafka.UpdateClusterConfigurationInput) (*kafka.UpdateClusterConfigurationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateClusterConfigurationWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateClusterKafkaVersion(ctx context.Context, input *kafka.UpdateClusterKafkaVersionInput) (*kafka.UpdateClusterKafkaVersionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateClusterKafkaVersionWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateConfiguration(ctx context.Context, input *kafka.UpdateConfigurationInput) (*kafka.UpdateConfigurationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateConfigurationWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateMonitoring(ctx context.Context, input *kafka.UpdateMonitoringInput) (*kafka.UpdateMonitoringOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateMonitoringWithContext(ctx, input)

	return output, internal.EncodeError(err)
}
