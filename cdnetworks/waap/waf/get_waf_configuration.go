// This file is auto-generated, don't edit it. Thanks.
package waf

import (
	"github.com/alibabacloud-go/tea/tea"
)

type GetWafConfigurationRequest struct {
	// {"en":"Domain list.","zh_CN":"域名列表。"}
	DomainList []*string `json:"domainList,omitempty" xml:"domainList,omitempty" require:"true" type:"Repeated"`
}

func (s GetWafConfigurationRequest) String() string {
	return tea.Prettify(s)
}

func (s GetWafConfigurationRequest) GoString() string {
	return s.String()
}

func (s *GetWafConfigurationRequest) SetDomainList(v []*string) *GetWafConfigurationRequest {
	s.DomainList = v
	return s
}

type GetWafConfigurationResponse struct {
	// {"dictionary":"belong=WAAP-MS-Ext|dict=waap_retCodeEnum","en":"Please refer to the error code for exceptions.","zh_CN":"请参照错误码。"}
	Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
	// {"en":"Description.","zh_CN":"描述信息。"}
	Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
	// {"en":"Data.","zh_CN":"出参数据。"}
	Data *GetWafConfigurationResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
}

func (s GetWafConfigurationResponse) String() string {
	return tea.Prettify(s)
}

func (s GetWafConfigurationResponse) GoString() string {
	return s.String()
}

func (s *GetWafConfigurationResponse) SetCode(v string) *GetWafConfigurationResponse {
	s.Code = &v
	return s
}

func (s *GetWafConfigurationResponse) SetMsg(v string) *GetWafConfigurationResponse {
	s.Msg = &v
	return s
}

func (s *GetWafConfigurationResponse) SetData(v *GetWafConfigurationResponseData) *GetWafConfigurationResponse {
	s.Data = v
	return s
}

type GetWafConfigurationResponseData struct {
	// {"en":"Array.","zh_CN":"数组。"}
	Array []*GetWafConfigurationResponseDataArray `json:"array,omitempty" xml:"array,omitempty" require:"true" type:"Repeated"`
}

func (s GetWafConfigurationResponseData) String() string {
	return tea.Prettify(s)
}

func (s GetWafConfigurationResponseData) GoString() string {
	return s.String()
}

func (s *GetWafConfigurationResponseData) SetArray(v []*GetWafConfigurationResponseDataArray) *GetWafConfigurationResponseData {
	s.Array = v
	return s
}

type GetWafConfigurationResponseDataArray struct {
	// {"en":"Domain.","zh_CN":"域名。"}
	Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
	// {"en":"Policy switch.ON: Enable.OFF: Disable.","zh_CN":"策略开关。\nON：开启。\nOFF：关闭。","exampleValue":"ON,OFF"}
	ConfigSwitch *string `json:"configSwitch,omitempty" xml:"configSwitch,omitempty" require:"true"`
	// {"en":"Basic configuration.","zh_CN":"基础配置。"}
	ConfBasic *GetWafConfigurationResponseDataArrayConfBasic `json:"confBasic,omitempty" xml:"confBasic,omitempty" require:"true" type:"Struct"`
	// {"en":"Rule list.","zh_CN":"规则列表。"}
	RuleList []*GetWafConfigurationResponseDataArrayRuleList `json:"ruleList,omitempty" xml:"ruleList,omitempty" require:"true" type:"Repeated"`
	// {"en":"Scan protection configuration.","zh_CN":"扫面防护配置。"}
	ScanProtection *GetWafConfigurationResponseDataArrayScanProtection `json:"scanProtection,omitempty" xml:"scanProtection,omitempty" require:"true" type:"Struct"`
}

func (s GetWafConfigurationResponseDataArray) String() string {
	return tea.Prettify(s)
}

func (s GetWafConfigurationResponseDataArray) GoString() string {
	return s.String()
}

func (s *GetWafConfigurationResponseDataArray) SetDomain(v string) *GetWafConfigurationResponseDataArray {
	s.Domain = &v
	return s
}

func (s *GetWafConfigurationResponseDataArray) SetConfigSwitch(v string) *GetWafConfigurationResponseDataArray {
	s.ConfigSwitch = &v
	return s
}

func (s *GetWafConfigurationResponseDataArray) SetConfBasic(v *GetWafConfigurationResponseDataArrayConfBasic) *GetWafConfigurationResponseDataArray {
	s.ConfBasic = v
	return s
}

func (s *GetWafConfigurationResponseDataArray) SetRuleList(v []*GetWafConfigurationResponseDataArrayRuleList) *GetWafConfigurationResponseDataArray {
	s.RuleList = v
	return s
}

func (s *GetWafConfigurationResponseDataArray) SetScanProtection(v *GetWafConfigurationResponseDataArrayScanProtection) *GetWafConfigurationResponseDataArray {
	s.ScanProtection = v
	return s
}

type GetWafConfigurationResponseDataArrayConfBasic struct {
	// {"en":"Protection Mode.\nBLOCK: Block the attack request directly.\nLOG: Only log the attack request without blocking it.","zh_CN":"防护模式。\nBLOCK：直接拦截攻击请求。\nLOG：记录日志，不拦截攻击请求。","exampleValue":"BLOCK,LOG"}
	DefendMode *string `json:"defendMode,omitempty" xml:"defendMode,omitempty" require:"true"`
	// {"en":"Ruleset Mode.\nMANUAL:  Check Ruleset update and all Recommendations on the Console, decide to apply them or not, all of these must be done by yourself manually.\nAUTO: Automatically upgrade the Ruleset to the latest version and apply the Recommendations learned from your website traffic to Exception, which can keep your website with high-level security anytime.","zh_CN":"规则集模式。\nMANUAL：规则集有更新，或系统自动学习网站流量生成规则例外建议时，需自主评估是否更新配置。\nAUTO：规则集有更新，或系统自动学习网站流量生成规则例外建议时，自动更新配置。","exampleValue":"MANUAL,AUTO"}
	RuleUpdateMode *string `json:"ruleUpdateMode,omitempty" xml:"ruleUpdateMode,omitempty" require:"true"`
}

func (s GetWafConfigurationResponseDataArrayConfBasic) String() string {
	return tea.Prettify(s)
}

func (s GetWafConfigurationResponseDataArrayConfBasic) GoString() string {
	return s.String()
}

func (s *GetWafConfigurationResponseDataArrayConfBasic) SetDefendMode(v string) *GetWafConfigurationResponseDataArrayConfBasic {
	s.DefendMode = &v
	return s
}

func (s *GetWafConfigurationResponseDataArrayConfBasic) SetRuleUpdateMode(v string) *GetWafConfigurationResponseDataArrayConfBasic {
	s.RuleUpdateMode = &v
	return s
}

type GetWafConfigurationResponseDataArrayRuleList struct {
	// {"en":"WAF rule ID.","zh_CN":"WAF规则ID。"}
	RuleId *int `json:"ruleId,omitempty" xml:"ruleId,omitempty" require:"true"`
	// {"en":"Rule name.","zh_CN":"规则名称。"}
	Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
	// {"en":"Rule actions.\nBLOCK: Deny request by a default 403 response.\nLOG: Log request and continue further detections.\nOFF: Select if you do not a policy take effect.","zh_CN":"规则动作。\nBLOCK：阻断请求并响应403。\nLOG：记录请求的攻击行为，并继续做进一步的评估。\nOFF：对应规则或策略不生效。","exampleValue":"BLOCK,LOG,OFF"}
	Mode *string `json:"mode,omitempty" xml:"mode,omitempty" require:"true"`
	// {"en":"Rule exceptions.","zh_CN":"规则例外。"}
	ExceptionList []*GetWafConfigurationResponseDataArrayRuleListExceptionList `json:"exceptionList,omitempty" xml:"exceptionList,omitempty" require:"true" type:"Repeated"`
}

func (s GetWafConfigurationResponseDataArrayRuleList) String() string {
	return tea.Prettify(s)
}

func (s GetWafConfigurationResponseDataArrayRuleList) GoString() string {
	return s.String()
}

func (s *GetWafConfigurationResponseDataArrayRuleList) SetRuleId(v int) *GetWafConfigurationResponseDataArrayRuleList {
	s.RuleId = &v
	return s
}

func (s *GetWafConfigurationResponseDataArrayRuleList) SetName(v string) *GetWafConfigurationResponseDataArrayRuleList {
	s.Name = &v
	return s
}

func (s *GetWafConfigurationResponseDataArrayRuleList) SetMode(v string) *GetWafConfigurationResponseDataArrayRuleList {
	s.Mode = &v
	return s
}

func (s *GetWafConfigurationResponseDataArrayRuleList) SetExceptionList(v []*GetWafConfigurationResponseDataArrayRuleListExceptionList) *GetWafConfigurationResponseDataArrayRuleList {
	s.ExceptionList = v
	return s
}

type GetWafConfigurationResponseDataArrayRuleListExceptionList struct {
	// {"en":"Matching conditions.\nip: IP\npath: Path\nuri: URI\nurlParamName: URI Parameter Name\nurlParamValue: URI Parameter Value\nuserAgent: User Agent\nhttpHeaderName: Request Header Name\nhttpHeaderValue: Request Header Value\ncookie: Cookie\nbody: Body\nbodyParamName: Body Parameter Name\nbodyParamValue: Body Parameter Value","zh_CN":"匹配条件。\nip：IP\npath：路径\nuri：URI\nurlParamName：URI参数名\nurlParamValue：URI参数值\nuserAgent：User Agent\nhttpHeaderName：请求头部名称\nhttpHeaderValue：请求头部值\ncookie：Cookie\nbody：Body\nbodyParamName：Body参数名\nbodyParamValue：Body参数值","exampleValue":"ip,path,uri,urlParamName,urlParamValue,userAgent,httpHeaderName,httpHeaderValue,cookie,body,bodyParamName,bodyParamValue"}
	Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
	// {"en":"Match type,IP can only be EQUAL.\nEQUAL: Equal\nCONTAIN: Contains\nREGEX: Regular match","zh_CN":"匹配类型，IP只能是等于。\nEQUAL：等于\nCONTAIN：包含\nREGEX：正则匹配","exampleValue":"EQUAL,CONTAIN,REGEX"}
	MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
	// {"en":"Rule exceptions.\nWhen matchType=EQUAL, case-sensitive, path and uri must start with \"/\", and body can only pass one value;\nWhen matchType=REGEX, only one value can be passed.","zh_CN":"规则例外内容。\nmatchType=EQUAL时，大小写敏感，path和uri必须以\"/\"开头，body只能传一个值；\nmatchType=REGEX时，只能传一个值。"}
	ContentList []*string `json:"contentList,omitempty" xml:"contentList,omitempty" require:"true" type:"Repeated"`
}

func (s GetWafConfigurationResponseDataArrayRuleListExceptionList) String() string {
	return tea.Prettify(s)
}

func (s GetWafConfigurationResponseDataArrayRuleListExceptionList) GoString() string {
	return s.String()
}

func (s *GetWafConfigurationResponseDataArrayRuleListExceptionList) SetType(v string) *GetWafConfigurationResponseDataArrayRuleListExceptionList {
	s.Type = &v
	return s
}

func (s *GetWafConfigurationResponseDataArrayRuleListExceptionList) SetMatchType(v string) *GetWafConfigurationResponseDataArrayRuleListExceptionList {
	s.MatchType = &v
	return s
}

func (s *GetWafConfigurationResponseDataArrayRuleListExceptionList) SetContentList(v []*string) *GetWafConfigurationResponseDataArrayRuleListExceptionList {
	s.ContentList = v
	return s
}

type GetWafConfigurationResponseDataArrayScanProtection struct {
	// {"en":"Scanning tool detection configuration.","zh_CN":"扫描工具检测配置。"}
	ScanToolsConfig *GetWafConfigurationResponseDataArrayScanProtectionScanToolsConfig `json:"scanToolsConfig,omitempty" xml:"scanToolsConfig,omitempty" require:"true" type:"Struct"`
	// {"en":"Repeated violation detection configuration.","zh_CN":"反复违规检测配置。"}
	RepeatedViolationConfig *GetWafConfigurationResponseDataArrayScanProtectionRepeatedViolationConfig `json:"repeatedViolationConfig,omitempty" xml:"repeatedViolationConfig,omitempty" require:"true" type:"Struct"`
	// {"en":"Directory probing detection configuration.","zh_CN":"目录嗅探检测配置。"}
	DirectoryProbingConfig *GetWafConfigurationResponseDataArrayScanProtectionDirectoryProbingConfig `json:"directoryProbingConfig,omitempty" xml:"directoryProbingConfig,omitempty" require:"true" type:"Struct"`
}

func (s GetWafConfigurationResponseDataArrayScanProtection) String() string {
	return tea.Prettify(s)
}

func (s GetWafConfigurationResponseDataArrayScanProtection) GoString() string {
	return s.String()
}

func (s *GetWafConfigurationResponseDataArrayScanProtection) SetScanToolsConfig(v *GetWafConfigurationResponseDataArrayScanProtectionScanToolsConfig) *GetWafConfigurationResponseDataArrayScanProtection {
	s.ScanToolsConfig = v
	return s
}

func (s *GetWafConfigurationResponseDataArrayScanProtection) SetRepeatedViolationConfig(v *GetWafConfigurationResponseDataArrayScanProtectionRepeatedViolationConfig) *GetWafConfigurationResponseDataArrayScanProtection {
	s.RepeatedViolationConfig = v
	return s
}

func (s *GetWafConfigurationResponseDataArrayScanProtection) SetDirectoryProbingConfig(v *GetWafConfigurationResponseDataArrayScanProtectionDirectoryProbingConfig) *GetWafConfigurationResponseDataArrayScanProtection {
	s.DirectoryProbingConfig = v
	return s
}

type GetWafConfigurationResponseDataArrayScanProtectionScanToolsConfig struct {
	// {"en":"Action.\nNO_USE:Not Used.\nLOG:Log.\nBLOCK:Deny.","zh_CN":"处理动作。\nNO_USE：不使用。\nLOG：监控。\nBLOCK：拦截。","exampleValue":"NO_USE,LOG,BLOCK"}
	Action *string `json:"action,omitempty" xml:"action,omitempty" require:"true"`
}

func (s GetWafConfigurationResponseDataArrayScanProtectionScanToolsConfig) String() string {
	return tea.Prettify(s)
}

func (s GetWafConfigurationResponseDataArrayScanProtectionScanToolsConfig) GoString() string {
	return s.String()
}

func (s *GetWafConfigurationResponseDataArrayScanProtectionScanToolsConfig) SetAction(v string) *GetWafConfigurationResponseDataArrayScanProtectionScanToolsConfig {
	s.Action = &v
	return s
}

type GetWafConfigurationResponseDataArrayScanProtectionRepeatedViolationConfig struct {
	// {"en":"Action.\nNO_USE:Not Used.\nLOG:Log.\nBLOCK:Deny.","zh_CN":"处理动作。\nNO_USE：不使用。\nLOG：监控。\nBLOCK：拦截。","exampleValue":"NO_USE,LOG,BLOCK"}
	Action *string `json:"action,omitempty" xml:"action,omitempty" require:"true"`
	// {"en":"Statistical subject. \nIP: IP. \nIP_JA3: IP and JA3 fingerprint.","zh_CN":"统计对象。\nIP：IP。\nIP_JA3：IP+JA3指纹。","exampleValue":"IP,IP_JA3"}
	Target *string `json:"target,omitempty" xml:"target,omitempty" require:"true"`
	// {"en":"Time range, in seconds.","zh_CN":"时间范围，单位秒。"}
	Period *int `json:"period,omitempty" xml:"period,omitempty" require:"true"`
	// {"en":"Number of WAF built-in rule triggers.","zh_CN":"WAF内置规则触发种类数。"}
	WafRuleTypeCount *int `json:"wafRuleTypeCount,omitempty" xml:"wafRuleTypeCount,omitempty" require:"true"`
	// {"en":"Number of block actions.","zh_CN":"拦截次数"}
	BlockCount *int `json:"blockCount,omitempty" xml:"blockCount,omitempty" require:"true"`
	// {"en":"Handling action duration, in seconds.","zh_CN":"处理动作持续时间，单位秒。"}
	Duration *int `json:"duration,omitempty" xml:"duration,omitempty" require:"true"`
}

func (s GetWafConfigurationResponseDataArrayScanProtectionRepeatedViolationConfig) String() string {
	return tea.Prettify(s)
}

func (s GetWafConfigurationResponseDataArrayScanProtectionRepeatedViolationConfig) GoString() string {
	return s.String()
}

func (s *GetWafConfigurationResponseDataArrayScanProtectionRepeatedViolationConfig) SetAction(v string) *GetWafConfigurationResponseDataArrayScanProtectionRepeatedViolationConfig {
	s.Action = &v
	return s
}

func (s *GetWafConfigurationResponseDataArrayScanProtectionRepeatedViolationConfig) SetTarget(v string) *GetWafConfigurationResponseDataArrayScanProtectionRepeatedViolationConfig {
	s.Target = &v
	return s
}

func (s *GetWafConfigurationResponseDataArrayScanProtectionRepeatedViolationConfig) SetPeriod(v int) *GetWafConfigurationResponseDataArrayScanProtectionRepeatedViolationConfig {
	s.Period = &v
	return s
}

func (s *GetWafConfigurationResponseDataArrayScanProtectionRepeatedViolationConfig) SetWafRuleTypeCount(v int) *GetWafConfigurationResponseDataArrayScanProtectionRepeatedViolationConfig {
	s.WafRuleTypeCount = &v
	return s
}

func (s *GetWafConfigurationResponseDataArrayScanProtectionRepeatedViolationConfig) SetBlockCount(v int) *GetWafConfigurationResponseDataArrayScanProtectionRepeatedViolationConfig {
	s.BlockCount = &v
	return s
}

func (s *GetWafConfigurationResponseDataArrayScanProtectionRepeatedViolationConfig) SetDuration(v int) *GetWafConfigurationResponseDataArrayScanProtectionRepeatedViolationConfig {
	s.Duration = &v
	return s
}

type GetWafConfigurationResponseDataArrayScanProtectionDirectoryProbingConfig struct {
	// {"en":"Action.\nNO_USE:Not Used.\nLOG:Log.\nBLOCK:Deny.","zh_CN":"处理动作。\nNO_USE：不使用。\nLOG：监控。\nBLOCK：拦截。","exampleValue":"NO_USE,LOG,BLOCK"}
	Action *string `json:"action,omitempty" xml:"action,omitempty" require:"true"`
	// {"en":"Statistical subject. \nIP: IP. \nIP_JA3: IP and JA3 fingerprint.","zh_CN":"统计对象。\nIP：IP。\nIP_JA3：IP+JA3指纹。","exampleValue":"IP,IP_JA3"}
	Target *string `json:"target,omitempty" xml:"target,omitempty" require:"true"`
	// {"en":"Time range, in seconds.","zh_CN":"时间范围，单位秒。"}
	Period *int `json:"period,omitempty" xml:"period,omitempty" require:"true"`
	// {"en":"Number of requests.","zh_CN":"请求次数。"}
	RequestCountThreshold *int `json:"requestCountThreshold,omitempty" xml:"requestCountThreshold,omitempty" require:"true"`
	// {"en":"Number of non-existent directory requests.","zh_CN":"请求不存在的目录数量。"}
	NonExistentDirectoryThreshold *int `json:"nonExistentDirectoryThreshold,omitempty" xml:"nonExistentDirectoryThreshold,omitempty" require:"true"`
	// {"en":"Proportion of 404 status codes.","zh_CN":"404状态码占比。"}
	Rate404Threshold *int `json:"rate404Threshold,omitempty" xml:"rate404Threshold,omitempty" require:"true"`
	// {"en":"Handling action duration, in seconds.","zh_CN":"处理动作持续时间，单位秒。"}
	Duration *int `json:"duration,omitempty" xml:"duration,omitempty" require:"true"`
}

func (s GetWafConfigurationResponseDataArrayScanProtectionDirectoryProbingConfig) String() string {
	return tea.Prettify(s)
}

func (s GetWafConfigurationResponseDataArrayScanProtectionDirectoryProbingConfig) GoString() string {
	return s.String()
}

func (s *GetWafConfigurationResponseDataArrayScanProtectionDirectoryProbingConfig) SetAction(v string) *GetWafConfigurationResponseDataArrayScanProtectionDirectoryProbingConfig {
	s.Action = &v
	return s
}

func (s *GetWafConfigurationResponseDataArrayScanProtectionDirectoryProbingConfig) SetTarget(v string) *GetWafConfigurationResponseDataArrayScanProtectionDirectoryProbingConfig {
	s.Target = &v
	return s
}

func (s *GetWafConfigurationResponseDataArrayScanProtectionDirectoryProbingConfig) SetPeriod(v int) *GetWafConfigurationResponseDataArrayScanProtectionDirectoryProbingConfig {
	s.Period = &v
	return s
}

func (s *GetWafConfigurationResponseDataArrayScanProtectionDirectoryProbingConfig) SetRequestCountThreshold(v int) *GetWafConfigurationResponseDataArrayScanProtectionDirectoryProbingConfig {
	s.RequestCountThreshold = &v
	return s
}

func (s *GetWafConfigurationResponseDataArrayScanProtectionDirectoryProbingConfig) SetNonExistentDirectoryThreshold(v int) *GetWafConfigurationResponseDataArrayScanProtectionDirectoryProbingConfig {
	s.NonExistentDirectoryThreshold = &v
	return s
}

func (s *GetWafConfigurationResponseDataArrayScanProtectionDirectoryProbingConfig) SetRate404Threshold(v int) *GetWafConfigurationResponseDataArrayScanProtectionDirectoryProbingConfig {
	s.Rate404Threshold = &v
	return s
}

func (s *GetWafConfigurationResponseDataArrayScanProtectionDirectoryProbingConfig) SetDuration(v int) *GetWafConfigurationResponseDataArrayScanProtectionDirectoryProbingConfig {
	s.Duration = &v
	return s
}
