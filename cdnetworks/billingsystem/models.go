package billingsystem

// This file is auto-generated, don't edit it. Thanks.
import (
  "github.com/alibabacloud-go/tea/tea"
)

type QueryBillingDetailsOfComputingServiceRequest struct {
}

func (s QueryBillingDetailsOfComputingServiceRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryBillingDetailsOfComputingServiceRequest) GoString() string {
  return s.String()
}

type QueryBillingDetailsOfComputingServiceRequestHeader struct {
}

func (s QueryBillingDetailsOfComputingServiceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryBillingDetailsOfComputingServiceRequestHeader) GoString() string {
  return s.String()
}

type QueryBillingDetailsOfComputingServicePaths struct {
}

func (s QueryBillingDetailsOfComputingServicePaths) String() string {
  return tea.Prettify(s)
}

func (s QueryBillingDetailsOfComputingServicePaths) GoString() string {
  return s.String()
}

type QueryBillingDetailsOfComputingServiceParameters struct {
  // {"en":"Starting time of use within the accounting period,format:yyyy-MM-dd","zh_CN":"账期开始时间，格式yyyy-MM-dd"}
  StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty" require:"true"`
  // {"en":"The end time of the accounting period, in the format yyyy-MM-dd. If not filled in, it will be the current time.","zh_CN":"账期结束时间，格式yyyy-MM-dd，如果未填写则为当前时间"}
  EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
}

func (s QueryBillingDetailsOfComputingServiceParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryBillingDetailsOfComputingServiceParameters) GoString() string {
  return s.String()
}

func (s *QueryBillingDetailsOfComputingServiceParameters) SetStartTime(v string) *QueryBillingDetailsOfComputingServiceParameters {
  s.StartTime = &v
  return s
}

func (s *QueryBillingDetailsOfComputingServiceParameters) SetEndTime(v string) *QueryBillingDetailsOfComputingServiceParameters {
  s.EndTime = &v
  return s
}

type QueryBillingDetailsOfComputingServiceResponse struct {
  // {"en":"Instance computing power information array","zh_CN":"实例算力信息数组"}
  Servers []*QueryBillingDetailsOfComputingServiceResponseServers `json:"servers,omitempty" xml:"servers,omitempty" require:"true" type:"Repeated"`
}

func (s QueryBillingDetailsOfComputingServiceResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryBillingDetailsOfComputingServiceResponse) GoString() string {
  return s.String()
}

func (s *QueryBillingDetailsOfComputingServiceResponse) SetServers(v []*QueryBillingDetailsOfComputingServiceResponseServers) *QueryBillingDetailsOfComputingServiceResponse {
  s.Servers = v
  return s
}

type QueryBillingDetailsOfComputingServiceResponseServers struct     {
  // {"en":"Instance id","zh_CN":"实例id"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {"en":"Instance name","zh_CN":"实例名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"Instance IPv4","zh_CN":"实例IPv4"}
  Ipv4 []*string `json:"ipv4,omitempty" xml:"ipv4,omitempty" require:"true" type:"Repeated"`
  // {"en":"Instance IPv6","zh_CN":"实例IPv6"}
  Ipv6 []*string `json:"ipv6,omitempty" xml:"ipv6,omitempty" require:"true" type:"Repeated"`
  // {"en":"Instance type:1---Bare metal instance;-1:Virtual machine instance","zh_CN":"实例类型\n1：裸机实例,-1：虚拟机实例"}
  IsBm *int `json:"isBm,omitempty" xml:"isBm,omitempty" require:"true"`
  // {"en":"Instance Status:\nRUNNING\nSTOPPED\nERROR\nDELETED\nRESTARTING\nSTARTING\nSTOPPING\nSNAPSHOTTING\nREBUILDING\nMIGRATING","zh_CN":"实例状态\nRUNNING运行状态\nSTOPPED停机\nERROR错误\nDELETED已销毁\nRESTARTING  重启中\nSTARTING  启动中\nSTOPPING  停止中\nSNAPSHOTTING  制作快照镜像中\nREBUILDING 重建中\nMIGRATING 迁移中"}
  State *string `json:"state,omitempty" xml:"state,omitempty" require:"true"`
  // {"en":"Node name","zh_CN":"节点名称"}
  NodeName *string `json:"nodeName,omitempty" xml:"nodeName,omitempty" require:"true"`
  // {"en":"Node Chinese Name","zh_CN":"节点中文名称"}
  NodeNameCn *string `json:"nodeNameCn,omitempty" xml:"nodeNameCn,omitempty" require:"true"`
  // {"en":"Charge region","zh_CN":"计费区域，14个取值：\n中国大陆\n亚太\n香港\n台湾\n美洲\n欧洲\n中东\n非洲\n台湾\n香港\n非洲\n南美\n澳大利亚\n印度"}
  ChargeRegion *string `json:"chargeRegion,omitempty" xml:"chargeRegion,omitempty" require:"true"`
  // {"en":"Instance spec","zh_CN":"实例规格"}
  InstanceType *string `json:"instanceType,omitempty" xml:"instanceType,omitempty" require:"true"`
  // {"en":"Instance configuration, including CPU, memory, and disk specifications information of the virtual machine","zh_CN":"实例配置，包含虚拟机的CPU、内存、磁盘规格信息"}
  ServerFeature *string `json:"serverFeature,omitempty" xml:"serverFeature,omitempty" require:"true"`
  // {"en":"Starting time of use within the accounting period","zh_CN":"账期内开始使用时间，格式yyyy-MM-dd"}
  PeriodStartTime *string `json:"periodStartTime,omitempty" xml:"periodStartTime,omitempty" require:"true"`
  // {"en":"End of usage time within the accounting period","zh_CN":"账期内结束时间，格式yyyy-MM-dd"}
  PeriodEndTime *string `json:"periodEndTime,omitempty" xml:"periodEndTime,omitempty" require:"true"`
  // {"en":"Usage days within the accounting period","zh_CN":"账期内使用天数"}
  PeriodDays *int `json:"periodDays,omitempty" xml:"periodDays,omitempty" require:"true"`
  // {"en":"Instance usage during the accounting period","zh_CN":"账期内实例用量"}
  PeriodCount *int64 `json:"periodCount,omitempty" xml:"periodCount,omitempty" require:"true"`
}

func (s QueryBillingDetailsOfComputingServiceResponseServers) String() string {
  return tea.Prettify(s)
}

func (s QueryBillingDetailsOfComputingServiceResponseServers) GoString() string {
  return s.String()
}

func (s *QueryBillingDetailsOfComputingServiceResponseServers) SetId(v string) *QueryBillingDetailsOfComputingServiceResponseServers {
  s.Id = &v
  return s
}

func (s *QueryBillingDetailsOfComputingServiceResponseServers) SetName(v string) *QueryBillingDetailsOfComputingServiceResponseServers {
  s.Name = &v
  return s
}

func (s *QueryBillingDetailsOfComputingServiceResponseServers) SetIpv4(v []*string) *QueryBillingDetailsOfComputingServiceResponseServers {
  s.Ipv4 = v
  return s
}

func (s *QueryBillingDetailsOfComputingServiceResponseServers) SetIpv6(v []*string) *QueryBillingDetailsOfComputingServiceResponseServers {
  s.Ipv6 = v
  return s
}

func (s *QueryBillingDetailsOfComputingServiceResponseServers) SetIsBm(v int) *QueryBillingDetailsOfComputingServiceResponseServers {
  s.IsBm = &v
  return s
}

func (s *QueryBillingDetailsOfComputingServiceResponseServers) SetState(v string) *QueryBillingDetailsOfComputingServiceResponseServers {
  s.State = &v
  return s
}

func (s *QueryBillingDetailsOfComputingServiceResponseServers) SetNodeName(v string) *QueryBillingDetailsOfComputingServiceResponseServers {
  s.NodeName = &v
  return s
}

func (s *QueryBillingDetailsOfComputingServiceResponseServers) SetNodeNameCn(v string) *QueryBillingDetailsOfComputingServiceResponseServers {
  s.NodeNameCn = &v
  return s
}

func (s *QueryBillingDetailsOfComputingServiceResponseServers) SetChargeRegion(v string) *QueryBillingDetailsOfComputingServiceResponseServers {
  s.ChargeRegion = &v
  return s
}

func (s *QueryBillingDetailsOfComputingServiceResponseServers) SetInstanceType(v string) *QueryBillingDetailsOfComputingServiceResponseServers {
  s.InstanceType = &v
  return s
}

func (s *QueryBillingDetailsOfComputingServiceResponseServers) SetServerFeature(v string) *QueryBillingDetailsOfComputingServiceResponseServers {
  s.ServerFeature = &v
  return s
}

func (s *QueryBillingDetailsOfComputingServiceResponseServers) SetPeriodStartTime(v string) *QueryBillingDetailsOfComputingServiceResponseServers {
  s.PeriodStartTime = &v
  return s
}

func (s *QueryBillingDetailsOfComputingServiceResponseServers) SetPeriodEndTime(v string) *QueryBillingDetailsOfComputingServiceResponseServers {
  s.PeriodEndTime = &v
  return s
}

func (s *QueryBillingDetailsOfComputingServiceResponseServers) SetPeriodDays(v int) *QueryBillingDetailsOfComputingServiceResponseServers {
  s.PeriodDays = &v
  return s
}

func (s *QueryBillingDetailsOfComputingServiceResponseServers) SetPeriodCount(v int64) *QueryBillingDetailsOfComputingServiceResponseServers {
  s.PeriodCount = &v
  return s
}

type QueryBillingDetailsOfComputingServiceResponseHeader struct {
}

func (s QueryBillingDetailsOfComputingServiceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryBillingDetailsOfComputingServiceResponseHeader) GoString() string {
  return s.String()
}




