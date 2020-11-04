// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.
// Copyright (c) 2020 Banzai Cloud All rights reserved.

package mturk

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/mturk"
	"github.com/aws/aws-sdk-go/service/mturk/mturkiface"

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
	client mturkiface.MTurkAPI

	sessionFactory SessionFactory
}

func NewActivities(sess *session.Session, config ...*aws.Config) *Activities {
	client := mturk.New(sess, config...)
	return &Activities{client: client}
}

func NewActivitiesWithSessionFactory(sessionFactory SessionFactory) *Activities {
	return &Activities{sessionFactory: sessionFactory}
}

func (a *Activities) getClient(ctx context.Context) (mturkiface.MTurkAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}

	return mturk.New(sess), nil
}

func (a *Activities) ApproveAssignment(ctx context.Context, input *mturk.ApproveAssignmentInput) (*mturk.ApproveAssignmentOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ApproveAssignmentWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) AssociateQualificationWithWorker(ctx context.Context, input *mturk.AssociateQualificationWithWorkerInput) (*mturk.AssociateQualificationWithWorkerOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.AssociateQualificationWithWorkerWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateAdditionalAssignmentsForHIT(ctx context.Context, input *mturk.CreateAdditionalAssignmentsForHITInput) (*mturk.CreateAdditionalAssignmentsForHITOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateAdditionalAssignmentsForHITWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateHIT(ctx context.Context, input *mturk.CreateHITInput) (*mturk.CreateHITOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateHITWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateHITType(ctx context.Context, input *mturk.CreateHITTypeInput) (*mturk.CreateHITTypeOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateHITTypeWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateHITWithHITType(ctx context.Context, input *mturk.CreateHITWithHITTypeInput) (*mturk.CreateHITWithHITTypeOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateHITWithHITTypeWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateQualificationType(ctx context.Context, input *mturk.CreateQualificationTypeInput) (*mturk.CreateQualificationTypeOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateQualificationTypeWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateWorkerBlock(ctx context.Context, input *mturk.CreateWorkerBlockInput) (*mturk.CreateWorkerBlockOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateWorkerBlockWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteHIT(ctx context.Context, input *mturk.DeleteHITInput) (*mturk.DeleteHITOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteHITWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteQualificationType(ctx context.Context, input *mturk.DeleteQualificationTypeInput) (*mturk.DeleteQualificationTypeOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteQualificationTypeWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteWorkerBlock(ctx context.Context, input *mturk.DeleteWorkerBlockInput) (*mturk.DeleteWorkerBlockOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteWorkerBlockWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DisassociateQualificationFromWorker(ctx context.Context, input *mturk.DisassociateQualificationFromWorkerInput) (*mturk.DisassociateQualificationFromWorkerOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DisassociateQualificationFromWorkerWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetAccountBalance(ctx context.Context, input *mturk.GetAccountBalanceInput) (*mturk.GetAccountBalanceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetAccountBalanceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetAssignment(ctx context.Context, input *mturk.GetAssignmentInput) (*mturk.GetAssignmentOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetAssignmentWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetFileUploadURL(ctx context.Context, input *mturk.GetFileUploadURLInput) (*mturk.GetFileUploadURLOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetFileUploadURLWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetHIT(ctx context.Context, input *mturk.GetHITInput) (*mturk.GetHITOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetHITWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetQualificationScore(ctx context.Context, input *mturk.GetQualificationScoreInput) (*mturk.GetQualificationScoreOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetQualificationScoreWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetQualificationType(ctx context.Context, input *mturk.GetQualificationTypeInput) (*mturk.GetQualificationTypeOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetQualificationTypeWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListAssignmentsForHIT(ctx context.Context, input *mturk.ListAssignmentsForHITInput) (*mturk.ListAssignmentsForHITOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListAssignmentsForHITWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListBonusPayments(ctx context.Context, input *mturk.ListBonusPaymentsInput) (*mturk.ListBonusPaymentsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListBonusPaymentsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListHITs(ctx context.Context, input *mturk.ListHITsInput) (*mturk.ListHITsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListHITsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListHITsForQualificationType(ctx context.Context, input *mturk.ListHITsForQualificationTypeInput) (*mturk.ListHITsForQualificationTypeOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListHITsForQualificationTypeWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListQualificationRequests(ctx context.Context, input *mturk.ListQualificationRequestsInput) (*mturk.ListQualificationRequestsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListQualificationRequestsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListQualificationTypes(ctx context.Context, input *mturk.ListQualificationTypesInput) (*mturk.ListQualificationTypesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListQualificationTypesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListReviewPolicyResultsForHIT(ctx context.Context, input *mturk.ListReviewPolicyResultsForHITInput) (*mturk.ListReviewPolicyResultsForHITOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListReviewPolicyResultsForHITWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListReviewableHITs(ctx context.Context, input *mturk.ListReviewableHITsInput) (*mturk.ListReviewableHITsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListReviewableHITsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListWorkerBlocks(ctx context.Context, input *mturk.ListWorkerBlocksInput) (*mturk.ListWorkerBlocksOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListWorkerBlocksWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListWorkersWithQualificationType(ctx context.Context, input *mturk.ListWorkersWithQualificationTypeInput) (*mturk.ListWorkersWithQualificationTypeOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListWorkersWithQualificationTypeWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) NotifyWorkers(ctx context.Context, input *mturk.NotifyWorkersInput) (*mturk.NotifyWorkersOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.NotifyWorkersWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) RejectAssignment(ctx context.Context, input *mturk.RejectAssignmentInput) (*mturk.RejectAssignmentOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.RejectAssignmentWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) SendBonus(ctx context.Context, input *mturk.SendBonusInput) (*mturk.SendBonusOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.SendBonusWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) SendTestEventNotification(ctx context.Context, input *mturk.SendTestEventNotificationInput) (*mturk.SendTestEventNotificationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.SendTestEventNotificationWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateExpirationForHIT(ctx context.Context, input *mturk.UpdateExpirationForHITInput) (*mturk.UpdateExpirationForHITOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateExpirationForHITWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateHITReviewStatus(ctx context.Context, input *mturk.UpdateHITReviewStatusInput) (*mturk.UpdateHITReviewStatusOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateHITReviewStatusWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateHITTypeOfHIT(ctx context.Context, input *mturk.UpdateHITTypeOfHITInput) (*mturk.UpdateHITTypeOfHITOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateHITTypeOfHITWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateNotificationSettings(ctx context.Context, input *mturk.UpdateNotificationSettingsInput) (*mturk.UpdateNotificationSettingsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateNotificationSettingsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateQualificationType(ctx context.Context, input *mturk.UpdateQualificationTypeInput) (*mturk.UpdateQualificationTypeOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateQualificationTypeWithContext(ctx, input)

	return output, internal.EncodeError(err)
}
