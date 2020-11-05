// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7

package wafstub

import (
	"github.com/aws/aws-sdk-go/service/waf"
	"go.uber.org/cadence/workflow"

	"github.com/banzaicloud/cadence-aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	CreateByteMatchSet(ctx workflow.Context, input *waf.CreateByteMatchSetInput) (*waf.CreateByteMatchSetOutput, error)
	CreateByteMatchSetAsync(ctx workflow.Context, input *waf.CreateByteMatchSetInput) *WAFCreateByteMatchSetFuture

	CreateGeoMatchSet(ctx workflow.Context, input *waf.CreateGeoMatchSetInput) (*waf.CreateGeoMatchSetOutput, error)
	CreateGeoMatchSetAsync(ctx workflow.Context, input *waf.CreateGeoMatchSetInput) *WAFCreateGeoMatchSetFuture

	CreateIPSet(ctx workflow.Context, input *waf.CreateIPSetInput) (*waf.CreateIPSetOutput, error)
	CreateIPSetAsync(ctx workflow.Context, input *waf.CreateIPSetInput) *WAFCreateIPSetFuture

	CreateRateBasedRule(ctx workflow.Context, input *waf.CreateRateBasedRuleInput) (*waf.CreateRateBasedRuleOutput, error)
	CreateRateBasedRuleAsync(ctx workflow.Context, input *waf.CreateRateBasedRuleInput) *WAFCreateRateBasedRuleFuture

	CreateRegexMatchSet(ctx workflow.Context, input *waf.CreateRegexMatchSetInput) (*waf.CreateRegexMatchSetOutput, error)
	CreateRegexMatchSetAsync(ctx workflow.Context, input *waf.CreateRegexMatchSetInput) *WAFCreateRegexMatchSetFuture

	CreateRegexPatternSet(ctx workflow.Context, input *waf.CreateRegexPatternSetInput) (*waf.CreateRegexPatternSetOutput, error)
	CreateRegexPatternSetAsync(ctx workflow.Context, input *waf.CreateRegexPatternSetInput) *WAFCreateRegexPatternSetFuture

	CreateRule(ctx workflow.Context, input *waf.CreateRuleInput) (*waf.CreateRuleOutput, error)
	CreateRuleAsync(ctx workflow.Context, input *waf.CreateRuleInput) *WAFCreateRuleFuture

	CreateRuleGroup(ctx workflow.Context, input *waf.CreateRuleGroupInput) (*waf.CreateRuleGroupOutput, error)
	CreateRuleGroupAsync(ctx workflow.Context, input *waf.CreateRuleGroupInput) *WAFCreateRuleGroupFuture

	CreateSizeConstraintSet(ctx workflow.Context, input *waf.CreateSizeConstraintSetInput) (*waf.CreateSizeConstraintSetOutput, error)
	CreateSizeConstraintSetAsync(ctx workflow.Context, input *waf.CreateSizeConstraintSetInput) *WAFCreateSizeConstraintSetFuture

	CreateSqlInjectionMatchSet(ctx workflow.Context, input *waf.CreateSqlInjectionMatchSetInput) (*waf.CreateSqlInjectionMatchSetOutput, error)
	CreateSqlInjectionMatchSetAsync(ctx workflow.Context, input *waf.CreateSqlInjectionMatchSetInput) *WAFCreateSqlInjectionMatchSetFuture

	CreateWebACL(ctx workflow.Context, input *waf.CreateWebACLInput) (*waf.CreateWebACLOutput, error)
	CreateWebACLAsync(ctx workflow.Context, input *waf.CreateWebACLInput) *WAFCreateWebACLFuture

	CreateWebACLMigrationStack(ctx workflow.Context, input *waf.CreateWebACLMigrationStackInput) (*waf.CreateWebACLMigrationStackOutput, error)
	CreateWebACLMigrationStackAsync(ctx workflow.Context, input *waf.CreateWebACLMigrationStackInput) *WAFCreateWebACLMigrationStackFuture

	CreateXssMatchSet(ctx workflow.Context, input *waf.CreateXssMatchSetInput) (*waf.CreateXssMatchSetOutput, error)
	CreateXssMatchSetAsync(ctx workflow.Context, input *waf.CreateXssMatchSetInput) *WAFCreateXssMatchSetFuture

	DeleteByteMatchSet(ctx workflow.Context, input *waf.DeleteByteMatchSetInput) (*waf.DeleteByteMatchSetOutput, error)
	DeleteByteMatchSetAsync(ctx workflow.Context, input *waf.DeleteByteMatchSetInput) *WAFDeleteByteMatchSetFuture

	DeleteGeoMatchSet(ctx workflow.Context, input *waf.DeleteGeoMatchSetInput) (*waf.DeleteGeoMatchSetOutput, error)
	DeleteGeoMatchSetAsync(ctx workflow.Context, input *waf.DeleteGeoMatchSetInput) *WAFDeleteGeoMatchSetFuture

	DeleteIPSet(ctx workflow.Context, input *waf.DeleteIPSetInput) (*waf.DeleteIPSetOutput, error)
	DeleteIPSetAsync(ctx workflow.Context, input *waf.DeleteIPSetInput) *WAFDeleteIPSetFuture

	DeleteLoggingConfiguration(ctx workflow.Context, input *waf.DeleteLoggingConfigurationInput) (*waf.DeleteLoggingConfigurationOutput, error)
	DeleteLoggingConfigurationAsync(ctx workflow.Context, input *waf.DeleteLoggingConfigurationInput) *WAFDeleteLoggingConfigurationFuture

	DeletePermissionPolicy(ctx workflow.Context, input *waf.DeletePermissionPolicyInput) (*waf.DeletePermissionPolicyOutput, error)
	DeletePermissionPolicyAsync(ctx workflow.Context, input *waf.DeletePermissionPolicyInput) *WAFDeletePermissionPolicyFuture

	DeleteRateBasedRule(ctx workflow.Context, input *waf.DeleteRateBasedRuleInput) (*waf.DeleteRateBasedRuleOutput, error)
	DeleteRateBasedRuleAsync(ctx workflow.Context, input *waf.DeleteRateBasedRuleInput) *WAFDeleteRateBasedRuleFuture

	DeleteRegexMatchSet(ctx workflow.Context, input *waf.DeleteRegexMatchSetInput) (*waf.DeleteRegexMatchSetOutput, error)
	DeleteRegexMatchSetAsync(ctx workflow.Context, input *waf.DeleteRegexMatchSetInput) *WAFDeleteRegexMatchSetFuture

	DeleteRegexPatternSet(ctx workflow.Context, input *waf.DeleteRegexPatternSetInput) (*waf.DeleteRegexPatternSetOutput, error)
	DeleteRegexPatternSetAsync(ctx workflow.Context, input *waf.DeleteRegexPatternSetInput) *WAFDeleteRegexPatternSetFuture

	DeleteRule(ctx workflow.Context, input *waf.DeleteRuleInput) (*waf.DeleteRuleOutput, error)
	DeleteRuleAsync(ctx workflow.Context, input *waf.DeleteRuleInput) *WAFDeleteRuleFuture

	DeleteRuleGroup(ctx workflow.Context, input *waf.DeleteRuleGroupInput) (*waf.DeleteRuleGroupOutput, error)
	DeleteRuleGroupAsync(ctx workflow.Context, input *waf.DeleteRuleGroupInput) *WAFDeleteRuleGroupFuture

	DeleteSizeConstraintSet(ctx workflow.Context, input *waf.DeleteSizeConstraintSetInput) (*waf.DeleteSizeConstraintSetOutput, error)
	DeleteSizeConstraintSetAsync(ctx workflow.Context, input *waf.DeleteSizeConstraintSetInput) *WAFDeleteSizeConstraintSetFuture

	DeleteSqlInjectionMatchSet(ctx workflow.Context, input *waf.DeleteSqlInjectionMatchSetInput) (*waf.DeleteSqlInjectionMatchSetOutput, error)
	DeleteSqlInjectionMatchSetAsync(ctx workflow.Context, input *waf.DeleteSqlInjectionMatchSetInput) *WAFDeleteSqlInjectionMatchSetFuture

	DeleteWebACL(ctx workflow.Context, input *waf.DeleteWebACLInput) (*waf.DeleteWebACLOutput, error)
	DeleteWebACLAsync(ctx workflow.Context, input *waf.DeleteWebACLInput) *WAFDeleteWebACLFuture

	DeleteXssMatchSet(ctx workflow.Context, input *waf.DeleteXssMatchSetInput) (*waf.DeleteXssMatchSetOutput, error)
	DeleteXssMatchSetAsync(ctx workflow.Context, input *waf.DeleteXssMatchSetInput) *WAFDeleteXssMatchSetFuture

	GetByteMatchSet(ctx workflow.Context, input *waf.GetByteMatchSetInput) (*waf.GetByteMatchSetOutput, error)
	GetByteMatchSetAsync(ctx workflow.Context, input *waf.GetByteMatchSetInput) *WAFGetByteMatchSetFuture

	GetChangeToken(ctx workflow.Context, input *waf.GetChangeTokenInput) (*waf.GetChangeTokenOutput, error)
	GetChangeTokenAsync(ctx workflow.Context, input *waf.GetChangeTokenInput) *WAFGetChangeTokenFuture

	GetChangeTokenStatus(ctx workflow.Context, input *waf.GetChangeTokenStatusInput) (*waf.GetChangeTokenStatusOutput, error)
	GetChangeTokenStatusAsync(ctx workflow.Context, input *waf.GetChangeTokenStatusInput) *WAFGetChangeTokenStatusFuture

	GetGeoMatchSet(ctx workflow.Context, input *waf.GetGeoMatchSetInput) (*waf.GetGeoMatchSetOutput, error)
	GetGeoMatchSetAsync(ctx workflow.Context, input *waf.GetGeoMatchSetInput) *WAFGetGeoMatchSetFuture

	GetIPSet(ctx workflow.Context, input *waf.GetIPSetInput) (*waf.GetIPSetOutput, error)
	GetIPSetAsync(ctx workflow.Context, input *waf.GetIPSetInput) *WAFGetIPSetFuture

	GetLoggingConfiguration(ctx workflow.Context, input *waf.GetLoggingConfigurationInput) (*waf.GetLoggingConfigurationOutput, error)
	GetLoggingConfigurationAsync(ctx workflow.Context, input *waf.GetLoggingConfigurationInput) *WAFGetLoggingConfigurationFuture

	GetPermissionPolicy(ctx workflow.Context, input *waf.GetPermissionPolicyInput) (*waf.GetPermissionPolicyOutput, error)
	GetPermissionPolicyAsync(ctx workflow.Context, input *waf.GetPermissionPolicyInput) *WAFGetPermissionPolicyFuture

	GetRateBasedRule(ctx workflow.Context, input *waf.GetRateBasedRuleInput) (*waf.GetRateBasedRuleOutput, error)
	GetRateBasedRuleAsync(ctx workflow.Context, input *waf.GetRateBasedRuleInput) *WAFGetRateBasedRuleFuture

	GetRateBasedRuleManagedKeys(ctx workflow.Context, input *waf.GetRateBasedRuleManagedKeysInput) (*waf.GetRateBasedRuleManagedKeysOutput, error)
	GetRateBasedRuleManagedKeysAsync(ctx workflow.Context, input *waf.GetRateBasedRuleManagedKeysInput) *WAFGetRateBasedRuleManagedKeysFuture

	GetRegexMatchSet(ctx workflow.Context, input *waf.GetRegexMatchSetInput) (*waf.GetRegexMatchSetOutput, error)
	GetRegexMatchSetAsync(ctx workflow.Context, input *waf.GetRegexMatchSetInput) *WAFGetRegexMatchSetFuture

	GetRegexPatternSet(ctx workflow.Context, input *waf.GetRegexPatternSetInput) (*waf.GetRegexPatternSetOutput, error)
	GetRegexPatternSetAsync(ctx workflow.Context, input *waf.GetRegexPatternSetInput) *WAFGetRegexPatternSetFuture

	GetRule(ctx workflow.Context, input *waf.GetRuleInput) (*waf.GetRuleOutput, error)
	GetRuleAsync(ctx workflow.Context, input *waf.GetRuleInput) *WAFGetRuleFuture

	GetRuleGroup(ctx workflow.Context, input *waf.GetRuleGroupInput) (*waf.GetRuleGroupOutput, error)
	GetRuleGroupAsync(ctx workflow.Context, input *waf.GetRuleGroupInput) *WAFGetRuleGroupFuture

	GetSampledRequests(ctx workflow.Context, input *waf.GetSampledRequestsInput) (*waf.GetSampledRequestsOutput, error)
	GetSampledRequestsAsync(ctx workflow.Context, input *waf.GetSampledRequestsInput) *WAFGetSampledRequestsFuture

	GetSizeConstraintSet(ctx workflow.Context, input *waf.GetSizeConstraintSetInput) (*waf.GetSizeConstraintSetOutput, error)
	GetSizeConstraintSetAsync(ctx workflow.Context, input *waf.GetSizeConstraintSetInput) *WAFGetSizeConstraintSetFuture

	GetSqlInjectionMatchSet(ctx workflow.Context, input *waf.GetSqlInjectionMatchSetInput) (*waf.GetSqlInjectionMatchSetOutput, error)
	GetSqlInjectionMatchSetAsync(ctx workflow.Context, input *waf.GetSqlInjectionMatchSetInput) *WAFGetSqlInjectionMatchSetFuture

	GetWebACL(ctx workflow.Context, input *waf.GetWebACLInput) (*waf.GetWebACLOutput, error)
	GetWebACLAsync(ctx workflow.Context, input *waf.GetWebACLInput) *WAFGetWebACLFuture

	GetXssMatchSet(ctx workflow.Context, input *waf.GetXssMatchSetInput) (*waf.GetXssMatchSetOutput, error)
	GetXssMatchSetAsync(ctx workflow.Context, input *waf.GetXssMatchSetInput) *WAFGetXssMatchSetFuture

	ListActivatedRulesInRuleGroup(ctx workflow.Context, input *waf.ListActivatedRulesInRuleGroupInput) (*waf.ListActivatedRulesInRuleGroupOutput, error)
	ListActivatedRulesInRuleGroupAsync(ctx workflow.Context, input *waf.ListActivatedRulesInRuleGroupInput) *WAFListActivatedRulesInRuleGroupFuture

	ListByteMatchSets(ctx workflow.Context, input *waf.ListByteMatchSetsInput) (*waf.ListByteMatchSetsOutput, error)
	ListByteMatchSetsAsync(ctx workflow.Context, input *waf.ListByteMatchSetsInput) *WAFListByteMatchSetsFuture

	ListGeoMatchSets(ctx workflow.Context, input *waf.ListGeoMatchSetsInput) (*waf.ListGeoMatchSetsOutput, error)
	ListGeoMatchSetsAsync(ctx workflow.Context, input *waf.ListGeoMatchSetsInput) *WAFListGeoMatchSetsFuture

	ListIPSets(ctx workflow.Context, input *waf.ListIPSetsInput) (*waf.ListIPSetsOutput, error)
	ListIPSetsAsync(ctx workflow.Context, input *waf.ListIPSetsInput) *WAFListIPSetsFuture

	ListLoggingConfigurations(ctx workflow.Context, input *waf.ListLoggingConfigurationsInput) (*waf.ListLoggingConfigurationsOutput, error)
	ListLoggingConfigurationsAsync(ctx workflow.Context, input *waf.ListLoggingConfigurationsInput) *WAFListLoggingConfigurationsFuture

	ListRateBasedRules(ctx workflow.Context, input *waf.ListRateBasedRulesInput) (*waf.ListRateBasedRulesOutput, error)
	ListRateBasedRulesAsync(ctx workflow.Context, input *waf.ListRateBasedRulesInput) *WAFListRateBasedRulesFuture

	ListRegexMatchSets(ctx workflow.Context, input *waf.ListRegexMatchSetsInput) (*waf.ListRegexMatchSetsOutput, error)
	ListRegexMatchSetsAsync(ctx workflow.Context, input *waf.ListRegexMatchSetsInput) *WAFListRegexMatchSetsFuture

	ListRegexPatternSets(ctx workflow.Context, input *waf.ListRegexPatternSetsInput) (*waf.ListRegexPatternSetsOutput, error)
	ListRegexPatternSetsAsync(ctx workflow.Context, input *waf.ListRegexPatternSetsInput) *WAFListRegexPatternSetsFuture

	ListRuleGroups(ctx workflow.Context, input *waf.ListRuleGroupsInput) (*waf.ListRuleGroupsOutput, error)
	ListRuleGroupsAsync(ctx workflow.Context, input *waf.ListRuleGroupsInput) *WAFListRuleGroupsFuture

	ListRules(ctx workflow.Context, input *waf.ListRulesInput) (*waf.ListRulesOutput, error)
	ListRulesAsync(ctx workflow.Context, input *waf.ListRulesInput) *WAFListRulesFuture

	ListSizeConstraintSets(ctx workflow.Context, input *waf.ListSizeConstraintSetsInput) (*waf.ListSizeConstraintSetsOutput, error)
	ListSizeConstraintSetsAsync(ctx workflow.Context, input *waf.ListSizeConstraintSetsInput) *WAFListSizeConstraintSetsFuture

	ListSqlInjectionMatchSets(ctx workflow.Context, input *waf.ListSqlInjectionMatchSetsInput) (*waf.ListSqlInjectionMatchSetsOutput, error)
	ListSqlInjectionMatchSetsAsync(ctx workflow.Context, input *waf.ListSqlInjectionMatchSetsInput) *WAFListSqlInjectionMatchSetsFuture

	ListSubscribedRuleGroups(ctx workflow.Context, input *waf.ListSubscribedRuleGroupsInput) (*waf.ListSubscribedRuleGroupsOutput, error)
	ListSubscribedRuleGroupsAsync(ctx workflow.Context, input *waf.ListSubscribedRuleGroupsInput) *WAFListSubscribedRuleGroupsFuture

	ListTagsForResource(ctx workflow.Context, input *waf.ListTagsForResourceInput) (*waf.ListTagsForResourceOutput, error)
	ListTagsForResourceAsync(ctx workflow.Context, input *waf.ListTagsForResourceInput) *WAFListTagsForResourceFuture

	ListWebACLs(ctx workflow.Context, input *waf.ListWebACLsInput) (*waf.ListWebACLsOutput, error)
	ListWebACLsAsync(ctx workflow.Context, input *waf.ListWebACLsInput) *WAFListWebACLsFuture

	ListXssMatchSets(ctx workflow.Context, input *waf.ListXssMatchSetsInput) (*waf.ListXssMatchSetsOutput, error)
	ListXssMatchSetsAsync(ctx workflow.Context, input *waf.ListXssMatchSetsInput) *WAFListXssMatchSetsFuture

	PutLoggingConfiguration(ctx workflow.Context, input *waf.PutLoggingConfigurationInput) (*waf.PutLoggingConfigurationOutput, error)
	PutLoggingConfigurationAsync(ctx workflow.Context, input *waf.PutLoggingConfigurationInput) *WAFPutLoggingConfigurationFuture

	PutPermissionPolicy(ctx workflow.Context, input *waf.PutPermissionPolicyInput) (*waf.PutPermissionPolicyOutput, error)
	PutPermissionPolicyAsync(ctx workflow.Context, input *waf.PutPermissionPolicyInput) *WAFPutPermissionPolicyFuture

	TagResource(ctx workflow.Context, input *waf.TagResourceInput) (*waf.TagResourceOutput, error)
	TagResourceAsync(ctx workflow.Context, input *waf.TagResourceInput) *WAFTagResourceFuture

	UntagResource(ctx workflow.Context, input *waf.UntagResourceInput) (*waf.UntagResourceOutput, error)
	UntagResourceAsync(ctx workflow.Context, input *waf.UntagResourceInput) *WAFUntagResourceFuture

	UpdateByteMatchSet(ctx workflow.Context, input *waf.UpdateByteMatchSetInput) (*waf.UpdateByteMatchSetOutput, error)
	UpdateByteMatchSetAsync(ctx workflow.Context, input *waf.UpdateByteMatchSetInput) *WAFUpdateByteMatchSetFuture

	UpdateGeoMatchSet(ctx workflow.Context, input *waf.UpdateGeoMatchSetInput) (*waf.UpdateGeoMatchSetOutput, error)
	UpdateGeoMatchSetAsync(ctx workflow.Context, input *waf.UpdateGeoMatchSetInput) *WAFUpdateGeoMatchSetFuture

	UpdateIPSet(ctx workflow.Context, input *waf.UpdateIPSetInput) (*waf.UpdateIPSetOutput, error)
	UpdateIPSetAsync(ctx workflow.Context, input *waf.UpdateIPSetInput) *WAFUpdateIPSetFuture

	UpdateRateBasedRule(ctx workflow.Context, input *waf.UpdateRateBasedRuleInput) (*waf.UpdateRateBasedRuleOutput, error)
	UpdateRateBasedRuleAsync(ctx workflow.Context, input *waf.UpdateRateBasedRuleInput) *WAFUpdateRateBasedRuleFuture

	UpdateRegexMatchSet(ctx workflow.Context, input *waf.UpdateRegexMatchSetInput) (*waf.UpdateRegexMatchSetOutput, error)
	UpdateRegexMatchSetAsync(ctx workflow.Context, input *waf.UpdateRegexMatchSetInput) *WAFUpdateRegexMatchSetFuture

	UpdateRegexPatternSet(ctx workflow.Context, input *waf.UpdateRegexPatternSetInput) (*waf.UpdateRegexPatternSetOutput, error)
	UpdateRegexPatternSetAsync(ctx workflow.Context, input *waf.UpdateRegexPatternSetInput) *WAFUpdateRegexPatternSetFuture

	UpdateRule(ctx workflow.Context, input *waf.UpdateRuleInput) (*waf.UpdateRuleOutput, error)
	UpdateRuleAsync(ctx workflow.Context, input *waf.UpdateRuleInput) *WAFUpdateRuleFuture

	UpdateRuleGroup(ctx workflow.Context, input *waf.UpdateRuleGroupInput) (*waf.UpdateRuleGroupOutput, error)
	UpdateRuleGroupAsync(ctx workflow.Context, input *waf.UpdateRuleGroupInput) *WAFUpdateRuleGroupFuture

	UpdateSizeConstraintSet(ctx workflow.Context, input *waf.UpdateSizeConstraintSetInput) (*waf.UpdateSizeConstraintSetOutput, error)
	UpdateSizeConstraintSetAsync(ctx workflow.Context, input *waf.UpdateSizeConstraintSetInput) *WAFUpdateSizeConstraintSetFuture

	UpdateSqlInjectionMatchSet(ctx workflow.Context, input *waf.UpdateSqlInjectionMatchSetInput) (*waf.UpdateSqlInjectionMatchSetOutput, error)
	UpdateSqlInjectionMatchSetAsync(ctx workflow.Context, input *waf.UpdateSqlInjectionMatchSetInput) *WAFUpdateSqlInjectionMatchSetFuture

	UpdateWebACL(ctx workflow.Context, input *waf.UpdateWebACLInput) (*waf.UpdateWebACLOutput, error)
	UpdateWebACLAsync(ctx workflow.Context, input *waf.UpdateWebACLInput) *WAFUpdateWebACLFuture

	UpdateXssMatchSet(ctx workflow.Context, input *waf.UpdateXssMatchSetInput) (*waf.UpdateXssMatchSetOutput, error)
	UpdateXssMatchSetAsync(ctx workflow.Context, input *waf.UpdateXssMatchSetInput) *WAFUpdateXssMatchSetFuture
}

func NewClient() Client {
	return &stub{}
}
