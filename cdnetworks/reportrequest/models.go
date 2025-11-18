package reportrequest

// This file is auto-generated, don't edit it. Thanks.
import (
  "github.com/alibabacloud-go/tea/tea"
)

type ReportRequestQuicServiceRequest struct {
  // {"en":"Start time:
  // 1. The time format is yyyy-MM-dd, for example, 2021-10-10;
  // 2. It cannot be greater than the current time
  // 3. Data for the most recent six months (183 days) can be obtained.", "zh_CN":"开始时间：
  //             时间格式为yyyy-MM-dd，例如，2021-10-10；
  //             不能大于当前时间
  //             最多可获取最近半年（183天）的数据。"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty"`
  // {"en":"End time:
  // 1. The time format is yyyy-MM-dd
  // 2. The end time must be greater than the start time. If the end time is greater than the current time, the current time is used.
  // 3. If both dateFrom and dateTo are not passed, the default query is the past day; if only one is not passed, an exception is thrown
  // 4. The maximum query time interval allowed is 7 days, that is, the difference between dateFrom and dateTo cannot exceed 7 days (you can contact technical support to adjust)", "zh_CN":"结束时间：
  //             时间格式为yyyy-MM-dd
  //             结束时间需大于开始时间，结束时间如果大于当前时间，取当前时间。
  //             dateFrom，dateTo二者都未传，默认查询过去的1天；如仅有一个未传，抛异常
  //             允许查询最大时间间隔：7天，即dateFrom和dateTo相差不能超过7天（可联系技术支持调整）。"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty"`
  // {"en":"Domain name:
  // 1. The upper limit of the number of domain names that can be passed is 20 by default (you can contact technical support to adjust it).
  // 2. Automatically filter out illegal domain names (if an illegal domain name is passed, it will be filtered out, and the query result will only return data for legal domain names).
  // 3. When this parameter is not passed, all domain names under the account are queried by default, but an error message will be prompted when the number of domain names under the account exceeds the upper limit.", "zh_CN":"域名：
  //             可传递域名数量上限默认为20个（可联系技术支持调整）。
  //             自动过滤掉非法域名（如传递非法域名，会被过滤掉，查询结果只返回合法域名的数据）
  //             未传递该入参时，默认查询账号下所有域名，但当账号下域名数量超过上限时提示错误"}
  Domain []*string `json:"domain,omitempty" xml:"domain,omitempty" type:"Repeated"`
}

func (s ReportRequestQuicServiceRequest) String() string {
  return tea.Prettify(s)
}

func (s ReportRequestQuicServiceRequest) GoString() string {
  return s.String()
}

func (s *ReportRequestQuicServiceRequest) SetDateFrom(v string) *ReportRequestQuicServiceRequest {
  s.DateFrom = &v
  return s
}

func (s *ReportRequestQuicServiceRequest) SetDateTo(v string) *ReportRequestQuicServiceRequest {
  s.DateTo = &v
  return s
}

func (s *ReportRequestQuicServiceRequest) SetDomain(v []*string) *ReportRequestQuicServiceRequest {
  s.Domain = v
  return s
}

type ReportRequestQuicServiceResponse struct {
  // {"en":"code", "zh_CN":"code"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"message", "zh_CN":"message"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  Data []*ReportRequestQuicServiceResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s ReportRequestQuicServiceResponse) String() string {
  return tea.Prettify(s)
}

func (s ReportRequestQuicServiceResponse) GoString() string {
  return s.String()
}

func (s *ReportRequestQuicServiceResponse) SetCode(v string) *ReportRequestQuicServiceResponse {
  s.Code = &v
  return s
}

func (s *ReportRequestQuicServiceResponse) SetMessage(v string) *ReportRequestQuicServiceResponse {
  s.Message = &v
  return s
}

func (s *ReportRequestQuicServiceResponse) SetData(v []*ReportRequestQuicServiceResponseData) *ReportRequestQuicServiceResponse {
  s.Data = v
  return s
}

type ReportRequestQuicServiceResponseData struct     {
  // {"en":"Time, in the format of yyyy-MM-dd HH:mm:ss", "zh_CN":"时间，格式为yyyy-MM-dd HH:mm:ss"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {"en":"Number of requests, in units", "zh_CN":"请求个数，单位 个"}
  TotalRequest *string `json:"totalRequest,omitempty" xml:"totalRequest,omitempty" require:"true"`
  // {"en":"Number of requests, in units", "zh_CN":"请求个数，单位 个"}
  QuicRequest *string `json:"quicRequest,omitempty" xml:"quicRequest,omitempty" require:"true"`
}

func (s ReportRequestQuicServiceResponseData) String() string {
  return tea.Prettify(s)
}

func (s ReportRequestQuicServiceResponseData) GoString() string {
  return s.String()
}

func (s *ReportRequestQuicServiceResponseData) SetTimestamp(v string) *ReportRequestQuicServiceResponseData {
  s.Timestamp = &v
  return s
}

func (s *ReportRequestQuicServiceResponseData) SetTotalRequest(v string) *ReportRequestQuicServiceResponseData {
  s.TotalRequest = &v
  return s
}

func (s *ReportRequestQuicServiceResponseData) SetQuicRequest(v string) *ReportRequestQuicServiceResponseData {
  s.QuicRequest = &v
  return s
}

type ReportRequestQuicServicePaths struct {
}

func (s ReportRequestQuicServicePaths) String() string {
  return tea.Prettify(s)
}

func (s ReportRequestQuicServicePaths) GoString() string {
  return s.String()
}

type ReportRequestQuicServiceParameters struct {
}

func (s ReportRequestQuicServiceParameters) String() string {
  return tea.Prettify(s)
}

func (s ReportRequestQuicServiceParameters) GoString() string {
  return s.String()
}

type ReportRequestQuicServiceRequestHeader struct {
}

func (s ReportRequestQuicServiceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportRequestQuicServiceRequestHeader) GoString() string {
  return s.String()
}

type ReportRequestQuicServiceResponseHeader struct {
}

func (s ReportRequestQuicServiceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportRequestQuicServiceResponseHeader) GoString() string {
  return s.String()
}




type ReportUserRequestCountryServiceRequest struct {
  // {"en":"Start time:
  // 1. The time format is yyyy-MM-ddTHH:mm:ss±HH:mm. Please note: ±HH:mm is the time zone offset, which can be adjusted according to your data needs, for example, +00:00 represents UTC time, +08:00 represents East 8th District, and -05:00 represents West 5th District. 2024-01-15T10:30:45+00:00 means UTC time January 15, 2024 10:30:45 AM;
  // 2. Can not exceed the current time;
  // 3. The latest half year (183 days) data can be obtained at most.", "zh_CN":"开始时间:
  // 1.时间格式为 yyyy-MM-ddTHH:mm:ss±HH:mm。请注意：±HH:mm 为时区偏移量，可根据您的数据需要进行调整，例如 +00:00 代表 UTC 时间，+08:00 代表东八区，-05:00 代表西五区。2024-01-15T10:30:45+00:00，表示UTC 时间 2024 年 1 月 15 日上午 10 点 30 分 45 秒;
  // 2.不能大于当前时间;
  // 3.最多可获取最近半年(183天)的数据。"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty"`
  // {"en":"End time:
  // 1. The time format is yyyy-MM-ddTHH:mm:ss±HH:mm. Please note: ±HH:mm is the time zone offset, which can be adjusted according to your data needs, for example, +00:00 represents UTC time, +08:00 represents East 8th District, and -05:00 represents West 5th District. 2024-01-15T10:30:45+00:00 means UTC time January 15, 2024 10:30:45 AM;
  // 2. The end time is greater than the start time;
  // 3. If the end time is greater than the current time, the current time is taken.
  // 4. DateFrom and dateTo are not uploaded, defaulting to query the past 24 hours; if only one is not uploaded, throw an exception;
  // 5. The default query interval is 7 days, Maximum query interval allowed: 31 days, that is, the difference between dateFrom and dateTo can not exceed 31 days.", "zh_CN":"结束时间:
  // 1.时间格式为 yyyy-MM-ddTHH:mm:ss±HH:mm。请注意：±HH:mm 为时区偏移量，可根据您的数据需要进行调整，例如 +00:00 代表 UTC 时间，+08:00 代表东八区，-05:00 代表西五区。2024-01-15T10:30:45+00:00，表示UTC 时间 2024 年 1 月 15 日上午 10 点 30 分 45 秒;
  // 2.结束时间需大于开始时间;
  // 3.结束时间如果大于当前时间,取当前时间;
  // 4.dateFrom,dateTo二者都未传,默认查询过去的24小时;如仅有一个未传,抛异常;
  // 5.默认查询间隔7天,允许查询最大间隔:31天,即dateFrom和dateTo相差不能超过31天。"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty"`
  // {"en":"Domains:
  // 1. Domain is not uploaded: Query all domain names of the account (More than 20 domains will error,you can contact technical support for adjustment);
  // 2. Domain is uploaded: Up to 20 domains are supported(you can contact technical support for adjustment).", "zh_CN":"域名:
  // 1.未传递domain时:查询账号下所有全部域名(域名超过20个则报错,可联系技术支持调整);
  // 2.有传递domain时:域名最多支持传20个(可联系技术支持调整)。"}
  Domain []*string `json:"domain,omitempty" xml:"domain,omitempty" type:"Repeated"`
  // {"en":"Acceleration area:
  // 1. Acceleration areaCode is not uploaded: Query all acceleration areas by default.", "zh_CN":"加速区域:
  // 未传递areaCode时,默认查询所有加速区域。"}
  AreaCode []*string `json:"areaCode,omitempty" xml:"areaCode,omitempty" type:"Repeated"`
  // {"en":"Data granularity:
  // 5m: 5 minute granularity;
  // 1h: 1 hour granularity;
  // 1d: 1 day granularity; Default value is 1d.", "zh_CN":"数据粒度:
  // 5m:5分钟粒度;
  // 1h:1小时粒度;
  // 1d:1天粒度。不传默认1天粒度"}
  DataInterval *string `json:"dataInterval,omitempty" xml:"dataInterval,omitempty"`
  // {"en":"Country area:
  // 1. countryCode is not uploaded: Query all country areas by default;
  // 2. countryCode is uploaded: Multiple can be uploaded, such as cn, in.  Please refer to the appendix description section of the overview page.", "zh_CN":"国家区域(含中国台湾、中国澳门、中国香港、中国大陆):
  // 1.未传递countryCode时:查询全部国家区域;
  // 2.有传递countryCode时:可传多个,如cn,in。可传递的值详见概览页附录说明章节"}
  CountryCode []*string `json:"countryCode,omitempty" xml:"countryCode,omitempty" type:"Repeated"`
  // {"en":"Grouped dimension:
  // 1. The optional values are domain, countryCode;Multiple values can be uploaded;
  // 2. If no value is uploaded: Aggregate all data by default.", "zh_CN":"分组维度
  // 可选值为domain、countryCode,可传入多个值;
  // 有传入则按照该维度展示明细数据;
  // 没传默认全部聚合。"}
  GroupBy []*string `json:"groupBy,omitempty" xml:"groupBy,omitempty" type:"Repeated"`
}

func (s ReportUserRequestCountryServiceRequest) String() string {
  return tea.Prettify(s)
}

func (s ReportUserRequestCountryServiceRequest) GoString() string {
  return s.String()
}

func (s *ReportUserRequestCountryServiceRequest) SetDateFrom(v string) *ReportUserRequestCountryServiceRequest {
  s.DateFrom = &v
  return s
}

func (s *ReportUserRequestCountryServiceRequest) SetDateTo(v string) *ReportUserRequestCountryServiceRequest {
  s.DateTo = &v
  return s
}

func (s *ReportUserRequestCountryServiceRequest) SetDomain(v []*string) *ReportUserRequestCountryServiceRequest {
  s.Domain = v
  return s
}

func (s *ReportUserRequestCountryServiceRequest) SetAreaCode(v []*string) *ReportUserRequestCountryServiceRequest {
  s.AreaCode = v
  return s
}

func (s *ReportUserRequestCountryServiceRequest) SetDataInterval(v string) *ReportUserRequestCountryServiceRequest {
  s.DataInterval = &v
  return s
}

func (s *ReportUserRequestCountryServiceRequest) SetCountryCode(v []*string) *ReportUserRequestCountryServiceRequest {
  s.CountryCode = v
  return s
}

func (s *ReportUserRequestCountryServiceRequest) SetGroupBy(v []*string) *ReportUserRequestCountryServiceRequest {
  s.GroupBy = v
  return s
}

type ReportUserRequestCountryServiceResponse struct {
  Result []*ReportUserRequestCountryServiceResponseResult `json:"result,omitempty" xml:"result,omitempty" require:"true" type:"Repeated"`
}

func (s ReportUserRequestCountryServiceResponse) String() string {
  return tea.Prettify(s)
}

func (s ReportUserRequestCountryServiceResponse) GoString() string {
  return s.String()
}

func (s *ReportUserRequestCountryServiceResponse) SetResult(v []*ReportUserRequestCountryServiceResponseResult) *ReportUserRequestCountryServiceResponse {
  s.Result = v
  return s
}

type ReportUserRequestCountryServiceResponseResult struct     {
  // {"en":"Domain", "zh_CN":"域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  CountryData []*ReportUserRequestCountryServiceResponseResultCountryData `json:"countryData,omitempty" xml:"countryData,omitempty" require:"true" type:"Repeated"`
}

func (s ReportUserRequestCountryServiceResponseResult) String() string {
  return tea.Prettify(s)
}

func (s ReportUserRequestCountryServiceResponseResult) GoString() string {
  return s.String()
}

func (s *ReportUserRequestCountryServiceResponseResult) SetDomain(v string) *ReportUserRequestCountryServiceResponseResult {
  s.Domain = &v
  return s
}

func (s *ReportUserRequestCountryServiceResponseResult) SetCountryData(v []*ReportUserRequestCountryServiceResponseResultCountryData) *ReportUserRequestCountryServiceResponseResult {
  s.CountryData = v
  return s
}

type ReportUserRequestCountryServiceResponseResultCountryData struct     {
  // {"en":"Country area", "zh_CN":"国家区域"}
  CountryCode *string `json:"countryCode,omitempty" xml:"countryCode,omitempty" require:"true"`
  // {"en":"Total number of requests.", "zh_CN":"国家的请求总数"}
  TotalRequest *string `json:"totalRequest,omitempty" xml:"totalRequest,omitempty" require:"true"`
  RequestData []*ReportUserRequestCountryServiceResponseResultCountryDataRequestData `json:"requestData,omitempty" xml:"requestData,omitempty" require:"true" type:"Repeated"`
}

func (s ReportUserRequestCountryServiceResponseResultCountryData) String() string {
  return tea.Prettify(s)
}

func (s ReportUserRequestCountryServiceResponseResultCountryData) GoString() string {
  return s.String()
}

func (s *ReportUserRequestCountryServiceResponseResultCountryData) SetCountryCode(v string) *ReportUserRequestCountryServiceResponseResultCountryData {
  s.CountryCode = &v
  return s
}

func (s *ReportUserRequestCountryServiceResponseResultCountryData) SetTotalRequest(v string) *ReportUserRequestCountryServiceResponseResultCountryData {
  s.TotalRequest = &v
  return s
}

func (s *ReportUserRequestCountryServiceResponseResultCountryData) SetRequestData(v []*ReportUserRequestCountryServiceResponseResultCountryDataRequestData) *ReportUserRequestCountryServiceResponseResultCountryData {
  s.RequestData = v
  return s
}

type ReportUserRequestCountryServiceResponseResultCountryDataRequestData struct     {
  // {"en":"Time:
  // 
  // 1. When the data query granularity is 5m, then the format is yyyy-MM-dd HH:mm; ach time slice value represents the value within the previous time granularity range. The first time slice of the day is yyyy-MM-dd 00:05, and the last one is (yyyy-MM-dd+1) 00:00;
  // 2. When the data query granularity is 1h, the format is yyyy-MM-dd HH; Each time slice value represents the value within the previous time granularity range. The first time slice of the day is yyyy-MM-dd 01, and the last one is (yyyy-MM-dd+1) 00;
  // 3. When the data query granularity is 1d, the format is yyyy-MM-dd; Each time slice value represents the value of the day;
  // 4. Return the time slices that contained in start time and in end time.", "zh_CN":"时间,
  // 查询的数据粒度为5m时,格式为yyyy-MM-dd HH:mm;每一个时间片数据值代表的是前一个时间粒度范围内的数据值。一天开始的时间片是yyyy-MM-dd 00:05,最后一个时间片是(yyyy-MM-dd+1) 00:00;
  // 查询的数据粒度为1h时,格式为yyyy-MM-dd HH;每一个时间片数据值代表的是前一个时间粒度范围内的数据值。一天开始的时间片是yyyy-MM-dd 01,最后一个时间片是(yyyy-MM-dd+1) 00;
  // 查询的数据粒度为1d时,格式为yyyy-MM-dd;每一个时间片数据值代表的该天内的数据值;
  // 返回开始时间和结束时间包含的时间片。"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {"en":"Number of requests.", "zh_CN":"请求数"}
  Request *string `json:"request,omitempty" xml:"request,omitempty" require:"true"`
}

func (s ReportUserRequestCountryServiceResponseResultCountryDataRequestData) String() string {
  return tea.Prettify(s)
}

func (s ReportUserRequestCountryServiceResponseResultCountryDataRequestData) GoString() string {
  return s.String()
}

func (s *ReportUserRequestCountryServiceResponseResultCountryDataRequestData) SetTimestamp(v string) *ReportUserRequestCountryServiceResponseResultCountryDataRequestData {
  s.Timestamp = &v
  return s
}

func (s *ReportUserRequestCountryServiceResponseResultCountryDataRequestData) SetRequest(v string) *ReportUserRequestCountryServiceResponseResultCountryDataRequestData {
  s.Request = &v
  return s
}

type ReportUserRequestCountryServicePaths struct {
}

func (s ReportUserRequestCountryServicePaths) String() string {
  return tea.Prettify(s)
}

func (s ReportUserRequestCountryServicePaths) GoString() string {
  return s.String()
}

type ReportUserRequestCountryServiceParameters struct {
}

func (s ReportUserRequestCountryServiceParameters) String() string {
  return tea.Prettify(s)
}

func (s ReportUserRequestCountryServiceParameters) GoString() string {
  return s.String()
}

type ReportUserRequestCountryServiceRequestHeader struct {
}

func (s ReportUserRequestCountryServiceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportUserRequestCountryServiceRequestHeader) GoString() string {
  return s.String()
}

type ReportUserRequestCountryServiceResponseHeader struct {
}

func (s ReportUserRequestCountryServiceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportUserRequestCountryServiceResponseHeader) GoString() string {
  return s.String()
}




type QueryMultidomainsIPV6OrIPV4RequestsRequest struct {
  // {"en":"Start time: 
  //     1. Time format is yyyy-MM-ddTHH:mm:ss+08:00, for example, 2016-12-02T10:00:00+08:00 (10:00 on 2nd of December 2016, Beijing Time); 
  //     2. No bigger than the current time. 3. QueryMultidomainsIPV6OrIPV4RequestsData in the last 183 days at most can be queried.", "zh_CN":"开始时间:
  // 
  // 1.时间格式为yyyy-MM-ddTHH:mm:ss+08:00,例如,2016-12-02T10:00:00+08:00(为北京时间2016年12月2日10点0分0秒);
  // 2.不能大于当前时间
  // 3.最多可获取最近半年(183天)的数据。"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty"`
  // {"en":"End time: 
  //     1. the time format is 2016-12-02T10:00:00+08:00 
  //     2. End time should be greater than start time. If the end time is greater than current time, current time will be used. 
  //     3. If both fields of dataFrom and dateTo are left empty, then data in the last 24 hours will be queried by default; if only one field is filled in and one is left empty, then exception will be occur. 
  //     4. Allowable maximum time range for query:7 day, means the period between dateFrom to dateTo should not exceed 7 day (can be adjusted by contacting technical support).", "zh_CN":"结束时间:
  // 
  // 1.时间格式yyyy-MM-ddTHH:mm:ss+08:00
  // 
  // 2.结束时间需大于开始时间,结束时间如果大于当前时间,取当前时间。
  // 
  // 3.dateFrom,dateTo二者都未传,默认查询过去的24小时;如仅有一个未传,抛异常
  // 
  // 4.允许查询最大时间间隔:7天,即dateFrom和dateTo相差不能超过7天。(可联系技术支持调整)"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty"`
  // {"en":"Domain: 
  //     1. The maximum number of domian is 200 by default (Technical Support Adjustment can be contacted); 
  //     2. automatic filtering invalid domain name (if pass illegal domain name, can be filtered, query result only returns the data of valid domain name).", "zh_CN":"域名:
  // 
  // 1.可传递域名数量上限默认为200个(可联系技术支持调整);
  // 
  // 2.自动过滤掉无效域名(如传递非法域名,会被过滤掉,查询结果只返回有效域名的数据)。"}
  Domain []*string `json:"domain,omitempty" xml:"domain,omitempty" type:"Repeated"`
  // {"en":"QueryMultidomainsIPV6OrIPV4RequestsData granularity 
  //     1. Support 5m (5 minute granularity), 1h (1 hour granularity); 
  //     2. The default is 5m;", "zh_CN":"数据粒度:
  // 
  // 1.支持5m(5分钟)、1h(1小时)
  // 
  // 2.不传默认5m。"}
  DataInterval *string `json:"dataInterval,omitempty" xml:"dataInterval,omitempty"`
  // {"en":"IP type 
  //     1.The optional values are IPv6 and IPv4 
  //     2.If let this parameter empty,it will query all IP type", "zh_CN":"IP类型:
  // 
  // 1.可选值为 IPV6、IPV4
  // 
  // 2.不传默认查询全部"}
  IpType *string `json:"ipType,omitempty" xml:"ipType,omitempty"`
  // {"en":"1.If let this parameter empty,it will query all rigion 2.Optional values cn, am, apac, au, emea, hk, tw, etc", "zh_CN":"1. 不传查全部
  // 2. 可选值 cn、am、apac、au、emea、hk、tw 等"}
  Region []*string `json:"region,omitempty" xml:"region,omitempty" type:"Repeated"`
}

func (s QueryMultidomainsIPV6OrIPV4RequestsRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryMultidomainsIPV6OrIPV4RequestsRequest) GoString() string {
  return s.String()
}

func (s *QueryMultidomainsIPV6OrIPV4RequestsRequest) SetDateFrom(v string) *QueryMultidomainsIPV6OrIPV4RequestsRequest {
  s.DateFrom = &v
  return s
}

func (s *QueryMultidomainsIPV6OrIPV4RequestsRequest) SetDateTo(v string) *QueryMultidomainsIPV6OrIPV4RequestsRequest {
  s.DateTo = &v
  return s
}

func (s *QueryMultidomainsIPV6OrIPV4RequestsRequest) SetDomain(v []*string) *QueryMultidomainsIPV6OrIPV4RequestsRequest {
  s.Domain = v
  return s
}

func (s *QueryMultidomainsIPV6OrIPV4RequestsRequest) SetDataInterval(v string) *QueryMultidomainsIPV6OrIPV4RequestsRequest {
  s.DataInterval = &v
  return s
}

func (s *QueryMultidomainsIPV6OrIPV4RequestsRequest) SetIpType(v string) *QueryMultidomainsIPV6OrIPV4RequestsRequest {
  s.IpType = &v
  return s
}

func (s *QueryMultidomainsIPV6OrIPV4RequestsRequest) SetRegion(v []*string) *QueryMultidomainsIPV6OrIPV4RequestsRequest {
  s.Region = v
  return s
}

type QueryMultidomainsIPV6OrIPV4RequestsResponse struct {
  // {"en":"request result status code", "zh_CN":"请求结果状态码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"request result information", "zh_CN":"请求结果信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  QueryMultidomainsIPV6OrIPV4RequestsData []*QueryMultidomainsIPV6OrIPV4RequestsData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s QueryMultidomainsIPV6OrIPV4RequestsResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryMultidomainsIPV6OrIPV4RequestsResponse) GoString() string {
  return s.String()
}

func (s *QueryMultidomainsIPV6OrIPV4RequestsResponse) SetCode(v string) *QueryMultidomainsIPV6OrIPV4RequestsResponse {
  s.Code = &v
  return s
}

func (s *QueryMultidomainsIPV6OrIPV4RequestsResponse) SetMessage(v string) *QueryMultidomainsIPV6OrIPV4RequestsResponse {
  s.Message = &v
  return s
}

func (s *QueryMultidomainsIPV6OrIPV4RequestsResponse) SetData(v []*QueryMultidomainsIPV6OrIPV4RequestsData) *QueryMultidomainsIPV6OrIPV4RequestsResponse {
  s.QueryMultidomainsIPV6OrIPV4RequestsData = v
  return s
}

type QueryMultidomainsIPV6OrIPV4RequestsData struct {
  // {"en":"Domain", "zh_CN":"域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  DetailList []*QueryMultidomainsIPV6OrIPV4RequestsDataDetailList `json:"detailList,omitempty" xml:"detailList,omitempty" require:"true" type:"Repeated"`
}

func (s QueryMultidomainsIPV6OrIPV4RequestsData) String() string {
  return tea.Prettify(s)
}

func (s QueryMultidomainsIPV6OrIPV4RequestsData) GoString() string {
  return s.String()
}

func (s *QueryMultidomainsIPV6OrIPV4RequestsData) SetDomain(v string) *QueryMultidomainsIPV6OrIPV4RequestsData {
  s.Domain = &v
  return s
}

func (s *QueryMultidomainsIPV6OrIPV4RequestsData) SetDetailList(v []*QueryMultidomainsIPV6OrIPV4RequestsDataDetailList) *QueryMultidomainsIPV6OrIPV4RequestsData {
  s.DetailList = v
  return s
}

type QueryMultidomainsIPV6OrIPV4RequestsDataDetailList struct     {
  // {"en":"Timestamp, The timestamp between the start time and the end time.", "zh_CN":"时间片
  // 返回开始时间和结束时间包含的时间片。"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {"en":"Number of requests", "zh_CN":"请求个数,单位 个"}
  Value *string `json:"value,omitempty" xml:"value,omitempty" require:"true"`
}

func (s QueryMultidomainsIPV6OrIPV4RequestsDataDetailList) String() string {
  return tea.Prettify(s)
}

func (s QueryMultidomainsIPV6OrIPV4RequestsDataDetailList) GoString() string {
  return s.String()
}

func (s *QueryMultidomainsIPV6OrIPV4RequestsDataDetailList) SetTimestamp(v string) *QueryMultidomainsIPV6OrIPV4RequestsDataDetailList {
  s.Timestamp = &v
  return s
}

func (s *QueryMultidomainsIPV6OrIPV4RequestsDataDetailList) SetValue(v string) *QueryMultidomainsIPV6OrIPV4RequestsDataDetailList {
  s.Value = &v
  return s
}

type QueryMultidomainsIPV6OrIPV4RequestsPaths struct {
}

func (s QueryMultidomainsIPV6OrIPV4RequestsPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryMultidomainsIPV6OrIPV4RequestsPaths) GoString() string {
  return s.String()
}

type QueryMultidomainsIPV6OrIPV4RequestsParameters struct {
}

func (s QueryMultidomainsIPV6OrIPV4RequestsParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryMultidomainsIPV6OrIPV4RequestsParameters) GoString() string {
  return s.String()
}

type QueryMultidomainsIPV6OrIPV4RequestsRequestHeader struct {
}

func (s QueryMultidomainsIPV6OrIPV4RequestsRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryMultidomainsIPV6OrIPV4RequestsRequestHeader) GoString() string {
  return s.String()
}

type QueryMultidomainsIPV6OrIPV4RequestsResponseHeader struct {
}

func (s QueryMultidomainsIPV6OrIPV4RequestsResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryMultidomainsIPV6OrIPV4RequestsResponseHeader) GoString() string {
  return s.String()
}




type QueryDomainTotalRequestRequest struct {
  // {"en":"Domain listDomain number limits can be adjusted depending on different accounts. The default value is 1000(if you want to adjust,please, contact technical support)","zh_CN":"域名列表1.域名个数限制根据账号可调，默认为1000个（可联系技术支持下单调整）；"}
  DomainList *QueryDomainTotalRequestRequestDomainList `json:"domain-list,omitempty" xml:"domain-list,omitempty" require:"true" type:"Struct"`
}

func (s QueryDomainTotalRequestRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryDomainTotalRequestRequest) GoString() string {
  return s.String()
}

func (s *QueryDomainTotalRequestRequest) SetDomainList(v *QueryDomainTotalRequestRequestDomainList) *QueryDomainTotalRequestRequest {
  s.DomainList = v
  return s
}

type QueryDomainTotalRequestRequestDomainList struct {
  // {"en":"Domain","zh_CN":"域名"}
  DomainName []*string `json:"domain-name,omitempty" xml:"domain-name,omitempty" type:"Repeated"`
}

func (s QueryDomainTotalRequestRequestDomainList) String() string {
  return tea.Prettify(s)
}

func (s QueryDomainTotalRequestRequestDomainList) GoString() string {
  return s.String()
}

func (s *QueryDomainTotalRequestRequestDomainList) SetDomainName(v []*string) *QueryDomainTotalRequestRequestDomainList {
  s.DomainName = v
  return s
}

type QueryDomainTotalRequestRequestHeader struct {
}

func (s QueryDomainTotalRequestRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryDomainTotalRequestRequestHeader) GoString() string {
  return s.String()
}

type QueryDomainTotalRequestPaths struct {
}

func (s QueryDomainTotalRequestPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryDomainTotalRequestPaths) GoString() string {
  return s.String()
}

type QueryDomainTotalRequestParameters struct {
  // {"en":"Start time\n1.The format is yyyy-MM-ddTHH:mm:ss+08:00;\n2.And smaller than the current time and dateTo;\n3.Period between dataFrom and dateTo should not be longer than 31 days;","zh_CN":"开始时间\n1.格式为yyyy-MM-ddTHH:mm:ss+08:00；\n2.并且小于当前时间和dateTo；\n3.dateFrom和dateTo相差不能超过31天；\n4.只能查询最近2年内数据。"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty" require:"true"`
  // {"en":"End time\n1.The format is yyyy-MM-ddTHH:mm:ss+08:00\n2.Must be greater than dateFrom;\n3.If it is greater than the current time, then the current time will be assigned as the value;","zh_CN":"结束时间\n1.格式为yyyy-MM-ddTHH:mm:ss+08:00；\n2.必须大于dateFrom；\n3.如果大于当前时间，则重新赋值为当前时间；"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty" require:"true"`
  // {"en":"Data granularity\n1.fiveminutes: five minutes, hourly: one hour, daily: one day;\n2.If not specified, daily is set as the default value;\n3.If fiveminutes is specified as the value, then data is returned in the granularity of actual configuration when there is specific configuration of the data collecting granularity for the customer","zh_CN":"数据粒度\n1.fiveminutes：5分钟，hourly：1小时，daily：1天；\n2.不传递，默认为daily；\n3.传递fiveminutes时，若客户数据采集粒度有特殊配置将按实际配置粒度返回。"}
  Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s QueryDomainTotalRequestParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryDomainTotalRequestParameters) GoString() string {
  return s.String()
}

func (s *QueryDomainTotalRequestParameters) SetDateFrom(v string) *QueryDomainTotalRequestParameters {
  s.DateFrom = &v
  return s
}

func (s *QueryDomainTotalRequestParameters) SetDateTo(v string) *QueryDomainTotalRequestParameters {
  s.DateTo = &v
  return s
}

func (s *QueryDomainTotalRequestParameters) SetType(v string) *QueryDomainTotalRequestParameters {
  s.Type = &v
  return s
}

type QueryDomainTotalRequestResponse struct {
  // {"en":"Total requests","zh_CN":"总请求数"}
  HitSummary *int `json:"hit-summary,omitempty" xml:"hit-summary,omitempty" require:"true"`
  // {"en":"hitData","zh_CN":"请求数据"}
  HitData []*QueryDomainTotalRequestResponseHitData `json:"hit-data,omitempty" xml:"hit-data,omitempty" require:"true" type:"Repeated"`
}

func (s QueryDomainTotalRequestResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryDomainTotalRequestResponse) GoString() string {
  return s.String()
}

func (s *QueryDomainTotalRequestResponse) SetHitSummary(v int) *QueryDomainTotalRequestResponse {
  s.HitSummary = &v
  return s
}

func (s *QueryDomainTotalRequestResponse) SetHitData(v []*QueryDomainTotalRequestResponseHitData) *QueryDomainTotalRequestResponse {
  s.HitData = v
  return s
}

type QueryDomainTotalRequestResponseHitData struct     {
  // {"en":"Date\nWhen the querying data granularity is fiveminutes, the format is yyyy-MM-dd HH:mm; the data value of every time slice represents the data value within the previous time granularity range. The first time slice of the day is yyyy-MM-dd 00:05 AM, and the last one is yyyy-MM-dd 24:00;When the data query granularity is hourly, the format is yyyy-MM-dd HH; the data value of every time slice represents the data value within the previous time granularity range. The first time slice of the day is yyyy-MM-dd 00:01, and the last one is yyyy-MM-dd 24;When the querying data granularity is daily, the format is yyyy-MM-dd; the data value of every time slice represents the value of the data;Return the time slices contained in start time and in end time","zh_CN":"时间\n1.查询的数据粒度为fiveminutes时，格式为yyyy-MM-dd HH:mm；每一个时间片数据值代表的是前一个时间粒度范围内的数据值。一天开始的时间片是yyyy-MM-dd 00:05，最后一个时间片是yyyy-MM-dd 24:00。\n2.查询的数据粒度为hourly时，格式为yyyy-MM-dd HH；每一个时间片数据值代表的是前一个时间粒度范围内的数据值。一天开始的时间片是yyyy-MM-dd 01，最后一个时间片是yyyy-MM-dd 24\n3.查询的数据粒度为daily时，格式为yyyy-MM-dd；每一个时间片数据值代表的该天内的数据值；\n4.返回开始时间和结束时间包含的时间片。"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {"en":"Total number of requests, More than 6 digits are displayed in scientific notation,e.g.:1642565=1.642565E6","zh_CN":"请求数,超过6位数的以科学计数展示，例：1642565=1.642565E6"}
  Hit *int `json:"hit,omitempty" xml:"hit,omitempty" require:"true"`
}

func (s QueryDomainTotalRequestResponseHitData) String() string {
  return tea.Prettify(s)
}

func (s QueryDomainTotalRequestResponseHitData) GoString() string {
  return s.String()
}

func (s *QueryDomainTotalRequestResponseHitData) SetTimestamp(v string) *QueryDomainTotalRequestResponseHitData {
  s.Timestamp = &v
  return s
}

func (s *QueryDomainTotalRequestResponseHitData) SetHit(v int) *QueryDomainTotalRequestResponseHitData {
  s.Hit = &v
  return s
}

type QueryDomainTotalRequestResponseHeader struct {
}

func (s QueryDomainTotalRequestResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryDomainTotalRequestResponseHeader) GoString() string {
  return s.String()
}




type QueryRequestHitRatioRequest struct {
  // {"en":"From date:
  //         1.The time format is yyyy-MM-ddTHH:mm:ss±HH:mm. Please note: ±HH:mm is the time zone offset, which can be adjusted according to your data needs, for example, +00:00 represents UTC time, +08:00 represents East 8th District, and -05:00 represents West 5th District. 2024-01-15T10:30:45+00:00 means UTC time January 15, 2024 10:30:45 AM
  //         2.Cannot exceed current time
  //         3.The most recent six-month (183 days) data are available.", "zh_CN":"开始时间:
  //         1.时间格式为 yyyy-MM-ddTHH:mm:ss±HH:mm。请注意：±HH:mm 为时区偏移量，可根据您的数据需要进行调整，例如 +00:00 代表 UTC 时间，+08:00 代表东八区，-05:00 代表西五区。2024-01-15T10:30:45+00:00，表示UTC 时间 2024 年 1 月 15 日上午 10 点 30 分 45 秒；
  //         2.不能大于当前时间
  //         3.最多可获取最近半年(183天)的数据。"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty"`
  // {"en":"To time:
  //         1.The time format is yyyy-MM-ddTHH:mm:ss±HH:mm. Please note: ±HH:mm is the time zone offset, which can be adjusted according to your data needs, for example, +00:00 represents UTC time, +08:00 represents East 8th District, and -05:00 represents West 5th District. 2024-01-15T10:30:45+00:00 means UTC time January 15, 2024 10:30:45 AM
  //         2.The end time needs to be greater than the start time. If the end time is greater than the current time, take the current time.
  //         3.Date from, Date to both, the default query past 24 hours; If there is only one unsent, throw an exception
  //         4.Maximum allowed query time interval: 31 days, Date from and dateTo, not more than 31 days", "zh_CN":"结束时间:
  //         1.时间格式为 yyyy-MM-ddTHH:mm:ss±HH:mm。请注意：±HH:mm 为时区偏移量，可根据您的数据需要进行调整，例如 +00:00 代表 UTC 时间，+08:00 代表东八区，-05:00 代表西五区。2024-01-15T10:30:45+00:00，表示UTC 时间 2024 年 1 月 15 日上午 10 点 30 分 45 秒
  //         2.结束时间需大于开始时间,结束时间如果大于当前时间,取当前时间。
  //         3.dateFrom,dateTo二者都未传,默认查询过去的24小时；如仅有一个未传,抛异常
  //         4.允许查询最大时间间隔:31天,即dateFrom和dateTo相差不能超过31天"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty"`
  // {"en":"Domain:
  //         1.The maximum number of deliverable domain names is 200 by default
  // 		2.Automatically filter out illegal domain names (pass illegal domain names, will be filtered out, the query results only return the data of legitimate domain names)
  //         3.The default query accounts for all domains if the number of domain names exceeds the upper limit when the entry is not delivered. If the number of domain names in the account exceeds the limit, an error is raised.", "zh_CN":"域名:
  //         1.可传递域名数量上限默认为200个
  //         2.自动过滤掉非法域名(如传递非法域名,会被过滤掉,查询结果只返回合法域名的数据)
  //         3.未传递该入参时,默认查询账号下所有域名,但当账号下域名数量超过上限时提示错误。"}
  Domain []*string `json:"domain,omitempty" xml:"domain,omitempty" require:"true" type:"Repeated"`
  // {"en":"Data granularity:
  // 
  // 1m: 1 minute granularity.
  // 
  // 5m: 5 minute granularity.
  // 
  // 1h: 1 hour granularity.
  // 
  // 1d: 1 day granularity. Default value is 1d.", "zh_CN":"数据粒度：
  // 
  // 支持1m（1分钟）、5m（5分钟）、1h（1小时）、1d（天）
  // 不传默认1d。"}
  DataInterval *string `json:"dataInterval,omitempty" xml:"dataInterval,omitempty"`
}

func (s QueryRequestHitRatioRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryRequestHitRatioRequest) GoString() string {
  return s.String()
}

func (s *QueryRequestHitRatioRequest) SetDateFrom(v string) *QueryRequestHitRatioRequest {
  s.DateFrom = &v
  return s
}

func (s *QueryRequestHitRatioRequest) SetDateTo(v string) *QueryRequestHitRatioRequest {
  s.DateTo = &v
  return s
}

func (s *QueryRequestHitRatioRequest) SetDomain(v []*string) *QueryRequestHitRatioRequest {
  s.Domain = v
  return s
}

func (s *QueryRequestHitRatioRequest) SetDataInterval(v string) *QueryRequestHitRatioRequest {
  s.DataInterval = &v
  return s
}

type QueryRequestHitRatioResponse struct {
  // {"en":"request result status code", "zh_CN":"请求结果状态码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Request result information", "zh_CN":"请求结果信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Detailed data on the results of the request", "zh_CN":"请求结果的详细数据"}
  Data []*QueryRequestHitRatioResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s QueryRequestHitRatioResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryRequestHitRatioResponse) GoString() string {
  return s.String()
}

func (s *QueryRequestHitRatioResponse) SetCode(v string) *QueryRequestHitRatioResponse {
  s.Code = &v
  return s
}

func (s *QueryRequestHitRatioResponse) SetMessage(v string) *QueryRequestHitRatioResponse {
  s.Message = &v
  return s
}

func (s *QueryRequestHitRatioResponse) SetData(v []*QueryRequestHitRatioResponseData) *QueryRequestHitRatioResponse {
  s.Data = v
  return s
}

type QueryRequestHitRatioResponseData struct     {
  // {"en":"Actually processed time. yyyy-MM-dd HH:mm format", "zh_CN":"实际查询时间,格式 yyyy-MM-dd HH:mm"}
  RealDate *string `json:"realDate,omitempty" xml:"realDate,omitempty" require:"true"`
  // {"en":"Average of total hit ratio.", "zh_CN":"总命中率的平均值"}
  TotalAvg *float64 `json:"totalAvg,omitempty" xml:"totalAvg,omitempty" require:"true"`
  // {"en":"hitRatioDatas", "zh_CN":"缓存命中率数据"}
  HitRatioDatas []*QueryRequestHitRatioResponseDataHitRatioDatas `json:"hitRatioDatas,omitempty" xml:"hitRatioDatas,omitempty" require:"true" type:"Repeated"`
}

func (s QueryRequestHitRatioResponseData) String() string {
  return tea.Prettify(s)
}

func (s QueryRequestHitRatioResponseData) GoString() string {
  return s.String()
}

func (s *QueryRequestHitRatioResponseData) SetRealDate(v string) *QueryRequestHitRatioResponseData {
  s.RealDate = &v
  return s
}

func (s *QueryRequestHitRatioResponseData) SetTotalAvg(v float64) *QueryRequestHitRatioResponseData {
  s.TotalAvg = &v
  return s
}

func (s *QueryRequestHitRatioResponseData) SetHitRatioDatas(v []*QueryRequestHitRatioResponseDataHitRatioDatas) *QueryRequestHitRatioResponseData {
  s.HitRatioDatas = v
  return s
}

type QueryRequestHitRatioResponseDataHitRatioDatas struct     {
  // {"en":"timetamp
  // 1. When the data granularity of the query is fiveminutes, the format is yyyy-MM-dd HH:MM; Each time slice data value represents the data value in the previous time granularity range, For example yyyy-MM-dd 00:05 represents data in the range from 00:00 to 00:05.
  // 2.The data granularity of query is hourly, the format is yyyy-MM-dd HH. Each time slice data value represents data values in the previous time granularity range such as yyyy-MM-dd 01 that represent data from 00 to 01.
  // 3. the data granularity of the query is daily, the format is yyyy-MM-dd; Each time slice data value represents the data value for that day.
  // 4.Returns the timetamp contained in start time and end time.", "zh_CN":"时间，
  // 
  // 查询的数据粒度为1m时，格式为yyyy-MM-dd HH:mm；每一个时间片数据值代表的是前一个时间粒度范围内的数据值。一天开始的时间片是yyyy-MM-dd 00:01，最后一个时间片是（yyyy-MM-dd+1） 00:00；
  // 查询的数据粒度为5m时，格式为yyyy-MM-dd HH:mm；每一个时间片数据值代表的是前一个时间粒度范围内的数据值。一天开始的时间片是yyyy-MM-dd 00:05，最后一个时间片是（yyyy-MM-dd+1） 00:00；
  // 查询的数据粒度为1h时，格式为yyyy-MM-dd HH；每一个时间片数据值代表的是前一个时间粒度范围内的数据值。一天开始的时间片是yyyy-MM-dd 01，最后一个时间片是（yyyy-MM-dd+1） 00；
  // 返回开始时间和结束时间包含的时间片。"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {"en":"Cache hit ratio,keep 4 decimal places", "zh_CN":"缓存命中率,保留4位小数"}
  HitRatio *string `json:"hitRatio,omitempty" xml:"hitRatio,omitempty" require:"true"`
}

func (s QueryRequestHitRatioResponseDataHitRatioDatas) String() string {
  return tea.Prettify(s)
}

func (s QueryRequestHitRatioResponseDataHitRatioDatas) GoString() string {
  return s.String()
}

func (s *QueryRequestHitRatioResponseDataHitRatioDatas) SetTimestamp(v string) *QueryRequestHitRatioResponseDataHitRatioDatas {
  s.Timestamp = &v
  return s
}

func (s *QueryRequestHitRatioResponseDataHitRatioDatas) SetHitRatio(v string) *QueryRequestHitRatioResponseDataHitRatioDatas {
  s.HitRatio = &v
  return s
}

type QueryRequestHitRatioPaths struct {
}

func (s QueryRequestHitRatioPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryRequestHitRatioPaths) GoString() string {
  return s.String()
}

type QueryRequestHitRatioParameters struct {
}

func (s QueryRequestHitRatioParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryRequestHitRatioParameters) GoString() string {
  return s.String()
}

type QueryRequestHitRatioRequestHeader struct {
}

func (s QueryRequestHitRatioRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryRequestHitRatioRequestHeader) GoString() string {
  return s.String()
}

type QueryRequestHitRatioResponseHeader struct {
}

func (s QueryRequestHitRatioResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryRequestHitRatioResponseHeader) GoString() string {
  return s.String()
}




type QueryRequestBySpecificProtocolRequest struct {
  // {"en":"Start time:
  // 1.The time format is yyyy-MM-ddTHH:mm:ss±HH:mm. Please note: ±HH:mm is the time zone offset, which can be adjusted according to your data needs, for example, +00:00 represents UTC time, +08:00 represents East 8th District, and -05:00 represents West 5th District. 2024-01-15T10:30:45+00:00 means UTC time January 15, 2024 10:30:45 AM;
  // 2.Must be a time that is 183 days earlier than the current time, and the time must be earlier than the current time and dateTo;
  // 3.Period between dataFrom and dateTo cannot be longer than 7 days(technical support can be contacted to adjust);
  // 4.dateFrom and dateTo can be either both are specified or neither is specifies;
  // 5.If neither dateFrom nor dateTo is specified, then by default, data in the last 24 hour is queried", "zh_CN":"开始时间:
  // 1.时间格式为 yyyy-MM-ddTHH:mm:ss±HH:mm。请注意：±HH:mm 为时区偏移量，可根据您的数据需要进行调整，例如 +00:00 代表 UTC 时间，+08:00 代表东八区，-05:00 代表西五区。2024-01-15T10:30:45+00:00，表示UTC 时间 2024 年 1 月 15 日上午 10 点 30 分 45 秒；
  // 2.必须大于当前时间-183天，并且小于当前时间和dateTo；
  // 3.dateFrom和dateTo相差不能超过7天（可联系技术支持调整）；
  // 4.dateFrom和dateTo要么都传递，要么都不传递；
  // 5.dateFrom和dateTo都未传递，则默认查询过去24小时的数据"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty"`
  // {"en":"End time
  // 1.The time format is yyyy-MM-ddTHH:mm:ss±HH:mm. Please note: ±HH:mm is the time zone offset, which can be adjusted according to your data needs, for example, +00:00 represents UTC time, +08:00 represents East 8th District, and -05:00 represents West 5th District. 2024-01-15T10:30:45+00:00 means UTC time January 15, 2024 10:30:45 AM;
  // 2.Must be greater than dateFrom;
  // 3.If it's greater than the current time, then the current time is assigned as the value;", "zh_CN":"结束时间
  // 1.时间格式为 yyyy-MM-ddTHH:mm:ss±HH:mm。请注意：±HH:mm 为时区偏移量，可根据您的数据需要进行调整，例如 +00:00 代表 UTC 时间，+08:00 代表东八区，-05:00 代表西五区。2024-01-15T10:30:45+00:00，表示UTC 时间 2024 年 1 月 15 日上午 10 点 30 分 45 秒；
  // 2.必须大于dateFrom；
  // 3.如果大于当前时间，则重新赋值为当前时间；"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty"`
  // {"en":"Domain names, domain number limits can be adjusted depending on different accounts. The default value is 20", "zh_CN":"域名，域名个数限制根据账号可调，默认为20个"}
  Domain []*string `json:"domain,omitempty" xml:"domain,omitempty" type:"Repeated"`
  // {"en":"Data granularity: 1m , 5m, support 5m", "zh_CN":"数据粒度 : 1m  , 5m ，默认5m"}
  DataInterval *string `json:"dataInterval,omitempty" xml:"dataInterval,omitempty"`
  // {"en":"Transmission protocol
  // 
  // 1.Options: http, https;
  // 2.https is used as the default value is no value specified;
  // 3.httpRequestData is displayed if http is queried, and httpsRequestData is displayed if https is queried;", "zh_CN":"传输协议
  // 1.可选值为http、https；
  // 2.不传默认查询https；
  // 3.查询http时出参展示httpRequestData，查询https时出参展示httpsRequestData；"}
  ProtocolType *string `json:"protocolType,omitempty" xml:"protocolType,omitempty"`
  // {"en":"Group dimension
  // 
  // 1.The value can be selected is domain;
  // 2.The data is displayed according to the specified dimension;", "zh_CN":"分组维度
  // 1.可选值为domain；
  // 2.有传入则按照该维度展示明细数据；"}
  GroupBy []*string `json:"groupBy,omitempty" xml:"groupBy,omitempty" type:"Repeated"`
}

func (s QueryRequestBySpecificProtocolRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryRequestBySpecificProtocolRequest) GoString() string {
  return s.String()
}

func (s *QueryRequestBySpecificProtocolRequest) SetDateFrom(v string) *QueryRequestBySpecificProtocolRequest {
  s.DateFrom = &v
  return s
}

func (s *QueryRequestBySpecificProtocolRequest) SetDateTo(v string) *QueryRequestBySpecificProtocolRequest {
  s.DateTo = &v
  return s
}

func (s *QueryRequestBySpecificProtocolRequest) SetDomain(v []*string) *QueryRequestBySpecificProtocolRequest {
  s.Domain = v
  return s
}

func (s *QueryRequestBySpecificProtocolRequest) SetDataInterval(v string) *QueryRequestBySpecificProtocolRequest {
  s.DataInterval = &v
  return s
}

func (s *QueryRequestBySpecificProtocolRequest) SetProtocolType(v string) *QueryRequestBySpecificProtocolRequest {
  s.ProtocolType = &v
  return s
}

func (s *QueryRequestBySpecificProtocolRequest) SetGroupBy(v []*string) *QueryRequestBySpecificProtocolRequest {
  s.GroupBy = v
  return s
}

type QueryRequestBySpecificProtocolResponse struct {
  Result []*QueryRequestBySpecificProtocolResponseResult `json:"result,omitempty" xml:"result,omitempty" require:"true" type:"Repeated"`
}

func (s QueryRequestBySpecificProtocolResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryRequestBySpecificProtocolResponse) GoString() string {
  return s.String()
}

func (s *QueryRequestBySpecificProtocolResponse) SetResult(v []*QueryRequestBySpecificProtocolResponseResult) *QueryRequestBySpecificProtocolResponse {
  s.Result = v
  return s
}

type QueryRequestBySpecificProtocolResponseResult struct     {
  // {"en":"Domain", "zh_CN":"域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  HttpsRequestData []*QueryRequestBySpecificProtocolResponseResultHttpsRequestData `json:"httpsRequestData,omitempty" xml:"httpsRequestData,omitempty" require:"true" type:"Repeated"`
}

func (s QueryRequestBySpecificProtocolResponseResult) String() string {
  return tea.Prettify(s)
}

func (s QueryRequestBySpecificProtocolResponseResult) GoString() string {
  return s.String()
}

func (s *QueryRequestBySpecificProtocolResponseResult) SetDomain(v string) *QueryRequestBySpecificProtocolResponseResult {
  s.Domain = &v
  return s
}

func (s *QueryRequestBySpecificProtocolResponseResult) SetHttpsRequestData(v []*QueryRequestBySpecificProtocolResponseResultHttpsRequestData) *QueryRequestBySpecificProtocolResponseResult {
  s.HttpsRequestData = v
  return s
}

type QueryRequestBySpecificProtocolResponseResultHttpsRequestData struct     {
  // {"en":"DateTime, the format is  yyyy-MM-dd HH:mm; the data value of every time slice represents the data  value within the previous time granularity range.", "zh_CN":"时间，格式为yyyy-MM-dd HH:mm；每一个时间片数据值代表的是前一个时间粒度范围内的数据值。比如yyyy-MM-dd 00:05，代表00:00到00:05范围内的数据。"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {"en":"Total number of requests", "zh_CN":"请求数"}
  Value *string `json:"value,omitempty" xml:"value,omitempty" require:"true"`
}

func (s QueryRequestBySpecificProtocolResponseResultHttpsRequestData) String() string {
  return tea.Prettify(s)
}

func (s QueryRequestBySpecificProtocolResponseResultHttpsRequestData) GoString() string {
  return s.String()
}

func (s *QueryRequestBySpecificProtocolResponseResultHttpsRequestData) SetTimestamp(v string) *QueryRequestBySpecificProtocolResponseResultHttpsRequestData {
  s.Timestamp = &v
  return s
}

func (s *QueryRequestBySpecificProtocolResponseResultHttpsRequestData) SetValue(v string) *QueryRequestBySpecificProtocolResponseResultHttpsRequestData {
  s.Value = &v
  return s
}

type QueryRequestBySpecificProtocolPaths struct {
}

func (s QueryRequestBySpecificProtocolPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryRequestBySpecificProtocolPaths) GoString() string {
  return s.String()
}

type QueryRequestBySpecificProtocolParameters struct {
}

func (s QueryRequestBySpecificProtocolParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryRequestBySpecificProtocolParameters) GoString() string {
  return s.String()
}

type QueryRequestBySpecificProtocolRequestHeader struct {
}

func (s QueryRequestBySpecificProtocolRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryRequestBySpecificProtocolRequestHeader) GoString() string {
  return s.String()
}

type QueryRequestBySpecificProtocolResponseHeader struct {
}

func (s QueryRequestBySpecificProtocolResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryRequestBySpecificProtocolResponseHeader) GoString() string {
  return s.String()
}




type QueryIPV6RequestOfeachISPandProvinceRequest struct {
  // {"en":"Start time:\n1. Time format is yyyy-MM-ddTHH:mm:ss+08:00, for example, 2016-12-02T10:00:00+08:00 (10:00 on 2nd of December 2016, Beijing Time);\n2. No bigger than the current time;\n3. Data in the last 183 days at most can be queried.","zh_CN":"开始时间：\n1.时间格式为yyyy-MM-ddTHH:mm:ss+08:00，例如，2016-12-02T10:00:00+08:00（为北京时间2016年12月2日10点0分0秒）;\n2.不能大于当前时间;\n3.最多可获取最近半年（183天）的数据。"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty"`
  // {"en":"End time: 1. the time format is 2016-12-02T10:00:00+08:00;\n2.End time should be greater than start time. If the end time is greater than current time, current time will be used;\n3.If both fields of dataFrom and dateTo are left empty, then data in the last 24 hours will be queried by default; if only one field is filled in and one is left empty, then exception will be occur;\n4.Allowable maximum time range for query: 1 day, means the period between dateFrom to dateTo should not exceed 1 day (can be adjusted by contacting technical support).","zh_CN":"结束时间：\n1.时间格式yyyy-MM-ddTHH:mm:ss+08:00;\n2.结束时间需大于开始时间，结束时间如果大于当前时间，取当前时间;\n3.dateFrom，dateTo二者都未传，默认查询过去的24小时，如仅有一个未传，抛异常;\n4.允许查询最大时间间隔：1天，即dateFrom和dateTo相差不能超过1天。（可联系技术支持调整）"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty"`
  // {"en":"Domain name: \n1. The maximum number of TLDs is 20 by default (Technical Support Adjustment can be contacted);\n2. automatic filtering invalid domain name (if pass illegal domain name, can be filtered, query result only returns the data of valid domain name).","zh_CN":"域名：\n1.可传递域名数量上限默认为200个（可联系技术支持调整）;\n2.自动过滤掉无效域名（如传递非法域名，会被过滤掉，查询结果只返回有效域名的数据）。"}
  Domain []*string `json:"domain,omitempty" xml:"domain,omitempty" type:"Repeated"`
  // {"defaultValue":"5m","en":"Data granularity \n1. Support 5m (5 minute granularity), 1h (1 hour granularity); 2. The default is 5m;","zh_CN":"数据粒度：\n1.支持5m（5分钟）、1h（1小时）\n2.不传默认5m。","exampleValue":"5m,1h"}
  DataInterval *string `json:"dataInterval,omitempty" xml:"dataInterval,omitempty"`
  // {"en":"Province\n\n1.Province is not upload: Query all provinces and aggregate the returned data according to all provinces;\n2.Province is upload: Provinces can transmit Chinese or code. Please refer to the appendix description section of the overview page for the provincial information code table.\n\n3.Support language request header Accept Language, only support zh-CN and en-US, default to zh-CN. Accept Language: en-US, both the province and isp input and return are in code, otherwise the return is in Chinese.","zh_CN":"省份\n\n1.未传递province时：查询所有省份，返回的数据按照所有省份聚合。\n\n2.有传递province时：省份 可传中文或code。省份信息码表详见概览页附录说明章节\n\n3.支持语言请求头Accept-Language，只支持zh-CN、en-US，默认为zh-CN。Accept-Language：en-US时，省份及运营商 入参及返回都为code，否则返回的为中文。"}
  Province []*string `json:"province,omitempty" xml:"province,omitempty" type:"Repeated"`
  // {"en":"ISP:\n1.ISP is not upload: Query all ISPs and aggregate the returned data according to all ISPs;\n2.ISP is upload: Isp can transmit Chinese or code. Please refer to the appendix description section of the overview page for the ISP information code table.","zh_CN":"运营商：\n1.未传递isp时：查询所有isp，返回的数据按照所有运营商聚合。\n2.有传递isp时：运营商 可传中文或code。运营商信息码表详见概览页附录说明章节"}
  Isp []*string `json:"isp,omitempty" xml:"isp,omitempty" type:"Repeated"`
  // {"en":"IP type:\n1.The optional values are IPv6 and IPv4;\n2.If let this parameter empty,it will query all IP type.","zh_CN":"IP类型：\n1.可选值为 IPV6、IPV4;\n2.不传默认查询全部","exampleValue":"IPV6,IPV4"}
  IPType *string `json:"IPType,omitempty" xml:"IPType,omitempty"`
  // {"en":"Grouped dimension:\n1.Aggregation date by default;\n2.the optional value is domain,province,isp,allow to send multi option ;\n3.send the Grouped dimension represent the need to display details by their corresponding values.For example, when groupBy is isp, the ISP dimension needs to be displayed in detail. When an ISP is not passed, it represents an aggregate date and would not return the ISP node.Provinces and domains have the same logic. For example, by passing 'groupBy': ['domain', 'province'], the ISP node under ispData does not need to return. {domain:'www.aaaa.com','ispData': [{'isp','China Telecom','provinceData': [...]}}","zh_CN":"分组关键词：\n1.默认聚合展示;\n2.可选值为domain.province.isp，可传入多个值;\n3.传入关键词则代表需要按照关键词对应的值展示明细; 例如groupBy传入isp，则isp维度需要明细展示;当没有传递isp，则代表isp聚合展示，同时isp节点则不返回。其他province和domain相同逻辑。 例如：传递'groupBy':   ['domain','province']，则ispData下的isp节点无需返回。 { 'domain': 'www.aaaa.com', 'ispData': [ { 'isp':   '中国电信', 'provinceData': [....] }]}"}
  GroupBy []*string `json:"groupBy,omitempty" xml:"groupBy,omitempty" type:"Repeated"`
}

func (s QueryIPV6RequestOfeachISPandProvinceRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryIPV6RequestOfeachISPandProvinceRequest) GoString() string {
  return s.String()
}

func (s *QueryIPV6RequestOfeachISPandProvinceRequest) SetDateFrom(v string) *QueryIPV6RequestOfeachISPandProvinceRequest {
  s.DateFrom = &v
  return s
}

func (s *QueryIPV6RequestOfeachISPandProvinceRequest) SetDateTo(v string) *QueryIPV6RequestOfeachISPandProvinceRequest {
  s.DateTo = &v
  return s
}

func (s *QueryIPV6RequestOfeachISPandProvinceRequest) SetDomain(v []*string) *QueryIPV6RequestOfeachISPandProvinceRequest {
  s.Domain = v
  return s
}

func (s *QueryIPV6RequestOfeachISPandProvinceRequest) SetDataInterval(v string) *QueryIPV6RequestOfeachISPandProvinceRequest {
  s.DataInterval = &v
  return s
}

func (s *QueryIPV6RequestOfeachISPandProvinceRequest) SetProvince(v []*string) *QueryIPV6RequestOfeachISPandProvinceRequest {
  s.Province = v
  return s
}

func (s *QueryIPV6RequestOfeachISPandProvinceRequest) SetIsp(v []*string) *QueryIPV6RequestOfeachISPandProvinceRequest {
  s.Isp = v
  return s
}

func (s *QueryIPV6RequestOfeachISPandProvinceRequest) SetIPType(v string) *QueryIPV6RequestOfeachISPandProvinceRequest {
  s.IPType = &v
  return s
}

func (s *QueryIPV6RequestOfeachISPandProvinceRequest) SetGroupBy(v []*string) *QueryIPV6RequestOfeachISPandProvinceRequest {
  s.GroupBy = v
  return s
}

type QueryIPV6RequestOfeachISPandProvinceRequestHeader struct {
}

func (s QueryIPV6RequestOfeachISPandProvinceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryIPV6RequestOfeachISPandProvinceRequestHeader) GoString() string {
  return s.String()
}

type QueryIPV6RequestOfeachISPandProvincePaths struct {
}

func (s QueryIPV6RequestOfeachISPandProvincePaths) String() string {
  return tea.Prettify(s)
}

func (s QueryIPV6RequestOfeachISPandProvincePaths) GoString() string {
  return s.String()
}

type QueryIPV6RequestOfeachISPandProvinceParameters struct {
}

func (s QueryIPV6RequestOfeachISPandProvinceParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryIPV6RequestOfeachISPandProvinceParameters) GoString() string {
  return s.String()
}

type QueryIPV6RequestOfeachISPandProvinceResponse struct {
  // {"en":"","zh_CN":""}
  Result []*QueryIPV6RequestOfeachISPandProvinceResponseResult `json:"result,omitempty" xml:"result,omitempty" require:"true" type:"Repeated"`
}

func (s QueryIPV6RequestOfeachISPandProvinceResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryIPV6RequestOfeachISPandProvinceResponse) GoString() string {
  return s.String()
}

func (s *QueryIPV6RequestOfeachISPandProvinceResponse) SetResult(v []*QueryIPV6RequestOfeachISPandProvinceResponseResult) *QueryIPV6RequestOfeachISPandProvinceResponse {
  s.Result = v
  return s
}

type QueryIPV6RequestOfeachISPandProvinceResponseResult struct     {
  // {"en":"domain","zh_CN":"域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {"en":"","zh_CN":""}
  IspData []*QueryIPV6RequestOfeachISPandProvinceResponseResultIspData `json:"ispData,omitempty" xml:"ispData,omitempty" require:"true" type:"Repeated"`
}

func (s QueryIPV6RequestOfeachISPandProvinceResponseResult) String() string {
  return tea.Prettify(s)
}

func (s QueryIPV6RequestOfeachISPandProvinceResponseResult) GoString() string {
  return s.String()
}

func (s *QueryIPV6RequestOfeachISPandProvinceResponseResult) SetDomain(v string) *QueryIPV6RequestOfeachISPandProvinceResponseResult {
  s.Domain = &v
  return s
}

func (s *QueryIPV6RequestOfeachISPandProvinceResponseResult) SetIspData(v []*QueryIPV6RequestOfeachISPandProvinceResponseResultIspData) *QueryIPV6RequestOfeachISPandProvinceResponseResult {
  s.IspData = v
  return s
}

type QueryIPV6RequestOfeachISPandProvinceResponseResultIspData struct     {
  // {"en":"isp","zh_CN":"运营商"}
  Isp *string `json:"isp,omitempty" xml:"isp,omitempty" require:"true"`
  // {"en":"","zh_CN":""}
  ProvinceData []*QueryIPV6RequestOfeachISPandProvinceResponseResultIspDataProvinceData `json:"provinceData,omitempty" xml:"provinceData,omitempty" require:"true" type:"Repeated"`
}

func (s QueryIPV6RequestOfeachISPandProvinceResponseResultIspData) String() string {
  return tea.Prettify(s)
}

func (s QueryIPV6RequestOfeachISPandProvinceResponseResultIspData) GoString() string {
  return s.String()
}

func (s *QueryIPV6RequestOfeachISPandProvinceResponseResultIspData) SetIsp(v string) *QueryIPV6RequestOfeachISPandProvinceResponseResultIspData {
  s.Isp = &v
  return s
}

func (s *QueryIPV6RequestOfeachISPandProvinceResponseResultIspData) SetProvinceData(v []*QueryIPV6RequestOfeachISPandProvinceResponseResultIspDataProvinceData) *QueryIPV6RequestOfeachISPandProvinceResponseResultIspData {
  s.ProvinceData = v
  return s
}

type QueryIPV6RequestOfeachISPandProvinceResponseResultIspDataProvinceData struct     {
  // {"en":"province","zh_CN":"省份"}
  Province *string `json:"province,omitempty" xml:"province,omitempty" require:"true"`
  // {"en":"","zh_CN":""}
  RequestData []*QueryIPV6RequestOfeachISPandProvinceResponseResultIspDataProvinceDataRequestData `json:"requestData,omitempty" xml:"requestData,omitempty" require:"true" type:"Repeated"`
}

func (s QueryIPV6RequestOfeachISPandProvinceResponseResultIspDataProvinceData) String() string {
  return tea.Prettify(s)
}

func (s QueryIPV6RequestOfeachISPandProvinceResponseResultIspDataProvinceData) GoString() string {
  return s.String()
}

func (s *QueryIPV6RequestOfeachISPandProvinceResponseResultIspDataProvinceData) SetProvince(v string) *QueryIPV6RequestOfeachISPandProvinceResponseResultIspDataProvinceData {
  s.Province = &v
  return s
}

func (s *QueryIPV6RequestOfeachISPandProvinceResponseResultIspDataProvinceData) SetRequestData(v []*QueryIPV6RequestOfeachISPandProvinceResponseResultIspDataProvinceDataRequestData) *QueryIPV6RequestOfeachISPandProvinceResponseResultIspDataProvinceData {
  s.RequestData = v
  return s
}

type QueryIPV6RequestOfeachISPandProvinceResponseResultIspDataProvinceDataRequestData struct     {
  // {"en":"1.When the data query granularity is 5m, then the format is yyyy-MM-dd HH:mm (the data value of every time slice represents the data value within the previous time granularity range. The first time slice of the day is yyyy-MM-dd 12:05 AM, and the last one is (yyyy-MM-dd+1) 00:00 );\n2.When the data query granularity is hourly, the format is yyyy-MM-dd HH; the data value of every time slice represents the data value within the previous time granularity range. The first time slice of the day is yyyy-MM-dd 01, and the last one is yyyy-MM-dd 24;\n3.Return the time slices that contained in start time and in end time.","zh_CN":"时间，\n1.查询的数据粒度为5m时，格式为yyyy-MM-dd HH:mm；每一个时间片数据值代表的是前一个时间粒度范围内的数据值。一天开始的时间片是yyyy-MM-dd 00:05，最后一个时间片是（yyyy-MM-dd+1） 00:00；\n2.查询的数据粒度为1h时，格式为yyyy-MM-dd HH；每一个时间片数据值代表的是前一个时间粒度范围内的数据值。一天开始的时间片是yyyy-MM-dd 01，最后一个时间片是（yyyy-MM-dd+1） 00；\n3.返回开始时间和结束时间包含的时间片。"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {"en":"Requests","zh_CN":"请求数"}
  Value *string `json:"value,omitempty" xml:"value,omitempty" require:"true"`
}

func (s QueryIPV6RequestOfeachISPandProvinceResponseResultIspDataProvinceDataRequestData) String() string {
  return tea.Prettify(s)
}

func (s QueryIPV6RequestOfeachISPandProvinceResponseResultIspDataProvinceDataRequestData) GoString() string {
  return s.String()
}

func (s *QueryIPV6RequestOfeachISPandProvinceResponseResultIspDataProvinceDataRequestData) SetTimestamp(v string) *QueryIPV6RequestOfeachISPandProvinceResponseResultIspDataProvinceDataRequestData {
  s.Timestamp = &v
  return s
}

func (s *QueryIPV6RequestOfeachISPandProvinceResponseResultIspDataProvinceDataRequestData) SetValue(v string) *QueryIPV6RequestOfeachISPandProvinceResponseResultIspDataProvinceDataRequestData {
  s.Value = &v
  return s
}

type QueryIPV6RequestOfeachISPandProvinceResponseHeader struct {
}

func (s QueryIPV6RequestOfeachISPandProvinceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryIPV6RequestOfeachISPandProvinceResponseHeader) GoString() string {
  return s.String()
}




type QueryRequestNumbersUnderShieldPoPRequest struct {
  // {"en":"Start time:
  // 
  // 1.The format is yyyy-MM-ddTHH:mm:ss+08:00;
  // 2.Must be a time that is 183 days earlier than the current time, and the time must be earlier than the current time and dateTo;
  // 3.Period between dataFrom and dateTo cannot be longer than 7 days(technical support can be contacted to adjust);
  // 4.dateFrom and dateTo can be either both are specified or neither is specifies;
  // 5.If neither dateFrom nor dateTo is specified, then by default, data in the last 24 hour is queried", "zh_CN":"开始时间:
  // 1.格式为yyyy-MM-ddTHH:mm:ss+08:00;
  // 2.必须大于当前时间-183天,并且小于当前时间和dateTo;
  // 3.dateFrom和dateTo相差不能超过7天(可联系技术支持调整);
  // 4.dateFrom和dateTo要么都传递,要么都不传递;
  // 5.dateFrom和dateTo都未传递,则默认查询过去24小时的数据"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty"`
  // {"en":"End time:
  // 
  // 1.The format is yyyy-MM-ddTHH:mm:ss+08:00;
  // 2.Must be greater than dateFrom;
  // 3.If it's greater than the current time, then the current time is assigned as the value;", "zh_CN":"结束时间:
  // 1.格式为yyyy-MM-ddTHH:mm:ss+08:00;
  // 2.必须大于dateFrom;
  // 3.如果大于当前时间,则重新赋值为当前时间;"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty"`
  // {"en":"Domain names, domain number limits can be adjusted depending on different accounts. The default value is 20", "zh_CN":"域名,域名个数限制根据账号可调,默认为20个"}
  Domain []*string `json:"domain,omitempty" xml:"domain,omitempty" require:"true" type:"Repeated"`
  // {"en":"Data granularity, support 5m (granularity of 5 minutes)", "zh_CN":"数据粒度,支持5m: 5分钟粒度"}
  DataInterval *string `json:"dataInterval,omitempty" xml:"dataInterval,omitempty"`
  // {"en":"Group dimension
  // 
  // 1.The value can be selected is domain;
  // 2.The data is displayed according to the specified dimension;", "zh_CN":"分组维度
  // 1.可选值为domain;
  // 2.有传入则按照该维度展示明细数据;"}
  GroupBy []*string `json:"groupBy,omitempty" xml:"groupBy,omitempty" type:"Repeated"`
}

func (s QueryRequestNumbersUnderShieldPoPRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryRequestNumbersUnderShieldPoPRequest) GoString() string {
  return s.String()
}

func (s *QueryRequestNumbersUnderShieldPoPRequest) SetDateFrom(v string) *QueryRequestNumbersUnderShieldPoPRequest {
  s.DateFrom = &v
  return s
}

func (s *QueryRequestNumbersUnderShieldPoPRequest) SetDateTo(v string) *QueryRequestNumbersUnderShieldPoPRequest {
  s.DateTo = &v
  return s
}

func (s *QueryRequestNumbersUnderShieldPoPRequest) SetDomain(v []*string) *QueryRequestNumbersUnderShieldPoPRequest {
  s.Domain = v
  return s
}

func (s *QueryRequestNumbersUnderShieldPoPRequest) SetDataInterval(v string) *QueryRequestNumbersUnderShieldPoPRequest {
  s.DataInterval = &v
  return s
}

func (s *QueryRequestNumbersUnderShieldPoPRequest) SetGroupBy(v []*string) *QueryRequestNumbersUnderShieldPoPRequest {
  s.GroupBy = v
  return s
}

type QueryRequestNumbersUnderShieldPoPResponse struct {
  // {"en":"result", "zh_CN":"结果"}
  Result []*QueryRequestNumbersUnderShieldPoPResponseResult `json:"result,omitempty" xml:"result,omitempty" require:"true" type:"Repeated"`
}

func (s QueryRequestNumbersUnderShieldPoPResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryRequestNumbersUnderShieldPoPResponse) GoString() string {
  return s.String()
}

func (s *QueryRequestNumbersUnderShieldPoPResponse) SetResult(v []*QueryRequestNumbersUnderShieldPoPResponseResult) *QueryRequestNumbersUnderShieldPoPResponse {
  s.Result = v
  return s
}

type QueryRequestNumbersUnderShieldPoPResponseResult struct     {
  // {"en":"Domain", "zh_CN":"域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {"en":"Peak Time", "zh_CN":"峰值时间"}
  PeakTime *string `json:"peakTime,omitempty" xml:"peakTime,omitempty" require:"true"`
  // {"en":"Peak Request", "zh_CN":"请求数峰值"}
  PeakValue *string `json:"peakValue,omitempty" xml:"peakValue,omitempty" require:"true"`
  // {"en":"Total Request", "zh_CN":"总请求数"}
  TotalRequest *string `json:"totalRequest,omitempty" xml:"totalRequest,omitempty" require:"true"`
  RequestData []*QueryRequestNumbersUnderShieldPoPResponseResultRequestData `json:"requestData,omitempty" xml:"requestData,omitempty" require:"true" type:"Repeated"`
}

func (s QueryRequestNumbersUnderShieldPoPResponseResult) String() string {
  return tea.Prettify(s)
}

func (s QueryRequestNumbersUnderShieldPoPResponseResult) GoString() string {
  return s.String()
}

func (s *QueryRequestNumbersUnderShieldPoPResponseResult) SetDomain(v string) *QueryRequestNumbersUnderShieldPoPResponseResult {
  s.Domain = &v
  return s
}

func (s *QueryRequestNumbersUnderShieldPoPResponseResult) SetPeakTime(v string) *QueryRequestNumbersUnderShieldPoPResponseResult {
  s.PeakTime = &v
  return s
}

func (s *QueryRequestNumbersUnderShieldPoPResponseResult) SetPeakValue(v string) *QueryRequestNumbersUnderShieldPoPResponseResult {
  s.PeakValue = &v
  return s
}

func (s *QueryRequestNumbersUnderShieldPoPResponseResult) SetTotalRequest(v string) *QueryRequestNumbersUnderShieldPoPResponseResult {
  s.TotalRequest = &v
  return s
}

func (s *QueryRequestNumbersUnderShieldPoPResponseResult) SetRequestData(v []*QueryRequestNumbersUnderShieldPoPResponseResultRequestData) *QueryRequestNumbersUnderShieldPoPResponseResult {
  s.RequestData = v
  return s
}

type QueryRequestNumbersUnderShieldPoPResponseResultRequestData struct     {
  // {"en":"DateTime, the format is yyyy-MM-dd HH:mm; the data value of every time slice represents the data  value within the previous time granularity range.", "zh_CN":"时间,格式为yyyy-MM-dd HH:mm;每一个时间片数据值代表的是前一个时间粒度范围内的数据值。比如yyyy-MM-dd 00:05,代表00:00到00:05范围内的数据。"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {"en":"Total number of requests", "zh_CN":"请求数"}
  Value *string `json:"value,omitempty" xml:"value,omitempty" require:"true"`
}

func (s QueryRequestNumbersUnderShieldPoPResponseResultRequestData) String() string {
  return tea.Prettify(s)
}

func (s QueryRequestNumbersUnderShieldPoPResponseResultRequestData) GoString() string {
  return s.String()
}

func (s *QueryRequestNumbersUnderShieldPoPResponseResultRequestData) SetTimestamp(v string) *QueryRequestNumbersUnderShieldPoPResponseResultRequestData {
  s.Timestamp = &v
  return s
}

func (s *QueryRequestNumbersUnderShieldPoPResponseResultRequestData) SetValue(v string) *QueryRequestNumbersUnderShieldPoPResponseResultRequestData {
  s.Value = &v
  return s
}

type QueryRequestNumbersUnderShieldPoPPaths struct {
}

func (s QueryRequestNumbersUnderShieldPoPPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryRequestNumbersUnderShieldPoPPaths) GoString() string {
  return s.String()
}

type QueryRequestNumbersUnderShieldPoPParameters struct {
}

func (s QueryRequestNumbersUnderShieldPoPParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryRequestNumbersUnderShieldPoPParameters) GoString() string {
  return s.String()
}

type QueryRequestNumbersUnderShieldPoPRequestHeader struct {
}

func (s QueryRequestNumbersUnderShieldPoPRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryRequestNumbersUnderShieldPoPRequestHeader) GoString() string {
  return s.String()
}

type QueryRequestNumbersUnderShieldPoPResponseHeader struct {
}

func (s QueryRequestNumbersUnderShieldPoPResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryRequestNumbersUnderShieldPoPResponseHeader) GoString() string {
  return s.String()
}




type QueryEdgeRequestHitRatioServiceRequest struct {
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

func (s QueryEdgeRequestHitRatioServiceRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryEdgeRequestHitRatioServiceRequest) GoString() string {
  return s.String()
}

func (s *QueryEdgeRequestHitRatioServiceRequest) SetDateFrom(v string) *QueryEdgeRequestHitRatioServiceRequest {
  s.DateFrom = &v
  return s
}

func (s *QueryEdgeRequestHitRatioServiceRequest) SetDateTo(v string) *QueryEdgeRequestHitRatioServiceRequest {
  s.DateTo = &v
  return s
}

func (s *QueryEdgeRequestHitRatioServiceRequest) SetDomain(v []*string) *QueryEdgeRequestHitRatioServiceRequest {
  s.Domain = v
  return s
}

type QueryEdgeRequestHitRatioServiceResponse struct {
  Result []*QueryEdgeRequestHitRatioServiceResponseResult `json:"result,omitempty" xml:"result,omitempty" require:"true" type:"Repeated"`
}

func (s QueryEdgeRequestHitRatioServiceResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryEdgeRequestHitRatioServiceResponse) GoString() string {
  return s.String()
}

func (s *QueryEdgeRequestHitRatioServiceResponse) SetResult(v []*QueryEdgeRequestHitRatioServiceResponseResult) *QueryEdgeRequestHitRatioServiceResponse {
  s.Result = v
  return s
}

type QueryEdgeRequestHitRatioServiceResponseResult struct     {
  // {"en":"Actually processed time.  yyyy-MM-dd HH:mm format", "zh_CN":"实际查询时间，格式 yyyy-MM-dd HH:mm"}
  RealDate *string `json:"realDate,omitempty" xml:"realDate,omitempty" require:"true"`
  // {"en":"Average of total edge node hit ratio", "zh_CN":"总边缘节点命中率的平均值,2位小数"}
  TotalAvg *string `json:"totalAvg,omitempty" xml:"totalAvg,omitempty" require:"true"`
  // {"en":"Detailed data on the results of the request", "zh_CN":"请求结果的详细数据"}
  HitRatioDatas []*QueryEdgeRequestHitRatioServiceResponseResultHitRatioDatas `json:"hitRatioDatas,omitempty" xml:"hitRatioDatas,omitempty" require:"true" type:"Repeated"`
}

func (s QueryEdgeRequestHitRatioServiceResponseResult) String() string {
  return tea.Prettify(s)
}

func (s QueryEdgeRequestHitRatioServiceResponseResult) GoString() string {
  return s.String()
}

func (s *QueryEdgeRequestHitRatioServiceResponseResult) SetRealDate(v string) *QueryEdgeRequestHitRatioServiceResponseResult {
  s.RealDate = &v
  return s
}

func (s *QueryEdgeRequestHitRatioServiceResponseResult) SetTotalAvg(v string) *QueryEdgeRequestHitRatioServiceResponseResult {
  s.TotalAvg = &v
  return s
}

func (s *QueryEdgeRequestHitRatioServiceResponseResult) SetHitRatioDatas(v []*QueryEdgeRequestHitRatioServiceResponseResultHitRatioDatas) *QueryEdgeRequestHitRatioServiceResponseResult {
  s.HitRatioDatas = v
  return s
}

type QueryEdgeRequestHitRatioServiceResponseResultHitRatioDatas struct     {
  // {"en":"timestamp", "zh_CN":"时间，格式为yyyy-MM-dd HH:mm；每一个时间片数据值代表的是前一个时间粒度范围内的数据值。一天开始的时间片是yyyy-MM-dd 00:01，最后一个时间片是（yyyy-MM-dd+1） 00:00；"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {"en":"edge node hit ratio,keep 4 decimal places", "zh_CN":"边缘节点请求数命中率，保留4位小数"}
  EdgeHitRatio *string `json:"edgeHitRatio,omitempty" xml:"edgeHitRatio,omitempty" require:"true"`
}

func (s QueryEdgeRequestHitRatioServiceResponseResultHitRatioDatas) String() string {
  return tea.Prettify(s)
}

func (s QueryEdgeRequestHitRatioServiceResponseResultHitRatioDatas) GoString() string {
  return s.String()
}

func (s *QueryEdgeRequestHitRatioServiceResponseResultHitRatioDatas) SetTimestamp(v string) *QueryEdgeRequestHitRatioServiceResponseResultHitRatioDatas {
  s.Timestamp = &v
  return s
}

func (s *QueryEdgeRequestHitRatioServiceResponseResultHitRatioDatas) SetEdgeHitRatio(v string) *QueryEdgeRequestHitRatioServiceResponseResultHitRatioDatas {
  s.EdgeHitRatio = &v
  return s
}

type QueryEdgeRequestHitRatioServicePaths struct {
}

func (s QueryEdgeRequestHitRatioServicePaths) String() string {
  return tea.Prettify(s)
}

func (s QueryEdgeRequestHitRatioServicePaths) GoString() string {
  return s.String()
}

type QueryEdgeRequestHitRatioServiceParameters struct {
}

func (s QueryEdgeRequestHitRatioServiceParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryEdgeRequestHitRatioServiceParameters) GoString() string {
  return s.String()
}

type QueryEdgeRequestHitRatioServiceRequestHeader struct {
}

func (s QueryEdgeRequestHitRatioServiceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryEdgeRequestHitRatioServiceRequestHeader) GoString() string {
  return s.String()
}

type QueryEdgeRequestHitRatioServiceResponseHeader struct {
}

func (s QueryEdgeRequestHitRatioServiceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryEdgeRequestHitRatioServiceResponseHeader) GoString() string {
  return s.String()
}




