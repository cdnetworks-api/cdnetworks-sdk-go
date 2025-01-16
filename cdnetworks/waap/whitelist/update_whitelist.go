// This file is auto-generated, don't edit it. Thanks.
package whitelist

import (
	"github.com/alibabacloud-go/tea/tea"
)

type UpdateWhitelistRuleRequest struct {
	// {"en":"Rule ID.", "zh_CN":"规则ID。"}
	Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
	// {"en":"Rule name, maximum 50 characters.
	//  Does not support special characters and spaces.", "zh_CN":"规则名称，最多50个字符。
	// 不支持特殊字符和空格。"}
	RuleName *string `json:"ruleName,omitempty" xml:"ruleName,omitempty"`
	// {"en":"Description, maximum 200 characters.", "zh_CN":"描述，最多200个字符。"}
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// {"en":"Match conditions, at least one, at most five.", "zh_CN":"匹配条件，至少一个，至多五个。"}
	Conditions *WhitelistRuleCondition `json:"conditions,omitempty" xml:"conditions,omitempty"`
}

func (s UpdateWhitelistRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateWhitelistRuleRequest) GoString() string {
	return s.String()
}

func (s *UpdateWhitelistRuleRequest) SetId(v string) *UpdateWhitelistRuleRequest {
	s.Id = &v
	return s
}

func (s *UpdateWhitelistRuleRequest) SetRuleName(v string) *UpdateWhitelistRuleRequest {
	s.RuleName = &v
	return s
}

func (s *UpdateWhitelistRuleRequest) SetDescription(v string) *UpdateWhitelistRuleRequest {
	s.Description = &v
	return s
}

func (s *UpdateWhitelistRuleRequest) SetConditions(v *WhitelistRuleCondition) *UpdateWhitelistRuleRequest {
	s.Conditions = v
	return s
}

type UpdateWhitelistRuleResponse struct {
	// {'en':'Please refer to the error code for exceptions.', 'zh_CN':'请参照错误码。','dictionary':'belong=WAAP-MS-Ext|dict=waap_retCodeEnum'}
	Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
	// {"en":"Description.", "zh_CN":"描述信息。"}
	Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s UpdateWhitelistRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateWhitelistRuleResponse) GoString() string {
	return s.String()
}

func (s *UpdateWhitelistRuleResponse) SetCode(v string) *UpdateWhitelistRuleResponse {
	s.Code = &v
	return s
}

func (s *UpdateWhitelistRuleResponse) SetMessage(v string) *UpdateWhitelistRuleResponse {
	s.Message = &v
	return s
}
