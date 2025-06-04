package share_customizerule

import "github.com/alibabacloud-go/tea/tea"

type ShareCustomizeRuleCondition struct {
	// {"en":"Request Method.\nWhen the business scenario is API,this matching condition is not supported.","zh_CN":"请求方法，匹配类型不可重复。\n当业务场景为API业务时不支持此匹配条件。"}
	MethodConditions []*MethodCondition `json:"methodConditions,omitempty" xml:"methodConditions,omitempty" type:"Repeated"`
	// {"en":"JA3 Fingerprint, match type cannot be repeated.","zh_CN":"JA3指纹，匹配类型不可重复。"}
	Ja3Conditions []*Ja3Condition `json:"ja3Conditions,omitempty" xml:"ja3Conditions,omitempty" type:"Repeated"`
	// {"en":"Geo, match type cannot be repeated.","zh_CN":"区域，匹配类型不可重复。"}
	AreaConditions []*AreaCondition `json:"areaConditions,omitempty" xml:"areaConditions,omitempty" type:"Repeated"`
	// {"en":"IP/CIDR, match type cannot be repeated.","zh_CN":"IP/IP段，匹配类型不可重复。"}
	IpOrIpsConditions []*IpOrIpsCondition `json:"ipOrIpsConditions,omitempty" xml:"ipOrIpsConditions,omitempty" type:"Repeated"`
	// {"en":"URI, match type cannot be repeated.\nWhen the business scenario is API, this matching condition is not supported.","zh_CN":"URI，匹配类型不可重复。\n当业务场景为API业务时不支持此匹配条件。"}
	UriConditions []*UriCondition `json:"uriConditions,omitempty" xml:"uriConditions,omitempty" type:"Repeated"`
	// {"en":"Path, match type cannot be repeated.\nWhen the business scenario is API, this matching condition is not supported.","zh_CN":"路径，匹配类型不可重复。\n当业务场景为API业务时不支持此匹配条件。"}
	PathConditions []*PathCondition `json:"pathConditions,omitempty" xml:"pathConditions,omitempty" type:"Repeated"`
	// {"en":"User Agent, match type cannot be repeated.","zh_CN":"User-Agent，匹配类型不可重复。"}
	UaConditions []*UaCondition `json:"uaConditions,omitempty" xml:"uaConditions,omitempty" type:"Repeated"`
	// {"en":"Request Header, match type can be repeated.","zh_CN":"请求头，匹配类型可重复。"}
	HeaderConditions []*HeaderCondition `json:"headerConditions,omitempty" xml:"headerConditions,omitempty" type:"Repeated"`
	// {"en":"Referer, match type cannot be repeated.","zh_CN":"Referer，匹配类型不可重复。"}
	RefererConditions []*RefererCondition `json:"refererConditions,omitempty" xml:"refererConditions,omitempty" type:"Repeated"`
	// {"en":"JA4 Fingerprint, match type cannot be repeated.","zh_CN":"JA4指纹，匹配类型不可重复。"}
	Ja4Conditions []*Ja4Condition `json:"ja4Conditions,omitempty" xml:"ja4Conditions,omitempty" type:"Repeated"`
}

func (s ShareCustomizeRuleCondition) String() string {
	return tea.Prettify(s)
}

func (s ShareCustomizeRuleCondition) GoString() string {
	return s.String()
}

func (s *ShareCustomizeRuleCondition) SetMethodConditions(v []*MethodCondition) *ShareCustomizeRuleCondition {
	s.MethodConditions = v
	return s
}

func (s *ShareCustomizeRuleCondition) SetJa3Conditions(v []*Ja3Condition) *ShareCustomizeRuleCondition {
	s.Ja3Conditions = v
	return s
}

func (s *ShareCustomizeRuleCondition) SetAreaConditions(v []*AreaCondition) *ShareCustomizeRuleCondition {
	s.AreaConditions = v
	return s
}

func (s *ShareCustomizeRuleCondition) SetIpOrIpsConditions(v []*IpOrIpsCondition) *ShareCustomizeRuleCondition {
	s.IpOrIpsConditions = v
	return s
}

func (s *ShareCustomizeRuleCondition) SetUriConditions(v []*UriCondition) *ShareCustomizeRuleCondition {
	s.UriConditions = v
	return s
}

func (s *ShareCustomizeRuleCondition) SetPathConditions(v []*PathCondition) *ShareCustomizeRuleCondition {
	s.PathConditions = v
	return s
}

func (s *ShareCustomizeRuleCondition) SetUaConditions(v []*UaCondition) *ShareCustomizeRuleCondition {
	s.UaConditions = v
	return s
}

func (s *ShareCustomizeRuleCondition) SetHeaderConditions(v []*HeaderCondition) *ShareCustomizeRuleCondition {
	s.HeaderConditions = v
	return s
}

func (s *ShareCustomizeRuleCondition) SetRefererConditions(v []*RefererCondition) *ShareCustomizeRuleCondition {
	s.RefererConditions = v
	return s
}

func (s *ShareCustomizeRuleCondition) SetJa4Conditions(v []*Ja4Condition) *ShareCustomizeRuleCondition {
	s.Ja4Conditions = v
	return s
}

type MethodCondition struct {
	// {"en":"Match type.\nEQUAL:Equals\nNOT_EQUAL:Does not equal","zh_CN":"匹配类型。\nEQUAL：等于\nNOT_EQUAL：不等于"}
	MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
	// {"en":"Request method.Supported values: GET/POST/DELETE/PUT/HEAD/OPTIONS/COPY.","zh_CN":"请求方法。支持的值：GET/POST/DELETE/PUT/HEAD/OPTIONS/COPY。"}
	RequestMethod []*string `json:"requestMethod,omitempty" xml:"requestMethod,omitempty" require:"true" type:"Repeated"`
}

func (s MethodCondition) String() string {
	return tea.Prettify(s)
}

func (s MethodCondition) GoString() string {
	return s.String()
}

func (s *MethodCondition) SetMatchType(v string) *MethodCondition {
	s.MatchType = &v
	return s
}

func (s *MethodCondition) SetRequestMethod(v []*string) *MethodCondition {
	s.RequestMethod = v
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

type AreaCondition struct {
	// {"en":"Match type.\nEQUAL:Equals\nNOT_EQUAL:Does not equal","zh_CN":"匹配类型。\nEQUAL：等于\nNOT_EQUAL：不等于"}
	MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
	// {"dictionary":"belong=WAAP-MS-Ext|dict=waap_areaCityAndCountry","en":"Geo.","zh_CN":"区域。"}
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

type IpOrIpsCondition struct {
	// {"en":"Match type.\nEQUAL:Equals\nNOT_EQUAL:Does not equal","zh_CN":"匹配类型。\nEQUAL：等于\nNOT_EQUAL：不等于"}
	MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
	// {"en":"IP/CIDR, maximum 500 IP/CIDR.","zh_CN":"IP/IP段，最多500个IP/IP段。"}
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

type UriCondition struct {
	// {"en":"Match type.\nEQUAL: Equals, URI case sensitive\nNOT_EQUAL: Does not equal, URI case sensitive\nCONTAIN: Contains, URI case insensitive\nNOT_CONTAIN: Does not Contains, URI case insensitive\nREGEX: Regex match, URI case insensitive\nNOT_REGEX: Regular does not match, URI case insensitive\nSTART_WITH: Starts with, URI case insensitive\nEND_WITH: Ends with, URI case insensitive\nWILDCARD: Wildcard matches, URI case insensitive,* represents zero or more arbitrary characters, ? represents any single character\nNOT_WILDCARD: Wildcard does not match, URI case insensitive,* represents zero or more arbitrary characters, ? represents any single character","zh_CN":"匹配类型。\nEQUAL：等于，URI大小写敏感\nNOT_EQUAL：不等于，URI大小写敏感\nCONTAIN：包含，URI大小写不敏感\nNOT_CONTAIN：不包含，URI大小写不敏感\nREGEX：匹配正则，URI大小写不敏感\nNOT_REGEX：正则不匹配，URI大小写不敏感\nSTART_WITH：开头是，URI大小写不敏感\nEND_WITH：结尾是，URI大小写不敏感\nWILDCARD：通配符匹配，URI大小写不敏感，*代表零个或多个任意字符，?代表任意单个字符\nNOT_WILDCARD：通配符不匹配，URI大小写不敏感，*代表零个或多个任意字符，?代表任意单个字符"}
	MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
	// {"en":"URI.When match type is EQUAL/NOT_EQUAL/START_WITH/END_WITH, uri needs to start with \"/\", and includes parameters.When the match type is REGEX/NOT_REGEX, only one value is allowed.Example: /test.html?id=1.","zh_CN":"URI。当匹配类型为等于/不等于/开头是/结尾是，URI必须以”/“开头，含参数。当匹配类型为正则/正则不匹配，则只允许只有一个值。示例：/test.html?id=1。"}
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

type PathCondition struct {
	// {"en":"Match type.\nEQUAL: Equals, path case sensitive\nNOT_EQUAL: Does not equal, path case sensitive\nCONTAIN: Contains, path case insensitive\nNOT_CONTAIN: Does not Contains, path case insensitive\nREGEX: Regex match, path case insensitive\nNOT_REGEX: Regular does not match, path case sensitive\nSTART_WITH: Starts with, path case sensitive\nEND_WITH: Ends with, path case sensitive\nWILDCARD: Wildcard matches, path case sensitive,* represents zero or more arbitrary characters, ? represents any single character.\nNOT_WILDCARD: Wildcard does not match, path case sensitive,* represents zero or more arbitrary characters, ? represents any single character","zh_CN":"匹配类型。\nEQUAL：等于，路径大小写敏感\nNOT_EQUAL：不等于，路径大小写敏感\nCONTAIN：包含，路径大小写不敏感\nNOT_CONTAIN：不包含，路径大小写不敏感\nREGEX：匹配正则，路径大小写不敏感\nNOT_REGEX：正则不匹配，路径大小写不敏感\nSTART_WITH：开头是，路径大小写不敏感\nEND_WITH：结尾是，路径大小写不敏感\nWILDCARD：通配符匹配，路径大小写不敏感，*代表零个或多个任意字符，?代表任意单个字符\nNOT_WILDCARD：通配符不匹配，路径大小写不敏感，*代表零个或多个任意字符，?代表任意单个字符"}
	MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
	// {"en":"Path.When match type is EQUAL/NOT_EQUAL/START_WITH/END_WITH, path needs to start with \"/\", and no parameters.When the match type is REGEX/NOT_REGEX, only one value is allowed.Example: /test.html.","zh_CN":"路径。当匹配类型为等于/不等于/开头是/结尾是，路径必须以“/”开头，不含参数。当匹配类型为正则/正则不匹配，则只允许只有一个值。示例：/test.html。"}
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

type UaCondition struct {
	// {"en":"Match type.\nEQUAL: Equals, user agent case sensitive\nNOT_EQUAL: Does not equal, user agent case sensitive\nCONTAIN: Contains, user agent case insensitive\nNOT_CONTAIN: Does not Contains, user agent case insensitive\nNONE:Empty or non-existent\nREGEX: Regex match, user agent case insensitive\nNOT_REGEX: Wegular does not match, user agent case insensitive\nSTART_WITH: Starts with, user agent case insensitive\nEND_WITH: Ends with, user agent case insensitive\nWILDCARD: Wildcard matches, user agent case insensitive,* represents zero or more arbitrary characters, ? represents any single character\nNOT_WILDCARD: Wildcard does not match, user agent case insensitive,* represents zero or more arbitrary characters, ? represents any single character","zh_CN":"匹配类型。\nEQUAL：等于，User-Agent大小写敏感\nNOT_EQUAL：不等于，User-Agent大小写敏感\nCONTAIN：包含，User-Agent大小写不敏感\nNOT_CONTAIN：不包含，User-Agent大小写不敏感\nNONE：为空或不存在\nREGEX：匹配正则，User-Agent大小写不敏感\nNOT_REGEX：正则不匹配，User-Agent大小写不敏感\nSTART_WITH：开头是，User-Agent大小写不敏感\nEND_WITH：结尾是，User-Agent大小写不敏感\nWILDCARD：通配符匹配，User-Agent大小写不敏感，*代表零个或多个任意字符，?代表任意单个字符\nNOT_WILDCARD：通配符不匹配，User-Agent大小写不敏感，*代表零个或多个任意字符，?代表任意单个字符"}
	MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
	// {"en":"User agent.When the match type is REGEX/NOT_REGEX, only one value is allowed.Example: go-Http-client/1.1.","zh_CN":"User-Agent。当匹配类型为正则/正则不匹配，则只允许只有一个值。示例：go-Http-client/1.1。"}
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

type HeaderCondition struct {
	// {"en":"Match type.\nEQUAL: Equals, request header values case sensitive\nNOT_EQUAL: Does not equal, request header values case sensitive\nCONTAIN: Contains, request header values case insensitive\nNOT_CONTAIN: Does not Contains, request header values case insensitive\nNONE: Empty or non-existent\nREGEX: Regex match, request header values case insensitive\nNOT_REGEX: Regular does not match, request header values case insensitive\nSTART_WITH: Starts with, request header values case insensitive\nEND_WITH: Ends with, request header values case insensitive\nWILDCARD: Wildcard matches, request header values case insensitive,* represents zero or more arbitrary characters, ? represents any single character\nNOT_WILDCARD: Wildcard does not match, request header values case insensitive,* represents zero or more arbitrary characters, ? represents any single character","zh_CN":"匹配类型。\nEQUAL：等于，头部值大小写敏感\nNOT_EQUAL：不等于，头部值大小写敏感\nCONTAIN：包含，头部值大小写不敏感\nNOT_CONTAIN：不包含，头部值大小写不敏感\nNONE：为空或不存在\nREGEX：匹配正则，头部值大小写不敏感\nNOT_REGEX：正则不匹配，头部值大小写不敏感\nSTART_WITH：开头是，头部值大小写不敏感\nEND_WITH：结尾是，头部值大小写不敏感\nWILDCARD：通配符匹配，头部值大小写不敏感，*代表零个或多个任意字符，?代表任意单个字符\nNOT_WILDCARD：通配符不匹配，头部值大小写不敏感，*代表零个或多个任意字符，?代表任意单个字符"}
	MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
	// {"en":"Header value.When the match type is REGEX/NOT_REGEX, only one value is allowed.","zh_CN":"头部值。当匹配类型为正则/正则不匹配，则只允许只有一个值。"}
	ValueList []*string `json:"valueList,omitempty" xml:"valueList,omitempty" require:"true" type:"Repeated"`
	// {"en":"Header name,case insensitive,up to 100 characters.\nExample: Accept.","zh_CN":"头部名称，大小写不敏感，最多100个字符。\n示例：Accept。"}
	Key *string `json:"key,omitempty" xml:"key,omitempty" require:"true"`
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

func (s *HeaderCondition) SetValueList(v []*string) *HeaderCondition {
	s.ValueList = v
	return s
}

func (s *HeaderCondition) SetKey(v string) *HeaderCondition {
	s.Key = &v
	return s
}

type RefererCondition struct {
	// {"en":"Referer.When the match type is REGEX/NOT_REGEX, only one value is allowed.Example: http://test.com.","zh_CN":"Referer。当匹配类型为正则/正则不匹配，则只允许只有一个值。示例：http://test.com。"}
	Referer []*string `json:"referer,omitempty" xml:"referer,omitempty" require:"true" type:"Repeated"`
	// {"en":"Match type.\nEQUAL: Equals, referer case sensitive\nNOT_EQUAL: Does not equal, referer case sensitive\nCONTAIN: Contains, referer case insensitive\nNOT_CONTAIN: Does not Contains, referer case insensitive\nNONE:Empty or non-existent\nREGEX: Regex match, referer case insensitive\nNOT_REGEX: Regular does not match, referer case insensitive\nSTART_WITH: Starts with, referer case insensitive\nEND_WITH: Ends with, referer case insensitive\nWILDCARD: Wildcard matches, referer case insensitive,* represents zero or more arbitrary characters, ? represents any single characte\nNOT_WILDCARD: Wildcard does not match, referer case insensitive,* represents zero or more arbitrary characters, ? represents any single character","zh_CN":"匹配类型。\nEQUAL：等于，referer大小写敏感\nNOT_EQUAL：不等于，referer大小写敏感\nCONTAIN：包含，referer大小写不敏感\nNOT_CONTAIN：不包含，referer大小写不敏感\nNONE：为空或不存在\nREGEX：匹配正则，referer大小写不敏感\nNOT_REGEX：正则不匹配，referer大小写不敏感\nSTART_WITH：开头是，referer大小写不敏感\nEND_WITH：结尾是，referer大小写不敏感\nWILDCARD：通配符匹配，referer大小写不敏感，*代表零个或多个任意字符，?代表任意单个字符\nNOT_WILDCARD：通配符不匹配，referer大小写不敏感，*代表零个或多个任意字符，?代表任意单个字符"}
	MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
}

func (s RefererCondition) String() string {
	return tea.Prettify(s)
}

func (s RefererCondition) GoString() string {
	return s.String()
}

func (s *RefererCondition) SetReferer(v []*string) *RefererCondition {
	s.Referer = v
	return s
}

func (s *RefererCondition) SetMatchType(v string) *RefererCondition {
	s.MatchType = &v
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
