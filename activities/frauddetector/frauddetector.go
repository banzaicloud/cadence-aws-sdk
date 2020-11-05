// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7

package frauddetector

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/frauddetector"
	"github.com/aws/aws-sdk-go/service/frauddetector/frauddetectoriface"

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
	client frauddetectoriface.FraudDetectorAPI

	sessionFactory SessionFactory
}

func NewActivities(sess *session.Session, config ...*aws.Config) *Activities {
	client := frauddetector.New(sess, config...)
	return &Activities{client: client}
}

func NewActivitiesWithSessionFactory(sessionFactory SessionFactory) *Activities {
	return &Activities{sessionFactory: sessionFactory}
}

func (a *Activities) getClient(ctx context.Context) (frauddetectoriface.FraudDetectorAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}

	return frauddetector.New(sess), nil
}

func (a *Activities) BatchCreateVariable(ctx context.Context, input *frauddetector.BatchCreateVariableInput) (*frauddetector.BatchCreateVariableOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.BatchCreateVariableWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) BatchGetVariable(ctx context.Context, input *frauddetector.BatchGetVariableInput) (*frauddetector.BatchGetVariableOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.BatchGetVariableWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateDetectorVersion(ctx context.Context, input *frauddetector.CreateDetectorVersionInput) (*frauddetector.CreateDetectorVersionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateDetectorVersionWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateModel(ctx context.Context, input *frauddetector.CreateModelInput) (*frauddetector.CreateModelOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateModelWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateModelVersion(ctx context.Context, input *frauddetector.CreateModelVersionInput) (*frauddetector.CreateModelVersionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateModelVersionWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateRule(ctx context.Context, input *frauddetector.CreateRuleInput) (*frauddetector.CreateRuleOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateRuleWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateVariable(ctx context.Context, input *frauddetector.CreateVariableInput) (*frauddetector.CreateVariableOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateVariableWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteDetector(ctx context.Context, input *frauddetector.DeleteDetectorInput) (*frauddetector.DeleteDetectorOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteDetectorWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteDetectorVersion(ctx context.Context, input *frauddetector.DeleteDetectorVersionInput) (*frauddetector.DeleteDetectorVersionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteDetectorVersionWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteEvent(ctx context.Context, input *frauddetector.DeleteEventInput) (*frauddetector.DeleteEventOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteEventWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteRule(ctx context.Context, input *frauddetector.DeleteRuleInput) (*frauddetector.DeleteRuleOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteRuleWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeDetector(ctx context.Context, input *frauddetector.DescribeDetectorInput) (*frauddetector.DescribeDetectorOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeDetectorWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeModelVersions(ctx context.Context, input *frauddetector.DescribeModelVersionsInput) (*frauddetector.DescribeModelVersionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeModelVersionsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetDetectorVersion(ctx context.Context, input *frauddetector.GetDetectorVersionInput) (*frauddetector.GetDetectorVersionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetDetectorVersionWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetDetectors(ctx context.Context, input *frauddetector.GetDetectorsInput) (*frauddetector.GetDetectorsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetDetectorsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetEntityTypes(ctx context.Context, input *frauddetector.GetEntityTypesInput) (*frauddetector.GetEntityTypesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetEntityTypesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetEventPrediction(ctx context.Context, input *frauddetector.GetEventPredictionInput) (*frauddetector.GetEventPredictionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetEventPredictionWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetEventTypes(ctx context.Context, input *frauddetector.GetEventTypesInput) (*frauddetector.GetEventTypesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetEventTypesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetExternalModels(ctx context.Context, input *frauddetector.GetExternalModelsInput) (*frauddetector.GetExternalModelsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetExternalModelsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetKMSEncryptionKey(ctx context.Context, input *frauddetector.GetKMSEncryptionKeyInput) (*frauddetector.GetKMSEncryptionKeyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetKMSEncryptionKeyWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetLabels(ctx context.Context, input *frauddetector.GetLabelsInput) (*frauddetector.GetLabelsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetLabelsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetModelVersion(ctx context.Context, input *frauddetector.GetModelVersionInput) (*frauddetector.GetModelVersionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetModelVersionWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetModels(ctx context.Context, input *frauddetector.GetModelsInput) (*frauddetector.GetModelsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetModelsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetOutcomes(ctx context.Context, input *frauddetector.GetOutcomesInput) (*frauddetector.GetOutcomesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetOutcomesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetRules(ctx context.Context, input *frauddetector.GetRulesInput) (*frauddetector.GetRulesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetRulesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetVariables(ctx context.Context, input *frauddetector.GetVariablesInput) (*frauddetector.GetVariablesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetVariablesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListTagsForResource(ctx context.Context, input *frauddetector.ListTagsForResourceInput) (*frauddetector.ListTagsForResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListTagsForResourceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) PutDetector(ctx context.Context, input *frauddetector.PutDetectorInput) (*frauddetector.PutDetectorOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.PutDetectorWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) PutEntityType(ctx context.Context, input *frauddetector.PutEntityTypeInput) (*frauddetector.PutEntityTypeOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.PutEntityTypeWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) PutEventType(ctx context.Context, input *frauddetector.PutEventTypeInput) (*frauddetector.PutEventTypeOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.PutEventTypeWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) PutExternalModel(ctx context.Context, input *frauddetector.PutExternalModelInput) (*frauddetector.PutExternalModelOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.PutExternalModelWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) PutKMSEncryptionKey(ctx context.Context, input *frauddetector.PutKMSEncryptionKeyInput) (*frauddetector.PutKMSEncryptionKeyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.PutKMSEncryptionKeyWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) PutLabel(ctx context.Context, input *frauddetector.PutLabelInput) (*frauddetector.PutLabelOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.PutLabelWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) PutOutcome(ctx context.Context, input *frauddetector.PutOutcomeInput) (*frauddetector.PutOutcomeOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.PutOutcomeWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) TagResource(ctx context.Context, input *frauddetector.TagResourceInput) (*frauddetector.TagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.TagResourceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UntagResource(ctx context.Context, input *frauddetector.UntagResourceInput) (*frauddetector.UntagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UntagResourceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateDetectorVersion(ctx context.Context, input *frauddetector.UpdateDetectorVersionInput) (*frauddetector.UpdateDetectorVersionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateDetectorVersionWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateDetectorVersionMetadata(ctx context.Context, input *frauddetector.UpdateDetectorVersionMetadataInput) (*frauddetector.UpdateDetectorVersionMetadataOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateDetectorVersionMetadataWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateDetectorVersionStatus(ctx context.Context, input *frauddetector.UpdateDetectorVersionStatusInput) (*frauddetector.UpdateDetectorVersionStatusOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateDetectorVersionStatusWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateModel(ctx context.Context, input *frauddetector.UpdateModelInput) (*frauddetector.UpdateModelOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateModelWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateModelVersion(ctx context.Context, input *frauddetector.UpdateModelVersionInput) (*frauddetector.UpdateModelVersionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateModelVersionWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateModelVersionStatus(ctx context.Context, input *frauddetector.UpdateModelVersionStatusInput) (*frauddetector.UpdateModelVersionStatusOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateModelVersionStatusWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateRuleMetadata(ctx context.Context, input *frauddetector.UpdateRuleMetadataInput) (*frauddetector.UpdateRuleMetadataOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateRuleMetadataWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateRuleVersion(ctx context.Context, input *frauddetector.UpdateRuleVersionInput) (*frauddetector.UpdateRuleVersionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateRuleVersionWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateVariable(ctx context.Context, input *frauddetector.UpdateVariableInput) (*frauddetector.UpdateVariableOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateVariableWithContext(ctx, input)

	return output, internal.EncodeError(err)
}
