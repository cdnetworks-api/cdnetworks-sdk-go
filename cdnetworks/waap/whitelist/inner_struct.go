package whitelist

import "github.com/alibabacloud-go/tea/tea"

type WhitelistRuleCondition struct {
	// {"en":"IP/CIDR match conditions.", "zh_CN":"IP/IP段匹配条件。"}
	IpOrIpsConditions []*IpOrIpsCondition `json:"ipOrIpsConditions,omitempty" xml:"ipOrIpsConditions,omitempty" require:"true" type:"Repeated"`
	// {"en":"Path match conditions.", "zh_CN":"路径匹配条件。"}
	PathConditions []*PathCondition `json:"pathConditions,omitempty" xml:"pathConditions,omitempty" require:"true" type:"Repeated"`
	// {"en":"URI match conditions.", "zh_CN":"URI匹配条件。"}
	UriConditions []*UriCondition `json:"uriConditions,omitempty" xml:"uriConditions,omitempty" require:"true" type:"Repeated"`
	// {"en":"User agent match conditions.", "zh_CN":"User-Agent 匹配条件。"}
	UaConditions []*UaCondition `json:"uaConditions,omitempty" xml:"uaConditions,omitempty" require:"true" type:"Repeated"`
	// {"en":"Referer match conditions.", "zh_CN":"Referer 匹配条件。"}
	RefererConditions []*RefererCondition `json:"refererConditions,omitempty" xml:"refererConditions,omitempty" require:"true" type:"Repeated"`
	// {"en":"Request header match conditions.", "zh_CN":"请求头匹配条件。"}
	HeaderConditions []*HeaderCondition `json:"headerConditions,omitempty" xml:"headerConditions,omitempty" require:"true" type:"Repeated"`
}

func (s WhitelistRuleCondition) String() string {
	return tea.Prettify(s)
}

func (s WhitelistRuleCondition) GoString() string {
	return s.String()
}

func (s *WhitelistRuleCondition) SetIpOrIpsConditions(v []*IpOrIpsCondition) *WhitelistRuleCondition {
	s.IpOrIpsConditions = v
	return s
}

func (s *WhitelistRuleCondition) SetPathConditions(v []*PathCondition) *WhitelistRuleCondition {
	s.PathConditions = v
	return s
}

func (s *WhitelistRuleCondition) SetUriConditions(v []*UriCondition) *WhitelistRuleCondition {
	s.UriConditions = v
	return s
}

func (s *WhitelistRuleCondition) SetUaConditions(v []*UaCondition) *WhitelistRuleCondition {
	s.UaConditions = v
	return s
}

func (s *WhitelistRuleCondition) SetRefererConditions(v []*RefererCondition) *WhitelistRuleCondition {
	s.RefererConditions = v
	return s
}

func (s *WhitelistRuleCondition) SetHeaderConditions(v []*HeaderCondition) *WhitelistRuleCondition {
	s.HeaderConditions = v
	return s
}

type IpOrIpsCondition struct {
	// {"en":"Match type.
	// EQUAL: Equals
	// NOT_EQUAL: Does not equal", "zh_CN":"匹配类型。
	// EQUAL：等于
	// NOT_EQUAL：不等于"}
	MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
	// {"en":"IP/CIDR.", "zh_CN":"IP/IP段。"}
	IpOrIps []*string `json:"ipOrIps,omitempty" xml:"ipOrIps,omitempty" require:"true" type:"Repeated"`
}

func (s IpOrIpsCondition) String() string {
	return tea.Prettify(s)
}

func (s IpOrIpsCondition) GoString() string {
	return s.String()
}

func (s *IpOrIpsCondition) SetMatchType(v string) *IpOrIpsCondition {
	s.MatchType = &v
	return s
}

func (s *IpOrIpsCondition) SetIpOrIps(v []*string) *IpOrIpsCondition {
	s.IpOrIps = v
	return s
}

type PathCondition struct {
	// {"en":"Match type.
	// EQUAL: Equals
	// NOT_EQUAL: Does not equal
	// CONTAIN: Contains
	// NOT_CONTAIN: Does not Contains
	// REGEX: Regex match", "zh_CN":"匹配类型。
	// EQUAL：等于
	// NOT_EQUAL：不等于
	// CONTAIN：包含
	// NOT_CONTAIN：不包含
	// REGEX：匹配正则"}
	MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
	// {"en":"Path.", "zh_CN":"路径。"}
	Paths []*string `json:"paths,omitempty" xml:"paths,omitempty" require:"true" type:"Repeated"`
}

func (s PathCondition) String() string {
	return tea.Prettify(s)
}

func (s PathCondition) GoString() string {
	return s.String()
}

func (s *PathCondition) SetMatchType(v string) *PathCondition {
	s.MatchType = &v
	return s
}

func (s *PathCondition) SetPaths(v []*string) *PathCondition {
	s.Paths = v
	return s
}

type UriCondition struct {
	// {"en":"Match type.
	// EQUAL: Equals
	// NOT_EQUAL: Does not equal
	// CONTAIN: Contains
	// NOT_CONTAIN: Does not Contains
	// REGEX: Regex match", "zh_CN":"匹配类型。
	// EQUAL：等于
	// NOT_EQUAL：不等于
	// CONTAIN：包含
	// NOT_CONTAIN：不包含
	// REGEX：匹配正则"}
	MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
	// {"en":"URI.", "zh_CN":"URI。"}
	Uri []*string `json:"uri,omitempty" xml:"uri,omitempty" require:"true" type:"Repeated"`
}

func (s UriCondition) String() string {
	return tea.Prettify(s)
}

func (s UriCondition) GoString() string {
	return s.String()
}

func (s *UriCondition) SetMatchType(v string) *UriCondition {
	s.MatchType = &v
	return s
}

func (s *UriCondition) SetUri(v []*string) *UriCondition {
	s.Uri = v
	return s
}

type UaCondition struct {
	// {"en":"Match type.
	// EQUAL: Equals
	// NOT_EQUAL: Does not equal
	// CONTAIN: Contains
	// NOT_CONTAIN: Does not Contains
	// REGEX: Regex match", "zh_CN":"匹配类型。
	// EQUAL：等于
	// NOT_EQUAL：不等于
	// CONTAIN：包含
	// NOT_CONTAIN：不包含
	// REGEX：匹配正则"}
	MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
	// {"en":"User agent.", "zh_CN":"User-Agent。"}
	Ua []*string `json:"ua,omitempty" xml:"ua,omitempty" require:"true" type:"Repeated"`
}

func (s UaCondition) String() string {
	return tea.Prettify(s)
}

func (s UaCondition) GoString() string {
	return s.String()
}

func (s *UaCondition) SetMatchType(v string) *UaCondition {
	s.MatchType = &v
	return s
}

func (s *UaCondition) SetUa(v []*string) *UaCondition {
	s.Ua = v
	return s
}

type RefererCondition struct {
	// {"en":"Match type.
	// EQUAL: Equals
	// NOT_EQUAL: Does not equal
	// CONTAIN: Contains
	// NOT_CONTAIN: Does not Contains
	// REGEX: Regex match", "zh_CN":"匹配类型。
	// EQUAL：等于
	// NOT_EQUAL：不等于
	// CONTAIN：包含
	// NOT_CONTAIN：不包含
	// REGEX：匹配正则"}
	MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
	// {"en":"Referer.", "zh_CN":"Referer。"}
	Referer []*string `json:"referer,omitempty" xml:"referer,omitempty" require:"true" type:"Repeated"`
}

func (s RefererCondition) String() string {
	return tea.Prettify(s)
}

func (s RefererCondition) GoString() string {
	return s.String()
}

func (s *RefererCondition) SetMatchType(v string) *RefererCondition {
	s.MatchType = &v
	return s
}

func (s *RefererCondition) SetReferer(v []*string) *RefererCondition {
	s.Referer = v
	return s
}

type HeaderCondition struct {
	// {"en":"Match type.
	// EQUAL: Equals
	// NOT_EQUAL: Does not equal
	// CONTAIN: Contains
	// NOT_CONTAIN: Does not Contains
	// REGEX: Regex match", "zh_CN":"匹配类型。
	// EQUAL：等于
	// NOT_EQUAL：不等于
	// CONTAIN：包含
	// NOT_CONTAIN：不包含
	// REGEX：匹配正则"}
	MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
	// {"en":"Request header key.", "zh_CN":"头部名称。"}
	Key *string `json:"key,omitempty" xml:"key,omitempty" require:"true"`
	// {"en":"List of request header values.", "zh_CN":"头部值列表。"}
	ValueList []*string `json:"valueList,omitempty" xml:"valueList,omitempty" require:"true" type:"Repeated"`
}

func (s HeaderCondition) String() string {
	return tea.Prettify(s)
}

func (s HeaderCondition) GoString() string {
	return s.String()
}

func (s *HeaderCondition) SetMatchType(v string) *HeaderCondition {
	s.MatchType = &v
	return s
}

func (s *HeaderCondition) SetKey(v string) *HeaderCondition {
	s.Key = &v
	return s
}

func (s *HeaderCondition) SetValueList(v []*string) *HeaderCondition {
	s.ValueList = v
	return s
}
