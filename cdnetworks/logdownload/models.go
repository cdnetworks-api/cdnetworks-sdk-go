package logdownload

// This file is auto-generated, don't edit it. Thanks.
import (
  "github.com/alibabacloud-go/tea/tea"
)

type QueryDomainLogDownloadAddressRequest struct {
  // {"en":"Domain list:
  //     1.The default upper limit of the number of domain names is 20 (you can contact technical support for adjustment), The maximum recommended cap is 500 .", "zh_CN":"域名列表:
  //     1.域名数量默认上限为20个(可联系技术支持调整), 最大上限为500."}
  QueryDomainLogDownloadAddressDomainList *QueryDomainLogDownloadAddressDomainList `json:"domain-list,omitempty" xml:"domain-list,omitempty"`
  // {"en":"Start time:
  //     1.The format is yyyy-MM-ddTHH:mm:ss+08:00 ;
  //     2.Must be smaller than the current time and dateTo ;
  //     3.Period between dataFrom and dateTo cannot be longer than 31 days .", "zh_CN":"开始时间:
  //     1.格式为yyyy-MM-ddTHH:mm:ss+08:00 ;
  //     2.必须小于当前时间和dateTo ;
  //     3.dateFrom和dateTo相差不能超过31天(可联系技术支持调整) ;
  //     4.只能查询最近2年内数据.(实际可查询的日志范围,取决于域名配置的日志保留天数) ."}
  Datefrom *string `json:"datefrom,omitempty" xml:"datefrom,omitempty" require:"true"`
  // {"en":"End time, format is yyyy-MM-ddTHH:mm:ss+08:00(The actual range of logs that can be queried depends on the number of days of log retention configured by the domain name).", "zh_CN":"结束时间,格式为yyyy-MM-ddTHH:mm:ss+08:00(实际可查询的日志范围,取决于域名配置的日志保留天数)."}
  Dateto *string `json:"dateto,omitempty" xml:"dateto,omitempty" require:"true"`
  // {"en":"Log type, optional values:cdn,bot,ddos;Multiple are separated by English commas. If they are not transmitted, the data will be queried by logtype = cdn,bot by default.", "zh_CN":"日志类型,可选值: cdn,bot,ddos ;
  // 多个通过英文逗号分隔,若未传则默认按logtype=cdn,bot查询数据"}
  LogType *string `json:"logType,omitempty" xml:"logType,omitempty"`
  // {"en":"areaCode", "zh_CN":"加速区域"}
  AreaCode *string `json:"areaCode,omitempty" xml:"areaCode,omitempty"`
}

func (s QueryDomainLogDownloadAddressRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryDomainLogDownloadAddressRequest) GoString() string {
  return s.String()
}

func (s *QueryDomainLogDownloadAddressRequest) SetDomainList(v *QueryDomainLogDownloadAddressDomainList) *QueryDomainLogDownloadAddressRequest {
  s.QueryDomainLogDownloadAddressDomainList = v
  return s
}

func (s *QueryDomainLogDownloadAddressRequest) SetDatefrom(v string) *QueryDomainLogDownloadAddressRequest {
  s.Datefrom = &v
  return s
}

func (s *QueryDomainLogDownloadAddressRequest) SetDateto(v string) *QueryDomainLogDownloadAddressRequest {
  s.Dateto = &v
  return s
}

func (s *QueryDomainLogDownloadAddressRequest) SetLogType(v string) *QueryDomainLogDownloadAddressRequest {
  s.LogType = &v
  return s
}

func (s *QueryDomainLogDownloadAddressRequest) SetAreaCode(v string) *QueryDomainLogDownloadAddressRequest {
  s.AreaCode = &v
  return s
}

type QueryDomainLogDownloadAddressDomainList struct {
  // {"en":"Domain", "zh_CN":"域名"}
  DomainName []*string `json:"domain-name,omitempty" xml:"domain-name,omitempty" require:"true" type:"Repeated"`
}

func (s QueryDomainLogDownloadAddressDomainList) String() string {
  return tea.Prettify(s)
}

func (s QueryDomainLogDownloadAddressDomainList) GoString() string {
  return s.String()
}

func (s *QueryDomainLogDownloadAddressDomainList) SetDomainName(v []*string) *QueryDomainLogDownloadAddressDomainList {
  s.DomainName = v
  return s
}

type QueryDomainLogDownloadAddressResponse struct {
  Logs []*QueryDomainLogDownloadAddressResponseLogs `json:"logs,omitempty" xml:"logs,omitempty" require:"true" type:"Repeated"`
}

func (s QueryDomainLogDownloadAddressResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryDomainLogDownloadAddressResponse) GoString() string {
  return s.String()
}

func (s *QueryDomainLogDownloadAddressResponse) SetLogs(v []*QueryDomainLogDownloadAddressResponseLogs) *QueryDomainLogDownloadAddressResponse {
  s.Logs = v
  return s
}

type QueryDomainLogDownloadAddressResponseLogs struct     {
  // {"en":"Domain", "zh_CN":"域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {"en":"areaCode", "zh_CN":"加速区域"}
  AreaCode *string `json:"areaCode,omitempty" xml:"areaCode,omitempty" require:"true"`
  Files []*QueryDomainLogDownloadAddressResponseLogsFiles `json:"files,omitempty" xml:"files,omitempty" require:"true" type:"Repeated"`
}

func (s QueryDomainLogDownloadAddressResponseLogs) String() string {
  return tea.Prettify(s)
}

func (s QueryDomainLogDownloadAddressResponseLogs) GoString() string {
  return s.String()
}

func (s *QueryDomainLogDownloadAddressResponseLogs) SetDomain(v string) *QueryDomainLogDownloadAddressResponseLogs {
  s.Domain = &v
  return s
}

func (s *QueryDomainLogDownloadAddressResponseLogs) SetAreaCode(v string) *QueryDomainLogDownloadAddressResponseLogs {
  s.AreaCode = &v
  return s
}

func (s *QueryDomainLogDownloadAddressResponseLogs) SetFiles(v []*QueryDomainLogDownloadAddressResponseLogsFiles) *QueryDomainLogDownloadAddressResponseLogs {
  s.Files = v
  return s
}

type QueryDomainLogDownloadAddressResponseLogsFiles struct     {
  // {"en":"The start time of log file, format is yyyy-MM-dd-HHmm", "zh_CN":"日志文件的开始时间,格式为yyyy-MM-dd-HHmm"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty" require:"true"`
  // {"en":"The end time of log file, format is yyyy-MM-dd-HHmm", "zh_CN":"日志文件的结束时间,格式为yyyy-MM-dd-HHmm"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty" require:"true"`
  // {"en":"Download address of log file", "zh_CN":"日志文件下载地址"}
  LogUrl *string `json:"logUrl,omitempty" xml:"logUrl,omitempty" require:"true"`
  // {"en":"Size of log file", "zh_CN":"日志文件大小"}
  FileSize *int32 `json:"fileSize,omitempty" xml:"fileSize,omitempty" require:"true"`
}

func (s QueryDomainLogDownloadAddressResponseLogsFiles) String() string {
  return tea.Prettify(s)
}

func (s QueryDomainLogDownloadAddressResponseLogsFiles) GoString() string {
  return s.String()
}

func (s *QueryDomainLogDownloadAddressResponseLogsFiles) SetDateFrom(v string) *QueryDomainLogDownloadAddressResponseLogsFiles {
  s.DateFrom = &v
  return s
}

func (s *QueryDomainLogDownloadAddressResponseLogsFiles) SetDateTo(v string) *QueryDomainLogDownloadAddressResponseLogsFiles {
  s.DateTo = &v
  return s
}

func (s *QueryDomainLogDownloadAddressResponseLogsFiles) SetLogUrl(v string) *QueryDomainLogDownloadAddressResponseLogsFiles {
  s.LogUrl = &v
  return s
}

func (s *QueryDomainLogDownloadAddressResponseLogsFiles) SetFileSize(v int32) *QueryDomainLogDownloadAddressResponseLogsFiles {
  s.FileSize = &v
  return s
}

type QueryDomainLogDownloadAddressPaths struct {
}

func (s QueryDomainLogDownloadAddressPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryDomainLogDownloadAddressPaths) GoString() string {
  return s.String()
}

type QueryDomainLogDownloadAddressParameters struct {
}

func (s QueryDomainLogDownloadAddressParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryDomainLogDownloadAddressParameters) GoString() string {
  return s.String()
}

type QueryDomainLogDownloadAddressRequestHeader struct {
}

func (s QueryDomainLogDownloadAddressRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryDomainLogDownloadAddressRequestHeader) GoString() string {
  return s.String()
}

type QueryDomainLogDownloadAddressResponseHeader struct {
}

func (s QueryDomainLogDownloadAddressResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryDomainLogDownloadAddressResponseHeader) GoString() string {
  return s.String()
}




