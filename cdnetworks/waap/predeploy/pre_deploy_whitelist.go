// This file is auto-generated, don't edit it. Thanks.
package predeploy

import (
	"github.com/alibabacloud-go/tea/tea"
)

type PreDeployWhitelistConfigurationRequest struct {
	// {"en":"Domain list.","zh_CN":"域名列表。"}
	Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
	// {"en":"Policy switch.ON: Enable.OFF: Disable.","zh_CN":"策略开关。\nON：开启。\nOFF：关闭。","exampleValue":"ON,OFF"}
	ConfigSwitch *string `json:"configSwitch,omitempty" xml:"configSwitch,omitempty" require:"true"`
	// {"en":"Rule list.","zh_CN":"规则列表。"}
	RuleList []*WhitelistRule `json:"ruleList,omitempty" xml:"ruleList,omitempty" require:"true" type:"Repeated"`
}

func (s PreDeployWhitelistConfigurationRequest) String() string {
	return tea.Prettify(s)
}

func (s PreDeployWhitelistConfigurationRequest) GoString() string {
	return s.String()
}

func (s *PreDeployWhitelistConfigurationRequest) SetDomain(v string) *PreDeployWhitelistConfigurationRequest {
	s.Domain = &v
	return s
}

func (s *PreDeployWhitelistConfigurationRequest) SetConfigSwitch(v string) *PreDeployWhitelistConfigurationRequest {
	s.ConfigSwitch = &v
	return s
}

func (s *PreDeployWhitelistConfigurationRequest) SetRuleList(v []*WhitelistRule) *PreDeployWhitelistConfigurationRequest {
	s.RuleList = v
	return s
}

type WhitelistRule struct {
	// {"en":"Rule name, maximum 50 characters.\nDoes not support special characters and spaces.","zh_CN":"规则名称，最多50个字符。\n不支持特殊字符和空格。"}
	RuleName *string `json:"ruleName,omitempty" xml:"ruleName,omitempty" require:"true"`
	// {"en":"Description, maximum 200 characters.","zh_CN":"描述，最多200个字符。"}
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// {"en":"Match conditions, at least one, at most five.","zh_CN":"匹配条件，至少一个，至多五个。"}
	Conditions *WhitelistRuleConditions `json:"conditions,omitempty" xml:"conditions,omitempty" require:"true" type:"Struct"`
}

func (s WhitelistRule) String() string {
	return tea.Prettify(s)
}

func (s WhitelistRule) GoString() string {
	return s.String()
}

func (s *WhitelistRule) SetRuleName(v string) *WhitelistRule {
	s.RuleName = &v
	return s
}

func (s *WhitelistRule) SetDescription(v string) *WhitelistRule {
	s.Description = &v
	return s
}

func (s *WhitelistRule) SetConditions(v *WhitelistRuleConditions) *WhitelistRule {
	s.Conditions = v
	return s
}

type WhitelistRuleConditions struct {
	// {"en":"IP/CIDR match conditions, match type cannot be repeated.","zh_CN":"IP/IP段匹配条件，匹配类型不可重复。"}
	IpOrIpsConditions []*IpOrIpsConditions `json:"ipOrIpsConditions,omitempty" xml:"ipOrIpsConditions,omitempty" type:"Repeated"`
	// {"en":"Path match conditions, match type cannot be repeated.","zh_CN":"路径匹配条件，匹配类型不可重复。"}
	PathConditions []*PathConditions `json:"pathConditions,omitempty" xml:"pathConditions,omitempty" type:"Repeated"`
	// {"en":"URI match conditions, match type cannot be repeated.","zh_CN":"URI匹配条件，匹配类型不可重复。"}
	UriConditions []*UriConditions `json:"uriConditions,omitempty" xml:"uriConditions,omitempty" type:"Repeated"`
	// {"en":"User agent match conditions, match type cannot be repeated.","zh_CN":"User-Agent 匹配条件，匹配类型不可重复。"}
	UaConditions []*UaConditions `json:"uaConditions,omitempty" xml:"uaConditions,omitempty" type:"Repeated"`
	// {"en":"Referer match conditions, match type cannot be repeated.","zh_CN":"Referer 匹配条件，匹配类型不可重复。"}
	RefererConditions []*RefererConditions `json:"refererConditions,omitempty" xml:"refererConditions,omitempty" type:"Repeated"`
	// {"en":"Request header match conditions.","zh_CN":"请求头匹配条件。"}
	HeaderConditions []*HeaderConditions `json:"headerConditions,omitempty" xml:"headerConditions,omitempty" type:"Repeated"`
}

func (s WhitelistRuleConditions) String() string {
	return tea.Prettify(s)
}

func (s WhitelistRuleConditions) GoString() string {
	return s.String()
}

func (s *WhitelistRuleConditions) SetIpOrIpsConditions(v []*IpOrIpsConditions) *WhitelistRuleConditions {
	s.IpOrIpsConditions = v
	return s
}

func (s *WhitelistRuleConditions) SetPathConditions(v []*PathConditions) *WhitelistRuleConditions {
	s.PathConditions = v
	return s
}

func (s *WhitelistRuleConditions) SetUriConditions(v []*UriConditions) *WhitelistRuleConditions {
	s.UriConditions = v
	return s
}

func (s *WhitelistRuleConditions) SetUaConditions(v []*UaConditions) *WhitelistRuleConditions {
	s.UaConditions = v
	return s
}

func (s *WhitelistRuleConditions) SetRefererConditions(v []*RefererConditions) *WhitelistRuleConditions {
	s.RefererConditions = v
	return s
}

func (s *WhitelistRuleConditions) SetHeaderConditions(v []*HeaderConditions) *WhitelistRuleConditions {
	s.HeaderConditions = v
	return s
}

type PreDeployWhitelistConfigurationResponse struct {
	// {"dictionary":"belong=WAAP-MS-Ext|dict=waap_retCodeEnum","en":"Please refer to the error code for exceptions.","zh_CN":"请参照错误码。"}
	Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
	// {"en":"Description.","zh_CN":"描述信息。"}
	Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
	// {"en":"Data.","zh_CN":"出参数据。"}
	Data *PreDeployResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
}

func (s PreDeployWhitelistConfigurationResponse) String() string {
	return tea.Prettify(s)
}

func (s PreDeployWhitelistConfigurationResponse) GoString() string {
	return s.String()
}

func (s *PreDeployWhitelistConfigurationResponse) SetCode(v string) *PreDeployWhitelistConfigurationResponse {
	s.Code = &v
	return s
}

func (s *PreDeployWhitelistConfigurationResponse) SetMsg(v string) *PreDeployWhitelistConfigurationResponse {
	s.Msg = &v
	return s
}

func (s *PreDeployWhitelistConfigurationResponse) SetData(v *PreDeployResponseData) *PreDeployWhitelistConfigurationResponse {
	s.Data = v
	return s
}
