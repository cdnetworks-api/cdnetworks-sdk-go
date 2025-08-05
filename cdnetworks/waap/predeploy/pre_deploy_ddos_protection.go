// This file is auto-generated, don't edit it. Thanks.
package predeploy

import (
	"github.com/alibabacloud-go/tea/tea"
)

type PreDeployDDoSProtectionConfigurationRequest struct {
	// {"en":"Domain.","zh_CN":"域名。"}
	Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
	// {"en":"Basic switch/mode information.","zh_CN":"基础开关/模式信息。"}
	DdosProtectSwitch *DdosProtectSwitch `json:"ddosProtectSwitch,omitempty" xml:"ddosProtectSwitch,omitempty" type:"Struct"`
	// {"en":"Built-In rules, unprovided rules will take effect according to current production configuration.","zh_CN":"内置规则，未提供的规则将按生产配置生效。"}
	BuiltInRules []*DDoSBuiltInRule `json:"builtInRules,omitempty" xml:"builtInRules,omitempty" type:"Repeated"`
}

func (s PreDeployDDoSProtectionConfigurationRequest) String() string {
	return tea.Prettify(s)
}

func (s PreDeployDDoSProtectionConfigurationRequest) GoString() string {
	return s.String()
}

func (s *PreDeployDDoSProtectionConfigurationRequest) SetDomain(v string) *PreDeployDDoSProtectionConfigurationRequest {
	s.Domain = &v
	return s
}

func (s *PreDeployDDoSProtectionConfigurationRequest) SetDdosProtectSwitch(v *DdosProtectSwitch) *PreDeployDDoSProtectionConfigurationRequest {
	s.DdosProtectSwitch = v
	return s
}

func (s *PreDeployDDoSProtectionConfigurationRequest) SetBuiltInRules(v []*DDoSBuiltInRule) *PreDeployDDoSProtectionConfigurationRequest {
	s.BuiltInRules = v
	return s
}

type DdosProtectSwitch struct {
	// {"en":"Layer 7 HTTP DDoS protection switch.\nON: Enable.\nOFF: Disable.","zh_CN":"7层HTTP DDoS防护开关。\nON：开启。\nOFF：关闭。","exampleValue":"ON,OFF"}
	L7DdosSwitch *string `json:"l7DdosSwitch,omitempty" xml:"l7DdosSwitch,omitempty" require:"true"`
	// {"en":"Protection mode.\nLOOSE: loose.\nMODERATE: moderate.\nSTRICT: strict.","zh_CN":"防护模式。\nLOOSE：宽松。\nMODERATE：适中。\nSTRICT：严格。","exampleValue":"LOOSE,MODERATE,STRICT"}
	ProtectMode *string `json:"protectMode,omitempty" xml:"protectMode,omitempty" require:"true"`
	// {"en":"Built in protective switch.\nON: Enable.\nOFF: Disable.","zh_CN":"内置防护开关。\nON：开启。\nOFF：关闭。","exampleValue":"ON,OFF"}
	InnerSwitch *string `json:"innerSwitch,omitempty" xml:"innerSwitch,omitempty" require:"true"`
}

func (s DdosProtectSwitch) String() string {
	return tea.Prettify(s)
}

func (s DdosProtectSwitch) GoString() string {
	return s.String()
}

func (s *DdosProtectSwitch) SetL7DdosSwitch(v string) *DdosProtectSwitch {
	s.L7DdosSwitch = &v
	return s
}

func (s *DdosProtectSwitch) SetProtectMode(v string) *DdosProtectSwitch {
	s.ProtectMode = &v
	return s
}

func (s *DdosProtectSwitch) SetInnerSwitch(v string) *DdosProtectSwitch {
	s.InnerSwitch = &v
	return s
}

type DDoSBuiltInRule struct {
	// {"en":"rule ID.","zh_CN":"规则ID。"}
	RuleId *string `json:"ruleId,omitempty" xml:"ruleId,omitempty" require:"true"`
	// {"en":"Security level.\nDEFAULT_ENABLE: default enabled.\nATTACK_ENABLE: enable during attack.\nBASE_CLOSE: basic off.\nCLOSE: permanently closed.","zh_CN":"安全级别。\nDEFAULT_ENABLE：默认启用。\nATTACK_ENABLE：攻击时启用。\nBASE_CLOSE：基本关闭。\nCLOSE：永久关闭。","exampleValue":"DEFAULT_ENABLE,ATTACK_ENABLE,BASE_CLOSE,CLOSE"}
	SecurityLevel *string `json:"securityLevel,omitempty" xml:"securityLevel,omitempty" require:"true"`
	// {"en":"Mode.\nBLOCK: Protect(Default).\nRR: Protect(Managed).\nLOG: Monitor.\nDENIED: Connection denied.","zh_CN":"模式。\nBLOCK：防护（拦截）。\nRR：防护（托管）。\nLOG：观察。\nDENIED: 拒绝连接。","exampleValue":"BLOCK,RR,LOG,nDENIED"}
	Action *string `json:"action,omitempty" xml:"action,omitempty" require:"true"`
}

func (s DDoSBuiltInRule) String() string {
	return tea.Prettify(s)
}

func (s DDoSBuiltInRule) GoString() string {
	return s.String()
}

func (s *DDoSBuiltInRule) SetRuleId(v string) *DDoSBuiltInRule {
	s.RuleId = &v
	return s
}

func (s *DDoSBuiltInRule) SetSecurityLevel(v string) *DDoSBuiltInRule {
	s.SecurityLevel = &v
	return s
}

func (s *DDoSBuiltInRule) SetAction(v string) *DDoSBuiltInRule {
	s.Action = &v
	return s
}

type PreDeployDDoSProtectionConfigurationResponse struct {
	// {"dictionary":"belong=WAAP-MS-Ext|dict=waap_retCodeEnum","en":"Please refer to the error code for exceptions.","zh_CN":"请参照错误码。"}
	Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
	// {"en":"Description.","zh_CN":"描述信息。"}
	Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
	// {"en":"Data.","zh_CN":"出参数据。"}
	Data *PreDeployDDoSProtectionConfigurationResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
}

func (s PreDeployDDoSProtectionConfigurationResponse) String() string {
	return tea.Prettify(s)
}

func (s PreDeployDDoSProtectionConfigurationResponse) GoString() string {
	return s.String()
}

func (s *PreDeployDDoSProtectionConfigurationResponse) SetCode(v string) *PreDeployDDoSProtectionConfigurationResponse {
	s.Code = &v
	return s
}

func (s *PreDeployDDoSProtectionConfigurationResponse) SetMsg(v string) *PreDeployDDoSProtectionConfigurationResponse {
	s.Msg = &v
	return s
}

func (s *PreDeployDDoSProtectionConfigurationResponse) SetData(v *PreDeployDDoSProtectionConfigurationResponseData) *PreDeployDDoSProtectionConfigurationResponse {
	s.Data = v
	return s
}

type PreDeployDDoSProtectionConfigurationResponseData struct {
	// {"en":"Pre-deployment id.","zh_CN":"预部署id。"}
	PreId *string `json:"preId,omitempty" xml:"preId,omitempty" require:"true"`
}

func (s PreDeployDDoSProtectionConfigurationResponseData) String() string {
	return tea.Prettify(s)
}

func (s PreDeployDDoSProtectionConfigurationResponseData) GoString() string {
	return s.String()
}

func (s *PreDeployDDoSProtectionConfigurationResponseData) SetPreId(v string) *PreDeployDDoSProtectionConfigurationResponseData {
	s.PreId = &v
	return s
}
