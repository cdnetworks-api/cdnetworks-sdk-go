package reportflow

// This file is auto-generated, don't edit it. Thanks.
import (
  "github.com/alibabacloud-go/tea/tea"
)

type GetTrafficRequestsAndPeakBandwidthForMultiDomainsRequest struct {
  // {"en":"Start time:\n1. The time format is yyyy-MM-ddTHH:mm:ss±HH:mm. Please note: ±HH:mm is the time zone offset, which can be adjusted according to your data needs, for example, +00:00 represents UTC time, +08:00 represents East 8th District, and -05:00 represents West 5th District. 2024-01-15T10:30:45+00:00 means UTC time January 15, 2024 10:30:45 AM;\n2. No bigger than the current time;\n3. Data in the last 183 days at most can be queried.","zh_CN":"开始时间：\n1.时间格式为 yyyy-MM-ddTHH:mm:ss±HH:mm。请注意：±HH:mm 为时区偏移量，可根据您的数据需要进行调整，例如 +00:00 代表 UTC 时间，+08:00 代表东八区，-05:00 代表西五区。2024-01-15T10:30:45+00:00，表示UTC 时间 2024 年 1 月 15 日上午 10 点 30 分 45 秒；\n2.不能大于当前时间;\n3.最多可获取最近半年（183天）的数据。"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty"`
  // {"en":"End time:\n1. The time format is yyyy-MM-ddTHH:mm:ss±HH:mm. Please note: ±HH:mm is the time zone offset, which can be adjusted according to your data needs, for example, +00:00 represents UTC time, +08:00 represents East 8th District, and -05:00 represents West 5th District. 2024-01-15T10:30:45+00:00 means UTC time January 15, 2024 10:30:45 AM;\n2. End time should be greater than start time. If the end time is greater than current time, current time will be used;\n3. If both fields of dataFrom and dateTo are left empty, then data in the last 24 hours will be queried by default, if only one field is filled in and one is left empty, then exception will be occur;\n4. Allowable maximum time range for query: 1 day, means the period between dateFrom to dateTo should not exceed 1 day (can be adjusted by contacting technical support up to 31 days).","zh_CN":"结束时间：\n1.时间格式为 yyyy-MM-ddTHH:mm:ss±HH:mm。请注意：±HH:mm 为时区偏移量，可根据您的数据需要进行调整，例如 +00:00 代表 UTC 时间，+08:00 代表东八区，-05:00 代表西五区。2024-01-15T10:30:45+00:00，表示UTC 时间 2024 年 1 月 15 日上午 10 点 30 分 45 秒;\n2.结束时间需大于开始时间，结束时间如果大于当前时间，取当前时间;\n3.dateFrom，dateTo二者都未传，默认查询过去的24小时，如仅有一个未传，抛异常;\n4.允许查询最大时间间隔：1天，即dateFrom和dateTo相差不能超过1天（可联系技术支持调整，最大不超过31天）。"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty"`
  // {"en":"Domain name:\n1. The default upper limit to domains that can be entered is 1000 (Contact technical support to adjust);\n2. All domains under the account are queried if this input parameter is not specified, but if the number of domains under the account exceeds limits, no query will be done (Error).","zh_CN":"域名：\n1.可传递域名数量上限默认1000个（可联系技术支持调整）;\n2.未传递该入参时查询账号下所有域名，但当账号下域名数量超过限制时不可查询（报错）。"}
  Domain []*string `json:"domain,omitempty" xml:"domain,omitempty" type:"Repeated"`
  // {"en":"Group keywords:\n1. By default, group data will be displayed;\n2. If there are keywords entered, value details shall be displayed by keywords;\nIf domain is specified to groupBy, it means results are returned according to domains;\n3. Only domain can be specified.","zh_CN":"分组关键词：\n1.默认聚合展示；\n2.传入关键词则代表需要按照关键词对应的值展示明细；\n例如groupBy传domain，则代表返回按照domain明细展开。\n3.只能传递domain。"}
  GroupBy []*string `json:"groupBy,omitempty" xml:"groupBy,omitempty" type:"Repeated"`
}

func (s GetTrafficRequestsAndPeakBandwidthForMultiDomainsRequest) String() string {
  return tea.Prettify(s)
}

func (s GetTrafficRequestsAndPeakBandwidthForMultiDomainsRequest) GoString() string {
  return s.String()
}

func (s *GetTrafficRequestsAndPeakBandwidthForMultiDomainsRequest) SetDateFrom(v string) *GetTrafficRequestsAndPeakBandwidthForMultiDomainsRequest {
  s.DateFrom = &v
  return s
}

func (s *GetTrafficRequestsAndPeakBandwidthForMultiDomainsRequest) SetDateTo(v string) *GetTrafficRequestsAndPeakBandwidthForMultiDomainsRequest {
  s.DateTo = &v
  return s
}

func (s *GetTrafficRequestsAndPeakBandwidthForMultiDomainsRequest) SetDomain(v []*string) *GetTrafficRequestsAndPeakBandwidthForMultiDomainsRequest {
  s.Domain = v
  return s
}

func (s *GetTrafficRequestsAndPeakBandwidthForMultiDomainsRequest) SetGroupBy(v []*string) *GetTrafficRequestsAndPeakBandwidthForMultiDomainsRequest {
  s.GroupBy = v
  return s
}

type GetTrafficRequestsAndPeakBandwidthForMultiDomainsRequestHeader struct {
}

func (s GetTrafficRequestsAndPeakBandwidthForMultiDomainsRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetTrafficRequestsAndPeakBandwidthForMultiDomainsRequestHeader) GoString() string {
  return s.String()
}

type GetTrafficRequestsAndPeakBandwidthForMultiDomainsPaths struct {
}

func (s GetTrafficRequestsAndPeakBandwidthForMultiDomainsPaths) String() string {
  return tea.Prettify(s)
}

func (s GetTrafficRequestsAndPeakBandwidthForMultiDomainsPaths) GoString() string {
  return s.String()
}

type GetTrafficRequestsAndPeakBandwidthForMultiDomainsParameters struct {
}

func (s GetTrafficRequestsAndPeakBandwidthForMultiDomainsParameters) String() string {
  return tea.Prettify(s)
}

func (s GetTrafficRequestsAndPeakBandwidthForMultiDomainsParameters) GoString() string {
  return s.String()
}

type GetTrafficRequestsAndPeakBandwidthForMultiDomainsResponse struct {
  // {"en":"","zh_CN":""}
  Result []*GetTrafficRequestsAndPeakBandwidthForMultiDomainsResponseResult `json:"result,omitempty" xml:"result,omitempty" require:"true" type:"Repeated"`
}

func (s GetTrafficRequestsAndPeakBandwidthForMultiDomainsResponse) String() string {
  return tea.Prettify(s)
}

func (s GetTrafficRequestsAndPeakBandwidthForMultiDomainsResponse) GoString() string {
  return s.String()
}

func (s *GetTrafficRequestsAndPeakBandwidthForMultiDomainsResponse) SetResult(v []*GetTrafficRequestsAndPeakBandwidthForMultiDomainsResponseResult) *GetTrafficRequestsAndPeakBandwidthForMultiDomainsResponse {
  s.Result = v
  return s
}

type GetTrafficRequestsAndPeakBandwidthForMultiDomainsResponseResult struct     {
  // {"en":"Domain","zh_CN":"域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {"en":"Total traffic, unit: MB, retain two decimals, example (74099.92)","zh_CN":"总流量，单位:MB ，保留2位小数，示例 ( 74099.92 )"}
  TotalTraffic *string `json:"totalTraffic,omitempty" xml:"totalTraffic,omitempty" require:"true"`
  // {"en":"Total requests","zh_CN":"总请求数"}
  TotalRequests *string `json:"totalRequests,omitempty" xml:"totalRequests,omitempty" require:"true"`
  // {"en":"Peak bandwidth, unit: Mbps, retain two decimals, example (74099.92)","zh_CN":"峰值带宽，单位: Mbps，保留2位小数，示例 （931556.21）"}
  PeakBandwidth *string `json:"peakBandwidth,omitempty" xml:"peakBandwidth,omitempty" require:"true"`
  // {"en":"Peak time of bandwidth, example (2019-02-13 18:01)","zh_CN":"峰值时间，示例（2019-02-13 18:01）"}
  PeakBandwidthTimestamp *string `json:"peakBandwidthTimestamp,omitempty" xml:"peakBandwidthTimestamp,omitempty" require:"true"`
  // {"en":"Peak requests","zh_CN":"请求数峰值"}
  PeakRequests *string `json:"peakRequests,omitempty" xml:"peakRequests,omitempty" require:"true"`
  // {"en":"Peak time of requests","zh_CN":"请求数峰值时间"}
  PeakRequestsTimestamp *string `json:"peakRequestsTimestamp,omitempty" xml:"peakRequestsTimestamp,omitempty" require:"true"`
  // {"en":"","zh_CN":""}
  DataSeries []*GetTrafficRequestsAndPeakBandwidthForMultiDomainsResponseResultDataSeries `json:"dataSeries,omitempty" xml:"dataSeries,omitempty" require:"true" type:"Repeated"`
}

func (s GetTrafficRequestsAndPeakBandwidthForMultiDomainsResponseResult) String() string {
  return tea.Prettify(s)
}

func (s GetTrafficRequestsAndPeakBandwidthForMultiDomainsResponseResult) GoString() string {
  return s.String()
}

func (s *GetTrafficRequestsAndPeakBandwidthForMultiDomainsResponseResult) SetDomain(v string) *GetTrafficRequestsAndPeakBandwidthForMultiDomainsResponseResult {
  s.Domain = &v
  return s
}

func (s *GetTrafficRequestsAndPeakBandwidthForMultiDomainsResponseResult) SetTotalTraffic(v string) *GetTrafficRequestsAndPeakBandwidthForMultiDomainsResponseResult {
  s.TotalTraffic = &v
  return s
}

func (s *GetTrafficRequestsAndPeakBandwidthForMultiDomainsResponseResult) SetTotalRequests(v string) *GetTrafficRequestsAndPeakBandwidthForMultiDomainsResponseResult {
  s.TotalRequests = &v
  return s
}

func (s *GetTrafficRequestsAndPeakBandwidthForMultiDomainsResponseResult) SetPeakBandwidth(v string) *GetTrafficRequestsAndPeakBandwidthForMultiDomainsResponseResult {
  s.PeakBandwidth = &v
  return s
}

func (s *GetTrafficRequestsAndPeakBandwidthForMultiDomainsResponseResult) SetPeakBandwidthTimestamp(v string) *GetTrafficRequestsAndPeakBandwidthForMultiDomainsResponseResult {
  s.PeakBandwidthTimestamp = &v
  return s
}

func (s *GetTrafficRequestsAndPeakBandwidthForMultiDomainsResponseResult) SetPeakRequests(v string) *GetTrafficRequestsAndPeakBandwidthForMultiDomainsResponseResult {
  s.PeakRequests = &v
  return s
}

func (s *GetTrafficRequestsAndPeakBandwidthForMultiDomainsResponseResult) SetPeakRequestsTimestamp(v string) *GetTrafficRequestsAndPeakBandwidthForMultiDomainsResponseResult {
  s.PeakRequestsTimestamp = &v
  return s
}

func (s *GetTrafficRequestsAndPeakBandwidthForMultiDomainsResponseResult) SetDataSeries(v []*GetTrafficRequestsAndPeakBandwidthForMultiDomainsResponseResultDataSeries) *GetTrafficRequestsAndPeakBandwidthForMultiDomainsResponseResult {
  s.DataSeries = v
  return s
}

type GetTrafficRequestsAndPeakBandwidthForMultiDomainsResponseResultDataSeries struct     {
  // {"en":"Timestamp1. When the data granularity of query is 5m, the format is yyyy-mm-dd HH: MM;Each time slice data value represents the data value in the previous time granularity range.The time slice at the beginning of the day is yyyy-mm-dd 00:05, and the last time slice is (yyyy-mm-dd +1) 00:00;2. Return the time slice of start time and end time.","zh_CN":"时间片\n1.查询的数据粒度为5m时，格式为yyyy-MM-dd HH:mm；每一个时间片数据值代表的是前一个时间粒度范围内的数据值。一天开始的时间片是yyyy-MM-dd 00:05，最后一个时间片是（yyyy-MM-dd+1） 00:00;\n2.返回开始时间和结束时间包含的时间片。"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {"en":"Traffic value:\nUnit: MB, retain two decimals,example (34099.92)","zh_CN":"流量值：\n\n单位MB，保留2位小数，示例 (34099.92)"}
  Traffic *string `json:"traffic,omitempty" xml:"traffic,omitempty" require:"true"`
  // {"en":"Bandwidth, unit: Mbps, retain two decimals, example (931556.21)","zh_CN":"带宽值，单位: Mbps，保留2位小数，示例 （931556.21）"}
  Bandwidth *string `json:"bandwidth,omitempty" xml:"bandwidth,omitempty" require:"true"`
  // {"en":"Requests","zh_CN":"请求数"}
  Requests *string `json:"requests,omitempty" xml:"requests,omitempty" require:"true"`
}

func (s GetTrafficRequestsAndPeakBandwidthForMultiDomainsResponseResultDataSeries) String() string {
  return tea.Prettify(s)
}

func (s GetTrafficRequestsAndPeakBandwidthForMultiDomainsResponseResultDataSeries) GoString() string {
  return s.String()
}

func (s *GetTrafficRequestsAndPeakBandwidthForMultiDomainsResponseResultDataSeries) SetTimestamp(v string) *GetTrafficRequestsAndPeakBandwidthForMultiDomainsResponseResultDataSeries {
  s.Timestamp = &v
  return s
}

func (s *GetTrafficRequestsAndPeakBandwidthForMultiDomainsResponseResultDataSeries) SetTraffic(v string) *GetTrafficRequestsAndPeakBandwidthForMultiDomainsResponseResultDataSeries {
  s.Traffic = &v
  return s
}

func (s *GetTrafficRequestsAndPeakBandwidthForMultiDomainsResponseResultDataSeries) SetBandwidth(v string) *GetTrafficRequestsAndPeakBandwidthForMultiDomainsResponseResultDataSeries {
  s.Bandwidth = &v
  return s
}

func (s *GetTrafficRequestsAndPeakBandwidthForMultiDomainsResponseResultDataSeries) SetRequests(v string) *GetTrafficRequestsAndPeakBandwidthForMultiDomainsResponseResultDataSeries {
  s.Requests = &v
  return s
}

type GetTrafficRequestsAndPeakBandwidthForMultiDomainsResponseHeader struct {
}

func (s GetTrafficRequestsAndPeakBandwidthForMultiDomainsResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetTrafficRequestsAndPeakBandwidthForMultiDomainsResponseHeader) GoString() string {
  return s.String()
}




type GetOriginTrafficAndRequestsForMultiDomainsRequest struct {
  // {"en":"Start time:\n1. The time format is yyyy-MM-ddTHH:mm:ss±HH:mm. Please note: ±HH:mm is the time zone offset, which can be adjusted according to your data needs, for example, +00:00 represents UTC time, +08:00 represents East 8th District, and -05:00 represents West 5th District. 2024-01-15T10:30:45+00:00 means UTC time January 15, 2024 10:30:45 AM;\n2. Not greater than the current time;\n3. The data can be queried is the last week (183 days).","zh_CN":"开始时间：\n1.时间格式为 yyyy-MM-ddTHH:mm:ss±HH:mm。请注意：±HH:mm 为时区偏移量，可根据您的数据需要进行调整，例如 +00:00 代表 UTC 时间，+08:00 代表东八区，-05:00 代表西五区。2024-01-15T10:30:45+00:00，表示UTC 时间 2024 年 1 月 15 日上午 10 点 30 分 45 秒；\n2.不能大于当前时间；\n3.最多可获取最近半年（183天）的数据。"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty"`
  // {"en":"End time:\n1. The time format is yyyy-MM-ddTHH:mm:ss±HH:mm. Please note: ±HH:mm is the time zone offset, which can be adjusted according to your data needs, for example, +00:00 represents UTC time, +08:00 represents East 8th District, and -05:00 represents West 5th District. 2024-01-15T10:30:45+00:00 means UTC time January 15, 2024 10:30:45 AM;\n2. End time should be greater than start time. If the end time is greater than current time, current time will be used.\n3. If both fields of dataFrom and dateTo are left empty, then data in the last 24 hours will be queried by default; if one field is filled and one is left empty, then exception will occur.\n4. Maximum query time interval allowed: 1 day by default, that is, the difference between dateFrom and dateTo cannot exceed 1 day (you can contact technical support to adjust it, the maximum adjustment is 31 days)","zh_CN":"结束时间：\n1.时间格式为 yyyy-MM-ddTHH:mm:ss±HH:mm。请注意：±HH:mm 为时区偏移量，可根据您的数据需要进行调整，例如 +00:00 代表 UTC 时间，+08:00 代表东八区，-05:00 代表西五区。2024-01-15T10:30:45+00:00，表示UTC 时间 2024 年 1 月 15 日上午 10 点 30 分 45 秒;\n2.结束时间需大于开始时间，结束时间如果大于当前时间，取当前时间。\n3.dateFrom，dateTo二者都未传，默认查询过去的24小时；如仅有一个未传，抛异常\n4.允许查询最大时间间隔：默认1天，即dateFrom和dateTo相差不能超过1天（可联系技术支持调整，最大调整到31天）。"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty"`
  // {"defaultValue":"5m","en":"Data granularity:\n1. Support 5m (granularity of 5 minutes) and 1m","zh_CN":"数据粒度：\n1、支持5m（5分钟）、1m（1分钟）","exampleValue":"1m,5m"}
  Granularity *string `json:"granularity,omitempty" xml:"granularity,omitempty"`
  // {"en":"Domain name:\n1. The default upper limit of domains that can be entered is 200 (if you want to adjust, please, contact technical support);\n2. All domains under the account will be queried if this input parameter is not specified. But if the number of domains under the account exceeds the limits, no query will be executed (Error)","zh_CN":"域名：\n1、可传递域名数量上限默认为200个（可联系技术支持调整）；\n2、未传递该入参时查询账号下所有域名，但当账号下域名数量超过限制时不可查询（报错）。"}
  Domain []*string `json:"domain,omitempty" xml:"domain,omitempty" type:"Repeated"`
  // {"en":"Grouping keywords:\n1. By default, data will be displayed by group;\n2. If there are keywords entered, value details shall be displayed by keywords; If groupBy is specified as domain, it means the results are returned according to domains.\n3. Only domain can be specified","zh_CN":"分组关键词：\n1、默认聚合展示；\n2、传入关键词则代表需要按照关键词对应的值展示明细；\n例如groupBy传domain，则代表返回按照domain明细展开。\n3、只能传递domain。","exampleValue":"domain"}
  GroupBy []*string `json:"groupBy,omitempty" xml:"groupBy,omitempty" type:"Repeated"`
  // {"defaultValue":"0","en":"Input 0 returns all data, input 1 only returns source site data","zh_CN":"入参 0 则返回全部数据，入参 1 则只返回回源站数据","exampleValue":"0,1"}
  OriginOnly *int `json:"originOnly,omitempty" xml:"originOnly,omitempty"`
}

func (s GetOriginTrafficAndRequestsForMultiDomainsRequest) String() string {
  return tea.Prettify(s)
}

func (s GetOriginTrafficAndRequestsForMultiDomainsRequest) GoString() string {
  return s.String()
}

func (s *GetOriginTrafficAndRequestsForMultiDomainsRequest) SetDateFrom(v string) *GetOriginTrafficAndRequestsForMultiDomainsRequest {
  s.DateFrom = &v
  return s
}

func (s *GetOriginTrafficAndRequestsForMultiDomainsRequest) SetDateTo(v string) *GetOriginTrafficAndRequestsForMultiDomainsRequest {
  s.DateTo = &v
  return s
}

func (s *GetOriginTrafficAndRequestsForMultiDomainsRequest) SetGranularity(v string) *GetOriginTrafficAndRequestsForMultiDomainsRequest {
  s.Granularity = &v
  return s
}

func (s *GetOriginTrafficAndRequestsForMultiDomainsRequest) SetDomain(v []*string) *GetOriginTrafficAndRequestsForMultiDomainsRequest {
  s.Domain = v
  return s
}

func (s *GetOriginTrafficAndRequestsForMultiDomainsRequest) SetGroupBy(v []*string) *GetOriginTrafficAndRequestsForMultiDomainsRequest {
  s.GroupBy = v
  return s
}

func (s *GetOriginTrafficAndRequestsForMultiDomainsRequest) SetOriginOnly(v int) *GetOriginTrafficAndRequestsForMultiDomainsRequest {
  s.OriginOnly = &v
  return s
}

type GetOriginTrafficAndRequestsForMultiDomainsRequestHeader struct {
}

func (s GetOriginTrafficAndRequestsForMultiDomainsRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetOriginTrafficAndRequestsForMultiDomainsRequestHeader) GoString() string {
  return s.String()
}

type GetOriginTrafficAndRequestsForMultiDomainsPaths struct {
}

func (s GetOriginTrafficAndRequestsForMultiDomainsPaths) String() string {
  return tea.Prettify(s)
}

func (s GetOriginTrafficAndRequestsForMultiDomainsPaths) GoString() string {
  return s.String()
}

type GetOriginTrafficAndRequestsForMultiDomainsParameters struct {
}

func (s GetOriginTrafficAndRequestsForMultiDomainsParameters) String() string {
  return tea.Prettify(s)
}

func (s GetOriginTrafficAndRequestsForMultiDomainsParameters) GoString() string {
  return s.String()
}

type GetOriginTrafficAndRequestsForMultiDomainsResponse struct {
  // {"en":"","zh_CN":""}
  Result []*GetOriginTrafficAndRequestsForMultiDomainsResponseResult `json:"result,omitempty" xml:"result,omitempty" require:"true" type:"Repeated"`
}

func (s GetOriginTrafficAndRequestsForMultiDomainsResponse) String() string {
  return tea.Prettify(s)
}

func (s GetOriginTrafficAndRequestsForMultiDomainsResponse) GoString() string {
  return s.String()
}

func (s *GetOriginTrafficAndRequestsForMultiDomainsResponse) SetResult(v []*GetOriginTrafficAndRequestsForMultiDomainsResponseResult) *GetOriginTrafficAndRequestsForMultiDomainsResponse {
  s.Result = v
  return s
}

type GetOriginTrafficAndRequestsForMultiDomainsResponseResult struct     {
  // {"en":"Domain","zh_CN":"域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {"en":"Total traffic, Unit:  MB, retain two decimals, example (74099.92)","zh_CN":"总流量，单位MB，保留2位小数，示例：(74099.92)"}
  TotalTraffic *string `json:"totalTraffic,omitempty" xml:"totalTraffic,omitempty" require:"true"`
  // {"en":"Total requests","zh_CN":"总请求数"}
  TotalRequests *string `json:"totalRequests,omitempty" xml:"totalRequests,omitempty" require:"true"`
  // {"en":"Peak of Requests","zh_CN":"请求数峰值"}
  PeakRequests *string `json:"peakRequests,omitempty" xml:"peakRequests,omitempty" require:"true"`
  // {"en":"Peak time of Request","zh_CN":"请求数峰值时间"}
  PeakRequestsTimestamp *string `json:"peakRequestsTimestamp,omitempty" xml:"peakRequestsTimestamp,omitempty" require:"true"`
  // {"en":"Peak bandwidth, Unit: Mbps, retain two decimals, example (74099.92)","zh_CN":"带宽峰值，单位Mbps，保留2位小数，示例 （931556.21）"}
  PeakBandwidth *string `json:"peakBandwidth,omitempty" xml:"peakBandwidth,omitempty" require:"true"`
  // {"en":"Peak time of Bandwidth","zh_CN":"带宽峰值时间"}
  PeakBandwidthTimestamp *string `json:"peakBandwidthTimestamp,omitempty" xml:"peakBandwidthTimestamp,omitempty" require:"true"`
  // {"en":"","zh_CN":""}
  DataSeries []*GetOriginTrafficAndRequestsForMultiDomainsResponseResultDataSeries `json:"dataSeries,omitempty" xml:"dataSeries,omitempty" require:"true" type:"Repeated"`
}

func (s GetOriginTrafficAndRequestsForMultiDomainsResponseResult) String() string {
  return tea.Prettify(s)
}

func (s GetOriginTrafficAndRequestsForMultiDomainsResponseResult) GoString() string {
  return s.String()
}

func (s *GetOriginTrafficAndRequestsForMultiDomainsResponseResult) SetDomain(v string) *GetOriginTrafficAndRequestsForMultiDomainsResponseResult {
  s.Domain = &v
  return s
}

func (s *GetOriginTrafficAndRequestsForMultiDomainsResponseResult) SetTotalTraffic(v string) *GetOriginTrafficAndRequestsForMultiDomainsResponseResult {
  s.TotalTraffic = &v
  return s
}

func (s *GetOriginTrafficAndRequestsForMultiDomainsResponseResult) SetTotalRequests(v string) *GetOriginTrafficAndRequestsForMultiDomainsResponseResult {
  s.TotalRequests = &v
  return s
}

func (s *GetOriginTrafficAndRequestsForMultiDomainsResponseResult) SetPeakRequests(v string) *GetOriginTrafficAndRequestsForMultiDomainsResponseResult {
  s.PeakRequests = &v
  return s
}

func (s *GetOriginTrafficAndRequestsForMultiDomainsResponseResult) SetPeakRequestsTimestamp(v string) *GetOriginTrafficAndRequestsForMultiDomainsResponseResult {
  s.PeakRequestsTimestamp = &v
  return s
}

func (s *GetOriginTrafficAndRequestsForMultiDomainsResponseResult) SetPeakBandwidth(v string) *GetOriginTrafficAndRequestsForMultiDomainsResponseResult {
  s.PeakBandwidth = &v
  return s
}

func (s *GetOriginTrafficAndRequestsForMultiDomainsResponseResult) SetPeakBandwidthTimestamp(v string) *GetOriginTrafficAndRequestsForMultiDomainsResponseResult {
  s.PeakBandwidthTimestamp = &v
  return s
}

func (s *GetOriginTrafficAndRequestsForMultiDomainsResponseResult) SetDataSeries(v []*GetOriginTrafficAndRequestsForMultiDomainsResponseResultDataSeries) *GetOriginTrafficAndRequestsForMultiDomainsResponseResult {
  s.DataSeries = v
  return s
}

type GetOriginTrafficAndRequestsForMultiDomainsResponseResultDataSeries struct     {
  // {"en":"1. When the querying data granularity is 5m, then the format is yyyy-MM-dd HH:mm; the data value of every time slice represents the data value within the previous time granularity range. The first time slice of the day is yyyy-MM-dd 12:05 AM, and the last one is (yyyy-MM-dd+1) 00:00.\n2. Return the time slices that contained in start time and in end time.","zh_CN":"1.查询的数据粒度为5m时，格式为yyyy-MM-dd HH:mm；每一个时间片数据值代表的是前一个时间粒度范围内的数据值。一天开始的时间片是yyyy-MM-dd 00:05，最后一个时间片是（yyyy-MM-dd+1）00:00。\n2.返回开始时间和结束时间包含的时间片。"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {"en":"Traffic value:\nUnit:  MB, retain two decimals, example (34099.92)","zh_CN":"流量值：\n\n单位MB，保留2位小数，示例：(34099.92)"}
  Traffic *string `json:"traffic,omitempty" xml:"traffic,omitempty" require:"true"`
  // {"en":"Bandwidth value:\n\nUnit: Mbps, retain two decimals, example: (31556.21)","zh_CN":"带宽值：\n\n单位Mbps，保留2位小数, 示例：(31556.21)"}
  Bandwidth *string `json:"bandwidth,omitempty" xml:"bandwidth,omitempty" require:"true"`
  // {"en":"Total number of requests","zh_CN":"请求数"}
  Requests *string `json:"requests,omitempty" xml:"requests,omitempty" require:"true"`
}

func (s GetOriginTrafficAndRequestsForMultiDomainsResponseResultDataSeries) String() string {
  return tea.Prettify(s)
}

func (s GetOriginTrafficAndRequestsForMultiDomainsResponseResultDataSeries) GoString() string {
  return s.String()
}

func (s *GetOriginTrafficAndRequestsForMultiDomainsResponseResultDataSeries) SetTimestamp(v string) *GetOriginTrafficAndRequestsForMultiDomainsResponseResultDataSeries {
  s.Timestamp = &v
  return s
}

func (s *GetOriginTrafficAndRequestsForMultiDomainsResponseResultDataSeries) SetTraffic(v string) *GetOriginTrafficAndRequestsForMultiDomainsResponseResultDataSeries {
  s.Traffic = &v
  return s
}

func (s *GetOriginTrafficAndRequestsForMultiDomainsResponseResultDataSeries) SetBandwidth(v string) *GetOriginTrafficAndRequestsForMultiDomainsResponseResultDataSeries {
  s.Bandwidth = &v
  return s
}

func (s *GetOriginTrafficAndRequestsForMultiDomainsResponseResultDataSeries) SetRequests(v string) *GetOriginTrafficAndRequestsForMultiDomainsResponseResultDataSeries {
  s.Requests = &v
  return s
}

type GetOriginTrafficAndRequestsForMultiDomainsResponseHeader struct {
}

func (s GetOriginTrafficAndRequestsForMultiDomainsResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetOriginTrafficAndRequestsForMultiDomainsResponseHeader) GoString() string {
  return s.String()
}




type QueryTrafficForMultiDomainRequest struct {
  // {"en":"Domain list:
  //   Domain number limits can be adjusted depending on different accounts. The default value is 1000(if you want to adjust,please, contact technical support)", "zh_CN":"域名列表:
  //   1.域名个数限制根据账号可调,默认为1000个(可联系技术支持下单调整);"}
  QueryTrafficForMultiDomainDomainList *QueryTrafficForMultiDomainDomainList `json:"domain-list,omitempty" xml:"domain-list,omitempty" require:"true"`
}

func (s QueryTrafficForMultiDomainRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryTrafficForMultiDomainRequest) GoString() string {
  return s.String()
}

func (s *QueryTrafficForMultiDomainRequest) SetDomainList(v *QueryTrafficForMultiDomainDomainList) *QueryTrafficForMultiDomainRequest {
  s.QueryTrafficForMultiDomainDomainList = v
  return s
}

type QueryTrafficForMultiDomainDomainList struct {
  // {"en":"Domain", "zh_CN":"域名"}
  DomainName []*string `json:"domain-name,omitempty" xml:"domain-name,omitempty" require:"true" type:"Repeated"`
}

func (s QueryTrafficForMultiDomainDomainList) String() string {
  return tea.Prettify(s)
}

func (s QueryTrafficForMultiDomainDomainList) GoString() string {
  return s.String()
}

func (s *QueryTrafficForMultiDomainDomainList) SetDomainName(v []*string) *QueryTrafficForMultiDomainDomainList {
  s.DomainName = v
  return s
}

type QueryTrafficForMultiDomainResponse struct {
  // {"en":"Total traffic. Keep two digits of decimals. Unit: MB", "zh_CN":"总流量,保留2位小数,单位为MB"}
  FlowSummary *int `json:"flow-summary,omitempty" xml:"flow-summary,omitempty" require:"true"`
  // {"en":"flowData", "zh_CN":"流量数据"}
  FlowData []*QueryTrafficForMultiDomainResponseFlowData `json:"flow-data,omitempty" xml:"flow-data,omitempty" require:"true" type:"Repeated"`
}

func (s QueryTrafficForMultiDomainResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryTrafficForMultiDomainResponse) GoString() string {
  return s.String()
}

func (s *QueryTrafficForMultiDomainResponse) SetFlowSummary(v int) *QueryTrafficForMultiDomainResponse {
  s.FlowSummary = &v
  return s
}

func (s *QueryTrafficForMultiDomainResponse) SetFlowData(v []*QueryTrafficForMultiDomainResponseFlowData) *QueryTrafficForMultiDomainResponse {
  s.FlowData = v
  return s
}

type QueryTrafficForMultiDomainResponseFlowData struct     {
  // {"en":"Date:
  //   1.When the data query granularity is fiveminutes, the format is yyyy-MM-dd HH:mm; the data value of every time slice represents the data value within the previous time granularity range. The first time slice of the day is yyyy-MM-dd 00:05 AM, and the last one is yyyy-MM-dd 24:00.
  //   2.When the data query granularity is hourly, the format is yyyy-MM-dd HH; the data value of every time slice represents the data value within the previous time granularity range. The first time slice of the day is yyyy-MM-dd 00:01, and the last one is yyyy-MM-dd 24.
  //   3.When the data query granularity is daily, the format is yyyy-MM-dd; the data value of every time slice represents the value of the daily data.Return the time slice contained in start time and the time slice contained in end time", "zh_CN":"时间
  //   1.查询的数据粒度为fiveminutes时,格式为yyyy-MM-dd HH:mm;每一个时间片数据值代表的是前一个时间粒度范围内的数据值。一天开始的时间片是yyyy-MM-dd 00:05,最后一个时间片是yyyy-MM-dd 24:00。
  //   2.查询的数据粒度为hourly时,格式为yyyy-MM-dd HH;每一个时间片数据值代表的是前一个时间粒度范围内的数据值。一天开始的时间片是yyyy-MM-dd 01,最后一个时间片是yyyy-MM-dd 24。
  //   3.查询的数据粒度为daily时,格式为yyyy-MM-dd;每一个时间片数据值代表的该天内的数据值。
  //   4.返回开始时间和结束时间包含的时间片。"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {"en":"Traffic. Keep two digits of decimals. Unit: MB", "zh_CN":"流量,保留2位小数,单位为MB"}
  Flow *int `json:"flow,omitempty" xml:"flow,omitempty" require:"true"`
}

func (s QueryTrafficForMultiDomainResponseFlowData) String() string {
  return tea.Prettify(s)
}

func (s QueryTrafficForMultiDomainResponseFlowData) GoString() string {
  return s.String()
}

func (s *QueryTrafficForMultiDomainResponseFlowData) SetTimestamp(v string) *QueryTrafficForMultiDomainResponseFlowData {
  s.Timestamp = &v
  return s
}

func (s *QueryTrafficForMultiDomainResponseFlowData) SetFlow(v int) *QueryTrafficForMultiDomainResponseFlowData {
  s.Flow = &v
  return s
}

type QueryTrafficForMultiDomainPaths struct {
}

func (s QueryTrafficForMultiDomainPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryTrafficForMultiDomainPaths) GoString() string {
  return s.String()
}

type QueryTrafficForMultiDomainParameters struct {
  // {"en":"Start time
  // 1.The time format is yyyy-MM-ddTHH:mm:ss±HH:mm. Please note: ±HH:mm is the time zone offset, which can be adjusted according to your data needs, for example, +00:00 represents UTC time, +08:00 represents East 8th District, and -05:00 represents West 5th District. 2024-01-15T10:30:45+00:00 means UTC time January 15, 2024 10:30:45 AM;
  // 2.And smaller than the current time and 'dateTo';
  // 3.Period between 'dataFrom' and 'dateTo' cannot be longer than 31 days", "zh_CN":"开始时间
  // 1.时间格式为 yyyy-MM-ddTHH:mm:ss±HH:mm。请注意：±HH:mm 为时区偏移量，可根据您的数据需要进行调整，例如 +00:00 代表 UTC 时间，+08:00 代表东八区，-05:00 代表西五区。2024-01-15T10:30:45+00:00，表示UTC 时间 2024 年 1 月 15 日上午 10 点 30 分 45 秒;
  // 2.并且小于当前时间和dateTo;
  // 3.dateFrom和dateTo相差不能超过31天(可联系技术支持调整);4.只能查询最近2年内数据。"}
  Datefrom *string `json:"datefrom,omitempty" xml:"datefrom,omitempty" require:"true"`
  // {"en":"End time
  // 1.The time format is yyyy-MM-ddTHH:mm:ss±HH:mm. Please note: ±HH:mm is the time zone offset, which can be adjusted according to your data needs, for example, +00:00 represents UTC time, +08:00 represents East 8th District, and -05:00 represents West 5th District. 2024-01-15T10:30:45+00:00 means UTC time January 15, 2024 10:30:45 AM;
  // 2.Must be greater than 'dateFrom';
  // 3.If it's greater than the current time, then the current time is assigned as the value", "zh_CN":"结束时间
  // 1.时间格式为 yyyy-MM-ddTHH:mm:ss±HH:mm。请注意：±HH:mm 为时区偏移量，可根据您的数据需要进行调整，例如 +00:00 代表 UTC 时间，+08:00 代表东八区，-05:00 代表西五区。2024-01-15T10:30:45+00:00，表示UTC 时间 2024 年 1 月 15 日上午 10 点 30 分 45 秒;
  // 2.必须大于dateFrom;
  // 3.如果大于当前时间,则重新赋值为当前时间;"}
  Dateto *string `json:"dateto,omitempty" xml:"dateto,omitempty" require:"true"`
  // {"en":"Data granularity
  // 1.fiveminutes: five minutes, hourly: one hour, daily: one day;
  // 2.If not specified, daily is set as the default value;
  // 3.If fiveminutes is specified as the value, then data is returned in actual configured granularity when there is specific configuration to data collecting granularity for the customer.", "zh_CN":"数据粒度
  // 1.fiveminutes:5分钟,hourly:1小时,daily:1天;
  // 2.不传递,默认为daily;
  // 3.传递fiveminutes时,若客户数据采集粒度有特殊配置将按实际配置粒度返回。"}
  Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s QueryTrafficForMultiDomainParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryTrafficForMultiDomainParameters) GoString() string {
  return s.String()
}

func (s *QueryTrafficForMultiDomainParameters) SetDatefrom(v string) *QueryTrafficForMultiDomainParameters {
  s.Datefrom = &v
  return s
}

func (s *QueryTrafficForMultiDomainParameters) SetDateto(v string) *QueryTrafficForMultiDomainParameters {
  s.Dateto = &v
  return s
}

func (s *QueryTrafficForMultiDomainParameters) SetType(v string) *QueryTrafficForMultiDomainParameters {
  s.Type = &v
  return s
}

type QueryTrafficForMultiDomainRequestHeader struct {
}

func (s QueryTrafficForMultiDomainRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryTrafficForMultiDomainRequestHeader) GoString() string {
  return s.String()
}

type QueryTrafficForMultiDomainResponseHeader struct {
}

func (s QueryTrafficForMultiDomainResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryTrafficForMultiDomainResponseHeader) GoString() string {
  return s.String()
}




type QueryOutputTrafficUnderShieldPoPRequest struct {
  // {"en":"Start time:
  // 1.The format is yyyy-MM-ddTHH:mm:ss+08:00;
  // 2.Must be a time that is 183 days earlier than the current time, and the time must be earlier than the current time and dateTo;
  // 3.Period between dataFrom and dateTo cannot be longer than 7 days(technical support can be contacted to adjust);
  // 4.dateFrom and dateTo can be either both are specified or neither is specifies;
  // 5.If neither dateFrom nor dateTo is specified, then by default, data in the last 24 hour is queried", "zh_CN":"开始时间
  // 1.格式为yyyy-MM-ddTHH:mm:ss+08:00;
  // 2.必须大于当前时间-183天,并且小于当前时间和dateTo;
  // 3.dateFrom和dateTo相差不能超过7天(可联系技术支持调整);
  // 4.dateFrom和dateTo要么都传递,要么都不传递;
  // 5.dateFrom和dateTo都未传递,则默认查询过去24小时的数据"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty"`
  // {"en":"End time
  // 
  // 1.The format is yyyy-MM-ddTHH:mm:ss+08:00;
  // 2.Must be greater than dateFrom;
  // 3.If it&rsquo;s greater than the current time, then the current time is assigned as the value;", "zh_CN":"结束时间
  // 1.格式为yyyy-MM-ddTHH:mm:ss+08:00;
  // 2.必须大于dateFrom;
  // 3.如果大于当前时间,则重新赋值为当前时间;"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty"`
  // {"en":"Domain names, domain number limits can be adjusted depending on different accounts. The default value is 20", "zh_CN":"域名,域名个数限制根据账号可调,默认为20个"}
  Domain []*string `json:"domain,omitempty" xml:"domain,omitempty" require:"true" type:"Repeated"`
  // {"en":"Data granularity, support 5m (granularity of 5 minutes)", "zh_CN":"数据粒度,支持5m:5分钟粒度"}
  DataInterval *string `json:"dataInterval,omitempty" xml:"dataInterval,omitempty"`
  // {"en":"Group dimension
  // 
  // 1.The value can be selected is domain;
  // 2.The data is displayed according to the specified dimension;", "zh_CN":"分组维度
  // 1.可选值为domain;
  // 2.有传入则按照该维度展示明细数据;"}
  GroupBy []*string `json:"groupBy,omitempty" xml:"groupBy,omitempty" type:"Repeated"`
}

func (s QueryOutputTrafficUnderShieldPoPRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryOutputTrafficUnderShieldPoPRequest) GoString() string {
  return s.String()
}

func (s *QueryOutputTrafficUnderShieldPoPRequest) SetDateFrom(v string) *QueryOutputTrafficUnderShieldPoPRequest {
  s.DateFrom = &v
  return s
}

func (s *QueryOutputTrafficUnderShieldPoPRequest) SetDateTo(v string) *QueryOutputTrafficUnderShieldPoPRequest {
  s.DateTo = &v
  return s
}

func (s *QueryOutputTrafficUnderShieldPoPRequest) SetDomain(v []*string) *QueryOutputTrafficUnderShieldPoPRequest {
  s.Domain = v
  return s
}

func (s *QueryOutputTrafficUnderShieldPoPRequest) SetDataInterval(v string) *QueryOutputTrafficUnderShieldPoPRequest {
  s.DataInterval = &v
  return s
}

func (s *QueryOutputTrafficUnderShieldPoPRequest) SetGroupBy(v []*string) *QueryOutputTrafficUnderShieldPoPRequest {
  s.GroupBy = v
  return s
}

type QueryOutputTrafficUnderShieldPoPResponse struct {
  // {"en":"result", "zh_CN":"结果"}
  Result []*QueryOutputTrafficUnderShieldPoPResponseResult `json:"result,omitempty" xml:"result,omitempty" require:"true" type:"Repeated"`
}

func (s QueryOutputTrafficUnderShieldPoPResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryOutputTrafficUnderShieldPoPResponse) GoString() string {
  return s.String()
}

func (s *QueryOutputTrafficUnderShieldPoPResponse) SetResult(v []*QueryOutputTrafficUnderShieldPoPResponseResult) *QueryOutputTrafficUnderShieldPoPResponse {
  s.Result = v
  return s
}

type QueryOutputTrafficUnderShieldPoPResponseResult struct     {
  // {"en":"Domain", "zh_CN":"域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {"en":"Total traffic", "zh_CN":"总流量"}
  TotalFlow *string `json:"totalFlow,omitempty" xml:"totalFlow,omitempty" require:"true"`
  // {"en":"flowData", "zh_CN":"流量值数据"}
  FlowData []*QueryOutputTrafficUnderShieldPoPResponseResultFlowData `json:"flowData,omitempty" xml:"flowData,omitempty" require:"true" type:"Repeated"`
}

func (s QueryOutputTrafficUnderShieldPoPResponseResult) String() string {
  return tea.Prettify(s)
}

func (s QueryOutputTrafficUnderShieldPoPResponseResult) GoString() string {
  return s.String()
}

func (s *QueryOutputTrafficUnderShieldPoPResponseResult) SetDomain(v string) *QueryOutputTrafficUnderShieldPoPResponseResult {
  s.Domain = &v
  return s
}

func (s *QueryOutputTrafficUnderShieldPoPResponseResult) SetTotalFlow(v string) *QueryOutputTrafficUnderShieldPoPResponseResult {
  s.TotalFlow = &v
  return s
}

func (s *QueryOutputTrafficUnderShieldPoPResponseResult) SetFlowData(v []*QueryOutputTrafficUnderShieldPoPResponseResultFlowData) *QueryOutputTrafficUnderShieldPoPResponseResult {
  s.FlowData = v
  return s
}

type QueryOutputTrafficUnderShieldPoPResponseResultFlowData struct     {
  // {"en":"DateTime, the format is yyyy-MM-dd HH:mm; the data value of every time slice represents the data value within the previous time granularity range. The first time slice of the day is yyyy-MM-dd 12:05 AM, and the last one is (yyyy-MM-dd+1) 00:00.", "zh_CN":"时间,格式为yyyy-MM-dd HH:mm;每一个时间片数据值代表的是前一个时间粒度范围内的数据值。一天开始的时间片是yyyy-MM-dd &nbsp; 00:05,最后一个时间片是(yyyy-MM-dd+1)&nbsp;00:00。"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {"en":"Traffic unit is MB and 2 digits &nbsp; of decimals allowed", "zh_CN":"流量值,单位MB,保留2位小数"}
  Value *string `json:"value,omitempty" xml:"value,omitempty" require:"true"`
}

func (s QueryOutputTrafficUnderShieldPoPResponseResultFlowData) String() string {
  return tea.Prettify(s)
}

func (s QueryOutputTrafficUnderShieldPoPResponseResultFlowData) GoString() string {
  return s.String()
}

func (s *QueryOutputTrafficUnderShieldPoPResponseResultFlowData) SetTimestamp(v string) *QueryOutputTrafficUnderShieldPoPResponseResultFlowData {
  s.Timestamp = &v
  return s
}

func (s *QueryOutputTrafficUnderShieldPoPResponseResultFlowData) SetValue(v string) *QueryOutputTrafficUnderShieldPoPResponseResultFlowData {
  s.Value = &v
  return s
}

type QueryOutputTrafficUnderShieldPoPPaths struct {
}

func (s QueryOutputTrafficUnderShieldPoPPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryOutputTrafficUnderShieldPoPPaths) GoString() string {
  return s.String()
}

type QueryOutputTrafficUnderShieldPoPParameters struct {
}

func (s QueryOutputTrafficUnderShieldPoPParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryOutputTrafficUnderShieldPoPParameters) GoString() string {
  return s.String()
}

type QueryOutputTrafficUnderShieldPoPRequestHeader struct {
}

func (s QueryOutputTrafficUnderShieldPoPRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryOutputTrafficUnderShieldPoPRequestHeader) GoString() string {
  return s.String()
}

type QueryOutputTrafficUnderShieldPoPResponseHeader struct {
}

func (s QueryOutputTrafficUnderShieldPoPResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryOutputTrafficUnderShieldPoPResponseHeader) GoString() string {
  return s.String()
}




type GetTrafficByProtocolRequest struct {
  // {"en":"Start time:\n1.The time format is yyyy-MM-ddTHH:mm:ss±HH:mm. Please note: ±HH:mm is the time zone offset, which can be adjusted according to your data needs, for example, +00:00 represents UTC time, +08:00 represents East 8th District, and -05:00 represents West 5th District. 2024-01-15T10:30:45+00:00 means UTC time January 15, 2024 10:30:45 AM;\n2.Must be no more than 183 days prior to the current time, and must be earlier than both the current time and `dateTo`.\n3.The period between `dateFrom` and `dateTo` cannot exceed 7 days (contact technical support for adjustments).\n4.Either both `dateFrom` and `dateTo` must be specified, or neither should be specified.\n5.If neither dateFrom nor dateTo is specified, then by default, data for the last 24 hours is queried.","zh_CN":"开始时间\n1.时间格式为 yyyy-MM-ddTHH:mm:ss±HH:mm。请注意：±HH:mm 为时区偏移量，可根据您的数据需要进行调整，例如 +00:00 代表 UTC 时间，+08:00 代表东八区，-05:00 代表西五区。2024-01-15T10:30:45+00:00，表示UTC 时间 2024 年 1 月 15 日上午 10 点 30 分 45 秒;\n2.必须大于当前时间-183天,并且小于当前时间和dateTo;\n3.dateFrom和dateTo相差不能超过7天(可联系技术支持调整);\n4.dateFrom和dateTo要么都传递,要么都不传递;\n5.dateFrom和dateTo都未传递,则默认查询过去24小时的数据"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty"`
  // {"en":"End time:\n1.The time format is yyyy-MM-ddTHH:mm:ss±HH:mm. Please note: ±HH:mm is the time zone offset, which can be adjusted according to your data needs, for example, +00:00 represents UTC time, +08:00 represents East 8th District, and -05:00 represents West 5th District. 2024-01-15T10:30:45+00:00 means UTC time January 15, 2024 10:30:45 AM.\n2.Must be greater than dateFrom.\n3.If it is greater than the current time, it will be reset to the current time.","zh_CN":"结束时间\n1.时间格式为 yyyy-MM-ddTHH:mm:ss±HH:mm。请注意：±HH:mm 为时区偏移量，可根据您的数据需要进行调整，例如 +00:00 代表 UTC 时间，+08:00 代表东八区，-05:00 代表西五区。2024-01-15T10:30:45+00:00，表示UTC 时间 2024 年 1 月 15 日上午 10 点 30 分 45 秒;\n2.必须大于dateFrom;\n3.如果大于当前时间,则重新赋值为当前时间;"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty"`
  // {"en":"Domain names. The number of allowed domains can be adjusted based on the account type, with a default limit of 20.","zh_CN":"域名,域名个数限制根据账号可调,默认为20个"}
  Domain []*string `json:"domain,omitempty" xml:"domain,omitempty" require:"true" type:"Repeated"`
  // {"defaultValue":"5m","en":"**Data granularity**\n1m: 1minute\n5m: 5minutes\n1h: 1hour\n1d: 1day","zh_CN":"数据粒度：\n1m：1分钟粒度\n5m：5分钟粒度\n1h：1小时粒度\n1d：1天粒度","exampleValue":"1m,5m,1h,1d"}
  Granularity *string `json:"granularity,omitempty" xml:"granularity,omitempty"`
  // {"defaultValue":"https","en":"Transmission protocol:\n1.Options: http, https.\n2.If no value is specified, https is used as the default.\n3.If http is queried, `httpFlowData` is displayed in the response; if https is queried, `httpsFlowData` is displayed.","zh_CN":"传输协议\n1.可选值为http、https;\n2.不传默认查询https;\n3.查询http时出参展示httpFlowData,查询https时出参展示httpsFlowData;","exampleValue":"http,https"}
  ProtocolType *string `json:"protocolType,omitempty" xml:"protocolType,omitempty"`
  // {"en":"Group dimension:\n1.Only 'domain' is a valid value.\n2.If provided, detailed data will be displayed according to this dimension.","zh_CN":"分组维度\n1.可选值为domain;\n2.有传入则按照该维度展示明细数据;","exampleValue":"domain"}
  GroupBy []*string `json:"groupBy,omitempty" xml:"groupBy,omitempty" type:"Repeated"`
}

func (s GetTrafficByProtocolRequest) String() string {
  return tea.Prettify(s)
}

func (s GetTrafficByProtocolRequest) GoString() string {
  return s.String()
}

func (s *GetTrafficByProtocolRequest) SetDateFrom(v string) *GetTrafficByProtocolRequest {
  s.DateFrom = &v
  return s
}

func (s *GetTrafficByProtocolRequest) SetDateTo(v string) *GetTrafficByProtocolRequest {
  s.DateTo = &v
  return s
}

func (s *GetTrafficByProtocolRequest) SetDomain(v []*string) *GetTrafficByProtocolRequest {
  s.Domain = v
  return s
}

func (s *GetTrafficByProtocolRequest) SetGranularity(v string) *GetTrafficByProtocolRequest {
  s.Granularity = &v
  return s
}

func (s *GetTrafficByProtocolRequest) SetProtocolType(v string) *GetTrafficByProtocolRequest {
  s.ProtocolType = &v
  return s
}

func (s *GetTrafficByProtocolRequest) SetGroupBy(v []*string) *GetTrafficByProtocolRequest {
  s.GroupBy = v
  return s
}

type GetTrafficByProtocolRequestHeader struct {
}

func (s GetTrafficByProtocolRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetTrafficByProtocolRequestHeader) GoString() string {
  return s.String()
}

type GetTrafficByProtocolPaths struct {
}

func (s GetTrafficByProtocolPaths) String() string {
  return tea.Prettify(s)
}

func (s GetTrafficByProtocolPaths) GoString() string {
  return s.String()
}

type GetTrafficByProtocolParameters struct {
}

func (s GetTrafficByProtocolParameters) String() string {
  return tea.Prettify(s)
}

func (s GetTrafficByProtocolParameters) GoString() string {
  return s.String()
}

type GetTrafficByProtocolResponse struct {
  // {"en":"The query result set.","zh_CN":"结果"}
  Result []*GetTrafficByProtocolResponseResult `json:"result,omitempty" xml:"result,omitempty" require:"true" type:"Repeated"`
}

func (s GetTrafficByProtocolResponse) String() string {
  return tea.Prettify(s)
}

func (s GetTrafficByProtocolResponse) GoString() string {
  return s.String()
}

func (s *GetTrafficByProtocolResponse) SetResult(v []*GetTrafficByProtocolResponseResult) *GetTrafficByProtocolResponse {
  s.Result = v
  return s
}

type GetTrafficByProtocolResponseResult struct     {
  // {"en":"The domain name.","zh_CN":"域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {"en":"Time-series data list.","zh_CN":"时间序列数据列表"}
  DataSeries []*GetTrafficByProtocolResponseResultDataSeries `json:"dataSeries,omitempty" xml:"dataSeries,omitempty" require:"true" type:"Repeated"`
}

func (s GetTrafficByProtocolResponseResult) String() string {
  return tea.Prettify(s)
}

func (s GetTrafficByProtocolResponseResult) GoString() string {
  return s.String()
}

func (s *GetTrafficByProtocolResponseResult) SetDomain(v string) *GetTrafficByProtocolResponseResult {
  s.Domain = &v
  return s
}

func (s *GetTrafficByProtocolResponseResult) SetDataSeries(v []*GetTrafficByProtocolResponseResultDataSeries) *GetTrafficByProtocolResponseResult {
  s.DataSeries = v
  return s
}

type GetTrafficByProtocolResponseResultDataSeries struct     {
  // {"en":"Timestamp, in yyyy-MM-dd HH:mm format. Each timestamp's data value represents the aggregated data within the preceding time granularity interval. The first timestamp for a day is yyyy-MM-dd 00:05, and the last is (yyyy-MM-dd+1) 00:00.","zh_CN":"时间,格式为yyyy-MM-dd HH:mm;每一个时间片数据值代表的是前一个时间粒度范围内的数据值。一天开始的时间片是yyyy-MM-dd 00:05,最后一个时间片是(yyyy-MM-dd+1)00:00。"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {"en":"Traffic value:\nUnit: MB. Retains two decimal places.","zh_CN":"流量值：\n\n单位MB，保留2位小数"}
  Traffic *string `json:"traffic,omitempty" xml:"traffic,omitempty" require:"true"`
}

func (s GetTrafficByProtocolResponseResultDataSeries) String() string {
  return tea.Prettify(s)
}

func (s GetTrafficByProtocolResponseResultDataSeries) GoString() string {
  return s.String()
}

func (s *GetTrafficByProtocolResponseResultDataSeries) SetTimestamp(v string) *GetTrafficByProtocolResponseResultDataSeries {
  s.Timestamp = &v
  return s
}

func (s *GetTrafficByProtocolResponseResultDataSeries) SetTraffic(v string) *GetTrafficByProtocolResponseResultDataSeries {
  s.Traffic = &v
  return s
}

type GetTrafficByProtocolResponseHeader struct {
}

func (s GetTrafficByProtocolResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetTrafficByProtocolResponseHeader) GoString() string {
  return s.String()
}




type GetTrafficAndBandwidthByClientCountryRequest struct {
  // {"en":"Starting time\n\n1.The time format is yyyy-MM-ddTHH:mm:ss±HH:mm. Please note: ±HH:mm is the time zone offset, which can be adjusted according to your data needs, for example, +00:00 represents UTC time, +08:00 represents East 8th District, and -05:00 represents West 5th District. 2024-01-15T10:30:45+00:00 means UTC time January 15, 2024 10:30:45 AM; \n2. Cannot be greater than the current time\n3. Get up to the last six months (183 days) of data.","zh_CN":"开始时间\n1.时间格式为 yyyy-MM-ddTHH:mm:ss±HH:mm。请注意：±HH:mm 为时区偏移量，可根据您的数据需要进行调整，例如 +00:00 代表 UTC 时间，+08:00 代表东八区，-05:00 代表西五区。2024-01-15T10:30:45+00:00，表示UTC 时间 2024 年 1 月 15 日上午 10 点 30 分 45 秒; \n2. 不能大于当前时间\n3. 最多可获取最近半年(183天)的数据。"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty"`
  // {"en":"End Time:\n1.The time format is yyyy-MM-ddTHH:mm:ss±HH:mm. Please note: ±HH:mm is the time zone offset, which can be adjusted according to your data needs, for example, +00:00 represents UTC time, +08:00 represents East 8th District, and -05:00 represents West 5th District. 2024-01-15T10:30:45+00:00 means UTC time January 15, 2024 10:30:45 AM; \n2. The end time needs to be greater than the start time. If the end time is greater than the current time, the current time is taken.\n3. dateFrom, dateTo are not passed, the default query for the past 24 hours; if there is only one untransmitted, throw an exception\n4. Allow query maximum time interval: 7 days, that is, the difference between dateFrom and dateTo can&rsquo;t exceed 7 days (can contact technical support adjustment, up to 31 days).","zh_CN":"结束时间:\n1.时间格式为 yyyy-MM-ddTHH:mm:ss±HH:mm。请注意：±HH:mm 为时区偏移量，可根据您的数据需要进行调整，例如 +00:00 代表 UTC 时间，+08:00 代表东八区，-05:00 代表西五区。2024-01-15T10:30:45+00:00，表示UTC 时间 2024 年 1 月 15 日上午 10 点 30 分 45 秒; \n2. 结束时间需大于开始时间,结束时间如果大于当前时间,取当前时间。\n3. dateFrom,dateTo二者都未传,默认查询过去的24小时;如仅有一个未传,抛异常\n4. 允许查询最大时间间隔:7天,即dateFrom和dateTo相差不能超过7天(可联系技术支持调整，最长31天)。"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty"`
  // {"defaultValue":"5m","en":"Data granularity:\n1.Support 5m (5 minutes granularity),1d (1 day granularity)\n2.Do not pass the default to 5m","zh_CN":"数据粒度:\n1.支持5m(5分钟粒度),1d(天粒度)\n2.不传默认为5m","exampleValue":"5m,1d"}
  Granularity *string `json:"granularity,omitempty" xml:"granularity,omitempty"`
  // {"en":"Domains:\n1.Domain is not uploaded: Query all domain names of the account (More than 20 domains will error, you can contact technical support for adjustment)\n2.Domain is uploaded: Up to 20 domains are supported (you can contact technical support for adjustment)","zh_CN":"域名：\n1.未传递domain时：查询账号下所有全部域名(域名超过20个则报错，可联系技术支持调整)\n2.有传递domain时：域名最多支持传20个（可联系技术支持调整）"}
  Domain []*string `json:"domain,omitempty" xml:"domain,omitempty" require:"true" type:"Repeated"`
  // {"dictionary":"belong=BCS-CC-API|dict=regionCodeList","en":"Country code:\n\n1. If not specified, all countries and regions will be queried by default\n2. Support language request header Accept-Language, only supports zh-CN and en-US, the default is zh-CN. When Accept-Language is en-US, the country and region returned are all in English, otherwise they are returned in Chinese.","zh_CN":"国家地区代号:\n1.不传默认查询全部国家地区\n2. 支持语言请求头Accept-Language，只支持zh-CN、en-US，默认为zh-CN。Accept-Language：en-US时，国家地区返回都为英文，否则返回的为中文。"}
  CountryCode []*string `json:"countryCode,omitempty" xml:"countryCode,omitempty" type:"Repeated"`
  // {"en":"Grouping dimension:\n1.Can pass in single or multiple values. The aggregatedOversea and country cannot be passed at the same time;\n2.If there is an incoming, the detailed data will be displayed according to the dimension:\ndomain: Group display by domain name dimension;\ncountry: Group display by country dimension;\naggregatedOversea: Group display according to domestic and overseas dimensions.\n3.The result hierarchy is fixed in order, and the order of the parameters does not affect the order of the returned results. For example: 'groupBy': ['domain','country'] and 'groupBy': ['country','domain'] return the same result.","zh_CN":"分组维度：\n1.可传入单个或多个值,其中不能同时传aggregatedOversea 和 country;\n2.有传入则按照该维度展示明细数据:\n3.domain:按照域名维度进行分组展示;\n4.domain:country:按照国家维度进行分组展示;\n5.aggregatedOversea:按照国内 和 海外维度进行分组展示\n6.返回结果层级顺序固定,入参顺序不影响返回结果顺序。例如:groupBy: [domain,country]与groupBy: [country,domain]返回结果一样。","exampleValue":"domain,country,aggregatedOverse"}
  GroupBy []*string `json:"groupBy,omitempty" xml:"groupBy,omitempty" type:"Repeated"`
}

func (s GetTrafficAndBandwidthByClientCountryRequest) String() string {
  return tea.Prettify(s)
}

func (s GetTrafficAndBandwidthByClientCountryRequest) GoString() string {
  return s.String()
}

func (s *GetTrafficAndBandwidthByClientCountryRequest) SetDateFrom(v string) *GetTrafficAndBandwidthByClientCountryRequest {
  s.DateFrom = &v
  return s
}

func (s *GetTrafficAndBandwidthByClientCountryRequest) SetDateTo(v string) *GetTrafficAndBandwidthByClientCountryRequest {
  s.DateTo = &v
  return s
}

func (s *GetTrafficAndBandwidthByClientCountryRequest) SetGranularity(v string) *GetTrafficAndBandwidthByClientCountryRequest {
  s.Granularity = &v
  return s
}

func (s *GetTrafficAndBandwidthByClientCountryRequest) SetDomain(v []*string) *GetTrafficAndBandwidthByClientCountryRequest {
  s.Domain = v
  return s
}

func (s *GetTrafficAndBandwidthByClientCountryRequest) SetCountryCode(v []*string) *GetTrafficAndBandwidthByClientCountryRequest {
  s.CountryCode = v
  return s
}

func (s *GetTrafficAndBandwidthByClientCountryRequest) SetGroupBy(v []*string) *GetTrafficAndBandwidthByClientCountryRequest {
  s.GroupBy = v
  return s
}

type GetTrafficAndBandwidthByClientCountryRequestHeader struct {
}

func (s GetTrafficAndBandwidthByClientCountryRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetTrafficAndBandwidthByClientCountryRequestHeader) GoString() string {
  return s.String()
}

type GetTrafficAndBandwidthByClientCountryPaths struct {
}

func (s GetTrafficAndBandwidthByClientCountryPaths) String() string {
  return tea.Prettify(s)
}

func (s GetTrafficAndBandwidthByClientCountryPaths) GoString() string {
  return s.String()
}

type GetTrafficAndBandwidthByClientCountryParameters struct {
}

func (s GetTrafficAndBandwidthByClientCountryParameters) String() string {
  return tea.Prettify(s)
}

func (s GetTrafficAndBandwidthByClientCountryParameters) GoString() string {
  return s.String()
}

type GetTrafficAndBandwidthByClientCountryResponse struct {
  // {"en":"Result","zh_CN":"结果"}
  Result []*GetTrafficAndBandwidthByClientCountryResponseResult `json:"result,omitempty" xml:"result,omitempty" require:"true" type:"Repeated"`
}

func (s GetTrafficAndBandwidthByClientCountryResponse) String() string {
  return tea.Prettify(s)
}

func (s GetTrafficAndBandwidthByClientCountryResponse) GoString() string {
  return s.String()
}

func (s *GetTrafficAndBandwidthByClientCountryResponse) SetResult(v []*GetTrafficAndBandwidthByClientCountryResponseResult) *GetTrafficAndBandwidthByClientCountryResponse {
  s.Result = v
  return s
}

type GetTrafficAndBandwidthByClientCountryResponseResult struct     {
  // {"en":"Domain","zh_CN":"域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {"en":"","zh_CN":""}
  CountryData []*GetTrafficAndBandwidthByClientCountryResponseResultCountryData `json:"countryData,omitempty" xml:"countryData,omitempty" require:"true" type:"Repeated"`
}

func (s GetTrafficAndBandwidthByClientCountryResponseResult) String() string {
  return tea.Prettify(s)
}

func (s GetTrafficAndBandwidthByClientCountryResponseResult) GoString() string {
  return s.String()
}

func (s *GetTrafficAndBandwidthByClientCountryResponseResult) SetDomain(v string) *GetTrafficAndBandwidthByClientCountryResponseResult {
  s.Domain = &v
  return s
}

func (s *GetTrafficAndBandwidthByClientCountryResponseResult) SetCountryData(v []*GetTrafficAndBandwidthByClientCountryResponseResultCountryData) *GetTrafficAndBandwidthByClientCountryResponseResult {
  s.CountryData = v
  return s
}

type GetTrafficAndBandwidthByClientCountryResponseResultCountryData struct     {
  // {"en":"Country code","zh_CN":"国家地区代号"}
  CountryCode *string `json:"countryCode,omitempty" xml:"countryCode,omitempty" require:"true"`
  // {"en":"Country name","zh_CN":"国家地区名称"}
  CountryName *string `json:"countryName,omitempty" xml:"countryName,omitempty" require:"true"`
  // {"en":"Summary of traffic in national regions: Summary of traffic flow in a single country region during the query period, unit of measure MB, retain two decimals","zh_CN":"国家地区流量汇总:单个国家地区流量在查询时段内的流量汇总值,计量单位MB,保留2位小数"}
  TotalTraffic *string `json:"totalTraffic,omitempty" xml:"totalTraffic,omitempty" require:"true"`
  // {"en":"National regional traffic ratio: the proportion (percentage) of traffic value in a single country region during the query period, retain two decimals","zh_CN":"国家地区流量占比:单个国家地区流量在查询时段内的流量值的占比(百分比),保留2位小数"}
  TrafficPercentage *string `json:"trafficPercentage,omitempty" xml:"trafficPercentage,omitempty" require:"true"`
  // {"en":"","zh_CN":""}
  DataSeries []*GetTrafficAndBandwidthByClientCountryResponseResultCountryDataDataSeries `json:"dataSeries,omitempty" xml:"dataSeries,omitempty" require:"true" type:"Repeated"`
}

func (s GetTrafficAndBandwidthByClientCountryResponseResultCountryData) String() string {
  return tea.Prettify(s)
}

func (s GetTrafficAndBandwidthByClientCountryResponseResultCountryData) GoString() string {
  return s.String()
}

func (s *GetTrafficAndBandwidthByClientCountryResponseResultCountryData) SetCountryCode(v string) *GetTrafficAndBandwidthByClientCountryResponseResultCountryData {
  s.CountryCode = &v
  return s
}

func (s *GetTrafficAndBandwidthByClientCountryResponseResultCountryData) SetCountryName(v string) *GetTrafficAndBandwidthByClientCountryResponseResultCountryData {
  s.CountryName = &v
  return s
}

func (s *GetTrafficAndBandwidthByClientCountryResponseResultCountryData) SetTotalTraffic(v string) *GetTrafficAndBandwidthByClientCountryResponseResultCountryData {
  s.TotalTraffic = &v
  return s
}

func (s *GetTrafficAndBandwidthByClientCountryResponseResultCountryData) SetTrafficPercentage(v string) *GetTrafficAndBandwidthByClientCountryResponseResultCountryData {
  s.TrafficPercentage = &v
  return s
}

func (s *GetTrafficAndBandwidthByClientCountryResponseResultCountryData) SetDataSeries(v []*GetTrafficAndBandwidthByClientCountryResponseResultCountryDataDataSeries) *GetTrafficAndBandwidthByClientCountryResponseResultCountryData {
  s.DataSeries = v
  return s
}

type GetTrafficAndBandwidthByClientCountryResponseResultCountryDataDataSeries struct     {
  // {"en":"1. When the data size of the query is 5m, the format is yyyy-MM-dd HH:mm; each time slice data value represents the data value in the previous time granularity range. The time slice starting at the beginning of the day is yyyy-MM-dd 00:05, and the last time slice is (yyyy-MM-dd+1) 00:00.\n2. Returns the time slice contained in the start time and end time.","zh_CN":"1. 查询的数据粒度为5m时,格式为yyyy-MM-dd  HH:mm;每一个时间片数据值代表的是前一个时间粒度范围内的数据值。一天开始的时间片是yyyy-MM-dd  00:05,最后一个时间片是(yyyy-MM-dd+1) 00:00。\n2. 返回开始时间和结束时间包含的时间片。"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {"en":"Traffic value:\n\nUnit:  MB, retain two decimals","zh_CN":"流量值：\n\n单位MB，保留2位小数"}
  Traffic *string `json:"traffic,omitempty" xml:"traffic,omitempty" require:"true"`
  // {"en":"Bandwidth value. \nUnit: Mbps , retain two decimals","zh_CN":"带宽值,单位Mbps,保留2位小数"}
  Bandwidth *string `json:"bandwidth,omitempty" xml:"bandwidth,omitempty" require:"true"`
}

func (s GetTrafficAndBandwidthByClientCountryResponseResultCountryDataDataSeries) String() string {
  return tea.Prettify(s)
}

func (s GetTrafficAndBandwidthByClientCountryResponseResultCountryDataDataSeries) GoString() string {
  return s.String()
}

func (s *GetTrafficAndBandwidthByClientCountryResponseResultCountryDataDataSeries) SetTimestamp(v string) *GetTrafficAndBandwidthByClientCountryResponseResultCountryDataDataSeries {
  s.Timestamp = &v
  return s
}

func (s *GetTrafficAndBandwidthByClientCountryResponseResultCountryDataDataSeries) SetTraffic(v string) *GetTrafficAndBandwidthByClientCountryResponseResultCountryDataDataSeries {
  s.Traffic = &v
  return s
}

func (s *GetTrafficAndBandwidthByClientCountryResponseResultCountryDataDataSeries) SetBandwidth(v string) *GetTrafficAndBandwidthByClientCountryResponseResultCountryDataDataSeries {
  s.Bandwidth = &v
  return s
}

type GetTrafficAndBandwidthByClientCountryResponseHeader struct {
}

func (s GetTrafficAndBandwidthByClientCountryResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetTrafficAndBandwidthByClientCountryResponseHeader) GoString() string {
  return s.String()
}




type WctQueryRequest struct {
  // {"en":"cust_en_name of sub-client.\nWhen a merged-account wants to  view the information of the subclient,the cust_en_name is required.","zh_CN":"合并账号下的某个客户的英文名，当合并账号要查看子客户的信息时，必须填写子客户的英文名"}
  Cust *string `json:"cust,omitempty" xml:"cust,omitempty"`
  // {"en":"Specifies the query date:\n1)With format yyyy-mm-dd.\n2)If not specified,it means today as default.","zh_CN":"查询的日期，日期格式为yyyy-mm-dd,不选或者为空时默认为当天；"}
  Date *string `json:"date,omitempty" xml:"date,omitempty"`
  // {"en":"1)Must work with 'enddate' and they  specify the query date scope.\n2)With format yyyy-mm-dd.\n3)If there is a 'date' parameter,this parameter will be omitted.","zh_CN":"查询的起始日期 ,日期格式为yyyy-mm-dd；此参数需与enddate参数配合,若存在date参数,则该参数无效"}
  Startdate *string `json:"startdate,omitempty" xml:"startdate,omitempty"`
  // {"en":"1)Must work with 'startdate' and they  specify the query date scope.\n2)With format yyyy-mm-dd\n3)If there is a 'date' parameter,this parameter will be omitted.","zh_CN":"查询的结束日期 ,日期格式为yyyy-mm-dd；此参数需与startdate参数配合,若存在date参数,则该参数无效。"}
  Enddate *string `json:"enddate,omitempty" xml:"enddate,omitempty"`
  // {"en":"GMT time zone, parameter format: GMT+09:00 means east 9th zone, GMT-09:00 means west 9th zone, if not transmitted, the default is local time zone (east 8th zone).","zh_CN":"格林尼治时区，参数格式 GMT+09:00 表示东九区，GMT-09:00 表示西9区，不传则默认为本地时区（东八区）"}
  Timezone *string `json:"timezone,omitempty" xml:"timezone,omitempty"`
  // {"en":"trans;hevc;audio_trans;encap_trans;vframe;file_op;trans_nbhd;hevc_nbhd;video_enhance;drm.","zh_CN":"转码操作类型：trans(H.264视频转码);hevc(H.265视频转码);audio_trans(音频转码);encap_trans(转封装);vframe(截图);file_op(文件处理);trans_nbhd(H.264智控高清视频转码);hevc_nbhd(H.265智控高清视频转码);video_enhance(AI视频增强);drm(DRM加密)"}
  Transcoding *string `json:"transcoding,omitempty" xml:"transcoding,omitempty"`
  // {"en":"sd240;sd480;sd720;hd1080;2k;4k.","zh_CN":"梯度: sd240;sd480;sd720;hd1080;2k;4k"}
  TranscodingType *string `json:"transcodingType,omitempty" xml:"transcodingType,omitempty"`
  // {"en":"Space; for multiple values, please use a semicolon ';' to separate them.","zh_CN":"空间，多个值请用英文分号“;”分割"}
  Space *string `json:"space,omitempty" xml:"space,omitempty"`
  // {"en":"Source provider of the feature. If there are multiple values, please separate them with a semicolon ';'.","zh_CN":"功能来源，提供方。多个值请用英文分号“;”分割"}
  Provider *string `json:"provider,omitempty" xml:"provider,omitempty"`
  // {"en":"1)If there are multiple inputs,use ';' as separator.For example,u can use 'region=cn;apac' to query data of cn and apac region.\n2)If not specified, it means all the regions.","zh_CN":"查询的加速区域的缩写，多个区域请用英文分号';'分隔开，如查询大陆及亚太区域，参数填写为：'region=cn;apac'。不选或者为空时默认为全部区域。"}
  Region *string `json:"region,omitempty" xml:"region,omitempty"`
  // {"en":"The response format:\n1)optional values:xml, json.\n2)'xml' as default.","zh_CN":"返回结果格式,支持格式为xml和json,默认为xml"}
  Dataformat *string `json:"dataformat,omitempty" xml:"dataformat,omitempty"`
}

func (s WctQueryRequest) String() string {
  return tea.Prettify(s)
}

func (s WctQueryRequest) GoString() string {
  return s.String()
}

func (s *WctQueryRequest) SetCust(v string) *WctQueryRequest {
  s.Cust = &v
  return s
}

func (s *WctQueryRequest) SetDate(v string) *WctQueryRequest {
  s.Date = &v
  return s
}

func (s *WctQueryRequest) SetStartdate(v string) *WctQueryRequest {
  s.Startdate = &v
  return s
}

func (s *WctQueryRequest) SetEnddate(v string) *WctQueryRequest {
  s.Enddate = &v
  return s
}

func (s *WctQueryRequest) SetTimezone(v string) *WctQueryRequest {
  s.Timezone = &v
  return s
}

func (s *WctQueryRequest) SetTranscoding(v string) *WctQueryRequest {
  s.Transcoding = &v
  return s
}

func (s *WctQueryRequest) SetTranscodingType(v string) *WctQueryRequest {
  s.TranscodingType = &v
  return s
}

func (s *WctQueryRequest) SetSpace(v string) *WctQueryRequest {
  s.Space = &v
  return s
}

func (s *WctQueryRequest) SetProvider(v string) *WctQueryRequest {
  s.Provider = &v
  return s
}

func (s *WctQueryRequest) SetRegion(v string) *WctQueryRequest {
  s.Region = &v
  return s
}

func (s *WctQueryRequest) SetDataformat(v string) *WctQueryRequest {
  s.Dataformat = &v
  return s
}

type WctQueryRequestHeader struct {
}

func (s WctQueryRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s WctQueryRequestHeader) GoString() string {
  return s.String()
}

type WctQueryPaths struct {
}

func (s WctQueryPaths) String() string {
  return tea.Prettify(s)
}

func (s WctQueryPaths) GoString() string {
  return s.String()
}

type WctQueryParameters struct {
}

func (s WctQueryParameters) String() string {
  return tea.Prettify(s)
}

func (s WctQueryParameters) GoString() string {
  return s.String()
}

type WctQueryResponse struct {
  // {"en":"provider","zh_CN":"结果"}
  Provider *WctQueryResponseProvider `json:"provider,omitempty" xml:"provider,omitempty" require:"true" type:"Struct"`
}

func (s WctQueryResponse) String() string {
  return tea.Prettify(s)
}

func (s WctQueryResponse) GoString() string {
  return s.String()
}

func (s *WctQueryResponse) SetProvider(v *WctQueryResponseProvider) *WctQueryResponse {
  s.Provider = v
  return s
}

type WctQueryResponseProvider struct {
  // {"en":"tenant","zh_CN":"租户"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"type","zh_CN":"接口类型"}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
  // {"en":"data","zh_CN":"转码数据"}
  Date *WctQueryResponseProviderDate `json:"date,omitempty" xml:"date,omitempty" require:"true" type:"Struct"`
}

func (s WctQueryResponseProvider) String() string {
  return tea.Prettify(s)
}

func (s WctQueryResponseProvider) GoString() string {
  return s.String()
}

func (s *WctQueryResponseProvider) SetName(v string) *WctQueryResponseProvider {
  s.Name = &v
  return s
}

func (s *WctQueryResponseProvider) SetType(v string) *WctQueryResponseProvider {
  s.Type = &v
  return s
}

func (s *WctQueryResponseProvider) SetDate(v *WctQueryResponseProviderDate) *WctQueryResponseProvider {
  s.Date = v
  return s
}

type WctQueryResponseProviderDate struct {
  // {"en":"startdate","zh_CN":"开始时间"}
  Startdate *string `json:"startdate,omitempty" xml:"startdate,omitempty" require:"true"`
  // {"en":"enddate","zh_CN":"结束时间"}
  Enddate *string `json:"enddate,omitempty" xml:"enddate,omitempty" require:"true"`
  // {"en":"customerId","zh_CN":"客户id"}
  CustomerId *string `json:"customerId,omitempty" xml:"customerId,omitempty" require:"true"`
  // {"en":"transcoding","zh_CN":"转码类型"}
  Transcoding *string `json:"transcoding,omitempty" xml:"transcoding,omitempty" require:"true"`
  // {"en":"transcodingType","zh_CN":"梯度"}
  TranscodingType *string `json:"transcodingType,omitempty" xml:"transcodingType,omitempty" require:"true"`
  // {"en":"unit","zh_CN":"单位"}
  Unit *string `json:"unit,omitempty" xml:"unit,omitempty" require:"true"`
  // {"en":"total","zh_CN":"总数"}
  Total *string `json:"total,omitempty" xml:"total,omitempty" require:"true"`
  // {"en":"wct","zh_CN":"转码数据"}
  Wct *WctQueryResponseProviderDateWct `json:"wct,omitempty" xml:"wct,omitempty" require:"true" type:"Struct"`
}

func (s WctQueryResponseProviderDate) String() string {
  return tea.Prettify(s)
}

func (s WctQueryResponseProviderDate) GoString() string {
  return s.String()
}

func (s *WctQueryResponseProviderDate) SetStartdate(v string) *WctQueryResponseProviderDate {
  s.Startdate = &v
  return s
}

func (s *WctQueryResponseProviderDate) SetEnddate(v string) *WctQueryResponseProviderDate {
  s.Enddate = &v
  return s
}

func (s *WctQueryResponseProviderDate) SetCustomerId(v string) *WctQueryResponseProviderDate {
  s.CustomerId = &v
  return s
}

func (s *WctQueryResponseProviderDate) SetTranscoding(v string) *WctQueryResponseProviderDate {
  s.Transcoding = &v
  return s
}

func (s *WctQueryResponseProviderDate) SetTranscodingType(v string) *WctQueryResponseProviderDate {
  s.TranscodingType = &v
  return s
}

func (s *WctQueryResponseProviderDate) SetUnit(v string) *WctQueryResponseProviderDate {
  s.Unit = &v
  return s
}

func (s *WctQueryResponseProviderDate) SetTotal(v string) *WctQueryResponseProviderDate {
  s.Total = &v
  return s
}

func (s *WctQueryResponseProviderDate) SetWct(v *WctQueryResponseProviderDateWct) *WctQueryResponseProviderDate {
  s.Wct = v
  return s
}

type WctQueryResponseProviderDateWct struct {
  // {"en":"type","zh_CN":"类型"}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
  // {"en":"total","zh_CN":"总数"}
  Total *string `json:"total,omitempty" xml:"total,omitempty" require:"true"`
  // {"en":"data","zh_CN":"明细数据"}
  Data []*WctQueryResponseProviderDateWctData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s WctQueryResponseProviderDateWct) String() string {
  return tea.Prettify(s)
}

func (s WctQueryResponseProviderDateWct) GoString() string {
  return s.String()
}

func (s *WctQueryResponseProviderDateWct) SetType(v string) *WctQueryResponseProviderDateWct {
  s.Type = &v
  return s
}

func (s *WctQueryResponseProviderDateWct) SetTotal(v string) *WctQueryResponseProviderDateWct {
  s.Total = &v
  return s
}

func (s *WctQueryResponseProviderDateWct) SetData(v []*WctQueryResponseProviderDateWctData) *WctQueryResponseProviderDateWct {
  s.Data = v
  return s
}

type WctQueryResponseProviderDateWctData struct     {
  // {"en":"timestamp","zh_CN":"时间点"}
  Time *string `json:"time,omitempty" xml:"time,omitempty" require:"true"`
  // {"en":"hit count","zh_CN":"明细数"}
  Text *string `json:"text,omitempty" xml:"text,omitempty" require:"true"`
}

func (s WctQueryResponseProviderDateWctData) String() string {
  return tea.Prettify(s)
}

func (s WctQueryResponseProviderDateWctData) GoString() string {
  return s.String()
}

func (s *WctQueryResponseProviderDateWctData) SetTime(v string) *WctQueryResponseProviderDateWctData {
  s.Time = &v
  return s
}

func (s *WctQueryResponseProviderDateWctData) SetText(v string) *WctQueryResponseProviderDateWctData {
  s.Text = &v
  return s
}

type WctQueryResponseHeader struct {
}

func (s WctQueryResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s WctQueryResponseHeader) GoString() string {
  return s.String()
}




type QueryRequestNumberAndPeakBandwidthForMultiDomainRequest struct {
  // {"en":"Start time:
  // 1. The time format is yyyy-MM-ddTHH:mm:ss±HH:mm. Please note: ±HH:mm is the time zone offset, which can be adjusted according to your data needs, for example, +00:00 represents UTC time, +08:00 represents East 8th District, and -05:00 represents West 5th District. 2024-01-15T10:30:45+00:00 means UTC time January 15, 2024 10:30:45 AM;
  // 2. No bigger than the current time.
  // 3. Data in the last 183 days at most can be queried.", "zh_CN":"开始时间：
  // 1.时间格式为 yyyy-MM-ddTHH:mm:ss±HH:mm。请注意：±HH:mm 为时区偏移量，可根据您的数据需要进行调整，例如 +00:00 代表 UTC 时间，+08:00 代表东八区，-05:00 代表西五区。2024-01-15T10:30:45+00:00，表示UTC 时间 2024 年 1 月 15 日上午 10 点 30 分 45 秒；
  // 2.不能大于当前时间
  // 3.最多可获取最近半年（183天）的数据。"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty"`
  // {"en":"End time:
  // 1. The time format is yyyy-MM-ddTHH:mm:ss±HH:mm. Please note: ±HH:mm is the time zone offset, which can be adjusted according to your data needs, for example, +00:00 represents UTC time, +08:00 represents East 8th District, and -05:00 represents West 5th District. 2024-01-15T10:30:45+00:00 means UTC time January 15, 2024 10:30:45 AM
  // 2. End time should be greater than start time. If the end time is greater than current time, current time will be used.
  // 3. If both fields of dataFrom and dateTo are left empty, then data in the last 24 hours will be queried by default; if only one field is filled in and one is left empty, then exception will be occur.
  // 4. Allowable maximum time range for query: 1 day, means the period between dateFrom to dateTo should not exceed 1 day (can be adjusted by contacting technical support up to 31 days).", "zh_CN":"结束时间：
  // 1.时间格式为 yyyy-MM-ddTHH:mm:ss±HH:mm。请注意：±HH:mm 为时区偏移量，可根据您的数据需要进行调整，例如 +00:00 代表 UTC 时间，+08:00 代表东八区，-05:00 代表西五区。2024-01-15T10:30:45+00:00，表示UTC 时间 2024 年 1 月 15 日上午 10 点 30 分 45 秒
  // 2.结束时间需大于开始时间，结束时间如果大于当前时间，取当前时间。
  // 3.dateFrom，dateTo二者都未传，默认查询过去的24小时；如仅有一个未传，抛异常
  // 4.允许查询最大时间间隔：1天，即dateFrom和dateTo相差不能超过1天（可联系技术支持调整，最大不超过31天）。"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty"`
  // {"en":"Domain name:
  // 1. The default upper limit to domains that can be entered is 200 (can be adjusted by contacting technical support);
  // 2. All domains under the account are queried if this input parameter is not specified, but if the number of domains under the account exceeds limits, no query will be done (Error).", "zh_CN":"域名：
  // 1.可传递域名数量上限默认为200个（可联系技术支持调整）；
  // 2.未传递该入参时查询账号下所有域名，但当账号下域名数量超过限制时不可查询（报错）。"}
  Domain []*string `json:"domain,omitempty" xml:"domain,omitempty" type:"Repeated"`
  // {"en":"Group keywords:
  // 1. By default, group data will be displayed;
  // 2. If there are keywords entered, value details shall be displayed by keywords;
  // If domain is specified to groupBy, it means results are returned according to domains.
  // 3. Only domain can be specified.", "zh_CN":"分组关键词：
  // 1.默认聚合展示；
  // 2.传入关键词则代表需要按照关键词对应的值展示明细；
  // 例如groupBy传domain，则代表返回按照domain明细展开。
  // 3.只能传递domain。"}
  GroupBy []*string `json:"groupBy,omitempty" xml:"groupBy,omitempty" type:"Repeated"`
}

func (s QueryRequestNumberAndPeakBandwidthForMultiDomainRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryRequestNumberAndPeakBandwidthForMultiDomainRequest) GoString() string {
  return s.String()
}

func (s *QueryRequestNumberAndPeakBandwidthForMultiDomainRequest) SetDateFrom(v string) *QueryRequestNumberAndPeakBandwidthForMultiDomainRequest {
  s.DateFrom = &v
  return s
}

func (s *QueryRequestNumberAndPeakBandwidthForMultiDomainRequest) SetDateTo(v string) *QueryRequestNumberAndPeakBandwidthForMultiDomainRequest {
  s.DateTo = &v
  return s
}

func (s *QueryRequestNumberAndPeakBandwidthForMultiDomainRequest) SetDomain(v []*string) *QueryRequestNumberAndPeakBandwidthForMultiDomainRequest {
  s.Domain = v
  return s
}

func (s *QueryRequestNumberAndPeakBandwidthForMultiDomainRequest) SetGroupBy(v []*string) *QueryRequestNumberAndPeakBandwidthForMultiDomainRequest {
  s.GroupBy = v
  return s
}

type QueryRequestNumberAndPeakBandwidthForMultiDomainResponse struct {
  // {"en":"", "zh_CN":""}
  Result []*QueryRequestNumberAndPeakBandwidthForMultiDomainResponseResult `json:"result,omitempty" xml:"result,omitempty" require:"true" type:"Repeated"`
}

func (s QueryRequestNumberAndPeakBandwidthForMultiDomainResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryRequestNumberAndPeakBandwidthForMultiDomainResponse) GoString() string {
  return s.String()
}

func (s *QueryRequestNumberAndPeakBandwidthForMultiDomainResponse) SetResult(v []*QueryRequestNumberAndPeakBandwidthForMultiDomainResponseResult) *QueryRequestNumberAndPeakBandwidthForMultiDomainResponse {
  s.Result = v
  return s
}

type QueryRequestNumberAndPeakBandwidthForMultiDomainResponseResult struct     {
  // {"en":"Domain", "zh_CN":"域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {"en":"totalFlow, unit :MB, 2 decimal places reserved, example (74099.92)", "zh_CN":"总流量，单位:MB ，保留2位小数，示例 ( 74099.92 )"}
  TotalFlow *string `json:"totalFlow,omitempty" xml:"totalFlow,omitempty" require:"true"`
  // {"en":"totalRequest", "zh_CN":"总请求数"}
  TotalRequest *string `json:"totalRequest,omitempty" xml:"totalRequest,omitempty" require:"true"`
  // {"en":"peakBandwidth, unit :Mbps, 2 decimal places reserved, example (74099.92)", "zh_CN":"峰值带宽，单位: Mbps，保留2位小数，示例 （931556.21）"}
  PeakBandwidth *string `json:"peakBandwidth,omitempty" xml:"peakBandwidth,omitempty" require:"true"`
  // {"en":"peakTime", "zh_CN":"峰值时间，示例（2019-02-13 18:01）"}
  PeakTime *string `json:"peakTime,omitempty" xml:"peakTime,omitempty" require:"true"`
  // {"en":"Peak request", "zh_CN":"请求数峰值"}
  PeakRequest *string `json:"peakRequest,omitempty" xml:"peakRequest,omitempty" require:"true"`
  // {"en":"Peak time of request", "zh_CN":"请求数峰值时间"}
  PeakRequestTime *string `json:"peakRequestTime,omitempty" xml:"peakRequestTime,omitempty" require:"true"`
  // {"en":"", "zh_CN":""}
  FlowRequestData []*QueryRequestNumberAndPeakBandwidthForMultiDomainResponseResultFlowRequestData `json:"flowRequestData,omitempty" xml:"flowRequestData,omitempty" require:"true" type:"Repeated"`
}

func (s QueryRequestNumberAndPeakBandwidthForMultiDomainResponseResult) String() string {
  return tea.Prettify(s)
}

func (s QueryRequestNumberAndPeakBandwidthForMultiDomainResponseResult) GoString() string {
  return s.String()
}

func (s *QueryRequestNumberAndPeakBandwidthForMultiDomainResponseResult) SetDomain(v string) *QueryRequestNumberAndPeakBandwidthForMultiDomainResponseResult {
  s.Domain = &v
  return s
}

func (s *QueryRequestNumberAndPeakBandwidthForMultiDomainResponseResult) SetTotalFlow(v string) *QueryRequestNumberAndPeakBandwidthForMultiDomainResponseResult {
  s.TotalFlow = &v
  return s
}

func (s *QueryRequestNumberAndPeakBandwidthForMultiDomainResponseResult) SetTotalRequest(v string) *QueryRequestNumberAndPeakBandwidthForMultiDomainResponseResult {
  s.TotalRequest = &v
  return s
}

func (s *QueryRequestNumberAndPeakBandwidthForMultiDomainResponseResult) SetPeakBandwidth(v string) *QueryRequestNumberAndPeakBandwidthForMultiDomainResponseResult {
  s.PeakBandwidth = &v
  return s
}

func (s *QueryRequestNumberAndPeakBandwidthForMultiDomainResponseResult) SetPeakTime(v string) *QueryRequestNumberAndPeakBandwidthForMultiDomainResponseResult {
  s.PeakTime = &v
  return s
}

func (s *QueryRequestNumberAndPeakBandwidthForMultiDomainResponseResult) SetPeakRequest(v string) *QueryRequestNumberAndPeakBandwidthForMultiDomainResponseResult {
  s.PeakRequest = &v
  return s
}

func (s *QueryRequestNumberAndPeakBandwidthForMultiDomainResponseResult) SetPeakRequestTime(v string) *QueryRequestNumberAndPeakBandwidthForMultiDomainResponseResult {
  s.PeakRequestTime = &v
  return s
}

func (s *QueryRequestNumberAndPeakBandwidthForMultiDomainResponseResult) SetFlowRequestData(v []*QueryRequestNumberAndPeakBandwidthForMultiDomainResponseResultFlowRequestData) *QueryRequestNumberAndPeakBandwidthForMultiDomainResponseResult {
  s.FlowRequestData = v
  return s
}

type QueryRequestNumberAndPeakBandwidthForMultiDomainResponseResultFlowRequestData struct     {
  // {"en":"timestamp:
  // 1. When the data granularity of query is 5m, the format is yyyy-mm-dd HH: MM;Each time slice data value represents the data value in the previous time granularity range.The time slice at the beginning of the day is yyyy-mm-dd 00:05, and the last time slice is (yyyy-mm-dd +1) 00:00.
  // 2. Return the time slice of start time and end time.", "zh_CN":"时间片：
  // 1.查询的数据粒度为5m时，格式为yyyy-MM-dd HH:mm；每一个时间片数据值代表的是前一个时间粒度范围内的数据值。一天开始的时间片是yyyy-MM-dd 00:05，最后一个时间片是（yyyy-MM-dd+1） 00:00。
  // 2.返回开始时间和结束时间包含的时间片。"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {"en":"flow", "zh_CN":"流量值，单位MB，保留2位小数；"}
  Flow *string `json:"flow,omitempty" xml:"flow,omitempty" require:"true"`
  // {"en":"Bandwidth, unit: Mbps, 2 decimal places reserved, example (931556.21)", "zh_CN":"带宽值，单位: Mbps，保留2位小数，示例 （931556.21）"}
  Bandwidth *string `json:"bandwidth,omitempty" xml:"bandwidth,omitempty" require:"true"`
  // {"en":"request", "zh_CN":"请求数"}
  Request *string `json:"request,omitempty" xml:"request,omitempty" require:"true"`
}

func (s QueryRequestNumberAndPeakBandwidthForMultiDomainResponseResultFlowRequestData) String() string {
  return tea.Prettify(s)
}

func (s QueryRequestNumberAndPeakBandwidthForMultiDomainResponseResultFlowRequestData) GoString() string {
  return s.String()
}

func (s *QueryRequestNumberAndPeakBandwidthForMultiDomainResponseResultFlowRequestData) SetTimestamp(v string) *QueryRequestNumberAndPeakBandwidthForMultiDomainResponseResultFlowRequestData {
  s.Timestamp = &v
  return s
}

func (s *QueryRequestNumberAndPeakBandwidthForMultiDomainResponseResultFlowRequestData) SetFlow(v string) *QueryRequestNumberAndPeakBandwidthForMultiDomainResponseResultFlowRequestData {
  s.Flow = &v
  return s
}

func (s *QueryRequestNumberAndPeakBandwidthForMultiDomainResponseResultFlowRequestData) SetBandwidth(v string) *QueryRequestNumberAndPeakBandwidthForMultiDomainResponseResultFlowRequestData {
  s.Bandwidth = &v
  return s
}

func (s *QueryRequestNumberAndPeakBandwidthForMultiDomainResponseResultFlowRequestData) SetRequest(v string) *QueryRequestNumberAndPeakBandwidthForMultiDomainResponseResultFlowRequestData {
  s.Request = &v
  return s
}

type QueryRequestNumberAndPeakBandwidthForMultiDomainPaths struct {
}

func (s QueryRequestNumberAndPeakBandwidthForMultiDomainPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryRequestNumberAndPeakBandwidthForMultiDomainPaths) GoString() string {
  return s.String()
}

type QueryRequestNumberAndPeakBandwidthForMultiDomainParameters struct {
}

func (s QueryRequestNumberAndPeakBandwidthForMultiDomainParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryRequestNumberAndPeakBandwidthForMultiDomainParameters) GoString() string {
  return s.String()
}

type QueryRequestNumberAndPeakBandwidthForMultiDomainRequestHeader struct {
}

func (s QueryRequestNumberAndPeakBandwidthForMultiDomainRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryRequestNumberAndPeakBandwidthForMultiDomainRequestHeader) GoString() string {
  return s.String()
}

type QueryRequestNumberAndPeakBandwidthForMultiDomainResponseHeader struct {
}

func (s QueryRequestNumberAndPeakBandwidthForMultiDomainResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryRequestNumberAndPeakBandwidthForMultiDomainResponseHeader) GoString() string {
  return s.String()
}




type QueryStaticDynamicTrafficByProtocolRequest struct {
  // {"en":"Starting time:
  // 
  // 1. The time format is yyyy-MM-ddTHH:mm:ss+08:00, for example, 2016-12-02T10:00:00+08:00 (for Beijing time, December 2, 2016, 10:00:00) );
  // 
  // 2. Cannot be greater than the current time
  // 
  // 3. Get up to the last six months (183 days) of data.", "zh_CN":"开始时间：
  // 1.   时间格式为yyyy-MM-ddTHH:mm:ss+08:00，例如，2016-12-02T10:00:00+08:00（为北京时间2016年12月2日10点0分0秒）；
  // 2. 不能大于当前时间
  // 3.  最多可获取最近半年（183天）的数据。"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty"`
  // {"en":"End Time:
  // 
  // 1. Time format 2016-12-02T10:00:00+08:00
  // 
  // 2. The end time needs to be greater than the start time. If the end time is greater than the current time, the current time is taken.
  // 
  // 3. dateFrom, dateTo are not passed, the default query for the past 1 hour; if only one is not passed, throw an exception
  // 
  // 4. Allow query maximum time interval: 1 day, that is, the difference between dateFrom and dateTo can't exceed 1 day (can contact technical support adjustment).", "zh_CN":"结束时间：
  // 1.   时间格式2016-12-02T10:00:00+08:00
  // 2.  结束时间需大于开始时间，结束时间如果大于当前时间，取当前时间。
  // 3.  dateFrom，dateTo二者都未传，默认查询过去的1小时；如仅有一个未传，抛异常
  // 4.  允许查询最大时间间隔：1天，即dateFrom和dateTo相差不能超过1天（可联系技术支持调整）。"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty"`
  // {"en":"Data granularity:
  // 
  // 1. Support 5m (5 minutes granularity)
  // 
  // 2. Do not pass the default to 5m", "zh_CN":"数据粒度：
  // 1. 支持5m（5分钟粒度）
  // 2.  不传默认为5m"}
  DataInterval *string `json:"dataInterval,omitempty" xml:"dataInterval,omitempty"`
  // {"en":"domain name:
  // 
  // 1. The maximum number of passable domain names is 250 by default (can be contacted by technical support);
  // 
  // 2. Query all the domain names under the account when the entry is not passed, but you cannot query (error) when the number of domain names under the account exceeds the limit.", "zh_CN":"域名：
  // 1. 可传递域名数量上限默认为250个（可联系技术支持调整）；
  // 2.  未传递该入参时查询账号下所有域名，但当账号下域名数量超过限制时不可查询（报错）。"}
  Domain []*string `json:"domain,omitempty" xml:"domain,omitempty" type:"Repeated"`
  // {"en":"agreement type
  // 
  // 1. By default, query all the sum of the protocols;
  // 
  // 2. Support https, http query.", "zh_CN":"协议类型
  // 1. 默认查询所有协议总和；
  // 2. 支持https、http查询。"}
  ProtocolType *string `json:"protocolType,omitempty" xml:"protocolType,omitempty"`
  // {"en":"Group keywords:
  // 
  // 1. Default aggregate display;
  // 
  // 2. The incoming keyword means that the details need to be displayed according to the value corresponding to the keyword; for example, groupBy passes the domain, which means that the return is expanded according to the domain details.
  // 
  // 3. Only domain can be passed.", "zh_CN":"分组关键词：
  // 1. 默认聚合展示；
  // 2. 传入关键词则代表需要按照关键词对应的值展示明细；例如groupBy传domain，则代表返回按照domain明细展开。
  // 3. 只能传递domain。"}
  GroupBy []*string `json:"groupBy,omitempty" xml:"groupBy,omitempty" type:"Repeated"`
  // {"en":"Traffic data unit
  // 
  // 1. MB, keep two decimals
  // 
  // 2. Byte
  // 
  // 3. Do not pass the default to MB", "zh_CN":"流量数据单位
  // 1. MB, 保留两位小数
  // 2. Byte
  // 3. 不传默认为MB"}
  DataUnit *string `json:"dataUnit,omitempty" xml:"dataUnit,omitempty"`
}

func (s QueryStaticDynamicTrafficByProtocolRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryStaticDynamicTrafficByProtocolRequest) GoString() string {
  return s.String()
}

func (s *QueryStaticDynamicTrafficByProtocolRequest) SetDateFrom(v string) *QueryStaticDynamicTrafficByProtocolRequest {
  s.DateFrom = &v
  return s
}

func (s *QueryStaticDynamicTrafficByProtocolRequest) SetDateTo(v string) *QueryStaticDynamicTrafficByProtocolRequest {
  s.DateTo = &v
  return s
}

func (s *QueryStaticDynamicTrafficByProtocolRequest) SetDataInterval(v string) *QueryStaticDynamicTrafficByProtocolRequest {
  s.DataInterval = &v
  return s
}

func (s *QueryStaticDynamicTrafficByProtocolRequest) SetDomain(v []*string) *QueryStaticDynamicTrafficByProtocolRequest {
  s.Domain = v
  return s
}

func (s *QueryStaticDynamicTrafficByProtocolRequest) SetProtocolType(v string) *QueryStaticDynamicTrafficByProtocolRequest {
  s.ProtocolType = &v
  return s
}

func (s *QueryStaticDynamicTrafficByProtocolRequest) SetGroupBy(v []*string) *QueryStaticDynamicTrafficByProtocolRequest {
  s.GroupBy = v
  return s
}

func (s *QueryStaticDynamicTrafficByProtocolRequest) SetDataUnit(v string) *QueryStaticDynamicTrafficByProtocolRequest {
  s.DataUnit = &v
  return s
}

type QueryStaticDynamicTrafficByProtocolResponse struct {
  Result []*QueryStaticDynamicTrafficByProtocolResponseResult `json:"result,omitempty" xml:"result,omitempty" require:"true" type:"Repeated"`
}

func (s QueryStaticDynamicTrafficByProtocolResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryStaticDynamicTrafficByProtocolResponse) GoString() string {
  return s.String()
}

func (s *QueryStaticDynamicTrafficByProtocolResponse) SetResult(v []*QueryStaticDynamicTrafficByProtocolResponseResult) *QueryStaticDynamicTrafficByProtocolResponse {
  s.Result = v
  return s
}

type QueryStaticDynamicTrafficByProtocolResponseResult struct     {
  // {"en":"Domain", "zh_CN":"域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  FlowData []*QueryStaticDynamicTrafficByProtocolResponseResultFlowData `json:"flowData,omitempty" xml:"flowData,omitempty" require:"true" type:"Repeated"`
}

func (s QueryStaticDynamicTrafficByProtocolResponseResult) String() string {
  return tea.Prettify(s)
}

func (s QueryStaticDynamicTrafficByProtocolResponseResult) GoString() string {
  return s.String()
}

func (s *QueryStaticDynamicTrafficByProtocolResponseResult) SetDomain(v string) *QueryStaticDynamicTrafficByProtocolResponseResult {
  s.Domain = &v
  return s
}

func (s *QueryStaticDynamicTrafficByProtocolResponseResult) SetFlowData(v []*QueryStaticDynamicTrafficByProtocolResponseResultFlowData) *QueryStaticDynamicTrafficByProtocolResponseResult {
  s.FlowData = v
  return s
}

type QueryStaticDynamicTrafficByProtocolResponseResultFlowData struct     {
  // {"en":"time
  // 
  //           1. When the data size of the query is 5m, the format is yyyy-MM-dd HH:mm; each time slice data value represents the data value in the previous time granularity range. The time slice starting at the beginning of the day is yyyy-MM-dd 00:05, and the last time slice is (yyyy-MM-dd+1) 00:00.
  // 
  //           2. Returns the time slice contained between the start time and the end time.", "zh_CN":"时间
  //           1.  查询的数据粒度为5m时，格式为yyyy-MM-dd HH:mm；每一个时间片数据值代表的是前一个时间粒度范围内的数据值。一天开始的时间片是yyyy-MM-dd 00:05，最后一个时间片是（yyyy-MM-dd+1）   00:00。
  //           2.  返回开始时间和结束时间之间包含的时间片。"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {"en":"Static inflow traffic value of the edge node, in MB, retaining 2 decimal places;", "zh_CN":"边缘节点的静态流入流量值，单位MB，保留2位小数；"}
  EdgeStaticUp *string `json:"edgeStaticUp,omitempty" xml:"edgeStaticUp,omitempty" require:"true"`
  // {"en":"Static outflow traffic value of the edge node, in MB, retaining 2 decimal places;", "zh_CN":"边缘节点的静态流出流量值，单位MB，保留2位小数；"}
  EdgeStaticDown *string `json:"edgeStaticDown,omitempty" xml:"edgeStaticDown,omitempty" require:"true"`
  // {"en":"Dynamic inflow value of the edge node, in MB, retaining 2 decimal places;", "zh_CN":"边缘节点的动态流入流量值，单位MB，保留2位小数；"}
  EdgeDynamicUp *string `json:"edgeDynamicUp,omitempty" xml:"edgeDynamicUp,omitempty" require:"true"`
  // {"en":"Dynamic outflow traffic value of the edge node, in MB, retaining 2 decimal places;", "zh_CN":"边缘节点的动态流出流量值，单位MB，保留2位小数；"}
  EdgeDynamicDown *string `json:"edgeDynamicDown,omitempty" xml:"edgeDynamicDown,omitempty" require:"true"`
  // {"en":"The static outgoing traffic value of the parent node, in MB, retains 2 decimal places;", "zh_CN":"父节点的静态流出流量值，单位MB，保留2位小数；"}
  ParentStaticDown *string `json:"parentStaticDown,omitempty" xml:"parentStaticDown,omitempty" require:"true"`
  // {"en":"The dynamic outflow value of the parent node, in MB, retains 2 decimal places;", "zh_CN":"父节点的动态流出流量值，单位MB，保留2位小数；"}
  ParentDynamicDown *string `json:"parentDynamicDown,omitempty" xml:"parentDynamicDown,omitempty" require:"true"`
}

func (s QueryStaticDynamicTrafficByProtocolResponseResultFlowData) String() string {
  return tea.Prettify(s)
}

func (s QueryStaticDynamicTrafficByProtocolResponseResultFlowData) GoString() string {
  return s.String()
}

func (s *QueryStaticDynamicTrafficByProtocolResponseResultFlowData) SetTimestamp(v string) *QueryStaticDynamicTrafficByProtocolResponseResultFlowData {
  s.Timestamp = &v
  return s
}

func (s *QueryStaticDynamicTrafficByProtocolResponseResultFlowData) SetEdgeStaticUp(v string) *QueryStaticDynamicTrafficByProtocolResponseResultFlowData {
  s.EdgeStaticUp = &v
  return s
}

func (s *QueryStaticDynamicTrafficByProtocolResponseResultFlowData) SetEdgeStaticDown(v string) *QueryStaticDynamicTrafficByProtocolResponseResultFlowData {
  s.EdgeStaticDown = &v
  return s
}

func (s *QueryStaticDynamicTrafficByProtocolResponseResultFlowData) SetEdgeDynamicUp(v string) *QueryStaticDynamicTrafficByProtocolResponseResultFlowData {
  s.EdgeDynamicUp = &v
  return s
}

func (s *QueryStaticDynamicTrafficByProtocolResponseResultFlowData) SetEdgeDynamicDown(v string) *QueryStaticDynamicTrafficByProtocolResponseResultFlowData {
  s.EdgeDynamicDown = &v
  return s
}

func (s *QueryStaticDynamicTrafficByProtocolResponseResultFlowData) SetParentStaticDown(v string) *QueryStaticDynamicTrafficByProtocolResponseResultFlowData {
  s.ParentStaticDown = &v
  return s
}

func (s *QueryStaticDynamicTrafficByProtocolResponseResultFlowData) SetParentDynamicDown(v string) *QueryStaticDynamicTrafficByProtocolResponseResultFlowData {
  s.ParentDynamicDown = &v
  return s
}

type QueryStaticDynamicTrafficByProtocolPaths struct {
}

func (s QueryStaticDynamicTrafficByProtocolPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryStaticDynamicTrafficByProtocolPaths) GoString() string {
  return s.String()
}

type QueryStaticDynamicTrafficByProtocolParameters struct {
}

func (s QueryStaticDynamicTrafficByProtocolParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryStaticDynamicTrafficByProtocolParameters) GoString() string {
  return s.String()
}

type QueryStaticDynamicTrafficByProtocolRequestHeader struct {
}

func (s QueryStaticDynamicTrafficByProtocolRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryStaticDynamicTrafficByProtocolRequestHeader) GoString() string {
  return s.String()
}

type QueryStaticDynamicTrafficByProtocolResponseHeader struct {
}

func (s QueryStaticDynamicTrafficByProtocolResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryStaticDynamicTrafficByProtocolResponseHeader) GoString() string {
  return s.String()
}




type GetTotalTrafficForAllDomainsRequest struct {
}

func (s GetTotalTrafficForAllDomainsRequest) String() string {
  return tea.Prettify(s)
}

func (s GetTotalTrafficForAllDomainsRequest) GoString() string {
  return s.String()
}

type GetTotalTrafficForAllDomainsRequestHeader struct {
}

func (s GetTotalTrafficForAllDomainsRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetTotalTrafficForAllDomainsRequestHeader) GoString() string {
  return s.String()
}

type GetTotalTrafficForAllDomainsPaths struct {
}

func (s GetTotalTrafficForAllDomainsPaths) String() string {
  return tea.Prettify(s)
}

func (s GetTotalTrafficForAllDomainsPaths) GoString() string {
  return s.String()
}

type GetTotalTrafficForAllDomainsParameters struct {
  // {"en":"Start time 1.The time format is yyyy-MM-ddTHH:mm:ss±HH:mm. Please note: ±HH:mm is the time zone offset, which can be adjusted according to your data needs, for example, +00:00 represents UTC time, +08:00 represents East 8th District, and -05:00 represents West 5th District. 2024-01-15T10:30:45+00:00 means UTC time January 15, 2024 10:30:45 AM; 2.Must be smaller than the current time and dateto; 3.Period between datafrom and dateto cannot be longer than 31 days","zh_CN":"开始时间1.时间格式为 yyyy-MM-ddTHH:mm:ss±HH:mm。请注意：±HH:mm 为时区偏移量，可根据您的数据需要进行调整，例如 +00:00 代表 UTC 时间，+08:00 代表东八区，-05:00 代表西五区。2024-01-15T10:30:45+00:00，表示UTC 时间 2024 年 1 月 15 日上午 10 点 30 分 45 秒;2.必须小于当前时间和dateto;3.dateFrom和dateTo相差不能超过31天;4.只能查询最近2年内数据。"}
  Datefrom *string `json:"datefrom,omitempty" xml:"datefrom,omitempty" require:"true"`
  // {"en":"End time\n1.The time format is yyyy-MM-ddTHH:mm:ss±HH:mm. Please note: ±HH:mm is the time zone offset, which can be adjusted according to your data needs, for example, +00:00 represents UTC time, +08:00 represents East 8th District, and -05:00 represents West 5th District. 2024-01-15T10:30:45+00:00 means UTC time January 15, 2024 10:30:45 AM;\n2.Must be greater than 'datefrom';\n3.if it's greater than the current time, then the current time is assigned as the value","zh_CN":"结束时间\n1.时间格式为 yyyy-MM-ddTHH:mm:ss±HH:mm。请注意：±HH:mm 为时区偏移量，可根据您的数据需要进行调整，例如 +00:00 代表 UTC 时间，+08:00 代表东八区，-05:00 代表西五区。2024-01-15T10:30:45+00:00，表示UTC 时间 2024 年 1 月 15 日上午 10 点 30 分 45 秒;\n2.必须大于datefrom;如果大于当前时间,则重新赋值为当前时间;"}
  Dateto *string `json:"dateto,omitempty" xml:"dateto,omitempty" require:"true"`
  // {"defaultValue":"daily","en":"Data granularity\n1.fiveminutes: five minutes, hourly: one hour, daily: one day;\n2.If not specified, daily is set as the default value;\n3.If fiveminutes is specified as the value, then data is returned in actual configured granularity when there is a specific configuration on data collecting granularity for the custome","zh_CN":"数据粒度\n1.fiveminutes:5分钟,hourly:1小时,daily:1天;\n2.不传递,默认为daily;\n3.传递fiveminutes时,若客户数据采集粒度有特殊配置将按实际配置粒度返回。","exampleValue":"fiveminutes,hourly,daily"}
  Granularity *string `json:"granularity,omitempty" xml:"granularity,omitempty"`
}

func (s GetTotalTrafficForAllDomainsParameters) String() string {
  return tea.Prettify(s)
}

func (s GetTotalTrafficForAllDomainsParameters) GoString() string {
  return s.String()
}

func (s *GetTotalTrafficForAllDomainsParameters) SetDatefrom(v string) *GetTotalTrafficForAllDomainsParameters {
  s.Datefrom = &v
  return s
}

func (s *GetTotalTrafficForAllDomainsParameters) SetDateto(v string) *GetTotalTrafficForAllDomainsParameters {
  s.Dateto = &v
  return s
}

func (s *GetTotalTrafficForAllDomainsParameters) SetGranularity(v string) *GetTotalTrafficForAllDomainsParameters {
  s.Granularity = &v
  return s
}

type GetTotalTrafficForAllDomainsResponse struct {
  // {"en":"Total traffic. Unit: MB, retain two decimals","zh_CN":"总流量,保留2位小数,单位为MB"}
  TotalTraffic *string `json:"totalTraffic,omitempty" xml:"totalTraffic,omitempty" require:"true"`
  // {"en":"Traffic data","zh_CN":"流量数据"}
  DataSeries []*GetTotalTrafficForAllDomainsResponseDataSeries `json:"dataSeries,omitempty" xml:"dataSeries,omitempty" require:"true" type:"Repeated"`
}

func (s GetTotalTrafficForAllDomainsResponse) String() string {
  return tea.Prettify(s)
}

func (s GetTotalTrafficForAllDomainsResponse) GoString() string {
  return s.String()
}

func (s *GetTotalTrafficForAllDomainsResponse) SetTotalTraffic(v string) *GetTotalTrafficForAllDomainsResponse {
  s.TotalTraffic = &v
  return s
}

func (s *GetTotalTrafficForAllDomainsResponse) SetDataSeries(v []*GetTotalTrafficForAllDomainsResponseDataSeries) *GetTotalTrafficForAllDomainsResponse {
  s.DataSeries = v
  return s
}

type GetTotalTrafficForAllDomainsResponseDataSeries struct     {
  // {"en":"Time:\n1.When the data query granularity is fiveminutes, the format is yyyy-MM-dd HH:mm; the data value of every time slice represents the data value within the previous time granularity range. The first time slice of the day is yyyy-MM-dd 00:05 AM, and the last one is yyyy-MM-dd 24:00.\n2.When the data query granularity is hourly, the format is yyyy-MM-dd HH; the data value of every time slice represents the data value within the previous time granularity range. The first time slice of the day is yyyy-MM-dd 00:01, and the last one is yyyy-MM-dd 24.\n3.When the data query granularity is daily, the format is yyyy-MM-dd; the data value of every time slice represents the value of the daily data.\n4.Return the time slice contained in start time and the time slice contained in end time","zh_CN":"时间: \n1.查询的数据粒度为fiveminutes时,格式为yyyy-MM-dd HH:mm;每一个时间片数据值代表的是前一个时间粒度范围内的数据值。一天开始的时间片是yyyy-MM-dd 00:05,最后一个时间片是yyyy-MM-dd 24:00。\n2.查询的数据粒度为hourly时,格式为yyyy-MM-dd HH;每一个时间片数据值代表的是前一个时间粒度范围内的数据值。一天开始的时间片是yyyy-MM-dd 01,最后一个时间片是yyyy-MM-dd 24。\n3.查询的数据粒度为daily时,格式为yyyy-MM-dd;每一个时间片数据值代表的该天内的数据值。\n4.返回开始时间和结束时间包含的时间片。"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {"en":"Traffic.  Unit: MB ,retain two decimals","zh_CN":"流量,保留2位小数,单位为MB"}
  Traffic *string `json:"traffic,omitempty" xml:"traffic,omitempty" require:"true"`
}

func (s GetTotalTrafficForAllDomainsResponseDataSeries) String() string {
  return tea.Prettify(s)
}

func (s GetTotalTrafficForAllDomainsResponseDataSeries) GoString() string {
  return s.String()
}

func (s *GetTotalTrafficForAllDomainsResponseDataSeries) SetTimestamp(v string) *GetTotalTrafficForAllDomainsResponseDataSeries {
  s.Timestamp = &v
  return s
}

func (s *GetTotalTrafficForAllDomainsResponseDataSeries) SetTraffic(v string) *GetTotalTrafficForAllDomainsResponseDataSeries {
  s.Traffic = &v
  return s
}

type GetTotalTrafficForAllDomainsResponseHeader struct {
}

func (s GetTotalTrafficForAllDomainsResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetTotalTrafficForAllDomainsResponseHeader) GoString() string {
  return s.String()
}




type QueryTotalTrafficInByteForAllDomainsUnderDWAByteRequest struct {
  // {"en":"Start time:
  // 1. Time format is yyyy-MM-ddTHH:mm:ss+08:00,
  // 2. No bigger than the current time.
  // 3. Data in the last 2 years at most can be queried.", "zh_CN":"开始时间：
  // 1.时间格式为yyyy-MM-ddTHH:mm:ss+08:00
  // 2.不能大于当前时间
  // 3.最多可获取最近2年内的数据"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty"`
  // {"en":"End time:
  // 1. the time format is yyyy-MM-ddTHH:mm:ss+08:00
  // 2. End time should be greater than start time. If the end time is greater than current time, current time will be used.
  // 3. If both fields of dataFrom and dateTo are left empty, then data in the last 1day will be queried by default; if only one field is filled in and one is left empty, then exception will be occur.
  // 4. Allowable maximum time range for query: 31days, means the period between dateFrom to dateTo should not exceed 31days (can be adjusted by contacting technical support).", "zh_CN":"结束时间：
  // 1.时间格式yyyy-MM-ddTHH:mm:ss+08:00
  // 2.结束时间需大于开始时间，结束时间如果大于当前时间，取当前时间，
  // 3.dateFrom，dateTo二者都未传，默认查询过去的24小时；如仅有一个未传，抛异常；
  // 4.允许查询最大时间间隔31天：，即dateFrom和dateTo相差不能超过31天。（可联系技术支持调整）"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty"`
  // {"en":"Data granularity: 5m(5-minute granularity), 1h(1-hour granularity), 1d(1-day granularity)", "zh_CN":"数据粒度： 5m(5分钟粒度)，1h(1小时粒度), 1d (1天粒度)"}
  DataInterval *string `json:"dataInterval,omitempty" xml:"dataInterval,omitempty"`
}

func (s QueryTotalTrafficInByteForAllDomainsUnderDWAByteRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryTotalTrafficInByteForAllDomainsUnderDWAByteRequest) GoString() string {
  return s.String()
}

func (s *QueryTotalTrafficInByteForAllDomainsUnderDWAByteRequest) SetDateFrom(v string) *QueryTotalTrafficInByteForAllDomainsUnderDWAByteRequest {
  s.DateFrom = &v
  return s
}

func (s *QueryTotalTrafficInByteForAllDomainsUnderDWAByteRequest) SetDateTo(v string) *QueryTotalTrafficInByteForAllDomainsUnderDWAByteRequest {
  s.DateTo = &v
  return s
}

func (s *QueryTotalTrafficInByteForAllDomainsUnderDWAByteRequest) SetDataInterval(v string) *QueryTotalTrafficInByteForAllDomainsUnderDWAByteRequest {
  s.DataInterval = &v
  return s
}

type QueryTotalTrafficInByteForAllDomainsUnderDWAByteResponse struct {
  // {"en":"The result status code of this request", "zh_CN":"请求结果状态码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"The result info of this request", "zh_CN":"请求结果信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  Data []*QueryTotalTrafficInByteForAllDomainsUnderDWAByteResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s QueryTotalTrafficInByteForAllDomainsUnderDWAByteResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryTotalTrafficInByteForAllDomainsUnderDWAByteResponse) GoString() string {
  return s.String()
}

func (s *QueryTotalTrafficInByteForAllDomainsUnderDWAByteResponse) SetCode(v string) *QueryTotalTrafficInByteForAllDomainsUnderDWAByteResponse {
  s.Code = &v
  return s
}

func (s *QueryTotalTrafficInByteForAllDomainsUnderDWAByteResponse) SetMessage(v string) *QueryTotalTrafficInByteForAllDomainsUnderDWAByteResponse {
  s.Message = &v
  return s
}

func (s *QueryTotalTrafficInByteForAllDomainsUnderDWAByteResponse) SetData(v []*QueryTotalTrafficInByteForAllDomainsUnderDWAByteResponseData) *QueryTotalTrafficInByteForAllDomainsUnderDWAByteResponse {
  s.Data = v
  return s
}

type QueryTotalTrafficInByteForAllDomainsUnderDWAByteResponseData struct     {
  // {"en":"flowSummary", "zh_CN":"总流量"}
  FlowSummary *string `json:"flowSummary,omitempty" xml:"flowSummary,omitempty" require:"true"`
  FlowData []*QueryTotalTrafficInByteForAllDomainsUnderDWAByteResponseDataFlowData `json:"flowData,omitempty" xml:"flowData,omitempty" require:"true" type:"Repeated"`
}

func (s QueryTotalTrafficInByteForAllDomainsUnderDWAByteResponseData) String() string {
  return tea.Prettify(s)
}

func (s QueryTotalTrafficInByteForAllDomainsUnderDWAByteResponseData) GoString() string {
  return s.String()
}

func (s *QueryTotalTrafficInByteForAllDomainsUnderDWAByteResponseData) SetFlowSummary(v string) *QueryTotalTrafficInByteForAllDomainsUnderDWAByteResponseData {
  s.FlowSummary = &v
  return s
}

func (s *QueryTotalTrafficInByteForAllDomainsUnderDWAByteResponseData) SetFlowData(v []*QueryTotalTrafficInByteForAllDomainsUnderDWAByteResponseDataFlowData) *QueryTotalTrafficInByteForAllDomainsUnderDWAByteResponseData {
  s.FlowData = v
  return s
}

type QueryTotalTrafficInByteForAllDomainsUnderDWAByteResponseDataFlowData struct     {
  // {"en":"timestamp", "zh_CN":"时间"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {"en":"flow", "zh_CN":"流量"}
  Flow *string `json:"flow,omitempty" xml:"flow,omitempty" require:"true"`
}

func (s QueryTotalTrafficInByteForAllDomainsUnderDWAByteResponseDataFlowData) String() string {
  return tea.Prettify(s)
}

func (s QueryTotalTrafficInByteForAllDomainsUnderDWAByteResponseDataFlowData) GoString() string {
  return s.String()
}

func (s *QueryTotalTrafficInByteForAllDomainsUnderDWAByteResponseDataFlowData) SetTimestamp(v string) *QueryTotalTrafficInByteForAllDomainsUnderDWAByteResponseDataFlowData {
  s.Timestamp = &v
  return s
}

func (s *QueryTotalTrafficInByteForAllDomainsUnderDWAByteResponseDataFlowData) SetFlow(v string) *QueryTotalTrafficInByteForAllDomainsUnderDWAByteResponseDataFlowData {
  s.Flow = &v
  return s
}

type QueryTotalTrafficInByteForAllDomainsUnderDWABytePaths struct {
}

func (s QueryTotalTrafficInByteForAllDomainsUnderDWABytePaths) String() string {
  return tea.Prettify(s)
}

func (s QueryTotalTrafficInByteForAllDomainsUnderDWABytePaths) GoString() string {
  return s.String()
}

type QueryTotalTrafficInByteForAllDomainsUnderDWAByteParameters struct {
}

func (s QueryTotalTrafficInByteForAllDomainsUnderDWAByteParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryTotalTrafficInByteForAllDomainsUnderDWAByteParameters) GoString() string {
  return s.String()
}

type QueryTotalTrafficInByteForAllDomainsUnderDWAByteRequestHeader struct {
}

func (s QueryTotalTrafficInByteForAllDomainsUnderDWAByteRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryTotalTrafficInByteForAllDomainsUnderDWAByteRequestHeader) GoString() string {
  return s.String()
}

type QueryTotalTrafficInByteForAllDomainsUnderDWAByteResponseHeader struct {
}

func (s QueryTotalTrafficInByteForAllDomainsUnderDWAByteResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryTotalTrafficInByteForAllDomainsUnderDWAByteResponseHeader) GoString() string {
  return s.String()
}




type QueryKeywordaquanEdgereportRequest struct {
  // {"en":"cust_en_name of sub-client. When a merged-account wants to view the information of the subclient,the cust_en_name is required.", "zh_CN":"合并账号下的某个客户的英文名，当合并账号要查看子客户的信息时，必须填写子客户的英文名"}
  Cust *string `json:"cust,omitempty" xml:"cust,omitempty"`
  // {"en":"Start time:
  // 1. Time format is yyyy-MM-dd, for example:2016-12-02;
  // 2. No bigger than the current time.", "zh_CN":"Start time:
  // 1. Time format is yyyy-MM-dd, for example:2016-12-02;
  // 2. No bigger than the current time."}
  Startdate *string `json:"startdate,omitempty" xml:"startdate,omitempty"`
  // {"en":"End time:
  // 1. The time format is 2016-12-02
  // 2. End time should be greater than start time. If the end time is greater than current time, current time will be used.
  // 3. Allowable maximum time range for query: 1 day, means the period between dateFrom to dateTo should not exceed 1 day (can be adjusted by contacting technical support).", "zh_CN":"End time:
  // 1. The time format is 2016-12-02
  // 2. End time should be greater than start time. If the end time is greater than current time, current time will be used.
  // 3. Allowable maximum time range for query: 1 day, means the period between dateFrom to dateTo should not exceed 1 day (can be adjusted by contacting technical support)."}
  Enddate *string `json:"enddate,omitempty" xml:"enddate,omitempty"`
  // {"en":"domains that been queried: 
  // 1.If there are multiple inputs,use ';' as separator. 
  // 2.Max num:500", "zh_CN":"domains that been queried: 
  // 1.If there are multiple inputs,use ';' as separator. 
  // 2.Max num:500"}
  Channel *string `json:"channel,omitempty" xml:"channel,omitempty" require:"true"`
  // {"en":"Format : GMT+08:00.  Default GMT+08:00.", "zh_CN":"Format : GMT+08:00.  Default GMT+08:00."}
  Timezone *string `json:"timezone,omitempty" xml:"timezone,omitempty"`
}

func (s QueryKeywordaquanEdgereportRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryKeywordaquanEdgereportRequest) GoString() string {
  return s.String()
}

func (s *QueryKeywordaquanEdgereportRequest) SetCust(v string) *QueryKeywordaquanEdgereportRequest {
  s.Cust = &v
  return s
}

func (s *QueryKeywordaquanEdgereportRequest) SetStartdate(v string) *QueryKeywordaquanEdgereportRequest {
  s.Startdate = &v
  return s
}

func (s *QueryKeywordaquanEdgereportRequest) SetEnddate(v string) *QueryKeywordaquanEdgereportRequest {
  s.Enddate = &v
  return s
}

func (s *QueryKeywordaquanEdgereportRequest) SetChannel(v string) *QueryKeywordaquanEdgereportRequest {
  s.Channel = &v
  return s
}

func (s *QueryKeywordaquanEdgereportRequest) SetTimezone(v string) *QueryKeywordaquanEdgereportRequest {
  s.Timezone = &v
  return s
}

type QueryKeywordaquanEdgereportResponse struct {
  // {'en':'status code', 'zh_CN':'状态码'}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {'en':'message', 'zh_CN':'消息'}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {'en':'bandwidth', 'zh_CN':'统计数据'}
  Data []*QueryKeywordaquanEdgereportResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s QueryKeywordaquanEdgereportResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryKeywordaquanEdgereportResponse) GoString() string {
  return s.String()
}

func (s *QueryKeywordaquanEdgereportResponse) SetCode(v string) *QueryKeywordaquanEdgereportResponse {
  s.Code = &v
  return s
}

func (s *QueryKeywordaquanEdgereportResponse) SetMessage(v string) *QueryKeywordaquanEdgereportResponse {
  s.Message = &v
  return s
}

func (s *QueryKeywordaquanEdgereportResponse) SetData(v []*QueryKeywordaquanEdgereportResponseData) *QueryKeywordaquanEdgereportResponse {
  s.Data = v
  return s
}

type QueryKeywordaquanEdgereportResponseData struct     {
  // {'en':'domain', 'zh_CN':'域名'}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {'en':'contract', 'zh_CN':'合同号'}
  Contract *string `json:"contract,omitempty" xml:"contract,omitempty" require:"true"`
  // {'en':'peakValue', 'zh_CN':'峰值带宽'}
  PeakValue *string `json:"peakValue,omitempty" xml:"peakValue,omitempty" require:"true"`
  // {'en':'peakAvg', 'zh_CN':'峰值平均'}
  PeakAvg *string `json:"peakAvg,omitempty" xml:"peakAvg,omitempty" require:"true"`
  // {'en':'totalFlow', 'zh_CN':'总流量'}
  TotalFlow *string `json:"totalFlow,omitempty" xml:"totalFlow,omitempty" require:"true"`
}

func (s QueryKeywordaquanEdgereportResponseData) String() string {
  return tea.Prettify(s)
}

func (s QueryKeywordaquanEdgereportResponseData) GoString() string {
  return s.String()
}

func (s *QueryKeywordaquanEdgereportResponseData) SetName(v string) *QueryKeywordaquanEdgereportResponseData {
  s.Name = &v
  return s
}

func (s *QueryKeywordaquanEdgereportResponseData) SetContract(v string) *QueryKeywordaquanEdgereportResponseData {
  s.Contract = &v
  return s
}

func (s *QueryKeywordaquanEdgereportResponseData) SetPeakValue(v string) *QueryKeywordaquanEdgereportResponseData {
  s.PeakValue = &v
  return s
}

func (s *QueryKeywordaquanEdgereportResponseData) SetPeakAvg(v string) *QueryKeywordaquanEdgereportResponseData {
  s.PeakAvg = &v
  return s
}

func (s *QueryKeywordaquanEdgereportResponseData) SetTotalFlow(v string) *QueryKeywordaquanEdgereportResponseData {
  s.TotalFlow = &v
  return s
}

type QueryKeywordaquanEdgereportPaths struct {
}

func (s QueryKeywordaquanEdgereportPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryKeywordaquanEdgereportPaths) GoString() string {
  return s.String()
}

type QueryKeywordaquanEdgereportParameters struct {
}

func (s QueryKeywordaquanEdgereportParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryKeywordaquanEdgereportParameters) GoString() string {
  return s.String()
}

type QueryKeywordaquanEdgereportRequestHeader struct {
}

func (s QueryKeywordaquanEdgereportRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryKeywordaquanEdgereportRequestHeader) GoString() string {
  return s.String()
}

type QueryKeywordaquanEdgereportResponseHeader struct {
}

func (s QueryKeywordaquanEdgereportResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryKeywordaquanEdgereportResponseHeader) GoString() string {
  return s.String()
}




type GetTrafficAndRequestsHitRatioByIspProvince1minGranularityRequest struct {
  // {"en":"Start date:\n1.The time format is yyyy-MM-ddTHH:mm:ss±HH:mm. Please note: ±HH:mm is the time zone offset, which can be adjusted according to your data needs, for example, +00:00 represents UTC time, +08:00 represents East 8th District, and -05:00 represents West 5th District. 2024-01-15T10:30:45+00:00 means UTC time January 15, 2024 10:30:45 AM; \n2. Cannot exceed current time\n3. The most recent six-month (183 days) data are available.","zh_CN":"开始时间:\n1.时间格式为 yyyy-MM-ddTHH:mm:ss±HH:mm。请注意：±HH:mm 为时区偏移量，可根据您的数据需要进行调整，例如 +00:00 代表 UTC 时间，+08:00 代表东八区，-05:00 代表西五区。2024-01-15T10:30:45+00:00，表示UTC 时间 2024 年 1 月 15 日上午 10 点 30 分 45 秒; \n2.不能大于当前时间\n3.最多可获取最近半年（183天）的数据。"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty"`
  // {"en":"End time:\n1.The time format is yyyy-MM-ddTHH:mm:ss±HH:mm. Please note: ±HH:mm is the time zone offset, which can be adjusted according to your data needs, for example, +00:00 represents UTC time, +08:00 represents East 8th District, and -05:00 represents West 5th District. 2024-01-15T10:30:45+00:00 means UTC time January 15, 2024 10:30:45 AM; \n2. The end time needs to be greater than the start time. If the end time is greater than the current time, take the current time.\n3. Date from, Date to both, the default query past 24 hours; If there is only one unsent, throw an exception\n4. Maximum allowed query time interval: 24 hours (with technical support adjustments), meaning that the difference between Date from and dateTo cannot exceed 24 hours.","zh_CN":"结束时间:\n1.时间格式为 yyyy-MM-ddTHH:mm:ss±HH:mm。请注意：±HH:mm 为时区偏移量，可根据您的数据需要进行调整，例如 +00:00 代表 UTC 时间，+08:00 代表东八区，-05:00 代表西五区。2024-01-15T10:30:45+00:00，表示UTC 时间 2024 年 1 月 15 日上午 10 点 30 分 45 秒; \n2.结束时间需大于开始时间，结束时间如果大于当前时间，取当前时间。\n3.dateFrom，dateTo二者都未传，默认查询过去的24小时;如仅有一个未传，抛异常\n4.允许查询最大时间间隔:24小时（可联系技术支持调整），即dateFrom和dateTo相差不能超过24小时。"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty"`
  // {"en":"Domain name:\n1. The maximum number of TLDs that can be delivered is 20 by default (contact technical support adjustment);\n2. Automatically filter out illegal domain names (pass illegal domain names, will be filtered out, the query results only return the data of legitimate domain names)\n3. Domain name exceeding limit, misstatement","zh_CN":"域名:\n1.可传递域名数量上限默认为20个（可联系技术支持调整）;\n2.自动过滤掉非法域名（如传递非法域名，会被过滤掉，查询结果只返回合法域名的数据）\n3.域名超过上限，提示错误"}
  Domain []*string `json:"domain,omitempty" xml:"domain,omitempty" type:"Repeated"`
  // {"defaultValue":"1m","en":"Data granularity:\n1. default 1m;\n2. 1m (1 minute), 5m (5 minutes)","zh_CN":"数据粒度:\n1.不传默认1m;\n2.支持1m（1分钟）、5m（5分钟）","exampleValue":"1m,5m"}
  Granularity *string `json:"granularity,omitempty" xml:"granularity,omitempty"`
  // {"dictionary":"belong=Portal-CDN-Service|dict=si-provinceCodeList","en":"Province:\n1.If no province  is specified: Query all provinces and aggregate the returned data according to all provinces\n2.Province specified: Send province code, multiple codes can be sent.","zh_CN":"省份：\n1.未传递province时：查询所有省份，返回的数据按照所有省份聚合。\n2.有传递province时：省份传code，可传多个。"}
  Province []*string `json:"province,omitempty" xml:"province,omitempty" type:"Repeated"`
  // {"dictionary":"belong=Portal-CDN-Service|dict=ISP_CODE_BY_SI_ISP","en":"ISP:\n1. If no ISP is specified: all ISPs will be queried, and the returned data will be aggregated by all ISPs\n2. ISP specified: Send ISP code, multiple codes can be provided","zh_CN":"运营商：\n1.未传递isp时：查询所有isp，返回的数据按照所有运营商聚合\n2.有传递isp时：传递运营商code，可传多个"}
  Isp []*string `json:"isp,omitempty" xml:"isp,omitempty" type:"Repeated"`
  // {"defaultValue":"flow","en":"query dimensionality:\n1. Optional value flow, request\n2. Default flow\n3. Flow: Flow Unit MB, keep two decimal places;\n4. Request: number of Request","zh_CN":"查询维度:\n1.可选值 flow、request\n2.传默认 flow\n3.flow:流量，单位MB，保留两位小数;\n4.request:请求数","exampleValue":"flow,request"}
  QueryBy *string `json:"queryBy,omitempty" xml:"queryBy,omitempty"`
  // {"defaultValue":"domain","en":"Optional: \ndomain, all, If it is empty, it defaults to returning by domain dimension;\nIf all is passed, merge and return according to the query domain name.","zh_CN":"可选项：\ndomain、all, 为空则默认为按domain维度返回;\n若传递all，则按查询域名合并返回","exampleValue":"domain,all"}
  GroupBy *string `json:"groupBy,omitempty" xml:"groupBy,omitempty"`
}

func (s GetTrafficAndRequestsHitRatioByIspProvince1minGranularityRequest) String() string {
  return tea.Prettify(s)
}

func (s GetTrafficAndRequestsHitRatioByIspProvince1minGranularityRequest) GoString() string {
  return s.String()
}

func (s *GetTrafficAndRequestsHitRatioByIspProvince1minGranularityRequest) SetDateFrom(v string) *GetTrafficAndRequestsHitRatioByIspProvince1minGranularityRequest {
  s.DateFrom = &v
  return s
}

func (s *GetTrafficAndRequestsHitRatioByIspProvince1minGranularityRequest) SetDateTo(v string) *GetTrafficAndRequestsHitRatioByIspProvince1minGranularityRequest {
  s.DateTo = &v
  return s
}

func (s *GetTrafficAndRequestsHitRatioByIspProvince1minGranularityRequest) SetDomain(v []*string) *GetTrafficAndRequestsHitRatioByIspProvince1minGranularityRequest {
  s.Domain = v
  return s
}

func (s *GetTrafficAndRequestsHitRatioByIspProvince1minGranularityRequest) SetGranularity(v string) *GetTrafficAndRequestsHitRatioByIspProvince1minGranularityRequest {
  s.Granularity = &v
  return s
}

func (s *GetTrafficAndRequestsHitRatioByIspProvince1minGranularityRequest) SetProvince(v []*string) *GetTrafficAndRequestsHitRatioByIspProvince1minGranularityRequest {
  s.Province = v
  return s
}

func (s *GetTrafficAndRequestsHitRatioByIspProvince1minGranularityRequest) SetIsp(v []*string) *GetTrafficAndRequestsHitRatioByIspProvince1minGranularityRequest {
  s.Isp = v
  return s
}

func (s *GetTrafficAndRequestsHitRatioByIspProvince1minGranularityRequest) SetQueryBy(v string) *GetTrafficAndRequestsHitRatioByIspProvince1minGranularityRequest {
  s.QueryBy = &v
  return s
}

func (s *GetTrafficAndRequestsHitRatioByIspProvince1minGranularityRequest) SetGroupBy(v string) *GetTrafficAndRequestsHitRatioByIspProvince1minGranularityRequest {
  s.GroupBy = &v
  return s
}

type GetTrafficAndRequestsHitRatioByIspProvince1minGranularityRequestHeader struct {
}

func (s GetTrafficAndRequestsHitRatioByIspProvince1minGranularityRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetTrafficAndRequestsHitRatioByIspProvince1minGranularityRequestHeader) GoString() string {
  return s.String()
}

type GetTrafficAndRequestsHitRatioByIspProvince1minGranularityPaths struct {
}

func (s GetTrafficAndRequestsHitRatioByIspProvince1minGranularityPaths) String() string {
  return tea.Prettify(s)
}

func (s GetTrafficAndRequestsHitRatioByIspProvince1minGranularityPaths) GoString() string {
  return s.String()
}

type GetTrafficAndRequestsHitRatioByIspProvince1minGranularityParameters struct {
}

func (s GetTrafficAndRequestsHitRatioByIspProvince1minGranularityParameters) String() string {
  return tea.Prettify(s)
}

func (s GetTrafficAndRequestsHitRatioByIspProvince1minGranularityParameters) GoString() string {
  return s.String()
}

type GetTrafficAndRequestsHitRatioByIspProvince1minGranularityResponse struct {
  // {"en":"Request result status code","zh_CN":"请求结果状态码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Request result information","zh_CN":"请求结果信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Detailed data on the results of the request","zh_CN":"请求结果的详细数据"}
  Data []*GetTrafficAndRequestsHitRatioByIspProvince1minGranularityResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s GetTrafficAndRequestsHitRatioByIspProvince1minGranularityResponse) String() string {
  return tea.Prettify(s)
}

func (s GetTrafficAndRequestsHitRatioByIspProvince1minGranularityResponse) GoString() string {
  return s.String()
}

func (s *GetTrafficAndRequestsHitRatioByIspProvince1minGranularityResponse) SetCode(v string) *GetTrafficAndRequestsHitRatioByIspProvince1minGranularityResponse {
  s.Code = &v
  return s
}

func (s *GetTrafficAndRequestsHitRatioByIspProvince1minGranularityResponse) SetMessage(v string) *GetTrafficAndRequestsHitRatioByIspProvince1minGranularityResponse {
  s.Message = &v
  return s
}

func (s *GetTrafficAndRequestsHitRatioByIspProvince1minGranularityResponse) SetData(v []*GetTrafficAndRequestsHitRatioByIspProvince1minGranularityResponseData) *GetTrafficAndRequestsHitRatioByIspProvince1minGranularityResponse {
  s.Data = v
  return s
}

type GetTrafficAndRequestsHitRatioByIspProvince1minGranularityResponseData struct     {
  // {"en":"Domain","zh_CN":"域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {"en":"","zh_CN":""}
  DataSeries []*GetTrafficAndRequestsHitRatioByIspProvince1minGranularityResponseDataDataSeries `json:"dataSeries,omitempty" xml:"dataSeries,omitempty" require:"true" type:"Repeated"`
}

func (s GetTrafficAndRequestsHitRatioByIspProvince1minGranularityResponseData) String() string {
  return tea.Prettify(s)
}

func (s GetTrafficAndRequestsHitRatioByIspProvince1minGranularityResponseData) GoString() string {
  return s.String()
}

func (s *GetTrafficAndRequestsHitRatioByIspProvince1minGranularityResponseData) SetDomain(v string) *GetTrafficAndRequestsHitRatioByIspProvince1minGranularityResponseData {
  s.Domain = &v
  return s
}

func (s *GetTrafficAndRequestsHitRatioByIspProvince1minGranularityResponseData) SetDataSeries(v []*GetTrafficAndRequestsHitRatioByIspProvince1minGranularityResponseDataDataSeries) *GetTrafficAndRequestsHitRatioByIspProvince1minGranularityResponseData {
  s.DataSeries = v
  return s
}

type GetTrafficAndRequestsHitRatioByIspProvince1minGranularityResponseDataDataSeries struct     {
  // {"en":"Time, in yyyy-MM-dd HH:MM","zh_CN":"时间，格式为yyyy-MM-dd HH:mm"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {"en":"Hit data:\nFlow: Flow, keep two decimal places;\nRequest: number of Request","zh_CN":"命中数据:\n1.flow:流量，保留两位小数;\n2.request:请求数"}
  HitValue *string `json:"hitValue,omitempty" xml:"hitValue,omitempty" require:"true"`
  // {"en":"Hit rate, keep four decimal places","zh_CN":"命中率，保留四位小数"}
  HitRate *string `json:"hitRate,omitempty" xml:"hitRate,omitempty" require:"true"`
}

func (s GetTrafficAndRequestsHitRatioByIspProvince1minGranularityResponseDataDataSeries) String() string {
  return tea.Prettify(s)
}

func (s GetTrafficAndRequestsHitRatioByIspProvince1minGranularityResponseDataDataSeries) GoString() string {
  return s.String()
}

func (s *GetTrafficAndRequestsHitRatioByIspProvince1minGranularityResponseDataDataSeries) SetTimestamp(v string) *GetTrafficAndRequestsHitRatioByIspProvince1minGranularityResponseDataDataSeries {
  s.Timestamp = &v
  return s
}

func (s *GetTrafficAndRequestsHitRatioByIspProvince1minGranularityResponseDataDataSeries) SetHitValue(v string) *GetTrafficAndRequestsHitRatioByIspProvince1minGranularityResponseDataDataSeries {
  s.HitValue = &v
  return s
}

func (s *GetTrafficAndRequestsHitRatioByIspProvince1minGranularityResponseDataDataSeries) SetHitRate(v string) *GetTrafficAndRequestsHitRatioByIspProvince1minGranularityResponseDataDataSeries {
  s.HitRate = &v
  return s
}

type GetTrafficAndRequestsHitRatioByIspProvince1minGranularityResponseHeader struct {
}

func (s GetTrafficAndRequestsHitRatioByIspProvince1minGranularityResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetTrafficAndRequestsHitRatioByIspProvince1minGranularityResponseHeader) GoString() string {
  return s.String()
}




type QueryEdgeByteHitRatioServiceRequest struct {
  // {"en":"Start time:
  // 1. The format is yyyyy-MM-ddTHH: mm: SS + 08:00, for example, 2016-12-02T10:00 + 08:00 (10:00:00 Beijing time on December 2, 2016);
  // 2. Can not exceed the current time;
  // 3. The latest half year (183 days) data can be obtained at most.;", "zh_CN":"开始时间：
  // 1.格式为yyyy-MM-ddTHH:mm:ss+08:00，例如，2016-12-02T10:00:00+08:00（为北京时间2016年12月2日10点0分0秒）；
  // 2.不能大于当前时间；
  // 3.最多可获取最近半年（183天）的数据。"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty"`
  // {"en":"End time:
  // 1. The 1format is yyyy-MM-ddTHH:mm:ss+08:00;
  // 2. The end time is greater than the start time.
  // 3. If the end time is greater than the current time, the current time is taken.
  // 4. DateFrom and dateTo are not uploaded, defaulting to query the past 24 hours; if only one is not uploaded, throw an exception;
  // 5. Maximum query interval allowed: 7 days, that is, the difference between dateFrom and dateTo can not exceed 7 days ", "zh_CN":"结束时间：
  // 1.格式为yyyy-MM-ddTHH:mm:ss+08:00；
  // 2.结束时间需大于开始时间；
  // 3.结束时间如果大于当前时间，取当前时间；
  // 4.dateFrom，dateTo二者都未传，默认查询过去的24小时；如仅有一个未传，抛异常；
  // 5.允许查询最大间隔：7天。"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty"`
  // {"en":"Domains:
  // 1.Domain is not uploaded: Query all domain names of the account (More than 200 domains will error , you can contact technical support for adjustment);
  // 2.Domain is uploaded: Up to 200 domains are supported (you can contact technical support for adjustment).", "zh_CN":"域名：
  // 1.未传递domain时：查询账号下所有全部域名(域名超过200个则报错，可联系技术支持调整)；
  // 2.有传递domain时：域名最多支持传200个（可联系技术支持调整）。"}
  Domain []*string `json:"domain,omitempty" xml:"domain,omitempty" type:"Repeated"`
}

func (s QueryEdgeByteHitRatioServiceRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryEdgeByteHitRatioServiceRequest) GoString() string {
  return s.String()
}

func (s *QueryEdgeByteHitRatioServiceRequest) SetDateFrom(v string) *QueryEdgeByteHitRatioServiceRequest {
  s.DateFrom = &v
  return s
}

func (s *QueryEdgeByteHitRatioServiceRequest) SetDateTo(v string) *QueryEdgeByteHitRatioServiceRequest {
  s.DateTo = &v
  return s
}

func (s *QueryEdgeByteHitRatioServiceRequest) SetDomain(v []*string) *QueryEdgeByteHitRatioServiceRequest {
  s.Domain = v
  return s
}

type QueryEdgeByteHitRatioServiceResponse struct {
  Result []*QueryEdgeByteHitRatioServiceResponseResult `json:"result,omitempty" xml:"result,omitempty" require:"true" type:"Repeated"`
}

func (s QueryEdgeByteHitRatioServiceResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryEdgeByteHitRatioServiceResponse) GoString() string {
  return s.String()
}

func (s *QueryEdgeByteHitRatioServiceResponse) SetResult(v []*QueryEdgeByteHitRatioServiceResponseResult) *QueryEdgeByteHitRatioServiceResponse {
  s.Result = v
  return s
}

type QueryEdgeByteHitRatioServiceResponseResult struct     {
  // {"en":"Actually processed time.  yyyy-MM-dd HH:mm format", "zh_CN":"实际查询时间，格式 yyyy-MM-dd HH:mm"}
  RealDate *string `json:"realDate,omitempty" xml:"realDate,omitempty" require:"true"`
  // {"en":"Average of total edge node hit ratio", "zh_CN":"总边缘节点命中率的平均值,2位小数"}
  TotalAvg *string `json:"totalAvg,omitempty" xml:"totalAvg,omitempty" require:"true"`
  // {"en":"Detailed data on the results of the request", "zh_CN":"请求结果的详细数据"}
  HitRatioDatas []*QueryEdgeByteHitRatioServiceResponseResultHitRatioDatas `json:"hitRatioDatas,omitempty" xml:"hitRatioDatas,omitempty" require:"true" type:"Repeated"`
}

func (s QueryEdgeByteHitRatioServiceResponseResult) String() string {
  return tea.Prettify(s)
}

func (s QueryEdgeByteHitRatioServiceResponseResult) GoString() string {
  return s.String()
}

func (s *QueryEdgeByteHitRatioServiceResponseResult) SetRealDate(v string) *QueryEdgeByteHitRatioServiceResponseResult {
  s.RealDate = &v
  return s
}

func (s *QueryEdgeByteHitRatioServiceResponseResult) SetTotalAvg(v string) *QueryEdgeByteHitRatioServiceResponseResult {
  s.TotalAvg = &v
  return s
}

func (s *QueryEdgeByteHitRatioServiceResponseResult) SetHitRatioDatas(v []*QueryEdgeByteHitRatioServiceResponseResultHitRatioDatas) *QueryEdgeByteHitRatioServiceResponseResult {
  s.HitRatioDatas = v
  return s
}

type QueryEdgeByteHitRatioServiceResponseResultHitRatioDatas struct     {
  // {"en":"timestamp", "zh_CN":"时间，格式为yyyy-MM-dd HH:mm；每一个时间片数据值代表的是前一个时间粒度范围内的数据值。一天开始的时间片是yyyy-MM-dd 00:01，最后一个时间片是（yyyy-MM-dd+1） 00:00；"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {"en":"edge node hit ratio,keep 4 decimal places", "zh_CN":"边缘节点缓存字节命中率，保留4位小数"}
  EdgeHitRatio *string `json:"edgeHitRatio,omitempty" xml:"edgeHitRatio,omitempty" require:"true"`
}

func (s QueryEdgeByteHitRatioServiceResponseResultHitRatioDatas) String() string {
  return tea.Prettify(s)
}

func (s QueryEdgeByteHitRatioServiceResponseResultHitRatioDatas) GoString() string {
  return s.String()
}

func (s *QueryEdgeByteHitRatioServiceResponseResultHitRatioDatas) SetTimestamp(v string) *QueryEdgeByteHitRatioServiceResponseResultHitRatioDatas {
  s.Timestamp = &v
  return s
}

func (s *QueryEdgeByteHitRatioServiceResponseResultHitRatioDatas) SetEdgeHitRatio(v string) *QueryEdgeByteHitRatioServiceResponseResultHitRatioDatas {
  s.EdgeHitRatio = &v
  return s
}

type QueryEdgeByteHitRatioServicePaths struct {
}

func (s QueryEdgeByteHitRatioServicePaths) String() string {
  return tea.Prettify(s)
}

func (s QueryEdgeByteHitRatioServicePaths) GoString() string {
  return s.String()
}

type QueryEdgeByteHitRatioServiceParameters struct {
}

func (s QueryEdgeByteHitRatioServiceParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryEdgeByteHitRatioServiceParameters) GoString() string {
  return s.String()
}

type QueryEdgeByteHitRatioServiceRequestHeader struct {
}

func (s QueryEdgeByteHitRatioServiceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryEdgeByteHitRatioServiceRequestHeader) GoString() string {
  return s.String()
}

type QueryEdgeByteHitRatioServiceResponseHeader struct {
}

func (s QueryEdgeByteHitRatioServiceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryEdgeByteHitRatioServiceResponseHeader) GoString() string {
  return s.String()
}




type QueryTotalTrafficForAllDomainsUnderDWARequest struct {
}

func (s QueryTotalTrafficForAllDomainsUnderDWARequest) String() string {
  return tea.Prettify(s)
}

func (s QueryTotalTrafficForAllDomainsUnderDWARequest) GoString() string {
  return s.String()
}

type QueryTotalTrafficForAllDomainsUnderDWAResponse struct {
  // {"en":"Total traffic. Keep two digits of decimals. Unit: MB", "zh_CN":"总流量,保留2位小数,单位为MB"}
  FlowSummary *string `json:"flow-summary,omitempty" xml:"flow-summary,omitempty" require:"true"`
  // {"en":"flowData", "zh_CN":"流量数据"}
  FlowData []*QueryTotalTrafficForAllDomainsUnderDWAResponseFlowData `json:"flow-data,omitempty" xml:"flow-data,omitempty" require:"true" type:"Repeated"`
}

func (s QueryTotalTrafficForAllDomainsUnderDWAResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryTotalTrafficForAllDomainsUnderDWAResponse) GoString() string {
  return s.String()
}

func (s *QueryTotalTrafficForAllDomainsUnderDWAResponse) SetFlowSummary(v string) *QueryTotalTrafficForAllDomainsUnderDWAResponse {
  s.FlowSummary = &v
  return s
}

func (s *QueryTotalTrafficForAllDomainsUnderDWAResponse) SetFlowData(v []*QueryTotalTrafficForAllDomainsUnderDWAResponseFlowData) *QueryTotalTrafficForAllDomainsUnderDWAResponse {
  s.FlowData = v
  return s
}

type QueryTotalTrafficForAllDomainsUnderDWAResponseFlowData struct     {
  // {"en":"Date:
  // 1.When the data query granularity is fiveminutes, the format is yyyy-MM-dd HH:mm; the data value of every time slice represents the data value within the previous time granularity range. The first time slice of the day is yyyy-MM-dd 00:05 AM, and the last one is yyyy-MM-dd 24:00.
  // 2.When the data query granularity is hourly, the format is yyyy-MM-dd HH; the data value of every time slice represents the data value within the previous time granularity range. The first time slice of the day is yyyy-MM-dd 00:01, and the last one is yyyy-MM-dd 24.
  // 3.When the data query granularity is daily, the format is yyyy-MM-dd; the data value of every time slice represents the value of the daily data.
  // 4.Return the time slice contained in start time and the time slice contained in end time", "zh_CN":"时间
  // 1.查询的数据粒度为fiveminutes时,格式为yyyy-MM-dd HH:mm;每一个时间片数据值代表的是前一个时间粒度范围内的数据值。一天开始的时间片是yyyy-MM-dd 00:05,最后一个时间片是yyyy-MM-dd 24:00。
  // 2.查询的数据粒度为hourly时,格式为yyyy-MM-dd HH;每一个时间片数据值代表的是前一个时间粒度范围内的数据值。一天开始的时间片是yyyy-MM-dd 01,最后一个时间片是yyyy-MM-dd 24。
  // 3.查询的数据粒度为daily时,格式为yyyy-MM-dd;每一个时间片数据值代表的该天内的数据值。
  // 4.返回开始时间和结束时间包含的时间片。"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {"en":"Traffic. Keep two digits of decimals. Unit: MB", "zh_CN":"流量,保留2位小数,单位为MB"}
  Flow *string `json:"flow,omitempty" xml:"flow,omitempty" require:"true"`
}

func (s QueryTotalTrafficForAllDomainsUnderDWAResponseFlowData) String() string {
  return tea.Prettify(s)
}

func (s QueryTotalTrafficForAllDomainsUnderDWAResponseFlowData) GoString() string {
  return s.String()
}

func (s *QueryTotalTrafficForAllDomainsUnderDWAResponseFlowData) SetTimestamp(v string) *QueryTotalTrafficForAllDomainsUnderDWAResponseFlowData {
  s.Timestamp = &v
  return s
}

func (s *QueryTotalTrafficForAllDomainsUnderDWAResponseFlowData) SetFlow(v string) *QueryTotalTrafficForAllDomainsUnderDWAResponseFlowData {
  s.Flow = &v
  return s
}

type QueryTotalTrafficForAllDomainsUnderDWAPaths struct {
}

func (s QueryTotalTrafficForAllDomainsUnderDWAPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryTotalTrafficForAllDomainsUnderDWAPaths) GoString() string {
  return s.String()
}

type QueryTotalTrafficForAllDomainsUnderDWAParameters struct {
  // {"en":"Start time:
  // 1.The format is yyyy-MM-ddTHH:mm:ss+08:00;
  // 2.Must be smaller than the current time and 'dateto';
  // 3.Period between 'datafrom' and 'dateto' cannot be longer than 31 days", "zh_CN":"开始时间:
  // 1.格式为yyyy-MM-ddTHH:mm:ss+08:00;
  // 2.必须小于当前时间和dateto;
  // 3.dateFrom和dateTo相差不能超过31天(可联系技术支持调整);4.只能查询最近2年内数据。"}
  Datefrom *string `json:"datefrom,omitempty" xml:"datefrom,omitempty" require:"true"`
  // {"en":"End time:
  // 1.The format is yyyy-MM-ddTHH:mm:ss+08:00;
  // 2.Must be greater than 'datefrom'; 
  // 3.if it's greater than the current time, then the current time is assigned as the value", "zh_CN":"结束时间:
  // 1.格式为yyyy-MM-ddTHH:mm:ss+08:00;
  // 2.必须大于datefrom;如果大于当前时间,则重新赋值为当前时间;"}
  Dateto *string `json:"dateto,omitempty" xml:"dateto,omitempty" require:"true"`
  // {"en":"Data granularity:
  // 1.fiveminutes: five minutes, hourly: one hour, daily: one day;
  // 2.If not specified, daily is set as the default value;
  // 3.If fiveminutes is specified as the value, then data is returned in actual configured granularity when there is a specific configuration on data collecting granularity for the custome", "zh_CN":"数据粒度:
  // 1.fiveminutes:5分钟,hourly:1小时,daily:1天;
  // 2.不传递,默认为daily;
  // 3.传递fiveminutes时,若客户数据采集粒度有特殊配置将按实际配置粒度返回。"}
  Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s QueryTotalTrafficForAllDomainsUnderDWAParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryTotalTrafficForAllDomainsUnderDWAParameters) GoString() string {
  return s.String()
}

func (s *QueryTotalTrafficForAllDomainsUnderDWAParameters) SetDatefrom(v string) *QueryTotalTrafficForAllDomainsUnderDWAParameters {
  s.Datefrom = &v
  return s
}

func (s *QueryTotalTrafficForAllDomainsUnderDWAParameters) SetDateto(v string) *QueryTotalTrafficForAllDomainsUnderDWAParameters {
  s.Dateto = &v
  return s
}

func (s *QueryTotalTrafficForAllDomainsUnderDWAParameters) SetType(v string) *QueryTotalTrafficForAllDomainsUnderDWAParameters {
  s.Type = &v
  return s
}

type QueryTotalTrafficForAllDomainsUnderDWARequestHeader struct {
}

func (s QueryTotalTrafficForAllDomainsUnderDWARequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryTotalTrafficForAllDomainsUnderDWARequestHeader) GoString() string {
  return s.String()
}

type QueryTotalTrafficForAllDomainsUnderDWAResponseHeader struct {
}

func (s QueryTotalTrafficForAllDomainsUnderDWAResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryTotalTrafficForAllDomainsUnderDWAResponseHeader) GoString() string {
  return s.String()
}




type QueryDirectoryBandwidthTrafficUnderLivestreamDomainRequest struct {
  // {"en":"Start time:
  // 1.The format is yyyy-MM-ddTHH:mm:ss+08:00;
  // 2.Must be smaller than the current time and dateTo;
  // 3.Period between dataFrom and dateTo cannot be longer than 7 days(technical support can be contacted to adjust);4.You can only query data for the last 2 years.", "zh_CN":"开始时间
  // 1.格式为yyyy-MM-ddTHH:mm:ss+08:00;
  // 2.必须小于当前时间和dateTo;
  // 3.dateFrom和dateTo相差不能超过7天(可联系技术支持调整);4.只能查询最近2年内数据。"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty" require:"true"`
  // {"en":"End time:
  // 1.The format is yyyy-MM-ddTHH:mm:ss+08:00;
  // 2.Must be greater than dateFrom;
  // 3.If it's greater than the current time, then the current time is assigned as the value;", "zh_CN":"结束时间
  // 1.格式为yyyy-MM-ddTHH:mm:ss+08:00;
  // 2.必须大于dateFrom;
  // 3.如果大于当前时间,则重新赋值为当前时间;"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty" require:"true"`
  DomainDir []*QueryDirectoryBandwidthTrafficUnderLivestreamDomainRequestDomainDir `json:"domainDir,omitempty" xml:"domainDir,omitempty" require:"true" type:"Repeated"`
  // {"en":"Query protocol:
  // 1.Means query statistics with http protocol or with https protocol;
  // 2.Values can be selected: http or https;
  // 3.Empty means both http and https;", "zh_CN":"查询协议:
  // 1.代表统计http协议或https协议;
  // 2.可选值:http 或 https;
  // 3.不传默认代表不区分http和https;"}
  ProtocolType *string `json:"protocolType,omitempty" xml:"protocolType,omitempty"`
  // {"en":"Query granularity:
  // 1.Empty value means 5 minutes of granularity by default;
  // 2.Values can be selected: 5m or 1m;
  // 3.Option of 1m can be configured in flowDirDataIntervalConfig depending on differrent accounts;", "zh_CN":"查询粒度:
  // 1.不传默认代表5分钟粒度;
  // 2.可选值:5m 或 1m;
  // 3.1m选项根据账号在数据字典flowDirDataIntervalConfig中可配;"}
  DataInterval *string `json:"dataInterval,omitempty" xml:"dataInterval,omitempty"`
  // {"en":"Query type
  // 
  // 1.Values can be selected: flow or bandwidth;
  // 
  // 2.bandwidth means to get bandwidth value, flow means to get the flow vaule. By default, data in all regions is queried;", "zh_CN":"查询类型
  // 1.可选值:flow 或bandwidth;
  // 2.bandwidth代表获取带宽值,flow代表获取流量值,默认查询所有区域的数据;"}
  DataType *string `json:"dataType,omitempty" xml:"dataType,omitempty" require:"true"`
}

func (s QueryDirectoryBandwidthTrafficUnderLivestreamDomainRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryDirectoryBandwidthTrafficUnderLivestreamDomainRequest) GoString() string {
  return s.String()
}

func (s *QueryDirectoryBandwidthTrafficUnderLivestreamDomainRequest) SetDateFrom(v string) *QueryDirectoryBandwidthTrafficUnderLivestreamDomainRequest {
  s.DateFrom = &v
  return s
}

func (s *QueryDirectoryBandwidthTrafficUnderLivestreamDomainRequest) SetDateTo(v string) *QueryDirectoryBandwidthTrafficUnderLivestreamDomainRequest {
  s.DateTo = &v
  return s
}

func (s *QueryDirectoryBandwidthTrafficUnderLivestreamDomainRequest) SetDomainDir(v []*QueryDirectoryBandwidthTrafficUnderLivestreamDomainRequestDomainDir) *QueryDirectoryBandwidthTrafficUnderLivestreamDomainRequest {
  s.DomainDir = v
  return s
}

func (s *QueryDirectoryBandwidthTrafficUnderLivestreamDomainRequest) SetProtocolType(v string) *QueryDirectoryBandwidthTrafficUnderLivestreamDomainRequest {
  s.ProtocolType = &v
  return s
}

func (s *QueryDirectoryBandwidthTrafficUnderLivestreamDomainRequest) SetDataInterval(v string) *QueryDirectoryBandwidthTrafficUnderLivestreamDomainRequest {
  s.DataInterval = &v
  return s
}

func (s *QueryDirectoryBandwidthTrafficUnderLivestreamDomainRequest) SetDataType(v string) *QueryDirectoryBandwidthTrafficUnderLivestreamDomainRequest {
  s.DataType = &v
  return s
}

type QueryDirectoryBandwidthTrafficUnderLivestreamDomainRequestDomainDir struct     {
  // {"en":"Domain:
  // 1.Need to meet the regular expression rules that are used to validate domains;
  // 2.Domain number limits can be adjusted depending on different accounts. The default value is 20;", "zh_CN":"域名:
  // 1.需要满足域名的正则校验;
  // 2.域名个数限制根据账号可调,默认为20个;"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {"en":"Table of contents:
  // 1.Empty value means to query all directories under the domain
  // 2.Directory limits can be adjusted depending on different accounts. The default value is 200(this limit applies to the empty value);
  // 3.Invalid directories are not returned", "zh_CN":"目录:
  // 1.目录个数限制根据账号可调,默认为200个;
  // 2.不传代表查询该域名下的所有目录,同时接受目录个数限制;
  // 3.无效的目录不返回"}
  Dir []*string `json:"dir,omitempty" xml:"dir,omitempty" type:"Repeated"`
}

func (s QueryDirectoryBandwidthTrafficUnderLivestreamDomainRequestDomainDir) String() string {
  return tea.Prettify(s)
}

func (s QueryDirectoryBandwidthTrafficUnderLivestreamDomainRequestDomainDir) GoString() string {
  return s.String()
}

func (s *QueryDirectoryBandwidthTrafficUnderLivestreamDomainRequestDomainDir) SetDomain(v string) *QueryDirectoryBandwidthTrafficUnderLivestreamDomainRequestDomainDir {
  s.Domain = &v
  return s
}

func (s *QueryDirectoryBandwidthTrafficUnderLivestreamDomainRequestDomainDir) SetDir(v []*string) *QueryDirectoryBandwidthTrafficUnderLivestreamDomainRequestDomainDir {
  s.Dir = v
  return s
}

type QueryDirectoryBandwidthTrafficUnderLivestreamDomainResponse struct {
  // {"en":"result", "zh_CN":"结果"}
  Result []*QueryDirectoryBandwidthTrafficUnderLivestreamDomainResponseResult `json:"result,omitempty" xml:"result,omitempty" require:"true" type:"Repeated"`
}

func (s QueryDirectoryBandwidthTrafficUnderLivestreamDomainResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryDirectoryBandwidthTrafficUnderLivestreamDomainResponse) GoString() string {
  return s.String()
}

func (s *QueryDirectoryBandwidthTrafficUnderLivestreamDomainResponse) SetResult(v []*QueryDirectoryBandwidthTrafficUnderLivestreamDomainResponseResult) *QueryDirectoryBandwidthTrafficUnderLivestreamDomainResponse {
  s.Result = v
  return s
}

type QueryDirectoryBandwidthTrafficUnderLivestreamDomainResponseResult struct     {
  // {"en":"Domain", "zh_CN":"域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  Details []*QueryDirectoryBandwidthTrafficUnderLivestreamDomainResponseResultDetails `json:"details,omitempty" xml:"details,omitempty" require:"true" type:"Repeated"`
}

func (s QueryDirectoryBandwidthTrafficUnderLivestreamDomainResponseResult) String() string {
  return tea.Prettify(s)
}

func (s QueryDirectoryBandwidthTrafficUnderLivestreamDomainResponseResult) GoString() string {
  return s.String()
}

func (s *QueryDirectoryBandwidthTrafficUnderLivestreamDomainResponseResult) SetDomain(v string) *QueryDirectoryBandwidthTrafficUnderLivestreamDomainResponseResult {
  s.Domain = &v
  return s
}

func (s *QueryDirectoryBandwidthTrafficUnderLivestreamDomainResponseResult) SetDetails(v []*QueryDirectoryBandwidthTrafficUnderLivestreamDomainResponseResultDetails) *QueryDirectoryBandwidthTrafficUnderLivestreamDomainResponseResult {
  s.Details = v
  return s
}

type QueryDirectoryBandwidthTrafficUnderLivestreamDomainResponseResultDetails struct     {
  // {"en":"Directory", "zh_CN":"目录"}
  Dir *string `json:"dir,omitempty" xml:"dir,omitempty" require:"true"`
  // {"en":"Return when the dataType is flow; Total flow of every stream. Unit is MB and 2 digits of decimals allowed;", "zh_CN":"当dataType为flow时返回;每路流的总流量单位MB,保留2位小数;"}
  TotalFlow *string `json:"totalFlow,omitempty" xml:"totalFlow,omitempty" require:"true"`
  // {"en":"Return when the dataType is bandwidth;
  // 
  // Bandwidth peak value of every stream within specified time. Unit is Mbps, two decimals digits;", "zh_CN":"当dataType为bandwidth时返回;
  // 每路流在该时间段内的带宽峰值单位Mbps,保留2位小数;"}
  BandwidthPeakValue *string `json:"bandwidthPeakValue,omitempty" xml:"bandwidthPeakValue,omitempty" require:"true"`
  Details []*QueryDirectoryBandwidthTrafficUnderLivestreamDomainResponseResultDetailsDetails `json:"details,omitempty" xml:"details,omitempty" require:"true" type:"Repeated"`
}

func (s QueryDirectoryBandwidthTrafficUnderLivestreamDomainResponseResultDetails) String() string {
  return tea.Prettify(s)
}

func (s QueryDirectoryBandwidthTrafficUnderLivestreamDomainResponseResultDetails) GoString() string {
  return s.String()
}

func (s *QueryDirectoryBandwidthTrafficUnderLivestreamDomainResponseResultDetails) SetDir(v string) *QueryDirectoryBandwidthTrafficUnderLivestreamDomainResponseResultDetails {
  s.Dir = &v
  return s
}

func (s *QueryDirectoryBandwidthTrafficUnderLivestreamDomainResponseResultDetails) SetTotalFlow(v string) *QueryDirectoryBandwidthTrafficUnderLivestreamDomainResponseResultDetails {
  s.TotalFlow = &v
  return s
}

func (s *QueryDirectoryBandwidthTrafficUnderLivestreamDomainResponseResultDetails) SetBandwidthPeakValue(v string) *QueryDirectoryBandwidthTrafficUnderLivestreamDomainResponseResultDetails {
  s.BandwidthPeakValue = &v
  return s
}

func (s *QueryDirectoryBandwidthTrafficUnderLivestreamDomainResponseResultDetails) SetDetails(v []*QueryDirectoryBandwidthTrafficUnderLivestreamDomainResponseResultDetailsDetails) *QueryDirectoryBandwidthTrafficUnderLivestreamDomainResponseResultDetails {
  s.Details = v
  return s
}

type QueryDirectoryBandwidthTrafficUnderLivestreamDomainResponseResultDetailsDetails struct     {
  // {"en":"Date
  // 1.When the data query granularity is 1m, the format is yyyy-MM-dd HH:mm; the data value of every time slice represents the data value within the previous time granularity range. The first time slice of the day is yyyy-MM-dd 00:01, and the last one is (yyyy-MM-dd+1) 00:00;
  // 2.When the data query granularity is 5m, the format is yyyy-MM-dd HH:mm; the data value of every time slice represents the data value within the previous time granularity range. The first time slice of the day is yyyy-MM-dd 12:05 AM, and the last one is (yyyy-MM-dd+1) 00:00;
  // 3.Return the time slice contained in start time and the time slice contained in end time.", "zh_CN":"时间
  // 1.查询的数据粒度为1m时,格式为yyyy-MM-dd HH:mm;每一个时间片数据值代表的是前一个时间粒度范围内的数据值。一天开始的时间片是yyyy-MM-dd  00:01,最后一个时间片是(yyyy-MM-dd+1)00:00;
  // 2.查询的数据粒度为5m时,格式为yyyy-MM-dd HH:mm;每一个时间片数据值代表的是前一个时间粒度范围内的数据值。一天开始的时间片是yyyy-MM-dd  00:05,最后一个时间片是(yyyy-MM-dd+1)00:00;
  // 3.返回开始时间和结束时间包含的时间片。"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {"en":"Data of the time point, two  decimal digits", "zh_CN":"时间点的数据,保留2位小数"}
  Value *string `json:"value,omitempty" xml:"value,omitempty" require:"true"`
}

func (s QueryDirectoryBandwidthTrafficUnderLivestreamDomainResponseResultDetailsDetails) String() string {
  return tea.Prettify(s)
}

func (s QueryDirectoryBandwidthTrafficUnderLivestreamDomainResponseResultDetailsDetails) GoString() string {
  return s.String()
}

func (s *QueryDirectoryBandwidthTrafficUnderLivestreamDomainResponseResultDetailsDetails) SetTimestamp(v string) *QueryDirectoryBandwidthTrafficUnderLivestreamDomainResponseResultDetailsDetails {
  s.Timestamp = &v
  return s
}

func (s *QueryDirectoryBandwidthTrafficUnderLivestreamDomainResponseResultDetailsDetails) SetValue(v string) *QueryDirectoryBandwidthTrafficUnderLivestreamDomainResponseResultDetailsDetails {
  s.Value = &v
  return s
}

type QueryDirectoryBandwidthTrafficUnderLivestreamDomainPaths struct {
}

func (s QueryDirectoryBandwidthTrafficUnderLivestreamDomainPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryDirectoryBandwidthTrafficUnderLivestreamDomainPaths) GoString() string {
  return s.String()
}

type QueryDirectoryBandwidthTrafficUnderLivestreamDomainParameters struct {
}

func (s QueryDirectoryBandwidthTrafficUnderLivestreamDomainParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryDirectoryBandwidthTrafficUnderLivestreamDomainParameters) GoString() string {
  return s.String()
}

type QueryDirectoryBandwidthTrafficUnderLivestreamDomainRequestHeader struct {
}

func (s QueryDirectoryBandwidthTrafficUnderLivestreamDomainRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryDirectoryBandwidthTrafficUnderLivestreamDomainRequestHeader) GoString() string {
  return s.String()
}

type QueryDirectoryBandwidthTrafficUnderLivestreamDomainResponseHeader struct {
}

func (s QueryDirectoryBandwidthTrafficUnderLivestreamDomainResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryDirectoryBandwidthTrafficUnderLivestreamDomainResponseHeader) GoString() string {
  return s.String()
}




type QuerySpecificPoPTrafficRequest struct {
  // {"en":"Control of returned   results
  // 
  //         1.         Option: all;
  // 
  //         2.         Only url, num are returned in the query results, and when all   is specified as the value, then returns url, num, totalNum, rate.", "zh_CN":"开始时间：
  //         1.       时间格式为yyyy-MM-ddTHH:mm:ss+08:00，例如，2016-12-02T10:00:00+08:00（为北京时间2016年12月2日10点0分0秒）；
  //         2.       不能大于当前时间
  //         3.       最多可获取最近2个月（61天）的数据。"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty"`
  // {"en":"End time:
  // 
  // 
  // 
  //         1. time format 2016-12-02T10:00:00+08:00
  // 
  // 
  // 
  //         2. the end time is greater than the start time. If the end time is greater than the current time, the current time is taken.
  // 
  // 
  // 
  //         3. dateFrom, dateTo two have not been passed, the default query for the past 15 minutes; if there is only one unsent, throw exception.
  // 
  // 
  // 
  //         4. allow maximum query interval: 1 days, that is, the difference between dateFrom and dateTo can not exceed 1 days (contact technical support adjustment).", "zh_CN":"结束时间：
  //         1.       时间格式2016-12-02T10:00:00+08:00
  //         2.       结束时间需大于开始时间，结束时间如果大于当前时间，取当前时间。
  //         3.       dateFrom，dateTo二者都未传，默认查询过去的15分钟；如仅有一个未传，抛异常
  //         4.       允许查询最大时间间隔：1天，即dateFrom和dateTo相差不能超过1天（可联系技术支持调整）。"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty"`
  // {"en":"End time:
  // 
  // 
  // 
  //         1. time format 2016-12-02T10:00:00+08:00
  // 
  // 
  // 
  //         2. the end time is greater than the start time. If the end time is greater than the current time, the current time is taken.
  // 
  // 
  // 
  //         3. dateFrom, dateTo two have not been passed, the default query for the past 15 minutes; if there is only one unsent, throw exception.
  // 
  // 
  // 
  //         4. allow maximum query interval: 1 days, that is, the difference between dateFrom and dateTo can not exceed 1 days (contact technical support adjustment).", "zh_CN":"数据粒度：
  //         1.       支持5m（5分钟粒度）
  //         2.       不传默认为5m"}
  DataInterval *string `json:"dataInterval,omitempty" xml:"dataInterval,omitempty"`
  // {"en":"Domain name:
  // 
  // 
  // 
  //         1. the maximum number of domains can be transferred to 200 by default (with technical support adjustment).
  // 
  // 
  // 
  //         2. no domain name is querying for the account without passing the entry, but no queries can be made when the number of domain names exceeds the limit.", "zh_CN":"域名：
  //         1.       可传递域名数量上限默认为200个（可联系技术支持调整）；
  //         2.       未传递该入参时查询账号下所有域名，但当账号下域名数量超过限制时不可查询（报错）。"}
  Domain []*string `json:"domain,omitempty" xml:"domain,omitempty" type:"Repeated"`
  // {"en":"Group keywords:
  // 
  //         1. Default aggregate display;
  // 
  //         2. The incoming keyword means that the details need to be displayed according to the value corresponding to the keyword; for example, groupBy passes the domain, which means that the return is expanded according to the domain details.
  // 
  //         3. Only domain can be passed.", "zh_CN":"分组关键词：
  //         1.       默认聚合展示；
  //         2.       传入关键词则代表需要按照关键词对应的值展示明细；例如groupBy传domain，则代表返回按照domain明细展开。
  //         3.       只能传递domain。"}
  GroupBy []*string `json:"groupBy,omitempty" xml:"groupBy,omitempty" type:"Repeated"`
  // {"en":"Traffic data unit
  // 
  //         1. MB, keep two decimals
  // 
  //         2. Byte
  // 
  //         3. Do not pass the default to MB", "zh_CN":"流量数据单位
  //         1. MB, 保留两位小数
  //         2. Byte
  //         3. 不传默认为MB"}
  DataUnit *string `json:"dataUnit,omitempty" xml:"dataUnit,omitempty"`
}

func (s QuerySpecificPoPTrafficRequest) String() string {
  return tea.Prettify(s)
}

func (s QuerySpecificPoPTrafficRequest) GoString() string {
  return s.String()
}

func (s *QuerySpecificPoPTrafficRequest) SetDateFrom(v string) *QuerySpecificPoPTrafficRequest {
  s.DateFrom = &v
  return s
}

func (s *QuerySpecificPoPTrafficRequest) SetDateTo(v string) *QuerySpecificPoPTrafficRequest {
  s.DateTo = &v
  return s
}

func (s *QuerySpecificPoPTrafficRequest) SetDataInterval(v string) *QuerySpecificPoPTrafficRequest {
  s.DataInterval = &v
  return s
}

func (s *QuerySpecificPoPTrafficRequest) SetDomain(v []*string) *QuerySpecificPoPTrafficRequest {
  s.Domain = v
  return s
}

func (s *QuerySpecificPoPTrafficRequest) SetGroupBy(v []*string) *QuerySpecificPoPTrafficRequest {
  s.GroupBy = v
  return s
}

func (s *QuerySpecificPoPTrafficRequest) SetDataUnit(v string) *QuerySpecificPoPTrafficRequest {
  s.DataUnit = &v
  return s
}

type QuerySpecificPoPTrafficResponse struct {
  Result []*QuerySpecificPoPTrafficResponseResult `json:"result,omitempty" xml:"result,omitempty" require:"true" type:"Repeated"`
}

func (s QuerySpecificPoPTrafficResponse) String() string {
  return tea.Prettify(s)
}

func (s QuerySpecificPoPTrafficResponse) GoString() string {
  return s.String()
}

func (s *QuerySpecificPoPTrafficResponse) SetResult(v []*QuerySpecificPoPTrafficResponseResult) *QuerySpecificPoPTrafficResponse {
  s.Result = v
  return s
}

type QuerySpecificPoPTrafficResponseResult struct     {
  // {"en":"Total number of   requests(not returned by default, and is returned only when the input   parameter of result is all)", "zh_CN":"域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  PopData []*QuerySpecificPoPTrafficResponseResultPopData `json:"popData,omitempty" xml:"popData,omitempty" require:"true" type:"Repeated"`
}

func (s QuerySpecificPoPTrafficResponseResult) String() string {
  return tea.Prettify(s)
}

func (s QuerySpecificPoPTrafficResponseResult) GoString() string {
  return s.String()
}

func (s *QuerySpecificPoPTrafficResponseResult) SetDomain(v string) *QuerySpecificPoPTrafficResponseResult {
  s.Domain = &v
  return s
}

func (s *QuerySpecificPoPTrafficResponseResult) SetPopData(v []*QuerySpecificPoPTrafficResponseResultPopData) *QuerySpecificPoPTrafficResponseResult {
  s.PopData = v
  return s
}

type QuerySpecificPoPTrafficResponseResultPopData struct     {
  // {"en":"Success rate of   downloads, four decimal digits allowed(not returned by default, and is   returned only when the input parameter of result is all)", "zh_CN":"POP名称"}
  Pop *string `json:"pop,omitempty" xml:"pop,omitempty" require:"true"`
  FlowData []*QuerySpecificPoPTrafficResponseResultPopDataFlowData `json:"flowData,omitempty" xml:"flowData,omitempty" require:"true" type:"Repeated"`
}

func (s QuerySpecificPoPTrafficResponseResultPopData) String() string {
  return tea.Prettify(s)
}

func (s QuerySpecificPoPTrafficResponseResultPopData) GoString() string {
  return s.String()
}

func (s *QuerySpecificPoPTrafficResponseResultPopData) SetPop(v string) *QuerySpecificPoPTrafficResponseResultPopData {
  s.Pop = &v
  return s
}

func (s *QuerySpecificPoPTrafficResponseResultPopData) SetFlowData(v []*QuerySpecificPoPTrafficResponseResultPopDataFlowData) *QuerySpecificPoPTrafficResponseResultPopData {
  s.FlowData = v
  return s
}

type QuerySpecificPoPTrafficResponseResultPopDataFlowData struct     {
  // {"en":"time
  //                     The data granularity of 1. queries is 5m, and the format is yyyy-MM-dd HH:mm; each time slice data value represents the data values in the previous time granularity range. The beginning of the day is yyyy-MM-dd 00:05, and the last time slice is (yyyy-MM-dd+1) 00:00.
  //                     2. return the time slice between the start time and the end time.", "zh_CN":"时间
  //                     1.       查询的数据粒度为5m时，格式为yyyy-MM-dd   HH:mm；每一个时间片数据值代表的是前一个时间粒度范围内的数据值。一天开始的时间片是yyyy-MM-dd 00:05，最后一个时间片是（yyyy-MM-dd+1）   00:00。
  //                     2.       返回开始时间和结束时间之间包含的时间片。"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {"en":"Static inflow value of the edge node,Unit:Btye.", "zh_CN":"边缘节点的静态流入流量值，单位为Btye"}
  EdgeStaticUp *string `json:"edgeStaticUp,omitempty" xml:"edgeStaticUp,omitempty" require:"true"`
  // {"en":"Static outflow value of the edge node, in units of Bitye", "zh_CN":"边缘节点的静态流出流量值，单位为Btye"}
  EdgeStaticDown *string `json:"edgeStaticDown,omitempty" xml:"edgeStaticDown,omitempty" require:"true"`
  // {"en":"Dynamic inflow value of the edge node, in units of Bitye;", "zh_CN":"边缘节点的动态流入流量值，单位为Btye；"}
  EdgeDynamicUp *string `json:"edgeDynamicUp,omitempty" xml:"edgeDynamicUp,omitempty" require:"true"`
  // {"en":"Dynamic outflow value of the edge node, in units of Bitye;", "zh_CN":"边缘节点的动态流出流量值，单位为Btye；"}
  EdgeDynamicDown *string `json:"edgeDynamicDown,omitempty" xml:"edgeDynamicDown,omitempty" require:"true"`
}

func (s QuerySpecificPoPTrafficResponseResultPopDataFlowData) String() string {
  return tea.Prettify(s)
}

func (s QuerySpecificPoPTrafficResponseResultPopDataFlowData) GoString() string {
  return s.String()
}

func (s *QuerySpecificPoPTrafficResponseResultPopDataFlowData) SetTimestamp(v string) *QuerySpecificPoPTrafficResponseResultPopDataFlowData {
  s.Timestamp = &v
  return s
}

func (s *QuerySpecificPoPTrafficResponseResultPopDataFlowData) SetEdgeStaticUp(v string) *QuerySpecificPoPTrafficResponseResultPopDataFlowData {
  s.EdgeStaticUp = &v
  return s
}

func (s *QuerySpecificPoPTrafficResponseResultPopDataFlowData) SetEdgeStaticDown(v string) *QuerySpecificPoPTrafficResponseResultPopDataFlowData {
  s.EdgeStaticDown = &v
  return s
}

func (s *QuerySpecificPoPTrafficResponseResultPopDataFlowData) SetEdgeDynamicUp(v string) *QuerySpecificPoPTrafficResponseResultPopDataFlowData {
  s.EdgeDynamicUp = &v
  return s
}

func (s *QuerySpecificPoPTrafficResponseResultPopDataFlowData) SetEdgeDynamicDown(v string) *QuerySpecificPoPTrafficResponseResultPopDataFlowData {
  s.EdgeDynamicDown = &v
  return s
}

type QuerySpecificPoPTrafficPaths struct {
}

func (s QuerySpecificPoPTrafficPaths) String() string {
  return tea.Prettify(s)
}

func (s QuerySpecificPoPTrafficPaths) GoString() string {
  return s.String()
}

type QuerySpecificPoPTrafficParameters struct {
}

func (s QuerySpecificPoPTrafficParameters) String() string {
  return tea.Prettify(s)
}

func (s QuerySpecificPoPTrafficParameters) GoString() string {
  return s.String()
}

type QuerySpecificPoPTrafficRequestHeader struct {
}

func (s QuerySpecificPoPTrafficRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QuerySpecificPoPTrafficRequestHeader) GoString() string {
  return s.String()
}

type QuerySpecificPoPTrafficResponseHeader struct {
}

func (s QuerySpecificPoPTrafficResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QuerySpecificPoPTrafficResponseHeader) GoString() string {
  return s.String()
}




type QueryTotalTrafficForMultiDomainsRequest struct {
  // {"en":"Domain list.\nDomain number limits can be adjusted depending on different accounts. The default value is 1000(if you want to adjust,please, contact technical support)","zh_CN":"域名列表\n1.域名个数限制根据账号可调,默认为1000个(可联系技术支持下单调整);"}
  DomainList *QueryTotalTrafficForMultiDomainsRequestDomainList `json:"domain-list,omitempty" xml:"domain-list,omitempty" type:"Struct"`
}

func (s QueryTotalTrafficForMultiDomainsRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryTotalTrafficForMultiDomainsRequest) GoString() string {
  return s.String()
}

func (s *QueryTotalTrafficForMultiDomainsRequest) SetDomainList(v *QueryTotalTrafficForMultiDomainsRequestDomainList) *QueryTotalTrafficForMultiDomainsRequest {
  s.DomainList = v
  return s
}

type QueryTotalTrafficForMultiDomainsRequestDomainList struct {
  // {"en":"Domain","zh_CN":"域名"}
  DomainName []*string `json:"domain-name,omitempty" xml:"domain-name,omitempty" type:"Repeated"`
}

func (s QueryTotalTrafficForMultiDomainsRequestDomainList) String() string {
  return tea.Prettify(s)
}

func (s QueryTotalTrafficForMultiDomainsRequestDomainList) GoString() string {
  return s.String()
}

func (s *QueryTotalTrafficForMultiDomainsRequestDomainList) SetDomainName(v []*string) *QueryTotalTrafficForMultiDomainsRequestDomainList {
  s.DomainName = v
  return s
}

type QueryTotalTrafficForMultiDomainsRequestHeader struct {
}

func (s QueryTotalTrafficForMultiDomainsRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryTotalTrafficForMultiDomainsRequestHeader) GoString() string {
  return s.String()
}

type QueryTotalTrafficForMultiDomainsPaths struct {
}

func (s QueryTotalTrafficForMultiDomainsPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryTotalTrafficForMultiDomainsPaths) GoString() string {
  return s.String()
}

type QueryTotalTrafficForMultiDomainsParameters struct {
  // {"en":"Start time\n1.The time format is yyyy-MM-ddTHH:mm:ss±HH:mm. Please note: ±HH:mm is the time zone offset, which can be adjusted according to your data needs, for example, +00:00 represents UTC time, +08:00 represents East 8th District, and -05:00 represents West 5th District. 2024-01-15T10:30:45+00:00 means UTC time January 15, 2024 10:30:45 AM;\n2.And smaller than the current time and 'dateTo';\n3.Period between 'dataFrom' and 'dateTo' cannot be longer than 31 days","zh_CN":"开始时间\n1.时间格式为 yyyy-MM-ddTHH:mm:ss±HH:mm。请注意：±HH:mm 为时区偏移量，可根据您的数据需要进行调整，例如 +00:00 代表 UTC 时间，+08:00 代表东八区，-05:00 代表西五区。2024-01-15T10:30:45+00:00，表示UTC 时间 2024 年 1 月 15 日上午 10 点 30 分 45 秒;\n2.并且小于当前时间和dateTo;\n3.dateFrom和dateTo相差不能超过31天;4.只能查询最近2年内数据。"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty" require:"true"`
  // {"en":"End time\n1.The time format is yyyy-MM-ddTHH:mm:ss±HH:mm. Please note: ±HH:mm is the time zone offset, which can be adjusted according to your data needs, for example, +00:00 represents UTC time, +08:00 represents East 8th District, and -05:00 represents West 5th District. 2024-01-15T10:30:45+00:00 means UTC time January 15, 2024 10:30:45 AM;\n2.Must be greater than 'dateFrom';\n3.If it's greater than the current time, then the current time is assigned as the value","zh_CN":"结束时间\n1.时间格式为 yyyy-MM-ddTHH:mm:ss±HH:mm。请注意：±HH:mm 为时区偏移量，可根据您的数据需要进行调整，例如 +00:00 代表 UTC 时间，+08:00 代表东八区，-05:00 代表西五区。2024-01-15T10:30:45+00:00，表示UTC 时间 2024 年 1 月 15 日上午 10 点 30 分 45 秒;\n2.必须大于dateFrom;\n3.如果大于当前时间,则重新赋值为当前时间;"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty" require:"true"`
  // {"defaultValue":"daily","en":"Data granularity\n1.fiveminutes: five minutes, hourly: one hour, daily: one day;\n2.If not specified, daily is set as the default value;\n3.If fiveminutes is specified as the value, then data is returned in actual configured granularity when there is specific configuration to data collecting granularity for the customer.","zh_CN":"数据粒度\n1.fiveminutes:5分钟,hourly:1小时,daily:1天;\n2.不传递,默认为daily;\n3.传递fiveminutes时,若客户数据采集粒度有特殊配置将按实际配置粒度返回。","exampleValue":"fiveminutes,hourly,daily"}
  Granularity *string `json:"granularity,omitempty" xml:"granularity,omitempty"`
}

func (s QueryTotalTrafficForMultiDomainsParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryTotalTrafficForMultiDomainsParameters) GoString() string {
  return s.String()
}

func (s *QueryTotalTrafficForMultiDomainsParameters) SetDateFrom(v string) *QueryTotalTrafficForMultiDomainsParameters {
  s.DateFrom = &v
  return s
}

func (s *QueryTotalTrafficForMultiDomainsParameters) SetDateTo(v string) *QueryTotalTrafficForMultiDomainsParameters {
  s.DateTo = &v
  return s
}

func (s *QueryTotalTrafficForMultiDomainsParameters) SetGranularity(v string) *QueryTotalTrafficForMultiDomainsParameters {
  s.Granularity = &v
  return s
}

type QueryTotalTrafficForMultiDomainsResponse struct {
  // {"en":"Total traffic. Unit:  MB, retain two decimals","zh_CN":"总流量, 单位MB，保留2位小数"}
  TotalTraffic *int `json:"totalTraffic,omitempty" xml:"totalTraffic,omitempty" require:"true"`
  // {"en":"dataSeries","zh_CN":"流量数据"}
  DataSeries []*QueryTotalTrafficForMultiDomainsResponseDataSeries `json:"dataSeries,omitempty" xml:"dataSeries,omitempty" require:"true" type:"Repeated"`
}

func (s QueryTotalTrafficForMultiDomainsResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryTotalTrafficForMultiDomainsResponse) GoString() string {
  return s.String()
}

func (s *QueryTotalTrafficForMultiDomainsResponse) SetTotalTraffic(v int) *QueryTotalTrafficForMultiDomainsResponse {
  s.TotalTraffic = &v
  return s
}

func (s *QueryTotalTrafficForMultiDomainsResponse) SetDataSeries(v []*QueryTotalTrafficForMultiDomainsResponseDataSeries) *QueryTotalTrafficForMultiDomainsResponse {
  s.DataSeries = v
  return s
}

type QueryTotalTrafficForMultiDomainsResponseDataSeries struct     {
  // {"en":"Date\n1.When the data query granularity is fiveminutes, the format is yyyy-MM-dd HH:mm; the data value of every time slice represents the data value within the previous time granularity range. The first time slice of the day is yyyy-MM-dd 00:05 AM, and the last one is yyyy-MM-dd 24:00.\n2.When the data query granularity is hourly, the format is yyyy-MM-dd HH; the data value of every time slice represents the data value within the previous time granularity range. The first time slice of the day is yyyy-MM-dd 00:01, and the last one is yyyy-MM-dd 24.\n3.When the data query granularity is daily, the format is yyyy-MM-dd; the data value of every time slice represents the value of the daily data.Return the time slice contained in start time and the time slice contained in end time","zh_CN":"时间\n1.查询的数据粒度为fiveminutes时,格式为yyyy-MM-dd HH:mm;每一个时间片数据值代表的是前一个时间粒度范围内的数据值。一天开始的时间片是yyyy-MM-dd 00:05,最后一个时间片是yyyy-MM-dd 24:00。\n2.查询的数据粒度为hourly时,格式为yyyy-MM-dd HH;每一个时间片数据值代表的是前一个时间粒度范围内的数据值。一天开始的时间片是yyyy-MM-dd 01,最后一个时间片是yyyy-MM-dd 24。\n3.查询的数据粒度为daily时,格式为yyyy-MM-dd;每一个时间片数据值代表的该天内的数据值。\n4.返回开始时间和结束时间包含的时间片。"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {"en":"Traffic value:\nUnit:  MB, retain two decimals","zh_CN":"流量值：\n\n单位MB，保留2位小数"}
  Traffic *string `json:"traffic,omitempty" xml:"traffic,omitempty" require:"true"`
}

func (s QueryTotalTrafficForMultiDomainsResponseDataSeries) String() string {
  return tea.Prettify(s)
}

func (s QueryTotalTrafficForMultiDomainsResponseDataSeries) GoString() string {
  return s.String()
}

func (s *QueryTotalTrafficForMultiDomainsResponseDataSeries) SetTimestamp(v string) *QueryTotalTrafficForMultiDomainsResponseDataSeries {
  s.Timestamp = &v
  return s
}

func (s *QueryTotalTrafficForMultiDomainsResponseDataSeries) SetTraffic(v string) *QueryTotalTrafficForMultiDomainsResponseDataSeries {
  s.Traffic = &v
  return s
}

type QueryTotalTrafficForMultiDomainsResponseHeader struct {
}

func (s QueryTotalTrafficForMultiDomainsResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryTotalTrafficForMultiDomainsResponseHeader) GoString() string {
  return s.String()
}




type QueryStreamTrafficUnderMALBDomainRequest struct {
  // {"en":"Starting time
  // 			1. The format is yyyy-MM-ddTHH:mm:ss+08:00;
  // 			2. Must be less than the current time and dateTo;
  // 			3. The difference between dateFrom and dateTo cannot exceed 1 day(technical support can be contacted to adjust);
  // 			4. Only data within the last 6 months can be queried.", "zh_CN":"开始时间
  // 			1.格式为yyyy-MM-ddTHH:mm:ss+08:00；
  // 			2.必须小于当前时间和dateTo；
  // 			3.dateFrom和dateTo相差不能超过1天(可联系技术支持调整)；
  // 			4.只能查询最近6个月内数据。"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty" require:"true"`
  // {"en":"End time:
  // 			1. The format is yyyy-MM-ddTHH:mm:ss+08:00;
  // 			2. Must be greater than dateFrom;
  // 			3. If it's greater than the current time, then the current time is assigned as the value;", "zh_CN":"结束时间
  // 			1.格式为yyyy-MM-ddTHH:mm:ss+08:00；
  // 			2.必须大于dateFrom；
  // 			3.如果大于当前时间，则重新赋值为当前时间；"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty" require:"true"`
  // {"en":"Domain and stream name list", "zh_CN":"域名和流名组：
  // 						1.可传递的域名和流名组数量上限默认为20组(可联系技术支持调整)；
  // 						2.域名domain：一组域名和流名组中只能传递单个域名，且域名必须传递；
  // 						3.流名stream：一组域名和流名组下(即单个域名下)可传递的流名数量上限默认为2000个(可联系技术支持调整)，流名未传递时默认查询域名下所有流名，但当域名下流名数量超过限制时不可查询(报错)。"}
  DomainStream []*QueryStreamTrafficUnderMALBDomainRequestDomainStream `json:"domainStream,omitempty" xml:"domainStream,omitempty" require:"true" type:"Repeated"`
  // {"en":"Data granularity:
  // 				1. 1m: 1-minute granularity, 5m: 5-minute granularity;
  // 				2. Data with granularity of 5 minutes is queried by default, if you need to query the granularity of 1 minute, please contact technical support for configuration", "zh_CN":"数据粒度
  // 				1.1m：1分钟粒度，5m：5分钟粒度；
  // 				2.默认查询5分钟粒度数据，若需要查询1m粒度请联系技术支持进行特殊配置"}
  DataInterval *string `json:"dataInterval,omitempty" xml:"dataInterval,omitempty"`
  // {"en":"Data type: flow, bandwidth", "zh_CN":"数据类型，flow：流量，bandwidth：带宽"}
  DataType *string `json:"dataType,omitempty" xml:"dataType,omitempty" require:"true"`
  // {"en":"Filter nullname stream:
  // 			1.Enter optional value '0' or '1';
  // 			2. The parameter 0 does not filter the data whose data is null, and the parameter 1 filter the data whose data is null.
  // 			3. the default value is 0;", "zh_CN":"是否过滤空流名:
  // 			1.入参可选值 '0' 或 '1' ；
  // 			2.传参 0 不过滤流名为空的数据，传参 1 则过滤流名为空的数据 ；
  // 			3.默认值为 0 ；"}
  FilterEmptyStream *int `json:"filterEmptyStream,omitempty" xml:"filterEmptyStream,omitempty"`
}

func (s QueryStreamTrafficUnderMALBDomainRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryStreamTrafficUnderMALBDomainRequest) GoString() string {
  return s.String()
}

func (s *QueryStreamTrafficUnderMALBDomainRequest) SetDateFrom(v string) *QueryStreamTrafficUnderMALBDomainRequest {
  s.DateFrom = &v
  return s
}

func (s *QueryStreamTrafficUnderMALBDomainRequest) SetDateTo(v string) *QueryStreamTrafficUnderMALBDomainRequest {
  s.DateTo = &v
  return s
}

func (s *QueryStreamTrafficUnderMALBDomainRequest) SetDomainStream(v []*QueryStreamTrafficUnderMALBDomainRequestDomainStream) *QueryStreamTrafficUnderMALBDomainRequest {
  s.DomainStream = v
  return s
}

func (s *QueryStreamTrafficUnderMALBDomainRequest) SetDataInterval(v string) *QueryStreamTrafficUnderMALBDomainRequest {
  s.DataInterval = &v
  return s
}

func (s *QueryStreamTrafficUnderMALBDomainRequest) SetDataType(v string) *QueryStreamTrafficUnderMALBDomainRequest {
  s.DataType = &v
  return s
}

func (s *QueryStreamTrafficUnderMALBDomainRequest) SetFilterEmptyStream(v int) *QueryStreamTrafficUnderMALBDomainRequest {
  s.FilterEmptyStream = &v
  return s
}

type QueryStreamTrafficUnderMALBDomainRequestDomainStream struct     {
  // {"en":"Domain names, domain number limits can be adjusted depending on different accounts. The default value is 20", "zh_CN":"可传递域名数量上限默认为20个(可联系技术支持调整)"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {"en":"Stream name:
  // 				1. If this field is not specified, it means all domains under the domain are queried;
  // 				2. Number of streams can be adjusted depending on different accounts. The default value is 2000(this limit applies to the empty value);", "zh_CN":"流名：'发布点'+'流名'。例如：live/test-20180101-test ,其中live是发布点，test-20180101-test是流名。"}
  Stream []*string `json:"stream,omitempty" xml:"stream,omitempty" type:"Repeated"`
}

func (s QueryStreamTrafficUnderMALBDomainRequestDomainStream) String() string {
  return tea.Prettify(s)
}

func (s QueryStreamTrafficUnderMALBDomainRequestDomainStream) GoString() string {
  return s.String()
}

func (s *QueryStreamTrafficUnderMALBDomainRequestDomainStream) SetDomain(v string) *QueryStreamTrafficUnderMALBDomainRequestDomainStream {
  s.Domain = &v
  return s
}

func (s *QueryStreamTrafficUnderMALBDomainRequestDomainStream) SetStream(v []*string) *QueryStreamTrafficUnderMALBDomainRequestDomainStream {
  s.Stream = v
  return s
}

type QueryStreamTrafficUnderMALBDomainResponse struct {
  Result []*QueryStreamTrafficUnderMALBDomainResponseResult `json:"result,omitempty" xml:"result,omitempty" require:"true" type:"Repeated"`
}

func (s QueryStreamTrafficUnderMALBDomainResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryStreamTrafficUnderMALBDomainResponse) GoString() string {
  return s.String()
}

func (s *QueryStreamTrafficUnderMALBDomainResponse) SetResult(v []*QueryStreamTrafficUnderMALBDomainResponseResult) *QueryStreamTrafficUnderMALBDomainResponse {
  s.Result = v
  return s
}

type QueryStreamTrafficUnderMALBDomainResponseResult struct     {
  // {"en":"Domain", "zh_CN":"域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  Details []*QueryStreamTrafficUnderMALBDomainResponseResultDetails `json:"details,omitempty" xml:"details,omitempty" require:"true" type:"Repeated"`
}

func (s QueryStreamTrafficUnderMALBDomainResponseResult) String() string {
  return tea.Prettify(s)
}

func (s QueryStreamTrafficUnderMALBDomainResponseResult) GoString() string {
  return s.String()
}

func (s *QueryStreamTrafficUnderMALBDomainResponseResult) SetDomain(v string) *QueryStreamTrafficUnderMALBDomainResponseResult {
  s.Domain = &v
  return s
}

func (s *QueryStreamTrafficUnderMALBDomainResponseResult) SetDetails(v []*QueryStreamTrafficUnderMALBDomainResponseResultDetails) *QueryStreamTrafficUnderMALBDomainResponseResult {
  s.Details = v
  return s
}

type QueryStreamTrafficUnderMALBDomainResponseResultDetails struct     {
  // {"en":"Stream name", "zh_CN":"流名"}
  Stream *string `json:"stream,omitempty" xml:"stream,omitempty" require:"true"`
  // {"en":"Total traffic:
  // 					1. Keep two digits of decimals. Unit: MB;
  // 					2. Return when the input parameter of datatype is flow.", "zh_CN":"总流量
  // 					1.保留2位小数，单位为MB;
  // 					2.当入参dataType为flow时，返回。"}
  TotalFlow *string `json:"totalFlow,omitempty" xml:"totalFlow,omitempty"`
  // {"en":"Peak value of bandwidth
  // 					1. Keep two digits of decimals. Unit: Mbps
  // 					2. Return when the input dataType is bandwidth.", "zh_CN":"峰值带宽
  // 					1.保留2位小数，单位Mbps
  // 					2.当入参dataType为bandwidth时，返回。"}
  BandwidthPeakValue *string `json:"bandwidthPeakValue,omitempty" xml:"bandwidthPeakValue,omitempty"`
  Details []*QueryStreamTrafficUnderMALBDomainResponseResultDetailsDetails `json:"details,omitempty" xml:"details,omitempty" require:"true" type:"Repeated"`
}

func (s QueryStreamTrafficUnderMALBDomainResponseResultDetails) String() string {
  return tea.Prettify(s)
}

func (s QueryStreamTrafficUnderMALBDomainResponseResultDetails) GoString() string {
  return s.String()
}

func (s *QueryStreamTrafficUnderMALBDomainResponseResultDetails) SetStream(v string) *QueryStreamTrafficUnderMALBDomainResponseResultDetails {
  s.Stream = &v
  return s
}

func (s *QueryStreamTrafficUnderMALBDomainResponseResultDetails) SetTotalFlow(v string) *QueryStreamTrafficUnderMALBDomainResponseResultDetails {
  s.TotalFlow = &v
  return s
}

func (s *QueryStreamTrafficUnderMALBDomainResponseResultDetails) SetBandwidthPeakValue(v string) *QueryStreamTrafficUnderMALBDomainResponseResultDetails {
  s.BandwidthPeakValue = &v
  return s
}

func (s *QueryStreamTrafficUnderMALBDomainResponseResultDetails) SetDetails(v []*QueryStreamTrafficUnderMALBDomainResponseResultDetailsDetails) *QueryStreamTrafficUnderMALBDomainResponseResultDetails {
  s.Details = v
  return s
}

type QueryStreamTrafficUnderMALBDomainResponseResultDetailsDetails struct     {
  // {"en":"Date:
  // 						1. When the data query granularity is 1m, the format is yyyy-MM-dd HH:mm; the data value of every time slice represents the data value within the previous time granularity range. The first time slice of the day is yyyy-MM-dd 00:01, and the last one is (yyyy-MM-dd+1) 00:00.
  // 						2. When the data query granularity is 5m, the format is yyyy-MM-dd HH:mm; the data value of every time slice represents the data value within the previous time granularity range. The first time slice of the day is yyyy-MM-dd 12:05 AM, and the last one is (yyyy-MM-dd+1) 00:00.", "zh_CN":"时间
  // 						1.查询的数据粒度为1m时，格式为yyyy-MM-dd HH:mm；每一个时间片数据值代表的是前一个时间粒度范围内的数据值。一天开始的时间片是yyyy-MM-dd 00:01，最后一个时间片是(yyyy-MM-dd+1)&nbsp;00:00。
  // 						2.查询的数据粒度为5m时，格式为yyyy-MM-dd HH:mm；每一个时间片数据值代表的是前一个时间粒度范围内的数据值。一天开始的时间片是yyyy-MM-dd 00:05，最后一个时间片是(yyyy-MM-dd+1)&nbsp;00:00。"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {"en":"Data of the time point:
  // 						1. Two decimal digits allowed;
  // 						2. When the input parameter of dataType is flow, the value is the flow/traffic value, with MB as the unit;
  // 						3. When the input parameter of dataType is bandwidth, the value is the bandwidth value, with Mbps as the unit;", "zh_CN":"时间点的数据
  // 						1.保留2位小数；
  // 						2.当入参dataType为flow时，value值为流量，单位为MB；
  // 						3.当入参dataType为bandwidth时，value值为带宽值，单位为Mbps；"}
  Value *string `json:"value,omitempty" xml:"value,omitempty" require:"true"`
}

func (s QueryStreamTrafficUnderMALBDomainResponseResultDetailsDetails) String() string {
  return tea.Prettify(s)
}

func (s QueryStreamTrafficUnderMALBDomainResponseResultDetailsDetails) GoString() string {
  return s.String()
}

func (s *QueryStreamTrafficUnderMALBDomainResponseResultDetailsDetails) SetTimestamp(v string) *QueryStreamTrafficUnderMALBDomainResponseResultDetailsDetails {
  s.Timestamp = &v
  return s
}

func (s *QueryStreamTrafficUnderMALBDomainResponseResultDetailsDetails) SetValue(v string) *QueryStreamTrafficUnderMALBDomainResponseResultDetailsDetails {
  s.Value = &v
  return s
}

type QueryStreamTrafficUnderMALBDomainPaths struct {
}

func (s QueryStreamTrafficUnderMALBDomainPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryStreamTrafficUnderMALBDomainPaths) GoString() string {
  return s.String()
}

type QueryStreamTrafficUnderMALBDomainParameters struct {
}

func (s QueryStreamTrafficUnderMALBDomainParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryStreamTrafficUnderMALBDomainParameters) GoString() string {
  return s.String()
}

type QueryStreamTrafficUnderMALBDomainRequestHeader struct {
}

func (s QueryStreamTrafficUnderMALBDomainRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryStreamTrafficUnderMALBDomainRequestHeader) GoString() string {
  return s.String()
}

type QueryStreamTrafficUnderMALBDomainResponseHeader struct {
}

func (s QueryStreamTrafficUnderMALBDomainResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryStreamTrafficUnderMALBDomainResponseHeader) GoString() string {
  return s.String()
}




type QueryDataTransferForAllDomainByteRequest struct {
  // {"en":"Start time:
  // 1. Time format is yyyy-MM-ddTHH:mm:ss+08:00,
  // 2. No bigger than the current time.
  // 3. Data in the last 2 years at most can be queried.", "zh_CN":"开始时间:
  // 1.时间格式为yyyy-MM-ddTHH:mm:ss+08:00
  // 2.不能大于当前时间
  // 3.最多可获取最近2年内的数据"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty"`
  // {"en":"End time:
  // 1. The time format is yyyy-MM-ddTHH:mm:ss+08:00
  // 2. End time should be greater than start time. If the end time is greater than current time, current time will be used.
  // 3. If both fields of dataFrom and dateTo are left empty, then data in the last 1day will be queried by default; if only one field is filled in and one is left empty, then exception will be occur.
  // 4. Allowable maximum time range for query: 31days, means the period between dateFrom to dateTo should not exceed 31days (can be adjusted by contacting technical support).", "zh_CN":"结束时间:
  // 1.时间格式yyyy-MM-ddTHH:mm:ss+08:00
  // 2.结束时间需大于开始时间,结束时间如果大于当前时间,取当前时间,
  // 3.dateFrom,dateTo二者都未传,默认查询过去的24小时;如仅有一个未传,抛异常;
  // 4.允许查询最大时间间隔31天:,即dateFrom和dateTo相差不能超过31天。(可联系技术支持调整)"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty"`
  // {"en":"Data granularity, 5m: 5-minute   granularity, 1h: 1-hour granularity, 1d: 1-day granularity ", "zh_CN":"数据粒度, 5m:5分钟粒度,1h:1小时粒度, 1d: 1天粒度"}
  DataInterval *string `json:"dataInterval,omitempty" xml:"dataInterval,omitempty"`
}

func (s QueryDataTransferForAllDomainByteRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryDataTransferForAllDomainByteRequest) GoString() string {
  return s.String()
}

func (s *QueryDataTransferForAllDomainByteRequest) SetDateFrom(v string) *QueryDataTransferForAllDomainByteRequest {
  s.DateFrom = &v
  return s
}

func (s *QueryDataTransferForAllDomainByteRequest) SetDateTo(v string) *QueryDataTransferForAllDomainByteRequest {
  s.DateTo = &v
  return s
}

func (s *QueryDataTransferForAllDomainByteRequest) SetDataInterval(v string) *QueryDataTransferForAllDomainByteRequest {
  s.DataInterval = &v
  return s
}

type QueryDataTransferForAllDomainByteResponse struct {
  // {"en":"request result status code", "zh_CN":"请求结果状态码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Request result information", "zh_CN":"请求结果信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  Data []*QueryDataTransferForAllDomainByteResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s QueryDataTransferForAllDomainByteResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryDataTransferForAllDomainByteResponse) GoString() string {
  return s.String()
}

func (s *QueryDataTransferForAllDomainByteResponse) SetCode(v string) *QueryDataTransferForAllDomainByteResponse {
  s.Code = &v
  return s
}

func (s *QueryDataTransferForAllDomainByteResponse) SetMessage(v string) *QueryDataTransferForAllDomainByteResponse {
  s.Message = &v
  return s
}

func (s *QueryDataTransferForAllDomainByteResponse) SetData(v []*QueryDataTransferForAllDomainByteResponseData) *QueryDataTransferForAllDomainByteResponse {
  s.Data = v
  return s
}

type QueryDataTransferForAllDomainByteResponseData struct     {
  // {"en":"flowSummary", "zh_CN":"总流量"}
  FlowSummary *string `json:"flowSummary,omitempty" xml:"flowSummary,omitempty" require:"true"`
  FlowData []*QueryDataTransferForAllDomainByteResponseDataFlowData `json:"flowData,omitempty" xml:"flowData,omitempty" require:"true" type:"Repeated"`
}

func (s QueryDataTransferForAllDomainByteResponseData) String() string {
  return tea.Prettify(s)
}

func (s QueryDataTransferForAllDomainByteResponseData) GoString() string {
  return s.String()
}

func (s *QueryDataTransferForAllDomainByteResponseData) SetFlowSummary(v string) *QueryDataTransferForAllDomainByteResponseData {
  s.FlowSummary = &v
  return s
}

func (s *QueryDataTransferForAllDomainByteResponseData) SetFlowData(v []*QueryDataTransferForAllDomainByteResponseDataFlowData) *QueryDataTransferForAllDomainByteResponseData {
  s.FlowData = v
  return s
}

type QueryDataTransferForAllDomainByteResponseDataFlowData struct     {
  // {"en":"timestamp", "zh_CN":"时间"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {"en":"flow", "zh_CN":"流量"}
  Flow *string `json:"flow,omitempty" xml:"flow,omitempty" require:"true"`
}

func (s QueryDataTransferForAllDomainByteResponseDataFlowData) String() string {
  return tea.Prettify(s)
}

func (s QueryDataTransferForAllDomainByteResponseDataFlowData) GoString() string {
  return s.String()
}

func (s *QueryDataTransferForAllDomainByteResponseDataFlowData) SetTimestamp(v string) *QueryDataTransferForAllDomainByteResponseDataFlowData {
  s.Timestamp = &v
  return s
}

func (s *QueryDataTransferForAllDomainByteResponseDataFlowData) SetFlow(v string) *QueryDataTransferForAllDomainByteResponseDataFlowData {
  s.Flow = &v
  return s
}

type QueryDataTransferForAllDomainBytePaths struct {
}

func (s QueryDataTransferForAllDomainBytePaths) String() string {
  return tea.Prettify(s)
}

func (s QueryDataTransferForAllDomainBytePaths) GoString() string {
  return s.String()
}

type QueryDataTransferForAllDomainByteParameters struct {
}

func (s QueryDataTransferForAllDomainByteParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryDataTransferForAllDomainByteParameters) GoString() string {
  return s.String()
}

type QueryDataTransferForAllDomainByteRequestHeader struct {
}

func (s QueryDataTransferForAllDomainByteRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryDataTransferForAllDomainByteRequestHeader) GoString() string {
  return s.String()
}

type QueryDataTransferForAllDomainByteResponseHeader struct {
}

func (s QueryDataTransferForAllDomainByteResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryDataTransferForAllDomainByteResponseHeader) GoString() string {
  return s.String()
}




type GetBandwidthSavingRatioRequest struct {
  // {"en":"Start time:\n1.The time format is yyyy-MM-ddTHH:mm:ss±HH:mm. Please note: ±HH:mm is the time zone offset, which can be adjusted according to your data needs, for example, +00:00 represents UTC time, +08:00 represents East 8th District, and -05:00 represents West 5th District. 2024-01-15T10:30:45+00:00 means UTC time January 15, 2024 10:30:45 AM; \n2.Cannot exceed current time;\n3.The most recent six-month (183 days) data are available.","zh_CN":"开始时间:\n1.时间格式为 yyyy-MM-ddTHH:mm:ss±HH:mm。请注意：±HH:mm 为时区偏移量，可根据您的数据需要进行调整，例如 +00:00 代表 UTC 时间，+08:00 代表东八区，-05:00 代表西五区。2024-01-15T10:30:45+00:00，表示UTC 时间 2024 年 1 月 15 日上午 10 点 30 分 45 秒; \n2.不能大于当前时间\n3.最多可获取最近半年(183天)的数据。"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty"`
  // {"en":"End time:\n1.The time format is yyyy-MM-ddTHH:mm:ss±HH:mm. Please note: ±HH:mm is the time zone offset, which can be adjusted according to your data needs, for example, +00:00 represents UTC time, +08:00 represents East 8th District, and -05:00 represents West 5th District. 2024-01-15T10:30:45+00:00 means UTC time January 15, 2024 10:30:45 AM; \n2.The end time needs to be greater than the start time. If the end time is greater than the current time, take the current time;\n3.Date from, Date to both, the default query past 24 hours; If there is only one unsent, throw an exception;\n4.Maximum allowed query time interval: 31 days, Date from and dateTo, not more than 31 days","zh_CN":"结束时间:\n1.时间格式为 yyyy-MM-ddTHH:mm:ss±HH:mm。请注意：±HH:mm 为时区偏移量，可根据您的数据需要进行调整，例如 +00:00 代表 UTC 时间，+08:00 代表东八区，-05:00 代表西五区。2024-01-15T10:30:45+00:00，表示UTC 时间 2024 年 1 月 15 日上午 10 点 30 分 45 秒; \n2.结束时间需大于开始时间,结束时间如果大于当前时间,取当前时间。\n3.dateFrom,dateTo二者都未传,默认查询过去的24小时;如仅有一个未传,抛异常\n4.允许查询最大时间间隔:31天,即dateFrom和dateTo相差不能超过31天"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty"`
  // {"en":"Domain:\n1.The maximum number of deliverable domain names is 200 by default;\n2.Automatically filter out illegal domain names (pass illegal domain names, will be filtered out, the query results only return the data of legitimate domain names);\n3.The default query accounts for all domains if the number of domain names exceeds the upper limit when the entry is not delivered. If the number of domain names in the account exceeds the limit, an error is raised.","zh_CN":"域名:\n1.可传递域名数量上限默认为200个\n2.自动过滤掉非法域名(如传递非法域名,会被过滤掉,查询结果只返回合法域名的数据)\n3.未传递该入参时,默认查询账号下所有域名,但当账号下域名数量超过上限时提示错误。"}
  Domain []*string `json:"domain,omitempty" xml:"domain,omitempty" require:"true" type:"Repeated"`
  // {"defaultValue":"1d","en":"Time interval of data: 5m (5 min), 1h (1 hour), 1d (1 day);\nThe default is 1d.","zh_CN":"数据粒度:\n1.支持5m(5分钟)、1h(1小时)、1d(天)\n2.不传默认1d。","exampleValue":"5m,1h,1d"}
  Granularity *string `json:"granularity,omitempty" xml:"granularity,omitempty"`
}

func (s GetBandwidthSavingRatioRequest) String() string {
  return tea.Prettify(s)
}

func (s GetBandwidthSavingRatioRequest) GoString() string {
  return s.String()
}

func (s *GetBandwidthSavingRatioRequest) SetDateFrom(v string) *GetBandwidthSavingRatioRequest {
  s.DateFrom = &v
  return s
}

func (s *GetBandwidthSavingRatioRequest) SetDateTo(v string) *GetBandwidthSavingRatioRequest {
  s.DateTo = &v
  return s
}

func (s *GetBandwidthSavingRatioRequest) SetDomain(v []*string) *GetBandwidthSavingRatioRequest {
  s.Domain = v
  return s
}

func (s *GetBandwidthSavingRatioRequest) SetGranularity(v string) *GetBandwidthSavingRatioRequest {
  s.Granularity = &v
  return s
}

type GetBandwidthSavingRatioRequestHeader struct {
}

func (s GetBandwidthSavingRatioRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetBandwidthSavingRatioRequestHeader) GoString() string {
  return s.String()
}

type GetBandwidthSavingRatioPaths struct {
}

func (s GetBandwidthSavingRatioPaths) String() string {
  return tea.Prettify(s)
}

func (s GetBandwidthSavingRatioPaths) GoString() string {
  return s.String()
}

type GetBandwidthSavingRatioParameters struct {
}

func (s GetBandwidthSavingRatioParameters) String() string {
  return tea.Prettify(s)
}

func (s GetBandwidthSavingRatioParameters) GoString() string {
  return s.String()
}

type GetBandwidthSavingRatioResponse struct {
  // {"en":"Request result status code","zh_CN":"请求结果状态码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Request result information","zh_CN":"请求结果信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"","zh_CN":""}
  Data []*GetBandwidthSavingRatioResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s GetBandwidthSavingRatioResponse) String() string {
  return tea.Prettify(s)
}

func (s GetBandwidthSavingRatioResponse) GoString() string {
  return s.String()
}

func (s *GetBandwidthSavingRatioResponse) SetCode(v string) *GetBandwidthSavingRatioResponse {
  s.Code = &v
  return s
}

func (s *GetBandwidthSavingRatioResponse) SetMessage(v string) *GetBandwidthSavingRatioResponse {
  s.Message = &v
  return s
}

func (s *GetBandwidthSavingRatioResponse) SetData(v []*GetBandwidthSavingRatioResponseData) *GetBandwidthSavingRatioResponse {
  s.Data = v
  return s
}

type GetBandwidthSavingRatioResponseData struct     {
  // {"en":"Actually processed time.  yyyy-MM-dd HH:mm format","zh_CN":"实际查询时间,格式 yyyy-MM-dd HH:mm"}
  QueryTime *string `json:"queryTime,omitempty" xml:"queryTime,omitempty" require:"true"`
  // {"en":"Average of total saving of bandwidth.","zh_CN":"总节省带宽的平均值"}
  AvgSavedBandwidthRatio *GetBandwidthSavingRatioResponseDataAvgSavedBandwidthRatio `json:"avgSavedBandwidthRatio,omitempty" xml:"avgSavedBandwidthRatio,omitempty" require:"true" type:"Struct"`
  // {"en":"","zh_CN":""}
  DataSeries []*GetBandwidthSavingRatioResponseDataDataSeries `json:"dataSeries,omitempty" xml:"dataSeries,omitempty" require:"true" type:"Repeated"`
}

func (s GetBandwidthSavingRatioResponseData) String() string {
  return tea.Prettify(s)
}

func (s GetBandwidthSavingRatioResponseData) GoString() string {
  return s.String()
}

func (s *GetBandwidthSavingRatioResponseData) SetQueryTime(v string) *GetBandwidthSavingRatioResponseData {
  s.QueryTime = &v
  return s
}

func (s *GetBandwidthSavingRatioResponseData) SetAvgSavedBandwidthRatio(v *GetBandwidthSavingRatioResponseDataAvgSavedBandwidthRatio) *GetBandwidthSavingRatioResponseData {
  s.AvgSavedBandwidthRatio = v
  return s
}

func (s *GetBandwidthSavingRatioResponseData) SetDataSeries(v []*GetBandwidthSavingRatioResponseDataDataSeries) *GetBandwidthSavingRatioResponseData {
  s.DataSeries = v
  return s
}

type GetBandwidthSavingRatioResponseDataAvgSavedBandwidthRatio struct {
}

func (s GetBandwidthSavingRatioResponseDataAvgSavedBandwidthRatio) String() string {
  return tea.Prettify(s)
}

func (s GetBandwidthSavingRatioResponseDataAvgSavedBandwidthRatio) GoString() string {
  return s.String()
}

type GetBandwidthSavingRatioResponseDataDataSeries struct     {
  // {"en":"Timetamp\n1. When the data granularity of the query is fiveminutes, the format is yyyy-MM-dd HH:MM; Each time slice data value represents the data value in the previous time granularity range, For example yyyy-MM-dd 00:05 represents data in the range from 00:00 to 00:05.\n2.The data granularity of query is hourly, the format is yyyy-MM-dd HH. Each time slice data value represents data values in the previous time granularity range such as yyyy-MM-dd 01 that represent data from 00 to 01.\n3. the data granularity of the query is daily, the format is yyyy-MM-dd; Each time slice data value represents the data value for that day.\n4.Returns the timetamp contained in start time and end time.","zh_CN":"时间片\n1.查询的数据粒度为fiveminutes时,格式为yyyy-MM-dd HH:mm;每一个时间片数据值代表的是前一个时间粒度范围内的数据值,比如yyyy-MM-dd 00:05,代表00:00到00:05范围内的数据。\n2.查询的数据粒度为hourly时,格式为yyyy-MM-dd HH;每一个时间片数据值代表的是前一个时间粒度范围内的数据值,比如yyyy-MM-dd 01,代表00到01之间的数据。\n3.查询的数据粒度为daily时,格式为yyyy-MM-dd;每一个时间片数据值代表的该天内的数据值。\n4.返回开始时间和结束时间包含的时间片。"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {"en":"Saving of bandwidth, retain four decimals","zh_CN":"节省带宽,保留4位小数"}
  SavedBandwidthRatio *string `json:"savedBandwidthRatio,omitempty" xml:"savedBandwidthRatio,omitempty" require:"true"`
}

func (s GetBandwidthSavingRatioResponseDataDataSeries) String() string {
  return tea.Prettify(s)
}

func (s GetBandwidthSavingRatioResponseDataDataSeries) GoString() string {
  return s.String()
}

func (s *GetBandwidthSavingRatioResponseDataDataSeries) SetTimestamp(v string) *GetBandwidthSavingRatioResponseDataDataSeries {
  s.Timestamp = &v
  return s
}

func (s *GetBandwidthSavingRatioResponseDataDataSeries) SetSavedBandwidthRatio(v string) *GetBandwidthSavingRatioResponseDataDataSeries {
  s.SavedBandwidthRatio = &v
  return s
}

type GetBandwidthSavingRatioResponseHeader struct {
}

func (s GetBandwidthSavingRatioResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetBandwidthSavingRatioResponseHeader) GoString() string {
  return s.String()
}




type QueryMultipleStreamTrafficAndBandwidthUnderTheDomainRequest struct {
  // {"en":"From date:
  // 1. The time format is yyyy-MM-ddTHH:MM:ss+08:00. For example :2019-01-01T10:00:00+08:00 (10:00 on December 2, 2018 10:00:00:00 UTC+8)
  // 2. Cannot exceed current time
  // 3. The most recent six-month (183 days) data are available.", "zh_CN":"开始时间:
  // 1.格式为yyyy-MM-ddTHH:mm:ss+08:00；
  // 2.必须小于当前时间和dateTo；
  // 3.dateFrom和dateTo相差不能超过 10 分钟 ；
  // 4.只能查询最近半年内数据。
  // 5.dateFrom 和 dateTo 都不填则默认查询过去 10 分钟的数据"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty"`
  // {"en":"To time:
  // 1. The time format is yyyy-MM-ddTHH:MM:ss+08:00. For example :2019-01-01T10:00:00+08:00 (10:00 on December 2, 2018 10:00:00:00 UTC+8)
  // 2. The end time needs to be greater than the start time. If the end time is greater than the current time, take the current time.
  // 3. Date from, Date to both, the default query past 24 hours; If there is only one unsent, throw an exception", "zh_CN":"
  // 1.格式为yyyy-MM-ddTHH:mm:ss+08:00；
  // 2.必须大于dateFrom；
  // 3.如果大于当前时间，则重新赋值为当前时间；"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty"`
  // {"en":"Domain and stream group:
  // 1. the maximum number of transmissible domain names and stream name groups is 20 by default 
  // 3. stream name stream : the maximum number of stream names that can be passed under a group of domain names and stream name groups ( that is, under a single domain name ) defaults to 2000 ( which can be adjusted with technical support ).", "zh_CN":"域名和流名组:
  // 
  // 1.可传递的域名和流名组数量上限默认为20组(可联系技术支持调整)；
  // 2.域名domain:一组域名和流名组中只能传递单个域名，且域名必须传递；
  // 3.流名stream:一组域名和流名组下(即单个域名下)可传递的流名数量上限默认为2000个(可联系技术支持调整)，流名未传递时默认查询域名下所有流名，但当域名下流名数量超过限制时不可查询(报错)。"}
  DomainStream []*QueryMultipleStreamTrafficAndBandwidthUnderTheDomainRequestDomainStream `json:"domainStream,omitempty" xml:"domainStream,omitempty" require:"true" type:"Repeated"`
  // {"en":"data granularity
  // 1m: 1 min granularity, 5m: 5 min granularity
  // Default 5-minute granularity data query", "zh_CN":"数据粒度
  // 1.1m: 1分钟粒度，5m: 5分钟粒度
  // 2.默认查询5分钟粒度数据"}
  DataInterval *string `json:"dataInterval,omitempty" xml:"dataInterval,omitempty"`
  // {"en":"Data type, flow: traffic, bandwidth: bandwidth", "zh_CN":"数据类型，flow:流量，bandwidth:带宽"}
  DataType *string `json:"dataType,omitempty" xml:"dataType,omitempty"`
  // {"en":"Filter empty stream name:
  // 1. The input optional value '0' or '1';
  // 2. The data whose parameter 0 does not filter flow is null, and the data whose parameter 1 filter stream is null.
  // 3. The default value is 0;", "zh_CN":"是否过滤空流名
  // 1.入参可选值 '0' 或 '1' ；
  // 2.传参 0 不过滤流名为空的数据，传参 1 则过滤流名为空的数据 ；
  // 3.默认值为 0；"}
  FilterEmptyStream *int `json:"filterEmptyStream,omitempty" xml:"filterEmptyStream,omitempty"`
  // {"en":"Query string, multiple queries are not supported, and fuzzy matching is used by default, allowing regular expression queries", "zh_CN":"查询字符串，不支持多次查询，默认使用模糊匹配，允许正则表达式查询"}
  Keyword *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
  // {"en":"Whether to return the number of online users:
  //     1.If true is passed, the number of online users will only be returned when dataInterval=1m. Does not return when dataInterval=5min;
  //     2.If no or false is passed,the default is not to return the number of online users.", "zh_CN":"是否返回在线人数: 
  //     1.传true, 仅当dataInterval=1m时,返回在线人数,5min时不返回;
  //     2.不传或传false，则默认不返回在线人数."}
  IsReturnOnline *string `json:"isReturnOnline,omitempty" xml:"isReturnOnline,omitempty"`
}

func (s QueryMultipleStreamTrafficAndBandwidthUnderTheDomainRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryMultipleStreamTrafficAndBandwidthUnderTheDomainRequest) GoString() string {
  return s.String()
}

func (s *QueryMultipleStreamTrafficAndBandwidthUnderTheDomainRequest) SetDateFrom(v string) *QueryMultipleStreamTrafficAndBandwidthUnderTheDomainRequest {
  s.DateFrom = &v
  return s
}

func (s *QueryMultipleStreamTrafficAndBandwidthUnderTheDomainRequest) SetDateTo(v string) *QueryMultipleStreamTrafficAndBandwidthUnderTheDomainRequest {
  s.DateTo = &v
  return s
}

func (s *QueryMultipleStreamTrafficAndBandwidthUnderTheDomainRequest) SetDomainStream(v []*QueryMultipleStreamTrafficAndBandwidthUnderTheDomainRequestDomainStream) *QueryMultipleStreamTrafficAndBandwidthUnderTheDomainRequest {
  s.DomainStream = v
  return s
}

func (s *QueryMultipleStreamTrafficAndBandwidthUnderTheDomainRequest) SetDataInterval(v string) *QueryMultipleStreamTrafficAndBandwidthUnderTheDomainRequest {
  s.DataInterval = &v
  return s
}

func (s *QueryMultipleStreamTrafficAndBandwidthUnderTheDomainRequest) SetDataType(v string) *QueryMultipleStreamTrafficAndBandwidthUnderTheDomainRequest {
  s.DataType = &v
  return s
}

func (s *QueryMultipleStreamTrafficAndBandwidthUnderTheDomainRequest) SetFilterEmptyStream(v int) *QueryMultipleStreamTrafficAndBandwidthUnderTheDomainRequest {
  s.FilterEmptyStream = &v
  return s
}

func (s *QueryMultipleStreamTrafficAndBandwidthUnderTheDomainRequest) SetKeyword(v string) *QueryMultipleStreamTrafficAndBandwidthUnderTheDomainRequest {
  s.Keyword = &v
  return s
}

func (s *QueryMultipleStreamTrafficAndBandwidthUnderTheDomainRequest) SetIsReturnOnline(v string) *QueryMultipleStreamTrafficAndBandwidthUnderTheDomainRequest {
  s.IsReturnOnline = &v
  return s
}

type QueryMultipleStreamTrafficAndBandwidthUnderTheDomainRequestDomainStream struct     {
  // {"en":"Domain:
  // 1. The maximum number of deliverable domain names is 200 by default
  // 2. Automatically filter out illegal domain names (pass illegal domain names, will be filtered out, the query results only return the data of legitimate domain names)", "zh_CN":"域名:
  // 1.可传递域名数量上限默认为200个
  // 2.自动过滤掉非法域名(如传递非法域名，会被过滤掉，查询结果只返回合法域名的数据)"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {"en":"Celebrity: Publishing Point  Stream Name. Example: live/test-20180101-test where live is a publishing point and test-20180101-test is a stream name
  // No, the default queries all stream data in the specified domain name", "zh_CN":"
  // 流名:'发布点'+'流名'。例如:live/test-20180101-test,其中live是发布点,test-20180101-test是流名;
  // 不传，默认查询指定域名下的所有流的数据"}
  Stream []*string `json:"stream,omitempty" xml:"stream,omitempty" type:"Repeated"`
}

func (s QueryMultipleStreamTrafficAndBandwidthUnderTheDomainRequestDomainStream) String() string {
  return tea.Prettify(s)
}

func (s QueryMultipleStreamTrafficAndBandwidthUnderTheDomainRequestDomainStream) GoString() string {
  return s.String()
}

func (s *QueryMultipleStreamTrafficAndBandwidthUnderTheDomainRequestDomainStream) SetDomain(v string) *QueryMultipleStreamTrafficAndBandwidthUnderTheDomainRequestDomainStream {
  s.Domain = &v
  return s
}

func (s *QueryMultipleStreamTrafficAndBandwidthUnderTheDomainRequestDomainStream) SetStream(v []*string) *QueryMultipleStreamTrafficAndBandwidthUnderTheDomainRequestDomainStream {
  s.Stream = v
  return s
}

type QueryMultipleStreamTrafficAndBandwidthUnderTheDomainResponse struct {
  // {"en":"request result status code", "zh_CN":"请求结果状态码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Request result information", "zh_CN":"请求结果信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Detailed data on the results of the request", "zh_CN":"请求结果的详细数据"}
  Data []*QueryMultipleStreamTrafficAndBandwidthUnderTheDomainResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s QueryMultipleStreamTrafficAndBandwidthUnderTheDomainResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryMultipleStreamTrafficAndBandwidthUnderTheDomainResponse) GoString() string {
  return s.String()
}

func (s *QueryMultipleStreamTrafficAndBandwidthUnderTheDomainResponse) SetCode(v string) *QueryMultipleStreamTrafficAndBandwidthUnderTheDomainResponse {
  s.Code = &v
  return s
}

func (s *QueryMultipleStreamTrafficAndBandwidthUnderTheDomainResponse) SetMessage(v string) *QueryMultipleStreamTrafficAndBandwidthUnderTheDomainResponse {
  s.Message = &v
  return s
}

func (s *QueryMultipleStreamTrafficAndBandwidthUnderTheDomainResponse) SetData(v []*QueryMultipleStreamTrafficAndBandwidthUnderTheDomainResponseData) *QueryMultipleStreamTrafficAndBandwidthUnderTheDomainResponse {
  s.Data = v
  return s
}

type QueryMultipleStreamTrafficAndBandwidthUnderTheDomainResponseData struct     {
  // {"en":"domain", "zh_CN":"域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  DomainOfStreamList []*QueryMultipleStreamTrafficAndBandwidthUnderTheDomainResponseDataDomainOfStreamList `json:"domainOfStreamList,omitempty" xml:"domainOfStreamList,omitempty" require:"true" type:"Repeated"`
}

func (s QueryMultipleStreamTrafficAndBandwidthUnderTheDomainResponseData) String() string {
  return tea.Prettify(s)
}

func (s QueryMultipleStreamTrafficAndBandwidthUnderTheDomainResponseData) GoString() string {
  return s.String()
}

func (s *QueryMultipleStreamTrafficAndBandwidthUnderTheDomainResponseData) SetDomain(v string) *QueryMultipleStreamTrafficAndBandwidthUnderTheDomainResponseData {
  s.Domain = &v
  return s
}

func (s *QueryMultipleStreamTrafficAndBandwidthUnderTheDomainResponseData) SetDomainOfStreamList(v []*QueryMultipleStreamTrafficAndBandwidthUnderTheDomainResponseDataDomainOfStreamList) *QueryMultipleStreamTrafficAndBandwidthUnderTheDomainResponseData {
  s.DomainOfStreamList = v
  return s
}

type QueryMultipleStreamTrafficAndBandwidthUnderTheDomainResponseDataDomainOfStreamList struct     {
  // {"en":"stream name", "zh_CN":"流名"}
  Stream *string `json:"stream,omitempty" xml:"stream,omitempty" require:"true"`
  // {"en":"aggregate flow
  // 1. Keep 2 decimal places, in MB;
  // 2. When entering dataType is flow.
  // 3. Only data with traffic greater than0;is returned.", "zh_CN":"总流量
  // 1.保留2位小数，单位为MB;
  // 2.当入参dataType为flow时，返回。
  // 3.仅返回流量大于 0  的数据。"}
  TotalFlow *string `json:"totalFlow,omitempty" xml:"totalFlow,omitempty" require:"true"`
  // {"en":"peak bandwidth
  // 1. keep 2 decimal places in Mbps
  // 2. when entering dataType is bandwidth, return.
  // 3. only returns data with bandwidth greater than 0.", "zh_CN":"峰值带宽
  // 1.保留2位小数，单位Mbps
  // 2.当入参dataType为bandwidth时，返回。
  // 3.仅返回带宽大于 0 的数据。"}
  BandwidthPeakValue *string `json:"bandwidthPeakValue,omitempty" xml:"bandwidthPeakValue,omitempty" require:"true"`
  DetailList []*QueryMultipleStreamTrafficAndBandwidthUnderTheDomainResponseDataDomainOfStreamListDetailList `json:"detailList,omitempty" xml:"detailList,omitempty" require:"true" type:"Repeated"`
}

func (s QueryMultipleStreamTrafficAndBandwidthUnderTheDomainResponseDataDomainOfStreamList) String() string {
  return tea.Prettify(s)
}

func (s QueryMultipleStreamTrafficAndBandwidthUnderTheDomainResponseDataDomainOfStreamList) GoString() string {
  return s.String()
}

func (s *QueryMultipleStreamTrafficAndBandwidthUnderTheDomainResponseDataDomainOfStreamList) SetStream(v string) *QueryMultipleStreamTrafficAndBandwidthUnderTheDomainResponseDataDomainOfStreamList {
  s.Stream = &v
  return s
}

func (s *QueryMultipleStreamTrafficAndBandwidthUnderTheDomainResponseDataDomainOfStreamList) SetTotalFlow(v string) *QueryMultipleStreamTrafficAndBandwidthUnderTheDomainResponseDataDomainOfStreamList {
  s.TotalFlow = &v
  return s
}

func (s *QueryMultipleStreamTrafficAndBandwidthUnderTheDomainResponseDataDomainOfStreamList) SetBandwidthPeakValue(v string) *QueryMultipleStreamTrafficAndBandwidthUnderTheDomainResponseDataDomainOfStreamList {
  s.BandwidthPeakValue = &v
  return s
}

func (s *QueryMultipleStreamTrafficAndBandwidthUnderTheDomainResponseDataDomainOfStreamList) SetDetailList(v []*QueryMultipleStreamTrafficAndBandwidthUnderTheDomainResponseDataDomainOfStreamListDetailList) *QueryMultipleStreamTrafficAndBandwidthUnderTheDomainResponseDataDomainOfStreamList {
  s.DetailList = v
  return s
}

type QueryMultipleStreamTrafficAndBandwidthUnderTheDomainResponseDataDomainOfStreamListDetailList struct     {
  // {"en":"time
  // 1. Each time slice data value represents the data value in the previous time granularity range. 
  // 2. The data granularity of the query is 5m, the format is yyyy-MM-dd HH:MM; Each time slice data value represents the data value in the previous time granularity range.", "zh_CN":"时间
  // 1.查询的数据粒度为1m时，格式为yyyy-MM-dd HH:mm；每一个时间片数据值代表的是前一个时间粒度范围内的数据值。一天开始的时间片是yyyy-MM-dd 00:01，最后一个时间片是(yyyy-MM-dd+1) 00:00。
  // 2.查询的数据粒度为5m时，格式为yyyy-MM-dd HH:mm；每一个时间片数据值代表的是前一个时间粒度范围内的数据值。一天开始的时间片是yyyy-MM-dd 00:05，最后一个时间片是(yyyy-MM-dd+1) 00:00。"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {"en":"Data at the point in time:
  // 1. keep 2 decimal places;
  // 2. When the dataType is flow, value is flow in MB;
  // 3. value is bandwidth value in Mbps when dataType is bandwidth.", "zh_CN":"时间点的数据
  // 1.保留2位小数；
  // 2.当入参dataType为flow时，value值为流量，单位为MB；
  // 3.当入参dataType为bandwidth时，value值为带宽值，单位为Mbps；"}
  Value *string `json:"value,omitempty" xml:"value,omitempty" require:"true"`
  // {"en":"Number of online users at a time point.", "zh_CN":"时间点的在线人数"}
  OnlineUser *int `json:"onlineUser,omitempty" xml:"onlineUser,omitempty" require:"true"`
}

func (s QueryMultipleStreamTrafficAndBandwidthUnderTheDomainResponseDataDomainOfStreamListDetailList) String() string {
  return tea.Prettify(s)
}

func (s QueryMultipleStreamTrafficAndBandwidthUnderTheDomainResponseDataDomainOfStreamListDetailList) GoString() string {
  return s.String()
}

func (s *QueryMultipleStreamTrafficAndBandwidthUnderTheDomainResponseDataDomainOfStreamListDetailList) SetTimestamp(v string) *QueryMultipleStreamTrafficAndBandwidthUnderTheDomainResponseDataDomainOfStreamListDetailList {
  s.Timestamp = &v
  return s
}

func (s *QueryMultipleStreamTrafficAndBandwidthUnderTheDomainResponseDataDomainOfStreamListDetailList) SetValue(v string) *QueryMultipleStreamTrafficAndBandwidthUnderTheDomainResponseDataDomainOfStreamListDetailList {
  s.Value = &v
  return s
}

func (s *QueryMultipleStreamTrafficAndBandwidthUnderTheDomainResponseDataDomainOfStreamListDetailList) SetOnlineUser(v int) *QueryMultipleStreamTrafficAndBandwidthUnderTheDomainResponseDataDomainOfStreamListDetailList {
  s.OnlineUser = &v
  return s
}

type QueryMultipleStreamTrafficAndBandwidthUnderTheDomainPaths struct {
}

func (s QueryMultipleStreamTrafficAndBandwidthUnderTheDomainPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryMultipleStreamTrafficAndBandwidthUnderTheDomainPaths) GoString() string {
  return s.String()
}

type QueryMultipleStreamTrafficAndBandwidthUnderTheDomainParameters struct {
}

func (s QueryMultipleStreamTrafficAndBandwidthUnderTheDomainParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryMultipleStreamTrafficAndBandwidthUnderTheDomainParameters) GoString() string {
  return s.String()
}

type QueryMultipleStreamTrafficAndBandwidthUnderTheDomainRequestHeader struct {
}

func (s QueryMultipleStreamTrafficAndBandwidthUnderTheDomainRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryMultipleStreamTrafficAndBandwidthUnderTheDomainRequestHeader) GoString() string {
  return s.String()
}

type QueryMultipleStreamTrafficAndBandwidthUnderTheDomainResponseHeader struct {
}

func (s QueryMultipleStreamTrafficAndBandwidthUnderTheDomainResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryMultipleStreamTrafficAndBandwidthUnderTheDomainResponseHeader) GoString() string {
  return s.String()
}




type QueryTrafficForSelectedBillingAreaRequest struct {
  // {"en":"Start time
  // 
  // 1.The format is yyyy-MM-ddTHH:mm:ss+08:00;
  // 2.And smaller than the current time and dateTo;
  // 3.Period between dataFrom and dateTo cannot be longer than 31 days;4.You can only query data for the last 2 years(technical support can be contacted to adjust).", "zh_CN":"开始时间
  // 1.格式为yyyy-MM-ddTHH:mm:ss+08:00;
  // 2.并且小于当前时间和dateTo;
  // 3.dateFrom和dateTo相差不能超过31天(可联系技术支持调整);4.只能查询最近2年内数据。"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty" require:"true"`
  // {"en":"End time
  // 1.The format is yyyy-MM-ddTHH:mm:ss+08:00;
  // 2.Must be greater than dateFrom;
  // 3.If it's greater than the current time, then the current time is assigned as the value;", "zh_CN":"结束时间
  // 1.格式为yyyy-MM-ddTHH:mm:ss+08:00;
  // 2.必须大于dateFrom;
  // 3.如果大于当前时间,则重新赋值为当前时间;"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty" require:"true"`
  // {"en":"Domain list, and all domains are queried if this field is not specified;", "zh_CN":"域名列表,不传递则查询全部域名;"}
  Domain []*string `json:"domain,omitempty" xml:"domain,omitempty" type:"Repeated"`
  // {"en":"Data granularity
  // 
  // 1.Options: 5m(5 minutes), 1h(1 hour) and 1d(1 day);
  // 2.Default value of 1d is used if the field is not specified;
  // 3.If 5m is specified as the value, then data is returned in actual configured granularity when there is specific configuration to data collecting granularity for the customer.", "zh_CN":"数据粒度
  // 1.可选值为:5m(5分钟)、1h(1小时)、1d(1天);
  // 2.不传时默认为1d;
  // 3.传递5m时,若客户数据采集粒度有特殊配置将按实际配置粒度返回。"}
  DataInterval *string `json:"dataInterval,omitempty" xml:"dataInterval,omitempty"`
  // {"en":"Queried area
  // 1.Use  English semicolon to separate two areas;
  // 2.Options: cn, nc, ov, apac, am, euna, emea, sa, af, au, hk, tw 
  //  ...;", "zh_CN":"查询区域
  // 1.多个区域使用英文分号分隔
  // 2.可选值为:cn、nc、ov、apac、am、euna、emea、sa、af、au、hk、tw等;"}
  AreaCode *string `json:"areaCode,omitempty" xml:"areaCode,omitempty"`
}

func (s QueryTrafficForSelectedBillingAreaRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryTrafficForSelectedBillingAreaRequest) GoString() string {
  return s.String()
}

func (s *QueryTrafficForSelectedBillingAreaRequest) SetDateFrom(v string) *QueryTrafficForSelectedBillingAreaRequest {
  s.DateFrom = &v
  return s
}

func (s *QueryTrafficForSelectedBillingAreaRequest) SetDateTo(v string) *QueryTrafficForSelectedBillingAreaRequest {
  s.DateTo = &v
  return s
}

func (s *QueryTrafficForSelectedBillingAreaRequest) SetDomain(v []*string) *QueryTrafficForSelectedBillingAreaRequest {
  s.Domain = v
  return s
}

func (s *QueryTrafficForSelectedBillingAreaRequest) SetDataInterval(v string) *QueryTrafficForSelectedBillingAreaRequest {
  s.DataInterval = &v
  return s
}

func (s *QueryTrafficForSelectedBillingAreaRequest) SetAreaCode(v string) *QueryTrafficForSelectedBillingAreaRequest {
  s.AreaCode = &v
  return s
}

type QueryTrafficForSelectedBillingAreaResponse struct {
  // {"en":"Total traffic, 2 decimal places, in MB", "zh_CN":"总流量,保留2位小数,单位为MB"}
  FlowSummary *int `json:"flow-summary,omitempty" xml:"flow-summary,omitempty" require:"true"`
  FlowData []*QueryTrafficForSelectedBillingAreaResponseFlowData `json:"flow-data,omitempty" xml:"flow-data,omitempty" require:"true" type:"Repeated"`
}

func (s QueryTrafficForSelectedBillingAreaResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryTrafficForSelectedBillingAreaResponse) GoString() string {
  return s.String()
}

func (s *QueryTrafficForSelectedBillingAreaResponse) SetFlowSummary(v int) *QueryTrafficForSelectedBillingAreaResponse {
  s.FlowSummary = &v
  return s
}

func (s *QueryTrafficForSelectedBillingAreaResponse) SetFlowData(v []*QueryTrafficForSelectedBillingAreaResponseFlowData) *QueryTrafficForSelectedBillingAreaResponse {
  s.FlowData = v
  return s
}

type QueryTrafficForSelectedBillingAreaResponseFlowData struct     {
  // {"en":"time
  // 1. When the data granularity of the query is fiveminutes, the format is yyyy-MM-dd HH: MM; Each time slice data value represents the data value in the previous time granularity range. The time slice at the beginning of the day is yyyy-MM-dd 00:05, and the last time slice is yyyy-MM-dd 24:00.
  // 2. When the data granularity of query is hourly, the format is yyyy-MM-dd HH. Each time slice data value represents the data value in the previous time granularity range. The time slice at the beginning of the day is yyyy-MM-dd 01, and the last time slice is yyyy-MM-dd 24.
  // 3. when the data granularity of query is daily, the format is yyyy-MM-dd; The value of data for each time slice represents the value of the data within that day;
  // 4. return the time slice contained in the start and end times.", "zh_CN":"时间
  // 1.查询的数据粒度为fiveminutes时,格式为yyyy-MM-dd HH:mm;每一个时间片数据值代表的是前一个时间粒度范围内的数据值。一天开始的时间片是yyyy-MM-dd 00:05,最后一个时间片是yyyy-MM-dd 24:00。
  // 2.查询的数据粒度为hourly时,格式为yyyy-MM-dd HH;每一个时间片数据值代表的是前一个时间粒度范围内的数据值。一天开始的时间片是yyyy-MM-dd 01,最后一个时间片是yyyy-MM-dd 24。
  // 3.查询的数据粒度为daily时,格式为yyyy-MM-dd;每一个时间片数据值代表的该天内的数据值;
  // 4.返回开始时间和结束时间包含的时间片。"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {"en":"Traffic, 2 decimal places, in MB", "zh_CN":"流量,保留2位小数,单位为MB"}
  Flow *int `json:"flow,omitempty" xml:"flow,omitempty" require:"true"`
}

func (s QueryTrafficForSelectedBillingAreaResponseFlowData) String() string {
  return tea.Prettify(s)
}

func (s QueryTrafficForSelectedBillingAreaResponseFlowData) GoString() string {
  return s.String()
}

func (s *QueryTrafficForSelectedBillingAreaResponseFlowData) SetTimestamp(v string) *QueryTrafficForSelectedBillingAreaResponseFlowData {
  s.Timestamp = &v
  return s
}

func (s *QueryTrafficForSelectedBillingAreaResponseFlowData) SetFlow(v int) *QueryTrafficForSelectedBillingAreaResponseFlowData {
  s.Flow = &v
  return s
}

type QueryTrafficForSelectedBillingAreaPaths struct {
}

func (s QueryTrafficForSelectedBillingAreaPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryTrafficForSelectedBillingAreaPaths) GoString() string {
  return s.String()
}

type QueryTrafficForSelectedBillingAreaParameters struct {
}

func (s QueryTrafficForSelectedBillingAreaParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryTrafficForSelectedBillingAreaParameters) GoString() string {
  return s.String()
}

type QueryTrafficForSelectedBillingAreaRequestHeader struct {
}

func (s QueryTrafficForSelectedBillingAreaRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryTrafficForSelectedBillingAreaRequestHeader) GoString() string {
  return s.String()
}

type QueryTrafficForSelectedBillingAreaResponseHeader struct {
}

func (s QueryTrafficForSelectedBillingAreaResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryTrafficForSelectedBillingAreaResponseHeader) GoString() string {
  return s.String()
}




