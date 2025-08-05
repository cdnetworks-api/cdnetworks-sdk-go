package ratelimit

import "github.com/alibabacloud-go/tea/tea"

type IpOrIpsCondition struct {
	// {'en':'Match type.
	// EQUAL:Equal
	// NOT_EQUAL:Does not equal', 'zh_CN':'匹配类型。
	// EQUAL：等于
	// NOT_EQUAL：不等于'}
	MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
	// {'en':'IP/CIDR, maximum 300 IP/CIDR.', 'zh_CN':'IP/IP段，最多300个IP/IP段。'}
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
	// {'en':'Match type.
	// EQUAL:Equals,header value case sensitive
	// NOT_EQUAL:Does not equal,header value case sensitive
	// CONTAIN:Contains,header value case insensitive
	// NOT_CONTAIN:Does not contains,header value case insensitive
	// REGEX:Regex match,header value case insensitive
	// NONE:Empty or non-existent', 'zh_CN':'匹配类型。
	// EQUAL：等于，头部值大小写敏感
	// NOT_EQUAL：不等于，头部值大小写敏感
	// CONTAIN：包含，头部值大小写不敏感
	// NOT_CONTAIN：不包含，头部值大小写不敏感
	// REGEX：正则，头部值大小写不敏感
	// NONE：为空或不存在'}
	MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
	// {'en':'Header name,case insensitive,maximum 100 characters.
	// Example: Accept.', 'zh_CN':'头部名称，大小写不敏感，最多100个字符。
	// 示例：Accept。'}
	Key *string `json:"key,omitempty" xml:"key,omitempty" require:"true"`
	// {'en':'Header value.
	// When the match type is regex match, only one value is allowed.', 'zh_CN':'头部值。
	// 当匹配类型为正则，则只允许只有一个值。'}
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
	// {'en':'Match type.
	// EQUAL:Equals,uri case sensitive
	// NOT_EQUAL:Does not equal,uri case sensitive
	// CONTAIN:Contains,uri case insensitive
	// NOT_CONTAIN:Does not contains,uri case insensitive
	// REGEX:Regex match,uri case insensitive', 'zh_CN':'匹配类型。
	// EQUAL：等于，URI大小写敏感
	// NOT_EQUAL：不等于，URI大小写敏感
	// CONTAIN：包含，URI大小写不敏感
	// NOT_CONTAIN：不包含，URI大小写不敏感
	// REGEX：正则，URI大小写不敏感'}
	MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
	// {'en':'URI.
	// When match type is EQUAL/NOT_EQUAL, uri needs to start with "/", and includes parameters.
	// When the match type is regex match, only one value is allowed.
	// Example: /test.html?id=1.', 'zh_CN':'URI。
	// 当匹配类型为等于/不等于，URI必须以”/“开头，含参数。
	// 当匹配类型为正则，则只允许只有一个值。
	// 示例：/test.html?id=1。'}
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
	// {'en':'Match type.
	// EQUAL:Equals,ua case sensitive
	// NOT_EQUAL:Does not equal,ua case sensitive
	// CONTAIN:Contains,ua case insensitive
	// NOT_CONTAIN:Does not contains,ua case insensitive
	// REGEX:Regex match,ua case insensitive
	// NONE:Empty or non-existent', 'zh_CN':'匹配类型。
	// EQUAL：等于，ua大小写敏感
	// NOT_EQUAL：不等于，ua大小写敏感
	// CONTAIN：包含，ua大小写不敏感
	// NOT_CONTAIN：不包含，ua大小写不敏感
	// REGEX：正则，ua大小写不敏感
	// NONE：为空或不存在'}
	MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
	// {'en':'User agent.
	// When the match type is regex match, only one value is allowed.
	// Example: go-Http-client/1.1.', 'zh_CN':'User-Agent。
	// 当匹配类型为正则，则只允许只有一个值。
	// 示例：go-Http-client/1.1。'}
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
	// {'en':'Match type.
	// EQUAL:Equals,referer case sensitive
	// NOT_EQUAL:Does not equal,referer case sensitive
	// CONTAIN:Contains,referer case insensitive
	// NOT_CONTAIN:Does not contains,referer case insensitive
	// REGEX:Regex match,referer case insensitive
	// NONE:Empty or non-existent', 'zh_CN':'匹配类型。
	// EQUAL：等于，referer大小写敏感
	// NOT_EQUAL：不等于，referer大小写敏感
	// CONTAIN：包含，referer大小写不敏感
	// NOT_CONTAIN：不包含，referer大小写不敏感
	// REGEX：正则，referer大小写不敏感
	// NONE：为空或不存在'}
	MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
	// {'en':'Referer.
	// When the match type is regex match, only one value is allowed. Example: http://test.com.', 'zh_CN':'Referer。
	// 当匹配类型为正则，则只允许只有一个值。
	// 示例：http://test.com。'}
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

type StatusCodeCondition struct {
	// {'en':'Match type.
	// EQUAL:Equal
	// NOT_EQUAL:Does not equal', 'zh_CN':'匹配类型。
	// EQUAL：等于
	// NOT_EQUAL：不等于'}
	MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
	// {'en':'Response Code.', 'zh_CN':'状态码。'}
	StatusCode []*string `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true" type:"Repeated"`
}

func (s StatusCodeCondition) String() string {
	return tea.Prettify(s)
}

func (s StatusCodeCondition) GoString() string {
	return s.String()
}

func (s *StatusCodeCondition) SetMatchType(v string) *StatusCodeCondition {
	s.MatchType = &v
	return s
}

func (s *StatusCodeCondition) SetStatusCode(v []*string) *StatusCodeCondition {
	s.StatusCode = v
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
	// 支持的值：GET/POST/DELETE/PUT/HEAD/OPTIONS/COPY。'}
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

type RateLimitRuleCondition struct {
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
	// {'en':'URI ParameterI, match type cannot be repeated.
	// When the business scenario is API, this matching condition is not supported.', 'zh_CN':'URI参数，匹配类型不可重复。
	// 当业务场景为网站业务时不支持此匹配条件。'}
	UriParamConditions []*UriParamCondition `json:"uriParamConditions,omitempty" xml:"uriParamConditions,omitempty" type:"Repeated"`
	// {'en':'User Agent, match type cannot be repeated.', 'zh_CN':'User-Agent，匹配类型不可重复。'}
	UaConditions []*UaCondition `json:"uaConditions,omitempty" xml:"uaConditions,omitempty" type:"Repeated"`
	// {'en':'Request Method.
	// When the business scenario is API,this matching condition is not supported.', 'zh_CN':'请求方法，匹配类型不可重复。
	// 当业务场景为API业务时不支持此匹配条件。'}
	MethodConditions []*RequestMethodCondition `json:"methodConditions,omitempty" xml:"methodConditions,omitempty" type:"Repeated"`
	// {'en':'Referer, match type cannot be repeated.', 'zh_CN':'Referer，匹配类型不可重复。'}
	RefererConditions []*RefererCondition `json:"refererConditions,omitempty" xml:"refererConditions,omitempty" type:"Repeated"`
	// {'en':'Request Header, match type can be repeated.', 'zh_CN':'请求头，匹配类型可重复。'}
	HeaderConditions []*HeaderCondition `json:"headerConditions,omitempty" xml:"headerConditions,omitempty" type:"Repeated"`
	// {'en':'Geo,match type cannot be repeated.', 'zh_CN':'区域，匹配类型不可重复。'}
	AreaConditions []*AreaCondition `json:"areaConditions,omitempty" xml:"areaConditions,omitempty" type:"Repeated"`
	// {'en':'Response Code, match type cannot be repeated.', 'zh_CN':'状态码，匹配类型不可重复。'}
	StatusCodeConditions []*StatusCodeCondition `json:"statusCodeConditions,omitempty" xml:"statusCodeConditions,omitempty" type:"Repeated"`
	// {'en':'HTTP/S, match type cannot be repeated.', 'zh_CN':'应用层协议，匹配类型不可重复。'}
	SchemeConditions []*SchemeCondition `json:"schemeConditions,omitempty" xml:"schemeConditions,omitempty" type:"Repeated"`
	// {"en":"JA3 Fingerprint, match type cannot be repeated.","zh_CN":"JA3指纹，匹配类型不可重复。"}
	Ja3Conditions []*Ja3Condition `json:"ja3Conditions,omitempty" xml:"ja3Conditions,omitempty" type:"Repeated"`
	// {"en":"JA4 Fingerprint, match type cannot be repeated.","zh_CN":"JA4指纹，匹配类型不可重复。"}
	Ja4Conditions []*Ja4Condition `json:"ja4Conditions,omitempty" xml:"ja4Conditions,omitempty" type:"Repeated"`
}

func (s RateLimitRuleCondition) String() string {
	return tea.Prettify(s)
}

func (s RateLimitRuleCondition) GoString() string {
	return s.String()
}

func (s *RateLimitRuleCondition) SetIpOrIpsConditions(v []*IpOrIpsCondition) *RateLimitRuleCondition {
	s.IpOrIpsConditions = v
	return s
}

func (s *RateLimitRuleCondition) SetPathConditions(v []*PathCondition) *RateLimitRuleCondition {
	s.PathConditions = v
	return s
}

func (s *RateLimitRuleCondition) SetUriConditions(v []*UriCondition) *RateLimitRuleCondition {
	s.UriConditions = v
	return s
}

func (s *RateLimitRuleCondition) SetUriParamConditions(v []*UriParamCondition) *RateLimitRuleCondition {
	s.UriParamConditions = v
	return s
}

func (s *RateLimitRuleCondition) SetUaConditions(v []*UaCondition) *RateLimitRuleCondition {
	s.UaConditions = v
	return s
}

func (s *RateLimitRuleCondition) SetMethodConditions(v []*RequestMethodCondition) *RateLimitRuleCondition {
	s.MethodConditions = v
	return s
}

func (s *RateLimitRuleCondition) SetRefererConditions(v []*RefererCondition) *RateLimitRuleCondition {
	s.RefererConditions = v
	return s
}

func (s *RateLimitRuleCondition) SetHeaderConditions(v []*HeaderCondition) *RateLimitRuleCondition {
	s.HeaderConditions = v
	return s
}

func (s *RateLimitRuleCondition) SetAreaConditions(v []*AreaCondition) *RateLimitRuleCondition {
	s.AreaConditions = v
	return s
}

func (s *RateLimitRuleCondition) SetStatusCodeConditions(v []*StatusCodeCondition) *RateLimitRuleCondition {
	s.StatusCodeConditions = v
	return s
}

func (s *RateLimitRuleCondition) SetSchemeConditions(v []*SchemeCondition) *RateLimitRuleCondition {
	s.SchemeConditions = v
	return s
}

type UriParamCondition struct {
	// {'en':'Match type.
	// EQUAL:Equals,param value case sensitive
	// NOT_EQUAL:Does not equal,param value case sensitive
	// CONTAIN:Contains,param value case insensitive
	// NOT_CONTAIN:Does not contains,param value case insensitive
	// REGEX:Regex match,param value case insensitive
	// NONE:Empty or non-existent', 'zh_CN':'匹配类型。
	// EQUAL：等于，参数值大小写敏感
	// NOT_EQUAL：不等于，参数值大小写敏感
	// CONTAIN：包含，参数值大小写不敏感
	// NOT_CONTAIN：不包含，参数值大小写不敏感
	// REGEX：正则，参数值大小写不敏感
	// NONE：为空或不存在'}
	MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
	// {'en':'Param name,case sensitive,maximum 100 characters.
	// Example: id.', 'zh_CN':'参数名，大小写敏感，最多100个字符。
	// 示例：id。'}
	ParamName *string `json:"paramName,omitempty" xml:"paramName,omitempty" require:"true"`
	// {'en':'Param value.', 'zh_CN':'参数值。'}
	ParamValue []*string `json:"paramValue,omitempty" xml:"paramValue,omitempty" require:"true" type:"Repeated"`
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

type RateLimitEffective struct {
	// {'en':'Effective.
	// MON:Monday
	// TUE:Tuesday
	// WED:Wednesday
	// THU:Thursday
	// FRI:Friday
	// SAT:Saturday
	// SUN:Sunday', 'zh_CN':'周期。
	// MON：星期一
	// TUE：星期二
	// WED：星期三
	// THU：星期四
	// FRI：星期五
	// SAT：星期六
	// SUN：星期天'}
	Effective []*string `json:"effective,omitempty" xml:"effective,omitempty" require:"true" type:"Repeated"`
	// {'en':'Start time, format: HH:mm.', 'zh_CN':'开始时间，格式：HH:mm。'}
	Start *string `json:"start,omitempty" xml:"start,omitempty" require:"true"`
	// {'en':'End time, format: HH:mm.', 'zh_CN':'结束时间，格式：HH:mm。'}
	End *string `json:"end,omitempty" xml:"end,omitempty" require:"true"`
	// {'en':'Timezone,default value: GTM+8.', 'zh_CN':'时区，默认：GTM+8。','dictionary':'belong=WAAP-MS-Ext|dict=waap_timezone'}
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

type PathCondition struct {
	// {'en':'Match type.
	// EQUAL:Equals,path case sensitive
	// NOT_EQUAL:Does not equal,path case sensitive
	// CONTAIN:Contains,path case insensitive
	// NOT_CONTAIN:Does not contains,path case insensitive
	// REGEX:Regex match,path case insensitive', 'zh_CN':'匹配类型。
	// EQUAL：等于，路径大小写敏感
	// NOT_EQUAL：不等于，路径大小写敏感
	// CONTAIN：包含，路径大小写不敏感
	// NOT_CONTAIN：不包含，路径大小写不敏感
	// REGEX：正则，路径大小写不敏感'}
	MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
	// {'en':'Path.
	// When match type is EQUAL/NOT_EQUAL, path needs to start with "/", and no parameters.
	// When the match type is regex match, only one value is allowed.
	// Example: /test.html.', 'zh_CN':'路径。
	// 当匹配类型为等于/不等于，路径必须以“/”开头，不含参数。
	// 当匹配类型为正则，则只允许只有一个值。
	// 示例：/test.html。'}
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

type SchemeCondition struct {
	// {'en':'Match type.
	// EQUAL:Equal
	// NOT_EQUAL:Does not equal', 'zh_CN':'匹配类型。
	// EQUAL：等于
	// NOT_EQUAL：不等于'}
	MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
	// {'en':'HTTP/S.
	// Supported values: HTTP/HTTPS.', 'zh_CN':'应用层协议。
	// 支持的值：HTTP/HTTPS。'}
	Scheme []*string `json:"scheme,omitempty" xml:"scheme,omitempty" require:"true" type:"Repeated"`
}

func (s SchemeCondition) String() string {
	return tea.Prettify(s)
}

func (s SchemeCondition) GoString() string {
	return s.String()
}

func (s *SchemeCondition) SetMatchType(v string) *SchemeCondition {
	s.MatchType = &v
	return s
}

func (s *SchemeCondition) SetScheme(v []*string) *SchemeCondition {
	s.Scheme = v
	return s
}

type Ja3Condition struct {
	// {"en":"Match type.\nEQUAL: Equals\nNOT_EQUAL: Does not equal","zh_CN":"匹配类型。\nEQUAL：等于 \nNOT_EQUAL：不等于"}
	MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
	// {"en":"JA3 Fingerprint List, maximum 300 JA3 Fingerprint.When the match type is EQUAL/NOT_EQUAL, each item's character length must be 32 and can only include numbers and lowercase letters.","zh_CN":"JA3指纹列表，最多300个JA3指纹。当匹配类型为等于/不等于时，每项字符长度必须为32，且仅限数字和小写字母。"}
	Ja3List []*string `json:"ja3List,omitempty" xml:"ja3List,omitempty" require:"true" type:"Repeated"`
}

func (s Ja3Condition) String() string {
	return tea.Prettify(s)
}

func (s Ja3Condition) GoString() string {
	return s.String()
}

func (s *Ja3Condition) SetMatchType(v string) *Ja3Condition {
	s.MatchType = &v
	return s
}

func (s *Ja3Condition) SetJa3List(v []*string) *Ja3Condition {
	s.Ja3List = v
	return s
}

type Ja4Condition struct {
	// {"en":"JA4 Fingerprint List, maximum 300 JA4 Fingerprint.When the match type is EQUAL/NOT_EQUAL, each item's format must be 10 characters + 12 characters + 12 characters, separated by underscores, and can only include underscores, numbers, and lowercase letters.When the match type is CONTAIN/NOT_CONTAIN/START_WITH/END_WITH, each item is only allowed to include underscores, numbers, and lowercase letters.When the match type is WILDCARD/NOT_WILDCARD, each item, aside from  ** and ?, is only allowed to include underscores, numbers, and lowercase letters.","zh_CN":"JA4指纹列表，最多300个JA4指纹。当匹配类型为等于/不等于时，每项格式必须为10位字符+12位字符+12位字符，中间以“_”分隔，且仅限下划线、数字和小写字母。当匹配类型为包含/不包含/开头是/结尾是时，每项只允许输入下划线、数字和小写字母。当匹配类型为通配符匹配/通配符不匹配时，每项除*和?外，只允许输入下划线、数字和小写字母。"}
	Ja4List []*string `json:"ja4List,omitempty" xml:"ja4List,omitempty" require:"true" type:"Repeated"`
	// {"en":"Match type. \nEQUAL: Equals\nNOT_EQUAL: Does not equal\nCONTAIN: Contains\nNOT_CONTAIN: Does not Contains\nSTART_WITH: Starts with\nEND_WITH: Ends with\nWILDCARD: Wildcard matches, ** represents zero or more arbitrary characters, ? represents any single character\nNOT_WILDCARD: Wildcard does not match, ** represents zero or more arbitrary characters, ? represents any single character","zh_CN":"匹配类型。\nEQUAL：等于\nNOT_EQUAL：不等于\nCONTAIN：包含\nNOT_CONTAIN：不包含\nSTART_WITH：开头是\nEND_WITH：结尾是\nWILDCARD：通配符匹配，*代表零个或多个任意字符，?代表任意单个字符\nNOT_WILDCARD：通配符不匹配，*代表零个或多个任意字符，?代表任意单个字符"}
	MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
}

func (s Ja4Condition) String() string {
	return tea.Prettify(s)
}

func (s Ja4Condition) GoString() string {
	return s.String()
}

func (s *Ja4Condition) SetJa4List(v []*string) *Ja4Condition {
	s.Ja4List = v
	return s
}

func (s *Ja4Condition) SetMatchType(v string) *Ja4Condition {
	s.MatchType = &v
	return s
}
