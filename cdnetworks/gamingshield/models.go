package gamingshield

// This file is auto-generated, don't edit it. Thanks.
import (
  "github.com/alibabacloud-go/tea/tea"
)

type UpdateDistributionRuleRequest struct {
  // {"en":"Website Group ID","zh_CN":"网站分组ID"}
  GroupId *string `json:"groupId,omitempty" xml:"groupId,omitempty" require:"true"`
  // {"en":"Switch.\nON: On\nOFF: Off","zh_CN":"开关。\nON：开\nOFF：关"}
  RuleSwitch *string `json:"ruleSwitch,omitempty" xml:"ruleSwitch,omitempty" require:"true"`
  // {"en":"Request uri,when fromType=SPECIFY, groupRuleFromList required.\nALL: All requests\nSPECIFY: Specified uri","zh_CN":"请求uri，当fromType=SPECIFY时，groupRuleFromList必填。\nALL：所有请求\nSPECIFY：指定uri"}
  FromType *string `json:"fromType,omitempty" xml:"fromType,omitempty" require:"true"`
  // {"en":"Grouping rule request uri,when matchType=EQUAL/STARTWITH, content must start with/.","zh_CN":"分组规则请求uri，当matchType=EQUAL/START_WITH时，content必须以/开头。"}
  GroupRuleFromList []*UpdateDistributionRuleRequestGroupRuleFromList `json:"groupRuleFromList,omitempty" xml:"groupRuleFromList,omitempty" type:"Repeated"`
  // {"en":"Target uri type,when from Type=SPECIY, toContent required and it must begin with /,when toType=INDEX/SAME, toContent is empty.\nINDEX: Home page\nSPECIFY: Specified uri\nSAME: Inherited request uri","zh_CN":"目标uri类型，toType=SPECIFY时，toContent必填且必须以/开头，当toType=INDEX/SAME时，content为空。\nINDEX：首页\nSPECIFY：指定uri\nSAME：继承请求路"}
  ToType *string `json:"toType,omitempty" xml:"toType,omitempty" require:"true"`
  // {"en":"Target uri,only supports characters [a-zA-Z0-9. - _/=?:&].","zh_CN":"目标uri，仅支持字符 [a-zA-Z0-9.-_/=?:&]。"}
  ToContent *string `json:"toContent,omitempty" xml:"toContent,omitempty"`
  // {"en":"Whether to carry a query string.\nTRUE: Yes\nFALSE: No","zh_CN":"是否携带查询串。\nTRUE：是\nFALSE：否"}
  WithQs *string `json:"withQs,omitempty" xml:"withQs,omitempty" require:"true"`
}

func (s UpdateDistributionRuleRequest) String() string {
  return tea.Prettify(s)
}

func (s UpdateDistributionRuleRequest) GoString() string {
  return s.String()
}

func (s *UpdateDistributionRuleRequest) SetGroupId(v string) *UpdateDistributionRuleRequest {
  s.GroupId = &v
  return s
}

func (s *UpdateDistributionRuleRequest) SetRuleSwitch(v string) *UpdateDistributionRuleRequest {
  s.RuleSwitch = &v
  return s
}

func (s *UpdateDistributionRuleRequest) SetFromType(v string) *UpdateDistributionRuleRequest {
  s.FromType = &v
  return s
}

func (s *UpdateDistributionRuleRequest) SetGroupRuleFromList(v []*UpdateDistributionRuleRequestGroupRuleFromList) *UpdateDistributionRuleRequest {
  s.GroupRuleFromList = v
  return s
}

func (s *UpdateDistributionRuleRequest) SetToType(v string) *UpdateDistributionRuleRequest {
  s.ToType = &v
  return s
}

func (s *UpdateDistributionRuleRequest) SetToContent(v string) *UpdateDistributionRuleRequest {
  s.ToContent = &v
  return s
}

func (s *UpdateDistributionRuleRequest) SetWithQs(v string) *UpdateDistributionRuleRequest {
  s.WithQs = &v
  return s
}

type UpdateDistributionRuleRequestGroupRuleFromList struct     {
  // {"en":"Request uri match type.\nEQUAL: equal to\nCONTAIN: contains\nSTART_WITH: starts with,\nEND_WITH: ends with","zh_CN":"请求uri匹配类型。\nEQUAL：等于\nCONTAIN：包含\nSTART_WITH：开头是\nEND_WITH：结尾是"}
  MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty"`
  // {"en":"Request uri,only supports characters [a-zA-Z0-9. - _/=?:&].","zh_CN":"请求uri，仅支持字符 [a-zA-Z0-9.-_/=?:&]。"}
  Content *string `json:"content,omitempty" xml:"content,omitempty"`
}

func (s UpdateDistributionRuleRequestGroupRuleFromList) String() string {
  return tea.Prettify(s)
}

func (s UpdateDistributionRuleRequestGroupRuleFromList) GoString() string {
  return s.String()
}

func (s *UpdateDistributionRuleRequestGroupRuleFromList) SetMatchType(v string) *UpdateDistributionRuleRequestGroupRuleFromList {
  s.MatchType = &v
  return s
}

func (s *UpdateDistributionRuleRequestGroupRuleFromList) SetContent(v string) *UpdateDistributionRuleRequestGroupRuleFromList {
  s.Content = &v
  return s
}

type UpdateDistributionRuleRequestHeader struct {
  // {"en":"Security service type. Please enter a specific service type, if you purchase multiple security services.","zh_CN":"安全服务类型。有使用多个不同的安全服务时，需要填写具体的服务类型。"}
  ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty"`
}

func (s UpdateDistributionRuleRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s UpdateDistributionRuleRequestHeader) GoString() string {
  return s.String()
}

func (s *UpdateDistributionRuleRequestHeader) SetServiceType(v string) *UpdateDistributionRuleRequestHeader {
  s.ServiceType = &v
  return s
}

type UpdateDistributionRulePaths struct {
}

func (s UpdateDistributionRulePaths) String() string {
  return tea.Prettify(s)
}

func (s UpdateDistributionRulePaths) GoString() string {
  return s.String()
}

type UpdateDistributionRuleParameters struct {
}

func (s UpdateDistributionRuleParameters) String() string {
  return tea.Prettify(s)
}

func (s UpdateDistributionRuleParameters) GoString() string {
  return s.String()
}

type UpdateDistributionRuleResponse struct {
  // {"dictionary":"belong=WAAP-MS-Ext|dict=waap_retCodeEnum","en":"Please refer to the error code for exceptions.","zh_CN":"请参照错误码。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Description.","zh_CN":"描述信息。"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
}

func (s UpdateDistributionRuleResponse) String() string {
  return tea.Prettify(s)
}

func (s UpdateDistributionRuleResponse) GoString() string {
  return s.String()
}

func (s *UpdateDistributionRuleResponse) SetCode(v string) *UpdateDistributionRuleResponse {
  s.Code = &v
  return s
}

func (s *UpdateDistributionRuleResponse) SetMsg(v string) *UpdateDistributionRuleResponse {
  s.Msg = &v
  return s
}

type UpdateDistributionRuleResponseHeader struct {
}

func (s UpdateDistributionRuleResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s UpdateDistributionRuleResponseHeader) GoString() string {
  return s.String()
}




type CreateWebsiteGroupRequest struct {
  // {"en":"Name.","zh_CN":"名称。"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"Remark.","zh_CN":"备注。"}
  Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
  // {"en":"Port configuration.","zh_CN":"端口配置。"}
  PortRelList []*CreateWebsiteGroupRequestPortRelList `json:"portRelList,omitempty" xml:"portRelList,omitempty" require:"true" type:"Repeated"`
  // {"en":"List of Associate hostname.","zh_CN":"关联域名列表。"}
  DomainInfoList []*CreateWebsiteGroupRequestDomainInfoList `json:"domainInfoList,omitempty" xml:"domainInfoList,omitempty" require:"true" type:"Repeated"`
}

func (s CreateWebsiteGroupRequest) String() string {
  return tea.Prettify(s)
}

func (s CreateWebsiteGroupRequest) GoString() string {
  return s.String()
}

func (s *CreateWebsiteGroupRequest) SetName(v string) *CreateWebsiteGroupRequest {
  s.Name = &v
  return s
}

func (s *CreateWebsiteGroupRequest) SetRemark(v string) *CreateWebsiteGroupRequest {
  s.Remark = &v
  return s
}

func (s *CreateWebsiteGroupRequest) SetPortRelList(v []*CreateWebsiteGroupRequestPortRelList) *CreateWebsiteGroupRequest {
  s.PortRelList = v
  return s
}

func (s *CreateWebsiteGroupRequest) SetDomainInfoList(v []*CreateWebsiteGroupRequestDomainInfoList) *CreateWebsiteGroupRequest {
  s.DomainInfoList = v
  return s
}

type CreateWebsiteGroupRequestPortRelList struct     {
  // {"en":"Port list.\n2000~2099:Except for 2012,2080\n5000~5099\n8200~8999:Except for 8245\n9000~9999:Except for 9119,9331,9518,9851,9852,9996\n15000~15999\n17000~17999\n18000~18099\n50000~50199","zh_CN":"端口列表。\n2000~2099：除2012,2080\n5000~5099\n8200~8999：除8245\n9000~9999：除9119，9331，9518，9851，9852，9996\n15000~15999\n17000~17999\n18000~18099\n50000~50199"}
  PortList []*int `json:"portList,omitempty" xml:"portList,omitempty" require:"true" type:"Repeated"`
  // {"en":"Protocol.\nHTTP: HTTP protocol\nHTTPS: HTTPS protocol","zh_CN":"协议。\nHTTP：HTTP协议\nHTTPS：HTTPS协议"}
  Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty" require:"true"`
}

func (s CreateWebsiteGroupRequestPortRelList) String() string {
  return tea.Prettify(s)
}

func (s CreateWebsiteGroupRequestPortRelList) GoString() string {
  return s.String()
}

func (s *CreateWebsiteGroupRequestPortRelList) SetPortList(v []*int) *CreateWebsiteGroupRequestPortRelList {
  s.PortList = v
  return s
}

func (s *CreateWebsiteGroupRequestPortRelList) SetProtocol(v string) *CreateWebsiteGroupRequestPortRelList {
  s.Protocol = &v
  return s
}

type CreateWebsiteGroupRequestDomainInfoList struct     {
  // {"en":"Hostname","zh_CN":"域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {"en":"Serial number. The smaller the number, the higher the priority; will be automatically converted to a consecutive serial number.","zh_CN":"序号。越小越靠前，会自动转成连续的序号"}
  Index *int `json:"index,omitempty" xml:"index,omitempty" require:"true"`
}

func (s CreateWebsiteGroupRequestDomainInfoList) String() string {
  return tea.Prettify(s)
}

func (s CreateWebsiteGroupRequestDomainInfoList) GoString() string {
  return s.String()
}

func (s *CreateWebsiteGroupRequestDomainInfoList) SetDomain(v string) *CreateWebsiteGroupRequestDomainInfoList {
  s.Domain = &v
  return s
}

func (s *CreateWebsiteGroupRequestDomainInfoList) SetIndex(v int) *CreateWebsiteGroupRequestDomainInfoList {
  s.Index = &v
  return s
}

type CreateWebsiteGroupRequestHeader struct {
  // {"en":"Security service type. Please enter a specific service type, if you purchase multiple security services.","zh_CN":"安全服务类型。有使用多个不同的安全服务时，需要填写具体的服务类型。"}
  ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty"`
}

func (s CreateWebsiteGroupRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s CreateWebsiteGroupRequestHeader) GoString() string {
  return s.String()
}

func (s *CreateWebsiteGroupRequestHeader) SetServiceType(v string) *CreateWebsiteGroupRequestHeader {
  s.ServiceType = &v
  return s
}

type CreateWebsiteGroupPaths struct {
}

func (s CreateWebsiteGroupPaths) String() string {
  return tea.Prettify(s)
}

func (s CreateWebsiteGroupPaths) GoString() string {
  return s.String()
}

type CreateWebsiteGroupParameters struct {
}

func (s CreateWebsiteGroupParameters) String() string {
  return tea.Prettify(s)
}

func (s CreateWebsiteGroupParameters) GoString() string {
  return s.String()
}

type CreateWebsiteGroupResponse struct {
  // {"dictionary":"belong=WAAP-MS-Ext|dict=waap_retCodeEnum","en":"Please refer to the error code for exceptions.","zh_CN":"请参照错误码。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Description.","zh_CN":"描述信息。"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"Data.","zh_CN":"出参数据。"}
  Data *string `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s CreateWebsiteGroupResponse) String() string {
  return tea.Prettify(s)
}

func (s CreateWebsiteGroupResponse) GoString() string {
  return s.String()
}

func (s *CreateWebsiteGroupResponse) SetCode(v string) *CreateWebsiteGroupResponse {
  s.Code = &v
  return s
}

func (s *CreateWebsiteGroupResponse) SetMsg(v string) *CreateWebsiteGroupResponse {
  s.Msg = &v
  return s
}

func (s *CreateWebsiteGroupResponse) SetData(v string) *CreateWebsiteGroupResponse {
  s.Data = &v
  return s
}

type CreateWebsiteGroupResponseHeader struct {
}

func (s CreateWebsiteGroupResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s CreateWebsiteGroupResponseHeader) GoString() string {
  return s.String()
}




type ListGroupDetailsRequest struct {
  // {"en":"ID List.","zh_CN":"ID列表。"}
  IdList []*string `json:"idList,omitempty" xml:"idList,omitempty" type:"Repeated"`
  // {"en":"Name list,supports fuzzy queries.","zh_CN":"名称列表，支持模糊查询。"}
  NameList []*string `json:"nameList,omitempty" xml:"nameList,omitempty" type:"Repeated"`
  // {"en":"Hostname list.","zh_CN":"域名列表。"}
  DomainList []*string `json:"domainList,omitempty" xml:"domainList,omitempty" type:"Repeated"`
}

func (s ListGroupDetailsRequest) String() string {
  return tea.Prettify(s)
}

func (s ListGroupDetailsRequest) GoString() string {
  return s.String()
}

func (s *ListGroupDetailsRequest) SetIdList(v []*string) *ListGroupDetailsRequest {
  s.IdList = v
  return s
}

func (s *ListGroupDetailsRequest) SetNameList(v []*string) *ListGroupDetailsRequest {
  s.NameList = v
  return s
}

func (s *ListGroupDetailsRequest) SetDomainList(v []*string) *ListGroupDetailsRequest {
  s.DomainList = v
  return s
}

type ListGroupDetailsRequestHeader struct {
  // {"en":"Security service type. Please enter a specific service type, if you purchase multiple security services.","zh_CN":"安全服务类型。有使用多个不同的安全服务时，需要填写具体的服务类型。"}
  ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty"`
}

func (s ListGroupDetailsRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ListGroupDetailsRequestHeader) GoString() string {
  return s.String()
}

func (s *ListGroupDetailsRequestHeader) SetServiceType(v string) *ListGroupDetailsRequestHeader {
  s.ServiceType = &v
  return s
}

type ListGroupDetailsPaths struct {
}

func (s ListGroupDetailsPaths) String() string {
  return tea.Prettify(s)
}

func (s ListGroupDetailsPaths) GoString() string {
  return s.String()
}

type ListGroupDetailsParameters struct {
}

func (s ListGroupDetailsParameters) String() string {
  return tea.Prettify(s)
}

func (s ListGroupDetailsParameters) GoString() string {
  return s.String()
}

type ListGroupDetailsResponse struct {
  // {"dictionary":"belong=WAAP-MS-Ext|dict=waap_retCodeEnum","en":"Please refer to the error code for exceptions.","zh_CN":"请参照错误码。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Description.","zh_CN":"描述信息。"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"Data.","zh_CN":"出参数据。"}
  Data []*ListGroupDetailsResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s ListGroupDetailsResponse) String() string {
  return tea.Prettify(s)
}

func (s ListGroupDetailsResponse) GoString() string {
  return s.String()
}

func (s *ListGroupDetailsResponse) SetCode(v string) *ListGroupDetailsResponse {
  s.Code = &v
  return s
}

func (s *ListGroupDetailsResponse) SetMsg(v string) *ListGroupDetailsResponse {
  s.Msg = &v
  return s
}

func (s *ListGroupDetailsResponse) SetData(v []*ListGroupDetailsResponseData) *ListGroupDetailsResponse {
  s.Data = v
  return s
}

type ListGroupDetailsResponseData struct     {
  // {"en":"ID.","zh_CN":"ID。"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {"en":"Name.","zh_CN":"名称。"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"Remark.","zh_CN":"备注。"}
  Remark *string `json:"remark,omitempty" xml:"remark,omitempty" require:"true"`
  // {"en":"Port configuration.","zh_CN":"端口配置。"}
  PortRelList []*ListGroupDetailsResponseDataPortRelList `json:"portRelList,omitempty" xml:"portRelList,omitempty" require:"true" type:"Repeated"`
  // {"en":"List of Associate hostname.","zh_CN":"关联域名列表。"}
  DomainInfoList []*ListGroupDetailsResponseDataDomainInfoList `json:"domainInfoList,omitempty" xml:"domainInfoList,omitempty" require:"true" type:"Repeated"`
  // {"en":"Whether there is an entry hostname associated with it.","zh_CN":"是否关联有入口域名。"}
  RelEntryDomain *bool `json:"relEntryDomain,omitempty" xml:"relEntryDomain,omitempty" require:"true"`
}

func (s ListGroupDetailsResponseData) String() string {
  return tea.Prettify(s)
}

func (s ListGroupDetailsResponseData) GoString() string {
  return s.String()
}

func (s *ListGroupDetailsResponseData) SetId(v string) *ListGroupDetailsResponseData {
  s.Id = &v
  return s
}

func (s *ListGroupDetailsResponseData) SetName(v string) *ListGroupDetailsResponseData {
  s.Name = &v
  return s
}

func (s *ListGroupDetailsResponseData) SetRemark(v string) *ListGroupDetailsResponseData {
  s.Remark = &v
  return s
}

func (s *ListGroupDetailsResponseData) SetPortRelList(v []*ListGroupDetailsResponseDataPortRelList) *ListGroupDetailsResponseData {
  s.PortRelList = v
  return s
}

func (s *ListGroupDetailsResponseData) SetDomainInfoList(v []*ListGroupDetailsResponseDataDomainInfoList) *ListGroupDetailsResponseData {
  s.DomainInfoList = v
  return s
}

func (s *ListGroupDetailsResponseData) SetRelEntryDomain(v bool) *ListGroupDetailsResponseData {
  s.RelEntryDomain = &v
  return s
}

type ListGroupDetailsResponseDataPortRelList struct     {
  // {"en":"Port list.","zh_CN":"端口列表。"}
  PortList []*int `json:"portList,omitempty" xml:"portList,omitempty" require:"true" type:"Repeated"`
  // {"en":"Protocol.\nHTTP: HTTP protocol\nHTTPS: HTTPS protocol","zh_CN":"协议。\nHTTP：HTTP协议\nHTTPS：HTTPS协议"}
  Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty" require:"true"`
}

func (s ListGroupDetailsResponseDataPortRelList) String() string {
  return tea.Prettify(s)
}

func (s ListGroupDetailsResponseDataPortRelList) GoString() string {
  return s.String()
}

func (s *ListGroupDetailsResponseDataPortRelList) SetPortList(v []*int) *ListGroupDetailsResponseDataPortRelList {
  s.PortList = v
  return s
}

func (s *ListGroupDetailsResponseDataPortRelList) SetProtocol(v string) *ListGroupDetailsResponseDataPortRelList {
  s.Protocol = &v
  return s
}

type ListGroupDetailsResponseDataDomainInfoList struct     {
  // {"en":"Hostname","zh_CN":"域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {"en":"Serial number","zh_CN":"序号"}
  Index *int `json:"index,omitempty" xml:"index,omitempty" require:"true"`
}

func (s ListGroupDetailsResponseDataDomainInfoList) String() string {
  return tea.Prettify(s)
}

func (s ListGroupDetailsResponseDataDomainInfoList) GoString() string {
  return s.String()
}

func (s *ListGroupDetailsResponseDataDomainInfoList) SetDomain(v string) *ListGroupDetailsResponseDataDomainInfoList {
  s.Domain = &v
  return s
}

func (s *ListGroupDetailsResponseDataDomainInfoList) SetIndex(v int) *ListGroupDetailsResponseDataDomainInfoList {
  s.Index = &v
  return s
}

type ListGroupDetailsResponseHeader struct {
}

func (s ListGroupDetailsResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ListGroupDetailsResponseHeader) GoString() string {
  return s.String()
}




type ListNavigationRulesRequest struct {
  // {"en":"Website group ID list.","zh_CN":"网站分组ID列表。"}
  GroupIdList []*string `json:"groupIdList,omitempty" xml:"groupIdList,omitempty" type:"Repeated"`
  // {"en":"Website group name list,supports fuzzy queries.","zh_CN":"网站分组名称列表，支持模糊查询。"}
  GroupNameList []*string `json:"groupNameList,omitempty" xml:"groupNameList,omitempty" type:"Repeated"`
}

func (s ListNavigationRulesRequest) String() string {
  return tea.Prettify(s)
}

func (s ListNavigationRulesRequest) GoString() string {
  return s.String()
}

func (s *ListNavigationRulesRequest) SetGroupIdList(v []*string) *ListNavigationRulesRequest {
  s.GroupIdList = v
  return s
}

func (s *ListNavigationRulesRequest) SetGroupNameList(v []*string) *ListNavigationRulesRequest {
  s.GroupNameList = v
  return s
}

type ListNavigationRulesRequestHeader struct {
  // {"en":"Security service type. Please enter a specific service type, if you purchase multiple security services.","zh_CN":"安全服务类型。有使用多个不同的安全服务时，需要填写具体的服务类型。"}
  ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty"`
}

func (s ListNavigationRulesRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ListNavigationRulesRequestHeader) GoString() string {
  return s.String()
}

func (s *ListNavigationRulesRequestHeader) SetServiceType(v string) *ListNavigationRulesRequestHeader {
  s.ServiceType = &v
  return s
}

type ListNavigationRulesPaths struct {
}

func (s ListNavigationRulesPaths) String() string {
  return tea.Prettify(s)
}

func (s ListNavigationRulesPaths) GoString() string {
  return s.String()
}

type ListNavigationRulesParameters struct {
}

func (s ListNavigationRulesParameters) String() string {
  return tea.Prettify(s)
}

func (s ListNavigationRulesParameters) GoString() string {
  return s.String()
}

type ListNavigationRulesResponse struct {
  // {"dictionary":"belong=WAAP-MS-Ext|dict=waap_retCodeEnum","en":"Please refer to the error code for exceptions.","zh_CN":"请参照错误码。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Description.","zh_CN":"描述信息。"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"Data.","zh_CN":"出参数据。"}
  Data []*ListNavigationRulesResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s ListNavigationRulesResponse) String() string {
  return tea.Prettify(s)
}

func (s ListNavigationRulesResponse) GoString() string {
  return s.String()
}

func (s *ListNavigationRulesResponse) SetCode(v string) *ListNavigationRulesResponse {
  s.Code = &v
  return s
}

func (s *ListNavigationRulesResponse) SetMsg(v string) *ListNavigationRulesResponse {
  s.Msg = &v
  return s
}

func (s *ListNavigationRulesResponse) SetData(v []*ListNavigationRulesResponseData) *ListNavigationRulesResponse {
  s.Data = v
  return s
}

type ListNavigationRulesResponseData struct     {
  // {"en":"ID.","zh_CN":"ID。"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {"en":"The website group ID.","zh_CN":"网站分组ID。"}
  GroupId *string `json:"groupId,omitempty" xml:"groupId,omitempty" require:"true"`
  // {"en":"Group name.","zh_CN":"分组名称。"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"Switch.\nON: On\nOFF: Off","zh_CN":"开关。\nON：开\nOFF：关"}
  RuleSwitch *string `json:"ruleSwitch,omitempty" xml:"ruleSwitch,omitempty" require:"true"`
  // {"en":"Request uri type.\nALL: All requests\nSPECIFY: Specified uri","zh_CN":"请求uri类型。\nALL：所有请求\nSPECIFY：指定uri"}
  FromType *string `json:"fromType,omitempty" xml:"fromType,omitempty" require:"true"`
  // {"en":"Grouping rule request uri.","zh_CN":"分组规则请求uri。"}
  GroupRuleFromList []*ListNavigationRulesResponseDataGroupRuleFromList `json:"groupRuleFromList,omitempty" xml:"groupRuleFromList,omitempty" require:"true" type:"Repeated"`
  // {"en":"Target uri type.\nINDEX: Home page\nSPECIFY: Specified uri\nSAME: Inherited request uri","zh_CN":"目标uri类型。\nINDEX：首页\nSPECIFY：指定uri\nSAME：继承请求uri"}
  ToType *string `json:"toType,omitempty" xml:"toType,omitempty" require:"true"`
  // {"en":"Target uri.","zh_CN":"目标uri。"}
  ToContent *string `json:"toContent,omitempty" xml:"toContent,omitempty" require:"true"`
  // {"en":"Whether to carry a query string.\nTRUE: Yes\nFALSE: No","zh_CN":"是否携带查询串。\nTRUE：是\nFALSE：否"}
  WithQs *string `json:"withQs,omitempty" xml:"withQs,omitempty" require:"true"`
}

func (s ListNavigationRulesResponseData) String() string {
  return tea.Prettify(s)
}

func (s ListNavigationRulesResponseData) GoString() string {
  return s.String()
}

func (s *ListNavigationRulesResponseData) SetId(v string) *ListNavigationRulesResponseData {
  s.Id = &v
  return s
}

func (s *ListNavigationRulesResponseData) SetGroupId(v string) *ListNavigationRulesResponseData {
  s.GroupId = &v
  return s
}

func (s *ListNavigationRulesResponseData) SetName(v string) *ListNavigationRulesResponseData {
  s.Name = &v
  return s
}

func (s *ListNavigationRulesResponseData) SetRuleSwitch(v string) *ListNavigationRulesResponseData {
  s.RuleSwitch = &v
  return s
}

func (s *ListNavigationRulesResponseData) SetFromType(v string) *ListNavigationRulesResponseData {
  s.FromType = &v
  return s
}

func (s *ListNavigationRulesResponseData) SetGroupRuleFromList(v []*ListNavigationRulesResponseDataGroupRuleFromList) *ListNavigationRulesResponseData {
  s.GroupRuleFromList = v
  return s
}

func (s *ListNavigationRulesResponseData) SetToType(v string) *ListNavigationRulesResponseData {
  s.ToType = &v
  return s
}

func (s *ListNavigationRulesResponseData) SetToContent(v string) *ListNavigationRulesResponseData {
  s.ToContent = &v
  return s
}

func (s *ListNavigationRulesResponseData) SetWithQs(v string) *ListNavigationRulesResponseData {
  s.WithQs = &v
  return s
}

type ListNavigationRulesResponseDataGroupRuleFromList struct     {
  // {"en":"Request uri match type.\nEQUAL: equal to\nCONTAIN: contains\nSTART_WITH: starts with\nEND_WITH: ends with","zh_CN":"请求uri匹配类型。\nEQUAL：等于\nCONTAIN：包含\nSTART_WITH：开头是\nEND_WITH：结尾是"}
  MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
  // {"en":"Request uri.","zh_CN":"请求uri。"}
  Content *string `json:"content,omitempty" xml:"content,omitempty" require:"true"`
}

func (s ListNavigationRulesResponseDataGroupRuleFromList) String() string {
  return tea.Prettify(s)
}

func (s ListNavigationRulesResponseDataGroupRuleFromList) GoString() string {
  return s.String()
}

func (s *ListNavigationRulesResponseDataGroupRuleFromList) SetMatchType(v string) *ListNavigationRulesResponseDataGroupRuleFromList {
  s.MatchType = &v
  return s
}

func (s *ListNavigationRulesResponseDataGroupRuleFromList) SetContent(v string) *ListNavigationRulesResponseDataGroupRuleFromList {
  s.Content = &v
  return s
}

type ListNavigationRulesResponseHeader struct {
}

func (s ListNavigationRulesResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ListNavigationRulesResponseHeader) GoString() string {
  return s.String()
}




type UpdateEntryDomainRequest struct {
  // {"en":"ID.","zh_CN":"ID。"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {"en":"Website group ID","zh_CN":"网站分组ID"}
  GroupId *string `json:"groupId,omitempty" xml:"groupId,omitempty" require:"true"`
  // {"en":"Port configuration.","zh_CN":"端口配置。"}
  PortRelList []*UpdateEntryDomainRequestPortRelList `json:"portRelList,omitempty" xml:"portRelList,omitempty" require:"true" type:"Repeated"`
}

func (s UpdateEntryDomainRequest) String() string {
  return tea.Prettify(s)
}

func (s UpdateEntryDomainRequest) GoString() string {
  return s.String()
}

func (s *UpdateEntryDomainRequest) SetId(v string) *UpdateEntryDomainRequest {
  s.Id = &v
  return s
}

func (s *UpdateEntryDomainRequest) SetGroupId(v string) *UpdateEntryDomainRequest {
  s.GroupId = &v
  return s
}

func (s *UpdateEntryDomainRequest) SetPortRelList(v []*UpdateEntryDomainRequestPortRelList) *UpdateEntryDomainRequest {
  s.PortRelList = v
  return s
}

type UpdateEntryDomainRequestPortRelList struct     {
  // {"en":"Port list.\n2000~2099:Except for 2012,2080\n5000~5099\n8200~8999:Except for 8245\n9000~9999:Except for 9119,9331,9518,9851,9852,9996\n15000~15999\n17000~17999\n18000~18099\n50000~50199","zh_CN":"端口列表。\n2000~2099：除2012,2080\n5000~5099\n8200~8999：除8245\n9000~9999：除9119，9331，9518，9851，9852，9996\n15000~15999\n17000~17999\n18000~18099\n50000~50199"}
  PortList []*int `json:"portList,omitempty" xml:"portList,omitempty" require:"true" type:"Repeated"`
  // {"en":"Protocol.\nHTTP: HTTP protocol\nHTTPS: HTTPS protocol","zh_CN":"协议。\nHTTP：HTTP协议\nHTTPS：HTTPS协议"}
  Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty" require:"true"`
}

func (s UpdateEntryDomainRequestPortRelList) String() string {
  return tea.Prettify(s)
}

func (s UpdateEntryDomainRequestPortRelList) GoString() string {
  return s.String()
}

func (s *UpdateEntryDomainRequestPortRelList) SetPortList(v []*int) *UpdateEntryDomainRequestPortRelList {
  s.PortList = v
  return s
}

func (s *UpdateEntryDomainRequestPortRelList) SetProtocol(v string) *UpdateEntryDomainRequestPortRelList {
  s.Protocol = &v
  return s
}

type UpdateEntryDomainRequestHeader struct {
  // {"en":"Security service type. Please enter a specific service type, if you purchase multiple security services.","zh_CN":"安全服务类型。有使用多个不同的安全服务时，需要填写具体的服务类型。"}
  ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty"`
}

func (s UpdateEntryDomainRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s UpdateEntryDomainRequestHeader) GoString() string {
  return s.String()
}

func (s *UpdateEntryDomainRequestHeader) SetServiceType(v string) *UpdateEntryDomainRequestHeader {
  s.ServiceType = &v
  return s
}

type UpdateEntryDomainPaths struct {
}

func (s UpdateEntryDomainPaths) String() string {
  return tea.Prettify(s)
}

func (s UpdateEntryDomainPaths) GoString() string {
  return s.String()
}

type UpdateEntryDomainParameters struct {
}

func (s UpdateEntryDomainParameters) String() string {
  return tea.Prettify(s)
}

func (s UpdateEntryDomainParameters) GoString() string {
  return s.String()
}

type UpdateEntryDomainResponse struct {
  // {"dictionary":"belong=WAAP-MS-Ext|dict=waap_retCodeEnum","en":"Please refer to the error code for exceptions.","zh_CN":"请参照错误码。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Description.","zh_CN":"描述信息。"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
}

func (s UpdateEntryDomainResponse) String() string {
  return tea.Prettify(s)
}

func (s UpdateEntryDomainResponse) GoString() string {
  return s.String()
}

func (s *UpdateEntryDomainResponse) SetCode(v string) *UpdateEntryDomainResponse {
  s.Code = &v
  return s
}

func (s *UpdateEntryDomainResponse) SetMsg(v string) *UpdateEntryDomainResponse {
  s.Msg = &v
  return s
}

type UpdateEntryDomainResponseHeader struct {
}

func (s UpdateEntryDomainResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s UpdateEntryDomainResponseHeader) GoString() string {
  return s.String()
}




type UpdateWebsiteGroupRequest struct {
  // {"en":"ID.","zh_CN":"ID。"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {"en":"Remark.","zh_CN":"备注。"}
  Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
  // {"en":"Port configuration.","zh_CN":"端口配置。"}
  PortRelList []*UpdateWebsiteGroupRequestPortRelList `json:"portRelList,omitempty" xml:"portRelList,omitempty" require:"true" type:"Repeated"`
  // {"en":"List of Associate hostname.","zh_CN":"关联域名列表。"}
  DomainInfoList []*UpdateWebsiteGroupRequestDomainInfoList `json:"domainInfoList,omitempty" xml:"domainInfoList,omitempty" require:"true" type:"Repeated"`
}

func (s UpdateWebsiteGroupRequest) String() string {
  return tea.Prettify(s)
}

func (s UpdateWebsiteGroupRequest) GoString() string {
  return s.String()
}

func (s *UpdateWebsiteGroupRequest) SetId(v string) *UpdateWebsiteGroupRequest {
  s.Id = &v
  return s
}

func (s *UpdateWebsiteGroupRequest) SetRemark(v string) *UpdateWebsiteGroupRequest {
  s.Remark = &v
  return s
}

func (s *UpdateWebsiteGroupRequest) SetPortRelList(v []*UpdateWebsiteGroupRequestPortRelList) *UpdateWebsiteGroupRequest {
  s.PortRelList = v
  return s
}

func (s *UpdateWebsiteGroupRequest) SetDomainInfoList(v []*UpdateWebsiteGroupRequestDomainInfoList) *UpdateWebsiteGroupRequest {
  s.DomainInfoList = v
  return s
}

type UpdateWebsiteGroupRequestPortRelList struct     {
  // {"en":"Port list.\n2000~2099:Except for 2012,2080\n5000~5099\n8200~8999:Except for 8245\n9000~9999:Except for 9119,9331,9518,9851,9852,9996\n15000~15999\n17000~17999\n18000~18099\n50000~50199","zh_CN":"端口列表。\n2000~2099：除2012,2080\n5000~5099\n8200~8999：除8245\n9000~9999：除9119，9331，9518，9851，9852，9996\n15000~15999\n17000~17999\n18000~18099\n50000~50199"}
  PortList []*int `json:"portList,omitempty" xml:"portList,omitempty" require:"true" type:"Repeated"`
  // {"en":"Protocol.\nHTTP: HTTP protocol\nHTTPS: HTTPS protocol","zh_CN":"协议。\nHTTP：HTTP协议\nHTTPS：HTTPS协议"}
  Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty" require:"true"`
}

func (s UpdateWebsiteGroupRequestPortRelList) String() string {
  return tea.Prettify(s)
}

func (s UpdateWebsiteGroupRequestPortRelList) GoString() string {
  return s.String()
}

func (s *UpdateWebsiteGroupRequestPortRelList) SetPortList(v []*int) *UpdateWebsiteGroupRequestPortRelList {
  s.PortList = v
  return s
}

func (s *UpdateWebsiteGroupRequestPortRelList) SetProtocol(v string) *UpdateWebsiteGroupRequestPortRelList {
  s.Protocol = &v
  return s
}

type UpdateWebsiteGroupRequestDomainInfoList struct     {
  // {"en":"Hostname","zh_CN":"域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {"en":"Serial number. The smaller the number, the higher the priority; will be automatically converted to a consecutive serial number.","zh_CN":"序号。越小越靠前，会自动转成连续的序号"}
  Index *int `json:"index,omitempty" xml:"index,omitempty" require:"true"`
}

func (s UpdateWebsiteGroupRequestDomainInfoList) String() string {
  return tea.Prettify(s)
}

func (s UpdateWebsiteGroupRequestDomainInfoList) GoString() string {
  return s.String()
}

func (s *UpdateWebsiteGroupRequestDomainInfoList) SetDomain(v string) *UpdateWebsiteGroupRequestDomainInfoList {
  s.Domain = &v
  return s
}

func (s *UpdateWebsiteGroupRequestDomainInfoList) SetIndex(v int) *UpdateWebsiteGroupRequestDomainInfoList {
  s.Index = &v
  return s
}

type UpdateWebsiteGroupRequestHeader struct {
  // {"en":"Security service type. Please enter a specific service type, if you purchase multiple security services.","zh_CN":"安全服务类型。有使用多个不同的安全服务时，需要填写具体的服务类型。"}
  ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty"`
}

func (s UpdateWebsiteGroupRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s UpdateWebsiteGroupRequestHeader) GoString() string {
  return s.String()
}

func (s *UpdateWebsiteGroupRequestHeader) SetServiceType(v string) *UpdateWebsiteGroupRequestHeader {
  s.ServiceType = &v
  return s
}

type UpdateWebsiteGroupPaths struct {
}

func (s UpdateWebsiteGroupPaths) String() string {
  return tea.Prettify(s)
}

func (s UpdateWebsiteGroupPaths) GoString() string {
  return s.String()
}

type UpdateWebsiteGroupParameters struct {
}

func (s UpdateWebsiteGroupParameters) String() string {
  return tea.Prettify(s)
}

func (s UpdateWebsiteGroupParameters) GoString() string {
  return s.String()
}

type UpdateWebsiteGroupResponse struct {
  // {"dictionary":"belong=WAAP-MS-Ext|dict=waap_retCodeEnum","en":"Please refer to the error code for exceptions.","zh_CN":"请参照错误码。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Description.","zh_CN":"描述信息。"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
}

func (s UpdateWebsiteGroupResponse) String() string {
  return tea.Prettify(s)
}

func (s UpdateWebsiteGroupResponse) GoString() string {
  return s.String()
}

func (s *UpdateWebsiteGroupResponse) SetCode(v string) *UpdateWebsiteGroupResponse {
  s.Code = &v
  return s
}

func (s *UpdateWebsiteGroupResponse) SetMsg(v string) *UpdateWebsiteGroupResponse {
  s.Msg = &v
  return s
}

type UpdateWebsiteGroupResponseHeader struct {
}

func (s UpdateWebsiteGroupResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s UpdateWebsiteGroupResponseHeader) GoString() string {
  return s.String()
}




type DeleteEntryDomainRequest struct {
  // {"en":"ID list.","zh_CN":"ID列表。"}
  IdList []*string `json:"idList,omitempty" xml:"idList,omitempty" require:"true" type:"Repeated"`
}

func (s DeleteEntryDomainRequest) String() string {
  return tea.Prettify(s)
}

func (s DeleteEntryDomainRequest) GoString() string {
  return s.String()
}

func (s *DeleteEntryDomainRequest) SetIdList(v []*string) *DeleteEntryDomainRequest {
  s.IdList = v
  return s
}

type DeleteEntryDomainRequestHeader struct {
  // {"en":"Security service type. Please enter a specific service type, if you purchase multiple security services.","zh_CN":"安全服务类型。有使用多个不同的安全服务时，需要填写具体的服务类型。"}
  ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty"`
}

func (s DeleteEntryDomainRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteEntryDomainRequestHeader) GoString() string {
  return s.String()
}

func (s *DeleteEntryDomainRequestHeader) SetServiceType(v string) *DeleteEntryDomainRequestHeader {
  s.ServiceType = &v
  return s
}

type DeleteEntryDomainPaths struct {
}

func (s DeleteEntryDomainPaths) String() string {
  return tea.Prettify(s)
}

func (s DeleteEntryDomainPaths) GoString() string {
  return s.String()
}

type DeleteEntryDomainParameters struct {
}

func (s DeleteEntryDomainParameters) String() string {
  return tea.Prettify(s)
}

func (s DeleteEntryDomainParameters) GoString() string {
  return s.String()
}

type DeleteEntryDomainResponse struct {
  // {"dictionary":"belong=WAAP-MS-Ext|dict=waap_retCodeEnum","en":"Please refer to the error code for exceptions.","zh_CN":"请参照错误码。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Description.","zh_CN":"描述信息。"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
}

func (s DeleteEntryDomainResponse) String() string {
  return tea.Prettify(s)
}

func (s DeleteEntryDomainResponse) GoString() string {
  return s.String()
}

func (s *DeleteEntryDomainResponse) SetCode(v string) *DeleteEntryDomainResponse {
  s.Code = &v
  return s
}

func (s *DeleteEntryDomainResponse) SetMsg(v string) *DeleteEntryDomainResponse {
  s.Msg = &v
  return s
}

type DeleteEntryDomainResponseHeader struct {
}

func (s DeleteEntryDomainResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteEntryDomainResponseHeader) GoString() string {
  return s.String()
}




type AddEntryDomainRequest struct {
  // {"en":"Hostname.","zh_CN":"域名。"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {"en":"Website group ID","zh_CN":"网站分组ID"}
  GroupId *string `json:"groupId,omitempty" xml:"groupId,omitempty" require:"true"`
  // {"en":"Port configuration.","zh_CN":"端口配置。"}
  PortRelList []*AddEntryDomainRequestPortRelList `json:"portRelList,omitempty" xml:"portRelList,omitempty" require:"true" type:"Repeated"`
}

func (s AddEntryDomainRequest) String() string {
  return tea.Prettify(s)
}

func (s AddEntryDomainRequest) GoString() string {
  return s.String()
}

func (s *AddEntryDomainRequest) SetDomain(v string) *AddEntryDomainRequest {
  s.Domain = &v
  return s
}

func (s *AddEntryDomainRequest) SetGroupId(v string) *AddEntryDomainRequest {
  s.GroupId = &v
  return s
}

func (s *AddEntryDomainRequest) SetPortRelList(v []*AddEntryDomainRequestPortRelList) *AddEntryDomainRequest {
  s.PortRelList = v
  return s
}

type AddEntryDomainRequestPortRelList struct     {
  // {"en":"Port list.\n2000~2099:Except for 2012,2080\n5000~5099\n8200~8999:Except for 8245\n9000~9999:Except for 9119,9331,9518,9851,9852,9996\n15000~15999\n17000~17999\n18000~18099\n50000~50199","zh_CN":"端口列表。\n2000~2099：除2012,2080\n5000~5099\n8200~8999：除8245\n9000~9999：除9119，9331，9518，9851，9852，9996\n15000~15999\n17000~17999\n18000~18099\n50000~50199"}
  PortList []*int `json:"portList,omitempty" xml:"portList,omitempty" require:"true" type:"Repeated"`
  // {"en":"Protocol.\nHTTP: HTTP protocol\nHTTPS: HTTPS protocol","zh_CN":"协议。\nHTTP：HTTP协议\nHTTPS：HTTPS协议"}
  Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty" require:"true"`
}

func (s AddEntryDomainRequestPortRelList) String() string {
  return tea.Prettify(s)
}

func (s AddEntryDomainRequestPortRelList) GoString() string {
  return s.String()
}

func (s *AddEntryDomainRequestPortRelList) SetPortList(v []*int) *AddEntryDomainRequestPortRelList {
  s.PortList = v
  return s
}

func (s *AddEntryDomainRequestPortRelList) SetProtocol(v string) *AddEntryDomainRequestPortRelList {
  s.Protocol = &v
  return s
}

type AddEntryDomainRequestHeader struct {
  // {"en":"Security service type. Please enter a specific service type, if you purchase multiple security services.","zh_CN":"安全服务类型。有使用多个不同的安全服务时，需要填写具体的服务类型。"}
  ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty"`
}

func (s AddEntryDomainRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s AddEntryDomainRequestHeader) GoString() string {
  return s.String()
}

func (s *AddEntryDomainRequestHeader) SetServiceType(v string) *AddEntryDomainRequestHeader {
  s.ServiceType = &v
  return s
}

type AddEntryDomainPaths struct {
}

func (s AddEntryDomainPaths) String() string {
  return tea.Prettify(s)
}

func (s AddEntryDomainPaths) GoString() string {
  return s.String()
}

type AddEntryDomainParameters struct {
}

func (s AddEntryDomainParameters) String() string {
  return tea.Prettify(s)
}

func (s AddEntryDomainParameters) GoString() string {
  return s.String()
}

type AddEntryDomainResponse struct {
  // {"dictionary":"belong=WAAP-MS-Ext|dict=waap_retCodeEnum","en":"Please refer to the error code for exceptions.","zh_CN":"请参照错误码。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Description.","zh_CN":"描述信息。"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"Data.","zh_CN":"出参数据。"}
  Data *string `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s AddEntryDomainResponse) String() string {
  return tea.Prettify(s)
}

func (s AddEntryDomainResponse) GoString() string {
  return s.String()
}

func (s *AddEntryDomainResponse) SetCode(v string) *AddEntryDomainResponse {
  s.Code = &v
  return s
}

func (s *AddEntryDomainResponse) SetMsg(v string) *AddEntryDomainResponse {
  s.Msg = &v
  return s
}

func (s *AddEntryDomainResponse) SetData(v string) *AddEntryDomainResponse {
  s.Data = &v
  return s
}

type AddEntryDomainResponseHeader struct {
}

func (s AddEntryDomainResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s AddEntryDomainResponseHeader) GoString() string {
  return s.String()
}




type DeleteWebsiteGroupRequest struct {
  // {"en":"ID list.","zh_CN":"ID列表。"}
  IdList []*string `json:"idList,omitempty" xml:"idList,omitempty" require:"true" type:"Repeated"`
}

func (s DeleteWebsiteGroupRequest) String() string {
  return tea.Prettify(s)
}

func (s DeleteWebsiteGroupRequest) GoString() string {
  return s.String()
}

func (s *DeleteWebsiteGroupRequest) SetIdList(v []*string) *DeleteWebsiteGroupRequest {
  s.IdList = v
  return s
}

type DeleteWebsiteGroupRequestHeader struct {
  // {"en":"Security service type. Please enter a specific service type, if you purchase multiple security services.","zh_CN":"安全服务类型。有使用多个不同的安全服务时，需要填写具体的服务类型。"}
  ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty"`
}

func (s DeleteWebsiteGroupRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteWebsiteGroupRequestHeader) GoString() string {
  return s.String()
}

func (s *DeleteWebsiteGroupRequestHeader) SetServiceType(v string) *DeleteWebsiteGroupRequestHeader {
  s.ServiceType = &v
  return s
}

type DeleteWebsiteGroupPaths struct {
}

func (s DeleteWebsiteGroupPaths) String() string {
  return tea.Prettify(s)
}

func (s DeleteWebsiteGroupPaths) GoString() string {
  return s.String()
}

type DeleteWebsiteGroupParameters struct {
}

func (s DeleteWebsiteGroupParameters) String() string {
  return tea.Prettify(s)
}

func (s DeleteWebsiteGroupParameters) GoString() string {
  return s.String()
}

type DeleteWebsiteGroupResponse struct {
  // {"dictionary":"belong=WAAP-MS-Ext|dict=waap_retCodeEnum","en":"Please refer to the error code for exceptions.","zh_CN":"请参照错误码。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Description.","zh_CN":"描述信息。"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
}

func (s DeleteWebsiteGroupResponse) String() string {
  return tea.Prettify(s)
}

func (s DeleteWebsiteGroupResponse) GoString() string {
  return s.String()
}

func (s *DeleteWebsiteGroupResponse) SetCode(v string) *DeleteWebsiteGroupResponse {
  s.Code = &v
  return s
}

func (s *DeleteWebsiteGroupResponse) SetMsg(v string) *DeleteWebsiteGroupResponse {
  s.Msg = &v
  return s
}

type DeleteWebsiteGroupResponseHeader struct {
}

func (s DeleteWebsiteGroupResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteWebsiteGroupResponseHeader) GoString() string {
  return s.String()
}




type RestartEntryDomainRequest struct {
  // {"en":"Hostname list.","zh_CN":"域名列表。"}
  DomainList []*string `json:"domainList,omitempty" xml:"domainList,omitempty" require:"true" type:"Repeated"`
}

func (s RestartEntryDomainRequest) String() string {
  return tea.Prettify(s)
}

func (s RestartEntryDomainRequest) GoString() string {
  return s.String()
}

func (s *RestartEntryDomainRequest) SetDomainList(v []*string) *RestartEntryDomainRequest {
  s.DomainList = v
  return s
}

type RestartEntryDomainRequestHeader struct {
  // {"en":"Security service type. Please enter a specific service type, if you purchase multiple security services.","zh_CN":"安全服务类型。有使用多个不同的安全服务时，需要填写具体的服务类型。"}
  ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty"`
}

func (s RestartEntryDomainRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s RestartEntryDomainRequestHeader) GoString() string {
  return s.String()
}

func (s *RestartEntryDomainRequestHeader) SetServiceType(v string) *RestartEntryDomainRequestHeader {
  s.ServiceType = &v
  return s
}

type RestartEntryDomainPaths struct {
}

func (s RestartEntryDomainPaths) String() string {
  return tea.Prettify(s)
}

func (s RestartEntryDomainPaths) GoString() string {
  return s.String()
}

type RestartEntryDomainParameters struct {
}

func (s RestartEntryDomainParameters) String() string {
  return tea.Prettify(s)
}

func (s RestartEntryDomainParameters) GoString() string {
  return s.String()
}

type RestartEntryDomainResponse struct {
  // {"dictionary":"belong=WAAP-MS-Ext|dict=waap_retCodeEnum","en":"Please refer to the error code for exceptions.","zh_CN":"请参照错误码。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Description.","zh_CN":"描述信息。"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
}

func (s RestartEntryDomainResponse) String() string {
  return tea.Prettify(s)
}

func (s RestartEntryDomainResponse) GoString() string {
  return s.String()
}

func (s *RestartEntryDomainResponse) SetCode(v string) *RestartEntryDomainResponse {
  s.Code = &v
  return s
}

func (s *RestartEntryDomainResponse) SetMsg(v string) *RestartEntryDomainResponse {
  s.Msg = &v
  return s
}

type RestartEntryDomainResponseHeader struct {
}

func (s RestartEntryDomainResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s RestartEntryDomainResponseHeader) GoString() string {
  return s.String()
}




type ListEntryDomainsRequest struct {
  // {"en":"ID List","zh_CN":"ID列表"}
  IdList []*string `json:"idList,omitempty" xml:"idList,omitempty" type:"Repeated"`
  // {"en":"List of entry hostname.","zh_CN":"入口域名列表。"}
  DomainList []*string `json:"domainList,omitempty" xml:"domainList,omitempty" type:"Repeated"`
  // {"en":"Website group ID list","zh_CN":"网站分组ID列表"}
  GroupIdList []*string `json:"groupIdList,omitempty" xml:"groupIdList,omitempty" type:"Repeated"`
  // {"en":"Website group name list,supports fuzzy queries.","zh_CN":"网站分组名称列表，支持模糊查询。"}
  GroupNameList []*string `json:"groupNameList,omitempty" xml:"groupNameList,omitempty" type:"Repeated"`
}

func (s ListEntryDomainsRequest) String() string {
  return tea.Prettify(s)
}

func (s ListEntryDomainsRequest) GoString() string {
  return s.String()
}

func (s *ListEntryDomainsRequest) SetIdList(v []*string) *ListEntryDomainsRequest {
  s.IdList = v
  return s
}

func (s *ListEntryDomainsRequest) SetDomainList(v []*string) *ListEntryDomainsRequest {
  s.DomainList = v
  return s
}

func (s *ListEntryDomainsRequest) SetGroupIdList(v []*string) *ListEntryDomainsRequest {
  s.GroupIdList = v
  return s
}

func (s *ListEntryDomainsRequest) SetGroupNameList(v []*string) *ListEntryDomainsRequest {
  s.GroupNameList = v
  return s
}

type ListEntryDomainsRequestHeader struct {
  // {"en":"Security service type. Please enter a specific service type, if you purchase multiple security services.","zh_CN":"安全服务类型。有使用多个不同的安全服务时，需要填写具体的服务类型。"}
  ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty"`
}

func (s ListEntryDomainsRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ListEntryDomainsRequestHeader) GoString() string {
  return s.String()
}

func (s *ListEntryDomainsRequestHeader) SetServiceType(v string) *ListEntryDomainsRequestHeader {
  s.ServiceType = &v
  return s
}

type ListEntryDomainsPaths struct {
}

func (s ListEntryDomainsPaths) String() string {
  return tea.Prettify(s)
}

func (s ListEntryDomainsPaths) GoString() string {
  return s.String()
}

type ListEntryDomainsParameters struct {
}

func (s ListEntryDomainsParameters) String() string {
  return tea.Prettify(s)
}

func (s ListEntryDomainsParameters) GoString() string {
  return s.String()
}

type ListEntryDomainsResponse struct {
  // {"dictionary":"belong=WAAP-MS-Ext|dict=waap_retCodeEnum","en":"Please refer to the error code for exceptions.","zh_CN":"请参照错误码。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Description.","zh_CN":"描述信息。"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"Data.","zh_CN":"出参数据。"}
  Data []*ListEntryDomainsResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s ListEntryDomainsResponse) String() string {
  return tea.Prettify(s)
}

func (s ListEntryDomainsResponse) GoString() string {
  return s.String()
}

func (s *ListEntryDomainsResponse) SetCode(v string) *ListEntryDomainsResponse {
  s.Code = &v
  return s
}

func (s *ListEntryDomainsResponse) SetMsg(v string) *ListEntryDomainsResponse {
  s.Msg = &v
  return s
}

func (s *ListEntryDomainsResponse) SetData(v []*ListEntryDomainsResponseData) *ListEntryDomainsResponse {
  s.Data = v
  return s
}

type ListEntryDomainsResponseData struct     {
  // {"en":"ID.","zh_CN":"ID。"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {"en":"Hostname.","zh_CN":"域名。"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {"en":"Related website grouping information.","zh_CN":"关联网站分组信息。"}
  GroupInfo *ListEntryDomainsResponseDataGroupInfo `json:"groupInfo,omitempty" xml:"groupInfo,omitempty" require:"true" type:"Struct"`
  // {"en":"Port configuration.","zh_CN":"端口配置。"}
  PortRelList []*ListEntryDomainsResponseDataPortRelList `json:"portRelList,omitempty" xml:"portRelList,omitempty" require:"true" type:"Repeated"`
  // {"en":"Status.\nENABLE: Enable\nDISABLE: Disable","zh_CN":"状态。\nENABLE：启用\nDISABLE：禁用"}
  Status *string `json:"status,omitempty" xml:"status,omitempty" require:"true"`
}

func (s ListEntryDomainsResponseData) String() string {
  return tea.Prettify(s)
}

func (s ListEntryDomainsResponseData) GoString() string {
  return s.String()
}

func (s *ListEntryDomainsResponseData) SetId(v string) *ListEntryDomainsResponseData {
  s.Id = &v
  return s
}

func (s *ListEntryDomainsResponseData) SetDomain(v string) *ListEntryDomainsResponseData {
  s.Domain = &v
  return s
}

func (s *ListEntryDomainsResponseData) SetGroupInfo(v *ListEntryDomainsResponseDataGroupInfo) *ListEntryDomainsResponseData {
  s.GroupInfo = v
  return s
}

func (s *ListEntryDomainsResponseData) SetPortRelList(v []*ListEntryDomainsResponseDataPortRelList) *ListEntryDomainsResponseData {
  s.PortRelList = v
  return s
}

func (s *ListEntryDomainsResponseData) SetStatus(v string) *ListEntryDomainsResponseData {
  s.Status = &v
  return s
}

type ListEntryDomainsResponseDataGroupInfo struct {
  // {"en":"Website group ID.","zh_CN":"网站分组ID。"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {"en":"Name.","zh_CN":"名称。"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
}

func (s ListEntryDomainsResponseDataGroupInfo) String() string {
  return tea.Prettify(s)
}

func (s ListEntryDomainsResponseDataGroupInfo) GoString() string {
  return s.String()
}

func (s *ListEntryDomainsResponseDataGroupInfo) SetId(v string) *ListEntryDomainsResponseDataGroupInfo {
  s.Id = &v
  return s
}

func (s *ListEntryDomainsResponseDataGroupInfo) SetName(v string) *ListEntryDomainsResponseDataGroupInfo {
  s.Name = &v
  return s
}

type ListEntryDomainsResponseDataPortRelList struct     {
  // {"en":"Port list.\n2000~2099:Except for 2012,2080\n5000~5099\n8200~8999:Except for 8245\n9000~9999:Except for 9119,9331,9518,9851,9852,9996\n15000~15999\n17000~17999\n18000~18099\n50000~50199","zh_CN":"端口列表。\n2000~2099：除2012,2080\n5000~5099\n8200~8999：除8245\n9000~9999：除9119，9331，9518，9851，9852，9996\n15000~15999\n17000~17999\n18000~18099\n50000~50199"}
  PortList []*int `json:"portList,omitempty" xml:"portList,omitempty" require:"true" type:"Repeated"`
  // {"en":"Protocol.\nHTTP: HTTP protocol\nHTTPS: HTTPS protocol","zh_CN":"协议。\nHTTP：HTTP协议\nHTTPS：HTTPS协议"}
  Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty" require:"true"`
}

func (s ListEntryDomainsResponseDataPortRelList) String() string {
  return tea.Prettify(s)
}

func (s ListEntryDomainsResponseDataPortRelList) GoString() string {
  return s.String()
}

func (s *ListEntryDomainsResponseDataPortRelList) SetPortList(v []*int) *ListEntryDomainsResponseDataPortRelList {
  s.PortList = v
  return s
}

func (s *ListEntryDomainsResponseDataPortRelList) SetProtocol(v string) *ListEntryDomainsResponseDataPortRelList {
  s.Protocol = &v
  return s
}

type ListEntryDomainsResponseHeader struct {
}

func (s ListEntryDomainsResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ListEntryDomainsResponseHeader) GoString() string {
  return s.String()
}




