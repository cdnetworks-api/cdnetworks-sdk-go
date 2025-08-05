// This file is auto-generated, don't edit it. Thanks.
package predeploy

import (
	"github.com/alibabacloud-go/tea/tea"
)

type PreDeployCustomRuleConfigurationRequest struct {
	// {"en":"Domain list.","zh_CN":"域名列表。"}
	Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
	// {"en":"Policy switch.ON: Enable.OFF: Disable.","zh_CN":"策略开关。\nON：开启。\nOFF：关闭。","exampleValue":"ON,OFF"}
	ConfigSwitch *string `json:"configSwitch,omitempty" xml:"configSwitch,omitempty" require:"true"`
	// {"en":"Rule list.","zh_CN":"规则列表。"}
	RuleList []*CustomRule `json:"ruleList,omitempty" xml:"ruleList,omitempty" require:"true" type:"Repeated"`
}

func (s PreDeployCustomRuleConfigurationRequest) String() string {
	return tea.Prettify(s)
}

func (s PreDeployCustomRuleConfigurationRequest) GoString() string {
	return s.String()
}

func (s *PreDeployCustomRuleConfigurationRequest) SetDomain(v string) *PreDeployCustomRuleConfigurationRequest {
	s.Domain = &v
	return s
}

func (s *PreDeployCustomRuleConfigurationRequest) SetConfigSwitch(v string) *PreDeployCustomRuleConfigurationRequest {
	s.ConfigSwitch = &v
	return s
}

func (s *PreDeployCustomRuleConfigurationRequest) SetRuleList(v []*CustomRule) *PreDeployCustomRuleConfigurationRequest {
	s.RuleList = v
	return s
}

type CustomRule struct {
	// {"en":"Rule Name, maximum 50 characters.\nDoes not support special characters and spaces.","zh_CN":"规则名称，最多50个字符。\n不支持特殊字符和空格。"}
	RuleName *string `json:"ruleName,omitempty" xml:"ruleName,omitempty" require:"true"`
	// {"en":"Description, maximum 200 characters.","zh_CN":"规则描述，最多200个字符。"}
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// {"en":"Protected target.\nWEB:Website\nAPI:API","zh_CN":"业务场景。\nWEB：网站业务\nAPI：API业务","exampleValue":"WEB,API"}
	Scene *string `json:"scene,omitempty" xml:"scene,omitempty" require:"true"`
	// {"en":"API ID under API business, multiple separated by ; sign.\nWhen the protected target is APIThis field is required.","zh_CN":"API业务下的API ID，多个用 ; 隔开。\n当业务场景为API业务时此字段必填。"}
	ApiId *string `json:"apiId,omitempty" xml:"apiId,omitempty"`
	// {"en":"Action.\nNO_USE:Not Used\nLOG:Log\nDELAY:Delay\nBLOCK:Deny\nRESET:Reset Connection","zh_CN":"处理动作。\nNO_USE：不使用\nLOG：监控\nDELAY：延迟响应\nBLOCK：拦截\nRESET:断开连接","exampleValue":"NO_USE,LOG,DELAY,BLOCK,RESET"}
	Act *string `json:"act,omitempty" xml:"act,omitempty" require:"true"`
	// {"en":"Match Conditions.","zh_CN":"匹配条件。"}
	Condition *CustomRuleCondition `json:"condition,omitempty" xml:"condition,omitempty" require:"true" type:"Struct"`
}

func (s CustomRule) String() string {
	return tea.Prettify(s)
}

func (s CustomRule) GoString() string {
	return s.String()
}

func (s *CustomRule) SetRuleName(v string) *CustomRule {
	s.RuleName = &v
	return s
}

func (s *CustomRule) SetDescription(v string) *CustomRule {
	s.Description = &v
	return s
}

func (s *CustomRule) SetScene(v string) *CustomRule {
	s.Scene = &v
	return s
}

func (s *CustomRule) SetApiId(v string) *CustomRule {
	s.ApiId = &v
	return s
}

func (s *CustomRule) SetAct(v string) *CustomRule {
	s.Act = &v
	return s
}

func (s *CustomRule) SetCondition(v *CustomRuleCondition) *CustomRule {
	s.Condition = v
	return s
}

type CustomRuleCondition struct {
	// {"en":"Request Method.\nWhen the business scenario is API,this matching condition is not supported.","zh_CN":"请求方法，匹配类型不可重复。\n当业务场景为API业务时不支持此匹配条件。"}
	MethodConditions []*MethodConditions `json:"methodConditions,omitempty" xml:"methodConditions,omitempty" type:"Repeated"`
	// {"en":"Geo, match type cannot be repeated.","zh_CN":"区域，匹配类型不可重复。"}
	AreaConditions []*AreaConditions `json:"areaConditions,omitempty" xml:"areaConditions,omitempty" type:"Repeated"`
	// {"en":"IP/CIDR, match type cannot be repeated.","zh_CN":"IP/IP段，匹配类型不可重复。"}
	IpOrIpsConditions []*IpOrIpsConditions `json:"ipOrIpsConditions,omitempty" xml:"ipOrIpsConditions,omitempty" type:"Repeated"`
	// {"en":"URI, match type cannot be repeated.\nWhen the business scenario is API, this matching condition is not supported.","zh_CN":"URI，匹配类型不可重复。\n当业务场景为API业务时不支持此匹配条件。"}
	UriConditions []*UriConditions `json:"uriConditions,omitempty" xml:"uriConditions,omitempty" type:"Repeated"`
	// {"en":"Path, match type cannot be repeated.\nWhen the business scenario is API, this matching condition is not supported.","zh_CN":"路径，匹配类型不可重复。\n当业务场景为API业务时不支持此匹配条件。"}
	PathConditions []*PathConditions `json:"pathConditions,omitempty" xml:"pathConditions,omitempty" type:"Repeated"`
	// {"en":"URI Parameter, match type cannot be repeated.\nWhen the business scenario is API, this matching condition is not supported.","zh_CN":"URI参数，匹配类型不可重复。\n当业务场景为网站业务时不支持此匹配条件。"}
	UriParamConditions []*UriParamConditions `json:"uriParamConditions,omitempty" xml:"uriParamConditions,omitempty" type:"Repeated"`
	// {"en":"User Agent, match type cannot be repeated.","zh_CN":"User-Agent，匹配类型不可重复。"}
	UaConditions []*UaConditions `json:"uaConditions,omitempty" xml:"uaConditions,omitempty" type:"Repeated"`
	// {"en":"Request Header, match type can be repeated.","zh_CN":"请求头，匹配类型可重复。"}
	HeaderConditions []*HeaderConditions `json:"headerConditions,omitempty" xml:"headerConditions,omitempty" type:"Repeated"`
	// {"en":"Referer, match type cannot be repeated.","zh_CN":"Referer，匹配类型不可重复。"}
	RefererConditions []*RefererConditions `json:"refererConditions,omitempty" xml:"refererConditions,omitempty" type:"Repeated"`
	// {"en":"JA3 Fingerprint, match type cannot be repeated.","zh_CN":"JA3指纹，匹配类型不可重复。"}
	Ja3Conditions []*Ja3Conditions `json:"ja3Conditions,omitempty" xml:"ja3Conditions,omitempty" type:"Repeated"`
	// {"en":"JA4 Fingerprint, match type cannot be repeated.","zh_CN":"JA4指纹，匹配类型不可重复。"}
	Ja4Conditions []*Ja4Conditions `json:"ja4Conditions,omitempty" xml:"ja4Conditions,omitempty" type:"Repeated"`
}

func (s CustomRuleCondition) String() string {
	return tea.Prettify(s)
}

func (s CustomRuleCondition) GoString() string {
	return s.String()
}

func (s *CustomRuleCondition) SetMethodConditions(v []*MethodConditions) *CustomRuleCondition {
	s.MethodConditions = v
	return s
}

func (s *CustomRuleCondition) SetAreaConditions(v []*AreaConditions) *CustomRuleCondition {
	s.AreaConditions = v
	return s
}

func (s *CustomRuleCondition) SetIpOrIpsConditions(v []*IpOrIpsConditions) *CustomRuleCondition {
	s.IpOrIpsConditions = v
	return s
}

func (s *CustomRuleCondition) SetUriConditions(v []*UriConditions) *CustomRuleCondition {
	s.UriConditions = v
	return s
}

func (s *CustomRuleCondition) SetPathConditions(v []*PathConditions) *CustomRuleCondition {
	s.PathConditions = v
	return s
}

func (s *CustomRuleCondition) SetUriParamConditions(v []*UriParamConditions) *CustomRuleCondition {
	s.UriParamConditions = v
	return s
}

func (s *CustomRuleCondition) SetUaConditions(v []*UaConditions) *CustomRuleCondition {
	s.UaConditions = v
	return s
}

func (s *CustomRuleCondition) SetHeaderConditions(v []*HeaderConditions) *CustomRuleCondition {
	s.HeaderConditions = v
	return s
}

func (s *CustomRuleCondition) SetRefererConditions(v []*RefererConditions) *CustomRuleCondition {
	s.RefererConditions = v
	return s
}

func (s *CustomRuleCondition) SetJa3Conditions(v []*Ja3Conditions) *CustomRuleCondition {
	s.Ja3Conditions = v
	return s
}

func (s *CustomRuleCondition) SetJa4Conditions(v []*Ja4Conditions) *CustomRuleCondition {
	s.Ja4Conditions = v
	return s
}

type PreDeployCustomRuleConfigurationResponse struct {
	// {"dictionary":"belong=WAAP-MS-Ext|dict=waap_retCodeEnum","en":"Please refer to the error code for exceptions.","zh_CN":"请参照错误码。"}
	Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
	// {"en":"Description.","zh_CN":"描述信息。"}
	Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
	// {"en":"Data.","zh_CN":"出参数据。"}
	Data *PreDeployResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
}

func (s PreDeployCustomRuleConfigurationResponse) String() string {
	return tea.Prettify(s)
}

func (s PreDeployCustomRuleConfigurationResponse) GoString() string {
	return s.String()
}

func (s *PreDeployCustomRuleConfigurationResponse) SetCode(v string) *PreDeployCustomRuleConfigurationResponse {
	s.Code = &v
	return s
}

func (s *PreDeployCustomRuleConfigurationResponse) SetMsg(v string) *PreDeployCustomRuleConfigurationResponse {
	s.Msg = &v
	return s
}

func (s *PreDeployCustomRuleConfigurationResponse) SetData(v *PreDeployResponseData) *PreDeployCustomRuleConfigurationResponse {
	s.Data = v
	return s
}
