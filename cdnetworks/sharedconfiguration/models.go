package sharedconfiguration

// This file is auto-generated, don't edit it. Thanks.
import (
  "github.com/alibabacloud-go/tea/tea"
)

type ListSharedWhitelistRulesRequest struct {
  // {"en":"Rule name, fuzzy query.","zh_CN":"规则名称，模糊查询。"}
  RuleName *string `json:"ruleName,omitempty" xml:"ruleName,omitempty"`
}

func (s ListSharedWhitelistRulesRequest) String() string {
  return tea.Prettify(s)
}

func (s ListSharedWhitelistRulesRequest) GoString() string {
  return s.String()
}

func (s *ListSharedWhitelistRulesRequest) SetRuleName(v string) *ListSharedWhitelistRulesRequest {
  s.RuleName = &v
  return s
}

type ListSharedWhitelistRulesRequestHeader struct {
  // {"dictionary":"belong=WAAP-MS-Ext|dict=waap_serviceType","en":"Security service type. Please enter a specific service type, if you purchase multiple security services.","zh_CN":"安全服务类型。有使用多个不同的安全服务时，需要填写具体的服务类型。"}
  ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty"`
}

func (s ListSharedWhitelistRulesRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ListSharedWhitelistRulesRequestHeader) GoString() string {
  return s.String()
}

func (s *ListSharedWhitelistRulesRequestHeader) SetServiceType(v string) *ListSharedWhitelistRulesRequestHeader {
  s.ServiceType = &v
  return s
}

type ListSharedWhitelistRulesPaths struct {
}

func (s ListSharedWhitelistRulesPaths) String() string {
  return tea.Prettify(s)
}

func (s ListSharedWhitelistRulesPaths) GoString() string {
  return s.String()
}

type ListSharedWhitelistRulesParameters struct {
}

func (s ListSharedWhitelistRulesParameters) String() string {
  return tea.Prettify(s)
}

func (s ListSharedWhitelistRulesParameters) GoString() string {
  return s.String()
}

type ListSharedWhitelistRulesResponse struct {
  // {"dictionary":"belong=WAAP-MS-Ext|dict=waap_retCodeEnum","en":"Please refer to the error code for exceptions.","zh_CN":"请参照错误码。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Description.","zh_CN":"描述信息。"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"Data.","zh_CN":"出参数据。"}
  Data []*ListSharedWhitelistRulesResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s ListSharedWhitelistRulesResponse) String() string {
  return tea.Prettify(s)
}

func (s ListSharedWhitelistRulesResponse) GoString() string {
  return s.String()
}

func (s *ListSharedWhitelistRulesResponse) SetCode(v string) *ListSharedWhitelistRulesResponse {
  s.Code = &v
  return s
}

func (s *ListSharedWhitelistRulesResponse) SetMsg(v string) *ListSharedWhitelistRulesResponse {
  s.Msg = &v
  return s
}

func (s *ListSharedWhitelistRulesResponse) SetData(v []*ListSharedWhitelistRulesResponseData) *ListSharedWhitelistRulesResponse {
  s.Data = v
  return s
}

type ListSharedWhitelistRulesResponseData struct     {
  // {"en":"Rule ID.","zh_CN":"规则ID。"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {"en":"Rule name.","zh_CN":"规则名称。"}
  RuleName *string `json:"ruleName,omitempty" xml:"ruleName,omitempty" require:"true"`
  // {"en":"Description.","zh_CN":"规则描述。"}
  Description *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
  // {"en":"creator.","zh_CN":"创建者。"}
  Creator *string `json:"creator,omitempty" xml:"creator,omitempty" require:"true"`
  // {"en":"Number of hostnames associated.","zh_CN":"关联的域名数。"}
  ImportedDomainCount *int `json:"importedDomainCount,omitempty" xml:"importedDomainCount,omitempty" require:"true"`
  // {"en":"Associated hostname.","zh_CN":"关联的域名。"}
  ImportedDomain []*string `json:"importedDomain,omitempty" xml:"importedDomain,omitempty" require:"true" type:"Repeated"`
  // {"en":"Match conditions, At least one, at most five.","zh_CN":"匹配条件，至少一个，至多五个。"}
  Conditions *ListSharedWhitelistRulesResponseDataConditions `json:"conditions,omitempty" xml:"conditions,omitempty" require:"true" type:"Struct"`
  // {"en":"Created date, format: yyyy-MM-dd HH:mm:ss.","zh_CN":"创建时间，格式：yyyy-MM-dd HH:mm:ss。"}
  CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty" require:"true"`
  // {"en":"Update date,format: yyyy-MM-dd HH:mm:ss.","zh_CN":"更新时间，格式：yyyy-MM-dd HH:mm:ss。"}
  UpdateTime *string `json:"updateTime,omitempty" xml:"updateTime,omitempty" require:"true"`
}

func (s ListSharedWhitelistRulesResponseData) String() string {
  return tea.Prettify(s)
}

func (s ListSharedWhitelistRulesResponseData) GoString() string {
  return s.String()
}

func (s *ListSharedWhitelistRulesResponseData) SetId(v string) *ListSharedWhitelistRulesResponseData {
  s.Id = &v
  return s
}

func (s *ListSharedWhitelistRulesResponseData) SetRuleName(v string) *ListSharedWhitelistRulesResponseData {
  s.RuleName = &v
  return s
}

func (s *ListSharedWhitelistRulesResponseData) SetDescription(v string) *ListSharedWhitelistRulesResponseData {
  s.Description = &v
  return s
}

func (s *ListSharedWhitelistRulesResponseData) SetCreator(v string) *ListSharedWhitelistRulesResponseData {
  s.Creator = &v
  return s
}

func (s *ListSharedWhitelistRulesResponseData) SetImportedDomainCount(v int) *ListSharedWhitelistRulesResponseData {
  s.ImportedDomainCount = &v
  return s
}

func (s *ListSharedWhitelistRulesResponseData) SetImportedDomain(v []*string) *ListSharedWhitelistRulesResponseData {
  s.ImportedDomain = v
  return s
}

func (s *ListSharedWhitelistRulesResponseData) SetConditions(v *ListSharedWhitelistRulesResponseDataConditions) *ListSharedWhitelistRulesResponseData {
  s.Conditions = v
  return s
}

func (s *ListSharedWhitelistRulesResponseData) SetCreateTime(v string) *ListSharedWhitelistRulesResponseData {
  s.CreateTime = &v
  return s
}

func (s *ListSharedWhitelistRulesResponseData) SetUpdateTime(v string) *ListSharedWhitelistRulesResponseData {
  s.UpdateTime = &v
  return s
}

type ListSharedWhitelistRulesResponseDataConditions struct {
  // {"en":"IP/CIDR match conditions.","zh_CN":"IP/IP段匹配条件。"}
  IpOrIpsConditions []*ListSharedWhitelistRulesResponseDataConditionsIpOrIpsConditions `json:"ipOrIpsConditions,omitempty" xml:"ipOrIpsConditions,omitempty" require:"true" type:"Repeated"`
  // {"en":"Path match conditions.","zh_CN":"路径匹配条件。"}
  PathConditions []*ListSharedWhitelistRulesResponseDataConditionsPathConditions `json:"pathConditions,omitempty" xml:"pathConditions,omitempty" require:"true" type:"Repeated"`
  // {"en":"URI match conditions.","zh_CN":"URI匹配条件。"}
  UriConditions []*ListSharedWhitelistRulesResponseDataConditionsUriConditions `json:"uriConditions,omitempty" xml:"uriConditions,omitempty" require:"true" type:"Repeated"`
  // {"en":"User agent match conditions.","zh_CN":"User-Agent 匹配条件。"}
  UaConditions []*ListSharedWhitelistRulesResponseDataConditionsUaConditions `json:"uaConditions,omitempty" xml:"uaConditions,omitempty" require:"true" type:"Repeated"`
  // {"en":"Referer match conditions.","zh_CN":"Referer 匹配条件。"}
  RefererConditions []*ListSharedWhitelistRulesResponseDataConditionsRefererConditions `json:"refererConditions,omitempty" xml:"refererConditions,omitempty" require:"true" type:"Repeated"`
  // {"en":"Request header match conditions.","zh_CN":"请求头匹配条件。"}
  HeaderConditions []*ListSharedWhitelistRulesResponseDataConditionsHeaderConditions `json:"headerConditions,omitempty" xml:"headerConditions,omitempty" require:"true" type:"Repeated"`
}

func (s ListSharedWhitelistRulesResponseDataConditions) String() string {
  return tea.Prettify(s)
}

func (s ListSharedWhitelistRulesResponseDataConditions) GoString() string {
  return s.String()
}

func (s *ListSharedWhitelistRulesResponseDataConditions) SetIpOrIpsConditions(v []*ListSharedWhitelistRulesResponseDataConditionsIpOrIpsConditions) *ListSharedWhitelistRulesResponseDataConditions {
  s.IpOrIpsConditions = v
  return s
}

func (s *ListSharedWhitelistRulesResponseDataConditions) SetPathConditions(v []*ListSharedWhitelistRulesResponseDataConditionsPathConditions) *ListSharedWhitelistRulesResponseDataConditions {
  s.PathConditions = v
  return s
}

func (s *ListSharedWhitelistRulesResponseDataConditions) SetUriConditions(v []*ListSharedWhitelistRulesResponseDataConditionsUriConditions) *ListSharedWhitelistRulesResponseDataConditions {
  s.UriConditions = v
  return s
}

func (s *ListSharedWhitelistRulesResponseDataConditions) SetUaConditions(v []*ListSharedWhitelistRulesResponseDataConditionsUaConditions) *ListSharedWhitelistRulesResponseDataConditions {
  s.UaConditions = v
  return s
}

func (s *ListSharedWhitelistRulesResponseDataConditions) SetRefererConditions(v []*ListSharedWhitelistRulesResponseDataConditionsRefererConditions) *ListSharedWhitelistRulesResponseDataConditions {
  s.RefererConditions = v
  return s
}

func (s *ListSharedWhitelistRulesResponseDataConditions) SetHeaderConditions(v []*ListSharedWhitelistRulesResponseDataConditionsHeaderConditions) *ListSharedWhitelistRulesResponseDataConditions {
  s.HeaderConditions = v
  return s
}

type ListSharedWhitelistRulesResponseDataConditionsIpOrIpsConditions struct     {
  // {"en":"Match type.\nEQUAL: Equals\nNOT_EQUAL: Does not equal","zh_CN":"匹配类型。\nEQUAL：等于\nNOT_EQUAL：不等于","exampleValue":"EQUAL,NOT_EQUAL"}
  MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
  // {"en":"IP/CIDR.","zh_CN":"IP/IP段。"}
  IpOrIps []*string `json:"ipOrIps,omitempty" xml:"ipOrIps,omitempty" require:"true" type:"Repeated"`
}

func (s ListSharedWhitelistRulesResponseDataConditionsIpOrIpsConditions) String() string {
  return tea.Prettify(s)
}

func (s ListSharedWhitelistRulesResponseDataConditionsIpOrIpsConditions) GoString() string {
  return s.String()
}

func (s *ListSharedWhitelistRulesResponseDataConditionsIpOrIpsConditions) SetMatchType(v string) *ListSharedWhitelistRulesResponseDataConditionsIpOrIpsConditions {
  s.MatchType = &v
  return s
}

func (s *ListSharedWhitelistRulesResponseDataConditionsIpOrIpsConditions) SetIpOrIps(v []*string) *ListSharedWhitelistRulesResponseDataConditionsIpOrIpsConditions {
  s.IpOrIps = v
  return s
}

type ListSharedWhitelistRulesResponseDataConditionsPathConditions struct     {
  // {"en":"Match type.\nEQUAL: Equals\nNOT_EQUAL: Does not equal\nCONTAIN: Contains\nNOT_CONTAIN: Does not Contains\nREGEX: Regex match\nNOT_REGEX: regular does not match\nSTART_WITH: starts with\nEND_WITH: ends with\nWILDCARD: wildcard matches\nNOT_WILDCARD: wildcard does not match","zh_CN":"匹配类型。\nEQUAL：等于\nNOT_EQUAL：不等于\nCONTAIN：包含\nNOT_CONTAIN：不包含\nREGEX：匹配正则\nNOT_REGEX：正则不匹配\nSTART_WITH：开头是\nEND_WITH：结尾是\nWILDCARD：通配符匹配\nNOT_WILDCARD：通配符不匹配","exampleValue":"EQUAL,NOT_EQUAL,CONTAIN,NOT_CONTAIN,REGEX,NOT_REGEX,START_WITH,END_WITH,WILDCARD,NOT_WILDCARD"}
  MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
  // {"en":"Path.","zh_CN":"路径。"}
  ListSharedWhitelistRulesPaths []*string `json:"paths,omitempty" xml:"paths,omitempty" require:"true" type:"Repeated"`
}

func (s ListSharedWhitelistRulesResponseDataConditionsPathConditions) String() string {
  return tea.Prettify(s)
}

func (s ListSharedWhitelistRulesResponseDataConditionsPathConditions) GoString() string {
  return s.String()
}

func (s *ListSharedWhitelistRulesResponseDataConditionsPathConditions) SetMatchType(v string) *ListSharedWhitelistRulesResponseDataConditionsPathConditions {
  s.MatchType = &v
  return s
}

func (s *ListSharedWhitelistRulesResponseDataConditionsPathConditions) SetListSharedWhitelistRulesPaths(v []*string) *ListSharedWhitelistRulesResponseDataConditionsPathConditions {
  s.ListSharedWhitelistRulesPaths = v
  return s
}

type ListSharedWhitelistRulesResponseDataConditionsUriConditions struct     {
  // {"en":"Match type.\nEQUAL: Equals\nNOT_EQUAL: Does not equal\nCONTAIN: Contains\nNOT_CONTAIN: Does not Contains\nREGEX: Regex match\nNOT_REGEX: regular does not match\nSTART_WITH: starts with\nEND_WITH: ends with\nWILDCARD: wildcard matches\nNOT_WILDCARD: wildcard does not match","zh_CN":"匹配类型。\nEQUAL：等于\nNOT_EQUAL：不等于\nCONTAIN：包含\nNOT_CONTAIN：不包含\nREGEX：匹配正则\nNOT_REGEX：正则不匹配\nSTART_WITH：开头是\nEND_WITH：结尾是\nWILDCARD：通配符匹配\nNOT_WILDCARD：通配符不匹配","exampleValue":"EQUAL,NOT_EQUAL,CONTAIN,NOT_CONTAIN,REGEX,NOT_REGEX,START_WITH,END_WITH,WILDCARD,NOT_WILDCARD"}
  MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
  // {"en":"URI.","zh_CN":"URI。"}
  Uri []*string `json:"uri,omitempty" xml:"uri,omitempty" require:"true" type:"Repeated"`
}

func (s ListSharedWhitelistRulesResponseDataConditionsUriConditions) String() string {
  return tea.Prettify(s)
}

func (s ListSharedWhitelistRulesResponseDataConditionsUriConditions) GoString() string {
  return s.String()
}

func (s *ListSharedWhitelistRulesResponseDataConditionsUriConditions) SetMatchType(v string) *ListSharedWhitelistRulesResponseDataConditionsUriConditions {
  s.MatchType = &v
  return s
}

func (s *ListSharedWhitelistRulesResponseDataConditionsUriConditions) SetUri(v []*string) *ListSharedWhitelistRulesResponseDataConditionsUriConditions {
  s.Uri = v
  return s
}

type ListSharedWhitelistRulesResponseDataConditionsUaConditions struct     {
  // {"en":"Match type.\nEQUAL: Equals\nNOT_EQUAL: Does not equal\nCONTAIN: Contains\nNOT_CONTAIN: Does not Contains\nREGEX: Regex match\nNOT_REGEX: regular does not match\nSTART_WITH: starts with\nEND_WITH: ends with\nWILDCARD: wildcard matches\nNOT_WILDCARD: wildcard does not match","zh_CN":"匹配类型。\nEQUAL：等于\nNOT_EQUAL：不等于\nCONTAIN：包含\nNOT_CONTAIN：不包含\nREGEX：匹配正则\nNOT_REGEX：正则不匹配\nSTART_WITH：开头是\nEND_WITH：结尾是\nWILDCARD：通配符匹配\nNOT_WILDCARD：通配符不匹配","exampleValue":"EQUAL,NOT_EQUAL,CONTAIN,NOT_CONTAIN,REGEX,NOT_REGEX,START_WITH,END_WITH,WILDCARD,NOT_WILDCARD"}
  MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
  // {"en":"User agent.","zh_CN":"User-Agent。"}
  Ua []*string `json:"ua,omitempty" xml:"ua,omitempty" require:"true" type:"Repeated"`
}

func (s ListSharedWhitelistRulesResponseDataConditionsUaConditions) String() string {
  return tea.Prettify(s)
}

func (s ListSharedWhitelistRulesResponseDataConditionsUaConditions) GoString() string {
  return s.String()
}

func (s *ListSharedWhitelistRulesResponseDataConditionsUaConditions) SetMatchType(v string) *ListSharedWhitelistRulesResponseDataConditionsUaConditions {
  s.MatchType = &v
  return s
}

func (s *ListSharedWhitelistRulesResponseDataConditionsUaConditions) SetUa(v []*string) *ListSharedWhitelistRulesResponseDataConditionsUaConditions {
  s.Ua = v
  return s
}

type ListSharedWhitelistRulesResponseDataConditionsRefererConditions struct     {
  // {"en":"Match type.\nEQUAL: Equals\nNOT_EQUAL: Does not equal\nCONTAIN: Contains\nNOT_CONTAIN: Does not Contains\nREGEX: Regex match\nNOT_REGEX: regular does not match\nSTART_WITH: starts with\nEND_WITH: ends with\nWILDCARD: wildcard matches\nNOT_WILDCARD: wildcard does not match","zh_CN":"匹配类型。\nEQUAL：等于\nNOT_EQUAL：不等于\nCONTAIN：包含\nNOT_CONTAIN：不包含\nREGEX：匹配正则\nNOT_REGEX：正则不匹配\nSTART_WITH：开头是\nEND_WITH：结尾是\nWILDCARD：通配符匹配\nNOT_WILDCARD：通配符不匹配","exampleValue":"EQUAL,NOT_EQUAL,CONTAIN,NOT_CONTAIN,REGEX,NOT_REGEX,START_WITH,END_WITH,WILDCARD,NOT_WILDCARD"}
  MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
  // {"en":"Referer.","zh_CN":"Referer。"}
  Referer []*string `json:"referer,omitempty" xml:"referer,omitempty" require:"true" type:"Repeated"`
}

func (s ListSharedWhitelistRulesResponseDataConditionsRefererConditions) String() string {
  return tea.Prettify(s)
}

func (s ListSharedWhitelistRulesResponseDataConditionsRefererConditions) GoString() string {
  return s.String()
}

func (s *ListSharedWhitelistRulesResponseDataConditionsRefererConditions) SetMatchType(v string) *ListSharedWhitelistRulesResponseDataConditionsRefererConditions {
  s.MatchType = &v
  return s
}

func (s *ListSharedWhitelistRulesResponseDataConditionsRefererConditions) SetReferer(v []*string) *ListSharedWhitelistRulesResponseDataConditionsRefererConditions {
  s.Referer = v
  return s
}

type ListSharedWhitelistRulesResponseDataConditionsHeaderConditions struct     {
  // {"en":"Match type.\nEQUAL: Equals\nNOT_EQUAL: Does not equal\nCONTAIN: Contains\nNOT_CONTAIN: Does not Contains\nREGEX: Regex match\nNOT_REGEX: regular does not match\nSTART_WITH: starts with\nEND_WITH: ends with\nWILDCARD: wildcard matches\nNOT_WILDCARD: wildcard does not match","zh_CN":"匹配类型。\nEQUAL：等于\nNOT_EQUAL：不等于\nCONTAIN：包含\nNOT_CONTAIN：不包含\nREGEX：匹配正则\nNOT_REGEX：正则不匹配\nSTART_WITH：开头是\nEND_WITH：结尾是\nWILDCARD：通配符匹配\nNOT_WILDCARD：通配符不匹配","exampleValue":"EQUAL,NOT_EQUAL,CONTAIN,NOT_CONTAIN,REGEX,NOT_REGEX,START_WITH,END_WITH,WILDCARD,NOT_WILDCARD"}
  MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
  // {"en":"Request header name.","zh_CN":"头部名称。"}
  Key *string `json:"key,omitempty" xml:"key,omitempty" require:"true"`
  // {"en":"List of request header values.","zh_CN":"头部值列表。"}
  ValueList []*string `json:"valueList,omitempty" xml:"valueList,omitempty" require:"true" type:"Repeated"`
}

func (s ListSharedWhitelistRulesResponseDataConditionsHeaderConditions) String() string {
  return tea.Prettify(s)
}

func (s ListSharedWhitelistRulesResponseDataConditionsHeaderConditions) GoString() string {
  return s.String()
}

func (s *ListSharedWhitelistRulesResponseDataConditionsHeaderConditions) SetMatchType(v string) *ListSharedWhitelistRulesResponseDataConditionsHeaderConditions {
  s.MatchType = &v
  return s
}

func (s *ListSharedWhitelistRulesResponseDataConditionsHeaderConditions) SetKey(v string) *ListSharedWhitelistRulesResponseDataConditionsHeaderConditions {
  s.Key = &v
  return s
}

func (s *ListSharedWhitelistRulesResponseDataConditionsHeaderConditions) SetValueList(v []*string) *ListSharedWhitelistRulesResponseDataConditionsHeaderConditions {
  s.ValueList = v
  return s
}

type ListSharedWhitelistRulesResponseHeader struct {
}

func (s ListSharedWhitelistRulesResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ListSharedWhitelistRulesResponseHeader) GoString() string {
  return s.String()
}




type DisassociateShareCustomizeBotsRequest struct {
  // {"en":"Share configuration ID.","zh_CN":"共享配置ID。"}
  ShareId *string `json:"shareId,omitempty" xml:"shareId,omitempty" require:"true"`
  // {"en":"List of hostname to be disassociated.","zh_CN":"待取消关联域名列表。"}
  DomainList []*string `json:"domainList,omitempty" xml:"domainList,omitempty" require:"true" type:"Repeated"`
}

func (s DisassociateShareCustomizeBotsRequest) String() string {
  return tea.Prettify(s)
}

func (s DisassociateShareCustomizeBotsRequest) GoString() string {
  return s.String()
}

func (s *DisassociateShareCustomizeBotsRequest) SetShareId(v string) *DisassociateShareCustomizeBotsRequest {
  s.ShareId = &v
  return s
}

func (s *DisassociateShareCustomizeBotsRequest) SetDomainList(v []*string) *DisassociateShareCustomizeBotsRequest {
  s.DomainList = v
  return s
}

type DisassociateShareCustomizeBotsRequestHeader struct {
  // {"en":"Security service type. Please enter a specific service type, if you purchase multiple security services.","zh_CN":"安全服务类型。有使用多个不同的安全服务时，需要填写具体的服务类型。"}
  ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty"`
}

func (s DisassociateShareCustomizeBotsRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s DisassociateShareCustomizeBotsRequestHeader) GoString() string {
  return s.String()
}

func (s *DisassociateShareCustomizeBotsRequestHeader) SetServiceType(v string) *DisassociateShareCustomizeBotsRequestHeader {
  s.ServiceType = &v
  return s
}

type DisassociateShareCustomizeBotsPaths struct {
}

func (s DisassociateShareCustomizeBotsPaths) String() string {
  return tea.Prettify(s)
}

func (s DisassociateShareCustomizeBotsPaths) GoString() string {
  return s.String()
}

type DisassociateShareCustomizeBotsParameters struct {
}

func (s DisassociateShareCustomizeBotsParameters) String() string {
  return tea.Prettify(s)
}

func (s DisassociateShareCustomizeBotsParameters) GoString() string {
  return s.String()
}

type DisassociateShareCustomizeBotsResponse struct {
  // {"dictionary":"belong=WAAP-MS-Ext|dict=waap_retCodeEnum","en":"Please refer to the error code for exceptions.","zh_CN":"请参照错误码。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Description.","zh_CN":"描述信息。"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
}

func (s DisassociateShareCustomizeBotsResponse) String() string {
  return tea.Prettify(s)
}

func (s DisassociateShareCustomizeBotsResponse) GoString() string {
  return s.String()
}

func (s *DisassociateShareCustomizeBotsResponse) SetCode(v string) *DisassociateShareCustomizeBotsResponse {
  s.Code = &v
  return s
}

func (s *DisassociateShareCustomizeBotsResponse) SetMsg(v string) *DisassociateShareCustomizeBotsResponse {
  s.Msg = &v
  return s
}

type DisassociateShareCustomizeBotsResponseHeader struct {
}

func (s DisassociateShareCustomizeBotsResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s DisassociateShareCustomizeBotsResponseHeader) GoString() string {
  return s.String()
}




type CreateSharedRateLimitingRuleRequest struct {
  // {"en":"Rule Name, maximum 100 characters.Does not support # and &.","zh_CN":"规则名称，最多100个字符。不支持 # 和 &。"}
  RuleName *string `json:"ruleName,omitempty" xml:"ruleName,omitempty" require:"true"`
  // {"en":"Description, maximum 1000 characters.","zh_CN":"规则描述，最多1000个字符。"}
  Description *string `json:"description,omitempty" xml:"description,omitempty"`
  // {"en":"Client identifier.\nIP:Client IP\nIP_UA:Client IP and User-Agent\nCOOKIE:Cookie\nIP_COOKIE:Client IP and Cookie\nHEADER:Request Header\nWhen there is a status code in the matching condition,this client identifier is not supported.\nIP_HEADER:Client IP and Request Header\nWhen there is a status code in the matching condition,this client identifier is not supported .","zh_CN":"统计粒度。\nIP：客户端IP\nIP_UA：客户端IP和User-Agent\nCOOKIE：Cookie\nIP_COOKIE：客户端IP和Cookie\nHEADER：请求头，当匹配条件中存在状态码时不支持此统计粒度\nIP_HEADER：客户端IP和请求头，当匹配条件中存在状态码时不支持此统计粒度"}
  StatisticalItem *string `json:"statisticalItem,omitempty" xml:"statisticalItem,omitempty" require:"true"`
  // {"en":"Statistical key value.\nWhen the client identifier is cookie/header value, the corresponding key value needs to be entered.","zh_CN":"统计key值。\n当统计粒度cookie/header值，需要输入对应的key值。"}
  StatisticsKey *string `json:"statisticsKey,omitempty" xml:"statisticsKey,omitempty"`
  // {"en":"Statistics period, unit: seconds, the range is 1 - 3600.","zh_CN":"统计周期，单位：秒，范围为 1 - 3600。"}
  StatisticalPeriod *int `json:"statisticalPeriod,omitempty" xml:"statisticalPeriod,omitempty" require:"true"`
  // {"en":"Trigger threshold, unit: times.","zh_CN":"触发阈值，单位：次。"}
  TriggerThreshold *int `json:"triggerThreshold,omitempty" xml:"triggerThreshold,omitempty" require:"true"`
  // {"en":"Action duration, unit: seconds, the range is 10 - 604800.","zh_CN":"处理动作持续时间，单位：秒，范围为 10 - 604800。"}
  InterceptTime *int `json:"interceptTime,omitempty" xml:"interceptTime,omitempty" require:"true"`
  // {"en":"Cycle effective status.\nPERMANENT:All time\nWITHOUT:Excluded time\nWITHIN:Selected time","zh_CN":"周期生效状态。\nPERMANENT：永久生效\nWITHOUT：周期内不生效\nWITHIN：周期内生效"}
  EffectiveStatus *string `json:"effectiveStatus,omitempty" xml:"effectiveStatus,omitempty" require:"true"`
  // {"en":"Effective time period.\nWhen the effective status is effective within the cycle or not effective within the cycle, this field must have a value.","zh_CN":"规则生效周期。\n生效状态为周期内生效或周期内不生效时，此字段必须有值。"}
  RateLimitEffective *CreateSharedRateLimitingRuleRequestRateLimitEffective `json:"rateLimitEffective,omitempty" xml:"rateLimitEffective,omitempty" type:"Struct"`
  // {"en":"Action.\nNO_USE:Not Used\nLOG:Log\nCOOKIE:Cookie verification\nJS_CHECK:Javascript verification\nDELAY:Delay\nBLOCK:Deny\nRESET:Reset Connection\nCUSTOM_ACTION_ID:Fill in the custom action id of the corresponding action\nWhen there is a status code in the matching condition, the supported actions are Log, Deny, and Reset Connection.","zh_CN":"处理动作。\nNO_USE：不使用\nLOG：监控\nCOOKIE：Cookie校验\nJS_CHECK：JavaScript校验\nDELAY：延迟响应\nBLOCK：拦截\nRESET：断开连接\nCUSTOM_ACTION_ID：传入对应的自定义响应动作的ID\n当匹配条件中存在状态码时，支持处理动作为监控、拦截、断开连接。"}
  Action *string `json:"action,omitempty" xml:"action,omitempty" require:"true"`
  // {"en":"Matching conditions.","zh_CN":"匹配条件。"}
  RateLimitRuleCondition *CreateSharedRateLimitingRuleRequestRateLimitRuleCondition `json:"rateLimitRuleCondition,omitempty" xml:"rateLimitRuleCondition,omitempty" require:"true" type:"Struct"`
}

func (s CreateSharedRateLimitingRuleRequest) String() string {
  return tea.Prettify(s)
}

func (s CreateSharedRateLimitingRuleRequest) GoString() string {
  return s.String()
}

func (s *CreateSharedRateLimitingRuleRequest) SetRuleName(v string) *CreateSharedRateLimitingRuleRequest {
  s.RuleName = &v
  return s
}

func (s *CreateSharedRateLimitingRuleRequest) SetDescription(v string) *CreateSharedRateLimitingRuleRequest {
  s.Description = &v
  return s
}

func (s *CreateSharedRateLimitingRuleRequest) SetStatisticalItem(v string) *CreateSharedRateLimitingRuleRequest {
  s.StatisticalItem = &v
  return s
}

func (s *CreateSharedRateLimitingRuleRequest) SetStatisticsKey(v string) *CreateSharedRateLimitingRuleRequest {
  s.StatisticsKey = &v
  return s
}

func (s *CreateSharedRateLimitingRuleRequest) SetStatisticalPeriod(v int) *CreateSharedRateLimitingRuleRequest {
  s.StatisticalPeriod = &v
  return s
}

func (s *CreateSharedRateLimitingRuleRequest) SetTriggerThreshold(v int) *CreateSharedRateLimitingRuleRequest {
  s.TriggerThreshold = &v
  return s
}

func (s *CreateSharedRateLimitingRuleRequest) SetInterceptTime(v int) *CreateSharedRateLimitingRuleRequest {
  s.InterceptTime = &v
  return s
}

func (s *CreateSharedRateLimitingRuleRequest) SetEffectiveStatus(v string) *CreateSharedRateLimitingRuleRequest {
  s.EffectiveStatus = &v
  return s
}

func (s *CreateSharedRateLimitingRuleRequest) SetRateLimitEffective(v *CreateSharedRateLimitingRuleRequestRateLimitEffective) *CreateSharedRateLimitingRuleRequest {
  s.RateLimitEffective = v
  return s
}

func (s *CreateSharedRateLimitingRuleRequest) SetAction(v string) *CreateSharedRateLimitingRuleRequest {
  s.Action = &v
  return s
}

func (s *CreateSharedRateLimitingRuleRequest) SetRateLimitRuleCondition(v *CreateSharedRateLimitingRuleRequestRateLimitRuleCondition) *CreateSharedRateLimitingRuleRequest {
  s.RateLimitRuleCondition = v
  return s
}

type CreateSharedRateLimitingRuleRequestRateLimitEffective struct {
  // {"en":"Effective.\nMON:Monday\nTUE:Tuesday\nWED:Wednesday\nTHU:Thursday\nFRI:Friday\nSAT:Saturday\nSUN:Sunday","zh_CN":"周期。\nMON：星期一\nTUE：星期二\nWED：星期三\nTHU：星期四\nFRI：星期五\nSAT：星期六\nSUN：星期天"}
  Effective []*string `json:"effective,omitempty" xml:"effective,omitempty" require:"true" type:"Repeated"`
  // {"en":"Start time, format: HH:mm.","zh_CN":"开始时间，格式：HH:mm。"}
  Start *string `json:"start,omitempty" xml:"start,omitempty" require:"true"`
  // {"en":"End time, format: HH:mm.","zh_CN":"结束时间，格式：HH:mm。"}
  End *string `json:"end,omitempty" xml:"end,omitempty" require:"true"`
  // {"dictionary":"belong=WAAP-MS-Ext|dict=waap_timezone","en":"Timezone,default value: GMT+8.","zh_CN":"时区，默认：GMT+8。"}
  Timezone *string `json:"timezone,omitempty" xml:"timezone,omitempty" require:"true"`
}

func (s CreateSharedRateLimitingRuleRequestRateLimitEffective) String() string {
  return tea.Prettify(s)
}

func (s CreateSharedRateLimitingRuleRequestRateLimitEffective) GoString() string {
  return s.String()
}

func (s *CreateSharedRateLimitingRuleRequestRateLimitEffective) SetEffective(v []*string) *CreateSharedRateLimitingRuleRequestRateLimitEffective {
  s.Effective = v
  return s
}

func (s *CreateSharedRateLimitingRuleRequestRateLimitEffective) SetStart(v string) *CreateSharedRateLimitingRuleRequestRateLimitEffective {
  s.Start = &v
  return s
}

func (s *CreateSharedRateLimitingRuleRequestRateLimitEffective) SetEnd(v string) *CreateSharedRateLimitingRuleRequestRateLimitEffective {
  s.End = &v
  return s
}

func (s *CreateSharedRateLimitingRuleRequestRateLimitEffective) SetTimezone(v string) *CreateSharedRateLimitingRuleRequestRateLimitEffective {
  s.Timezone = &v
  return s
}

type CreateSharedRateLimitingRuleRequestRateLimitRuleCondition struct {
  // {"en":"IP/CIDR, match type cannot be repeated.","zh_CN":"IP/IP段，匹配类型不可重复。"}
  IpOrIpsConditions []*CreateSharedRateLimitingRuleRequestRateLimitRuleConditionIpOrIpsConditions `json:"ipOrIpsConditions,omitempty" xml:"ipOrIpsConditions,omitempty" type:"Repeated"`
  // {"en":"Path, match type cannot be repeated.\nWhen the business scenario is API, this matching condition is not supported.","zh_CN":"路径，匹配类型不可重复。\n当业务场景为API业务时不支持此匹配条件。"}
  PathConditions []*CreateSharedRateLimitingRuleRequestRateLimitRuleConditionPathConditions `json:"pathConditions,omitempty" xml:"pathConditions,omitempty" type:"Repeated"`
  // {"en":"URI, match type cannot be repeated.\nWhen the business scenario is API, this matching condition is not supported.","zh_CN":"URI，匹配类型不可重复。\n当业务场景为API业务时不支持此匹配条件。"}
  UriConditions []*CreateSharedRateLimitingRuleRequestRateLimitRuleConditionUriConditions `json:"uriConditions,omitempty" xml:"uriConditions,omitempty" type:"Repeated"`
  // {"en":"User Agent, match type cannot be repeated.","zh_CN":"User-Agent，匹配类型不可重复。"}
  UaConditions []*CreateSharedRateLimitingRuleRequestRateLimitRuleConditionUaConditions `json:"uaConditions,omitempty" xml:"uaConditions,omitempty" type:"Repeated"`
  // {"en":"Request Method.\nWhen the business scenario is API,this matching condition is not supported.","zh_CN":"请求方法，匹配类型不可重复。\n当业务场景为API业务时不支持此匹配条件。"}
  MethodConditions []*CreateSharedRateLimitingRuleRequestRateLimitRuleConditionMethodConditions `json:"methodConditions,omitempty" xml:"methodConditions,omitempty" type:"Repeated"`
  // {"en":"Referer, match type cannot be repeated.","zh_CN":"Referer，匹配类型不可重复。"}
  RefererConditions []*CreateSharedRateLimitingRuleRequestRateLimitRuleConditionRefererConditions `json:"refererConditions,omitempty" xml:"refererConditions,omitempty" type:"Repeated"`
  // {"en":"Request Header, match type can be repeated.","zh_CN":"请求头，匹配类型可重复。"}
  HeaderConditions []*CreateSharedRateLimitingRuleRequestRateLimitRuleConditionHeaderConditions `json:"headerConditions,omitempty" xml:"headerConditions,omitempty" type:"Repeated"`
  // {"en":"Geo, match type cannot be repeated.","zh_CN":"区域，匹配类型不可重复。"}
  AreaConditions []*CreateSharedRateLimitingRuleRequestRateLimitRuleConditionAreaConditions `json:"areaConditions,omitempty" xml:"areaConditions,omitempty" type:"Repeated"`
  // {"en":"HTTP/S, match type cannot be repeated.","zh_CN":"应用层协议，匹配类型不可重复。"}
  SchemeConditions []*CreateSharedRateLimitingRuleRequestRateLimitRuleConditionSchemeConditions `json:"schemeConditions,omitempty" xml:"schemeConditions,omitempty" type:"Repeated"`
  // {"en":"Response Code, match type cannot be repeated.","zh_CN":"状态码，匹配类型不可重复。"}
  StatusCodeConditions []*CreateSharedRateLimitingRuleRequestRateLimitRuleConditionStatusCodeConditions `json:"statusCodeConditions,omitempty" xml:"statusCodeConditions,omitempty" type:"Repeated"`
  // {"en":"JA3 Fingerprint, match type cannot be repeated.","zh_CN":"JA3指纹，匹配类型不可重复。"}
  Ja3Conditions []*CreateSharedRateLimitingRuleRequestRateLimitRuleConditionJa3Conditions `json:"ja3Conditions,omitempty" xml:"ja3Conditions,omitempty" type:"Repeated"`
  // {"en":"JA4 Fingerprint, match type cannot be repeated.","zh_CN":"JA4指纹，匹配类型不可重复。"}
  Ja4Conditions []*CreateSharedRateLimitingRuleRequestRateLimitRuleConditionJa4Conditions `json:"ja4Conditions,omitempty" xml:"ja4Conditions,omitempty" type:"Repeated"`
}

func (s CreateSharedRateLimitingRuleRequestRateLimitRuleCondition) String() string {
  return tea.Prettify(s)
}

func (s CreateSharedRateLimitingRuleRequestRateLimitRuleCondition) GoString() string {
  return s.String()
}

func (s *CreateSharedRateLimitingRuleRequestRateLimitRuleCondition) SetIpOrIpsConditions(v []*CreateSharedRateLimitingRuleRequestRateLimitRuleConditionIpOrIpsConditions) *CreateSharedRateLimitingRuleRequestRateLimitRuleCondition {
  s.IpOrIpsConditions = v
  return s
}

func (s *CreateSharedRateLimitingRuleRequestRateLimitRuleCondition) SetPathConditions(v []*CreateSharedRateLimitingRuleRequestRateLimitRuleConditionPathConditions) *CreateSharedRateLimitingRuleRequestRateLimitRuleCondition {
  s.PathConditions = v
  return s
}

func (s *CreateSharedRateLimitingRuleRequestRateLimitRuleCondition) SetUriConditions(v []*CreateSharedRateLimitingRuleRequestRateLimitRuleConditionUriConditions) *CreateSharedRateLimitingRuleRequestRateLimitRuleCondition {
  s.UriConditions = v
  return s
}

func (s *CreateSharedRateLimitingRuleRequestRateLimitRuleCondition) SetUaConditions(v []*CreateSharedRateLimitingRuleRequestRateLimitRuleConditionUaConditions) *CreateSharedRateLimitingRuleRequestRateLimitRuleCondition {
  s.UaConditions = v
  return s
}

func (s *CreateSharedRateLimitingRuleRequestRateLimitRuleCondition) SetMethodConditions(v []*CreateSharedRateLimitingRuleRequestRateLimitRuleConditionMethodConditions) *CreateSharedRateLimitingRuleRequestRateLimitRuleCondition {
  s.MethodConditions = v
  return s
}

func (s *CreateSharedRateLimitingRuleRequestRateLimitRuleCondition) SetRefererConditions(v []*CreateSharedRateLimitingRuleRequestRateLimitRuleConditionRefererConditions) *CreateSharedRateLimitingRuleRequestRateLimitRuleCondition {
  s.RefererConditions = v
  return s
}

func (s *CreateSharedRateLimitingRuleRequestRateLimitRuleCondition) SetHeaderConditions(v []*CreateSharedRateLimitingRuleRequestRateLimitRuleConditionHeaderConditions) *CreateSharedRateLimitingRuleRequestRateLimitRuleCondition {
  s.HeaderConditions = v
  return s
}

func (s *CreateSharedRateLimitingRuleRequestRateLimitRuleCondition) SetAreaConditions(v []*CreateSharedRateLimitingRuleRequestRateLimitRuleConditionAreaConditions) *CreateSharedRateLimitingRuleRequestRateLimitRuleCondition {
  s.AreaConditions = v
  return s
}

func (s *CreateSharedRateLimitingRuleRequestRateLimitRuleCondition) SetSchemeConditions(v []*CreateSharedRateLimitingRuleRequestRateLimitRuleConditionSchemeConditions) *CreateSharedRateLimitingRuleRequestRateLimitRuleCondition {
  s.SchemeConditions = v
  return s
}

func (s *CreateSharedRateLimitingRuleRequestRateLimitRuleCondition) SetStatusCodeConditions(v []*CreateSharedRateLimitingRuleRequestRateLimitRuleConditionStatusCodeConditions) *CreateSharedRateLimitingRuleRequestRateLimitRuleCondition {
  s.StatusCodeConditions = v
  return s
}

func (s *CreateSharedRateLimitingRuleRequestRateLimitRuleCondition) SetJa3Conditions(v []*CreateSharedRateLimitingRuleRequestRateLimitRuleConditionJa3Conditions) *CreateSharedRateLimitingRuleRequestRateLimitRuleCondition {
  s.Ja3Conditions = v
  return s
}

func (s *CreateSharedRateLimitingRuleRequestRateLimitRuleCondition) SetJa4Conditions(v []*CreateSharedRateLimitingRuleRequestRateLimitRuleConditionJa4Conditions) *CreateSharedRateLimitingRuleRequestRateLimitRuleCondition {
  s.Ja4Conditions = v
  return s
}

type CreateSharedRateLimitingRuleRequestRateLimitRuleConditionIpOrIpsConditions struct     {
  // {"en":"Match type.\nEQUAL:Equals\nNOT_EQUAL:Does not equal","zh_CN":"匹配类型。\nEQUAL：等于\nNOT_EQUAL：不等于","exampleValue":"EQUAL,NOT_EQUAL"}
  MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
  // {"en":"IP/CIDR, maximum 500 IP/CIDR.","zh_CN":"IP/IP段，最多500个IP/IP段。"}
  IpOrIps []*string `json:"ipOrIps,omitempty" xml:"ipOrIps,omitempty" require:"true" type:"Repeated"`
}

func (s CreateSharedRateLimitingRuleRequestRateLimitRuleConditionIpOrIpsConditions) String() string {
  return tea.Prettify(s)
}

func (s CreateSharedRateLimitingRuleRequestRateLimitRuleConditionIpOrIpsConditions) GoString() string {
  return s.String()
}

func (s *CreateSharedRateLimitingRuleRequestRateLimitRuleConditionIpOrIpsConditions) SetMatchType(v string) *CreateSharedRateLimitingRuleRequestRateLimitRuleConditionIpOrIpsConditions {
  s.MatchType = &v
  return s
}

func (s *CreateSharedRateLimitingRuleRequestRateLimitRuleConditionIpOrIpsConditions) SetIpOrIps(v []*string) *CreateSharedRateLimitingRuleRequestRateLimitRuleConditionIpOrIpsConditions {
  s.IpOrIps = v
  return s
}

type CreateSharedRateLimitingRuleRequestRateLimitRuleConditionPathConditions struct     {
  // {"en":"Match type.\nEQUAL: Equals, path case sensitive\nNOT_EQUAL: Does not equal, path case sensitive\nCONTAIN: Contains, path case insensitive\nNOT_CONTAIN: Does not Contains, path case insensitive\nREGEX: Regex match, path case insensitive\nNOT_REGEX: Regular does not match, path case sensitive\nSTART_WITH: Starts with, path case sensitive\nEND_WITH: Ends with, path case sensitive\nWILDCARD: Wildcard matches, path case sensitive,* represents zero or more arbitrary characters, ? represents any single character.\nNOT_WILDCARD: Wildcard does not match, path case sensitive,* represents zero or more arbitrary characters, ? represents any single character","zh_CN":"匹配类型。\nEQUAL：等于，路径大小写敏感\nNOT_EQUAL：不等于，路径大小写敏感\nCONTAIN：包含，路径大小写不敏感\nNOT_CONTAIN：不包含，路径大小写不敏感\nREGEX：匹配正则，路径大小写不敏感\nNOT_REGEX：正则不匹配，路径大小写不敏感\nSTART_WITH：开头是，路径大小写不敏感\nEND_WITH：结尾是，路径大小写不敏感\nWILDCARD：通配符匹配，路径大小写不敏感，*代表零个或多个任意字符，?代表任意单个字符\nNOT_WILDCARD：通配符不匹配，路径大小写不敏感，*代表零个或多个任意字符，?代表任意单个字符","exampleValue":"EQUAL,NOT_EQUAL,CONTAIN,NOT_CONTAIN,REGEX,NOT_REGEX,START_WITH,END_WITH,WILDCARD,NOT_WILDCARD"}
  MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
  // {"en":"Path.\nWhen match type is EQUAL/NOT_EQUAL/START_WITH/END_WITH, path needs to start with \"/\", and no parameters.\nWhen the match type is REGEX/NOT_REGEX, only one value is allowed.\nExample: /test.html.","zh_CN":"路径。\n当匹配类型为等于/不等于/开头是/结尾是，路径必须以“/”开头，不含参数。\n当匹配类型为正则/正则不匹配，则只允许只有一个值。\n示例：/test.html。"}
  CreateSharedRateLimitingRulePaths []*string `json:"paths,omitempty" xml:"paths,omitempty" require:"true" type:"Repeated"`
}

func (s CreateSharedRateLimitingRuleRequestRateLimitRuleConditionPathConditions) String() string {
  return tea.Prettify(s)
}

func (s CreateSharedRateLimitingRuleRequestRateLimitRuleConditionPathConditions) GoString() string {
  return s.String()
}

func (s *CreateSharedRateLimitingRuleRequestRateLimitRuleConditionPathConditions) SetMatchType(v string) *CreateSharedRateLimitingRuleRequestRateLimitRuleConditionPathConditions {
  s.MatchType = &v
  return s
}

func (s *CreateSharedRateLimitingRuleRequestRateLimitRuleConditionPathConditions) SetCreateSharedRateLimitingRulePaths(v []*string) *CreateSharedRateLimitingRuleRequestRateLimitRuleConditionPathConditions {
  s.CreateSharedRateLimitingRulePaths = v
  return s
}

type CreateSharedRateLimitingRuleRequestRateLimitRuleConditionUriConditions struct     {
  // {"en":"Match type.\nEQUAL: Equals, URI case sensitive\nNOT_EQUAL: Does not equal, URI case sensitive\nCONTAIN: Contains, URI case insensitive\nNOT_CONTAIN: Does not Contains, URI case insensitive\nREGEX: Regex match, URI case insensitive\nNOT_REGEX: Regular does not match, URI case insensitive\nSTART_WITH: Starts with, URI case insensitive\nEND_WITH: Ends with, URI case insensitive\nWILDCARD: Wildcard matches, URI case insensitive,* represents zero or more arbitrary characters, ? represents any single character\nNOT_WILDCARD: Wildcard does not match, URI case insensitive,* represents zero or more arbitrary characters, ? represents any single character","zh_CN":"匹配类型。\nEQUAL：等于，URI大小写敏感\nNOT_EQUAL：不等于，URI大小写敏感\nCONTAIN：包含，URI大小写不敏感\nNOT_CONTAIN：不包含，URI大小写不敏感\nREGEX：匹配正则，URI大小写不敏感\nNOT_REGEX：正则不匹配，URI大小写不敏感\nSTART_WITH：开头是，URI大小写不敏感\nEND_WITH：结尾是，URI大小写不敏感\nWILDCARD：通配符匹配，URI大小写不敏感，*代表零个或多个任意字符，?代表任意单个字符\nNOT_WILDCARD：通配符不匹配，URI大小写不敏感，*代表零个或多个任意字符，?代表任意单个字符","exampleValue":"EQUAL,NOT_EQUAL,CONTAIN,NOT_CONTAIN,REGEX,NOT_REGEX,START_WITH,END_WITH,WILDCARD,NOT_WILDCARD"}
  MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
  // {"en":"URI.\nWhen match type is EQUAL/NOT_EQUAL/START_WITH/END_WITH, uri needs to start with \"/\", and includes parameters.\nWhen the match type is REGEX/NOT_REGEX, only one value is allowed.\nExample: /test.html?id=1.","zh_CN":"URI。\n当匹配类型为等于/不等于/开头是/结尾是，URI必须以”/“开头，含参数。\n当匹配类型为正则/正则不匹配，则只允许只有一个值。\n示例：/test.html?id=1。"}
  Uri []*string `json:"uri,omitempty" xml:"uri,omitempty" require:"true" type:"Repeated"`
}

func (s CreateSharedRateLimitingRuleRequestRateLimitRuleConditionUriConditions) String() string {
  return tea.Prettify(s)
}

func (s CreateSharedRateLimitingRuleRequestRateLimitRuleConditionUriConditions) GoString() string {
  return s.String()
}

func (s *CreateSharedRateLimitingRuleRequestRateLimitRuleConditionUriConditions) SetMatchType(v string) *CreateSharedRateLimitingRuleRequestRateLimitRuleConditionUriConditions {
  s.MatchType = &v
  return s
}

func (s *CreateSharedRateLimitingRuleRequestRateLimitRuleConditionUriConditions) SetUri(v []*string) *CreateSharedRateLimitingRuleRequestRateLimitRuleConditionUriConditions {
  s.Uri = v
  return s
}

type CreateSharedRateLimitingRuleRequestRateLimitRuleConditionUaConditions struct     {
  // {"en":"Match type.\nEQUAL: Equals, user agent case sensitive\nNOT_EQUAL: Does not equal, user agent case sensitive\nCONTAIN: Contains, user agent case insensitive\nNOT_CONTAIN: Does not Contains, user agent case insensitive\nNONE:Empty or non-existent\nREGEX: Regex match, user agent case insensitive\nNOT_REGEX: Regular does not match, user agent case insensitive\nSTART_WITH: Starts with, user agent case insensitive\nEND_WITH: Ends with, user agent case insensitive\nWILDCARD: Wildcard matches, user agent case insensitive,* represents zero or more arbitrary characters, ? represents any single character\nNOT_WILDCARD: Wildcard does not match, user agent case insensitive,* represents zero or more arbitrary characters, ? represents any single character","zh_CN":"匹配类型。\nEQUAL：等于，User-Agent大小写敏感\nNOT_EQUAL：不等于，User-Agent大小写敏感\nCONTAIN：包含，User-Agent大小写不敏感\nNOT_CONTAIN：不包含，User-Agent大小写不敏感\nNONE：为空或不存在\nREGEX：匹配正则，User-Agent大小写不敏感\nNOT_REGEX：正则不匹配，User-Agent大小写不敏感\nSTART_WITH：开头是，User-Agent大小写不敏感\nEND_WITH：结尾是，User-Agent大小写不敏感\nWILDCARD：通配符匹配，User-Agent大小写不敏感，*代表零个或多个任意字符，?代表任意单个字符\nNOT_WILDCARD：通配符不匹配，User-Agent大小写不敏感，*代表零个或多个任意字符，?代表任意单个字符","exampleValue":"EQUAL,NOT_EQUAL,CONTAIN,NOT_CONTAIN,NONE,REGEX,NOT_REGEX,START_WITH,END_WITH,WILDCARD,NOT_WILDCARD"}
  MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
  // {"en":"User agent.\nWhen the match type is REGEX/NOT_REGEX, only one value is allowed.\nExample: go-Http-client/1.1.","zh_CN":"User-Agent。\n当匹配类型为正则/正则不匹配，则只允许只有一个值。\n示例：go-Http-client/1.1。"}
  Ua []*string `json:"ua,omitempty" xml:"ua,omitempty" require:"true" type:"Repeated"`
}

func (s CreateSharedRateLimitingRuleRequestRateLimitRuleConditionUaConditions) String() string {
  return tea.Prettify(s)
}

func (s CreateSharedRateLimitingRuleRequestRateLimitRuleConditionUaConditions) GoString() string {
  return s.String()
}

func (s *CreateSharedRateLimitingRuleRequestRateLimitRuleConditionUaConditions) SetMatchType(v string) *CreateSharedRateLimitingRuleRequestRateLimitRuleConditionUaConditions {
  s.MatchType = &v
  return s
}

func (s *CreateSharedRateLimitingRuleRequestRateLimitRuleConditionUaConditions) SetUa(v []*string) *CreateSharedRateLimitingRuleRequestRateLimitRuleConditionUaConditions {
  s.Ua = v
  return s
}

type CreateSharedRateLimitingRuleRequestRateLimitRuleConditionMethodConditions struct     {
  // {"en":"Match type.\nEQUAL:Equals\nNOT_EQUAL:Does not equal","zh_CN":"匹配类型。\nEQUAL：等于\nNOT_EQUAL：不等于","exampleValue":"EQUAL,NOT_EQUAL"}
  MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
  // {"en":"Request method.\nSupported values: GET/POST/DELETE/PUT/HEAD/OPTIONS/COPY.","zh_CN":"请求方法。\n支持的值：GET/POST/DELETE/PUT/HEAD/OPTIONS/COPY。","exampleValue":"GET,POST,DELETE,PUT,HEAD,OPTIONS,COPY"}
  RequestMethod []*string `json:"requestMethod,omitempty" xml:"requestMethod,omitempty" require:"true" type:"Repeated"`
}

func (s CreateSharedRateLimitingRuleRequestRateLimitRuleConditionMethodConditions) String() string {
  return tea.Prettify(s)
}

func (s CreateSharedRateLimitingRuleRequestRateLimitRuleConditionMethodConditions) GoString() string {
  return s.String()
}

func (s *CreateSharedRateLimitingRuleRequestRateLimitRuleConditionMethodConditions) SetMatchType(v string) *CreateSharedRateLimitingRuleRequestRateLimitRuleConditionMethodConditions {
  s.MatchType = &v
  return s
}

func (s *CreateSharedRateLimitingRuleRequestRateLimitRuleConditionMethodConditions) SetRequestMethod(v []*string) *CreateSharedRateLimitingRuleRequestRateLimitRuleConditionMethodConditions {
  s.RequestMethod = v
  return s
}

type CreateSharedRateLimitingRuleRequestRateLimitRuleConditionRefererConditions struct     {
  // {"en":"Referer.\nWhen the match type is REGEX/NOT_REGEX, only one value is allowed.\nExample: http://test.com.","zh_CN":"Referer。\n当匹配类型为正则/正则不匹配，则只允许只有一个值。\n示例：http://test.com。"}
  Referer []*string `json:"referer,omitempty" xml:"referer,omitempty" require:"true" type:"Repeated"`
  // {"en":"Match type.\nEQUAL: Equals, referer case sensitive\nNOT_EQUAL: Does not equal, referer case sensitive\nCONTAIN: Contains, referer case insensitive\nNOT_CONTAIN: Does not Contains, referer case insensitive\nNONE:Empty or non-existent\nREGEX: Regex match, referer case insensitive\nNOT_REGEX: Regular does not match, referer case insensitive\nSTART_WITH: Starts with, referer case insensitive\nEND_WITH: Ends with, referer case insensitive\nWILDCARD: Wildcard matches, referer case insensitive,* represents zero or more arbitrary characters, ? represents any single characte\nNOT_WILDCARD: Wildcard does not match, referer case insensitive,* represents zero or more arbitrary characters, ? represents any single character","zh_CN":"匹配类型。\nEQUAL：等于，referer大小写敏感\nNOT_EQUAL：不等于，referer大小写敏感\nCONTAIN：包含，referer大小写不敏感\nNOT_CONTAIN：不包含，referer大小写不敏感\nNONE：为空或不存在\nREGEX：匹配正则，referer大小写不敏感\nNOT_REGEX：正则不匹配，referer大小写不敏感\nSTART_WITH：开头是，referer大小写不敏感\nEND_WITH：结尾是，referer大小写不敏感\nWILDCARD：通配符匹配，referer大小写不敏感，*代表零个或多个任意字符，?代表任意单个字符\nNOT_WILDCARD：通配符不匹配，referer大小写不敏感，*代表零个或多个任意字符，?代表任意单个字符","exampleValue":"EQUAL,NOT_EQUAL,CONTAIN,NOT_CONTAIN,NONE,REGEX,NOT_REGEX,START_WITH,END_WITH,WILDCARD,NOT_WILDCARD"}
  MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
}

func (s CreateSharedRateLimitingRuleRequestRateLimitRuleConditionRefererConditions) String() string {
  return tea.Prettify(s)
}

func (s CreateSharedRateLimitingRuleRequestRateLimitRuleConditionRefererConditions) GoString() string {
  return s.String()
}

func (s *CreateSharedRateLimitingRuleRequestRateLimitRuleConditionRefererConditions) SetReferer(v []*string) *CreateSharedRateLimitingRuleRequestRateLimitRuleConditionRefererConditions {
  s.Referer = v
  return s
}

func (s *CreateSharedRateLimitingRuleRequestRateLimitRuleConditionRefererConditions) SetMatchType(v string) *CreateSharedRateLimitingRuleRequestRateLimitRuleConditionRefererConditions {
  s.MatchType = &v
  return s
}

type CreateSharedRateLimitingRuleRequestRateLimitRuleConditionHeaderConditions struct     {
  // {"en":"Match type.EQUAL: Equals, request header values case sensitiveNOT_EQUAL: Does not equal, request header values case sensitiveCONTAIN: Contains, request header values case insensitiveNOT_CONTAIN: Does not Contains, request header values case insensitiveNONE: Empty or non-existentREGEX: Regex match, request header values case insensitiveNOT_REGEX: Regular does not match, request header values case insensitiveSTART_WITH: Starts with, request header values case insensitiveEND_WITH: Ends with, request header values case insensitiveWILDCARD: Wildcard matches, request header values case insensitive,* represents zero or more arbitrary characters, ? represents any single characterNOT_WILDCARD: Wildcard does not match, request header values case insensitive,* represents zero or more arbitrary characters, ? represents any single character","zh_CN":"匹配类型。EQUAL：等于，头部值大小写敏感NOT_EQUAL：不等于，头部值大小写敏感CONTAIN：包含，头部值大小写不敏感NOT_CONTAIN：不包含，头部值大小写不敏感NONE：为空或不存在REGEX：匹配正则，头部值大小写不敏感NOT_REGEX：正则不匹配，头部值大小写不敏感START_WITH：开头是，头部值大小写不敏感END_WITH：结尾是，头部值大小写不敏感WILDCARD：通配符匹配，头部值大小写不敏感，*代表零个或多个任意字符，?代表任意单个字符NOT_WILDCARD：通配符不匹配，头部值大小写不敏感，*代表零个或多个任意字符，?代表任意单个字符","exampleValue":"EQUAL,NOT_EQUAL,CONTAIN,NOT_CONTAIN,NONE,REGEX,NOT_REGEX,START_WITH,END_WITH,WILDCARD,NOT_WILDCARD"}
  MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
  // {"en":"Header value.\nWhen the match type is REGEX/NOT_REGEX, only one value is allowed.","zh_CN":"头部值。\n当匹配类型为正则/正则不匹配，则只允许只有一个值。"}
  ValueList []*string `json:"valueList,omitempty" xml:"valueList,omitempty" require:"true" type:"Repeated"`
  // {"en":"Header name,case insensitive,up to 100 characters.\nExample: Accept.","zh_CN":"头部名称，大小写不敏感，最多100个字符。\n示例：Accept。"}
  Key *string `json:"key,omitempty" xml:"key,omitempty" require:"true"`
}

func (s CreateSharedRateLimitingRuleRequestRateLimitRuleConditionHeaderConditions) String() string {
  return tea.Prettify(s)
}

func (s CreateSharedRateLimitingRuleRequestRateLimitRuleConditionHeaderConditions) GoString() string {
  return s.String()
}

func (s *CreateSharedRateLimitingRuleRequestRateLimitRuleConditionHeaderConditions) SetMatchType(v string) *CreateSharedRateLimitingRuleRequestRateLimitRuleConditionHeaderConditions {
  s.MatchType = &v
  return s
}

func (s *CreateSharedRateLimitingRuleRequestRateLimitRuleConditionHeaderConditions) SetValueList(v []*string) *CreateSharedRateLimitingRuleRequestRateLimitRuleConditionHeaderConditions {
  s.ValueList = v
  return s
}

func (s *CreateSharedRateLimitingRuleRequestRateLimitRuleConditionHeaderConditions) SetKey(v string) *CreateSharedRateLimitingRuleRequestRateLimitRuleConditionHeaderConditions {
  s.Key = &v
  return s
}

type CreateSharedRateLimitingRuleRequestRateLimitRuleConditionAreaConditions struct     {
  // {"en":"Match type.\nEQUAL:Equals\nNOT_EQUAL:Does not equal","zh_CN":"匹配类型。\nEQUAL：等于\nNOT_EQUAL：不等于","exampleValue":"EQUAL,NOT_EQUAL"}
  MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
  // {"dictionary":"belong=WAAP-MS-Ext|dict=waap_areaCityAndCountry","en":"Geo.","zh_CN":"区域。"}
  Areas []*string `json:"areas,omitempty" xml:"areas,omitempty" require:"true" type:"Repeated"`
}

func (s CreateSharedRateLimitingRuleRequestRateLimitRuleConditionAreaConditions) String() string {
  return tea.Prettify(s)
}

func (s CreateSharedRateLimitingRuleRequestRateLimitRuleConditionAreaConditions) GoString() string {
  return s.String()
}

func (s *CreateSharedRateLimitingRuleRequestRateLimitRuleConditionAreaConditions) SetMatchType(v string) *CreateSharedRateLimitingRuleRequestRateLimitRuleConditionAreaConditions {
  s.MatchType = &v
  return s
}

func (s *CreateSharedRateLimitingRuleRequestRateLimitRuleConditionAreaConditions) SetAreas(v []*string) *CreateSharedRateLimitingRuleRequestRateLimitRuleConditionAreaConditions {
  s.Areas = v
  return s
}

type CreateSharedRateLimitingRuleRequestRateLimitRuleConditionSchemeConditions struct     {
  // {"en":"Match type.\nEQUAL:Equals\nNOT_EQUAL:Does not equal","zh_CN":"匹配类型。\nEQUAL：等于\nNOT_EQUAL：不等于","exampleValue":"EQUAL,NOT_EQUAL"}
  MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
  // {"en":"HTTP/S.\nSupported values: HTTP/HTTPS.","zh_CN":"应用层协议。\n支持的值：HTTP/HTTPS。","exampleValue":"HTTP,HTTPS"}
  Scheme []*string `json:"scheme,omitempty" xml:"scheme,omitempty" require:"true" type:"Repeated"`
}

func (s CreateSharedRateLimitingRuleRequestRateLimitRuleConditionSchemeConditions) String() string {
  return tea.Prettify(s)
}

func (s CreateSharedRateLimitingRuleRequestRateLimitRuleConditionSchemeConditions) GoString() string {
  return s.String()
}

func (s *CreateSharedRateLimitingRuleRequestRateLimitRuleConditionSchemeConditions) SetMatchType(v string) *CreateSharedRateLimitingRuleRequestRateLimitRuleConditionSchemeConditions {
  s.MatchType = &v
  return s
}

func (s *CreateSharedRateLimitingRuleRequestRateLimitRuleConditionSchemeConditions) SetScheme(v []*string) *CreateSharedRateLimitingRuleRequestRateLimitRuleConditionSchemeConditions {
  s.Scheme = v
  return s
}

type CreateSharedRateLimitingRuleRequestRateLimitRuleConditionStatusCodeConditions struct     {
  // {"en":"Match type.\nEQUAL:Equals\nNOT_EQUAL:Does not equal","zh_CN":"匹配类型。\nEQUAL：等于\nNOT_EQUAL：不等于","exampleValue":"EQUAL,NOT_EQUAL"}
  MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
  // {"en":"Response Code.","zh_CN":"状态码。"}
  StatusCode []*string `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true" type:"Repeated"`
}

func (s CreateSharedRateLimitingRuleRequestRateLimitRuleConditionStatusCodeConditions) String() string {
  return tea.Prettify(s)
}

func (s CreateSharedRateLimitingRuleRequestRateLimitRuleConditionStatusCodeConditions) GoString() string {
  return s.String()
}

func (s *CreateSharedRateLimitingRuleRequestRateLimitRuleConditionStatusCodeConditions) SetMatchType(v string) *CreateSharedRateLimitingRuleRequestRateLimitRuleConditionStatusCodeConditions {
  s.MatchType = &v
  return s
}

func (s *CreateSharedRateLimitingRuleRequestRateLimitRuleConditionStatusCodeConditions) SetStatusCode(v []*string) *CreateSharedRateLimitingRuleRequestRateLimitRuleConditionStatusCodeConditions {
  s.StatusCode = v
  return s
}

type CreateSharedRateLimitingRuleRequestRateLimitRuleConditionJa3Conditions struct     {
  // {"en":"Match type.\nEQUAL: Equals\nNOT_EQUAL: Does not equal","zh_CN":"匹配类型。\nEQUAL：等于 \nNOT_EQUAL：不等于","exampleValue":"EQUAL,NOT_EQUAL"}
  MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
  // {"en":"JA3 Fingerprint List, maximum 300 JA3 Fingerprint.\nWhen the match type is EQUAL/NOT_EQUAL, each item's character length must be 32 and can only include numbers and lowercase letters.","zh_CN":"JA3指纹列表，最多300个JA3指纹。\n当匹配类型为等于/不等于时，每项字符长度必须为32，且仅限数字和小写字母。"}
  Ja3List []*string `json:"ja3List,omitempty" xml:"ja3List,omitempty" require:"true" type:"Repeated"`
}

func (s CreateSharedRateLimitingRuleRequestRateLimitRuleConditionJa3Conditions) String() string {
  return tea.Prettify(s)
}

func (s CreateSharedRateLimitingRuleRequestRateLimitRuleConditionJa3Conditions) GoString() string {
  return s.String()
}

func (s *CreateSharedRateLimitingRuleRequestRateLimitRuleConditionJa3Conditions) SetMatchType(v string) *CreateSharedRateLimitingRuleRequestRateLimitRuleConditionJa3Conditions {
  s.MatchType = &v
  return s
}

func (s *CreateSharedRateLimitingRuleRequestRateLimitRuleConditionJa3Conditions) SetJa3List(v []*string) *CreateSharedRateLimitingRuleRequestRateLimitRuleConditionJa3Conditions {
  s.Ja3List = v
  return s
}

type CreateSharedRateLimitingRuleRequestRateLimitRuleConditionJa4Conditions struct     {
  // {"en":"Match type. \nEQUAL: Equals\nNOT_EQUAL: Does not equal\nCONTAIN: Contains\nNOT_CONTAIN: Does not Contains\nSTART_WITH: Starts with\nEND_WITH: Ends with\nWILDCARD: Wildcard matches, ** represents zero or more arbitrary characters, ? represents any single character\nNOT_WILDCARD: Wildcard does not match, ** represents zero or more arbitrary characters, ? represents any single character","zh_CN":"匹配类型。\nEQUAL：等于\nNOT_EQUAL：不等于\nCONTAIN：包含\nNOT_CONTAIN：不包含\nSTART_WITH：开头是\nEND_WITH：结尾是\nWILDCARD：通配符匹配，*代表零个或多个任意字符，?代表任意单个字符\nNOT_WILDCARD：通配符不匹配，*代表零个或多个任意字符，?代表任意单个字符","exampleValue":"EQUAL,NOT_EQUAL,CONTAIN,NOT_CONTAIN,START_WITH,END_WITH,WILDCARD,NOT_WILDCARD"}
  MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
  // {"en":"JA4 Fingerprint List, maximum 300 JA4 Fingerprint.\nWhen the match type is EQUAL/NOT_EQUAL, each item's format must be 10 characters + 12 characters + 12 characters, separated by underscores, and can only include underscores, numbers, and lowercase letters.\nWhen the match type is CONTAIN/NOT_CONTAIN/START_WITH/END_WITH, each item is only allowed to include underscores, numbers, and lowercase letters.\nWhen the match type is WILDCARD/NOT_WILDCARD, each item, aside from  ** and ?, is only allowed to include underscores, numbers, and lowercase letters.","zh_CN":"JA4指纹列表，最多300个JA4指纹。\n当匹配类型为等于/不等于时，每项格式必须为10位字符+12位字符+12位字符，中间以“_”分隔，且仅限下划线、数字和小写字母。\n当匹配类型为包含/不包含/开头是/结尾是时，每项只允许输入下划线、数字和小写字母。\n当匹配类型为通配符匹配/通配符不匹配时，每项除*和?外，只允许输入下划线、数字和小写字母。"}
  Ja4List []*string `json:"ja4List,omitempty" xml:"ja4List,omitempty" require:"true" type:"Repeated"`
}

func (s CreateSharedRateLimitingRuleRequestRateLimitRuleConditionJa4Conditions) String() string {
  return tea.Prettify(s)
}

func (s CreateSharedRateLimitingRuleRequestRateLimitRuleConditionJa4Conditions) GoString() string {
  return s.String()
}

func (s *CreateSharedRateLimitingRuleRequestRateLimitRuleConditionJa4Conditions) SetMatchType(v string) *CreateSharedRateLimitingRuleRequestRateLimitRuleConditionJa4Conditions {
  s.MatchType = &v
  return s
}

func (s *CreateSharedRateLimitingRuleRequestRateLimitRuleConditionJa4Conditions) SetJa4List(v []*string) *CreateSharedRateLimitingRuleRequestRateLimitRuleConditionJa4Conditions {
  s.Ja4List = v
  return s
}

type CreateSharedRateLimitingRuleRequestHeader struct {
  // {"en":"Security service type. Please enter a specific service type, if you purchase multiple security services.","zh_CN":"安全服务类型。有使用多个不同的安全服务时，需要填写具体的服务类型。"}
  ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty"`
}

func (s CreateSharedRateLimitingRuleRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s CreateSharedRateLimitingRuleRequestHeader) GoString() string {
  return s.String()
}

func (s *CreateSharedRateLimitingRuleRequestHeader) SetServiceType(v string) *CreateSharedRateLimitingRuleRequestHeader {
  s.ServiceType = &v
  return s
}

type CreateSharedRateLimitingRulePaths struct {
}

func (s CreateSharedRateLimitingRulePaths) String() string {
  return tea.Prettify(s)
}

func (s CreateSharedRateLimitingRulePaths) GoString() string {
  return s.String()
}

type CreateSharedRateLimitingRuleParameters struct {
}

func (s CreateSharedRateLimitingRuleParameters) String() string {
  return tea.Prettify(s)
}

func (s CreateSharedRateLimitingRuleParameters) GoString() string {
  return s.String()
}

type CreateSharedRateLimitingRuleResponse struct {
  // {"dictionary":"belong=WAAP-MS-Ext|dict=waap_retCodeEnum","en":"Please refer to the error code for exceptions.","zh_CN":"请参照错误码。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Description.","zh_CN":"描述信息。"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"Rule ID.","zh_CN":"规则ID。"}
  Data *string `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s CreateSharedRateLimitingRuleResponse) String() string {
  return tea.Prettify(s)
}

func (s CreateSharedRateLimitingRuleResponse) GoString() string {
  return s.String()
}

func (s *CreateSharedRateLimitingRuleResponse) SetCode(v string) *CreateSharedRateLimitingRuleResponse {
  s.Code = &v
  return s
}

func (s *CreateSharedRateLimitingRuleResponse) SetMsg(v string) *CreateSharedRateLimitingRuleResponse {
  s.Msg = &v
  return s
}

func (s *CreateSharedRateLimitingRuleResponse) SetData(v string) *CreateSharedRateLimitingRuleResponse {
  s.Data = &v
  return s
}

type CreateSharedRateLimitingRuleResponseHeader struct {
}

func (s CreateSharedRateLimitingRuleResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s CreateSharedRateLimitingRuleResponseHeader) GoString() string {
  return s.String()
}




type DisassociateShareCustomizeRuleRequest struct {
  // {"en":"Share rule ID.","zh_CN":"共享规则ID。"}
  ShareId *string `json:"shareId,omitempty" xml:"shareId,omitempty" require:"true"`
  // {"en":"List of domain to be disassociated.","zh_CN":"待取消关联域名列表。"}
  DomainList []*string `json:"domainList,omitempty" xml:"domainList,omitempty" require:"true" type:"Repeated"`
}

func (s DisassociateShareCustomizeRuleRequest) String() string {
  return tea.Prettify(s)
}

func (s DisassociateShareCustomizeRuleRequest) GoString() string {
  return s.String()
}

func (s *DisassociateShareCustomizeRuleRequest) SetShareId(v string) *DisassociateShareCustomizeRuleRequest {
  s.ShareId = &v
  return s
}

func (s *DisassociateShareCustomizeRuleRequest) SetDomainList(v []*string) *DisassociateShareCustomizeRuleRequest {
  s.DomainList = v
  return s
}

type DisassociateShareCustomizeRuleRequestHeader struct {
  // {"en":"Security service type. Please enter a specific service type, if you purchase multiple security services.","zh_CN":"安全服务类型。有使用多个不同的安全服务时，需要填写具体的服务类型。"}
  ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty"`
}

func (s DisassociateShareCustomizeRuleRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s DisassociateShareCustomizeRuleRequestHeader) GoString() string {
  return s.String()
}

func (s *DisassociateShareCustomizeRuleRequestHeader) SetServiceType(v string) *DisassociateShareCustomizeRuleRequestHeader {
  s.ServiceType = &v
  return s
}

type DisassociateShareCustomizeRulePaths struct {
}

func (s DisassociateShareCustomizeRulePaths) String() string {
  return tea.Prettify(s)
}

func (s DisassociateShareCustomizeRulePaths) GoString() string {
  return s.String()
}

type DisassociateShareCustomizeRuleParameters struct {
}

func (s DisassociateShareCustomizeRuleParameters) String() string {
  return tea.Prettify(s)
}

func (s DisassociateShareCustomizeRuleParameters) GoString() string {
  return s.String()
}

type DisassociateShareCustomizeRuleResponse struct {
  // {"dictionary":"belong=WAAP-MS-Ext|dict=waap_retCodeEnum","en":"Please refer to the error code for exceptions.","zh_CN":"请参照错误码。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Description.","zh_CN":"描述信息。"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
}

func (s DisassociateShareCustomizeRuleResponse) String() string {
  return tea.Prettify(s)
}

func (s DisassociateShareCustomizeRuleResponse) GoString() string {
  return s.String()
}

func (s *DisassociateShareCustomizeRuleResponse) SetCode(v string) *DisassociateShareCustomizeRuleResponse {
  s.Code = &v
  return s
}

func (s *DisassociateShareCustomizeRuleResponse) SetMsg(v string) *DisassociateShareCustomizeRuleResponse {
  s.Msg = &v
  return s
}

type DisassociateShareCustomizeRuleResponseHeader struct {
}

func (s DisassociateShareCustomizeRuleResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s DisassociateShareCustomizeRuleResponseHeader) GoString() string {
  return s.String()
}




type ListCustomActionsRequest struct {
}

func (s ListCustomActionsRequest) String() string {
  return tea.Prettify(s)
}

func (s ListCustomActionsRequest) GoString() string {
  return s.String()
}

type ListCustomActionsRequestHeader struct {
  // {"en":"Security service type. Please enter a specific service type, if you purchase multiple security services.","zh_CN":"安全服务类型。有使用多个不同的安全服务时，需要填写具体的服务类型。"}
  ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty"`
}

func (s ListCustomActionsRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ListCustomActionsRequestHeader) GoString() string {
  return s.String()
}

func (s *ListCustomActionsRequestHeader) SetServiceType(v string) *ListCustomActionsRequestHeader {
  s.ServiceType = &v
  return s
}

type ListCustomActionsPaths struct {
}

func (s ListCustomActionsPaths) String() string {
  return tea.Prettify(s)
}

func (s ListCustomActionsPaths) GoString() string {
  return s.String()
}

type ListCustomActionsParameters struct {
}

func (s ListCustomActionsParameters) String() string {
  return tea.Prettify(s)
}

func (s ListCustomActionsParameters) GoString() string {
  return s.String()
}

type ListCustomActionsResponse struct {
  // {"dictionary":"belong=WAAP-MS-Ext|dict=waap_retCodeEnum","en":"Please refer to the error code for exceptions.","zh_CN":"请参照错误码。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Description.","zh_CN":"描述信息。"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"Data.","zh_CN":"出参数据。"}
  Data []*ListCustomActionsResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s ListCustomActionsResponse) String() string {
  return tea.Prettify(s)
}

func (s ListCustomActionsResponse) GoString() string {
  return s.String()
}

func (s *ListCustomActionsResponse) SetCode(v string) *ListCustomActionsResponse {
  s.Code = &v
  return s
}

func (s *ListCustomActionsResponse) SetMsg(v string) *ListCustomActionsResponse {
  s.Msg = &v
  return s
}

func (s *ListCustomActionsResponse) SetData(v []*ListCustomActionsResponseData) *ListCustomActionsResponse {
  s.Data = v
  return s
}

type ListCustomActionsResponseData struct     {
  // {"en":"Custom response ID.","zh_CN":"自定义响应ID。"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {"en":"Main account.","zh_CN":"主账号。"}
  MainAccount *string `json:"mainAccount,omitempty" xml:"mainAccount,omitempty" require:"true"`
  // {"en":"Creator.","zh_CN":"创建者。"}
  Creator *string `json:"creator,omitempty" xml:"creator,omitempty" require:"true"`
  // {"en":"Action name,maximum 50 characters.\nDoes not support special characters and spaces.","zh_CN":"动作名称，最多50个字符。\n不支持特殊字符和空格。"}
  ActionName *string `json:"actionName,omitempty" xml:"actionName,omitempty" require:"true"`
  // {"en":"Description,maximum 200 characters.","zh_CN":"描述，最多200个字符。"}
  Description *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
  // {"en":"Status code.\nValid value range:200,204,206,400,401,403,404,500,501,503.","zh_CN":"状态码。\n有效值范围：200，204，206，400，401，403，404，500，501，503。","exampleValue":"200,204,206,400,401,403,404,500,501,503"}
  StatusCode *string `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
  // {"en":"Response Content-Type type,multiple separated by ; sign.","zh_CN":"响应Content-Type类型，多个以 ; 隔开。"}
  ContentType *string `json:"contentType,omitempty" xml:"contentType,omitempty" require:"true"`
  // {"en":"Response content definition, the size of the response content cannot exceed 16KB, if you need to insert static resources such as pictures, it is recommended to use links to import. The following interception information parameters are supported:\n{url}} : Displays the URL information when intercepting\n{client_ip}} : Displays the intercepted user IP information\n{time}} : Displays the intercepted time point\n{event_id} : Displays the ID information of this interception event","zh_CN":"自定义响应内容定义，响应内容大小不能超过16KB，如需插入图片等静态资源，建议采用链接引入。\n支持引用下列拦截信息参数：\n{url} ：显示拦截时的URL信息\n{client_ip} ：显示被拦截的用户IP信息\n{time} ：显示被拦截的时间点\n{event_id} ：显示本次拦截事件的ID信息"}
  CusBlockCnt *string `json:"cusBlockCnt,omitempty" xml:"cusBlockCnt,omitempty" require:"true"`
  // {"en":"List of associated domains.","zh_CN":"关联域名列表。"}
  DomainList *string `json:"domainList,omitempty" xml:"domainList,omitempty" require:"true"`
}

func (s ListCustomActionsResponseData) String() string {
  return tea.Prettify(s)
}

func (s ListCustomActionsResponseData) GoString() string {
  return s.String()
}

func (s *ListCustomActionsResponseData) SetId(v string) *ListCustomActionsResponseData {
  s.Id = &v
  return s
}

func (s *ListCustomActionsResponseData) SetMainAccount(v string) *ListCustomActionsResponseData {
  s.MainAccount = &v
  return s
}

func (s *ListCustomActionsResponseData) SetCreator(v string) *ListCustomActionsResponseData {
  s.Creator = &v
  return s
}

func (s *ListCustomActionsResponseData) SetActionName(v string) *ListCustomActionsResponseData {
  s.ActionName = &v
  return s
}

func (s *ListCustomActionsResponseData) SetDescription(v string) *ListCustomActionsResponseData {
  s.Description = &v
  return s
}

func (s *ListCustomActionsResponseData) SetStatusCode(v string) *ListCustomActionsResponseData {
  s.StatusCode = &v
  return s
}

func (s *ListCustomActionsResponseData) SetContentType(v string) *ListCustomActionsResponseData {
  s.ContentType = &v
  return s
}

func (s *ListCustomActionsResponseData) SetCusBlockCnt(v string) *ListCustomActionsResponseData {
  s.CusBlockCnt = &v
  return s
}

func (s *ListCustomActionsResponseData) SetDomainList(v string) *ListCustomActionsResponseData {
  s.DomainList = &v
  return s
}

type ListCustomActionsResponseHeader struct {
}

func (s ListCustomActionsResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ListCustomActionsResponseHeader) GoString() string {
  return s.String()
}




type DisassociateDmsShareServiceFeatureRequest struct {
  // {"en":"Share configuration ID.","zh_CN":"共享配置ID。"}
  ShareId *string `json:"shareId,omitempty" xml:"shareId,omitempty" require:"true"`
  // {"en":"List of domain to be disassociated.","zh_CN":"待解除关联域名列表。"}
  DisAssDomainList []*string `json:"disAssDomainList,omitempty" xml:"disAssDomainList,omitempty" require:"true" type:"Repeated"`
}

func (s DisassociateDmsShareServiceFeatureRequest) String() string {
  return tea.Prettify(s)
}

func (s DisassociateDmsShareServiceFeatureRequest) GoString() string {
  return s.String()
}

func (s *DisassociateDmsShareServiceFeatureRequest) SetShareId(v string) *DisassociateDmsShareServiceFeatureRequest {
  s.ShareId = &v
  return s
}

func (s *DisassociateDmsShareServiceFeatureRequest) SetDisAssDomainList(v []*string) *DisassociateDmsShareServiceFeatureRequest {
  s.DisAssDomainList = v
  return s
}

type DisassociateDmsShareServiceFeatureRequestHeader struct {
  // {"en":"Security service type. Please enter a specific service type, if you purchase multiple security services.","zh_CN":"安全服务类型。有使用多个不同的安全服务时，需要填写具体的服务类型。"}
  ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty"`
}

func (s DisassociateDmsShareServiceFeatureRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s DisassociateDmsShareServiceFeatureRequestHeader) GoString() string {
  return s.String()
}

func (s *DisassociateDmsShareServiceFeatureRequestHeader) SetServiceType(v string) *DisassociateDmsShareServiceFeatureRequestHeader {
  s.ServiceType = &v
  return s
}

type DisassociateDmsShareServiceFeaturePaths struct {
}

func (s DisassociateDmsShareServiceFeaturePaths) String() string {
  return tea.Prettify(s)
}

func (s DisassociateDmsShareServiceFeaturePaths) GoString() string {
  return s.String()
}

type DisassociateDmsShareServiceFeatureParameters struct {
}

func (s DisassociateDmsShareServiceFeatureParameters) String() string {
  return tea.Prettify(s)
}

func (s DisassociateDmsShareServiceFeatureParameters) GoString() string {
  return s.String()
}

type DisassociateDmsShareServiceFeatureResponse struct {
  // {"dictionary":"belong=WAAP-MS-Ext|dict=waap_retCodeEnum","en":"Please refer to the error code for exceptions.","zh_CN":"请参照错误码。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Description.","zh_CN":"描述信息。"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
}

func (s DisassociateDmsShareServiceFeatureResponse) String() string {
  return tea.Prettify(s)
}

func (s DisassociateDmsShareServiceFeatureResponse) GoString() string {
  return s.String()
}

func (s *DisassociateDmsShareServiceFeatureResponse) SetCode(v string) *DisassociateDmsShareServiceFeatureResponse {
  s.Code = &v
  return s
}

func (s *DisassociateDmsShareServiceFeatureResponse) SetMsg(v string) *DisassociateDmsShareServiceFeatureResponse {
  s.Msg = &v
  return s
}

type DisassociateDmsShareServiceFeatureResponseHeader struct {
}

func (s DisassociateDmsShareServiceFeatureResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s DisassociateDmsShareServiceFeatureResponseHeader) GoString() string {
  return s.String()
}




type AssociateShareRateLimitRequest struct {
  // {"en":"Share rule ID.","zh_CN":"共享规则ID。"}
  ShareId *string `json:"shareId,omitempty" xml:"shareId,omitempty" require:"true"`
  // {"en":"List of domain to be associated.","zh_CN":"待关联域名列表。"}
  DomainList []*string `json:"domainList,omitempty" xml:"domainList,omitempty" require:"true" type:"Repeated"`
}

func (s AssociateShareRateLimitRequest) String() string {
  return tea.Prettify(s)
}

func (s AssociateShareRateLimitRequest) GoString() string {
  return s.String()
}

func (s *AssociateShareRateLimitRequest) SetShareId(v string) *AssociateShareRateLimitRequest {
  s.ShareId = &v
  return s
}

func (s *AssociateShareRateLimitRequest) SetDomainList(v []*string) *AssociateShareRateLimitRequest {
  s.DomainList = v
  return s
}

type AssociateShareRateLimitRequestHeader struct {
  // {"en":"Security service type. Please enter a specific service type, if you purchase multiple security services.","zh_CN":"安全服务类型。有使用多个不同的安全服务时，需要填写具体的服务类型。"}
  ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty"`
}

func (s AssociateShareRateLimitRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s AssociateShareRateLimitRequestHeader) GoString() string {
  return s.String()
}

func (s *AssociateShareRateLimitRequestHeader) SetServiceType(v string) *AssociateShareRateLimitRequestHeader {
  s.ServiceType = &v
  return s
}

type AssociateShareRateLimitPaths struct {
}

func (s AssociateShareRateLimitPaths) String() string {
  return tea.Prettify(s)
}

func (s AssociateShareRateLimitPaths) GoString() string {
  return s.String()
}

type AssociateShareRateLimitParameters struct {
}

func (s AssociateShareRateLimitParameters) String() string {
  return tea.Prettify(s)
}

func (s AssociateShareRateLimitParameters) GoString() string {
  return s.String()
}

type AssociateShareRateLimitResponse struct {
  // {"dictionary":"belong=WAAP-MS-Ext|dict=waap_retCodeEnum","en":"Please refer to the error code for exceptions.","zh_CN":"请参照错误码。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Description.","zh_CN":"描述信息。"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
}

func (s AssociateShareRateLimitResponse) String() string {
  return tea.Prettify(s)
}

func (s AssociateShareRateLimitResponse) GoString() string {
  return s.String()
}

func (s *AssociateShareRateLimitResponse) SetCode(v string) *AssociateShareRateLimitResponse {
  s.Code = &v
  return s
}

func (s *AssociateShareRateLimitResponse) SetMsg(v string) *AssociateShareRateLimitResponse {
  s.Msg = &v
  return s
}

type AssociateShareRateLimitResponseHeader struct {
}

func (s AssociateShareRateLimitResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s AssociateShareRateLimitResponseHeader) GoString() string {
  return s.String()
}




type AssociateDmsShareServiceFeatureRequest struct {
  // {"en":"Share configuration ID.","zh_CN":"共享配置ID。"}
  ShareId *string `json:"shareId,omitempty" xml:"shareId,omitempty" require:"true"`
  // {"en":"List of domain to be associated.","zh_CN":"待关联域名列表。"}
  AssDomainList []*string `json:"assDomainList,omitempty" xml:"assDomainList,omitempty" require:"true" type:"Repeated"`
}

func (s AssociateDmsShareServiceFeatureRequest) String() string {
  return tea.Prettify(s)
}

func (s AssociateDmsShareServiceFeatureRequest) GoString() string {
  return s.String()
}

func (s *AssociateDmsShareServiceFeatureRequest) SetShareId(v string) *AssociateDmsShareServiceFeatureRequest {
  s.ShareId = &v
  return s
}

func (s *AssociateDmsShareServiceFeatureRequest) SetAssDomainList(v []*string) *AssociateDmsShareServiceFeatureRequest {
  s.AssDomainList = v
  return s
}

type AssociateDmsShareServiceFeatureRequestHeader struct {
  // {"en":"Security service type. Please enter a specific service type, if you purchase multiple security services.","zh_CN":"安全服务类型。有使用多个不同的安全服务时，需要填写具体的服务类型。"}
  ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty"`
}

func (s AssociateDmsShareServiceFeatureRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s AssociateDmsShareServiceFeatureRequestHeader) GoString() string {
  return s.String()
}

func (s *AssociateDmsShareServiceFeatureRequestHeader) SetServiceType(v string) *AssociateDmsShareServiceFeatureRequestHeader {
  s.ServiceType = &v
  return s
}

type AssociateDmsShareServiceFeaturePaths struct {
}

func (s AssociateDmsShareServiceFeaturePaths) String() string {
  return tea.Prettify(s)
}

func (s AssociateDmsShareServiceFeaturePaths) GoString() string {
  return s.String()
}

type AssociateDmsShareServiceFeatureParameters struct {
}

func (s AssociateDmsShareServiceFeatureParameters) String() string {
  return tea.Prettify(s)
}

func (s AssociateDmsShareServiceFeatureParameters) GoString() string {
  return s.String()
}

type AssociateDmsShareServiceFeatureResponse struct {
  // {"dictionary":"belong=WAAP-MS-Ext|dict=waap_retCodeEnum","en":"Please refer to the error code for exceptions.","zh_CN":"请参照错误码。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Description.","zh_CN":"描述信息。"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
}

func (s AssociateDmsShareServiceFeatureResponse) String() string {
  return tea.Prettify(s)
}

func (s AssociateDmsShareServiceFeatureResponse) GoString() string {
  return s.String()
}

func (s *AssociateDmsShareServiceFeatureResponse) SetCode(v string) *AssociateDmsShareServiceFeatureResponse {
  s.Code = &v
  return s
}

func (s *AssociateDmsShareServiceFeatureResponse) SetMsg(v string) *AssociateDmsShareServiceFeatureResponse {
  s.Msg = &v
  return s
}

type AssociateDmsShareServiceFeatureResponseHeader struct {
}

func (s AssociateDmsShareServiceFeatureResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s AssociateDmsShareServiceFeatureResponseHeader) GoString() string {
  return s.String()
}




type AssociateSharedWhitelistRuleRequest struct {
  // {"en":"Share rule ID.","zh_CN":"共享规则ID。"}
  ShareId *string `json:"shareId,omitempty" xml:"shareId,omitempty" require:"true"`
  // {"en":"List of domain to be associated.","zh_CN":"待关联域名列表。"}
  DomainList []*string `json:"domainList,omitempty" xml:"domainList,omitempty" require:"true" type:"Repeated"`
}

func (s AssociateSharedWhitelistRuleRequest) String() string {
  return tea.Prettify(s)
}

func (s AssociateSharedWhitelistRuleRequest) GoString() string {
  return s.String()
}

func (s *AssociateSharedWhitelistRuleRequest) SetShareId(v string) *AssociateSharedWhitelistRuleRequest {
  s.ShareId = &v
  return s
}

func (s *AssociateSharedWhitelistRuleRequest) SetDomainList(v []*string) *AssociateSharedWhitelistRuleRequest {
  s.DomainList = v
  return s
}

type AssociateSharedWhitelistRuleRequestHeader struct {
  // {"en":"Security service type. Please enter a specific service type, if you purchase multiple security services.","zh_CN":"安全服务类型。有使用多个不同的安全服务时，需要填写具体的服务类型。"}
  ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty"`
}

func (s AssociateSharedWhitelistRuleRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s AssociateSharedWhitelistRuleRequestHeader) GoString() string {
  return s.String()
}

func (s *AssociateSharedWhitelistRuleRequestHeader) SetServiceType(v string) *AssociateSharedWhitelistRuleRequestHeader {
  s.ServiceType = &v
  return s
}

type AssociateSharedWhitelistRulePaths struct {
}

func (s AssociateSharedWhitelistRulePaths) String() string {
  return tea.Prettify(s)
}

func (s AssociateSharedWhitelistRulePaths) GoString() string {
  return s.String()
}

type AssociateSharedWhitelistRuleParameters struct {
}

func (s AssociateSharedWhitelistRuleParameters) String() string {
  return tea.Prettify(s)
}

func (s AssociateSharedWhitelistRuleParameters) GoString() string {
  return s.String()
}

type AssociateSharedWhitelistRuleResponse struct {
  // {"dictionary":"belong=WAAP-MS-Ext|dict=waap_retCodeEnum","en":"Please refer to the error code for exceptions.","zh_CN":"请参照错误码。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Description.","zh_CN":"描述信息。"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
}

func (s AssociateSharedWhitelistRuleResponse) String() string {
  return tea.Prettify(s)
}

func (s AssociateSharedWhitelistRuleResponse) GoString() string {
  return s.String()
}

func (s *AssociateSharedWhitelistRuleResponse) SetCode(v string) *AssociateSharedWhitelistRuleResponse {
  s.Code = &v
  return s
}

func (s *AssociateSharedWhitelistRuleResponse) SetMsg(v string) *AssociateSharedWhitelistRuleResponse {
  s.Msg = &v
  return s
}

type AssociateSharedWhitelistRuleResponseHeader struct {
}

func (s AssociateSharedWhitelistRuleResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s AssociateSharedWhitelistRuleResponseHeader) GoString() string {
  return s.String()
}




type CreateSharedWAFRuleExceptionRequest struct {
  // {"en":"Exception name,maximum 50 character.\nDoes not support # and &.","zh_CN":"例外名称，最多50个字符。\n不支持 # 和 &。"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"Exception description,maximum 200 characters.","zh_CN":"例外描述，最多200个字符。"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty"`
  // {"en":"Matching conditions.\nip: IP\npath: Path\nuri: URI\nurlParamName: URI Parameter Name\nurlParamValue: URI Parameter Value\nuserAgent: User Agent\nhttpHeaderName: Request Header Name\nhttpHeaderValue: Request Header Value\ncookie: Cookie\nbody: Body\nbodyParamName: Body Parameter Name\nbodyParamValue: Body Parameter Value","zh_CN":"匹配条件。\nip：IP\npath：路径\nuri：URI\nurlParamName：URI参数名\nurlParamValue：URI参数值\nuserAgent：User Agent\nhttpHeaderName：请求头部名称\nhttpHeaderValue：请求头部值\ncookie：Cookie\nbody：Body\nbodyParamName：Body参数名\nbodyParamValue：Body参数值","exampleValue":"ip,path,uri,urlParamName,urlParamValue,userAgent,httpHeaderName,httpHeaderValue,cookie,body,bodyParamName,bodyParamValue"}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
  // {"en":"Match type,IP can only be EQUAL.\nEQUAL: Equal\nCONTAIN: Contains\nREGEX: Regular match","zh_CN":"匹配类型，IP只能是等于。\nEQUAL：等于\nCONTAIN：包含\nREGEX：正则匹配","exampleValue":"EQUAL,CONTAIN,REGEX"}
  MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
  // {"en":"Rule exceptions.\nWhen matchType=EQUAL, case-sensitive, path and uri must start with \"/\", and body can only pass one value;\nWhen matchType=REGEX, only one value can be passed","zh_CN":"规则例外内容。\nmatchType=EQUAL时，大小写敏感，path和uri必须以\"/\"开头，body只能传一个值；\nmatchType=REGEX时，只能传一个值。"}
  ContentList []*string `json:"contentList,omitempty" xml:"contentList,omitempty" require:"true" type:"Repeated"`
}

func (s CreateSharedWAFRuleExceptionRequest) String() string {
  return tea.Prettify(s)
}

func (s CreateSharedWAFRuleExceptionRequest) GoString() string {
  return s.String()
}

func (s *CreateSharedWAFRuleExceptionRequest) SetName(v string) *CreateSharedWAFRuleExceptionRequest {
  s.Name = &v
  return s
}

func (s *CreateSharedWAFRuleExceptionRequest) SetMsg(v string) *CreateSharedWAFRuleExceptionRequest {
  s.Msg = &v
  return s
}

func (s *CreateSharedWAFRuleExceptionRequest) SetType(v string) *CreateSharedWAFRuleExceptionRequest {
  s.Type = &v
  return s
}

func (s *CreateSharedWAFRuleExceptionRequest) SetMatchType(v string) *CreateSharedWAFRuleExceptionRequest {
  s.MatchType = &v
  return s
}

func (s *CreateSharedWAFRuleExceptionRequest) SetContentList(v []*string) *CreateSharedWAFRuleExceptionRequest {
  s.ContentList = v
  return s
}

type CreateSharedWAFRuleExceptionRequestHeader struct {
  // {"dictionary":"belong=WAAP-MS-Ext|dict=waap_serviceType","en":"Security service type. Please enter a specific service type, if you purchase multiple security services.","zh_CN":"安全服务类型。有使用多个不同的安全服务时，需要填写具体的服务类型。"}
  ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty"`
}

func (s CreateSharedWAFRuleExceptionRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s CreateSharedWAFRuleExceptionRequestHeader) GoString() string {
  return s.String()
}

func (s *CreateSharedWAFRuleExceptionRequestHeader) SetServiceType(v string) *CreateSharedWAFRuleExceptionRequestHeader {
  s.ServiceType = &v
  return s
}

type CreateSharedWAFRuleExceptionPaths struct {
}

func (s CreateSharedWAFRuleExceptionPaths) String() string {
  return tea.Prettify(s)
}

func (s CreateSharedWAFRuleExceptionPaths) GoString() string {
  return s.String()
}

type CreateSharedWAFRuleExceptionParameters struct {
}

func (s CreateSharedWAFRuleExceptionParameters) String() string {
  return tea.Prettify(s)
}

func (s CreateSharedWAFRuleExceptionParameters) GoString() string {
  return s.String()
}

type CreateSharedWAFRuleExceptionResponse struct {
  // {"dictionary":"belong=WAAP-MS-Ext|dict=waap_retCodeEnum","en":"Please refer to the error code for exceptions.","zh_CN":"请参照错误码。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Description.","zh_CN":"描述信息。"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"Rule exception ID.","zh_CN":"规则例外ID。"}
  Data *string `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s CreateSharedWAFRuleExceptionResponse) String() string {
  return tea.Prettify(s)
}

func (s CreateSharedWAFRuleExceptionResponse) GoString() string {
  return s.String()
}

func (s *CreateSharedWAFRuleExceptionResponse) SetCode(v string) *CreateSharedWAFRuleExceptionResponse {
  s.Code = &v
  return s
}

func (s *CreateSharedWAFRuleExceptionResponse) SetMsg(v string) *CreateSharedWAFRuleExceptionResponse {
  s.Msg = &v
  return s
}

func (s *CreateSharedWAFRuleExceptionResponse) SetData(v string) *CreateSharedWAFRuleExceptionResponse {
  s.Data = &v
  return s
}

type CreateSharedWAFRuleExceptionResponseHeader struct {
}

func (s CreateSharedWAFRuleExceptionResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s CreateSharedWAFRuleExceptionResponseHeader) GoString() string {
  return s.String()
}




type ListSharedCustomRulesRequest struct {
  // {"en":"Rule name, fuzzy query.","zh_CN":"规则名称，模糊查询。"}
  RuleName *string `json:"ruleName,omitempty" xml:"ruleName,omitempty"`
}

func (s ListSharedCustomRulesRequest) String() string {
  return tea.Prettify(s)
}

func (s ListSharedCustomRulesRequest) GoString() string {
  return s.String()
}

func (s *ListSharedCustomRulesRequest) SetRuleName(v string) *ListSharedCustomRulesRequest {
  s.RuleName = &v
  return s
}

type ListSharedCustomRulesRequestHeader struct {
  // {"en":"Security service type. Please enter a specific service type, if you purchase multiple security services.","zh_CN":"安全服务类型。有使用多个不同的安全服务时，需要填写具体的服务类型。"}
  ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty"`
}

func (s ListSharedCustomRulesRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ListSharedCustomRulesRequestHeader) GoString() string {
  return s.String()
}

func (s *ListSharedCustomRulesRequestHeader) SetServiceType(v string) *ListSharedCustomRulesRequestHeader {
  s.ServiceType = &v
  return s
}

type ListSharedCustomRulesPaths struct {
}

func (s ListSharedCustomRulesPaths) String() string {
  return tea.Prettify(s)
}

func (s ListSharedCustomRulesPaths) GoString() string {
  return s.String()
}

type ListSharedCustomRulesParameters struct {
}

func (s ListSharedCustomRulesParameters) String() string {
  return tea.Prettify(s)
}

func (s ListSharedCustomRulesParameters) GoString() string {
  return s.String()
}

type ListSharedCustomRulesResponse struct {
  // {"en":"Description.","zh_CN":"描述信息。"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"dictionary":"belong=WAAP-MS-Ext|dict=waap_retCodeEnum","en":"Please refer to the error code for exceptions.","zh_CN":"请参照错误码。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Data.","zh_CN":"出参数据。"}
  Data []*ListSharedCustomRulesResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s ListSharedCustomRulesResponse) String() string {
  return tea.Prettify(s)
}

func (s ListSharedCustomRulesResponse) GoString() string {
  return s.String()
}

func (s *ListSharedCustomRulesResponse) SetMsg(v string) *ListSharedCustomRulesResponse {
  s.Msg = &v
  return s
}

func (s *ListSharedCustomRulesResponse) SetCode(v string) *ListSharedCustomRulesResponse {
  s.Code = &v
  return s
}

func (s *ListSharedCustomRulesResponse) SetData(v []*ListSharedCustomRulesResponseData) *ListSharedCustomRulesResponse {
  s.Data = v
  return s
}

type ListSharedCustomRulesResponseData struct     {
  // {"en":"Matching conditions.","zh_CN":"匹配条件。"}
  Condition *ListSharedCustomRulesResponseDataCondition `json:"condition,omitempty" xml:"condition,omitempty" require:"true" type:"Struct"`
  // {"en":"Creator.","zh_CN":"创建者。"}
  Creator *string `json:"creator,omitempty" xml:"creator,omitempty" require:"true"`
  // {"en":"Action.\nNO_USE: Not Used\nLOG: Log\nDELAY: Delay\nBLOCK: Deny\nRESET: Reset Connection\nCUSTOM_ACTION_ID:Fill in the custom action id of the corresponding action","zh_CN":"处理动作。\nNO_USE：不使用\nLOG：监控\nDELAY：延迟响应\nBLOCK：拦截\nRESET：断开连接\nCUSTOM_ACTION_ID：传入对应的自定义响应动作的ID","exampleValue":"NO_USE,LOG,DELAY,BLOCK,RESET"}
  Act *string `json:"act,omitempty" xml:"act,omitempty" require:"true"`
  // {"en":"Create time, format: yyyy-MM-dd HH:mm:ss.","zh_CN":"创建时间，格式：yyyy-MM-dd HH:mm:ss。"}
  CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty" require:"true"`
  // {"en":"Number of associated hostnames.","zh_CN":"关联域名数。"}
  RelationDomainCount *int `json:"relationDomainCount,omitempty" xml:"relationDomainCount,omitempty" require:"true"`
  // {"en":"Rule Name.","zh_CN":"规则名称。"}
  RuleName *string `json:"ruleName,omitempty" xml:"ruleName,omitempty" require:"true"`
  // {"en":"Description.","zh_CN":"规则描述。"}
  Description *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
  // {"en":"Update time, format: yyyy-MM-dd HH:mm:ss.","zh_CN":"更新时间，格式：yyyy-MM-dd HH:mm:ss。"}
  UpdateTime *string `json:"updateTime,omitempty" xml:"updateTime,omitempty" require:"true"`
  // {"en":"Rule ID.","zh_CN":"规则ID。"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
}

func (s ListSharedCustomRulesResponseData) String() string {
  return tea.Prettify(s)
}

func (s ListSharedCustomRulesResponseData) GoString() string {
  return s.String()
}

func (s *ListSharedCustomRulesResponseData) SetCondition(v *ListSharedCustomRulesResponseDataCondition) *ListSharedCustomRulesResponseData {
  s.Condition = v
  return s
}

func (s *ListSharedCustomRulesResponseData) SetCreator(v string) *ListSharedCustomRulesResponseData {
  s.Creator = &v
  return s
}

func (s *ListSharedCustomRulesResponseData) SetAct(v string) *ListSharedCustomRulesResponseData {
  s.Act = &v
  return s
}

func (s *ListSharedCustomRulesResponseData) SetCreateTime(v string) *ListSharedCustomRulesResponseData {
  s.CreateTime = &v
  return s
}

func (s *ListSharedCustomRulesResponseData) SetRelationDomainCount(v int) *ListSharedCustomRulesResponseData {
  s.RelationDomainCount = &v
  return s
}

func (s *ListSharedCustomRulesResponseData) SetRuleName(v string) *ListSharedCustomRulesResponseData {
  s.RuleName = &v
  return s
}

func (s *ListSharedCustomRulesResponseData) SetDescription(v string) *ListSharedCustomRulesResponseData {
  s.Description = &v
  return s
}

func (s *ListSharedCustomRulesResponseData) SetUpdateTime(v string) *ListSharedCustomRulesResponseData {
  s.UpdateTime = &v
  return s
}

func (s *ListSharedCustomRulesResponseData) SetId(v string) *ListSharedCustomRulesResponseData {
  s.Id = &v
  return s
}

type ListSharedCustomRulesResponseDataCondition struct {
  // {"en":"Request Method.","zh_CN":"请求方法。"}
  MethodConditions []*ListSharedCustomRulesResponseDataConditionMethodConditions `json:"methodConditions,omitempty" xml:"methodConditions,omitempty" require:"true" type:"Repeated"`
  // {"en":"Geo.","zh_CN":"区域。"}
  AreaConditions []*ListSharedCustomRulesResponseDataConditionAreaConditions `json:"areaConditions,omitempty" xml:"areaConditions,omitempty" require:"true" type:"Repeated"`
  // {"en":"IP/CIDR.","zh_CN":"IP/IP段。"}
  IpOrIpsConditions []*ListSharedCustomRulesResponseDataConditionIpOrIpsConditions `json:"ipOrIpsConditions,omitempty" xml:"ipOrIpsConditions,omitempty" require:"true" type:"Repeated"`
  // {"en":"URI.","zh_CN":"URI。"}
  UriConditions []*ListSharedCustomRulesResponseDataConditionUriConditions `json:"uriConditions,omitempty" xml:"uriConditions,omitempty" require:"true" type:"Repeated"`
  // {"en":"Path.","zh_CN":"路径。"}
  PathConditions []*ListSharedCustomRulesResponseDataConditionPathConditions `json:"pathConditions,omitempty" xml:"pathConditions,omitempty" require:"true" type:"Repeated"`
  // {"en":"User Agent.","zh_CN":"User-Agent。"}
  UaConditions []*ListSharedCustomRulesResponseDataConditionUaConditions `json:"uaConditions,omitempty" xml:"uaConditions,omitempty" require:"true" type:"Repeated"`
  // {"en":"Request Header.","zh_CN":"请求头。"}
  HeaderConditions []*ListSharedCustomRulesResponseDataConditionHeaderConditions `json:"headerConditions,omitempty" xml:"headerConditions,omitempty" require:"true" type:"Repeated"`
  // {"en":"Referer.","zh_CN":"Referer。"}
  RefererConditions []*ListSharedCustomRulesResponseDataConditionRefererConditions `json:"refererConditions,omitempty" xml:"refererConditions,omitempty" require:"true" type:"Repeated"`
  // {"en":"JA3 Fingerprint, match type cannot be repeated.","zh_CN":"JA3指纹，匹配类型不可重复。"}
  Ja3Conditions []*ListSharedCustomRulesResponseDataConditionJa3Conditions `json:"ja3Conditions,omitempty" xml:"ja3Conditions,omitempty" require:"true" type:"Repeated"`
  // {"en":"JA4 Fingerprint, match type cannot be repeated.","zh_CN":"JA4指纹，匹配类型不可重复。"}
  Ja4Conditions []*ListSharedCustomRulesResponseDataConditionJa4Conditions `json:"ja4Conditions,omitempty" xml:"ja4Conditions,omitempty" require:"true" type:"Repeated"`
}

func (s ListSharedCustomRulesResponseDataCondition) String() string {
  return tea.Prettify(s)
}

func (s ListSharedCustomRulesResponseDataCondition) GoString() string {
  return s.String()
}

func (s *ListSharedCustomRulesResponseDataCondition) SetMethodConditions(v []*ListSharedCustomRulesResponseDataConditionMethodConditions) *ListSharedCustomRulesResponseDataCondition {
  s.MethodConditions = v
  return s
}

func (s *ListSharedCustomRulesResponseDataCondition) SetAreaConditions(v []*ListSharedCustomRulesResponseDataConditionAreaConditions) *ListSharedCustomRulesResponseDataCondition {
  s.AreaConditions = v
  return s
}

func (s *ListSharedCustomRulesResponseDataCondition) SetIpOrIpsConditions(v []*ListSharedCustomRulesResponseDataConditionIpOrIpsConditions) *ListSharedCustomRulesResponseDataCondition {
  s.IpOrIpsConditions = v
  return s
}

func (s *ListSharedCustomRulesResponseDataCondition) SetUriConditions(v []*ListSharedCustomRulesResponseDataConditionUriConditions) *ListSharedCustomRulesResponseDataCondition {
  s.UriConditions = v
  return s
}

func (s *ListSharedCustomRulesResponseDataCondition) SetPathConditions(v []*ListSharedCustomRulesResponseDataConditionPathConditions) *ListSharedCustomRulesResponseDataCondition {
  s.PathConditions = v
  return s
}

func (s *ListSharedCustomRulesResponseDataCondition) SetUaConditions(v []*ListSharedCustomRulesResponseDataConditionUaConditions) *ListSharedCustomRulesResponseDataCondition {
  s.UaConditions = v
  return s
}

func (s *ListSharedCustomRulesResponseDataCondition) SetHeaderConditions(v []*ListSharedCustomRulesResponseDataConditionHeaderConditions) *ListSharedCustomRulesResponseDataCondition {
  s.HeaderConditions = v
  return s
}

func (s *ListSharedCustomRulesResponseDataCondition) SetRefererConditions(v []*ListSharedCustomRulesResponseDataConditionRefererConditions) *ListSharedCustomRulesResponseDataCondition {
  s.RefererConditions = v
  return s
}

func (s *ListSharedCustomRulesResponseDataCondition) SetJa3Conditions(v []*ListSharedCustomRulesResponseDataConditionJa3Conditions) *ListSharedCustomRulesResponseDataCondition {
  s.Ja3Conditions = v
  return s
}

func (s *ListSharedCustomRulesResponseDataCondition) SetJa4Conditions(v []*ListSharedCustomRulesResponseDataConditionJa4Conditions) *ListSharedCustomRulesResponseDataCondition {
  s.Ja4Conditions = v
  return s
}

type ListSharedCustomRulesResponseDataConditionMethodConditions struct     {
  // {"en":"Match type.\nEQUAL:Equal\nNOT_EQUAL:Does not equal","zh_CN":"匹配类型。\nEQUAL：等于\nNOT_EQUAL：不等于","exampleValue":"EQUAL,NOT_EQUAL"}
  MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
  // {"en":"Request method.\nSupported values: GET/POST/DELETE/PUT/HEAD/OPTIONS/COPY.","zh_CN":"请求方法。\n支持的值：GET/POST/DELETE/PUT/HEAD/OPTIONS/COPY","exampleValue":"GET,POST,DELETE,PUT,HEAD,OPTIONS,COPY"}
  RequestMethod []*string `json:"requestMethod,omitempty" xml:"requestMethod,omitempty" require:"true" type:"Repeated"`
}

func (s ListSharedCustomRulesResponseDataConditionMethodConditions) String() string {
  return tea.Prettify(s)
}

func (s ListSharedCustomRulesResponseDataConditionMethodConditions) GoString() string {
  return s.String()
}

func (s *ListSharedCustomRulesResponseDataConditionMethodConditions) SetMatchType(v string) *ListSharedCustomRulesResponseDataConditionMethodConditions {
  s.MatchType = &v
  return s
}

func (s *ListSharedCustomRulesResponseDataConditionMethodConditions) SetRequestMethod(v []*string) *ListSharedCustomRulesResponseDataConditionMethodConditions {
  s.RequestMethod = v
  return s
}

type ListSharedCustomRulesResponseDataConditionAreaConditions struct     {
  // {"en":"Match type.\nEQUAL:Equal\nNOT_EQUAL:Does not equal","zh_CN":"匹配类型。\nEQUAL：等于\nNOT_EQUAL：不等于","exampleValue":"EQUAL,NOT_EQUAL"}
  MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
  // {"dictionary":"belong=WAAP-MS-Ext|dict=waap_areaCityAndCountry","en":"Geo.","zh_CN":"区域。"}
  Areas []*string `json:"areas,omitempty" xml:"areas,omitempty" require:"true" type:"Repeated"`
}

func (s ListSharedCustomRulesResponseDataConditionAreaConditions) String() string {
  return tea.Prettify(s)
}

func (s ListSharedCustomRulesResponseDataConditionAreaConditions) GoString() string {
  return s.String()
}

func (s *ListSharedCustomRulesResponseDataConditionAreaConditions) SetMatchType(v string) *ListSharedCustomRulesResponseDataConditionAreaConditions {
  s.MatchType = &v
  return s
}

func (s *ListSharedCustomRulesResponseDataConditionAreaConditions) SetAreas(v []*string) *ListSharedCustomRulesResponseDataConditionAreaConditions {
  s.Areas = v
  return s
}

type ListSharedCustomRulesResponseDataConditionIpOrIpsConditions struct     {
  // {"en":"Match type.\nEQUAL:Equal\nNOT_EQUAL:Does not equal","zh_CN":"匹配类型。\nEQUAL：等于\nNOT_EQUAL：不等于","exampleValue":"EQUAL,NOT_EQUAL"}
  MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
  // {"en":"IP/CIDR.","zh_CN":"IP/IP段。"}
  IpOrIps []*string `json:"ipOrIps,omitempty" xml:"ipOrIps,omitempty" require:"true" type:"Repeated"`
}

func (s ListSharedCustomRulesResponseDataConditionIpOrIpsConditions) String() string {
  return tea.Prettify(s)
}

func (s ListSharedCustomRulesResponseDataConditionIpOrIpsConditions) GoString() string {
  return s.String()
}

func (s *ListSharedCustomRulesResponseDataConditionIpOrIpsConditions) SetMatchType(v string) *ListSharedCustomRulesResponseDataConditionIpOrIpsConditions {
  s.MatchType = &v
  return s
}

func (s *ListSharedCustomRulesResponseDataConditionIpOrIpsConditions) SetIpOrIps(v []*string) *ListSharedCustomRulesResponseDataConditionIpOrIpsConditions {
  s.IpOrIps = v
  return s
}

type ListSharedCustomRulesResponseDataConditionUriConditions struct     {
  // {"en":"Match type.\nEQUAL: equal to\nNOT_EQUAL: not equal to\nCONTAIN: contains\nNOT_CONTAIN: does not contain\nREGEX: regular\nNOT_REGEX: regular does not match\nSTART_WITH: starts with\nEND_WITH: ends with\nWILDCARD: wildcard matches\nNOT_WILDCARD: wildcard does not match","zh_CN":"匹配类型。\nEQUAL：等于\nNOT_EQUAL：不等于\nCONTAIN：包含\nNOT_CONTAIN：不包含\nREGEX：正则\nNOT_REGEX：正则不匹配\nSTART_WITH：开头是\nEND_WITH：结尾是\nWILDCARD：通配符匹配\nNOT_WILDCARD：通配符不匹配","exampleValue":"EQUAL,NOT_EQUAL,CONTAIN,NOT_CONTAIN,REGEX,NOT_REGEX,START_WITH,END_WITH,WILDCARD,NOT_WILDCARD"}
  MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
  // {"en":"URI.","zh_CN":"URI。"}
  Uri []*string `json:"uri,omitempty" xml:"uri,omitempty" require:"true" type:"Repeated"`
}

func (s ListSharedCustomRulesResponseDataConditionUriConditions) String() string {
  return tea.Prettify(s)
}

func (s ListSharedCustomRulesResponseDataConditionUriConditions) GoString() string {
  return s.String()
}

func (s *ListSharedCustomRulesResponseDataConditionUriConditions) SetMatchType(v string) *ListSharedCustomRulesResponseDataConditionUriConditions {
  s.MatchType = &v
  return s
}

func (s *ListSharedCustomRulesResponseDataConditionUriConditions) SetUri(v []*string) *ListSharedCustomRulesResponseDataConditionUriConditions {
  s.Uri = v
  return s
}

type ListSharedCustomRulesResponseDataConditionPathConditions struct     {
  // {"en":"Match type.\nEQUAL: equal to\nNOT_EQUAL: not equal to\nCONTAIN: contains\nNOT_CONTAIN: does not contain\nREGEX: regular\nNOT_REGEX: regular does not match\nSTART_WITH: starts with\nEND_WITH: ends with\nWILDCARD: wildcard matches\nNOT_WILDCARD: wildcard does not match","zh_CN":"匹配类型。\nEQUAL：等于\nNOT_EQUAL：不等于\nCONTAIN：包含\nNOT_CONTAIN：不包含\nREGEX：正则\nNOT_REGEX：正则不匹配\nSTART_WITH：开头是\nEND_WITH：结尾是\nWILDCARD：通配符匹配\nNOT_WILDCARD：通配符不匹配","exampleValue":"EQUAL,NOT_EQUAL,CONTAIN,NOT_CONTAIN,REGEX,NOT_REGEX,START_WITH,END_WITH,WILDCARD,NOT_WILDCARD"}
  MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
  // {"en":"Path.","zh_CN":"路径。"}
  ListSharedCustomRulesPaths []*string `json:"paths,omitempty" xml:"paths,omitempty" require:"true" type:"Repeated"`
}

func (s ListSharedCustomRulesResponseDataConditionPathConditions) String() string {
  return tea.Prettify(s)
}

func (s ListSharedCustomRulesResponseDataConditionPathConditions) GoString() string {
  return s.String()
}

func (s *ListSharedCustomRulesResponseDataConditionPathConditions) SetMatchType(v string) *ListSharedCustomRulesResponseDataConditionPathConditions {
  s.MatchType = &v
  return s
}

func (s *ListSharedCustomRulesResponseDataConditionPathConditions) SetListSharedCustomRulesPaths(v []*string) *ListSharedCustomRulesResponseDataConditionPathConditions {
  s.ListSharedCustomRulesPaths = v
  return s
}

type ListSharedCustomRulesResponseDataConditionUaConditions struct     {
  // {"en":"Match type.\nEQUAL: equal to\nNOT_EQUAL: not equal to\nCONTAIN: contains\nNOT_CONTAIN: does not contain\nREGEX: regular\nNONE: empty or does not exist\nNOT_REGEX: regular does not match\nSTART_WITH: starts with\nEND_WITH: ends with\nWILDCARD: wildcard matches\nNOT_WILDCARD: wildcard does not match","zh_CN":"匹配类型。\nEQUAL：等于\nNOT_EQUAL：不等于\nCONTAIN：包含\nNOT_CONTAIN：不包含\nREGEX：正则\nNONE：为空或不存在\nNOT_REGEX：正则不匹配\nSTART_WITH：开头是\nEND_WITH：结尾是\nWILDCARD：通配符匹配\nNOT_WILDCARD：通配符不匹配","exampleValue":"EQUAL,NOT_EQUAL,CONTAIN,NOT_CONTAIN,NONE,REGEX,NOT_REGEX,START_WITH,END_WITH,WILDCARD,NOT_WILDCARD"}
  MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
  // {"en":"User-Agent.","zh_CN":"User-Agent。"}
  Ua []*string `json:"ua,omitempty" xml:"ua,omitempty" require:"true" type:"Repeated"`
}

func (s ListSharedCustomRulesResponseDataConditionUaConditions) String() string {
  return tea.Prettify(s)
}

func (s ListSharedCustomRulesResponseDataConditionUaConditions) GoString() string {
  return s.String()
}

func (s *ListSharedCustomRulesResponseDataConditionUaConditions) SetMatchType(v string) *ListSharedCustomRulesResponseDataConditionUaConditions {
  s.MatchType = &v
  return s
}

func (s *ListSharedCustomRulesResponseDataConditionUaConditions) SetUa(v []*string) *ListSharedCustomRulesResponseDataConditionUaConditions {
  s.Ua = v
  return s
}

type ListSharedCustomRulesResponseDataConditionHeaderConditions struct     {
  // {"en":"Match type.\nEQUAL: equal to\nNOT_EQUAL: not equal to\nCONTAIN: contains\nNOT_CONTAIN: does not contain\nREGEX: regular\nNONE: empty or does not exist\nNOT_REGEX: regular does not match\nSTART_WITH: starts with\nEND_WITH: ends with\nWILDCARD: wildcard matches\nNOT_WILDCARD: wildcard does not match","zh_CN":"匹配类型。\nEQUAL：等于\nNOT_EQUAL：不等于\nCONTAIN：包含\nNOT_CONTAIN：不包含\nREGEX：正则\nNONE：为空或不存在\nNOT_REGEX：正则不匹配\nSTART_WITH：开头是\nEND_WITH：结尾是\nWILDCARD：通配符匹配\nNOT_WILDCARD：通配符不匹配","exampleValue":"EQUAL,NOT_EQUAL,CONTAIN,NOT_CONTAIN,NONE,REGEX,NOT_REGEX,START_WITH,END_WITH,WILDCARD,NOT_WILDCARD"}
  MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
  // {"en":"Header value.","zh_CN":"头部值。"}
  ValueList []*string `json:"valueList,omitempty" xml:"valueList,omitempty" require:"true" type:"Repeated"`
  // {"en":"Request header name.","zh_CN":"头部名称。"}
  Key *string `json:"key,omitempty" xml:"key,omitempty" require:"true"`
}

func (s ListSharedCustomRulesResponseDataConditionHeaderConditions) String() string {
  return tea.Prettify(s)
}

func (s ListSharedCustomRulesResponseDataConditionHeaderConditions) GoString() string {
  return s.String()
}

func (s *ListSharedCustomRulesResponseDataConditionHeaderConditions) SetMatchType(v string) *ListSharedCustomRulesResponseDataConditionHeaderConditions {
  s.MatchType = &v
  return s
}

func (s *ListSharedCustomRulesResponseDataConditionHeaderConditions) SetValueList(v []*string) *ListSharedCustomRulesResponseDataConditionHeaderConditions {
  s.ValueList = v
  return s
}

func (s *ListSharedCustomRulesResponseDataConditionHeaderConditions) SetKey(v string) *ListSharedCustomRulesResponseDataConditionHeaderConditions {
  s.Key = &v
  return s
}

type ListSharedCustomRulesResponseDataConditionRefererConditions struct     {
  // {"en":"Referer.","zh_CN":"Referer。"}
  Referer []*string `json:"referer,omitempty" xml:"referer,omitempty" require:"true" type:"Repeated"`
  // {"en":"Match type.\nEQUAL: equal to\nNOT_EQUAL: not equal to\nCONTAIN: contains\nNOT_CONTAIN: does not contain\nREGEX: regular\nNONE: empty or does not exist\nNOT_REGEX: regular does not match\nSTART_WITH: starts with\nEND_WITH: ends with\nWILDCARD: wildcard matches\nNOT_WILDCARD: wildcard does not match","zh_CN":"匹配类型。\nEQUAL：等于\nNOT_EQUAL：不等于\nCONTAIN：包含\nNOT_CONTAIN：不包含\nREGEX：正则\nNONE：为空或不存在\nNOT_REGEX：正则不匹配\nSTART_WITH：开头是\nEND_WITH：结尾是\nWILDCARD：通配符匹配\nNOT_WILDCARD：通配符不匹配","exampleValue":"EQUAL,NOT_EQUAL,CONTAIN,NOT_CONTAIN,NONE,REGEX,NOT_REGEX,START_WITH,END_WITH,WILDCARD,NOT_WILDCARD"}
  MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
}

func (s ListSharedCustomRulesResponseDataConditionRefererConditions) String() string {
  return tea.Prettify(s)
}

func (s ListSharedCustomRulesResponseDataConditionRefererConditions) GoString() string {
  return s.String()
}

func (s *ListSharedCustomRulesResponseDataConditionRefererConditions) SetReferer(v []*string) *ListSharedCustomRulesResponseDataConditionRefererConditions {
  s.Referer = v
  return s
}

func (s *ListSharedCustomRulesResponseDataConditionRefererConditions) SetMatchType(v string) *ListSharedCustomRulesResponseDataConditionRefererConditions {
  s.MatchType = &v
  return s
}

type ListSharedCustomRulesResponseDataConditionJa3Conditions struct     {
  // {"en":"Match type.\nEQUAL: Equals\nNOT_EQUAL: Does not equal","zh_CN":"匹配类型。\nEQUAL：等于 \nNOT_EQUAL：不等于","exampleValue":"EQUAL,NOT_EQUAL"}
  MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
  // {"en":"JA3 Fingerprint List.","zh_CN":"JA3指纹列表。"}
  Ja3List []*string `json:"ja3List,omitempty" xml:"ja3List,omitempty" require:"true" type:"Repeated"`
}

func (s ListSharedCustomRulesResponseDataConditionJa3Conditions) String() string {
  return tea.Prettify(s)
}

func (s ListSharedCustomRulesResponseDataConditionJa3Conditions) GoString() string {
  return s.String()
}

func (s *ListSharedCustomRulesResponseDataConditionJa3Conditions) SetMatchType(v string) *ListSharedCustomRulesResponseDataConditionJa3Conditions {
  s.MatchType = &v
  return s
}

func (s *ListSharedCustomRulesResponseDataConditionJa3Conditions) SetJa3List(v []*string) *ListSharedCustomRulesResponseDataConditionJa3Conditions {
  s.Ja3List = v
  return s
}

type ListSharedCustomRulesResponseDataConditionJa4Conditions struct     {
  // {"en":"Match type. \nEQUAL: Equals\nNOT_EQUAL: Does not equal\nCONTAIN: Contains\nNOT_CONTAIN: Does not Contains\nSTART_WITH: Starts with\nEND_WITH: Ends with\nWILDCARD: Wildcard matches, ** represents zero or more arbitrary characters, ? represents any single character\nNOT_WILDCARD: Wildcard does not match, ** represents zero or more arbitrary characters, ? represents any single character","zh_CN":"匹配类型。\nEQUAL：等于\nNOT_EQUAL：不等于\nCONTAIN：包含\nNOT_CONTAIN：不包含\nSTART_WITH：开头是\nEND_WITH：结尾是\nWILDCARD：通配符匹配，*代表零个或多个任意字符，?代表任意单个字符\nNOT_WILDCARD：通配符不匹配，*代表零个或多个任意字符，?代表任意单个字符","exampleValue":"EQUAL,NOT_EQUAL,CONTAIN,NOT_CONTAIN,START_WITH,END_WITH,WILDCARD,NOT_WILDCARD"}
  MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
  // {"en":"JA4 Fingerprint List.","zh_CN":"JA4指纹列表。"}
  Ja4List []*string `json:"ja4List,omitempty" xml:"ja4List,omitempty" require:"true" type:"Repeated"`
}

func (s ListSharedCustomRulesResponseDataConditionJa4Conditions) String() string {
  return tea.Prettify(s)
}

func (s ListSharedCustomRulesResponseDataConditionJa4Conditions) GoString() string {
  return s.String()
}

func (s *ListSharedCustomRulesResponseDataConditionJa4Conditions) SetMatchType(v string) *ListSharedCustomRulesResponseDataConditionJa4Conditions {
  s.MatchType = &v
  return s
}

func (s *ListSharedCustomRulesResponseDataConditionJa4Conditions) SetJa4List(v []*string) *ListSharedCustomRulesResponseDataConditionJa4Conditions {
  s.Ja4List = v
  return s
}

type ListSharedCustomRulesResponseHeader struct {
}

func (s ListSharedCustomRulesResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ListSharedCustomRulesResponseHeader) GoString() string {
  return s.String()
}




type DisassociateShareRateLimitRequest struct {
  // {"en":"Share rule ID.","zh_CN":"共享规则ID。"}
  ShareId *string `json:"shareId,omitempty" xml:"shareId,omitempty" require:"true"`
  // {"en":"List of domain to be disassociated.","zh_CN":"待取消关联域名列表。"}
  DomainList []*string `json:"domainList,omitempty" xml:"domainList,omitempty" require:"true" type:"Repeated"`
}

func (s DisassociateShareRateLimitRequest) String() string {
  return tea.Prettify(s)
}

func (s DisassociateShareRateLimitRequest) GoString() string {
  return s.String()
}

func (s *DisassociateShareRateLimitRequest) SetShareId(v string) *DisassociateShareRateLimitRequest {
  s.ShareId = &v
  return s
}

func (s *DisassociateShareRateLimitRequest) SetDomainList(v []*string) *DisassociateShareRateLimitRequest {
  s.DomainList = v
  return s
}

type DisassociateShareRateLimitRequestHeader struct {
  // {"en":"Security service type. Please enter a specific service type, if you purchase multiple security services.","zh_CN":"安全服务类型。有使用多个不同的安全服务时，需要填写具体的服务类型。"}
  ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty"`
}

func (s DisassociateShareRateLimitRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s DisassociateShareRateLimitRequestHeader) GoString() string {
  return s.String()
}

func (s *DisassociateShareRateLimitRequestHeader) SetServiceType(v string) *DisassociateShareRateLimitRequestHeader {
  s.ServiceType = &v
  return s
}

type DisassociateShareRateLimitPaths struct {
}

func (s DisassociateShareRateLimitPaths) String() string {
  return tea.Prettify(s)
}

func (s DisassociateShareRateLimitPaths) GoString() string {
  return s.String()
}

type DisassociateShareRateLimitParameters struct {
}

func (s DisassociateShareRateLimitParameters) String() string {
  return tea.Prettify(s)
}

func (s DisassociateShareRateLimitParameters) GoString() string {
  return s.String()
}

type DisassociateShareRateLimitResponse struct {
  // {"dictionary":"belong=WAAP-MS-Ext|dict=waap_retCodeEnum","en":"Please refer to the error code for exceptions.","zh_CN":"请参照错误码。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Description.","zh_CN":"描述信息。"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
}

func (s DisassociateShareRateLimitResponse) String() string {
  return tea.Prettify(s)
}

func (s DisassociateShareRateLimitResponse) GoString() string {
  return s.String()
}

func (s *DisassociateShareRateLimitResponse) SetCode(v string) *DisassociateShareRateLimitResponse {
  s.Code = &v
  return s
}

func (s *DisassociateShareRateLimitResponse) SetMsg(v string) *DisassociateShareRateLimitResponse {
  s.Msg = &v
  return s
}

type DisassociateShareRateLimitResponseHeader struct {
}

func (s DisassociateShareRateLimitResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s DisassociateShareRateLimitResponseHeader) GoString() string {
  return s.String()
}




type UpdateSharedWAFRuleExceptionRequest struct {
  // {"en":"Rule exception ID.","zh_CN":"规则例外ID。"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {"en":"Exception name, maximum 50 characters.\nDoes not support # and &.","zh_CN":"例外名称，最多50个字符。\n不支持 # 和 &。"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"Exception description,maximum 200 characters.","zh_CN":"例外描述，最多200个字符。"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty"`
  // {"en":"Matching conditions.\nip: IP\npath: Path\nuri: URI\nurlParamName: URI Parameter Name\nurlParamValue: URI Parameter Value\nuserAgent: User Agent\nhttpHeaderName: Request Header Name\nhttpHeaderValue: Request Header Value\ncookie: Cookie\nbody: Body\nbodyParamName: Body Parameter Name\nbodyParamValue: Body Parameter Value","zh_CN":"匹配条件。\nip：IP\npath：路径\nuri：URI\nurlParamName：URI参数名\nurlParamValue：URI参数值\nuserAgent：User Agent\nhttpHeaderName：请求头部名称\nhttpHeaderValue：请求头部值\ncookie：Cookie\nbody：Body\nbodyParamName：Body参数名\nbodyParamValue：Body参数值","exampleValue":"ip,path,uri,urlParamName,urlParamValue,userAgent,httpHeaderName,httpHeaderValue,cookie,body,bodyParamName,bodyParamValue"}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
  // {"en":"Match type,IP can only be EQUAL.\nEQUAL: Equal\nCONTAIN: Contains\nREGEX: Regular match","zh_CN":"匹配类型，IP只能是等于。\nEQUAL：等于\nCONTAIN：包含\nREGEX：正则匹配","exampleValue":"EQUAL,CONTAIN,REGEX"}
  MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
  // {"en":"Rule exceptions.\nWhen matchType=EQUAL, case-sensitive, path and uri must start with \"/\", and body can only pass one value;\nWhen matchType=REGEX, only one value can be passed","zh_CN":"规则例外内容。\nmatchType=EQUAL时，大小写敏感，path和uri必须以\"/\"开头，body只能传一个值；\nmatchType=REGEX时，只能传一个值。"}
  ContentList []*string `json:"contentList,omitempty" xml:"contentList,omitempty" require:"true" type:"Repeated"`
}

func (s UpdateSharedWAFRuleExceptionRequest) String() string {
  return tea.Prettify(s)
}

func (s UpdateSharedWAFRuleExceptionRequest) GoString() string {
  return s.String()
}

func (s *UpdateSharedWAFRuleExceptionRequest) SetId(v string) *UpdateSharedWAFRuleExceptionRequest {
  s.Id = &v
  return s
}

func (s *UpdateSharedWAFRuleExceptionRequest) SetName(v string) *UpdateSharedWAFRuleExceptionRequest {
  s.Name = &v
  return s
}

func (s *UpdateSharedWAFRuleExceptionRequest) SetMsg(v string) *UpdateSharedWAFRuleExceptionRequest {
  s.Msg = &v
  return s
}

func (s *UpdateSharedWAFRuleExceptionRequest) SetType(v string) *UpdateSharedWAFRuleExceptionRequest {
  s.Type = &v
  return s
}

func (s *UpdateSharedWAFRuleExceptionRequest) SetMatchType(v string) *UpdateSharedWAFRuleExceptionRequest {
  s.MatchType = &v
  return s
}

func (s *UpdateSharedWAFRuleExceptionRequest) SetContentList(v []*string) *UpdateSharedWAFRuleExceptionRequest {
  s.ContentList = v
  return s
}

type UpdateSharedWAFRuleExceptionRequestHeader struct {
  // {"en":"Security service type. Please enter a specific service type, if you purchase multiple security services.","zh_CN":"安全服务类型。有使用多个不同的安全服务时，需要填写具体的服务类型。"}
  ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty"`
}

func (s UpdateSharedWAFRuleExceptionRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s UpdateSharedWAFRuleExceptionRequestHeader) GoString() string {
  return s.String()
}

func (s *UpdateSharedWAFRuleExceptionRequestHeader) SetServiceType(v string) *UpdateSharedWAFRuleExceptionRequestHeader {
  s.ServiceType = &v
  return s
}

type UpdateSharedWAFRuleExceptionPaths struct {
}

func (s UpdateSharedWAFRuleExceptionPaths) String() string {
  return tea.Prettify(s)
}

func (s UpdateSharedWAFRuleExceptionPaths) GoString() string {
  return s.String()
}

type UpdateSharedWAFRuleExceptionParameters struct {
}

func (s UpdateSharedWAFRuleExceptionParameters) String() string {
  return tea.Prettify(s)
}

func (s UpdateSharedWAFRuleExceptionParameters) GoString() string {
  return s.String()
}

type UpdateSharedWAFRuleExceptionResponse struct {
  // {"dictionary":"belong=WAAP-MS-Ext|dict=waap_retCodeEnum","en":"Please refer to the error code for exceptions.","zh_CN":"请参照错误码。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Description.","zh_CN":"描述信息。"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"Modification success indicator.","zh_CN":"修改成功标识。"}
  Data *bool `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s UpdateSharedWAFRuleExceptionResponse) String() string {
  return tea.Prettify(s)
}

func (s UpdateSharedWAFRuleExceptionResponse) GoString() string {
  return s.String()
}

func (s *UpdateSharedWAFRuleExceptionResponse) SetCode(v string) *UpdateSharedWAFRuleExceptionResponse {
  s.Code = &v
  return s
}

func (s *UpdateSharedWAFRuleExceptionResponse) SetMsg(v string) *UpdateSharedWAFRuleExceptionResponse {
  s.Msg = &v
  return s
}

func (s *UpdateSharedWAFRuleExceptionResponse) SetData(v bool) *UpdateSharedWAFRuleExceptionResponse {
  s.Data = &v
  return s
}

type UpdateSharedWAFRuleExceptionResponseHeader struct {
}

func (s UpdateSharedWAFRuleExceptionResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s UpdateSharedWAFRuleExceptionResponseHeader) GoString() string {
  return s.String()
}




type UpdateCustomActionRequest struct {
  // {"en":"Response content definition, the size of the response content cannot exceed 16KB, if you need to insert static resources such as pictures, it is recommended to use links to import.\nThe following interception information parameters are supported:\n{url} : Displays the URL information when intercepting\n{client_ip} : Displays the intercepted user IP information\n{time} : Displays the intercepted time point\n{event_id} : Displays the ID information of this interception event","zh_CN":"自定义响应内容定义，响应内容大小不能超过16KB，如需插入图片等静态资源，建议采用链接引入。\n支持引用下列拦截信息参数：\n{url}  ：显示拦截时的URL信息\n{client_ip}  ：显示被拦截的用户IP信息\n{time}  ：显示被拦截的时间点\n{event_id}  ：显示本次拦截事件的ID信息"}
  CusBlockCnt *string `json:"cusBlockCnt,omitempty" xml:"cusBlockCnt,omitempty"`
  // {"en":"Description,maximum 200 characters.","zh_CN":"描述，最多200个字符。"}
  Description *string `json:"description,omitempty" xml:"description,omitempty"`
  // {"en":"Custom response ID.","zh_CN":"自定义响应ID。"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {"en":"Response Content-Type type,multiple separated by ; sign.","zh_CN":"响应Content-Type类型，多个以 ; 隔开。"}
  ContentType *string `json:"contentType,omitempty" xml:"contentType,omitempty"`
  // {"en":"Action name,maximum 50 characters.\nDoes not support # and &.","zh_CN":"动作名称，最多50个字符。\n不支持 # 和 &。"}
  ActionName *string `json:"actionName,omitempty" xml:"actionName,omitempty"`
  // {"en":"Status code.Valid value range:100 - 299,400 - 999.","zh_CN":"状态码。有效值范围：100 - 299，400 - 999。"}
  StatusCode *string `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s UpdateCustomActionRequest) String() string {
  return tea.Prettify(s)
}

func (s UpdateCustomActionRequest) GoString() string {
  return s.String()
}

func (s *UpdateCustomActionRequest) SetCusBlockCnt(v string) *UpdateCustomActionRequest {
  s.CusBlockCnt = &v
  return s
}

func (s *UpdateCustomActionRequest) SetDescription(v string) *UpdateCustomActionRequest {
  s.Description = &v
  return s
}

func (s *UpdateCustomActionRequest) SetId(v string) *UpdateCustomActionRequest {
  s.Id = &v
  return s
}

func (s *UpdateCustomActionRequest) SetContentType(v string) *UpdateCustomActionRequest {
  s.ContentType = &v
  return s
}

func (s *UpdateCustomActionRequest) SetActionName(v string) *UpdateCustomActionRequest {
  s.ActionName = &v
  return s
}

func (s *UpdateCustomActionRequest) SetStatusCode(v string) *UpdateCustomActionRequest {
  s.StatusCode = &v
  return s
}

type UpdateCustomActionRequestHeader struct {
  // {"en":"Security service type. Please enter a specific service type, if you purchase multiple security services.","zh_CN":"安全服务类型。有使用多个不同的安全服务时，需要填写具体的服务类型。"}
  ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty"`
}

func (s UpdateCustomActionRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s UpdateCustomActionRequestHeader) GoString() string {
  return s.String()
}

func (s *UpdateCustomActionRequestHeader) SetServiceType(v string) *UpdateCustomActionRequestHeader {
  s.ServiceType = &v
  return s
}

type UpdateCustomActionPaths struct {
}

func (s UpdateCustomActionPaths) String() string {
  return tea.Prettify(s)
}

func (s UpdateCustomActionPaths) GoString() string {
  return s.String()
}

type UpdateCustomActionParameters struct {
}

func (s UpdateCustomActionParameters) String() string {
  return tea.Prettify(s)
}

func (s UpdateCustomActionParameters) GoString() string {
  return s.String()
}

type UpdateCustomActionResponse struct {
  // {"en":"Description.","zh_CN":"描述信息。"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"dictionary":"belong=WAAP-MS-Ext|dict=waap_retCodeEnum","en":"Please refer to the error code for exceptions.","zh_CN":"请参照错误码。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
}

func (s UpdateCustomActionResponse) String() string {
  return tea.Prettify(s)
}

func (s UpdateCustomActionResponse) GoString() string {
  return s.String()
}

func (s *UpdateCustomActionResponse) SetMsg(v string) *UpdateCustomActionResponse {
  s.Msg = &v
  return s
}

func (s *UpdateCustomActionResponse) SetCode(v string) *UpdateCustomActionResponse {
  s.Code = &v
  return s
}

type UpdateCustomActionResponseHeader struct {
}

func (s UpdateCustomActionResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s UpdateCustomActionResponseHeader) GoString() string {
  return s.String()
}




type DeleteCustomActionsRequest struct {
  // {"en":"Custom response ID list.","zh_CN":"自定义响应ID列表。"}
  IdList []*string `json:"idList,omitempty" xml:"idList,omitempty" require:"true" type:"Repeated"`
}

func (s DeleteCustomActionsRequest) String() string {
  return tea.Prettify(s)
}

func (s DeleteCustomActionsRequest) GoString() string {
  return s.String()
}

func (s *DeleteCustomActionsRequest) SetIdList(v []*string) *DeleteCustomActionsRequest {
  s.IdList = v
  return s
}

type DeleteCustomActionsRequestHeader struct {
  // {"en":"Security service type. Please enter a specific service type, if you purchase multiple security services.","zh_CN":"安全服务类型。有使用多个不同的安全服务时，需要填写具体的服务类型。"}
  ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty"`
}

func (s DeleteCustomActionsRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteCustomActionsRequestHeader) GoString() string {
  return s.String()
}

func (s *DeleteCustomActionsRequestHeader) SetServiceType(v string) *DeleteCustomActionsRequestHeader {
  s.ServiceType = &v
  return s
}

type DeleteCustomActionsPaths struct {
}

func (s DeleteCustomActionsPaths) String() string {
  return tea.Prettify(s)
}

func (s DeleteCustomActionsPaths) GoString() string {
  return s.String()
}

type DeleteCustomActionsParameters struct {
}

func (s DeleteCustomActionsParameters) String() string {
  return tea.Prettify(s)
}

func (s DeleteCustomActionsParameters) GoString() string {
  return s.String()
}

type DeleteCustomActionsResponse struct {
  // {"dictionary":"belong=WAAP-MS-Ext|dict=waap_retCodeEnum","en":"Please refer to the error code for exceptions.","zh_CN":"请参照错误码。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Description.","zh_CN":"描述信息。"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
}

func (s DeleteCustomActionsResponse) String() string {
  return tea.Prettify(s)
}

func (s DeleteCustomActionsResponse) GoString() string {
  return s.String()
}

func (s *DeleteCustomActionsResponse) SetCode(v string) *DeleteCustomActionsResponse {
  s.Code = &v
  return s
}

func (s *DeleteCustomActionsResponse) SetMsg(v string) *DeleteCustomActionsResponse {
  s.Msg = &v
  return s
}

type DeleteCustomActionsResponseHeader struct {
}

func (s DeleteCustomActionsResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteCustomActionsResponseHeader) GoString() string {
  return s.String()
}




type QueryAppApiExceptionListRequest struct {
  // {"en":"Whether to return the number of associated hostnames.Default:false.","zh_CN":"是否需要返回关联域名数，默认为false"}
  NeedDomain *bool `json:"needDomain,omitempty" xml:"needDomain,omitempty"`
}

func (s QueryAppApiExceptionListRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryAppApiExceptionListRequest) GoString() string {
  return s.String()
}

func (s *QueryAppApiExceptionListRequest) SetNeedDomain(v bool) *QueryAppApiExceptionListRequest {
  s.NeedDomain = &v
  return s
}

type QueryAppApiExceptionListRequestHeader struct {
  // {"en":"Security service type. Please enter a specific service type, if you purchase multiple security services.","zh_CN":"安全服务类型。有使用多个不同的安全服务时，需要填写具体的服务类型。"}
  ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty"`
}

func (s QueryAppApiExceptionListRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryAppApiExceptionListRequestHeader) GoString() string {
  return s.String()
}

func (s *QueryAppApiExceptionListRequestHeader) SetServiceType(v string) *QueryAppApiExceptionListRequestHeader {
  s.ServiceType = &v
  return s
}

type QueryAppApiExceptionListPaths struct {
}

func (s QueryAppApiExceptionListPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryAppApiExceptionListPaths) GoString() string {
  return s.String()
}

type QueryAppApiExceptionListParameters struct {
}

func (s QueryAppApiExceptionListParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryAppApiExceptionListParameters) GoString() string {
  return s.String()
}

type QueryAppApiExceptionListResponse struct {
  // {"dictionary":"belong=WAAP-MS-Ext|dict=waap_retCodeEnum","en":"Please refer to the error code for exceptions.","zh_CN":"请参照错误码。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Description.","zh_CN":"描述信息。"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"Data.","zh_CN":"出参数据。"}
  Data []*QueryAppApiExceptionListResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s QueryAppApiExceptionListResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryAppApiExceptionListResponse) GoString() string {
  return s.String()
}

func (s *QueryAppApiExceptionListResponse) SetCode(v string) *QueryAppApiExceptionListResponse {
  s.Code = &v
  return s
}

func (s *QueryAppApiExceptionListResponse) SetMsg(v string) *QueryAppApiExceptionListResponse {
  s.Msg = &v
  return s
}

func (s *QueryAppApiExceptionListResponse) SetData(v []*QueryAppApiExceptionListResponseData) *QueryAppApiExceptionListResponse {
  s.Data = v
  return s
}

type QueryAppApiExceptionListResponseData struct     {
  // {"en":"Rule ID","zh_CN":"规则ID"}
  RuleId *string `json:"ruleId,omitempty" xml:"ruleId,omitempty" require:"true"`
  // {"en":"Rule name, duplicate not allowed.","zh_CN":"规则名称，不允许重复"}
  RuleName *string `json:"ruleName,omitempty" xml:"ruleName,omitempty" require:"true"`
  // {"en":"Characteristic.","zh_CN":"特征。"}
  Config *QueryAppApiExceptionListResponseDataConfig `json:"config,omitempty" xml:"config,omitempty" require:"true" type:"Struct"`
  // {"en":"Type enumeration, for example:\nNATIVE_ APP        remarks: Native App\nCallback_ API       remarks: Callback API\nOther_ API         remarks: Other program API\nHYBRID_ APP     remarks: Hybrid APP","zh_CN":"类型枚举，例如：\nNATIVE_APP          备注：原生APP\nCALLBACK_API      备注：回调API\nOTHER_API            备注：其他程序API\nHYBRID_APP          备注：混合APP","exampleValue":"NATIVE_APP,CALLBACK_API,OTHER_API,HYBRID_APP"}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
  // {"en":"Number of associated hostnames","zh_CN":"关联域名数量"}
  DomainNum *int `json:"domainNum,omitempty" xml:"domainNum,omitempty" require:"true"`
  // {"en":"Creator","zh_CN":"创建者"}
  Creator *string `json:"creator,omitempty" xml:"creator,omitempty" require:"true"`
}

func (s QueryAppApiExceptionListResponseData) String() string {
  return tea.Prettify(s)
}

func (s QueryAppApiExceptionListResponseData) GoString() string {
  return s.String()
}

func (s *QueryAppApiExceptionListResponseData) SetRuleId(v string) *QueryAppApiExceptionListResponseData {
  s.RuleId = &v
  return s
}

func (s *QueryAppApiExceptionListResponseData) SetRuleName(v string) *QueryAppApiExceptionListResponseData {
  s.RuleName = &v
  return s
}

func (s *QueryAppApiExceptionListResponseData) SetConfig(v *QueryAppApiExceptionListResponseDataConfig) *QueryAppApiExceptionListResponseData {
  s.Config = v
  return s
}

func (s *QueryAppApiExceptionListResponseData) SetType(v string) *QueryAppApiExceptionListResponseData {
  s.Type = &v
  return s
}

func (s *QueryAppApiExceptionListResponseData) SetDomainNum(v int) *QueryAppApiExceptionListResponseData {
  s.DomainNum = &v
  return s
}

func (s *QueryAppApiExceptionListResponseData) SetCreator(v string) *QueryAppApiExceptionListResponseData {
  s.Creator = &v
  return s
}

type QueryAppApiExceptionListResponseDataConfig struct {
  // {"en":"Matching condition name.\nPATH         Path\nURI            URI\nUA             User-Agent\nREFERER    Referer\nHEADER    Request Header","zh_CN":"匹配条件名称。\nPATH                       路径\nURI                          URI\nUA                           User-Agent\nREFERER                  Referer\nHEAD                      Request Header","exampleValue":"PATH,URI,UA,REFERER,HEAD"}
  Condition *string `json:"condition,omitempty" xml:"condition,omitempty" require:"true"`
  // {"en":"Matching condition function.\nCONTAIN                  contains\nNOT_CONTAIN         does not contain\nEQUAL                      equals\nNOT_EQUAL             does not equal\nEMPTY                      does not exist or has no value(When the condition is a path or URI, this value is not optional)\nREGEX                      regex match","zh_CN":"匹配条件函数。\nCONTAIN                  包含\nNOT_CONTAIN         不包含\nEQUAL                      等于\nNOT_EQUAL             不等于\nEMPTY                      为空或不存在（当condition为路径或者URI时，该值不可选）\nREGEX                       匹配正则","exampleValue":"CONTAIN,NOT_CONTAIN,EQUAL,NOT_EQUAL,EMPTY,REGEX"}
  MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
  // {"en":"Condition value list, When matchType is EMPTY, the value is empty;  When matchType is REGEX, only one item is set","zh_CN":"条件值列表，当matchType为EMPTY时，该值为空；当matchType为REGEX时，只能传一条"}
  ConditionValue []*string `json:"conditionValue,omitempty" xml:"conditionValue,omitempty" require:"true" type:"Repeated"`
  // {"en":"Head value, If condition=HEAD, then this field is mandatory.","zh_CN":"头部值，如果condition为HEAD，则该字段必填"}
  Key *string `json:"key,omitempty" xml:"key,omitempty" require:"true"`
}

func (s QueryAppApiExceptionListResponseDataConfig) String() string {
  return tea.Prettify(s)
}

func (s QueryAppApiExceptionListResponseDataConfig) GoString() string {
  return s.String()
}

func (s *QueryAppApiExceptionListResponseDataConfig) SetCondition(v string) *QueryAppApiExceptionListResponseDataConfig {
  s.Condition = &v
  return s
}

func (s *QueryAppApiExceptionListResponseDataConfig) SetMatchType(v string) *QueryAppApiExceptionListResponseDataConfig {
  s.MatchType = &v
  return s
}

func (s *QueryAppApiExceptionListResponseDataConfig) SetConditionValue(v []*string) *QueryAppApiExceptionListResponseDataConfig {
  s.ConditionValue = v
  return s
}

func (s *QueryAppApiExceptionListResponseDataConfig) SetKey(v string) *QueryAppApiExceptionListResponseDataConfig {
  s.Key = &v
  return s
}

type QueryAppApiExceptionListResponseHeader struct {
}

func (s QueryAppApiExceptionListResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryAppApiExceptionListResponseHeader) GoString() string {
  return s.String()
}




type CreateSharedCustomRuleRequest struct {
  // {"en":"Matching conditions.\nExcept for header conditions, there can only be at most one record per match type under each type of condition.","zh_CN":"匹配条件。\n除了请求头条件，其它类型的条件下一种匹配类型最多只能有一条记录。"}
  Condition *CreateSharedCustomRuleRequestCondition `json:"condition,omitempty" xml:"condition,omitempty" require:"true" type:"Struct"`
  // {"en":"Action.\nNO_USE: Not Used\nLOG: Log\nDELAY: Delay\nBLOCK: Deny\nRESET: Reset Connection\nCUSTOM_ACTION_ID:Fill in the custom action id of the corresponding action","zh_CN":"处理动作。\nNO_USE：不使用\nLOG：监控\nDELAY：延迟响应\nBLOCK：拦截\nRESET：断开连接\nCUSTOM_ACTION_ID：传入对应的自定义响应动作的ID","exampleValue":"NO_USE,LOG,DELAY,BLOCK,RESET"}
  Act *string `json:"act,omitempty" xml:"act,omitempty" require:"true"`
  // {"en":"Rule Name, maximum 100 characters.Does not support # and &.","zh_CN":"规则名称，最多100个字符。不支持 # 和 &。"}
  RuleName *string `json:"ruleName,omitempty" xml:"ruleName,omitempty" require:"true"`
  // {"en":"Description, maximum 1000 characters.","zh_CN":"规则描述，最多1000个字符。"}
  Description *string `json:"description,omitempty" xml:"description,omitempty"`
}

func (s CreateSharedCustomRuleRequest) String() string {
  return tea.Prettify(s)
}

func (s CreateSharedCustomRuleRequest) GoString() string {
  return s.String()
}

func (s *CreateSharedCustomRuleRequest) SetCondition(v *CreateSharedCustomRuleRequestCondition) *CreateSharedCustomRuleRequest {
  s.Condition = v
  return s
}

func (s *CreateSharedCustomRuleRequest) SetAct(v string) *CreateSharedCustomRuleRequest {
  s.Act = &v
  return s
}

func (s *CreateSharedCustomRuleRequest) SetRuleName(v string) *CreateSharedCustomRuleRequest {
  s.RuleName = &v
  return s
}

func (s *CreateSharedCustomRuleRequest) SetDescription(v string) *CreateSharedCustomRuleRequest {
  s.Description = &v
  return s
}

type CreateSharedCustomRuleRequestCondition struct {
  // {"en":"Request Method.\nWhen the business scenario is API,this matching condition is not supported.","zh_CN":"请求方法，匹配类型不可重复。\n当业务场景为API业务时不支持此匹配条件。"}
  MethodConditions []*CreateSharedCustomRuleRequestConditionMethodConditions `json:"methodConditions,omitempty" xml:"methodConditions,omitempty" type:"Repeated"`
  // {"en":"Geo, match type cannot be repeated.","zh_CN":"区域，匹配类型不可重复。"}
  AreaConditions []*CreateSharedCustomRuleRequestConditionAreaConditions `json:"areaConditions,omitempty" xml:"areaConditions,omitempty" type:"Repeated"`
  // {"en":"IP/CIDR, match type cannot be repeated.","zh_CN":"IP/IP段，匹配类型不可重复。"}
  IpOrIpsConditions []*CreateSharedCustomRuleRequestConditionIpOrIpsConditions `json:"ipOrIpsConditions,omitempty" xml:"ipOrIpsConditions,omitempty" type:"Repeated"`
  // {"en":"URI, match type cannot be repeated.\nWhen the business scenario is API, this matching condition is not supported.","zh_CN":"URI，匹配类型不可重复。\n当业务场景为API业务时不支持此匹配条件。"}
  UriConditions []*CreateSharedCustomRuleRequestConditionUriConditions `json:"uriConditions,omitempty" xml:"uriConditions,omitempty" type:"Repeated"`
  // {"en":"Path, match type cannot be repeated.\nWhen the business scenario is API, this matching condition is not supported.","zh_CN":"路径，匹配类型不可重复。\n当业务场景为API业务时不支持此匹配条件。"}
  PathConditions []*CreateSharedCustomRuleRequestConditionPathConditions `json:"pathConditions,omitempty" xml:"pathConditions,omitempty" type:"Repeated"`
  // {"en":"User Agent, match type cannot be repeated.","zh_CN":"User-Agent，匹配类型不可重复。"}
  UaConditions []*CreateSharedCustomRuleRequestConditionUaConditions `json:"uaConditions,omitempty" xml:"uaConditions,omitempty" type:"Repeated"`
  // {"en":"Request Header, match type can be repeated.","zh_CN":"请求头，匹配类型可重复。"}
  HeaderConditions []*CreateSharedCustomRuleRequestConditionHeaderConditions `json:"headerConditions,omitempty" xml:"headerConditions,omitempty" type:"Repeated"`
  // {"en":"Referer, match type cannot be repeated.","zh_CN":"Referer，匹配类型不可重复。"}
  RefererConditions []*CreateSharedCustomRuleRequestConditionRefererConditions `json:"refererConditions,omitempty" xml:"refererConditions,omitempty" type:"Repeated"`
  // {"en":"JA3 Fingerprint, match type cannot be repeated.","zh_CN":"JA3指纹，匹配类型不可重复。"}
  Ja3Conditions []*CreateSharedCustomRuleRequestConditionJa3Conditions `json:"ja3Conditions,omitempty" xml:"ja3Conditions,omitempty" type:"Repeated"`
  // {"en":"JA4 Fingerprint, match type cannot be repeated.","zh_CN":"JA4指纹，匹配类型不可重复。"}
  Ja4Conditions []*CreateSharedCustomRuleRequestConditionJa4Conditions `json:"ja4Conditions,omitempty" xml:"ja4Conditions,omitempty" type:"Repeated"`
}

func (s CreateSharedCustomRuleRequestCondition) String() string {
  return tea.Prettify(s)
}

func (s CreateSharedCustomRuleRequestCondition) GoString() string {
  return s.String()
}

func (s *CreateSharedCustomRuleRequestCondition) SetMethodConditions(v []*CreateSharedCustomRuleRequestConditionMethodConditions) *CreateSharedCustomRuleRequestCondition {
  s.MethodConditions = v
  return s
}

func (s *CreateSharedCustomRuleRequestCondition) SetAreaConditions(v []*CreateSharedCustomRuleRequestConditionAreaConditions) *CreateSharedCustomRuleRequestCondition {
  s.AreaConditions = v
  return s
}

func (s *CreateSharedCustomRuleRequestCondition) SetIpOrIpsConditions(v []*CreateSharedCustomRuleRequestConditionIpOrIpsConditions) *CreateSharedCustomRuleRequestCondition {
  s.IpOrIpsConditions = v
  return s
}

func (s *CreateSharedCustomRuleRequestCondition) SetUriConditions(v []*CreateSharedCustomRuleRequestConditionUriConditions) *CreateSharedCustomRuleRequestCondition {
  s.UriConditions = v
  return s
}

func (s *CreateSharedCustomRuleRequestCondition) SetPathConditions(v []*CreateSharedCustomRuleRequestConditionPathConditions) *CreateSharedCustomRuleRequestCondition {
  s.PathConditions = v
  return s
}

func (s *CreateSharedCustomRuleRequestCondition) SetUaConditions(v []*CreateSharedCustomRuleRequestConditionUaConditions) *CreateSharedCustomRuleRequestCondition {
  s.UaConditions = v
  return s
}

func (s *CreateSharedCustomRuleRequestCondition) SetHeaderConditions(v []*CreateSharedCustomRuleRequestConditionHeaderConditions) *CreateSharedCustomRuleRequestCondition {
  s.HeaderConditions = v
  return s
}

func (s *CreateSharedCustomRuleRequestCondition) SetRefererConditions(v []*CreateSharedCustomRuleRequestConditionRefererConditions) *CreateSharedCustomRuleRequestCondition {
  s.RefererConditions = v
  return s
}

func (s *CreateSharedCustomRuleRequestCondition) SetJa3Conditions(v []*CreateSharedCustomRuleRequestConditionJa3Conditions) *CreateSharedCustomRuleRequestCondition {
  s.Ja3Conditions = v
  return s
}

func (s *CreateSharedCustomRuleRequestCondition) SetJa4Conditions(v []*CreateSharedCustomRuleRequestConditionJa4Conditions) *CreateSharedCustomRuleRequestCondition {
  s.Ja4Conditions = v
  return s
}

type CreateSharedCustomRuleRequestConditionMethodConditions struct     {
  // {"en":"Match type.\nEQUAL:Equals\nNOT_EQUAL:Does not equal","zh_CN":"匹配类型。\nEQUAL：等于\nNOT_EQUAL：不等于","exampleValue":"EQUAL,NOT_EQUAL"}
  MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
  // {"en":"Request method.\nSupported values: GET/POST/DELETE/PUT/HEAD/OPTIONS/COPY.","zh_CN":"请求方法。\n支持的值：GET/POST/DELETE/PUT/HEAD/OPTIONS/COPY。","exampleValue":"GET,POST,DELETE,PUT,HEAD,OPTIONS,COPY"}
  RequestMethod []*string `json:"requestMethod,omitempty" xml:"requestMethod,omitempty" require:"true" type:"Repeated"`
}

func (s CreateSharedCustomRuleRequestConditionMethodConditions) String() string {
  return tea.Prettify(s)
}

func (s CreateSharedCustomRuleRequestConditionMethodConditions) GoString() string {
  return s.String()
}

func (s *CreateSharedCustomRuleRequestConditionMethodConditions) SetMatchType(v string) *CreateSharedCustomRuleRequestConditionMethodConditions {
  s.MatchType = &v
  return s
}

func (s *CreateSharedCustomRuleRequestConditionMethodConditions) SetRequestMethod(v []*string) *CreateSharedCustomRuleRequestConditionMethodConditions {
  s.RequestMethod = v
  return s
}

type CreateSharedCustomRuleRequestConditionAreaConditions struct     {
  // {"en":"Match type.\nEQUAL:Equals\nNOT_EQUAL:Does not equal","zh_CN":"匹配类型。\nEQUAL：等于\nNOT_EQUAL：不等于","exampleValue":"EQUAL,NOT_EQUAL"}
  MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
  // {"dictionary":"belong=WAAP-MS-Ext|dict=waap_areaCityAndCountry","en":"Geo.","zh_CN":"区域。"}
  Areas []*string `json:"areas,omitempty" xml:"areas,omitempty" require:"true" type:"Repeated"`
}

func (s CreateSharedCustomRuleRequestConditionAreaConditions) String() string {
  return tea.Prettify(s)
}

func (s CreateSharedCustomRuleRequestConditionAreaConditions) GoString() string {
  return s.String()
}

func (s *CreateSharedCustomRuleRequestConditionAreaConditions) SetMatchType(v string) *CreateSharedCustomRuleRequestConditionAreaConditions {
  s.MatchType = &v
  return s
}

func (s *CreateSharedCustomRuleRequestConditionAreaConditions) SetAreas(v []*string) *CreateSharedCustomRuleRequestConditionAreaConditions {
  s.Areas = v
  return s
}

type CreateSharedCustomRuleRequestConditionIpOrIpsConditions struct     {
  // {"en":"Match type.\nEQUAL:Equals\nNOT_EQUAL:Does not equal","zh_CN":"匹配类型。\nEQUAL：等于\nNOT_EQUAL：不等于","exampleValue":"EQUAL,NOT_EQUAL"}
  MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
  // {"en":"IP/CIDR, maximum 500 IP/CIDR.","zh_CN":"IP/IP段，最多500个IP/IP段。"}
  IpOrIps []*string `json:"ipOrIps,omitempty" xml:"ipOrIps,omitempty" require:"true" type:"Repeated"`
}

func (s CreateSharedCustomRuleRequestConditionIpOrIpsConditions) String() string {
  return tea.Prettify(s)
}

func (s CreateSharedCustomRuleRequestConditionIpOrIpsConditions) GoString() string {
  return s.String()
}

func (s *CreateSharedCustomRuleRequestConditionIpOrIpsConditions) SetMatchType(v string) *CreateSharedCustomRuleRequestConditionIpOrIpsConditions {
  s.MatchType = &v
  return s
}

func (s *CreateSharedCustomRuleRequestConditionIpOrIpsConditions) SetIpOrIps(v []*string) *CreateSharedCustomRuleRequestConditionIpOrIpsConditions {
  s.IpOrIps = v
  return s
}

type CreateSharedCustomRuleRequestConditionUriConditions struct     {
  // {"en":"Match type.\nEQUAL: Equals, URI case sensitive\nNOT_EQUAL: Does not equal, URI case sensitive\nCONTAIN: Contains, URI case insensitive\nNOT_CONTAIN: Does not Contains, URI case insensitive\nREGEX: Regex match, URI case insensitive\nNOT_REGEX: Regular does not match, URI case insensitive\nSTART_WITH: Starts with, URI case insensitive\nEND_WITH: Ends with, URI case insensitive\nWILDCARD: Wildcard matches, URI case insensitive,* represents zero or more arbitrary characters, ? represents any single character\nNOT_WILDCARD: Wildcard does not match, URI case insensitive,* represents zero or more arbitrary characters, ? represents any single character","zh_CN":"匹配类型。\nEQUAL：等于，URI大小写敏感\nNOT_EQUAL：不等于，URI大小写敏感\nCONTAIN：包含，URI大小写不敏感\nNOT_CONTAIN：不包含，URI大小写不敏感\nREGEX：匹配正则，URI大小写不敏感\nNOT_REGEX：正则不匹配，URI大小写不敏感\nSTART_WITH：开头是，URI大小写不敏感\nEND_WITH：结尾是，URI大小写不敏感\nWILDCARD：通配符匹配，URI大小写不敏感，*代表零个或多个任意字符，?代表任意单个字符\nNOT_WILDCARD：通配符不匹配，URI大小写不敏感，*代表零个或多个任意字符，?代表任意单个字符","exampleValue":"EQUAL,NOT_EQUAL,CONTAIN,NOT_CONTAIN,REGEX,NOT_REGEX,START_WITH,END_WITH,WILDCARD,NOT_WILDCARD"}
  MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
  // {"en":"URI.\nWhen match type is EQUAL/NOT_EQUAL/START_WITH/END_WITH, uri needs to start with \"/\", and includes parameters.\nWhen the match type is REGEX/NOT_REGEX, only one value is allowed.\nExample: /test.html?id=1.","zh_CN":"URI。\n当匹配类型为等于/不等于/开头是/结尾是，URI必须以”/“开头，含参数。\n当匹配类型为正则/正则不匹配，则只允许只有一个值。\n示例：/test.html?id=1。"}
  Uri []*string `json:"uri,omitempty" xml:"uri,omitempty" require:"true" type:"Repeated"`
}

func (s CreateSharedCustomRuleRequestConditionUriConditions) String() string {
  return tea.Prettify(s)
}

func (s CreateSharedCustomRuleRequestConditionUriConditions) GoString() string {
  return s.String()
}

func (s *CreateSharedCustomRuleRequestConditionUriConditions) SetMatchType(v string) *CreateSharedCustomRuleRequestConditionUriConditions {
  s.MatchType = &v
  return s
}

func (s *CreateSharedCustomRuleRequestConditionUriConditions) SetUri(v []*string) *CreateSharedCustomRuleRequestConditionUriConditions {
  s.Uri = v
  return s
}

type CreateSharedCustomRuleRequestConditionPathConditions struct     {
  // {"en":"Match type.\nEQUAL: Equals, path case sensitive\nNOT_EQUAL: Does not equal, path case sensitive\nCONTAIN: Contains, path case insensitive\nNOT_CONTAIN: Does not Contains, path case insensitive\nREGEX: Regex match, path case insensitive\nNOT_REGEX: Regular does not match, path case sensitive\nSTART_WITH: Starts with, path case sensitive\nEND_WITH: Ends with, path case sensitive\nWILDCARD: Wildcard matches, path case sensitive,* represents zero or more arbitrary characters, ? represents any single character.\nNOT_WILDCARD: Wildcard does not match, path case sensitive,* represents zero or more arbitrary characters, ? represents any single character","zh_CN":"匹配类型。\nEQUAL：等于，路径大小写敏感\nNOT_EQUAL：不等于，路径大小写敏感\nCONTAIN：包含，路径大小写不敏感\nNOT_CONTAIN：不包含，路径大小写不敏感\nREGEX：匹配正则，路径大小写不敏感\nNOT_REGEX：正则不匹配，路径大小写不敏感\nSTART_WITH：开头是，路径大小写不敏感\nEND_WITH：结尾是，路径大小写不敏感\nWILDCARD：通配符匹配，路径大小写不敏感，*代表零个或多个任意字符，?代表任意单个字符\nNOT_WILDCARD：通配符不匹配，路径大小写不敏感，*代表零个或多个任意字符，?代表任意单个字符","exampleValue":"EQUAL,NOT_EQUAL,CONTAIN,NOT_CONTAIN,REGEX,NOT_REGEX,START_WITH,END_WITH,WILDCARD,NOT_WILDCARD"}
  MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
  // {"en":"Path.\nWhen match type is EQUAL/NOT_EQUAL/START_WITH/END_WITH, path needs to start with \"/\", and no parameters.\nWhen the match type is REGEX/NOT_REGEX, only one value is allowed.\nExample: /test.html.","zh_CN":"路径。\n当匹配类型为等于/不等于/开头是/结尾是，路径必须以“/”开头，不含参数。\n当匹配类型为正则/正则不匹配，则只允许只有一个值。\n示例：/test.html。"}
  CreateSharedCustomRulePaths []*string `json:"paths,omitempty" xml:"paths,omitempty" require:"true" type:"Repeated"`
}

func (s CreateSharedCustomRuleRequestConditionPathConditions) String() string {
  return tea.Prettify(s)
}

func (s CreateSharedCustomRuleRequestConditionPathConditions) GoString() string {
  return s.String()
}

func (s *CreateSharedCustomRuleRequestConditionPathConditions) SetMatchType(v string) *CreateSharedCustomRuleRequestConditionPathConditions {
  s.MatchType = &v
  return s
}

func (s *CreateSharedCustomRuleRequestConditionPathConditions) SetCreateSharedCustomRulePaths(v []*string) *CreateSharedCustomRuleRequestConditionPathConditions {
  s.CreateSharedCustomRulePaths = v
  return s
}

type CreateSharedCustomRuleRequestConditionUaConditions struct     {
  // {"en":"Match type.\nEQUAL: Equals, user agent case sensitive\nNOT_EQUAL: Does not equal, user agent case sensitive\nCONTAIN: Contains, user agent case insensitive\nNOT_CONTAIN: Does not Contains, user agent case insensitive\nNONE:Empty or non-existent\nREGEX: Regex match, user agent case insensitive\nNOT_REGEX: Wegular does not match, user agent case insensitive\nSTART_WITH: Starts with, user agent case insensitive\nEND_WITH: Ends with, user agent case insensitive\nWILDCARD: Wildcard matches, user agent case insensitive,* represents zero or more arbitrary characters, ? represents any single character\nNOT_WILDCARD: Wildcard does not match, user agent case insensitive,* represents zero or more arbitrary characters, ? represents any single character","zh_CN":"匹配类型。\nEQUAL：等于，User-Agent大小写敏感\nNOT_EQUAL：不等于，User-Agent大小写敏感\nCONTAIN：包含，User-Agent大小写不敏感\nNOT_CONTAIN：不包含，User-Agent大小写不敏感\nNONE：为空或不存在\nREGEX：匹配正则，User-Agent大小写不敏感\nNOT_REGEX：正则不匹配，User-Agent大小写不敏感\nSTART_WITH：开头是，User-Agent大小写不敏感\nEND_WITH：结尾是，User-Agent大小写不敏感\nWILDCARD：通配符匹配，User-Agent大小写不敏感，*代表零个或多个任意字符，?代表任意单个字符\nNOT_WILDCARD：通配符不匹配，User-Agent大小写不敏感，*代表零个或多个任意字符，?代表任意单个字符","exampleValue":"EQUAL,NOT_EQUAL,CONTAIN,NOT_CONTAIN,NONE,REGEX,NOT_REGEX,START_WITH,END_WITH,WILDCARD,NOT_WILDCARD"}
  MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
  // {"en":"User agent.\nWhen the match type is REGEX/NOT_REGEX, only one value is allowed.\nExample: go-Http-client/1.1.","zh_CN":"User-Agent。\n当匹配类型为正则/正则不匹配，则只允许只有一个值。\n示例：go-Http-client/1.1。"}
  Ua []*string `json:"ua,omitempty" xml:"ua,omitempty" require:"true" type:"Repeated"`
}

func (s CreateSharedCustomRuleRequestConditionUaConditions) String() string {
  return tea.Prettify(s)
}

func (s CreateSharedCustomRuleRequestConditionUaConditions) GoString() string {
  return s.String()
}

func (s *CreateSharedCustomRuleRequestConditionUaConditions) SetMatchType(v string) *CreateSharedCustomRuleRequestConditionUaConditions {
  s.MatchType = &v
  return s
}

func (s *CreateSharedCustomRuleRequestConditionUaConditions) SetUa(v []*string) *CreateSharedCustomRuleRequestConditionUaConditions {
  s.Ua = v
  return s
}

type CreateSharedCustomRuleRequestConditionHeaderConditions struct     {
  // {"en":"Match type.\nEQUAL: Equals, request header values case sensitive\nNOT_EQUAL: Does not equal, request header values case sensitive\nCONTAIN: Contains, request header values case insensitive\nNOT_CONTAIN: Does not Contains, request header values case insensitive\nNONE: Empty or non-existent\nREGEX: Regex match, request header values case insensitive\nNOT_REGEX: Regular does not match, request header values case insensitive\nSTART_WITH: Starts with, request header values case insensitive\nEND_WITH: Ends with, request header values case insensitive\nWILDCARD: Wildcard matches, request header values case insensitive,* represents zero or more arbitrary characters, ? represents any single character\nNOT_WILDCARD: Wildcard does not match, request header values case insensitive,* represents zero or more arbitrary characters, ? represents any single character","zh_CN":"匹配类型。\nEQUAL：等于，头部值大小写敏感\nNOT_EQUAL：不等于，头部值大小写敏感\nCONTAIN：包含，头部值大小写不敏感\nNOT_CONTAIN：不包含，头部值大小写不敏感\nNONE：为空或不存在\nREGEX：匹配正则，头部值大小写不敏感\nNOT_REGEX：正则不匹配，头部值大小写不敏感\nSTART_WITH：开头是，头部值大小写不敏感\nEND_WITH：结尾是，头部值大小写不敏感\nWILDCARD：通配符匹配，头部值大小写不敏感，*代表零个或多个任意字符，?代表任意单个字符\nNOT_WILDCARD：通配符不匹配，头部值大小写不敏感，*代表零个或多个任意字符，?代表任意单个字符","exampleValue":"EQUAL,NOT_EQUAL,CONTAIN,NOT_CONTAIN,NONE,REGEX,NOT_REGEX,START_WITH,END_WITH,WILDCARD,NOT_WILDCARD"}
  MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
  // {"en":"Header value.\nWhen the match type is REGEX/NOT_REGEX, only one value is allowed.","zh_CN":"头部值。\n当匹配类型为正则/正则不匹配，则只允许只有一个值。"}
  ValueList []*string `json:"valueList,omitempty" xml:"valueList,omitempty" require:"true" type:"Repeated"`
  // {"en":"Header name,case insensitive,up to 100 characters.\nExample: Accept.","zh_CN":"头部名称，大小写不敏感，最多100个字符。\n示例：Accept。"}
  Key *string `json:"key,omitempty" xml:"key,omitempty" require:"true"`
}

func (s CreateSharedCustomRuleRequestConditionHeaderConditions) String() string {
  return tea.Prettify(s)
}

func (s CreateSharedCustomRuleRequestConditionHeaderConditions) GoString() string {
  return s.String()
}

func (s *CreateSharedCustomRuleRequestConditionHeaderConditions) SetMatchType(v string) *CreateSharedCustomRuleRequestConditionHeaderConditions {
  s.MatchType = &v
  return s
}

func (s *CreateSharedCustomRuleRequestConditionHeaderConditions) SetValueList(v []*string) *CreateSharedCustomRuleRequestConditionHeaderConditions {
  s.ValueList = v
  return s
}

func (s *CreateSharedCustomRuleRequestConditionHeaderConditions) SetKey(v string) *CreateSharedCustomRuleRequestConditionHeaderConditions {
  s.Key = &v
  return s
}

type CreateSharedCustomRuleRequestConditionRefererConditions struct     {
  // {"en":"Referer.\nWhen the match type is REGEX/NOT_REGEX, only one value is allowed.\nExample: http://test.com.","zh_CN":"Referer。\n当匹配类型为正则/正则不匹配，则只允许只有一个值。\n示例：http://test.com。"}
  Referer []*string `json:"referer,omitempty" xml:"referer,omitempty" require:"true" type:"Repeated"`
  // {"en":"Match type.\nEQUAL: Equals, referer case sensitive\nNOT_EQUAL: Does not equal, referer case sensitive\nCONTAIN: Contains, referer case insensitive\nNOT_CONTAIN: Does not Contains, referer case insensitive\nNONE:Empty or non-existent\nREGEX: Regex match, referer case insensitive\nNOT_REGEX: Regular does not match, referer case insensitive\nSTART_WITH: Starts with, referer case insensitive\nEND_WITH: Ends with, referer case insensitive\nWILDCARD: Wildcard matches, referer case insensitive,* represents zero or more arbitrary characters, ? represents any single characte\nNOT_WILDCARD: Wildcard does not match, referer case insensitive,* represents zero or more arbitrary characters, ? represents any single character","zh_CN":"匹配类型。\nEQUAL：等于，referer大小写敏感\nNOT_EQUAL：不等于，referer大小写敏感\nCONTAIN：包含，referer大小写不敏感\nNOT_CONTAIN：不包含，referer大小写不敏感\nNONE：为空或不存在\nREGEX：匹配正则，referer大小写不敏感\nNOT_REGEX：正则不匹配，referer大小写不敏感\nSTART_WITH：开头是，referer大小写不敏感\nEND_WITH：结尾是，referer大小写不敏感\nWILDCARD：通配符匹配，referer大小写不敏感，*代表零个或多个任意字符，?代表任意单个字符\nNOT_WILDCARD：通配符不匹配，referer大小写不敏感，*代表零个或多个任意字符，?代表任意单个字符","exampleValue":"EQUAL,NOT_EQUAL,CONTAIN,NOT_CONTAIN,NONE,REGEX,NOT_REGEX,START_WITH,END_WITH,WILDCARD,NOT_WILDCARD"}
  MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
}

func (s CreateSharedCustomRuleRequestConditionRefererConditions) String() string {
  return tea.Prettify(s)
}

func (s CreateSharedCustomRuleRequestConditionRefererConditions) GoString() string {
  return s.String()
}

func (s *CreateSharedCustomRuleRequestConditionRefererConditions) SetReferer(v []*string) *CreateSharedCustomRuleRequestConditionRefererConditions {
  s.Referer = v
  return s
}

func (s *CreateSharedCustomRuleRequestConditionRefererConditions) SetMatchType(v string) *CreateSharedCustomRuleRequestConditionRefererConditions {
  s.MatchType = &v
  return s
}

type CreateSharedCustomRuleRequestConditionJa3Conditions struct     {
  // {"en":"Match type.\nEQUAL: Equals\nNOT_EQUAL: Does not equal","zh_CN":"匹配类型。\nEQUAL：等于 \nNOT_EQUAL：不等于","exampleValue":"EQUAL,NOT_EQUAL"}
  MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
  // {"en":"JA3 Fingerprint List, maximum 300 JA3 Fingerprint.\nWhen the match type is EQUAL/NOT_EQUAL, each item's character length must be 32 and can only include numbers and lowercase letters.","zh_CN":"JA3指纹列表，最多300个JA3指纹。\n当匹配类型为等于/不等于时，每项字符长度必须为32，且仅限数字和小写字母。"}
  Ja3List []*string `json:"ja3List,omitempty" xml:"ja3List,omitempty" require:"true" type:"Repeated"`
}

func (s CreateSharedCustomRuleRequestConditionJa3Conditions) String() string {
  return tea.Prettify(s)
}

func (s CreateSharedCustomRuleRequestConditionJa3Conditions) GoString() string {
  return s.String()
}

func (s *CreateSharedCustomRuleRequestConditionJa3Conditions) SetMatchType(v string) *CreateSharedCustomRuleRequestConditionJa3Conditions {
  s.MatchType = &v
  return s
}

func (s *CreateSharedCustomRuleRequestConditionJa3Conditions) SetJa3List(v []*string) *CreateSharedCustomRuleRequestConditionJa3Conditions {
  s.Ja3List = v
  return s
}

type CreateSharedCustomRuleRequestConditionJa4Conditions struct     {
  // {"en":"Match type. \nEQUAL: Equals\nNOT_EQUAL: Does not equal\nCONTAIN: Contains\nNOT_CONTAIN: Does not Contains\nSTART_WITH: Starts with\nEND_WITH: Ends with\nWILDCARD: Wildcard matches, ** represents zero or more arbitrary characters, ? represents any single character\nNOT_WILDCARD: Wildcard does not match, ** represents zero or more arbitrary characters, ? represents any single character","zh_CN":"匹配类型。\nEQUAL：等于\nNOT_EQUAL：不等于\nCONTAIN：包含\nNOT_CONTAIN：不包含\nSTART_WITH：开头是\nEND_WITH：结尾是\nWILDCARD：通配符匹配，*代表零个或多个任意字符，?代表任意单个字符\nNOT_WILDCARD：通配符不匹配，*代表零个或多个任意字符，?代表任意单个字符","exampleValue":"EQUAL,NOT_EQUAL,CONTAIN,NOT_CONTAIN,START_WITH,END_WITH,WILDCARD,NOT_WILDCARD"}
  MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
  // {"en":"JA4 Fingerprint List, maximum 300 JA4 Fingerprint.\nWhen the match type is EQUAL/NOT_EQUAL, each item's format must be 10 characters + 12 characters + 12 characters, separated by underscores, and can only include underscores, numbers, and lowercase letters.\nWhen the match type is CONTAIN/NOT_CONTAIN/START_WITH/END_WITH, each item is only allowed to include underscores, numbers, and lowercase letters.\nWhen the match type is WILDCARD/NOT_WILDCARD, each item, aside from  ** and ?, is only allowed to include underscores, numbers, and lowercase letters.","zh_CN":"JA4指纹列表，最多300个JA4指纹。\n当匹配类型为等于/不等于时，每项格式必须为10位字符+12位字符+12位字符，中间以“_”分隔，且仅限下划线、数字和小写字母。\n当匹配类型为包含/不包含/开头是/结尾是时，每项只允许输入下划线、数字和小写字母。\n当匹配类型为通配符匹配/通配符不匹配时，每项除*和?外，只允许输入下划线、数字和小写字母。"}
  Ja4List []*string `json:"ja4List,omitempty" xml:"ja4List,omitempty" require:"true" type:"Repeated"`
}

func (s CreateSharedCustomRuleRequestConditionJa4Conditions) String() string {
  return tea.Prettify(s)
}

func (s CreateSharedCustomRuleRequestConditionJa4Conditions) GoString() string {
  return s.String()
}

func (s *CreateSharedCustomRuleRequestConditionJa4Conditions) SetMatchType(v string) *CreateSharedCustomRuleRequestConditionJa4Conditions {
  s.MatchType = &v
  return s
}

func (s *CreateSharedCustomRuleRequestConditionJa4Conditions) SetJa4List(v []*string) *CreateSharedCustomRuleRequestConditionJa4Conditions {
  s.Ja4List = v
  return s
}

type CreateSharedCustomRuleRequestHeader struct {
  // {"en":"Security service type. Please enter a specific service type, if you purchase multiple security services.","zh_CN":"安全服务类型。有使用多个不同的安全服务时，需要填写具体的服务类型。"}
  ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty"`
}

func (s CreateSharedCustomRuleRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s CreateSharedCustomRuleRequestHeader) GoString() string {
  return s.String()
}

func (s *CreateSharedCustomRuleRequestHeader) SetServiceType(v string) *CreateSharedCustomRuleRequestHeader {
  s.ServiceType = &v
  return s
}

type CreateSharedCustomRulePaths struct {
}

func (s CreateSharedCustomRulePaths) String() string {
  return tea.Prettify(s)
}

func (s CreateSharedCustomRulePaths) GoString() string {
  return s.String()
}

type CreateSharedCustomRuleParameters struct {
}

func (s CreateSharedCustomRuleParameters) String() string {
  return tea.Prettify(s)
}

func (s CreateSharedCustomRuleParameters) GoString() string {
  return s.String()
}

type CreateSharedCustomRuleResponse struct {
  // {"en":"Description.","zh_CN":"描述信息。"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"dictionary":"belong=WAAP-MS-Ext|dict=waap_retCodeEnum","en":"Please refer to the error code for exceptions.","zh_CN":"请参照错误码。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Rule ID.","zh_CN":"规则ID。"}
  Data *string `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s CreateSharedCustomRuleResponse) String() string {
  return tea.Prettify(s)
}

func (s CreateSharedCustomRuleResponse) GoString() string {
  return s.String()
}

func (s *CreateSharedCustomRuleResponse) SetMsg(v string) *CreateSharedCustomRuleResponse {
  s.Msg = &v
  return s
}

func (s *CreateSharedCustomRuleResponse) SetCode(v string) *CreateSharedCustomRuleResponse {
  s.Code = &v
  return s
}

func (s *CreateSharedCustomRuleResponse) SetData(v string) *CreateSharedCustomRuleResponse {
  s.Data = &v
  return s
}

type CreateSharedCustomRuleResponseHeader struct {
}

func (s CreateSharedCustomRuleResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s CreateSharedCustomRuleResponseHeader) GoString() string {
  return s.String()
}




type CreateCustomActionRequest struct {
  // {"en":"Action name,maximum 50 characters.\nDoes not support # and &.","zh_CN":"动作名称，最多50个字符。\n不支持 # 和 &。"}
  ActionName *string `json:"actionName,omitempty" xml:"actionName,omitempty" require:"true"`
  // {"en":"Description,maximum 200 characters.","zh_CN":"描述，最多200个字符。"}
  Description *string `json:"description,omitempty" xml:"description,omitempty"`
  // {"en":"Status code.\nValid value range:200,204,206,400,401,403,404,500,501,503.","zh_CN":"状态码。\n有效值范围：200，204，206，400，401，403，404，500，501，503。"}
  StatusCode *string `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
  // {"en":"Response Content-Type type,multiple separated by ; sign.","zh_CN":"响应Content-Type类型，多个以 ; 隔开。"}
  ContentType *string `json:"contentType,omitempty" xml:"contentType,omitempty" require:"true"`
  // {"en":"Response content definition, the size of the response content cannot exceed 16KB, if you need to insert static resources such as pictures, it is recommended to use links to import.\nThe following interception information parameters are supported:\n{url} : Displays the URL information when intercepting\n{client_ip} : Displays the intercepted user IP information\n{time} : Displays the intercepted time point\n{event_id} : Displays the ID information of this interception event","zh_CN":"自定义响应内容定义，响应内容大小不能超过16KB，如需插入图片等静态资源，建议采用链接引入。\n支持引用下列拦截信息参数：\n{url} ：显示拦截时的URL信息\n{client_ip} ：显示被拦截的用户IP信息\n{time} ：显示被拦截的时间点\n{event_id} ：显示本次拦截事件的ID信息"}
  CusBlockCnt *string `json:"cusBlockCnt,omitempty" xml:"cusBlockCnt,omitempty" require:"true"`
}

func (s CreateCustomActionRequest) String() string {
  return tea.Prettify(s)
}

func (s CreateCustomActionRequest) GoString() string {
  return s.String()
}

func (s *CreateCustomActionRequest) SetActionName(v string) *CreateCustomActionRequest {
  s.ActionName = &v
  return s
}

func (s *CreateCustomActionRequest) SetDescription(v string) *CreateCustomActionRequest {
  s.Description = &v
  return s
}

func (s *CreateCustomActionRequest) SetStatusCode(v string) *CreateCustomActionRequest {
  s.StatusCode = &v
  return s
}

func (s *CreateCustomActionRequest) SetContentType(v string) *CreateCustomActionRequest {
  s.ContentType = &v
  return s
}

func (s *CreateCustomActionRequest) SetCusBlockCnt(v string) *CreateCustomActionRequest {
  s.CusBlockCnt = &v
  return s
}

type CreateCustomActionRequestHeader struct {
  // {"en":"Security service type. Please enter a specific service type, if you purchase multiple security services.","zh_CN":"安全服务类型。有使用多个不同的安全服务时，需要填写具体的服务类型。"}
  ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty"`
}

func (s CreateCustomActionRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s CreateCustomActionRequestHeader) GoString() string {
  return s.String()
}

func (s *CreateCustomActionRequestHeader) SetServiceType(v string) *CreateCustomActionRequestHeader {
  s.ServiceType = &v
  return s
}

type CreateCustomActionPaths struct {
}

func (s CreateCustomActionPaths) String() string {
  return tea.Prettify(s)
}

func (s CreateCustomActionPaths) GoString() string {
  return s.String()
}

type CreateCustomActionParameters struct {
}

func (s CreateCustomActionParameters) String() string {
  return tea.Prettify(s)
}

func (s CreateCustomActionParameters) GoString() string {
  return s.String()
}

type CreateCustomActionResponse struct {
  // {"dictionary":"belong=WAAP-MS-Ext|dict=waap_retCodeEnum","en":"Please refer to the error code for exceptions.","zh_CN":"请参照错误码。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Description.","zh_CN":"描述信息。"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"Customize response ID.","zh_CN":"自定义响应ID。"}
  Data *string `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s CreateCustomActionResponse) String() string {
  return tea.Prettify(s)
}

func (s CreateCustomActionResponse) GoString() string {
  return s.String()
}

func (s *CreateCustomActionResponse) SetCode(v string) *CreateCustomActionResponse {
  s.Code = &v
  return s
}

func (s *CreateCustomActionResponse) SetMsg(v string) *CreateCustomActionResponse {
  s.Msg = &v
  return s
}

func (s *CreateCustomActionResponse) SetData(v string) *CreateCustomActionResponse {
  s.Data = &v
  return s
}

type CreateCustomActionResponseHeader struct {
}

func (s CreateCustomActionResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s CreateCustomActionResponseHeader) GoString() string {
  return s.String()
}




type UpdateSharedWhitelistRuleRequest struct {
  // {"en":"Rule ID.","zh_CN":"规则ID。"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {"en":"Rule name, maximum 100 characters.Does not support # and &.","zh_CN":"规则名称，最多100个字符。不支持 # 和 &。"}
  RuleName *string `json:"ruleName,omitempty" xml:"ruleName,omitempty"`
  // {"en":"Description, maximum 1000 characters.","zh_CN":"描述，最多1000个字符。"}
  Description *string `json:"description,omitempty" xml:"description,omitempty"`
  // {"en":"Match conditions, at least one, at most five.","zh_CN":"匹配条件，至少一个，至多五个。"}
  Conditions *UpdateSharedWhitelistRuleRequestConditions `json:"conditions,omitempty" xml:"conditions,omitempty" type:"Struct"`
}

func (s UpdateSharedWhitelistRuleRequest) String() string {
  return tea.Prettify(s)
}

func (s UpdateSharedWhitelistRuleRequest) GoString() string {
  return s.String()
}

func (s *UpdateSharedWhitelistRuleRequest) SetId(v string) *UpdateSharedWhitelistRuleRequest {
  s.Id = &v
  return s
}

func (s *UpdateSharedWhitelistRuleRequest) SetRuleName(v string) *UpdateSharedWhitelistRuleRequest {
  s.RuleName = &v
  return s
}

func (s *UpdateSharedWhitelistRuleRequest) SetDescription(v string) *UpdateSharedWhitelistRuleRequest {
  s.Description = &v
  return s
}

func (s *UpdateSharedWhitelistRuleRequest) SetConditions(v *UpdateSharedWhitelistRuleRequestConditions) *UpdateSharedWhitelistRuleRequest {
  s.Conditions = v
  return s
}

type UpdateSharedWhitelistRuleRequestConditions struct {
  // {"en":"IP/CIDR match conditions, match type cannot be repeated.","zh_CN":"IP/IP段匹配条件，匹配类型不可重复。"}
  IpOrIpsConditions []*UpdateSharedWhitelistRuleRequestConditionsIpOrIpsConditions `json:"ipOrIpsConditions,omitempty" xml:"ipOrIpsConditions,omitempty" type:"Repeated"`
  // {"en":"Path match conditions, match type cannot be repeated.","zh_CN":"路径匹配条件，匹配类型不可重复。"}
  PathConditions []*UpdateSharedWhitelistRuleRequestConditionsPathConditions `json:"pathConditions,omitempty" xml:"pathConditions,omitempty" type:"Repeated"`
  // {"en":"URI match conditions, match type cannot be repeated.","zh_CN":"URI匹配条件，匹配类型不可重复。"}
  UriConditions []*UpdateSharedWhitelistRuleRequestConditionsUriConditions `json:"uriConditions,omitempty" xml:"uriConditions,omitempty" type:"Repeated"`
  // {"en":"User agent match conditions, match type cannot be repeated.","zh_CN":"User-Agent 匹配条件，匹配类型不可重复。"}
  UaConditions []*UpdateSharedWhitelistRuleRequestConditionsUaConditions `json:"uaConditions,omitempty" xml:"uaConditions,omitempty" type:"Repeated"`
  // {"en":"Referer match conditions, match type cannot be repeated.","zh_CN":"Referer 匹配条件，匹配类型不可重复。"}
  RefererConditions []*UpdateSharedWhitelistRuleRequestConditionsRefererConditions `json:"refererConditions,omitempty" xml:"refererConditions,omitempty" type:"Repeated"`
  // {"en":"Request header match conditions.","zh_CN":"请求头匹配条件。"}
  HeaderConditions []*UpdateSharedWhitelistRuleRequestConditionsHeaderConditions `json:"headerConditions,omitempty" xml:"headerConditions,omitempty" type:"Repeated"`
}

func (s UpdateSharedWhitelistRuleRequestConditions) String() string {
  return tea.Prettify(s)
}

func (s UpdateSharedWhitelistRuleRequestConditions) GoString() string {
  return s.String()
}

func (s *UpdateSharedWhitelistRuleRequestConditions) SetIpOrIpsConditions(v []*UpdateSharedWhitelistRuleRequestConditionsIpOrIpsConditions) *UpdateSharedWhitelistRuleRequestConditions {
  s.IpOrIpsConditions = v
  return s
}

func (s *UpdateSharedWhitelistRuleRequestConditions) SetPathConditions(v []*UpdateSharedWhitelistRuleRequestConditionsPathConditions) *UpdateSharedWhitelistRuleRequestConditions {
  s.PathConditions = v
  return s
}

func (s *UpdateSharedWhitelistRuleRequestConditions) SetUriConditions(v []*UpdateSharedWhitelistRuleRequestConditionsUriConditions) *UpdateSharedWhitelistRuleRequestConditions {
  s.UriConditions = v
  return s
}

func (s *UpdateSharedWhitelistRuleRequestConditions) SetUaConditions(v []*UpdateSharedWhitelistRuleRequestConditionsUaConditions) *UpdateSharedWhitelistRuleRequestConditions {
  s.UaConditions = v
  return s
}

func (s *UpdateSharedWhitelistRuleRequestConditions) SetRefererConditions(v []*UpdateSharedWhitelistRuleRequestConditionsRefererConditions) *UpdateSharedWhitelistRuleRequestConditions {
  s.RefererConditions = v
  return s
}

func (s *UpdateSharedWhitelistRuleRequestConditions) SetHeaderConditions(v []*UpdateSharedWhitelistRuleRequestConditionsHeaderConditions) *UpdateSharedWhitelistRuleRequestConditions {
  s.HeaderConditions = v
  return s
}

type UpdateSharedWhitelistRuleRequestConditionsIpOrIpsConditions struct     {
  // {"en":"Match type.\nEQUAL: Equals\nNOT_EQUAL: Does not equal","zh_CN":"匹配类型。\nEQUAL：等于\nNOT_EQUAL：不等于","exampleValue":"EQUAL,NOT_EQUAL"}
  MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
  // {"en":"IP/CIDR, maximum 500 IP/CIDR.","zh_CN":"IP/IP段，最多500个IP/IP段。"}
  IpOrIps []*string `json:"ipOrIps,omitempty" xml:"ipOrIps,omitempty" require:"true" type:"Repeated"`
}

func (s UpdateSharedWhitelistRuleRequestConditionsIpOrIpsConditions) String() string {
  return tea.Prettify(s)
}

func (s UpdateSharedWhitelistRuleRequestConditionsIpOrIpsConditions) GoString() string {
  return s.String()
}

func (s *UpdateSharedWhitelistRuleRequestConditionsIpOrIpsConditions) SetMatchType(v string) *UpdateSharedWhitelistRuleRequestConditionsIpOrIpsConditions {
  s.MatchType = &v
  return s
}

func (s *UpdateSharedWhitelistRuleRequestConditionsIpOrIpsConditions) SetIpOrIps(v []*string) *UpdateSharedWhitelistRuleRequestConditionsIpOrIpsConditions {
  s.IpOrIps = v
  return s
}

type UpdateSharedWhitelistRuleRequestConditionsPathConditions struct     {
  // {"en":"Match type.\nEQUAL: Equals, path case sensitive\nNOT_EQUAL: Does not equal, path case sensitive\nCONTAIN: Contains, path case insensitive\nNOT_CONTAIN: Does not Contains, path case insensitive\nREGEX: Regex match, path case insensitive\nNOT_REGEX: Regular does not match, path case sensitive\nSTART_WITH: Starts with, path case sensitive\nEND_WITH: Ends with, path case sensitive\nWILDCARD: Wildcard matches, path case sensitive,* represents zero or more arbitrary characters, ? represents any single character.\nNOT_WILDCARD: Wildcard does not match, path case sensitive,* represents zero or more arbitrary characters, ? represents any single character","zh_CN":"匹配类型。\nEQUAL：等于，路径大小写敏感\nNOT_EQUAL：不等于，路径大小写敏感\nCONTAIN：包含，路径大小写不敏感\nNOT_CONTAIN：不包含，路径大小写不敏感\nREGEX：匹配正则，路径大小写不敏感\nNOT_REGEX：正则不匹配，路径大小写不敏感\nSTART_WITH：开头是，路径大小写不敏感\nEND_WITH：结尾是，路径大小写不敏感\nWILDCARD：通配符匹配，路径大小写不敏感，*代表零个或多个任意字符，?代表任意单个字符\nNOT_WILDCARD：通配符不匹配，路径大小写不敏感，*代表零个或多个任意字符，?代表任意单个字符","exampleValue":"EQUAL,NOT_EQUAL,CONTAIN,NOT_CONTAIN,REGEX,NOT_REGEX,START_WITH,END_WITH,WILDCARD,NOT_WILDCARD"}
  MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
  // {"en":"Path.\nWhen match type is EQUAL/NOT_EQUAL/START_WITH/END_WITH, path needs to start with \"/\", and no parameters.\nWhen the match type is REGEX/NOT_REGEX, only one value is allowed.\nExample: /test.html.","zh_CN":"路径。\n当匹配类型为等于/不等于/开头是/结尾是，路径必须以“/”开头，不含参数。\n当匹配类型为正则/正则不匹配，则只允许只有一个值。\n示例：/test.html。"}
  UpdateSharedWhitelistRulePaths []*string `json:"paths,omitempty" xml:"paths,omitempty" require:"true" type:"Repeated"`
}

func (s UpdateSharedWhitelistRuleRequestConditionsPathConditions) String() string {
  return tea.Prettify(s)
}

func (s UpdateSharedWhitelistRuleRequestConditionsPathConditions) GoString() string {
  return s.String()
}

func (s *UpdateSharedWhitelistRuleRequestConditionsPathConditions) SetMatchType(v string) *UpdateSharedWhitelistRuleRequestConditionsPathConditions {
  s.MatchType = &v
  return s
}

func (s *UpdateSharedWhitelistRuleRequestConditionsPathConditions) SetUpdateSharedWhitelistRulePaths(v []*string) *UpdateSharedWhitelistRuleRequestConditionsPathConditions {
  s.UpdateSharedWhitelistRulePaths = v
  return s
}

type UpdateSharedWhitelistRuleRequestConditionsUriConditions struct     {
  // {"en":"Match type.\nEQUAL: Equals, URI case sensitive\nNOT_EQUAL: Does not equal, URI case sensitive\nCONTAIN: Contains, URI case insensitive\nNOT_CONTAIN: Does not Contains, URI case insensitive\nREGEX: Regex match, URI case insensitive\nNOT_REGEX: Regular does not match, URI case insensitive\nSTART_WITH: Starts with, URI case insensitive\nEND_WITH: Ends with, URI case insensitive\nWILDCARD: Wildcard matches, URI case insensitive,* represents zero or more arbitrary characters, ? represents any single character\nNOT_WILDCARD: Wildcard does not match, URI case insensitive,* represents zero or more arbitrary characters, ? represents any single character","zh_CN":"匹配类型。\nEQUAL：等于，URI大小写敏感\nNOT_EQUAL：不等于，URI大小写敏感\nCONTAIN：包含，URI大小写不敏感\nNOT_CONTAIN：不包含，URI大小写不敏感\nREGEX：匹配正则，URI大小写不敏感\nNOT_REGEX：正则不匹配，URI大小写不敏感\nSTART_WITH：开头是，URI大小写不敏感\nEND_WITH：结尾是，URI大小写不敏感\nWILDCARD：通配符匹配，URI大小写不敏感，*代表零个或多个任意字符，?代表任意单个字符\nNOT_WILDCARD：通配符不匹配，URI大小写不敏感，*代表零个或多个任意字符，?代表任意单个字符","exampleValue":"EQUAL,NOT_EQUAL,CONTAIN,NOT_CONTAIN,REGEX,NOT_REGEX,START_WITH,END_WITH,WILDCARD,NOT_WILDCARD"}
  MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
  // {"en":"URI.\nWhen match type is EQUAL/NOT_EQUAL/START_WITH/END_WITH, uri needs to start with \"/\", and includes parameters.\nWhen the match type is REGEX/NOT_REGEX, only one value is allowed.\nExample: /test.html?id=1.","zh_CN":"URI。\n当匹配类型为等于/不等于/开头是/结尾是，URI必须以”/“开头，含参数。\n当匹配类型为正则/正则不匹配，则只允许只有一个值。\n示例：/test.html?id=1。"}
  Uri []*string `json:"uri,omitempty" xml:"uri,omitempty" require:"true" type:"Repeated"`
}

func (s UpdateSharedWhitelistRuleRequestConditionsUriConditions) String() string {
  return tea.Prettify(s)
}

func (s UpdateSharedWhitelistRuleRequestConditionsUriConditions) GoString() string {
  return s.String()
}

func (s *UpdateSharedWhitelistRuleRequestConditionsUriConditions) SetMatchType(v string) *UpdateSharedWhitelistRuleRequestConditionsUriConditions {
  s.MatchType = &v
  return s
}

func (s *UpdateSharedWhitelistRuleRequestConditionsUriConditions) SetUri(v []*string) *UpdateSharedWhitelistRuleRequestConditionsUriConditions {
  s.Uri = v
  return s
}

type UpdateSharedWhitelistRuleRequestConditionsUaConditions struct     {
  // {"en":"Match type.\nEQUAL: Equals, user agent case sensitive\nNOT_EQUAL: Does not equal, user agent case sensitive\nCONTAIN: Contains, user agent case insensitive\nNOT_CONTAIN: Does not Contains, user agent case insensitive\nREGEX: Regex match, user agent case insensitive\nNOT_REGEX: Regular does not match, user agent case insensitive\nSTART_WITH: Starts with, user agent case insensitive\nEND_WITH: Ends with, user agent case insensitive\nWILDCARD: Wildcard matches, user agent case insensitive,* represents zero or more arbitrary characters, ? represents any single character\nNOT_WILDCARD: Wildcard does not match, user agent case insensitive,* represents zero or more arbitrary characters, ? represents any single character","zh_CN":"匹配类型。\nEQUAL：等于，User-Agent大小写敏感\nNOT_EQUAL：不等于，User-Agent大小写敏感\nCONTAIN：包含，User-Agent大小写不敏感\nNOT_CONTAIN：不包含，User-Agent大小写不敏感\nREGEX：匹配正则，User-Agent大小写不敏感\nNOT_REGEX：正则不匹配，User-Agent大小写不敏感\nSTART_WITH：开头是，User-Agent大小写不敏感\nEND_WITH：结尾是，User-Agent大小写不敏感\nWILDCARD：通配符匹配，User-Agent大小写不敏感，*代表零个或多个任意字符，?代表任意单个字符\nNOT_WILDCARD：通配符不匹配，User-Agent大小写不敏感，*代表零个或多个任意字符，?代表任意单个字符","exampleValue":"EQUAL,NOT_EQUAL,CONTAIN,NOT_CONTAIN,REGEX,NOT_REGEX,START_WITH,END_WITH,WILDCARD,NOT_WILDCARD"}
  MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
  // {"en":"User agent.\nWhen the match type is REGEX/NOT_REGEX, only one value is allowed.\nExample: go-Http-client/1.1.","zh_CN":"User-Agent。\n当匹配类型为正则/正则不匹配，则只允许只有一个值。\n示例：go-Http-client/1.1。"}
  Ua []*string `json:"ua,omitempty" xml:"ua,omitempty" require:"true" type:"Repeated"`
}

func (s UpdateSharedWhitelistRuleRequestConditionsUaConditions) String() string {
  return tea.Prettify(s)
}

func (s UpdateSharedWhitelistRuleRequestConditionsUaConditions) GoString() string {
  return s.String()
}

func (s *UpdateSharedWhitelistRuleRequestConditionsUaConditions) SetMatchType(v string) *UpdateSharedWhitelistRuleRequestConditionsUaConditions {
  s.MatchType = &v
  return s
}

func (s *UpdateSharedWhitelistRuleRequestConditionsUaConditions) SetUa(v []*string) *UpdateSharedWhitelistRuleRequestConditionsUaConditions {
  s.Ua = v
  return s
}

type UpdateSharedWhitelistRuleRequestConditionsRefererConditions struct     {
  // {"en":"Match type.\nEQUAL: Equals, referer case sensitive\nNOT_EQUAL: Does not equal, referer case sensitive\nCONTAIN: Contains, referer case insensitive\nNOT_CONTAIN: Does not Contains, referer case insensitive\nREGEX: Regex match, referer case insensitive\nNOT_REGEX: Regular does not match, referer case insensitive\nSTART_WITH: Starts with, referer case insensitive\nEND_WITH: Ends with, referer case insensitive\nWILDCARD: Wildcard matches, referer case insensitive,* represents zero or more arbitrary characters, ? represents any single characte\nNOT_WILDCARD: Wildcard does not match, referer case insensitive,* represents zero or more arbitrary characters, ? represents any single character","zh_CN":"匹配类型。\nEQUAL：等于，referer大小写敏感\nNOT_EQUAL：不等于，referer大小写敏感\nCONTAIN：包含，referer大小写不敏感\nNOT_CONTAIN：不包含，referer大小写不敏感\nREGEX：匹配正则，referer大小写不敏感\nNOT_REGEX：正则不匹配，referer大小写不敏感\nSTART_WITH：开头是，referer大小写不敏感\nEND_WITH：结尾是，referer大小写不敏感\nWILDCARD：通配符匹配，referer大小写不敏感，*代表零个或多个任意字符，?代表任意单个字符\nNOT_WILDCARD：通配符不匹配，referer大小写不敏感，*代表零个或多个任意字符，?代表任意单个字符","exampleValue":"EQUAL,NOT_EQUAL,CONTAIN,NOT_CONTAIN,REGEX,NOT_REGEX,START_WITH,END_WITH,WILDCARD,NOT_WILDCARD"}
  MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
  // {"en":"Referer.\nWhen the match type is REGEX/NOT_REGEX, only one value is allowed.\nExample: http://test.com.","zh_CN":"Referer。\n当匹配类型为正则/正则不匹配，则只允许只有一个值。\n示例：http://test.com。"}
  Referer []*string `json:"referer,omitempty" xml:"referer,omitempty" require:"true" type:"Repeated"`
}

func (s UpdateSharedWhitelistRuleRequestConditionsRefererConditions) String() string {
  return tea.Prettify(s)
}

func (s UpdateSharedWhitelistRuleRequestConditionsRefererConditions) GoString() string {
  return s.String()
}

func (s *UpdateSharedWhitelistRuleRequestConditionsRefererConditions) SetMatchType(v string) *UpdateSharedWhitelistRuleRequestConditionsRefererConditions {
  s.MatchType = &v
  return s
}

func (s *UpdateSharedWhitelistRuleRequestConditionsRefererConditions) SetReferer(v []*string) *UpdateSharedWhitelistRuleRequestConditionsRefererConditions {
  s.Referer = v
  return s
}

type UpdateSharedWhitelistRuleRequestConditionsHeaderConditions struct     {
  // {"en":"Match type.\nEQUAL: Equals, request header values case sensitive\nNOT_EQUAL: Does not equal, request header values case sensitive\nCONTAIN: Contains, request header values case insensitive\nNOT_CONTAIN: Does not Contains, request header values case insensitive\nREGEX: Regex match, request header values case insensitive\nNOT_REGEX: Regular does not match, request header values case insensitive\nSTART_WITH: Starts with, request header values case insensitive\nEND_WITH: Ends with, request header values case insensitive\nWILDCARD: Wildcard matches, request header values case insensitive,* represents zero or more arbitrary characters, ? represents any single character\nNOT_WILDCARD: Wildcard does not match, request header values case insensitive,* represents zero or more arbitrary characters, ? represents any single character","zh_CN":"匹配类型。\nEQUAL：等于，头部值大小写敏感\nNOT_EQUAL：不等于，头部值大小写敏感\nCONTAIN：包含，头部值大小写不敏感\nNOT_CONTAIN：不包含，头部值大小写不敏感\nREGEX：匹配正则，头部值大小写不敏感\nNOT_REGEX：正则不匹配，头部值大小写不敏感\nSTART_WITH：开头是，头部值大小写不敏感\nEND_WITH：结尾是，头部值大小写不敏感\nWILDCARD：通配符匹配，头部值大小写不敏感，*代表零个或多个任意字符，?代表任意单个字符\nNOT_WILDCARD：通配符不匹配，头部值大小写不敏感，*代表零个或多个任意字符，?代表任意单个字符","exampleValue":"EQUAL,NOT_EQUAL,CONTAIN,NOT_CONTAIN,REGEX,NOT_REGEX,START_WITH,END_WITH,WILDCARD,NOT_WILDCARD"}
  MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
  // {"en":"Header name,case insensitive,up to 100 characters.\nExample: Accept.","zh_CN":"头部名称，大小写不敏感，最多100个字符。\n示例：Accept。"}
  Key *string `json:"key,omitempty" xml:"key,omitempty" require:"true"`
  // {"en":"Header value.\nWhen the match type is REGEX/NOT_REGEX, only one value is allowed.","zh_CN":"头部值。\n当匹配类型为正则/正则不匹配，则只允许只有一个值。"}
  ValueList []*string `json:"valueList,omitempty" xml:"valueList,omitempty" require:"true" type:"Repeated"`
}

func (s UpdateSharedWhitelistRuleRequestConditionsHeaderConditions) String() string {
  return tea.Prettify(s)
}

func (s UpdateSharedWhitelistRuleRequestConditionsHeaderConditions) GoString() string {
  return s.String()
}

func (s *UpdateSharedWhitelistRuleRequestConditionsHeaderConditions) SetMatchType(v string) *UpdateSharedWhitelistRuleRequestConditionsHeaderConditions {
  s.MatchType = &v
  return s
}

func (s *UpdateSharedWhitelistRuleRequestConditionsHeaderConditions) SetKey(v string) *UpdateSharedWhitelistRuleRequestConditionsHeaderConditions {
  s.Key = &v
  return s
}

func (s *UpdateSharedWhitelistRuleRequestConditionsHeaderConditions) SetValueList(v []*string) *UpdateSharedWhitelistRuleRequestConditionsHeaderConditions {
  s.ValueList = v
  return s
}

type UpdateSharedWhitelistRuleRequestHeader struct {
  // {"en":"Security service type. Please enter a specific service type, if you purchase multiple security services.","zh_CN":"安全服务类型。有使用多个不同的安全服务时，需要填写具体的服务类型。"}
  ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty"`
}

func (s UpdateSharedWhitelistRuleRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s UpdateSharedWhitelistRuleRequestHeader) GoString() string {
  return s.String()
}

func (s *UpdateSharedWhitelistRuleRequestHeader) SetServiceType(v string) *UpdateSharedWhitelistRuleRequestHeader {
  s.ServiceType = &v
  return s
}

type UpdateSharedWhitelistRulePaths struct {
}

func (s UpdateSharedWhitelistRulePaths) String() string {
  return tea.Prettify(s)
}

func (s UpdateSharedWhitelistRulePaths) GoString() string {
  return s.String()
}

type UpdateSharedWhitelistRuleParameters struct {
}

func (s UpdateSharedWhitelistRuleParameters) String() string {
  return tea.Prettify(s)
}

func (s UpdateSharedWhitelistRuleParameters) GoString() string {
  return s.String()
}

type UpdateSharedWhitelistRuleResponse struct {
  // {"dictionary":"belong=WAAP-MS-Ext|dict=waap_retCodeEnum","en":"Please refer to the error code for exceptions.","zh_CN":"请参照错误码。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Description.","zh_CN":"描述信息。"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
}

func (s UpdateSharedWhitelistRuleResponse) String() string {
  return tea.Prettify(s)
}

func (s UpdateSharedWhitelistRuleResponse) GoString() string {
  return s.String()
}

func (s *UpdateSharedWhitelistRuleResponse) SetCode(v string) *UpdateSharedWhitelistRuleResponse {
  s.Code = &v
  return s
}

func (s *UpdateSharedWhitelistRuleResponse) SetMsg(v string) *UpdateSharedWhitelistRuleResponse {
  s.Msg = &v
  return s
}

type UpdateSharedWhitelistRuleResponseHeader struct {
}

func (s UpdateSharedWhitelistRuleResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s UpdateSharedWhitelistRuleResponseHeader) GoString() string {
  return s.String()
}




type QueryAppApiExceptionFeatureDetailRequest struct {
  // {"en":"Rule Id","zh_CN":"规则 ID"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
}

func (s QueryAppApiExceptionFeatureDetailRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryAppApiExceptionFeatureDetailRequest) GoString() string {
  return s.String()
}

func (s *QueryAppApiExceptionFeatureDetailRequest) SetId(v string) *QueryAppApiExceptionFeatureDetailRequest {
  s.Id = &v
  return s
}

type QueryAppApiExceptionFeatureDetailRequestHeader struct {
  // {"en":"Security service type. Please enter a specific service type, if you purchase multiple security services.","zh_CN":"安全服务类型。有使用多个不同的安全服务时，需要填写具体的服务类型。"}
  ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty"`
}

func (s QueryAppApiExceptionFeatureDetailRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryAppApiExceptionFeatureDetailRequestHeader) GoString() string {
  return s.String()
}

func (s *QueryAppApiExceptionFeatureDetailRequestHeader) SetServiceType(v string) *QueryAppApiExceptionFeatureDetailRequestHeader {
  s.ServiceType = &v
  return s
}

type QueryAppApiExceptionFeatureDetailPaths struct {
}

func (s QueryAppApiExceptionFeatureDetailPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryAppApiExceptionFeatureDetailPaths) GoString() string {
  return s.String()
}

type QueryAppApiExceptionFeatureDetailParameters struct {
}

func (s QueryAppApiExceptionFeatureDetailParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryAppApiExceptionFeatureDetailParameters) GoString() string {
  return s.String()
}

type QueryAppApiExceptionFeatureDetailResponse struct {
  // {"dictionary":"belong=WAAP-MS-Ext|dict=waap_retCodeEnum","en":"Please refer to the error code for exceptions.","zh_CN":"请参照错误码。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Description.","zh_CN":"描述信息。"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"Data.","zh_CN":"出参数据。"}
  Data *QueryAppApiExceptionFeatureDetailResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
}

func (s QueryAppApiExceptionFeatureDetailResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryAppApiExceptionFeatureDetailResponse) GoString() string {
  return s.String()
}

func (s *QueryAppApiExceptionFeatureDetailResponse) SetCode(v string) *QueryAppApiExceptionFeatureDetailResponse {
  s.Code = &v
  return s
}

func (s *QueryAppApiExceptionFeatureDetailResponse) SetMsg(v string) *QueryAppApiExceptionFeatureDetailResponse {
  s.Msg = &v
  return s
}

func (s *QueryAppApiExceptionFeatureDetailResponse) SetData(v *QueryAppApiExceptionFeatureDetailResponseData) *QueryAppApiExceptionFeatureDetailResponse {
  s.Data = v
  return s
}

type QueryAppApiExceptionFeatureDetailResponseData struct {
  // {"en":"Rule name, duplicate not allowed.","zh_CN":"规则名称，不允许重复"}
  RuleName *string `json:"ruleName,omitempty" xml:"ruleName,omitempty" require:"true"`
  // {"en":"Characteristic.","zh_CN":"特征。"}
  Config *QueryAppApiExceptionFeatureDetailResponseDataConfig `json:"config,omitempty" xml:"config,omitempty" require:"true" type:"Struct"`
  // {"en":"Type enumeration, for example:\nNATIVE_ APP        remarks: Native App\nCallback_ API       remarks: Callback API\nOther_ API         remarks: Other program API\nHYBRID_ APP     remarks: Hybrid APP","zh_CN":"类型枚举，例如：\nNATIVE_APP          备注：原生APP\nCALLBACK_API      备注：回调API\nOTHER_API            备注：其他程序API\nHYBRID_APP          备注：混合APP","exampleValue":"NATIVE_APP,CALLBACK_API,OTHER_API,HYBRID_APP"}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
}

func (s QueryAppApiExceptionFeatureDetailResponseData) String() string {
  return tea.Prettify(s)
}

func (s QueryAppApiExceptionFeatureDetailResponseData) GoString() string {
  return s.String()
}

func (s *QueryAppApiExceptionFeatureDetailResponseData) SetRuleName(v string) *QueryAppApiExceptionFeatureDetailResponseData {
  s.RuleName = &v
  return s
}

func (s *QueryAppApiExceptionFeatureDetailResponseData) SetConfig(v *QueryAppApiExceptionFeatureDetailResponseDataConfig) *QueryAppApiExceptionFeatureDetailResponseData {
  s.Config = v
  return s
}

func (s *QueryAppApiExceptionFeatureDetailResponseData) SetType(v string) *QueryAppApiExceptionFeatureDetailResponseData {
  s.Type = &v
  return s
}

type QueryAppApiExceptionFeatureDetailResponseDataConfig struct {
  // {"en":"Matching condition name.\nPATH         Path\nURI            URI\nUA             User-Agent\nREFERER    Referer\nHEADER    Request Header","zh_CN":"匹配条件名称。\nPATH                       路径\nURI                          URI\nUA                           User-Agent\nREFERER                  Referer\nHEAD                      Request Header","exampleValue":"PATH,URI,UA,REFERER,HEAD"}
  Condition *string `json:"condition,omitempty" xml:"condition,omitempty" require:"true"`
  // {"en":"Matching condition function.\nCONTAIN                  contains\nNOT_CONTAIN         does not contain\nEQUAL                      equals\nNOT_EQUAL             does not equal\nEMPTY                      does not exist or has no value(When the condition is a path or URI, this value is not optional)\nREGEX                      regex match","zh_CN":"匹配条件函数。\nCONTAIN                  包含\nNOT_CONTAIN         不包含\nEQUAL                      等于\nNOT_EQUAL             不等于\nEMPTY                      为空或不存在（当condition为路径或者URI时，该值不可选）\nREGEX                       匹配正则","exampleValue":"CONTAIN,NOT_CONTAIN,EQUAL,NOT_EQUAL,EMPTY,REGEX"}
  MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
  // {"en":"Condition value list, When matchType is EMPTY, the value is empty;  When matchType is REGEX, only one item is set","zh_CN":"条件值列表，当matchType为EMPTY时，该值为空；当matchType为REGEX时，只能传一条"}
  ConditionValue []*string `json:"conditionValue,omitempty" xml:"conditionValue,omitempty" require:"true" type:"Repeated"`
  // {"en":"Head value, If condition=HEAD, then this field is mandatory.","zh_CN":"头部值，如果condition为HEAD，则该字段必填"}
  Key *string `json:"key,omitempty" xml:"key,omitempty" require:"true"`
}

func (s QueryAppApiExceptionFeatureDetailResponseDataConfig) String() string {
  return tea.Prettify(s)
}

func (s QueryAppApiExceptionFeatureDetailResponseDataConfig) GoString() string {
  return s.String()
}

func (s *QueryAppApiExceptionFeatureDetailResponseDataConfig) SetCondition(v string) *QueryAppApiExceptionFeatureDetailResponseDataConfig {
  s.Condition = &v
  return s
}

func (s *QueryAppApiExceptionFeatureDetailResponseDataConfig) SetMatchType(v string) *QueryAppApiExceptionFeatureDetailResponseDataConfig {
  s.MatchType = &v
  return s
}

func (s *QueryAppApiExceptionFeatureDetailResponseDataConfig) SetConditionValue(v []*string) *QueryAppApiExceptionFeatureDetailResponseDataConfig {
  s.ConditionValue = v
  return s
}

func (s *QueryAppApiExceptionFeatureDetailResponseDataConfig) SetKey(v string) *QueryAppApiExceptionFeatureDetailResponseDataConfig {
  s.Key = &v
  return s
}

type QueryAppApiExceptionFeatureDetailResponseHeader struct {
}

func (s QueryAppApiExceptionFeatureDetailResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryAppApiExceptionFeatureDetailResponseHeader) GoString() string {
  return s.String()
}




type DeleteAppApiExceptionFeatureRequest struct {
  // {"en":"Rule Ids","zh_CN":"规则ID列表"}
  IdList []*string `json:"idList,omitempty" xml:"idList,omitempty" require:"true" type:"Repeated"`
}

func (s DeleteAppApiExceptionFeatureRequest) String() string {
  return tea.Prettify(s)
}

func (s DeleteAppApiExceptionFeatureRequest) GoString() string {
  return s.String()
}

func (s *DeleteAppApiExceptionFeatureRequest) SetIdList(v []*string) *DeleteAppApiExceptionFeatureRequest {
  s.IdList = v
  return s
}

type DeleteAppApiExceptionFeatureRequestHeader struct {
  // {"en":"Security service type. Please enter a specific service type, if you purchase multiple security services.","zh_CN":"安全服务类型。有使用多个不同的安全服务时，需要填写具体的服务类型。"}
  ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty"`
}

func (s DeleteAppApiExceptionFeatureRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteAppApiExceptionFeatureRequestHeader) GoString() string {
  return s.String()
}

func (s *DeleteAppApiExceptionFeatureRequestHeader) SetServiceType(v string) *DeleteAppApiExceptionFeatureRequestHeader {
  s.ServiceType = &v
  return s
}

type DeleteAppApiExceptionFeaturePaths struct {
}

func (s DeleteAppApiExceptionFeaturePaths) String() string {
  return tea.Prettify(s)
}

func (s DeleteAppApiExceptionFeaturePaths) GoString() string {
  return s.String()
}

type DeleteAppApiExceptionFeatureParameters struct {
}

func (s DeleteAppApiExceptionFeatureParameters) String() string {
  return tea.Prettify(s)
}

func (s DeleteAppApiExceptionFeatureParameters) GoString() string {
  return s.String()
}

type DeleteAppApiExceptionFeatureResponse struct {
  // {"dictionary":"belong=WAAP-MS-Ext|dict=waap_retCodeEnum","en":"Please refer to the error code for exceptions.","zh_CN":"请参照错误码。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Description.","zh_CN":"描述信息。"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"Data.","zh_CN":"出参数据。"}
  Data *bool `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s DeleteAppApiExceptionFeatureResponse) String() string {
  return tea.Prettify(s)
}

func (s DeleteAppApiExceptionFeatureResponse) GoString() string {
  return s.String()
}

func (s *DeleteAppApiExceptionFeatureResponse) SetCode(v string) *DeleteAppApiExceptionFeatureResponse {
  s.Code = &v
  return s
}

func (s *DeleteAppApiExceptionFeatureResponse) SetMsg(v string) *DeleteAppApiExceptionFeatureResponse {
  s.Msg = &v
  return s
}

func (s *DeleteAppApiExceptionFeatureResponse) SetData(v bool) *DeleteAppApiExceptionFeatureResponse {
  s.Data = &v
  return s
}

type DeleteAppApiExceptionFeatureResponseHeader struct {
}

func (s DeleteAppApiExceptionFeatureResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteAppApiExceptionFeatureResponseHeader) GoString() string {
  return s.String()
}




type ListSharedRateLimitingRulesRequest struct {
  // {"en":"Rule name.","zh_CN":"规则名称。"}
  RuleName *string `json:"ruleName,omitempty" xml:"ruleName,omitempty"`
}

func (s ListSharedRateLimitingRulesRequest) String() string {
  return tea.Prettify(s)
}

func (s ListSharedRateLimitingRulesRequest) GoString() string {
  return s.String()
}

func (s *ListSharedRateLimitingRulesRequest) SetRuleName(v string) *ListSharedRateLimitingRulesRequest {
  s.RuleName = &v
  return s
}

type ListSharedRateLimitingRulesRequestHeader struct {
  // {"en":"Security service type. Please enter a specific service type, if you purchase multiple security services.","zh_CN":"安全服务类型。有使用多个不同的安全服务时，需要填写具体的服务类型。"}
  ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty"`
}

func (s ListSharedRateLimitingRulesRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ListSharedRateLimitingRulesRequestHeader) GoString() string {
  return s.String()
}

func (s *ListSharedRateLimitingRulesRequestHeader) SetServiceType(v string) *ListSharedRateLimitingRulesRequestHeader {
  s.ServiceType = &v
  return s
}

type ListSharedRateLimitingRulesPaths struct {
}

func (s ListSharedRateLimitingRulesPaths) String() string {
  return tea.Prettify(s)
}

func (s ListSharedRateLimitingRulesPaths) GoString() string {
  return s.String()
}

type ListSharedRateLimitingRulesParameters struct {
}

func (s ListSharedRateLimitingRulesParameters) String() string {
  return tea.Prettify(s)
}

func (s ListSharedRateLimitingRulesParameters) GoString() string {
  return s.String()
}

type ListSharedRateLimitingRulesResponse struct {
  // {"dictionary":"belong=WAAP-MS-Ext|dict=waap_retCodeEnum","en":"Please refer to the error code for exceptions.","zh_CN":"请参照错误码。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Description.","zh_CN":"描述信息。"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"Data.","zh_CN":"出参数据。"}
  Data []*ListSharedRateLimitingRulesResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s ListSharedRateLimitingRulesResponse) String() string {
  return tea.Prettify(s)
}

func (s ListSharedRateLimitingRulesResponse) GoString() string {
  return s.String()
}

func (s *ListSharedRateLimitingRulesResponse) SetCode(v string) *ListSharedRateLimitingRulesResponse {
  s.Code = &v
  return s
}

func (s *ListSharedRateLimitingRulesResponse) SetMsg(v string) *ListSharedRateLimitingRulesResponse {
  s.Msg = &v
  return s
}

func (s *ListSharedRateLimitingRulesResponse) SetData(v []*ListSharedRateLimitingRulesResponseData) *ListSharedRateLimitingRulesResponse {
  s.Data = v
  return s
}

type ListSharedRateLimitingRulesResponseData struct     {
  // {"en":"Rule ID.","zh_CN":"规则ID。"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {"en":"Creator.","zh_CN":"创建者。"}
  Creator *string `json:"creator,omitempty" xml:"creator,omitempty" require:"true"`
  // {"en":"Rule Name.","zh_CN":"规则名称。"}
  RuleName *string `json:"ruleName,omitempty" xml:"ruleName,omitempty" require:"true"`
  // {"en":"Description.","zh_CN":"规则描述。"}
  Description *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
  // {"en":"Statistical stage.\nREQUEST:Request\nRESPONSE:Response","zh_CN":"统计阶段。\nREQUEST：请求\nRESPONSE：响应"}
  StatisticalStage *string `json:"statisticalStage,omitempty" xml:"statisticalStage,omitempty" require:"true"`
  // {"en":"Client identifier.\nIP:Client IP\nIP_UA:Client IP and User-Agent\nCOOKIE:Cookie\nIP_COOKIE:Client IP and Cookie\nHEADER:Request Header\nIP_HEADER:Client IP and Request Header","zh_CN":"统计粒度。\nIP：客户端IP\nIP_UA：客户端IP和User-Agent\nCOOKIE：Cookie\nIP_COOKIE：客户端IP和Cookie\nHEADER：请求头\nIP_HEADER：客户端IP和请求头"}
  StatisticalItem *string `json:"statisticalItem,omitempty" xml:"statisticalItem,omitempty" require:"true"`
  // {"en":"Statistical key value.","zh_CN":"统计key值。"}
  StatisticsKey *string `json:"statisticsKey,omitempty" xml:"statisticsKey,omitempty" require:"true"`
  // {"en":"Statistics period, unit: seconds.","zh_CN":"统计周期，单位：秒。"}
  StatisticalPeriod *int `json:"statisticalPeriod,omitempty" xml:"statisticalPeriod,omitempty" require:"true"`
  // {"en":"Trigger threshold, unit: times.","zh_CN":"触发阈值，单位：次。"}
  TriggerThreshold *int `json:"triggerThreshold,omitempty" xml:"triggerThreshold,omitempty" require:"true"`
  // {"en":"Action duration, unit: seconds.","zh_CN":"处理动作持续时间，单位：秒。"}
  InterceptTime *int `json:"interceptTime,omitempty" xml:"interceptTime,omitempty" require:"true"`
  // {"en":"Cycle effective status.\nPERMANENT:All time\nWITHOUT:Excluded time\nWITHIN:Selected time","zh_CN":"周期生效状态。\nPERMANENT：永久生效\nWITHOUT：周期内不生效\nWITHIN：周期内生效"}
  EffectiveStatus *string `json:"effectiveStatus,omitempty" xml:"effectiveStatus,omitempty" require:"true"`
  // {"en":"Effective time period.","zh_CN":"规则生效周期。"}
  RateLimitEffective *ListSharedRateLimitingRulesResponseDataRateLimitEffective `json:"rateLimitEffective,omitempty" xml:"rateLimitEffective,omitempty" require:"true" type:"Struct"`
  // {"en":"Action.\nNO_USE:Not Used\nLOG:Log\nCOOKIE:Cookie verification\nJS_CHECK:Javascript verification\nDELAY:Delay\nBLOCK:Deny\nRESET:Reset Connection\nCUSTOM_ACTION_ID:Fill in the custom action id of the corresponding action","zh_CN":"处理动作。\nNO_USE：不使用\nLOG：监控\nCOOKIE：Cookie校验\nJS_CHECK：JavaScript校验\nDELAY：延迟响应\nBLOCK：拦截\nRESET：断开连接\nCUSTOM_ACTION_ID：传入对应的自定义响应动作的ID"}
  Action *string `json:"action,omitempty" xml:"action,omitempty" require:"true"`
  // {"en":"Matching conditions.","zh_CN":"匹配条件。"}
  RateLimitRuleCondition *ListSharedRateLimitingRulesResponseDataRateLimitRuleCondition `json:"rateLimitRuleCondition,omitempty" xml:"rateLimitRuleCondition,omitempty" require:"true" type:"Struct"`
  // {"en":"Associated hostnames.","zh_CN":"关联域名数。"}
  CountShareRateLimitDomain *int `json:"countShareRateLimitDomain,omitempty" xml:"countShareRateLimitDomain,omitempty" require:"true"`
  // {"en":"List of associated hostnames.","zh_CN":"关联域名列表。"}
  ShareRateLimitRelDomains []*string `json:"shareRateLimitRelDomains,omitempty" xml:"shareRateLimitRelDomains,omitempty" require:"true" type:"Repeated"`
  // {"en":"Update time.","zh_CN":"更新时间。"}
  UpdateTime *string `json:"updateTime,omitempty" xml:"updateTime,omitempty" require:"true"`
}

func (s ListSharedRateLimitingRulesResponseData) String() string {
  return tea.Prettify(s)
}

func (s ListSharedRateLimitingRulesResponseData) GoString() string {
  return s.String()
}

func (s *ListSharedRateLimitingRulesResponseData) SetId(v string) *ListSharedRateLimitingRulesResponseData {
  s.Id = &v
  return s
}

func (s *ListSharedRateLimitingRulesResponseData) SetCreator(v string) *ListSharedRateLimitingRulesResponseData {
  s.Creator = &v
  return s
}

func (s *ListSharedRateLimitingRulesResponseData) SetRuleName(v string) *ListSharedRateLimitingRulesResponseData {
  s.RuleName = &v
  return s
}

func (s *ListSharedRateLimitingRulesResponseData) SetDescription(v string) *ListSharedRateLimitingRulesResponseData {
  s.Description = &v
  return s
}

func (s *ListSharedRateLimitingRulesResponseData) SetStatisticalStage(v string) *ListSharedRateLimitingRulesResponseData {
  s.StatisticalStage = &v
  return s
}

func (s *ListSharedRateLimitingRulesResponseData) SetStatisticalItem(v string) *ListSharedRateLimitingRulesResponseData {
  s.StatisticalItem = &v
  return s
}

func (s *ListSharedRateLimitingRulesResponseData) SetStatisticsKey(v string) *ListSharedRateLimitingRulesResponseData {
  s.StatisticsKey = &v
  return s
}

func (s *ListSharedRateLimitingRulesResponseData) SetStatisticalPeriod(v int) *ListSharedRateLimitingRulesResponseData {
  s.StatisticalPeriod = &v
  return s
}

func (s *ListSharedRateLimitingRulesResponseData) SetTriggerThreshold(v int) *ListSharedRateLimitingRulesResponseData {
  s.TriggerThreshold = &v
  return s
}

func (s *ListSharedRateLimitingRulesResponseData) SetInterceptTime(v int) *ListSharedRateLimitingRulesResponseData {
  s.InterceptTime = &v
  return s
}

func (s *ListSharedRateLimitingRulesResponseData) SetEffectiveStatus(v string) *ListSharedRateLimitingRulesResponseData {
  s.EffectiveStatus = &v
  return s
}

func (s *ListSharedRateLimitingRulesResponseData) SetRateLimitEffective(v *ListSharedRateLimitingRulesResponseDataRateLimitEffective) *ListSharedRateLimitingRulesResponseData {
  s.RateLimitEffective = v
  return s
}

func (s *ListSharedRateLimitingRulesResponseData) SetAction(v string) *ListSharedRateLimitingRulesResponseData {
  s.Action = &v
  return s
}

func (s *ListSharedRateLimitingRulesResponseData) SetRateLimitRuleCondition(v *ListSharedRateLimitingRulesResponseDataRateLimitRuleCondition) *ListSharedRateLimitingRulesResponseData {
  s.RateLimitRuleCondition = v
  return s
}

func (s *ListSharedRateLimitingRulesResponseData) SetCountShareRateLimitDomain(v int) *ListSharedRateLimitingRulesResponseData {
  s.CountShareRateLimitDomain = &v
  return s
}

func (s *ListSharedRateLimitingRulesResponseData) SetShareRateLimitRelDomains(v []*string) *ListSharedRateLimitingRulesResponseData {
  s.ShareRateLimitRelDomains = v
  return s
}

func (s *ListSharedRateLimitingRulesResponseData) SetUpdateTime(v string) *ListSharedRateLimitingRulesResponseData {
  s.UpdateTime = &v
  return s
}

type ListSharedRateLimitingRulesResponseDataRateLimitEffective struct {
  // {"en":"effective.\nMON:Monday\nTUE:Tuesday\nWED:Wednesday\nTHU:Thursday\nFRI:Friday\nSAT:Saturday\nSUN:Sunday","zh_CN":"周期。\nMON：星期一\nTUE：星期二\nWED：星期三\nTHU：星期四\nFRI：星期五\nSAT：星期六\nSUN：星期天"}
  Effective []*string `json:"effective,omitempty" xml:"effective,omitempty" require:"true" type:"Repeated"`
  // {"en":"Start time, format: HH:mm.","zh_CN":"开始时间，格式：HH:mm。"}
  Start *string `json:"start,omitempty" xml:"start,omitempty" require:"true"`
  // {"en":"End time, format: HH:mm.","zh_CN":"结束时间，格式：HH:mm。"}
  End *string `json:"end,omitempty" xml:"end,omitempty" require:"true"`
  // {"dictionary":"belong=WAAP-MS-Ext|dict=waap_timezone","en":"Timezone,default value: GMT+8.","zh_CN":"时区，默认：GMT+8。"}
  Timezone *string `json:"timezone,omitempty" xml:"timezone,omitempty" require:"true"`
}

func (s ListSharedRateLimitingRulesResponseDataRateLimitEffective) String() string {
  return tea.Prettify(s)
}

func (s ListSharedRateLimitingRulesResponseDataRateLimitEffective) GoString() string {
  return s.String()
}

func (s *ListSharedRateLimitingRulesResponseDataRateLimitEffective) SetEffective(v []*string) *ListSharedRateLimitingRulesResponseDataRateLimitEffective {
  s.Effective = v
  return s
}

func (s *ListSharedRateLimitingRulesResponseDataRateLimitEffective) SetStart(v string) *ListSharedRateLimitingRulesResponseDataRateLimitEffective {
  s.Start = &v
  return s
}

func (s *ListSharedRateLimitingRulesResponseDataRateLimitEffective) SetEnd(v string) *ListSharedRateLimitingRulesResponseDataRateLimitEffective {
  s.End = &v
  return s
}

func (s *ListSharedRateLimitingRulesResponseDataRateLimitEffective) SetTimezone(v string) *ListSharedRateLimitingRulesResponseDataRateLimitEffective {
  s.Timezone = &v
  return s
}

type ListSharedRateLimitingRulesResponseDataRateLimitRuleCondition struct {
  // {"en":"IP/CIDR, match type cannot be repeated.","zh_CN":"IP/IP段，匹配类型不可重复。"}
  IpOrIpsConditions []*ListSharedRateLimitingRulesResponseDataRateLimitRuleConditionIpOrIpsConditions `json:"ipOrIpsConditions,omitempty" xml:"ipOrIpsConditions,omitempty" require:"true" type:"Repeated"`
  // {"en":"Path, match type cannot be repeated.\nWhen the business scenario is API, this matching condition is not supported.","zh_CN":"路径，匹配类型不可重复。\n当业务场景为API业务时不支持此匹配条件。"}
  PathConditions []*ListSharedRateLimitingRulesResponseDataRateLimitRuleConditionPathConditions `json:"pathConditions,omitempty" xml:"pathConditions,omitempty" require:"true" type:"Repeated"`
  // {"en":"URI, match type cannot be repeated.\nWhen the business scenario is API, this matching condition is not supported.","zh_CN":"URI，匹配类型不可重复。\n当业务场景为API业务时不支持此匹配条件。"}
  UriConditions []*ListSharedRateLimitingRulesResponseDataRateLimitRuleConditionUriConditions `json:"uriConditions,omitempty" xml:"uriConditions,omitempty" require:"true" type:"Repeated"`
  // {"en":"User Agent, match type cannot be repeated.","zh_CN":"User-Agent，匹配类型不可重复。"}
  UaConditions []*ListSharedRateLimitingRulesResponseDataRateLimitRuleConditionUaConditions `json:"uaConditions,omitempty" xml:"uaConditions,omitempty" require:"true" type:"Repeated"`
  // {"en":"Request Method.\nWhen the business scenario is API,this matching condition is not supported.","zh_CN":"请求方法，匹配类型不可重复。\n当业务场景为API业务时不支持此匹配条件。"}
  MethodConditions []*ListSharedRateLimitingRulesResponseDataRateLimitRuleConditionMethodConditions `json:"methodConditions,omitempty" xml:"methodConditions,omitempty" require:"true" type:"Repeated"`
  // {"en":"Referer, match type cannot be repeated.","zh_CN":"Referer，匹配类型不可重复。"}
  RefererConditions []*ListSharedRateLimitingRulesResponseDataRateLimitRuleConditionRefererConditions `json:"refererConditions,omitempty" xml:"refererConditions,omitempty" require:"true" type:"Repeated"`
  // {"en":"Request Header, match type can be repeated.","zh_CN":"请求头，匹配类型可重复。"}
  HeaderConditions []*ListSharedRateLimitingRulesResponseDataRateLimitRuleConditionHeaderConditions `json:"headerConditions,omitempty" xml:"headerConditions,omitempty" require:"true" type:"Repeated"`
  // {"en":"Geo, match type cannot be repeated.","zh_CN":"区域，匹配类型不可重复。"}
  AreaConditions []*ListSharedRateLimitingRulesResponseDataRateLimitRuleConditionAreaConditions `json:"areaConditions,omitempty" xml:"areaConditions,omitempty" require:"true" type:"Repeated"`
  // {"en":"HTTP/S, match type cannot be repeated.","zh_CN":"应用层协议，匹配类型不可重复。"}
  SchemeConditions []*ListSharedRateLimitingRulesResponseDataRateLimitRuleConditionSchemeConditions `json:"schemeConditions,omitempty" xml:"schemeConditions,omitempty" require:"true" type:"Repeated"`
  // {"en":"Response Code, match type cannot be repeated.","zh_CN":"状态码，匹配类型不可重复。"}
  StatusCodeConditions []*ListSharedRateLimitingRulesResponseDataRateLimitRuleConditionStatusCodeConditions `json:"statusCodeConditions,omitempty" xml:"statusCodeConditions,omitempty" require:"true" type:"Repeated"`
  // {"en":"JA3 Fingerprint, match type cannot be repeated.","zh_CN":"JA3指纹，匹配类型不可重复。"}
  Ja3Conditions []*ListSharedRateLimitingRulesResponseDataRateLimitRuleConditionJa3Conditions `json:"ja3Conditions,omitempty" xml:"ja3Conditions,omitempty" require:"true" type:"Repeated"`
  // {"en":"JA4 Fingerprint, match type cannot be repeated.","zh_CN":"JA4指纹，匹配类型不可重复。"}
  Ja4Conditions []*ListSharedRateLimitingRulesResponseDataRateLimitRuleConditionJa4Conditions `json:"ja4Conditions,omitempty" xml:"ja4Conditions,omitempty" require:"true" type:"Repeated"`
}

func (s ListSharedRateLimitingRulesResponseDataRateLimitRuleCondition) String() string {
  return tea.Prettify(s)
}

func (s ListSharedRateLimitingRulesResponseDataRateLimitRuleCondition) GoString() string {
  return s.String()
}

func (s *ListSharedRateLimitingRulesResponseDataRateLimitRuleCondition) SetIpOrIpsConditions(v []*ListSharedRateLimitingRulesResponseDataRateLimitRuleConditionIpOrIpsConditions) *ListSharedRateLimitingRulesResponseDataRateLimitRuleCondition {
  s.IpOrIpsConditions = v
  return s
}

func (s *ListSharedRateLimitingRulesResponseDataRateLimitRuleCondition) SetPathConditions(v []*ListSharedRateLimitingRulesResponseDataRateLimitRuleConditionPathConditions) *ListSharedRateLimitingRulesResponseDataRateLimitRuleCondition {
  s.PathConditions = v
  return s
}

func (s *ListSharedRateLimitingRulesResponseDataRateLimitRuleCondition) SetUriConditions(v []*ListSharedRateLimitingRulesResponseDataRateLimitRuleConditionUriConditions) *ListSharedRateLimitingRulesResponseDataRateLimitRuleCondition {
  s.UriConditions = v
  return s
}

func (s *ListSharedRateLimitingRulesResponseDataRateLimitRuleCondition) SetUaConditions(v []*ListSharedRateLimitingRulesResponseDataRateLimitRuleConditionUaConditions) *ListSharedRateLimitingRulesResponseDataRateLimitRuleCondition {
  s.UaConditions = v
  return s
}

func (s *ListSharedRateLimitingRulesResponseDataRateLimitRuleCondition) SetMethodConditions(v []*ListSharedRateLimitingRulesResponseDataRateLimitRuleConditionMethodConditions) *ListSharedRateLimitingRulesResponseDataRateLimitRuleCondition {
  s.MethodConditions = v
  return s
}

func (s *ListSharedRateLimitingRulesResponseDataRateLimitRuleCondition) SetRefererConditions(v []*ListSharedRateLimitingRulesResponseDataRateLimitRuleConditionRefererConditions) *ListSharedRateLimitingRulesResponseDataRateLimitRuleCondition {
  s.RefererConditions = v
  return s
}

func (s *ListSharedRateLimitingRulesResponseDataRateLimitRuleCondition) SetHeaderConditions(v []*ListSharedRateLimitingRulesResponseDataRateLimitRuleConditionHeaderConditions) *ListSharedRateLimitingRulesResponseDataRateLimitRuleCondition {
  s.HeaderConditions = v
  return s
}

func (s *ListSharedRateLimitingRulesResponseDataRateLimitRuleCondition) SetAreaConditions(v []*ListSharedRateLimitingRulesResponseDataRateLimitRuleConditionAreaConditions) *ListSharedRateLimitingRulesResponseDataRateLimitRuleCondition {
  s.AreaConditions = v
  return s
}

func (s *ListSharedRateLimitingRulesResponseDataRateLimitRuleCondition) SetSchemeConditions(v []*ListSharedRateLimitingRulesResponseDataRateLimitRuleConditionSchemeConditions) *ListSharedRateLimitingRulesResponseDataRateLimitRuleCondition {
  s.SchemeConditions = v
  return s
}

func (s *ListSharedRateLimitingRulesResponseDataRateLimitRuleCondition) SetStatusCodeConditions(v []*ListSharedRateLimitingRulesResponseDataRateLimitRuleConditionStatusCodeConditions) *ListSharedRateLimitingRulesResponseDataRateLimitRuleCondition {
  s.StatusCodeConditions = v
  return s
}

func (s *ListSharedRateLimitingRulesResponseDataRateLimitRuleCondition) SetJa3Conditions(v []*ListSharedRateLimitingRulesResponseDataRateLimitRuleConditionJa3Conditions) *ListSharedRateLimitingRulesResponseDataRateLimitRuleCondition {
  s.Ja3Conditions = v
  return s
}

func (s *ListSharedRateLimitingRulesResponseDataRateLimitRuleCondition) SetJa4Conditions(v []*ListSharedRateLimitingRulesResponseDataRateLimitRuleConditionJa4Conditions) *ListSharedRateLimitingRulesResponseDataRateLimitRuleCondition {
  s.Ja4Conditions = v
  return s
}

type ListSharedRateLimitingRulesResponseDataRateLimitRuleConditionIpOrIpsConditions struct     {
  // {"en":"Match type.\nEQUAL:Equals\nNOT_EQUAL:Does not equal","zh_CN":"匹配类型。\nEQUAL：等于\nNOT_EQUAL：不等于","exampleValue":"EQUAL,NOT_EQUAL"}
  MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
  // {"en":"IP/CIDR, maximum 500 IP/CIDR.","zh_CN":"IP/IP段，最多500个IP/IP段。"}
  IpOrIps []*string `json:"ipOrIps,omitempty" xml:"ipOrIps,omitempty" require:"true" type:"Repeated"`
}

func (s ListSharedRateLimitingRulesResponseDataRateLimitRuleConditionIpOrIpsConditions) String() string {
  return tea.Prettify(s)
}

func (s ListSharedRateLimitingRulesResponseDataRateLimitRuleConditionIpOrIpsConditions) GoString() string {
  return s.String()
}

func (s *ListSharedRateLimitingRulesResponseDataRateLimitRuleConditionIpOrIpsConditions) SetMatchType(v string) *ListSharedRateLimitingRulesResponseDataRateLimitRuleConditionIpOrIpsConditions {
  s.MatchType = &v
  return s
}

func (s *ListSharedRateLimitingRulesResponseDataRateLimitRuleConditionIpOrIpsConditions) SetIpOrIps(v []*string) *ListSharedRateLimitingRulesResponseDataRateLimitRuleConditionIpOrIpsConditions {
  s.IpOrIps = v
  return s
}

type ListSharedRateLimitingRulesResponseDataRateLimitRuleConditionPathConditions struct     {
  // {"en":"Match type.\nEQUAL: Equals, path case sensitive\nNOT_EQUAL: Does not equal, path case sensitive\nCONTAIN: Contains, path case insensitive\nNOT_CONTAIN: Does not Contains, path case insensitive\nREGEX: Regex match, path case insensitive\nNOT_REGEX: Regular does not match, path case sensitive\nSTART_WITH: Starts with, path case sensitive\nEND_WITH: Ends with, path case sensitive\nWILDCARD: Wildcard matches, path case sensitive,* represents zero or more arbitrary characters, ? represents any single character.\nNOT_WILDCARD: Wildcard does not match, path case sensitive,* represents zero or more arbitrary characters, ? represents any single character","zh_CN":"匹配类型。\nEQUAL：等于，路径大小写敏感\nNOT_EQUAL：不等于，路径大小写敏感\nCONTAIN：包含，路径大小写不敏感\nNOT_CONTAIN：不包含，路径大小写不敏感\nREGEX：匹配正则，路径大小写不敏感\nNOT_REGEX：正则不匹配，路径大小写不敏感\nSTART_WITH：开头是，路径大小写不敏感\nEND_WITH：结尾是，路径大小写不敏感\nWILDCARD：通配符匹配，路径大小写不敏感，*代表零个或多个任意字符，?代表任意单个字符\nNOT_WILDCARD：通配符不匹配，路径大小写不敏感，*代表零个或多个任意字符，?代表任意单个字符","exampleValue":"EQUAL,NOT_EQUAL,CONTAIN,NOT_CONTAIN,REGEX,NOT_REGEX,START_WITH,END_WITH,WILDCARD,NOT_WILDCARD"}
  MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
  // {"en":"Path.\nWhen match type is EQUAL/NOT_EQUAL/START_WITH/END_WITH, path needs to start with \"/\", and no parameters.\nWhen the match type is REGEX/NOT_REGEX, only one value is allowed.\nExample: /test.html.","zh_CN":"路径。\n当匹配类型为等于/不等于/开头是/结尾是，路径必须以“/”开头，不含参数。\n当匹配类型为正则/正则不匹配，则只允许只有一个值。\n示例：/test.html。"}
  ListSharedRateLimitingRulesPaths []*string `json:"paths,omitempty" xml:"paths,omitempty" require:"true" type:"Repeated"`
}

func (s ListSharedRateLimitingRulesResponseDataRateLimitRuleConditionPathConditions) String() string {
  return tea.Prettify(s)
}

func (s ListSharedRateLimitingRulesResponseDataRateLimitRuleConditionPathConditions) GoString() string {
  return s.String()
}

func (s *ListSharedRateLimitingRulesResponseDataRateLimitRuleConditionPathConditions) SetMatchType(v string) *ListSharedRateLimitingRulesResponseDataRateLimitRuleConditionPathConditions {
  s.MatchType = &v
  return s
}

func (s *ListSharedRateLimitingRulesResponseDataRateLimitRuleConditionPathConditions) SetListSharedRateLimitingRulesPaths(v []*string) *ListSharedRateLimitingRulesResponseDataRateLimitRuleConditionPathConditions {
  s.ListSharedRateLimitingRulesPaths = v
  return s
}

type ListSharedRateLimitingRulesResponseDataRateLimitRuleConditionUriConditions struct     {
  // {"en":"Match type.\nEQUAL: Equals, URI case sensitive\nNOT_EQUAL: Does not equal, URI case sensitive\nCONTAIN: Contains, URI case insensitive\nNOT_CONTAIN: Does not Contains, URI case insensitive\nREGEX: Regex match, URI case insensitive\nNOT_REGEX: Regular does not match, URI case insensitive\nSTART_WITH: Starts with, URI case insensitive\nEND_WITH: Ends with, URI case insensitive\nWILDCARD: Wildcard matches, URI case insensitive,* represents zero or more arbitrary characters, ? represents any single character\nNOT_WILDCARD: Wildcard does not match, URI case insensitive,* represents zero or more arbitrary characters, ? represents any single character","zh_CN":"匹配类型。\nEQUAL：等于，URI大小写敏感\nNOT_EQUAL：不等于，URI大小写敏感\nCONTAIN：包含，URI大小写不敏感\nNOT_CONTAIN：不包含，URI大小写不敏感\nREGEX：匹配正则，URI大小写不敏感\nNOT_REGEX：正则不匹配，URI大小写不敏感\nSTART_WITH：开头是，URI大小写不敏感\nEND_WITH：结尾是，URI大小写不敏感\nWILDCARD：通配符匹配，URI大小写不敏感，*代表零个或多个任意字符，?代表任意单个字符\nNOT_WILDCARD：通配符不匹配，URI大小写不敏感，*代表零个或多个任意字符，?代表任意单个字符","exampleValue":"EQUAL,NOT_EQUAL,CONTAIN,NOT_CONTAIN,REGEX,NOT_REGEX,START_WITH,END_WITH,WILDCARD,NOT_WILDCARD"}
  MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
  // {"en":"URI.\nWhen match type is EQUAL/NOT_EQUAL/START_WITH/END_WITH, uri needs to start with \"/\", and includes parameters.\nWhen the match type is REGEX/NOT_REGEX, only one value is allowed.\nExample: /test.html?id=1.","zh_CN":"URI。\n当匹配类型为等于/不等于/开头是/结尾是，URI必须以”/“开头，含参数。\n当匹配类型为正则/正则不匹配，则只允许只有一个值。\n示例：/test.html?id=1。"}
  Uri []*string `json:"uri,omitempty" xml:"uri,omitempty" require:"true" type:"Repeated"`
}

func (s ListSharedRateLimitingRulesResponseDataRateLimitRuleConditionUriConditions) String() string {
  return tea.Prettify(s)
}

func (s ListSharedRateLimitingRulesResponseDataRateLimitRuleConditionUriConditions) GoString() string {
  return s.String()
}

func (s *ListSharedRateLimitingRulesResponseDataRateLimitRuleConditionUriConditions) SetMatchType(v string) *ListSharedRateLimitingRulesResponseDataRateLimitRuleConditionUriConditions {
  s.MatchType = &v
  return s
}

func (s *ListSharedRateLimitingRulesResponseDataRateLimitRuleConditionUriConditions) SetUri(v []*string) *ListSharedRateLimitingRulesResponseDataRateLimitRuleConditionUriConditions {
  s.Uri = v
  return s
}

type ListSharedRateLimitingRulesResponseDataRateLimitRuleConditionUaConditions struct     {
  // {"en":"Match type.\nEQUAL: Equals, user agent case sensitive\nNOT_EQUAL: Does not equal, user agent case sensitive\nCONTAIN: Contains, user agent case insensitive\nNOT_CONTAIN: Does not Contains, user agent case insensitive\nNONE:Empty or non-existent\nREGEX: Regex match, user agent case insensitive\nNOT_REGEX: Regular does not match, user agent case insensitive\nSTART_WITH: Starts with, user agent case insensitive\nEND_WITH: Ends with, user agent case insensitive\nWILDCARD: Wildcard matches, user agent case insensitive,* represents zero or more arbitrary characters, ? represents any single character\nNOT_WILDCARD: Wildcard does not match, user agent case insensitive,* represents zero or more arbitrary characters, ? represents any single character","zh_CN":"匹配类型。\nEQUAL：等于，User-Agent大小写敏感\nNOT_EQUAL：不等于，User-Agent大小写敏感\nCONTAIN：包含，User-Agent大小写不敏感\nNOT_CONTAIN：不包含，User-Agent大小写不敏感\nNONE：为空或不存在\nREGEX：匹配正则，User-Agent大小写不敏感\nNOT_REGEX：正则不匹配，User-Agent大小写不敏感\nSTART_WITH：开头是，User-Agent大小写不敏感\nEND_WITH：结尾是，User-Agent大小写不敏感\nWILDCARD：通配符匹配，User-Agent大小写不敏感，*代表零个或多个任意字符，?代表任意单个字符\nNOT_WILDCARD：通配符不匹配，User-Agent大小写不敏感，*代表零个或多个任意字符，?代表任意单个字符","exampleValue":"EQUAL,NOT_EQUAL,CONTAIN,NOT_CONTAIN,NONE,REGEX,NOT_REGEX,START_WITH,END_WITH,WILDCARD,NOT_WILDCARD"}
  MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
  // {"en":"User agent.\nWhen the match type is REGEX/NOT_REGEX, only one value is allowed.\nExample: go-Http-client/1.1.","zh_CN":"User-Agent。\n当匹配类型为正则/正则不匹配，则只允许只有一个值。\n示例：go-Http-client/1.1。"}
  Ua []*string `json:"ua,omitempty" xml:"ua,omitempty" require:"true" type:"Repeated"`
}

func (s ListSharedRateLimitingRulesResponseDataRateLimitRuleConditionUaConditions) String() string {
  return tea.Prettify(s)
}

func (s ListSharedRateLimitingRulesResponseDataRateLimitRuleConditionUaConditions) GoString() string {
  return s.String()
}

func (s *ListSharedRateLimitingRulesResponseDataRateLimitRuleConditionUaConditions) SetMatchType(v string) *ListSharedRateLimitingRulesResponseDataRateLimitRuleConditionUaConditions {
  s.MatchType = &v
  return s
}

func (s *ListSharedRateLimitingRulesResponseDataRateLimitRuleConditionUaConditions) SetUa(v []*string) *ListSharedRateLimitingRulesResponseDataRateLimitRuleConditionUaConditions {
  s.Ua = v
  return s
}

type ListSharedRateLimitingRulesResponseDataRateLimitRuleConditionMethodConditions struct     {
  // {"en":"Match type.\nEQUAL:Equals\nNOT_EQUAL:Does not equal","zh_CN":"匹配类型。\nEQUAL：等于\nNOT_EQUAL：不等于","exampleValue":"EQUAL,NOT_EQUAL"}
  MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
  // {"en":"Request method.\nSupported values: GET/POST/DELETE/PUT/HEAD/OPTIONS/COPY.","zh_CN":"请求方法。\n支持的值：GET/POST/DELETE/PUT/HEAD/OPTIONS/COPY。","exampleValue":"GET,POST,DELETE,PUT,HEAD,OPTIONS,COPY"}
  RequestMethod []*string `json:"requestMethod,omitempty" xml:"requestMethod,omitempty" require:"true" type:"Repeated"`
}

func (s ListSharedRateLimitingRulesResponseDataRateLimitRuleConditionMethodConditions) String() string {
  return tea.Prettify(s)
}

func (s ListSharedRateLimitingRulesResponseDataRateLimitRuleConditionMethodConditions) GoString() string {
  return s.String()
}

func (s *ListSharedRateLimitingRulesResponseDataRateLimitRuleConditionMethodConditions) SetMatchType(v string) *ListSharedRateLimitingRulesResponseDataRateLimitRuleConditionMethodConditions {
  s.MatchType = &v
  return s
}

func (s *ListSharedRateLimitingRulesResponseDataRateLimitRuleConditionMethodConditions) SetRequestMethod(v []*string) *ListSharedRateLimitingRulesResponseDataRateLimitRuleConditionMethodConditions {
  s.RequestMethod = v
  return s
}

type ListSharedRateLimitingRulesResponseDataRateLimitRuleConditionRefererConditions struct     {
  // {"en":"Referer.\nWhen the match type is REGEX/NOT_REGEX, only one value is allowed.\nExample: http://test.com.","zh_CN":"Referer。\n当匹配类型为正则/正则不匹配，则只允许只有一个值。\n示例：http://test.com。"}
  Referer []*string `json:"referer,omitempty" xml:"referer,omitempty" require:"true" type:"Repeated"`
  // {"en":"Match type.\nEQUAL: Equals, referer case sensitive\nNOT_EQUAL: Does not equal, referer case sensitive\nCONTAIN: Contains, referer case insensitive\nNOT_CONTAIN: Does not Contains, referer case insensitive\nNONE:Empty or non-existent\nREGEX: Regex match, referer case insensitive\nNOT_REGEX: Regular does not match, referer case insensitive\nSTART_WITH: Starts with, referer case insensitive\nEND_WITH: Ends with, referer case insensitive\nWILDCARD: Wildcard matches, referer case insensitive,* represents zero or more arbitrary characters, ? represents any single characte\nNOT_WILDCARD: Wildcard does not match, referer case insensitive,* represents zero or more arbitrary characters, ? represents any single character","zh_CN":"匹配类型。\nEQUAL：等于，referer大小写敏感\nNOT_EQUAL：不等于，referer大小写敏感\nCONTAIN：包含，referer大小写不敏感\nNOT_CONTAIN：不包含，referer大小写不敏感\nNONE：为空或不存在\nREGEX：匹配正则，referer大小写不敏感\nNOT_REGEX：正则不匹配，referer大小写不敏感\nSTART_WITH：开头是，referer大小写不敏感\nEND_WITH：结尾是，referer大小写不敏感\nWILDCARD：通配符匹配，referer大小写不敏感，*代表零个或多个任意字符，?代表任意单个字符\nNOT_WILDCARD：通配符不匹配，referer大小写不敏感，*代表零个或多个任意字符，?代表任意单个字符","exampleValue":"EQUAL,NOT_EQUAL,CONTAIN,NOT_CONTAIN,NONE,REGEX,NOT_REGEX,START_WITH,END_WITH,WILDCARD,NOT_WILDCARD"}
  MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
}

func (s ListSharedRateLimitingRulesResponseDataRateLimitRuleConditionRefererConditions) String() string {
  return tea.Prettify(s)
}

func (s ListSharedRateLimitingRulesResponseDataRateLimitRuleConditionRefererConditions) GoString() string {
  return s.String()
}

func (s *ListSharedRateLimitingRulesResponseDataRateLimitRuleConditionRefererConditions) SetReferer(v []*string) *ListSharedRateLimitingRulesResponseDataRateLimitRuleConditionRefererConditions {
  s.Referer = v
  return s
}

func (s *ListSharedRateLimitingRulesResponseDataRateLimitRuleConditionRefererConditions) SetMatchType(v string) *ListSharedRateLimitingRulesResponseDataRateLimitRuleConditionRefererConditions {
  s.MatchType = &v
  return s
}

type ListSharedRateLimitingRulesResponseDataRateLimitRuleConditionHeaderConditions struct     {
  // {"en":"Match type.EQUAL: Equals, request header values case sensitiveNOT_EQUAL: Does not equal, request header values case sensitiveCONTAIN: Contains, request header values case insensitiveNOT_CONTAIN: Does not Contains, request header values case insensitiveNONE: Empty or non-existentREGEX: Regex match, request header values case insensitiveNOT_REGEX: Regular does not match, request header values case insensitiveSTART_WITH: Starts with, request header values case insensitiveEND_WITH: Ends with, request header values case insensitiveWILDCARD: Wildcard matches, request header values case insensitive,* represents zero or more arbitrary characters, ? represents any single characterNOT_WILDCARD: Wildcard does not match, request header values case insensitive,* represents zero or more arbitrary characters, ? represents any single character","zh_CN":"匹配类型。EQUAL：等于，头部值大小写敏感NOT_EQUAL：不等于，头部值大小写敏感CONTAIN：包含，头部值大小写不敏感NOT_CONTAIN：不包含，头部值大小写不敏感NONE：为空或不存在REGEX：匹配正则，头部值大小写不敏感NOT_REGEX：正则不匹配，头部值大小写不敏感START_WITH：开头是，头部值大小写不敏感END_WITH：结尾是，头部值大小写不敏感WILDCARD：通配符匹配，头部值大小写不敏感，*代表零个或多个任意字符，?代表任意单个字符NOT_WILDCARD：通配符不匹配，头部值大小写不敏感，*代表零个或多个任意字符，?代表任意单个字符","exampleValue":"EQUAL,NOT_EQUAL,CONTAIN,NOT_CONTAIN,NONE,REGEX,NOT_REGEX,START_WITH,END_WITH,WILDCARD,NOT_WILDCARD"}
  MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
  // {"en":"Header value.\nWhen the match type is REGEX/NOT_REGEX, only one value is allowed.","zh_CN":"头部值。\n当匹配类型为正则/正则不匹配，则只允许只有一个值。"}
  ValueList []*string `json:"valueList,omitempty" xml:"valueList,omitempty" require:"true" type:"Repeated"`
  // {"en":"Header name,case insensitive,up to 100 characters.\nExample: Accept.","zh_CN":"头部名称，大小写不敏感，最多100个字符。\n示例：Accept。"}
  Key *string `json:"key,omitempty" xml:"key,omitempty" require:"true"`
}

func (s ListSharedRateLimitingRulesResponseDataRateLimitRuleConditionHeaderConditions) String() string {
  return tea.Prettify(s)
}

func (s ListSharedRateLimitingRulesResponseDataRateLimitRuleConditionHeaderConditions) GoString() string {
  return s.String()
}

func (s *ListSharedRateLimitingRulesResponseDataRateLimitRuleConditionHeaderConditions) SetMatchType(v string) *ListSharedRateLimitingRulesResponseDataRateLimitRuleConditionHeaderConditions {
  s.MatchType = &v
  return s
}

func (s *ListSharedRateLimitingRulesResponseDataRateLimitRuleConditionHeaderConditions) SetValueList(v []*string) *ListSharedRateLimitingRulesResponseDataRateLimitRuleConditionHeaderConditions {
  s.ValueList = v
  return s
}

func (s *ListSharedRateLimitingRulesResponseDataRateLimitRuleConditionHeaderConditions) SetKey(v string) *ListSharedRateLimitingRulesResponseDataRateLimitRuleConditionHeaderConditions {
  s.Key = &v
  return s
}

type ListSharedRateLimitingRulesResponseDataRateLimitRuleConditionAreaConditions struct     {
  // {"en":"Match type.\nEQUAL:Equals\nNOT_EQUAL:Does not equal","zh_CN":"匹配类型。\nEQUAL：等于\nNOT_EQUAL：不等于","exampleValue":"EQUAL,NOT_EQUAL"}
  MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
  // {"dictionary":"belong=WAAP-MS-Ext|dict=waap_areaCityAndCountry","en":"Geo.","zh_CN":"区域。"}
  Areas []*string `json:"areas,omitempty" xml:"areas,omitempty" require:"true" type:"Repeated"`
}

func (s ListSharedRateLimitingRulesResponseDataRateLimitRuleConditionAreaConditions) String() string {
  return tea.Prettify(s)
}

func (s ListSharedRateLimitingRulesResponseDataRateLimitRuleConditionAreaConditions) GoString() string {
  return s.String()
}

func (s *ListSharedRateLimitingRulesResponseDataRateLimitRuleConditionAreaConditions) SetMatchType(v string) *ListSharedRateLimitingRulesResponseDataRateLimitRuleConditionAreaConditions {
  s.MatchType = &v
  return s
}

func (s *ListSharedRateLimitingRulesResponseDataRateLimitRuleConditionAreaConditions) SetAreas(v []*string) *ListSharedRateLimitingRulesResponseDataRateLimitRuleConditionAreaConditions {
  s.Areas = v
  return s
}

type ListSharedRateLimitingRulesResponseDataRateLimitRuleConditionSchemeConditions struct     {
  // {"en":"Match type.\nEQUAL:Equals\nNOT_EQUAL:Does not equal","zh_CN":"匹配类型。\nEQUAL：等于\nNOT_EQUAL：不等于","exampleValue":"EQUAL,NOT_EQUAL"}
  MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
  // {"en":"HTTP/S.\nSupported values: HTTP/HTTPS.","zh_CN":"应用层协议。\n支持的值：HTTP/HTTPS。","exampleValue":"HTTP,HTTPS"}
  Scheme []*string `json:"scheme,omitempty" xml:"scheme,omitempty" require:"true" type:"Repeated"`
}

func (s ListSharedRateLimitingRulesResponseDataRateLimitRuleConditionSchemeConditions) String() string {
  return tea.Prettify(s)
}

func (s ListSharedRateLimitingRulesResponseDataRateLimitRuleConditionSchemeConditions) GoString() string {
  return s.String()
}

func (s *ListSharedRateLimitingRulesResponseDataRateLimitRuleConditionSchemeConditions) SetMatchType(v string) *ListSharedRateLimitingRulesResponseDataRateLimitRuleConditionSchemeConditions {
  s.MatchType = &v
  return s
}

func (s *ListSharedRateLimitingRulesResponseDataRateLimitRuleConditionSchemeConditions) SetScheme(v []*string) *ListSharedRateLimitingRulesResponseDataRateLimitRuleConditionSchemeConditions {
  s.Scheme = v
  return s
}

type ListSharedRateLimitingRulesResponseDataRateLimitRuleConditionStatusCodeConditions struct     {
  // {"en":"Match type.\nEQUAL:Equals\nNOT_EQUAL:Does not equal","zh_CN":"匹配类型。\nEQUAL：等于\nNOT_EQUAL：不等于","exampleValue":"EQUAL,NOT_EQUAL"}
  MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
  // {"en":"Response Code.","zh_CN":"状态码。"}
  StatusCode []*string `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true" type:"Repeated"`
}

func (s ListSharedRateLimitingRulesResponseDataRateLimitRuleConditionStatusCodeConditions) String() string {
  return tea.Prettify(s)
}

func (s ListSharedRateLimitingRulesResponseDataRateLimitRuleConditionStatusCodeConditions) GoString() string {
  return s.String()
}

func (s *ListSharedRateLimitingRulesResponseDataRateLimitRuleConditionStatusCodeConditions) SetMatchType(v string) *ListSharedRateLimitingRulesResponseDataRateLimitRuleConditionStatusCodeConditions {
  s.MatchType = &v
  return s
}

func (s *ListSharedRateLimitingRulesResponseDataRateLimitRuleConditionStatusCodeConditions) SetStatusCode(v []*string) *ListSharedRateLimitingRulesResponseDataRateLimitRuleConditionStatusCodeConditions {
  s.StatusCode = v
  return s
}

type ListSharedRateLimitingRulesResponseDataRateLimitRuleConditionJa3Conditions struct     {
  // {"en":"Match type.\nEQUAL: Equals\nNOT_EQUAL: Does not equal","zh_CN":"匹配类型。\nEQUAL：等于 \nNOT_EQUAL：不等于","exampleValue":"EQUAL,NOT_EQUAL"}
  MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
  // {"en":"JA3 Fingerprint List, maximum 300 JA3 Fingerprint.\nWhen the match type is EQUAL/NOT_EQUAL, each item's character length must be 32 and can only include numbers and lowercase letters.","zh_CN":"JA3指纹列表，最多300个JA3指纹。\n当匹配类型为等于/不等于时，每项字符长度必须为32，且仅限数字和小写字母。"}
  Ja3List []*string `json:"ja3List,omitempty" xml:"ja3List,omitempty" require:"true" type:"Repeated"`
}

func (s ListSharedRateLimitingRulesResponseDataRateLimitRuleConditionJa3Conditions) String() string {
  return tea.Prettify(s)
}

func (s ListSharedRateLimitingRulesResponseDataRateLimitRuleConditionJa3Conditions) GoString() string {
  return s.String()
}

func (s *ListSharedRateLimitingRulesResponseDataRateLimitRuleConditionJa3Conditions) SetMatchType(v string) *ListSharedRateLimitingRulesResponseDataRateLimitRuleConditionJa3Conditions {
  s.MatchType = &v
  return s
}

func (s *ListSharedRateLimitingRulesResponseDataRateLimitRuleConditionJa3Conditions) SetJa3List(v []*string) *ListSharedRateLimitingRulesResponseDataRateLimitRuleConditionJa3Conditions {
  s.Ja3List = v
  return s
}

type ListSharedRateLimitingRulesResponseDataRateLimitRuleConditionJa4Conditions struct     {
  // {"en":"Match type. \nEQUAL: Equals\nNOT_EQUAL: Does not equal\nCONTAIN: Contains\nNOT_CONTAIN: Does not Contains\nSTART_WITH: Starts with\nEND_WITH: Ends with\nWILDCARD: Wildcard matches, ** represents zero or more arbitrary characters, ? represents any single character\nNOT_WILDCARD: Wildcard does not match, ** represents zero or more arbitrary characters, ? represents any single character","zh_CN":"匹配类型。\nEQUAL：等于\nNOT_EQUAL：不等于\nCONTAIN：包含\nNOT_CONTAIN：不包含\nSTART_WITH：开头是\nEND_WITH：结尾是\nWILDCARD：通配符匹配，*代表零个或多个任意字符，?代表任意单个字符\nNOT_WILDCARD：通配符不匹配，*代表零个或多个任意字符，?代表任意单个字符","exampleValue":"EQUAL,NOT_EQUAL,CONTAIN,NOT_CONTAIN,START_WITH,END_WITH,WILDCARD,NOT_WILDCARD"}
  MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
  // {"en":"JA4 Fingerprint List, maximum 300 JA4 Fingerprint.\nWhen the match type is EQUAL/NOT_EQUAL, each item's format must be 10 characters + 12 characters + 12 characters, separated by underscores, and can only include underscores, numbers, and lowercase letters.\nWhen the match type is CONTAIN/NOT_CONTAIN/START_WITH/END_WITH, each item is only allowed to include underscores, numbers, and lowercase letters.\nWhen the match type is WILDCARD/NOT_WILDCARD, each item, aside from  ** and ?, is only allowed to include underscores, numbers, and lowercase letters.","zh_CN":"JA4指纹列表，最多300个JA4指纹。\n当匹配类型为等于/不等于时，每项格式必须为10位字符+12位字符+12位字符，中间以“_”分隔，且仅限下划线、数字和小写字母。\n当匹配类型为包含/不包含/开头是/结尾是时，每项只允许输入下划线、数字和小写字母。\n当匹配类型为通配符匹配/通配符不匹配时，每项除*和?外，只允许输入下划线、数字和小写字母。"}
  Ja4List []*string `json:"ja4List,omitempty" xml:"ja4List,omitempty" require:"true" type:"Repeated"`
}

func (s ListSharedRateLimitingRulesResponseDataRateLimitRuleConditionJa4Conditions) String() string {
  return tea.Prettify(s)
}

func (s ListSharedRateLimitingRulesResponseDataRateLimitRuleConditionJa4Conditions) GoString() string {
  return s.String()
}

func (s *ListSharedRateLimitingRulesResponseDataRateLimitRuleConditionJa4Conditions) SetMatchType(v string) *ListSharedRateLimitingRulesResponseDataRateLimitRuleConditionJa4Conditions {
  s.MatchType = &v
  return s
}

func (s *ListSharedRateLimitingRulesResponseDataRateLimitRuleConditionJa4Conditions) SetJa4List(v []*string) *ListSharedRateLimitingRulesResponseDataRateLimitRuleConditionJa4Conditions {
  s.Ja4List = v
  return s
}

type ListSharedRateLimitingRulesResponseHeader struct {
}

func (s ListSharedRateLimitingRulesResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ListSharedRateLimitingRulesResponseHeader) GoString() string {
  return s.String()
}




type DeleteSharedCustomRulesRequest struct {
  // {"en":"Rule ID list.","zh_CN":"规则ID列表。"}
  IdList []*string `json:"idList,omitempty" xml:"idList,omitempty" require:"true" type:"Repeated"`
}

func (s DeleteSharedCustomRulesRequest) String() string {
  return tea.Prettify(s)
}

func (s DeleteSharedCustomRulesRequest) GoString() string {
  return s.String()
}

func (s *DeleteSharedCustomRulesRequest) SetIdList(v []*string) *DeleteSharedCustomRulesRequest {
  s.IdList = v
  return s
}

type DeleteSharedCustomRulesRequestHeader struct {
  // {"dictionary":"belong=WAAP-MS-Ext|dict=waap_serviceType","en":"Security service type. Please enter a specific service type, if you purchase multiple security services.","zh_CN":"安全服务类型。有使用多个不同的安全服务时，需要填写具体的服务类型。"}
  ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty"`
}

func (s DeleteSharedCustomRulesRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteSharedCustomRulesRequestHeader) GoString() string {
  return s.String()
}

func (s *DeleteSharedCustomRulesRequestHeader) SetServiceType(v string) *DeleteSharedCustomRulesRequestHeader {
  s.ServiceType = &v
  return s
}

type DeleteSharedCustomRulesPaths struct {
}

func (s DeleteSharedCustomRulesPaths) String() string {
  return tea.Prettify(s)
}

func (s DeleteSharedCustomRulesPaths) GoString() string {
  return s.String()
}

type DeleteSharedCustomRulesParameters struct {
}

func (s DeleteSharedCustomRulesParameters) String() string {
  return tea.Prettify(s)
}

func (s DeleteSharedCustomRulesParameters) GoString() string {
  return s.String()
}

type DeleteSharedCustomRulesResponse struct {
  // {"dictionary":"belong=WAAP-MS-Ext|dict=waap_retCodeEnum","en":"Please refer to the error code for exceptions.","zh_CN":"请参照错误码。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Description.","zh_CN":"描述信息。"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"Data.","zh_CN":"出参数据。"}
  Data *bool `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s DeleteSharedCustomRulesResponse) String() string {
  return tea.Prettify(s)
}

func (s DeleteSharedCustomRulesResponse) GoString() string {
  return s.String()
}

func (s *DeleteSharedCustomRulesResponse) SetCode(v string) *DeleteSharedCustomRulesResponse {
  s.Code = &v
  return s
}

func (s *DeleteSharedCustomRulesResponse) SetMsg(v string) *DeleteSharedCustomRulesResponse {
  s.Msg = &v
  return s
}

func (s *DeleteSharedCustomRulesResponse) SetData(v bool) *DeleteSharedCustomRulesResponse {
  s.Data = &v
  return s
}

type DeleteSharedCustomRulesResponseHeader struct {
}

func (s DeleteSharedCustomRulesResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteSharedCustomRulesResponseHeader) GoString() string {
  return s.String()
}




type DeleteShareCustomizeBotsRequest struct {
  // {"en":"Rule ID list.","zh_CN":"规则ID列表。"}
  IdList []*string `json:"idList,omitempty" xml:"idList,omitempty" require:"true" type:"Repeated"`
}

func (s DeleteShareCustomizeBotsRequest) String() string {
  return tea.Prettify(s)
}

func (s DeleteShareCustomizeBotsRequest) GoString() string {
  return s.String()
}

func (s *DeleteShareCustomizeBotsRequest) SetIdList(v []*string) *DeleteShareCustomizeBotsRequest {
  s.IdList = v
  return s
}

type DeleteShareCustomizeBotsRequestHeader struct {
  // {"en":"Security service type. Please enter a specific service type, if you purchase multiple security services.","zh_CN":"安全服务类型。有使用多个不同的安全服务时，需要填写具体的服务类型。"}
  ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty"`
}

func (s DeleteShareCustomizeBotsRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteShareCustomizeBotsRequestHeader) GoString() string {
  return s.String()
}

func (s *DeleteShareCustomizeBotsRequestHeader) SetServiceType(v string) *DeleteShareCustomizeBotsRequestHeader {
  s.ServiceType = &v
  return s
}

type DeleteShareCustomizeBotsPaths struct {
}

func (s DeleteShareCustomizeBotsPaths) String() string {
  return tea.Prettify(s)
}

func (s DeleteShareCustomizeBotsPaths) GoString() string {
  return s.String()
}

type DeleteShareCustomizeBotsParameters struct {
}

func (s DeleteShareCustomizeBotsParameters) String() string {
  return tea.Prettify(s)
}

func (s DeleteShareCustomizeBotsParameters) GoString() string {
  return s.String()
}

type DeleteShareCustomizeBotsResponse struct {
  // {"dictionary":"belong=WAAP-MS-Ext|dict=waap_retCodeEnum","en":"Please refer to the error code for exceptions.","zh_CN":"请参照错误码。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Description.","zh_CN":"描述信息。"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"Data.","zh_CN":"出参数据。"}
  Data *bool `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s DeleteShareCustomizeBotsResponse) String() string {
  return tea.Prettify(s)
}

func (s DeleteShareCustomizeBotsResponse) GoString() string {
  return s.String()
}

func (s *DeleteShareCustomizeBotsResponse) SetCode(v string) *DeleteShareCustomizeBotsResponse {
  s.Code = &v
  return s
}

func (s *DeleteShareCustomizeBotsResponse) SetMsg(v string) *DeleteShareCustomizeBotsResponse {
  s.Msg = &v
  return s
}

func (s *DeleteShareCustomizeBotsResponse) SetData(v bool) *DeleteShareCustomizeBotsResponse {
  s.Data = &v
  return s
}

type DeleteShareCustomizeBotsResponseHeader struct {
}

func (s DeleteShareCustomizeBotsResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteShareCustomizeBotsResponseHeader) GoString() string {
  return s.String()
}




type ListSharedWAFRuleExceptionsRequest struct {
}

func (s ListSharedWAFRuleExceptionsRequest) String() string {
  return tea.Prettify(s)
}

func (s ListSharedWAFRuleExceptionsRequest) GoString() string {
  return s.String()
}

type ListSharedWAFRuleExceptionsRequestHeader struct {
  // {"en":"Security service type. Please enter a specific service type, if you purchase multiple security services.","zh_CN":"安全服务类型。有使用多个不同的安全服务时，需要填写具体的服务类型。"}
  ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty"`
}

func (s ListSharedWAFRuleExceptionsRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ListSharedWAFRuleExceptionsRequestHeader) GoString() string {
  return s.String()
}

func (s *ListSharedWAFRuleExceptionsRequestHeader) SetServiceType(v string) *ListSharedWAFRuleExceptionsRequestHeader {
  s.ServiceType = &v
  return s
}

type ListSharedWAFRuleExceptionsPaths struct {
}

func (s ListSharedWAFRuleExceptionsPaths) String() string {
  return tea.Prettify(s)
}

func (s ListSharedWAFRuleExceptionsPaths) GoString() string {
  return s.String()
}

type ListSharedWAFRuleExceptionsParameters struct {
}

func (s ListSharedWAFRuleExceptionsParameters) String() string {
  return tea.Prettify(s)
}

func (s ListSharedWAFRuleExceptionsParameters) GoString() string {
  return s.String()
}

type ListSharedWAFRuleExceptionsResponse struct {
  // {"dictionary":"belong=WAAP-MS-Ext|dict=waap_retCodeEnum","en":"Please refer to the error code for exceptions.","zh_CN":"请参照错误码。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Description.","zh_CN":"描述信息。"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"Shared configuration WAF rule exception list.","zh_CN":"共享配置WAF规则例外列表。"}
  Data []*ListSharedWAFRuleExceptionsResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s ListSharedWAFRuleExceptionsResponse) String() string {
  return tea.Prettify(s)
}

func (s ListSharedWAFRuleExceptionsResponse) GoString() string {
  return s.String()
}

func (s *ListSharedWAFRuleExceptionsResponse) SetCode(v string) *ListSharedWAFRuleExceptionsResponse {
  s.Code = &v
  return s
}

func (s *ListSharedWAFRuleExceptionsResponse) SetMsg(v string) *ListSharedWAFRuleExceptionsResponse {
  s.Msg = &v
  return s
}

func (s *ListSharedWAFRuleExceptionsResponse) SetData(v []*ListSharedWAFRuleExceptionsResponseData) *ListSharedWAFRuleExceptionsResponse {
  s.Data = v
  return s
}

type ListSharedWAFRuleExceptionsResponseData struct     {
  // {"en":"Rule exception ID.","zh_CN":"规则例外ID。"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {"en":"Exception name,maximum 50 characters.\nDoes not support special characters and spaces.","zh_CN":"例外名称，最多50个字符。\n不支持特殊字符和空格。"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"Exception description,maximum 200 characters..","zh_CN":"例外描述，最多200个字符。"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"Matching conditions.\nip: IP\npath: Path\nuri: URI\nurlParamName: URI Parameter Name\nurlParamValue: URI Parameter Value\nuserAgent: User Agent\nhttpHeaderName: Request Header Name\nhttpHeaderValue: Request Header Value\ncookie: Cookie\nbody: Body\nbodyParamName: Body Parameter Name\nbodyParamValue: Body Parameter Value","zh_CN":"匹配条件。\nip：IP\npath：路径\nuri：URI\nurlParamName：URI参数名\nurlParamValue：URI参数值\nuserAgent：User Agent\nhttpHeaderName：请求头部名称\nhttpHeaderValue：请求头部值\ncookie：Cookie\nbody：Body\nbodyParamName：Body参数名\nbodyParamValue：Body参数值","exampleValue":"ip,path,uri,urlParamName,urlParamValue,userAgent,httpHeaderName,httpHeaderValue,cookie,body,bodyParamName,bodyParamValue"}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
  // {"en":"Match type,IP can only be EQUAL.\nEQUAL: Equal\nCONTAIN: Contains\nREGEX: Regular match","zh_CN":"匹配类型，IP只能是等于。\nEQUAL：等于\nCONTAIN：包含\nREGEX：正则匹配","exampleValue":"EQUAL,CONTAIN,REGEX"}
  MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
  // {"en":"Rule exceptions.\nWhen matchType=EQUAL, case-sensitive.","zh_CN":"规则例外内容。\nmatchType=EQUAL时，大小写敏感。"}
  ContentList []*string `json:"contentList,omitempty" xml:"contentList,omitempty" require:"true" type:"Repeated"`
  // {"en":"Creator login.","zh_CN":"创建者登录名。"}
  Creator *string `json:"creator,omitempty" xml:"creator,omitempty" require:"true"`
  // {"en":"Creation time.","zh_CN":"创建时间。"}
  CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty" require:"true"`
  // {"en":"Update time.","zh_CN":"更新时间。"}
  UpdateTime *string `json:"updateTime,omitempty" xml:"updateTime,omitempty" require:"true"`
  // {"en":"List of associated hostnames and rule IDs.","zh_CN":"关联的域名以及规则ID列表。"}
  DomainRuleList []*ListSharedWAFRuleExceptionsResponseDataDomainRuleList `json:"domainRuleList,omitempty" xml:"domainRuleList,omitempty" require:"true" type:"Repeated"`
}

func (s ListSharedWAFRuleExceptionsResponseData) String() string {
  return tea.Prettify(s)
}

func (s ListSharedWAFRuleExceptionsResponseData) GoString() string {
  return s.String()
}

func (s *ListSharedWAFRuleExceptionsResponseData) SetId(v string) *ListSharedWAFRuleExceptionsResponseData {
  s.Id = &v
  return s
}

func (s *ListSharedWAFRuleExceptionsResponseData) SetName(v string) *ListSharedWAFRuleExceptionsResponseData {
  s.Name = &v
  return s
}

func (s *ListSharedWAFRuleExceptionsResponseData) SetMsg(v string) *ListSharedWAFRuleExceptionsResponseData {
  s.Msg = &v
  return s
}

func (s *ListSharedWAFRuleExceptionsResponseData) SetType(v string) *ListSharedWAFRuleExceptionsResponseData {
  s.Type = &v
  return s
}

func (s *ListSharedWAFRuleExceptionsResponseData) SetMatchType(v string) *ListSharedWAFRuleExceptionsResponseData {
  s.MatchType = &v
  return s
}

func (s *ListSharedWAFRuleExceptionsResponseData) SetContentList(v []*string) *ListSharedWAFRuleExceptionsResponseData {
  s.ContentList = v
  return s
}

func (s *ListSharedWAFRuleExceptionsResponseData) SetCreator(v string) *ListSharedWAFRuleExceptionsResponseData {
  s.Creator = &v
  return s
}

func (s *ListSharedWAFRuleExceptionsResponseData) SetCreateTime(v string) *ListSharedWAFRuleExceptionsResponseData {
  s.CreateTime = &v
  return s
}

func (s *ListSharedWAFRuleExceptionsResponseData) SetUpdateTime(v string) *ListSharedWAFRuleExceptionsResponseData {
  s.UpdateTime = &v
  return s
}

func (s *ListSharedWAFRuleExceptionsResponseData) SetDomainRuleList(v []*ListSharedWAFRuleExceptionsResponseDataDomainRuleList) *ListSharedWAFRuleExceptionsResponseData {
  s.DomainRuleList = v
  return s
}

type ListSharedWAFRuleExceptionsResponseDataDomainRuleList struct     {
  // {"en":"Domain.","zh_CN":"域名。"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {"en":"WAF rule ID list.","zh_CN":"WAF规则ID列表。"}
  RuleIdList []*int `json:"ruleIdList,omitempty" xml:"ruleIdList,omitempty" require:"true" type:"Repeated"`
}

func (s ListSharedWAFRuleExceptionsResponseDataDomainRuleList) String() string {
  return tea.Prettify(s)
}

func (s ListSharedWAFRuleExceptionsResponseDataDomainRuleList) GoString() string {
  return s.String()
}

func (s *ListSharedWAFRuleExceptionsResponseDataDomainRuleList) SetDomain(v string) *ListSharedWAFRuleExceptionsResponseDataDomainRuleList {
  s.Domain = &v
  return s
}

func (s *ListSharedWAFRuleExceptionsResponseDataDomainRuleList) SetRuleIdList(v []*int) *ListSharedWAFRuleExceptionsResponseDataDomainRuleList {
  s.RuleIdList = v
  return s
}

type ListSharedWAFRuleExceptionsResponseHeader struct {
}

func (s ListSharedWAFRuleExceptionsResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ListSharedWAFRuleExceptionsResponseHeader) GoString() string {
  return s.String()
}




type QueryAppApiExceptionFeatureReferencedDomainsRequest struct {
  // {"en":"Rule Id, Only supports single rule queries","zh_CN":"规则ID，仅支持单个规则查询"}
  IdList []*string `json:"idList,omitempty" xml:"idList,omitempty" require:"true" type:"Repeated"`
}

func (s QueryAppApiExceptionFeatureReferencedDomainsRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryAppApiExceptionFeatureReferencedDomainsRequest) GoString() string {
  return s.String()
}

func (s *QueryAppApiExceptionFeatureReferencedDomainsRequest) SetIdList(v []*string) *QueryAppApiExceptionFeatureReferencedDomainsRequest {
  s.IdList = v
  return s
}

type QueryAppApiExceptionFeatureReferencedDomainsRequestHeader struct {
  // {"en":"Security service type. Please enter a specific service type, if you purchase multiple security services.","zh_CN":"安全服务类型。有使用多个不同的安全服务时，需要填写具体的服务类型。"}
  ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty"`
}

func (s QueryAppApiExceptionFeatureReferencedDomainsRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryAppApiExceptionFeatureReferencedDomainsRequestHeader) GoString() string {
  return s.String()
}

func (s *QueryAppApiExceptionFeatureReferencedDomainsRequestHeader) SetServiceType(v string) *QueryAppApiExceptionFeatureReferencedDomainsRequestHeader {
  s.ServiceType = &v
  return s
}

type QueryAppApiExceptionFeatureReferencedDomainsPaths struct {
}

func (s QueryAppApiExceptionFeatureReferencedDomainsPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryAppApiExceptionFeatureReferencedDomainsPaths) GoString() string {
  return s.String()
}

type QueryAppApiExceptionFeatureReferencedDomainsParameters struct {
}

func (s QueryAppApiExceptionFeatureReferencedDomainsParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryAppApiExceptionFeatureReferencedDomainsParameters) GoString() string {
  return s.String()
}

type QueryAppApiExceptionFeatureReferencedDomainsResponse struct {
  // {"dictionary":"belong=WAAP-MS-Ext|dict=waap_retCodeEnum","en":"Please refer to the error code for exceptions.","zh_CN":"请参照错误码。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Description.","zh_CN":"描述信息。"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"Data.","zh_CN":"出参数据。"}
  Data []*string `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s QueryAppApiExceptionFeatureReferencedDomainsResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryAppApiExceptionFeatureReferencedDomainsResponse) GoString() string {
  return s.String()
}

func (s *QueryAppApiExceptionFeatureReferencedDomainsResponse) SetCode(v string) *QueryAppApiExceptionFeatureReferencedDomainsResponse {
  s.Code = &v
  return s
}

func (s *QueryAppApiExceptionFeatureReferencedDomainsResponse) SetMsg(v string) *QueryAppApiExceptionFeatureReferencedDomainsResponse {
  s.Msg = &v
  return s
}

func (s *QueryAppApiExceptionFeatureReferencedDomainsResponse) SetData(v []*string) *QueryAppApiExceptionFeatureReferencedDomainsResponse {
  s.Data = v
  return s
}

type QueryAppApiExceptionFeatureReferencedDomainsResponseHeader struct {
}

func (s QueryAppApiExceptionFeatureReferencedDomainsResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryAppApiExceptionFeatureReferencedDomainsResponseHeader) GoString() string {
  return s.String()
}




type DeleteSharedWhitelistRuleRequest struct {
  // {"en":"Rule ID List.","zh_CN":"规则ID列表。"}
  IdList []*string `json:"idList,omitempty" xml:"idList,omitempty" require:"true" type:"Repeated"`
}

func (s DeleteSharedWhitelistRuleRequest) String() string {
  return tea.Prettify(s)
}

func (s DeleteSharedWhitelistRuleRequest) GoString() string {
  return s.String()
}

func (s *DeleteSharedWhitelistRuleRequest) SetIdList(v []*string) *DeleteSharedWhitelistRuleRequest {
  s.IdList = v
  return s
}

type DeleteSharedWhitelistRuleRequestHeader struct {
  // {"dictionary":"belong=WAAP-MS-Ext|dict=waap_serviceType","en":"Security service type. Please enter a specific service type, if you purchase multiple security services.","zh_CN":"安全服务类型。有使用多个不同的安全服务时，需要填写具体的服务类型。"}
  ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty"`
}

func (s DeleteSharedWhitelistRuleRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteSharedWhitelistRuleRequestHeader) GoString() string {
  return s.String()
}

func (s *DeleteSharedWhitelistRuleRequestHeader) SetServiceType(v string) *DeleteSharedWhitelistRuleRequestHeader {
  s.ServiceType = &v
  return s
}

type DeleteSharedWhitelistRulePaths struct {
}

func (s DeleteSharedWhitelistRulePaths) String() string {
  return tea.Prettify(s)
}

func (s DeleteSharedWhitelistRulePaths) GoString() string {
  return s.String()
}

type DeleteSharedWhitelistRuleParameters struct {
}

func (s DeleteSharedWhitelistRuleParameters) String() string {
  return tea.Prettify(s)
}

func (s DeleteSharedWhitelistRuleParameters) GoString() string {
  return s.String()
}

type DeleteSharedWhitelistRuleResponse struct {
  // {"dictionary":"belong=WAAP-MS-Ext|dict=waap_retCodeEnum","en":"Please refer to the error code for exceptions.","zh_CN":"请参照错误码。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Description.","zh_CN":"描述信息。"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
}

func (s DeleteSharedWhitelistRuleResponse) String() string {
  return tea.Prettify(s)
}

func (s DeleteSharedWhitelistRuleResponse) GoString() string {
  return s.String()
}

func (s *DeleteSharedWhitelistRuleResponse) SetCode(v string) *DeleteSharedWhitelistRuleResponse {
  s.Code = &v
  return s
}

func (s *DeleteSharedWhitelistRuleResponse) SetMsg(v string) *DeleteSharedWhitelistRuleResponse {
  s.Msg = &v
  return s
}

type DeleteSharedWhitelistRuleResponseHeader struct {
}

func (s DeleteSharedWhitelistRuleResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteSharedWhitelistRuleResponseHeader) GoString() string {
  return s.String()
}




type UpdateShareConfigurationsAppApiExceptionFeatureRequest struct {
  // {"en":"Rule ID","zh_CN":"规则ID"}
  RuleId *string `json:"ruleId,omitempty" xml:"ruleId,omitempty" require:"true"`
  // {"en":"Rule name, duplicate not allowed.","zh_CN":"规则名称，不允许重复"}
  RuleName *string `json:"ruleName,omitempty" xml:"ruleName,omitempty" require:"true"`
  // {"en":"Type enumeration, for example:\nNATIVE_ APP        remarks: Native App\nCallback_ API       remarks: Callback API\nOther_ API         remarks: Other program API\nHYBRID_ APP     remarks: Hybrid APP","zh_CN":"类型枚举，例如：\nNATIVE_APP          备注：原生APP\nCALLBACK_API      备注：回调API\nOTHER_API            备注：其他程序API\nHYBRID_APP          备注：混合APP","exampleValue":"NATIVE_APP,CALLBACK_API,OTHER_API,HYBRID_APP"}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
  // {"en":"Characteristic.","zh_CN":"特征。"}
  Config *UpdateShareConfigurationsAppApiExceptionFeatureRequestConfig `json:"config,omitempty" xml:"config,omitempty" require:"true" type:"Struct"`
}

func (s UpdateShareConfigurationsAppApiExceptionFeatureRequest) String() string {
  return tea.Prettify(s)
}

func (s UpdateShareConfigurationsAppApiExceptionFeatureRequest) GoString() string {
  return s.String()
}

func (s *UpdateShareConfigurationsAppApiExceptionFeatureRequest) SetRuleId(v string) *UpdateShareConfigurationsAppApiExceptionFeatureRequest {
  s.RuleId = &v
  return s
}

func (s *UpdateShareConfigurationsAppApiExceptionFeatureRequest) SetRuleName(v string) *UpdateShareConfigurationsAppApiExceptionFeatureRequest {
  s.RuleName = &v
  return s
}

func (s *UpdateShareConfigurationsAppApiExceptionFeatureRequest) SetType(v string) *UpdateShareConfigurationsAppApiExceptionFeatureRequest {
  s.Type = &v
  return s
}

func (s *UpdateShareConfigurationsAppApiExceptionFeatureRequest) SetConfig(v *UpdateShareConfigurationsAppApiExceptionFeatureRequestConfig) *UpdateShareConfigurationsAppApiExceptionFeatureRequest {
  s.Config = v
  return s
}

type UpdateShareConfigurationsAppApiExceptionFeatureRequestConfig struct {
  // {"en":"Matching condition name.\nPATH         Path\nURI            URI\nUA             User-Agent\nREFERER    Referer\nHEADER    Request Header","zh_CN":"匹配条件名称。\nPATH                       路径\nURI                          URI\nUA                           User-Agent\nREFERER                  Referer\nHEAD                      Request Header","exampleValue":"PATH,URI,UA,REFERER,HEAD"}
  Condition *string `json:"condition,omitempty" xml:"condition,omitempty" require:"true"`
  // {"en":"Matching condition function.\nCONTAIN                  contains\nNOT_CONTAIN         does not contain\nEQUAL                      equals\nNOT_EQUAL             does not equal\nEMPTY                      does not exist or has no value(When the condition is a path or URI, this value is not optional)\nREGEX                      regex match","zh_CN":"匹配条件函数。\nCONTAIN                  包含\nNOT_CONTAIN         不包含\nEQUAL                      等于\nNOT_EQUAL             不等于\nEMPTY                      为空或不存在（当condition为路径或者URI时，该值不可选）\nREGEX                       匹配正则","exampleValue":"CONTAIN,NOT_CONTAIN,EQUAL,NOT_EQUAL,EMPTY,REGEX"}
  MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
  // {"en":"Condition value list, When matchType is EMPTY, the value is empty;  When matchType is REGEX, only one item is set","zh_CN":"条件值列表，当matchType为EMPTY时，该值为空；当matchType为REGEX时，只能传一条"}
  ConditionValue []*string `json:"conditionValue,omitempty" xml:"conditionValue,omitempty" require:"true" type:"Repeated"`
  // {"en":"Head value, If condition=HEAD, then this field is mandatory, Otherwise, it is empty.","zh_CN":"头部值，如果condition为HEAD，则该字段必填，否则为空。"}
  Key *string `json:"key,omitempty" xml:"key,omitempty"`
}

func (s UpdateShareConfigurationsAppApiExceptionFeatureRequestConfig) String() string {
  return tea.Prettify(s)
}

func (s UpdateShareConfigurationsAppApiExceptionFeatureRequestConfig) GoString() string {
  return s.String()
}

func (s *UpdateShareConfigurationsAppApiExceptionFeatureRequestConfig) SetCondition(v string) *UpdateShareConfigurationsAppApiExceptionFeatureRequestConfig {
  s.Condition = &v
  return s
}

func (s *UpdateShareConfigurationsAppApiExceptionFeatureRequestConfig) SetMatchType(v string) *UpdateShareConfigurationsAppApiExceptionFeatureRequestConfig {
  s.MatchType = &v
  return s
}

func (s *UpdateShareConfigurationsAppApiExceptionFeatureRequestConfig) SetConditionValue(v []*string) *UpdateShareConfigurationsAppApiExceptionFeatureRequestConfig {
  s.ConditionValue = v
  return s
}

func (s *UpdateShareConfigurationsAppApiExceptionFeatureRequestConfig) SetKey(v string) *UpdateShareConfigurationsAppApiExceptionFeatureRequestConfig {
  s.Key = &v
  return s
}

type UpdateShareConfigurationsAppApiExceptionFeatureRequestHeader struct {
  // {"en":"Security service type. Please enter a specific service type, if you purchase multiple security services.","zh_CN":"安全服务类型。有使用多个不同的安全服务时，需要填写具体的服务类型。"}
  ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty"`
}

func (s UpdateShareConfigurationsAppApiExceptionFeatureRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s UpdateShareConfigurationsAppApiExceptionFeatureRequestHeader) GoString() string {
  return s.String()
}

func (s *UpdateShareConfigurationsAppApiExceptionFeatureRequestHeader) SetServiceType(v string) *UpdateShareConfigurationsAppApiExceptionFeatureRequestHeader {
  s.ServiceType = &v
  return s
}

type UpdateShareConfigurationsAppApiExceptionFeaturePaths struct {
}

func (s UpdateShareConfigurationsAppApiExceptionFeaturePaths) String() string {
  return tea.Prettify(s)
}

func (s UpdateShareConfigurationsAppApiExceptionFeaturePaths) GoString() string {
  return s.String()
}

type UpdateShareConfigurationsAppApiExceptionFeatureParameters struct {
}

func (s UpdateShareConfigurationsAppApiExceptionFeatureParameters) String() string {
  return tea.Prettify(s)
}

func (s UpdateShareConfigurationsAppApiExceptionFeatureParameters) GoString() string {
  return s.String()
}

type UpdateShareConfigurationsAppApiExceptionFeatureResponse struct {
  // {"dictionary":"belong=WAAP-MS-Ext|dict=waap_retCodeEnum","en":"Please refer to the error code for exceptions.","zh_CN":"请参照错误码。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Description.","zh_CN":"描述信息。"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
}

func (s UpdateShareConfigurationsAppApiExceptionFeatureResponse) String() string {
  return tea.Prettify(s)
}

func (s UpdateShareConfigurationsAppApiExceptionFeatureResponse) GoString() string {
  return s.String()
}

func (s *UpdateShareConfigurationsAppApiExceptionFeatureResponse) SetCode(v string) *UpdateShareConfigurationsAppApiExceptionFeatureResponse {
  s.Code = &v
  return s
}

func (s *UpdateShareConfigurationsAppApiExceptionFeatureResponse) SetMsg(v string) *UpdateShareConfigurationsAppApiExceptionFeatureResponse {
  s.Msg = &v
  return s
}

type UpdateShareConfigurationsAppApiExceptionFeatureResponseHeader struct {
}

func (s UpdateShareConfigurationsAppApiExceptionFeatureResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s UpdateShareConfigurationsAppApiExceptionFeatureResponseHeader) GoString() string {
  return s.String()
}




type DeleteSharedRateLimitingRuleRequest struct {
  // {"en":"Rule ID List.","zh_CN":"规则ID列表。"}
  Ids []*string `json:"ids,omitempty" xml:"ids,omitempty" require:"true" type:"Repeated"`
}

func (s DeleteSharedRateLimitingRuleRequest) String() string {
  return tea.Prettify(s)
}

func (s DeleteSharedRateLimitingRuleRequest) GoString() string {
  return s.String()
}

func (s *DeleteSharedRateLimitingRuleRequest) SetIds(v []*string) *DeleteSharedRateLimitingRuleRequest {
  s.Ids = v
  return s
}

type DeleteSharedRateLimitingRuleRequestHeader struct {
  // {"en":"Security service type. Please enter a specific service type, if you purchase multiple security services.","zh_CN":"安全服务类型。有使用多个不同的安全服务时，需要填写具体的服务类型。"}
  ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty"`
}

func (s DeleteSharedRateLimitingRuleRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteSharedRateLimitingRuleRequestHeader) GoString() string {
  return s.String()
}

func (s *DeleteSharedRateLimitingRuleRequestHeader) SetServiceType(v string) *DeleteSharedRateLimitingRuleRequestHeader {
  s.ServiceType = &v
  return s
}

type DeleteSharedRateLimitingRulePaths struct {
}

func (s DeleteSharedRateLimitingRulePaths) String() string {
  return tea.Prettify(s)
}

func (s DeleteSharedRateLimitingRulePaths) GoString() string {
  return s.String()
}

type DeleteSharedRateLimitingRuleParameters struct {
}

func (s DeleteSharedRateLimitingRuleParameters) String() string {
  return tea.Prettify(s)
}

func (s DeleteSharedRateLimitingRuleParameters) GoString() string {
  return s.String()
}

type DeleteSharedRateLimitingRuleResponse struct {
  // {"dictionary":"belong=WAAP-MS-Ext|dict=waap_retCodeEnum","en":"Please refer to the error code for exceptions.","zh_CN":"请参照错误码。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Description.","zh_CN":"描述信息。"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
}

func (s DeleteSharedRateLimitingRuleResponse) String() string {
  return tea.Prettify(s)
}

func (s DeleteSharedRateLimitingRuleResponse) GoString() string {
  return s.String()
}

func (s *DeleteSharedRateLimitingRuleResponse) SetCode(v string) *DeleteSharedRateLimitingRuleResponse {
  s.Code = &v
  return s
}

func (s *DeleteSharedRateLimitingRuleResponse) SetMsg(v string) *DeleteSharedRateLimitingRuleResponse {
  s.Msg = &v
  return s
}

type DeleteSharedRateLimitingRuleResponseHeader struct {
}

func (s DeleteSharedRateLimitingRuleResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteSharedRateLimitingRuleResponseHeader) GoString() string {
  return s.String()
}




type UpdateSharedCustomRulesRequest struct {
  // {"en":"Matching conditions. Except for header conditions, there can only be at most one record per match type under each type of condition.","zh_CN":"匹配条件。除了请求头条件，其它类型的条件下一种匹配类型最多只能有一条记录。"}
  Condition *UpdateSharedCustomRulesRequestCondition `json:"condition,omitempty" xml:"condition,omitempty" type:"Struct"`
  // {"en":"Action.\nNO_USE: Not Used\nLOG: Log\nDELAY: Delay\nBLOCK: Deny\nRESET: Reset Connection\nCUSTOM_ACTION_ID:Fill in the custom action id of the corresponding action","zh_CN":"处理动作。\nNO_USE：不使用\nLOG：监控\nDELAY：延迟响应\nBLOCK：拦截\nRESET：断开连接\nCUSTOM_ACTION_ID：传入对应的自定义响应动作的ID","exampleValue":"NO_USE,LOG,DELAY,BLOCK,RESET"}
  Act *string `json:"act,omitempty" xml:"act,omitempty"`
  // {"en":"Rule Name, maximum 100 characters.Does not support # and &.","zh_CN":"规则名称，最多100个字符。不支持 # 和 &。"}
  RuleName *string `json:"ruleName,omitempty" xml:"ruleName,omitempty"`
  // {"en":"Description, maximum 1000 characters.","zh_CN":"规则描述，最多1000个字符。"}
  Description *string `json:"description,omitempty" xml:"description,omitempty"`
  // {"en":"Rule ID.","zh_CN":"规则ID。"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
}

func (s UpdateSharedCustomRulesRequest) String() string {
  return tea.Prettify(s)
}

func (s UpdateSharedCustomRulesRequest) GoString() string {
  return s.String()
}

func (s *UpdateSharedCustomRulesRequest) SetCondition(v *UpdateSharedCustomRulesRequestCondition) *UpdateSharedCustomRulesRequest {
  s.Condition = v
  return s
}

func (s *UpdateSharedCustomRulesRequest) SetAct(v string) *UpdateSharedCustomRulesRequest {
  s.Act = &v
  return s
}

func (s *UpdateSharedCustomRulesRequest) SetRuleName(v string) *UpdateSharedCustomRulesRequest {
  s.RuleName = &v
  return s
}

func (s *UpdateSharedCustomRulesRequest) SetDescription(v string) *UpdateSharedCustomRulesRequest {
  s.Description = &v
  return s
}

func (s *UpdateSharedCustomRulesRequest) SetId(v string) *UpdateSharedCustomRulesRequest {
  s.Id = &v
  return s
}

type UpdateSharedCustomRulesRequestCondition struct {
  // {"en":"Request Method.\nWhen the business scenario is API,this matching condition is not supported.","zh_CN":"请求方法，匹配类型不可重复。\n当业务场景为API业务时不支持此匹配条件。"}
  MethodConditions []*UpdateSharedCustomRulesRequestConditionMethodConditions `json:"methodConditions,omitempty" xml:"methodConditions,omitempty" type:"Repeated"`
  // {"en":"Geo, match type cannot be repeated.","zh_CN":"区域，匹配类型不可重复。"}
  AreaConditions []*UpdateSharedCustomRulesRequestConditionAreaConditions `json:"areaConditions,omitempty" xml:"areaConditions,omitempty" type:"Repeated"`
  // {"en":"IP/CIDR, match type cannot be repeated.","zh_CN":"IP/IP段，匹配类型不可重复。"}
  IpOrIpsConditions []*UpdateSharedCustomRulesRequestConditionIpOrIpsConditions `json:"ipOrIpsConditions,omitempty" xml:"ipOrIpsConditions,omitempty" type:"Repeated"`
  // {"en":"URI, match type cannot be repeated.\nWhen the business scenario is API, this matching condition is not supported.","zh_CN":"URI，匹配类型不可重复。\n当业务场景为API业务时不支持此匹配条件。"}
  UriConditions []*UpdateSharedCustomRulesRequestConditionUriConditions `json:"uriConditions,omitempty" xml:"uriConditions,omitempty" type:"Repeated"`
  // {"en":"Path, match type cannot be repeated.\nWhen the business scenario is API, this matching condition is not supported.","zh_CN":"路径，匹配类型不可重复。\n当业务场景为API业务时不支持此匹配条件。"}
  PathConditions []*UpdateSharedCustomRulesRequestConditionPathConditions `json:"pathConditions,omitempty" xml:"pathConditions,omitempty" type:"Repeated"`
  // {"en":"User Agent, match type cannot be repeated.","zh_CN":"User-Agent，匹配类型不可重复。"}
  UaConditions []*UpdateSharedCustomRulesRequestConditionUaConditions `json:"uaConditions,omitempty" xml:"uaConditions,omitempty" type:"Repeated"`
  // {"en":"Request Header, match type can be repeated.","zh_CN":"请求头，匹配类型可重复。"}
  HeaderConditions []*UpdateSharedCustomRulesRequestConditionHeaderConditions `json:"headerConditions,omitempty" xml:"headerConditions,omitempty" type:"Repeated"`
  // {"en":"Referer, match type cannot be repeated.","zh_CN":"Referer，匹配类型不可重复。"}
  RefererConditions []*UpdateSharedCustomRulesRequestConditionRefererConditions `json:"refererConditions,omitempty" xml:"refererConditions,omitempty" type:"Repeated"`
  // {"en":"JA3 Fingerprint, match type cannot be repeated.","zh_CN":"JA3指纹，匹配类型不可重复。"}
  Ja3Conditions []*UpdateSharedCustomRulesRequestConditionJa3Conditions `json:"ja3Conditions,omitempty" xml:"ja3Conditions,omitempty" type:"Repeated"`
  // {"en":"JA4 Fingerprint, match type cannot be repeated.","zh_CN":"JA4指纹，匹配类型不可重复。"}
  Ja4Conditions []*UpdateSharedCustomRulesRequestConditionJa4Conditions `json:"ja4Conditions,omitempty" xml:"ja4Conditions,omitempty" type:"Repeated"`
}

func (s UpdateSharedCustomRulesRequestCondition) String() string {
  return tea.Prettify(s)
}

func (s UpdateSharedCustomRulesRequestCondition) GoString() string {
  return s.String()
}

func (s *UpdateSharedCustomRulesRequestCondition) SetMethodConditions(v []*UpdateSharedCustomRulesRequestConditionMethodConditions) *UpdateSharedCustomRulesRequestCondition {
  s.MethodConditions = v
  return s
}

func (s *UpdateSharedCustomRulesRequestCondition) SetAreaConditions(v []*UpdateSharedCustomRulesRequestConditionAreaConditions) *UpdateSharedCustomRulesRequestCondition {
  s.AreaConditions = v
  return s
}

func (s *UpdateSharedCustomRulesRequestCondition) SetIpOrIpsConditions(v []*UpdateSharedCustomRulesRequestConditionIpOrIpsConditions) *UpdateSharedCustomRulesRequestCondition {
  s.IpOrIpsConditions = v
  return s
}

func (s *UpdateSharedCustomRulesRequestCondition) SetUriConditions(v []*UpdateSharedCustomRulesRequestConditionUriConditions) *UpdateSharedCustomRulesRequestCondition {
  s.UriConditions = v
  return s
}

func (s *UpdateSharedCustomRulesRequestCondition) SetPathConditions(v []*UpdateSharedCustomRulesRequestConditionPathConditions) *UpdateSharedCustomRulesRequestCondition {
  s.PathConditions = v
  return s
}

func (s *UpdateSharedCustomRulesRequestCondition) SetUaConditions(v []*UpdateSharedCustomRulesRequestConditionUaConditions) *UpdateSharedCustomRulesRequestCondition {
  s.UaConditions = v
  return s
}

func (s *UpdateSharedCustomRulesRequestCondition) SetHeaderConditions(v []*UpdateSharedCustomRulesRequestConditionHeaderConditions) *UpdateSharedCustomRulesRequestCondition {
  s.HeaderConditions = v
  return s
}

func (s *UpdateSharedCustomRulesRequestCondition) SetRefererConditions(v []*UpdateSharedCustomRulesRequestConditionRefererConditions) *UpdateSharedCustomRulesRequestCondition {
  s.RefererConditions = v
  return s
}

func (s *UpdateSharedCustomRulesRequestCondition) SetJa3Conditions(v []*UpdateSharedCustomRulesRequestConditionJa3Conditions) *UpdateSharedCustomRulesRequestCondition {
  s.Ja3Conditions = v
  return s
}

func (s *UpdateSharedCustomRulesRequestCondition) SetJa4Conditions(v []*UpdateSharedCustomRulesRequestConditionJa4Conditions) *UpdateSharedCustomRulesRequestCondition {
  s.Ja4Conditions = v
  return s
}

type UpdateSharedCustomRulesRequestConditionMethodConditions struct     {
  // {"en":"Match type.\nEQUAL:Equals\nNOT_EQUAL:Does not equal","zh_CN":"匹配类型。\nEQUAL：等于\nNOT_EQUAL：不等于","exampleValue":"EQUAL,NOT_EQUAL"}
  MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
  // {"en":"Request method.\nSupported values: GET/POST/DELETE/PUT/HEAD/OPTIONS/COPY.","zh_CN":"请求方法。\n支持的值：GET/POST/DELETE/PUT/HEAD/OPTIONS/COPY。","exampleValue":"GET,POST,DELETE,PUT,HEAD,OPTIONS,COPY"}
  RequestMethod []*string `json:"requestMethod,omitempty" xml:"requestMethod,omitempty" require:"true" type:"Repeated"`
}

func (s UpdateSharedCustomRulesRequestConditionMethodConditions) String() string {
  return tea.Prettify(s)
}

func (s UpdateSharedCustomRulesRequestConditionMethodConditions) GoString() string {
  return s.String()
}

func (s *UpdateSharedCustomRulesRequestConditionMethodConditions) SetMatchType(v string) *UpdateSharedCustomRulesRequestConditionMethodConditions {
  s.MatchType = &v
  return s
}

func (s *UpdateSharedCustomRulesRequestConditionMethodConditions) SetRequestMethod(v []*string) *UpdateSharedCustomRulesRequestConditionMethodConditions {
  s.RequestMethod = v
  return s
}

type UpdateSharedCustomRulesRequestConditionAreaConditions struct     {
  // {"en":"Match type.\nEQUAL:Equals\nNOT_EQUAL:Does not equal","zh_CN":"匹配类型。\nEQUAL：等于\nNOT_EQUAL：不等于","exampleValue":"EQUAL,NOT_EQUAL"}
  MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
  // {"dictionary":"belong=WAAP-MS-Ext|dict=waap_areaCityAndCountry","en":"Geo.","zh_CN":"区域。"}
  Areas []*string `json:"areas,omitempty" xml:"areas,omitempty" require:"true" type:"Repeated"`
}

func (s UpdateSharedCustomRulesRequestConditionAreaConditions) String() string {
  return tea.Prettify(s)
}

func (s UpdateSharedCustomRulesRequestConditionAreaConditions) GoString() string {
  return s.String()
}

func (s *UpdateSharedCustomRulesRequestConditionAreaConditions) SetMatchType(v string) *UpdateSharedCustomRulesRequestConditionAreaConditions {
  s.MatchType = &v
  return s
}

func (s *UpdateSharedCustomRulesRequestConditionAreaConditions) SetAreas(v []*string) *UpdateSharedCustomRulesRequestConditionAreaConditions {
  s.Areas = v
  return s
}

type UpdateSharedCustomRulesRequestConditionIpOrIpsConditions struct     {
  // {"en":"Match type.\nEQUAL:Equals\nNOT_EQUAL:Does not equal","zh_CN":"匹配类型。\nEQUAL：等于\nNOT_EQUAL：不等于","exampleValue":"EQUAL,NOT_EQUAL"}
  MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
  // {"en":"IP/CIDR, maximum 500 IP/CIDR.","zh_CN":"IP/IP段，最多500个IP/IP段。"}
  IpOrIps []*string `json:"ipOrIps,omitempty" xml:"ipOrIps,omitempty" require:"true" type:"Repeated"`
}

func (s UpdateSharedCustomRulesRequestConditionIpOrIpsConditions) String() string {
  return tea.Prettify(s)
}

func (s UpdateSharedCustomRulesRequestConditionIpOrIpsConditions) GoString() string {
  return s.String()
}

func (s *UpdateSharedCustomRulesRequestConditionIpOrIpsConditions) SetMatchType(v string) *UpdateSharedCustomRulesRequestConditionIpOrIpsConditions {
  s.MatchType = &v
  return s
}

func (s *UpdateSharedCustomRulesRequestConditionIpOrIpsConditions) SetIpOrIps(v []*string) *UpdateSharedCustomRulesRequestConditionIpOrIpsConditions {
  s.IpOrIps = v
  return s
}

type UpdateSharedCustomRulesRequestConditionUriConditions struct     {
  // {"en":"Match type.\nEQUAL: Equals, URI case sensitive\nNOT_EQUAL: Does not equal, URI case sensitive\nCONTAIN: Contains, URI case insensitive\nNOT_CONTAIN: Does not Contains, URI case insensitive\nREGEX: Regex match, URI case insensitive\nNOT_REGEX: Regular does not match, URI case insensitive\nSTART_WITH: Starts with, URI case insensitive\nEND_WITH: Ends with, URI case insensitive\nWILDCARD: Wildcard matches, URI case insensitive,* represents zero or more arbitrary characters, ? represents any single character\nNOT_WILDCARD: Wildcard does not match, URI case insensitive,* represents zero or more arbitrary characters, ? represents any single character","zh_CN":"匹配类型。\nEQUAL：等于，URI大小写敏感\nNOT_EQUAL：不等于，URI大小写敏感\nCONTAIN：包含，URI大小写不敏感\nNOT_CONTAIN：不包含，URI大小写不敏感\nREGEX：匹配正则，URI大小写不敏感\nNOT_REGEX：正则不匹配，URI大小写不敏感\nSTART_WITH：开头是，URI大小写不敏感\nEND_WITH：结尾是，URI大小写不敏感\nWILDCARD：通配符匹配，URI大小写不敏感，*代表零个或多个任意字符，?代表任意单个字符\nNOT_WILDCARD：通配符不匹配，URI大小写不敏感，*代表零个或多个任意字符，?代表任意单个字符","exampleValue":"EQUAL,NOT_EQUAL,CONTAIN,NOT_CONTAIN,REGEX,NOT_REGEX,START_WITH,END_WITH,WILDCARD,NOT_WILDCARD"}
  MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
  // {"en":"URI.\nWhen match type is EQUAL/NOT_EQUAL/START_WITH/END_WITH, uri needs to start with \"/\", and includes parameters.\nWhen the match type is REGEX/NOT_REGEX, only one value is allowed.\nExample: /test.html?id=1.","zh_CN":"URI。\n当匹配类型为等于/不等于/开头是/结尾是，URI必须以”/“开头，含参数。\n当匹配类型为正则/正则不匹配，则只允许只有一个值。\n示例：/test.html?id=1。"}
  Uri []*string `json:"uri,omitempty" xml:"uri,omitempty" require:"true" type:"Repeated"`
}

func (s UpdateSharedCustomRulesRequestConditionUriConditions) String() string {
  return tea.Prettify(s)
}

func (s UpdateSharedCustomRulesRequestConditionUriConditions) GoString() string {
  return s.String()
}

func (s *UpdateSharedCustomRulesRequestConditionUriConditions) SetMatchType(v string) *UpdateSharedCustomRulesRequestConditionUriConditions {
  s.MatchType = &v
  return s
}

func (s *UpdateSharedCustomRulesRequestConditionUriConditions) SetUri(v []*string) *UpdateSharedCustomRulesRequestConditionUriConditions {
  s.Uri = v
  return s
}

type UpdateSharedCustomRulesRequestConditionPathConditions struct     {
  // {"en":"Match type.\nEQUAL: Equals, path case sensitive\nNOT_EQUAL: Does not equal, path case sensitive\nCONTAIN: Contains, path case insensitive\nNOT_CONTAIN: Does not Contains, path case insensitive\nREGEX: Regex match, path case insensitive\nNOT_REGEX: Regular does not match, path case sensitive\nSTART_WITH: Starts with, path case sensitive\nEND_WITH: Ends with, path case sensitive\nWILDCARD: Wildcard matches, path case sensitive,* represents zero or more arbitrary characters, ? represents any single character.\nNOT_WILDCARD: Wildcard does not match, path case sensitive,* represents zero or more arbitrary characters, ? represents any single character","zh_CN":"匹配类型。\nEQUAL：等于，路径大小写敏感\nNOT_EQUAL：不等于，路径大小写敏感\nCONTAIN：包含，路径大小写不敏感\nNOT_CONTAIN：不包含，路径大小写不敏感\nREGEX：匹配正则，路径大小写不敏感\nNOT_REGEX：正则不匹配，路径大小写不敏感\nSTART_WITH：开头是，路径大小写不敏感\nEND_WITH：结尾是，路径大小写不敏感\nWILDCARD：通配符匹配，路径大小写不敏感，*代表零个或多个任意字符，?代表任意单个字符\nNOT_WILDCARD：通配符不匹配，路径大小写不敏感，*代表零个或多个任意字符，?代表任意单个字符","exampleValue":"EQUAL,NOT_EQUAL,CONTAIN,NOT_CONTAIN,REGEX,NOT_REGEX,START_WITH,END_WITH,WILDCARD,NOT_WILDCARD"}
  MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
  // {"en":"Path.\nWhen match type is EQUAL/NOT_EQUAL/START_WITH/END_WITH, path needs to start with \"/\", and no parameters.\nWhen the match type is REGEX/NOT_REGEX, only one value is allowed.\nExample: /test.html.","zh_CN":"路径。\n当匹配类型为等于/不等于/开头是/结尾是，路径必须以“/”开头，不含参数。\n当匹配类型为正则/正则不匹配，则只允许只有一个值。\n示例：/test.html。"}
  UpdateSharedCustomRulesPaths []*string `json:"paths,omitempty" xml:"paths,omitempty" require:"true" type:"Repeated"`
}

func (s UpdateSharedCustomRulesRequestConditionPathConditions) String() string {
  return tea.Prettify(s)
}

func (s UpdateSharedCustomRulesRequestConditionPathConditions) GoString() string {
  return s.String()
}

func (s *UpdateSharedCustomRulesRequestConditionPathConditions) SetMatchType(v string) *UpdateSharedCustomRulesRequestConditionPathConditions {
  s.MatchType = &v
  return s
}

func (s *UpdateSharedCustomRulesRequestConditionPathConditions) SetUpdateSharedCustomRulesPaths(v []*string) *UpdateSharedCustomRulesRequestConditionPathConditions {
  s.UpdateSharedCustomRulesPaths = v
  return s
}

type UpdateSharedCustomRulesRequestConditionUaConditions struct     {
  // {"en":"Match type.\nEQUAL: Equals, user agent case sensitive\nNOT_EQUAL: Does not equal, user agent case sensitive\nCONTAIN: Contains, user agent case insensitive\nNOT_CONTAIN: Does not Contains, user agent case insensitive\nNONE:Empty or non-existent\nREGEX: Regex match, user agent case insensitive\nNOT_REGEX: Regular does not match, user agent case insensitive\nSTART_WITH: Starts with, user agent case insensitive\nEND_WITH: Ends with, user agent case insensitive\nWILDCARD: Wildcard matches, user agent case insensitive,* represents zero or more arbitrary characters, ? represents any single character\nNOT_WILDCARD: Wildcard does not match, user agent case insensitive,* represents zero or more arbitrary characters, ? represents any single character","zh_CN":"匹配类型。\nEQUAL：等于，User-Agent大小写敏感\nNOT_EQUAL：不等于，User-Agent大小写敏感\nCONTAIN：包含，User-Agent大小写不敏感\nNOT_CONTAIN：不包含，User-Agent大小写不敏感\nNONE：为空或不存在\nREGEX：匹配正则，User-Agent大小写不敏感\nNOT_REGEX：正则不匹配，User-Agent大小写不敏感\nSTART_WITH：开头是，User-Agent大小写不敏感\nEND_WITH：结尾是，User-Agent大小写不敏感\nWILDCARD：通配符匹配，User-Agent大小写不敏感，*代表零个或多个任意字符，?代表任意单个字符\nNOT_WILDCARD：通配符不匹配，User-Agent大小写不敏感，*代表零个或多个任意字符，?代表任意单个字符","exampleValue":"EQUAL,NOT_EQUAL,CONTAIN,NOT_CONTAIN,NONE,REGEX,NOT_REGEX,START_WITH,END_WITH,WILDCARD,NOT_WILDCARD"}
  MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
  // {"en":"User agent.\nWhen the match type is REGEX/NOT_REGEX, only one value is allowed.\nExample: go-Http-client/1.1.","zh_CN":"User-Agent。\n当匹配类型为正则/正则不匹配，则只允许只有一个值。\n示例：go-Http-client/1.1。"}
  Ua []*string `json:"ua,omitempty" xml:"ua,omitempty" require:"true" type:"Repeated"`
}

func (s UpdateSharedCustomRulesRequestConditionUaConditions) String() string {
  return tea.Prettify(s)
}

func (s UpdateSharedCustomRulesRequestConditionUaConditions) GoString() string {
  return s.String()
}

func (s *UpdateSharedCustomRulesRequestConditionUaConditions) SetMatchType(v string) *UpdateSharedCustomRulesRequestConditionUaConditions {
  s.MatchType = &v
  return s
}

func (s *UpdateSharedCustomRulesRequestConditionUaConditions) SetUa(v []*string) *UpdateSharedCustomRulesRequestConditionUaConditions {
  s.Ua = v
  return s
}

type UpdateSharedCustomRulesRequestConditionHeaderConditions struct     {
  // {"en":"Match type.\nEQUAL: Equals, request header values case sensitive\nNOT_EQUAL: Does not equal, request header values case sensitive\nCONTAIN: Contains, request header values case insensitive\nNOT_CONTAIN: Does not Contains, request header values case insensitive\nNONE: Empty or non-existent\nREGEX: Regex match, request header values case insensitive\nNOT_REGEX: Regular does not match, request header values case insensitive\nSTART_WITH: Starts with, request header values case insensitive\nEND_WITH: Ends with, request header values case insensitive\nWILDCARD: Wildcard matches, request header values case insensitive,* represents zero or more arbitrary characters, ? represents any single character\nNOT_WILDCARD: Wildcard does not match, request header values case insensitive,* represents zero or more arbitrary characters, ? represents any single character","zh_CN":"匹配类型。\nEQUAL：等于，头部值大小写敏感\nNOT_EQUAL：不等于，头部值大小写敏感\nCONTAIN：包含，头部值大小写不敏感\nNOT_CONTAIN：不包含，头部值大小写不敏感\nNONE：为空或不存在\nREGEX：匹配正则，头部值大小写不敏感\nNOT_REGEX：正则不匹配，头部值大小写不敏感\nSTART_WITH：开头是，头部值大小写不敏感\nEND_WITH：结尾是，头部值大小写不敏感\nWILDCARD：通配符匹配，头部值大小写不敏感，*代表零个或多个任意字符，?代表任意单个字符\nNOT_WILDCARD：通配符不匹配，头部值大小写不敏感，*代表零个或多个任意字符，?代表任意单个字符","exampleValue":"EQUAL,NOT_EQUAL,CONTAIN,NOT_CONTAIN,NONE,REGEX,NOT_REGEX,START_WITH,END_WITH,WILDCARD,NOT_WILDCARD"}
  MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
  // {"en":"Header value.\nWhen the match type is REGEX/NOT_REGEX, only one value is allowed.","zh_CN":"头部值。\n当匹配类型为正则/正则不匹配，则只允许只有一个值。"}
  ValueList []*string `json:"valueList,omitempty" xml:"valueList,omitempty" require:"true" type:"Repeated"`
  // {"en":"Header name,case insensitive,up to 100 characters.\nExample: Accept.","zh_CN":"头部名称，大小写不敏感，最多100个字符。\n示例：Accept。"}
  Key *string `json:"key,omitempty" xml:"key,omitempty" require:"true"`
}

func (s UpdateSharedCustomRulesRequestConditionHeaderConditions) String() string {
  return tea.Prettify(s)
}

func (s UpdateSharedCustomRulesRequestConditionHeaderConditions) GoString() string {
  return s.String()
}

func (s *UpdateSharedCustomRulesRequestConditionHeaderConditions) SetMatchType(v string) *UpdateSharedCustomRulesRequestConditionHeaderConditions {
  s.MatchType = &v
  return s
}

func (s *UpdateSharedCustomRulesRequestConditionHeaderConditions) SetValueList(v []*string) *UpdateSharedCustomRulesRequestConditionHeaderConditions {
  s.ValueList = v
  return s
}

func (s *UpdateSharedCustomRulesRequestConditionHeaderConditions) SetKey(v string) *UpdateSharedCustomRulesRequestConditionHeaderConditions {
  s.Key = &v
  return s
}

type UpdateSharedCustomRulesRequestConditionRefererConditions struct     {
  // {"en":"Referer.\nWhen the match type is REGEX/NOT_REGEX, only one value is allowed.\nExample: http://test.com.","zh_CN":"Referer。\n当匹配类型为正则/正则不匹配，则只允许只有一个值。\n示例：http://test.com。"}
  Referer []*string `json:"referer,omitempty" xml:"referer,omitempty" require:"true" type:"Repeated"`
  // {"en":"Match type.\nEQUAL: Equals, referer case sensitive\nNOT_EQUAL: Does not equal, referer case sensitive\nCONTAIN: Contains, referer case insensitive\nNOT_CONTAIN: Does not Contains, referer case insensitive\nNONE:Empty or non-existent\nREGEX: Regex match, referer case insensitive\nNOT_REGEX: Regular does not match, referer case insensitive\nSTART_WITH: Starts with, referer case insensitive\nEND_WITH: Ends with, referer case insensitive\nWILDCARD: Wildcard matches, referer case insensitive,* represents zero or more arbitrary characters, ? represents any single characte\nNOT_WILDCARD: Wildcard does not match, referer case insensitive,* represents zero or more arbitrary characters, ? represents any single character","zh_CN":"匹配类型。\nEQUAL：等于，referer大小写敏感\nNOT_EQUAL：不等于，referer大小写敏感\nCONTAIN：包含，referer大小写不敏感\nNOT_CONTAIN：不包含，referer大小写不敏感\nNONE：为空或不存在\nREGEX：匹配正则，referer大小写不敏感\nNOT_REGEX：正则不匹配，referer大小写不敏感\nSTART_WITH：开头是，referer大小写不敏感\nEND_WITH：结尾是，referer大小写不敏感\nWILDCARD：通配符匹配，referer大小写不敏感，*代表零个或多个任意字符，?代表任意单个字符\nNOT_WILDCARD：通配符不匹配，referer大小写不敏感，*代表零个或多个任意字符，?代表任意单个字符","exampleValue":"EQUAL,NOT_EQUAL,CONTAIN,NOT_CONTAIN,NONE,REGEX,NOT_REGEX,START_WITH,END_WITH,WILDCARD,NOT_WILDCARD"}
  MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
}

func (s UpdateSharedCustomRulesRequestConditionRefererConditions) String() string {
  return tea.Prettify(s)
}

func (s UpdateSharedCustomRulesRequestConditionRefererConditions) GoString() string {
  return s.String()
}

func (s *UpdateSharedCustomRulesRequestConditionRefererConditions) SetReferer(v []*string) *UpdateSharedCustomRulesRequestConditionRefererConditions {
  s.Referer = v
  return s
}

func (s *UpdateSharedCustomRulesRequestConditionRefererConditions) SetMatchType(v string) *UpdateSharedCustomRulesRequestConditionRefererConditions {
  s.MatchType = &v
  return s
}

type UpdateSharedCustomRulesRequestConditionJa3Conditions struct     {
  // {"en":"Match type.\nEQUAL: Equals\nNOT_EQUAL: Does not equal","zh_CN":"匹配类型。\nEQUAL：等于 \nNOT_EQUAL：不等于","exampleValue":"EQUAL,NOT_EQUAL"}
  MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
  // {"en":"JA3 Fingerprint List, maximum 300 JA3 Fingerprint.\nWhen the match type is EQUAL/NOT_EQUAL, each item's character length must be 32 and can only include numbers and lowercase letters.","zh_CN":"JA3指纹列表，最多300个JA3指纹。\n当匹配类型为等于/不等于时，每项字符长度必须为32，且仅限数字和小写字母。"}
  Ja3List []*string `json:"ja3List,omitempty" xml:"ja3List,omitempty" require:"true" type:"Repeated"`
}

func (s UpdateSharedCustomRulesRequestConditionJa3Conditions) String() string {
  return tea.Prettify(s)
}

func (s UpdateSharedCustomRulesRequestConditionJa3Conditions) GoString() string {
  return s.String()
}

func (s *UpdateSharedCustomRulesRequestConditionJa3Conditions) SetMatchType(v string) *UpdateSharedCustomRulesRequestConditionJa3Conditions {
  s.MatchType = &v
  return s
}

func (s *UpdateSharedCustomRulesRequestConditionJa3Conditions) SetJa3List(v []*string) *UpdateSharedCustomRulesRequestConditionJa3Conditions {
  s.Ja3List = v
  return s
}

type UpdateSharedCustomRulesRequestConditionJa4Conditions struct     {
  // {"en":"Match type. \nEQUAL: Equals\nNOT_EQUAL: Does not equal\nCONTAIN: Contains\nNOT_CONTAIN: Does not Contains\nSTART_WITH: Starts with\nEND_WITH: Ends with\nWILDCARD: Wildcard matches, ** represents zero or more arbitrary characters, ? represents any single character\nNOT_WILDCARD: Wildcard does not match, ** represents zero or more arbitrary characters, ? represents any single character","zh_CN":"匹配类型。\nEQUAL：等于\nNOT_EQUAL：不等于\nCONTAIN：包含\nNOT_CONTAIN：不包含\nSTART_WITH：开头是\nEND_WITH：结尾是\nWILDCARD：通配符匹配，*代表零个或多个任意字符，?代表任意单个字符\nNOT_WILDCARD：通配符不匹配，*代表零个或多个任意字符，?代表任意单个字符","exampleValue":"EQUAL,NOT_EQUAL,CONTAIN,NOT_CONTAIN,START_WITH,END_WITH,WILDCARD,NOT_WILDCARD"}
  MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
  // {"en":"JA4 Fingerprint List, maximum 300 JA4 Fingerprint.\nWhen the match type is EQUAL/NOT_EQUAL, each item's format must be 10 characters + 12 characters + 12 characters, separated by underscores, and can only include underscores, numbers, and lowercase letters.\nWhen the match type is CONTAIN/NOT_CONTAIN/START_WITH/END_WITH, each item is only allowed to include underscores, numbers, and lowercase letters.\nWhen the match type is WILDCARD/NOT_WILDCARD, each item, aside from  ** and ?, is only allowed to include underscores, numbers, and lowercase letters.","zh_CN":"JA4指纹列表，最多300个JA4指纹。\n当匹配类型为等于/不等于时，每项格式必须为10位字符+12位字符+12位字符，中间以“_”分隔，且仅限下划线、数字和小写字母。\n当匹配类型为包含/不包含/开头是/结尾是时，每项只允许输入下划线、数字和小写字母。\n当匹配类型为通配符匹配/通配符不匹配时，每项除*和?外，只允许输入下划线、数字和小写字母。"}
  Ja4List []*string `json:"ja4List,omitempty" xml:"ja4List,omitempty" require:"true" type:"Repeated"`
}

func (s UpdateSharedCustomRulesRequestConditionJa4Conditions) String() string {
  return tea.Prettify(s)
}

func (s UpdateSharedCustomRulesRequestConditionJa4Conditions) GoString() string {
  return s.String()
}

func (s *UpdateSharedCustomRulesRequestConditionJa4Conditions) SetMatchType(v string) *UpdateSharedCustomRulesRequestConditionJa4Conditions {
  s.MatchType = &v
  return s
}

func (s *UpdateSharedCustomRulesRequestConditionJa4Conditions) SetJa4List(v []*string) *UpdateSharedCustomRulesRequestConditionJa4Conditions {
  s.Ja4List = v
  return s
}

type UpdateSharedCustomRulesRequestHeader struct {
  // {"en":"Security service type. Please enter a specific service type, if you purchase multiple security services.","zh_CN":"安全服务类型。有使用多个不同的安全服务时，需要填写具体的服务类型。"}
  ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty"`
}

func (s UpdateSharedCustomRulesRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s UpdateSharedCustomRulesRequestHeader) GoString() string {
  return s.String()
}

func (s *UpdateSharedCustomRulesRequestHeader) SetServiceType(v string) *UpdateSharedCustomRulesRequestHeader {
  s.ServiceType = &v
  return s
}

type UpdateSharedCustomRulesPaths struct {
}

func (s UpdateSharedCustomRulesPaths) String() string {
  return tea.Prettify(s)
}

func (s UpdateSharedCustomRulesPaths) GoString() string {
  return s.String()
}

type UpdateSharedCustomRulesParameters struct {
}

func (s UpdateSharedCustomRulesParameters) String() string {
  return tea.Prettify(s)
}

func (s UpdateSharedCustomRulesParameters) GoString() string {
  return s.String()
}

type UpdateSharedCustomRulesResponse struct {
  // {"en":"Description.","zh_CN":"描述信息。"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"dictionary":"belong=WAAP-MS-Ext|dict=waap_retCodeEnum","en":"Please refer to the error code for exceptions.","zh_CN":"请参照错误码。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Data.","zh_CN":"出参数据。"}
  Data *bool `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s UpdateSharedCustomRulesResponse) String() string {
  return tea.Prettify(s)
}

func (s UpdateSharedCustomRulesResponse) GoString() string {
  return s.String()
}

func (s *UpdateSharedCustomRulesResponse) SetMsg(v string) *UpdateSharedCustomRulesResponse {
  s.Msg = &v
  return s
}

func (s *UpdateSharedCustomRulesResponse) SetCode(v string) *UpdateSharedCustomRulesResponse {
  s.Code = &v
  return s
}

func (s *UpdateSharedCustomRulesResponse) SetData(v bool) *UpdateSharedCustomRulesResponse {
  s.Data = &v
  return s
}

type UpdateSharedCustomRulesResponseHeader struct {
}

func (s UpdateSharedCustomRulesResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s UpdateSharedCustomRulesResponseHeader) GoString() string {
  return s.String()
}




type AssociateShareCustomizeBotsRequest struct {
  // {"en":"Share configuration ID.","zh_CN":"共享配置ID。"}
  ShareId *string `json:"shareId,omitempty" xml:"shareId,omitempty" require:"true"`
  // {"en":"List of hostname to be associated.","zh_CN":"待关联域名列表。"}
  DomainList []*string `json:"domainList,omitempty" xml:"domainList,omitempty" require:"true" type:"Repeated"`
}

func (s AssociateShareCustomizeBotsRequest) String() string {
  return tea.Prettify(s)
}

func (s AssociateShareCustomizeBotsRequest) GoString() string {
  return s.String()
}

func (s *AssociateShareCustomizeBotsRequest) SetShareId(v string) *AssociateShareCustomizeBotsRequest {
  s.ShareId = &v
  return s
}

func (s *AssociateShareCustomizeBotsRequest) SetDomainList(v []*string) *AssociateShareCustomizeBotsRequest {
  s.DomainList = v
  return s
}

type AssociateShareCustomizeBotsRequestHeader struct {
  // {"en":"Security service type. Please enter a specific service type, if you purchase multiple security services.","zh_CN":"安全服务类型。有使用多个不同的安全服务时，需要填写具体的服务类型。"}
  ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty"`
}

func (s AssociateShareCustomizeBotsRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s AssociateShareCustomizeBotsRequestHeader) GoString() string {
  return s.String()
}

func (s *AssociateShareCustomizeBotsRequestHeader) SetServiceType(v string) *AssociateShareCustomizeBotsRequestHeader {
  s.ServiceType = &v
  return s
}

type AssociateShareCustomizeBotsPaths struct {
}

func (s AssociateShareCustomizeBotsPaths) String() string {
  return tea.Prettify(s)
}

func (s AssociateShareCustomizeBotsPaths) GoString() string {
  return s.String()
}

type AssociateShareCustomizeBotsParameters struct {
}

func (s AssociateShareCustomizeBotsParameters) String() string {
  return tea.Prettify(s)
}

func (s AssociateShareCustomizeBotsParameters) GoString() string {
  return s.String()
}

type AssociateShareCustomizeBotsResponse struct {
  // {"dictionary":"belong=WAAP-MS-Ext|dict=waap_retCodeEnum","en":"Please refer to the error code for exceptions.","zh_CN":"请参照错误码。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Description.","zh_CN":"描述信息。"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
}

func (s AssociateShareCustomizeBotsResponse) String() string {
  return tea.Prettify(s)
}

func (s AssociateShareCustomizeBotsResponse) GoString() string {
  return s.String()
}

func (s *AssociateShareCustomizeBotsResponse) SetCode(v string) *AssociateShareCustomizeBotsResponse {
  s.Code = &v
  return s
}

func (s *AssociateShareCustomizeBotsResponse) SetMsg(v string) *AssociateShareCustomizeBotsResponse {
  s.Msg = &v
  return s
}

type AssociateShareCustomizeBotsResponseHeader struct {
}

func (s AssociateShareCustomizeBotsResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s AssociateShareCustomizeBotsResponseHeader) GoString() string {
  return s.String()
}




type DeleteSharedWAFRuleExceptionRequest struct {
  // {"en":"Rule exception ID list.","zh_CN":"规则例外ID列表。"}
  IdList []*string `json:"idList,omitempty" xml:"idList,omitempty" require:"true" type:"Repeated"`
}

func (s DeleteSharedWAFRuleExceptionRequest) String() string {
  return tea.Prettify(s)
}

func (s DeleteSharedWAFRuleExceptionRequest) GoString() string {
  return s.String()
}

func (s *DeleteSharedWAFRuleExceptionRequest) SetIdList(v []*string) *DeleteSharedWAFRuleExceptionRequest {
  s.IdList = v
  return s
}

type DeleteSharedWAFRuleExceptionRequestHeader struct {
  // {"en":"Security service type. Please enter a specific service type, if you purchase multiple security services.","zh_CN":"安全服务类型。有使用多个不同的安全服务时，需要填写具体的服务类型。"}
  ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty"`
}

func (s DeleteSharedWAFRuleExceptionRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteSharedWAFRuleExceptionRequestHeader) GoString() string {
  return s.String()
}

func (s *DeleteSharedWAFRuleExceptionRequestHeader) SetServiceType(v string) *DeleteSharedWAFRuleExceptionRequestHeader {
  s.ServiceType = &v
  return s
}

type DeleteSharedWAFRuleExceptionPaths struct {
}

func (s DeleteSharedWAFRuleExceptionPaths) String() string {
  return tea.Prettify(s)
}

func (s DeleteSharedWAFRuleExceptionPaths) GoString() string {
  return s.String()
}

type DeleteSharedWAFRuleExceptionParameters struct {
}

func (s DeleteSharedWAFRuleExceptionParameters) String() string {
  return tea.Prettify(s)
}

func (s DeleteSharedWAFRuleExceptionParameters) GoString() string {
  return s.String()
}

type DeleteSharedWAFRuleExceptionResponse struct {
  // {"dictionary":"belong=WAAP-MS-Ext|dict=waap_retCodeEnum","en":"Please refer to the error code for exceptions.","zh_CN":"请参照错误码。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Description.","zh_CN":"描述信息。"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"Delete success indicator.","zh_CN":"删除成功标识。"}
  Data *bool `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s DeleteSharedWAFRuleExceptionResponse) String() string {
  return tea.Prettify(s)
}

func (s DeleteSharedWAFRuleExceptionResponse) GoString() string {
  return s.String()
}

func (s *DeleteSharedWAFRuleExceptionResponse) SetCode(v string) *DeleteSharedWAFRuleExceptionResponse {
  s.Code = &v
  return s
}

func (s *DeleteSharedWAFRuleExceptionResponse) SetMsg(v string) *DeleteSharedWAFRuleExceptionResponse {
  s.Msg = &v
  return s
}

func (s *DeleteSharedWAFRuleExceptionResponse) SetData(v bool) *DeleteSharedWAFRuleExceptionResponse {
  s.Data = &v
  return s
}

type DeleteSharedWAFRuleExceptionResponseHeader struct {
}

func (s DeleteSharedWAFRuleExceptionResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteSharedWAFRuleExceptionResponseHeader) GoString() string {
  return s.String()
}




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

type ListShareCustomizeBotsRequestHeader struct {
  // {"en":"Security service type. Please enter a specific service type, if you purchase multiple security services.","zh_CN":"安全服务类型。有使用多个不同的安全服务时，需要填写具体的服务类型。"}
  ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty"`
}

func (s ListShareCustomizeBotsRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ListShareCustomizeBotsRequestHeader) GoString() string {
  return s.String()
}

func (s *ListShareCustomizeBotsRequestHeader) SetServiceType(v string) *ListShareCustomizeBotsRequestHeader {
  s.ServiceType = &v
  return s
}

type ListShareCustomizeBotsPaths struct {
}

func (s ListShareCustomizeBotsPaths) String() string {
  return tea.Prettify(s)
}

func (s ListShareCustomizeBotsPaths) GoString() string {
  return s.String()
}

type ListShareCustomizeBotsParameters struct {
}

func (s ListShareCustomizeBotsParameters) String() string {
  return tea.Prettify(s)
}

func (s ListShareCustomizeBotsParameters) GoString() string {
  return s.String()
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

type ListShareCustomizeBotsResponseData struct     {
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

type ListShareCustomizeBotsResponseDataConditionList struct     {
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

type ListShareCustomizeBotsResponseHeader struct {
}

func (s ListShareCustomizeBotsResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ListShareCustomizeBotsResponseHeader) GoString() string {
  return s.String()
}




type CreateAppApiExceptionFeatureRequest struct {
  // {"en":"Rule name, duplicate not allowed.","zh_CN":"规则名称，不允许重复"}
  RuleName *string `json:"ruleName,omitempty" xml:"ruleName,omitempty" require:"true"`
  // {"en":"Type enumeration, for example:\nNATIVE_ APP        remarks: Native App\nCallback_ API       remarks: Callback API\nOther_ API         remarks: Other program API\nHYBRID_ APP     remarks: Hybrid APP","zh_CN":"类型枚举，例如：\nNATIVE_APP          备注：原生APP\nCALLBACK_API      备注：回调API\nOTHER_API            备注：其他程序API\nHYBRID_APP          备注：混合APP","exampleValue":"NATIVE_APP,CALLBACK_API,OTHER_API,HYBRID_APP"}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
  // {"en":"Characteristic.","zh_CN":"特征。"}
  Config *CreateAppApiExceptionFeatureRequestConfig `json:"config,omitempty" xml:"config,omitempty" require:"true" type:"Struct"`
}

func (s CreateAppApiExceptionFeatureRequest) String() string {
  return tea.Prettify(s)
}

func (s CreateAppApiExceptionFeatureRequest) GoString() string {
  return s.String()
}

func (s *CreateAppApiExceptionFeatureRequest) SetRuleName(v string) *CreateAppApiExceptionFeatureRequest {
  s.RuleName = &v
  return s
}

func (s *CreateAppApiExceptionFeatureRequest) SetType(v string) *CreateAppApiExceptionFeatureRequest {
  s.Type = &v
  return s
}

func (s *CreateAppApiExceptionFeatureRequest) SetConfig(v *CreateAppApiExceptionFeatureRequestConfig) *CreateAppApiExceptionFeatureRequest {
  s.Config = v
  return s
}

type CreateAppApiExceptionFeatureRequestConfig struct {
  // {"en":"Matching condition name.\nPATH         Path\nURI            URI\nUA             User-Agent\nREFERER    Referer\nHEADER    Request Header","zh_CN":"匹配条件名称。\nPATH                       路径\nURI                          URI\nUA                           User-Agent\nREFERER                  Referer\nHEAD                      Request Header","exampleValue":"PATH,URI,UA,REFERER,HEAD"}
  Condition *string `json:"condition,omitempty" xml:"condition,omitempty" require:"true"`
  // {"en":"Matching condition function.\nCONTAIN                  contains\nNOT_CONTAIN         does not contain\nEQUAL                      equals\nNOT_EQUAL             does not equal\nEMPTY                      does not exist or has no value(When the condition is a path or URI, this value is not optional)\nREGEX                      regex match","zh_CN":"匹配条件函数。\nCONTAIN                  包含\nNOT_CONTAIN         不包含\nEQUAL                      等于\nNOT_EQUAL             不等于\nEMPTY                      为空或不存在（当condition为路径或者URI时，该值不可选）\nREGEX                       匹配正则","exampleValue":"CONTAIN,NOT_CONTAIN,EQUAL,NOT_EQUAL,EMPTY,REGEX"}
  MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
  // {"en":"Condition value list, When matchType is EMPTY, the value is empty;  When matchType is REGEX, only one item is set","zh_CN":"条件值列表，当matchType为EMPTY时，该值为空；当matchType为REGEX时，只能传一条"}
  ConditionValue []*string `json:"conditionValue,omitempty" xml:"conditionValue,omitempty" require:"true" type:"Repeated"`
  // {"en":"Head value, If condition=HEAD, then this field is mandatory, Otherwise, it is empty.","zh_CN":"头部值，如果condition为HEAD，则该字段必填，否则为空。"}
  Key *string `json:"key,omitempty" xml:"key,omitempty"`
}

func (s CreateAppApiExceptionFeatureRequestConfig) String() string {
  return tea.Prettify(s)
}

func (s CreateAppApiExceptionFeatureRequestConfig) GoString() string {
  return s.String()
}

func (s *CreateAppApiExceptionFeatureRequestConfig) SetCondition(v string) *CreateAppApiExceptionFeatureRequestConfig {
  s.Condition = &v
  return s
}

func (s *CreateAppApiExceptionFeatureRequestConfig) SetMatchType(v string) *CreateAppApiExceptionFeatureRequestConfig {
  s.MatchType = &v
  return s
}

func (s *CreateAppApiExceptionFeatureRequestConfig) SetConditionValue(v []*string) *CreateAppApiExceptionFeatureRequestConfig {
  s.ConditionValue = v
  return s
}

func (s *CreateAppApiExceptionFeatureRequestConfig) SetKey(v string) *CreateAppApiExceptionFeatureRequestConfig {
  s.Key = &v
  return s
}

type CreateAppApiExceptionFeatureRequestHeader struct {
  // {"en":"Security service type. Please enter a specific service type, if you purchase multiple security services.","zh_CN":"安全服务类型。有使用多个不同的安全服务时，需要填写具体的服务类型。"}
  ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty"`
}

func (s CreateAppApiExceptionFeatureRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s CreateAppApiExceptionFeatureRequestHeader) GoString() string {
  return s.String()
}

func (s *CreateAppApiExceptionFeatureRequestHeader) SetServiceType(v string) *CreateAppApiExceptionFeatureRequestHeader {
  s.ServiceType = &v
  return s
}

type CreateAppApiExceptionFeaturePaths struct {
}

func (s CreateAppApiExceptionFeaturePaths) String() string {
  return tea.Prettify(s)
}

func (s CreateAppApiExceptionFeaturePaths) GoString() string {
  return s.String()
}

type CreateAppApiExceptionFeatureParameters struct {
}

func (s CreateAppApiExceptionFeatureParameters) String() string {
  return tea.Prettify(s)
}

func (s CreateAppApiExceptionFeatureParameters) GoString() string {
  return s.String()
}

type CreateAppApiExceptionFeatureResponse struct {
  // {"dictionary":"belong=WAAP-MS-Ext|dict=waap_retCodeEnum","en":"Please refer to the error code for exceptions.","zh_CN":"请参照错误码。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Description.","zh_CN":"描述信息。"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
}

func (s CreateAppApiExceptionFeatureResponse) String() string {
  return tea.Prettify(s)
}

func (s CreateAppApiExceptionFeatureResponse) GoString() string {
  return s.String()
}

func (s *CreateAppApiExceptionFeatureResponse) SetCode(v string) *CreateAppApiExceptionFeatureResponse {
  s.Code = &v
  return s
}

func (s *CreateAppApiExceptionFeatureResponse) SetMsg(v string) *CreateAppApiExceptionFeatureResponse {
  s.Msg = &v
  return s
}

type CreateAppApiExceptionFeatureResponseHeader struct {
}

func (s CreateAppApiExceptionFeatureResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s CreateAppApiExceptionFeatureResponseHeader) GoString() string {
  return s.String()
}




type DisassociateSharedWhitelistRuleRequest struct {
  // {"en":"Share rule ID.","zh_CN":"共享规则ID。"}
  ShareId *string `json:"shareId,omitempty" xml:"shareId,omitempty" require:"true"`
  // {"en":"List of domain to be disassociated.","zh_CN":"待取消关联域名列表。"}
  DomainList []*string `json:"domainList,omitempty" xml:"domainList,omitempty" require:"true" type:"Repeated"`
}

func (s DisassociateSharedWhitelistRuleRequest) String() string {
  return tea.Prettify(s)
}

func (s DisassociateSharedWhitelistRuleRequest) GoString() string {
  return s.String()
}

func (s *DisassociateSharedWhitelistRuleRequest) SetShareId(v string) *DisassociateSharedWhitelistRuleRequest {
  s.ShareId = &v
  return s
}

func (s *DisassociateSharedWhitelistRuleRequest) SetDomainList(v []*string) *DisassociateSharedWhitelistRuleRequest {
  s.DomainList = v
  return s
}

type DisassociateSharedWhitelistRuleRequestHeader struct {
  // {"en":"Security service type. Please enter a specific service type, if you purchase multiple security services.","zh_CN":"安全服务类型。有使用多个不同的安全服务时，需要填写具体的服务类型。"}
  ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty"`
}

func (s DisassociateSharedWhitelistRuleRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s DisassociateSharedWhitelistRuleRequestHeader) GoString() string {
  return s.String()
}

func (s *DisassociateSharedWhitelistRuleRequestHeader) SetServiceType(v string) *DisassociateSharedWhitelistRuleRequestHeader {
  s.ServiceType = &v
  return s
}

type DisassociateSharedWhitelistRulePaths struct {
}

func (s DisassociateSharedWhitelistRulePaths) String() string {
  return tea.Prettify(s)
}

func (s DisassociateSharedWhitelistRulePaths) GoString() string {
  return s.String()
}

type DisassociateSharedWhitelistRuleParameters struct {
}

func (s DisassociateSharedWhitelistRuleParameters) String() string {
  return tea.Prettify(s)
}

func (s DisassociateSharedWhitelistRuleParameters) GoString() string {
  return s.String()
}

type DisassociateSharedWhitelistRuleResponse struct {
  // {"dictionary":"belong=WAAP-MS-Ext|dict=waap_retCodeEnum","en":"Please refer to the error code for exceptions.","zh_CN":"请参照错误码。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Description.","zh_CN":"描述信息。"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
}

func (s DisassociateSharedWhitelistRuleResponse) String() string {
  return tea.Prettify(s)
}

func (s DisassociateSharedWhitelistRuleResponse) GoString() string {
  return s.String()
}

func (s *DisassociateSharedWhitelistRuleResponse) SetCode(v string) *DisassociateSharedWhitelistRuleResponse {
  s.Code = &v
  return s
}

func (s *DisassociateSharedWhitelistRuleResponse) SetMsg(v string) *DisassociateSharedWhitelistRuleResponse {
  s.Msg = &v
  return s
}

type DisassociateSharedWhitelistRuleResponseHeader struct {
}

func (s DisassociateSharedWhitelistRuleResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s DisassociateSharedWhitelistRuleResponseHeader) GoString() string {
  return s.String()
}




type UpdateSharedRateLimitingRuleRequest struct {
  // {"en":"Rule ID.","zh_CN":"规则ID。"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {"en":"Rule Name, maximum 100 characters.Does not support # and &.","zh_CN":"规则名称，最多100个字符。 不支持 # 和 &。"}
  RuleName *string `json:"ruleName,omitempty" xml:"ruleName,omitempty"`
  // {"en":"Description, maximum 1000 characters.","zh_CN":"规则描述，最多1000个字符。"}
  Description *string `json:"description,omitempty" xml:"description,omitempty"`
  // {"en":"Client identifier.\nIP:Client IP\nIP_UA:Client IP and User-Agent\nCOOKIE:Cookie\nIP_COOKIE:Client IP and Cookie\nHEADER:Request Header\nWhen there is a status code in the matching condition,this client identifier is not supported.\nIP_HEADER:Client IP and Request Header\nWhen there is a status code in the matching condition,this client identifier is not supported .","zh_CN":"统计粒度。\nIP：客户端IP\nIP_UA：客户端IP和User-Agent\nCOOKIE：Cookie\nIP_COOKIE：客户端IP和Cookie\nHEADER：请求头，当匹配条件中存在状态码时不支持此统计粒度\nIP_HEADER：客户端IP和请求头，当匹配条件中存在状态码时不支持此统计粒度"}
  StatisticalItem *string `json:"statisticalItem,omitempty" xml:"statisticalItem,omitempty"`
  // {"en":"Statistical key value.\nWhen the client identifier is cookie/header value, the corresponding key value needs to be entered.","zh_CN":"统计key值。\n当统计粒度cookie/header值，需要输入对应的key值。"}
  StatisticsKey *string `json:"statisticsKey,omitempty" xml:"statisticsKey,omitempty"`
  // {"en":"Statistics period, unit: seconds, the range is 1 - 3600.","zh_CN":"统计周期，单位：秒，范围为 1 - 3600。"}
  StatisticalPeriod *int `json:"statisticalPeriod,omitempty" xml:"statisticalPeriod,omitempty"`
  // {"en":"Trigger threshold, unit: times.","zh_CN":"触发阈值，单位：次。"}
  TriggerThreshold *int `json:"triggerThreshold,omitempty" xml:"triggerThreshold,omitempty"`
  // {"en":"Action duration, unit: seconds, the range is 10 - 604800.","zh_CN":"处理动作持续时间，单位：秒，范围为 10 - 604800。"}
  InterceptTime *int `json:"interceptTime,omitempty" xml:"interceptTime,omitempty"`
  // {"en":"Cycle effective status.\nPERMANENT:All time\nWITHOUT:Excluded time\nWITHIN:Selected time","zh_CN":"周期生效状态。\nPERMANENT：永久生效\nWITHOUT：周期内不生效\nWITHIN：周期内生效"}
  EffectiveStatus *string `json:"effectiveStatus,omitempty" xml:"effectiveStatus,omitempty"`
  // {"en":"Effective time period.\nWhen the effective status is effective within the cycle or not effective within the cycle, this field must have a value.","zh_CN":"规则生效周期。\n生效状态为周期内生效或周期内不生效时，此字段必须有值。"}
  RateLimitEffective *UpdateSharedRateLimitingRuleRequestRateLimitEffective `json:"rateLimitEffective,omitempty" xml:"rateLimitEffective,omitempty" type:"Struct"`
  // {"en":"Action.\nNO_USE:Not Used\nLOG:Log\nCOOKIE:Cookie verification\nJS_CHECK:Javascript verification\nDELAY:Delay\nBLOCK:Deny\nRESET:Reset Connection\nCUSTOM_ACTION_ID:Fill in the custom action id of the corresponding action\n(When there is a status code in the matching condition, the supported actions are Log, Deny, and Reset Connection.)","zh_CN":"处理动作。\nNO_USE：不使用\nLOG：监控\nCOOKIE：Cookie校验\nJS_CHECK：JavaScript校验\nDELAY：延迟响应\nBLOCK：拦截\nRESET：断开连接\nCUSTOM_ACTION_ID：传入对应的自定义响应动作的ID\n当匹配条件中存在状态码时，支持处理动作为监控、拦截、断开连接。"}
  Action *string `json:"action,omitempty" xml:"action,omitempty"`
  // {"en":"Matching conditions.","zh_CN":"匹配条件。"}
  RateLimitRuleCondition *UpdateSharedRateLimitingRuleRequestRateLimitRuleCondition `json:"rateLimitRuleCondition,omitempty" xml:"rateLimitRuleCondition,omitempty" type:"Struct"`
}

func (s UpdateSharedRateLimitingRuleRequest) String() string {
  return tea.Prettify(s)
}

func (s UpdateSharedRateLimitingRuleRequest) GoString() string {
  return s.String()
}

func (s *UpdateSharedRateLimitingRuleRequest) SetId(v string) *UpdateSharedRateLimitingRuleRequest {
  s.Id = &v
  return s
}

func (s *UpdateSharedRateLimitingRuleRequest) SetRuleName(v string) *UpdateSharedRateLimitingRuleRequest {
  s.RuleName = &v
  return s
}

func (s *UpdateSharedRateLimitingRuleRequest) SetDescription(v string) *UpdateSharedRateLimitingRuleRequest {
  s.Description = &v
  return s
}

func (s *UpdateSharedRateLimitingRuleRequest) SetStatisticalItem(v string) *UpdateSharedRateLimitingRuleRequest {
  s.StatisticalItem = &v
  return s
}

func (s *UpdateSharedRateLimitingRuleRequest) SetStatisticsKey(v string) *UpdateSharedRateLimitingRuleRequest {
  s.StatisticsKey = &v
  return s
}

func (s *UpdateSharedRateLimitingRuleRequest) SetStatisticalPeriod(v int) *UpdateSharedRateLimitingRuleRequest {
  s.StatisticalPeriod = &v
  return s
}

func (s *UpdateSharedRateLimitingRuleRequest) SetTriggerThreshold(v int) *UpdateSharedRateLimitingRuleRequest {
  s.TriggerThreshold = &v
  return s
}

func (s *UpdateSharedRateLimitingRuleRequest) SetInterceptTime(v int) *UpdateSharedRateLimitingRuleRequest {
  s.InterceptTime = &v
  return s
}

func (s *UpdateSharedRateLimitingRuleRequest) SetEffectiveStatus(v string) *UpdateSharedRateLimitingRuleRequest {
  s.EffectiveStatus = &v
  return s
}

func (s *UpdateSharedRateLimitingRuleRequest) SetRateLimitEffective(v *UpdateSharedRateLimitingRuleRequestRateLimitEffective) *UpdateSharedRateLimitingRuleRequest {
  s.RateLimitEffective = v
  return s
}

func (s *UpdateSharedRateLimitingRuleRequest) SetAction(v string) *UpdateSharedRateLimitingRuleRequest {
  s.Action = &v
  return s
}

func (s *UpdateSharedRateLimitingRuleRequest) SetRateLimitRuleCondition(v *UpdateSharedRateLimitingRuleRequestRateLimitRuleCondition) *UpdateSharedRateLimitingRuleRequest {
  s.RateLimitRuleCondition = v
  return s
}

type UpdateSharedRateLimitingRuleRequestRateLimitEffective struct {
  // {"en":"Effective.\nMON:Monday\nTUE:Tuesday\nWED:Wednesday\nTHU:Thursday\nFRI:Friday\nSAT:Saturday\nSUN:Sunday","zh_CN":"周期。\nMON：星期一\nTUE：星期二\nWED：星期三\nTHU：星期四\nFRI：星期五\nSAT：星期六\nSUN：星期天"}
  Effective []*string `json:"effective,omitempty" xml:"effective,omitempty" require:"true" type:"Repeated"`
  // {"en":"Start time, format: HH:mm.","zh_CN":"开始时间，格式：HH:mm。"}
  Start *string `json:"start,omitempty" xml:"start,omitempty" require:"true"`
  // {"en":"End time, format: HH:mm.","zh_CN":"结束时间，格式：HH:mm。"}
  End *string `json:"end,omitempty" xml:"end,omitempty" require:"true"`
  // {"dictionary":"belong=WAAP-MS-Ext|dict=waap_timezone","en":"Timezone,default value: GMT+8.","zh_CN":"时区，默认：GMT+8。"}
  Timezone *string `json:"timezone,omitempty" xml:"timezone,omitempty" require:"true"`
}

func (s UpdateSharedRateLimitingRuleRequestRateLimitEffective) String() string {
  return tea.Prettify(s)
}

func (s UpdateSharedRateLimitingRuleRequestRateLimitEffective) GoString() string {
  return s.String()
}

func (s *UpdateSharedRateLimitingRuleRequestRateLimitEffective) SetEffective(v []*string) *UpdateSharedRateLimitingRuleRequestRateLimitEffective {
  s.Effective = v
  return s
}

func (s *UpdateSharedRateLimitingRuleRequestRateLimitEffective) SetStart(v string) *UpdateSharedRateLimitingRuleRequestRateLimitEffective {
  s.Start = &v
  return s
}

func (s *UpdateSharedRateLimitingRuleRequestRateLimitEffective) SetEnd(v string) *UpdateSharedRateLimitingRuleRequestRateLimitEffective {
  s.End = &v
  return s
}

func (s *UpdateSharedRateLimitingRuleRequestRateLimitEffective) SetTimezone(v string) *UpdateSharedRateLimitingRuleRequestRateLimitEffective {
  s.Timezone = &v
  return s
}

type UpdateSharedRateLimitingRuleRequestRateLimitRuleCondition struct {
  // {"en":"IP/CIDR, match type cannot be repeated.","zh_CN":"IP/IP段，匹配类型不可重复。"}
  IpOrIpsConditions []*UpdateSharedRateLimitingRuleRequestRateLimitRuleConditionIpOrIpsConditions `json:"ipOrIpsConditions,omitempty" xml:"ipOrIpsConditions,omitempty" type:"Repeated"`
  // {"en":"Path, match type cannot be repeated.\nWhen the business scenario is API, this matching condition is not supported.","zh_CN":"路径，匹配类型不可重复。\n当业务场景为API业务时不支持此匹配条件。"}
  PathConditions []*UpdateSharedRateLimitingRuleRequestRateLimitRuleConditionPathConditions `json:"pathConditions,omitempty" xml:"pathConditions,omitempty" type:"Repeated"`
  // {"en":"URI, match type cannot be repeated.\nWhen the business scenario is API, this matching condition is not supported.","zh_CN":"URI，匹配类型不可重复。\n当业务场景为API业务时不支持此匹配条件。"}
  UriConditions []*UpdateSharedRateLimitingRuleRequestRateLimitRuleConditionUriConditions `json:"uriConditions,omitempty" xml:"uriConditions,omitempty" type:"Repeated"`
  // {"en":"User Agent, match type cannot be repeated.","zh_CN":"User-Agent，匹配类型不可重复。"}
  UaConditions []*UpdateSharedRateLimitingRuleRequestRateLimitRuleConditionUaConditions `json:"uaConditions,omitempty" xml:"uaConditions,omitempty" type:"Repeated"`
  // {"en":"Request Method.\nWhen the business scenario is API,this matching condition is not supported.","zh_CN":"请求方法，匹配类型不可重复。\n当业务场景为API业务时不支持此匹配条件。"}
  MethodConditions []*UpdateSharedRateLimitingRuleRequestRateLimitRuleConditionMethodConditions `json:"methodConditions,omitempty" xml:"methodConditions,omitempty" type:"Repeated"`
  // {"en":"Referer, match type cannot be repeated.","zh_CN":"Referer，匹配类型不可重复。"}
  RefererConditions []*UpdateSharedRateLimitingRuleRequestRateLimitRuleConditionRefererConditions `json:"refererConditions,omitempty" xml:"refererConditions,omitempty" type:"Repeated"`
  // {"en":"Request Header, match type can be repeated.","zh_CN":"请求头，匹配类型可重复。"}
  HeaderConditions []*UpdateSharedRateLimitingRuleRequestRateLimitRuleConditionHeaderConditions `json:"headerConditions,omitempty" xml:"headerConditions,omitempty" type:"Repeated"`
  // {"en":"Geo, match type cannot be repeated.","zh_CN":"区域，匹配类型不可重复。"}
  AreaConditions []*UpdateSharedRateLimitingRuleRequestRateLimitRuleConditionAreaConditions `json:"areaConditions,omitempty" xml:"areaConditions,omitempty" type:"Repeated"`
  // {"en":"HTTP/S, match type cannot be repeated.","zh_CN":"应用层协议，匹配类型不可重复。"}
  SchemeConditions []*UpdateSharedRateLimitingRuleRequestRateLimitRuleConditionSchemeConditions `json:"schemeConditions,omitempty" xml:"schemeConditions,omitempty" type:"Repeated"`
  // {"en":"Response Code, match type cannot be repeated.","zh_CN":"状态码，匹配类型不可重复。"}
  StatusCodeConditions []*UpdateSharedRateLimitingRuleRequestRateLimitRuleConditionStatusCodeConditions `json:"statusCodeConditions,omitempty" xml:"statusCodeConditions,omitempty" type:"Repeated"`
  // {"en":"JA3 Fingerprint, match type cannot be repeated.","zh_CN":"JA3指纹，匹配类型不可重复。"}
  Ja3Conditions []*UpdateSharedRateLimitingRuleRequestRateLimitRuleConditionJa3Conditions `json:"ja3Conditions,omitempty" xml:"ja3Conditions,omitempty" type:"Repeated"`
  // {"en":"JA4 Fingerprint, match type cannot be repeated.","zh_CN":"JA4指纹，匹配类型不可重复。"}
  Ja4Conditions []*UpdateSharedRateLimitingRuleRequestRateLimitRuleConditionJa4Conditions `json:"ja4Conditions,omitempty" xml:"ja4Conditions,omitempty" type:"Repeated"`
}

func (s UpdateSharedRateLimitingRuleRequestRateLimitRuleCondition) String() string {
  return tea.Prettify(s)
}

func (s UpdateSharedRateLimitingRuleRequestRateLimitRuleCondition) GoString() string {
  return s.String()
}

func (s *UpdateSharedRateLimitingRuleRequestRateLimitRuleCondition) SetIpOrIpsConditions(v []*UpdateSharedRateLimitingRuleRequestRateLimitRuleConditionIpOrIpsConditions) *UpdateSharedRateLimitingRuleRequestRateLimitRuleCondition {
  s.IpOrIpsConditions = v
  return s
}

func (s *UpdateSharedRateLimitingRuleRequestRateLimitRuleCondition) SetPathConditions(v []*UpdateSharedRateLimitingRuleRequestRateLimitRuleConditionPathConditions) *UpdateSharedRateLimitingRuleRequestRateLimitRuleCondition {
  s.PathConditions = v
  return s
}

func (s *UpdateSharedRateLimitingRuleRequestRateLimitRuleCondition) SetUriConditions(v []*UpdateSharedRateLimitingRuleRequestRateLimitRuleConditionUriConditions) *UpdateSharedRateLimitingRuleRequestRateLimitRuleCondition {
  s.UriConditions = v
  return s
}

func (s *UpdateSharedRateLimitingRuleRequestRateLimitRuleCondition) SetUaConditions(v []*UpdateSharedRateLimitingRuleRequestRateLimitRuleConditionUaConditions) *UpdateSharedRateLimitingRuleRequestRateLimitRuleCondition {
  s.UaConditions = v
  return s
}

func (s *UpdateSharedRateLimitingRuleRequestRateLimitRuleCondition) SetMethodConditions(v []*UpdateSharedRateLimitingRuleRequestRateLimitRuleConditionMethodConditions) *UpdateSharedRateLimitingRuleRequestRateLimitRuleCondition {
  s.MethodConditions = v
  return s
}

func (s *UpdateSharedRateLimitingRuleRequestRateLimitRuleCondition) SetRefererConditions(v []*UpdateSharedRateLimitingRuleRequestRateLimitRuleConditionRefererConditions) *UpdateSharedRateLimitingRuleRequestRateLimitRuleCondition {
  s.RefererConditions = v
  return s
}

func (s *UpdateSharedRateLimitingRuleRequestRateLimitRuleCondition) SetHeaderConditions(v []*UpdateSharedRateLimitingRuleRequestRateLimitRuleConditionHeaderConditions) *UpdateSharedRateLimitingRuleRequestRateLimitRuleCondition {
  s.HeaderConditions = v
  return s
}

func (s *UpdateSharedRateLimitingRuleRequestRateLimitRuleCondition) SetAreaConditions(v []*UpdateSharedRateLimitingRuleRequestRateLimitRuleConditionAreaConditions) *UpdateSharedRateLimitingRuleRequestRateLimitRuleCondition {
  s.AreaConditions = v
  return s
}

func (s *UpdateSharedRateLimitingRuleRequestRateLimitRuleCondition) SetSchemeConditions(v []*UpdateSharedRateLimitingRuleRequestRateLimitRuleConditionSchemeConditions) *UpdateSharedRateLimitingRuleRequestRateLimitRuleCondition {
  s.SchemeConditions = v
  return s
}

func (s *UpdateSharedRateLimitingRuleRequestRateLimitRuleCondition) SetStatusCodeConditions(v []*UpdateSharedRateLimitingRuleRequestRateLimitRuleConditionStatusCodeConditions) *UpdateSharedRateLimitingRuleRequestRateLimitRuleCondition {
  s.StatusCodeConditions = v
  return s
}

func (s *UpdateSharedRateLimitingRuleRequestRateLimitRuleCondition) SetJa3Conditions(v []*UpdateSharedRateLimitingRuleRequestRateLimitRuleConditionJa3Conditions) *UpdateSharedRateLimitingRuleRequestRateLimitRuleCondition {
  s.Ja3Conditions = v
  return s
}

func (s *UpdateSharedRateLimitingRuleRequestRateLimitRuleCondition) SetJa4Conditions(v []*UpdateSharedRateLimitingRuleRequestRateLimitRuleConditionJa4Conditions) *UpdateSharedRateLimitingRuleRequestRateLimitRuleCondition {
  s.Ja4Conditions = v
  return s
}

type UpdateSharedRateLimitingRuleRequestRateLimitRuleConditionIpOrIpsConditions struct     {
  // {"en":"Match type.\nEQUAL:Equals\nNOT_EQUAL:Does not equal","zh_CN":"匹配类型。\nEQUAL：等于\nNOT_EQUAL：不等于","exampleValue":"EQUAL,NOT_EQUAL"}
  MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
  // {"en":"IP/CIDR, maximum 500 IP/CIDR.","zh_CN":"IP/IP段，最多500个IP/IP段。"}
  IpOrIps []*string `json:"ipOrIps,omitempty" xml:"ipOrIps,omitempty" require:"true" type:"Repeated"`
}

func (s UpdateSharedRateLimitingRuleRequestRateLimitRuleConditionIpOrIpsConditions) String() string {
  return tea.Prettify(s)
}

func (s UpdateSharedRateLimitingRuleRequestRateLimitRuleConditionIpOrIpsConditions) GoString() string {
  return s.String()
}

func (s *UpdateSharedRateLimitingRuleRequestRateLimitRuleConditionIpOrIpsConditions) SetMatchType(v string) *UpdateSharedRateLimitingRuleRequestRateLimitRuleConditionIpOrIpsConditions {
  s.MatchType = &v
  return s
}

func (s *UpdateSharedRateLimitingRuleRequestRateLimitRuleConditionIpOrIpsConditions) SetIpOrIps(v []*string) *UpdateSharedRateLimitingRuleRequestRateLimitRuleConditionIpOrIpsConditions {
  s.IpOrIps = v
  return s
}

type UpdateSharedRateLimitingRuleRequestRateLimitRuleConditionPathConditions struct     {
  // {"en":"Match type.\nEQUAL: Equals, path case sensitive\nNOT_EQUAL: Does not equal, path case sensitive\nCONTAIN: Contains, path case insensitive\nNOT_CONTAIN: Does not Contains, path case insensitive\nREGEX: Regex match, path case insensitive\nNOT_REGEX: Regular does not match, path case sensitive\nSTART_WITH: Starts with, path case sensitive\nEND_WITH: Ends with, path case sensitive\nWILDCARD: Wildcard matches, path case sensitive,* represents zero or more arbitrary characters, ? represents any single character.\nNOT_WILDCARD: Wildcard does not match, path case sensitive,* represents zero or more arbitrary characters, ? represents any single character","zh_CN":"匹配类型。\nEQUAL：等于，路径大小写敏感\nNOT_EQUAL：不等于，路径大小写敏感\nCONTAIN：包含，路径大小写不敏感\nNOT_CONTAIN：不包含，路径大小写不敏感\nREGEX：匹配正则，路径大小写不敏感\nNOT_REGEX：正则不匹配，路径大小写不敏感\nSTART_WITH：开头是，路径大小写不敏感\nEND_WITH：结尾是，路径大小写不敏感\nWILDCARD：通配符匹配，路径大小写不敏感，*代表零个或多个任意字符，?代表任意单个字符\nNOT_WILDCARD：通配符不匹配，路径大小写不敏感，*代表零个或多个任意字符，?代表任意单个字符","exampleValue":"EQUAL,NOT_EQUAL,CONTAIN,NOT_CONTAIN,REGEX,NOT_REGEX,START_WITH,END_WITH,WILDCARD,NOT_WILDCARD"}
  MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
  // {"en":"Path.\nWhen match type is EQUAL/NOT_EQUAL/START_WITH/END_WITH, path needs to start with \"/\", and no parameters.\nWhen the match type is REGEX/NOT_REGEX, only one value is allowed.\nExample: /test.html.","zh_CN":"路径。\n当匹配类型为等于/不等于/开头是/结尾是，路径必须以“/”开头，不含参数。\n当匹配类型为正则/正则不匹配，则只允许只有一个值。\n示例：/test.html。"}
  UpdateSharedRateLimitingRulePaths []*string `json:"paths,omitempty" xml:"paths,omitempty" require:"true" type:"Repeated"`
}

func (s UpdateSharedRateLimitingRuleRequestRateLimitRuleConditionPathConditions) String() string {
  return tea.Prettify(s)
}

func (s UpdateSharedRateLimitingRuleRequestRateLimitRuleConditionPathConditions) GoString() string {
  return s.String()
}

func (s *UpdateSharedRateLimitingRuleRequestRateLimitRuleConditionPathConditions) SetMatchType(v string) *UpdateSharedRateLimitingRuleRequestRateLimitRuleConditionPathConditions {
  s.MatchType = &v
  return s
}

func (s *UpdateSharedRateLimitingRuleRequestRateLimitRuleConditionPathConditions) SetUpdateSharedRateLimitingRulePaths(v []*string) *UpdateSharedRateLimitingRuleRequestRateLimitRuleConditionPathConditions {
  s.UpdateSharedRateLimitingRulePaths = v
  return s
}

type UpdateSharedRateLimitingRuleRequestRateLimitRuleConditionUriConditions struct     {
  // {"en":"Match type.\nEQUAL: Equals, URI case sensitive\nNOT_EQUAL: Does not equal, URI case sensitive\nCONTAIN: Contains, URI case insensitive\nNOT_CONTAIN: Does not Contains, URI case insensitive\nREGEX: Regex match, URI case insensitive\nNOT_REGEX: Regular does not match, URI case insensitive\nSTART_WITH: Starts with, URI case insensitive\nEND_WITH: Ends with, URI case insensitive\nWILDCARD: Wildcard matches, URI case insensitive,* represents zero or more arbitrary characters, ? represents any single character\nNOT_WILDCARD: Wildcard does not match, URI case insensitive,* represents zero or more arbitrary characters, ? represents any single character","zh_CN":"匹配类型。\nEQUAL：等于，URI大小写敏感\nNOT_EQUAL：不等于，URI大小写敏感\nCONTAIN：包含，URI大小写不敏感\nNOT_CONTAIN：不包含，URI大小写不敏感\nREGEX：匹配正则，URI大小写不敏感\nNOT_REGEX：正则不匹配，URI大小写不敏感\nSTART_WITH：开头是，URI大小写不敏感\nEND_WITH：结尾是，URI大小写不敏感\nWILDCARD：通配符匹配，URI大小写不敏感，*代表零个或多个任意字符，?代表任意单个字符\nNOT_WILDCARD：通配符不匹配，URI大小写不敏感，*代表零个或多个任意字符，?代表任意单个字符","exampleValue":"EQUAL,NOT_EQUAL,CONTAIN,NOT_CONTAIN,REGEX,NOT_REGEX,START_WITH,END_WITH,WILDCARD,NOT_WILDCARD"}
  MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
  // {"en":"URI.\nWhen match type is EQUAL/NOT_EQUAL/START_WITH/END_WITH, uri needs to start with \"/\", and includes parameters.\nWhen the match type is REGEX/NOT_REGEX, only one value is allowed.\nExample: /test.html?id=1.","zh_CN":"URI。\n当匹配类型为等于/不等于/开头是/结尾是，URI必须以”/“开头，含参数。\n当匹配类型为正则/正则不匹配，则只允许只有一个值。\n示例：/test.html?id=1。"}
  Uri []*string `json:"uri,omitempty" xml:"uri,omitempty" require:"true" type:"Repeated"`
}

func (s UpdateSharedRateLimitingRuleRequestRateLimitRuleConditionUriConditions) String() string {
  return tea.Prettify(s)
}

func (s UpdateSharedRateLimitingRuleRequestRateLimitRuleConditionUriConditions) GoString() string {
  return s.String()
}

func (s *UpdateSharedRateLimitingRuleRequestRateLimitRuleConditionUriConditions) SetMatchType(v string) *UpdateSharedRateLimitingRuleRequestRateLimitRuleConditionUriConditions {
  s.MatchType = &v
  return s
}

func (s *UpdateSharedRateLimitingRuleRequestRateLimitRuleConditionUriConditions) SetUri(v []*string) *UpdateSharedRateLimitingRuleRequestRateLimitRuleConditionUriConditions {
  s.Uri = v
  return s
}

type UpdateSharedRateLimitingRuleRequestRateLimitRuleConditionUaConditions struct     {
  // {"en":"Match type.\nEQUAL: Equals, user agent case sensitive\nNOT_EQUAL: Does not equal, user agent case sensitive\nCONTAIN: Contains, user agent case insensitive\nNOT_CONTAIN: Does not Contains, user agent case insensitive\nNONE:Empty or non-existent\nREGEX: Regex match, user agent case insensitive\nNOT_REGEX: Regular does not match, user agent case insensitive\nSTART_WITH: Starts with, user agent case insensitive\nEND_WITH: Ends with, user agent case insensitive\nWILDCARD: Wildcard matches, user agent case insensitive,* represents zero or more arbitrary characters, ? represents any single character\nNOT_WILDCARD: Wildcard does not match, user agent case insensitive,* represents zero or more arbitrary characters, ? represents any single character","zh_CN":"匹配类型。\nEQUAL：等于，User-Agent大小写敏感\nNOT_EQUAL：不等于，User-Agent大小写敏感\nCONTAIN：包含，User-Agent大小写不敏感\nNOT_CONTAIN：不包含，User-Agent大小写不敏感\nNONE：为空或不存在\nREGEX：匹配正则，User-Agent大小写不敏感\nNOT_REGEX：正则不匹配，User-Agent大小写不敏感\nSTART_WITH：开头是，User-Agent大小写不敏感\nEND_WITH：结尾是，User-Agent大小写不敏感\nWILDCARD：通配符匹配，User-Agent大小写不敏感，*代表零个或多个任意字符，?代表任意单个字符\nNOT_WILDCARD：通配符不匹配，User-Agent大小写不敏感，*代表零个或多个任意字符，?代表任意单个字符","exampleValue":"EQUAL,NOT_EQUAL,CONTAIN,NOT_CONTAIN,NONE,REGEX,NOT_REGEX,START_WITH,END_WITH,WILDCARD,NOT_WILDCARD"}
  MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
  // {"en":"User agent.\nWhen the match type is REGEX/NOT_REGEX, only one value is allowed.\nExample: go-Http-client/1.1.","zh_CN":"User-Agent。\n当匹配类型为正则/正则不匹配，则只允许只有一个值。\n示例：go-Http-client/1.1。"}
  Ua []*string `json:"ua,omitempty" xml:"ua,omitempty" require:"true" type:"Repeated"`
}

func (s UpdateSharedRateLimitingRuleRequestRateLimitRuleConditionUaConditions) String() string {
  return tea.Prettify(s)
}

func (s UpdateSharedRateLimitingRuleRequestRateLimitRuleConditionUaConditions) GoString() string {
  return s.String()
}

func (s *UpdateSharedRateLimitingRuleRequestRateLimitRuleConditionUaConditions) SetMatchType(v string) *UpdateSharedRateLimitingRuleRequestRateLimitRuleConditionUaConditions {
  s.MatchType = &v
  return s
}

func (s *UpdateSharedRateLimitingRuleRequestRateLimitRuleConditionUaConditions) SetUa(v []*string) *UpdateSharedRateLimitingRuleRequestRateLimitRuleConditionUaConditions {
  s.Ua = v
  return s
}

type UpdateSharedRateLimitingRuleRequestRateLimitRuleConditionMethodConditions struct     {
  // {"en":"Match type.\nEQUAL:Equals\nNOT_EQUAL:Does not equal","zh_CN":"匹配类型。\nEQUAL：等于\nNOT_EQUAL：不等于","exampleValue":"EQUAL,NOT_EQUAL"}
  MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
  // {"en":"Request method.\nSupported values: GET/POST/DELETE/PUT/HEAD/OPTIONS/COPY.","zh_CN":"请求方法。\n支持的值：GET/POST/DELETE/PUT/HEAD/OPTIONS/COPY。","exampleValue":"GET,POST,DELETE,PUT,HEAD,OPTIONS,COPY"}
  RequestMethod []*string `json:"requestMethod,omitempty" xml:"requestMethod,omitempty" require:"true" type:"Repeated"`
}

func (s UpdateSharedRateLimitingRuleRequestRateLimitRuleConditionMethodConditions) String() string {
  return tea.Prettify(s)
}

func (s UpdateSharedRateLimitingRuleRequestRateLimitRuleConditionMethodConditions) GoString() string {
  return s.String()
}

func (s *UpdateSharedRateLimitingRuleRequestRateLimitRuleConditionMethodConditions) SetMatchType(v string) *UpdateSharedRateLimitingRuleRequestRateLimitRuleConditionMethodConditions {
  s.MatchType = &v
  return s
}

func (s *UpdateSharedRateLimitingRuleRequestRateLimitRuleConditionMethodConditions) SetRequestMethod(v []*string) *UpdateSharedRateLimitingRuleRequestRateLimitRuleConditionMethodConditions {
  s.RequestMethod = v
  return s
}

type UpdateSharedRateLimitingRuleRequestRateLimitRuleConditionRefererConditions struct     {
  // {"en":"Referer.\nWhen the match type is REGEX/NOT_REGEX, only one value is allowed.\nExample: http://test.com.","zh_CN":"Referer。\n当匹配类型为正则/正则不匹配，则只允许只有一个值。\n示例：http://test.com。"}
  Referer []*string `json:"referer,omitempty" xml:"referer,omitempty" require:"true" type:"Repeated"`
  // {"en":"Match type.\nEQUAL: Equals, referer case sensitive\nNOT_EQUAL: Does not equal, referer case sensitive\nCONTAIN: Contains, referer case insensitive\nNOT_CONTAIN: Does not Contains, referer case insensitive\nNONE:Empty or non-existent\nREGEX: Regex match, referer case insensitive\nNOT_REGEX: Regular does not match, referer case insensitive\nSTART_WITH: Starts with, referer case insensitive\nEND_WITH: Ends with, referer case insensitive\nWILDCARD: Wildcard matches, referer case insensitive,* represents zero or more arbitrary characters, ? represents any single characte\nNOT_WILDCARD: Wildcard does not match, referer case insensitive,* represents zero or more arbitrary characters, ? represents any single character","zh_CN":"匹配类型。\nEQUAL：等于，referer大小写敏感\nNOT_EQUAL：不等于，referer大小写敏感\nCONTAIN：包含，referer大小写不敏感\nNOT_CONTAIN：不包含，referer大小写不敏感\nNONE：为空或不存在\nREGEX：匹配正则，referer大小写不敏感\nNOT_REGEX：正则不匹配，referer大小写不敏感\nSTART_WITH：开头是，referer大小写不敏感\nEND_WITH：结尾是，referer大小写不敏感\nWILDCARD：通配符匹配，referer大小写不敏感，*代表零个或多个任意字符，?代表任意单个字符\nNOT_WILDCARD：通配符不匹配，referer大小写不敏感，*代表零个或多个任意字符，?代表任意单个字符","exampleValue":"EQUAL,NOT_EQUAL,CONTAIN,NOT_CONTAIN,NONE,REGEX,NOT_REGEX,START_WITH,END_WITH,WILDCARD,NOT_WILDCARD"}
  MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
}

func (s UpdateSharedRateLimitingRuleRequestRateLimitRuleConditionRefererConditions) String() string {
  return tea.Prettify(s)
}

func (s UpdateSharedRateLimitingRuleRequestRateLimitRuleConditionRefererConditions) GoString() string {
  return s.String()
}

func (s *UpdateSharedRateLimitingRuleRequestRateLimitRuleConditionRefererConditions) SetReferer(v []*string) *UpdateSharedRateLimitingRuleRequestRateLimitRuleConditionRefererConditions {
  s.Referer = v
  return s
}

func (s *UpdateSharedRateLimitingRuleRequestRateLimitRuleConditionRefererConditions) SetMatchType(v string) *UpdateSharedRateLimitingRuleRequestRateLimitRuleConditionRefererConditions {
  s.MatchType = &v
  return s
}

type UpdateSharedRateLimitingRuleRequestRateLimitRuleConditionHeaderConditions struct     {
  // {"en":"Match type.EQUAL: Equals, request header values case sensitiveNOT_EQUAL: Does not equal, request header values case sensitiveCONTAIN: Contains, request header values case insensitiveNOT_CONTAIN: Does not Contains, request header values case insensitiveNONE: Empty or non-existentREGEX: Regex match, request header values case insensitiveNOT_REGEX: Regular does not match, request header values case insensitiveSTART_WITH: Starts with, request header values case insensitiveEND_WITH: Ends with, request header values case insensitiveWILDCARD: Wildcard matches, request header values case insensitive,* represents zero or more arbitrary characters, ? represents any single characterNOT_WILDCARD: Wildcard does not match, request header values case insensitive,* represents zero or more arbitrary characters, ? represents any single character","zh_CN":"匹配类型。EQUAL：等于，头部值大小写敏感NOT_EQUAL：不等于，头部值大小写敏感CONTAIN：包含，头部值大小写不敏感NOT_CONTAIN：不包含，头部值大小写不敏感NONE：为空或不存在REGEX：匹配正则，头部值大小写不敏感NOT_REGEX：正则不匹配，头部值大小写不敏感START_WITH：开头是，头部值大小写不敏感END_WITH：结尾是，头部值大小写不敏感WILDCARD：通配符匹配，头部值大小写不敏感，*代表零个或多个任意字符，?代表任意单个字符NOT_WILDCARD：通配符不匹配，头部值大小写不敏感，*代表零个或多个任意字符，?代表任意单个字符","exampleValue":"EQUAL,NOT_EQUAL,CONTAIN,NOT_CONTAIN,NONE,REGEX,NOT_REGEX,START_WITH,END_WITH,WILDCARD,NOT_WILDCARD"}
  MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
  // {"en":"Header value.\nWhen the match type is REGEX/NOT_REGEX, only one value is allowed.","zh_CN":"头部值。\n当匹配类型为正则/正则不匹配，则只允许只有一个值。"}
  ValueList []*string `json:"valueList,omitempty" xml:"valueList,omitempty" require:"true" type:"Repeated"`
  // {"en":"Header name,case insensitive,up to 100 characters.\nExample: Accept.","zh_CN":"头部名称，大小写不敏感，最多100个字符。\n示例：Accept。"}
  Key *string `json:"key,omitempty" xml:"key,omitempty" require:"true"`
}

func (s UpdateSharedRateLimitingRuleRequestRateLimitRuleConditionHeaderConditions) String() string {
  return tea.Prettify(s)
}

func (s UpdateSharedRateLimitingRuleRequestRateLimitRuleConditionHeaderConditions) GoString() string {
  return s.String()
}

func (s *UpdateSharedRateLimitingRuleRequestRateLimitRuleConditionHeaderConditions) SetMatchType(v string) *UpdateSharedRateLimitingRuleRequestRateLimitRuleConditionHeaderConditions {
  s.MatchType = &v
  return s
}

func (s *UpdateSharedRateLimitingRuleRequestRateLimitRuleConditionHeaderConditions) SetValueList(v []*string) *UpdateSharedRateLimitingRuleRequestRateLimitRuleConditionHeaderConditions {
  s.ValueList = v
  return s
}

func (s *UpdateSharedRateLimitingRuleRequestRateLimitRuleConditionHeaderConditions) SetKey(v string) *UpdateSharedRateLimitingRuleRequestRateLimitRuleConditionHeaderConditions {
  s.Key = &v
  return s
}

type UpdateSharedRateLimitingRuleRequestRateLimitRuleConditionAreaConditions struct     {
  // {"en":"Match type.\nEQUAL:Equals\nNOT_EQUAL:Does not equal","zh_CN":"匹配类型。\nEQUAL：等于\nNOT_EQUAL：不等于","exampleValue":"EQUAL,NOT_EQUAL"}
  MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
  // {"dictionary":"belong=WAAP-MS-Ext|dict=waap_areaCityAndCountry","en":"Geo.","zh_CN":"区域。"}
  Areas []*string `json:"areas,omitempty" xml:"areas,omitempty" require:"true" type:"Repeated"`
}

func (s UpdateSharedRateLimitingRuleRequestRateLimitRuleConditionAreaConditions) String() string {
  return tea.Prettify(s)
}

func (s UpdateSharedRateLimitingRuleRequestRateLimitRuleConditionAreaConditions) GoString() string {
  return s.String()
}

func (s *UpdateSharedRateLimitingRuleRequestRateLimitRuleConditionAreaConditions) SetMatchType(v string) *UpdateSharedRateLimitingRuleRequestRateLimitRuleConditionAreaConditions {
  s.MatchType = &v
  return s
}

func (s *UpdateSharedRateLimitingRuleRequestRateLimitRuleConditionAreaConditions) SetAreas(v []*string) *UpdateSharedRateLimitingRuleRequestRateLimitRuleConditionAreaConditions {
  s.Areas = v
  return s
}

type UpdateSharedRateLimitingRuleRequestRateLimitRuleConditionSchemeConditions struct     {
  // {"en":"Match type.\nEQUAL:Equals\nNOT_EQUAL:Does not equal","zh_CN":"匹配类型。\nEQUAL：等于\nNOT_EQUAL：不等于","exampleValue":"EQUAL,NOT_EQUAL"}
  MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
  // {"en":"HTTP/S.\nSupported values: HTTP/HTTPS.","zh_CN":"应用层协议。\n支持的值：HTTP/HTTPS。","exampleValue":"HTTP,HTTPS"}
  Scheme []*string `json:"scheme,omitempty" xml:"scheme,omitempty" require:"true" type:"Repeated"`
}

func (s UpdateSharedRateLimitingRuleRequestRateLimitRuleConditionSchemeConditions) String() string {
  return tea.Prettify(s)
}

func (s UpdateSharedRateLimitingRuleRequestRateLimitRuleConditionSchemeConditions) GoString() string {
  return s.String()
}

func (s *UpdateSharedRateLimitingRuleRequestRateLimitRuleConditionSchemeConditions) SetMatchType(v string) *UpdateSharedRateLimitingRuleRequestRateLimitRuleConditionSchemeConditions {
  s.MatchType = &v
  return s
}

func (s *UpdateSharedRateLimitingRuleRequestRateLimitRuleConditionSchemeConditions) SetScheme(v []*string) *UpdateSharedRateLimitingRuleRequestRateLimitRuleConditionSchemeConditions {
  s.Scheme = v
  return s
}

type UpdateSharedRateLimitingRuleRequestRateLimitRuleConditionStatusCodeConditions struct     {
  // {"en":"Match type.\nEQUAL:Equals\nNOT_EQUAL:Does not equal","zh_CN":"匹配类型。\nEQUAL：等于\nNOT_EQUAL：不等于","exampleValue":"EQUAL,NOT_EQUAL"}
  MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
  // {"en":"Response Code.","zh_CN":"状态码。"}
  StatusCode []*string `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true" type:"Repeated"`
}

func (s UpdateSharedRateLimitingRuleRequestRateLimitRuleConditionStatusCodeConditions) String() string {
  return tea.Prettify(s)
}

func (s UpdateSharedRateLimitingRuleRequestRateLimitRuleConditionStatusCodeConditions) GoString() string {
  return s.String()
}

func (s *UpdateSharedRateLimitingRuleRequestRateLimitRuleConditionStatusCodeConditions) SetMatchType(v string) *UpdateSharedRateLimitingRuleRequestRateLimitRuleConditionStatusCodeConditions {
  s.MatchType = &v
  return s
}

func (s *UpdateSharedRateLimitingRuleRequestRateLimitRuleConditionStatusCodeConditions) SetStatusCode(v []*string) *UpdateSharedRateLimitingRuleRequestRateLimitRuleConditionStatusCodeConditions {
  s.StatusCode = v
  return s
}

type UpdateSharedRateLimitingRuleRequestRateLimitRuleConditionJa3Conditions struct     {
  // {"en":"Match type.\nEQUAL: Equals\nNOT_EQUAL: Does not equal","zh_CN":"匹配类型。\nEQUAL：等于 \nNOT_EQUAL：不等于","exampleValue":"EQUAL,NOT_EQUAL"}
  MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
  // {"en":"JA3 Fingerprint List, maximum 300 JA3 Fingerprint.\nWhen the match type is EQUAL/NOT_EQUAL, each item's character length must be 32 and can only include numbers and lowercase letters.","zh_CN":"JA3指纹列表，最多300个JA3指纹。\n当匹配类型为等于/不等于时，每项字符长度必须为32，且仅限数字和小写字母。"}
  Ja3List []*string `json:"ja3List,omitempty" xml:"ja3List,omitempty" require:"true" type:"Repeated"`
}

func (s UpdateSharedRateLimitingRuleRequestRateLimitRuleConditionJa3Conditions) String() string {
  return tea.Prettify(s)
}

func (s UpdateSharedRateLimitingRuleRequestRateLimitRuleConditionJa3Conditions) GoString() string {
  return s.String()
}

func (s *UpdateSharedRateLimitingRuleRequestRateLimitRuleConditionJa3Conditions) SetMatchType(v string) *UpdateSharedRateLimitingRuleRequestRateLimitRuleConditionJa3Conditions {
  s.MatchType = &v
  return s
}

func (s *UpdateSharedRateLimitingRuleRequestRateLimitRuleConditionJa3Conditions) SetJa3List(v []*string) *UpdateSharedRateLimitingRuleRequestRateLimitRuleConditionJa3Conditions {
  s.Ja3List = v
  return s
}

type UpdateSharedRateLimitingRuleRequestRateLimitRuleConditionJa4Conditions struct     {
  // {"en":"Match type. \nEQUAL: Equals\nNOT_EQUAL: Does not equal\nCONTAIN: Contains\nNOT_CONTAIN: Does not Contains\nSTART_WITH: Starts with\nEND_WITH: Ends with\nWILDCARD: Wildcard matches, ** represents zero or more arbitrary characters, ? represents any single character\nNOT_WILDCARD: Wildcard does not match, ** represents zero or more arbitrary characters, ? represents any single character","zh_CN":"匹配类型。\nEQUAL：等于\nNOT_EQUAL：不等于\nCONTAIN：包含\nNOT_CONTAIN：不包含\nSTART_WITH：开头是\nEND_WITH：结尾是\nWILDCARD：通配符匹配，*代表零个或多个任意字符，?代表任意单个字符\nNOT_WILDCARD：通配符不匹配，*代表零个或多个任意字符，?代表任意单个字符","exampleValue":"EQUAL,NOT_EQUAL,CONTAIN,NOT_CONTAIN,START_WITH,END_WITH,WILDCARD,NOT_WILDCARD"}
  MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
  // {"en":"JA4 Fingerprint List, maximum 300 JA4 Fingerprint.\nWhen the match type is EQUAL/NOT_EQUAL, each item's format must be 10 characters + 12 characters + 12 characters, separated by underscores, and can only include underscores, numbers, and lowercase letters.\nWhen the match type is CONTAIN/NOT_CONTAIN/START_WITH/END_WITH, each item is only allowed to include underscores, numbers, and lowercase letters.\nWhen the match type is WILDCARD/NOT_WILDCARD, each item, aside from  ** and ?, is only allowed to include underscores, numbers, and lowercase letters.","zh_CN":"JA4指纹列表，最多300个JA4指纹。\n当匹配类型为等于/不等于时，每项格式必须为10位字符+12位字符+12位字符，中间以“_”分隔，且仅限下划线、数字和小写字母。\n当匹配类型为包含/不包含/开头是/结尾是时，每项只允许输入下划线、数字和小写字母。\n当匹配类型为通配符匹配/通配符不匹配时，每项除*和?外，只允许输入下划线、数字和小写字母。"}
  Ja4List []*string `json:"ja4List,omitempty" xml:"ja4List,omitempty" require:"true" type:"Repeated"`
}

func (s UpdateSharedRateLimitingRuleRequestRateLimitRuleConditionJa4Conditions) String() string {
  return tea.Prettify(s)
}

func (s UpdateSharedRateLimitingRuleRequestRateLimitRuleConditionJa4Conditions) GoString() string {
  return s.String()
}

func (s *UpdateSharedRateLimitingRuleRequestRateLimitRuleConditionJa4Conditions) SetMatchType(v string) *UpdateSharedRateLimitingRuleRequestRateLimitRuleConditionJa4Conditions {
  s.MatchType = &v
  return s
}

func (s *UpdateSharedRateLimitingRuleRequestRateLimitRuleConditionJa4Conditions) SetJa4List(v []*string) *UpdateSharedRateLimitingRuleRequestRateLimitRuleConditionJa4Conditions {
  s.Ja4List = v
  return s
}

type UpdateSharedRateLimitingRuleRequestHeader struct {
  // {"en":"Security service type. Please enter a specific service type, if you purchase multiple security services.","zh_CN":"安全服务类型。有使用多个不同的安全服务时，需要填写具体的服务类型。"}
  ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty"`
}

func (s UpdateSharedRateLimitingRuleRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s UpdateSharedRateLimitingRuleRequestHeader) GoString() string {
  return s.String()
}

func (s *UpdateSharedRateLimitingRuleRequestHeader) SetServiceType(v string) *UpdateSharedRateLimitingRuleRequestHeader {
  s.ServiceType = &v
  return s
}

type UpdateSharedRateLimitingRulePaths struct {
}

func (s UpdateSharedRateLimitingRulePaths) String() string {
  return tea.Prettify(s)
}

func (s UpdateSharedRateLimitingRulePaths) GoString() string {
  return s.String()
}

type UpdateSharedRateLimitingRuleParameters struct {
}

func (s UpdateSharedRateLimitingRuleParameters) String() string {
  return tea.Prettify(s)
}

func (s UpdateSharedRateLimitingRuleParameters) GoString() string {
  return s.String()
}

type UpdateSharedRateLimitingRuleResponse struct {
  // {"dictionary":"belong=WAAP-MS-Ext|dict=waap_retCodeEnum","en":"Please refer to the error code for exceptions.","zh_CN":"请参照错误码。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Description.","zh_CN":"描述信息。"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
}

func (s UpdateSharedRateLimitingRuleResponse) String() string {
  return tea.Prettify(s)
}

func (s UpdateSharedRateLimitingRuleResponse) GoString() string {
  return s.String()
}

func (s *UpdateSharedRateLimitingRuleResponse) SetCode(v string) *UpdateSharedRateLimitingRuleResponse {
  s.Code = &v
  return s
}

func (s *UpdateSharedRateLimitingRuleResponse) SetMsg(v string) *UpdateSharedRateLimitingRuleResponse {
  s.Msg = &v
  return s
}

type UpdateSharedRateLimitingRuleResponseHeader struct {
}

func (s UpdateSharedRateLimitingRuleResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s UpdateSharedRateLimitingRuleResponseHeader) GoString() string {
  return s.String()
}




type AssociateShareCustomizeRuleRequest struct {
  // {"en":"Share rule ID.","zh_CN":"共享规则ID。"}
  ShareId *string `json:"shareId,omitempty" xml:"shareId,omitempty" require:"true"`
  // {"en":"List of domain to associate.","zh_CN":"待关联域名列表。"}
  DomainList []*string `json:"domainList,omitempty" xml:"domainList,omitempty" require:"true" type:"Repeated"`
}

func (s AssociateShareCustomizeRuleRequest) String() string {
  return tea.Prettify(s)
}

func (s AssociateShareCustomizeRuleRequest) GoString() string {
  return s.String()
}

func (s *AssociateShareCustomizeRuleRequest) SetShareId(v string) *AssociateShareCustomizeRuleRequest {
  s.ShareId = &v
  return s
}

func (s *AssociateShareCustomizeRuleRequest) SetDomainList(v []*string) *AssociateShareCustomizeRuleRequest {
  s.DomainList = v
  return s
}

type AssociateShareCustomizeRuleRequestHeader struct {
  // {"en":"Security service type. Please enter a specific service type, if you purchase multiple security services.","zh_CN":"安全服务类型。有使用多个不同的安全服务时，需要填写具体的服务类型。"}
  ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty"`
}

func (s AssociateShareCustomizeRuleRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s AssociateShareCustomizeRuleRequestHeader) GoString() string {
  return s.String()
}

func (s *AssociateShareCustomizeRuleRequestHeader) SetServiceType(v string) *AssociateShareCustomizeRuleRequestHeader {
  s.ServiceType = &v
  return s
}

type AssociateShareCustomizeRulePaths struct {
}

func (s AssociateShareCustomizeRulePaths) String() string {
  return tea.Prettify(s)
}

func (s AssociateShareCustomizeRulePaths) GoString() string {
  return s.String()
}

type AssociateShareCustomizeRuleParameters struct {
}

func (s AssociateShareCustomizeRuleParameters) String() string {
  return tea.Prettify(s)
}

func (s AssociateShareCustomizeRuleParameters) GoString() string {
  return s.String()
}

type AssociateShareCustomizeRuleResponse struct {
  // {"dictionary":"belong=WAAP-MS-Ext|dict=waap_retCodeEnum","en":"Please refer to the error code for exceptions.","zh_CN":"请参照错误码。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Description.","zh_CN":"描述信息。"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
}

func (s AssociateShareCustomizeRuleResponse) String() string {
  return tea.Prettify(s)
}

func (s AssociateShareCustomizeRuleResponse) GoString() string {
  return s.String()
}

func (s *AssociateShareCustomizeRuleResponse) SetCode(v string) *AssociateShareCustomizeRuleResponse {
  s.Code = &v
  return s
}

func (s *AssociateShareCustomizeRuleResponse) SetMsg(v string) *AssociateShareCustomizeRuleResponse {
  s.Msg = &v
  return s
}

type AssociateShareCustomizeRuleResponseHeader struct {
}

func (s AssociateShareCustomizeRuleResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s AssociateShareCustomizeRuleResponseHeader) GoString() string {
  return s.String()
}




type CreateSharedWhitelistRuleRequest struct {
  // {"en":"Rule name, maximum 100 characters.Does not support # and &.","zh_CN":"规则名称，最多100个字符。不支持 # 和 &。"}
  RuleName *string `json:"ruleName,omitempty" xml:"ruleName,omitempty" require:"true"`
  // {"en":"Description, maximum 1000 characters.","zh_CN":"描述，最多1000个字符。"}
  Description *string `json:"description,omitempty" xml:"description,omitempty"`
  // {"en":"Match conditions, at least one, at most five.","zh_CN":"匹配条件，至少一个，至多五个。"}
  Conditions *CreateSharedWhitelistRuleRequestConditions `json:"conditions,omitempty" xml:"conditions,omitempty" require:"true" type:"Struct"`
}

func (s CreateSharedWhitelistRuleRequest) String() string {
  return tea.Prettify(s)
}

func (s CreateSharedWhitelistRuleRequest) GoString() string {
  return s.String()
}

func (s *CreateSharedWhitelistRuleRequest) SetRuleName(v string) *CreateSharedWhitelistRuleRequest {
  s.RuleName = &v
  return s
}

func (s *CreateSharedWhitelistRuleRequest) SetDescription(v string) *CreateSharedWhitelistRuleRequest {
  s.Description = &v
  return s
}

func (s *CreateSharedWhitelistRuleRequest) SetConditions(v *CreateSharedWhitelistRuleRequestConditions) *CreateSharedWhitelistRuleRequest {
  s.Conditions = v
  return s
}

type CreateSharedWhitelistRuleRequestConditions struct {
  // {"en":"IP/CIDR match conditions, match type cannot be repeated.","zh_CN":"IP/IP段匹配条件，匹配类型不可重复。"}
  IpOrIpsConditions []*CreateSharedWhitelistRuleRequestConditionsIpOrIpsConditions `json:"ipOrIpsConditions,omitempty" xml:"ipOrIpsConditions,omitempty" type:"Repeated"`
  // {"en":"Path match conditions, match type cannot be repeated.","zh_CN":"路径匹配条件，匹配类型不可重复。"}
  PathConditions []*CreateSharedWhitelistRuleRequestConditionsPathConditions `json:"pathConditions,omitempty" xml:"pathConditions,omitempty" type:"Repeated"`
  // {"en":"URI match conditions, match type cannot be repeated.","zh_CN":"URI匹配条件，匹配类型不可重复。"}
  UriConditions []*CreateSharedWhitelistRuleRequestConditionsUriConditions `json:"uriConditions,omitempty" xml:"uriConditions,omitempty" type:"Repeated"`
  // {"en":"User agent match conditions, match type cannot be repeated.","zh_CN":"User-Agent 匹配条件，匹配类型不可重复。"}
  UaConditions []*CreateSharedWhitelistRuleRequestConditionsUaConditions `json:"uaConditions,omitempty" xml:"uaConditions,omitempty" type:"Repeated"`
  // {"en":"Referer match conditions, match type cannot be repeated.","zh_CN":"Referer 匹配条件，匹配类型不可重复。"}
  RefererConditions []*CreateSharedWhitelistRuleRequestConditionsRefererConditions `json:"refererConditions,omitempty" xml:"refererConditions,omitempty" type:"Repeated"`
  // {"en":"Request header match conditions.","zh_CN":"请求头匹配条件。"}
  HeaderConditions []*CreateSharedWhitelistRuleRequestConditionsHeaderConditions `json:"headerConditions,omitempty" xml:"headerConditions,omitempty" type:"Repeated"`
}

func (s CreateSharedWhitelistRuleRequestConditions) String() string {
  return tea.Prettify(s)
}

func (s CreateSharedWhitelistRuleRequestConditions) GoString() string {
  return s.String()
}

func (s *CreateSharedWhitelistRuleRequestConditions) SetIpOrIpsConditions(v []*CreateSharedWhitelistRuleRequestConditionsIpOrIpsConditions) *CreateSharedWhitelistRuleRequestConditions {
  s.IpOrIpsConditions = v
  return s
}

func (s *CreateSharedWhitelistRuleRequestConditions) SetPathConditions(v []*CreateSharedWhitelistRuleRequestConditionsPathConditions) *CreateSharedWhitelistRuleRequestConditions {
  s.PathConditions = v
  return s
}

func (s *CreateSharedWhitelistRuleRequestConditions) SetUriConditions(v []*CreateSharedWhitelistRuleRequestConditionsUriConditions) *CreateSharedWhitelistRuleRequestConditions {
  s.UriConditions = v
  return s
}

func (s *CreateSharedWhitelistRuleRequestConditions) SetUaConditions(v []*CreateSharedWhitelistRuleRequestConditionsUaConditions) *CreateSharedWhitelistRuleRequestConditions {
  s.UaConditions = v
  return s
}

func (s *CreateSharedWhitelistRuleRequestConditions) SetRefererConditions(v []*CreateSharedWhitelistRuleRequestConditionsRefererConditions) *CreateSharedWhitelistRuleRequestConditions {
  s.RefererConditions = v
  return s
}

func (s *CreateSharedWhitelistRuleRequestConditions) SetHeaderConditions(v []*CreateSharedWhitelistRuleRequestConditionsHeaderConditions) *CreateSharedWhitelistRuleRequestConditions {
  s.HeaderConditions = v
  return s
}

type CreateSharedWhitelistRuleRequestConditionsIpOrIpsConditions struct     {
  // {"en":"Match type.\nEQUAL: Equals\nNOT_EQUAL: Does not equal","zh_CN":"匹配类型。\nEQUAL：等于\nNOT_EQUAL：不等于","exampleValue":"EQUAL,NOT_EQUAL"}
  MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
  // {"en":"IP/CIDR, maximum 500 IP/CIDR.","zh_CN":"IP/IP段，最多500个IP/IP段。"}
  IpOrIps []*string `json:"ipOrIps,omitempty" xml:"ipOrIps,omitempty" require:"true" type:"Repeated"`
}

func (s CreateSharedWhitelistRuleRequestConditionsIpOrIpsConditions) String() string {
  return tea.Prettify(s)
}

func (s CreateSharedWhitelistRuleRequestConditionsIpOrIpsConditions) GoString() string {
  return s.String()
}

func (s *CreateSharedWhitelistRuleRequestConditionsIpOrIpsConditions) SetMatchType(v string) *CreateSharedWhitelistRuleRequestConditionsIpOrIpsConditions {
  s.MatchType = &v
  return s
}

func (s *CreateSharedWhitelistRuleRequestConditionsIpOrIpsConditions) SetIpOrIps(v []*string) *CreateSharedWhitelistRuleRequestConditionsIpOrIpsConditions {
  s.IpOrIps = v
  return s
}

type CreateSharedWhitelistRuleRequestConditionsPathConditions struct     {
  // {"en":"Match type.\nEQUAL: Equals, path case sensitive\nNOT_EQUAL: Does not equal, path case sensitive\nCONTAIN: Contains, path case insensitive\nNOT_CONTAIN: Does not Contains, path case insensitive\nREGEX: Regex match, path case insensitive\nNOT_REGEX: Regular does not match, path case sensitive\nSTART_WITH: Starts with, path case sensitive\nEND_WITH: Ends with, path case sensitive\nWILDCARD: Wildcard matches, path case sensitive,* represents zero or more arbitrary characters, ? represents any single character.\nNOT_WILDCARD: Wildcard does not match, path case sensitive,* represents zero or more arbitrary characters, ? represents any single character","zh_CN":"匹配类型。\nEQUAL：等于，路径大小写敏感\nNOT_EQUAL：不等于，路径大小写敏感\nCONTAIN：包含，路径大小写不敏感\nNOT_CONTAIN：不包含，路径大小写不敏感\nREGEX：匹配正则，路径大小写不敏感\nNOT_REGEX：正则不匹配，路径大小写不敏感\nSTART_WITH：开头是，路径大小写不敏感\nEND_WITH：结尾是，路径大小写不敏感\nWILDCARD：通配符匹配，路径大小写不敏感，*代表零个或多个任意字符，?代表任意单个字符\nNOT_WILDCARD：通配符不匹配，路径大小写不敏感，*代表零个或多个任意字符，?代表任意单个字符","exampleValue":"EQUAL,NOT_EQUAL,CONTAIN,NOT_CONTAIN,REGEX,NOT_REGEX,START_WITH,END_WITH,WILDCARD,NOT_WILDCARD"}
  MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
  // {"en":"Path.\nWhen match type is EQUAL/NOT_EQUAL/START_WITH/END_WITH, path needs to start with \"/\", and no parameters.\nWhen the match type is REGEX/NOT_REGEX, only one value is allowed.\nExample: /test.html.","zh_CN":"路径。\n当匹配类型为等于/不等于/开头是/结尾是，路径必须以“/”开头，不含参数。\n当匹配类型为正则/正则不匹配，则只允许只有一个值。\n示例：/test.html。"}
  CreateSharedWhitelistRulePaths []*string `json:"paths,omitempty" xml:"paths,omitempty" require:"true" type:"Repeated"`
}

func (s CreateSharedWhitelistRuleRequestConditionsPathConditions) String() string {
  return tea.Prettify(s)
}

func (s CreateSharedWhitelistRuleRequestConditionsPathConditions) GoString() string {
  return s.String()
}

func (s *CreateSharedWhitelistRuleRequestConditionsPathConditions) SetMatchType(v string) *CreateSharedWhitelistRuleRequestConditionsPathConditions {
  s.MatchType = &v
  return s
}

func (s *CreateSharedWhitelistRuleRequestConditionsPathConditions) SetCreateSharedWhitelistRulePaths(v []*string) *CreateSharedWhitelistRuleRequestConditionsPathConditions {
  s.CreateSharedWhitelistRulePaths = v
  return s
}

type CreateSharedWhitelistRuleRequestConditionsUriConditions struct     {
  // {"en":"Match type.\nEQUAL: Equals, URI case sensitive\nNOT_EQUAL: Does not equal, URI case sensitive\nCONTAIN: Contains, URI case insensitive\nNOT_CONTAIN: Does not Contains, URI case insensitive\nREGEX: Regex match, URI case insensitive\nNOT_REGEX: Regular does not match, URI case insensitive\nSTART_WITH: Starts with, URI case insensitive\nEND_WITH: Ends with, URI case insensitive\nWILDCARD: Wildcard matches, URI case insensitive,* represents zero or more arbitrary characters, ? represents any single character\nNOT_WILDCARD: Wildcard does not match, URI case insensitive,* represents zero or more arbitrary characters, ? represents any single character","zh_CN":"匹配类型。\nEQUAL：等于，URI大小写敏感\nNOT_EQUAL：不等于，URI大小写敏感\nCONTAIN：包含，URI大小写不敏感\nNOT_CONTAIN：不包含，URI大小写不敏感\nREGEX：匹配正则，URI大小写不敏感\nNOT_REGEX：正则不匹配，URI大小写不敏感\nSTART_WITH：开头是，URI大小写不敏感\nEND_WITH：结尾是，URI大小写不敏感\nWILDCARD：通配符匹配，URI大小写不敏感，*代表零个或多个任意字符，?代表任意单个字符\nNOT_WILDCARD：通配符不匹配，URI大小写不敏感，*代表零个或多个任意字符，?代表任意单个字符","exampleValue":"EQUAL,NOT_EQUAL,CONTAIN,NOT_CONTAIN,REGEX,NOT_REGEX,START_WITH,END_WITH,WILDCARD,NOT_WILDCARD"}
  MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
  // {"en":"URI.\nWhen match type is EQUAL/NOT_EQUAL/START_WITH/END_WITH, uri needs to start with \"/\", and includes parameters.\nWhen the match type is REGEX/NOT_REGEX, only one value is allowed.\nExample: /test.html?id=1.","zh_CN":"URI。\n当匹配类型为等于/不等于/开头是/结尾是，URI必须以”/“开头，含参数。\n当匹配类型为正则/正则不匹配，则只允许只有一个值。\n示例：/test.html?id=1。"}
  Uri []*string `json:"uri,omitempty" xml:"uri,omitempty" require:"true" type:"Repeated"`
}

func (s CreateSharedWhitelistRuleRequestConditionsUriConditions) String() string {
  return tea.Prettify(s)
}

func (s CreateSharedWhitelistRuleRequestConditionsUriConditions) GoString() string {
  return s.String()
}

func (s *CreateSharedWhitelistRuleRequestConditionsUriConditions) SetMatchType(v string) *CreateSharedWhitelistRuleRequestConditionsUriConditions {
  s.MatchType = &v
  return s
}

func (s *CreateSharedWhitelistRuleRequestConditionsUriConditions) SetUri(v []*string) *CreateSharedWhitelistRuleRequestConditionsUriConditions {
  s.Uri = v
  return s
}

type CreateSharedWhitelistRuleRequestConditionsUaConditions struct     {
  // {"en":"Match type.\nEQUAL: Equals, user agent case sensitive\nNOT_EQUAL: Does not equal, user agent case sensitive\nCONTAIN: Contains, user agent case insensitive\nNOT_CONTAIN: Does not Contains, user agent case insensitive\nREGEX: Regex match, user agent case insensitive\nNOT_REGEX: Regular does not match, user agent case insensitive\nSTART_WITH: Starts with, user agent case insensitive\nEND_WITH: Ends with, user agent case insensitive\nWILDCARD: Wildcard matches, user agent case insensitive,* represents zero or more arbitrary characters, ? represents any single character\nNOT_WILDCARD: Wildcard does not match, user agent case insensitive,* represents zero or more arbitrary characters, ? represents any single character","zh_CN":"匹配类型。\nEQUAL：等于，User-Agent大小写敏感\nNOT_EQUAL：不等于，User-Agent大小写敏感\nCONTAIN：包含，User-Agent大小写不敏感\nNOT_CONTAIN：不包含，User-Agent大小写不敏感\nREGEX：匹配正则，User-Agent大小写不敏感\nNOT_REGEX：正则不匹配，User-Agent大小写不敏感\nSTART_WITH：开头是，User-Agent大小写不敏感\nEND_WITH：结尾是，User-Agent大小写不敏感\nWILDCARD：通配符匹配，User-Agent大小写不敏感，*代表零个或多个任意字符，?代表任意单个字符\nNOT_WILDCARD：通配符不匹配，User-Agent大小写不敏感，*代表零个或多个任意字符，?代表任意单个字符","exampleValue":"EQUAL,NOT_EQUAL,CONTAIN,NOT_CONTAIN,REGEX,NOT_REGEX,START_WITH,END_WITH,WILDCARD,NOT_WILDCARD"}
  MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
  // {"en":"User agent.\nWhen the match type is REGEX/NOT_REGEX, only one value is allowed.\nExample: go-Http-client/1.1.","zh_CN":"User-Agent。\n当匹配类型为正则/正则不匹配，则只允许只有一个值。\n示例：go-Http-client/1.1。"}
  Ua []*string `json:"ua,omitempty" xml:"ua,omitempty" require:"true" type:"Repeated"`
}

func (s CreateSharedWhitelistRuleRequestConditionsUaConditions) String() string {
  return tea.Prettify(s)
}

func (s CreateSharedWhitelistRuleRequestConditionsUaConditions) GoString() string {
  return s.String()
}

func (s *CreateSharedWhitelistRuleRequestConditionsUaConditions) SetMatchType(v string) *CreateSharedWhitelistRuleRequestConditionsUaConditions {
  s.MatchType = &v
  return s
}

func (s *CreateSharedWhitelistRuleRequestConditionsUaConditions) SetUa(v []*string) *CreateSharedWhitelistRuleRequestConditionsUaConditions {
  s.Ua = v
  return s
}

type CreateSharedWhitelistRuleRequestConditionsRefererConditions struct     {
  // {"en":"Match type.\nEQUAL: Equals, referer case sensitive\nNOT_EQUAL: Does not equal, referer case sensitive\nCONTAIN: Contains, referer case insensitive\nNOT_CONTAIN: Does not Contains, referer case insensitive\nREGEX: Regex match, referer case insensitive\nNOT_REGEX: Regular does not match, referer case insensitive\nSTART_WITH: Starts with, referer case insensitive\nEND_WITH: Ends with, referer case insensitive\nWILDCARD: Wildcard matches, referer case insensitive,* represents zero or more arbitrary characters, ? represents any single characte\nNOT_WILDCARD: Wildcard does not match, referer case insensitive,* represents zero or more arbitrary characters, ? represents any single character","zh_CN":"匹配类型。\nEQUAL：等于，referer大小写敏感\nNOT_EQUAL：不等于，referer大小写敏感\nCONTAIN：包含，referer大小写不敏感\nNOT_CONTAIN：不包含，referer大小写不敏感\nREGEX：匹配正则，referer大小写不敏感\nNOT_REGEX：正则不匹配，referer大小写不敏感\nSTART_WITH：开头是，referer大小写不敏感\nEND_WITH：结尾是，referer大小写不敏感\nWILDCARD：通配符匹配，referer大小写不敏感，*代表零个或多个任意字符，?代表任意单个字符\nNOT_WILDCARD：通配符不匹配，referer大小写不敏感，*代表零个或多个任意字符，?代表任意单个字符","exampleValue":"EQUAL,NOT_EQUAL,CONTAIN,NOT_CONTAIN,REGEX,NOT_REGEX,START_WITH,END_WITH,WILDCARD,NOT_WILDCARD"}
  MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
  // {"en":"Referer.\nWhen the match type is REGEX/NOT_REGEX, only one value is allowed.\nExample: http://test.com.","zh_CN":"Referer。\n当匹配类型为正则/正则不匹配，则只允许只有一个值。\n示例：http://test.com。"}
  Referer []*string `json:"referer,omitempty" xml:"referer,omitempty" require:"true" type:"Repeated"`
}

func (s CreateSharedWhitelistRuleRequestConditionsRefererConditions) String() string {
  return tea.Prettify(s)
}

func (s CreateSharedWhitelistRuleRequestConditionsRefererConditions) GoString() string {
  return s.String()
}

func (s *CreateSharedWhitelistRuleRequestConditionsRefererConditions) SetMatchType(v string) *CreateSharedWhitelistRuleRequestConditionsRefererConditions {
  s.MatchType = &v
  return s
}

func (s *CreateSharedWhitelistRuleRequestConditionsRefererConditions) SetReferer(v []*string) *CreateSharedWhitelistRuleRequestConditionsRefererConditions {
  s.Referer = v
  return s
}

type CreateSharedWhitelistRuleRequestConditionsHeaderConditions struct     {
  // {"en":"Match type.\nEQUAL: Equals, request header values case sensitive\nNOT_EQUAL: Does not equal, request header values case sensitive\nCONTAIN: Contains, request header values case insensitive\nNOT_CONTAIN: Does not Contains, request header values case insensitive\nREGEX: Regex match, request header values case insensitive\nNOT_REGEX: Regular does not match, request header values case insensitive\nSTART_WITH: Starts with, request header values case insensitive\nEND_WITH: Ends with, request header values case insensitive\nWILDCARD: Wildcard matches, request header values case insensitive,* represents zero or more arbitrary characters, ? represents any single character\nNOT_WILDCARD: Wildcard does not match, request header values case insensitive,* represents zero or more arbitrary characters, ? represents any single character","zh_CN":"匹配类型。\nEQUAL：等于，头部值大小写敏感\nNOT_EQUAL：不等于，头部值大小写敏感\nCONTAIN：包含，头部值大小写不敏感\nNOT_CONTAIN：不包含，头部值大小写不敏感\nREGEX：匹配正则，头部值大小写不敏感\nNOT_REGEX：正则不匹配，头部值大小写不敏感\nSTART_WITH：开头是，头部值大小写不敏感\nEND_WITH：结尾是，头部值大小写不敏感\nWILDCARD：通配符匹配，头部值大小写不敏感，*代表零个或多个任意字符，?代表任意单个字符\nNOT_WILDCARD：通配符不匹配，头部值大小写不敏感，*代表零个或多个任意字符，?代表任意单个字符","exampleValue":"EQUAL,NOT_EQUAL,CONTAIN,NOT_CONTAIN,REGEX,NOT_REGEX,START_WITH,END_WITH,WILDCARD,NOT_WILDCARD"}
  MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
  // {"en":"Header name,case insensitive,up to 100 characters.\nExample: Accept.","zh_CN":"头部名称，大小写不敏感，最多100个字符。\n示例：Accept。"}
  Key *string `json:"key,omitempty" xml:"key,omitempty" require:"true"`
  // {"en":"Header value.\nWhen the match type is REGEX/NOT_REGEX, only one value is allowed.","zh_CN":"头部值。\n当匹配类型为正则/正则不匹配，则只允许只有一个值。"}
  ValueList []*string `json:"valueList,omitempty" xml:"valueList,omitempty" require:"true" type:"Repeated"`
}

func (s CreateSharedWhitelistRuleRequestConditionsHeaderConditions) String() string {
  return tea.Prettify(s)
}

func (s CreateSharedWhitelistRuleRequestConditionsHeaderConditions) GoString() string {
  return s.String()
}

func (s *CreateSharedWhitelistRuleRequestConditionsHeaderConditions) SetMatchType(v string) *CreateSharedWhitelistRuleRequestConditionsHeaderConditions {
  s.MatchType = &v
  return s
}

func (s *CreateSharedWhitelistRuleRequestConditionsHeaderConditions) SetKey(v string) *CreateSharedWhitelistRuleRequestConditionsHeaderConditions {
  s.Key = &v
  return s
}

func (s *CreateSharedWhitelistRuleRequestConditionsHeaderConditions) SetValueList(v []*string) *CreateSharedWhitelistRuleRequestConditionsHeaderConditions {
  s.ValueList = v
  return s
}

type CreateSharedWhitelistRuleRequestHeader struct {
  // {"en":"Security service type. Please enter a specific service type, if you purchase multiple security services.","zh_CN":"安全服务类型。有使用多个不同的安全服务时，需要填写具体的服务类型。"}
  ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty"`
}

func (s CreateSharedWhitelistRuleRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s CreateSharedWhitelistRuleRequestHeader) GoString() string {
  return s.String()
}

func (s *CreateSharedWhitelistRuleRequestHeader) SetServiceType(v string) *CreateSharedWhitelistRuleRequestHeader {
  s.ServiceType = &v
  return s
}

type CreateSharedWhitelistRulePaths struct {
}

func (s CreateSharedWhitelistRulePaths) String() string {
  return tea.Prettify(s)
}

func (s CreateSharedWhitelistRulePaths) GoString() string {
  return s.String()
}

type CreateSharedWhitelistRuleParameters struct {
}

func (s CreateSharedWhitelistRuleParameters) String() string {
  return tea.Prettify(s)
}

func (s CreateSharedWhitelistRuleParameters) GoString() string {
  return s.String()
}

type CreateSharedWhitelistRuleResponse struct {
  // {"dictionary":"belong=WAAP-MS-Ext|dict=waap_retCodeEnum","en":"Please refer to the error code for exceptions.","zh_CN":"请参照错误码。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Description.","zh_CN":"描述信息。"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"Rule ID.","zh_CN":"规则ID。"}
  Data *string `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s CreateSharedWhitelistRuleResponse) String() string {
  return tea.Prettify(s)
}

func (s CreateSharedWhitelistRuleResponse) GoString() string {
  return s.String()
}

func (s *CreateSharedWhitelistRuleResponse) SetCode(v string) *CreateSharedWhitelistRuleResponse {
  s.Code = &v
  return s
}

func (s *CreateSharedWhitelistRuleResponse) SetMsg(v string) *CreateSharedWhitelistRuleResponse {
  s.Msg = &v
  return s
}

func (s *CreateSharedWhitelistRuleResponse) SetData(v string) *CreateSharedWhitelistRuleResponse {
  s.Data = &v
  return s
}

type CreateSharedWhitelistRuleResponseHeader struct {
}

func (s CreateSharedWhitelistRuleResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s CreateSharedWhitelistRuleResponseHeader) GoString() string {
  return s.String()
}




