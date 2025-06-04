package sharedconfiguration

// This file is auto-generated, don't edit it. Thanks.
import (
  "github.com/alibabacloud-go/tea/tea"
)

type ListSharedWhitelistRulesRequest struct {
  // {"en":"Rule name, fuzzy query.", "zh_CN":"规则名称，模糊查询。"}
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

type ListSharedWhitelistRulesResponse struct {
  // {'en':'Please refer to the error code for exceptions.', 'zh_CN':'请参照错误码。','dictionary':'belong=WAAP-MS-Ext|dict=waap_retCodeEnum'}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Description.", "zh_CN":"描述信息。"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"Data.", "zh_CN":"出参数据。"}
  Data []*ListSharedWhitelistRulesCommonShareWhitelistVO `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
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

func (s *ListSharedWhitelistRulesResponse) SetData(v []*ListSharedWhitelistRulesCommonShareWhitelistVO) *ListSharedWhitelistRulesResponse {
  s.Data = v
  return s
}

type ListSharedWhitelistRulesCommonShareWhitelistVO struct {
  // {"en":"Rule ID.", "zh_CN":"规则ID。"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {"en":"Rule name.", "zh_CN":"规则名称。"}
  RuleName *string `json:"ruleName,omitempty" xml:"ruleName,omitempty" require:"true"`
  // {"en":"Description.", "zh_CN":"规则描述。"}
  Description *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
  // {"en":"creator.", "zh_CN":"创建者。"}
  Creator *string `json:"creator,omitempty" xml:"creator,omitempty" require:"true"`
  // {"en":"Number of hostnames associated.", "zh_CN":"关联的域名数。"}
  ImportedDomainCount *int32 `json:"importedDomainCount,omitempty" xml:"importedDomainCount,omitempty" require:"true"`
  // {"en":"Associated hostname.", "zh_CN":"关联的域名。"}
  ImportedDomain []*string `json:"importedDomain,omitempty" xml:"importedDomain,omitempty" require:"true" type:"Repeated"`
  // {"en":"Match conditions, At least one, at most five.", "zh_CN":"匹配条件，至少一个，至多五个。"}
  Conditions *ListSharedWhitelistRulesWhitelistRuleCondition `json:"conditions,omitempty" xml:"conditions,omitempty" require:"true"`
  // {"en":"Created date, format: yyyy-MM-dd HH:mm:ss.", "zh_CN":"创建时间，格式：yyyy-MM-dd HH:mm:ss。"}
  CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty" require:"true"`
  // {"en":"Update date,format: yyyy-MM-dd HH:mm:ss.", "zh_CN":"更新时间，格式：yyyy-MM-dd HH:mm:ss。"}
  UpdateTime *string `json:"updateTime,omitempty" xml:"updateTime,omitempty" require:"true"`
}

func (s ListSharedWhitelistRulesCommonShareWhitelistVO) String() string {
  return tea.Prettify(s)
}

func (s ListSharedWhitelistRulesCommonShareWhitelistVO) GoString() string {
  return s.String()
}

func (s *ListSharedWhitelistRulesCommonShareWhitelistVO) SetId(v string) *ListSharedWhitelistRulesCommonShareWhitelistVO {
  s.Id = &v
  return s
}

func (s *ListSharedWhitelistRulesCommonShareWhitelistVO) SetRuleName(v string) *ListSharedWhitelistRulesCommonShareWhitelistVO {
  s.RuleName = &v
  return s
}

func (s *ListSharedWhitelistRulesCommonShareWhitelistVO) SetDescription(v string) *ListSharedWhitelistRulesCommonShareWhitelistVO {
  s.Description = &v
  return s
}

func (s *ListSharedWhitelistRulesCommonShareWhitelistVO) SetCreator(v string) *ListSharedWhitelistRulesCommonShareWhitelistVO {
  s.Creator = &v
  return s
}

func (s *ListSharedWhitelistRulesCommonShareWhitelistVO) SetImportedDomainCount(v int32) *ListSharedWhitelistRulesCommonShareWhitelistVO {
  s.ImportedDomainCount = &v
  return s
}

func (s *ListSharedWhitelistRulesCommonShareWhitelistVO) SetImportedDomain(v []*string) *ListSharedWhitelistRulesCommonShareWhitelistVO {
  s.ImportedDomain = v
  return s
}

func (s *ListSharedWhitelistRulesCommonShareWhitelistVO) SetConditions(v *ListSharedWhitelistRulesWhitelistRuleCondition) *ListSharedWhitelistRulesCommonShareWhitelistVO {
  s.Conditions = v
  return s
}

func (s *ListSharedWhitelistRulesCommonShareWhitelistVO) SetCreateTime(v string) *ListSharedWhitelistRulesCommonShareWhitelistVO {
  s.CreateTime = &v
  return s
}

func (s *ListSharedWhitelistRulesCommonShareWhitelistVO) SetUpdateTime(v string) *ListSharedWhitelistRulesCommonShareWhitelistVO {
  s.UpdateTime = &v
  return s
}

type ListSharedWhitelistRulesWhitelistRuleCondition struct {
  // {"en":"IP/CIDR match conditions.", "zh_CN":"IP/IP段匹配条件。"}
  IpOrIpsConditions []*ListSharedWhitelistRulesIpOrIpsCondition `json:"ipOrIpsConditions,omitempty" xml:"ipOrIpsConditions,omitempty" require:"true" type:"Repeated"`
  // {"en":"Path match conditions.", "zh_CN":"路径匹配条件。"}
  PathConditions []*ListSharedWhitelistRulesPathCondition `json:"pathConditions,omitempty" xml:"pathConditions,omitempty" require:"true" type:"Repeated"`
  // {"en":"URI match conditions.", "zh_CN":"URI匹配条件。"}
  UriConditions []*ListSharedWhitelistRulesUriCondition `json:"uriConditions,omitempty" xml:"uriConditions,omitempty" require:"true" type:"Repeated"`
  // {"en":"User agent match conditions.", "zh_CN":"User-Agent 匹配条件。"}
  UaConditions []*ListSharedWhitelistRulesUaCondition `json:"uaConditions,omitempty" xml:"uaConditions,omitempty" require:"true" type:"Repeated"`
  // {"en":"Referer match conditions.", "zh_CN":"Referer 匹配条件。"}
  RefererConditions []*ListSharedWhitelistRulesRefererCondition `json:"refererConditions,omitempty" xml:"refererConditions,omitempty" require:"true" type:"Repeated"`
  // {"en":"Request header match conditions.", "zh_CN":"请求头匹配条件。"}
  HeaderConditions []*ListSharedWhitelistRulesHeaderCondition `json:"headerConditions,omitempty" xml:"headerConditions,omitempty" require:"true" type:"Repeated"`
}

func (s ListSharedWhitelistRulesWhitelistRuleCondition) String() string {
  return tea.Prettify(s)
}

func (s ListSharedWhitelistRulesWhitelistRuleCondition) GoString() string {
  return s.String()
}

func (s *ListSharedWhitelistRulesWhitelistRuleCondition) SetIpOrIpsConditions(v []*ListSharedWhitelistRulesIpOrIpsCondition) *ListSharedWhitelistRulesWhitelistRuleCondition {
  s.IpOrIpsConditions = v
  return s
}

func (s *ListSharedWhitelistRulesWhitelistRuleCondition) SetPathConditions(v []*ListSharedWhitelistRulesPathCondition) *ListSharedWhitelistRulesWhitelistRuleCondition {
  s.PathConditions = v
  return s
}

func (s *ListSharedWhitelistRulesWhitelistRuleCondition) SetUriConditions(v []*ListSharedWhitelistRulesUriCondition) *ListSharedWhitelistRulesWhitelistRuleCondition {
  s.UriConditions = v
  return s
}

func (s *ListSharedWhitelistRulesWhitelistRuleCondition) SetUaConditions(v []*ListSharedWhitelistRulesUaCondition) *ListSharedWhitelistRulesWhitelistRuleCondition {
  s.UaConditions = v
  return s
}

func (s *ListSharedWhitelistRulesWhitelistRuleCondition) SetRefererConditions(v []*ListSharedWhitelistRulesRefererCondition) *ListSharedWhitelistRulesWhitelistRuleCondition {
  s.RefererConditions = v
  return s
}

func (s *ListSharedWhitelistRulesWhitelistRuleCondition) SetHeaderConditions(v []*ListSharedWhitelistRulesHeaderCondition) *ListSharedWhitelistRulesWhitelistRuleCondition {
  s.HeaderConditions = v
  return s
}

type ListSharedWhitelistRulesIpOrIpsCondition struct {
  // {"en":"Match type.
  // EQUAL: Equals
  // NOT_EQUAL: Does not equal", "zh_CN":"匹配类型。
  // EQUAL：等于
  // NOT_EQUAL：不等于"}
  MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
  // {"en":"IP/CIDR.", "zh_CN":"IP/IP段。"}
  IpOrIps []*string `json:"ipOrIps,omitempty" xml:"ipOrIps,omitempty" require:"true" type:"Repeated"`
}

func (s ListSharedWhitelistRulesIpOrIpsCondition) String() string {
  return tea.Prettify(s)
}

func (s ListSharedWhitelistRulesIpOrIpsCondition) GoString() string {
  return s.String()
}

func (s *ListSharedWhitelistRulesIpOrIpsCondition) SetMatchType(v string) *ListSharedWhitelistRulesIpOrIpsCondition {
  s.MatchType = &v
  return s
}

func (s *ListSharedWhitelistRulesIpOrIpsCondition) SetIpOrIps(v []*string) *ListSharedWhitelistRulesIpOrIpsCondition {
  s.IpOrIps = v
  return s
}

type ListSharedWhitelistRulesPathCondition struct {
  // {"en":"Match type.
  // EQUAL: Equals
  // NOT_EQUAL: Does not equal
  // CONTAIN: Contains
  // NOT_CONTAIN: Does not Contains
  // REGEX: Regex match
  // NOT_REGEX: regular does not match
  // START_WITH: starts with
  // END_WITH: ends with
  // WILDCARD: wildcard matches
  // NOT_WILDCARD: wildcard does not match", "zh_CN":"匹配类型。
  // EQUAL：等于
  // NOT_EQUAL：不等于
  // CONTAIN：包含
  // NOT_CONTAIN：不包含
  // REGEX：匹配正则
  // NOT_REGEX：正则不匹配
  // START_WITH：开头是
  // END_WITH：结尾是
  // WILDCARD：通配符匹配
  // NOT_WILDCARD：通配符不匹配"}
  MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
  // {"en":"Path.", "zh_CN":"路径。"}
  ListSharedWhitelistRulesPaths []*string `json:"paths,omitempty" xml:"paths,omitempty" require:"true" type:"Repeated"`
}

func (s ListSharedWhitelistRulesPathCondition) String() string {
  return tea.Prettify(s)
}

func (s ListSharedWhitelistRulesPathCondition) GoString() string {
  return s.String()
}

func (s *ListSharedWhitelistRulesPathCondition) SetMatchType(v string) *ListSharedWhitelistRulesPathCondition {
  s.MatchType = &v
  return s
}

func (s *ListSharedWhitelistRulesPathCondition) SetListSharedWhitelistRulesPaths(v []*string) *ListSharedWhitelistRulesPathCondition {
  s.ListSharedWhitelistRulesPaths = v
  return s
}

type ListSharedWhitelistRulesUriCondition struct {
  // {"en":"Match type.
  // EQUAL: Equals
  // NOT_EQUAL: Does not equal
  // CONTAIN: Contains
  // NOT_CONTAIN: Does not Contains
  // REGEX: Regex match
  // NOT_REGEX: regular does not match
  // START_WITH: starts with
  // END_WITH: ends with
  // WILDCARD: wildcard matches
  // NOT_WILDCARD: wildcard does not match", "zh_CN":"匹配类型。
  // EQUAL：等于
  // NOT_EQUAL：不等于
  // CONTAIN：包含
  // NOT_CONTAIN：不包含
  // REGEX：匹配正则
  // NOT_REGEX：正则不匹配
  // START_WITH：开头是
  // END_WITH：结尾是
  // WILDCARD：通配符匹配
  // NOT_WILDCARD：通配符不匹配"}
  MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
  // {"en":"URI.", "zh_CN":"URI。"}
  Uri []*string `json:"uri,omitempty" xml:"uri,omitempty" require:"true" type:"Repeated"`
}

func (s ListSharedWhitelistRulesUriCondition) String() string {
  return tea.Prettify(s)
}

func (s ListSharedWhitelistRulesUriCondition) GoString() string {
  return s.String()
}

func (s *ListSharedWhitelistRulesUriCondition) SetMatchType(v string) *ListSharedWhitelistRulesUriCondition {
  s.MatchType = &v
  return s
}

func (s *ListSharedWhitelistRulesUriCondition) SetUri(v []*string) *ListSharedWhitelistRulesUriCondition {
  s.Uri = v
  return s
}

type ListSharedWhitelistRulesUaCondition struct {
  // {"en":"Match type.
  // EQUAL: Equals
  // NOT_EQUAL: Does not equal
  // CONTAIN: Contains
  // NOT_CONTAIN: Does not Contains
  // REGEX: Regex match
  // NOT_REGEX: regular does not match
  // START_WITH: starts with
  // END_WITH: ends with
  // WILDCARD: wildcard matches
  // NOT_WILDCARD: wildcard does not match", "zh_CN":"匹配类型。
  // EQUAL：等于
  // NOT_EQUAL：不等于
  // CONTAIN：包含
  // NOT_CONTAIN：不包含
  // REGEX：匹配正则
  // NOT_REGEX：正则不匹配
  // START_WITH：开头是
  // END_WITH：结尾是
  // WILDCARD：通配符匹配
  // NOT_WILDCARD：通配符不匹配"}
  MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
  // {"en":"User agent.", "zh_CN":"User-Agent。"}
  Ua []*string `json:"ua,omitempty" xml:"ua,omitempty" require:"true" type:"Repeated"`
}

func (s ListSharedWhitelistRulesUaCondition) String() string {
  return tea.Prettify(s)
}

func (s ListSharedWhitelistRulesUaCondition) GoString() string {
  return s.String()
}

func (s *ListSharedWhitelistRulesUaCondition) SetMatchType(v string) *ListSharedWhitelistRulesUaCondition {
  s.MatchType = &v
  return s
}

func (s *ListSharedWhitelistRulesUaCondition) SetUa(v []*string) *ListSharedWhitelistRulesUaCondition {
  s.Ua = v
  return s
}

type ListSharedWhitelistRulesRefererCondition struct {
  // {"en":"Match type.
  // EQUAL: Equals
  // NOT_EQUAL: Does not equal
  // CONTAIN: Contains
  // NOT_CONTAIN: Does not Contains
  // REGEX: Regex match
  // NOT_REGEX: regular does not match
  // START_WITH: starts with
  // END_WITH: ends with
  // WILDCARD: wildcard matches
  // NOT_WILDCARD: wildcard does not match", "zh_CN":"匹配类型。
  // EQUAL：等于
  // NOT_EQUAL：不等于
  // CONTAIN：包含
  // NOT_CONTAIN：不包含
  // REGEX：匹配正则
  // NOT_REGEX：正则不匹配
  // START_WITH：开头是
  // END_WITH：结尾是
  // WILDCARD：通配符匹配
  // NOT_WILDCARD：通配符不匹配"}
  MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
  // {"en":"Referer.", "zh_CN":"Referer。"}
  Referer []*string `json:"referer,omitempty" xml:"referer,omitempty" require:"true" type:"Repeated"`
}

func (s ListSharedWhitelistRulesRefererCondition) String() string {
  return tea.Prettify(s)
}

func (s ListSharedWhitelistRulesRefererCondition) GoString() string {
  return s.String()
}

func (s *ListSharedWhitelistRulesRefererCondition) SetMatchType(v string) *ListSharedWhitelistRulesRefererCondition {
  s.MatchType = &v
  return s
}

func (s *ListSharedWhitelistRulesRefererCondition) SetReferer(v []*string) *ListSharedWhitelistRulesRefererCondition {
  s.Referer = v
  return s
}

type ListSharedWhitelistRulesHeaderCondition struct {
  // {"en":"Match type.
  // EQUAL: Equals
  // NOT_EQUAL: Does not equal
  // CONTAIN: Contains
  // NOT_CONTAIN: Does not Contains
  // REGEX: Regex match
  // NOT_REGEX: regular does not match
  // START_WITH: starts with
  // END_WITH: ends with
  // WILDCARD: wildcard matches
  // NOT_WILDCARD: wildcard does not match", "zh_CN":"匹配类型。
  // EQUAL：等于
  // NOT_EQUAL：不等于
  // CONTAIN：包含
  // NOT_CONTAIN：不包含
  // REGEX：匹配正则
  // NOT_REGEX：正则不匹配
  // START_WITH：开头是
  // END_WITH：结尾是
  // WILDCARD：通配符匹配
  // NOT_WILDCARD：通配符不匹配"}
  MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
  // {"en":"Request header name.", "zh_CN":"头部名称。"}
  Key *string `json:"key,omitempty" xml:"key,omitempty" require:"true"`
  // {"en":"List of request header values.", "zh_CN":"头部值列表。"}
  ValueList []*string `json:"valueList,omitempty" xml:"valueList,omitempty" require:"true" type:"Repeated"`
}

func (s ListSharedWhitelistRulesHeaderCondition) String() string {
  return tea.Prettify(s)
}

func (s ListSharedWhitelistRulesHeaderCondition) GoString() string {
  return s.String()
}

func (s *ListSharedWhitelistRulesHeaderCondition) SetMatchType(v string) *ListSharedWhitelistRulesHeaderCondition {
  s.MatchType = &v
  return s
}

func (s *ListSharedWhitelistRulesHeaderCondition) SetKey(v string) *ListSharedWhitelistRulesHeaderCondition {
  s.Key = &v
  return s
}

func (s *ListSharedWhitelistRulesHeaderCondition) SetValueList(v []*string) *ListSharedWhitelistRulesHeaderCondition {
  s.ValueList = v
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

type ListSharedWhitelistRulesRequestHeader struct {
  // {"zh_CN":"安全服务类型。有使用多个不同的安全服务时，需要填写具体的服务类型。","en":"Security service type. Please enter a specific service type, if you purchase multiple security services.","dictionary":"belong=WAAP-MS-Ext|dict=waap_serviceType"}
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

type ListSharedWhitelistRulesResponseHeader struct {
}

func (s ListSharedWhitelistRulesResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ListSharedWhitelistRulesResponseHeader) GoString() string {
  return s.String()
}




type DisassociateShareCustomizeBotsRequest struct {
  // {"zh_CN":"共享配置ID。","en":"Share configuration ID."}
  ShareId *string `json:"shareId,omitempty" xml:"shareId,omitempty" require:"true"`
  // {"zh_CN":"待取消关联域名列表。","en":"List of hostname to be disassociated."}
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

type DisassociateShareCustomizeBotsResponse struct {
  // {'en':'Please refer to the error code for exceptions.', 'zh_CN':'请参照错误码。','dictionary':'belong=WAAP-MS-Ext|dict=waap_retCodeEnum'}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {'en':'Description.', 'zh_CN':'描述信息。'}
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

type DisassociateShareCustomizeBotsRequestHeader struct {
  // {"zh_CN":"安全服务类型。有使用多个不同的安全服务时，需要填写具体的服务类型。","en":"Security service type. Please enter a specific service type, if you purchase multiple security services.","dictionary":"belong=WAAP-MS-Ext|dict=waap_serviceType"}
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

type DisassociateShareCustomizeBotsResponseHeader struct {
}

func (s DisassociateShareCustomizeBotsResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s DisassociateShareCustomizeBotsResponseHeader) GoString() string {
  return s.String()
}




type CreateSharedRateLimitingRuleRequest struct {
  // {'en':'Rule Name, maximum 50 characters.
  // Does not support special characters and spaces.', 'zh_CN':'规则名称，最多50个字符。
  // 不支持特殊字符和空格。'}
  RuleName *string `json:"ruleName,omitempty" xml:"ruleName,omitempty" require:"true"`
  // {'en':'Description, maximum 200 characters.', 'zh_CN':'规则描述，最多200个字符。'}
  Description *string `json:"description,omitempty" xml:"description,omitempty"`
  // {'en':'Client identifier.
  // IP:Client IP
  // IP_UA:Client IP and User-Agent
  // COOKIE:Cookie
  // IP_COOKIE:Client IP and Cookie
  // HEADER:Request Header
  // When there is a status code in the matching condition,this client identifier is not supported.
  // IP_HEADER:Client IP and Request Header
  // When there is a status code in the matching condition,this client identifier is not supported .', 'zh_CN':'统计粒度。
  // IP：客户端IP
  // IP_UA：客户端IP和User-Agent
  // COOKIE：Cookie
  // IP_COOKIE：客户端IP和Cookie
  // HEADER：请求头，当匹配条件中存在状态码时不支持此统计粒度
  // IP_HEADER：客户端IP和请求头，当匹配条件中存在状态码时不支持此统计粒度'}
  StatisticalItem *string `json:"statisticalItem,omitempty" xml:"statisticalItem,omitempty" require:"true"`
  // {'en':'Statistical key value.
  // When the client identifier is cookie/header value, the corresponding key value needs to be entered.', 'zh_CN':'统计key值。
  // 当统计粒度cookie/header值，需要输入对应的key值。'}
  StatisticsKey *string `json:"statisticsKey,omitempty" xml:"statisticsKey,omitempty"`
  // {'en':'Statistics period, unit: seconds, the range is 1 - 3600.', 'zh_CN':'统计周期，单位：秒，范围为 1 - 3600。'}
  StatisticalPeriod *int32 `json:"statisticalPeriod,omitempty" xml:"statisticalPeriod,omitempty" require:"true"`
  // {'en':'Trigger threshold, unit: times.', 'zh_CN':'触发阈值，单位：次。'}
  TriggerThreshold *int32 `json:"triggerThreshold,omitempty" xml:"triggerThreshold,omitempty" require:"true"`
  // {'en':'Action duration, unit: seconds, the range is 10 - 604800.', 'zh_CN':'处理动作持续时间，单位：秒，范围为 10 - 604800。'}
  InterceptTime *int32 `json:"interceptTime,omitempty" xml:"interceptTime,omitempty" require:"true"`
  // {'en':'Cycle effective status.
  // PERMANENT:All time
  // WITHOUT:Excluded time
  // WITHIN:Selected time', 'zh_CN':'周期生效状态。
  // PERMANENT：永久生效
  // WITHOUT：周期内不生效
  // WITHIN：周期内生效'}
  EffectiveStatus *string `json:"effectiveStatus,omitempty" xml:"effectiveStatus,omitempty" require:"true"`
  // {'en':'Effective time period.
  // When the effective status is effective within the cycle or not effective within the cycle, this field must have a value.', 'zh_CN':'规则生效周期。
  // 生效状态为周期内生效或周期内不生效时，此字段必须有值。'}
  CreateSharedRateLimitingRuleRateLimitEffective *CreateSharedRateLimitingRuleRateLimitEffective `json:"rateLimitEffective,omitempty" xml:"rateLimitEffective,omitempty"`
  // {'en':'Action.
  // NO_USE:Not Used
  // LOG:Log
  // COOKIE:Cookie verification
  // JS_CHECK:Javascript verification
  // DELAY:Delay
  // BLOCK:Deny
  // RESET:Reset Connection
  // Custom response ID:Custom response ID
  // When there is a status code in the matching condition, the supported actions are Log, Deny, and Reset Connection.', 'zh_CN':'处理动作。
  // NO_USE：不使用
  // LOG：监控
  // COOKIE：Cookie校验
  // JS_CHECK：JavaScript校验
  // DELAY：延迟响应
  // BLOCK：拦截
  // RESET：断开连接
  // 自定义响应ID：自定义响应ID
  // 当匹配条件中存在状态码时，支持处理动作为监控、拦截、断开连接。'}
  Action *string `json:"action,omitempty" xml:"action,omitempty" require:"true"`
  // {'en':'Matching conditions.', 'zh_CN':'匹配条件。'}
  RateLimitRuleCondition *CreateSharedRateLimitingRuleShareRateLimitRuleCondition `json:"rateLimitRuleCondition,omitempty" xml:"rateLimitRuleCondition,omitempty" require:"true"`
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

func (s *CreateSharedRateLimitingRuleRequest) SetStatisticalPeriod(v int32) *CreateSharedRateLimitingRuleRequest {
  s.StatisticalPeriod = &v
  return s
}

func (s *CreateSharedRateLimitingRuleRequest) SetTriggerThreshold(v int32) *CreateSharedRateLimitingRuleRequest {
  s.TriggerThreshold = &v
  return s
}

func (s *CreateSharedRateLimitingRuleRequest) SetInterceptTime(v int32) *CreateSharedRateLimitingRuleRequest {
  s.InterceptTime = &v
  return s
}

func (s *CreateSharedRateLimitingRuleRequest) SetEffectiveStatus(v string) *CreateSharedRateLimitingRuleRequest {
  s.EffectiveStatus = &v
  return s
}

func (s *CreateSharedRateLimitingRuleRequest) SetRateLimitEffective(v *CreateSharedRateLimitingRuleRateLimitEffective) *CreateSharedRateLimitingRuleRequest {
  s.CreateSharedRateLimitingRuleRateLimitEffective = v
  return s
}

func (s *CreateSharedRateLimitingRuleRequest) SetAction(v string) *CreateSharedRateLimitingRuleRequest {
  s.Action = &v
  return s
}

func (s *CreateSharedRateLimitingRuleRequest) SetRateLimitRuleCondition(v *CreateSharedRateLimitingRuleShareRateLimitRuleCondition) *CreateSharedRateLimitingRuleRequest {
  s.RateLimitRuleCondition = v
  return s
}

type CreateSharedRateLimitingRuleResponse struct {
  // {'en':'Please refer to the error code for exceptions.', 'zh_CN':'请参照错误码。','dictionary':'belong=WAAP-MS-Ext|dict=waap_retCodeEnum'}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {'en':'Description.', 'zh_CN':'描述信息。'}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {'en':'Rule ID.', 'zh_CN':'规则ID。'}
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

type CreateSharedRateLimitingRuleIpOrIpsCondition struct {
  // {'en':'Match type.
  // EQUAL:Equals
  // NOT_EQUAL:Does not equal', 'zh_CN':'匹配类型。
  // EQUAL：等于
  // NOT_EQUAL：不等于'}
  MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
  // {'en':'IP/CIDR, maximum 500 IP/CIDR.', 'zh_CN':'IP/IP段，最多500个IP/IP段。'}
  IpOrIps []*string `json:"ipOrIps,omitempty" xml:"ipOrIps,omitempty" require:"true" type:"Repeated"`
}

func (s CreateSharedRateLimitingRuleIpOrIpsCondition) String() string {
  return tea.Prettify(s)
}

func (s CreateSharedRateLimitingRuleIpOrIpsCondition) GoString() string {
  return s.String()
}

func (s *CreateSharedRateLimitingRuleIpOrIpsCondition) SetMatchType(v string) *CreateSharedRateLimitingRuleIpOrIpsCondition {
  s.MatchType = &v
  return s
}

func (s *CreateSharedRateLimitingRuleIpOrIpsCondition) SetIpOrIps(v []*string) *CreateSharedRateLimitingRuleIpOrIpsCondition {
  s.IpOrIps = v
  return s
}

type CreateSharedRateLimitingRuleHeaderCondition struct {
  // {"en":"Match type.
  // EQUAL: Equals, request header values case sensitive
  // NOT_EQUAL: Does not equal, request header values case sensitive
  // CONTAIN: Contains, request header values case insensitive
  // NOT_CONTAIN: Does not Contains, request header values case insensitive
  // NONE: Empty or non-existent
  // REGEX: Regex match, request header values case insensitive
  // NOT_REGEX: Regular does not match, request header values case insensitive
  // START_WITH: Starts with, request header values case insensitive
  // END_WITH: Ends with, request header values case insensitive
  // WILDCARD: Wildcard matches, request header values case insensitive, ** represents zero or more arbitrary characters, ? represents any single character
  // NOT_WILDCARD: Wildcard does not match, request header values case insensitive, ** represents zero or more arbitrary characters, ? represents any single character", "zh_CN":"匹配类型。
  // EQUAL：等于，头部值大小写敏感
  // NOT_EQUAL：不等于，头部值大小写敏感
  // CONTAIN：包含，头部值大小写不敏感
  // NOT_CONTAIN：不包含，头部值大小写不敏感
  // NONE：为空或不存在
  // REGEX：匹配正则，头部值大小写不敏感
  // NOT_REGEX：正则不匹配，头部值大小写不敏感
  // START_WITH：开头是，头部值大小写不敏感
  // END_WITH：结尾是，头部值大小写不敏感
  // WILDCARD：通配符匹配，头部值大小写不敏感，*代表零个或多个任意字符，?代表任意单个字符
  // NOT_WILDCARD：通配符不匹配，头部值大小写不敏感，*代表零个或多个任意字符，?代表任意单个字符"}
  MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
  // {'en':'Header name,case insensitive,up to 100 characters.
  // Example: Accept.', 'zh_CN':'头部名称，大小写不敏感，最多100个字符。
  // 示例：Accept。'}
  Key *string `json:"key,omitempty" xml:"key,omitempty" require:"true"`
  // {'en':'Header value.
  // When the match type is REGEX/NOT_REGEX, only one value is allowed.', 'zh_CN':'头部值。
  // 当匹配类型为正则/正则不匹配，则只允许只有一个值。'}
  ValueList []*string `json:"valueList,omitempty" xml:"valueList,omitempty" require:"true" type:"Repeated"`
}

func (s CreateSharedRateLimitingRuleHeaderCondition) String() string {
  return tea.Prettify(s)
}

func (s CreateSharedRateLimitingRuleHeaderCondition) GoString() string {
  return s.String()
}

func (s *CreateSharedRateLimitingRuleHeaderCondition) SetMatchType(v string) *CreateSharedRateLimitingRuleHeaderCondition {
  s.MatchType = &v
  return s
}

func (s *CreateSharedRateLimitingRuleHeaderCondition) SetKey(v string) *CreateSharedRateLimitingRuleHeaderCondition {
  s.Key = &v
  return s
}

func (s *CreateSharedRateLimitingRuleHeaderCondition) SetValueList(v []*string) *CreateSharedRateLimitingRuleHeaderCondition {
  s.ValueList = v
  return s
}

type CreateSharedRateLimitingRuleAreaCondition struct {
  // {'en':'Match type.
  // EQUAL:Equals
  // NOT_EQUAL:Does not equal', 'zh_CN':'匹配类型。
  // EQUAL：等于
  // NOT_EQUAL：不等于'}
  MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
  // {'en':'Geo.', 'zh_CN':'区域。','dictionary':'belong=WAAP-MS-Ext|dict=waap_areaCityAndCountry'}
  Areas []*string `json:"areas,omitempty" xml:"areas,omitempty" require:"true" type:"Repeated"`
}

func (s CreateSharedRateLimitingRuleAreaCondition) String() string {
  return tea.Prettify(s)
}

func (s CreateSharedRateLimitingRuleAreaCondition) GoString() string {
  return s.String()
}

func (s *CreateSharedRateLimitingRuleAreaCondition) SetMatchType(v string) *CreateSharedRateLimitingRuleAreaCondition {
  s.MatchType = &v
  return s
}

func (s *CreateSharedRateLimitingRuleAreaCondition) SetAreas(v []*string) *CreateSharedRateLimitingRuleAreaCondition {
  s.Areas = v
  return s
}

type CreateSharedRateLimitingRuleUriCondition struct {
  // {"en":"Match type.
  // EQUAL: Equals, URI case sensitive
  // NOT_EQUAL: Does not equal, URI case sensitive
  // CONTAIN: Contains, URI case insensitive
  // NOT_CONTAIN: Does not Contains, URI case insensitive
  // REGEX: Regex match, URI case insensitive
  // NOT_REGEX: Regular does not match, URI case insensitive
  // START_WITH: Starts with, URI case insensitive
  // END_WITH: Ends with, URI case insensitive
  // WILDCARD: Wildcard matches, URI case insensitive, ** represents zero or more arbitrary characters, ? represents any single character
  // NOT_WILDCARD: Wildcard does not match, URI case insensitive, ** represents zero or more arbitrary characters, ? represents any single character", "zh_CN":"匹配类型。
  // EQUAL：等于，URI大小写敏感
  // NOT_EQUAL：不等于，URI大小写敏感
  // CONTAIN：包含，URI大小写不敏感
  // NOT_CONTAIN：不包含，URI大小写不敏感
  // REGEX：匹配正则，URI大小写不敏感
  // NOT_REGEX：正则不匹配，URI大小写不敏感
  // START_WITH：开头是，URI大小写不敏感
  // END_WITH：结尾是，URI大小写不敏感
  // WILDCARD：通配符匹配，URI大小写不敏感，*代表零个或多个任意字符，?代表任意单个字符
  // NOT_WILDCARD：通配符不匹配，URI大小写不敏感，*代表零个或多个任意字符，?代表任意单个字符"}
  MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
  // {'en':'URI.
  // When match type is EQUAL/NOT_EQUAL/START_WITH/END_WITH, uri needs to start with "/", and includes parameters.
  // When the match type is REGEX/NOT_REGEX, only one value is allowed. 
  // Example: /test.html?id=1.', 'zh_CN':'URI。
  // 当匹配类型为等于/不等于/开头是/结尾是，URI必须以”/“开头，含参数。
  // 当匹配类型为正则/正则不匹配，则只允许只有一个值。
  // 示例：/test.html?id=1。'}
  Uri []*string `json:"uri,omitempty" xml:"uri,omitempty" require:"true" type:"Repeated"`
}

func (s CreateSharedRateLimitingRuleUriCondition) String() string {
  return tea.Prettify(s)
}

func (s CreateSharedRateLimitingRuleUriCondition) GoString() string {
  return s.String()
}

func (s *CreateSharedRateLimitingRuleUriCondition) SetMatchType(v string) *CreateSharedRateLimitingRuleUriCondition {
  s.MatchType = &v
  return s
}

func (s *CreateSharedRateLimitingRuleUriCondition) SetUri(v []*string) *CreateSharedRateLimitingRuleUriCondition {
  s.Uri = v
  return s
}

type CreateSharedRateLimitingRuleUaCondition struct {
  // {"en":"Match type.
  // EQUAL: Equals, user agent case sensitive
  // NOT_EQUAL: Does not equal, user agent case sensitive
  // CONTAIN: Contains, user agent case insensitive
  // NOT_CONTAIN: Does not Contains, user agent case insensitive
  // NONE:Empty or non-existent
  // REGEX: Regex match, user agent case insensitive
  // NOT_REGEX: Regular does not match, user agent case insensitive
  // START_WITH: Starts with, user agent case insensitive
  // END_WITH: Ends with, user agent case insensitive
  // WILDCARD: Wildcard matches, user agent case insensitive, ** represents zero or more arbitrary characters, ? represents any single character
  // NOT_WILDCARD: Wildcard does not match, user agent case insensitive, ** represents zero or more arbitrary characters, ? represents any single character", "zh_CN":"匹配类型。
  // EQUAL：等于，User-Agent大小写敏感
  // NOT_EQUAL：不等于，User-Agent大小写敏感
  // CONTAIN：包含，User-Agent大小写不敏感
  // NOT_CONTAIN：不包含，User-Agent大小写不敏感
  // NONE：为空或不存在
  // REGEX：匹配正则，User-Agent大小写不敏感
  // NOT_REGEX：正则不匹配，User-Agent大小写不敏感
  // START_WITH：开头是，User-Agent大小写不敏感
  // END_WITH：结尾是，User-Agent大小写不敏感
  // WILDCARD：通配符匹配，User-Agent大小写不敏感，*代表零个或多个任意字符，?代表任意单个字符
  // NOT_WILDCARD：通配符不匹配，User-Agent大小写不敏感，*代表零个或多个任意字符，?代表任意单个字符"}
  MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
  // {'en':'User agent.
  // When the match type is REGEX/NOT_REGEX, only one value is allowed. 
  // Example: go-Http-client/1.1.', 'zh_CN':'User-Agent。
  // 当匹配类型为正则/正则不匹配，则只允许只有一个值。
  // 示例：go-Http-client/1.1。'}
  Ua []*string `json:"ua,omitempty" xml:"ua,omitempty" require:"true" type:"Repeated"`
}

func (s CreateSharedRateLimitingRuleUaCondition) String() string {
  return tea.Prettify(s)
}

func (s CreateSharedRateLimitingRuleUaCondition) GoString() string {
  return s.String()
}

func (s *CreateSharedRateLimitingRuleUaCondition) SetMatchType(v string) *CreateSharedRateLimitingRuleUaCondition {
  s.MatchType = &v
  return s
}

func (s *CreateSharedRateLimitingRuleUaCondition) SetUa(v []*string) *CreateSharedRateLimitingRuleUaCondition {
  s.Ua = v
  return s
}

type CreateSharedRateLimitingRuleRefererCondition struct {
  // {"en":"Match type.
  // EQUAL: Equals, referer case sensitive
  // NOT_EQUAL: Does not equal, referer case sensitive
  // CONTAIN: Contains, referer case insensitive
  // NOT_CONTAIN: Does not Contains, referer case insensitive
  // NONE:Empty or non-existent
  // REGEX: Regex match, referer case insensitive
  // NOT_REGEX: Regular does not match, referer case insensitive
  // START_WITH: Starts with, referer case insensitive
  // END_WITH: Ends with, referer case insensitive
  // WILDCARD: Wildcard matches, referer case insensitive, ** represents zero or more arbitrary characters, ? represents any single characte
  // NOT_WILDCARD: Wildcard does not match, referer case insensitive, ** represents zero or more arbitrary characters, ? represents any single character", "zh_CN":"匹配类型。
  // EQUAL：等于，referer大小写敏感
  // NOT_EQUAL：不等于，referer大小写敏感
  // CONTAIN：包含，referer大小写不敏感
  // NOT_CONTAIN：不包含，referer大小写不敏感
  // NONE：为空或不存在
  // REGEX：匹配正则，referer大小写不敏感
  // NOT_REGEX：正则不匹配，referer大小写不敏感
  // START_WITH：开头是，referer大小写不敏感
  // END_WITH：结尾是，referer大小写不敏感
  // WILDCARD：通配符匹配，referer大小写不敏感，*代表零个或多个任意字符，?代表任意单个字符
  // NOT_WILDCARD：通配符不匹配，referer大小写不敏感，*代表零个或多个任意字符，?代表任意单个字符"}
  MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
  // {'en':'Referer.
  // When the match type is REGEX/NOT_REGEX, only one value is allowed. 
  // Example: http://test.com.', 'zh_CN':'Referer。
  // 当匹配类型为正则/正则不匹配，则只允许只有一个值。
  // 示例：http://test.com。'}
  Referer []*string `json:"referer,omitempty" xml:"referer,omitempty" require:"true" type:"Repeated"`
}

func (s CreateSharedRateLimitingRuleRefererCondition) String() string {
  return tea.Prettify(s)
}

func (s CreateSharedRateLimitingRuleRefererCondition) GoString() string {
  return s.String()
}

func (s *CreateSharedRateLimitingRuleRefererCondition) SetMatchType(v string) *CreateSharedRateLimitingRuleRefererCondition {
  s.MatchType = &v
  return s
}

func (s *CreateSharedRateLimitingRuleRefererCondition) SetReferer(v []*string) *CreateSharedRateLimitingRuleRefererCondition {
  s.Referer = v
  return s
}

type CreateSharedRateLimitingRuleStatusCodeCondition struct {
  // {'en':'Match type.
  // EQUAL:Equals
  // NOT_EQUAL:Does not equal', 'zh_CN':'匹配类型。
  // EQUAL：等于
  // NOT_EQUAL：不等于'}
  MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
  // {'en':'Response Code.', 'zh_CN':'状态码。'}
  StatusCode []*string `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true" type:"Repeated"`
}

func (s CreateSharedRateLimitingRuleStatusCodeCondition) String() string {
  return tea.Prettify(s)
}

func (s CreateSharedRateLimitingRuleStatusCodeCondition) GoString() string {
  return s.String()
}

func (s *CreateSharedRateLimitingRuleStatusCodeCondition) SetMatchType(v string) *CreateSharedRateLimitingRuleStatusCodeCondition {
  s.MatchType = &v
  return s
}

func (s *CreateSharedRateLimitingRuleStatusCodeCondition) SetStatusCode(v []*string) *CreateSharedRateLimitingRuleStatusCodeCondition {
  s.StatusCode = v
  return s
}

type CreateSharedRateLimitingRuleRequestMethodCondition struct {
  // {'en':'Match type.
  // EQUAL:Equals
  // NOT_EQUAL:Does not equal', 'zh_CN':'匹配类型。
  // EQUAL：等于
  // NOT_EQUAL：不等于'}
  MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
  // {'en':'Request method.
  // Supported values: GET/POST/DELETE/PUT/HEAD/OPTIONS/COPY.', 'zh_CN':'请求方法。
  // 支持的值：GET/POST/DELETE/PUT/HEAD/OPTIONS/COPY。'}
  RequestMethod []*string `json:"requestMethod,omitempty" xml:"requestMethod,omitempty" require:"true" type:"Repeated"`
}

func (s CreateSharedRateLimitingRuleRequestMethodCondition) String() string {
  return tea.Prettify(s)
}

func (s CreateSharedRateLimitingRuleRequestMethodCondition) GoString() string {
  return s.String()
}

func (s *CreateSharedRateLimitingRuleRequestMethodCondition) SetMatchType(v string) *CreateSharedRateLimitingRuleRequestMethodCondition {
  s.MatchType = &v
  return s
}

func (s *CreateSharedRateLimitingRuleRequestMethodCondition) SetRequestMethod(v []*string) *CreateSharedRateLimitingRuleRequestMethodCondition {
  s.RequestMethod = v
  return s
}

type CreateSharedRateLimitingRuleSchemeCondition struct {
  // {'en':'Match type.
  // EQUAL:Equals
  // NOT_EQUAL:Does not equal', 'zh_CN':'匹配类型。
  // EQUAL：等于
  // NOT_EQUAL：不等于'}
  MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
  // {'en':'HTTP/S.
  // Supported values: HTTP/HTTPS.', 'zh_CN':'应用层协议。
  // 支持的值：HTTP/HTTPS。'}
  Scheme []*string `json:"scheme,omitempty" xml:"scheme,omitempty" require:"true" type:"Repeated"`
}

func (s CreateSharedRateLimitingRuleSchemeCondition) String() string {
  return tea.Prettify(s)
}

func (s CreateSharedRateLimitingRuleSchemeCondition) GoString() string {
  return s.String()
}

func (s *CreateSharedRateLimitingRuleSchemeCondition) SetMatchType(v string) *CreateSharedRateLimitingRuleSchemeCondition {
  s.MatchType = &v
  return s
}

func (s *CreateSharedRateLimitingRuleSchemeCondition) SetScheme(v []*string) *CreateSharedRateLimitingRuleSchemeCondition {
  s.Scheme = v
  return s
}

type CreateSharedRateLimitingRuleShareRateLimitRuleCondition struct {
  // {'en':'IP/CIDR, match type cannot be repeated.', 'zh_CN':'IP/IP段，匹配类型不可重复。'}
  IpOrIpsConditions []*CreateSharedRateLimitingRuleIpOrIpsCondition `json:"ipOrIpsConditions,omitempty" xml:"ipOrIpsConditions,omitempty" type:"Repeated"`
  // {'en':'Path, match type cannot be repeated.
  // When the business scenario is API, this matching condition is not supported.', 'zh_CN':'路径，匹配类型不可重复。
  // 当业务场景为API业务时不支持此匹配条件。'}
  PathConditions []*CreateSharedRateLimitingRulePathCondition `json:"pathConditions,omitempty" xml:"pathConditions,omitempty" type:"Repeated"`
  // {'en':'URI, match type cannot be repeated.
  // 
  // When the business scenario is API, this matching condition is not supported.', 'zh_CN':'URI，匹配类型不可重复。
  // 当业务场景为API业务时不支持此匹配条件。'}
  UriConditions []*CreateSharedRateLimitingRuleUriCondition `json:"uriConditions,omitempty" xml:"uriConditions,omitempty" type:"Repeated"`
  // {'en':'User Agent, match type cannot be repeated.', 'zh_CN':'User-Agent，匹配类型不可重复。'}
  UaConditions []*CreateSharedRateLimitingRuleUaCondition `json:"uaConditions,omitempty" xml:"uaConditions,omitempty" type:"Repeated"`
  // {'en':'Request Method.
  // When the business scenario is API,this matching condition is not supported.', 'zh_CN':'请求方法，匹配类型不可重复。
  // 当业务场景为API业务时不支持此匹配条件。'}
  MethodConditions []*CreateSharedRateLimitingRuleRequestMethodCondition `json:"methodConditions,omitempty" xml:"methodConditions,omitempty" type:"Repeated"`
  // {'en':'Referer, match type cannot be repeated.', 'zh_CN':'Referer，匹配类型不可重复。'}
  RefererConditions []*CreateSharedRateLimitingRuleRefererCondition `json:"refererConditions,omitempty" xml:"refererConditions,omitempty" type:"Repeated"`
  // {'en':'Request Header, match type can be repeated.', 'zh_CN':'请求头，匹配类型可重复。'}
  HeaderConditions []*CreateSharedRateLimitingRuleHeaderCondition `json:"headerConditions,omitempty" xml:"headerConditions,omitempty" type:"Repeated"`
  // {'en':'Geo,match type cannot be repeated.', 'zh_CN':'区域，匹配类型不可重复。'}
  AreaConditions []*CreateSharedRateLimitingRuleAreaCondition `json:"areaConditions,omitempty" xml:"areaConditions,omitempty" type:"Repeated"`
  // {'en':'Response Code, match type cannot be repeated.', 'zh_CN':'状态码，匹配类型不可重复。'}
  StatusCodeConditions []*CreateSharedRateLimitingRuleStatusCodeCondition `json:"statusCodeConditions,omitempty" xml:"statusCodeConditions,omitempty" type:"Repeated"`
  // {'en':'HTTP/S, match type cannot be repeated.', 'zh_CN':'应用层协议，匹配类型不可重复。'}
  SchemeConditions []*CreateSharedRateLimitingRuleSchemeCondition `json:"schemeConditions,omitempty" xml:"schemeConditions,omitempty" type:"Repeated"`
}

func (s CreateSharedRateLimitingRuleShareRateLimitRuleCondition) String() string {
  return tea.Prettify(s)
}

func (s CreateSharedRateLimitingRuleShareRateLimitRuleCondition) GoString() string {
  return s.String()
}

func (s *CreateSharedRateLimitingRuleShareRateLimitRuleCondition) SetIpOrIpsConditions(v []*CreateSharedRateLimitingRuleIpOrIpsCondition) *CreateSharedRateLimitingRuleShareRateLimitRuleCondition {
  s.IpOrIpsConditions = v
  return s
}

func (s *CreateSharedRateLimitingRuleShareRateLimitRuleCondition) SetPathConditions(v []*CreateSharedRateLimitingRulePathCondition) *CreateSharedRateLimitingRuleShareRateLimitRuleCondition {
  s.PathConditions = v
  return s
}

func (s *CreateSharedRateLimitingRuleShareRateLimitRuleCondition) SetUriConditions(v []*CreateSharedRateLimitingRuleUriCondition) *CreateSharedRateLimitingRuleShareRateLimitRuleCondition {
  s.UriConditions = v
  return s
}

func (s *CreateSharedRateLimitingRuleShareRateLimitRuleCondition) SetUaConditions(v []*CreateSharedRateLimitingRuleUaCondition) *CreateSharedRateLimitingRuleShareRateLimitRuleCondition {
  s.UaConditions = v
  return s
}

func (s *CreateSharedRateLimitingRuleShareRateLimitRuleCondition) SetMethodConditions(v []*CreateSharedRateLimitingRuleRequestMethodCondition) *CreateSharedRateLimitingRuleShareRateLimitRuleCondition {
  s.MethodConditions = v
  return s
}

func (s *CreateSharedRateLimitingRuleShareRateLimitRuleCondition) SetRefererConditions(v []*CreateSharedRateLimitingRuleRefererCondition) *CreateSharedRateLimitingRuleShareRateLimitRuleCondition {
  s.RefererConditions = v
  return s
}

func (s *CreateSharedRateLimitingRuleShareRateLimitRuleCondition) SetHeaderConditions(v []*CreateSharedRateLimitingRuleHeaderCondition) *CreateSharedRateLimitingRuleShareRateLimitRuleCondition {
  s.HeaderConditions = v
  return s
}

func (s *CreateSharedRateLimitingRuleShareRateLimitRuleCondition) SetAreaConditions(v []*CreateSharedRateLimitingRuleAreaCondition) *CreateSharedRateLimitingRuleShareRateLimitRuleCondition {
  s.AreaConditions = v
  return s
}

func (s *CreateSharedRateLimitingRuleShareRateLimitRuleCondition) SetStatusCodeConditions(v []*CreateSharedRateLimitingRuleStatusCodeCondition) *CreateSharedRateLimitingRuleShareRateLimitRuleCondition {
  s.StatusCodeConditions = v
  return s
}

func (s *CreateSharedRateLimitingRuleShareRateLimitRuleCondition) SetSchemeConditions(v []*CreateSharedRateLimitingRuleSchemeCondition) *CreateSharedRateLimitingRuleShareRateLimitRuleCondition {
  s.SchemeConditions = v
  return s
}

type CreateSharedRateLimitingRuleRateLimitEffective struct {
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

func (s CreateSharedRateLimitingRuleRateLimitEffective) String() string {
  return tea.Prettify(s)
}

func (s CreateSharedRateLimitingRuleRateLimitEffective) GoString() string {
  return s.String()
}

func (s *CreateSharedRateLimitingRuleRateLimitEffective) SetEffective(v []*string) *CreateSharedRateLimitingRuleRateLimitEffective {
  s.Effective = v
  return s
}

func (s *CreateSharedRateLimitingRuleRateLimitEffective) SetStart(v string) *CreateSharedRateLimitingRuleRateLimitEffective {
  s.Start = &v
  return s
}

func (s *CreateSharedRateLimitingRuleRateLimitEffective) SetEnd(v string) *CreateSharedRateLimitingRuleRateLimitEffective {
  s.End = &v
  return s
}

func (s *CreateSharedRateLimitingRuleRateLimitEffective) SetTimezone(v string) *CreateSharedRateLimitingRuleRateLimitEffective {
  s.Timezone = &v
  return s
}

type CreateSharedRateLimitingRulePathCondition struct {
  // {"en":"Match type.
  // EQUAL: Equals, path case sensitive
  // NOT_EQUAL: Does not equal, path case sensitive
  // CONTAIN: Contains, path case insensitive
  // NOT_CONTAIN: Does not Contains, path case insensitive
  // REGEX: Regex match, path case insensitive
  // NOT_REGEX: Regular does not match, path case sensitive
  // START_WITH: Starts with, path case sensitive
  // END_WITH: Ends with, path case sensitive
  // WILDCARD: Wildcard matches, path case sensitive, ** represents zero or more arbitrary characters, ? represents any single character.
  // NOT_WILDCARD: Wildcard does not match, path case sensitive, ** represents zero or more arbitrary characters, ? represents any single character ", "zh_CN":"匹配类型。
  // EQUAL：等于，路径大小写敏感
  // NOT_EQUAL：不等于，路径大小写敏感
  // CONTAIN：包含，路径大小写不敏感
  // NOT_CONTAIN：不包含，路径大小写不敏感
  // REGEX：匹配正则，路径大小写不敏感
  // NOT_REGEX：正则不匹配，路径大小写不敏感
  // START_WITH：开头是，路径大小写不敏感
  // END_WITH：结尾是，路径大小写不敏感
  // WILDCARD：通配符匹配，路径大小写不敏感，*代表零个或多个任意字符，?代表任意单个字符
  // NOT_WILDCARD：通配符不匹配，路径大小写不敏感，*代表零个或多个任意字符，?代表任意单个字符"}
  MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
  // {'en':'Path.
  // When match type is EQUAL/NOT_EQUAL/START_WITH/END_WITH, path needs to start with "/", and no parameters.
  // When the match type is REGEX/NOT_REGEX, only one value is allowed. 
  // Example: /test.html.', 'zh_CN':'路径。
  // 当匹配类型为等于/不等于/开头是/结尾是，路径必须以“/”开头，不含参数。
  // 当匹配类型为正则/正则不匹配，则只允许只有一个值。
  // 示例：/test.html。'}
  CreateSharedRateLimitingRulePaths []*string `json:"paths,omitempty" xml:"paths,omitempty" require:"true" type:"Repeated"`
}

func (s CreateSharedRateLimitingRulePathCondition) String() string {
  return tea.Prettify(s)
}

func (s CreateSharedRateLimitingRulePathCondition) GoString() string {
  return s.String()
}

func (s *CreateSharedRateLimitingRulePathCondition) SetMatchType(v string) *CreateSharedRateLimitingRulePathCondition {
  s.MatchType = &v
  return s
}

func (s *CreateSharedRateLimitingRulePathCondition) SetCreateSharedRateLimitingRulePaths(v []*string) *CreateSharedRateLimitingRulePathCondition {
  s.CreateSharedRateLimitingRulePaths = v
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

type CreateSharedRateLimitingRuleRequestHeader struct {
  // {"zh_CN":"安全服务类型。有使用多个不同的安全服务时，需要填写具体的服务类型。","en":"Security service type. Please enter a specific service type, if you purchase multiple security services.","dictionary":"belong=WAAP-MS-Ext|dict=waap_serviceType"}
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

type CreateSharedRateLimitingRuleResponseHeader struct {
}

func (s CreateSharedRateLimitingRuleResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s CreateSharedRateLimitingRuleResponseHeader) GoString() string {
  return s.String()
}




type DisassociateShareCustomizeRuleRequest struct {
  // {"zh_CN":"共享规则ID。","en":"Share rule ID."}
  ShareId *string `json:"shareId,omitempty" xml:"shareId,omitempty" require:"true"`
  // {"zh_CN":"待取消关联域名列表。","en":"List of hostname to be disassociated."}
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

type DisassociateShareCustomizeRuleResponse struct {
  // {'en':'Please refer to the error code for exceptions.', 'zh_CN':'请参照错误码。','dictionary':'belong=WAAP-MS-Ext|dict=waap_retCodeEnum'}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {'en':'Description.', 'zh_CN':'描述信息。'}
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

type DisassociateShareCustomizeRuleRequestHeader struct {
  // {"zh_CN":"安全服务类型。有使用多个不同的安全服务时，需要填写具体的服务类型。","en":"Security service type. Please enter a specific service type, if you purchase multiple security services.","dictionary":"belong=WAAP-MS-Ext|dict=waap_serviceType"}
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

type ListCustomActionsResponse struct {
  // {'en':'Please refer to the error code for exceptions.', 'zh_CN':'请参照错误码。','dictionary':'belong=WAAP-MS-Ext|dict=waap_retCodeEnum'}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {'en':'Description.', 'zh_CN':'描述信息。'}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {'en':'Data.', 'zh_CN':'出参数据。'}
  Data []*ListCustomActionsCommonShareCustomizeActionVO `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
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

func (s *ListCustomActionsResponse) SetData(v []*ListCustomActionsCommonShareCustomizeActionVO) *ListCustomActionsResponse {
  s.Data = v
  return s
}

type ListCustomActionsCommonShareCustomizeActionVO struct {
  // {'en':'Custom response ID.', 'zh_CN':'自定义响应ID。'}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {'en':'Main account.', 'zh_CN':'主账号。'}
  MainAccount *string `json:"mainAccount,omitempty" xml:"mainAccount,omitempty" require:"true"`
  // {'en':'Creator.', 'zh_CN':'创建者。'}
  Creator *string `json:"creator,omitempty" xml:"creator,omitempty" require:"true"`
  // {'en':'Action name,maximum 50 characters.
  // Does not support special characters and spaces.', 'zh_CN':'动作名称，最多50个字符。
  // 不支持特殊字符和空格。'}
  ActionName *string `json:"actionName,omitempty" xml:"actionName,omitempty" require:"true"`
  // {'en':'Description,maximum 200 characters.', 'zh_CN':'描述，最多200个字符。'}
  Description *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
  // {'en':'Status code.
  // Valid value range:200,204,206,400,401,403,404,500,501,503.', 'zh_CN':'状态码。
  // 有效值范围：200，204，206，400，401，403，404，500，501，503。'}
  StatusCode *string `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
  // {'en':'Response Content-Type type,multiple separated by ; sign.', 'zh_CN':'响应Content-Type类型，多个以 ; 隔开。'}
  ContentType *string `json:"contentType,omitempty" xml:"contentType,omitempty" require:"true"`
  // {'en':'Response content definition, the size of the response content cannot exceed 16KB, if you need to insert static resources such as pictures, it is recommended to use links to import. The following interception information parameters are supported: 
  //  {url}} : Displays the URL information when intercepting 
  //  {client_ip}} : Displays the intercepted user IP information 
  //  {time}} : Displays the intercepted time point 
  //  {event_id} : Displays the ID information of this interception event', 'zh_CN':'自定义响应内容定义，响应内容大小不能超过16KB，如需插入图片等静态资源，建议采用链接引入。
  // 支持引用下列拦截信息参数：
  // {url} ：显示拦截时的URL信息
  // {client_ip} ：显示被拦截的用户IP信息
  // {time} ：显示被拦截的时间点
  // {event_id} ：显示本次拦截事件的ID信息'}
  CusBlockCnt *string `json:"cusBlockCnt,omitempty" xml:"cusBlockCnt,omitempty" require:"true"`
  // {"zh_CN":"关联域名列表。","en":"List of associated hostnames."}
  DomainList *string `json:"domainList,omitempty" xml:"domainList,omitempty" require:"true"`
}

func (s ListCustomActionsCommonShareCustomizeActionVO) String() string {
  return tea.Prettify(s)
}

func (s ListCustomActionsCommonShareCustomizeActionVO) GoString() string {
  return s.String()
}

func (s *ListCustomActionsCommonShareCustomizeActionVO) SetId(v string) *ListCustomActionsCommonShareCustomizeActionVO {
  s.Id = &v
  return s
}

func (s *ListCustomActionsCommonShareCustomizeActionVO) SetMainAccount(v string) *ListCustomActionsCommonShareCustomizeActionVO {
  s.MainAccount = &v
  return s
}

func (s *ListCustomActionsCommonShareCustomizeActionVO) SetCreator(v string) *ListCustomActionsCommonShareCustomizeActionVO {
  s.Creator = &v
  return s
}

func (s *ListCustomActionsCommonShareCustomizeActionVO) SetActionName(v string) *ListCustomActionsCommonShareCustomizeActionVO {
  s.ActionName = &v
  return s
}

func (s *ListCustomActionsCommonShareCustomizeActionVO) SetDescription(v string) *ListCustomActionsCommonShareCustomizeActionVO {
  s.Description = &v
  return s
}

func (s *ListCustomActionsCommonShareCustomizeActionVO) SetStatusCode(v string) *ListCustomActionsCommonShareCustomizeActionVO {
  s.StatusCode = &v
  return s
}

func (s *ListCustomActionsCommonShareCustomizeActionVO) SetContentType(v string) *ListCustomActionsCommonShareCustomizeActionVO {
  s.ContentType = &v
  return s
}

func (s *ListCustomActionsCommonShareCustomizeActionVO) SetCusBlockCnt(v string) *ListCustomActionsCommonShareCustomizeActionVO {
  s.CusBlockCnt = &v
  return s
}

func (s *ListCustomActionsCommonShareCustomizeActionVO) SetDomainList(v string) *ListCustomActionsCommonShareCustomizeActionVO {
  s.DomainList = &v
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

type ListCustomActionsRequestHeader struct {
  // {"zh_CN":"安全服务类型。有使用多个不同的安全服务时，需要填写具体的服务类型。","en":"Security service type. Please enter a specific service type, if you purchase multiple security services.","dictionary":"belong=WAAP-MS-Ext|dict=waap_serviceType"}
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

type ListCustomActionsResponseHeader struct {
}

func (s ListCustomActionsResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ListCustomActionsResponseHeader) GoString() string {
  return s.String()
}




type DisassociateDmsShareServiceFeatureRequest struct {
  // {"zh_CN":"共享配置ID。","en":"Share configuration ID."}
  ShareId *string `json:"shareId,omitempty" xml:"shareId,omitempty" require:"true"`
  // {"zh_CN":"待解除关联域名列表。","en":"List of hostname to be disassociated."}
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

type DisassociateDmsShareServiceFeatureResponse struct {
  // {'en':'Please refer to the error code for exceptions.', 'zh_CN':'请参照错误码。','dictionary':'belong=WAAP-MS-Ext|dict=waap_retCodeEnum'}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {'en':'Description.', 'zh_CN':'描述信息。'}
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

type DisassociateDmsShareServiceFeatureRequestHeader struct {
  // {"zh_CN":"安全服务类型。有使用多个不同的安全服务时，需要填写具体的服务类型。","en":"Security service type. Please enter a specific service type, if you purchase multiple security services.","dictionary":"belong=WAAP-MS-Ext|dict=waap_serviceType"}
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

type DisassociateDmsShareServiceFeatureResponseHeader struct {
}

func (s DisassociateDmsShareServiceFeatureResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s DisassociateDmsShareServiceFeatureResponseHeader) GoString() string {
  return s.String()
}




type AssociateShareRateLimitRequest struct {
  // {"zh_CN":"共享规则ID。","en":"Share rule ID."}
  ShareId *string `json:"shareId,omitempty" xml:"shareId,omitempty" require:"true"`
  // {"zh_CN":"待关联域名列表。","en":"List of hostname to be associated."}
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

type AssociateShareRateLimitResponse struct {
  // {'en':'Please refer to the error code for exceptions.', 'zh_CN':'请参照错误码。','dictionary':'belong=WAAP-MS-Ext|dict=waap_retCodeEnum'}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {'en':'Description.', 'zh_CN':'描述信息。'}
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

type AssociateShareRateLimitRequestHeader struct {
  // {"zh_CN":"安全服务类型。有使用多个不同的安全服务时，需要填写具体的服务类型。","en":"Security service type. Please enter a specific service type, if you purchase multiple security services.","dictionary":"belong=WAAP-MS-Ext|dict=waap_serviceType"}
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

type AssociateShareRateLimitResponseHeader struct {
}

func (s AssociateShareRateLimitResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s AssociateShareRateLimitResponseHeader) GoString() string {
  return s.String()
}




type AssociateDmsShareServiceFeatureRequest struct {
  // {"zh_CN":"共享配置ID。","en":"Share configuration ID."}
  ShareId *string `json:"shareId,omitempty" xml:"shareId,omitempty" require:"true"`
  // {"zh_CN":"待关联域名列表。","en":"List of hostname to be associated."}
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

type AssociateDmsShareServiceFeatureResponse struct {
  // {'en':'Please refer to the error code for exceptions.', 'zh_CN':'请参照错误码。','dictionary':'belong=WAAP-MS-Ext|dict=waap_retCodeEnum'}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {'en':'Description.', 'zh_CN':'描述信息。'}
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

type AssociateDmsShareServiceFeatureRequestHeader struct {
  // {"zh_CN":"安全服务类型。有使用多个不同的安全服务时，需要填写具体的服务类型。","en":"Security service type. Please enter a specific service type, if you purchase multiple security services.","dictionary":"belong=WAAP-MS-Ext|dict=waap_serviceType"}
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

type AssociateDmsShareServiceFeatureResponseHeader struct {
}

func (s AssociateDmsShareServiceFeatureResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s AssociateDmsShareServiceFeatureResponseHeader) GoString() string {
  return s.String()
}




type AssociateSharedWhitelistRuleRequest struct {
  // {"zh_CN":"共享规则ID。","en":"Share rule ID."}
  ShareId *string `json:"shareId,omitempty" xml:"shareId,omitempty" require:"true"`
  // {"zh_CN":"待关联域名列表。","en":"List of hostname to be associated."}
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

type AssociateSharedWhitelistRuleResponse struct {
  // {'en':'Please refer to the error code for exceptions.', 'zh_CN':'请参照错误码。','dictionary':'belong=WAAP-MS-Ext|dict=waap_retCodeEnum'}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {'en':'Description.', 'zh_CN':'描述信息。'}
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

type AssociateSharedWhitelistRuleRequestHeader struct {
  // {"zh_CN":"安全服务类型。有使用多个不同的安全服务时，需要填写具体的服务类型。","en":"Security service type. Please enter a specific service type, if you purchase multiple security services.","dictionary":"belong=WAAP-MS-Ext|dict=waap_serviceType"}
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

type AssociateSharedWhitelistRuleResponseHeader struct {
}

func (s AssociateSharedWhitelistRuleResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s AssociateSharedWhitelistRuleResponseHeader) GoString() string {
  return s.String()
}




type CreateSharedWAFRuleExceptionRequest struct {
  // {'en':'Exception name,maximum 50 character.
  // Does not support special characters and spaces.', 'zh_CN':'例外名称，最多50个字符。
  // 不支持特殊字符和空格。'}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {'en':'Exception description,maximum 200 characters.', 'zh_CN':'例外描述，最多200个字符。'}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty"`
  // {'en':'Matching conditions.
  // ip: IP
  // path: Path
  // uri: URI
  // urlParamName: URI Parameter Name
  // urlParamValue: URI Parameter Value
  // userAgent: User Agent
  // httpHeaderName: Request Header Name
  // httpHeaderValue: Request Header Value
  // cookie: Cookie
  // body: Body
  // bodyParamName: Body Parameter Name
  // bodyParamValue: Body Parameter Value', 'zh_CN':'匹配条件。
  // ip：IP
  // path：路径
  // uri：URI
  // urlParamName：URI参数名
  // urlParamValue：URI参数值
  // userAgent：User Agent
  // httpHeaderName：请求头部名称
  // httpHeaderValue：请求头部值
  // cookie：Cookie
  // body：Body
  // bodyParamName：Body参数名
  // bodyParamValue：Body参数值'}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
  // {'en':'Match type,IP can only be EQUAL.
  // EQUAL: Equal
  // CONTAIN: Contains
  // REGEX: Regular match', 'zh_CN':'匹配类型，IP只能是等于。
  // EQUAL：等于
  // CONTAIN：包含
  // REGEX：正则匹配'}
  MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
  // {'en':'Rule exceptions.
  // When matchType=EQUAL, case-sensitive, path and uri must start with "/", and body can only pass one value;
  // When matchType=REGEX, only one value can be passed', 'zh_CN':'规则例外内容。
  // matchType=EQUAL时，大小写敏感，path和uri必须以"/"开头，body只能传一个值；
  // matchType=REGEX时，只能传一个值。'}
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

type CreateSharedWAFRuleExceptionResponse struct {
  // {'en':'Please refer to the error code for exceptions.', 'zh_CN':'请参照错误码。','dictionary':'belong=WAAP-MS-Ext|dict=waap_retCodeEnum'}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {'en':'Description.', 'zh_CN':'描述信息。'}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {'en':'Rule exception ID.', 'zh_CN':'规则例外ID。'}
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

type CreateSharedWAFRuleExceptionRequestHeader struct {
  // {"zh_CN":"安全服务类型。有使用多个不同的安全服务时，需要填写具体的服务类型。","en":"Security service type. Please enter a specific service type, if you purchase multiple security services.","dictionary":"belong=WAAP-MS-Ext|dict=waap_serviceType"}
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
  // {"en":"Action.
  // NO_USE: Not Used
  // LOG: Log
  // DELAY: Delay
  // BLOCK: Deny
  // RESET: Reset Connection","zh_CN":"处理动作。
  // NO_USE：不使用
  // LOG：监控
  // DELAY：延迟响应
  // BLOCK：拦截
  // RESET：断开连接"}
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
  // {"en":"JA3 Fingerprint, match type cannot be repeated.","zh_CN":"JA3指纹，匹配类型不可重复。"}
  Ja3Conditions []*ListSharedCustomRulesResponseDataConditionJa3Conditions `json:"ja3Conditions,omitempty" xml:"ja3Conditions,omitempty" require:"true" type:"Repeated"`
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

func (s *ListSharedCustomRulesResponseDataCondition) SetJa3Conditions(v []*ListSharedCustomRulesResponseDataConditionJa3Conditions) *ListSharedCustomRulesResponseDataCondition {
  s.Ja3Conditions = v
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

func (s *ListSharedCustomRulesResponseDataCondition) SetJa4Conditions(v []*ListSharedCustomRulesResponseDataConditionJa4Conditions) *ListSharedCustomRulesResponseDataCondition {
  s.Ja4Conditions = v
  return s
}

type ListSharedCustomRulesResponseDataConditionMethodConditions struct     {
  // {"en":"Match type.
  // EQUAL:Equal
  // NOT_EQUAL:Does not equal","zh_CN":"匹配类型。
  // EQUAL：等于
  // NOT_EQUAL：不等于"}
  MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
  // {"en":"Request method.
  // Supported values: GET/POST/DELETE/PUT/HEAD/OPTIONS/COPY.","zh_CN":"请求方法。
  // 支持的值：GET/POST/DELETE/PUT/HEAD/OPTIONS/COPY"}
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

type ListSharedCustomRulesResponseDataConditionJa3Conditions struct     {
  // {"en":"Match type.
  // EQUAL: Equals
  // NOT_EQUAL: Does not equal","zh_CN":"匹配类型。
  // EQUAL：等于 
  // NOT_EQUAL：不等于"}
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

type ListSharedCustomRulesResponseDataConditionAreaConditions struct     {
  // {"en":"Match type.
  // EQUAL:Equal
  // NOT_EQUAL:Does not equal","zh_CN":"匹配类型。
  // EQUAL：等于
  // NOT_EQUAL：不等于"}
  MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
  // {"en":"Geo.","zh_CN":"区域。"}
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
  // {"en":"Match type.
  // EQUAL:Equal
  // NOT_EQUAL:Does not equal","zh_CN":"匹配类型。
  // EQUAL：等于
  // NOT_EQUAL：不等于"}
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
  // {"en":"Match type.
  // EQUAL: equal to
  // NOT_EQUAL: not equal to
  // CONTAIN: contains
  // NOT_CONTAIN: does not contain
  // REGEX: regular
  // NOT_REGEX: regular does not match
  // START_WITH: starts with
  // END_WITH: ends with
  // WILDCARD: wildcard matches
  // NOT_WILDCARD: wildcard does not match","zh_CN":"匹配类型。
  // EQUAL：等于
  // NOT_EQUAL：不等于
  // CONTAIN：包含
  // NOT_CONTAIN：不包含
  // REGEX：正则
  // NOT_REGEX：正则不匹配
  // START_WITH：开头是
  // END_WITH：结尾是
  // WILDCARD：通配符匹配
  // NOT_WILDCARD：通配符不匹配"}
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
  // {"en":"Match type.
  // EQUAL: equal to
  // NOT_EQUAL: not equal to
  // CONTAIN: contains
  // NOT_CONTAIN: does not contain
  // REGEX: regular
  // NOT_REGEX: regular does not match
  // START_WITH: starts with
  // END_WITH: ends with
  // WILDCARD: wildcard matches
  // NOT_WILDCARD: wildcard does not match","zh_CN":"匹配类型。
  // EQUAL：等于
  // NOT_EQUAL：不等于
  // CONTAIN：包含
  // NOT_CONTAIN：不包含
  // REGEX：正则
  // NOT_REGEX：正则不匹配
  // START_WITH：开头是
  // END_WITH：结尾是
  // WILDCARD：通配符匹配
  // NOT_WILDCARD：通配符不匹配"}
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
  // {"en":"Match type.
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
  // NOT_WILDCARD: wildcard does not match","zh_CN":"匹配类型。
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
  // NOT_WILDCARD：通配符不匹配"}
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
  // {"en":"Match type.
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
  // NOT_WILDCARD: wildcard does not match","zh_CN":"匹配类型。
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
  // NOT_WILDCARD：通配符不匹配"}
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
  // {"en":"Match type.
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
  // NOT_WILDCARD: wildcard does not match","zh_CN":"匹配类型。
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
  // NOT_WILDCARD：通配符不匹配"}
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

type ListSharedCustomRulesResponseDataConditionJa4Conditions struct     {
  // {"en":"JA4 Fingerprint List.","zh_CN":"JA4指纹列表。"}
  Ja4List []*string `json:"ja4List,omitempty" xml:"ja4List,omitempty" require:"true" type:"Repeated"`
  // {"en":"Match type. 
  // EQUAL: Equals
  // NOT_EQUAL: Does not equal
  // CONTAIN: Contains
  // NOT_CONTAIN: Does not Contains
  // START_WITH: Starts with
  // END_WITH: Ends with
  // WILDCARD: Wildcard matches, ** represents zero or more arbitrary characters, ? represents any single character
  // NOT_WILDCARD: Wildcard does not match, ** represents zero or more arbitrary characters, ? represents any single character","zh_CN":"匹配类型。
  // EQUAL：等于
  // NOT_EQUAL：不等于
  // CONTAIN：包含
  // NOT_CONTAIN：不包含
  // START_WITH：开头是
  // END_WITH：结尾是
  // WILDCARD：通配符匹配，*代表零个或多个任意字符，?代表任意单个字符
  // NOT_WILDCARD：通配符不匹配，*代表零个或多个任意字符，?代表任意单个字符"}
  MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
}

func (s ListSharedCustomRulesResponseDataConditionJa4Conditions) String() string {
  return tea.Prettify(s)
}

func (s ListSharedCustomRulesResponseDataConditionJa4Conditions) GoString() string {
  return s.String()
}

func (s *ListSharedCustomRulesResponseDataConditionJa4Conditions) SetJa4List(v []*string) *ListSharedCustomRulesResponseDataConditionJa4Conditions {
  s.Ja4List = v
  return s
}

func (s *ListSharedCustomRulesResponseDataConditionJa4Conditions) SetMatchType(v string) *ListSharedCustomRulesResponseDataConditionJa4Conditions {
  s.MatchType = &v
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
  // {"zh_CN":"共享规则ID。","en":"Share rule ID."}
  ShareId *string `json:"shareId,omitempty" xml:"shareId,omitempty" require:"true"`
  // {"zh_CN":"待取消关联域名列表。","en":"List of hostname to be disassociated."}
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

type DisassociateShareRateLimitResponse struct {
  // {'en':'Please refer to the error code for exceptions.', 'zh_CN':'请参照错误码。','dictionary':'belong=WAAP-MS-Ext|dict=waap_retCodeEnum'}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {'en':'Description.', 'zh_CN':'描述信息。'}
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

type DisassociateShareRateLimitRequestHeader struct {
  // {"zh_CN":"安全服务类型。有使用多个不同的安全服务时，需要填写具体的服务类型。","en":"Security service type. Please enter a specific service type, if you purchase multiple security services.","dictionary":"belong=WAAP-MS-Ext|dict=waap_serviceType"}
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

type DisassociateShareRateLimitResponseHeader struct {
}

func (s DisassociateShareRateLimitResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s DisassociateShareRateLimitResponseHeader) GoString() string {
  return s.String()
}




type UpdateSharedWAFRuleExceptionRequest struct {
  // {'en':'Rule exception ID.', 'zh_CN':'规则例外ID。'}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {'en':'Exception name,maximum 50 character.
  // Does not support special characters and spaces.', 'zh_CN':'例外名称，最多50个字符。
  // 不支持特殊字符和空格。'}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {'en':'Exception description,maximum 200 characters.', 'zh_CN':'例外描述，最多200个字符。'}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty"`
  // {'en':'Matching conditions.
  // ip: IP
  // path: Path
  // uri: URI
  // urlParamName: URI Parameter Name
  // urlParamValue: URI Parameter Value
  // userAgent: User Agent
  // httpHeaderName: Request Header Name
  // httpHeaderValue: Request Header Value
  // cookie: Cookie
  // body: Body
  // bodyParamName: Body Parameter Name
  // bodyParamValue: Body Parameter Value', 'zh_CN':'匹配条件。
  // ip：IP
  // path：路径
  // uri：URI
  // urlParamName：URI参数名
  // urlParamValue：URI参数值
  // userAgent：User Agent
  // httpHeaderName：请求头部名称
  // httpHeaderValue：请求头部值
  // cookie：Cookie
  // body：Body
  // bodyParamName：Body参数名
  // bodyParamValue：Body参数值'}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
  // {'en':'Match type,IP can only be EQUAL.
  // EQUAL: Equal
  // CONTAIN: Contains
  // REGEX: Regular match', 'zh_CN':'匹配类型，IP只能是等于。
  // EQUAL：等于
  // CONTAIN：包含
  // REGEX：正则匹配'}
  MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
  // {'en':'Rule exceptions.
  // When matchType=EQUAL, case-sensitive, path and uri must start with "/", and body can only pass one value;
  // When matchType=REGEX, only one value can be passed', 'zh_CN':'规则例外内容。
  // matchType=EQUAL时，大小写敏感，path和uri必须以"/"开头，body只能传一个值；
  // matchType=REGEX时，只能传一个值。'}
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

type UpdateSharedWAFRuleExceptionResponse struct {
  // {'en':'Please refer to the error code for exceptions.', 'zh_CN':'请参照错误码。','dictionary':'belong=WAAP-MS-Ext|dict=waap_retCodeEnum'}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {'en':'Description.', 'zh_CN':'描述信息。'}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {'en':'Modification success indicator.', 'zh_CN':'修改成功标识。'}
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

type UpdateSharedWAFRuleExceptionRequestHeader struct {
  // {"zh_CN":"安全服务类型。有使用多个不同的安全服务时，需要填写具体的服务类型。","en":"Security service type. Please enter a specific service type, if you purchase multiple security services.","dictionary":"belong=WAAP-MS-Ext|dict=waap_serviceType"}
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

type UpdateSharedWAFRuleExceptionResponseHeader struct {
}

func (s UpdateSharedWAFRuleExceptionResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s UpdateSharedWAFRuleExceptionResponseHeader) GoString() string {
  return s.String()
}




type UpdateCustomActionRequest struct {
  // {'en':'Custom response ID.', 'zh_CN':'自定义响应ID。'}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {'en':'Action name,maximum 50 characters.
  // Does not support special characters and spaces.', 'zh_CN':'动作名称，最多50个字符。
  // 不支持特殊字符和空格。'}
  ActionName *string `json:"actionName,omitempty" xml:"actionName,omitempty"`
  // {'en':'Description,maximum 200 characters.', 'zh_CN':'描述，最多200个字符。'}
  Description *string `json:"description,omitempty" xml:"description,omitempty"`
  // {'en':'Status code.
  // Valid value range:200,204,206,400,401,403,404,500,501,503.', 'zh_CN':'状态码。
  // 有效值范围：200，204，206，400，401，403，404，500，501，503。'}
  StatusCode *string `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  // {'en':'Response Content-Type type,multiple separated by ; sign.', 'zh_CN':'响应Content-Type类型，多个以 ; 隔开。'}
  ContentType *string `json:"contentType,omitempty" xml:"contentType,omitempty"`
  // {'en':'Response content definition, the size of the response content cannot exceed 16KB, if you need to insert static resources such as pictures, it is recommended to use links to import. 
  // The following interception information parameters are supported: 
  //  {url} : Displays the URL information when intercepting 
  //  {client_ip} : Displays the intercepted user IP information 
  //  {time} : Displays the intercepted time point 
  //  {event_id} : Displays the ID information of this interception event', 'zh_CN':'自定义响应内容定义，响应内容大小不能超过16KB，如需插入图片等静态资源，建议采用链接引入。
  // 支持引用下列拦截信息参数：
  // {url}  ：显示拦截时的URL信息
  // {client_ip}  ：显示被拦截的用户IP信息
  // {time}  ：显示被拦截的时间点
  // {event_id}  ：显示本次拦截事件的ID信息'}
  CusBlockCnt *string `json:"cusBlockCnt,omitempty" xml:"cusBlockCnt,omitempty"`
}

func (s UpdateCustomActionRequest) String() string {
  return tea.Prettify(s)
}

func (s UpdateCustomActionRequest) GoString() string {
  return s.String()
}

func (s *UpdateCustomActionRequest) SetId(v string) *UpdateCustomActionRequest {
  s.Id = &v
  return s
}

func (s *UpdateCustomActionRequest) SetActionName(v string) *UpdateCustomActionRequest {
  s.ActionName = &v
  return s
}

func (s *UpdateCustomActionRequest) SetDescription(v string) *UpdateCustomActionRequest {
  s.Description = &v
  return s
}

func (s *UpdateCustomActionRequest) SetStatusCode(v string) *UpdateCustomActionRequest {
  s.StatusCode = &v
  return s
}

func (s *UpdateCustomActionRequest) SetContentType(v string) *UpdateCustomActionRequest {
  s.ContentType = &v
  return s
}

func (s *UpdateCustomActionRequest) SetCusBlockCnt(v string) *UpdateCustomActionRequest {
  s.CusBlockCnt = &v
  return s
}

type UpdateCustomActionResponse struct {
  // {'en':'Please refer to the error code for exceptions.', 'zh_CN':'请参照错误码。','dictionary':'belong=WAAP-MS-Ext|dict=waap_retCodeEnum'}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {'en':'Description.', 'zh_CN':'描述信息。'}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
}

func (s UpdateCustomActionResponse) String() string {
  return tea.Prettify(s)
}

func (s UpdateCustomActionResponse) GoString() string {
  return s.String()
}

func (s *UpdateCustomActionResponse) SetCode(v string) *UpdateCustomActionResponse {
  s.Code = &v
  return s
}

func (s *UpdateCustomActionResponse) SetMsg(v string) *UpdateCustomActionResponse {
  s.Msg = &v
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

type UpdateCustomActionRequestHeader struct {
  // {"zh_CN":"安全服务类型。有使用多个不同的安全服务时，需要填写具体的服务类型。","en":"Security service type. Please enter a specific service type, if you purchase multiple security services.","dictionary":"belong=WAAP-MS-Ext|dict=waap_serviceType"}
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

type UpdateCustomActionResponseHeader struct {
}

func (s UpdateCustomActionResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s UpdateCustomActionResponseHeader) GoString() string {
  return s.String()
}




type DeleteCustomActionsRequest struct {
  // {'en':'Custom response ID list.', 'zh_CN':'自定义响应ID列表。'}
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

type DeleteCustomActionsResponse struct {
  // {'en':'Please refer to the error code for exceptions.', 'zh_CN':'请参照错误码。','dictionary':'belong=WAAP-MS-Ext|dict=waap_retCodeEnum'}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {'en':'Description.', 'zh_CN':'描述信息。'}
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

type DeleteCustomActionsRequestHeader struct {
  // {"zh_CN":"安全服务类型。有使用多个不同的安全服务时，需要填写具体的服务类型。","en":"Security service type. Please enter a specific service type, if you purchase multiple security services.","dictionary":"belong=WAAP-MS-Ext|dict=waap_serviceType"}
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

type DeleteCustomActionsResponseHeader struct {
}

func (s DeleteCustomActionsResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteCustomActionsResponseHeader) GoString() string {
  return s.String()
}




type QueryAppApiExceptionListRequest struct {
  // {'en':'Do you need to return the number of associated hostnames? The default is false', 'zh_CN':'是否需要返回关联域名数，默认为false'}
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

type QueryAppApiExceptionListResponse struct {
  // {'en':'Please refer to the error code for exceptions.', 'zh_CN':'请参照错误码。','dictionary':'belong=WAAP-MS-Ext|dict=waap_retCodeEnum'}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {'en':'Description.', 'zh_CN':'描述信息。'}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {'en':'Data.', 'zh_CN':'出参数据。'}
  Data []*QueryAppApiExceptionListDmsShareServiceFeatureListVO `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
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

func (s *QueryAppApiExceptionListResponse) SetData(v []*QueryAppApiExceptionListDmsShareServiceFeatureListVO) *QueryAppApiExceptionListResponse {
  s.Data = v
  return s
}

type QueryAppApiExceptionListDmsShareServiceFeatureListVO struct {
  // {'en':'Rule ID', 'zh_CN':'规则ID'}
  RuleId *string `json:"ruleId,omitempty" xml:"ruleId,omitempty" require:"true"`
  // {'en':'Rule name, duplicate not allowed.', 'zh_CN':'规则名称，不允许重复'}
  RuleName *string `json:"ruleName,omitempty" xml:"ruleName,omitempty" require:"true"`
  // {'en':'Characteristic.', 'zh_CN':'特征。'}
  Config *QueryAppApiExceptionListDmsShareServiceFeatureDetailVO `json:"config,omitempty" xml:"config,omitempty" require:"true"`
  // {
  // 'en':'Type enumeration, for example:
  // NATIVE_ APP        remarks: Native App
  // Callback_ API      	 remarks: Callback API
  // Other_ API       	  remarks: Other program API
  // HYBRID_ APP  	   remarks: Hybrid APP',
  // 'zh_CN':'类型枚举，例如：
  // NATIVE_APP          备注：原生APP
  // CALLBACK_API      备注：回调API
  // OTHER_API            备注：其他程序API
  // HYBRID_APP          备注：混合APP'}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
  // {'en':'Number of associated hostnames', 'zh_CN':'关联域名数量'}
  DomainNum *int32 `json:"domainNum,omitempty" xml:"domainNum,omitempty" require:"true"`
  // {'en':'Creator', 'zh_CN':'创建者'}
  Creator *string `json:"creator,omitempty" xml:"creator,omitempty" require:"true"`
}

func (s QueryAppApiExceptionListDmsShareServiceFeatureListVO) String() string {
  return tea.Prettify(s)
}

func (s QueryAppApiExceptionListDmsShareServiceFeatureListVO) GoString() string {
  return s.String()
}

func (s *QueryAppApiExceptionListDmsShareServiceFeatureListVO) SetRuleId(v string) *QueryAppApiExceptionListDmsShareServiceFeatureListVO {
  s.RuleId = &v
  return s
}

func (s *QueryAppApiExceptionListDmsShareServiceFeatureListVO) SetRuleName(v string) *QueryAppApiExceptionListDmsShareServiceFeatureListVO {
  s.RuleName = &v
  return s
}

func (s *QueryAppApiExceptionListDmsShareServiceFeatureListVO) SetConfig(v *QueryAppApiExceptionListDmsShareServiceFeatureDetailVO) *QueryAppApiExceptionListDmsShareServiceFeatureListVO {
  s.Config = v
  return s
}

func (s *QueryAppApiExceptionListDmsShareServiceFeatureListVO) SetType(v string) *QueryAppApiExceptionListDmsShareServiceFeatureListVO {
  s.Type = &v
  return s
}

func (s *QueryAppApiExceptionListDmsShareServiceFeatureListVO) SetDomainNum(v int32) *QueryAppApiExceptionListDmsShareServiceFeatureListVO {
  s.DomainNum = &v
  return s
}

func (s *QueryAppApiExceptionListDmsShareServiceFeatureListVO) SetCreator(v string) *QueryAppApiExceptionListDmsShareServiceFeatureListVO {
  s.Creator = &v
  return s
}

type QueryAppApiExceptionListDmsShareServiceFeatureDetailVO struct {
  // {'en':'Matching condition name.
  // PATH         Path
  // URI            URI
  // UA             User-Agent
  // REFERER    Referer
  // HEADER    Request Header
  // ', 'zh_CN':'匹配条件名称。
  // PATH                       路径
  // URI                          URI
  // UA                           User-Agent
  // REFERER                  Referer
  // HEAD                      Request Header
  // '}
  Condition *string `json:"condition,omitempty" xml:"condition,omitempty" require:"true"`
  // {'en':'Matching condition function.
  // CONTAIN                  contains
  // NOT_CONTAIN         does not contain
  // EQUAL                      equals
  // NOT_EQUAL             does not equal
  // EMPTY                      does not exist or has no value(When the condition is a path or URI, this value is not optional)
  // REGEX                      regex match
  // ', 'zh_CN':'匹配条件函数。
  // CONTAIN                  包含
  // NOT_CONTAIN         不包含 
  // EQUAL                      等于
  // NOT_EQUAL             不等于
  // EMPTY                      为空或不存在（当condition为路径或者URI时，该值不可选）
  // REGEX                       匹配正则 
  // '}
  MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
  // {'en':'Condition value list, When matchType is EMPTY, the value is empty;  When matchType is REGEX, only one item is set', 'zh_CN':'条件值列表，当matchType为EMPTY时，该值为空；当matchType为REGEX时，只能传一条'}
  ConditionValue []*string `json:"conditionValue,omitempty" xml:"conditionValue,omitempty" require:"true" type:"Repeated"`
  // {'en':'Head value, If condition=HEAD, then this field is mandatory.', 'zh_CN':'头部值，如果condition为HEAD，则该字段必填'}
  Key *string `json:"key,omitempty" xml:"key,omitempty" require:"true"`
}

func (s QueryAppApiExceptionListDmsShareServiceFeatureDetailVO) String() string {
  return tea.Prettify(s)
}

func (s QueryAppApiExceptionListDmsShareServiceFeatureDetailVO) GoString() string {
  return s.String()
}

func (s *QueryAppApiExceptionListDmsShareServiceFeatureDetailVO) SetCondition(v string) *QueryAppApiExceptionListDmsShareServiceFeatureDetailVO {
  s.Condition = &v
  return s
}

func (s *QueryAppApiExceptionListDmsShareServiceFeatureDetailVO) SetMatchType(v string) *QueryAppApiExceptionListDmsShareServiceFeatureDetailVO {
  s.MatchType = &v
  return s
}

func (s *QueryAppApiExceptionListDmsShareServiceFeatureDetailVO) SetConditionValue(v []*string) *QueryAppApiExceptionListDmsShareServiceFeatureDetailVO {
  s.ConditionValue = v
  return s
}

func (s *QueryAppApiExceptionListDmsShareServiceFeatureDetailVO) SetKey(v string) *QueryAppApiExceptionListDmsShareServiceFeatureDetailVO {
  s.Key = &v
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

type QueryAppApiExceptionListRequestHeader struct {
  // {"zh_CN":"安全服务类型。有使用多个不同的安全服务时，需要填写具体的服务类型。","en":"Security service type. Please enter a specific service type, if you purchase multiple security services.","dictionary":"belong=WAAP-MS-Ext|dict=waap_serviceType"}
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

type QueryAppApiExceptionListResponseHeader struct {
}

func (s QueryAppApiExceptionListResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryAppApiExceptionListResponseHeader) GoString() string {
  return s.String()
}




type CreateCustomActionRequest struct {
  // {'en':'Action name,maximum 50 characters.
  // Does not support special characters and spaces.', 'zh_CN':'动作名称，最多50个字符。
  // 不支持特殊字符和空格。'}
  ActionName *string `json:"actionName,omitempty" xml:"actionName,omitempty" require:"true"`
  // {'en':'Description,maximum 200 characters.', 'zh_CN':'描述，最多200个字符。'}
  Description *string `json:"description,omitempty" xml:"description,omitempty"`
  // {'en':'Status code.
  // Valid value range:200,204,206,400,401,403,404,500,501,503.', 'zh_CN':'状态码。
  // 有效值范围：200，204，206，400，401，403，404，500，501，503。'}
  StatusCode *string `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
  // {'en':'Response Content-Type type,multiple separated by ; sign.', 'zh_CN':'响应Content-Type类型，多个以 ; 隔开。'}
  ContentType *string `json:"contentType,omitempty" xml:"contentType,omitempty" require:"true"`
  // {'en':'Response content definition, the size of the response content cannot exceed 16KB, if you need to insert static resources such as pictures, it is recommended to use links to import. 
  // The following interception information parameters are supported: 
  //  {url} : Displays the URL information when intercepting 
  //  {client_ip} : Displays the intercepted user IP information 
  //  {time} : Displays the intercepted time point 
  //  {event_id} : Displays the ID information of this interception event', 'zh_CN':'自定义响应内容定义，响应内容大小不能超过16KB，如需插入图片等静态资源，建议采用链接引入。
  // 支持引用下列拦截信息参数：
  // {url} ：显示拦截时的URL信息
  // {client_ip} ：显示被拦截的用户IP信息
  // {time} ：显示被拦截的时间点
  // {event_id} ：显示本次拦截事件的ID信息'}
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

type CreateCustomActionResponse struct {
  // {'en':'Please refer to the error code for exceptions.', 'zh_CN':'请参照错误码。','dictionary':'belong=WAAP-MS-Ext|dict=waap_retCodeEnum'}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {'en':'Description.', 'zh_CN':'描述信息。'}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {'en':'Customize response ID.', 'zh_CN':'自定义响应ID。'}
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

type CreateCustomActionRequestHeader struct {
  // {"zh_CN":"安全服务类型。有使用多个不同的安全服务时，需要填写具体的服务类型。","en":"Security service type. Please enter a specific service type, if you purchase multiple security services.","dictionary":"belong=WAAP-MS-Ext|dict=waap_serviceType"}
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

type CreateCustomActionResponseHeader struct {
}

func (s CreateCustomActionResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s CreateCustomActionResponseHeader) GoString() string {
  return s.String()
}




type UpdateSharedWhitelistRuleRequest struct {
  // {"en":"Rule ID.", "zh_CN":"规则ID。"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {"en":"Rule name, maximum 50 characters.
  //  Does not support special characters and spaces.", "zh_CN":"规则名称，最多50个字符。
  // 不支持特殊字符和空格。"}
  RuleName *string `json:"ruleName,omitempty" xml:"ruleName,omitempty"`
  // {"en":"Description, maximum 200 characters.", "zh_CN":"描述，最多200个字符。"}
  Description *string `json:"description,omitempty" xml:"description,omitempty"`
  // {"en":"Match conditions, at least one, at most five.", "zh_CN":"匹配条件，至少一个，至多五个。"}
  Conditions *UpdateSharedWhitelistRuleWhitelistRuleCondition `json:"conditions,omitempty" xml:"conditions,omitempty"`
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

func (s *UpdateSharedWhitelistRuleRequest) SetConditions(v *UpdateSharedWhitelistRuleWhitelistRuleCondition) *UpdateSharedWhitelistRuleRequest {
  s.Conditions = v
  return s
}

type UpdateSharedWhitelistRuleWhitelistRuleCondition struct {
  // {"en":"IP/CIDR match conditions, match type cannot be repeated.", "zh_CN":"IP/IP段匹配条件，匹配类型不可重复。"}
  IpOrIpsConditions []*UpdateSharedWhitelistRuleIpOrIpsCondition `json:"ipOrIpsConditions,omitempty" xml:"ipOrIpsConditions,omitempty" type:"Repeated"`
  // {"en":"Path match conditions, match type cannot be repeated.", "zh_CN":"路径匹配条件，匹配类型不可重复。"}
  PathConditions []*UpdateSharedWhitelistRulePathCondition `json:"pathConditions,omitempty" xml:"pathConditions,omitempty" type:"Repeated"`
  // {"en":"URI match conditions, match type cannot be repeated.", "zh_CN":"URI匹配条件，匹配类型不可重复。"}
  UriConditions []*UpdateSharedWhitelistRuleUriCondition `json:"uriConditions,omitempty" xml:"uriConditions,omitempty" type:"Repeated"`
  // {"en":"User agent match conditions, match type cannot be repeated.", "zh_CN":"User-Agent 匹配条件，匹配类型不可重复。"}
  UaConditions []*UpdateSharedWhitelistRuleUaCondition `json:"uaConditions,omitempty" xml:"uaConditions,omitempty" type:"Repeated"`
  // {"en":"Referer match conditions, match type cannot be repeated.", "zh_CN":"Referer 匹配条件，匹配类型不可重复。"}
  RefererConditions []*UpdateSharedWhitelistRuleRefererCondition `json:"refererConditions,omitempty" xml:"refererConditions,omitempty" type:"Repeated"`
  // {"en":"Request header match conditions.", "zh_CN":"请求头匹配条件。"}
  HeaderConditions []*UpdateSharedWhitelistRuleHeaderCondition `json:"headerConditions,omitempty" xml:"headerConditions,omitempty" type:"Repeated"`
}

func (s UpdateSharedWhitelistRuleWhitelistRuleCondition) String() string {
  return tea.Prettify(s)
}

func (s UpdateSharedWhitelistRuleWhitelistRuleCondition) GoString() string {
  return s.String()
}

func (s *UpdateSharedWhitelistRuleWhitelistRuleCondition) SetIpOrIpsConditions(v []*UpdateSharedWhitelistRuleIpOrIpsCondition) *UpdateSharedWhitelistRuleWhitelistRuleCondition {
  s.IpOrIpsConditions = v
  return s
}

func (s *UpdateSharedWhitelistRuleWhitelistRuleCondition) SetPathConditions(v []*UpdateSharedWhitelistRulePathCondition) *UpdateSharedWhitelistRuleWhitelistRuleCondition {
  s.PathConditions = v
  return s
}

func (s *UpdateSharedWhitelistRuleWhitelistRuleCondition) SetUriConditions(v []*UpdateSharedWhitelistRuleUriCondition) *UpdateSharedWhitelistRuleWhitelistRuleCondition {
  s.UriConditions = v
  return s
}

func (s *UpdateSharedWhitelistRuleWhitelistRuleCondition) SetUaConditions(v []*UpdateSharedWhitelistRuleUaCondition) *UpdateSharedWhitelistRuleWhitelistRuleCondition {
  s.UaConditions = v
  return s
}

func (s *UpdateSharedWhitelistRuleWhitelistRuleCondition) SetRefererConditions(v []*UpdateSharedWhitelistRuleRefererCondition) *UpdateSharedWhitelistRuleWhitelistRuleCondition {
  s.RefererConditions = v
  return s
}

func (s *UpdateSharedWhitelistRuleWhitelistRuleCondition) SetHeaderConditions(v []*UpdateSharedWhitelistRuleHeaderCondition) *UpdateSharedWhitelistRuleWhitelistRuleCondition {
  s.HeaderConditions = v
  return s
}

type UpdateSharedWhitelistRuleIpOrIpsCondition struct {
  // {"en":"Match type.
  // EQUAL: Equals
  // NOT_EQUAL: Does not equal", "zh_CN":"匹配类型。
  // EQUAL：等于
  // NOT_EQUAL：不等于"}
  MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
  // {"en":"IP/CIDR, maximum 500 IP/CIDR.", "zh_CN":"IP/IP段，最多500个IP/IP段。"}
  IpOrIps []*string `json:"ipOrIps,omitempty" xml:"ipOrIps,omitempty" require:"true" type:"Repeated"`
}

func (s UpdateSharedWhitelistRuleIpOrIpsCondition) String() string {
  return tea.Prettify(s)
}

func (s UpdateSharedWhitelistRuleIpOrIpsCondition) GoString() string {
  return s.String()
}

func (s *UpdateSharedWhitelistRuleIpOrIpsCondition) SetMatchType(v string) *UpdateSharedWhitelistRuleIpOrIpsCondition {
  s.MatchType = &v
  return s
}

func (s *UpdateSharedWhitelistRuleIpOrIpsCondition) SetIpOrIps(v []*string) *UpdateSharedWhitelistRuleIpOrIpsCondition {
  s.IpOrIps = v
  return s
}

type UpdateSharedWhitelistRulePathCondition struct {
  // {"en":"Match type.
  // EQUAL: Equals, path case sensitive
  // NOT_EQUAL: Does not equal, path case sensitive
  // CONTAIN: Contains, path case insensitive
  // NOT_CONTAIN: Does not Contains, path case insensitive
  // REGEX: Regex match, path case insensitive
  // NOT_REGEX: Regular does not match, path case sensitive
  // START_WITH: Starts with, path case sensitive
  // END_WITH: Ends with, path case sensitive
  // WILDCARD: Wildcard matches, path case sensitive, ** represents zero or more arbitrary characters, ? represents any single character.
  // NOT_WILDCARD: Wildcard does not match, path case sensitive, ** represents zero or more arbitrary characters, ? represents any single character ", "zh_CN":"匹配类型。
  // EQUAL：等于，路径大小写敏感
  // NOT_EQUAL：不等于，路径大小写敏感
  // CONTAIN：包含，路径大小写不敏感
  // NOT_CONTAIN：不包含，路径大小写不敏感
  // REGEX：匹配正则，路径大小写不敏感
  // NOT_REGEX：正则不匹配，路径大小写不敏感
  // START_WITH：开头是，路径大小写不敏感
  // END_WITH：结尾是，路径大小写不敏感
  // WILDCARD：通配符匹配，路径大小写不敏感，*代表零个或多个任意字符，?代表任意单个字符
  // NOT_WILDCARD：通配符不匹配，路径大小写不敏感，*代表零个或多个任意字符，?代表任意单个字符"}
  MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
  // {'en':'Path.
  // When match type is EQUAL/NOT_EQUAL/START_WITH/END_WITH, path needs to start with "/", and no parameters.
  // When the match type is REGEX/NOT_REGEX, only one value is allowed. 
  // Example: /test.html.', 'zh_CN':'路径。
  // 当匹配类型为等于/不等于/开头是/结尾是，路径必须以“/”开头，不含参数。
  // 当匹配类型为正则/正则不匹配，则只允许只有一个值。
  // 示例：/test.html。'}
  UpdateSharedWhitelistRulePaths []*string `json:"paths,omitempty" xml:"paths,omitempty" require:"true" type:"Repeated"`
}

func (s UpdateSharedWhitelistRulePathCondition) String() string {
  return tea.Prettify(s)
}

func (s UpdateSharedWhitelistRulePathCondition) GoString() string {
  return s.String()
}

func (s *UpdateSharedWhitelistRulePathCondition) SetMatchType(v string) *UpdateSharedWhitelistRulePathCondition {
  s.MatchType = &v
  return s
}

func (s *UpdateSharedWhitelistRulePathCondition) SetUpdateSharedWhitelistRulePaths(v []*string) *UpdateSharedWhitelistRulePathCondition {
  s.UpdateSharedWhitelistRulePaths = v
  return s
}

type UpdateSharedWhitelistRuleUriCondition struct {
  // {"en":"Match type.
  // EQUAL: Equals, URI case sensitive
  // NOT_EQUAL: Does not equal, URI case sensitive
  // CONTAIN: Contains, URI case insensitive
  // NOT_CONTAIN: Does not Contains, URI case insensitive
  // REGEX: Regex match, URI case insensitive
  // NOT_REGEX: Regular does not match, URI case insensitive
  // START_WITH: Starts with, URI case insensitive
  // END_WITH: Ends with, URI case insensitive
  // WILDCARD: Wildcard matches, URI case insensitive, ** represents zero or more arbitrary characters, ? represents any single character
  // NOT_WILDCARD: Wildcard does not match, URI case insensitive, ** represents zero or more arbitrary characters, ? represents any single character", "zh_CN":"匹配类型。
  // EQUAL：等于，URI大小写敏感
  // NOT_EQUAL：不等于，URI大小写敏感
  // CONTAIN：包含，URI大小写不敏感
  // NOT_CONTAIN：不包含，URI大小写不敏感
  // REGEX：匹配正则，URI大小写不敏感
  // NOT_REGEX：正则不匹配，URI大小写不敏感
  // START_WITH：开头是，URI大小写不敏感
  // END_WITH：结尾是，URI大小写不敏感
  // WILDCARD：通配符匹配，URI大小写不敏感，*代表零个或多个任意字符，?代表任意单个字符
  // NOT_WILDCARD：通配符不匹配，URI大小写不敏感，*代表零个或多个任意字符，?代表任意单个字符"}
  MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
  // {'en':'URI.
  // When match type is EQUAL/NOT_EQUAL/START_WITH/END_WITH, uri needs to start with "/", and includes parameters.
  // When the match type is REGEX/NOT_REGEX, only one value is allowed. 
  // Example: /test.html?id=1.', 'zh_CN':'URI。
  // 当匹配类型为等于/不等于/开头是/结尾是，URI必须以”/“开头，含参数。
  // 当匹配类型为正则/正则不匹配，则只允许只有一个值。
  // 示例：/test.html?id=1。'}
  Uri []*string `json:"uri,omitempty" xml:"uri,omitempty" require:"true" type:"Repeated"`
}

func (s UpdateSharedWhitelistRuleUriCondition) String() string {
  return tea.Prettify(s)
}

func (s UpdateSharedWhitelistRuleUriCondition) GoString() string {
  return s.String()
}

func (s *UpdateSharedWhitelistRuleUriCondition) SetMatchType(v string) *UpdateSharedWhitelistRuleUriCondition {
  s.MatchType = &v
  return s
}

func (s *UpdateSharedWhitelistRuleUriCondition) SetUri(v []*string) *UpdateSharedWhitelistRuleUriCondition {
  s.Uri = v
  return s
}

type UpdateSharedWhitelistRuleUaCondition struct {
  // {"en":"Match type.
  // EQUAL: Equals, user agent case sensitive
  // NOT_EQUAL: Does not equal, user agent case sensitive
  // CONTAIN: Contains, user agent case insensitive
  // NOT_CONTAIN: Does not Contains, user agent case insensitive
  // REGEX: Regex match, user agent case insensitive
  // NOT_REGEX: Regular does not match, user agent case insensitive
  // START_WITH: Starts with, user agent case insensitive
  // END_WITH: Ends with, user agent case insensitive
  // WILDCARD: Wildcard matches, user agent case insensitive, ** represents zero or more arbitrary characters, ? represents any single character
  // NOT_WILDCARD: Wildcard does not match, user agent case insensitive, ** represents zero or more arbitrary characters, ? represents any single character", "zh_CN":"匹配类型。
  // EQUAL：等于，User-Agent大小写敏感
  // NOT_EQUAL：不等于，User-Agent大小写敏感
  // CONTAIN：包含，User-Agent大小写不敏感
  // NOT_CONTAIN：不包含，User-Agent大小写不敏感
  // REGEX：匹配正则，User-Agent大小写不敏感
  // NOT_REGEX：正则不匹配，User-Agent大小写不敏感
  // START_WITH：开头是，User-Agent大小写不敏感
  // END_WITH：结尾是，User-Agent大小写不敏感
  // WILDCARD：通配符匹配，User-Agent大小写不敏感，*代表零个或多个任意字符，?代表任意单个字符
  // NOT_WILDCARD：通配符不匹配，User-Agent大小写不敏感，*代表零个或多个任意字符，?代表任意单个字符"}
  MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
  // {'en':'User agent.
  // When the match type is REGEX/NOT_REGEX, only one value is allowed. 
  // Example: go-Http-client/1.1.', 'zh_CN':'User-Agent。
  // 当匹配类型为正则/正则不匹配，则只允许只有一个值。
  // 示例：go-Http-client/1.1。'}
  Ua []*string `json:"ua,omitempty" xml:"ua,omitempty" require:"true" type:"Repeated"`
}

func (s UpdateSharedWhitelistRuleUaCondition) String() string {
  return tea.Prettify(s)
}

func (s UpdateSharedWhitelistRuleUaCondition) GoString() string {
  return s.String()
}

func (s *UpdateSharedWhitelistRuleUaCondition) SetMatchType(v string) *UpdateSharedWhitelistRuleUaCondition {
  s.MatchType = &v
  return s
}

func (s *UpdateSharedWhitelistRuleUaCondition) SetUa(v []*string) *UpdateSharedWhitelistRuleUaCondition {
  s.Ua = v
  return s
}

type UpdateSharedWhitelistRuleRefererCondition struct {
  // {"en":"Match type.
  // EQUAL: Equals, referer case sensitive
  // NOT_EQUAL: Does not equal, referer case sensitive
  // CONTAIN: Contains, referer case insensitive
  // NOT_CONTAIN: Does not Contains, referer case insensitive
  // REGEX: Regex match, referer case insensitive
  // NOT_REGEX: Regular does not match, referer case insensitive
  // START_WITH: Starts with, referer case insensitive
  // END_WITH: Ends with, referer case insensitive
  // WILDCARD: Wildcard matches, referer case insensitive, ** represents zero or more arbitrary characters, ? represents any single characte
  // NOT_WILDCARD: Wildcard does not match, referer case insensitive, ** represents zero or more arbitrary characters, ? represents any single character", "zh_CN":"匹配类型。
  // EQUAL：等于，referer大小写敏感
  // NOT_EQUAL：不等于，referer大小写敏感
  // CONTAIN：包含，referer大小写不敏感
  // NOT_CONTAIN：不包含，referer大小写不敏感
  // REGEX：匹配正则，referer大小写不敏感
  // NOT_REGEX：正则不匹配，referer大小写不敏感
  // START_WITH：开头是，referer大小写不敏感
  // END_WITH：结尾是，referer大小写不敏感
  // WILDCARD：通配符匹配，referer大小写不敏感，*代表零个或多个任意字符，?代表任意单个字符
  // NOT_WILDCARD：通配符不匹配，referer大小写不敏感，*代表零个或多个任意字符，?代表任意单个字符"}
  MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
  // {'en':'Referer.
  // When the match type is REGEX/NOT_REGEX, only one value is allowed. 
  // Example: http://test.com.', 'zh_CN':'Referer。
  // 当匹配类型为正则/正则不匹配，则只允许只有一个值。
  // 示例：http://test.com。'}
  Referer []*string `json:"referer,omitempty" xml:"referer,omitempty" require:"true" type:"Repeated"`
}

func (s UpdateSharedWhitelistRuleRefererCondition) String() string {
  return tea.Prettify(s)
}

func (s UpdateSharedWhitelistRuleRefererCondition) GoString() string {
  return s.String()
}

func (s *UpdateSharedWhitelistRuleRefererCondition) SetMatchType(v string) *UpdateSharedWhitelistRuleRefererCondition {
  s.MatchType = &v
  return s
}

func (s *UpdateSharedWhitelistRuleRefererCondition) SetReferer(v []*string) *UpdateSharedWhitelistRuleRefererCondition {
  s.Referer = v
  return s
}

type UpdateSharedWhitelistRuleHeaderCondition struct {
  // {"en":"Match type.
  // EQUAL: Equals, request header values case sensitive
  // NOT_EQUAL: Does not equal, request header values case sensitive
  // CONTAIN: Contains, request header values case insensitive
  // NOT_CONTAIN: Does not Contains, request header values case insensitive
  // REGEX: Regex match, request header values case insensitive
  // NOT_REGEX: Regular does not match, request header values case insensitive
  // START_WITH: Starts with, request header values case insensitive
  // END_WITH: Ends with, request header values case insensitive
  // WILDCARD: Wildcard matches, request header values case insensitive, ** represents zero or more arbitrary characters, ? represents any single character
  // NOT_WILDCARD: Wildcard does not match, request header values case insensitive, ** represents zero or more arbitrary characters, ? represents any single character", "zh_CN":"匹配类型。
  // EQUAL：等于，头部值大小写敏感
  // NOT_EQUAL：不等于，头部值大小写敏感
  // CONTAIN：包含，头部值大小写不敏感
  // NOT_CONTAIN：不包含，头部值大小写不敏感
  // REGEX：匹配正则，头部值大小写不敏感
  // NOT_REGEX：正则不匹配，头部值大小写不敏感
  // START_WITH：开头是，头部值大小写不敏感
  // END_WITH：结尾是，头部值大小写不敏感
  // WILDCARD：通配符匹配，头部值大小写不敏感，*代表零个或多个任意字符，?代表任意单个字符
  // NOT_WILDCARD：通配符不匹配，头部值大小写不敏感，*代表零个或多个任意字符，?代表任意单个字符"}
  MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
  // {'en':'Header name,case insensitive,up to 100 characters.
  // Example: Accept.', 'zh_CN':'头部名称，大小写不敏感，最多100个字符。
  // 示例：Accept。'}
  Key *string `json:"key,omitempty" xml:"key,omitempty" require:"true"`
  // {'en':'Header value.
  // When the match type is REGEX/NOT_REGEX, only one value is allowed.', 'zh_CN':'头部值。
  // 当匹配类型为正则/正则不匹配，则只允许只有一个值。'}
  ValueList []*string `json:"valueList,omitempty" xml:"valueList,omitempty" require:"true" type:"Repeated"`
}

func (s UpdateSharedWhitelistRuleHeaderCondition) String() string {
  return tea.Prettify(s)
}

func (s UpdateSharedWhitelistRuleHeaderCondition) GoString() string {
  return s.String()
}

func (s *UpdateSharedWhitelistRuleHeaderCondition) SetMatchType(v string) *UpdateSharedWhitelistRuleHeaderCondition {
  s.MatchType = &v
  return s
}

func (s *UpdateSharedWhitelistRuleHeaderCondition) SetKey(v string) *UpdateSharedWhitelistRuleHeaderCondition {
  s.Key = &v
  return s
}

func (s *UpdateSharedWhitelistRuleHeaderCondition) SetValueList(v []*string) *UpdateSharedWhitelistRuleHeaderCondition {
  s.ValueList = v
  return s
}

type UpdateSharedWhitelistRuleResponse struct {
  // {'en':'Please refer to the error code for exceptions.', 'zh_CN':'请参照错误码。','dictionary':'belong=WAAP-MS-Ext|dict=waap_retCodeEnum'}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Description.", "zh_CN":"描述信息。"}
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

type UpdateSharedWhitelistRuleRequestHeader struct {
  // {"zh_CN":"安全服务类型。有使用多个不同的安全服务时，需要填写具体的服务类型。","en":"Security service type. Please enter a specific service type, if you purchase multiple security services.","dictionary":"belong=WAAP-MS-Ext|dict=waap_serviceType"}
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

type UpdateSharedWhitelistRuleResponseHeader struct {
}

func (s UpdateSharedWhitelistRuleResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s UpdateSharedWhitelistRuleResponseHeader) GoString() string {
  return s.String()
}




type QueryAppApiExceptionFeatureDetailRequest struct {
  // {'en':'Rule Id', 'zh_CN':'规则 ID'}
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

type QueryAppApiExceptionFeatureDetailResponse struct {
  // {'en':'Please refer to the error code for exceptions.', 'zh_CN':'请参照错误码。','dictionary':'belong=WAAP-MS-Ext|dict=waap_retCodeEnum'}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {'en':'Description.', 'zh_CN':'描述信息。'}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {'en':'Data.', 'zh_CN':'出参数据。'}
  Data *QueryAppApiExceptionFeatureDetailDmsShareServiceFeatureVO `json:"data,omitempty" xml:"data,omitempty" require:"true"`
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

func (s *QueryAppApiExceptionFeatureDetailResponse) SetData(v *QueryAppApiExceptionFeatureDetailDmsShareServiceFeatureVO) *QueryAppApiExceptionFeatureDetailResponse {
  s.Data = v
  return s
}

type QueryAppApiExceptionFeatureDetailDmsShareServiceFeatureVO struct {
  // {'en':'Rule name, duplicate not allowed.', 'zh_CN':'规则名称，不允许重复'}
  RuleName *string `json:"ruleName,omitempty" xml:"ruleName,omitempty" require:"true"`
  // {'en':'Characteristic.', 'zh_CN':'特征。'}
  Config *QueryAppApiExceptionFeatureDetailDmsShareServiceFeatureDetailVO `json:"config,omitempty" xml:"config,omitempty" require:"true"`
  // {
  // 'en':'Type enumeration, for example:
  // NATIVE_ APP        remarks: Native App
  // Callback_ API      	 remarks: Callback API
  // Other_ API       	  remarks: Other program API
  // HYBRID_ APP  	   remarks: Hybrid APP',
  // 'zh_CN':'类型枚举，例如：
  // NATIVE_APP          备注：原生APP
  // CALLBACK_API      备注：回调API
  // OTHER_API            备注：其他程序API
  // HYBRID_APP          备注：混合APP'}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
}

func (s QueryAppApiExceptionFeatureDetailDmsShareServiceFeatureVO) String() string {
  return tea.Prettify(s)
}

func (s QueryAppApiExceptionFeatureDetailDmsShareServiceFeatureVO) GoString() string {
  return s.String()
}

func (s *QueryAppApiExceptionFeatureDetailDmsShareServiceFeatureVO) SetRuleName(v string) *QueryAppApiExceptionFeatureDetailDmsShareServiceFeatureVO {
  s.RuleName = &v
  return s
}

func (s *QueryAppApiExceptionFeatureDetailDmsShareServiceFeatureVO) SetConfig(v *QueryAppApiExceptionFeatureDetailDmsShareServiceFeatureDetailVO) *QueryAppApiExceptionFeatureDetailDmsShareServiceFeatureVO {
  s.Config = v
  return s
}

func (s *QueryAppApiExceptionFeatureDetailDmsShareServiceFeatureVO) SetType(v string) *QueryAppApiExceptionFeatureDetailDmsShareServiceFeatureVO {
  s.Type = &v
  return s
}

type QueryAppApiExceptionFeatureDetailDmsShareServiceFeatureDetailVO struct {
  // {'en':'Matching condition name.
  // PATH         Path
  // URI            URI
  // UA             User-Agent
  // REFERER    Referer
  // HEADER    Request Header
  // ', 'zh_CN':'匹配条件名称。
  // PATH                       路径
  // URI                          URI
  // UA                           User-Agent
  // REFERER                  Referer
  // HEAD                      Request Header
  // '}
  Condition *string `json:"condition,omitempty" xml:"condition,omitempty" require:"true"`
  // {'en':'Matching condition function.
  // CONTAIN                  contains
  // NOT_CONTAIN         does not contain
  // EQUAL                      equals
  // NOT_EQUAL             does not equal
  // EMPTY                      does not exist or has no value(When the condition is a path or URI, this value is not optional)
  // REGEX                      regex match
  // ', 'zh_CN':'匹配条件函数。
  // CONTAIN                  包含
  // NOT_CONTAIN         不包含 
  // EQUAL                      等于
  // NOT_EQUAL             不等于
  // EMPTY                      为空或不存在（当condition为路径或者URI时，该值不可选）
  // REGEX                       匹配正则 
  // '}
  MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
  // {'en':'Condition value list, When matchType is EMPTY, the value is empty;  When matchType is REGEX, only one item is set', 'zh_CN':'条件值列表，当matchType为EMPTY时，该值为空；当matchType为REGEX时，只能传一条'}
  ConditionValue []*string `json:"conditionValue,omitempty" xml:"conditionValue,omitempty" require:"true" type:"Repeated"`
  // {'en':'Head value, If condition=HEAD, then this field is mandatory.', 'zh_CN':'头部值，如果condition为HEAD，则该字段必填'}
  Key *string `json:"key,omitempty" xml:"key,omitempty" require:"true"`
}

func (s QueryAppApiExceptionFeatureDetailDmsShareServiceFeatureDetailVO) String() string {
  return tea.Prettify(s)
}

func (s QueryAppApiExceptionFeatureDetailDmsShareServiceFeatureDetailVO) GoString() string {
  return s.String()
}

func (s *QueryAppApiExceptionFeatureDetailDmsShareServiceFeatureDetailVO) SetCondition(v string) *QueryAppApiExceptionFeatureDetailDmsShareServiceFeatureDetailVO {
  s.Condition = &v
  return s
}

func (s *QueryAppApiExceptionFeatureDetailDmsShareServiceFeatureDetailVO) SetMatchType(v string) *QueryAppApiExceptionFeatureDetailDmsShareServiceFeatureDetailVO {
  s.MatchType = &v
  return s
}

func (s *QueryAppApiExceptionFeatureDetailDmsShareServiceFeatureDetailVO) SetConditionValue(v []*string) *QueryAppApiExceptionFeatureDetailDmsShareServiceFeatureDetailVO {
  s.ConditionValue = v
  return s
}

func (s *QueryAppApiExceptionFeatureDetailDmsShareServiceFeatureDetailVO) SetKey(v string) *QueryAppApiExceptionFeatureDetailDmsShareServiceFeatureDetailVO {
  s.Key = &v
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

type QueryAppApiExceptionFeatureDetailRequestHeader struct {
  // {"zh_CN":"安全服务类型。有使用多个不同的安全服务时，需要填写具体的服务类型。","en":"Security service type. Please enter a specific service type, if you purchase multiple security services.","dictionary":"belong=WAAP-MS-Ext|dict=waap_serviceType"}
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

type QueryAppApiExceptionFeatureDetailResponseHeader struct {
}

func (s QueryAppApiExceptionFeatureDetailResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryAppApiExceptionFeatureDetailResponseHeader) GoString() string {
  return s.String()
}




type DeleteAppApiExceptionFeatureRequest struct {
  // {'en':'Rule Ids', 'zh_CN':'规则ID列表'}
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

type DeleteAppApiExceptionFeatureResponse struct {
  // {'en':'Please refer to the error code for exceptions.', 'zh_CN':'请参照错误码。','dictionary':'belong=WAAP-MS-Ext|dict=waap_retCodeEnum'}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {'en':'Description.', 'zh_CN':'描述信息。'}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {'en':'Data.', 'zh_CN':'出参数据。'}
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

type DeleteAppApiExceptionFeatureRequestHeader struct {
  // {"zh_CN":"安全服务类型。有使用多个不同的安全服务时，需要填写具体的服务类型。","en":"Security service type. Please enter a specific service type, if you purchase multiple security services.","dictionary":"belong=WAAP-MS-Ext|dict=waap_serviceType"}
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

type DeleteAppApiExceptionFeatureResponseHeader struct {
}

func (s DeleteAppApiExceptionFeatureResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteAppApiExceptionFeatureResponseHeader) GoString() string {
  return s.String()
}




type ListSharedRateLimitingRulesRequest struct {
  // {'en':'Rule name.', 'zh_CN':'规则名称。'}
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

type ListSharedRateLimitingRulesResponse struct {
  // {'en':'Please refer to the error code for exceptions.', 'zh_CN':'请参照错误码。','dictionary':'belong=WAAP-MS-Ext|dict=waap_retCodeEnum'}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {'en':'Description.', 'zh_CN':'描述信息。'}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {'en':'Data.', 'zh_CN':'出参数据。'}
  Data []*ListSharedRateLimitingRulesCommonShareRateLimitRuleVO `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
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

func (s *ListSharedRateLimitingRulesResponse) SetData(v []*ListSharedRateLimitingRulesCommonShareRateLimitRuleVO) *ListSharedRateLimitingRulesResponse {
  s.Data = v
  return s
}

type ListSharedRateLimitingRulesRefererCondition struct {
  // {'en':'Match type.
  // EQUAL:Equals
  // NOT_EQUAL:Does not equal
  // CONTAIN:Contains
  // NOT_CONTAIN:Does not contains
  // REGEX:Regex match
  // NONE:Empty or non-existent
  // NOT_REGEX: regular does not match
  // START_WITH: starts with
  // END_WITH: ends with
  // WILDCARD: wildcard matches
  // NOT_WILDCARD: wildcard does not match', 'zh_CN':'匹配类型。
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
  // NOT_WILDCARD：通配符不匹配'}
  MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
  // {'en':'Referer.', 'zh_CN':'Referer。'}
  Referer []*string `json:"referer,omitempty" xml:"referer,omitempty" type:"Repeated"`
}

func (s ListSharedRateLimitingRulesRefererCondition) String() string {
  return tea.Prettify(s)
}

func (s ListSharedRateLimitingRulesRefererCondition) GoString() string {
  return s.String()
}

func (s *ListSharedRateLimitingRulesRefererCondition) SetMatchType(v string) *ListSharedRateLimitingRulesRefererCondition {
  s.MatchType = &v
  return s
}

func (s *ListSharedRateLimitingRulesRefererCondition) SetReferer(v []*string) *ListSharedRateLimitingRulesRefererCondition {
  s.Referer = v
  return s
}

type ListSharedRateLimitingRulesAreaCondition struct {
  // {'en':'Match type.
  // EQUAL:Equal
  // NOT_EQUAL:Does not equal', 'zh_CN':'匹配类型。
  // EQUAL：等于
  // NOT_EQUAL：不等于'}
  MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
  // {'en':'Geo.', 'zh_CN':'区域。','dictionary':'belong=WAAP-MS-Ext|dict=waap_areaCityAndCountry'}
  Areas []*string `json:"areas,omitempty" xml:"areas,omitempty" require:"true" type:"Repeated"`
}

func (s ListSharedRateLimitingRulesAreaCondition) String() string {
  return tea.Prettify(s)
}

func (s ListSharedRateLimitingRulesAreaCondition) GoString() string {
  return s.String()
}

func (s *ListSharedRateLimitingRulesAreaCondition) SetMatchType(v string) *ListSharedRateLimitingRulesAreaCondition {
  s.MatchType = &v
  return s
}

func (s *ListSharedRateLimitingRulesAreaCondition) SetAreas(v []*string) *ListSharedRateLimitingRulesAreaCondition {
  s.Areas = v
  return s
}

type ListSharedRateLimitingRulesIpOrIpsCondition struct {
  // {'en':'Match type.
  // EQUAL:Equal
  // NOT_EQUAL:Does not equal', 'zh_CN':'匹配类型。
  // EQUAL：等于
  // NOT_EQUAL：不等于'}
  MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
  // {'en':'IP/CIDR.', 'zh_CN':'IP/IP段。'}
  IpOrIps []*string `json:"ipOrIps,omitempty" xml:"ipOrIps,omitempty" require:"true" type:"Repeated"`
}

func (s ListSharedRateLimitingRulesIpOrIpsCondition) String() string {
  return tea.Prettify(s)
}

func (s ListSharedRateLimitingRulesIpOrIpsCondition) GoString() string {
  return s.String()
}

func (s *ListSharedRateLimitingRulesIpOrIpsCondition) SetMatchType(v string) *ListSharedRateLimitingRulesIpOrIpsCondition {
  s.MatchType = &v
  return s
}

func (s *ListSharedRateLimitingRulesIpOrIpsCondition) SetIpOrIps(v []*string) *ListSharedRateLimitingRulesIpOrIpsCondition {
  s.IpOrIps = v
  return s
}

type ListSharedRateLimitingRulesRequestMethodCondition struct {
  // {'en':'Match type.
  // EQUAL:Equal
  // NOT_EQUAL:Does not equal', 'zh_CN':'匹配类型。
  // EQUAL：等于
  // NOT_EQUAL：不等于'}
  MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
  // {'en':'Request method.', 'zh_CN':'请求方法。'}
  RequestMethod []*string `json:"requestMethod,omitempty" xml:"requestMethod,omitempty" require:"true" type:"Repeated"`
}

func (s ListSharedRateLimitingRulesRequestMethodCondition) String() string {
  return tea.Prettify(s)
}

func (s ListSharedRateLimitingRulesRequestMethodCondition) GoString() string {
  return s.String()
}

func (s *ListSharedRateLimitingRulesRequestMethodCondition) SetMatchType(v string) *ListSharedRateLimitingRulesRequestMethodCondition {
  s.MatchType = &v
  return s
}

func (s *ListSharedRateLimitingRulesRequestMethodCondition) SetRequestMethod(v []*string) *ListSharedRateLimitingRulesRequestMethodCondition {
  s.RequestMethod = v
  return s
}

type ListSharedRateLimitingRulesCommonShareRateLimitRuleVO struct {
  // {'en':'Rule ID.', 'zh_CN':'规则ID。'}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {'en':'Creator.', 'zh_CN':'创建者。'}
  Creator *string `json:"creator,omitempty" xml:"creator,omitempty" require:"true"`
  // {'en':'Rule Name.', 'zh_CN':'规则名称。'}
  RuleName *string `json:"ruleName,omitempty" xml:"ruleName,omitempty" require:"true"`
  // {'en':'Description.', 'zh_CN':'规则描述。'}
  Description *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
  // {'en':'Statistical stage.
  // REQUEST:Request
  // RESPONSE:Response', 'zh_CN':'统计阶段。
  // REQUEST：请求
  // RESPONSE：响应'}
  StatisticalStage *string `json:"statisticalStage,omitempty" xml:"statisticalStage,omitempty" require:"true"`
  // {'en':'Client identifier.
  // IP:Client IP
  // IP_UA:Client IP and User-Agent
  // COOKIE:Cookie
  // IP_COOKIE:Client IP and Cookie
  // HEADER:Request Header
  // IP_HEADER:Client IP and Request Header', 'zh_CN':'统计粒度。
  // IP：客户端IP
  // IP_UA：客户端IP和User-Agent
  // COOKIE：Cookie
  // IP_COOKIE：客户端IP和Cookie
  // HEADER：请求头
  // IP_HEADER：客户端IP和请求头'}
  StatisticalItem *string `json:"statisticalItem,omitempty" xml:"statisticalItem,omitempty" require:"true"`
  // {'en':'Statistical key value.', 'zh_CN':'统计key值。'}
  StatisticsKey *string `json:"statisticsKey,omitempty" xml:"statisticsKey,omitempty" require:"true"`
  // {'en':'Statistics period, unit: seconds.', 'zh_CN':'统计周期，单位：秒。'}
  StatisticalPeriod *int32 `json:"statisticalPeriod,omitempty" xml:"statisticalPeriod,omitempty" require:"true"`
  // {'en':'Trigger threshold, unit: times.', 'zh_CN':'触发阈值，单位：次。'}
  TriggerThreshold *int32 `json:"triggerThreshold,omitempty" xml:"triggerThreshold,omitempty" require:"true"`
  // {'en':'Action duration, unit: seconds.', 'zh_CN':'处理动作持续时间，单位：秒。'}
  InterceptTime *int32 `json:"interceptTime,omitempty" xml:"interceptTime,omitempty" require:"true"`
  // {'en':'Cycle effective status.
  // PERMANENT:All time
  // WITHOUT:Excluded time
  // WITHIN:Selected time', 'zh_CN':'周期生效状态。
  // PERMANENT：永久生效
  // WITHOUT：周期内不生效
  // WITHIN：周期内生效'}
  EffectiveStatus *string `json:"effectiveStatus,omitempty" xml:"effectiveStatus,omitempty" require:"true"`
  // {'en':'Effective time period.', 'zh_CN':'规则生效周期。'}
  ListSharedRateLimitingRulesRateLimitEffective *ListSharedRateLimitingRulesRateLimitEffective `json:"rateLimitEffective,omitempty" xml:"rateLimitEffective,omitempty" require:"true"`
  // {'en':'Action.
  // NO_USE:Not Used
  // LOG:Log
  // COOKIE:Cookie verification
  // JS_CHECK:Javascript verification
  // DELAY:Delay
  // BLOCK:Deny
  // RESET:Reset Connection
  // Custom action primary key id:Custom action primary key id', 'zh_CN':'处理动作。
  // NO_USE：不使用
  // LOG：监控
  // COOKIE：Cookie校验
  // JS_CHECK：JavaScript校验
  // DELAY：延迟响应
  // BLOCK：拦截
  // RESET：断开连接
  // 自定义处理动作主键id：自定义处理动作主键id'}
  Action *string `json:"action,omitempty" xml:"action,omitempty" require:"true"`
  // {'en':'Matching conditions.', 'zh_CN':'匹配条件。'}
  RateLimitRuleCondition *ListSharedRateLimitingRulesShareRateLimitRuleCondition `json:"rateLimitRuleCondition,omitempty" xml:"rateLimitRuleCondition,omitempty" require:"true"`
  // {'en':'Associated hostnames.', 'zh_CN':'关联域名数。'}
  CountShareRateLimitDomain *int32 `json:"countShareRateLimitDomain,omitempty" xml:"countShareRateLimitDomain,omitempty" require:"true"`
  // {'en':'List of associated hostnames.', 'zh_CN':'关联域名列表。'}
  ShareRateLimitRelDomains []*string `json:"shareRateLimitRelDomains,omitempty" xml:"shareRateLimitRelDomains,omitempty" require:"true" type:"Repeated"`
  // {'en':'Update time.', 'zh_CN':'更新时间。'}
  UpdateTime *string `json:"updateTime,omitempty" xml:"updateTime,omitempty" require:"true"`
}

func (s ListSharedRateLimitingRulesCommonShareRateLimitRuleVO) String() string {
  return tea.Prettify(s)
}

func (s ListSharedRateLimitingRulesCommonShareRateLimitRuleVO) GoString() string {
  return s.String()
}

func (s *ListSharedRateLimitingRulesCommonShareRateLimitRuleVO) SetId(v string) *ListSharedRateLimitingRulesCommonShareRateLimitRuleVO {
  s.Id = &v
  return s
}

func (s *ListSharedRateLimitingRulesCommonShareRateLimitRuleVO) SetCreator(v string) *ListSharedRateLimitingRulesCommonShareRateLimitRuleVO {
  s.Creator = &v
  return s
}

func (s *ListSharedRateLimitingRulesCommonShareRateLimitRuleVO) SetRuleName(v string) *ListSharedRateLimitingRulesCommonShareRateLimitRuleVO {
  s.RuleName = &v
  return s
}

func (s *ListSharedRateLimitingRulesCommonShareRateLimitRuleVO) SetDescription(v string) *ListSharedRateLimitingRulesCommonShareRateLimitRuleVO {
  s.Description = &v
  return s
}

func (s *ListSharedRateLimitingRulesCommonShareRateLimitRuleVO) SetStatisticalStage(v string) *ListSharedRateLimitingRulesCommonShareRateLimitRuleVO {
  s.StatisticalStage = &v
  return s
}

func (s *ListSharedRateLimitingRulesCommonShareRateLimitRuleVO) SetStatisticalItem(v string) *ListSharedRateLimitingRulesCommonShareRateLimitRuleVO {
  s.StatisticalItem = &v
  return s
}

func (s *ListSharedRateLimitingRulesCommonShareRateLimitRuleVO) SetStatisticsKey(v string) *ListSharedRateLimitingRulesCommonShareRateLimitRuleVO {
  s.StatisticsKey = &v
  return s
}

func (s *ListSharedRateLimitingRulesCommonShareRateLimitRuleVO) SetStatisticalPeriod(v int32) *ListSharedRateLimitingRulesCommonShareRateLimitRuleVO {
  s.StatisticalPeriod = &v
  return s
}

func (s *ListSharedRateLimitingRulesCommonShareRateLimitRuleVO) SetTriggerThreshold(v int32) *ListSharedRateLimitingRulesCommonShareRateLimitRuleVO {
  s.TriggerThreshold = &v
  return s
}

func (s *ListSharedRateLimitingRulesCommonShareRateLimitRuleVO) SetInterceptTime(v int32) *ListSharedRateLimitingRulesCommonShareRateLimitRuleVO {
  s.InterceptTime = &v
  return s
}

func (s *ListSharedRateLimitingRulesCommonShareRateLimitRuleVO) SetEffectiveStatus(v string) *ListSharedRateLimitingRulesCommonShareRateLimitRuleVO {
  s.EffectiveStatus = &v
  return s
}

func (s *ListSharedRateLimitingRulesCommonShareRateLimitRuleVO) SetRateLimitEffective(v *ListSharedRateLimitingRulesRateLimitEffective) *ListSharedRateLimitingRulesCommonShareRateLimitRuleVO {
  s.ListSharedRateLimitingRulesRateLimitEffective = v
  return s
}

func (s *ListSharedRateLimitingRulesCommonShareRateLimitRuleVO) SetAction(v string) *ListSharedRateLimitingRulesCommonShareRateLimitRuleVO {
  s.Action = &v
  return s
}

func (s *ListSharedRateLimitingRulesCommonShareRateLimitRuleVO) SetRateLimitRuleCondition(v *ListSharedRateLimitingRulesShareRateLimitRuleCondition) *ListSharedRateLimitingRulesCommonShareRateLimitRuleVO {
  s.RateLimitRuleCondition = v
  return s
}

func (s *ListSharedRateLimitingRulesCommonShareRateLimitRuleVO) SetCountShareRateLimitDomain(v int32) *ListSharedRateLimitingRulesCommonShareRateLimitRuleVO {
  s.CountShareRateLimitDomain = &v
  return s
}

func (s *ListSharedRateLimitingRulesCommonShareRateLimitRuleVO) SetShareRateLimitRelDomains(v []*string) *ListSharedRateLimitingRulesCommonShareRateLimitRuleVO {
  s.ShareRateLimitRelDomains = v
  return s
}

func (s *ListSharedRateLimitingRulesCommonShareRateLimitRuleVO) SetUpdateTime(v string) *ListSharedRateLimitingRulesCommonShareRateLimitRuleVO {
  s.UpdateTime = &v
  return s
}

type ListSharedRateLimitingRulesPathCondition struct {
  // {'en':'Match type.
  // EQUAL:Equals
  // NOT_EQUAL:Does not equal
  // CONTAIN:Contains
  // NOT_CONTAIN:Does not contains
  // REGEX:Regex match
  // NOT_REGEX: regular does not match
  // START_WITH: starts with
  // END_WITH: ends with
  // WILDCARD: wildcard matches
  // NOT_WILDCARD: wildcard does not match', 'zh_CN':'匹配类型。
  // EQUAL：等于
  // NOT_EQUAL：不等于
  // CONTAIN：包含
  // NOT_CONTAIN：不包含
  // REGEX：正则
  // NOT_REGEX：正则不匹配
  // START_WITH：开头是
  // END_WITH：结尾是
  // WILDCARD：通配符匹配
  // NOT_WILDCARD：通配符不匹配'}
  MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
  // {'en':'Path.', 'zh_CN':'路径。'}
  ListSharedRateLimitingRulesPaths []*string `json:"paths,omitempty" xml:"paths,omitempty" require:"true" type:"Repeated"`
}

func (s ListSharedRateLimitingRulesPathCondition) String() string {
  return tea.Prettify(s)
}

func (s ListSharedRateLimitingRulesPathCondition) GoString() string {
  return s.String()
}

func (s *ListSharedRateLimitingRulesPathCondition) SetMatchType(v string) *ListSharedRateLimitingRulesPathCondition {
  s.MatchType = &v
  return s
}

func (s *ListSharedRateLimitingRulesPathCondition) SetListSharedRateLimitingRulesPaths(v []*string) *ListSharedRateLimitingRulesPathCondition {
  s.ListSharedRateLimitingRulesPaths = v
  return s
}

type ListSharedRateLimitingRulesUaCondition struct {
  // {'en':'Match type.
  // EQUAL:Equals
  // NOT_EQUAL:Does not equal
  // CONTAIN:Contains
  // NOT_CONTAIN:Does not contains
  // REGEX:Regex match
  // NONE:Empty or non-existent
  // NOT_REGEX: regular does not match
  // START_WITH: starts with
  // END_WITH: ends with
  // WILDCARD: wildcard matches
  // NOT_WILDCARD: wildcard does not match', 'zh_CN':'匹配类型。
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
  // NOT_WILDCARD：通配符不匹配'}
  MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
  // {'en':'User agent.', 'zh_CN':'User-Agent。'}
  Ua []*string `json:"ua,omitempty" xml:"ua,omitempty" require:"true" type:"Repeated"`
}

func (s ListSharedRateLimitingRulesUaCondition) String() string {
  return tea.Prettify(s)
}

func (s ListSharedRateLimitingRulesUaCondition) GoString() string {
  return s.String()
}

func (s *ListSharedRateLimitingRulesUaCondition) SetMatchType(v string) *ListSharedRateLimitingRulesUaCondition {
  s.MatchType = &v
  return s
}

func (s *ListSharedRateLimitingRulesUaCondition) SetUa(v []*string) *ListSharedRateLimitingRulesUaCondition {
  s.Ua = v
  return s
}

type ListSharedRateLimitingRulesShareRateLimitRuleCondition struct {
  // {'en':'IP/CIDR.', 'zh_CN':'IP/IP段。'}
  IpOrIpsConditions []*ListSharedRateLimitingRulesIpOrIpsCondition `json:"ipOrIpsConditions,omitempty" xml:"ipOrIpsConditions,omitempty" require:"true" type:"Repeated"`
  // {'en':'Path.', 'zh_CN':'路径。'}
  PathConditions []*ListSharedRateLimitingRulesPathCondition `json:"pathConditions,omitempty" xml:"pathConditions,omitempty" require:"true" type:"Repeated"`
  // {'en':'URI.', 'zh_CN':'URI。'}
  UriConditions []*ListSharedRateLimitingRulesUriCondition `json:"uriConditions,omitempty" xml:"uriConditions,omitempty" require:"true" type:"Repeated"`
  // {'en':'User Agent.', 'zh_CN':'User-Agent。'}
  UaConditions []*ListSharedRateLimitingRulesUaCondition `json:"uaConditions,omitempty" xml:"uaConditions,omitempty" require:"true" type:"Repeated"`
  // {'en':'Request Method.', 'zh_CN':'请求方法。'}
  MethodConditions []*ListSharedRateLimitingRulesRequestMethodCondition `json:"methodConditions,omitempty" xml:"methodConditions,omitempty" require:"true" type:"Repeated"`
  // {'en':'Referer.', 'zh_CN':'Referer。'}
  RefererConditions []*ListSharedRateLimitingRulesRefererCondition `json:"refererConditions,omitempty" xml:"refererConditions,omitempty" require:"true" type:"Repeated"`
  // {'en':'Request Header.', 'zh_CN':'请求头。'}
  HeaderConditions []*ListSharedRateLimitingRulesHeaderCondition `json:"headerConditions,omitempty" xml:"headerConditions,omitempty" require:"true" type:"Repeated"`
  // {'en':'Geo.', 'zh_CN':'区域。'}
  AreaConditions []*ListSharedRateLimitingRulesAreaCondition `json:"areaConditions,omitempty" xml:"areaConditions,omitempty" require:"true" type:"Repeated"`
  // {'en':'Response Code.', 'zh_CN':'状态码。'}
  StatusCodeConditions []*ListSharedRateLimitingRulesStatusCodeCondition `json:"statusCodeConditions,omitempty" xml:"statusCodeConditions,omitempty" require:"true" type:"Repeated"`
  // {'en':'HTTP/S.', 'zh_CN':'应用层协议。'}
  SchemeConditions []*ListSharedRateLimitingRulesSchemeCondition `json:"schemeConditions,omitempty" xml:"schemeConditions,omitempty" require:"true" type:"Repeated"`
}

func (s ListSharedRateLimitingRulesShareRateLimitRuleCondition) String() string {
  return tea.Prettify(s)
}

func (s ListSharedRateLimitingRulesShareRateLimitRuleCondition) GoString() string {
  return s.String()
}

func (s *ListSharedRateLimitingRulesShareRateLimitRuleCondition) SetIpOrIpsConditions(v []*ListSharedRateLimitingRulesIpOrIpsCondition) *ListSharedRateLimitingRulesShareRateLimitRuleCondition {
  s.IpOrIpsConditions = v
  return s
}

func (s *ListSharedRateLimitingRulesShareRateLimitRuleCondition) SetPathConditions(v []*ListSharedRateLimitingRulesPathCondition) *ListSharedRateLimitingRulesShareRateLimitRuleCondition {
  s.PathConditions = v
  return s
}

func (s *ListSharedRateLimitingRulesShareRateLimitRuleCondition) SetUriConditions(v []*ListSharedRateLimitingRulesUriCondition) *ListSharedRateLimitingRulesShareRateLimitRuleCondition {
  s.UriConditions = v
  return s
}

func (s *ListSharedRateLimitingRulesShareRateLimitRuleCondition) SetUaConditions(v []*ListSharedRateLimitingRulesUaCondition) *ListSharedRateLimitingRulesShareRateLimitRuleCondition {
  s.UaConditions = v
  return s
}

func (s *ListSharedRateLimitingRulesShareRateLimitRuleCondition) SetMethodConditions(v []*ListSharedRateLimitingRulesRequestMethodCondition) *ListSharedRateLimitingRulesShareRateLimitRuleCondition {
  s.MethodConditions = v
  return s
}

func (s *ListSharedRateLimitingRulesShareRateLimitRuleCondition) SetRefererConditions(v []*ListSharedRateLimitingRulesRefererCondition) *ListSharedRateLimitingRulesShareRateLimitRuleCondition {
  s.RefererConditions = v
  return s
}

func (s *ListSharedRateLimitingRulesShareRateLimitRuleCondition) SetHeaderConditions(v []*ListSharedRateLimitingRulesHeaderCondition) *ListSharedRateLimitingRulesShareRateLimitRuleCondition {
  s.HeaderConditions = v
  return s
}

func (s *ListSharedRateLimitingRulesShareRateLimitRuleCondition) SetAreaConditions(v []*ListSharedRateLimitingRulesAreaCondition) *ListSharedRateLimitingRulesShareRateLimitRuleCondition {
  s.AreaConditions = v
  return s
}

func (s *ListSharedRateLimitingRulesShareRateLimitRuleCondition) SetStatusCodeConditions(v []*ListSharedRateLimitingRulesStatusCodeCondition) *ListSharedRateLimitingRulesShareRateLimitRuleCondition {
  s.StatusCodeConditions = v
  return s
}

func (s *ListSharedRateLimitingRulesShareRateLimitRuleCondition) SetSchemeConditions(v []*ListSharedRateLimitingRulesSchemeCondition) *ListSharedRateLimitingRulesShareRateLimitRuleCondition {
  s.SchemeConditions = v
  return s
}

type ListSharedRateLimitingRulesUriCondition struct {
  // {'en':'Match type.
  // EQUAL:Equals
  // NOT_EQUAL:Does not equal
  // CONTAIN:Contains
  // NOT_CONTAIN:Does not contains
  // REGEX:Regex match
  // NOT_REGEX: regular does not match
  // START_WITH: starts with
  // END_WITH: ends with
  // WILDCARD: wildcard matches
  // NOT_WILDCARD: wildcard does not match', 'zh_CN':'匹配类型。
  // EQUAL：等于
  // NOT_EQUAL：不等于
  // CONTAIN：包含
  // NOT_CONTAIN：不包含
  // REGEX：正则
  // NOT_REGEX：正则不匹配
  // START_WITH：开头是
  // END_WITH：结尾是
  // WILDCARD：通配符匹配
  // NOT_WILDCARD：通配符不匹配'}
  MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
  // {'en':'URI.', 'zh_CN':'URI。'}
  Uri []*string `json:"uri,omitempty" xml:"uri,omitempty" require:"true" type:"Repeated"`
}

func (s ListSharedRateLimitingRulesUriCondition) String() string {
  return tea.Prettify(s)
}

func (s ListSharedRateLimitingRulesUriCondition) GoString() string {
  return s.String()
}

func (s *ListSharedRateLimitingRulesUriCondition) SetMatchType(v string) *ListSharedRateLimitingRulesUriCondition {
  s.MatchType = &v
  return s
}

func (s *ListSharedRateLimitingRulesUriCondition) SetUri(v []*string) *ListSharedRateLimitingRulesUriCondition {
  s.Uri = v
  return s
}

type ListSharedRateLimitingRulesStatusCodeCondition struct {
  // {'en':'Match type.
  // EQUAL:Equal
  // NOT_EQUAL:Does not equal', 'zh_CN':'匹配类型。
  // EQUAL：等于
  // NOT_EQUAL：不等于'}
  MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
  // {'en':'Response Code.', 'zh_CN':'状态码。'}
  StatusCode []*string `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true" type:"Repeated"`
}

func (s ListSharedRateLimitingRulesStatusCodeCondition) String() string {
  return tea.Prettify(s)
}

func (s ListSharedRateLimitingRulesStatusCodeCondition) GoString() string {
  return s.String()
}

func (s *ListSharedRateLimitingRulesStatusCodeCondition) SetMatchType(v string) *ListSharedRateLimitingRulesStatusCodeCondition {
  s.MatchType = &v
  return s
}

func (s *ListSharedRateLimitingRulesStatusCodeCondition) SetStatusCode(v []*string) *ListSharedRateLimitingRulesStatusCodeCondition {
  s.StatusCode = v
  return s
}

type ListSharedRateLimitingRulesSchemeCondition struct {
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

func (s ListSharedRateLimitingRulesSchemeCondition) String() string {
  return tea.Prettify(s)
}

func (s ListSharedRateLimitingRulesSchemeCondition) GoString() string {
  return s.String()
}

func (s *ListSharedRateLimitingRulesSchemeCondition) SetMatchType(v string) *ListSharedRateLimitingRulesSchemeCondition {
  s.MatchType = &v
  return s
}

func (s *ListSharedRateLimitingRulesSchemeCondition) SetScheme(v []*string) *ListSharedRateLimitingRulesSchemeCondition {
  s.Scheme = v
  return s
}

type ListSharedRateLimitingRulesHeaderCondition struct {
  // {'en':'Match type.
  // EQUAL:Equals
  // NOT_EQUAL:Does not equal
  // CONTAIN:Contains
  // NOT_CONTAIN:Does not contains
  // REGEX:Regex match
  // NONE:Empty or non-existent
  // NOT_REGEX: regular does not match
  // START_WITH: starts with
  // END_WITH: ends with
  // WILDCARD: wildcard matches
  // NOT_WILDCARD: wildcard does not match', 'zh_CN':'匹配类型。
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
  // NOT_WILDCARD：通配符不匹配'}
  MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
  // {'en':'Header name.', 'zh_CN':'头部名称。'}
  Key *string `json:"key,omitempty" xml:"key,omitempty" require:"true"`
  // {'en':'Header value.', 'zh_CN':'头部值。'}
  ValueList []*string `json:"valueList,omitempty" xml:"valueList,omitempty" require:"true" type:"Repeated"`
}

func (s ListSharedRateLimitingRulesHeaderCondition) String() string {
  return tea.Prettify(s)
}

func (s ListSharedRateLimitingRulesHeaderCondition) GoString() string {
  return s.String()
}

func (s *ListSharedRateLimitingRulesHeaderCondition) SetMatchType(v string) *ListSharedRateLimitingRulesHeaderCondition {
  s.MatchType = &v
  return s
}

func (s *ListSharedRateLimitingRulesHeaderCondition) SetKey(v string) *ListSharedRateLimitingRulesHeaderCondition {
  s.Key = &v
  return s
}

func (s *ListSharedRateLimitingRulesHeaderCondition) SetValueList(v []*string) *ListSharedRateLimitingRulesHeaderCondition {
  s.ValueList = v
  return s
}

type ListSharedRateLimitingRulesRateLimitEffective struct {
  // {'en':'effective.
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

func (s ListSharedRateLimitingRulesRateLimitEffective) String() string {
  return tea.Prettify(s)
}

func (s ListSharedRateLimitingRulesRateLimitEffective) GoString() string {
  return s.String()
}

func (s *ListSharedRateLimitingRulesRateLimitEffective) SetEffective(v []*string) *ListSharedRateLimitingRulesRateLimitEffective {
  s.Effective = v
  return s
}

func (s *ListSharedRateLimitingRulesRateLimitEffective) SetStart(v string) *ListSharedRateLimitingRulesRateLimitEffective {
  s.Start = &v
  return s
}

func (s *ListSharedRateLimitingRulesRateLimitEffective) SetEnd(v string) *ListSharedRateLimitingRulesRateLimitEffective {
  s.End = &v
  return s
}

func (s *ListSharedRateLimitingRulesRateLimitEffective) SetTimezone(v string) *ListSharedRateLimitingRulesRateLimitEffective {
  s.Timezone = &v
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

type ListSharedRateLimitingRulesRequestHeader struct {
  // {"zh_CN":"安全服务类型。有使用多个不同的安全服务时，需要填写具体的服务类型。","en":"Security service type. Please enter a specific service type, if you purchase multiple security services.","dictionary":"belong=WAAP-MS-Ext|dict=waap_serviceType"}
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

type ListSharedRateLimitingRulesResponseHeader struct {
}

func (s ListSharedRateLimitingRulesResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ListSharedRateLimitingRulesResponseHeader) GoString() string {
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

type ListSharedWAFRuleExceptionsResponse struct {
  // {'en':'Please refer to the error code for exceptions.', 'zh_CN':'请参照错误码。','dictionary':'belong=WAAP-MS-Ext|dict=waap_retCodeEnum'}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {'en':'Description.', 'zh_CN':'描述信息。'}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {'en':'Shared configuration WAF rule exception list.', 'zh_CN':'共享配置WAF规则例外列表。'}
  Data []*ListSharedWAFRuleExceptionsWafUserShareExceptionListVO `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
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

func (s *ListSharedWAFRuleExceptionsResponse) SetData(v []*ListSharedWAFRuleExceptionsWafUserShareExceptionListVO) *ListSharedWAFRuleExceptionsResponse {
  s.Data = v
  return s
}

type ListSharedWAFRuleExceptionsWafUserShareExceptionListVO struct {
  // {'en':'Rule exception ID.', 'zh_CN':'规则例外ID。'}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {'en':'Exception name,maximum 50 characters.
  // Does not support special characters and spaces.', 'zh_CN':'例外名称，最多50个字符。
  // 不支持特殊字符和空格。'}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {'en':'Exception description,maximum 200 characters..', 'zh_CN':'例外描述，最多200个字符。'}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {'en':'Matching conditions.
  // ip: IP
  // path: Path
  // uri: URI
  // urlParamName: URI Parameter Name
  // urlParamValue: URI Parameter Value
  // userAgent: User Agent
  // httpHeaderName: Request Header Name
  // httpHeaderValue: Request Header Value
  // cookie: Cookie
  // body: Body
  // bodyParamName: Body Parameter Name
  // bodyParamValue: Body Parameter Value', 'zh_CN':'匹配条件。
  // ip：IP
  // path：路径
  // uri：URI
  // urlParamName：URI参数名
  // urlParamValue：URI参数值
  // userAgent：User Agent
  // httpHeaderName：请求头部名称
  // httpHeaderValue：请求头部值
  // cookie：Cookie
  // body：Body
  // bodyParamName：Body参数名
  // bodyParamValue：Body参数值'}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
  // {'en':'Match type,IP can only be EQUAL.
  // EQUAL: Equal
  // CONTAIN: Contains
  // REGEX: Regular match', 'zh_CN':'匹配类型，IP只能是等于。
  // EQUAL：等于
  // CONTAIN：包含
  // REGEX：正则匹配'}
  MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
  // {'en':'Rule exceptions.
  // When matchType=EQUAL, case-sensitive.', 'zh_CN':'规则例外内容。
  // matchType=EQUAL时，大小写敏感。
  // '}
  ContentList []*string `json:"contentList,omitempty" xml:"contentList,omitempty" require:"true" type:"Repeated"`
  // {'en':'Creator login.', 'zh_CN':'创建者登录名。'}
  Creator *string `json:"creator,omitempty" xml:"creator,omitempty" require:"true"`
  // {'en':'Creation time.', 'zh_CN':'创建时间。'}
  CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty" require:"true"`
  // {'en':'Update time.', 'zh_CN':'更新时间。'}
  UpdateTime *string `json:"updateTime,omitempty" xml:"updateTime,omitempty" require:"true"`
  // {'en':'List of associated hostnames and rule IDs.', 'zh_CN':'关联的域名以及规则ID列表。'}
  DomainRuleList []*ListSharedWAFRuleExceptionsWafShareExceptionDomainRule `json:"domainRuleList,omitempty" xml:"domainRuleList,omitempty" require:"true" type:"Repeated"`
}

func (s ListSharedWAFRuleExceptionsWafUserShareExceptionListVO) String() string {
  return tea.Prettify(s)
}

func (s ListSharedWAFRuleExceptionsWafUserShareExceptionListVO) GoString() string {
  return s.String()
}

func (s *ListSharedWAFRuleExceptionsWafUserShareExceptionListVO) SetId(v string) *ListSharedWAFRuleExceptionsWafUserShareExceptionListVO {
  s.Id = &v
  return s
}

func (s *ListSharedWAFRuleExceptionsWafUserShareExceptionListVO) SetName(v string) *ListSharedWAFRuleExceptionsWafUserShareExceptionListVO {
  s.Name = &v
  return s
}

func (s *ListSharedWAFRuleExceptionsWafUserShareExceptionListVO) SetMsg(v string) *ListSharedWAFRuleExceptionsWafUserShareExceptionListVO {
  s.Msg = &v
  return s
}

func (s *ListSharedWAFRuleExceptionsWafUserShareExceptionListVO) SetType(v string) *ListSharedWAFRuleExceptionsWafUserShareExceptionListVO {
  s.Type = &v
  return s
}

func (s *ListSharedWAFRuleExceptionsWafUserShareExceptionListVO) SetMatchType(v string) *ListSharedWAFRuleExceptionsWafUserShareExceptionListVO {
  s.MatchType = &v
  return s
}

func (s *ListSharedWAFRuleExceptionsWafUserShareExceptionListVO) SetContentList(v []*string) *ListSharedWAFRuleExceptionsWafUserShareExceptionListVO {
  s.ContentList = v
  return s
}

func (s *ListSharedWAFRuleExceptionsWafUserShareExceptionListVO) SetCreator(v string) *ListSharedWAFRuleExceptionsWafUserShareExceptionListVO {
  s.Creator = &v
  return s
}

func (s *ListSharedWAFRuleExceptionsWafUserShareExceptionListVO) SetCreateTime(v string) *ListSharedWAFRuleExceptionsWafUserShareExceptionListVO {
  s.CreateTime = &v
  return s
}

func (s *ListSharedWAFRuleExceptionsWafUserShareExceptionListVO) SetUpdateTime(v string) *ListSharedWAFRuleExceptionsWafUserShareExceptionListVO {
  s.UpdateTime = &v
  return s
}

func (s *ListSharedWAFRuleExceptionsWafUserShareExceptionListVO) SetDomainRuleList(v []*ListSharedWAFRuleExceptionsWafShareExceptionDomainRule) *ListSharedWAFRuleExceptionsWafUserShareExceptionListVO {
  s.DomainRuleList = v
  return s
}

type ListSharedWAFRuleExceptionsWafShareExceptionDomainRule struct {
  // {'en':'Hostname.', 'zh_CN':'域名。'}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {'en':'WAF rule ID list.', 'zh_CN':'WAF规则ID列表。'}
  RuleIdList []*int32 `json:"ruleIdList,omitempty" xml:"ruleIdList,omitempty" require:"true" type:"Repeated"`
}

func (s ListSharedWAFRuleExceptionsWafShareExceptionDomainRule) String() string {
  return tea.Prettify(s)
}

func (s ListSharedWAFRuleExceptionsWafShareExceptionDomainRule) GoString() string {
  return s.String()
}

func (s *ListSharedWAFRuleExceptionsWafShareExceptionDomainRule) SetDomain(v string) *ListSharedWAFRuleExceptionsWafShareExceptionDomainRule {
  s.Domain = &v
  return s
}

func (s *ListSharedWAFRuleExceptionsWafShareExceptionDomainRule) SetRuleIdList(v []*int32) *ListSharedWAFRuleExceptionsWafShareExceptionDomainRule {
  s.RuleIdList = v
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

type ListSharedWAFRuleExceptionsRequestHeader struct {
  // {"zh_CN":"安全服务类型。有使用多个不同的安全服务时，需要填写具体的服务类型。","en":"Security service type. Please enter a specific service type, if you purchase multiple security services.","dictionary":"belong=WAAP-MS-Ext|dict=waap_serviceType"}
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

type ListSharedWAFRuleExceptionsResponseHeader struct {
}

func (s ListSharedWAFRuleExceptionsResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ListSharedWAFRuleExceptionsResponseHeader) GoString() string {
  return s.String()
}




type QueryAppApiExceptionFeatureReferencedHostnamesRequest struct {
  // {'en':'Rule Id, Only supports single rule queries', 'zh_CN':'规则ID，仅支持单个规则查询'}
  IdList []*string `json:"idList,omitempty" xml:"idList,omitempty" require:"true" type:"Repeated"`
}

func (s QueryAppApiExceptionFeatureReferencedHostnamesRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryAppApiExceptionFeatureReferencedHostnamesRequest) GoString() string {
  return s.String()
}

func (s *QueryAppApiExceptionFeatureReferencedHostnamesRequest) SetIdList(v []*string) *QueryAppApiExceptionFeatureReferencedHostnamesRequest {
  s.IdList = v
  return s
}

type QueryAppApiExceptionFeatureReferencedHostnamesResponse struct {
  // {'en':'Please refer to the error code for exceptions.', 'zh_CN':'请参照错误码。','dictionary':'belong=WAAP-MS-Ext|dict=waap_retCodeEnum'}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {'en':'Description.', 'zh_CN':'描述信息。'}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {'en':'Data.', 'zh_CN':'出参数据。'}
  Data []*string `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s QueryAppApiExceptionFeatureReferencedHostnamesResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryAppApiExceptionFeatureReferencedHostnamesResponse) GoString() string {
  return s.String()
}

func (s *QueryAppApiExceptionFeatureReferencedHostnamesResponse) SetCode(v string) *QueryAppApiExceptionFeatureReferencedHostnamesResponse {
  s.Code = &v
  return s
}

func (s *QueryAppApiExceptionFeatureReferencedHostnamesResponse) SetMsg(v string) *QueryAppApiExceptionFeatureReferencedHostnamesResponse {
  s.Msg = &v
  return s
}

func (s *QueryAppApiExceptionFeatureReferencedHostnamesResponse) SetData(v []*string) *QueryAppApiExceptionFeatureReferencedHostnamesResponse {
  s.Data = v
  return s
}

type QueryAppApiExceptionFeatureReferencedHostnamesPaths struct {
}

func (s QueryAppApiExceptionFeatureReferencedHostnamesPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryAppApiExceptionFeatureReferencedHostnamesPaths) GoString() string {
  return s.String()
}

type QueryAppApiExceptionFeatureReferencedHostnamesParameters struct {
}

func (s QueryAppApiExceptionFeatureReferencedHostnamesParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryAppApiExceptionFeatureReferencedHostnamesParameters) GoString() string {
  return s.String()
}

type QueryAppApiExceptionFeatureReferencedHostnamesRequestHeader struct {
  // {"zh_CN":"安全服务类型。有使用多个不同的安全服务时，需要填写具体的服务类型。","en":"Security service type. Please enter a specific service type, if you purchase multiple security services.","dictionary":"belong=WAAP-MS-Ext|dict=waap_serviceType"}
  ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty"`
}

func (s QueryAppApiExceptionFeatureReferencedHostnamesRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryAppApiExceptionFeatureReferencedHostnamesRequestHeader) GoString() string {
  return s.String()
}

func (s *QueryAppApiExceptionFeatureReferencedHostnamesRequestHeader) SetServiceType(v string) *QueryAppApiExceptionFeatureReferencedHostnamesRequestHeader {
  s.ServiceType = &v
  return s
}

type QueryAppApiExceptionFeatureReferencedHostnamesResponseHeader struct {
}

func (s QueryAppApiExceptionFeatureReferencedHostnamesResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryAppApiExceptionFeatureReferencedHostnamesResponseHeader) GoString() string {
  return s.String()
}




type DeleteSharedWhitelistRuleRequest struct {
  // {"en":"Rule ID List.", "zh_CN":"规则ID列表。"}
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

type DeleteSharedWhitelistRuleResponse struct {
  // {'en':'Please refer to the error code for exceptions.', 'zh_CN':'请参照错误码。','dictionary':'belong=WAAP-MS-Ext|dict=waap_retCodeEnum'}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Description.", "zh_CN":"描述信息。"}
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

type DeleteSharedWhitelistRuleRequestHeader struct {
  // {"zh_CN":"安全服务类型。有使用多个不同的安全服务时，需要填写具体的服务类型。","en":"Security service type. Please enter a specific service type, if you purchase multiple security services.","dictionary":"belong=WAAP-MS-Ext|dict=waap_serviceType"}
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

type DeleteSharedWhitelistRuleResponseHeader struct {
}

func (s DeleteSharedWhitelistRuleResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteSharedWhitelistRuleResponseHeader) GoString() string {
  return s.String()
}




type UpdateShareConfigurationsAppApiExceptionFeatureRequest struct {
  // {'en':'Rule ID', 'zh_CN':'规则ID'}
  RuleId *string `json:"ruleId,omitempty" xml:"ruleId,omitempty" require:"true"`
  // {'en':'Rule name, duplicate not allowed.', 'zh_CN':'规则名称，不允许重复'}
  RuleName *string `json:"ruleName,omitempty" xml:"ruleName,omitempty" require:"true"`
  // {
  // 'en':'Type enumeration, for example:
  // NATIVE_ APP        remarks: Native App
  // Callback_ API      	 remarks: Callback API
  // Other_ API       	  remarks: Other program API
  // HYBRID_ APP  	   remarks: Hybrid APP',
  // 'zh_CN':'类型枚举，例如：
  // NATIVE_APP          备注：原生APP
  // CALLBACK_API      备注：回调API
  // OTHER_API            备注：其他程序API
  // HYBRID_APP          备注：混合APP'}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
  // {'en':'Characteristic.', 'zh_CN':'特征。'}
  Config *UpdateShareConfigurationsAppApiExceptionFeatureDmsShareServiceFeatureUpdateDTO `json:"config,omitempty" xml:"config,omitempty" require:"true"`
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

func (s *UpdateShareConfigurationsAppApiExceptionFeatureRequest) SetConfig(v *UpdateShareConfigurationsAppApiExceptionFeatureDmsShareServiceFeatureUpdateDTO) *UpdateShareConfigurationsAppApiExceptionFeatureRequest {
  s.Config = v
  return s
}

type UpdateShareConfigurationsAppApiExceptionFeatureDmsShareServiceFeatureUpdateDTO struct {
  // {'en':'Matching condition name.
  // PATH         Path
  // URI            URI
  // UA             User-Agent
  // REFERER    Referer
  // HEADER    Request Header
  // ', 'zh_CN':'匹配条件名称。
  // PATH                       路径
  // URI                          URI
  // UA                           User-Agent
  // REFERER                  Referer
  // HEAD                      Request Header
  // '}
  Condition *string `json:"condition,omitempty" xml:"condition,omitempty" require:"true"`
  // {'en':'Matching condition function.
  // CONTAIN                  contains
  // NOT_CONTAIN         does not contain
  // EQUAL                      equals
  // NOT_EQUAL             does not equal
  // EMPTY                      does not exist or has no value(When the condition is a path or URI, this value is not optional)
  // REGEX                      regex match
  // ', 'zh_CN':'匹配条件函数。
  // CONTAIN                  包含
  // NOT_CONTAIN         不包含 
  // EQUAL                      等于
  // NOT_EQUAL             不等于
  // EMPTY                      为空或不存在（当condition为路径或者URI时，该值不可选）
  // REGEX                       匹配正则 
  // '}
  MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
  // {'en':'Condition value list, When matchType is EMPTY, the value is empty;  When matchType is REGEX, only one item is set', 'zh_CN':'条件值列表，当matchType为EMPTY时，该值为空；当matchType为REGEX时，只能传一条'}
  ConditionValue []*string `json:"conditionValue,omitempty" xml:"conditionValue,omitempty" require:"true" type:"Repeated"`
  // {'en':'Head value, If condition=HEAD, then this field is mandatory, Otherwise, it is empty.', 'zh_CN':'头部值，如果condition为HEAD，则该字段必填，否则为空。'}
  Key *string `json:"key,omitempty" xml:"key,omitempty"`
}

func (s UpdateShareConfigurationsAppApiExceptionFeatureDmsShareServiceFeatureUpdateDTO) String() string {
  return tea.Prettify(s)
}

func (s UpdateShareConfigurationsAppApiExceptionFeatureDmsShareServiceFeatureUpdateDTO) GoString() string {
  return s.String()
}

func (s *UpdateShareConfigurationsAppApiExceptionFeatureDmsShareServiceFeatureUpdateDTO) SetCondition(v string) *UpdateShareConfigurationsAppApiExceptionFeatureDmsShareServiceFeatureUpdateDTO {
  s.Condition = &v
  return s
}

func (s *UpdateShareConfigurationsAppApiExceptionFeatureDmsShareServiceFeatureUpdateDTO) SetMatchType(v string) *UpdateShareConfigurationsAppApiExceptionFeatureDmsShareServiceFeatureUpdateDTO {
  s.MatchType = &v
  return s
}

func (s *UpdateShareConfigurationsAppApiExceptionFeatureDmsShareServiceFeatureUpdateDTO) SetConditionValue(v []*string) *UpdateShareConfigurationsAppApiExceptionFeatureDmsShareServiceFeatureUpdateDTO {
  s.ConditionValue = v
  return s
}

func (s *UpdateShareConfigurationsAppApiExceptionFeatureDmsShareServiceFeatureUpdateDTO) SetKey(v string) *UpdateShareConfigurationsAppApiExceptionFeatureDmsShareServiceFeatureUpdateDTO {
  s.Key = &v
  return s
}

type UpdateShareConfigurationsAppApiExceptionFeatureResponse struct {
  // {'en':'Please refer to the error code for exceptions.', 'zh_CN':'请参照错误码。','dictionary':'belong=WAAP-MS-Ext|dict=waap_retCodeEnum'}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {'en':'Description.', 'zh_CN':'描述信息。'}
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

type UpdateShareConfigurationsAppApiExceptionFeatureRequestHeader struct {
  // {"zh_CN":"安全服务类型。有使用多个不同的安全服务时，需要填写具体的服务类型。","en":"Security service type. Please enter a specific service type, if you purchase multiple security services.","dictionary":"belong=WAAP-MS-Ext|dict=waap_serviceType"}
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

type UpdateShareConfigurationsAppApiExceptionFeatureResponseHeader struct {
}

func (s UpdateShareConfigurationsAppApiExceptionFeatureResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s UpdateShareConfigurationsAppApiExceptionFeatureResponseHeader) GoString() string {
  return s.String()
}




type DeleteSharedRateLimitingRuleRequest struct {
  // {'en':'Rule ID List.', 'zh_CN':'规则ID列表。'}
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

type DeleteSharedRateLimitingRuleResponse struct {
  // {'en':'Please refer to the error code for exceptions.', 'zh_CN':'请参照错误码。','dictionary':'belong=WAAP-MS-Ext|dict=waap_retCodeEnum'}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {'en':'Description.', 'zh_CN':'描述信息。'}
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

type DeleteSharedRateLimitingRuleRequestHeader struct {
  // {"zh_CN":"安全服务类型。有使用多个不同的安全服务时，需要填写具体的服务类型。","en":"Security service type. Please enter a specific service type, if you purchase multiple security services.","dictionary":"belong=WAAP-MS-Ext|dict=waap_serviceType"}
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

type DeleteSharedRateLimitingRuleResponseHeader struct {
}

func (s DeleteSharedRateLimitingRuleResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteSharedRateLimitingRuleResponseHeader) GoString() string {
  return s.String()
}




type AssociateShareCustomizeBotsRequest struct {
  // {"zh_CN":"共享配置ID。","en":"Share configuration ID."}
  ShareId *string `json:"shareId,omitempty" xml:"shareId,omitempty" require:"true"`
  // {"zh_CN":"待关联域名列表。","en":"List of hostname to be associated."}
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

type AssociateShareCustomizeBotsResponse struct {
  // {'en':'Please refer to the error code for exceptions.', 'zh_CN':'请参照错误码。','dictionary':'belong=WAAP-MS-Ext|dict=waap_retCodeEnum'}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {'en':'Description.', 'zh_CN':'描述信息。'}
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

type AssociateShareCustomizeBotsRequestHeader struct {
  // {"zh_CN":"安全服务类型。有使用多个不同的安全服务时，需要填写具体的服务类型。","en":"Security service type. Please enter a specific service type, if you purchase multiple security services.","dictionary":"belong=WAAP-MS-Ext|dict=waap_serviceType"}
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

type AssociateShareCustomizeBotsResponseHeader struct {
}

func (s AssociateShareCustomizeBotsResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s AssociateShareCustomizeBotsResponseHeader) GoString() string {
  return s.String()
}




type DeleteSharedWAFRuleExceptionRequest struct {
  // {'en':'Rule exception ID list.', 'zh_CN':'规则例外ID列表。'}
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

type DeleteSharedWAFRuleExceptionResponse struct {
  // {'en':'Please refer to the error code for exceptions.', 'zh_CN':'请参照错误码。','dictionary':'belong=WAAP-MS-Ext|dict=waap_retCodeEnum'}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {'en':'Description.', 'zh_CN':'描述信息。'}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {'en':'Delete success indicator.', 'zh_CN':'删除成功标识。'}
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

type DeleteSharedWAFRuleExceptionRequestHeader struct {
  // {"zh_CN":"安全服务类型。有使用多个不同的安全服务时，需要填写具体的服务类型。","en":"Security service type. Please enter a specific service type, if you purchase multiple security services.","dictionary":"belong=WAAP-MS-Ext|dict=waap_serviceType"}
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

type DeleteSharedWAFRuleExceptionResponseHeader struct {
}

func (s DeleteSharedWAFRuleExceptionResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteSharedWAFRuleExceptionResponseHeader) GoString() string {
  return s.String()
}




type CreateAppApiExceptionFeatureRequest struct {
  // {'en':'Rule name, duplicate not allowed.', 'zh_CN':'规则名称，不允许重复'}
  RuleName *string `json:"ruleName,omitempty" xml:"ruleName,omitempty" require:"true"`
  // {
  // 'en':'Type enumeration, for example:
  // NATIVE_ APP        remarks: Native App
  // Callback_ API      	 remarks: Callback API
  // Other_ API       	  remarks: Other program API
  // HYBRID_ APP  	   remarks: Hybrid APP',
  // 'zh_CN':'类型枚举，例如：
  // NATIVE_APP          备注：原生APP
  // CALLBACK_API      备注：回调API
  // OTHER_API            备注：其他程序API
  // HYBRID_APP          备注：混合APP'}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
  // {'en':'Characteristic.', 'zh_CN':'特征。'}
  Config *CreateAppApiExceptionFeatureDmsShareServiceFeatureAddDTO `json:"config,omitempty" xml:"config,omitempty" require:"true"`
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

func (s *CreateAppApiExceptionFeatureRequest) SetConfig(v *CreateAppApiExceptionFeatureDmsShareServiceFeatureAddDTO) *CreateAppApiExceptionFeatureRequest {
  s.Config = v
  return s
}

type CreateAppApiExceptionFeatureDmsShareServiceFeatureAddDTO struct {
  // {'en':'Matching condition name.
  // PATH         Path
  // URI            URI
  // UA             User-Agent
  // REFERER    Referer
  // HEADER    Request Header
  // ', 'zh_CN':'匹配条件名称。
  // PATH                       路径
  // URI                          URI
  // UA                           User-Agent
  // REFERER                  Referer
  // HEAD                      Request Header
  // '}
  Condition *string `json:"condition,omitempty" xml:"condition,omitempty" require:"true"`
  // {'en':'Matching condition function.
  // CONTAIN                  contains
  // NOT_CONTAIN         does not contain
  // EQUAL                      equals
  // NOT_EQUAL             does not equal
  // EMPTY                      does not exist or has no value(When the condition is a path or URI, this value is not optional)
  // REGEX                      regex match
  // ', 'zh_CN':'匹配条件函数。
  // CONTAIN                  包含
  // NOT_CONTAIN         不包含 
  // EQUAL                      等于
  // NOT_EQUAL             不等于
  // EMPTY                      为空或不存在（当condition为路径或者URI时，该值不可选）
  // REGEX                       匹配正则 
  // '}
  MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
  // {'en':'Condition value list, When matchType is EMPTY, the value is empty;  When matchType is REGEX, only one item is set', 'zh_CN':'条件值列表，当matchType为EMPTY时，该值为空；当matchType为REGEX时，只能传一条'}
  ConditionValue []*string `json:"conditionValue,omitempty" xml:"conditionValue,omitempty" require:"true" type:"Repeated"`
  // {'en':'Head value, If condition=HEAD, then this field is mandatory, Otherwise, it is empty.', 'zh_CN':'头部值，如果condition为HEAD，则该字段必填，否则为空。'}
  Key *string `json:"key,omitempty" xml:"key,omitempty"`
}

func (s CreateAppApiExceptionFeatureDmsShareServiceFeatureAddDTO) String() string {
  return tea.Prettify(s)
}

func (s CreateAppApiExceptionFeatureDmsShareServiceFeatureAddDTO) GoString() string {
  return s.String()
}

func (s *CreateAppApiExceptionFeatureDmsShareServiceFeatureAddDTO) SetCondition(v string) *CreateAppApiExceptionFeatureDmsShareServiceFeatureAddDTO {
  s.Condition = &v
  return s
}

func (s *CreateAppApiExceptionFeatureDmsShareServiceFeatureAddDTO) SetMatchType(v string) *CreateAppApiExceptionFeatureDmsShareServiceFeatureAddDTO {
  s.MatchType = &v
  return s
}

func (s *CreateAppApiExceptionFeatureDmsShareServiceFeatureAddDTO) SetConditionValue(v []*string) *CreateAppApiExceptionFeatureDmsShareServiceFeatureAddDTO {
  s.ConditionValue = v
  return s
}

func (s *CreateAppApiExceptionFeatureDmsShareServiceFeatureAddDTO) SetKey(v string) *CreateAppApiExceptionFeatureDmsShareServiceFeatureAddDTO {
  s.Key = &v
  return s
}

type CreateAppApiExceptionFeatureResponse struct {
  // {'en':'Please refer to the error code for exceptions.', 'zh_CN':'请参照错误码。','dictionary':'belong=WAAP-MS-Ext|dict=waap_retCodeEnum'}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {'en':'Description.', 'zh_CN':'描述信息。'}
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

type CreateAppApiExceptionFeatureRequestHeader struct {
  // {"zh_CN":"安全服务类型。有使用多个不同的安全服务时，需要填写具体的服务类型。","en":"Security service type. Please enter a specific service type, if you purchase multiple security services.","dictionary":"belong=WAAP-MS-Ext|dict=waap_serviceType"}
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

type CreateAppApiExceptionFeatureResponseHeader struct {
}

func (s CreateAppApiExceptionFeatureResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s CreateAppApiExceptionFeatureResponseHeader) GoString() string {
  return s.String()
}




type DisassociateSharedWhitelistRuleRequest struct {
  // {"zh_CN":"共享规则ID。","en":"Share rule ID."}
  ShareId *string `json:"shareId,omitempty" xml:"shareId,omitempty" require:"true"`
  // {"zh_CN":"待取消关联域名列表。","en":"List of hostname to be disassociated."}
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

type DisassociateSharedWhitelistRuleResponse struct {
  // {'en':'Please refer to the error code for exceptions.', 'zh_CN':'请参照错误码。','dictionary':'belong=WAAP-MS-Ext|dict=waap_retCodeEnum'}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {'en':'Description.', 'zh_CN':'描述信息。'}
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

type DisassociateSharedWhitelistRuleRequestHeader struct {
  // {"zh_CN":"安全服务类型。有使用多个不同的安全服务时，需要填写具体的服务类型。","en":"Security service type. Please enter a specific service type, if you purchase multiple security services.","dictionary":"belong=WAAP-MS-Ext|dict=waap_serviceType"}
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

type DisassociateSharedWhitelistRuleResponseHeader struct {
}

func (s DisassociateSharedWhitelistRuleResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s DisassociateSharedWhitelistRuleResponseHeader) GoString() string {
  return s.String()
}




type UpdateSharedRateLimitingRuleRequest struct {
  // {'en':'Rule ID.', 'zh_CN':'规则ID。'}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {'en':'Rule Name, maximum 50 characters.
  // Does not support special characters and spaces.', 'zh_CN':'规则名称，最多50个字符。
  // 不支持特殊字符和空格。'}
  RuleName *string `json:"ruleName,omitempty" xml:"ruleName,omitempty"`
  // {'en':'Description, maximum 200 characters.', 'zh_CN':'规则描述，最多200个字符。'}
  Description *string `json:"description,omitempty" xml:"description,omitempty"`
  // {'en':'Client identifier.
  // IP:Client IP
  // IP_UA:Client IP and User-Agent
  // COOKIE:Cookie
  // IP_COOKIE:Client IP and Cookie
  // HEADER:Request Header
  // When there is a status code in the matching condition,this client identifier is not supported.
  // IP_HEADER:Client IP and Request Header
  // When there is a status code in the matching condition,this client identifier is not supported .', 'zh_CN':'统计粒度。
  // IP：客户端IP
  // IP_UA：客户端IP和User-Agent
  // COOKIE：Cookie
  // IP_COOKIE：客户端IP和Cookie
  // HEADER：请求头，当匹配条件中存在状态码时不支持此统计粒度
  // IP_HEADER：客户端IP和请求头，当匹配条件中存在状态码时不支持此统计粒度'}
  StatisticalItem *string `json:"statisticalItem,omitempty" xml:"statisticalItem,omitempty"`
  // {'en':'Statistical key value.
  // When the client identifier is cookie/header value, the corresponding key value needs to be entered.', 'zh_CN':'统计key值。
  // 当统计粒度cookie/header值，需要输入对应的key值。'}
  StatisticsKey *string `json:"statisticsKey,omitempty" xml:"statisticsKey,omitempty"`
  // {'en':'Statistics period, unit: seconds, the range is 1 - 3600.', 'zh_CN':'统计周期，单位：秒，范围为 1 - 3600。'}
  StatisticalPeriod *int32 `json:"statisticalPeriod,omitempty" xml:"statisticalPeriod,omitempty"`
  // {'en':'Trigger threshold, unit: times.', 'zh_CN':'触发阈值，单位：次。'}
  TriggerThreshold *int32 `json:"triggerThreshold,omitempty" xml:"triggerThreshold,omitempty"`
  // {'en':'Action duration, unit: seconds, the range is 10 - 604800.', 'zh_CN':'处理动作持续时间，单位：秒，范围为 10 - 604800。'}
  InterceptTime *int32 `json:"interceptTime,omitempty" xml:"interceptTime,omitempty"`
  // {'en':'Cycle effective status.
  // PERMANENT:All time
  // WITHOUT:Excluded time
  // WITHIN:Selected time', 'zh_CN':'周期生效状态。
  // PERMANENT：永久生效
  // WITHOUT：周期内不生效
  // WITHIN：周期内生效'}
  EffectiveStatus *string `json:"effectiveStatus,omitempty" xml:"effectiveStatus,omitempty"`
  // {'en':'Effective time period.
  // When the effective status is effective within the cycle or not effective within the cycle, this field must have a value.', 'zh_CN':'规则生效周期。
  // 生效状态为周期内生效或周期内不生效时，此字段必须有值。'}
  UpdateSharedRateLimitingRuleRateLimitEffective *UpdateSharedRateLimitingRuleRateLimitEffective `json:"rateLimitEffective,omitempty" xml:"rateLimitEffective,omitempty"`
  // {'en':'Action.
  // NO_USE:Not Used
  // LOG:Log
  // COOKIE:Cookie verification
  // JS_CHECK:Javascript verification
  // DELAY:Delay
  // BLOCK:Deny
  // RESET:Reset Connection
  // Custom response ID:Custom response ID
  // (When there is a status code in the matching condition, the supported actions are Log, Deny, and Reset Connection.)', 'zh_CN':'处理动作。
  // NO_USE：不使用
  // LOG：监控
  // COOKIE：Cookie校验
  // JS_CHECK：JavaScript校验
  // DELAY：延迟响应
  // BLOCK：拦截
  // RESET：断开连接
  // 自定义响应ID：自定义响应ID
  // 当匹配条件中存在状态码时，支持处理动作为监控、拦截、断开连接。'}
  Action *string `json:"action,omitempty" xml:"action,omitempty"`
  // {'en':'Matching conditions.', 'zh_CN':'匹配条件。'}
  RateLimitRuleCondition *UpdateSharedRateLimitingRuleShareRateLimitRuleCondition `json:"rateLimitRuleCondition,omitempty" xml:"rateLimitRuleCondition,omitempty"`
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

func (s *UpdateSharedRateLimitingRuleRequest) SetStatisticalPeriod(v int32) *UpdateSharedRateLimitingRuleRequest {
  s.StatisticalPeriod = &v
  return s
}

func (s *UpdateSharedRateLimitingRuleRequest) SetTriggerThreshold(v int32) *UpdateSharedRateLimitingRuleRequest {
  s.TriggerThreshold = &v
  return s
}

func (s *UpdateSharedRateLimitingRuleRequest) SetInterceptTime(v int32) *UpdateSharedRateLimitingRuleRequest {
  s.InterceptTime = &v
  return s
}

func (s *UpdateSharedRateLimitingRuleRequest) SetEffectiveStatus(v string) *UpdateSharedRateLimitingRuleRequest {
  s.EffectiveStatus = &v
  return s
}

func (s *UpdateSharedRateLimitingRuleRequest) SetRateLimitEffective(v *UpdateSharedRateLimitingRuleRateLimitEffective) *UpdateSharedRateLimitingRuleRequest {
  s.UpdateSharedRateLimitingRuleRateLimitEffective = v
  return s
}

func (s *UpdateSharedRateLimitingRuleRequest) SetAction(v string) *UpdateSharedRateLimitingRuleRequest {
  s.Action = &v
  return s
}

func (s *UpdateSharedRateLimitingRuleRequest) SetRateLimitRuleCondition(v *UpdateSharedRateLimitingRuleShareRateLimitRuleCondition) *UpdateSharedRateLimitingRuleRequest {
  s.RateLimitRuleCondition = v
  return s
}

type UpdateSharedRateLimitingRuleResponse struct {
  // {'en':'Please refer to the error code for exceptions.', 'zh_CN':'请参照错误码。','dictionary':'belong=WAAP-MS-Ext|dict=waap_retCodeEnum'}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {'en':'Description.', 'zh_CN':'描述信息。'}
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

type UpdateSharedRateLimitingRuleIpOrIpsCondition struct {
  // {'en':'Match type.
  // EQUAL:Equals
  // NOT_EQUAL:Does not equal', 'zh_CN':'匹配类型。
  // EQUAL：等于
  // NOT_EQUAL：不等于'}
  MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
  // {'en':'IP/CIDR, maximum 500 IP/CIDR.', 'zh_CN':'IP/IP段，最多500个IP/IP段。'}
  IpOrIps []*string `json:"ipOrIps,omitempty" xml:"ipOrIps,omitempty" require:"true" type:"Repeated"`
}

func (s UpdateSharedRateLimitingRuleIpOrIpsCondition) String() string {
  return tea.Prettify(s)
}

func (s UpdateSharedRateLimitingRuleIpOrIpsCondition) GoString() string {
  return s.String()
}

func (s *UpdateSharedRateLimitingRuleIpOrIpsCondition) SetMatchType(v string) *UpdateSharedRateLimitingRuleIpOrIpsCondition {
  s.MatchType = &v
  return s
}

func (s *UpdateSharedRateLimitingRuleIpOrIpsCondition) SetIpOrIps(v []*string) *UpdateSharedRateLimitingRuleIpOrIpsCondition {
  s.IpOrIps = v
  return s
}

type UpdateSharedRateLimitingRuleHeaderCondition struct {
  // {"en":"Match type.
  // EQUAL: Equals, request header values case sensitive
  // NOT_EQUAL: Does not equal, request header values case sensitive
  // CONTAIN: Contains, request header values case insensitive
  // NOT_CONTAIN: Does not Contains, request header values case insensitive
  // NONE: Empty or non-existent
  // REGEX: Regex match, request header values case insensitive
  // NOT_REGEX: Regular does not match, request header values case insensitive
  // START_WITH: Starts with, request header values case insensitive
  // END_WITH: Ends with, request header values case insensitive
  // WILDCARD: Wildcard matches, request header values case insensitive, ** represents zero or more arbitrary characters, ? represents any single character
  // NOT_WILDCARD: Wildcard does not match, request header values case insensitive, ** represents zero or more arbitrary characters, ? represents any single character", "zh_CN":"匹配类型。
  // EQUAL：等于，头部值大小写敏感
  // NOT_EQUAL：不等于，头部值大小写敏感
  // CONTAIN：包含，头部值大小写不敏感
  // NOT_CONTAIN：不包含，头部值大小写不敏感
  // NONE：为空或不存在
  // REGEX：匹配正则，头部值大小写不敏感
  // NOT_REGEX：正则不匹配，头部值大小写不敏感
  // START_WITH：开头是，头部值大小写不敏感
  // END_WITH：结尾是，头部值大小写不敏感
  // WILDCARD：通配符匹配，头部值大小写不敏感，*代表零个或多个任意字符，?代表任意单个字符
  // NOT_WILDCARD：通配符不匹配，头部值大小写不敏感，*代表零个或多个任意字符，?代表任意单个字符"}
  MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
  // {'en':'Header name,case insensitive,up to 100 characters.
  // Example: Accept.', 'zh_CN':'头部名称，大小写不敏感，最多100个字符。
  // 示例：Accept。'}
  Key *string `json:"key,omitempty" xml:"key,omitempty" require:"true"`
  // {'en':'Header value.
  // When the match type is REGEX/NOT_REGEX, only one value is allowed.', 'zh_CN':'头部值。
  // 当匹配类型为正则/正则不匹配，则只允许只有一个值。'}
  ValueList []*string `json:"valueList,omitempty" xml:"valueList,omitempty" require:"true" type:"Repeated"`
}

func (s UpdateSharedRateLimitingRuleHeaderCondition) String() string {
  return tea.Prettify(s)
}

func (s UpdateSharedRateLimitingRuleHeaderCondition) GoString() string {
  return s.String()
}

func (s *UpdateSharedRateLimitingRuleHeaderCondition) SetMatchType(v string) *UpdateSharedRateLimitingRuleHeaderCondition {
  s.MatchType = &v
  return s
}

func (s *UpdateSharedRateLimitingRuleHeaderCondition) SetKey(v string) *UpdateSharedRateLimitingRuleHeaderCondition {
  s.Key = &v
  return s
}

func (s *UpdateSharedRateLimitingRuleHeaderCondition) SetValueList(v []*string) *UpdateSharedRateLimitingRuleHeaderCondition {
  s.ValueList = v
  return s
}

type UpdateSharedRateLimitingRuleAreaCondition struct {
  // {'en':'Match type.
  // EQUAL:Equals
  // NOT_EQUAL:Does not equal', 'zh_CN':'匹配类型。
  // EQUAL：等于
  // NOT_EQUAL：不等于'}
  MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
  // {'en':'Geo.', 'zh_CN':'区域。','dictionary':'belong=WAAP-MS-Ext|dict=waap_areaCityAndCountry'}
  Areas []*string `json:"areas,omitempty" xml:"areas,omitempty" require:"true" type:"Repeated"`
}

func (s UpdateSharedRateLimitingRuleAreaCondition) String() string {
  return tea.Prettify(s)
}

func (s UpdateSharedRateLimitingRuleAreaCondition) GoString() string {
  return s.String()
}

func (s *UpdateSharedRateLimitingRuleAreaCondition) SetMatchType(v string) *UpdateSharedRateLimitingRuleAreaCondition {
  s.MatchType = &v
  return s
}

func (s *UpdateSharedRateLimitingRuleAreaCondition) SetAreas(v []*string) *UpdateSharedRateLimitingRuleAreaCondition {
  s.Areas = v
  return s
}

type UpdateSharedRateLimitingRuleUriCondition struct {
  // {"en":"Match type.
  // EQUAL: Equals, URI case sensitive
  // NOT_EQUAL: Does not equal, URI case sensitive
  // CONTAIN: Contains, URI case insensitive
  // NOT_CONTAIN: Does not Contains, URI case insensitive
  // REGEX: Regex match, URI case insensitive
  // NOT_REGEX: Regular does not match, URI case insensitive
  // START_WITH: Starts with, URI case insensitive
  // END_WITH: Ends with, URI case insensitive
  // WILDCARD: Wildcard matches, URI case insensitive, ** represents zero or more arbitrary characters, ? represents any single character
  // NOT_WILDCARD: Wildcard does not match, URI case insensitive, ** represents zero or more arbitrary characters, ? represents any single character", "zh_CN":"匹配类型。
  // EQUAL：等于，URI大小写敏感
  // NOT_EQUAL：不等于，URI大小写敏感
  // CONTAIN：包含，URI大小写不敏感
  // NOT_CONTAIN：不包含，URI大小写不敏感
  // REGEX：匹配正则，URI大小写不敏感
  // NOT_REGEX：正则不匹配，URI大小写不敏感
  // START_WITH：开头是，URI大小写不敏感
  // END_WITH：结尾是，URI大小写不敏感
  // WILDCARD：通配符匹配，URI大小写不敏感，*代表零个或多个任意字符，?代表任意单个字符
  // NOT_WILDCARD：通配符不匹配，URI大小写不敏感，*代表零个或多个任意字符，?代表任意单个字符"}
  MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
  // {'en':'URI.
  // When match type is EQUAL/NOT_EQUAL/START_WITH/END_WITH, uri needs to start with "/", and includes parameters.
  // When the match type is REGEX/NOT_REGEX, only one value is allowed. 
  // Example: /test.html?id=1.', 'zh_CN':'URI。
  // 当匹配类型为等于/不等于/开头是/结尾是，URI必须以”/“开头，含参数。
  // 当匹配类型为正则/正则不匹配，则只允许只有一个值。
  // 示例：/test.html?id=1。'}
  Uri []*string `json:"uri,omitempty" xml:"uri,omitempty" require:"true" type:"Repeated"`
}

func (s UpdateSharedRateLimitingRuleUriCondition) String() string {
  return tea.Prettify(s)
}

func (s UpdateSharedRateLimitingRuleUriCondition) GoString() string {
  return s.String()
}

func (s *UpdateSharedRateLimitingRuleUriCondition) SetMatchType(v string) *UpdateSharedRateLimitingRuleUriCondition {
  s.MatchType = &v
  return s
}

func (s *UpdateSharedRateLimitingRuleUriCondition) SetUri(v []*string) *UpdateSharedRateLimitingRuleUriCondition {
  s.Uri = v
  return s
}

type UpdateSharedRateLimitingRuleUaCondition struct {
  // {"en":"Match type.
  // EQUAL: Equals, user agent case sensitive
  // NOT_EQUAL: Does not equal, user agent case sensitive
  // CONTAIN: Contains, user agent case insensitive
  // NOT_CONTAIN: Does not Contains, user agent case insensitive
  // NONE:Empty or non-existent
  // REGEX: Regex match, user agent case insensitive
  // NOT_REGEX: Regular does not match, user agent case insensitive
  // START_WITH: Starts with, user agent case insensitive
  // END_WITH: Ends with, user agent case insensitive
  // WILDCARD: Wildcard matches, user agent case insensitive, ** represents zero or more arbitrary characters, ? represents any single character
  // NOT_WILDCARD: Wildcard does not match, user agent case insensitive, ** represents zero or more arbitrary characters, ? represents any single character", "zh_CN":"匹配类型。
  // EQUAL：等于，User-Agent大小写敏感
  // NOT_EQUAL：不等于，User-Agent大小写敏感
  // CONTAIN：包含，User-Agent大小写不敏感
  // NOT_CONTAIN：不包含，User-Agent大小写不敏感
  // NONE：为空或不存在
  // REGEX：匹配正则，User-Agent大小写不敏感
  // NOT_REGEX：正则不匹配，User-Agent大小写不敏感
  // START_WITH：开头是，User-Agent大小写不敏感
  // END_WITH：结尾是，User-Agent大小写不敏感
  // WILDCARD：通配符匹配，User-Agent大小写不敏感，*代表零个或多个任意字符，?代表任意单个字符
  // NOT_WILDCARD：通配符不匹配，User-Agent大小写不敏感，*代表零个或多个任意字符，?代表任意单个字符"}
  MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
  // {'en':'User agent.
  // When the match type is REGEX/NOT_REGEX, only one value is allowed. 
  // Example: go-Http-client/1.1.', 'zh_CN':'User-Agent。
  // 当匹配类型为正则/正则不匹配，则只允许只有一个值。
  // 示例：go-Http-client/1.1。'}
  Ua []*string `json:"ua,omitempty" xml:"ua,omitempty" require:"true" type:"Repeated"`
}

func (s UpdateSharedRateLimitingRuleUaCondition) String() string {
  return tea.Prettify(s)
}

func (s UpdateSharedRateLimitingRuleUaCondition) GoString() string {
  return s.String()
}

func (s *UpdateSharedRateLimitingRuleUaCondition) SetMatchType(v string) *UpdateSharedRateLimitingRuleUaCondition {
  s.MatchType = &v
  return s
}

func (s *UpdateSharedRateLimitingRuleUaCondition) SetUa(v []*string) *UpdateSharedRateLimitingRuleUaCondition {
  s.Ua = v
  return s
}

type UpdateSharedRateLimitingRuleRefererCondition struct {
  // {"en":"Match type.
  // EQUAL: Equals, referer case sensitive
  // NOT_EQUAL: Does not equal, referer case sensitive
  // CONTAIN: Contains, referer case insensitive
  // NOT_CONTAIN: Does not Contains, referer case insensitive
  // NONE:Empty or non-existent
  // REGEX: Regex match, referer case insensitive
  // NOT_REGEX: Regular does not match, referer case insensitive
  // START_WITH: Starts with, referer case insensitive
  // END_WITH: Ends with, referer case insensitive
  // WILDCARD: Wildcard matches, referer case insensitive, ** represents zero or more arbitrary characters, ? represents any single characte
  // NOT_WILDCARD: Wildcard does not match, referer case insensitive, ** represents zero or more arbitrary characters, ? represents any single character", "zh_CN":"匹配类型。
  // EQUAL：等于，referer大小写敏感
  // NOT_EQUAL：不等于，referer大小写敏感
  // CONTAIN：包含，referer大小写不敏感
  // NOT_CONTAIN：不包含，referer大小写不敏感
  // NONE：为空或不存在
  // REGEX：匹配正则，referer大小写不敏感
  // NOT_REGEX：正则不匹配，referer大小写不敏感
  // START_WITH：开头是，referer大小写不敏感
  // END_WITH：结尾是，referer大小写不敏感
  // WILDCARD：通配符匹配，referer大小写不敏感，*代表零个或多个任意字符，?代表任意单个字符
  // NOT_WILDCARD：通配符不匹配，referer大小写不敏感，*代表零个或多个任意字符，?代表任意单个字符"}
  MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
  // {'en':'Referer.
  // When the match type is REGEX/NOT_REGEX, only one value is allowed. 
  // Example: http://test.com.', 'zh_CN':'Referer。
  // 当匹配类型为正则/正则不匹配，则只允许只有一个值。
  // 示例：http://test.com。'}
  Referer []*string `json:"referer,omitempty" xml:"referer,omitempty" require:"true" type:"Repeated"`
}

func (s UpdateSharedRateLimitingRuleRefererCondition) String() string {
  return tea.Prettify(s)
}

func (s UpdateSharedRateLimitingRuleRefererCondition) GoString() string {
  return s.String()
}

func (s *UpdateSharedRateLimitingRuleRefererCondition) SetMatchType(v string) *UpdateSharedRateLimitingRuleRefererCondition {
  s.MatchType = &v
  return s
}

func (s *UpdateSharedRateLimitingRuleRefererCondition) SetReferer(v []*string) *UpdateSharedRateLimitingRuleRefererCondition {
  s.Referer = v
  return s
}

type UpdateSharedRateLimitingRuleStatusCodeCondition struct {
  // {'en':'Match type.
  // EQUAL:Equals
  // NOT_EQUAL:Does not equal', 'zh_CN':'匹配类型。
  // EQUAL：等于
  // NOT_EQUAL：不等于'}
  MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
  // {'en':'Response Code.', 'zh_CN':'状态码。'}
  StatusCode []*string `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true" type:"Repeated"`
}

func (s UpdateSharedRateLimitingRuleStatusCodeCondition) String() string {
  return tea.Prettify(s)
}

func (s UpdateSharedRateLimitingRuleStatusCodeCondition) GoString() string {
  return s.String()
}

func (s *UpdateSharedRateLimitingRuleStatusCodeCondition) SetMatchType(v string) *UpdateSharedRateLimitingRuleStatusCodeCondition {
  s.MatchType = &v
  return s
}

func (s *UpdateSharedRateLimitingRuleStatusCodeCondition) SetStatusCode(v []*string) *UpdateSharedRateLimitingRuleStatusCodeCondition {
  s.StatusCode = v
  return s
}

type UpdateSharedRateLimitingRuleRequestMethodCondition struct {
  // {'en':'Match type.
  // EQUAL:Equals
  // NOT_EQUAL:Does not equal', 'zh_CN':'匹配类型。
  // EQUAL：等于
  // NOT_EQUAL：不等于'}
  MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
  // {'en':'Request method.
  // Supported values: GET/POST/DELETE/PUT/HEAD/OPTIONS/COPY.', 'zh_CN':'请求方法。
  // 支持的值：GET/POST/DELETE/PUT/HEAD/OPTIONS/COPY。'}
  RequestMethod []*string `json:"requestMethod,omitempty" xml:"requestMethod,omitempty" require:"true" type:"Repeated"`
}

func (s UpdateSharedRateLimitingRuleRequestMethodCondition) String() string {
  return tea.Prettify(s)
}

func (s UpdateSharedRateLimitingRuleRequestMethodCondition) GoString() string {
  return s.String()
}

func (s *UpdateSharedRateLimitingRuleRequestMethodCondition) SetMatchType(v string) *UpdateSharedRateLimitingRuleRequestMethodCondition {
  s.MatchType = &v
  return s
}

func (s *UpdateSharedRateLimitingRuleRequestMethodCondition) SetRequestMethod(v []*string) *UpdateSharedRateLimitingRuleRequestMethodCondition {
  s.RequestMethod = v
  return s
}

type UpdateSharedRateLimitingRuleSchemeCondition struct {
  // {'en':'Match type.
  // EQUAL:Equals
  // NOT_EQUAL:Does not equal', 'zh_CN':'匹配类型。
  // EQUAL：等于
  // NOT_EQUAL：不等于'}
  MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
  // {'en':'HTTP/S.
  // Supported values: HTTP/HTTPS.', 'zh_CN':'应用层协议。
  // 支持的值：HTTP/HTTPS。'}
  Scheme []*string `json:"scheme,omitempty" xml:"scheme,omitempty" require:"true" type:"Repeated"`
}

func (s UpdateSharedRateLimitingRuleSchemeCondition) String() string {
  return tea.Prettify(s)
}

func (s UpdateSharedRateLimitingRuleSchemeCondition) GoString() string {
  return s.String()
}

func (s *UpdateSharedRateLimitingRuleSchemeCondition) SetMatchType(v string) *UpdateSharedRateLimitingRuleSchemeCondition {
  s.MatchType = &v
  return s
}

func (s *UpdateSharedRateLimitingRuleSchemeCondition) SetScheme(v []*string) *UpdateSharedRateLimitingRuleSchemeCondition {
  s.Scheme = v
  return s
}

type UpdateSharedRateLimitingRuleShareRateLimitRuleCondition struct {
  // {'en':'IP/CIDR, match type cannot be repeated.', 'zh_CN':'IP/IP段，匹配类型不可重复。'}
  IpOrIpsConditions []*UpdateSharedRateLimitingRuleIpOrIpsCondition `json:"ipOrIpsConditions,omitempty" xml:"ipOrIpsConditions,omitempty" type:"Repeated"`
  // {'en':'Path, match type cannot be repeated.
  // When the business scenario is API, this matching condition is not supported.', 'zh_CN':'路径，匹配类型不可重复。
  // 当业务场景为API业务时不支持此匹配条件。'}
  PathConditions []*UpdateSharedRateLimitingRulePathCondition `json:"pathConditions,omitempty" xml:"pathConditions,omitempty" type:"Repeated"`
  // {'en':'URI, match type cannot be repeated.
  // 
  // When the business scenario is API, this matching condition is not supported.', 'zh_CN':'URI，匹配类型不可重复。
  // 当业务场景为API业务时不支持此匹配条件。'}
  UriConditions []*UpdateSharedRateLimitingRuleUriCondition `json:"uriConditions,omitempty" xml:"uriConditions,omitempty" type:"Repeated"`
  // {'en':'User Agent, match type cannot be repeated.', 'zh_CN':'User-Agent，匹配类型不可重复。'}
  UaConditions []*UpdateSharedRateLimitingRuleUaCondition `json:"uaConditions,omitempty" xml:"uaConditions,omitempty" type:"Repeated"`
  // {'en':'Request Method.
  // When the business scenario is API,this matching condition is not supported.', 'zh_CN':'请求方法，匹配类型不可重复。
  // 当业务场景为API业务时不支持此匹配条件。'}
  MethodConditions []*UpdateSharedRateLimitingRuleRequestMethodCondition `json:"methodConditions,omitempty" xml:"methodConditions,omitempty" type:"Repeated"`
  // {'en':'Referer, match type cannot be repeated.', 'zh_CN':'Referer，匹配类型不可重复。'}
  RefererConditions []*UpdateSharedRateLimitingRuleRefererCondition `json:"refererConditions,omitempty" xml:"refererConditions,omitempty" type:"Repeated"`
  // {'en':'Request Header, match type can be repeated.', 'zh_CN':'请求头，匹配类型可重复。'}
  HeaderConditions []*UpdateSharedRateLimitingRuleHeaderCondition `json:"headerConditions,omitempty" xml:"headerConditions,omitempty" type:"Repeated"`
  // {'en':'Geo,match type cannot be repeated.', 'zh_CN':'区域，匹配类型不可重复。'}
  AreaConditions []*UpdateSharedRateLimitingRuleAreaCondition `json:"areaConditions,omitempty" xml:"areaConditions,omitempty" type:"Repeated"`
  // {'en':'Response Code, match type cannot be repeated.', 'zh_CN':'状态码，匹配类型不可重复。'}
  StatusCodeConditions []*UpdateSharedRateLimitingRuleStatusCodeCondition `json:"statusCodeConditions,omitempty" xml:"statusCodeConditions,omitempty" type:"Repeated"`
  // {'en':'HTTP/S, match type cannot be repeated.', 'zh_CN':'应用层协议，匹配类型不可重复。'}
  SchemeConditions []*UpdateSharedRateLimitingRuleSchemeCondition `json:"schemeConditions,omitempty" xml:"schemeConditions,omitempty" type:"Repeated"`
}

func (s UpdateSharedRateLimitingRuleShareRateLimitRuleCondition) String() string {
  return tea.Prettify(s)
}

func (s UpdateSharedRateLimitingRuleShareRateLimitRuleCondition) GoString() string {
  return s.String()
}

func (s *UpdateSharedRateLimitingRuleShareRateLimitRuleCondition) SetIpOrIpsConditions(v []*UpdateSharedRateLimitingRuleIpOrIpsCondition) *UpdateSharedRateLimitingRuleShareRateLimitRuleCondition {
  s.IpOrIpsConditions = v
  return s
}

func (s *UpdateSharedRateLimitingRuleShareRateLimitRuleCondition) SetPathConditions(v []*UpdateSharedRateLimitingRulePathCondition) *UpdateSharedRateLimitingRuleShareRateLimitRuleCondition {
  s.PathConditions = v
  return s
}

func (s *UpdateSharedRateLimitingRuleShareRateLimitRuleCondition) SetUriConditions(v []*UpdateSharedRateLimitingRuleUriCondition) *UpdateSharedRateLimitingRuleShareRateLimitRuleCondition {
  s.UriConditions = v
  return s
}

func (s *UpdateSharedRateLimitingRuleShareRateLimitRuleCondition) SetUaConditions(v []*UpdateSharedRateLimitingRuleUaCondition) *UpdateSharedRateLimitingRuleShareRateLimitRuleCondition {
  s.UaConditions = v
  return s
}

func (s *UpdateSharedRateLimitingRuleShareRateLimitRuleCondition) SetMethodConditions(v []*UpdateSharedRateLimitingRuleRequestMethodCondition) *UpdateSharedRateLimitingRuleShareRateLimitRuleCondition {
  s.MethodConditions = v
  return s
}

func (s *UpdateSharedRateLimitingRuleShareRateLimitRuleCondition) SetRefererConditions(v []*UpdateSharedRateLimitingRuleRefererCondition) *UpdateSharedRateLimitingRuleShareRateLimitRuleCondition {
  s.RefererConditions = v
  return s
}

func (s *UpdateSharedRateLimitingRuleShareRateLimitRuleCondition) SetHeaderConditions(v []*UpdateSharedRateLimitingRuleHeaderCondition) *UpdateSharedRateLimitingRuleShareRateLimitRuleCondition {
  s.HeaderConditions = v
  return s
}

func (s *UpdateSharedRateLimitingRuleShareRateLimitRuleCondition) SetAreaConditions(v []*UpdateSharedRateLimitingRuleAreaCondition) *UpdateSharedRateLimitingRuleShareRateLimitRuleCondition {
  s.AreaConditions = v
  return s
}

func (s *UpdateSharedRateLimitingRuleShareRateLimitRuleCondition) SetStatusCodeConditions(v []*UpdateSharedRateLimitingRuleStatusCodeCondition) *UpdateSharedRateLimitingRuleShareRateLimitRuleCondition {
  s.StatusCodeConditions = v
  return s
}

func (s *UpdateSharedRateLimitingRuleShareRateLimitRuleCondition) SetSchemeConditions(v []*UpdateSharedRateLimitingRuleSchemeCondition) *UpdateSharedRateLimitingRuleShareRateLimitRuleCondition {
  s.SchemeConditions = v
  return s
}

type UpdateSharedRateLimitingRuleRateLimitEffective struct {
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

func (s UpdateSharedRateLimitingRuleRateLimitEffective) String() string {
  return tea.Prettify(s)
}

func (s UpdateSharedRateLimitingRuleRateLimitEffective) GoString() string {
  return s.String()
}

func (s *UpdateSharedRateLimitingRuleRateLimitEffective) SetEffective(v []*string) *UpdateSharedRateLimitingRuleRateLimitEffective {
  s.Effective = v
  return s
}

func (s *UpdateSharedRateLimitingRuleRateLimitEffective) SetStart(v string) *UpdateSharedRateLimitingRuleRateLimitEffective {
  s.Start = &v
  return s
}

func (s *UpdateSharedRateLimitingRuleRateLimitEffective) SetEnd(v string) *UpdateSharedRateLimitingRuleRateLimitEffective {
  s.End = &v
  return s
}

func (s *UpdateSharedRateLimitingRuleRateLimitEffective) SetTimezone(v string) *UpdateSharedRateLimitingRuleRateLimitEffective {
  s.Timezone = &v
  return s
}

type UpdateSharedRateLimitingRulePathCondition struct {
  // {"en":"Match type.
  // EQUAL: Equals, path case sensitive
  // NOT_EQUAL: Does not equal, path case sensitive
  // CONTAIN: Contains, path case insensitive
  // NOT_CONTAIN: Does not Contains, path case insensitive
  // REGEX: Regex match, path case insensitive
  // NOT_REGEX: Regular does not match, path case sensitive
  // START_WITH: Starts with, path case sensitive
  // END_WITH: Ends with, path case sensitive
  // WILDCARD: Wildcard matches, path case sensitive, ** represents zero or more arbitrary characters, ? represents any single character.
  // NOT_WILDCARD: Wildcard does not match, path case sensitive, ** represents zero or more arbitrary characters, ? represents any single character ", "zh_CN":"匹配类型。
  // EQUAL：等于，路径大小写敏感
  // NOT_EQUAL：不等于，路径大小写敏感
  // CONTAIN：包含，路径大小写不敏感
  // NOT_CONTAIN：不包含，路径大小写不敏感
  // REGEX：匹配正则，路径大小写不敏感
  // NOT_REGEX：正则不匹配，路径大小写不敏感
  // START_WITH：开头是，路径大小写不敏感
  // END_WITH：结尾是，路径大小写不敏感
  // WILDCARD：通配符匹配，路径大小写不敏感，*代表零个或多个任意字符，?代表任意单个字符
  // NOT_WILDCARD：通配符不匹配，路径大小写不敏感，*代表零个或多个任意字符，?代表任意单个字符"}
  MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
  // {'en':'Path.
  // When match type is EQUAL/NOT_EQUAL/START_WITH/END_WITH, path needs to start with "/", and no parameters.
  // When the match type is REGEX/NOT_REGEX, only one value is allowed. 
  // Example: /test.html.', 'zh_CN':'路径。
  // 当匹配类型为等于/不等于/开头是/结尾是，路径必须以“/”开头，不含参数。
  // 当匹配类型为正则/正则不匹配，则只允许只有一个值。
  // 示例：/test.html。'}
  UpdateSharedRateLimitingRulePaths []*string `json:"paths,omitempty" xml:"paths,omitempty" require:"true" type:"Repeated"`
}

func (s UpdateSharedRateLimitingRulePathCondition) String() string {
  return tea.Prettify(s)
}

func (s UpdateSharedRateLimitingRulePathCondition) GoString() string {
  return s.String()
}

func (s *UpdateSharedRateLimitingRulePathCondition) SetMatchType(v string) *UpdateSharedRateLimitingRulePathCondition {
  s.MatchType = &v
  return s
}

func (s *UpdateSharedRateLimitingRulePathCondition) SetUpdateSharedRateLimitingRulePaths(v []*string) *UpdateSharedRateLimitingRulePathCondition {
  s.UpdateSharedRateLimitingRulePaths = v
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

type UpdateSharedRateLimitingRuleRequestHeader struct {
  // {"zh_CN":"安全服务类型。有使用多个不同的安全服务时，需要填写具体的服务类型。","en":"Security service type. Please enter a specific service type, if you purchase multiple security services.","dictionary":"belong=WAAP-MS-Ext|dict=waap_serviceType"}
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

type UpdateSharedRateLimitingRuleResponseHeader struct {
}

func (s UpdateSharedRateLimitingRuleResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s UpdateSharedRateLimitingRuleResponseHeader) GoString() string {
  return s.String()
}




type AssociateShareCustomizeRuleRequest struct {
  // {"zh_CN":"共享规则ID。","en":"Share rule ID."}
  ShareId *string `json:"shareId,omitempty" xml:"shareId,omitempty" require:"true"`
  // {"zh_CN":"待关联域名列表。","en":"List of hostname to be associated."}
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

type AssociateShareCustomizeRuleResponse struct {
  // {'en':'Please refer to the error code for exceptions.', 'zh_CN':'请参照错误码。','dictionary':'belong=WAAP-MS-Ext|dict=waap_retCodeEnum'}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {'en':'Description.', 'zh_CN':'描述信息。'}
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

type AssociateShareCustomizeRuleRequestHeader struct {
  // {"zh_CN":"安全服务类型。有使用多个不同的安全服务时，需要填写具体的服务类型。","en":"Security service type. Please enter a specific service type, if you purchase multiple security services.","dictionary":"belong=WAAP-MS-Ext|dict=waap_serviceType"}
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

type AssociateShareCustomizeRuleResponseHeader struct {
}

func (s AssociateShareCustomizeRuleResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s AssociateShareCustomizeRuleResponseHeader) GoString() string {
  return s.String()
}




type CreateSharedWhitelistRuleRequest struct {
  // {"en":"Rule name, maximum 50 characters.
  //  Does not support special characters and spaces.", "zh_CN":"规则名称，最多50个字符。
  // 不支持特殊字符和空格。"}
  RuleName *string `json:"ruleName,omitempty" xml:"ruleName,omitempty" require:"true"`
  // {"en":"Description, maximum 200 characters.", "zh_CN":"描述，最多200个字符。"}
  Description *string `json:"description,omitempty" xml:"description,omitempty"`
  // {"en":"Match conditions, at least one, at most five.", "zh_CN":"匹配条件，至少一个，至多五个。"}
  Conditions *CreateSharedWhitelistRuleWhitelistRuleCondition `json:"conditions,omitempty" xml:"conditions,omitempty" require:"true"`
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

func (s *CreateSharedWhitelistRuleRequest) SetConditions(v *CreateSharedWhitelistRuleWhitelistRuleCondition) *CreateSharedWhitelistRuleRequest {
  s.Conditions = v
  return s
}

type CreateSharedWhitelistRuleWhitelistRuleCondition struct {
  // {"en":"IP/CIDR match conditions, match type cannot be repeated.", "zh_CN":"IP/IP段匹配条件，匹配类型不可重复。"}
  IpOrIpsConditions []*CreateSharedWhitelistRuleIpOrIpsCondition `json:"ipOrIpsConditions,omitempty" xml:"ipOrIpsConditions,omitempty" type:"Repeated"`
  // {"en":"Path match conditions, match type cannot be repeated.", "zh_CN":"路径匹配条件，匹配类型不可重复。"}
  PathConditions []*CreateSharedWhitelistRulePathCondition `json:"pathConditions,omitempty" xml:"pathConditions,omitempty" type:"Repeated"`
  // {"en":"URI match conditions, match type cannot be repeated.", "zh_CN":"URI匹配条件，匹配类型不可重复。"}
  UriConditions []*CreateSharedWhitelistRuleUriCondition `json:"uriConditions,omitempty" xml:"uriConditions,omitempty" type:"Repeated"`
  // {"en":"User agent match conditions, match type cannot be repeated.", "zh_CN":"User-Agent 匹配条件，匹配类型不可重复。"}
  UaConditions []*CreateSharedWhitelistRuleUaCondition `json:"uaConditions,omitempty" xml:"uaConditions,omitempty" type:"Repeated"`
  // {"en":"Referer match conditions, match type cannot be repeated.", "zh_CN":"Referer 匹配条件，匹配类型不可重复。"}
  RefererConditions []*CreateSharedWhitelistRuleRefererCondition `json:"refererConditions,omitempty" xml:"refererConditions,omitempty" type:"Repeated"`
  // {"en":"Request header match conditions.", "zh_CN":"请求头匹配条件。"}
  HeaderConditions []*CreateSharedWhitelistRuleHeaderCondition `json:"headerConditions,omitempty" xml:"headerConditions,omitempty" type:"Repeated"`
}

func (s CreateSharedWhitelistRuleWhitelistRuleCondition) String() string {
  return tea.Prettify(s)
}

func (s CreateSharedWhitelistRuleWhitelistRuleCondition) GoString() string {
  return s.String()
}

func (s *CreateSharedWhitelistRuleWhitelistRuleCondition) SetIpOrIpsConditions(v []*CreateSharedWhitelistRuleIpOrIpsCondition) *CreateSharedWhitelistRuleWhitelistRuleCondition {
  s.IpOrIpsConditions = v
  return s
}

func (s *CreateSharedWhitelistRuleWhitelistRuleCondition) SetPathConditions(v []*CreateSharedWhitelistRulePathCondition) *CreateSharedWhitelistRuleWhitelistRuleCondition {
  s.PathConditions = v
  return s
}

func (s *CreateSharedWhitelistRuleWhitelistRuleCondition) SetUriConditions(v []*CreateSharedWhitelistRuleUriCondition) *CreateSharedWhitelistRuleWhitelistRuleCondition {
  s.UriConditions = v
  return s
}

func (s *CreateSharedWhitelistRuleWhitelistRuleCondition) SetUaConditions(v []*CreateSharedWhitelistRuleUaCondition) *CreateSharedWhitelistRuleWhitelistRuleCondition {
  s.UaConditions = v
  return s
}

func (s *CreateSharedWhitelistRuleWhitelistRuleCondition) SetRefererConditions(v []*CreateSharedWhitelistRuleRefererCondition) *CreateSharedWhitelistRuleWhitelistRuleCondition {
  s.RefererConditions = v
  return s
}

func (s *CreateSharedWhitelistRuleWhitelistRuleCondition) SetHeaderConditions(v []*CreateSharedWhitelistRuleHeaderCondition) *CreateSharedWhitelistRuleWhitelistRuleCondition {
  s.HeaderConditions = v
  return s
}

type CreateSharedWhitelistRuleIpOrIpsCondition struct {
  // {"en":"Match type.
  // EQUAL: Equals
  // NOT_EQUAL: Does not equal", "zh_CN":"匹配类型。
  // EQUAL：等于
  // NOT_EQUAL：不等于"}
  MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
  // {"en":"IP/CIDR, maximum 500 IP/CIDR.", "zh_CN":"IP/IP段，最多500个IP/IP段。"}
  IpOrIps []*string `json:"ipOrIps,omitempty" xml:"ipOrIps,omitempty" require:"true" type:"Repeated"`
}

func (s CreateSharedWhitelistRuleIpOrIpsCondition) String() string {
  return tea.Prettify(s)
}

func (s CreateSharedWhitelistRuleIpOrIpsCondition) GoString() string {
  return s.String()
}

func (s *CreateSharedWhitelistRuleIpOrIpsCondition) SetMatchType(v string) *CreateSharedWhitelistRuleIpOrIpsCondition {
  s.MatchType = &v
  return s
}

func (s *CreateSharedWhitelistRuleIpOrIpsCondition) SetIpOrIps(v []*string) *CreateSharedWhitelistRuleIpOrIpsCondition {
  s.IpOrIps = v
  return s
}

type CreateSharedWhitelistRulePathCondition struct {
  // {"en":"Match type.
  // EQUAL: Equals, path case sensitive
  // NOT_EQUAL: Does not equal, path case sensitive
  // CONTAIN: Contains, path case insensitive
  // NOT_CONTAIN: Does not Contains, path case insensitive
  // REGEX: Regex match, path case insensitive
  // NOT_REGEX: Regular does not match, path case sensitive
  // START_WITH: Starts with, path case sensitive
  // END_WITH: Ends with, path case sensitive
  // WILDCARD: Wildcard matches, path case sensitive, ** represents zero or more arbitrary characters, ? represents any single character.
  // NOT_WILDCARD: Wildcard does not match, path case sensitive, ** represents zero or more arbitrary characters, ? represents any single character ", "zh_CN":"匹配类型。
  // EQUAL：等于，路径大小写敏感
  // NOT_EQUAL：不等于，路径大小写敏感
  // CONTAIN：包含，路径大小写不敏感
  // NOT_CONTAIN：不包含，路径大小写不敏感
  // REGEX：匹配正则，路径大小写不敏感
  // NOT_REGEX：正则不匹配，路径大小写不敏感
  // START_WITH：开头是，路径大小写不敏感
  // END_WITH：结尾是，路径大小写不敏感
  // WILDCARD：通配符匹配，路径大小写不敏感，*代表零个或多个任意字符，?代表任意单个字符
  // NOT_WILDCARD：通配符不匹配，路径大小写不敏感，*代表零个或多个任意字符，?代表任意单个字符"}
  MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
  // {'en':'Path.
  // When match type is EQUAL/NOT_EQUAL/START_WITH/END_WITH, path needs to start with "/", and no parameters.
  // When the match type is REGEX/NOT_REGEX, only one value is allowed. 
  // Example: /test.html.', 'zh_CN':'路径。
  // 当匹配类型为等于/不等于/开头是/结尾是，路径必须以“/”开头，不含参数。
  // 当匹配类型为正则/正则不匹配，则只允许只有一个值。
  // 示例：/test.html。'}
  CreateSharedWhitelistRulePaths []*string `json:"paths,omitempty" xml:"paths,omitempty" require:"true" type:"Repeated"`
}

func (s CreateSharedWhitelistRulePathCondition) String() string {
  return tea.Prettify(s)
}

func (s CreateSharedWhitelistRulePathCondition) GoString() string {
  return s.String()
}

func (s *CreateSharedWhitelistRulePathCondition) SetMatchType(v string) *CreateSharedWhitelistRulePathCondition {
  s.MatchType = &v
  return s
}

func (s *CreateSharedWhitelistRulePathCondition) SetCreateSharedWhitelistRulePaths(v []*string) *CreateSharedWhitelistRulePathCondition {
  s.CreateSharedWhitelistRulePaths = v
  return s
}

type CreateSharedWhitelistRuleUriCondition struct {
  // {"en":"Match type.
  // EQUAL: Equals, URI case sensitive
  // NOT_EQUAL: Does not equal, URI case sensitive
  // CONTAIN: Contains, URI case insensitive
  // NOT_CONTAIN: Does not Contains, URI case insensitive
  // REGEX: Regex match, URI case insensitive
  // NOT_REGEX: Regular does not match, URI case insensitive
  // START_WITH: Starts with, URI case insensitive
  // END_WITH: Ends with, URI case insensitive
  // WILDCARD: Wildcard matches, URI case insensitive, ** represents zero or more arbitrary characters, ? represents any single character
  // NOT_WILDCARD: Wildcard does not match, URI case insensitive, ** represents zero or more arbitrary characters, ? represents any single character", "zh_CN":"匹配类型。
  // EQUAL：等于，URI大小写敏感
  // NOT_EQUAL：不等于，URI大小写敏感
  // CONTAIN：包含，URI大小写不敏感
  // NOT_CONTAIN：不包含，URI大小写不敏感
  // REGEX：匹配正则，URI大小写不敏感
  // NOT_REGEX：正则不匹配，URI大小写不敏感
  // START_WITH：开头是，URI大小写不敏感
  // END_WITH：结尾是，URI大小写不敏感
  // WILDCARD：通配符匹配，URI大小写不敏感，*代表零个或多个任意字符，?代表任意单个字符
  // NOT_WILDCARD：通配符不匹配，URI大小写不敏感，*代表零个或多个任意字符，?代表任意单个字符"}
  MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
  // {'en':'URI.
  // When match type is EQUAL/NOT_EQUAL/START_WITH/END_WITH, uri needs to start with "/", and includes parameters.
  // When the match type is REGEX/NOT_REGEX, only one value is allowed. 
  // Example: /test.html?id=1.', 'zh_CN':'URI。
  // 当匹配类型为等于/不等于/开头是/结尾是，URI必须以”/“开头，含参数。
  // 当匹配类型为正则/正则不匹配，则只允许只有一个值。
  // 示例：/test.html?id=1。'}
  Uri []*string `json:"uri,omitempty" xml:"uri,omitempty" require:"true" type:"Repeated"`
}

func (s CreateSharedWhitelistRuleUriCondition) String() string {
  return tea.Prettify(s)
}

func (s CreateSharedWhitelistRuleUriCondition) GoString() string {
  return s.String()
}

func (s *CreateSharedWhitelistRuleUriCondition) SetMatchType(v string) *CreateSharedWhitelistRuleUriCondition {
  s.MatchType = &v
  return s
}

func (s *CreateSharedWhitelistRuleUriCondition) SetUri(v []*string) *CreateSharedWhitelistRuleUriCondition {
  s.Uri = v
  return s
}

type CreateSharedWhitelistRuleUaCondition struct {
  // {"en":"Match type.
  // EQUAL: Equals, user agent case sensitive
  // NOT_EQUAL: Does not equal, user agent case sensitive
  // CONTAIN: Contains, user agent case insensitive
  // NOT_CONTAIN: Does not Contains, user agent case insensitive
  // REGEX: Regex match, user agent case insensitive
  // NOT_REGEX: Regular does not match, user agent case insensitive
  // START_WITH: Starts with, user agent case insensitive
  // END_WITH: Ends with, user agent case insensitive
  // WILDCARD: Wildcard matches, user agent case insensitive, ** represents zero or more arbitrary characters, ? represents any single character
  // NOT_WILDCARD: Wildcard does not match, user agent case insensitive, ** represents zero or more arbitrary characters, ? represents any single character", "zh_CN":"匹配类型。
  // EQUAL：等于，User-Agent大小写敏感
  // NOT_EQUAL：不等于，User-Agent大小写敏感
  // CONTAIN：包含，User-Agent大小写不敏感
  // NOT_CONTAIN：不包含，User-Agent大小写不敏感
  // REGEX：匹配正则，User-Agent大小写不敏感
  // NOT_REGEX：正则不匹配，User-Agent大小写不敏感
  // START_WITH：开头是，User-Agent大小写不敏感
  // END_WITH：结尾是，User-Agent大小写不敏感
  // WILDCARD：通配符匹配，User-Agent大小写不敏感，*代表零个或多个任意字符，?代表任意单个字符
  // NOT_WILDCARD：通配符不匹配，User-Agent大小写不敏感，*代表零个或多个任意字符，?代表任意单个字符"}
  MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
  // {'en':'User agent.
  // When the match type is REGEX/NOT_REGEX, only one value is allowed. 
  // Example: go-Http-client/1.1.', 'zh_CN':'User-Agent。
  // 当匹配类型为正则/正则不匹配，则只允许只有一个值。
  // 示例：go-Http-client/1.1。'}
  Ua []*string `json:"ua,omitempty" xml:"ua,omitempty" require:"true" type:"Repeated"`
}

func (s CreateSharedWhitelistRuleUaCondition) String() string {
  return tea.Prettify(s)
}

func (s CreateSharedWhitelistRuleUaCondition) GoString() string {
  return s.String()
}

func (s *CreateSharedWhitelistRuleUaCondition) SetMatchType(v string) *CreateSharedWhitelistRuleUaCondition {
  s.MatchType = &v
  return s
}

func (s *CreateSharedWhitelistRuleUaCondition) SetUa(v []*string) *CreateSharedWhitelistRuleUaCondition {
  s.Ua = v
  return s
}

type CreateSharedWhitelistRuleRefererCondition struct {
  // {"en":"Match type.
  // EQUAL: Equals, referer case sensitive
  // NOT_EQUAL: Does not equal, referer case sensitive
  // CONTAIN: Contains, referer case insensitive
  // NOT_CONTAIN: Does not Contains, referer case insensitive
  // REGEX: Regex match, referer case insensitive
  // NOT_REGEX: Regular does not match, referer case insensitive
  // START_WITH: Starts with, referer case insensitive
  // END_WITH: Ends with, referer case insensitive
  // WILDCARD: Wildcard matches, referer case insensitive, ** represents zero or more arbitrary characters, ? represents any single characte
  // NOT_WILDCARD: Wildcard does not match, referer case insensitive, ** represents zero or more arbitrary characters, ? represents any single character", "zh_CN":"匹配类型。
  // EQUAL：等于，referer大小写敏感
  // NOT_EQUAL：不等于，referer大小写敏感
  // CONTAIN：包含，referer大小写不敏感
  // NOT_CONTAIN：不包含，referer大小写不敏感
  // REGEX：匹配正则，referer大小写不敏感
  // NOT_REGEX：正则不匹配，referer大小写不敏感
  // START_WITH：开头是，referer大小写不敏感
  // END_WITH：结尾是，referer大小写不敏感
  // WILDCARD：通配符匹配，referer大小写不敏感，*代表零个或多个任意字符，?代表任意单个字符
  // NOT_WILDCARD：通配符不匹配，referer大小写不敏感，*代表零个或多个任意字符，?代表任意单个字符"}
  MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
  // {'en':'Referer.
  // When the match type is REGEX/NOT_REGEX, only one value is allowed. 
  // Example: http://test.com.', 'zh_CN':'Referer。
  // 当匹配类型为正则/正则不匹配，则只允许只有一个值。
  // 示例：http://test.com。'}
  Referer []*string `json:"referer,omitempty" xml:"referer,omitempty" require:"true" type:"Repeated"`
}

func (s CreateSharedWhitelistRuleRefererCondition) String() string {
  return tea.Prettify(s)
}

func (s CreateSharedWhitelistRuleRefererCondition) GoString() string {
  return s.String()
}

func (s *CreateSharedWhitelistRuleRefererCondition) SetMatchType(v string) *CreateSharedWhitelistRuleRefererCondition {
  s.MatchType = &v
  return s
}

func (s *CreateSharedWhitelistRuleRefererCondition) SetReferer(v []*string) *CreateSharedWhitelistRuleRefererCondition {
  s.Referer = v
  return s
}

type CreateSharedWhitelistRuleHeaderCondition struct {
  // {"en":"Match type.
  // EQUAL: Equals, request header values case sensitive
  // NOT_EQUAL: Does not equal, request header values case sensitive
  // CONTAIN: Contains, request header values case insensitive
  // NOT_CONTAIN: Does not Contains, request header values case insensitive
  // REGEX: Regex match, request header values case insensitive
  // NOT_REGEX: Regular does not match, request header values case insensitive
  // START_WITH: Starts with, request header values case insensitive
  // END_WITH: Ends with, request header values case insensitive
  // WILDCARD: Wildcard matches, request header values case insensitive, ** represents zero or more arbitrary characters, ? represents any single character
  // NOT_WILDCARD: Wildcard does not match, request header values case insensitive, ** represents zero or more arbitrary characters, ? represents any single character", "zh_CN":"匹配类型。
  // EQUAL：等于，头部值大小写敏感
  // NOT_EQUAL：不等于，头部值大小写敏感
  // CONTAIN：包含，头部值大小写不敏感
  // NOT_CONTAIN：不包含，头部值大小写不敏感
  // REGEX：匹配正则，头部值大小写不敏感
  // NOT_REGEX：正则不匹配，头部值大小写不敏感
  // START_WITH：开头是，头部值大小写不敏感
  // END_WITH：结尾是，头部值大小写不敏感
  // WILDCARD：通配符匹配，头部值大小写不敏感，*代表零个或多个任意字符，?代表任意单个字符
  // NOT_WILDCARD：通配符不匹配，头部值大小写不敏感，*代表零个或多个任意字符，?代表任意单个字符"}
  MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
  // {'en':'Header name,case insensitive,up to 100 characters.
  // Example: Accept.', 'zh_CN':'头部名称，大小写不敏感，最多100个字符。
  // 示例：Accept。'}
  Key *string `json:"key,omitempty" xml:"key,omitempty" require:"true"`
  // {'en':'Header value.
  // When the match type is REGEX/NOT_REGEX, only one value is allowed.', 'zh_CN':'头部值。
  // 当匹配类型为正则/正则不匹配，则只允许只有一个值。'}
  ValueList []*string `json:"valueList,omitempty" xml:"valueList,omitempty" require:"true" type:"Repeated"`
}

func (s CreateSharedWhitelistRuleHeaderCondition) String() string {
  return tea.Prettify(s)
}

func (s CreateSharedWhitelistRuleHeaderCondition) GoString() string {
  return s.String()
}

func (s *CreateSharedWhitelistRuleHeaderCondition) SetMatchType(v string) *CreateSharedWhitelistRuleHeaderCondition {
  s.MatchType = &v
  return s
}

func (s *CreateSharedWhitelistRuleHeaderCondition) SetKey(v string) *CreateSharedWhitelistRuleHeaderCondition {
  s.Key = &v
  return s
}

func (s *CreateSharedWhitelistRuleHeaderCondition) SetValueList(v []*string) *CreateSharedWhitelistRuleHeaderCondition {
  s.ValueList = v
  return s
}

type CreateSharedWhitelistRuleResponse struct {
  // {'en':'Please refer to the error code for exceptions.', 'zh_CN':'请参照错误码。','dictionary':'belong=WAAP-MS-Ext|dict=waap_retCodeEnum'}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Description.", "zh_CN":"描述信息。"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"Rule ID.", "zh_CN":"规则ID。"}
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

type CreateSharedWhitelistRuleRequestHeader struct {
  // {"zh_CN":"安全服务类型。有使用多个不同的安全服务时，需要填写具体的服务类型。","en":"Security service type. Please enter a specific service type, if you purchase multiple security services.","dictionary":"belong=WAAP-MS-Ext|dict=waap_serviceType"}
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

type CreateSharedWhitelistRuleResponseHeader struct {
}

func (s CreateSharedWhitelistRuleResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s CreateSharedWhitelistRuleResponseHeader) GoString() string {
  return s.String()
}




