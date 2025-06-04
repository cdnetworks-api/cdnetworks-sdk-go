// This file is auto-generated, don't edit it. Thanks.
package share_whitelist

import (
	"github.com/alibabacloud-go/tea/tea"
)

type CreateShareWhitelistRuleRequest struct {
	// {"en":"","zh_CN":""}
	RelationDomainList []*string `json:"relationDomainList,omitempty" xml:"relationDomainList,omitempty" type:"Repeated"`
	// {"en":"Rule name, maximum 50 characters.\nDoes not support special characters and spaces.","zh_CN":"规则名称，最多50个字符。\n不支持特殊字符和空格。"}
	RuleName *string `json:"ruleName,omitempty" xml:"ruleName,omitempty" require:"true"`
	// {"en":"Description, maximum 200 characters.","zh_CN":"描述，最多200个字符。"}
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// {"en":"Match conditions, at least one, at most five.","zh_CN":"匹配条件，至少一个，至多五个。"}
	Conditions *ShareWhitelistRuleConditions `json:"conditions,omitempty" xml:"conditions,omitempty" require:"true" type:"Struct"`
}

func (s CreateShareWhitelistRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateShareWhitelistRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateShareWhitelistRuleRequest) SetRelationDomainList(v []*string) *CreateShareWhitelistRuleRequest {
	s.RelationDomainList = v
	return s
}

func (s *CreateShareWhitelistRuleRequest) SetRuleName(v string) *CreateShareWhitelistRuleRequest {
	s.RuleName = &v
	return s
}

func (s *CreateShareWhitelistRuleRequest) SetDescription(v string) *CreateShareWhitelistRuleRequest {
	s.Description = &v
	return s
}

func (s *CreateShareWhitelistRuleRequest) SetConditions(v *ShareWhitelistRuleConditions) *CreateShareWhitelistRuleRequest {
	s.Conditions = v
	return s
}

type CreateShareWhitelistRuleResponse struct {
	// {"en":"Description.","zh_CN":"描述信息。"}
	Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
	// {"dictionary":"belong=WAAP-MS-Ext|dict=waap_retCodeEnum","en":"Please refer to the error code for exceptions.","zh_CN":"请参照错误码。"}
	Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
	// {"en":"Rule ID.","zh_CN":"规则ID。"}
	Data *string `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s CreateShareWhitelistRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateShareWhitelistRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateShareWhitelistRuleResponse) SetMsg(v string) *CreateShareWhitelistRuleResponse {
	s.Msg = &v
	return s
}

func (s *CreateShareWhitelistRuleResponse) SetCode(v string) *CreateShareWhitelistRuleResponse {
	s.Code = &v
	return s
}

func (s *CreateShareWhitelistRuleResponse) SetData(v string) *CreateShareWhitelistRuleResponse {
	s.Data = &v
	return s
}
