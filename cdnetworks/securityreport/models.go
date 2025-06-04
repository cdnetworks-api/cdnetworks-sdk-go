package securityreport

// This file is auto-generated, don't edit it. Thanks.
import (
  "github.com/alibabacloud-go/tea/tea"
)

type L7DdosTrendRequest struct {
  // {"en":"Start time, format: yyyy-MM-dd HH:mm:ss.
  // The time range does not exceed 30 days.", "zh_CN":"开始时间，格式：yyyy-MM-dd HH:mm:ss。
  // 时间范围不超过30天。"}
  StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty" require:"true"`
  // {"en":"End time, format: yyyy-MM-dd HH:mm:ss.
  // The time range does not exceed 30 days)", "zh_CN":"结束时间，格式：yyyy-MM-dd HH:mm:ss。
  // 时间范围不超过30天。"}
  EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty" require:"true"`
  // {"en":"Hostname list.", "zh_CN":"域名数组。"}
  Domains []*string `json:"domains,omitempty" xml:"domains,omitempty" require:"true" type:"Repeated"`
}

func (s L7DdosTrendRequest) String() string {
  return tea.Prettify(s)
}

func (s L7DdosTrendRequest) GoString() string {
  return s.String()
}

func (s *L7DdosTrendRequest) SetStartTime(v string) *L7DdosTrendRequest {
  s.StartTime = &v
  return s
}

func (s *L7DdosTrendRequest) SetEndTime(v string) *L7DdosTrendRequest {
  s.EndTime = &v
  return s
}

func (s *L7DdosTrendRequest) SetDomains(v []*string) *L7DdosTrendRequest {
  s.Domains = v
  return s
}

type L7DdosTrendResponse struct {
  // {'en':'Please refer to the error code for exceptions.', 'zh_CN':'请参照错误码。','dictionary':'belong=WAAP-MS-Ext|dict=waap_retCodeEnum'}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Description.", "zh_CN":"描述信息。"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"Data.", "zh_CN":"出参数据。"}
  Data []*L7DdosTrendAttackedUrl `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
  // {"en":"L7DdosTrendPeak Attack QPS", "zh_CN":"CC攻击QPS峰值"}
  Peak_qps []*L7DdosTrendPeak `json:"peak_qps,omitempty" xml:"peak_qps,omitempty" require:"true" type:"Repeated"`
}

func (s L7DdosTrendResponse) String() string {
  return tea.Prettify(s)
}

func (s L7DdosTrendResponse) GoString() string {
  return s.String()
}

func (s *L7DdosTrendResponse) SetCode(v string) *L7DdosTrendResponse {
  s.Code = &v
  return s
}

func (s *L7DdosTrendResponse) SetMsg(v string) *L7DdosTrendResponse {
  s.Msg = &v
  return s
}

func (s *L7DdosTrendResponse) SetData(v []*L7DdosTrendAttackedUrl) *L7DdosTrendResponse {
  s.Data = v
  return s
}

func (s *L7DdosTrendResponse) SetPeak_qps(v []*L7DdosTrendPeak) *L7DdosTrendResponse {
  s.Peak_qps = v
  return s
}

type L7DdosTrendAttackedUrl struct {
  // {"en":"All Requests(QPS)", "zh_CN":"所有请求（QPS）"}
  All_count *int64 `json:"all_count,omitempty" xml:"all_count,omitempty" require:"true"`
  // {"en":"L7 DDoS Attack(QPS)", "zh_CN":"CC攻击（QPS）"}
  Attack_count *int64 `json:"attack_count,omitempty" xml:"attack_count,omitempty" require:"true"`
  // {"en":"Managed Ruleset Protection(QPS)", "zh_CN":"内置规则防护（QPS）"}
  Def_count *int64 `json:"def_count,omitempty" xml:"def_count,omitempty" require:"true"`
  // {"en":"Adaptive DDoS Protection(QPS)", "zh_CN":"AI智能防护（QPS）"}
  Ai_count *int64 `json:"ai_count,omitempty" xml:"ai_count,omitempty" require:"true"`
  // {"en":"IP Block(QPS)", "zh_CN":"IP封禁（QPS）"}
  Ip_count *int64 `json:"ip_count,omitempty" xml:"ip_count,omitempty" require:"true"`
  // {"en":"Geo Block(QPS)", "zh_CN":"区域封禁（QPS）"}
  Area_count *int64 `json:"area_count,omitempty" xml:"area_count,omitempty" require:"true"`
  // {"en":"Custom Rules(QPS)", "zh_CN":"自定义规则（QPS）"}
  Rule_count *int64 `json:"rule_count,omitempty" xml:"rule_count,omitempty" require:"true"`
  // {"en":"Rate Limit(QPS)", "zh_CN":"频率限制（QPS）"}
  Limit_count *int64 `json:"limit_count,omitempty" xml:"limit_count,omitempty" require:"true"`
  // {"en":"Time.", "zh_CN":"时间。"}
  Time *string `json:"time,omitempty" xml:"time,omitempty" require:"true"`
}

func (s L7DdosTrendAttackedUrl) String() string {
  return tea.Prettify(s)
}

func (s L7DdosTrendAttackedUrl) GoString() string {
  return s.String()
}

func (s *L7DdosTrendAttackedUrl) SetAll_count(v int64) *L7DdosTrendAttackedUrl {
  s.All_count = &v
  return s
}

func (s *L7DdosTrendAttackedUrl) SetAttack_count(v int64) *L7DdosTrendAttackedUrl {
  s.Attack_count = &v
  return s
}

func (s *L7DdosTrendAttackedUrl) SetDef_count(v int64) *L7DdosTrendAttackedUrl {
  s.Def_count = &v
  return s
}

func (s *L7DdosTrendAttackedUrl) SetAi_count(v int64) *L7DdosTrendAttackedUrl {
  s.Ai_count = &v
  return s
}

func (s *L7DdosTrendAttackedUrl) SetIp_count(v int64) *L7DdosTrendAttackedUrl {
  s.Ip_count = &v
  return s
}

func (s *L7DdosTrendAttackedUrl) SetArea_count(v int64) *L7DdosTrendAttackedUrl {
  s.Area_count = &v
  return s
}

func (s *L7DdosTrendAttackedUrl) SetRule_count(v int64) *L7DdosTrendAttackedUrl {
  s.Rule_count = &v
  return s
}

func (s *L7DdosTrendAttackedUrl) SetLimit_count(v int64) *L7DdosTrendAttackedUrl {
  s.Limit_count = &v
  return s
}

func (s *L7DdosTrendAttackedUrl) SetTime(v string) *L7DdosTrendAttackedUrl {
  s.Time = &v
  return s
}

type L7DdosTrendPeak struct {
  // {"en":"L7DdosTrendPeak Attack Value", "zh_CN":"峰值（QPS）"}
  Peak_value *int64 `json:"peak_value,omitempty" xml:"peak_value,omitempty" require:"true"`
  // {"en":"L7DdosTrendPeak Attack Time", "zh_CN":"峰值时间"}
  Peak_time *string `json:"peak_time,omitempty" xml:"peak_time,omitempty" require:"true"`
}

func (s L7DdosTrendPeak) String() string {
  return tea.Prettify(s)
}

func (s L7DdosTrendPeak) GoString() string {
  return s.String()
}

func (s *L7DdosTrendPeak) SetPeak_value(v int64) *L7DdosTrendPeak {
  s.Peak_value = &v
  return s
}

func (s *L7DdosTrendPeak) SetPeak_time(v string) *L7DdosTrendPeak {
  s.Peak_time = &v
  return s
}

type L7DdosTrendPaths struct {
}

func (s L7DdosTrendPaths) String() string {
  return tea.Prettify(s)
}

func (s L7DdosTrendPaths) GoString() string {
  return s.String()
}

type L7DdosTrendParameters struct {
}

func (s L7DdosTrendParameters) String() string {
  return tea.Prettify(s)
}

func (s L7DdosTrendParameters) GoString() string {
  return s.String()
}

type L7DdosTrendRequestHeader struct {
  // {"zh_CN":"安全服务类型。有使用多个不同的安全服务时，需要填写具体的服务类型。","en":"Security service type. Please enter a specific service type, if you purchase multiple security services.","dictionary":"belong=WAAP-MS-Ext|dict=waap_serviceType"}
  ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty"`
}

func (s L7DdosTrendRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s L7DdosTrendRequestHeader) GoString() string {
  return s.String()
}

func (s *L7DdosTrendRequestHeader) SetServiceType(v string) *L7DdosTrendRequestHeader {
  s.ServiceType = &v
  return s
}

type L7DdosTrendResponseHeader struct {
}

func (s L7DdosTrendResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s L7DdosTrendResponseHeader) GoString() string {
  return s.String()
}




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
  // {'en':'Top rankings, default value: 10, max: 1000.', 'zh_CN':'取前几排名，默认：10，上限：1000。'}
  Top *int32 `json:"top,omitempty" xml:"top,omitempty"`
  // {'en':'Multiple choice. Disposal result, default value: all results.
  //  mitigated: Mitegaed requests.
  //  monitored: monitored requests.', 'zh_CN':'多选。处置结果，默认：展示所有结果。
  //  mitigated：已抵御请求数。
  //  monitored：观察请求数。'}
  ActType []*string `json:"actType,omitempty" xml:"actType,omitempty" type:"Repeated"`
  // {'en':'Start time, format: yyyy-MM-dd HH:mm:ss.
  // The time range does not exceed 31 days.', 'zh_CN':'起始时间，格式：yyyy-MM-dd HH:mm:ss。
  // 时间范围不超过31天。'}
  StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty" require:"true"`
  // {'en':'End time, format: yyyy-MM-dd HH:mm:ss.
  // The time range does not exceed 31 days.', 'zh_CN':'截止时间，格式：yyyy-MM-dd HH:mm:ss。
  // 时间范围不超过31天。'}
  EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty" require:"true"`
  // {'en':'Hostname list, if not specified, it means all the hostnames of the account.', 'zh_CN':'域名列表，未指定时查询账号下的所有域名。'}
  Domains []*string `json:"domains,omitempty" xml:"domains,omitempty" type:"Repeated"`
}

func (s GetTriggeredCustomRulesRequest) String() string {
  return tea.Prettify(s)
}

func (s GetTriggeredCustomRulesRequest) GoString() string {
  return s.String()
}

func (s *GetTriggeredCustomRulesRequest) SetTop(v int32) *GetTriggeredCustomRulesRequest {
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

type GetTriggeredCustomRulesResponse struct {
  // {'en':'Please refer to the error code for exceptions.', 'zh_CN':'请参照错误码。','dictionary':'belong=WAAP-MS-Ext|dict=waap_retCodeEnum'}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {'en':'Description.', 'zh_CN':'描述信息。'}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {'en':'Data.', 'zh_CN':'出参数据。'}
  Data []*GetTriggeredCustomRulesRuleTopDTO `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
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

func (s *GetTriggeredCustomRulesResponse) SetData(v []*GetTriggeredCustomRulesRuleTopDTO) *GetTriggeredCustomRulesResponse {
  s.Data = v
  return s
}

type GetTriggeredCustomRulesRuleHitDTO struct {
  // {'en':'Action.', 'zh_CN':'采取动作。'}
  Act *string `json:"act,omitempty" xml:"act,omitempty" require:"true"`
  // {'en':'Hit times.', 'zh_CN':'命中数。'}
  Value *int64 `json:"value,omitempty" xml:"value,omitempty" require:"true"`
}

func (s GetTriggeredCustomRulesRuleHitDTO) String() string {
  return tea.Prettify(s)
}

func (s GetTriggeredCustomRulesRuleHitDTO) GoString() string {
  return s.String()
}

func (s *GetTriggeredCustomRulesRuleHitDTO) SetAct(v string) *GetTriggeredCustomRulesRuleHitDTO {
  s.Act = &v
  return s
}

func (s *GetTriggeredCustomRulesRuleHitDTO) SetValue(v int64) *GetTriggeredCustomRulesRuleHitDTO {
  s.Value = &v
  return s
}

type GetTriggeredCustomRulesRuleTopDTO struct {
  // {'en':'Rule ID.', 'zh_CN':'规则ID。'}
  RuleId *string `json:"ruleId,omitempty" xml:"ruleId,omitempty" require:"true"`
  // {'en':'Rule name.', 'zh_CN':'规则名称。'}
  RuleName *string `json:"ruleName,omitempty" xml:"ruleName,omitempty" require:"true"`
  // {'en':'Protected target.', 'zh_CN':'业务场景。'}
  Scene *string `json:"scene,omitempty" xml:"scene,omitempty" require:"true"`
  // {'en':'Trigger times, sort by times.', 'zh_CN':'触发次数，按次数排序。'}
  Hits []*GetTriggeredCustomRulesRuleHitDTO `json:"hits,omitempty" xml:"hits,omitempty" require:"true" type:"Repeated"`
}

func (s GetTriggeredCustomRulesRuleTopDTO) String() string {
  return tea.Prettify(s)
}

func (s GetTriggeredCustomRulesRuleTopDTO) GoString() string {
  return s.String()
}

func (s *GetTriggeredCustomRulesRuleTopDTO) SetRuleId(v string) *GetTriggeredCustomRulesRuleTopDTO {
  s.RuleId = &v
  return s
}

func (s *GetTriggeredCustomRulesRuleTopDTO) SetRuleName(v string) *GetTriggeredCustomRulesRuleTopDTO {
  s.RuleName = &v
  return s
}

func (s *GetTriggeredCustomRulesRuleTopDTO) SetScene(v string) *GetTriggeredCustomRulesRuleTopDTO {
  s.Scene = &v
  return s
}

func (s *GetTriggeredCustomRulesRuleTopDTO) SetHits(v []*GetTriggeredCustomRulesRuleHitDTO) *GetTriggeredCustomRulesRuleTopDTO {
  s.Hits = v
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

type GetTriggeredCustomRulesRequestHeader struct {
  // {'en':'The language of response data, default value: en.
  // zh_CN: Chinese
  // en: English', 'zh_CN':'返回内容的语言版本，默认值: en。
  // zh_CN：中文
  // en：英文'}
  Language *string `json:"language,omitempty" xml:"language,omitempty"`
  // {"zh_CN":"安全服务类型。有使用多个不同的安全服务时，需要填写具体的服务类型。","en":"Security service type. Please enter a specific service type, if you purchase multiple security services.","dictionary":"belong=WAAP-MS-Ext|dict=waap_serviceType"}
  ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty"`
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

type GetTriggeredCustomRulesResponseHeader struct {
}

func (s GetTriggeredCustomRulesResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetTriggeredCustomRulesResponseHeader) GoString() string {
  return s.String()
}




type L4DdosEventRequest struct {
  // {"en":"Start time, format: yyyy-MM-dd HH:mm:ss.
  // The time range does not exceed 30 days.", "zh_CN":"开始时间，格式：yyyy-MM-dd HH:mm:ss。
  // 时间范围不超过30天。"}
  StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty" require:"true"`
  // {"en":"End time, format: yyyy-MM-dd HH:mm:ss.
  // The time range does not exceed 30 days)", "zh_CN":"结束时间，格式：yyyy-MM-dd HH:mm:ss。
  // 时间范围不超过30天。"}
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

type L4DdosEventResponse struct {
  // {'en':'Please refer to the error code for exceptions.', 'zh_CN':'请参照错误码。','dictionary':'belong=WAAP-MS-Ext|dict=waap_retCodeEnum'}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Description.", "zh_CN":"描述信息。"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"Data.", "zh_CN":"出参数据。"}
  Data []*L4DdosEventAttackedUrl `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
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

func (s *L4DdosEventResponse) SetData(v []*L4DdosEventAttackedUrl) *L4DdosEventResponse {
  s.Data = v
  return s
}

type L4DdosEventAttackedUrl struct {
  // {"en":"Duration.", "zh_CN":"持续时间。"}
  Duration *string `json:"duration,omitempty" xml:"duration,omitempty" require:"true"`
  // {"en":"Start time.", "zh_CN":"开始时间。"}
  Start_time *string `json:"start_time,omitempty" xml:"start_time,omitempty" require:"true"`
  // {"en":"Attack type.", "zh_CN":"攻击类型。"}
  Attack_type []*string `json:"attack_type,omitempty" xml:"attack_type,omitempty" require:"true" type:"Repeated"`
  // {"en":"End time.", "zh_CN":"结束时间。"}
  End_time *string `json:"end_time,omitempty" xml:"end_time,omitempty" require:"true"`
  // {"en":"Max time.", "zh_CN":"峰值时间。"}
  Max_time *string `json:"max_time,omitempty" xml:"max_time,omitempty" require:"true"`
  // {"en":"Status.", "zh_CN":"状态 1.清洗结束 0.清洗中。"}
  Status *string `json:"status,omitempty" xml:"status,omitempty" require:"true"`
  // {"en":"Peak bandwidth.", "zh_CN":"峰值（Mbps）。"}
  Max_value *string `json:"max_value,omitempty" xml:"max_value,omitempty" require:"true"`
}

func (s L4DdosEventAttackedUrl) String() string {
  return tea.Prettify(s)
}

func (s L4DdosEventAttackedUrl) GoString() string {
  return s.String()
}

func (s *L4DdosEventAttackedUrl) SetDuration(v string) *L4DdosEventAttackedUrl {
  s.Duration = &v
  return s
}

func (s *L4DdosEventAttackedUrl) SetStart_time(v string) *L4DdosEventAttackedUrl {
  s.Start_time = &v
  return s
}

func (s *L4DdosEventAttackedUrl) SetAttack_type(v []*string) *L4DdosEventAttackedUrl {
  s.Attack_type = v
  return s
}

func (s *L4DdosEventAttackedUrl) SetEnd_time(v string) *L4DdosEventAttackedUrl {
  s.End_time = &v
  return s
}

func (s *L4DdosEventAttackedUrl) SetMax_time(v string) *L4DdosEventAttackedUrl {
  s.Max_time = &v
  return s
}

func (s *L4DdosEventAttackedUrl) SetStatus(v string) *L4DdosEventAttackedUrl {
  s.Status = &v
  return s
}

func (s *L4DdosEventAttackedUrl) SetMax_value(v string) *L4DdosEventAttackedUrl {
  s.Max_value = &v
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

type L4DdosEventRequestHeader struct {
  // {"zh_CN":"安全服务类型。有使用多个不同的安全服务时，需要填写具体的服务类型。","en":"Security service type. Please enter a specific service type, if you purchase multiple security services.","dictionary":"belong=WAAP-MS-Ext|dict=waap_serviceType"}
  ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty"`
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




type L7DdosEventsRequest struct {
  // {"en":"Start time, format: yyyy-MM-dd HH:mm:ss.
  // The time range does not exceed 30 days.", "zh_CN":"开始时间，格式：yyyy-MM-dd HH:mm:ss。
  // 时间范围不超过30天。"}
  StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty" require:"true"`
  // {"en":"End time, format: yyyy-MM-dd HH:mm:ss.
  // The time range does not exceed 30 days)", "zh_CN":"结束时间，格式：yyyy-MM-dd HH:mm:ss。
  // 时间范围不超过30天。"}
  EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty" require:"true"`
  // {"en":"Hostname list.", "zh_CN":"域名数组。"}
  Hostnames []*string `json:"Hostnames,omitempty" xml:"Hostnames,omitempty" require:"true" type:"Repeated"`
}

func (s L7DdosEventsRequest) String() string {
  return tea.Prettify(s)
}

func (s L7DdosEventsRequest) GoString() string {
  return s.String()
}

func (s *L7DdosEventsRequest) SetStartTime(v string) *L7DdosEventsRequest {
  s.StartTime = &v
  return s
}

func (s *L7DdosEventsRequest) SetEndTime(v string) *L7DdosEventsRequest {
  s.EndTime = &v
  return s
}

func (s *L7DdosEventsRequest) SetHostnames(v []*string) *L7DdosEventsRequest {
  s.Hostnames = v
  return s
}

type L7DdosEventsResponse struct {
  // {'en':'Please refer to the error code for exceptions.', 'zh_CN':'请参照错误码。','dictionary':'belong=WAAP-MS-Ext|dict=waap_retCodeEnum'}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Description.", "zh_CN":"描述信息。"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"Data.", "zh_CN":"出参数据。"}
  Data []*L7DdosEventsAttackedUrl `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s L7DdosEventsResponse) String() string {
  return tea.Prettify(s)
}

func (s L7DdosEventsResponse) GoString() string {
  return s.String()
}

func (s *L7DdosEventsResponse) SetCode(v string) *L7DdosEventsResponse {
  s.Code = &v
  return s
}

func (s *L7DdosEventsResponse) SetMsg(v string) *L7DdosEventsResponse {
  s.Msg = &v
  return s
}

func (s *L7DdosEventsResponse) SetData(v []*L7DdosEventsAttackedUrl) *L7DdosEventsResponse {
  s.Data = v
  return s
}

type L7DdosEventsAttackedUrl struct {
  // {"en":"Duration(minutes)", "zh_CN":"持续时间。"}
  Duration *string `json:"duration,omitempty" xml:"duration,omitempty" require:"true"`
  // {"en":"Start Time", "zh_CN":"开始时间。"}
  Start_time *string `json:"start_time,omitempty" xml:"start_time,omitempty" require:"true"`
  // {"en":"Attacked Requests", "zh_CN":"攻击次数。"}
  Total_count *string `json:"total_count,omitempty" xml:"total_count,omitempty" require:"true"`
  // {"en":"End time", "zh_CN":"结束时间。"}
  End_time *string `json:"end_time,omitempty" xml:"end_time,omitempty" require:"true"`
  // {"en":"Attacked Hostnames", "zh_CN":"被攻击域名。"}
  Channel *string `json:"channel,omitempty" xml:"channel,omitempty" require:"true"`
  // {"en":"Peak Time", "zh_CN":"峰值时间。"}
  Max_time *string `json:"max_time,omitempty" xml:"max_time,omitempty" require:"true"`
  // {"en":"Status.(1. Mitigated 0. Mitigating)", "zh_CN":"状态。（ 1.清洗结束 0.清洗中）"}
  Status *string `json:"status,omitempty" xml:"status,omitempty" require:"true"`
}

func (s L7DdosEventsAttackedUrl) String() string {
  return tea.Prettify(s)
}

func (s L7DdosEventsAttackedUrl) GoString() string {
  return s.String()
}

func (s *L7DdosEventsAttackedUrl) SetDuration(v string) *L7DdosEventsAttackedUrl {
  s.Duration = &v
  return s
}

func (s *L7DdosEventsAttackedUrl) SetStart_time(v string) *L7DdosEventsAttackedUrl {
  s.Start_time = &v
  return s
}

func (s *L7DdosEventsAttackedUrl) SetTotal_count(v string) *L7DdosEventsAttackedUrl {
  s.Total_count = &v
  return s
}

func (s *L7DdosEventsAttackedUrl) SetEnd_time(v string) *L7DdosEventsAttackedUrl {
  s.End_time = &v
  return s
}

func (s *L7DdosEventsAttackedUrl) SetChannel(v string) *L7DdosEventsAttackedUrl {
  s.Channel = &v
  return s
}

func (s *L7DdosEventsAttackedUrl) SetMax_time(v string) *L7DdosEventsAttackedUrl {
  s.Max_time = &v
  return s
}

func (s *L7DdosEventsAttackedUrl) SetStatus(v string) *L7DdosEventsAttackedUrl {
  s.Status = &v
  return s
}

type L7DdosEventsPaths struct {
}

func (s L7DdosEventsPaths) String() string {
  return tea.Prettify(s)
}

func (s L7DdosEventsPaths) GoString() string {
  return s.String()
}

type L7DdosEventsParameters struct {
}

func (s L7DdosEventsParameters) String() string {
  return tea.Prettify(s)
}

func (s L7DdosEventsParameters) GoString() string {
  return s.String()
}

type L7DdosEventsRequestHeader struct {
  // {"zh_CN":"安全服务类型。有使用多个不同的安全服务时，需要填写具体的服务类型。","en":"Security service type. Please enter a specific service type, if you purchase multiple security services.","dictionary":"belong=WAAP-MS-Ext|dict=waap_serviceType"}
  ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty"`
}

func (s L7DdosEventsRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s L7DdosEventsRequestHeader) GoString() string {
  return s.String()
}

func (s *L7DdosEventsRequestHeader) SetServiceType(v string) *L7DdosEventsRequestHeader {
  s.ServiceType = &v
  return s
}

type L7DdosEventsResponseHeader struct {
}

func (s L7DdosEventsResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s L7DdosEventsResponseHeader) GoString() string {
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
  // {'en':'Top rankings, default value: 10, max: 1000.', 'zh_CN':'取前几排名，默认：10，上限：1000。'}
  Top *int32 `json:"top,omitempty" xml:"top,omitempty"`
  // {'en':'Multiple choice. Disposal result, default value: all results.
  //  mitigated: Mitegaed requests.
  //  monitored: monitored requests.', 'zh_CN':'多选。处置结果，默认：展示所有结果。
  //  mitigated：已抵御请求数。
  //  monitored：观察请求数。'}
  ActType []*string `json:"actType,omitempty" xml:"actType,omitempty" type:"Repeated"`
  // {'en':'Start time, format: yyyy-MM-dd HH:mm:ss.
  // The time range does not exceed 31 days.', 'zh_CN':'起始时间，格式：yyyy-MM-dd HH:mm:ss。
  // 时间范围不超过31天。'}
  StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty" require:"true"`
  // {'en':'End time, format: yyyy-MM-dd HH:mm:ss.
  // The time range does not exceed 31 days.', 'zh_CN':'截止时间，格式：yyyy-MM-dd HH:mm:ss。
  // 时间范围不超过31天。'}
  EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty" require:"true"`
  // {'en':'Hostname list, if not specified, it means all the hostnames of the account.', 'zh_CN':'域名列表，未指定时查询账号下的所有域名。'}
  Domains []*string `json:"domains,omitempty" xml:"domains,omitempty" type:"Repeated"`
}

func (s GetTopAttackSourcesForGlobalRequest) String() string {
  return tea.Prettify(s)
}

func (s GetTopAttackSourcesForGlobalRequest) GoString() string {
  return s.String()
}

func (s *GetTopAttackSourcesForGlobalRequest) SetTop(v int32) *GetTopAttackSourcesForGlobalRequest {
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

type GetTopAttackSourcesForGlobalResponse struct {
  // {'en':'Please refer to the error code for exceptions.', 'zh_CN':'请参照错误码。','dictionary':'belong=WAAP-MS-Ext|dict=waap_retCodeEnum'}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {'en':'Description.', 'zh_CN':'描述信息。'}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {'en':'Data.', 'zh_CN':'出参数据。'}
  Data []*GetTopAttackSourcesForGlobalAreaTopDTO `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
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

func (s *GetTopAttackSourcesForGlobalResponse) SetData(v []*GetTopAttackSourcesForGlobalAreaTopDTO) *GetTopAttackSourcesForGlobalResponse {
  s.Data = v
  return s
}

type GetTopAttackSourcesForGlobalAreaTopDTO struct {
  // {'en':'Source country or area.', 'zh_CN':'来源国家或地区。'}
  Area *string `json:"area,omitempty" xml:"area,omitempty" require:"true"`
  // {'en':'Attack requests.', 'zh_CN':'攻击请求数。'}
  Value *int64 `json:"value,omitempty" xml:"value,omitempty" require:"true"`
}

func (s GetTopAttackSourcesForGlobalAreaTopDTO) String() string {
  return tea.Prettify(s)
}

func (s GetTopAttackSourcesForGlobalAreaTopDTO) GoString() string {
  return s.String()
}

func (s *GetTopAttackSourcesForGlobalAreaTopDTO) SetArea(v string) *GetTopAttackSourcesForGlobalAreaTopDTO {
  s.Area = &v
  return s
}

func (s *GetTopAttackSourcesForGlobalAreaTopDTO) SetValue(v int64) *GetTopAttackSourcesForGlobalAreaTopDTO {
  s.Value = &v
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

type GetTopAttackSourcesForGlobalRequestHeader struct {
  // {'en':'The language of response data, default value: en.
  // zh_CN: Chinese
  // en: English', 'zh_CN':'返回内容的语言版本，默认值: en。
  // zh_CN：中文
  // en：英文'}
  Language *string `json:"language,omitempty" xml:"language,omitempty"`
  // {"zh_CN":"安全服务类型。有使用多个不同的安全服务时，需要填写具体的服务类型。","en":"Security service type. Please enter a specific service type, if you purchase multiple security services.","dictionary":"belong=WAAP-MS-Ext|dict=waap_serviceType"}
  ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty"`
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
  // {'en':'Start time, format: yyyy-MM-dd HH:mm:ss.
  // The time range does not exceed 31 days.', 'zh_CN':'开始时间，格式：yyyy-MM-dd HH:mm:ss。
  // 时间范围不超过31天。'}
  StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty" require:"true"`
  // {'en':'End time, format: yyyy-MM-dd HH:mm:ss.
  // The time range does not exceed 31 days.', 'zh_CN':'结束时间，格式：yyyy-MM-dd HH:mm:ss。
  // 时间范围不超过31天。'}
  EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty" require:"true"`
  // {'en':'Multiple choice. Disposal result, default value: all results.
  //  mitigated: Mitegaed requests.
  //  monitored: monitored requests.', 'zh_CN':'多选。处置结果，默认：展示所有结果。
  //  mitigated：已抵御请求数。
  //  monitored：观察请求数。'}
  ActType []*string `json:"actType,omitempty" xml:"actType,omitempty" type:"Repeated"`
  // {'en':'Hostname list, if not specified, it means all the hostnames of the account.', 'zh_CN':'域名列表，未指定时查询账号下的所有域名。'}
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

type GetTrendsByrpsResponse struct {
  // {'en':'Please refer to the error code for exceptions.', 'zh_CN':'请参照错误码。','dictionary':'belong=WAAP-MS-Ext|dict=waap_retCodeEnum'}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {'en':'Description.', 'zh_CN':'描述信息。'}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {'en':'Data.', 'zh_CN':'出参数据。'}
  Data []*GetTrendsByrpsEventTrendDTO `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
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

func (s *GetTrendsByrpsResponse) SetData(v []*GetTrendsByrpsEventTrendDTO) *GetTrendsByrpsResponse {
  s.Data = v
  return s
}

type GetTrendsByrpsEventTypeDTO struct {
  // {'en':'Attack type.
  //  BLOCK: IP/Geo Block
  //  DMS_DEFEND: DDoS Protection
  //  WAF_DEFEND: WAF
  //  BOT_MANAGE: Bot Management
  //  API_DEFEND: API Security
  //  INTELLIGENCE: Threat Intelligence
  //  RATE_LIMIT: Rate Limiting
  //  CUSTOMIZE_RULE: Custom Rules', 'zh_CN':'攻击类型。
  //  BLOCK：IP/区域封禁
  //  DMS_DEFEND：DDoS防护
  //  WAF_DEFEND：WAF
  //  BOT_MANAGE：Bot管理
  //  API_DEFEND：API安全
  //  INTELLIGENCE：威胁情报
  //  RATE_LIMIT：频率限制
  //  CUSTOMIZE_RULE：自定义规则'}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {'en':'Number of attack requests of this type(per second).', 'zh_CN':'该类型攻击请求数（每秒均值）。'}
  Value *int64 `json:"value,omitempty" xml:"value,omitempty" require:"true"`
}

func (s GetTrendsByrpsEventTypeDTO) String() string {
  return tea.Prettify(s)
}

func (s GetTrendsByrpsEventTypeDTO) GoString() string {
  return s.String()
}

func (s *GetTrendsByrpsEventTypeDTO) SetCode(v string) *GetTrendsByrpsEventTypeDTO {
  s.Code = &v
  return s
}

func (s *GetTrendsByrpsEventTypeDTO) SetValue(v int64) *GetTrendsByrpsEventTypeDTO {
  s.Value = &v
  return s
}

type GetTrendsByrpsEventTrendDTO struct {
  // {'en':'Time point(yyyy-MM-dd HH-mm-ss).', 'zh_CN':'时间点（yyyy-MM-dd HH-mm-ss）。'}
  TimePoint *string `json:"timePoint,omitempty" xml:"timePoint,omitempty" require:"true"`
  // {'en':'Total requests(per second).', 'zh_CN':'总请求数（每秒均值）。'}
  Total *int64 `json:"total,omitempty" xml:"total,omitempty" require:"true"`
  // {'en':'Attack requests(per second).', 'zh_CN':'攻击请求数（每秒均值）。'}
  Attack *int64 `json:"attack,omitempty" xml:"attack,omitempty" require:"true"`
  // {'en':'Mitigated requests.', 'zh_CN':'已抵御请求数。'}
  Mitigated *int64 `json:"mitigated,omitempty" xml:"mitigated,omitempty" require:"true"`
  // {'en':'Monitored requests.', 'zh_CN':'观察请求数。'}
  Monitored *int64 `json:"monitored,omitempty" xml:"monitored,omitempty" require:"true"`
  // {'en':'Whitelist requests(per second).', 'zh_CN':'白名单请求数（每秒均值）。'}
  Whitelist *int64 `json:"whitelist,omitempty" xml:"whitelist,omitempty" require:"true"`
  // {'en':'Attack type classification requests(per second).', 'zh_CN':'攻击类型分类请求数（每秒均值）。'}
  Distribution []*GetTrendsByrpsEventTypeDTO `json:"distribution,omitempty" xml:"distribution,omitempty" require:"true" type:"Repeated"`
}

func (s GetTrendsByrpsEventTrendDTO) String() string {
  return tea.Prettify(s)
}

func (s GetTrendsByrpsEventTrendDTO) GoString() string {
  return s.String()
}

func (s *GetTrendsByrpsEventTrendDTO) SetTimePoint(v string) *GetTrendsByrpsEventTrendDTO {
  s.TimePoint = &v
  return s
}

func (s *GetTrendsByrpsEventTrendDTO) SetTotal(v int64) *GetTrendsByrpsEventTrendDTO {
  s.Total = &v
  return s
}

func (s *GetTrendsByrpsEventTrendDTO) SetAttack(v int64) *GetTrendsByrpsEventTrendDTO {
  s.Attack = &v
  return s
}

func (s *GetTrendsByrpsEventTrendDTO) SetMitigated(v int64) *GetTrendsByrpsEventTrendDTO {
  s.Mitigated = &v
  return s
}

func (s *GetTrendsByrpsEventTrendDTO) SetMonitored(v int64) *GetTrendsByrpsEventTrendDTO {
  s.Monitored = &v
  return s
}

func (s *GetTrendsByrpsEventTrendDTO) SetWhitelist(v int64) *GetTrendsByrpsEventTrendDTO {
  s.Whitelist = &v
  return s
}

func (s *GetTrendsByrpsEventTrendDTO) SetDistribution(v []*GetTrendsByrpsEventTypeDTO) *GetTrendsByrpsEventTrendDTO {
  s.Distribution = v
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

type GetTrendsByrpsRequestHeader struct {
  // {"zh_CN":"安全服务类型。有使用多个不同的安全服务时，需要填写具体的服务类型。","en":"Security service type. Please enter a specific service type, if you purchase multiple security services.","dictionary":"belong=WAAP-MS-Ext|dict=waap_serviceType"}
  ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty"`
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




type GetL7DdosAnalysisAttackIpListV2Request struct {
  // {"en":"Start time, format: yyyy-MM-dd HH:mm:ss.
  // The time range does not exceed 30 days.", "zh_CN":"开始时间，格式：yyyy-MM-dd HH:mm:ss。
  // 时间范围不超过30天。"}
  StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty" require:"true"`
  // {"en":"End time, format: yyyy-MM-dd HH:mm:ss.
  // The time range does not exceed 30 days)", "zh_CN":"结束时间，格式：yyyy-MM-dd HH:mm:ss。
  // 时间范围不超过30天。"}
  EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty" require:"true"`
  // {"en":"domains list.", "zh_CN":"域名数组。"}
  Domains []*string `json:"domains,omitempty" xml:"domains,omitempty" require:"true" type:"Repeated"`
  // {"en":"Top num.", "zh_CN":"默认10条。"}
  TopNum *int `json:"topNum,omitempty" xml:"topNum,omitempty"`
}

func (s GetL7DdosAnalysisAttackIpListV2Request) String() string {
  return tea.Prettify(s)
}

func (s GetL7DdosAnalysisAttackIpListV2Request) GoString() string {
  return s.String()
}

func (s *GetL7DdosAnalysisAttackIpListV2Request) SetStartTime(v string) *GetL7DdosAnalysisAttackIpListV2Request {
  s.StartTime = &v
  return s
}

func (s *GetL7DdosAnalysisAttackIpListV2Request) SetEndTime(v string) *GetL7DdosAnalysisAttackIpListV2Request {
  s.EndTime = &v
  return s
}

func (s *GetL7DdosAnalysisAttackIpListV2Request) SetDomains(v []*string) *GetL7DdosAnalysisAttackIpListV2Request {
  s.Domains = v
  return s
}

func (s *GetL7DdosAnalysisAttackIpListV2Request) SetTopNum(v int) *GetL7DdosAnalysisAttackIpListV2Request {
  s.TopNum = &v
  return s
}

type GetL7DdosAnalysisAttackIpListV2Response struct {
  // {'en':'Please refer to the error code for exceptions.', 'zh_CN':'请参照错误码。','dictionary':'belong=WAAP-MS-Ext|dict=waap_retCodeEnum'}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Description.", "zh_CN":"描述信息。"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"Data.", "zh_CN":"出参数据。"}
  Data []*GetL7DdosAnalysisAttackIpListV2AttackIp `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s GetL7DdosAnalysisAttackIpListV2Response) String() string {
  return tea.Prettify(s)
}

func (s GetL7DdosAnalysisAttackIpListV2Response) GoString() string {
  return s.String()
}

func (s *GetL7DdosAnalysisAttackIpListV2Response) SetCode(v string) *GetL7DdosAnalysisAttackIpListV2Response {
  s.Code = &v
  return s
}

func (s *GetL7DdosAnalysisAttackIpListV2Response) SetMsg(v string) *GetL7DdosAnalysisAttackIpListV2Response {
  s.Msg = &v
  return s
}

func (s *GetL7DdosAnalysisAttackIpListV2Response) SetData(v []*GetL7DdosAnalysisAttackIpListV2AttackIp) *GetL7DdosAnalysisAttackIpListV2Response {
  s.Data = v
  return s
}

type GetL7DdosAnalysisAttackIpListV2AttackIp struct {
  // {"en":"ip.", "zh_CN":"攻击ip。"}
  Ip *string `json:"ip,omitempty" xml:"ip,omitempty" require:"true"`
  // {"en":"All count.", "zh_CN":"总请求数。"}
  All_count *int64 `json:"all_count,omitempty" xml:"all_count,omitempty" require:"true"`
  // {"en":"Observed Requests.", "zh_CN":"观察请求数。"}
  AlarmCount *int64 `json:"alarmCount,omitempty" xml:"alarmCount,omitempty" require:"true"`
  // {"en":"Observed Ratio.", "zh_CN":"观察请求占比。"}
  AlarmRatio *float64 `json:"alarmRatio,omitempty" xml:"alarmRatio,omitempty" require:"true"`
  // {"en":"Resisted Requests.", "zh_CN":"已抵御请求数。"}
  Count *int64 `json:"count,omitempty" xml:"count,omitempty" require:"true"`
  // {"en":"Resisted Ratio.", "zh_CN":"已抵御请求占比。"}
  Ratio *float64 `json:"ratio,omitempty" xml:"ratio,omitempty" require:"true"`
  // {"en":"Country - Chinese.", "zh_CN":"国家——中文。"}
  Ip_country_cn *string `json:"ip_country_cn,omitempty" xml:"ip_country_cn,omitempty" require:"true"`
  // {"en":"Province -  Chinese.", "zh_CN":"省份——中文。"}
  Ip_province_cn *string `json:"ip_province_cn,omitempty" xml:"ip_province_cn,omitempty" require:"true"`
  // {"en":"city - English.", "zh_CN":"城市——英文。"}
  Ip_city_en *string `json:"ip_city_en,omitempty" xml:"ip_city_en,omitempty" require:"true"`
  // {"en":"Country - English.", "zh_CN":"国家——英文。"}
  Ip_country_en *string `json:"ip_country_en,omitempty" xml:"ip_country_en,omitempty" require:"true"`
  // {"en":"Province - English.", "zh_CN":"省份——英文。"}
  Ip_province_en *string `json:"ip_province_en,omitempty" xml:"ip_province_en,omitempty" require:"true"`
}

func (s GetL7DdosAnalysisAttackIpListV2AttackIp) String() string {
  return tea.Prettify(s)
}

func (s GetL7DdosAnalysisAttackIpListV2AttackIp) GoString() string {
  return s.String()
}

func (s *GetL7DdosAnalysisAttackIpListV2AttackIp) SetIp(v string) *GetL7DdosAnalysisAttackIpListV2AttackIp {
  s.Ip = &v
  return s
}

func (s *GetL7DdosAnalysisAttackIpListV2AttackIp) SetAll_count(v int64) *GetL7DdosAnalysisAttackIpListV2AttackIp {
  s.All_count = &v
  return s
}

func (s *GetL7DdosAnalysisAttackIpListV2AttackIp) SetAlarmCount(v int64) *GetL7DdosAnalysisAttackIpListV2AttackIp {
  s.AlarmCount = &v
  return s
}

func (s *GetL7DdosAnalysisAttackIpListV2AttackIp) SetAlarmRatio(v float64) *GetL7DdosAnalysisAttackIpListV2AttackIp {
  s.AlarmRatio = &v
  return s
}

func (s *GetL7DdosAnalysisAttackIpListV2AttackIp) SetCount(v int64) *GetL7DdosAnalysisAttackIpListV2AttackIp {
  s.Count = &v
  return s
}

func (s *GetL7DdosAnalysisAttackIpListV2AttackIp) SetRatio(v float64) *GetL7DdosAnalysisAttackIpListV2AttackIp {
  s.Ratio = &v
  return s
}

func (s *GetL7DdosAnalysisAttackIpListV2AttackIp) SetIp_country_cn(v string) *GetL7DdosAnalysisAttackIpListV2AttackIp {
  s.Ip_country_cn = &v
  return s
}

func (s *GetL7DdosAnalysisAttackIpListV2AttackIp) SetIp_province_cn(v string) *GetL7DdosAnalysisAttackIpListV2AttackIp {
  s.Ip_province_cn = &v
  return s
}

func (s *GetL7DdosAnalysisAttackIpListV2AttackIp) SetIp_city_en(v string) *GetL7DdosAnalysisAttackIpListV2AttackIp {
  s.Ip_city_en = &v
  return s
}

func (s *GetL7DdosAnalysisAttackIpListV2AttackIp) SetIp_country_en(v string) *GetL7DdosAnalysisAttackIpListV2AttackIp {
  s.Ip_country_en = &v
  return s
}

func (s *GetL7DdosAnalysisAttackIpListV2AttackIp) SetIp_province_en(v string) *GetL7DdosAnalysisAttackIpListV2AttackIp {
  s.Ip_province_en = &v
  return s
}

type GetL7DdosAnalysisAttackIpListV2Paths struct {
}

func (s GetL7DdosAnalysisAttackIpListV2Paths) String() string {
  return tea.Prettify(s)
}

func (s GetL7DdosAnalysisAttackIpListV2Paths) GoString() string {
  return s.String()
}

type GetL7DdosAnalysisAttackIpListV2Parameters struct {
}

func (s GetL7DdosAnalysisAttackIpListV2Parameters) String() string {
  return tea.Prettify(s)
}

func (s GetL7DdosAnalysisAttackIpListV2Parameters) GoString() string {
  return s.String()
}

type GetL7DdosAnalysisAttackIpListV2RequestHeader struct {
  // {"zh_CN":"安全服务类型。有使用多个不同的安全服务时，需要填写具体的服务类型。","en":"Security service type. Please enter a specific service type, if you purchase multiple security services.","dictionary":"belong=WAAP-MS-Ext|dict=waap_serviceType"}
  ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty"`
}

func (s GetL7DdosAnalysisAttackIpListV2RequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetL7DdosAnalysisAttackIpListV2RequestHeader) GoString() string {
  return s.String()
}

func (s *GetL7DdosAnalysisAttackIpListV2RequestHeader) SetServiceType(v string) *GetL7DdosAnalysisAttackIpListV2RequestHeader {
  s.ServiceType = &v
  return s
}

type GetL7DdosAnalysisAttackIpListV2ResponseHeader struct {
}

func (s GetL7DdosAnalysisAttackIpListV2ResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetL7DdosAnalysisAttackIpListV2ResponseHeader) GoString() string {
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
  // {'en':'Top rankings, default value: 10, max: 1000.', 'zh_CN':'取前几排名，默认：10，上限：1000。'}
  Top *int32 `json:"top,omitempty" xml:"top,omitempty"`
  // {'en':'Multiple choice. Disposal result, default value: all results.
  //  mitigated: Mitegaed requests.
  //  monitored: monitored requests.', 'zh_CN':'多选。处置结果，默认：展示所有结果。
  //  mitigated：已抵御请求数。
  //  monitored：观察请求数。'}
  ActType []*string `json:"actType,omitempty" xml:"actType,omitempty" type:"Repeated"`
  // {'en':'Start time, format: yyyy-MM-dd HH:mm:ss.
  // The time range does not exceed 31 days.', 'zh_CN':'起始时间，格式：yyyy-MM-dd HH:mm:ss。
  // 时间范围不超过31天。'}
  StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty" require:"true"`
  // {'en':'End time, format: yyyy-MM-dd HH:mm:ss.
  // The time range does not exceed 31 days.', 'zh_CN':'截止时间，格式：yyyy-MM-dd HH:mm:ss。
  // 时间范围不超过31天。'}
  EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty" require:"true"`
  // {'en':'Hostname list, if not specified, it means all the hostnames of the account.', 'zh_CN':'域名列表，未指定时查询账号下的所有域名。'}
  Domains []*string `json:"domains,omitempty" xml:"domains,omitempty" type:"Repeated"`
}

func (s GetTriggeredDDoSManagedRulesRequest) String() string {
  return tea.Prettify(s)
}

func (s GetTriggeredDDoSManagedRulesRequest) GoString() string {
  return s.String()
}

func (s *GetTriggeredDDoSManagedRulesRequest) SetTop(v int32) *GetTriggeredDDoSManagedRulesRequest {
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

type GetTriggeredDDoSManagedRulesResponse struct {
  // {'en':'Please refer to the error code for exceptions.', 'zh_CN':'请参照错误码。','dictionary':'belong=WAAP-MS-Ext|dict=waap_retCodeEnum'}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {'en':'Description.', 'zh_CN':'描述信息。'}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {'en':'Data.', 'zh_CN':'出参数据。'}
  Data []*GetTriggeredDDoSManagedRulesRuleTopDTO `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
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

func (s *GetTriggeredDDoSManagedRulesResponse) SetData(v []*GetTriggeredDDoSManagedRulesRuleTopDTO) *GetTriggeredDDoSManagedRulesResponse {
  s.Data = v
  return s
}

type GetTriggeredDDoSManagedRulesRuleHitDTO struct {
  // {'en':'Action.', 'zh_CN':'采取动作。'}
  Act *string `json:"act,omitempty" xml:"act,omitempty" require:"true"`
  // {'en':'Hit times.', 'zh_CN':'命中数。'}
  Value *int64 `json:"value,omitempty" xml:"value,omitempty" require:"true"`
}

func (s GetTriggeredDDoSManagedRulesRuleHitDTO) String() string {
  return tea.Prettify(s)
}

func (s GetTriggeredDDoSManagedRulesRuleHitDTO) GoString() string {
  return s.String()
}

func (s *GetTriggeredDDoSManagedRulesRuleHitDTO) SetAct(v string) *GetTriggeredDDoSManagedRulesRuleHitDTO {
  s.Act = &v
  return s
}

func (s *GetTriggeredDDoSManagedRulesRuleHitDTO) SetValue(v int64) *GetTriggeredDDoSManagedRulesRuleHitDTO {
  s.Value = &v
  return s
}

type GetTriggeredDDoSManagedRulesRuleTopDTO struct {
  // {'en':'Rule ID.', 'zh_CN':'规则ID。'}
  RuleId *string `json:"ruleId,omitempty" xml:"ruleId,omitempty" require:"true"`
  // {'en':'Rule name.', 'zh_CN':'规则名称。'}
  RuleName *string `json:"ruleName,omitempty" xml:"ruleName,omitempty" require:"true"`
  // {'en':'Rule type.', 'zh_CN':'规则类型。'}
  RuleType *string `json:"ruleType,omitempty" xml:"ruleType,omitempty" require:"true"`
  // {'en':'System recommended action.', 'zh_CN':'系统推荐动作。'}
  Action *string `json:"action,omitempty" xml:"action,omitempty" require:"true"`
  // {'en':'Trigger times, sort by times.', 'zh_CN':'触发次数，按次数排序。'}
  Hits []*GetTriggeredDDoSManagedRulesRuleHitDTO `json:"hits,omitempty" xml:"hits,omitempty" require:"true" type:"Repeated"`
}

func (s GetTriggeredDDoSManagedRulesRuleTopDTO) String() string {
  return tea.Prettify(s)
}

func (s GetTriggeredDDoSManagedRulesRuleTopDTO) GoString() string {
  return s.String()
}

func (s *GetTriggeredDDoSManagedRulesRuleTopDTO) SetRuleId(v string) *GetTriggeredDDoSManagedRulesRuleTopDTO {
  s.RuleId = &v
  return s
}

func (s *GetTriggeredDDoSManagedRulesRuleTopDTO) SetRuleName(v string) *GetTriggeredDDoSManagedRulesRuleTopDTO {
  s.RuleName = &v
  return s
}

func (s *GetTriggeredDDoSManagedRulesRuleTopDTO) SetRuleType(v string) *GetTriggeredDDoSManagedRulesRuleTopDTO {
  s.RuleType = &v
  return s
}

func (s *GetTriggeredDDoSManagedRulesRuleTopDTO) SetAction(v string) *GetTriggeredDDoSManagedRulesRuleTopDTO {
  s.Action = &v
  return s
}

func (s *GetTriggeredDDoSManagedRulesRuleTopDTO) SetHits(v []*GetTriggeredDDoSManagedRulesRuleHitDTO) *GetTriggeredDDoSManagedRulesRuleTopDTO {
  s.Hits = v
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

type GetTriggeredDDoSManagedRulesRequestHeader struct {
  // {'en':'The language of response data, default value: en.
  // zh_CN: Chinese
  // en: English', 'zh_CN':'返回内容的语言版本，默认值: en。
  // zh_CN：中文
  // en：英文'}
  Language *string `json:"language,omitempty" xml:"language,omitempty"`
  // {"zh_CN":"安全服务类型。有使用多个不同的安全服务时，需要填写具体的服务类型。","en":"Security service type. Please enter a specific service type, if you purchase multiple security services.","dictionary":"belong=WAAP-MS-Ext|dict=waap_serviceType"}
  ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty"`
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

type GetTriggeredDDoSManagedRulesResponseHeader struct {
}

func (s GetTriggeredDDoSManagedRulesResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetTriggeredDDoSManagedRulesResponseHeader) GoString() string {
  return s.String()
}




type GetTopAttackSourcesForChinaMainlandRequest struct {
  // {'en':'Top rankings, default value: 10, max: 1000.', 'zh_CN':'取前几排名，默认：10，上限：1000。'}
  Top *int32 `json:"top,omitempty" xml:"top,omitempty"`
  // {'en':'Multiple choice. Disposal result, default value: all results.
  //  mitigated: Mitegaed requests.
  //  monitored: monitored requests.', 'zh_CN':'多选。处置结果，默认：展示所有结果。
  //  mitigated：已抵御请求数。
  //  monitored：观察请求数。'}
  ActType []*string `json:"actType,omitempty" xml:"actType,omitempty" type:"Repeated"`
  // {'en':'Start time, format: yyyy-MM-dd HH:mm:ss.
  // The time range does not exceed 31 days.', 'zh_CN':'起始时间，格式：yyyy-MM-dd HH:mm:ss。
  // 时间范围不超过31天。'}
  StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty" require:"true"`
  // {'en':'End time, format: yyyy-MM-dd HH:mm:ss.
  // The time range does not exceed 31 days.', 'zh_CN':'截止时间，格式：yyyy-MM-dd HH:mm:ss。
  // 时间范围不超过31天。'}
  EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty" require:"true"`
  // {'en':'Hostname list, if not specified, it means all the hostnames of the account.', 'zh_CN':'域名列表，未指定时查询账号下的所有域名。'}
  Domains []*string `json:"domains,omitempty" xml:"domains,omitempty" type:"Repeated"`
}

func (s GetTopAttackSourcesForChinaMainlandRequest) String() string {
  return tea.Prettify(s)
}

func (s GetTopAttackSourcesForChinaMainlandRequest) GoString() string {
  return s.String()
}

func (s *GetTopAttackSourcesForChinaMainlandRequest) SetTop(v int32) *GetTopAttackSourcesForChinaMainlandRequest {
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

type GetTopAttackSourcesForChinaMainlandResponse struct {
  // {'en':'Please refer to the error code for exceptions.', 'zh_CN':'请参照错误码。','dictionary':'belong=WAAP-MS-Ext|dict=waap_retCodeEnum'}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {'en':'Description.', 'zh_CN':'描述信息。'}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {'en':'Data.', 'zh_CN':'出参数据。'}
  Data []*GetTopAttackSourcesForChinaMainlandAreaTopDTO `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
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

func (s *GetTopAttackSourcesForChinaMainlandResponse) SetData(v []*GetTopAttackSourcesForChinaMainlandAreaTopDTO) *GetTopAttackSourcesForChinaMainlandResponse {
  s.Data = v
  return s
}

type GetTopAttackSourcesForChinaMainlandAreaTopDTO struct {
  // {'en':'Source area(domestic province).', 'zh_CN':'来源地区（国内省份）。'}
  Area *string `json:"area,omitempty" xml:"area,omitempty" require:"true"`
  // {'en':'Source city.', 'zh_CN':'来源城市。'}
  City *string `json:"city,omitempty" xml:"city,omitempty" require:"true"`
  // {'en':'Attack requests.', 'zh_CN':'攻击请求数。'}
  Value *int64 `json:"value,omitempty" xml:"value,omitempty" require:"true"`
}

func (s GetTopAttackSourcesForChinaMainlandAreaTopDTO) String() string {
  return tea.Prettify(s)
}

func (s GetTopAttackSourcesForChinaMainlandAreaTopDTO) GoString() string {
  return s.String()
}

func (s *GetTopAttackSourcesForChinaMainlandAreaTopDTO) SetArea(v string) *GetTopAttackSourcesForChinaMainlandAreaTopDTO {
  s.Area = &v
  return s
}

func (s *GetTopAttackSourcesForChinaMainlandAreaTopDTO) SetCity(v string) *GetTopAttackSourcesForChinaMainlandAreaTopDTO {
  s.City = &v
  return s
}

func (s *GetTopAttackSourcesForChinaMainlandAreaTopDTO) SetValue(v int64) *GetTopAttackSourcesForChinaMainlandAreaTopDTO {
  s.Value = &v
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

type GetTopAttackSourcesForChinaMainlandRequestHeader struct {
  // {'en':'The language of response data, default value: en.
  // zh_CN: Chinese
  // en: English', 'zh_CN':'返回内容的语言版本，默认值: en。
  // zh_CN：中文
  // en：英文'}
  Language *string `json:"language,omitempty" xml:"language,omitempty"`
  // {"zh_CN":"安全服务类型。有使用多个不同的安全服务时，需要填写具体的服务类型。","en":"Security service type. Please enter a specific service type, if you purchase multiple security services.","dictionary":"belong=WAAP-MS-Ext|dict=waap_serviceType"}
  ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty"`
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




type DmsThreatenAnalysisAttackedUrlsRequest struct {
  // {"en":"start time", "zh_CN":"开始时间"}
  StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty" require:"true"`
  // {"en":"end time", "zh_CN":"结束时间"}
  EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty" require:"true"`
  // {"en":"hostname", "zh_CN":"域名，多个用;分隔"}
  Hostname *string `json:"hostname,omitempty" xml:"hostname,omitempty" require:"true"`
  // {"en":"top num", "zh_CN":"默认10条"}
  TopNum *int `json:"topNum,omitempty" xml:"topNum,omitempty"`
}

func (s DmsThreatenAnalysisAttackedUrlsRequest) String() string {
  return tea.Prettify(s)
}

func (s DmsThreatenAnalysisAttackedUrlsRequest) GoString() string {
  return s.String()
}

func (s *DmsThreatenAnalysisAttackedUrlsRequest) SetStartTime(v string) *DmsThreatenAnalysisAttackedUrlsRequest {
  s.StartTime = &v
  return s
}

func (s *DmsThreatenAnalysisAttackedUrlsRequest) SetEndTime(v string) *DmsThreatenAnalysisAttackedUrlsRequest {
  s.EndTime = &v
  return s
}

func (s *DmsThreatenAnalysisAttackedUrlsRequest) SetHostname(v string) *DmsThreatenAnalysisAttackedUrlsRequest {
  s.Hostname = &v
  return s
}

func (s *DmsThreatenAnalysisAttackedUrlsRequest) SetTopNum(v int) *DmsThreatenAnalysisAttackedUrlsRequest {
  s.TopNum = &v
  return s
}

type DmsThreatenAnalysisAttackedUrlsResponse struct {
  // {'en':'Please refer to the error code for exceptions.', 'zh_CN':'请参照错误码。','dictionary':'belong=WAAP-MS-Ext|dict=waap_retCodeEnum'}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Description.", "zh_CN":"描述信息。"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"Data.", "zh_CN":"出参数据。"}
  Data map[string]interface{} `json:"data,omitempty" xml:"data,omitempty" require:"true"`
  // {"en":"url", "zh_CN":"url"}
  Url *string `json:"url,omitempty" xml:"url,omitempty" require:"true"`
  // {"en":"alarm count", "zh_CN":"告警数"}
  AlarmCount *int64 `json:"alarmCount,omitempty" xml:"alarmCount,omitempty" require:"true"`
  // {"en":"all count", "zh_CN":"总数"}
  Count *int64 `json:"count,omitempty" xml:"count,omitempty" require:"true"`
}

func (s DmsThreatenAnalysisAttackedUrlsResponse) String() string {
  return tea.Prettify(s)
}

func (s DmsThreatenAnalysisAttackedUrlsResponse) GoString() string {
  return s.String()
}

func (s *DmsThreatenAnalysisAttackedUrlsResponse) SetCode(v string) *DmsThreatenAnalysisAttackedUrlsResponse {
  s.Code = &v
  return s
}

func (s *DmsThreatenAnalysisAttackedUrlsResponse) SetMsg(v string) *DmsThreatenAnalysisAttackedUrlsResponse {
  s.Msg = &v
  return s
}

func (s *DmsThreatenAnalysisAttackedUrlsResponse) SetData(v map[string]interface{}) *DmsThreatenAnalysisAttackedUrlsResponse {
  s.Data = v
  return s
}

func (s *DmsThreatenAnalysisAttackedUrlsResponse) SetUrl(v string) *DmsThreatenAnalysisAttackedUrlsResponse {
  s.Url = &v
  return s
}

func (s *DmsThreatenAnalysisAttackedUrlsResponse) SetAlarmCount(v int64) *DmsThreatenAnalysisAttackedUrlsResponse {
  s.AlarmCount = &v
  return s
}

func (s *DmsThreatenAnalysisAttackedUrlsResponse) SetCount(v int64) *DmsThreatenAnalysisAttackedUrlsResponse {
  s.Count = &v
  return s
}

type DmsThreatenAnalysisAttackedUrlsPaths struct {
}

func (s DmsThreatenAnalysisAttackedUrlsPaths) String() string {
  return tea.Prettify(s)
}

func (s DmsThreatenAnalysisAttackedUrlsPaths) GoString() string {
  return s.String()
}

type DmsThreatenAnalysisAttackedUrlsParameters struct {
}

func (s DmsThreatenAnalysisAttackedUrlsParameters) String() string {
  return tea.Prettify(s)
}

func (s DmsThreatenAnalysisAttackedUrlsParameters) GoString() string {
  return s.String()
}

type DmsThreatenAnalysisAttackedUrlsRequestHeader struct {
  // {"zh_CN":"安全服务类型。有使用多个不同的安全服务时，需要填写具体的服务类型。","en":"Security service type. Please enter a specific service type, if you purchase multiple security services.","dictionary":"belong=WAAP-MS-Ext|dict=waap_serviceType"}
  ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty"`
}

func (s DmsThreatenAnalysisAttackedUrlsRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s DmsThreatenAnalysisAttackedUrlsRequestHeader) GoString() string {
  return s.String()
}

func (s *DmsThreatenAnalysisAttackedUrlsRequestHeader) SetServiceType(v string) *DmsThreatenAnalysisAttackedUrlsRequestHeader {
  s.ServiceType = &v
  return s
}

type DmsThreatenAnalysisAttackedUrlsResponseHeader struct {
}

func (s DmsThreatenAnalysisAttackedUrlsResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s DmsThreatenAnalysisAttackedUrlsResponseHeader) GoString() string {
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
  // {'en':'Top rankings, default: 100, max: 1000.', 'zh_CN':'取前几排名，默认100，上限1000。'}
  Top *int32 `json:"top,omitempty" xml:"top,omitempty"`
  // {'en':'Multiple choice. Disposal result, default value: all results.
  //  mitigated: Mitegaed requests.
  //  monitored: monitored requests.', 'zh_CN':'多选。处置结果，默认：展示所有结果。
  //  mitigated：已抵御请求数。
  //  monitored：观察请求数。'}
  ActType []*string `json:"actType,omitempty" xml:"actType,omitempty" type:"Repeated"`
  // {'en':'Start time, format: yyyy-MM-dd HH:mm:ss.
  // The time range does not exceed 31 days.', 'zh_CN':'起始时间，格式：yyyy-MM-dd HH:mm:ss。
  // 时间范围不超过31天。'}
  StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty" require:"true"`
  // {'en':'End time, format: yyyy-MM-dd HH:mm:ss.
  // The time range does not exceed 31 days.', 'zh_CN':'截止时间，格式：yyyy-MM-dd HH:mm:ss。
  // 时间范围不超过31天。'}
  EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty" require:"true"`
  // {'en':'Hostname list, if not specified, it means all the hostnames of the account.', 'zh_CN':'域名列表，未指定时查询账号下的所有域名。'}
  Domains []*string `json:"domains,omitempty" xml:"domains,omitempty" type:"Repeated"`
}

func (s GetTopAttackTargetsByPathRequest) String() string {
  return tea.Prettify(s)
}

func (s GetTopAttackTargetsByPathRequest) GoString() string {
  return s.String()
}

func (s *GetTopAttackTargetsByPathRequest) SetTop(v int32) *GetTopAttackTargetsByPathRequest {
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

type GetTopAttackTargetsByPathResponse struct {
  // {'en':'Please refer to the error code for exceptions.', 'zh_CN':'请参照错误码。','dictionary':'belong=WAAP-MS-Ext|dict=waap_retCodeEnum'}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {'en':'Description.', 'zh_CN':'描述信息。'}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {'en':'Data.', 'zh_CN':'出参数据。'}
  Data []*GetTopAttackTargetsByPathUrlTopDTO `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
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

func (s *GetTopAttackTargetsByPathResponse) SetData(v []*GetTopAttackTargetsByPathUrlTopDTO) *GetTopAttackTargetsByPathResponse {
  s.Data = v
  return s
}

type GetTopAttackTargetsByPathUrlTopDTO struct {
  // {'en':'URL.', 'zh_CN':'URL。'}
  Url *string `json:"url,omitempty" xml:"url,omitempty" require:"true"`
  // {'en':'Attack requests.', 'zh_CN':'攻击请求数。'}
  Attack *int64 `json:"attack,omitempty" xml:"attack,omitempty" require:"true"`
}

func (s GetTopAttackTargetsByPathUrlTopDTO) String() string {
  return tea.Prettify(s)
}

func (s GetTopAttackTargetsByPathUrlTopDTO) GoString() string {
  return s.String()
}

func (s *GetTopAttackTargetsByPathUrlTopDTO) SetUrl(v string) *GetTopAttackTargetsByPathUrlTopDTO {
  s.Url = &v
  return s
}

func (s *GetTopAttackTargetsByPathUrlTopDTO) SetAttack(v int64) *GetTopAttackTargetsByPathUrlTopDTO {
  s.Attack = &v
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

type GetTopAttackTargetsByPathRequestHeader struct {
  // {"zh_CN":"安全服务类型。有使用多个不同的安全服务时，需要填写具体的服务类型。","en":"Security service type. Please enter a specific service type, if you purchase multiple security services.","dictionary":"belong=WAAP-MS-Ext|dict=waap_serviceType"}
  ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty"`
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
  // {'en':'Start time, format: yyyy-MM-dd HH:mm:ss.
  // The time range does not exceed 31 days. ', 'zh_CN':'开始时间，格式：yyyy-MM-dd HH:mm:ss。
  // 时间范围不超过31天。'}
  StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty" require:"true"`
  // {'en':'End time, format: yyyy-MM-dd HH:mm:ss.
  // The time range does not exceed 31 days. ', 'zh_CN':'结束时间，格式：yyyy-MM-dd HH:mm:ss。
  // 时间范围不超过31天。'}
  EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty" require:"true"`
  // {'en':'Hostname list, if  not specified, it means all the hostnames of the account.', 'zh_CN':'域名列表，未指定时查询账号下的所有域名。'}
  DomainList []*string `json:"domainList,omitempty" xml:"domainList,omitempty" type:"Repeated"`
  // {'en':'Policy type list. 
  //  DMS_DEFEND: DDoS Protection
  //  BOT_MANAGE: Bot Management
  //  API_DEFEND: API Security
  //  WAF_DEFEND: WAF
  //  BLOCK: IP/Geo Block
  //  CUSTOMIZE_RULE: Custom Rules
  //  RATE_LIMIT: Rate Limiting
  //  INTELLIGENCE: Threat Intelligence', 'zh_CN':'策略类型列表。
  // DMS_DEFEND：DDoS防护
  //  BOT_MANAGE：Bot管理
  //  API_DEFEND：API安全
  //  WAF_DEFEND：WAF
  //  BLOCK：IP/区域封禁
  //  CUSTOMIZE_RULE：自定义规则
  //  RATE_LIMIT：频率限制
  //  INTELLIGENCE：威胁情报'}
  PolicyTypeList []*string `json:"policyTypeList,omitempty" xml:"policyTypeList,omitempty" type:"Repeated"`
  // {'en':'The number of records per page, default value:10, maximum value 2000.', 'zh_CN':'每页显示的条目数，默认：10，最大：2000。'}
  PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
  // {'en':'The current page, default value: 1.', 'zh_CN':'当前第几页，默认：1。'}
  CurrentPage *int32 `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
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

func (s *ListAttackLogsRequest) SetPageSize(v int32) *ListAttackLogsRequest {
  s.PageSize = &v
  return s
}

func (s *ListAttackLogsRequest) SetCurrentPage(v int32) *ListAttackLogsRequest {
  s.CurrentPage = &v
  return s
}

type ListAttackLogsResponse struct {
  // {'en':'Please refer to the error code for exceptions.', 'zh_CN':'请参照错误码。','dictionary':'belong=WAAP-MS-Ext|dict=waap_retCodeEnum'}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {'en':'Description.', 'zh_CN':'描述信息。'}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {'en':'Data.', 'zh_CN':'出参数据。'}
  Data *ListAttackLogsPageBean `json:"data,omitempty" xml:"data,omitempty" require:"true"`
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

func (s *ListAttackLogsResponse) SetData(v *ListAttackLogsPageBean) *ListAttackLogsResponse {
  s.Data = v
  return s
}

type ListAttackLogsAttackLogSimpleDto struct {
  // {'en':'Attack time, format: yyyy-MM-dd HH:mm:ss.', 'zh_CN':'攻击时间，格式：yyyy-MM-dd HH:mm:ss。'}
  AttackTime *string `json:"attackTime,omitempty" xml:"attackTime,omitempty" require:"true"`
  // {'en':'Client IP location.', 'zh_CN':'客户端IP。'}
  ClientIp *string `json:"clientIp,omitempty" xml:"clientIp,omitempty" require:"true"`
  // {'en':'Hostname.', 'zh_CN':'域名。'}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {'en':'Event ID.', 'zh_CN':'事件ID。'}
  Uuid *string `json:"uuid,omitempty" xml:"uuid,omitempty" require:"true"`
  // {'en':'Requeset ID.', 'zh_CN':'请求ID。'}
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty" require:"true"`
  // {'en':'Policy type. 
  //  DMS_DEFEND: DDoS Protection
  //  BOT_MANAGE: Bot Management
  //  API_DEFEND: API Security
  //  WAF_DEFEND: WAF
  //  BLOCK: IP/Geo Block
  //  CUSTOMIZE_RULE: Custom Rules
  //  RATE_LIMIT: Rate Limiting
  //  INTELLIGENCE: Threat Intelligence', 'zh_CN':'策略类型。
  //  DMS_DEFEND：DDoS防护
  //  BOT_MANAGE：Bot管理
  //  API_DEFEND：API安全
  //  WAF_DEFEND：WAF
  //  BLOCK：IP/区域封禁
  //  CUSTOMIZE_RULE：自定义规则
  //  RATE_LIMIT：频率限制
  //  INTELLIGENCE：威胁情报'}
  PolicyType *string `json:"policyType,omitempty" xml:"policyType,omitempty" require:"true"`
  // {'en':'Path.', 'zh_CN':'路径。'}
  Path *string `json:"path,omitempty" xml:"path,omitempty" require:"true"`
  // {'en':'Status code.', 'zh_CN':'状态码。'}
  StatusCode *string `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
  // {'en':'Action.', 'zh_CN':'处理动作。'}
  Action *string `json:"action,omitempty" xml:"action,omitempty" require:"true"`
}

func (s ListAttackLogsAttackLogSimpleDto) String() string {
  return tea.Prettify(s)
}

func (s ListAttackLogsAttackLogSimpleDto) GoString() string {
  return s.String()
}

func (s *ListAttackLogsAttackLogSimpleDto) SetAttackTime(v string) *ListAttackLogsAttackLogSimpleDto {
  s.AttackTime = &v
  return s
}

func (s *ListAttackLogsAttackLogSimpleDto) SetClientIp(v string) *ListAttackLogsAttackLogSimpleDto {
  s.ClientIp = &v
  return s
}

func (s *ListAttackLogsAttackLogSimpleDto) SetDomain(v string) *ListAttackLogsAttackLogSimpleDto {
  s.Domain = &v
  return s
}

func (s *ListAttackLogsAttackLogSimpleDto) SetUuid(v string) *ListAttackLogsAttackLogSimpleDto {
  s.Uuid = &v
  return s
}

func (s *ListAttackLogsAttackLogSimpleDto) SetRequestId(v string) *ListAttackLogsAttackLogSimpleDto {
  s.RequestId = &v
  return s
}

func (s *ListAttackLogsAttackLogSimpleDto) SetPolicyType(v string) *ListAttackLogsAttackLogSimpleDto {
  s.PolicyType = &v
  return s
}

func (s *ListAttackLogsAttackLogSimpleDto) SetPath(v string) *ListAttackLogsAttackLogSimpleDto {
  s.Path = &v
  return s
}

func (s *ListAttackLogsAttackLogSimpleDto) SetStatusCode(v string) *ListAttackLogsAttackLogSimpleDto {
  s.StatusCode = &v
  return s
}

func (s *ListAttackLogsAttackLogSimpleDto) SetAction(v string) *ListAttackLogsAttackLogSimpleDto {
  s.Action = &v
  return s
}

type ListAttackLogsPageBean struct {
  // {'en':'Return content.', 'zh_CN':'返回内容。'}
  List *ListAttackLogsAttackLogSimpleDto `json:"list,omitempty" xml:"list,omitempty" require:"true"`
  // {'en':'The current page number.', 'zh_CN':'当前页码。'}
  PageNum *int64 `json:"pageNum,omitempty" xml:"pageNum,omitempty" require:"true"`
  // {'en':'The number of records per page.', 'zh_CN':'每页显示的条目数。'}
  PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty" require:"true"`
  // {'en':'The total number of records.', 'zh_CN':'总记录数。'}
  TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty" require:"true"`
  // {'en':'The total number of pages.', 'zh_CN':'总页数。'}
  TotalPageCount *int64 `json:"totalPageCount,omitempty" xml:"totalPageCount,omitempty" require:"true"`
}

func (s ListAttackLogsPageBean) String() string {
  return tea.Prettify(s)
}

func (s ListAttackLogsPageBean) GoString() string {
  return s.String()
}

func (s *ListAttackLogsPageBean) SetList(v *ListAttackLogsAttackLogSimpleDto) *ListAttackLogsPageBean {
  s.List = v
  return s
}

func (s *ListAttackLogsPageBean) SetPageNum(v int64) *ListAttackLogsPageBean {
  s.PageNum = &v
  return s
}

func (s *ListAttackLogsPageBean) SetPageSize(v int64) *ListAttackLogsPageBean {
  s.PageSize = &v
  return s
}

func (s *ListAttackLogsPageBean) SetTotalCount(v int64) *ListAttackLogsPageBean {
  s.TotalCount = &v
  return s
}

func (s *ListAttackLogsPageBean) SetTotalPageCount(v int64) *ListAttackLogsPageBean {
  s.TotalPageCount = &v
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

type ListAttackLogsRequestHeader struct {
  // {'en':'The language of response data, default value: en.
  // zh_CN: Chinese
  // en: English', 'zh_CN':'返回内容的语言版本，默认值: en。
  // zh_CN：中文
  // en：英文'}
  Language *string `json:"language,omitempty" xml:"language,omitempty"`
  // {"zh_CN":"安全服务类型。有使用多个不同的安全服务时，需要填写具体的服务类型。","en":"Security service type. Please enter a specific service type, if you purchase multiple security services.","dictionary":"belong=WAAP-MS-Ext|dict=waap_serviceType"}
  ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty"`
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

type ListAttackLogsResponseHeader struct {
}

func (s ListAttackLogsResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ListAttackLogsResponseHeader) GoString() string {
  return s.String()
}




type GetTopAttackSourcesForClientIPsRequest struct {
  // {'en':'Top rankings, default value: 10, max: 1000.', 'zh_CN':'取前几排名，默认：10，上限：1000。'}
  Top *int32 `json:"top,omitempty" xml:"top,omitempty"`
  // {'en':'Multiple choice. Disposal result, default value: all results.
  //  mitigated: Mitegaed requests.
  //  monitored: monitored requests.', 'zh_CN':'多选。处置结果，默认：展示所有结果。
  //  mitigated：已抵御请求数。
  //  monitored：观察请求数。'}
  ActType []*string `json:"actType,omitempty" xml:"actType,omitempty" type:"Repeated"`
  // {'en':'Start time, format: yyyy-MM-dd HH:mm:ss.
  // The time range does not exceed 31 days.', 'zh_CN':'起始时间，格式：yyyy-MM-dd HH:mm:ss。
  // 时间范围不超过31天。'}
  StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty" require:"true"`
  // {'en':'End time, format: yyyy-MM-dd HH:mm:ss.
  // The time range does not exceed 31 days.', 'zh_CN':'截止时间，格式：yyyy-MM-dd HH:mm:ss。
  // 时间范围不超过31天。'}
  EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty" require:"true"`
  // {'en':'Hostname list, if not specified, it means all the hostnames of the account.', 'zh_CN':'域名列表，未指定时查询账号下的所有域名。'}
  Domains []*string `json:"domains,omitempty" xml:"domains,omitempty" type:"Repeated"`
}

func (s GetTopAttackSourcesForClientIPsRequest) String() string {
  return tea.Prettify(s)
}

func (s GetTopAttackSourcesForClientIPsRequest) GoString() string {
  return s.String()
}

func (s *GetTopAttackSourcesForClientIPsRequest) SetTop(v int32) *GetTopAttackSourcesForClientIPsRequest {
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

type GetTopAttackSourcesForClientIPsResponse struct {
  // {'en':'Please refer to the error code for exceptions.', 'zh_CN':'请参照错误码。','dictionary':'belong=WAAP-MS-Ext|dict=waap_retCodeEnum'}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {'en':'Description.', 'zh_CN':'描述信息。'}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {'en':'Data.', 'zh_CN':'出参数据。'}
  Data []*GetTopAttackSourcesForClientIPsAreaTopDTO `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
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

func (s *GetTopAttackSourcesForClientIPsResponse) SetData(v []*GetTopAttackSourcesForClientIPsAreaTopDTO) *GetTopAttackSourcesForClientIPsResponse {
  s.Data = v
  return s
}

type GetTopAttackSourcesForClientIPsAreaTopDTO struct {
  // {'en':'Source area(country or area, province).', 'zh_CN':'来源地区（国家或地区、省份）。'}
  Area *string `json:"area,omitempty" xml:"area,omitempty" require:"true"`
  // {'en':'Attack requests.', 'zh_CN':'攻击请求数。'}
  Value *int64 `json:"value,omitempty" xml:"value,omitempty" require:"true"`
  // {'en':'Source IP.', 'zh_CN':'来源IP。'}
  Ip *string `json:"ip,omitempty" xml:"ip,omitempty" require:"true"`
}

func (s GetTopAttackSourcesForClientIPsAreaTopDTO) String() string {
  return tea.Prettify(s)
}

func (s GetTopAttackSourcesForClientIPsAreaTopDTO) GoString() string {
  return s.String()
}

func (s *GetTopAttackSourcesForClientIPsAreaTopDTO) SetArea(v string) *GetTopAttackSourcesForClientIPsAreaTopDTO {
  s.Area = &v
  return s
}

func (s *GetTopAttackSourcesForClientIPsAreaTopDTO) SetValue(v int64) *GetTopAttackSourcesForClientIPsAreaTopDTO {
  s.Value = &v
  return s
}

func (s *GetTopAttackSourcesForClientIPsAreaTopDTO) SetIp(v string) *GetTopAttackSourcesForClientIPsAreaTopDTO {
  s.Ip = &v
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

type GetTopAttackSourcesForClientIPsRequestHeader struct {
  // {'en':'The language of response data, default value: en.
  // zh_CN: Chinese
  // en: English', 'zh_CN':'返回内容的语言版本，默认值: en。
  // zh_CN：中文
  // en：英文'}
  Language *string `json:"language,omitempty" xml:"language,omitempty"`
  // {"zh_CN":"安全服务类型。有使用多个不同的安全服务时，需要填写具体的服务类型。","en":"Security service type. Please enter a specific service type, if you purchase multiple security services.","dictionary":"belong=WAAP-MS-Ext|dict=waap_serviceType"}
  ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty"`
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

type GetTopAttackSourcesForClientIPsResponseHeader struct {
}

func (s GetTopAttackSourcesForClientIPsResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetTopAttackSourcesForClientIPsResponseHeader) GoString() string {
  return s.String()
}




type GetSummaryRequestsRequest struct {
  // {'en':'Start time, format: yyyy-MM-dd HH:mm:ss.
  // The time range does not exceed 31 days.', 'zh_CN':'起始时间，格式：yyyy-MM-dd HH:mm:ss。
  // 时间范围不超过31天。'}
  StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty" require:"true"`
  // {'en':'End time, format: yyyy-MM-dd HH:mm:ss.
  // The time range does not exceed 31 days.', 'zh_CN':'截止时间，格式：yyyy-MM-dd HH:mm:ss。
  // 时间范围不超过31天。'}
  EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty" require:"true"`
  // {'en':'Hostname list, if not specified, it means all the hostnames of the account.', 'zh_CN':'域名列表，未指定时查询账号下的所有域名。'}
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

type GetSummaryRequestsResponse struct {
  // {'en':'Please refer to the error code for exceptions.', 'zh_CN':'请参照错误码。','dictionary':'belong=WAAP-MS-Ext|dict=waap_retCodeEnum'}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {'en':'Description.', 'zh_CN':'描述信息。'}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {'en':'Data.', 'zh_CN':'出参数据。'}
  Data *GetSummaryRequestsEventDTO `json:"data,omitempty" xml:"data,omitempty" require:"true"`
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

func (s *GetSummaryRequestsResponse) SetData(v *GetSummaryRequestsEventDTO) *GetSummaryRequestsResponse {
  s.Data = v
  return s
}

type GetSummaryRequestsEventDTO struct {
  // {'en':'Total requests.', 'zh_CN':'总请求数。'}
  Total *int64 `json:"total,omitempty" xml:"total,omitempty" require:"true"`
  // {'en':'Attack requests,sum of resisted requests and observed requests.', 'zh_CN':'攻击请求数，已抵御请求数与观察请求数之和。'}
  Attack *int64 `json:"attack,omitempty" xml:"attack,omitempty" require:"true"`
  // {'en':'Whitelist requests.', 'zh_CN':'白名单请求数。'}
  Whitelist *int64 `json:"whitelist,omitempty" xml:"whitelist,omitempty" require:"true"`
  // {'en':'Resisted requests.', 'zh_CN':'已抵御请求数。'}
  Resisted *int64 `json:"resisted,omitempty" xml:"resisted,omitempty" require:"true"`
  // {'en':'Observed requests.', 'zh_CN':'观察请求数。'}
  Observed *int64 `json:"observed,omitempty" xml:"observed,omitempty" require:"true"`
}

func (s GetSummaryRequestsEventDTO) String() string {
  return tea.Prettify(s)
}

func (s GetSummaryRequestsEventDTO) GoString() string {
  return s.String()
}

func (s *GetSummaryRequestsEventDTO) SetTotal(v int64) *GetSummaryRequestsEventDTO {
  s.Total = &v
  return s
}

func (s *GetSummaryRequestsEventDTO) SetAttack(v int64) *GetSummaryRequestsEventDTO {
  s.Attack = &v
  return s
}

func (s *GetSummaryRequestsEventDTO) SetWhitelist(v int64) *GetSummaryRequestsEventDTO {
  s.Whitelist = &v
  return s
}

func (s *GetSummaryRequestsEventDTO) SetResisted(v int64) *GetSummaryRequestsEventDTO {
  s.Resisted = &v
  return s
}

func (s *GetSummaryRequestsEventDTO) SetObserved(v int64) *GetSummaryRequestsEventDTO {
  s.Observed = &v
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

type GetSummaryRequestsRequestHeader struct {
  // {"zh_CN":"安全服务类型。有使用多个不同的安全服务时，需要填写具体的服务类型。","en":"Security service type. Please enter a specific service type, if you purchase multiple security services.","dictionary":"belong=WAAP-MS-Ext|dict=waap_serviceType"}
  ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty"`
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
  // {"en":"Start date format yyyy-MM-dd HH:mm:ss", "zh_CN":"起始日期 格式 yyyy-MM-dd HH:mm:ss"}
  Startdate *string `json:"startdate,omitempty" xml:"startdate,omitempty" require:"true"`
  // {"en":"End date format yyyy-MM-dd HH:mm:ss", "zh_CN":"结束日期 格式 yyyy-MM-dd HH:mm:ss"}
  Enddate *string `json:"enddate,omitempty" xml:"enddate,omitempty" require:"true"`
  // {"en":"Time zone, such as 28800000. The default is 28800000", "zh_CN":"时区，如：28800000 默认为：28800000"}
  Msec *string `json:"msec,omitempty" xml:"msec,omitempty"`
  // {"en":"Business type: dms by default", "zh_CN":"业务类型,默认为dms"}
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

type QueryDDoSAttackDetailsResponse struct {
  // {"en":"error response message", "zh_CN":"错误响应信息"}
  ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty" require:"true"`
  // {'en':'result', 'zh_CN':'结果'}
  Data []*QueryDDoSAttackDetailsResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
  // {"en":"error response code", "zh_CN":"错误响应码"}
  ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty" require:"true"`
  // {"en":"response code", "zh_CN":"响应码"}
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
  // {'en':'attack peak value', 'zh_CN':'攻击峰值'}
  TotalFlow *string `json:"totalFlow,omitempty" xml:"totalFlow,omitempty" require:"true"`
  // {'en':'IP', 'zh_CN':'IP'}
  Ip *string `json:"ip,omitempty" xml:"ip,omitempty" require:"true"`
  // {'en':'attack time', 'zh_CN':'攻击时间'}
  Time *string `json:"time,omitempty" xml:"time,omitempty" require:"true"`
  // {'en':'attack type', 'zh_CN':'攻击类型'}
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

type QueryDDoSAttackDetailsRequestHeader struct {
}

func (s QueryDDoSAttackDetailsRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryDDoSAttackDetailsRequestHeader) GoString() string {
  return s.String()
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
  // {'en':'Event ID.', 'zh_CN':'事件ID。'}
  Uuid *string `json:"uuid,omitempty" xml:"uuid,omitempty" require:"true"`
  // {'en':'Hostname.', 'zh_CN':'域名。'}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {'en':'Attack time, format: yyyy-MM-dd HH:mm:ss.', 'zh_CN':'攻击时间，格式：yyyy-MM-dd HH:mm:ss。'}
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

type GetOriginalRequestInformationResponse struct {
  // {'en':'Please refer to the error code for exceptions.', 'zh_CN':'请参照错误码。','dictionary':'belong=WAAP-MS-Ext|dict=waap_retCodeEnum'}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {'en':'Description.', 'zh_CN':'描述信息。'}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {'en':'Data.', 'zh_CN':'出参数据。'}
  Data *GetOriginalRequestInformationAttackLogSeniorInfoDto `json:"data,omitempty" xml:"data,omitempty" require:"true"`
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

func (s *GetOriginalRequestInformationResponse) SetData(v *GetOriginalRequestInformationAttackLogSeniorInfoDto) *GetOriginalRequestInformationResponse {
  s.Data = v
  return s
}

type GetOriginalRequestInformationAttackLogSeniorInfoDto struct {
  // {'en':'Request header information.', 'zh_CN':'请求头信息。'}
  RequestHead *string `json:"requestHead,omitempty" xml:"requestHead,omitempty" require:"true"`
}

func (s GetOriginalRequestInformationAttackLogSeniorInfoDto) String() string {
  return tea.Prettify(s)
}

func (s GetOriginalRequestInformationAttackLogSeniorInfoDto) GoString() string {
  return s.String()
}

func (s *GetOriginalRequestInformationAttackLogSeniorInfoDto) SetRequestHead(v string) *GetOriginalRequestInformationAttackLogSeniorInfoDto {
  s.RequestHead = &v
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

type GetOriginalRequestInformationRequestHeader struct {
  // {'en':'The language of response data, default value: en.
  // zh_CN: Chinese
  // en: English', 'zh_CN':'返回内容的语言版本，默认值: en。
  // zh_CN：中文
  // en：英文'}
  Language *string `json:"language,omitempty" xml:"language,omitempty"`
  // {"zh_CN":"安全服务类型。有使用多个不同的安全服务时，需要填写具体的服务类型。","en":"Security service type. Please enter a specific service type, if you purchase multiple security services.","dictionary":"belong=WAAP-MS-Ext|dict=waap_serviceType"}
  ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty"`
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
  // {"en":"Start time, format: yyyy-MM-dd HH:mm:ss.
  // The time range does not exceed 30 days.", "zh_CN":"开始时间，格式：yyyy-MM-dd HH:mm:ss。
  // 时间范围不超过30天。"}
  StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty" require:"true"`
  // {"en":"End time, format: yyyy-MM-dd HH:mm:ss.
  // The time range does not exceed 30 days)", "zh_CN":"结束时间，格式：yyyy-MM-dd HH:mm:ss。
  // 时间范围不超过30天。"}
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

type L4DdosTrendResponse struct {
  // {'en':'Please refer to the error code for exceptions.', 'zh_CN':'请参照错误码。','dictionary':'belong=WAAP-MS-Ext|dict=waap_retCodeEnum'}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Description.", "zh_CN":"描述信息。"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"Data.", "zh_CN":"出参数据。"}
  Data []*L4DdosTrendAttackedUrl `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
  // {"en":"Maximum attack bps.", "zh_CN":"DDoS攻击带宽峰值"}
  Peak_bps []*L4DdosTrendPeak `json:"peak_bps,omitempty" xml:"peak_bps,omitempty" require:"true" type:"Repeated"`
  // {"en":"Maximum attack qps.", "zh_CN":"DDoS攻击包速峰值"}
  Peak_pps []*L4DdosTrendPeak `json:"peak_pps,omitempty" xml:"peak_pps,omitempty" require:"true" type:"Repeated"`
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

func (s *L4DdosTrendResponse) SetData(v []*L4DdosTrendAttackedUrl) *L4DdosTrendResponse {
  s.Data = v
  return s
}

func (s *L4DdosTrendResponse) SetPeak_bps(v []*L4DdosTrendPeak) *L4DdosTrendResponse {
  s.Peak_bps = v
  return s
}

func (s *L4DdosTrendResponse) SetPeak_pps(v []*L4DdosTrendPeak) *L4DdosTrendResponse {
  s.Peak_pps = v
  return s
}

type L4DdosTrendAttackedUrl struct {
  // {"en":"Normal bandwidth.(bps)", "zh_CN":"正常带宽。（bps）"}
  Normal_bps *string `json:"normal_bps,omitempty" xml:"normal_bps,omitempty" require:"true"`
  // {"en":"Total cleaning bandwidth.", "zh_CN":"总清洗带宽。"}
  Attack_bps *string `json:"attack_bps,omitempty" xml:"attack_bps,omitempty" require:"true"`
  // {"en":"Normal packages.(pps)", "zh_CN":"正常包数。（pps）"}
  Normal_pps *string `json:"normal_pps,omitempty" xml:"normal_pps,omitempty" require:"true"`
  // {"en":"Total  packages.", "zh_CN":"总清洗包数。"}
  Attack_pps *string `json:"attack_pps,omitempty" xml:"attack_pps,omitempty" require:"true"`
  // {"en":"SYN Flood.(bps)", "zh_CN":"SYN Flood。（bps）"}
  Syn_bps *string `json:"syn_bps,omitempty" xml:"syn_bps,omitempty" require:"true"`
  // {"en":"SYN Flood.(pps)", "zh_CN":"SYN Flood。（pps）"}
  Syn_pps *string `json:"syn_pps,omitempty" xml:"syn_pps,omitempty" require:"true"`
  // {"en":"ACK Flood.(bps)", "zh_CN":"ACK Flood。（bps）"}
  Ack_bps *string `json:"ack_bps,omitempty" xml:"ack_bps,omitempty" require:"true"`
  // {"en":"ACK Flood.(pps)", "zh_CN":"ACK Flood。（pps）"}
  Ack_pps *string `json:"ack_pps,omitempty" xml:"ack_pps,omitempty" require:"true"`
  // {"en":"UDP Flood. (bps)", "zh_CN":"UDP Flood。 (bps)"}
  Udp_bps *string `json:"udp_bps,omitempty" xml:"udp_bps,omitempty" require:"true"`
  // {"en":"UDP Flood. (pps)", "zh_CN":"UDP Flood。 (pps)"}
  Udp_pps *string `json:"udp_pps,omitempty" xml:"udp_pps,omitempty" require:"true"`
  // {"en":"ICMP Flood. (bps)", "zh_CN":"ICMP Flood。 (bps)"}
  Icmp_bps *string `json:"icmp_bps,omitempty" xml:"icmp_bps,omitempty" require:"true"`
  // {"en":"ICMP Flood. (pps)", "zh_CN":"ICMP Flood。 (pps)"}
  Icmp_pps *string `json:"icmp_pps,omitempty" xml:"icmp_pps,omitempty" require:"true"`
  // {"en":"Other Flood. (pps)", "zh_CN":"Other Flood。 (pps)"}
  Other_pps *string `json:"other_pps,omitempty" xml:"other_pps,omitempty" require:"true"`
  // {"en":"Other Flood.(bps)", "zh_CN":"Other Flood。(bps)"}
  Other_bps *string `json:"other_bps,omitempty" xml:"other_bps,omitempty" require:"true"`
  // {"en":"Time.", "zh_CN":"时间。"}
  Time *string `json:"time,omitempty" xml:"time,omitempty" require:"true"`
}

func (s L4DdosTrendAttackedUrl) String() string {
  return tea.Prettify(s)
}

func (s L4DdosTrendAttackedUrl) GoString() string {
  return s.String()
}

func (s *L4DdosTrendAttackedUrl) SetNormal_bps(v string) *L4DdosTrendAttackedUrl {
  s.Normal_bps = &v
  return s
}

func (s *L4DdosTrendAttackedUrl) SetAttack_bps(v string) *L4DdosTrendAttackedUrl {
  s.Attack_bps = &v
  return s
}

func (s *L4DdosTrendAttackedUrl) SetNormal_pps(v string) *L4DdosTrendAttackedUrl {
  s.Normal_pps = &v
  return s
}

func (s *L4DdosTrendAttackedUrl) SetAttack_pps(v string) *L4DdosTrendAttackedUrl {
  s.Attack_pps = &v
  return s
}

func (s *L4DdosTrendAttackedUrl) SetSyn_bps(v string) *L4DdosTrendAttackedUrl {
  s.Syn_bps = &v
  return s
}

func (s *L4DdosTrendAttackedUrl) SetSyn_pps(v string) *L4DdosTrendAttackedUrl {
  s.Syn_pps = &v
  return s
}

func (s *L4DdosTrendAttackedUrl) SetAck_bps(v string) *L4DdosTrendAttackedUrl {
  s.Ack_bps = &v
  return s
}

func (s *L4DdosTrendAttackedUrl) SetAck_pps(v string) *L4DdosTrendAttackedUrl {
  s.Ack_pps = &v
  return s
}

func (s *L4DdosTrendAttackedUrl) SetUdp_bps(v string) *L4DdosTrendAttackedUrl {
  s.Udp_bps = &v
  return s
}

func (s *L4DdosTrendAttackedUrl) SetUdp_pps(v string) *L4DdosTrendAttackedUrl {
  s.Udp_pps = &v
  return s
}

func (s *L4DdosTrendAttackedUrl) SetIcmp_bps(v string) *L4DdosTrendAttackedUrl {
  s.Icmp_bps = &v
  return s
}

func (s *L4DdosTrendAttackedUrl) SetIcmp_pps(v string) *L4DdosTrendAttackedUrl {
  s.Icmp_pps = &v
  return s
}

func (s *L4DdosTrendAttackedUrl) SetOther_pps(v string) *L4DdosTrendAttackedUrl {
  s.Other_pps = &v
  return s
}

func (s *L4DdosTrendAttackedUrl) SetOther_bps(v string) *L4DdosTrendAttackedUrl {
  s.Other_bps = &v
  return s
}

func (s *L4DdosTrendAttackedUrl) SetTime(v string) *L4DdosTrendAttackedUrl {
  s.Time = &v
  return s
}

type L4DdosTrendPeak struct {
  // {"en":"Maximum attack time.", "zh_CN":"峰值时间"}
  Peak_time *string `json:"peak_time,omitempty" xml:"peak_time,omitempty" require:"true"`
  // {"en":"Maximum attack value.", "zh_CN":"峰值"}
  Peak_value *int64 `json:"peak_value,omitempty" xml:"peak_value,omitempty" require:"true"`
}

func (s L4DdosTrendPeak) String() string {
  return tea.Prettify(s)
}

func (s L4DdosTrendPeak) GoString() string {
  return s.String()
}

func (s *L4DdosTrendPeak) SetPeak_time(v string) *L4DdosTrendPeak {
  s.Peak_time = &v
  return s
}

func (s *L4DdosTrendPeak) SetPeak_value(v int64) *L4DdosTrendPeak {
  s.Peak_value = &v
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

type L4DdosTrendRequestHeader struct {
  // {"zh_CN":"安全服务类型。有使用多个不同的安全服务时，需要填写具体的服务类型。","en":"Security service type. Please enter a specific service type, if you purchase multiple security services.","dictionary":"belong=WAAP-MS-Ext|dict=waap_serviceType"}
  ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty"`
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




type GetTopAttackTargetsByHostnameRequest struct {
  // {'en':'Top rankings, default value: 100, max: 1000.', 'zh_CN':'取前几排名，默认：100，上限：1000。'}
  Top *int32 `json:"top,omitempty" xml:"top,omitempty"`
  // {'en':'Multiple choice. Disposal result, default value: all results.
  //  mitigated: Mitegaed requests.
  //  monitored: monitored requests.', 'zh_CN':'多选。处置结果，默认：展示所有结果。
  //  mitigated：已抵御请求数。
  //  monitored：观察请求数。'}
  ActType []*string `json:"actType,omitempty" xml:"actType,omitempty" type:"Repeated"`
  // {'en':'Start time, format: yyyy-MM-dd HH:mm:ss.
  // The time range does not exceed 31 days.', 'zh_CN':'起始时间，格式：yyyy-MM-dd HH:mm:ss。
  // 时间范围不超过31天。'}
  StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty" require:"true"`
  // {'en':'End time, format: yyyy-MM-dd HH:mm:ss.
  // The time range does not exceed 31 days.', 'zh_CN':'截止时间，格式：yyyy-MM-dd HH:mm:ss。
  // 时间范围不超过31天。'}
  EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty" require:"true"`
  // {'en':'Hostname list, if not specified, it means all the hostnames of the account.', 'zh_CN':'域名列表，未指定时查询账号下的所有域名。'}
  Domains []*string `json:"domains,omitempty" xml:"domains,omitempty" type:"Repeated"`
}

func (s GetTopAttackTargetsByHostnameRequest) String() string {
  return tea.Prettify(s)
}

func (s GetTopAttackTargetsByHostnameRequest) GoString() string {
  return s.String()
}

func (s *GetTopAttackTargetsByHostnameRequest) SetTop(v int32) *GetTopAttackTargetsByHostnameRequest {
  s.Top = &v
  return s
}

func (s *GetTopAttackTargetsByHostnameRequest) SetActType(v []*string) *GetTopAttackTargetsByHostnameRequest {
  s.ActType = v
  return s
}

func (s *GetTopAttackTargetsByHostnameRequest) SetStartTime(v string) *GetTopAttackTargetsByHostnameRequest {
  s.StartTime = &v
  return s
}

func (s *GetTopAttackTargetsByHostnameRequest) SetEndTime(v string) *GetTopAttackTargetsByHostnameRequest {
  s.EndTime = &v
  return s
}

func (s *GetTopAttackTargetsByHostnameRequest) SetDomains(v []*string) *GetTopAttackTargetsByHostnameRequest {
  s.Domains = v
  return s
}

type GetTopAttackTargetsByHostnameResponse struct {
  // {'en':'Please refer to the error code for exceptions.', 'zh_CN':'请参照错误码。','dictionary':'belong=WAAP-MS-Ext|dict=waap_retCodeEnum'}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {'en':'Description.', 'zh_CN':'描述信息。'}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {'en':'Data.', 'zh_CN':'出参数据。'}
  Data []*GetTopAttackTargetsByHostnameDomainTopDTO `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s GetTopAttackTargetsByHostnameResponse) String() string {
  return tea.Prettify(s)
}

func (s GetTopAttackTargetsByHostnameResponse) GoString() string {
  return s.String()
}

func (s *GetTopAttackTargetsByHostnameResponse) SetCode(v string) *GetTopAttackTargetsByHostnameResponse {
  s.Code = &v
  return s
}

func (s *GetTopAttackTargetsByHostnameResponse) SetMsg(v string) *GetTopAttackTargetsByHostnameResponse {
  s.Msg = &v
  return s
}

func (s *GetTopAttackTargetsByHostnameResponse) SetData(v []*GetTopAttackTargetsByHostnameDomainTopDTO) *GetTopAttackTargetsByHostnameResponse {
  s.Data = v
  return s
}

type GetTopAttackTargetsByHostnameDomainTopDTO struct {
  // {'en':'Hostname.', 'zh_CN':'域名。'}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {'en':'Attack requests.', 'zh_CN':'攻击请求数。'}
  Attack *int64 `json:"attack,omitempty" xml:"attack,omitempty" require:"true"`
}

func (s GetTopAttackTargetsByHostnameDomainTopDTO) String() string {
  return tea.Prettify(s)
}

func (s GetTopAttackTargetsByHostnameDomainTopDTO) GoString() string {
  return s.String()
}

func (s *GetTopAttackTargetsByHostnameDomainTopDTO) SetDomain(v string) *GetTopAttackTargetsByHostnameDomainTopDTO {
  s.Domain = &v
  return s
}

func (s *GetTopAttackTargetsByHostnameDomainTopDTO) SetAttack(v int64) *GetTopAttackTargetsByHostnameDomainTopDTO {
  s.Attack = &v
  return s
}

type GetTopAttackTargetsByHostnamePaths struct {
}

func (s GetTopAttackTargetsByHostnamePaths) String() string {
  return tea.Prettify(s)
}

func (s GetTopAttackTargetsByHostnamePaths) GoString() string {
  return s.String()
}

type GetTopAttackTargetsByHostnameParameters struct {
}

func (s GetTopAttackTargetsByHostnameParameters) String() string {
  return tea.Prettify(s)
}

func (s GetTopAttackTargetsByHostnameParameters) GoString() string {
  return s.String()
}

type GetTopAttackTargetsByHostnameRequestHeader struct {
  // {'en':'The language of response data, default value: en.
  // zh_CN: Chinese
  // en: English', 'zh_CN':'返回内容的语言版本，默认值: en。
  // zh_CN：中文
  // en：英文'}
  Language *string `json:"language,omitempty" xml:"language,omitempty"`
  // {"zh_CN":"安全服务类型。有使用多个不同的安全服务时，需要填写具体的服务类型。","en":"Security service type. Please enter a specific service type, if you purchase multiple security services.","dictionary":"belong=WAAP-MS-Ext|dict=waap_serviceType"}
  ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty"`
}

func (s GetTopAttackTargetsByHostnameRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetTopAttackTargetsByHostnameRequestHeader) GoString() string {
  return s.String()
}

func (s *GetTopAttackTargetsByHostnameRequestHeader) SetLanguage(v string) *GetTopAttackTargetsByHostnameRequestHeader {
  s.Language = &v
  return s
}

func (s *GetTopAttackTargetsByHostnameRequestHeader) SetServiceType(v string) *GetTopAttackTargetsByHostnameRequestHeader {
  s.ServiceType = &v
  return s
}

type GetTopAttackTargetsByHostnameResponseHeader struct {
}

func (s GetTopAttackTargetsByHostnameResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetTopAttackTargetsByHostnameResponseHeader) GoString() string {
  return s.String()
}




type GetTopPoliciesTriggeredRequest struct {
  // {'en':'Start time, format: yyyy-MM-dd HH:mm:ss.
  // The time range does not exceed 31 days.', 'zh_CN':'起始时间，格式：yyyy-MM-dd HH:mm:ss。
  // 时间范围不超过31天。'}
  StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty" require:"true"`
  // {'en':'End time, format: yyyy-MM-dd HH:mm:ss.
  // The time range does not exceed 31 days.', 'zh_CN':'截止时间，格式：yyyy-MM-dd HH:mm:ss。
  // 时间范围不超过31天。'}
  EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty" require:"true"`
  // {'en':'Hostname list, if not specified, it means all the hostnames of the account.', 'zh_CN':'域名列表，未指定时查询账号下的所有域名。'}
  Domains []*string `json:"domains,omitempty" xml:"domains,omitempty" type:"Repeated"`
  // {'en':'Multiple choice. Disposal result, default value: all results.
  //  mitigated: Mitegaed requests.
  //  monitored: monitored requests.', 'zh_CN':'多选。处置结果，默认：展示所有结果。
  //  mitigated：已抵御请求数。
  //  monitored：观察请求数。'}
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

type GetTopPoliciesTriggeredResponse struct {
  // {'en':'Please refer to the error code for exceptions.', 'zh_CN':'请参照错误码。','dictionary':'belong=WAAP-MS-Ext|dict=waap_retCodeEnum'}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {'en':'Description.', 'zh_CN':'描述信息。'}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {'en':'Data.', 'zh_CN':'出参数据。'}
  Data []*GetTopPoliciesTriggeredEventTypeDTO `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
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

func (s *GetTopPoliciesTriggeredResponse) SetData(v []*GetTopPoliciesTriggeredEventTypeDTO) *GetTopPoliciesTriggeredResponse {
  s.Data = v
  return s
}

type GetTopPoliciesTriggeredEventTypeDTO struct {
  // {'en':'Attack type.
  //  BLOCK: IP/Geo Block
  //  DMS_DEFEND: DDoS Protection
  //  WAF_DEFEND: WAF
  //  BOT_MANAGE: Bot Management
  //  API_DEFEND: API Security
  //  INTELLIGENCE: Threat Intelligence
  //  RATE_LIMIT: Rate Limiting
  //  CUSTOMIZE_RULE: Custom Rules', 'zh_CN':'攻击类型。
  //  BLOCK：IP/区域封禁
  //  DMS_DEFEND：DDoS防护
  //  WAF_DEFEND：WAF
  //  BOT_MANAGE：Bot管理
  //  API_DEFEND：API安全
  //  INTELLIGENCE：威胁情报
  //  RATE_LIMIT：频率限制
  //  CUSTOMIZE_RULE：自定义规则'}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {'en':'Number of requests that triggered the policy type.', 'zh_CN':'触发该策略类型的请求数。'}
  Value *int64 `json:"value,omitempty" xml:"value,omitempty" require:"true"`
}

func (s GetTopPoliciesTriggeredEventTypeDTO) String() string {
  return tea.Prettify(s)
}

func (s GetTopPoliciesTriggeredEventTypeDTO) GoString() string {
  return s.String()
}

func (s *GetTopPoliciesTriggeredEventTypeDTO) SetCode(v string) *GetTopPoliciesTriggeredEventTypeDTO {
  s.Code = &v
  return s
}

func (s *GetTopPoliciesTriggeredEventTypeDTO) SetValue(v int64) *GetTopPoliciesTriggeredEventTypeDTO {
  s.Value = &v
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

type GetTopPoliciesTriggeredRequestHeader struct {
  // {"zh_CN":"安全服务类型。有使用多个不同的安全服务时，需要填写具体的服务类型。","en":"Security service type. Please enter a specific service type, if you purchase multiple security services.","dictionary":"belong=WAAP-MS-Ext|dict=waap_serviceType"}
  ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty"`
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




type GetL7DdosAnalysisAttackedUrlListV2Request struct {
  // {"en":"Start time, format: yyyy-MM-dd HH:mm:ss.
  // The time range does not exceed 30 days.", "zh_CN":"开始时间，格式：yyyy-MM-dd HH:mm:ss。
  // 时间范围不超过30天。"}
  StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty" require:"true"`
  // {"en":"End time, format: yyyy-MM-dd HH:mm:ss.
  // The time range does not exceed 30 days)", "zh_CN":"结束时间，格式：yyyy-MM-dd HH:mm:ss。
  // 时间范围不超过30天。"}
  EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty" require:"true"`
  // {"en":"Hostname list.", "zh_CN":"域名数组。"}
  Hostnames []*string `json:"Hostnames,omitempty" xml:"Hostnames,omitempty" require:"true" type:"Repeated"`
  // {"en":"Top num.", "zh_CN":"默认10条。"}
  TopNum *int `json:"topNum,omitempty" xml:"topNum,omitempty"`
}

func (s GetL7DdosAnalysisAttackedUrlListV2Request) String() string {
  return tea.Prettify(s)
}

func (s GetL7DdosAnalysisAttackedUrlListV2Request) GoString() string {
  return s.String()
}

func (s *GetL7DdosAnalysisAttackedUrlListV2Request) SetStartTime(v string) *GetL7DdosAnalysisAttackedUrlListV2Request {
  s.StartTime = &v
  return s
}

func (s *GetL7DdosAnalysisAttackedUrlListV2Request) SetEndTime(v string) *GetL7DdosAnalysisAttackedUrlListV2Request {
  s.EndTime = &v
  return s
}

func (s *GetL7DdosAnalysisAttackedUrlListV2Request) SetHostnames(v []*string) *GetL7DdosAnalysisAttackedUrlListV2Request {
  s.Hostnames = v
  return s
}

func (s *GetL7DdosAnalysisAttackedUrlListV2Request) SetTopNum(v int) *GetL7DdosAnalysisAttackedUrlListV2Request {
  s.TopNum = &v
  return s
}

type GetL7DdosAnalysisAttackedUrlListV2Response struct {
  // {'en':'Please refer to the error code for exceptions.', 'zh_CN':'请参照错误码。','dictionary':'belong=WAAP-MS-Ext|dict=waap_retCodeEnum'}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Description.", "zh_CN":"描述信息。"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"Data.", "zh_CN":"出参数据。"}
  Data []*GetL7DdosAnalysisAttackedUrlListV2AttackedUrl `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s GetL7DdosAnalysisAttackedUrlListV2Response) String() string {
  return tea.Prettify(s)
}

func (s GetL7DdosAnalysisAttackedUrlListV2Response) GoString() string {
  return s.String()
}

func (s *GetL7DdosAnalysisAttackedUrlListV2Response) SetCode(v string) *GetL7DdosAnalysisAttackedUrlListV2Response {
  s.Code = &v
  return s
}

func (s *GetL7DdosAnalysisAttackedUrlListV2Response) SetMsg(v string) *GetL7DdosAnalysisAttackedUrlListV2Response {
  s.Msg = &v
  return s
}

func (s *GetL7DdosAnalysisAttackedUrlListV2Response) SetData(v []*GetL7DdosAnalysisAttackedUrlListV2AttackedUrl) *GetL7DdosAnalysisAttackedUrlListV2Response {
  s.Data = v
  return s
}

type GetL7DdosAnalysisAttackedUrlListV2AttackedUrl struct {
  // {"en":"Observed Requests.", "zh_CN":"观察请求数。"}
  AlarmCount *int64 `json:"alarmCount,omitempty" xml:"alarmCount,omitempty" require:"true"`
  // {"en":"Resisted Requests.", "zh_CN":"已抵御请求数。"}
  Count *int64 `json:"count,omitempty" xml:"count,omitempty" require:"true"`
  // {"en":"URL.", "zh_CN":"URL。"}
  Url *string `json:"url,omitempty" xml:"url,omitempty" require:"true"`
}

func (s GetL7DdosAnalysisAttackedUrlListV2AttackedUrl) String() string {
  return tea.Prettify(s)
}

func (s GetL7DdosAnalysisAttackedUrlListV2AttackedUrl) GoString() string {
  return s.String()
}

func (s *GetL7DdosAnalysisAttackedUrlListV2AttackedUrl) SetAlarmCount(v int64) *GetL7DdosAnalysisAttackedUrlListV2AttackedUrl {
  s.AlarmCount = &v
  return s
}

func (s *GetL7DdosAnalysisAttackedUrlListV2AttackedUrl) SetCount(v int64) *GetL7DdosAnalysisAttackedUrlListV2AttackedUrl {
  s.Count = &v
  return s
}

func (s *GetL7DdosAnalysisAttackedUrlListV2AttackedUrl) SetUrl(v string) *GetL7DdosAnalysisAttackedUrlListV2AttackedUrl {
  s.Url = &v
  return s
}

type GetL7DdosAnalysisAttackedUrlListV2Paths struct {
}

func (s GetL7DdosAnalysisAttackedUrlListV2Paths) String() string {
  return tea.Prettify(s)
}

func (s GetL7DdosAnalysisAttackedUrlListV2Paths) GoString() string {
  return s.String()
}

type GetL7DdosAnalysisAttackedUrlListV2Parameters struct {
}

func (s GetL7DdosAnalysisAttackedUrlListV2Parameters) String() string {
  return tea.Prettify(s)
}

func (s GetL7DdosAnalysisAttackedUrlListV2Parameters) GoString() string {
  return s.String()
}

type GetL7DdosAnalysisAttackedUrlListV2RequestHeader struct {
  // {"zh_CN":"安全服务类型。有使用多个不同的安全服务时，需要填写具体的服务类型。","en":"Security service type. Please enter a specific service type, if you purchase multiple security services.","dictionary":"belong=WAAP-MS-Ext|dict=waap_serviceType"}
  ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty"`
}

func (s GetL7DdosAnalysisAttackedUrlListV2RequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetL7DdosAnalysisAttackedUrlListV2RequestHeader) GoString() string {
  return s.String()
}

func (s *GetL7DdosAnalysisAttackedUrlListV2RequestHeader) SetServiceType(v string) *GetL7DdosAnalysisAttackedUrlListV2RequestHeader {
  s.ServiceType = &v
  return s
}

type GetL7DdosAnalysisAttackedUrlListV2ResponseHeader struct {
}

func (s GetL7DdosAnalysisAttackedUrlListV2ResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetL7DdosAnalysisAttackedUrlListV2ResponseHeader) GoString() string {
  return s.String()
}




type GetTriggeredRateLimitingRulesRequest struct {
  // {'en':'Top rankings, default value: 10, max: 1000.', 'zh_CN':'取前几排名，默认：10，上限：1000。'}
  Top *int32 `json:"top,omitempty" xml:"top,omitempty"`
  // {'en':'Multiple choice. Disposal result, default value: all results.
  //  mitigated: Mitegaed requests.
  //  monitored: monitored requests.', 'zh_CN':'多选。处置结果，默认：展示所有结果。
  //  mitigated：已抵御请求数。
  //  monitored：观察请求数。'}
  ActType []*string `json:"actType,omitempty" xml:"actType,omitempty" type:"Repeated"`
  // {'en':'Start time, format: yyyy-MM-dd HH:mm:ss.
  // The time range does not exceed 31 days.', 'zh_CN':'起始时间，格式：yyyy-MM-dd HH:mm:ss。
  // 时间范围不超过31天。'}
  StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty" require:"true"`
  // {'en':'End time, format: yyyy-MM-dd HH:mm:ss.
  // The time range does not exceed 31 days.', 'zh_CN':'截止时间，格式：yyyy-MM-dd HH:mm:ss。
  // 时间范围不超过31天。'}
  EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty" require:"true"`
  // {'en':'Hostname list, if not specified, it means all the hostnames of the account.', 'zh_CN':'域名列表，未指定时查询账号下的所有域名。'}
  Domains []*string `json:"domains,omitempty" xml:"domains,omitempty" type:"Repeated"`
}

func (s GetTriggeredRateLimitingRulesRequest) String() string {
  return tea.Prettify(s)
}

func (s GetTriggeredRateLimitingRulesRequest) GoString() string {
  return s.String()
}

func (s *GetTriggeredRateLimitingRulesRequest) SetTop(v int32) *GetTriggeredRateLimitingRulesRequest {
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

type GetTriggeredRateLimitingRulesResponse struct {
  // {'en':'Please refer to the error code for exceptions.', 'zh_CN':'请参照错误码。','dictionary':'belong=WAAP-MS-Ext|dict=waap_retCodeEnum'}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {'en':'Description.', 'zh_CN':'描述信息。'}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {'en':'Data.', 'zh_CN':'出参数据。'}
  Data []*GetTriggeredRateLimitingRulesRuleTopDTO `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
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

func (s *GetTriggeredRateLimitingRulesResponse) SetData(v []*GetTriggeredRateLimitingRulesRuleTopDTO) *GetTriggeredRateLimitingRulesResponse {
  s.Data = v
  return s
}

type GetTriggeredRateLimitingRulesRuleHitDTO struct {
  // {'en':'Action.', 'zh_CN':'采取动作。'}
  Act *string `json:"act,omitempty" xml:"act,omitempty" require:"true"`
  // {'en':'Hit times.', 'zh_CN':'命中数。'}
  Value *int64 `json:"value,omitempty" xml:"value,omitempty" require:"true"`
}

func (s GetTriggeredRateLimitingRulesRuleHitDTO) String() string {
  return tea.Prettify(s)
}

func (s GetTriggeredRateLimitingRulesRuleHitDTO) GoString() string {
  return s.String()
}

func (s *GetTriggeredRateLimitingRulesRuleHitDTO) SetAct(v string) *GetTriggeredRateLimitingRulesRuleHitDTO {
  s.Act = &v
  return s
}

func (s *GetTriggeredRateLimitingRulesRuleHitDTO) SetValue(v int64) *GetTriggeredRateLimitingRulesRuleHitDTO {
  s.Value = &v
  return s
}

type GetTriggeredRateLimitingRulesRuleTopDTO struct {
  // {'en':'Rule ID.', 'zh_CN':'规则ID。'}
  RuleId *string `json:"ruleId,omitempty" xml:"ruleId,omitempty" require:"true"`
  // {'en':'Rule name.', 'zh_CN':'规则名称。'}
  RuleName *string `json:"ruleName,omitempty" xml:"ruleName,omitempty" require:"true"`
  // {'en':'Protected target.', 'zh_CN':'业务场景。'}
  Scene *string `json:"scene,omitempty" xml:"scene,omitempty" require:"true"`
  // {'en':'Trigger times, sort by times.', 'zh_CN':'触发次数，按次数排序。'}
  Hits []*GetTriggeredRateLimitingRulesRuleHitDTO `json:"hits,omitempty" xml:"hits,omitempty" require:"true" type:"Repeated"`
}

func (s GetTriggeredRateLimitingRulesRuleTopDTO) String() string {
  return tea.Prettify(s)
}

func (s GetTriggeredRateLimitingRulesRuleTopDTO) GoString() string {
  return s.String()
}

func (s *GetTriggeredRateLimitingRulesRuleTopDTO) SetRuleId(v string) *GetTriggeredRateLimitingRulesRuleTopDTO {
  s.RuleId = &v
  return s
}

func (s *GetTriggeredRateLimitingRulesRuleTopDTO) SetRuleName(v string) *GetTriggeredRateLimitingRulesRuleTopDTO {
  s.RuleName = &v
  return s
}

func (s *GetTriggeredRateLimitingRulesRuleTopDTO) SetScene(v string) *GetTriggeredRateLimitingRulesRuleTopDTO {
  s.Scene = &v
  return s
}

func (s *GetTriggeredRateLimitingRulesRuleTopDTO) SetHits(v []*GetTriggeredRateLimitingRulesRuleHitDTO) *GetTriggeredRateLimitingRulesRuleTopDTO {
  s.Hits = v
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

type GetTriggeredRateLimitingRulesRequestHeader struct {
  // {'en':'The language of response data, default value: en.
  // zh_CN: Chinese
  // en: English', 'zh_CN':'返回内容的语言版本，默认值: en。
  // zh_CN：中文
  // en：英文'}
  Language *string `json:"language,omitempty" xml:"language,omitempty"`
  // {"zh_CN":"安全服务类型。有使用多个不同的安全服务时，需要填写具体的服务类型。","en":"Security service type. Please enter a specific service type, if you purchase multiple security services.","dictionary":"belong=WAAP-MS-Ext|dict=waap_serviceType"}
  ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty"`
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
  // {'en':'Top rankings, default value: 10, max: 1000.', 'zh_CN':'取前几排名，默认：10，上限：1000。'}
  Top *int32 `json:"top,omitempty" xml:"top,omitempty"`
  // {'en':'Multiple choice. Disposal result, default value: all results.
  //  mitigated: Mitegaed requests.
  //  monitored: monitored requests.', 'zh_CN':'多选。处置结果，默认：展示所有结果。
  //  mitigated：已抵御请求数。
  //  monitored：观察请求数。'}
  ActType []*string `json:"actType,omitempty" xml:"actType,omitempty" type:"Repeated"`
  // {'en':'Start time, format: yyyy-MM-dd HH:mm:ss.
  // The time range does not exceed 31 days.', 'zh_CN':'起始时间，格式：yyyy-MM-dd HH:mm:ss。
  // 时间范围不超过31天。'}
  StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty" require:"true"`
  // {'en':'End time, format: yyyy-MM-dd HH:mm:ss.
  // The time range does not exceed 31 days.', 'zh_CN':'截止时间，格式：yyyy-MM-dd HH:mm:ss。
  // 时间范围不超过31天。'}
  EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty" require:"true"`
  // {'en':'Hostname list, if not specified, it means all the hostnames of the account.', 'zh_CN':'域名列表，未指定时查询账号下的所有域名。'}
  Domains []*string `json:"domains,omitempty" xml:"domains,omitempty" type:"Repeated"`
}

func (s GetTriggeredWAFManagedRulesRequest) String() string {
  return tea.Prettify(s)
}

func (s GetTriggeredWAFManagedRulesRequest) GoString() string {
  return s.String()
}

func (s *GetTriggeredWAFManagedRulesRequest) SetTop(v int32) *GetTriggeredWAFManagedRulesRequest {
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

type GetTriggeredWAFManagedRulesResponse struct {
  // {'en':'Please refer to the error code for exceptions.', 'zh_CN':'请参照错误码。','dictionary':'belong=WAAP-MS-Ext|dict=waap_retCodeEnum'}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {'en':'Description.', 'zh_CN':'描述信息。'}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {'en':'Data.', 'zh_CN':'出参数据。'}
  Data []*GetTriggeredWAFManagedRulesRuleTopDTO `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
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

func (s *GetTriggeredWAFManagedRulesResponse) SetData(v []*GetTriggeredWAFManagedRulesRuleTopDTO) *GetTriggeredWAFManagedRulesResponse {
  s.Data = v
  return s
}

type GetTriggeredWAFManagedRulesRuleHitDTO struct {
  // {'en':'Action.', 'zh_CN':'采取动作。'}
  Act *string `json:"act,omitempty" xml:"act,omitempty" require:"true"`
  // {'en':'Hit times.', 'zh_CN':'命中数。'}
  Value *int64 `json:"value,omitempty" xml:"value,omitempty" require:"true"`
}

func (s GetTriggeredWAFManagedRulesRuleHitDTO) String() string {
  return tea.Prettify(s)
}

func (s GetTriggeredWAFManagedRulesRuleHitDTO) GoString() string {
  return s.String()
}

func (s *GetTriggeredWAFManagedRulesRuleHitDTO) SetAct(v string) *GetTriggeredWAFManagedRulesRuleHitDTO {
  s.Act = &v
  return s
}

func (s *GetTriggeredWAFManagedRulesRuleHitDTO) SetValue(v int64) *GetTriggeredWAFManagedRulesRuleHitDTO {
  s.Value = &v
  return s
}

type GetTriggeredWAFManagedRulesRuleTopDTO struct {
  // {'en':'Rule ID.', 'zh_CN':'规则ID。'}
  RuleId *string `json:"ruleId,omitempty" xml:"ruleId,omitempty" require:"true"`
  // {'en':'Rule name.', 'zh_CN':'规则名称。'}
  RuleName *string `json:"ruleName,omitempty" xml:"ruleName,omitempty" require:"true"`
  // {'en':'Rule type.', 'zh_CN':'规则类型。'}
  RuleType *string `json:"ruleType,omitempty" xml:"ruleType,omitempty" require:"true"`
  // {'en':'System recommended action.', 'zh_CN':'系统推荐动作。'}
  Action *string `json:"action,omitempty" xml:"action,omitempty" require:"true"`
  // {'en':'Trigger times, sort by times.', 'zh_CN':'触发次数，按次数排序。'}
  Hits []*GetTriggeredWAFManagedRulesRuleHitDTO `json:"hits,omitempty" xml:"hits,omitempty" require:"true" type:"Repeated"`
}

func (s GetTriggeredWAFManagedRulesRuleTopDTO) String() string {
  return tea.Prettify(s)
}

func (s GetTriggeredWAFManagedRulesRuleTopDTO) GoString() string {
  return s.String()
}

func (s *GetTriggeredWAFManagedRulesRuleTopDTO) SetRuleId(v string) *GetTriggeredWAFManagedRulesRuleTopDTO {
  s.RuleId = &v
  return s
}

func (s *GetTriggeredWAFManagedRulesRuleTopDTO) SetRuleName(v string) *GetTriggeredWAFManagedRulesRuleTopDTO {
  s.RuleName = &v
  return s
}

func (s *GetTriggeredWAFManagedRulesRuleTopDTO) SetRuleType(v string) *GetTriggeredWAFManagedRulesRuleTopDTO {
  s.RuleType = &v
  return s
}

func (s *GetTriggeredWAFManagedRulesRuleTopDTO) SetAction(v string) *GetTriggeredWAFManagedRulesRuleTopDTO {
  s.Action = &v
  return s
}

func (s *GetTriggeredWAFManagedRulesRuleTopDTO) SetHits(v []*GetTriggeredWAFManagedRulesRuleHitDTO) *GetTriggeredWAFManagedRulesRuleTopDTO {
  s.Hits = v
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

type GetTriggeredWAFManagedRulesRequestHeader struct {
  // {'en':'The language of response data, default value: en.
  // zh_CN: Chinese
  // en: English', 'zh_CN':'返回内容的语言版本，默认值: en。
  // zh_CN：中文
  // en：英文'}
  Language *string `json:"language,omitempty" xml:"language,omitempty"`
  // {"zh_CN":"安全服务类型。有使用多个不同的安全服务时，需要填写具体的服务类型。","en":"Security service type. Please enter a specific service type, if you purchase multiple security services.","dictionary":"belong=WAAP-MS-Ext|dict=waap_serviceType"}
  ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty"`
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
  // {'en':'Start time, format: yyyy-MM-dd HH:mm:ss.
  // The time range does not exceed 31 days.', 'zh_CN':'开始时间，格式：yyyy-MM-dd HH:mm:ss。
  // 时间范围不超过31天。'}
  StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty" require:"true"`
  // {'en':'End time, format: yyyy-MM-dd HH:mm:ss.
  // The time range does not exceed 31 days.', 'zh_CN':'结束时间，格式：yyyy-MM-dd HH:mm:ss。
  // 时间范围不超过31天。'}
  EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty" require:"true"`
  // {'en':'Multiple choice. Disposal result, default value: all results.
  //  mitigated: Mitegaed requests.
  //  monitored: monitored requests.', 'zh_CN':'多选。处置结果，默认：展示所有结果。
  //  mitigated：已抵御请求数。
  //  monitored：观察请求数。'}
  ActType []*string `json:"actType,omitempty" xml:"actType,omitempty" type:"Repeated"`
  // {'en':'Hostname list, if not specified, it means all the hostnames of the account.', 'zh_CN':'域名列表，未指定时查询账号下的所有域名。'}
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

type QueryEventTrendResponse struct {
  // {'en':'Please refer to the error code for exceptions.', 'zh_CN':'请参照错误码。','dictionary':'belong=WAAP-MS-Ext|dict=waap_retCodeEnum'}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {'en':'Description.', 'zh_CN':'描述信息。'}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {'en':'Data.', 'zh_CN':'出参数据。'}
  Data []*QueryEventTrendEventTrendDTO `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
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

func (s *QueryEventTrendResponse) SetData(v []*QueryEventTrendEventTrendDTO) *QueryEventTrendResponse {
  s.Data = v
  return s
}

type QueryEventTrendEventTypeDTO struct {
  // {'en':'Policy type.
  //  BLOCK: IP/Geo Block
  //  DMS_DEFEND: DDoS Protection
  //  WAF_DEFEND: WAF
  //  BOT_MANAGE: Bot Management
  //  API_DEFEND: API Security
  //  INTELLIGENCE: Threat Intelligence
  //  RATE_LIMIT: Rate Limiting
  //  CUSTOMIZE_RULE: Custom Rules', 'zh_CN':'策略类型。
  //  BLOCK：IP/区域封禁
  //  DMS_DEFEND：DDoS防护
  //  WAF_DEFEND：WAF
  //  BOT_MANAGE：Bot管理
  //  API_DEFEND：API安全
  //  INTELLIGENCE：威胁情报
  //  RATE_LIMIT：频率限制
  //  CUSTOMIZE_RULE：自定义规则'}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {'en':'Number of policy requests of this type.', 'zh_CN':'该策略类型请求数。'}
  Value *int64 `json:"value,omitempty" xml:"value,omitempty" require:"true"`
}

func (s QueryEventTrendEventTypeDTO) String() string {
  return tea.Prettify(s)
}

func (s QueryEventTrendEventTypeDTO) GoString() string {
  return s.String()
}

func (s *QueryEventTrendEventTypeDTO) SetCode(v string) *QueryEventTrendEventTypeDTO {
  s.Code = &v
  return s
}

func (s *QueryEventTrendEventTypeDTO) SetValue(v int64) *QueryEventTrendEventTypeDTO {
  s.Value = &v
  return s
}

type QueryEventTrendEventTrendDTO struct {
  // {'en':'Time point, format: yyyy-MM-dd HH-mm-ss.', 'zh_CN':'时间点，格式：yyyy-MM-dd HH-mm-ss。'}
  TimePoint *string `json:"timePoint,omitempty" xml:"timePoint,omitempty" require:"true"`
  // {'en':'Total requests.', 'zh_CN':'总请求数。'}
  Total *int64 `json:"total,omitempty" xml:"total,omitempty" require:"true"`
  // {'en':'Attack requests.', 'zh_CN':'攻击请求数。'}
  Attack *int64 `json:"attack,omitempty" xml:"attack,omitempty" require:"true"`
  // {'en':'Mitigated requests.', 'zh_CN':'已抵御请求数。'}
  Mitigated *int64 `json:"mitigated,omitempty" xml:"mitigated,omitempty" require:"true"`
  // {'en':'Monitored requests.', 'zh_CN':'观察请求数。'}
  Monitored *int64 `json:"monitored,omitempty" xml:"monitored,omitempty" require:"true"`
  // {'en':'Whitelist requests.', 'zh_CN':'白名单请求数。'}
  Whitelist *int64 `json:"whitelist,omitempty" xml:"whitelist,omitempty" require:"true"`
  // {'en':'Policy type classification requests.', 'zh_CN':'策略类型分类请求数。'}
  Distribution []*QueryEventTrendEventTypeDTO `json:"distribution,omitempty" xml:"distribution,omitempty" require:"true" type:"Repeated"`
}

func (s QueryEventTrendEventTrendDTO) String() string {
  return tea.Prettify(s)
}

func (s QueryEventTrendEventTrendDTO) GoString() string {
  return s.String()
}

func (s *QueryEventTrendEventTrendDTO) SetTimePoint(v string) *QueryEventTrendEventTrendDTO {
  s.TimePoint = &v
  return s
}

func (s *QueryEventTrendEventTrendDTO) SetTotal(v int64) *QueryEventTrendEventTrendDTO {
  s.Total = &v
  return s
}

func (s *QueryEventTrendEventTrendDTO) SetAttack(v int64) *QueryEventTrendEventTrendDTO {
  s.Attack = &v
  return s
}

func (s *QueryEventTrendEventTrendDTO) SetMitigated(v int64) *QueryEventTrendEventTrendDTO {
  s.Mitigated = &v
  return s
}

func (s *QueryEventTrendEventTrendDTO) SetMonitored(v int64) *QueryEventTrendEventTrendDTO {
  s.Monitored = &v
  return s
}

func (s *QueryEventTrendEventTrendDTO) SetWhitelist(v int64) *QueryEventTrendEventTrendDTO {
  s.Whitelist = &v
  return s
}

func (s *QueryEventTrendEventTrendDTO) SetDistribution(v []*QueryEventTrendEventTypeDTO) *QueryEventTrendEventTrendDTO {
  s.Distribution = v
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

type QueryEventTrendRequestHeader struct {
  // {"zh_CN":"安全服务类型。有使用多个不同的安全服务时，需要填写具体的服务类型。","en":"Security service type. Please enter a specific service type, if you purchase multiple security services.","dictionary":"belong=WAAP-MS-Ext|dict=waap_serviceType"}
  ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty"`
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
  ControlGroupCode *string `json:"controlGroupCode,omitempty" xml:"controlGroupCode,omitempty"`
  // {"en":"End date, format: yyyy-MM-dd.The time range does not exceed 31 days, only supports querying data within one year..","zh_CN":"结束日期，格式：yyyy-MM-dd。时间范围不超过31天，仅支持查询一年内的数据。"}
  EndDate *string `json:"endDate,omitempty" xml:"endDate,omitempty" require:"true"`
  // {"en":"Type of Control Group.ALL: All  PART: Custom category  DOMAIN_ITEM: Contract category","zh_CN":"Control Group的类型。ALL： 所有  PART：自定义类 DOMAIN_ITEM： 合同类"}
  ControlGroupType *string `json:"controlGroupType,omitempty" xml:"controlGroupType,omitempty"`
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
  // {"zh_CN":"安全服务类型。有使用多个不同的安全服务时，需要填写具体的服务类型。","en":"Security service type. Please enter a specific service type, if you purchase multiple security services.","dictionary":"belong=WAAP-MS-Ext|dict=waap_serviceType"}
  ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty"`
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
  // {"en":"Please refer to the error code for exceptions.","zh_CN":"请参照错误码。"}
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




type GetL7DdosAnalysisAttackedHostnameListV2Request struct {
  // {"en":"Start time, format: yyyy-MM-dd HH:mm:ss.
  // The time range does not exceed 30 days.", "zh_CN":"开始时间，格式：yyyy-MM-dd HH:mm:ss。
  // 时间范围不超过30天。"}
  StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty" require:"true"`
  // {"en":"End time, format: yyyy-MM-dd HH:mm:ss.
  // The time range does not exceed 30 days)", "zh_CN":"结束时间，格式：yyyy-MM-dd HH:mm:ss。
  // 时间范围不超过30天。"}
  EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty" require:"true"`
  // {"en":"Hostname list.", "zh_CN":"域名数组"}
  Hostnames []*string `json:"Hostnames,omitempty" xml:"Hostnames,omitempty" require:"true" type:"Repeated"`
  // {"en":"Top num.", "zh_CN":"默认10条"}
  TopNum *int `json:"topNum,omitempty" xml:"topNum,omitempty"`
}

func (s GetL7DdosAnalysisAttackedHostnameListV2Request) String() string {
  return tea.Prettify(s)
}

func (s GetL7DdosAnalysisAttackedHostnameListV2Request) GoString() string {
  return s.String()
}

func (s *GetL7DdosAnalysisAttackedHostnameListV2Request) SetStartTime(v string) *GetL7DdosAnalysisAttackedHostnameListV2Request {
  s.StartTime = &v
  return s
}

func (s *GetL7DdosAnalysisAttackedHostnameListV2Request) SetEndTime(v string) *GetL7DdosAnalysisAttackedHostnameListV2Request {
  s.EndTime = &v
  return s
}

func (s *GetL7DdosAnalysisAttackedHostnameListV2Request) SetHostnames(v []*string) *GetL7DdosAnalysisAttackedHostnameListV2Request {
  s.Hostnames = v
  return s
}

func (s *GetL7DdosAnalysisAttackedHostnameListV2Request) SetTopNum(v int) *GetL7DdosAnalysisAttackedHostnameListV2Request {
  s.TopNum = &v
  return s
}

type GetL7DdosAnalysisAttackedHostnameListV2Response struct {
  // {'en':'Please refer to the error code for exceptions.', 'zh_CN':'请参照错误码。','dictionary':'belong=WAAP-MS-Ext|dict=waap_retCodeEnum'}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Description.", "zh_CN":"描述信息。"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"Data.", "zh_CN":"出参数据。"}
  Data []*GetL7DdosAnalysisAttackedHostnameListV2AttackedDomain `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s GetL7DdosAnalysisAttackedHostnameListV2Response) String() string {
  return tea.Prettify(s)
}

func (s GetL7DdosAnalysisAttackedHostnameListV2Response) GoString() string {
  return s.String()
}

func (s *GetL7DdosAnalysisAttackedHostnameListV2Response) SetCode(v string) *GetL7DdosAnalysisAttackedHostnameListV2Response {
  s.Code = &v
  return s
}

func (s *GetL7DdosAnalysisAttackedHostnameListV2Response) SetMsg(v string) *GetL7DdosAnalysisAttackedHostnameListV2Response {
  s.Msg = &v
  return s
}

func (s *GetL7DdosAnalysisAttackedHostnameListV2Response) SetData(v []*GetL7DdosAnalysisAttackedHostnameListV2AttackedDomain) *GetL7DdosAnalysisAttackedHostnameListV2Response {
  s.Data = v
  return s
}

type GetL7DdosAnalysisAttackedHostnameListV2AttackedDomain struct {
  // {"en":"All count.", "zh_CN":"总请求。"}
  All_count *int64 `json:"all_count,omitempty" xml:"all_count,omitempty" require:"true"`
  // {"en":"Observed Requests..", "zh_CN":"观察请求数。"}
  AlarmCount *int64 `json:"alarmCount,omitempty" xml:"alarmCount,omitempty" require:"true"`
  // {"en":"Observed Ratio.", "zh_CN":"观察请求占比。"}
  AlarmRatio *float64 `json:"alarmRatio,omitempty" xml:"alarmRatio,omitempty" require:"true"`
  // {"en":"Resisted Requests..", "zh_CN":"已抵御请求数。"}
  Count *int64 `json:"count,omitempty" xml:"count,omitempty" require:"true"`
  // {"en":"Resisted Ratio.", "zh_CN":"已抵御请求占比。"}
  Ratio *float64 `json:"ratio,omitempty" xml:"ratio,omitempty" require:"true"`
  // {"en":"Hostname.", "zh_CN":"域名。"}
  Hostname *string `json:"Hostname,omitempty" xml:"Hostname,omitempty" require:"true"`
}

func (s GetL7DdosAnalysisAttackedHostnameListV2AttackedDomain) String() string {
  return tea.Prettify(s)
}

func (s GetL7DdosAnalysisAttackedHostnameListV2AttackedDomain) GoString() string {
  return s.String()
}

func (s *GetL7DdosAnalysisAttackedHostnameListV2AttackedDomain) SetAll_count(v int64) *GetL7DdosAnalysisAttackedHostnameListV2AttackedDomain {
  s.All_count = &v
  return s
}

func (s *GetL7DdosAnalysisAttackedHostnameListV2AttackedDomain) SetAlarmCount(v int64) *GetL7DdosAnalysisAttackedHostnameListV2AttackedDomain {
  s.AlarmCount = &v
  return s
}

func (s *GetL7DdosAnalysisAttackedHostnameListV2AttackedDomain) SetAlarmRatio(v float64) *GetL7DdosAnalysisAttackedHostnameListV2AttackedDomain {
  s.AlarmRatio = &v
  return s
}

func (s *GetL7DdosAnalysisAttackedHostnameListV2AttackedDomain) SetCount(v int64) *GetL7DdosAnalysisAttackedHostnameListV2AttackedDomain {
  s.Count = &v
  return s
}

func (s *GetL7DdosAnalysisAttackedHostnameListV2AttackedDomain) SetRatio(v float64) *GetL7DdosAnalysisAttackedHostnameListV2AttackedDomain {
  s.Ratio = &v
  return s
}

func (s *GetL7DdosAnalysisAttackedHostnameListV2AttackedDomain) SetHostname(v string) *GetL7DdosAnalysisAttackedHostnameListV2AttackedDomain {
  s.Hostname = &v
  return s
}

type GetL7DdosAnalysisAttackedHostnameListV2Paths struct {
}

func (s GetL7DdosAnalysisAttackedHostnameListV2Paths) String() string {
  return tea.Prettify(s)
}

func (s GetL7DdosAnalysisAttackedHostnameListV2Paths) GoString() string {
  return s.String()
}

type GetL7DdosAnalysisAttackedHostnameListV2Parameters struct {
}

func (s GetL7DdosAnalysisAttackedHostnameListV2Parameters) String() string {
  return tea.Prettify(s)
}

func (s GetL7DdosAnalysisAttackedHostnameListV2Parameters) GoString() string {
  return s.String()
}

type GetL7DdosAnalysisAttackedHostnameListV2RequestHeader struct {
  // {"zh_CN":"安全服务类型。有使用多个不同的安全服务时，需要填写具体的服务类型。","en":"Security service type. Please enter a specific service type, if you purchase multiple security services.","dictionary":"belong=WAAP-MS-Ext|dict=waap_serviceType"}
  ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty"`
}

func (s GetL7DdosAnalysisAttackedHostnameListV2RequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetL7DdosAnalysisAttackedHostnameListV2RequestHeader) GoString() string {
  return s.String()
}

func (s *GetL7DdosAnalysisAttackedHostnameListV2RequestHeader) SetServiceType(v string) *GetL7DdosAnalysisAttackedHostnameListV2RequestHeader {
  s.ServiceType = &v
  return s
}

type GetL7DdosAnalysisAttackedHostnameListV2ResponseHeader struct {
}

func (s GetL7DdosAnalysisAttackedHostnameListV2ResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetL7DdosAnalysisAttackedHostnameListV2ResponseHeader) GoString() string {
  return s.String()
}




type GetInfrastructureLogListRequest struct {
  // {'en':'Start time, format:yyyy-MM-dd HH:mm:ss. 
  // Only supports querying logs of the past month, the query time range cannot exceed 24 hours.', 'zh_CN':'开始时间，格式：yyyy-MM-dd HH:mm:ss。
  // 仅支持查询近一个月的日志，查询时间范围不能超过24小时。'}
  StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty" require:"true"`
  // {'en':'End time, format:yyyy-MM-dd HH:mm:ss. 
  // Only supports querying logs of the past month, the query time range cannot exceed 24 hours.', 'zh_CN':'结束时间，格式：yyyy-MM-dd HH:mm:ss。
  // 仅支持查询近一个月的日志，查询时间范围不能超过24小时，格式：yyyy-MM-dd HH:mm:ss。'}
  EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty" require:"true"`
  // {'en':'Client IP.', 'zh_CN':'客户端IP。'}
  ClientIps []*string `json:"clientIps,omitempty" xml:"clientIps,omitempty" type:"Repeated"`
  // {'en':'Node IP.', 'zh_CN':'节点IP。'}
  TargetIps []*string `json:"targetIps,omitempty" xml:"targetIps,omitempty" type:"Repeated"`
  // {'en':'Rule Name.', 'zh_CN':'规则名称。'}
  RuleName *string `json:"ruleName,omitempty" xml:"ruleName,omitempty"`
  // {'en':'Policy Name.
  // rule_protection:Managed Ruleset Protection', 'zh_CN':'策略名称。
  // rule_protection：内置防护规则'}
  PolicyNames []*string `json:"policyNames,omitempty" xml:"policyNames,omitempty" require:"true" type:"Repeated"`
  // {'en':'Action.
  // connection_denied:Connection Refused', 'zh_CN':'处理动作。
  // connection_denied：拒绝连接'}
  Actions []*string `json:"actions,omitempty" xml:"actions,omitempty" require:"true" type:"Repeated"`
  // {'en':'The number of records per page, default value:10', 'zh_CN':'每页显示的条目数。默认值：10'}
  PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
  // {'en':'Current page, default:1', 'zh_CN':'当前第几页。默认值：1'}
  CurrentPage *int32 `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
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

func (s *GetInfrastructureLogListRequest) SetPageSize(v int32) *GetInfrastructureLogListRequest {
  s.PageSize = &v
  return s
}

func (s *GetInfrastructureLogListRequest) SetCurrentPage(v int32) *GetInfrastructureLogListRequest {
  s.CurrentPage = &v
  return s
}

type GetInfrastructureLogListResponse struct {
  // {'en':'Please refer to the error code for exceptions.', 'zh_CN':'请参照错误码。','dictionary':'belong=WAAP-MS-Ext|dict=waap_retCodeEnum'}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {'en':'Description.', 'zh_CN':'描述信息。'}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {'en':'Data.', 'zh_CN':'出参数据。'}
  Data *GetInfrastructureLogListPageVO `json:"data,omitempty" xml:"data,omitempty" require:"true"`
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

func (s *GetInfrastructureLogListResponse) SetData(v *GetInfrastructureLogListPageVO) *GetInfrastructureLogListResponse {
  s.Data = v
  return s
}

type GetInfrastructureLogListPageVO struct {
  // {'en':'The current page number.', 'zh_CN':'当前页码。'}
  Current *int64 `json:"current,omitempty" xml:"current,omitempty" require:"true"`
  // {'en':'The number of records per page.', 'zh_CN':'每页显示的条目数。'}
  Size *int64 `json:"size,omitempty" xml:"size,omitempty" require:"true"`
  // {'en':'Pages.', 'zh_CN':'页数。'}
  Pages *int64 `json:"pages,omitempty" xml:"pages,omitempty" require:"true"`
  // {'en':'The total number of records.', 'zh_CN':'总记录数。'}
  Total *int64 `json:"total,omitempty" xml:"total,omitempty" require:"true"`
  // {'en':'Records.', 'zh_CN':'记录。'}
  Records *GetInfrastructureLogListCsecInfrastructureLogVO `json:"records,omitempty" xml:"records,omitempty" require:"true"`
}

func (s GetInfrastructureLogListPageVO) String() string {
  return tea.Prettify(s)
}

func (s GetInfrastructureLogListPageVO) GoString() string {
  return s.String()
}

func (s *GetInfrastructureLogListPageVO) SetCurrent(v int64) *GetInfrastructureLogListPageVO {
  s.Current = &v
  return s
}

func (s *GetInfrastructureLogListPageVO) SetSize(v int64) *GetInfrastructureLogListPageVO {
  s.Size = &v
  return s
}

func (s *GetInfrastructureLogListPageVO) SetPages(v int64) *GetInfrastructureLogListPageVO {
  s.Pages = &v
  return s
}

func (s *GetInfrastructureLogListPageVO) SetTotal(v int64) *GetInfrastructureLogListPageVO {
  s.Total = &v
  return s
}

func (s *GetInfrastructureLogListPageVO) SetRecords(v *GetInfrastructureLogListCsecInfrastructureLogVO) *GetInfrastructureLogListPageVO {
  s.Records = v
  return s
}

type GetInfrastructureLogListCsecInfrastructureLogVO struct {
  // {'en':'Time.', 'zh_CN':'时间。'}
  AttackTime *string `json:"attackTime,omitempty" xml:"attackTime,omitempty" require:"true"`
  // {'en':'Client IP.', 'zh_CN':'客户端IP。'}
  ClientIp *string `json:"clientIp,omitempty" xml:"clientIp,omitempty" require:"true"`
  // {'en':'Node IP.', 'zh_CN':'节点IP。'}
  TargetIp *string `json:"targetIp,omitempty" xml:"targetIp,omitempty" require:"true"`
  // {'en':'Policy Name.', 'zh_CN':'策略名称。'}
  PolicyName *string `json:"policyName,omitempty" xml:"policyName,omitempty" require:"true"`
  // {'en':'Rule Name.', 'zh_CN':'规则名称。'}
  RuleName *string `json:"ruleName,omitempty" xml:"ruleName,omitempty" require:"true"`
  // {'en':'Action.', 'zh_CN':'处理动作。'}
  Action *string `json:"action,omitempty" xml:"action,omitempty" require:"true"`
  // {'en':'IP site.', 'zh_CN':'IP地理位置。'}
  IpSite *string `json:"ipSite,omitempty" xml:"ipSite,omitempty" require:"true"`
  // {'en':'Explanation.', 'zh_CN':'异常说明。'}
  Content *string `json:"content,omitempty" xml:"content,omitempty" require:"true"`
  // {'en':'Target port.', 'zh_CN':'目标端口。'}
  TargetPort *string `json:"targetPort,omitempty" xml:"targetPort,omitempty" require:"true"`
}

func (s GetInfrastructureLogListCsecInfrastructureLogVO) String() string {
  return tea.Prettify(s)
}

func (s GetInfrastructureLogListCsecInfrastructureLogVO) GoString() string {
  return s.String()
}

func (s *GetInfrastructureLogListCsecInfrastructureLogVO) SetAttackTime(v string) *GetInfrastructureLogListCsecInfrastructureLogVO {
  s.AttackTime = &v
  return s
}

func (s *GetInfrastructureLogListCsecInfrastructureLogVO) SetClientIp(v string) *GetInfrastructureLogListCsecInfrastructureLogVO {
  s.ClientIp = &v
  return s
}

func (s *GetInfrastructureLogListCsecInfrastructureLogVO) SetTargetIp(v string) *GetInfrastructureLogListCsecInfrastructureLogVO {
  s.TargetIp = &v
  return s
}

func (s *GetInfrastructureLogListCsecInfrastructureLogVO) SetPolicyName(v string) *GetInfrastructureLogListCsecInfrastructureLogVO {
  s.PolicyName = &v
  return s
}

func (s *GetInfrastructureLogListCsecInfrastructureLogVO) SetRuleName(v string) *GetInfrastructureLogListCsecInfrastructureLogVO {
  s.RuleName = &v
  return s
}

func (s *GetInfrastructureLogListCsecInfrastructureLogVO) SetAction(v string) *GetInfrastructureLogListCsecInfrastructureLogVO {
  s.Action = &v
  return s
}

func (s *GetInfrastructureLogListCsecInfrastructureLogVO) SetIpSite(v string) *GetInfrastructureLogListCsecInfrastructureLogVO {
  s.IpSite = &v
  return s
}

func (s *GetInfrastructureLogListCsecInfrastructureLogVO) SetContent(v string) *GetInfrastructureLogListCsecInfrastructureLogVO {
  s.Content = &v
  return s
}

func (s *GetInfrastructureLogListCsecInfrastructureLogVO) SetTargetPort(v string) *GetInfrastructureLogListCsecInfrastructureLogVO {
  s.TargetPort = &v
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

type GetInfrastructureLogListRequestHeader struct {
  // {"zh_CN":"安全服务类型。有使用多个不同的安全服务时，需要填写具体的服务类型。","en":"Security service type. Please enter a specific service type, if you purchase multiple security services.","dictionary":"belong=WAAP-MS-Ext|dict=waap_serviceType"}
  ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty"`
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
  // {'en':'Event ID.', 'zh_CN':'事件ID。'}
  Uuid *string `json:"uuid,omitempty" xml:"uuid,omitempty" require:"true"`
  // {'en':'Request ID.', 'zh_CN':'请求ID。'}
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty" require:"true"`
  // {'en':'Hostname.', 'zh_CN':'域名。'}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {'en':'Policy type. 
  //  DMS_DEFEND: DDoS Protection
  //  BOT_MANAGE: Bot Management
  //  API_DEFEND: API Security
  //  WAF_DEFEND: WAF
  //  BLOCK: IP/Geo Block
  //  CUSTOMIZE_RULE: Custom Rules
  //  RATE_LIMIT: Rate Limiting
  //  INTELLIGENCE: Threat Intelligence', 'zh_CN':'策略类型。
  //  DMS_DEFEND：DDoS防护
  //  BOT_MANAGE：Bot管理
  //  API_DEFEND：API安全
  //  WAF_DEFEND：WAF
  //  BLOCK：IP/区域封禁
  //  CUSTOMIZE_RULE：自定义规则
  //  RATE_LIMIT：频率限制
  //  INTELLIGENCE：威胁情报'}
  PolicyType *string `json:"policyType,omitempty" xml:"policyType,omitempty" require:"true"`
  // {'en':'Attack time, format: yyyy-MM-dd HH:mm:ss.', 'zh_CN':'攻击时间，时间格式：yyyy-MM-dd HH:mm:ss。'}
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

type QueryAttackLogDetailInfoResponse struct {
  // {'en':'Please refer to the error code for exceptions.', 'zh_CN':'请参照错误码。','dictionary':'belong=WAAP-MS-Ext|dict=waap_retCodeEnum'}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {'en':'Description.', 'zh_CN':'描述信息。'}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {'en':'Data.', 'zh_CN':'出参数据。'}
  Data *QueryAttackLogDetailInfoAttackLogAllDetailDto `json:"data,omitempty" xml:"data,omitempty" require:"true"`
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

func (s *QueryAttackLogDetailInfoResponse) SetData(v *QueryAttackLogDetailInfoAttackLogAllDetailDto) *QueryAttackLogDetailInfoResponse {
  s.Data = v
  return s
}

type QueryAttackLogDetailInfoAttackLogDetailDto struct {
  // {'en':'The policy type of hit interception, empty when not intercepted. 
  //  DMS_DEFEND: DDoS Protection
  //  BOT_MANAGE: Bot Management
  //  API_DEFEND: API Security
  //  WAF_DEFEND: WAF
  //  BLOCK: IP/Geo Block
  //  CUSTOMIZE_RULE: Custom Rules
  //  RATE_LIMIT: Rate Limiting
  //  INTELLIGENCE: Threat Intelligence', 'zh_CN':'命中拦截的策略类型，没有被拦截时为空。
  //  DMS_DEFEND：DDoS防护
  //  BOT_MANAGE：Bot管理
  //  API_DEFEND：API安全
  //  WAF_DEFEND：WAF
  //  BLOCK：IP/区域封禁
  //  CUSTOMIZE_RULE：自定义规则
  //  RATE_LIMIT：频率限制
  //  INTELLIGENCE：威胁情报'}
  BlockPolicyName *string `json:"blockPolicyName,omitempty" xml:"blockPolicyName,omitempty" require:"true"`
  // {'en':'IP/Geo Block.', 'zh_CN':'IP区域封禁。'}
  BLOCK *QueryAttackLogDetailInfoAttackLogIpBlockDto `json:"BLOCK,omitempty" xml:"BLOCK,omitempty" require:"true"`
  // {'en':'WAF.', 'zh_CN':'WAF。'}
  WAF_DEFEND []*QueryAttackLogDetailInfoAttackLogWafRuleDto `json:"WAF_DEFEND,omitempty" xml:"WAF_DEFEND,omitempty" require:"true" type:"Repeated"`
  // {'en':'DDoS Protection.', 'zh_CN':'DDos防护。'}
  DMS_DEFEND []*QueryAttackLogDetailInfoAttackLogDdosRuleDto `json:"DMS_DEFEND,omitempty" xml:"DMS_DEFEND,omitempty" require:"true" type:"Repeated"`
  // {'en':'Bot Management.', 'zh_CN':'Bot管理。'}
  BOT_MANAGE []*QueryAttackLogDetailInfoAttackLogBotRuleDto `json:"BOT_MANAGE,omitempty" xml:"BOT_MANAGE,omitempty" require:"true" type:"Repeated"`
  // {'en':'API Security.', 'zh_CN':'API安全。'}
  API_DEFEND []*QueryAttackLogDetailInfoAttackLogApiRuleDto `json:"API_DEFEND,omitempty" xml:"API_DEFEND,omitempty" require:"true" type:"Repeated"`
  // {'en':'Threat Intelligence.', 'zh_CN':'威胁情报。'}
  INTELLIGENCE []*QueryAttackLogDetailInfoAttackLogIntelligentDto `json:"INTELLIGENCE,omitempty" xml:"INTELLIGENCE,omitempty" require:"true" type:"Repeated"`
  // {'en':'Rate Limiting.', 'zh_CN':'频率限制。'}
  RATE_LIMIT []*QueryAttackLogDetailInfoAttackLogRateLimitDto `json:"RATE_LIMIT,omitempty" xml:"RATE_LIMIT,omitempty" require:"true" type:"Repeated"`
  // {'en':'Custom Rules.', 'zh_CN':'自定义规则。'}
  CUSTOMIZE_RULE []*QueryAttackLogDetailInfoAttackLogCustomizeRuleDto `json:"CUSTOMIZE_RULE,omitempty" xml:"CUSTOMIZE_RULE,omitempty" require:"true" type:"Repeated"`
}

func (s QueryAttackLogDetailInfoAttackLogDetailDto) String() string {
  return tea.Prettify(s)
}

func (s QueryAttackLogDetailInfoAttackLogDetailDto) GoString() string {
  return s.String()
}

func (s *QueryAttackLogDetailInfoAttackLogDetailDto) SetBlockPolicyName(v string) *QueryAttackLogDetailInfoAttackLogDetailDto {
  s.BlockPolicyName = &v
  return s
}

func (s *QueryAttackLogDetailInfoAttackLogDetailDto) SetBLOCK(v *QueryAttackLogDetailInfoAttackLogIpBlockDto) *QueryAttackLogDetailInfoAttackLogDetailDto {
  s.BLOCK = v
  return s
}

func (s *QueryAttackLogDetailInfoAttackLogDetailDto) SetWAF_DEFEND(v []*QueryAttackLogDetailInfoAttackLogWafRuleDto) *QueryAttackLogDetailInfoAttackLogDetailDto {
  s.WAF_DEFEND = v
  return s
}

func (s *QueryAttackLogDetailInfoAttackLogDetailDto) SetDMS_DEFEND(v []*QueryAttackLogDetailInfoAttackLogDdosRuleDto) *QueryAttackLogDetailInfoAttackLogDetailDto {
  s.DMS_DEFEND = v
  return s
}

func (s *QueryAttackLogDetailInfoAttackLogDetailDto) SetBOT_MANAGE(v []*QueryAttackLogDetailInfoAttackLogBotRuleDto) *QueryAttackLogDetailInfoAttackLogDetailDto {
  s.BOT_MANAGE = v
  return s
}

func (s *QueryAttackLogDetailInfoAttackLogDetailDto) SetAPI_DEFEND(v []*QueryAttackLogDetailInfoAttackLogApiRuleDto) *QueryAttackLogDetailInfoAttackLogDetailDto {
  s.API_DEFEND = v
  return s
}

func (s *QueryAttackLogDetailInfoAttackLogDetailDto) SetINTELLIGENCE(v []*QueryAttackLogDetailInfoAttackLogIntelligentDto) *QueryAttackLogDetailInfoAttackLogDetailDto {
  s.INTELLIGENCE = v
  return s
}

func (s *QueryAttackLogDetailInfoAttackLogDetailDto) SetRATE_LIMIT(v []*QueryAttackLogDetailInfoAttackLogRateLimitDto) *QueryAttackLogDetailInfoAttackLogDetailDto {
  s.RATE_LIMIT = v
  return s
}

func (s *QueryAttackLogDetailInfoAttackLogDetailDto) SetCUSTOMIZE_RULE(v []*QueryAttackLogDetailInfoAttackLogCustomizeRuleDto) *QueryAttackLogDetailInfoAttackLogDetailDto {
  s.CUSTOMIZE_RULE = v
  return s
}

type QueryAttackLogDetailInfoAttackLogAllDetailDto struct {
  // {'en':'Basic Information.', 'zh_CN':'基础信息。'}
  BasicInfo *QueryAttackLogDetailInfoAttackLogBasicInfoDto `json:"basicInfo,omitempty" xml:"basicInfo,omitempty" require:"true"`
  // {'en':'Details.', 'zh_CN':'详细信息。'}
  DetailInfo *QueryAttackLogDetailInfoAttackLogDetailDto `json:"detailInfo,omitempty" xml:"detailInfo,omitempty" require:"true"`
}

func (s QueryAttackLogDetailInfoAttackLogAllDetailDto) String() string {
  return tea.Prettify(s)
}

func (s QueryAttackLogDetailInfoAttackLogAllDetailDto) GoString() string {
  return s.String()
}

func (s *QueryAttackLogDetailInfoAttackLogAllDetailDto) SetBasicInfo(v *QueryAttackLogDetailInfoAttackLogBasicInfoDto) *QueryAttackLogDetailInfoAttackLogAllDetailDto {
  s.BasicInfo = v
  return s
}

func (s *QueryAttackLogDetailInfoAttackLogAllDetailDto) SetDetailInfo(v *QueryAttackLogDetailInfoAttackLogDetailDto) *QueryAttackLogDetailInfoAttackLogAllDetailDto {
  s.DetailInfo = v
  return s
}

type QueryAttackLogDetailInfoAttackLogWafRuleDto struct {
  // {'en':'Rule type.', 'zh_CN':'规则类型。'}
  RuleType *string `json:"ruleType,omitempty" xml:"ruleType,omitempty" require:"true"`
  // {'en':'Rule action.', 'zh_CN':'规则动作。'}
  RuleAction *string `json:"ruleAction,omitempty" xml:"ruleAction,omitempty" require:"true"`
  // {'en':'Rule ID.', 'zh_CN':'规则ID。'}
  RuleId *string `json:"ruleId,omitempty" xml:"ruleId,omitempty" require:"true"`
  // {'en':'Rule name.', 'zh_CN':'规则名称。'}
  RuleName *string `json:"ruleName,omitempty" xml:"ruleName,omitempty" require:"true"`
  // {'en':'Rule description.', 'zh_CN':'规则描述。'}
  RuleDesc *string `json:"ruleDesc,omitempty" xml:"ruleDesc,omitempty" require:"true"`
  // {'en':'Hit request area.', 'zh_CN':'命中请求区域。'}
  MatchArea *string `json:"matchArea,omitempty" xml:"matchArea,omitempty" require:"true"`
  // {'en':'Hit request content.', 'zh_CN':'命中请求内容。'}
  MatchContent *string `json:"matchContent,omitempty" xml:"matchContent,omitempty" require:"true"`
}

func (s QueryAttackLogDetailInfoAttackLogWafRuleDto) String() string {
  return tea.Prettify(s)
}

func (s QueryAttackLogDetailInfoAttackLogWafRuleDto) GoString() string {
  return s.String()
}

func (s *QueryAttackLogDetailInfoAttackLogWafRuleDto) SetRuleType(v string) *QueryAttackLogDetailInfoAttackLogWafRuleDto {
  s.RuleType = &v
  return s
}

func (s *QueryAttackLogDetailInfoAttackLogWafRuleDto) SetRuleAction(v string) *QueryAttackLogDetailInfoAttackLogWafRuleDto {
  s.RuleAction = &v
  return s
}

func (s *QueryAttackLogDetailInfoAttackLogWafRuleDto) SetRuleId(v string) *QueryAttackLogDetailInfoAttackLogWafRuleDto {
  s.RuleId = &v
  return s
}

func (s *QueryAttackLogDetailInfoAttackLogWafRuleDto) SetRuleName(v string) *QueryAttackLogDetailInfoAttackLogWafRuleDto {
  s.RuleName = &v
  return s
}

func (s *QueryAttackLogDetailInfoAttackLogWafRuleDto) SetRuleDesc(v string) *QueryAttackLogDetailInfoAttackLogWafRuleDto {
  s.RuleDesc = &v
  return s
}

func (s *QueryAttackLogDetailInfoAttackLogWafRuleDto) SetMatchArea(v string) *QueryAttackLogDetailInfoAttackLogWafRuleDto {
  s.MatchArea = &v
  return s
}

func (s *QueryAttackLogDetailInfoAttackLogWafRuleDto) SetMatchContent(v string) *QueryAttackLogDetailInfoAttackLogWafRuleDto {
  s.MatchContent = &v
  return s
}

type QueryAttackLogDetailInfoAttackLogApiRuleDto struct {
  // {'en':'Rule action.', 'zh_CN':'规则动作。'}
  RuleAction *string `json:"ruleAction,omitempty" xml:"ruleAction,omitempty" require:"true"`
  // {'en':'Policy name.', 'zh_CN':'策略名称。'}
  PolicyName *string `json:"policyName,omitempty" xml:"policyName,omitempty" require:"true"`
  // {'en':'API name.', 'zh_CN':'API名称。'}
  ApiName *string `json:"apiName,omitempty" xml:"apiName,omitempty" require:"true"`
  // {'en':'API ID.', 'zh_CN':'API ID。'}
  ApiId *string `json:"apiId,omitempty" xml:"apiId,omitempty" require:"true"`
  // {'en':'English instructions.', 'zh_CN':'英文说明。'}
  ExplainEn *string `json:"explainEn,omitempty" xml:"explainEn,omitempty" require:"true"`
  // {'en':'Chinese instructions, only supported in language=zh_CN', 'zh_CN':'中文说明，仅当language=zh_CN支持。'}
  ExplainCn *string `json:"explainCn,omitempty" xml:"explainCn,omitempty" require:"true"`
}

func (s QueryAttackLogDetailInfoAttackLogApiRuleDto) String() string {
  return tea.Prettify(s)
}

func (s QueryAttackLogDetailInfoAttackLogApiRuleDto) GoString() string {
  return s.String()
}

func (s *QueryAttackLogDetailInfoAttackLogApiRuleDto) SetRuleAction(v string) *QueryAttackLogDetailInfoAttackLogApiRuleDto {
  s.RuleAction = &v
  return s
}

func (s *QueryAttackLogDetailInfoAttackLogApiRuleDto) SetPolicyName(v string) *QueryAttackLogDetailInfoAttackLogApiRuleDto {
  s.PolicyName = &v
  return s
}

func (s *QueryAttackLogDetailInfoAttackLogApiRuleDto) SetApiName(v string) *QueryAttackLogDetailInfoAttackLogApiRuleDto {
  s.ApiName = &v
  return s
}

func (s *QueryAttackLogDetailInfoAttackLogApiRuleDto) SetApiId(v string) *QueryAttackLogDetailInfoAttackLogApiRuleDto {
  s.ApiId = &v
  return s
}

func (s *QueryAttackLogDetailInfoAttackLogApiRuleDto) SetExplainEn(v string) *QueryAttackLogDetailInfoAttackLogApiRuleDto {
  s.ExplainEn = &v
  return s
}

func (s *QueryAttackLogDetailInfoAttackLogApiRuleDto) SetExplainCn(v string) *QueryAttackLogDetailInfoAttackLogApiRuleDto {
  s.ExplainCn = &v
  return s
}

type QueryAttackLogDetailInfoAttackLogIntelligentDto struct {
  // {'en':'Threat type.', 'zh_CN':'情报类型。'}
  IntelligentType *string `json:"intelligentType,omitempty" xml:"intelligentType,omitempty" require:"true"`
  // {'en':'Rule action.', 'zh_CN':'规则动作。'}
  RuleAction *string `json:"ruleAction,omitempty" xml:"ruleAction,omitempty" require:"true"`
  // {'en':'Intelligence module.', 'zh_CN':'情报模块。'}
  IntelligentModule *string `json:"intelligentModule,omitempty" xml:"intelligentModule,omitempty" require:"true"`
}

func (s QueryAttackLogDetailInfoAttackLogIntelligentDto) String() string {
  return tea.Prettify(s)
}

func (s QueryAttackLogDetailInfoAttackLogIntelligentDto) GoString() string {
  return s.String()
}

func (s *QueryAttackLogDetailInfoAttackLogIntelligentDto) SetIntelligentType(v string) *QueryAttackLogDetailInfoAttackLogIntelligentDto {
  s.IntelligentType = &v
  return s
}

func (s *QueryAttackLogDetailInfoAttackLogIntelligentDto) SetRuleAction(v string) *QueryAttackLogDetailInfoAttackLogIntelligentDto {
  s.RuleAction = &v
  return s
}

func (s *QueryAttackLogDetailInfoAttackLogIntelligentDto) SetIntelligentModule(v string) *QueryAttackLogDetailInfoAttackLogIntelligentDto {
  s.IntelligentModule = &v
  return s
}

type QueryAttackLogDetailInfoAttackLogCustomizeRuleDto struct {
  // {'en':'Rule action.', 'zh_CN':'规则动作。'}
  RuleAction *string `json:"ruleAction,omitempty" xml:"ruleAction,omitempty" require:"true"`
  // {'en':'API name.', 'zh_CN':'API名称。'}
  ApiName *string `json:"apiName,omitempty" xml:"apiName,omitempty" require:"true"`
  // {'en':'API ID.', 'zh_CN':'API ID。'}
  ApiId *string `json:"apiId,omitempty" xml:"apiId,omitempty" require:"true"`
  // {'en':'Rule name.', 'zh_CN':'规则名称。'}
  RuleName *string `json:"ruleName,omitempty" xml:"ruleName,omitempty" require:"true"`
  // {'en':'Rule ID.', 'zh_CN':'规则ID。'}
  RuleId *string `json:"ruleId,omitempty" xml:"ruleId,omitempty" require:"true"`
  // {'en':'Rule description.', 'zh_CN':'规则描述。'}
  RuleDesc *string `json:"ruleDesc,omitempty" xml:"ruleDesc,omitempty" require:"true"`
}

func (s QueryAttackLogDetailInfoAttackLogCustomizeRuleDto) String() string {
  return tea.Prettify(s)
}

func (s QueryAttackLogDetailInfoAttackLogCustomizeRuleDto) GoString() string {
  return s.String()
}

func (s *QueryAttackLogDetailInfoAttackLogCustomizeRuleDto) SetRuleAction(v string) *QueryAttackLogDetailInfoAttackLogCustomizeRuleDto {
  s.RuleAction = &v
  return s
}

func (s *QueryAttackLogDetailInfoAttackLogCustomizeRuleDto) SetApiName(v string) *QueryAttackLogDetailInfoAttackLogCustomizeRuleDto {
  s.ApiName = &v
  return s
}

func (s *QueryAttackLogDetailInfoAttackLogCustomizeRuleDto) SetApiId(v string) *QueryAttackLogDetailInfoAttackLogCustomizeRuleDto {
  s.ApiId = &v
  return s
}

func (s *QueryAttackLogDetailInfoAttackLogCustomizeRuleDto) SetRuleName(v string) *QueryAttackLogDetailInfoAttackLogCustomizeRuleDto {
  s.RuleName = &v
  return s
}

func (s *QueryAttackLogDetailInfoAttackLogCustomizeRuleDto) SetRuleId(v string) *QueryAttackLogDetailInfoAttackLogCustomizeRuleDto {
  s.RuleId = &v
  return s
}

func (s *QueryAttackLogDetailInfoAttackLogCustomizeRuleDto) SetRuleDesc(v string) *QueryAttackLogDetailInfoAttackLogCustomizeRuleDto {
  s.RuleDesc = &v
  return s
}

type QueryAttackLogDetailInfoAttackLogBasicInfoDto struct {
  // {'en':'Request ID.', 'zh_CN':'请求ID。'}
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty" require:"true"`
  // {'en':'Event ID.', 'zh_CN':'事件ID。'}
  Uuid *string `json:"uuid,omitempty" xml:"uuid,omitempty" require:"true"`
  // {'en':'Request Method.', 'zh_CN':'请求方法。'}
  RequestMethod *string `json:"requestMethod,omitempty" xml:"requestMethod,omitempty" require:"true"`
  // {'en':'Http Version.', 'zh_CN':'HTTP版本。'}
  HttpVersion *string `json:"httpVersion,omitempty" xml:"httpVersion,omitempty" require:"true"`
  // {'en':'Path.', 'zh_CN':'路径。'}
  Path *string `json:"path,omitempty" xml:"path,omitempty" require:"true"`
  // {'en':'URI.', 'zh_CN':'URI。'}
  Uri *string `json:"uri,omitempty" xml:"uri,omitempty" require:"true"`
  // {'en':'User-Agent.', 'zh_CN':'User-Agent。'}
  UserAgent *string `json:"userAgent,omitempty" xml:"userAgent,omitempty" require:"true"`
  // {'en':'Referer.', 'zh_CN':'Referer。'}
  Referer *string `json:"referer,omitempty" xml:"referer,omitempty" require:"true"`
  // {'en':'Status Code.', 'zh_CN':'状态码。'}
  StatusCode *string `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
  // {'en':'Client IP.', 'zh_CN':'客户端IP。'}
  ClientIp *string `json:"clientIp,omitempty" xml:"clientIp,omitempty" require:"true"`
  // {'en':'IP geolocation.', 'zh_CN':'IP地理位置。'}
  IpLocation *string `json:"ipLocation,omitempty" xml:"ipLocation,omitempty" require:"true"`
  // {'en':'Client ID.', 'zh_CN':'客户端ID。'}
  UserFinger *string `json:"userFinger,omitempty" xml:"userFinger,omitempty" require:"true"`
  // {'en':'Browser fingerprint.', 'zh_CN':'浏览器指纹。'}
  BrowserFinger *string `json:"browserFinger,omitempty" xml:"browserFinger,omitempty" require:"true"`
  // {'en':'Device fingerprint.', 'zh_CN':'设备指纹。'}
  DeviceFinger *string `json:"deviceFinger,omitempty" xml:"deviceFinger,omitempty" require:"true"`
  // {'en':'APP name.', 'zh_CN':'APP名称。'}
  AppName *string `json:"appName,omitempty" xml:"appName,omitempty" require:"true"`
  // {'en':'APP version.', 'zh_CN':'APP版本。'}
  AppVersion *string `json:"appVersion,omitempty" xml:"appVersion,omitempty" require:"true"`
  // {'en':'APP ID.', 'zh_CN':'APP ID。'}
  AppId *string `json:"appId,omitempty" xml:"appId,omitempty" require:"true"`
  // {'en':'APP package name.', 'zh_CN':'APP包名。'}
  AppPackageName *string `json:"appPackageName,omitempty" xml:"appPackageName,omitempty" require:"true"`
  // {'en':'SDK version.', 'zh_CN':'SDK版本。'}
  SdkVersion *string `json:"sdkVersion,omitempty" xml:"sdkVersion,omitempty" require:"true"`
}

func (s QueryAttackLogDetailInfoAttackLogBasicInfoDto) String() string {
  return tea.Prettify(s)
}

func (s QueryAttackLogDetailInfoAttackLogBasicInfoDto) GoString() string {
  return s.String()
}

func (s *QueryAttackLogDetailInfoAttackLogBasicInfoDto) SetRequestId(v string) *QueryAttackLogDetailInfoAttackLogBasicInfoDto {
  s.RequestId = &v
  return s
}

func (s *QueryAttackLogDetailInfoAttackLogBasicInfoDto) SetUuid(v string) *QueryAttackLogDetailInfoAttackLogBasicInfoDto {
  s.Uuid = &v
  return s
}

func (s *QueryAttackLogDetailInfoAttackLogBasicInfoDto) SetRequestMethod(v string) *QueryAttackLogDetailInfoAttackLogBasicInfoDto {
  s.RequestMethod = &v
  return s
}

func (s *QueryAttackLogDetailInfoAttackLogBasicInfoDto) SetHttpVersion(v string) *QueryAttackLogDetailInfoAttackLogBasicInfoDto {
  s.HttpVersion = &v
  return s
}

func (s *QueryAttackLogDetailInfoAttackLogBasicInfoDto) SetPath(v string) *QueryAttackLogDetailInfoAttackLogBasicInfoDto {
  s.Path = &v
  return s
}

func (s *QueryAttackLogDetailInfoAttackLogBasicInfoDto) SetUri(v string) *QueryAttackLogDetailInfoAttackLogBasicInfoDto {
  s.Uri = &v
  return s
}

func (s *QueryAttackLogDetailInfoAttackLogBasicInfoDto) SetUserAgent(v string) *QueryAttackLogDetailInfoAttackLogBasicInfoDto {
  s.UserAgent = &v
  return s
}

func (s *QueryAttackLogDetailInfoAttackLogBasicInfoDto) SetReferer(v string) *QueryAttackLogDetailInfoAttackLogBasicInfoDto {
  s.Referer = &v
  return s
}

func (s *QueryAttackLogDetailInfoAttackLogBasicInfoDto) SetStatusCode(v string) *QueryAttackLogDetailInfoAttackLogBasicInfoDto {
  s.StatusCode = &v
  return s
}

func (s *QueryAttackLogDetailInfoAttackLogBasicInfoDto) SetClientIp(v string) *QueryAttackLogDetailInfoAttackLogBasicInfoDto {
  s.ClientIp = &v
  return s
}

func (s *QueryAttackLogDetailInfoAttackLogBasicInfoDto) SetIpLocation(v string) *QueryAttackLogDetailInfoAttackLogBasicInfoDto {
  s.IpLocation = &v
  return s
}

func (s *QueryAttackLogDetailInfoAttackLogBasicInfoDto) SetUserFinger(v string) *QueryAttackLogDetailInfoAttackLogBasicInfoDto {
  s.UserFinger = &v
  return s
}

func (s *QueryAttackLogDetailInfoAttackLogBasicInfoDto) SetBrowserFinger(v string) *QueryAttackLogDetailInfoAttackLogBasicInfoDto {
  s.BrowserFinger = &v
  return s
}

func (s *QueryAttackLogDetailInfoAttackLogBasicInfoDto) SetDeviceFinger(v string) *QueryAttackLogDetailInfoAttackLogBasicInfoDto {
  s.DeviceFinger = &v
  return s
}

func (s *QueryAttackLogDetailInfoAttackLogBasicInfoDto) SetAppName(v string) *QueryAttackLogDetailInfoAttackLogBasicInfoDto {
  s.AppName = &v
  return s
}

func (s *QueryAttackLogDetailInfoAttackLogBasicInfoDto) SetAppVersion(v string) *QueryAttackLogDetailInfoAttackLogBasicInfoDto {
  s.AppVersion = &v
  return s
}

func (s *QueryAttackLogDetailInfoAttackLogBasicInfoDto) SetAppId(v string) *QueryAttackLogDetailInfoAttackLogBasicInfoDto {
  s.AppId = &v
  return s
}

func (s *QueryAttackLogDetailInfoAttackLogBasicInfoDto) SetAppPackageName(v string) *QueryAttackLogDetailInfoAttackLogBasicInfoDto {
  s.AppPackageName = &v
  return s
}

func (s *QueryAttackLogDetailInfoAttackLogBasicInfoDto) SetSdkVersion(v string) *QueryAttackLogDetailInfoAttackLogBasicInfoDto {
  s.SdkVersion = &v
  return s
}

type QueryAttackLogDetailInfoAttackLogIpBlockDto struct {
  // {'en':'Policy name.', 'zh_CN':'策略名称。'}
  PolicyName *string `json:"policyName,omitempty" xml:"policyName,omitempty" require:"true"`
  // {'en':'Rule action.', 'zh_CN':'规则动作。'}
  RuleAction *string `json:"ruleAction,omitempty" xml:"ruleAction,omitempty" require:"true"`
  // {'en':'Description.', 'zh_CN':'说明。'}
  Explain *string `json:"explain,omitempty" xml:"explain,omitempty" require:"true"`
}

func (s QueryAttackLogDetailInfoAttackLogIpBlockDto) String() string {
  return tea.Prettify(s)
}

func (s QueryAttackLogDetailInfoAttackLogIpBlockDto) GoString() string {
  return s.String()
}

func (s *QueryAttackLogDetailInfoAttackLogIpBlockDto) SetPolicyName(v string) *QueryAttackLogDetailInfoAttackLogIpBlockDto {
  s.PolicyName = &v
  return s
}

func (s *QueryAttackLogDetailInfoAttackLogIpBlockDto) SetRuleAction(v string) *QueryAttackLogDetailInfoAttackLogIpBlockDto {
  s.RuleAction = &v
  return s
}

func (s *QueryAttackLogDetailInfoAttackLogIpBlockDto) SetExplain(v string) *QueryAttackLogDetailInfoAttackLogIpBlockDto {
  s.Explain = &v
  return s
}

type QueryAttackLogDetailInfoAttackLogDdosRuleDto struct {
  // {'en':'Policy name.', 'zh_CN':'策略名称。'}
  PolicyName *string `json:"policyName,omitempty" xml:"policyName,omitempty" require:"true"`
  // {'en':'Rule action.', 'zh_CN':'规则动作。'}
  RuleAction *string `json:"ruleAction,omitempty" xml:"ruleAction,omitempty" require:"true"`
  // {'en':'Rule ID.', 'zh_CN':'规则ID。'}
  RuleId *string `json:"ruleId,omitempty" xml:"ruleId,omitempty" require:"true"`
  // {'en':'Rule name.', 'zh_CN':'规则名称。'}
  RuleName *string `json:"ruleName,omitempty" xml:"ruleName,omitempty" require:"true"`
  // {'en':'Description.', 'zh_CN':'说明。'}
  Explain *string `json:"explain,omitempty" xml:"explain,omitempty" require:"true"`
}

func (s QueryAttackLogDetailInfoAttackLogDdosRuleDto) String() string {
  return tea.Prettify(s)
}

func (s QueryAttackLogDetailInfoAttackLogDdosRuleDto) GoString() string {
  return s.String()
}

func (s *QueryAttackLogDetailInfoAttackLogDdosRuleDto) SetPolicyName(v string) *QueryAttackLogDetailInfoAttackLogDdosRuleDto {
  s.PolicyName = &v
  return s
}

func (s *QueryAttackLogDetailInfoAttackLogDdosRuleDto) SetRuleAction(v string) *QueryAttackLogDetailInfoAttackLogDdosRuleDto {
  s.RuleAction = &v
  return s
}

func (s *QueryAttackLogDetailInfoAttackLogDdosRuleDto) SetRuleId(v string) *QueryAttackLogDetailInfoAttackLogDdosRuleDto {
  s.RuleId = &v
  return s
}

func (s *QueryAttackLogDetailInfoAttackLogDdosRuleDto) SetRuleName(v string) *QueryAttackLogDetailInfoAttackLogDdosRuleDto {
  s.RuleName = &v
  return s
}

func (s *QueryAttackLogDetailInfoAttackLogDdosRuleDto) SetExplain(v string) *QueryAttackLogDetailInfoAttackLogDdosRuleDto {
  s.Explain = &v
  return s
}

type QueryAttackLogDetailInfoAttackLogBotRuleDto struct {
  // {'en':'Policy name.', 'zh_CN':'策略名称。'}
  PolicyName *string `json:"policyName,omitempty" xml:"policyName,omitempty" require:"true"`
  // {'en':'Rule action.', 'zh_CN':'规则动作。'}
  RuleAction *string `json:"ruleAction,omitempty" xml:"ruleAction,omitempty" require:"true"`
  // {'en':'Rule name.', 'zh_CN':'规则名称。'}
  RuleName *string `json:"ruleName,omitempty" xml:"ruleName,omitempty" require:"true"`
  // {'en':'Bot Category.', 'zh_CN':'Bot分类。'}
  BotCategory *string `json:"botCategory,omitempty" xml:"botCategory,omitempty" require:"true"`
  // {'en':'Bot tag.', 'zh_CN':'Bot标签。'}
  BotLabel *string `json:"botLabel,omitempty" xml:"botLabel,omitempty" require:"true"`
  // {'en':'Event description.', 'zh_CN':'事件描述。'}
  EventDesc *string `json:"eventDesc,omitempty" xml:"eventDesc,omitempty" require:"true"`
}

func (s QueryAttackLogDetailInfoAttackLogBotRuleDto) String() string {
  return tea.Prettify(s)
}

func (s QueryAttackLogDetailInfoAttackLogBotRuleDto) GoString() string {
  return s.String()
}

func (s *QueryAttackLogDetailInfoAttackLogBotRuleDto) SetPolicyName(v string) *QueryAttackLogDetailInfoAttackLogBotRuleDto {
  s.PolicyName = &v
  return s
}

func (s *QueryAttackLogDetailInfoAttackLogBotRuleDto) SetRuleAction(v string) *QueryAttackLogDetailInfoAttackLogBotRuleDto {
  s.RuleAction = &v
  return s
}

func (s *QueryAttackLogDetailInfoAttackLogBotRuleDto) SetRuleName(v string) *QueryAttackLogDetailInfoAttackLogBotRuleDto {
  s.RuleName = &v
  return s
}

func (s *QueryAttackLogDetailInfoAttackLogBotRuleDto) SetBotCategory(v string) *QueryAttackLogDetailInfoAttackLogBotRuleDto {
  s.BotCategory = &v
  return s
}

func (s *QueryAttackLogDetailInfoAttackLogBotRuleDto) SetBotLabel(v string) *QueryAttackLogDetailInfoAttackLogBotRuleDto {
  s.BotLabel = &v
  return s
}

func (s *QueryAttackLogDetailInfoAttackLogBotRuleDto) SetEventDesc(v string) *QueryAttackLogDetailInfoAttackLogBotRuleDto {
  s.EventDesc = &v
  return s
}

type QueryAttackLogDetailInfoAttackLogRateLimitDto struct {
  // {'en':'Rule action.', 'zh_CN':'规则动作。'}
  RuleAction *string `json:"ruleAction,omitempty" xml:"ruleAction,omitempty" require:"true"`
  // {'en':'API name.', 'zh_CN':'API名称。'}
  ApiName *string `json:"apiName,omitempty" xml:"apiName,omitempty" require:"true"`
  // {'en':'API ID.', 'zh_CN':'API ID。'}
  ApiId *string `json:"apiId,omitempty" xml:"apiId,omitempty" require:"true"`
  // {'en':'Rule name.', 'zh_CN':'规则名称。'}
  RuleName *string `json:"ruleName,omitempty" xml:"ruleName,omitempty" require:"true"`
  // {'en':'Rule ID.', 'zh_CN':'规则ID。'}
  RuleId *string `json:"ruleId,omitempty" xml:"ruleId,omitempty" require:"true"`
  // {'en':'Rule description.', 'zh_CN':'规则描述。'}
  RuleDesc *string `json:"ruleDesc,omitempty" xml:"ruleDesc,omitempty" require:"true"`
}

func (s QueryAttackLogDetailInfoAttackLogRateLimitDto) String() string {
  return tea.Prettify(s)
}

func (s QueryAttackLogDetailInfoAttackLogRateLimitDto) GoString() string {
  return s.String()
}

func (s *QueryAttackLogDetailInfoAttackLogRateLimitDto) SetRuleAction(v string) *QueryAttackLogDetailInfoAttackLogRateLimitDto {
  s.RuleAction = &v
  return s
}

func (s *QueryAttackLogDetailInfoAttackLogRateLimitDto) SetApiName(v string) *QueryAttackLogDetailInfoAttackLogRateLimitDto {
  s.ApiName = &v
  return s
}

func (s *QueryAttackLogDetailInfoAttackLogRateLimitDto) SetApiId(v string) *QueryAttackLogDetailInfoAttackLogRateLimitDto {
  s.ApiId = &v
  return s
}

func (s *QueryAttackLogDetailInfoAttackLogRateLimitDto) SetRuleName(v string) *QueryAttackLogDetailInfoAttackLogRateLimitDto {
  s.RuleName = &v
  return s
}

func (s *QueryAttackLogDetailInfoAttackLogRateLimitDto) SetRuleId(v string) *QueryAttackLogDetailInfoAttackLogRateLimitDto {
  s.RuleId = &v
  return s
}

func (s *QueryAttackLogDetailInfoAttackLogRateLimitDto) SetRuleDesc(v string) *QueryAttackLogDetailInfoAttackLogRateLimitDto {
  s.RuleDesc = &v
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

type QueryAttackLogDetailInfoRequestHeader struct {
  // {'en':'The language of response data, default value: en.
  // zh_CN: Chinese
  // en: English', 'zh_CN':'返回内容的语言版本，默认值: en。
  // zh_CN：中文
  // en：英文'}
  Language *string `json:"language,omitempty" xml:"language,omitempty"`
  // {"zh_CN":"安全服务类型。有使用多个不同的安全服务时，需要填写具体的服务类型。","en":"Security service type. Please enter a specific service type, if you purchase multiple security services.","dictionary":"belong=WAAP-MS-Ext|dict=waap_serviceType"}
  ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty"`
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




