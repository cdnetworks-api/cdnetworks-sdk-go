package customeralarm

// This file is auto-generated, don't edit it. Thanks.
import (
  "github.com/alibabacloud-go/tea/tea"
)

type QueryInstanceDowntimeNotificationRequest struct {
}

func (s QueryInstanceDowntimeNotificationRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryInstanceDowntimeNotificationRequest) GoString() string {
  return s.String()
}

type QueryInstanceDowntimeNotificationResponse struct {
  Events []*QueryInstanceDowntimeNotificationEvent `json:"events,omitempty" xml:"events,omitempty" require:"true" type:"Repeated"`
}

func (s QueryInstanceDowntimeNotificationResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryInstanceDowntimeNotificationResponse) GoString() string {
  return s.String()
}

func (s *QueryInstanceDowntimeNotificationResponse) SetEvents(v []*QueryInstanceDowntimeNotificationEvent) *QueryInstanceDowntimeNotificationResponse {
  s.Events = v
  return s
}

type QueryInstanceDowntimeNotificationEvent struct {
  // {"en":"node name", "zh_CN":"节点英文名称"}
  NodeName *string `json:"nodeName,omitempty" xml:"nodeName,omitempty" require:"true"`
  // {"en":"event name", "zh_CN":"事件名称"}
  QueryInstanceDowntimeNotificationEvent *string `json:"event,omitempty" xml:"event,omitempty" require:"true"`
  // {"en":"instnace fault time", "zh_CN":"宕机时间"}
  Time *string `json:"time,omitempty" xml:"time,omitempty" require:"true"`
  Servers []*QueryInstanceDowntimeNotificationServer `json:"servers,omitempty" xml:"servers,omitempty" require:"true" type:"Repeated"`
}

func (s QueryInstanceDowntimeNotificationEvent) String() string {
  return tea.Prettify(s)
}

func (s QueryInstanceDowntimeNotificationEvent) GoString() string {
  return s.String()
}

func (s *QueryInstanceDowntimeNotificationEvent) SetNodeName(v string) *QueryInstanceDowntimeNotificationEvent {
  s.NodeName = &v
  return s
}

func (s *QueryInstanceDowntimeNotificationEvent) SetEvent(v string) *QueryInstanceDowntimeNotificationEvent {
  s.QueryInstanceDowntimeNotificationEvent = &v
  return s
}

func (s *QueryInstanceDowntimeNotificationEvent) SetTime(v string) *QueryInstanceDowntimeNotificationEvent {
  s.Time = &v
  return s
}

func (s *QueryInstanceDowntimeNotificationEvent) SetServers(v []*QueryInstanceDowntimeNotificationServer) *QueryInstanceDowntimeNotificationEvent {
  s.Servers = v
  return s
}

type QueryInstanceDowntimeNotificationServer struct {
  // {"en":"instance id", "zh_CN":"云主机id"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {"en":"instance name", "zh_CN":"云主机名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"ipv4 array", "zh_CN":"ipv4 集合"}
  Ipv4 []*string `json:"ipv4,omitempty" xml:"ipv4,omitempty" require:"true" type:"Repeated"`
  // {"en":"ipv6 array", "zh_CN":"ipv6 集合"}
  Ipv6 []*string `json:"ipv6,omitempty" xml:"ipv6,omitempty" require:"true" type:"Repeated"`
}

func (s QueryInstanceDowntimeNotificationServer) String() string {
  return tea.Prettify(s)
}

func (s QueryInstanceDowntimeNotificationServer) GoString() string {
  return s.String()
}

func (s *QueryInstanceDowntimeNotificationServer) SetId(v string) *QueryInstanceDowntimeNotificationServer {
  s.Id = &v
  return s
}

func (s *QueryInstanceDowntimeNotificationServer) SetName(v string) *QueryInstanceDowntimeNotificationServer {
  s.Name = &v
  return s
}

func (s *QueryInstanceDowntimeNotificationServer) SetIpv4(v []*string) *QueryInstanceDowntimeNotificationServer {
  s.Ipv4 = v
  return s
}

func (s *QueryInstanceDowntimeNotificationServer) SetIpv6(v []*string) *QueryInstanceDowntimeNotificationServer {
  s.Ipv6 = v
  return s
}

type QueryInstanceDowntimeNotificationPaths struct {
}

func (s QueryInstanceDowntimeNotificationPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryInstanceDowntimeNotificationPaths) GoString() string {
  return s.String()
}

type QueryInstanceDowntimeNotificationParameters struct {
  // {"en":"node name", "zh_CN":"节点英文名称，多个逗号隔开"}
  NodeName *string `json:"nodeName,omitempty" xml:"nodeName,omitempty"`
}

func (s QueryInstanceDowntimeNotificationParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryInstanceDowntimeNotificationParameters) GoString() string {
  return s.String()
}

func (s *QueryInstanceDowntimeNotificationParameters) SetNodeName(v string) *QueryInstanceDowntimeNotificationParameters {
  s.NodeName = &v
  return s
}

type QueryInstanceDowntimeNotificationRequestHeader struct {
}

func (s QueryInstanceDowntimeNotificationRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryInstanceDowntimeNotificationRequestHeader) GoString() string {
  return s.String()
}

type QueryInstanceDowntimeNotificationResponseHeader struct {
}

func (s QueryInstanceDowntimeNotificationResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryInstanceDowntimeNotificationResponseHeader) GoString() string {
  return s.String()
}




type QueryPoPsMaintenanceNotificationRequest struct {
}

func (s QueryPoPsMaintenanceNotificationRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryPoPsMaintenanceNotificationRequest) GoString() string {
  return s.String()
}

type QueryPoPsMaintenanceNotificationResponse struct {
  // {"en":"Node name", "zh_CN":"节点英文名称"}
  NodeName *string `json:"nodeName,omitempty" xml:"nodeName,omitempty" require:"true"`
  // {"en":"Event:
  // CUTOVER
  // WITHDRAWAL
  // ", "zh_CN":"事件
  // CUTOVER     节点割接
  // WITHDRAWAL  节点退用"}
  Event *string `json:"event,omitempty" xml:"event,omitempty" require:"true"`
  // {"en":"Event status:
  // Scheduled
  // Executing
  // Executed
  // Canceled
  // ", "zh_CN":"事件状态
  // Scheduled    计划
  // Executing     执行中
  // Executed     执行完毕
  // Canceled     取消"}
  Status *string `json:"status,omitempty" xml:"status,omitempty" require:"true"`
  // {"en":"Cutover start time in yyyy-MM-dd HH:mm:ss.
  //     If it is a node retirement, then it is the retirement date.
  //     ", "zh_CN":"割接开始时间，格式 yyyy-MM-dd HH:mm:ss ，如果是节点退用，则是退用日期."}
  StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty" require:"true"`
  // {"en":"Cutover end time in yyyy-MM-dd HH:mm:ss.
  // If it is a node retirement, then the value is empty.", "zh_CN":"割接结束时间，格式yyyy-MM-dd HH:mm:ss
  // 如果是节点退用，则该值为空."}
  EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty" require:"true"`
  // {"en":"Instance information array", "zh_CN":"实例信息数组"}
  Servers []*QueryPoPsMaintenanceNotificationServer `json:"servers,omitempty" xml:"servers,omitempty" require:"true" type:"Repeated"`
}

func (s QueryPoPsMaintenanceNotificationResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryPoPsMaintenanceNotificationResponse) GoString() string {
  return s.String()
}

func (s *QueryPoPsMaintenanceNotificationResponse) SetNodeName(v string) *QueryPoPsMaintenanceNotificationResponse {
  s.NodeName = &v
  return s
}

func (s *QueryPoPsMaintenanceNotificationResponse) SetEvent(v string) *QueryPoPsMaintenanceNotificationResponse {
  s.Event = &v
  return s
}

func (s *QueryPoPsMaintenanceNotificationResponse) SetStatus(v string) *QueryPoPsMaintenanceNotificationResponse {
  s.Status = &v
  return s
}

func (s *QueryPoPsMaintenanceNotificationResponse) SetStartTime(v string) *QueryPoPsMaintenanceNotificationResponse {
  s.StartTime = &v
  return s
}

func (s *QueryPoPsMaintenanceNotificationResponse) SetEndTime(v string) *QueryPoPsMaintenanceNotificationResponse {
  s.EndTime = &v
  return s
}

func (s *QueryPoPsMaintenanceNotificationResponse) SetServers(v []*QueryPoPsMaintenanceNotificationServer) *QueryPoPsMaintenanceNotificationResponse {
  s.Servers = v
  return s
}

type QueryPoPsMaintenanceNotificationServer struct {
  // {"en":"Instance ID", "zh_CN":"实例id"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {"en":"Instance name", "zh_CN":"实例名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"Instance IPv4", "zh_CN":"实例IPv4"}
  Ipv4 []*string `json:"ipv4,omitempty" xml:"ipv4,omitempty" require:"true" type:"Repeated"`
  // {"en":"Instance IPv6", "zh_CN":"实例IPv6"}
  Ipv6 []*string `json:"ipv6,omitempty" xml:"ipv6,omitempty" require:"true" type:"Repeated"`
  // {"en":"Instance intranet IP", "zh_CN":"实例内网IP"}
  PrivateIp []*string `json:"privateIp,omitempty" xml:"privateIp,omitempty" require:"true" type:"Repeated"`
}

func (s QueryPoPsMaintenanceNotificationServer) String() string {
  return tea.Prettify(s)
}

func (s QueryPoPsMaintenanceNotificationServer) GoString() string {
  return s.String()
}

func (s *QueryPoPsMaintenanceNotificationServer) SetId(v string) *QueryPoPsMaintenanceNotificationServer {
  s.Id = &v
  return s
}

func (s *QueryPoPsMaintenanceNotificationServer) SetName(v string) *QueryPoPsMaintenanceNotificationServer {
  s.Name = &v
  return s
}

func (s *QueryPoPsMaintenanceNotificationServer) SetIpv4(v []*string) *QueryPoPsMaintenanceNotificationServer {
  s.Ipv4 = v
  return s
}

func (s *QueryPoPsMaintenanceNotificationServer) SetIpv6(v []*string) *QueryPoPsMaintenanceNotificationServer {
  s.Ipv6 = v
  return s
}

func (s *QueryPoPsMaintenanceNotificationServer) SetPrivateIp(v []*string) *QueryPoPsMaintenanceNotificationServer {
  s.PrivateIp = v
  return s
}

type QueryPoPsMaintenanceNotificationPaths struct {
}

func (s QueryPoPsMaintenanceNotificationPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryPoPsMaintenanceNotificationPaths) GoString() string {
  return s.String()
}

type QueryPoPsMaintenanceNotificationParameters struct {
  // {"en":"Node name, optional;
  // If it is multiple, it needs to be separated by a semicolon;
  // If not specified, query all nodes.
  // If the specified node customer is not in use or there is no cutover and dropout event, the node will not be returned in the response message.
  // ", "zh_CN":"节点英文名称，可选；
  // 如果是多个，需要用半角逗号隔开;
  // 如果未指定，则查询所有的节点。
  // 如果指定的节点客户未使用，或者没有割接退用事件，则在响应消息中不返回该节点。"}
  NodeName *string `json:"nodeName,omitempty" xml:"nodeName,omitempty"`
}

func (s QueryPoPsMaintenanceNotificationParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryPoPsMaintenanceNotificationParameters) GoString() string {
  return s.String()
}

func (s *QueryPoPsMaintenanceNotificationParameters) SetNodeName(v string) *QueryPoPsMaintenanceNotificationParameters {
  s.NodeName = &v
  return s
}

type QueryPoPsMaintenanceNotificationRequestHeader struct {
}

func (s QueryPoPsMaintenanceNotificationRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryPoPsMaintenanceNotificationRequestHeader) GoString() string {
  return s.String()
}

type QueryPoPsMaintenanceNotificationResponseHeader struct {
}

func (s QueryPoPsMaintenanceNotificationResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryPoPsMaintenanceNotificationResponseHeader) GoString() string {
  return s.String()
}




