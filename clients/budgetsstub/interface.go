// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.
// Copyright (c) 2020 Banzai Cloud All rights reserved.

package budgetsstub

import (
	"github.com/aws/aws-sdk-go/service/budgets"
	"go.uber.org/cadence/workflow"

	"github.com/banzaicloud/cadence-aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	CreateBudget(ctx workflow.Context, input *budgets.CreateBudgetInput) (*budgets.CreateBudgetOutput, error)
	CreateBudgetAsync(ctx workflow.Context, input *budgets.CreateBudgetInput) *BudgetsCreateBudgetFuture

	CreateBudgetAction(ctx workflow.Context, input *budgets.CreateBudgetActionInput) (*budgets.CreateBudgetActionOutput, error)
	CreateBudgetActionAsync(ctx workflow.Context, input *budgets.CreateBudgetActionInput) *BudgetsCreateBudgetActionFuture

	CreateNotification(ctx workflow.Context, input *budgets.CreateNotificationInput) (*budgets.CreateNotificationOutput, error)
	CreateNotificationAsync(ctx workflow.Context, input *budgets.CreateNotificationInput) *BudgetsCreateNotificationFuture

	CreateSubscriber(ctx workflow.Context, input *budgets.CreateSubscriberInput) (*budgets.CreateSubscriberOutput, error)
	CreateSubscriberAsync(ctx workflow.Context, input *budgets.CreateSubscriberInput) *BudgetsCreateSubscriberFuture

	DeleteBudget(ctx workflow.Context, input *budgets.DeleteBudgetInput) (*budgets.DeleteBudgetOutput, error)
	DeleteBudgetAsync(ctx workflow.Context, input *budgets.DeleteBudgetInput) *BudgetsDeleteBudgetFuture

	DeleteBudgetAction(ctx workflow.Context, input *budgets.DeleteBudgetActionInput) (*budgets.DeleteBudgetActionOutput, error)
	DeleteBudgetActionAsync(ctx workflow.Context, input *budgets.DeleteBudgetActionInput) *BudgetsDeleteBudgetActionFuture

	DeleteNotification(ctx workflow.Context, input *budgets.DeleteNotificationInput) (*budgets.DeleteNotificationOutput, error)
	DeleteNotificationAsync(ctx workflow.Context, input *budgets.DeleteNotificationInput) *BudgetsDeleteNotificationFuture

	DeleteSubscriber(ctx workflow.Context, input *budgets.DeleteSubscriberInput) (*budgets.DeleteSubscriberOutput, error)
	DeleteSubscriberAsync(ctx workflow.Context, input *budgets.DeleteSubscriberInput) *BudgetsDeleteSubscriberFuture

	DescribeBudget(ctx workflow.Context, input *budgets.DescribeBudgetInput) (*budgets.DescribeBudgetOutput, error)
	DescribeBudgetAsync(ctx workflow.Context, input *budgets.DescribeBudgetInput) *BudgetsDescribeBudgetFuture

	DescribeBudgetAction(ctx workflow.Context, input *budgets.DescribeBudgetActionInput) (*budgets.DescribeBudgetActionOutput, error)
	DescribeBudgetActionAsync(ctx workflow.Context, input *budgets.DescribeBudgetActionInput) *BudgetsDescribeBudgetActionFuture

	DescribeBudgetActionHistories(ctx workflow.Context, input *budgets.DescribeBudgetActionHistoriesInput) (*budgets.DescribeBudgetActionHistoriesOutput, error)
	DescribeBudgetActionHistoriesAsync(ctx workflow.Context, input *budgets.DescribeBudgetActionHistoriesInput) *BudgetsDescribeBudgetActionHistoriesFuture

	DescribeBudgetActionsForAccount(ctx workflow.Context, input *budgets.DescribeBudgetActionsForAccountInput) (*budgets.DescribeBudgetActionsForAccountOutput, error)
	DescribeBudgetActionsForAccountAsync(ctx workflow.Context, input *budgets.DescribeBudgetActionsForAccountInput) *BudgetsDescribeBudgetActionsForAccountFuture

	DescribeBudgetActionsForBudget(ctx workflow.Context, input *budgets.DescribeBudgetActionsForBudgetInput) (*budgets.DescribeBudgetActionsForBudgetOutput, error)
	DescribeBudgetActionsForBudgetAsync(ctx workflow.Context, input *budgets.DescribeBudgetActionsForBudgetInput) *BudgetsDescribeBudgetActionsForBudgetFuture

	DescribeBudgetPerformanceHistory(ctx workflow.Context, input *budgets.DescribeBudgetPerformanceHistoryInput) (*budgets.DescribeBudgetPerformanceHistoryOutput, error)
	DescribeBudgetPerformanceHistoryAsync(ctx workflow.Context, input *budgets.DescribeBudgetPerformanceHistoryInput) *BudgetsDescribeBudgetPerformanceHistoryFuture

	DescribeBudgets(ctx workflow.Context, input *budgets.DescribeBudgetsInput) (*budgets.DescribeBudgetsOutput, error)
	DescribeBudgetsAsync(ctx workflow.Context, input *budgets.DescribeBudgetsInput) *BudgetsDescribeBudgetsFuture

	DescribeNotificationsForBudget(ctx workflow.Context, input *budgets.DescribeNotificationsForBudgetInput) (*budgets.DescribeNotificationsForBudgetOutput, error)
	DescribeNotificationsForBudgetAsync(ctx workflow.Context, input *budgets.DescribeNotificationsForBudgetInput) *BudgetsDescribeNotificationsForBudgetFuture

	DescribeSubscribersForNotification(ctx workflow.Context, input *budgets.DescribeSubscribersForNotificationInput) (*budgets.DescribeSubscribersForNotificationOutput, error)
	DescribeSubscribersForNotificationAsync(ctx workflow.Context, input *budgets.DescribeSubscribersForNotificationInput) *BudgetsDescribeSubscribersForNotificationFuture

	ExecuteBudgetAction(ctx workflow.Context, input *budgets.ExecuteBudgetActionInput) (*budgets.ExecuteBudgetActionOutput, error)
	ExecuteBudgetActionAsync(ctx workflow.Context, input *budgets.ExecuteBudgetActionInput) *BudgetsExecuteBudgetActionFuture

	UpdateBudget(ctx workflow.Context, input *budgets.UpdateBudgetInput) (*budgets.UpdateBudgetOutput, error)
	UpdateBudgetAsync(ctx workflow.Context, input *budgets.UpdateBudgetInput) *BudgetsUpdateBudgetFuture

	UpdateBudgetAction(ctx workflow.Context, input *budgets.UpdateBudgetActionInput) (*budgets.UpdateBudgetActionOutput, error)
	UpdateBudgetActionAsync(ctx workflow.Context, input *budgets.UpdateBudgetActionInput) *BudgetsUpdateBudgetActionFuture

	UpdateNotification(ctx workflow.Context, input *budgets.UpdateNotificationInput) (*budgets.UpdateNotificationOutput, error)
	UpdateNotificationAsync(ctx workflow.Context, input *budgets.UpdateNotificationInput) *BudgetsUpdateNotificationFuture

	UpdateSubscriber(ctx workflow.Context, input *budgets.UpdateSubscriberInput) (*budgets.UpdateSubscriberOutput, error)
	UpdateSubscriberAsync(ctx workflow.Context, input *budgets.UpdateSubscriberInput) *BudgetsUpdateSubscriberFuture
}

func NewClient() Client {
	return &stub{}
}
