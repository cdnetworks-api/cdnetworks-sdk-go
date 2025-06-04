package reportflow

// This file is auto-generated, don't edit it. Thanks.
import (
  "github.com/alibabacloud-go/tea/tea"
)

type QueryTrafficForMultiDomainRequest struct {
  // {"en":"Domain list:
  //   Domain number limits can be adjusted depending on different accounts. The default value is 20(if you want to adjust,please, contact technical support)", "zh_CN":"域名列表:
  //   1.域名个数限制根据账号可调,默认为20个(可联系技术支持下单调整);"}
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
  // 1.The format is yyyy-MM-ddTHH:mm:ss+08:00;
  // 2.And smaller than the current time and 'dateTo';
  // 3.Period between 'dataFrom' and 'dateTo' cannot be longer than 31 days", "zh_CN":"开始时间
  // 1.格式为yyyy-MM-ddTHH:mm:ss+08:00;
  // 2.并且小于当前时间和dateTo;
  // 3.dateFrom和dateTo相差不能超过31天(可联系技术支持调整);4.只能查询最近2年内数据。"}
  Datefrom *string `json:"datefrom,omitempty" xml:"datefrom,omitempty" require:"true"`
  // {"en":"End time
  // 1.The format is yyyy-MM-ddTHH:mm:ss+08:00;
  // 2.Must be greater than 'dateFrom';
  // 3.If it's greater than the current time, then the current time is assigned as the value", "zh_CN":"结束时间
  // 1.格式为yyyy-MM-ddTHH:mm:ss+08:00;
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




type ReportFlowDomainCountryServiceRequest struct {
  // {"en":"Starting time
  // 
  // 1. The time format is yyyy-MM-ddTHH:mm:ss+08:00, for example, 2016-12-02T10:00:00+08:00 (for Beijing time, December 2, 2016, 10:00:00);
  // 2. Cannot be greater than the current time
  // 3. Get up to the last six months (183 days) of data.", "zh_CN":"开始时间
  // 1. 时间格式为yyyy-MM-ddTHH:mm:ss+08:00,例如,2016-12-02T10:00:00+08:00(为北京时间2016年12月2日10点0分0秒);
  // 2. 不能大于当前时间
  // 3. 最多可获取最近半年(183天)的数据。"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty"`
  // {"en":"End Time:
  // 1. Time format 2016-12-02T10:00:00+08:00
  // 2. The end time needs to be greater than the start time. If the end time is greater than the current time, the current time is taken.
  // 3. dateFrom, dateTo are not passed, the default query for the past 24 hours; if there is only one untransmitted, throw an exception
  // 4. Allow query maximum time interval: 7 days, that is, the difference between dateFrom and dateTo can&rsquo;t exceed 7 days (can contact technical support adjustment, up to 31 days).", "zh_CN":"结束时间:
  // 1. 时间格式2016-12-02T10:00:00+08:00
  // 2. 结束时间需大于开始时间,结束时间如果大于当前时间,取当前时间。
  // 3. dateFrom,dateTo二者都未传,默认查询过去的24小时;如仅有一个未传,抛异常
  // 4. 允许查询最大时间间隔:7天,即dateFrom和dateTo相差不能超过7天(可联系技术支持调整，最长31天)。"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty"`
  // {"en":"Data granularity:
  // 1. Support 5m (5 minutes granularity),1d (1 day granularity)
  // 2. Do not pass the default to 5m", "zh_CN":"数据粒度:
  // 1. 支持5m(5分钟粒度),1d(天粒度)
  // 2. 不传默认为5m"}
  DataInterval *string `json:"dataInterval,omitempty" xml:"dataInterval,omitempty"`
  // {"en":"domain name:
  // 
  // 1. The maximum number of deliverable domain names is 20 by default (can be contacted by technical support);
  // 2. Automatically filter out illegal domain names (such as passing illegal domain names, they will be filtered out, and the query results only return data of legitimate domain names).", "zh_CN":"域名:
  // 1. 可传递域名数量上限默认为20个(可联系技术支持调整);
  // 2. 自动过滤掉非法域名(如传递非法域名,会被过滤掉,查询结果只返回合法域名的数据)。"}
  Domain []*string `json:"domain,omitempty" xml:"domain,omitempty" require:"true" type:"Repeated"`
  // {"en":"Country code:
  // 
  // 1. Do not pass the default query for all countries and regions;
  // 2. The values that can be passed are detailed in the Countrycode list of appendix table on the API Overview page", "zh_CN":"国家地区代号:
  // 1. 不传默认查询全部国家地区;
  // 2. 可传递的值详见概览页国家地区列表说明。"}
  CountryCode []*string `json:"countryCode,omitempty" xml:"countryCode,omitempty" type:"Repeated"`
  // {"en":"Grouping dimension:
  // 
  // 1. Optional values domain, country and aggregatedOversea, can pass in single or multiple values. The aggregatedOversea and country cannot be passed at the same time;
  // 
  // 2. If there is an incoming, the detailed data will be displayed according to the dimension:
  // domain: Group display by domain name dimension;
  // country: Group display by country dimension;
  // aggregatedOversea: Group display according to domestic and overseas dimensions.
  // 3. The result hierarchy is fixed in order, and the order of the parameters does not affect the order of the returned results. For example: 'groupBy': ['domain','country'] and 'groupBy': ['country','domain'] return the same result.",
  // "zh_CN":"1. 可选值domain、country、aggregatedOversea,可传入单个或多个值,其中不能同时传aggregatedOversea 和 country;
  // 2. 有传入则按照该维度展示明细数据:
  //   1.domain:按照域名维度进行分组展示;
  //   2.domain:country:按照国家维度进行分组展示;
  //   3.aggregatedOversea:按照国内 和 海外维度进行分组展示
  // 3. 返回结果层级顺序固定,入参顺序不影响返回结果顺序。例如:groupBy: [domain,country]与groupBy: [country,domain]返回结果一样。"}
  GroupBy []*string `json:"groupBy,omitempty" xml:"groupBy,omitempty" type:"Repeated"`
}

func (s ReportFlowDomainCountryServiceRequest) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowDomainCountryServiceRequest) GoString() string {
  return s.String()
}

func (s *ReportFlowDomainCountryServiceRequest) SetDateFrom(v string) *ReportFlowDomainCountryServiceRequest {
  s.DateFrom = &v
  return s
}

func (s *ReportFlowDomainCountryServiceRequest) SetDateTo(v string) *ReportFlowDomainCountryServiceRequest {
  s.DateTo = &v
  return s
}

func (s *ReportFlowDomainCountryServiceRequest) SetDataInterval(v string) *ReportFlowDomainCountryServiceRequest {
  s.DataInterval = &v
  return s
}

func (s *ReportFlowDomainCountryServiceRequest) SetDomain(v []*string) *ReportFlowDomainCountryServiceRequest {
  s.Domain = v
  return s
}

func (s *ReportFlowDomainCountryServiceRequest) SetCountryCode(v []*string) *ReportFlowDomainCountryServiceRequest {
  s.CountryCode = v
  return s
}

func (s *ReportFlowDomainCountryServiceRequest) SetGroupBy(v []*string) *ReportFlowDomainCountryServiceRequest {
  s.GroupBy = v
  return s
}

type ReportFlowDomainCountryServiceResponse struct {
  // {"en":"result", "zh_CN":"结果"}
  Result []*ReportFlowDomainCountryServiceResponseResult `json:"result,omitempty" xml:"result,omitempty" require:"true" type:"Repeated"`
}

func (s ReportFlowDomainCountryServiceResponse) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowDomainCountryServiceResponse) GoString() string {
  return s.String()
}

func (s *ReportFlowDomainCountryServiceResponse) SetResult(v []*ReportFlowDomainCountryServiceResponseResult) *ReportFlowDomainCountryServiceResponse {
  s.Result = v
  return s
}

type ReportFlowDomainCountryServiceResponseResult struct     {
  // {"en":"Domain", "zh_CN":"域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  CountryData []*ReportFlowDomainCountryServiceResponseResultCountryData `json:"countryData,omitempty" xml:"countryData,omitempty" require:"true" type:"Repeated"`
}

func (s ReportFlowDomainCountryServiceResponseResult) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowDomainCountryServiceResponseResult) GoString() string {
  return s.String()
}

func (s *ReportFlowDomainCountryServiceResponseResult) SetDomain(v string) *ReportFlowDomainCountryServiceResponseResult {
  s.Domain = &v
  return s
}

func (s *ReportFlowDomainCountryServiceResponseResult) SetCountryData(v []*ReportFlowDomainCountryServiceResponseResultCountryData) *ReportFlowDomainCountryServiceResponseResult {
  s.CountryData = v
  return s
}

type ReportFlowDomainCountryServiceResponseResultCountryData struct     {
  // {"en":"Country code", "zh_CN":"国家地区代号"}
  CountryCode *string `json:"countryCode,omitempty" xml:"countryCode,omitempty" require:"true"`
  // {"en":"Country name", "zh_CN":"国家地区名称"}
  CountryName *string `json:"countryName,omitempty" xml:"countryName,omitempty" require:"true"`
  // {"en":"Summary of traffic in national regions: Summary of traffic flow in a single country region during the query period, unit of measure MB, retaining 2 decimal places", "zh_CN":"国家地区流量汇总:单个国家地区流量在查询时段内的流量汇总值,计量单位MB,保留2位小数"}
  FlowSum *string `json:"flowSum,omitempty" xml:"flowSum,omitempty" require:"true"`
  // {"en":"National regional traffic ratio: the proportion (percentage) of traffic value in a single country region during the query period, retaining 2 decimal places", "zh_CN":"国家地区流量占比:单个国家地区流量在查询时段内的流量值的占比(百分比),保留2位小数"}
  FlowPercentage *string `json:"flowPercentage,omitempty" xml:"flowPercentage,omitempty" require:"true"`
  FlowData []*ReportFlowDomainCountryServiceResponseResultCountryDataFlowData `json:"flowData,omitempty" xml:"flowData,omitempty" require:"true" type:"Repeated"`
}

func (s ReportFlowDomainCountryServiceResponseResultCountryData) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowDomainCountryServiceResponseResultCountryData) GoString() string {
  return s.String()
}

func (s *ReportFlowDomainCountryServiceResponseResultCountryData) SetCountryCode(v string) *ReportFlowDomainCountryServiceResponseResultCountryData {
  s.CountryCode = &v
  return s
}

func (s *ReportFlowDomainCountryServiceResponseResultCountryData) SetCountryName(v string) *ReportFlowDomainCountryServiceResponseResultCountryData {
  s.CountryName = &v
  return s
}

func (s *ReportFlowDomainCountryServiceResponseResultCountryData) SetFlowSum(v string) *ReportFlowDomainCountryServiceResponseResultCountryData {
  s.FlowSum = &v
  return s
}

func (s *ReportFlowDomainCountryServiceResponseResultCountryData) SetFlowPercentage(v string) *ReportFlowDomainCountryServiceResponseResultCountryData {
  s.FlowPercentage = &v
  return s
}

func (s *ReportFlowDomainCountryServiceResponseResultCountryData) SetFlowData(v []*ReportFlowDomainCountryServiceResponseResultCountryDataFlowData) *ReportFlowDomainCountryServiceResponseResultCountryData {
  s.FlowData = v
  return s
}

type ReportFlowDomainCountryServiceResponseResultCountryDataFlowData struct     {
  // {"en":"1. When the data size of the query is 5m, the format is yyyy-MM-dd HH:mm; each time slice data value represents the data value in the previous time granularity range. The time slice starting at the beginning of the day is yyyy-MM-dd 00:05, and the last time slice is (yyyy-MM-dd+1) 00:00.
  // 2. Returns the time slice contained in the start time and end time.", "zh_CN":"1. 查询的数据粒度为5m时,格式为yyyy-MM-dd  HH:mm;每一个时间片数据值代表的是前一个时间粒度范围内的数据值。一天开始的时间片是yyyy-MM-dd  00:05,最后一个时间片是(yyyy-MM-dd+1) 00:00。2. 返回开始时间和结束时间包含的时间片。"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {"en":"Flow value, in megabytes, 2 decimal places", "zh_CN":"流量值,计量单位MB,保留2位小数"}
  Value *string `json:"value,omitempty" xml:"value,omitempty" require:"true"`
  // {"en":"Bandwidth value. Unit is Mbps and 2 digits of decimals are allowed.", "zh_CN":"带宽值,单位Mbps,保留2位小数"}
  Bandwidth *string `json:"bandwidth,omitempty" xml:"bandwidth,omitempty" require:"true"`
}

func (s ReportFlowDomainCountryServiceResponseResultCountryDataFlowData) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowDomainCountryServiceResponseResultCountryDataFlowData) GoString() string {
  return s.String()
}

func (s *ReportFlowDomainCountryServiceResponseResultCountryDataFlowData) SetTimestamp(v string) *ReportFlowDomainCountryServiceResponseResultCountryDataFlowData {
  s.Timestamp = &v
  return s
}

func (s *ReportFlowDomainCountryServiceResponseResultCountryDataFlowData) SetValue(v string) *ReportFlowDomainCountryServiceResponseResultCountryDataFlowData {
  s.Value = &v
  return s
}

func (s *ReportFlowDomainCountryServiceResponseResultCountryDataFlowData) SetBandwidth(v string) *ReportFlowDomainCountryServiceResponseResultCountryDataFlowData {
  s.Bandwidth = &v
  return s
}

type ReportFlowDomainCountryServicePaths struct {
}

func (s ReportFlowDomainCountryServicePaths) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowDomainCountryServicePaths) GoString() string {
  return s.String()
}

type ReportFlowDomainCountryServiceParameters struct {
}

func (s ReportFlowDomainCountryServiceParameters) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowDomainCountryServiceParameters) GoString() string {
  return s.String()
}

type ReportFlowDomainCountryServiceRequestHeader struct {
}

func (s ReportFlowDomainCountryServiceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowDomainCountryServiceRequestHeader) GoString() string {
  return s.String()
}

type ReportFlowDomainCountryServiceResponseHeader struct {
}

func (s ReportFlowDomainCountryServiceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportFlowDomainCountryServiceResponseHeader) GoString() string {
  return s.String()
}




type QuerySumUpTrafficUnderAccountRequest struct {
}

func (s QuerySumUpTrafficUnderAccountRequest) String() string {
  return tea.Prettify(s)
}

func (s QuerySumUpTrafficUnderAccountRequest) GoString() string {
  return s.String()
}

type QuerySumUpTrafficUnderAccountResponse struct {
  // {"en":"Total traffic. Keep two digits of decimals. Unit: MB", "zh_CN":"总流量,保留2位小数,单位为MB"}
  FlowSummary *string `json:"flow-summary,omitempty" xml:"flow-summary,omitempty" require:"true"`
  // {"en":"flowData", "zh_CN":"流量数据"}
  FlowData []*QuerySumUpTrafficUnderAccountResponseFlowData `json:"flow-data,omitempty" xml:"flow-data,omitempty" require:"true" type:"Repeated"`
}

func (s QuerySumUpTrafficUnderAccountResponse) String() string {
  return tea.Prettify(s)
}

func (s QuerySumUpTrafficUnderAccountResponse) GoString() string {
  return s.String()
}

func (s *QuerySumUpTrafficUnderAccountResponse) SetFlowSummary(v string) *QuerySumUpTrafficUnderAccountResponse {
  s.FlowSummary = &v
  return s
}

func (s *QuerySumUpTrafficUnderAccountResponse) SetFlowData(v []*QuerySumUpTrafficUnderAccountResponseFlowData) *QuerySumUpTrafficUnderAccountResponse {
  s.FlowData = v
  return s
}

type QuerySumUpTrafficUnderAccountResponseFlowData struct     {
  // {"en":"Date
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

func (s QuerySumUpTrafficUnderAccountResponseFlowData) String() string {
  return tea.Prettify(s)
}

func (s QuerySumUpTrafficUnderAccountResponseFlowData) GoString() string {
  return s.String()
}

func (s *QuerySumUpTrafficUnderAccountResponseFlowData) SetTimestamp(v string) *QuerySumUpTrafficUnderAccountResponseFlowData {
  s.Timestamp = &v
  return s
}

func (s *QuerySumUpTrafficUnderAccountResponseFlowData) SetFlow(v string) *QuerySumUpTrafficUnderAccountResponseFlowData {
  s.Flow = &v
  return s
}

type QuerySumUpTrafficUnderAccountPaths struct {
}

func (s QuerySumUpTrafficUnderAccountPaths) String() string {
  return tea.Prettify(s)
}

func (s QuerySumUpTrafficUnderAccountPaths) GoString() string {
  return s.String()
}

type QuerySumUpTrafficUnderAccountParameters struct {
  // {"en":"Start time
  // 1.The format is yyyy-MM-ddTHH:mm:ss+08:00;
  // 2.Must be smaller than the current time and 'dateto';
  // 3.Period between 'datafrom' and 'dateto' cannot be longer than 31 days", "zh_CN":"开始时间
  // 1.格式为yyyy-MM-ddTHH:mm:ss+08:00;
  // 2.必须小于当前时间和dateto;
  // 3.dateFrom和dateTo相差不能超过31天(可联系技术支持调整);4.只能查询最近2年内数据。"}
  Datefrom *string `json:"datefrom,omitempty" xml:"datefrom,omitempty" require:"true"`
  // {"en":"End time
  // 1.The format is yyyy-MM-ddTHH:mm:ss+08:00;
  // 2.Must be greater than 'datefrom'; 
  // 3.if it's greater than the current time, then the current time is assigned as the value", "zh_CN":"结束时间
  // 1.格式为yyyy-MM-ddTHH:mm:ss+08:00;
  // 2.必须大于datefrom;如果大于当前时间,则重新赋值为当前时间;"}
  Dateto *string `json:"dateto,omitempty" xml:"dateto,omitempty" require:"true"`
  // {"en":"Data granularity
  // 1.fiveminutes: five minutes, hourly: one hour, daily: one day;
  // 2.If not specified, daily is set as the default value;
  // 3.If fiveminutes is specified as the value, then data is returned in actual configured granularity when there is a specific configuration on data collecting granularity for the custome", "zh_CN":"数据粒度
  // 1.fiveminutes:5分钟,hourly:1小时,daily:1天;
  // 2.不传递,默认为daily;
  // 3.传递fiveminutes时,若客户数据采集粒度有特殊配置将按实际配置粒度返回。"}
  Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s QuerySumUpTrafficUnderAccountParameters) String() string {
  return tea.Prettify(s)
}

func (s QuerySumUpTrafficUnderAccountParameters) GoString() string {
  return s.String()
}

func (s *QuerySumUpTrafficUnderAccountParameters) SetDatefrom(v string) *QuerySumUpTrafficUnderAccountParameters {
  s.Datefrom = &v
  return s
}

func (s *QuerySumUpTrafficUnderAccountParameters) SetDateto(v string) *QuerySumUpTrafficUnderAccountParameters {
  s.Dateto = &v
  return s
}

func (s *QuerySumUpTrafficUnderAccountParameters) SetType(v string) *QuerySumUpTrafficUnderAccountParameters {
  s.Type = &v
  return s
}

type QuerySumUpTrafficUnderAccountRequestHeader struct {
}

func (s QuerySumUpTrafficUnderAccountRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QuerySumUpTrafficUnderAccountRequestHeader) GoString() string {
  return s.String()
}

type QuerySumUpTrafficUnderAccountResponseHeader struct {
}

func (s QuerySumUpTrafficUnderAccountResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QuerySumUpTrafficUnderAccountResponseHeader) GoString() string {
  return s.String()
}




type QueryTrafficRequestInTotalAndPeakValueRequest struct {
  // {"en":"Start time:
  // 1. Time format is yyyy-MM-ddTHH:mm:ss+08:00, for example, 2016-12-02T10:00:00+08:00 (10:00 on 2nd of December 2016, Beijing Time);
  // 2. No bigger than the current time;
  // 3. Data in the last 183 days at most can be queried.", "zh_CN":"开始时间：
  // 1.时间格式为yyyy-MM-ddTHH:mm:ss+08:00，例如，2016-12-02T10:00:00+08:00（为北京时间2016年12月2日10点0分0秒）；
  // 2.不能大于当前时间;
  // 3.最多可获取最近半年（183天）的数据。"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty"`
  // {"en":"End time:
  // 1. the time format is 2016-12-02T10:00:00+08:00;
  // 2. End time should be greater than start time. If the end time is greater than current time, current time will be used;
  // 3.  If both fields of dataFrom and dateTo are left empty, then data in the last 24 hours will be queried by default, if only one field is filled in and one is left empty, then exception will be occur;
  // 4. Allowable maximum time range for query: 1 day, means the period between dateFrom to dateTo should not exceed 1 day (can be adjusted by contacting technical support up to 31 days).", "zh_CN":"结束时间：
  // 1.时间格式2016-12-02T10:00:00+08:00;
  // 2.结束时间需大于开始时间，结束时间如果大于当前时间，取当前时间;
  // 3.dateFrom，dateTo二者都未传，默认查询过去的24小时，如仅有一个未传，抛异常;
  // 4.允许查询最大时间间隔：1天，即dateFrom和dateTo相差不能超过1天（可联系技术支持调整，最大不超过31天）。"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty"`
  // {"en":"Domain name:
  // 1. The default upper limit to domains that can be entered is 200 (Contact technical support to adjust, the upper limit is 500);
  // 2. All domains under the account are queried if this input parameter is not specified, but if the number of domains under the account exceeds limits, no query will be done (Error).", "zh_CN":"域名：
  // 1.可传递域名数量上限默认200个（可联系技术支持调整，最高上限500）;
  // 2.未传递该入参时查询账号下所有域名，但当账号下域名数量超过限制时不可查询（报错）。"}
  Domain []*string `json:"domain,omitempty" xml:"domain,omitempty" type:"Repeated"`
  // {"en":"Group keywords:
  // 1. By default, group data will be displayed;
  // 2. If there are keywords entered, value details shall be displayed by keywords;
  // If domain is specified to groupBy, it means results are returned according to domains;
  // 3. Only domain can be specified.", "zh_CN":"分组关键词：
  // 1.默认聚合展示；
  // 2.传入关键词则代表需要按照关键词对应的值展示明细；
  // 例如groupBy传domain，则代表返回按照domain明细展开。
  // 3.只能传递domain。"}
  GroupBy []*string `json:"groupBy,omitempty" xml:"groupBy,omitempty" type:"Repeated"`
}

func (s QueryTrafficRequestInTotalAndPeakValueRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryTrafficRequestInTotalAndPeakValueRequest) GoString() string {
  return s.String()
}

func (s *QueryTrafficRequestInTotalAndPeakValueRequest) SetDateFrom(v string) *QueryTrafficRequestInTotalAndPeakValueRequest {
  s.DateFrom = &v
  return s
}

func (s *QueryTrafficRequestInTotalAndPeakValueRequest) SetDateTo(v string) *QueryTrafficRequestInTotalAndPeakValueRequest {
  s.DateTo = &v
  return s
}

func (s *QueryTrafficRequestInTotalAndPeakValueRequest) SetDomain(v []*string) *QueryTrafficRequestInTotalAndPeakValueRequest {
  s.Domain = v
  return s
}

func (s *QueryTrafficRequestInTotalAndPeakValueRequest) SetGroupBy(v []*string) *QueryTrafficRequestInTotalAndPeakValueRequest {
  s.GroupBy = v
  return s
}

type QueryTrafficRequestInTotalAndPeakValueResponse struct {
  // {"en":"", "zh_CN":""}
  Result []*QueryTrafficRequestInTotalAndPeakValueResponseResult `json:"result,omitempty" xml:"result,omitempty" require:"true" type:"Repeated"`
}

func (s QueryTrafficRequestInTotalAndPeakValueResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryTrafficRequestInTotalAndPeakValueResponse) GoString() string {
  return s.String()
}

func (s *QueryTrafficRequestInTotalAndPeakValueResponse) SetResult(v []*QueryTrafficRequestInTotalAndPeakValueResponseResult) *QueryTrafficRequestInTotalAndPeakValueResponse {
  s.Result = v
  return s
}

type QueryTrafficRequestInTotalAndPeakValueResponseResult struct     {
  // {"en":"Domain", "zh_CN":"域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {"en":"totalFlow, unit: MB, 2 decimal places reserved, example (74099.92)", "zh_CN":"总流量，单位:MB ，保留2位小数，示例 ( 74099.92 )"}
  TotalFlow *string `json:"totalFlow,omitempty" xml:"totalFlow,omitempty" require:"true"`
  // {"en":"totalRequest", "zh_CN":"总请求数"}
  TotalRequest *string `json:"totalRequest,omitempty" xml:"totalRequest,omitempty" require:"true"`
  // {"en":"peakBandwidth, unit: Mbps, 2 decimal places reserved, example (74099.92)", "zh_CN":"峰值带宽，单位: Mbps，保留2位小数，示例 （931556.21）"}
  PeakBandwidth *string `json:"peakBandwidth,omitempty" xml:"peakBandwidth,omitempty" require:"true"`
  // {"en":"peakTime", "zh_CN":"峰值时间，示例（2019-02-13 18:01）"}
  PeakTime *string `json:"peakTime,omitempty" xml:"peakTime,omitempty" require:"true"`
  // {"en":"Peak request", "zh_CN":"请求数峰值"}
  PeakRequest *string `json:"peakRequest,omitempty" xml:"peakRequest,omitempty" require:"true"`
  // {"en":"Peak time of request", "zh_CN":"请求数峰值时间"}
  PeakRequestTime *string `json:"peakRequestTime,omitempty" xml:"peakRequestTime,omitempty" require:"true"`
  // {"en":"", "zh_CN":""}
  FlowRequestData []*QueryTrafficRequestInTotalAndPeakValueResponseResultFlowRequestData `json:"flowRequestData,omitempty" xml:"flowRequestData,omitempty" require:"true" type:"Repeated"`
}

func (s QueryTrafficRequestInTotalAndPeakValueResponseResult) String() string {
  return tea.Prettify(s)
}

func (s QueryTrafficRequestInTotalAndPeakValueResponseResult) GoString() string {
  return s.String()
}

func (s *QueryTrafficRequestInTotalAndPeakValueResponseResult) SetDomain(v string) *QueryTrafficRequestInTotalAndPeakValueResponseResult {
  s.Domain = &v
  return s
}

func (s *QueryTrafficRequestInTotalAndPeakValueResponseResult) SetTotalFlow(v string) *QueryTrafficRequestInTotalAndPeakValueResponseResult {
  s.TotalFlow = &v
  return s
}

func (s *QueryTrafficRequestInTotalAndPeakValueResponseResult) SetTotalRequest(v string) *QueryTrafficRequestInTotalAndPeakValueResponseResult {
  s.TotalRequest = &v
  return s
}

func (s *QueryTrafficRequestInTotalAndPeakValueResponseResult) SetPeakBandwidth(v string) *QueryTrafficRequestInTotalAndPeakValueResponseResult {
  s.PeakBandwidth = &v
  return s
}

func (s *QueryTrafficRequestInTotalAndPeakValueResponseResult) SetPeakTime(v string) *QueryTrafficRequestInTotalAndPeakValueResponseResult {
  s.PeakTime = &v
  return s
}

func (s *QueryTrafficRequestInTotalAndPeakValueResponseResult) SetPeakRequest(v string) *QueryTrafficRequestInTotalAndPeakValueResponseResult {
  s.PeakRequest = &v
  return s
}

func (s *QueryTrafficRequestInTotalAndPeakValueResponseResult) SetPeakRequestTime(v string) *QueryTrafficRequestInTotalAndPeakValueResponseResult {
  s.PeakRequestTime = &v
  return s
}

func (s *QueryTrafficRequestInTotalAndPeakValueResponseResult) SetFlowRequestData(v []*QueryTrafficRequestInTotalAndPeakValueResponseResultFlowRequestData) *QueryTrafficRequestInTotalAndPeakValueResponseResult {
  s.FlowRequestData = v
  return s
}

type QueryTrafficRequestInTotalAndPeakValueResponseResultFlowRequestData struct     {
  // {"en":"timestamp
  // 1. When the data granularity of query is 5m, the format is yyyy-mm-dd HH: MM;Each time slice data value represents the data value in the previous time granularity range.The time slice at the beginning of the day is yyyy-mm-dd 00:05, and the last time slice is (yyyy-mm-dd +1) 00:00;
  // 2. Return the time slice of start time and end time.", "zh_CN":"时间片
  // 1.查询的数据粒度为5m时，格式为yyyy-MM-dd HH:mm；每一个时间片数据值代表的是前一个时间粒度范围内的数据值。一天开始的时间片是yyyy-MM-dd 00:05，最后一个时间片是（yyyy-MM-dd+1） 00:00;
  // 2.返回开始时间和结束时间包含的时间片。"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {"en":"flow", "zh_CN":"流量值，单位MB，保留2位小数；"}
  Flow *string `json:"flow,omitempty" xml:"flow,omitempty" require:"true"`
  // {"en":"Bandwidth, unit: Mbps, 2 decimal places reserved, example (931556.21)", "zh_CN":"带宽值，单位: Mbps，保留2位小数，示例 （931556.21）"}
  Bandwidth *string `json:"bandwidth,omitempty" xml:"bandwidth,omitempty" require:"true"`
  // {"en":"request", "zh_CN":"请求数"}
  Request *string `json:"request,omitempty" xml:"request,omitempty" require:"true"`
}

func (s QueryTrafficRequestInTotalAndPeakValueResponseResultFlowRequestData) String() string {
  return tea.Prettify(s)
}

func (s QueryTrafficRequestInTotalAndPeakValueResponseResultFlowRequestData) GoString() string {
  return s.String()
}

func (s *QueryTrafficRequestInTotalAndPeakValueResponseResultFlowRequestData) SetTimestamp(v string) *QueryTrafficRequestInTotalAndPeakValueResponseResultFlowRequestData {
  s.Timestamp = &v
  return s
}

func (s *QueryTrafficRequestInTotalAndPeakValueResponseResultFlowRequestData) SetFlow(v string) *QueryTrafficRequestInTotalAndPeakValueResponseResultFlowRequestData {
  s.Flow = &v
  return s
}

func (s *QueryTrafficRequestInTotalAndPeakValueResponseResultFlowRequestData) SetBandwidth(v string) *QueryTrafficRequestInTotalAndPeakValueResponseResultFlowRequestData {
  s.Bandwidth = &v
  return s
}

func (s *QueryTrafficRequestInTotalAndPeakValueResponseResultFlowRequestData) SetRequest(v string) *QueryTrafficRequestInTotalAndPeakValueResponseResultFlowRequestData {
  s.Request = &v
  return s
}

type QueryTrafficRequestInTotalAndPeakValuePaths struct {
}

func (s QueryTrafficRequestInTotalAndPeakValuePaths) String() string {
  return tea.Prettify(s)
}

func (s QueryTrafficRequestInTotalAndPeakValuePaths) GoString() string {
  return s.String()
}

type QueryTrafficRequestInTotalAndPeakValueParameters struct {
}

func (s QueryTrafficRequestInTotalAndPeakValueParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryTrafficRequestInTotalAndPeakValueParameters) GoString() string {
  return s.String()
}

type QueryTrafficRequestInTotalAndPeakValueRequestHeader struct {
}

func (s QueryTrafficRequestInTotalAndPeakValueRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryTrafficRequestInTotalAndPeakValueRequestHeader) GoString() string {
  return s.String()
}

type QueryTrafficRequestInTotalAndPeakValueResponseHeader struct {
}

func (s QueryTrafficRequestInTotalAndPeakValueResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryTrafficRequestInTotalAndPeakValueResponseHeader) GoString() string {
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




type QueryDomainTotalTrafficRequest struct {
  // {"en":"Domain list.
  //   Domain number limits can be adjusted depending on different accounts. The default value is 20(if you want to adjust,please, contact technical support)", "zh_CN":"域名列表
  //   1.域名个数限制根据账号可调,默认为20个(可联系技术支持下单调整);"}
  QueryDomainTotalTrafficDomainList *QueryDomainTotalTrafficDomainList `json:"domain-list,omitempty" xml:"domain-list,omitempty" require:"true"`
}

func (s QueryDomainTotalTrafficRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryDomainTotalTrafficRequest) GoString() string {
  return s.String()
}

func (s *QueryDomainTotalTrafficRequest) SetDomainList(v *QueryDomainTotalTrafficDomainList) *QueryDomainTotalTrafficRequest {
  s.QueryDomainTotalTrafficDomainList = v
  return s
}

type QueryDomainTotalTrafficDomainList struct {
  // {"en":"Domain", "zh_CN":"域名"}
  DomainName []*string `json:"domain-name,omitempty" xml:"domain-name,omitempty" require:"true" type:"Repeated"`
}

func (s QueryDomainTotalTrafficDomainList) String() string {
  return tea.Prettify(s)
}

func (s QueryDomainTotalTrafficDomainList) GoString() string {
  return s.String()
}

func (s *QueryDomainTotalTrafficDomainList) SetDomainName(v []*string) *QueryDomainTotalTrafficDomainList {
  s.DomainName = v
  return s
}

type QueryDomainTotalTrafficResponse struct {
  // {"en":"Total traffic. Keep two digits of decimals. Unit: MB", "zh_CN":"总流量,保留2位小数,单位为MB"}
  FlowSummary *int `json:"flow-summary,omitempty" xml:"flow-summary,omitempty" require:"true"`
  // {"en":"flowData", "zh_CN":"流量数据"}
  FlowData []*QueryDomainTotalTrafficResponseFlowData `json:"flow-data,omitempty" xml:"flow-data,omitempty" require:"true" type:"Repeated"`
}

func (s QueryDomainTotalTrafficResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryDomainTotalTrafficResponse) GoString() string {
  return s.String()
}

func (s *QueryDomainTotalTrafficResponse) SetFlowSummary(v int) *QueryDomainTotalTrafficResponse {
  s.FlowSummary = &v
  return s
}

func (s *QueryDomainTotalTrafficResponse) SetFlowData(v []*QueryDomainTotalTrafficResponseFlowData) *QueryDomainTotalTrafficResponse {
  s.FlowData = v
  return s
}

type QueryDomainTotalTrafficResponseFlowData struct     {
  // {"en":"Date
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

func (s QueryDomainTotalTrafficResponseFlowData) String() string {
  return tea.Prettify(s)
}

func (s QueryDomainTotalTrafficResponseFlowData) GoString() string {
  return s.String()
}

func (s *QueryDomainTotalTrafficResponseFlowData) SetTimestamp(v string) *QueryDomainTotalTrafficResponseFlowData {
  s.Timestamp = &v
  return s
}

func (s *QueryDomainTotalTrafficResponseFlowData) SetFlow(v int) *QueryDomainTotalTrafficResponseFlowData {
  s.Flow = &v
  return s
}

type QueryDomainTotalTrafficPaths struct {
}

func (s QueryDomainTotalTrafficPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryDomainTotalTrafficPaths) GoString() string {
  return s.String()
}

type QueryDomainTotalTrafficParameters struct {
  // {"en":"Start time
  // 1.The format is yyyy-MM-ddTHH:mm:ss+08:00;
  // 2.And smaller than the current time and 'dateTo';
  // 3.Period between 'dataFrom' and 'dateTo' cannot be longer than 31 days", "zh_CN":"开始时间
  // 1.格式为yyyy-MM-ddTHH:mm:ss+08:00;
  // 2.并且小于当前时间和dateTo;
  // 3.dateFrom和dateTo相差不能超过31天;4.只能查询最近2年内数据。"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty" require:"true"`
  // {"en":"End time
  // 1.The format is yyyy-MM-ddTHH:mm:ss+08:00;
  // 2.Must be greater than 'dateFrom';
  // 3.If it's greater than the current time, then the current time is assigned as the value", "zh_CN":"结束时间
  // 1.格式为yyyy-MM-ddTHH:mm:ss+08:00;
  // 2.必须大于dateFrom;
  // 3.如果大于当前时间,则重新赋值为当前时间;"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty" require:"true"`
  // {"en":"Data granularity
  // 1.fiveminutes: five minutes, hourly: one hour, daily: one day;
  // 2.If not specified, daily is set as the default value;
  // 3.If fiveminutes is specified as the value, then data is returned in actual configured granularity when there is specific configuration to data collecting granularity for the customer.", "zh_CN":"数据粒度
  // 1.fiveminutes:5分钟,hourly:1小时,daily:1天;
  // 2.不传递,默认为daily;
  // 3.传递fiveminutes时,若客户数据采集粒度有特殊配置将按实际配置粒度返回。"}
  Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s QueryDomainTotalTrafficParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryDomainTotalTrafficParameters) GoString() string {
  return s.String()
}

func (s *QueryDomainTotalTrafficParameters) SetDateFrom(v string) *QueryDomainTotalTrafficParameters {
  s.DateFrom = &v
  return s
}

func (s *QueryDomainTotalTrafficParameters) SetDateTo(v string) *QueryDomainTotalTrafficParameters {
  s.DateTo = &v
  return s
}

func (s *QueryDomainTotalTrafficParameters) SetType(v string) *QueryDomainTotalTrafficParameters {
  s.Type = &v
  return s
}

type QueryDomainTotalTrafficRequestHeader struct {
}

func (s QueryDomainTotalTrafficRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryDomainTotalTrafficRequestHeader) GoString() string {
  return s.String()
}

type QueryDomainTotalTrafficResponseHeader struct {
}

func (s QueryDomainTotalTrafficResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryDomainTotalTrafficResponseHeader) GoString() string {
  return s.String()
}




type QueryRequestNumberAndPeakBandwidthForMultiDomainRequest struct {
  // {"en":"Start time:
  // 1. Time format is yyyy-MM-ddTHH:mm:ss+08:00, for example, 2016-12-02T10:00:00+08:00 (10:00 on 2nd of December 2016, Beijing Time);
  // 2. No bigger than the current time.
  // 3. Data in the last 183 days at most can be queried.", "zh_CN":"开始时间：
  // 1.时间格式为yyyy-MM-ddTHH:mm:ss+08:00，例如，2016-12-02T10:00:00+08:00（为北京时间2016年12月2日10点0分0秒）；
  // 2.不能大于当前时间
  // 3.最多可获取最近半年（183天）的数据。"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty"`
  // {"en":"End time:
  // 1. the time format is 2016-12-02T10:00:00+08:00
  // 2. End time should be greater than start time. If the end time is greater than current time, current time will be used.
  // 3. If both fields of dataFrom and dateTo are left empty, then data in the last 24 hours will be queried by default; if only one field is filled in and one is left empty, then exception will be occur.
  // 4. Allowable maximum time range for query: 1 day, means the period between dateFrom to dateTo should not exceed 1 day (can be adjusted by contacting technical support up to 31 days).", "zh_CN":"结束时间：
  // 1.时间格式2016-12-02T10:00:00+08:00
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




type QueryBacktoOriginTrafficAndRequestRequest struct {
  // {'en':'Start time: 
  // 1. Start time: time format is yyyy-MM-ddTHH:mm:ss+08:00, for example, 2016-12-02T10:00:00+08:00 (December 2rd, 2016, 10:00 a.m., Beijing Time); 
  // 2. Not greater than the current time 
  // 3. The data can be queried is the last week (183 days).', 'zh_CN':'开始时间：
  // 1.时间格式为yyyy-MM-ddTHH:mm:ss+08:00，例如，2016-12-02T10:00:00+08:00（为北京时间2016年12月2日10点0分0秒）；
  // 2.不能大于当前时间
  // 3.最多可获取最近半年（183天）的数据。'}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty"`
  // {'en':'End time: 
  // 1. The time format is 2016-12-02T10:00:00+08:00 
  // 2. End time should be greater than start time. If the end time is greater than current time, current time will be used. 
  // 3. If both fields of dataFrom and dateTo are left empty, then data in the last 24 hours will be queried by default; if one field is filled and one is left empty, then exception will occur. 
  // 4. Maximum query time interval allowed: 1 day by default, that is, the difference between dateFrom and dateTo cannot exceed 1 day (you can contact technical support to adjust it, the maximum adjustment is 31 days)', 'zh_CN':'结束时间：
  // 1.时间格式2016-12-02T10:00:00+08:00
  // 2.结束时间需大于开始时间，结束时间如果大于当前时间，取当前时间。
  // 4.dateFrom，dateTo二者都未传，默认查询过去的24小时；如仅有一个未传，抛异常
  // 5.允许查询最大时间间隔：默认1天，即dateFrom和dateTo相差不能超过1天（可联系技术支持调整，最大调整到31天）。'}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty"`
  // {'en':'Data granularity: 
  // 1. Support 5m (granularity of 5 minutes) and 1m
  // 2. 5m by default if the value is empty', 'zh_CN':'数据粒度：
  // 1、支持5m（5分钟）。和 1m（1分钟）
  // 2、不传默认5m。'}
  DataInterval *string `json:"dataInterval,omitempty" xml:"dataInterval,omitempty"`
  // {'en':'Domain name: 
  // 1. The default upper limit of domains that can be entered is 200 (if you want to adjust, please, contact technical support); 
  // 2. All domains under the account will be queried if this input parameter is not specified. But if the number of domains under the account exceeds the limits, no query will be executed (Error)', 'zh_CN':'域名：
  // 1、可传递域名数量上限默认为200个（可联系技术支持调整）；
  // 2、未传递该入参时查询账号下所有域名，但当账号下域名数量超过限制时不可查询（报错）。'}
  Domain []*string `json:"domain,omitempty" xml:"domain,omitempty" type:"Repeated"`
  // {'en':'Grouping keywords:
  // 1. By default, data will be displayed by group;
  // 2. If there are keywords entered, value details shall be displayed by keywords; If groupBy is specified as domain, it means the results are returned according to domains. 
  // 3. Only domain can be specified', 'zh_CN':'分组关键词：
  // 1、默认聚合展示；
  // &nbsp; 2、传入关键词则代表需要按照关键词对应的值展示明细；
  // 例如groupBy传domain，则代表返回按照domain明细展开。
  // 3、只能传递domain。'}
  GroupBy []*string `json:"groupBy,omitempty" xml:"groupBy,omitempty" type:"Repeated"`
  // {'en':'Optional values 0, 1. 
  // Input 0 returns all data, input 1 only returns source site data, default is0&nbsp;', 'zh_CN':'可选值 0、1 。入参 0 则返回全部数据，入参 1 则只返回回源站数据，默认为 0'}
  BacksrcOnly *int `json:"backsrcOnly,omitempty" xml:"backsrcOnly,omitempty"`
}

func (s QueryBacktoOriginTrafficAndRequestRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryBacktoOriginTrafficAndRequestRequest) GoString() string {
  return s.String()
}

func (s *QueryBacktoOriginTrafficAndRequestRequest) SetDateFrom(v string) *QueryBacktoOriginTrafficAndRequestRequest {
  s.DateFrom = &v
  return s
}

func (s *QueryBacktoOriginTrafficAndRequestRequest) SetDateTo(v string) *QueryBacktoOriginTrafficAndRequestRequest {
  s.DateTo = &v
  return s
}

func (s *QueryBacktoOriginTrafficAndRequestRequest) SetDataInterval(v string) *QueryBacktoOriginTrafficAndRequestRequest {
  s.DataInterval = &v
  return s
}

func (s *QueryBacktoOriginTrafficAndRequestRequest) SetDomain(v []*string) *QueryBacktoOriginTrafficAndRequestRequest {
  s.Domain = v
  return s
}

func (s *QueryBacktoOriginTrafficAndRequestRequest) SetGroupBy(v []*string) *QueryBacktoOriginTrafficAndRequestRequest {
  s.GroupBy = v
  return s
}

func (s *QueryBacktoOriginTrafficAndRequestRequest) SetBacksrcOnly(v int) *QueryBacktoOriginTrafficAndRequestRequest {
  s.BacksrcOnly = &v
  return s
}

type QueryBacktoOriginTrafficAndRequestResponse struct {
  Result []*QueryBacktoOriginTrafficAndRequestResponseResult `json:"result,omitempty" xml:"result,omitempty" require:"true" type:"Repeated"`
}

func (s QueryBacktoOriginTrafficAndRequestResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryBacktoOriginTrafficAndRequestResponse) GoString() string {
  return s.String()
}

func (s *QueryBacktoOriginTrafficAndRequestResponse) SetResult(v []*QueryBacktoOriginTrafficAndRequestResponseResult) *QueryBacktoOriginTrafficAndRequestResponse {
  s.Result = v
  return s
}

type QueryBacktoOriginTrafficAndRequestResponseResult struct     {
  // {'en':'Domain', 'zh_CN':'域名'}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {'en':'totalFlow, unit :MB, 2 decimal places reserved, example (74099.92)', 'zh_CN':'总流量，单位MB，保留两位小时'}
  TotalFlow *string `json:"totalFlow,omitempty" xml:"totalFlow,omitempty" require:"true"`
  // {'en':'totalRequest', 'zh_CN':'总请求数'}
  TotalRequest *string `json:"totalRequest,omitempty" xml:"totalRequest,omitempty" require:"true"`
  // {'en':'peak of Request', 'zh_CN':'请求数峰值'}
  PeakRequest *string `json:"peakRequest,omitempty" xml:"peakRequest,omitempty" require:"true"`
  // {'en':'peaktime of Request', 'zh_CN':'请求数峰值时间'}
  PeakRequestTime *string `json:"peakRequestTime,omitempty" xml:"peakRequestTime,omitempty" require:"true"`
  // {'en':'peakBandwidth, unit :Mbps, 2 decimal places reserved, example (74099.92)', 'zh_CN':'带宽峰值，单位Mbps，保留2位小数，示例 （931556.21）'}
  PeakBandwidth *string `json:"peakBandwidth,omitempty" xml:"peakBandwidth,omitempty" require:"true"`
  // {'en':'Peak time of Bandwidth', 'zh_CN':'带宽峰值时间'}
  PeakBandwidthTime *string `json:"peakBandwidthTime,omitempty" xml:"peakBandwidthTime,omitempty" require:"true"`
  FlowRequestOriginData []*QueryBacktoOriginTrafficAndRequestResponseResultFlowRequestOriginData `json:"flowRequestOriginData,omitempty" xml:"flowRequestOriginData,omitempty" require:"true" type:"Repeated"`
}

func (s QueryBacktoOriginTrafficAndRequestResponseResult) String() string {
  return tea.Prettify(s)
}

func (s QueryBacktoOriginTrafficAndRequestResponseResult) GoString() string {
  return s.String()
}

func (s *QueryBacktoOriginTrafficAndRequestResponseResult) SetDomain(v string) *QueryBacktoOriginTrafficAndRequestResponseResult {
  s.Domain = &v
  return s
}

func (s *QueryBacktoOriginTrafficAndRequestResponseResult) SetTotalFlow(v string) *QueryBacktoOriginTrafficAndRequestResponseResult {
  s.TotalFlow = &v
  return s
}

func (s *QueryBacktoOriginTrafficAndRequestResponseResult) SetTotalRequest(v string) *QueryBacktoOriginTrafficAndRequestResponseResult {
  s.TotalRequest = &v
  return s
}

func (s *QueryBacktoOriginTrafficAndRequestResponseResult) SetPeakRequest(v string) *QueryBacktoOriginTrafficAndRequestResponseResult {
  s.PeakRequest = &v
  return s
}

func (s *QueryBacktoOriginTrafficAndRequestResponseResult) SetPeakRequestTime(v string) *QueryBacktoOriginTrafficAndRequestResponseResult {
  s.PeakRequestTime = &v
  return s
}

func (s *QueryBacktoOriginTrafficAndRequestResponseResult) SetPeakBandwidth(v string) *QueryBacktoOriginTrafficAndRequestResponseResult {
  s.PeakBandwidth = &v
  return s
}

func (s *QueryBacktoOriginTrafficAndRequestResponseResult) SetPeakBandwidthTime(v string) *QueryBacktoOriginTrafficAndRequestResponseResult {
  s.PeakBandwidthTime = &v
  return s
}

func (s *QueryBacktoOriginTrafficAndRequestResponseResult) SetFlowRequestOriginData(v []*QueryBacktoOriginTrafficAndRequestResponseResultFlowRequestOriginData) *QueryBacktoOriginTrafficAndRequestResponseResult {
  s.FlowRequestOriginData = v
  return s
}

type QueryBacktoOriginTrafficAndRequestResponseResultFlowRequestOriginData struct     {
  // {'en':'1. When the querying data granularity is 5m, then the format is yyyy-MM-dd HH:mm; the data value of every time slice represents the data value within the previous time granularity range. The first time slice of the day is yyyy-MM-dd 12:05 AM, and the last one is (yyyy-MM-dd+1) 00:00. 
  //         2. Return the time slices that contained in start time and in end time.', 'zh_CN':'1.查询的数据粒度为5m时，格式为yyyy-MM-dd &nbsp; HH:mm；每一个时间片数据值代表的是前一个时间粒度范围内的数据值。一天开始的时间片是yyyy-MM-dd 00:05，最后一个时间片是（yyyy-MM-dd+1） &nbsp; 00:00。
  //         2.返回开始时间和结束时间包含的时间片。'}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {'en':'Traffic unit is MB and keep the number to two decimal places', 'zh_CN':'流量值，单位MB，保留2位小数；'}
  Flow *string `json:"flow,omitempty" xml:"flow,omitempty" require:"true"`
  // {'en':'Bandwidth, unit: Mbps, 2 decimal places reserved, example (931556.21)', 'zh_CN':'带宽，单位：Mbps，保留两位小数'}
  Bandwidth *string `json:"bandwidth,omitempty" xml:"bandwidth,omitempty" require:"true"`
  // {'en':'Total number of requests', 'zh_CN':'请求数'}
  Request *string `json:"request,omitempty" xml:"request,omitempty" require:"true"`
}

func (s QueryBacktoOriginTrafficAndRequestResponseResultFlowRequestOriginData) String() string {
  return tea.Prettify(s)
}

func (s QueryBacktoOriginTrafficAndRequestResponseResultFlowRequestOriginData) GoString() string {
  return s.String()
}

func (s *QueryBacktoOriginTrafficAndRequestResponseResultFlowRequestOriginData) SetTimestamp(v string) *QueryBacktoOriginTrafficAndRequestResponseResultFlowRequestOriginData {
  s.Timestamp = &v
  return s
}

func (s *QueryBacktoOriginTrafficAndRequestResponseResultFlowRequestOriginData) SetFlow(v string) *QueryBacktoOriginTrafficAndRequestResponseResultFlowRequestOriginData {
  s.Flow = &v
  return s
}

func (s *QueryBacktoOriginTrafficAndRequestResponseResultFlowRequestOriginData) SetBandwidth(v string) *QueryBacktoOriginTrafficAndRequestResponseResultFlowRequestOriginData {
  s.Bandwidth = &v
  return s
}

func (s *QueryBacktoOriginTrafficAndRequestResponseResultFlowRequestOriginData) SetRequest(v string) *QueryBacktoOriginTrafficAndRequestResponseResultFlowRequestOriginData {
  s.Request = &v
  return s
}

type QueryBacktoOriginTrafficAndRequestPaths struct {
}

func (s QueryBacktoOriginTrafficAndRequestPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryBacktoOriginTrafficAndRequestPaths) GoString() string {
  return s.String()
}

type QueryBacktoOriginTrafficAndRequestParameters struct {
}

func (s QueryBacktoOriginTrafficAndRequestParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryBacktoOriginTrafficAndRequestParameters) GoString() string {
  return s.String()
}

type QueryBacktoOriginTrafficAndRequestRequestHeader struct {
}

func (s QueryBacktoOriginTrafficAndRequestRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryBacktoOriginTrafficAndRequestRequestHeader) GoString() string {
  return s.String()
}

type QueryBacktoOriginTrafficAndRequestResponseHeader struct {
}

func (s QueryBacktoOriginTrafficAndRequestResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryBacktoOriginTrafficAndRequestResponseHeader) GoString() string {
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




type QueryRequesBandwidthtSavingRatioRequest struct {
  // {"en":"From date:
  // 1.The time format is yyyy-MM-ddTHH:MM:ss+08:00. For example :2019-01-01T10:00:00+08:00 (10:00 on December 2, 2018 10:00:00:00 UTC+8);
  // 2.Cannot exceed current time;
  // 3.The most recent six-month (183 days) data are available.", "zh_CN":"开始时间:
  // 1.时间格式为yyyy-MM-ddTHH:mm:ss+08:00,例如,2019-01-01T10:00:00+08:00(为北京时间2018年12月2日10点0分0秒);
  // 2.不能大于当前时间
  // 3.最多可获取最近半年(183天)的数据。"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty"`
  // {"en":"To time:
  // 1.The time format is yyyy-MM-ddTHH:MM:ss+08:00. For example :2019-01-01T10:00:00+08:00 (10:00 on December 2, 2018 10:00:00:00 UTC+8);
  // 2.The end time needs to be greater than the start time. If the end time is greater than the current time, take the current time;
  // 3.Date from, Date to both, the default query past 24 hours; If there is only one unsent, throw an exception;
  // 4.Maximum allowed query time interval: 31 days, Date from and dateTo, not more than 31 days", "zh_CN":"结束时间:
  // 1.时间格式2019-01-02T10:00:00+08:00
  // 2.结束时间需大于开始时间,结束时间如果大于当前时间,取当前时间。
  // 3.dateFrom,dateTo二者都未传,默认查询过去的24小时;如仅有一个未传,抛异常
  // 4.允许查询最大时间间隔:31天,即dateFrom和dateTo相差不能超过31天"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty"`
  // {"en":"Domain:
  // 1.The maximum number of deliverable domain names is 200 by default;
  // 2.Automatically filter out illegal domain names (pass illegal domain names, will be filtered out, the query results only return the data of legitimate domain names);
  // 3.The default query accounts for all domains if the number of domain names exceeds the upper limit when the entry is not delivered. If the number of domain names in the account exceeds the limit, an error is raised.", "zh_CN":"域名:
  // 1.可传递域名数量上限默认为200个
  // 2.自动过滤掉非法域名(如传递非法域名,会被过滤掉,查询结果只返回合法域名的数据)
  // 3.未传递该入参时,默认查询账号下所有域名,但当账号下域名数量超过上限时提示错误。"}
  Domain []*string `json:"domain,omitempty" xml:"domain,omitempty" require:"true" type:"Repeated"`
  // {"en":"Time interval of data: 5m (5 min), 1h (1 hour), 1d (1 day); 
  // 	The default is 1d.", "zh_CN":"数据粒度:
  // 1.支持5m(5分钟)、1h(1小时)、1d(天)
  // 2.不传默认1d。"}
  DataInterval *string `json:"dataInterval,omitempty" xml:"dataInterval,omitempty"`
}

func (s QueryRequesBandwidthtSavingRatioRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryRequesBandwidthtSavingRatioRequest) GoString() string {
  return s.String()
}

func (s *QueryRequesBandwidthtSavingRatioRequest) SetDateFrom(v string) *QueryRequesBandwidthtSavingRatioRequest {
  s.DateFrom = &v
  return s
}

func (s *QueryRequesBandwidthtSavingRatioRequest) SetDateTo(v string) *QueryRequesBandwidthtSavingRatioRequest {
  s.DateTo = &v
  return s
}

func (s *QueryRequesBandwidthtSavingRatioRequest) SetDomain(v []*string) *QueryRequesBandwidthtSavingRatioRequest {
  s.Domain = v
  return s
}

func (s *QueryRequesBandwidthtSavingRatioRequest) SetDataInterval(v string) *QueryRequesBandwidthtSavingRatioRequest {
  s.DataInterval = &v
  return s
}

type QueryRequesBandwidthtSavingRatioResponse struct {
  // {"en":"request result status code", "zh_CN":"请求结果状态码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Request result information", "zh_CN":"请求结果信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  Data []*QueryRequesBandwidthtSavingRatioResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s QueryRequesBandwidthtSavingRatioResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryRequesBandwidthtSavingRatioResponse) GoString() string {
  return s.String()
}

func (s *QueryRequesBandwidthtSavingRatioResponse) SetCode(v string) *QueryRequesBandwidthtSavingRatioResponse {
  s.Code = &v
  return s
}

func (s *QueryRequesBandwidthtSavingRatioResponse) SetMessage(v string) *QueryRequesBandwidthtSavingRatioResponse {
  s.Message = &v
  return s
}

func (s *QueryRequesBandwidthtSavingRatioResponse) SetData(v []*QueryRequesBandwidthtSavingRatioResponseData) *QueryRequesBandwidthtSavingRatioResponse {
  s.Data = v
  return s
}

type QueryRequesBandwidthtSavingRatioResponseData struct     {
  // {"en":"Actually processed time.  yyyy-MM-dd HH:mm format", "zh_CN":"实际查询时间,格式 yyyy-MM-dd HH:mm"}
  RealDate *string `json:"realDate,omitempty" xml:"realDate,omitempty" require:"true"`
  // {"en":"Average of total saving of bandwidth.", "zh_CN":"总节省带宽的平均值"}
  TotalAvg *int `json:"totalAvg,omitempty" xml:"totalAvg,omitempty" require:"true"`
  SavingBandwidthDatas []*QueryRequesBandwidthtSavingRatioResponseDataSavingBandwidthDatas `json:"savingBandwidthDatas,omitempty" xml:"savingBandwidthDatas,omitempty" require:"true" type:"Repeated"`
}

func (s QueryRequesBandwidthtSavingRatioResponseData) String() string {
  return tea.Prettify(s)
}

func (s QueryRequesBandwidthtSavingRatioResponseData) GoString() string {
  return s.String()
}

func (s *QueryRequesBandwidthtSavingRatioResponseData) SetRealDate(v string) *QueryRequesBandwidthtSavingRatioResponseData {
  s.RealDate = &v
  return s
}

func (s *QueryRequesBandwidthtSavingRatioResponseData) SetTotalAvg(v int) *QueryRequesBandwidthtSavingRatioResponseData {
  s.TotalAvg = &v
  return s
}

func (s *QueryRequesBandwidthtSavingRatioResponseData) SetSavingBandwidthDatas(v []*QueryRequesBandwidthtSavingRatioResponseDataSavingBandwidthDatas) *QueryRequesBandwidthtSavingRatioResponseData {
  s.SavingBandwidthDatas = v
  return s
}

type QueryRequesBandwidthtSavingRatioResponseDataSavingBandwidthDatas struct     {
  // {"en":"timetamp
  // 1. When the data granularity of the query is fiveminutes, the format is yyyy-MM-dd HH:MM; Each time slice data value represents the data value in the previous time granularity range, For example yyyy-MM-dd 00:05 represents data in the range from 00:00 to 00:05.
  // 2.The data granularity of query is hourly, the format is yyyy-MM-dd HH. Each time slice data value represents data values in the previous time granularity range such as yyyy-MM-dd 01 that represent data from 00 to 01.
  // 3. the data granularity of the query is daily, the format is yyyy-MM-dd; Each time slice data value represents the data value for that day.
  // 4.Returns the timetamp contained in start time and end time.", "zh_CN":"时间片
  // 1.查询的数据粒度为fiveminutes时,格式为yyyy-MM-dd HH:mm;每一个时间片数据值代表的是前一个时间粒度范围内的数据值,比如yyyy-MM-dd 00:05,代表00:00到00:05范围内的数据。
  // 2.查询的数据粒度为hourly时,格式为yyyy-MM-dd HH;每一个时间片数据值代表的是前一个时间粒度范围内的数据值,比如yyyy-MM-dd 01,代表00到01之间的数据。
  // 3.查询的数据粒度为daily时,格式为yyyy-MM-dd;每一个时间片数据值代表的该天内的数据值。
  // 4.返回开始时间和结束时间包含的时间片。"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {"en":"saving of bandwidth,keep 4 decimal places", "zh_CN":"节省带宽,保留4位小数"}
  SavingBandwidth *string `json:"savingBandwidth,omitempty" xml:"savingBandwidth,omitempty" require:"true"`
}

func (s QueryRequesBandwidthtSavingRatioResponseDataSavingBandwidthDatas) String() string {
  return tea.Prettify(s)
}

func (s QueryRequesBandwidthtSavingRatioResponseDataSavingBandwidthDatas) GoString() string {
  return s.String()
}

func (s *QueryRequesBandwidthtSavingRatioResponseDataSavingBandwidthDatas) SetTimestamp(v string) *QueryRequesBandwidthtSavingRatioResponseDataSavingBandwidthDatas {
  s.Timestamp = &v
  return s
}

func (s *QueryRequesBandwidthtSavingRatioResponseDataSavingBandwidthDatas) SetSavingBandwidth(v string) *QueryRequesBandwidthtSavingRatioResponseDataSavingBandwidthDatas {
  s.SavingBandwidth = &v
  return s
}

type QueryRequesBandwidthtSavingRatioPaths struct {
}

func (s QueryRequesBandwidthtSavingRatioPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryRequesBandwidthtSavingRatioPaths) GoString() string {
  return s.String()
}

type QueryRequesBandwidthtSavingRatioParameters struct {
}

func (s QueryRequesBandwidthtSavingRatioParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryRequesBandwidthtSavingRatioParameters) GoString() string {
  return s.String()
}

type QueryRequesBandwidthtSavingRatioRequestHeader struct {
}

func (s QueryRequesBandwidthtSavingRatioRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryRequesBandwidthtSavingRatioRequestHeader) GoString() string {
  return s.String()
}

type QueryRequesBandwidthtSavingRatioResponseHeader struct {
}

func (s QueryRequesBandwidthtSavingRatioResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryRequesBandwidthtSavingRatioResponseHeader) GoString() string {
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




type QueryTrafficBySpecificProtocolRequest struct {
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
  // {"en":"End time:
  // 1.The format is yyyy-MM-ddTHH:mm:ss+08:00;
  // 2.Must be greater than dateFrom;
  // 3.If it's greater than the current time, then the current time is assigned as the value;", "zh_CN":"结束时间
  // 1.格式为yyyy-MM-ddTHH:mm:ss+08:00;
  // 2.必须大于dateFrom;
  // 3.如果大于当前时间,则重新赋值为当前时间;"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty"`
  // {"en":"Domain names, domain number limits can be adjusted depending on different accounts. The default value is 20", "zh_CN":"域名,域名个数限制根据账号可调,默认为20个"}
  Domain []*string `json:"domain,omitempty" xml:"domain,omitempty" require:"true" type:"Repeated"`
  // {"en":"Data granularity:
  // Support for 1m(1 minutes), 5m (5 minutes), 1h (1 hour), 1d (1 day)", "zh_CN":"数据粒度:
  // 支持1m(1分钟),5m(5分钟),1h(1小时),1d(1天)"}
  DataInterval *string `json:"dataInterval,omitempty" xml:"dataInterval,omitempty"`
  // {"en":"Transmission protocol:
  // 1.Options: http, https;
  // 2.https is used as the default value is no value specified;
  // 3.httpFlowData is displayed if http is queried, and httpsFlowData is displayed if https is queried;", "zh_CN":"传输协议
  // 1.可选值为http、https;
  // 2.不传默认查询https;
  // 3.查询http时出参展示httpFlowData,查询https时出参展示httpsFlowData;"}
  ProtocolType *string `json:"protocolType,omitempty" xml:"protocolType,omitempty"`
  // {"en":"Group dimension:
  // 1.The value can be selected is domain;
  // 2.The data is displayed according to the specified dimension;", "zh_CN":"分组维度
  // 1.可选值为domain;
  // 2.有传入则按照该维度展示明细数据;"}
  GroupBy []*string `json:"groupBy,omitempty" xml:"groupBy,omitempty" type:"Repeated"`
}

func (s QueryTrafficBySpecificProtocolRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryTrafficBySpecificProtocolRequest) GoString() string {
  return s.String()
}

func (s *QueryTrafficBySpecificProtocolRequest) SetDateFrom(v string) *QueryTrafficBySpecificProtocolRequest {
  s.DateFrom = &v
  return s
}

func (s *QueryTrafficBySpecificProtocolRequest) SetDateTo(v string) *QueryTrafficBySpecificProtocolRequest {
  s.DateTo = &v
  return s
}

func (s *QueryTrafficBySpecificProtocolRequest) SetDomain(v []*string) *QueryTrafficBySpecificProtocolRequest {
  s.Domain = v
  return s
}

func (s *QueryTrafficBySpecificProtocolRequest) SetDataInterval(v string) *QueryTrafficBySpecificProtocolRequest {
  s.DataInterval = &v
  return s
}

func (s *QueryTrafficBySpecificProtocolRequest) SetProtocolType(v string) *QueryTrafficBySpecificProtocolRequest {
  s.ProtocolType = &v
  return s
}

func (s *QueryTrafficBySpecificProtocolRequest) SetGroupBy(v []*string) *QueryTrafficBySpecificProtocolRequest {
  s.GroupBy = v
  return s
}

type QueryTrafficBySpecificProtocolResponse struct {
  // {"en":"result", "zh_CN":"结果"}
  Result []*QueryTrafficBySpecificProtocolResponseResult `json:"result,omitempty" xml:"result,omitempty" require:"true" type:"Repeated"`
}

func (s QueryTrafficBySpecificProtocolResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryTrafficBySpecificProtocolResponse) GoString() string {
  return s.String()
}

func (s *QueryTrafficBySpecificProtocolResponse) SetResult(v []*QueryTrafficBySpecificProtocolResponseResult) *QueryTrafficBySpecificProtocolResponse {
  s.Result = v
  return s
}

type QueryTrafficBySpecificProtocolResponseResult struct     {
  // {"en":"Domain", "zh_CN":"域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  HttpsFlowData []*QueryTrafficBySpecificProtocolResponseResultHttpsFlowData `json:"httpsFlowData,omitempty" xml:"httpsFlowData,omitempty" require:"true" type:"Repeated"`
}

func (s QueryTrafficBySpecificProtocolResponseResult) String() string {
  return tea.Prettify(s)
}

func (s QueryTrafficBySpecificProtocolResponseResult) GoString() string {
  return s.String()
}

func (s *QueryTrafficBySpecificProtocolResponseResult) SetDomain(v string) *QueryTrafficBySpecificProtocolResponseResult {
  s.Domain = &v
  return s
}

func (s *QueryTrafficBySpecificProtocolResponseResult) SetHttpsFlowData(v []*QueryTrafficBySpecificProtocolResponseResultHttpsFlowData) *QueryTrafficBySpecificProtocolResponseResult {
  s.HttpsFlowData = v
  return s
}

type QueryTrafficBySpecificProtocolResponseResultHttpsFlowData struct     {
  // {"en":"DateTime: the format is yyyy-MM-dd HH:mm; the data value of every time slice represents the data value within the previous time granularity range. The first time slice of the day is yyyy-MM-dd 12:05 AM, and the last one is (yyyy-MM-dd+1) 00:00.", "zh_CN":"时间,格式为yyyy-MM-dd HH:mm;每一个时间片数据值代表的是前一个时间粒度范围内的数据值。一天开始的时间片是yyyy-MM-dd 00:05,最后一个时间片是(yyyy-MM-dd+1)00:00。"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {"en":"Traffic unit is MB and 2 digits  of decimals allowed", "zh_CN":"流量值,单位MB,保留2位小数"}
  Value *string `json:"value,omitempty" xml:"value,omitempty" require:"true"`
}

func (s QueryTrafficBySpecificProtocolResponseResultHttpsFlowData) String() string {
  return tea.Prettify(s)
}

func (s QueryTrafficBySpecificProtocolResponseResultHttpsFlowData) GoString() string {
  return s.String()
}

func (s *QueryTrafficBySpecificProtocolResponseResultHttpsFlowData) SetTimestamp(v string) *QueryTrafficBySpecificProtocolResponseResultHttpsFlowData {
  s.Timestamp = &v
  return s
}

func (s *QueryTrafficBySpecificProtocolResponseResultHttpsFlowData) SetValue(v string) *QueryTrafficBySpecificProtocolResponseResultHttpsFlowData {
  s.Value = &v
  return s
}

type QueryTrafficBySpecificProtocolPaths struct {
}

func (s QueryTrafficBySpecificProtocolPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryTrafficBySpecificProtocolPaths) GoString() string {
  return s.String()
}

type QueryTrafficBySpecificProtocolParameters struct {
}

func (s QueryTrafficBySpecificProtocolParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryTrafficBySpecificProtocolParameters) GoString() string {
  return s.String()
}

type QueryTrafficBySpecificProtocolRequestHeader struct {
}

func (s QueryTrafficBySpecificProtocolRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryTrafficBySpecificProtocolRequestHeader) GoString() string {
  return s.String()
}

type QueryTrafficBySpecificProtocolResponseHeader struct {
}

func (s QueryTrafficBySpecificProtocolResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryTrafficBySpecificProtocolResponseHeader) GoString() string {
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




type QueryISPProvinceHitRateRequest struct {
  // {"en":"Start date:
  // 1. The time format is yyyy-MM-ddTHH:MM:ss+08:00. For example, 2019-01-01T10:00:00+08:00
  // 2. Cannot exceed current time
  // 3. The most recent six-month (183 days) data are available.", "zh_CN":"开始时间:
  // 1.时间格式为yyyy-MM-ddTHH:mm:ss+08:00，例如，2019-01-01T10:00:00+08:00（为北京时间2019年01月01日10点0分0秒）;
  // 2.不能大于当前时间
  // 3.最多可获取最近半年（183天）的数据。"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty"`
  // {"en":"End time:
  // 1. The time format is yyyy-MM-ddTHH:MM:ss+08:00. For example, 2019-01-01T10:00:00+08:00
  // 2. The end time needs to be greater than the start time. If the end time is greater than the current time, take the current time.
  // 3. Date from, Date to both, the default query past 24 hours; If there is only one unsent, throw an exception
  // 4. Maximum allowed query time interval: 24 hours (with technical support adjustments), meaning that the difference between Date from and dateTo cannot exceed 24 hours.", "zh_CN":"结束时间:
  // 1.时间格式为yyyy-MM-ddTHH:mm:ss+08:00，例如 2016-12-02T10:00:00+08:00
  // 2.结束时间需大于开始时间，结束时间如果大于当前时间，取当前时间。
  // 3.dateFrom，dateTo二者都未传，默认查询过去的24小时;如仅有一个未传，抛异常
  // 4.允许查询最大时间间隔:24小时（可联系技术支持调整），即dateFrom和dateTo相差不能超过24小时。"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty"`
  // {"en":"Domain name:
  // 1. The maximum number of TLDs that can be delivered is 20 by default (contact technical support adjustment);
  // 2. Automatically filter out illegal domain names (pass illegal domain names, will be filtered out, the query results only return the data of legitimate domain names)
  // 3. Domain name exceeding limit, misstatement", "zh_CN":"域名:
  // 1.可传递域名数量上限默认为20个（可联系技术支持调整）;
  // 2.自动过滤掉非法域名（如传递非法域名，会被过滤掉，查询结果只返回合法域名的数据）
  // 3.域名超过上限，提示错误"}
  Domain []*string `json:"domain,omitempty" xml:"domain,omitempty" type:"Repeated"`
  // {"en":"Data granularity: 
  // 1. default 1m
  // 2. 1m (1 minute), 5m (5 minutes)", "zh_CN":"数据粒度:
  // 1.不传默认1m
  // 2.支持1m（1分钟）、5m（5分钟）"}
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
  // {"en":"query dimensionality:
  // 1. Optional value flow, request
  // 2. Default flow
  // 3. Flow: Flow, keep two decimal places;
  // 4. Request: number of Request", "zh_CN":"查询维度:
  // 1.可选值 flow、request
  // 2.传默认 flow
  // 3.flow:流量，保留两位小数;
  // 4.request:请求数"}
  QueryBy *string `json:"queryBy,omitempty" xml:"queryBy,omitempty"`
  // {"en":"Optional: domain, all, If it is empty, it defaults to returning by domain dimension;
  // If all is passed, merge and return according to the query domain name.", "zh_CN":"可选项：domain、all, 为空则默认为按domain维度返回;
  // 若传递all，则按查询域名合并返回"}
  GroupBy *string `json:"groupBy,omitempty" xml:"groupBy,omitempty"`
}

func (s QueryISPProvinceHitRateRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryISPProvinceHitRateRequest) GoString() string {
  return s.String()
}

func (s *QueryISPProvinceHitRateRequest) SetDateFrom(v string) *QueryISPProvinceHitRateRequest {
  s.DateFrom = &v
  return s
}

func (s *QueryISPProvinceHitRateRequest) SetDateTo(v string) *QueryISPProvinceHitRateRequest {
  s.DateTo = &v
  return s
}

func (s *QueryISPProvinceHitRateRequest) SetDomain(v []*string) *QueryISPProvinceHitRateRequest {
  s.Domain = v
  return s
}

func (s *QueryISPProvinceHitRateRequest) SetDataInterval(v string) *QueryISPProvinceHitRateRequest {
  s.DataInterval = &v
  return s
}

func (s *QueryISPProvinceHitRateRequest) SetProvince(v []*string) *QueryISPProvinceHitRateRequest {
  s.Province = v
  return s
}

func (s *QueryISPProvinceHitRateRequest) SetIsp(v []*string) *QueryISPProvinceHitRateRequest {
  s.Isp = v
  return s
}

func (s *QueryISPProvinceHitRateRequest) SetQueryBy(v string) *QueryISPProvinceHitRateRequest {
  s.QueryBy = &v
  return s
}

func (s *QueryISPProvinceHitRateRequest) SetGroupBy(v string) *QueryISPProvinceHitRateRequest {
  s.GroupBy = &v
  return s
}

type QueryISPProvinceHitRateResponse struct {
  // {"en":"request result status code", "zh_CN":"请求结果状态码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Request result information", "zh_CN":"请求结果信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Detailed data on the results of the request", "zh_CN":"请求结果的详细数据"}
  Data []*QueryISPProvinceHitRateResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s QueryISPProvinceHitRateResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryISPProvinceHitRateResponse) GoString() string {
  return s.String()
}

func (s *QueryISPProvinceHitRateResponse) SetCode(v string) *QueryISPProvinceHitRateResponse {
  s.Code = &v
  return s
}

func (s *QueryISPProvinceHitRateResponse) SetMessage(v string) *QueryISPProvinceHitRateResponse {
  s.Message = &v
  return s
}

func (s *QueryISPProvinceHitRateResponse) SetData(v []*QueryISPProvinceHitRateResponseData) *QueryISPProvinceHitRateResponse {
  s.Data = v
  return s
}

type QueryISPProvinceHitRateResponseData struct     {
  // {"en":"domain", "zh_CN":"域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  DetailList []*QueryISPProvinceHitRateResponseDataDetailList `json:"detailList,omitempty" xml:"detailList,omitempty" require:"true" type:"Repeated"`
}

func (s QueryISPProvinceHitRateResponseData) String() string {
  return tea.Prettify(s)
}

func (s QueryISPProvinceHitRateResponseData) GoString() string {
  return s.String()
}

func (s *QueryISPProvinceHitRateResponseData) SetDomain(v string) *QueryISPProvinceHitRateResponseData {
  s.Domain = &v
  return s
}

func (s *QueryISPProvinceHitRateResponseData) SetDetailList(v []*QueryISPProvinceHitRateResponseDataDetailList) *QueryISPProvinceHitRateResponseData {
  s.DetailList = v
  return s
}

type QueryISPProvinceHitRateResponseDataDetailList struct     {
  // {"en":"time, in yyyy-MM-dd HH:MM", "zh_CN":"时间，格式为yyyy-MM-dd HH:mm"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {"en":"Hit data:
  // Flow: Flow, keep two decimal places;
  // Request: number of Request", "zh_CN":"命中数据:
  // 1.flow:流量，保留两位小数;
  // 2.request:请求数"}
  HitValue *string `json:"hitValue,omitempty" xml:"hitValue,omitempty" require:"true"`
  // {"en":"Hit rate, keep four decimal places", "zh_CN":"命中率，保留四位小数"}
  HitRate *string `json:"hitRate,omitempty" xml:"hitRate,omitempty" require:"true"`
}

func (s QueryISPProvinceHitRateResponseDataDetailList) String() string {
  return tea.Prettify(s)
}

func (s QueryISPProvinceHitRateResponseDataDetailList) GoString() string {
  return s.String()
}

func (s *QueryISPProvinceHitRateResponseDataDetailList) SetTimestamp(v string) *QueryISPProvinceHitRateResponseDataDetailList {
  s.Timestamp = &v
  return s
}

func (s *QueryISPProvinceHitRateResponseDataDetailList) SetHitValue(v string) *QueryISPProvinceHitRateResponseDataDetailList {
  s.HitValue = &v
  return s
}

func (s *QueryISPProvinceHitRateResponseDataDetailList) SetHitRate(v string) *QueryISPProvinceHitRateResponseDataDetailList {
  s.HitRate = &v
  return s
}

type QueryISPProvinceHitRatePaths struct {
}

func (s QueryISPProvinceHitRatePaths) String() string {
  return tea.Prettify(s)
}

func (s QueryISPProvinceHitRatePaths) GoString() string {
  return s.String()
}

type QueryISPProvinceHitRateParameters struct {
}

func (s QueryISPProvinceHitRateParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryISPProvinceHitRateParameters) GoString() string {
  return s.String()
}

type QueryISPProvinceHitRateRequestHeader struct {
}

func (s QueryISPProvinceHitRateRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryISPProvinceHitRateRequestHeader) GoString() string {
  return s.String()
}

type QueryISPProvinceHitRateResponseHeader struct {
}

func (s QueryISPProvinceHitRateResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryISPProvinceHitRateResponseHeader) GoString() string {
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




