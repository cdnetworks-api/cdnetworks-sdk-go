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
	Array []*WAFDataArray `json:"array,omitempty" xml:"array,omitempty" require:"true" type:"Repeated"`
}

func (s GetWafConfigurationResponseData) String() string {
	return tea.Prettify(s)
}

func (s GetWafConfigurationResponseData) GoString() string {
	return s.String()
}

func (s *GetWafConfigurationResponseData) SetArray(v []*WAFDataArray) *GetWafConfigurationResponseData {
	s.Array = v
	return s
}

type WAFDataArray struct {
	// {"en":"Domain.","zh_CN":"域名。"}
	Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
	// {"en":"Policy switch.ON: Enable.OFF: Disable.","zh_CN":"策略开关。\nON：开启。\nOFF：关闭。","exampleValue":"ON,OFF"}
	ConfigSwitch *string `json:"configSwitch,omitempty" xml:"configSwitch,omitempty" require:"true"`
	// {"en":"Basic configuration.","zh_CN":"基础配置。"}
	ConfBasic *ConfBasicVO `json:"confBasic,omitempty" xml:"confBasic,omitempty" require:"true" type:"Struct"`
	// {"en":"Rule list.","zh_CN":"规则列表。"}
	RuleList []*WAFRuleVO `json:"ruleList,omitempty" xml:"ruleList,omitempty" require:"true" type:"Repeated"`
}

func (s WAFDataArray) String() string {
	return tea.Prettify(s)
}

func (s WAFDataArray) GoString() string {
	return s.String()
}

func (s *WAFDataArray) SetDomain(v string) *WAFDataArray {
	s.Domain = &v
	return s
}

func (s *WAFDataArray) SetConfigSwitch(v string) *WAFDataArray {
	s.ConfigSwitch = &v
	return s
}

func (s *WAFDataArray) SetConfBasic(v *ConfBasicVO) *WAFDataArray {
	s.ConfBasic = v
	return s
}

func (s *WAFDataArray) SetRuleList(v []*WAFRuleVO) *WAFDataArray {
	s.RuleList = v
	return s
}

type ConfBasicVO struct {
	// {"en":"Protection Mode.\nBLOCK: Block the attack request directly.\nLOG: Only log the attack request without blocking it.","zh_CN":"防护模式。\nBLOCK：直接拦截攻击请求。\nLOG：记录日志，不拦截攻击请求。","exampleValue":"BLOCK,LOG"}
	DefendMode *string `json:"defendMode,omitempty" xml:"defendMode,omitempty" require:"true"`
	// {"en":"Ruleset Mode.\nMANUAL:  Check Ruleset update and all Recommendations on the Console, decide to apply them or not, all of these must be done by yourself manually.\nAUTO: Automatically upgrade the Ruleset to the latest version and apply the Recommendations learned from your website traffic to Exception, which can keep your website with high-level security anytime.","zh_CN":"规则集模式。\nMANUAL：规则集有更新，或系统自动学习网站流量生成规则例外建议时，需自主评估是否更新配置。\nAUTO：规则集有更新，或系统自动学习网站流量生成规则例外建议时，自动更新配置。","exampleValue":"MANUAL,AUTO"}
	RuleUpdateMode *string `json:"ruleUpdateMode,omitempty" xml:"ruleUpdateMode,omitempty" require:"true"`
}

func (s ConfBasicVO) String() string {
	return tea.Prettify(s)
}

func (s ConfBasicVO) GoString() string {
	return s.String()
}

func (s *ConfBasicVO) SetDefendMode(v string) *ConfBasicVO {
	s.DefendMode = &v
	return s
}

func (s *ConfBasicVO) SetRuleUpdateMode(v string) *ConfBasicVO {
	s.RuleUpdateMode = &v
	return s
}

type WAFRuleVO struct {
	// {"en":"WAF rule ID.","zh_CN":"WAF规则ID。"}
	RuleId *int    `json:"ruleId,omitempty" xml:"ruleId,omitempty" require:"true"`
	Name   *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
	// {"en":"Rule actions.\nBLOCK: Deny request by a default 403 response.\nLOG: Log request and continue further detections.\nOFF: Select if you do not a policy take effect.","zh_CN":"规则动作。\nBLOCK：阻断请求并响应403。\nLOG：记录请求的攻击行为，并继续做进一步的评估。\nOFF：对应规则或策略不生效。","exampleValue":"BLOCK,LOG,OFF"}
	Mode *string `json:"mode,omitempty" xml:"mode,omitempty" require:"true"`
	// {"en":"Rule exceptions.","zh_CN":"规则例外。"}
	ExceptionList []*GetWafConfigurationResponseDataArrayRuleListExceptionList `json:"exceptionList,omitempty" xml:"exceptionList,omitempty" require:"true" type:"Repeated"`
}

func (s WAFRuleVO) String() string {
	return tea.Prettify(s)
}

func (s WAFRuleVO) GoString() string {
	return s.String()
}

func (s *WAFRuleVO) SetRuleId(v int) *WAFRuleVO {
	s.RuleId = &v
	return s
}

func (s *WAFRuleVO) SetMode(v string) *WAFRuleVO {
	s.Mode = &v
	return s
}

func (s *WAFRuleVO) SetName(v string) *WAFRuleVO {
	s.Name = &v
	return s
}

func (s *WAFRuleVO) SetExceptionList(v []*GetWafConfigurationResponseDataArrayRuleListExceptionList) *WAFRuleVO {
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
