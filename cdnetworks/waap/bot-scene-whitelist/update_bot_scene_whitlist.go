// This file is auto-generated, don't edit it. Thanks.
package bot_scene_whitelist

import (
	"github.com/alibabacloud-go/tea/tea"
)

type UpdateSpecificClientTrafficBypassRequest struct {
	// {"en":"Whitelist ID.","zh_CN":"白名单ID。"}
	Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
	// {"en":"Hostname.","zh_CN":"域名。"}
	Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
	// {"en":"Whitelist name, does not support # and & , maximum 50 characters.","zh_CN":"白名单名称，不支持 # 和 & ，最多50个字符。"}
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// {"en":"Description, maximum 200 characters.","zh_CN":"描述，最多200个字符。"}
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// {"en":"List of matching conditions.","zh_CN":"匹配条件列表。"}
	Conditions []*UpdateSpecificClientTrafficBypassRequestConditions `json:"conditions,omitempty" xml:"conditions,omitempty" type:"Repeated"`
}

func (s UpdateSpecificClientTrafficBypassRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateSpecificClientTrafficBypassRequest) GoString() string {
	return s.String()
}

func (s *UpdateSpecificClientTrafficBypassRequest) SetId(v string) *UpdateSpecificClientTrafficBypassRequest {
	s.Id = &v
	return s
}

func (s *UpdateSpecificClientTrafficBypassRequest) SetDomain(v string) *UpdateSpecificClientTrafficBypassRequest {
	s.Domain = &v
	return s
}

func (s *UpdateSpecificClientTrafficBypassRequest) SetName(v string) *UpdateSpecificClientTrafficBypassRequest {
	s.Name = &v
	return s
}

func (s *UpdateSpecificClientTrafficBypassRequest) SetDescription(v string) *UpdateSpecificClientTrafficBypassRequest {
	s.Description = &v
	return s
}

func (s *UpdateSpecificClientTrafficBypassRequest) SetConditions(v []*UpdateSpecificClientTrafficBypassRequestConditions) *UpdateSpecificClientTrafficBypassRequest {
	s.Conditions = v
	return s
}

type UpdateSpecificClientTrafficBypassRequestConditions struct {
	// {"en":"Matching condition name.\nIP_IPS: IP/CIDR\nPATH: Path\nURI: Path with parameters\nHEADER: Request Header\nUA: User Agent\nREQUEST_METHOD: Request Method\nREFERER: Referer","zh_CN":"匹配条件名称。\nIP_IPS：IP/IP段\nPATH：路径\nURI：带参数路径\nHEADER：请求头\nUA: User-Agent\nREQUEST_METHOD: 请求方法\nREFERER: Referer"}
	MatchName *string `json:"matchName,omitempty" xml:"matchName,omitempty" require:"true"`
	// {"en":"When matchName is IP_IPS, maximum 300 IP/CIDR in match value list, the optional value of matchType is:\nEQUAL: Equals\nNOT_EQUAL: Does not equal\nWhen matchName is a URI, the optional value of matchType is:\nEQUAL: Equals, the matching value is case-sensitive and needs to start with \"/\" and include parameters\nNOT_EQUAL: Does not equal, the matching value is case-sensitive, needs to start with \"/\", and contains parameters\nCONTAIN: Contains, match values are case insensitive\nNOT_CONTAIN: Does not contains, match values are case insensitive\nREGEX: Regex match, , match values are case insensitive and only one value is allowed\nNOT_REGEX: Regular does not match, match values are case insensitive and only one value is allowed\nSTART_WITH: Starts with, case insensitive\nEND_WITH: Ends with, case insensitive\nWILDCARD: Wildcard matches, case insensitive,* represents zero or more arbitrary characters, ? represents any single character\nNOT_WILDCARD: Wildcard does not match, URI case insensitive,* represents zero or more arbitrary characters, ? represents any single character\nWhen matchName is PATH, the optional value of matchType is:\nEQUAL: Equals, the matching value is case-sensitive and needs to start with \"/\" , and does not contain parameters\nNOT_EQUAL: Does not equal, the matching value is case-sensitive, needs to start with \"/\", and does not contain parameters\nCONTAIN: Contains, match values are case insensitive\nNOT_CONTAIN: Does not contains, match values are case insensitive\nREGEX: Regex match,  match values are case insensitive and only one value is allowed\nNOT_REGEX: Regular does not match, match values are case insensitive and only one value is allowed\nSTART_WITH: Starts with, match values are case insensitive\nEND_WITH: Ends with, match values are case insensitive\nWILDCARD: Wildcard matches, match values are case insensitive,* represents zero or more arbitrary characters, ? represents any single character\nNOT_WILDCARD: Wildcard does not match, match values are case insensitive,* represents zero or more arbitrary characters, ? represents any single character\nWhen matchName is HEADER, the optional value of matchType is:\nEQUAL: Equals, match values are case sensitive\nNOT_EQUAL: Does not equal, the matching value is case-sensitive\nCONTAIN: Contains, match values are case insensitive\nNOT_CONTAIN: Does not contains, match values are case insensitive\nREGEX: Regex match, match values are case insensitive and only one value is allowed\nNONE: Empty or does not exist\nNOT_REGEX: Regular does not match, match values are case insensitive and only one value is allowed\nSTART_WITH: Starts with, match values are case insensitive\nEND_WITH: Ends with, match values are case insensitive\nWILDCARD: Wildcard matches, match values are case insensitive,* represents zero or more arbitrary characters, ? represents any single character\nNOT_WILDCARD: Wildcard does not match, match values are case insensitive,* represents zero or more arbitrary characters, ? represents any single character\nWhen matchName is UA, the optional value of matchType is:\nEQUAL: Equals, match values are case sensitive\nNOT_EQUAL: Does not equal, the matching value is case-sensitive\nCONTAIN: Contains, match values are case insensitive\nNOT_CONTAIN: Does not contains, match values are case insensitive\nREGEX: Regex match, match values are case insensitive and only one value is allowed\nNONE: Empty or does not exist\nNOT_REGEX: Regular does not match, match values are case insensitive and only one value is allowed\nSTART_WITH: Starts with, match values are case insensitive\nEND_WITH: Ends with, match values are case insensitive\nWILDCARD: Wildcard matches, match values are case insensitive,* represents zero or more arbitrary characters, ? represents any single character\nNOT_WILDCARD: Wildcard does not match, match values are case insensitive,* represents zero or more arbitrary characters, ? represents any single character\nWhen matchName is REFERER, the optional value of matchType is:\nEQUAL: Equals, match values are case sensitive\nNOT_EQUAL: Does not equal, the matching value is case-sensitive\nCONTAIN: Contains, match values are case insensitive\nNOT_CONTAIN: Does not contains, match values are case insensitive\nREGEX: Regex match, match values are case insensitive and only one value is allowed\nNONE: Empty or does not exist\nNOT_REGEX: Regular does not match, match values are case insensitive and only one value is allowed\nSTART_WITH: Starts with, match values are case insensitive\nEND_WITH: Ends with, match values are case insensitive\nWILDCARD: Wildcard matches, match values are case insensitive,* represents zero or more arbitrary characters, ? represents any single character\nNOT_WILDCARD: Wildcard does not match, match values are case insensitive,* represents zero or more arbitrary characters, ? represents any single character\nWhen matchName is REQUEST_METHOD, the optional value of matchType is:\nEQUAL: Equals\nNOT_EQUAL: Does not equal","zh_CN":"当matchName是IP_IPS时，匹配值列表中最多300个IP或IP段，matchType可选值为：\nEQUAL：等于\nNOT_EQUAL：不等于\n当matchName是URI时，matchType可选值为：\nEQUAL：等于，匹配值大小写敏感，需要以\"/\"开头，含参数\nNOT_EQUAL：不等于，匹配值大小写敏感，需要以\"/\"开头，含参数\nCONTAIN：包含，匹配值大小写不敏感\nNOT_CONTAIN：不包含，匹配值大小写不敏感\nREGEX：匹配正则，大小写不敏感，匹配值只允许有一个值\nNOT_REGEX：正则不匹配，大小写不敏感，匹配值只允许有一个值\nSTART_WITH：开头是，大小写不敏感\nEND_WITH：结尾是，大小写不敏感\nWILDCARD：通配符匹配，大小写不敏感，*代表零个或多个任意字符，?代表任意单个字符\nNOT_WILDCARD：通配符不匹配，大小写不敏感，*代表零个或多个任意字符，?代表任意单个字符\n当matchName是PATH时，matchType可选值为：\nEQUAL：等于，匹配值大小写敏感，需要以\"/\"开头，不含参数\nNOT_EQUAL：不等于，匹配值大小写敏感，需要以\"/\"开头，不含参数\nCONTAIN：包含，匹配值大小写不敏感\nNOT_CONTAIN：不包含，匹配值大小写不敏感\nREGEX：匹配正则，大小写不敏感，匹配值大小写不敏感，只允许有一个值\nNOT_REGEX：正则不匹配，大小写不敏感，匹配值只允许有一个值\nSTART_WITH：开头是，大小写不敏感\nEND_WITH：结尾是，大小写不敏感\nWILDCARD：通配符匹配，大小写不敏感，*代表零个或多个任意字符，?代表任意单个字符\nNOT_WILDCARD：通配符不匹配，大小写不敏感，*代表零个或多个任意字符，?代表任意单个字符\n当matchName是HEADER时，matchType可选值为：\nEQUAL：等于，匹配值大小写敏感\nNOT_EQUAL：不等于，匹配值大小写敏感\nCONTAIN：包含，匹配值大小写不敏感\nNOT_CONTAIN：不包含，匹配值大小写不敏感\nREGEX：匹配正则，大小写不敏感，匹配值大小写不敏感，只允许有一个值\nNONE：为空或不存在\nNOT_REGEX：正则不匹配，大小写不敏感，匹配值只允许有一个值\nSTART_WITH：开头是，大小写不敏感\nEND_WITH：结尾是，大小写不敏感\nWILDCARD：通配符匹配，大小写不敏感，*代表零个或多个任意字符，?代表任意单个字符\nNOT_WILDCARD：通配符不匹配，大小写不敏感，*代表零个或多个任意字符，?代表任意单个字符\n当matchName是UA时，matchType可选值为：\nEQUAL：等于，匹配值大小写敏感\nNOT_EQUAL：不等于，匹配值大小写敏感\nCONTAIN：包含，匹配值大小写不敏感\nNOT_CONTAIN：不包含，匹配值大小写不敏感\nREGEX：匹配正则，大小写不敏感，匹配值大小写不敏感，只允许有一个值\nNONE：为空或不存在\nNOT_REGEX：正则不匹配，大小写不敏感，匹配值只允许有一个值\nSTART_WITH：开头是，大小写不敏感\nEND_WITH：结尾是，大小写不敏感\nWILDCARD：通配符匹配，大小写不敏感，*代表零个或多个任意字符，?代表任意单个字符\nNOT_WILDCARD：通配符不匹配，大小写不敏感，*代表零个或多个任意字符，?代表任意单个字符\n当matchName是REFERER时，matchType可选值为：\nEQUAL：等于，匹配值大小写敏感\nNOT_EQUAL：不等于，匹配值大小写敏感\nCONTAIN：包含，匹配值大小写不敏感\nNOT_CONTAIN：不包含，匹配值大小写不敏感\nREGEX：匹配正则，大小写不敏感，匹配值大小写不敏感，只允许有一个值\nNONE：为空或不存在\nNOT_REGEX：正则不匹配，大小写不敏感，匹配值只允许有一个值\nSTART_WITH：开头是，大小写不敏感\nEND_WITH：结尾是，大小写不敏感\nWILDCARD：通配符匹配，大小写不敏感，*代表零个或多个任意字符，?代表任意单个字符\nNOT_WILDCARD：通配符不匹配，大小写不敏感，*代表零个或多个任意字符，?代表任意单个字符\n当matchName是REQUEST_METHOD时，matchType可选值为：\nEQUAL：等于\nNOT_EQUAL：不等于"}
	MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
	// {"en":"Mathing key, this value is not empty and valid only when matchName=HEADER.\nMaximum 100 characters, case insensitive.","zh_CN":"条件key，只有当matchName=HEADER时，该值不为空且有效。\n最多100个字符，大小写不敏感。"}
	MatchKey *string `json:"matchKey,omitempty" xml:"matchKey,omitempty"`
	// {"en":"List of matching values..When matchName is REQUEST_METHOD, the optional values for matchValueList are GET/POST/DELETE/PUT/HEAD/OPTIONS/COPY.","zh_CN":"匹配值列表。当matchName是REQUEST_METHOD时，matchValueList可选值为GET/POST/DELETE/PUT/HEAD/OPTIONS/COPY。"}
	MatchValueList []*string `json:"matchValueList,omitempty" xml:"matchValueList,omitempty" type:"Repeated"`
}

func (s UpdateSpecificClientTrafficBypassRequestConditions) String() string {
	return tea.Prettify(s)
}

func (s UpdateSpecificClientTrafficBypassRequestConditions) GoString() string {
	return s.String()
}

func (s *UpdateSpecificClientTrafficBypassRequestConditions) SetMatchName(v string) *UpdateSpecificClientTrafficBypassRequestConditions {
	s.MatchName = &v
	return s
}

func (s *UpdateSpecificClientTrafficBypassRequestConditions) SetMatchType(v string) *UpdateSpecificClientTrafficBypassRequestConditions {
	s.MatchType = &v
	return s
}

func (s *UpdateSpecificClientTrafficBypassRequestConditions) SetMatchKey(v string) *UpdateSpecificClientTrafficBypassRequestConditions {
	s.MatchKey = &v
	return s
}

func (s *UpdateSpecificClientTrafficBypassRequestConditions) SetMatchValueList(v []*string) *UpdateSpecificClientTrafficBypassRequestConditions {
	s.MatchValueList = v
	return s
}

type UpdateSpecificClientTrafficBypassResponse struct {
	// {"dictionary":"belong=WAAP-MS-Ext|dict=waap_retCodeEnum","en":"Please refer to the error code for exceptions.","zh_CN":"请参照错误码。"}
	Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
	// {"en":"Description.","zh_CN":"描述信息。"}
	Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
}

func (s UpdateSpecificClientTrafficBypassResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateSpecificClientTrafficBypassResponse) GoString() string {
	return s.String()
}

func (s *UpdateSpecificClientTrafficBypassResponse) SetCode(v string) *UpdateSpecificClientTrafficBypassResponse {
	s.Code = &v
	return s
}

func (s *UpdateSpecificClientTrafficBypassResponse) SetMsg(v string) *UpdateSpecificClientTrafficBypassResponse {
	s.Msg = &v
	return s
}
