package reportpv

// This file is auto-generated, don't edit it. Thanks.
import (
  "github.com/alibabacloud-go/tea/tea"
)

type QueryDomainUVRequest struct {
  // {"en":"Start date:
  //         1. The time format is yyyy-MM-ddTHH:MM:ss+08:00, for example, 2016-12-02T10:00:00+08:00 (Beijing Time 2 December 2016 10:0 min 0 seconds);
  //         2. Not Greater Than The Current Time
  //         3. Data for the last six months (183 days) are available at most.", "zh_CN":"开始时间：
  //         1.时间格式为yyyy-MM-ddTHH:mm:ss+08:00，例如，2016-12-02T10:00:00+08:00（为北京时间2016年12月2日10点0分0秒）；
  //         2.不能大于当前时间
  //         3.最多可获取最近半年（183天）的数据。"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty"`
  // {"en":"End time:
  //         1. Time format yyyy-MM-ddTHH:MM:ss+08:00
  //         2. The end time must be greater than the start time. if the end time is greater than the current time, take the current time.
  //         3. Date from, dateTo, neither passed, default query past 24 hours; If there is only one unsent, throw an exception
  //         4. Maximum time interval allowed for queries: 31 days, i.e. the difference between Date from and dateTo cannot exceed 31 days. (Could contact technical support adjustment)", "zh_CN":"结束时间：
  //         1.时间格式yyyy-MM-ddTHH:mm:ss+08:00
  //         2.结束时间需大于开始时间，结束时间如果大于当前时间，取当前时间。
  //         3.dateFrom，dateTo二者都未传，默认查询过去的24小时；如仅有一个未传，抛异常
  //         4.允许查询最大时间间隔：31天，即dateFrom和dateTo相差不能超过31天。（可联系技术支持调整）"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty"`
  // {"en":"Data granularity
  //         1. Support 1h (1 hour granularity), 1d (1 day granularity);
  //         2. The default is 1h;", "zh_CN":"数据粒度
  //         1.支持1h（1小时粒度）、1d（1天粒度）；
  //         2.不传默认为1h；"}
  DataInterval *string `json:"dataInterval,omitempty" xml:"dataInterval,omitempty"`
  // {"en":"Domain name:
  //         1. The maximum number of TLDs is 20 by default (Technical Support Adjustment can be contacted);
  //         2. Automatic filtering invalid domain name (if pass illegal domain name, can be filtered, query result only returns the data of valid domain name).
  //         3. No default lookup of all domain names", "zh_CN":"域名：
  //         1.可传递域名数量上限默认为20个（可联系技术支持调整）；
  //         2.自动过滤掉无效域名（如传递非法域名，会被过滤，查询结果只返回有效域名的数据）。
  //         3.不传默认查全部域名"}
  Domain []*string `json:"domain,omitempty" xml:"domain,omitempty" type:"Repeated"`
}

func (s QueryDomainUVRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryDomainUVRequest) GoString() string {
  return s.String()
}

func (s *QueryDomainUVRequest) SetDateFrom(v string) *QueryDomainUVRequest {
  s.DateFrom = &v
  return s
}

func (s *QueryDomainUVRequest) SetDateTo(v string) *QueryDomainUVRequest {
  s.DateTo = &v
  return s
}

func (s *QueryDomainUVRequest) SetDataInterval(v string) *QueryDomainUVRequest {
  s.DataInterval = &v
  return s
}

func (s *QueryDomainUVRequest) SetDomain(v []*string) *QueryDomainUVRequest {
  s.Domain = v
  return s
}

type QueryDomainUVResponse struct {
  // {"en":"result", "zh_CN":"结果"}
  Result []*QueryDomainUVResponseResult `json:"result,omitempty" xml:"result,omitempty" require:"true" type:"Repeated"`
}

func (s QueryDomainUVResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryDomainUVResponse) GoString() string {
  return s.String()
}

func (s *QueryDomainUVResponse) SetResult(v []*QueryDomainUVResponseResult) *QueryDomainUVResponse {
  s.Result = v
  return s
}

type QueryDomainUVResponseResult struct     {
  // {"en":"Domain", "zh_CN":"域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {"en":"Peak Time", "zh_CN":"峰值时间"}
  PeakTime *string `json:"peakTime,omitempty" xml:"peakTime,omitempty" require:"true"`
  // {"en":"Peak value of UV", "zh_CN":"UV峰值"}
  PeakUV *string `json:"peakUV,omitempty" xml:"peakUV,omitempty" require:"true"`
  // {"en":"Total UV", "zh_CN":"UV总数"}
  TotalUV *string `json:"totalUV,omitempty" xml:"totalUV,omitempty" require:"true"`
  // {"en":"uvData", "zh_CN":"UV数"}
  UvData []*QueryDomainUVResponseResultUvData `json:"uvData,omitempty" xml:"uvData,omitempty" require:"true" type:"Repeated"`
}

func (s QueryDomainUVResponseResult) String() string {
  return tea.Prettify(s)
}

func (s QueryDomainUVResponseResult) GoString() string {
  return s.String()
}

func (s *QueryDomainUVResponseResult) SetDomain(v string) *QueryDomainUVResponseResult {
  s.Domain = &v
  return s
}

func (s *QueryDomainUVResponseResult) SetPeakTime(v string) *QueryDomainUVResponseResult {
  s.PeakTime = &v
  return s
}

func (s *QueryDomainUVResponseResult) SetPeakUV(v string) *QueryDomainUVResponseResult {
  s.PeakUV = &v
  return s
}

func (s *QueryDomainUVResponseResult) SetTotalUV(v string) *QueryDomainUVResponseResult {
  s.TotalUV = &v
  return s
}

func (s *QueryDomainUVResponseResult) SetUvData(v []*QueryDomainUVResponseResultUvData) *QueryDomainUVResponseResult {
  s.UvData = v
  return s
}

type QueryDomainUVResponseResultUvData struct     {
  // {"en":"time
  //         1. When the data granularity of the query is 1h, the format is yyyy-MM-dd HH; Each time slice data value represents the data value in the previous time granularity range. The time slice at the beginning of the day is yyyy-MM-dd 01, and the last time slice is (yyyy-MM-dd+1) 00;
  //         2. When the data granularity of the query is 1d, the format is yyyy-MM-dd; Each time slice data value represents the value of the data for that date.
  //         3. Return the time slice contained in the start and end times.", "zh_CN":"时间
  //         1.查询的数据粒度为1h时，格式为yyyy-MM-dd HH；每一个时间片数据值代表的是前一个时间粒度范围内的数据值。一天开始的时间片是yyyy-MM-dd 01，最后一个时间片是（yyyy-MM-dd+1） 00；
  //         2.查询的数据粒度为1d时，格式为yyyy-MM-dd；每一个时间片数据值代表的该天内的数据值。
  //         3.返回开始时间和结束时间包含的时间片。"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {"en":"UV", "zh_CN":"UV数"}
  Value *string `json:"value,omitempty" xml:"value,omitempty" require:"true"`
}

func (s QueryDomainUVResponseResultUvData) String() string {
  return tea.Prettify(s)
}

func (s QueryDomainUVResponseResultUvData) GoString() string {
  return s.String()
}

func (s *QueryDomainUVResponseResultUvData) SetTimestamp(v string) *QueryDomainUVResponseResultUvData {
  s.Timestamp = &v
  return s
}

func (s *QueryDomainUVResponseResultUvData) SetValue(v string) *QueryDomainUVResponseResultUvData {
  s.Value = &v
  return s
}

type QueryDomainUVPaths struct {
}

func (s QueryDomainUVPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryDomainUVPaths) GoString() string {
  return s.String()
}

type QueryDomainUVParameters struct {
}

func (s QueryDomainUVParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryDomainUVParameters) GoString() string {
  return s.String()
}

type QueryDomainUVRequestHeader struct {
}

func (s QueryDomainUVRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryDomainUVRequestHeader) GoString() string {
  return s.String()
}

type QueryDomainUVResponseHeader struct {
}

func (s QueryDomainUVResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryDomainUVResponseHeader) GoString() string {
  return s.String()
}




