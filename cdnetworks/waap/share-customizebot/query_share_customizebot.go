// This file is auto-generated, don't edit it. Thanks.
package share_customizebot

import (
	"github.com/alibabacloud-go/tea/tea"
)

type ListShareCustomizeBotsRequest struct {
	// {"en":"Bot name,fuzzy query.","zh_CN":"Bot名称，模糊查询。"}
	BotName *string `json:"botName,omitempty" xml:"botName,omitempty"`
}

func (s ListShareCustomizeBotsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListShareCustomizeBotsRequest) GoString() string {
	return s.String()
}

func (s *ListShareCustomizeBotsRequest) SetBotName(v string) *ListShareCustomizeBotsRequest {
	s.BotName = &v
	return s
}

type ListShareCustomizeBotsResponse struct {
	// {"dictionary":"belong=WAAP-MS-Ext|dict=waap_retCodeEnum","en":"Please refer to the error code for exceptions.","zh_CN":"请参照错误码。"}
	Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
	// {"en":"Description.","zh_CN":"描述信息。"}
	Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
	// {"en":"Data.","zh_CN":"出参数据。"}
	Data []*ListShareCustomizeBotsResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s ListShareCustomizeBotsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListShareCustomizeBotsResponse) GoString() string {
	return s.String()
}

func (s *ListShareCustomizeBotsResponse) SetCode(v string) *ListShareCustomizeBotsResponse {
	s.Code = &v
	return s
}

func (s *ListShareCustomizeBotsResponse) SetMsg(v string) *ListShareCustomizeBotsResponse {
	s.Msg = &v
	return s
}

func (s *ListShareCustomizeBotsResponse) SetData(v []*ListShareCustomizeBotsResponseData) *ListShareCustomizeBotsResponse {
	s.Data = v
	return s
}

type ListShareCustomizeBotsResponseData struct {
	// {"en":"Rule ID.","zh_CN":"规则ID。"}
	Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
	// {"en":"Rule name.","zh_CN":"规则名称。"}
	BotName *string `json:"botName,omitempty" xml:"botName,omitempty" require:"true"`
	// {"en":"Description.","zh_CN":"规则描述。"}
	BotDescription *string `json:"botDescription,omitempty" xml:"botDescription,omitempty" require:"true"`
	// {"en":"Actions.\nBLOCK: block\nLOG: log\nACCEPT: release","zh_CN":"处理动作。\nBLOCK：拦截\nLOG：监控\nACCEPT：放行"}
	BotAct *string `json:"botAct,omitempty" xml:"botAct,omitempty" require:"true"`
	// {"en":"Matching conditions.","zh_CN":"匹配条件。"}
	ConditionList []*ListShareCustomizeBotsResponseDataConditionList `json:"conditionList,omitempty" xml:"conditionList,omitempty" require:"true" type:"Repeated"`
	// {"en":"List of associated hostnames.","zh_CN":"关联的域名列表。"}
	RelaDomainList []*string `json:"relaDomainList,omitempty" xml:"relaDomainList,omitempty" require:"true" type:"Repeated"`
	// {"en":"Creator.","zh_CN":"创建者。"}
	Creator *string `json:"creator,omitempty" xml:"creator,omitempty" require:"true"`
}

func (s ListShareCustomizeBotsResponseData) String() string {
	return tea.Prettify(s)
}

func (s ListShareCustomizeBotsResponseData) GoString() string {
	return s.String()
}

func (s *ListShareCustomizeBotsResponseData) SetId(v string) *ListShareCustomizeBotsResponseData {
	s.Id = &v
	return s
}

func (s *ListShareCustomizeBotsResponseData) SetBotName(v string) *ListShareCustomizeBotsResponseData {
	s.BotName = &v
	return s
}

func (s *ListShareCustomizeBotsResponseData) SetBotDescription(v string) *ListShareCustomizeBotsResponseData {
	s.BotDescription = &v
	return s
}

func (s *ListShareCustomizeBotsResponseData) SetBotAct(v string) *ListShareCustomizeBotsResponseData {
	s.BotAct = &v
	return s
}

func (s *ListShareCustomizeBotsResponseData) SetConditionList(v []*ListShareCustomizeBotsResponseDataConditionList) *ListShareCustomizeBotsResponseData {
	s.ConditionList = v
	return s
}

func (s *ListShareCustomizeBotsResponseData) SetRelaDomainList(v []*string) *ListShareCustomizeBotsResponseData {
	s.RelaDomainList = v
	return s
}

func (s *ListShareCustomizeBotsResponseData) SetCreator(v string) *ListShareCustomizeBotsResponseData {
	s.Creator = &v
	return s
}

type ListShareCustomizeBotsResponseDataConditionList struct {
	// {"en":"Matching condition name. \nIP_IPS: IP/IP segment \nJA3: JA3 Fingerprint\nJA4: JA4 Fingerprint\nUA: User-agent \nHEADER: Request Header \nASN: AS Number \nCLIENT_GROUP: Client Group \nPUBLIC_BOT: Public Bots","zh_CN":"匹配条件名称。 \nIP_IPS：IP/IP段 \nJA3：JA3指纹\nJA4：JA4指纹\nUA：User-agent \nHEADER：请求头 \nASN：AS号 \nCLIENT_GROUP：客户端分组 \nPUBLIC_BOT：公开Bots"}
	ConditionName *string `json:"conditionName,omitempty" xml:"conditionName,omitempty" require:"true"`
	// {"en":"Condition value list.","zh_CN":"条件值列表。"}
	ConditionValueList []*string `json:"conditionValueList,omitempty" xml:"conditionValueList,omitempty" require:"true" type:"Repeated"`
	// {"en":"Matching condition function.\nEQUAL: Equals\nNOT_EQUAL: Does not equal\nCONTAIN: Contains\nNOT_CONTAIN: Does not contain\nNONE: Empty or non-existent\nREGEX: Regex match\nNOT_REGEX: Does not match regex\nSTART_WITH: Starts with\nEND_WITH: Ends with\nWILDCARD: Wildcard matches, * represents zero or more arbitrary characters, ? represents any single character\nNOT_WILDCARD: Wildcard does not match, * represents zero or more arbitrary characters, ? represents any single character","zh_CN":"匹配条件函数。\nEQUAL：等于\nNOT_EQUAL：不等于\nCONTAIN：包含\nNOT_CONTAIN：不包含\nNONE：为空或不存在\nREGEX：匹配正则\nNOT_REGEX：正则不匹配\nSTART_WITH：开头是\nEND_WITH：结尾是\nWILDCARD：通配符匹配，*代表零个或多个任意字符，?代表任意单个字符\nNOT_WILDCARD：通配符不匹配，*代表零个或多个任意字符，?代表任意单个字符"}
	ConditionFunc *string `json:"conditionFunc,omitempty" xml:"conditionFunc,omitempty" require:"true"`
	// {"en":"Request header name.","zh_CN":"头部名称。"}
	ConditionKey *string `json:"conditionKey,omitempty" xml:"conditionKey,omitempty" require:"true"`
}

func (s ListShareCustomizeBotsResponseDataConditionList) String() string {
	return tea.Prettify(s)
}

func (s ListShareCustomizeBotsResponseDataConditionList) GoString() string {
	return s.String()
}

func (s *ListShareCustomizeBotsResponseDataConditionList) SetConditionName(v string) *ListShareCustomizeBotsResponseDataConditionList {
	s.ConditionName = &v
	return s
}

func (s *ListShareCustomizeBotsResponseDataConditionList) SetConditionValueList(v []*string) *ListShareCustomizeBotsResponseDataConditionList {
	s.ConditionValueList = v
	return s
}

func (s *ListShareCustomizeBotsResponseDataConditionList) SetConditionFunc(v string) *ListShareCustomizeBotsResponseDataConditionList {
	s.ConditionFunc = &v
	return s
}

func (s *ListShareCustomizeBotsResponseDataConditionList) SetConditionKey(v string) *ListShareCustomizeBotsResponseDataConditionList {
	s.ConditionKey = &v
	return s
}
