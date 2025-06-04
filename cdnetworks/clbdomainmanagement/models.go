package clbdomainmanagement

// This file is auto-generated, don't edit it. Thanks.
import (
  "github.com/alibabacloud-go/tea/tea"
)

type CreateClbDomainRequest struct {
  // {"en":"The zone id", "zh_CN":"zone的ID"}
  ZoneId *int `json:"zoneId,omitempty" xml:"zoneId,omitempty" require:"true"`
  // {"en":"domain name", "zh_CN":"clbDomain名称"}
  DomainName *string `json:"domainName,omitempty" xml:"domainName,omitempty" require:"true"`
  // {"en":"Dynamic Selection Type,0:GSLB = Shortest lantency, 1:RR = Round Robin, 2:Random = Random", "zh_CN":"动态选择类型, 0:GSLB = Shortest lantency, 1:RR = Round Robin, 2:Random = Random"}
  LoadBalancingType *int `json:"loadBalancingType,omitempty" xml:"loadBalancingType,omitempty" require:"true"`
  // {"en":"Portal CLB Basic Settings > Number of answer", "zh_CN":"Portal CLB 基础设置 > Number of answer"}
  AnswerCount *int `json:"answerCount,omitempty" xml:"answerCount,omitempty" require:"true"`
  // {"en":"TTL", "zh_CN":"TTL"}
  AnswerTtl *int `json:"answerTtl,omitempty" xml:"answerTtl,omitempty" require:"true"`
  // {"en":"health checker, can be null", "zh_CN":"监控检查,可以为null"}
  Healthchecker *CreateClbDomainRequestHealthchecker `json:"healthchecker,omitempty" xml:"healthchecker,omitempty" type:"Struct"`
  // {"en":"delay start time(second), only use when healthchecker not null", "zh_CN":"延迟开始时间,在healthchecker不为空时使用"}
  DelayStartTime *int `json:"delayStartTime,omitempty" xml:"delayStartTime,omitempty"`
  // {"en":"slow Start Period(second), only use when healthchecker not null", "zh_CN":"慢启动期,在healthchecker不为空时使用"}
  SlowStartPeriod *int `json:"slowStartPeriod,omitempty" xml:"slowStartPeriod,omitempty"`
  // {"en":"domain vip list, can be empty []", "zh_CN":"domain的VIP列表, 可为空[]"}
  DomainVips []*CreateClbDomainRequestDomainVips `json:"domainVips,omitempty" xml:"domainVips,omitempty" require:"true" type:"Repeated"`
  // {"en":"policy config, can be empty '[]'", "zh_CN":"策略配置，可以为空 []"}
  Policies []*CreateClbDomainRequestPolicies `json:"policies,omitempty" xml:"policies,omitempty" require:"true" type:"Repeated"`
  // {"en":"etc policy config, default = null(Dyanamic Selection From All IPs), only one need", "zh_CN":"Etc策略配置, default = null(Dyanamic Selection From All IPs)，只能配置一个"}
  EtcPolicies []*CreateClbDomainRequestEtcPolicies `json:"etcPolicies,omitempty" xml:"etcPolicies,omitempty" require:"true" type:"Repeated"`
}

func (s CreateClbDomainRequest) String() string {
  return tea.Prettify(s)
}

func (s CreateClbDomainRequest) GoString() string {
  return s.String()
}

func (s *CreateClbDomainRequest) SetZoneId(v int) *CreateClbDomainRequest {
  s.ZoneId = &v
  return s
}

func (s *CreateClbDomainRequest) SetDomainName(v string) *CreateClbDomainRequest {
  s.DomainName = &v
  return s
}

func (s *CreateClbDomainRequest) SetLoadBalancingType(v int) *CreateClbDomainRequest {
  s.LoadBalancingType = &v
  return s
}

func (s *CreateClbDomainRequest) SetAnswerCount(v int) *CreateClbDomainRequest {
  s.AnswerCount = &v
  return s
}

func (s *CreateClbDomainRequest) SetAnswerTtl(v int) *CreateClbDomainRequest {
  s.AnswerTtl = &v
  return s
}

func (s *CreateClbDomainRequest) SetHealthchecker(v *CreateClbDomainRequestHealthchecker) *CreateClbDomainRequest {
  s.Healthchecker = v
  return s
}

func (s *CreateClbDomainRequest) SetDelayStartTime(v int) *CreateClbDomainRequest {
  s.DelayStartTime = &v
  return s
}

func (s *CreateClbDomainRequest) SetSlowStartPeriod(v int) *CreateClbDomainRequest {
  s.SlowStartPeriod = &v
  return s
}

func (s *CreateClbDomainRequest) SetDomainVips(v []*CreateClbDomainRequestDomainVips) *CreateClbDomainRequest {
  s.DomainVips = v
  return s
}

func (s *CreateClbDomainRequest) SetPolicies(v []*CreateClbDomainRequestPolicies) *CreateClbDomainRequest {
  s.Policies = v
  return s
}

func (s *CreateClbDomainRequest) SetEtcPolicies(v []*CreateClbDomainRequestEtcPolicies) *CreateClbDomainRequest {
  s.EtcPolicies = v
  return s
}

type CreateClbDomainRequestHealthchecker struct {
  // {"en":"health checker id", "zh_CN":"监控检查ID"}
  Id *int `json:"id,omitempty" xml:"id,omitempty" require:"true"`
}

func (s CreateClbDomainRequestHealthchecker) String() string {
  return tea.Prettify(s)
}

func (s CreateClbDomainRequestHealthchecker) GoString() string {
  return s.String()
}

func (s *CreateClbDomainRequestHealthchecker) SetId(v int) *CreateClbDomainRequestHealthchecker {
  s.Id = &v
  return s
}

type CreateClbDomainRequestDomainVips struct     {
  // {"en":"Server's ID, get from servers API, If healthchecker is set, the server must also have healthchecker set", "zh_CN":"Server的ID，从serversAPI获取，如果设置了healthchecker，server也必须有设置的healthchecker"}
  ServerId *int `json:"serverId,omitempty" xml:"serverId,omitempty" require:"true"`
}

func (s CreateClbDomainRequestDomainVips) String() string {
  return tea.Prettify(s)
}

func (s CreateClbDomainRequestDomainVips) GoString() string {
  return s.String()
}

func (s *CreateClbDomainRequestDomainVips) SetServerId(v int) *CreateClbDomainRequestDomainVips {
  s.ServerId = &v
  return s
}

type CreateClbDomainRequestPolicies struct     {
  // {"en":"queue number", "zh_CN":"排序号"}
  Sequence *int `json:"sequence,omitempty" xml:"sequence,omitempty" require:"true"`
  // {"en":"policy name", "zh_CN":"策略名称"}
  PolicyName *string `json:"policyName,omitempty" xml:"policyName,omitempty"`
  // {"en":"enable", "zh_CN":"是否可用"}
  IsEnabled *bool `json:"isEnabled,omitempty" xml:"isEnabled,omitempty" require:"true"`
  // {"en":"condition", "zh_CN":"条件"}
  Condition *CreateClbDomainRequestPoliciesCondition `json:"condition,omitempty" xml:"condition,omitempty" require:"true" type:"Struct"`
  // {"en":"action, only one need", "zh_CN":"行动,只能配置一个"}
  Actions []*CreateClbDomainRequestPoliciesActions `json:"actions,omitempty" xml:"actions,omitempty" require:"true" type:"Repeated"`
}

func (s CreateClbDomainRequestPolicies) String() string {
  return tea.Prettify(s)
}

func (s CreateClbDomainRequestPolicies) GoString() string {
  return s.String()
}

func (s *CreateClbDomainRequestPolicies) SetSequence(v int) *CreateClbDomainRequestPolicies {
  s.Sequence = &v
  return s
}

func (s *CreateClbDomainRequestPolicies) SetPolicyName(v string) *CreateClbDomainRequestPolicies {
  s.PolicyName = &v
  return s
}

func (s *CreateClbDomainRequestPolicies) SetIsEnabled(v bool) *CreateClbDomainRequestPolicies {
  s.IsEnabled = &v
  return s
}

func (s *CreateClbDomainRequestPolicies) SetCondition(v *CreateClbDomainRequestPoliciesCondition) *CreateClbDomainRequestPolicies {
  s.Condition = v
  return s
}

func (s *CreateClbDomainRequestPolicies) SetActions(v []*CreateClbDomainRequestPoliciesActions) *CreateClbDomainRequestPolicies {
  s.Actions = v
  return s
}

type CreateClbDomainRequestPoliciesCondition struct {
  // {"en":"condition type, 0:IP RANGE, 1: GEO DATA, 2 Ratio(%)", "zh_CN":"条件类型, 0:IP RANGE, 1: GEO DATA, 2 Ratio(%)"}
  ConditionType *int `json:"conditionType,omitempty" xml:"conditionType,omitempty" require:"true"`
  // {"en":"IP range,comma separated,only useful if condition type is 0(ip range), eg: 1.2.3.4,1.5.6.0/24", "zh_CN":"IP范围,逗号分隔,只有在contidion type为0(ip range)时有用, 例:1.2.3.4,1.5.6.0/24"}
  IpRange *string `json:"ipRange,omitempty" xml:"ipRange,omitempty"`
  // {"en":"invert, only useful if condition type is 0(ip range) or 1(geo data) ", "zh_CN":"反转, 只有在contidion type为0(ip range)或1(geo data)时有用"}
  Invert *bool `json:"invert,omitempty" xml:"invert,omitempty"`
  // {"en":"region id list,only useful if condition type is 1(geo data), get by regions api", "zh_CN":"地理位置的ID列表，只有在contidion type为1(geo data)时有用，从regions接口获取"}
  Region []*string `json:"region,omitempty" xml:"region,omitempty" type:"Repeated"`
  // {"en":"config view id list,only useful if condition type is 1(geo data), get by cvs api", "zh_CN":"配置视图ID列表，只有在contidion type为1(geo data)时有用，从cvs接口获取"}
  Cv []*string `json:"cv,omitempty" xml:"cv,omitempty" type:"Repeated"`
  // {"en":"isp id list,only useful if condition type is 1(geo data), get by isps api", "zh_CN":"供应商ID列表，只有在contidion type为1(geo data)时有用，从isps接口获取"}
  Isp []*string `json:"isp,omitempty" xml:"isp,omitempty" type:"Repeated"`
  // {"en":"as list, only useful if condition type is 1(geo data)", "zh_CN":"as列表，只有在contidion type为1(geo data)时有用"}
  Asn []*int `json:"asn,omitempty" xml:"asn,omitempty" type:"Repeated"`
  // {"en":"percent,Only use conditionType = 2(Ratio(%)), eg: 70", "zh_CN":"比例，只有在conditionType=2 (Ratio%)时有用，如: 70"}
  Percent *int `json:"percent,omitempty" xml:"percent,omitempty"`
}

func (s CreateClbDomainRequestPoliciesCondition) String() string {
  return tea.Prettify(s)
}

func (s CreateClbDomainRequestPoliciesCondition) GoString() string {
  return s.String()
}

func (s *CreateClbDomainRequestPoliciesCondition) SetConditionType(v int) *CreateClbDomainRequestPoliciesCondition {
  s.ConditionType = &v
  return s
}

func (s *CreateClbDomainRequestPoliciesCondition) SetIpRange(v string) *CreateClbDomainRequestPoliciesCondition {
  s.IpRange = &v
  return s
}

func (s *CreateClbDomainRequestPoliciesCondition) SetInvert(v bool) *CreateClbDomainRequestPoliciesCondition {
  s.Invert = &v
  return s
}

func (s *CreateClbDomainRequestPoliciesCondition) SetRegion(v []*string) *CreateClbDomainRequestPoliciesCondition {
  s.Region = v
  return s
}

func (s *CreateClbDomainRequestPoliciesCondition) SetCv(v []*string) *CreateClbDomainRequestPoliciesCondition {
  s.Cv = v
  return s
}

func (s *CreateClbDomainRequestPoliciesCondition) SetIsp(v []*string) *CreateClbDomainRequestPoliciesCondition {
  s.Isp = v
  return s
}

func (s *CreateClbDomainRequestPoliciesCondition) SetAsn(v []*int) *CreateClbDomainRequestPoliciesCondition {
  s.Asn = v
  return s
}

func (s *CreateClbDomainRequestPoliciesCondition) SetPercent(v int) *CreateClbDomainRequestPoliciesCondition {
  s.Percent = &v
  return s
}

type CreateClbDomainRequestPoliciesActions struct     {
  // {"en":"action type, 0:deny, 1:vip, 2:cname,3:a", "zh_CN":"行动类别: 0:deny, 1:vip, 2:cname,3:a"}
  ActionType *int `json:"actionType,omitempty" xml:"actionType,omitempty" require:"true"`
  // {"en":"vip ID, only useful if actionType=1(vip), get in domainVips serverId", "zh_CN":"VIP ID,只有在action type为1(vip)时有用，是属性domainVips中的serverId值"}
  Vips []*int `json:"vips,omitempty" xml:"vips,omitempty" require:"true" type:"Repeated"`
  // {"en":"only useful if actionType=3(a) or actionType=2(cname), else is null", "zh_CN":"只有在actionType=3(a)或actionType=2(cname)时有用，或者为null"}
  Answer []*CreateClbDomainRequestPoliciesActionsAnswer `json:"answer,omitempty" xml:"answer,omitempty" require:"true" type:"Repeated"`
}

func (s CreateClbDomainRequestPoliciesActions) String() string {
  return tea.Prettify(s)
}

func (s CreateClbDomainRequestPoliciesActions) GoString() string {
  return s.String()
}

func (s *CreateClbDomainRequestPoliciesActions) SetActionType(v int) *CreateClbDomainRequestPoliciesActions {
  s.ActionType = &v
  return s
}

func (s *CreateClbDomainRequestPoliciesActions) SetVips(v []*int) *CreateClbDomainRequestPoliciesActions {
  s.Vips = v
  return s
}

func (s *CreateClbDomainRequestPoliciesActions) SetAnswer(v []*CreateClbDomainRequestPoliciesActionsAnswer) *CreateClbDomainRequestPoliciesActions {
  s.Answer = v
  return s
}

type CreateClbDomainRequestPoliciesActionsAnswer struct     {
  // {"en":"when actionType=3,this is IP, if actionType=2,that is domain name,eg: 1.2.3.4 or test.com", "zh_CN":"当actionType=3时，值为IP地址，当actionType=2时，值为域名地址"}
  Data *string `json:"data,omitempty" xml:"data,omitempty" require:"true"`
  // {"en":"TTL", "zh_CN":"TTL"}
  Ttl *int `json:"ttl,omitempty" xml:"ttl,omitempty" require:"true"`
}

func (s CreateClbDomainRequestPoliciesActionsAnswer) String() string {
  return tea.Prettify(s)
}

func (s CreateClbDomainRequestPoliciesActionsAnswer) GoString() string {
  return s.String()
}

func (s *CreateClbDomainRequestPoliciesActionsAnswer) SetData(v string) *CreateClbDomainRequestPoliciesActionsAnswer {
  s.Data = &v
  return s
}

func (s *CreateClbDomainRequestPoliciesActionsAnswer) SetTtl(v int) *CreateClbDomainRequestPoliciesActionsAnswer {
  s.Ttl = &v
  return s
}

type CreateClbDomainRequestEtcPolicies struct     {
  // {"en":"action", "zh_CN":"行动"}
  Actions []*CreateClbDomainRequestEtcPoliciesActions `json:"actions,omitempty" xml:"actions,omitempty" require:"true" type:"Repeated"`
}

func (s CreateClbDomainRequestEtcPolicies) String() string {
  return tea.Prettify(s)
}

func (s CreateClbDomainRequestEtcPolicies) GoString() string {
  return s.String()
}

func (s *CreateClbDomainRequestEtcPolicies) SetActions(v []*CreateClbDomainRequestEtcPoliciesActions) *CreateClbDomainRequestEtcPolicies {
  s.Actions = v
  return s
}

type CreateClbDomainRequestEtcPoliciesActions struct     {
  // {"en":"action type, 0:deny, 2:cname,3:a", "zh_CN":"行动类别: 0:deny, 2:cname,3:a"}
  ActionType *int `json:"actionType,omitempty" xml:"actionType,omitempty" require:"true"`
  // {"en":"only useful if actionType=3(a) or actionType=2(cname), default is null", "zh_CN":"只有在actionType=3(a)或actionType=2(cname)时有用,默认为null"}
  Answer []*CreateClbDomainRequestEtcPoliciesActionsAnswer `json:"answer,omitempty" xml:"answer,omitempty" require:"true" type:"Repeated"`
}

func (s CreateClbDomainRequestEtcPoliciesActions) String() string {
  return tea.Prettify(s)
}

func (s CreateClbDomainRequestEtcPoliciesActions) GoString() string {
  return s.String()
}

func (s *CreateClbDomainRequestEtcPoliciesActions) SetActionType(v int) *CreateClbDomainRequestEtcPoliciesActions {
  s.ActionType = &v
  return s
}

func (s *CreateClbDomainRequestEtcPoliciesActions) SetAnswer(v []*CreateClbDomainRequestEtcPoliciesActionsAnswer) *CreateClbDomainRequestEtcPoliciesActions {
  s.Answer = v
  return s
}

type CreateClbDomainRequestEtcPoliciesActionsAnswer struct     {
  // {"en":"when actionType=3,this is IP, if actionType=2,that is domain name,eg: 1.2.3.4 or test.com", "zh_CN":"当actionType=3时，值为IP地址，当actionType=2时，值为域名地址"}
  Data *string `json:"data,omitempty" xml:"data,omitempty" require:"true"`
  // {"en":"TTL", "zh_CN":"TTL"}
  Ttl *int `json:"ttl,omitempty" xml:"ttl,omitempty" require:"true"`
}

func (s CreateClbDomainRequestEtcPoliciesActionsAnswer) String() string {
  return tea.Prettify(s)
}

func (s CreateClbDomainRequestEtcPoliciesActionsAnswer) GoString() string {
  return s.String()
}

func (s *CreateClbDomainRequestEtcPoliciesActionsAnswer) SetData(v string) *CreateClbDomainRequestEtcPoliciesActionsAnswer {
  s.Data = &v
  return s
}

func (s *CreateClbDomainRequestEtcPoliciesActionsAnswer) SetTtl(v int) *CreateClbDomainRequestEtcPoliciesActionsAnswer {
  s.Ttl = &v
  return s
}

type CreateClbDomainResponse struct {
  // {"en":"return data", "zh_CN":"返回值"}
  Data *CreateClbDomainResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
  // {"en":"code,success is 0", "zh_CN":"状态码,成功为0"}
  Code *int `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"error message or Success", "zh_CN":"错误信息或Success"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s CreateClbDomainResponse) String() string {
  return tea.Prettify(s)
}

func (s CreateClbDomainResponse) GoString() string {
  return s.String()
}

func (s *CreateClbDomainResponse) SetData(v *CreateClbDomainResponseData) *CreateClbDomainResponse {
  s.Data = v
  return s
}

func (s *CreateClbDomainResponse) SetCode(v int) *CreateClbDomainResponse {
  s.Code = &v
  return s
}

func (s *CreateClbDomainResponse) SetMessage(v string) *CreateClbDomainResponse {
  s.Message = &v
  return s
}

type CreateClbDomainResponseData struct {
  // {'en':'customer info','zh_CN':'客户信息'}
  Customer *CreateClbDomainResponseDataCustomer `json:"customer,omitempty" xml:"customer,omitempty" require:"true" type:"Struct"`
  // {"en":"The zone id", "zh_CN":"zone的ID"}
  ZoneId *int `json:"zoneId,omitempty" xml:"zoneId,omitempty" require:"true"`
  // {"en":"The domain id", "zh_CN":"clbDomain的ID"}
  DomainId *int `json:"domainId,omitempty" xml:"domainId,omitempty" require:"true"`
  // {"en":"zone name", "zh_CN":"域名"}
  ZoneName *string `json:"zoneName,omitempty" xml:"zoneName,omitempty" require:"true"`
  // {"en":"domain name", "zh_CN":"clbDomain名称"}
  DomainName *string `json:"domainName,omitempty" xml:"domainName,omitempty" require:"true"`
  // {"en":"The unicode name of zone", "zh_CN":"Zone的Unicode名称"}
  ZoneUnicode *string `json:"zoneUnicode,omitempty" xml:"zoneUnicode,omitempty" require:"true"`
  // {"en":"The unicode name of domain", "zh_CN":"clbDomain的Unicode名称"}
  DomainUnicode *string `json:"domainUnicode,omitempty" xml:"domainUnicode,omitempty" require:"true"`
  // {"en":"STATUS CODE ( 0: NEW OR MODIFY ; 1:PENDING TO STAGE; 2: IN STAGE; 3: PENDING TO PRODUCTION;4: IN PRODUCTION; -2: ERROR TO STAGE; -4: ERROR TO PRODUCTION )", "zh_CN":"状态码（0: NEW; 1:PENDING TO STAGE; 2: IN STAGE; 3: PENDING TO PRODUCTION;4: IN PRODUCTION; -2: ERROR TO STAGE; -4: ERROR TO PRODUCTION）"}
  ZoneStatusCode *int `json:"zoneStatusCode,omitempty" xml:"zoneStatusCode,omitempty" require:"true"`
  // {"en":"status code description", "zh_CN":"状态码说明"}
  ZoneStatusText *string `json:"zoneStatusText,omitempty" xml:"zoneStatusText,omitempty" require:"true"`
  // {"en":"Published time, UTC time(eg: 2012-11-12T02:08:34Z )", "zh_CN":"发布时间，UTC时间(如: 2012-11-12T02:08:34Z )"}
  ZoneDatePublished *string `json:"zoneDatePublished,omitempty" xml:"zoneDatePublished,omitempty" require:"true"`
  // {"en":"create time, UTC time(eg: 2012-11-12T02:08:34Z )", "zh_CN":"创建时间，UTC时间(如: 2012-11-12T02:08:34Z )"}
  DateCreated *string `json:"dateCreated,omitempty" xml:"dateCreated,omitempty" require:"true"`
  // {"en":"modify time, UTC time(eg: 2012-11-12T02:08:34Z )", "zh_CN":"修改时间，UTC时间(如: 2012-11-12T02:08:34Z )"}
  DateModified *string `json:"dateModified,omitempty" xml:"dateModified,omitempty" require:"true"`
  // {"en":"Dynamic Selection Type,0:GSLB = Shortest lantency, 1:RR = Round Robin, 2:Random = Random", "zh_CN":"动态选择类型, 0:GSLB = Shortest lantency, 1:RR = Round Robin, 2:Random = Random"}
  LoadBalancingType *int `json:"loadBalancingType,omitempty" xml:"loadBalancingType,omitempty" require:"true"`
  // {"en":"Dynamic Selection Type Text", "zh_CN":"动态选择类型名称"}
  LoadBalancingTypeText *string `json:"loadBalancingTypeText,omitempty" xml:"loadBalancingTypeText,omitempty" require:"true"`
  // {"en":"Portal CLB Basic Settings > Number of answer", "zh_CN":"Portal CLB 基础设置 > Number of answer"}
  AnswerCount *int `json:"answerCount,omitempty" xml:"answerCount,omitempty" require:"true"`
  // {"en":"TTL", "zh_CN":"TTL"}
  AnswerTtl *int `json:"answerTtl,omitempty" xml:"answerTtl,omitempty" require:"true"`
  // {"en":"health checker, can be null", "zh_CN":"监控检查,可以为null"}
  Healthchecker *CreateClbDomainResponseDataHealthchecker `json:"healthchecker,omitempty" xml:"healthchecker,omitempty" require:"true" type:"Struct"`
  // {"en":"delay start time(second), only use when healthchecker not null", "zh_CN":"延迟开始时间,在healthchecker不为空时使用"}
  DelayStartTime *int `json:"delayStartTime,omitempty" xml:"delayStartTime,omitempty" require:"true"`
  // {"en":"slow Start Period(second), only use when healthchecker not null", "zh_CN":"慢启动期,在healthchecker不为空时使用"}
  SlowStartPeriod *int `json:"slowStartPeriod,omitempty" xml:"slowStartPeriod,omitempty" require:"true"`
  // {"en":"domain vip list", "zh_CN":"domain的VIP列表"}
  DomainVips []*CreateClbDomainResponseDataDomainVips `json:"domainVips,omitempty" xml:"domainVips,omitempty" require:"true" type:"Repeated"`
  // {"en":"policy config", "zh_CN":"策略配置"}
  Policies []*CreateClbDomainResponseDataPolicies `json:"policies,omitempty" xml:"policies,omitempty" require:"true" type:"Repeated"`
  // {"en":"etc policy config, default = null(Dyanamic Selection From All IPs)", "zh_CN":"Etc策略配置, default = null(Dyanamic Selection From All IPs)"}
  EtcPolicies []*CreateClbDomainResponseDataEtcPolicies `json:"etcPolicies,omitempty" xml:"etcPolicies,omitempty" require:"true" type:"Repeated"`
}

func (s CreateClbDomainResponseData) String() string {
  return tea.Prettify(s)
}

func (s CreateClbDomainResponseData) GoString() string {
  return s.String()
}

func (s *CreateClbDomainResponseData) SetCustomer(v *CreateClbDomainResponseDataCustomer) *CreateClbDomainResponseData {
  s.Customer = v
  return s
}

func (s *CreateClbDomainResponseData) SetZoneId(v int) *CreateClbDomainResponseData {
  s.ZoneId = &v
  return s
}

func (s *CreateClbDomainResponseData) SetDomainId(v int) *CreateClbDomainResponseData {
  s.DomainId = &v
  return s
}

func (s *CreateClbDomainResponseData) SetZoneName(v string) *CreateClbDomainResponseData {
  s.ZoneName = &v
  return s
}

func (s *CreateClbDomainResponseData) SetDomainName(v string) *CreateClbDomainResponseData {
  s.DomainName = &v
  return s
}

func (s *CreateClbDomainResponseData) SetZoneUnicode(v string) *CreateClbDomainResponseData {
  s.ZoneUnicode = &v
  return s
}

func (s *CreateClbDomainResponseData) SetDomainUnicode(v string) *CreateClbDomainResponseData {
  s.DomainUnicode = &v
  return s
}

func (s *CreateClbDomainResponseData) SetZoneStatusCode(v int) *CreateClbDomainResponseData {
  s.ZoneStatusCode = &v
  return s
}

func (s *CreateClbDomainResponseData) SetZoneStatusText(v string) *CreateClbDomainResponseData {
  s.ZoneStatusText = &v
  return s
}

func (s *CreateClbDomainResponseData) SetZoneDatePublished(v string) *CreateClbDomainResponseData {
  s.ZoneDatePublished = &v
  return s
}

func (s *CreateClbDomainResponseData) SetDateCreated(v string) *CreateClbDomainResponseData {
  s.DateCreated = &v
  return s
}

func (s *CreateClbDomainResponseData) SetDateModified(v string) *CreateClbDomainResponseData {
  s.DateModified = &v
  return s
}

func (s *CreateClbDomainResponseData) SetLoadBalancingType(v int) *CreateClbDomainResponseData {
  s.LoadBalancingType = &v
  return s
}

func (s *CreateClbDomainResponseData) SetLoadBalancingTypeText(v string) *CreateClbDomainResponseData {
  s.LoadBalancingTypeText = &v
  return s
}

func (s *CreateClbDomainResponseData) SetAnswerCount(v int) *CreateClbDomainResponseData {
  s.AnswerCount = &v
  return s
}

func (s *CreateClbDomainResponseData) SetAnswerTtl(v int) *CreateClbDomainResponseData {
  s.AnswerTtl = &v
  return s
}

func (s *CreateClbDomainResponseData) SetHealthchecker(v *CreateClbDomainResponseDataHealthchecker) *CreateClbDomainResponseData {
  s.Healthchecker = v
  return s
}

func (s *CreateClbDomainResponseData) SetDelayStartTime(v int) *CreateClbDomainResponseData {
  s.DelayStartTime = &v
  return s
}

func (s *CreateClbDomainResponseData) SetSlowStartPeriod(v int) *CreateClbDomainResponseData {
  s.SlowStartPeriod = &v
  return s
}

func (s *CreateClbDomainResponseData) SetDomainVips(v []*CreateClbDomainResponseDataDomainVips) *CreateClbDomainResponseData {
  s.DomainVips = v
  return s
}

func (s *CreateClbDomainResponseData) SetPolicies(v []*CreateClbDomainResponseDataPolicies) *CreateClbDomainResponseData {
  s.Policies = v
  return s
}

func (s *CreateClbDomainResponseData) SetEtcPolicies(v []*CreateClbDomainResponseDataEtcPolicies) *CreateClbDomainResponseData {
  s.EtcPolicies = v
  return s
}

type CreateClbDomainResponseDataCustomer struct {
  // {'en':'customer name','zh_CN':'客户名称'}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
}

func (s CreateClbDomainResponseDataCustomer) String() string {
  return tea.Prettify(s)
}

func (s CreateClbDomainResponseDataCustomer) GoString() string {
  return s.String()
}

func (s *CreateClbDomainResponseDataCustomer) SetName(v string) *CreateClbDomainResponseDataCustomer {
  s.Name = &v
  return s
}

type CreateClbDomainResponseDataHealthchecker struct {
  // {"en":"health checker id", "zh_CN":"监控检查ID"}
  Id *int `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {"en":"health checker name", "zh_CN":"监控检查名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
}

func (s CreateClbDomainResponseDataHealthchecker) String() string {
  return tea.Prettify(s)
}

func (s CreateClbDomainResponseDataHealthchecker) GoString() string {
  return s.String()
}

func (s *CreateClbDomainResponseDataHealthchecker) SetId(v int) *CreateClbDomainResponseDataHealthchecker {
  s.Id = &v
  return s
}

func (s *CreateClbDomainResponseDataHealthchecker) SetName(v string) *CreateClbDomainResponseDataHealthchecker {
  s.Name = &v
  return s
}

type CreateClbDomainResponseDataDomainVips struct     {
  // {"en":"Servergroup's ID", "zh_CN":"Servergroup的ID"}
  DomainVipId *int `json:"domainVipId,omitempty" xml:"domainVipId,omitempty" require:"true"`
  // {"en":"Servergroup's ID", "zh_CN":"Servergroup的ID"}
  ServergroupId *int `json:"servergroupId,omitempty" xml:"servergroupId,omitempty" require:"true"`
  // {"en":"server group's name", "zh_CN":"server group的名称"}
  ServergroupName *string `json:"servergroupName,omitempty" xml:"servergroupName,omitempty" require:"true"`
  // {"en":"Server's ID", "zh_CN":"Server的ID"}
  ServerId *int `json:"serverId,omitempty" xml:"serverId,omitempty" require:"true"`
  // {"en":"server name", "zh_CN":"服务名"}
  ServerName *string `json:"serverName,omitempty" xml:"serverName,omitempty" require:"true"`
  // {"en":"server group region id", "zh_CN":"server group的地理ID"}
  Region *string `json:"region,omitempty" xml:"region,omitempty" require:"true"`
  // {"en":"server group region name", "zh_CN":"server group的地理名称"}
  RegionName *string `json:"regionName,omitempty" xml:"regionName,omitempty" require:"true"`
  // {"en":"IP", "zh_CN":"ip地址"}
  Ip *string `json:"ip,omitempty" xml:"ip,omitempty" require:"true"`
  // {"en":"enable", "zh_CN":"是否可用"}
  IsEnabled *bool `json:"isEnabled,omitempty" xml:"isEnabled,omitempty" require:"true"`
  // {"en":"health checkers", "zh_CN":"监控检查列表"}
  Healthcheckers []*CreateClbDomainResponseDataDomainVipsHealthcheckers `json:"healthcheckers,omitempty" xml:"healthcheckers,omitempty" require:"true" type:"Repeated"`
}

func (s CreateClbDomainResponseDataDomainVips) String() string {
  return tea.Prettify(s)
}

func (s CreateClbDomainResponseDataDomainVips) GoString() string {
  return s.String()
}

func (s *CreateClbDomainResponseDataDomainVips) SetDomainVipId(v int) *CreateClbDomainResponseDataDomainVips {
  s.DomainVipId = &v
  return s
}

func (s *CreateClbDomainResponseDataDomainVips) SetServergroupId(v int) *CreateClbDomainResponseDataDomainVips {
  s.ServergroupId = &v
  return s
}

func (s *CreateClbDomainResponseDataDomainVips) SetServergroupName(v string) *CreateClbDomainResponseDataDomainVips {
  s.ServergroupName = &v
  return s
}

func (s *CreateClbDomainResponseDataDomainVips) SetServerId(v int) *CreateClbDomainResponseDataDomainVips {
  s.ServerId = &v
  return s
}

func (s *CreateClbDomainResponseDataDomainVips) SetServerName(v string) *CreateClbDomainResponseDataDomainVips {
  s.ServerName = &v
  return s
}

func (s *CreateClbDomainResponseDataDomainVips) SetRegion(v string) *CreateClbDomainResponseDataDomainVips {
  s.Region = &v
  return s
}

func (s *CreateClbDomainResponseDataDomainVips) SetRegionName(v string) *CreateClbDomainResponseDataDomainVips {
  s.RegionName = &v
  return s
}

func (s *CreateClbDomainResponseDataDomainVips) SetIp(v string) *CreateClbDomainResponseDataDomainVips {
  s.Ip = &v
  return s
}

func (s *CreateClbDomainResponseDataDomainVips) SetIsEnabled(v bool) *CreateClbDomainResponseDataDomainVips {
  s.IsEnabled = &v
  return s
}

func (s *CreateClbDomainResponseDataDomainVips) SetHealthcheckers(v []*CreateClbDomainResponseDataDomainVipsHealthcheckers) *CreateClbDomainResponseDataDomainVips {
  s.Healthcheckers = v
  return s
}

type CreateClbDomainResponseDataDomainVipsHealthcheckers struct     {
  // {"en":"health checker id", "zh_CN":"监控检查ID"}
  Id *int `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {"en":"health checker name", "zh_CN":"监控检查名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
}

func (s CreateClbDomainResponseDataDomainVipsHealthcheckers) String() string {
  return tea.Prettify(s)
}

func (s CreateClbDomainResponseDataDomainVipsHealthcheckers) GoString() string {
  return s.String()
}

func (s *CreateClbDomainResponseDataDomainVipsHealthcheckers) SetId(v int) *CreateClbDomainResponseDataDomainVipsHealthcheckers {
  s.Id = &v
  return s
}

func (s *CreateClbDomainResponseDataDomainVipsHealthcheckers) SetName(v string) *CreateClbDomainResponseDataDomainVipsHealthcheckers {
  s.Name = &v
  return s
}

type CreateClbDomainResponseDataPolicies struct     {
  // {"en":"queue number", "zh_CN":"排序号"}
  Sequence *int `json:"sequence,omitempty" xml:"sequence,omitempty" require:"true"`
  // {"en":"policy name", "zh_CN":"策略名称"}
  PolicyName *string `json:"policyName,omitempty" xml:"policyName,omitempty" require:"true"`
  // {"en":"enable", "zh_CN":"是否可用"}
  IsEnabled *bool `json:"isEnabled,omitempty" xml:"isEnabled,omitempty" require:"true"`
  // {"en":"condition", "zh_CN":"条件"}
  Condition *CreateClbDomainResponseDataPoliciesCondition `json:"condition,omitempty" xml:"condition,omitempty" require:"true" type:"Struct"`
  // {"en":"action", "zh_CN":"行动"}
  Actions []*CreateClbDomainResponseDataPoliciesActions `json:"actions,omitempty" xml:"actions,omitempty" require:"true" type:"Repeated"`
}

func (s CreateClbDomainResponseDataPolicies) String() string {
  return tea.Prettify(s)
}

func (s CreateClbDomainResponseDataPolicies) GoString() string {
  return s.String()
}

func (s *CreateClbDomainResponseDataPolicies) SetSequence(v int) *CreateClbDomainResponseDataPolicies {
  s.Sequence = &v
  return s
}

func (s *CreateClbDomainResponseDataPolicies) SetPolicyName(v string) *CreateClbDomainResponseDataPolicies {
  s.PolicyName = &v
  return s
}

func (s *CreateClbDomainResponseDataPolicies) SetIsEnabled(v bool) *CreateClbDomainResponseDataPolicies {
  s.IsEnabled = &v
  return s
}

func (s *CreateClbDomainResponseDataPolicies) SetCondition(v *CreateClbDomainResponseDataPoliciesCondition) *CreateClbDomainResponseDataPolicies {
  s.Condition = v
  return s
}

func (s *CreateClbDomainResponseDataPolicies) SetActions(v []*CreateClbDomainResponseDataPoliciesActions) *CreateClbDomainResponseDataPolicies {
  s.Actions = v
  return s
}

type CreateClbDomainResponseDataPoliciesCondition struct {
  // {"en":"The condition id", "zh_CN":"条件ID"}
  ConditionId *int `json:"conditionId,omitempty" xml:"conditionId,omitempty" require:"true"`
  // {"en":"condition type, 0:IP RANGE, 1: GEO DATA, 2 Ratio(%)", "zh_CN":"条件类型, 0:IP RANGE, 1: GEO DATA, 2 Ratio(%)"}
  ConditionType *int `json:"conditionType,omitempty" xml:"conditionType,omitempty" require:"true"`
  // {"en":"IP range,comma separated,only useful if condition type is 0(ip range), eg: 1.2.3.4,1.5.6.0/24", "zh_CN":"IP范围,逗号分隔,只有在contidion type为0(ip range)时有用, 例:1.2.3.4,1.5.6.0/24"}
  IpRange *string `json:"ipRange,omitempty" xml:"ipRange,omitempty" require:"true"`
  // {"en":"invert, only useful if condition type is 0(ip range) or 1(geo data) ", "zh_CN":"反转, 只有在contidion type为0(ip range)或1(geo data)时有用"}
  Invert *bool `json:"invert,omitempty" xml:"invert,omitempty" require:"true"`
  // {"en":"region id list,only useful if condition type is 1(geo data), get by regions api", "zh_CN":"地理位置的ID列表，只有在contidion type为1(geo data)时有用，从regions接口获取"}
  Region []*string `json:"region,omitempty" xml:"region,omitempty" require:"true" type:"Repeated"`
  // {"en":"config view id list,only useful if condition type is 1(geo data), get by cvs api", "zh_CN":"配置视图ID列表，只有在contidion type为1(geo data)时有用，从cvs接口获取"}
  Cv []*string `json:"cv,omitempty" xml:"cv,omitempty" require:"true" type:"Repeated"`
  // {"en":"isp id list,only useful if condition type is 1(geo data), get by isps api", "zh_CN":"供应商ID列表，只有在contidion type为1(geo data)时有用，从isps接口获取"}
  Isp []*string `json:"isp,omitempty" xml:"isp,omitempty" require:"true" type:"Repeated"`
  // {"en":"as list, only useful if condition type is 1(geo data)", "zh_CN":"as列表，只有在contidion type为1(geo data)时有用"}
  Asn []*int `json:"asn,omitempty" xml:"asn,omitempty" require:"true" type:"Repeated"`
  // {"en":"percent,Only use conditionType = 2(Ratio(%)), eg: 70", "zh_CN":"比例，只有在conditionType=2 (Ratio%)时有用，如: 70"}
  Percent *int `json:"percent,omitempty" xml:"percent,omitempty" require:"true"`
}

func (s CreateClbDomainResponseDataPoliciesCondition) String() string {
  return tea.Prettify(s)
}

func (s CreateClbDomainResponseDataPoliciesCondition) GoString() string {
  return s.String()
}

func (s *CreateClbDomainResponseDataPoliciesCondition) SetConditionId(v int) *CreateClbDomainResponseDataPoliciesCondition {
  s.ConditionId = &v
  return s
}

func (s *CreateClbDomainResponseDataPoliciesCondition) SetConditionType(v int) *CreateClbDomainResponseDataPoliciesCondition {
  s.ConditionType = &v
  return s
}

func (s *CreateClbDomainResponseDataPoliciesCondition) SetIpRange(v string) *CreateClbDomainResponseDataPoliciesCondition {
  s.IpRange = &v
  return s
}

func (s *CreateClbDomainResponseDataPoliciesCondition) SetInvert(v bool) *CreateClbDomainResponseDataPoliciesCondition {
  s.Invert = &v
  return s
}

func (s *CreateClbDomainResponseDataPoliciesCondition) SetRegion(v []*string) *CreateClbDomainResponseDataPoliciesCondition {
  s.Region = v
  return s
}

func (s *CreateClbDomainResponseDataPoliciesCondition) SetCv(v []*string) *CreateClbDomainResponseDataPoliciesCondition {
  s.Cv = v
  return s
}

func (s *CreateClbDomainResponseDataPoliciesCondition) SetIsp(v []*string) *CreateClbDomainResponseDataPoliciesCondition {
  s.Isp = v
  return s
}

func (s *CreateClbDomainResponseDataPoliciesCondition) SetAsn(v []*int) *CreateClbDomainResponseDataPoliciesCondition {
  s.Asn = v
  return s
}

func (s *CreateClbDomainResponseDataPoliciesCondition) SetPercent(v int) *CreateClbDomainResponseDataPoliciesCondition {
  s.Percent = &v
  return s
}

type CreateClbDomainResponseDataPoliciesActions struct     {
  // {"en":"queue number", "zh_CN":"排序号"}
  Sequence *int `json:"sequence,omitempty" xml:"sequence,omitempty" require:"true"`
  // {"en":"The action id", "zh_CN":"行动ID"}
  ActionId *int `json:"actionId,omitempty" xml:"actionId,omitempty" require:"true"`
  // {"en":"action type, 0:deny, 1:vip, 2:cname,3:a", "zh_CN":"行动类别: 0:deny, 1:vip, 2:cname,3:a"}
  ActionType *int `json:"actionType,omitempty" xml:"actionType,omitempty" require:"true"`
  // {"en":"vip ID, only useful if actionType=1(vip), get in domainVips serverId", "zh_CN":"VIP ID,只有在action type为1(vip)时有用，是属性domainVips中的serverId值"}
  Vips []*int `json:"vips,omitempty" xml:"vips,omitempty" require:"true" type:"Repeated"`
  // {"en":"only useful if actionType=3(a) or actionType=2(cname)", "zh_CN":"只有在actionType=3(a)或actionType=2(cname)时有用"}
  Answer []*CreateClbDomainResponseDataPoliciesActionsAnswer `json:"answer,omitempty" xml:"answer,omitempty" require:"true" type:"Repeated"`
}

func (s CreateClbDomainResponseDataPoliciesActions) String() string {
  return tea.Prettify(s)
}

func (s CreateClbDomainResponseDataPoliciesActions) GoString() string {
  return s.String()
}

func (s *CreateClbDomainResponseDataPoliciesActions) SetSequence(v int) *CreateClbDomainResponseDataPoliciesActions {
  s.Sequence = &v
  return s
}

func (s *CreateClbDomainResponseDataPoliciesActions) SetActionId(v int) *CreateClbDomainResponseDataPoliciesActions {
  s.ActionId = &v
  return s
}

func (s *CreateClbDomainResponseDataPoliciesActions) SetActionType(v int) *CreateClbDomainResponseDataPoliciesActions {
  s.ActionType = &v
  return s
}

func (s *CreateClbDomainResponseDataPoliciesActions) SetVips(v []*int) *CreateClbDomainResponseDataPoliciesActions {
  s.Vips = v
  return s
}

func (s *CreateClbDomainResponseDataPoliciesActions) SetAnswer(v []*CreateClbDomainResponseDataPoliciesActionsAnswer) *CreateClbDomainResponseDataPoliciesActions {
  s.Answer = v
  return s
}

type CreateClbDomainResponseDataPoliciesActionsAnswer struct     {
  // {"en":"actionType=3, A or actionType=2, CNAME", "zh_CN":"actionType=3,则值为A,actionType=2,则值为CNAME"}
  Rtype *string `json:"rtype,omitempty" xml:"rtype,omitempty" require:"true"`
  // {"en":"when actionType=3,this is IP, if actionType=2,that is domain name,eg: 1.2.3.4 or test.com", "zh_CN":"当actionType=3时，值为IP地址，当actionType=2时，值为域名地址"}
  Data *string `json:"data,omitempty" xml:"data,omitempty" require:"true"`
  // {"en":"TTL", "zh_CN":"TTL"}
  Ttl *int `json:"ttl,omitempty" xml:"ttl,omitempty" require:"true"`
}

func (s CreateClbDomainResponseDataPoliciesActionsAnswer) String() string {
  return tea.Prettify(s)
}

func (s CreateClbDomainResponseDataPoliciesActionsAnswer) GoString() string {
  return s.String()
}

func (s *CreateClbDomainResponseDataPoliciesActionsAnswer) SetRtype(v string) *CreateClbDomainResponseDataPoliciesActionsAnswer {
  s.Rtype = &v
  return s
}

func (s *CreateClbDomainResponseDataPoliciesActionsAnswer) SetData(v string) *CreateClbDomainResponseDataPoliciesActionsAnswer {
  s.Data = &v
  return s
}

func (s *CreateClbDomainResponseDataPoliciesActionsAnswer) SetTtl(v int) *CreateClbDomainResponseDataPoliciesActionsAnswer {
  s.Ttl = &v
  return s
}

type CreateClbDomainResponseDataEtcPolicies struct     {
  // {"en":"policy name(always)", "zh_CN":"策略名称(always)"}
  PolicyName *string `json:"policyName,omitempty" xml:"policyName,omitempty" require:"true"`
  // {"en":"action", "zh_CN":"行动"}
  Actions []*CreateClbDomainResponseDataEtcPoliciesActions `json:"actions,omitempty" xml:"actions,omitempty" require:"true" type:"Repeated"`
}

func (s CreateClbDomainResponseDataEtcPolicies) String() string {
  return tea.Prettify(s)
}

func (s CreateClbDomainResponseDataEtcPolicies) GoString() string {
  return s.String()
}

func (s *CreateClbDomainResponseDataEtcPolicies) SetPolicyName(v string) *CreateClbDomainResponseDataEtcPolicies {
  s.PolicyName = &v
  return s
}

func (s *CreateClbDomainResponseDataEtcPolicies) SetActions(v []*CreateClbDomainResponseDataEtcPoliciesActions) *CreateClbDomainResponseDataEtcPolicies {
  s.Actions = v
  return s
}

type CreateClbDomainResponseDataEtcPoliciesActions struct     {
  // {"en":"action type, 0:deny, 2:cname,3:a", "zh_CN":"行动类别: 0:deny, 2:cname,3:a"}
  ActionType *int `json:"actionType,omitempty" xml:"actionType,omitempty" require:"true"`
  // {"en":"only useful if actionType=3(a) or actionType=2(cname), default is null", "zh_CN":"只有在actionType=3(a)或actionType=2(cname)时有用,默认为null"}
  Answer []*CreateClbDomainResponseDataEtcPoliciesActionsAnswer `json:"answer,omitempty" xml:"answer,omitempty" require:"true" type:"Repeated"`
}

func (s CreateClbDomainResponseDataEtcPoliciesActions) String() string {
  return tea.Prettify(s)
}

func (s CreateClbDomainResponseDataEtcPoliciesActions) GoString() string {
  return s.String()
}

func (s *CreateClbDomainResponseDataEtcPoliciesActions) SetActionType(v int) *CreateClbDomainResponseDataEtcPoliciesActions {
  s.ActionType = &v
  return s
}

func (s *CreateClbDomainResponseDataEtcPoliciesActions) SetAnswer(v []*CreateClbDomainResponseDataEtcPoliciesActionsAnswer) *CreateClbDomainResponseDataEtcPoliciesActions {
  s.Answer = v
  return s
}

type CreateClbDomainResponseDataEtcPoliciesActionsAnswer struct     {
  // {"en":"actionType=3, A or actionType=2, CNAME", "zh_CN":"actionType=3,则值为A,actionType=2,则值为CNAME"}
  Rtype *string `json:"rtype,omitempty" xml:"rtype,omitempty" require:"true"`
  // {"en":"when actionType=3,this is IP, if actionType=2,that is domain name,eg: 1.2.3.4 or test.com", "zh_CN":"当actionType=3时，值为IP地址，当actionType=2时，值为域名地址"}
  Data *string `json:"data,omitempty" xml:"data,omitempty" require:"true"`
  // {"en":"TTL", "zh_CN":"TTL"}
  Ttl *int `json:"ttl,omitempty" xml:"ttl,omitempty" require:"true"`
}

func (s CreateClbDomainResponseDataEtcPoliciesActionsAnswer) String() string {
  return tea.Prettify(s)
}

func (s CreateClbDomainResponseDataEtcPoliciesActionsAnswer) GoString() string {
  return s.String()
}

func (s *CreateClbDomainResponseDataEtcPoliciesActionsAnswer) SetRtype(v string) *CreateClbDomainResponseDataEtcPoliciesActionsAnswer {
  s.Rtype = &v
  return s
}

func (s *CreateClbDomainResponseDataEtcPoliciesActionsAnswer) SetData(v string) *CreateClbDomainResponseDataEtcPoliciesActionsAnswer {
  s.Data = &v
  return s
}

func (s *CreateClbDomainResponseDataEtcPoliciesActionsAnswer) SetTtl(v int) *CreateClbDomainResponseDataEtcPoliciesActionsAnswer {
  s.Ttl = &v
  return s
}

type CreateClbDomainPaths struct {
}

func (s CreateClbDomainPaths) String() string {
  return tea.Prettify(s)
}

func (s CreateClbDomainPaths) GoString() string {
  return s.String()
}

type CreateClbDomainParameters struct {
}

func (s CreateClbDomainParameters) String() string {
  return tea.Prettify(s)
}

func (s CreateClbDomainParameters) GoString() string {
  return s.String()
}

type CreateClbDomainRequestHeader struct {
}

func (s CreateClbDomainRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s CreateClbDomainRequestHeader) GoString() string {
  return s.String()
}

type CreateClbDomainResponseHeader struct {
}

func (s CreateClbDomainResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s CreateClbDomainResponseHeader) GoString() string {
  return s.String()
}




type UpdateClbDomainByIDRequest struct {
  // {"en":"The zone id", "zh_CN":"zone的ID"}
  ZoneId *int `json:"zoneId,omitempty" xml:"zoneId,omitempty" require:"true"`
  // {"en":"domain name", "zh_CN":"clbDomain名称"}
  DomainName *string `json:"domainName,omitempty" xml:"domainName,omitempty" require:"true"`
  // {"en":"Dynamic Selection Type,0:GSLB = Shortest lantency, 1:RR = Round Robin, 2:Random = Random", "zh_CN":"动态选择类型, 0:GSLB = Shortest lantency, 1:RR = Round Robin, 2:Random = Random"}
  LoadBalancingType *int `json:"loadBalancingType,omitempty" xml:"loadBalancingType,omitempty" require:"true"`
  // {"en":"Portal CLB Basic Settings > Number of answer", "zh_CN":"Portal CLB 基础设置 > Number of answer"}
  AnswerCount *int `json:"answerCount,omitempty" xml:"answerCount,omitempty" require:"true"`
  // {"en":"TTL", "zh_CN":"TTL"}
  AnswerTtl *int `json:"answerTtl,omitempty" xml:"answerTtl,omitempty" require:"true"`
  // {"en":"health checker, can be null", "zh_CN":"监控检查,可以为null"}
  Healthchecker *UpdateClbDomainByIDRequestHealthchecker `json:"healthchecker,omitempty" xml:"healthchecker,omitempty" type:"Struct"`
  // {"en":"delay start time(second), only use when healthchecker not null", "zh_CN":"延迟开始时间,在healthchecker不为空时使用"}
  DelayStartTime *int `json:"delayStartTime,omitempty" xml:"delayStartTime,omitempty"`
  // {"en":"slow Start Period(second), only use when healthchecker not null", "zh_CN":"慢启动期,在healthchecker不为空时使用"}
  SlowStartPeriod *int `json:"slowStartPeriod,omitempty" xml:"slowStartPeriod,omitempty"`
  // {"en":"domain vip list, can be empty []", "zh_CN":"domain的VIP列表, 可为空[]"}
  DomainVips []*UpdateClbDomainByIDRequestDomainVips `json:"domainVips,omitempty" xml:"domainVips,omitempty" require:"true" type:"Repeated"`
  // {"en":"policy config, can be empty '[]'", "zh_CN":"策略配置，可以为空 []"}
  Policies []*UpdateClbDomainByIDRequestPolicies `json:"policies,omitempty" xml:"policies,omitempty" require:"true" type:"Repeated"`
  // {"en":"etc policy config, default = null(Dyanamic Selection From All IPs), only one need", "zh_CN":"Etc策略配置, default = null(Dyanamic Selection From All IPs)，只能配置一个"}
  EtcPolicies []*UpdateClbDomainByIDRequestEtcPolicies `json:"etcPolicies,omitempty" xml:"etcPolicies,omitempty" type:"Repeated"`
}

func (s UpdateClbDomainByIDRequest) String() string {
  return tea.Prettify(s)
}

func (s UpdateClbDomainByIDRequest) GoString() string {
  return s.String()
}

func (s *UpdateClbDomainByIDRequest) SetZoneId(v int) *UpdateClbDomainByIDRequest {
  s.ZoneId = &v
  return s
}

func (s *UpdateClbDomainByIDRequest) SetDomainName(v string) *UpdateClbDomainByIDRequest {
  s.DomainName = &v
  return s
}

func (s *UpdateClbDomainByIDRequest) SetLoadBalancingType(v int) *UpdateClbDomainByIDRequest {
  s.LoadBalancingType = &v
  return s
}

func (s *UpdateClbDomainByIDRequest) SetAnswerCount(v int) *UpdateClbDomainByIDRequest {
  s.AnswerCount = &v
  return s
}

func (s *UpdateClbDomainByIDRequest) SetAnswerTtl(v int) *UpdateClbDomainByIDRequest {
  s.AnswerTtl = &v
  return s
}

func (s *UpdateClbDomainByIDRequest) SetHealthchecker(v *UpdateClbDomainByIDRequestHealthchecker) *UpdateClbDomainByIDRequest {
  s.Healthchecker = v
  return s
}

func (s *UpdateClbDomainByIDRequest) SetDelayStartTime(v int) *UpdateClbDomainByIDRequest {
  s.DelayStartTime = &v
  return s
}

func (s *UpdateClbDomainByIDRequest) SetSlowStartPeriod(v int) *UpdateClbDomainByIDRequest {
  s.SlowStartPeriod = &v
  return s
}

func (s *UpdateClbDomainByIDRequest) SetDomainVips(v []*UpdateClbDomainByIDRequestDomainVips) *UpdateClbDomainByIDRequest {
  s.DomainVips = v
  return s
}

func (s *UpdateClbDomainByIDRequest) SetPolicies(v []*UpdateClbDomainByIDRequestPolicies) *UpdateClbDomainByIDRequest {
  s.Policies = v
  return s
}

func (s *UpdateClbDomainByIDRequest) SetEtcPolicies(v []*UpdateClbDomainByIDRequestEtcPolicies) *UpdateClbDomainByIDRequest {
  s.EtcPolicies = v
  return s
}

type UpdateClbDomainByIDRequestHealthchecker struct {
  // {"en":"health checker id", "zh_CN":"监控检查ID"}
  Id *int `json:"id,omitempty" xml:"id,omitempty" require:"true"`
}

func (s UpdateClbDomainByIDRequestHealthchecker) String() string {
  return tea.Prettify(s)
}

func (s UpdateClbDomainByIDRequestHealthchecker) GoString() string {
  return s.String()
}

func (s *UpdateClbDomainByIDRequestHealthchecker) SetId(v int) *UpdateClbDomainByIDRequestHealthchecker {
  s.Id = &v
  return s
}

type UpdateClbDomainByIDRequestDomainVips struct     {
  // {"en":"Server's ID, get from servers API, If healthchecker is set, the server must also have healthchecker set", "zh_CN":"Server的ID，从serversAPI获取，如果设置了healthchecker，server也必须有设置的healthchecker"}
  ServerId *int `json:"serverId,omitempty" xml:"serverId,omitempty" require:"true"`
}

func (s UpdateClbDomainByIDRequestDomainVips) String() string {
  return tea.Prettify(s)
}

func (s UpdateClbDomainByIDRequestDomainVips) GoString() string {
  return s.String()
}

func (s *UpdateClbDomainByIDRequestDomainVips) SetServerId(v int) *UpdateClbDomainByIDRequestDomainVips {
  s.ServerId = &v
  return s
}

type UpdateClbDomainByIDRequestPolicies struct     {
  // {"en":"queue number", "zh_CN":"排序号"}
  Sequence *int `json:"sequence,omitempty" xml:"sequence,omitempty" require:"true"`
  // {"en":"policy name", "zh_CN":"策略名称"}
  PolicyName *string `json:"policyName,omitempty" xml:"policyName,omitempty"`
  // {"en":"enable", "zh_CN":"是否可用"}
  IsEnabled *bool `json:"isEnabled,omitempty" xml:"isEnabled,omitempty" require:"true"`
  // {"en":"condition", "zh_CN":"条件"}
  Condition *UpdateClbDomainByIDRequestPoliciesCondition `json:"condition,omitempty" xml:"condition,omitempty" require:"true" type:"Struct"`
  // {"en":"action, only one need", "zh_CN":"行动,只能配置一个"}
  Actions []*UpdateClbDomainByIDRequestPoliciesActions `json:"actions,omitempty" xml:"actions,omitempty" require:"true" type:"Repeated"`
}

func (s UpdateClbDomainByIDRequestPolicies) String() string {
  return tea.Prettify(s)
}

func (s UpdateClbDomainByIDRequestPolicies) GoString() string {
  return s.String()
}

func (s *UpdateClbDomainByIDRequestPolicies) SetSequence(v int) *UpdateClbDomainByIDRequestPolicies {
  s.Sequence = &v
  return s
}

func (s *UpdateClbDomainByIDRequestPolicies) SetPolicyName(v string) *UpdateClbDomainByIDRequestPolicies {
  s.PolicyName = &v
  return s
}

func (s *UpdateClbDomainByIDRequestPolicies) SetIsEnabled(v bool) *UpdateClbDomainByIDRequestPolicies {
  s.IsEnabled = &v
  return s
}

func (s *UpdateClbDomainByIDRequestPolicies) SetCondition(v *UpdateClbDomainByIDRequestPoliciesCondition) *UpdateClbDomainByIDRequestPolicies {
  s.Condition = v
  return s
}

func (s *UpdateClbDomainByIDRequestPolicies) SetActions(v []*UpdateClbDomainByIDRequestPoliciesActions) *UpdateClbDomainByIDRequestPolicies {
  s.Actions = v
  return s
}

type UpdateClbDomainByIDRequestPoliciesCondition struct {
  // {"en":"condition type, 0:IP RANGE, 1: GEO DATA, 2 Ratio(%)", "zh_CN":"条件类型, 0:IP RANGE, 1: GEO DATA, 2 Ratio(%)"}
  ConditionType *int `json:"conditionType,omitempty" xml:"conditionType,omitempty" require:"true"`
  // {"en":"IP range,comma separated,only useful if condition type is 0(ip range), eg: 1.2.3.4,1.5.6.0/24", "zh_CN":"IP范围,逗号分隔,只有在contidion type为0(ip range)时有用, 例:1.2.3.4,1.5.6.0/24"}
  IpRange *string `json:"ipRange,omitempty" xml:"ipRange,omitempty"`
  // {"en":"invert, only useful if condition type is 0(ip range) or 1(geo data) ", "zh_CN":"反转, 只有在contidion type为0(ip range)或1(geo data)时有用"}
  Invert *bool `json:"invert,omitempty" xml:"invert,omitempty"`
  // {"en":"region id list,only useful if condition type is 1(geo data), get by regions api", "zh_CN":"地理位置的ID列表，只有在contidion type为1(geo data)时有用，从regions接口获取"}
  Region []*string `json:"region,omitempty" xml:"region,omitempty" type:"Repeated"`
  // {"en":"config view id list,only useful if condition type is 1(geo data), get by cvs api", "zh_CN":"配置视图ID列表，只有在contidion type为1(geo data)时有用，从cvs接口获取"}
  Cv []*string `json:"cv,omitempty" xml:"cv,omitempty" type:"Repeated"`
  // {"en":"isp id list,only useful if condition type is 1(geo data), get by isps api", "zh_CN":"供应商ID列表，只有在contidion type为1(geo data)时有用，从isps接口获取"}
  Isp []*string `json:"isp,omitempty" xml:"isp,omitempty" type:"Repeated"`
  // {"en":"as list, only useful if condition type is 1(geo data)", "zh_CN":"as列表，只有在contidion type为1(geo data)时有用"}
  Asn []*int `json:"asn,omitempty" xml:"asn,omitempty" type:"Repeated"`
  // {"en":"percent,Only use conditionType = 2(Ratio(%)), eg: 70", "zh_CN":"比例，只有在conditionType=2 (Ratio%)时有用，如: 70"}
  Percent *int `json:"percent,omitempty" xml:"percent,omitempty"`
}

func (s UpdateClbDomainByIDRequestPoliciesCondition) String() string {
  return tea.Prettify(s)
}

func (s UpdateClbDomainByIDRequestPoliciesCondition) GoString() string {
  return s.String()
}

func (s *UpdateClbDomainByIDRequestPoliciesCondition) SetConditionType(v int) *UpdateClbDomainByIDRequestPoliciesCondition {
  s.ConditionType = &v
  return s
}

func (s *UpdateClbDomainByIDRequestPoliciesCondition) SetIpRange(v string) *UpdateClbDomainByIDRequestPoliciesCondition {
  s.IpRange = &v
  return s
}

func (s *UpdateClbDomainByIDRequestPoliciesCondition) SetInvert(v bool) *UpdateClbDomainByIDRequestPoliciesCondition {
  s.Invert = &v
  return s
}

func (s *UpdateClbDomainByIDRequestPoliciesCondition) SetRegion(v []*string) *UpdateClbDomainByIDRequestPoliciesCondition {
  s.Region = v
  return s
}

func (s *UpdateClbDomainByIDRequestPoliciesCondition) SetCv(v []*string) *UpdateClbDomainByIDRequestPoliciesCondition {
  s.Cv = v
  return s
}

func (s *UpdateClbDomainByIDRequestPoliciesCondition) SetIsp(v []*string) *UpdateClbDomainByIDRequestPoliciesCondition {
  s.Isp = v
  return s
}

func (s *UpdateClbDomainByIDRequestPoliciesCondition) SetAsn(v []*int) *UpdateClbDomainByIDRequestPoliciesCondition {
  s.Asn = v
  return s
}

func (s *UpdateClbDomainByIDRequestPoliciesCondition) SetPercent(v int) *UpdateClbDomainByIDRequestPoliciesCondition {
  s.Percent = &v
  return s
}

type UpdateClbDomainByIDRequestPoliciesActions struct     {
  // {"en":"action type, 0:deny, 1:vip, 2:cname,3:a", "zh_CN":"行动类别: 0:deny, 1:vip, 2:cname,3:a"}
  ActionType *int `json:"actionType,omitempty" xml:"actionType,omitempty" require:"true"`
  // {"en":"vip ID, only useful if actionType=1(vip), get in domainVips serverId", "zh_CN":"VIP ID,只有在action type为1(vip)时有用，是属性domainVips中的serverId值"}
  Vips []*int `json:"vips,omitempty" xml:"vips,omitempty" require:"true" type:"Repeated"`
  // {"en":"only useful if actionType=3(a) or actionType=2(cname), else is null", "zh_CN":"只有在actionType=3(a)或actionType=2(cname)时有用，或者为null"}
  Answer []*UpdateClbDomainByIDRequestPoliciesActionsAnswer `json:"answer,omitempty" xml:"answer,omitempty" require:"true" type:"Repeated"`
}

func (s UpdateClbDomainByIDRequestPoliciesActions) String() string {
  return tea.Prettify(s)
}

func (s UpdateClbDomainByIDRequestPoliciesActions) GoString() string {
  return s.String()
}

func (s *UpdateClbDomainByIDRequestPoliciesActions) SetActionType(v int) *UpdateClbDomainByIDRequestPoliciesActions {
  s.ActionType = &v
  return s
}

func (s *UpdateClbDomainByIDRequestPoliciesActions) SetVips(v []*int) *UpdateClbDomainByIDRequestPoliciesActions {
  s.Vips = v
  return s
}

func (s *UpdateClbDomainByIDRequestPoliciesActions) SetAnswer(v []*UpdateClbDomainByIDRequestPoliciesActionsAnswer) *UpdateClbDomainByIDRequestPoliciesActions {
  s.Answer = v
  return s
}

type UpdateClbDomainByIDRequestPoliciesActionsAnswer struct     {
  // {"en":"when actionType=3,this is IP, if actionType=2,that is domain name,eg: 1.2.3.4 or test.com", "zh_CN":"当actionType=3时，值为IP地址，当actionType=2时，值为域名地址"}
  Data *string `json:"data,omitempty" xml:"data,omitempty" require:"true"`
  // {"en":"TTL", "zh_CN":"TTL"}
  Ttl *int `json:"ttl,omitempty" xml:"ttl,omitempty" require:"true"`
}

func (s UpdateClbDomainByIDRequestPoliciesActionsAnswer) String() string {
  return tea.Prettify(s)
}

func (s UpdateClbDomainByIDRequestPoliciesActionsAnswer) GoString() string {
  return s.String()
}

func (s *UpdateClbDomainByIDRequestPoliciesActionsAnswer) SetData(v string) *UpdateClbDomainByIDRequestPoliciesActionsAnswer {
  s.Data = &v
  return s
}

func (s *UpdateClbDomainByIDRequestPoliciesActionsAnswer) SetTtl(v int) *UpdateClbDomainByIDRequestPoliciesActionsAnswer {
  s.Ttl = &v
  return s
}

type UpdateClbDomainByIDRequestEtcPolicies struct     {
  // {"en":"action", "zh_CN":"行动"}
  Actions []*UpdateClbDomainByIDRequestEtcPoliciesActions `json:"actions,omitempty" xml:"actions,omitempty" require:"true" type:"Repeated"`
}

func (s UpdateClbDomainByIDRequestEtcPolicies) String() string {
  return tea.Prettify(s)
}

func (s UpdateClbDomainByIDRequestEtcPolicies) GoString() string {
  return s.String()
}

func (s *UpdateClbDomainByIDRequestEtcPolicies) SetActions(v []*UpdateClbDomainByIDRequestEtcPoliciesActions) *UpdateClbDomainByIDRequestEtcPolicies {
  s.Actions = v
  return s
}

type UpdateClbDomainByIDRequestEtcPoliciesActions struct     {
  // {"en":"action type, 0:deny, 2:cname,3:a", "zh_CN":"行动类别: 0:deny, 2:cname,3:a"}
  ActionType *int `json:"actionType,omitempty" xml:"actionType,omitempty" require:"true"`
  // {"en":"only useful if actionType=3(a) or actionType=2(cname), default is null", "zh_CN":"只有在actionType=3(a)或actionType=2(cname)时有用,默认为null"}
  Answer []*UpdateClbDomainByIDRequestEtcPoliciesActionsAnswer `json:"answer,omitempty" xml:"answer,omitempty" require:"true" type:"Repeated"`
}

func (s UpdateClbDomainByIDRequestEtcPoliciesActions) String() string {
  return tea.Prettify(s)
}

func (s UpdateClbDomainByIDRequestEtcPoliciesActions) GoString() string {
  return s.String()
}

func (s *UpdateClbDomainByIDRequestEtcPoliciesActions) SetActionType(v int) *UpdateClbDomainByIDRequestEtcPoliciesActions {
  s.ActionType = &v
  return s
}

func (s *UpdateClbDomainByIDRequestEtcPoliciesActions) SetAnswer(v []*UpdateClbDomainByIDRequestEtcPoliciesActionsAnswer) *UpdateClbDomainByIDRequestEtcPoliciesActions {
  s.Answer = v
  return s
}

type UpdateClbDomainByIDRequestEtcPoliciesActionsAnswer struct     {
  // {"en":"when actionType=3,this is IP, if actionType=2,that is domain name,eg: 1.2.3.4 or test.com", "zh_CN":"当actionType=3时，值为IP地址，当actionType=2时，值为域名地址"}
  Data *string `json:"data,omitempty" xml:"data,omitempty" require:"true"`
  // {"en":"TTL", "zh_CN":"TTL"}
  Ttl *int `json:"ttl,omitempty" xml:"ttl,omitempty" require:"true"`
}

func (s UpdateClbDomainByIDRequestEtcPoliciesActionsAnswer) String() string {
  return tea.Prettify(s)
}

func (s UpdateClbDomainByIDRequestEtcPoliciesActionsAnswer) GoString() string {
  return s.String()
}

func (s *UpdateClbDomainByIDRequestEtcPoliciesActionsAnswer) SetData(v string) *UpdateClbDomainByIDRequestEtcPoliciesActionsAnswer {
  s.Data = &v
  return s
}

func (s *UpdateClbDomainByIDRequestEtcPoliciesActionsAnswer) SetTtl(v int) *UpdateClbDomainByIDRequestEtcPoliciesActionsAnswer {
  s.Ttl = &v
  return s
}

type UpdateClbDomainByIDResponse struct {
  // {"en":"return data", "zh_CN":"返回值"}
  Data *UpdateClbDomainByIDResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
  // {"en":"code,success is 0", "zh_CN":"状态码,成功为0"}
  Code *int `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"error message or Success", "zh_CN":"错误信息或Success"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s UpdateClbDomainByIDResponse) String() string {
  return tea.Prettify(s)
}

func (s UpdateClbDomainByIDResponse) GoString() string {
  return s.String()
}

func (s *UpdateClbDomainByIDResponse) SetData(v *UpdateClbDomainByIDResponseData) *UpdateClbDomainByIDResponse {
  s.Data = v
  return s
}

func (s *UpdateClbDomainByIDResponse) SetCode(v int) *UpdateClbDomainByIDResponse {
  s.Code = &v
  return s
}

func (s *UpdateClbDomainByIDResponse) SetMessage(v string) *UpdateClbDomainByIDResponse {
  s.Message = &v
  return s
}

type UpdateClbDomainByIDResponseData struct {
  // {'en':'customer info','zh_CN':'客户信息'}
  Customer *UpdateClbDomainByIDResponseDataCustomer `json:"customer,omitempty" xml:"customer,omitempty" require:"true" type:"Struct"`
  // {"en":"The zone id", "zh_CN":"zone的ID"}
  ZoneId *int `json:"zoneId,omitempty" xml:"zoneId,omitempty" require:"true"`
  // {"en":"The domain id", "zh_CN":"clbDomain的ID"}
  DomainId *int `json:"domainId,omitempty" xml:"domainId,omitempty" require:"true"`
  // {"en":"zone name", "zh_CN":"域名"}
  ZoneName *string `json:"zoneName,omitempty" xml:"zoneName,omitempty" require:"true"`
  // {"en":"domain name", "zh_CN":"clbDomain名称"}
  DomainName *string `json:"domainName,omitempty" xml:"domainName,omitempty" require:"true"`
  // {"en":"The unicode name of zone", "zh_CN":"Zone的Unicode名称"}
  ZoneUnicode *string `json:"zoneUnicode,omitempty" xml:"zoneUnicode,omitempty" require:"true"`
  // {"en":"The unicode name of domain", "zh_CN":"clbDomain的Unicode名称"}
  DomainUnicode *string `json:"domainUnicode,omitempty" xml:"domainUnicode,omitempty" require:"true"`
  // {"en":"STATUS CODE ( 0: NEW OR MODIFY ; 1:PENDING TO STAGE; 2: IN STAGE; 3: PENDING TO PRODUCTION;4: IN PRODUCTION; -2: ERROR TO STAGE; -4: ERROR TO PRODUCTION )", "zh_CN":"状态码（0: NEW; 1:PENDING TO STAGE; 2: IN STAGE; 3: PENDING TO PRODUCTION;4: IN PRODUCTION; -2: ERROR TO STAGE; -4: ERROR TO PRODUCTION）"}
  ZoneStatusCode *int `json:"zoneStatusCode,omitempty" xml:"zoneStatusCode,omitempty" require:"true"`
  // {"en":"status code description", "zh_CN":"状态码说明"}
  ZoneStatusText *string `json:"zoneStatusText,omitempty" xml:"zoneStatusText,omitempty" require:"true"`
  // {"en":"Published time, UTC time(eg: 2012-11-12T02:08:34Z )", "zh_CN":"发布时间，UTC时间(如: 2012-11-12T02:08:34Z )"}
  ZoneDatePublished *string `json:"zoneDatePublished,omitempty" xml:"zoneDatePublished,omitempty" require:"true"`
  // {"en":"create time, UTC time(eg: 2012-11-12T02:08:34Z )", "zh_CN":"创建时间，UTC时间(如: 2012-11-12T02:08:34Z )"}
  DateCreated *string `json:"dateCreated,omitempty" xml:"dateCreated,omitempty" require:"true"`
  // {"en":"modify time, UTC time(eg: 2012-11-12T02:08:34Z )", "zh_CN":"修改时间，UTC时间(如: 2012-11-12T02:08:34Z )"}
  DateModified *string `json:"dateModified,omitempty" xml:"dateModified,omitempty" require:"true"`
  // {"en":"Dynamic Selection Type,0:GSLB = Shortest lantency, 1:RR = Round Robin, 2:Random = Random", "zh_CN":"动态选择类型, 0:GSLB = Shortest lantency, 1:RR = Round Robin, 2:Random = Random"}
  LoadBalancingType *int `json:"loadBalancingType,omitempty" xml:"loadBalancingType,omitempty" require:"true"`
  // {"en":"Dynamic Selection Type Text", "zh_CN":"动态选择类型名称"}
  LoadBalancingTypeText *string `json:"loadBalancingTypeText,omitempty" xml:"loadBalancingTypeText,omitempty" require:"true"`
  // {"en":"Portal CLB Basic Settings > Number of answer", "zh_CN":"Portal CLB 基础设置 > Number of answer"}
  AnswerCount *int `json:"answerCount,omitempty" xml:"answerCount,omitempty" require:"true"`
  // {"en":"TTL", "zh_CN":"TTL"}
  AnswerTtl *int `json:"answerTtl,omitempty" xml:"answerTtl,omitempty" require:"true"`
  // {"en":"health checker, can be null", "zh_CN":"监控检查,可以为null"}
  Healthchecker *UpdateClbDomainByIDResponseDataHealthchecker `json:"healthchecker,omitempty" xml:"healthchecker,omitempty" require:"true" type:"Struct"`
  // {"en":"delay start time(second), only use when healthchecker not null", "zh_CN":"延迟开始时间,在healthchecker不为空时使用"}
  DelayStartTime *int `json:"delayStartTime,omitempty" xml:"delayStartTime,omitempty" require:"true"`
  // {"en":"slow Start Period(second), only use when healthchecker not null", "zh_CN":"慢启动期,在healthchecker不为空时使用"}
  SlowStartPeriod *int `json:"slowStartPeriod,omitempty" xml:"slowStartPeriod,omitempty" require:"true"`
  // {"en":"domain vip list", "zh_CN":"domain的VIP列表"}
  DomainVips []*UpdateClbDomainByIDResponseDataDomainVips `json:"domainVips,omitempty" xml:"domainVips,omitempty" require:"true" type:"Repeated"`
  // {"en":"policy config", "zh_CN":"策略配置"}
  Policies []*UpdateClbDomainByIDResponseDataPolicies `json:"policies,omitempty" xml:"policies,omitempty" require:"true" type:"Repeated"`
  // {"en":"etc policy config, default = null(Dyanamic Selection From All IPs)", "zh_CN":"Etc策略配置, default = null(Dyanamic Selection From All IPs)"}
  EtcPolicies []*UpdateClbDomainByIDResponseDataEtcPolicies `json:"etcPolicies,omitempty" xml:"etcPolicies,omitempty" require:"true" type:"Repeated"`
}

func (s UpdateClbDomainByIDResponseData) String() string {
  return tea.Prettify(s)
}

func (s UpdateClbDomainByIDResponseData) GoString() string {
  return s.String()
}

func (s *UpdateClbDomainByIDResponseData) SetCustomer(v *UpdateClbDomainByIDResponseDataCustomer) *UpdateClbDomainByIDResponseData {
  s.Customer = v
  return s
}

func (s *UpdateClbDomainByIDResponseData) SetZoneId(v int) *UpdateClbDomainByIDResponseData {
  s.ZoneId = &v
  return s
}

func (s *UpdateClbDomainByIDResponseData) SetDomainId(v int) *UpdateClbDomainByIDResponseData {
  s.DomainId = &v
  return s
}

func (s *UpdateClbDomainByIDResponseData) SetZoneName(v string) *UpdateClbDomainByIDResponseData {
  s.ZoneName = &v
  return s
}

func (s *UpdateClbDomainByIDResponseData) SetDomainName(v string) *UpdateClbDomainByIDResponseData {
  s.DomainName = &v
  return s
}

func (s *UpdateClbDomainByIDResponseData) SetZoneUnicode(v string) *UpdateClbDomainByIDResponseData {
  s.ZoneUnicode = &v
  return s
}

func (s *UpdateClbDomainByIDResponseData) SetDomainUnicode(v string) *UpdateClbDomainByIDResponseData {
  s.DomainUnicode = &v
  return s
}

func (s *UpdateClbDomainByIDResponseData) SetZoneStatusCode(v int) *UpdateClbDomainByIDResponseData {
  s.ZoneStatusCode = &v
  return s
}

func (s *UpdateClbDomainByIDResponseData) SetZoneStatusText(v string) *UpdateClbDomainByIDResponseData {
  s.ZoneStatusText = &v
  return s
}

func (s *UpdateClbDomainByIDResponseData) SetZoneDatePublished(v string) *UpdateClbDomainByIDResponseData {
  s.ZoneDatePublished = &v
  return s
}

func (s *UpdateClbDomainByIDResponseData) SetDateCreated(v string) *UpdateClbDomainByIDResponseData {
  s.DateCreated = &v
  return s
}

func (s *UpdateClbDomainByIDResponseData) SetDateModified(v string) *UpdateClbDomainByIDResponseData {
  s.DateModified = &v
  return s
}

func (s *UpdateClbDomainByIDResponseData) SetLoadBalancingType(v int) *UpdateClbDomainByIDResponseData {
  s.LoadBalancingType = &v
  return s
}

func (s *UpdateClbDomainByIDResponseData) SetLoadBalancingTypeText(v string) *UpdateClbDomainByIDResponseData {
  s.LoadBalancingTypeText = &v
  return s
}

func (s *UpdateClbDomainByIDResponseData) SetAnswerCount(v int) *UpdateClbDomainByIDResponseData {
  s.AnswerCount = &v
  return s
}

func (s *UpdateClbDomainByIDResponseData) SetAnswerTtl(v int) *UpdateClbDomainByIDResponseData {
  s.AnswerTtl = &v
  return s
}

func (s *UpdateClbDomainByIDResponseData) SetHealthchecker(v *UpdateClbDomainByIDResponseDataHealthchecker) *UpdateClbDomainByIDResponseData {
  s.Healthchecker = v
  return s
}

func (s *UpdateClbDomainByIDResponseData) SetDelayStartTime(v int) *UpdateClbDomainByIDResponseData {
  s.DelayStartTime = &v
  return s
}

func (s *UpdateClbDomainByIDResponseData) SetSlowStartPeriod(v int) *UpdateClbDomainByIDResponseData {
  s.SlowStartPeriod = &v
  return s
}

func (s *UpdateClbDomainByIDResponseData) SetDomainVips(v []*UpdateClbDomainByIDResponseDataDomainVips) *UpdateClbDomainByIDResponseData {
  s.DomainVips = v
  return s
}

func (s *UpdateClbDomainByIDResponseData) SetPolicies(v []*UpdateClbDomainByIDResponseDataPolicies) *UpdateClbDomainByIDResponseData {
  s.Policies = v
  return s
}

func (s *UpdateClbDomainByIDResponseData) SetEtcPolicies(v []*UpdateClbDomainByIDResponseDataEtcPolicies) *UpdateClbDomainByIDResponseData {
  s.EtcPolicies = v
  return s
}

type UpdateClbDomainByIDResponseDataCustomer struct {
  // {'en':'customer name','zh_CN':'客户名称'}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
}

func (s UpdateClbDomainByIDResponseDataCustomer) String() string {
  return tea.Prettify(s)
}

func (s UpdateClbDomainByIDResponseDataCustomer) GoString() string {
  return s.String()
}

func (s *UpdateClbDomainByIDResponseDataCustomer) SetName(v string) *UpdateClbDomainByIDResponseDataCustomer {
  s.Name = &v
  return s
}

type UpdateClbDomainByIDResponseDataHealthchecker struct {
  // {"en":"health checker id", "zh_CN":"监控检查ID"}
  Id *int `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {"en":"health checker name", "zh_CN":"监控检查名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
}

func (s UpdateClbDomainByIDResponseDataHealthchecker) String() string {
  return tea.Prettify(s)
}

func (s UpdateClbDomainByIDResponseDataHealthchecker) GoString() string {
  return s.String()
}

func (s *UpdateClbDomainByIDResponseDataHealthchecker) SetId(v int) *UpdateClbDomainByIDResponseDataHealthchecker {
  s.Id = &v
  return s
}

func (s *UpdateClbDomainByIDResponseDataHealthchecker) SetName(v string) *UpdateClbDomainByIDResponseDataHealthchecker {
  s.Name = &v
  return s
}

type UpdateClbDomainByIDResponseDataDomainVips struct     {
  // {"en":"Servergroup's ID", "zh_CN":"Servergroup的ID"}
  DomainVipId *int `json:"domainVipId,omitempty" xml:"domainVipId,omitempty" require:"true"`
  // {"en":"Servergroup's ID", "zh_CN":"Servergroup的ID"}
  ServergroupId *int `json:"servergroupId,omitempty" xml:"servergroupId,omitempty" require:"true"`
  // {"en":"server group's name", "zh_CN":"server group的名称"}
  ServergroupName *string `json:"servergroupName,omitempty" xml:"servergroupName,omitempty" require:"true"`
  // {"en":"Server's ID", "zh_CN":"Server的ID"}
  ServerId *int `json:"serverId,omitempty" xml:"serverId,omitempty" require:"true"`
  // {"en":"server name", "zh_CN":"服务名"}
  ServerName *string `json:"serverName,omitempty" xml:"serverName,omitempty" require:"true"`
  // {"en":"server group region id", "zh_CN":"server group的地理ID"}
  Region *string `json:"region,omitempty" xml:"region,omitempty" require:"true"`
  // {"en":"server group region name", "zh_CN":"server group的地理名称"}
  RegionName *string `json:"regionName,omitempty" xml:"regionName,omitempty" require:"true"`
  // {"en":"IP", "zh_CN":"ip地址"}
  Ip *string `json:"ip,omitempty" xml:"ip,omitempty" require:"true"`
  // {"en":"enable", "zh_CN":"是否可用"}
  IsEnabled *bool `json:"isEnabled,omitempty" xml:"isEnabled,omitempty" require:"true"`
  // {"en":"health checkers", "zh_CN":"监控检查列表"}
  Healthcheckers []*UpdateClbDomainByIDResponseDataDomainVipsHealthcheckers `json:"healthcheckers,omitempty" xml:"healthcheckers,omitempty" require:"true" type:"Repeated"`
}

func (s UpdateClbDomainByIDResponseDataDomainVips) String() string {
  return tea.Prettify(s)
}

func (s UpdateClbDomainByIDResponseDataDomainVips) GoString() string {
  return s.String()
}

func (s *UpdateClbDomainByIDResponseDataDomainVips) SetDomainVipId(v int) *UpdateClbDomainByIDResponseDataDomainVips {
  s.DomainVipId = &v
  return s
}

func (s *UpdateClbDomainByIDResponseDataDomainVips) SetServergroupId(v int) *UpdateClbDomainByIDResponseDataDomainVips {
  s.ServergroupId = &v
  return s
}

func (s *UpdateClbDomainByIDResponseDataDomainVips) SetServergroupName(v string) *UpdateClbDomainByIDResponseDataDomainVips {
  s.ServergroupName = &v
  return s
}

func (s *UpdateClbDomainByIDResponseDataDomainVips) SetServerId(v int) *UpdateClbDomainByIDResponseDataDomainVips {
  s.ServerId = &v
  return s
}

func (s *UpdateClbDomainByIDResponseDataDomainVips) SetServerName(v string) *UpdateClbDomainByIDResponseDataDomainVips {
  s.ServerName = &v
  return s
}

func (s *UpdateClbDomainByIDResponseDataDomainVips) SetRegion(v string) *UpdateClbDomainByIDResponseDataDomainVips {
  s.Region = &v
  return s
}

func (s *UpdateClbDomainByIDResponseDataDomainVips) SetRegionName(v string) *UpdateClbDomainByIDResponseDataDomainVips {
  s.RegionName = &v
  return s
}

func (s *UpdateClbDomainByIDResponseDataDomainVips) SetIp(v string) *UpdateClbDomainByIDResponseDataDomainVips {
  s.Ip = &v
  return s
}

func (s *UpdateClbDomainByIDResponseDataDomainVips) SetIsEnabled(v bool) *UpdateClbDomainByIDResponseDataDomainVips {
  s.IsEnabled = &v
  return s
}

func (s *UpdateClbDomainByIDResponseDataDomainVips) SetHealthcheckers(v []*UpdateClbDomainByIDResponseDataDomainVipsHealthcheckers) *UpdateClbDomainByIDResponseDataDomainVips {
  s.Healthcheckers = v
  return s
}

type UpdateClbDomainByIDResponseDataDomainVipsHealthcheckers struct     {
  // {"en":"health checker id", "zh_CN":"监控检查ID"}
  Id *int `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {"en":"health checker name", "zh_CN":"监控检查名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
}

func (s UpdateClbDomainByIDResponseDataDomainVipsHealthcheckers) String() string {
  return tea.Prettify(s)
}

func (s UpdateClbDomainByIDResponseDataDomainVipsHealthcheckers) GoString() string {
  return s.String()
}

func (s *UpdateClbDomainByIDResponseDataDomainVipsHealthcheckers) SetId(v int) *UpdateClbDomainByIDResponseDataDomainVipsHealthcheckers {
  s.Id = &v
  return s
}

func (s *UpdateClbDomainByIDResponseDataDomainVipsHealthcheckers) SetName(v string) *UpdateClbDomainByIDResponseDataDomainVipsHealthcheckers {
  s.Name = &v
  return s
}

type UpdateClbDomainByIDResponseDataPolicies struct     {
  // {"en":"queue number", "zh_CN":"排序号"}
  Sequence *int `json:"sequence,omitempty" xml:"sequence,omitempty" require:"true"`
  // {"en":"policy name", "zh_CN":"策略名称"}
  PolicyName *string `json:"policyName,omitempty" xml:"policyName,omitempty" require:"true"`
  // {"en":"enable", "zh_CN":"是否可用"}
  IsEnabled *bool `json:"isEnabled,omitempty" xml:"isEnabled,omitempty" require:"true"`
  // {"en":"condition", "zh_CN":"条件"}
  Condition *UpdateClbDomainByIDResponseDataPoliciesCondition `json:"condition,omitempty" xml:"condition,omitempty" require:"true" type:"Struct"`
  // {"en":"action", "zh_CN":"行动"}
  Actions []*UpdateClbDomainByIDResponseDataPoliciesActions `json:"actions,omitempty" xml:"actions,omitempty" require:"true" type:"Repeated"`
}

func (s UpdateClbDomainByIDResponseDataPolicies) String() string {
  return tea.Prettify(s)
}

func (s UpdateClbDomainByIDResponseDataPolicies) GoString() string {
  return s.String()
}

func (s *UpdateClbDomainByIDResponseDataPolicies) SetSequence(v int) *UpdateClbDomainByIDResponseDataPolicies {
  s.Sequence = &v
  return s
}

func (s *UpdateClbDomainByIDResponseDataPolicies) SetPolicyName(v string) *UpdateClbDomainByIDResponseDataPolicies {
  s.PolicyName = &v
  return s
}

func (s *UpdateClbDomainByIDResponseDataPolicies) SetIsEnabled(v bool) *UpdateClbDomainByIDResponseDataPolicies {
  s.IsEnabled = &v
  return s
}

func (s *UpdateClbDomainByIDResponseDataPolicies) SetCondition(v *UpdateClbDomainByIDResponseDataPoliciesCondition) *UpdateClbDomainByIDResponseDataPolicies {
  s.Condition = v
  return s
}

func (s *UpdateClbDomainByIDResponseDataPolicies) SetActions(v []*UpdateClbDomainByIDResponseDataPoliciesActions) *UpdateClbDomainByIDResponseDataPolicies {
  s.Actions = v
  return s
}

type UpdateClbDomainByIDResponseDataPoliciesCondition struct {
  // {"en":"The condition id", "zh_CN":"条件ID"}
  ConditionId *int `json:"conditionId,omitempty" xml:"conditionId,omitempty" require:"true"`
  // {"en":"condition type, 0:IP RANGE, 1: GEO DATA, 2 Ratio(%)", "zh_CN":"条件类型, 0:IP RANGE, 1: GEO DATA, 2 Ratio(%)"}
  ConditionType *int `json:"conditionType,omitempty" xml:"conditionType,omitempty" require:"true"`
  // {"en":"IP range,comma separated,only useful if condition type is 0(ip range), eg: 1.2.3.4,1.5.6.0/24", "zh_CN":"IP范围,逗号分隔,只有在contidion type为0(ip range)时有用, 例:1.2.3.4,1.5.6.0/24"}
  IpRange *string `json:"ipRange,omitempty" xml:"ipRange,omitempty" require:"true"`
  // {"en":"invert, only useful if condition type is 0(ip range) or 1(geo data) ", "zh_CN":"反转, 只有在contidion type为0(ip range)或1(geo data)时有用"}
  Invert *bool `json:"invert,omitempty" xml:"invert,omitempty" require:"true"`
  // {"en":"region id list,only useful if condition type is 1(geo data), get by regions api", "zh_CN":"地理位置的ID列表，只有在contidion type为1(geo data)时有用，从regions接口获取"}
  Region []*string `json:"region,omitempty" xml:"region,omitempty" require:"true" type:"Repeated"`
  // {"en":"config view id list,only useful if condition type is 1(geo data), get by cvs api", "zh_CN":"配置视图ID列表，只有在contidion type为1(geo data)时有用，从cvs接口获取"}
  Cv []*string `json:"cv,omitempty" xml:"cv,omitempty" require:"true" type:"Repeated"`
  // {"en":"isp id list,only useful if condition type is 1(geo data), get by isps api", "zh_CN":"供应商ID列表，只有在contidion type为1(geo data)时有用，从isps接口获取"}
  Isp []*string `json:"isp,omitempty" xml:"isp,omitempty" require:"true" type:"Repeated"`
  // {"en":"as list, only useful if condition type is 1(geo data)", "zh_CN":"as列表，只有在contidion type为1(geo data)时有用"}
  Asn []*int `json:"asn,omitempty" xml:"asn,omitempty" require:"true" type:"Repeated"`
  // {"en":"percent,Only use conditionType = 2(Ratio(%)), eg: 70", "zh_CN":"比例，只有在conditionType=2 (Ratio%)时有用，如: 70"}
  Percent *int `json:"percent,omitempty" xml:"percent,omitempty" require:"true"`
}

func (s UpdateClbDomainByIDResponseDataPoliciesCondition) String() string {
  return tea.Prettify(s)
}

func (s UpdateClbDomainByIDResponseDataPoliciesCondition) GoString() string {
  return s.String()
}

func (s *UpdateClbDomainByIDResponseDataPoliciesCondition) SetConditionId(v int) *UpdateClbDomainByIDResponseDataPoliciesCondition {
  s.ConditionId = &v
  return s
}

func (s *UpdateClbDomainByIDResponseDataPoliciesCondition) SetConditionType(v int) *UpdateClbDomainByIDResponseDataPoliciesCondition {
  s.ConditionType = &v
  return s
}

func (s *UpdateClbDomainByIDResponseDataPoliciesCondition) SetIpRange(v string) *UpdateClbDomainByIDResponseDataPoliciesCondition {
  s.IpRange = &v
  return s
}

func (s *UpdateClbDomainByIDResponseDataPoliciesCondition) SetInvert(v bool) *UpdateClbDomainByIDResponseDataPoliciesCondition {
  s.Invert = &v
  return s
}

func (s *UpdateClbDomainByIDResponseDataPoliciesCondition) SetRegion(v []*string) *UpdateClbDomainByIDResponseDataPoliciesCondition {
  s.Region = v
  return s
}

func (s *UpdateClbDomainByIDResponseDataPoliciesCondition) SetCv(v []*string) *UpdateClbDomainByIDResponseDataPoliciesCondition {
  s.Cv = v
  return s
}

func (s *UpdateClbDomainByIDResponseDataPoliciesCondition) SetIsp(v []*string) *UpdateClbDomainByIDResponseDataPoliciesCondition {
  s.Isp = v
  return s
}

func (s *UpdateClbDomainByIDResponseDataPoliciesCondition) SetAsn(v []*int) *UpdateClbDomainByIDResponseDataPoliciesCondition {
  s.Asn = v
  return s
}

func (s *UpdateClbDomainByIDResponseDataPoliciesCondition) SetPercent(v int) *UpdateClbDomainByIDResponseDataPoliciesCondition {
  s.Percent = &v
  return s
}

type UpdateClbDomainByIDResponseDataPoliciesActions struct     {
  // {"en":"queue number", "zh_CN":"排序号"}
  Sequence *int `json:"sequence,omitempty" xml:"sequence,omitempty" require:"true"`
  // {"en":"The action id", "zh_CN":"行动ID"}
  ActionId *int `json:"actionId,omitempty" xml:"actionId,omitempty" require:"true"`
  // {"en":"action type, 0:deny, 1:vip, 2:cname,3:a", "zh_CN":"行动类别: 0:deny, 1:vip, 2:cname,3:a"}
  ActionType *int `json:"actionType,omitempty" xml:"actionType,omitempty" require:"true"`
  // {"en":"vip ID, only useful if actionType=1(vip), get in domainVips serverId", "zh_CN":"VIP ID,只有在action type为1(vip)时有用，是属性domainVips中的serverId值"}
  Vips []*int `json:"vips,omitempty" xml:"vips,omitempty" require:"true" type:"Repeated"`
  // {"en":"only useful if actionType=3(a) or actionType=2(cname)", "zh_CN":"只有在actionType=3(a)或actionType=2(cname)时有用"}
  Answer []*UpdateClbDomainByIDResponseDataPoliciesActionsAnswer `json:"answer,omitempty" xml:"answer,omitempty" require:"true" type:"Repeated"`
}

func (s UpdateClbDomainByIDResponseDataPoliciesActions) String() string {
  return tea.Prettify(s)
}

func (s UpdateClbDomainByIDResponseDataPoliciesActions) GoString() string {
  return s.String()
}

func (s *UpdateClbDomainByIDResponseDataPoliciesActions) SetSequence(v int) *UpdateClbDomainByIDResponseDataPoliciesActions {
  s.Sequence = &v
  return s
}

func (s *UpdateClbDomainByIDResponseDataPoliciesActions) SetActionId(v int) *UpdateClbDomainByIDResponseDataPoliciesActions {
  s.ActionId = &v
  return s
}

func (s *UpdateClbDomainByIDResponseDataPoliciesActions) SetActionType(v int) *UpdateClbDomainByIDResponseDataPoliciesActions {
  s.ActionType = &v
  return s
}

func (s *UpdateClbDomainByIDResponseDataPoliciesActions) SetVips(v []*int) *UpdateClbDomainByIDResponseDataPoliciesActions {
  s.Vips = v
  return s
}

func (s *UpdateClbDomainByIDResponseDataPoliciesActions) SetAnswer(v []*UpdateClbDomainByIDResponseDataPoliciesActionsAnswer) *UpdateClbDomainByIDResponseDataPoliciesActions {
  s.Answer = v
  return s
}

type UpdateClbDomainByIDResponseDataPoliciesActionsAnswer struct     {
  // {"en":"actionType=3, A or actionType=2, CNAME", "zh_CN":"actionType=3,则值为A,actionType=2,则值为CNAME"}
  Rtype *string `json:"rtype,omitempty" xml:"rtype,omitempty" require:"true"`
  // {"en":"when actionType=3,this is IP, if actionType=2,that is domain name,eg: 1.2.3.4 or test.com", "zh_CN":"当actionType=3时，值为IP地址，当actionType=2时，值为域名地址"}
  Data *string `json:"data,omitempty" xml:"data,omitempty" require:"true"`
  // {"en":"TTL", "zh_CN":"TTL"}
  Ttl *int `json:"ttl,omitempty" xml:"ttl,omitempty" require:"true"`
}

func (s UpdateClbDomainByIDResponseDataPoliciesActionsAnswer) String() string {
  return tea.Prettify(s)
}

func (s UpdateClbDomainByIDResponseDataPoliciesActionsAnswer) GoString() string {
  return s.String()
}

func (s *UpdateClbDomainByIDResponseDataPoliciesActionsAnswer) SetRtype(v string) *UpdateClbDomainByIDResponseDataPoliciesActionsAnswer {
  s.Rtype = &v
  return s
}

func (s *UpdateClbDomainByIDResponseDataPoliciesActionsAnswer) SetData(v string) *UpdateClbDomainByIDResponseDataPoliciesActionsAnswer {
  s.Data = &v
  return s
}

func (s *UpdateClbDomainByIDResponseDataPoliciesActionsAnswer) SetTtl(v int) *UpdateClbDomainByIDResponseDataPoliciesActionsAnswer {
  s.Ttl = &v
  return s
}

type UpdateClbDomainByIDResponseDataEtcPolicies struct     {
  // {"en":"policy name(always)", "zh_CN":"策略名称(always)"}
  PolicyName *string `json:"policyName,omitempty" xml:"policyName,omitempty" require:"true"`
  // {"en":"action", "zh_CN":"行动"}
  Actions []*UpdateClbDomainByIDResponseDataEtcPoliciesActions `json:"actions,omitempty" xml:"actions,omitempty" require:"true" type:"Repeated"`
}

func (s UpdateClbDomainByIDResponseDataEtcPolicies) String() string {
  return tea.Prettify(s)
}

func (s UpdateClbDomainByIDResponseDataEtcPolicies) GoString() string {
  return s.String()
}

func (s *UpdateClbDomainByIDResponseDataEtcPolicies) SetPolicyName(v string) *UpdateClbDomainByIDResponseDataEtcPolicies {
  s.PolicyName = &v
  return s
}

func (s *UpdateClbDomainByIDResponseDataEtcPolicies) SetActions(v []*UpdateClbDomainByIDResponseDataEtcPoliciesActions) *UpdateClbDomainByIDResponseDataEtcPolicies {
  s.Actions = v
  return s
}

type UpdateClbDomainByIDResponseDataEtcPoliciesActions struct     {
  // {"en":"action type, 0:deny, 2:cname,3:a", "zh_CN":"行动类别: 0:deny, 2:cname,3:a"}
  ActionType *int `json:"actionType,omitempty" xml:"actionType,omitempty" require:"true"`
  // {"en":"only useful if actionType=3(a) or actionType=2(cname), default is null", "zh_CN":"只有在actionType=3(a)或actionType=2(cname)时有用,默认为null"}
  Answer []*UpdateClbDomainByIDResponseDataEtcPoliciesActionsAnswer `json:"answer,omitempty" xml:"answer,omitempty" require:"true" type:"Repeated"`
}

func (s UpdateClbDomainByIDResponseDataEtcPoliciesActions) String() string {
  return tea.Prettify(s)
}

func (s UpdateClbDomainByIDResponseDataEtcPoliciesActions) GoString() string {
  return s.String()
}

func (s *UpdateClbDomainByIDResponseDataEtcPoliciesActions) SetActionType(v int) *UpdateClbDomainByIDResponseDataEtcPoliciesActions {
  s.ActionType = &v
  return s
}

func (s *UpdateClbDomainByIDResponseDataEtcPoliciesActions) SetAnswer(v []*UpdateClbDomainByIDResponseDataEtcPoliciesActionsAnswer) *UpdateClbDomainByIDResponseDataEtcPoliciesActions {
  s.Answer = v
  return s
}

type UpdateClbDomainByIDResponseDataEtcPoliciesActionsAnswer struct     {
  // {"en":"actionType=3, A or actionType=2, CNAME", "zh_CN":"actionType=3,则值为A,actionType=2,则值为CNAME"}
  Rtype *string `json:"rtype,omitempty" xml:"rtype,omitempty" require:"true"`
  // {"en":"when actionType=3,this is IP, if actionType=2,that is domain name,eg: 1.2.3.4 or test.com", "zh_CN":"当actionType=3时，值为IP地址，当actionType=2时，值为域名地址"}
  Data *string `json:"data,omitempty" xml:"data,omitempty" require:"true"`
  // {"en":"TTL", "zh_CN":"TTL"}
  Ttl *int `json:"ttl,omitempty" xml:"ttl,omitempty" require:"true"`
}

func (s UpdateClbDomainByIDResponseDataEtcPoliciesActionsAnswer) String() string {
  return tea.Prettify(s)
}

func (s UpdateClbDomainByIDResponseDataEtcPoliciesActionsAnswer) GoString() string {
  return s.String()
}

func (s *UpdateClbDomainByIDResponseDataEtcPoliciesActionsAnswer) SetRtype(v string) *UpdateClbDomainByIDResponseDataEtcPoliciesActionsAnswer {
  s.Rtype = &v
  return s
}

func (s *UpdateClbDomainByIDResponseDataEtcPoliciesActionsAnswer) SetData(v string) *UpdateClbDomainByIDResponseDataEtcPoliciesActionsAnswer {
  s.Data = &v
  return s
}

func (s *UpdateClbDomainByIDResponseDataEtcPoliciesActionsAnswer) SetTtl(v int) *UpdateClbDomainByIDResponseDataEtcPoliciesActionsAnswer {
  s.Ttl = &v
  return s
}

type UpdateClbDomainByIDPaths struct {
  // {"en":"CLB Domain's ID", "zh_CN":"CLB Domain的ID"}
  DomainId *int `json:"domainId,omitempty" xml:"domainId,omitempty" require:"true"`
}

func (s UpdateClbDomainByIDPaths) String() string {
  return tea.Prettify(s)
}

func (s UpdateClbDomainByIDPaths) GoString() string {
  return s.String()
}

func (s *UpdateClbDomainByIDPaths) SetDomainId(v int) *UpdateClbDomainByIDPaths {
  s.DomainId = &v
  return s
}

type UpdateClbDomainByIDParameters struct {
}

func (s UpdateClbDomainByIDParameters) String() string {
  return tea.Prettify(s)
}

func (s UpdateClbDomainByIDParameters) GoString() string {
  return s.String()
}

type UpdateClbDomainByIDRequestHeader struct {
}

func (s UpdateClbDomainByIDRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s UpdateClbDomainByIDRequestHeader) GoString() string {
  return s.String()
}

type UpdateClbDomainByIDResponseHeader struct {
}

func (s UpdateClbDomainByIDResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s UpdateClbDomainByIDResponseHeader) GoString() string {
  return s.String()
}




type GetClbDomainListRequest struct {
}

func (s GetClbDomainListRequest) String() string {
  return tea.Prettify(s)
}

func (s GetClbDomainListRequest) GoString() string {
  return s.String()
}

type GetClbDomainListResponse struct {
  // {"en":"return data", "zh_CN":"返回值"}
  Data []*GetClbDomainListResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
  // {"en":"code,success is 0", "zh_CN":"状态码,成功为0"}
  Code *int `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"error message or Success", "zh_CN":"错误信息或Success"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s GetClbDomainListResponse) String() string {
  return tea.Prettify(s)
}

func (s GetClbDomainListResponse) GoString() string {
  return s.String()
}

func (s *GetClbDomainListResponse) SetData(v []*GetClbDomainListResponseData) *GetClbDomainListResponse {
  s.Data = v
  return s
}

func (s *GetClbDomainListResponse) SetCode(v int) *GetClbDomainListResponse {
  s.Code = &v
  return s
}

func (s *GetClbDomainListResponse) SetMessage(v string) *GetClbDomainListResponse {
  s.Message = &v
  return s
}

type GetClbDomainListResponseData struct     {
  // {"en":"Pagination info", "zh_CN":"分页信息"}
  PageInfo *GetClbDomainListResponseDataPageInfo `json:"pageInfo,omitempty" xml:"pageInfo,omitempty" require:"true" type:"Struct"`
  // {"en":"return data", "zh_CN":"返回值"}
  Results []*GetClbDomainListResponseDataResults `json:"results,omitempty" xml:"results,omitempty" require:"true" type:"Repeated"`
}

func (s GetClbDomainListResponseData) String() string {
  return tea.Prettify(s)
}

func (s GetClbDomainListResponseData) GoString() string {
  return s.String()
}

func (s *GetClbDomainListResponseData) SetPageInfo(v *GetClbDomainListResponseDataPageInfo) *GetClbDomainListResponseData {
  s.PageInfo = v
  return s
}

func (s *GetClbDomainListResponseData) SetResults(v []*GetClbDomainListResponseDataResults) *GetClbDomainListResponseData {
  s.Results = v
  return s
}

type GetClbDomainListResponseDataPageInfo struct {
  // {"en":"total pages", "zh_CN":"总页数"}
  Pages *int `json:"pages,omitempty" xml:"pages,omitempty" require:"true"`
  // {"en":"total count", "zh_CN":"总数量"}
  Count *int `json:"count,omitempty" xml:"count,omitempty" require:"true"`
  // {"en":"num of per page", "zh_CN":"每页数量"}
  PageSize *int `json:"pageSize,omitempty" xml:"pageSize,omitempty" require:"true"`
  // {"en":"current page", "zh_CN":"当前页"}
  Page *int `json:"page,omitempty" xml:"page,omitempty" require:"true"`
}

func (s GetClbDomainListResponseDataPageInfo) String() string {
  return tea.Prettify(s)
}

func (s GetClbDomainListResponseDataPageInfo) GoString() string {
  return s.String()
}

func (s *GetClbDomainListResponseDataPageInfo) SetPages(v int) *GetClbDomainListResponseDataPageInfo {
  s.Pages = &v
  return s
}

func (s *GetClbDomainListResponseDataPageInfo) SetCount(v int) *GetClbDomainListResponseDataPageInfo {
  s.Count = &v
  return s
}

func (s *GetClbDomainListResponseDataPageInfo) SetPageSize(v int) *GetClbDomainListResponseDataPageInfo {
  s.PageSize = &v
  return s
}

func (s *GetClbDomainListResponseDataPageInfo) SetPage(v int) *GetClbDomainListResponseDataPageInfo {
  s.Page = &v
  return s
}

type GetClbDomainListResponseDataResults struct     {
  // {"en":"The zone id", "zh_CN":"zone的ID"}
  ZoneId *int `json:"zoneId,omitempty" xml:"zoneId,omitempty" require:"true"`
  // {"en":"The domain id", "zh_CN":"clbDomain的ID"}
  DomainId *int `json:"domainId,omitempty" xml:"domainId,omitempty" require:"true"`
  // {"en":"zone name", "zh_CN":"域名"}
  ZoneName *string `json:"zoneName,omitempty" xml:"zoneName,omitempty" require:"true"`
  // {"en":"domain name", "zh_CN":"clbDomain名称"}
  DomainName *string `json:"domainName,omitempty" xml:"domainName,omitempty" require:"true"`
  // {"en":"STATUS CODE ( 0: NEW OR MODIFY ; 1:PENDING TO STAGE; 2: IN STAGE; 3: PENDING TO PRODUCTION;4: IN PRODUCTION; -2: ERROR TO STAGE; -4: ERROR TO PRODUCTION )", "zh_CN":"状态码（0: NEW; 1:PENDING TO STAGE; 2: IN STAGE; 3: PENDING TO PRODUCTION;4: IN PRODUCTION; -2: ERROR TO STAGE; -4: ERROR TO PRODUCTION）"}
  ZoneStatusCode *int `json:"zoneStatusCode,omitempty" xml:"zoneStatusCode,omitempty" require:"true"`
  // {"en":"status code description", "zh_CN":"状态码说明"}
  ZoneStatusText *string `json:"zoneStatusText,omitempty" xml:"zoneStatusText,omitempty" require:"true"`
  // {"en":"total vip count", "zh_CN":"总vip数量"}
  TotalVipCount *int `json:"totalVipCount,omitempty" xml:"totalVipCount,omitempty" require:"true"`
  // {"en":"failed vip count", "zh_CN":"失败vip数量"}
  FailedVipCount *int `json:"failedVipCount,omitempty" xml:"failedVipCount,omitempty" require:"true"`
  // {"en":"Published time, UTC time(eg: 2012-11-12T02:08:34Z )", "zh_CN":"发布时间，UTC时间(如: 2012-11-12T02:08:34Z )"}
  ZoneDatePublished *string `json:"zoneDatePublished,omitempty" xml:"zoneDatePublished,omitempty" require:"true"`
  // {"en":"The unicode name of zone", "zh_CN":"Zone的Unicode名称"}
  ZoneUnicode *string `json:"zoneUnicode,omitempty" xml:"zoneUnicode,omitempty" require:"true"`
  // {"en":"The unicode name of domain", "zh_CN":"clbDomain的Unicode名称"}
  DomainUnicode *string `json:"domainUnicode,omitempty" xml:"domainUnicode,omitempty" require:"true"`
  // {"en":"vip list", "zh_CN":"vip列表"}
  Vips []*GetClbDomainListResponseDataResultsVips `json:"vips,omitempty" xml:"vips,omitempty" require:"true" type:"Repeated"`
}

func (s GetClbDomainListResponseDataResults) String() string {
  return tea.Prettify(s)
}

func (s GetClbDomainListResponseDataResults) GoString() string {
  return s.String()
}

func (s *GetClbDomainListResponseDataResults) SetZoneId(v int) *GetClbDomainListResponseDataResults {
  s.ZoneId = &v
  return s
}

func (s *GetClbDomainListResponseDataResults) SetDomainId(v int) *GetClbDomainListResponseDataResults {
  s.DomainId = &v
  return s
}

func (s *GetClbDomainListResponseDataResults) SetZoneName(v string) *GetClbDomainListResponseDataResults {
  s.ZoneName = &v
  return s
}

func (s *GetClbDomainListResponseDataResults) SetDomainName(v string) *GetClbDomainListResponseDataResults {
  s.DomainName = &v
  return s
}

func (s *GetClbDomainListResponseDataResults) SetZoneStatusCode(v int) *GetClbDomainListResponseDataResults {
  s.ZoneStatusCode = &v
  return s
}

func (s *GetClbDomainListResponseDataResults) SetZoneStatusText(v string) *GetClbDomainListResponseDataResults {
  s.ZoneStatusText = &v
  return s
}

func (s *GetClbDomainListResponseDataResults) SetTotalVipCount(v int) *GetClbDomainListResponseDataResults {
  s.TotalVipCount = &v
  return s
}

func (s *GetClbDomainListResponseDataResults) SetFailedVipCount(v int) *GetClbDomainListResponseDataResults {
  s.FailedVipCount = &v
  return s
}

func (s *GetClbDomainListResponseDataResults) SetZoneDatePublished(v string) *GetClbDomainListResponseDataResults {
  s.ZoneDatePublished = &v
  return s
}

func (s *GetClbDomainListResponseDataResults) SetZoneUnicode(v string) *GetClbDomainListResponseDataResults {
  s.ZoneUnicode = &v
  return s
}

func (s *GetClbDomainListResponseDataResults) SetDomainUnicode(v string) *GetClbDomainListResponseDataResults {
  s.DomainUnicode = &v
  return s
}

func (s *GetClbDomainListResponseDataResults) SetVips(v []*GetClbDomainListResponseDataResultsVips) *GetClbDomainListResponseDataResults {
  s.Vips = v
  return s
}

type GetClbDomainListResponseDataResultsVips struct     {
  // {"en":"IP", "zh_CN":"ip地址"}
  Ip *string `json:"ip,omitempty" xml:"ip,omitempty" require:"true"`
  // {"en":"enable", "zh_CN":"是否可用"}
  IsEnabled *bool `json:"isEnabled,omitempty" xml:"isEnabled,omitempty" require:"true"`
  // {"en":"health checkers", "zh_CN":"监控检查列表"}
  Healthcheckers []*GetClbDomainListResponseDataResultsVipsHealthcheckers `json:"healthcheckers,omitempty" xml:"healthcheckers,omitempty" require:"true" type:"Repeated"`
}

func (s GetClbDomainListResponseDataResultsVips) String() string {
  return tea.Prettify(s)
}

func (s GetClbDomainListResponseDataResultsVips) GoString() string {
  return s.String()
}

func (s *GetClbDomainListResponseDataResultsVips) SetIp(v string) *GetClbDomainListResponseDataResultsVips {
  s.Ip = &v
  return s
}

func (s *GetClbDomainListResponseDataResultsVips) SetIsEnabled(v bool) *GetClbDomainListResponseDataResultsVips {
  s.IsEnabled = &v
  return s
}

func (s *GetClbDomainListResponseDataResultsVips) SetHealthcheckers(v []*GetClbDomainListResponseDataResultsVipsHealthcheckers) *GetClbDomainListResponseDataResultsVips {
  s.Healthcheckers = v
  return s
}

type GetClbDomainListResponseDataResultsVipsHealthcheckers struct     {
  // {"en":"health checker id", "zh_CN":"监控检查ID"}
  Id *int `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {"en":"health checker name", "zh_CN":"监控检查名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
}

func (s GetClbDomainListResponseDataResultsVipsHealthcheckers) String() string {
  return tea.Prettify(s)
}

func (s GetClbDomainListResponseDataResultsVipsHealthcheckers) GoString() string {
  return s.String()
}

func (s *GetClbDomainListResponseDataResultsVipsHealthcheckers) SetId(v int) *GetClbDomainListResponseDataResultsVipsHealthcheckers {
  s.Id = &v
  return s
}

func (s *GetClbDomainListResponseDataResultsVipsHealthcheckers) SetName(v string) *GetClbDomainListResponseDataResultsVipsHealthcheckers {
  s.Name = &v
  return s
}

type GetClbDomainListPaths struct {
}

func (s GetClbDomainListPaths) String() string {
  return tea.Prettify(s)
}

func (s GetClbDomainListPaths) GoString() string {
  return s.String()
}

type GetClbDomainListParameters struct {
  // {"en":"The zone ID you want to filter", "zh_CN":"要过滤的zone的ID"}
  ZoneId *int `json:"zoneId,omitempty" xml:"zoneId,omitempty"`
  // {"en":"Keywords of CLB Domain name", "zh_CN":"CLB Domain名称的关键字"}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // {"en":"page number, default is 1", "zh_CN":"页码,默认为1"}
  Page *int `json:"page,omitempty" xml:"page,omitempty"`
  // {"en":"Page size of each response, default is 25", "zh_CN":"分页大小,默认为25"}
  PageSize *int `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s GetClbDomainListParameters) String() string {
  return tea.Prettify(s)
}

func (s GetClbDomainListParameters) GoString() string {
  return s.String()
}

func (s *GetClbDomainListParameters) SetZoneId(v int) *GetClbDomainListParameters {
  s.ZoneId = &v
  return s
}

func (s *GetClbDomainListParameters) SetName(v string) *GetClbDomainListParameters {
  s.Name = &v
  return s
}

func (s *GetClbDomainListParameters) SetPage(v int) *GetClbDomainListParameters {
  s.Page = &v
  return s
}

func (s *GetClbDomainListParameters) SetPageSize(v int) *GetClbDomainListParameters {
  s.PageSize = &v
  return s
}

type GetClbDomainListRequestHeader struct {
}

func (s GetClbDomainListRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetClbDomainListRequestHeader) GoString() string {
  return s.String()
}

type GetClbDomainListResponseHeader struct {
}

func (s GetClbDomainListResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetClbDomainListResponseHeader) GoString() string {
  return s.String()
}




type GetClbDomainByIDRequest struct {
}

func (s GetClbDomainByIDRequest) String() string {
  return tea.Prettify(s)
}

func (s GetClbDomainByIDRequest) GoString() string {
  return s.String()
}

type GetClbDomainByIDResponse struct {
  // {"en":"return data", "zh_CN":"返回值"}
  Data *GetClbDomainByIDResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
  // {"en":"code,success is 0", "zh_CN":"状态码,成功为0"}
  Code *int `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"error message or Success", "zh_CN":"错误信息或Success"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s GetClbDomainByIDResponse) String() string {
  return tea.Prettify(s)
}

func (s GetClbDomainByIDResponse) GoString() string {
  return s.String()
}

func (s *GetClbDomainByIDResponse) SetData(v *GetClbDomainByIDResponseData) *GetClbDomainByIDResponse {
  s.Data = v
  return s
}

func (s *GetClbDomainByIDResponse) SetCode(v int) *GetClbDomainByIDResponse {
  s.Code = &v
  return s
}

func (s *GetClbDomainByIDResponse) SetMessage(v string) *GetClbDomainByIDResponse {
  s.Message = &v
  return s
}

type GetClbDomainByIDResponseData struct {
  // {'en':'customer info','zh_CN':'客户信息'}
  Customer *GetClbDomainByIDResponseDataCustomer `json:"customer,omitempty" xml:"customer,omitempty" require:"true" type:"Struct"`
  // {"en":"The zone id", "zh_CN":"zone的ID"}
  ZoneId *int `json:"zoneId,omitempty" xml:"zoneId,omitempty" require:"true"`
  // {"en":"The domain id", "zh_CN":"clbDomain的ID"}
  DomainId *int `json:"domainId,omitempty" xml:"domainId,omitempty" require:"true"`
  // {"en":"zone name", "zh_CN":"域名"}
  ZoneName *string `json:"zoneName,omitempty" xml:"zoneName,omitempty" require:"true"`
  // {"en":"domain name", "zh_CN":"clbDomain名称"}
  DomainName *string `json:"domainName,omitempty" xml:"domainName,omitempty" require:"true"`
  // {"en":"The unicode name of zone", "zh_CN":"Zone的Unicode名称"}
  ZoneUnicode *string `json:"zoneUnicode,omitempty" xml:"zoneUnicode,omitempty" require:"true"`
  // {"en":"The unicode name of domain", "zh_CN":"clbDomain的Unicode名称"}
  DomainUnicode *string `json:"domainUnicode,omitempty" xml:"domainUnicode,omitempty" require:"true"`
  // {"en":"STATUS CODE ( 0: NEW OR MODIFY ; 1:PENDING TO STAGE; 2: IN STAGE; 3: PENDING TO PRODUCTION;4: IN PRODUCTION; -2: ERROR TO STAGE; -4: ERROR TO PRODUCTION )", "zh_CN":"状态码（0: NEW; 1:PENDING TO STAGE; 2: IN STAGE; 3: PENDING TO PRODUCTION;4: IN PRODUCTION; -2: ERROR TO STAGE; -4: ERROR TO PRODUCTION）"}
  ZoneStatusCode *int `json:"zoneStatusCode,omitempty" xml:"zoneStatusCode,omitempty" require:"true"`
  // {"en":"status code description", "zh_CN":"状态码说明"}
  ZoneStatusText *string `json:"zoneStatusText,omitempty" xml:"zoneStatusText,omitempty" require:"true"`
  // {"en":"Published time, UTC time(eg: 2012-11-12T02:08:34Z )", "zh_CN":"发布时间，UTC时间(如: 2012-11-12T02:08:34Z )"}
  ZoneDatePublished *string `json:"zoneDatePublished,omitempty" xml:"zoneDatePublished,omitempty" require:"true"`
  // {"en":"create time, UTC time(eg: 2012-11-12T02:08:34Z )", "zh_CN":"创建时间，UTC时间(如: 2012-11-12T02:08:34Z )"}
  DateCreated *string `json:"dateCreated,omitempty" xml:"dateCreated,omitempty" require:"true"`
  // {"en":"modify time, UTC time(eg: 2012-11-12T02:08:34Z )", "zh_CN":"修改时间，UTC时间(如: 2012-11-12T02:08:34Z )"}
  DateModified *string `json:"dateModified,omitempty" xml:"dateModified,omitempty" require:"true"`
  // {"en":"Dynamic Selection Type,0:GSLB = Shortest lantency, 1:RR = Round Robin, 2:Random = Random", "zh_CN":"动态选择类型, 0:GSLB = Shortest lantency, 1:RR = Round Robin, 2:Random = Random"}
  LoadBalancingType *int `json:"loadBalancingType,omitempty" xml:"loadBalancingType,omitempty" require:"true"`
  // {"en":"Dynamic Selection Type Text", "zh_CN":"动态选择类型名称"}
  LoadBalancingTypeText *string `json:"loadBalancingTypeText,omitempty" xml:"loadBalancingTypeText,omitempty" require:"true"`
  // {"en":"Portal CLB Basic Settings > Number of answer", "zh_CN":"Portal CLB 基础设置 > Number of answer"}
  AnswerCount *int `json:"answerCount,omitempty" xml:"answerCount,omitempty" require:"true"`
  // {"en":"TTL", "zh_CN":"TTL"}
  AnswerTtl *int `json:"answerTtl,omitempty" xml:"answerTtl,omitempty" require:"true"`
  // {"en":"health checker, can be null", "zh_CN":"监控检查,可以为null"}
  Healthchecker *GetClbDomainByIDResponseDataHealthchecker `json:"healthchecker,omitempty" xml:"healthchecker,omitempty" require:"true" type:"Struct"`
  // {"en":"delay start time(second), only use when healthchecker not null", "zh_CN":"延迟开始时间,在healthchecker不为空时使用"}
  DelayStartTime *int `json:"delayStartTime,omitempty" xml:"delayStartTime,omitempty" require:"true"`
  // {"en":"slow Start Period(second), only use when healthchecker not null", "zh_CN":"慢启动期,在healthchecker不为空时使用"}
  SlowStartPeriod *int `json:"slowStartPeriod,omitempty" xml:"slowStartPeriod,omitempty" require:"true"`
  // {"en":"domain vip list", "zh_CN":"domain的VIP列表"}
  DomainVips []*GetClbDomainByIDResponseDataDomainVips `json:"domainVips,omitempty" xml:"domainVips,omitempty" require:"true" type:"Repeated"`
  // {"en":"policy config", "zh_CN":"策略配置"}
  Policies []*GetClbDomainByIDResponseDataPolicies `json:"policies,omitempty" xml:"policies,omitempty" require:"true" type:"Repeated"`
  // {"en":"etc policy config, default = null(Dyanamic Selection From All IPs)", "zh_CN":"Etc策略配置, default = null(Dyanamic Selection From All IPs)"}
  EtcPolicies []*GetClbDomainByIDResponseDataEtcPolicies `json:"etcPolicies,omitempty" xml:"etcPolicies,omitempty" require:"true" type:"Repeated"`
}

func (s GetClbDomainByIDResponseData) String() string {
  return tea.Prettify(s)
}

func (s GetClbDomainByIDResponseData) GoString() string {
  return s.String()
}

func (s *GetClbDomainByIDResponseData) SetCustomer(v *GetClbDomainByIDResponseDataCustomer) *GetClbDomainByIDResponseData {
  s.Customer = v
  return s
}

func (s *GetClbDomainByIDResponseData) SetZoneId(v int) *GetClbDomainByIDResponseData {
  s.ZoneId = &v
  return s
}

func (s *GetClbDomainByIDResponseData) SetDomainId(v int) *GetClbDomainByIDResponseData {
  s.DomainId = &v
  return s
}

func (s *GetClbDomainByIDResponseData) SetZoneName(v string) *GetClbDomainByIDResponseData {
  s.ZoneName = &v
  return s
}

func (s *GetClbDomainByIDResponseData) SetDomainName(v string) *GetClbDomainByIDResponseData {
  s.DomainName = &v
  return s
}

func (s *GetClbDomainByIDResponseData) SetZoneUnicode(v string) *GetClbDomainByIDResponseData {
  s.ZoneUnicode = &v
  return s
}

func (s *GetClbDomainByIDResponseData) SetDomainUnicode(v string) *GetClbDomainByIDResponseData {
  s.DomainUnicode = &v
  return s
}

func (s *GetClbDomainByIDResponseData) SetZoneStatusCode(v int) *GetClbDomainByIDResponseData {
  s.ZoneStatusCode = &v
  return s
}

func (s *GetClbDomainByIDResponseData) SetZoneStatusText(v string) *GetClbDomainByIDResponseData {
  s.ZoneStatusText = &v
  return s
}

func (s *GetClbDomainByIDResponseData) SetZoneDatePublished(v string) *GetClbDomainByIDResponseData {
  s.ZoneDatePublished = &v
  return s
}

func (s *GetClbDomainByIDResponseData) SetDateCreated(v string) *GetClbDomainByIDResponseData {
  s.DateCreated = &v
  return s
}

func (s *GetClbDomainByIDResponseData) SetDateModified(v string) *GetClbDomainByIDResponseData {
  s.DateModified = &v
  return s
}

func (s *GetClbDomainByIDResponseData) SetLoadBalancingType(v int) *GetClbDomainByIDResponseData {
  s.LoadBalancingType = &v
  return s
}

func (s *GetClbDomainByIDResponseData) SetLoadBalancingTypeText(v string) *GetClbDomainByIDResponseData {
  s.LoadBalancingTypeText = &v
  return s
}

func (s *GetClbDomainByIDResponseData) SetAnswerCount(v int) *GetClbDomainByIDResponseData {
  s.AnswerCount = &v
  return s
}

func (s *GetClbDomainByIDResponseData) SetAnswerTtl(v int) *GetClbDomainByIDResponseData {
  s.AnswerTtl = &v
  return s
}

func (s *GetClbDomainByIDResponseData) SetHealthchecker(v *GetClbDomainByIDResponseDataHealthchecker) *GetClbDomainByIDResponseData {
  s.Healthchecker = v
  return s
}

func (s *GetClbDomainByIDResponseData) SetDelayStartTime(v int) *GetClbDomainByIDResponseData {
  s.DelayStartTime = &v
  return s
}

func (s *GetClbDomainByIDResponseData) SetSlowStartPeriod(v int) *GetClbDomainByIDResponseData {
  s.SlowStartPeriod = &v
  return s
}

func (s *GetClbDomainByIDResponseData) SetDomainVips(v []*GetClbDomainByIDResponseDataDomainVips) *GetClbDomainByIDResponseData {
  s.DomainVips = v
  return s
}

func (s *GetClbDomainByIDResponseData) SetPolicies(v []*GetClbDomainByIDResponseDataPolicies) *GetClbDomainByIDResponseData {
  s.Policies = v
  return s
}

func (s *GetClbDomainByIDResponseData) SetEtcPolicies(v []*GetClbDomainByIDResponseDataEtcPolicies) *GetClbDomainByIDResponseData {
  s.EtcPolicies = v
  return s
}

type GetClbDomainByIDResponseDataCustomer struct {
  // {'en':'customer name','zh_CN':'客户名称'}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
}

func (s GetClbDomainByIDResponseDataCustomer) String() string {
  return tea.Prettify(s)
}

func (s GetClbDomainByIDResponseDataCustomer) GoString() string {
  return s.String()
}

func (s *GetClbDomainByIDResponseDataCustomer) SetName(v string) *GetClbDomainByIDResponseDataCustomer {
  s.Name = &v
  return s
}

type GetClbDomainByIDResponseDataHealthchecker struct {
  // {"en":"health checker id", "zh_CN":"监控检查ID"}
  Id *int `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {"en":"health checker name", "zh_CN":"监控检查名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
}

func (s GetClbDomainByIDResponseDataHealthchecker) String() string {
  return tea.Prettify(s)
}

func (s GetClbDomainByIDResponseDataHealthchecker) GoString() string {
  return s.String()
}

func (s *GetClbDomainByIDResponseDataHealthchecker) SetId(v int) *GetClbDomainByIDResponseDataHealthchecker {
  s.Id = &v
  return s
}

func (s *GetClbDomainByIDResponseDataHealthchecker) SetName(v string) *GetClbDomainByIDResponseDataHealthchecker {
  s.Name = &v
  return s
}

type GetClbDomainByIDResponseDataDomainVips struct     {
  // {"en":"Servergroup's ID", "zh_CN":"Servergroup的ID"}
  DomainVipId *int `json:"domainVipId,omitempty" xml:"domainVipId,omitempty" require:"true"`
  // {"en":"Servergroup's ID", "zh_CN":"Servergroup的ID"}
  ServergroupId *int `json:"servergroupId,omitempty" xml:"servergroupId,omitempty" require:"true"`
  // {"en":"server group's name", "zh_CN":"server group的名称"}
  ServergroupName *string `json:"servergroupName,omitempty" xml:"servergroupName,omitempty" require:"true"`
  // {"en":"Server's ID", "zh_CN":"Server的ID"}
  ServerId *int `json:"serverId,omitempty" xml:"serverId,omitempty" require:"true"`
  // {"en":"server name", "zh_CN":"服务名"}
  ServerName *string `json:"serverName,omitempty" xml:"serverName,omitempty" require:"true"`
  // {"en":"server group region id", "zh_CN":"server group的地理ID"}
  Region *string `json:"region,omitempty" xml:"region,omitempty" require:"true"`
  // {"en":"server group region name", "zh_CN":"server group的地理名称"}
  RegionName *string `json:"regionName,omitempty" xml:"regionName,omitempty" require:"true"`
  // {"en":"IP", "zh_CN":"ip地址"}
  Ip *string `json:"ip,omitempty" xml:"ip,omitempty" require:"true"`
  // {"en":"enable", "zh_CN":"是否可用"}
  IsEnabled *bool `json:"isEnabled,omitempty" xml:"isEnabled,omitempty" require:"true"`
  // {"en":"health checkers", "zh_CN":"监控检查列表"}
  Healthcheckers []*GetClbDomainByIDResponseDataDomainVipsHealthcheckers `json:"healthcheckers,omitempty" xml:"healthcheckers,omitempty" require:"true" type:"Repeated"`
}

func (s GetClbDomainByIDResponseDataDomainVips) String() string {
  return tea.Prettify(s)
}

func (s GetClbDomainByIDResponseDataDomainVips) GoString() string {
  return s.String()
}

func (s *GetClbDomainByIDResponseDataDomainVips) SetDomainVipId(v int) *GetClbDomainByIDResponseDataDomainVips {
  s.DomainVipId = &v
  return s
}

func (s *GetClbDomainByIDResponseDataDomainVips) SetServergroupId(v int) *GetClbDomainByIDResponseDataDomainVips {
  s.ServergroupId = &v
  return s
}

func (s *GetClbDomainByIDResponseDataDomainVips) SetServergroupName(v string) *GetClbDomainByIDResponseDataDomainVips {
  s.ServergroupName = &v
  return s
}

func (s *GetClbDomainByIDResponseDataDomainVips) SetServerId(v int) *GetClbDomainByIDResponseDataDomainVips {
  s.ServerId = &v
  return s
}

func (s *GetClbDomainByIDResponseDataDomainVips) SetServerName(v string) *GetClbDomainByIDResponseDataDomainVips {
  s.ServerName = &v
  return s
}

func (s *GetClbDomainByIDResponseDataDomainVips) SetRegion(v string) *GetClbDomainByIDResponseDataDomainVips {
  s.Region = &v
  return s
}

func (s *GetClbDomainByIDResponseDataDomainVips) SetRegionName(v string) *GetClbDomainByIDResponseDataDomainVips {
  s.RegionName = &v
  return s
}

func (s *GetClbDomainByIDResponseDataDomainVips) SetIp(v string) *GetClbDomainByIDResponseDataDomainVips {
  s.Ip = &v
  return s
}

func (s *GetClbDomainByIDResponseDataDomainVips) SetIsEnabled(v bool) *GetClbDomainByIDResponseDataDomainVips {
  s.IsEnabled = &v
  return s
}

func (s *GetClbDomainByIDResponseDataDomainVips) SetHealthcheckers(v []*GetClbDomainByIDResponseDataDomainVipsHealthcheckers) *GetClbDomainByIDResponseDataDomainVips {
  s.Healthcheckers = v
  return s
}

type GetClbDomainByIDResponseDataDomainVipsHealthcheckers struct     {
  // {"en":"health checker id", "zh_CN":"监控检查ID"}
  Id *int `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {"en":"health checker name", "zh_CN":"监控检查名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
}

func (s GetClbDomainByIDResponseDataDomainVipsHealthcheckers) String() string {
  return tea.Prettify(s)
}

func (s GetClbDomainByIDResponseDataDomainVipsHealthcheckers) GoString() string {
  return s.String()
}

func (s *GetClbDomainByIDResponseDataDomainVipsHealthcheckers) SetId(v int) *GetClbDomainByIDResponseDataDomainVipsHealthcheckers {
  s.Id = &v
  return s
}

func (s *GetClbDomainByIDResponseDataDomainVipsHealthcheckers) SetName(v string) *GetClbDomainByIDResponseDataDomainVipsHealthcheckers {
  s.Name = &v
  return s
}

type GetClbDomainByIDResponseDataPolicies struct     {
  // {"en":"queue number", "zh_CN":"排序号"}
  Sequence *int `json:"sequence,omitempty" xml:"sequence,omitempty" require:"true"`
  // {"en":"policy name", "zh_CN":"策略名称"}
  PolicyName *string `json:"policyName,omitempty" xml:"policyName,omitempty" require:"true"`
  // {"en":"enable", "zh_CN":"是否可用"}
  IsEnabled *bool `json:"isEnabled,omitempty" xml:"isEnabled,omitempty" require:"true"`
  // {"en":"condition", "zh_CN":"条件"}
  Condition *GetClbDomainByIDResponseDataPoliciesCondition `json:"condition,omitempty" xml:"condition,omitempty" require:"true" type:"Struct"`
  // {"en":"action", "zh_CN":"行动"}
  Actions []*GetClbDomainByIDResponseDataPoliciesActions `json:"actions,omitempty" xml:"actions,omitempty" require:"true" type:"Repeated"`
}

func (s GetClbDomainByIDResponseDataPolicies) String() string {
  return tea.Prettify(s)
}

func (s GetClbDomainByIDResponseDataPolicies) GoString() string {
  return s.String()
}

func (s *GetClbDomainByIDResponseDataPolicies) SetSequence(v int) *GetClbDomainByIDResponseDataPolicies {
  s.Sequence = &v
  return s
}

func (s *GetClbDomainByIDResponseDataPolicies) SetPolicyName(v string) *GetClbDomainByIDResponseDataPolicies {
  s.PolicyName = &v
  return s
}

func (s *GetClbDomainByIDResponseDataPolicies) SetIsEnabled(v bool) *GetClbDomainByIDResponseDataPolicies {
  s.IsEnabled = &v
  return s
}

func (s *GetClbDomainByIDResponseDataPolicies) SetCondition(v *GetClbDomainByIDResponseDataPoliciesCondition) *GetClbDomainByIDResponseDataPolicies {
  s.Condition = v
  return s
}

func (s *GetClbDomainByIDResponseDataPolicies) SetActions(v []*GetClbDomainByIDResponseDataPoliciesActions) *GetClbDomainByIDResponseDataPolicies {
  s.Actions = v
  return s
}

type GetClbDomainByIDResponseDataPoliciesCondition struct {
  // {"en":"The condition id", "zh_CN":"条件ID"}
  ConditionId *int `json:"conditionId,omitempty" xml:"conditionId,omitempty" require:"true"`
  // {"en":"condition type, 0:IP RANGE, 1: GEO DATA, 2 Ratio(%)", "zh_CN":"条件类型, 0:IP RANGE, 1: GEO DATA, 2 Ratio(%)"}
  ConditionType *int `json:"conditionType,omitempty" xml:"conditionType,omitempty" require:"true"`
  // {"en":"IP range,comma separated,only useful if condition type is 0(ip range), eg: 1.2.3.4,1.5.6.0/24", "zh_CN":"IP范围,逗号分隔,只有在contidion type为0(ip range)时有用, 例:1.2.3.4,1.5.6.0/24"}
  IpRange *string `json:"ipRange,omitempty" xml:"ipRange,omitempty" require:"true"`
  // {"en":"invert, only useful if condition type is 0(ip range) or 1(geo data) ", "zh_CN":"反转, 只有在contidion type为0(ip range)或1(geo data)时有用"}
  Invert *bool `json:"invert,omitempty" xml:"invert,omitempty" require:"true"`
  // {"en":"region id list,only useful if condition type is 1(geo data), get by regions api", "zh_CN":"地理位置的ID列表，只有在contidion type为1(geo data)时有用，从regions接口获取"}
  Region []*string `json:"region,omitempty" xml:"region,omitempty" require:"true" type:"Repeated"`
  // {"en":"config view id list,only useful if condition type is 1(geo data), get by cvs api", "zh_CN":"配置视图ID列表，只有在contidion type为1(geo data)时有用，从cvs接口获取"}
  Cv []*string `json:"cv,omitempty" xml:"cv,omitempty" require:"true" type:"Repeated"`
  // {"en":"isp id list,only useful if condition type is 1(geo data), get by isps api", "zh_CN":"供应商ID列表，只有在contidion type为1(geo data)时有用，从isps接口获取"}
  Isp []*string `json:"isp,omitempty" xml:"isp,omitempty" require:"true" type:"Repeated"`
  // {"en":"as list, only useful if condition type is 1(geo data)", "zh_CN":"as列表，只有在contidion type为1(geo data)时有用"}
  Asn []*int `json:"asn,omitempty" xml:"asn,omitempty" require:"true" type:"Repeated"`
  // {"en":"percent,Only use conditionType = 2(Ratio(%)), eg: 70", "zh_CN":"比例，只有在conditionType=2 (Ratio%)时有用，如: 70"}
  Percent *int `json:"percent,omitempty" xml:"percent,omitempty" require:"true"`
}

func (s GetClbDomainByIDResponseDataPoliciesCondition) String() string {
  return tea.Prettify(s)
}

func (s GetClbDomainByIDResponseDataPoliciesCondition) GoString() string {
  return s.String()
}

func (s *GetClbDomainByIDResponseDataPoliciesCondition) SetConditionId(v int) *GetClbDomainByIDResponseDataPoliciesCondition {
  s.ConditionId = &v
  return s
}

func (s *GetClbDomainByIDResponseDataPoliciesCondition) SetConditionType(v int) *GetClbDomainByIDResponseDataPoliciesCondition {
  s.ConditionType = &v
  return s
}

func (s *GetClbDomainByIDResponseDataPoliciesCondition) SetIpRange(v string) *GetClbDomainByIDResponseDataPoliciesCondition {
  s.IpRange = &v
  return s
}

func (s *GetClbDomainByIDResponseDataPoliciesCondition) SetInvert(v bool) *GetClbDomainByIDResponseDataPoliciesCondition {
  s.Invert = &v
  return s
}

func (s *GetClbDomainByIDResponseDataPoliciesCondition) SetRegion(v []*string) *GetClbDomainByIDResponseDataPoliciesCondition {
  s.Region = v
  return s
}

func (s *GetClbDomainByIDResponseDataPoliciesCondition) SetCv(v []*string) *GetClbDomainByIDResponseDataPoliciesCondition {
  s.Cv = v
  return s
}

func (s *GetClbDomainByIDResponseDataPoliciesCondition) SetIsp(v []*string) *GetClbDomainByIDResponseDataPoliciesCondition {
  s.Isp = v
  return s
}

func (s *GetClbDomainByIDResponseDataPoliciesCondition) SetAsn(v []*int) *GetClbDomainByIDResponseDataPoliciesCondition {
  s.Asn = v
  return s
}

func (s *GetClbDomainByIDResponseDataPoliciesCondition) SetPercent(v int) *GetClbDomainByIDResponseDataPoliciesCondition {
  s.Percent = &v
  return s
}

type GetClbDomainByIDResponseDataPoliciesActions struct     {
  // {"en":"queue number", "zh_CN":"排序号"}
  Sequence *int `json:"sequence,omitempty" xml:"sequence,omitempty" require:"true"`
  // {"en":"The action id", "zh_CN":"行动ID"}
  ActionId *int `json:"actionId,omitempty" xml:"actionId,omitempty" require:"true"`
  // {"en":"action type, 0:deny, 1:vip, 2:cname,3:a", "zh_CN":"行动类别: 0:deny, 1:vip, 2:cname,3:a"}
  ActionType *int `json:"actionType,omitempty" xml:"actionType,omitempty" require:"true"`
  // {"en":"vip ID, only useful if actionType=1(vip), get in domainVips serverId", "zh_CN":"VIP ID,只有在action type为1(vip)时有用，是属性domainVips中的serverId值"}
  Vips []*int `json:"vips,omitempty" xml:"vips,omitempty" require:"true" type:"Repeated"`
  // {"en":"only useful if actionType=3(a) or actionType=2(cname)", "zh_CN":"只有在actionType=3(a)或actionType=2(cname)时有用"}
  Answer []*GetClbDomainByIDResponseDataPoliciesActionsAnswer `json:"answer,omitempty" xml:"answer,omitempty" require:"true" type:"Repeated"`
}

func (s GetClbDomainByIDResponseDataPoliciesActions) String() string {
  return tea.Prettify(s)
}

func (s GetClbDomainByIDResponseDataPoliciesActions) GoString() string {
  return s.String()
}

func (s *GetClbDomainByIDResponseDataPoliciesActions) SetSequence(v int) *GetClbDomainByIDResponseDataPoliciesActions {
  s.Sequence = &v
  return s
}

func (s *GetClbDomainByIDResponseDataPoliciesActions) SetActionId(v int) *GetClbDomainByIDResponseDataPoliciesActions {
  s.ActionId = &v
  return s
}

func (s *GetClbDomainByIDResponseDataPoliciesActions) SetActionType(v int) *GetClbDomainByIDResponseDataPoliciesActions {
  s.ActionType = &v
  return s
}

func (s *GetClbDomainByIDResponseDataPoliciesActions) SetVips(v []*int) *GetClbDomainByIDResponseDataPoliciesActions {
  s.Vips = v
  return s
}

func (s *GetClbDomainByIDResponseDataPoliciesActions) SetAnswer(v []*GetClbDomainByIDResponseDataPoliciesActionsAnswer) *GetClbDomainByIDResponseDataPoliciesActions {
  s.Answer = v
  return s
}

type GetClbDomainByIDResponseDataPoliciesActionsAnswer struct     {
  // {"en":"actionType=3, A or actionType=2, CNAME", "zh_CN":"actionType=3,则值为A,actionType=2,则值为CNAME"}
  Rtype *string `json:"rtype,omitempty" xml:"rtype,omitempty" require:"true"`
  // {"en":"when actionType=3,this is IP, if actionType=2,that is domain name,eg: 1.2.3.4 or test.com", "zh_CN":"当actionType=3时，值为IP地址，当actionType=2时，值为域名地址"}
  Data *string `json:"data,omitempty" xml:"data,omitempty" require:"true"`
  // {"en":"TTL", "zh_CN":"TTL"}
  Ttl *int `json:"ttl,omitempty" xml:"ttl,omitempty" require:"true"`
}

func (s GetClbDomainByIDResponseDataPoliciesActionsAnswer) String() string {
  return tea.Prettify(s)
}

func (s GetClbDomainByIDResponseDataPoliciesActionsAnswer) GoString() string {
  return s.String()
}

func (s *GetClbDomainByIDResponseDataPoliciesActionsAnswer) SetRtype(v string) *GetClbDomainByIDResponseDataPoliciesActionsAnswer {
  s.Rtype = &v
  return s
}

func (s *GetClbDomainByIDResponseDataPoliciesActionsAnswer) SetData(v string) *GetClbDomainByIDResponseDataPoliciesActionsAnswer {
  s.Data = &v
  return s
}

func (s *GetClbDomainByIDResponseDataPoliciesActionsAnswer) SetTtl(v int) *GetClbDomainByIDResponseDataPoliciesActionsAnswer {
  s.Ttl = &v
  return s
}

type GetClbDomainByIDResponseDataEtcPolicies struct     {
  // {"en":"policy name(always)", "zh_CN":"策略名称(always)"}
  PolicyName *string `json:"policyName,omitempty" xml:"policyName,omitempty" require:"true"`
  // {"en":"action", "zh_CN":"行动"}
  Actions []*GetClbDomainByIDResponseDataEtcPoliciesActions `json:"actions,omitempty" xml:"actions,omitempty" require:"true" type:"Repeated"`
}

func (s GetClbDomainByIDResponseDataEtcPolicies) String() string {
  return tea.Prettify(s)
}

func (s GetClbDomainByIDResponseDataEtcPolicies) GoString() string {
  return s.String()
}

func (s *GetClbDomainByIDResponseDataEtcPolicies) SetPolicyName(v string) *GetClbDomainByIDResponseDataEtcPolicies {
  s.PolicyName = &v
  return s
}

func (s *GetClbDomainByIDResponseDataEtcPolicies) SetActions(v []*GetClbDomainByIDResponseDataEtcPoliciesActions) *GetClbDomainByIDResponseDataEtcPolicies {
  s.Actions = v
  return s
}

type GetClbDomainByIDResponseDataEtcPoliciesActions struct     {
  // {"en":"action type, 0:deny, 2:cname,3:a", "zh_CN":"行动类别: 0:deny, 2:cname,3:a"}
  ActionType *int `json:"actionType,omitempty" xml:"actionType,omitempty" require:"true"`
  // {"en":"only useful if actionType=3(a) or actionType=2(cname), default is null", "zh_CN":"只有在actionType=3(a)或actionType=2(cname)时有用,默认为null"}
  Answer []*GetClbDomainByIDResponseDataEtcPoliciesActionsAnswer `json:"answer,omitempty" xml:"answer,omitempty" require:"true" type:"Repeated"`
}

func (s GetClbDomainByIDResponseDataEtcPoliciesActions) String() string {
  return tea.Prettify(s)
}

func (s GetClbDomainByIDResponseDataEtcPoliciesActions) GoString() string {
  return s.String()
}

func (s *GetClbDomainByIDResponseDataEtcPoliciesActions) SetActionType(v int) *GetClbDomainByIDResponseDataEtcPoliciesActions {
  s.ActionType = &v
  return s
}

func (s *GetClbDomainByIDResponseDataEtcPoliciesActions) SetAnswer(v []*GetClbDomainByIDResponseDataEtcPoliciesActionsAnswer) *GetClbDomainByIDResponseDataEtcPoliciesActions {
  s.Answer = v
  return s
}

type GetClbDomainByIDResponseDataEtcPoliciesActionsAnswer struct     {
  // {"en":"actionType=3, A or actionType=2, CNAME", "zh_CN":"actionType=3,则值为A,actionType=2,则值为CNAME"}
  Rtype *string `json:"rtype,omitempty" xml:"rtype,omitempty" require:"true"`
  // {"en":"when actionType=3,this is IP, if actionType=2,that is domain name,eg: 1.2.3.4 or test.com", "zh_CN":"当actionType=3时，值为IP地址，当actionType=2时，值为域名地址"}
  Data *string `json:"data,omitempty" xml:"data,omitempty" require:"true"`
  // {"en":"TTL", "zh_CN":"TTL"}
  Ttl *int `json:"ttl,omitempty" xml:"ttl,omitempty" require:"true"`
}

func (s GetClbDomainByIDResponseDataEtcPoliciesActionsAnswer) String() string {
  return tea.Prettify(s)
}

func (s GetClbDomainByIDResponseDataEtcPoliciesActionsAnswer) GoString() string {
  return s.String()
}

func (s *GetClbDomainByIDResponseDataEtcPoliciesActionsAnswer) SetRtype(v string) *GetClbDomainByIDResponseDataEtcPoliciesActionsAnswer {
  s.Rtype = &v
  return s
}

func (s *GetClbDomainByIDResponseDataEtcPoliciesActionsAnswer) SetData(v string) *GetClbDomainByIDResponseDataEtcPoliciesActionsAnswer {
  s.Data = &v
  return s
}

func (s *GetClbDomainByIDResponseDataEtcPoliciesActionsAnswer) SetTtl(v int) *GetClbDomainByIDResponseDataEtcPoliciesActionsAnswer {
  s.Ttl = &v
  return s
}

type GetClbDomainByIDPaths struct {
  // {"en":"CLB Domain's ID", "zh_CN":"CLB Domain的ID"}
  DomainId *int `json:"domainId,omitempty" xml:"domainId,omitempty" require:"true"`
}

func (s GetClbDomainByIDPaths) String() string {
  return tea.Prettify(s)
}

func (s GetClbDomainByIDPaths) GoString() string {
  return s.String()
}

func (s *GetClbDomainByIDPaths) SetDomainId(v int) *GetClbDomainByIDPaths {
  s.DomainId = &v
  return s
}

type GetClbDomainByIDParameters struct {
}

func (s GetClbDomainByIDParameters) String() string {
  return tea.Prettify(s)
}

func (s GetClbDomainByIDParameters) GoString() string {
  return s.String()
}

type GetClbDomainByIDRequestHeader struct {
}

func (s GetClbDomainByIDRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetClbDomainByIDRequestHeader) GoString() string {
  return s.String()
}

type GetClbDomainByIDResponseHeader struct {
}

func (s GetClbDomainByIDResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetClbDomainByIDResponseHeader) GoString() string {
  return s.String()
}




type DeleteSpecifiedClbDomainRequest struct {
}

func (s DeleteSpecifiedClbDomainRequest) String() string {
  return tea.Prettify(s)
}

func (s DeleteSpecifiedClbDomainRequest) GoString() string {
  return s.String()
}

type DeleteSpecifiedClbDomainResponse struct {
  // {"en":"return data", "zh_CN":"返回值"}
  Data *string `json:"data,omitempty" xml:"data,omitempty" require:"true"`
  // {"en":"code,success is 0", "zh_CN":"状态码,成功为0"}
  Code *int `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"error message or Success", "zh_CN":"错误信息或Success"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s DeleteSpecifiedClbDomainResponse) String() string {
  return tea.Prettify(s)
}

func (s DeleteSpecifiedClbDomainResponse) GoString() string {
  return s.String()
}

func (s *DeleteSpecifiedClbDomainResponse) SetData(v string) *DeleteSpecifiedClbDomainResponse {
  s.Data = &v
  return s
}

func (s *DeleteSpecifiedClbDomainResponse) SetCode(v int) *DeleteSpecifiedClbDomainResponse {
  s.Code = &v
  return s
}

func (s *DeleteSpecifiedClbDomainResponse) SetMessage(v string) *DeleteSpecifiedClbDomainResponse {
  s.Message = &v
  return s
}

type DeleteSpecifiedClbDomainPaths struct {
  // {"en":"CLB Domain's ID", "zh_CN":"CLB Domain的ID"}
  DomainId *int `json:"domainId,omitempty" xml:"domainId,omitempty" require:"true"`
}

func (s DeleteSpecifiedClbDomainPaths) String() string {
  return tea.Prettify(s)
}

func (s DeleteSpecifiedClbDomainPaths) GoString() string {
  return s.String()
}

func (s *DeleteSpecifiedClbDomainPaths) SetDomainId(v int) *DeleteSpecifiedClbDomainPaths {
  s.DomainId = &v
  return s
}

type DeleteSpecifiedClbDomainParameters struct {
}

func (s DeleteSpecifiedClbDomainParameters) String() string {
  return tea.Prettify(s)
}

func (s DeleteSpecifiedClbDomainParameters) GoString() string {
  return s.String()
}

type DeleteSpecifiedClbDomainRequestHeader struct {
}

func (s DeleteSpecifiedClbDomainRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteSpecifiedClbDomainRequestHeader) GoString() string {
  return s.String()
}

type DeleteSpecifiedClbDomainResponseHeader struct {
}

func (s DeleteSpecifiedClbDomainResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteSpecifiedClbDomainResponseHeader) GoString() string {
  return s.String()
}




