package incidentlogs

// This file is auto-generated, don't edit it. Thanks.
import (
  "github.com/alibabacloud-go/tea/tea"
)

type DownloadAttackLogRequest struct {
  // {"en":"Start time 2020-01-01 00:00:00  interval between start time and end time < = 1 day", "zh_CN":"开始时间 2020-01-01 00:00:00 开始时间与结束时间间隔<=1天"}
  StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty" require:"true"`
  // {"en":"End time 2020-01-02 00:00:00  interval between start time and end time < = 1 day", "zh_CN":"结束时间 2020-01-02 00:00:00 开始时间与结束时间间隔<=1天"}
  EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty" require:"true"`
  // {"en":"language default:cn
  //     cn Chinese
  //     en English", "zh_CN":"语言 默认为中文
  //     cn 中文
  //     en 英文"}
  Language *string `json:"language,omitempty" xml:"language,omitempty"`
  // {"en":"Domain name. Multiple values are separated by English commas.If blank, query all.", "zh_CN":"域名，多个值使用英文逗号分隔，域名为空查询全部"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty"`
  // {"en":"Client IP, multiple are separated by English commas", "zh_CN":"客户端IP，多个使用英文逗号分隔"}
  ClientIp *string `json:"clientIp,omitempty" xml:"clientIp,omitempty"`
  // {"en":"Protection strategy, multiple are separated by English commas
  // 001 built in protection
  // 002 HTTP watermark protection
  // 003 access control
  // 004 frequency limit
  // 005 four layer DDoS protection
  // ", "zh_CN":"防护策略，多个使用英文逗号分隔
  // 001  内置防护
  // 002  HTTP水印防护
  // 003  访问控制
  // 004  频率限制
  // 005  四层DDoS防护
  // "}
  AttackType *string `json:"attackType,omitempty" xml:"attackType,omitempty" require:"true"`
  // {"en":"Processing actions, multiple are separated by English commas, and the input parameter is empty to query all
  // Intercept: 1
  // Man machine verification: 2
  // IntelligentBlocking: 3
  // Alarm: 4
  // Custom block page: 5
  // Disconnect: 6
  // ", "zh_CN":"处理动作，多个使用英文逗号分隔,入参为空查询全部
  // 拦截：1
  // 人机校验：2
  // 智能封禁：3
  // 告警：4
  // 跳转友好界面：5
  // 断开连接：6
  // "}
  ProcessAction *string `json:"processAction,omitempty" xml:"processAction,omitempty" require:"true"`
  // {"en":"uri", "zh_CN":"uri"}
  Uri *string `json:"uri,omitempty" xml:"uri,omitempty"`
  // {"en":"Whether the URIs match exactly, 1 yes", "zh_CN":"uri是否精确匹配，1是"}
  AccurateUri *int `json:"accurateUri,omitempty" xml:"accurateUri,omitempty"`
  // {"en":"userAgent", "zh_CN":"userAgent"}
  UserAgent *string `json:"userAgent,omitempty" xml:"userAgent,omitempty"`
  // {"en":"Whether useragent exactly matches, 1 yes", "zh_CN":"useragent 是否精确匹配, 1 yes"}
  AccurateUserAgent *int `json:"accurateUserAgent,omitempty" xml:"accurateUserAgent,omitempty"`
  // {"en":"Status codes, multiple separated by commas", "zh_CN":"状态码，多个使用逗号分隔"}
  StateCode *string `json:"stateCode,omitempty" xml:"stateCode,omitempty"`
  // {"en":"referer", "zh_CN":"referer"}
  Referer *string `json:"referer,omitempty" xml:"referer,omitempty"`
  // {"en":"Request mode: multiple are separated by commas (get, post, head, options, put, delete, trace, connect)", "zh_CN":"请求方式，多个使用逗号分隔（GET,POST,HEAD,OPTIONS,PUT,DELETE,TRACE,CONNECT）"}
  RequestMode *string `json:"requestMode,omitempty" xml:"requestMode,omitempty"`
}

func (s DownloadAttackLogRequest) String() string {
  return tea.Prettify(s)
}

func (s DownloadAttackLogRequest) GoString() string {
  return s.String()
}

func (s *DownloadAttackLogRequest) SetStartTime(v string) *DownloadAttackLogRequest {
  s.StartTime = &v
  return s
}

func (s *DownloadAttackLogRequest) SetEndTime(v string) *DownloadAttackLogRequest {
  s.EndTime = &v
  return s
}

func (s *DownloadAttackLogRequest) SetLanguage(v string) *DownloadAttackLogRequest {
  s.Language = &v
  return s
}

func (s *DownloadAttackLogRequest) SetDomain(v string) *DownloadAttackLogRequest {
  s.Domain = &v
  return s
}

func (s *DownloadAttackLogRequest) SetClientIp(v string) *DownloadAttackLogRequest {
  s.ClientIp = &v
  return s
}

func (s *DownloadAttackLogRequest) SetAttackType(v string) *DownloadAttackLogRequest {
  s.AttackType = &v
  return s
}

func (s *DownloadAttackLogRequest) SetProcessAction(v string) *DownloadAttackLogRequest {
  s.ProcessAction = &v
  return s
}

func (s *DownloadAttackLogRequest) SetUri(v string) *DownloadAttackLogRequest {
  s.Uri = &v
  return s
}

func (s *DownloadAttackLogRequest) SetAccurateUri(v int) *DownloadAttackLogRequest {
  s.AccurateUri = &v
  return s
}

func (s *DownloadAttackLogRequest) SetUserAgent(v string) *DownloadAttackLogRequest {
  s.UserAgent = &v
  return s
}

func (s *DownloadAttackLogRequest) SetAccurateUserAgent(v int) *DownloadAttackLogRequest {
  s.AccurateUserAgent = &v
  return s
}

func (s *DownloadAttackLogRequest) SetStateCode(v string) *DownloadAttackLogRequest {
  s.StateCode = &v
  return s
}

func (s *DownloadAttackLogRequest) SetReferer(v string) *DownloadAttackLogRequest {
  s.Referer = &v
  return s
}

func (s *DownloadAttackLogRequest) SetRequestMode(v string) *DownloadAttackLogRequest {
  s.RequestMode = &v
  return s
}

type DownloadAttackLogResponse struct {
  // {"en":"accessTime", "zh_CN":"访问时间"}
  AccessTime *string `json:"accessTime,omitempty" xml:"accessTime,omitempty" require:"true"`
  // {"en":"clientIp", "zh_CN":"客户端ip"}
  ClientIp *string `json:"clientIp,omitempty" xml:"clientIp,omitempty" require:"true"`
  // {"en":"domain", "zh_CN":"域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {"en":"requestMethod", "zh_CN":"请求方式"}
  RequestMethod *string `json:"requestMethod,omitempty" xml:"requestMethod,omitempty" require:"true"`
  // {"en":"uri", "zh_CN":"uri"}
  Uri *string `json:"uri,omitempty" xml:"uri,omitempty" require:"true"`
  // {"en":"referer", "zh_CN":"referer"}
  Referer *string `json:"referer,omitempty" xml:"referer,omitempty" require:"true"`
  // {"en":"userAgent", "zh_CN":"userAgent"}
  UserAgent *string `json:"userAgent,omitempty" xml:"userAgent,omitempty" require:"true"`
  // {"en":"stateCode", "zh_CN":"状态码"}
  StateCode *string `json:"stateCode,omitempty" xml:"stateCode,omitempty" require:"true"`
  // {"en":"dstPort", "zh_CN":"端口"}
  DstPort *string `json:"dstPort,omitempty" xml:"dstPort,omitempty" require:"true"`
  // {"en":"agreement", "zh_CN":"协议"}
  Agreement *string `json:"agreement,omitempty" xml:"agreement,omitempty" require:"true"`
  // {"en":"ruleName", "zh_CN":"规则名称"}
  RuleName *string `json:"ruleName,omitempty" xml:"ruleName,omitempty" require:"true"`
  // {'en':'processAction','zh_CN':'处理动作'}
  ProcessAction *string `json:"processAction,omitempty" xml:"processAction,omitempty" require:"true"`
  // {"en":"ip location", "zh_CN":"ip地理位置"}
  IpLocation *string `json:"ipLocation,omitempty" xml:"ipLocation,omitempty" require:"true"`
  // {"en":"httpVersion", "zh_CN":"http版本"}
  HttpVersion *string `json:"httpVersion,omitempty" xml:"httpVersion,omitempty" require:"true"`
}

func (s DownloadAttackLogResponse) String() string {
  return tea.Prettify(s)
}

func (s DownloadAttackLogResponse) GoString() string {
  return s.String()
}

func (s *DownloadAttackLogResponse) SetAccessTime(v string) *DownloadAttackLogResponse {
  s.AccessTime = &v
  return s
}

func (s *DownloadAttackLogResponse) SetClientIp(v string) *DownloadAttackLogResponse {
  s.ClientIp = &v
  return s
}

func (s *DownloadAttackLogResponse) SetDomain(v string) *DownloadAttackLogResponse {
  s.Domain = &v
  return s
}

func (s *DownloadAttackLogResponse) SetRequestMethod(v string) *DownloadAttackLogResponse {
  s.RequestMethod = &v
  return s
}

func (s *DownloadAttackLogResponse) SetUri(v string) *DownloadAttackLogResponse {
  s.Uri = &v
  return s
}

func (s *DownloadAttackLogResponse) SetReferer(v string) *DownloadAttackLogResponse {
  s.Referer = &v
  return s
}

func (s *DownloadAttackLogResponse) SetUserAgent(v string) *DownloadAttackLogResponse {
  s.UserAgent = &v
  return s
}

func (s *DownloadAttackLogResponse) SetStateCode(v string) *DownloadAttackLogResponse {
  s.StateCode = &v
  return s
}

func (s *DownloadAttackLogResponse) SetDstPort(v string) *DownloadAttackLogResponse {
  s.DstPort = &v
  return s
}

func (s *DownloadAttackLogResponse) SetAgreement(v string) *DownloadAttackLogResponse {
  s.Agreement = &v
  return s
}

func (s *DownloadAttackLogResponse) SetRuleName(v string) *DownloadAttackLogResponse {
  s.RuleName = &v
  return s
}

func (s *DownloadAttackLogResponse) SetProcessAction(v string) *DownloadAttackLogResponse {
  s.ProcessAction = &v
  return s
}

func (s *DownloadAttackLogResponse) SetIpLocation(v string) *DownloadAttackLogResponse {
  s.IpLocation = &v
  return s
}

func (s *DownloadAttackLogResponse) SetHttpVersion(v string) *DownloadAttackLogResponse {
  s.HttpVersion = &v
  return s
}

type DownloadAttackLogPaths struct {
}

func (s DownloadAttackLogPaths) String() string {
  return tea.Prettify(s)
}

func (s DownloadAttackLogPaths) GoString() string {
  return s.String()
}

type DownloadAttackLogParameters struct {
}

func (s DownloadAttackLogParameters) String() string {
  return tea.Prettify(s)
}

func (s DownloadAttackLogParameters) GoString() string {
  return s.String()
}

type DownloadAttackLogRequestHeader struct {
}

func (s DownloadAttackLogRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s DownloadAttackLogRequestHeader) GoString() string {
  return s.String()
}

type DownloadAttackLogResponseHeader struct {
}

func (s DownloadAttackLogResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s DownloadAttackLogResponseHeader) GoString() string {
  return s.String()
}




type QueryAttackLogDetailsRequest struct {
  // {"en":"Start time 2020-01-01 00:00:00    interval between start time and end time < = one day", "zh_CN":"开始时间  2020-01-01 00:00:00  开始时间与结束时间间隔<=1天"}
  StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty" require:"true"`
  // {"en":"End time 2020-01-02 00:00:00   interval between start time and end time < = one day", "zh_CN":"结束时间  2020-01-02 00:00:00  开始时间与结束时间间隔<=1天"}
  EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty" require:"true"`
  // {"en":"language default:cn
  //     cn Chinese
  //     en English", "zh_CN":"语言 默认为中文
  //     cn 中文
  //     en 英文"}
  Language *string `json:"language,omitempty" xml:"language,omitempty"`
  // {"en":"Domain name. Multiple values are separated by English commas.If blank, query all.", "zh_CN":"域名，多个值使用英文逗号分隔，域名为空查询全部"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty"`
  // {"en":"Client IP, multiple are separated by English commas", "zh_CN":"客户端IP，多个使用英文逗号分隔"}
  ClientIp *string `json:"clientIp,omitempty" xml:"clientIp,omitempty"`
  // {"en":"Protection strategy, multiple are separated by English commas
  // 001 built in protection
  // 002 HTTP watermark protection
  // 003 access control
  // 004 frequency limit
  // 005 four layer DDoS protection
  // ", "zh_CN":"防护策略，多个使用英文逗号分隔
  // 001  内置防护
  // 002  HTTP水印防护
  // 003  访问控制
  // 004  频率限制
  // 005  四层DDoS防护
  // "}
  AttackType *string `json:"attackType,omitempty" xml:"attackType,omitempty" require:"true"`
  // {"en":"Processing actions, multiple are separated by English commas, and the input parameter is empty to query all
  // Intercept: 1
  // Man machine verification: 2
  // IntelligentBlocking: 3
  // Alarm: 4
  // Custom block page: 5
  // Disconnect: 6
  // ", "zh_CN":"处理动作，多个使用英文逗号分隔,入参为空查询全部
  // 拦截：1
  // 人机校验：2
  // 智能封禁：3
  // 告警：4
  // 跳转友好界面：5
  // 断开连接：6
  // 注：防护策略和处理动作必填"}
  ProcessAction *string `json:"processAction,omitempty" xml:"processAction,omitempty" require:"true"`
  // {"en":"uri", "zh_CN":"uri"}
  Uri *string `json:"uri,omitempty" xml:"uri,omitempty"`
  // {"en":"Whether the URIs match exactly, 1 yes", "zh_CN":"uri是否精确匹配，1是"}
  AccurateUri *int `json:"accurateUri,omitempty" xml:"accurateUri,omitempty"`
  // {"en":"userAgent", "zh_CN":"userAgent"}
  UserAgent *string `json:"userAgent,omitempty" xml:"userAgent,omitempty"`
  // {"en":"Whether useragent exactly matches, 1 yes", "zh_CN":"Whether useragent exactly matches, 1 yes"}
  AccurateUserAgent *int `json:"accurateUserAgent,omitempty" xml:"accurateUserAgent,omitempty"`
  // {"en":"Status codes, multiple separated by commas", "zh_CN":"状态码，多个使用逗号分隔"}
  StateCode *string `json:"stateCode,omitempty" xml:"stateCode,omitempty"`
  // {"en":"referer", "zh_CN":"referer"}
  Referer *string `json:"referer,omitempty" xml:"referer,omitempty"`
  // {"en":"Request mode: multiple are separated by commas (get, post, head, options, put, delete, trace, connect)", "zh_CN":"请求方式，多个使用逗号分隔（GET,POST,HEAD,OPTIONS,PUT,DELETE,TRACE,CONNECT）"}
  RequestMode *string `json:"requestMode,omitempty" xml:"requestMode,omitempty"`
}

func (s QueryAttackLogDetailsRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryAttackLogDetailsRequest) GoString() string {
  return s.String()
}

func (s *QueryAttackLogDetailsRequest) SetStartTime(v string) *QueryAttackLogDetailsRequest {
  s.StartTime = &v
  return s
}

func (s *QueryAttackLogDetailsRequest) SetEndTime(v string) *QueryAttackLogDetailsRequest {
  s.EndTime = &v
  return s
}

func (s *QueryAttackLogDetailsRequest) SetLanguage(v string) *QueryAttackLogDetailsRequest {
  s.Language = &v
  return s
}

func (s *QueryAttackLogDetailsRequest) SetDomain(v string) *QueryAttackLogDetailsRequest {
  s.Domain = &v
  return s
}

func (s *QueryAttackLogDetailsRequest) SetClientIp(v string) *QueryAttackLogDetailsRequest {
  s.ClientIp = &v
  return s
}

func (s *QueryAttackLogDetailsRequest) SetAttackType(v string) *QueryAttackLogDetailsRequest {
  s.AttackType = &v
  return s
}

func (s *QueryAttackLogDetailsRequest) SetProcessAction(v string) *QueryAttackLogDetailsRequest {
  s.ProcessAction = &v
  return s
}

func (s *QueryAttackLogDetailsRequest) SetUri(v string) *QueryAttackLogDetailsRequest {
  s.Uri = &v
  return s
}

func (s *QueryAttackLogDetailsRequest) SetAccurateUri(v int) *QueryAttackLogDetailsRequest {
  s.AccurateUri = &v
  return s
}

func (s *QueryAttackLogDetailsRequest) SetUserAgent(v string) *QueryAttackLogDetailsRequest {
  s.UserAgent = &v
  return s
}

func (s *QueryAttackLogDetailsRequest) SetAccurateUserAgent(v int) *QueryAttackLogDetailsRequest {
  s.AccurateUserAgent = &v
  return s
}

func (s *QueryAttackLogDetailsRequest) SetStateCode(v string) *QueryAttackLogDetailsRequest {
  s.StateCode = &v
  return s
}

func (s *QueryAttackLogDetailsRequest) SetReferer(v string) *QueryAttackLogDetailsRequest {
  s.Referer = &v
  return s
}

func (s *QueryAttackLogDetailsRequest) SetRequestMode(v string) *QueryAttackLogDetailsRequest {
  s.RequestMode = &v
  return s
}

type QueryAttackLogDetailsResponse struct {
  // {"en":"response code", "zh_CN":"响应码，响应成功为200，响应失败见状态码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"response message", "zh_CN":"响应信息，响应成功为success"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  Data *QueryAttackLogDetailsResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
}

func (s QueryAttackLogDetailsResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryAttackLogDetailsResponse) GoString() string {
  return s.String()
}

func (s *QueryAttackLogDetailsResponse) SetCode(v string) *QueryAttackLogDetailsResponse {
  s.Code = &v
  return s
}

func (s *QueryAttackLogDetailsResponse) SetMsg(v string) *QueryAttackLogDetailsResponse {
  s.Msg = &v
  return s
}

func (s *QueryAttackLogDetailsResponse) SetData(v *QueryAttackLogDetailsResponseData) *QueryAttackLogDetailsResponse {
  s.Data = v
  return s
}

type QueryAttackLogDetailsResponseData struct {
  // {'en':'accessTime','zh_CN':'访问时间'}
  AccessTime *string `json:"accessTime,omitempty" xml:"accessTime,omitempty" require:"true"`
  // {'en':'clientIp','zh_CN':'客户端IP'}
  ClientIp *string `json:"clientIp,omitempty" xml:"clientIp,omitempty" require:"true"`
  // {'en':'domain','zh_CN':'域名'}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {'en':'attackType','zh_CN':'攻击类型'}
  AttackType *string `json:"attackType,omitempty" xml:"attackType,omitempty" require:"true"`
  // {'en':'processAction','zh_CN':'处理动作'}
  ProcessAction *string `json:"processAction,omitempty" xml:"processAction,omitempty" require:"true"`
  // {'en':'proStrategy','zh_CN':'防护策略'}
  ProStrategy *string `json:"proStrategy,omitempty" xml:"proStrategy,omitempty" require:"true"`
  // {'en':'stateCode','zh_CN':'状态码'}
  StateCode *string `json:"stateCode,omitempty" xml:"stateCode,omitempty" require:"true"`
  HttpRequestInfo *QueryAttackLogDetailsResponseDataHttpRequestInfo `json:"httpRequestInfo,omitempty" xml:"httpRequestInfo,omitempty" require:"true" type:"Struct"`
  TcpRequestInfo *QueryAttackLogDetailsResponseDataTcpRequestInfo `json:"tcpRequestInfo,omitempty" xml:"tcpRequestInfo,omitempty" require:"true" type:"Struct"`
}

func (s QueryAttackLogDetailsResponseData) String() string {
  return tea.Prettify(s)
}

func (s QueryAttackLogDetailsResponseData) GoString() string {
  return s.String()
}

func (s *QueryAttackLogDetailsResponseData) SetAccessTime(v string) *QueryAttackLogDetailsResponseData {
  s.AccessTime = &v
  return s
}

func (s *QueryAttackLogDetailsResponseData) SetClientIp(v string) *QueryAttackLogDetailsResponseData {
  s.ClientIp = &v
  return s
}

func (s *QueryAttackLogDetailsResponseData) SetDomain(v string) *QueryAttackLogDetailsResponseData {
  s.Domain = &v
  return s
}

func (s *QueryAttackLogDetailsResponseData) SetAttackType(v string) *QueryAttackLogDetailsResponseData {
  s.AttackType = &v
  return s
}

func (s *QueryAttackLogDetailsResponseData) SetProcessAction(v string) *QueryAttackLogDetailsResponseData {
  s.ProcessAction = &v
  return s
}

func (s *QueryAttackLogDetailsResponseData) SetProStrategy(v string) *QueryAttackLogDetailsResponseData {
  s.ProStrategy = &v
  return s
}

func (s *QueryAttackLogDetailsResponseData) SetStateCode(v string) *QueryAttackLogDetailsResponseData {
  s.StateCode = &v
  return s
}

func (s *QueryAttackLogDetailsResponseData) SetHttpRequestInfo(v *QueryAttackLogDetailsResponseDataHttpRequestInfo) *QueryAttackLogDetailsResponseData {
  s.HttpRequestInfo = v
  return s
}

func (s *QueryAttackLogDetailsResponseData) SetTcpRequestInfo(v *QueryAttackLogDetailsResponseDataTcpRequestInfo) *QueryAttackLogDetailsResponseData {
  s.TcpRequestInfo = v
  return s
}

type QueryAttackLogDetailsResponseDataHttpRequestInfo struct {
  // {'en':'locationIp','zh_CN':'ip地址'}
  LocationIp *string `json:"locationIp,omitempty" xml:"locationIp,omitempty" require:"true"`
  // {'en':'httpVersion','zh_CN':'HTTP版本'}
  HttpVersion *string `json:"httpVersion,omitempty" xml:"httpVersion,omitempty" require:"true"`
  // {'en':'requestMethod','zh_CN':'请求方法'}
  RequestMethod *string `json:"requestMethod,omitempty" xml:"requestMethod,omitempty" require:"true"`
  // {'en':'uri','zh_CN':'uri'}
  Uri *string `json:"uri,omitempty" xml:"uri,omitempty" require:"true"`
  // {'en':'referer','zh_CN':'referer'}
  Referer *string `json:"referer,omitempty" xml:"referer,omitempty" require:"true"`
  // {'en':'userAgent','zh_CN':'userAgent'}
  UserAgent *string `json:"userAgent,omitempty" xml:"userAgent,omitempty" require:"true"`
}

func (s QueryAttackLogDetailsResponseDataHttpRequestInfo) String() string {
  return tea.Prettify(s)
}

func (s QueryAttackLogDetailsResponseDataHttpRequestInfo) GoString() string {
  return s.String()
}

func (s *QueryAttackLogDetailsResponseDataHttpRequestInfo) SetLocationIp(v string) *QueryAttackLogDetailsResponseDataHttpRequestInfo {
  s.LocationIp = &v
  return s
}

func (s *QueryAttackLogDetailsResponseDataHttpRequestInfo) SetHttpVersion(v string) *QueryAttackLogDetailsResponseDataHttpRequestInfo {
  s.HttpVersion = &v
  return s
}

func (s *QueryAttackLogDetailsResponseDataHttpRequestInfo) SetRequestMethod(v string) *QueryAttackLogDetailsResponseDataHttpRequestInfo {
  s.RequestMethod = &v
  return s
}

func (s *QueryAttackLogDetailsResponseDataHttpRequestInfo) SetUri(v string) *QueryAttackLogDetailsResponseDataHttpRequestInfo {
  s.Uri = &v
  return s
}

func (s *QueryAttackLogDetailsResponseDataHttpRequestInfo) SetReferer(v string) *QueryAttackLogDetailsResponseDataHttpRequestInfo {
  s.Referer = &v
  return s
}

func (s *QueryAttackLogDetailsResponseDataHttpRequestInfo) SetUserAgent(v string) *QueryAttackLogDetailsResponseDataHttpRequestInfo {
  s.UserAgent = &v
  return s
}

type QueryAttackLogDetailsResponseDataTcpRequestInfo struct {
  // {'en':'locationIp','zh_CN':'IP地理位置'}
  LocationIp *string `json:"locationIp,omitempty" xml:"locationIp,omitempty" require:"true"`
  // {'en':'locationIpEn','zh_CN':'IP地理位置英文名称'}
  LocationIpEn *string `json:"locationIpEn,omitempty" xml:"locationIpEn,omitempty" require:"true"`
  // {'en':'dsPort','zh_CN':'端口'}
  DsPort *string `json:"dsPort,omitempty" xml:"dsPort,omitempty" require:"true"`
  // {'en':'agreement','zh_CN':'协议'}
  Agreement *string `json:"agreement,omitempty" xml:"agreement,omitempty" require:"true"`
}

func (s QueryAttackLogDetailsResponseDataTcpRequestInfo) String() string {
  return tea.Prettify(s)
}

func (s QueryAttackLogDetailsResponseDataTcpRequestInfo) GoString() string {
  return s.String()
}

func (s *QueryAttackLogDetailsResponseDataTcpRequestInfo) SetLocationIp(v string) *QueryAttackLogDetailsResponseDataTcpRequestInfo {
  s.LocationIp = &v
  return s
}

func (s *QueryAttackLogDetailsResponseDataTcpRequestInfo) SetLocationIpEn(v string) *QueryAttackLogDetailsResponseDataTcpRequestInfo {
  s.LocationIpEn = &v
  return s
}

func (s *QueryAttackLogDetailsResponseDataTcpRequestInfo) SetDsPort(v string) *QueryAttackLogDetailsResponseDataTcpRequestInfo {
  s.DsPort = &v
  return s
}

func (s *QueryAttackLogDetailsResponseDataTcpRequestInfo) SetAgreement(v string) *QueryAttackLogDetailsResponseDataTcpRequestInfo {
  s.Agreement = &v
  return s
}

type QueryAttackLogDetailsPaths struct {
}

func (s QueryAttackLogDetailsPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryAttackLogDetailsPaths) GoString() string {
  return s.String()
}

type QueryAttackLogDetailsParameters struct {
}

func (s QueryAttackLogDetailsParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryAttackLogDetailsParameters) GoString() string {
  return s.String()
}

type QueryAttackLogDetailsRequestHeader struct {
}

func (s QueryAttackLogDetailsRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryAttackLogDetailsRequestHeader) GoString() string {
  return s.String()
}

type QueryAttackLogDetailsResponseHeader struct {
}

func (s QueryAttackLogDetailsResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryAttackLogDetailsResponseHeader) GoString() string {
  return s.String()
}




type GetWAFAttackLogRequest struct {
  // {"en":"Domain, array.", "zh_CN":"域名，数组。"}
  Domains []*string `json:"domains,omitempty" xml:"domains,omitempty" require:"true" type:"Repeated"`
  // {"en":"Start time, yyyy-MM-dd HH:mm:ss.", "zh_CN":"开始时间，yyyy-MM-dd HH:mm:ss。"}
  StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty" require:"true"`
  // {"en":"End time, yyyy-MM-dd HH:mm:ss.", "zh_CN":"结束时间，yyyy-MM-dd HH:mm:ss。"}
  EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty" require:"true"`
  // {"en":"Time zone, GMT+8 by default.", "zh_CN":"时区，默认GMT+8，即“GMT+8”。"}
  TimeZone *string `json:"timeZone,omitempty" xml:"timeZone,omitempty"`
  // {"en":"Action,default value: 1 and 2.
  //     1: Block
  //     2: GetWAFAttackLogLog", "zh_CN":"处理动作，默认值：1和2。
  //     1：拦截
  //     2：监控"}
  Acts []*string `json:"acts,omitempty" xml:"acts,omitempty" type:"Repeated"`
  // {"en":"Attack type, array. [SQL Injection: WAF_SQLI,
  // X-Site Scripting: WAF_XSS,
  // File Inclusion: WAF_RFI,
  // File Uploading: WAF_FILE_UPLOAD,
  // Command Injection: WAF_CMDI,
  // Directory Traversal: WAF_DIR_TRAVERSAL,
  // 3rd Party Component Exploit: WAF_THIRDCOMP,
  // XPATH/LDAP/SSI Injection: WAF_XPATH_LDAP_SSI,
  // Malicious Scanning: WAF_SCANNER,
  // Webshell Uploading: WAF_SHELL_BACKDOOR,
  // Unauthorized Download: WAF_FILE_DOWNLOAD,
  // Illegal Access(RFC): WAF_ILLEGAL,
  // Illegal HTTP Method(RFC): WAF_INVALID_METHOD,
  // Illegal HTTP Version(RFC): WAF_INVALID_PROTOCOL,
  // HTTP Buffer Overflow(RFC): WAF_BUF_OVERFLOW,
  // Server Sensitive Info Leakage[remarks:combination]: WAF_SERVER_ERROR_LEAKAGE,WAF_SENS_DATA_LEAKAGE,
  // XML Injection: WAF_XXE,
  // Custom Rules: WAF_CUSTOM_RULE,
  // Cookie Protection: WAF_COOKIE_PROTECT,
  // CSRF: WAF_CSRF,
  // Adware Insertion Prevention: WAF_UAD,
  // Webshell Access Detection: WAF_WEBSHELL,
  // Threat Intelligence: WAF_THREAT_INTELLIGENCE,
  // Credential Stuffing: WAF_HIT_LIB,
  // Server-side Request Forge: WAF_SSRF,
  // Malicious Dig: WAF_COIN,
  // Access Control[remarks:combination]: WAF_BLACK_IP,WAF_BLACK_URL,WAF_BLACK_UA,WAF_BLACK_HEADER,WAF_ACCESS_CONTROL,
  // Rate Limiting: WAF_FORCE_CRACKING,
  // Response Code Rate Limiting: WAF_BLACK_STATUS,
  // IP Repeated Violations: WAF_DYNAMIC_BLACK_IP]", "zh_CN":"攻击类型，数组。[SQL注入：WAF_SQLI,
  // XSS跨站：WAF_XSS,
  // 文件包含：WAF_RFI,
  // 文件上传：WAF_FILE_UPLOAD,
  // 命令注入：WAF_CMDI,
  // 目录遍历：WAF_DIR_TRAVERSAL,
  // 第三方组件漏洞：WAF_THIRDCOMP,
  // XPATH/LDAP/SSI注入：WAF_XPATH_LDAP_SSI,
  // 扫描器：WAF_SCANNER,
  // 木马后门：WAF_SHELL_BACKDOOR,
  // 非法下载：WAF_FILE_DOWNLOAD,
  // 非法请求(合规)：WAF_ILLEGAL,
  // 非法请求方法(合规)：WAF_INVALID_METHOD,
  // 非法请求协议(合规)：WAF_INVALID_PROTOCOL,
  // 缓冲区溢出(合规)：WAF_BUF_OVERFLOW,
  // 服务器信息泄露[备:多个组合]：WAF_SERVER_ERROR_LEAKAGE,WAF_SENS_DATA_LEAKAGE,
  // XML注入：WAF_XXE,
  // 自定义规则：WAF_CUSTOM_RULE,
  // Cookie防护：WAF_COOKIE_PROTECT,
  // CSRF：WAF_CSRF,
  // 恶意广告检测：WAF_UAD,
  // 后门识别：WAF_WEBSHELL,
  // 威胁情报：WAF_THREAT_INTELLIGENCE,
  // 撞库：WAF_HIT_LIB,
  // 服务端请求伪造：WAF_SSRF,
  // 恶意挖矿：WAF_COIN,
  // 访问控制[备:多个组合]：WAF_BLACK_IP,WAF_BLACK_URL,WAF_BLACK_UA,WAF_BLACK_HEADER,WAF_ACCESS_CONTROL,
  // 频率限制：WAF_FORCE_CRACKING,
  // 状态码限速：WAF_BLACK_STATUS,
  // 攻击IP惩罚：WAF_DYNAMIC_BLACK_IP]"}
  AttackTypes []*string `json:"attackTypes,omitempty" xml:"attackTypes,omitempty" type:"Repeated"`
  // {"en":"Client IP.", "zh_CN":"客户端IP。"}
  Ips []*string `json:"ips,omitempty" xml:"ips,omitempty" type:"Repeated"`
  // {"en":"The logic of ip match conditions, default value 3.
  //     1: Equal
  //     2: Not Equal
  //     3: Included
  //     4: Excluded", "zh_CN":"IP条件匹配逻辑，默认值3。
  //     1：等于
  //     2：不等于
  //     3：包含
  //     4：不包含"}
  IpCondition *int `json:"ipCondition,omitempty" xml:"ipCondition,omitempty"`
  // {"en":"The geographical location of IP.", "zh_CN":"IP地理位置。"}
  IpLocation *string `json:"ipLocation,omitempty" xml:"ipLocation,omitempty"`
  // {"en":"Request Path.", "zh_CN":"请求路径。"}
  Urls []*string `json:"urls,omitempty" xml:"urls,omitempty" type:"Repeated"`
  // {"en":"The logic of path match conditions, default value 3.
  //     1: Equal
  //     2: Not Equal
  //     3: Included
  //     4: Excluded", "zh_CN":"路径条件匹配逻辑，默认值3。
  //     1：等于
  //     2：不等于
  //     3：包含
  //     4：不包含"}
  UrlCondition *int `json:"urlCondition,omitempty" xml:"urlCondition,omitempty"`
  // {"en":"HTTP request header: Referer.", "zh_CN":"请求Referer。"}
  Referer *string `json:"referer,omitempty" xml:"referer,omitempty"`
  // {"en":"The logic of referer match conditions, default value 3.
  //     1: Equal
  //     2: Not Equal
  //     3: Included
  //     4: Excluded", "zh_CN":"Referer条件匹配逻辑，默认值3。
  //     1：等于
  //     2：不等于
  //     3：包含
  //     4：不包含"}
  RefererCondition *int `json:"refererCondition,omitempty" xml:"refererCondition,omitempty"`
  // {"en":"HTTP request header: User-Agent.", "zh_CN":"请求User-Agent。"}
  UserAgent *string `json:"userAgent,omitempty" xml:"userAgent,omitempty"`
  // {"en":"The logic of user-agent match conditions, default value 3.
  //     1: Equal
  //     2: Not Equal
  //     3: Included
  //     4: Excluded", "zh_CN":"User-Agent条件匹配逻辑，默认值3。
  //     1：等于
  //     2：不等于
  //     3：包含
  //     4：不包含"}
  UserAgentCondition *int `json:"userAgentCondition,omitempty" xml:"userAgentCondition,omitempty"`
  // {"en":"HTTP version, such as:[HTTP/1.1].", "zh_CN":"HTTP版本，比如:[HTTP/1.1]。"}
  HttpVerisions []*string `json:"httpVerisions,omitempty" xml:"httpVerisions,omitempty" type:"Repeated"`
  // {"en":"HTTP response status code.", "zh_CN":"HTTP响应状态码。"}
  StatusCodes []*string `json:"statusCodes,omitempty" xml:"statusCodes,omitempty" type:"Repeated"`
  // {"en":"The logic of status code match conditions, default value 3.
  //     1: Equal
  //     2: Not Equal
  //     3: Included
  //     4: Excluded", "zh_CN":"HTTP状态码条件匹配逻辑，默认值3。
  //     1：等于
  //     2：不等于
  //     3：包含
  //     4：不包含"}
  StatusCodeCondition *int `json:"statusCodeCondition,omitempty" xml:"statusCodeCondition,omitempty"`
  // {"en":"Event ID.", "zh_CN":"事件ID。"}
  EventId *string `json:"eventId,omitempty" xml:"eventId,omitempty"`
  // {"en":"Rule ID.", "zh_CN":"规则ID。"}
  RuleIds []*string `json:"ruleIds,omitempty" xml:"ruleIds,omitempty" type:"Repeated"`
  // {"en":"The logic of rule id match conditions, default value 3.
  //     1: Equal
  //     2: Not Equal
  //     3: Included
  //     4: Excluded", "zh_CN":"规则ID条件匹配逻辑，默认值3。
  //     1：等于
  //     2：不等于
  //     3：包含
  //     4：不包含"}
  RuleIdCondition *int `json:"ruleIdCondition,omitempty" xml:"ruleIdCondition,omitempty"`
  // {"en":"Rule name.", "zh_CN":"规则名称。"}
  RuleName *string `json:"ruleName,omitempty" xml:"ruleName,omitempty"`
  // {"en":"The language of response data, default value: cn.
  //     cn: Chinese
  //     en: English", "zh_CN":"返回内容的语言版本，默认值：cn。
  //     cn：中文
  //     en：英文"}
  Lang *string `json:"lang,omitempty" xml:"lang,omitempty"`
  // {"en":"Current page.", "zh_CN":"当前页码。"}
  PageNum *int `json:"pageNum,omitempty" xml:"pageNum,omitempty" require:"true"`
  // {"en":"Records per page, max 1000 records.", "zh_CN":"每页日志条数, 最大1000。"}
  PageSize *int `json:"pageSize,omitempty" xml:"pageSize,omitempty" require:"true"`
}

func (s GetWAFAttackLogRequest) String() string {
  return tea.Prettify(s)
}

func (s GetWAFAttackLogRequest) GoString() string {
  return s.String()
}

func (s *GetWAFAttackLogRequest) SetDomains(v []*string) *GetWAFAttackLogRequest {
  s.Domains = v
  return s
}

func (s *GetWAFAttackLogRequest) SetStartTime(v string) *GetWAFAttackLogRequest {
  s.StartTime = &v
  return s
}

func (s *GetWAFAttackLogRequest) SetEndTime(v string) *GetWAFAttackLogRequest {
  s.EndTime = &v
  return s
}

func (s *GetWAFAttackLogRequest) SetTimeZone(v string) *GetWAFAttackLogRequest {
  s.TimeZone = &v
  return s
}

func (s *GetWAFAttackLogRequest) SetActs(v []*string) *GetWAFAttackLogRequest {
  s.Acts = v
  return s
}

func (s *GetWAFAttackLogRequest) SetAttackTypes(v []*string) *GetWAFAttackLogRequest {
  s.AttackTypes = v
  return s
}

func (s *GetWAFAttackLogRequest) SetIps(v []*string) *GetWAFAttackLogRequest {
  s.Ips = v
  return s
}

func (s *GetWAFAttackLogRequest) SetIpCondition(v int) *GetWAFAttackLogRequest {
  s.IpCondition = &v
  return s
}

func (s *GetWAFAttackLogRequest) SetIpLocation(v string) *GetWAFAttackLogRequest {
  s.IpLocation = &v
  return s
}

func (s *GetWAFAttackLogRequest) SetUrls(v []*string) *GetWAFAttackLogRequest {
  s.Urls = v
  return s
}

func (s *GetWAFAttackLogRequest) SetUrlCondition(v int) *GetWAFAttackLogRequest {
  s.UrlCondition = &v
  return s
}

func (s *GetWAFAttackLogRequest) SetReferer(v string) *GetWAFAttackLogRequest {
  s.Referer = &v
  return s
}

func (s *GetWAFAttackLogRequest) SetRefererCondition(v int) *GetWAFAttackLogRequest {
  s.RefererCondition = &v
  return s
}

func (s *GetWAFAttackLogRequest) SetUserAgent(v string) *GetWAFAttackLogRequest {
  s.UserAgent = &v
  return s
}

func (s *GetWAFAttackLogRequest) SetUserAgentCondition(v int) *GetWAFAttackLogRequest {
  s.UserAgentCondition = &v
  return s
}

func (s *GetWAFAttackLogRequest) SetHttpVerisions(v []*string) *GetWAFAttackLogRequest {
  s.HttpVerisions = v
  return s
}

func (s *GetWAFAttackLogRequest) SetStatusCodes(v []*string) *GetWAFAttackLogRequest {
  s.StatusCodes = v
  return s
}

func (s *GetWAFAttackLogRequest) SetStatusCodeCondition(v int) *GetWAFAttackLogRequest {
  s.StatusCodeCondition = &v
  return s
}

func (s *GetWAFAttackLogRequest) SetEventId(v string) *GetWAFAttackLogRequest {
  s.EventId = &v
  return s
}

func (s *GetWAFAttackLogRequest) SetRuleIds(v []*string) *GetWAFAttackLogRequest {
  s.RuleIds = v
  return s
}

func (s *GetWAFAttackLogRequest) SetRuleIdCondition(v int) *GetWAFAttackLogRequest {
  s.RuleIdCondition = &v
  return s
}

func (s *GetWAFAttackLogRequest) SetRuleName(v string) *GetWAFAttackLogRequest {
  s.RuleName = &v
  return s
}

func (s *GetWAFAttackLogRequest) SetLang(v string) *GetWAFAttackLogRequest {
  s.Lang = &v
  return s
}

func (s *GetWAFAttackLogRequest) SetPageNum(v int) *GetWAFAttackLogRequest {
  s.PageNum = &v
  return s
}

func (s *GetWAFAttackLogRequest) SetPageSize(v int) *GetWAFAttackLogRequest {
  s.PageSize = &v
  return s
}

type GetWAFAttackLogResponse struct {
  // {"en":"Return 0 means success, please see <Error code> to check other status code.", "zh_CN":"0状态码表示请求成功，其他状态码说明请参见《错误码》。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Error message or Success.", "zh_CN":"错误信息或Success。"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Return data.", "zh_CN":"返回值。"}
  Data *GetWAFAttackLogLogList `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s GetWAFAttackLogResponse) String() string {
  return tea.Prettify(s)
}

func (s GetWAFAttackLogResponse) GoString() string {
  return s.String()
}

func (s *GetWAFAttackLogResponse) SetCode(v string) *GetWAFAttackLogResponse {
  s.Code = &v
  return s
}

func (s *GetWAFAttackLogResponse) SetMessage(v string) *GetWAFAttackLogResponse {
  s.Message = &v
  return s
}

func (s *GetWAFAttackLogResponse) SetData(v *GetWAFAttackLogLogList) *GetWAFAttackLogResponse {
  s.Data = v
  return s
}

type GetWAFAttackLogLogList struct {
  // {'en':'Current page.', 'zh_CN':'当前页码。'}
  CurrentPage *int `json:"currentPage,omitempty" xml:"currentPage,omitempty" require:"true"`
  // {'en':'Records per page.', 'zh_CN':'每页日志条数。'}
  PageSize *int `json:"pageSize,omitempty" xml:"pageSize,omitempty" require:"true"`
  // {'en':'Total number of records.', 'zh_CN':'总记录数量。'}
  TotalRecords *int `json:"totalRecords,omitempty" xml:"totalRecords,omitempty" require:"true"`
  // {'en':'Total pages.', 'zh_CN':'合计页数。'}
  TotalPages *int `json:"totalPages,omitempty" xml:"totalPages,omitempty" require:"true"`
  // {'en':'GetWAFAttackLogLog list.', 'zh_CN':'日志列表。'}
  GetWAFAttackLogLog []*GetWAFAttackLogLog `json:"log,omitempty" xml:"log,omitempty" require:"true" type:"Repeated"`
}

func (s GetWAFAttackLogLogList) String() string {
  return tea.Prettify(s)
}

func (s GetWAFAttackLogLogList) GoString() string {
  return s.String()
}

func (s *GetWAFAttackLogLogList) SetCurrentPage(v int) *GetWAFAttackLogLogList {
  s.CurrentPage = &v
  return s
}

func (s *GetWAFAttackLogLogList) SetPageSize(v int) *GetWAFAttackLogLogList {
  s.PageSize = &v
  return s
}

func (s *GetWAFAttackLogLogList) SetTotalRecords(v int) *GetWAFAttackLogLogList {
  s.TotalRecords = &v
  return s
}

func (s *GetWAFAttackLogLogList) SetTotalPages(v int) *GetWAFAttackLogLogList {
  s.TotalPages = &v
  return s
}

func (s *GetWAFAttackLogLogList) SetLog(v []*GetWAFAttackLogLog) *GetWAFAttackLogLogList {
  s.GetWAFAttackLogLog = v
  return s
}

type GetWAFAttackLogLog struct {
  // {'en':'Request referer.', 'zh_CN':'请求Referer。'}
  Referer *string `json:"referer,omitempty" xml:"referer,omitempty" require:"true"`
  // {'en':'HTTP request method.', 'zh_CN':'HTTP请求方式。'}
  Method *string `json:"method,omitempty" xml:"method,omitempty" require:"true"`
  // {'en':'HTTP response status code.', 'zh_CN':'HTTP响应状态码。'}
  ResponseCode *string `json:"responseCode,omitempty" xml:"responseCode,omitempty" require:"true"`
  // {'en':'Attack type.', 'zh_CN':'攻击类型。'}
  AttackType *string `json:"attackType,omitempty" xml:"attackType,omitempty" require:"true"`
  // {'en':'Rule name.', 'zh_CN':'规则名称。'}
  RuleName *string `json:"ruleName,omitempty" xml:"ruleName,omitempty" require:"true"`
  // {'en':'Client IP.', 'zh_CN':'客户端IP。'}
  Ip *string `json:"ip,omitempty" xml:"ip,omitempty" require:"true"`
  // {'en':'Event ID.', 'zh_CN':'事件ID。'}
  Uuid *string `json:"uuid,omitempty" xml:"uuid,omitempty" require:"true"`
  // {'en':'Request Path.', 'zh_CN':'请求路径。'}
  Url *string `json:"url,omitempty" xml:"url,omitempty" require:"true"`
  // {'en':'Client ID.', 'zh_CN':'客户端ID。'}
  ClientId *string `json:"clientId,omitempty" xml:"clientId,omitempty" require:"true"`
  // {'en':'Detail information.', 'zh_CN':'详细信息。'}
  Content *string `json:"content,omitempty" xml:"content,omitempty" require:"true"`
  // {'en':'Rule ID.', 'zh_CN':'规则ID。'}
  RuleId *string `json:"ruleId,omitempty" xml:"ruleId,omitempty" require:"true"`
  // {'en':'Event ID.', 'zh_CN':'事件ID。'}
  EventId *string `json:"eventId,omitempty" xml:"eventId,omitempty" require:"true"`
  // {'en':'HTTP version.', 'zh_CN':'HTTP版本。'}
  HttpVersion *string `json:"httpVersion,omitempty" xml:"httpVersion,omitempty" require:"true"`
  // {'en':'Match area.', 'zh_CN':'匹配区域。'}
  Zone *string `json:"zone,omitempty" xml:"zone,omitempty" require:"true"`
  // {'en':'Domain.', 'zh_CN':'域名。'}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {'en':'Action.
  //         1: Block
  //         2: GetWAFAttackLogLog', 'zh_CN':'处理动作。
  //         1：拦截
  //         2：监控'}
  Action *string `json:"action,omitempty" xml:"action,omitempty" require:"true"`
  // {'en':'The geographical location of IP.', 'zh_CN':'IP地理位置。'}
  Location *string `json:"location,omitempty" xml:"location,omitempty" require:"true"`
  // {'en':'Time.', 'zh_CN':'时间。'}
  Time *string `json:"time,omitempty" xml:"time,omitempty" require:"true"`
  // {'en':'Request User-Agent.', 'zh_CN':'请求User-Agent。'}
  UserAgent *string `json:"userAgent,omitempty" xml:"userAgent,omitempty" require:"true"`
  // {'en':'Domain.', 'zh_CN':'域名。'}
  DetailDomain *string `json:"detailDomain,omitempty" xml:"detailDomain,omitempty" require:"true"`
}

func (s GetWAFAttackLogLog) String() string {
  return tea.Prettify(s)
}

func (s GetWAFAttackLogLog) GoString() string {
  return s.String()
}

func (s *GetWAFAttackLogLog) SetReferer(v string) *GetWAFAttackLogLog {
  s.Referer = &v
  return s
}

func (s *GetWAFAttackLogLog) SetMethod(v string) *GetWAFAttackLogLog {
  s.Method = &v
  return s
}

func (s *GetWAFAttackLogLog) SetResponseCode(v string) *GetWAFAttackLogLog {
  s.ResponseCode = &v
  return s
}

func (s *GetWAFAttackLogLog) SetAttackType(v string) *GetWAFAttackLogLog {
  s.AttackType = &v
  return s
}

func (s *GetWAFAttackLogLog) SetRuleName(v string) *GetWAFAttackLogLog {
  s.RuleName = &v
  return s
}

func (s *GetWAFAttackLogLog) SetIp(v string) *GetWAFAttackLogLog {
  s.Ip = &v
  return s
}

func (s *GetWAFAttackLogLog) SetUuid(v string) *GetWAFAttackLogLog {
  s.Uuid = &v
  return s
}

func (s *GetWAFAttackLogLog) SetUrl(v string) *GetWAFAttackLogLog {
  s.Url = &v
  return s
}

func (s *GetWAFAttackLogLog) SetClientId(v string) *GetWAFAttackLogLog {
  s.ClientId = &v
  return s
}

func (s *GetWAFAttackLogLog) SetContent(v string) *GetWAFAttackLogLog {
  s.Content = &v
  return s
}

func (s *GetWAFAttackLogLog) SetRuleId(v string) *GetWAFAttackLogLog {
  s.RuleId = &v
  return s
}

func (s *GetWAFAttackLogLog) SetEventId(v string) *GetWAFAttackLogLog {
  s.EventId = &v
  return s
}

func (s *GetWAFAttackLogLog) SetHttpVersion(v string) *GetWAFAttackLogLog {
  s.HttpVersion = &v
  return s
}

func (s *GetWAFAttackLogLog) SetZone(v string) *GetWAFAttackLogLog {
  s.Zone = &v
  return s
}

func (s *GetWAFAttackLogLog) SetDomain(v string) *GetWAFAttackLogLog {
  s.Domain = &v
  return s
}

func (s *GetWAFAttackLogLog) SetAction(v string) *GetWAFAttackLogLog {
  s.Action = &v
  return s
}

func (s *GetWAFAttackLogLog) SetLocation(v string) *GetWAFAttackLogLog {
  s.Location = &v
  return s
}

func (s *GetWAFAttackLogLog) SetTime(v string) *GetWAFAttackLogLog {
  s.Time = &v
  return s
}

func (s *GetWAFAttackLogLog) SetUserAgent(v string) *GetWAFAttackLogLog {
  s.UserAgent = &v
  return s
}

func (s *GetWAFAttackLogLog) SetDetailDomain(v string) *GetWAFAttackLogLog {
  s.DetailDomain = &v
  return s
}

type GetWAFAttackLogPaths struct {
}

func (s GetWAFAttackLogPaths) String() string {
  return tea.Prettify(s)
}

func (s GetWAFAttackLogPaths) GoString() string {
  return s.String()
}

type GetWAFAttackLogParameters struct {
}

func (s GetWAFAttackLogParameters) String() string {
  return tea.Prettify(s)
}

func (s GetWAFAttackLogParameters) GoString() string {
  return s.String()
}

type GetWAFAttackLogRequestHeader struct {
}

func (s GetWAFAttackLogRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetWAFAttackLogRequestHeader) GoString() string {
  return s.String()
}

type GetWAFAttackLogResponseHeader struct {
}

func (s GetWAFAttackLogResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetWAFAttackLogResponseHeader) GoString() string {
  return s.String()
}




