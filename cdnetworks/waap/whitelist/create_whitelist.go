// This file is auto-generated, don't edit it. Thanks.
package whitelist

import (
	"github.com/alibabacloud-go/tea/tea"
)

type CreateWhitelistRuleRequest struct {
	// {"en":"Hostname.", "zh_CN":"域名。"}
	Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
	// {"en":"Rule name, maximum 50 characters.
	//  Does not support special characters and spaces.", "zh_CN":"规则名称，最多50个字符。
	// 不支持特殊字符和空格。"}
	RuleName *string `json:"ruleName,omitempty" xml:"ruleName,omitempty" require:"true"`
	// {"en":"Description, maximum 200 characters.", "zh_CN":"描述，最多200个字符。"}
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// {"en":"Match conditions, at least one, at most five.", "zh_CN":"匹配条件，至少一个，至多五个。"}
	Conditions *WhitelistRuleCondition `json:"conditions,omitempty" xml:"conditions,omitempty" require:"true"`
}

func (s CreateWhitelistRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateWhitelistRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateWhitelistRuleRequest) SetDomain(v string) *CreateWhitelistRuleRequest {
	s.Domain = &v
	return s
}

func (s *CreateWhitelistRuleRequest) SetRuleName(v string) *CreateWhitelistRuleRequest {
	s.RuleName = &v
	return s
}

func (s *CreateWhitelistRuleRequest) SetDescription(v string) *CreateWhitelistRuleRequest {
	s.Description = &v
	return s
}

func (s *CreateWhitelistRuleRequest) SetConditions(v *WhitelistRuleCondition) *CreateWhitelistRuleRequest {
	s.Conditions = v
	return s
}

type CreateWhitelistRuleResponse struct {
	// {'en':'Please refer to the error code for exceptions.', 'zh_CN':'请参照错误码。','dictionary':'belong=WAAP-MS-Ext|dict=waap_retCodeEnum'}
	Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
	// {"en":"Description.", "zh_CN":"描述信息。"}
	Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
	// {"en":"Rule ID.", "zh_CN":"规则 ID。"}
	Data *string `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s CreateWhitelistRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateWhitelistRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateWhitelistRuleResponse) SetCode(v string) *CreateWhitelistRuleResponse {
	s.Code = &v
	return s
}

func (s *CreateWhitelistRuleResponse) SetMessage(v string) *CreateWhitelistRuleResponse {
	s.Message = &v
	return s
}

func (s *CreateWhitelistRuleResponse) SetData(v string) *CreateWhitelistRuleResponse {
	s.Data = &v
	return s
}
