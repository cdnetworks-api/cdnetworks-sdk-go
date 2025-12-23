package reportvisitor

// This file is auto-generated, don't edit it. Thanks.
import (
  "github.com/alibabacloud-go/tea/tea"
)

type ReportVisitorCustomTopDailyServiceRequest struct {
  // {"en":"Start time:
  // The time format is yyyy-MM-dd, for example, 2021-10-10;
  // It cannot be greater than the current time
  // You can get data for the last three months (90 days) at most.", "zh_CN":"开始时间：
  // 时间格式为yyyy-MM-dd，例如，2021-10-10；
  // 不能大于当前时间
  // 最多可获取最近三个月（90天）的数据。"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty"`
  // {"en":"End time:
  // 1. The time format is yyyy-MM-dd, the end time must be greater than the start time. If the end time is greater than the current time, the current time is used.
  // 2. If both dateFrom and dateTo are not transmitted, the default query is the past 7 days. If only one is not transmitted, an exception is thrown;
  // 3. The maximum query time interval allowed is 31 days, that is, the difference between dateFrom and dateTo cannot exceed 31 days (you can contact technical support to adjust)", "zh_CN":"结束时间：
  // 时间格式为yyyy-MM-dd
  // 结束时间需大于开始时间，结束时间如果大于当前时间，取当前时间。
  // dateFrom，dateTo二者都未传，默认查询过去的7天；如仅有一个未传，抛异常
  // 允许查询最大时间间隔：31天，即dateFrom和dateTo相差不能超过31天（可联系技术支持调整）。"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty"`
  // {"en":"Domain name:
  // 1. The upper limit of the number of domain names that can be passed is 100 by default (you can contact technical support to adjust it);
  // 2. Automatically filter out illegal domain names (if an illegal domain name is passed, it will be filtered out, and the query result will only return data for legal domain names);
  // 3. When this parameter is not passed, all domain names under the account will be queried by default, but an error will be prompted when the number of domain names under the account exceeds the upper limit.", "zh_CN":"域名：
  // 可传递域名数量上限默认为100个（可联系技术支持调整）。
  // 自动过滤掉非法域名（如传递非法域名，会被过滤掉，查询结果只返回合法域名的数据）
  // 未传递该入参时，默认查询账号下所有域名，但当账号下域名数量超过上限时提示错误"}
  Domain []*string `json:"domain,omitempty" xml:"domain,omitempty" type:"Repeated"`
  // {"en":"Number of TOPs:
  // 1. If not specified, the default is TOP 10;
  // 2. The maximum number of TOPs is 100.", "zh_CN":"TOP个数：
  // 不传默认TOP 10
  // 最大TOP 100"}
  Top *string `json:"top,omitempty" xml:"top,omitempty"`
  // {"en":"Sorting:
  // 1. Optional values: request, flow
  // 2. Do not pass the default request", "zh_CN":"排序：
  // 1、可选值为：request, flow
  // 2、不传默认request"}
  OrderBy *string `json:"orderBy,omitempty" xml:"orderBy,omitempty"`
}

func (s ReportVisitorCustomTopDailyServiceRequest) String() string {
  return tea.Prettify(s)
}

func (s ReportVisitorCustomTopDailyServiceRequest) GoString() string {
  return s.String()
}

func (s *ReportVisitorCustomTopDailyServiceRequest) SetDateFrom(v string) *ReportVisitorCustomTopDailyServiceRequest {
  s.DateFrom = &v
  return s
}

func (s *ReportVisitorCustomTopDailyServiceRequest) SetDateTo(v string) *ReportVisitorCustomTopDailyServiceRequest {
  s.DateTo = &v
  return s
}

func (s *ReportVisitorCustomTopDailyServiceRequest) SetDomain(v []*string) *ReportVisitorCustomTopDailyServiceRequest {
  s.Domain = v
  return s
}

func (s *ReportVisitorCustomTopDailyServiceRequest) SetTop(v string) *ReportVisitorCustomTopDailyServiceRequest {
  s.Top = &v
  return s
}

func (s *ReportVisitorCustomTopDailyServiceRequest) SetOrderBy(v string) *ReportVisitorCustomTopDailyServiceRequest {
  s.OrderBy = &v
  return s
}

type ReportVisitorCustomTopDailyServiceResponse struct {
  // {"en":"request result status code", "zh_CN":"请求结果状态码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Request result information", "zh_CN":"请求结果信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Detailed data on the results of the request", "zh_CN":"请求结果的详细数据"}
  ReportVisitorCustomTopDailyServiceData []*ReportVisitorCustomTopDailyServiceData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s ReportVisitorCustomTopDailyServiceResponse) String() string {
  return tea.Prettify(s)
}

func (s ReportVisitorCustomTopDailyServiceResponse) GoString() string {
  return s.String()
}

func (s *ReportVisitorCustomTopDailyServiceResponse) SetCode(v string) *ReportVisitorCustomTopDailyServiceResponse {
  s.Code = &v
  return s
}

func (s *ReportVisitorCustomTopDailyServiceResponse) SetMessage(v string) *ReportVisitorCustomTopDailyServiceResponse {
  s.Message = &v
  return s
}

func (s *ReportVisitorCustomTopDailyServiceResponse) SetData(v []*ReportVisitorCustomTopDailyServiceData) *ReportVisitorCustomTopDailyServiceResponse {
  s.ReportVisitorCustomTopDailyServiceData = v
  return s
}

type ReportVisitorCustomTopDailyServiceData struct {
  // {"en":"Top", "zh_CN":"top排名"}
  Top *string `json:"top,omitempty" xml:"top,omitempty" require:"true"`
  // {"en":"IP", "zh_CN":"ip"}
  Ip *string `json:"ip,omitempty" xml:"ip,omitempty" require:"true"`
  // {"en":"Total traffic: The unit of measurement is MB, with 2 decimal places.", "zh_CN":"总流量：计量单位MB，保留2位小数"}
  TotalFlow *string `json:"totalFlow,omitempty" xml:"totalFlow,omitempty" require:"true"`
  // {"en":"Total requests", "zh_CN":"总请求数"}
  TotalRequest *string `json:"totalRequest,omitempty" xml:"totalRequest,omitempty" require:"true"`
}

func (s ReportVisitorCustomTopDailyServiceData) String() string {
  return tea.Prettify(s)
}

func (s ReportVisitorCustomTopDailyServiceData) GoString() string {
  return s.String()
}

func (s *ReportVisitorCustomTopDailyServiceData) SetTop(v string) *ReportVisitorCustomTopDailyServiceData {
  s.Top = &v
  return s
}

func (s *ReportVisitorCustomTopDailyServiceData) SetIp(v string) *ReportVisitorCustomTopDailyServiceData {
  s.Ip = &v
  return s
}

func (s *ReportVisitorCustomTopDailyServiceData) SetTotalFlow(v string) *ReportVisitorCustomTopDailyServiceData {
  s.TotalFlow = &v
  return s
}

func (s *ReportVisitorCustomTopDailyServiceData) SetTotalRequest(v string) *ReportVisitorCustomTopDailyServiceData {
  s.TotalRequest = &v
  return s
}

type ReportVisitorCustomTopDailyServicePaths struct {
}

func (s ReportVisitorCustomTopDailyServicePaths) String() string {
  return tea.Prettify(s)
}

func (s ReportVisitorCustomTopDailyServicePaths) GoString() string {
  return s.String()
}

type ReportVisitorCustomTopDailyServiceParameters struct {
}

func (s ReportVisitorCustomTopDailyServiceParameters) String() string {
  return tea.Prettify(s)
}

func (s ReportVisitorCustomTopDailyServiceParameters) GoString() string {
  return s.String()
}

type ReportVisitorCustomTopDailyServiceRequestHeader struct {
}

func (s ReportVisitorCustomTopDailyServiceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportVisitorCustomTopDailyServiceRequestHeader) GoString() string {
  return s.String()
}

type ReportVisitorCustomTopDailyServiceResponseHeader struct {
}

func (s ReportVisitorCustomTopDailyServiceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportVisitorCustomTopDailyServiceResponseHeader) GoString() string {
  return s.String()
}




type StatisticalAnalysisOfCountriesRequest struct {
  // {"en":"Start date:
  // 1. The time format is yyyy-MM-ddTHH:MM:ss+08:00. For example, 2019-01-01T10:00:00+08:00
  // 2. Cannot exceed current time
  // 3. The most recent six-month (183 days) data are available.", "zh_CN":"开始时间:
  // 1.时间格式为yyyy-MM-ddTHH:mm:ss+08:00,例如,2021-05-19T10:00:00+08:00(为北京时间2021年5月19日10点0分0秒)
  // 2.不能大于当前时间
  // 3.最多可获取最近半年(183天)的数据"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty"`
  // {"en":"End time:
  // 1. The time format is yyyy-MM-ddTHH:MM:ss+08:00. For example, 2019-01-01T10:00:00+08:00
  // 2. The end time needs to be greater than the start time. If the end time is greater than the current time, take the current time.
  // 3. Date from, Date to both, the default query past 24 hours; If there is only one unsent, throw an exception
  // 4. Maximum allowed query time interval: 31days (with technical support adjustments), meaning that the difference between Date from and dateTo cannot exceed 31days.", "zh_CN":"结束时间:
  // 1.时间格式yyyy-MM-ddTHH:mm:ss+08:00
  // 2.结束时间需大于开始时间,结束时间如果大于当前时间,取当前时间
  // 3.dateFrom,dateTo二者都未传,默认查询过去的1天;如仅有一个未传,抛异常
  // 4.允许查询最大时间间隔:31天,即dateFrom和dateTo相差不能超过31天。(可联系技术支持调整)"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty"`
  // {"en":"Domain name:
  // 1. The maximum number of TLDs that can be delivered is 20 by default (contact technical support adjustment);
  // 2. Automatically filter out illegal domain names (pass illegal domain names, will be filtered out, the query results only return the data of legitimate domain names)", "zh_CN":"域名:
  // 1. 可传递域名数量上限默认为20个(可联系技术支持调整);
  // 2. 自动过滤掉无效域名(如传递非法域名,会被过滤掉,查询结果只返回有效域名的数据)。"}
  Domain []*string `json:"domain,omitempty" xml:"domain,omitempty" type:"Repeated"`
  // {"en":"Order:
  // 1. The optional values are: request, successRequest, failRequest, totalTraffic, pv, uv, ip ;If left blank, default value is request", "zh_CN":"排序:
  // 1.可选值为:request, successRequest, failRequest, totalTraffic, pv, uv, ip
  // 不填默认为 request"}
  OrderBy *string `json:"orderBy,omitempty" xml:"orderBy,omitempty"`
  // {"en":"Area code:
  // 1.Default query for all countries and regions;
  // 2.The transferable values can be found in the description of the country/region list on the overview page.", "zh_CN":"地区代号:
  // 不传默认查询全部国家地区;
  // 可传递的值详见概览页国家地区列表说明。"}
  AreaCode []*string `json:"areaCode,omitempty" xml:"areaCode,omitempty" type:"Repeated"`
}

func (s StatisticalAnalysisOfCountriesRequest) String() string {
  return tea.Prettify(s)
}

func (s StatisticalAnalysisOfCountriesRequest) GoString() string {
  return s.String()
}

func (s *StatisticalAnalysisOfCountriesRequest) SetDateFrom(v string) *StatisticalAnalysisOfCountriesRequest {
  s.DateFrom = &v
  return s
}

func (s *StatisticalAnalysisOfCountriesRequest) SetDateTo(v string) *StatisticalAnalysisOfCountriesRequest {
  s.DateTo = &v
  return s
}

func (s *StatisticalAnalysisOfCountriesRequest) SetDomain(v []*string) *StatisticalAnalysisOfCountriesRequest {
  s.Domain = v
  return s
}

func (s *StatisticalAnalysisOfCountriesRequest) SetOrderBy(v string) *StatisticalAnalysisOfCountriesRequest {
  s.OrderBy = &v
  return s
}

func (s *StatisticalAnalysisOfCountriesRequest) SetAreaCode(v []*string) *StatisticalAnalysisOfCountriesRequest {
  s.AreaCode = v
  return s
}

type StatisticalAnalysisOfCountriesResponse struct {
  // {"en":"Request Result Status Code", "zh_CN":"请求结果状态码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Request Result Information", "zh_CN":"请求结果信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Detailed data of request results", "zh_CN":"请求结果的详细数据"}
  StatisticalAnalysisOfCountriesData []*StatisticalAnalysisOfCountriesData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s StatisticalAnalysisOfCountriesResponse) String() string {
  return tea.Prettify(s)
}

func (s StatisticalAnalysisOfCountriesResponse) GoString() string {
  return s.String()
}

func (s *StatisticalAnalysisOfCountriesResponse) SetCode(v string) *StatisticalAnalysisOfCountriesResponse {
  s.Code = &v
  return s
}

func (s *StatisticalAnalysisOfCountriesResponse) SetMessage(v string) *StatisticalAnalysisOfCountriesResponse {
  s.Message = &v
  return s
}

func (s *StatisticalAnalysisOfCountriesResponse) SetData(v []*StatisticalAnalysisOfCountriesData) *StatisticalAnalysisOfCountriesResponse {
  s.StatisticalAnalysisOfCountriesData = v
  return s
}

type StatisticalAnalysisOfCountriesData struct {
  // {"en":"Area code", "zh_CN":"地区代号"}
  AreaCode *string `json:"areaCode,omitempty" xml:"areaCode,omitempty" require:"true"`
  // {"en":"Area nam", "zh_CN":"地区名称"}
  AreaName *string `json:"areaName,omitempty" xml:"areaName,omitempty" require:"true"`
  // {"en":"Requests of area", "zh_CN":"地区对应的请求数"}
  Request *string `json:"request,omitempty" xml:"request,omitempty" require:"true"`
  // {"en":"Sucss requests of area", "zh_CN":"地区对应的成功请求数"}
  SuccessRequest *string `json:"successRequest,omitempty" xml:"successRequest,omitempty" require:"true"`
  // {"en":"Fail requests of area", "zh_CN":"地区对应的失败请求数"}
  FailRequest *string `json:"failRequest,omitempty" xml:"failRequest,omitempty" require:"true"`
  // {"en":"Successful request rate of the area", "zh_CN":"地区对应的成功请求率"}
  SuccessRate *string `json:"successRate,omitempty" xml:"successRate,omitempty" require:"true"`
  // {"en":"Total traffic of the area", "zh_CN":"地区对应的总流量,单位MB"}
  TotalTraffic *string `json:"totalTraffic,omitempty" xml:"totalTraffic,omitempty" require:"true"`
  // {"en":"UV of the area", "zh_CN":"地区对应的UV"}
  UniqueVisitors *string `json:"uniqueVisitors,omitempty" xml:"uniqueVisitors,omitempty" require:"true"`
  // {"en":"PV of the area", "zh_CN":"地区对应的PV"}
  PageViews *string `json:"pageViews,omitempty" xml:"pageViews,omitempty" require:"true"`
  // {"en":"IP amount of the area", "zh_CN":"地区对应的IP数"}
  IpAmount *string `json:"ipAmount,omitempty" xml:"ipAmount,omitempty" require:"true"`
}

func (s StatisticalAnalysisOfCountriesData) String() string {
  return tea.Prettify(s)
}

func (s StatisticalAnalysisOfCountriesData) GoString() string {
  return s.String()
}

func (s *StatisticalAnalysisOfCountriesData) SetAreaCode(v string) *StatisticalAnalysisOfCountriesData {
  s.AreaCode = &v
  return s
}

func (s *StatisticalAnalysisOfCountriesData) SetAreaName(v string) *StatisticalAnalysisOfCountriesData {
  s.AreaName = &v
  return s
}

func (s *StatisticalAnalysisOfCountriesData) SetRequest(v string) *StatisticalAnalysisOfCountriesData {
  s.Request = &v
  return s
}

func (s *StatisticalAnalysisOfCountriesData) SetSuccessRequest(v string) *StatisticalAnalysisOfCountriesData {
  s.SuccessRequest = &v
  return s
}

func (s *StatisticalAnalysisOfCountriesData) SetFailRequest(v string) *StatisticalAnalysisOfCountriesData {
  s.FailRequest = &v
  return s
}

func (s *StatisticalAnalysisOfCountriesData) SetSuccessRate(v string) *StatisticalAnalysisOfCountriesData {
  s.SuccessRate = &v
  return s
}

func (s *StatisticalAnalysisOfCountriesData) SetTotalTraffic(v string) *StatisticalAnalysisOfCountriesData {
  s.TotalTraffic = &v
  return s
}

func (s *StatisticalAnalysisOfCountriesData) SetUniqueVisitors(v string) *StatisticalAnalysisOfCountriesData {
  s.UniqueVisitors = &v
  return s
}

func (s *StatisticalAnalysisOfCountriesData) SetPageViews(v string) *StatisticalAnalysisOfCountriesData {
  s.PageViews = &v
  return s
}

func (s *StatisticalAnalysisOfCountriesData) SetIpAmount(v string) *StatisticalAnalysisOfCountriesData {
  s.IpAmount = &v
  return s
}

type StatisticalAnalysisOfCountriesPaths struct {
}

func (s StatisticalAnalysisOfCountriesPaths) String() string {
  return tea.Prettify(s)
}

func (s StatisticalAnalysisOfCountriesPaths) GoString() string {
  return s.String()
}

type StatisticalAnalysisOfCountriesParameters struct {
}

func (s StatisticalAnalysisOfCountriesParameters) String() string {
  return tea.Prettify(s)
}

func (s StatisticalAnalysisOfCountriesParameters) GoString() string {
  return s.String()
}

type StatisticalAnalysisOfCountriesRequestHeader struct {
}

func (s StatisticalAnalysisOfCountriesRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s StatisticalAnalysisOfCountriesRequestHeader) GoString() string {
  return s.String()
}

type StatisticalAnalysisOfCountriesResponseHeader struct {
}

func (s StatisticalAnalysisOfCountriesResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s StatisticalAnalysisOfCountriesResponseHeader) GoString() string {
  return s.String()
}




type StatisticalAnalysisOfChinaMainlandProvincevisitorsRequest struct {
  // {"en":"Start date:
  // 1. The time format is yyyy-MM-ddTHH:MM:ss+08:00. For example, 2019-01-01T10:00:00+08:00
  // 2. Cannot exceed current time
  // 3. The most recent six-month (183 days) data are available.", "zh_CN":"开始时间:
  // 1.时间格式为yyyy-MM-ddTHH:mm:ss+08:00,例如,2021-05-19T10:00:00+08:00(为北京时间2021年5月19日10点0分0秒)
  // 2.不能大于当前时间
  // 3.最多可获取最近半年(183天)的数据"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty"`
  // {"en":"End time:
  // 1. The time format is yyyy-MM-ddTHH:MM:ss+08:00. For example, 2019-01-01T10:00:00+08:00
  // 2. The end time needs to be greater than the start time. If the end time is greater than the current time, take the current time.
  // 3. Date from, Date to both, the default query past 24 hours; If there is only one unsent, throw an exception
  // 4. Maximum allowed query time interval: 31days (with technical support adjustments), meaning that the difference between Date from and dateTo cannot exceed 31days.", "zh_CN":"结束时间:
  // 1.时间格式yyyy-MM-ddTHH:mm:ss+08:00
  // 2.结束时间需大于开始时间,结束时间如果大于当前时间,取当前时间
  // 3.dateFrom,dateTo二者都未传,默认查询过去的1天;如仅有一个未传,抛异常
  // 4.允许查询最大时间间隔:31天,即dateFrom和dateTo相差不能超过31天。(可联系技术支持调整)"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty"`
  // {"en":"Domain name:
  // 1. The maximum number of TLDs that can be delivered is 20 by default (contact technical support adjustment);
  // 2. Automatically filter out illegal domain names (pass illegal domain names, will be filtered out, the query results only return the data of legitimate domain names)", "zh_CN":"域名:
  // 1.可传递域名数量上限默认为20个(可联系技术支持调整);
  // 2.自动过滤掉无效域名(如传递非法域名,会被过滤掉,查询结果只返回有效域名的数据)。"}
  Domain []*string `json:"domain,omitempty" xml:"domain,omitempty" type:"Repeated"`
  // {"en":"Order:
  // 1. The optional values are: request, successRequest, failRequest, totalTraffic, pv, uv, ip ;If left blank, default value is request", "zh_CN":"排序:
  // 1.可选值为:request, successRequest, failRequest, totalTraffic, pv, uv, ip; 不填默认为 request"}
  OrderBy *string `json:"orderBy,omitempty" xml:"orderBy,omitempty"`
  // {"en":"Area code:
  // 1.Default query for all countries and regions;
  // 2.The transferable values can be found in the description of the country/region list on the overview page.", "zh_CN":"地区代号:
  // 1.不传默认查询全部省份地区;
  // 2.可传递的值详见概览页省份地区列表说明。"}
  AreaCode []*string `json:"areaCode,omitempty" xml:"areaCode,omitempty" type:"Repeated"`
}

func (s StatisticalAnalysisOfChinaMainlandProvincevisitorsRequest) String() string {
  return tea.Prettify(s)
}

func (s StatisticalAnalysisOfChinaMainlandProvincevisitorsRequest) GoString() string {
  return s.String()
}

func (s *StatisticalAnalysisOfChinaMainlandProvincevisitorsRequest) SetDateFrom(v string) *StatisticalAnalysisOfChinaMainlandProvincevisitorsRequest {
  s.DateFrom = &v
  return s
}

func (s *StatisticalAnalysisOfChinaMainlandProvincevisitorsRequest) SetDateTo(v string) *StatisticalAnalysisOfChinaMainlandProvincevisitorsRequest {
  s.DateTo = &v
  return s
}

func (s *StatisticalAnalysisOfChinaMainlandProvincevisitorsRequest) SetDomain(v []*string) *StatisticalAnalysisOfChinaMainlandProvincevisitorsRequest {
  s.Domain = v
  return s
}

func (s *StatisticalAnalysisOfChinaMainlandProvincevisitorsRequest) SetOrderBy(v string) *StatisticalAnalysisOfChinaMainlandProvincevisitorsRequest {
  s.OrderBy = &v
  return s
}

func (s *StatisticalAnalysisOfChinaMainlandProvincevisitorsRequest) SetAreaCode(v []*string) *StatisticalAnalysisOfChinaMainlandProvincevisitorsRequest {
  s.AreaCode = v
  return s
}

type StatisticalAnalysisOfChinaMainlandProvincevisitorsResponse struct {
  // {"en":"Request Result Status Code", "zh_CN":"请求结果状态码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Request Result Information", "zh_CN":"请求结果信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Detailed data of request results", "zh_CN":"请求结果的详细数据"}
  StatisticalAnalysisOfChinaMainlandProvincevisitorsData []*StatisticalAnalysisOfChinaMainlandProvincevisitorsData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s StatisticalAnalysisOfChinaMainlandProvincevisitorsResponse) String() string {
  return tea.Prettify(s)
}

func (s StatisticalAnalysisOfChinaMainlandProvincevisitorsResponse) GoString() string {
  return s.String()
}

func (s *StatisticalAnalysisOfChinaMainlandProvincevisitorsResponse) SetCode(v string) *StatisticalAnalysisOfChinaMainlandProvincevisitorsResponse {
  s.Code = &v
  return s
}

func (s *StatisticalAnalysisOfChinaMainlandProvincevisitorsResponse) SetMessage(v string) *StatisticalAnalysisOfChinaMainlandProvincevisitorsResponse {
  s.Message = &v
  return s
}

func (s *StatisticalAnalysisOfChinaMainlandProvincevisitorsResponse) SetData(v []*StatisticalAnalysisOfChinaMainlandProvincevisitorsData) *StatisticalAnalysisOfChinaMainlandProvincevisitorsResponse {
  s.StatisticalAnalysisOfChinaMainlandProvincevisitorsData = v
  return s
}

type StatisticalAnalysisOfChinaMainlandProvincevisitorsData struct {
  // {"en":"Area code", "zh_CN":"地区代号"}
  AreaCode *string `json:"areaCode,omitempty" xml:"areaCode,omitempty" require:"true"`
  // {"en":"Area name", "zh_CN":"地区名称"}
  AreaName *string `json:"areaName,omitempty" xml:"areaName,omitempty" require:"true"`
  // {"en":"Requests of area", "zh_CN":"地区对应的请求数"}
  Request *string `json:"request,omitempty" xml:"request,omitempty" require:"true"`
  // {"en":"Sucss requests of area", "zh_CN":"地区对应的成功请求数"}
  SuccessRequest *string `json:"successRequest,omitempty" xml:"successRequest,omitempty" require:"true"`
  // {"en":"Fail requests of area", "zh_CN":"地区对应的失败请求数"}
  FailRequest *string `json:"failRequest,omitempty" xml:"failRequest,omitempty" require:"true"`
  // {"en":"Successful request rate of the area", "zh_CN":"地区对应的成功请求率"}
  SuccessRate *string `json:"successRate,omitempty" xml:"successRate,omitempty" require:"true"`
  // {"en":"Total traffic of the area", "zh_CN":"地区对应的总流量,单位MB"}
  TotalTraffic *string `json:"totalTraffic,omitempty" xml:"totalTraffic,omitempty" require:"true"`
  // {"en":"UV of the area", "zh_CN":"地区对应的UV"}
  UniqueVisitors *string `json:"uniqueVisitors,omitempty" xml:"uniqueVisitors,omitempty" require:"true"`
  // {"en":"PV of the area", "zh_CN":"地区对应的PV"}
  PageViews *string `json:"pageViews,omitempty" xml:"pageViews,omitempty" require:"true"`
  // {"en":"IP amount of the area", "zh_CN":"地区对应的IP数"}
  IpAmount *string `json:"ipAmount,omitempty" xml:"ipAmount,omitempty" require:"true"`
}

func (s StatisticalAnalysisOfChinaMainlandProvincevisitorsData) String() string {
  return tea.Prettify(s)
}

func (s StatisticalAnalysisOfChinaMainlandProvincevisitorsData) GoString() string {
  return s.String()
}

func (s *StatisticalAnalysisOfChinaMainlandProvincevisitorsData) SetAreaCode(v string) *StatisticalAnalysisOfChinaMainlandProvincevisitorsData {
  s.AreaCode = &v
  return s
}

func (s *StatisticalAnalysisOfChinaMainlandProvincevisitorsData) SetAreaName(v string) *StatisticalAnalysisOfChinaMainlandProvincevisitorsData {
  s.AreaName = &v
  return s
}

func (s *StatisticalAnalysisOfChinaMainlandProvincevisitorsData) SetRequest(v string) *StatisticalAnalysisOfChinaMainlandProvincevisitorsData {
  s.Request = &v
  return s
}

func (s *StatisticalAnalysisOfChinaMainlandProvincevisitorsData) SetSuccessRequest(v string) *StatisticalAnalysisOfChinaMainlandProvincevisitorsData {
  s.SuccessRequest = &v
  return s
}

func (s *StatisticalAnalysisOfChinaMainlandProvincevisitorsData) SetFailRequest(v string) *StatisticalAnalysisOfChinaMainlandProvincevisitorsData {
  s.FailRequest = &v
  return s
}

func (s *StatisticalAnalysisOfChinaMainlandProvincevisitorsData) SetSuccessRate(v string) *StatisticalAnalysisOfChinaMainlandProvincevisitorsData {
  s.SuccessRate = &v
  return s
}

func (s *StatisticalAnalysisOfChinaMainlandProvincevisitorsData) SetTotalTraffic(v string) *StatisticalAnalysisOfChinaMainlandProvincevisitorsData {
  s.TotalTraffic = &v
  return s
}

func (s *StatisticalAnalysisOfChinaMainlandProvincevisitorsData) SetUniqueVisitors(v string) *StatisticalAnalysisOfChinaMainlandProvincevisitorsData {
  s.UniqueVisitors = &v
  return s
}

func (s *StatisticalAnalysisOfChinaMainlandProvincevisitorsData) SetPageViews(v string) *StatisticalAnalysisOfChinaMainlandProvincevisitorsData {
  s.PageViews = &v
  return s
}

func (s *StatisticalAnalysisOfChinaMainlandProvincevisitorsData) SetIpAmount(v string) *StatisticalAnalysisOfChinaMainlandProvincevisitorsData {
  s.IpAmount = &v
  return s
}

type StatisticalAnalysisOfChinaMainlandProvincevisitorsPaths struct {
}

func (s StatisticalAnalysisOfChinaMainlandProvincevisitorsPaths) String() string {
  return tea.Prettify(s)
}

func (s StatisticalAnalysisOfChinaMainlandProvincevisitorsPaths) GoString() string {
  return s.String()
}

type StatisticalAnalysisOfChinaMainlandProvincevisitorsParameters struct {
}

func (s StatisticalAnalysisOfChinaMainlandProvincevisitorsParameters) String() string {
  return tea.Prettify(s)
}

func (s StatisticalAnalysisOfChinaMainlandProvincevisitorsParameters) GoString() string {
  return s.String()
}

type StatisticalAnalysisOfChinaMainlandProvincevisitorsRequestHeader struct {
}

func (s StatisticalAnalysisOfChinaMainlandProvincevisitorsRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s StatisticalAnalysisOfChinaMainlandProvincevisitorsRequestHeader) GoString() string {
  return s.String()
}

type StatisticalAnalysisOfChinaMainlandProvincevisitorsResponseHeader struct {
}

func (s StatisticalAnalysisOfChinaMainlandProvincevisitorsResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s StatisticalAnalysisOfChinaMainlandProvincevisitorsResponseHeader) GoString() string {
  return s.String()
}




type QueryTotalNumberofUniqueIPUnderSingleDomainRequest struct {
  // {"en":"Domain, maximum number is 1","zh_CN":"域名，数量上限1个"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {"en":"Stream name:\n1. Limit to the number of streams can be adjusted depending on different accounts. The default value is 20;\n2. All streams are queried by default is this field is left empty and at the same time, the upper limit of the number of streams can be set.","zh_CN":"流名：\n1.流名个数限制根据账号可调，默认为20个；\n2.不传时默认查询域名下所有流名，同时受流名数量上限限制；"}
  Stream *string `json:"stream,omitempty" xml:"stream,omitempty"`
}

func (s QueryTotalNumberofUniqueIPUnderSingleDomainRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryTotalNumberofUniqueIPUnderSingleDomainRequest) GoString() string {
  return s.String()
}

func (s *QueryTotalNumberofUniqueIPUnderSingleDomainRequest) SetDomain(v string) *QueryTotalNumberofUniqueIPUnderSingleDomainRequest {
  s.Domain = &v
  return s
}

func (s *QueryTotalNumberofUniqueIPUnderSingleDomainRequest) SetStream(v string) *QueryTotalNumberofUniqueIPUnderSingleDomainRequest {
  s.Stream = &v
  return s
}

type QueryTotalNumberofUniqueIPUnderSingleDomainRequestHeader struct {
}

func (s QueryTotalNumberofUniqueIPUnderSingleDomainRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryTotalNumberofUniqueIPUnderSingleDomainRequestHeader) GoString() string {
  return s.String()
}

type QueryTotalNumberofUniqueIPUnderSingleDomainPaths struct {
}

func (s QueryTotalNumberofUniqueIPUnderSingleDomainPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryTotalNumberofUniqueIPUnderSingleDomainPaths) GoString() string {
  return s.String()
}

type QueryTotalNumberofUniqueIPUnderSingleDomainParameters struct {
  // {"en":"Start time\n1. The format is yyyy-MM-ddTHH:mm:ss+08:00;\n2. Must be smaller than the current time and dateTo;\n3. Period between dataFrom and dateTo cannot be longer than 3 days(technical support can be contacted to adjust);\n4. You can only query data for the last 2 years.","zh_CN":"开始时间\n1.格式为yyyy-MM-ddTHH:mm:ss+08:00；\n2.必须小于当前时间和dateTo；\n3.dateFrom和dateTo相差不能超过3天（可联系技术支持调整）；\n4.只能查询最近2年内数据。"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty" require:"true"`
  // {"en":"End time\n1. The format is yyyy-MM-ddTHH:mm:ss+08:00;\n2. Must be greater than dateFrom;\n3. The query range must include 00:00:00 of a certain day to query the data of that day. For example, if the query range includes 2017-11-07 00:00:00, the data of 2017-11-07 can be queried.","zh_CN":"结束时间\n1.格式为yyyy-MM-ddTHH:mm:ss+08:00；\n2.必须大于dateFrom；\n3.查询范围必须包含某一天的00:00:00，才能查询到当天数据，例如查询范围包含 2017-11-07 00:00:00 可以查询到2017-11-07当天数据。"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty" require:"true"`
}

func (s QueryTotalNumberofUniqueIPUnderSingleDomainParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryTotalNumberofUniqueIPUnderSingleDomainParameters) GoString() string {
  return s.String()
}

func (s *QueryTotalNumberofUniqueIPUnderSingleDomainParameters) SetDateFrom(v string) *QueryTotalNumberofUniqueIPUnderSingleDomainParameters {
  s.DateFrom = &v
  return s
}

func (s *QueryTotalNumberofUniqueIPUnderSingleDomainParameters) SetDateTo(v string) *QueryTotalNumberofUniqueIPUnderSingleDomainParameters {
  s.DateTo = &v
  return s
}

type QueryTotalNumberofUniqueIPUnderSingleDomainResponse struct {
  // {"en":"","zh_CN":""}
  Result []*QueryTotalNumberofUniqueIPUnderSingleDomainResponseResult `json:"result,omitempty" xml:"result,omitempty" require:"true" type:"Repeated"`
}

func (s QueryTotalNumberofUniqueIPUnderSingleDomainResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryTotalNumberofUniqueIPUnderSingleDomainResponse) GoString() string {
  return s.String()
}

func (s *QueryTotalNumberofUniqueIPUnderSingleDomainResponse) SetResult(v []*QueryTotalNumberofUniqueIPUnderSingleDomainResponseResult) *QueryTotalNumberofUniqueIPUnderSingleDomainResponse {
  s.Result = v
  return s
}

type QueryTotalNumberofUniqueIPUnderSingleDomainResponseResult struct     {
  // {"en":"Domain","zh_CN":"域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {"en":"","zh_CN":""}
  Details []*QueryTotalNumberofUniqueIPUnderSingleDomainResponseResultDetails `json:"details,omitempty" xml:"details,omitempty" require:"true" type:"Repeated"`
}

func (s QueryTotalNumberofUniqueIPUnderSingleDomainResponseResult) String() string {
  return tea.Prettify(s)
}

func (s QueryTotalNumberofUniqueIPUnderSingleDomainResponseResult) GoString() string {
  return s.String()
}

func (s *QueryTotalNumberofUniqueIPUnderSingleDomainResponseResult) SetDomain(v string) *QueryTotalNumberofUniqueIPUnderSingleDomainResponseResult {
  s.Domain = &v
  return s
}

func (s *QueryTotalNumberofUniqueIPUnderSingleDomainResponseResult) SetDetails(v []*QueryTotalNumberofUniqueIPUnderSingleDomainResponseResultDetails) *QueryTotalNumberofUniqueIPUnderSingleDomainResponseResult {
  s.Details = v
  return s
}

type QueryTotalNumberofUniqueIPUnderSingleDomainResponseResultDetails struct     {
  // {"en":"Stream name","zh_CN":"流名"}
  Stream *string `json:"stream,omitempty" xml:"stream,omitempty" require:"true"`
  // {"en":"Time, format is yyyy-MM-dd","zh_CN":"时间，格式为yyyy-MM-dd"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {"en":"Number of unique IP addresses","zh_CN":"独立IP数"}
  Total *int `json:"total,omitempty" xml:"total,omitempty" require:"true"`
}

func (s QueryTotalNumberofUniqueIPUnderSingleDomainResponseResultDetails) String() string {
  return tea.Prettify(s)
}

func (s QueryTotalNumberofUniqueIPUnderSingleDomainResponseResultDetails) GoString() string {
  return s.String()
}

func (s *QueryTotalNumberofUniqueIPUnderSingleDomainResponseResultDetails) SetStream(v string) *QueryTotalNumberofUniqueIPUnderSingleDomainResponseResultDetails {
  s.Stream = &v
  return s
}

func (s *QueryTotalNumberofUniqueIPUnderSingleDomainResponseResultDetails) SetTimestamp(v string) *QueryTotalNumberofUniqueIPUnderSingleDomainResponseResultDetails {
  s.Timestamp = &v
  return s
}

func (s *QueryTotalNumberofUniqueIPUnderSingleDomainResponseResultDetails) SetTotal(v int) *QueryTotalNumberofUniqueIPUnderSingleDomainResponseResultDetails {
  s.Total = &v
  return s
}

type QueryTotalNumberofUniqueIPUnderSingleDomainResponseHeader struct {
}

func (s QueryTotalNumberofUniqueIPUnderSingleDomainResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryTotalNumberofUniqueIPUnderSingleDomainResponseHeader) GoString() string {
  return s.String()
}




