package cloudtrail

// This file is auto-generated, don't edit it. Thanks.
import (
  "github.com/alibabacloud-go/tea/tea"
)

type QueryAccountTrailLogRequest struct {
  // {"en":"Start time:
  // 1. The format is yyyyy-MM-ddTHH: mm: SS + 08:00, for example, 2023-12-20T10:00 + 08:00 (10:0:00 Beijing time on December 20, 2016);
  // 2. Can not exceed the current time;
  // 3. The latest one year data can be obtained at most.", "zh_CN":"开始时间：
  // 1.格式为yyyy-MM-ddTHH:mm:ss+08:00，例如，2023-12-20T10:00:00+08:00（为北京时间2023年12月20日10点0分0秒）；
  // 2.不能大于当前时间；
  // 3.最多可获取最近1年 的数据。"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty" require:"true"`
  // {"en":"End time:
  // 1. The 1format is yyyy-MM-ddTHH:mm:ss+08:00;
  // 2. The end time is greater than the start time.
  // 3. If the end time is greater than the current time, the current time is taken.
  // 4.  Maximum query interval allowed: 7 days, that is, the difference between dateFrom and dateTo can not exceed 7 days. ", "zh_CN":"	结束时间：
  // 1.格式为yyyy-MM-ddTHH:mm:ss+08:00；
  // 2.结束时间需大于开始时间；
  // 3.结束时间如果大于当前时间，取当前时间；
  // 4.允许查询最大间隔：7天，即dateFrom和dateTo相差不能超过7天。"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty" require:"true"`
}

func (s QueryAccountTrailLogRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryAccountTrailLogRequest) GoString() string {
  return s.String()
}

func (s *QueryAccountTrailLogRequest) SetDateFrom(v string) *QueryAccountTrailLogRequest {
  s.DateFrom = &v
  return s
}

func (s *QueryAccountTrailLogRequest) SetDateTo(v string) *QueryAccountTrailLogRequest {
  s.DateTo = &v
  return s
}

type QueryAccountTrailLogReferencedResourceView struct {
  // {"en":"Resource Type.", "zh_CN":"资源类型"}
  ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty" require:"true"`
  // {"en":"Resource Name.", "zh_CN":"资源名称"}
  ResourceNames []*string `json:"resourceNames,omitempty" xml:"resourceNames,omitempty" require:"true" type:"Repeated"`
}

func (s QueryAccountTrailLogReferencedResourceView) String() string {
  return tea.Prettify(s)
}

func (s QueryAccountTrailLogReferencedResourceView) GoString() string {
  return s.String()
}

func (s *QueryAccountTrailLogReferencedResourceView) SetResourceType(v string) *QueryAccountTrailLogReferencedResourceView {
  s.ResourceType = &v
  return s
}

func (s *QueryAccountTrailLogReferencedResourceView) SetResourceNames(v []*string) *QueryAccountTrailLogReferencedResourceView {
  s.ResourceNames = v
  return s
}

type QueryAccountTrailLogUserIdentityView struct {
  // {"en":"Account login name.", "zh_CN":"账号登录名"}
  LoginName *string `json:"loginName,omitempty" xml:"loginName,omitempty" require:"true"`
  // {"en":"IDP internal account.", "zh_CN":"idp内部切换账号"}
  IdpLoginName []*string `json:"idpLoginName,omitempty" xml:"idpLoginName,omitempty" require:"true" type:"Repeated"`
  // {"en":"IDP name.", "zh_CN":"	idp名称"}
  ClientName []*string `json:"clientName,omitempty" xml:"clientName,omitempty" require:"true" type:"Repeated"`
}

func (s QueryAccountTrailLogUserIdentityView) String() string {
  return tea.Prettify(s)
}

func (s QueryAccountTrailLogUserIdentityView) GoString() string {
  return s.String()
}

func (s *QueryAccountTrailLogUserIdentityView) SetLoginName(v string) *QueryAccountTrailLogUserIdentityView {
  s.LoginName = &v
  return s
}

func (s *QueryAccountTrailLogUserIdentityView) SetIdpLoginName(v []*string) *QueryAccountTrailLogUserIdentityView {
  s.IdpLoginName = v
  return s
}

func (s *QueryAccountTrailLogUserIdentityView) SetClientName(v []*string) *QueryAccountTrailLogUserIdentityView {
  s.ClientName = v
  return s
}

type QueryAccountTrailLogResultView struct {
  // {"en":"Event code.", "zh_CN":"事件code"}
  EventCode *string `json:"eventCode,omitempty" xml:"eventCode,omitempty" require:"true"`
  // {"en":"Event Chinese Name.", "zh_CN":"事件中文名称"}
  EventCnName *string `json:"eventCnName,omitempty" xml:"eventCnName,omitempty" require:"true"`
  // {"en":"Event English Name.", "zh_CN":"事件英文名称"}
  EventEnName *string `json:"eventEnName,omitempty" xml:"eventEnName,omitempty" require:"true"`
  // {"en":"Event Source.", "zh_CN":"事件来源"}
  EventSource *string `json:"eventSource,omitempty" xml:"eventSource,omitempty" require:"true"`
  // {"en":"Event time.", "zh_CN":"事件发生时间点"}
  EventTime *string `json:"eventTime,omitempty" xml:"eventTime,omitempty" require:"true"`
  // {"en":"Request ID.", "zh_CN":"请求ID"}
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty" require:"true"`
  // {"en":"The list of resources impacted by the event.", "zh_CN":"事件影响的资源列表。"}
  ReferencedResources []*QueryAccountTrailLogReferencedResourceView `json:"referencedResources,omitempty" xml:"referencedResources,omitempty" require:"true" type:"Repeated"`
  // {"en":"The product code related to the event.", "zh_CN":"事件相关的产品code。"}
  ProductCode *string `json:"productCode,omitempty" xml:"productCode,omitempty" require:"true"`
  // {"en":"The source IP address of the initiating event.", "zh_CN":"事件发起的源IP。"}
  SourceIpAddress *string `json:"sourceIpAddress,omitempty" xml:"sourceIpAddress,omitempty" require:"true"`
  // {"en":"The client proxy identity for sending API requests.", "zh_CN":"发送API请求的客户端代理标识。"}
  UserAgent *string `json:"userAgent,omitempty" xml:"userAgent,omitempty" require:"true"`
  // {"en":"The identity information of the requester.", "zh_CN":"请求者的身份信息。"}
  UserIdentity *QueryAccountTrailLogUserIdentityView `json:"userIdentity,omitempty" xml:"userIdentity,omitempty" require:"true"`
}

func (s QueryAccountTrailLogResultView) String() string {
  return tea.Prettify(s)
}

func (s QueryAccountTrailLogResultView) GoString() string {
  return s.String()
}

func (s *QueryAccountTrailLogResultView) SetEventCode(v string) *QueryAccountTrailLogResultView {
  s.EventCode = &v
  return s
}

func (s *QueryAccountTrailLogResultView) SetEventCnName(v string) *QueryAccountTrailLogResultView {
  s.EventCnName = &v
  return s
}

func (s *QueryAccountTrailLogResultView) SetEventEnName(v string) *QueryAccountTrailLogResultView {
  s.EventEnName = &v
  return s
}

func (s *QueryAccountTrailLogResultView) SetEventSource(v string) *QueryAccountTrailLogResultView {
  s.EventSource = &v
  return s
}

func (s *QueryAccountTrailLogResultView) SetEventTime(v string) *QueryAccountTrailLogResultView {
  s.EventTime = &v
  return s
}

func (s *QueryAccountTrailLogResultView) SetRequestId(v string) *QueryAccountTrailLogResultView {
  s.RequestId = &v
  return s
}

func (s *QueryAccountTrailLogResultView) SetReferencedResources(v []*QueryAccountTrailLogReferencedResourceView) *QueryAccountTrailLogResultView {
  s.ReferencedResources = v
  return s
}

func (s *QueryAccountTrailLogResultView) SetProductCode(v string) *QueryAccountTrailLogResultView {
  s.ProductCode = &v
  return s
}

func (s *QueryAccountTrailLogResultView) SetSourceIpAddress(v string) *QueryAccountTrailLogResultView {
  s.SourceIpAddress = &v
  return s
}

func (s *QueryAccountTrailLogResultView) SetUserAgent(v string) *QueryAccountTrailLogResultView {
  s.UserAgent = &v
  return s
}

func (s *QueryAccountTrailLogResultView) SetUserIdentity(v *QueryAccountTrailLogUserIdentityView) *QueryAccountTrailLogResultView {
  s.UserIdentity = v
  return s
}

type QueryAccountTrailLogResponse struct {
  // {"en":"API response code", "zh_CN":"API响应码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"API response message", "zh_CN":"API响应消息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Audit log query results", "zh_CN":"审计日志查询结果"}
  Data []*QueryAccountTrailLogResultView `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s QueryAccountTrailLogResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryAccountTrailLogResponse) GoString() string {
  return s.String()
}

func (s *QueryAccountTrailLogResponse) SetCode(v string) *QueryAccountTrailLogResponse {
  s.Code = &v
  return s
}

func (s *QueryAccountTrailLogResponse) SetMessage(v string) *QueryAccountTrailLogResponse {
  s.Message = &v
  return s
}

func (s *QueryAccountTrailLogResponse) SetData(v []*QueryAccountTrailLogResultView) *QueryAccountTrailLogResponse {
  s.Data = v
  return s
}

type QueryAccountTrailLogPaths struct {
}

func (s QueryAccountTrailLogPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryAccountTrailLogPaths) GoString() string {
  return s.String()
}

type QueryAccountTrailLogParameters struct {
}

func (s QueryAccountTrailLogParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryAccountTrailLogParameters) GoString() string {
  return s.String()
}

type QueryAccountTrailLogRequestHeader struct {
}

func (s QueryAccountTrailLogRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryAccountTrailLogRequestHeader) GoString() string {
  return s.String()
}

type QueryAccountTrailLogResponseHeader struct {
}

func (s QueryAccountTrailLogResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryAccountTrailLogResponseHeader) GoString() string {
  return s.String()
}




