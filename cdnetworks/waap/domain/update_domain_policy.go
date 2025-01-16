// This file is auto-generated, don't edit it. Thanks.
package domain

import (
	"github.com/alibabacloud-go/tea/tea"
)

type ModifyPolicyStatusRequest struct {
	// {"en":"Hostname list.", "zh_CN":"域名列表。"}
	DomainList []*string `json:"domainList,omitempty" xml:"domainList,omitempty" require:"true" type:"Repeated"`
	// {"en":"Bot management switch.
	// ON: Enabled
	// OFF: Disabled", "zh_CN":"Bot管理开关。
	// ON：开启
	// OFF：关闭"}
	BotManageSwitch *string `json:"botManageSwitch,omitempty" xml:"botManageSwitch,omitempty"`
	// {"en":"Custom rules switch.
	// ON: Enabled
	// OFF: Disabled", "zh_CN":"自定义规则开关。
	// ON：开启
	// OFF：关闭"}
	CustomizeRuleSwitch *string `json:"customizeRuleSwitch,omitempty" xml:"customizeRuleSwitch,omitempty"`
	// {"en":"IP/Geo blocking switch.
	// ON: Enabled
	// OFF: Disabled", "zh_CN":"IP/区域封禁开关。
	// ON：开启
	// OFF：关闭"}
	BlockSwitch *string `json:"blockSwitch,omitempty" xml:"blockSwitch,omitempty"`
	// {"en":"Rate limiting switch.
	// ON: Enabled
	// OFF: Disabled", "zh_CN":"频率限制开关。
	// ON：开启
	// OFF：关闭"}
	RateLimitSwitch *string `json:"rateLimitSwitch,omitempty" xml:"rateLimitSwitch,omitempty"`
	// {"en":"Whitelist switch.
	// ON: Enabled
	// OFF: Disabled", "zh_CN":"白名单开关。
	// ON：开启
	// OFF：关闭"}
	WhitelistSwitch *string `json:"whitelistSwitch,omitempty" xml:"whitelistSwitch,omitempty"`
	// {"en":"API security switch.
	// ON: Enabled
	// OFF: Disabled", "zh_CN":"API安全开关。
	// ON：开启
	// OFF：关闭"}
	ApiDefendSwitch *string `json:"apiDefendSwitch,omitempty" xml:"apiDefendSwitch,omitempty"`
	// {"en":"DDoS protection switch.
	// ON: Enabled
	// OFF: Disabled", "zh_CN":"DDoS防护开关。
	// ON：开启
	// OFF：关闭"}
	DmsDefendSwitch *string `json:"dmsDefendSwitch,omitempty" xml:"dmsDefendSwitch,omitempty"`
	// {"en":"Threat intelligence switch.
	// ON: Enabled
	// OFF: Disabled", "zh_CN":"威胁情报开关。
	// ON：开启
	// OFF：关闭"}
	IntelligenceSwitch *string `json:"intelligenceSwitch,omitempty" xml:"intelligenceSwitch,omitempty"`
	// {"en":"WAF protection switch.
	// ON: Enabled
	// OFF: Disabled", "zh_CN":"WAF防护开关。
	// ON：开启
	// OFF：关闭"}
	WafDefendSwitch *string `json:"wafDefendSwitch,omitempty" xml:"wafDefendSwitch,omitempty"`
}

func (s ModifyPolicyStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyPolicyStatusRequest) GoString() string {
	return s.String()
}

func (s *ModifyPolicyStatusRequest) SetDomainList(v []*string) *ModifyPolicyStatusRequest {
	s.DomainList = v
	return s
}

func (s *ModifyPolicyStatusRequest) SetBotManageSwitch(v string) *ModifyPolicyStatusRequest {
	s.BotManageSwitch = &v
	return s
}

func (s *ModifyPolicyStatusRequest) SetCustomizeRuleSwitch(v string) *ModifyPolicyStatusRequest {
	s.CustomizeRuleSwitch = &v
	return s
}

func (s *ModifyPolicyStatusRequest) SetBlockSwitch(v string) *ModifyPolicyStatusRequest {
	s.BlockSwitch = &v
	return s
}

func (s *ModifyPolicyStatusRequest) SetRateLimitSwitch(v string) *ModifyPolicyStatusRequest {
	s.RateLimitSwitch = &v
	return s
}

func (s *ModifyPolicyStatusRequest) SetWhitelistSwitch(v string) *ModifyPolicyStatusRequest {
	s.WhitelistSwitch = &v
	return s
}

func (s *ModifyPolicyStatusRequest) SetApiDefendSwitch(v string) *ModifyPolicyStatusRequest {
	s.ApiDefendSwitch = &v
	return s
}

func (s *ModifyPolicyStatusRequest) SetDmsDefendSwitch(v string) *ModifyPolicyStatusRequest {
	s.DmsDefendSwitch = &v
	return s
}

func (s *ModifyPolicyStatusRequest) SetIntelligenceSwitch(v string) *ModifyPolicyStatusRequest {
	s.IntelligenceSwitch = &v
	return s
}

func (s *ModifyPolicyStatusRequest) SetWafDefendSwitch(v string) *ModifyPolicyStatusRequest {
	s.WafDefendSwitch = &v
	return s
}

type ModifyPolicyStatusResponse struct {
	// {'en':'Please refer to the error code for exceptions.', 'zh_CN':'请参照错误码。','dictionary':'belong=WAAP-MS-Ext|dict=waap_retCodeEnum'}
	Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
	// {"en":"Description.", "zh_CN":"描述信息。"}
	Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s ModifyPolicyStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyPolicyStatusResponse) GoString() string {
	return s.String()
}

func (s *ModifyPolicyStatusResponse) SetCode(v string) *ModifyPolicyStatusResponse {
	s.Code = &v
	return s
}

func (s *ModifyPolicyStatusResponse) SetMessage(v string) *ModifyPolicyStatusResponse {
	s.Message = &v
	return s
}
