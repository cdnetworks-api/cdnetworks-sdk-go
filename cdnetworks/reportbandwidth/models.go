package reportbandwidth

// This file is auto-generated, don't edit it. Thanks.
import (
  "github.com/alibabacloud-go/tea/tea"
)

type PerzoneBillingRequest struct {
  // {"en":"cust_en_name of sub-client.
  // When a merged-account wants to  view the information of the subclient,the cust_en_name is required.", "zh_CN":"合并账号下的某个客户的英文名，当合并账号要查看子客户的信息时，必须填写子客户的英文名"}
  Cust *string `json:"cust,omitempty" xml:"cust,omitempty"`
  // {"en":"Specifies the query date:
  // 1.With format yyyy-mm-dd.
  // 2.If not Specifies,it means today as default.", "zh_CN":"查询的日期，日期格式为yyyy-mm-dd,不选或者为空时默认为当天；"}
  Date *string `json:"date,omitempty" xml:"date,omitempty"`
  // {"en":"1.Must work with 'enddate' and they specify the query date scope. 2.With format yyyy-mm-dd. 3.If there is a 'date' parameter,this parameter will be omitted.", "zh_CN":"查询的起始日期，日期格式为yyyy-mm-dd；此参数需与enddate参数配合,若存在date参数,则该参数无效"}
  StartDate *string `json:"startDate,omitempty" xml:"startDate,omitempty"`
  // {"en":"1.Must work with 'startdate' and they specify the query date scope. 2.With format yyyy-mm-dd 3.If there is a 'date' parameter,this parameter will be omitted.", "zh_CN":"查询的结束日期,日期格式为yyyy-mm-dd；此参数需与startdate参数配合,若存在date参数,则该参数无效。"}
  EndDate *string `json:"endDate,omitempty" xml:"endDate,omitempty"`
  // {"en":"In Greenwich time zone, the value GMT+09:00 means East Nine, GMT-09:00 means West Nine, and the default is East Eight.", "zh_CN":"格林尼治时区，参数格式 GMT+09:00 表示东九区，GMT-09:00 表示西9区，不传则默认为本地时区（东八区）"}
  Timezone *string `json:"timezone,omitempty" xml:"timezone,omitempty"`
  // {"en":"domains that been queried:
  // 1.If there are multiple inputs,use  ';' as separator.
  // 2.If not specified, it means all the domains of the account .", "zh_CN":"查询的频道，多个频道值请用英文分号“;”，不选或者为空时默认为所查询客户的所有频道"}
  Channel *string `json:"channel,omitempty" xml:"channel,omitempty"`
  // {"en":"acceleration type.
  // 1.If there are multiple inputs,use ';' as separator.
  // 2.If not specified or specified as 'all', it means all the accetypes.", "zh_CN":"加速类型参数，如accetype=web。多个请用英文分号“;”分隔开，不填或值为all表示所有类型"}
  Accetype *string `json:"accetype,omitempty" xml:"accetype,omitempty"`
  // {"en":"Whether it is a DWA product, 1: yes; 0: no, if not fill in, the default is 0", "zh_CN":"是否DWA产品，1:是;0:否,不填默认0"}
  Dwa *string `json:"dwa,omitempty" xml:"dwa,omitempty"`
  // {"en":"Data type (1: bandwidth|2: traffic) default 1", "zh_CN":"数据类型（1:带宽|2:流量）默认1"}
  Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s PerzoneBillingRequest) String() string {
  return tea.Prettify(s)
}

func (s PerzoneBillingRequest) GoString() string {
  return s.String()
}

func (s *PerzoneBillingRequest) SetCust(v string) *PerzoneBillingRequest {
  s.Cust = &v
  return s
}

func (s *PerzoneBillingRequest) SetDate(v string) *PerzoneBillingRequest {
  s.Date = &v
  return s
}

func (s *PerzoneBillingRequest) SetStartDate(v string) *PerzoneBillingRequest {
  s.StartDate = &v
  return s
}

func (s *PerzoneBillingRequest) SetEndDate(v string) *PerzoneBillingRequest {
  s.EndDate = &v
  return s
}

func (s *PerzoneBillingRequest) SetTimezone(v string) *PerzoneBillingRequest {
  s.Timezone = &v
  return s
}

func (s *PerzoneBillingRequest) SetChannel(v string) *PerzoneBillingRequest {
  s.Channel = &v
  return s
}

func (s *PerzoneBillingRequest) SetAccetype(v string) *PerzoneBillingRequest {
  s.Accetype = &v
  return s
}

func (s *PerzoneBillingRequest) SetDwa(v string) *PerzoneBillingRequest {
  s.Dwa = &v
  return s
}

func (s *PerzoneBillingRequest) SetType(v string) *PerzoneBillingRequest {
  s.Type = &v
  return s
}

type PerzoneBillingResponse struct {
  // {'en':'provider', 'zh_CN':'结果'}
  Provider *PerzoneBillingResponseProvider `json:"provider,omitempty" xml:"provider,omitempty" require:"true" type:"Struct"`
}

func (s PerzoneBillingResponse) String() string {
  return tea.Prettify(s)
}

func (s PerzoneBillingResponse) GoString() string {
  return s.String()
}

func (s *PerzoneBillingResponse) SetProvider(v *PerzoneBillingResponseProvider) *PerzoneBillingResponse {
  s.Provider = v
  return s
}

type PerzoneBillingResponseProvider struct {
  // {'en':'tenant', 'zh_CN':'租户'}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {'en':'type', 'zh_CN':'接口类型'}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
  // {'en':'resultType', 'zh_CN':'统计类型'}
  ResultType *string `json:"resultType,omitempty" xml:"resultType,omitempty" require:"true"`
  // {'en':'data', 'zh_CN':'perzone带宽数据'}
  Date *PerzoneBillingResponseProviderDate `json:"date,omitempty" xml:"date,omitempty" require:"true" type:"Struct"`
}

func (s PerzoneBillingResponseProvider) String() string {
  return tea.Prettify(s)
}

func (s PerzoneBillingResponseProvider) GoString() string {
  return s.String()
}

func (s *PerzoneBillingResponseProvider) SetName(v string) *PerzoneBillingResponseProvider {
  s.Name = &v
  return s
}

func (s *PerzoneBillingResponseProvider) SetType(v string) *PerzoneBillingResponseProvider {
  s.Type = &v
  return s
}

func (s *PerzoneBillingResponseProvider) SetResultType(v string) *PerzoneBillingResponseProvider {
  s.ResultType = &v
  return s
}

func (s *PerzoneBillingResponseProvider) SetDate(v *PerzoneBillingResponseProviderDate) *PerzoneBillingResponseProvider {
  s.Date = v
  return s
}

type PerzoneBillingResponseProviderDate struct {
  // {'en':'startdate', 'zh_CN':'开始时间'}
  StartDate *string `json:"startDate,omitempty" xml:"startDate,omitempty" require:"true"`
  // {'en':'enddate', 'zh_CN':'结束时间'}
  EndDate *string `json:"endDate,omitempty" xml:"endDate,omitempty" require:"true"`
  // {'en':'type', 'zh_CN':'数据类型1:带宽|2:流量'}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
  // {'en':'channel', 'zh_CN':'频道'}
  AreaMap *PerzoneBillingResponseProviderDateAreaMap `json:"areaMap,omitempty" xml:"areaMap,omitempty" require:"true" type:"Struct"`
}

func (s PerzoneBillingResponseProviderDate) String() string {
  return tea.Prettify(s)
}

func (s PerzoneBillingResponseProviderDate) GoString() string {
  return s.String()
}

func (s *PerzoneBillingResponseProviderDate) SetStartDate(v string) *PerzoneBillingResponseProviderDate {
  s.StartDate = &v
  return s
}

func (s *PerzoneBillingResponseProviderDate) SetEndDate(v string) *PerzoneBillingResponseProviderDate {
  s.EndDate = &v
  return s
}

func (s *PerzoneBillingResponseProviderDate) SetType(v string) *PerzoneBillingResponseProviderDate {
  s.Type = &v
  return s
}

func (s *PerzoneBillingResponseProviderDate) SetAreaMap(v *PerzoneBillingResponseProviderDateAreaMap) *PerzoneBillingResponseProviderDate {
  s.AreaMap = v
  return s
}

type PerzoneBillingResponseProviderDateAreaMap struct {
  // {'en':'perzone', 'zh_CN':'perzone数据'}
  Perzone *PerzoneBillingResponseProviderDateAreaMapPerzone `json:"perzone,omitempty" xml:"perzone,omitempty" require:"true" type:"Struct"`
}

func (s PerzoneBillingResponseProviderDateAreaMap) String() string {
  return tea.Prettify(s)
}

func (s PerzoneBillingResponseProviderDateAreaMap) GoString() string {
  return s.String()
}

func (s *PerzoneBillingResponseProviderDateAreaMap) SetPerzone(v *PerzoneBillingResponseProviderDateAreaMapPerzone) *PerzoneBillingResponseProviderDateAreaMap {
  s.Perzone = v
  return s
}

type PerzoneBillingResponseProviderDateAreaMapPerzone struct {
  // {'en':'name', 'zh_CN':'perzone区域'}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {'en':'peakValue', 'zh_CN':'峰值'}
  PeakValue *string `json:"peakValue,omitempty" xml:"peakValue,omitempty" require:"true"`
  // {'en':'peakTime', 'zh_CN':'峰值时间点'}
  PeakTime *string `json:"peakTime,omitempty" xml:"peakTime,omitempty" require:"true"`
  // {'en':'totalFlow', 'zh_CN':'总流量'}
  TotalFlow *string `json:"totalFlow,omitempty" xml:"totalFlow,omitempty" require:"true"`
  // {'en':'bandwidth', 'zh_CN':'带宽数据'}
  Detail []*PerzoneBillingResponseProviderDateAreaMapPerzoneDetail `json:"detail,omitempty" xml:"detail,omitempty" require:"true" type:"Repeated"`
}

func (s PerzoneBillingResponseProviderDateAreaMapPerzone) String() string {
  return tea.Prettify(s)
}

func (s PerzoneBillingResponseProviderDateAreaMapPerzone) GoString() string {
  return s.String()
}

func (s *PerzoneBillingResponseProviderDateAreaMapPerzone) SetName(v string) *PerzoneBillingResponseProviderDateAreaMapPerzone {
  s.Name = &v
  return s
}

func (s *PerzoneBillingResponseProviderDateAreaMapPerzone) SetPeakValue(v string) *PerzoneBillingResponseProviderDateAreaMapPerzone {
  s.PeakValue = &v
  return s
}

func (s *PerzoneBillingResponseProviderDateAreaMapPerzone) SetPeakTime(v string) *PerzoneBillingResponseProviderDateAreaMapPerzone {
  s.PeakTime = &v
  return s
}

func (s *PerzoneBillingResponseProviderDateAreaMapPerzone) SetTotalFlow(v string) *PerzoneBillingResponseProviderDateAreaMapPerzone {
  s.TotalFlow = &v
  return s
}

func (s *PerzoneBillingResponseProviderDateAreaMapPerzone) SetDetail(v []*PerzoneBillingResponseProviderDateAreaMapPerzoneDetail) *PerzoneBillingResponseProviderDateAreaMapPerzone {
  s.Detail = v
  return s
}

type PerzoneBillingResponseProviderDateAreaMapPerzoneDetail struct     {
  // {'en':'timestamp', 'zh_CN':'时间点'}
  Time *string `json:"time,omitempty" xml:"time,omitempty" require:"true"`
  // {'en':'bandwidth', 'zh_CN':'带宽'}
  Text *string `json:"text,omitempty" xml:"text,omitempty" require:"true"`
}

func (s PerzoneBillingResponseProviderDateAreaMapPerzoneDetail) String() string {
  return tea.Prettify(s)
}

func (s PerzoneBillingResponseProviderDateAreaMapPerzoneDetail) GoString() string {
  return s.String()
}

func (s *PerzoneBillingResponseProviderDateAreaMapPerzoneDetail) SetTime(v string) *PerzoneBillingResponseProviderDateAreaMapPerzoneDetail {
  s.Time = &v
  return s
}

func (s *PerzoneBillingResponseProviderDateAreaMapPerzoneDetail) SetText(v string) *PerzoneBillingResponseProviderDateAreaMapPerzoneDetail {
  s.Text = &v
  return s
}

type PerzoneBillingPaths struct {
}

func (s PerzoneBillingPaths) String() string {
  return tea.Prettify(s)
}

func (s PerzoneBillingPaths) GoString() string {
  return s.String()
}

type PerzoneBillingParameters struct {
}

func (s PerzoneBillingParameters) String() string {
  return tea.Prettify(s)
}

func (s PerzoneBillingParameters) GoString() string {
  return s.String()
}

type PerzoneBillingRequestHeader struct {
}

func (s PerzoneBillingRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s PerzoneBillingRequestHeader) GoString() string {
  return s.String()
}

type PerzoneBillingResponseHeader struct {
}

func (s PerzoneBillingResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s PerzoneBillingResponseHeader) GoString() string {
  return s.String()
}




type QueryBandwidthofOriginminutelyRequest struct {
  // {"en":"Start Time:
  // 1. The time format is yyyy-MM-ddTHH:MM:ss+08:00, for example, 2019-01-01T10:00:00+08:00 (Beijing time on January 1, 2019 at 10:00 am to 0 seconds);
  // 2. No more than the current time;
  // 3. Up to 6 months (183 days) of data are available.", "zh_CN":"开始时间：
  // 1.时间格式为yyyy-MM-ddTHH:mm:ss+08:00，例如，2019-01-01T10:00:00+08:00（为北京时间2019年1月1日10点0分0秒）；
  // 2.不能大于当前时间；
  // 3.最多可获取最近半年（183天）的数据。"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty"`
  // {"en":"End Time:
  // 1. Time Format 2019-01-02T10:00:00+08:00;
  // 2. the end time should be greater than the start time. if the end time is greater than the current time, take the current time;
  // 3. dateFrom, dateTo, both are not sent, default query past 24 hours; If only one is not sent, throw exception;
  // 4. Allow maximum query interval: 7 days, i.e., 7 days between Dategroup and dateTo. Do not exceed 7 days.", "zh_CN":"结束时间：
  // 1.时间格式2019-01-02T10:00:00+08:00；
  // 2.结束时间需大于开始时间，结束时间如果大于当前时间，取当前时间；
  // 3.dateFrom，dateTo二者都未传，默认查询过去的24小时；如仅有一个未传，抛异常；
  // 4.允许查询最大时间间隔：7天，即dateFrom和dateTo相差不能超过7天，不支持调整"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty"`
  // {"en":"Data granularity:
  // 1. Support for 1m (1 minute), 5m (5 minutes), 1h (1 hour);
  // 2. do not pass default 5m.
  // Data granularity, default to 5m", "zh_CN":"数据粒度：
  // 1.支持1m（1分钟）、5m（5分钟）、1h（1小时）；
  // 2.不传默认5m。
  // 
  // 数据粒度，默认为5m"}
  DataInterval *string `json:"dataInterval,omitempty" xml:"dataInterval,omitempty"`
  // {"en":"Domain name:
  // 1. the maximum number of transitive domain names is 50 by default;
  // 2. Automatically filter out illegal domain names (e.g. passing illegal domain names will be filtered out, and the search results will only return the legal domain name data)
  // 3. If left blank, all domain names will be obtained. If the total number of domain names exceeds the upper limit, an error will be reported.", "zh_CN":"域名：
  // 1.可传递域名数量上限默认为50（可联系技术支持调整）；
  // 2.自动过滤掉非法域名（如传递非法域名，会被过滤掉，查询结果只返回合法域名的数据）。
  // 3. 若未填写默认查询全部域名，全部域名超出域名上限报错"}
  Domain []*string `json:"domain,omitempty" xml:"domain,omitempty" type:"Repeated"`
  // {"en":"Acceleration zone:
  // 1. do not pass the default query to all areas;
  // 2. currently only the leaflet area is supported externally;
  // 3. optional values: cn, apac, am, emea.", "zh_CN":"加速区域：
  // 1.不传默认查询全部区域；
  // 2.目前对外只支持传单区域；
  // 3.可选值：cn、apac、am、emea。"}
  Region *string `json:"region,omitempty" xml:"region,omitempty"`
  // {"en":"grouped dimension:
  // 1. the optional value is domain;
  // 2. If incoming data is shown in accordance with the dimension.", "zh_CN":"分组维度：
  // 1.可选值为domain；
  // 2.有传入则按照该维度展示明细数据。"}
  GroupBy *string `json:"groupBy,omitempty" xml:"groupBy,omitempty"`
}

func (s QueryBandwidthofOriginminutelyRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryBandwidthofOriginminutelyRequest) GoString() string {
  return s.String()
}

func (s *QueryBandwidthofOriginminutelyRequest) SetDateFrom(v string) *QueryBandwidthofOriginminutelyRequest {
  s.DateFrom = &v
  return s
}

func (s *QueryBandwidthofOriginminutelyRequest) SetDateTo(v string) *QueryBandwidthofOriginminutelyRequest {
  s.DateTo = &v
  return s
}

func (s *QueryBandwidthofOriginminutelyRequest) SetDataInterval(v string) *QueryBandwidthofOriginminutelyRequest {
  s.DataInterval = &v
  return s
}

func (s *QueryBandwidthofOriginminutelyRequest) SetDomain(v []*string) *QueryBandwidthofOriginminutelyRequest {
  s.Domain = v
  return s
}

func (s *QueryBandwidthofOriginminutelyRequest) SetRegion(v string) *QueryBandwidthofOriginminutelyRequest {
  s.Region = &v
  return s
}

func (s *QueryBandwidthofOriginminutelyRequest) SetGroupBy(v string) *QueryBandwidthofOriginminutelyRequest {
  s.GroupBy = &v
  return s
}

type QueryBandwidthofOriginminutelyResponse struct {
  // {"en":"Request result status code", "zh_CN":"请求结果状态码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Request result information", "zh_CN":"请求结果信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Detailed data on request results", "zh_CN":"请求结果的详细数据"}
  Data []*QueryBandwidthofOriginminutelyResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s QueryBandwidthofOriginminutelyResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryBandwidthofOriginminutelyResponse) GoString() string {
  return s.String()
}

func (s *QueryBandwidthofOriginminutelyResponse) SetCode(v string) *QueryBandwidthofOriginminutelyResponse {
  s.Code = &v
  return s
}

func (s *QueryBandwidthofOriginminutelyResponse) SetMessage(v string) *QueryBandwidthofOriginminutelyResponse {
  s.Message = &v
  return s
}

func (s *QueryBandwidthofOriginminutelyResponse) SetData(v []*QueryBandwidthofOriginminutelyResponseData) *QueryBandwidthofOriginminutelyResponse {
  s.Data = v
  return s
}

type QueryBandwidthofOriginminutelyResponseData struct     {
  // {"en":"Domain", "zh_CN":"域名信息"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {"en":"Peak bandwidth Mbps, example (931556.21 Mbps)", "zh_CN":"峰值带宽 Mbps，示例 （931556.21 Mbps）"}
  PeakValue *string `json:"peakValue,omitempty" xml:"peakValue,omitempty" require:"true"`
  // {"en":"Peak Time, Example (2019-02-13 18:01)", "zh_CN":"峰值时间，示例(2019-02-13 18:01)"}
  PeakTime *string `json:"peakTime,omitempty" xml:"peakTime,omitempty" require:"true"`
  // {"en":"Total return flow, example ( 74099.92 MB )", "zh_CN":"回源总流量，示例 ( 74099.92 MB )"}
  Total *string `json:"total,omitempty" xml:"total,omitempty" require:"true"`
  OriginBandwidthData []*QueryBandwidthofOriginminutelyResponseDataOriginBandwidthData `json:"originBandwidthData,omitempty" xml:"originBandwidthData,omitempty" require:"true" type:"Repeated"`
}

func (s QueryBandwidthofOriginminutelyResponseData) String() string {
  return tea.Prettify(s)
}

func (s QueryBandwidthofOriginminutelyResponseData) GoString() string {
  return s.String()
}

func (s *QueryBandwidthofOriginminutelyResponseData) SetDomain(v string) *QueryBandwidthofOriginminutelyResponseData {
  s.Domain = &v
  return s
}

func (s *QueryBandwidthofOriginminutelyResponseData) SetPeakValue(v string) *QueryBandwidthofOriginminutelyResponseData {
  s.PeakValue = &v
  return s
}

func (s *QueryBandwidthofOriginminutelyResponseData) SetPeakTime(v string) *QueryBandwidthofOriginminutelyResponseData {
  s.PeakTime = &v
  return s
}

func (s *QueryBandwidthofOriginminutelyResponseData) SetTotal(v string) *QueryBandwidthofOriginminutelyResponseData {
  s.Total = &v
  return s
}

func (s *QueryBandwidthofOriginminutelyResponseData) SetOriginBandwidthData(v []*QueryBandwidthofOriginminutelyResponseDataOriginBandwidthData) *QueryBandwidthofOriginminutelyResponseData {
  s.OriginBandwidthData = v
  return s
}

type QueryBandwidthofOriginminutelyResponseDataOriginBandwidthData struct     {
  // {"en":"The granularity of data is 1 minute, and the format is yyyy-MM-dd HH:MM.", "zh_CN":"数据粒度为1分钟，格式为yyyy-MM-dd HH:mm。"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {"en":"Return the source bandwidth value, in Mbps, 2 decimal places reserved.", "zh_CN":"回源带宽值，单位Mbps，保留2位小数。"}
  Value *string `json:"value,omitempty" xml:"value,omitempty" require:"true"`
}

func (s QueryBandwidthofOriginminutelyResponseDataOriginBandwidthData) String() string {
  return tea.Prettify(s)
}

func (s QueryBandwidthofOriginminutelyResponseDataOriginBandwidthData) GoString() string {
  return s.String()
}

func (s *QueryBandwidthofOriginminutelyResponseDataOriginBandwidthData) SetTimestamp(v string) *QueryBandwidthofOriginminutelyResponseDataOriginBandwidthData {
  s.Timestamp = &v
  return s
}

func (s *QueryBandwidthofOriginminutelyResponseDataOriginBandwidthData) SetValue(v string) *QueryBandwidthofOriginminutelyResponseDataOriginBandwidthData {
  s.Value = &v
  return s
}

type QueryBandwidthofOriginminutelyPaths struct {
}

func (s QueryBandwidthofOriginminutelyPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryBandwidthofOriginminutelyPaths) GoString() string {
  return s.String()
}

type QueryBandwidthofOriginminutelyParameters struct {
}

func (s QueryBandwidthofOriginminutelyParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryBandwidthofOriginminutelyParameters) GoString() string {
  return s.String()
}

type QueryBandwidthofOriginminutelyRequestHeader struct {
}

func (s QueryBandwidthofOriginminutelyRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryBandwidthofOriginminutelyRequestHeader) GoString() string {
  return s.String()
}

type QueryBandwidthofOriginminutelyResponseHeader struct {
}

func (s QueryBandwidthofOriginminutelyResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryBandwidthofOriginminutelyResponseHeader) GoString() string {
  return s.String()
}




type BandwidthChannelRequest struct {
  // {"en":"cust_en_name of sub-client.
  // When a merged-account wants to  view the information of the subclient,the cust_en_name is required.", "zh_CN":"合并账号下的某个客户的英文名，当合并账号要查看子客户的信息时，必须填写子客户的英文名"}
  Cust *string `json:"cust,omitempty" xml:"cust,omitempty"`
  // {"en":"Specifies the query date:
  // 1.With format yyyy-mm-dd.
  // 2.If not Specifies,it means today as default.", "zh_CN":"查询的日期，日期格式为yyyy-mm-dd,不选或者为空时默认为当天；"}
  Date *string `json:"date,omitempty" xml:"date,omitempty"`
  // {"en":"1.Must work with 'enddate' and they  specify the query date scope. 
  // 2.With format yyyy-mm-dd hh:MM.If 'hh:MM' not specified,it means '00:01'.
  // 3.If there is a 'date' parameter,this parameter will be omitted.", "zh_CN":"查询的起始日期,精确到分钟,日期格式为yyyy-mm-dd hh:MM若没有输入时、分，则时分默认为00:01；此参数需与enddate参数配合,若存在date参数,则该参数无效"}
  Startdate *string `json:"startdate,omitempty" xml:"startdate,omitempty"`
  // {"en":"1.Must work with 'startdate' and they  specify the query date scope. 
  // 2.With format yyyy-mm-dd hh:MM.If 'hh:MM' not specified,it means '24:00'.
  // 3.If there is a 'date' parameter,this parameter will be omitted.", "zh_CN":"查询的结束日期,精确到分钟,日期格式为yyyy-mm-dd hh:MM,若没有输入时、分，则时分默认为24:00；此参数需与startdate参数配合,若存在date参数,则该参数无效。"}
  Enddate *string `json:"enddate,omitempty" xml:"enddate,omitempty"`
  // {"en":"domains that been queried:
  // 1.If there are multiple inputs,use  ';' as separator.
  // 2.If not specified, it means all the domains of the account .", "zh_CN":"查询的频道，多个频道值请用英文分号';'，不选或者为空时默认为所查询客户的所有频道"}
  Channel *string `json:"channel,omitempty" xml:"channel,omitempty"`
  // {"en":"1.If there are multiple inputs,use ';' as separator.For example,u can use 'region=cn;apac' to query data of cn and apac region.
  // 2.If not specified, it means all the regions.", "zh_CN":"查询的加速区域的缩写，多个区域请用英文分号';'分隔开，如查询大陆及亚太区域，参数填写为：'region=cn;apac'。不选或者为空时默认为全部区域。"}
  Region *string `json:"region,omitempty" xml:"region,omitempty"`
  // {"en":"acceleration type.
  // 1.If there are multiple inputs,use ';' as separator.
  // 2.If not specified or specified as 'all', it means all the accetypes.", "zh_CN":"加速类型参数，如accetype=web。多个请用英文分号';'分隔开，不填或值为all表示所有类型"}
  Accetype *string `json:"accetype,omitempty" xml:"accetype,omitempty"`
  // {"en":"The response format:
  // 1.optional values:xml, json.
  // 2.'xml' as default.", "zh_CN":"返回结果格式,支持格式为xml和json,默认为xml"}
  Dataformat *string `json:"dataformat,omitempty" xml:"dataformat,omitempty"`
  // {"en":"Specifies if  the 'channel' parameter should be exactly matched:
  // 1.'true' as default.
  // 2. If not 'true',it will query data of channels that ends with any item of input 'channel's.", "zh_CN":"&nbsp;频道是否完全匹配,为true时，必须填写完整的域名(此时会过滤用户输入的无效或重复频道,所有输入频道都无效时返403.。不为true时，显示以用户输入的频道为结尾的所有频道。默认为true"}
  IsExactMatch *string `json:"isExactMatch,omitempty" xml:"isExactMatch,omitempty"`
  // {"en":"1.If there isp multiple inputs,use ';' as demimeter.
  // 2.optional values of isp: refers to the ISP-section of appendix.
  // 3. If not specified,means all the isp.", "zh_CN":"&nbsp;要查询的运营商的缩写，多个isp请用英文分号';'分隔开。运营商的缩写格式参考附录：具体运行商（ISP）信息的代号。备注：只有当地区只写了'cn'时，填写isp信息才有效。不选或者为空时默认为所有isp。"}
  Isp *string `json:"isp,omitempty" xml:"isp,omitempty"`
  // {"en":"Display statistic result in merged or separate way.
  // 1.If specified 1,get the merged result.
  // 2.If specified 2,get the separate result.
  // 3.If specified 3,get both merged result and separate result.
  // 4.If not specified,means '1'.", "zh_CN":"&nbsp;结果的显示是否提供合并值。填写1时：只提供合并结果；填写2时：只提供拆分值；填写3时：既提供合并值，又提供拆分值。不选或者为空时默认为'1'。"}
  ResultType *string `json:"resultType,omitempty" xml:"resultType,omitempty"`
  // {"en":"If return the flow details:
  // choose 1: Return 
  // choose 0: Not return 
  // the default is 0", "zh_CN":"needFlow 是否需要返回流量明细，1：需要；0：不需要。默认为0."}
  NeedFlow *string `json:"needFlow,omitempty" xml:"needFlow,omitempty"`
  // {"en":"flowInfo:displays bandwidth peak, peak time, and total flow information;", "zh_CN":"flowInfo ：展示带宽峰值、峰值时间、总流量信息；"}
  OptionalFields *string `json:"optionalFields,omitempty" xml:"optionalFields,omitempty"`
}

func (s BandwidthChannelRequest) String() string {
  return tea.Prettify(s)
}

func (s BandwidthChannelRequest) GoString() string {
  return s.String()
}

func (s *BandwidthChannelRequest) SetCust(v string) *BandwidthChannelRequest {
  s.Cust = &v
  return s
}

func (s *BandwidthChannelRequest) SetDate(v string) *BandwidthChannelRequest {
  s.Date = &v
  return s
}

func (s *BandwidthChannelRequest) SetStartdate(v string) *BandwidthChannelRequest {
  s.Startdate = &v
  return s
}

func (s *BandwidthChannelRequest) SetEnddate(v string) *BandwidthChannelRequest {
  s.Enddate = &v
  return s
}

func (s *BandwidthChannelRequest) SetChannel(v string) *BandwidthChannelRequest {
  s.Channel = &v
  return s
}

func (s *BandwidthChannelRequest) SetRegion(v string) *BandwidthChannelRequest {
  s.Region = &v
  return s
}

func (s *BandwidthChannelRequest) SetAccetype(v string) *BandwidthChannelRequest {
  s.Accetype = &v
  return s
}

func (s *BandwidthChannelRequest) SetDataformat(v string) *BandwidthChannelRequest {
  s.Dataformat = &v
  return s
}

func (s *BandwidthChannelRequest) SetIsExactMatch(v string) *BandwidthChannelRequest {
  s.IsExactMatch = &v
  return s
}

func (s *BandwidthChannelRequest) SetIsp(v string) *BandwidthChannelRequest {
  s.Isp = &v
  return s
}

func (s *BandwidthChannelRequest) SetResultType(v string) *BandwidthChannelRequest {
  s.ResultType = &v
  return s
}

func (s *BandwidthChannelRequest) SetNeedFlow(v string) *BandwidthChannelRequest {
  s.NeedFlow = &v
  return s
}

func (s *BandwidthChannelRequest) SetOptionalFields(v string) *BandwidthChannelRequest {
  s.OptionalFields = &v
  return s
}

type BandwidthChannelResponse struct {
  // {'en':'provider', 'zh_CN':'结果'}
  Provider *BandwidthChannelResponseProvider `json:"provider,omitempty" xml:"provider,omitempty" require:"true" type:"Struct"`
}

func (s BandwidthChannelResponse) String() string {
  return tea.Prettify(s)
}

func (s BandwidthChannelResponse) GoString() string {
  return s.String()
}

func (s *BandwidthChannelResponse) SetProvider(v *BandwidthChannelResponseProvider) *BandwidthChannelResponse {
  s.Provider = v
  return s
}

type BandwidthChannelResponseProvider struct {
  // {'en':'tenant', 'zh_CN':'租户'}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {'en':'type', 'zh_CN':'接口类型'}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
  // {'en':'resultType', 'zh_CN':'统计类型'}
  ResultType *string `json:"resultType,omitempty" xml:"resultType,omitempty" require:"true"`
  // {'en':'data', 'zh_CN':'频道带宽数据'}
  Date *BandwidthChannelResponseProviderDate `json:"date,omitempty" xml:"date,omitempty" require:"true" type:"Struct"`
}

func (s BandwidthChannelResponseProvider) String() string {
  return tea.Prettify(s)
}

func (s BandwidthChannelResponseProvider) GoString() string {
  return s.String()
}

func (s *BandwidthChannelResponseProvider) SetName(v string) *BandwidthChannelResponseProvider {
  s.Name = &v
  return s
}

func (s *BandwidthChannelResponseProvider) SetType(v string) *BandwidthChannelResponseProvider {
  s.Type = &v
  return s
}

func (s *BandwidthChannelResponseProvider) SetResultType(v string) *BandwidthChannelResponseProvider {
  s.ResultType = &v
  return s
}

func (s *BandwidthChannelResponseProvider) SetDate(v *BandwidthChannelResponseProviderDate) *BandwidthChannelResponseProvider {
  s.Date = v
  return s
}

type BandwidthChannelResponseProviderDate struct {
  // {'en':'startdate', 'zh_CN':'开始时间'}
  Startdate *string `json:"startdate,omitempty" xml:"startdate,omitempty" require:"true"`
  // {'en':'enddate', 'zh_CN':'结束时间'}
  Enddate *string `json:"enddate,omitempty" xml:"enddate,omitempty" require:"true"`
  // {'en':'channel', 'zh_CN':'频道'}
  Channel *BandwidthChannelResponseProviderDateChannel `json:"channel,omitempty" xml:"channel,omitempty" require:"true" type:"Struct"`
}

func (s BandwidthChannelResponseProviderDate) String() string {
  return tea.Prettify(s)
}

func (s BandwidthChannelResponseProviderDate) GoString() string {
  return s.String()
}

func (s *BandwidthChannelResponseProviderDate) SetStartdate(v string) *BandwidthChannelResponseProviderDate {
  s.Startdate = &v
  return s
}

func (s *BandwidthChannelResponseProviderDate) SetEnddate(v string) *BandwidthChannelResponseProviderDate {
  s.Enddate = &v
  return s
}

func (s *BandwidthChannelResponseProviderDate) SetChannel(v *BandwidthChannelResponseProviderDateChannel) *BandwidthChannelResponseProviderDate {
  s.Channel = v
  return s
}

type BandwidthChannelResponseProviderDateChannel struct {
  // {'en':'channel', 'zh_CN':'频道'}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {'en':'bandwidth', 'zh_CN':'带宽数据'}
  Bandwidth []*BandwidthChannelResponseProviderDateChannelBandwidth `json:"bandwidth,omitempty" xml:"bandwidth,omitempty" require:"true" type:"Repeated"`
}

func (s BandwidthChannelResponseProviderDateChannel) String() string {
  return tea.Prettify(s)
}

func (s BandwidthChannelResponseProviderDateChannel) GoString() string {
  return s.String()
}

func (s *BandwidthChannelResponseProviderDateChannel) SetName(v string) *BandwidthChannelResponseProviderDateChannel {
  s.Name = &v
  return s
}

func (s *BandwidthChannelResponseProviderDateChannel) SetBandwidth(v []*BandwidthChannelResponseProviderDateChannelBandwidth) *BandwidthChannelResponseProviderDateChannel {
  s.Bandwidth = v
  return s
}

type BandwidthChannelResponseProviderDateChannelBandwidth struct     {
  // {'en':'timestamp', 'zh_CN':'时间点'}
  Time *string `json:"time,omitempty" xml:"time,omitempty" require:"true"`
  // {'en':'bandwidth', 'zh_CN':'带宽，单位Mbps'}
  Text *string `json:"text,omitempty" xml:"text,omitempty" require:"true"`
}

func (s BandwidthChannelResponseProviderDateChannelBandwidth) String() string {
  return tea.Prettify(s)
}

func (s BandwidthChannelResponseProviderDateChannelBandwidth) GoString() string {
  return s.String()
}

func (s *BandwidthChannelResponseProviderDateChannelBandwidth) SetTime(v string) *BandwidthChannelResponseProviderDateChannelBandwidth {
  s.Time = &v
  return s
}

func (s *BandwidthChannelResponseProviderDateChannelBandwidth) SetText(v string) *BandwidthChannelResponseProviderDateChannelBandwidth {
  s.Text = &v
  return s
}

type BandwidthChannelPaths struct {
}

func (s BandwidthChannelPaths) String() string {
  return tea.Prettify(s)
}

func (s BandwidthChannelPaths) GoString() string {
  return s.String()
}

type BandwidthChannelParameters struct {
}

func (s BandwidthChannelParameters) String() string {
  return tea.Prettify(s)
}

func (s BandwidthChannelParameters) GoString() string {
  return s.String()
}

type BandwidthChannelRequestHeader struct {
}

func (s BandwidthChannelRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s BandwidthChannelRequestHeader) GoString() string {
  return s.String()
}

type BandwidthChannelResponseHeader struct {
}

func (s BandwidthChannelResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s BandwidthChannelResponseHeader) GoString() string {
  return s.String()
}




type GetBandwidthLogRequest struct {
  // {"en":"cust_en_name of sub-client.
  // When a merged-account wants to  view the information of the subclient,the cust_en_name is required.", "zh_CN":"合并账号下的某个客户的英文名，当合并账号要查看子客户的信息时，必须填写子客户的英文名"}
  Cust *string `json:"cust,omitempty" xml:"cust,omitempty"`
  // {"en":"Specifies the query date:
  // 1.With format yyyy-mm-dd.
  // 2.If not specified,it means today as default.", "zh_CN":"查询的日期，日期格式为yyyy-mm-dd,不选或者为空时默认为当天；"}
  Date *string `json:"date,omitempty" xml:"date,omitempty"`
  // {"en":"1.Must work with 'enddate' and they  specify the query date scope.
  // 2.With format yyyy-mm-dd hh:MM.If 'hh:MM' not specified,it means '00:01'.
  // 3.If there is a 'date' parameter,this parameter will be omitted.", "zh_CN":"查询的起始日期,精确到分钟,日期格式为yyyy-mm-dd hh:MM若没有输入时、分，则时分默认为00:01；此参数需与enddate参数配合,若存在date参数,则该参数无效"}
  Startdate *string `json:"startdate,omitempty" xml:"startdate,omitempty"`
  // {"en":"1.Must work with 'startdate' and they  specify the query date scope.
  // 2.With format yyyy-mm-dd hh:MM.If 'hh:MM' not specified,it means '24:00'.
  // 3.If there is a 'date' parameter,this parameter will be omitted.", "zh_CN":"查询的结束日期,精确到分钟,日期格式为yyyy-mm-dd hh:MM,若没有输入时、分，则时分默认为24:00；此参数需与startdate参数配合,若存在date参数,则该参数无效。"}
  Enddate *string `json:"enddate,omitempty" xml:"enddate,omitempty"`
  // {"en":"GMT time zone, parameter format: GMT+09:00 means east 9th zone, GMT-09:00 means west 9th zone, if not transmitted, the default is local time zone (east 8th zone).", "zh_CN":"格林尼治时区，参数格式 GMT+09:00 表示东九区，GMT-09:00 表示西9区，不传则默认为本地时区（东八区）"}
  Timezone *string `json:"timezone,omitempty" xml:"timezone,omitempty"`
  // {"en":"domains that been queried:
  // 1.If there are multiple inputs,use  ';' as separator.
  // 2.If not specified, it means all the domains of the account .", "zh_CN":"查询的频道，多个频道值请用英文分号';'，不选或者为空时默认为所查询客户的所有频道"}
  Channel *string `json:"channel,omitempty" xml:"channel,omitempty"`
  // {"en":"1.If there are multiple inputs,use ';' as separator.For example,u can use 'region=cn;apac' to query data of cn and apac region.
  // 2.If not specified, it means all the regions.", "zh_CN":"查询的加速区域的缩写，多个区域请用英文分号';'分隔开，如查询大陆及亚太区域，参数填写为：'region=cn;apac'。不选或者为空时默认为全部区域。"}
  Region *string `json:"region,omitempty" xml:"region,omitempty"`
  // {"en":"acceleration type.
  // 1.If there are multiple inputs,use ';' as separator.
  // 2.If not specified or specified as 'all', it means all the accetypes.", "zh_CN":"加速类型参数，如accetype=web。多个请用英文分号';'分隔开，不填或值为all表示所有类型"}
  Accetype *string `json:"accetype,omitempty" xml:"accetype,omitempty"`
  // {"en":"The response format:
  // 1.optional values:xml, json.
  // 2.'xml' as default.", "zh_CN":"返回结果格式,支持格式为xml和json,默认为xml"}
  Dataformat *string `json:"dataformat,omitempty" xml:"dataformat,omitempty"`
  // {"en":"Display statistic result in merged or separate way.
  // 1.If specified 1,get the merged result.
  // 2.If specified 2,get the separate result.
  // 3.If specified 3,get both merged result and separate result.
  // 4.If not specified,means '1'.", "zh_CN":"&nbsp;结果的显示是否提供合并值。填写1时：只提供合并结果；填写2时：只提供拆分值；填写3时：既提供合并值，又提供拆分值。不选或者为空时默认为'1'。"}
  ResultType *string `json:"resultType,omitempty" xml:"resultType,omitempty"`
  // {"en":"The unit of the flow value in the details, the default is Mbps. Optional byte (byte) or bps", "zh_CN":"明细中的流量值单位，默认为Mbps。可选 byte (字节)或者bps"}
  FlowUnit *string `json:"flowUnit,omitempty" xml:"flowUnit,omitempty"`
  // {"en":"If specified 0,the time return as 00:05--24:00,If specified 1,the time return as 00:00--23:55,If not specified, the default value is 0", "zh_CN":"当值为0:返回00:05--24:00；当值为1:返回00:00--23:55。不传默认为:0"}
  TimeFromZero *string `json:"timeFromZero,omitempty" xml:"timeFromZero,omitempty"`
}

func (s GetBandwidthLogRequest) String() string {
  return tea.Prettify(s)
}

func (s GetBandwidthLogRequest) GoString() string {
  return s.String()
}

func (s *GetBandwidthLogRequest) SetCust(v string) *GetBandwidthLogRequest {
  s.Cust = &v
  return s
}

func (s *GetBandwidthLogRequest) SetDate(v string) *GetBandwidthLogRequest {
  s.Date = &v
  return s
}

func (s *GetBandwidthLogRequest) SetStartdate(v string) *GetBandwidthLogRequest {
  s.Startdate = &v
  return s
}

func (s *GetBandwidthLogRequest) SetEnddate(v string) *GetBandwidthLogRequest {
  s.Enddate = &v
  return s
}

func (s *GetBandwidthLogRequest) SetTimezone(v string) *GetBandwidthLogRequest {
  s.Timezone = &v
  return s
}

func (s *GetBandwidthLogRequest) SetChannel(v string) *GetBandwidthLogRequest {
  s.Channel = &v
  return s
}

func (s *GetBandwidthLogRequest) SetRegion(v string) *GetBandwidthLogRequest {
  s.Region = &v
  return s
}

func (s *GetBandwidthLogRequest) SetAccetype(v string) *GetBandwidthLogRequest {
  s.Accetype = &v
  return s
}

func (s *GetBandwidthLogRequest) SetDataformat(v string) *GetBandwidthLogRequest {
  s.Dataformat = &v
  return s
}

func (s *GetBandwidthLogRequest) SetResultType(v string) *GetBandwidthLogRequest {
  s.ResultType = &v
  return s
}

func (s *GetBandwidthLogRequest) SetFlowUnit(v string) *GetBandwidthLogRequest {
  s.FlowUnit = &v
  return s
}

func (s *GetBandwidthLogRequest) SetTimeFromZero(v string) *GetBandwidthLogRequest {
  s.TimeFromZero = &v
  return s
}

type GetBandwidthLogResponse struct {
  // {'en':'provider', 'zh_CN':'结果'}
  Provider *GetBandwidthLogResponseProvider `json:"provider,omitempty" xml:"provider,omitempty" require:"true" type:"Struct"`
}

func (s GetBandwidthLogResponse) String() string {
  return tea.Prettify(s)
}

func (s GetBandwidthLogResponse) GoString() string {
  return s.String()
}

func (s *GetBandwidthLogResponse) SetProvider(v *GetBandwidthLogResponseProvider) *GetBandwidthLogResponse {
  s.Provider = v
  return s
}

type GetBandwidthLogResponseProvider struct {
  // {'en':'tenant', 'zh_CN':'租户'}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {'en':'type', 'zh_CN':'接口类型'}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
  // {'en':'resultType', 'zh_CN':'统计类型'}
  ResultType *string `json:"resultType,omitempty" xml:"resultType,omitempty" require:"true"`
  // {'en':'data', 'zh_CN':'频道带宽数据'}
  Date *GetBandwidthLogResponseProviderDate `json:"date,omitempty" xml:"date,omitempty" require:"true" type:"Struct"`
}

func (s GetBandwidthLogResponseProvider) String() string {
  return tea.Prettify(s)
}

func (s GetBandwidthLogResponseProvider) GoString() string {
  return s.String()
}

func (s *GetBandwidthLogResponseProvider) SetName(v string) *GetBandwidthLogResponseProvider {
  s.Name = &v
  return s
}

func (s *GetBandwidthLogResponseProvider) SetType(v string) *GetBandwidthLogResponseProvider {
  s.Type = &v
  return s
}

func (s *GetBandwidthLogResponseProvider) SetResultType(v string) *GetBandwidthLogResponseProvider {
  s.ResultType = &v
  return s
}

func (s *GetBandwidthLogResponseProvider) SetDate(v *GetBandwidthLogResponseProviderDate) *GetBandwidthLogResponseProvider {
  s.Date = v
  return s
}

type GetBandwidthLogResponseProviderDate struct {
  // {'en':'startdate', 'zh_CN':'开始时间'}
  Startdate *string `json:"startdate,omitempty" xml:"startdate,omitempty" require:"true"`
  // {'en':'enddate', 'zh_CN':'结束时间'}
  Enddate *string `json:"enddate,omitempty" xml:"enddate,omitempty" require:"true"`
  // {'en':'channel', 'zh_CN':'频道'}
  Channel *GetBandwidthLogResponseProviderDateChannel `json:"channel,omitempty" xml:"channel,omitempty" require:"true" type:"Struct"`
}

func (s GetBandwidthLogResponseProviderDate) String() string {
  return tea.Prettify(s)
}

func (s GetBandwidthLogResponseProviderDate) GoString() string {
  return s.String()
}

func (s *GetBandwidthLogResponseProviderDate) SetStartdate(v string) *GetBandwidthLogResponseProviderDate {
  s.Startdate = &v
  return s
}

func (s *GetBandwidthLogResponseProviderDate) SetEnddate(v string) *GetBandwidthLogResponseProviderDate {
  s.Enddate = &v
  return s
}

func (s *GetBandwidthLogResponseProviderDate) SetChannel(v *GetBandwidthLogResponseProviderDateChannel) *GetBandwidthLogResponseProviderDate {
  s.Channel = v
  return s
}

type GetBandwidthLogResponseProviderDateChannel struct {
  // {'en':'channel', 'zh_CN':'频道'}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {'en':'bandwidth', 'zh_CN':'带宽数据'}
  Bandwidth []*GetBandwidthLogResponseProviderDateChannelBandwidth `json:"bandwidth,omitempty" xml:"bandwidth,omitempty" require:"true" type:"Repeated"`
}

func (s GetBandwidthLogResponseProviderDateChannel) String() string {
  return tea.Prettify(s)
}

func (s GetBandwidthLogResponseProviderDateChannel) GoString() string {
  return s.String()
}

func (s *GetBandwidthLogResponseProviderDateChannel) SetName(v string) *GetBandwidthLogResponseProviderDateChannel {
  s.Name = &v
  return s
}

func (s *GetBandwidthLogResponseProviderDateChannel) SetBandwidth(v []*GetBandwidthLogResponseProviderDateChannelBandwidth) *GetBandwidthLogResponseProviderDateChannel {
  s.Bandwidth = v
  return s
}

type GetBandwidthLogResponseProviderDateChannelBandwidth struct     {
  // {'en':'timestamp', 'zh_CN':'时间点'}
  Time *string `json:"time,omitempty" xml:"time,omitempty" require:"true"`
  // {'en':'bandwidth', 'zh_CN':'带宽'}
  Text *string `json:"text,omitempty" xml:"text,omitempty" require:"true"`
}

func (s GetBandwidthLogResponseProviderDateChannelBandwidth) String() string {
  return tea.Prettify(s)
}

func (s GetBandwidthLogResponseProviderDateChannelBandwidth) GoString() string {
  return s.String()
}

func (s *GetBandwidthLogResponseProviderDateChannelBandwidth) SetTime(v string) *GetBandwidthLogResponseProviderDateChannelBandwidth {
  s.Time = &v
  return s
}

func (s *GetBandwidthLogResponseProviderDateChannelBandwidth) SetText(v string) *GetBandwidthLogResponseProviderDateChannelBandwidth {
  s.Text = &v
  return s
}

type GetBandwidthLogPaths struct {
}

func (s GetBandwidthLogPaths) String() string {
  return tea.Prettify(s)
}

func (s GetBandwidthLogPaths) GoString() string {
  return s.String()
}

type GetBandwidthLogParameters struct {
}

func (s GetBandwidthLogParameters) String() string {
  return tea.Prettify(s)
}

func (s GetBandwidthLogParameters) GoString() string {
  return s.String()
}

type GetBandwidthLogRequestHeader struct {
}

func (s GetBandwidthLogRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetBandwidthLogRequestHeader) GoString() string {
  return s.String()
}

type GetBandwidthLogResponseHeader struct {
}

func (s GetBandwidthLogResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetBandwidthLogResponseHeader) GoString() string {
  return s.String()
}




type QueryRealTimeBandwidthForMultiDomainRequest struct {
  // {'en':'Start time: 
  //     1.Start time: time format is yyyy-MM-ddTHH:mm:ss+08:00, for example, 2016-12-02T10:00:00+08:00 (December 2rd, 2016, 10:00 a.m., Beijing Time); 
  //     2.Not greater than the current time;
  //     3.The most recent half-year (183 days) data can be obtained
  // ', 'zh_CN':'开始时间：
  // 
  // 时间格式为yyyy-MM-ddTHH:mm:ss+08:00，例如，2019-01-01T10:00:00+08:00（为北京时间2019年1月1日10点0分0秒）；
  // 不能大于当前时间
  // 最多可获取最近半年（183天）的数据。'}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty"`
  // {'en':'End time: 
  //     1.The time format is 2016-12-02T10:00:00+08:00 
  //     2.End time should be greater than start time. If the end time is greater than current time, current time will be used. 
  //     3.If both fields of dataFrom and dateTo are left empty, then data in the last 1 hours will be queried by default; if one field is filled and one is left empty, then exception will occur. 
  //     4.Maximum time range allowable for query: The default value is 1 hour, that is, the difference between dateFrom and dateTo cannot exceed 1 hour (you can contact technical support to adjust it, the maximum is 31 days)
  // ', 'zh_CN':'结束时间：
  // 
  // 时间格式为yyyy-MM-ddTHH:mm:ss+08:00
  // 结束时间需大于开始时间，结束时间如果大于当前时间，取当前时间。
  // dateFrom，dateTo二者都未传，默认查询过去的1小时；如仅有一个未传，抛异常
  // 允许查询最大时间间隔：默认1小时，即dateFrom和dateTo相差不能超过1小时（可联系技术支持调整，最长31天）。'}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty"`
  // {'en':'Data granularity: 
  //     1. Support 1m (1 minute granularity),5m (5 minutes granularity) 
  //     2. The default value is 1m
  // ', 'zh_CN':'数据粒度：不传默认1m
  // 
  // 支持1m（1分钟）、5m（5分钟）'}
  DataInterval *string `json:"dataInterval,omitempty" xml:"dataInterval,omitempty"`
  // {'en':'Domain: 
  //     1.The default upper limit of domains that can be entered is 20 (if you want to adjust, please, contact technical support); 
  //     2.Automatically filter out illegal domains (illegal domains will be filtered out, the query results only return the data of legitimate domains)
  // ', 'zh_CN':'域名：
  // 
  // 可传递域名数量上限默认为20（可联系技术支持调整）；
  // 自动过滤掉非法域名（如传递非法域名，会被过滤掉，查询结果只返回合法域名的数据）'}
  Domain []*string `json:"domain,omitempty" xml:"domain,omitempty" require:"true" type:"Repeated"`
  // {'en':'The optional value of the grouping dimension is domain; if it is passed in, detailed data will be displayed according to this dimension;', 'zh_CN':'分组维度
  // 
  // 可选值为domain；
  // 有传入则按照该维度展示明细数据；'}
  GroupBy *string `json:"groupBy,omitempty" xml:"groupBy,omitempty"`
}

func (s QueryRealTimeBandwidthForMultiDomainRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryRealTimeBandwidthForMultiDomainRequest) GoString() string {
  return s.String()
}

func (s *QueryRealTimeBandwidthForMultiDomainRequest) SetDateFrom(v string) *QueryRealTimeBandwidthForMultiDomainRequest {
  s.DateFrom = &v
  return s
}

func (s *QueryRealTimeBandwidthForMultiDomainRequest) SetDateTo(v string) *QueryRealTimeBandwidthForMultiDomainRequest {
  s.DateTo = &v
  return s
}

func (s *QueryRealTimeBandwidthForMultiDomainRequest) SetDataInterval(v string) *QueryRealTimeBandwidthForMultiDomainRequest {
  s.DataInterval = &v
  return s
}

func (s *QueryRealTimeBandwidthForMultiDomainRequest) SetDomain(v []*string) *QueryRealTimeBandwidthForMultiDomainRequest {
  s.Domain = v
  return s
}

func (s *QueryRealTimeBandwidthForMultiDomainRequest) SetGroupBy(v string) *QueryRealTimeBandwidthForMultiDomainRequest {
  s.GroupBy = &v
  return s
}

type QueryRealTimeBandwidthForMultiDomainResponse struct {
  // {'en':'request result status code', 'zh_CN':'请求结果状态码'}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {'en':'request result information', 'zh_CN':'请求结果信息'}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  Data []*QueryRealTimeBandwidthForMultiDomainResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s QueryRealTimeBandwidthForMultiDomainResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryRealTimeBandwidthForMultiDomainResponse) GoString() string {
  return s.String()
}

func (s *QueryRealTimeBandwidthForMultiDomainResponse) SetCode(v string) *QueryRealTimeBandwidthForMultiDomainResponse {
  s.Code = &v
  return s
}

func (s *QueryRealTimeBandwidthForMultiDomainResponse) SetMessage(v string) *QueryRealTimeBandwidthForMultiDomainResponse {
  s.Message = &v
  return s
}

func (s *QueryRealTimeBandwidthForMultiDomainResponse) SetData(v []*QueryRealTimeBandwidthForMultiDomainResponseData) *QueryRealTimeBandwidthForMultiDomainResponse {
  s.Data = v
  return s
}

type QueryRealTimeBandwidthForMultiDomainResponseData struct     {
  // {'en':' Domain name. If you do not select domain name group Dimension, this field is a semicolon-separated string of all domain names.', 'zh_CN':'域名，如果不选择域名分组维度，该字段为所有域名以分号分隔的字符串'}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {'en':'Peak Bandwidth,unit is Mbps,example(9811.21Mbps)', 'zh_CN':'峰值带宽，单位Mbps，示例 （9811.21Mbps)'}
  PeakValue *string `json:"peakValue,omitempty" xml:"peakValue,omitempty" require:"true"`
  // {'en':'Time of peak bandwidth,example(2019-02-13 18:01)', 'zh_CN':'峰值时间，示例（2019-02-13 18:01）'}
  PeakTime *string `json:"peakTime,omitempty" xml:"peakTime,omitempty" require:"true"`
  // {'en':'Edge total traffic,example(74099.91MB)', 'zh_CN':'边缘总流量，单位MB，示例 ( 74099.91MB)'}
  Total *string `json:"total,omitempty" xml:"total,omitempty" require:"true"`
  BandwidthData []*QueryRealTimeBandwidthForMultiDomainResponseDataBandwidthData `json:"bandwidthData,omitempty" xml:"bandwidthData,omitempty" require:"true" type:"Repeated"`
}

func (s QueryRealTimeBandwidthForMultiDomainResponseData) String() string {
  return tea.Prettify(s)
}

func (s QueryRealTimeBandwidthForMultiDomainResponseData) GoString() string {
  return s.String()
}

func (s *QueryRealTimeBandwidthForMultiDomainResponseData) SetDomain(v string) *QueryRealTimeBandwidthForMultiDomainResponseData {
  s.Domain = &v
  return s
}

func (s *QueryRealTimeBandwidthForMultiDomainResponseData) SetPeakValue(v string) *QueryRealTimeBandwidthForMultiDomainResponseData {
  s.PeakValue = &v
  return s
}

func (s *QueryRealTimeBandwidthForMultiDomainResponseData) SetPeakTime(v string) *QueryRealTimeBandwidthForMultiDomainResponseData {
  s.PeakTime = &v
  return s
}

func (s *QueryRealTimeBandwidthForMultiDomainResponseData) SetTotal(v string) *QueryRealTimeBandwidthForMultiDomainResponseData {
  s.Total = &v
  return s
}

func (s *QueryRealTimeBandwidthForMultiDomainResponseData) SetBandwidthData(v []*QueryRealTimeBandwidthForMultiDomainResponseDataBandwidthData) *QueryRealTimeBandwidthForMultiDomainResponseData {
  s.BandwidthData = v
  return s
}

type QueryRealTimeBandwidthForMultiDomainResponseDataBandwidthData struct     {
  // {'en':'Time: 1. When the data query granularity is 1m, then the format is yyyy-MM-dd HH:mm; 
  //     Each time slice value represents the value within the previous time granularity range. The first time slice of the day is yyyy-MM-dd 00:01, and the last one is (yyyy-MM-dd+1) 00:00; 
  //     2. Return the time slices that contained in start time and in end time.
  // ', 'zh_CN':'格式为yyyy-MM-dd HH:mm；每一个时间片数据值代表的是前一个时间粒度范围内的数据值。
  // 
  // 一天开始的时间片是yyyy-MM-dd 00:01，最后一个时间片是第二天（yyyy-MM-dd） 00:00。
  // 
  // 返回开始时间和结束时间包含的时间片'}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {'en':'Edge bandwidth,the unit is Mbps,keep 2 decimal places', 'zh_CN':'带宽值，单位Mbps，保留2位小数。'}
  Value *string `json:"value,omitempty" xml:"value,omitempty" require:"true"`
}

func (s QueryRealTimeBandwidthForMultiDomainResponseDataBandwidthData) String() string {
  return tea.Prettify(s)
}

func (s QueryRealTimeBandwidthForMultiDomainResponseDataBandwidthData) GoString() string {
  return s.String()
}

func (s *QueryRealTimeBandwidthForMultiDomainResponseDataBandwidthData) SetTimestamp(v string) *QueryRealTimeBandwidthForMultiDomainResponseDataBandwidthData {
  s.Timestamp = &v
  return s
}

func (s *QueryRealTimeBandwidthForMultiDomainResponseDataBandwidthData) SetValue(v string) *QueryRealTimeBandwidthForMultiDomainResponseDataBandwidthData {
  s.Value = &v
  return s
}

type QueryRealTimeBandwidthForMultiDomainPaths struct {
}

func (s QueryRealTimeBandwidthForMultiDomainPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryRealTimeBandwidthForMultiDomainPaths) GoString() string {
  return s.String()
}

type QueryRealTimeBandwidthForMultiDomainParameters struct {
}

func (s QueryRealTimeBandwidthForMultiDomainParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryRealTimeBandwidthForMultiDomainParameters) GoString() string {
  return s.String()
}

type QueryRealTimeBandwidthForMultiDomainRequestHeader struct {
}

func (s QueryRealTimeBandwidthForMultiDomainRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryRealTimeBandwidthForMultiDomainRequestHeader) GoString() string {
  return s.String()
}

type QueryRealTimeBandwidthForMultiDomainResponseHeader struct {
}

func (s QueryRealTimeBandwidthForMultiDomainResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryRealTimeBandwidthForMultiDomainResponseHeader) GoString() string {
  return s.String()
}




type QueryCPSBandwidthRequest struct {
  // {"en":"cust_en_name.If there are multiple inputs,use ';' as separator", "zh_CN":"客户的英文名，多个';'隔开"}
  Cust *string `json:"cust,omitempty" xml:"cust,omitempty"`
  // {"en":"1.Must work with 'enddate' and they specify the query date scope.
  // 2.With format yyyy-mm-dd.
  // 3.If there is a 'date' parameter,this parameter will be omitted.", "zh_CN":"查询的起始日期,日期格式为yyyy-mm-dd；
  // 
  // 此参数需与enddate参数配合,若存在date参数,则该参数无效"}
  StartDate *string `json:"startDate,omitempty" xml:"startDate,omitempty"`
  // {"en":"1.Must work with 'startdate' and they specify the query date scope.
  // 
  // 2.With format yyyy-mm-dd 3.If there is a 'date' parameter,this parameter will be omitted.", "zh_CN":"查询的结束日期,日期格式为yyyy-mm-dd；
  // 
  // 此参数需与startdate参数配合,若存在date参数,则该参数无效"}
  EndDate *string `json:"endDate,omitempty" xml:"endDate,omitempty"`
  // {"en":"domains that been queried:
  // 1.If there are multiple inputs,use ';' as separator.
  // 2.If not specified, it means all the domains of the account .", "zh_CN":"查询的频道，多个频道值请用英文分号';'，
  // 
  // 返回多个频道的汇总值，不选或者为空时默认为所查询客户的所有频道"}
  Channel *string `json:"channel,omitempty" xml:"channel,omitempty"`
  // {"en":"GMT time zone, parameter format: GMT+09:00 means east 9th zone, GMT-09:00 means west 9th zone, if not transmitted, the default is local time zone (east 8th zone).", "zh_CN":"格林尼治时区，参数格式 GMT+09:00 表示东九区，
  // 
  // GMT-09:00 表示西9区，不传则默认为本地时区（东八区）"}
  Timezone *string `json:"timezone,omitempty" xml:"timezone,omitempty"`
  // {"en":"Input 'flowInfo', which will show the bandwidth peak, peak time, total flow. not selected or null&nbsp;will not be showed", "zh_CN":"入参flowInfo，将展示带宽峰值、峰值时间、总流量,不选或者为空默认不展示"}
  OptionalFields *string `json:"optionalFields,omitempty" xml:"optionalFields,omitempty"`
  // {"en":"This parameter represents the type of the dedicated line. It can have multiple values. Use the semicolon as a separator if there are multiple values. hk: represents China Premium Service; jp: represents China Premium Service-Basic. eu: not applicable to CDNW. all: not applicable to CDNW.", "zh_CN":"专线类型:hk;jp，支持多个，用分号隔开。hk:中港；jp:中日；eu:中欧；all:所有，默认为：hk"}
  LineType *string `json:"lineType,omitempty" xml:"lineType,omitempty"`
}

func (s QueryCPSBandwidthRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryCPSBandwidthRequest) GoString() string {
  return s.String()
}

func (s *QueryCPSBandwidthRequest) SetCust(v string) *QueryCPSBandwidthRequest {
  s.Cust = &v
  return s
}

func (s *QueryCPSBandwidthRequest) SetStartDate(v string) *QueryCPSBandwidthRequest {
  s.StartDate = &v
  return s
}

func (s *QueryCPSBandwidthRequest) SetEndDate(v string) *QueryCPSBandwidthRequest {
  s.EndDate = &v
  return s
}

func (s *QueryCPSBandwidthRequest) SetChannel(v string) *QueryCPSBandwidthRequest {
  s.Channel = &v
  return s
}

func (s *QueryCPSBandwidthRequest) SetTimezone(v string) *QueryCPSBandwidthRequest {
  s.Timezone = &v
  return s
}

func (s *QueryCPSBandwidthRequest) SetOptionalFields(v string) *QueryCPSBandwidthRequest {
  s.OptionalFields = &v
  return s
}

func (s *QueryCPSBandwidthRequest) SetLineType(v string) *QueryCPSBandwidthRequest {
  s.LineType = &v
  return s
}

type QueryCPSBandwidthResponse struct {
  // {'en':'code', 'zh_CN':'返回编码'}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {'en':'message', 'zh_CN':'返回消息'}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {'en':'data', 'zh_CN':'频道带宽数据'}
  Data *QueryCPSBandwidthResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
}

func (s QueryCPSBandwidthResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryCPSBandwidthResponse) GoString() string {
  return s.String()
}

func (s *QueryCPSBandwidthResponse) SetCode(v string) *QueryCPSBandwidthResponse {
  s.Code = &v
  return s
}

func (s *QueryCPSBandwidthResponse) SetMessage(v string) *QueryCPSBandwidthResponse {
  s.Message = &v
  return s
}

func (s *QueryCPSBandwidthResponse) SetData(v *QueryCPSBandwidthResponseData) *QueryCPSBandwidthResponse {
  s.Data = v
  return s
}

type QueryCPSBandwidthResponseData struct {
  // {'en':'totalFlow', 'zh_CN':'总流量，单位GB'}
  TotalFlow *string `json:"totalFlow,omitempty" xml:"totalFlow,omitempty" require:"true"`
  // {'en':'peakTime', 'zh_CN':'峰值时间'}
  PeakTime *string `json:"peakTime,omitempty" xml:"peakTime,omitempty" require:"true"`
  // {'en':'peakvalue', 'zh_CN':'带宽峰值，单位Mbps'}
  PeakValue *string `json:"peakValue,omitempty" xml:"peakValue,omitempty" require:"true"`
  // {'en':'detail', 'zh_CN':'时点带宽数据'}
  Detail *QueryCPSBandwidthResponseDataDetail `json:"detail,omitempty" xml:"detail,omitempty" require:"true" type:"Struct"`
}

func (s QueryCPSBandwidthResponseData) String() string {
  return tea.Prettify(s)
}

func (s QueryCPSBandwidthResponseData) GoString() string {
  return s.String()
}

func (s *QueryCPSBandwidthResponseData) SetTotalFlow(v string) *QueryCPSBandwidthResponseData {
  s.TotalFlow = &v
  return s
}

func (s *QueryCPSBandwidthResponseData) SetPeakTime(v string) *QueryCPSBandwidthResponseData {
  s.PeakTime = &v
  return s
}

func (s *QueryCPSBandwidthResponseData) SetPeakValue(v string) *QueryCPSBandwidthResponseData {
  s.PeakValue = &v
  return s
}

func (s *QueryCPSBandwidthResponseData) SetDetail(v *QueryCPSBandwidthResponseDataDetail) *QueryCPSBandwidthResponseData {
  s.Detail = v
  return s
}

type QueryCPSBandwidthResponseDataDetail struct {
  // {'en':'timestamp', 'zh_CN':'时间点'}
  Time *string `json:"time,omitempty" xml:"time,omitempty" require:"true"`
  // {'en':'bandwidth', 'zh_CN':'带宽'}
  Text *string `json:"text,omitempty" xml:"text,omitempty" require:"true"`
}

func (s QueryCPSBandwidthResponseDataDetail) String() string {
  return tea.Prettify(s)
}

func (s QueryCPSBandwidthResponseDataDetail) GoString() string {
  return s.String()
}

func (s *QueryCPSBandwidthResponseDataDetail) SetTime(v string) *QueryCPSBandwidthResponseDataDetail {
  s.Time = &v
  return s
}

func (s *QueryCPSBandwidthResponseDataDetail) SetText(v string) *QueryCPSBandwidthResponseDataDetail {
  s.Text = &v
  return s
}

type QueryCPSBandwidthPaths struct {
}

func (s QueryCPSBandwidthPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryCPSBandwidthPaths) GoString() string {
  return s.String()
}

type QueryCPSBandwidthParameters struct {
}

func (s QueryCPSBandwidthParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryCPSBandwidthParameters) GoString() string {
  return s.String()
}

type QueryCPSBandwidthRequestHeader struct {
}

func (s QueryCPSBandwidthRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryCPSBandwidthRequestHeader) GoString() string {
  return s.String()
}

type QueryCPSBandwidthResponseHeader struct {
}

func (s QueryCPSBandwidthResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryCPSBandwidthResponseHeader) GoString() string {
  return s.String()
}




type QueryP2PBandwidthRequest struct {
  // {"en":"cust_en_name of sub-client.
  // When a merged-account wants to view the information of the subclient,the cust_en_name is required.", "zh_CN":"合并账号下的某个客户的英文名，当合并账号要查看子客户的信息时，必须填写子客户的英文名"}
  Cust *string `json:"cust,omitempty" xml:"cust,omitempty"`
  // {"en":"Specifies the query date: 1.With format yyyy-mm-dd. 2.If not specified,it means today as default.", "zh_CN":"查询的日期，日期格式为yyyy-mm-dd,不选或者为空时默认为当天；"}
  Date *string `json:"date,omitempty" xml:"date,omitempty"`
  // {"en":"1.Must work with 'enddate' and they specify the query date scope. 2.With format yyyy-mm-dd. 3.If there is a 'date' parameter,this parameter will be omitted.", "zh_CN":"查询的起始日期,日期格式为yyyy-mm-dd；此参数需与enddate参数配合,若存在date参数,则该参数无效"}
  Startdate *string `json:"startdate,omitempty" xml:"startdate,omitempty"`
  // {"en":"1.Must work with 'startdate' and they specify the query date scope.
  // 2.With format yyyy-mm-dd
  // 3.If there is a 'date' parameter,this parameter will be omitted.", "zh_CN":"查询的结束日期,日期格式为yyyy-mm-dd此参数需与startdate参数配合,若存在date参数,则该参数无效。"}
  Enddate *string `json:"enddate,omitempty" xml:"enddate,omitempty"`
  // {"en":"domains that been queried.
  // 1.If there are multiple inputs,use ';' as separator.
  // 2.If not specified, it means all the domains of the account .", "zh_CN":"查询的频道，多个频道值请用英文分号';'，不选或者为空时默认为所查询客户的所有频道"}
  Channel *string `json:"channel,omitempty" xml:"channel,omitempty"`
  // {"en":"acceleration type.
  // 1.If there are multiple inputs,use ';' as separator.
  // 2.If not specified or specified as 'all', it means all the accetypes.", "zh_CN":"加速类型参数，如accetype=web。多个请用英文分号';'分隔开，不填或值为all表示所有类型"}
  Accetype *string `json:"accetype,omitempty" xml:"accetype,omitempty"`
  // {"en":"The response format:
  // 1.optional values:xml, json.
  // 2.'xml' as default.", "zh_CN":"返回结果格式,支持格式为xml和json,默认为xml"}
  Dataformat *string `json:"dataformat,omitempty" xml:"dataformat,omitempty"`
  // {"en":"Whether the channel is completely matched. When it is true, the complete domain name must be filled in (invalid or duplicate channels entered by the user will be filtered at this time, and 403 will be returned if all input channels are invalid. If it is not true, the display will end with the channel entered by the user All channels. Defaults to true", "zh_CN":"频道是否完全匹配,为true时，必须填写完整的域名(此时会过滤用户输入的无效或重复频道,所有输入频道都无效时返403.。不为true时，显示以用户输入的频道为结尾的所有频道。默认为true"}
  IsExactMatch *string `json:"isExactMatch,omitempty" xml:"isExactMatch,omitempty"`
  // {"en":"A value of 1 means log bandwidth is returned, 0 means cdn bandwidth is returned, and the default value is 0", "zh_CN":"值为1，表示返回的是log带宽，0表示返回的是cdn带宽，默认值为0"}
  IsLog *string `json:"isLog,omitempty" xml:"isLog,omitempty"`
  // {"en":"Control whether to return box bandwidth, the default is 0, when box=1, return cdn(log) bandwidth+p2p bandwidth+box bandwidth; when box=0, return cdn(log) bandwidth+p2p bandwidth", "zh_CN":"控制是否返回盒子带宽,默认为0,box=1时，返回cdn(log)带宽+p2p带宽+box带宽;box=0时，返回cdn(log)带宽+p2p带宽"}
  Box *string `json:"box,omitempty" xml:"box,omitempty"`
  // {"en":"Control whether to return flow details, the default is 0 and no return, return flow details when passing 1", "zh_CN":"控制是否返回流量明细，默认为0不返回，传1时返回流量明细"}
  Flow *string `json:"flow,omitempty" xml:"flow,omitempty"`
  // {"en":"1.If there are multiple inputs,use ';' as separator.For example,u can use 'region=cn;apac' to query data of cn and apac region.
  // 2.If not specified, it means all the regions.", "zh_CN":"查询的加速区域的缩写，多个区域请用英文分号';'分隔开，如查询大陆及亚太区域，参数填写为：'region=cn;apac'。不选或者为空时默认为全部区域。"}
  Region *string `json:"region,omitempty" xml:"region,omitempty"`
}

func (s QueryP2PBandwidthRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryP2PBandwidthRequest) GoString() string {
  return s.String()
}

func (s *QueryP2PBandwidthRequest) SetCust(v string) *QueryP2PBandwidthRequest {
  s.Cust = &v
  return s
}

func (s *QueryP2PBandwidthRequest) SetDate(v string) *QueryP2PBandwidthRequest {
  s.Date = &v
  return s
}

func (s *QueryP2PBandwidthRequest) SetStartdate(v string) *QueryP2PBandwidthRequest {
  s.Startdate = &v
  return s
}

func (s *QueryP2PBandwidthRequest) SetEnddate(v string) *QueryP2PBandwidthRequest {
  s.Enddate = &v
  return s
}

func (s *QueryP2PBandwidthRequest) SetChannel(v string) *QueryP2PBandwidthRequest {
  s.Channel = &v
  return s
}

func (s *QueryP2PBandwidthRequest) SetAccetype(v string) *QueryP2PBandwidthRequest {
  s.Accetype = &v
  return s
}

func (s *QueryP2PBandwidthRequest) SetDataformat(v string) *QueryP2PBandwidthRequest {
  s.Dataformat = &v
  return s
}

func (s *QueryP2PBandwidthRequest) SetIsExactMatch(v string) *QueryP2PBandwidthRequest {
  s.IsExactMatch = &v
  return s
}

func (s *QueryP2PBandwidthRequest) SetIsLog(v string) *QueryP2PBandwidthRequest {
  s.IsLog = &v
  return s
}

func (s *QueryP2PBandwidthRequest) SetBox(v string) *QueryP2PBandwidthRequest {
  s.Box = &v
  return s
}

func (s *QueryP2PBandwidthRequest) SetFlow(v string) *QueryP2PBandwidthRequest {
  s.Flow = &v
  return s
}

func (s *QueryP2PBandwidthRequest) SetRegion(v string) *QueryP2PBandwidthRequest {
  s.Region = &v
  return s
}

type QueryP2PBandwidthResponse struct {
  // {'en':'provider', 'zh_CN':'结果'}
  Provider *QueryP2PBandwidthResponseProvider `json:"provider,omitempty" xml:"provider,omitempty" require:"true" type:"Struct"`
}

func (s QueryP2PBandwidthResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryP2PBandwidthResponse) GoString() string {
  return s.String()
}

func (s *QueryP2PBandwidthResponse) SetProvider(v *QueryP2PBandwidthResponseProvider) *QueryP2PBandwidthResponse {
  s.Provider = v
  return s
}

type QueryP2PBandwidthResponseProvider struct {
  // {'en':'tenant', 'zh_CN':'租户'}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {'en':'type', 'zh_CN':'接口类型'}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
  // {'en':'data', 'zh_CN':'p2p带宽数据'}
  Date *QueryP2PBandwidthResponseProviderDate `json:"date,omitempty" xml:"date,omitempty" require:"true" type:"Struct"`
}

func (s QueryP2PBandwidthResponseProvider) String() string {
  return tea.Prettify(s)
}

func (s QueryP2PBandwidthResponseProvider) GoString() string {
  return s.String()
}

func (s *QueryP2PBandwidthResponseProvider) SetName(v string) *QueryP2PBandwidthResponseProvider {
  s.Name = &v
  return s
}

func (s *QueryP2PBandwidthResponseProvider) SetType(v string) *QueryP2PBandwidthResponseProvider {
  s.Type = &v
  return s
}

func (s *QueryP2PBandwidthResponseProvider) SetDate(v *QueryP2PBandwidthResponseProviderDate) *QueryP2PBandwidthResponseProvider {
  s.Date = v
  return s
}

type QueryP2PBandwidthResponseProviderDate struct {
  // {'en':'startdate', 'zh_CN':'开始时间'}
  Startdate *string `json:"startdate,omitempty" xml:"startdate,omitempty" require:"true"`
  // {'en':'enddate', 'zh_CN':'结束时间'}
  Enddate *string `json:"enddate,omitempty" xml:"enddate,omitempty" require:"true"`
  // {'en':'channel', 'zh_CN':'频道'}
  Channel *QueryP2PBandwidthResponseProviderDateChannel `json:"channel,omitempty" xml:"channel,omitempty" require:"true" type:"Struct"`
}

func (s QueryP2PBandwidthResponseProviderDate) String() string {
  return tea.Prettify(s)
}

func (s QueryP2PBandwidthResponseProviderDate) GoString() string {
  return s.String()
}

func (s *QueryP2PBandwidthResponseProviderDate) SetStartdate(v string) *QueryP2PBandwidthResponseProviderDate {
  s.Startdate = &v
  return s
}

func (s *QueryP2PBandwidthResponseProviderDate) SetEnddate(v string) *QueryP2PBandwidthResponseProviderDate {
  s.Enddate = &v
  return s
}

func (s *QueryP2PBandwidthResponseProviderDate) SetChannel(v *QueryP2PBandwidthResponseProviderDateChannel) *QueryP2PBandwidthResponseProviderDate {
  s.Channel = v
  return s
}

type QueryP2PBandwidthResponseProviderDateChannel struct {
  // {'en':'channel', 'zh_CN':'频道'}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {'en':'bandwidth', 'zh_CN':'p2p带宽数据'}
  Bandwidth []*QueryP2PBandwidthResponseProviderDateChannelBandwidth `json:"bandwidth,omitempty" xml:"bandwidth,omitempty" require:"true" type:"Repeated"`
}

func (s QueryP2PBandwidthResponseProviderDateChannel) String() string {
  return tea.Prettify(s)
}

func (s QueryP2PBandwidthResponseProviderDateChannel) GoString() string {
  return s.String()
}

func (s *QueryP2PBandwidthResponseProviderDateChannel) SetName(v string) *QueryP2PBandwidthResponseProviderDateChannel {
  s.Name = &v
  return s
}

func (s *QueryP2PBandwidthResponseProviderDateChannel) SetBandwidth(v []*QueryP2PBandwidthResponseProviderDateChannelBandwidth) *QueryP2PBandwidthResponseProviderDateChannel {
  s.Bandwidth = v
  return s
}

type QueryP2PBandwidthResponseProviderDateChannelBandwidth struct     {
  // {'en':'time', 'zh_CN':'时间点，格式 yyyy-MM-dd hh:mm:ss'}
  Time *string `json:"time,omitempty" xml:"time,omitempty" require:"true"`
  // {'en':'p2pBandWidth', 'zh_CN':'cdn带宽，单位 Mbps'}
  P2p *string `json:"p2p,omitempty" xml:"p2p,omitempty" require:"true"`
  // {'en':'cdnBandWidth', 'zh_CN':'cdn带宽，单位 Mbps'}
  Cdn *string `json:"cdn,omitempty" xml:"cdn,omitempty" require:"true"`
  // {'en':'boxBandWidth', 'zh_CN':'p2sp带宽，单位 Mbps'}
  Box *string `json:"box,omitempty" xml:"box,omitempty" require:"true"`
  // {'en':'totalBandWidth', 'zh_CN':'总带宽，单位 Mbps'}
  Total *string `json:"total,omitempty" xml:"total,omitempty" require:"true"`
}

func (s QueryP2PBandwidthResponseProviderDateChannelBandwidth) String() string {
  return tea.Prettify(s)
}

func (s QueryP2PBandwidthResponseProviderDateChannelBandwidth) GoString() string {
  return s.String()
}

func (s *QueryP2PBandwidthResponseProviderDateChannelBandwidth) SetTime(v string) *QueryP2PBandwidthResponseProviderDateChannelBandwidth {
  s.Time = &v
  return s
}

func (s *QueryP2PBandwidthResponseProviderDateChannelBandwidth) SetP2p(v string) *QueryP2PBandwidthResponseProviderDateChannelBandwidth {
  s.P2p = &v
  return s
}

func (s *QueryP2PBandwidthResponseProviderDateChannelBandwidth) SetCdn(v string) *QueryP2PBandwidthResponseProviderDateChannelBandwidth {
  s.Cdn = &v
  return s
}

func (s *QueryP2PBandwidthResponseProviderDateChannelBandwidth) SetBox(v string) *QueryP2PBandwidthResponseProviderDateChannelBandwidth {
  s.Box = &v
  return s
}

func (s *QueryP2PBandwidthResponseProviderDateChannelBandwidth) SetTotal(v string) *QueryP2PBandwidthResponseProviderDateChannelBandwidth {
  s.Total = &v
  return s
}

type QueryP2PBandwidthPaths struct {
}

func (s QueryP2PBandwidthPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryP2PBandwidthPaths) GoString() string {
  return s.String()
}

type QueryP2PBandwidthParameters struct {
}

func (s QueryP2PBandwidthParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryP2PBandwidthParameters) GoString() string {
  return s.String()
}

type QueryP2PBandwidthRequestHeader struct {
}

func (s QueryP2PBandwidthRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryP2PBandwidthRequestHeader) GoString() string {
  return s.String()
}

type QueryP2PBandwidthResponseHeader struct {
}

func (s QueryP2PBandwidthResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryP2PBandwidthResponseHeader) GoString() string {
  return s.String()
}




type ReportCountryServerBandwidthServiceRequest struct {
  // {"en":"Start time:
  // 
  // 1. The format is yyyyy-MM-ddTHH: mm: SS + 08:00, for example, 2016-12-02T10:00 + 08:00 (10:0:00 Beijing time on December 2, 2016);
  // 
  // 2. Can not exceed the current time;
  // 
  // 3. The latest half year (183 days) data can be obtained at most.", "zh_CN":"开始时间:
  // 1.格式为yyyy-MM-ddTHH:mm:ss+08:00,例如,2016-12-02T10:00:00+08:00(为北京时间2016年12月2日10点0分0秒);
  // 2.不能大于当前时间;
  // 3.最多可获取最近半年(183天)的数据。"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty"`
  // {"en":"End time:
  // 
  // 1. The 1format is yyyy-MM-ddTHH:mm:ss+08:00;
  // 
  // 2. The end time is greater than the start time.
  // 
  // 3. If the end time is greater than the current time, the current time is taken.
  // 
  // 4. DateFrom and dateTo are not uploaded, defaulting to query the past 24 hours; if only one is not uploaded, throw an exception;
  // 
  // 5. Maximum query interval allowed: 7 days, that is, the difference between dateFrom and dateTo can not exceed 7 days.  ", "zh_CN":"结束时间:
  // 1.格式为yyyy-MM-ddTHH:mm:ss+08:00;
  // 2.结束时间需大于开始时间;
  // 3.结束时间如果大于当前时间,取当前时间;
  // 4.dateFrom,dateTo二者都未传,默认查询过去的24小时;如仅有一个未传,抛异常;
  // 5.允许查询最大间隔:7天,即dateFrom和dateTo相差不能超过7天。"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty"`
  // {"en":"Domains:
  // 
  // 1.Domain is not uploaded: Query all domain names of the account (More than 20 domains will error,you can contact technical support for adjustment);
  // 
  // 2.Domain is uploaded: Up to 20 domains are supported(you can contact technical support for adjustment).", "zh_CN":"域名:
  // 
  // 1.未传递domain时:查询账号下所有全部域名(域名超过20个则报错,可联系技术支持调整);
  // 
  // 2.有传递domain时:域名最多支持传20个(可联系技术支持调整)。"}
  Domain []*string `json:"domain,omitempty" xml:"domain,omitempty" type:"Repeated"`
  // {"en":"Country area:
  // 
  // 1. countryCode is not uploaded: Query all country areas by default;
  // 
  // 2. countryCode is uploaded: Multiple can be uploaded, such as cn, in.  Please refer to the appendix description section of the overview page.", "zh_CN":"国家区域(含中国台湾、中国澳门、中国香港、中国大陆):
  // 
  // 1.未传递countryCode时:查询全部国家区域;
  // 
  // 2.有传递countryCode时:可传多个,如cn,in。可传递的值详见概览页附录说明章节"}
  CountryCode []*string `json:"countryCode,omitempty" xml:"countryCode,omitempty" type:"Repeated"`
  // {"en":"Data granularity:
  // 
  // 5m: 5 minute granularity; Default value is 5m.
  // 
  // 1h: 1 hour granularity.", "zh_CN":"数据粒度:
  // 
  // 5m:5分钟粒度。不传默认5分钟粒度
  // 
  // 1h:1小时粒度;"}
  DataInterval *string `json:"dataInterval,omitempty" xml:"dataInterval,omitempty"`
}

func (s ReportCountryServerBandwidthServiceRequest) String() string {
  return tea.Prettify(s)
}

func (s ReportCountryServerBandwidthServiceRequest) GoString() string {
  return s.String()
}

func (s *ReportCountryServerBandwidthServiceRequest) SetDateFrom(v string) *ReportCountryServerBandwidthServiceRequest {
  s.DateFrom = &v
  return s
}

func (s *ReportCountryServerBandwidthServiceRequest) SetDateTo(v string) *ReportCountryServerBandwidthServiceRequest {
  s.DateTo = &v
  return s
}

func (s *ReportCountryServerBandwidthServiceRequest) SetDomain(v []*string) *ReportCountryServerBandwidthServiceRequest {
  s.Domain = v
  return s
}

func (s *ReportCountryServerBandwidthServiceRequest) SetCountryCode(v []*string) *ReportCountryServerBandwidthServiceRequest {
  s.CountryCode = v
  return s
}

func (s *ReportCountryServerBandwidthServiceRequest) SetDataInterval(v string) *ReportCountryServerBandwidthServiceRequest {
  s.DataInterval = &v
  return s
}

type ReportCountryServerBandwidthServiceResponse struct {
  // {"en":"Domain", "zh_CN":"域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  CountryData []*ReportCountryServerBandwidthServiceResponseCountryData `json:"countryData,omitempty" xml:"countryData,omitempty" require:"true" type:"Repeated"`
}

func (s ReportCountryServerBandwidthServiceResponse) String() string {
  return tea.Prettify(s)
}

func (s ReportCountryServerBandwidthServiceResponse) GoString() string {
  return s.String()
}

func (s *ReportCountryServerBandwidthServiceResponse) SetDomain(v string) *ReportCountryServerBandwidthServiceResponse {
  s.Domain = &v
  return s
}

func (s *ReportCountryServerBandwidthServiceResponse) SetCountryData(v []*ReportCountryServerBandwidthServiceResponseCountryData) *ReportCountryServerBandwidthServiceResponse {
  s.CountryData = v
  return s
}

type ReportCountryServerBandwidthServiceResponseCountryData struct     {
  // {"en":"Country area", "zh_CN":"国家地区"}
  CountryCode *string `json:"countryCode,omitempty" xml:"countryCode,omitempty" require:"true"`
  DetailList []*ReportCountryServerBandwidthServiceResponseCountryDataDetailList `json:"detailList,omitempty" xml:"detailList,omitempty" require:"true" type:"Repeated"`
}

func (s ReportCountryServerBandwidthServiceResponseCountryData) String() string {
  return tea.Prettify(s)
}

func (s ReportCountryServerBandwidthServiceResponseCountryData) GoString() string {
  return s.String()
}

func (s *ReportCountryServerBandwidthServiceResponseCountryData) SetCountryCode(v string) *ReportCountryServerBandwidthServiceResponseCountryData {
  s.CountryCode = &v
  return s
}

func (s *ReportCountryServerBandwidthServiceResponseCountryData) SetDetailList(v []*ReportCountryServerBandwidthServiceResponseCountryDataDetailList) *ReportCountryServerBandwidthServiceResponseCountryData {
  s.DetailList = v
  return s
}

type ReportCountryServerBandwidthServiceResponseCountryDataDetailList struct     {
  // {"en":"Time:
  // 
  // 1. When the data query granularity is 5m, then the format is yyyy-MM-dd HH:mm; ach time slice value represents the value within the previous time granularity range. The first time slice of the day is yyyy-MM-dd 00:05, and the last one is (yyyy-MM-dd+1) 00:00;
  // 
  // 2. When the data query granularity is 1h, the format is yyyy-MM-dd HH; Each time slice value represents the value within the previous time granularity range. The first time slice of the day is yyyy-MM-dd 01, and the last one is (yyyy-MM-dd+1) 00;
  // 
  // 3. Return the time slices that contained in start time and in end time.", "zh_CN":"时间,
  // 查询的数据粒度为5m时,格式为yyyy-MM-dd HH:mm;每一个时间片数据值代表的是前一个时间粒度范围内的数据值。一天开始的时间片是yyyy-MM-dd 00:05,最后一个时间片是(yyyy-MM-dd+1) 00:00;
  // 查询的数据粒度为1h时,格式为yyyy-MM-dd HH;每一个时间片数据值代表的是前一个时间粒度范围内的数据值。一天开始的时间片是yyyy-MM-dd 01,最后一个时间片是(yyyy-MM-dd+1) 00;
  // 返回开始时间和结束时间包含的时间片。"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {"en":"Bandwidth value, unit Mbps", "zh_CN":"时间点对应的边缘带宽数据,单位Mbps"}
  Value *string `json:"value,omitempty" xml:"value,omitempty" require:"true"`
  // {"en":"Flow, unit MB", "zh_CN":"时间点对应的边缘流量,单位MB"}
  Flow *string `json:"flow,omitempty" xml:"flow,omitempty" require:"true"`
}

func (s ReportCountryServerBandwidthServiceResponseCountryDataDetailList) String() string {
  return tea.Prettify(s)
}

func (s ReportCountryServerBandwidthServiceResponseCountryDataDetailList) GoString() string {
  return s.String()
}

func (s *ReportCountryServerBandwidthServiceResponseCountryDataDetailList) SetTimestamp(v string) *ReportCountryServerBandwidthServiceResponseCountryDataDetailList {
  s.Timestamp = &v
  return s
}

func (s *ReportCountryServerBandwidthServiceResponseCountryDataDetailList) SetValue(v string) *ReportCountryServerBandwidthServiceResponseCountryDataDetailList {
  s.Value = &v
  return s
}

func (s *ReportCountryServerBandwidthServiceResponseCountryDataDetailList) SetFlow(v string) *ReportCountryServerBandwidthServiceResponseCountryDataDetailList {
  s.Flow = &v
  return s
}

type ReportCountryServerBandwidthServicePaths struct {
}

func (s ReportCountryServerBandwidthServicePaths) String() string {
  return tea.Prettify(s)
}

func (s ReportCountryServerBandwidthServicePaths) GoString() string {
  return s.String()
}

type ReportCountryServerBandwidthServiceParameters struct {
}

func (s ReportCountryServerBandwidthServiceParameters) String() string {
  return tea.Prettify(s)
}

func (s ReportCountryServerBandwidthServiceParameters) GoString() string {
  return s.String()
}

type ReportCountryServerBandwidthServiceRequestHeader struct {
}

func (s ReportCountryServerBandwidthServiceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportCountryServerBandwidthServiceRequestHeader) GoString() string {
  return s.String()
}

type ReportCountryServerBandwidthServiceResponseHeader struct {
}

func (s ReportCountryServerBandwidthServiceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportCountryServerBandwidthServiceResponseHeader) GoString() string {
  return s.String()
}




type QueryMultiDomainsIPV6OrIPV4BandwidthRequest struct {
  // {"en":"Start time: 
  //     1. Time format is yyyy-MM-ddTHH:mm:ss+08:00, for example, 2016-12-02T10:00:00+08:00 (10:00 on 2nd of December 2016, Beijing Time); 
  //     2. No bigger than the current time. 
  //     3. QueryMultiDomainsIPV6OrIPV4BandwidthData in the last 183 days at most can be queried.", "zh_CN":"开始时间:
  // 1.时间格式为yyyy-MM-ddTHH:mm:ss+08:00,例如,2016-12-02T10:00:00+08:00(为北京时间2016年12月2日10点0分0秒);
  // 2.不能大于当前时间
  // 3.最多可获取最近半年(183天)的数据。"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty"`
  // {"en":"End time: 
  //     1. the time format is 2016-12-02T10:00:00+08:00 
  //     2. End time should be greater than start time. If the end time is greater than current time, current time will be used. 
  //     3. If both fields of dataFrom and dateTo are left empty, then data in the last 24 hours will be queried by default; if only one field is filled in and one is left empty, then exception will be occur. 
  //     4. Allowable maximum time range for query:7 day, means the period between dateFrom to dateTo should not exceed 7 day (can be adjusted by contacting technical support).", "zh_CN":"结束时间:
  // 1.时间格式yyyy-MM-ddTHH:mm:ss+08:00
  // 2.结束时间需大于开始时间,结束时间如果大于当前时间,取当前时间。
  // 3.dateFrom,dateTo二者都未传,默认查询过去的24小时;如仅有一个未传,抛异常
  // 4.允许查询最大时间间隔:7天,即dateFrom和dateTo相差不能超过7天。(可联系技术支持调整)"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty"`
  // {"en":"Domain: 
  //     1. The maximum number of domian is 200 by default (Technical Support Adjustment can be contacted); 
  //     2. automatic filtering invalid domain name (if pass illegal domain name, can be filtered, query result only returns the data of valid domain name).", "zh_CN":"域名:
  // 
  // 1.可传递域名数量上限默认为200个(可联系技术支持调整);
  // 2.自动过滤掉无效域名(如传递非法域名,会被过滤掉,查询结果只返回有效域名的数据)。"}
  Domain []*string `json:"domain,omitempty" xml:"domain,omitempty" type:"Repeated"`
  // {"en":"QueryMultiDomainsIPV6OrIPV4BandwidthData granularity 
  //     1. Support 5m (5 minute granularity), 1h (1 hour granularity); 
  //     2. The default is 5m", "zh_CN":"数据粒度:
  // 1.支持5m(5分钟)、1h(1小时)
  // 2.不传默认5m。"}
  DataInterval *string `json:"dataInterval,omitempty" xml:"dataInterval,omitempty"`
  // {"en":"IP type:
  //     1.The optional values are IPv6 and IPv4 
  //     2.If let this parameter empty,it will query all IP type;", "zh_CN":"IP类型:
  // 1.可选值为 IPV6、IPV4
  // 2.不传默认查询全部"}
  IpType *string `json:"ipType,omitempty" xml:"ipType,omitempty"`
  // {"en":"1.If let this parameter empty,it will query all rigion;
  // 	2.Optional values cn, am, apac, au, emea, hk, tw, etc", "zh_CN":"1. 不传查全部;
  // 	2. 可选值 cn、am、apac、au、emea、hk、tw 等"}
  Region []*string `json:"region,omitempty" xml:"region,omitempty" type:"Repeated"`
}

func (s QueryMultiDomainsIPV6OrIPV4BandwidthRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryMultiDomainsIPV6OrIPV4BandwidthRequest) GoString() string {
  return s.String()
}

func (s *QueryMultiDomainsIPV6OrIPV4BandwidthRequest) SetDateFrom(v string) *QueryMultiDomainsIPV6OrIPV4BandwidthRequest {
  s.DateFrom = &v
  return s
}

func (s *QueryMultiDomainsIPV6OrIPV4BandwidthRequest) SetDateTo(v string) *QueryMultiDomainsIPV6OrIPV4BandwidthRequest {
  s.DateTo = &v
  return s
}

func (s *QueryMultiDomainsIPV6OrIPV4BandwidthRequest) SetDomain(v []*string) *QueryMultiDomainsIPV6OrIPV4BandwidthRequest {
  s.Domain = v
  return s
}

func (s *QueryMultiDomainsIPV6OrIPV4BandwidthRequest) SetDataInterval(v string) *QueryMultiDomainsIPV6OrIPV4BandwidthRequest {
  s.DataInterval = &v
  return s
}

func (s *QueryMultiDomainsIPV6OrIPV4BandwidthRequest) SetIpType(v string) *QueryMultiDomainsIPV6OrIPV4BandwidthRequest {
  s.IpType = &v
  return s
}

func (s *QueryMultiDomainsIPV6OrIPV4BandwidthRequest) SetRegion(v []*string) *QueryMultiDomainsIPV6OrIPV4BandwidthRequest {
  s.Region = v
  return s
}

type QueryMultiDomainsIPV6OrIPV4BandwidthResponse struct {
  // {"en":"The result status code of this request", "zh_CN":"请求结果状态码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"The result info of this request", "zh_CN":"请求结果信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Detailed data on the result of the request", "zh_CN":"请求结果的详细数据"}
  QueryMultiDomainsIPV6OrIPV4BandwidthData []*QueryMultiDomainsIPV6OrIPV4BandwidthData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s QueryMultiDomainsIPV6OrIPV4BandwidthResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryMultiDomainsIPV6OrIPV4BandwidthResponse) GoString() string {
  return s.String()
}

func (s *QueryMultiDomainsIPV6OrIPV4BandwidthResponse) SetCode(v string) *QueryMultiDomainsIPV6OrIPV4BandwidthResponse {
  s.Code = &v
  return s
}

func (s *QueryMultiDomainsIPV6OrIPV4BandwidthResponse) SetMessage(v string) *QueryMultiDomainsIPV6OrIPV4BandwidthResponse {
  s.Message = &v
  return s
}

func (s *QueryMultiDomainsIPV6OrIPV4BandwidthResponse) SetData(v []*QueryMultiDomainsIPV6OrIPV4BandwidthData) *QueryMultiDomainsIPV6OrIPV4BandwidthResponse {
  s.QueryMultiDomainsIPV6OrIPV4BandwidthData = v
  return s
}

type QueryMultiDomainsIPV6OrIPV4BandwidthData struct {
  // {"en":"domain", "zh_CN":"域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  DetailList []*QueryMultiDomainsIPV6OrIPV4BandwidthDataDetailList `json:"detailList,omitempty" xml:"detailList,omitempty" require:"true" type:"Repeated"`
}

func (s QueryMultiDomainsIPV6OrIPV4BandwidthData) String() string {
  return tea.Prettify(s)
}

func (s QueryMultiDomainsIPV6OrIPV4BandwidthData) GoString() string {
  return s.String()
}

func (s *QueryMultiDomainsIPV6OrIPV4BandwidthData) SetDomain(v string) *QueryMultiDomainsIPV6OrIPV4BandwidthData {
  s.Domain = &v
  return s
}

func (s *QueryMultiDomainsIPV6OrIPV4BandwidthData) SetDetailList(v []*QueryMultiDomainsIPV6OrIPV4BandwidthDataDetailList) *QueryMultiDomainsIPV6OrIPV4BandwidthData {
  s.DetailList = v
  return s
}

type QueryMultiDomainsIPV6OrIPV4BandwidthDataDetailList struct     {
  // {"en":"Timestamp: The timestamp between the start time and the end time.", "zh_CN":"时间片: 返回开始时间和结束时间包含的时间片。"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {"en":"Bandwidth,Mbps", "zh_CN":"带宽,Mbps"}
  Value *string `json:"value,omitempty" xml:"value,omitempty" require:"true"`
}

func (s QueryMultiDomainsIPV6OrIPV4BandwidthDataDetailList) String() string {
  return tea.Prettify(s)
}

func (s QueryMultiDomainsIPV6OrIPV4BandwidthDataDetailList) GoString() string {
  return s.String()
}

func (s *QueryMultiDomainsIPV6OrIPV4BandwidthDataDetailList) SetTimestamp(v string) *QueryMultiDomainsIPV6OrIPV4BandwidthDataDetailList {
  s.Timestamp = &v
  return s
}

func (s *QueryMultiDomainsIPV6OrIPV4BandwidthDataDetailList) SetValue(v string) *QueryMultiDomainsIPV6OrIPV4BandwidthDataDetailList {
  s.Value = &v
  return s
}

type QueryMultiDomainsIPV6OrIPV4BandwidthPaths struct {
}

func (s QueryMultiDomainsIPV6OrIPV4BandwidthPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryMultiDomainsIPV6OrIPV4BandwidthPaths) GoString() string {
  return s.String()
}

type QueryMultiDomainsIPV6OrIPV4BandwidthParameters struct {
}

func (s QueryMultiDomainsIPV6OrIPV4BandwidthParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryMultiDomainsIPV6OrIPV4BandwidthParameters) GoString() string {
  return s.String()
}

type QueryMultiDomainsIPV6OrIPV4BandwidthRequestHeader struct {
}

func (s QueryMultiDomainsIPV6OrIPV4BandwidthRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryMultiDomainsIPV6OrIPV4BandwidthRequestHeader) GoString() string {
  return s.String()
}

type QueryMultiDomainsIPV6OrIPV4BandwidthResponseHeader struct {
}

func (s QueryMultiDomainsIPV6OrIPV4BandwidthResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryMultiDomainsIPV6OrIPV4BandwidthResponseHeader) GoString() string {
  return s.String()
}




type QueryBandwidthForMultiDomainRequest struct {
  // {"en":"Domain list.
  // Domain number limits can be adjusted depending on different accounts. The default value is 20(if you want to adjust,please, contact technical support)", "zh_CN":"域名列表
  // 1.域名个数限制根据账号可调,默认为20个(可联系技术支持下单调整);"}
  QueryBandwidthForMultiDomainDomainList *QueryBandwidthForMultiDomainDomainList `json:"domain-list,omitempty" xml:"domain-list,omitempty" require:"true"`
}

func (s QueryBandwidthForMultiDomainRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryBandwidthForMultiDomainRequest) GoString() string {
  return s.String()
}

func (s *QueryBandwidthForMultiDomainRequest) SetDomainList(v *QueryBandwidthForMultiDomainDomainList) *QueryBandwidthForMultiDomainRequest {
  s.QueryBandwidthForMultiDomainDomainList = v
  return s
}

type QueryBandwidthForMultiDomainDomainList struct {
  // {"en":"Domain", "zh_CN":"域名"}
  DomainName []*string `json:"domain-name,omitempty" xml:"domain-name,omitempty" require:"true" type:"Repeated"`
}

func (s QueryBandwidthForMultiDomainDomainList) String() string {
  return tea.Prettify(s)
}

func (s QueryBandwidthForMultiDomainDomainList) GoString() string {
  return s.String()
}

func (s *QueryBandwidthForMultiDomainDomainList) SetDomainName(v []*string) *QueryBandwidthForMultiDomainDomainList {
  s.DomainName = v
  return s
}

type QueryBandwidthForMultiDomainResponse struct {
  BandwidthReport []*QueryBandwidthForMultiDomainResponseBandwidthReport `json:"bandwidthReport,omitempty" xml:"bandwidthReport,omitempty" require:"true" type:"Repeated"`
}

func (s QueryBandwidthForMultiDomainResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryBandwidthForMultiDomainResponse) GoString() string {
  return s.String()
}

func (s *QueryBandwidthForMultiDomainResponse) SetBandwidthReport(v []*QueryBandwidthForMultiDomainResponseBandwidthReport) *QueryBandwidthForMultiDomainResponse {
  s.BandwidthReport = v
  return s
}

type QueryBandwidthForMultiDomainResponseBandwidthReport struct     {
  // {"en":"Date:
  // 1.When the querying data granularity is fiveminutes, the format is yyyy-MM-dd HH:mm; the data value of every time slice represents the data value within the previous time granularity range. The first time slice of the day is yyyy-MM-dd 00:05 AM, and the last one is yyyy-MM-dd 24:00;
  // 2.When the querying data granularity is hourly, the format is yyyy-MM-dd HH; the data value of every time slice represents the data value within the previous time granularity range. The first time slice of the day is yyyy-MM-dd 01, and the last one is yyyy-MM-dd 24;
  // 3.When the querying data granularity is daily, the format is yyyy-MM-dd; the data value of every time slice represents the value of the daily data;
  // 4.Return the time slice contained in start time and in end time", "zh_CN":"时间:
  // 1.查询的数据粒度为fiveminutes时,格式为yyyy-MM-dd HH:mm;每一个时间片数据值代表的是前一个时间粒度范围内的数据值。一天开始的时间片是yyyy-MM-dd 00:05,最后一个时间片是yyyy-MM-dd 24:00。
  // 2.查询的数据粒度为hourly时,格式为yyyy-MM-dd HH;每一个时间片数据值代表的是前一个时间粒度范围内的数据值。一天开始的时间片是yyyy-MM-dd 01,最后一个时间片是yyyy-MM-dd 24。
  // 3.查询的数据粒度为daily时,格式为yyyy-MM-dd;每一个时间片数据值代表的该天内的数据值;
  // 4.返回开始时间和结束时间包含的时间片。"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {"en":"Bandwidth, keep the number to four decimal places. Unit: Mbps", "zh_CN":"带宽,保留4位小数,单位为Mbps"}
  Bandwidth *int `json:"bandwidth,omitempty" xml:"bandwidth,omitempty" require:"true"`
}

func (s QueryBandwidthForMultiDomainResponseBandwidthReport) String() string {
  return tea.Prettify(s)
}

func (s QueryBandwidthForMultiDomainResponseBandwidthReport) GoString() string {
  return s.String()
}

func (s *QueryBandwidthForMultiDomainResponseBandwidthReport) SetTimestamp(v string) *QueryBandwidthForMultiDomainResponseBandwidthReport {
  s.Timestamp = &v
  return s
}

func (s *QueryBandwidthForMultiDomainResponseBandwidthReport) SetBandwidth(v int) *QueryBandwidthForMultiDomainResponseBandwidthReport {
  s.Bandwidth = &v
  return s
}

type QueryBandwidthForMultiDomainPaths struct {
}

func (s QueryBandwidthForMultiDomainPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryBandwidthForMultiDomainPaths) GoString() string {
  return s.String()
}

type QueryBandwidthForMultiDomainParameters struct {
  // {"en":"Start time
  // 1.The format is yyyy-MM-ddTHH:mm:ss+08:00;
  // 2.And smaller than the current time and dateTo;
  // 3.Period between dataFrom and dateTo cannot be longer than 31 days", "zh_CN":"开始时间
  // 1.格式为yyyy-MM-ddTHH:mm:ss+08:00;
  // 2.并且小于当前时间和dateTo;
  // 3.dateFrom和dateTo相差不能超过31天(可联系技术支持调整);4.只能查询最近2年内数据。"}
  Datefrom *string `json:"datefrom,omitempty" xml:"datefrom,omitempty" require:"true"`
  // {"en":"End time
  // 1.The format is yyyy-MM-ddTHH:mm:ss+08:00;
  // 2.Must be greater than dateFrom;
  // 3.If it's greater than the current time, then the current time will be assigned as the value", "zh_CN":"结束时间
  // 1.格式为yyyy-MM-ddTHH:mm:ss+08:00;
  // 2.必须大于dateFrom;
  // 3.如果大于当前时间,则重新赋值为当前时间;"}
  Dateto *string `json:"dateto,omitempty" xml:"dateto,omitempty" require:"true"`
  // {"en":"Data granularity
  // fiveminutes: five minutes, hourly: one hour, daily: one day;If not specified, daily is set as the default value", "zh_CN":"数据粒度
  // 1.fiveminutes:5分钟,hourly:1小时,daily:1天;
  // 2.不传递,默认为daily;
  // 3.传递fiveminutes时,若客户数据采集粒度有特殊配置将按实际配置粒度返回。"}
  Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s QueryBandwidthForMultiDomainParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryBandwidthForMultiDomainParameters) GoString() string {
  return s.String()
}

func (s *QueryBandwidthForMultiDomainParameters) SetDatefrom(v string) *QueryBandwidthForMultiDomainParameters {
  s.Datefrom = &v
  return s
}

func (s *QueryBandwidthForMultiDomainParameters) SetDateto(v string) *QueryBandwidthForMultiDomainParameters {
  s.Dateto = &v
  return s
}

func (s *QueryBandwidthForMultiDomainParameters) SetType(v string) *QueryBandwidthForMultiDomainParameters {
  s.Type = &v
  return s
}

type QueryBandwidthForMultiDomainRequestHeader struct {
}

func (s QueryBandwidthForMultiDomainRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryBandwidthForMultiDomainRequestHeader) GoString() string {
  return s.String()
}

type QueryBandwidthForMultiDomainResponseHeader struct {
}

func (s QueryBandwidthForMultiDomainResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryBandwidthForMultiDomainResponseHeader) GoString() string {
  return s.String()
}




type QueryBandwidthbyISPProvinceRequest struct {
  // {'en':'Start time', 'zh_CN':'开始时间：
  // &nbsp; 1.格式为yyyy-MM-ddTHH:mm:ss+08:00，例如，2016-12-02T10:00:00+08:00（为北京时间2016年12月2日10点0分0秒）；
  // &nbsp; 2.不能大于当前时间；
  // &nbsp; 3.最多可获取最近半年（183天）的数据。'}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty"`
  // {'en':'End time:
  // 
  // The 1. format is yyyy-MM-ddTHH:mm:ss+08:00;
  // 
  // 2. the end time is larger than the start time.
  // 
  // 3. if the end time is greater than the current time, take the current time.
  // 
  // 4. DateFrom and dateTo are not uploaded, default query for the past 24 hours; if only one is not uploaded, throw an exception;
  // 
  // 5. Maximum query interval allowed: 31 days, that is, the difference between dateFrom and dateTo can not exceed 31 days (technical support can be contacted to adjust).', 'zh_CN':'结束时间：
  // &nbsp; 1.格式为yyyy-MM-ddTHH:mm:ss+08:00；
  // &nbsp; 2.结束时间需大于开始时间；
  // &nbsp; 3.结束时间如果大于当前时间，取当前时间；
  // &nbsp; 4.dateFrom，dateTo二者都未传，默认查询过去的24小时；如仅有一个未传，抛异常；
  // &nbsp; 5.允许查询最大间隔：31天，即dateFrom和dateTo相差不能超过31天（可联系技术支持调整）。'}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty"`
  // {'en':'Data granularity
  // 
  // 1.        fiveminutes:   five minutes, hourly: one hour, daily: one day;
  // 
  // 2.        If   not specified, daily is set as the default value;
  // 
  // 3.        If   fiveminutes is specified as the value, then data is returned in actual   configured granularity when there is specific configuration to data collecting   granularity for the customer.', 'zh_CN':'数据粒度：
  // 1.支持5m（5分钟）
  // 2.不传默认5m。'}
  DataInterval *string `json:"dataInterval,omitempty" xml:"dataInterval,omitempty"`
  // {'en':'area:
  // 
  // 1. Do not pass the default query all areas.
  // 
  // 2. Support to pass multiple regions, the optional values are: cn, hk, ov, apac, am, emea, fg.', 'zh_CN':'区域：
  // 1.不传默认查询全部区域。
  // 2.支持传递多个区域，可选值为：cn、hk、ov、apac、am、emea、fg。'}
  AreaCode []*string `json:"areaCode,omitempty" xml:"areaCode,omitempty" type:"Repeated"`
  // {'en':'domain:
  // 
  // 1. The maximum number of passable domain names is 20 by default (you can contact technical support adjustment).
  // 
  // 2. Query all the domain names under the account when the entry is not passed, but you cannot query (error) when the number of domain names under the account exceeds the limit.', 'zh_CN':'域名：
  // 1.可传递域名数量上限默认为20个（可联系技术支持调整），
  // 2.未传递该入参时查询账号下所有域名，但当账号下域名数量超过限制时不可查询（报错）。'}
  Domain []*string `json:"domain,omitempty" xml:"domain,omitempty" type:"Repeated"`
  // {'en':'By default, all provinces are queried, and the provinces are transferred. The province information code table is detailed in the appendix of the overview page.', 'zh_CN':'默认查询全部省份，传递省份，省份信息码表详见概览页附录说明章节'}
  Province []*string `json:"province,omitempty" xml:"province,omitempty" type:"Repeated"`
  // {'en':'The default query is for all operators, the operators and operator information codes are detailed in the appendix of the overview page.', 'zh_CN':'默认查询全部运营商，传递运营商，运营商信息码表详见概览页附录说明章节'}
  Isp []*string `json:"isp,omitempty" xml:"isp,omitempty" type:"Repeated"`
}

func (s QueryBandwidthbyISPProvinceRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryBandwidthbyISPProvinceRequest) GoString() string {
  return s.String()
}

func (s *QueryBandwidthbyISPProvinceRequest) SetDateFrom(v string) *QueryBandwidthbyISPProvinceRequest {
  s.DateFrom = &v
  return s
}

func (s *QueryBandwidthbyISPProvinceRequest) SetDateTo(v string) *QueryBandwidthbyISPProvinceRequest {
  s.DateTo = &v
  return s
}

func (s *QueryBandwidthbyISPProvinceRequest) SetDataInterval(v string) *QueryBandwidthbyISPProvinceRequest {
  s.DataInterval = &v
  return s
}

func (s *QueryBandwidthbyISPProvinceRequest) SetAreaCode(v []*string) *QueryBandwidthbyISPProvinceRequest {
  s.AreaCode = v
  return s
}

func (s *QueryBandwidthbyISPProvinceRequest) SetDomain(v []*string) *QueryBandwidthbyISPProvinceRequest {
  s.Domain = v
  return s
}

func (s *QueryBandwidthbyISPProvinceRequest) SetProvince(v []*string) *QueryBandwidthbyISPProvinceRequest {
  s.Province = v
  return s
}

func (s *QueryBandwidthbyISPProvinceRequest) SetIsp(v []*string) *QueryBandwidthbyISPProvinceRequest {
  s.Isp = v
  return s
}

type QueryBandwidthbyISPProvinceResponse struct {
  // {'en':'result', 'zh_CN':'结果'}
  Result []*QueryBandwidthbyISPProvinceResponseResult `json:"result,omitempty" xml:"result,omitempty" require:"true" type:"Repeated"`
}

func (s QueryBandwidthbyISPProvinceResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryBandwidthbyISPProvinceResponse) GoString() string {
  return s.String()
}

func (s *QueryBandwidthbyISPProvinceResponse) SetResult(v []*QueryBandwidthbyISPProvinceResponseResult) *QueryBandwidthbyISPProvinceResponse {
  s.Result = v
  return s
}

type QueryBandwidthbyISPProvinceResponseResult struct     {
  BandwidthData []*QueryBandwidthbyISPProvinceResponseResultBandwidthData `json:"bandwidthData,omitempty" xml:"bandwidthData,omitempty" require:"true" type:"Repeated"`
}

func (s QueryBandwidthbyISPProvinceResponseResult) String() string {
  return tea.Prettify(s)
}

func (s QueryBandwidthbyISPProvinceResponseResult) GoString() string {
  return s.String()
}

func (s *QueryBandwidthbyISPProvinceResponseResult) SetBandwidthData(v []*QueryBandwidthbyISPProvinceResponseResultBandwidthData) *QueryBandwidthbyISPProvinceResponseResult {
  s.BandwidthData = v
  return s
}

type QueryBandwidthbyISPProvinceResponseResultBandwidthData struct     {
  // {'en':'Date
  // 
  // 1.        When   the data query granularity is fiveminutes, the format is yyyy-MM-dd HH:mm;   the data value of every time slice represents the data value within the   previous time granularity range. 
  // 2.        Return   the time slice contained in start time and the time slice contained in end   time.', 'zh_CN':'时间
  // 1.查询的数据粒度为5m时，格式为yyyy-MM-dd &nbsp; HH:mm；每一个时间片数据值代表的是前一个时间粒度范围内的数据值，比如yyyy-MM-dd 00:05，代表00:00到00:05范围内的数据。
  // 2.返回开始时间和结束时间包含的时间片。'}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {'en':'The bandwidth value, the corresponding value value of each time slice, shows the cumulative data value represented by this time slice.
  // 
  // 
  // 
  // Unit Mbps, keep 2 decimal digits.', 'zh_CN':'带宽值，每个时间片的对应value值都展示这个时间片代表的完整数据累计值。
  // 单位Mbps，保留2位小数。'}
  Value *string `json:"value,omitempty" xml:"value,omitempty" require:"true"`
}

func (s QueryBandwidthbyISPProvinceResponseResultBandwidthData) String() string {
  return tea.Prettify(s)
}

func (s QueryBandwidthbyISPProvinceResponseResultBandwidthData) GoString() string {
  return s.String()
}

func (s *QueryBandwidthbyISPProvinceResponseResultBandwidthData) SetTimestamp(v string) *QueryBandwidthbyISPProvinceResponseResultBandwidthData {
  s.Timestamp = &v
  return s
}

func (s *QueryBandwidthbyISPProvinceResponseResultBandwidthData) SetValue(v string) *QueryBandwidthbyISPProvinceResponseResultBandwidthData {
  s.Value = &v
  return s
}

type QueryBandwidthbyISPProvincePaths struct {
}

func (s QueryBandwidthbyISPProvincePaths) String() string {
  return tea.Prettify(s)
}

func (s QueryBandwidthbyISPProvincePaths) GoString() string {
  return s.String()
}

type QueryBandwidthbyISPProvinceParameters struct {
}

func (s QueryBandwidthbyISPProvinceParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryBandwidthbyISPProvinceParameters) GoString() string {
  return s.String()
}

type QueryBandwidthbyISPProvinceRequestHeader struct {
}

func (s QueryBandwidthbyISPProvinceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryBandwidthbyISPProvinceRequestHeader) GoString() string {
  return s.String()
}

type QueryBandwidthbyISPProvinceResponseHeader struct {
}

func (s QueryBandwidthbyISPProvinceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryBandwidthbyISPProvinceResponseHeader) GoString() string {
  return s.String()
}




type QueryIPV6BandwidthOfeachISPandProvinceRequest struct {
  // {"en":"Start time
  // 
  // 1.The time format is yyyy-MM-ddTHH:mm:ss+08:00, for example, 2016-12-02T10:00:00+08:00 (for Beijing time, December 2, 2016, 10:00:00);
  // 2.Cannot be greater than the current time
  // 3.Get up to the last six months (183 days) of data.", "zh_CN":"开始时间:
  // 
  // 1.时间格式为yyyy-MM-ddTHH:mm:ss+08:00,例如,2016-12-02T10:00:00+08:00(为北京时间2016年12月2日10点0分0秒);
  // 
  // 2.不能大于当前时间
  // 
  // 3.最多可获取最近半年(183天)的数据"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty"`
  // {"en":"End time
  // 1.The format is yyyy-MM-ddTHH:mm:ss+08:00;
  // 2.Must be greater than dateFrom; if it's greater than the current time, then the current time is assigned as the value;
  // 3.If both fields of dataFrom and dateTo are left empty, then data in the last 24 hours will be queried by default;
  // 4.Allowable maximum time range for query: 1 day, means the period between dateFrom to dateTo should not exceed 1 day (can be adjusted by contacting technical support up to 31 days)", "zh_CN":"结束时间:
  // 
  // 1.时间格式yyyy-MM-ddTHH:mm:ss+08:00
  // 
  // 2.结束时间需大于开始时间,结束时间如果大于当前时间,取当前时间。
  // 
  // 3.dateFrom,dateTo二者都未传,默认查询过去的24小时;如仅有一个未传,抛异常
  // 
  // 4.允许查询最大时间间隔:1天,即dateFrom和dateTo相差不能超过1天。(可联系技术支持调整)"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty"`
  // {"en":"Domain:
  // 1.The default upper limit to domains that can be entered is 200 (Contact technical support to adjust);
  // 2.Automatically filter out illegal domains (illegal domains will be filtered out, the query results only return the data of legitimate domains)", "zh_CN":"域名:
  // 
  // 1.可传递域名数量上限默认为200个(可联系技术支持调整);
  // 
  // 2.自动过滤掉无效域名(如传递非法域名,会被过滤掉,查询结果只返回有效域名的数据)"}
  Domain []*string `json:"domain,omitempty" xml:"domain,omitempty" type:"Repeated"`
  // {"en":"Data granularity:
  // 1. Support 5m (5 minutes granularity),1h (1 hour granularity)
  // 2. The default value is 5m", "zh_CN":"数据粒度:
  // 
  // 1.支持5m(5分钟)、1h(1小时)
  // 
  // 2.不传默认5m。"}
  DataInterval *string `json:"dataInterval,omitempty" xml:"dataInterval,omitempty"`
  // {"en":"Province
  // 
  // 1.Province is not upload: Query all provinces and aggregate the returned data according to all provinces; 
  // 2.Province is upload: Provinces can transmit Chinese or code. Please refer to the appendix description section of the overview page for the provincial information code table.
  // 
  // 3.Support language request header Accept Language, only support zh-CN and en-US, default to zh-CN. Accept Language: en-US, both the province and isp input and return are in code, otherwise the return is in Chinese.", "zh_CN":"省份
  // 
  // 1.未传递province时：查询所有省份，返回的数据按照所有省份聚合。
  // 
  // 2.有传递province时：省份 可传中文或code。省份信息码表详见概览页附录说明章节
  // 
  // 3.支持语言请求头Accept-Language，只支持zh-CN、en-US，默认为zh-CN。Accept-Language：en-US时，省份及运营商 入参及返回都为code，否则返回的为中文。"}
  Province []*string `json:"province,omitempty" xml:"province,omitempty" type:"Repeated"`
  // {"en":"ISP:
  // 1.ISP is not upload: Query all ISPs and aggregate the returned data according to all ISPs; 
  // 2.ISPs is upload: Isp can transmit Chinese or code. Please refer to the appendix description section of the overview page for the ISP information code table.", "zh_CN":"运营商：
  // 1.未传递isp时：查询所有isp，返回的数据按照所有运营商聚合。 
  // 2.有传递isp时：运营商 可传中文或code。运营商信息码表详见概览页附录说明章节"}
  Isp []*string `json:"isp,omitempty" xml:"isp,omitempty" type:"Repeated"`
  // {"en":"IP type
  // 1.Optional values:IPV6,IPV4
  // 2.Query all IP type by default.", "zh_CN":"IP类型:
  // 
  // 1.可选值为 IPV6.IPV4
  // 2.不传默认查询全部"}
  IPType *string `json:"IPType,omitempty" xml:"IPType,omitempty"`
  // {"en":"Group dimension:
  // 1.Default response aggregation data without group
  // 2.Optional values:domain,province,isp,you can pass in single or multiple values
  // 3.The detailed data will be displayed according to the dimension.For example,if the dimension is isp,the detail data will be group by each isp.", "zh_CN":"分组关键词:
  // 1.默认聚合展示;
  // 2.可选值为domain.province.isp,可传入多个值;
  // 3.传入关键词则代表需要按照关键词对应的值展示明细; 例如groupBy传入isp,则isp维度需要明细展示;当没有传递isp,则代表isp聚合展示,同时isp节点则不返回。其他province和domain相同逻辑。 例如:传递'groupBy':   ['domain','province'],则ispData下的isp节点无需返回。 { 'domain': 'www.aaaa.com', 'ispData': [ { 'isp':   '中国电信', 'provinceData': [....] }]}"}
  GroupBy []*string `json:"groupBy,omitempty" xml:"groupBy,omitempty" type:"Repeated"`
}

func (s QueryIPV6BandwidthOfeachISPandProvinceRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryIPV6BandwidthOfeachISPandProvinceRequest) GoString() string {
  return s.String()
}

func (s *QueryIPV6BandwidthOfeachISPandProvinceRequest) SetDateFrom(v string) *QueryIPV6BandwidthOfeachISPandProvinceRequest {
  s.DateFrom = &v
  return s
}

func (s *QueryIPV6BandwidthOfeachISPandProvinceRequest) SetDateTo(v string) *QueryIPV6BandwidthOfeachISPandProvinceRequest {
  s.DateTo = &v
  return s
}

func (s *QueryIPV6BandwidthOfeachISPandProvinceRequest) SetDomain(v []*string) *QueryIPV6BandwidthOfeachISPandProvinceRequest {
  s.Domain = v
  return s
}

func (s *QueryIPV6BandwidthOfeachISPandProvinceRequest) SetDataInterval(v string) *QueryIPV6BandwidthOfeachISPandProvinceRequest {
  s.DataInterval = &v
  return s
}

func (s *QueryIPV6BandwidthOfeachISPandProvinceRequest) SetProvince(v []*string) *QueryIPV6BandwidthOfeachISPandProvinceRequest {
  s.Province = v
  return s
}

func (s *QueryIPV6BandwidthOfeachISPandProvinceRequest) SetIsp(v []*string) *QueryIPV6BandwidthOfeachISPandProvinceRequest {
  s.Isp = v
  return s
}

func (s *QueryIPV6BandwidthOfeachISPandProvinceRequest) SetIPType(v string) *QueryIPV6BandwidthOfeachISPandProvinceRequest {
  s.IPType = &v
  return s
}

func (s *QueryIPV6BandwidthOfeachISPandProvinceRequest) SetGroupBy(v []*string) *QueryIPV6BandwidthOfeachISPandProvinceRequest {
  s.GroupBy = v
  return s
}

type QueryIPV6BandwidthOfeachISPandProvinceResponse struct {
  // {"en":"result", "zh_CN":"结果"}
  Result []*QueryIPV6BandwidthOfeachISPandProvinceResponseResult `json:"result,omitempty" xml:"result,omitempty" require:"true" type:"Repeated"`
}

func (s QueryIPV6BandwidthOfeachISPandProvinceResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryIPV6BandwidthOfeachISPandProvinceResponse) GoString() string {
  return s.String()
}

func (s *QueryIPV6BandwidthOfeachISPandProvinceResponse) SetResult(v []*QueryIPV6BandwidthOfeachISPandProvinceResponseResult) *QueryIPV6BandwidthOfeachISPandProvinceResponse {
  s.Result = v
  return s
}

type QueryIPV6BandwidthOfeachISPandProvinceResponseResult struct     {
  // {"en":"Domain", "zh_CN":"域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  IspData []*QueryIPV6BandwidthOfeachISPandProvinceResponseResultIspData `json:"ispData,omitempty" xml:"ispData,omitempty" require:"true" type:"Repeated"`
}

func (s QueryIPV6BandwidthOfeachISPandProvinceResponseResult) String() string {
  return tea.Prettify(s)
}

func (s QueryIPV6BandwidthOfeachISPandProvinceResponseResult) GoString() string {
  return s.String()
}

func (s *QueryIPV6BandwidthOfeachISPandProvinceResponseResult) SetDomain(v string) *QueryIPV6BandwidthOfeachISPandProvinceResponseResult {
  s.Domain = &v
  return s
}

func (s *QueryIPV6BandwidthOfeachISPandProvinceResponseResult) SetIspData(v []*QueryIPV6BandwidthOfeachISPandProvinceResponseResultIspData) *QueryIPV6BandwidthOfeachISPandProvinceResponseResult {
  s.IspData = v
  return s
}

type QueryIPV6BandwidthOfeachISPandProvinceResponseResultIspData struct     {
  // {"en":"ISP", "zh_CN":"运营商"}
  Isp *string `json:"isp,omitempty" xml:"isp,omitempty" require:"true"`
  ProvinceData []*QueryIPV6BandwidthOfeachISPandProvinceResponseResultIspDataProvinceData `json:"provinceData,omitempty" xml:"provinceData,omitempty" require:"true" type:"Repeated"`
}

func (s QueryIPV6BandwidthOfeachISPandProvinceResponseResultIspData) String() string {
  return tea.Prettify(s)
}

func (s QueryIPV6BandwidthOfeachISPandProvinceResponseResultIspData) GoString() string {
  return s.String()
}

func (s *QueryIPV6BandwidthOfeachISPandProvinceResponseResultIspData) SetIsp(v string) *QueryIPV6BandwidthOfeachISPandProvinceResponseResultIspData {
  s.Isp = &v
  return s
}

func (s *QueryIPV6BandwidthOfeachISPandProvinceResponseResultIspData) SetProvinceData(v []*QueryIPV6BandwidthOfeachISPandProvinceResponseResultIspDataProvinceData) *QueryIPV6BandwidthOfeachISPandProvinceResponseResultIspData {
  s.ProvinceData = v
  return s
}

type QueryIPV6BandwidthOfeachISPandProvinceResponseResultIspDataProvinceData struct     {
  // {"en":"Province", "zh_CN":"省份"}
  Province *string `json:"province,omitempty" xml:"province,omitempty" require:"true"`
  BandwidthData []*QueryIPV6BandwidthOfeachISPandProvinceResponseResultIspDataProvinceDataBandwidthData `json:"bandwidthData,omitempty" xml:"bandwidthData,omitempty" require:"true" type:"Repeated"`
}

func (s QueryIPV6BandwidthOfeachISPandProvinceResponseResultIspDataProvinceData) String() string {
  return tea.Prettify(s)
}

func (s QueryIPV6BandwidthOfeachISPandProvinceResponseResultIspDataProvinceData) GoString() string {
  return s.String()
}

func (s *QueryIPV6BandwidthOfeachISPandProvinceResponseResultIspDataProvinceData) SetProvince(v string) *QueryIPV6BandwidthOfeachISPandProvinceResponseResultIspDataProvinceData {
  s.Province = &v
  return s
}

func (s *QueryIPV6BandwidthOfeachISPandProvinceResponseResultIspDataProvinceData) SetBandwidthData(v []*QueryIPV6BandwidthOfeachISPandProvinceResponseResultIspDataProvinceDataBandwidthData) *QueryIPV6BandwidthOfeachISPandProvinceResponseResultIspDataProvinceData {
  s.BandwidthData = v
  return s
}

type QueryIPV6BandwidthOfeachISPandProvinceResponseResultIspDataProvinceDataBandwidthData struct     {
  // {"en":"Time
  // 1. When the data size of the query is 5m, the format is yyyy-MM-dd HH:mm; each time slice data value represents the data value in the previous time granularity range. The time slice starting at the beginning of the day is yyyy-MM-dd 00:05, and the last time slice is (yyyy-MM-dd+1) 00:00.
  // 2. When the data granularity of the query is 1h, the format is yyyy-MM-dd HH; each time slice data value represents the data value within the previous time granularity range. The time slice starting at the beginning of the day is yyyy-MM-dd 01, and the last time slice is (yyyy-MM-dd+1) 00.
  // 3. Returns the time slice contained in the start time and end time.", "zh_CN":"时间,
  // 查询的数据粒度为5m时,格式为yyyy-MM-dd HH:mm;每一个时间片数据值代表的是前一个时间粒度范围内的数据值。一天开始的时间片是yyyy-MM-dd 00:05,最后一个时间片是(yyyy-MM-dd+1) 00:00;
  // 查询的数据粒度为1h时,格式为yyyy-MM-dd HH;每一个时间片数据值代表的是前一个时间粒度范围内的数据值。一天开始的时间片是yyyy-MM-dd 01,最后一个时间片是(yyyy-MM-dd+1) 00;
  // 返回开始时间和结束时间包含的时间片。"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {"en":"Bandwidth,unit is Mbps,Keep 2 decimal places", "zh_CN":"带宽值,单位Mbps,保留2位小数。"}
  Value *string `json:"value,omitempty" xml:"value,omitempty" require:"true"`
}

func (s QueryIPV6BandwidthOfeachISPandProvinceResponseResultIspDataProvinceDataBandwidthData) String() string {
  return tea.Prettify(s)
}

func (s QueryIPV6BandwidthOfeachISPandProvinceResponseResultIspDataProvinceDataBandwidthData) GoString() string {
  return s.String()
}

func (s *QueryIPV6BandwidthOfeachISPandProvinceResponseResultIspDataProvinceDataBandwidthData) SetTimestamp(v string) *QueryIPV6BandwidthOfeachISPandProvinceResponseResultIspDataProvinceDataBandwidthData {
  s.Timestamp = &v
  return s
}

func (s *QueryIPV6BandwidthOfeachISPandProvinceResponseResultIspDataProvinceDataBandwidthData) SetValue(v string) *QueryIPV6BandwidthOfeachISPandProvinceResponseResultIspDataProvinceDataBandwidthData {
  s.Value = &v
  return s
}

type QueryIPV6BandwidthOfeachISPandProvincePaths struct {
}

func (s QueryIPV6BandwidthOfeachISPandProvincePaths) String() string {
  return tea.Prettify(s)
}

func (s QueryIPV6BandwidthOfeachISPandProvincePaths) GoString() string {
  return s.String()
}

type QueryIPV6BandwidthOfeachISPandProvinceParameters struct {
}

func (s QueryIPV6BandwidthOfeachISPandProvinceParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryIPV6BandwidthOfeachISPandProvinceParameters) GoString() string {
  return s.String()
}

type QueryIPV6BandwidthOfeachISPandProvinceRequestHeader struct {
}

func (s QueryIPV6BandwidthOfeachISPandProvinceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryIPV6BandwidthOfeachISPandProvinceRequestHeader) GoString() string {
  return s.String()
}

type QueryIPV6BandwidthOfeachISPandProvinceResponseHeader struct {
}

func (s QueryIPV6BandwidthOfeachISPandProvinceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryIPV6BandwidthOfeachISPandProvinceResponseHeader) GoString() string {
  return s.String()
}




type BandwidthMiddleRequest struct {
  // {"en":"cust_en_name of sub-client.
  // When a merged-account wants to  view the information of the subclient,the cust_en_name is required.", "zh_CN":"合并账号下的某个客户的英文名，当合并账号要查看子客户的信息时，必须填写子客户的英文名"}
  Cust *string `json:"cust,omitempty" xml:"cust,omitempty"`
  // {"en":"Specifies the query date:
  // 1.With format yyyy-mm-dd.
  // 2.If not specified,it means today as default.", "zh_CN":"查询的日期，日期格式为yyyy-mm-dd,不选或者为空时默认为当天；"}
  Date *string `json:"date,omitempty" xml:"date,omitempty"`
  // {"en":"1.Must work with 'enddate' and they  specify the query date scope. 
  // 2.With format yyyy-mm-dd.
  // 3.If there is a 'date' parameter,this parameter will be omitted.", "zh_CN":"查询的起始日期，日期格式为yyyy-mm-dd；此参数需与enddate参数配合,若存在date参数,则该参数无效"}
  Startdate *string `json:"startdate,omitempty" xml:"startdate,omitempty"`
  // {"en":"1.Must work with 'startdate' and they  specify the query date scope. 
  // 2.With format yyyy-mm-dd.
  // 3.If there is a 'date' parameter,this parameter will be omitted.", "zh_CN":"查询的结束日期,日期格式为yyyy-mm-dd；此参数需与startdate参数配合,若存在date参数,则该参数无效。"}
  Enddate *string `json:"enddate,omitempty" xml:"enddate,omitempty"`
  // {"en":"domains that been queried:
  // 1.If there are multiple inputs,use  ';' as separator.
  // 2.If not specified, it means all the domains of the account .", "zh_CN":"查询的频道，多个频道值请用英文分号“;”，不选或者为空时默认为所查询客户的所有频道"}
  Channel *string `json:"channel,omitempty" xml:"channel,omitempty"`
  // {"en":"1.If there are multiple inputs,use ';' as separator.For example,u can use 'region=cn;apac' to query data of cn and apac region.
  // 2.If not specified, it means all the regions.", "zh_CN":"查询的加速区域的缩写，多个区域请用英文分号“;”分隔开，如查询大陆及亚太区域，参数填写为：“region=cn;apac”。不选或者为空时默认为全部区域。"}
  Region *string `json:"region,omitempty" xml:"region,omitempty"`
  // {"en":"acceleration type:
  // 1.If there are multiple inputs,use ';' as separator.
  // 2.If not specified or specified as 'all', it means all the accetypes.", "zh_CN":"加速类型参数，如accetype=web。多个请用英文分号“;”分隔开，不填或值为all表示所有类型"}
  Accetype *string `json:"accetype,omitempty" xml:"accetype,omitempty"`
  // {"en":"The response format:
  // 1.optional values:xml, json.
  // 2.'xml' as default.", "zh_CN":"返回结果格式,支持格式为xml和json,默认为xml"}
  Dataformat *string `json:"dataformat,omitempty" xml:"dataformat,omitempty"`
  // {"en":"Specifies if  the 'channel' parameter should be exactly matched:
  // 1.'true' as default.
  // 2. If not 'true',it will query data of channels that ends with any item of input 'channel's.", "zh_CN":" 频道是否完全匹配,为true时，必须填写完整的域名(此时会过滤用户输入的无效或重复频道,所有输入频道都无效时返403)。不为true时，显示以用户输入的频道为结尾的所有频道。默认为true"}
  IsExactMatch *string `json:"isExactMatch,omitempty" xml:"isExactMatch,omitempty"`
  // {"en":"Display statistic result in merged or separate way:
  // 1.If specified 1, get the merged result.
  // 2.If  specified 2,get the separate result.
  // 3.If specifed 3,get both merged result and separate result.
  // 4.If not specified,means '1'.", "zh_CN":" 结果的显示是否提供合并值。填写1时：只提供合并结果；填写2时：只提供拆分值；填写3时：既提供合并值，又提供拆分值。不选或者为空时默认为“1”。"}
  ResultType *string `json:"resultType,omitempty" xml:"resultType,omitempty"`
  // {"en":"Different data types:
  // 1.optional values: hit,miss.
  // 2.If  there are multiple inputs,use ';' as separator.
  // 3.If not specified,it means the merged value of all types.", "zh_CN":"可选值：hit（hit带宽），miss（miss带宽）。多个以英文分号分隔，不选或为空默认展示合计值。"}
  Type *string `json:"type,omitempty" xml:"type,omitempty"`
  // {"en":"Different father types:
  // 1.optional values: stafu,dynfu,all
  // 2.If there are multiple inputs,use ';' as separator.
  // 3.If not specified,it means 'all'.", "zh_CN":"可选值：stafu（静态父），dynfu（动态父），all(全选)；默认all（全选）。"}
  FatherType *string `json:"fatherType,omitempty" xml:"fatherType,omitempty"`
  // {"en":"Return traffic details, 1: mandatory; 0: optional. The default value is 0.", "zh_CN":"返回流量明细，1：需要；0：不需要。默认为0."}
  NeedFlow *string `json:"needFlow,omitempty" xml:"needFlow,omitempty"`
}

func (s BandwidthMiddleRequest) String() string {
  return tea.Prettify(s)
}

func (s BandwidthMiddleRequest) GoString() string {
  return s.String()
}

func (s *BandwidthMiddleRequest) SetCust(v string) *BandwidthMiddleRequest {
  s.Cust = &v
  return s
}

func (s *BandwidthMiddleRequest) SetDate(v string) *BandwidthMiddleRequest {
  s.Date = &v
  return s
}

func (s *BandwidthMiddleRequest) SetStartdate(v string) *BandwidthMiddleRequest {
  s.Startdate = &v
  return s
}

func (s *BandwidthMiddleRequest) SetEnddate(v string) *BandwidthMiddleRequest {
  s.Enddate = &v
  return s
}

func (s *BandwidthMiddleRequest) SetChannel(v string) *BandwidthMiddleRequest {
  s.Channel = &v
  return s
}

func (s *BandwidthMiddleRequest) SetRegion(v string) *BandwidthMiddleRequest {
  s.Region = &v
  return s
}

func (s *BandwidthMiddleRequest) SetAccetype(v string) *BandwidthMiddleRequest {
  s.Accetype = &v
  return s
}

func (s *BandwidthMiddleRequest) SetDataformat(v string) *BandwidthMiddleRequest {
  s.Dataformat = &v
  return s
}

func (s *BandwidthMiddleRequest) SetIsExactMatch(v string) *BandwidthMiddleRequest {
  s.IsExactMatch = &v
  return s
}

func (s *BandwidthMiddleRequest) SetResultType(v string) *BandwidthMiddleRequest {
  s.ResultType = &v
  return s
}

func (s *BandwidthMiddleRequest) SetType(v string) *BandwidthMiddleRequest {
  s.Type = &v
  return s
}

func (s *BandwidthMiddleRequest) SetFatherType(v string) *BandwidthMiddleRequest {
  s.FatherType = &v
  return s
}

func (s *BandwidthMiddleRequest) SetNeedFlow(v string) *BandwidthMiddleRequest {
  s.NeedFlow = &v
  return s
}

type BandwidthMiddleResponse struct {
  // {'en':'provider', 'zh_CN':'结果'}
  Provider *BandwidthMiddleResponseProvider `json:"provider,omitempty" xml:"provider,omitempty" require:"true" type:"Struct"`
}

func (s BandwidthMiddleResponse) String() string {
  return tea.Prettify(s)
}

func (s BandwidthMiddleResponse) GoString() string {
  return s.String()
}

func (s *BandwidthMiddleResponse) SetProvider(v *BandwidthMiddleResponseProvider) *BandwidthMiddleResponse {
  s.Provider = v
  return s
}

type BandwidthMiddleResponseProvider struct {
  // {'en':'tenant', 'zh_CN':'租户'}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {'en':'type', 'zh_CN':'接口类型'}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
  // {'en':'resultType', 'zh_CN':'统计类型'}
  ResultType *string `json:"resultType,omitempty" xml:"resultType,omitempty" require:"true"`
  // {'en':'data', 'zh_CN':'中间缓存带宽数据'}
  Date *BandwidthMiddleResponseProviderDate `json:"date,omitempty" xml:"date,omitempty" require:"true" type:"Struct"`
}

func (s BandwidthMiddleResponseProvider) String() string {
  return tea.Prettify(s)
}

func (s BandwidthMiddleResponseProvider) GoString() string {
  return s.String()
}

func (s *BandwidthMiddleResponseProvider) SetName(v string) *BandwidthMiddleResponseProvider {
  s.Name = &v
  return s
}

func (s *BandwidthMiddleResponseProvider) SetType(v string) *BandwidthMiddleResponseProvider {
  s.Type = &v
  return s
}

func (s *BandwidthMiddleResponseProvider) SetResultType(v string) *BandwidthMiddleResponseProvider {
  s.ResultType = &v
  return s
}

func (s *BandwidthMiddleResponseProvider) SetDate(v *BandwidthMiddleResponseProviderDate) *BandwidthMiddleResponseProvider {
  s.Date = v
  return s
}

type BandwidthMiddleResponseProviderDate struct {
  // {'en':'startdate', 'zh_CN':'开始时间'}
  Startdate *string `json:"startdate,omitempty" xml:"startdate,omitempty" require:"true"`
  // {'en':'enddate', 'zh_CN':'结束时间'}
  Enddate *string `json:"enddate,omitempty" xml:"enddate,omitempty" require:"true"`
  // {'en':'channel', 'zh_CN':'频道'}
  Channel *BandwidthMiddleResponseProviderDateChannel `json:"channel,omitempty" xml:"channel,omitempty" require:"true" type:"Struct"`
}

func (s BandwidthMiddleResponseProviderDate) String() string {
  return tea.Prettify(s)
}

func (s BandwidthMiddleResponseProviderDate) GoString() string {
  return s.String()
}

func (s *BandwidthMiddleResponseProviderDate) SetStartdate(v string) *BandwidthMiddleResponseProviderDate {
  s.Startdate = &v
  return s
}

func (s *BandwidthMiddleResponseProviderDate) SetEnddate(v string) *BandwidthMiddleResponseProviderDate {
  s.Enddate = &v
  return s
}

func (s *BandwidthMiddleResponseProviderDate) SetChannel(v *BandwidthMiddleResponseProviderDateChannel) *BandwidthMiddleResponseProviderDate {
  s.Channel = v
  return s
}

type BandwidthMiddleResponseProviderDateChannel struct {
  // {'en':'channel', 'zh_CN':'频道'}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {'en':'bandwidth', 'zh_CN':'中间缓存带宽数据'}
  Bandwidth []*BandwidthMiddleResponseProviderDateChannelBandwidth `json:"bandwidth,omitempty" xml:"bandwidth,omitempty" require:"true" type:"Repeated"`
}

func (s BandwidthMiddleResponseProviderDateChannel) String() string {
  return tea.Prettify(s)
}

func (s BandwidthMiddleResponseProviderDateChannel) GoString() string {
  return s.String()
}

func (s *BandwidthMiddleResponseProviderDateChannel) SetName(v string) *BandwidthMiddleResponseProviderDateChannel {
  s.Name = &v
  return s
}

func (s *BandwidthMiddleResponseProviderDateChannel) SetBandwidth(v []*BandwidthMiddleResponseProviderDateChannelBandwidth) *BandwidthMiddleResponseProviderDateChannel {
  s.Bandwidth = v
  return s
}

type BandwidthMiddleResponseProviderDateChannelBandwidth struct     {
  // {'en':'timestamp', 'zh_CN':'时间点'}
  Time *string `json:"time,omitempty" xml:"time,omitempty" require:"true"`
  // {'en':'summary', 'zh_CN':'带宽合计'}
  Summary *string `json:"summary,omitempty" xml:"summary,omitempty" require:"true"`
  // {'en':'stafuHit', 'zh_CN':'静态Hit带宽'}
  StafuHit *string `json:"stafuHit,omitempty" xml:"stafuHit,omitempty" require:"true"`
  // {'en':'dynfuHit', 'zh_CN':'动态Hit带宽'}
  DynfuHit *string `json:"dynfuHit,omitempty" xml:"dynfuHit,omitempty" require:"true"`
  // {'en':'stafuMiss', 'zh_CN':'静态Miss带宽'}
  StafuMiss *string `json:"stafuMiss,omitempty" xml:"stafuMiss,omitempty" require:"true"`
  // {'en':'dynfuMiss', 'zh_CN':'动态Miss带宽'}
  DynfuMiss *string `json:"dynfuMiss,omitempty" xml:"dynfuMiss,omitempty" require:"true"`
  // {'en':'thirdWs', 'zh_CN':'带宽'}
  ThirdWs *string `json:"thirdWs,omitempty" xml:"thirdWs,omitempty" require:"true"`
}

func (s BandwidthMiddleResponseProviderDateChannelBandwidth) String() string {
  return tea.Prettify(s)
}

func (s BandwidthMiddleResponseProviderDateChannelBandwidth) GoString() string {
  return s.String()
}

func (s *BandwidthMiddleResponseProviderDateChannelBandwidth) SetTime(v string) *BandwidthMiddleResponseProviderDateChannelBandwidth {
  s.Time = &v
  return s
}

func (s *BandwidthMiddleResponseProviderDateChannelBandwidth) SetSummary(v string) *BandwidthMiddleResponseProviderDateChannelBandwidth {
  s.Summary = &v
  return s
}

func (s *BandwidthMiddleResponseProviderDateChannelBandwidth) SetStafuHit(v string) *BandwidthMiddleResponseProviderDateChannelBandwidth {
  s.StafuHit = &v
  return s
}

func (s *BandwidthMiddleResponseProviderDateChannelBandwidth) SetDynfuHit(v string) *BandwidthMiddleResponseProviderDateChannelBandwidth {
  s.DynfuHit = &v
  return s
}

func (s *BandwidthMiddleResponseProviderDateChannelBandwidth) SetStafuMiss(v string) *BandwidthMiddleResponseProviderDateChannelBandwidth {
  s.StafuMiss = &v
  return s
}

func (s *BandwidthMiddleResponseProviderDateChannelBandwidth) SetDynfuMiss(v string) *BandwidthMiddleResponseProviderDateChannelBandwidth {
  s.DynfuMiss = &v
  return s
}

func (s *BandwidthMiddleResponseProviderDateChannelBandwidth) SetThirdWs(v string) *BandwidthMiddleResponseProviderDateChannelBandwidth {
  s.ThirdWs = &v
  return s
}

type BandwidthMiddlePaths struct {
}

func (s BandwidthMiddlePaths) String() string {
  return tea.Prettify(s)
}

func (s BandwidthMiddlePaths) GoString() string {
  return s.String()
}

type BandwidthMiddleParameters struct {
}

func (s BandwidthMiddleParameters) String() string {
  return tea.Prettify(s)
}

func (s BandwidthMiddleParameters) GoString() string {
  return s.String()
}

type BandwidthMiddleRequestHeader struct {
}

func (s BandwidthMiddleRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s BandwidthMiddleRequestHeader) GoString() string {
  return s.String()
}

type BandwidthMiddleResponseHeader struct {
}

func (s BandwidthMiddleResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s BandwidthMiddleResponseHeader) GoString() string {
  return s.String()
}




type QueryDomainBandwidthRequest struct {
  // {"en":"Domain list.
  // Domain number limits can be adjusted depending on different accounts. The default value is 20(if you want to adjust,please, contact technical support)", "zh_CN":"域名列表
  // 1.域名个数限制根据账号可调,默认为20个（可联系技术支持下单调整）;"}
  QueryDomainBandwidthDomainList *QueryDomainBandwidthDomainList `json:"domain-list,omitempty" xml:"domain-list,omitempty" require:"true"`
}

func (s QueryDomainBandwidthRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryDomainBandwidthRequest) GoString() string {
  return s.String()
}

func (s *QueryDomainBandwidthRequest) SetDomainList(v *QueryDomainBandwidthDomainList) *QueryDomainBandwidthRequest {
  s.QueryDomainBandwidthDomainList = v
  return s
}

type QueryDomainBandwidthDomainList struct {
  // {"en":"Domain", "zh_CN":"域名"}
  DomainName []*string `json:"domain-name,omitempty" xml:"domain-name,omitempty" require:"true" type:"Repeated"`
}

func (s QueryDomainBandwidthDomainList) String() string {
  return tea.Prettify(s)
}

func (s QueryDomainBandwidthDomainList) GoString() string {
  return s.String()
}

func (s *QueryDomainBandwidthDomainList) SetDomainName(v []*string) *QueryDomainBandwidthDomainList {
  s.DomainName = v
  return s
}

type QueryDomainBandwidthResponse struct {
  BandwidthReport []*QueryDomainBandwidthResponseBandwidthReport `json:"bandwidthReport,omitempty" xml:"bandwidthReport,omitempty" require:"true" type:"Repeated"`
}

func (s QueryDomainBandwidthResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryDomainBandwidthResponse) GoString() string {
  return s.String()
}

func (s *QueryDomainBandwidthResponse) SetBandwidthReport(v []*QueryDomainBandwidthResponseBandwidthReport) *QueryDomainBandwidthResponse {
  s.BandwidthReport = v
  return s
}

type QueryDomainBandwidthResponseBandwidthReport struct     {
  // {"en":"Date
  // When the querying data granularity is fiveminutes, the format is yyyy-MM-dd HH:mm; the data value of every time slice represents the data value within the previous time granularity range. The first time slice of the day is yyyy-MM-dd 00:05 AM, and the last one is yyyy-MM-dd 24:00;When the querying data granularity is hourly, the format is yyyy-MM-dd HH; the data value of every time slice represents the data value within the previous time granularity range. The first time slice of the day is yyyy-MM-dd 01, and the last one is yyyy-MM-dd 24;When the querying data granularity is daily, the format is yyyy-MM-dd; the data value of every time slice represents the value of the daily data;Return the time slice contained in start time and in end time", "zh_CN":"时间
  // 1.查询的数据粒度为fiveminutes时,格式为yyyy-MM-dd HH:mm;每一个时间片数据值代表的是前一个时间粒度范围内的数据值。一天开始的时间片是yyyy-MM-dd 00:05,最后一个时间片是yyyy-MM-dd 24:00。
  // 2.查询的数据粒度为hourly时,格式为yyyy-MM-dd HH;每一个时间片数据值代表的是前一个时间粒度范围内的数据值。一天开始的时间片是yyyy-MM-dd 01,最后一个时间片是yyyy-MM-dd 24。
  // 3.查询的数据粒度为daily时,格式为yyyy-MM-dd;每一个时间片数据值代表的该天内的数据值;
  // 4.返回开始时间和结束时间包含的时间片。"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {"en":"Bandwidth, keep the number to four decimal places. Unit: Mbps", "zh_CN":"带宽,保留4位小数,单位为Mbps"}
  Bandwidth *int `json:"bandwidth,omitempty" xml:"bandwidth,omitempty" require:"true"`
}

func (s QueryDomainBandwidthResponseBandwidthReport) String() string {
  return tea.Prettify(s)
}

func (s QueryDomainBandwidthResponseBandwidthReport) GoString() string {
  return s.String()
}

func (s *QueryDomainBandwidthResponseBandwidthReport) SetTimestamp(v string) *QueryDomainBandwidthResponseBandwidthReport {
  s.Timestamp = &v
  return s
}

func (s *QueryDomainBandwidthResponseBandwidthReport) SetBandwidth(v int) *QueryDomainBandwidthResponseBandwidthReport {
  s.Bandwidth = &v
  return s
}

type QueryDomainBandwidthPaths struct {
}

func (s QueryDomainBandwidthPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryDomainBandwidthPaths) GoString() string {
  return s.String()
}

type QueryDomainBandwidthParameters struct {
  // {"en":"Start time
  // 1.The format is yyyy-MM-ddTHH:mm:ss+08:00;
  // 2.And smaller than the current time and dateTo;
  // 3.Period between dataFrom and dateTo cannot be longer than 31 days", "zh_CN":"开始时间
  // 1.格式为yyyy-MM-ddTHH:mm:ss+08:00;
  // 2.并且小于当前时间和dateTo;
  // 3.dateFrom和dateTo相差不能超过31天（可联系技术支持调整）;4.只能查询最近2年内数据。"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty" require:"true"`
  // {"en":"End time
  // 1.The format is yyyy-MM-ddTHH:mm:ss+08:00;
  // 2.Must be greater than dateFrom;
  // 3.If it's greater than the current time, then the current time will be assigned as the value", "zh_CN":"结束时间
  // 1.格式为yyyy-MM-ddTHH:mm:ss+08:00;
  // 2.必须大于dateFrom;
  // 3.如果大于当前时间,则重新赋值为当前时间;"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty" require:"true"`
  // {"en":"Data granularity
  // fiveminutes: five minutes, hourly: one hour, daily: one day;If not specified, daily is set as the default value", "zh_CN":"数据粒度
  // 1.fiveminutes:5分钟,hourly:1小时,daily:1天;
  // 2.不传递,默认为daily;
  // 3.传递fiveminutes时,若客户数据采集粒度有特殊配置将按实际配置粒度返回。"}
  Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s QueryDomainBandwidthParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryDomainBandwidthParameters) GoString() string {
  return s.String()
}

func (s *QueryDomainBandwidthParameters) SetDateFrom(v string) *QueryDomainBandwidthParameters {
  s.DateFrom = &v
  return s
}

func (s *QueryDomainBandwidthParameters) SetDateTo(v string) *QueryDomainBandwidthParameters {
  s.DateTo = &v
  return s
}

func (s *QueryDomainBandwidthParameters) SetType(v string) *QueryDomainBandwidthParameters {
  s.Type = &v
  return s
}

type QueryDomainBandwidthRequestHeader struct {
}

func (s QueryDomainBandwidthRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryDomainBandwidthRequestHeader) GoString() string {
  return s.String()
}

type QueryDomainBandwidthResponseHeader struct {
}

func (s QueryDomainBandwidthResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryDomainBandwidthResponseHeader) GoString() string {
  return s.String()
}




