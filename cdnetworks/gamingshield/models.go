package gamingshield

// This file is auto-generated, don't edit it. Thanks.
import (
  "github.com/alibabacloud-go/tea/tea"
)

type SaveNavigationRuleRequest struct {
  // {"zh_CN":"网站分组ID","en":"Website Group ID"}
  GroupId *string `json:"groupId,omitempty" xml:"groupId,omitempty" require:"true"`
  // {"zh_CN":"开关。
  // ON：开
  // OFF：关","en":"Switch. 
  // ON: On
  // OFF: Off"}
  RuleSwitch *string `json:"ruleSwitch,omitempty" xml:"ruleSwitch,omitempty" require:"true"`
  // {"zh_CN":"请求uri，当fromType=SPECIFY时，groupRuleFromList必填。
  // ALL：所有请求
  // SPECIFY：指定uri","en":"Request uri,when fromType=SPECIFY, groupRuleFromList required. 
  // ALL: All requests
  // SPECIFY: Specified uri"}
  FromType *string `json:"fromType,omitempty" xml:"fromType,omitempty" require:"true"`
  // {"zh_CN":"分组规则请求uri，当matchType=EQUAL/START_WITH时，content必须以/开头。","en":"Grouping rule request uri,when matchType=EQUAL/STARTWITH, content must start with/."}
  GroupRuleFromList []*SaveNavigationRuleGroupRuleFromSaveDTO `json:"groupRuleFromList,omitempty" xml:"groupRuleFromList,omitempty" type:"Repeated"`
  // {"zh_CN":"目标uri类型，toType=SPECIFY时，toContent必填且必须以/开头，当toType=INDEX/SAME时，content为空。
  // INDEX：首页
  // SPECIFY：指定uri
  // SAME：继承请求路","en":"Target uri type,when from Type=SPECIY, toContent required and it must begin with /,when toType=INDEX/SAME, toContent is empty.
  // INDEX: Home page
  // SPECIFY: Specified uri
  // SAME: Inherited request uri"}
  ToType *string `json:"toType,omitempty" xml:"toType,omitempty" require:"true"`
  // {"zh_CN":"目标uri，仅支持字符 [a-zA-Z0-9.-_/=?:&]。","en":"Target uri,only supports characters [a-zA-Z0-9. - _/=?:&]."}
  ToContent *string `json:"toContent,omitempty" xml:"toContent,omitempty"`
  // {"zh_CN":"是否携带查询串。
  // TRUE：是
  // FALSE：否","en":"Whether to carry a query string. 
  // TRUE: Yes 
  // FALSE: No"}
  WithQs *string `json:"withQs,omitempty" xml:"withQs,omitempty" require:"true"`
}

func (s SaveNavigationRuleRequest) String() string {
  return tea.Prettify(s)
}

func (s SaveNavigationRuleRequest) GoString() string {
  return s.String()
}

func (s *SaveNavigationRuleRequest) SetGroupId(v string) *SaveNavigationRuleRequest {
  s.GroupId = &v
  return s
}

func (s *SaveNavigationRuleRequest) SetRuleSwitch(v string) *SaveNavigationRuleRequest {
  s.RuleSwitch = &v
  return s
}

func (s *SaveNavigationRuleRequest) SetFromType(v string) *SaveNavigationRuleRequest {
  s.FromType = &v
  return s
}

func (s *SaveNavigationRuleRequest) SetGroupRuleFromList(v []*SaveNavigationRuleGroupRuleFromSaveDTO) *SaveNavigationRuleRequest {
  s.GroupRuleFromList = v
  return s
}

func (s *SaveNavigationRuleRequest) SetToType(v string) *SaveNavigationRuleRequest {
  s.ToType = &v
  return s
}

func (s *SaveNavigationRuleRequest) SetToContent(v string) *SaveNavigationRuleRequest {
  s.ToContent = &v
  return s
}

func (s *SaveNavigationRuleRequest) SetWithQs(v string) *SaveNavigationRuleRequest {
  s.WithQs = &v
  return s
}

type SaveNavigationRuleResponse struct {
  // {'en':'Please refer to the error code for exceptions.', 'zh_CN':'请参照错误码。','dictionary':'belong=WAAP-MS-Ext|dict=waap_retCodeEnum'}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {'en':'Description.', 'zh_CN':'描述信息。'}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
}

func (s SaveNavigationRuleResponse) String() string {
  return tea.Prettify(s)
}

func (s SaveNavigationRuleResponse) GoString() string {
  return s.String()
}

func (s *SaveNavigationRuleResponse) SetCode(v string) *SaveNavigationRuleResponse {
  s.Code = &v
  return s
}

func (s *SaveNavigationRuleResponse) SetMsg(v string) *SaveNavigationRuleResponse {
  s.Msg = &v
  return s
}

type SaveNavigationRuleGroupRuleFromSaveDTO struct {
  // {"zh_CN":"请求uri匹配类型。
  // EQUAL：等于
  // CONTAIN：包含
  // START_WITH：开头是
  // END_WITH：结尾是","en":"Request uri match type.
  // EQUAL: equal to
  // CONTAIN: contains
  // START_WITH: starts with,
  // END_WITH: ends with"}
  MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty"`
  // {"zh_CN":"请求uri，仅支持字符 [a-zA-Z0-9.-_/=?:&]。","en":"Request uri,only supports characters [a-zA-Z0-9. - _/=?:&]."}
  Content *string `json:"content,omitempty" xml:"content,omitempty"`
}

func (s SaveNavigationRuleGroupRuleFromSaveDTO) String() string {
  return tea.Prettify(s)
}

func (s SaveNavigationRuleGroupRuleFromSaveDTO) GoString() string {
  return s.String()
}

func (s *SaveNavigationRuleGroupRuleFromSaveDTO) SetMatchType(v string) *SaveNavigationRuleGroupRuleFromSaveDTO {
  s.MatchType = &v
  return s
}

func (s *SaveNavigationRuleGroupRuleFromSaveDTO) SetContent(v string) *SaveNavigationRuleGroupRuleFromSaveDTO {
  s.Content = &v
  return s
}

type SaveNavigationRulePaths struct {
}

func (s SaveNavigationRulePaths) String() string {
  return tea.Prettify(s)
}

func (s SaveNavigationRulePaths) GoString() string {
  return s.String()
}

type SaveNavigationRuleParameters struct {
}

func (s SaveNavigationRuleParameters) String() string {
  return tea.Prettify(s)
}

func (s SaveNavigationRuleParameters) GoString() string {
  return s.String()
}

type SaveNavigationRuleRequestHeader struct {
  // {"zh_CN":"安全服务类型。有使用多个不同的安全服务时，需要填写具体的服务类型。","en":"Security service type. Please enter a specific service type, if you purchase multiple security services.","dictionary":"belong=WAAP-MS-Ext|dict=waap_serviceType"}
  ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty"`
}

func (s SaveNavigationRuleRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s SaveNavigationRuleRequestHeader) GoString() string {
  return s.String()
}

func (s *SaveNavigationRuleRequestHeader) SetServiceType(v string) *SaveNavigationRuleRequestHeader {
  s.ServiceType = &v
  return s
}

type SaveNavigationRuleResponseHeader struct {
}

func (s SaveNavigationRuleResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s SaveNavigationRuleResponseHeader) GoString() string {
  return s.String()
}




type CreateWebsiteGroupRequest struct {
  // {'en':'Name.','zh_CN':'名称。'}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {'en':'Remark.','zh_CN':'备注。'}
  Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
  // {"zh_CN":"端口配置。","en":"Port configuration."}
  PortRelList []*CreateWebsiteGroupGroupPortRelAddDTO `json:"portRelList,omitempty" xml:"portRelList,omitempty" require:"true" type:"Repeated"`
  // {"zh_CN":"关联域名列表。","en":"List of Associate hostname."}
  DomainList []*string `json:"domainList,omitempty" xml:"domainList,omitempty" require:"true" type:"Repeated"`
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

func (s *CreateWebsiteGroupRequest) SetPortRelList(v []*CreateWebsiteGroupGroupPortRelAddDTO) *CreateWebsiteGroupRequest {
  s.PortRelList = v
  return s
}

func (s *CreateWebsiteGroupRequest) SetDomainList(v []*string) *CreateWebsiteGroupRequest {
  s.DomainList = v
  return s
}

type CreateWebsiteGroupResponse struct {
  // {'en':'Please refer to the error code for exceptions.', 'zh_CN':'请参照错误码。','dictionary':'belong=WAAP-MS-Ext|dict=waap_retCodeEnum'}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {'en':'Description.', 'zh_CN':'描述信息。'}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {'en':'Data.', 'zh_CN':'出参数据。'}
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

type CreateWebsiteGroupGroupPortRelAddDTO struct {
  // {"zh_CN":"端口列表。
  //     2000~2099：除2012,2080
  // 5000~5099
  // 8200~8999：除8245
  // 9000~9999：除9119，9331，9518，9851，9852，9996
  // 15000~15999
  // 17000~17999
  // 18000~18099
  // 50000~50199","en":"Port list.
  // 2000~2099:Except for 2012,2080
  // 5000~5099
  // 8200~8999:Except for 8245
  // 9000~9999:Except for 9119,9331,9518,9851,9852,9996
  // 15000~15999
  // 17000~17999
  // 18000~18099
  // 50000~50199"}
  PortList []*int32 `json:"portList,omitempty" xml:"portList,omitempty" require:"true" type:"Repeated"`
  // {"zh_CN":"协议。
  // HTTP：HTTP协议
  // HTTPS：HTTPS协议","en":"Protocol. 
  // HTTP: HTTP protocol 
  // HTTPS: HTTPS protocol"}
  Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty" require:"true"`
}

func (s CreateWebsiteGroupGroupPortRelAddDTO) String() string {
  return tea.Prettify(s)
}

func (s CreateWebsiteGroupGroupPortRelAddDTO) GoString() string {
  return s.String()
}

func (s *CreateWebsiteGroupGroupPortRelAddDTO) SetPortList(v []*int32) *CreateWebsiteGroupGroupPortRelAddDTO {
  s.PortList = v
  return s
}

func (s *CreateWebsiteGroupGroupPortRelAddDTO) SetProtocol(v string) *CreateWebsiteGroupGroupPortRelAddDTO {
  s.Protocol = &v
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

type CreateWebsiteGroupRequestHeader struct {
  // {"zh_CN":"安全服务类型。有使用多个不同的安全服务时，需要填写具体的服务类型。","en":"Security service type. Please enter a specific service type, if you purchase multiple security services.","dictionary":"belong=WAAP-MS-Ext|dict=waap_serviceType"}
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

type CreateWebsiteGroupResponseHeader struct {
}

func (s CreateWebsiteGroupResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s CreateWebsiteGroupResponseHeader) GoString() string {
  return s.String()
}




type ListGroupDetailsRequest struct {
  // {'en':'ID List.', 'zh_CN':'ID列表。'}
  IdList []*string `json:"idList,omitempty" xml:"idList,omitempty" type:"Repeated"`
  // {'en':'Name list,supports fuzzy queries.','zh_CN':'名称列表，支持模糊查询。'}
  NameList []*string `json:"nameList,omitempty" xml:"nameList,omitempty" type:"Repeated"`
  // {'en':'Hostname list.','zh_CN':'域名列表。'}
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

type ListGroupDetailsResponse struct {
  // {'en':'Please refer to the error code for exceptions.', 'zh_CN':'请参照错误码。','dictionary':'belong=WAAP-MS-Ext|dict=waap_retCodeEnum'}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {'en':'Description.', 'zh_CN':'描述信息。'}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {'en':'Data.', 'zh_CN':'出参数据。'}
  Data []*ListGroupDetailsGroupInfoDetailVO `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
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

func (s *ListGroupDetailsResponse) SetData(v []*ListGroupDetailsGroupInfoDetailVO) *ListGroupDetailsResponse {
  s.Data = v
  return s
}

type ListGroupDetailsGroupInfoDetailVO struct {
  // {"zh_CN":"ID。","en":"ID."}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {"zh_CN":"名称。","en":"Name."}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"zh_CN":"备注。","en":"Remark."}
  Remark *string `json:"remark,omitempty" xml:"remark,omitempty" require:"true"`
  // {"zh_CN":"端口配置。","en":"Port configuration."}
  PortRelList []*ListGroupDetailsGroupPortRelVO `json:"portRelList,omitempty" xml:"portRelList,omitempty" require:"true" type:"Repeated"`
  // {"zh_CN":"关联域名列表。","en":"List of Associate hostname."}
  DomainList []*string `json:"domainList,omitempty" xml:"domainList,omitempty" require:"true" type:"Repeated"`
  // {"zh_CN":"是否关联有入口域名。","en":"Whether there is an entry hostname associated with it."}
  RelEntryDomain *bool `json:"relEntryDomain,omitempty" xml:"relEntryDomain,omitempty" require:"true"`
}

func (s ListGroupDetailsGroupInfoDetailVO) String() string {
  return tea.Prettify(s)
}

func (s ListGroupDetailsGroupInfoDetailVO) GoString() string {
  return s.String()
}

func (s *ListGroupDetailsGroupInfoDetailVO) SetId(v string) *ListGroupDetailsGroupInfoDetailVO {
  s.Id = &v
  return s
}

func (s *ListGroupDetailsGroupInfoDetailVO) SetName(v string) *ListGroupDetailsGroupInfoDetailVO {
  s.Name = &v
  return s
}

func (s *ListGroupDetailsGroupInfoDetailVO) SetRemark(v string) *ListGroupDetailsGroupInfoDetailVO {
  s.Remark = &v
  return s
}

func (s *ListGroupDetailsGroupInfoDetailVO) SetPortRelList(v []*ListGroupDetailsGroupPortRelVO) *ListGroupDetailsGroupInfoDetailVO {
  s.PortRelList = v
  return s
}

func (s *ListGroupDetailsGroupInfoDetailVO) SetDomainList(v []*string) *ListGroupDetailsGroupInfoDetailVO {
  s.DomainList = v
  return s
}

func (s *ListGroupDetailsGroupInfoDetailVO) SetRelEntryDomain(v bool) *ListGroupDetailsGroupInfoDetailVO {
  s.RelEntryDomain = &v
  return s
}

type ListGroupDetailsGroupPortRelVO struct {
  // {"zh_CN":"端口列表。","en":"Port list."}
  PortList []*int32 `json:"portList,omitempty" xml:"portList,omitempty" require:"true" type:"Repeated"`
  // {"zh_CN":"协议。
  // HTTP：HTTP协议
  // HTTPS：HTTPS协议","en":"Protocol. 
  // HTTP: HTTP protocol 
  // HTTPS: HTTPS protocol"}
  Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty" require:"true"`
}

func (s ListGroupDetailsGroupPortRelVO) String() string {
  return tea.Prettify(s)
}

func (s ListGroupDetailsGroupPortRelVO) GoString() string {
  return s.String()
}

func (s *ListGroupDetailsGroupPortRelVO) SetPortList(v []*int32) *ListGroupDetailsGroupPortRelVO {
  s.PortList = v
  return s
}

func (s *ListGroupDetailsGroupPortRelVO) SetProtocol(v string) *ListGroupDetailsGroupPortRelVO {
  s.Protocol = &v
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

type ListGroupDetailsRequestHeader struct {
  // {"zh_CN":"安全服务类型。有使用多个不同的安全服务时，需要填写具体的服务类型。","en":"Security service type. Please enter a specific service type, if you purchase multiple security services.","dictionary":"belong=WAAP-MS-Ext|dict=waap_serviceType"}
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

type ListGroupDetailsResponseHeader struct {
}

func (s ListGroupDetailsResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ListGroupDetailsResponseHeader) GoString() string {
  return s.String()
}




type ListNavigationRulesRequest struct {
  // {'en':'Website group ID list.','zh_CN':'网站分组ID列表。'}
  GroupIdList []*string `json:"groupIdList,omitempty" xml:"groupIdList,omitempty" type:"Repeated"`
  // {'en':'Website group name list,supports fuzzy queries.','zh_CN':'网站分组名称列表，支持模糊查询。'}
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

type ListNavigationRulesResponse struct {
  // {'en':'Please refer to the error code for exceptions.', 'zh_CN':'请参照错误码。','dictionary':'belong=WAAP-MS-Ext|dict=waap_retCodeEnum'}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {'en':'Description.', 'zh_CN':'描述信息。'}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {'en':'Data.', 'zh_CN':'出参数据。'}
  Data []*ListNavigationRulesGroupRuleVO `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
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

func (s *ListNavigationRulesResponse) SetData(v []*ListNavigationRulesGroupRuleVO) *ListNavigationRulesResponse {
  s.Data = v
  return s
}

type ListNavigationRulesGroupRuleFromVO struct {
  // {"zh_CN":"请求uri匹配类型。
  // EQUAL：等于
  // CONTAIN：包含
  // START_WITH：开头是
  // END_WITH：结尾是","en":"Request uri match type. 
  // EQUAL: equal to
  // CONTAIN: contains
  // START_WITH: starts with
  // END_WITH: ends with"}
  MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
  // {"zh_CN":"请求uri。","en":"Request uri."}
  Content *string `json:"content,omitempty" xml:"content,omitempty" require:"true"`
}

func (s ListNavigationRulesGroupRuleFromVO) String() string {
  return tea.Prettify(s)
}

func (s ListNavigationRulesGroupRuleFromVO) GoString() string {
  return s.String()
}

func (s *ListNavigationRulesGroupRuleFromVO) SetMatchType(v string) *ListNavigationRulesGroupRuleFromVO {
  s.MatchType = &v
  return s
}

func (s *ListNavigationRulesGroupRuleFromVO) SetContent(v string) *ListNavigationRulesGroupRuleFromVO {
  s.Content = &v
  return s
}

type ListNavigationRulesGroupRuleVO struct {
  // {"zh_CN":"ID。","en":"ID."}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {"zh_CN":"网站分组ID。","en":"The website group ID."}
  GroupId *string `json:"groupId,omitempty" xml:"groupId,omitempty" require:"true"`
  // {"zh_CN":"分组名称。","en":"Group name."}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"zh_CN":"开关。
  // ON：开
  // OFF：关","en":"Switch. 
  // ON: On
  // OFF: Off"}
  RuleSwitch *string `json:"ruleSwitch,omitempty" xml:"ruleSwitch,omitempty" require:"true"`
  // {"zh_CN":"请求uri类型。
  // ALL：所有请求
  // SPECIFY：指定uri","en":"Request uri type. 
  // ALL: All requests 
  // SPECIFY: Specified uri"}
  FromType *string `json:"fromType,omitempty" xml:"fromType,omitempty" require:"true"`
  // {"zh_CN":"分组规则请求uri。","en":"Grouping rule request uri."}
  GroupRuleFromList []*ListNavigationRulesGroupRuleFromVO `json:"groupRuleFromList,omitempty" xml:"groupRuleFromList,omitempty" require:"true" type:"Repeated"`
  // {"zh_CN":"目标uri类型。
  // INDEX：首页
  // SPECIFY：指定uri
  // SAME：继承请求uri","en":"Target uri type. 
  // INDEX: Home page
  // SPECIFY: Specified uri
  // SAME: Inherited request uri"}
  ToType *string `json:"toType,omitempty" xml:"toType,omitempty" require:"true"`
  // {"zh_CN":"目标uri。","en":"Target uri."}
  ToContent *string `json:"toContent,omitempty" xml:"toContent,omitempty" require:"true"`
  // {"zh_CN":"是否携带查询串。
  // TRUE：是
  // FALSE：否","en":"Whether to carry a query string. 
  // TRUE: Yes 
  // FALSE: No"}
  WithQs *string `json:"withQs,omitempty" xml:"withQs,omitempty" require:"true"`
}

func (s ListNavigationRulesGroupRuleVO) String() string {
  return tea.Prettify(s)
}

func (s ListNavigationRulesGroupRuleVO) GoString() string {
  return s.String()
}

func (s *ListNavigationRulesGroupRuleVO) SetId(v string) *ListNavigationRulesGroupRuleVO {
  s.Id = &v
  return s
}

func (s *ListNavigationRulesGroupRuleVO) SetGroupId(v string) *ListNavigationRulesGroupRuleVO {
  s.GroupId = &v
  return s
}

func (s *ListNavigationRulesGroupRuleVO) SetName(v string) *ListNavigationRulesGroupRuleVO {
  s.Name = &v
  return s
}

func (s *ListNavigationRulesGroupRuleVO) SetRuleSwitch(v string) *ListNavigationRulesGroupRuleVO {
  s.RuleSwitch = &v
  return s
}

func (s *ListNavigationRulesGroupRuleVO) SetFromType(v string) *ListNavigationRulesGroupRuleVO {
  s.FromType = &v
  return s
}

func (s *ListNavigationRulesGroupRuleVO) SetGroupRuleFromList(v []*ListNavigationRulesGroupRuleFromVO) *ListNavigationRulesGroupRuleVO {
  s.GroupRuleFromList = v
  return s
}

func (s *ListNavigationRulesGroupRuleVO) SetToType(v string) *ListNavigationRulesGroupRuleVO {
  s.ToType = &v
  return s
}

func (s *ListNavigationRulesGroupRuleVO) SetToContent(v string) *ListNavigationRulesGroupRuleVO {
  s.ToContent = &v
  return s
}

func (s *ListNavigationRulesGroupRuleVO) SetWithQs(v string) *ListNavigationRulesGroupRuleVO {
  s.WithQs = &v
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

type ListNavigationRulesRequestHeader struct {
  // {"zh_CN":"安全服务类型。有使用多个不同的安全服务时，需要填写具体的服务类型。","en":"Security service type. Please enter a specific service type, if you purchase multiple security services.","dictionary":"belong=WAAP-MS-Ext|dict=waap_serviceType"}
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

type ListNavigationRulesResponseHeader struct {
}

func (s ListNavigationRulesResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ListNavigationRulesResponseHeader) GoString() string {
  return s.String()
}




type UpdateEntryDomainRequest struct {
  // {'en':'ID.', 'zh_CN':'ID。'}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {"zh_CN":"网站分组ID","en":"Website group ID"}
  GroupId *string `json:"groupId,omitempty" xml:"groupId,omitempty" require:"true"`
  // {"zh_CN":"端口配置。","en":"Port configuration."}
  PortRelList []*UpdateEntryDomainEntryDomainPortRelEditDTO `json:"portRelList,omitempty" xml:"portRelList,omitempty" require:"true" type:"Repeated"`
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

func (s *UpdateEntryDomainRequest) SetPortRelList(v []*UpdateEntryDomainEntryDomainPortRelEditDTO) *UpdateEntryDomainRequest {
  s.PortRelList = v
  return s
}

type UpdateEntryDomainResponse struct {
  // {'en':'Please refer to the error code for exceptions.', 'zh_CN':'请参照错误码。','dictionary':'belong=WAAP-MS-Ext|dict=waap_retCodeEnum'}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {'en':'Description.', 'zh_CN':'描述信息。'}
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

type UpdateEntryDomainEntryDomainPortRelEditDTO struct {
  // {"zh_CN":"端口列表。
  //     2000~2099：除2012,2080
  // 5000~5099
  // 8200~8999：除8245
  // 9000~9999：除9119，9331，9518，9851，9852，9996
  // 15000~15999
  // 17000~17999
  // 18000~18099
  // 50000~50199","en":"Port list.
  // 2000~2099:Except for 2012,2080
  // 5000~5099
  // 8200~8999:Except for 8245
  // 9000~9999:Except for 9119,9331,9518,9851,9852,9996
  // 15000~15999
  // 17000~17999
  // 18000~18099
  // 50000~50199"}
  PortList []*int32 `json:"portList,omitempty" xml:"portList,omitempty" require:"true" type:"Repeated"`
  // {"zh_CN":"协议。
  // HTTP：HTTP协议
  // HTTPS：HTTPS协议","en":"Protocol. 
  // HTTP: HTTP protocol 
  // HTTPS: HTTPS protocol"}
  Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty" require:"true"`
}

func (s UpdateEntryDomainEntryDomainPortRelEditDTO) String() string {
  return tea.Prettify(s)
}

func (s UpdateEntryDomainEntryDomainPortRelEditDTO) GoString() string {
  return s.String()
}

func (s *UpdateEntryDomainEntryDomainPortRelEditDTO) SetPortList(v []*int32) *UpdateEntryDomainEntryDomainPortRelEditDTO {
  s.PortList = v
  return s
}

func (s *UpdateEntryDomainEntryDomainPortRelEditDTO) SetProtocol(v string) *UpdateEntryDomainEntryDomainPortRelEditDTO {
  s.Protocol = &v
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

type UpdateEntryDomainRequestHeader struct {
  // {"zh_CN":"安全服务类型。有使用多个不同的安全服务时，需要填写具体的服务类型。","en":"Security service type. Please enter a specific service type, if you purchase multiple security services.","dictionary":"belong=WAAP-MS-Ext|dict=waap_serviceType"}
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

type UpdateEntryDomainResponseHeader struct {
}

func (s UpdateEntryDomainResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s UpdateEntryDomainResponseHeader) GoString() string {
  return s.String()
}




type UpdateWebsiteGroupRequest struct {
  // {'en':'ID.', 'zh_CN':'ID。'}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {'en':'Remark.','zh_CN':'备注。'}
  Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
  // {"zh_CN":"端口配置。","en":"Port configuration."}
  PortRelList []*UpdateWebsiteGroupGroupPortRelEditDTO `json:"portRelList,omitempty" xml:"portRelList,omitempty" require:"true" type:"Repeated"`
  // {"zh_CN":"关联域名列表。","en":"List of Associate hostname."}
  DomainList []*string `json:"domainList,omitempty" xml:"domainList,omitempty" require:"true" type:"Repeated"`
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

func (s *UpdateWebsiteGroupRequest) SetPortRelList(v []*UpdateWebsiteGroupGroupPortRelEditDTO) *UpdateWebsiteGroupRequest {
  s.PortRelList = v
  return s
}

func (s *UpdateWebsiteGroupRequest) SetDomainList(v []*string) *UpdateWebsiteGroupRequest {
  s.DomainList = v
  return s
}

type UpdateWebsiteGroupResponse struct {
  // {'en':'Please refer to the error code for exceptions.', 'zh_CN':'请参照错误码。','dictionary':'belong=WAAP-MS-Ext|dict=waap_retCodeEnum'}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {'en':'Description.', 'zh_CN':'描述信息。'}
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

type UpdateWebsiteGroupGroupPortRelEditDTO struct {
  // {"zh_CN":"端口列表。
  //     2000~2099：除2012,2080
  // 5000~5099
  // 8200~8999：除8245
  // 9000~9999：除9119，9331，9518，9851，9852，9996
  // 15000~15999
  // 17000~17999
  // 18000~18099
  // 50000~50199","en":"Port list.
  // 2000~2099:Except for 2012,2080
  // 5000~5099
  // 8200~8999:Except for 8245
  // 9000~9999:Except for 9119,9331,9518,9851,9852,9996
  // 15000~15999
  // 17000~17999
  // 18000~18099
  // 50000~50199"}
  PortList []*int32 `json:"portList,omitempty" xml:"portList,omitempty" require:"true" type:"Repeated"`
  // {"zh_CN":"协议。
  // HTTP：HTTP协议
  // HTTPS：HTTPS协议","en":"Protocol. 
  // HTTP: HTTP protocol 
  // HTTPS: HTTPS protocol"}
  Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty" require:"true"`
}

func (s UpdateWebsiteGroupGroupPortRelEditDTO) String() string {
  return tea.Prettify(s)
}

func (s UpdateWebsiteGroupGroupPortRelEditDTO) GoString() string {
  return s.String()
}

func (s *UpdateWebsiteGroupGroupPortRelEditDTO) SetPortList(v []*int32) *UpdateWebsiteGroupGroupPortRelEditDTO {
  s.PortList = v
  return s
}

func (s *UpdateWebsiteGroupGroupPortRelEditDTO) SetProtocol(v string) *UpdateWebsiteGroupGroupPortRelEditDTO {
  s.Protocol = &v
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

type UpdateWebsiteGroupRequestHeader struct {
  // {"zh_CN":"安全服务类型。有使用多个不同的安全服务时，需要填写具体的服务类型。","en":"Security service type. Please enter a specific service type, if you purchase multiple security services.","dictionary":"belong=WAAP-MS-Ext|dict=waap_serviceType"}
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

type UpdateWebsiteGroupResponseHeader struct {
}

func (s UpdateWebsiteGroupResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s UpdateWebsiteGroupResponseHeader) GoString() string {
  return s.String()
}




type DeleteEntryDomainRequest struct {
  // {'en':'ID list.', 'zh_CN':'ID列表。'}
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

type DeleteEntryDomainResponse struct {
  // {'en':'Please refer to the error code for exceptions.', 'zh_CN':'请参照错误码。','dictionary':'belong=WAAP-MS-Ext|dict=waap_retCodeEnum'}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {'en':'Description.', 'zh_CN':'描述信息。'}
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

type DeleteEntryDomainRequestHeader struct {
  // {"zh_CN":"安全服务类型。有使用多个不同的安全服务时，需要填写具体的服务类型。","en":"Security service type. Please enter a specific service type, if you purchase multiple security services.","dictionary":"belong=WAAP-MS-Ext|dict=waap_serviceType"}
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

type DeleteEntryDomainResponseHeader struct {
}

func (s DeleteEntryDomainResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteEntryDomainResponseHeader) GoString() string {
  return s.String()
}




type AddEntryDomainRequest struct {
  // {"zh_CN":"域名。","en":"Hostname."}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {"zh_CN":"网站分组ID","en":"Website group ID"}
  GroupId *string `json:"groupId,omitempty" xml:"groupId,omitempty" require:"true"`
  // {"zh_CN":"端口配置。","en":"Port configuration."}
  PortRelList []*AddEntryDomainEntryDomainPortRelAddDTO `json:"portRelList,omitempty" xml:"portRelList,omitempty" require:"true" type:"Repeated"`
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

func (s *AddEntryDomainRequest) SetPortRelList(v []*AddEntryDomainEntryDomainPortRelAddDTO) *AddEntryDomainRequest {
  s.PortRelList = v
  return s
}

type AddEntryDomainResponse struct {
  // {'en':'Please refer to the error code for exceptions.', 'zh_CN':'请参照错误码。','dictionary':'belong=WAAP-MS-Ext|dict=waap_retCodeEnum'}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {'en':'Description.', 'zh_CN':'描述信息。'}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {'en':'Data.', 'zh_CN':'出参数据。'}
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

type AddEntryDomainEntryDomainPortRelAddDTO struct {
  // {"zh_CN":"端口列表。
  //     2000~2099：除2012,2080
  // 5000~5099
  // 8200~8999：除8245
  // 9000~9999：除9119，9331，9518，9851，9852，9996
  // 15000~15999
  // 17000~17999
  // 18000~18099
  // 50000~50199","en":"Port list.
  // 2000~2099:Except for 2012,2080
  // 5000~5099
  // 8200~8999:Except for 8245
  // 9000~9999:Except for 9119,9331,9518,9851,9852,9996
  // 15000~15999
  // 17000~17999
  // 18000~18099
  // 50000~50199"}
  PortList []*int32 `json:"portList,omitempty" xml:"portList,omitempty" require:"true" type:"Repeated"`
  // {"zh_CN":"协议。
  // HTTP：HTTP协议
  // HTTPS：HTTPS协议","en":"Protocol. 
  // HTTP: HTTP protocol 
  // HTTPS: HTTPS protocol"}
  Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty" require:"true"`
}

func (s AddEntryDomainEntryDomainPortRelAddDTO) String() string {
  return tea.Prettify(s)
}

func (s AddEntryDomainEntryDomainPortRelAddDTO) GoString() string {
  return s.String()
}

func (s *AddEntryDomainEntryDomainPortRelAddDTO) SetPortList(v []*int32) *AddEntryDomainEntryDomainPortRelAddDTO {
  s.PortList = v
  return s
}

func (s *AddEntryDomainEntryDomainPortRelAddDTO) SetProtocol(v string) *AddEntryDomainEntryDomainPortRelAddDTO {
  s.Protocol = &v
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

type AddEntryDomainRequestHeader struct {
  // {"zh_CN":"安全服务类型。有使用多个不同的安全服务时，需要填写具体的服务类型。","en":"Security service type. Please enter a specific service type, if you purchase multiple security services.","dictionary":"belong=WAAP-MS-Ext|dict=waap_serviceType"}
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

type AddEntryDomainResponseHeader struct {
}

func (s AddEntryDomainResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s AddEntryDomainResponseHeader) GoString() string {
  return s.String()
}




type DeleteWebsiteGroupRequest struct {
  // {'en':'ID list.', 'zh_CN':'ID列表。'}
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

type DeleteWebsiteGroupResponse struct {
  // {'en':'Please refer to the error code for exceptions.', 'zh_CN':'请参照错误码。','dictionary':'belong=WAAP-MS-Ext|dict=waap_retCodeEnum'}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {'en':'Description.', 'zh_CN':'描述信息。'}
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

type DeleteWebsiteGroupRequestHeader struct {
  // {"zh_CN":"安全服务类型。有使用多个不同的安全服务时，需要填写具体的服务类型。","en":"Security service type. Please enter a specific service type, if you purchase multiple security services.","dictionary":"belong=WAAP-MS-Ext|dict=waap_serviceType"}
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

type DeleteWebsiteGroupResponseHeader struct {
}

func (s DeleteWebsiteGroupResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteWebsiteGroupResponseHeader) GoString() string {
  return s.String()
}




type RestartEntryDomainRequest struct {
  // {'en':'Hostname list.', 'zh_CN':'域名列表。'}
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

type RestartEntryDomainResponse struct {
  // {'en':'Please refer to the error code for exceptions.', 'zh_CN':'请参照错误码。','dictionary':'belong=WAAP-MS-Ext|dict=waap_retCodeEnum'}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {'en':'Description.', 'zh_CN':'描述信息。'}
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

type RestartEntryDomainRequestHeader struct {
  // {"zh_CN":"安全服务类型。有使用多个不同的安全服务时，需要填写具体的服务类型。","en":"Security service type. Please enter a specific service type, if you purchase multiple security services.","dictionary":"belong=WAAP-MS-Ext|dict=waap_serviceType"}
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

type RestartEntryDomainResponseHeader struct {
}

func (s RestartEntryDomainResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s RestartEntryDomainResponseHeader) GoString() string {
  return s.String()
}




type ListEntryDomainsRequest struct {
  // {"zh_CN":"ID列表","en":"ID List"}
  IdList []*string `json:"idList,omitempty" xml:"idList,omitempty" type:"Repeated"`
  // {"zh_CN":"入口域名列表。","en":"List of entry hostname."}
  DomainList []*string `json:"domainList,omitempty" xml:"domainList,omitempty" type:"Repeated"`
  // {"zh_CN":"网站分组ID列表","en":"Website group ID list"}
  GroupIdList []*string `json:"groupIdList,omitempty" xml:"groupIdList,omitempty" type:"Repeated"`
  // {'en':'Website group name list,supports fuzzy queries.','zh_CN':'网站分组名称列表，支持模糊查询。'}
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

type ListEntryDomainsResponse struct {
  // {'en':'Please refer to the error code for exceptions.', 'zh_CN':'请参照错误码。','dictionary':'belong=WAAP-MS-Ext|dict=waap_retCodeEnum'}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {'en':'Description.', 'zh_CN':'描述信息。'}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {'en':'Data.', 'zh_CN':'出参数据。'}
  Data []*ListEntryDomainsEntryDomainDetailVO `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
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

func (s *ListEntryDomainsResponse) SetData(v []*ListEntryDomainsEntryDomainDetailVO) *ListEntryDomainsResponse {
  s.Data = v
  return s
}

type ListEntryDomainsEntryDomainDetailVO struct {
  // {"zh_CN":"ID。","en":"ID."}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {"zh_CN":"域名。","en":"Hostname."}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {"zh_CN":"关联网站分组信息。","en":"Related website grouping information."}
  GroupInfo *ListEntryDomainsGroupInfoVO `json:"groupInfo,omitempty" xml:"groupInfo,omitempty" require:"true"`
  // {"zh_CN":"端口配置。","en":"Port configuration."}
  PortRelList []*ListEntryDomainsEntryDomainPortRelVO `json:"portRelList,omitempty" xml:"portRelList,omitempty" require:"true" type:"Repeated"`
  // {"zh_CN":"状态。
  // ENABLE：启用
  // DISABLE：禁用","en":"Status. 
  // ENABLE: Enable 
  // DISABLE: Disable"}
  Status *string `json:"status,omitempty" xml:"status,omitempty" require:"true"`
}

func (s ListEntryDomainsEntryDomainDetailVO) String() string {
  return tea.Prettify(s)
}

func (s ListEntryDomainsEntryDomainDetailVO) GoString() string {
  return s.String()
}

func (s *ListEntryDomainsEntryDomainDetailVO) SetId(v string) *ListEntryDomainsEntryDomainDetailVO {
  s.Id = &v
  return s
}

func (s *ListEntryDomainsEntryDomainDetailVO) SetDomain(v string) *ListEntryDomainsEntryDomainDetailVO {
  s.Domain = &v
  return s
}

func (s *ListEntryDomainsEntryDomainDetailVO) SetGroupInfo(v *ListEntryDomainsGroupInfoVO) *ListEntryDomainsEntryDomainDetailVO {
  s.GroupInfo = v
  return s
}

func (s *ListEntryDomainsEntryDomainDetailVO) SetPortRelList(v []*ListEntryDomainsEntryDomainPortRelVO) *ListEntryDomainsEntryDomainDetailVO {
  s.PortRelList = v
  return s
}

func (s *ListEntryDomainsEntryDomainDetailVO) SetStatus(v string) *ListEntryDomainsEntryDomainDetailVO {
  s.Status = &v
  return s
}

type ListEntryDomainsGroupInfoVO struct {
  // {'en':'Website group ID.', 'zh_CN':'网站分组ID。'}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {'en':'Name.','zh_CN':'名称。'}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
}

func (s ListEntryDomainsGroupInfoVO) String() string {
  return tea.Prettify(s)
}

func (s ListEntryDomainsGroupInfoVO) GoString() string {
  return s.String()
}

func (s *ListEntryDomainsGroupInfoVO) SetId(v string) *ListEntryDomainsGroupInfoVO {
  s.Id = &v
  return s
}

func (s *ListEntryDomainsGroupInfoVO) SetName(v string) *ListEntryDomainsGroupInfoVO {
  s.Name = &v
  return s
}

type ListEntryDomainsEntryDomainPortRelVO struct {
  // {"zh_CN":"端口列表。
  //     2000~2099：除2012,2080
  // 5000~5099
  // 8200~8999：除8245
  // 9000~9999：除9119，9331，9518，9851，9852，9996
  // 15000~15999
  // 17000~17999
  // 18000~18099
  // 50000~50199","en":"Port list.
  // 2000~2099:Except for 2012,2080
  // 5000~5099
  // 8200~8999:Except for 8245
  // 9000~9999:Except for 9119,9331,9518,9851,9852,9996
  // 15000~15999
  // 17000~17999
  // 18000~18099
  // 50000~50199"}
  PortList []*int32 `json:"portList,omitempty" xml:"portList,omitempty" require:"true" type:"Repeated"`
  // {"zh_CN":"协议。
  // HTTP：HTTP协议
  // HTTPS：HTTPS协议","en":"Protocol. 
  // HTTP: HTTP protocol 
  // HTTPS: HTTPS protocol"}
  Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty" require:"true"`
}

func (s ListEntryDomainsEntryDomainPortRelVO) String() string {
  return tea.Prettify(s)
}

func (s ListEntryDomainsEntryDomainPortRelVO) GoString() string {
  return s.String()
}

func (s *ListEntryDomainsEntryDomainPortRelVO) SetPortList(v []*int32) *ListEntryDomainsEntryDomainPortRelVO {
  s.PortList = v
  return s
}

func (s *ListEntryDomainsEntryDomainPortRelVO) SetProtocol(v string) *ListEntryDomainsEntryDomainPortRelVO {
  s.Protocol = &v
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

type ListEntryDomainsRequestHeader struct {
  // {"zh_CN":"安全服务类型。有使用多个不同的安全服务时，需要填写具体的服务类型。","en":"Security service type. Please enter a specific service type, if you purchase multiple security services.","dictionary":"belong=WAAP-MS-Ext|dict=waap_serviceType"}
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

type ListEntryDomainsResponseHeader struct {
}

func (s ListEntryDomainsResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ListEntryDomainsResponseHeader) GoString() string {
  return s.String()
}




