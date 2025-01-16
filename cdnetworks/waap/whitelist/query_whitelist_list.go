// This file is auto-generated, don't edit it. Thanks.
package whitelist

import (
	"github.com/alibabacloud-go/tea/tea"
)

type ListWhitelistRulesRequest struct {
	// {"en":"Hostname list.", "zh_CN":"域名列表。"}
	DomainList []*string `json:"domainList,omitempty" xml:"domainList,omitempty" require:"true" type:"Repeated"`
	// {"en":"Rule name, fuzzy query.", "zh_CN":"规则名称，模糊查询。"}
	RuleName *string `json:"ruleName,omitempty" xml:"ruleName,omitempty"`
}

func (s ListWhitelistRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListWhitelistRulesRequest) GoString() string {
	return s.String()
}

func (s *ListWhitelistRulesRequest) SetDomainList(v []*string) *ListWhitelistRulesRequest {
	s.DomainList = v
	return s
}

func (s *ListWhitelistRulesRequest) SetRuleName(v string) *ListWhitelistRulesRequest {
	s.RuleName = &v
	return s
}

type ListWhitelistRulesResponse struct {
	// {'en':'Please refer to the error code for exceptions.', 'zh_CN':'请参照错误码。','dictionary':'belong=WAAP-MS-Ext|dict=waap_retCodeEnum'}
	Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
	// {"en":"Description.", "zh_CN":"描述信息。"}
	Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
	// {"en":"Data.", "zh_CN":"出参数据。"}
	Data []*CommonWhitelistVO `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s ListWhitelistRulesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListWhitelistRulesResponse) GoString() string {
	return s.String()
}

func (s *ListWhitelistRulesResponse) SetCode(v string) *ListWhitelistRulesResponse {
	s.Code = &v
	return s
}

func (s *ListWhitelistRulesResponse) SetMessage(v string) *ListWhitelistRulesResponse {
	s.Message = &v
	return s
}

func (s *ListWhitelistRulesResponse) SetData(v []*CommonWhitelistVO) *ListWhitelistRulesResponse {
	s.Data = v
	return s
}

type CommonWhitelistVO struct {
	// {"en":"Rule ID.", "zh_CN":"规则ID。"}
	Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
	// {"en":"Hostname.", "zh_CN":"域名。"}
	Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
	// {"en":"Rule name.", "zh_CN":"规则名称。"}
	RuleName *string `json:"ruleName,omitempty" xml:"ruleName,omitempty" require:"true"`
	// {"en":"Description.", "zh_CN":"描述。"}
	Description *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	// {"en":"Match conditions, at least one, at most five.", "zh_CN":"匹配条件，至少一个，至多五个。"}
	Conditions *WhitelistRuleCondition `json:"conditions,omitempty" xml:"conditions,omitempty" require:"true"`
	// {"en":"Created date, format: yyyy-MM-dd HH:mm:ss.", "zh_CN":"创建时间，格式：yyyy-MM-dd HH:mm:ss。"}
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty" require:"true"`
	// {"en":"Update date, format: yyyy-MM-dd HH:mm:ss.", "zh_CN":"更新时间，格式：yyyy-MM-dd HH:mm:ss。"}
	UpdateTime *string `json:"updateTime,omitempty" xml:"updateTime,omitempty" require:"true"`
}

func (s CommonWhitelistVO) String() string {
	return tea.Prettify(s)
}

func (s CommonWhitelistVO) GoString() string {
	return s.String()
}

func (s *CommonWhitelistVO) SetId(v string) *CommonWhitelistVO {
	s.Id = &v
	return s
}

func (s *CommonWhitelistVO) SetDomain(v string) *CommonWhitelistVO {
	s.Domain = &v
	return s
}

func (s *CommonWhitelistVO) SetRuleName(v string) *CommonWhitelistVO {
	s.RuleName = &v
	return s
}

func (s *CommonWhitelistVO) SetDescription(v string) *CommonWhitelistVO {
	s.Description = &v
	return s
}

func (s *CommonWhitelistVO) SetConditions(v *WhitelistRuleCondition) *CommonWhitelistVO {
	s.Conditions = v
	return s
}

func (s *CommonWhitelistVO) SetCreateTime(v string) *CommonWhitelistVO {
	s.CreateTime = &v
	return s
}

func (s *CommonWhitelistVO) SetUpdateTime(v string) *CommonWhitelistVO {
	s.UpdateTime = &v
	return s
}
