package reportstatuscode

// This file is auto-generated, don't edit it. Thanks.
import (
  "github.com/alibabacloud-go/tea/tea"
)

type QueryIPV6StatusOfeachISPandProvinceRequest struct {
  // {"en":"Start time: 
  //     1.The time format is yyyy-MM-ddTHH:mm:ss±HH:mm. Please note: ±HH:mm is the time zone offset, which can be adjusted according to your data needs, for example, +00:00 represents UTC time, +08:00 represents East 8th District, and -05:00 represents West 5th District. 2024-01-15T10:30:45+00:00 means UTC time January 15, 2024 10:30:45 AM;
  //     2. No bigger than the current time. 
  //     3. Data in the last 183 days at most can be queried.", "zh_CN":"开始时间：
  // 1.时间格式为 yyyy-MM-ddTHH:mm:ss±HH:mm。请注意：±HH:mm 为时区偏移量，可根据您的数据需要进行调整，例如 +00:00 代表 UTC 时间，+08:00 代表东八区，-05:00 代表西五区。2024-01-15T10:30:45+00:00，表示UTC 时间 2024 年 1 月 15 日上午 10 点 30 分 45 秒;
  // 2.不能大于当前时间
  // 3.最多可获取最近半年（183天）的数据。"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty"`
  // {"en":"End time: 
  //     1.The time format is yyyy-MM-ddTHH:mm:ss±HH:mm. Please note: ±HH:mm is the time zone offset, which can be adjusted according to your data needs, for example, +00:00 represents UTC time, +08:00 represents East 8th District, and -05:00 represents West 5th District. 2024-01-15T10:30:45+00:00 means UTC time January 15, 2024 10:30:45 AM;
  //     2. End time should be greater than start time. If the end time is greater than current time, current time will be used. 
  //     3. If both fields of dataFrom and dateTo are left empty, then data in the last 24 hours will be queried by default; if only one field is filled in and one is left empty, then exception will be occur. 
  //     4. Allowable maximum time range for query: 1 day, means the period between dateFrom to dateTo should not exceed 1 day (can be adjusted by contacting technical support).", "zh_CN":"结束时间：
  // 1.时间格式为 yyyy-MM-ddTHH:mm:ss±HH:mm。请注意：±HH:mm 为时区偏移量，可根据您的数据需要进行调整，例如 +00:00 代表 UTC 时间，+08:00 代表东八区，-05:00 代表西五区。2024-01-15T10:30:45+00:00，表示UTC 时间 2024 年 1 月 15 日上午 10 点 30 分 45 秒;
  // 2.结束时间需大于开始时间，结束时间如果大于当前时间，取当前时间。
  // 3.dateFrom，dateTo二者都未传，默认查询过去的24小时；如仅有一个未传，抛异常
  // 4.允许查询最大时间间隔：1天，即dateFrom和dateTo相差不能超过1天。（可联系技术支持调整）"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty"`
  // {"en":"Domain name: 
  // 	1. The maximum number of domains is 200 by default (Technical Support Adjustment can be contacted); 
  //     2. Automatic filtering invalid domain name (if pass illegal domain name, can be filtered, query result only returns the data of valid domain name).", "zh_CN":"域名：
  // 1.可传递域名数量上限默认为200个（可联系技术支持调整）；
  // 2.自动过滤掉无效域名（如传递非法域名，会被过滤掉，查询结果只返回有效域名的数据）。"}
  Domain []*string `json:"domain,omitempty" xml:"domain,omitempty" require:"true" type:"Repeated"`
  // {"en":"Data granularity:
  //     1. Support 5m (5 minute granularity), 1h (1 hour granularity); 
  // 	2. The default is 5m", "zh_CN":"数据粒度：
  // 1.支持5m（5分钟）、1h（1小时）
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
  // {"en":"IP type: 
  // 	1.The optional values are IPv6 and IPv4 
  // 	2.If let this parameter empty,it will query all IP type", "zh_CN":"IP类型：
  // 1.可选值为 IPV6、IPV4
  // 2.不传默认查询全部"}
  IPType *string `json:"IPType,omitempty" xml:"IPType,omitempty"`
  // {"en":"Grouped dimension: 
  //     1.Aggregation date by default. 
  //     2.the optional value is domain,province,isp,allow to send multi option 
  //     3.send the Grouped dimension represent the need to display details by their corresponding values.For example, when groupBy is isp, the ISP dimension needs to be displayed in detail. When an ISP is not passed, it represents an aggregate date and would not return the ISP node. Provinces and domains have the same logic. For example, by passing 'groupBy': ['domain', 'province'], the ISP node under ispData does not need to return. {domain:'www.aaaa.com','ispData': [{'isp','China Telecom','provinceData': [...]}]}", "zh_CN":"分组关键词：
  // 1.默认聚合展示；
  // 2.可选值为domain.province.isp，可传入多个值；
  // 3.传入关键词则代表需要按照关键词对应的值展示明细； 例如groupBy传入isp，则isp维度需要明细展示；当没有传递isp，则代表isp聚合展示，同时isp节点则不返回。其他province和domain相同逻辑。 例如：传递'groupBy':   ['domain','province']，则ispData下的isp节点无需返回。 { 'domain': 'www.aaaa.com', 'ispData': [ { 'isp':   '中国电信', 'provinceData': [....] }]}"}
  GroupBy []*string `json:"groupBy,omitempty" xml:"groupBy,omitempty" type:"Repeated"`
}

func (s QueryIPV6StatusOfeachISPandProvinceRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryIPV6StatusOfeachISPandProvinceRequest) GoString() string {
  return s.String()
}

func (s *QueryIPV6StatusOfeachISPandProvinceRequest) SetDateFrom(v string) *QueryIPV6StatusOfeachISPandProvinceRequest {
  s.DateFrom = &v
  return s
}

func (s *QueryIPV6StatusOfeachISPandProvinceRequest) SetDateTo(v string) *QueryIPV6StatusOfeachISPandProvinceRequest {
  s.DateTo = &v
  return s
}

func (s *QueryIPV6StatusOfeachISPandProvinceRequest) SetDomain(v []*string) *QueryIPV6StatusOfeachISPandProvinceRequest {
  s.Domain = v
  return s
}

func (s *QueryIPV6StatusOfeachISPandProvinceRequest) SetDataInterval(v string) *QueryIPV6StatusOfeachISPandProvinceRequest {
  s.DataInterval = &v
  return s
}

func (s *QueryIPV6StatusOfeachISPandProvinceRequest) SetProvince(v []*string) *QueryIPV6StatusOfeachISPandProvinceRequest {
  s.Province = v
  return s
}

func (s *QueryIPV6StatusOfeachISPandProvinceRequest) SetIsp(v []*string) *QueryIPV6StatusOfeachISPandProvinceRequest {
  s.Isp = v
  return s
}

func (s *QueryIPV6StatusOfeachISPandProvinceRequest) SetIPType(v string) *QueryIPV6StatusOfeachISPandProvinceRequest {
  s.IPType = &v
  return s
}

func (s *QueryIPV6StatusOfeachISPandProvinceRequest) SetGroupBy(v []*string) *QueryIPV6StatusOfeachISPandProvinceRequest {
  s.GroupBy = v
  return s
}

type QueryIPV6StatusOfeachISPandProvinceResponse struct {
  Result []*QueryIPV6StatusOfeachISPandProvinceResponseResult `json:"result,omitempty" xml:"result,omitempty" require:"true" type:"Repeated"`
}

func (s QueryIPV6StatusOfeachISPandProvinceResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryIPV6StatusOfeachISPandProvinceResponse) GoString() string {
  return s.String()
}

func (s *QueryIPV6StatusOfeachISPandProvinceResponse) SetResult(v []*QueryIPV6StatusOfeachISPandProvinceResponseResult) *QueryIPV6StatusOfeachISPandProvinceResponse {
  s.Result = v
  return s
}

type QueryIPV6StatusOfeachISPandProvinceResponseResult struct     {
  // {"en":"Domain", "zh_CN":"域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  IspData []*QueryIPV6StatusOfeachISPandProvinceResponseResultIspData `json:"ispData,omitempty" xml:"ispData,omitempty" require:"true" type:"Repeated"`
}

func (s QueryIPV6StatusOfeachISPandProvinceResponseResult) String() string {
  return tea.Prettify(s)
}

func (s QueryIPV6StatusOfeachISPandProvinceResponseResult) GoString() string {
  return s.String()
}

func (s *QueryIPV6StatusOfeachISPandProvinceResponseResult) SetDomain(v string) *QueryIPV6StatusOfeachISPandProvinceResponseResult {
  s.Domain = &v
  return s
}

func (s *QueryIPV6StatusOfeachISPandProvinceResponseResult) SetIspData(v []*QueryIPV6StatusOfeachISPandProvinceResponseResultIspData) *QueryIPV6StatusOfeachISPandProvinceResponseResult {
  s.IspData = v
  return s
}

type QueryIPV6StatusOfeachISPandProvinceResponseResultIspData struct     {
  // {"en":"ISP", "zh_CN":"运营商"}
  Isp *string `json:"isp,omitempty" xml:"isp,omitempty" require:"true"`
  ProvinceData []*QueryIPV6StatusOfeachISPandProvinceResponseResultIspDataProvinceData `json:"provinceData,omitempty" xml:"provinceData,omitempty" require:"true" type:"Repeated"`
}

func (s QueryIPV6StatusOfeachISPandProvinceResponseResultIspData) String() string {
  return tea.Prettify(s)
}

func (s QueryIPV6StatusOfeachISPandProvinceResponseResultIspData) GoString() string {
  return s.String()
}

func (s *QueryIPV6StatusOfeachISPandProvinceResponseResultIspData) SetIsp(v string) *QueryIPV6StatusOfeachISPandProvinceResponseResultIspData {
  s.Isp = &v
  return s
}

func (s *QueryIPV6StatusOfeachISPandProvinceResponseResultIspData) SetProvinceData(v []*QueryIPV6StatusOfeachISPandProvinceResponseResultIspDataProvinceData) *QueryIPV6StatusOfeachISPandProvinceResponseResultIspData {
  s.ProvinceData = v
  return s
}

type QueryIPV6StatusOfeachISPandProvinceResponseResultIspDataProvinceData struct     {
  // {"en":"Province", "zh_CN":"省份"}
  Province *string `json:"province,omitempty" xml:"province,omitempty" require:"true"`
  StatusCodeData []*QueryIPV6StatusOfeachISPandProvinceResponseResultIspDataProvinceDataStatusCodeData `json:"statusCodeData,omitempty" xml:"statusCodeData,omitempty" require:"true" type:"Repeated"`
}

func (s QueryIPV6StatusOfeachISPandProvinceResponseResultIspDataProvinceData) String() string {
  return tea.Prettify(s)
}

func (s QueryIPV6StatusOfeachISPandProvinceResponseResultIspDataProvinceData) GoString() string {
  return s.String()
}

func (s *QueryIPV6StatusOfeachISPandProvinceResponseResultIspDataProvinceData) SetProvince(v string) *QueryIPV6StatusOfeachISPandProvinceResponseResultIspDataProvinceData {
  s.Province = &v
  return s
}

func (s *QueryIPV6StatusOfeachISPandProvinceResponseResultIspDataProvinceData) SetStatusCodeData(v []*QueryIPV6StatusOfeachISPandProvinceResponseResultIspDataProvinceDataStatusCodeData) *QueryIPV6StatusOfeachISPandProvinceResponseResultIspDataProvinceData {
  s.StatusCodeData = v
  return s
}

type QueryIPV6StatusOfeachISPandProvinceResponseResultIspDataProvinceDataStatusCodeData struct     {
  // {"en":"Status Code", "zh_CN":"状态码类型"}
  StatusCode *string `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
  RequestData []*QueryIPV6StatusOfeachISPandProvinceResponseResultIspDataProvinceDataStatusCodeDataRequestData `json:"requestData,omitempty" xml:"requestData,omitempty" require:"true" type:"Repeated"`
}

func (s QueryIPV6StatusOfeachISPandProvinceResponseResultIspDataProvinceDataStatusCodeData) String() string {
  return tea.Prettify(s)
}

func (s QueryIPV6StatusOfeachISPandProvinceResponseResultIspDataProvinceDataStatusCodeData) GoString() string {
  return s.String()
}

func (s *QueryIPV6StatusOfeachISPandProvinceResponseResultIspDataProvinceDataStatusCodeData) SetStatusCode(v string) *QueryIPV6StatusOfeachISPandProvinceResponseResultIspDataProvinceDataStatusCodeData {
  s.StatusCode = &v
  return s
}

func (s *QueryIPV6StatusOfeachISPandProvinceResponseResultIspDataProvinceDataStatusCodeData) SetRequestData(v []*QueryIPV6StatusOfeachISPandProvinceResponseResultIspDataProvinceDataStatusCodeDataRequestData) *QueryIPV6StatusOfeachISPandProvinceResponseResultIspDataProvinceDataStatusCodeData {
  s.RequestData = v
  return s
}

type QueryIPV6StatusOfeachISPandProvinceResponseResultIspDataProvinceDataStatusCodeDataRequestData struct     {
  // {"en":"Time granularity is 5m, the format is yyyy-MM-dd HH:mm", "zh_CN":"数据粒度为5分钟，格式为yyyy-MM-dd HH:mm"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {"en":"Requests of status code", "zh_CN":"状态码请求数"}
  Value *string `json:"value,omitempty" xml:"value,omitempty" require:"true"`
}

func (s QueryIPV6StatusOfeachISPandProvinceResponseResultIspDataProvinceDataStatusCodeDataRequestData) String() string {
  return tea.Prettify(s)
}

func (s QueryIPV6StatusOfeachISPandProvinceResponseResultIspDataProvinceDataStatusCodeDataRequestData) GoString() string {
  return s.String()
}

func (s *QueryIPV6StatusOfeachISPandProvinceResponseResultIspDataProvinceDataStatusCodeDataRequestData) SetTimestamp(v string) *QueryIPV6StatusOfeachISPandProvinceResponseResultIspDataProvinceDataStatusCodeDataRequestData {
  s.Timestamp = &v
  return s
}

func (s *QueryIPV6StatusOfeachISPandProvinceResponseResultIspDataProvinceDataStatusCodeDataRequestData) SetValue(v string) *QueryIPV6StatusOfeachISPandProvinceResponseResultIspDataProvinceDataStatusCodeDataRequestData {
  s.Value = &v
  return s
}

type QueryIPV6StatusOfeachISPandProvincePaths struct {
}

func (s QueryIPV6StatusOfeachISPandProvincePaths) String() string {
  return tea.Prettify(s)
}

func (s QueryIPV6StatusOfeachISPandProvincePaths) GoString() string {
  return s.String()
}

type QueryIPV6StatusOfeachISPandProvinceParameters struct {
}

func (s QueryIPV6StatusOfeachISPandProvinceParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryIPV6StatusOfeachISPandProvinceParameters) GoString() string {
  return s.String()
}

type QueryIPV6StatusOfeachISPandProvinceRequestHeader struct {
}

func (s QueryIPV6StatusOfeachISPandProvinceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryIPV6StatusOfeachISPandProvinceRequestHeader) GoString() string {
  return s.String()
}

type QueryIPV6StatusOfeachISPandProvinceResponseHeader struct {
}

func (s QueryIPV6StatusOfeachISPandProvinceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryIPV6StatusOfeachISPandProvinceResponseHeader) GoString() string {
  return s.String()
}




type QueryStatusCodeDistributioninCountriesRequest struct {
  // {"en":"Start time
  // 
  // 1. The format is yyyyy-MM-ddTHH:mm:SS+08:00, for example, 2016-12-02T10:00+08:00 (10:00:00 Beijing time on December 2, 2016);
  // 2. can not exceed the current time;
  // 3. the latest half year (183 days) data can be obtained at most.", "zh_CN":"开始时间
  // 1.格式为yyyy-MM-ddTHH:mm:ss+08:00,例如,2016-12-02T10:00:00+08:00(为北京时间2016年12月2日10点0分0秒);
  // 2.不能大于当前时间;
  // 3.最多可获取最近半年(183天)的数据。"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty"`
  // {"en":"End time:
  // 1. The format is yyyy-MM-ddTHH:mm:ss+08:00;
  // 2. The end time is greater than the start time. If the end time is greater than the current time, the current time is taken.
  // 3. DateFrom and dateTo are not uploaded, defaulting to query the past 24 hours; if only one is not uploaded, throw an exception;
  // 4. Maximum query interval allowed: 7 days, that is, the difference between dateFrom and dateTo can not exceed 7 days (technical support can be contacted to adjust). ", "zh_CN":"结束时间
  // 1.格式为yyyy-MM-ddTHH:mm:ss+08:00;
  // 2.结束时间需大于开始时间,结束时间如果大于当前时间,取当前时间;
  // 3.dateFrom,dateTo二者都未传,默认查询过去的24小时;如仅有一个未传,抛异常;
  // 4.允许查询最大间隔:7天,即dateFrom和dateTo相差不能超过7天(可联系技术支持调整)。"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty"`
  // {"en":"Domain names, domain number limits can be adjusted depending on different accounts. The default value is 20", "zh_CN":"域名:可传递域名数量上限默认为20个(可联系技术支持调整),未传递报错"}
  Domain []*string `json:"domain,omitempty" xml:"domain,omitempty" type:"Repeated"`
  // {"en":"Data granularity, 1m: 1-minute  5m: 5-minute  granularity, 1h: 1-hour granularity, 1d: 1-day granularity", "zh_CN":"数据粒度,1m: 1分钟粒度, 5m:5分钟粒度,1h:1小时粒度,1d:1天粒度"}
  DataInterval *string `json:"dataInterval,omitempty" xml:"dataInterval,omitempty"`
  // {"en":"Country code 1) By default, all countries  are queried; 2) For specific country and region names, please refer to the appendix chapter on the overview page).","zh_CN":"国家代号: 1) 没传返回所有国家; 2)具体国家和地区名称请参考概览页面的附录章节."}
  CountryCode []*string `json:"countryCode,omitempty" xml:"countryCode,omitempty" type:"Repeated"`
  // {"en":"Group dimension
  // 
  // 1.        Options   are domain, province, isp, and more than one value can be entered;
  // 
  // 2.        The   data is displayed according to the specified dimension;", "zh_CN":"分组维度
  // 1.可选值为domain、country,可传入多个值;
  // 2.有传入则按照该维度展示明细数据;"}
  GroupBy []*string `json:"groupBy,omitempty" xml:"groupBy,omitempty" type:"Repeated"`
}

func (s QueryStatusCodeDistributioninCountriesRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryStatusCodeDistributioninCountriesRequest) GoString() string {
  return s.String()
}

func (s *QueryStatusCodeDistributioninCountriesRequest) SetDateFrom(v string) *QueryStatusCodeDistributioninCountriesRequest {
  s.DateFrom = &v
  return s
}

func (s *QueryStatusCodeDistributioninCountriesRequest) SetDateTo(v string) *QueryStatusCodeDistributioninCountriesRequest {
  s.DateTo = &v
  return s
}

func (s *QueryStatusCodeDistributioninCountriesRequest) SetDomain(v []*string) *QueryStatusCodeDistributioninCountriesRequest {
  s.Domain = v
  return s
}

func (s *QueryStatusCodeDistributioninCountriesRequest) SetDataInterval(v string) *QueryStatusCodeDistributioninCountriesRequest {
  s.DataInterval = &v
  return s
}

func (s *QueryStatusCodeDistributioninCountriesRequest) SetCountryCode(v []*string) *QueryStatusCodeDistributioninCountriesRequest {
  s.CountryCode = v
  return s
}

func (s *QueryStatusCodeDistributioninCountriesRequest) SetGroupBy(v []*string) *QueryStatusCodeDistributioninCountriesRequest {
  s.GroupBy = v
  return s
}

type QueryStatusCodeDistributioninCountriesResponse struct {
  // {"en":"request result status code", "zh_CN":"请求结果状态码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"request result information", "zh_CN":"请求结果信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  Data []*QueryStatusCodeDistributioninCountriesResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s QueryStatusCodeDistributioninCountriesResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryStatusCodeDistributioninCountriesResponse) GoString() string {
  return s.String()
}

func (s *QueryStatusCodeDistributioninCountriesResponse) SetCode(v string) *QueryStatusCodeDistributioninCountriesResponse {
  s.Code = &v
  return s
}

func (s *QueryStatusCodeDistributioninCountriesResponse) SetMessage(v string) *QueryStatusCodeDistributioninCountriesResponse {
  s.Message = &v
  return s
}

func (s *QueryStatusCodeDistributioninCountriesResponse) SetData(v []*QueryStatusCodeDistributioninCountriesResponseData) *QueryStatusCodeDistributioninCountriesResponse {
  s.Data = v
  return s
}

type QueryStatusCodeDistributioninCountriesResponseData struct     {
  // {"en":"Domain", "zh_CN":"域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  CountryData []*QueryStatusCodeDistributioninCountriesResponseDataCountryData `json:"countryData,omitempty" xml:"countryData,omitempty" require:"true" type:"Repeated"`
}

func (s QueryStatusCodeDistributioninCountriesResponseData) String() string {
  return tea.Prettify(s)
}

func (s QueryStatusCodeDistributioninCountriesResponseData) GoString() string {
  return s.String()
}

func (s *QueryStatusCodeDistributioninCountriesResponseData) SetDomain(v string) *QueryStatusCodeDistributioninCountriesResponseData {
  s.Domain = &v
  return s
}

func (s *QueryStatusCodeDistributioninCountriesResponseData) SetCountryData(v []*QueryStatusCodeDistributioninCountriesResponseDataCountryData) *QueryStatusCodeDistributioninCountriesResponseData {
  s.CountryData = v
  return s
}

type QueryStatusCodeDistributioninCountriesResponseDataCountryData struct     {
  // {"en":"country code", "zh_CN":"国家代码"}
  CountryCode *string `json:"countryCode,omitempty" xml:"countryCode,omitempty" require:"true"`
  // {"en":"country name", "zh_CN":"国家名称"}
  CountryName *string `json:"countryName,omitempty" xml:"countryName,omitempty" require:"true"`
  StatusCodeData []*QueryStatusCodeDistributioninCountriesResponseDataCountryDataStatusCodeData `json:"statusCodeData,omitempty" xml:"statusCodeData,omitempty" require:"true" type:"Repeated"`
}

func (s QueryStatusCodeDistributioninCountriesResponseDataCountryData) String() string {
  return tea.Prettify(s)
}

func (s QueryStatusCodeDistributioninCountriesResponseDataCountryData) GoString() string {
  return s.String()
}

func (s *QueryStatusCodeDistributioninCountriesResponseDataCountryData) SetCountryCode(v string) *QueryStatusCodeDistributioninCountriesResponseDataCountryData {
  s.CountryCode = &v
  return s
}

func (s *QueryStatusCodeDistributioninCountriesResponseDataCountryData) SetCountryName(v string) *QueryStatusCodeDistributioninCountriesResponseDataCountryData {
  s.CountryName = &v
  return s
}

func (s *QueryStatusCodeDistributioninCountriesResponseDataCountryData) SetStatusCodeData(v []*QueryStatusCodeDistributioninCountriesResponseDataCountryDataStatusCodeData) *QueryStatusCodeDistributioninCountriesResponseDataCountryData {
  s.StatusCodeData = v
  return s
}

type QueryStatusCodeDistributioninCountriesResponseDataCountryDataStatusCodeData struct     {
  // {"en":"status code", "zh_CN":"状态码"}
  StatusCode *string `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
  RequestData []*QueryStatusCodeDistributioninCountriesResponseDataCountryDataStatusCodeDataRequestData `json:"requestData,omitempty" xml:"requestData,omitempty" require:"true" type:"Repeated"`
}

func (s QueryStatusCodeDistributioninCountriesResponseDataCountryDataStatusCodeData) String() string {
  return tea.Prettify(s)
}

func (s QueryStatusCodeDistributioninCountriesResponseDataCountryDataStatusCodeData) GoString() string {
  return s.String()
}

func (s *QueryStatusCodeDistributioninCountriesResponseDataCountryDataStatusCodeData) SetStatusCode(v string) *QueryStatusCodeDistributioninCountriesResponseDataCountryDataStatusCodeData {
  s.StatusCode = &v
  return s
}

func (s *QueryStatusCodeDistributioninCountriesResponseDataCountryDataStatusCodeData) SetRequestData(v []*QueryStatusCodeDistributioninCountriesResponseDataCountryDataStatusCodeDataRequestData) *QueryStatusCodeDistributioninCountriesResponseDataCountryDataStatusCodeData {
  s.RequestData = v
  return s
}

type QueryStatusCodeDistributioninCountriesResponseDataCountryDataStatusCodeDataRequestData struct     {
  // {"en":"Time,
  // 1.        When   the data query granularity is 5m, the format is yyyy-MM-dd HH:mm; the data   value of every time slice represents the data value within the previous time   granularity range. The first time slice of the day is yyyy-MM-dd 12:05 AM,   and the last one is (yyyy-MM-dd+1) 00:00;
  // 2.        When   the data query granularity is 1h, the format is yyyy-MM-dd HH; the data value   of every time slice represents the data value within the previous time   granularity range. The first time slice of the day is yyyy-MM-dd 00:01, and   the last one is (yyyy-MM-dd+1) 00;
  // 3.        Return   the time slice contained in start time and the time slice contained in end   time.", "zh_CN":"时间,
  // 1.查询的数据粒度为5m时,格式为yyyy-MM-dd HH:mm;每一个时间片数据值代表的是前一个时间粒度范围内的数据值。一天开始的时间片是yyyy-MM-dd 00:05,最后一个时间片是(yyyy-MM-dd+1)00:00;
  // 2.查询的数据粒度为1h时,格式为yyyy-MM-dd HH;每一个时间片数据值代表的是前一个时间粒度范围内的数据值。一天开始的时间片是yyyy-MM-dd 01,最后一个时间片是(yyyy-MM-dd+1)&nbsp;00;
  // 3.返回开始时间和结束时间包含的时间片。"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {"en":"Number of requests of the status   code", "zh_CN":"状态码对应的请求数"}
  Value *string `json:"value,omitempty" xml:"value,omitempty" require:"true"`
}

func (s QueryStatusCodeDistributioninCountriesResponseDataCountryDataStatusCodeDataRequestData) String() string {
  return tea.Prettify(s)
}

func (s QueryStatusCodeDistributioninCountriesResponseDataCountryDataStatusCodeDataRequestData) GoString() string {
  return s.String()
}

func (s *QueryStatusCodeDistributioninCountriesResponseDataCountryDataStatusCodeDataRequestData) SetTimestamp(v string) *QueryStatusCodeDistributioninCountriesResponseDataCountryDataStatusCodeDataRequestData {
  s.Timestamp = &v
  return s
}

func (s *QueryStatusCodeDistributioninCountriesResponseDataCountryDataStatusCodeDataRequestData) SetValue(v string) *QueryStatusCodeDistributioninCountriesResponseDataCountryDataStatusCodeDataRequestData {
  s.Value = &v
  return s
}

type QueryStatusCodeDistributioninCountriesPaths struct {
}

func (s QueryStatusCodeDistributioninCountriesPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryStatusCodeDistributioninCountriesPaths) GoString() string {
  return s.String()
}

type QueryStatusCodeDistributioninCountriesParameters struct {
}

func (s QueryStatusCodeDistributioninCountriesParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryStatusCodeDistributioninCountriesParameters) GoString() string {
  return s.String()
}

type QueryStatusCodeDistributioninCountriesRequestHeader struct {
}

func (s QueryStatusCodeDistributioninCountriesRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryStatusCodeDistributioninCountriesRequestHeader) GoString() string {
  return s.String()
}

type QueryStatusCodeDistributioninCountriesResponseHeader struct {
}

func (s QueryStatusCodeDistributioninCountriesResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryStatusCodeDistributioninCountriesResponseHeader) GoString() string {
  return s.String()
}




type QueryOriginStatusCodeDistributionRequest struct {
  // {"en":"Start time\n1.The time format is yyyy-MM-ddTHH:mm:ss±HH:mm. Please note: ±HH:mm is the time zone offset, which can be adjusted according to your data needs, for example, +00:00 represents UTC time, +08:00 represents East 8th District, and -05:00 represents West 5th District. 2024-01-15T10:30:45+00:00 means UTC time January 15, 2024 10:30:45 AM;\n2.Must be a time that is 183 days earlier than the current time, and the time must be earlier than the current time and dateTo;\n3.Period between dataFrom and dateTo cannot be longer than 7 days;\n4dateFrom and dateTo can be either both are specified or neither is specifies;\n5.If neither dateFrom nor dateTo is specified, then by default, data in the last 24 hour is queried","zh_CN":"开始时间\n\n1.时间格式为 yyyy-MM-ddTHH:mm:ss±HH:mm。请注意：±HH:mm 为时区偏移量，可根据您的数据需要进行调整，例如 +00:00 代表 UTC 时间，+08:00 代表东八区，-05:00 代表西五区。2024-01-15T10:30:45+00:00，表示UTC 时间 2024 年 1 月 15 日上午 10 点 30 分 45 秒;\n\n2.必须大于当前时间-183天,并且小于当前时间和dateTo;\n\n3.dateFrom和dateTo默认时间跨度最大为7天(可联系技术支持调整，最大31天);\n\n4.dateFrom和dateTo要么都传递,要么都不传递;\n\n5.dateFrom和dateTo都未传递,则默认查询过去24小时的数据"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty"`
  // {"en":"End time\n\n1.The time format is yyyy-MM-ddTHH:mm:ss±HH:mm. Please note: ±HH:mm is the time zone offset, which can be adjusted according to your data needs, for example, +00:00 represents UTC time, +08:00 represents East 8th District, and -05:00 represents West 5th District. 2024-01-15T10:30:45+00:00 means UTC time January 15, 2024 10:30:45 AM;\n2.Must be greater than dateFrom; if it's greater than the current time, then the current time is assigned as the value;","zh_CN":"结束时间\n1.时间格式为 yyyy-MM-ddTHH:mm:ss±HH:mm。请注意：±HH:mm 为时区偏移量，可根据您的数据需要进行调整，例如 +00:00 代表 UTC 时间，+08:00 代表东八区，-05:00 代表西五区。2024-01-15T10:30:45+00:00，表示UTC 时间 2024 年 1 月 15 日上午 10 点 30 分 45 秒;\n2.必须大于dateFrom;如果大于当前时间,则重新赋值为当前时间;"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty"`
  // {"en":"1.The time format is yyyy-MM-ddTHH:mm:ss±HH:mm. Please note: ±HH:mm is the time zone offset, which can be adjusted according to your data needs, for example, +00:00 represents UTC time, +08:00 represents East 8th District, and -05:00 represents West 5th District. 2024-01-15T10:30:45+00:00 means UTC time January 15, 2024 10:30:45 AM;\n\n2.Must be a time that is 183 days earlier than the current time, and the time must be earlier than the current time and dateTo;\n\n3.The default time span of dateFrom and dateTo is up to 7 days (you can contact technical support to adjust it to a maximum of 31 days);\n\n4.dateFrom and dateTo can be either both are specified or neither is specifies;\n\n5.If neither dateFrom nor dateTo is specified, then by default, data in the last 24 hour is queried","zh_CN":"域名,域名个数限制根据账号可调,默认为20个"}
  Domain []*string `json:"domain,omitempty" xml:"domain,omitempty" require:"true" type:"Repeated"`
  // {"en":"Data granularity, 5m: granularity of 5 minutes","zh_CN":"数据粒度,5m:5分钟粒度"}
  DataInterval *string `json:"dataInterval,omitempty" xml:"dataInterval,omitempty"`
  // {"en":"Group dimension\n\n1.The value can be selected is domain;\n2.The data is displayed according to the specified dimension;","zh_CN":"分组维度\n1.可选值为domain;\n2.有传入则按照该维度展示明细数据;"}
  GroupBy []*string `json:"groupBy,omitempty" xml:"groupBy,omitempty" type:"Repeated"`
  // {"en":"Optional values 0, 1. Default is 0\nInput parameter 1 returns the total number of requests to the source, and 0 only returns the number of requests to the source station","zh_CN":"可选值 0, 1 。默认为 0\n入参 1 则返回全部回源请求数,入参 0 则只返回回源站请求数"}
  BacksrcOnly *QueryOriginStatusCodeDistributionRequestBacksrcOnly `json:"backsrcOnly,omitempty" xml:"backsrcOnly,omitempty" type:"Struct"`
  // {"en":"Query dimension. Optional values: statusCode , statusCodeType. Default value is statuscode.\n1.statusCode: returns the status code details;\n2.statusCodeType: returns the requests of each status code type (such as the number of requests corresponding to success, redirect, not modified, permission, not found, server error, and other)","zh_CN":"查询维度,可选值:statusCode, statusCodeType;不传默认statusCode\n1.statusCode :返回状态码明细;\n2.statusCodeType:返回状态码类型对应明细(如Success, Redirect, Not-Modified, Permission, Not-Found, Server Error, Other对应的请求数)"}
  QueryBy *string `json:"queryBy,omitempty" xml:"queryBy,omitempty"`
}

func (s QueryOriginStatusCodeDistributionRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryOriginStatusCodeDistributionRequest) GoString() string {
  return s.String()
}

func (s *QueryOriginStatusCodeDistributionRequest) SetDateFrom(v string) *QueryOriginStatusCodeDistributionRequest {
  s.DateFrom = &v
  return s
}

func (s *QueryOriginStatusCodeDistributionRequest) SetDateTo(v string) *QueryOriginStatusCodeDistributionRequest {
  s.DateTo = &v
  return s
}

func (s *QueryOriginStatusCodeDistributionRequest) SetDomain(v []*string) *QueryOriginStatusCodeDistributionRequest {
  s.Domain = v
  return s
}

func (s *QueryOriginStatusCodeDistributionRequest) SetDataInterval(v string) *QueryOriginStatusCodeDistributionRequest {
  s.DataInterval = &v
  return s
}

func (s *QueryOriginStatusCodeDistributionRequest) SetGroupBy(v []*string) *QueryOriginStatusCodeDistributionRequest {
  s.GroupBy = v
  return s
}

func (s *QueryOriginStatusCodeDistributionRequest) SetBacksrcOnly(v *QueryOriginStatusCodeDistributionRequestBacksrcOnly) *QueryOriginStatusCodeDistributionRequest {
  s.BacksrcOnly = v
  return s
}

func (s *QueryOriginStatusCodeDistributionRequest) SetQueryBy(v string) *QueryOriginStatusCodeDistributionRequest {
  s.QueryBy = &v
  return s
}

type QueryOriginStatusCodeDistributionRequestBacksrcOnly struct {
}

func (s QueryOriginStatusCodeDistributionRequestBacksrcOnly) String() string {
  return tea.Prettify(s)
}

func (s QueryOriginStatusCodeDistributionRequestBacksrcOnly) GoString() string {
  return s.String()
}

type QueryOriginStatusCodeDistributionRequestHeader struct {
}

func (s QueryOriginStatusCodeDistributionRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryOriginStatusCodeDistributionRequestHeader) GoString() string {
  return s.String()
}

type QueryOriginStatusCodeDistributionPaths struct {
}

func (s QueryOriginStatusCodeDistributionPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryOriginStatusCodeDistributionPaths) GoString() string {
  return s.String()
}

type QueryOriginStatusCodeDistributionParameters struct {
}

func (s QueryOriginStatusCodeDistributionParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryOriginStatusCodeDistributionParameters) GoString() string {
  return s.String()
}

type QueryOriginStatusCodeDistributionResponse struct {
  // {"en":"result","zh_CN":"结果"}
  Result []*QueryOriginStatusCodeDistributionResponseResult `json:"result,omitempty" xml:"result,omitempty" require:"true" type:"Repeated"`
}

func (s QueryOriginStatusCodeDistributionResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryOriginStatusCodeDistributionResponse) GoString() string {
  return s.String()
}

func (s *QueryOriginStatusCodeDistributionResponse) SetResult(v []*QueryOriginStatusCodeDistributionResponseResult) *QueryOriginStatusCodeDistributionResponse {
  s.Result = v
  return s
}

type QueryOriginStatusCodeDistributionResponseResult struct     {
  // {"en":"Domain","zh_CN":"域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {"en":"Success, Redirect, Not-Modified, Permission, Not-Found, Server Error, Other","zh_CN":"Success, Redirect, Not-Modified, Permission, Not-Found, Server Error, Other"}
  StatusCodeType *string `json:"statusCodeType,omitempty" xml:"statusCodeType,omitempty" require:"true"`
  // {"en":"statusCodeOriginData","zh_CN":"回源状态码数据"}
  StatusCodeOriginData []*QueryOriginStatusCodeDistributionResponseResultStatusCodeOriginData `json:"statusCodeOriginData,omitempty" xml:"statusCodeOriginData,omitempty" require:"true" type:"Repeated"`
}

func (s QueryOriginStatusCodeDistributionResponseResult) String() string {
  return tea.Prettify(s)
}

func (s QueryOriginStatusCodeDistributionResponseResult) GoString() string {
  return s.String()
}

func (s *QueryOriginStatusCodeDistributionResponseResult) SetDomain(v string) *QueryOriginStatusCodeDistributionResponseResult {
  s.Domain = &v
  return s
}

func (s *QueryOriginStatusCodeDistributionResponseResult) SetStatusCodeType(v string) *QueryOriginStatusCodeDistributionResponseResult {
  s.StatusCodeType = &v
  return s
}

func (s *QueryOriginStatusCodeDistributionResponseResult) SetStatusCodeOriginData(v []*QueryOriginStatusCodeDistributionResponseResultStatusCodeOriginData) *QueryOriginStatusCodeDistributionResponseResult {
  s.StatusCodeOriginData = v
  return s
}

type QueryOriginStatusCodeDistributionResponseResultStatusCodeOriginData struct     {
  // {"en":"Status code","zh_CN":"状态码"}
  StatusCode *string `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
  // {"en":"requestData","zh_CN":"数据"}
  RequestData []*QueryOriginStatusCodeDistributionResponseResultStatusCodeOriginDataRequestData `json:"requestData,omitempty" xml:"requestData,omitempty" require:"true" type:"Repeated"`
}

func (s QueryOriginStatusCodeDistributionResponseResultStatusCodeOriginData) String() string {
  return tea.Prettify(s)
}

func (s QueryOriginStatusCodeDistributionResponseResultStatusCodeOriginData) GoString() string {
  return s.String()
}

func (s *QueryOriginStatusCodeDistributionResponseResultStatusCodeOriginData) SetStatusCode(v string) *QueryOriginStatusCodeDistributionResponseResultStatusCodeOriginData {
  s.StatusCode = &v
  return s
}

func (s *QueryOriginStatusCodeDistributionResponseResultStatusCodeOriginData) SetRequestData(v []*QueryOriginStatusCodeDistributionResponseResultStatusCodeOriginDataRequestData) *QueryOriginStatusCodeDistributionResponseResultStatusCodeOriginData {
  s.RequestData = v
  return s
}

type QueryOriginStatusCodeDistributionResponseResultStatusCodeOriginDataRequestData struct     {
  // {"en":"DateTime, the format is yyyy-MM-dd HH:mm; the data value of every time slice represents the data value within the previous time granularity range.","zh_CN":"时间,格式为yyyy-MM-dd HH:mm;每一个时间片数据值代表的是前一个时间粒度范围内的数据值。每一个时间片数据值代表的是前一个时间粒度范围内的数据值,比如yyyy-MM-dd 00:05,代表00:00到00:05范围内的数据。"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {"en":"Number of requests of the status  code","zh_CN":"状态码对应的请求数"}
  Value *string `json:"value,omitempty" xml:"value,omitempty" require:"true"`
}

func (s QueryOriginStatusCodeDistributionResponseResultStatusCodeOriginDataRequestData) String() string {
  return tea.Prettify(s)
}

func (s QueryOriginStatusCodeDistributionResponseResultStatusCodeOriginDataRequestData) GoString() string {
  return s.String()
}

func (s *QueryOriginStatusCodeDistributionResponseResultStatusCodeOriginDataRequestData) SetTimestamp(v string) *QueryOriginStatusCodeDistributionResponseResultStatusCodeOriginDataRequestData {
  s.Timestamp = &v
  return s
}

func (s *QueryOriginStatusCodeDistributionResponseResultStatusCodeOriginDataRequestData) SetValue(v string) *QueryOriginStatusCodeDistributionResponseResultStatusCodeOriginDataRequestData {
  s.Value = &v
  return s
}

type QueryOriginStatusCodeDistributionResponseHeader struct {
}

func (s QueryOriginStatusCodeDistributionResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryOriginStatusCodeDistributionResponseHeader) GoString() string {
  return s.String()
}




type QueryISPProvinceStatusCodeRequest struct {
  // {"en":"Start date:
  // 1.The time format is yyyy-MM-ddTHH:MM:ss+08:00. For example, 2019-01-01T10:00:00+08:00
  // 2.Cannot exceed current time
  // 3.The most recent six-month (183 days) data are available.", "zh_CN":"开始时间：
  // 1.时间格式为yyyy-MM-ddTHH:mm:ss+08:00，例如，2019-01-01T10:00:00+08:00(为北京时间2019年01月01日10点0分0秒)；
  // 2.不能大于当前时间
  // 3.最多可获取最近半年(183天)的数据。"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty"`
  // {"en":"End time:
  // 1.The time format is yyyy-MM-ddTHH:MM:ss+08:00. For example, 2019-01-01T10:00:00+08:00
  // 2.The end time needs to be greater than the start time. If the end time is greater than the current time, take the current time.
  // 3.Date from, Date to both, the default query past 24 hours; If there is only one unsent, throw an exception
  // 4.Maximum allowed query time interval: 24 hours (with technical support adjustments), meaning that the difference between Date from and dateTo cannot exceed 24 hours.", "zh_CN":"结束时间：
  // 1.时间格式为yyyy-MM-ddTHH:mm:ss+08:00，例如，2019-01-01T10:00:00+08:00(为北京时间2019年01月01日10点0分0秒)
  // 2.结束时间需大于开始时间，结束时间如果大于当前时间，取当前时间。
  // 3.dateFrom，dateTo二者都未传，默认查询过去的24小时；如仅有一个未传，抛异常
  // 4.允许查询最大时间间隔：24小时(可联系技术支持调整)，即dateFrom和dateTo相差不能超过24小时。"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty"`
  // {"en":"Domain name:
  // 1.The maximum number of TLDs that can be delivered is 20 by default (contact technical support adjustment);
  // 2.Automatically filter out illegal domain names (pass illegal domain names, will be filtered out, the query results only return the data of legitimate domain names)
  // 3.If left blank, all domain names will be obtained. If the total number of domain names exceeds the upper limit, an error will be reported.", "zh_CN":"域名：
  // 1.可传递域名数量上限默认为20个(可联系技术支持调整)；
  // 2.自动过滤掉非法域名(如传递非法域名，会被过滤掉，查询结果只返回合法域名的数据)
  // 3.若未填写默认查询全部域名，全部域名超出域名上限报错"}
  Domain []*string `json:"domain,omitempty" xml:"domain,omitempty" type:"Repeated"`
  // {"en":"Data granularity:  
  // 1.The default is 1m;
  // 2.Support 1m (1 minute), 5m (5 minutes)", "zh_CN":"数据粒度：
  // 1.不传默认1m
  // 2.支持1m(1分钟)、5m(5分钟)"}
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
  // 1.Optional value StatusCode, 2XX, 3XX, 4XX, 5XX, No default statusCode
  // 2.StatusCode: returns the status code details;
  // 3.2XX:Returns the 2XX return status code summary and the 2 start return status code detail data;
  // 4.3XX:Returns the 2XX return status code summary and the 3 start return status code detail data;
  // 5.4XX:Returns the 2XX return status code summary and the 4 start return status code detail data;
  // 6.5XX:Returns the 2XX return status code summary and the 5 start return status code detail data;", "zh_CN":"查询维度:
  // 1.可选值 statusCode 2XX 3XX 4XX 5XX, 不传默认 statusCode
  // 2.statusCode ：返回请求状态码明细；
  // 3.2XX:返回 2XX 状态码汇总及各 2 开头状态码明细数据；
  // 4.3XX:返回 3XX 状态码汇总及各 2 开头状态码明细数据；
  // 5.4XX:返回 4XX 状态码汇总及各 4 开头状态码明细数据；
  // 6.5XX:返回 5XX 状态码汇总及各 5 开头状态码明细数据；"}
  QueryBy *string `json:"queryBy,omitempty" xml:"queryBy,omitempty"`
  // {"en":"Optional: domain, all, If it is empty, it defaults to returning by domain dimension;
  // If all is passed, merge and return according to the query domain name.", "zh_CN":"可选项：domain、all, 为空则默认为按domain维度返回;
  // 若传递all，则按查询域名合并返回"}
  GroupBy *string `json:"groupBy,omitempty" xml:"groupBy,omitempty"`
}

func (s QueryISPProvinceStatusCodeRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryISPProvinceStatusCodeRequest) GoString() string {
  return s.String()
}

func (s *QueryISPProvinceStatusCodeRequest) SetDateFrom(v string) *QueryISPProvinceStatusCodeRequest {
  s.DateFrom = &v
  return s
}

func (s *QueryISPProvinceStatusCodeRequest) SetDateTo(v string) *QueryISPProvinceStatusCodeRequest {
  s.DateTo = &v
  return s
}

func (s *QueryISPProvinceStatusCodeRequest) SetDomain(v []*string) *QueryISPProvinceStatusCodeRequest {
  s.Domain = v
  return s
}

func (s *QueryISPProvinceStatusCodeRequest) SetDataInterval(v string) *QueryISPProvinceStatusCodeRequest {
  s.DataInterval = &v
  return s
}

func (s *QueryISPProvinceStatusCodeRequest) SetProvince(v []*string) *QueryISPProvinceStatusCodeRequest {
  s.Province = v
  return s
}

func (s *QueryISPProvinceStatusCodeRequest) SetIsp(v []*string) *QueryISPProvinceStatusCodeRequest {
  s.Isp = v
  return s
}

func (s *QueryISPProvinceStatusCodeRequest) SetQueryBy(v string) *QueryISPProvinceStatusCodeRequest {
  s.QueryBy = &v
  return s
}

func (s *QueryISPProvinceStatusCodeRequest) SetGroupBy(v string) *QueryISPProvinceStatusCodeRequest {
  s.GroupBy = &v
  return s
}

type QueryISPProvinceStatusCodeResponse struct {
  // {"en":"request result status code", "zh_CN":"请求结果状态码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Request result information", "zh_CN":"请求结果信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Detailed data on the results of the request", "zh_CN":"请求结果的详细数据"}
  Data []*QueryISPProvinceStatusCodeResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s QueryISPProvinceStatusCodeResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryISPProvinceStatusCodeResponse) GoString() string {
  return s.String()
}

func (s *QueryISPProvinceStatusCodeResponse) SetCode(v string) *QueryISPProvinceStatusCodeResponse {
  s.Code = &v
  return s
}

func (s *QueryISPProvinceStatusCodeResponse) SetMessage(v string) *QueryISPProvinceStatusCodeResponse {
  s.Message = &v
  return s
}

func (s *QueryISPProvinceStatusCodeResponse) SetData(v []*QueryISPProvinceStatusCodeResponseData) *QueryISPProvinceStatusCodeResponse {
  s.Data = v
  return s
}

type QueryISPProvinceStatusCodeResponseData struct     {
  // {"en":"domain", "zh_CN":"域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  StatusCodeDataList []*QueryISPProvinceStatusCodeResponseDataStatusCodeDataList `json:"statusCodeDataList,omitempty" xml:"statusCodeDataList,omitempty" require:"true" type:"Repeated"`
}

func (s QueryISPProvinceStatusCodeResponseData) String() string {
  return tea.Prettify(s)
}

func (s QueryISPProvinceStatusCodeResponseData) GoString() string {
  return s.String()
}

func (s *QueryISPProvinceStatusCodeResponseData) SetDomain(v string) *QueryISPProvinceStatusCodeResponseData {
  s.Domain = &v
  return s
}

func (s *QueryISPProvinceStatusCodeResponseData) SetStatusCodeDataList(v []*QueryISPProvinceStatusCodeResponseDataStatusCodeDataList) *QueryISPProvinceStatusCodeResponseData {
  s.StatusCodeDataList = v
  return s
}

type QueryISPProvinceStatusCodeResponseDataStatusCodeDataList struct     {
  // {"en":"StatusCode", "zh_CN":"状态码"}
  StatusCode *string `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
  DetailList []*QueryISPProvinceStatusCodeResponseDataStatusCodeDataListDetailList `json:"detailList,omitempty" xml:"detailList,omitempty" require:"true" type:"Repeated"`
}

func (s QueryISPProvinceStatusCodeResponseDataStatusCodeDataList) String() string {
  return tea.Prettify(s)
}

func (s QueryISPProvinceStatusCodeResponseDataStatusCodeDataList) GoString() string {
  return s.String()
}

func (s *QueryISPProvinceStatusCodeResponseDataStatusCodeDataList) SetStatusCode(v string) *QueryISPProvinceStatusCodeResponseDataStatusCodeDataList {
  s.StatusCode = &v
  return s
}

func (s *QueryISPProvinceStatusCodeResponseDataStatusCodeDataList) SetDetailList(v []*QueryISPProvinceStatusCodeResponseDataStatusCodeDataListDetailList) *QueryISPProvinceStatusCodeResponseDataStatusCodeDataList {
  s.DetailList = v
  return s
}

type QueryISPProvinceStatusCodeResponseDataStatusCodeDataListDetailList struct     {
  // {"en":"time, in yyyy-MM-dd HH:MM", "zh_CN":"时间，格式为yyyy-MM-dd HH:mm"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {"en":"Total number of  requests", "zh_CN":"总请求数"}
  Value *string `json:"value,omitempty" xml:"value,omitempty" require:"true"`
}

func (s QueryISPProvinceStatusCodeResponseDataStatusCodeDataListDetailList) String() string {
  return tea.Prettify(s)
}

func (s QueryISPProvinceStatusCodeResponseDataStatusCodeDataListDetailList) GoString() string {
  return s.String()
}

func (s *QueryISPProvinceStatusCodeResponseDataStatusCodeDataListDetailList) SetTimestamp(v string) *QueryISPProvinceStatusCodeResponseDataStatusCodeDataListDetailList {
  s.Timestamp = &v
  return s
}

func (s *QueryISPProvinceStatusCodeResponseDataStatusCodeDataListDetailList) SetValue(v string) *QueryISPProvinceStatusCodeResponseDataStatusCodeDataListDetailList {
  s.Value = &v
  return s
}

type QueryISPProvinceStatusCodePaths struct {
}

func (s QueryISPProvinceStatusCodePaths) String() string {
  return tea.Prettify(s)
}

func (s QueryISPProvinceStatusCodePaths) GoString() string {
  return s.String()
}

type QueryISPProvinceStatusCodeParameters struct {
}

func (s QueryISPProvinceStatusCodeParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryISPProvinceStatusCodeParameters) GoString() string {
  return s.String()
}

type QueryISPProvinceStatusCodeRequestHeader struct {
}

func (s QueryISPProvinceStatusCodeRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryISPProvinceStatusCodeRequestHeader) GoString() string {
  return s.String()
}

type QueryISPProvinceStatusCodeResponseHeader struct {
}

func (s QueryISPProvinceStatusCodeResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryISPProvinceStatusCodeResponseHeader) GoString() string {
  return s.String()
}




type ReportStatusCodeRealTimeEdgeServiceRequest struct {
  // {'en':'Start time:
  // 1. The time format is yyyy-MM-ddTHH:mm:ss±HH:mm. Please note: ±HH:mm is the time zone offset, which can be adjusted according to your data needs, for example, +00:00 represents UTC time, +08:00 represents East 8th District, and -05:00 represents West 5th District. 2024-01-15T10:30:45+00:00 means UTC time January 15, 2024 10:30:45 AM;
  // 2. Not greater than the current time
  // 3. The most recent half-year (183 days) data can be obtained', 'zh_CN':'开始时间：
  // 1.时间格式为 yyyy-MM-ddTHH:mm:ss±HH:mm。请注意：±HH:mm 为时区偏移量，可根据您的数据需要进行调整，例如 +00:00 代表 UTC 时间，+08:00 代表东八区，-05:00 代表西五区。2024-01-15T10:30:45+00:00，表示UTC 时间 2024 年 1 月 15 日上午 10 点 30 分 45 秒；
  // 2.不能大于当前时间
  // 3.最多可获取最近半年（183天）的数据。'}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty"`
  // {'en':'End time:
  // 1. The time format is yyyy-MM-ddTHH:mm:ss±HH:mm. Please note: ±HH:mm is the time zone offset, which can be adjusted according to your data needs, for example, +00:00 represents UTC time, +08:00 represents East 8th District, and -05:00 represents West 5th District. 2024-01-15T10:30:45+00:00 means UTC time January 15, 2024 10:30:45 AM
  // 2. End time should be greater than start time. If the end time is greater than current time, current time will be used.
  // 3. If both fields of dataFrom and dateTo are left empty, then data in the last 1 hours will be queried by default; if one field is filled and one is left empty, then exception will occur.
  // 4. Maximum time range allowable for query: 1 hour, means the period between dateFrom to dateTo should not exceed 1 hour', 'zh_CN':'结束时间：
  // 1.时间格式为 yyyy-MM-ddTHH:mm:ss±HH:mm。请注意：±HH:mm 为时区偏移量，可根据您的数据需要进行调整，例如 +00:00 代表 UTC 时间，+08:00 代表东八区，-05:00 代表西五区。2024-01-15T10:30:45+00:00，表示UTC 时间 2024 年 1 月 15 日上午 10 点 30 分 45 秒
  // 2.结束时间需大于开始时间，结束时间如果大于当前时间，取当前时间。
  // 3.dateFrom，dateTo二者都未传，默认查询过去的1小时；如仅有一个未传，抛异常
  // 4.允许查询最大时间间隔：1小时（可联系技术支持调整），即dateFrom和dateTo相差不能超过1小时。'}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty"`
  // {'en':'Domain:
  // 1. The default upper limit of domains that can be entered is 20 (if you want to adjust, please, contact technical support);
  // 2. Automatically filter out illegal domain names (pass illegal domain names, will be filtered out, the query results only return the data of legitimate domain names)
  // 3. Domain name exceeding limit, misstatement', 'zh_CN':'域名：
  // 1.可传递域名数量上限默认为20个（可联系技术支持调整）；
  // 2.自动过滤掉非法域名（如传递非法域名，会被过滤掉，查询结果只返回合法域名的数据）
  // 3.域名超过上限，报错'}
  Domain []*string `json:"domain,omitempty" xml:"domain,omitempty" require:"true" type:"Repeated"`
  // {'en':'Data granularity:
  // 1. default 1m
  // 2. 1m (1 minute), 5m (5 minutes)', 'zh_CN':'数据粒度：不传默认1m
  // 1.支持1m（1分钟）、5m（5分钟）'}
  DataInterval *string `json:"dataInterval,omitempty" xml:"dataInterval,omitempty"`
}

func (s ReportStatusCodeRealTimeEdgeServiceRequest) String() string {
  return tea.Prettify(s)
}

func (s ReportStatusCodeRealTimeEdgeServiceRequest) GoString() string {
  return s.String()
}

func (s *ReportStatusCodeRealTimeEdgeServiceRequest) SetDateFrom(v string) *ReportStatusCodeRealTimeEdgeServiceRequest {
  s.DateFrom = &v
  return s
}

func (s *ReportStatusCodeRealTimeEdgeServiceRequest) SetDateTo(v string) *ReportStatusCodeRealTimeEdgeServiceRequest {
  s.DateTo = &v
  return s
}

func (s *ReportStatusCodeRealTimeEdgeServiceRequest) SetDomain(v []*string) *ReportStatusCodeRealTimeEdgeServiceRequest {
  s.Domain = v
  return s
}

func (s *ReportStatusCodeRealTimeEdgeServiceRequest) SetDataInterval(v string) *ReportStatusCodeRealTimeEdgeServiceRequest {
  s.DataInterval = &v
  return s
}

type ReportStatusCodeRealTimeEdgeServiceResponse struct {
  // {'en':'result', 'zh_CN':'结果'}
  Result []*ReportStatusCodeRealTimeEdgeServiceResponseResult `json:"result,omitempty" xml:"result,omitempty" require:"true" type:"Repeated"`
}

func (s ReportStatusCodeRealTimeEdgeServiceResponse) String() string {
  return tea.Prettify(s)
}

func (s ReportStatusCodeRealTimeEdgeServiceResponse) GoString() string {
  return s.String()
}

func (s *ReportStatusCodeRealTimeEdgeServiceResponse) SetResult(v []*ReportStatusCodeRealTimeEdgeServiceResponseResult) *ReportStatusCodeRealTimeEdgeServiceResponse {
  s.Result = v
  return s
}

type ReportStatusCodeRealTimeEdgeServiceResponseResult struct     {
  // {'en':'statusCodeData', 'zh_CN':'状态码数据'}
  StatusCodeData []*ReportStatusCodeRealTimeEdgeServiceResponseResultStatusCodeData `json:"statusCodeData,omitempty" xml:"statusCodeData,omitempty" require:"true" type:"Repeated"`
}

func (s ReportStatusCodeRealTimeEdgeServiceResponseResult) String() string {
  return tea.Prettify(s)
}

func (s ReportStatusCodeRealTimeEdgeServiceResponseResult) GoString() string {
  return s.String()
}

func (s *ReportStatusCodeRealTimeEdgeServiceResponseResult) SetStatusCodeData(v []*ReportStatusCodeRealTimeEdgeServiceResponseResultStatusCodeData) *ReportStatusCodeRealTimeEdgeServiceResponseResult {
  s.StatusCodeData = v
  return s
}

type ReportStatusCodeRealTimeEdgeServiceResponseResultStatusCodeData struct     {
  // {'en':'Status code', 'zh_CN':'状态码'}
  StatusCode *string `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
  // {'en':'requestData', 'zh_CN':'请求数数据'}
  RequestData []*ReportStatusCodeRealTimeEdgeServiceResponseResultStatusCodeDataRequestData `json:"requestData,omitempty" xml:"requestData,omitempty" require:"true" type:"Repeated"`
}

func (s ReportStatusCodeRealTimeEdgeServiceResponseResultStatusCodeData) String() string {
  return tea.Prettify(s)
}

func (s ReportStatusCodeRealTimeEdgeServiceResponseResultStatusCodeData) GoString() string {
  return s.String()
}

func (s *ReportStatusCodeRealTimeEdgeServiceResponseResultStatusCodeData) SetStatusCode(v string) *ReportStatusCodeRealTimeEdgeServiceResponseResultStatusCodeData {
  s.StatusCode = &v
  return s
}

func (s *ReportStatusCodeRealTimeEdgeServiceResponseResultStatusCodeData) SetRequestData(v []*ReportStatusCodeRealTimeEdgeServiceResponseResultStatusCodeDataRequestData) *ReportStatusCodeRealTimeEdgeServiceResponseResultStatusCodeData {
  s.RequestData = v
  return s
}

type ReportStatusCodeRealTimeEdgeServiceResponseResultStatusCodeDataRequestData struct     {
  // {'en':'DateTime, the format is   yyyy-MM-dd HH:mm; the data value of every time slice represents the data   value within the previous time granularity range.', 'zh_CN':'时间，格式为yyyy-MM-dd HH:mm；每一个时间片数据值代表的是前一个时间粒度范围内的数据值。比如yyyy-MM-dd 00:05，代表00:00到00:05范围内的数据。'}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {'en':'Number of requests for status codes', 'zh_CN':'状态码对应的请求数'}
  Value *string `json:"value,omitempty" xml:"value,omitempty" require:"true"`
}

func (s ReportStatusCodeRealTimeEdgeServiceResponseResultStatusCodeDataRequestData) String() string {
  return tea.Prettify(s)
}

func (s ReportStatusCodeRealTimeEdgeServiceResponseResultStatusCodeDataRequestData) GoString() string {
  return s.String()
}

func (s *ReportStatusCodeRealTimeEdgeServiceResponseResultStatusCodeDataRequestData) SetTimestamp(v string) *ReportStatusCodeRealTimeEdgeServiceResponseResultStatusCodeDataRequestData {
  s.Timestamp = &v
  return s
}

func (s *ReportStatusCodeRealTimeEdgeServiceResponseResultStatusCodeDataRequestData) SetValue(v string) *ReportStatusCodeRealTimeEdgeServiceResponseResultStatusCodeDataRequestData {
  s.Value = &v
  return s
}

type ReportStatusCodeRealTimeEdgeServicePaths struct {
}

func (s ReportStatusCodeRealTimeEdgeServicePaths) String() string {
  return tea.Prettify(s)
}

func (s ReportStatusCodeRealTimeEdgeServicePaths) GoString() string {
  return s.String()
}

type ReportStatusCodeRealTimeEdgeServiceParameters struct {
}

func (s ReportStatusCodeRealTimeEdgeServiceParameters) String() string {
  return tea.Prettify(s)
}

func (s ReportStatusCodeRealTimeEdgeServiceParameters) GoString() string {
  return s.String()
}

type ReportStatusCodeRealTimeEdgeServiceRequestHeader struct {
}

func (s ReportStatusCodeRealTimeEdgeServiceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportStatusCodeRealTimeEdgeServiceRequestHeader) GoString() string {
  return s.String()
}

type ReportStatusCodeRealTimeEdgeServiceResponseHeader struct {
}

func (s ReportStatusCodeRealTimeEdgeServiceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportStatusCodeRealTimeEdgeServiceResponseHeader) GoString() string {
  return s.String()
}




type QueryStatusCodeDistributionOfeachISPandProvinceRequest struct {
  // {"en":"Start Time:\n\n1.The Time format is yyyy-MM-ddTHH:mm:ss±HH:mm. Please note: ±HH:mm is the time zone offset, which can be adjusted according to your data needs, for example, +00:00 represents UTC time, +08:00 represents East 8th District, and -05:00 represents West 5th District. 2024-01-15T10:30:45+00:00 means UTC time January 15, 2024 10:30:45 AM\n\n2.Cannot exceed the current time\n\n3.Up to the past six months (183 days) of data can be obtained","zh_CN":"开始时间：\n\n1.时间格式为 yyyy-MM-ddTHH:mm:ss±HH:mm。请注意：±HH:mm 为时区偏移量，如 +00:00 代表 UTC 时间，+08:00 代表东八区，2024-01-15T10:30:45+00:00，表示UTC 时间 2024 年 1 月 15 日上午 10 点 30 分 45 秒\n\n2.不能大于当前时间\n\n3.最多可获取最近半年（183天）的数据"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty"`
  // {"en":"End time\n1. The time format is yyyy-MM-ddTHH:mm:ss±HH:mm. Please note: ±HH:mm is the time zone offset, which can be adjusted according to your data needs, for example, +00:00 represents UTC time, +08:00 represents East 8th District, and -05:00 represents West 5th District. 2024-01-15T10:30:45+00:00 means UTC time January 15, 2024 10:30:45 AM;\n2. the end time is greater than the start time. If the end time is greater than the current time, the current time is taken.\n3. DateFrom and dateTo are not uploaded, defaulting to query the past 24 hours; if only one is not uploaded, throw an exception;\n4. Maximum query interval allowed: 7 days, that is, the difference between dateFrom and dateTo can not exceed 7 days (technical support can be contacted to adjust).","zh_CN":"结束时间\n1.时间格式为 yyyy-MM-ddTHH:mm:ss±HH:mm。请注意：±HH:mm 为时区偏移量，可根据您的数据需要进行调整，例如 +00:00 代表 UTC 时间，+08:00 代表东八区，-05:00 代表西五区。2024-01-15T10:30:45+00:00，表示UTC 时间 2024 年 1 月 15 日上午 10 点 30 分 45 秒；\n2.结束时间需大于开始时间,结束时间如果大于当前时间,取当前时间;\n3.dateFrom,dateTo二者都未传,默认查询过去的24小时;如仅有一个未传,抛异常;\n4.允许查询最大间隔:7天,即dateFrom和dateTo相差不能超过7天(可联系技术支持调整)。"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty"`
  // {"en":"Domain: when domain is not passed: default is all domain names, maximum supported domain names are 20 (can be adjusted by contacting technical support)","zh_CN":"域名：当domain没有传时:默认为全部域名,最大域名支持20个(可联系技术支持调整)"}
  Domain []*string `json:"domain,omitempty" xml:"domain,omitempty" type:"Repeated"`
  // {"en":"Data granularity, 1m: 1-minute 5m: 5-minute granularity, 1h: 1-hour granularity","zh_CN":"数据粒度,1m: 1分钟粒度, 5m:5分钟粒度,1h:1小时粒度"}
  DataInterval *string `json:"dataInterval,omitempty" xml:"dataInterval,omitempty"`
  // {"en":"Province\n\n1.Province is not upload: Query all provinces and aggregate the returned data according to all provinces;\n2.Province is upload: Provinces can transmit Chinese or code. Please refer to the appendix description section of the overview page for the provincial information code table.\n\n3.Support language request header Accept Language, only support zh-CN and en-US, default to zh-CN. Accept Language: en-US, both the province and isp input and return are in code, otherwise the return is in Chinese.","zh_CN":"省份\n\n1.未传递province时：查询所有省份，返回的数据按照所有省份聚合。\n\n2.有传递province时：省份 可传中文或code。省份信息码表详见概览页附录说明章节\n\n3.支持语言请求头Accept-Language，只支持zh-CN、en-US，默认为zh-CN。Accept-Language：en-US时，省份及运营商 入参及返回都为code，否则返回的为中文。"}
  Province []*string `json:"province,omitempty" xml:"province,omitempty" type:"Repeated"`
  // {"en":"ISP:\n1.ISP is not upload: Query all ISPs and aggregate the returned data according to all ISPs;\n2.ISPs is upload: Isp can transmit Chinese or code. Please refer to the appendix description section of the overview page for the ISP information code table.","zh_CN":"运营商：\n1.未传递isp时：查询所有isp，返回的数据按照所有运营商聚合。\n2.有传递isp时：运营商 可传中文或code。运营商信息码表详见概览页附录说明章节"}
  Isp []*string `json:"isp,omitempty" xml:"isp,omitempty" type:"Repeated"`
  // {"en":"Group dimension\n\n1.Options are domain, province, isp, and more than one value can be entered;\n2.The data is displayed according to the specified dimension;","zh_CN":"分组维度\n1.可选值为domain、province、isp,可传入多个值;\n2.有传入则按照该维度展示明细数据;"}
  GroupBy []*string `json:"groupBy,omitempty" xml:"groupBy,omitempty" type:"Repeated"`
}

func (s QueryStatusCodeDistributionOfeachISPandProvinceRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryStatusCodeDistributionOfeachISPandProvinceRequest) GoString() string {
  return s.String()
}

func (s *QueryStatusCodeDistributionOfeachISPandProvinceRequest) SetDateFrom(v string) *QueryStatusCodeDistributionOfeachISPandProvinceRequest {
  s.DateFrom = &v
  return s
}

func (s *QueryStatusCodeDistributionOfeachISPandProvinceRequest) SetDateTo(v string) *QueryStatusCodeDistributionOfeachISPandProvinceRequest {
  s.DateTo = &v
  return s
}

func (s *QueryStatusCodeDistributionOfeachISPandProvinceRequest) SetDomain(v []*string) *QueryStatusCodeDistributionOfeachISPandProvinceRequest {
  s.Domain = v
  return s
}

func (s *QueryStatusCodeDistributionOfeachISPandProvinceRequest) SetDataInterval(v string) *QueryStatusCodeDistributionOfeachISPandProvinceRequest {
  s.DataInterval = &v
  return s
}

func (s *QueryStatusCodeDistributionOfeachISPandProvinceRequest) SetProvince(v []*string) *QueryStatusCodeDistributionOfeachISPandProvinceRequest {
  s.Province = v
  return s
}

func (s *QueryStatusCodeDistributionOfeachISPandProvinceRequest) SetIsp(v []*string) *QueryStatusCodeDistributionOfeachISPandProvinceRequest {
  s.Isp = v
  return s
}

func (s *QueryStatusCodeDistributionOfeachISPandProvinceRequest) SetGroupBy(v []*string) *QueryStatusCodeDistributionOfeachISPandProvinceRequest {
  s.GroupBy = v
  return s
}

type QueryStatusCodeDistributionOfeachISPandProvinceRequestHeader struct {
}

func (s QueryStatusCodeDistributionOfeachISPandProvinceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryStatusCodeDistributionOfeachISPandProvinceRequestHeader) GoString() string {
  return s.String()
}

type QueryStatusCodeDistributionOfeachISPandProvincePaths struct {
}

func (s QueryStatusCodeDistributionOfeachISPandProvincePaths) String() string {
  return tea.Prettify(s)
}

func (s QueryStatusCodeDistributionOfeachISPandProvincePaths) GoString() string {
  return s.String()
}

type QueryStatusCodeDistributionOfeachISPandProvinceParameters struct {
}

func (s QueryStatusCodeDistributionOfeachISPandProvinceParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryStatusCodeDistributionOfeachISPandProvinceParameters) GoString() string {
  return s.String()
}

type QueryStatusCodeDistributionOfeachISPandProvinceResponse struct {
  // {"en":"","zh_CN":""}
  Result []*QueryStatusCodeDistributionOfeachISPandProvinceResponseResult `json:"result,omitempty" xml:"result,omitempty" require:"true" type:"Repeated"`
}

func (s QueryStatusCodeDistributionOfeachISPandProvinceResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryStatusCodeDistributionOfeachISPandProvinceResponse) GoString() string {
  return s.String()
}

func (s *QueryStatusCodeDistributionOfeachISPandProvinceResponse) SetResult(v []*QueryStatusCodeDistributionOfeachISPandProvinceResponseResult) *QueryStatusCodeDistributionOfeachISPandProvinceResponse {
  s.Result = v
  return s
}

type QueryStatusCodeDistributionOfeachISPandProvinceResponseResult struct     {
  // {"en":"Domain","zh_CN":"域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {"en":"","zh_CN":""}
  IspData []*QueryStatusCodeDistributionOfeachISPandProvinceResponseResultIspData `json:"ispData,omitempty" xml:"ispData,omitempty" require:"true" type:"Repeated"`
}

func (s QueryStatusCodeDistributionOfeachISPandProvinceResponseResult) String() string {
  return tea.Prettify(s)
}

func (s QueryStatusCodeDistributionOfeachISPandProvinceResponseResult) GoString() string {
  return s.String()
}

func (s *QueryStatusCodeDistributionOfeachISPandProvinceResponseResult) SetDomain(v string) *QueryStatusCodeDistributionOfeachISPandProvinceResponseResult {
  s.Domain = &v
  return s
}

func (s *QueryStatusCodeDistributionOfeachISPandProvinceResponseResult) SetIspData(v []*QueryStatusCodeDistributionOfeachISPandProvinceResponseResultIspData) *QueryStatusCodeDistributionOfeachISPandProvinceResponseResult {
  s.IspData = v
  return s
}

type QueryStatusCodeDistributionOfeachISPandProvinceResponseResultIspData struct     {
  // {"en":"ISP","zh_CN":"运营商"}
  Isp *string `json:"isp,omitempty" xml:"isp,omitempty" require:"true"`
  // {"en":"","zh_CN":""}
  ProvinceData []*QueryStatusCodeDistributionOfeachISPandProvinceResponseResultIspDataProvinceData `json:"provinceData,omitempty" xml:"provinceData,omitempty" require:"true" type:"Repeated"`
}

func (s QueryStatusCodeDistributionOfeachISPandProvinceResponseResultIspData) String() string {
  return tea.Prettify(s)
}

func (s QueryStatusCodeDistributionOfeachISPandProvinceResponseResultIspData) GoString() string {
  return s.String()
}

func (s *QueryStatusCodeDistributionOfeachISPandProvinceResponseResultIspData) SetIsp(v string) *QueryStatusCodeDistributionOfeachISPandProvinceResponseResultIspData {
  s.Isp = &v
  return s
}

func (s *QueryStatusCodeDistributionOfeachISPandProvinceResponseResultIspData) SetProvinceData(v []*QueryStatusCodeDistributionOfeachISPandProvinceResponseResultIspDataProvinceData) *QueryStatusCodeDistributionOfeachISPandProvinceResponseResultIspData {
  s.ProvinceData = v
  return s
}

type QueryStatusCodeDistributionOfeachISPandProvinceResponseResultIspDataProvinceData struct     {
  // {"en":"Province","zh_CN":"省份"}
  Province *string `json:"province,omitempty" xml:"province,omitempty" require:"true"`
  // {"en":"","zh_CN":""}
  StatusCodeData []*QueryStatusCodeDistributionOfeachISPandProvinceResponseResultIspDataProvinceDataStatusCodeData `json:"statusCodeData,omitempty" xml:"statusCodeData,omitempty" require:"true" type:"Repeated"`
}

func (s QueryStatusCodeDistributionOfeachISPandProvinceResponseResultIspDataProvinceData) String() string {
  return tea.Prettify(s)
}

func (s QueryStatusCodeDistributionOfeachISPandProvinceResponseResultIspDataProvinceData) GoString() string {
  return s.String()
}

func (s *QueryStatusCodeDistributionOfeachISPandProvinceResponseResultIspDataProvinceData) SetProvince(v string) *QueryStatusCodeDistributionOfeachISPandProvinceResponseResultIspDataProvinceData {
  s.Province = &v
  return s
}

func (s *QueryStatusCodeDistributionOfeachISPandProvinceResponseResultIspDataProvinceData) SetStatusCodeData(v []*QueryStatusCodeDistributionOfeachISPandProvinceResponseResultIspDataProvinceDataStatusCodeData) *QueryStatusCodeDistributionOfeachISPandProvinceResponseResultIspDataProvinceData {
  s.StatusCodeData = v
  return s
}

type QueryStatusCodeDistributionOfeachISPandProvinceResponseResultIspDataProvinceDataStatusCodeData struct     {
  // {"en":"Return specific status code details such as 200, 201, 500, as well as aggregated 1XX, 2XX, 3XX, 4XX, 5XX, all, OTHERS. Return when values are available.","zh_CN":"返回具体状态码明细如200,201,500，及聚合的1XX，2XX，3XX，4XX，5XX，all，OTHERS。有值时返回"}
  StatusCode *string `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
  // {"en":"","zh_CN":""}
  RequestData []*QueryStatusCodeDistributionOfeachISPandProvinceResponseResultIspDataProvinceDataStatusCodeDataRequestData `json:"requestData,omitempty" xml:"requestData,omitempty" require:"true" type:"Repeated"`
}

func (s QueryStatusCodeDistributionOfeachISPandProvinceResponseResultIspDataProvinceDataStatusCodeData) String() string {
  return tea.Prettify(s)
}

func (s QueryStatusCodeDistributionOfeachISPandProvinceResponseResultIspDataProvinceDataStatusCodeData) GoString() string {
  return s.String()
}

func (s *QueryStatusCodeDistributionOfeachISPandProvinceResponseResultIspDataProvinceDataStatusCodeData) SetStatusCode(v string) *QueryStatusCodeDistributionOfeachISPandProvinceResponseResultIspDataProvinceDataStatusCodeData {
  s.StatusCode = &v
  return s
}

func (s *QueryStatusCodeDistributionOfeachISPandProvinceResponseResultIspDataProvinceDataStatusCodeData) SetRequestData(v []*QueryStatusCodeDistributionOfeachISPandProvinceResponseResultIspDataProvinceDataStatusCodeDataRequestData) *QueryStatusCodeDistributionOfeachISPandProvinceResponseResultIspDataProvinceDataStatusCodeData {
  s.RequestData = v
  return s
}

type QueryStatusCodeDistributionOfeachISPandProvinceResponseResultIspDataProvinceDataStatusCodeDataRequestData struct     {
  // {"en":"Time,\n1.When the data query granularity is 1m, the format is yyyy-MM-dd HH:mm; the data value of every time slice represents the data value within the previous time granularity range. The first time slice of the day is yyyy-MM-dd 00:01 AM, and the last one is (yyyy-MM-dd+1) 00:00;\n2.When the data query granularity is 5m, the format is yyyy-MM-dd HH:mm; the data value of every time slice represents the data value within the previous time granularity range. The first time slice of the day is yyyy-MM-dd 12:05 AM, and the last one is (yyyy-MM-dd+1) 00:00;\n3.When the data query granularity is 1h, the format is yyyy-MM-dd HH; the data value of every time slice represents the data value within the previous time granularity range. The first time slice of the day is yyyy-MM-dd 00:01, and the last one is (yyyy-MM-dd+1) 00;\n4.Return the time slice contained in start time and the time slice contained in end time.","zh_CN":"时间,\n1.查询的数据粒度为1m时,格式为yyyy-MM-dd HH:mm;每一个时间片数据值代表的是前一个时间粒度范围内的数据值。一天开始的时间片是yyyy-MM-dd 00:01,最后一个时间片是(yyyy-MM-dd+1)00:00;\n2.查询的数据粒度为5m时,格式为yyyy-MM-dd HH:mm;每一个时间片数据值代表的是前一个时间粒度范围内的数据值。一天开始的时间片是yyyy-MM-dd 00:05,最后一个时间片是(yyyy-MM-dd+1)00:00;\n3.查询的数据粒度为1h时,格式为yyyy-MM-dd HH;每一个时间片数据值代表的是前一个时间粒度范围内的数据值。一天开始的时间片是yyyy-MM-dd 01,最后一个时间片是(yyyy-MM-dd+1)00;\n4.返回开始时间和结束时间包含的时间片。"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {"en":"Number of requests of the status code","zh_CN":"状态码对应的请求数"}
  Value *string `json:"value,omitempty" xml:"value,omitempty" require:"true"`
}

func (s QueryStatusCodeDistributionOfeachISPandProvinceResponseResultIspDataProvinceDataStatusCodeDataRequestData) String() string {
  return tea.Prettify(s)
}

func (s QueryStatusCodeDistributionOfeachISPandProvinceResponseResultIspDataProvinceDataStatusCodeDataRequestData) GoString() string {
  return s.String()
}

func (s *QueryStatusCodeDistributionOfeachISPandProvinceResponseResultIspDataProvinceDataStatusCodeDataRequestData) SetTimestamp(v string) *QueryStatusCodeDistributionOfeachISPandProvinceResponseResultIspDataProvinceDataStatusCodeDataRequestData {
  s.Timestamp = &v
  return s
}

func (s *QueryStatusCodeDistributionOfeachISPandProvinceResponseResultIspDataProvinceDataStatusCodeDataRequestData) SetValue(v string) *QueryStatusCodeDistributionOfeachISPandProvinceResponseResultIspDataProvinceDataStatusCodeDataRequestData {
  s.Value = &v
  return s
}

type QueryStatusCodeDistributionOfeachISPandProvinceResponseHeader struct {
}

func (s QueryStatusCodeDistributionOfeachISPandProvinceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryStatusCodeDistributionOfeachISPandProvinceResponseHeader) GoString() string {
  return s.String()
}




type ReportStatusCodeLogServiceRequest struct {
  // {'en':'From date:
  // 
  // 1.The time format is yyyy-MM-ddTHH:mm:ss±HH:mm. Please note: ±HH:mm is the time zone offset, which can be adjusted according to your data needs, for example, +00:00 represents UTC time, +08:00 represents East 8th District, and -05:00 represents West 5th District. 2024-01-15T10:30:45+00:00 means UTC time January 15, 2024 10:30:45 AM
  // 
  //  2.Cannot exceed current time
  // 
  // 3.The most recent six-month (183 days) data are available.', 'zh_CN':'开始时间：
  // 
  // 1.时间格式为 yyyy-MM-ddTHH:mm:ss±HH:mm。请注意：±HH:mm 为时区偏移量，可根据您的数据需要进行调整，例如 +00:00 代表 UTC 时间，+08:00 代表东八区，-05:00 代表西五区。2024-01-15T10:30:45+00:00，表示UTC 时间 2024 年 1 月 15 日上午 10 点 30 分 45 秒
  // 
  // 2.不能大于当前时间
  // 
  // 3.最多可获取最近半年（183天）的数据'}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty"`
  // {'en':'End time:
  // 
  // 1.The time format is yyyy-MM-ddTHH:mm:ss±HH:mm. Please note: ±HH:mm is the time zone offset, which can be adjusted according to your data needs, for example, +00:00 represents UTC time, +08:00 represents East 8th District, and -05:00 represents West 5th District. 2024-01-15T10:30:45+00:00 means UTC time January 15, 2024 10:30:45 AM
  // 
  // 2.The end time needs to be greater than the start time. If the end time is greater than the current time, take the current time.
  // 
  // 3.If both fields of dataFrom and dateTo are left empty,  the default query past 1 day; If there is only one unsent, throw an exception
  // 
  // 4.Maximum allowed query time interval: 31 days, Date from and dateTo, not more than 31 days', 'zh_CN':'结束时间：
  // 
  // 1.时间格式为 yyyy-MM-ddTHH:mm:ss±HH:mm。请注意：±HH:mm 为时区偏移量，可根据您的数据需要进行调整，例如 +00:00 代表 UTC 时间，+08:00 代表东八区，-05:00 代表西五区。2024-01-15T10:30:45+00:00，表示UTC 时间 2024 年 1 月 15 日上午 10 点 30 分 45 秒
  // 
  // 2.结束时间需大于开始时间，结束时间如果大于当前时间，取当前时间
  // 
  // 3.dateFrom，dateTo二者都未传，默认查询过去的1天；如仅有一个未传，抛异常
  // 
  // 4.允许查询最大时间间隔：31天，即dateFrom和dateTo相差不能超过31天。（可联系技术支持调整）'}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty"`
  // {'en':'Domain:
  // 
  // 1.Allowable maximum number of domain is 20 (can be adjusted by contacting technical support).
  // 
  // 2.Automatically filter out illegal domain names (pass illegal domain names, will be filtered out, the query results only return the data of legitimate domain names)', 'zh_CN':'域名：
  // 
  // 1、可传递域名数量上限默认为20个（可联系技术支持调整）；
  // 
  // 2、自动过滤掉无效域名（如传递非法域名，会被过滤掉，查询结果只返回有效域名的数据）。'}
  Domain []*string `json:"domain,omitempty" xml:"domain,omitempty" type:"Repeated"`
  // {'en':'1.Selection:domain
  // 2.If groupBy left empty, merge date of all domains', 'zh_CN':'1. 可选值：domain
  // 2. 不传默认聚合所有频道数据'}
  GroupBy []*string `json:"groupBy,omitempty" xml:"groupBy,omitempty" type:"Repeated"`
}

func (s ReportStatusCodeLogServiceRequest) String() string {
  return tea.Prettify(s)
}

func (s ReportStatusCodeLogServiceRequest) GoString() string {
  return s.String()
}

func (s *ReportStatusCodeLogServiceRequest) SetDateFrom(v string) *ReportStatusCodeLogServiceRequest {
  s.DateFrom = &v
  return s
}

func (s *ReportStatusCodeLogServiceRequest) SetDateTo(v string) *ReportStatusCodeLogServiceRequest {
  s.DateTo = &v
  return s
}

func (s *ReportStatusCodeLogServiceRequest) SetDomain(v []*string) *ReportStatusCodeLogServiceRequest {
  s.Domain = v
  return s
}

func (s *ReportStatusCodeLogServiceRequest) SetGroupBy(v []*string) *ReportStatusCodeLogServiceRequest {
  s.GroupBy = v
  return s
}

type ReportStatusCodeLogServiceResponse struct {
  // {'en':'code', 'zh_CN':'接口返回状态'}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {'en':'message', 'zh_CN':'接口返回信息'}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  Data []*ReportStatusCodeLogServiceResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s ReportStatusCodeLogServiceResponse) String() string {
  return tea.Prettify(s)
}

func (s ReportStatusCodeLogServiceResponse) GoString() string {
  return s.String()
}

func (s *ReportStatusCodeLogServiceResponse) SetCode(v string) *ReportStatusCodeLogServiceResponse {
  s.Code = &v
  return s
}

func (s *ReportStatusCodeLogServiceResponse) SetMessage(v string) *ReportStatusCodeLogServiceResponse {
  s.Message = &v
  return s
}

func (s *ReportStatusCodeLogServiceResponse) SetData(v []*ReportStatusCodeLogServiceResponseData) *ReportStatusCodeLogServiceResponse {
  s.Data = v
  return s
}

type ReportStatusCodeLogServiceResponseData struct     {
  // {'en':'Domain. If merge date of all domains will not return domain', 'zh_CN':'域名，聚合全部域名数据不返回该字段'}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  StatusCodeDataList []*ReportStatusCodeLogServiceResponseDataStatusCodeDataList `json:"statusCodeDataList,omitempty" xml:"statusCodeDataList,omitempty" require:"true" type:"Repeated"`
}

func (s ReportStatusCodeLogServiceResponseData) String() string {
  return tea.Prettify(s)
}

func (s ReportStatusCodeLogServiceResponseData) GoString() string {
  return s.String()
}

func (s *ReportStatusCodeLogServiceResponseData) SetDomain(v string) *ReportStatusCodeLogServiceResponseData {
  s.Domain = &v
  return s
}

func (s *ReportStatusCodeLogServiceResponseData) SetStatusCodeDataList(v []*ReportStatusCodeLogServiceResponseDataStatusCodeDataList) *ReportStatusCodeLogServiceResponseData {
  s.StatusCodeDataList = v
  return s
}

type ReportStatusCodeLogServiceResponseDataStatusCodeDataList struct     {
  // {'en':'Status code', 'zh_CN':'状态码'}
  StatusCode *string `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
  DetailList []*ReportStatusCodeLogServiceResponseDataStatusCodeDataListDetailList `json:"detailList,omitempty" xml:"detailList,omitempty" require:"true" type:"Repeated"`
}

func (s ReportStatusCodeLogServiceResponseDataStatusCodeDataList) String() string {
  return tea.Prettify(s)
}

func (s ReportStatusCodeLogServiceResponseDataStatusCodeDataList) GoString() string {
  return s.String()
}

func (s *ReportStatusCodeLogServiceResponseDataStatusCodeDataList) SetStatusCode(v string) *ReportStatusCodeLogServiceResponseDataStatusCodeDataList {
  s.StatusCode = &v
  return s
}

func (s *ReportStatusCodeLogServiceResponseDataStatusCodeDataList) SetDetailList(v []*ReportStatusCodeLogServiceResponseDataStatusCodeDataListDetailList) *ReportStatusCodeLogServiceResponseDataStatusCodeDataList {
  s.DetailList = v
  return s
}

type ReportStatusCodeLogServiceResponseDataStatusCodeDataListDetailList struct     {
  // {'en':'timestamp,Returns the timestamp between the start time and end time.Time format:
  //                                                                                  Hours: yyyy MM DD hh:00:00', 
  //                                                                                  'zh_CN':'时间片,返回开始时间和结束时间包含的时间片。
  //                                                                                  时间格式:
  //                                                                                  小时：yyyy-MM-dd HH:00:00'}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {'en':'Status code times', 'zh_CN':'状态码次数'}
  Value *string `json:"value,omitempty" xml:"value,omitempty" require:"true"`
}

func (s ReportStatusCodeLogServiceResponseDataStatusCodeDataListDetailList) String() string {
  return tea.Prettify(s)
}

func (s ReportStatusCodeLogServiceResponseDataStatusCodeDataListDetailList) GoString() string {
  return s.String()
}

func (s *ReportStatusCodeLogServiceResponseDataStatusCodeDataListDetailList) SetTimestamp(v string) *ReportStatusCodeLogServiceResponseDataStatusCodeDataListDetailList {
  s.Timestamp = &v
  return s
}

func (s *ReportStatusCodeLogServiceResponseDataStatusCodeDataListDetailList) SetValue(v string) *ReportStatusCodeLogServiceResponseDataStatusCodeDataListDetailList {
  s.Value = &v
  return s
}

type ReportStatusCodeLogServicePaths struct {
}

func (s ReportStatusCodeLogServicePaths) String() string {
  return tea.Prettify(s)
}

func (s ReportStatusCodeLogServicePaths) GoString() string {
  return s.String()
}

type ReportStatusCodeLogServiceParameters struct {
}

func (s ReportStatusCodeLogServiceParameters) String() string {
  return tea.Prettify(s)
}

func (s ReportStatusCodeLogServiceParameters) GoString() string {
  return s.String()
}

type ReportStatusCodeLogServiceRequestHeader struct {
}

func (s ReportStatusCodeLogServiceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportStatusCodeLogServiceRequestHeader) GoString() string {
  return s.String()
}

type ReportStatusCodeLogServiceResponseHeader struct {
}

func (s ReportStatusCodeLogServiceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportStatusCodeLogServiceResponseHeader) GoString() string {
  return s.String()
}




type QueryStatusCodeDistributionRequest struct {
  // {"en":"Start time:
  // 1.The time format is yyyy-MM-ddTHH:mm:ss±HH:mm. Please note: ±HH:mm is the time zone offset, which can be adjusted according to your data needs, for example, +00:00 represents UTC time, +08:00 represents East 8th District, and -05:00 represents West 5th District. 2024-01-15T10:30:45+00:00 means UTC time January 15, 2024 10:30:45 AM;
  // 2.Must be a time that is 183 days earlier than the current time, and the time must be earlier than the current time and dateTo;
  // 3.Period between dataFrom and dateTo cannot be longer than 7 days(technical support can be contacted to adjust);
  // 4.dateFrom and dateTo can be either both are specified or neither is specifies;
  // 5.If neither dateFrom nor dateTo is specified, then by default, data in the last 24 hour is queried", "zh_CN":"开始时间
  // 1.时间格式为 yyyy-MM-ddTHH:mm:ss±HH:mm。请注意：±HH:mm 为时区偏移量，可根据您的数据需要进行调整，例如 +00:00 代表 UTC 时间，+08:00 代表东八区，-05:00 代表西五区。2024-01-15T10:30:45+00:00，表示UTC 时间 2024 年 1 月 15 日上午 10 点 30 分 45 秒；
  // 2.必须大于当前时间-183天，并且小于当前时间和dateTo；
  // 3.dateFrom和dateTo相差不能超过7天(可联系技术支持调整)；
  // 4.dateFrom和dateTo要么都传递，要么都不传递；
  // 5.dateFrom和dateTo都未传递，则默认查询过去24小时的数据"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty"`
  // {"en":"End time:
  // 1.The time format is yyyy-MM-ddTHH:mm:ss±HH:mm. Please note: ±HH:mm is the time zone offset, which can be adjusted according to your data needs, for example, +00:00 represents UTC time, +08:00 represents East 8th District, and -05:00 represents West 5th District. 2024-01-15T10:30:45+00:00 means UTC time January 15, 2024 10:30:45 AM;
  // 2.Must be greater than dateFrom; if it's greater than the current time, then the current time is assigned as the value;", "zh_CN":"结束时间
  // 1.时间格式为 yyyy-MM-ddTHH:mm:ss±HH:mm。请注意：±HH:mm 为时区偏移量，可根据您的数据需要进行调整，例如 +00:00 代表 UTC 时间，+08:00 代表东八区，-05:00 代表西五区。2024-01-15T10:30:45+00:00，表示UTC 时间 2024 年 1 月 15 日上午 10 点 30 分 45 秒；
  // 2.必须大于dateFrom；如果大于当前时间，则重新赋值为当前时间；"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty"`
  // {"en":"Domain names:
  // 1.Domain number limits can be adjusted depending on different accounts. The default value is 20
  // 2.Query all domain names under account when this entry is not passed", "zh_CN":"域名：
  // 1.可传递域名数量上限默认为20个(可联系技术支持调整)；
  // 2.未传递该入参时查询账号下所有域名"}
  Domain []*string `json:"domain,omitempty" xml:"domain,omitempty" type:"Repeated"`
  // {"en":"Data granularity, 5m: granularity of 5 minutes", "zh_CN":"数据粒度，5m：5分钟粒度"}
  DataInterval *string `json:"dataInterval,omitempty" xml:"dataInterval,omitempty"`
  // {"en":"Group dimension
  // 1.The value can be selected is domain;
  // 2.The data is displayed according to the specified dimension;", "zh_CN":"分组维度
  // 1.可选值为domain；
  // 2.有传入则按照该维度展示明细数据；"}
  GroupBy []*string `json:"groupBy,omitempty" xml:"groupBy,omitempty" type:"Repeated"`
  // {"en":"1.If 0 is added, the value is true or false, which is false by default
  // 2.When the value of data adding is true, data is added to time points without data
  // 3.When the dataPadding value is false, no treatment will be done", "zh_CN":"是否补0，取值为true或false，默认为false
  // 当dataPadding取值为true时，对没有数据的时间点填充数据，取值为0
  // 当dataPadding取值为false时，不做处理"}
  DataPadding *bool `json:"dataPadding,omitempty" xml:"dataPadding,omitempty"`
  // {"en":"Query dimension. Optional values: statusCode, statusCodeType. Default value is statuscode.
  // 1.statusCode: returns the status code details;
  // 2.statusCodeType: returns the requests of eachstatus code type (such as the number of requests corresponding to success, redirect, not modified, permission, not found, server error, and other)", "zh_CN":"查询维度，可选值：statusCode、statusCodeType；不传默认statusCode
  // 1.statusCode ：返回状态码明细；
  // 2.statusCodeType：返回状态码类型对应明细(如Success、Redirect、Not-Modified、Permission、Not-Found、Server Error、Other对应的请求数)"}
  QueryBy *string `json:"queryBy,omitempty" xml:"queryBy,omitempty"`
}

func (s QueryStatusCodeDistributionRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryStatusCodeDistributionRequest) GoString() string {
  return s.String()
}

func (s *QueryStatusCodeDistributionRequest) SetDateFrom(v string) *QueryStatusCodeDistributionRequest {
  s.DateFrom = &v
  return s
}

func (s *QueryStatusCodeDistributionRequest) SetDateTo(v string) *QueryStatusCodeDistributionRequest {
  s.DateTo = &v
  return s
}

func (s *QueryStatusCodeDistributionRequest) SetDomain(v []*string) *QueryStatusCodeDistributionRequest {
  s.Domain = v
  return s
}

func (s *QueryStatusCodeDistributionRequest) SetDataInterval(v string) *QueryStatusCodeDistributionRequest {
  s.DataInterval = &v
  return s
}

func (s *QueryStatusCodeDistributionRequest) SetGroupBy(v []*string) *QueryStatusCodeDistributionRequest {
  s.GroupBy = v
  return s
}

func (s *QueryStatusCodeDistributionRequest) SetDataPadding(v bool) *QueryStatusCodeDistributionRequest {
  s.DataPadding = &v
  return s
}

func (s *QueryStatusCodeDistributionRequest) SetQueryBy(v string) *QueryStatusCodeDistributionRequest {
  s.QueryBy = &v
  return s
}

type QueryStatusCodeDistributionResponse struct {
  // {"en":"result", "zh_CN":"结果"}
  Result []*QueryStatusCodeDistributionResponseResult `json:"result,omitempty" xml:"result,omitempty" require:"true" type:"Repeated"`
}

func (s QueryStatusCodeDistributionResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryStatusCodeDistributionResponse) GoString() string {
  return s.String()
}

func (s *QueryStatusCodeDistributionResponse) SetResult(v []*QueryStatusCodeDistributionResponseResult) *QueryStatusCodeDistributionResponse {
  s.Result = v
  return s
}

type QueryStatusCodeDistributionResponseResult struct     {
  // {"en":"Domain", "zh_CN":"域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {"en":"statusCodeData", "zh_CN":"状态码数据"}
  StatusCodeData []*QueryStatusCodeDistributionResponseResultStatusCodeData `json:"statusCodeData,omitempty" xml:"statusCodeData,omitempty" require:"true" type:"Repeated"`
}

func (s QueryStatusCodeDistributionResponseResult) String() string {
  return tea.Prettify(s)
}

func (s QueryStatusCodeDistributionResponseResult) GoString() string {
  return s.String()
}

func (s *QueryStatusCodeDistributionResponseResult) SetDomain(v string) *QueryStatusCodeDistributionResponseResult {
  s.Domain = &v
  return s
}

func (s *QueryStatusCodeDistributionResponseResult) SetStatusCodeData(v []*QueryStatusCodeDistributionResponseResultStatusCodeData) *QueryStatusCodeDistributionResponseResult {
  s.StatusCodeData = v
  return s
}

type QueryStatusCodeDistributionResponseResultStatusCodeData struct     {
  // {"en":"Status code", "zh_CN":"状态码"}
  StatusCode *string `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
  // {"en":"totalRequest", "zh_CN":"总请求数"}
  TotalRequest *string `json:"totalRequest,omitempty" xml:"totalRequest,omitempty" require:"true"`
  // {"en":"Success, Redirect, Not-Modified, Permission, Not-Found, Server Error, Other", "zh_CN":"Success, Redirect, Not-Modified, Permission, Not-Found, Server Error, Other"}
  StatusCodeType *string `json:"statusCodeType,omitempty" xml:"statusCodeType,omitempty" require:"true"`
  // {"en":"requestData", "zh_CN":"请求数数据"}
  RequestData []*QueryStatusCodeDistributionResponseResultStatusCodeDataRequestData `json:"requestData,omitempty" xml:"requestData,omitempty" require:"true" type:"Repeated"`
}

func (s QueryStatusCodeDistributionResponseResultStatusCodeData) String() string {
  return tea.Prettify(s)
}

func (s QueryStatusCodeDistributionResponseResultStatusCodeData) GoString() string {
  return s.String()
}

func (s *QueryStatusCodeDistributionResponseResultStatusCodeData) SetStatusCode(v string) *QueryStatusCodeDistributionResponseResultStatusCodeData {
  s.StatusCode = &v
  return s
}

func (s *QueryStatusCodeDistributionResponseResultStatusCodeData) SetTotalRequest(v string) *QueryStatusCodeDistributionResponseResultStatusCodeData {
  s.TotalRequest = &v
  return s
}

func (s *QueryStatusCodeDistributionResponseResultStatusCodeData) SetStatusCodeType(v string) *QueryStatusCodeDistributionResponseResultStatusCodeData {
  s.StatusCodeType = &v
  return s
}

func (s *QueryStatusCodeDistributionResponseResultStatusCodeData) SetRequestData(v []*QueryStatusCodeDistributionResponseResultStatusCodeDataRequestData) *QueryStatusCodeDistributionResponseResultStatusCodeData {
  s.RequestData = v
  return s
}

type QueryStatusCodeDistributionResponseResultStatusCodeDataRequestData struct     {
  // {"en":"DateTime, the format is yyyy-MM-dd HH:mm; the data value of every time slice represents the data value within the previous time granularity range.", "zh_CN":"时间，格式为yyyy-MM-dd HH:mm；每一个时间片数据值代表的是前一个时间粒度范围内的数据值。比如yyyy-MM-dd 00:05，代表00:00到00:05范围内的数据。"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {"en":"Number of requests for status codes", "zh_CN":"状态码对应的请求数"}
  Value *string `json:"value,omitempty" xml:"value,omitempty" require:"true"`
}

func (s QueryStatusCodeDistributionResponseResultStatusCodeDataRequestData) String() string {
  return tea.Prettify(s)
}

func (s QueryStatusCodeDistributionResponseResultStatusCodeDataRequestData) GoString() string {
  return s.String()
}

func (s *QueryStatusCodeDistributionResponseResultStatusCodeDataRequestData) SetTimestamp(v string) *QueryStatusCodeDistributionResponseResultStatusCodeDataRequestData {
  s.Timestamp = &v
  return s
}

func (s *QueryStatusCodeDistributionResponseResultStatusCodeDataRequestData) SetValue(v string) *QueryStatusCodeDistributionResponseResultStatusCodeDataRequestData {
  s.Value = &v
  return s
}

type QueryStatusCodeDistributionPaths struct {
}

func (s QueryStatusCodeDistributionPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryStatusCodeDistributionPaths) GoString() string {
  return s.String()
}

type QueryStatusCodeDistributionParameters struct {
}

func (s QueryStatusCodeDistributionParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryStatusCodeDistributionParameters) GoString() string {
  return s.String()
}

type QueryStatusCodeDistributionRequestHeader struct {
}

func (s QueryStatusCodeDistributionRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryStatusCodeDistributionRequestHeader) GoString() string {
  return s.String()
}

type QueryStatusCodeDistributionResponseHeader struct {
}

func (s QueryStatusCodeDistributionResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryStatusCodeDistributionResponseHeader) GoString() string {
  return s.String()
}




type QueryRealTimeOriginStatusCodeRequest struct {
  // {"en":"Start date:
  // 1.The time format is yyyy-MM-ddTHH:mm:ss±HH:mm. Please note: ±HH:mm is the time zone offset, which can be adjusted according to your data needs, for example, +00:00 represents UTC time, +08:00 represents East 8th District, and -05:00 represents West 5th District. 2024-01-15T10:30:45+00:00 means UTC time January 15, 2024 10:30:45 AM
  // 2.Cannot exceed current time
  // 3.The most recent six-month (183 days) data are available.", "zh_CN":"开始时间：
  // 1.时间格式为 yyyy-MM-ddTHH:mm:ss±HH:mm。请注意：±HH:mm 为时区偏移量，可根据您的数据需要进行调整，例如 +00:00 代表 UTC 时间，+08:00 代表东八区，-05:00 代表西五区。2024-01-15T10:30:45+00:00，表示UTC 时间 2024 年 1 月 15 日上午 10 点 30 分 45 秒；
  // 2.不能大于当前时间
  // 3.最多可获取最近半年(183天)的数据。"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty"`
  // {"en":"End time:
  // 1.The time format is yyyy-MM-ddTHH:mm:ss±HH:mm. Please note: ±HH:mm is the time zone offset, which can be adjusted according to your data needs, for example, +00:00 represents UTC time, +08:00 represents East 8th District, and -05:00 represents West 5th District. 2024-01-15T10:30:45+00:00 means UTC time January 15, 2024 10:30:45 AM
  // 2.The end time needs to be greater than the start time. If the end time is greater than the current time, take the current time.
  // 3.Date from, Date to both, the default query past 24 hours; If there is only one unsent, throw an exception
  // 4.The default maximum query time interval is 24 hours (you can contact technical support to adjust it, up to 31 days)", "zh_CN":"结束时间：
  // 1.时间格式为 yyyy-MM-ddTHH:mm:ss±HH:mm。请注意：±HH:mm 为时区偏移量，可根据您的数据需要进行调整，例如 +00:00 代表 UTC 时间，+08:00 代表东八区，-05:00 代表西五区。2024-01-15T10:30:45+00:00，表示UTC 时间 2024 年 1 月 15 日上午 10 点 30 分 45 秒
  // 2.结束时间需大于开始时间，结束时间如果大于当前时间，取当前时间。
  // 3.dateFrom，dateTo二者都未传，默认查询过去的24小时；如仅有一个未传，抛异常
  // 4.默认允许查询最大时间间隔：24小时(可联系技术支持调整，最大31天)"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty"`
  // {"en":"Domain name:
  // 1.The maximum number of TLDs that can be delivered is 20 by default (contact technical support adjustment);
  // 2.Automatically filter out illegal domain names (pass illegal domain names, will be filtered out, the query results only return the data of legitimate domain names)
  // 3.Domain name exceeding limit, misstatement", "zh_CN":"域名：
  // 1.可传递域名数量上限默认为20个(可联系技术支持调整)；
  // 2.自动过滤掉非法域名(如传递非法域名，会被过滤掉，查询结果只返回合法域名的数据)
  // 3.域名超过上限，报错"}
  Domain []*string `json:"domain,omitempty" xml:"domain,omitempty" require:"true" type:"Repeated"`
  // {"en":"Data granularity:  
  // 1.The default is 1m ;
  // 2.Support 1m (1 minute), 5m (5 minutes)", "zh_CN":"数据粒度：
  // 1.不传默认1m
  // 2.支持1m(1分钟)、5m(5分钟)"}
  DataInterval *string `json:"dataInterval,omitempty" xml:"dataInterval,omitempty"`
  // {"en":"Query dimension:
  // 1.Optional value statusCode, statusCodeType, 2XX, 3XX, 4XX, 5XX. Default value is statusCode;
  // 2.StatusCode: returns the source status code details
  // 3.statusCodeType: returns the requests of eachstatus code type (such as the number of requests corresponding to success, redirect, not modified, permission, not found, server error, and other)
  // 4.2XX:Returns the 2XX return source status code summary and the 2 start return source status code detail data;
  // 5.3XX:Returns the 2XX return source status code summary and the 3 start return source status code detail data;
  // 6.4XX:Returns the 2XX return source status code summary and the 4 start return source status code detail data;
  // 7.5XX:Returns the 2XX return source status code summary and the 5 start return source status code detail data;", "zh_CN":"查询维度:
  // 1.可选值 statusCode，statusCodeType， 2XX ，3XX， 4XX， 5XX，不传默认 statusCode
  // 2.statusCode: 返回回源状态码明细；
  // 3.statusCodeType：返回状态码类型对应明细(如Success、Redirect、Not-Modified、Permission、Not-Found、Server Error、Other对应的请求数)
  // 4.2XX : 返回 2XX 回源状态码汇总及各 2 开头回源状态码明细数据；
  // 5.3XX : 返回 3XX 回源状态码汇总及各 2 开头回源状态码明细数据；
  // 6.4XX : 返回 4XX 回源状态码汇总及各 4 开头回源状态码明细数据；
  // 7.5XX : 返回 5XX回源状态码汇总及各 5 开头回源状态码明细数据；"}
  QueryBy *string `json:"queryBy,omitempty" xml:"queryBy,omitempty"`
}

func (s QueryRealTimeOriginStatusCodeRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryRealTimeOriginStatusCodeRequest) GoString() string {
  return s.String()
}

func (s *QueryRealTimeOriginStatusCodeRequest) SetDateFrom(v string) *QueryRealTimeOriginStatusCodeRequest {
  s.DateFrom = &v
  return s
}

func (s *QueryRealTimeOriginStatusCodeRequest) SetDateTo(v string) *QueryRealTimeOriginStatusCodeRequest {
  s.DateTo = &v
  return s
}

func (s *QueryRealTimeOriginStatusCodeRequest) SetDomain(v []*string) *QueryRealTimeOriginStatusCodeRequest {
  s.Domain = v
  return s
}

func (s *QueryRealTimeOriginStatusCodeRequest) SetDataInterval(v string) *QueryRealTimeOriginStatusCodeRequest {
  s.DataInterval = &v
  return s
}

func (s *QueryRealTimeOriginStatusCodeRequest) SetQueryBy(v string) *QueryRealTimeOriginStatusCodeRequest {
  s.QueryBy = &v
  return s
}

type QueryRealTimeOriginStatusCodeResponse struct {
  // {"en":"request result status code", "zh_CN":"请求结果状态码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Request result information", "zh_CN":"请求结果信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Detailed data on the results of the request", "zh_CN":"请求结果的详细数据"}
  Data []*QueryRealTimeOriginStatusCodeResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
  // {"en":"StatusCode", "zh_CN":"状态码"}
  StatusCode *string `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
  // {"en":"Success, Redirect, Not-Modified, Permission, Not-Found, Server Error, Other", "zh_CN":"Success、Redirect、Not-Modified、Permission、Not-Found、Server Error、Other"}
  StatusCodeType *string `json:"statusCodeType,omitempty" xml:"statusCodeType,omitempty" require:"true"`
}

func (s QueryRealTimeOriginStatusCodeResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryRealTimeOriginStatusCodeResponse) GoString() string {
  return s.String()
}

func (s *QueryRealTimeOriginStatusCodeResponse) SetCode(v string) *QueryRealTimeOriginStatusCodeResponse {
  s.Code = &v
  return s
}

func (s *QueryRealTimeOriginStatusCodeResponse) SetMessage(v string) *QueryRealTimeOriginStatusCodeResponse {
  s.Message = &v
  return s
}

func (s *QueryRealTimeOriginStatusCodeResponse) SetData(v []*QueryRealTimeOriginStatusCodeResponseData) *QueryRealTimeOriginStatusCodeResponse {
  s.Data = v
  return s
}

func (s *QueryRealTimeOriginStatusCodeResponse) SetStatusCode(v string) *QueryRealTimeOriginStatusCodeResponse {
  s.StatusCode = &v
  return s
}

func (s *QueryRealTimeOriginStatusCodeResponse) SetStatusCodeType(v string) *QueryRealTimeOriginStatusCodeResponse {
  s.StatusCodeType = &v
  return s
}

type QueryRealTimeOriginStatusCodeResponseData struct     {
  // {"en":"time, in yyyy-MM-dd HH:MM", "zh_CN":"时间，格式为yyyy-MM-dd HH:mm"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {"en":"Total number of return requests", "zh_CN":"回源总请求数"}
  Value *string `json:"value,omitempty" xml:"value,omitempty" require:"true"`
}

func (s QueryRealTimeOriginStatusCodeResponseData) String() string {
  return tea.Prettify(s)
}

func (s QueryRealTimeOriginStatusCodeResponseData) GoString() string {
  return s.String()
}

func (s *QueryRealTimeOriginStatusCodeResponseData) SetTimestamp(v string) *QueryRealTimeOriginStatusCodeResponseData {
  s.Timestamp = &v
  return s
}

func (s *QueryRealTimeOriginStatusCodeResponseData) SetValue(v string) *QueryRealTimeOriginStatusCodeResponseData {
  s.Value = &v
  return s
}

type QueryRealTimeOriginStatusCodePaths struct {
}

func (s QueryRealTimeOriginStatusCodePaths) String() string {
  return tea.Prettify(s)
}

func (s QueryRealTimeOriginStatusCodePaths) GoString() string {
  return s.String()
}

type QueryRealTimeOriginStatusCodeParameters struct {
}

func (s QueryRealTimeOriginStatusCodeParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryRealTimeOriginStatusCodeParameters) GoString() string {
  return s.String()
}

type QueryRealTimeOriginStatusCodeRequestHeader struct {
}

func (s QueryRealTimeOriginStatusCodeRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryRealTimeOriginStatusCodeRequestHeader) GoString() string {
  return s.String()
}

type QueryRealTimeOriginStatusCodeResponseHeader struct {
}

func (s QueryRealTimeOriginStatusCodeResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryRealTimeOriginStatusCodeResponseHeader) GoString() string {
  return s.String()
}




type QueryMultidomainsIPV6OrIPV4StatusCodeRequest struct {
  // {"en":"Start date:
  // 1.The time format is yyyy-MM-ddTHH:mm:ss±HH:mm. Please note: ±HH:mm is the time zone offset, which can be adjusted according to your data needs, for example, +00:00 represents UTC time, +08:00 represents East 8th District, and -05:00 represents West 5th District. 2024-01-15T10:30:45+00:00 means UTC time January 15, 2024 10:30:45 AM
  // 2.Cannot exceed current time
  // 3.The most recent six-month (183 days) data are available.", "zh_CN":"开始时间:
  // 1.时间格式为 yyyy-MM-ddTHH:mm:ss±HH:mm。请注意：±HH:mm 为时区偏移量，可根据您的数据需要进行调整，例如 +00:00 代表 UTC 时间，+08:00 代表东八区，-05:00 代表西五区。2024-01-15T10:30:45+00:00，表示UTC 时间 2024 年 1 月 15 日上午 10 点 30 分 45 秒;
  // 2.不能大于当前时间
  // 3.最多可获取最近半年(183天)的数据。"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty"`
  // {"en":"End time:
  // 1.The time format is yyyy-MM-ddTHH:mm:ss±HH:mm. Please note: ±HH:mm is the time zone offset, which can be adjusted according to your data needs, for example, +00:00 represents UTC time, +08:00 represents East 8th District, and -05:00 represents West 5th District. 2024-01-15T10:30:45+00:00 means UTC time January 15, 2024 10:30:45 AM
  // 2.The end time needs to be greater than the start time. If the end time is greater than the current time, take the current time.
  // 3.Date from, Date to both, the default query past 24 hours; If there is only one unsent, throw an exception
  // 4.Maximum allowed query time interval: 7 days (with technical support adjustments), meaning that the difference between Date from and dateTo cannot exceed 7 days.", "zh_CN":"结束时间:
  // 1.时间格式为 yyyy-MM-ddTHH:mm:ss±HH:mm。请注意：±HH:mm 为时区偏移量，可根据您的数据需要进行调整，例如 +00:00 代表 UTC 时间，+08:00 代表东八区，-05:00 代表西五区。2024-01-15T10:30:45+00:00，表示UTC 时间 2024 年 1 月 15 日上午 10 点 30 分 45 秒
  // 2.结束时间需大于开始时间,结束时间如果大于当前时间,取当前时间。
  // 3.dateFrom,dateTo二者都未传,默认查询过去的24小时;如仅有一个未传,抛异常
  // 4.允许查询最大时间间隔:7天,即dateFrom和dateTo相差不能超过7天。(可联系技术支持调整)"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty"`
  // {"en":"Domain:
  // 1. The maximum number of domains is 200 by default (Technical Support Adjustment can be contacted);
  // 2. Automatic filtering invalid domain name (if pass illegal domain name, can be filtered, query result only returns the data of valid domain name).", "zh_CN":"域名:
  // 
  // 1.可传递域名数量上限默认为200个(可联系技术支持调整);
  // 
  // 2.自动过滤掉无效域名(如传递非法域名,会被过滤掉,查询结果只返回有效域名的数据)。"}
  Domain []*string `json:"domain,omitempty" xml:"domain,omitempty" type:"Repeated"`
  // {"en":"QueryMultidomainsIPV6OrIPV4StatusCodeData granularity:
  // 1. Support 5m (5 minute granularity), 1h (1 hour granularity);
  // 2. The default is 5m", "zh_CN":"数据粒度:
  // 
  // 1.支持5m(5分钟), 1h(1小时)
  // 
  // 2.不传默认5m。"}
  DataInterval *string `json:"dataInterval,omitempty" xml:"dataInterval,omitempty"`
  // {"en":"IP type:
  // 1.The optional values are IPv6 and IPv4
  // 2.If let this parameter empty,it will query all IP type", "zh_CN":"IP类型:
  // 
  // 1.可选值为 IPV6, IPV4
  // 2.不传默认查询全部"}
  IpType *string `json:"ipType,omitempty" xml:"ipType,omitempty"`
  // {"en":"All region are queried by default, optional values are cn, am, apac, au, emea, hk, tw etc. ", "zh_CN":"1. 不传查全部
  // 2.可选值 cn, am, apac, au, emea, hk, tw 等"}
  Region []*string `json:"region,omitempty" xml:"region,omitempty" type:"Repeated"`
}

func (s QueryMultidomainsIPV6OrIPV4StatusCodeRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryMultidomainsIPV6OrIPV4StatusCodeRequest) GoString() string {
  return s.String()
}

func (s *QueryMultidomainsIPV6OrIPV4StatusCodeRequest) SetDateFrom(v string) *QueryMultidomainsIPV6OrIPV4StatusCodeRequest {
  s.DateFrom = &v
  return s
}

func (s *QueryMultidomainsIPV6OrIPV4StatusCodeRequest) SetDateTo(v string) *QueryMultidomainsIPV6OrIPV4StatusCodeRequest {
  s.DateTo = &v
  return s
}

func (s *QueryMultidomainsIPV6OrIPV4StatusCodeRequest) SetDomain(v []*string) *QueryMultidomainsIPV6OrIPV4StatusCodeRequest {
  s.Domain = v
  return s
}

func (s *QueryMultidomainsIPV6OrIPV4StatusCodeRequest) SetDataInterval(v string) *QueryMultidomainsIPV6OrIPV4StatusCodeRequest {
  s.DataInterval = &v
  return s
}

func (s *QueryMultidomainsIPV6OrIPV4StatusCodeRequest) SetIpType(v string) *QueryMultidomainsIPV6OrIPV4StatusCodeRequest {
  s.IpType = &v
  return s
}

func (s *QueryMultidomainsIPV6OrIPV4StatusCodeRequest) SetRegion(v []*string) *QueryMultidomainsIPV6OrIPV4StatusCodeRequest {
  s.Region = v
  return s
}

type QueryMultidomainsIPV6OrIPV4StatusCodeResponse struct {
  // {"en":"Request status code", "zh_CN":"请求结果状态码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Result message", "zh_CN":"请求结果信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Detail data of result", "zh_CN":"请求结果的详细数据"}
  QueryMultidomainsIPV6OrIPV4StatusCodeData []*QueryMultidomainsIPV6OrIPV4StatusCodeData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s QueryMultidomainsIPV6OrIPV4StatusCodeResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryMultidomainsIPV6OrIPV4StatusCodeResponse) GoString() string {
  return s.String()
}

func (s *QueryMultidomainsIPV6OrIPV4StatusCodeResponse) SetCode(v string) *QueryMultidomainsIPV6OrIPV4StatusCodeResponse {
  s.Code = &v
  return s
}

func (s *QueryMultidomainsIPV6OrIPV4StatusCodeResponse) SetMessage(v string) *QueryMultidomainsIPV6OrIPV4StatusCodeResponse {
  s.Message = &v
  return s
}

func (s *QueryMultidomainsIPV6OrIPV4StatusCodeResponse) SetData(v []*QueryMultidomainsIPV6OrIPV4StatusCodeData) *QueryMultidomainsIPV6OrIPV4StatusCodeResponse {
  s.QueryMultidomainsIPV6OrIPV4StatusCodeData = v
  return s
}

type QueryMultidomainsIPV6OrIPV4StatusCodeData struct {
  // {"en":"Domain", "zh_CN":"域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  DetailList []*QueryMultidomainsIPV6OrIPV4StatusCodeDataDetailList `json:"detailList,omitempty" xml:"detailList,omitempty" require:"true" type:"Repeated"`
}

func (s QueryMultidomainsIPV6OrIPV4StatusCodeData) String() string {
  return tea.Prettify(s)
}

func (s QueryMultidomainsIPV6OrIPV4StatusCodeData) GoString() string {
  return s.String()
}

func (s *QueryMultidomainsIPV6OrIPV4StatusCodeData) SetDomain(v string) *QueryMultidomainsIPV6OrIPV4StatusCodeData {
  s.Domain = &v
  return s
}

func (s *QueryMultidomainsIPV6OrIPV4StatusCodeData) SetDetailList(v []*QueryMultidomainsIPV6OrIPV4StatusCodeDataDetailList) *QueryMultidomainsIPV6OrIPV4StatusCodeData {
  s.DetailList = v
  return s
}

type QueryMultidomainsIPV6OrIPV4StatusCodeDataDetailList struct     {
  // {"en":"Time:
  //         the time slice included in the start time and end time", "zh_CN":"时间片
  // 返回开始时间和结束时间包含的时间片。"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {"en":"Number of status codes", "zh_CN":"状态码个数,单位 个"}
  Value *string `json:"value,omitempty" xml:"value,omitempty" require:"true"`
}

func (s QueryMultidomainsIPV6OrIPV4StatusCodeDataDetailList) String() string {
  return tea.Prettify(s)
}

func (s QueryMultidomainsIPV6OrIPV4StatusCodeDataDetailList) GoString() string {
  return s.String()
}

func (s *QueryMultidomainsIPV6OrIPV4StatusCodeDataDetailList) SetTimestamp(v string) *QueryMultidomainsIPV6OrIPV4StatusCodeDataDetailList {
  s.Timestamp = &v
  return s
}

func (s *QueryMultidomainsIPV6OrIPV4StatusCodeDataDetailList) SetValue(v string) *QueryMultidomainsIPV6OrIPV4StatusCodeDataDetailList {
  s.Value = &v
  return s
}

type QueryMultidomainsIPV6OrIPV4StatusCodePaths struct {
}

func (s QueryMultidomainsIPV6OrIPV4StatusCodePaths) String() string {
  return tea.Prettify(s)
}

func (s QueryMultidomainsIPV6OrIPV4StatusCodePaths) GoString() string {
  return s.String()
}

type QueryMultidomainsIPV6OrIPV4StatusCodeParameters struct {
}

func (s QueryMultidomainsIPV6OrIPV4StatusCodeParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryMultidomainsIPV6OrIPV4StatusCodeParameters) GoString() string {
  return s.String()
}

type QueryMultidomainsIPV6OrIPV4StatusCodeRequestHeader struct {
}

func (s QueryMultidomainsIPV6OrIPV4StatusCodeRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryMultidomainsIPV6OrIPV4StatusCodeRequestHeader) GoString() string {
  return s.String()
}

type QueryMultidomainsIPV6OrIPV4StatusCodeResponseHeader struct {
}

func (s QueryMultidomainsIPV6OrIPV4StatusCodeResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryMultidomainsIPV6OrIPV4StatusCodeResponseHeader) GoString() string {
  return s.String()
}




