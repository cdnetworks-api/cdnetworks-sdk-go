package predeploy

import "github.com/alibabacloud-go/tea/tea"

type MethodConditions struct {
	// {"en":"Match type.\nEQUAL:Equals\nNOT_EQUAL:Does not equal","zh_CN":"匹配类型。\nEQUAL：等于\nNOT_EQUAL：不等于","exampleValue":"EQUAL,NOT_EQUAL"}
	MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
	// {"en":"Request method.\nSupported values: GET/POST/DELETE/PUT/HEAD/OPTIONS/COPY.","zh_CN":"请求方法。\n支持的值：GET/POST/DELETE/PUT/HEAD/OPTIONS/COPY。","exampleValue":"GET,POST,DELETE,PUT,HEAD,OPTIONS,COPY"}
	RequestMethod []*string `json:"requestMethod,omitempty" xml:"requestMethod,omitempty" require:"true" type:"Repeated"`
}

func (s MethodConditions) String() string {
	return tea.Prettify(s)
}

func (s MethodConditions) GoString() string {
	return s.String()
}

func (s *MethodConditions) SetMatchType(v string) *MethodConditions {
	s.MatchType = &v
	return s
}

func (s *MethodConditions) SetRequestMethod(v []*string) *MethodConditions {
	s.RequestMethod = v
	return s
}

type AreaConditions struct {
	// {"en":"Match type.\nEQUAL:Equals\nNOT_EQUAL:Does not equal","zh_CN":"匹配类型。\nEQUAL：等于\nNOT_EQUAL：不等于","exampleValue":"EQUAL,NOT_EQUAL"}
	MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
	// {"dictionary":"belong=WAAP-MS-Ext|dict=waap_areaCityAndCountry","en":"Geo.","zh_CN":"区域。"}
	Areas []*string `json:"areas,omitempty" xml:"areas,omitempty" require:"true" type:"Repeated"`
}

func (s AreaConditions) String() string {
	return tea.Prettify(s)
}

func (s AreaConditions) GoString() string {
	return s.String()
}

func (s *AreaConditions) SetMatchType(v string) *AreaConditions {
	s.MatchType = &v
	return s
}

func (s *AreaConditions) SetAreas(v []*string) *AreaConditions {
	s.Areas = v
	return s
}

type IpOrIpsConditions struct {
	// {"en":"Match type.\nEQUAL:Equals\nNOT_EQUAL:Does not equal","zh_CN":"匹配类型。\nEQUAL：等于\nNOT_EQUAL：不等于","exampleValue":"EQUAL,NOT_EQUAL"}
	MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
	// {"en":"IP/CIDR, maximum 500 IP/CIDR.","zh_CN":"IP/IP段，最多500个IP/IP段。"}
	IpOrIps []*string `json:"ipOrIps,omitempty" xml:"ipOrIps,omitempty" require:"true" type:"Repeated"`
}

func (s IpOrIpsConditions) String() string {
	return tea.Prettify(s)
}

func (s IpOrIpsConditions) GoString() string {
	return s.String()
}

func (s *IpOrIpsConditions) SetMatchType(v string) *IpOrIpsConditions {
	s.MatchType = &v
	return s
}

func (s *IpOrIpsConditions) SetIpOrIps(v []*string) *IpOrIpsConditions {
	s.IpOrIps = v
	return s
}

type UriConditions struct {
	// {"en":"Match type.\nEQUAL: Equals, URI case sensitive\nNOT_EQUAL: Does not equal, URI case sensitive\nCONTAIN: Contains, URI case insensitive\nNOT_CONTAIN: Does not Contains, URI case insensitive\nREGEX: Regex match, URI case insensitive\nNOT_REGEX: Regular does not match, URI case insensitive\nSTART_WITH: Starts with, URI case insensitive\nEND_WITH: Ends with, URI case insensitive\nWILDCARD: Wildcard matches, URI case insensitive,* represents zero or more arbitrary characters, ? represents any single character\nNOT_WILDCARD: Wildcard does not match, URI case insensitive,* represents zero or more arbitrary characters, ? represents any single character","zh_CN":"匹配类型。\nEQUAL：等于，URI大小写敏感\nNOT_EQUAL：不等于，URI大小写敏感\nCONTAIN：包含，URI大小写不敏感\nNOT_CONTAIN：不包含，URI大小写不敏感\nREGEX：匹配正则，URI大小写不敏感\nNOT_REGEX：正则不匹配，URI大小写不敏感\nSTART_WITH：开头是，URI大小写不敏感\nEND_WITH：结尾是，URI大小写不敏感\nWILDCARD：通配符匹配，URI大小写不敏感，*代表零个或多个任意字符，?代表任意单个字符\nNOT_WILDCARD：通配符不匹配，URI大小写不敏感，*代表零个或多个任意字符，?代表任意单个字符","exampleValue":"EQUAL,NOT_EQUAL,CONTAIN,NOT_CONTAIN,REGEX,NOT_REGEX,START_WITH,END_WITH,WILDCARD,NOT_WILDCARD"}
	MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
	// {"en":"URI.\nWhen match type is EQUAL/NOT_EQUAL/START_WITH/END_WITH, uri needs to start with \"/\", and includes parameters.\nWhen the match type is REGEX/NOT_REGEX, only one value is allowed.\nExample: /test.html?id=1.","zh_CN":"URI。\n当匹配类型为等于/不等于/开头是/结尾是，URI必须以”/“开头，含参数。\n当匹配类型为正则/正则不匹配，则只允许只有一个值。\n示例：/test.html?id=1。"}
	Uri []*string `json:"uri,omitempty" xml:"uri,omitempty" require:"true" type:"Repeated"`
}

func (s UriConditions) String() string {
	return tea.Prettify(s)
}

func (s UriConditions) GoString() string {
	return s.String()
}

func (s *UriConditions) SetMatchType(v string) *UriConditions {
	s.MatchType = &v
	return s
}

func (s *UriConditions) SetUri(v []*string) *UriConditions {
	s.Uri = v
	return s
}

type PathConditions struct {
	// {"en":"Match type.\nEQUAL: Equals, path case sensitive\nNOT_EQUAL: Does not equal, path case sensitive\nCONTAIN: Contains, path case insensitive\nNOT_CONTAIN: Does not Contains, path case insensitive\nREGEX: Regex match, path case insensitive\nNOT_REGEX: Regular does not match, path case sensitive\nSTART_WITH: Starts with, path case sensitive\nEND_WITH: Ends with, path case sensitive\nWILDCARD: Wildcard matches, path case sensitive,* represents zero or more arbitrary characters, ? represents any single character.\nNOT_WILDCARD: Wildcard does not match, path case sensitive,* represents zero or more arbitrary characters, ? represents any single character","zh_CN":"匹配类型。\nEQUAL：等于，路径大小写敏感\nNOT_EQUAL：不等于，路径大小写敏感\nCONTAIN：包含，路径大小写不敏感\nNOT_CONTAIN：不包含，路径大小写不敏感\nREGEX：匹配正则，路径大小写不敏感\nNOT_REGEX：正则不匹配，路径大小写不敏感\nSTART_WITH：开头是，路径大小写不敏感\nEND_WITH：结尾是，路径大小写不敏感\nWILDCARD：通配符匹配，路径大小写不敏感，*代表零个或多个任意字符，?代表任意单个字符\nNOT_WILDCARD：通配符不匹配，路径大小写不敏感，*代表零个或多个任意字符，?代表任意单个字符","exampleValue":"EQUAL,NOT_EQUAL,CONTAIN,NOT_CONTAIN,REGEX,NOT_REGEX,START_WITH,END_WITH,WILDCARD,NOT_WILDCARD"}
	MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
	// {"en":"Path.\nWhen match type is EQUAL/NOT_EQUAL/START_WITH/END_WITH, path needs to start with \"/\", and no parameters.\nWhen the match type is REGEX/NOT_REGEX, only one value is allowed.\nExample: /test.html.","zh_CN":"路径。\n当匹配类型为等于/不等于/开头是/结尾是，路径必须以“/”开头，不含参数。\n当匹配类型为正则/正则不匹配，则只允许只有一个值。\n示例：/test.html。"}
	Paths []*string `json:"paths,omitempty" xml:"paths,omitempty" require:"true" type:"Repeated"`
}

func (s PathConditions) String() string {
	return tea.Prettify(s)
}

func (s PathConditions) GoString() string {
	return s.String()
}

func (s *PathConditions) SetMatchType(v string) *PathConditions {
	s.MatchType = &v
	return s
}

func (s *PathConditions) SetPaths(v []*string) *PathConditions {
	s.Paths = v
	return s
}

type UriParamConditions struct {
	// {"en":"Match type.\nEQUAL:Equals,param value case sensitive\nNOT_EQUAL:Does not equal,param value case sensitive\nCONTAIN:Contains,param value case insensitive\nNOT_CONTAIN:Does not contains,param value case insensitive\nREGEX:Regex match,param value case insensitive\nNONE:Empty or non-existent","zh_CN":"匹配类型。\nEQUAL：等于，参数值大小写敏感\nNOT_EQUAL：不等于，参数值大小写敏感\nCONTAIN：包含，参数值大小写不敏感\nNOT_CONTAIN：不包含，参数值大小写不敏感\nREGEX：正则，参数值大小写不敏感\nNONE：为空或不存在","exampleValue":"EQUAL,NOT_EQUAL,CONTAIN,NOT_CONTAIN,REGEX,NONE"}
	MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
	// {"en":"Param name,case sensitive,maximum 100 characters.\nExample: id.","zh_CN":"参数名，大小写敏感，最多100个字符。\n示例：id。"}
	ParamName *string `json:"paramName,omitempty" xml:"paramName,omitempty" require:"true"`
	// {"en":"Param value.","zh_CN":"参数值。"}
	ParamValue []*string `json:"paramValue,omitempty" xml:"paramValue,omitempty" require:"true" type:"Repeated"`
}

func (s UriParamConditions) String() string {
	return tea.Prettify(s)
}

func (s UriParamConditions) GoString() string {
	return s.String()
}

func (s *UriParamConditions) SetMatchType(v string) *UriParamConditions {
	s.MatchType = &v
	return s
}

func (s *UriParamConditions) SetParamName(v string) *UriParamConditions {
	s.ParamName = &v
	return s
}

func (s *UriParamConditions) SetParamValue(v []*string) *UriParamConditions {
	s.ParamValue = v
	return s
}

type UaConditions struct {
	// {"en":"Match type.\nEQUAL: Equals, user agent case sensitive\nNOT_EQUAL: Does not equal, user agent case sensitive\nCONTAIN: Contains, user agent case insensitive\nNOT_CONTAIN: Does not Contains, user agent case insensitive\nNONE:Empty or non-existent\nREGEX: Regex match, user agent case insensitive\nNOT_REGEX: Regular does not match, user agent case insensitive\nSTART_WITH: Starts with, user agent case insensitive\nEND_WITH: Ends with, user agent case insensitive\nWILDCARD: Wildcard matches, user agent case insensitive,* represents zero or more arbitrary characters, ? represents any single character\nNOT_WILDCARD: Wildcard does not match, user agent case insensitive,* represents zero or more arbitrary characters, ? represents any single character","zh_CN":"匹配类型。\nEQUAL：等于，User-Agent大小写敏感\nNOT_EQUAL：不等于，User-Agent大小写敏感\nCONTAIN：包含，User-Agent大小写不敏感\nNOT_CONTAIN：不包含，User-Agent大小写不敏感\nNONE：为空或不存在\nREGEX：匹配正则，User-Agent大小写不敏感\nNOT_REGEX：正则不匹配，User-Agent大小写不敏感\nSTART_WITH：开头是，User-Agent大小写不敏感\nEND_WITH：结尾是，User-Agent大小写不敏感\nWILDCARD：通配符匹配，User-Agent大小写不敏感，*代表零个或多个任意字符，?代表任意单个字符\nNOT_WILDCARD：通配符不匹配，User-Agent大小写不敏感，*代表零个或多个任意字符，?代表任意单个字符","exampleValue":"EQUAL,NOT_EQUAL,CONTAIN,NOT_CONTAIN,NONE,REGEX,NOT_REGEX,START_WITH,END_WITH,WILDCARD,NOT_WILDCARD"}
	MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
	// {"en":"User agent.\nWhen the match type is REGEX/NOT_REGEX, only one value is allowed.\nExample: go-Http-client/1.1.","zh_CN":"User-Agent。\n当匹配类型为正则/正则不匹配，则只允许只有一个值。\n示例：go-Http-client/1.1。"}
	Ua []*string `json:"ua,omitempty" xml:"ua,omitempty" require:"true" type:"Repeated"`
}

func (s UaConditions) String() string {
	return tea.Prettify(s)
}

func (s UaConditions) GoString() string {
	return s.String()
}

func (s *UaConditions) SetMatchType(v string) *UaConditions {
	s.MatchType = &v
	return s
}

func (s *UaConditions) SetUa(v []*string) *UaConditions {
	s.Ua = v
	return s
}

type HeaderConditions struct {
	// {"en":"Match type.\nEQUAL: Equals, request header values case sensitive.\nNOT_EQUAL: Does not equal, request header values case sensitive.\nCONTAIN: Contains, request header values case insensitive.\nNOT_CONTAIN: Does not Contains, request header values case insensitive.\nNONE: Empty or non-existent.\nREGEX: Regex match, request header values case insensitive.\nNOT_REGEX: Regular does not match, request header values case insensitive.\nSTART_WITH: Starts with, request header values case insensitive\nEND_WITH: Ends with, request header values case insensitive.\nWILDCARD: Wildcard matches, request header values case insensitive,* represents zero or more arbitrary characters, ? represents any single character.\nNOT_WILDCARD: Wildcard does not match, request header values case insensitive,* represents zero or more arbitrary characters, ? represents any single character.","zh_CN":"匹配类型。\nEQUAL：等于，头部值大小写敏感。\nNOT_EQUAL：不等于，头部值大小写敏感。\nCONTAIN：包含，头部值大小写不敏感。\nNOT_CONTAIN：不包含，头部值大小写不敏感。\nNONE：为空或不存在REGEX：匹配正则，头部值大小写不敏感。\nNOT_REGEX：正则不匹配，头部值大小写不敏感。\nSTART_WITH：开头是，头部值大小写不敏感。\nEND_WITH：结尾是，头部值大小写不敏感。\nWILDCARD：通配符匹配，头部值大小写不敏感，*代表零个或多个任意字符，?代表任意单个字符。\nNOT_WILDCARD：通配符不匹配，头部值大小写不敏感，*代表零个或多个任意字符，?代表任意单个字符。","exampleValue":"EQUAL,NOT_EQUAL,CONTAIN,NOT_CONTAIN,NONE,REGEX,NOT_REGEX,START_WITH,END_WITH,WILDCARD,NOT_WILDCARD"}
	MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
	// {"en":"Header value.\nWhen the match type is REGEX/NOT_REGEX, only one value is allowed.","zh_CN":"头部值。\n当匹配类型为正则/正则不匹配，则只允许只有一个值。"}
	ValueList []*string `json:"valueList,omitempty" xml:"valueList,omitempty" require:"true" type:"Repeated"`
	// {"en":"Header name,case insensitive,up to 100 characters.\nExample: Accept.","zh_CN":"头部名称，大小写不敏感，最多100个字符。\n示例：Accept。"}
	Key *string `json:"key,omitempty" xml:"key,omitempty" require:"true"`
}

func (s HeaderConditions) String() string {
	return tea.Prettify(s)
}

func (s HeaderConditions) GoString() string {
	return s.String()
}

func (s *HeaderConditions) SetMatchType(v string) *HeaderConditions {
	s.MatchType = &v
	return s
}

func (s *HeaderConditions) SetValueList(v []*string) *HeaderConditions {
	s.ValueList = v
	return s
}

func (s *HeaderConditions) SetKey(v string) *HeaderConditions {
	s.Key = &v
	return s
}

type RefererConditions struct {
	// {"en":"Referer.\nWhen the match type is REGEX/NOT_REGEX, only one value is allowed.\nExample: http://test.com.","zh_CN":"Referer。\n当匹配类型为正则/正则不匹配，则只允许只有一个值。\n示例：http://test.com。"}
	Referer []*string `json:"referer,omitempty" xml:"referer,omitempty" require:"true" type:"Repeated"`
	// {"en":"Match type.\nEQUAL: Equals, referer case sensitive\nNOT_EQUAL: Does not equal, referer case sensitive\nCONTAIN: Contains, referer case insensitive\nNOT_CONTAIN: Does not Contains, referer case insensitive\nNONE:Empty or non-existent\nREGEX: Regex match, referer case insensitive\nNOT_REGEX: Regular does not match, referer case insensitive\nSTART_WITH: Starts with, referer case insensitive\nEND_WITH: Ends with, referer case insensitive\nWILDCARD: Wildcard matches, referer case insensitive,* represents zero or more arbitrary characters, ? represents any single characte\nNOT_WILDCARD: Wildcard does not match, referer case insensitive,* represents zero or more arbitrary characters, ? represents any single character","zh_CN":"匹配类型。\nEQUAL：等于，referer大小写敏感\nNOT_EQUAL：不等于，referer大小写敏感\nCONTAIN：包含，referer大小写不敏感\nNOT_CONTAIN：不包含，referer大小写不敏感\nNONE：为空或不存在\nREGEX：匹配正则，referer大小写不敏感\nNOT_REGEX：正则不匹配，referer大小写不敏感\nSTART_WITH：开头是，referer大小写不敏感\nEND_WITH：结尾是，referer大小写不敏感\nWILDCARD：通配符匹配，referer大小写不敏感，*代表零个或多个任意字符，?代表任意单个字符\nNOT_WILDCARD：通配符不匹配，referer大小写不敏感，*代表零个或多个任意字符，?代表任意单个字符","exampleValue":"EQUAL,NOT_EQUAL,CONTAIN,NOT_CONTAIN,NONE,REGEX,NOT_REGEX,START_WITH,END_WITH,WILDCARD,NOT_WILDCARD"}
	MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
}

func (s RefererConditions) String() string {
	return tea.Prettify(s)
}

func (s RefererConditions) GoString() string {
	return s.String()
}

func (s *RefererConditions) SetReferer(v []*string) *RefererConditions {
	s.Referer = v
	return s
}

func (s *RefererConditions) SetMatchType(v string) *RefererConditions {
	s.MatchType = &v
	return s
}

type SchemeConditions struct {
	// {"en":"Match type.\nEQUAL:Equals\nNOT_EQUAL:Does not equal","zh_CN":"匹配类型。\nEQUAL：等于\nNOT_EQUAL：不等于","exampleValue":"EQUAL,NOT_EQUAL"}
	MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
	// {"en":"HTTP/S.\nSupported values: HTTP/HTTPS.","zh_CN":"应用层协议。\n支持的值：HTTP/HTTPS。","exampleValue":"HTTP,HTTPS"}
	Scheme []*string `json:"scheme,omitempty" xml:"scheme,omitempty" require:"true" type:"Repeated"`
}

func (s SchemeConditions) String() string {
	return tea.Prettify(s)
}

func (s SchemeConditions) GoString() string {
	return s.String()
}

func (s *SchemeConditions) SetMatchType(v string) *SchemeConditions {
	s.MatchType = &v
	return s
}

func (s *SchemeConditions) SetScheme(v []*string) *SchemeConditions {
	s.Scheme = v
	return s
}

type StatusCodeConditions struct {
	// {"en":"Match type.\nEQUAL:Equals\nNOT_EQUAL:Does not equal","zh_CN":"匹配类型。\nEQUAL：等于\nNOT_EQUAL：不等于","exampleValue":"EQUAL,NOT_EQUAL"}
	MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
	// {"en":"Response Code.","zh_CN":"状态码。"}
	StatusCode []*string `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true" type:"Repeated"`
}

func (s StatusCodeConditions) String() string {
	return tea.Prettify(s)
}

func (s StatusCodeConditions) GoString() string {
	return s.String()
}

func (s *StatusCodeConditions) SetMatchType(v string) *StatusCodeConditions {
	s.MatchType = &v
	return s
}

func (s *StatusCodeConditions) SetStatusCode(v []*string) *StatusCodeConditions {
	s.StatusCode = v
	return s
}

type Ja3Conditions struct {
	// {"en":"Match type.\nEQUAL: Equals\nNOT_EQUAL: Does not equal","zh_CN":"匹配类型。\nEQUAL：等于 \nNOT_EQUAL：不等于","exampleValue":"EQUAL,NOT_EQUAL"}
	MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
	// {"en":"JA3 Fingerprint List, maximum 300 JA3 Fingerprint.\nWhen the match type is EQUAL/NOT_EQUAL, each item's character length must be 32 and can only include numbers and lowercase letters.","zh_CN":"JA3指纹列表，最多300个JA3指纹。\n当匹配类型为等于/不等于时，每项字符长度必须为32，且仅限数字和小写字母。"}
	Ja3List []*string `json:"ja3List,omitempty" xml:"ja3List,omitempty" require:"true" type:"Repeated"`
}

func (s Ja3Conditions) String() string {
	return tea.Prettify(s)
}

func (s Ja3Conditions) GoString() string {
	return s.String()
}

func (s *Ja3Conditions) SetMatchType(v string) *Ja3Conditions {
	s.MatchType = &v
	return s
}

func (s *Ja3Conditions) SetJa3List(v []*string) *Ja3Conditions {
	s.Ja3List = v
	return s
}

type Ja4Conditions struct {
	// {"en":"Match type. \nEQUAL: Equals\nNOT_EQUAL: Does not equal\nCONTAIN: Contains\nNOT_CONTAIN: Does not Contains\nSTART_WITH: Starts with\nEND_WITH: Ends with\nWILDCARD: Wildcard matches, ** represents zero or more arbitrary characters, ? represents any single character\nNOT_WILDCARD: Wildcard does not match, ** represents zero or more arbitrary characters, ? represents any single character","zh_CN":"匹配类型。\nEQUAL：等于\nNOT_EQUAL：不等于\nCONTAIN：包含\nNOT_CONTAIN：不包含\nSTART_WITH：开头是\nEND_WITH：结尾是\nWILDCARD：通配符匹配，*代表零个或多个任意字符，?代表任意单个字符\nNOT_WILDCARD：通配符不匹配，*代表零个或多个任意字符，?代表任意单个字符","exampleValue":"EQUAL,NOT_EQUAL,CONTAIN,NOT_CONTAIN,START_WITH,END_WITH,WILDCARD,NOT_WILDCARD"}
	MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
	// {"en":"JA4 Fingerprint List, maximum 300 JA4 Fingerprint.\nWhen the match type is EQUAL/NOT_EQUAL, each item's format must be 10 characters + 12 characters + 12 characters, separated by underscores, and can only include underscores, numbers, and lowercase letters.\nWhen the match type is CONTAIN/NOT_CONTAIN/START_WITH/END_WITH, each item is only allowed to include underscores, numbers, and lowercase letters.\nWhen the match type is WILDCARD/NOT_WILDCARD, each item, aside from  ** and ?, is only allowed to include underscores, numbers, and lowercase letters.","zh_CN":"JA4指纹列表，最多300个JA4指纹。\n当匹配类型为等于/不等于时，每项格式必须为10位字符+12位字符+12位字符，中间以“_”分隔，且仅限下划线、数字和小写字母。\n当匹配类型为包含/不包含/开头是/结尾是时，每项只允许输入下划线、数字和小写字母。\n当匹配类型为通配符匹配/通配符不匹配时，每项除*和?外，只允许输入下划线、数字和小写字母。"}
	Ja4List []*string `json:"ja4List,omitempty" xml:"ja4List,omitempty" require:"true" type:"Repeated"`
}

func (s Ja4Conditions) String() string {
	return tea.Prettify(s)
}

func (s Ja4Conditions) GoString() string {
	return s.String()
}

func (s *Ja4Conditions) SetMatchType(v string) *Ja4Conditions {
	s.MatchType = &v
	return s
}

func (s *Ja4Conditions) SetJa4List(v []*string) *Ja4Conditions {
	s.Ja4List = v
	return s
}
