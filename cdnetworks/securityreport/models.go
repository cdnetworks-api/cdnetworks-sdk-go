package securityreport

// This file is auto-generated, don't edit it. Thanks.
import (
  "github.com/alibabacloud-go/tea/tea"
)

type GetbotrequestuseragentTopdataRequest struct {
  // {"en":"Domain.Separate by ';'.", "zh_CN":"域名。多个以;隔开"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {"en":"StartTime. Format: yyyy-MM-dd HH:mm:ss", "zh_CN":"开始时间。格式： yyyy-MM-dd HH:mm:ss"}
  StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty" require:"true"`
  // {"en":"EndTime. Format: yyyy-MM-dd HH:mm:ss", "zh_CN":"结束时间。格式： yyyy-MM-dd HH:mm:ss"}
  EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty" require:"true"`
  // {"en":"Time zone. Default 8, i.e.'GTM+8'", "zh_CN":"时区，默认8，即“GTM+8”"}
  TimeZone *int `json:"timeZone,omitempty" xml:"timeZone,omitempty"`
  // {"en":"Number of top values.Default value 10", "zh_CN":"排名最前值数目。默认10"}
  TopNum *int `json:"topNum,omitempty" xml:"topNum,omitempty"`
}

func (s GetbotrequestuseragentTopdataRequest) String() string {
  return tea.Prettify(s)
}

func (s GetbotrequestuseragentTopdataRequest) GoString() string {
  return s.String()
}

func (s *GetbotrequestuseragentTopdataRequest) SetDomain(v string) *GetbotrequestuseragentTopdataRequest {
  s.Domain = &v
  return s
}

func (s *GetbotrequestuseragentTopdataRequest) SetStartTime(v string) *GetbotrequestuseragentTopdataRequest {
  s.StartTime = &v
  return s
}

func (s *GetbotrequestuseragentTopdataRequest) SetEndTime(v string) *GetbotrequestuseragentTopdataRequest {
  s.EndTime = &v
  return s
}

func (s *GetbotrequestuseragentTopdataRequest) SetTimeZone(v int) *GetbotrequestuseragentTopdataRequest {
  s.TimeZone = &v
  return s
}

func (s *GetbotrequestuseragentTopdataRequest) SetTopNum(v int) *GetbotrequestuseragentTopdataRequest {
  s.TopNum = &v
  return s
}

type GetbotrequestuseragentTopdataResponse struct {
  // {"en":"Status code, success is '200'.", "zh_CN":"状态码，成功为“200”。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Return message, success is 'Success'.", "zh_CN":"返回信息，成功为“Success”。"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Data returned.", "zh_CN":"返回数据。"}
  Data []*GetbotrequestuseragentTopdataResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s GetbotrequestuseragentTopdataResponse) String() string {
  return tea.Prettify(s)
}

func (s GetbotrequestuseragentTopdataResponse) GoString() string {
  return s.String()
}

func (s *GetbotrequestuseragentTopdataResponse) SetCode(v string) *GetbotrequestuseragentTopdataResponse {
  s.Code = &v
  return s
}

func (s *GetbotrequestuseragentTopdataResponse) SetMessage(v string) *GetbotrequestuseragentTopdataResponse {
  s.Message = &v
  return s
}

func (s *GetbotrequestuseragentTopdataResponse) SetData(v []*GetbotrequestuseragentTopdataResponseData) *GetbotrequestuseragentTopdataResponse {
  s.Data = v
  return s
}

type GetbotrequestuseragentTopdataResponseData struct     {
  // {"en":"Statistical type.", "zh_CN":"统计类型。"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"Bot request count.", "zh_CN":"bot请求数。"}
  Count *int64 `json:"count,omitempty" xml:"count,omitempty" require:"true"`
}

func (s GetbotrequestuseragentTopdataResponseData) String() string {
  return tea.Prettify(s)
}

func (s GetbotrequestuseragentTopdataResponseData) GoString() string {
  return s.String()
}

func (s *GetbotrequestuseragentTopdataResponseData) SetName(v string) *GetbotrequestuseragentTopdataResponseData {
  s.Name = &v
  return s
}

func (s *GetbotrequestuseragentTopdataResponseData) SetCount(v int64) *GetbotrequestuseragentTopdataResponseData {
  s.Count = &v
  return s
}

type GetbotrequestuseragentTopdataPaths struct {
}

func (s GetbotrequestuseragentTopdataPaths) String() string {
  return tea.Prettify(s)
}

func (s GetbotrequestuseragentTopdataPaths) GoString() string {
  return s.String()
}

type GetbotrequestuseragentTopdataParameters struct {
}

func (s GetbotrequestuseragentTopdataParameters) String() string {
  return tea.Prettify(s)
}

func (s GetbotrequestuseragentTopdataParameters) GoString() string {
  return s.String()
}

type GetbotrequestuseragentTopdataRequestHeader struct {
}

func (s GetbotrequestuseragentTopdataRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetbotrequestuseragentTopdataRequestHeader) GoString() string {
  return s.String()
}

type GetbotrequestuseragentTopdataResponseHeader struct {
}

func (s GetbotrequestuseragentTopdataResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetbotrequestuseragentTopdataResponseHeader) GoString() string {
  return s.String()
}




type GetTriggeredCustomRulesRequest struct {
  // {"defaultValue":"10","en":"Top rankings, default value: 10, max: 1000.","zh_CN":"取前几排名，默认：10，上限：1000。"}
  Top *int `json:"top,omitempty" xml:"top,omitempty"`
  // {"en":"Multiple selection. Handling results, default: display all results.\nmitigated: Number of mitigated requests.\nmonitored: Number of observed requests.","zh_CN":"多选。处置结果，默认：展示所有结果。\nmitigated：已抵御请求数。\nmonitored：观察请求数。","exampleValue":"mitigated,monitored"}
  ActType []*string `json:"actType,omitempty" xml:"actType,omitempty" type:"Repeated"`
  // {"en":"Start time, format: yyyy-MM-dd HH:mm:ss.\nThe time range can not exceed 31 days.","zh_CN":"起始时间，格式：yyyy-MM-dd HH:mm:ss。\n时间范围不超过31天。"}
  StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty" require:"true"`
  // {"en":"End time, format: yyyy-MM-dd HH:mm:ss.\nThe time range can not exceed 31 days.","zh_CN":"截止时间，格式：yyyy-MM-dd HH:mm:ss。\n时间范围不超过31天。"}
  EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty" require:"true"`
  // {"en":"Domain list. Queries all domains under the account when not specified.","zh_CN":"域名列表，未指定时查询账号下的所有域名。"}
  Domains []*string `json:"domains,omitempty" xml:"domains,omitempty" type:"Repeated"`
}

func (s GetTriggeredCustomRulesRequest) String() string {
  return tea.Prettify(s)
}

func (s GetTriggeredCustomRulesRequest) GoString() string {
  return s.String()
}

func (s *GetTriggeredCustomRulesRequest) SetTop(v int) *GetTriggeredCustomRulesRequest {
  s.Top = &v
  return s
}

func (s *GetTriggeredCustomRulesRequest) SetActType(v []*string) *GetTriggeredCustomRulesRequest {
  s.ActType = v
  return s
}

func (s *GetTriggeredCustomRulesRequest) SetStartTime(v string) *GetTriggeredCustomRulesRequest {
  s.StartTime = &v
  return s
}

func (s *GetTriggeredCustomRulesRequest) SetEndTime(v string) *GetTriggeredCustomRulesRequest {
  s.EndTime = &v
  return s
}

func (s *GetTriggeredCustomRulesRequest) SetDomains(v []*string) *GetTriggeredCustomRulesRequest {
  s.Domains = v
  return s
}

type GetTriggeredCustomRulesRequestHeader struct {
  // {"defaultValue":"en","en":"The language of response data, default value: en.\nzh_CN: Chinese\nen: English","zh_CN":"返回内容的语言版本，默认值: en。\nzh_CN：中文\nen：英文","exampleValue":"zh_CN,en"}
  Language *string `json:"language,omitempty" xml:"language,omitempty"`
  // {"dictionary":"belong=WAAP-MS-Ext|dict=waap_serviceType","en":"Security service type. Please enter a specific service type, if you purchase multiple security services.","zh_CN":"安全服务类型。有使用多个不同的安全服务时，需要填写具体的服务类型。"}
  ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty"`
  // {"defaultValue":"GMT+8","en":"Report Data Timezone:1. Indicates the timezone for the report data. lt must be relative to GMT andspecified in the format GMT+n where -12<=n <= 12. For example,timezone=GMT-7 and timezone=GMT+8.2. If the parameter is unspecified,results will be in the GMT+8 timezone.","zh_CN":"报表数据时区：1.请传递GMT时区。示例格式：GMT+N（其中12<=N <= 12），示例：timezone=GMT-7或timezone=GMT+8.2.若参数为传递，将默认按GMT+8时区查询"}
  Timezone *string `json:"timezone,omitempty" xml:"timezone,omitempty"`
}

func (s GetTriggeredCustomRulesRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetTriggeredCustomRulesRequestHeader) GoString() string {
  return s.String()
}

func (s *GetTriggeredCustomRulesRequestHeader) SetLanguage(v string) *GetTriggeredCustomRulesRequestHeader {
  s.Language = &v
  return s
}

func (s *GetTriggeredCustomRulesRequestHeader) SetServiceType(v string) *GetTriggeredCustomRulesRequestHeader {
  s.ServiceType = &v
  return s
}

func (s *GetTriggeredCustomRulesRequestHeader) SetTimezone(v string) *GetTriggeredCustomRulesRequestHeader {
  s.Timezone = &v
  return s
}

type GetTriggeredCustomRulesPaths struct {
}

func (s GetTriggeredCustomRulesPaths) String() string {
  return tea.Prettify(s)
}

func (s GetTriggeredCustomRulesPaths) GoString() string {
  return s.String()
}

type GetTriggeredCustomRulesParameters struct {
}

func (s GetTriggeredCustomRulesParameters) String() string {
  return tea.Prettify(s)
}

func (s GetTriggeredCustomRulesParameters) GoString() string {
  return s.String()
}

type GetTriggeredCustomRulesResponse struct {
  // {"dictionary":"belong=WAAP-MS-Ext|dict=waap_retCodeEnum","en":"Please refer to the error code for exceptions.","zh_CN":"请参照错误码。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Description.","zh_CN":"描述信息。"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"Data.","zh_CN":"出参数据。"}
  Data []*GetTriggeredCustomRulesResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s GetTriggeredCustomRulesResponse) String() string {
  return tea.Prettify(s)
}

func (s GetTriggeredCustomRulesResponse) GoString() string {
  return s.String()
}

func (s *GetTriggeredCustomRulesResponse) SetCode(v string) *GetTriggeredCustomRulesResponse {
  s.Code = &v
  return s
}

func (s *GetTriggeredCustomRulesResponse) SetMsg(v string) *GetTriggeredCustomRulesResponse {
  s.Msg = &v
  return s
}

func (s *GetTriggeredCustomRulesResponse) SetData(v []*GetTriggeredCustomRulesResponseData) *GetTriggeredCustomRulesResponse {
  s.Data = v
  return s
}

type GetTriggeredCustomRulesResponseData struct     {
  // {"en":"Rule ID.","zh_CN":"规则ID。"}
  RuleId *string `json:"ruleId,omitempty" xml:"ruleId,omitempty" require:"true"`
  // {"en":"Rule name.","zh_CN":"规则名称。"}
  RuleName *string `json:"ruleName,omitempty" xml:"ruleName,omitempty" require:"true"`
  // {"en":"Protected target.","zh_CN":"业务场景。"}
  Scene *string `json:"scene,omitempty" xml:"scene,omitempty" require:"true"`
  // {"en":"Trigger times, sort by times.","zh_CN":"触发次数，按次数排序。"}
  Hits []*GetTriggeredCustomRulesResponseDataHits `json:"hits,omitempty" xml:"hits,omitempty" require:"true" type:"Repeated"`
}

func (s GetTriggeredCustomRulesResponseData) String() string {
  return tea.Prettify(s)
}

func (s GetTriggeredCustomRulesResponseData) GoString() string {
  return s.String()
}

func (s *GetTriggeredCustomRulesResponseData) SetRuleId(v string) *GetTriggeredCustomRulesResponseData {
  s.RuleId = &v
  return s
}

func (s *GetTriggeredCustomRulesResponseData) SetRuleName(v string) *GetTriggeredCustomRulesResponseData {
  s.RuleName = &v
  return s
}

func (s *GetTriggeredCustomRulesResponseData) SetScene(v string) *GetTriggeredCustomRulesResponseData {
  s.Scene = &v
  return s
}

func (s *GetTriggeredCustomRulesResponseData) SetHits(v []*GetTriggeredCustomRulesResponseDataHits) *GetTriggeredCustomRulesResponseData {
  s.Hits = v
  return s
}

type GetTriggeredCustomRulesResponseDataHits struct     {
  // {"en":"Action.","zh_CN":"采取动作。"}
  Act *string `json:"act,omitempty" xml:"act,omitempty" require:"true"`
  // {"en":"Hit times.","zh_CN":"命中数。"}
  Value *int64 `json:"value,omitempty" xml:"value,omitempty" require:"true"`
}

func (s GetTriggeredCustomRulesResponseDataHits) String() string {
  return tea.Prettify(s)
}

func (s GetTriggeredCustomRulesResponseDataHits) GoString() string {
  return s.String()
}

func (s *GetTriggeredCustomRulesResponseDataHits) SetAct(v string) *GetTriggeredCustomRulesResponseDataHits {
  s.Act = &v
  return s
}

func (s *GetTriggeredCustomRulesResponseDataHits) SetValue(v int64) *GetTriggeredCustomRulesResponseDataHits {
  s.Value = &v
  return s
}

type GetTriggeredCustomRulesResponseHeader struct {
}

func (s GetTriggeredCustomRulesResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetTriggeredCustomRulesResponseHeader) GoString() string {
  return s.String()
}




type L4DdosEventRequest struct {
  // {"en":"Start time, format: yyyy-MM-dd HH:mm:ss.\nThe time range can not exceed 31 days.","zh_CN":"开始时间，格式：yyyy-MM-dd HH:mm:ss。\n时间范围不超过30天。"}
  StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty" require:"true"`
  // {"en":"End time, format: yyyy-MM-dd HH:mm:ss.\nThe time range can not exceed 31 days.","zh_CN":"结束时间，格式：yyyy-MM-dd HH:mm:ss。\n时间范围不超过30天。"}
  EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty" require:"true"`
}

func (s L4DdosEventRequest) String() string {
  return tea.Prettify(s)
}

func (s L4DdosEventRequest) GoString() string {
  return s.String()
}

func (s *L4DdosEventRequest) SetStartTime(v string) *L4DdosEventRequest {
  s.StartTime = &v
  return s
}

func (s *L4DdosEventRequest) SetEndTime(v string) *L4DdosEventRequest {
  s.EndTime = &v
  return s
}

type L4DdosEventRequestHeader struct {
  // {"dictionary":"belong=WAAP-MS-Ext|dict=waap_serviceType","en":"Security service type. Please enter a specific service type, if you purchase multiple security services.","zh_CN":"安全服务类型。有使用多个不同的安全服务时，需要填写具体的服务类型。"}
  ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty"`
  // {"defaultValue":"GMT+8","en":"Report Data Timezone:1. Indicates the timezone for the report data. lt must be relative to GMT andspecified in the format GMT+n where -12<=n <= 12. For example,timezone=GMT-7 and timezone=GMT+8.2. If the parameter is unspecified,results will be in the GMT+8 timezone.","zh_CN":"报表数据时区：1.请传递GMT时区。示例格式：GMT+N（其中12<=N <= 12），示例：timezone=GMT-7或timezone=GMT+8.2.若参数为传递，将默认按GMT+8时区查询"}
  Timezone *string `json:"timezone,omitempty" xml:"timezone,omitempty"`
}

func (s L4DdosEventRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s L4DdosEventRequestHeader) GoString() string {
  return s.String()
}

func (s *L4DdosEventRequestHeader) SetServiceType(v string) *L4DdosEventRequestHeader {
  s.ServiceType = &v
  return s
}

func (s *L4DdosEventRequestHeader) SetTimezone(v string) *L4DdosEventRequestHeader {
  s.Timezone = &v
  return s
}

type L4DdosEventPaths struct {
}

func (s L4DdosEventPaths) String() string {
  return tea.Prettify(s)
}

func (s L4DdosEventPaths) GoString() string {
  return s.String()
}

type L4DdosEventParameters struct {
}

func (s L4DdosEventParameters) String() string {
  return tea.Prettify(s)
}

func (s L4DdosEventParameters) GoString() string {
  return s.String()
}

type L4DdosEventResponse struct {
  // {"dictionary":"belong=WAAP-MS-Ext|dict=waap_retCodeEnum","en":"Please refer to the error code for exceptions.","zh_CN":"请参照错误码。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Description.","zh_CN":"描述信息。"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"Data.","zh_CN":"出参数据。"}
  Data []*L4DdosEventResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s L4DdosEventResponse) String() string {
  return tea.Prettify(s)
}

func (s L4DdosEventResponse) GoString() string {
  return s.String()
}

func (s *L4DdosEventResponse) SetCode(v string) *L4DdosEventResponse {
  s.Code = &v
  return s
}

func (s *L4DdosEventResponse) SetMsg(v string) *L4DdosEventResponse {
  s.Msg = &v
  return s
}

func (s *L4DdosEventResponse) SetData(v []*L4DdosEventResponseData) *L4DdosEventResponse {
  s.Data = v
  return s
}

type L4DdosEventResponseData struct     {
  // {"en":"Duration.","zh_CN":"持续时间。"}
  Duration *string `json:"duration,omitempty" xml:"duration,omitempty" require:"true"`
  // {"en":"Start time.","zh_CN":"开始时间。"}
  StartTime *string `json:"start_time,omitempty" xml:"start_time,omitempty" require:"true"`
  // {"en":"Attack type.","zh_CN":"攻击类型。"}
  AttackType []*string `json:"attack_type,omitempty" xml:"attack_type,omitempty" require:"true" type:"Repeated"`
  // {"en":"End time.","zh_CN":"结束时间。"}
  EndTime *string `json:"end_time,omitempty" xml:"end_time,omitempty" require:"true"`
  // {"en":"Peak Attack Time.","zh_CN":"峰值时间。"}
  MaxTime *string `json:"max_time,omitempty" xml:"max_time,omitempty" require:"true"`
  // {"en":"Status.1:DDoS Mitigation Completed; 0:DDoS Mitigation Active","zh_CN":"状态 1.清洗结束 0.清洗中。","exampleValue":"1,0"}
  Status *string `json:"status,omitempty" xml:"status,omitempty" require:"true"`
  // {"en":"Peak bandwidth.","zh_CN":"峰值（Mbps）。"}
  MaxValue *string `json:"max_value,omitempty" xml:"max_value,omitempty" require:"true"`
}

func (s L4DdosEventResponseData) String() string {
  return tea.Prettify(s)
}

func (s L4DdosEventResponseData) GoString() string {
  return s.String()
}

func (s *L4DdosEventResponseData) SetDuration(v string) *L4DdosEventResponseData {
  s.Duration = &v
  return s
}

func (s *L4DdosEventResponseData) SetStartTime(v string) *L4DdosEventResponseData {
  s.StartTime = &v
  return s
}

func (s *L4DdosEventResponseData) SetAttackType(v []*string) *L4DdosEventResponseData {
  s.AttackType = v
  return s
}

func (s *L4DdosEventResponseData) SetEndTime(v string) *L4DdosEventResponseData {
  s.EndTime = &v
  return s
}

func (s *L4DdosEventResponseData) SetMaxTime(v string) *L4DdosEventResponseData {
  s.MaxTime = &v
  return s
}

func (s *L4DdosEventResponseData) SetStatus(v string) *L4DdosEventResponseData {
  s.Status = &v
  return s
}

func (s *L4DdosEventResponseData) SetMaxValue(v string) *L4DdosEventResponseData {
  s.MaxValue = &v
  return s
}

type L4DdosEventResponseHeader struct {
}

func (s L4DdosEventResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s L4DdosEventResponseHeader) GoString() string {
  return s.String()
}




type GetDistributionOfWAFAttackSourceRequest struct {
  // {"en":"Domain, array.", "zh_CN":"域名，数组。"}
  Domains []*string `json:"domains,omitempty" xml:"domains,omitempty" require:"true" type:"Repeated"`
  // {"en":"Start time, yyyy-MM-dd HH:mm:ss.", "zh_CN":"开始时间，yyyy-MM-dd HH:mm:ss。"}
  StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty" require:"true"`
  // {"en":"End time, yyyy-MM-dd HH:mm:ss.", "zh_CN":"结束时间，yyyy-MM-dd HH:mm:ss。"}
  EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty" require:"true"`
  // {"en":"Time zone, GMT+8 by default.", "zh_CN":"时区，默认GMT+8，即“GMT+8”。"}
  TimeZone *string `json:"timeZone,omitempty" xml:"timeZone,omitempty"`
  // {'en':'Policy type, array. [protocol: Protocol Validation,
  // webShell: Webshell Access Detection,
  // other: Others Rules,
  // access: Access Control/Rate Limiting,
  // rule: WAF Rules]', 'zh_CN':'触发策略类型，数组。
  // [protocol：协议合规检测，
  // webShell：后门识别，
  // other：其他防护规则，
  // access：访问控制/限速，
  // rule：Web规则防护]'}
  Policys []*string `json:"policys,omitempty" xml:"policys,omitempty" type:"Repeated"`
  // {"en":"The language of response data, default value: cn.
  //     cn: Chinese
  //     en: English", "zh_CN":"返回内容的语言版本，默认值：cn。
  //     cn：中文
  //     en：英文"}
  Lang *string `json:"lang,omitempty" xml:"lang,omitempty"`
}

func (s GetDistributionOfWAFAttackSourceRequest) String() string {
  return tea.Prettify(s)
}

func (s GetDistributionOfWAFAttackSourceRequest) GoString() string {
  return s.String()
}

func (s *GetDistributionOfWAFAttackSourceRequest) SetDomains(v []*string) *GetDistributionOfWAFAttackSourceRequest {
  s.Domains = v
  return s
}

func (s *GetDistributionOfWAFAttackSourceRequest) SetStartTime(v string) *GetDistributionOfWAFAttackSourceRequest {
  s.StartTime = &v
  return s
}

func (s *GetDistributionOfWAFAttackSourceRequest) SetEndTime(v string) *GetDistributionOfWAFAttackSourceRequest {
  s.EndTime = &v
  return s
}

func (s *GetDistributionOfWAFAttackSourceRequest) SetTimeZone(v string) *GetDistributionOfWAFAttackSourceRequest {
  s.TimeZone = &v
  return s
}

func (s *GetDistributionOfWAFAttackSourceRequest) SetPolicys(v []*string) *GetDistributionOfWAFAttackSourceRequest {
  s.Policys = v
  return s
}

func (s *GetDistributionOfWAFAttackSourceRequest) SetLang(v string) *GetDistributionOfWAFAttackSourceRequest {
  s.Lang = &v
  return s
}

type GetDistributionOfWAFAttackSourceResponse struct {
  // {"en":"Return 0 means success, please see <Error code> to check other status code.", "zh_CN":"0状态码表示请求成功，其他状态码说明请参见《错误码》。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Error message or Success.", "zh_CN":"错误信息或Success。"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Return data.", "zh_CN":"返回值。"}
  Data []*GetDistributionOfWAFAttackSourceAttackArea `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s GetDistributionOfWAFAttackSourceResponse) String() string {
  return tea.Prettify(s)
}

func (s GetDistributionOfWAFAttackSourceResponse) GoString() string {
  return s.String()
}

func (s *GetDistributionOfWAFAttackSourceResponse) SetCode(v string) *GetDistributionOfWAFAttackSourceResponse {
  s.Code = &v
  return s
}

func (s *GetDistributionOfWAFAttackSourceResponse) SetMessage(v string) *GetDistributionOfWAFAttackSourceResponse {
  s.Message = &v
  return s
}

func (s *GetDistributionOfWAFAttackSourceResponse) SetData(v []*GetDistributionOfWAFAttackSourceAttackArea) *GetDistributionOfWAFAttackSourceResponse {
  s.Data = v
  return s
}

type GetDistributionOfWAFAttackSourceAttackArea struct {
  // {"en":"Country.", "zh_CN":"国家。"}
  Country *string `json:"country,omitempty" xml:"country,omitempty" require:"true"`
  // {"en":"Province.", "zh_CN":"省份。"}
  Province *string `json:"province,omitempty" xml:"province,omitempty" require:"true"`
  // {"en":"Attack requests.", "zh_CN":"攻击请求数。"}
  TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty" require:"true"`
}

func (s GetDistributionOfWAFAttackSourceAttackArea) String() string {
  return tea.Prettify(s)
}

func (s GetDistributionOfWAFAttackSourceAttackArea) GoString() string {
  return s.String()
}

func (s *GetDistributionOfWAFAttackSourceAttackArea) SetCountry(v string) *GetDistributionOfWAFAttackSourceAttackArea {
  s.Country = &v
  return s
}

func (s *GetDistributionOfWAFAttackSourceAttackArea) SetProvince(v string) *GetDistributionOfWAFAttackSourceAttackArea {
  s.Province = &v
  return s
}

func (s *GetDistributionOfWAFAttackSourceAttackArea) SetTotalCount(v int64) *GetDistributionOfWAFAttackSourceAttackArea {
  s.TotalCount = &v
  return s
}

type GetDistributionOfWAFAttackSourcePaths struct {
}

func (s GetDistributionOfWAFAttackSourcePaths) String() string {
  return tea.Prettify(s)
}

func (s GetDistributionOfWAFAttackSourcePaths) GoString() string {
  return s.String()
}

type GetDistributionOfWAFAttackSourceParameters struct {
}

func (s GetDistributionOfWAFAttackSourceParameters) String() string {
  return tea.Prettify(s)
}

func (s GetDistributionOfWAFAttackSourceParameters) GoString() string {
  return s.String()
}

type GetDistributionOfWAFAttackSourceRequestHeader struct {
}

func (s GetDistributionOfWAFAttackSourceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetDistributionOfWAFAttackSourceRequestHeader) GoString() string {
  return s.String()
}

type GetDistributionOfWAFAttackSourceResponseHeader struct {
}

func (s GetDistributionOfWAFAttackSourceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetDistributionOfWAFAttackSourceResponseHeader) GoString() string {
  return s.String()
}




type GetRuleStatusStatisticsRequest struct {
  // {'en':'Domain name.', 'zh_CN':'域名。'}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {"en":"The language of response data, default value: cn.
  //     cn: Chinese
  //     en: English", "zh_CN":"返回内容的语言版本，默认值：cn。
  //     cn：中文
  //     en：英文"}
  Lang *string `json:"lang,omitempty" xml:"lang,omitempty"`
}

func (s GetRuleStatusStatisticsRequest) String() string {
  return tea.Prettify(s)
}

func (s GetRuleStatusStatisticsRequest) GoString() string {
  return s.String()
}

func (s *GetRuleStatusStatisticsRequest) SetDomain(v string) *GetRuleStatusStatisticsRequest {
  s.Domain = &v
  return s
}

func (s *GetRuleStatusStatisticsRequest) SetLang(v string) *GetRuleStatusStatisticsRequest {
  s.Lang = &v
  return s
}

type GetRuleStatusStatisticsResponse struct {
  // {"en":"Return 0 means success, please see <Error code> to check other status code.", "zh_CN":"0状态码表示请求成功，其他状态码说明请参见《错误码》。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"error message or Success.", "zh_CN":"错误信息或Success。"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {'en':'return data.','zh_CN':'返回值。'}
  Data *GetRuleStatusStatisticsStatisticalRuleStatus `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s GetRuleStatusStatisticsResponse) String() string {
  return tea.Prettify(s)
}

func (s GetRuleStatusStatisticsResponse) GoString() string {
  return s.String()
}

func (s *GetRuleStatusStatisticsResponse) SetCode(v string) *GetRuleStatusStatisticsResponse {
  s.Code = &v
  return s
}

func (s *GetRuleStatusStatisticsResponse) SetMessage(v string) *GetRuleStatusStatisticsResponse {
  s.Message = &v
  return s
}

func (s *GetRuleStatusStatisticsResponse) SetData(v *GetRuleStatusStatisticsStatisticalRuleStatus) *GetRuleStatusStatisticsResponse {
  s.Data = v
  return s
}

type GetRuleStatusStatisticsStatisticalRuleStatus struct {
  // {'en':'Attack type name.', 'zh_CN':'攻击类型名称。'}
  AttackTypeName *string `json:"attackTypeName,omitempty" xml:"attackTypeName,omitempty" require:"true"`
  // {'en':'Attack type.', 'zh_CN':'攻击类型。'}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
  // {'en':'The number of rules that enable monitoring for this domain name.', 'zh_CN':'该域名开启监控的规则数。'}
  Log *int `json:"log,omitempty" xml:"log,omitempty" require:"true"`
  // {'en':'The number of blocked rules for this domain name.', 'zh_CN':'该域名开启拦截的规则数。'}
  Block *int `json:"block,omitempty" xml:"block,omitempty" require:"true"`
  // {'en':'The number of rules closed for this domain name.', 'zh_CN':'该域名关闭的规则数。'}
  Off *int `json:"off,omitempty" xml:"off,omitempty" require:"true"`
}

func (s GetRuleStatusStatisticsStatisticalRuleStatus) String() string {
  return tea.Prettify(s)
}

func (s GetRuleStatusStatisticsStatisticalRuleStatus) GoString() string {
  return s.String()
}

func (s *GetRuleStatusStatisticsStatisticalRuleStatus) SetAttackTypeName(v string) *GetRuleStatusStatisticsStatisticalRuleStatus {
  s.AttackTypeName = &v
  return s
}

func (s *GetRuleStatusStatisticsStatisticalRuleStatus) SetType(v string) *GetRuleStatusStatisticsStatisticalRuleStatus {
  s.Type = &v
  return s
}

func (s *GetRuleStatusStatisticsStatisticalRuleStatus) SetLog(v int) *GetRuleStatusStatisticsStatisticalRuleStatus {
  s.Log = &v
  return s
}

func (s *GetRuleStatusStatisticsStatisticalRuleStatus) SetBlock(v int) *GetRuleStatusStatisticsStatisticalRuleStatus {
  s.Block = &v
  return s
}

func (s *GetRuleStatusStatisticsStatisticalRuleStatus) SetOff(v int) *GetRuleStatusStatisticsStatisticalRuleStatus {
  s.Off = &v
  return s
}

type GetRuleStatusStatisticsPaths struct {
}

func (s GetRuleStatusStatisticsPaths) String() string {
  return tea.Prettify(s)
}

func (s GetRuleStatusStatisticsPaths) GoString() string {
  return s.String()
}

type GetRuleStatusStatisticsParameters struct {
}

func (s GetRuleStatusStatisticsParameters) String() string {
  return tea.Prettify(s)
}

func (s GetRuleStatusStatisticsParameters) GoString() string {
  return s.String()
}

type GetRuleStatusStatisticsRequestHeader struct {
}

func (s GetRuleStatusStatisticsRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetRuleStatusStatisticsRequestHeader) GoString() string {
  return s.String()
}

type GetRuleStatusStatisticsResponseHeader struct {
}

func (s GetRuleStatusStatisticsResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetRuleStatusStatisticsResponseHeader) GoString() string {
  return s.String()
}




type GetApiNumberRequest struct {
  // {"en":"List of API group.", "zh_CN":"API分组列表。"}
  ApiGroups []*string `json:"apiGroups,omitempty" xml:"apiGroups,omitempty" type:"Repeated"`
  // {"en":"List of API name.", "zh_CN":"API名称列表。"}
  ApiIds []*string `json:"apiIds,omitempty" xml:"apiIds,omitempty" type:"Repeated"`
  // {"en":"List of domain.", "zh_CN":"域名列表。"}
  Domains []*string `json:"domains,omitempty" xml:"domains,omitempty" type:"Repeated"`
  // {"en":"Start time, format: yyyy-MM-dd HH:mm:ss.", "zh_CN":"开始时间，格式：yyyy-MM-dd HH:mm:ss。"}
  StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty" require:"true"`
  // {"en":"End time, format: yyyy-MM-dd HH:mm:ss.", "zh_CN":"结束时间，格式：yyyy-MM-dd HH:mm:ss。"}
  EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty" require:"true"`
  // {"en":"Endpoint path.", "zh_CN":"前端路径。"}
  FrontPath *string `json:"frontPath,omitempty" xml:"frontPath,omitempty"`
}

func (s GetApiNumberRequest) String() string {
  return tea.Prettify(s)
}

func (s GetApiNumberRequest) GoString() string {
  return s.String()
}

func (s *GetApiNumberRequest) SetApiGroups(v []*string) *GetApiNumberRequest {
  s.ApiGroups = v
  return s
}

func (s *GetApiNumberRequest) SetApiIds(v []*string) *GetApiNumberRequest {
  s.ApiIds = v
  return s
}

func (s *GetApiNumberRequest) SetDomains(v []*string) *GetApiNumberRequest {
  s.Domains = v
  return s
}

func (s *GetApiNumberRequest) SetStartTime(v string) *GetApiNumberRequest {
  s.StartTime = &v
  return s
}

func (s *GetApiNumberRequest) SetEndTime(v string) *GetApiNumberRequest {
  s.EndTime = &v
  return s
}

func (s *GetApiNumberRequest) SetFrontPath(v string) *GetApiNumberRequest {
  s.FrontPath = &v
  return s
}

type GetApiNumberResponse struct {
  // {"en":"Return 200 means success.", "zh_CN":"200状态码表示请求成功。"}
  Code *int `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Message.", "zh_CN":"返回信息。"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"Data.", "zh_CN":"数据。"}
  Data *GetApiNumberVo `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s GetApiNumberResponse) String() string {
  return tea.Prettify(s)
}

func (s GetApiNumberResponse) GoString() string {
  return s.String()
}

func (s *GetApiNumberResponse) SetCode(v int) *GetApiNumberResponse {
  s.Code = &v
  return s
}

func (s *GetApiNumberResponse) SetMsg(v string) *GetApiNumberResponse {
  s.Msg = &v
  return s
}

func (s *GetApiNumberResponse) SetData(v *GetApiNumberVo) *GetApiNumberResponse {
  s.Data = v
  return s
}

type GetApiNumberVo struct {
  // {"en":"Activate.", "zh_CN":"激活。"}
  Active *int `json:"active,omitempty" xml:"active,omitempty" require:"true"`
  // {"en":"To be activated.", "zh_CN":"待激活。"}
  Inactive *int `json:"inactive,omitempty" xml:"inactive,omitempty" require:"true"`
  // {"en":"Deactivate.", "zh_CN":"停用。"}
  Stop *int `json:"stop,omitempty" xml:"stop,omitempty" require:"true"`
}

func (s GetApiNumberVo) String() string {
  return tea.Prettify(s)
}

func (s GetApiNumberVo) GoString() string {
  return s.String()
}

func (s *GetApiNumberVo) SetActive(v int) *GetApiNumberVo {
  s.Active = &v
  return s
}

func (s *GetApiNumberVo) SetInactive(v int) *GetApiNumberVo {
  s.Inactive = &v
  return s
}

func (s *GetApiNumberVo) SetStop(v int) *GetApiNumberVo {
  s.Stop = &v
  return s
}

type GetApiNumberPaths struct {
}

func (s GetApiNumberPaths) String() string {
  return tea.Prettify(s)
}

func (s GetApiNumberPaths) GoString() string {
  return s.String()
}

type GetApiNumberParameters struct {
}

func (s GetApiNumberParameters) String() string {
  return tea.Prettify(s)
}

func (s GetApiNumberParameters) GoString() string {
  return s.String()
}

type GetApiNumberRequestHeader struct {
}

func (s GetApiNumberRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetApiNumberRequestHeader) GoString() string {
  return s.String()
}

type GetApiNumberResponseHeader struct {
}

func (s GetApiNumberResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetApiNumberResponseHeader) GoString() string {
  return s.String()
}




type GetTotalRequestNumberRequest struct {
  // {"en":"List of API group.", "zh_CN":"API分组列表。"}
  ApiGroups []*string `json:"apiGroups,omitempty" xml:"apiGroups,omitempty" type:"Repeated"`
  // {"en":"List of API name.", "zh_CN":"API名称列表。"}
  ApiIds []*string `json:"apiIds,omitempty" xml:"apiIds,omitempty" type:"Repeated"`
  // {"en":"List of domain.", "zh_CN":"域名列表。"}
  Domains []*string `json:"domains,omitempty" xml:"domains,omitempty" type:"Repeated"`
  // {"en":"Start time, format: yyyy-MM-dd HH:mm:ss.", "zh_CN":"开始时间，格式：yyyy-MM-dd HH:mm:ss。"}
  StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty" require:"true"`
  // {"en":"End time, format: yyyy-MM-dd HH:mm:ss.", "zh_CN":"结束时间，格式：yyyy-MM-dd HH:mm:ss。"}
  EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty" require:"true"`
  // {"en":"Endpoint path.", "zh_CN":"前端路径。"}
  FrontPath *string `json:"frontPath,omitempty" xml:"frontPath,omitempty"`
}

func (s GetTotalRequestNumberRequest) String() string {
  return tea.Prettify(s)
}

func (s GetTotalRequestNumberRequest) GoString() string {
  return s.String()
}

func (s *GetTotalRequestNumberRequest) SetApiGroups(v []*string) *GetTotalRequestNumberRequest {
  s.ApiGroups = v
  return s
}

func (s *GetTotalRequestNumberRequest) SetApiIds(v []*string) *GetTotalRequestNumberRequest {
  s.ApiIds = v
  return s
}

func (s *GetTotalRequestNumberRequest) SetDomains(v []*string) *GetTotalRequestNumberRequest {
  s.Domains = v
  return s
}

func (s *GetTotalRequestNumberRequest) SetStartTime(v string) *GetTotalRequestNumberRequest {
  s.StartTime = &v
  return s
}

func (s *GetTotalRequestNumberRequest) SetEndTime(v string) *GetTotalRequestNumberRequest {
  s.EndTime = &v
  return s
}

func (s *GetTotalRequestNumberRequest) SetFrontPath(v string) *GetTotalRequestNumberRequest {
  s.FrontPath = &v
  return s
}

type GetTotalRequestNumberResponse struct {
  // {"en":"Return 200 means success.", "zh_CN":"200状态码表示请求成功。"}
  Code *int `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Message.", "zh_CN":"返回信息。"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"Data.", "zh_CN":"数据。"}
  Data *GetTotalRequestNumberVo `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s GetTotalRequestNumberResponse) String() string {
  return tea.Prettify(s)
}

func (s GetTotalRequestNumberResponse) GoString() string {
  return s.String()
}

func (s *GetTotalRequestNumberResponse) SetCode(v int) *GetTotalRequestNumberResponse {
  s.Code = &v
  return s
}

func (s *GetTotalRequestNumberResponse) SetMsg(v string) *GetTotalRequestNumberResponse {
  s.Msg = &v
  return s
}

func (s *GetTotalRequestNumberResponse) SetData(v *GetTotalRequestNumberVo) *GetTotalRequestNumberResponse {
  s.Data = v
  return s
}

type GetTotalRequestNumberVo struct {
  // {"en":"The total number of requests.", "zh_CN":"总请求数。"}
  Total *int `json:"total,omitempty" xml:"total,omitempty" require:"true"`
  // {"en":"The number of API requests detected and blocked by API Shield.", "zh_CN":"缓解次数。"}
  Attack *int `json:"attack,omitempty" xml:"attack,omitempty" require:"true"`
}

func (s GetTotalRequestNumberVo) String() string {
  return tea.Prettify(s)
}

func (s GetTotalRequestNumberVo) GoString() string {
  return s.String()
}

func (s *GetTotalRequestNumberVo) SetTotal(v int) *GetTotalRequestNumberVo {
  s.Total = &v
  return s
}

func (s *GetTotalRequestNumberVo) SetAttack(v int) *GetTotalRequestNumberVo {
  s.Attack = &v
  return s
}

type GetTotalRequestNumberPaths struct {
}

func (s GetTotalRequestNumberPaths) String() string {
  return tea.Prettify(s)
}

func (s GetTotalRequestNumberPaths) GoString() string {
  return s.String()
}

type GetTotalRequestNumberParameters struct {
}

func (s GetTotalRequestNumberParameters) String() string {
  return tea.Prettify(s)
}

func (s GetTotalRequestNumberParameters) GoString() string {
  return s.String()
}

type GetTotalRequestNumberRequestHeader struct {
}

func (s GetTotalRequestNumberRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetTotalRequestNumberRequestHeader) GoString() string {
  return s.String()
}

type GetTotalRequestNumberResponseHeader struct {
}

func (s GetTotalRequestNumberResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetTotalRequestNumberResponseHeader) GoString() string {
  return s.String()
}




type GetWAFAttackIPRequest struct {
  // {"en":"Domain, array.", "zh_CN":"域名，数组。"}
  Domains []*string `json:"domains,omitempty" xml:"domains,omitempty" require:"true" type:"Repeated"`
  // {"en":"Start time, yyyy-MM-dd HH:mm:ss.", "zh_CN":"开始时间，yyyy-MM-dd HH:mm:ss。"}
  StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty" require:"true"`
  // {"en":"End time, yyyy-MM-dd HH:mm:ss.", "zh_CN":"结束时间，yyyy-MM-dd HH:mm:ss。"}
  EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty" require:"true"`
  // {"en":"Time zone, GMT+8 by default.", "zh_CN":"时区，默认GMT+8，即“GMT+8”。"}
  TimeZone *string `json:"timeZone,omitempty" xml:"timeZone,omitempty"`
  // {'en':'Policy type, array. [protocol: Protocol Validation,
  // webShell: Webshell Access Detection,
  // other: Others Rules,
  // access: Access Control/Rate Limiting,
  // rule: WAF Rules]', 'zh_CN':'触发策略类型，数组。
  // [protocol：协议合规检测，
  // webShell：后门识别，
  // other：其他防护规则，
  // access：访问控制/限速，
  // rule：Web规则防护]'}
  Policys []*string `json:"policys,omitempty" xml:"policys,omitempty" type:"Repeated"`
  // {"en":"Attacker IP Top value, default Top 100, only support Top 1000.", "zh_CN":"攻击IP Top值，默认返回Top 100，最大只能查询 Top 1000。"}
  TopNum *int `json:"topNum,omitempty" xml:"topNum,omitempty"`
}

func (s GetWAFAttackIPRequest) String() string {
  return tea.Prettify(s)
}

func (s GetWAFAttackIPRequest) GoString() string {
  return s.String()
}

func (s *GetWAFAttackIPRequest) SetDomains(v []*string) *GetWAFAttackIPRequest {
  s.Domains = v
  return s
}

func (s *GetWAFAttackIPRequest) SetStartTime(v string) *GetWAFAttackIPRequest {
  s.StartTime = &v
  return s
}

func (s *GetWAFAttackIPRequest) SetEndTime(v string) *GetWAFAttackIPRequest {
  s.EndTime = &v
  return s
}

func (s *GetWAFAttackIPRequest) SetTimeZone(v string) *GetWAFAttackIPRequest {
  s.TimeZone = &v
  return s
}

func (s *GetWAFAttackIPRequest) SetPolicys(v []*string) *GetWAFAttackIPRequest {
  s.Policys = v
  return s
}

func (s *GetWAFAttackIPRequest) SetTopNum(v int) *GetWAFAttackIPRequest {
  s.TopNum = &v
  return s
}

type GetWAFAttackIPResponse struct {
  // {"en":"Return 0 means success, please see <Error code> to check other status code.", "zh_CN":"0状态码表示请求成功，其他状态码说明请参见《错误码》。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Error message or Success.", "zh_CN":"错误信息或Success。"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Return data.", "zh_CN":"返回值。"}
  Data []*GetWAFAttackIPAttackIp `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s GetWAFAttackIPResponse) String() string {
  return tea.Prettify(s)
}

func (s GetWAFAttackIPResponse) GoString() string {
  return s.String()
}

func (s *GetWAFAttackIPResponse) SetCode(v string) *GetWAFAttackIPResponse {
  s.Code = &v
  return s
}

func (s *GetWAFAttackIPResponse) SetMessage(v string) *GetWAFAttackIPResponse {
  s.Message = &v
  return s
}

func (s *GetWAFAttackIPResponse) SetData(v []*GetWAFAttackIPAttackIp) *GetWAFAttackIPResponse {
  s.Data = v
  return s
}

type GetWAFAttackIPAttackIp struct {
  // {"en":"Attacker IP.", "zh_CN":"攻击IP。"}
  Ip *string `json:"ip,omitempty" xml:"ip,omitempty" require:"true"`
  // {"en":"Block requests.", "zh_CN":"拦截请求数。"}
  BlockCount *int64 `json:"blockCount,omitempty" xml:"blockCount,omitempty" require:"true"`
  // {"en":"Attack requests.", "zh_CN":"攻击请求数。"}
  TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty" require:"true"`
  // {"en":"Log requests.", "zh_CN":"监控请求数。"}
  MonitorCount *int64 `json:"monitorCount,omitempty" xml:"monitorCount,omitempty" require:"true"`
  // {"en":"Country(English)", "zh_CN":"国家英文。"}
  CountryEn *string `json:"countryEn,omitempty" xml:"countryEn,omitempty" require:"true"`
  // {"en":"Country(Chinese)", "zh_CN":"国家中文。"}
  CountryCn *string `json:"countryCn,omitempty" xml:"countryCn,omitempty" require:"true"`
  // {"en":"Province(English)", "zh_CN":"省份英文。"}
  ProvinceEn *string `json:"provinceEn,omitempty" xml:"provinceEn,omitempty" require:"true"`
  // {"en":"Province(Chinese)", "zh_CN":"省份中文。"}
  ProvinceCn *string `json:"provinceCn,omitempty" xml:"provinceCn,omitempty" require:"true"`
  // {"en":"City(English)", "zh_CN":"城市英文。"}
  CityEn *string `json:"cityEn,omitempty" xml:"cityEn,omitempty" require:"true"`
  // {"en":"City(Chinese)", "zh_CN":"城市中文。"}
  CityCn *string `json:"cityCn,omitempty" xml:"cityCn,omitempty" require:"true"`
}

func (s GetWAFAttackIPAttackIp) String() string {
  return tea.Prettify(s)
}

func (s GetWAFAttackIPAttackIp) GoString() string {
  return s.String()
}

func (s *GetWAFAttackIPAttackIp) SetIp(v string) *GetWAFAttackIPAttackIp {
  s.Ip = &v
  return s
}

func (s *GetWAFAttackIPAttackIp) SetBlockCount(v int64) *GetWAFAttackIPAttackIp {
  s.BlockCount = &v
  return s
}

func (s *GetWAFAttackIPAttackIp) SetTotalCount(v int64) *GetWAFAttackIPAttackIp {
  s.TotalCount = &v
  return s
}

func (s *GetWAFAttackIPAttackIp) SetMonitorCount(v int64) *GetWAFAttackIPAttackIp {
  s.MonitorCount = &v
  return s
}

func (s *GetWAFAttackIPAttackIp) SetCountryEn(v string) *GetWAFAttackIPAttackIp {
  s.CountryEn = &v
  return s
}

func (s *GetWAFAttackIPAttackIp) SetCountryCn(v string) *GetWAFAttackIPAttackIp {
  s.CountryCn = &v
  return s
}

func (s *GetWAFAttackIPAttackIp) SetProvinceEn(v string) *GetWAFAttackIPAttackIp {
  s.ProvinceEn = &v
  return s
}

func (s *GetWAFAttackIPAttackIp) SetProvinceCn(v string) *GetWAFAttackIPAttackIp {
  s.ProvinceCn = &v
  return s
}

func (s *GetWAFAttackIPAttackIp) SetCityEn(v string) *GetWAFAttackIPAttackIp {
  s.CityEn = &v
  return s
}

func (s *GetWAFAttackIPAttackIp) SetCityCn(v string) *GetWAFAttackIPAttackIp {
  s.CityCn = &v
  return s
}

type GetWAFAttackIPPaths struct {
}

func (s GetWAFAttackIPPaths) String() string {
  return tea.Prettify(s)
}

func (s GetWAFAttackIPPaths) GoString() string {
  return s.String()
}

type GetWAFAttackIPParameters struct {
}

func (s GetWAFAttackIPParameters) String() string {
  return tea.Prettify(s)
}

func (s GetWAFAttackIPParameters) GoString() string {
  return s.String()
}

type GetWAFAttackIPRequestHeader struct {
}

func (s GetWAFAttackIPRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetWAFAttackIPRequestHeader) GoString() string {
  return s.String()
}

type GetWAFAttackIPResponseHeader struct {
}

func (s GetWAFAttackIPResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetWAFAttackIPResponseHeader) GoString() string {
  return s.String()
}




type GetTopAttackSourcesForGlobalRequest struct {
  // {"defaultValue":"10","en":"Top N rankings to retrieve, default: 10, maximum: 1000.","zh_CN":"取前几排名，默认：10，上限：1000。"}
  Top *int `json:"top,omitempty" xml:"top,omitempty"`
  // {"en":"Multiple selection. Handling results, default: display all results. mitigated: Number of mitigated requests. monitored: Number of observed requests.","zh_CN":"多选。处置结果，默认：展示所有结果。\nmitigated：已抵御请求数。\nmonitored：观察请求数。","exampleValue":"mitigated,monitored"}
  ActType []*string `json:"actType,omitempty" xml:"actType,omitempty" type:"Repeated"`
  // {"en":"Start time, format: yyyy-MM-dd HH:mm:ss. The time range should not exceed 31 days.","zh_CN":"起始时间，格式：yyyy-MM-dd HH:mm:ss。\n时间范围不超过31天。"}
  StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty" require:"true"`
  // {"en":"End time, format: yyyy-MM-dd HH:mm:ss. The time range should not exceed 31 days.","zh_CN":"截止时间，格式：yyyy-MM-dd HH:mm:ss。\n时间范围不超过31天。"}
  EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty" require:"true"`
  // {"en":"List of domains. If not specified, all domains under the account are queried.","zh_CN":"域名列表，未指定时查询账号下的所有域名。"}
  Domains []*string `json:"domains,omitempty" xml:"domains,omitempty" type:"Repeated"`
}

func (s GetTopAttackSourcesForGlobalRequest) String() string {
  return tea.Prettify(s)
}

func (s GetTopAttackSourcesForGlobalRequest) GoString() string {
  return s.String()
}

func (s *GetTopAttackSourcesForGlobalRequest) SetTop(v int) *GetTopAttackSourcesForGlobalRequest {
  s.Top = &v
  return s
}

func (s *GetTopAttackSourcesForGlobalRequest) SetActType(v []*string) *GetTopAttackSourcesForGlobalRequest {
  s.ActType = v
  return s
}

func (s *GetTopAttackSourcesForGlobalRequest) SetStartTime(v string) *GetTopAttackSourcesForGlobalRequest {
  s.StartTime = &v
  return s
}

func (s *GetTopAttackSourcesForGlobalRequest) SetEndTime(v string) *GetTopAttackSourcesForGlobalRequest {
  s.EndTime = &v
  return s
}

func (s *GetTopAttackSourcesForGlobalRequest) SetDomains(v []*string) *GetTopAttackSourcesForGlobalRequest {
  s.Domains = v
  return s
}

type GetTopAttackSourcesForGlobalRequestHeader struct {
  // {"defaultValue":"en","en":"The language of response data, default value: en.\nzh_CN: Chinese\nen: English","zh_CN":"返回内容的语言版本，默认值: en。\nzh_CN：中文\nen：英文","exampleValue":"zh_CN,en"}
  Language *string `json:"language,omitempty" xml:"language,omitempty"`
  // {"dictionary":"belong=WAAP-MS-Ext|dict=waap_serviceType","en":"Security service type. Please enter a specific service type, if you purchase multiple security services.","zh_CN":"安全服务类型。有使用多个不同的安全服务时，需要填写具体的服务类型。"}
  ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty"`
  // {"defaultValue":"GMT+8","en":"Report Data Timezone:1. Indicates the timezone for the report data. lt must be relative to GMT andspecified in the format GMT+n where -12<=n <= 12. For example,timezone=GMT-7 and timezone=GMT+8.2. If the parameter is unspecified,results will be in the GMT+8 timezone.","zh_CN":"报表数据时区：1.请传递GMT时区。示例格式：GMT+N（其中12<=N <= 12），示例：timezone=GMT-7或timezone=GMT+8.2.若参数为传递，将默认按GMT+8时区查询"}
  Timezone *string `json:"timezone,omitempty" xml:"timezone,omitempty"`
}

func (s GetTopAttackSourcesForGlobalRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetTopAttackSourcesForGlobalRequestHeader) GoString() string {
  return s.String()
}

func (s *GetTopAttackSourcesForGlobalRequestHeader) SetLanguage(v string) *GetTopAttackSourcesForGlobalRequestHeader {
  s.Language = &v
  return s
}

func (s *GetTopAttackSourcesForGlobalRequestHeader) SetServiceType(v string) *GetTopAttackSourcesForGlobalRequestHeader {
  s.ServiceType = &v
  return s
}

func (s *GetTopAttackSourcesForGlobalRequestHeader) SetTimezone(v string) *GetTopAttackSourcesForGlobalRequestHeader {
  s.Timezone = &v
  return s
}

type GetTopAttackSourcesForGlobalPaths struct {
}

func (s GetTopAttackSourcesForGlobalPaths) String() string {
  return tea.Prettify(s)
}

func (s GetTopAttackSourcesForGlobalPaths) GoString() string {
  return s.String()
}

type GetTopAttackSourcesForGlobalParameters struct {
}

func (s GetTopAttackSourcesForGlobalParameters) String() string {
  return tea.Prettify(s)
}

func (s GetTopAttackSourcesForGlobalParameters) GoString() string {
  return s.String()
}

type GetTopAttackSourcesForGlobalResponse struct {
  // {"dictionary":"belong=WAAP-MS-Ext|dict=waap_retCodeEnum","en":"Please refer to the error code for exceptions.","zh_CN":"请参照错误码。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Description.","zh_CN":"描述信息。"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"Data.","zh_CN":"出参数据。"}
  Data []*GetTopAttackSourcesForGlobalResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s GetTopAttackSourcesForGlobalResponse) String() string {
  return tea.Prettify(s)
}

func (s GetTopAttackSourcesForGlobalResponse) GoString() string {
  return s.String()
}

func (s *GetTopAttackSourcesForGlobalResponse) SetCode(v string) *GetTopAttackSourcesForGlobalResponse {
  s.Code = &v
  return s
}

func (s *GetTopAttackSourcesForGlobalResponse) SetMsg(v string) *GetTopAttackSourcesForGlobalResponse {
  s.Msg = &v
  return s
}

func (s *GetTopAttackSourcesForGlobalResponse) SetData(v []*GetTopAttackSourcesForGlobalResponseData) *GetTopAttackSourcesForGlobalResponse {
  s.Data = v
  return s
}

type GetTopAttackSourcesForGlobalResponseData struct     {
  // {"en":"Source country or area.","zh_CN":"来源国家或地区。"}
  Area *string `json:"area,omitempty" xml:"area,omitempty" require:"true"`
  // {"en":"Attack requests.","zh_CN":"攻击请求数。"}
  Value *int64 `json:"value,omitempty" xml:"value,omitempty" require:"true"`
}

func (s GetTopAttackSourcesForGlobalResponseData) String() string {
  return tea.Prettify(s)
}

func (s GetTopAttackSourcesForGlobalResponseData) GoString() string {
  return s.String()
}

func (s *GetTopAttackSourcesForGlobalResponseData) SetArea(v string) *GetTopAttackSourcesForGlobalResponseData {
  s.Area = &v
  return s
}

func (s *GetTopAttackSourcesForGlobalResponseData) SetValue(v int64) *GetTopAttackSourcesForGlobalResponseData {
  s.Value = &v
  return s
}

type GetTopAttackSourcesForGlobalResponseHeader struct {
}

func (s GetTopAttackSourcesForGlobalResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetTopAttackSourcesForGlobalResponseHeader) GoString() string {
  return s.String()
}




type GetApiEventLogsRequest struct {
  // {"en":"The action to execute when a rule is matched, array,default value:[\"1\",\"2\",\"3\"].
  //  1:Block
  //  2:Log
  //  3:Sign", "zh_CN":"触发规则时的处理动作，数组。默认值：[\"1\",\"2\",\"3\"]。
  // 1：拦截
  // 2：监控
  // 3：标记"}
  Acts []*string `json:"acts,omitempty" xml:"acts,omitempty" type:"Repeated"`
  // {"en":"API name.", "zh_CN":"API名称。"}
  ApiName *string `json:"apiName,omitempty" xml:"apiName,omitempty"`
  // {"en":"The logic of API name match conditions,default value: INCLUDE.
  //  EQUAL: Equal
  //  UNEQUAL: Not equal
  //  INCLUDE: Included
  //  EXCLUDE: Not included.", "zh_CN":"API名称参数条件。默认值：INCLUDE。
  // EQUAL：等于
  // UNEQUAL：不等于
  // INCLUDE：包含
  // EXCLUDE：不包含
  // "}
  ApiNameCondition *string `json:"apiNameCondition,omitempty" xml:"apiNameCondition,omitempty"`
  // {"en":"List of event type, array.
  // [API_ACB: Access control blacklist,
  // API_AUTH_FAILED: Authentication failed,
  // API_BODY_LIMIT: Noncompliance request body,
  // API_DEACTIVATE_CALL: Deactivate API call,
  // API_METHOD_LIMIT: Noncompliance request method,
  // API_PARAM_LIMIT: Noncompliance parameter,
  // API_QUOTA_LIMIT: Quota limit,
  // API_REQUEST_LIMIT: High concurrency Limit,
  // API_UNAUTH_OBJECT: Unauthorized object]", "zh_CN":"事件类型，数组。
  // [API_ACB：访问控制限制,
  // API_AUTH_FAILED：鉴权失败,
  // API_BODY_LIMIT：请求BODY违规,
  // API_DEACTIVATE_CALL：停用API调用,
  // API_METHOD_LIMIT：请求方法违规,
  // API_PARAM_LIMIT：请求参数违规,
  // API_QUOTA_LIMIT：配额封顶,
  // API_REQUEST_LIMIT：限流封顶,
  // API_UNAUTH_OBJECT：非授权调用]"}
  AttackTypes []*string `json:"attackTypes,omitempty" xml:"attackTypes,omitempty" type:"Repeated"`
  // {"en":"Client IP.", "zh_CN":"客户端IP。"}
  ClientIp *string `json:"clientIp,omitempty" xml:"clientIp,omitempty"`
  // {"en":"The logic of ip match conditions, default value:INCLUDE.
  //  EQUAL: Equal
  //  UNEQUAL:Not equal
  //  INCLUDE:Included
  //  EXCLUDE: Not included.", "zh_CN":"客户端IP参数条件。默认值：INCLUDE。
  // EQUAL：等于
  // UNEQUAL：不等于
  // INCLUDE：包含
  // EXCLUDE：不包含
  // "}
  ClientIpCondition *string `json:"clientIpCondition,omitempty" xml:"clientIpCondition,omitempty"`
  // {"en":"List of domain name.", "zh_CN":"域名列表。"}
  Domains []*string `json:"domains,omitempty" xml:"domains,omitempty" type:"Repeated"`
  // {"en":"End time, format: yyyy-MM-dd HH:mm:ss.", "zh_CN":"结束时间，格式：yyyy-MM-dd HH:mm:ss。"}
  EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty" require:"true"`
  // {"en":"Endpoint path.", "zh_CN":"前端路径。"}
  FrontPath *string `json:"frontPath,omitempty" xml:"frontPath,omitempty"`
  // {"en":"The logic of front path match conditions, default value:INCLUDE.
  //  EQUAL: Equal
  //  UNEQUAL: Not equal
  //  INCLUDE: Included
  //  EXCLUDE: Not included.", "zh_CN":"前端路径参数条件。默认值：INCLUDE。
  // EQUAL：等于
  // UNEQUAL：不等于
  // INCLUDE：包含
  // EXCLUDE：不包含
  // "}
  FrontPathCondition *string `json:"frontPathCondition,omitempty" xml:"frontPathCondition,omitempty"`
  // {"en":"IP location.", "zh_CN":"IP地理位置。"}
  IpLocation *string `json:"ipLocation,omitempty" xml:"ipLocation,omitempty"`
  // {"en":"Abnormal info.", "zh_CN":"异常说明。"}
  LogRemark *string `json:"logRemark,omitempty" xml:"logRemark,omitempty"`
  // {"en":"The current page number.", "zh_CN":"当前页码。"}
  PageNum *int `json:"pageNum,omitempty" xml:"pageNum,omitempty" require:"true"`
  // {"en":"The number of records per page.", "zh_CN":"每页日志条数。"}
  PageSize *int `json:"pageSize,omitempty" xml:"pageSize,omitempty" require:"true"`
  // {"en":"HTTP request header: Referer.", "zh_CN":"来源网址。"}
  Referer *string `json:"referer,omitempty" xml:"referer,omitempty"`
  // {"en":"The logic of referer match conditions, default value:INCLUDE.
  //  EQUAL: Equal
  //  UNEQUAL: Not equal
  //  INCLUDE: Included
  //  EXCLUDE: Not included.", "zh_CN":"Referer参数条件。默认值：INCLUDE。
  // EQUAL：等于
  // UNEQUAL：不等于
  // INCLUDE：包含
  // EXCLUDE：不包含
  // "}
  RefererCondition *string `json:"refererCondition,omitempty" xml:"refererCondition,omitempty"`
  // {"en":"Rule name.", "zh_CN":"规则名称。"}
  RuleName *string `json:"ruleName,omitempty" xml:"ruleName,omitempty"`
  // {"en":"Start time, format: yyyy-MM-dd HH:mm:ss.", "zh_CN":"开始时间，格式：yyyy-MM-dd HH:mm:ss。"}
  StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty" require:"true"`
  // {"en":"Status code.", "zh_CN":"状态码。"}
  StatusCode *string `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  // {"en":"The logic of status code match conditions, default value:INCLUDE.
  //  EQUAL: Equal
  //  UNEQUAL: Not equal
  //  INCLUDE: Included
  //  EXCLUDE: Not included.", "zh_CN":"状态码参数条件。默认值：INCLUDE。
  // EQUAL：等于
  // UNEQUAL：不等于
  // INCLUDE：包含
  // EXCLUDE：不包含
  // "}
  StatusCodeCondition *string `json:"statusCodeCondition,omitempty" xml:"statusCodeCondition,omitempty"`
  // {"en":"Rule type.", "zh_CN":"规则类型。"}
  TacticsType *string `json:"tacticsType,omitempty" xml:"tacticsType,omitempty"`
  // {"en":"HTTP request header: User-Agent.", "zh_CN":"User-Agent头。"}
  Ua *string `json:"ua,omitempty" xml:"ua,omitempty"`
  // {"en":"The logic of User-Agent match conditions, default value:INCLUDE.
  //  EQUAL: Equal
  //  UNEQUAL: Not equal
  //  INCLUDE: Included
  //  EXCLUDE: Not included.", "zh_CN":"UA参数条件。默认值：INCLUDE。
  // EQUAL：等于
  // UNEQUAL：不等于
  // INCLUDE：包含
  // EXCLUDE：不包含
  // "}
  UaConditon *string `json:"uaConditon,omitempty" xml:"uaConditon,omitempty"`
  // {"en":"Event ID.", "zh_CN":"事件ID。"}
  Uuid *string `json:"uuid,omitempty" xml:"uuid,omitempty"`
}

func (s GetApiEventLogsRequest) String() string {
  return tea.Prettify(s)
}

func (s GetApiEventLogsRequest) GoString() string {
  return s.String()
}

func (s *GetApiEventLogsRequest) SetActs(v []*string) *GetApiEventLogsRequest {
  s.Acts = v
  return s
}

func (s *GetApiEventLogsRequest) SetApiName(v string) *GetApiEventLogsRequest {
  s.ApiName = &v
  return s
}

func (s *GetApiEventLogsRequest) SetApiNameCondition(v string) *GetApiEventLogsRequest {
  s.ApiNameCondition = &v
  return s
}

func (s *GetApiEventLogsRequest) SetAttackTypes(v []*string) *GetApiEventLogsRequest {
  s.AttackTypes = v
  return s
}

func (s *GetApiEventLogsRequest) SetClientIp(v string) *GetApiEventLogsRequest {
  s.ClientIp = &v
  return s
}

func (s *GetApiEventLogsRequest) SetClientIpCondition(v string) *GetApiEventLogsRequest {
  s.ClientIpCondition = &v
  return s
}

func (s *GetApiEventLogsRequest) SetDomains(v []*string) *GetApiEventLogsRequest {
  s.Domains = v
  return s
}

func (s *GetApiEventLogsRequest) SetEndTime(v string) *GetApiEventLogsRequest {
  s.EndTime = &v
  return s
}

func (s *GetApiEventLogsRequest) SetFrontPath(v string) *GetApiEventLogsRequest {
  s.FrontPath = &v
  return s
}

func (s *GetApiEventLogsRequest) SetFrontPathCondition(v string) *GetApiEventLogsRequest {
  s.FrontPathCondition = &v
  return s
}

func (s *GetApiEventLogsRequest) SetIpLocation(v string) *GetApiEventLogsRequest {
  s.IpLocation = &v
  return s
}

func (s *GetApiEventLogsRequest) SetLogRemark(v string) *GetApiEventLogsRequest {
  s.LogRemark = &v
  return s
}

func (s *GetApiEventLogsRequest) SetPageNum(v int) *GetApiEventLogsRequest {
  s.PageNum = &v
  return s
}

func (s *GetApiEventLogsRequest) SetPageSize(v int) *GetApiEventLogsRequest {
  s.PageSize = &v
  return s
}

func (s *GetApiEventLogsRequest) SetReferer(v string) *GetApiEventLogsRequest {
  s.Referer = &v
  return s
}

func (s *GetApiEventLogsRequest) SetRefererCondition(v string) *GetApiEventLogsRequest {
  s.RefererCondition = &v
  return s
}

func (s *GetApiEventLogsRequest) SetRuleName(v string) *GetApiEventLogsRequest {
  s.RuleName = &v
  return s
}

func (s *GetApiEventLogsRequest) SetStartTime(v string) *GetApiEventLogsRequest {
  s.StartTime = &v
  return s
}

func (s *GetApiEventLogsRequest) SetStatusCode(v string) *GetApiEventLogsRequest {
  s.StatusCode = &v
  return s
}

func (s *GetApiEventLogsRequest) SetStatusCodeCondition(v string) *GetApiEventLogsRequest {
  s.StatusCodeCondition = &v
  return s
}

func (s *GetApiEventLogsRequest) SetTacticsType(v string) *GetApiEventLogsRequest {
  s.TacticsType = &v
  return s
}

func (s *GetApiEventLogsRequest) SetUa(v string) *GetApiEventLogsRequest {
  s.Ua = &v
  return s
}

func (s *GetApiEventLogsRequest) SetUaConditon(v string) *GetApiEventLogsRequest {
  s.UaConditon = &v
  return s
}

func (s *GetApiEventLogsRequest) SetUuid(v string) *GetApiEventLogsRequest {
  s.Uuid = &v
  return s
}

type GetApiEventLogsResponse struct {
  // {"en":"Return 200 means success.", "zh_CN":"200状态码表示请求成功。"}
  Code *int `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Message.", "zh_CN":"返回信息。"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"Data.", "zh_CN":"数据。"}
  Data *GetApiEventLogsPage `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s GetApiEventLogsResponse) String() string {
  return tea.Prettify(s)
}

func (s GetApiEventLogsResponse) GoString() string {
  return s.String()
}

func (s *GetApiEventLogsResponse) SetCode(v int) *GetApiEventLogsResponse {
  s.Code = &v
  return s
}

func (s *GetApiEventLogsResponse) SetMsg(v string) *GetApiEventLogsResponse {
  s.Msg = &v
  return s
}

func (s *GetApiEventLogsResponse) SetData(v *GetApiEventLogsPage) *GetApiEventLogsResponse {
  s.Data = v
  return s
}

type GetApiEventLogsPage struct {
  // {"en":"The current page number.", "zh_CN":"当前页码。"}
  PageNum *int `json:"pageNum,omitempty" xml:"pageNum,omitempty" require:"true"`
  // {"en":"The number of records per page.", "zh_CN":"每页日志条数。"}
  PageSize *int `json:"pageSize,omitempty" xml:"pageSize,omitempty" require:"true"`
  // {"en":"The total number of records.", "zh_CN":"总条数。"}
  Total *int `json:"total,omitempty" xml:"total,omitempty" require:"true"`
  // {"en":"List of event log.", "zh_CN":"事件日志列表。"}
  Records *GetApiEventLogsVo `json:"records,omitempty" xml:"records,omitempty" require:"true"`
}

func (s GetApiEventLogsPage) String() string {
  return tea.Prettify(s)
}

func (s GetApiEventLogsPage) GoString() string {
  return s.String()
}

func (s *GetApiEventLogsPage) SetPageNum(v int) *GetApiEventLogsPage {
  s.PageNum = &v
  return s
}

func (s *GetApiEventLogsPage) SetPageSize(v int) *GetApiEventLogsPage {
  s.PageSize = &v
  return s
}

func (s *GetApiEventLogsPage) SetTotal(v int) *GetApiEventLogsPage {
  s.Total = &v
  return s
}

func (s *GetApiEventLogsPage) SetRecords(v *GetApiEventLogsVo) *GetApiEventLogsPage {
  s.Records = v
  return s
}

type GetApiEventLogsVo struct {
  // {"en":"The action to execute when a rule is matched.
  //  1:Block
  //  2:Log
  //  3:Sign", "zh_CN":"触发规则时的处理动作。
  // 1：拦截
  // 2：监控
  // 3：标记"}
  Act *string `json:"act,omitempty" xml:"act,omitempty" require:"true"`
  // {"en":"The action description to execute when a rule is matched.", "zh_CN":"处理动作描述。"}
  ActDesc *string `json:"actDesc,omitempty" xml:"actDesc,omitempty" require:"true"`
  // {"en":"API name.", "zh_CN":"api名称。"}
  ApiName *string `json:"apiName,omitempty" xml:"apiName,omitempty" require:"true"`
  // {"en":"Event description.", "zh_CN":"事件描述。"}
  AttackDesc *string `json:"attackDesc,omitempty" xml:"attackDesc,omitempty" require:"true"`
  // {"en":"Event type.", "zh_CN":"事件类型。"}
  AttackType *string `json:"attackType,omitempty" xml:"attackType,omitempty" require:"true"`
  // {"en":"Client IP.", "zh_CN":"客户端IP。"}
  ClientIp *string `json:"clientIp,omitempty" xml:"clientIp,omitempty" require:"true"`
  // {"en":"Domain.", "zh_CN":"域名。"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {"en":"Abnormal info.", "zh_CN":"异常说明。"}
  ExceptionDesc *string `json:"exceptionDesc,omitempty" xml:"exceptionDesc,omitempty" require:"true"`
  // {"en":"Endpoint path.", "zh_CN":"前端路径。"}
  FrontPath *string `json:"frontPath,omitempty" xml:"frontPath,omitempty" require:"true"`
  // {"en":"HTTP version.", "zh_CN":"http版本。"}
  HttpVersion *string `json:"httpVersion,omitempty" xml:"httpVersion,omitempty" require:"true"`
  // {"en":"IP location.", "zh_CN":"IP 地理位置。"}
  IpLocation *string `json:"ipLocation,omitempty" xml:"ipLocation,omitempty" require:"true"`
  // {"en":"HTTP request header: Referer.", "zh_CN":"来源网址。"}
  Referer *string `json:"referer,omitempty" xml:"referer,omitempty" require:"true"`
  // {"en":"HTTP request header.", "zh_CN":"请求头信息。"}
  RequestHead *string `json:"requestHead,omitempty" xml:"requestHead,omitempty" require:"true"`
  // {"en":"HTTP request method.", "zh_CN":"请求方法。"}
  RequestMethod *string `json:"requestMethod,omitempty" xml:"requestMethod,omitempty" require:"true"`
  // {"en":"HTTP request URL.", "zh_CN":"请求URL。"}
  RequestUrl *string `json:"requestUrl,omitempty" xml:"requestUrl,omitempty" require:"true"`
  // {"en":"HTTP response header.", "zh_CN":"响应头信息。"}
  Response *string `json:"response,omitempty" xml:"response,omitempty" require:"true"`
  // {"en":"Rule name.", "zh_CN":"规则名称。"}
  RuleName *string `json:"ruleName,omitempty" xml:"ruleName,omitempty" require:"true"`
  // {"en":"Status code.", "zh_CN":"状态码。"}
  StatusCode *string `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
  // {"en":"Rule type.", "zh_CN":"规则类型。"}
  TacticsType *string `json:"tacticsType,omitempty" xml:"tacticsType,omitempty" require:"true"`
  // {"en":"Time, format: yyyy-MM-dd HH:mm:ss.", "zh_CN":"时间，格式：yyyy-MM-dd HH:mm:ss。"}
  Time *string `json:"time,omitempty" xml:"time,omitempty" require:"true"`
  // {"en":"HTTP request header: User-Agent.", "zh_CN":"User-Agent头。"}
  Ua *string `json:"ua,omitempty" xml:"ua,omitempty" require:"true"`
  // {"en":"Event ID.", "zh_CN":"事件ID。"}
  Uuid *string `json:"uuid,omitempty" xml:"uuid,omitempty" require:"true"`
}

func (s GetApiEventLogsVo) String() string {
  return tea.Prettify(s)
}

func (s GetApiEventLogsVo) GoString() string {
  return s.String()
}

func (s *GetApiEventLogsVo) SetAct(v string) *GetApiEventLogsVo {
  s.Act = &v
  return s
}

func (s *GetApiEventLogsVo) SetActDesc(v string) *GetApiEventLogsVo {
  s.ActDesc = &v
  return s
}

func (s *GetApiEventLogsVo) SetApiName(v string) *GetApiEventLogsVo {
  s.ApiName = &v
  return s
}

func (s *GetApiEventLogsVo) SetAttackDesc(v string) *GetApiEventLogsVo {
  s.AttackDesc = &v
  return s
}

func (s *GetApiEventLogsVo) SetAttackType(v string) *GetApiEventLogsVo {
  s.AttackType = &v
  return s
}

func (s *GetApiEventLogsVo) SetClientIp(v string) *GetApiEventLogsVo {
  s.ClientIp = &v
  return s
}

func (s *GetApiEventLogsVo) SetDomain(v string) *GetApiEventLogsVo {
  s.Domain = &v
  return s
}

func (s *GetApiEventLogsVo) SetExceptionDesc(v string) *GetApiEventLogsVo {
  s.ExceptionDesc = &v
  return s
}

func (s *GetApiEventLogsVo) SetFrontPath(v string) *GetApiEventLogsVo {
  s.FrontPath = &v
  return s
}

func (s *GetApiEventLogsVo) SetHttpVersion(v string) *GetApiEventLogsVo {
  s.HttpVersion = &v
  return s
}

func (s *GetApiEventLogsVo) SetIpLocation(v string) *GetApiEventLogsVo {
  s.IpLocation = &v
  return s
}

func (s *GetApiEventLogsVo) SetReferer(v string) *GetApiEventLogsVo {
  s.Referer = &v
  return s
}

func (s *GetApiEventLogsVo) SetRequestHead(v string) *GetApiEventLogsVo {
  s.RequestHead = &v
  return s
}

func (s *GetApiEventLogsVo) SetRequestMethod(v string) *GetApiEventLogsVo {
  s.RequestMethod = &v
  return s
}

func (s *GetApiEventLogsVo) SetRequestUrl(v string) *GetApiEventLogsVo {
  s.RequestUrl = &v
  return s
}

func (s *GetApiEventLogsVo) SetResponse(v string) *GetApiEventLogsVo {
  s.Response = &v
  return s
}

func (s *GetApiEventLogsVo) SetRuleName(v string) *GetApiEventLogsVo {
  s.RuleName = &v
  return s
}

func (s *GetApiEventLogsVo) SetStatusCode(v string) *GetApiEventLogsVo {
  s.StatusCode = &v
  return s
}

func (s *GetApiEventLogsVo) SetTacticsType(v string) *GetApiEventLogsVo {
  s.TacticsType = &v
  return s
}

func (s *GetApiEventLogsVo) SetTime(v string) *GetApiEventLogsVo {
  s.Time = &v
  return s
}

func (s *GetApiEventLogsVo) SetUa(v string) *GetApiEventLogsVo {
  s.Ua = &v
  return s
}

func (s *GetApiEventLogsVo) SetUuid(v string) *GetApiEventLogsVo {
  s.Uuid = &v
  return s
}

type GetApiEventLogsPaths struct {
}

func (s GetApiEventLogsPaths) String() string {
  return tea.Prettify(s)
}

func (s GetApiEventLogsPaths) GoString() string {
  return s.String()
}

type GetApiEventLogsParameters struct {
}

func (s GetApiEventLogsParameters) String() string {
  return tea.Prettify(s)
}

func (s GetApiEventLogsParameters) GoString() string {
  return s.String()
}

type GetApiEventLogsRequestHeader struct {
}

func (s GetApiEventLogsRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetApiEventLogsRequestHeader) GoString() string {
  return s.String()
}

type GetApiEventLogsResponseHeader struct {
}

func (s GetApiEventLogsResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetApiEventLogsResponseHeader) GoString() string {
  return s.String()
}




type GetWAFPolicyDetailsRequest struct {
  // {"en":"Domain, array.", "zh_CN":"域名，数组。"}
  Domains []*string `json:"domains,omitempty" xml:"domains,omitempty" require:"true" type:"Repeated"`
  // {"en":"Start time, yyyy-MM-dd HH:mm:ss.", "zh_CN":"开始时间，yyyy-MM-dd HH:mm:ss。"}
  StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty" require:"true"`
  // {"en":"End time, yyyy-MM-dd HH:mm:ss.", "zh_CN":"结束时间，yyyy-MM-dd HH:mm:ss。"}
  EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty" require:"true"`
  // {"en":"Rule ID.", "zh_CN":"规则ID。"}
  RuleId *string `json:"ruleId,omitempty" xml:"ruleId,omitempty" require:"true"`
  // {"en":"Policy type, array. [protocol: Protocol Validation,
  // webShell: Webshell Access Detection,
  // other: Others Rules,
  // access: Access Control/Rate Limiting,
  // rule: WAF Rules]", "zh_CN":"触发策略类型，数组。[protocol：协议合规检测，
  // webShell：后门识别，
  // other：其他防护规则，
  // access：访问控制/限速，
  // rule：Web规则防护]"}
  Policys []*string `json:"policys,omitempty" xml:"policys,omitempty" type:"Repeated"`
  // {"en":"Time zone, GMT+8 by default.", "zh_CN":"时区，默认GMT+8，即“GMT+8”。"}
  TimeZone *string `json:"timeZone,omitempty" xml:"timeZone,omitempty"`
  // {"en":"Action, default 1 and 2.
  //     1: Block
  //     2: Log", "zh_CN":"处理动作，默认1和2。
  //     1：拦截
  //     2：监控"}
  Acts []*string `json:"acts,omitempty" xml:"acts,omitempty" type:"Repeated"`
}

func (s GetWAFPolicyDetailsRequest) String() string {
  return tea.Prettify(s)
}

func (s GetWAFPolicyDetailsRequest) GoString() string {
  return s.String()
}

func (s *GetWAFPolicyDetailsRequest) SetDomains(v []*string) *GetWAFPolicyDetailsRequest {
  s.Domains = v
  return s
}

func (s *GetWAFPolicyDetailsRequest) SetStartTime(v string) *GetWAFPolicyDetailsRequest {
  s.StartTime = &v
  return s
}

func (s *GetWAFPolicyDetailsRequest) SetEndTime(v string) *GetWAFPolicyDetailsRequest {
  s.EndTime = &v
  return s
}

func (s *GetWAFPolicyDetailsRequest) SetRuleId(v string) *GetWAFPolicyDetailsRequest {
  s.RuleId = &v
  return s
}

func (s *GetWAFPolicyDetailsRequest) SetPolicys(v []*string) *GetWAFPolicyDetailsRequest {
  s.Policys = v
  return s
}

func (s *GetWAFPolicyDetailsRequest) SetTimeZone(v string) *GetWAFPolicyDetailsRequest {
  s.TimeZone = &v
  return s
}

func (s *GetWAFPolicyDetailsRequest) SetActs(v []*string) *GetWAFPolicyDetailsRequest {
  s.Acts = v
  return s
}

type GetWAFPolicyDetailsResponse struct {
  // {"en":"Return 0 means success, please see <Error code> to check other status code.", "zh_CN":"0状态码表示请求成功，其他状态码说明请参见《错误码》。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Error message or Success.", "zh_CN":"错误信息或Success。"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Return data.", "zh_CN":"返回值。"}
  Data *string `json:"data,omitempty" xml:"data,omitempty" require:"true"`
  // {"en":"Last attack time.", "zh_CN":"最近攻击时间。"}
  AttackTime *string `json:"attackTime,omitempty" xml:"attackTime,omitempty" require:"true"`
  // {"en":"Attack type.", "zh_CN":"攻击类型。"}
  AttackType *string `json:"attackType,omitempty" xml:"attackType,omitempty" require:"true"`
  // {"en":"Chinese name of attack type.", "zh_CN":"攻击类型中文名称。"}
  AttackTypeName *string `json:"attackTypeName,omitempty" xml:"attackTypeName,omitempty" require:"true"`
  // {"en":"English name of attack type.", "zh_CN":"攻击类型英文名称。"}
  AttackTypeNameEn *string `json:"attackTypeNameEn,omitempty" xml:"attackTypeNameEn,omitempty" require:"true"`
  // {"en":"Attack requests.", "zh_CN":"攻击次数。"}
  TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty" require:"true"`
  // {"en":"Rule ID.", "zh_CN":"规则ID。"}
  RuleId *string `json:"ruleId,omitempty" xml:"ruleId,omitempty" require:"true"`
  // {"en":"Rule name.", "zh_CN":"规则名称。"}
  RuleName *string `json:"ruleName,omitempty" xml:"ruleName,omitempty" require:"true"`
  // {"en":"Attacker IP.", "zh_CN":"攻击IP。"}
  Ips []*GetWAFPolicyDetailsAttackItemArray `json:"ips,omitempty" xml:"ips,omitempty" require:"true" type:"Repeated"`
  // {"en":"Attacked url.", "zh_CN":"受攻击URL。"}
  Urls []*GetWAFPolicyDetailsAttackUrlItemArray `json:"urls,omitempty" xml:"urls,omitempty" require:"true" type:"Repeated"`
}

func (s GetWAFPolicyDetailsResponse) String() string {
  return tea.Prettify(s)
}

func (s GetWAFPolicyDetailsResponse) GoString() string {
  return s.String()
}

func (s *GetWAFPolicyDetailsResponse) SetCode(v string) *GetWAFPolicyDetailsResponse {
  s.Code = &v
  return s
}

func (s *GetWAFPolicyDetailsResponse) SetMessage(v string) *GetWAFPolicyDetailsResponse {
  s.Message = &v
  return s
}

func (s *GetWAFPolicyDetailsResponse) SetData(v string) *GetWAFPolicyDetailsResponse {
  s.Data = &v
  return s
}

func (s *GetWAFPolicyDetailsResponse) SetAttackTime(v string) *GetWAFPolicyDetailsResponse {
  s.AttackTime = &v
  return s
}

func (s *GetWAFPolicyDetailsResponse) SetAttackType(v string) *GetWAFPolicyDetailsResponse {
  s.AttackType = &v
  return s
}

func (s *GetWAFPolicyDetailsResponse) SetAttackTypeName(v string) *GetWAFPolicyDetailsResponse {
  s.AttackTypeName = &v
  return s
}

func (s *GetWAFPolicyDetailsResponse) SetAttackTypeNameEn(v string) *GetWAFPolicyDetailsResponse {
  s.AttackTypeNameEn = &v
  return s
}

func (s *GetWAFPolicyDetailsResponse) SetTotalCount(v int64) *GetWAFPolicyDetailsResponse {
  s.TotalCount = &v
  return s
}

func (s *GetWAFPolicyDetailsResponse) SetRuleId(v string) *GetWAFPolicyDetailsResponse {
  s.RuleId = &v
  return s
}

func (s *GetWAFPolicyDetailsResponse) SetRuleName(v string) *GetWAFPolicyDetailsResponse {
  s.RuleName = &v
  return s
}

func (s *GetWAFPolicyDetailsResponse) SetIps(v []*GetWAFPolicyDetailsAttackItemArray) *GetWAFPolicyDetailsResponse {
  s.Ips = v
  return s
}

func (s *GetWAFPolicyDetailsResponse) SetUrls(v []*GetWAFPolicyDetailsAttackUrlItemArray) *GetWAFPolicyDetailsResponse {
  s.Urls = v
  return s
}

type GetWAFPolicyDetailsAttackItemArray struct {
  // {"en":"Attacker IP count", "zh_CN":"攻击IP记录数。"}
  Count *int64 `json:"count,omitempty" xml:"count,omitempty" require:"true"`
  // {"en":"Attacker IP array", "zh_CN":"攻击IP数组。"}
  List []*GetWAFPolicyDetailsAttackItem `json:"list,omitempty" xml:"list,omitempty" require:"true" type:"Repeated"`
}

func (s GetWAFPolicyDetailsAttackItemArray) String() string {
  return tea.Prettify(s)
}

func (s GetWAFPolicyDetailsAttackItemArray) GoString() string {
  return s.String()
}

func (s *GetWAFPolicyDetailsAttackItemArray) SetCount(v int64) *GetWAFPolicyDetailsAttackItemArray {
  s.Count = &v
  return s
}

func (s *GetWAFPolicyDetailsAttackItemArray) SetList(v []*GetWAFPolicyDetailsAttackItem) *GetWAFPolicyDetailsAttackItemArray {
  s.List = v
  return s
}

type GetWAFPolicyDetailsAttackItem struct {
  // {"en":"Attacker IP", "zh_CN":"攻击IP。"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"Last attack time.", "zh_CN":"最近攻击时间。"}
  AttackTime *string `json:"attackTime,omitempty" xml:"attackTime,omitempty" require:"true"`
  // {"en":"Attack requests.", "zh_CN":"攻击次数。"}
  TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty" require:"true"`
}

func (s GetWAFPolicyDetailsAttackItem) String() string {
  return tea.Prettify(s)
}

func (s GetWAFPolicyDetailsAttackItem) GoString() string {
  return s.String()
}

func (s *GetWAFPolicyDetailsAttackItem) SetName(v string) *GetWAFPolicyDetailsAttackItem {
  s.Name = &v
  return s
}

func (s *GetWAFPolicyDetailsAttackItem) SetAttackTime(v string) *GetWAFPolicyDetailsAttackItem {
  s.AttackTime = &v
  return s
}

func (s *GetWAFPolicyDetailsAttackItem) SetTotalCount(v int64) *GetWAFPolicyDetailsAttackItem {
  s.TotalCount = &v
  return s
}

type GetWAFPolicyDetailsAttackUrlItemArray struct {
  // {"en":"Attack requests count", "zh_CN":"受攻击URL记录数。"}
  Count *int64 `json:"count,omitempty" xml:"count,omitempty" require:"true"`
  // {"en":"Attack requests array", "zh_CN":"受攻击URL数组。"}
  List []*GetWAFPolicyDetailsAttackUrlItem `json:"list,omitempty" xml:"list,omitempty" require:"true" type:"Repeated"`
}

func (s GetWAFPolicyDetailsAttackUrlItemArray) String() string {
  return tea.Prettify(s)
}

func (s GetWAFPolicyDetailsAttackUrlItemArray) GoString() string {
  return s.String()
}

func (s *GetWAFPolicyDetailsAttackUrlItemArray) SetCount(v int64) *GetWAFPolicyDetailsAttackUrlItemArray {
  s.Count = &v
  return s
}

func (s *GetWAFPolicyDetailsAttackUrlItemArray) SetList(v []*GetWAFPolicyDetailsAttackUrlItem) *GetWAFPolicyDetailsAttackUrlItemArray {
  s.List = v
  return s
}

type GetWAFPolicyDetailsAttackUrlItem struct {
  // {"en":"URL", "zh_CN":"URL"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"Last attack time.", "zh_CN":"最近防护时间。"}
  AttackTime *string `json:"attackTime,omitempty" xml:"attackTime,omitempty" require:"true"`
  // {"en":"Attack requests.", "zh_CN":"攻击次数。"}
  TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty" require:"true"`
}

func (s GetWAFPolicyDetailsAttackUrlItem) String() string {
  return tea.Prettify(s)
}

func (s GetWAFPolicyDetailsAttackUrlItem) GoString() string {
  return s.String()
}

func (s *GetWAFPolicyDetailsAttackUrlItem) SetName(v string) *GetWAFPolicyDetailsAttackUrlItem {
  s.Name = &v
  return s
}

func (s *GetWAFPolicyDetailsAttackUrlItem) SetAttackTime(v string) *GetWAFPolicyDetailsAttackUrlItem {
  s.AttackTime = &v
  return s
}

func (s *GetWAFPolicyDetailsAttackUrlItem) SetTotalCount(v int64) *GetWAFPolicyDetailsAttackUrlItem {
  s.TotalCount = &v
  return s
}

type GetWAFPolicyDetailsPaths struct {
}

func (s GetWAFPolicyDetailsPaths) String() string {
  return tea.Prettify(s)
}

func (s GetWAFPolicyDetailsPaths) GoString() string {
  return s.String()
}

type GetWAFPolicyDetailsParameters struct {
}

func (s GetWAFPolicyDetailsParameters) String() string {
  return tea.Prettify(s)
}

func (s GetWAFPolicyDetailsParameters) GoString() string {
  return s.String()
}

type GetWAFPolicyDetailsRequestHeader struct {
}

func (s GetWAFPolicyDetailsRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetWAFPolicyDetailsRequestHeader) GoString() string {
  return s.String()
}

type GetWAFPolicyDetailsResponseHeader struct {
}

func (s GetWAFPolicyDetailsResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetWAFPolicyDetailsResponseHeader) GoString() string {
  return s.String()
}




type GetTrendsByrpsRequest struct {
  // {"en":"Start time, format: yyyy-MM-dd HH:mm:ss.\nThe time range can not exceed 31 days.","zh_CN":"开始时间，格式：yyyy-MM-dd HH:mm:ss。\n时间范围不超过31天。"}
  StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty" require:"true"`
  // {"en":"End time, format: yyyy-MM-dd HH:mm:ss.\nThe time range can not exceed 31 days.","zh_CN":"结束时间，格式：yyyy-MM-dd HH:mm:ss。\n时间范围不超过31天。"}
  EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty" require:"true"`
  // {"en":"Multiple selection. Handling results, default: display all results. \nmitigated: Number of mitigated requests. \nmonitored: Number of observed requests.","zh_CN":"多选。处置结果，默认：展示所有结果。\nmitigated：已抵御请求数。\nmonitored：观察请求数。","exampleValue":"mitigated,monitored"}
  ActType []*string `json:"actType,omitempty" xml:"actType,omitempty" type:"Repeated"`
  // {"en":"Domain list. Queries all domains under the account when not specified.","zh_CN":"域名列表，未指定时查询账号下的所有域名。"}
  Domains []*string `json:"domains,omitempty" xml:"domains,omitempty" type:"Repeated"`
}

func (s GetTrendsByrpsRequest) String() string {
  return tea.Prettify(s)
}

func (s GetTrendsByrpsRequest) GoString() string {
  return s.String()
}

func (s *GetTrendsByrpsRequest) SetStartTime(v string) *GetTrendsByrpsRequest {
  s.StartTime = &v
  return s
}

func (s *GetTrendsByrpsRequest) SetEndTime(v string) *GetTrendsByrpsRequest {
  s.EndTime = &v
  return s
}

func (s *GetTrendsByrpsRequest) SetActType(v []*string) *GetTrendsByrpsRequest {
  s.ActType = v
  return s
}

func (s *GetTrendsByrpsRequest) SetDomains(v []*string) *GetTrendsByrpsRequest {
  s.Domains = v
  return s
}

type GetTrendsByrpsRequestHeader struct {
  // {"dictionary":"belong=WAAP-MS-Ext|dict=waap_serviceType","en":"Security service type. Please enter a specific service type, if you purchase multiple security services.","zh_CN":"安全服务类型。有使用多个不同的安全服务时，需要填写具体的服务类型。"}
  ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty"`
  // {"defaultValue":"GMT+8","en":"Report Data Timezone:1. Indicates the timezone for the report data. lt must be relative to GMT andspecified in the format GMT+n where -12<=n <= 12. For example,timezone=GMT-7 and timezone=GMT+8.2. If the parameter is unspecified,results will be in the GMT+8 timezone.","zh_CN":"报表数据时区：1.请传递GMT时区。示例格式：GMT+N（其中12<=N <= 12），示例：timezone=GMT-7或timezone=GMT+8.2.若参数为传递，将默认按GMT+8时区查询"}
  Timezone *string `json:"timezone,omitempty" xml:"timezone,omitempty"`
}

func (s GetTrendsByrpsRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetTrendsByrpsRequestHeader) GoString() string {
  return s.String()
}

func (s *GetTrendsByrpsRequestHeader) SetServiceType(v string) *GetTrendsByrpsRequestHeader {
  s.ServiceType = &v
  return s
}

func (s *GetTrendsByrpsRequestHeader) SetTimezone(v string) *GetTrendsByrpsRequestHeader {
  s.Timezone = &v
  return s
}

type GetTrendsByrpsPaths struct {
}

func (s GetTrendsByrpsPaths) String() string {
  return tea.Prettify(s)
}

func (s GetTrendsByrpsPaths) GoString() string {
  return s.String()
}

type GetTrendsByrpsParameters struct {
}

func (s GetTrendsByrpsParameters) String() string {
  return tea.Prettify(s)
}

func (s GetTrendsByrpsParameters) GoString() string {
  return s.String()
}

type GetTrendsByrpsResponse struct {
  // {"dictionary":"belong=WAAP-MS-Ext|dict=waap_retCodeEnum","en":"Please refer to the error code for exceptions.","zh_CN":"请参照错误码。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Description.","zh_CN":"描述信息。"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"Data.","zh_CN":"出参数据。"}
  Data []*GetTrendsByrpsResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s GetTrendsByrpsResponse) String() string {
  return tea.Prettify(s)
}

func (s GetTrendsByrpsResponse) GoString() string {
  return s.String()
}

func (s *GetTrendsByrpsResponse) SetCode(v string) *GetTrendsByrpsResponse {
  s.Code = &v
  return s
}

func (s *GetTrendsByrpsResponse) SetMsg(v string) *GetTrendsByrpsResponse {
  s.Msg = &v
  return s
}

func (s *GetTrendsByrpsResponse) SetData(v []*GetTrendsByrpsResponseData) *GetTrendsByrpsResponse {
  s.Data = v
  return s
}

type GetTrendsByrpsResponseData struct     {
  // {"en":"Time point(yyyy-MM-dd HH-mm-ss).","zh_CN":"时间点（yyyy-MM-dd HH-mm-ss）。"}
  TimePoint *string `json:"timePoint,omitempty" xml:"timePoint,omitempty" require:"true"`
  // {"en":"Total requests(per second).","zh_CN":"总请求数（每秒均值）。"}
  Total *int64 `json:"total,omitempty" xml:"total,omitempty" require:"true"`
  // {"en":"Attack requests(per second).","zh_CN":"攻击请求数（每秒均值）。"}
  Attack *int64 `json:"attack,omitempty" xml:"attack,omitempty" require:"true"`
  // {"en":"Mitigated requests.","zh_CN":"已抵御请求数。"}
  Mitigated *int64 `json:"mitigated,omitempty" xml:"mitigated,omitempty" require:"true"`
  // {"en":"Monitored requests.","zh_CN":"观察请求数。"}
  Monitored *int64 `json:"monitored,omitempty" xml:"monitored,omitempty" require:"true"`
  // {"en":"Whitelist requests(per second).","zh_CN":"白名单请求数（每秒均值）。"}
  Whitelist *int64 `json:"whitelist,omitempty" xml:"whitelist,omitempty" require:"true"`
  // {"en":"Attack type classification requests(per second).","zh_CN":"攻击类型分类请求数（每秒均值）。"}
  Distribution []*GetTrendsByrpsResponseDataDistribution `json:"distribution,omitempty" xml:"distribution,omitempty" require:"true" type:"Repeated"`
}

func (s GetTrendsByrpsResponseData) String() string {
  return tea.Prettify(s)
}

func (s GetTrendsByrpsResponseData) GoString() string {
  return s.String()
}

func (s *GetTrendsByrpsResponseData) SetTimePoint(v string) *GetTrendsByrpsResponseData {
  s.TimePoint = &v
  return s
}

func (s *GetTrendsByrpsResponseData) SetTotal(v int64) *GetTrendsByrpsResponseData {
  s.Total = &v
  return s
}

func (s *GetTrendsByrpsResponseData) SetAttack(v int64) *GetTrendsByrpsResponseData {
  s.Attack = &v
  return s
}

func (s *GetTrendsByrpsResponseData) SetMitigated(v int64) *GetTrendsByrpsResponseData {
  s.Mitigated = &v
  return s
}

func (s *GetTrendsByrpsResponseData) SetMonitored(v int64) *GetTrendsByrpsResponseData {
  s.Monitored = &v
  return s
}

func (s *GetTrendsByrpsResponseData) SetWhitelist(v int64) *GetTrendsByrpsResponseData {
  s.Whitelist = &v
  return s
}

func (s *GetTrendsByrpsResponseData) SetDistribution(v []*GetTrendsByrpsResponseDataDistribution) *GetTrendsByrpsResponseData {
  s.Distribution = v
  return s
}

type GetTrendsByrpsResponseDataDistribution struct     {
  // {"en":"Attack type.\nBLOCK: IP/Geo Block\nDMS_DEFEND: DDoS Protection\nWAF_DEFEND: WAF\nBOT_MANAGE: Bot Management\nAPI_DEFEND: API Security\nINTELLIGENCE: Threat Intelligence\nRATE_LIMIT: Rate Limiting\nCUSTOMIZE_RULE: Custom Rules","zh_CN":"攻击类型。\nBLOCK：IP/区域封禁\nDMS_DEFEND：DDoS防护\nWAF_DEFEND：WAF\nBOT_MANAGE：Bot管理\nAPI_DEFEND：API安全\nINTELLIGENCE：威胁情报\nRATE_LIMIT：频率限制\nCUSTOMIZE_RULE：自定义规则","exampleValue":"BLOCK,DMS_DEFEND,WAF_DEFEND,BOT_MANAGE,API_DEFEND,INTELLIGENCE,RATE_LIMIT,CUSTOMIZE_RULE"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Number of attack requests of this type(per second).","zh_CN":"该类型攻击请求数（每秒均值）。"}
  Value *int64 `json:"value,omitempty" xml:"value,omitempty" require:"true"`
}

func (s GetTrendsByrpsResponseDataDistribution) String() string {
  return tea.Prettify(s)
}

func (s GetTrendsByrpsResponseDataDistribution) GoString() string {
  return s.String()
}

func (s *GetTrendsByrpsResponseDataDistribution) SetCode(v string) *GetTrendsByrpsResponseDataDistribution {
  s.Code = &v
  return s
}

func (s *GetTrendsByrpsResponseDataDistribution) SetValue(v int64) *GetTrendsByrpsResponseDataDistribution {
  s.Value = &v
  return s
}

type GetTrendsByrpsResponseHeader struct {
}

func (s GetTrendsByrpsResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetTrendsByrpsResponseHeader) GoString() string {
  return s.String()
}




type GetRequestTrendRequest struct {
  // {"en":"List of API group.", "zh_CN":"API分组列表。"}
  ApiGroups []*string `json:"apiGroups,omitempty" xml:"apiGroups,omitempty" type:"Repeated"`
  // {"en":"List of API name.", "zh_CN":"API名称列表。"}
  ApiIds []*string `json:"apiIds,omitempty" xml:"apiIds,omitempty" type:"Repeated"`
  // {"en":"List of domain.", "zh_CN":"域名列表。"}
  Domains []*string `json:"domains,omitempty" xml:"domains,omitempty" type:"Repeated"`
  // {"en":"Start time, format: yyyy-MM-dd HH:mm:ss.", "zh_CN":"开始时间，格式：yyyy-MM-dd HH:mm:ss。"}
  StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty" require:"true"`
  // {"en":"End time, format: yyyy-MM-dd HH:mm:ss.", "zh_CN":"结束时间，格式：yyyy-MM-dd HH:mm:ss。"}
  EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty" require:"true"`
  // {"en":"Endpoint path.", "zh_CN":"前端路径。"}
  FrontPath *string `json:"frontPath,omitempty" xml:"frontPath,omitempty"`
}

func (s GetRequestTrendRequest) String() string {
  return tea.Prettify(s)
}

func (s GetRequestTrendRequest) GoString() string {
  return s.String()
}

func (s *GetRequestTrendRequest) SetApiGroups(v []*string) *GetRequestTrendRequest {
  s.ApiGroups = v
  return s
}

func (s *GetRequestTrendRequest) SetApiIds(v []*string) *GetRequestTrendRequest {
  s.ApiIds = v
  return s
}

func (s *GetRequestTrendRequest) SetDomains(v []*string) *GetRequestTrendRequest {
  s.Domains = v
  return s
}

func (s *GetRequestTrendRequest) SetStartTime(v string) *GetRequestTrendRequest {
  s.StartTime = &v
  return s
}

func (s *GetRequestTrendRequest) SetEndTime(v string) *GetRequestTrendRequest {
  s.EndTime = &v
  return s
}

func (s *GetRequestTrendRequest) SetFrontPath(v string) *GetRequestTrendRequest {
  s.FrontPath = &v
  return s
}

type GetRequestTrendResponse struct {
  // {"en":"Return 200 means success.", "zh_CN":"200状态码表示请求成功。"}
  Code *int `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Message.", "zh_CN":"返回信息。"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"List of data.", "zh_CN":"数据列表。"}
  Data []*GetRequestTrendVo `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s GetRequestTrendResponse) String() string {
  return tea.Prettify(s)
}

func (s GetRequestTrendResponse) GoString() string {
  return s.String()
}

func (s *GetRequestTrendResponse) SetCode(v int) *GetRequestTrendResponse {
  s.Code = &v
  return s
}

func (s *GetRequestTrendResponse) SetMsg(v string) *GetRequestTrendResponse {
  s.Msg = &v
  return s
}

func (s *GetRequestTrendResponse) SetData(v []*GetRequestTrendVo) *GetRequestTrendResponse {
  s.Data = v
  return s
}

type GetRequestTrendVo struct {
  // {"en":"The total number of high concurrency limit.", "zh_CN":"限流封顶总数。"}
  FlowCelling *int `json:"flowCelling,omitempty" xml:"flowCelling,omitempty" require:"true"`
  // {"en":"The total number of noncompliance request body.", "zh_CN":"请求Body违规总数。"}
  IllegalRequestBody *int `json:"illegalRequestBody,omitempty" xml:"illegalRequestBody,omitempty" require:"true"`
  // {"en":"The total number of authentication failed.", "zh_CN":"鉴权失败总数。"}
  AuthFailed *int `json:"authFailed,omitempty" xml:"authFailed,omitempty" require:"true"`
  // {"en":"The total number of deactivate API call.", "zh_CN":"停用API调用总数。"}
  ApiDeactivateCall *int `json:"apiDeactivateCall,omitempty" xml:"apiDeactivateCall,omitempty" require:"true"`
  // {"en":"The total number of access control blacklist.", "zh_CN":"访问控制限制总数。"}
  AccessControlBlack *int `json:"accessControlBlack,omitempty" xml:"accessControlBlack,omitempty" require:"true"`
  // {"en":"The total number of back-to-source.", "zh_CN":"回源数。"}
  BackOrigin *int `json:"backOrigin,omitempty" xml:"backOrigin,omitempty" require:"true"`
  // {"en":"The total number of requests.", "zh_CN":"总请求数。"}
  TotalRequest *int `json:"totalRequest,omitempty" xml:"totalRequest,omitempty" require:"true"`
  // {"en":"The total number of noncompliance request method.", "zh_CN":"请求方法违规总数。"}
  RequestMethodLimit *int `json:"requestMethodLimit,omitempty" xml:"requestMethodLimit,omitempty" require:"true"`
  // {"en":"The total number of quota limit.", "zh_CN":"配额封顶总数。"}
  QuotaCelling *int `json:"quotaCelling,omitempty" xml:"quotaCelling,omitempty" require:"true"`
  // {"en":"The total number of remission.", "zh_CN":"缓解数。"}
  Remission *int `json:"remission,omitempty" xml:"remission,omitempty" require:"true"`
  // {"en":"The total number of unauthorized calls.", "zh_CN":"非授权调用总数。"}
  WithoutAuth *int `json:"withoutAuth,omitempty" xml:"withoutAuth,omitempty" require:"true"`
  // {"en":"The time point of request trend.", "zh_CN":"请求趋势时间点。"}
  Time *int `json:"time,omitempty" xml:"time,omitempty" require:"true"`
}

func (s GetRequestTrendVo) String() string {
  return tea.Prettify(s)
}

func (s GetRequestTrendVo) GoString() string {
  return s.String()
}

func (s *GetRequestTrendVo) SetFlowCelling(v int) *GetRequestTrendVo {
  s.FlowCelling = &v
  return s
}

func (s *GetRequestTrendVo) SetIllegalRequestBody(v int) *GetRequestTrendVo {
  s.IllegalRequestBody = &v
  return s
}

func (s *GetRequestTrendVo) SetAuthFailed(v int) *GetRequestTrendVo {
  s.AuthFailed = &v
  return s
}

func (s *GetRequestTrendVo) SetApiDeactivateCall(v int) *GetRequestTrendVo {
  s.ApiDeactivateCall = &v
  return s
}

func (s *GetRequestTrendVo) SetAccessControlBlack(v int) *GetRequestTrendVo {
  s.AccessControlBlack = &v
  return s
}

func (s *GetRequestTrendVo) SetBackOrigin(v int) *GetRequestTrendVo {
  s.BackOrigin = &v
  return s
}

func (s *GetRequestTrendVo) SetTotalRequest(v int) *GetRequestTrendVo {
  s.TotalRequest = &v
  return s
}

func (s *GetRequestTrendVo) SetRequestMethodLimit(v int) *GetRequestTrendVo {
  s.RequestMethodLimit = &v
  return s
}

func (s *GetRequestTrendVo) SetQuotaCelling(v int) *GetRequestTrendVo {
  s.QuotaCelling = &v
  return s
}

func (s *GetRequestTrendVo) SetRemission(v int) *GetRequestTrendVo {
  s.Remission = &v
  return s
}

func (s *GetRequestTrendVo) SetWithoutAuth(v int) *GetRequestTrendVo {
  s.WithoutAuth = &v
  return s
}

func (s *GetRequestTrendVo) SetTime(v int) *GetRequestTrendVo {
  s.Time = &v
  return s
}

type GetRequestTrendPaths struct {
}

func (s GetRequestTrendPaths) String() string {
  return tea.Prettify(s)
}

func (s GetRequestTrendPaths) GoString() string {
  return s.String()
}

type GetRequestTrendParameters struct {
}

func (s GetRequestTrendParameters) String() string {
  return tea.Prettify(s)
}

func (s GetRequestTrendParameters) GoString() string {
  return s.String()
}

type GetRequestTrendRequestHeader struct {
}

func (s GetRequestTrendRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetRequestTrendRequestHeader) GoString() string {
  return s.String()
}

type GetRequestTrendResponseHeader struct {
}

func (s GetRequestTrendResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetRequestTrendResponseHeader) GoString() string {
  return s.String()
}




type QueryCCAttackDetailsRequest struct {
  // {"en":"Start date format yyyy MM DD HH: mm: SS", "zh_CN":"起始日期 格式 yyyy-MM-dd HH:mm:ss"}
  Startdate *string `json:"startdate,omitempty" xml:"startdate,omitempty" require:"true"`
  // {"en":"End date format yyyy MM DD HH: mm: SS", "zh_CN":"结束日期 格式 yyyy-MM-dd HH:mm:ss"}
  Enddate *string `json:"enddate,omitempty" xml:"enddate,omitempty" require:"true"`
  // {"en":"Domain names, multiple separated by semicolons", "zh_CN":"域名,多个用分号分隔"}
  Channel *string `json:"channel,omitempty" xml:"channel,omitempty" require:"true"`
  // {"en":"Time zone offset,default:28800000", "zh_CN":"时区，默认为28800000"}
  TimeZone *string `json:"timeZone,omitempty" xml:"timeZone,omitempty"`
}

func (s QueryCCAttackDetailsRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryCCAttackDetailsRequest) GoString() string {
  return s.String()
}

func (s *QueryCCAttackDetailsRequest) SetStartdate(v string) *QueryCCAttackDetailsRequest {
  s.Startdate = &v
  return s
}

func (s *QueryCCAttackDetailsRequest) SetEnddate(v string) *QueryCCAttackDetailsRequest {
  s.Enddate = &v
  return s
}

func (s *QueryCCAttackDetailsRequest) SetChannel(v string) *QueryCCAttackDetailsRequest {
  s.Channel = &v
  return s
}

func (s *QueryCCAttackDetailsRequest) SetTimeZone(v string) *QueryCCAttackDetailsRequest {
  s.TimeZone = &v
  return s
}

type QueryCCAttackDetailsResponse struct {
  // {"en":"IP home, Chinese", "zh_CN":"ip归属地，中文"}
  Ip_country_cn *string `json:"ip_country_cn,omitempty" xml:"ip_country_cn,omitempty" require:"true"`
  // {"en":"IP home, English", "zh_CN":"ip归属地，英文"}
  Ip_country_en *string `json:"ip_country_en,omitempty" xml:"ip_country_en,omitempty" require:"true"`
  // {"en":"Attacked domain name", "zh_CN":"被攻击域名"}
  Attackdomain *string `json:"attackdomain,omitempty" xml:"attackdomain,omitempty" require:"true"`
  // {"en":"country code", "zh_CN":"国家编码"}
  Cn_country_code *string `json:"cn_country_code,omitempty" xml:"cn_country_code,omitempty" require:"true"`
  // {"en":"Country code, English", "zh_CN":"国家编码，英文"}
  En_country_code *string `json:"en_country_code,omitempty" xml:"en_country_code,omitempty" require:"true"`
  // {"en":"count", "zh_CN":"攻击数量"}
  Count *int `json:"count,omitempty" xml:"count,omitempty" require:"true"`
  // {"en":"City, English", "zh_CN":"城市，英文"}
  En_city *string `json:"en_city,omitempty" xml:"en_city,omitempty" require:"true"`
  // {"en":"City, Chinese", "zh_CN":"城市，中文"}
  Cn_city *string `json:"cn_city,omitempty" xml:"cn_city,omitempty" require:"true"`
  // {"en":"attackip", "zh_CN":"攻击ip"}
  Attackip *string `json:"attackip,omitempty" xml:"attackip,omitempty" require:"true"`
  // {"en":"IP Province, Chinese", "zh_CN":"ip所在省，中文"}
  Ip_province_cn *string `json:"ip_province_cn,omitempty" xml:"ip_province_cn,omitempty" require:"true"`
  // {"en":"IP Province, English", "zh_CN":"ip所在省，英文"}
  Ip_province_en *string `json:"ip_province_en,omitempty" xml:"ip_province_en,omitempty" require:"true"`
  // {"en":"IP City, English", "zh_CN":"ip所在城市，英文"}
  Ip_city_en *string `json:"ip_city_en,omitempty" xml:"ip_city_en,omitempty" require:"true"`
  // {"en":"IP City, Chinese", "zh_CN":"ip所在城市，中文"}
  Ip_city_cn *string `json:"ip_city_cn,omitempty" xml:"ip_city_cn,omitempty" require:"true"`
  // {"en":"time", "zh_CN":"时间"}
  Time *string `json:"time,omitempty" xml:"time,omitempty" require:"true"`
}

func (s QueryCCAttackDetailsResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryCCAttackDetailsResponse) GoString() string {
  return s.String()
}

func (s *QueryCCAttackDetailsResponse) SetIp_country_cn(v string) *QueryCCAttackDetailsResponse {
  s.Ip_country_cn = &v
  return s
}

func (s *QueryCCAttackDetailsResponse) SetIp_country_en(v string) *QueryCCAttackDetailsResponse {
  s.Ip_country_en = &v
  return s
}

func (s *QueryCCAttackDetailsResponse) SetAttackdomain(v string) *QueryCCAttackDetailsResponse {
  s.Attackdomain = &v
  return s
}

func (s *QueryCCAttackDetailsResponse) SetCn_country_code(v string) *QueryCCAttackDetailsResponse {
  s.Cn_country_code = &v
  return s
}

func (s *QueryCCAttackDetailsResponse) SetEn_country_code(v string) *QueryCCAttackDetailsResponse {
  s.En_country_code = &v
  return s
}

func (s *QueryCCAttackDetailsResponse) SetCount(v int) *QueryCCAttackDetailsResponse {
  s.Count = &v
  return s
}

func (s *QueryCCAttackDetailsResponse) SetEn_city(v string) *QueryCCAttackDetailsResponse {
  s.En_city = &v
  return s
}

func (s *QueryCCAttackDetailsResponse) SetCn_city(v string) *QueryCCAttackDetailsResponse {
  s.Cn_city = &v
  return s
}

func (s *QueryCCAttackDetailsResponse) SetAttackip(v string) *QueryCCAttackDetailsResponse {
  s.Attackip = &v
  return s
}

func (s *QueryCCAttackDetailsResponse) SetIp_province_cn(v string) *QueryCCAttackDetailsResponse {
  s.Ip_province_cn = &v
  return s
}

func (s *QueryCCAttackDetailsResponse) SetIp_province_en(v string) *QueryCCAttackDetailsResponse {
  s.Ip_province_en = &v
  return s
}

func (s *QueryCCAttackDetailsResponse) SetIp_city_en(v string) *QueryCCAttackDetailsResponse {
  s.Ip_city_en = &v
  return s
}

func (s *QueryCCAttackDetailsResponse) SetIp_city_cn(v string) *QueryCCAttackDetailsResponse {
  s.Ip_city_cn = &v
  return s
}

func (s *QueryCCAttackDetailsResponse) SetTime(v string) *QueryCCAttackDetailsResponse {
  s.Time = &v
  return s
}

type QueryCCAttackDetailsPaths struct {
}

func (s QueryCCAttackDetailsPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryCCAttackDetailsPaths) GoString() string {
  return s.String()
}

type QueryCCAttackDetailsParameters struct {
}

func (s QueryCCAttackDetailsParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryCCAttackDetailsParameters) GoString() string {
  return s.String()
}

type QueryCCAttackDetailsRequestHeader struct {
}

func (s QueryCCAttackDetailsRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryCCAttackDetailsRequestHeader) GoString() string {
  return s.String()
}

type QueryCCAttackDetailsResponseHeader struct {
}

func (s QueryCCAttackDetailsResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryCCAttackDetailsResponseHeader) GoString() string {
  return s.String()
}




type GetBotRequestRefererTopDataRequest struct {
  // {"en":"Domain.Separate by ';'.", "zh_CN":"域名。多个以;隔开"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {"en":"StartTime. Format: yyyy-MM-dd HH:mm:ss", "zh_CN":"开始时间。格式： yyyy-MM-dd HH:mm:ss"}
  StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty" require:"true"`
  // {"en":"EndTime. Format: yyyy-MM-dd HH:mm:ss", "zh_CN":"结束时间。格式： yyyy-MM-dd HH:mm:ss"}
  EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty" require:"true"`
  // {"en":"Time zone. Default 8, i.e.'GTM+8'", "zh_CN":"时区，默认8，即“GTM+8”"}
  TimeZone *int `json:"timeZone,omitempty" xml:"timeZone,omitempty"`
  // {"en":"Number of top values.Default value 10", "zh_CN":"排名最前值数目。默认10"}
  TopNum *int `json:"topNum,omitempty" xml:"topNum,omitempty"`
}

func (s GetBotRequestRefererTopDataRequest) String() string {
  return tea.Prettify(s)
}

func (s GetBotRequestRefererTopDataRequest) GoString() string {
  return s.String()
}

func (s *GetBotRequestRefererTopDataRequest) SetDomain(v string) *GetBotRequestRefererTopDataRequest {
  s.Domain = &v
  return s
}

func (s *GetBotRequestRefererTopDataRequest) SetStartTime(v string) *GetBotRequestRefererTopDataRequest {
  s.StartTime = &v
  return s
}

func (s *GetBotRequestRefererTopDataRequest) SetEndTime(v string) *GetBotRequestRefererTopDataRequest {
  s.EndTime = &v
  return s
}

func (s *GetBotRequestRefererTopDataRequest) SetTimeZone(v int) *GetBotRequestRefererTopDataRequest {
  s.TimeZone = &v
  return s
}

func (s *GetBotRequestRefererTopDataRequest) SetTopNum(v int) *GetBotRequestRefererTopDataRequest {
  s.TopNum = &v
  return s
}

type GetBotRequestRefererTopDataResponse struct {
  // {"en":"Status code, success is '200'.", "zh_CN":"状态码，成功为“200”。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Return message, success is 'Success'.", "zh_CN":"返回信息，成功为“Success”。"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Data rturned", "zh_CN":"返回数据"}
  Data []*GetBotRequestRefererTopDataResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s GetBotRequestRefererTopDataResponse) String() string {
  return tea.Prettify(s)
}

func (s GetBotRequestRefererTopDataResponse) GoString() string {
  return s.String()
}

func (s *GetBotRequestRefererTopDataResponse) SetCode(v string) *GetBotRequestRefererTopDataResponse {
  s.Code = &v
  return s
}

func (s *GetBotRequestRefererTopDataResponse) SetMessage(v string) *GetBotRequestRefererTopDataResponse {
  s.Message = &v
  return s
}

func (s *GetBotRequestRefererTopDataResponse) SetData(v []*GetBotRequestRefererTopDataResponseData) *GetBotRequestRefererTopDataResponse {
  s.Data = v
  return s
}

type GetBotRequestRefererTopDataResponseData struct     {
  // {"en":"Statistical type.", "zh_CN":"统计类型。"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"Bot request count", "zh_CN":"Bot请求数。"}
  Count *int64 `json:"count,omitempty" xml:"count,omitempty" require:"true"`
}

func (s GetBotRequestRefererTopDataResponseData) String() string {
  return tea.Prettify(s)
}

func (s GetBotRequestRefererTopDataResponseData) GoString() string {
  return s.String()
}

func (s *GetBotRequestRefererTopDataResponseData) SetName(v string) *GetBotRequestRefererTopDataResponseData {
  s.Name = &v
  return s
}

func (s *GetBotRequestRefererTopDataResponseData) SetCount(v int64) *GetBotRequestRefererTopDataResponseData {
  s.Count = &v
  return s
}

type GetBotRequestRefererTopDataPaths struct {
}

func (s GetBotRequestRefererTopDataPaths) String() string {
  return tea.Prettify(s)
}

func (s GetBotRequestRefererTopDataPaths) GoString() string {
  return s.String()
}

type GetBotRequestRefererTopDataParameters struct {
}

func (s GetBotRequestRefererTopDataParameters) String() string {
  return tea.Prettify(s)
}

func (s GetBotRequestRefererTopDataParameters) GoString() string {
  return s.String()
}

type GetBotRequestRefererTopDataRequestHeader struct {
}

func (s GetBotRequestRefererTopDataRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetBotRequestRefererTopDataRequestHeader) GoString() string {
  return s.String()
}

type GetBotRequestRefererTopDataResponseHeader struct {
}

func (s GetBotRequestRefererTopDataResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetBotRequestRefererTopDataResponseHeader) GoString() string {
  return s.String()
}




type GetRiskEventTop5Request struct {
  // {"en":"List of API groups.", "zh_CN":"API分组列表。"}
  ApiGroups []*string `json:"apiGroups,omitempty" xml:"apiGroups,omitempty" type:"Repeated"`
  // {"en":"List of API names.", "zh_CN":"API名称列表。"}
  ApiIds []*string `json:"apiIds,omitempty" xml:"apiIds,omitempty" type:"Repeated"`
  // {"en":"List of domains.", "zh_CN":"域名列表。"}
  Domains []*string `json:"domains,omitempty" xml:"domains,omitempty" type:"Repeated"`
  // {"en":"Start time, format: yyyy-MM-dd HH:mm:ss.", "zh_CN":"开始时间，格式：yyyy-MM-dd HH:mm:ss。"}
  StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty" require:"true"`
  // {"en":"End time, format: yyyy-MM-dd HH:mm:ss.", "zh_CN":"结束时间，格式：yyyy-MM-dd HH:mm:ss。"}
  EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty" require:"true"`
  // {"en":"Endpoint path.", "zh_CN":"前端路径。"}
  FrontPath *string `json:"frontPath,omitempty" xml:"frontPath,omitempty"`
}

func (s GetRiskEventTop5Request) String() string {
  return tea.Prettify(s)
}

func (s GetRiskEventTop5Request) GoString() string {
  return s.String()
}

func (s *GetRiskEventTop5Request) SetApiGroups(v []*string) *GetRiskEventTop5Request {
  s.ApiGroups = v
  return s
}

func (s *GetRiskEventTop5Request) SetApiIds(v []*string) *GetRiskEventTop5Request {
  s.ApiIds = v
  return s
}

func (s *GetRiskEventTop5Request) SetDomains(v []*string) *GetRiskEventTop5Request {
  s.Domains = v
  return s
}

func (s *GetRiskEventTop5Request) SetStartTime(v string) *GetRiskEventTop5Request {
  s.StartTime = &v
  return s
}

func (s *GetRiskEventTop5Request) SetEndTime(v string) *GetRiskEventTop5Request {
  s.EndTime = &v
  return s
}

func (s *GetRiskEventTop5Request) SetFrontPath(v string) *GetRiskEventTop5Request {
  s.FrontPath = &v
  return s
}

type GetRiskEventTop5Response struct {
  // {"en":"Return 200 means success.", "zh_CN":"200状态码表示请求成功。"}
  Code *int `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Message.", "zh_CN":"返回信息。"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"Data.", "zh_CN":"数据。"}
  Data []*GetRiskEventTop5Vo `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s GetRiskEventTop5Response) String() string {
  return tea.Prettify(s)
}

func (s GetRiskEventTop5Response) GoString() string {
  return s.String()
}

func (s *GetRiskEventTop5Response) SetCode(v int) *GetRiskEventTop5Response {
  s.Code = &v
  return s
}

func (s *GetRiskEventTop5Response) SetMsg(v string) *GetRiskEventTop5Response {
  s.Msg = &v
  return s
}

func (s *GetRiskEventTop5Response) SetData(v []*GetRiskEventTop5Vo) *GetRiskEventTop5Response {
  s.Data = v
  return s
}

type GetRiskEventTop5Vo struct {
  // {"en":"Attack type.
  // accessControlBlack:Access Control Blacklist
  // authFailed:Authentication Failed
  // illegalRequestBody:Noncompliance Request Body
  // apiDeactivateCall:Deactivate API call
  // requestMethodLimit:Noncompliance Request Method
  // illegalRequestParam:Noncompliance Parameter
  // quotaCelling:Quota Limit
  // flowCelling:High Concurrency Limit
  // withoutAuth:Unauthorized Object.
  // ", "zh_CN":"攻击类型。
  // accessControlBlack:访问控制限制
  // authFailed:鉴权失败
  // illegalRequestBody:请求Body违规
  // apiDeactivateCall:停用API调用
  // requestMethodLimit:请求方法违规
  // illegalRequestParam:请求参数违规
  // quotaCelling:配额封顶
  // flowCelling:限流封顶
  // withoutAuth:非授权调用。"}
  AttackType *string `json:"attackType,omitempty" xml:"attackType,omitempty" require:"true"`
  // {"en":"The number of certain risk records.", "zh_CN":"某种风险记录数。"}
  Count *int `json:"count,omitempty" xml:"count,omitempty" require:"true"`
}

func (s GetRiskEventTop5Vo) String() string {
  return tea.Prettify(s)
}

func (s GetRiskEventTop5Vo) GoString() string {
  return s.String()
}

func (s *GetRiskEventTop5Vo) SetAttackType(v string) *GetRiskEventTop5Vo {
  s.AttackType = &v
  return s
}

func (s *GetRiskEventTop5Vo) SetCount(v int) *GetRiskEventTop5Vo {
  s.Count = &v
  return s
}

type GetRiskEventTop5Paths struct {
}

func (s GetRiskEventTop5Paths) String() string {
  return tea.Prettify(s)
}

func (s GetRiskEventTop5Paths) GoString() string {
  return s.String()
}

type GetRiskEventTop5Parameters struct {
}

func (s GetRiskEventTop5Parameters) String() string {
  return tea.Prettify(s)
}

func (s GetRiskEventTop5Parameters) GoString() string {
  return s.String()
}

type GetRiskEventTop5RequestHeader struct {
}

func (s GetRiskEventTop5RequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetRiskEventTop5RequestHeader) GoString() string {
  return s.String()
}

type GetRiskEventTop5ResponseHeader struct {
}

func (s GetRiskEventTop5ResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetRiskEventTop5ResponseHeader) GoString() string {
  return s.String()
}




type GetTriggeredDDoSManagedRulesRequest struct {
  // {"defaultValue":"10","en":"Top rankings, default value: 10, max: 1000.","zh_CN":"取前几排名，默认：10，上限：1000。"}
  Top *int `json:"top,omitempty" xml:"top,omitempty"`
  // {"en":"Multiple selection. Handling results, default: display all results.\nmitigated: Number of mitigated requests.\nmonitored: Number of observed requests.","zh_CN":"多选。处置结果，默认：展示所有结果。\nmitigated：已抵御请求数。\nmonitored：观察请求数。","exampleValue":"mitigated,monitored"}
  ActType []*string `json:"actType,omitempty" xml:"actType,omitempty" type:"Repeated"`
  // {"en":"Start time, format: yyyy-MM-dd HH:mm:ss.\nThe time range can not exceed 31 days.","zh_CN":"起始时间，格式：yyyy-MM-dd HH:mm:ss。\n时间范围不超过31天。"}
  StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty" require:"true"`
  // {"en":"End time, format: yyyy-MM-dd HH:mm:ss.\nThe time range can not exceed 31 days.","zh_CN":"截止时间，格式：yyyy-MM-dd HH:mm:ss。\n时间范围不超过31天。"}
  EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty" require:"true"`
  // {"en":"Domain list. Queries all domains under the account when not specified.","zh_CN":"域名列表，未指定时查询账号下的所有域名。"}
  Domains []*string `json:"domains,omitempty" xml:"domains,omitempty" type:"Repeated"`
}

func (s GetTriggeredDDoSManagedRulesRequest) String() string {
  return tea.Prettify(s)
}

func (s GetTriggeredDDoSManagedRulesRequest) GoString() string {
  return s.String()
}

func (s *GetTriggeredDDoSManagedRulesRequest) SetTop(v int) *GetTriggeredDDoSManagedRulesRequest {
  s.Top = &v
  return s
}

func (s *GetTriggeredDDoSManagedRulesRequest) SetActType(v []*string) *GetTriggeredDDoSManagedRulesRequest {
  s.ActType = v
  return s
}

func (s *GetTriggeredDDoSManagedRulesRequest) SetStartTime(v string) *GetTriggeredDDoSManagedRulesRequest {
  s.StartTime = &v
  return s
}

func (s *GetTriggeredDDoSManagedRulesRequest) SetEndTime(v string) *GetTriggeredDDoSManagedRulesRequest {
  s.EndTime = &v
  return s
}

func (s *GetTriggeredDDoSManagedRulesRequest) SetDomains(v []*string) *GetTriggeredDDoSManagedRulesRequest {
  s.Domains = v
  return s
}

type GetTriggeredDDoSManagedRulesRequestHeader struct {
  // {"defaultValue":"en","en":"The language of response data, default value: en.\nzh_CN: Chinese\nen: English","zh_CN":"返回内容的语言版本，默认值: en。\nzh_CN：中文\nen：英文","exampleValue":"zh_CN,en"}
  Language *string `json:"language,omitempty" xml:"language,omitempty"`
  // {"dictionary":"belong=WAAP-MS-Ext|dict=waap_serviceType","en":"Security service type. Please enter a specific service type, if you purchase multiple security services.","zh_CN":"安全服务类型。有使用多个不同的安全服务时，需要填写具体的服务类型。"}
  ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty"`
  // {"defaultValue":"GMT+8","en":"Report Data Timezone:1. Indicates the timezone for the report data. lt must be relative to GMT andspecified in the format GMT+n where -12<=n <= 12. For example,timezone=GMT-7 and timezone=GMT+8.2. If the parameter is unspecified,results will be in the GMT+8 timezone.","zh_CN":"报表数据时区：1.请传递GMT时区。示例格式：GMT+N（其中12<=N <= 12），示例：timezone=GMT-7或timezone=GMT+8.2.若参数为传递，将默认按GMT+8时区查询"}
  Timezone *string `json:"timezone,omitempty" xml:"timezone,omitempty"`
}

func (s GetTriggeredDDoSManagedRulesRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetTriggeredDDoSManagedRulesRequestHeader) GoString() string {
  return s.String()
}

func (s *GetTriggeredDDoSManagedRulesRequestHeader) SetLanguage(v string) *GetTriggeredDDoSManagedRulesRequestHeader {
  s.Language = &v
  return s
}

func (s *GetTriggeredDDoSManagedRulesRequestHeader) SetServiceType(v string) *GetTriggeredDDoSManagedRulesRequestHeader {
  s.ServiceType = &v
  return s
}

func (s *GetTriggeredDDoSManagedRulesRequestHeader) SetTimezone(v string) *GetTriggeredDDoSManagedRulesRequestHeader {
  s.Timezone = &v
  return s
}

type GetTriggeredDDoSManagedRulesPaths struct {
}

func (s GetTriggeredDDoSManagedRulesPaths) String() string {
  return tea.Prettify(s)
}

func (s GetTriggeredDDoSManagedRulesPaths) GoString() string {
  return s.String()
}

type GetTriggeredDDoSManagedRulesParameters struct {
}

func (s GetTriggeredDDoSManagedRulesParameters) String() string {
  return tea.Prettify(s)
}

func (s GetTriggeredDDoSManagedRulesParameters) GoString() string {
  return s.String()
}

type GetTriggeredDDoSManagedRulesResponse struct {
  // {"dictionary":"belong=WAAP-MS-Ext|dict=waap_retCodeEnum","en":"Please refer to the error code for exceptions.","zh_CN":"请参照错误码。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Description.","zh_CN":"描述信息。"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"Data.","zh_CN":"出参数据。"}
  Data []*GetTriggeredDDoSManagedRulesResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s GetTriggeredDDoSManagedRulesResponse) String() string {
  return tea.Prettify(s)
}

func (s GetTriggeredDDoSManagedRulesResponse) GoString() string {
  return s.String()
}

func (s *GetTriggeredDDoSManagedRulesResponse) SetCode(v string) *GetTriggeredDDoSManagedRulesResponse {
  s.Code = &v
  return s
}

func (s *GetTriggeredDDoSManagedRulesResponse) SetMsg(v string) *GetTriggeredDDoSManagedRulesResponse {
  s.Msg = &v
  return s
}

func (s *GetTriggeredDDoSManagedRulesResponse) SetData(v []*GetTriggeredDDoSManagedRulesResponseData) *GetTriggeredDDoSManagedRulesResponse {
  s.Data = v
  return s
}

type GetTriggeredDDoSManagedRulesResponseData struct     {
  // {"en":"Rule ID.","zh_CN":"规则ID。"}
  RuleId *string `json:"ruleId,omitempty" xml:"ruleId,omitempty" require:"true"`
  // {"en":"Rule name.","zh_CN":"规则名称。"}
  RuleName *string `json:"ruleName,omitempty" xml:"ruleName,omitempty" require:"true"`
  // {"en":"Rule type.","zh_CN":"规则类型。"}
  RuleType *string `json:"ruleType,omitempty" xml:"ruleType,omitempty" require:"true"`
  // {"en":"System recommended action.","zh_CN":"系统推荐动作。"}
  Action *string `json:"action,omitempty" xml:"action,omitempty" require:"true"`
  // {"en":"Trigger times, sort by times.","zh_CN":"触发次数，按次数排序。"}
  Hits []*GetTriggeredDDoSManagedRulesResponseDataHits `json:"hits,omitempty" xml:"hits,omitempty" require:"true" type:"Repeated"`
}

func (s GetTriggeredDDoSManagedRulesResponseData) String() string {
  return tea.Prettify(s)
}

func (s GetTriggeredDDoSManagedRulesResponseData) GoString() string {
  return s.String()
}

func (s *GetTriggeredDDoSManagedRulesResponseData) SetRuleId(v string) *GetTriggeredDDoSManagedRulesResponseData {
  s.RuleId = &v
  return s
}

func (s *GetTriggeredDDoSManagedRulesResponseData) SetRuleName(v string) *GetTriggeredDDoSManagedRulesResponseData {
  s.RuleName = &v
  return s
}

func (s *GetTriggeredDDoSManagedRulesResponseData) SetRuleType(v string) *GetTriggeredDDoSManagedRulesResponseData {
  s.RuleType = &v
  return s
}

func (s *GetTriggeredDDoSManagedRulesResponseData) SetAction(v string) *GetTriggeredDDoSManagedRulesResponseData {
  s.Action = &v
  return s
}

func (s *GetTriggeredDDoSManagedRulesResponseData) SetHits(v []*GetTriggeredDDoSManagedRulesResponseDataHits) *GetTriggeredDDoSManagedRulesResponseData {
  s.Hits = v
  return s
}

type GetTriggeredDDoSManagedRulesResponseDataHits struct     {
  // {"en":"Action.","zh_CN":"采取动作。"}
  Act *string `json:"act,omitempty" xml:"act,omitempty" require:"true"`
  // {"en":"Hit times.","zh_CN":"命中数。"}
  Value *int64 `json:"value,omitempty" xml:"value,omitempty" require:"true"`
}

func (s GetTriggeredDDoSManagedRulesResponseDataHits) String() string {
  return tea.Prettify(s)
}

func (s GetTriggeredDDoSManagedRulesResponseDataHits) GoString() string {
  return s.String()
}

func (s *GetTriggeredDDoSManagedRulesResponseDataHits) SetAct(v string) *GetTriggeredDDoSManagedRulesResponseDataHits {
  s.Act = &v
  return s
}

func (s *GetTriggeredDDoSManagedRulesResponseDataHits) SetValue(v int64) *GetTriggeredDDoSManagedRulesResponseDataHits {
  s.Value = &v
  return s
}

type GetTriggeredDDoSManagedRulesResponseHeader struct {
}

func (s GetTriggeredDDoSManagedRulesResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetTriggeredDDoSManagedRulesResponseHeader) GoString() string {
  return s.String()
}




type GetTopAttackSourcesForChinaMainlandRequest struct {
  // {"defaultValue":"10","en":"Top rankings, default value: 10, max: 1000.","zh_CN":"取前几排名，默认：10，上限：1000。"}
  Top *int `json:"top,omitempty" xml:"top,omitempty"`
  // {"en":"Multiple selection. Handling results, default: display all results.\nmitigated: Number of mitigated requests.\nmonitored: Number of observed requests.","zh_CN":"多选。处置结果，默认：展示所有结果。\nmitigated：已抵御请求数。\nmonitored：观察请求数。","exampleValue":"mitigated,monitored"}
  ActType []*string `json:"actType,omitempty" xml:"actType,omitempty" type:"Repeated"`
  // {"en":"Start time, format: yyyy-MM-dd HH:mm:ss.\nThe time range can not exceed 31 days.","zh_CN":"起始时间，格式：yyyy-MM-dd HH:mm:ss。\n时间范围不超过31天。"}
  StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty" require:"true"`
  // {"en":"End time, format: yyyy-MM-dd HH:mm:ss.\nThe time range can not exceed 31 days.","zh_CN":"截止时间，格式：yyyy-MM-dd HH:mm:ss。\n时间范围不超过31天。"}
  EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty" require:"true"`
  // {"en":"Domain list. Queries all domains under the account when not specified.","zh_CN":"域名列表，未指定时查询账号下的所有域名。"}
  Domains []*string `json:"domains,omitempty" xml:"domains,omitempty" type:"Repeated"`
}

func (s GetTopAttackSourcesForChinaMainlandRequest) String() string {
  return tea.Prettify(s)
}

func (s GetTopAttackSourcesForChinaMainlandRequest) GoString() string {
  return s.String()
}

func (s *GetTopAttackSourcesForChinaMainlandRequest) SetTop(v int) *GetTopAttackSourcesForChinaMainlandRequest {
  s.Top = &v
  return s
}

func (s *GetTopAttackSourcesForChinaMainlandRequest) SetActType(v []*string) *GetTopAttackSourcesForChinaMainlandRequest {
  s.ActType = v
  return s
}

func (s *GetTopAttackSourcesForChinaMainlandRequest) SetStartTime(v string) *GetTopAttackSourcesForChinaMainlandRequest {
  s.StartTime = &v
  return s
}

func (s *GetTopAttackSourcesForChinaMainlandRequest) SetEndTime(v string) *GetTopAttackSourcesForChinaMainlandRequest {
  s.EndTime = &v
  return s
}

func (s *GetTopAttackSourcesForChinaMainlandRequest) SetDomains(v []*string) *GetTopAttackSourcesForChinaMainlandRequest {
  s.Domains = v
  return s
}

type GetTopAttackSourcesForChinaMainlandRequestHeader struct {
  // {"defaultValue":"en","en":"The language of response data, default value: en.\nzh_CN: Chinese\nen: English","zh_CN":"返回内容的语言版本，默认值: en。\nzh_CN：中文\nen：英文","exampleValue":"zh_CN,en"}
  Language *string `json:"language,omitempty" xml:"language,omitempty"`
  // {"dictionary":"belong=WAAP-MS-Ext|dict=waap_serviceType","en":"Security service type. Please enter a specific service type, if you purchase multiple security services.","zh_CN":"安全服务类型。有使用多个不同的安全服务时，需要填写具体的服务类型。"}
  ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty"`
  // {"defaultValue":"GMT+8","en":"Report Data Timezone:1. Indicates the timezone for the report data. lt must be relative to GMT andspecified in the format GMT+n where -12<=n <= 12. For example,timezone=GMT-7 and timezone=GMT+8.2. If the parameter is unspecified,results will be in the GMT+8 timezone.","zh_CN":"报表数据时区：1.请传递GMT时区。示例格式：GMT+N（其中12<=N <= 12），示例：timezone=GMT-7或timezone=GMT+8.2.若参数为传递，将默认按GMT+8时区查询"}
  Timezone *string `json:"timezone,omitempty" xml:"timezone,omitempty"`
}

func (s GetTopAttackSourcesForChinaMainlandRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetTopAttackSourcesForChinaMainlandRequestHeader) GoString() string {
  return s.String()
}

func (s *GetTopAttackSourcesForChinaMainlandRequestHeader) SetLanguage(v string) *GetTopAttackSourcesForChinaMainlandRequestHeader {
  s.Language = &v
  return s
}

func (s *GetTopAttackSourcesForChinaMainlandRequestHeader) SetServiceType(v string) *GetTopAttackSourcesForChinaMainlandRequestHeader {
  s.ServiceType = &v
  return s
}

func (s *GetTopAttackSourcesForChinaMainlandRequestHeader) SetTimezone(v string) *GetTopAttackSourcesForChinaMainlandRequestHeader {
  s.Timezone = &v
  return s
}

type GetTopAttackSourcesForChinaMainlandPaths struct {
}

func (s GetTopAttackSourcesForChinaMainlandPaths) String() string {
  return tea.Prettify(s)
}

func (s GetTopAttackSourcesForChinaMainlandPaths) GoString() string {
  return s.String()
}

type GetTopAttackSourcesForChinaMainlandParameters struct {
}

func (s GetTopAttackSourcesForChinaMainlandParameters) String() string {
  return tea.Prettify(s)
}

func (s GetTopAttackSourcesForChinaMainlandParameters) GoString() string {
  return s.String()
}

type GetTopAttackSourcesForChinaMainlandResponse struct {
  // {"dictionary":"belong=WAAP-MS-Ext|dict=waap_retCodeEnum","en":"Please refer to the error code for exceptions.","zh_CN":"请参照错误码。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Description.","zh_CN":"描述信息。"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"Data.","zh_CN":"出参数据。"}
  Data []*GetTopAttackSourcesForChinaMainlandResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s GetTopAttackSourcesForChinaMainlandResponse) String() string {
  return tea.Prettify(s)
}

func (s GetTopAttackSourcesForChinaMainlandResponse) GoString() string {
  return s.String()
}

func (s *GetTopAttackSourcesForChinaMainlandResponse) SetCode(v string) *GetTopAttackSourcesForChinaMainlandResponse {
  s.Code = &v
  return s
}

func (s *GetTopAttackSourcesForChinaMainlandResponse) SetMsg(v string) *GetTopAttackSourcesForChinaMainlandResponse {
  s.Msg = &v
  return s
}

func (s *GetTopAttackSourcesForChinaMainlandResponse) SetData(v []*GetTopAttackSourcesForChinaMainlandResponseData) *GetTopAttackSourcesForChinaMainlandResponse {
  s.Data = v
  return s
}

type GetTopAttackSourcesForChinaMainlandResponseData struct     {
  // {"en":"Source area(domestic province).","zh_CN":"来源地区（国内省份）。"}
  Area *string `json:"area,omitempty" xml:"area,omitempty" require:"true"`
  // {"en":"Source city.","zh_CN":"来源城市。"}
  City *string `json:"city,omitempty" xml:"city,omitempty" require:"true"`
  // {"en":"Attack requests.","zh_CN":"攻击请求数。"}
  Value *int64 `json:"value,omitempty" xml:"value,omitempty" require:"true"`
}

func (s GetTopAttackSourcesForChinaMainlandResponseData) String() string {
  return tea.Prettify(s)
}

func (s GetTopAttackSourcesForChinaMainlandResponseData) GoString() string {
  return s.String()
}

func (s *GetTopAttackSourcesForChinaMainlandResponseData) SetArea(v string) *GetTopAttackSourcesForChinaMainlandResponseData {
  s.Area = &v
  return s
}

func (s *GetTopAttackSourcesForChinaMainlandResponseData) SetCity(v string) *GetTopAttackSourcesForChinaMainlandResponseData {
  s.City = &v
  return s
}

func (s *GetTopAttackSourcesForChinaMainlandResponseData) SetValue(v int64) *GetTopAttackSourcesForChinaMainlandResponseData {
  s.Value = &v
  return s
}

type GetTopAttackSourcesForChinaMainlandResponseHeader struct {
}

func (s GetTopAttackSourcesForChinaMainlandResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetTopAttackSourcesForChinaMainlandResponseHeader) GoString() string {
  return s.String()
}




type GetWAFAttackEventRequest struct {
  // {'en':'Domain, array.', 'zh_CN':'域名，数组。'}
  Domains []*string `json:"domains,omitempty" xml:"domains,omitempty" require:"true" type:"Repeated"`
  // {'en':'Start time, yyyy-MM-dd HH:mm:ss.', 'zh_CN':'开始时间，yyyy-MM-dd HH:mm:ss。'}
  StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty" require:"true"`
  // {'en':'End time, yyyy-MM-dd HH:mm:ss.', 'zh_CN':'结束时间，yyyy-MM-dd HH:mm:ss。'}
  EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty" require:"true"`
  // {'en':'Time zone, GMT+8 by default.', 'zh_CN':'时区，默认GMT+8，即“GMT+8”。'}
  TimeZone *string `json:"timeZone,omitempty" xml:"timeZone,omitempty"`
}

func (s GetWAFAttackEventRequest) String() string {
  return tea.Prettify(s)
}

func (s GetWAFAttackEventRequest) GoString() string {
  return s.String()
}

func (s *GetWAFAttackEventRequest) SetDomains(v []*string) *GetWAFAttackEventRequest {
  s.Domains = v
  return s
}

func (s *GetWAFAttackEventRequest) SetStartTime(v string) *GetWAFAttackEventRequest {
  s.StartTime = &v
  return s
}

func (s *GetWAFAttackEventRequest) SetEndTime(v string) *GetWAFAttackEventRequest {
  s.EndTime = &v
  return s
}

func (s *GetWAFAttackEventRequest) SetTimeZone(v string) *GetWAFAttackEventRequest {
  s.TimeZone = &v
  return s
}

type GetWAFAttackEventResponse struct {
  // {"en":"Return 0 means success, please see <Error code> to check other status code.", "zh_CN":"0状态码表示请求成功，其他状态码说明请参见《错误码》。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Error message or Success.", "zh_CN":"错误信息或Success。"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Return data.", "zh_CN":"返回值。"}
  Data *GetWAFAttackEventAttackEventList `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s GetWAFAttackEventResponse) String() string {
  return tea.Prettify(s)
}

func (s GetWAFAttackEventResponse) GoString() string {
  return s.String()
}

func (s *GetWAFAttackEventResponse) SetCode(v string) *GetWAFAttackEventResponse {
  s.Code = &v
  return s
}

func (s *GetWAFAttackEventResponse) SetMessage(v string) *GetWAFAttackEventResponse {
  s.Message = &v
  return s
}

func (s *GetWAFAttackEventResponse) SetData(v *GetWAFAttackEventAttackEventList) *GetWAFAttackEventResponse {
  s.Data = v
  return s
}

type GetWAFAttackEventAttackEventList struct {
  // {"en":"Attack events.", "zh_CN":"攻击事件。"}
  GetWAFAttackEventAttackEventList []*GetWAFAttackEventAttackEvent `json:"attackEventList,omitempty" xml:"attackEventList,omitempty" require:"true" type:"Repeated"`
  // {"en":"Number of attack events.", "zh_CN":"攻击事件数。"}
  TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty" require:"true"`
}

func (s GetWAFAttackEventAttackEventList) String() string {
  return tea.Prettify(s)
}

func (s GetWAFAttackEventAttackEventList) GoString() string {
  return s.String()
}

func (s *GetWAFAttackEventAttackEventList) SetAttackEventList(v []*GetWAFAttackEventAttackEvent) *GetWAFAttackEventAttackEventList {
  s.GetWAFAttackEventAttackEventList = v
  return s
}

func (s *GetWAFAttackEventAttackEventList) SetTotalCount(v int64) *GetWAFAttackEventAttackEventList {
  s.TotalCount = &v
  return s
}

type GetWAFAttackEventAttackEvent struct {
  // {"en":"Attack requests.", "zh_CN":"攻击请求数。"}
  AttackCount *int64 `json:"attackCount,omitempty" xml:"attackCount,omitempty" require:"true"`
  // {"en":"Start time of attack.", "zh_CN":"攻击事件起始时间。"}
  StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty" require:"true"`
  // {"en":"End time of attack.", "zh_CN":"攻击事件截止时间。"}
  EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty" require:"true"`
  // {"en":"Attack type.", "zh_CN":"攻击类型。"}
  AttackType *string `json:"attackType,omitempty" xml:"attackType,omitempty" require:"true"`
  // {"en":"Attacker Ip.", "zh_CN":"攻击IP。"}
  Ip *string `json:"ip,omitempty" xml:"ip,omitempty" require:"true"`
  // {"en":"Attack event type.", "zh_CN":"攻击事件类型。"}
  AttackEventType *string `json:"attackEventType,omitempty" xml:"attackEventType,omitempty" require:"true"`
  // {"en":"English name of attack type.", "zh_CN":"攻击类别英文名称。"}
  AttackTypeNameEn *string `json:"attackTypeNameEn,omitempty" xml:"attackTypeNameEn,omitempty" require:"true"`
  // {"en":"Chinese name of attack type.", "zh_CN":"攻击类别中文名称。"}
  AttackTypeName *string `json:"attackTypeName,omitempty" xml:"attackTypeName,omitempty" require:"true"`
  // {"en":"English name of attack event.", "zh_CN":"攻击事件英文名称。"}
  EventNameEn *string `json:"eventNameEn,omitempty" xml:"eventNameEn,omitempty" require:"true"`
  // {"en":"Chinese name of attack event.", "zh_CN":"攻击事件中文名称。"}
  EventName *string `json:"eventName,omitempty" xml:"eventName,omitempty" require:"true"`
}

func (s GetWAFAttackEventAttackEvent) String() string {
  return tea.Prettify(s)
}

func (s GetWAFAttackEventAttackEvent) GoString() string {
  return s.String()
}

func (s *GetWAFAttackEventAttackEvent) SetAttackCount(v int64) *GetWAFAttackEventAttackEvent {
  s.AttackCount = &v
  return s
}

func (s *GetWAFAttackEventAttackEvent) SetStartTime(v string) *GetWAFAttackEventAttackEvent {
  s.StartTime = &v
  return s
}

func (s *GetWAFAttackEventAttackEvent) SetEndTime(v string) *GetWAFAttackEventAttackEvent {
  s.EndTime = &v
  return s
}

func (s *GetWAFAttackEventAttackEvent) SetAttackType(v string) *GetWAFAttackEventAttackEvent {
  s.AttackType = &v
  return s
}

func (s *GetWAFAttackEventAttackEvent) SetIp(v string) *GetWAFAttackEventAttackEvent {
  s.Ip = &v
  return s
}

func (s *GetWAFAttackEventAttackEvent) SetAttackEventType(v string) *GetWAFAttackEventAttackEvent {
  s.AttackEventType = &v
  return s
}

func (s *GetWAFAttackEventAttackEvent) SetAttackTypeNameEn(v string) *GetWAFAttackEventAttackEvent {
  s.AttackTypeNameEn = &v
  return s
}

func (s *GetWAFAttackEventAttackEvent) SetAttackTypeName(v string) *GetWAFAttackEventAttackEvent {
  s.AttackTypeName = &v
  return s
}

func (s *GetWAFAttackEventAttackEvent) SetEventNameEn(v string) *GetWAFAttackEventAttackEvent {
  s.EventNameEn = &v
  return s
}

func (s *GetWAFAttackEventAttackEvent) SetEventName(v string) *GetWAFAttackEventAttackEvent {
  s.EventName = &v
  return s
}

type GetWAFAttackEventPaths struct {
}

func (s GetWAFAttackEventPaths) String() string {
  return tea.Prettify(s)
}

func (s GetWAFAttackEventPaths) GoString() string {
  return s.String()
}

type GetWAFAttackEventParameters struct {
}

func (s GetWAFAttackEventParameters) String() string {
  return tea.Prettify(s)
}

func (s GetWAFAttackEventParameters) GoString() string {
  return s.String()
}

type GetWAFAttackEventRequestHeader struct {
}

func (s GetWAFAttackEventRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetWAFAttackEventRequestHeader) GoString() string {
  return s.String()
}

type GetWAFAttackEventResponseHeader struct {
}

func (s GetWAFAttackEventResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetWAFAttackEventResponseHeader) GoString() string {
  return s.String()
}




type GetRiskEventNumberRequest struct {
  // {"en":"List of API group.", "zh_CN":"API分组列表。"}
  ApiGroups []*string `json:"apiGroups,omitempty" xml:"apiGroups,omitempty" type:"Repeated"`
  // {"en":"List of API name.", "zh_CN":"API名称列表。"}
  ApiIds []*string `json:"apiIds,omitempty" xml:"apiIds,omitempty" type:"Repeated"`
  // {"en":"List of domain.", "zh_CN":"域名列表。"}
  Domains []*string `json:"domains,omitempty" xml:"domains,omitempty" type:"Repeated"`
  // {"en":"Start time, format: yyyy-MM-dd HH:mm:ss.", "zh_CN":"开始时间，格式：yyyy-MM-dd HH:mm:ss。"}
  StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty" require:"true"`
  // {"en":"End time, format: yyyy-MM-dd HH:mm:ss.", "zh_CN":"结束时间，格式：yyyy-MM-dd HH:mm:ss。"}
  EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty" require:"true"`
  // {"en":"Endpoint path.", "zh_CN":"前端路径。"}
  FrontPath *string `json:"frontPath,omitempty" xml:"frontPath,omitempty"`
}

func (s GetRiskEventNumberRequest) String() string {
  return tea.Prettify(s)
}

func (s GetRiskEventNumberRequest) GoString() string {
  return s.String()
}

func (s *GetRiskEventNumberRequest) SetApiGroups(v []*string) *GetRiskEventNumberRequest {
  s.ApiGroups = v
  return s
}

func (s *GetRiskEventNumberRequest) SetApiIds(v []*string) *GetRiskEventNumberRequest {
  s.ApiIds = v
  return s
}

func (s *GetRiskEventNumberRequest) SetDomains(v []*string) *GetRiskEventNumberRequest {
  s.Domains = v
  return s
}

func (s *GetRiskEventNumberRequest) SetStartTime(v string) *GetRiskEventNumberRequest {
  s.StartTime = &v
  return s
}

func (s *GetRiskEventNumberRequest) SetEndTime(v string) *GetRiskEventNumberRequest {
  s.EndTime = &v
  return s
}

func (s *GetRiskEventNumberRequest) SetFrontPath(v string) *GetRiskEventNumberRequest {
  s.FrontPath = &v
  return s
}

type GetRiskEventNumberResponse struct {
  // {"en":"Return 200 means success.", "zh_CN":"200状态码表示请求成功。"}
  Code *int `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Message.", "zh_CN":"返回信息。"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"Data.", "zh_CN":"数据。"}
  Data *GetRiskEventNumberVo `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s GetRiskEventNumberResponse) String() string {
  return tea.Prettify(s)
}

func (s GetRiskEventNumberResponse) GoString() string {
  return s.String()
}

func (s *GetRiskEventNumberResponse) SetCode(v int) *GetRiskEventNumberResponse {
  s.Code = &v
  return s
}

func (s *GetRiskEventNumberResponse) SetMsg(v string) *GetRiskEventNumberResponse {
  s.Msg = &v
  return s
}

func (s *GetRiskEventNumberResponse) SetData(v *GetRiskEventNumberVo) *GetRiskEventNumberResponse {
  s.Data = v
  return s
}

type GetRiskEventNumberVo struct {
  // {"en":"The total number of risk events.", "zh_CN":"风险事件总数。"}
  Total *int `json:"total,omitempty" xml:"total,omitempty" require:"true"`
  // {"en":"List of risk event.", "zh_CN":"风险事件列表。"}
  Data []*GetRiskEventNumberVoData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s GetRiskEventNumberVo) String() string {
  return tea.Prettify(s)
}

func (s GetRiskEventNumberVo) GoString() string {
  return s.String()
}

func (s *GetRiskEventNumberVo) SetTotal(v int) *GetRiskEventNumberVo {
  s.Total = &v
  return s
}

func (s *GetRiskEventNumberVo) SetData(v []*GetRiskEventNumberVoData) *GetRiskEventNumberVo {
  s.Data = v
  return s
}

type GetRiskEventNumberVoData struct {
  // {"en":"The time point of risk event, format: yyyy-MM-dd HH:mm:ss.", "zh_CN":"风险事件时间点，格式：yyyy-MM-dd HH:mm:ss。"}
  Time *string `json:"time,omitempty" xml:"time,omitempty" require:"true"`
  // {"en":"The number of risks at a certain point in time.", "zh_CN":"某时间点的风险数。"}
  Count *int `json:"count,omitempty" xml:"count,omitempty" require:"true"`
}

func (s GetRiskEventNumberVoData) String() string {
  return tea.Prettify(s)
}

func (s GetRiskEventNumberVoData) GoString() string {
  return s.String()
}

func (s *GetRiskEventNumberVoData) SetTime(v string) *GetRiskEventNumberVoData {
  s.Time = &v
  return s
}

func (s *GetRiskEventNumberVoData) SetCount(v int) *GetRiskEventNumberVoData {
  s.Count = &v
  return s
}

type GetRiskEventNumberPaths struct {
}

func (s GetRiskEventNumberPaths) String() string {
  return tea.Prettify(s)
}

func (s GetRiskEventNumberPaths) GoString() string {
  return s.String()
}

type GetRiskEventNumberParameters struct {
}

func (s GetRiskEventNumberParameters) String() string {
  return tea.Prettify(s)
}

func (s GetRiskEventNumberParameters) GoString() string {
  return s.String()
}

type GetRiskEventNumberRequestHeader struct {
}

func (s GetRiskEventNumberRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetRiskEventNumberRequestHeader) GoString() string {
  return s.String()
}

type GetRiskEventNumberResponseHeader struct {
}

func (s GetRiskEventNumberResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetRiskEventNumberResponseHeader) GoString() string {
  return s.String()
}




type GetWAFAttackedDomainRequest struct {
  // {'en':'Domain, array.', 'zh_CN':'域名，数组'}
  Domains []*string `json:"domains,omitempty" xml:"domains,omitempty" require:"true" type:"Repeated"`
  // {'en':'Start time, yyyy-MM-dd HH:mm:ss.', 'zh_CN':'开始时间，yyyy-MM-dd HH:mm:ss。'}
  StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty" require:"true"`
  // {'en':'End time, yyyy-MM-dd HH:mm:ss.', 'zh_CN':'结束时间，yyyy-MM-dd HH:mm:ss。'}
  EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty" require:"true"`
  // {'en':'Time zone, GMT+8 by default.', 'zh_CN':'时区，默认GMT+8，即“GMT+8”。'}
  TimeZone *string `json:"timeZone,omitempty" xml:"timeZone,omitempty"`
  // {'en':'Policy type, array. [protocol: Protocol Validation,
  // webShell: Webshell Access Detection,
  // other: Others Rules,
  // access: Access Control/Rate Limiting,
  // rule: WAF Rules]', 'zh_CN':'触发策略类型，数组 
  // [protocol：协议合规检测，
  // webShell：后门识别，
  // other：其他防护规则，
  // access：访问控制/限速，
  // rule：Web规则防护]'}
  Policys []*string `json:"policys,omitempty" xml:"policys,omitempty" type:"Repeated"`
}

func (s GetWAFAttackedDomainRequest) String() string {
  return tea.Prettify(s)
}

func (s GetWAFAttackedDomainRequest) GoString() string {
  return s.String()
}

func (s *GetWAFAttackedDomainRequest) SetDomains(v []*string) *GetWAFAttackedDomainRequest {
  s.Domains = v
  return s
}

func (s *GetWAFAttackedDomainRequest) SetStartTime(v string) *GetWAFAttackedDomainRequest {
  s.StartTime = &v
  return s
}

func (s *GetWAFAttackedDomainRequest) SetEndTime(v string) *GetWAFAttackedDomainRequest {
  s.EndTime = &v
  return s
}

func (s *GetWAFAttackedDomainRequest) SetTimeZone(v string) *GetWAFAttackedDomainRequest {
  s.TimeZone = &v
  return s
}

func (s *GetWAFAttackedDomainRequest) SetPolicys(v []*string) *GetWAFAttackedDomainRequest {
  s.Policys = v
  return s
}

type GetWAFAttackedDomainResponse struct {
  // {"en":"Return 0 means success, please see <Error code> to check other status code.", "zh_CN":"0状态码表示请求成功，其他状态码说明请参见《错误码》。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Error message or Success.", "zh_CN":"错误信息或Success。"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Return data.", "zh_CN":"返回值。"}
  Data *GetWAFAttackedDomainAttackedDomainList `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s GetWAFAttackedDomainResponse) String() string {
  return tea.Prettify(s)
}

func (s GetWAFAttackedDomainResponse) GoString() string {
  return s.String()
}

func (s *GetWAFAttackedDomainResponse) SetCode(v string) *GetWAFAttackedDomainResponse {
  s.Code = &v
  return s
}

func (s *GetWAFAttackedDomainResponse) SetMessage(v string) *GetWAFAttackedDomainResponse {
  s.Message = &v
  return s
}

func (s *GetWAFAttackedDomainResponse) SetData(v *GetWAFAttackedDomainAttackedDomainList) *GetWAFAttackedDomainResponse {
  s.Data = v
  return s
}

type GetWAFAttackedDomainAttackedDomainList struct {
  // {"en":"Total requests.", "zh_CN":"检测请求数。"}
  AllTotalCount *int64 `json:"allTotalCount,omitempty" xml:"allTotalCount,omitempty" require:"true"`
  // {"en":"Block requests.", "zh_CN":"拦截请求数。"}
  AllBlockCount *int64 `json:"allBlockCount,omitempty" xml:"allBlockCount,omitempty" require:"true"`
  // {"en":"Amount of list field.", "zh_CN":"List字段数据量。"}
  Count *int `json:"count,omitempty" xml:"count,omitempty" require:"true"`
  // {"en":"Attacked domains.", "zh_CN":"受攻击域名。"}
  List []*GetWAFAttackedDomainAttackedDomain `json:"list,omitempty" xml:"list,omitempty" require:"true" type:"Repeated"`
}

func (s GetWAFAttackedDomainAttackedDomainList) String() string {
  return tea.Prettify(s)
}

func (s GetWAFAttackedDomainAttackedDomainList) GoString() string {
  return s.String()
}

func (s *GetWAFAttackedDomainAttackedDomainList) SetAllTotalCount(v int64) *GetWAFAttackedDomainAttackedDomainList {
  s.AllTotalCount = &v
  return s
}

func (s *GetWAFAttackedDomainAttackedDomainList) SetAllBlockCount(v int64) *GetWAFAttackedDomainAttackedDomainList {
  s.AllBlockCount = &v
  return s
}

func (s *GetWAFAttackedDomainAttackedDomainList) SetCount(v int) *GetWAFAttackedDomainAttackedDomainList {
  s.Count = &v
  return s
}

func (s *GetWAFAttackedDomainAttackedDomainList) SetList(v []*GetWAFAttackedDomainAttackedDomain) *GetWAFAttackedDomainAttackedDomainList {
  s.List = v
  return s
}

type GetWAFAttackedDomainAttackedDomain struct {
  // {"en":"Attacked domains.", "zh_CN":"受攻击域名。"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {"en":"Attack type top 1.", "zh_CN":"Top1攻击类型。"}
  AttackTypeTop1 *string `json:"attackTypeTop1,omitempty" xml:"attackTypeTop1,omitempty" require:"true"`
  // {"en":"Total requests.", "zh_CN":"检测请求数。"}
  TotalHitCount *int64 `json:"totalHitCount,omitempty" xml:"totalHitCount,omitempty" require:"true"`
  // {"en":"Attack requests", "zh_CN":"攻击请求数。"}
  TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty" require:"true"`
  // {"en":"Block requests.", "zh_CN":"拦截请求数。"}
  BlockCount *int64 `json:"blockCount,omitempty" xml:"blockCount,omitempty" require:"true"`
  // {"en":"Log requests.", "zh_CN":"监控请求数。"}
  AlertCount *int64 `json:"alertCount,omitempty" xml:"alertCount,omitempty" require:"true"`
}

func (s GetWAFAttackedDomainAttackedDomain) String() string {
  return tea.Prettify(s)
}

func (s GetWAFAttackedDomainAttackedDomain) GoString() string {
  return s.String()
}

func (s *GetWAFAttackedDomainAttackedDomain) SetDomain(v string) *GetWAFAttackedDomainAttackedDomain {
  s.Domain = &v
  return s
}

func (s *GetWAFAttackedDomainAttackedDomain) SetAttackTypeTop1(v string) *GetWAFAttackedDomainAttackedDomain {
  s.AttackTypeTop1 = &v
  return s
}

func (s *GetWAFAttackedDomainAttackedDomain) SetTotalHitCount(v int64) *GetWAFAttackedDomainAttackedDomain {
  s.TotalHitCount = &v
  return s
}

func (s *GetWAFAttackedDomainAttackedDomain) SetTotalCount(v int64) *GetWAFAttackedDomainAttackedDomain {
  s.TotalCount = &v
  return s
}

func (s *GetWAFAttackedDomainAttackedDomain) SetBlockCount(v int64) *GetWAFAttackedDomainAttackedDomain {
  s.BlockCount = &v
  return s
}

func (s *GetWAFAttackedDomainAttackedDomain) SetAlertCount(v int64) *GetWAFAttackedDomainAttackedDomain {
  s.AlertCount = &v
  return s
}

type GetWAFAttackedDomainPaths struct {
}

func (s GetWAFAttackedDomainPaths) String() string {
  return tea.Prettify(s)
}

func (s GetWAFAttackedDomainPaths) GoString() string {
  return s.String()
}

type GetWAFAttackedDomainParameters struct {
}

func (s GetWAFAttackedDomainParameters) String() string {
  return tea.Prettify(s)
}

func (s GetWAFAttackedDomainParameters) GoString() string {
  return s.String()
}

type GetWAFAttackedDomainRequestHeader struct {
}

func (s GetWAFAttackedDomainRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetWAFAttackedDomainRequestHeader) GoString() string {
  return s.String()
}

type GetWAFAttackedDomainResponseHeader struct {
}

func (s GetWAFAttackedDomainResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetWAFAttackedDomainResponseHeader) GoString() string {
  return s.String()
}




type GetBotRequestOverviewDataRequest struct {
  // {"en":"Domain.Separate by ';'.", "zh_CN":"域名。多个以;隔开"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {"en":"StartTime. Format: yyyy-MM-dd HH:mm:ss", "zh_CN":"开始时间。格式： yyyy-MM-dd HH:mm:ss"}
  StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty" require:"true"`
  // {"en":"EndTime. Format: yyyy-MM-dd HH:mm:ss", "zh_CN":"结束时间。格式： yyyy-MM-dd HH:mm:ss"}
  EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty" require:"true"`
  // {"en":"Time zone, default 8, i.e.'GTM+8'", "zh_CN":"时区，默认8，即“GTM+8”"}
  TimeZone *int `json:"timeZone,omitempty" xml:"timeZone,omitempty"`
}

func (s GetBotRequestOverviewDataRequest) String() string {
  return tea.Prettify(s)
}

func (s GetBotRequestOverviewDataRequest) GoString() string {
  return s.String()
}

func (s *GetBotRequestOverviewDataRequest) SetDomain(v string) *GetBotRequestOverviewDataRequest {
  s.Domain = &v
  return s
}

func (s *GetBotRequestOverviewDataRequest) SetStartTime(v string) *GetBotRequestOverviewDataRequest {
  s.StartTime = &v
  return s
}

func (s *GetBotRequestOverviewDataRequest) SetEndTime(v string) *GetBotRequestOverviewDataRequest {
  s.EndTime = &v
  return s
}

func (s *GetBotRequestOverviewDataRequest) SetTimeZone(v int) *GetBotRequestOverviewDataRequest {
  s.TimeZone = &v
  return s
}

type GetBotRequestOverviewDataResponse struct {
  // {"en":"Status code, success is '200'.", "zh_CN":"状态码，成功为“200”。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Return message, success is 'Success'.", "zh_CN":"返回信息，成功为“Success”。"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Data returned.", "zh_CN":"返回数据。"}
  Data *GetBotRequestOverviewDataResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
}

func (s GetBotRequestOverviewDataResponse) String() string {
  return tea.Prettify(s)
}

func (s GetBotRequestOverviewDataResponse) GoString() string {
  return s.String()
}

func (s *GetBotRequestOverviewDataResponse) SetCode(v string) *GetBotRequestOverviewDataResponse {
  s.Code = &v
  return s
}

func (s *GetBotRequestOverviewDataResponse) SetMessage(v string) *GetBotRequestOverviewDataResponse {
  s.Message = &v
  return s
}

func (s *GetBotRequestOverviewDataResponse) SetData(v *GetBotRequestOverviewDataResponseData) *GetBotRequestOverviewDataResponse {
  s.Data = v
  return s
}

type GetBotRequestOverviewDataResponseData struct {
  // {"en":"Known bot type request count.", "zh_CN":"已知Bot类型请求数。"}
  GoodBotRequest *int64 `json:"goodBotRequest,omitempty" xml:"goodBotRequest,omitempty" require:"true"`
  // {"en":"Relief bot attack count.", "zh_CN":"缓解Bot攻击数。"}
  ReliefRequest *int64 `json:"reliefRequest,omitempty" xml:"reliefRequest,omitempty" require:"true"`
  // {"en":"Total request count.", "zh_CN":"总请求数。"}
  TotalRequest *int64 `json:"totalRequest,omitempty" xml:"totalRequest,omitempty" require:"true"`
  // {"en":"Unknown bot type request count.", "zh_CN":"未知Bot类型请求数。"}
  UnknowBotRequest *int64 `json:"unknowBotRequest,omitempty" xml:"unknowBotRequest,omitempty" require:"true"`
}

func (s GetBotRequestOverviewDataResponseData) String() string {
  return tea.Prettify(s)
}

func (s GetBotRequestOverviewDataResponseData) GoString() string {
  return s.String()
}

func (s *GetBotRequestOverviewDataResponseData) SetGoodBotRequest(v int64) *GetBotRequestOverviewDataResponseData {
  s.GoodBotRequest = &v
  return s
}

func (s *GetBotRequestOverviewDataResponseData) SetReliefRequest(v int64) *GetBotRequestOverviewDataResponseData {
  s.ReliefRequest = &v
  return s
}

func (s *GetBotRequestOverviewDataResponseData) SetTotalRequest(v int64) *GetBotRequestOverviewDataResponseData {
  s.TotalRequest = &v
  return s
}

func (s *GetBotRequestOverviewDataResponseData) SetUnknowBotRequest(v int64) *GetBotRequestOverviewDataResponseData {
  s.UnknowBotRequest = &v
  return s
}

type GetBotRequestOverviewDataPaths struct {
}

func (s GetBotRequestOverviewDataPaths) String() string {
  return tea.Prettify(s)
}

func (s GetBotRequestOverviewDataPaths) GoString() string {
  return s.String()
}

type GetBotRequestOverviewDataParameters struct {
}

func (s GetBotRequestOverviewDataParameters) String() string {
  return tea.Prettify(s)
}

func (s GetBotRequestOverviewDataParameters) GoString() string {
  return s.String()
}

type GetBotRequestOverviewDataRequestHeader struct {
}

func (s GetBotRequestOverviewDataRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetBotRequestOverviewDataRequestHeader) GoString() string {
  return s.String()
}

type GetBotRequestOverviewDataResponseHeader struct {
}

func (s GetBotRequestOverviewDataResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetBotRequestOverviewDataResponseHeader) GoString() string {
  return s.String()
}




type GetTopAttackTargetsByPathRequest struct {
  // {"defaultValue":"100","en":"Top rankings, default: 100, max: 1000.","zh_CN":"取前几排名，默认100，上限1000。"}
  Top *int `json:"top,omitempty" xml:"top,omitempty"`
  // {"en":"Multiple selection. Handling results, default: display all results.\nmitigated: Number of mitigated requests.\nmonitored: Number of observed requests.","zh_CN":"多选。处置结果，默认：展示所有结果。\nmitigated：已抵御请求数。\nmonitored：观察请求数。","exampleValue":"mitigated,monitored"}
  ActType []*string `json:"actType,omitempty" xml:"actType,omitempty" type:"Repeated"`
  // {"en":"Start time, format: yyyy-MM-dd HH:mm:ss.\nThe time range can not exceed 31 days.","zh_CN":"起始时间，格式：yyyy-MM-dd HH:mm:ss。\n时间范围不超过31天。"}
  StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty" require:"true"`
  // {"en":"End time, format: yyyy-MM-dd HH:mm:ss.\nThe time range can not exceed 31 days.","zh_CN":"截止时间，格式：yyyy-MM-dd HH:mm:ss。\n时间范围不超过31天。"}
  EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty" require:"true"`
  // {"en":"Domain list. Queries all domains under the account when not specified.","zh_CN":"域名列表，未指定时查询账号下的所有域名。"}
  Domains []*string `json:"domains,omitempty" xml:"domains,omitempty" type:"Repeated"`
}

func (s GetTopAttackTargetsByPathRequest) String() string {
  return tea.Prettify(s)
}

func (s GetTopAttackTargetsByPathRequest) GoString() string {
  return s.String()
}

func (s *GetTopAttackTargetsByPathRequest) SetTop(v int) *GetTopAttackTargetsByPathRequest {
  s.Top = &v
  return s
}

func (s *GetTopAttackTargetsByPathRequest) SetActType(v []*string) *GetTopAttackTargetsByPathRequest {
  s.ActType = v
  return s
}

func (s *GetTopAttackTargetsByPathRequest) SetStartTime(v string) *GetTopAttackTargetsByPathRequest {
  s.StartTime = &v
  return s
}

func (s *GetTopAttackTargetsByPathRequest) SetEndTime(v string) *GetTopAttackTargetsByPathRequest {
  s.EndTime = &v
  return s
}

func (s *GetTopAttackTargetsByPathRequest) SetDomains(v []*string) *GetTopAttackTargetsByPathRequest {
  s.Domains = v
  return s
}

type GetTopAttackTargetsByPathRequestHeader struct {
  // {"dictionary":"belong=WAAP-MS-Ext|dict=waap_serviceType","en":"Security service type. Please enter a specific service type, if you purchase multiple security services.","zh_CN":"安全服务类型。有使用多个不同的安全服务时，需要填写具体的服务类型。"}
  ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty"`
  // {"defaultValue":"GMT+8","en":"Report Data Timezone:1. Indicates the timezone for the report data. lt must be relative to GMT andspecified in the format GMT+n where -12<=n <= 12. For example,timezone=GMT-7 and timezone=GMT+8.2. If the parameter is unspecified,results will be in the GMT+8 timezone.","zh_CN":"报表数据时区：1.请传递GMT时区。示例格式：GMT+N（其中12<=N <= 12），示例：timezone=GMT-7或timezone=GMT+8.2.若参数为传递，将默认按GMT+8时区查询"}
  Timezone *string `json:"timezone,omitempty" xml:"timezone,omitempty"`
}

func (s GetTopAttackTargetsByPathRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetTopAttackTargetsByPathRequestHeader) GoString() string {
  return s.String()
}

func (s *GetTopAttackTargetsByPathRequestHeader) SetServiceType(v string) *GetTopAttackTargetsByPathRequestHeader {
  s.ServiceType = &v
  return s
}

func (s *GetTopAttackTargetsByPathRequestHeader) SetTimezone(v string) *GetTopAttackTargetsByPathRequestHeader {
  s.Timezone = &v
  return s
}

type GetTopAttackTargetsByPathPaths struct {
}

func (s GetTopAttackTargetsByPathPaths) String() string {
  return tea.Prettify(s)
}

func (s GetTopAttackTargetsByPathPaths) GoString() string {
  return s.String()
}

type GetTopAttackTargetsByPathParameters struct {
}

func (s GetTopAttackTargetsByPathParameters) String() string {
  return tea.Prettify(s)
}

func (s GetTopAttackTargetsByPathParameters) GoString() string {
  return s.String()
}

type GetTopAttackTargetsByPathResponse struct {
  // {"dictionary":"belong=WAAP-MS-Ext|dict=waap_retCodeEnum","en":"Please refer to the error code for exceptions.","zh_CN":"请参照错误码。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Description.","zh_CN":"描述信息。"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"Data.","zh_CN":"出参数据。"}
  Data []*GetTopAttackTargetsByPathResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s GetTopAttackTargetsByPathResponse) String() string {
  return tea.Prettify(s)
}

func (s GetTopAttackTargetsByPathResponse) GoString() string {
  return s.String()
}

func (s *GetTopAttackTargetsByPathResponse) SetCode(v string) *GetTopAttackTargetsByPathResponse {
  s.Code = &v
  return s
}

func (s *GetTopAttackTargetsByPathResponse) SetMsg(v string) *GetTopAttackTargetsByPathResponse {
  s.Msg = &v
  return s
}

func (s *GetTopAttackTargetsByPathResponse) SetData(v []*GetTopAttackTargetsByPathResponseData) *GetTopAttackTargetsByPathResponse {
  s.Data = v
  return s
}

type GetTopAttackTargetsByPathResponseData struct     {
  // {"en":"URL.","zh_CN":"URL。"}
  Url *string `json:"url,omitempty" xml:"url,omitempty" require:"true"`
  // {"en":"Attack requests.","zh_CN":"攻击请求数。"}
  Attack *int64 `json:"attack,omitempty" xml:"attack,omitempty" require:"true"`
}

func (s GetTopAttackTargetsByPathResponseData) String() string {
  return tea.Prettify(s)
}

func (s GetTopAttackTargetsByPathResponseData) GoString() string {
  return s.String()
}

func (s *GetTopAttackTargetsByPathResponseData) SetUrl(v string) *GetTopAttackTargetsByPathResponseData {
  s.Url = &v
  return s
}

func (s *GetTopAttackTargetsByPathResponseData) SetAttack(v int64) *GetTopAttackTargetsByPathResponseData {
  s.Attack = &v
  return s
}

type GetTopAttackTargetsByPathResponseHeader struct {
}

func (s GetTopAttackTargetsByPathResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetTopAttackTargetsByPathResponseHeader) GoString() string {
  return s.String()
}




type GetActiveApiTop5Request struct {
  // {"en":"List of API group.", "zh_CN":"API分组列表。"}
  ApiGroups []*string `json:"apiGroups,omitempty" xml:"apiGroups,omitempty" type:"Repeated"`
  // {"en":"List of API name.", "zh_CN":"API名称列表。"}
  ApiIds []*string `json:"apiIds,omitempty" xml:"apiIds,omitempty" type:"Repeated"`
  // {"en":"List of domain.", "zh_CN":"域名列表。"}
  Domains []*string `json:"domains,omitempty" xml:"domains,omitempty" type:"Repeated"`
  // {"en":"Start time, format: yyyy-MM-dd HH:mm:ss.", "zh_CN":"开始时间，格式：yyyy-MM-dd HH:mm:ss。"}
  StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty" require:"true"`
  // {"en":"End time, format: yyyy-MM-dd HH:mm:ss.", "zh_CN":"结束时间，格式：yyyy-MM-dd HH:mm:ss。"}
  EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty" require:"true"`
  // {"en":"Endpoint path.", "zh_CN":"前端路径。"}
  FrontPath *string `json:"frontPath,omitempty" xml:"frontPath,omitempty"`
}

func (s GetActiveApiTop5Request) String() string {
  return tea.Prettify(s)
}

func (s GetActiveApiTop5Request) GoString() string {
  return s.String()
}

func (s *GetActiveApiTop5Request) SetApiGroups(v []*string) *GetActiveApiTop5Request {
  s.ApiGroups = v
  return s
}

func (s *GetActiveApiTop5Request) SetApiIds(v []*string) *GetActiveApiTop5Request {
  s.ApiIds = v
  return s
}

func (s *GetActiveApiTop5Request) SetDomains(v []*string) *GetActiveApiTop5Request {
  s.Domains = v
  return s
}

func (s *GetActiveApiTop5Request) SetStartTime(v string) *GetActiveApiTop5Request {
  s.StartTime = &v
  return s
}

func (s *GetActiveApiTop5Request) SetEndTime(v string) *GetActiveApiTop5Request {
  s.EndTime = &v
  return s
}

func (s *GetActiveApiTop5Request) SetFrontPath(v string) *GetActiveApiTop5Request {
  s.FrontPath = &v
  return s
}

type GetActiveApiTop5Response struct {
  // {"en":"Return 200 means success.", "zh_CN":"200状态码表示请求成功。"}
  Code *int `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Message.", "zh_CN":"返回信息。"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"Data.", "zh_CN":"数据。"}
  Data []*GetActiveApiTop5Vo `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s GetActiveApiTop5Response) String() string {
  return tea.Prettify(s)
}

func (s GetActiveApiTop5Response) GoString() string {
  return s.String()
}

func (s *GetActiveApiTop5Response) SetCode(v int) *GetActiveApiTop5Response {
  s.Code = &v
  return s
}

func (s *GetActiveApiTop5Response) SetMsg(v string) *GetActiveApiTop5Response {
  s.Msg = &v
  return s
}

func (s *GetActiveApiTop5Response) SetData(v []*GetActiveApiTop5Vo) *GetActiveApiTop5Response {
  s.Data = v
  return s
}

type GetActiveApiTop5Vo struct {
  // {"en":"API name.", "zh_CN":"API名称。"}
  ApiName *string `json:"apiName,omitempty" xml:"apiName,omitempty" require:"true"`
  // {"en":"The total number of requests.", "zh_CN":"调用总次数。"}
  Count *int `json:"count,omitempty" xml:"count,omitempty" require:"true"`
  // {"en":"HTTP request URL.", "zh_CN":"URL地址。"}
  Url *string `json:"url,omitempty" xml:"url,omitempty" require:"true"`
}

func (s GetActiveApiTop5Vo) String() string {
  return tea.Prettify(s)
}

func (s GetActiveApiTop5Vo) GoString() string {
  return s.String()
}

func (s *GetActiveApiTop5Vo) SetApiName(v string) *GetActiveApiTop5Vo {
  s.ApiName = &v
  return s
}

func (s *GetActiveApiTop5Vo) SetCount(v int) *GetActiveApiTop5Vo {
  s.Count = &v
  return s
}

func (s *GetActiveApiTop5Vo) SetUrl(v string) *GetActiveApiTop5Vo {
  s.Url = &v
  return s
}

type GetActiveApiTop5Paths struct {
}

func (s GetActiveApiTop5Paths) String() string {
  return tea.Prettify(s)
}

func (s GetActiveApiTop5Paths) GoString() string {
  return s.String()
}

type GetActiveApiTop5Parameters struct {
}

func (s GetActiveApiTop5Parameters) String() string {
  return tea.Prettify(s)
}

func (s GetActiveApiTop5Parameters) GoString() string {
  return s.String()
}

type GetActiveApiTop5RequestHeader struct {
}

func (s GetActiveApiTop5RequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetActiveApiTop5RequestHeader) GoString() string {
  return s.String()
}

type GetActiveApiTop5ResponseHeader struct {
}

func (s GetActiveApiTop5ResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetActiveApiTop5ResponseHeader) GoString() string {
  return s.String()
}




type GetBotRequestSourceIPTopDataRequest struct {
  // {"en":"Domain.Separate by ';'.", "zh_CN":"域名。多个以;隔开"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {"en":"StartTime. Format: yyyy-MM-dd HH:mm:ss", "zh_CN":"开始时间。格式： yyyy-MM-dd HH:mm:ss"}
  StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty" require:"true"`
  // {"en":"EndTime. Format: yyyy-MM-dd HH:mm:ss", "zh_CN":"结束时间。格式： yyyy-MM-dd HH:mm:ss"}
  EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty" require:"true"`
  // {"en":"Time zone, default 8, i.e. 'GTM+8'", "zh_CN":"时区，默认8，即“GTM+8”"}
  TimeZone *int `json:"timeZone,omitempty" xml:"timeZone,omitempty"`
  // {"en":"Language type. Default cn. 
  //  en:English 
  //  cn:Chinese ", "zh_CN":"语言类型。 默认cn 
  //  en：英文 
  //  cn：中文"}
  Lang *string `json:"lang,omitempty" xml:"lang,omitempty"`
  // {"en":"Number of top values.Default value 10", "zh_CN":"排名最前值数目。默认10"}
  TopNum *int `json:"topNum,omitempty" xml:"topNum,omitempty"`
}

func (s GetBotRequestSourceIPTopDataRequest) String() string {
  return tea.Prettify(s)
}

func (s GetBotRequestSourceIPTopDataRequest) GoString() string {
  return s.String()
}

func (s *GetBotRequestSourceIPTopDataRequest) SetDomain(v string) *GetBotRequestSourceIPTopDataRequest {
  s.Domain = &v
  return s
}

func (s *GetBotRequestSourceIPTopDataRequest) SetStartTime(v string) *GetBotRequestSourceIPTopDataRequest {
  s.StartTime = &v
  return s
}

func (s *GetBotRequestSourceIPTopDataRequest) SetEndTime(v string) *GetBotRequestSourceIPTopDataRequest {
  s.EndTime = &v
  return s
}

func (s *GetBotRequestSourceIPTopDataRequest) SetTimeZone(v int) *GetBotRequestSourceIPTopDataRequest {
  s.TimeZone = &v
  return s
}

func (s *GetBotRequestSourceIPTopDataRequest) SetLang(v string) *GetBotRequestSourceIPTopDataRequest {
  s.Lang = &v
  return s
}

func (s *GetBotRequestSourceIPTopDataRequest) SetTopNum(v int) *GetBotRequestSourceIPTopDataRequest {
  s.TopNum = &v
  return s
}

type GetBotRequestSourceIPTopDataResponse struct {
  // {"en":"Status code, success is '200'.", "zh_CN":"状态码，成功为“200”。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Return message, success is 'Success'.", "zh_CN":"返回信息，成功为“Success”。"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Data returned.", "zh_CN":"返回数据。"}
  Data []*GetBotRequestSourceIPTopDataResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s GetBotRequestSourceIPTopDataResponse) String() string {
  return tea.Prettify(s)
}

func (s GetBotRequestSourceIPTopDataResponse) GoString() string {
  return s.String()
}

func (s *GetBotRequestSourceIPTopDataResponse) SetCode(v string) *GetBotRequestSourceIPTopDataResponse {
  s.Code = &v
  return s
}

func (s *GetBotRequestSourceIPTopDataResponse) SetMessage(v string) *GetBotRequestSourceIPTopDataResponse {
  s.Message = &v
  return s
}

func (s *GetBotRequestSourceIPTopDataResponse) SetData(v []*GetBotRequestSourceIPTopDataResponseData) *GetBotRequestSourceIPTopDataResponse {
  s.Data = v
  return s
}

type GetBotRequestSourceIPTopDataResponseData struct     {
  // {"en":"Statistical type.", "zh_CN":"统计类型。"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"Area.", "zh_CN":"地区。"}
  Province *string `json:"province,omitempty" xml:"province,omitempty" require:"true"`
  // {"en":"Bot request count.", "zh_CN":"Bot请求数。"}
  Count *int64 `json:"count,omitempty" xml:"count,omitempty" require:"true"`
}

func (s GetBotRequestSourceIPTopDataResponseData) String() string {
  return tea.Prettify(s)
}

func (s GetBotRequestSourceIPTopDataResponseData) GoString() string {
  return s.String()
}

func (s *GetBotRequestSourceIPTopDataResponseData) SetName(v string) *GetBotRequestSourceIPTopDataResponseData {
  s.Name = &v
  return s
}

func (s *GetBotRequestSourceIPTopDataResponseData) SetProvince(v string) *GetBotRequestSourceIPTopDataResponseData {
  s.Province = &v
  return s
}

func (s *GetBotRequestSourceIPTopDataResponseData) SetCount(v int64) *GetBotRequestSourceIPTopDataResponseData {
  s.Count = &v
  return s
}

type GetBotRequestSourceIPTopDataPaths struct {
}

func (s GetBotRequestSourceIPTopDataPaths) String() string {
  return tea.Prettify(s)
}

func (s GetBotRequestSourceIPTopDataPaths) GoString() string {
  return s.String()
}

type GetBotRequestSourceIPTopDataParameters struct {
}

func (s GetBotRequestSourceIPTopDataParameters) String() string {
  return tea.Prettify(s)
}

func (s GetBotRequestSourceIPTopDataParameters) GoString() string {
  return s.String()
}

type GetBotRequestSourceIPTopDataRequestHeader struct {
}

func (s GetBotRequestSourceIPTopDataRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetBotRequestSourceIPTopDataRequestHeader) GoString() string {
  return s.String()
}

type GetBotRequestSourceIPTopDataResponseHeader struct {
}

func (s GetBotRequestSourceIPTopDataResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetBotRequestSourceIPTopDataResponseHeader) GoString() string {
  return s.String()
}




type CcAttackQpsForQueryRequest struct {
  // {"en":"start time format: "yyyy-MM-dd HH:mm:ss"", "zh_CN":"开始时间("yyyy-MM-dd HH:mm:ss")"}
  Startdate *string `json:"startdate,omitempty" xml:"startdate,omitempty"`
  // {"en":"packageId or acctype should be selected at least one.", "zh_CN":"套餐ID: packageId和acctype至少传一个,但不能同时传"}
  PackageId *string `json:"packageId,omitempty" xml:"packageId,omitempty"`
  // {"en":"need Detail : 1
  // no need Detail: 0 default : 1", "zh_CN":"是否需要查看域名或是转发规则带宽的详细信息：0：不需要；1：需要，默认需要"}
  NeedDetail *string `json:"needDetail,omitempty" xml:"needDetail,omitempty"`
  // {"en":"end time format: "yyyy-MM-dd HH:mm:ss"", "zh_CN":"结束时间("yyyy-MM-dd HH:mm:ss")"}
  Enddate *string `json:"enddate,omitempty" xml:"enddate,omitempty"`
  // {"en":"domain, split by semicolon(";").", "zh_CN":"域名，支持多个用英文半角分号分隔；不传默认查询全部域名"}
  Domains *string `json:"domains,omitempty" xml:"domains,omitempty"`
  // {"en":"custom code", "zh_CN":"客户英文名"}
  CustomCode *string `json:"customCode,omitempty" xml:"customCode,omitempty"`
  // {"en":"packageId or acctype should be selected at least one. acctype( Only One can be selected): gess，fsa，app-s，dms-https，wss, dms， wss-https，s-appa， wsa，esa，wsa-https", "zh_CN":"packageId和acctype不能同时传且至少传一个；产品外部服务类型,只支持传1个:gess，fsa，app-s，dms-https，wss, dms， wss-https，s-appa， wsa，esa，wsa-https"}
  Acctype *string `json:"acctype,omitempty" xml:"acctype,omitempty"`
}

func (s CcAttackQpsForQueryRequest) String() string {
  return tea.Prettify(s)
}

func (s CcAttackQpsForQueryRequest) GoString() string {
  return s.String()
}

func (s *CcAttackQpsForQueryRequest) SetStartdate(v string) *CcAttackQpsForQueryRequest {
  s.Startdate = &v
  return s
}

func (s *CcAttackQpsForQueryRequest) SetPackageId(v string) *CcAttackQpsForQueryRequest {
  s.PackageId = &v
  return s
}

func (s *CcAttackQpsForQueryRequest) SetNeedDetail(v string) *CcAttackQpsForQueryRequest {
  s.NeedDetail = &v
  return s
}

func (s *CcAttackQpsForQueryRequest) SetEnddate(v string) *CcAttackQpsForQueryRequest {
  s.Enddate = &v
  return s
}

func (s *CcAttackQpsForQueryRequest) SetDomains(v string) *CcAttackQpsForQueryRequest {
  s.Domains = &v
  return s
}

func (s *CcAttackQpsForQueryRequest) SetCustomCode(v string) *CcAttackQpsForQueryRequest {
  s.CustomCode = &v
  return s
}

func (s *CcAttackQpsForQueryRequest) SetAcctype(v string) *CcAttackQpsForQueryRequest {
  s.Acctype = &v
  return s
}

type CcAttackQpsForQueryResponse struct {
  // {"en":"response message", "zh_CN":"响应信息，成功时为success"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"Time of CC attack", "zh_CN":"统计CC攻击的时间"}
  Time *string `json:"time,omitempty" xml:"time,omitempty" require:"true"`
  // {"en":"Number of total requests", "zh_CN":"总请求数"}
  Hit *int64 `json:"hit,omitempty" xml:"hit,omitempty" require:"true"`
  // {"en":"Number of CC attack", "zh_CN":"CC攻击次数"}
  Hitdeny *int64 `json:"hitdeny,omitempty" xml:"hitdeny,omitempty" require:"true"`
  // {"en":"status code 200: success", "zh_CN":"状态码，成功为200；失败见“错误码”信息"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
}

func (s CcAttackQpsForQueryResponse) String() string {
  return tea.Prettify(s)
}

func (s CcAttackQpsForQueryResponse) GoString() string {
  return s.String()
}

func (s *CcAttackQpsForQueryResponse) SetMsg(v string) *CcAttackQpsForQueryResponse {
  s.Msg = &v
  return s
}

func (s *CcAttackQpsForQueryResponse) SetTime(v string) *CcAttackQpsForQueryResponse {
  s.Time = &v
  return s
}

func (s *CcAttackQpsForQueryResponse) SetHit(v int64) *CcAttackQpsForQueryResponse {
  s.Hit = &v
  return s
}

func (s *CcAttackQpsForQueryResponse) SetHitdeny(v int64) *CcAttackQpsForQueryResponse {
  s.Hitdeny = &v
  return s
}

func (s *CcAttackQpsForQueryResponse) SetCode(v string) *CcAttackQpsForQueryResponse {
  s.Code = &v
  return s
}

type CcAttackQpsForQueryPaths struct {
}

func (s CcAttackQpsForQueryPaths) String() string {
  return tea.Prettify(s)
}

func (s CcAttackQpsForQueryPaths) GoString() string {
  return s.String()
}

type CcAttackQpsForQueryParameters struct {
}

func (s CcAttackQpsForQueryParameters) String() string {
  return tea.Prettify(s)
}

func (s CcAttackQpsForQueryParameters) GoString() string {
  return s.String()
}

type CcAttackQpsForQueryRequestHeader struct {
}

func (s CcAttackQpsForQueryRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s CcAttackQpsForQueryRequestHeader) GoString() string {
  return s.String()
}

type CcAttackQpsForQueryResponseHeader struct {
}

func (s CcAttackQpsForQueryResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s CcAttackQpsForQueryResponseHeader) GoString() string {
  return s.String()
}




type ListAttackLogsRequest struct {
  // {"en":"Start time, format: yyyy-MM-dd HH:mm:ss.\nThe time range can not exceed 31 days.","zh_CN":"开始时间，格式：yyyy-MM-dd HH:mm:ss。\n时间范围不超过31天。"}
  StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty" require:"true"`
  // {"en":"End time, format: yyyy-MM-dd HH:mm:ss.\nThe time range can not exceed 31 days.","zh_CN":"结束时间，格式：yyyy-MM-dd HH:mm:ss。\n时间范围不超过31天。"}
  EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty" require:"true"`
  // {"en":"Domain list. Queries all domains under the account when not specified.","zh_CN":"域名列表，未指定时查询账号下的所有域名。"}
  DomainList []*string `json:"domainList,omitempty" xml:"domainList,omitempty" type:"Repeated"`
  // {"en":"Policy type list.\nDMS_DEFEND: DDoS Protection\nBOT_MANAGE: Bot Management\nAPI_DEFEND: API Security\nWAF_DEFEND: WAF\nBLOCK: IP/Geo Block\nCUSTOMIZE_RULE: Custom Rules\nRATE_LIMIT: Rate Limiting\nINTELLIGENCE: Threat Intelligence","zh_CN":"策略类型列表。\nDMS_DEFEND：DDoS防护\nBOT_MANAGE：Bot管理\nAPI_DEFEND：API安全\nWAF_DEFEND：WAF\nBLOCK：IP/区域封禁\nCUSTOMIZE_RULE：自定义规则\nRATE_LIMIT：频率限制\nINTELLIGENCE：威胁情报","exampleValue":"BLOCK,DMS_DEFEND,WAF_DEFEND,BOT_MANAGE,API_DEFEND,INTELLIGENCE,RATE_LIMIT,CUSTOMIZE_RULE"}
  PolicyTypeList []*string `json:"policyTypeList,omitempty" xml:"policyTypeList,omitempty" type:"Repeated"`
  // {"en":"The number of records per page, default value:10, maximum value 2000.","zh_CN":"每页显示的条目数，默认：10，最大：2000。"}
  PageSize *int `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
  // {"defaultValue":"1","en":"The current page, default value: 1.","zh_CN":"当前第几页，默认：1。"}
  CurrentPage *int `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
}

func (s ListAttackLogsRequest) String() string {
  return tea.Prettify(s)
}

func (s ListAttackLogsRequest) GoString() string {
  return s.String()
}

func (s *ListAttackLogsRequest) SetStartTime(v string) *ListAttackLogsRequest {
  s.StartTime = &v
  return s
}

func (s *ListAttackLogsRequest) SetEndTime(v string) *ListAttackLogsRequest {
  s.EndTime = &v
  return s
}

func (s *ListAttackLogsRequest) SetDomainList(v []*string) *ListAttackLogsRequest {
  s.DomainList = v
  return s
}

func (s *ListAttackLogsRequest) SetPolicyTypeList(v []*string) *ListAttackLogsRequest {
  s.PolicyTypeList = v
  return s
}

func (s *ListAttackLogsRequest) SetPageSize(v int) *ListAttackLogsRequest {
  s.PageSize = &v
  return s
}

func (s *ListAttackLogsRequest) SetCurrentPage(v int) *ListAttackLogsRequest {
  s.CurrentPage = &v
  return s
}

type ListAttackLogsRequestHeader struct {
  // {"defaultValue":"en","en":"The language of response data, default value: en.\nzh_CN: Chinese\nen: English","zh_CN":"返回内容的语言版本，默认值: en。\nzh_CN：中文\nen：英文","exampleValue":"zh_CN,en"}
  Language *string `json:"language,omitempty" xml:"language,omitempty"`
  // {"dictionary":"belong=WAAP-MS-Ext|dict=waap_serviceType","en":"Security service type. Please enter a specific service type, if you purchase multiple security services.","zh_CN":"安全服务类型。有使用多个不同的安全服务时，需要填写具体的服务类型。"}
  ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty"`
  // {"defaultValue":"GMT+8","en":"Report Data Timezone:\r\n\r\n1. Indicates the timezone for the report data. lt must be relative to GMT and\r\nspecified in the format GMT+n where -12<=n <= 12. For example,timezone=GMT-7 and timezone=GMT+8.\r\n\r\n2. If the parameter is unspecified,results will be in the GMT timezone.","zh_CN":"报表数据时区：1.请传递GMT时区。示例格式：GMT+N（其中12<=N <= 12），示例：timezone=GMT-7或timezone=GMT+8.2.若参数为传递，将默认按GMT+8时区查询"}
  Timezone *string `json:"timezone,omitempty" xml:"timezone,omitempty"`
}

func (s ListAttackLogsRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ListAttackLogsRequestHeader) GoString() string {
  return s.String()
}

func (s *ListAttackLogsRequestHeader) SetLanguage(v string) *ListAttackLogsRequestHeader {
  s.Language = &v
  return s
}

func (s *ListAttackLogsRequestHeader) SetServiceType(v string) *ListAttackLogsRequestHeader {
  s.ServiceType = &v
  return s
}

func (s *ListAttackLogsRequestHeader) SetTimezone(v string) *ListAttackLogsRequestHeader {
  s.Timezone = &v
  return s
}

type ListAttackLogsPaths struct {
}

func (s ListAttackLogsPaths) String() string {
  return tea.Prettify(s)
}

func (s ListAttackLogsPaths) GoString() string {
  return s.String()
}

type ListAttackLogsParameters struct {
}

func (s ListAttackLogsParameters) String() string {
  return tea.Prettify(s)
}

func (s ListAttackLogsParameters) GoString() string {
  return s.String()
}

type ListAttackLogsResponse struct {
  // {"dictionary":"belong=WAAP-MS-Ext|dict=waap_retCodeEnum","en":"Please refer to the error code for exceptions.","zh_CN":"请参照错误码。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Description.","zh_CN":"描述信息。"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"Data.","zh_CN":"出参数据。"}
  Data *ListAttackLogsResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
}

func (s ListAttackLogsResponse) String() string {
  return tea.Prettify(s)
}

func (s ListAttackLogsResponse) GoString() string {
  return s.String()
}

func (s *ListAttackLogsResponse) SetCode(v string) *ListAttackLogsResponse {
  s.Code = &v
  return s
}

func (s *ListAttackLogsResponse) SetMsg(v string) *ListAttackLogsResponse {
  s.Msg = &v
  return s
}

func (s *ListAttackLogsResponse) SetData(v *ListAttackLogsResponseData) *ListAttackLogsResponse {
  s.Data = v
  return s
}

type ListAttackLogsResponseData struct {
  // {"en":"Return content.","zh_CN":"返回内容。"}
  List *ListAttackLogsResponseDataList `json:"list,omitempty" xml:"list,omitempty" require:"true" type:"Struct"`
  // {"en":"The current page number.","zh_CN":"当前页码。"}
  PageNum *int64 `json:"pageNum,omitempty" xml:"pageNum,omitempty" require:"true"`
  // {"en":"The number of records per page.","zh_CN":"每页显示的条目数。"}
  PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty" require:"true"`
  // {"en":"The total number of records.","zh_CN":"总记录数。"}
  TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty" require:"true"`
  // {"en":"The total number of pages.","zh_CN":"总页数。"}
  TotalPageCount *int64 `json:"totalPageCount,omitempty" xml:"totalPageCount,omitempty" require:"true"`
}

func (s ListAttackLogsResponseData) String() string {
  return tea.Prettify(s)
}

func (s ListAttackLogsResponseData) GoString() string {
  return s.String()
}

func (s *ListAttackLogsResponseData) SetList(v *ListAttackLogsResponseDataList) *ListAttackLogsResponseData {
  s.List = v
  return s
}

func (s *ListAttackLogsResponseData) SetPageNum(v int64) *ListAttackLogsResponseData {
  s.PageNum = &v
  return s
}

func (s *ListAttackLogsResponseData) SetPageSize(v int64) *ListAttackLogsResponseData {
  s.PageSize = &v
  return s
}

func (s *ListAttackLogsResponseData) SetTotalCount(v int64) *ListAttackLogsResponseData {
  s.TotalCount = &v
  return s
}

func (s *ListAttackLogsResponseData) SetTotalPageCount(v int64) *ListAttackLogsResponseData {
  s.TotalPageCount = &v
  return s
}

type ListAttackLogsResponseDataList struct {
  // {"en":"Attack time, format: yyyy-MM-dd HH:mm:ss.","zh_CN":"攻击时间，格式：yyyy-MM-dd HH:mm:ss。"}
  AttackTime *string `json:"attackTime,omitempty" xml:"attackTime,omitempty" require:"true"`
  // {"en":"Client IP.","zh_CN":"客户端IP。"}
  ClientIp *string `json:"clientIp,omitempty" xml:"clientIp,omitempty" require:"true"`
  // {"en":"Domain.","zh_CN":"域名。"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {"en":"Event ID.","zh_CN":"事件ID。"}
  Uuid *string `json:"uuid,omitempty" xml:"uuid,omitempty" require:"true"`
  // {"en":"Request ID.","zh_CN":"请求ID。"}
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty" require:"true"`
  // {"en":"Policy type.\nDMS_DEFEND: DDoS Protection\nBOT_MANAGE: Bot Management\nAPI_DEFEND: API Security\nWAF_DEFEND: WAF\nBLOCK: IP/Geo Block\nCUSTOMIZE_RULE: Custom Rules\nRATE_LIMIT: Rate Limiting\nINTELLIGENCE: Threat Intelligence","zh_CN":"策略类型。\nDMS_DEFEND：DDoS防护\nBOT_MANAGE：Bot管理\nAPI_DEFEND：API安全\nWAF_DEFEND：WAF\nBLOCK：IP/区域封禁\nCUSTOMIZE_RULE：自定义规则\nRATE_LIMIT：频率限制\nINTELLIGENCE：威胁情报","exampleValue":"BLOCK,DMS_DEFEND,WAF_DEFEND,BOT_MANAGE,API_DEFEND,INTELLIGENCE,RATE_LIMIT,CUSTOMIZE_RULE"}
  PolicyType *string `json:"policyType,omitempty" xml:"policyType,omitempty" require:"true"`
  // {"en":"Path.","zh_CN":"路径。"}
  Path *string `json:"path,omitempty" xml:"path,omitempty" require:"true"`
  // {"en":"Status code.","zh_CN":"状态码。"}
  StatusCode *string `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
  // {"en":"Action.","zh_CN":"处理动作。"}
  Action *string `json:"action,omitempty" xml:"action,omitempty" require:"true"`
}

func (s ListAttackLogsResponseDataList) String() string {
  return tea.Prettify(s)
}

func (s ListAttackLogsResponseDataList) GoString() string {
  return s.String()
}

func (s *ListAttackLogsResponseDataList) SetAttackTime(v string) *ListAttackLogsResponseDataList {
  s.AttackTime = &v
  return s
}

func (s *ListAttackLogsResponseDataList) SetClientIp(v string) *ListAttackLogsResponseDataList {
  s.ClientIp = &v
  return s
}

func (s *ListAttackLogsResponseDataList) SetDomain(v string) *ListAttackLogsResponseDataList {
  s.Domain = &v
  return s
}

func (s *ListAttackLogsResponseDataList) SetUuid(v string) *ListAttackLogsResponseDataList {
  s.Uuid = &v
  return s
}

func (s *ListAttackLogsResponseDataList) SetRequestId(v string) *ListAttackLogsResponseDataList {
  s.RequestId = &v
  return s
}

func (s *ListAttackLogsResponseDataList) SetPolicyType(v string) *ListAttackLogsResponseDataList {
  s.PolicyType = &v
  return s
}

func (s *ListAttackLogsResponseDataList) SetPath(v string) *ListAttackLogsResponseDataList {
  s.Path = &v
  return s
}

func (s *ListAttackLogsResponseDataList) SetStatusCode(v string) *ListAttackLogsResponseDataList {
  s.StatusCode = &v
  return s
}

func (s *ListAttackLogsResponseDataList) SetAction(v string) *ListAttackLogsResponseDataList {
  s.Action = &v
  return s
}

type ListAttackLogsResponseHeader struct {
}

func (s ListAttackLogsResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ListAttackLogsResponseHeader) GoString() string {
  return s.String()
}




type GetTopAttackSourcesForClientIPsRequest struct {
  // {"defaultValue":"10","en":"Top rankings, default value: 10, max: 1000.","zh_CN":"取前几排名，默认：10，上限：1000。"}
  Top *int `json:"top,omitempty" xml:"top,omitempty"`
  // {"en":"Multiple selection. Handling results, default: display all results.\nmitigated: Number of mitigated requests.\nmonitored: Number of observed requests.","zh_CN":"多选。处置结果，默认：展示所有结果。\nmitigated：已抵御请求数。\nmonitored：观察请求数。","exampleValue":"mitigated,monitored"}
  ActType []*string `json:"actType,omitempty" xml:"actType,omitempty" type:"Repeated"`
  // {"en":"Start time, format: yyyy-MM-dd HH:mm:ss.\nThe time range can not exceed 31 days.","zh_CN":"起始时间，格式：yyyy-MM-dd HH:mm:ss。\n时间范围不超过31天。"}
  StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty" require:"true"`
  // {"en":"End time, format: yyyy-MM-dd HH:mm:ss.\nThe time range can not exceed 31 days.","zh_CN":"截止时间，格式：yyyy-MM-dd HH:mm:ss。\n时间范围不超过31天。"}
  EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty" require:"true"`
  // {"en":"Domain list. Queries all domains under the account when not specified.","zh_CN":"域名列表，未指定时查询账号下的所有域名。"}
  Domains []*string `json:"domains,omitempty" xml:"domains,omitempty" type:"Repeated"`
}

func (s GetTopAttackSourcesForClientIPsRequest) String() string {
  return tea.Prettify(s)
}

func (s GetTopAttackSourcesForClientIPsRequest) GoString() string {
  return s.String()
}

func (s *GetTopAttackSourcesForClientIPsRequest) SetTop(v int) *GetTopAttackSourcesForClientIPsRequest {
  s.Top = &v
  return s
}

func (s *GetTopAttackSourcesForClientIPsRequest) SetActType(v []*string) *GetTopAttackSourcesForClientIPsRequest {
  s.ActType = v
  return s
}

func (s *GetTopAttackSourcesForClientIPsRequest) SetStartTime(v string) *GetTopAttackSourcesForClientIPsRequest {
  s.StartTime = &v
  return s
}

func (s *GetTopAttackSourcesForClientIPsRequest) SetEndTime(v string) *GetTopAttackSourcesForClientIPsRequest {
  s.EndTime = &v
  return s
}

func (s *GetTopAttackSourcesForClientIPsRequest) SetDomains(v []*string) *GetTopAttackSourcesForClientIPsRequest {
  s.Domains = v
  return s
}

type GetTopAttackSourcesForClientIPsRequestHeader struct {
  // {"defaultValue":"en","en":"The language of response data, default value: en.\nzh_CN: Chinese\nen: English","zh_CN":"返回内容的语言版本，默认值: en。\nzh_CN：中文\nen：英文","exampleValue":"zh_CN,en"}
  Language *string `json:"language,omitempty" xml:"language,omitempty"`
  // {"dictionary":"belong=WAAP-MS-Ext|dict=waap_serviceType","en":"Security service type. Please enter a specific service type, if you purchase multiple security services.","zh_CN":"安全服务类型。有使用多个不同的安全服务时，需要填写具体的服务类型。"}
  ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty"`
  // {"defaultValue":"GMT+8","en":"Report Data Timezone:1. Indicates the timezone for the report data. lt must be relative to GMT andspecified in the format GMT+n where -12<=n <= 12. For example,timezone=GMT-7 and timezone=GMT+8.2. If the parameter is unspecified,results will be in the GMT+8 timezone.","zh_CN":"报表数据时区：1.请传递GMT时区。示例格式：GMT+N（其中12<=N <= 12），示例：timezone=GMT-7或timezone=GMT+8.2.若参数为传递，将默认按GMT+8时区查询"}
  Timezone *string `json:"timezone,omitempty" xml:"timezone,omitempty"`
}

func (s GetTopAttackSourcesForClientIPsRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetTopAttackSourcesForClientIPsRequestHeader) GoString() string {
  return s.String()
}

func (s *GetTopAttackSourcesForClientIPsRequestHeader) SetLanguage(v string) *GetTopAttackSourcesForClientIPsRequestHeader {
  s.Language = &v
  return s
}

func (s *GetTopAttackSourcesForClientIPsRequestHeader) SetServiceType(v string) *GetTopAttackSourcesForClientIPsRequestHeader {
  s.ServiceType = &v
  return s
}

func (s *GetTopAttackSourcesForClientIPsRequestHeader) SetTimezone(v string) *GetTopAttackSourcesForClientIPsRequestHeader {
  s.Timezone = &v
  return s
}

type GetTopAttackSourcesForClientIPsPaths struct {
}

func (s GetTopAttackSourcesForClientIPsPaths) String() string {
  return tea.Prettify(s)
}

func (s GetTopAttackSourcesForClientIPsPaths) GoString() string {
  return s.String()
}

type GetTopAttackSourcesForClientIPsParameters struct {
}

func (s GetTopAttackSourcesForClientIPsParameters) String() string {
  return tea.Prettify(s)
}

func (s GetTopAttackSourcesForClientIPsParameters) GoString() string {
  return s.String()
}

type GetTopAttackSourcesForClientIPsResponse struct {
  // {"dictionary":"belong=WAAP-MS-Ext|dict=waap_retCodeEnum","en":"Please refer to the error code for exceptions.","zh_CN":"请参照错误码。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Description.","zh_CN":"描述信息。"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"Data.","zh_CN":"出参数据。"}
  Data []*GetTopAttackSourcesForClientIPsResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s GetTopAttackSourcesForClientIPsResponse) String() string {
  return tea.Prettify(s)
}

func (s GetTopAttackSourcesForClientIPsResponse) GoString() string {
  return s.String()
}

func (s *GetTopAttackSourcesForClientIPsResponse) SetCode(v string) *GetTopAttackSourcesForClientIPsResponse {
  s.Code = &v
  return s
}

func (s *GetTopAttackSourcesForClientIPsResponse) SetMsg(v string) *GetTopAttackSourcesForClientIPsResponse {
  s.Msg = &v
  return s
}

func (s *GetTopAttackSourcesForClientIPsResponse) SetData(v []*GetTopAttackSourcesForClientIPsResponseData) *GetTopAttackSourcesForClientIPsResponse {
  s.Data = v
  return s
}

type GetTopAttackSourcesForClientIPsResponseData struct     {
  // {"en":"Source area(country or area, province).","zh_CN":"来源地区（国家或地区、省份）。"}
  Area *string `json:"area,omitempty" xml:"area,omitempty" require:"true"`
  // {"en":"Attack requests.","zh_CN":"攻击请求数。"}
  Value *int64 `json:"value,omitempty" xml:"value,omitempty" require:"true"`
  // {"en":"Client IP.","zh_CN":"来源IP。"}
  Ip *string `json:"ip,omitempty" xml:"ip,omitempty" require:"true"`
}

func (s GetTopAttackSourcesForClientIPsResponseData) String() string {
  return tea.Prettify(s)
}

func (s GetTopAttackSourcesForClientIPsResponseData) GoString() string {
  return s.String()
}

func (s *GetTopAttackSourcesForClientIPsResponseData) SetArea(v string) *GetTopAttackSourcesForClientIPsResponseData {
  s.Area = &v
  return s
}

func (s *GetTopAttackSourcesForClientIPsResponseData) SetValue(v int64) *GetTopAttackSourcesForClientIPsResponseData {
  s.Value = &v
  return s
}

func (s *GetTopAttackSourcesForClientIPsResponseData) SetIp(v string) *GetTopAttackSourcesForClientIPsResponseData {
  s.Ip = &v
  return s
}

type GetTopAttackSourcesForClientIPsResponseHeader struct {
}

func (s GetTopAttackSourcesForClientIPsResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetTopAttackSourcesForClientIPsResponseHeader) GoString() string {
  return s.String()
}




type GetSummaryRequestsRequest struct {
  // {"en":"Start time, format: yyyy-MM-dd HH:mm:ss.\nThe time range can not exceed 31 days.","zh_CN":"起始时间，格式：yyyy-MM-dd HH:mm:ss。\n时间范围不超过31天。"}
  StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty" require:"true"`
  // {"en":"End time, format: yyyy-MM-dd HH:mm:ss.\nThe time range can not exceed 31 days.","zh_CN":"截止时间，格式：yyyy-MM-dd HH:mm:ss。\n时间范围不超过31天。"}
  EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty" require:"true"`
  // {"en":"Domain list. Queries all domains under the account when not specified.","zh_CN":"域名列表，未指定时查询账号下的所有域名。"}
  Domains []*string `json:"domains,omitempty" xml:"domains,omitempty" type:"Repeated"`
}

func (s GetSummaryRequestsRequest) String() string {
  return tea.Prettify(s)
}

func (s GetSummaryRequestsRequest) GoString() string {
  return s.String()
}

func (s *GetSummaryRequestsRequest) SetStartTime(v string) *GetSummaryRequestsRequest {
  s.StartTime = &v
  return s
}

func (s *GetSummaryRequestsRequest) SetEndTime(v string) *GetSummaryRequestsRequest {
  s.EndTime = &v
  return s
}

func (s *GetSummaryRequestsRequest) SetDomains(v []*string) *GetSummaryRequestsRequest {
  s.Domains = v
  return s
}

type GetSummaryRequestsRequestHeader struct {
  // {"dictionary":"belong=WAAP-MS-Ext|dict=waap_serviceType","en":"Security service type. Please enter a specific service type, if you purchase multiple security services.","zh_CN":"安全服务类型。有使用多个不同的安全服务时，需要填写具体的服务类型。"}
  ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty"`
  // {"defaultValue":"GMT+8","en":"Report Data Timezone:1. Indicates the timezone for the report data. lt must be relative to GMT andspecified in the format GMT+n where -12<=n <= 12. For example,timezone=GMT-7 and timezone=GMT+8.2. If the parameter is unspecified,results will be in the GMT+8 timezone.","zh_CN":"报表数据时区：1.请传递GMT时区。示例格式：GMT+N（其中12<=N <= 12），示例：timezone=GMT-7或timezone=GMT+8.2.若参数为传递，将默认按GMT+8时区查询"}
  Timezone *string `json:"timezone,omitempty" xml:"timezone,omitempty"`
}

func (s GetSummaryRequestsRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetSummaryRequestsRequestHeader) GoString() string {
  return s.String()
}

func (s *GetSummaryRequestsRequestHeader) SetServiceType(v string) *GetSummaryRequestsRequestHeader {
  s.ServiceType = &v
  return s
}

func (s *GetSummaryRequestsRequestHeader) SetTimezone(v string) *GetSummaryRequestsRequestHeader {
  s.Timezone = &v
  return s
}

type GetSummaryRequestsPaths struct {
}

func (s GetSummaryRequestsPaths) String() string {
  return tea.Prettify(s)
}

func (s GetSummaryRequestsPaths) GoString() string {
  return s.String()
}

type GetSummaryRequestsParameters struct {
}

func (s GetSummaryRequestsParameters) String() string {
  return tea.Prettify(s)
}

func (s GetSummaryRequestsParameters) GoString() string {
  return s.String()
}

type GetSummaryRequestsResponse struct {
  // {"dictionary":"belong=WAAP-MS-Ext|dict=waap_retCodeEnum","en":"Please refer to the error code for exceptions.","zh_CN":"请参照错误码。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Description.","zh_CN":"描述信息。"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"Data.","zh_CN":"出参数据。"}
  Data *GetSummaryRequestsResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
}

func (s GetSummaryRequestsResponse) String() string {
  return tea.Prettify(s)
}

func (s GetSummaryRequestsResponse) GoString() string {
  return s.String()
}

func (s *GetSummaryRequestsResponse) SetCode(v string) *GetSummaryRequestsResponse {
  s.Code = &v
  return s
}

func (s *GetSummaryRequestsResponse) SetMsg(v string) *GetSummaryRequestsResponse {
  s.Msg = &v
  return s
}

func (s *GetSummaryRequestsResponse) SetData(v *GetSummaryRequestsResponseData) *GetSummaryRequestsResponse {
  s.Data = v
  return s
}

type GetSummaryRequestsResponseData struct {
  // {"en":"Total requests.","zh_CN":"总请求数。"}
  Total *int64 `json:"total,omitempty" xml:"total,omitempty" require:"true"`
  // {"en":"Attack requests,sum of resisted requests and observed requests.","zh_CN":"攻击请求数，已抵御请求数与观察请求数之和。"}
  Attack *int64 `json:"attack,omitempty" xml:"attack,omitempty" require:"true"`
  // {"en":"Whitelist requests.","zh_CN":"白名单请求数。"}
  Whitelist *int64 `json:"whitelist,omitempty" xml:"whitelist,omitempty" require:"true"`
  // {"en":"Resisted requests.","zh_CN":"已抵御请求数。"}
  Resisted *int64 `json:"resisted,omitempty" xml:"resisted,omitempty" require:"true"`
  // {"en":"Observed requests.","zh_CN":"观察请求数。"}
  Observed *int64 `json:"observed,omitempty" xml:"observed,omitempty" require:"true"`
}

func (s GetSummaryRequestsResponseData) String() string {
  return tea.Prettify(s)
}

func (s GetSummaryRequestsResponseData) GoString() string {
  return s.String()
}

func (s *GetSummaryRequestsResponseData) SetTotal(v int64) *GetSummaryRequestsResponseData {
  s.Total = &v
  return s
}

func (s *GetSummaryRequestsResponseData) SetAttack(v int64) *GetSummaryRequestsResponseData {
  s.Attack = &v
  return s
}

func (s *GetSummaryRequestsResponseData) SetWhitelist(v int64) *GetSummaryRequestsResponseData {
  s.Whitelist = &v
  return s
}

func (s *GetSummaryRequestsResponseData) SetResisted(v int64) *GetSummaryRequestsResponseData {
  s.Resisted = &v
  return s
}

func (s *GetSummaryRequestsResponseData) SetObserved(v int64) *GetSummaryRequestsResponseData {
  s.Observed = &v
  return s
}

type GetSummaryRequestsResponseHeader struct {
}

func (s GetSummaryRequestsResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetSummaryRequestsResponseHeader) GoString() string {
  return s.String()
}




type GetActTypeDistributionDataRequest struct {
  // {"en":"Domain.Separate by ';'.", "zh_CN":"域名。多个以;隔开"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {"en":"StartTime.Format: yyyy-MM-dd HH:mm:ss", "zh_CN":"开始时间。格式： yyyy-MM-dd HH:mm:ss"}
  StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty" require:"true"`
  // {"en":"EndTime. format: yyyy-MM-dd HH:mm:ss", "zh_CN":"结束时间。格式： yyyy-MM-dd HH:mm:ss"}
  EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty" require:"true"`
  // {"en":"Time zone. Default 8, i.e.'GTM+8'", "zh_CN":"时区，默认8，即“GTM+8”"}
  TimeZone *int `json:"timeZone,omitempty" xml:"timeZone,omitempty"`
  // {"en":"Language type. Default cn. 
  //  en:English 
  //  cn:Chinese ", "zh_CN":"语言类型。 默认cn 
  //  en：英文 
  //  cn：中文"}
  Lang *string `json:"lang,omitempty" xml:"lang,omitempty"`
}

func (s GetActTypeDistributionDataRequest) String() string {
  return tea.Prettify(s)
}

func (s GetActTypeDistributionDataRequest) GoString() string {
  return s.String()
}

func (s *GetActTypeDistributionDataRequest) SetDomain(v string) *GetActTypeDistributionDataRequest {
  s.Domain = &v
  return s
}

func (s *GetActTypeDistributionDataRequest) SetStartTime(v string) *GetActTypeDistributionDataRequest {
  s.StartTime = &v
  return s
}

func (s *GetActTypeDistributionDataRequest) SetEndTime(v string) *GetActTypeDistributionDataRequest {
  s.EndTime = &v
  return s
}

func (s *GetActTypeDistributionDataRequest) SetTimeZone(v int) *GetActTypeDistributionDataRequest {
  s.TimeZone = &v
  return s
}

func (s *GetActTypeDistributionDataRequest) SetLang(v string) *GetActTypeDistributionDataRequest {
  s.Lang = &v
  return s
}

type GetActTypeDistributionDataResponse struct {
  // {"en":"Status code, success is '200'.", "zh_CN":"状态码，成功为“200”。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Return message, success is 'Success'.", "zh_CN":"返回信息，成功为“Success”。"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Data returned", "zh_CN":"返回数据"}
  Data []*GetActTypeDistributionDataResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s GetActTypeDistributionDataResponse) String() string {
  return tea.Prettify(s)
}

func (s GetActTypeDistributionDataResponse) GoString() string {
  return s.String()
}

func (s *GetActTypeDistributionDataResponse) SetCode(v string) *GetActTypeDistributionDataResponse {
  s.Code = &v
  return s
}

func (s *GetActTypeDistributionDataResponse) SetMessage(v string) *GetActTypeDistributionDataResponse {
  s.Message = &v
  return s
}

func (s *GetActTypeDistributionDataResponse) SetData(v []*GetActTypeDistributionDataResponseData) *GetActTypeDistributionDataResponse {
  s.Data = v
  return s
}

type GetActTypeDistributionDataResponseData struct     {
  // {"en":"Action status.", "zh_CN":"处置状态。"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"Action times.", "zh_CN":"处置次数。"}
  Value *int64 `json:"value,omitempty" xml:"value,omitempty" require:"true"`
}

func (s GetActTypeDistributionDataResponseData) String() string {
  return tea.Prettify(s)
}

func (s GetActTypeDistributionDataResponseData) GoString() string {
  return s.String()
}

func (s *GetActTypeDistributionDataResponseData) SetName(v string) *GetActTypeDistributionDataResponseData {
  s.Name = &v
  return s
}

func (s *GetActTypeDistributionDataResponseData) SetValue(v int64) *GetActTypeDistributionDataResponseData {
  s.Value = &v
  return s
}

type GetActTypeDistributionDataPaths struct {
}

func (s GetActTypeDistributionDataPaths) String() string {
  return tea.Prettify(s)
}

func (s GetActTypeDistributionDataPaths) GoString() string {
  return s.String()
}

type GetActTypeDistributionDataParameters struct {
}

func (s GetActTypeDistributionDataParameters) String() string {
  return tea.Prettify(s)
}

func (s GetActTypeDistributionDataParameters) GoString() string {
  return s.String()
}

type GetActTypeDistributionDataRequestHeader struct {
}

func (s GetActTypeDistributionDataRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetActTypeDistributionDataRequestHeader) GoString() string {
  return s.String()
}

type GetActTypeDistributionDataResponseHeader struct {
}

func (s GetActTypeDistributionDataResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetActTypeDistributionDataResponseHeader) GoString() string {
  return s.String()
}




type GetBotRuleTypeTopDataRequest struct {
  // {"en":"Domain.Separate by ';'.", "zh_CN":"域名。多个以;隔开"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {"en":"StartTime.Format:yyyy-MM-dd HH:mm:ss", "zh_CN":"开始时间。格式： yyyy-MM-dd HH:mm:ss"}
  StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty" require:"true"`
  // {"en":"EndTime. Format: yyyy-MM-dd HH:mm:ss", "zh_CN":"结束时间。格式： yyyy-MM-dd HH:mm:ss"}
  EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty" require:"true"`
  // {"en":"Time zone. Default 8, i.e.'GTM+8'.", "zh_CN":"时区，默认8，即“GTM+8”"}
  TimeZone *int `json:"timeZone,omitempty" xml:"timeZone,omitempty"`
  // {"en":"Language type. Default cn. 
  //  en:English 
  //  cn:Chinese ", "zh_CN":"语言类型。 默认cn 
  //  en：英文 
  //  cn：中文"}
  Lang *string `json:"lang,omitempty" xml:"lang,omitempty"`
}

func (s GetBotRuleTypeTopDataRequest) String() string {
  return tea.Prettify(s)
}

func (s GetBotRuleTypeTopDataRequest) GoString() string {
  return s.String()
}

func (s *GetBotRuleTypeTopDataRequest) SetDomain(v string) *GetBotRuleTypeTopDataRequest {
  s.Domain = &v
  return s
}

func (s *GetBotRuleTypeTopDataRequest) SetStartTime(v string) *GetBotRuleTypeTopDataRequest {
  s.StartTime = &v
  return s
}

func (s *GetBotRuleTypeTopDataRequest) SetEndTime(v string) *GetBotRuleTypeTopDataRequest {
  s.EndTime = &v
  return s
}

func (s *GetBotRuleTypeTopDataRequest) SetTimeZone(v int) *GetBotRuleTypeTopDataRequest {
  s.TimeZone = &v
  return s
}

func (s *GetBotRuleTypeTopDataRequest) SetLang(v string) *GetBotRuleTypeTopDataRequest {
  s.Lang = &v
  return s
}

type GetBotRuleTypeTopDataResponse struct {
  // {"en":"Status code, success is '200'.", "zh_CN":"状态码，成功为“200”。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Return message, success is 'Success'.", "zh_CN":"返回信息，成功为“Success”。"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Data returned.", "zh_CN":"返回数据"}
  Data []*GetBotRuleTypeTopDataResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s GetBotRuleTypeTopDataResponse) String() string {
  return tea.Prettify(s)
}

func (s GetBotRuleTypeTopDataResponse) GoString() string {
  return s.String()
}

func (s *GetBotRuleTypeTopDataResponse) SetCode(v string) *GetBotRuleTypeTopDataResponse {
  s.Code = &v
  return s
}

func (s *GetBotRuleTypeTopDataResponse) SetMessage(v string) *GetBotRuleTypeTopDataResponse {
  s.Message = &v
  return s
}

func (s *GetBotRuleTypeTopDataResponse) SetData(v []*GetBotRuleTypeTopDataResponseData) *GetBotRuleTypeTopDataResponse {
  s.Data = v
  return s
}

type GetBotRuleTypeTopDataResponseData struct     {
  // {"en":"Rule name.", "zh_CN":"规则名。"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"Trigger Times.", "zh_CN":"触发次数。"}
  Value *int64 `json:"value,omitempty" xml:"value,omitempty" require:"true"`
}

func (s GetBotRuleTypeTopDataResponseData) String() string {
  return tea.Prettify(s)
}

func (s GetBotRuleTypeTopDataResponseData) GoString() string {
  return s.String()
}

func (s *GetBotRuleTypeTopDataResponseData) SetName(v string) *GetBotRuleTypeTopDataResponseData {
  s.Name = &v
  return s
}

func (s *GetBotRuleTypeTopDataResponseData) SetValue(v int64) *GetBotRuleTypeTopDataResponseData {
  s.Value = &v
  return s
}

type GetBotRuleTypeTopDataPaths struct {
}

func (s GetBotRuleTypeTopDataPaths) String() string {
  return tea.Prettify(s)
}

func (s GetBotRuleTypeTopDataPaths) GoString() string {
  return s.String()
}

type GetBotRuleTypeTopDataParameters struct {
}

func (s GetBotRuleTypeTopDataParameters) String() string {
  return tea.Prettify(s)
}

func (s GetBotRuleTypeTopDataParameters) GoString() string {
  return s.String()
}

type GetBotRuleTypeTopDataRequestHeader struct {
}

func (s GetBotRuleTypeTopDataRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetBotRuleTypeTopDataRequestHeader) GoString() string {
  return s.String()
}

type GetBotRuleTypeTopDataResponseHeader struct {
}

func (s GetBotRuleTypeTopDataResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetBotRuleTypeTopDataResponseHeader) GoString() string {
  return s.String()
}




type GetWAFAttackTrendRequest struct {
  // {"en":"Domain, array.", "zh_CN":"域名，数组。"}
  Domains []*string `json:"domains,omitempty" xml:"domains,omitempty" require:"true" type:"Repeated"`
  // {"en":"Start time, yyyy-MM-dd HH:mm:ss.", "zh_CN":"开始时间，yyyy-MM-dd HH:mm:ss。"}
  StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty" require:"true"`
  // {"en":"End time, yyyy-MM-dd HH:mm:ss.", "zh_CN":"结束时间，yyyy-MM-dd HH:mm:ss。"}
  EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty" require:"true"`
  // {"en":"Time zone, GMT+8 by default.", "zh_CN":"时区，默认GMT+8，即“GMT+8”。"}
  TimeZone *string `json:"timeZone,omitempty" xml:"timeZone,omitempty"`
  // {"en":"Attack trend granularity, default value: 5m.
  // 5m: 5 minutes
  // 1h: hour
  // 1d: day", "zh_CN":"攻击趋势粒度，默认值：5m。
  // 5m：5分钟
  // 1h：小时
  // 1d：天"}
  IntervalExpression *string `json:"intervalExpression,omitempty" xml:"intervalExpression,omitempty"`
}

func (s GetWAFAttackTrendRequest) String() string {
  return tea.Prettify(s)
}

func (s GetWAFAttackTrendRequest) GoString() string {
  return s.String()
}

func (s *GetWAFAttackTrendRequest) SetDomains(v []*string) *GetWAFAttackTrendRequest {
  s.Domains = v
  return s
}

func (s *GetWAFAttackTrendRequest) SetStartTime(v string) *GetWAFAttackTrendRequest {
  s.StartTime = &v
  return s
}

func (s *GetWAFAttackTrendRequest) SetEndTime(v string) *GetWAFAttackTrendRequest {
  s.EndTime = &v
  return s
}

func (s *GetWAFAttackTrendRequest) SetTimeZone(v string) *GetWAFAttackTrendRequest {
  s.TimeZone = &v
  return s
}

func (s *GetWAFAttackTrendRequest) SetIntervalExpression(v string) *GetWAFAttackTrendRequest {
  s.IntervalExpression = &v
  return s
}

type GetWAFAttackTrendResponse struct {
  // {"en":"Return 0 means success, please see <Error code> to check other status code.", "zh_CN":"0状态码表示请求成功，其他状态码说明请参见《错误码》。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Error message or Success.", "zh_CN":"错误信息或Success。"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Return data.", "zh_CN":"返回值。"}
  Data *GetWAFAttackTrendAttackRequestTrend `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s GetWAFAttackTrendResponse) String() string {
  return tea.Prettify(s)
}

func (s GetWAFAttackTrendResponse) GoString() string {
  return s.String()
}

func (s *GetWAFAttackTrendResponse) SetCode(v string) *GetWAFAttackTrendResponse {
  s.Code = &v
  return s
}

func (s *GetWAFAttackTrendResponse) SetMessage(v string) *GetWAFAttackTrendResponse {
  s.Message = &v
  return s
}

func (s *GetWAFAttackTrendResponse) SetData(v *GetWAFAttackTrendAttackRequestTrend) *GetWAFAttackTrendResponse {
  s.Data = v
  return s
}

type GetWAFAttackTrendAttackRequestTrend struct {
  // {"en":"Total requests", "zh_CN":"检测请求数。"}
  Hit []*GetWAFAttackTrendHitTrend `json:"hit,omitempty" xml:"hit,omitempty" require:"true" type:"Repeated"`
  // {"en":"Attack requests", "zh_CN":"攻击请求数。"}
  Attack []*GetWAFAttackTrendAttackTrend `json:"attack,omitempty" xml:"attack,omitempty" require:"true" type:"Repeated"`
  // {"en":"Web rules", "zh_CN":"Web规则防护。"}
  Rule []*GetWAFAttackTrendRuleTrend `json:"rule,omitempty" xml:"rule,omitempty" require:"true" type:"Repeated"`
  // {"en":"Protocol validation", "zh_CN":"协议合规检测。"}
  Protocol []*GetWAFAttackTrendProtocolTrend `json:"protocol,omitempty" xml:"protocol,omitempty" require:"true" type:"Repeated"`
  // {"en":"Webshell access detection", "zh_CN":"后门识别。"}
  WebShell []*GetWAFAttackTrendWebShellTrend `json:"webShell,omitempty" xml:"webShell,omitempty" require:"true" type:"Repeated"`
  // {"en":"Access control/Rate limiting", "zh_CN":"访问控制/限速。"}
  Access []*GetWAFAttackTrendAccessTrend `json:"access,omitempty" xml:"access,omitempty" require:"true" type:"Repeated"`
  // {"en":"Other rules。", "zh_CN":"其他防护规则。"}
  Other []*GetWAFAttackTrendOtherTrend `json:"other,omitempty" xml:"other,omitempty" require:"true" type:"Repeated"`
}

func (s GetWAFAttackTrendAttackRequestTrend) String() string {
  return tea.Prettify(s)
}

func (s GetWAFAttackTrendAttackRequestTrend) GoString() string {
  return s.String()
}

func (s *GetWAFAttackTrendAttackRequestTrend) SetHit(v []*GetWAFAttackTrendHitTrend) *GetWAFAttackTrendAttackRequestTrend {
  s.Hit = v
  return s
}

func (s *GetWAFAttackTrendAttackRequestTrend) SetAttack(v []*GetWAFAttackTrendAttackTrend) *GetWAFAttackTrendAttackRequestTrend {
  s.Attack = v
  return s
}

func (s *GetWAFAttackTrendAttackRequestTrend) SetRule(v []*GetWAFAttackTrendRuleTrend) *GetWAFAttackTrendAttackRequestTrend {
  s.Rule = v
  return s
}

func (s *GetWAFAttackTrendAttackRequestTrend) SetProtocol(v []*GetWAFAttackTrendProtocolTrend) *GetWAFAttackTrendAttackRequestTrend {
  s.Protocol = v
  return s
}

func (s *GetWAFAttackTrendAttackRequestTrend) SetWebShell(v []*GetWAFAttackTrendWebShellTrend) *GetWAFAttackTrendAttackRequestTrend {
  s.WebShell = v
  return s
}

func (s *GetWAFAttackTrendAttackRequestTrend) SetAccess(v []*GetWAFAttackTrendAccessTrend) *GetWAFAttackTrendAttackRequestTrend {
  s.Access = v
  return s
}

func (s *GetWAFAttackTrendAttackRequestTrend) SetOther(v []*GetWAFAttackTrendOtherTrend) *GetWAFAttackTrendAttackRequestTrend {
  s.Other = v
  return s
}

type GetWAFAttackTrendHitTrend struct {
  // {"en":"Time, yyyy-MM-dd HH:mm:ss.", "zh_CN":"时间，yyyy-MM-dd HH:mm:ss。"}
  Time *string `json:"time,omitempty" xml:"time,omitempty" require:"true"`
  // {"en":"Total requests.", "zh_CN":"检测请求数。"}
  Text *string `json:"text,omitempty" xml:"text,omitempty" require:"true"`
}

func (s GetWAFAttackTrendHitTrend) String() string {
  return tea.Prettify(s)
}

func (s GetWAFAttackTrendHitTrend) GoString() string {
  return s.String()
}

func (s *GetWAFAttackTrendHitTrend) SetTime(v string) *GetWAFAttackTrendHitTrend {
  s.Time = &v
  return s
}

func (s *GetWAFAttackTrendHitTrend) SetText(v string) *GetWAFAttackTrendHitTrend {
  s.Text = &v
  return s
}

type GetWAFAttackTrendAttackTrend struct {
  // {"en":"Time, yyyy-MM-dd HH:mm:ss.", "zh_CN":"时间，yyyy-MM-dd HH:mm:ss。"}
  Time *string `json:"time,omitempty" xml:"time,omitempty" require:"true"`
  // {"en":"Attack requests.", "zh_CN":"攻击请求数。"}
  Text *string `json:"text,omitempty" xml:"text,omitempty" require:"true"`
}

func (s GetWAFAttackTrendAttackTrend) String() string {
  return tea.Prettify(s)
}

func (s GetWAFAttackTrendAttackTrend) GoString() string {
  return s.String()
}

func (s *GetWAFAttackTrendAttackTrend) SetTime(v string) *GetWAFAttackTrendAttackTrend {
  s.Time = &v
  return s
}

func (s *GetWAFAttackTrendAttackTrend) SetText(v string) *GetWAFAttackTrendAttackTrend {
  s.Text = &v
  return s
}

type GetWAFAttackTrendRuleTrend struct {
  // {"en":"Time, yyyy-MM-dd HH:mm:ss.", "zh_CN":"时间，yyyy-MM-dd HH:mm:ss。"}
  Time *string `json:"time,omitempty" xml:"time,omitempty" require:"true"`
  // {"en":"Web rules.", "zh_CN":"Web规则防护。"}
  Text *string `json:"text,omitempty" xml:"text,omitempty" require:"true"`
}

func (s GetWAFAttackTrendRuleTrend) String() string {
  return tea.Prettify(s)
}

func (s GetWAFAttackTrendRuleTrend) GoString() string {
  return s.String()
}

func (s *GetWAFAttackTrendRuleTrend) SetTime(v string) *GetWAFAttackTrendRuleTrend {
  s.Time = &v
  return s
}

func (s *GetWAFAttackTrendRuleTrend) SetText(v string) *GetWAFAttackTrendRuleTrend {
  s.Text = &v
  return s
}

type GetWAFAttackTrendProtocolTrend struct {
  // {"en":"Time, yyyy-MM-dd HH:mm:ss.", "zh_CN":"时间，yyyy-MM-dd HH:mm:ss。"}
  Time *string `json:"time,omitempty" xml:"time,omitempty" require:"true"`
  // {"en":"Protocol validation.", "zh_CN":"协议合规检测。"}
  Text *string `json:"text,omitempty" xml:"text,omitempty" require:"true"`
}

func (s GetWAFAttackTrendProtocolTrend) String() string {
  return tea.Prettify(s)
}

func (s GetWAFAttackTrendProtocolTrend) GoString() string {
  return s.String()
}

func (s *GetWAFAttackTrendProtocolTrend) SetTime(v string) *GetWAFAttackTrendProtocolTrend {
  s.Time = &v
  return s
}

func (s *GetWAFAttackTrendProtocolTrend) SetText(v string) *GetWAFAttackTrendProtocolTrend {
  s.Text = &v
  return s
}

type GetWAFAttackTrendWebShellTrend struct {
  // {"en":"Time, yyyy-MM-dd HH:mm:ss.", "zh_CN":"时间，yyyy-MM-dd HH:mm:ss。"}
  Time *string `json:"time,omitempty" xml:"time,omitempty" require:"true"`
  // {"en":"Webshell access detection.", "zh_CN":"后门识别。"}
  Text *string `json:"text,omitempty" xml:"text,omitempty" require:"true"`
}

func (s GetWAFAttackTrendWebShellTrend) String() string {
  return tea.Prettify(s)
}

func (s GetWAFAttackTrendWebShellTrend) GoString() string {
  return s.String()
}

func (s *GetWAFAttackTrendWebShellTrend) SetTime(v string) *GetWAFAttackTrendWebShellTrend {
  s.Time = &v
  return s
}

func (s *GetWAFAttackTrendWebShellTrend) SetText(v string) *GetWAFAttackTrendWebShellTrend {
  s.Text = &v
  return s
}

type GetWAFAttackTrendAccessTrend struct {
  // {"en":"Time, yyyy-MM-dd HH:mm:ss.", "zh_CN":"时间，yyyy-MM-dd HH:mm:ss。"}
  Time *string `json:"time,omitempty" xml:"time,omitempty" require:"true"`
  // {"en":"Access control/Rate limiting.", "zh_CN":"访问控制/限速。"}
  Text *string `json:"text,omitempty" xml:"text,omitempty" require:"true"`
}

func (s GetWAFAttackTrendAccessTrend) String() string {
  return tea.Prettify(s)
}

func (s GetWAFAttackTrendAccessTrend) GoString() string {
  return s.String()
}

func (s *GetWAFAttackTrendAccessTrend) SetTime(v string) *GetWAFAttackTrendAccessTrend {
  s.Time = &v
  return s
}

func (s *GetWAFAttackTrendAccessTrend) SetText(v string) *GetWAFAttackTrendAccessTrend {
  s.Text = &v
  return s
}

type GetWAFAttackTrendOtherTrend struct {
  // {"en":"Time, yyyy-MM-dd HH:mm:ss.", "zh_CN":"时间，yyyy-MM-dd HH:mm:ss。"}
  Time *string `json:"time,omitempty" xml:"time,omitempty" require:"true"`
  // {"en":"Other rules.", "zh_CN":"其他防护规则。"}
  Text *string `json:"text,omitempty" xml:"text,omitempty" require:"true"`
}

func (s GetWAFAttackTrendOtherTrend) String() string {
  return tea.Prettify(s)
}

func (s GetWAFAttackTrendOtherTrend) GoString() string {
  return s.String()
}

func (s *GetWAFAttackTrendOtherTrend) SetTime(v string) *GetWAFAttackTrendOtherTrend {
  s.Time = &v
  return s
}

func (s *GetWAFAttackTrendOtherTrend) SetText(v string) *GetWAFAttackTrendOtherTrend {
  s.Text = &v
  return s
}

type GetWAFAttackTrendPaths struct {
}

func (s GetWAFAttackTrendPaths) String() string {
  return tea.Prettify(s)
}

func (s GetWAFAttackTrendPaths) GoString() string {
  return s.String()
}

type GetWAFAttackTrendParameters struct {
}

func (s GetWAFAttackTrendParameters) String() string {
  return tea.Prettify(s)
}

func (s GetWAFAttackTrendParameters) GoString() string {
  return s.String()
}

type GetWAFAttackTrendRequestHeader struct {
}

func (s GetWAFAttackTrendRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetWAFAttackTrendRequestHeader) GoString() string {
  return s.String()
}

type GetWAFAttackTrendResponseHeader struct {
}

func (s GetWAFAttackTrendResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetWAFAttackTrendResponseHeader) GoString() string {
  return s.String()
}




type QueryDDoSAttackDetailsRequest struct {
  // {"en":"Start date format yyyy-MM-dd HH:mm:ss","zh_CN":"起始日期 格式 yyyy-MM-dd HH:mm:ss"}
  Startdate *string `json:"startdate,omitempty" xml:"startdate,omitempty" require:"true"`
  // {"en":"End date format yyyy-MM-dd HH:mm:ss","zh_CN":"结束日期 格式 yyyy-MM-dd HH:mm:ss"}
  Enddate *string `json:"enddate,omitempty" xml:"enddate,omitempty" require:"true"`
  // {"en":"Time zone, such as 28800000. The default is 28800000","zh_CN":"时区，如：28800000 默认为：28800000"}
  Msec *string `json:"msec,omitempty" xml:"msec,omitempty"`
  // {"en":"Business type: dms by default","zh_CN":"业务类型,默认为dms"}
  AcceType *string `json:"acceType,omitempty" xml:"acceType,omitempty"`
}

func (s QueryDDoSAttackDetailsRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryDDoSAttackDetailsRequest) GoString() string {
  return s.String()
}

func (s *QueryDDoSAttackDetailsRequest) SetStartdate(v string) *QueryDDoSAttackDetailsRequest {
  s.Startdate = &v
  return s
}

func (s *QueryDDoSAttackDetailsRequest) SetEnddate(v string) *QueryDDoSAttackDetailsRequest {
  s.Enddate = &v
  return s
}

func (s *QueryDDoSAttackDetailsRequest) SetMsec(v string) *QueryDDoSAttackDetailsRequest {
  s.Msec = &v
  return s
}

func (s *QueryDDoSAttackDetailsRequest) SetAcceType(v string) *QueryDDoSAttackDetailsRequest {
  s.AcceType = &v
  return s
}

type QueryDDoSAttackDetailsRequestHeader struct {
}

func (s QueryDDoSAttackDetailsRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryDDoSAttackDetailsRequestHeader) GoString() string {
  return s.String()
}

type QueryDDoSAttackDetailsPaths struct {
}

func (s QueryDDoSAttackDetailsPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryDDoSAttackDetailsPaths) GoString() string {
  return s.String()
}

type QueryDDoSAttackDetailsParameters struct {
}

func (s QueryDDoSAttackDetailsParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryDDoSAttackDetailsParameters) GoString() string {
  return s.String()
}

type QueryDDoSAttackDetailsResponse struct {
  // {"en":"error response message","zh_CN":"错误响应信息"}
  ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty" require:"true"`
  // {"en":"result","zh_CN":"结果"}
  Data []*QueryDDoSAttackDetailsResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
  // {"en":"error response code","zh_CN":"错误响应码"}
  ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty" require:"true"`
  // {"en":"response code","zh_CN":"响应码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
}

func (s QueryDDoSAttackDetailsResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryDDoSAttackDetailsResponse) GoString() string {
  return s.String()
}

func (s *QueryDDoSAttackDetailsResponse) SetErrorMessage(v string) *QueryDDoSAttackDetailsResponse {
  s.ErrorMessage = &v
  return s
}

func (s *QueryDDoSAttackDetailsResponse) SetData(v []*QueryDDoSAttackDetailsResponseData) *QueryDDoSAttackDetailsResponse {
  s.Data = v
  return s
}

func (s *QueryDDoSAttackDetailsResponse) SetErrorCode(v string) *QueryDDoSAttackDetailsResponse {
  s.ErrorCode = &v
  return s
}

func (s *QueryDDoSAttackDetailsResponse) SetCode(v string) *QueryDDoSAttackDetailsResponse {
  s.Code = &v
  return s
}

type QueryDDoSAttackDetailsResponseData struct     {
  // {"en":"attack peak value","zh_CN":"攻击峰值"}
  TotalFlow *string `json:"totalFlow,omitempty" xml:"totalFlow,omitempty" require:"true"`
  // {"en":"IP","zh_CN":"IP"}
  Ip *string `json:"ip,omitempty" xml:"ip,omitempty" require:"true"`
  // {"en":"attack time","zh_CN":"攻击时间"}
  Time *string `json:"time,omitempty" xml:"time,omitempty" require:"true"`
  // {"en":"attack type","zh_CN":"攻击类型"}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
}

func (s QueryDDoSAttackDetailsResponseData) String() string {
  return tea.Prettify(s)
}

func (s QueryDDoSAttackDetailsResponseData) GoString() string {
  return s.String()
}

func (s *QueryDDoSAttackDetailsResponseData) SetTotalFlow(v string) *QueryDDoSAttackDetailsResponseData {
  s.TotalFlow = &v
  return s
}

func (s *QueryDDoSAttackDetailsResponseData) SetIp(v string) *QueryDDoSAttackDetailsResponseData {
  s.Ip = &v
  return s
}

func (s *QueryDDoSAttackDetailsResponseData) SetTime(v string) *QueryDDoSAttackDetailsResponseData {
  s.Time = &v
  return s
}

func (s *QueryDDoSAttackDetailsResponseData) SetType(v string) *QueryDDoSAttackDetailsResponseData {
  s.Type = &v
  return s
}

type QueryDDoSAttackDetailsResponseHeader struct {
}

func (s QueryDDoSAttackDetailsResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryDDoSAttackDetailsResponseHeader) GoString() string {
  return s.String()
}




type GetWAFRequestAndAttackEventRequest struct {
  // {"en":"Domain, array.", "zh_CN":"域名，数组。"}
  Domains []*string `json:"domains,omitempty" xml:"domains,omitempty" require:"true" type:"Repeated"`
  // {"en":"Start time, yyyy-MM-dd HH:mm:ss.", "zh_CN":"开始时间，yyyy-MM-dd HH:mm:ss。"}
  StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty" require:"true"`
  // {"en":"End time, yyyy-MM-dd HH:mm:ss.", "zh_CN":"结束时间，yyyy-MM-dd HH:mm:ss。"}
  EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty" require:"true"`
  // {"en":"Time zone, GMT+8 by default.", "zh_CN":"时区，默认GTM+8，即“GTM+8”。"}
  TimeZone *string `json:"timeZone,omitempty" xml:"timeZone,omitempty"`
}

func (s GetWAFRequestAndAttackEventRequest) String() string {
  return tea.Prettify(s)
}

func (s GetWAFRequestAndAttackEventRequest) GoString() string {
  return s.String()
}

func (s *GetWAFRequestAndAttackEventRequest) SetDomains(v []*string) *GetWAFRequestAndAttackEventRequest {
  s.Domains = v
  return s
}

func (s *GetWAFRequestAndAttackEventRequest) SetStartTime(v string) *GetWAFRequestAndAttackEventRequest {
  s.StartTime = &v
  return s
}

func (s *GetWAFRequestAndAttackEventRequest) SetEndTime(v string) *GetWAFRequestAndAttackEventRequest {
  s.EndTime = &v
  return s
}

func (s *GetWAFRequestAndAttackEventRequest) SetTimeZone(v string) *GetWAFRequestAndAttackEventRequest {
  s.TimeZone = &v
  return s
}

type GetWAFRequestAndAttackEventResponse struct {
  // {"en":"Return 0 means success, please see <Error code> to check other status code.", "zh_CN":"0状态码表示请求成功，其他状态码说明请参见《错误码》。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Error message or Success.", "zh_CN":"错误信息或Success。"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Return data.", "zh_CN":"返回值。"}
  Data *TotalRequest `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s GetWAFRequestAndAttackEventResponse) String() string {
  return tea.Prettify(s)
}

func (s GetWAFRequestAndAttackEventResponse) GoString() string {
  return s.String()
}

func (s *GetWAFRequestAndAttackEventResponse) SetCode(v string) *GetWAFRequestAndAttackEventResponse {
  s.Code = &v
  return s
}

func (s *GetWAFRequestAndAttackEventResponse) SetMessage(v string) *GetWAFRequestAndAttackEventResponse {
  s.Message = &v
  return s
}

func (s *GetWAFRequestAndAttackEventResponse) SetData(v *TotalRequest) *GetWAFRequestAndAttackEventResponse {
  s.Data = v
  return s
}

type TotalRequest struct {
  // {"en":"Total requests.", "zh_CN":"检测请求数。"}
  TestRequest *int64 `json:"testRequest,omitempty" xml:"testRequest,omitempty" require:"true"`
  // {"en":"Attack requests.", "zh_CN":"攻击请求数。"}
  AttackRequest *int64 `json:"attackRequest,omitempty" xml:"attackRequest,omitempty" require:"true"`
  // {"en":"Block requests.", "zh_CN":"拦截请求数。"}
  BlockRequest *int64 `json:"blockRequest,omitempty" xml:"blockRequest,omitempty" require:"true"`
  // {"en":"High risk threat requests.", "zh_CN":"高风险攻击请求数。"}
  HighRiskAttack *int64 `json:"highRiskAttack,omitempty" xml:"highRiskAttack,omitempty" require:"true"`
  // {"en":"Persistent threat requests.", "zh_CN":"持续性渗透请求数。"}
  ContinuePermeate *int64 `json:"continuePermeate,omitempty" xml:"continuePermeate,omitempty" require:"true"`
  // {"en":"High frequency threat requests.", "zh_CN":"高频率攻击请求数。"}
  HighFrequencyAttack *int64 `json:"highFrequencyAttack,omitempty" xml:"highFrequencyAttack,omitempty" require:"true"`
}

func (s TotalRequest) String() string {
  return tea.Prettify(s)
}

func (s TotalRequest) GoString() string {
  return s.String()
}

func (s *TotalRequest) SetTestRequest(v int64) *TotalRequest {
  s.TestRequest = &v
  return s
}

func (s *TotalRequest) SetAttackRequest(v int64) *TotalRequest {
  s.AttackRequest = &v
  return s
}

func (s *TotalRequest) SetBlockRequest(v int64) *TotalRequest {
  s.BlockRequest = &v
  return s
}

func (s *TotalRequest) SetHighRiskAttack(v int64) *TotalRequest {
  s.HighRiskAttack = &v
  return s
}

func (s *TotalRequest) SetContinuePermeate(v int64) *TotalRequest {
  s.ContinuePermeate = &v
  return s
}

func (s *TotalRequest) SetHighFrequencyAttack(v int64) *TotalRequest {
  s.HighFrequencyAttack = &v
  return s
}

type GetWAFRequestAndAttackEventPaths struct {
}

func (s GetWAFRequestAndAttackEventPaths) String() string {
  return tea.Prettify(s)
}

func (s GetWAFRequestAndAttackEventPaths) GoString() string {
  return s.String()
}

type GetWAFRequestAndAttackEventParameters struct {
}

func (s GetWAFRequestAndAttackEventParameters) String() string {
  return tea.Prettify(s)
}

func (s GetWAFRequestAndAttackEventParameters) GoString() string {
  return s.String()
}

type GetWAFRequestAndAttackEventRequestHeader struct {
}

func (s GetWAFRequestAndAttackEventRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetWAFRequestAndAttackEventRequestHeader) GoString() string {
  return s.String()
}

type GetWAFRequestAndAttackEventResponseHeader struct {
}

func (s GetWAFRequestAndAttackEventResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetWAFRequestAndAttackEventResponseHeader) GoString() string {
  return s.String()
}




type GetOriginalRequestInformationRequest struct {
  // {"en":"Event ID.","zh_CN":"事件ID。"}
  Uuid *string `json:"uuid,omitempty" xml:"uuid,omitempty" require:"true"`
  // {"en":"Domain.","zh_CN":"域名。"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {"en":"Attack time, format: yyyy-MM-dd HH:mm:ss.","zh_CN":"攻击时间，格式：yyyy-MM-dd HH:mm:ss。"}
  AttackTime *string `json:"attackTime,omitempty" xml:"attackTime,omitempty" require:"true"`
}

func (s GetOriginalRequestInformationRequest) String() string {
  return tea.Prettify(s)
}

func (s GetOriginalRequestInformationRequest) GoString() string {
  return s.String()
}

func (s *GetOriginalRequestInformationRequest) SetUuid(v string) *GetOriginalRequestInformationRequest {
  s.Uuid = &v
  return s
}

func (s *GetOriginalRequestInformationRequest) SetDomain(v string) *GetOriginalRequestInformationRequest {
  s.Domain = &v
  return s
}

func (s *GetOriginalRequestInformationRequest) SetAttackTime(v string) *GetOriginalRequestInformationRequest {
  s.AttackTime = &v
  return s
}

type GetOriginalRequestInformationRequestHeader struct {
  // {"defaultValue":"en","en":"The language of response data, default value: en.\nzh_CN: Chinese\nen: English","zh_CN":"返回内容的语言版本，默认值: en。\nzh_CN：中文\nen：英文","exampleValue":"zh_CN,en"}
  Language *string `json:"language,omitempty" xml:"language,omitempty"`
  // {"dictionary":"belong=WAAP-MS-Ext|dict=waap_serviceType","en":"Security service type. Please enter a specific service type, if you purchase multiple security services.","zh_CN":"安全服务类型。有使用多个不同的安全服务时，需要填写具体的服务类型。"}
  ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty"`
  // {"defaultValue":"GMT+8","en":"Report Data Timezone:1. Indicates the timezone for the report data. lt must be relative to GMT andspecified in the format GMT+n where -12<=n <= 12. For example,timezone=GMT-7 and timezone=GMT+8.2. If the parameter is unspecified,results will be in the GMT+8 timezone.","zh_CN":"报表数据时区：1.请传递GMT时区。示例格式：GMT+N（其中12<=N <= 12），示例：timezone=GMT-7或timezone=GMT+8.2.若参数为传递，将默认按GMT+8时区查询"}
  Timezone *string `json:"timezone,omitempty" xml:"timezone,omitempty"`
}

func (s GetOriginalRequestInformationRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetOriginalRequestInformationRequestHeader) GoString() string {
  return s.String()
}

func (s *GetOriginalRequestInformationRequestHeader) SetLanguage(v string) *GetOriginalRequestInformationRequestHeader {
  s.Language = &v
  return s
}

func (s *GetOriginalRequestInformationRequestHeader) SetServiceType(v string) *GetOriginalRequestInformationRequestHeader {
  s.ServiceType = &v
  return s
}

func (s *GetOriginalRequestInformationRequestHeader) SetTimezone(v string) *GetOriginalRequestInformationRequestHeader {
  s.Timezone = &v
  return s
}

type GetOriginalRequestInformationPaths struct {
}

func (s GetOriginalRequestInformationPaths) String() string {
  return tea.Prettify(s)
}

func (s GetOriginalRequestInformationPaths) GoString() string {
  return s.String()
}

type GetOriginalRequestInformationParameters struct {
}

func (s GetOriginalRequestInformationParameters) String() string {
  return tea.Prettify(s)
}

func (s GetOriginalRequestInformationParameters) GoString() string {
  return s.String()
}

type GetOriginalRequestInformationResponse struct {
  // {"dictionary":"belong=WAAP-MS-Ext|dict=waap_retCodeEnum","en":"Please refer to the error code for exceptions.","zh_CN":"请参照错误码。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Description.","zh_CN":"描述信息。"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"Data.","zh_CN":"出参数据。"}
  Data *GetOriginalRequestInformationResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
}

func (s GetOriginalRequestInformationResponse) String() string {
  return tea.Prettify(s)
}

func (s GetOriginalRequestInformationResponse) GoString() string {
  return s.String()
}

func (s *GetOriginalRequestInformationResponse) SetCode(v string) *GetOriginalRequestInformationResponse {
  s.Code = &v
  return s
}

func (s *GetOriginalRequestInformationResponse) SetMsg(v string) *GetOriginalRequestInformationResponse {
  s.Msg = &v
  return s
}

func (s *GetOriginalRequestInformationResponse) SetData(v *GetOriginalRequestInformationResponseData) *GetOriginalRequestInformationResponse {
  s.Data = v
  return s
}

type GetOriginalRequestInformationResponseData struct {
  // {"en":"Request header information.","zh_CN":"请求头信息。"}
  RequestHead *string `json:"requestHead,omitempty" xml:"requestHead,omitempty" require:"true"`
}

func (s GetOriginalRequestInformationResponseData) String() string {
  return tea.Prettify(s)
}

func (s GetOriginalRequestInformationResponseData) GoString() string {
  return s.String()
}

func (s *GetOriginalRequestInformationResponseData) SetRequestHead(v string) *GetOriginalRequestInformationResponseData {
  s.RequestHead = &v
  return s
}

type GetOriginalRequestInformationResponseHeader struct {
}

func (s GetOriginalRequestInformationResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetOriginalRequestInformationResponseHeader) GoString() string {
  return s.String()
}




type GetDomainBotVisitDetailsRequest struct {
  // {"en":"Domain.Separate by';'.", "zh_CN":"域名串，分号拼接"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {"en":"StartTime. Format: yyyy-MM-dd HH:mm:ss", "zh_CN":"开始时间。格式： yyyy-MM-dd HH:mm:ss"}
  StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty" require:"true"`
  // {"en":"EndTime. Format: yyyy-MM-dd HH:mm:ss", "zh_CN":"结束时间。格式： yyyy-MM-dd HH:mm:ss"}
  EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty" require:"true"`
  // {"en":"Time zone, default 8, i.e.'GTM+8'", "zh_CN":"时区，默认8，即“GTM+8”"}
  TimeZone *int `json:"timeZone,omitempty" xml:"timeZone,omitempty"`
  // {"en":"Language type. Default cn. 
  //  en:English 
  //  cn:Chinese ", "zh_CN":"语言类型。 默认cn 
  //  en：英文 
  //  cn：中文"}
  Lang *string `json:"lang,omitempty" xml:"lang,omitempty"`
}

func (s GetDomainBotVisitDetailsRequest) String() string {
  return tea.Prettify(s)
}

func (s GetDomainBotVisitDetailsRequest) GoString() string {
  return s.String()
}

func (s *GetDomainBotVisitDetailsRequest) SetDomain(v string) *GetDomainBotVisitDetailsRequest {
  s.Domain = &v
  return s
}

func (s *GetDomainBotVisitDetailsRequest) SetStartTime(v string) *GetDomainBotVisitDetailsRequest {
  s.StartTime = &v
  return s
}

func (s *GetDomainBotVisitDetailsRequest) SetEndTime(v string) *GetDomainBotVisitDetailsRequest {
  s.EndTime = &v
  return s
}

func (s *GetDomainBotVisitDetailsRequest) SetTimeZone(v int) *GetDomainBotVisitDetailsRequest {
  s.TimeZone = &v
  return s
}

func (s *GetDomainBotVisitDetailsRequest) SetLang(v string) *GetDomainBotVisitDetailsRequest {
  s.Lang = &v
  return s
}

type GetDomainBotVisitDetailsResponse struct {
  // {"en":"Status code, success is '200'.", "zh_CN":"状态码，成功为“200”。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Return message, success is 'Success'.", "zh_CN":"返回信息，成功为“Success”。"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Data returned.", "zh_CN":"返回数据。"}
  Data []*GetDomainBotVisitDetailsResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s GetDomainBotVisitDetailsResponse) String() string {
  return tea.Prettify(s)
}

func (s GetDomainBotVisitDetailsResponse) GoString() string {
  return s.String()
}

func (s *GetDomainBotVisitDetailsResponse) SetCode(v string) *GetDomainBotVisitDetailsResponse {
  s.Code = &v
  return s
}

func (s *GetDomainBotVisitDetailsResponse) SetMessage(v string) *GetDomainBotVisitDetailsResponse {
  s.Message = &v
  return s
}

func (s *GetDomainBotVisitDetailsResponse) SetData(v []*GetDomainBotVisitDetailsResponseData) *GetDomainBotVisitDetailsResponse {
  s.Data = v
  return s
}

type GetDomainBotVisitDetailsResponseData struct     {
  // {"en":"Domain.", "zh_CN":"域名。"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {"en":"Total request count.", "zh_CN":"总请求数。"}
  TotalRequest *int64 `json:"totalRequest,omitempty" xml:"totalRequest,omitempty" require:"true"`
  // {"en":"Bot request count.", "zh_CN":"Bot请求数。"}
  BotRequest *int64 `json:"botRequest,omitempty" xml:"botRequest,omitempty" require:"true"`
  // {"en":"Relief attack count.", "zh_CN":"缓解攻击数。"}
  ReliefAttack *int64 `json:"reliefAttack,omitempty" xml:"reliefAttack,omitempty" require:"true"`
  // {"en":"Total attack type count.", "zh_CN":"攻击类型总数。"}
  TypeTotal *int64 `json:"typeTotal,omitempty" xml:"typeTotal,omitempty" require:"true"`
  // {"en":"Top5 bot types.", "zh_CN":"Bot类型TOP5。"}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
}

func (s GetDomainBotVisitDetailsResponseData) String() string {
  return tea.Prettify(s)
}

func (s GetDomainBotVisitDetailsResponseData) GoString() string {
  return s.String()
}

func (s *GetDomainBotVisitDetailsResponseData) SetDomain(v string) *GetDomainBotVisitDetailsResponseData {
  s.Domain = &v
  return s
}

func (s *GetDomainBotVisitDetailsResponseData) SetTotalRequest(v int64) *GetDomainBotVisitDetailsResponseData {
  s.TotalRequest = &v
  return s
}

func (s *GetDomainBotVisitDetailsResponseData) SetBotRequest(v int64) *GetDomainBotVisitDetailsResponseData {
  s.BotRequest = &v
  return s
}

func (s *GetDomainBotVisitDetailsResponseData) SetReliefAttack(v int64) *GetDomainBotVisitDetailsResponseData {
  s.ReliefAttack = &v
  return s
}

func (s *GetDomainBotVisitDetailsResponseData) SetTypeTotal(v int64) *GetDomainBotVisitDetailsResponseData {
  s.TypeTotal = &v
  return s
}

func (s *GetDomainBotVisitDetailsResponseData) SetType(v string) *GetDomainBotVisitDetailsResponseData {
  s.Type = &v
  return s
}

type GetDomainBotVisitDetailsPaths struct {
}

func (s GetDomainBotVisitDetailsPaths) String() string {
  return tea.Prettify(s)
}

func (s GetDomainBotVisitDetailsPaths) GoString() string {
  return s.String()
}

type GetDomainBotVisitDetailsParameters struct {
}

func (s GetDomainBotVisitDetailsParameters) String() string {
  return tea.Prettify(s)
}

func (s GetDomainBotVisitDetailsParameters) GoString() string {
  return s.String()
}

type GetDomainBotVisitDetailsRequestHeader struct {
}

func (s GetDomainBotVisitDetailsRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetDomainBotVisitDetailsRequestHeader) GoString() string {
  return s.String()
}

type GetDomainBotVisitDetailsResponseHeader struct {
}

func (s GetDomainBotVisitDetailsResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetDomainBotVisitDetailsResponseHeader) GoString() string {
  return s.String()
}




type L4DdosTrendRequest struct {
  // {"en":"Start time, format: yyyy-MM-dd HH:mm:ss.\nThe time range can not exceed 31 days.","zh_CN":"开始时间，格式：yyyy-MM-dd HH:mm:ss。\n时间范围不超过30天。"}
  StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty" require:"true"`
  // {"en":"End time, format: yyyy-MM-dd HH:mm:ss.\nThe time range can not exceed 31 days.","zh_CN":"结束时间，格式：yyyy-MM-dd HH:mm:ss。\n时间范围不超过30天。"}
  EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty" require:"true"`
}

func (s L4DdosTrendRequest) String() string {
  return tea.Prettify(s)
}

func (s L4DdosTrendRequest) GoString() string {
  return s.String()
}

func (s *L4DdosTrendRequest) SetStartTime(v string) *L4DdosTrendRequest {
  s.StartTime = &v
  return s
}

func (s *L4DdosTrendRequest) SetEndTime(v string) *L4DdosTrendRequest {
  s.EndTime = &v
  return s
}

type L4DdosTrendRequestHeader struct {
  // {"dictionary":"belong=WAAP-MS-Ext|dict=waap_serviceType","en":"Security service type. Please enter a specific service type, if you purchase multiple security services.","zh_CN":"安全服务类型。有使用多个不同的安全服务时，需要填写具体的服务类型。"}
  ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty"`
  // {"defaultValue":"GMT+8","en":"Report Data Timezone:1. Indicates the timezone for the report data. lt must be relative to GMT andspecified in the format GMT+n where -12<=n <= 12. For example,timezone=GMT-7 and timezone=GMT+8.2. If the parameter is unspecified,results will be in the GMT+8 timezone.","zh_CN":"报表数据时区：1.请传递GMT时区。示例格式：GMT+N（其中12<=N <= 12），示例：timezone=GMT-7或timezone=GMT+8.2.若参数为传递，将默认按GMT+8时区查询"}
  Timezone *string `json:"timezone,omitempty" xml:"timezone,omitempty"`
}

func (s L4DdosTrendRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s L4DdosTrendRequestHeader) GoString() string {
  return s.String()
}

func (s *L4DdosTrendRequestHeader) SetServiceType(v string) *L4DdosTrendRequestHeader {
  s.ServiceType = &v
  return s
}

func (s *L4DdosTrendRequestHeader) SetTimezone(v string) *L4DdosTrendRequestHeader {
  s.Timezone = &v
  return s
}

type L4DdosTrendPaths struct {
}

func (s L4DdosTrendPaths) String() string {
  return tea.Prettify(s)
}

func (s L4DdosTrendPaths) GoString() string {
  return s.String()
}

type L4DdosTrendParameters struct {
}

func (s L4DdosTrendParameters) String() string {
  return tea.Prettify(s)
}

func (s L4DdosTrendParameters) GoString() string {
  return s.String()
}

type L4DdosTrendResponse struct {
  // {"dictionary":"belong=WAAP-MS-Ext|dict=waap_retCodeEnum","en":"Please refer to the error code for exceptions.","zh_CN":"请参照错误码。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Description.","zh_CN":"描述信息。"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"Data.","zh_CN":"出参数据。"}
  Data []*L4DdosTrendResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
  // {"en":"Maximum attack bps.","zh_CN":"DDoS攻击带宽峰值"}
  PeakBps []*L4DdosTrendResponsePeakBps `json:"peak_bps,omitempty" xml:"peak_bps,omitempty" require:"true" type:"Repeated"`
  // {"en":"Maximum attack qps.","zh_CN":"DDoS攻击包速峰值"}
  PeakPps []*L4DdosTrendResponsePeakPps `json:"peak_pps,omitempty" xml:"peak_pps,omitempty" require:"true" type:"Repeated"`
}

func (s L4DdosTrendResponse) String() string {
  return tea.Prettify(s)
}

func (s L4DdosTrendResponse) GoString() string {
  return s.String()
}

func (s *L4DdosTrendResponse) SetCode(v string) *L4DdosTrendResponse {
  s.Code = &v
  return s
}

func (s *L4DdosTrendResponse) SetMsg(v string) *L4DdosTrendResponse {
  s.Msg = &v
  return s
}

func (s *L4DdosTrendResponse) SetData(v []*L4DdosTrendResponseData) *L4DdosTrendResponse {
  s.Data = v
  return s
}

func (s *L4DdosTrendResponse) SetPeakBps(v []*L4DdosTrendResponsePeakBps) *L4DdosTrendResponse {
  s.PeakBps = v
  return s
}

func (s *L4DdosTrendResponse) SetPeakPps(v []*L4DdosTrendResponsePeakPps) *L4DdosTrendResponse {
  s.PeakPps = v
  return s
}

type L4DdosTrendResponseData struct     {
  // {"en":"Normal bandwidth.(bps)","zh_CN":"正常带宽。（bps）"}
  NormalBps *string `json:"normal_bps,omitempty" xml:"normal_bps,omitempty" require:"true"`
  // {"en":"Total cleaning bandwidth.","zh_CN":"总清洗带宽。"}
  AttackBps *string `json:"attack_bps,omitempty" xml:"attack_bps,omitempty" require:"true"`
  // {"en":"Normal packages.(pps)","zh_CN":"正常包数。（pps）"}
  NormalPps *string `json:"normal_pps,omitempty" xml:"normal_pps,omitempty" require:"true"`
  // {"en":"Total  packages.","zh_CN":"总清洗包数。"}
  AttackPps *string `json:"attack_pps,omitempty" xml:"attack_pps,omitempty" require:"true"`
  // {"en":"SYN Flood.(bps)","zh_CN":"SYN Flood。（bps）"}
  SynBps *string `json:"syn_bps,omitempty" xml:"syn_bps,omitempty" require:"true"`
  // {"en":"SYN Flood.(pps)","zh_CN":"SYN Flood。（pps）"}
  SynPps *string `json:"syn_pps,omitempty" xml:"syn_pps,omitempty" require:"true"`
  // {"en":"ACK Flood.(bps)","zh_CN":"ACK Flood。（bps）"}
  AckBps *string `json:"ack_bps,omitempty" xml:"ack_bps,omitempty" require:"true"`
  // {"en":"ACK Flood.(pps)","zh_CN":"ACK Flood。（pps）"}
  AckPps *string `json:"ack_pps,omitempty" xml:"ack_pps,omitempty" require:"true"`
  // {"en":"UDP Flood. (bps)","zh_CN":"UDP Flood。 (bps)"}
  UdpBps *string `json:"udp_bps,omitempty" xml:"udp_bps,omitempty" require:"true"`
  // {"en":"UDP Flood. (pps)","zh_CN":"UDP Flood。 (pps)"}
  UdpPps *string `json:"udp_pps,omitempty" xml:"udp_pps,omitempty" require:"true"`
  // {"en":"ICMP Flood. (bps)","zh_CN":"ICMP Flood。 (bps)"}
  IcmpBps *string `json:"icmp_bps,omitempty" xml:"icmp_bps,omitempty" require:"true"`
  // {"en":"ICMP Flood. (pps)","zh_CN":"ICMP Flood。 (pps)"}
  IcmpPps *string `json:"icmp_pps,omitempty" xml:"icmp_pps,omitempty" require:"true"`
  // {"en":"Other Flood. (pps)","zh_CN":"Other Flood。 (pps)"}
  OtherPps *string `json:"other_pps,omitempty" xml:"other_pps,omitempty" require:"true"`
  // {"en":"Other Flood.(bps)","zh_CN":"Other Flood。(bps)"}
  OtherBps *string `json:"other_bps,omitempty" xml:"other_bps,omitempty" require:"true"`
  // {"en":"Time.","zh_CN":"时间。"}
  Time *string `json:"time,omitempty" xml:"time,omitempty" require:"true"`
}

func (s L4DdosTrendResponseData) String() string {
  return tea.Prettify(s)
}

func (s L4DdosTrendResponseData) GoString() string {
  return s.String()
}

func (s *L4DdosTrendResponseData) SetNormalBps(v string) *L4DdosTrendResponseData {
  s.NormalBps = &v
  return s
}

func (s *L4DdosTrendResponseData) SetAttackBps(v string) *L4DdosTrendResponseData {
  s.AttackBps = &v
  return s
}

func (s *L4DdosTrendResponseData) SetNormalPps(v string) *L4DdosTrendResponseData {
  s.NormalPps = &v
  return s
}

func (s *L4DdosTrendResponseData) SetAttackPps(v string) *L4DdosTrendResponseData {
  s.AttackPps = &v
  return s
}

func (s *L4DdosTrendResponseData) SetSynBps(v string) *L4DdosTrendResponseData {
  s.SynBps = &v
  return s
}

func (s *L4DdosTrendResponseData) SetSynPps(v string) *L4DdosTrendResponseData {
  s.SynPps = &v
  return s
}

func (s *L4DdosTrendResponseData) SetAckBps(v string) *L4DdosTrendResponseData {
  s.AckBps = &v
  return s
}

func (s *L4DdosTrendResponseData) SetAckPps(v string) *L4DdosTrendResponseData {
  s.AckPps = &v
  return s
}

func (s *L4DdosTrendResponseData) SetUdpBps(v string) *L4DdosTrendResponseData {
  s.UdpBps = &v
  return s
}

func (s *L4DdosTrendResponseData) SetUdpPps(v string) *L4DdosTrendResponseData {
  s.UdpPps = &v
  return s
}

func (s *L4DdosTrendResponseData) SetIcmpBps(v string) *L4DdosTrendResponseData {
  s.IcmpBps = &v
  return s
}

func (s *L4DdosTrendResponseData) SetIcmpPps(v string) *L4DdosTrendResponseData {
  s.IcmpPps = &v
  return s
}

func (s *L4DdosTrendResponseData) SetOtherPps(v string) *L4DdosTrendResponseData {
  s.OtherPps = &v
  return s
}

func (s *L4DdosTrendResponseData) SetOtherBps(v string) *L4DdosTrendResponseData {
  s.OtherBps = &v
  return s
}

func (s *L4DdosTrendResponseData) SetTime(v string) *L4DdosTrendResponseData {
  s.Time = &v
  return s
}

type L4DdosTrendResponsePeakBps struct     {
  // {"en":"Maximum attack time.","zh_CN":"峰值时间"}
  PeakTime *string `json:"peak_time,omitempty" xml:"peak_time,omitempty" require:"true"`
  // {"en":"Maximum attack value.","zh_CN":"峰值"}
  PeakValue *int64 `json:"peak_value,omitempty" xml:"peak_value,omitempty" require:"true"`
}

func (s L4DdosTrendResponsePeakBps) String() string {
  return tea.Prettify(s)
}

func (s L4DdosTrendResponsePeakBps) GoString() string {
  return s.String()
}

func (s *L4DdosTrendResponsePeakBps) SetPeakTime(v string) *L4DdosTrendResponsePeakBps {
  s.PeakTime = &v
  return s
}

func (s *L4DdosTrendResponsePeakBps) SetPeakValue(v int64) *L4DdosTrendResponsePeakBps {
  s.PeakValue = &v
  return s
}

type L4DdosTrendResponsePeakPps struct     {
  // {"en":"Maximum attack time.","zh_CN":"峰值时间"}
  PeakTime *string `json:"peak_time,omitempty" xml:"peak_time,omitempty" require:"true"`
  // {"en":"Maximum attack value.","zh_CN":"峰值"}
  PeakValue *int64 `json:"peak_value,omitempty" xml:"peak_value,omitempty" require:"true"`
}

func (s L4DdosTrendResponsePeakPps) String() string {
  return tea.Prettify(s)
}

func (s L4DdosTrendResponsePeakPps) GoString() string {
  return s.String()
}

func (s *L4DdosTrendResponsePeakPps) SetPeakTime(v string) *L4DdosTrendResponsePeakPps {
  s.PeakTime = &v
  return s
}

func (s *L4DdosTrendResponsePeakPps) SetPeakValue(v int64) *L4DdosTrendResponsePeakPps {
  s.PeakValue = &v
  return s
}

type L4DdosTrendResponseHeader struct {
}

func (s L4DdosTrendResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s L4DdosTrendResponseHeader) GoString() string {
  return s.String()
}




type GetWAFDomainHistoryRequest struct {
  // {"en":"Start time, yyyy-MM-dd HH:mm:ss.", "zh_CN":"开始时间，yyyy-MM-dd HH:mm:ss。"}
  StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty" require:"true"`
  // {"en":"End time, yyyy-MM-dd HH:mm:ss.", "zh_CN":"结束时间，yyyy-MM-dd HH:mm:ss。"}
  EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty" require:"true"`
  // {"en":"Time zone, GMT+8 by default.", "zh_CN":"时区，默认GMT+8，即“GMT+8”。"}
  TimeZone *string `json:"timeZone,omitempty" xml:"timeZone,omitempty"`
}

func (s GetWAFDomainHistoryRequest) String() string {
  return tea.Prettify(s)
}

func (s GetWAFDomainHistoryRequest) GoString() string {
  return s.String()
}

func (s *GetWAFDomainHistoryRequest) SetStartTime(v string) *GetWAFDomainHistoryRequest {
  s.StartTime = &v
  return s
}

func (s *GetWAFDomainHistoryRequest) SetEndTime(v string) *GetWAFDomainHistoryRequest {
  s.EndTime = &v
  return s
}

func (s *GetWAFDomainHistoryRequest) SetTimeZone(v string) *GetWAFDomainHistoryRequest {
  s.TimeZone = &v
  return s
}

type GetWAFDomainHistoryResponse struct {
  // {"en":"Return 0 means success, please see <Error code> to check other status code.", "zh_CN":"0状态码表示请求成功，其他状态码说明请参见《错误码》。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Error message or Success.", "zh_CN":"错误信息或Success。"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Return data.", "zh_CN":"返回值。"}
  Data []*string `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s GetWAFDomainHistoryResponse) String() string {
  return tea.Prettify(s)
}

func (s GetWAFDomainHistoryResponse) GoString() string {
  return s.String()
}

func (s *GetWAFDomainHistoryResponse) SetCode(v string) *GetWAFDomainHistoryResponse {
  s.Code = &v
  return s
}

func (s *GetWAFDomainHistoryResponse) SetMessage(v string) *GetWAFDomainHistoryResponse {
  s.Message = &v
  return s
}

func (s *GetWAFDomainHistoryResponse) SetData(v []*string) *GetWAFDomainHistoryResponse {
  s.Data = v
  return s
}

type GetWAFDomainHistoryPaths struct {
}

func (s GetWAFDomainHistoryPaths) String() string {
  return tea.Prettify(s)
}

func (s GetWAFDomainHistoryPaths) GoString() string {
  return s.String()
}

type GetWAFDomainHistoryParameters struct {
}

func (s GetWAFDomainHistoryParameters) String() string {
  return tea.Prettify(s)
}

func (s GetWAFDomainHistoryParameters) GoString() string {
  return s.String()
}

type GetWAFDomainHistoryRequestHeader struct {
}

func (s GetWAFDomainHistoryRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetWAFDomainHistoryRequestHeader) GoString() string {
  return s.String()
}

type GetWAFDomainHistoryResponseHeader struct {
}

func (s GetWAFDomainHistoryResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetWAFDomainHistoryResponseHeader) GoString() string {
  return s.String()
}




type GetBotAccessURLTopDataRequest struct {
  // {"en":"Domain.Separate by ';'.", "zh_CN":"域名。多个以;隔开"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {"en":"StartTime. Format: yyyy-MM-dd HH:mm:ss", "zh_CN":"开始时间。格式： yyyy-MM-dd HH:mm:ss"}
  StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty" require:"true"`
  // {"en":"EndTime. Format: yyyy-MM-dd HH:mm:ss", "zh_CN":"结束时间。格式： yyyy-MM-dd HH:mm:ss"}
  EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty" require:"true"`
  // {"en":"Time zone, default 8, i.e. 'GTM+8'", "zh_CN":"时区，默认8，即“GTM+8”"}
  TimeZone *int `json:"timeZone,omitempty" xml:"timeZone,omitempty"`
  // {"en":"Number of top values.Default value 10", "zh_CN":"排名最前值数目。默认10"}
  TopNum *int `json:"topNum,omitempty" xml:"topNum,omitempty"`
}

func (s GetBotAccessURLTopDataRequest) String() string {
  return tea.Prettify(s)
}

func (s GetBotAccessURLTopDataRequest) GoString() string {
  return s.String()
}

func (s *GetBotAccessURLTopDataRequest) SetDomain(v string) *GetBotAccessURLTopDataRequest {
  s.Domain = &v
  return s
}

func (s *GetBotAccessURLTopDataRequest) SetStartTime(v string) *GetBotAccessURLTopDataRequest {
  s.StartTime = &v
  return s
}

func (s *GetBotAccessURLTopDataRequest) SetEndTime(v string) *GetBotAccessURLTopDataRequest {
  s.EndTime = &v
  return s
}

func (s *GetBotAccessURLTopDataRequest) SetTimeZone(v int) *GetBotAccessURLTopDataRequest {
  s.TimeZone = &v
  return s
}

func (s *GetBotAccessURLTopDataRequest) SetTopNum(v int) *GetBotAccessURLTopDataRequest {
  s.TopNum = &v
  return s
}

type GetBotAccessURLTopDataResponse struct {
  // {"en":"Status code, success is '200'.", "zh_CN":"状态码，成功为“200”。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Return message, success is 'Success'.", "zh_CN":"返回信息，成功为“Success”。"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Data returned.", "zh_CN":"返回数据。"}
  Data []*GetBotAccessURLTopDataResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s GetBotAccessURLTopDataResponse) String() string {
  return tea.Prettify(s)
}

func (s GetBotAccessURLTopDataResponse) GoString() string {
  return s.String()
}

func (s *GetBotAccessURLTopDataResponse) SetCode(v string) *GetBotAccessURLTopDataResponse {
  s.Code = &v
  return s
}

func (s *GetBotAccessURLTopDataResponse) SetMessage(v string) *GetBotAccessURLTopDataResponse {
  s.Message = &v
  return s
}

func (s *GetBotAccessURLTopDataResponse) SetData(v []*GetBotAccessURLTopDataResponseData) *GetBotAccessURLTopDataResponse {
  s.Data = v
  return s
}

type GetBotAccessURLTopDataResponseData struct     {
  // {"en":"URL.", "zh_CN":"统计类型"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"Bot request count.", "zh_CN":"Bot请求数。"}
  Count *int64 `json:"count,omitempty" xml:"count,omitempty" require:"true"`
}

func (s GetBotAccessURLTopDataResponseData) String() string {
  return tea.Prettify(s)
}

func (s GetBotAccessURLTopDataResponseData) GoString() string {
  return s.String()
}

func (s *GetBotAccessURLTopDataResponseData) SetName(v string) *GetBotAccessURLTopDataResponseData {
  s.Name = &v
  return s
}

func (s *GetBotAccessURLTopDataResponseData) SetCount(v int64) *GetBotAccessURLTopDataResponseData {
  s.Count = &v
  return s
}

type GetBotAccessURLTopDataPaths struct {
}

func (s GetBotAccessURLTopDataPaths) String() string {
  return tea.Prettify(s)
}

func (s GetBotAccessURLTopDataPaths) GoString() string {
  return s.String()
}

type GetBotAccessURLTopDataParameters struct {
}

func (s GetBotAccessURLTopDataParameters) String() string {
  return tea.Prettify(s)
}

func (s GetBotAccessURLTopDataParameters) GoString() string {
  return s.String()
}

type GetBotAccessURLTopDataRequestHeader struct {
}

func (s GetBotAccessURLTopDataRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetBotAccessURLTopDataRequestHeader) GoString() string {
  return s.String()
}

type GetBotAccessURLTopDataResponseHeader struct {
}

func (s GetBotAccessURLTopDataResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetBotAccessURLTopDataResponseHeader) GoString() string {
  return s.String()
}




type GetBotRequestTypeDistributeDataRequest struct {
  // {"en":"Domain.Separate by ';'.", "zh_CN":"域名。多个以;隔开"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {"en":"StartTime. Format: yyyy-MM-dd HH:mm:ss", "zh_CN":"开始时间。格式： yyyy-MM-dd HH:mm:ss"}
  StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty" require:"true"`
  // {"en":"EndTime. Format: yyyy-MM-dd HH:mm:ss", "zh_CN":"结束时间。格式： yyyy-MM-dd HH:mm:ss"}
  EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty" require:"true"`
  // {"en":"Time zone, default 8, i.e. 'GTM+8'", "zh_CN":"时区，默认8，即“GTM+8”"}
  TimeZone *int `json:"timeZone,omitempty" xml:"timeZone,omitempty"`
  // {"en":"Language type. Default cn. 
  //  en:English 
  //  cn:Chinese ", "zh_CN":"语言类型。 默认cn 
  //  en：英文 
  //  cn：中文"}
  Lang *string `json:"lang,omitempty" xml:"lang,omitempty"`
  // {"en":"Number of top values.Default value 10", "zh_CN":"排名最前值数目。默认10"}
  TopNum *int `json:"topNum,omitempty" xml:"topNum,omitempty"`
  // {"en":"Bot type.
  //  0:All 
  //  1:known type bot 
  //  2:unknown type bot", "zh_CN":"Bot类型。 
  //  0：全部 
  //  1：已知bot 
  //  2：未知bot"}
  Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetBotRequestTypeDistributeDataRequest) String() string {
  return tea.Prettify(s)
}

func (s GetBotRequestTypeDistributeDataRequest) GoString() string {
  return s.String()
}

func (s *GetBotRequestTypeDistributeDataRequest) SetDomain(v string) *GetBotRequestTypeDistributeDataRequest {
  s.Domain = &v
  return s
}

func (s *GetBotRequestTypeDistributeDataRequest) SetStartTime(v string) *GetBotRequestTypeDistributeDataRequest {
  s.StartTime = &v
  return s
}

func (s *GetBotRequestTypeDistributeDataRequest) SetEndTime(v string) *GetBotRequestTypeDistributeDataRequest {
  s.EndTime = &v
  return s
}

func (s *GetBotRequestTypeDistributeDataRequest) SetTimeZone(v int) *GetBotRequestTypeDistributeDataRequest {
  s.TimeZone = &v
  return s
}

func (s *GetBotRequestTypeDistributeDataRequest) SetLang(v string) *GetBotRequestTypeDistributeDataRequest {
  s.Lang = &v
  return s
}

func (s *GetBotRequestTypeDistributeDataRequest) SetTopNum(v int) *GetBotRequestTypeDistributeDataRequest {
  s.TopNum = &v
  return s
}

func (s *GetBotRequestTypeDistributeDataRequest) SetType(v string) *GetBotRequestTypeDistributeDataRequest {
  s.Type = &v
  return s
}

type GetBotRequestTypeDistributeDataResponse struct {
  // {"en":"Status code, success is '200'.", "zh_CN":"状态码，成功为“200”。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Return message, success is 'Success'.", "zh_CN":"返回信息，成功为“Success”。"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Data returned.", "zh_CN":"返回数据。"}
  Data []*GetBotRequestTypeDistributeDataResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s GetBotRequestTypeDistributeDataResponse) String() string {
  return tea.Prettify(s)
}

func (s GetBotRequestTypeDistributeDataResponse) GoString() string {
  return s.String()
}

func (s *GetBotRequestTypeDistributeDataResponse) SetCode(v string) *GetBotRequestTypeDistributeDataResponse {
  s.Code = &v
  return s
}

func (s *GetBotRequestTypeDistributeDataResponse) SetMessage(v string) *GetBotRequestTypeDistributeDataResponse {
  s.Message = &v
  return s
}

func (s *GetBotRequestTypeDistributeDataResponse) SetData(v []*GetBotRequestTypeDistributeDataResponseData) *GetBotRequestTypeDistributeDataResponse {
  s.Data = v
  return s
}

type GetBotRequestTypeDistributeDataResponseData struct     {
  // {"en":"Policies Triggered.", "zh_CN":"触发规则。"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"Bot Type. 
  //   1:known type bot 
  //  2:unknown type bot", "zh_CN":"Bot类型。 
  //  1：已知bot 
  //  2：未知bot"}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
  // {"en":"Bot request count.", "zh_CN":"Bot请求数。"}
  Value *int64 `json:"value,omitempty" xml:"value,omitempty" require:"true"`
}

func (s GetBotRequestTypeDistributeDataResponseData) String() string {
  return tea.Prettify(s)
}

func (s GetBotRequestTypeDistributeDataResponseData) GoString() string {
  return s.String()
}

func (s *GetBotRequestTypeDistributeDataResponseData) SetName(v string) *GetBotRequestTypeDistributeDataResponseData {
  s.Name = &v
  return s
}

func (s *GetBotRequestTypeDistributeDataResponseData) SetType(v string) *GetBotRequestTypeDistributeDataResponseData {
  s.Type = &v
  return s
}

func (s *GetBotRequestTypeDistributeDataResponseData) SetValue(v int64) *GetBotRequestTypeDistributeDataResponseData {
  s.Value = &v
  return s
}

type GetBotRequestTypeDistributeDataPaths struct {
}

func (s GetBotRequestTypeDistributeDataPaths) String() string {
  return tea.Prettify(s)
}

func (s GetBotRequestTypeDistributeDataPaths) GoString() string {
  return s.String()
}

type GetBotRequestTypeDistributeDataParameters struct {
}

func (s GetBotRequestTypeDistributeDataParameters) String() string {
  return tea.Prettify(s)
}

func (s GetBotRequestTypeDistributeDataParameters) GoString() string {
  return s.String()
}

type GetBotRequestTypeDistributeDataRequestHeader struct {
}

func (s GetBotRequestTypeDistributeDataRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetBotRequestTypeDistributeDataRequestHeader) GoString() string {
  return s.String()
}

type GetBotRequestTypeDistributeDataResponseHeader struct {
}

func (s GetBotRequestTypeDistributeDataResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetBotRequestTypeDistributeDataResponseHeader) GoString() string {
  return s.String()
}




type GetTopAttackTargetsByDomainRequest struct {
  // {"defaultValue":"100","en":"Top rankings, default value: 100, max: 1000.","zh_CN":"取前几排名，默认：100，上限：1000。"}
  Top *int `json:"top,omitempty" xml:"top,omitempty"`
  // {"en":"Multiple selection. Handling results, default: display all results.\nmitigated: Number of mitigated requests.\nmonitored: Number of observed requests.","zh_CN":"多选。处置结果，默认：展示所有结果。\nmitigated：已抵御请求数。\nmonitored：观察请求数。","exampleValue":"mitigated,monitored"}
  ActType []*string `json:"actType,omitempty" xml:"actType,omitempty" type:"Repeated"`
  // {"en":"Start time, format: yyyy-MM-dd HH:mm:ss.\nThe time range can not exceed 31 days.","zh_CN":"起始时间，格式：yyyy-MM-dd HH:mm:ss。\n时间范围不超过31天。"}
  StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty" require:"true"`
  // {"en":"End time, format: yyyy-MM-dd HH:mm:ss.\nThe time range can not exceed 31 days.","zh_CN":"截止时间，格式：yyyy-MM-dd HH:mm:ss。\n时间范围不超过31天。"}
  EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty" require:"true"`
  // {"en":"Domain list. Queries all domains under the account when not specified.","zh_CN":"域名列表，未指定时查询账号下的所有域名。"}
  Domains []*string `json:"domains,omitempty" xml:"domains,omitempty" type:"Repeated"`
}

func (s GetTopAttackTargetsByDomainRequest) String() string {
  return tea.Prettify(s)
}

func (s GetTopAttackTargetsByDomainRequest) GoString() string {
  return s.String()
}

func (s *GetTopAttackTargetsByDomainRequest) SetTop(v int) *GetTopAttackTargetsByDomainRequest {
  s.Top = &v
  return s
}

func (s *GetTopAttackTargetsByDomainRequest) SetActType(v []*string) *GetTopAttackTargetsByDomainRequest {
  s.ActType = v
  return s
}

func (s *GetTopAttackTargetsByDomainRequest) SetStartTime(v string) *GetTopAttackTargetsByDomainRequest {
  s.StartTime = &v
  return s
}

func (s *GetTopAttackTargetsByDomainRequest) SetEndTime(v string) *GetTopAttackTargetsByDomainRequest {
  s.EndTime = &v
  return s
}

func (s *GetTopAttackTargetsByDomainRequest) SetDomains(v []*string) *GetTopAttackTargetsByDomainRequest {
  s.Domains = v
  return s
}

type GetTopAttackTargetsByDomainRequestHeader struct {
  // {"defaultValue":"en","en":"The language of response data, default value: en.\nzh_CN: Chinese\nen: English","zh_CN":"返回内容的语言版本，默认值: en。\nzh_CN：中文\nen：英文","exampleValue":"zh_CN,en"}
  Language *string `json:"language,omitempty" xml:"language,omitempty"`
  // {"dictionary":"belong=WAAP-MS-Ext|dict=waap_serviceType","en":"Security service type. Please enter a specific service type, if you purchase multiple security services.","zh_CN":"安全服务类型。有使用多个不同的安全服务时，需要填写具体的服务类型。"}
  ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty"`
  // {"defaultValue":"GMT+8","en":"Report Data Timezone:1. Indicates the timezone for the report data. lt must be relative to GMT andspecified in the format GMT+n where -12<=n <= 12. For example,timezone=GMT-7 and timezone=GMT+8.2. If the parameter is unspecified,results will be in the GMT+8 timezone.","zh_CN":"报表数据时区：1.请传递GMT时区。示例格式：GMT+N（其中12<=N <= 12），示例：timezone=GMT-7或timezone=GMT+8.2.若参数为传递，将默认按GMT+8时区查询"}
  Timezone *string `json:"timezone,omitempty" xml:"timezone,omitempty"`
}

func (s GetTopAttackTargetsByDomainRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetTopAttackTargetsByDomainRequestHeader) GoString() string {
  return s.String()
}

func (s *GetTopAttackTargetsByDomainRequestHeader) SetLanguage(v string) *GetTopAttackTargetsByDomainRequestHeader {
  s.Language = &v
  return s
}

func (s *GetTopAttackTargetsByDomainRequestHeader) SetServiceType(v string) *GetTopAttackTargetsByDomainRequestHeader {
  s.ServiceType = &v
  return s
}

func (s *GetTopAttackTargetsByDomainRequestHeader) SetTimezone(v string) *GetTopAttackTargetsByDomainRequestHeader {
  s.Timezone = &v
  return s
}

type GetTopAttackTargetsByDomainPaths struct {
}

func (s GetTopAttackTargetsByDomainPaths) String() string {
  return tea.Prettify(s)
}

func (s GetTopAttackTargetsByDomainPaths) GoString() string {
  return s.String()
}

type GetTopAttackTargetsByDomainParameters struct {
}

func (s GetTopAttackTargetsByDomainParameters) String() string {
  return tea.Prettify(s)
}

func (s GetTopAttackTargetsByDomainParameters) GoString() string {
  return s.String()
}

type GetTopAttackTargetsByDomainResponse struct {
  // {"dictionary":"belong=WAAP-MS-Ext|dict=waap_retCodeEnum","en":"Please refer to the error code for exceptions.","zh_CN":"请参照错误码。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Description.","zh_CN":"描述信息。"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"Data.","zh_CN":"出参数据。"}
  Data []*GetTopAttackTargetsByDomainResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s GetTopAttackTargetsByDomainResponse) String() string {
  return tea.Prettify(s)
}

func (s GetTopAttackTargetsByDomainResponse) GoString() string {
  return s.String()
}

func (s *GetTopAttackTargetsByDomainResponse) SetCode(v string) *GetTopAttackTargetsByDomainResponse {
  s.Code = &v
  return s
}

func (s *GetTopAttackTargetsByDomainResponse) SetMsg(v string) *GetTopAttackTargetsByDomainResponse {
  s.Msg = &v
  return s
}

func (s *GetTopAttackTargetsByDomainResponse) SetData(v []*GetTopAttackTargetsByDomainResponseData) *GetTopAttackTargetsByDomainResponse {
  s.Data = v
  return s
}

type GetTopAttackTargetsByDomainResponseData struct     {
  // {"en":"Domain.","zh_CN":"域名。"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {"en":"Attack requests.","zh_CN":"攻击请求数。"}
  Attack *int64 `json:"attack,omitempty" xml:"attack,omitempty" require:"true"`
}

func (s GetTopAttackTargetsByDomainResponseData) String() string {
  return tea.Prettify(s)
}

func (s GetTopAttackTargetsByDomainResponseData) GoString() string {
  return s.String()
}

func (s *GetTopAttackTargetsByDomainResponseData) SetDomain(v string) *GetTopAttackTargetsByDomainResponseData {
  s.Domain = &v
  return s
}

func (s *GetTopAttackTargetsByDomainResponseData) SetAttack(v int64) *GetTopAttackTargetsByDomainResponseData {
  s.Attack = &v
  return s
}

type GetTopAttackTargetsByDomainResponseHeader struct {
}

func (s GetTopAttackTargetsByDomainResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetTopAttackTargetsByDomainResponseHeader) GoString() string {
  return s.String()
}




type GetTopPoliciesTriggeredRequest struct {
  // {"en":"Start time, format: yyyy-MM-dd HH:mm:ss.\nThe time range can not exceed 31 days.","zh_CN":"起始时间，格式：yyyy-MM-dd HH:mm:ss。\n时间范围不超过31天。"}
  StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty" require:"true"`
  // {"en":"End time, format: yyyy-MM-dd HH:mm:ss.\nThe time range can not exceed 31 days.","zh_CN":"截止时间，格式：yyyy-MM-dd HH:mm:ss。\n时间范围不超过31天。"}
  EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty" require:"true"`
  // {"en":"Domain list. Queries all domains under the account when not specified.","zh_CN":"域名列表，未指定时查询账号下的所有域名。"}
  Domains []*string `json:"domains,omitempty" xml:"domains,omitempty" type:"Repeated"`
  // {"en":"Multiple selection. Handling results, default: display all results.\nmitigated: Number of mitigated requests.\nmonitored: Number of observed requests.","zh_CN":"多选。处置结果，默认：展示所有结果。\nmitigated：已抵御请求数。\nmonitored：观察请求数。","exampleValue":"mitigated,monitored"}
  ActType []*string `json:"actType,omitempty" xml:"actType,omitempty" type:"Repeated"`
}

func (s GetTopPoliciesTriggeredRequest) String() string {
  return tea.Prettify(s)
}

func (s GetTopPoliciesTriggeredRequest) GoString() string {
  return s.String()
}

func (s *GetTopPoliciesTriggeredRequest) SetStartTime(v string) *GetTopPoliciesTriggeredRequest {
  s.StartTime = &v
  return s
}

func (s *GetTopPoliciesTriggeredRequest) SetEndTime(v string) *GetTopPoliciesTriggeredRequest {
  s.EndTime = &v
  return s
}

func (s *GetTopPoliciesTriggeredRequest) SetDomains(v []*string) *GetTopPoliciesTriggeredRequest {
  s.Domains = v
  return s
}

func (s *GetTopPoliciesTriggeredRequest) SetActType(v []*string) *GetTopPoliciesTriggeredRequest {
  s.ActType = v
  return s
}

type GetTopPoliciesTriggeredRequestHeader struct {
  // {"dictionary":"belong=WAAP-MS-Ext|dict=waap_serviceType","en":"Security service type. Please enter a specific service type, if you purchase multiple security services.","zh_CN":"安全服务类型。有使用多个不同的安全服务时，需要填写具体的服务类型。"}
  ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty"`
  // {"defaultValue":"GMT+8","en":"Report Data Timezone:1. Indicates the timezone for the report data. lt must be relative to GMT andspecified in the format GMT+n where -12<=n <= 12. For example,timezone=GMT-7 and timezone=GMT+8.2. If the parameter is unspecified,results will be in the GMT+8 timezone.","zh_CN":"报表数据时区：1.请传递GMT时区。示例格式：GMT+N（其中12<=N <= 12），示例：timezone=GMT-7或timezone=GMT+8.2.若参数为传递，将默认按GMT+8时区查询"}
  Timezone *string `json:"timezone,omitempty" xml:"timezone,omitempty"`
}

func (s GetTopPoliciesTriggeredRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetTopPoliciesTriggeredRequestHeader) GoString() string {
  return s.String()
}

func (s *GetTopPoliciesTriggeredRequestHeader) SetServiceType(v string) *GetTopPoliciesTriggeredRequestHeader {
  s.ServiceType = &v
  return s
}

func (s *GetTopPoliciesTriggeredRequestHeader) SetTimezone(v string) *GetTopPoliciesTriggeredRequestHeader {
  s.Timezone = &v
  return s
}

type GetTopPoliciesTriggeredPaths struct {
}

func (s GetTopPoliciesTriggeredPaths) String() string {
  return tea.Prettify(s)
}

func (s GetTopPoliciesTriggeredPaths) GoString() string {
  return s.String()
}

type GetTopPoliciesTriggeredParameters struct {
}

func (s GetTopPoliciesTriggeredParameters) String() string {
  return tea.Prettify(s)
}

func (s GetTopPoliciesTriggeredParameters) GoString() string {
  return s.String()
}

type GetTopPoliciesTriggeredResponse struct {
  // {"dictionary":"belong=WAAP-MS-Ext|dict=waap_retCodeEnum","en":"Please refer to the error code for exceptions.","zh_CN":"请参照错误码。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Description.","zh_CN":"描述信息。"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"Data.","zh_CN":"出参数据。"}
  Data []*GetTopPoliciesTriggeredResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s GetTopPoliciesTriggeredResponse) String() string {
  return tea.Prettify(s)
}

func (s GetTopPoliciesTriggeredResponse) GoString() string {
  return s.String()
}

func (s *GetTopPoliciesTriggeredResponse) SetCode(v string) *GetTopPoliciesTriggeredResponse {
  s.Code = &v
  return s
}

func (s *GetTopPoliciesTriggeredResponse) SetMsg(v string) *GetTopPoliciesTriggeredResponse {
  s.Msg = &v
  return s
}

func (s *GetTopPoliciesTriggeredResponse) SetData(v []*GetTopPoliciesTriggeredResponseData) *GetTopPoliciesTriggeredResponse {
  s.Data = v
  return s
}

type GetTopPoliciesTriggeredResponseData struct     {
  // {"en":"Attack type.\nBLOCK: IP/Geo Block\nDMS_DEFEND: DDoS Protection\nWAF_DEFEND: WAF\nBOT_MANAGE: Bot Management\nAPI_DEFEND: API Security\nINTELLIGENCE: Threat Intelligence\nRATE_LIMIT: Rate Limiting\nCUSTOMIZE_RULE: Custom Rules","zh_CN":"攻击类型。\nBLOCK：IP/区域封禁\nDMS_DEFEND：DDoS防护\nWAF_DEFEND：WAF\nBOT_MANAGE：Bot管理\nAPI_DEFEND：API安全\nINTELLIGENCE：威胁情报\nRATE_LIMIT：频率限制\nCUSTOMIZE_RULE：自定义规则","exampleValue":"BLOCK,DMS_DEFEND,WAF_DEFEND,BOT_MANAGE,API_DEFEND,INTELLIGENCE,RATE_LIMIT,CUSTOMIZE_RULE"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Number of requests that triggered the policy type.","zh_CN":"触发该策略类型的请求数。"}
  Value *int64 `json:"value,omitempty" xml:"value,omitempty" require:"true"`
}

func (s GetTopPoliciesTriggeredResponseData) String() string {
  return tea.Prettify(s)
}

func (s GetTopPoliciesTriggeredResponseData) GoString() string {
  return s.String()
}

func (s *GetTopPoliciesTriggeredResponseData) SetCode(v string) *GetTopPoliciesTriggeredResponseData {
  s.Code = &v
  return s
}

func (s *GetTopPoliciesTriggeredResponseData) SetValue(v int64) *GetTopPoliciesTriggeredResponseData {
  s.Value = &v
  return s
}

type GetTopPoliciesTriggeredResponseHeader struct {
}

func (s GetTopPoliciesTriggeredResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetTopPoliciesTriggeredResponseHeader) GoString() string {
  return s.String()
}




type GetDistributionOfWAFAttackTypeRequest struct {
  // {"en":"Domain, array.", "zh_CN":"域名，数组。"}
  Domains []*string `json:"domains,omitempty" xml:"domains,omitempty" require:"true" type:"Repeated"`
  // {"en":"Start time, yyyy-MM-dd HH:mm:ss.", "zh_CN":"开始时间，yyyy-MM-dd HH:mm:ss。"}
  StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty" require:"true"`
  // {"en":"End time, yyyy-MM-dd HH:mm:ss.", "zh_CN":"结束时间，yyyy-MM-dd HH:mm:ss。"}
  EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty" require:"true"`
  // {"en":"Time zone, GMT+8 by default.", "zh_CN":"时区，默认GMT+8，即“GMT+8”。"}
  TimeZone *string `json:"timeZone,omitempty" xml:"timeZone,omitempty"`
  // {"en":"Attack type, array. [SQL Injection: WAF_SQLI,
  // X-Site Scripting: WAF_XSS,
  // File Inclusion: WAF_RFI,
  // File Uploading: WAF_FILE_UPLOAD,
  // Command Injection: WAF_CMDI,
  // Directory Traversal: WAF_DIR_TRAVERSAL,
  // 3rd Party Component Exploit: WAF_THIRDCOMP,
  // XPATH/LDAP/SSI Injection: WAF_XPATH_LDAP_SSI,
  // Malicious Scanning: WAF_SCANNER,
  // Webshell Uploading: WAF_SHELL_BACKDOOR,
  // Unauthorized Download: WAF_FILE_DOWNLOAD,
  // Illegal Access(RFC): WAF_ILLEGAL,
  // Illegal HTTP Method(RFC): WAF_INVALID_METHOD,
  // Illegal HTTP Version(RFC): WAF_INVALID_PROTOCOL,
  // HTTP Buffer Overflow(RFC): WAF_BUF_OVERFLOW,
  // Server Sensitive Info Leakage[remarks:combination]: WAF_SERVER_ERROR_LEAKAGE,WAF_SENS_DATA_LEAKAGE,
  // XML Injection: WAF_XXE,
  // Custom Rules: WAF_CUSTOM_RULE,
  // Cookie Protection: WAF_COOKIE_PROTECT,
  // CSRF: WAF_CSRF,
  // Adware Insertion Prevention: WAF_UAD,
  // Webshell Access Detection: WAF_WEBSHELL,
  // Threat Intelligence: WAF_THREAT_INTELLIGENCE,
  // Credential Stuffing: WAF_HIT_LIB,
  // Server-side Request Forge: WAF_SSRF,
  // Malicious Dig: WAF_COIN,
  // Access Control[remarks:combination]: WAF_BLACK_IP,WAF_BLACK_URL,WAF_BLACK_UA,WAF_BLACK_HEADER,WAF_ACCESS_CONTROL,
  // Rate Limiting: WAF_FORCE_CRACKING,
  // Response Code Rate Limiting: WAF_BLACK_STATUS,
  // IP Repeated Violations: WAF_DYNAMIC_BLACK_IP]", "zh_CN":"攻击类型，数组。[SQL注入：WAF_SQLI,
  // XSS跨站：WAF_XSS,
  // 文件包含：WAF_RFI,
  // 文件上传：WAF_FILE_UPLOAD,
  // 命令注入：WAF_CMDI,
  // 目录遍历：WAF_DIR_TRAVERSAL,
  // 第三方组件漏洞：WAF_THIRDCOMP,
  // XPATH/LDAP/SSI注入：WAF_XPATH_LDAP_SSI,
  // 扫描器：WAF_SCANNER,
  // 木马后门：WAF_SHELL_BACKDOOR,
  // 非法下载：WAF_FILE_DOWNLOAD,
  // 非法请求(合规)：WAF_ILLEGAL,
  // 非法请求方法(合规)：WAF_INVALID_METHOD,
  // 非法请求协议(合规)：WAF_INVALID_PROTOCOL,
  // 缓冲区溢出(合规)：WAF_BUF_OVERFLOW,
  // 服务器信息泄露[备:多个组合]：WAF_SERVER_ERROR_LEAKAGE,WAF_SENS_DATA_LEAKAGE,
  // XML注入：WAF_XXE,
  // 自定义规则：WAF_CUSTOM_RULE,
  // Cookie防护：WAF_COOKIE_PROTECT,
  // CSRF：WAF_CSRF,
  // 恶意广告检测：WAF_UAD,
  // 后门识别：WAF_WEBSHELL,
  // 威胁情报：WAF_THREAT_INTELLIGENCE,
  // 撞库：WAF_HIT_LIB,
  // 服务端请求伪造：WAF_SSRF,
  // 恶意挖矿：WAF_COIN,
  // 访问控制[备:多个组合]：WAF_BLACK_IP,WAF_BLACK_URL,WAF_BLACK_UA,WAF_BLACK_HEADER,WAF_ACCESS_CONTROL,
  // 频率限制：WAF_FORCE_CRACKING,
  // 状态码限速：WAF_BLACK_STATUS,
  // 攻击IP惩罚：WAF_DYNAMIC_BLACK_IP]"}
  AttackTypes []*string `json:"attackTypes,omitempty" xml:"attackTypes,omitempty" type:"Repeated"`
  // {"en":"Action, defaule 1 and 2.
  // 1: Block
  // 2: Log", "zh_CN":"处理动作，默认1和2。
  // 1：拦截
  // 2：监控"}
  Acts []*string `json:"acts,omitempty" xml:"acts,omitempty" type:"Repeated"`
  // {"en":"Attack type top value, default 100.", "zh_CN":"攻击类型top值，默认100。"}
  TopNum *int `json:"topNum,omitempty" xml:"topNum,omitempty"`
}

func (s GetDistributionOfWAFAttackTypeRequest) String() string {
  return tea.Prettify(s)
}

func (s GetDistributionOfWAFAttackTypeRequest) GoString() string {
  return s.String()
}

func (s *GetDistributionOfWAFAttackTypeRequest) SetDomains(v []*string) *GetDistributionOfWAFAttackTypeRequest {
  s.Domains = v
  return s
}

func (s *GetDistributionOfWAFAttackTypeRequest) SetStartTime(v string) *GetDistributionOfWAFAttackTypeRequest {
  s.StartTime = &v
  return s
}

func (s *GetDistributionOfWAFAttackTypeRequest) SetEndTime(v string) *GetDistributionOfWAFAttackTypeRequest {
  s.EndTime = &v
  return s
}

func (s *GetDistributionOfWAFAttackTypeRequest) SetTimeZone(v string) *GetDistributionOfWAFAttackTypeRequest {
  s.TimeZone = &v
  return s
}

func (s *GetDistributionOfWAFAttackTypeRequest) SetAttackTypes(v []*string) *GetDistributionOfWAFAttackTypeRequest {
  s.AttackTypes = v
  return s
}

func (s *GetDistributionOfWAFAttackTypeRequest) SetActs(v []*string) *GetDistributionOfWAFAttackTypeRequest {
  s.Acts = v
  return s
}

func (s *GetDistributionOfWAFAttackTypeRequest) SetTopNum(v int) *GetDistributionOfWAFAttackTypeRequest {
  s.TopNum = &v
  return s
}

type GetDistributionOfWAFAttackTypeResponse struct {
  // {"en":"Return 0 means success, please see <Error code> to check other status code.", "zh_CN":"0状态码表示请求成功，其他状态码说明请参见《错误码》。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Error message or Success.", "zh_CN":"错误信息或Success。"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Return data.", "zh_CN":"返回值。"}
  Data *GetDistributionOfWAFAttackTypeAttackTypePage `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s GetDistributionOfWAFAttackTypeResponse) String() string {
  return tea.Prettify(s)
}

func (s GetDistributionOfWAFAttackTypeResponse) GoString() string {
  return s.String()
}

func (s *GetDistributionOfWAFAttackTypeResponse) SetCode(v string) *GetDistributionOfWAFAttackTypeResponse {
  s.Code = &v
  return s
}

func (s *GetDistributionOfWAFAttackTypeResponse) SetMessage(v string) *GetDistributionOfWAFAttackTypeResponse {
  s.Message = &v
  return s
}

func (s *GetDistributionOfWAFAttackTypeResponse) SetData(v *GetDistributionOfWAFAttackTypeAttackTypePage) *GetDistributionOfWAFAttackTypeResponse {
  s.Data = v
  return s
}

type GetDistributionOfWAFAttackTypeAttackTypePage struct {
  // {"en":"Attack requests.", "zh_CN":"攻击请求数。"}
  AllTotalCount *int64 `json:"allTotalCount,omitempty" xml:"allTotalCount,omitempty" require:"true"`
  // {"en":"Attack type.", "zh_CN":"攻击类型。"}
  Top []*GetDistributionOfWAFAttackTypeAttackTypeTop `json:"top,omitempty" xml:"top,omitempty" require:"true" type:"Repeated"`
}

func (s GetDistributionOfWAFAttackTypeAttackTypePage) String() string {
  return tea.Prettify(s)
}

func (s GetDistributionOfWAFAttackTypeAttackTypePage) GoString() string {
  return s.String()
}

func (s *GetDistributionOfWAFAttackTypeAttackTypePage) SetAllTotalCount(v int64) *GetDistributionOfWAFAttackTypeAttackTypePage {
  s.AllTotalCount = &v
  return s
}

func (s *GetDistributionOfWAFAttackTypeAttackTypePage) SetTop(v []*GetDistributionOfWAFAttackTypeAttackTypeTop) *GetDistributionOfWAFAttackTypeAttackTypePage {
  s.Top = v
  return s
}

type GetDistributionOfWAFAttackTypeAttackTypeTop struct {
  // {"en":"English name of attack type.", "zh_CN":"攻击类型英文。"}
  AttackTypeNameEn *string `json:"attackTypeNameEn,omitempty" xml:"attackTypeNameEn,omitempty" require:"true"`
  // {"en":"Chinese name of attack type.", "zh_CN":"攻击类型中文。"}
  AttackTypeName *string `json:"attackTypeName,omitempty" xml:"attackTypeName,omitempty" require:"true"`
  // {"en":"Attack requests.", "zh_CN":"攻击请求数。"}
  TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty" require:"true"`
}

func (s GetDistributionOfWAFAttackTypeAttackTypeTop) String() string {
  return tea.Prettify(s)
}

func (s GetDistributionOfWAFAttackTypeAttackTypeTop) GoString() string {
  return s.String()
}

func (s *GetDistributionOfWAFAttackTypeAttackTypeTop) SetAttackTypeNameEn(v string) *GetDistributionOfWAFAttackTypeAttackTypeTop {
  s.AttackTypeNameEn = &v
  return s
}

func (s *GetDistributionOfWAFAttackTypeAttackTypeTop) SetAttackTypeName(v string) *GetDistributionOfWAFAttackTypeAttackTypeTop {
  s.AttackTypeName = &v
  return s
}

func (s *GetDistributionOfWAFAttackTypeAttackTypeTop) SetTotalCount(v int64) *GetDistributionOfWAFAttackTypeAttackTypeTop {
  s.TotalCount = &v
  return s
}

type GetDistributionOfWAFAttackTypePaths struct {
}

func (s GetDistributionOfWAFAttackTypePaths) String() string {
  return tea.Prettify(s)
}

func (s GetDistributionOfWAFAttackTypePaths) GoString() string {
  return s.String()
}

type GetDistributionOfWAFAttackTypeParameters struct {
}

func (s GetDistributionOfWAFAttackTypeParameters) String() string {
  return tea.Prettify(s)
}

func (s GetDistributionOfWAFAttackTypeParameters) GoString() string {
  return s.String()
}

type GetDistributionOfWAFAttackTypeRequestHeader struct {
}

func (s GetDistributionOfWAFAttackTypeRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetDistributionOfWAFAttackTypeRequestHeader) GoString() string {
  return s.String()
}

type GetDistributionOfWAFAttackTypeResponseHeader struct {
}

func (s GetDistributionOfWAFAttackTypeResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetDistributionOfWAFAttackTypeResponseHeader) GoString() string {
  return s.String()
}




type GetWAFTriggerRuleListRequest struct {
  // {"en":"Domain, array.", "zh_CN":"域名，数组。"}
  Domains []*string `json:"domains,omitempty" xml:"domains,omitempty" require:"true" type:"Repeated"`
  // {"en":"Start time, yyyy-MM-dd HH:mm:ss.", "zh_CN":"开始时间，yyyy-MM-dd HH:mm:ss。"}
  StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty" require:"true"`
  // {"en":"End time, yyyy-MM-dd HH:mm:ss.", "zh_CN":"结束时间，yyyy-MM-dd HH:mm:ss。"}
  EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty" require:"true"`
  // {"en":"Time zone, GMT+8 by default.", "zh_CN":"时区，默认GMT+8，即“GMT+8”。"}
  TimeZone *string `json:"timeZone,omitempty" xml:"timeZone,omitempty"`
  // {"en":"Action, default 1 and 2.
  //     1: Block
  //     2: Log", "zh_CN":"处理动作，默认1和2。
  //     1：拦截
  //     2：监控"}
  Acts []*string `json:"acts,omitempty" xml:"acts,omitempty" type:"Repeated"`
  // {"en":"Policy type, array. [protocol: Protocol Validation,
  // webShell: Webshell Access Detection,
  // other: Others Rules,
  // access: Access Control/Rate Limiting,
  // rule: WAF Rules]", "zh_CN":"触发策略类型，数组。[protocol：协议合规检测，
  // webShell：后门识别，
  // other：其他防护规则，
  // access：访问控制/限速，
  // rule：Web规则防护]"}
  Policys []*string `json:"policys,omitempty" xml:"policys,omitempty" type:"Repeated"`
}

func (s GetWAFTriggerRuleListRequest) String() string {
  return tea.Prettify(s)
}

func (s GetWAFTriggerRuleListRequest) GoString() string {
  return s.String()
}

func (s *GetWAFTriggerRuleListRequest) SetDomains(v []*string) *GetWAFTriggerRuleListRequest {
  s.Domains = v
  return s
}

func (s *GetWAFTriggerRuleListRequest) SetStartTime(v string) *GetWAFTriggerRuleListRequest {
  s.StartTime = &v
  return s
}

func (s *GetWAFTriggerRuleListRequest) SetEndTime(v string) *GetWAFTriggerRuleListRequest {
  s.EndTime = &v
  return s
}

func (s *GetWAFTriggerRuleListRequest) SetTimeZone(v string) *GetWAFTriggerRuleListRequest {
  s.TimeZone = &v
  return s
}

func (s *GetWAFTriggerRuleListRequest) SetActs(v []*string) *GetWAFTriggerRuleListRequest {
  s.Acts = v
  return s
}

func (s *GetWAFTriggerRuleListRequest) SetPolicys(v []*string) *GetWAFTriggerRuleListRequest {
  s.Policys = v
  return s
}

type GetWAFTriggerRuleListResponse struct {
  // {"en":"Return 0 means success, please see <Error code> to check other status code.", "zh_CN":"0状态码表示请求成功，其他状态码说明请参见《错误码》。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Error message or Success.", "zh_CN":"错误信息或Success。"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Return data.", "zh_CN":"返回值。"}
  Data *GetWAFTriggerRuleListTriggerRule `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s GetWAFTriggerRuleListResponse) String() string {
  return tea.Prettify(s)
}

func (s GetWAFTriggerRuleListResponse) GoString() string {
  return s.String()
}

func (s *GetWAFTriggerRuleListResponse) SetCode(v string) *GetWAFTriggerRuleListResponse {
  s.Code = &v
  return s
}

func (s *GetWAFTriggerRuleListResponse) SetMessage(v string) *GetWAFTriggerRuleListResponse {
  s.Message = &v
  return s
}

func (s *GetWAFTriggerRuleListResponse) SetData(v *GetWAFTriggerRuleListTriggerRule) *GetWAFTriggerRuleListResponse {
  s.Data = v
  return s
}

type GetWAFTriggerRuleListTriggerRule struct {
  // {"en":"Action.
  //     1: Block
  //     2: Log", "zh_CN":"处理动作。
  //     1：拦截
  //     2：监控"}
  Act *string `json:"act,omitempty" xml:"act,omitempty" require:"true"`
  // {"en":"Rule name.", "zh_CN":"规则名称。"}
  RuleName *string `json:"ruleName,omitempty" xml:"ruleName,omitempty" require:"true"`
  // {"en":"Attack type.", "zh_CN":"攻击类型。"}
  AttackType *string `json:"attackType,omitempty" xml:"attackType,omitempty" require:"true"`
  // {"en":"Chinese name of attack type.", "zh_CN":"攻击类型中文名称。"}
  AttackTypeName *string `json:"attackTypeName,omitempty" xml:"attackTypeName,omitempty" require:"true"`
  // {"en":"English name of attack type.", "zh_CN":"攻击类型英文名称。"}
  AttackTypeNameEn *string `json:"attackTypeNameEn,omitempty" xml:"attackTypeNameEn,omitempty" require:"true"`
  // {"en":"Rule ID.", "zh_CN":"规则ID。"}
  RuleId *string `json:"ruleId,omitempty" xml:"ruleId,omitempty" require:"true"`
  // {"en":"Trigger requests.", "zh_CN":"触发次数。"}
  TotalCount *string `json:"totalCount,omitempty" xml:"totalCount,omitempty" require:"true"`
}

func (s GetWAFTriggerRuleListTriggerRule) String() string {
  return tea.Prettify(s)
}

func (s GetWAFTriggerRuleListTriggerRule) GoString() string {
  return s.String()
}

func (s *GetWAFTriggerRuleListTriggerRule) SetAct(v string) *GetWAFTriggerRuleListTriggerRule {
  s.Act = &v
  return s
}

func (s *GetWAFTriggerRuleListTriggerRule) SetRuleName(v string) *GetWAFTriggerRuleListTriggerRule {
  s.RuleName = &v
  return s
}

func (s *GetWAFTriggerRuleListTriggerRule) SetAttackType(v string) *GetWAFTriggerRuleListTriggerRule {
  s.AttackType = &v
  return s
}

func (s *GetWAFTriggerRuleListTriggerRule) SetAttackTypeName(v string) *GetWAFTriggerRuleListTriggerRule {
  s.AttackTypeName = &v
  return s
}

func (s *GetWAFTriggerRuleListTriggerRule) SetAttackTypeNameEn(v string) *GetWAFTriggerRuleListTriggerRule {
  s.AttackTypeNameEn = &v
  return s
}

func (s *GetWAFTriggerRuleListTriggerRule) SetRuleId(v string) *GetWAFTriggerRuleListTriggerRule {
  s.RuleId = &v
  return s
}

func (s *GetWAFTriggerRuleListTriggerRule) SetTotalCount(v string) *GetWAFTriggerRuleListTriggerRule {
  s.TotalCount = &v
  return s
}

type GetWAFTriggerRuleListPaths struct {
}

func (s GetWAFTriggerRuleListPaths) String() string {
  return tea.Prettify(s)
}

func (s GetWAFTriggerRuleListPaths) GoString() string {
  return s.String()
}

type GetWAFTriggerRuleListParameters struct {
}

func (s GetWAFTriggerRuleListParameters) String() string {
  return tea.Prettify(s)
}

func (s GetWAFTriggerRuleListParameters) GoString() string {
  return s.String()
}

type GetWAFTriggerRuleListRequestHeader struct {
}

func (s GetWAFTriggerRuleListRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetWAFTriggerRuleListRequestHeader) GoString() string {
  return s.String()
}

type GetWAFTriggerRuleListResponseHeader struct {
}

func (s GetWAFTriggerRuleListResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetWAFTriggerRuleListResponseHeader) GoString() string {
  return s.String()
}




type GetWAFAttackedURLRequest struct {
  // {'en':'Domain, array.', 'zh_CN':'域名，数组。'}
  Domains []*string `json:"domains,omitempty" xml:"domains,omitempty" require:"true" type:"Repeated"`
  // {'en':'Start time, yyyy-MM-dd HH:mm:ss.', 'zh_CN':'开始时间，yyyy-MM-dd HH:mm:ss。'}
  StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty" require:"true"`
  // {'en':'End time, yyyy-MM-dd HH:mm:ss.', 'zh_CN':'结束时间，yyyy-MM-dd HH:mm:ss。'}
  EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty" require:"true"`
  // {'en':'Time zone, GMT+8 by default.', 'zh_CN':'时区，默认GMT+8，即“GMT+8”。'}
  TimeZone *string `json:"timeZone,omitempty" xml:"timeZone,omitempty"`
  // {'en':'Policy type, array. [protocol: Protocol Validation,
  // webShell: Webshell Access Detection,
  // other: Others Rules,
  // access: Access Control/Rate Limiting,
  // rule: WAF Rules]', 'zh_CN':'触发策略类型，数组 
  // [protocol：协议合规检测，
  // webShell：后门识别，
  // other：其他防护规则，
  // access：访问控制/限速，
  // rule：Web规则防护]'}
  Policys []*string `json:"policys,omitempty" xml:"policys,omitempty" type:"Repeated"`
}

func (s GetWAFAttackedURLRequest) String() string {
  return tea.Prettify(s)
}

func (s GetWAFAttackedURLRequest) GoString() string {
  return s.String()
}

func (s *GetWAFAttackedURLRequest) SetDomains(v []*string) *GetWAFAttackedURLRequest {
  s.Domains = v
  return s
}

func (s *GetWAFAttackedURLRequest) SetStartTime(v string) *GetWAFAttackedURLRequest {
  s.StartTime = &v
  return s
}

func (s *GetWAFAttackedURLRequest) SetEndTime(v string) *GetWAFAttackedURLRequest {
  s.EndTime = &v
  return s
}

func (s *GetWAFAttackedURLRequest) SetTimeZone(v string) *GetWAFAttackedURLRequest {
  s.TimeZone = &v
  return s
}

func (s *GetWAFAttackedURLRequest) SetPolicys(v []*string) *GetWAFAttackedURLRequest {
  s.Policys = v
  return s
}

type GetWAFAttackedURLResponse struct {
  // {"en":"Return 0 means success, please see <Error code> to check other status code.", "zh_CN":"0状态码表示请求成功，其他状态码说明请参见《错误码》。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Error message or Success.", "zh_CN":"错误信息或Success。"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Return data.", "zh_CN":"返回值。"}
  Data *GetWAFAttackedURLAttackedUrlList `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s GetWAFAttackedURLResponse) String() string {
  return tea.Prettify(s)
}

func (s GetWAFAttackedURLResponse) GoString() string {
  return s.String()
}

func (s *GetWAFAttackedURLResponse) SetCode(v string) *GetWAFAttackedURLResponse {
  s.Code = &v
  return s
}

func (s *GetWAFAttackedURLResponse) SetMessage(v string) *GetWAFAttackedURLResponse {
  s.Message = &v
  return s
}

func (s *GetWAFAttackedURLResponse) SetData(v *GetWAFAttackedURLAttackedUrlList) *GetWAFAttackedURLResponse {
  s.Data = v
  return s
}

type GetWAFAttackedURLAttackedUrlList struct {
  // {"en":"Total requests.", "zh_CN":"检测请求数。"}
  AllTotalCount *int64 `json:"allTotalCount,omitempty" xml:"allTotalCount,omitempty" require:"true"`
  // {"en":"Block requests.", "zh_CN":"拦截请求数。"}
  AllBlockCount *int64 `json:"allBlockCount,omitempty" xml:"allBlockCount,omitempty" require:"true"`
  // {"en":"Attacked url.", "zh_CN":"受攻击url。"}
  List []*GetWAFAttackedURLAttackedUrl `json:"list,omitempty" xml:"list,omitempty" require:"true" type:"Repeated"`
}

func (s GetWAFAttackedURLAttackedUrlList) String() string {
  return tea.Prettify(s)
}

func (s GetWAFAttackedURLAttackedUrlList) GoString() string {
  return s.String()
}

func (s *GetWAFAttackedURLAttackedUrlList) SetAllTotalCount(v int64) *GetWAFAttackedURLAttackedUrlList {
  s.AllTotalCount = &v
  return s
}

func (s *GetWAFAttackedURLAttackedUrlList) SetAllBlockCount(v int64) *GetWAFAttackedURLAttackedUrlList {
  s.AllBlockCount = &v
  return s
}

func (s *GetWAFAttackedURLAttackedUrlList) SetList(v []*GetWAFAttackedURLAttackedUrl) *GetWAFAttackedURLAttackedUrlList {
  s.List = v
  return s
}

type GetWAFAttackedURLAttackedUrl struct {
  // {"en":"Attacked url", "zh_CN":"受攻击url。"}
  Url *string `json:"url,omitempty" xml:"url,omitempty" require:"true"`
  // {"en":"Attack type top 1.", "zh_CN":"Top1攻击类型。"}
  AttackTypeTop1 *string `json:"attackTypeTop1,omitempty" xml:"attackTypeTop1,omitempty" require:"true"`
  // {"en":"Total requests of attacked url.", "zh_CN":"受攻击URL检测请求数。"}
  TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty" require:"true"`
  // {"en":"Block requests of attacked url.", "zh_CN":"受攻击URL拦截请求数。"}
  BlockCount *int64 `json:"blockCount,omitempty" xml:"blockCount,omitempty" require:"true"`
  // {"en":"Log requests of attacked url.", "zh_CN":"受攻击URL监控请求数。"}
  AlertCount *int64 `json:"alertCount,omitempty" xml:"alertCount,omitempty" require:"true"`
}

func (s GetWAFAttackedURLAttackedUrl) String() string {
  return tea.Prettify(s)
}

func (s GetWAFAttackedURLAttackedUrl) GoString() string {
  return s.String()
}

func (s *GetWAFAttackedURLAttackedUrl) SetUrl(v string) *GetWAFAttackedURLAttackedUrl {
  s.Url = &v
  return s
}

func (s *GetWAFAttackedURLAttackedUrl) SetAttackTypeTop1(v string) *GetWAFAttackedURLAttackedUrl {
  s.AttackTypeTop1 = &v
  return s
}

func (s *GetWAFAttackedURLAttackedUrl) SetTotalCount(v int64) *GetWAFAttackedURLAttackedUrl {
  s.TotalCount = &v
  return s
}

func (s *GetWAFAttackedURLAttackedUrl) SetBlockCount(v int64) *GetWAFAttackedURLAttackedUrl {
  s.BlockCount = &v
  return s
}

func (s *GetWAFAttackedURLAttackedUrl) SetAlertCount(v int64) *GetWAFAttackedURLAttackedUrl {
  s.AlertCount = &v
  return s
}

type GetWAFAttackedURLPaths struct {
}

func (s GetWAFAttackedURLPaths) String() string {
  return tea.Prettify(s)
}

func (s GetWAFAttackedURLPaths) GoString() string {
  return s.String()
}

type GetWAFAttackedURLParameters struct {
}

func (s GetWAFAttackedURLParameters) String() string {
  return tea.Prettify(s)
}

func (s GetWAFAttackedURLParameters) GoString() string {
  return s.String()
}

type GetWAFAttackedURLRequestHeader struct {
}

func (s GetWAFAttackedURLRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetWAFAttackedURLRequestHeader) GoString() string {
  return s.String()
}

type GetWAFAttackedURLResponseHeader struct {
}

func (s GetWAFAttackedURLResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetWAFAttackedURLResponseHeader) GoString() string {
  return s.String()
}




type GetTriggeredRateLimitingRulesRequest struct {
  // {"defaultValue":"10","en":"Top rankings, default value: 10, max: 1000.","zh_CN":"取前几排名，默认：10，上限：1000。"}
  Top *int `json:"top,omitempty" xml:"top,omitempty"`
  // {"en":"Multiple selection. Handling results, default: display all results.\nmitigated: Number of mitigated requests.\nmonitored: Number of observed requests.","zh_CN":"多选。处置结果，默认：展示所有结果。\nmitigated：已抵御请求数。\nmonitored：观察请求数。","exampleValue":"mitigated,monitored"}
  ActType []*string `json:"actType,omitempty" xml:"actType,omitempty" type:"Repeated"`
  // {"en":"Start time, format: yyyy-MM-dd HH:mm:ss.\nThe time range can not exceed 31 days.","zh_CN":"起始时间，格式：yyyy-MM-dd HH:mm:ss。\n时间范围不超过31天。"}
  StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty" require:"true"`
  // {"en":"End time, format: yyyy-MM-dd HH:mm:ss.\nThe time range can not exceed 31 days.","zh_CN":"截止时间，格式：yyyy-MM-dd HH:mm:ss。\n时间范围不超过31天。"}
  EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty" require:"true"`
  // {"en":"Domain list. Queries all domains under the account when not specified.","zh_CN":"域名列表，未指定时查询账号下的所有域名。"}
  Domains []*string `json:"domains,omitempty" xml:"domains,omitempty" type:"Repeated"`
}

func (s GetTriggeredRateLimitingRulesRequest) String() string {
  return tea.Prettify(s)
}

func (s GetTriggeredRateLimitingRulesRequest) GoString() string {
  return s.String()
}

func (s *GetTriggeredRateLimitingRulesRequest) SetTop(v int) *GetTriggeredRateLimitingRulesRequest {
  s.Top = &v
  return s
}

func (s *GetTriggeredRateLimitingRulesRequest) SetActType(v []*string) *GetTriggeredRateLimitingRulesRequest {
  s.ActType = v
  return s
}

func (s *GetTriggeredRateLimitingRulesRequest) SetStartTime(v string) *GetTriggeredRateLimitingRulesRequest {
  s.StartTime = &v
  return s
}

func (s *GetTriggeredRateLimitingRulesRequest) SetEndTime(v string) *GetTriggeredRateLimitingRulesRequest {
  s.EndTime = &v
  return s
}

func (s *GetTriggeredRateLimitingRulesRequest) SetDomains(v []*string) *GetTriggeredRateLimitingRulesRequest {
  s.Domains = v
  return s
}

type GetTriggeredRateLimitingRulesRequestHeader struct {
  // {"defaultValue":"en","en":"The language of response data, default value: en.\nzh_CN: Chinese\nen: English","zh_CN":"返回内容的语言版本，默认值: en。\nzh_CN：中文\nen：英文","exampleValue":"zh_CN,en"}
  Language *string `json:"language,omitempty" xml:"language,omitempty"`
  // {"dictionary":"belong=WAAP-MS-Ext|dict=waap_serviceType","en":"Security service type. Please enter a specific service type, if you purchase multiple security services.","zh_CN":"安全服务类型。有使用多个不同的安全服务时，需要填写具体的服务类型。"}
  ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty"`
  // {"defaultValue":"GMT+8","en":"Report Data Timezone:1. Indicates the timezone for the report data. lt must be relative to GMT andspecified in the format GMT+n where -12<=n <= 12. For example,timezone=GMT-7 and timezone=GMT+8.2. If the parameter is unspecified,results will be in the GMT+8 timezone.","zh_CN":"报表数据时区：1.请传递GMT时区。示例格式：GMT+N（其中12<=N <= 12），示例：timezone=GMT-7或timezone=GMT+8.2.若参数为传递，将默认按GMT+8时区查询"}
  Timezone *string `json:"timezone,omitempty" xml:"timezone,omitempty"`
}

func (s GetTriggeredRateLimitingRulesRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetTriggeredRateLimitingRulesRequestHeader) GoString() string {
  return s.String()
}

func (s *GetTriggeredRateLimitingRulesRequestHeader) SetLanguage(v string) *GetTriggeredRateLimitingRulesRequestHeader {
  s.Language = &v
  return s
}

func (s *GetTriggeredRateLimitingRulesRequestHeader) SetServiceType(v string) *GetTriggeredRateLimitingRulesRequestHeader {
  s.ServiceType = &v
  return s
}

func (s *GetTriggeredRateLimitingRulesRequestHeader) SetTimezone(v string) *GetTriggeredRateLimitingRulesRequestHeader {
  s.Timezone = &v
  return s
}

type GetTriggeredRateLimitingRulesPaths struct {
}

func (s GetTriggeredRateLimitingRulesPaths) String() string {
  return tea.Prettify(s)
}

func (s GetTriggeredRateLimitingRulesPaths) GoString() string {
  return s.String()
}

type GetTriggeredRateLimitingRulesParameters struct {
}

func (s GetTriggeredRateLimitingRulesParameters) String() string {
  return tea.Prettify(s)
}

func (s GetTriggeredRateLimitingRulesParameters) GoString() string {
  return s.String()
}

type GetTriggeredRateLimitingRulesResponse struct {
  // {"dictionary":"belong=WAAP-MS-Ext|dict=waap_retCodeEnum","en":"Please refer to the error code for exceptions.","zh_CN":"请参照错误码。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Description.","zh_CN":"描述信息。"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"Data.","zh_CN":"出参数据。"}
  Data []*GetTriggeredRateLimitingRulesResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s GetTriggeredRateLimitingRulesResponse) String() string {
  return tea.Prettify(s)
}

func (s GetTriggeredRateLimitingRulesResponse) GoString() string {
  return s.String()
}

func (s *GetTriggeredRateLimitingRulesResponse) SetCode(v string) *GetTriggeredRateLimitingRulesResponse {
  s.Code = &v
  return s
}

func (s *GetTriggeredRateLimitingRulesResponse) SetMsg(v string) *GetTriggeredRateLimitingRulesResponse {
  s.Msg = &v
  return s
}

func (s *GetTriggeredRateLimitingRulesResponse) SetData(v []*GetTriggeredRateLimitingRulesResponseData) *GetTriggeredRateLimitingRulesResponse {
  s.Data = v
  return s
}

type GetTriggeredRateLimitingRulesResponseData struct     {
  // {"en":"Rule ID.","zh_CN":"规则ID。"}
  RuleId *string `json:"ruleId,omitempty" xml:"ruleId,omitempty" require:"true"`
  // {"en":"Rule name.","zh_CN":"规则名称。"}
  RuleName *string `json:"ruleName,omitempty" xml:"ruleName,omitempty" require:"true"`
  // {"en":"Protected target.","zh_CN":"业务场景。"}
  Scene *string `json:"scene,omitempty" xml:"scene,omitempty" require:"true"`
  // {"en":"Trigger times, sort by times.","zh_CN":"触发次数，按次数排序。"}
  Hits []*GetTriggeredRateLimitingRulesResponseDataHits `json:"hits,omitempty" xml:"hits,omitempty" require:"true" type:"Repeated"`
}

func (s GetTriggeredRateLimitingRulesResponseData) String() string {
  return tea.Prettify(s)
}

func (s GetTriggeredRateLimitingRulesResponseData) GoString() string {
  return s.String()
}

func (s *GetTriggeredRateLimitingRulesResponseData) SetRuleId(v string) *GetTriggeredRateLimitingRulesResponseData {
  s.RuleId = &v
  return s
}

func (s *GetTriggeredRateLimitingRulesResponseData) SetRuleName(v string) *GetTriggeredRateLimitingRulesResponseData {
  s.RuleName = &v
  return s
}

func (s *GetTriggeredRateLimitingRulesResponseData) SetScene(v string) *GetTriggeredRateLimitingRulesResponseData {
  s.Scene = &v
  return s
}

func (s *GetTriggeredRateLimitingRulesResponseData) SetHits(v []*GetTriggeredRateLimitingRulesResponseDataHits) *GetTriggeredRateLimitingRulesResponseData {
  s.Hits = v
  return s
}

type GetTriggeredRateLimitingRulesResponseDataHits struct     {
  // {"en":"Action.","zh_CN":"采取动作。"}
  Act *string `json:"act,omitempty" xml:"act,omitempty" require:"true"`
  // {"en":"Hit times.","zh_CN":"命中数。"}
  Value *int64 `json:"value,omitempty" xml:"value,omitempty" require:"true"`
}

func (s GetTriggeredRateLimitingRulesResponseDataHits) String() string {
  return tea.Prettify(s)
}

func (s GetTriggeredRateLimitingRulesResponseDataHits) GoString() string {
  return s.String()
}

func (s *GetTriggeredRateLimitingRulesResponseDataHits) SetAct(v string) *GetTriggeredRateLimitingRulesResponseDataHits {
  s.Act = &v
  return s
}

func (s *GetTriggeredRateLimitingRulesResponseDataHits) SetValue(v int64) *GetTriggeredRateLimitingRulesResponseDataHits {
  s.Value = &v
  return s
}

type GetTriggeredRateLimitingRulesResponseHeader struct {
}

func (s GetTriggeredRateLimitingRulesResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetTriggeredRateLimitingRulesResponseHeader) GoString() string {
  return s.String()
}




type GetConsumerNumberRequest struct {
  // {"en":"List of API group.", "zh_CN":"API分组列表。"}
  ApiGroups []*string `json:"apiGroups,omitempty" xml:"apiGroups,omitempty" type:"Repeated"`
  // {"en":"List of API name.", "zh_CN":"API名称列表。"}
  ApiIds []*string `json:"apiIds,omitempty" xml:"apiIds,omitempty" type:"Repeated"`
  // {"en":"List of domain.", "zh_CN":"域名列表。"}
  Domains []*string `json:"domains,omitempty" xml:"domains,omitempty" type:"Repeated"`
  // {"en":"Start time, format: yyyy-MM-dd HH:mm:ss.", "zh_CN":"开始时间，格式：yyyy-MM-dd HH:mm:ss。"}
  StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty" require:"true"`
  // {"en":"End time, format: yyyy-MM-dd HH:mm:ss.", "zh_CN":"结束时间，格式：yyyy-MM-dd HH:mm:ss。"}
  EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty" require:"true"`
  // {"en":"Endpoint path.", "zh_CN":"前端路径。"}
  FrontPath *string `json:"frontPath,omitempty" xml:"frontPath,omitempty"`
}

func (s GetConsumerNumberRequest) String() string {
  return tea.Prettify(s)
}

func (s GetConsumerNumberRequest) GoString() string {
  return s.String()
}

func (s *GetConsumerNumberRequest) SetApiGroups(v []*string) *GetConsumerNumberRequest {
  s.ApiGroups = v
  return s
}

func (s *GetConsumerNumberRequest) SetApiIds(v []*string) *GetConsumerNumberRequest {
  s.ApiIds = v
  return s
}

func (s *GetConsumerNumberRequest) SetDomains(v []*string) *GetConsumerNumberRequest {
  s.Domains = v
  return s
}

func (s *GetConsumerNumberRequest) SetStartTime(v string) *GetConsumerNumberRequest {
  s.StartTime = &v
  return s
}

func (s *GetConsumerNumberRequest) SetEndTime(v string) *GetConsumerNumberRequest {
  s.EndTime = &v
  return s
}

func (s *GetConsumerNumberRequest) SetFrontPath(v string) *GetConsumerNumberRequest {
  s.FrontPath = &v
  return s
}

type GetConsumerNumberResponse struct {
  // {"en":"Return 200 means success.", "zh_CN":"200状态码表示请求成功。"}
  Code *int `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Message.", "zh_CN":"返回信息。"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"Data.", "zh_CN":"数据。"}
  Data *GetConsumerNumberVo `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s GetConsumerNumberResponse) String() string {
  return tea.Prettify(s)
}

func (s GetConsumerNumberResponse) GoString() string {
  return s.String()
}

func (s *GetConsumerNumberResponse) SetCode(v int) *GetConsumerNumberResponse {
  s.Code = &v
  return s
}

func (s *GetConsumerNumberResponse) SetMsg(v string) *GetConsumerNumberResponse {
  s.Msg = &v
  return s
}

func (s *GetConsumerNumberResponse) SetData(v *GetConsumerNumberVo) *GetConsumerNumberResponse {
  s.Data = v
  return s
}

type GetConsumerNumberVo struct {
  // {"en":"The total number of consumers.", "zh_CN":"消费方总数。"}
  Total *int `json:"total,omitempty" xml:"total,omitempty" require:"true"`
  // {"en":"The total number of authorized consumers.", "zh_CN":"已授权消费方总数。"}
  Auth *int `json:"auth,omitempty" xml:"auth,omitempty" require:"true"`
}

func (s GetConsumerNumberVo) String() string {
  return tea.Prettify(s)
}

func (s GetConsumerNumberVo) GoString() string {
  return s.String()
}

func (s *GetConsumerNumberVo) SetTotal(v int) *GetConsumerNumberVo {
  s.Total = &v
  return s
}

func (s *GetConsumerNumberVo) SetAuth(v int) *GetConsumerNumberVo {
  s.Auth = &v
  return s
}

type GetConsumerNumberPaths struct {
}

func (s GetConsumerNumberPaths) String() string {
  return tea.Prettify(s)
}

func (s GetConsumerNumberPaths) GoString() string {
  return s.String()
}

type GetConsumerNumberParameters struct {
}

func (s GetConsumerNumberParameters) String() string {
  return tea.Prettify(s)
}

func (s GetConsumerNumberParameters) GoString() string {
  return s.String()
}

type GetConsumerNumberRequestHeader struct {
}

func (s GetConsumerNumberRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetConsumerNumberRequestHeader) GoString() string {
  return s.String()
}

type GetConsumerNumberResponseHeader struct {
}

func (s GetConsumerNumberResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetConsumerNumberResponseHeader) GoString() string {
  return s.String()
}




type GetTriggeredWAFManagedRulesRequest struct {
  // {"defaultValue":"10","en":"Top rankings, default value: 10, max: 1000.","zh_CN":"取前几排名，默认：10，上限：1000。"}
  Top *int `json:"top,omitempty" xml:"top,omitempty"`
  // {"en":"Multiple selection. Handling results, default: display all results.\nmitigated: Number of mitigated requests.\nmonitored: Number of observed requests.","zh_CN":"多选。处置结果，默认：展示所有结果。\nmitigated：已抵御请求数。\nmonitored：观察请求数。","exampleValue":"mitigated,monitored"}
  ActType []*string `json:"actType,omitempty" xml:"actType,omitempty" type:"Repeated"`
  // {"en":"Start time, format: yyyy-MM-dd HH:mm:ss.\nThe time range can not exceed 31 days.","zh_CN":"起始时间，格式：yyyy-MM-dd HH:mm:ss。\n时间范围不超过31天。"}
  StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty" require:"true"`
  // {"en":"End time, format: yyyy-MM-dd HH:mm:ss.\nThe time range can not exceed 31 days.","zh_CN":"截止时间，格式：yyyy-MM-dd HH:mm:ss。\n时间范围不超过31天。"}
  EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty" require:"true"`
  // {"en":"Domain list. Queries all domains under the account when not specified.","zh_CN":"域名列表，未指定时查询账号下的所有域名。"}
  Domains []*string `json:"domains,omitempty" xml:"domains,omitempty" type:"Repeated"`
}

func (s GetTriggeredWAFManagedRulesRequest) String() string {
  return tea.Prettify(s)
}

func (s GetTriggeredWAFManagedRulesRequest) GoString() string {
  return s.String()
}

func (s *GetTriggeredWAFManagedRulesRequest) SetTop(v int) *GetTriggeredWAFManagedRulesRequest {
  s.Top = &v
  return s
}

func (s *GetTriggeredWAFManagedRulesRequest) SetActType(v []*string) *GetTriggeredWAFManagedRulesRequest {
  s.ActType = v
  return s
}

func (s *GetTriggeredWAFManagedRulesRequest) SetStartTime(v string) *GetTriggeredWAFManagedRulesRequest {
  s.StartTime = &v
  return s
}

func (s *GetTriggeredWAFManagedRulesRequest) SetEndTime(v string) *GetTriggeredWAFManagedRulesRequest {
  s.EndTime = &v
  return s
}

func (s *GetTriggeredWAFManagedRulesRequest) SetDomains(v []*string) *GetTriggeredWAFManagedRulesRequest {
  s.Domains = v
  return s
}

type GetTriggeredWAFManagedRulesRequestHeader struct {
  // {"defaultValue":"en","en":"The language of response data, default value: en.\nzh_CN: Chinese\nen: English","zh_CN":"返回内容的语言版本，默认值: en。\nzh_CN：中文\nen：英文","exampleValue":"zh_CN,en"}
  Language *string `json:"language,omitempty" xml:"language,omitempty"`
  // {"dictionary":"belong=WAAP-MS-Ext|dict=waap_serviceType","en":"Security service type. Please enter a specific service type, if you purchase multiple security services.","zh_CN":"安全服务类型。有使用多个不同的安全服务时，需要填写具体的服务类型。"}
  ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty"`
  // {"defaultValue":"GMT+8","en":"Report Data Timezone:1. Indicates the timezone for the report data. lt must be relative to GMT andspecified in the format GMT+n where -12<=n <= 12. For example,timezone=GMT-7 and timezone=GMT+8.2. If the parameter is unspecified,results will be in the GMT+8 timezone.","zh_CN":"报表数据时区：1.请传递GMT时区。示例格式：GMT+N（其中12<=N <= 12），示例：timezone=GMT-7或timezone=GMT+8.2.若参数为传递，将默认按GMT+8时区查询"}
  Timezone *string `json:"timezone,omitempty" xml:"timezone,omitempty"`
}

func (s GetTriggeredWAFManagedRulesRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetTriggeredWAFManagedRulesRequestHeader) GoString() string {
  return s.String()
}

func (s *GetTriggeredWAFManagedRulesRequestHeader) SetLanguage(v string) *GetTriggeredWAFManagedRulesRequestHeader {
  s.Language = &v
  return s
}

func (s *GetTriggeredWAFManagedRulesRequestHeader) SetServiceType(v string) *GetTriggeredWAFManagedRulesRequestHeader {
  s.ServiceType = &v
  return s
}

func (s *GetTriggeredWAFManagedRulesRequestHeader) SetTimezone(v string) *GetTriggeredWAFManagedRulesRequestHeader {
  s.Timezone = &v
  return s
}

type GetTriggeredWAFManagedRulesPaths struct {
}

func (s GetTriggeredWAFManagedRulesPaths) String() string {
  return tea.Prettify(s)
}

func (s GetTriggeredWAFManagedRulesPaths) GoString() string {
  return s.String()
}

type GetTriggeredWAFManagedRulesParameters struct {
}

func (s GetTriggeredWAFManagedRulesParameters) String() string {
  return tea.Prettify(s)
}

func (s GetTriggeredWAFManagedRulesParameters) GoString() string {
  return s.String()
}

type GetTriggeredWAFManagedRulesResponse struct {
  // {"dictionary":"belong=WAAP-MS-Ext|dict=waap_retCodeEnum","en":"Please refer to the error code for exceptions.","zh_CN":"请参照错误码。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Description.","zh_CN":"描述信息。"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"Data.","zh_CN":"出参数据。"}
  Data []*GetTriggeredWAFManagedRulesResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s GetTriggeredWAFManagedRulesResponse) String() string {
  return tea.Prettify(s)
}

func (s GetTriggeredWAFManagedRulesResponse) GoString() string {
  return s.String()
}

func (s *GetTriggeredWAFManagedRulesResponse) SetCode(v string) *GetTriggeredWAFManagedRulesResponse {
  s.Code = &v
  return s
}

func (s *GetTriggeredWAFManagedRulesResponse) SetMsg(v string) *GetTriggeredWAFManagedRulesResponse {
  s.Msg = &v
  return s
}

func (s *GetTriggeredWAFManagedRulesResponse) SetData(v []*GetTriggeredWAFManagedRulesResponseData) *GetTriggeredWAFManagedRulesResponse {
  s.Data = v
  return s
}

type GetTriggeredWAFManagedRulesResponseData struct     {
  // {"en":"Rule ID.","zh_CN":"规则ID。"}
  RuleId *string `json:"ruleId,omitempty" xml:"ruleId,omitempty" require:"true"`
  // {"en":"Rule name.","zh_CN":"规则名称。"}
  RuleName *string `json:"ruleName,omitempty" xml:"ruleName,omitempty" require:"true"`
  // {"en":"Rule type.","zh_CN":"规则类型。"}
  RuleType *string `json:"ruleType,omitempty" xml:"ruleType,omitempty" require:"true"`
  // {"en":"System recommended action.","zh_CN":"系统推荐动作。"}
  Action *string `json:"action,omitempty" xml:"action,omitempty" require:"true"`
  // {"en":"Trigger times, sort by times.","zh_CN":"触发次数，按次数排序。"}
  Hits []*GetTriggeredWAFManagedRulesResponseDataHits `json:"hits,omitempty" xml:"hits,omitempty" require:"true" type:"Repeated"`
}

func (s GetTriggeredWAFManagedRulesResponseData) String() string {
  return tea.Prettify(s)
}

func (s GetTriggeredWAFManagedRulesResponseData) GoString() string {
  return s.String()
}

func (s *GetTriggeredWAFManagedRulesResponseData) SetRuleId(v string) *GetTriggeredWAFManagedRulesResponseData {
  s.RuleId = &v
  return s
}

func (s *GetTriggeredWAFManagedRulesResponseData) SetRuleName(v string) *GetTriggeredWAFManagedRulesResponseData {
  s.RuleName = &v
  return s
}

func (s *GetTriggeredWAFManagedRulesResponseData) SetRuleType(v string) *GetTriggeredWAFManagedRulesResponseData {
  s.RuleType = &v
  return s
}

func (s *GetTriggeredWAFManagedRulesResponseData) SetAction(v string) *GetTriggeredWAFManagedRulesResponseData {
  s.Action = &v
  return s
}

func (s *GetTriggeredWAFManagedRulesResponseData) SetHits(v []*GetTriggeredWAFManagedRulesResponseDataHits) *GetTriggeredWAFManagedRulesResponseData {
  s.Hits = v
  return s
}

type GetTriggeredWAFManagedRulesResponseDataHits struct     {
  // {"en":"Action.","zh_CN":"采取动作。"}
  Act *string `json:"act,omitempty" xml:"act,omitempty" require:"true"`
  // {"en":"Hit times.","zh_CN":"命中数。"}
  Value *int64 `json:"value,omitempty" xml:"value,omitempty" require:"true"`
}

func (s GetTriggeredWAFManagedRulesResponseDataHits) String() string {
  return tea.Prettify(s)
}

func (s GetTriggeredWAFManagedRulesResponseDataHits) GoString() string {
  return s.String()
}

func (s *GetTriggeredWAFManagedRulesResponseDataHits) SetAct(v string) *GetTriggeredWAFManagedRulesResponseDataHits {
  s.Act = &v
  return s
}

func (s *GetTriggeredWAFManagedRulesResponseDataHits) SetValue(v int64) *GetTriggeredWAFManagedRulesResponseDataHits {
  s.Value = &v
  return s
}

type GetTriggeredWAFManagedRulesResponseHeader struct {
}

func (s GetTriggeredWAFManagedRulesResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetTriggeredWAFManagedRulesResponseHeader) GoString() string {
  return s.String()
}




type GetBotRequestStatisticPerDomainRequest struct {
  // {"en":"Domain.Separate by';'.", "zh_CN":"域名串，分号拼接"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {"en":"StartTime. Format: yyyy-MM-dd HH:mm:ss", "zh_CN":"开始时间。格式： yyyy-MM-dd HH:mm:ss"}
  StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty" require:"true"`
  // {"en":"EndTime. Format: yyyy-MM-dd HH:mm:ss", "zh_CN":"结束时间。格式： yyyy-MM-dd HH:mm:ss"}
  EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty" require:"true"`
  // {"en":"Time zone, default 8, i.e.'GTM+8'", "zh_CN":"时区，默认8，即“GTM+8”"}
  TimeZone *int `json:"timeZone,omitempty" xml:"timeZone,omitempty"`
  // {"en":"Number of top values.Default value 10", "zh_CN":"排名最前值数目。默认10"}
  TopNum *int `json:"topNum,omitempty" xml:"topNum,omitempty"`
}

func (s GetBotRequestStatisticPerDomainRequest) String() string {
  return tea.Prettify(s)
}

func (s GetBotRequestStatisticPerDomainRequest) GoString() string {
  return s.String()
}

func (s *GetBotRequestStatisticPerDomainRequest) SetDomain(v string) *GetBotRequestStatisticPerDomainRequest {
  s.Domain = &v
  return s
}

func (s *GetBotRequestStatisticPerDomainRequest) SetStartTime(v string) *GetBotRequestStatisticPerDomainRequest {
  s.StartTime = &v
  return s
}

func (s *GetBotRequestStatisticPerDomainRequest) SetEndTime(v string) *GetBotRequestStatisticPerDomainRequest {
  s.EndTime = &v
  return s
}

func (s *GetBotRequestStatisticPerDomainRequest) SetTimeZone(v int) *GetBotRequestStatisticPerDomainRequest {
  s.TimeZone = &v
  return s
}

func (s *GetBotRequestStatisticPerDomainRequest) SetTopNum(v int) *GetBotRequestStatisticPerDomainRequest {
  s.TopNum = &v
  return s
}

type GetBotRequestStatisticPerDomainResponse struct {
  // {"en":"Status code, success is '200'.", "zh_CN":"状态码，成功为“200”。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Return message, success is 'Success'.", "zh_CN":"返回信息，成功为“Success”。"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Data returned.", "zh_CN":"返回数据。"}
  Data []*GetBotRequestStatisticPerDomainResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s GetBotRequestStatisticPerDomainResponse) String() string {
  return tea.Prettify(s)
}

func (s GetBotRequestStatisticPerDomainResponse) GoString() string {
  return s.String()
}

func (s *GetBotRequestStatisticPerDomainResponse) SetCode(v string) *GetBotRequestStatisticPerDomainResponse {
  s.Code = &v
  return s
}

func (s *GetBotRequestStatisticPerDomainResponse) SetMessage(v string) *GetBotRequestStatisticPerDomainResponse {
  s.Message = &v
  return s
}

func (s *GetBotRequestStatisticPerDomainResponse) SetData(v []*GetBotRequestStatisticPerDomainResponseData) *GetBotRequestStatisticPerDomainResponse {
  s.Data = v
  return s
}

type GetBotRequestStatisticPerDomainResponseData struct     {
  // {"en":"Statistical type.", "zh_CN":"统计类型。"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"Bot request count.", "zh_CN":"Bot请求数。"}
  Value *int64 `json:"value,omitempty" xml:"value,omitempty" require:"true"`
}

func (s GetBotRequestStatisticPerDomainResponseData) String() string {
  return tea.Prettify(s)
}

func (s GetBotRequestStatisticPerDomainResponseData) GoString() string {
  return s.String()
}

func (s *GetBotRequestStatisticPerDomainResponseData) SetName(v string) *GetBotRequestStatisticPerDomainResponseData {
  s.Name = &v
  return s
}

func (s *GetBotRequestStatisticPerDomainResponseData) SetValue(v int64) *GetBotRequestStatisticPerDomainResponseData {
  s.Value = &v
  return s
}

type GetBotRequestStatisticPerDomainPaths struct {
}

func (s GetBotRequestStatisticPerDomainPaths) String() string {
  return tea.Prettify(s)
}

func (s GetBotRequestStatisticPerDomainPaths) GoString() string {
  return s.String()
}

type GetBotRequestStatisticPerDomainParameters struct {
}

func (s GetBotRequestStatisticPerDomainParameters) String() string {
  return tea.Prettify(s)
}

func (s GetBotRequestStatisticPerDomainParameters) GoString() string {
  return s.String()
}

type GetBotRequestStatisticPerDomainRequestHeader struct {
}

func (s GetBotRequestStatisticPerDomainRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetBotRequestStatisticPerDomainRequestHeader) GoString() string {
  return s.String()
}

type GetBotRequestStatisticPerDomainResponseHeader struct {
}

func (s GetBotRequestStatisticPerDomainResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetBotRequestStatisticPerDomainResponseHeader) GoString() string {
  return s.String()
}




type GetBotRequestTrendsAndTriggerRulesDataRequest struct {
  // {"en":"Domain.Separate by';'.", "zh_CN":"域名。多个以;隔开"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {"en":"StartTime.Format: yyyy-MM-dd HH:mm:ss", "zh_CN":"开始时间。格式： yyyy-MM-dd HH:mm:ss"}
  StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty" require:"true"`
  // {"en":"EndTime. format: yyyy-MM-dd HH:mm:ss", "zh_CN":"结束时间。格式： yyyy-MM-dd HH:mm:ss"}
  EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty" require:"true"`
  // {"en":"Time zone. Default 8, i.e. 'GTM+8'", "zh_CN":"时区，默认8，即“GTM+8”"}
  TimeZone *int `json:"timeZone,omitempty" xml:"timeZone,omitempty"`
  // {"en":"Language type. Default cn. 
  //  en:English 
  //  cn:Chinese ", "zh_CN":"语言类型。 默认cn 
  //  en：英文 
  //  cn：中文"}
  Lang *string `json:"lang,omitempty" xml:"lang,omitempty"`
}

func (s GetBotRequestTrendsAndTriggerRulesDataRequest) String() string {
  return tea.Prettify(s)
}

func (s GetBotRequestTrendsAndTriggerRulesDataRequest) GoString() string {
  return s.String()
}

func (s *GetBotRequestTrendsAndTriggerRulesDataRequest) SetDomain(v string) *GetBotRequestTrendsAndTriggerRulesDataRequest {
  s.Domain = &v
  return s
}

func (s *GetBotRequestTrendsAndTriggerRulesDataRequest) SetStartTime(v string) *GetBotRequestTrendsAndTriggerRulesDataRequest {
  s.StartTime = &v
  return s
}

func (s *GetBotRequestTrendsAndTriggerRulesDataRequest) SetEndTime(v string) *GetBotRequestTrendsAndTriggerRulesDataRequest {
  s.EndTime = &v
  return s
}

func (s *GetBotRequestTrendsAndTriggerRulesDataRequest) SetTimeZone(v int) *GetBotRequestTrendsAndTriggerRulesDataRequest {
  s.TimeZone = &v
  return s
}

func (s *GetBotRequestTrendsAndTriggerRulesDataRequest) SetLang(v string) *GetBotRequestTrendsAndTriggerRulesDataRequest {
  s.Lang = &v
  return s
}

type GetBotRequestTrendsAndTriggerRulesDataResponse struct {
  // {"en":"Status code, success is '200'.", "zh_CN":"状态码，成功为“200”。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Return message, success is 'Success'.", "zh_CN":"返回信息，成功为“Success”。"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Data returned.", "zh_CN":"返回数据。"}
  Data *GetBotRequestTrendsAndTriggerRulesDataResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
}

func (s GetBotRequestTrendsAndTriggerRulesDataResponse) String() string {
  return tea.Prettify(s)
}

func (s GetBotRequestTrendsAndTriggerRulesDataResponse) GoString() string {
  return s.String()
}

func (s *GetBotRequestTrendsAndTriggerRulesDataResponse) SetCode(v string) *GetBotRequestTrendsAndTriggerRulesDataResponse {
  s.Code = &v
  return s
}

func (s *GetBotRequestTrendsAndTriggerRulesDataResponse) SetMessage(v string) *GetBotRequestTrendsAndTriggerRulesDataResponse {
  s.Message = &v
  return s
}

func (s *GetBotRequestTrendsAndTriggerRulesDataResponse) SetData(v *GetBotRequestTrendsAndTriggerRulesDataResponseData) *GetBotRequestTrendsAndTriggerRulesDataResponse {
  s.Data = v
  return s
}

type GetBotRequestTrendsAndTriggerRulesDataResponseData struct {
  // {"en":"Request Trend.", "zh_CN":"请求趋势。"}
  RequestTrend []*GetBotRequestTrendsAndTriggerRulesDataResponseDataRequestTrend `json:"requestTrend,omitempty" xml:"requestTrend,omitempty" require:"true" type:"Repeated"`
  // {"en":"Trigger Rule.", "zh_CN":"触发规则。"}
  TriggerRule []*GetBotRequestTrendsAndTriggerRulesDataResponseDataTriggerRule `json:"triggerRule,omitempty" xml:"triggerRule,omitempty" require:"true" type:"Repeated"`
}

func (s GetBotRequestTrendsAndTriggerRulesDataResponseData) String() string {
  return tea.Prettify(s)
}

func (s GetBotRequestTrendsAndTriggerRulesDataResponseData) GoString() string {
  return s.String()
}

func (s *GetBotRequestTrendsAndTriggerRulesDataResponseData) SetRequestTrend(v []*GetBotRequestTrendsAndTriggerRulesDataResponseDataRequestTrend) *GetBotRequestTrendsAndTriggerRulesDataResponseData {
  s.RequestTrend = v
  return s
}

func (s *GetBotRequestTrendsAndTriggerRulesDataResponseData) SetTriggerRule(v []*GetBotRequestTrendsAndTriggerRulesDataResponseDataTriggerRule) *GetBotRequestTrendsAndTriggerRulesDataResponseData {
  s.TriggerRule = v
  return s
}

type GetBotRequestTrendsAndTriggerRulesDataResponseDataRequestTrend struct     {
  // {"en":"Request time.", "zh_CN":"请求时间。"}
  Time *string `json:"time,omitempty" xml:"time,omitempty" require:"true"`
  // {"en":"Trend detail.", "zh_CN":"趋势详情。"}
  Detail []*GetBotRequestTrendsAndTriggerRulesDataResponseDataRequestTrendDetail `json:"detail,omitempty" xml:"detail,omitempty" require:"true" type:"Repeated"`
}

func (s GetBotRequestTrendsAndTriggerRulesDataResponseDataRequestTrend) String() string {
  return tea.Prettify(s)
}

func (s GetBotRequestTrendsAndTriggerRulesDataResponseDataRequestTrend) GoString() string {
  return s.String()
}

func (s *GetBotRequestTrendsAndTriggerRulesDataResponseDataRequestTrend) SetTime(v string) *GetBotRequestTrendsAndTriggerRulesDataResponseDataRequestTrend {
  s.Time = &v
  return s
}

func (s *GetBotRequestTrendsAndTriggerRulesDataResponseDataRequestTrend) SetDetail(v []*GetBotRequestTrendsAndTriggerRulesDataResponseDataRequestTrendDetail) *GetBotRequestTrendsAndTriggerRulesDataResponseDataRequestTrend {
  s.Detail = v
  return s
}

type GetBotRequestTrendsAndTriggerRulesDataResponseDataRequestTrendDetail struct     {
  // {"en":"Statistical type.", "zh_CN":"统计类型。"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"Request times.", "zh_CN":"请求数。"}
  Value *int64 `json:"value,omitempty" xml:"value,omitempty" require:"true"`
}

func (s GetBotRequestTrendsAndTriggerRulesDataResponseDataRequestTrendDetail) String() string {
  return tea.Prettify(s)
}

func (s GetBotRequestTrendsAndTriggerRulesDataResponseDataRequestTrendDetail) GoString() string {
  return s.String()
}

func (s *GetBotRequestTrendsAndTriggerRulesDataResponseDataRequestTrendDetail) SetName(v string) *GetBotRequestTrendsAndTriggerRulesDataResponseDataRequestTrendDetail {
  s.Name = &v
  return s
}

func (s *GetBotRequestTrendsAndTriggerRulesDataResponseDataRequestTrendDetail) SetValue(v int64) *GetBotRequestTrendsAndTriggerRulesDataResponseDataRequestTrendDetail {
  s.Value = &v
  return s
}

type GetBotRequestTrendsAndTriggerRulesDataResponseDataTriggerRule struct     {
  // {"en":"Request time", "zh_CN":"请求时间。"}
  Time *string `json:"time,omitempty" xml:"time,omitempty" require:"true"`
  // {"en":"Trigger rules detail.", "zh_CN":"触发规则详情。"}
  Detail []*GetBotRequestTrendsAndTriggerRulesDataResponseDataTriggerRuleDetail `json:"detail,omitempty" xml:"detail,omitempty" require:"true" type:"Repeated"`
}

func (s GetBotRequestTrendsAndTriggerRulesDataResponseDataTriggerRule) String() string {
  return tea.Prettify(s)
}

func (s GetBotRequestTrendsAndTriggerRulesDataResponseDataTriggerRule) GoString() string {
  return s.String()
}

func (s *GetBotRequestTrendsAndTriggerRulesDataResponseDataTriggerRule) SetTime(v string) *GetBotRequestTrendsAndTriggerRulesDataResponseDataTriggerRule {
  s.Time = &v
  return s
}

func (s *GetBotRequestTrendsAndTriggerRulesDataResponseDataTriggerRule) SetDetail(v []*GetBotRequestTrendsAndTriggerRulesDataResponseDataTriggerRuleDetail) *GetBotRequestTrendsAndTriggerRulesDataResponseDataTriggerRule {
  s.Detail = v
  return s
}

type GetBotRequestTrendsAndTriggerRulesDataResponseDataTriggerRuleDetail struct     {
  // {"en":"Statistical type.", "zh_CN":"统计类型。"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"Request times.", "zh_CN":"请求数。"}
  Value *int64 `json:"value,omitempty" xml:"value,omitempty" require:"true"`
}

func (s GetBotRequestTrendsAndTriggerRulesDataResponseDataTriggerRuleDetail) String() string {
  return tea.Prettify(s)
}

func (s GetBotRequestTrendsAndTriggerRulesDataResponseDataTriggerRuleDetail) GoString() string {
  return s.String()
}

func (s *GetBotRequestTrendsAndTriggerRulesDataResponseDataTriggerRuleDetail) SetName(v string) *GetBotRequestTrendsAndTriggerRulesDataResponseDataTriggerRuleDetail {
  s.Name = &v
  return s
}

func (s *GetBotRequestTrendsAndTriggerRulesDataResponseDataTriggerRuleDetail) SetValue(v int64) *GetBotRequestTrendsAndTriggerRulesDataResponseDataTriggerRuleDetail {
  s.Value = &v
  return s
}

type GetBotRequestTrendsAndTriggerRulesDataPaths struct {
}

func (s GetBotRequestTrendsAndTriggerRulesDataPaths) String() string {
  return tea.Prettify(s)
}

func (s GetBotRequestTrendsAndTriggerRulesDataPaths) GoString() string {
  return s.String()
}

type GetBotRequestTrendsAndTriggerRulesDataParameters struct {
}

func (s GetBotRequestTrendsAndTriggerRulesDataParameters) String() string {
  return tea.Prettify(s)
}

func (s GetBotRequestTrendsAndTriggerRulesDataParameters) GoString() string {
  return s.String()
}

type GetBotRequestTrendsAndTriggerRulesDataRequestHeader struct {
}

func (s GetBotRequestTrendsAndTriggerRulesDataRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetBotRequestTrendsAndTriggerRulesDataRequestHeader) GoString() string {
  return s.String()
}

type GetBotRequestTrendsAndTriggerRulesDataResponseHeader struct {
}

func (s GetBotRequestTrendsAndTriggerRulesDataResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetBotRequestTrendsAndTriggerRulesDataResponseHeader) GoString() string {
  return s.String()
}




type QueryEventTrendRequest struct {
  // {"en":"Start time, format: yyyy-MM-dd HH:mm:ss.\nThe time range can not exceed 31 days.","zh_CN":"开始时间，格式：yyyy-MM-dd HH:mm:ss。\n时间范围不超过31天。"}
  StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty" require:"true"`
  // {"en":"End time, format: yyyy-MM-dd HH:mm:ss.\nThe time range can not exceed 31 days.","zh_CN":"结束时间，格式：yyyy-MM-dd HH:mm:ss。\n时间范围不超过31天。"}
  EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty" require:"true"`
  // {"en":"Multiple selection. Handling results, default: display all results. \nmitigated: Number of mitigated requests. \nmonitored: Number of observed requests.","zh_CN":"多选。处置结果，默认：展示所有结果。\nmitigated：已抵御请求数。\nmonitored：观察请求数。","exampleValue":"mitigated,monitored"}
  ActType []*string `json:"actType,omitempty" xml:"actType,omitempty" type:"Repeated"`
  // {"en":"Domain list. Queries all domains under the account when not specified.","zh_CN":"域名列表，未指定时查询账号下的所有域名。"}
  Domains []*string `json:"domains,omitempty" xml:"domains,omitempty" type:"Repeated"`
}

func (s QueryEventTrendRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryEventTrendRequest) GoString() string {
  return s.String()
}

func (s *QueryEventTrendRequest) SetStartTime(v string) *QueryEventTrendRequest {
  s.StartTime = &v
  return s
}

func (s *QueryEventTrendRequest) SetEndTime(v string) *QueryEventTrendRequest {
  s.EndTime = &v
  return s
}

func (s *QueryEventTrendRequest) SetActType(v []*string) *QueryEventTrendRequest {
  s.ActType = v
  return s
}

func (s *QueryEventTrendRequest) SetDomains(v []*string) *QueryEventTrendRequest {
  s.Domains = v
  return s
}

type QueryEventTrendRequestHeader struct {
  // {"dictionary":"belong=WAAP-MS-Ext|dict=waap_serviceType","en":"Security service type. Please enter a specific service type, if you purchase multiple security services.","zh_CN":"安全服务类型。有使用多个不同的安全服务时，需要填写具体的服务类型。"}
  ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty"`
  // {"defaultValue":"GMT+8","en":"Report Data Timezone:1. Indicates the timezone for the report data. lt must be relative to GMT andspecified in the format GMT+n where -12<=n <= 12. For example,timezone=GMT-7 and timezone=GMT+8.2. If the parameter is unspecified,results will be in the GMT+8 timezone.","zh_CN":"报表数据时区：1.请传递GMT时区。示例格式：GMT+N（其中12<=N <= 12），示例：timezone=GMT-7或timezone=GMT+8.2.若参数为传递，将默认按GMT+8时区查询"}
  Timezone *string `json:"timezone,omitempty" xml:"timezone,omitempty"`
}

func (s QueryEventTrendRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryEventTrendRequestHeader) GoString() string {
  return s.String()
}

func (s *QueryEventTrendRequestHeader) SetServiceType(v string) *QueryEventTrendRequestHeader {
  s.ServiceType = &v
  return s
}

func (s *QueryEventTrendRequestHeader) SetTimezone(v string) *QueryEventTrendRequestHeader {
  s.Timezone = &v
  return s
}

type QueryEventTrendPaths struct {
}

func (s QueryEventTrendPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryEventTrendPaths) GoString() string {
  return s.String()
}

type QueryEventTrendParameters struct {
}

func (s QueryEventTrendParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryEventTrendParameters) GoString() string {
  return s.String()
}

type QueryEventTrendResponse struct {
  // {"dictionary":"belong=WAAP-MS-Ext|dict=waap_retCodeEnum","en":"Please refer to the error code for exceptions.","zh_CN":"请参照错误码。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Description.","zh_CN":"描述信息。"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"Data.","zh_CN":"出参数据。"}
  Data []*QueryEventTrendResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s QueryEventTrendResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryEventTrendResponse) GoString() string {
  return s.String()
}

func (s *QueryEventTrendResponse) SetCode(v string) *QueryEventTrendResponse {
  s.Code = &v
  return s
}

func (s *QueryEventTrendResponse) SetMsg(v string) *QueryEventTrendResponse {
  s.Msg = &v
  return s
}

func (s *QueryEventTrendResponse) SetData(v []*QueryEventTrendResponseData) *QueryEventTrendResponse {
  s.Data = v
  return s
}

type QueryEventTrendResponseData struct     {
  // {"en":"Time, format: yyyy-MM-dd HH-mm-ss.","zh_CN":"时间点，格式：yyyy-MM-dd HH-mm-ss。"}
  TimePoint *string `json:"timePoint,omitempty" xml:"timePoint,omitempty" require:"true"`
  // {"en":"Total requests.","zh_CN":"总请求数。"}
  Total *int64 `json:"total,omitempty" xml:"total,omitempty" require:"true"`
  // {"en":"Attack requests.","zh_CN":"攻击请求数。"}
  Attack *int64 `json:"attack,omitempty" xml:"attack,omitempty" require:"true"`
  // {"en":"Mitigated requests.","zh_CN":"已抵御请求数。"}
  Mitigated *int64 `json:"mitigated,omitempty" xml:"mitigated,omitempty" require:"true"`
  // {"en":"Monitored requests.","zh_CN":"观察请求数。"}
  Monitored *int64 `json:"monitored,omitempty" xml:"monitored,omitempty" require:"true"`
  // {"en":"Whitelist requests.","zh_CN":"白名单请求数。"}
  Whitelist *int64 `json:"whitelist,omitempty" xml:"whitelist,omitempty" require:"true"`
  // {"en":"Policy type classification requests.","zh_CN":"策略类型分类请求数。"}
  Distribution []*QueryEventTrendResponseDataDistribution `json:"distribution,omitempty" xml:"distribution,omitempty" require:"true" type:"Repeated"`
}

func (s QueryEventTrendResponseData) String() string {
  return tea.Prettify(s)
}

func (s QueryEventTrendResponseData) GoString() string {
  return s.String()
}

func (s *QueryEventTrendResponseData) SetTimePoint(v string) *QueryEventTrendResponseData {
  s.TimePoint = &v
  return s
}

func (s *QueryEventTrendResponseData) SetTotal(v int64) *QueryEventTrendResponseData {
  s.Total = &v
  return s
}

func (s *QueryEventTrendResponseData) SetAttack(v int64) *QueryEventTrendResponseData {
  s.Attack = &v
  return s
}

func (s *QueryEventTrendResponseData) SetMitigated(v int64) *QueryEventTrendResponseData {
  s.Mitigated = &v
  return s
}

func (s *QueryEventTrendResponseData) SetMonitored(v int64) *QueryEventTrendResponseData {
  s.Monitored = &v
  return s
}

func (s *QueryEventTrendResponseData) SetWhitelist(v int64) *QueryEventTrendResponseData {
  s.Whitelist = &v
  return s
}

func (s *QueryEventTrendResponseData) SetDistribution(v []*QueryEventTrendResponseDataDistribution) *QueryEventTrendResponseData {
  s.Distribution = v
  return s
}

type QueryEventTrendResponseDataDistribution struct     {
  // {"en":"Policy type.\nBLOCK: IP/Geo Block\nDMS_DEFEND: DDoS Protection\nWAF_DEFEND: WAF\nBOT_MANAGE: Bot Management\nAPI_DEFEND: API Security\nINTELLIGENCE: Threat Intelligence\nRATE_LIMIT: Rate Limiting\nCUSTOMIZE_RULE: Custom Rules","zh_CN":"策略类型。\nBLOCK：IP/区域封禁\nDMS_DEFEND：DDoS防护\nWAF_DEFEND：WAF\nBOT_MANAGE：Bot管理\nAPI_DEFEND：API安全\nINTELLIGENCE：威胁情报\nRATE_LIMIT：频率限制\nCUSTOMIZE_RULE：自定义规则","exampleValue":"BLOCK,DMS_DEFEND,WAF_DEFEND,BOT_MANAGE,API_DEFEND,INTELLIGENCE,RATE_LIMIT,CUSTOMIZE_RULE"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Number of policy requests of this type.","zh_CN":"该策略类型请求数。"}
  Value *int64 `json:"value,omitempty" xml:"value,omitempty" require:"true"`
}

func (s QueryEventTrendResponseDataDistribution) String() string {
  return tea.Prettify(s)
}

func (s QueryEventTrendResponseDataDistribution) GoString() string {
  return s.String()
}

func (s *QueryEventTrendResponseDataDistribution) SetCode(v string) *QueryEventTrendResponseDataDistribution {
  s.Code = &v
  return s
}

func (s *QueryEventTrendResponseDataDistribution) SetValue(v int64) *QueryEventTrendResponseDataDistribution {
  s.Value = &v
  return s
}

type QueryEventTrendResponseHeader struct {
}

func (s QueryEventTrendResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryEventTrendResponseHeader) GoString() string {
  return s.String()
}




type ListServiceUsageRequest struct {
  // {"en":"Code of Control Group.","zh_CN":"Control Group的code。"}
  ControlGroupCode *string `json:"controlGroupCode,omitempty" xml:"controlGroupCode,omitempty" require:"true"`
  // {"en":"End date, format: yyyy-MM-dd.The time range does not exceed 31 days, only supports querying data within one year..","zh_CN":"结束日期，格式：yyyy-MM-dd。时间范围不超过31天，仅支持查询一年内的数据。"}
  EndDate *string `json:"endDate,omitempty" xml:"endDate,omitempty" require:"true"`
  // {"en":"Type of Control Group.ALL: All  PART: Custom category  DOMAIN_ITEM: Contract category","zh_CN":"Control Group的类型。ALL： 所有  PART：自定义类 DOMAIN_ITEM： 合同类"}
  ControlGroupType *string `json:"controlGroupType,omitempty" xml:"controlGroupType,omitempty" require:"true"`
  // {"en":"Start date, format: yyyy-MM-dd.The time range does not exceed 31 days, only supports querying data within one year..","zh_CN":"开始日期，格式：yyyy-MM-dd。时间范围不超过31天，仅支持查询一年内的数据。"}
  StartDate *string `json:"startDate,omitempty" xml:"startDate,omitempty" require:"true"`
}

func (s ListServiceUsageRequest) String() string {
  return tea.Prettify(s)
}

func (s ListServiceUsageRequest) GoString() string {
  return s.String()
}

func (s *ListServiceUsageRequest) SetControlGroupCode(v string) *ListServiceUsageRequest {
  s.ControlGroupCode = &v
  return s
}

func (s *ListServiceUsageRequest) SetEndDate(v string) *ListServiceUsageRequest {
  s.EndDate = &v
  return s
}

func (s *ListServiceUsageRequest) SetControlGroupType(v string) *ListServiceUsageRequest {
  s.ControlGroupType = &v
  return s
}

func (s *ListServiceUsageRequest) SetStartDate(v string) *ListServiceUsageRequest {
  s.StartDate = &v
  return s
}

type ListServiceUsageRequestHeader struct {
  // {"en":"The language of response data, default value: en.zh_CN: Chinese en: English","zh_CN":"返回内容的语言版本，默认值: en。zh_CN：中文 en：英文"}
  Language *string `json:"language,omitempty" xml:"language,omitempty"`
  // {"dictionary":"belong=WAAP-MS-Ext|dict=waap_serviceType","en":"Security service type. Please enter a specific service type, if you purchase multiple security services.","zh_CN":"安全服务类型。有使用多个不同的安全服务时，需要填写具体的服务类型。"}
  ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty"`
  // {"defaultValue":"GMT+8","en":"Report Data Timezone:1. Indicates the timezone for the report data. lt must be relative to GMT andspecified in the format GMT+n where -12<=n <= 12. For example,timezone=GMT-7 and timezone=GMT+8.2. If the parameter is unspecified,results will be in the GMT+8 timezone.","zh_CN":"报表数据时区：1.请传递GMT时区。示例格式：GMT+N（其中12<=N <= 12），示例：timezone=GMT-7或timezone=GMT+8.2.若参数为传递，将默认按GMT+8时区查询"}
  Timezone *string `json:"timezone,omitempty" xml:"timezone,omitempty"`
}

func (s ListServiceUsageRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ListServiceUsageRequestHeader) GoString() string {
  return s.String()
}

func (s *ListServiceUsageRequestHeader) SetLanguage(v string) *ListServiceUsageRequestHeader {
  s.Language = &v
  return s
}

func (s *ListServiceUsageRequestHeader) SetServiceType(v string) *ListServiceUsageRequestHeader {
  s.ServiceType = &v
  return s
}

func (s *ListServiceUsageRequestHeader) SetTimezone(v string) *ListServiceUsageRequestHeader {
  s.Timezone = &v
  return s
}

type ListServiceUsagePaths struct {
}

func (s ListServiceUsagePaths) String() string {
  return tea.Prettify(s)
}

func (s ListServiceUsagePaths) GoString() string {
  return s.String()
}

type ListServiceUsageParameters struct {
}

func (s ListServiceUsageParameters) String() string {
  return tea.Prettify(s)
}

func (s ListServiceUsageParameters) GoString() string {
  return s.String()
}

type ListServiceUsageResponse struct {
  // {"en":"Description.","zh_CN":"描述信息。"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"dictionary":"belong=WAAP-MS-Ext|dict=waap_retCodeEnum","en":"Please refer to the error code for exceptions.","zh_CN":"请参照错误码。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Data.","zh_CN":"出参数据。"}
  Data []*ListServiceUsageResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s ListServiceUsageResponse) String() string {
  return tea.Prettify(s)
}

func (s ListServiceUsageResponse) GoString() string {
  return s.String()
}

func (s *ListServiceUsageResponse) SetMsg(v string) *ListServiceUsageResponse {
  s.Msg = &v
  return s
}

func (s *ListServiceUsageResponse) SetCode(v string) *ListServiceUsageResponse {
  s.Code = &v
  return s
}

func (s *ListServiceUsageResponse) SetData(v []*ListServiceUsageResponseData) *ListServiceUsageResponse {
  s.Data = v
  return s
}

type ListServiceUsageResponseData struct     {
  // {"en":"Primary Domain.","zh_CN":"主域名。"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {"en":"Subdomain list.","zh_CN":"子域名列表。"}
  SubDomainLogList []*ListServiceUsageResponseDataSubDomainLogList `json:"subDomainLogList,omitempty" xml:"subDomainLogList,omitempty" require:"true" type:"Repeated"`
}

func (s ListServiceUsageResponseData) String() string {
  return tea.Prettify(s)
}

func (s ListServiceUsageResponseData) GoString() string {
  return s.String()
}

func (s *ListServiceUsageResponseData) SetDomain(v string) *ListServiceUsageResponseData {
  s.Domain = &v
  return s
}

func (s *ListServiceUsageResponseData) SetSubDomainLogList(v []*ListServiceUsageResponseDataSubDomainLogList) *ListServiceUsageResponseData {
  s.SubDomainLogList = v
  return s
}

type ListServiceUsageResponseDataSubDomainLogList struct     {
  // {"en":"Subdomain.","zh_CN":"子域名。"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {"en":"Protection start time.","zh_CN":"防护接入时间。"}
  StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty" require:"true"`
  // {"en":"Protection end time.","zh_CN":"防护结束时间。"}
  EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty" require:"true"`
}

func (s ListServiceUsageResponseDataSubDomainLogList) String() string {
  return tea.Prettify(s)
}

func (s ListServiceUsageResponseDataSubDomainLogList) GoString() string {
  return s.String()
}

func (s *ListServiceUsageResponseDataSubDomainLogList) SetDomain(v string) *ListServiceUsageResponseDataSubDomainLogList {
  s.Domain = &v
  return s
}

func (s *ListServiceUsageResponseDataSubDomainLogList) SetStartTime(v string) *ListServiceUsageResponseDataSubDomainLogList {
  s.StartTime = &v
  return s
}

func (s *ListServiceUsageResponseDataSubDomainLogList) SetEndTime(v string) *ListServiceUsageResponseDataSubDomainLogList {
  s.EndTime = &v
  return s
}

type ListServiceUsageResponseHeader struct {
}

func (s ListServiceUsageResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ListServiceUsageResponseHeader) GoString() string {
  return s.String()
}




type GetInfrastructureLogListRequest struct {
  // {"en":"Start time, format:yyyy-MM-dd HH:mm:ss.\nOnly supports querying logs of the past month, the query time range cannot exceed 24 hours.","zh_CN":"开始时间，格式：yyyy-MM-dd HH:mm:ss。\n仅支持查询近一个月的日志，查询时间范围不能超过24小时。"}
  StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty" require:"true"`
  // {"en":"End time, format:yyyy-MM-dd HH:mm:ss.\nOnly supports querying logs of the past month, the query time range cannot exceed 24 hours.","zh_CN":"结束时间，格式：yyyy-MM-dd HH:mm:ss。\n仅支持查询近一个月的日志，查询时间范围不能超过24小时，格式：yyyy-MM-dd HH:mm:ss。"}
  EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty" require:"true"`
  // {"en":"Client IP.","zh_CN":"客户端IP。"}
  ClientIps []*string `json:"clientIps,omitempty" xml:"clientIps,omitempty" type:"Repeated"`
  // {"en":"Target IP.","zh_CN":"目标IP。"}
  TargetIps []*string `json:"targetIps,omitempty" xml:"targetIps,omitempty" type:"Repeated"`
  // {"en":"Rule Name.","zh_CN":"规则名称。"}
  RuleName *string `json:"ruleName,omitempty" xml:"ruleName,omitempty"`
  // {"en":"Policy Name.\nrule_protection:Managed Ruleset Protection","zh_CN":"策略名称。\nrule_protection：内置防护规则"}
  PolicyNames []*string `json:"policyNames,omitempty" xml:"policyNames,omitempty" require:"true" type:"Repeated"`
  // {"en":"Action.\nconnection_denied:Connection Refused","zh_CN":"处理动作。\nconnection_denied：拒绝连接"}
  Actions []*string `json:"actions,omitempty" xml:"actions,omitempty" require:"true" type:"Repeated"`
  // {"defaultValue":"10","en":"The number of records per page, default value:10","zh_CN":"每页显示的条目数。默认值：10"}
  PageSize *int `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
  // {"defaultValue":"1","en":"Current page, default:1","zh_CN":"当前第几页。默认值：1"}
  CurrentPage *int `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
}

func (s GetInfrastructureLogListRequest) String() string {
  return tea.Prettify(s)
}

func (s GetInfrastructureLogListRequest) GoString() string {
  return s.String()
}

func (s *GetInfrastructureLogListRequest) SetStartTime(v string) *GetInfrastructureLogListRequest {
  s.StartTime = &v
  return s
}

func (s *GetInfrastructureLogListRequest) SetEndTime(v string) *GetInfrastructureLogListRequest {
  s.EndTime = &v
  return s
}

func (s *GetInfrastructureLogListRequest) SetClientIps(v []*string) *GetInfrastructureLogListRequest {
  s.ClientIps = v
  return s
}

func (s *GetInfrastructureLogListRequest) SetTargetIps(v []*string) *GetInfrastructureLogListRequest {
  s.TargetIps = v
  return s
}

func (s *GetInfrastructureLogListRequest) SetRuleName(v string) *GetInfrastructureLogListRequest {
  s.RuleName = &v
  return s
}

func (s *GetInfrastructureLogListRequest) SetPolicyNames(v []*string) *GetInfrastructureLogListRequest {
  s.PolicyNames = v
  return s
}

func (s *GetInfrastructureLogListRequest) SetActions(v []*string) *GetInfrastructureLogListRequest {
  s.Actions = v
  return s
}

func (s *GetInfrastructureLogListRequest) SetPageSize(v int) *GetInfrastructureLogListRequest {
  s.PageSize = &v
  return s
}

func (s *GetInfrastructureLogListRequest) SetCurrentPage(v int) *GetInfrastructureLogListRequest {
  s.CurrentPage = &v
  return s
}

type GetInfrastructureLogListRequestHeader struct {
  // {"en":"Security service type. Please enter a specific service type, if you purchase multiple security services.","zh_CN":"安全服务类型。有使用多个不同的安全服务时，需要填写具体的服务类型。"}
  ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty"`
  // {"defaultValue":"GMT+8","en":"Report Data Timezone:1. Indicates the timezone for the report data. lt must be relative to GMT andspecified in the format GMT+n where -12<=n <= 12. For example,timezone=GMT-7 and timezone=GMT+8.2. If the parameter is unspecified,results will be in the GMT+8 timezone.","zh_CN":"报表数据时区：1.请传递GMT时区。示例格式：GMT+N（其中12<=N <= 12），示例：timezone=GMT-7或timezone=GMT+8.2.若参数为传递，将默认按GMT+8时区查询"}
  Timezone *string `json:"timezone,omitempty" xml:"timezone,omitempty"`
}

func (s GetInfrastructureLogListRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetInfrastructureLogListRequestHeader) GoString() string {
  return s.String()
}

func (s *GetInfrastructureLogListRequestHeader) SetServiceType(v string) *GetInfrastructureLogListRequestHeader {
  s.ServiceType = &v
  return s
}

func (s *GetInfrastructureLogListRequestHeader) SetTimezone(v string) *GetInfrastructureLogListRequestHeader {
  s.Timezone = &v
  return s
}

type GetInfrastructureLogListPaths struct {
}

func (s GetInfrastructureLogListPaths) String() string {
  return tea.Prettify(s)
}

func (s GetInfrastructureLogListPaths) GoString() string {
  return s.String()
}

type GetInfrastructureLogListParameters struct {
}

func (s GetInfrastructureLogListParameters) String() string {
  return tea.Prettify(s)
}

func (s GetInfrastructureLogListParameters) GoString() string {
  return s.String()
}

type GetInfrastructureLogListResponse struct {
  // {"dictionary":"belong=WAAP-MS-Ext|dict=waap_retCodeEnum","en":"Please refer to the error code for exceptions.","zh_CN":"请参照错误码。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Description.","zh_CN":"描述信息。"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"Data.","zh_CN":"出参数据。"}
  Data *GetInfrastructureLogListResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
}

func (s GetInfrastructureLogListResponse) String() string {
  return tea.Prettify(s)
}

func (s GetInfrastructureLogListResponse) GoString() string {
  return s.String()
}

func (s *GetInfrastructureLogListResponse) SetCode(v string) *GetInfrastructureLogListResponse {
  s.Code = &v
  return s
}

func (s *GetInfrastructureLogListResponse) SetMsg(v string) *GetInfrastructureLogListResponse {
  s.Msg = &v
  return s
}

func (s *GetInfrastructureLogListResponse) SetData(v *GetInfrastructureLogListResponseData) *GetInfrastructureLogListResponse {
  s.Data = v
  return s
}

type GetInfrastructureLogListResponseData struct {
  // {"en":"The current page number.","zh_CN":"当前页码。"}
  Current *int64 `json:"current,omitempty" xml:"current,omitempty" require:"true"`
  // {"en":"The number of records per page.","zh_CN":"每页显示的条目数。"}
  Size *int64 `json:"size,omitempty" xml:"size,omitempty" require:"true"`
  // {"en":"Pages.","zh_CN":"页数。"}
  Pages *int64 `json:"pages,omitempty" xml:"pages,omitempty" require:"true"`
  // {"en":"The total number of records.","zh_CN":"总记录数。"}
  Total *int64 `json:"total,omitempty" xml:"total,omitempty" require:"true"`
  // {"en":"Records.","zh_CN":"记录。"}
  Records *GetInfrastructureLogListResponseDataRecords `json:"records,omitempty" xml:"records,omitempty" require:"true" type:"Struct"`
}

func (s GetInfrastructureLogListResponseData) String() string {
  return tea.Prettify(s)
}

func (s GetInfrastructureLogListResponseData) GoString() string {
  return s.String()
}

func (s *GetInfrastructureLogListResponseData) SetCurrent(v int64) *GetInfrastructureLogListResponseData {
  s.Current = &v
  return s
}

func (s *GetInfrastructureLogListResponseData) SetSize(v int64) *GetInfrastructureLogListResponseData {
  s.Size = &v
  return s
}

func (s *GetInfrastructureLogListResponseData) SetPages(v int64) *GetInfrastructureLogListResponseData {
  s.Pages = &v
  return s
}

func (s *GetInfrastructureLogListResponseData) SetTotal(v int64) *GetInfrastructureLogListResponseData {
  s.Total = &v
  return s
}

func (s *GetInfrastructureLogListResponseData) SetRecords(v *GetInfrastructureLogListResponseDataRecords) *GetInfrastructureLogListResponseData {
  s.Records = v
  return s
}

type GetInfrastructureLogListResponseDataRecords struct {
  // {"en":"Time.","zh_CN":"时间。"}
  AttackTime *string `json:"attackTime,omitempty" xml:"attackTime,omitempty" require:"true"`
  // {"en":"Client IP.","zh_CN":"客户端IP。"}
  ClientIp *string `json:"clientIp,omitempty" xml:"clientIp,omitempty" require:"true"`
  // {"en":"Target IP.","zh_CN":"目标IP。"}
  TargetIp *string `json:"targetIp,omitempty" xml:"targetIp,omitempty" require:"true"`
  // {"en":"Policy Name.","zh_CN":"策略名称。"}
  PolicyName *string `json:"policyName,omitempty" xml:"policyName,omitempty" require:"true"`
  // {"en":"Rule Name.","zh_CN":"规则名称。"}
  RuleName *string `json:"ruleName,omitempty" xml:"ruleName,omitempty" require:"true"`
  // {"en":"Action.","zh_CN":"处理动作。"}
  Action *string `json:"action,omitempty" xml:"action,omitempty" require:"true"`
  // {"en":"IP site.","zh_CN":"IP地理位置。"}
  IpSite *string `json:"ipSite,omitempty" xml:"ipSite,omitempty" require:"true"`
  // {"en":"Explanation.","zh_CN":"异常说明。"}
  Content *string `json:"content,omitempty" xml:"content,omitempty" require:"true"`
  // {"en":"Target port.","zh_CN":"目标端口。"}
  TargetPort *string `json:"targetPort,omitempty" xml:"targetPort,omitempty" require:"true"`
}

func (s GetInfrastructureLogListResponseDataRecords) String() string {
  return tea.Prettify(s)
}

func (s GetInfrastructureLogListResponseDataRecords) GoString() string {
  return s.String()
}

func (s *GetInfrastructureLogListResponseDataRecords) SetAttackTime(v string) *GetInfrastructureLogListResponseDataRecords {
  s.AttackTime = &v
  return s
}

func (s *GetInfrastructureLogListResponseDataRecords) SetClientIp(v string) *GetInfrastructureLogListResponseDataRecords {
  s.ClientIp = &v
  return s
}

func (s *GetInfrastructureLogListResponseDataRecords) SetTargetIp(v string) *GetInfrastructureLogListResponseDataRecords {
  s.TargetIp = &v
  return s
}

func (s *GetInfrastructureLogListResponseDataRecords) SetPolicyName(v string) *GetInfrastructureLogListResponseDataRecords {
  s.PolicyName = &v
  return s
}

func (s *GetInfrastructureLogListResponseDataRecords) SetRuleName(v string) *GetInfrastructureLogListResponseDataRecords {
  s.RuleName = &v
  return s
}

func (s *GetInfrastructureLogListResponseDataRecords) SetAction(v string) *GetInfrastructureLogListResponseDataRecords {
  s.Action = &v
  return s
}

func (s *GetInfrastructureLogListResponseDataRecords) SetIpSite(v string) *GetInfrastructureLogListResponseDataRecords {
  s.IpSite = &v
  return s
}

func (s *GetInfrastructureLogListResponseDataRecords) SetContent(v string) *GetInfrastructureLogListResponseDataRecords {
  s.Content = &v
  return s
}

func (s *GetInfrastructureLogListResponseDataRecords) SetTargetPort(v string) *GetInfrastructureLogListResponseDataRecords {
  s.TargetPort = &v
  return s
}

type GetInfrastructureLogListResponseHeader struct {
}

func (s GetInfrastructureLogListResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetInfrastructureLogListResponseHeader) GoString() string {
  return s.String()
}




type GetBotRequestSourceDistributionDataRequest struct {
  // {"en":"Domain.Separate by ';'.", "zh_CN":"域名。多个以;隔开"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {"en":"StartTime. Format: yyyy-MM-dd HH:mm:ss", "zh_CN":"开始时间。格式： yyyy-MM-dd HH:mm:ss"}
  StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty" require:"true"`
  // {"en":"EndTime. Format: yyyy-MM-dd HH:mm:ss", "zh_CN":"结束时间。格式： yyyy-MM-dd HH:mm:ss"}
  EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty" require:"true"`
  // {"en":"Time zone. Default 8, i.e.'GTM+8'", "zh_CN":"时区，默认8，即“GTM+8”"}
  TimeZone *int `json:"timeZone,omitempty" xml:"timeZone,omitempty"`
  // {"en":"Language type. Default cn. 
  //  en:English 
  //  cn:Chinese ", "zh_CN":"语言类型。 默认cn 
  //  en：英文 
  //  cn：中文"}
  Lang *string `json:"lang,omitempty" xml:"lang,omitempty"`
  // {"en":"Number of top values.Default value 10", "zh_CN":"排名最前值数目。默认10"}
  TopNum *int `json:"topNum,omitempty" xml:"topNum,omitempty"`
  // {"en":"Area search type. 
  //  0:Global 
  //  1:China ", "zh_CN":"地区查询类型。 
  //  0：全球 
  //  1：中国"}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
}

func (s GetBotRequestSourceDistributionDataRequest) String() string {
  return tea.Prettify(s)
}

func (s GetBotRequestSourceDistributionDataRequest) GoString() string {
  return s.String()
}

func (s *GetBotRequestSourceDistributionDataRequest) SetDomain(v string) *GetBotRequestSourceDistributionDataRequest {
  s.Domain = &v
  return s
}

func (s *GetBotRequestSourceDistributionDataRequest) SetStartTime(v string) *GetBotRequestSourceDistributionDataRequest {
  s.StartTime = &v
  return s
}

func (s *GetBotRequestSourceDistributionDataRequest) SetEndTime(v string) *GetBotRequestSourceDistributionDataRequest {
  s.EndTime = &v
  return s
}

func (s *GetBotRequestSourceDistributionDataRequest) SetTimeZone(v int) *GetBotRequestSourceDistributionDataRequest {
  s.TimeZone = &v
  return s
}

func (s *GetBotRequestSourceDistributionDataRequest) SetLang(v string) *GetBotRequestSourceDistributionDataRequest {
  s.Lang = &v
  return s
}

func (s *GetBotRequestSourceDistributionDataRequest) SetTopNum(v int) *GetBotRequestSourceDistributionDataRequest {
  s.TopNum = &v
  return s
}

func (s *GetBotRequestSourceDistributionDataRequest) SetType(v string) *GetBotRequestSourceDistributionDataRequest {
  s.Type = &v
  return s
}

type GetBotRequestSourceDistributionDataResponse struct {
  // {"en":"Status code, success is '200'.", "zh_CN":"状态码，成功为“200”。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Return message, success is 'Success'.", "zh_CN":"返回信息，成功为“Success”。"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Data returned.", "zh_CN":"返回数据。"}
  Data []*GetBotRequestSourceDistributionDataResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s GetBotRequestSourceDistributionDataResponse) String() string {
  return tea.Prettify(s)
}

func (s GetBotRequestSourceDistributionDataResponse) GoString() string {
  return s.String()
}

func (s *GetBotRequestSourceDistributionDataResponse) SetCode(v string) *GetBotRequestSourceDistributionDataResponse {
  s.Code = &v
  return s
}

func (s *GetBotRequestSourceDistributionDataResponse) SetMessage(v string) *GetBotRequestSourceDistributionDataResponse {
  s.Message = &v
  return s
}

func (s *GetBotRequestSourceDistributionDataResponse) SetData(v []*GetBotRequestSourceDistributionDataResponseData) *GetBotRequestSourceDistributionDataResponse {
  s.Data = v
  return s
}

type GetBotRequestSourceDistributionDataResponseData struct     {
  // {"en":"Statistical type.", "zh_CN":"统计类型。"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"Bot request count", "zh_CN":"Bot请求数。"}
  Count *int64 `json:"count,omitempty" xml:"count,omitempty" require:"true"`
}

func (s GetBotRequestSourceDistributionDataResponseData) String() string {
  return tea.Prettify(s)
}

func (s GetBotRequestSourceDistributionDataResponseData) GoString() string {
  return s.String()
}

func (s *GetBotRequestSourceDistributionDataResponseData) SetName(v string) *GetBotRequestSourceDistributionDataResponseData {
  s.Name = &v
  return s
}

func (s *GetBotRequestSourceDistributionDataResponseData) SetCount(v int64) *GetBotRequestSourceDistributionDataResponseData {
  s.Count = &v
  return s
}

type GetBotRequestSourceDistributionDataPaths struct {
}

func (s GetBotRequestSourceDistributionDataPaths) String() string {
  return tea.Prettify(s)
}

func (s GetBotRequestSourceDistributionDataPaths) GoString() string {
  return s.String()
}

type GetBotRequestSourceDistributionDataParameters struct {
}

func (s GetBotRequestSourceDistributionDataParameters) String() string {
  return tea.Prettify(s)
}

func (s GetBotRequestSourceDistributionDataParameters) GoString() string {
  return s.String()
}

type GetBotRequestSourceDistributionDataRequestHeader struct {
}

func (s GetBotRequestSourceDistributionDataRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetBotRequestSourceDistributionDataRequestHeader) GoString() string {
  return s.String()
}

type GetBotRequestSourceDistributionDataResponseHeader struct {
}

func (s GetBotRequestSourceDistributionDataResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetBotRequestSourceDistributionDataResponseHeader) GoString() string {
  return s.String()
}




type QueryAttackLogDetailInfoRequest struct {
  // {"en":"Event ID.","zh_CN":"事件ID。"}
  Uuid *string `json:"uuid,omitempty" xml:"uuid,omitempty" require:"true"`
  // {"en":"Request ID.","zh_CN":"请求ID。"}
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty" require:"true"`
  // {"en":"Domain.","zh_CN":"域名。"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {"en":"Policy type.\nDMS_DEFEND: DDoS Protection\nBOT_MANAGE: Bot Management\nAPI_DEFEND: API Security\nWAF_DEFEND: WAF\nBLOCK: IP/Geo Block\nCUSTOMIZE_RULE: Custom Rules\nRATE_LIMIT: Rate Limiting\nINTELLIGENCE: Threat Intelligence","zh_CN":"策略类型。\nDMS_DEFEND：DDoS防护\nBOT_MANAGE：Bot管理\nAPI_DEFEND：API安全\nWAF_DEFEND：WAF\nBLOCK：IP/区域封禁\nCUSTOMIZE_RULE：自定义规则\nRATE_LIMIT：频率限制\nINTELLIGENCE：威胁情报","exampleValue":"BLOCK,DMS_DEFEND,WAF_DEFEND,BOT_MANAGE,API_DEFEND,INTELLIGENCE,RATE_LIMIT,CUSTOMIZE_RULE"}
  PolicyType *string `json:"policyType,omitempty" xml:"policyType,omitempty" require:"true"`
  // {"en":"Attack time, format: yyyy-MM-dd HH:mm:ss.","zh_CN":"攻击时间，时间格式：yyyy-MM-dd HH:mm:ss。"}
  AttackTime *string `json:"attackTime,omitempty" xml:"attackTime,omitempty" require:"true"`
}

func (s QueryAttackLogDetailInfoRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryAttackLogDetailInfoRequest) GoString() string {
  return s.String()
}

func (s *QueryAttackLogDetailInfoRequest) SetUuid(v string) *QueryAttackLogDetailInfoRequest {
  s.Uuid = &v
  return s
}

func (s *QueryAttackLogDetailInfoRequest) SetRequestId(v string) *QueryAttackLogDetailInfoRequest {
  s.RequestId = &v
  return s
}

func (s *QueryAttackLogDetailInfoRequest) SetDomain(v string) *QueryAttackLogDetailInfoRequest {
  s.Domain = &v
  return s
}

func (s *QueryAttackLogDetailInfoRequest) SetPolicyType(v string) *QueryAttackLogDetailInfoRequest {
  s.PolicyType = &v
  return s
}

func (s *QueryAttackLogDetailInfoRequest) SetAttackTime(v string) *QueryAttackLogDetailInfoRequest {
  s.AttackTime = &v
  return s
}

type QueryAttackLogDetailInfoRequestHeader struct {
  // {"defaultValue":"en","en":"The language of response data, default value: en.\nzh_CN: Chinese\nen: English","zh_CN":"返回内容的语言版本，默认值: en。\nzh_CN：中文\nen：英文","exampleValue":"zh_CN,en"}
  Language *string `json:"language,omitempty" xml:"language,omitempty"`
  // {"dictionary":"belong=WAAP-MS-Ext|dict=waap_serviceType","en":"Security service type. Please enter a specific service type, if you purchase multiple security services.","zh_CN":"安全服务类型。有使用多个不同的安全服务时，需要填写具体的服务类型。"}
  ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty"`
  // {"defaultValue":"GMT+8","en":"Report Data Timezone:1. Indicates the timezone for the report data. lt must be relative to GMT andspecified in the format GMT+n where -12<=n <= 12. For example,timezone=GMT-7 and timezone=GMT+8.2. If the parameter is unspecified,results will be in the GMT+8 timezone.","zh_CN":"报表数据时区：1.请传递GMT时区。示例格式：GMT+N（其中12<=N <= 12），示例：timezone=GMT-7或timezone=GMT+8.2.若参数为传递，将默认按GMT+8时区查询"}
  Timezone *string `json:"timezone,omitempty" xml:"timezone,omitempty"`
}

func (s QueryAttackLogDetailInfoRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryAttackLogDetailInfoRequestHeader) GoString() string {
  return s.String()
}

func (s *QueryAttackLogDetailInfoRequestHeader) SetLanguage(v string) *QueryAttackLogDetailInfoRequestHeader {
  s.Language = &v
  return s
}

func (s *QueryAttackLogDetailInfoRequestHeader) SetServiceType(v string) *QueryAttackLogDetailInfoRequestHeader {
  s.ServiceType = &v
  return s
}

func (s *QueryAttackLogDetailInfoRequestHeader) SetTimezone(v string) *QueryAttackLogDetailInfoRequestHeader {
  s.Timezone = &v
  return s
}

type QueryAttackLogDetailInfoPaths struct {
}

func (s QueryAttackLogDetailInfoPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryAttackLogDetailInfoPaths) GoString() string {
  return s.String()
}

type QueryAttackLogDetailInfoParameters struct {
}

func (s QueryAttackLogDetailInfoParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryAttackLogDetailInfoParameters) GoString() string {
  return s.String()
}

type QueryAttackLogDetailInfoResponse struct {
  // {"dictionary":"belong=WAAP-MS-Ext|dict=waap_retCodeEnum","en":"Please refer to the error code for exceptions.","zh_CN":"请参照错误码。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Description.","zh_CN":"描述信息。"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"Data.","zh_CN":"出参数据。"}
  Data *QueryAttackLogDetailInfoResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
}

func (s QueryAttackLogDetailInfoResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryAttackLogDetailInfoResponse) GoString() string {
  return s.String()
}

func (s *QueryAttackLogDetailInfoResponse) SetCode(v string) *QueryAttackLogDetailInfoResponse {
  s.Code = &v
  return s
}

func (s *QueryAttackLogDetailInfoResponse) SetMsg(v string) *QueryAttackLogDetailInfoResponse {
  s.Msg = &v
  return s
}

func (s *QueryAttackLogDetailInfoResponse) SetData(v *QueryAttackLogDetailInfoResponseData) *QueryAttackLogDetailInfoResponse {
  s.Data = v
  return s
}

type QueryAttackLogDetailInfoResponseData struct {
  // {"en":"Basic Information.","zh_CN":"基础信息。"}
  BasicInfo *QueryAttackLogDetailInfoResponseDataBasicInfo `json:"basicInfo,omitempty" xml:"basicInfo,omitempty" require:"true" type:"Struct"`
  // {"en":"Details.","zh_CN":"详细信息。"}
  DetailInfo *QueryAttackLogDetailInfoResponseDataDetailInfo `json:"detailInfo,omitempty" xml:"detailInfo,omitempty" require:"true" type:"Struct"`
}

func (s QueryAttackLogDetailInfoResponseData) String() string {
  return tea.Prettify(s)
}

func (s QueryAttackLogDetailInfoResponseData) GoString() string {
  return s.String()
}

func (s *QueryAttackLogDetailInfoResponseData) SetBasicInfo(v *QueryAttackLogDetailInfoResponseDataBasicInfo) *QueryAttackLogDetailInfoResponseData {
  s.BasicInfo = v
  return s
}

func (s *QueryAttackLogDetailInfoResponseData) SetDetailInfo(v *QueryAttackLogDetailInfoResponseDataDetailInfo) *QueryAttackLogDetailInfoResponseData {
  s.DetailInfo = v
  return s
}

type QueryAttackLogDetailInfoResponseDataBasicInfo struct {
  // {"en":"Request ID.","zh_CN":"请求ID。"}
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty" require:"true"`
  // {"en":"Event ID.","zh_CN":"事件ID。"}
  Uuid *string `json:"uuid,omitempty" xml:"uuid,omitempty" require:"true"`
  // {"en":"Request Method.","zh_CN":"请求方法。"}
  RequestMethod *string `json:"requestMethod,omitempty" xml:"requestMethod,omitempty" require:"true"`
  // {"en":"Http Version.","zh_CN":"HTTP版本。"}
  HttpVersion *string `json:"httpVersion,omitempty" xml:"httpVersion,omitempty" require:"true"`
  // {"en":"Path.","zh_CN":"路径。"}
  Path *string `json:"path,omitempty" xml:"path,omitempty" require:"true"`
  // {"en":"URI.","zh_CN":"URI。"}
  Uri *string `json:"uri,omitempty" xml:"uri,omitempty" require:"true"`
  // {"en":"User-Agent.","zh_CN":"User-Agent。"}
  UserAgent *string `json:"userAgent,omitempty" xml:"userAgent,omitempty" require:"true"`
  // {"en":"Referer.","zh_CN":"Referer。"}
  Referer *string `json:"referer,omitempty" xml:"referer,omitempty" require:"true"`
  // {"en":"Status Code.","zh_CN":"状态码。"}
  StatusCode *string `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
  // {"en":"Client IP.","zh_CN":"客户端IP。"}
  ClientIp *string `json:"clientIp,omitempty" xml:"clientIp,omitempty" require:"true"`
  // {"en":"IP geolocation.","zh_CN":"IP地理位置。"}
  IpLocation *string `json:"ipLocation,omitempty" xml:"ipLocation,omitempty" require:"true"`
  // {"en":"Client ID.","zh_CN":"客户端ID。"}
  UserFinger *string `json:"userFinger,omitempty" xml:"userFinger,omitempty" require:"true"`
  // {"en":"Browser fingerprint.","zh_CN":"浏览器指纹。"}
  BrowserFinger *string `json:"browserFinger,omitempty" xml:"browserFinger,omitempty" require:"true"`
  // {"en":"Device fingerprint.","zh_CN":"设备指纹。"}
  DeviceFinger *string `json:"deviceFinger,omitempty" xml:"deviceFinger,omitempty" require:"true"`
}

func (s QueryAttackLogDetailInfoResponseDataBasicInfo) String() string {
  return tea.Prettify(s)
}

func (s QueryAttackLogDetailInfoResponseDataBasicInfo) GoString() string {
  return s.String()
}

func (s *QueryAttackLogDetailInfoResponseDataBasicInfo) SetRequestId(v string) *QueryAttackLogDetailInfoResponseDataBasicInfo {
  s.RequestId = &v
  return s
}

func (s *QueryAttackLogDetailInfoResponseDataBasicInfo) SetUuid(v string) *QueryAttackLogDetailInfoResponseDataBasicInfo {
  s.Uuid = &v
  return s
}

func (s *QueryAttackLogDetailInfoResponseDataBasicInfo) SetRequestMethod(v string) *QueryAttackLogDetailInfoResponseDataBasicInfo {
  s.RequestMethod = &v
  return s
}

func (s *QueryAttackLogDetailInfoResponseDataBasicInfo) SetHttpVersion(v string) *QueryAttackLogDetailInfoResponseDataBasicInfo {
  s.HttpVersion = &v
  return s
}

func (s *QueryAttackLogDetailInfoResponseDataBasicInfo) SetPath(v string) *QueryAttackLogDetailInfoResponseDataBasicInfo {
  s.Path = &v
  return s
}

func (s *QueryAttackLogDetailInfoResponseDataBasicInfo) SetUri(v string) *QueryAttackLogDetailInfoResponseDataBasicInfo {
  s.Uri = &v
  return s
}

func (s *QueryAttackLogDetailInfoResponseDataBasicInfo) SetUserAgent(v string) *QueryAttackLogDetailInfoResponseDataBasicInfo {
  s.UserAgent = &v
  return s
}

func (s *QueryAttackLogDetailInfoResponseDataBasicInfo) SetReferer(v string) *QueryAttackLogDetailInfoResponseDataBasicInfo {
  s.Referer = &v
  return s
}

func (s *QueryAttackLogDetailInfoResponseDataBasicInfo) SetStatusCode(v string) *QueryAttackLogDetailInfoResponseDataBasicInfo {
  s.StatusCode = &v
  return s
}

func (s *QueryAttackLogDetailInfoResponseDataBasicInfo) SetClientIp(v string) *QueryAttackLogDetailInfoResponseDataBasicInfo {
  s.ClientIp = &v
  return s
}

func (s *QueryAttackLogDetailInfoResponseDataBasicInfo) SetIpLocation(v string) *QueryAttackLogDetailInfoResponseDataBasicInfo {
  s.IpLocation = &v
  return s
}

func (s *QueryAttackLogDetailInfoResponseDataBasicInfo) SetUserFinger(v string) *QueryAttackLogDetailInfoResponseDataBasicInfo {
  s.UserFinger = &v
  return s
}

func (s *QueryAttackLogDetailInfoResponseDataBasicInfo) SetBrowserFinger(v string) *QueryAttackLogDetailInfoResponseDataBasicInfo {
  s.BrowserFinger = &v
  return s
}

func (s *QueryAttackLogDetailInfoResponseDataBasicInfo) SetDeviceFinger(v string) *QueryAttackLogDetailInfoResponseDataBasicInfo {
  s.DeviceFinger = &v
  return s
}

type QueryAttackLogDetailInfoResponseDataDetailInfo struct {
  // {"en":"The policy type of hit interception, empty when not intercepted.\nDMS_DEFEND: DDoS Protection\nBOT_MANAGE: Bot Management\nAPI_DEFEND: API Security\nWAF_DEFEND: WAF\nBLOCK: IP/Geo Block\nCUSTOMIZE_RULE: Custom Rules\nRATE_LIMIT: Rate Limiting\nINTELLIGENCE: Threat Intelligence","zh_CN":"命中拦截的策略类型，没有被拦截时为空。\nDMS_DEFEND：DDoS防护\nBOT_MANAGE：Bot管理\nAPI_DEFEND：API安全\nWAF_DEFEND：WAF\nBLOCK：IP/区域封禁\nCUSTOMIZE_RULE：自定义规则\nRATE_LIMIT：频率限制\nINTELLIGENCE：威胁情报","exampleValue":"BLOCK,DMS_DEFEND,WAF_DEFEND,BOT_MANAGE,API_DEFEND,INTELLIGENCE,RATE_LIMIT,CUSTOMIZE_RULE"}
  BlockPolicyName *string `json:"blockPolicyName,omitempty" xml:"blockPolicyName,omitempty" require:"true"`
  // {"en":"IP/Geo Block.","zh_CN":"IP区域封禁。"}
  BLOCK *QueryAttackLogDetailInfoResponseDataDetailInfoBLOCK `json:"BLOCK,omitempty" xml:"BLOCK,omitempty" require:"true" type:"Struct"`
  // {"en":"WAF.","zh_CN":"WAF。"}
  WAFDEFEND []*QueryAttackLogDetailInfoResponseDataDetailInfoWAFDEFEND `json:"WAF_DEFEND,omitempty" xml:"WAF_DEFEND,omitempty" require:"true" type:"Repeated"`
  // {"en":"DDoS Protection.","zh_CN":"DDos防护。"}
  DMSDEFEND []*QueryAttackLogDetailInfoResponseDataDetailInfoDMSDEFEND `json:"DMS_DEFEND,omitempty" xml:"DMS_DEFEND,omitempty" require:"true" type:"Repeated"`
  // {"en":"Bot Management.","zh_CN":"Bot管理。"}
  BOTMANAGE []*QueryAttackLogDetailInfoResponseDataDetailInfoBOTMANAGE `json:"BOT_MANAGE,omitempty" xml:"BOT_MANAGE,omitempty" require:"true" type:"Repeated"`
  // {"en":"API Security.","zh_CN":"API安全。"}
  APIDEFEND []*QueryAttackLogDetailInfoResponseDataDetailInfoAPIDEFEND `json:"API_DEFEND,omitempty" xml:"API_DEFEND,omitempty" require:"true" type:"Repeated"`
  // {"en":"Threat Intelligence.","zh_CN":"威胁情报。"}
  INTELLIGENCE []*QueryAttackLogDetailInfoResponseDataDetailInfoINTELLIGENCE `json:"INTELLIGENCE,omitempty" xml:"INTELLIGENCE,omitempty" require:"true" type:"Repeated"`
  // {"en":"Rate Limiting.","zh_CN":"频率限制。"}
  RATELIMIT []*QueryAttackLogDetailInfoResponseDataDetailInfoRATELIMIT `json:"RATE_LIMIT,omitempty" xml:"RATE_LIMIT,omitempty" require:"true" type:"Repeated"`
  // {"en":"Custom Rules.","zh_CN":"自定义规则。"}
  CUSTOMIZERULE []*QueryAttackLogDetailInfoResponseDataDetailInfoCUSTOMIZERULE `json:"CUSTOMIZE_RULE,omitempty" xml:"CUSTOMIZE_RULE,omitempty" require:"true" type:"Repeated"`
}

func (s QueryAttackLogDetailInfoResponseDataDetailInfo) String() string {
  return tea.Prettify(s)
}

func (s QueryAttackLogDetailInfoResponseDataDetailInfo) GoString() string {
  return s.String()
}

func (s *QueryAttackLogDetailInfoResponseDataDetailInfo) SetBlockPolicyName(v string) *QueryAttackLogDetailInfoResponseDataDetailInfo {
  s.BlockPolicyName = &v
  return s
}

func (s *QueryAttackLogDetailInfoResponseDataDetailInfo) SetBLOCK(v *QueryAttackLogDetailInfoResponseDataDetailInfoBLOCK) *QueryAttackLogDetailInfoResponseDataDetailInfo {
  s.BLOCK = v
  return s
}

func (s *QueryAttackLogDetailInfoResponseDataDetailInfo) SetWAFDEFEND(v []*QueryAttackLogDetailInfoResponseDataDetailInfoWAFDEFEND) *QueryAttackLogDetailInfoResponseDataDetailInfo {
  s.WAFDEFEND = v
  return s
}

func (s *QueryAttackLogDetailInfoResponseDataDetailInfo) SetDMSDEFEND(v []*QueryAttackLogDetailInfoResponseDataDetailInfoDMSDEFEND) *QueryAttackLogDetailInfoResponseDataDetailInfo {
  s.DMSDEFEND = v
  return s
}

func (s *QueryAttackLogDetailInfoResponseDataDetailInfo) SetBOTMANAGE(v []*QueryAttackLogDetailInfoResponseDataDetailInfoBOTMANAGE) *QueryAttackLogDetailInfoResponseDataDetailInfo {
  s.BOTMANAGE = v
  return s
}

func (s *QueryAttackLogDetailInfoResponseDataDetailInfo) SetAPIDEFEND(v []*QueryAttackLogDetailInfoResponseDataDetailInfoAPIDEFEND) *QueryAttackLogDetailInfoResponseDataDetailInfo {
  s.APIDEFEND = v
  return s
}

func (s *QueryAttackLogDetailInfoResponseDataDetailInfo) SetINTELLIGENCE(v []*QueryAttackLogDetailInfoResponseDataDetailInfoINTELLIGENCE) *QueryAttackLogDetailInfoResponseDataDetailInfo {
  s.INTELLIGENCE = v
  return s
}

func (s *QueryAttackLogDetailInfoResponseDataDetailInfo) SetRATELIMIT(v []*QueryAttackLogDetailInfoResponseDataDetailInfoRATELIMIT) *QueryAttackLogDetailInfoResponseDataDetailInfo {
  s.RATELIMIT = v
  return s
}

func (s *QueryAttackLogDetailInfoResponseDataDetailInfo) SetCUSTOMIZERULE(v []*QueryAttackLogDetailInfoResponseDataDetailInfoCUSTOMIZERULE) *QueryAttackLogDetailInfoResponseDataDetailInfo {
  s.CUSTOMIZERULE = v
  return s
}

type QueryAttackLogDetailInfoResponseDataDetailInfoBLOCK struct {
  // {"en":"Policy name.","zh_CN":"策略名称。"}
  PolicyName *string `json:"policyName,omitempty" xml:"policyName,omitempty" require:"true"`
  // {"en":"Rule action.","zh_CN":"规则动作。"}
  RuleAction *string `json:"ruleAction,omitempty" xml:"ruleAction,omitempty" require:"true"`
  // {"en":"Description.","zh_CN":"说明。"}
  Explain *string `json:"explain,omitempty" xml:"explain,omitempty" require:"true"`
}

func (s QueryAttackLogDetailInfoResponseDataDetailInfoBLOCK) String() string {
  return tea.Prettify(s)
}

func (s QueryAttackLogDetailInfoResponseDataDetailInfoBLOCK) GoString() string {
  return s.String()
}

func (s *QueryAttackLogDetailInfoResponseDataDetailInfoBLOCK) SetPolicyName(v string) *QueryAttackLogDetailInfoResponseDataDetailInfoBLOCK {
  s.PolicyName = &v
  return s
}

func (s *QueryAttackLogDetailInfoResponseDataDetailInfoBLOCK) SetRuleAction(v string) *QueryAttackLogDetailInfoResponseDataDetailInfoBLOCK {
  s.RuleAction = &v
  return s
}

func (s *QueryAttackLogDetailInfoResponseDataDetailInfoBLOCK) SetExplain(v string) *QueryAttackLogDetailInfoResponseDataDetailInfoBLOCK {
  s.Explain = &v
  return s
}

type QueryAttackLogDetailInfoResponseDataDetailInfoWAFDEFEND struct     {
  // {"en":"Rule type.","zh_CN":"规则类型。"}
  RuleType *string `json:"ruleType,omitempty" xml:"ruleType,omitempty" require:"true"`
  // {"en":"Rule action.","zh_CN":"规则动作。"}
  RuleAction *string `json:"ruleAction,omitempty" xml:"ruleAction,omitempty" require:"true"`
  // {"en":"Rule ID.","zh_CN":"规则ID。"}
  RuleId *string `json:"ruleId,omitempty" xml:"ruleId,omitempty" require:"true"`
  // {"en":"Rule name.","zh_CN":"规则名称。"}
  RuleName *string `json:"ruleName,omitempty" xml:"ruleName,omitempty" require:"true"`
  // {"en":"Rule description.","zh_CN":"规则描述。"}
  RuleDesc *string `json:"ruleDesc,omitempty" xml:"ruleDesc,omitempty" require:"true"`
  // {"en":"Hit request area.","zh_CN":"命中请求区域。"}
  MatchArea *string `json:"matchArea,omitempty" xml:"matchArea,omitempty" require:"true"`
  // {"en":"Hit request content.","zh_CN":"命中请求内容。"}
  MatchContent *string `json:"matchContent,omitempty" xml:"matchContent,omitempty" require:"true"`
}

func (s QueryAttackLogDetailInfoResponseDataDetailInfoWAFDEFEND) String() string {
  return tea.Prettify(s)
}

func (s QueryAttackLogDetailInfoResponseDataDetailInfoWAFDEFEND) GoString() string {
  return s.String()
}

func (s *QueryAttackLogDetailInfoResponseDataDetailInfoWAFDEFEND) SetRuleType(v string) *QueryAttackLogDetailInfoResponseDataDetailInfoWAFDEFEND {
  s.RuleType = &v
  return s
}

func (s *QueryAttackLogDetailInfoResponseDataDetailInfoWAFDEFEND) SetRuleAction(v string) *QueryAttackLogDetailInfoResponseDataDetailInfoWAFDEFEND {
  s.RuleAction = &v
  return s
}

func (s *QueryAttackLogDetailInfoResponseDataDetailInfoWAFDEFEND) SetRuleId(v string) *QueryAttackLogDetailInfoResponseDataDetailInfoWAFDEFEND {
  s.RuleId = &v
  return s
}

func (s *QueryAttackLogDetailInfoResponseDataDetailInfoWAFDEFEND) SetRuleName(v string) *QueryAttackLogDetailInfoResponseDataDetailInfoWAFDEFEND {
  s.RuleName = &v
  return s
}

func (s *QueryAttackLogDetailInfoResponseDataDetailInfoWAFDEFEND) SetRuleDesc(v string) *QueryAttackLogDetailInfoResponseDataDetailInfoWAFDEFEND {
  s.RuleDesc = &v
  return s
}

func (s *QueryAttackLogDetailInfoResponseDataDetailInfoWAFDEFEND) SetMatchArea(v string) *QueryAttackLogDetailInfoResponseDataDetailInfoWAFDEFEND {
  s.MatchArea = &v
  return s
}

func (s *QueryAttackLogDetailInfoResponseDataDetailInfoWAFDEFEND) SetMatchContent(v string) *QueryAttackLogDetailInfoResponseDataDetailInfoWAFDEFEND {
  s.MatchContent = &v
  return s
}

type QueryAttackLogDetailInfoResponseDataDetailInfoDMSDEFEND struct     {
  // {"en":"Policy name.","zh_CN":"策略名称。"}
  PolicyName *string `json:"policyName,omitempty" xml:"policyName,omitempty" require:"true"`
  // {"en":"Rule action.","zh_CN":"规则动作。"}
  RuleAction *string `json:"ruleAction,omitempty" xml:"ruleAction,omitempty" require:"true"`
  // {"en":"Rule ID.","zh_CN":"规则ID。"}
  RuleId *string `json:"ruleId,omitempty" xml:"ruleId,omitempty" require:"true"`
  // {"en":"Rule name.","zh_CN":"规则名称。"}
  RuleName *string `json:"ruleName,omitempty" xml:"ruleName,omitempty" require:"true"`
  // {"en":"Description.","zh_CN":"说明。"}
  Explain *string `json:"explain,omitempty" xml:"explain,omitempty" require:"true"`
}

func (s QueryAttackLogDetailInfoResponseDataDetailInfoDMSDEFEND) String() string {
  return tea.Prettify(s)
}

func (s QueryAttackLogDetailInfoResponseDataDetailInfoDMSDEFEND) GoString() string {
  return s.String()
}

func (s *QueryAttackLogDetailInfoResponseDataDetailInfoDMSDEFEND) SetPolicyName(v string) *QueryAttackLogDetailInfoResponseDataDetailInfoDMSDEFEND {
  s.PolicyName = &v
  return s
}

func (s *QueryAttackLogDetailInfoResponseDataDetailInfoDMSDEFEND) SetRuleAction(v string) *QueryAttackLogDetailInfoResponseDataDetailInfoDMSDEFEND {
  s.RuleAction = &v
  return s
}

func (s *QueryAttackLogDetailInfoResponseDataDetailInfoDMSDEFEND) SetRuleId(v string) *QueryAttackLogDetailInfoResponseDataDetailInfoDMSDEFEND {
  s.RuleId = &v
  return s
}

func (s *QueryAttackLogDetailInfoResponseDataDetailInfoDMSDEFEND) SetRuleName(v string) *QueryAttackLogDetailInfoResponseDataDetailInfoDMSDEFEND {
  s.RuleName = &v
  return s
}

func (s *QueryAttackLogDetailInfoResponseDataDetailInfoDMSDEFEND) SetExplain(v string) *QueryAttackLogDetailInfoResponseDataDetailInfoDMSDEFEND {
  s.Explain = &v
  return s
}

type QueryAttackLogDetailInfoResponseDataDetailInfoBOTMANAGE struct     {
  // {"en":"Policy name.","zh_CN":"策略名称。"}
  PolicyName *string `json:"policyName,omitempty" xml:"policyName,omitempty" require:"true"`
  // {"en":"Rule action.","zh_CN":"规则动作。"}
  RuleAction *string `json:"ruleAction,omitempty" xml:"ruleAction,omitempty" require:"true"`
  // {"en":"Rule name.","zh_CN":"规则名称。"}
  RuleName *string `json:"ruleName,omitempty" xml:"ruleName,omitempty" require:"true"`
  // {"en":"Bot Category.","zh_CN":"Bot分类。"}
  BotCategory *string `json:"botCategory,omitempty" xml:"botCategory,omitempty" require:"true"`
  // {"en":"Bot tag.","zh_CN":"Bot标签。"}
  BotLabel *string `json:"botLabel,omitempty" xml:"botLabel,omitempty" require:"true"`
  // {"en":"Event description.","zh_CN":"事件描述。"}
  EventDesc *string `json:"eventDesc,omitempty" xml:"eventDesc,omitempty" require:"true"`
}

func (s QueryAttackLogDetailInfoResponseDataDetailInfoBOTMANAGE) String() string {
  return tea.Prettify(s)
}

func (s QueryAttackLogDetailInfoResponseDataDetailInfoBOTMANAGE) GoString() string {
  return s.String()
}

func (s *QueryAttackLogDetailInfoResponseDataDetailInfoBOTMANAGE) SetPolicyName(v string) *QueryAttackLogDetailInfoResponseDataDetailInfoBOTMANAGE {
  s.PolicyName = &v
  return s
}

func (s *QueryAttackLogDetailInfoResponseDataDetailInfoBOTMANAGE) SetRuleAction(v string) *QueryAttackLogDetailInfoResponseDataDetailInfoBOTMANAGE {
  s.RuleAction = &v
  return s
}

func (s *QueryAttackLogDetailInfoResponseDataDetailInfoBOTMANAGE) SetRuleName(v string) *QueryAttackLogDetailInfoResponseDataDetailInfoBOTMANAGE {
  s.RuleName = &v
  return s
}

func (s *QueryAttackLogDetailInfoResponseDataDetailInfoBOTMANAGE) SetBotCategory(v string) *QueryAttackLogDetailInfoResponseDataDetailInfoBOTMANAGE {
  s.BotCategory = &v
  return s
}

func (s *QueryAttackLogDetailInfoResponseDataDetailInfoBOTMANAGE) SetBotLabel(v string) *QueryAttackLogDetailInfoResponseDataDetailInfoBOTMANAGE {
  s.BotLabel = &v
  return s
}

func (s *QueryAttackLogDetailInfoResponseDataDetailInfoBOTMANAGE) SetEventDesc(v string) *QueryAttackLogDetailInfoResponseDataDetailInfoBOTMANAGE {
  s.EventDesc = &v
  return s
}

type QueryAttackLogDetailInfoResponseDataDetailInfoAPIDEFEND struct     {
  // {"en":"Rule action.","zh_CN":"规则动作。"}
  RuleAction *string `json:"ruleAction,omitempty" xml:"ruleAction,omitempty" require:"true"`
  // {"en":"Policy name.","zh_CN":"策略名称。"}
  PolicyName *string `json:"policyName,omitempty" xml:"policyName,omitempty" require:"true"`
  // {"en":"API name.","zh_CN":"API名称。"}
  ApiName *string `json:"apiName,omitempty" xml:"apiName,omitempty" require:"true"`
  // {"en":"API ID.","zh_CN":"API ID。"}
  ApiId *string `json:"apiId,omitempty" xml:"apiId,omitempty" require:"true"`
  // {"en":"English instructions.","zh_CN":"英文说明。"}
  ExplainEn *string `json:"explainEn,omitempty" xml:"explainEn,omitempty" require:"true"`
  // {"en":"Chinese instructions, only supported in language=zh_CN","zh_CN":"中文说明，仅当language=zh_CN支持。"}
  ExplainCn *string `json:"explainCn,omitempty" xml:"explainCn,omitempty" require:"true"`
}

func (s QueryAttackLogDetailInfoResponseDataDetailInfoAPIDEFEND) String() string {
  return tea.Prettify(s)
}

func (s QueryAttackLogDetailInfoResponseDataDetailInfoAPIDEFEND) GoString() string {
  return s.String()
}

func (s *QueryAttackLogDetailInfoResponseDataDetailInfoAPIDEFEND) SetRuleAction(v string) *QueryAttackLogDetailInfoResponseDataDetailInfoAPIDEFEND {
  s.RuleAction = &v
  return s
}

func (s *QueryAttackLogDetailInfoResponseDataDetailInfoAPIDEFEND) SetPolicyName(v string) *QueryAttackLogDetailInfoResponseDataDetailInfoAPIDEFEND {
  s.PolicyName = &v
  return s
}

func (s *QueryAttackLogDetailInfoResponseDataDetailInfoAPIDEFEND) SetApiName(v string) *QueryAttackLogDetailInfoResponseDataDetailInfoAPIDEFEND {
  s.ApiName = &v
  return s
}

func (s *QueryAttackLogDetailInfoResponseDataDetailInfoAPIDEFEND) SetApiId(v string) *QueryAttackLogDetailInfoResponseDataDetailInfoAPIDEFEND {
  s.ApiId = &v
  return s
}

func (s *QueryAttackLogDetailInfoResponseDataDetailInfoAPIDEFEND) SetExplainEn(v string) *QueryAttackLogDetailInfoResponseDataDetailInfoAPIDEFEND {
  s.ExplainEn = &v
  return s
}

func (s *QueryAttackLogDetailInfoResponseDataDetailInfoAPIDEFEND) SetExplainCn(v string) *QueryAttackLogDetailInfoResponseDataDetailInfoAPIDEFEND {
  s.ExplainCn = &v
  return s
}

type QueryAttackLogDetailInfoResponseDataDetailInfoINTELLIGENCE struct     {
  // {"en":"Threat type.","zh_CN":"情报类型。"}
  IntelligentType *string `json:"intelligentType,omitempty" xml:"intelligentType,omitempty" require:"true"`
  // {"en":"Rule action.","zh_CN":"规则动作。"}
  RuleAction *string `json:"ruleAction,omitempty" xml:"ruleAction,omitempty" require:"true"`
  // {"en":"Intelligence module.","zh_CN":"情报模块。"}
  IntelligentModule *string `json:"intelligentModule,omitempty" xml:"intelligentModule,omitempty" require:"true"`
}

func (s QueryAttackLogDetailInfoResponseDataDetailInfoINTELLIGENCE) String() string {
  return tea.Prettify(s)
}

func (s QueryAttackLogDetailInfoResponseDataDetailInfoINTELLIGENCE) GoString() string {
  return s.String()
}

func (s *QueryAttackLogDetailInfoResponseDataDetailInfoINTELLIGENCE) SetIntelligentType(v string) *QueryAttackLogDetailInfoResponseDataDetailInfoINTELLIGENCE {
  s.IntelligentType = &v
  return s
}

func (s *QueryAttackLogDetailInfoResponseDataDetailInfoINTELLIGENCE) SetRuleAction(v string) *QueryAttackLogDetailInfoResponseDataDetailInfoINTELLIGENCE {
  s.RuleAction = &v
  return s
}

func (s *QueryAttackLogDetailInfoResponseDataDetailInfoINTELLIGENCE) SetIntelligentModule(v string) *QueryAttackLogDetailInfoResponseDataDetailInfoINTELLIGENCE {
  s.IntelligentModule = &v
  return s
}

type QueryAttackLogDetailInfoResponseDataDetailInfoRATELIMIT struct     {
  // {"en":"Rule action.","zh_CN":"规则动作。"}
  RuleAction *string `json:"ruleAction,omitempty" xml:"ruleAction,omitempty" require:"true"`
  // {"en":"API name.","zh_CN":"API名称。"}
  ApiName *string `json:"apiName,omitempty" xml:"apiName,omitempty" require:"true"`
  // {"en":"API ID.","zh_CN":"API ID。"}
  ApiId *string `json:"apiId,omitempty" xml:"apiId,omitempty" require:"true"`
  // {"en":"Rule name.","zh_CN":"规则名称。"}
  RuleName *string `json:"ruleName,omitempty" xml:"ruleName,omitempty" require:"true"`
  // {"en":"Rule ID.","zh_CN":"规则ID。"}
  RuleId *string `json:"ruleId,omitempty" xml:"ruleId,omitempty" require:"true"`
  // {"en":"Rule description.","zh_CN":"规则描述。"}
  RuleDesc *string `json:"ruleDesc,omitempty" xml:"ruleDesc,omitempty" require:"true"`
}

func (s QueryAttackLogDetailInfoResponseDataDetailInfoRATELIMIT) String() string {
  return tea.Prettify(s)
}

func (s QueryAttackLogDetailInfoResponseDataDetailInfoRATELIMIT) GoString() string {
  return s.String()
}

func (s *QueryAttackLogDetailInfoResponseDataDetailInfoRATELIMIT) SetRuleAction(v string) *QueryAttackLogDetailInfoResponseDataDetailInfoRATELIMIT {
  s.RuleAction = &v
  return s
}

func (s *QueryAttackLogDetailInfoResponseDataDetailInfoRATELIMIT) SetApiName(v string) *QueryAttackLogDetailInfoResponseDataDetailInfoRATELIMIT {
  s.ApiName = &v
  return s
}

func (s *QueryAttackLogDetailInfoResponseDataDetailInfoRATELIMIT) SetApiId(v string) *QueryAttackLogDetailInfoResponseDataDetailInfoRATELIMIT {
  s.ApiId = &v
  return s
}

func (s *QueryAttackLogDetailInfoResponseDataDetailInfoRATELIMIT) SetRuleName(v string) *QueryAttackLogDetailInfoResponseDataDetailInfoRATELIMIT {
  s.RuleName = &v
  return s
}

func (s *QueryAttackLogDetailInfoResponseDataDetailInfoRATELIMIT) SetRuleId(v string) *QueryAttackLogDetailInfoResponseDataDetailInfoRATELIMIT {
  s.RuleId = &v
  return s
}

func (s *QueryAttackLogDetailInfoResponseDataDetailInfoRATELIMIT) SetRuleDesc(v string) *QueryAttackLogDetailInfoResponseDataDetailInfoRATELIMIT {
  s.RuleDesc = &v
  return s
}

type QueryAttackLogDetailInfoResponseDataDetailInfoCUSTOMIZERULE struct     {
  // {"en":"Rule action.","zh_CN":"规则动作。"}
  RuleAction *string `json:"ruleAction,omitempty" xml:"ruleAction,omitempty" require:"true"`
  // {"en":"API name.","zh_CN":"API名称。"}
  ApiName *string `json:"apiName,omitempty" xml:"apiName,omitempty" require:"true"`
  // {"en":"API ID.","zh_CN":"API ID。"}
  ApiId *string `json:"apiId,omitempty" xml:"apiId,omitempty" require:"true"`
  // {"en":"Rule name.","zh_CN":"规则名称。"}
  RuleName *string `json:"ruleName,omitempty" xml:"ruleName,omitempty" require:"true"`
  // {"en":"Rule ID.","zh_CN":"规则ID。"}
  RuleId *string `json:"ruleId,omitempty" xml:"ruleId,omitempty" require:"true"`
  // {"en":"Rule description.","zh_CN":"规则描述。"}
  RuleDesc *string `json:"ruleDesc,omitempty" xml:"ruleDesc,omitempty" require:"true"`
}

func (s QueryAttackLogDetailInfoResponseDataDetailInfoCUSTOMIZERULE) String() string {
  return tea.Prettify(s)
}

func (s QueryAttackLogDetailInfoResponseDataDetailInfoCUSTOMIZERULE) GoString() string {
  return s.String()
}

func (s *QueryAttackLogDetailInfoResponseDataDetailInfoCUSTOMIZERULE) SetRuleAction(v string) *QueryAttackLogDetailInfoResponseDataDetailInfoCUSTOMIZERULE {
  s.RuleAction = &v
  return s
}

func (s *QueryAttackLogDetailInfoResponseDataDetailInfoCUSTOMIZERULE) SetApiName(v string) *QueryAttackLogDetailInfoResponseDataDetailInfoCUSTOMIZERULE {
  s.ApiName = &v
  return s
}

func (s *QueryAttackLogDetailInfoResponseDataDetailInfoCUSTOMIZERULE) SetApiId(v string) *QueryAttackLogDetailInfoResponseDataDetailInfoCUSTOMIZERULE {
  s.ApiId = &v
  return s
}

func (s *QueryAttackLogDetailInfoResponseDataDetailInfoCUSTOMIZERULE) SetRuleName(v string) *QueryAttackLogDetailInfoResponseDataDetailInfoCUSTOMIZERULE {
  s.RuleName = &v
  return s
}

func (s *QueryAttackLogDetailInfoResponseDataDetailInfoCUSTOMIZERULE) SetRuleId(v string) *QueryAttackLogDetailInfoResponseDataDetailInfoCUSTOMIZERULE {
  s.RuleId = &v
  return s
}

func (s *QueryAttackLogDetailInfoResponseDataDetailInfoCUSTOMIZERULE) SetRuleDesc(v string) *QueryAttackLogDetailInfoResponseDataDetailInfoCUSTOMIZERULE {
  s.RuleDesc = &v
  return s
}

type QueryAttackLogDetailInfoResponseHeader struct {
}

func (s QueryAttackLogDetailInfoResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryAttackLogDetailInfoResponseHeader) GoString() string {
  return s.String()
}




type GetPrivacyStatusDistributionRequest struct {
  // {"en":"List of API group.", "zh_CN":"API分组列表。"}
  ApiGroups []*string `json:"apiGroups,omitempty" xml:"apiGroups,omitempty" type:"Repeated"`
  // {"en":"List of API name.", "zh_CN":"API名称列表。"}
  ApiIds []*string `json:"apiIds,omitempty" xml:"apiIds,omitempty" type:"Repeated"`
  // {"en":"List of domain.", "zh_CN":"域名列表。"}
  Domains []*string `json:"domains,omitempty" xml:"domains,omitempty" type:"Repeated"`
  // {"en":"Start time, format: yyyy-MM-dd HH:mm:ss.", "zh_CN":"开始时间，格式：yyyy-MM-dd HH:mm:ss。"}
  StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty" require:"true"`
  // {"en":"End time, format: yyyy-MM-dd HH:mm:ss.", "zh_CN":"结束时间，格式：yyyy-MM-dd HH:mm:ss。"}
  EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty" require:"true"`
  // {"en":"Endpoint path.", "zh_CN":"前端路径。"}
  FrontPath *string `json:"frontPath,omitempty" xml:"frontPath,omitempty"`
}

func (s GetPrivacyStatusDistributionRequest) String() string {
  return tea.Prettify(s)
}

func (s GetPrivacyStatusDistributionRequest) GoString() string {
  return s.String()
}

func (s *GetPrivacyStatusDistributionRequest) SetApiGroups(v []*string) *GetPrivacyStatusDistributionRequest {
  s.ApiGroups = v
  return s
}

func (s *GetPrivacyStatusDistributionRequest) SetApiIds(v []*string) *GetPrivacyStatusDistributionRequest {
  s.ApiIds = v
  return s
}

func (s *GetPrivacyStatusDistributionRequest) SetDomains(v []*string) *GetPrivacyStatusDistributionRequest {
  s.Domains = v
  return s
}

func (s *GetPrivacyStatusDistributionRequest) SetStartTime(v string) *GetPrivacyStatusDistributionRequest {
  s.StartTime = &v
  return s
}

func (s *GetPrivacyStatusDistributionRequest) SetEndTime(v string) *GetPrivacyStatusDistributionRequest {
  s.EndTime = &v
  return s
}

func (s *GetPrivacyStatusDistributionRequest) SetFrontPath(v string) *GetPrivacyStatusDistributionRequest {
  s.FrontPath = &v
  return s
}

type GetPrivacyStatusDistributionResponse struct {
  // {"en":"Return 200 means success.", "zh_CN":"200状态码表示请求成功。"}
  Code *int `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Message.", "zh_CN":"返回信息。"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"Data", "zh_CN":"数据。"}
  Data *GetPrivacyStatusDistributionVo `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s GetPrivacyStatusDistributionResponse) String() string {
  return tea.Prettify(s)
}

func (s GetPrivacyStatusDistributionResponse) GoString() string {
  return s.String()
}

func (s *GetPrivacyStatusDistributionResponse) SetCode(v int) *GetPrivacyStatusDistributionResponse {
  s.Code = &v
  return s
}

func (s *GetPrivacyStatusDistributionResponse) SetMsg(v string) *GetPrivacyStatusDistributionResponse {
  s.Msg = &v
  return s
}

func (s *GetPrivacyStatusDistributionResponse) SetData(v *GetPrivacyStatusDistributionVo) *GetPrivacyStatusDistributionResponse {
  s.Data = v
  return s
}

type GetPrivacyStatusDistributionVo struct {
  // {"en":"The number of private API.", "zh_CN":"私有API个数。"}
  PrivateNum *int `json:"privateNum,omitempty" xml:"privateNum,omitempty" require:"true"`
  // {"en":"The number of public API.", "zh_CN":"公共API个数。"}
  PublicNum *int `json:"publicNum,omitempty" xml:"publicNum,omitempty" require:"true"`
}

func (s GetPrivacyStatusDistributionVo) String() string {
  return tea.Prettify(s)
}

func (s GetPrivacyStatusDistributionVo) GoString() string {
  return s.String()
}

func (s *GetPrivacyStatusDistributionVo) SetPrivateNum(v int) *GetPrivacyStatusDistributionVo {
  s.PrivateNum = &v
  return s
}

func (s *GetPrivacyStatusDistributionVo) SetPublicNum(v int) *GetPrivacyStatusDistributionVo {
  s.PublicNum = &v
  return s
}

type GetPrivacyStatusDistributionPaths struct {
}

func (s GetPrivacyStatusDistributionPaths) String() string {
  return tea.Prettify(s)
}

func (s GetPrivacyStatusDistributionPaths) GoString() string {
  return s.String()
}

type GetPrivacyStatusDistributionParameters struct {
}

func (s GetPrivacyStatusDistributionParameters) String() string {
  return tea.Prettify(s)
}

func (s GetPrivacyStatusDistributionParameters) GoString() string {
  return s.String()
}

type GetPrivacyStatusDistributionRequestHeader struct {
}

func (s GetPrivacyStatusDistributionRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetPrivacyStatusDistributionRequestHeader) GoString() string {
  return s.String()
}

type GetPrivacyStatusDistributionResponseHeader struct {
}

func (s GetPrivacyStatusDistributionResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetPrivacyStatusDistributionResponseHeader) GoString() string {
  return s.String()
}




