package share_whitelist

import "github.com/alibabacloud-go/tea/tea"

type ShareWhitelistRuleConditions struct {
	// {"en":"IP/CIDR match conditions, match type cannot be repeated.","zh_CN":"IP/IP段匹配条件，匹配类型不可重复。"}
	IpOrIpsConditions []*IpOrIpsCondition `json:"ipOrIpsConditions,omitempty" xml:"ipOrIpsConditions,omitempty" type:"Repeated"`
	// {"en":"URI match conditions, match type cannot be repeated.","zh_CN":"URI匹配条件，匹配类型不可重复。"}
	UriConditions []*UriCondition `json:"uriConditions,omitempty" xml:"uriConditions,omitempty" type:"Repeated"`
	// {"en":"Path match conditions, match type cannot be repeated.","zh_CN":"路径匹配条件，匹配类型不可重复。"}
	PathConditions []*PathCondition `json:"pathConditions,omitempty" xml:"pathConditions,omitempty" type:"Repeated"`
	// {"en":"User agent match conditions, match type cannot be repeated.","zh_CN":"User-Agent 匹配条件，匹配类型不可重复。"}
	UaConditions []*UaCondition `json:"uaConditions,omitempty" xml:"uaConditions,omitempty" type:"Repeated"`
	// {"en":"Request header match conditions.","zh_CN":"请求头匹配条件。"}
	HeaderConditions []*HeaderCondition `json:"headerConditions,omitempty" xml:"headerConditions,omitempty" type:"Repeated"`
	// {"en":"Referer match conditions, match type cannot be repeated.","zh_CN":"Referer 匹配条件，匹配类型不可重复。"}
	RefererConditions []*RefererCondition `json:"refererConditions,omitempty" xml:"refererConditions,omitempty" type:"Repeated"`
}

func (s ShareWhitelistRuleConditions) String() string {
	return tea.Prettify(s)
}

func (s ShareWhitelistRuleConditions) GoString() string {
	return s.String()
}

func (s *ShareWhitelistRuleConditions) SetIpOrIpsConditions(v []*IpOrIpsCondition) *ShareWhitelistRuleConditions {
	s.IpOrIpsConditions = v
	return s
}

func (s *ShareWhitelistRuleConditions) SetUriConditions(v []*UriCondition) *ShareWhitelistRuleConditions {
	s.UriConditions = v
	return s
}

func (s *ShareWhitelistRuleConditions) SetPathConditions(v []*PathCondition) *ShareWhitelistRuleConditions {
	s.PathConditions = v
	return s
}

func (s *ShareWhitelistRuleConditions) SetUaConditions(v []*UaCondition) *ShareWhitelistRuleConditions {
	s.UaConditions = v
	return s
}

func (s *ShareWhitelistRuleConditions) SetHeaderConditions(v []*HeaderCondition) *ShareWhitelistRuleConditions {
	s.HeaderConditions = v
	return s
}

func (s *ShareWhitelistRuleConditions) SetRefererConditions(v []*RefererCondition) *ShareWhitelistRuleConditions {
	s.RefererConditions = v
	return s
}

type IpOrIpsCondition struct {
	// {"en":"Match type.\nEQUAL: Equals\nNOT_EQUAL: Does not equal","zh_CN":"匹配类型。\nEQUAL：等于\nNOT_EQUAL：不等于"}
	MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
	// {"en":"","zh_CN":""}
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
	// {"en":"","zh_CN":""}
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
	// {"en":"","zh_CN":""}
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
	// {"en":"Match type.\nEQUAL: Equals, user agent case sensitive\nNOT_EQUAL: Does not equal, user agent case sensitive\nCONTAIN: Contains, user agent case insensitive\nNOT_CONTAIN: Does not Contains, user agent case insensitive\nREGEX: Regex match, user agent case insensitive\nNOT_REGEX: Regular does not match, user agent case insensitive\nSTART_WITH: Starts with, user agent case insensitive\nEND_WITH: Ends with, user agent case insensitive\nWILDCARD: Wildcard matches, user agent case insensitive,* represents zero or more arbitrary characters, ? represents any single character\nNOT_WILDCARD: Wildcard does not match, user agent case insensitive,* represents zero or more arbitrary characters, ? represents any single character","zh_CN":"匹配类型。\nEQUAL：等于，User-Agent大小写敏感\nNOT_EQUAL：不等于，User-Agent大小写敏感\nCONTAIN：包含，User-Agent大小写不敏感\nNOT_CONTAIN：不包含，User-Agent大小写不敏感\nREGEX：匹配正则，User-Agent大小写不敏感\nNOT_REGEX：正则不匹配，User-Agent大小写不敏感\nSTART_WITH：开头是，User-Agent大小写不敏感\nEND_WITH：结尾是，User-Agent大小写不敏感\nWILDCARD：通配符匹配，User-Agent大小写不敏感，*代表零个或多个任意字符，?代表任意单个字符\nNOT_WILDCARD：通配符不匹配，User-Agent大小写不敏感，*代表零个或多个任意字符，?代表任意单个字符"}
	MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
	// {"en":"","zh_CN":""}
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
	// {"en":"Match type.\nEQUAL: Equals, request header values case sensitive\nNOT_EQUAL: Does not equal, request header values case sensitive\nCONTAIN: Contains, request header values case insensitive\nNOT_CONTAIN: Does not Contains, request header values case insensitive\nREGEX: Regex match, request header values case insensitive\nNOT_REGEX: Regular does not match, request header values case insensitive\nSTART_WITH: Starts with, request header values case insensitive\nEND_WITH: Ends with, request header values case insensitive\nWILDCARD: Wildcard matches, request header values case insensitive,* represents zero or more arbitrary characters, ? represents any single character\nNOT_WILDCARD: Wildcard does not match, request header values case insensitive,* represents zero or more arbitrary characters, ? represents any single character","zh_CN":"匹配类型。\nEQUAL：等于，头部值大小写敏感\nNOT_EQUAL：不等于，头部值大小写敏感\nCONTAIN：包含，头部值大小写不敏感\nNOT_CONTAIN：不包含，头部值大小写不敏感\nREGEX：匹配正则，头部值大小写不敏感\nNOT_REGEX：正则不匹配，头部值大小写不敏感\nSTART_WITH：开头是，头部值大小写不敏感\nEND_WITH：结尾是，头部值大小写不敏感\nWILDCARD：通配符匹配，头部值大小写不敏感，*代表零个或多个任意字符，?代表任意单个字符\nNOT_WILDCARD：通配符不匹配，头部值大小写不敏感，*代表零个或多个任意字符，?代表任意单个字符"}
	MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
	// {"en":"","zh_CN":""}
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
	// {"en":"","zh_CN":""}
	Referer []*string `json:"referer,omitempty" xml:"referer,omitempty" require:"true" type:"Repeated"`
	// {"en":"Match type.\nEQUAL: Equals, referer case sensitive\nNOT_EQUAL: Does not equal, referer case sensitive\nCONTAIN: Contains, referer case insensitive\nNOT_CONTAIN: Does not Contains, referer case insensitive\nREGEX: Regex match, referer case insensitive\nNOT_REGEX: Regular does not match, referer case insensitive\nSTART_WITH: Starts with, referer case insensitive\nEND_WITH: Ends with, referer case insensitive\nWILDCARD: Wildcard matches, referer case insensitive,* represents zero or more arbitrary characters, ? represents any single characte\nNOT_WILDCARD: Wildcard does not match, referer case insensitive,* represents zero or more arbitrary characters, ? represents any single character","zh_CN":"匹配类型。\nEQUAL：等于，referer大小写敏感\nNOT_EQUAL：不等于，referer大小写敏感\nCONTAIN：包含，referer大小写不敏感\nNOT_CONTAIN：不包含，referer大小写不敏感\nREGEX：匹配正则，referer大小写不敏感\nNOT_REGEX：正则不匹配，referer大小写不敏感\nSTART_WITH：开头是，referer大小写不敏感\nEND_WITH：结尾是，referer大小写不敏感\nWILDCARD：通配符匹配，referer大小写不敏感，*代表零个或多个任意字符，?代表任意单个字符\nNOT_WILDCARD：通配符不匹配，referer大小写不敏感，*代表零个或多个任意字符，?代表任意单个字符"}
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
