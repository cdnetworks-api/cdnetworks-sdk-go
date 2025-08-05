// This file is auto-generated, don't edit it. Thanks.
package predeploy

import (
	"github.com/alibabacloud-go/tea/tea"
)

type PreDeployRateLimitingConfigurationRequest struct {
	// {"en":"Domain list.","zh_CN":"域名列表。"}
	Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
	// {"en":"Policy switch.ON: Enable.OFF: Disable.","zh_CN":"策略开关。\nON：开启。\nOFF：关闭。","exampleValue":"ON,OFF"}
	ConfigSwitch *string `json:"configSwitch,omitempty" xml:"configSwitch,omitempty" require:"true"`
	// {"en":"Rule list.","zh_CN":"规则列表。"}
	RuleList []*RateLimitingRule `json:"ruleList,omitempty" xml:"ruleList,omitempty" require:"true" type:"Repeated"`
}

func (s PreDeployRateLimitingConfigurationRequest) String() string {
	return tea.Prettify(s)
}

func (s PreDeployRateLimitingConfigurationRequest) GoString() string {
	return s.String()
}

func (s *PreDeployRateLimitingConfigurationRequest) SetDomain(v string) *PreDeployRateLimitingConfigurationRequest {
	s.Domain = &v
	return s
}

func (s *PreDeployRateLimitingConfigurationRequest) SetConfigSwitch(v string) *PreDeployRateLimitingConfigurationRequest {
	s.ConfigSwitch = &v
	return s
}

func (s *PreDeployRateLimitingConfigurationRequest) SetRuleList(v []*RateLimitingRule) *PreDeployRateLimitingConfigurationRequest {
	s.RuleList = v
	return s
}

type RateLimitingRule struct {
	// {"en":"Rule Name, maximum 50 characters.\nDoes not support special characters and spaces.","zh_CN":"规则名称，最多50个字符。\n不支持特殊字符和空格。"}
	RuleName *string `json:"ruleName,omitempty" xml:"ruleName,omitempty" require:"true"`
	// {"en":"Description, maximum 200 characters.","zh_CN":"规则描述，最多200个字符。"}
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// {"en":"Protected target.\nWEB:Website\nAPI:API","zh_CN":"业务场景。\nWEB：网站业务\nAPI：API业务","exampleValue":"WEB,API"}
	Scene *string `json:"scene,omitempty" xml:"scene,omitempty" require:"true"`
	// {"en":"Client identifier.\nIP:Client IP\nIP_UA:Client IP and User-Agent\nCOOKIE:Cookie\nIP_COOKIE:Client IP and Cookie\nHEADER:Request Header\nWhen there is a status code in the matching condition,this client identifier is not supported.\nIP_HEADER:Client IP and Request Header\nWhen there is a status code in the matching condition,this client identifier is not supported .","zh_CN":"统计粒度。\nIP：客户端IP\nIP_UA：客户端IP和User-Agent\nCOOKIE：Cookie\nIP_COOKIE：客户端IP和Cookie\nHEADER：请求头，当匹配条件中存在状态码时不支持此统计粒度\nIP_HEADER：客户端IP和请求头，当匹配条件中存在状态码时不支持此统计粒度","exampleValue":"IP,IP_UA,COOKIE,IP_COOKIE,HEADER,IP_HEADER"}
	StatisticalItem *string `json:"statisticalItem,omitempty" xml:"statisticalItem,omitempty" require:"true"`
	// {"en":"Statistical key value.\nWhen the client identifier is cookie/header value, the corresponding key value needs to be entered.","zh_CN":"统计key值。\n当统计粒度cookie/header值，需要输入对应的key值。"}
	StatisticsKey *string `json:"statisticsKey,omitempty" xml:"statisticsKey,omitempty"`
	// {"en":"Statistics period, unit: seconds, the range is 1 - 3600.","zh_CN":"统计周期，单位：秒，范围为 1 - 3600。"}
	StatisticalPeriod *int `json:"statisticalPeriod,omitempty" xml:"statisticalPeriod,omitempty" require:"true"`
	// {"en":"Trigger threshold, unit: times.","zh_CN":"触发阈值，单位：次。"}
	TriggerThreshold *int `json:"triggerThreshold,omitempty" xml:"triggerThreshold,omitempty" require:"true"`
	// {"en":"Action duration, unit: seconds, the range is 10 - 604800.","zh_CN":"处理动作持续时间，单位：秒，范围为 10 - 604800。"}
	InterceptTime *int `json:"interceptTime,omitempty" xml:"interceptTime,omitempty" require:"true"`
	// {"en":"Cycle effective status.\nPERMANENT:All time\nWITHOUT:Excluded time\nWITHIN:Selected time","zh_CN":"周期生效状态。\nPERMANENT：永久生效\nWITHOUT：周期内不生效\nWITHIN：周期内生效","exampleValue":"PERMANENT,WITHOUT,WITHIN"}
	EffectiveStatus *string `json:"effectiveStatus,omitempty" xml:"effectiveStatus,omitempty" require:"true"`
	// {"en":"Effective time period.When the effective status is effective within the cycle or not effective within the cycle, this field must have a value.","zh_CN":"规则生效周期。\n生效状态为周期内生效或周期内不生效时，此字段必须有值。"}
	RateLimitEffective *RateLimitEffective `json:"rateLimitEffective,omitempty" xml:"rateLimitEffective,omitempty" type:"Struct"`
	// {"en":"API ID under API business, multiple separated by ; sign.\nWhen the protected target is APIThis field is required.","zh_CN":"API业务下的API ID，多个用 ; 隔开。\n当业务场景为API业务时此字段必填。"}
	AssetApiId *string `json:"assetApiId,omitempty" xml:"assetApiId,omitempty"`
	// {"en":"Action.\nNO_USE:Not Used\nLOG:Log\nCOOKIE:Cookie verification\nJS_CHECK:Javascript verification\nDELAY:Delay\nBLOCK:Deny\nRESET:Reset Connection\nCustom response ID:Custom response ID\nWhen there is a status code in the matching condition, the supported actions are Log, Deny, NO_USE, and Reset, Connection.","zh_CN":"处理动作。\nNO_USE：不使用\nLOG：监控\nCOOKIE：Cookie校验\nJS_CHECK：JavaScript校验\nDELAY：延迟响应\nBLOCK：拦截\nRESET：断开连接\n自定义响应ID：自定义响应ID\n当匹配条件中存在状态码时，支持处理动作为监控、拦截、不使用、断开连接。"}
	Action *string `json:"action,omitempty" xml:"action,omitempty" require:"true"`
	// {"en":"Matching conditions.","zh_CN":"匹配条件。"}
	RateLimitRuleCondition *RateLimitingRuleCondition `json:"rateLimitRuleCondition,omitempty" xml:"rateLimitRuleCondition,omitempty" require:"true" type:"Struct"`
}

func (s RateLimitingRule) String() string {
	return tea.Prettify(s)
}

func (s RateLimitingRule) GoString() string {
	return s.String()
}

func (s *RateLimitingRule) SetRuleName(v string) *RateLimitingRule {
	s.RuleName = &v
	return s
}

func (s *RateLimitingRule) SetDescription(v string) *RateLimitingRule {
	s.Description = &v
	return s
}

func (s *RateLimitingRule) SetScene(v string) *RateLimitingRule {
	s.Scene = &v
	return s
}

func (s *RateLimitingRule) SetStatisticalItem(v string) *RateLimitingRule {
	s.StatisticalItem = &v
	return s
}

func (s *RateLimitingRule) SetStatisticsKey(v string) *RateLimitingRule {
	s.StatisticsKey = &v
	return s
}

func (s *RateLimitingRule) SetStatisticalPeriod(v int) *RateLimitingRule {
	s.StatisticalPeriod = &v
	return s
}

func (s *RateLimitingRule) SetTriggerThreshold(v int) *RateLimitingRule {
	s.TriggerThreshold = &v
	return s
}

func (s *RateLimitingRule) SetInterceptTime(v int) *RateLimitingRule {
	s.InterceptTime = &v
	return s
}

func (s *RateLimitingRule) SetEffectiveStatus(v string) *RateLimitingRule {
	s.EffectiveStatus = &v
	return s
}

func (s *RateLimitingRule) SetRateLimitEffective(v *RateLimitEffective) *RateLimitingRule {
	s.RateLimitEffective = v
	return s
}

func (s *RateLimitingRule) SetAssetApiId(v string) *RateLimitingRule {
	s.AssetApiId = &v
	return s
}

func (s *RateLimitingRule) SetAction(v string) *RateLimitingRule {
	s.Action = &v
	return s
}

func (s *RateLimitingRule) SetRateLimitRuleCondition(v *RateLimitingRuleCondition) *RateLimitingRule {
	s.RateLimitRuleCondition = v
	return s
}

type RateLimitEffective struct {
	// {"en":"Effective.\nMON:Monday\nTUE:Tuesday\nWED:Wednesday\nTHU:Thursday\nFRI:Friday\nSAT:Saturday\nSUN:Sunday","zh_CN":"周期。\nMON：星期一\nTUE：星期二\nWED：星期三\nTHU：星期四\nFRI：星期五\nSAT：星期六\nSUN：星期天","exampleValue":"MON,TUE,WED,THU,FRI,SAT,SUN"}
	Effective []*string `json:"effective,omitempty" xml:"effective,omitempty" require:"true" type:"Repeated"`
	// {"en":"Start time, format: HH:mm.","zh_CN":"开始时间，格式：HH:mm。"}
	Start *string `json:"start,omitempty" xml:"start,omitempty" require:"true"`
	// {"en":"End time, format: HH:mm.","zh_CN":"结束时间，格式：HH:mm。"}
	End *string `json:"end,omitempty" xml:"end,omitempty" require:"true"`
	// {"dictionary":"belong=WAAP-MS-Ext|dict=waap_timezone","en":"Timezone,default value: GTM+8.","zh_CN":"时区，默认：GTM+8。"}
	Timezone *string `json:"timezone,omitempty" xml:"timezone,omitempty" require:"true"`
}

func (s RateLimitEffective) String() string {
	return tea.Prettify(s)
}

func (s RateLimitEffective) GoString() string {
	return s.String()
}

func (s *RateLimitEffective) SetEffective(v []*string) *RateLimitEffective {
	s.Effective = v
	return s
}

func (s *RateLimitEffective) SetStart(v string) *RateLimitEffective {
	s.Start = &v
	return s
}

func (s *RateLimitEffective) SetEnd(v string) *RateLimitEffective {
	s.End = &v
	return s
}

func (s *RateLimitEffective) SetTimezone(v string) *RateLimitEffective {
	s.Timezone = &v
	return s
}

type RateLimitingRuleCondition struct {
	// {"en":"IP/CIDR, match type cannot be repeated.","zh_CN":"IP/IP段，匹配类型不可重复。"}
	IpOrIpsConditions []*IpOrIpsConditions `json:"ipOrIpsConditions,omitempty" xml:"ipOrIpsConditions,omitempty" type:"Repeated"`
	// {"en":"Path, match type cannot be repeated.\nWhen the business scenario is API, this matching condition is not supported.","zh_CN":"路径，匹配类型不可重复。\n当业务场景为API业务时不支持此匹配条件。"}
	PathConditions []*PathConditions `json:"pathConditions,omitempty" xml:"pathConditions,omitempty" type:"Repeated"`
	// {"en":"URI, match type cannot be repeated.\nWhen the business scenario is API, this matching condition is not supported.","zh_CN":"URI，匹配类型不可重复。\n当业务场景为API业务时不支持此匹配条件。"}
	UriConditions []*UriConditions `json:"uriConditions,omitempty" xml:"uriConditions,omitempty" type:"Repeated"`
	// {"en":"URI Parameter, match type cannot be repeated.\nWhen the business scenario is API, this matching condition is not supported.","zh_CN":"URI参数，匹配类型不可重复。\n当业务场景为网站业务时不支持此匹配条件。"}
	UriParamConditions []*UriParamConditions `json:"uriParamConditions,omitempty" xml:"uriParamConditions,omitempty" type:"Repeated"`
	// {"en":"User Agent, match type cannot be repeated.","zh_CN":"User-Agent，匹配类型不可重复。"}
	UaConditions []*UaConditions `json:"uaConditions,omitempty" xml:"uaConditions,omitempty" type:"Repeated"`
	// {"en":"Request Method.\nWhen the business scenario is API,this matching condition is not supported.","zh_CN":"请求方法，匹配类型不可重复。\n当业务场景为API业务时不支持此匹配条件。"}
	MethodConditions []*MethodConditions `json:"methodConditions,omitempty" xml:"methodConditions,omitempty" type:"Repeated"`
	// {"en":"Referer, match type cannot be repeated.","zh_CN":"Referer，匹配类型不可重复。"}
	RefererConditions []*RefererConditions `json:"refererConditions,omitempty" xml:"refererConditions,omitempty" type:"Repeated"`
	// {"en":"Request Header, match type can be repeated.","zh_CN":"请求头，匹配类型可重复。"}
	HeaderConditions []*HeaderConditions `json:"headerConditions,omitempty" xml:"headerConditions,omitempty" type:"Repeated"`
	// {"en":"Geo, match type cannot be repeated.","zh_CN":"区域，匹配类型不可重复。"}
	AreaConditions []*AreaConditions `json:"areaConditions,omitempty" xml:"areaConditions,omitempty" type:"Repeated"`
	// {"en":"HTTP/S, match type cannot be repeated.","zh_CN":"应用层协议，匹配类型不可重复。"}
	SchemeConditions []*SchemeConditions `json:"schemeConditions,omitempty" xml:"schemeConditions,omitempty" type:"Repeated"`
	// {"en":"Response Code, match type cannot be repeated.","zh_CN":"状态码，匹配类型不可重复。"}
	StatusCodeConditions []*StatusCodeConditions `json:"statusCodeConditions,omitempty" xml:"statusCodeConditions,omitempty" type:"Repeated"`
	// {"en":"JA3 Fingerprint, match type cannot be repeated.","zh_CN":"JA3指纹，匹配类型不可重复。"}
	Ja3Conditions []*Ja3Conditions `json:"ja3Conditions,omitempty" xml:"ja3Conditions,omitempty" type:"Repeated"`
	// {"en":"JA4 Fingerprint, match type cannot be repeated.","zh_CN":"JA4指纹，匹配类型不可重复。"}
	Ja4Conditions []*Ja4Conditions `json:"ja4Conditions,omitempty" xml:"ja4Conditions,omitempty" type:"Repeated"`
}

func (s RateLimitingRuleCondition) String() string {
	return tea.Prettify(s)
}

func (s RateLimitingRuleCondition) GoString() string {
	return s.String()
}

func (s *RateLimitingRuleCondition) SetIpOrIpsConditions(v []*IpOrIpsConditions) *RateLimitingRuleCondition {
	s.IpOrIpsConditions = v
	return s
}

func (s *RateLimitingRuleCondition) SetPathConditions(v []*PathConditions) *RateLimitingRuleCondition {
	s.PathConditions = v
	return s
}

func (s *RateLimitingRuleCondition) SetUriConditions(v []*UriConditions) *RateLimitingRuleCondition {
	s.UriConditions = v
	return s
}

func (s *RateLimitingRuleCondition) SetUriParamConditions(v []*UriParamConditions) *RateLimitingRuleCondition {
	s.UriParamConditions = v
	return s
}

func (s *RateLimitingRuleCondition) SetUaConditions(v []*UaConditions) *RateLimitingRuleCondition {
	s.UaConditions = v
	return s
}

func (s *RateLimitingRuleCondition) SetMethodConditions(v []*MethodConditions) *RateLimitingRuleCondition {
	s.MethodConditions = v
	return s
}

func (s *RateLimitingRuleCondition) SetRefererConditions(v []*RefererConditions) *RateLimitingRuleCondition {
	s.RefererConditions = v
	return s
}

func (s *RateLimitingRuleCondition) SetHeaderConditions(v []*HeaderConditions) *RateLimitingRuleCondition {
	s.HeaderConditions = v
	return s
}

func (s *RateLimitingRuleCondition) SetAreaConditions(v []*AreaConditions) *RateLimitingRuleCondition {
	s.AreaConditions = v
	return s
}

func (s *RateLimitingRuleCondition) SetSchemeConditions(v []*SchemeConditions) *RateLimitingRuleCondition {
	s.SchemeConditions = v
	return s
}

func (s *RateLimitingRuleCondition) SetStatusCodeConditions(v []*StatusCodeConditions) *RateLimitingRuleCondition {
	s.StatusCodeConditions = v
	return s
}

func (s *RateLimitingRuleCondition) SetJa3Conditions(v []*Ja3Conditions) *RateLimitingRuleCondition {
	s.Ja3Conditions = v
	return s
}

func (s *RateLimitingRuleCondition) SetJa4Conditions(v []*Ja4Conditions) *RateLimitingRuleCondition {
	s.Ja4Conditions = v
	return s
}

type PreDeployRateLimitingConfigurationResponse struct {
	// {"dictionary":"belong=WAAP-MS-Ext|dict=waap_retCodeEnum","en":"Please refer to the error code for exceptions.","zh_CN":"请参照错误码。"}
	Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
	// {"en":"Description.","zh_CN":"描述信息。"}
	Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
	// {"en":"Data.","zh_CN":"出参数据。"}
	Data *PreDeployResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
}

func (s PreDeployRateLimitingConfigurationResponse) String() string {
	return tea.Prettify(s)
}

func (s PreDeployRateLimitingConfigurationResponse) GoString() string {
	return s.String()
}

func (s *PreDeployRateLimitingConfigurationResponse) SetCode(v string) *PreDeployRateLimitingConfigurationResponse {
	s.Code = &v
	return s
}

func (s *PreDeployRateLimitingConfigurationResponse) SetMsg(v string) *PreDeployRateLimitingConfigurationResponse {
	s.Msg = &v
	return s
}

func (s *PreDeployRateLimitingConfigurationResponse) SetData(v *PreDeployResponseData) *PreDeployRateLimitingConfigurationResponse {
	s.Data = v
	return s
}
