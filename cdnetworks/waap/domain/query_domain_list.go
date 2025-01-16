// This file is auto-generated, don't edit it. Thanks.
package domain

import (
	"github.com/alibabacloud-go/tea/tea"
)

type ListDomainInfoRequest struct {
	// {"en":"Protection status, If not specified, it means all the protection status.
	// PROTECTING: Protecting
	// UNPROTECTED: Unprotected", "zh_CN":"防护状态，未指定时查询所有防护状态。
	// PROTECTING：防护中
	// UNPROTECTED：未防护"}
	DefendStatus *string `json:"defendStatus,omitempty" xml:"defendStatus,omitempty"`
	// {"en":"Hostname list, if not specified, it means all the hostnames of the account.", "zh_CN":"域名列表，未指定时查询账号下的所有域名。"}
	DomainList []*string `json:"domainList,omitempty" xml:"domainList,omitempty" type:"Repeated"`
	// {"en":"DDoS protection switch, if not specified, it means all the status.
	// ON: Enabled
	// OFF: Disabled", "zh_CN":"DDoS防护开关，未指定时查询该开关所有状态。
	// ON：开启
	// OFF：关闭"}
	DmsDefendSwitch *string `json:"dmsDefendSwitch,omitempty" xml:"dmsDefendSwitch,omitempty"`
	// {"en":"Rate limiting switch, if not specified,  it means all the status.
	// ON: Enabled
	// OFF: Disabled", "zh_CN":"频率限制开关，未指定时查询该开关所有状态。
	// ON：开启
	// OFF：关闭"}
	RateLimitSwitch *string `json:"rateLimitSwitch,omitempty" xml:"rateLimitSwitch,omitempty"`
	// {"en":"IP/Geo blocking switch, if not specified,  it means all the status.
	// ON: Enabled
	// OFF: Disabled", "zh_CN":"IP/区域封禁开关，未指定时查询该开关所有状态。
	// ON：开启
	// OFF：关闭"}
	BlockSwitch *string `json:"blockSwitch,omitempty" xml:"blockSwitch,omitempty"`
	// {"en":"WAF protection switch, if not specified, it means all the status.
	// ON: Enabled
	// OFF: Disabled", "zh_CN":"WAF防护开关，未指定时查询该开关所有状态。
	// ON：开启
	// OFF：关闭"}
	WafDefendSwitch *string `json:"wafDefendSwitch,omitempty" xml:"wafDefendSwitch,omitempty"`
	// {"en":"Threat intelligence switch, if not specified,  it means all the status.
	// ON: Enabled
	// OFF: Disabled", "zh_CN":"威胁情报开关，未指定时查询该开关所有状态。
	// ON：开启
	// OFF：关闭"}
	IntelligenceSwitch *string `json:"intelligenceSwitch,omitempty" xml:"intelligenceSwitch,omitempty"`
	// {"en":"Whitelist switch, if not specified,  it means all the status.
	// ON: Enabled
	// OFF: Disabled", "zh_CN":"白名单开关，未指定时查询该开关所有状态。
	// ON：开启
	// OFF：关闭"}
	WhitelistSwitch *string `json:"whitelistSwitch,omitempty" xml:"whitelistSwitch,omitempty"`
	// {"en":"Bot management switch, if not specified,  it means all the status.
	// ON: Enabled
	// OFF: Disabled", "zh_CN":"Bot管理开关，未指定时查询该开关所有状态。
	// ON：开启
	// OFF：关闭"}
	BotManageSwitch *string `json:"botManageSwitch,omitempty" xml:"botManageSwitch,omitempty"`
	// {"en":"Custom rules switch, if not specified, it means all the status.
	// ON: Enabled
	// OFF: Disabled", "zh_CN":"自定义规则开关，未指定时查询该开关所有状态。
	// ON：开启
	// OFF：关闭"}
	CustomizeRuleSwitch *string `json:"customizeRuleSwitch,omitempty" xml:"customizeRuleSwitch,omitempty"`
	// {"en":"API security switch, if not specified, it means all the status.
	// ON: Enabled
	// OFF: Disabled", "zh_CN":"API安全开关，未指定时查询该开关所有状态。
	// ON：开启
	// OFF：关闭"}
	ApiDefendSwitch *string `json:"apiDefendSwitch,omitempty" xml:"apiDefendSwitch,omitempty"`
}

func (s ListDomainInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDomainInfoRequest) GoString() string {
	return s.String()
}

func (s *ListDomainInfoRequest) SetDefendStatus(v string) *ListDomainInfoRequest {
	s.DefendStatus = &v
	return s
}

func (s *ListDomainInfoRequest) SetDomainList(v []*string) *ListDomainInfoRequest {
	s.DomainList = v
	return s
}

func (s *ListDomainInfoRequest) SetDmsDefendSwitch(v string) *ListDomainInfoRequest {
	s.DmsDefendSwitch = &v
	return s
}

func (s *ListDomainInfoRequest) SetRateLimitSwitch(v string) *ListDomainInfoRequest {
	s.RateLimitSwitch = &v
	return s
}

func (s *ListDomainInfoRequest) SetBlockSwitch(v string) *ListDomainInfoRequest {
	s.BlockSwitch = &v
	return s
}

func (s *ListDomainInfoRequest) SetWafDefendSwitch(v string) *ListDomainInfoRequest {
	s.WafDefendSwitch = &v
	return s
}

func (s *ListDomainInfoRequest) SetIntelligenceSwitch(v string) *ListDomainInfoRequest {
	s.IntelligenceSwitch = &v
	return s
}

func (s *ListDomainInfoRequest) SetWhitelistSwitch(v string) *ListDomainInfoRequest {
	s.WhitelistSwitch = &v
	return s
}

func (s *ListDomainInfoRequest) SetBotManageSwitch(v string) *ListDomainInfoRequest {
	s.BotManageSwitch = &v
	return s
}

func (s *ListDomainInfoRequest) SetCustomizeRuleSwitch(v string) *ListDomainInfoRequest {
	s.CustomizeRuleSwitch = &v
	return s
}

func (s *ListDomainInfoRequest) SetApiDefendSwitch(v string) *ListDomainInfoRequest {
	s.ApiDefendSwitch = &v
	return s
}

type ListDomainInfoResponse struct {
	// {'en':'Please refer to the error code for exceptions.', 'zh_CN':'请参照错误码。','dictionary':'belong=WAAP-MS-Ext|dict=waap_retCodeEnum'}
	Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
	// {"en":"Description.", "zh_CN":"描述信息。"}
	Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
	// {"en":"Data.", "zh_CN":"出参数据。"}
	Data []*SysDomainInfoVO `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s ListDomainInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDomainInfoResponse) GoString() string {
	return s.String()
}

func (s *ListDomainInfoResponse) SetCode(v string) *ListDomainInfoResponse {
	s.Code = &v
	return s
}

func (s *ListDomainInfoResponse) SetMessage(v string) *ListDomainInfoResponse {
	s.Message = &v
	return s
}

func (s *ListDomainInfoResponse) SetData(v []*SysDomainInfoVO) *ListDomainInfoResponse {
	s.Data = v
	return s
}

type SysDomainInfoVO struct {
	// {"en":"ID.", "zh_CN":"ID。"}
	Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
	// {"en":"Hostname.", "zh_CN":"域名。"}
	Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
	// {"en":"Created time, format: yyyy-MM-dd HH:mm:ss.", "zh_CN":"接入时间，格式：yyyy-MM-dd HH:mm:ss。"}
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty" require:"true"`
	// {"en":"Deployment status.
	// DEPLOYING: Publishing
	// SUCCESS: Success", "zh_CN":"部署状态。
	// DEPLOYING：部署中
	// SUCCESS：部署成功"}
	DeployStatus *string `json:"deployStatus,omitempty" xml:"deployStatus,omitempty" require:"true"`
	// {"en":"IP/Geo blocking switch.
	// ON: Enabled
	// OFF: Disabled", "zh_CN":"IP/区域封禁开关。
	// ON：开启
	// OFF：关闭"}
	BlockSwitch *string `json:"blockSwitch,omitempty" xml:"blockSwitch,omitempty" require:"true"`
	// {"en":"Protection status.
	// PROTECTING: Protecting
	// UNPROTECTED: Unprotected", "zh_CN":"防护状态。
	// PROTECTING：防护中
	// UNPROTECTED：未防护"}
	DefendStatus *string `json:"defendStatus,omitempty" xml:"defendStatus,omitempty" require:"true"`
	// {"en":"DDoS protection switch.
	// ON: Enabled
	// OFF: Disabled", "zh_CN":"DDoS防护开关。
	// ON：开启
	// OFF：关闭"}
	DmsDefendSwitch *string `json:"dmsDefendSwitch,omitempty" xml:"dmsDefendSwitch,omitempty" require:"true"`
	// {"en":"Bot management switch.
	// ON: Enabled
	// OFF: Disabled", "zh_CN":"Bot管理开关。
	// ON：开启
	// OFF：关闭"}
	BotManageSwitch *string `json:"botManageSwitch,omitempty" xml:"botManageSwitch,omitempty" require:"true"`
	// {"en":"Custom rules switch.
	// ON: Enabled
	// OFF: Disabled", "zh_CN":"自定义规则开关。
	// ON：开启
	// OFF：关闭"}
	CustomizeRuleSwitch *string `json:"customizeRuleSwitch,omitempty" xml:"customizeRuleSwitch,omitempty" require:"true"`
	// {"en":"API security switch.
	// ON: Enabled
	// OFF: Disabled", "zh_CN":"API安全开关。
	// ON：开启
	// OFF：关闭"}
	ApiDefendSwitch *string `json:"apiDefendSwitch,omitempty" xml:"apiDefendSwitch,omitempty" require:"true"`
	// {"en":"Rate limiting switch.
	// ON: Enabled
	// OFF: Disabled", "zh_CN":"频率限制开关。
	// ON：开启
	// OFF：关闭"}
	RateLimitSwitch *string `json:"rateLimitSwitch,omitempty" xml:"rateLimitSwitch,omitempty" require:"true"`
	// {"en":"Whitelist switch.
	// ON: Enabled
	// OFF: Disabled", "zh_CN":"白名单开关。
	// ON：开启
	// OFF：关闭"}
	WhitelistSwitch *string `json:"whitelistSwitch,omitempty" xml:"whitelistSwitch,omitempty" require:"true"`
	// {"en":"Threat intelligence switch.
	// ON: Enabled
	// OFF: Disabled", "zh_CN":"威胁情报开关。
	// ON：开启
	// OFF：关闭"}
	IntelligenceSwitch *string `json:"intelligenceSwitch,omitempty" xml:"intelligenceSwitch,omitempty" require:"true"`
	// {"en":"WAF protection switch.
	// ON: Enabled
	// OFF: Disabled", "zh_CN":"WAF防护开关。
	// ON：开启
	// OFF：关闭"}
	WafDefendSwitch *string `json:"wafDefendSwitch,omitempty" xml:"wafDefendSwitch,omitempty" require:"true"`
}

func (s SysDomainInfoVO) String() string {
	return tea.Prettify(s)
}

func (s SysDomainInfoVO) GoString() string {
	return s.String()
}

func (s *SysDomainInfoVO) SetId(v string) *SysDomainInfoVO {
	s.Id = &v
	return s
}

func (s *SysDomainInfoVO) SetDomain(v string) *SysDomainInfoVO {
	s.Domain = &v
	return s
}

func (s *SysDomainInfoVO) SetCreateTime(v string) *SysDomainInfoVO {
	s.CreateTime = &v
	return s
}

func (s *SysDomainInfoVO) SetDeployStatus(v string) *SysDomainInfoVO {
	s.DeployStatus = &v
	return s
}

func (s *SysDomainInfoVO) SetBlockSwitch(v string) *SysDomainInfoVO {
	s.BlockSwitch = &v
	return s
}

func (s *SysDomainInfoVO) SetDefendStatus(v string) *SysDomainInfoVO {
	s.DefendStatus = &v
	return s
}

func (s *SysDomainInfoVO) SetDmsDefendSwitch(v string) *SysDomainInfoVO {
	s.DmsDefendSwitch = &v
	return s
}

func (s *SysDomainInfoVO) SetBotManageSwitch(v string) *SysDomainInfoVO {
	s.BotManageSwitch = &v
	return s
}

func (s *SysDomainInfoVO) SetCustomizeRuleSwitch(v string) *SysDomainInfoVO {
	s.CustomizeRuleSwitch = &v
	return s
}

func (s *SysDomainInfoVO) SetApiDefendSwitch(v string) *SysDomainInfoVO {
	s.ApiDefendSwitch = &v
	return s
}

func (s *SysDomainInfoVO) SetRateLimitSwitch(v string) *SysDomainInfoVO {
	s.RateLimitSwitch = &v
	return s
}

func (s *SysDomainInfoVO) SetWhitelistSwitch(v string) *SysDomainInfoVO {
	s.WhitelistSwitch = &v
	return s
}

func (s *SysDomainInfoVO) SetIntelligenceSwitch(v string) *SysDomainInfoVO {
	s.IntelligenceSwitch = &v
	return s
}

func (s *SysDomainInfoVO) SetWafDefendSwitch(v string) *SysDomainInfoVO {
	s.WafDefendSwitch = &v
	return s
}
