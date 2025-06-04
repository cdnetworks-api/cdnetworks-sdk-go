// This file is auto-generated, don't edit it. Thanks.
package share_whitelist

import (
	"github.com/alibabacloud-go/tea/tea"
)

type ListShareWhitelistRulesRequest struct {
	// {"en":"Rule name, fuzzy query.", "zh_CN":"规则名称，模糊查询。"}
	RuleName *string `json:"ruleName,omitempty" xml:"ruleName,omitempty"`
}

func (s ListShareWhitelistRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListShareWhitelistRulesRequest) GoString() string {
	return s.String()
}

func (s *ListShareWhitelistRulesRequest) SetRuleName(v string) *ListShareWhitelistRulesRequest {
	s.RuleName = &v
	return s
}

type ListShareWhitelistRulesResponse struct {
	// {'en':'Please refer to the error code for exceptions.', 'zh_CN':'请参照错误码。','dictionary':'belong=WAAP-MS-Ext|dict=waap_retCodeEnum'}
	Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
	// {"en":"Description.", "zh_CN":"描述信息。"}
	Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
	// {"en":"Data.", "zh_CN":"出参数据。"}
	Data []*CommonShareWhitelistVO `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s ListShareWhitelistRulesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListShareWhitelistRulesResponse) GoString() string {
	return s.String()
}

func (s *ListShareWhitelistRulesResponse) SetCode(v string) *ListShareWhitelistRulesResponse {
	s.Code = &v
	return s
}

func (s *ListShareWhitelistRulesResponse) SetMsg(v string) *ListShareWhitelistRulesResponse {
	s.Msg = &v
	return s
}

func (s *ListShareWhitelistRulesResponse) SetData(v []*CommonShareWhitelistVO) *ListShareWhitelistRulesResponse {
	s.Data = v
	return s
}

type CommonShareWhitelistVO struct {
	// {"en":"Rule ID.", "zh_CN":"规则ID。"}
	Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
	// {"en":"Rule name.", "zh_CN":"规则名称。"}
	RuleName *string `json:"ruleName,omitempty" xml:"ruleName,omitempty" require:"true"`
	// {"en":"Description.", "zh_CN":"规则描述。"}
	Description *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	// {"en":"creator.", "zh_CN":"创建者。"}
	Creator *string `json:"creator,omitempty" xml:"creator,omitempty" require:"true"`
	// {"en":"Associated hostname.", "zh_CN":"关联的域名。"}
	RelationDomainList []*string `json:"importedDomain,omitempty" xml:"importedDomain,omitempty" require:"true" type:"Repeated"`
	// {"en":"Match conditions, At least one, at most five.", "zh_CN":"匹配条件，至少一个，至多五个。"}
	Conditions *ShareWhitelistRuleConditions `json:"conditions,omitempty" xml:"conditions,omitempty" require:"true"`
	// {"en":"Created date, format: yyyy-MM-dd HH:mm:ss.", "zh_CN":"创建时间，格式：yyyy-MM-dd HH:mm:ss。"}
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty" require:"true"`
	// {"en":"Update date,format: yyyy-MM-dd HH:mm:ss.", "zh_CN":"更新时间，格式：yyyy-MM-dd HH:mm:ss。"}
	UpdateTime *string `json:"updateTime,omitempty" xml:"updateTime,omitempty" require:"true"`
}

func (s CommonShareWhitelistVO) String() string {
	return tea.Prettify(s)
}

func (s CommonShareWhitelistVO) GoString() string {
	return s.String()
}

func (s *CommonShareWhitelistVO) SetId(v string) *CommonShareWhitelistVO {
	s.Id = &v
	return s
}

func (s *CommonShareWhitelistVO) SetRuleName(v string) *CommonShareWhitelistVO {
	s.RuleName = &v
	return s
}

func (s *CommonShareWhitelistVO) SetDescription(v string) *CommonShareWhitelistVO {
	s.Description = &v
	return s
}

func (s *CommonShareWhitelistVO) SetCreator(v string) *CommonShareWhitelistVO {
	s.Creator = &v
	return s
}

func (s *CommonShareWhitelistVO) SetImportedDomain(v []*string) *CommonShareWhitelistVO {
	s.RelationDomainList = v
	return s
}

func (s *CommonShareWhitelistVO) SetConditions(v *ShareWhitelistRuleConditions) *CommonShareWhitelistVO {
	s.Conditions = v
	return s
}

func (s *CommonShareWhitelistVO) SetCreateTime(v string) *CommonShareWhitelistVO {
	s.CreateTime = &v
	return s
}

func (s *CommonShareWhitelistVO) SetUpdateTime(v string) *CommonShareWhitelistVO {
	s.UpdateTime = &v
	return s
}
