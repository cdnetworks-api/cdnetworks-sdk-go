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

type QueryInstanceDowntimeNotificationRequestHeader struct {
}

func (s QueryInstanceDowntimeNotificationRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryInstanceDowntimeNotificationRequestHeader) GoString() string {
  return s.String()
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
  // {"en":"node name","zh_CN":"节点英文名称，多个逗号隔开"}
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

type QueryInstanceDowntimeNotificationResponse struct {
  // {"en":"reponse code","zh_CN":"请求返回码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"reponse message","zh_CN":"请求返回信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"","zh_CN":""}
  Data *QueryInstanceDowntimeNotificationResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
}

func (s QueryInstanceDowntimeNotificationResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryInstanceDowntimeNotificationResponse) GoString() string {
  return s.String()
}

func (s *QueryInstanceDowntimeNotificationResponse) SetCode(v string) *QueryInstanceDowntimeNotificationResponse {
  s.Code = &v
  return s
}

func (s *QueryInstanceDowntimeNotificationResponse) SetMessage(v string) *QueryInstanceDowntimeNotificationResponse {
  s.Message = &v
  return s
}

func (s *QueryInstanceDowntimeNotificationResponse) SetData(v *QueryInstanceDowntimeNotificationResponseData) *QueryInstanceDowntimeNotificationResponse {
  s.Data = v
  return s
}

type QueryInstanceDowntimeNotificationResponseData struct {
  // {"en":"","zh_CN":""}
  Events []*QueryInstanceDowntimeNotificationResponseDataEvents `json:"events,omitempty" xml:"events,omitempty" require:"true" type:"Repeated"`
}

func (s QueryInstanceDowntimeNotificationResponseData) String() string {
  return tea.Prettify(s)
}

func (s QueryInstanceDowntimeNotificationResponseData) GoString() string {
  return s.String()
}

func (s *QueryInstanceDowntimeNotificationResponseData) SetEvents(v []*QueryInstanceDowntimeNotificationResponseDataEvents) *QueryInstanceDowntimeNotificationResponseData {
  s.Events = v
  return s
}

type QueryInstanceDowntimeNotificationResponseDataEvents struct     {
  // {"en":"node name","zh_CN":"节点英文名称"}
  NodeName *string `json:"nodeName,omitempty" xml:"nodeName,omitempty" require:"true"`
  // {"en":"event name","zh_CN":"事件名称"}
  Event *string `json:"event,omitempty" xml:"event,omitempty" require:"true"`
  // {"en":"instnace fault time","zh_CN":"宕机时间"}
  Time *string `json:"time,omitempty" xml:"time,omitempty" require:"true"`
  // {"en":"","zh_CN":""}
  Servers []*QueryInstanceDowntimeNotificationResponseDataEventsServers `json:"servers,omitempty" xml:"servers,omitempty" require:"true" type:"Repeated"`
}

func (s QueryInstanceDowntimeNotificationResponseDataEvents) String() string {
  return tea.Prettify(s)
}

func (s QueryInstanceDowntimeNotificationResponseDataEvents) GoString() string {
  return s.String()
}

func (s *QueryInstanceDowntimeNotificationResponseDataEvents) SetNodeName(v string) *QueryInstanceDowntimeNotificationResponseDataEvents {
  s.NodeName = &v
  return s
}

func (s *QueryInstanceDowntimeNotificationResponseDataEvents) SetEvent(v string) *QueryInstanceDowntimeNotificationResponseDataEvents {
  s.Event = &v
  return s
}

func (s *QueryInstanceDowntimeNotificationResponseDataEvents) SetTime(v string) *QueryInstanceDowntimeNotificationResponseDataEvents {
  s.Time = &v
  return s
}

func (s *QueryInstanceDowntimeNotificationResponseDataEvents) SetServers(v []*QueryInstanceDowntimeNotificationResponseDataEventsServers) *QueryInstanceDowntimeNotificationResponseDataEvents {
  s.Servers = v
  return s
}

type QueryInstanceDowntimeNotificationResponseDataEventsServers struct     {
  // {"en":"instance id","zh_CN":"云主机id"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {"en":"instance name","zh_CN":"云主机名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"ipv4 array","zh_CN":"ipv4 集合"}
  Ipv4 []*string `json:"ipv4,omitempty" xml:"ipv4,omitempty" require:"true" type:"Repeated"`
  // {"en":"ipv6 array","zh_CN":"ipv6 集合"}
  Ipv6 []*string `json:"ipv6,omitempty" xml:"ipv6,omitempty" require:"true" type:"Repeated"`
}

func (s QueryInstanceDowntimeNotificationResponseDataEventsServers) String() string {
  return tea.Prettify(s)
}

func (s QueryInstanceDowntimeNotificationResponseDataEventsServers) GoString() string {
  return s.String()
}

func (s *QueryInstanceDowntimeNotificationResponseDataEventsServers) SetId(v string) *QueryInstanceDowntimeNotificationResponseDataEventsServers {
  s.Id = &v
  return s
}

func (s *QueryInstanceDowntimeNotificationResponseDataEventsServers) SetName(v string) *QueryInstanceDowntimeNotificationResponseDataEventsServers {
  s.Name = &v
  return s
}

func (s *QueryInstanceDowntimeNotificationResponseDataEventsServers) SetIpv4(v []*string) *QueryInstanceDowntimeNotificationResponseDataEventsServers {
  s.Ipv4 = v
  return s
}

func (s *QueryInstanceDowntimeNotificationResponseDataEventsServers) SetIpv6(v []*string) *QueryInstanceDowntimeNotificationResponseDataEventsServers {
  s.Ipv6 = v
  return s
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

type QueryPoPsMaintenanceNotificationRequestHeader struct {
}

func (s QueryPoPsMaintenanceNotificationRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryPoPsMaintenanceNotificationRequestHeader) GoString() string {
  return s.String()
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
  // {"en":"Node name, optional;\nIf it is multiple, it needs to be separated by a semicolon;\nIf not specified, query all nodes.\nIf the specified node customer is not in use or there is no cutover and dropout event, the node will not be returned in the response message.","zh_CN":"节点英文名称，可选；\n如果是多个，需要用半角逗号隔开;\n如果未指定，则查询所有的节点。\n如果指定的节点客户未使用，或者没有割接退用事件，则在响应消息中不返回该节点。"}
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

type QueryPoPsMaintenanceNotificationResponse struct {
  // {"en":"reponse code","zh_CN":"请求返回码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"reponse message","zh_CN":"请求返回信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"reponse data","zh_CN":"请求返回数据"}
  Data *QueryPoPsMaintenanceNotificationResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
}

func (s QueryPoPsMaintenanceNotificationResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryPoPsMaintenanceNotificationResponse) GoString() string {
  return s.String()
}

func (s *QueryPoPsMaintenanceNotificationResponse) SetCode(v string) *QueryPoPsMaintenanceNotificationResponse {
  s.Code = &v
  return s
}

func (s *QueryPoPsMaintenanceNotificationResponse) SetMessage(v string) *QueryPoPsMaintenanceNotificationResponse {
  s.Message = &v
  return s
}

func (s *QueryPoPsMaintenanceNotificationResponse) SetData(v *QueryPoPsMaintenanceNotificationResponseData) *QueryPoPsMaintenanceNotificationResponse {
  s.Data = v
  return s
}

type QueryPoPsMaintenanceNotificationResponseData struct {
  // {"en":"","zh_CN":""}
  Events []*QueryPoPsMaintenanceNotificationResponseDataEvents `json:"events,omitempty" xml:"events,omitempty" require:"true" type:"Repeated"`
}

func (s QueryPoPsMaintenanceNotificationResponseData) String() string {
  return tea.Prettify(s)
}

func (s QueryPoPsMaintenanceNotificationResponseData) GoString() string {
  return s.String()
}

func (s *QueryPoPsMaintenanceNotificationResponseData) SetEvents(v []*QueryPoPsMaintenanceNotificationResponseDataEvents) *QueryPoPsMaintenanceNotificationResponseData {
  s.Events = v
  return s
}

type QueryPoPsMaintenanceNotificationResponseDataEvents struct     {
  // {"en":"Node name","zh_CN":"节点英文名称"}
  NodeName *string `json:"nodeName,omitempty" xml:"nodeName,omitempty" require:"true"`
  // {"en":"Event:CUTOVERWITHDRAWAL","zh_CN":"事件CUTOVER     节点割接WITHDRAWAL  节点退用"}
  Event *string `json:"event,omitempty" xml:"event,omitempty" require:"true"`
  // {"en":"Event status:ScheduledExecutingExecutedCanceled","zh_CN":"事件状态Scheduled    计划Executing     执行中Executed     执行完毕Canceled     取消"}
  Status *string `json:"status,omitempty" xml:"status,omitempty" require:"true"`
  // {"en":"Cutover start time in yyyy-MM-dd HH:mm:ss.If it is a node retirement, then it is the retirement date.","zh_CN":"割接开始时间，格式 yyyy-MM-dd HH:mm:ss ，如果是节点退用，则是退用日期."}
  StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty" require:"true"`
  // {"en":"Cutover end time in yyyy-MM-dd HH:mm:ss.If it is a node retirement, then the value is empty.","zh_CN":"割接结束时间，格式yyyy-MM-dd HH:mm:ss如果是节点退用，则该值为空."}
  EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty" require:"true"`
  // {"en":"Instance information array","zh_CN":"实例信息数组"}
  Servers []*QueryPoPsMaintenanceNotificationResponseDataEventsServers `json:"servers,omitempty" xml:"servers,omitempty" require:"true" type:"Repeated"`
}

func (s QueryPoPsMaintenanceNotificationResponseDataEvents) String() string {
  return tea.Prettify(s)
}

func (s QueryPoPsMaintenanceNotificationResponseDataEvents) GoString() string {
  return s.String()
}

func (s *QueryPoPsMaintenanceNotificationResponseDataEvents) SetNodeName(v string) *QueryPoPsMaintenanceNotificationResponseDataEvents {
  s.NodeName = &v
  return s
}

func (s *QueryPoPsMaintenanceNotificationResponseDataEvents) SetEvent(v string) *QueryPoPsMaintenanceNotificationResponseDataEvents {
  s.Event = &v
  return s
}

func (s *QueryPoPsMaintenanceNotificationResponseDataEvents) SetStatus(v string) *QueryPoPsMaintenanceNotificationResponseDataEvents {
  s.Status = &v
  return s
}

func (s *QueryPoPsMaintenanceNotificationResponseDataEvents) SetStartTime(v string) *QueryPoPsMaintenanceNotificationResponseDataEvents {
  s.StartTime = &v
  return s
}

func (s *QueryPoPsMaintenanceNotificationResponseDataEvents) SetEndTime(v string) *QueryPoPsMaintenanceNotificationResponseDataEvents {
  s.EndTime = &v
  return s
}

func (s *QueryPoPsMaintenanceNotificationResponseDataEvents) SetServers(v []*QueryPoPsMaintenanceNotificationResponseDataEventsServers) *QueryPoPsMaintenanceNotificationResponseDataEvents {
  s.Servers = v
  return s
}

type QueryPoPsMaintenanceNotificationResponseDataEventsServers struct     {
  // {"en":"Instance ID","zh_CN":"实例id"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {"en":"Instance name","zh_CN":"实例名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"Instance IPv4","zh_CN":"实例IPv4"}
  Ipv4 []*string `json:"ipv4,omitempty" xml:"ipv4,omitempty" require:"true" type:"Repeated"`
  // {"en":"Instance IPv6","zh_CN":"实例IPv6"}
  Ipv6 []*string `json:"ipv6,omitempty" xml:"ipv6,omitempty" require:"true" type:"Repeated"`
  // {"en":"Instance intranet IP","zh_CN":"实例内网IP"}
  PrivateIp []*string `json:"privateIp,omitempty" xml:"privateIp,omitempty" require:"true" type:"Repeated"`
}

func (s QueryPoPsMaintenanceNotificationResponseDataEventsServers) String() string {
  return tea.Prettify(s)
}

func (s QueryPoPsMaintenanceNotificationResponseDataEventsServers) GoString() string {
  return s.String()
}

func (s *QueryPoPsMaintenanceNotificationResponseDataEventsServers) SetId(v string) *QueryPoPsMaintenanceNotificationResponseDataEventsServers {
  s.Id = &v
  return s
}

func (s *QueryPoPsMaintenanceNotificationResponseDataEventsServers) SetName(v string) *QueryPoPsMaintenanceNotificationResponseDataEventsServers {
  s.Name = &v
  return s
}

func (s *QueryPoPsMaintenanceNotificationResponseDataEventsServers) SetIpv4(v []*string) *QueryPoPsMaintenanceNotificationResponseDataEventsServers {
  s.Ipv4 = v
  return s
}

func (s *QueryPoPsMaintenanceNotificationResponseDataEventsServers) SetIpv6(v []*string) *QueryPoPsMaintenanceNotificationResponseDataEventsServers {
  s.Ipv6 = v
  return s
}

func (s *QueryPoPsMaintenanceNotificationResponseDataEventsServers) SetPrivateIp(v []*string) *QueryPoPsMaintenanceNotificationResponseDataEventsServers {
  s.PrivateIp = v
  return s
}

type QueryPoPsMaintenanceNotificationResponseHeader struct {
}

func (s QueryPoPsMaintenanceNotificationResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryPoPsMaintenanceNotificationResponseHeader) GoString() string {
  return s.String()
}




