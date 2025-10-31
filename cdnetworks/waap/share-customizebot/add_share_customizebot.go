// This file is auto-generated, don't edit it. Thanks.
package share_customizebot

import (
	"github.com/alibabacloud-go/tea/tea"
)

type AddShareCustomizeBotTFRequest struct {
	// {"en":"Rule name, maximum 50 characters.Does not support # and &.","zh_CN":"规则名称，最多50个字符。\n不支持 # 和 &。"}
	BotName *string `json:"botName,omitempty" xml:"botName,omitempty" require:"true"`
	// {"en":"Description, maximum 200 characters.","zh_CN":"规则描述。最多200个字符。"}
	BotDescription *string `json:"botDescription,omitempty" xml:"botDescription,omitempty"`
	// {"en":"Actions.\nBLOCK: block\nLOG: log\nACCEPT: release","zh_CN":"处理动作。\nBLOCK：拦截\nLOG：监控\nACCEPT：放行"}
	BotAct *string `json:"botAct,omitempty" xml:"botAct,omitempty" require:"true"`
	// {"en":"Matching conditions.\nThere can only be at most one record per matching condition function under each type of condition.","zh_CN":"匹配条件。\n每个类型的条件下一种匹配条件函数最多只能有一条记录。"}
	ConditionList []*AddShareCustomizeBotTFRequestConditionList `json:"conditionList,omitempty" xml:"conditionList,omitempty" require:"true" type:"Repeated"`
	// {"en":"List of associated domain names.","zh_CN":"关联域名列表。"}
	RelaDomainList []*string `json:"relaDomainList,omitempty" xml:"relaDomainList,omitempty" type:"Repeated"`
}

func (s AddShareCustomizeBotTFRequest) String() string {
	return tea.Prettify(s)
}

func (s AddShareCustomizeBotTFRequest) GoString() string {
	return s.String()
}

func (s *AddShareCustomizeBotTFRequest) SetBotName(v string) *AddShareCustomizeBotTFRequest {
	s.BotName = &v
	return s
}

func (s *AddShareCustomizeBotTFRequest) SetBotDescription(v string) *AddShareCustomizeBotTFRequest {
	s.BotDescription = &v
	return s
}

func (s *AddShareCustomizeBotTFRequest) SetBotAct(v string) *AddShareCustomizeBotTFRequest {
	s.BotAct = &v
	return s
}

func (s *AddShareCustomizeBotTFRequest) SetConditionList(v []*AddShareCustomizeBotTFRequestConditionList) *AddShareCustomizeBotTFRequest {
	s.ConditionList = v
	return s
}

func (s *AddShareCustomizeBotTFRequest) SetRelaDomainList(v []*string) *AddShareCustomizeBotTFRequest {
	s.RelaDomainList = v
	return s
}

type AddShareCustomizeBotTFRequestConditionList struct {
	// {"en":"Matching condition name. \nIP_IPS: IP/IP segment \nJA3: JA3 Fingerprint\nJA4: JA4 Fingerprint\nUA: User-agent \nHEADER: Request Header \nASN: AS Number \nCLIENT_GROUP: Client Group \nPUBLIC_BOT: Public Bots","zh_CN":"匹配条件名称。 \nIP_IPS：IP/IP段 \nJA3：JA3指纹\nJA4：JA4指纹\nUA：User-agent \nHEADER：请求头 \nASN：AS号 \nCLIENT_GROUP：客户端分组 \nPUBLIC_BOT：公开Bots"}
	ConditionName *string `json:"conditionName,omitempty" xml:"conditionName,omitempty" require:"true"`
	// {"en":"Condition value list.\nWhen conditionName is IP_IPS, maximum 300 IP/CIDR.\nWhen conditionName is JA3, maximum 300 JA3 Fingerprint.\nWhen conditionName is JA4, maximum 300 JA4 Fingerprint.\nWhen conditionName is PUBLIC_BOT, Supported values: search_engine_bot/site_monitor/marketing_analysis/feed_fetcher/tool/page_preview.","zh_CN":"条件值列表。\n当conditionName是IP_IPS时，最多300个IP/CIDR。\n当conditionName是JA3时，最多300个JA3指纹。\n当conditionName是JA4时，最多300个JA4指纹。\n当conditionName为PUBLIC_BOT时，支持的值：search_engine_bot/site_monitor/marketing_analysis/feed_fetcher/tool/page_preview。"}
	ConditionValueList []*string `json:"conditionValueList,omitempty" xml:"conditionValueList,omitempty" require:"true" type:"Repeated"`
	// {"en":"Matching condition function.\nEQUAL: Equals\nNOT_EQUAL: Does not equal\nCONTAIN: Contains\nNOT_CONTAIN: Does not contain\nNONE: Empty or non-existent\nREGEX: Regex match\nNOT_REGEX: Regex does not match\nSTART_WITH: Starts with\nEND_WITH: Ends with\nWILDCARD: Wildcard matches, * represents zero or more arbitrary characters, ? represents any single character\nNOT_WILDCARD: Wildcard does not match, * represents zero or more arbitrary characters, ? represents any single character\n\nWhen conditionName is IP_IPS, the value can be: EQUAL,NOT_EQUAL \nWhen conditionName is JA3, the value can be: EQUAL,NOT_EQUAL\nWhen conditionName is JA4, the value can be: EQUAL,NOT_EQUAL,CONTAIN,NOT_CONTAIN,START_WITH,END_WITH,WILDCARD,NOT_WILDCARD\nWhen conditionName is UA, the value can be: EQUAL,NOT_EQUAL,CONTAIN,NOT_CONTAIN,NONE,REGEX,NOT_REGEX,START_WITH,END_WITH,WILDCARD,NOT_WILDCARD \nWhen conditionName is HEADER, the value can be: EQUAL,NOT_EQUAL,CONTAIN,NOT_CONTAIN,NONE,REGEX,NOT_REGEX,START_WITH,END_WITH,WILDCARD,NOT_WILDCARD \nWhen conditionName is ASN, the value can be: EQUAL,NOT_EQUAL \nWhen conditionName is CLIENT_GROUP, the value can be: EQUAL,NOT_EQUAL \nWhen conditionName is PUBLIC_BOT, the value can be: EQUAL","zh_CN":"匹配条件函数。\nEQUAL：等于\nNOT_EQUAL：不等于\nCONTAIN：包含\nNOT_CONTAIN：不包含\nNONE：为空或不存在\nREGEX：匹配正则\nNOT_REGEX：正则不匹配\nSTART_WITH：开头是\nEND_WITH：结尾是\nWILDCARD：通配符匹配，*代表零个或多个任意字符，?代表任意单个字符\nNOT_WILDCARD：通配符不匹配，*代表零个或多个任意字符，?代表任意单个字符\n\n当conditionName为IP_IPS时，取值范围为：EQUAL、NOT_EQUAL \n当conditionName为JA3时，取值范围为：EQUAL、NOT_EQUAL \n当conditionName为JA4时，取值范围为：EQUAL、NOT_EQUAL、CONTAIN、NOT_CONTAIN、START_WITH、END_WITH、WILDCARD、NOT_WILDCARD\n当conditionName为UA时，取值范围为：EQUAL、NOT_EQUAL、CONTAIN、NOT_CONTAIN、NONE、REGEX、NOT_REGEX、START_WITH、END_WITH、WILDCARD、NOT_WILDCARD\n当conditionName为HEADER时，取值范围为：EQUAL、NOT_EQUAL、CONTAIN、NOT_CONTAIN、NONE、REGEX、NOT_REGEX、START_WITH、END_WITH、WILDCARD、NOT_WILDCARD\n当conditionName为ASN时，取值范围为：EQUAL、NOT_EQUAL \n当conditionName为CLIENT_GROUP时，取值范围为：EQUAL、NOT_EQUAL \n当conditionName为PUBLIC_BOT时，取值范围为：EQUAL"}
	ConditionFunc *string `json:"conditionFunc,omitempty" xml:"conditionFunc,omitempty" require:"true"`
	// {"en":"Request header name. Case insensitive.","zh_CN":"头部名称，大小写不敏感"}
	ConditionKey *string `json:"conditionKey,omitempty" xml:"conditionKey,omitempty"`
}

func (s AddShareCustomizeBotTFRequestConditionList) String() string {
	return tea.Prettify(s)
}

func (s AddShareCustomizeBotTFRequestConditionList) GoString() string {
	return s.String()
}

func (s *AddShareCustomizeBotTFRequestConditionList) SetConditionName(v string) *AddShareCustomizeBotTFRequestConditionList {
	s.ConditionName = &v
	return s
}

func (s *AddShareCustomizeBotTFRequestConditionList) SetConditionValueList(v []*string) *AddShareCustomizeBotTFRequestConditionList {
	s.ConditionValueList = v
	return s
}

func (s *AddShareCustomizeBotTFRequestConditionList) SetConditionFunc(v string) *AddShareCustomizeBotTFRequestConditionList {
	s.ConditionFunc = &v
	return s
}

func (s *AddShareCustomizeBotTFRequestConditionList) SetConditionKey(v string) *AddShareCustomizeBotTFRequestConditionList {
	s.ConditionKey = &v
	return s
}

type AddShareCustomizeBotTFResponse struct {
	// {"en":"Description.","zh_CN":"描述信息。"}
	Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
	// {"dictionary":"belong=WAAP-MS-Ext|dict=waap_retCodeEnum","en":"Please refer to the error code for exceptions.","zh_CN":"请参照错误码。"}
	Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
	// {"en":"","zh_CN":""}
	Data *AddShareCustomizeBotTFResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
}

func (s AddShareCustomizeBotTFResponse) String() string {
	return tea.Prettify(s)
}

func (s AddShareCustomizeBotTFResponse) GoString() string {
	return s.String()
}

func (s *AddShareCustomizeBotTFResponse) SetMsg(v string) *AddShareCustomizeBotTFResponse {
	s.Msg = &v
	return s
}

func (s *AddShareCustomizeBotTFResponse) SetCode(v string) *AddShareCustomizeBotTFResponse {
	s.Code = &v
	return s
}

func (s *AddShareCustomizeBotTFResponse) SetData(v *AddShareCustomizeBotTFResponseData) *AddShareCustomizeBotTFResponse {
	s.Data = v
	return s
}

type AddShareCustomizeBotTFResponseData struct {
	// {"en":"Rule ID.","zh_CN":"规则ID。"}
	Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
	// {"en":"Rule name.","zh_CN":"规则名称。"}
	BotName *string `json:"botName,omitempty" xml:"botName,omitempty" require:"true"`
	// {"en":"Description.","zh_CN":"规则描述。"}
	BotDescription *string `json:"botDescription,omitempty" xml:"botDescription,omitempty" require:"true"`
	// {"en":"Actions.\nBLOCK: block\nLOG: log\nACCEPT: release","zh_CN":"处理动作。\nBLOCK：拦截\nLOG：监控\nACCEPT：放行"}
	BotAct *string `json:"botAct,omitempty" xml:"botAct,omitempty" require:"true"`
	// {"en":"Matching conditions.","zh_CN":"匹配条件。"}
	ConditionList []*AddShareCustomizeBotTFResponseDataConditionList `json:"conditionList,omitempty" xml:"conditionList,omitempty" require:"true" type:"Repeated"`
	// {"en":"List of associated hostnames.","zh_CN":"关联的域名列表。"}
	RelaDomainList []*string `json:"relaDomainList,omitempty" xml:"relaDomainList,omitempty" require:"true" type:"Repeated"`
}

func (s AddShareCustomizeBotTFResponseData) String() string {
	return tea.Prettify(s)
}

func (s AddShareCustomizeBotTFResponseData) GoString() string {
	return s.String()
}

func (s *AddShareCustomizeBotTFResponseData) SetId(v string) *AddShareCustomizeBotTFResponseData {
	s.Id = &v
	return s
}

func (s *AddShareCustomizeBotTFResponseData) SetBotName(v string) *AddShareCustomizeBotTFResponseData {
	s.BotName = &v
	return s
}

func (s *AddShareCustomizeBotTFResponseData) SetBotDescription(v string) *AddShareCustomizeBotTFResponseData {
	s.BotDescription = &v
	return s
}

func (s *AddShareCustomizeBotTFResponseData) SetBotAct(v string) *AddShareCustomizeBotTFResponseData {
	s.BotAct = &v
	return s
}

func (s *AddShareCustomizeBotTFResponseData) SetConditionList(v []*AddShareCustomizeBotTFResponseDataConditionList) *AddShareCustomizeBotTFResponseData {
	s.ConditionList = v
	return s
}

func (s *AddShareCustomizeBotTFResponseData) SetRelaDomainList(v []*string) *AddShareCustomizeBotTFResponseData {
	s.RelaDomainList = v
	return s
}

type AddShareCustomizeBotTFResponseDataConditionList struct {
	// {"en":"Matching condition name. \nIP_IPS: IP/IP segment \nJA3: JA3 Fingerprint\nJA4: JA4 Fingerprint\nUA: User-agent \nHEADER: Request Header \nASN: AS Number \nCLIENT_GROUP: Client Group \nPUBLIC_BOT: Public Bots","zh_CN":"匹配条件名称。 \nIP_IPS：IP/IP段 \nJA3：JA3指纹\nJA4：JA4指纹\nUA：User-agent \nHEADER：请求头 \nASN：AS号 \nCLIENT_GROUP：客户端分组 \nPUBLIC_BOT：公开Bots"}
	ConditionName *string `json:"conditionName,omitempty" xml:"conditionName,omitempty" require:"true"`
	// {"en":"Condition value list.","zh_CN":"条件值列表。"}
	ConditionValueList []*string `json:"conditionValueList,omitempty" xml:"conditionValueList,omitempty" require:"true" type:"Repeated"`
	// {"en":"Matching condition function.\nEQUAL: Equals\nNOT_EQUAL: Does not equal\nCONTAIN: Contains\nNOT_CONTAIN: Does not contain\nNONE: Empty or non-existent\nREGEX: Regex match\nNOT_REGEX: Does not match regex\nSTART_WITH: Starts with\nEND_WITH: Ends with\nWILDCARD: Wildcard matches, * represents zero or more arbitrary characters, ? represents any single character\nNOT_WILDCARD: Wildcard does not match, * represents zero or more arbitrary characters, ? represents any single character","zh_CN":"匹配条件函数。\nEQUAL：等于\nNOT_EQUAL：不等于\nCONTAIN：包含\nNOT_CONTAIN：不包含\nNONE：为空或不存在\nREGEX：匹配正则\nNOT_REGEX：正则不匹配\nSTART_WITH：开头是\nEND_WITH：结尾是\nWILDCARD：通配符匹配，*代表零个或多个任意字符，?代表任意单个字符\nNOT_WILDCARD：通配符不匹配，*代表零个或多个任意字符，?代表任意单个字符"}
	ConditionFunc *string `json:"conditionFunc,omitempty" xml:"conditionFunc,omitempty" require:"true"`
	// {"en":"Request header name.","zh_CN":"头部名称。"}
	ConditionKey *string `json:"conditionKey,omitempty" xml:"conditionKey,omitempty" require:"true"`
}

func (s AddShareCustomizeBotTFResponseDataConditionList) String() string {
	return tea.Prettify(s)
}

func (s AddShareCustomizeBotTFResponseDataConditionList) GoString() string {
	return s.String()
}

func (s *AddShareCustomizeBotTFResponseDataConditionList) SetConditionName(v string) *AddShareCustomizeBotTFResponseDataConditionList {
	s.ConditionName = &v
	return s
}

func (s *AddShareCustomizeBotTFResponseDataConditionList) SetConditionValueList(v []*string) *AddShareCustomizeBotTFResponseDataConditionList {
	s.ConditionValueList = v
	return s
}

func (s *AddShareCustomizeBotTFResponseDataConditionList) SetConditionFunc(v string) *AddShareCustomizeBotTFResponseDataConditionList {
	s.ConditionFunc = &v
	return s
}

func (s *AddShareCustomizeBotTFResponseDataConditionList) SetConditionKey(v string) *AddShareCustomizeBotTFResponseDataConditionList {
	s.ConditionKey = &v
	return s
}
