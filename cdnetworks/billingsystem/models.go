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

type QueryBillingDetailsOfComputingServiceServerBill struct {
  // {"en":"Instance id", "zh_CN":"实例id"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {"en":"Instance name", "zh_CN":"实例名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"Instance IPv4", "zh_CN":"实例IPv4"}
  Ipv4 []*string `json:"ipv4,omitempty" xml:"ipv4,omitempty" require:"true" type:"Repeated"`
  // {"en":"Instance IPv6", "zh_CN":"实例IPv6"}
  Ipv6 []*string `json:"ipv6,omitempty" xml:"ipv6,omitempty" require:"true" type:"Repeated"`
  // {"en":"Instance type:1---Bare metal instance;-1:Virtual machine instance", "zh_CN":"实例类型
  // 1：裸机实例,-1：虚拟机实例"}
  IsBm *int `json:"isBm,omitempty" xml:"isBm,omitempty" require:"true"`
  // {"en":"Instance Status:
  // RUNNING
  // STOPPED
  // ERROR
  // DELETED
  // RESTARTING
  // STARTING
  // STOPPING
  // SNAPSHOTTING
  // REBUILDING
  // MIGRATING
  // ", "zh_CN":"实例状态
  // RUNNING	运行状态
  // STOPPED	停机
  // ERROR	错误
  // DELETED	已销毁
  // RESTARTING  重启中
  // STARTING  启动中
  // STOPPING  停止中
  // SNAPSHOTTING  制作快照镜像中
  // REBUILDING 重建中
  // MIGRATING 迁移中"}
  State *string `json:"state,omitempty" xml:"state,omitempty" require:"true"`
  // {"en":"Node name", "zh_CN":"节点名称"}
  NodeName *string `json:"nodeName,omitempty" xml:"nodeName,omitempty" require:"true"`
  // {"en":"Node Chinese Name", "zh_CN":"节点中文名称"}
  NodeNameCn *string `json:"nodeNameCn,omitempty" xml:"nodeNameCn,omitempty" require:"true"`
  // {"en":"Charge region", "zh_CN":"计费区域，14个取值：
  // 中国大陆
  // 亚太
  // 香港
  // 台湾
  // 美洲
  // 欧洲
  // 中东
  // 非洲
  // 台湾 
  // 香港
  // 非洲
  // 南美
  // 澳大利亚
  // 印度"}
  ChargeRegion *string `json:"chargeRegion,omitempty" xml:"chargeRegion,omitempty" require:"true"`
  // {"en":"Instance spec", "zh_CN":"实例规格"}
  InstanceType *string `json:"instanceType,omitempty" xml:"instanceType,omitempty" require:"true"`
  // {"en":"Instance configuration, including CPU, memory, and disk specifications information of the virtual machine", "zh_CN":"实例配置，包含虚拟机的CPU、内存、磁盘规格信息"}
  ServerFeature *string `json:"serverFeature,omitempty" xml:"serverFeature,omitempty" require:"true"`
  // {"en":"Starting time of use within the accounting period", "zh_CN":"账期内开始使用时间，格式yyyy-MM-dd"}
  PeriodStartTime *string `json:"periodStartTime,omitempty" xml:"periodStartTime,omitempty" require:"true"`
  // {"en":"End of usage time within the accounting period", "zh_CN":"账期内结束时间，格式yyyy-MM-dd"}
  PeriodEndTime *string `json:"periodEndTime,omitempty" xml:"periodEndTime,omitempty" require:"true"`
  // {"en":"Usage days within the accounting period", "zh_CN":"账期内使用天数"}
  PeriodDays *int32 `json:"periodDays,omitempty" xml:"periodDays,omitempty" require:"true"`
  // {"en":"Instance usage during the accounting period", "zh_CN":"账期内实例用量"}
  PeriodCount *float32 `json:"periodCount,omitempty" xml:"periodCount,omitempty" require:"true"`
}

func (s QueryBillingDetailsOfComputingServiceServerBill) String() string {
  return tea.Prettify(s)
}

func (s QueryBillingDetailsOfComputingServiceServerBill) GoString() string {
  return s.String()
}

func (s *QueryBillingDetailsOfComputingServiceServerBill) SetId(v string) *QueryBillingDetailsOfComputingServiceServerBill {
  s.Id = &v
  return s
}

func (s *QueryBillingDetailsOfComputingServiceServerBill) SetName(v string) *QueryBillingDetailsOfComputingServiceServerBill {
  s.Name = &v
  return s
}

func (s *QueryBillingDetailsOfComputingServiceServerBill) SetIpv4(v []*string) *QueryBillingDetailsOfComputingServiceServerBill {
  s.Ipv4 = v
  return s
}

func (s *QueryBillingDetailsOfComputingServiceServerBill) SetIpv6(v []*string) *QueryBillingDetailsOfComputingServiceServerBill {
  s.Ipv6 = v
  return s
}

func (s *QueryBillingDetailsOfComputingServiceServerBill) SetIsBm(v int) *QueryBillingDetailsOfComputingServiceServerBill {
  s.IsBm = &v
  return s
}

func (s *QueryBillingDetailsOfComputingServiceServerBill) SetState(v string) *QueryBillingDetailsOfComputingServiceServerBill {
  s.State = &v
  return s
}

func (s *QueryBillingDetailsOfComputingServiceServerBill) SetNodeName(v string) *QueryBillingDetailsOfComputingServiceServerBill {
  s.NodeName = &v
  return s
}

func (s *QueryBillingDetailsOfComputingServiceServerBill) SetNodeNameCn(v string) *QueryBillingDetailsOfComputingServiceServerBill {
  s.NodeNameCn = &v
  return s
}

func (s *QueryBillingDetailsOfComputingServiceServerBill) SetChargeRegion(v string) *QueryBillingDetailsOfComputingServiceServerBill {
  s.ChargeRegion = &v
  return s
}

func (s *QueryBillingDetailsOfComputingServiceServerBill) SetInstanceType(v string) *QueryBillingDetailsOfComputingServiceServerBill {
  s.InstanceType = &v
  return s
}

func (s *QueryBillingDetailsOfComputingServiceServerBill) SetServerFeature(v string) *QueryBillingDetailsOfComputingServiceServerBill {
  s.ServerFeature = &v
  return s
}

func (s *QueryBillingDetailsOfComputingServiceServerBill) SetPeriodStartTime(v string) *QueryBillingDetailsOfComputingServiceServerBill {
  s.PeriodStartTime = &v
  return s
}

func (s *QueryBillingDetailsOfComputingServiceServerBill) SetPeriodEndTime(v string) *QueryBillingDetailsOfComputingServiceServerBill {
  s.PeriodEndTime = &v
  return s
}

func (s *QueryBillingDetailsOfComputingServiceServerBill) SetPeriodDays(v int32) *QueryBillingDetailsOfComputingServiceServerBill {
  s.PeriodDays = &v
  return s
}

func (s *QueryBillingDetailsOfComputingServiceServerBill) SetPeriodCount(v float32) *QueryBillingDetailsOfComputingServiceServerBill {
  s.PeriodCount = &v
  return s
}

type QueryBillingDetailsOfComputingServiceResponse struct {
  // {"en":"Instance computing power information array", "zh_CN":"实例算力信息数组"}
  Servers []*QueryBillingDetailsOfComputingServiceServerBill `json:"servers,omitempty" xml:"servers,omitempty" require:"true" type:"Repeated"`
}

func (s QueryBillingDetailsOfComputingServiceResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryBillingDetailsOfComputingServiceResponse) GoString() string {
  return s.String()
}

func (s *QueryBillingDetailsOfComputingServiceResponse) SetServers(v []*QueryBillingDetailsOfComputingServiceServerBill) *QueryBillingDetailsOfComputingServiceResponse {
  s.Servers = v
  return s
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
  // {"en":"Starting time of use within the accounting period,format:yyyy-MM-dd", "zh_CN":"账期开始时间，格式yyyy-MM-dd"}
  StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty" require:"true"`
  // {"en":"The end time of the accounting period, in the format yyyy-MM-dd. If not filled in, it will be the current time.", "zh_CN":"账期结束时间，格式yyyy-MM-dd，如果未填写则为当前时间"}
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

type QueryBillingDetailsOfComputingServiceRequestHeader struct {
}

func (s QueryBillingDetailsOfComputingServiceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryBillingDetailsOfComputingServiceRequestHeader) GoString() string {
  return s.String()
}

type QueryBillingDetailsOfComputingServiceResponseHeader struct {
}

func (s QueryBillingDetailsOfComputingServiceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryBillingDetailsOfComputingServiceResponseHeader) GoString() string {
  return s.String()
}




