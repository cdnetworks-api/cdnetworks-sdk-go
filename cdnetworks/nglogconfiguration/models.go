package nglogconfiguration

// This file is auto-generated, don't edit it. Thanks.
import (
  "github.com/alibabacloud-go/tea/tea"
)

type DeleteALogConfigurationPaths struct {
  // {"en" : "ID of a log configuration", "zh_CN": "日志配置的id。"}
  Id *int `json:"id,omitempty" xml:"id,omitempty" require:"true"`
}

func (s DeleteALogConfigurationPaths) String() string {
  return tea.Prettify(s)
}

func (s DeleteALogConfigurationPaths) GoString() string {
  return s.String()
}

func (s *DeleteALogConfigurationPaths) SetId(v int) *DeleteALogConfigurationPaths {
  s.Id = &v
  return s
}

type DeleteALogConfigurationParameters struct {
}

func (s DeleteALogConfigurationParameters) String() string {
  return tea.Prettify(s)
}

func (s DeleteALogConfigurationParameters) GoString() string {
  return s.String()
}

type DeleteALogConfigurationRequestHeader struct {
}

func (s DeleteALogConfigurationRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteALogConfigurationRequestHeader) GoString() string {
  return s.String()
}

type DeleteALogConfigurationRequest struct {
}

func (s DeleteALogConfigurationRequest) String() string {
  return tea.Prettify(s)
}

func (s DeleteALogConfigurationRequest) GoString() string {
  return s.String()
}

type DeleteALogConfigurationResponseHeader struct {
}

func (s DeleteALogConfigurationResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteALogConfigurationResponseHeader) GoString() string {
  return s.String()
}

type DeleteALogConfigurationResponse struct {
}

func (s DeleteALogConfigurationResponse) String() string {
  return tea.Prettify(s)
}

func (s DeleteALogConfigurationResponse) GoString() string {
  return s.String()
}




type GetAListOfLogConfigurationsPaths struct {
}

func (s GetAListOfLogConfigurationsPaths) String() string {
  return tea.Prettify(s)
}

func (s GetAListOfLogConfigurationsPaths) GoString() string {
  return s.String()
}

type GetAListOfLogConfigurationsParameters struct {
  // {"en" : "Search for hostnames affected by the log configuration.", "zh_CN": "根据域名查询相关的日志配置。"}
  Search *string `json:"search,omitempty" xml:"search,omitempty"`
  // {"en" : "Range: >= 0 
  // Index of the first log configuration to return.", "zh_CN": "默认值: 0 取值范围: >= 0 
  // 查询起始位置。"}
  Offset *int `json:"offset,omitempty" xml:"offset,omitempty"`
  // {"en" : "Default: 100 Range: [ 1 .. 200 ] 
  // Maximum number of log configurations to return.", "zh_CN": "默认值: 100 取值范围: <= 200 
  // 每次查询的最大条数。"}
  Limit *int `json:"limit,omitempty" xml:"limit,omitempty"`
}

func (s GetAListOfLogConfigurationsParameters) String() string {
  return tea.Prettify(s)
}

func (s GetAListOfLogConfigurationsParameters) GoString() string {
  return s.String()
}

func (s *GetAListOfLogConfigurationsParameters) SetSearch(v string) *GetAListOfLogConfigurationsParameters {
  s.Search = &v
  return s
}

func (s *GetAListOfLogConfigurationsParameters) SetOffset(v int) *GetAListOfLogConfigurationsParameters {
  s.Offset = &v
  return s
}

func (s *GetAListOfLogConfigurationsParameters) SetLimit(v int) *GetAListOfLogConfigurationsParameters {
  s.Limit = &v
  return s
}

type GetAListOfLogConfigurationsRequestHeader struct {
}

func (s GetAListOfLogConfigurationsRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetAListOfLogConfigurationsRequestHeader) GoString() string {
  return s.String()
}

type GetAListOfLogConfigurationsRequest struct {
}

func (s GetAListOfLogConfigurationsRequest) String() string {
  return tea.Prettify(s)
}

func (s GetAListOfLogConfigurationsRequest) GoString() string {
  return s.String()
}

type GetAListOfLogConfigurationsResponseHeader struct {
}

func (s GetAListOfLogConfigurationsResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetAListOfLogConfigurationsResponseHeader) GoString() string {
  return s.String()
}

type GetAListOfLogConfigurationsResponse struct {
  // {"en" : "List of log configurations.", "zh_CN": "日志配置列表。"}
  LogConfigs []*GetAListOfLogConfigurationsResponseLogConfigs `json:"logConfigs,omitempty" xml:"logConfigs,omitempty" require:"true" type:"Repeated"`
  // {"en" : "Range: >= 0 
  // Number of log configurations.", "zh_CN": "取值范围: >= 0 
  // 日志配置数量。"}
  Count *int `json:"count,omitempty" xml:"count,omitempty" require:"true"`
}

func (s GetAListOfLogConfigurationsResponse) String() string {
  return tea.Prettify(s)
}

func (s GetAListOfLogConfigurationsResponse) GoString() string {
  return s.String()
}

func (s *GetAListOfLogConfigurationsResponse) SetLogConfigs(v []*GetAListOfLogConfigurationsResponseLogConfigs) *GetAListOfLogConfigurationsResponse {
  s.LogConfigs = v
  return s
}

func (s *GetAListOfLogConfigurationsResponse) SetCount(v int) *GetAListOfLogConfigurationsResponse {
  s.Count = &v
  return s
}

type GetAListOfLogConfigurationsResponseLogConfigs struct     {
  // {"en" : "An ID representating the log configuration. Use the ID to get, update, or delete the configuration later.", "zh_CN": "日志配置唯一标识ID。可用该ID查询，更新或删除日志配置。"}
  LogConfigId *int `json:"logConfigId,omitempty" xml:"logConfigId,omitempty"`
  // {"en" : "A list of hostnames to which the log configuration applies. A hostname should be a fully qualified domain name like domain.com, a wildcard domain name with a leading '*' like *.domain.com, or '*' if you want the log configuration to apply to all hostnames.", "zh_CN": "适用该日志配置的域名列表。每个域名必须是完整的FQDN域名如domain.com，或者是带星号的泛域名如*.domain.com。如果该日志配置适用于所有域名，则直接用星号*表示。"}
  Hostnames []*string `json:"hostnames,omitempty" xml:"hostnames,omitempty" type:"Repeated"`
  // {"en" : "A description of the log configuration.", "zh_CN": "日志配置描述信息。"}
  Description *string `json:"description,omitempty" xml:"description,omitempty"`
  // {"en" : "RFC 3339 date indicating when the log configuration was last updated.", "zh_CN": "日志配置最近一次更新时间，以RFC 3339日期格式展示。"}
  LastUpdateTime *string `json:"lastUpdateTime,omitempty" xml:"lastUpdateTime,omitempty"`
  // {"en" : "RFC 3339 date indicating when the log configuration was created.", "zh_CN": "日志配置创建时间，以RFC 3339日期格式展示。"}
  CreationTime *string `json:"creationTime,omitempty" xml:"creationTime,omitempty"`
}

func (s GetAListOfLogConfigurationsResponseLogConfigs) String() string {
  return tea.Prettify(s)
}

func (s GetAListOfLogConfigurationsResponseLogConfigs) GoString() string {
  return s.String()
}

func (s *GetAListOfLogConfigurationsResponseLogConfigs) SetLogConfigId(v int) *GetAListOfLogConfigurationsResponseLogConfigs {
  s.LogConfigId = &v
  return s
}

func (s *GetAListOfLogConfigurationsResponseLogConfigs) SetHostnames(v []*string) *GetAListOfLogConfigurationsResponseLogConfigs {
  s.Hostnames = v
  return s
}

func (s *GetAListOfLogConfigurationsResponseLogConfigs) SetDescription(v string) *GetAListOfLogConfigurationsResponseLogConfigs {
  s.Description = &v
  return s
}

func (s *GetAListOfLogConfigurationsResponseLogConfigs) SetLastUpdateTime(v string) *GetAListOfLogConfigurationsResponseLogConfigs {
  s.LastUpdateTime = &v
  return s
}

func (s *GetAListOfLogConfigurationsResponseLogConfigs) SetCreationTime(v string) *GetAListOfLogConfigurationsResponseLogConfigs {
  s.CreationTime = &v
  return s
}




type CreateALogConfigurationRequest struct {
  // {"en":"A description of the log configuration.","zh_CN":"日志配置描述。"}
  Description *string `json:"description,omitempty" xml:"description,omitempty"`
  // {"en":"A list of hostnames to which the log configuration applies. A hostname should be a fully qualified domain name like domain.com, a wildcard domain name with a leading '*' like.domain.com, or '*' if you want the log configuration to apply to all hostnames. You should ensure that only one log configuration applies to each hostname.","zh_CN":"适用该日志配置的域名列表。每个域名必须是完整的FQDN域名如domain.com，或者是带星号的泛域名如*.domain.com。如果该日志配置适用于所有域名，则直接用星号*表示。每个域名有且只能有一个对应的日志配置。"}
  Hostnames []*string `json:"hostnames,omitempty" xml:"hostnames,omitempty" require:"true" type:"Repeated"`
  // {"en":"Default: %cltip %rmtuser  [%apachet] '%method %url %protocol' %statuscode %rspsize '%referer' '%ua' %rsptime\nFormat of the log entry. You can use the following variables to create a custom format:\n<table><tr><th>Variable</th><th>Meaning</th></tr><tr><td>%apachet</td><td>Time in the Common Log Format, for example, 03/Apr/2023:07:49:03 +0000</td></tr><tr><td>%cachestate</td><td>HIT, MISS, or REVALIDATE</td></tr><tr><td>%cltip</td><td>Client IP address</td></tr><tr><td>%cltisp</td><td>Client ISP</td></tr><tr><td>%cltport</td><td>Client port number</td></tr><tr><td>%cltregion</td><td>Client region</td></tr><tr><td>%cpu_ns</td><td>CPU time in nanoseconds for this request</td></tr><tr><td>%custom1</td><td>Refers to a custom log field with ID 1 you define using the custom_log_field directive.</td></tr>      <tr><td>%custom2</td><td>Refers to a custom log field with ID 2 you define using the custom_log_field directive.</td></tr>      <tr><td>%hostname</td><td>Host header</td></tr><tr><td>%method</td><td>HTTP method used to access the content, for example, GET</td></tr><tr><td>%protocol</td><td>HTTP/1.0 or HTTP/1.1 or HTTP2.0</td></tr><tr><td>%querystr</td><td>Query string</td></tr>      <tr><td>%referer</td><td>Referer request header</td></tr><tr><td>%reqhdrsize</td><td>Request header size in bytes</td></tr><tr><td>%reqrange</td><td>Range header in requests from client</td></tr><tr><td>%reqsize</td><td>Request size in bytes including the request line, header, and request body</td></tr><tr><td>%rmtuser</td><td>User name extracted from the Authorization header when basic authentication is used</td></tr><tr><td>%rspsize</td><td>HTTP response size, in bytes, including header and body, but not including TCP/IP/MAC</td></tr><tr><td>%rsptime</td><td>Response time in milliseconds</td></tr><tr><td>%samplerate</td><td>Number of requests corresponding to each log entry</td></tr><tr><td>%scheme</td><td>http or https</td></tr><tr><td>%statuscode</td><td>Response's status code, for example, 200</td></tr><tr><td>%srvip</td><td>IP address of the CDN server handling the request</td></tr><tr><td>%svrnode</td><td>Cache server node name</td></tr><tr><td>%tcprtt</td><td>Round-trip time between server and client in microseconds</td></tr>      <tr><td>%ua</td><td>User-Agent header</td></tr><tr><td>%uniqueid</td><td>A string uniquely identifying this request</td></tr>       <tr><td>%url</td><td>Full URL with query string, if any</td></tr><tr><td>%utctime</td><td>RFC 3339 timestamp of the response, for example, 2021-10-05T12:34:56Z</td></tr></table>\nIf any variable value contains non-printable characters (ASCII characters <=0x1F or >=0x7F), double quotes, or a backslash, each of those characters will be escaped to a format like '\xXX'. For example, double quotes will become '\x22' and backslash will become '\x5C'.\nNote that space is not escaped. You must quote variables that may contain spaces to facilitate parsing.","zh_CN":"默认值: %cltip %rmtuser  [%apachet] '%method %url %protocol' %statuscode %rspsize '%referer' '%ua' %rsptime\n日志格式。可使用以下变量来自定义日志输出格式：\n<table><tr><th>变量</th><th>变量说明</th></tr><tr><td>%apachet</td><td>CLF（Common Log Format）时间格式，例如03/Apr/2023:07:49:03 +0000。</td></tr><tr><td>%cachestate</td><td>缓存状态，取值为HIT，MISS 或者 REVALIDATE。</td></tr><tr><td>%cltip</td><td>客户端IP地址。</td></tr><tr><td>%cltisp</td><td>客户端所属运营商。</td></tr><tr><td>%cltport</td><td>客户端端口号。</td></tr><tr><td>%cltregion</td><td>客户端所属区域。</td></tr><tr><td>%cpu_ns</td><td>此次请求所消耗的CPU时间，以纳秒计算。</td></tr><tr><td>%custom1</td><td>如果您在加速项目中自定义了日志字段，可以用该变量来获取自定义日志字段1的值。自定义日志字段属于高级功能，如需使用请联系我们的技术支持。</td></tr>      <tr><td>%custom2</td><td>如果您在加速项目中自定义了日志字段，可以用该变量来获取自定义日志字段2的值。自定义日志字段属于高级功能，如需使用请联系我们的技术支持。</td></tr>      <tr><td>%hostname</td><td>Host 请求头。</td></tr><tr><td>%method</td><td>HTTP请求方法，例如 GET。</td></tr><tr><td>%protocol</td><td>HTTP/1.0，HTTP/1.1 或者 HTTP2.0协议。</td></tr><tr><td>%querystr</td><td>查询参数字符串。</td></tr>      <tr><td>%referer</td><td>Referer 请求头。</td></tr><tr><td>%reqhdrsize</td><td>请求头大小，单位为bytes。</td></tr><tr><td>%reqrange</td><td>客户端请求所携带的Range请求头。</td></tr><tr><td>%reqsize</td><td>请求的大小，单位为bytes。包括请求行，请求头和请求体。</td></tr><tr><td>%rmtuser</td><td>从Authorization请求头解析出的用户名。当您使用HTTP basic对用户请求进行鉴权时，可通过该变量获取用户名。</td></tr><tr><td>%rspsize</td><td>HTTP响应的大小，单位为bytes。包括响应头和响应体，但不包括TCP/IP/MAC头部开销。</td></tr><tr><td>%rsptime</td><td>响应时间，单位为毫秒。</td></tr><tr><td>%samplerate</td><td>采样率，即每几次请求输出一条日志。</td></tr><tr><td>%sheme</td><td>http或https协议。</td></tr><tr><td>%statuscode</td><td>响应状态码，例如200。</td></tr><tr><td>%srvip</td><td>处理此次请求的CDN服务器的IP地址。</td></tr><tr><td>%svrnode</td><td>CDN服务器所在的节点名称。</td></tr><tr><td>%tcprtt</td><td>服务器与客户端之间的RTT时长，单位为毫秒。</td></tr>      <tr><td>%ua</td><td>User-Agent 请求头。</td></tr><tr><td>%uniqueid</td><td>此次请求的唯一标识，字符串类型。</td></tr>       <tr><td>%url</td><td>完整的URL，包括查询参数字符串。</td></tr><tr><td>%utctime</td><td>响应时间，以RFC 3339格式表示，例如 2021-10-05T12:34:56Z。</td></tr></table>\n\n变量如果包含了不可打印的字符（ASCII码<=0x1F 或 >=0x7F的字符），双引号或者反斜杆，则这些字符会被转义，用'\xXX'的格式表示。例如，双引号以'\x22'表示，反斜杆以'\x5C'表示。需要注意的是，空格不会被转义。如果变量包含空格，则必须加上双引号，才能被正确解析。"}
  LogDownloadFormat *string `json:"logDownloadFormat,omitempty" xml:"logDownloadFormat,omitempty"`
  // {"en":"Default: 14 Range: [ 1 .. 30 ]\nNumber of days to store the logs.","zh_CN":"默认值: 14 取值范围: [ 1 .. 30 ]\n日志保存天数。"}
  LogDownloadStorageDays *int `json:"logDownloadStorageDays,omitempty" xml:"logDownloadStorageDays,omitempty"`
  // {"en":"Enum: 5,15,30,60,120,240,480,1440\nDefault: 30\nTime span of each log file in minutes. When the count of logs tend to increase drastically, logs within a same time span might be truncated and put into multiple log files. When there is significant delay of logs ingestion, the logs delayed might be written to separate log files.","zh_CN":"取值范围: 5,15,30,60,120,240,480,1440\n默认值: 30\n每几分钟生成日志文件。为避免单个日志文件过大，当日志条目增长过快时，同一个时间段内的日志可能被分割为多个文件。当日志入库发生显著延迟时，延迟入库的部分会被写入单独的日志文件。"}
  LogDownloadFileSpanMinutes *int `json:"logDownloadFileSpanMinutes,omitempty" xml:"logDownloadFileSpanMinutes,omitempty"`
}

func (s CreateALogConfigurationRequest) String() string {
  return tea.Prettify(s)
}

func (s CreateALogConfigurationRequest) GoString() string {
  return s.String()
}

func (s *CreateALogConfigurationRequest) SetDescription(v string) *CreateALogConfigurationRequest {
  s.Description = &v
  return s
}

func (s *CreateALogConfigurationRequest) SetHostnames(v []*string) *CreateALogConfigurationRequest {
  s.Hostnames = v
  return s
}

func (s *CreateALogConfigurationRequest) SetLogDownloadFormat(v string) *CreateALogConfigurationRequest {
  s.LogDownloadFormat = &v
  return s
}

func (s *CreateALogConfigurationRequest) SetLogDownloadStorageDays(v int) *CreateALogConfigurationRequest {
  s.LogDownloadStorageDays = &v
  return s
}

func (s *CreateALogConfigurationRequest) SetLogDownloadFileSpanMinutes(v int) *CreateALogConfigurationRequest {
  s.LogDownloadFileSpanMinutes = &v
  return s
}

type CreateALogConfigurationRequestHeader struct {
}

func (s CreateALogConfigurationRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s CreateALogConfigurationRequestHeader) GoString() string {
  return s.String()
}

type CreateALogConfigurationPaths struct {
}

func (s CreateALogConfigurationPaths) String() string {
  return tea.Prettify(s)
}

func (s CreateALogConfigurationPaths) GoString() string {
  return s.String()
}

type CreateALogConfigurationParameters struct {
}

func (s CreateALogConfigurationParameters) String() string {
  return tea.Prettify(s)
}

func (s CreateALogConfigurationParameters) GoString() string {
  return s.String()
}

type CreateALogConfigurationResponse struct {
}

func (s CreateALogConfigurationResponse) String() string {
  return tea.Prettify(s)
}

func (s CreateALogConfigurationResponse) GoString() string {
  return s.String()
}

type CreateALogConfigurationResponseHeader struct {
  // {"en":"Returns a URL pointing to the new log config created, if the request is accepted. The URL contains the ID of the new log config. </br> URL format: <code>{scheme}://{domain}/cdn/report/logConfigs/{logConfigId}</code> Example URL: <code>https://api.example.com/cdn/report/logConfigs/19472dwecwdd48323</code>","zh_CN":"当接口调用成功时，通过Location响应头返回新建的日志配置的URL。URL中包含日志配置的ID，可使用该ID调用'查询日志配置详情'接口来查看日志配置的详细信息。</br> URL格式：<code>{协议}://{域名}/cdn/report/logConfigs/{日志配置ID}</code> URL示例： <code>https://open.chinanetcenter.com/cdn/report/logConfigs/19472dwecwdd48323</code>"}
  Location *string `json:"Location,omitempty" xml:"Location,omitempty" require:"true"`
}

func (s CreateALogConfigurationResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s CreateALogConfigurationResponseHeader) GoString() string {
  return s.String()
}

func (s *CreateALogConfigurationResponseHeader) SetLocation(v string) *CreateALogConfigurationResponseHeader {
  s.Location = &v
  return s
}




type GetAccessLogsForHostnamesPaths struct {
}

func (s GetAccessLogsForHostnamesPaths) String() string {
  return tea.Prettify(s)
}

func (s GetAccessLogsForHostnamesPaths) GoString() string {
  return s.String()
}

type GetAccessLogsForHostnamesParameters struct {
  // {"en" : "RFC 3339 date indicating the beginning of the time period. It should be within the past 14 days. The time must be specified using the UTC timezone; it cannot be an offset. Example: startdate=2021-10-30T00:00:00Z ", "zh_CN": "查询范围的开始时间，以RFC 3339日期格式表示。必须使用UTC时区指定时间。示例：startdate=2019-10-30T00:00:00Z。 查询开始时间不得早于14天之前。"}
  Startdate *string `json:"startdate,omitempty" xml:"startdate,omitempty" require:"true"`
  // {"en" : "RFC 3339 date indicating the end of the time period. The time must be specified using the UTC timezone; it cannot be an offset. Example: enddate=2021-11-14T00:00:00Z ", "zh_CN": "查询范围的结束时间，以RFC 3339日期格式表示。必须使用UTC时区指定时间。示例：enddate=2019-11-14T00:00:00Z。"}
  Enddate *string `json:"enddate,omitempty" xml:"enddate,omitempty" require:"true"`
}

func (s GetAccessLogsForHostnamesParameters) String() string {
  return tea.Prettify(s)
}

func (s GetAccessLogsForHostnamesParameters) GoString() string {
  return s.String()
}

func (s *GetAccessLogsForHostnamesParameters) SetStartdate(v string) *GetAccessLogsForHostnamesParameters {
  s.Startdate = &v
  return s
}

func (s *GetAccessLogsForHostnamesParameters) SetEnddate(v string) *GetAccessLogsForHostnamesParameters {
  s.Enddate = &v
  return s
}

type GetAccessLogsForHostnamesRequestHeader struct {
}

func (s GetAccessLogsForHostnamesRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetAccessLogsForHostnamesRequestHeader) GoString() string {
  return s.String()
}

type GetAccessLogsForHostnamesRequest struct {
  // {"en" : "", "zh_CN": ""}
  Filters *GetAccessLogsForHostnamesRequestFilters `json:"filters,omitempty" xml:"filters,omitempty" type:"Struct"`
}

func (s GetAccessLogsForHostnamesRequest) String() string {
  return tea.Prettify(s)
}

func (s GetAccessLogsForHostnamesRequest) GoString() string {
  return s.String()
}

func (s *GetAccessLogsForHostnamesRequest) SetFilters(v *GetAccessLogsForHostnamesRequestFilters) *GetAccessLogsForHostnamesRequest {
  s.Filters = v
  return s
}

type GetAccessLogsForHostnamesRequestFilters struct {
  // {"en" : "Specify the hostnames that interest you.", "zh_CN": "指定加速域名。"}
  Hostnames []*string `json:"hostnames,omitempty" xml:"hostnames,omitempty" type:"Repeated"`
}

func (s GetAccessLogsForHostnamesRequestFilters) String() string {
  return tea.Prettify(s)
}

func (s GetAccessLogsForHostnamesRequestFilters) GoString() string {
  return s.String()
}

func (s *GetAccessLogsForHostnamesRequestFilters) SetHostnames(v []*string) *GetAccessLogsForHostnamesRequestFilters {
  s.Hostnames = v
  return s
}

type GetAccessLogsForHostnamesResponseHeader struct {
}

func (s GetAccessLogsForHostnamesResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetAccessLogsForHostnamesResponseHeader) GoString() string {
  return s.String()
}

type GetAccessLogsForHostnamesResponse struct {
  // {"en" : "List of objects describing logs you can download.", "zh_CN": "日志信息列表。"}
  Logs []*GetAccessLogsForHostnamesResponseLogs `json:"logs,omitempty" xml:"logs,omitempty" require:"true" type:"Repeated"`
}

func (s GetAccessLogsForHostnamesResponse) String() string {
  return tea.Prettify(s)
}

func (s GetAccessLogsForHostnamesResponse) GoString() string {
  return s.String()
}

func (s *GetAccessLogsForHostnamesResponse) SetLogs(v []*GetAccessLogsForHostnamesResponseLogs) *GetAccessLogsForHostnamesResponse {
  s.Logs = v
  return s
}

type GetAccessLogsForHostnamesResponseLogs struct     {
  // {"en" : "An RFC 3339 date indicate the beginning of the log file.", "zh_CN": "RFC 3339格式的日期，表示日志文件的开始时间。"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty"`
  // {"en" : "An RFC 3339 date indicating the end of the log file.", "zh_CN": "RFC 3339格式的日期，表示日志文件的结束时间。"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty"`
  // {"en" : "Range: >= 0 
  // Size of log file in megabytes.", "zh_CN": "取值范围: >= 0 
  // 日志文件的大小（MB）。"}
  FileSize *int `json:"fileSize,omitempty" xml:"fileSize,omitempty"`
  // {"en" : "URL to the log file. The file has been compressed using gzip. Refer to the API description for details about the format.", "zh_CN": "日志文件的下载链接。该文件已使用gzip压缩。"}
  LogUrl *string `json:"logUrl,omitempty" xml:"logUrl,omitempty"`
  // {"en" : "The log file contains entries for this hostname.", "zh_CN": "日志文件对应的加速域名。"}
  Hostname *string `json:"hostname,omitempty" xml:"hostname,omitempty"`
}

func (s GetAccessLogsForHostnamesResponseLogs) String() string {
  return tea.Prettify(s)
}

func (s GetAccessLogsForHostnamesResponseLogs) GoString() string {
  return s.String()
}

func (s *GetAccessLogsForHostnamesResponseLogs) SetDateFrom(v string) *GetAccessLogsForHostnamesResponseLogs {
  s.DateFrom = &v
  return s
}

func (s *GetAccessLogsForHostnamesResponseLogs) SetDateTo(v string) *GetAccessLogsForHostnamesResponseLogs {
  s.DateTo = &v
  return s
}

func (s *GetAccessLogsForHostnamesResponseLogs) SetFileSize(v int) *GetAccessLogsForHostnamesResponseLogs {
  s.FileSize = &v
  return s
}

func (s *GetAccessLogsForHostnamesResponseLogs) SetLogUrl(v string) *GetAccessLogsForHostnamesResponseLogs {
  s.LogUrl = &v
  return s
}

func (s *GetAccessLogsForHostnamesResponseLogs) SetHostname(v string) *GetAccessLogsForHostnamesResponseLogs {
  s.Hostname = &v
  return s
}




type GetALogConfigurationRequest struct {
}

func (s GetALogConfigurationRequest) String() string {
  return tea.Prettify(s)
}

func (s GetALogConfigurationRequest) GoString() string {
  return s.String()
}

type GetALogConfigurationRequestHeader struct {
}

func (s GetALogConfigurationRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetALogConfigurationRequestHeader) GoString() string {
  return s.String()
}

type GetALogConfigurationPaths struct {
  // {"en":"ID of a log configuration","zh_CN":"日志配置的id。"}
  Id *int `json:"id,omitempty" xml:"id,omitempty" require:"true"`
}

func (s GetALogConfigurationPaths) String() string {
  return tea.Prettify(s)
}

func (s GetALogConfigurationPaths) GoString() string {
  return s.String()
}

func (s *GetALogConfigurationPaths) SetId(v int) *GetALogConfigurationPaths {
  s.Id = &v
  return s
}

type GetALogConfigurationParameters struct {
}

func (s GetALogConfigurationParameters) String() string {
  return tea.Prettify(s)
}

func (s GetALogConfigurationParameters) GoString() string {
  return s.String()
}

type GetALogConfigurationResponse struct {
  // {"en":"Range: >= 1\nID of the log configuration","zh_CN":"取值范围: >= 1\n日志配置的ID。"}
  Id *int `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {"en":"Describes a log configuration","zh_CN":"日志配置。"}
  Configs *GetALogConfigurationResponseConfigs `json:"configs,omitempty" xml:"configs,omitempty" require:"true" type:"Struct"`
  // {"en":"RFC 3339 date indicating when the log configuration was last updated.","zh_CN":"日志配置最近一次更新时间，以RFC 3339日期格式展示。"}
  LastUpdateTime *string `json:"lastUpdateTime,omitempty" xml:"lastUpdateTime,omitempty" require:"true"`
  // {"en":"RFC 3339 date indicating when the log configuration was created.","zh_CN":"日志配置创建时间，以RFC 3339日期格式展示。"}
  CreationTime *string `json:"creationTime,omitempty" xml:"creationTime,omitempty" require:"true"`
}

func (s GetALogConfigurationResponse) String() string {
  return tea.Prettify(s)
}

func (s GetALogConfigurationResponse) GoString() string {
  return s.String()
}

func (s *GetALogConfigurationResponse) SetId(v int) *GetALogConfigurationResponse {
  s.Id = &v
  return s
}

func (s *GetALogConfigurationResponse) SetConfigs(v *GetALogConfigurationResponseConfigs) *GetALogConfigurationResponse {
  s.Configs = v
  return s
}

func (s *GetALogConfigurationResponse) SetLastUpdateTime(v string) *GetALogConfigurationResponse {
  s.LastUpdateTime = &v
  return s
}

func (s *GetALogConfigurationResponse) SetCreationTime(v string) *GetALogConfigurationResponse {
  s.CreationTime = &v
  return s
}

type GetALogConfigurationResponseConfigs struct {
  // {"en":"A description of the log configuration.","zh_CN":"日志配置描述信息。"}
  Description *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
  // {"en":"A list of hostnames to which the log configuration applies. A hostname should be a fully qualified domain name like domain.com, a wildcard domain name with a leading '*' like *.domain.com, or '*' if you want the log configuration to apply to all hostnames. You should ensure that only one log configuration applies to each hostname.","zh_CN":"适用该日志配置的域名列表。每个域名必须是完整的FQDN域名如domain.com，或者是带星号的泛域名如*.domain.com。如果该日志配置适用于所有域名，则直接用星号*表示。每个域名有且只能有一个对应的日志配置。"}
  Hostnames []*string `json:"hostnames,omitempty" xml:"hostnames,omitempty" require:"true" type:"Repeated"`
  // {"en":"Format of the log entry.","zh_CN":"日志格式。"}
  LogDownloadFormat *string `json:"logDownloadFormat,omitempty" xml:"logDownloadFormat,omitempty" require:"true"`
  // {"en":"Default: 14 Range: [ 1 .. 30 ]\nNumber of days to store the logs.","zh_CN":"默认值: 14 取值范围: [ 1 .. 30 ]\n日志保存天数。"}
  LogDownloadStorageDays *int `json:"logDownloadStorageDays,omitempty" xml:"logDownloadStorageDays,omitempty" require:"true"`
  // {"en":"Enum: 5,15,30,60,120,240,480,1440\nDefault: 30\nTime span of each log file in minutes.","zh_CN":"取值范围: 5,15,30,60,120,240,480,1440\n默认值: 30\n每几分钟生成日志文件。"}
  LogDownloadFileSpanMinutes *int `json:"logDownloadFileSpanMinutes,omitempty" xml:"logDownloadFileSpanMinutes,omitempty" require:"true"`
}

func (s GetALogConfigurationResponseConfigs) String() string {
  return tea.Prettify(s)
}

func (s GetALogConfigurationResponseConfigs) GoString() string {
  return s.String()
}

func (s *GetALogConfigurationResponseConfigs) SetDescription(v string) *GetALogConfigurationResponseConfigs {
  s.Description = &v
  return s
}

func (s *GetALogConfigurationResponseConfigs) SetHostnames(v []*string) *GetALogConfigurationResponseConfigs {
  s.Hostnames = v
  return s
}

func (s *GetALogConfigurationResponseConfigs) SetLogDownloadFormat(v string) *GetALogConfigurationResponseConfigs {
  s.LogDownloadFormat = &v
  return s
}

func (s *GetALogConfigurationResponseConfigs) SetLogDownloadStorageDays(v int) *GetALogConfigurationResponseConfigs {
  s.LogDownloadStorageDays = &v
  return s
}

func (s *GetALogConfigurationResponseConfigs) SetLogDownloadFileSpanMinutes(v int) *GetALogConfigurationResponseConfigs {
  s.LogDownloadFileSpanMinutes = &v
  return s
}

type GetALogConfigurationResponseHeader struct {
}

func (s GetALogConfigurationResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetALogConfigurationResponseHeader) GoString() string {
  return s.String()
}




type UpdateALogConfigurationRequest struct {
  // {"en":"A description of the log configuration.","zh_CN":"日志配置描述。"}
  Description *string `json:"description,omitempty" xml:"description,omitempty"`
  // {"en":"A list of hostnames to which the log configuration applies. A hostname should be a fully qualified domain name like domain.com, a wildcard domain name with a leading '*' like *.domain.com, or '*' if you want the log configuration to apply to all hostnames. You should ensure that only one log configuration applies to each hostname.","zh_CN":"适用该日志配置的域名列表。每个域名必须是完整的FQDN域名如domain.com，或者是带星号的泛域名如*.domain.com。如果该日志配置适用于所有域名，则直接用星号*表示。每个域名有且只能有一个对应的日志配置。"}
  Hostnames []*string `json:"hostnames,omitempty" xml:"hostnames,omitempty" require:"true" type:"Repeated"`
  // {"en":"Default: %cltip %rmtuser  [%apachet] '%method %url %protocol' %statuscode %rspsize '%referer' '%ua' %rsptime\nFormat of the log entry. You can use the following variables to create a custom format:\n<table><tr><th>Variable</th><th>Meaning</th></tr><tr><td>%apachet</td><td>Time in the Common Log Format, for example, 03/Apr/2023:07:49:03 +0000</td></tr><tr><td>%cachestate</td><td>HIT, MISS, or REVALIDATE</td></tr><tr><td>%cltip</td><td>Client IP address</td></tr><tr><td>%cltisp</td><td>Client ISP</td></tr><tr><td>%cltport</td><td>Client port number</td></tr><tr><td>%cltregion</td><td>Client region</td></tr><tr><td>%cpu_ns</td><td>CPU time in nanoseconds for this request</td></tr><tr><td>%custom1</td><td>Refers to a custom log field with ID 1 you define using the custom_log_field directive.</td></tr>      <tr><td>%custom2</td><td>Refers to a custom log field with ID 2 you define using the custom_log_field directive.</td></tr>      <tr><td>%hostname</td><td>Host header</td></tr><tr><td>%method</td><td>HTTP method used to access the content, for example, GET</td></tr><tr><td>%protocol</td><td>HTTP/1.0 or HTTP/1.1 or HTTP2.0</td></tr><tr><td>%querystr</td><td>Query string</td></tr>      <tr><td>%referer</td><td>Referer request header</td></tr><tr><td>%reqhdrsize</td><td>Request header size in bytes</td></tr><tr><td>%reqrange</td><td>Range header in requests from client</td></tr><tr><td>%reqsize</td><td>Request size in bytes including the request line, header, and request body</td></tr><tr><td>%rmtuser</td><td>User name extracted from the Authorization header when basic authentication is used</td></tr><tr><td>%rspsize</td><td>HTTP response size, in bytes, including header and body, but not including TCP/IP/MAC</td></tr><tr><td>%rsptime</td><td>Response time in milliseconds</td></tr><tr><td>%samplerate</td><td>Number of requests corresponding to each log entry</td></tr><tr><td>%scheme</td><td>http or https</td></tr><tr><td>%statuscode</td><td>Response's status code, for example, 200</td></tr><tr><td>%srvip</td><td>IP address of the CDN server handling the request</td></tr><tr><td>%svrnode</td><td>Cache server node name</td></tr><tr><td>%tcprtt</td><td>Round-trip time between server and client in microseconds</td></tr>      <tr><td>%ua</td><td>User-Agent header</td></tr><tr><td>%uniqueid</td><td>A string uniquely identifying this request</td></tr>       <tr><td>%url</td><td>Full URL with query string, if any</td></tr><tr><td>%utctime</td><td>RFC 3339 timestamp of the response, for example, 2021-10-05T12:34:56Z</td></tr></table>\nIf any variable value contains non-printable characters (ASCII characters <=0x1F or >=0x7F), double quotes, or a backslash, each of those characters will be escaped to a format like '\xXX'. For example, double quotes will become '\x22' and backslash will become '\x5C'.\nNote that space is not escaped. You must quote variables that may contain spaces to facilitate parsing.","zh_CN":"默认值: %cltip %rmtuser  [%apachet] '%method %url %protocol' %statuscode %rspsize '%referer' '%ua' %rsptime\n日志格式。可使用以下变量来自定义日志输出格式：\n<table><tr><th>变量</th><th>变量说明</th></tr><tr><td>%apachet</td><td>CLF（Common Log Format）时间格式，例如03/Apr/2023:07:49:03 +0000。</td></tr><tr><td>%cachestate</td><td>缓存状态，取值为HIT，MISS 或者 REVALIDATE。</td></tr><tr><td>%cltip</td><td>客户端IP地址。</td></tr><tr><td>%cltisp</td><td>客户端所属运营商。</td></tr><tr><td>%cltport</td><td>客户端端口号。</td></tr><tr><td>%cltregion</td><td>客户端所属区域。</td></tr><tr><td>%cpu_ns</td><td>此次请求所消耗的CPU时间，以纳秒计算。</td></tr><tr><td>%custom1</td><td>如果您在加速项目中自定义了日志字段，可以用该变量来获取自定义日志字段1的值。自定义日志字段属于高级功能，如需使用请联系我们的技术支持。</td></tr>      <tr><td>%custom2</td><td>如果您在加速项目中自定义了日志字段，可以用该变量来获取自定义日志字段2的值。自定义日志字段属于高级功能，如需使用请联系我们的技术支持。</td></tr>      <tr><td>%hostname</td><td>Host 请求头。</td></tr><tr><td>%method</td><td>HTTP请求方法，例如 GET。</td></tr><tr><td>%protocol</td><td>HTTP/1.0，HTTP/1.1 或者 HTTP2.0协议。</td></tr><tr><td>%querystr</td><td>查询参数字符串。</td></tr>      <tr><td>%referer</td><td>Referer 请求头。</td></tr><tr><td>%reqhdrsize</td><td>请求头大小，单位为bytes。</td></tr><tr><td>%reqrange</td><td>客户端请求所携带的Range请求头。</td></tr><tr><td>%reqsize</td><td>请求的大小，单位为bytes。包括请求行，请求头和请求体。</td></tr><tr><td>%rmtuser</td><td>从Authorization请求头解析出的用户名。当您使用HTTP basic对用户请求进行鉴权时，可通过该变量获取用户名。</td></tr><tr><td>%rspsize</td><td>HTTP响应的大小，单位为bytes。包括响应头和响应体，但不包括TCP/IP/MAC头部开销。</td></tr><tr><td>%rsptime</td><td>响应时间，单位为毫秒。</td></tr><tr><td>%samplerate</td><td>采样率，即每几次请求输出一条日志。</td></tr><tr><td>%sheme</td><td>http或https协议。</td></tr><tr><td>%statuscode</td><td>响应状态码，例如200。</td></tr><tr><td>%srvip</td><td>处理此次请求的CDN服务器的IP地址。</td></tr><tr><td>%svrnode</td><td>CDN服务器所在的节点名称。</td></tr><tr><td>%tcprtt</td><td>服务器与客户端之间的RTT时长，单位为毫秒。</td></tr>      <tr><td>%ua</td><td>User-Agent 请求头。</td></tr><tr><td>%uniqueid</td><td>此次请求的唯一标识，字符串类型。</td></tr>       <tr><td>%url</td><td>完整的URL，包括查询参数字符串。</td></tr><tr><td>%utctime</td><td>响应时间，以RFC 3339格式表示，例如 2021-10-05T12:34:56Z。</td></tr></table>\n\n变量如果包含了不可打印的字符（ASCII码<=0x1F 或 >=0x7F的字符），双引号或者反斜杆，则这些字符会被转义，用'\xXX'的格式表示。例如，双引号以'\x22'表示，反斜杆以'\x5C'表示。需要注意的是，空格不会被转义。如果变量包含空格，则必须加上双引号，才能被正确解析。"}
  LogDownloadFormat *string `json:"logDownloadFormat,omitempty" xml:"logDownloadFormat,omitempty"`
  // {"en":"Default: 14 Range: [ 1 .. 30 ]\nNumber of days to store the logs.","zh_CN":"默认值: 14 取值范围: [ 1 .. 30 ]\n日志保存天数。"}
  LogDownloadStorageDays *int `json:"logDownloadStorageDays,omitempty" xml:"logDownloadStorageDays,omitempty"`
  // {"en":"Enum: 5,15,30,60,120,240,480,1440\nDefault: 30\nTime span of each log file in minutes.","zh_CN":"取值范围: 5,15,30,60,120,240,480,1440\n默认值: 30\n每几分钟生成日志文件。"}
  LogDownloadFileSpanMinutes *int `json:"logDownloadFileSpanMinutes,omitempty" xml:"logDownloadFileSpanMinutes,omitempty"`
}

func (s UpdateALogConfigurationRequest) String() string {
  return tea.Prettify(s)
}

func (s UpdateALogConfigurationRequest) GoString() string {
  return s.String()
}

func (s *UpdateALogConfigurationRequest) SetDescription(v string) *UpdateALogConfigurationRequest {
  s.Description = &v
  return s
}

func (s *UpdateALogConfigurationRequest) SetHostnames(v []*string) *UpdateALogConfigurationRequest {
  s.Hostnames = v
  return s
}

func (s *UpdateALogConfigurationRequest) SetLogDownloadFormat(v string) *UpdateALogConfigurationRequest {
  s.LogDownloadFormat = &v
  return s
}

func (s *UpdateALogConfigurationRequest) SetLogDownloadStorageDays(v int) *UpdateALogConfigurationRequest {
  s.LogDownloadStorageDays = &v
  return s
}

func (s *UpdateALogConfigurationRequest) SetLogDownloadFileSpanMinutes(v int) *UpdateALogConfigurationRequest {
  s.LogDownloadFileSpanMinutes = &v
  return s
}

type UpdateALogConfigurationRequestHeader struct {
}

func (s UpdateALogConfigurationRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s UpdateALogConfigurationRequestHeader) GoString() string {
  return s.String()
}

type UpdateALogConfigurationPaths struct {
  // {"en":"ID of a log configuration","zh_CN":"日志配置的id。"}
  Id *int `json:"id,omitempty" xml:"id,omitempty" require:"true"`
}

func (s UpdateALogConfigurationPaths) String() string {
  return tea.Prettify(s)
}

func (s UpdateALogConfigurationPaths) GoString() string {
  return s.String()
}

func (s *UpdateALogConfigurationPaths) SetId(v int) *UpdateALogConfigurationPaths {
  s.Id = &v
  return s
}

type UpdateALogConfigurationParameters struct {
}

func (s UpdateALogConfigurationParameters) String() string {
  return tea.Prettify(s)
}

func (s UpdateALogConfigurationParameters) GoString() string {
  return s.String()
}

type UpdateALogConfigurationResponse struct {
}

func (s UpdateALogConfigurationResponse) String() string {
  return tea.Prettify(s)
}

func (s UpdateALogConfigurationResponse) GoString() string {
  return s.String()
}

type UpdateALogConfigurationResponseHeader struct {
}

func (s UpdateALogConfigurationResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s UpdateALogConfigurationResponseHeader) GoString() string {
  return s.String()
}




