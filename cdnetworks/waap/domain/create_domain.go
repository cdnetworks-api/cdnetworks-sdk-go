// This file is auto-generated, don't edit it. Thanks.
package domain

import (
	"github.com/alibabacloud-go/tea/tea"
)

type AccessDomainRequest struct {
	// {"en":"WAF.", "zh_CN":"WAF。"}
	WafDefendConfig *WafConfig `json:"wafDefendConfig,omitempty" xml:"wafDefendConfig,omitempty"`
	// {"en":"Custom rules.", "zh_CN":"自定义规则。"}
	CustomizeRuleConfig *CustomizeRuleConfig `json:"customizeRuleConfig,omitempty" xml:"customizeRuleConfig,omitempty"`
	// {"en":"API security.", "zh_CN":"API安全。"}
	ApiDefendConfig *APIDefendConfig `json:"apiDefendConfig,omitempty" xml:"apiDefendConfig,omitempty"`
	// {"en":"Whitelist.", "zh_CN":"白名单。"}
	WhitelistConfig *WhitelistConfig `json:"whitelistConfig,omitempty" xml:"whitelistConfig,omitempty"`
	// {"en":"Hostnames to be accessed.", "zh_CN":"目标域名。"}
	TargetDomains []*string `json:"targetDomains,omitempty" xml:"targetDomains,omitempty" require:"true" type:"Repeated"`
	// {"en":"IP/Geo blocking.", "zh_CN":"IP/区域封禁。"}
	BlockConfig *BlockConfig `json:"blockConfig,omitempty" xml:"blockConfig,omitempty"`
	// {"en":"DDoS protection.", "zh_CN":"DDoS防护。"}
	DmsDefendConfig *DMSConfig `json:"dmsDefendConfig,omitempty" xml:"dmsDefendConfig,omitempty"`
	// {"en":"Threat intelligence.", "zh_CN":"威胁情报。"}
	IntelligenceConfig *IntelligenceConfig `json:"intelligenceConfig,omitempty" xml:"intelligenceConfig,omitempty"`
	// {"en":"Bot management.", "zh_CN":"Bot管理。"}
	BotManageConfig *BOTConfig `json:"botManageConfig,omitempty" xml:"botManageConfig,omitempty"`
	// {"en":"Rate limiting.", "zh_CN":"频率限制。"}
	RateLimitConfig *RateLimitConfig `json:"rateLimitConfig,omitempty" xml:"rateLimitConfig,omitempty"`
}

func (s AccessDomainRequest) String() string {
	return tea.Prettify(s)
}

func (s AccessDomainRequest) GoString() string {
	return s.String()
}

func (s *AccessDomainRequest) SetWafDefendConfig(v *WafConfig) *AccessDomainRequest {
	s.WafDefendConfig = v
	return s
}

func (s *AccessDomainRequest) SetCustomizeRuleConfig(v *CustomizeRuleConfig) *AccessDomainRequest {
	s.CustomizeRuleConfig = v
	return s
}

func (s *AccessDomainRequest) SetApiDefendConfig(v *APIDefendConfig) *AccessDomainRequest {
	s.ApiDefendConfig = v
	return s
}

func (s *AccessDomainRequest) SetWhitelistConfig(v *WhitelistConfig) *AccessDomainRequest {
	s.WhitelistConfig = v
	return s
}

func (s *AccessDomainRequest) SetTargetDomains(v []*string) *AccessDomainRequest {
	s.TargetDomains = v
	return s
}

func (s *AccessDomainRequest) SetBlockConfig(v *BlockConfig) *AccessDomainRequest {
	s.BlockConfig = v
	return s
}

func (s *AccessDomainRequest) SetDmsDefendConfig(v *DMSConfig) *AccessDomainRequest {
	s.DmsDefendConfig = v
	return s
}

func (s *AccessDomainRequest) SetIntelligenceConfig(v *IntelligenceConfig) *AccessDomainRequest {
	s.IntelligenceConfig = v
	return s
}

func (s *AccessDomainRequest) SetBotManageConfig(v *BOTConfig) *AccessDomainRequest {
	s.BotManageConfig = v
	return s
}

func (s *AccessDomainRequest) SetRateLimitConfig(v *RateLimitConfig) *AccessDomainRequest {
	s.RateLimitConfig = v
	return s
}

type AccessDomainResponse struct {
	// {'en':'Please refer to the error code for exceptions.', 'zh_CN':'请参照错误码。','dictionary':'belong=WAAP-MS-Ext|dict=waap_retCodeEnum'}
	Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
	// {"en":"Description.", "zh_CN":"描述信息。"}
	Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s AccessDomainResponse) String() string {
	return tea.Prettify(s)
}

func (s AccessDomainResponse) GoString() string {
	return s.String()
}

func (s *AccessDomainResponse) SetCode(v string) *AccessDomainResponse {
	s.Code = &v
	return s
}

func (s *AccessDomainResponse) SetMsg(v string) *AccessDomainResponse {
	s.Message = &v
	return s
}
