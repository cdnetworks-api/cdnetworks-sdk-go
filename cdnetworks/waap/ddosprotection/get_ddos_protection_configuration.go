// This file is auto-generated, don't edit it. Thanks.
package ddosprotection

import (
	"github.com/alibabacloud-go/tea/tea"
)

type GetDDoSProtectionConfigurationRequest struct {
	// {"en":"Domain list.","zh_CN":"域名列表。"}
	DomainList []*string `json:"domainList,omitempty" xml:"domainList,omitempty" require:"true" type:"Repeated"`
}

func (s GetDDoSProtectionConfigurationRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDDoSProtectionConfigurationRequest) GoString() string {
	return s.String()
}

func (s *GetDDoSProtectionConfigurationRequest) SetDomainList(v []*string) *GetDDoSProtectionConfigurationRequest {
	s.DomainList = v
	return s
}

type GetDDoSProtectionConfigurationResponse struct {
	// {"dictionary":"belong=WAAP-MS-Ext|dict=waap_retCodeEnum","en":"Please refer to the error code for exceptions.","zh_CN":"请参照错误码。"}
	Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
	// {"en":"Description.","zh_CN":"描述信息。"}
	Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
	// {"en":"Data.","zh_CN":"出参数据。"}
	Data *GetDDoSProtectionConfigurationResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
}

func (s GetDDoSProtectionConfigurationResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDDoSProtectionConfigurationResponse) GoString() string {
	return s.String()
}

func (s *GetDDoSProtectionConfigurationResponse) SetCode(v string) *GetDDoSProtectionConfigurationResponse {
	s.Code = &v
	return s
}

func (s *GetDDoSProtectionConfigurationResponse) SetMsg(v string) *GetDDoSProtectionConfigurationResponse {
	s.Msg = &v
	return s
}

func (s *GetDDoSProtectionConfigurationResponse) SetData(v *GetDDoSProtectionConfigurationResponseData) *GetDDoSProtectionConfigurationResponse {
	s.Data = v
	return s
}

type GetDDoSProtectionConfigurationResponseData struct {
	// {"en":"Array.","zh_CN":"数组。"}
	Array []*DDoSProtectionDataArray `json:"array,omitempty" xml:"array,omitempty" require:"true" type:"Repeated"`
}

func (s GetDDoSProtectionConfigurationResponseData) String() string {
	return tea.Prettify(s)
}

func (s GetDDoSProtectionConfigurationResponseData) GoString() string {
	return s.String()
}

func (s *GetDDoSProtectionConfigurationResponseData) SetArray(v []*DDoSProtectionDataArray) *GetDDoSProtectionConfigurationResponseData {
	s.Array = v
	return s
}

type DDoSProtectionDataArray struct {
	// {"en":"Domain.","zh_CN":"域名。"}
	Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
	// {"en":"Basic switch/mode information.","zh_CN":"基础开关/模式信息。"}
	DdosProtectSwitch *DdosProtectSwitchVO `json:"ddosProtectSwitch,omitempty" xml:"ddosProtectSwitch,omitempty" require:"true" type:"Struct"`
	// {"en":"Built-In rules, unprovided rules will take effect according to production configuration.","zh_CN":"内置规则，未提供的规则将按生产配置生效。"}
	BuiltInRules []*BuiltInRulesVO `json:"builtInRules,omitempty" xml:"builtInRules,omitempty" require:"true" type:"Repeated"`
}

func (s DDoSProtectionDataArray) String() string {
	return tea.Prettify(s)
}

func (s DDoSProtectionDataArray) GoString() string {
	return s.String()
}

func (s *DDoSProtectionDataArray) SetDomain(v string) *DDoSProtectionDataArray {
	s.Domain = &v
	return s
}

func (s *DDoSProtectionDataArray) SetDdosProtectSwitch(v *DdosProtectSwitchVO) *DDoSProtectionDataArray {
	s.DdosProtectSwitch = v
	return s
}

func (s *DDoSProtectionDataArray) SetBuiltInRules(v []*BuiltInRulesVO) *DDoSProtectionDataArray {
	s.BuiltInRules = v
	return s
}

type DdosProtectSwitchVO struct {
	// {"en":"3/4 layer protection switch.\nON: Enable.\nOFF: Disable.","zh_CN":"3/4层防护开关。\nON：开启。\nOFF：关闭。","exampleValue":"ON,OFF"}
	L4DdosSwitch *string `json:"l4DdosSwitch,omitempty" xml:"l4DdosSwitch,omitempty" require:"true"`
	// {"en":"Layer 7 HTTP DDoS protection switch.\nON: Enable.\nOFF: Disable.","zh_CN":"7层HTTP DDoS防护开关。\nON：开启。\nOFF：关闭。","exampleValue":"ON,OFF"}
	L7DdosSwitch *string `json:"l7DdosSwitch,omitempty" xml:"l7DdosSwitch,omitempty" require:"true"`
	// {"en":"Protection mode.\nLOOSE: loose.\nMODERATE: moderate.\nSTRICT: strict.","zh_CN":"防护模式。\nLOOSE：宽松。\nMODERATE：适中。\nSTRICT：严格。","exampleValue":"LOOSE,MODERATE,STRICT"}
	ProtectMode *string `json:"protectMode,omitempty" xml:"protectMode,omitempty" require:"true"`
	// {"en":"Built in protective switch.\nON: Enable.\nOFF: Disable.","zh_CN":"内置防护开关。\nON：开启。\nOFF：关闭。","exampleValue":"ON,OFF"}
	InnerSwitch *string `json:"innerSwitch,omitempty" xml:"innerSwitch,omitempty" require:"true"`
	// {"en":"AI intelligent protection switch.\nON: Enable.\nOFF: Disable.","zh_CN":"AI智能防护开关。\nON：开启。\nOFF：关闭。","exampleValue":"ON,OFF"}
	AiSwitch *string `json:"aiSwitch,omitempty" xml:"aiSwitch,omitempty" require:"true"`
	// {"en":"AI rule action.\nBLOCK: BLOCK.\nLOG: Monitor.\nRR: DDoS managed challenge.","zh_CN":"AI规则动作。\nBLOCK：拦截。\nLOG：监控。\nRR：DDoS托管。"}
	AiAction *string `json:"aiAction,omitempty" xml:"aiAction,omitempty" require:"true"`
}

func (s DdosProtectSwitchVO) String() string {
	return tea.Prettify(s)
}

func (s DdosProtectSwitchVO) GoString() string {
	return s.String()
}

func (s *DdosProtectSwitchVO) SetL4DdosSwitch(v string) *DdosProtectSwitchVO {
	s.L4DdosSwitch = &v
	return s
}

func (s *DdosProtectSwitchVO) SetL7DdosSwitch(v string) *DdosProtectSwitchVO {
	s.L7DdosSwitch = &v
	return s
}

func (s *DdosProtectSwitchVO) SetProtectMode(v string) *DdosProtectSwitchVO {
	s.ProtectMode = &v
	return s
}

func (s *DdosProtectSwitchVO) SetInnerSwitch(v string) *DdosProtectSwitchVO {
	s.InnerSwitch = &v
	return s
}

func (s *DdosProtectSwitchVO) SetAiSwitch(v string) *DdosProtectSwitchVO {
	s.AiSwitch = &v
	return s
}

func (s *DdosProtectSwitchVO) SetAiAction(v string) *DdosProtectSwitchVO {
	s.AiAction = &v
	return s
}

type BuiltInRulesVO struct {
	// {"en":"rule ID.","zh_CN":"规则ID。"}
	RuleId *string `json:"ruleId,omitempty" xml:"ruleId,omitempty" require:"true"`
	// {"en":"Security level.\nDEFAULT_ENABLE: default enabled.\nATTACK_ENABLE: enable during attack.\nBASE_CLOSE: basic off.\nCLOSE: permanently closed.","zh_CN":"安全级别。\nDEFAULT_ENABLE：默认启用。\nATTACK_ENABLE：攻击时启用。\nBASE_CLOSE：基本关闭。\nCLOSE：永久关闭。","exampleValue":"DEFAULT_ENABLE,ATTACK_ENABLE,BASE_CLOSE,CLOSE"}
	SecurityLevel *string `json:"securityLevel,omitempty" xml:"securityLevel,omitempty" require:"true"`
	// {"en":"Mode.\nBLOCK: Protect(Default).\nRR: Protect(Managed).\nLOG: Monitor.\nDENIED: Connection denied.","zh_CN":"模式。\nBLOCK：防护（拦截）。\nRR：防护（托管）。\nLOG：观察。\nDENIED: 拒绝连接。","exampleValue":"BLOCK,RR,LOG,DENIED"}
	Action *string `json:"action,omitempty" xml:"action,omitempty" require:"true"`
	// {"en":"Chinese Rule Name.","zh_CN":"中文规则名称。"}
	RuleNameCn *string `json:"ruleNameCn,omitempty" xml:"ruleNameCn,omitempty" require:"true"`
	// {"en":"English Rule Name.","zh_CN":"英文规则名称。"}
	RuleNameEn *string `json:"ruleNameEn,omitempty" xml:"ruleNameEn,omitempty" require:"true"`
}

func (s BuiltInRulesVO) String() string {
	return tea.Prettify(s)
}

func (s BuiltInRulesVO) GoString() string {
	return s.String()
}

func (s *BuiltInRulesVO) SetRuleId(v string) *BuiltInRulesVO {
	s.RuleId = &v
	return s
}

func (s *BuiltInRulesVO) SetSecurityLevel(v string) *BuiltInRulesVO {
	s.SecurityLevel = &v
	return s
}

func (s *BuiltInRulesVO) SetAction(v string) *BuiltInRulesVO {
	s.Action = &v
	return s
}

func (s *BuiltInRulesVO) SetRuleNameCn(v string) *BuiltInRulesVO {
	s.RuleNameCn = &v
	return s
}

func (s *BuiltInRulesVO) SetRuleNameEn(v string) *BuiltInRulesVO {
	s.RuleNameEn = &v
	return s
}
