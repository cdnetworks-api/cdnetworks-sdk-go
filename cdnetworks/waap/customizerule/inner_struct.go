package customizerule

import "github.com/alibabacloud-go/tea/tea"

type CommonCustomizeRuleConditionDTO struct {
	// {'en':'IP/CIDR, match type cannot be repeated.', 'zh_CN':'IP/IP段，匹配类型不可重复。'}
	IpOrIpsConditions []*IpOrIpsCondition `json:"ipOrIpsConditions,omitempty" xml:"ipOrIpsConditions,omitempty" type:"Repeated"`
	// {'en':'Path, match type cannot be repeated.
	// When the business scenario is API, this matching condition is not supported.', 'zh_CN':'路径，匹配类型不可重复。
	// 当业务场景为API业务时不支持此匹配条件。'}
	PathConditions []*PathCondition `json:"pathConditions,omitempty" xml:"pathConditions,omitempty" type:"Repeated"`
	// {'en':'URI, match type cannot be repeated.
	// When the business scenario is API, this matching condition is not supported.', 'zh_CN':'URI，匹配类型不可重复。
	// 当业务场景为API业务时不支持此匹配条件。'}
	UriConditions []*UriCondition `json:"uriConditions,omitempty" xml:"uriConditions,omitempty" type:"Repeated"`
	// {'en':'URI Parameter, match type cannot be repeated.
	// When the business scenario is API, this matching condition is not supported.', 'zh_CN':'URI参数，匹配类型不可重复。
	// 当业务场景为网站业务时不支持此匹配条件。'}
	UriParamConditions []*UriParamCondition `json:"uriParamConditions,omitempty" xml:"uriParamConditions,omitempty" type:"Repeated"`
	// {'en':'User Agent, match type cannot be repeated.', 'zh_CN':'User-Agent，匹配类型不可重复。'}
	UaConditions []*UaCondition `json:"uaConditions,omitempty" xml:"uaConditions,omitempty" type:"Repeated"`
	// {'en':'Referer, match type cannot be repeated.', 'zh_CN':'Referer，匹配类型不可重复。'}
	RefererConditions []*RefererCondition `json:"refererConditions,omitempty" xml:"refererConditions,omitempty" type:"Repeated"`
	// {'en':'Request Header, match type can be repeated.', 'zh_CN':'请求头，匹配类型可重复。'}
	HeaderConditions []*HeaderCondition `json:"headerConditions,omitempty" xml:"headerConditions,omitempty" type:"Repeated"`
	// {'en':'Geo, match type cannot be repeated.', 'zh_CN':'区域，匹配类型不可重复。'}
	AreaConditions []*AreaCondition `json:"areaConditions,omitempty" xml:"areaConditions,omitempty" type:"Repeated"`
	// {'en':'Request Method.
	// When the business scenario is API,this matching condition is not supported.', 'zh_CN':'请求方法，匹配类型不可重复。
	// 当业务场景为API业务时不支持此匹配条件。'}
	MethodConditions []*RequestMethodCondition `json:"methodConditions,omitempty" xml:"methodConditions,omitempty" type:"Repeated"`
}

func (s CommonCustomizeRuleConditionDTO) String() string {
	return tea.Prettify(s)
}

func (s CommonCustomizeRuleConditionDTO) GoString() string {
	return s.String()
}

func (s *CommonCustomizeRuleConditionDTO) SetIpOrIpsConditions(v []*IpOrIpsCondition) *CommonCustomizeRuleConditionDTO {
	s.IpOrIpsConditions = v
	return s
}

func (s *CommonCustomizeRuleConditionDTO) SetPathConditions(v []*PathCondition) *CommonCustomizeRuleConditionDTO {
	s.PathConditions = v
	return s
}

func (s *CommonCustomizeRuleConditionDTO) SetUriConditions(v []*UriCondition) *CommonCustomizeRuleConditionDTO {
	s.UriConditions = v
	return s
}

func (s *CommonCustomizeRuleConditionDTO) SetUriParamConditions(v []*UriParamCondition) *CommonCustomizeRuleConditionDTO {
	s.UriParamConditions = v
	return s
}

func (s *CommonCustomizeRuleConditionDTO) SetUaConditions(v []*UaCondition) *CommonCustomizeRuleConditionDTO {
	s.UaConditions = v
	return s
}

func (s *CommonCustomizeRuleConditionDTO) SetRefererConditions(v []*RefererCondition) *CommonCustomizeRuleConditionDTO {
	s.RefererConditions = v
	return s
}

func (s *CommonCustomizeRuleConditionDTO) SetHeaderConditions(v []*HeaderCondition) *CommonCustomizeRuleConditionDTO {
	s.HeaderConditions = v
	return s
}

func (s *CommonCustomizeRuleConditionDTO) SetAreaConditions(v []*AreaCondition) *CommonCustomizeRuleConditionDTO {
	s.AreaConditions = v
	return s
}

func (s *CommonCustomizeRuleConditionDTO) SetMethodConditions(v []*RequestMethodCondition) *CommonCustomizeRuleConditionDTO {
	s.MethodConditions = v
	return s
}

type IpOrIpsCondition struct {
	// {'en':'Match type.
	// EQUAL:Equal
	// NOT_EQUAL:Does not equal', 'zh_CN':'匹配类型。
	// EQUAL：等于
	// NOT_EQUAL：不等于'}
	MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
	// {'en':'IP/CIDR.', 'zh_CN':'IP/IP段。'}
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

type HeaderCondition struct {
	// {"zh_CN":"匹配类型。
	// EQUAL：等于
	// NOT_EQUAL：不等于
	// CONTAIN：包含
	// NOT_CONTAIN：不包含
	// REGEX：正则
	// NONE：为空或不存在
	// NOT_REGEX：正则不匹配
	// START_WITH：开头是
	// END_WITH：结尾是
	// WILDCARD：通配符匹配
	// NOT_WILDCARD：通配符不匹配","en":"Match type.
	// EQUAL: equal to
	// NOT_EQUAL: not equal to
	// CONTAIN: contains
	// NOT_CONTAIN: does not contain
	// REGEX: regular
	// NONE: empty or does not exist
	// NOT_REGEX: regular does not match
	// START_WITH: starts with
	// END_WITH: ends with
	// WILDCARD: wildcard matches
	// NOT_WILDCARD: wildcard does not match"}
	MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
	// {"zh_CN":"头部名称。","en":"Request header name."}
	Key *string `json:"key,omitempty" xml:"key,omitempty" require:"true"`
	// {"zh_CN":"头部值。","en":"Header value."}
	ValueList []*string `json:"valueList,omitempty" xml:"valueList,omitempty" type:"Repeated"`
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

type AreaCondition struct {
	// {'en':'Match type.
	// EQUAL:Equal
	// NOT_EQUAL:Does not equal', 'zh_CN':'匹配类型。
	// EQUAL：等于
	// NOT_EQUAL：不等于'}
	MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
	// {'en':'Geo.', 'zh_CN':'区域。','dictionary':'belong=WAAP-MS-Ext|dict=waap_areaCityAndCountry'}
	Areas []*string `json:"areas,omitempty" xml:"areas,omitempty" require:"true" type:"Repeated"`
}

func (s AreaCondition) String() string {
	return tea.Prettify(s)
}

func (s AreaCondition) GoString() string {
	return s.String()
}

func (s *AreaCondition) SetMatchType(v string) *AreaCondition {
	s.MatchType = &v
	return s
}

func (s *AreaCondition) SetAreas(v []*string) *AreaCondition {
	s.Areas = v
	return s
}

type UriCondition struct {
	// {"zh_CN":"匹配类型。
	// EQUAL：等于
	// NOT_EQUAL：不等于
	// CONTAIN：包含
	// NOT_CONTAIN：不包含
	// REGEX：正则
	// NOT_REGEX：正则不匹配
	// START_WITH：开头是
	// END_WITH：结尾是
	// WILDCARD：通配符匹配
	// NOT_WILDCARD：通配符不匹配","en":"Match type.
	// EQUAL: equal to
	// NOT_EQUAL: not equal to
	// CONTAIN: contains
	// NOT_CONTAIN: does not contain
	// REGEX: regular
	// NOT_REGEX: regular does not match
	// START_WITH: starts with
	// END_WITH: ends with
	// WILDCARD: wildcard matches
	// NOT_WILDCARD: wildcard does not match"}
	MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
	// {"zh_CN":"URI。","en":"URI."}
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
	// {"zh_CN":"匹配类型。
	// EQUAL：等于
	// NOT_EQUAL：不等于
	// CONTAIN：包含
	// NOT_CONTAIN：不包含
	// REGEX：正则
	// NONE：为空或不存在
	// NOT_REGEX：正则不匹配
	// START_WITH：开头是
	// END_WITH：结尾是
	// WILDCARD：通配符匹配
	// NOT_WILDCARD：通配符不匹配","en":"Match type.
	// EQUAL: equal to
	// NOT_EQUAL: not equal to
	// CONTAIN: contains
	// NOT_CONTAIN: does not contain
	// REGEX: regular
	// NONE: empty or does not exist
	// NOT_REGEX: regular does not match
	// START_WITH: starts with
	// END_WITH: ends with
	// WILDCARD: wildcard matches
	// NOT_WILDCARD: wildcard does not match"}
	MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
	// {"zh_CN":"User-Agent。","en":"User-Agent."}
	Ua []*string `json:"ua,omitempty" xml:"ua,omitempty" type:"Repeated"`
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
	// {"zh_CN":"匹配类型。
	// EQUAL：等于
	// NOT_EQUAL：不等于
	// CONTAIN：包含
	// NOT_CONTAIN：不包含
	// REGEX：正则
	// NONE：为空或不存在
	// NOT_REGEX：正则不匹配
	// START_WITH：开头是
	// END_WITH：结尾是
	// WILDCARD：通配符匹配
	// NOT_WILDCARD：通配符不匹配","en":"Match type.
	// EQUAL: equal to
	// NOT_EQUAL: not equal to
	// CONTAIN: contains
	// NOT_CONTAIN: does not contain
	// REGEX: regular
	// NONE: empty or does not exist
	// NOT_REGEX: regular does not match
	// START_WITH: starts with
	// END_WITH: ends with
	// WILDCARD: wildcard matches
	// NOT_WILDCARD: wildcard does not match"}
	MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
	// {"zh_CN":"Referer。","en":"Referer."}
	Referer []*string `json:"referer,omitempty" xml:"referer,omitempty" type:"Repeated"`
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

type RequestMethodCondition struct {
	// {'en':'Match type.
	// EQUAL:Equal
	// NOT_EQUAL:Does not equal', 'zh_CN':'匹配类型。
	// EQUAL：等于
	// NOT_EQUAL：不等于'}
	MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
	// {'en':'Request method.
	// Supported values: GET/POST/DELETE/PUT/HEAD/OPTIONS/COPY.', 'zh_CN':'请求方法。
	// 支持的值：GET/POST/DELETE/PUT/HEAD/OPTIONS/COPY'}
	RequestMethod []*string `json:"requestMethod,omitempty" xml:"requestMethod,omitempty" require:"true" type:"Repeated"`
}

func (s RequestMethodCondition) String() string {
	return tea.Prettify(s)
}

func (s RequestMethodCondition) GoString() string {
	return s.String()
}

func (s *RequestMethodCondition) SetMatchType(v string) *RequestMethodCondition {
	s.MatchType = &v
	return s
}

func (s *RequestMethodCondition) SetRequestMethod(v []*string) *RequestMethodCondition {
	s.RequestMethod = v
	return s
}

type UriParamCondition struct {
	// {'en':'Match type.
	// EQUAL:Equals
	// NOT_EQUAL:Does not equal
	// CONTAIN:Contains
	// NOT_CONTAIN:Does not contains
	// REGEX:Regex match
	// NONE:Empty or non-existent', 'zh_CN':'匹配类型。
	// EQUAL：等于
	// NOT_EQUAL：不等于
	// CONTAIN：包含
	// NOT_CONTAIN：不包含
	// REGEX：正则
	// NONE：为空或不存在'}
	MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
	// {'en':'Param name.', 'zh_CN':'参数名。'}
	ParamName *string `json:"paramName,omitempty" xml:"paramName,omitempty" require:"true"`
	// {'en':'Param value.', 'zh_CN':'参数值。'}
	ParamValue []*string `json:"paramValue,omitempty" xml:"paramValue,omitempty" type:"Repeated"`
}

func (s UriParamCondition) String() string {
	return tea.Prettify(s)
}

func (s UriParamCondition) GoString() string {
	return s.String()
}

func (s *UriParamCondition) SetMatchType(v string) *UriParamCondition {
	s.MatchType = &v
	return s
}

func (s *UriParamCondition) SetParamName(v string) *UriParamCondition {
	s.ParamName = &v
	return s
}

func (s *UriParamCondition) SetParamValue(v []*string) *UriParamCondition {
	s.ParamValue = v
	return s
}

type PathCondition struct {
	// {"zh_CN":"匹配类型。
	// EQUAL：等于
	// NOT_EQUAL：不等于
	// CONTAIN：包含
	// NOT_CONTAIN：不包含
	// REGEX：正则
	// NOT_REGEX：正则不匹配
	// START_WITH：开头是
	// END_WITH：结尾是
	// WILDCARD：通配符匹配
	// NOT_WILDCARD：通配符不匹配","en":"Match type.
	// EQUAL: equal to
	// NOT_EQUAL: not equal to
	// CONTAIN: contains
	// NOT_CONTAIN: does not contain
	// REGEX: regular
	// NOT_REGEX: regular does not match
	// START_WITH: starts with
	// END_WITH: ends with
	// WILDCARD: wildcard matches
	// NOT_WILDCARD: wildcard does not match"}
	MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
	// {"zh_CN":"路径。","en":"Path."}
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
