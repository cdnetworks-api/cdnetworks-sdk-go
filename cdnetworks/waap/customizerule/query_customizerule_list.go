// This file is auto-generated, don't edit it. Thanks.
package customizerule

import (
	"github.com/alibabacloud-go/tea/tea"
)

type ListCustomRulesRequest struct {
	// {'en':'Hostname list.', 'zh_CN':'域名列表。'}
	DomainList []*string `json:"domainList,omitempty" xml:"domainList,omitempty" require:"true" type:"Repeated"`
	// {'en':'Rule name, fuzzy query.', 'zh_CN':'规则名称，模糊查询。'}
	RuleName *string `json:"ruleName,omitempty" xml:"ruleName,omitempty"`
}

func (s ListCustomRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListCustomRulesRequest) GoString() string {
	return s.String()
}

func (s *ListCustomRulesRequest) SetDomainList(v []*string) *ListCustomRulesRequest {
	s.DomainList = v
	return s
}

func (s *ListCustomRulesRequest) SetRuleName(v string) *ListCustomRulesRequest {
	s.RuleName = &v
	return s
}

type ListCustomRulesResponse struct {
	// {'en':'Please refer to the error code for exceptions.', 'zh_CN':'请参照错误码。','dictionary':'belong=WAAP-MS-Ext|dict=waap_retCodeEnum'}
	Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
	// {'en':'Description.', 'zh_CN':'描述信息。'}
	Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
	// {'en':'Data.', 'zh_CN':'出参数据。'}
	Data []*CommonCustomizeRuleVO `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s ListCustomRulesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListCustomRulesResponse) GoString() string {
	return s.String()
}

func (s *ListCustomRulesResponse) SetCode(v string) *ListCustomRulesResponse {
	s.Code = &v
	return s
}

func (s *ListCustomRulesResponse) SetMessage(v string) *ListCustomRulesResponse {
	s.Message = &v
	return s
}

func (s *ListCustomRulesResponse) SetData(v []*CommonCustomizeRuleVO) *ListCustomRulesResponse {
	s.Data = v
	return s
}

type CommonCustomizeRuleVO struct {
	// {'en':'Rule ID.', 'zh_CN':'规则ID。'}
	Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
	// {'en':'Hostname.', 'zh_CN':'域名。'}
	Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
	// {'en':'Rule name.', 'zh_CN':'规则名称。'}
	RuleName *string `json:"ruleName,omitempty" xml:"ruleName,omitempty" require:"true"`
	// {'en':'Description.', 'zh_CN':'描述。'}
	Description *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	// {'en':'Protected target.
	// WEB:Website
	// API:API', 'zh_CN':'业务场景。
	// WEB：WEB业务
	// API：API业务'}
	Scene *string `json:"scene,omitempty" xml:"scene,omitempty" require:"true"`
	// {'en':'API ID, multiple separated by ; sign.', 'zh_CN':'API业务下的API ID，多个用 ; 隔开。'}
	ApiId *string `json:"apiId,omitempty" xml:"apiId,omitempty" require:"true"`
	// {'en':'Action.
	// NO_USE:Not Used
	// LOG:Log
	// DELAY:Delay
	// BLOCK:Deny
	// RESET:Reset Connection', 'zh_CN':'处理动作。
	// NO_USE：不使用
	// LOG：监控
	// DELAY：延迟响应
	// BLOCK：拦截
	// RESET:断开连接'}
	Act *string `json:"act,omitempty" xml:"act,omitempty" require:"true"`
	// {'en':'Match Conditions.', 'zh_CN':'匹配条件。'}
	ConditionList *CommonCustomizeRuleConditionDTO `json:"conditionList,omitempty" xml:"conditionList,omitempty" require:"true"`
}

func (s CommonCustomizeRuleVO) String() string {
	return tea.Prettify(s)
}

func (s CommonCustomizeRuleVO) GoString() string {
	return s.String()
}

func (s *CommonCustomizeRuleVO) SetId(v string) *CommonCustomizeRuleVO {
	s.Id = &v
	return s
}

func (s *CommonCustomizeRuleVO) SetDomain(v string) *CommonCustomizeRuleVO {
	s.Domain = &v
	return s
}

func (s *CommonCustomizeRuleVO) SetRuleName(v string) *CommonCustomizeRuleVO {
	s.RuleName = &v
	return s
}

func (s *CommonCustomizeRuleVO) SetDescription(v string) *CommonCustomizeRuleVO {
	s.Description = &v
	return s
}

func (s *CommonCustomizeRuleVO) SetScene(v string) *CommonCustomizeRuleVO {
	s.Scene = &v
	return s
}

func (s *CommonCustomizeRuleVO) SetApiId(v string) *CommonCustomizeRuleVO {
	s.ApiId = &v
	return s
}

func (s *CommonCustomizeRuleVO) SetAct(v string) *CommonCustomizeRuleVO {
	s.Act = &v
	return s
}

func (s *CommonCustomizeRuleVO) SetConditionList(v *CommonCustomizeRuleConditionDTO) *CommonCustomizeRuleVO {
	s.ConditionList = v
	return s
}
