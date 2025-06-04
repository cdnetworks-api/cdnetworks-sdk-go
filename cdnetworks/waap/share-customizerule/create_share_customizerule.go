// This file is auto-generated, don't edit it. Thanks.
package share_customizerule

import (
	"github.com/alibabacloud-go/tea/tea"
)

type CreateSharedCustomRuleRequest struct {
	// {"en":"Matching conditions.\nExcept for header conditions, there can only be at most one record per match type under each type of condition.","zh_CN":"匹配条件。\n除了请求头条件，其它类型的条件下一种匹配类型最多只能有一条记录。"}
	Condition *ShareCustomizeRuleCondition `json:"condition,omitempty" xml:"condition,omitempty" require:"true" type:"Struct"`
	// {"en":"Action.\nNO_USE: Not Used\nLOG: Log\nDELAY: Delay\nBLOCK: Deny\nRESET: Reset Connection","zh_CN":"处理动作。\nNO_USE：不使用\nLOG：监控\nDELAY：延迟响应\nBLOCK：拦截\nRESET：断开连接"}
	Act *string `json:"act,omitempty" xml:"act,omitempty" require:"true"`
	// {"en":"List of hostname to be associated.","zh_CN":"待关联域名列表。"}
	RelationDomainList []*string `json:"relationDomainList,omitempty" xml:"relationDomainList,omitempty" type:"Repeated"`
	// {"en":"Rule Name, maximum 50 characters.\nDoes not support special characters and spaces.","zh_CN":"规则名称，最多50个字符。\n不支持特殊字符和空格。"}
	RuleName *string `json:"ruleName,omitempty" xml:"ruleName,omitempty" require:"true"`
	// {"en":"Description, maximum 200 characters.","zh_CN":"规则描述，最多200个字符。"}
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
}

func (s CreateSharedCustomRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSharedCustomRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateSharedCustomRuleRequest) SetCondition(v *ShareCustomizeRuleCondition) *CreateSharedCustomRuleRequest {
	s.Condition = v
	return s
}

func (s *CreateSharedCustomRuleRequest) SetAct(v string) *CreateSharedCustomRuleRequest {
	s.Act = &v
	return s
}

func (s *CreateSharedCustomRuleRequest) SetRelationDomainList(v []*string) *CreateSharedCustomRuleRequest {
	s.RelationDomainList = v
	return s
}

func (s *CreateSharedCustomRuleRequest) SetRuleName(v string) *CreateSharedCustomRuleRequest {
	s.RuleName = &v
	return s
}

func (s *CreateSharedCustomRuleRequest) SetDescription(v string) *CreateSharedCustomRuleRequest {
	s.Description = &v
	return s
}

type CreateSharedCustomRuleResponse struct {
	// {"en":"Description.","zh_CN":"描述信息。"}
	Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
	// {"dictionary":"belong=WAAP-MS-Ext|dict=waap_retCodeEnum","en":"Please refer to the error code for exceptions.","zh_CN":"请参照错误码。"}
	Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
	// {"en":"Rule ID.","zh_CN":"规则ID。"}
	Data *string `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s CreateSharedCustomRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSharedCustomRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateSharedCustomRuleResponse) SetMsg(v string) *CreateSharedCustomRuleResponse {
	s.Msg = &v
	return s
}

func (s *CreateSharedCustomRuleResponse) SetCode(v string) *CreateSharedCustomRuleResponse {
	s.Code = &v
	return s
}

func (s *CreateSharedCustomRuleResponse) SetData(v string) *CreateSharedCustomRuleResponse {
	s.Data = &v
	return s
}
