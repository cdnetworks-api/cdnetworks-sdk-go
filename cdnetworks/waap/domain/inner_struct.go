package domain

import "github.com/alibabacloud-go/tea/tea"

type WafConfig struct {
	// {"en":"Ruleset pattern.
	// MANUAL: Manual
	// AUTO: Automatic ", "zh_CN":"规则集模式。
	// MANUAL：手动
	// AUTO：自动"}
	RuleUpdateMode *string `json:"ruleUpdateMode,omitempty" xml:"ruleUpdateMode,omitempty" require:"true"`
	// {"en":"WAF protection switch.
	// ON: Enabled
	// OFF: Disabled", "zh_CN":"WAF防护开关。
	// ON：开启
	// OFF：关闭"}
	ConfigSwitch *string `json:"configSwitch,omitempty" xml:"configSwitch,omitempty" require:"true"`
	// {"en":"WAF protection Mode.
	// BLOCK: Interception
	// LOG: Observation", "zh_CN":"WAF防护模式。
	// BLOCK：拦截
	// LOG：观察"}
	DefendMode *string `json:"defendMode,omitempty" xml:"defendMode,omitempty" require:"true"`
}

func (s WafConfig) String() string {
	return tea.Prettify(s)
}

func (s WafConfig) GoString() string {
	return s.String()
}

func (s *WafConfig) SetRuleUpdateMode(v string) *WafConfig {
	s.RuleUpdateMode = &v
	return s
}

func (s *WafConfig) SetConfigSwitch(v string) *WafConfig {
	s.ConfigSwitch = &v
	return s
}

func (s *WafConfig) SetDefendMode(v string) *WafConfig {
	s.DefendMode = &v
	return s
}

type CustomizeRuleConfig struct {
	// {"en":"Custom rules switch.
	// ON: Enabled
	// OFF: Disabel", "zh_CN":"自定义规则开关。
	// ON：开启
	// OFF：关闭"}
	ConfigSwitch *string `json:"configSwitch,omitempty" xml:"configSwitch,omitempty" require:"true"`
}

func (s CustomizeRuleConfig) String() string {
	return tea.Prettify(s)
}

func (s CustomizeRuleConfig) GoString() string {
	return s.String()
}

func (s *CustomizeRuleConfig) SetConfigSwitch(v string) *CustomizeRuleConfig {
	s.ConfigSwitch = &v
	return s
}

type APIDefendConfig struct {
	// {"en":"API security switch.
	// ON: Enabled
	// OFF: Disabled", "zh_CN":"API安全开关。
	// ON：开启
	// OFF：关闭"}
	ConfigSwitch *string `json:"configSwitch,omitempty" xml:"configSwitch,omitempty" require:"true"`
}

func (s APIDefendConfig) String() string {
	return tea.Prettify(s)
}

func (s APIDefendConfig) GoString() string {
	return s.String()
}

func (s *APIDefendConfig) SetConfigSwitch(v string) *APIDefendConfig {
	s.ConfigSwitch = &v
	return s
}

type WhitelistConfig struct {
	// {"en":"Whitelist switch.
	// ON: Enabled
	// OFF: Disabled", "zh_CN":"白名单开关。
	// ON：开启
	// OFF：关闭"}
	ConfigSwitch *string `json:"configSwitch,omitempty" xml:"configSwitch,omitempty" require:"true"`
}

func (s WhitelistConfig) String() string {
	return tea.Prettify(s)
}

func (s WhitelistConfig) GoString() string {
	return s.String()
}

func (s *WhitelistConfig) SetConfigSwitch(v string) *WhitelistConfig {
	s.ConfigSwitch = &v
	return s
}

type BlockConfig struct {
	// {"en":"IP/Geo switch.
	// ON: Enabled
	// OFF: Disabled", "zh_CN":"IP/区域封禁开关。
	// ON：开启
	// OFF：关闭"}
	ConfigSwitch *string `json:"configSwitch,omitempty" xml:"configSwitch,omitempty" require:"true"`
}

func (s BlockConfig) String() string {
	return tea.Prettify(s)
}

func (s BlockConfig) GoString() string {
	return s.String()
}

func (s *BlockConfig) SetConfigSwitch(v string) *BlockConfig {
	s.ConfigSwitch = &v
	return s
}

type DMSConfig struct {
	// {"en":"DDoS protection switch.
	// ON: Enabled
	// OFF: Disabled", "zh_CN":"DDoS防护开关。
	// ON：开启
	// OFF：关闭"}
	ConfigSwitch *string `json:"configSwitch,omitempty" xml:"configSwitch,omitempty" require:"true"`
	// {"en":"DDoS protection mode.
	// LOOSE：Loose
	// MODERATE：Moderate
	// STRICT：Strict", "zh_CN":"DDoS防护模式。
	// LOOSE：宽松
	// MODERATE：适中
	// STRICT：严格"}
	ProtectionMode *string `json:"protectionMode,omitempty" xml:"protectionMode,omitempty" require:"true"`
	// {"en":"DDoS AI intelligent protection switch.
	// ON: Enabled
	// OFF: Disabled", "zh_CN":"DDoS AI智能防护开关。
	// ON：开启
	// OFF：关闭"}
	AiSwitch *string `json:"aiSwitch,omitempty" xml:"aiSwitch,omitempty" require:"true"`
}

func (s DMSConfig) String() string {
	return tea.Prettify(s)
}

func (s DMSConfig) GoString() string {
	return s.String()
}

func (s *DMSConfig) SetConfigSwitch(v string) *DMSConfig {
	s.ConfigSwitch = &v
	return s
}

func (s *DMSConfig) SetProtectionMode(v string) *DMSConfig {
	s.ProtectionMode = &v
	return s
}

func (s *DMSConfig) SetAiSwitch(v string) *DMSConfig {
	s.AiSwitch = &v
	return s
}

type IntelligenceConfig struct {
	// {"en":"Attack risk type action.", "zh_CN":"攻击风险类型处理动作。"}
	InfoCateAct *AttackRiskType `json:"infoCateAct,omitempty" xml:"infoCateAct,omitempty" require:"true"`
	// {"en":"Threat intelligence switch.
	// ON: Enabled
	// OFF: Disabled", "zh_CN":"威胁情报开关。
	// ON：开启
	// OFF：关闭"}
	ConfigSwitch *string `json:"configSwitch,omitempty" xml:"configSwitch,omitempty" require:"true"`
}

func (s IntelligenceConfig) String() string {
	return tea.Prettify(s)
}

func (s IntelligenceConfig) GoString() string {
	return s.String()
}

func (s *IntelligenceConfig) SetInfoCateAct(v *AttackRiskType) *IntelligenceConfig {
	s.InfoCateAct = v
	return s
}

func (s *IntelligenceConfig) SetConfigSwitch(v string) *IntelligenceConfig {
	s.ConfigSwitch = &v
	return s
}

type AttackRiskType struct {
	// {"en":"Attack resource risk action.
	// NO_USE: Not used
	// BLOCK: Deny
	// LOG: Log", "zh_CN":"攻击资源风险处理动作。
	// NO_USE：不使用
	// BLOCK：拦截
	// LOG：监控"}
	AttackSource *string `json:"attackSource,omitempty" xml:"attackSource,omitempty" require:"true"`
	// {"en":"Specific attack risk action.
	// NO_USE: Not used
	// BLOCK: Deny
	// LOG:Log", "zh_CN":"特定攻击风险处理动作。
	// NO_USE：不使用
	// BLOCK：拦截
	// LOG：监控"}
	SpecAttack *string `json:"specAttack,omitempty" xml:"specAttack,omitempty" require:"true"`
	// {"en":"Industry attack risk action.
	// NO_USE: Not used
	// BLOCK: Deny
	// LOG: Log", "zh_CN":"重点行业风险处理动作。
	// NO_USE：不使用
	// BLOCK：拦截
	// LOG：监控"}
	Industry *string `json:"industry,omitempty" xml:"industry,omitempty" require:"true"`
}

func (s AttackRiskType) String() string {
	return tea.Prettify(s)
}

func (s AttackRiskType) GoString() string {
	return s.String()
}

func (s *AttackRiskType) SetAttackSource(v string) *AttackRiskType {
	s.AttackSource = &v
	return s
}

func (s *AttackRiskType) SetSpecAttack(v string) *AttackRiskType {
	s.SpecAttack = &v
	return s
}

func (s *AttackRiskType) SetIndustry(v string) *AttackRiskType {
	s.Industry = &v
	return s
}

type BOTConfig struct {
	// {"en":"Bot management switch.
	//ON:Enabled
	//OFF:Disabled","zh_CN":"Bot管理开关。
	//ON：开启
	//OFF：关闭","exampleValue":"ON,OFF"}
	ConfigSwitch *string `json:"configSwitch,omitempty" xml:"configSwitch,omitempty" require:"true"`
	// {"en":"AI Bots action.
	//NO_USE: Not used
	//BLOCK: Deny
	//LOG: Log","zh_CN":"AI Bots处理动作。
	//NO_USE：不使用
	//BLOCK：拦截
	//LOG：监控","exampleValue":"NO_USE,LOG,BLOCK"}
	AiBotsAct *string `json:"aiBotsAct,omitempty" xml:"aiBotsAct,omitempty" require:"true"`
	// {"en":"Public Bots action.
	//NO_USE: not used
	//BLOCK: Deny
	//LOG: Log
	//ACCEPT: Skip","zh_CN":"公开Bots处理动作。
	//NO_USE：不使用
	//BLOCK：拦截
	//LOG：监控
	//ACCEPT：放行","exampleValue":"NO_USE,LOG,BLOCK,ACCEPT"}
	PublicBotsAct *string `json:"publicBotsAct,omitempty" xml:"publicBotsAct,omitempty" require:"true"`
	// {"en":"Definite Bots action.
	//NO_USE: Not used
	//BLOCK: Deny
	//LOG: Log","zh_CN":"绝对Bots处理动作。
	//NO_USE：不使用
	//BLOCK：拦截
	//LOG：监控","exampleValue":"NO_USE,LOG,BLOCK"}
	AbsoluteBotsAct *string `json:"absoluteBotsAct,omitempty" xml:"absoluteBotsAct,omitempty" require:"true"`
	// {"en":"Browser Bot defense.","zh_CN":"Web风险检测。"}
	WebRiskConfig *BotWebConfigDefault `json:"webRiskConfig,omitempty" xml:"webRiskConfig,omitempty" require:"true" type:"Struct"`
}

func (s BOTConfig) String() string {
	return tea.Prettify(s)
}

func (s BOTConfig) GoString() string {
	return s.String()
}

func (s *BOTConfig) SetPublicBotsAct(v string) *BOTConfig {
	s.PublicBotsAct = &v
	return s
}

func (s *BOTConfig) SetConfigSwitch(v string) *BOTConfig {
	s.ConfigSwitch = &v
	return s
}

func (s *BOTConfig) SetAiBotsAct(v string) *BOTConfig {
	s.AiBotsAct = &v
	return s
}

func (s *BOTConfig) SetAbsoluteBotsAct(v string) *BOTConfig {
	s.AbsoluteBotsAct = &v
	return s
}

func (s *BOTConfig) SetWebRiskConfig(v *BotWebConfigDefault) *BOTConfig {
	s.WebRiskConfig = v
	return s
}

type BotWebConfigDefault struct {
	// {"en":"Action.
	// NO_USE: Not used
	// BLOCK: Deny
	// LOG: Log", "zh_CN":"处理动作。
	// NO_USE:不使用
	// BLOCK:拦截
	// LOG:监控"}
	Act *string `json:"act,omitempty" xml:"act,omitempty" require:"true"`
}

func (s BotWebConfigDefault) String() string {
	return tea.Prettify(s)
}

func (s BotWebConfigDefault) GoString() string {
	return s.String()
}

func (s *BotWebConfigDefault) SetAct(v string) *BotWebConfigDefault {
	s.Act = &v
	return s
}

type RateLimitConfig struct {
	// {"en":"Rate limiting switch.
	// ON: Enabled
	// OFF: Disabled", "zh_CN":"频率限制开关。
	// ON：开启
	// OFF：关闭"}
	ConfigSwitch *string `json:"configSwitch,omitempty" xml:"configSwitch,omitempty" require:"true"`
}

func (s RateLimitConfig) String() string {
	return tea.Prettify(s)
}

func (s RateLimitConfig) GoString() string {
	return s.String()
}

func (s *RateLimitConfig) SetConfigSwitch(v string) *RateLimitConfig {
	s.ConfigSwitch = &v
	return s
}
