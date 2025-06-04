// This file is auto-generated, don't edit it. Thanks.
package share_customizerule

import (
	"github.com/alibabacloud-go/tea/tea"
)

type ListSharedCustomRulesRequest struct {
	// {"en":"Rule name, fuzzy query.","zh_CN":"规则名称，模糊查询。"}
	RuleName *string `json:"ruleName,omitempty" xml:"ruleName,omitempty"`
}

func (s ListSharedCustomRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSharedCustomRulesRequest) GoString() string {
	return s.String()
}

func (s *ListSharedCustomRulesRequest) SetRuleName(v string) *ListSharedCustomRulesRequest {
	s.RuleName = &v
	return s
}

type ListSharedCustomRulesResponse struct {
	// {"en":"Description.","zh_CN":"描述信息。"}
	Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
	// {"dictionary":"belong=WAAP-MS-Ext|dict=waap_retCodeEnum","en":"Please refer to the error code for exceptions.","zh_CN":"请参照错误码。"}
	Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
	// {"en":"Data.","zh_CN":"出参数据。"}
	Data []*ListSharedCustomRulesResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s ListSharedCustomRulesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSharedCustomRulesResponse) GoString() string {
	return s.String()
}

func (s *ListSharedCustomRulesResponse) SetMsg(v string) *ListSharedCustomRulesResponse {
	s.Msg = &v
	return s
}

func (s *ListSharedCustomRulesResponse) SetCode(v string) *ListSharedCustomRulesResponse {
	s.Code = &v
	return s
}

func (s *ListSharedCustomRulesResponse) SetData(v []*ListSharedCustomRulesResponseData) *ListSharedCustomRulesResponse {
	s.Data = v
	return s
}

type ListSharedCustomRulesResponseData struct {
	// {"en":"Matching conditions.","zh_CN":"匹配条件。"}
	Condition *ShareCustomizeRuleCondition `json:"condition,omitempty" xml:"condition,omitempty" require:"true" type:"Struct"`
	// {"en":"Creator.","zh_CN":"创建者。"}
	Creator *string `json:"creator,omitempty" xml:"creator,omitempty" require:"true"`
	// {"en":"Action.
	// NO_USE: Not Used
	// LOG: Log
	// DELAY: Delay
	// BLOCK: Deny
	// RESET: Reset Connection","zh_CN":"处理动作。
	// NO_USE：不使用
	// LOG：监控
	// DELAY：延迟响应
	// BLOCK：拦截
	// RESET：断开连接"}
	Act *string `json:"act,omitempty" xml:"act,omitempty" require:"true"`
	// {"en":"Create time, format: yyyy-MM-dd HH:mm:ss.","zh_CN":"创建时间，格式：yyyy-MM-dd HH:mm:ss。"}
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty" require:"true"`
	// {"en":"Associated hostname.", "zh_CN":"关联的域名。"}
	RelationDomainList []*string `json:"relationDomainList,omitempty" xml:"relationDomainList,omitempty" require:"true" type:"Repeated"`
	// {"en":"Rule Name.","zh_CN":"规则名称。"}
	RuleName *string `json:"ruleName,omitempty" xml:"ruleName,omitempty" require:"true"`
	// {"en":"Description.","zh_CN":"规则描述。"}
	Description *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	// {"en":"Update time, format: yyyy-MM-dd HH:mm:ss.","zh_CN":"更新时间，格式：yyyy-MM-dd HH:mm:ss。"}
	UpdateTime *string `json:"updateTime,omitempty" xml:"updateTime,omitempty" require:"true"`
	// {"en":"Rule ID.","zh_CN":"规则ID。"}
	Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
}

func (s ListSharedCustomRulesResponseData) String() string {
	return tea.Prettify(s)
}

func (s ListSharedCustomRulesResponseData) GoString() string {
	return s.String()
}

func (s *ListSharedCustomRulesResponseData) SetCondition(v *ShareCustomizeRuleCondition) *ListSharedCustomRulesResponseData {
	s.Condition = v
	return s
}

func (s *ListSharedCustomRulesResponseData) SetCreator(v string) *ListSharedCustomRulesResponseData {
	s.Creator = &v
	return s
}

func (s *ListSharedCustomRulesResponseData) SetAct(v string) *ListSharedCustomRulesResponseData {
	s.Act = &v
	return s
}

func (s *ListSharedCustomRulesResponseData) SetCreateTime(v string) *ListSharedCustomRulesResponseData {
	s.CreateTime = &v
	return s
}

func (s *ListSharedCustomRulesResponseData) SetRelationDomainList(v []*string) *ListSharedCustomRulesResponseData {
	s.RelationDomainList = v
	return s
}

func (s *ListSharedCustomRulesResponseData) SetRuleName(v string) *ListSharedCustomRulesResponseData {
	s.RuleName = &v
	return s
}

func (s *ListSharedCustomRulesResponseData) SetDescription(v string) *ListSharedCustomRulesResponseData {
	s.Description = &v
	return s
}

func (s *ListSharedCustomRulesResponseData) SetUpdateTime(v string) *ListSharedCustomRulesResponseData {
	s.UpdateTime = &v
	return s
}

func (s *ListSharedCustomRulesResponseData) SetId(v string) *ListSharedCustomRulesResponseData {
	s.Id = &v
	return s
}
