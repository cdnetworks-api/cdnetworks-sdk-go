// This file is auto-generated, don't edit it. Thanks.
package bot_scene_whitelist

import (
	"github.com/alibabacloud-go/tea/tea"
)

type ListSpecificClientTrafficBypassRequest struct {
	// {"en":"Hostname list.","zh_CN":"域名列表。"}
	DomainList []*string `json:"domainList,omitempty" xml:"domainList,omitempty" require:"true" type:"Repeated"`
}

func (s ListSpecificClientTrafficBypassRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSpecificClientTrafficBypassRequest) GoString() string {
	return s.String()
}

func (s *ListSpecificClientTrafficBypassRequest) SetDomainList(v []*string) *ListSpecificClientTrafficBypassRequest {
	s.DomainList = v
	return s
}

type ListSpecificClientTrafficBypassResponse struct {
	// {"dictionary":"belong=WAAP-MS-Ext|dict=waap_retCodeEnum","en":"Please refer to the error code for exceptions.","zh_CN":"请参照错误码。"}
	Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
	// {"en":"Description.","zh_CN":"描述信息。"}
	Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
	// {"en":"Data.","zh_CN":"出参数据。"}
	Data []*ListSpecificClientTrafficBypassResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s ListSpecificClientTrafficBypassResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSpecificClientTrafficBypassResponse) GoString() string {
	return s.String()
}

func (s *ListSpecificClientTrafficBypassResponse) SetCode(v string) *ListSpecificClientTrafficBypassResponse {
	s.Code = &v
	return s
}

func (s *ListSpecificClientTrafficBypassResponse) SetMsg(v string) *ListSpecificClientTrafficBypassResponse {
	s.Msg = &v
	return s
}

func (s *ListSpecificClientTrafficBypassResponse) SetData(v []*ListSpecificClientTrafficBypassResponseData) *ListSpecificClientTrafficBypassResponse {
	s.Data = v
	return s
}

type ListSpecificClientTrafficBypassResponseData struct {
	// {"en":"ID.","zh_CN":"ID。"}
	Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
	// {"en":"Hostname.","zh_CN":"域名。"}
	Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
	// {"en":"Whitelist name.","zh_CN":"白名单名称。"}
	Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
	// {"en":"Description.","zh_CN":"描述。"}
	Description *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	// {"en":"List of matching conditions.","zh_CN":"匹配条件列表。"}
	Conditions []*ListSpecificClientTrafficBypassResponseDataConditions `json:"conditions,omitempty" xml:"conditions,omitempty" require:"true" type:"Repeated"`
}

func (s ListSpecificClientTrafficBypassResponseData) String() string {
	return tea.Prettify(s)
}

func (s ListSpecificClientTrafficBypassResponseData) GoString() string {
	return s.String()
}

func (s *ListSpecificClientTrafficBypassResponseData) SetId(v string) *ListSpecificClientTrafficBypassResponseData {
	s.Id = &v
	return s
}

func (s *ListSpecificClientTrafficBypassResponseData) SetDomain(v string) *ListSpecificClientTrafficBypassResponseData {
	s.Domain = &v
	return s
}

func (s *ListSpecificClientTrafficBypassResponseData) SetName(v string) *ListSpecificClientTrafficBypassResponseData {
	s.Name = &v
	return s
}

func (s *ListSpecificClientTrafficBypassResponseData) SetDescription(v string) *ListSpecificClientTrafficBypassResponseData {
	s.Description = &v
	return s
}

func (s *ListSpecificClientTrafficBypassResponseData) SetConditions(v []*ListSpecificClientTrafficBypassResponseDataConditions) *ListSpecificClientTrafficBypassResponseData {
	s.Conditions = v
	return s
}

type ListSpecificClientTrafficBypassResponseDataConditions struct {
	// {"en":"Matching condition name.\nIP_IPS: IP/CIDR\nPATH: Path\nURI: Path with parameters\nHEADER: Request Header\nUA: User Agent\nREQUEST_METHOD: Request Method\nREFERER: Referer","zh_CN":"匹配条件名称。\nIP_IPS：IP/IP段\nPATH：路径\nURI：带参数路径\nHEADER：请求头\nUA: User-Agent\nREQUEST_METHOD: 请求方法\nREFERER: Referer"}
	MatchName *string `json:"matchName,omitempty" xml:"matchName,omitempty" require:"true"`
	// {"en":"When matchName is IP_IPS, maximum 300 IP/CIDR in match value list, the optional value of matchType is:\nEQUAL: Equals\nNOT_EQUAL: Does not equal\nWhen matchName is a URI, the optional value of matchType is:\nEQUAL: Equals, the matching value is case-sensitive and needs to start with \"/\" and include parameters\nNOT_EQUAL: Does not equal, the matching value is case-sensitive, needs to start with \"/\", and contains parameters\nCONTAIN: Contains, match values are case insensitive\nNOT_CONTAIN: Does not contains, match values are case insensitive\nREGEX: Regex match, only one value is allowed for the match value\nNOT_REGEX: regular does not match\nSTART_WITH: starts with\nEND_WITH: ends with\nWILDCARD: wildcard matches\nNOT_WILDCARD: wildcard does not match\nWhen matchName is PATH, the optional value of matchType is:\nEQUAL: Equals, the matching value is case-sensitive and needs to start with \"/\" , and does not contain parameters\nNOT_EQUAL: Does not equal, the matching value is case-sensitive, needs to start with \"/\", and does not contain parameters\nCONTAIN: Contains, match values are case insensitive\nNOT_CONTAIN: Does not contains, match values are case insensitive\nREGEX: Regex match, match values are case insensitive and only one value is allowed\nNOT_REGEX: regular does not match\nSTART_WITH: starts with\nEND_WITH: ends with\nWILDCARD: wildcard matches\nNOT_WILDCARD: wildcard does not match\nWhen matchName is HEADER, the optional value of matchType is:\nEQUAL: Equals, match values are case sensitive\nNOT_EQUAL: Does not equal, the matching value is case-sensitive\nCONTAIN: Contains, match values are case insensitive\nNOT_CONTAIN: Does not contains, match values are case insensitive\nREGEX: Regex match, match values are case insensitive and only one value is allowed\nNONE: Empty or does not exist\nNOT_REGEX: regular does not match\nSTART_WITH: starts with\nEND_WITH: ends with\nWILDCARD: wildcard matches\nNOT_WILDCARD: wildcard does not match\nWhen matchName is UA, the optional value of matchType is:\nEQUAL: Equals, match values are case sensitive\nNOT_EQUAL: Does not equal, the matching value is case-sensitive\nCONTAIN: Contains, match values are case insensitive\nNOT_CONTAIN: Does not contains, match values are case insensitive\nREGEX: Regex match, match values are case insensitive and only one value is allowed\nNONE: Empty or does not exist\nNOT_REGEX: regular does not match\nSTART_WITH: starts with\nEND_WITH: ends with\nWILDCARD: wildcard matches\nNOT_WILDCARD: wildcard does not match\nWhen matchName is REFERER, the optional value of matchType is:\nEQUAL: Equals, match values are case sensitive\nNOT_EQUAL: Does not equal, the matching value is case-sensitive\nCONTAIN: Contains, match values are case insensitive\nNOT_CONTAIN: Does not contains, match values are case insensitive\nREGEX: Regex match, match values are case insensitive and only one value is allowed\nNONE: Empty or does not exist\nNOT_REGEX: regular does not match\nSTART_WITH: starts with\nEND_WITH: ends with\nWILDCARD: wildcard matches\nNOT_WILDCARD: wildcard does not match\nWhen matchName is REQUEST_METHOD, the optional value of matchType is:\nEQUAL: Equals, match values are case sensitive\nNOT_EQUAL: Does not equal, the matching value is case-sensitive","zh_CN":"当matchName是IP_IPS时，匹配值列表中最多300个IP或IP段，matchType可选值为：\nEQUAL：等于\nNOT_EQUAL：不等于\n当matchName是URI时，matchType可选值为：\nEQUAL：等于匹配值大小写敏感，需要以\"/\"开头，含参数\nNOT_EQUAL：不等于，匹配值大小写敏感，需要以\"/\"开头，含参数\nCONTAIN：包含，匹配值大小写不敏感\nNOT_CONTAIN：不包含，匹配值大小写不敏感\nREGEX：匹配正则，匹配值只允许有一个值\nNOT_REGEX：正则不匹配\nSTART_WITH：开头是\nEND_WITH：结尾是\nWILDCARD：通配符匹配\nNOT_WILDCARD：通配符不匹配\n当matchName是PATH时，matchType可选值为：\nEQUAL：等于，匹配值大小写敏感，需要以\"/\"开头，不含参数\nNOT_EQUAL：不等于，匹配值大小写敏感，需要以\"/\"开头，不含参数\nCONTAIN：包含，匹配值大小写不敏感\nNOT_CONTAIN：不包含，匹配值大小写不敏感\nREGEX：匹配正则，匹配值大小写不敏感，只允许有一个值\nNOT_REGEX：正则不匹配\nSTART_WITH：开头是\nEND_WITH：结尾是\nWILDCARD：通配符匹配\nNOT_WILDCARD：通配符不匹配\n当matchName是HEADER时，matchType可选值为：\nEQUAL：等于，匹配值大小写敏感\nNOT_EQUAL：不等于，匹配值大小写敏感\nCONTAIN：包含，匹配值大小写不敏感\nNOT_CONTAIN：不包含，匹配值大小写不敏感\nREGEX：匹配正则，匹配值大小写不敏感，只允许有一个值\nNONE：为空或不存在\nNOT_REGEX：正则不匹配\nSTART_WITH：开头是\nEND_WITH：结尾是\nWILDCARD：通配符匹配\nNOT_WILDCARD：通配符不匹配\n当matchName是UA时，matchType可选值为：\nEQUAL：等于，匹配值大小写敏感\nNOT_EQUAL：不等于，匹配值大小写敏感\nCONTAIN：包含，匹配值大小写不敏感\nNOT_CONTAIN：不包含，匹配值大小写不敏感\nREGEX：匹配正则，匹配值大小写不敏感，只允许有一个值\nNONE：为空或不存在\nNOT_REGEX：正则不匹配\nSTART_WITH：开头是\nEND_WITH：结尾是\nWILDCARD：通配符匹配\nNOT_WILDCARD：通配符不匹配\n当matchName是REFERER时，matchType可选值为：\nEQUAL：等于，匹配值大小写敏感\nNOT_EQUAL：不等于，匹配值大小写敏感\nCONTAIN：包含，匹配值大小写不敏感\nNOT_CONTAIN：不包含，匹配值大小写不敏感\nREGEX：匹配正则，匹配值大小写不敏感，只允许有一个值\nNONE：为空或不存在\nNOT_REGEX：正则不匹配\nSTART_WITH：开头是\nEND_WITH：结尾是\nWILDCARD：通配符匹配\nNOT_WILDCARD：通配符不匹配\n当matchName是REQUEST_METHOD时，matchType可选值为：\nEQUAL：等于，匹配值大小写敏感\nNOT_EQUAL：不等于，匹配值大小写敏感"}
	MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
	// {"en":"Mathing key, this value is not empty and valid only when matchName=HEADER.\nMaximum 100 characters, case insensitive.","zh_CN":"条件key，只有当matchName=HEADER时，该值不为空且有效。\n最多100个字符，大小写不敏感。"}
	MatchKey *string `json:"matchKey,omitempty" xml:"matchKey,omitempty" require:"true"`
	// {"en":"List of matching values.","zh_CN":"匹配值列表。"}
	MatchValueList []*string `json:"matchValueList,omitempty" xml:"matchValueList,omitempty" require:"true" type:"Repeated"`
}

func (s ListSpecificClientTrafficBypassResponseDataConditions) String() string {
	return tea.Prettify(s)
}

func (s ListSpecificClientTrafficBypassResponseDataConditions) GoString() string {
	return s.String()
}

func (s *ListSpecificClientTrafficBypassResponseDataConditions) SetMatchName(v string) *ListSpecificClientTrafficBypassResponseDataConditions {
	s.MatchName = &v
	return s
}

func (s *ListSpecificClientTrafficBypassResponseDataConditions) SetMatchType(v string) *ListSpecificClientTrafficBypassResponseDataConditions {
	s.MatchType = &v
	return s
}

func (s *ListSpecificClientTrafficBypassResponseDataConditions) SetMatchKey(v string) *ListSpecificClientTrafficBypassResponseDataConditions {
	s.MatchKey = &v
	return s
}

func (s *ListSpecificClientTrafficBypassResponseDataConditions) SetMatchValueList(v []*string) *ListSpecificClientTrafficBypassResponseDataConditions {
	s.MatchValueList = v
	return s
}
