package cloudtrail

// This file is auto-generated, don't edit it. Thanks.
import (
  "github.com/alibabacloud-go/tea/tea"
)

type QueryAccountTrailLogRequest struct {
  // {"en":"Start Time:\n\n1.The Time format is yyyy-MM-ddTHH:mm:ss±HH:mm. Please note: ±HH:mm is the time zone offset, which can be adjusted according to your data needs, for example, +00:00 represents UTC time, +08:00 represents East 8th District, and -05:00 represents West 5th District. 2024-01-15T10:30:45+00:00 means UTC time January 15, 2024 10:30:45 AM\n\n2.Cannot exceed the current time\n\n3.Up to the past 1 year of data can be obtained","zh_CN":"开始时间：\n\n1.时间格式为 yyyy-MM-ddTHH:mm:ss±HH:mm。请注意：±HH:mm 为时区偏移量，如 +00:00 代表 UTC 时间，+08:00 代表东八区，2024-01-15T10:30:45+00:00，表示UTC 时间 2024 年 1 月 15 日上午 10 点 30 分 45 秒\n\n2.不能大于当前时间\n\n3.最多可获取最近1年的数据"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty" require:"true"`
  // {"en":"End time:\n\n1. The time format is yyyy-MM-ddTHH:mm:ss±HH:mm. Please note: ±HH:mm is the time zone offset, which can be adjusted according to your data needs, for example, +00:00 represents UTC time, +08:00 represents East 8th District, and -05:00 represents West 5th District. 2024-01-15T10:30:45+00:00 means UTC time January 15, 2024 10:30:45 AM\n\n2. The end time is greater than the start time\n\n3. If the end time is greater than the current time, the current time is taken\n\n4. DateFrom and dateTo are not uploaded, defaulting to query the past 24 hours; if only one is not uploaded, throw an exception\n\n5. Maximum query interval allowed: 7 days, that is, the difference between dateFrom and dateTo can not exceed 7 days","zh_CN":"结束时间：  \n1.时间格式为 yyyy-MM-ddTHH:mm:ss±HH:mm。请注意：±HH:mm 为时区偏移量，可根据您的数据需要进行调整，例如 +00:00 代表 UTC 时间，+08:00 代表东八区，-05:00 代表西五区。2024-01-15T10:30:45+00:00，表示UTC 时间 2024 年 1 月 15 日上午 10 点 30 分 45 秒  \n2.结束时间需大于开始时间  \n3.结束时间如果大于当前时间，取当前时间  \n4.dateFrom，dateTo二者都未传，默认查询过去的1小时；如仅有一个未传，抛异常  \n5.允许查询最大间隔：7天，即dateFrom和dateTo相差不能超过7天"}
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

type QueryAccountTrailLogRequestHeader struct {
}

func (s QueryAccountTrailLogRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryAccountTrailLogRequestHeader) GoString() string {
  return s.String()
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

type QueryAccountTrailLogResponse struct {
  // {"en":"API response code","zh_CN":"API响应码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"API response message","zh_CN":"API响应消息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Audit log query results","zh_CN":"审计日志查询结果"}
  Data []*QueryAccountTrailLogResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
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

func (s *QueryAccountTrailLogResponse) SetData(v []*QueryAccountTrailLogResponseData) *QueryAccountTrailLogResponse {
  s.Data = v
  return s
}

type QueryAccountTrailLogResponseData struct     {
  // {"en":"Event code.","zh_CN":"事件code"}
  EventCode *string `json:"eventCode,omitempty" xml:"eventCode,omitempty" require:"true"`
  // {"en":"Event Chinese Name.","zh_CN":"事件中文名称"}
  EventCnName *string `json:"eventCnName,omitempty" xml:"eventCnName,omitempty" require:"true"`
  // {"en":"Event English Name.","zh_CN":"事件英文名称"}
  EventEnName *string `json:"eventEnName,omitempty" xml:"eventEnName,omitempty" require:"true"`
  // {"en":"Event Source.","zh_CN":"事件来源"}
  EventSource *string `json:"eventSource,omitempty" xml:"eventSource,omitempty" require:"true"`
  // {"en":"Event time.","zh_CN":"事件发生时间点"}
  EventTime *string `json:"eventTime,omitempty" xml:"eventTime,omitempty" require:"true"`
  // {"en":"Request ID.","zh_CN":"请求ID"}
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty" require:"true"`
  // {"en":"The list of resources impacted by the event.","zh_CN":"事件影响的资源列表。"}
  ReferencedResources []*QueryAccountTrailLogResponseDataReferencedResources `json:"referencedResources,omitempty" xml:"referencedResources,omitempty" require:"true" type:"Repeated"`
  // {"en":"The product code related to the event.","zh_CN":"事件相关的产品code。"}
  ProductCode *string `json:"productCode,omitempty" xml:"productCode,omitempty" require:"true"`
  // {"en":"The source IP address of the initiating event.","zh_CN":"事件发起的源IP。"}
  SourceIpAddress *string `json:"sourceIpAddress,omitempty" xml:"sourceIpAddress,omitempty" require:"true"`
  // {"en":"The client proxy identity for sending API requests.","zh_CN":"发送API请求的客户端代理标识。"}
  UserAgent *string `json:"userAgent,omitempty" xml:"userAgent,omitempty" require:"true"`
  // {"en":"The identity information of the requester.","zh_CN":"请求者的身份信息。"}
  UserIdentity *QueryAccountTrailLogResponseDataUserIdentity `json:"userIdentity,omitempty" xml:"userIdentity,omitempty" require:"true" type:"Struct"`
}

func (s QueryAccountTrailLogResponseData) String() string {
  return tea.Prettify(s)
}

func (s QueryAccountTrailLogResponseData) GoString() string {
  return s.String()
}

func (s *QueryAccountTrailLogResponseData) SetEventCode(v string) *QueryAccountTrailLogResponseData {
  s.EventCode = &v
  return s
}

func (s *QueryAccountTrailLogResponseData) SetEventCnName(v string) *QueryAccountTrailLogResponseData {
  s.EventCnName = &v
  return s
}

func (s *QueryAccountTrailLogResponseData) SetEventEnName(v string) *QueryAccountTrailLogResponseData {
  s.EventEnName = &v
  return s
}

func (s *QueryAccountTrailLogResponseData) SetEventSource(v string) *QueryAccountTrailLogResponseData {
  s.EventSource = &v
  return s
}

func (s *QueryAccountTrailLogResponseData) SetEventTime(v string) *QueryAccountTrailLogResponseData {
  s.EventTime = &v
  return s
}

func (s *QueryAccountTrailLogResponseData) SetRequestId(v string) *QueryAccountTrailLogResponseData {
  s.RequestId = &v
  return s
}

func (s *QueryAccountTrailLogResponseData) SetReferencedResources(v []*QueryAccountTrailLogResponseDataReferencedResources) *QueryAccountTrailLogResponseData {
  s.ReferencedResources = v
  return s
}

func (s *QueryAccountTrailLogResponseData) SetProductCode(v string) *QueryAccountTrailLogResponseData {
  s.ProductCode = &v
  return s
}

func (s *QueryAccountTrailLogResponseData) SetSourceIpAddress(v string) *QueryAccountTrailLogResponseData {
  s.SourceIpAddress = &v
  return s
}

func (s *QueryAccountTrailLogResponseData) SetUserAgent(v string) *QueryAccountTrailLogResponseData {
  s.UserAgent = &v
  return s
}

func (s *QueryAccountTrailLogResponseData) SetUserIdentity(v *QueryAccountTrailLogResponseDataUserIdentity) *QueryAccountTrailLogResponseData {
  s.UserIdentity = v
  return s
}

type QueryAccountTrailLogResponseDataReferencedResources struct     {
  // {"en":"Resource Type.","zh_CN":"资源类型"}
  ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty" require:"true"`
  // {"en":"Resource Name.","zh_CN":"资源名称"}
  ResourceNames []*string `json:"resourceNames,omitempty" xml:"resourceNames,omitempty" require:"true" type:"Repeated"`
}

func (s QueryAccountTrailLogResponseDataReferencedResources) String() string {
  return tea.Prettify(s)
}

func (s QueryAccountTrailLogResponseDataReferencedResources) GoString() string {
  return s.String()
}

func (s *QueryAccountTrailLogResponseDataReferencedResources) SetResourceType(v string) *QueryAccountTrailLogResponseDataReferencedResources {
  s.ResourceType = &v
  return s
}

func (s *QueryAccountTrailLogResponseDataReferencedResources) SetResourceNames(v []*string) *QueryAccountTrailLogResponseDataReferencedResources {
  s.ResourceNames = v
  return s
}

type QueryAccountTrailLogResponseDataUserIdentity struct {
  // {"en":"Account login name.","zh_CN":"账号登录名"}
  LoginName *string `json:"loginName,omitempty" xml:"loginName,omitempty" require:"true"`
  // {"en":"IDP internal account.","zh_CN":"idp内部切换账号"}
  IdpLoginName []*string `json:"idpLoginName,omitempty" xml:"idpLoginName,omitempty" require:"true" type:"Repeated"`
  // {"en":"IDP name.","zh_CN":"idp名称"}
  ClientName []*string `json:"clientName,omitempty" xml:"clientName,omitempty" require:"true" type:"Repeated"`
}

func (s QueryAccountTrailLogResponseDataUserIdentity) String() string {
  return tea.Prettify(s)
}

func (s QueryAccountTrailLogResponseDataUserIdentity) GoString() string {
  return s.String()
}

func (s *QueryAccountTrailLogResponseDataUserIdentity) SetLoginName(v string) *QueryAccountTrailLogResponseDataUserIdentity {
  s.LoginName = &v
  return s
}

func (s *QueryAccountTrailLogResponseDataUserIdentity) SetIdpLoginName(v []*string) *QueryAccountTrailLogResponseDataUserIdentity {
  s.IdpLoginName = v
  return s
}

func (s *QueryAccountTrailLogResponseDataUserIdentity) SetClientName(v []*string) *QueryAccountTrailLogResponseDataUserIdentity {
  s.ClientName = v
  return s
}

type QueryAccountTrailLogResponseHeader struct {
}

func (s QueryAccountTrailLogResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryAccountTrailLogResponseHeader) GoString() string {
  return s.String()
}




