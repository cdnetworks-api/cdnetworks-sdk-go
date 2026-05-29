package basicmonitor

// This file is auto-generated, don't edit it. Thanks.
import (
  "github.com/alibabacloud-go/tea/tea"
)

type VMPQueryServerMetricRequest struct {
}

func (s VMPQueryServerMetricRequest) String() string {
  return tea.Prettify(s)
}

func (s VMPQueryServerMetricRequest) GoString() string {
  return s.String()
}

type VMPQueryServerMetricResponse struct {
  // {"en":"Instance information array", "zh_CN":"实例信息数组"}
  Servers []*VMPQueryServerMetricServer `json:"servers,omitempty" xml:"servers,omitempty" require:"true" type:"Repeated"`
}

func (s VMPQueryServerMetricResponse) String() string {
  return tea.Prettify(s)
}

func (s VMPQueryServerMetricResponse) GoString() string {
  return s.String()
}

func (s *VMPQueryServerMetricResponse) SetServers(v []*VMPQueryServerMetricServer) *VMPQueryServerMetricResponse {
  s.Servers = v
  return s
}

type VMPQueryServerMetricServer struct {
  // {"en":"instance id", "zh_CN":"实例唯一标识"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {"en":"instance ip", "zh_CN":"实例公网IP"}
  InstanceIp *string `json:"instanceIp,omitempty" xml:"instanceIp,omitempty" require:"true"`
  // {"en":"VMPQueryServerMetricBandwidth information", "zh_CN":"带宽信息"}
  Bandwidths []*VMPQueryServerMetricBandwidth `json:"bandwidths,omitempty" xml:"bandwidths,omitempty" require:"true" type:"Repeated"`
  // {"en":"Memory information", "zh_CN":"内存信息"}
  Mems []*VMPQueryServerMetricMem `json:"mems,omitempty" xml:"mems,omitempty" require:"true" type:"Repeated"`
  // {"en":"CPU information", "zh_CN":"CPU信息"}
  Cpus []*VMPQueryServerMetricCpu `json:"cpus,omitempty" xml:"cpus,omitempty" require:"true" type:"Repeated"`
}

func (s VMPQueryServerMetricServer) String() string {
  return tea.Prettify(s)
}

func (s VMPQueryServerMetricServer) GoString() string {
  return s.String()
}

func (s *VMPQueryServerMetricServer) SetId(v string) *VMPQueryServerMetricServer {
  s.Id = &v
  return s
}

func (s *VMPQueryServerMetricServer) SetInstanceIp(v string) *VMPQueryServerMetricServer {
  s.InstanceIp = &v
  return s
}

func (s *VMPQueryServerMetricServer) SetBandwidths(v []*VMPQueryServerMetricBandwidth) *VMPQueryServerMetricServer {
  s.Bandwidths = v
  return s
}

func (s *VMPQueryServerMetricServer) SetMems(v []*VMPQueryServerMetricMem) *VMPQueryServerMetricServer {
  s.Mems = v
  return s
}

func (s *VMPQueryServerMetricServer) SetCpus(v []*VMPQueryServerMetricCpu) *VMPQueryServerMetricServer {
  s.Cpus = v
  return s
}

type VMPQueryServerMetricBandwidth struct {
  // {"en":"The bandwidth collection time, formatted as YYYYMMDDHHMM, is 5 minutes granularity", "zh_CN":"带宽采集时间，格式yyyyMMddHHmm，为5分钟粒度"}
  StatTime *string `json:"statTime,omitempty" xml:"statTime,omitempty" require:"true"`
  // {"en":"Total outflow bandwidth (including intra-node flow) in Mbps", "zh_CN":"总流出带宽（含节点内流量），单位Mbps"}
  Out *string `json:"out,omitempty" xml:"out,omitempty" require:"true"`
  // {"en":"Total inflow bandwidth (including intra-node traffic) in Mbps", "zh_CN":"总流入带宽（含节点内流量），单位Mbps"}
  In *string `json:"in,omitempty" xml:"in,omitempty" require:"true"`
  // {"en":"Inflow bandwidth of external network (not including intra-node traffic), unit of Mbps", "zh_CN":"外网流入带宽（不含节点内流量），单位Mbps"}
  ExtIn *string `json:"extIn,omitempty" xml:"extIn,omitempty" require:"true"`
  // {"en":"Outflow bandwidth of the external network (excluding intra-node traffic), unit of Mbps", "zh_CN":"外网流出带宽（不含节点内流量），单位Mbps"}
  ExtOut *string `json:"extOut,omitempty" xml:"extOut,omitempty" require:"true"`
}

func (s VMPQueryServerMetricBandwidth) String() string {
  return tea.Prettify(s)
}

func (s VMPQueryServerMetricBandwidth) GoString() string {
  return s.String()
}

func (s *VMPQueryServerMetricBandwidth) SetStatTime(v string) *VMPQueryServerMetricBandwidth {
  s.StatTime = &v
  return s
}

func (s *VMPQueryServerMetricBandwidth) SetOut(v string) *VMPQueryServerMetricBandwidth {
  s.Out = &v
  return s
}

func (s *VMPQueryServerMetricBandwidth) SetIn(v string) *VMPQueryServerMetricBandwidth {
  s.In = &v
  return s
}

func (s *VMPQueryServerMetricBandwidth) SetExtIn(v string) *VMPQueryServerMetricBandwidth {
  s.ExtIn = &v
  return s
}

func (s *VMPQueryServerMetricBandwidth) SetExtOut(v string) *VMPQueryServerMetricBandwidth {
  s.ExtOut = &v
  return s
}

type VMPQueryServerMetricMem struct {
  // {"en":"The memory data collection time, in the format YYYYMMDDHHMM, is 5 minutes granularity", "zh_CN":"内存数据采集时间，格式yyyyMMddHHmm，为5分钟粒度"}
  StatTime *string `json:"statTime,omitempty" xml:"statTime,omitempty" require:"true"`
  // {"en":"Memory usage in MB", "zh_CN":"内存使用量，单位为MB"}
  Used *string `json:"used,omitempty" xml:"used,omitempty" require:"true"`
  // {"en":"Total memory, in MB", "zh_CN":"总内存量，单位为MB"}
  Total *string `json:"total,omitempty" xml:"total,omitempty" require:"true"`
}

func (s VMPQueryServerMetricMem) String() string {
  return tea.Prettify(s)
}

func (s VMPQueryServerMetricMem) GoString() string {
  return s.String()
}

func (s *VMPQueryServerMetricMem) SetStatTime(v string) *VMPQueryServerMetricMem {
  s.StatTime = &v
  return s
}

func (s *VMPQueryServerMetricMem) SetUsed(v string) *VMPQueryServerMetricMem {
  s.Used = &v
  return s
}

func (s *VMPQueryServerMetricMem) SetTotal(v string) *VMPQueryServerMetricMem {
  s.Total = &v
  return s
}

type VMPQueryServerMetricCpu struct {
  // {"en":"CPU usage collection time in the format YYYYMMDDHHMM, with a granularity of 5 minutes", "zh_CN":"cpu使用率采集时间，格式yyyyMMddHHmm，为5分钟粒度"}
  StatTime *string `json:"statTime,omitempty" xml:"statTime,omitempty" require:"true"`
  // {"en":"CPU utilization, with a value of 0-1, such as 0.25, represents 25% utilization", "zh_CN":"cpu使用率，取值0-1，如0.25，该值则表示使用率为25%"}
  Usage *string `json:"usage,omitempty" xml:"usage,omitempty" require:"true"`
}

func (s VMPQueryServerMetricCpu) String() string {
  return tea.Prettify(s)
}

func (s VMPQueryServerMetricCpu) GoString() string {
  return s.String()
}

func (s *VMPQueryServerMetricCpu) SetStatTime(v string) *VMPQueryServerMetricCpu {
  s.StatTime = &v
  return s
}

func (s *VMPQueryServerMetricCpu) SetUsage(v string) *VMPQueryServerMetricCpu {
  s.Usage = &v
  return s
}

type VMPQueryServerMetricPaths struct {
}

func (s VMPQueryServerMetricPaths) String() string {
  return tea.Prettify(s)
}

func (s VMPQueryServerMetricPaths) GoString() string {
  return s.String()
}

type VMPQueryServerMetricParameters struct {
  // {"en":"The ID of the virtual machine to query. At most 30 queries can be made at a time, ids 
  //  are separated by character  ','", "zh_CN":"云主机ID。单次最多查询 20 条 ID，ID 之间用半角逗号字符','隔开。"}
  Ids *string `json:"ids,omitempty" xml:"ids,omitempty" require:"true"`
  // {"en":"Monitoring data type of query, value:
  // CPU: CPU utilization
  // VMPQueryServerMetricMem: memory
  // VMPQueryServerMetricBandwidth: Traffic data", "zh_CN":"查询的监控数据类型，取值：
  // cpu： cpu使用率
  // mem：内存
  // bandwidth：流量数据"}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
  // {"en":"Query time range in format: yyyymmddhhmm-yyyymmddhhmm
  // Query data for nearly 90 days, a single query no more than 3 days.
  // For example: 202001201730-202001201930 means to query 2020-01-20 17:30 to 19:30 monitoring data.", "zh_CN":"查询时间范围，格式：yyyyMMddHHmm- yyyyMMddHHmm
  // 查询近90天的数据，单次查询不超过10天。
  // 例如：202001201730-202001201930表示查询2020-01-20 17:30到19:30的监控数据。"}
  StatTime *string `json:"statTime,omitempty" xml:"statTime,omitempty" require:"true"`
  // {"en":"The ISP code  ", "zh_CN":"当查询流量数据时可填。实例所属运营商：dx-电信；wt-网通；yd-移动。一次仅允许传入一个运营商参数"}
  Carrier *string `json:"carrier,omitempty" xml:"carrier,omitempty"`
}

func (s VMPQueryServerMetricParameters) String() string {
  return tea.Prettify(s)
}

func (s VMPQueryServerMetricParameters) GoString() string {
  return s.String()
}

func (s *VMPQueryServerMetricParameters) SetIds(v string) *VMPQueryServerMetricParameters {
  s.Ids = &v
  return s
}

func (s *VMPQueryServerMetricParameters) SetType(v string) *VMPQueryServerMetricParameters {
  s.Type = &v
  return s
}

func (s *VMPQueryServerMetricParameters) SetStatTime(v string) *VMPQueryServerMetricParameters {
  s.StatTime = &v
  return s
}

func (s *VMPQueryServerMetricParameters) SetCarrier(v string) *VMPQueryServerMetricParameters {
  s.Carrier = &v
  return s
}

type VMPQueryServerMetricRequestHeader struct {
}

func (s VMPQueryServerMetricRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s VMPQueryServerMetricRequestHeader) GoString() string {
  return s.String()
}

type VMPQueryServerMetricResponseHeader struct {
}

func (s VMPQueryServerMetricResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s VMPQueryServerMetricResponseHeader) GoString() string {
  return s.String()
}




