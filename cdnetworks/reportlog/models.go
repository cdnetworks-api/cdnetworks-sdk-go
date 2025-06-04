package reportlog

// This file is auto-generated, don't edit it. Thanks.
import (
  "github.com/alibabacloud-go/tea/tea"
)

type GetRequestsByProcessTimeRequest struct {
}

func (s GetRequestsByProcessTimeRequest) String() string {
  return tea.Prettify(s)
}

func (s GetRequestsByProcessTimeRequest) GoString() string {
  return s.String()
}

type GetRequestsByProcessTimeResponse struct {
  // {"en":"return data", "zh_CN":"返回值"}
  Data []*GetRequestsByProcessTimeResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
  // {"en":"code,success is 0", "zh_CN":"状态码,成功为0"}
  Code *int `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"error message or Success", "zh_CN":"错误信息或Success"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s GetRequestsByProcessTimeResponse) String() string {
  return tea.Prettify(s)
}

func (s GetRequestsByProcessTimeResponse) GoString() string {
  return s.String()
}

func (s *GetRequestsByProcessTimeResponse) SetData(v []*GetRequestsByProcessTimeResponseData) *GetRequestsByProcessTimeResponse {
  s.Data = v
  return s
}

func (s *GetRequestsByProcessTimeResponse) SetCode(v int) *GetRequestsByProcessTimeResponse {
  s.Code = &v
  return s
}

func (s *GetRequestsByProcessTimeResponse) SetMessage(v string) *GetRequestsByProcessTimeResponse {
  s.Message = &v
  return s
}

type GetRequestsByProcessTimeResponseData struct     {
  // {"en":"timestamp (accurate to the second)", "zh_CN":"时间戳，精确到秒"}
  Ts *int `json:"ts,omitempty" xml:"ts,omitempty" require:"true"`
  // {"en":"Counts less than 1ms", "zh_CN":"小于1毫秒的数量统计"}
  Lt1ms *int `json:"lt1ms,omitempty" xml:"lt1ms,omitempty" require:"true"`
  // {"en":"Counts greater than 1ms and less than 10ms", "zh_CN":"大于1ms小于10ms的数量统计"}
  Ge1ms *int `json:"ge1ms,omitempty" xml:"ge1ms,omitempty" require:"true"`
  // {"en":"Counts greater than 10ms", "zh_CN":"大于10ms的数量统计"}
  Ge10ms *int `json:"ge10ms,omitempty" xml:"ge10ms,omitempty" require:"true"`
}

func (s GetRequestsByProcessTimeResponseData) String() string {
  return tea.Prettify(s)
}

func (s GetRequestsByProcessTimeResponseData) GoString() string {
  return s.String()
}

func (s *GetRequestsByProcessTimeResponseData) SetTs(v int) *GetRequestsByProcessTimeResponseData {
  s.Ts = &v
  return s
}

func (s *GetRequestsByProcessTimeResponseData) SetLt1ms(v int) *GetRequestsByProcessTimeResponseData {
  s.Lt1ms = &v
  return s
}

func (s *GetRequestsByProcessTimeResponseData) SetGe1ms(v int) *GetRequestsByProcessTimeResponseData {
  s.Ge1ms = &v
  return s
}

func (s *GetRequestsByProcessTimeResponseData) SetGe10ms(v int) *GetRequestsByProcessTimeResponseData {
  s.Ge10ms = &v
  return s
}

type GetRequestsByProcessTimePaths struct {
}

func (s GetRequestsByProcessTimePaths) String() string {
  return tea.Prettify(s)
}

func (s GetRequestsByProcessTimePaths) GoString() string {
  return s.String()
}

type GetRequestsByProcessTimeParameters struct {
  // {"en":"from time(timestamp) (accurate to the second)", "zh_CN":"起始时间，timestamp类型，精确到秒"}
  From *int `json:"from,omitempty" xml:"from,omitempty" require:"true"`
  // {"en":"to time(timestamp) (accurate to the second)", "zh_CN":"结束时间，timestamp类型，精确到秒"}
  To *int `json:"to,omitempty" xml:"to,omitempty" require:"true"`
  // {"en":"interval time, default is 'hourly'. Support interval:oneminute: 1 min.(Range <= 1 day);fiveminutes: 5 min(Range <= 7 days); hourly: 1 hr; daily: 1 day; monthly: 1 month; all: Without interval(Combine data into one, Range>=1hr)", "zh_CN":"间隔时间，默认为hourly。支持的类型有:oneminute: 1 min.(Range <= 1 day);fiveminutes: 5 min(Range <= 7 days); hourly: 1 hr; daily: 1 day; monthly: 1 month; all: Without interval(合并数据成一条返回，Range>=1hr)"}
  Interval *string `json:"interval,omitempty" xml:"interval,omitempty"`
  // {"en":"zone names, separeated by ',', case insensitive", "zh_CN":"zone名称字符串，英文逗号分隔，不区分大小写"}
  Zone *string `json:"zone,omitempty" xml:"zone,omitempty"`
  // {"en":"timezone is necessary when interval=daily/monthly/all, the default timezone is 0 (UTC+0)", "zh_CN":"时区，当interval为daily/monthly/all时需要判断时区。"}
  Timezone *int `json:"timezone,omitempty" xml:"timezone,omitempty"`
}

func (s GetRequestsByProcessTimeParameters) String() string {
  return tea.Prettify(s)
}

func (s GetRequestsByProcessTimeParameters) GoString() string {
  return s.String()
}

func (s *GetRequestsByProcessTimeParameters) SetFrom(v int) *GetRequestsByProcessTimeParameters {
  s.From = &v
  return s
}

func (s *GetRequestsByProcessTimeParameters) SetTo(v int) *GetRequestsByProcessTimeParameters {
  s.To = &v
  return s
}

func (s *GetRequestsByProcessTimeParameters) SetInterval(v string) *GetRequestsByProcessTimeParameters {
  s.Interval = &v
  return s
}

func (s *GetRequestsByProcessTimeParameters) SetZone(v string) *GetRequestsByProcessTimeParameters {
  s.Zone = &v
  return s
}

func (s *GetRequestsByProcessTimeParameters) SetTimezone(v int) *GetRequestsByProcessTimeParameters {
  s.Timezone = &v
  return s
}

type GetRequestsByProcessTimeRequestHeader struct {
}

func (s GetRequestsByProcessTimeRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetRequestsByProcessTimeRequestHeader) GoString() string {
  return s.String()
}

type GetRequestsByProcessTimeResponseHeader struct {
}

func (s GetRequestsByProcessTimeResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetRequestsByProcessTimeResponseHeader) GoString() string {
  return s.String()
}




type GetRequestsByResponseTypeRequest struct {
}

func (s GetRequestsByResponseTypeRequest) String() string {
  return tea.Prettify(s)
}

func (s GetRequestsByResponseTypeRequest) GoString() string {
  return s.String()
}

type GetRequestsByResponseTypeResponse struct {
  // {"en":"return data", "zh_CN":"返回值"}
  Data []*GetRequestsByResponseTypeResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
  // {"en":"code,success is 0", "zh_CN":"状态码,成功为0"}
  Code *int `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"error message or Success", "zh_CN":"错误信息或Success"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s GetRequestsByResponseTypeResponse) String() string {
  return tea.Prettify(s)
}

func (s GetRequestsByResponseTypeResponse) GoString() string {
  return s.String()
}

func (s *GetRequestsByResponseTypeResponse) SetData(v []*GetRequestsByResponseTypeResponseData) *GetRequestsByResponseTypeResponse {
  s.Data = v
  return s
}

func (s *GetRequestsByResponseTypeResponse) SetCode(v int) *GetRequestsByResponseTypeResponse {
  s.Code = &v
  return s
}

func (s *GetRequestsByResponseTypeResponse) SetMessage(v string) *GetRequestsByResponseTypeResponse {
  s.Message = &v
  return s
}

type GetRequestsByResponseTypeResponseData struct     {
  // {"en":"timestamp (accurate to the second)", "zh_CN":"时间戳，精确到秒"}
  Ts *int `json:"ts,omitempty" xml:"ts,omitempty" require:"true"`
  // {"en":"success count", "zh_CN":"成功数量"}
  Success *int `json:"success,omitempty" xml:"success,omitempty" require:"true"`
  // {"en":"nx domain count", "zh_CN":"nxDomain数量"}
  SuccessNxDomain *int `json:"successNxDomain,omitempty" xml:"successNxDomain,omitempty" require:"true"`
  // {"en":"error count", "zh_CN":"失败数量"}
  Error *int `json:"error,omitempty" xml:"error,omitempty" require:"true"`
  // {"en":"error invalid query count", "zh_CN":"错误无效查询数量"}
  ErrorInvalidQuery *int `json:"errorInvalidQuery,omitempty" xml:"errorInvalidQuery,omitempty" require:"true"`
}

func (s GetRequestsByResponseTypeResponseData) String() string {
  return tea.Prettify(s)
}

func (s GetRequestsByResponseTypeResponseData) GoString() string {
  return s.String()
}

func (s *GetRequestsByResponseTypeResponseData) SetTs(v int) *GetRequestsByResponseTypeResponseData {
  s.Ts = &v
  return s
}

func (s *GetRequestsByResponseTypeResponseData) SetSuccess(v int) *GetRequestsByResponseTypeResponseData {
  s.Success = &v
  return s
}

func (s *GetRequestsByResponseTypeResponseData) SetSuccessNxDomain(v int) *GetRequestsByResponseTypeResponseData {
  s.SuccessNxDomain = &v
  return s
}

func (s *GetRequestsByResponseTypeResponseData) SetError(v int) *GetRequestsByResponseTypeResponseData {
  s.Error = &v
  return s
}

func (s *GetRequestsByResponseTypeResponseData) SetErrorInvalidQuery(v int) *GetRequestsByResponseTypeResponseData {
  s.ErrorInvalidQuery = &v
  return s
}

type GetRequestsByResponseTypePaths struct {
}

func (s GetRequestsByResponseTypePaths) String() string {
  return tea.Prettify(s)
}

func (s GetRequestsByResponseTypePaths) GoString() string {
  return s.String()
}

type GetRequestsByResponseTypeParameters struct {
  // {"en":"from time(timestamp) (accurate to the second)", "zh_CN":"起始时间，timestamp类型，精确到秒"}
  From *int `json:"from,omitempty" xml:"from,omitempty" require:"true"`
  // {"en":"to time(timestamp) (accurate to the second)", "zh_CN":"结束时间，timestamp类型，精确到秒"}
  To *int `json:"to,omitempty" xml:"to,omitempty" require:"true"`
  // {"en":"interval time, default is 'hourly'. Support interval:oneminute: 1 min.(Range <= 1 day);fiveminutes: 5 min(Range <= 7 days); hourly: 1 hr; daily: 1 day; monthly: 1 month; all: Without interval(Combine data into one, Range>=1hr)", "zh_CN":"间隔时间，默认为hourly。支持的类型有:oneminute: 1 min.(Range <= 1 day);fiveminutes: 5 min(Range <= 7 days); hourly: 1 hr; daily: 1 day; monthly: 1 month; all: Without interval(合并数据成一条返回，Range>=1hr)"}
  Interval *string `json:"interval,omitempty" xml:"interval,omitempty"`
  // {"en":"zone names, separeated by ',', case insensitive", "zh_CN":"zone名称字符串，英文逗号分隔，不区分大小写"}
  Zone *string `json:"zone,omitempty" xml:"zone,omitempty"`
  // {"en":"timezone is necessary when interval=daily/monthly/all, the default timezone is 0 (UTC+0)", "zh_CN":"时区，当interval为daily/monthly/all时需要判断时区。"}
  Timezone *int `json:"timezone,omitempty" xml:"timezone,omitempty"`
}

func (s GetRequestsByResponseTypeParameters) String() string {
  return tea.Prettify(s)
}

func (s GetRequestsByResponseTypeParameters) GoString() string {
  return s.String()
}

func (s *GetRequestsByResponseTypeParameters) SetFrom(v int) *GetRequestsByResponseTypeParameters {
  s.From = &v
  return s
}

func (s *GetRequestsByResponseTypeParameters) SetTo(v int) *GetRequestsByResponseTypeParameters {
  s.To = &v
  return s
}

func (s *GetRequestsByResponseTypeParameters) SetInterval(v string) *GetRequestsByResponseTypeParameters {
  s.Interval = &v
  return s
}

func (s *GetRequestsByResponseTypeParameters) SetZone(v string) *GetRequestsByResponseTypeParameters {
  s.Zone = &v
  return s
}

func (s *GetRequestsByResponseTypeParameters) SetTimezone(v int) *GetRequestsByResponseTypeParameters {
  s.Timezone = &v
  return s
}

type GetRequestsByResponseTypeRequestHeader struct {
}

func (s GetRequestsByResponseTypeRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetRequestsByResponseTypeRequestHeader) GoString() string {
  return s.String()
}

type GetRequestsByResponseTypeResponseHeader struct {
}

func (s GetRequestsByResponseTypeResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetRequestsByResponseTypeResponseHeader) GoString() string {
  return s.String()
}




type GetRequestsByLocationDistributionRequest struct {
}

func (s GetRequestsByLocationDistributionRequest) String() string {
  return tea.Prettify(s)
}

func (s GetRequestsByLocationDistributionRequest) GoString() string {
  return s.String()
}

type GetRequestsByLocationDistributionResponse struct {
  // {"en":"return data", "zh_CN":"返回值"}
  Data []*GetRequestsByLocationDistributionResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
  // {"en":"code,success is 0", "zh_CN":"状态码,成功为0"}
  Code *int `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"error message or Success", "zh_CN":"错误信息或Success"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s GetRequestsByLocationDistributionResponse) String() string {
  return tea.Prettify(s)
}

func (s GetRequestsByLocationDistributionResponse) GoString() string {
  return s.String()
}

func (s *GetRequestsByLocationDistributionResponse) SetData(v []*GetRequestsByLocationDistributionResponseData) *GetRequestsByLocationDistributionResponse {
  s.Data = v
  return s
}

func (s *GetRequestsByLocationDistributionResponse) SetCode(v int) *GetRequestsByLocationDistributionResponse {
  s.Code = &v
  return s
}

func (s *GetRequestsByLocationDistributionResponse) SetMessage(v string) *GetRequestsByLocationDistributionResponse {
  s.Message = &v
  return s
}

type GetRequestsByLocationDistributionResponseData struct     {
  // {"en":"timestamp (accurate to the second)", "zh_CN":"时间戳，精确到秒"}
  Ts *int `json:"ts,omitempty" xml:"ts,omitempty" require:"true"`
  // {"en":"data", "zh_CN":"值"}
  Results []*GetRequestsByLocationDistributionResponseDataResults `json:"results,omitempty" xml:"results,omitempty" require:"true" type:"Repeated"`
}

func (s GetRequestsByLocationDistributionResponseData) String() string {
  return tea.Prettify(s)
}

func (s GetRequestsByLocationDistributionResponseData) GoString() string {
  return s.String()
}

func (s *GetRequestsByLocationDistributionResponseData) SetTs(v int) *GetRequestsByLocationDistributionResponseData {
  s.Ts = &v
  return s
}

func (s *GetRequestsByLocationDistributionResponseData) SetResults(v []*GetRequestsByLocationDistributionResponseDataResults) *GetRequestsByLocationDistributionResponseData {
  s.Results = v
  return s
}

type GetRequestsByLocationDistributionResponseDataResults struct     {
  // {"en":"fsn, can get name map with getRegionList", "zh_CN":"fsn值，可以从接口<获取区域列表>获取对应的映射"}
  Region *string `json:"region,omitempty" xml:"region,omitempty" require:"true"`
  // {"en":"count", "zh_CN":"数量"}
  Count *int `json:"count,omitempty" xml:"count,omitempty" require:"true"`
}

func (s GetRequestsByLocationDistributionResponseDataResults) String() string {
  return tea.Prettify(s)
}

func (s GetRequestsByLocationDistributionResponseDataResults) GoString() string {
  return s.String()
}

func (s *GetRequestsByLocationDistributionResponseDataResults) SetRegion(v string) *GetRequestsByLocationDistributionResponseDataResults {
  s.Region = &v
  return s
}

func (s *GetRequestsByLocationDistributionResponseDataResults) SetCount(v int) *GetRequestsByLocationDistributionResponseDataResults {
  s.Count = &v
  return s
}

type GetRequestsByLocationDistributionPaths struct {
}

func (s GetRequestsByLocationDistributionPaths) String() string {
  return tea.Prettify(s)
}

func (s GetRequestsByLocationDistributionPaths) GoString() string {
  return s.String()
}

type GetRequestsByLocationDistributionParameters struct {
  // {"en":"from time(timestamp) (accurate to the second)", "zh_CN":"起始时间，timestamp类型，精确到秒"}
  From *int `json:"from,omitempty" xml:"from,omitempty" require:"true"`
  // {"en":"to time(timestamp) (accurate to the second)", "zh_CN":"结束时间，timestamp类型，精确到秒"}
  To *int `json:"to,omitempty" xml:"to,omitempty" require:"true"`
  // {"en":"interval time, default is 'hourly'. Support interval:oneminute: 1 min.(Range <= 1 day);fiveminutes: 5 min(Range <= 7 days); hourly: 1 hr; daily: 1 day; monthly: 1 month; all: Without interval(Combine data into one, Range>=1hr)", "zh_CN":"间隔时间，默认为hourly。支持的类型有:oneminute: 1 min.(Range <= 1 day);fiveminutes: 5 min(Range <= 7 days); hourly: 1 hr; daily: 1 day; monthly: 1 month; all: Without interval(合并数据成一条返回，Range>=1hr)"}
  Interval *string `json:"interval,omitempty" xml:"interval,omitempty"`
  // {"en":"zone names, separeated by ',', case insensitive", "zh_CN":"zone名称字符串，英文逗号分隔，不区分大小写"}
  Zone *string `json:"zone,omitempty" xml:"zone,omitempty"`
  // {"en":"timezone is necessary when interval=daily/monthly/all, the default timezone is 0 (UTC+0)", "zh_CN":"时区，当interval为daily/monthly/all时需要判断时区。"}
  Timezone *int `json:"timezone,omitempty" xml:"timezone,omitempty"`
}

func (s GetRequestsByLocationDistributionParameters) String() string {
  return tea.Prettify(s)
}

func (s GetRequestsByLocationDistributionParameters) GoString() string {
  return s.String()
}

func (s *GetRequestsByLocationDistributionParameters) SetFrom(v int) *GetRequestsByLocationDistributionParameters {
  s.From = &v
  return s
}

func (s *GetRequestsByLocationDistributionParameters) SetTo(v int) *GetRequestsByLocationDistributionParameters {
  s.To = &v
  return s
}

func (s *GetRequestsByLocationDistributionParameters) SetInterval(v string) *GetRequestsByLocationDistributionParameters {
  s.Interval = &v
  return s
}

func (s *GetRequestsByLocationDistributionParameters) SetZone(v string) *GetRequestsByLocationDistributionParameters {
  s.Zone = &v
  return s
}

func (s *GetRequestsByLocationDistributionParameters) SetTimezone(v int) *GetRequestsByLocationDistributionParameters {
  s.Timezone = &v
  return s
}

type GetRequestsByLocationDistributionRequestHeader struct {
}

func (s GetRequestsByLocationDistributionRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetRequestsByLocationDistributionRequestHeader) GoString() string {
  return s.String()
}

type GetRequestsByLocationDistributionResponseHeader struct {
}

func (s GetRequestsByLocationDistributionResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetRequestsByLocationDistributionResponseHeader) GoString() string {
  return s.String()
}




type GetBotAttackIncidentLogDataRequest struct {
  // {"en":"Domain.Separate by ';'.", "zh_CN":"域名。多个以;隔开。"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty"`
  // {"en":"StartTime. Format: yyyy-MM-dd HH:mm:ss", "zh_CN":"开始时间。格式： yyyy-MM-dd HH:mm:ss"}
  StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty" require:"true"`
  // {"en":"EndTime. Format: yyyy-MM-dd HH:mm:ss", "zh_CN":"结束时间。格式： yyyy-MM-dd HH:mm:ss"}
  EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty" require:"true"`
  // {"en":"Time zone. Default 8, i.e.'GTM+8'.", "zh_CN":"时区，默认8，即“GTM+8”"}
  TimeZone *int `json:"timeZone,omitempty" xml:"timeZone,omitempty"`
  // {"en":"Language type. Default cn. 
  //  en:English 
  //  cn:Chinese ", "zh_CN":"语言类型。 默认cn 
  //  en：英文 
  //  cn：中文"}
  Lang *string `json:"lang,omitempty" xml:"lang,omitempty"`
  // {"en":"Attack Type,separate by ';'.", "zh_CN":"攻击类型。多个以;隔开。"}
  AttackType *string `json:"attackType,omitempty" xml:"attackType,omitempty"`
  // {"en":"Action. 
  //  1:Interception 
  //  2:Log 
  //  7:Flag 
  //  8:Captcha", "zh_CN":"处理动作。 
  //  1：拦截 
  //  2：监控 
  //  7：攻击标记 
  //  8：验证码"}
  Act *string `json:"act,omitempty" xml:"act,omitempty"`
  // {"en":"IP location.", "zh_CN":"IP地理位置。"}
  Location *string `json:"location,omitempty" xml:"location,omitempty"`
  // {"en":"Client IP.", "zh_CN":"客户端IP。"}
  Ip *string `json:"ip,omitempty" xml:"ip,omitempty"`
  // {"en":"URI.", "zh_CN":"URI。"}
  Url *string `json:"url,omitempty" xml:"url,omitempty"`
  // {"en":"Referer.", "zh_CN":"Referer。"}
  Referer *string `json:"referer,omitempty" xml:"referer,omitempty"`
  // {"en":"Status code.", "zh_CN":"状态码。"}
  StatusCode *string `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  // {"en":"User-Agent.", "zh_CN":"User-Agent。"}
  UserAgent *string `json:"userAgent,omitempty" xml:"userAgent,omitempty"`
  // {"en":"Event ID.", "zh_CN":"事件ID。"}
  Uuid *string `json:"uuid,omitempty" xml:"uuid,omitempty"`
  // {"en":"The number of entries displayed per page.Maximum limit 10,000.", "zh_CN":"每页显示的条目数。最大限制10,000。"}
  PageSize *int `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
  // {"en":"Current page number.", "zh_CN":"当前页码。"}
  CurrentPage *int `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
  // {"en":"Client ID.", "zh_CN":"客户端ID。"}
  ClientId *string `json:"clientId,omitempty" xml:"clientId,omitempty"`
  // {"en":"Browser Fingerprint.", "zh_CN":"浏览器指纹。"}
  BrowserFp *string `json:"browserFp,omitempty" xml:"browserFp,omitempty"`
  // {"en":"Rule name.", "zh_CN":"规则名。"}
  RuleName *string `json:"ruleName,omitempty" xml:"ruleName,omitempty"`
  // {"en":"Bot rule name.", "zh_CN":"Bot规则名。"}
  BotRuleName *string `json:"botRuleName,omitempty" xml:"botRuleName,omitempty"`
  // {"en":"Query criteria matching method of 'Client IP'.Default value: 3. 
  //  1:equal 
  //  2:not equal 
  //  3:Include 
  //  4:Not Include.", "zh_CN":"'客户端IP'查询条件匹配方式。默认值：3。 
  //  1：相等 
  //  2：不相等 
  //  3：包含 
  //  4：不包含"}
  IpCondition *int `json:"ipCondition,omitempty" xml:"ipCondition,omitempty"`
  // {"en":"Query criteria matching method of 'URI'. Default value: 3 
  //  1:equal 
  //  2:not equal 
  //  3:Include 
  //  4:Not Include.", "zh_CN":"'客户端URL'查询条件匹配方式。默认值：3。 
  //  1：相等 
  //  2：不相等 
  //  3：包含 
  //  4：不包含"}
  UrlCondition *int `json:"urlCondition,omitempty" xml:"urlCondition,omitempty"`
  // {"en":"Query criteria matching method of 'Referer'. Default value: 3 
  //  1:equal 
  //  2:not equal 
  //  3:Include 
  //  4:Not Include.", "zh_CN":"'客户端Rerfer'查询条件匹配方式。默认值：3。 
  //  1：相等 
  //  2：不相等 
  //  3：包含 
  //  4：不包含"}
  RefererCondition *int `json:"refererCondition,omitempty" xml:"refererCondition,omitempty"`
  // {"en":"Query criteria matching method of 'Response code'.Default value: 3 
  //  1:equal 
  //  2:not equal 
  //  3:Include 
  //  4:Not Include.", "zh_CN":"'客户端状态码'查询条件匹配方式。默认值：3。 
  //  1：相等 
  //  2：不相等 
  //  3：包含 
  //  4：不包含"}
  StatusCodeConditon *int `json:"statusCodeConditon,omitempty" xml:"statusCodeConditon,omitempty"`
  // {"en":"Query criteria matching method of 'User-Agent'.Default value: 3 
  //  1:equal 
  //  2:not equal 
  //  3:Include 
  //  4:Not Include. ", "zh_CN":"'客户端UA'查询条件匹配方式。默认值：3。 
  //  1：相等 
  //  2：不相等 
  //  3：包含 
  //  4：不包含"}
  UserAgentCondition *int `json:"userAgentCondition,omitempty" xml:"userAgentCondition,omitempty"`
}

func (s GetBotAttackIncidentLogDataRequest) String() string {
  return tea.Prettify(s)
}

func (s GetBotAttackIncidentLogDataRequest) GoString() string {
  return s.String()
}

func (s *GetBotAttackIncidentLogDataRequest) SetDomain(v string) *GetBotAttackIncidentLogDataRequest {
  s.Domain = &v
  return s
}

func (s *GetBotAttackIncidentLogDataRequest) SetStartTime(v string) *GetBotAttackIncidentLogDataRequest {
  s.StartTime = &v
  return s
}

func (s *GetBotAttackIncidentLogDataRequest) SetEndTime(v string) *GetBotAttackIncidentLogDataRequest {
  s.EndTime = &v
  return s
}

func (s *GetBotAttackIncidentLogDataRequest) SetTimeZone(v int) *GetBotAttackIncidentLogDataRequest {
  s.TimeZone = &v
  return s
}

func (s *GetBotAttackIncidentLogDataRequest) SetLang(v string) *GetBotAttackIncidentLogDataRequest {
  s.Lang = &v
  return s
}

func (s *GetBotAttackIncidentLogDataRequest) SetAttackType(v string) *GetBotAttackIncidentLogDataRequest {
  s.AttackType = &v
  return s
}

func (s *GetBotAttackIncidentLogDataRequest) SetAct(v string) *GetBotAttackIncidentLogDataRequest {
  s.Act = &v
  return s
}

func (s *GetBotAttackIncidentLogDataRequest) SetLocation(v string) *GetBotAttackIncidentLogDataRequest {
  s.Location = &v
  return s
}

func (s *GetBotAttackIncidentLogDataRequest) SetIp(v string) *GetBotAttackIncidentLogDataRequest {
  s.Ip = &v
  return s
}

func (s *GetBotAttackIncidentLogDataRequest) SetUrl(v string) *GetBotAttackIncidentLogDataRequest {
  s.Url = &v
  return s
}

func (s *GetBotAttackIncidentLogDataRequest) SetReferer(v string) *GetBotAttackIncidentLogDataRequest {
  s.Referer = &v
  return s
}

func (s *GetBotAttackIncidentLogDataRequest) SetStatusCode(v string) *GetBotAttackIncidentLogDataRequest {
  s.StatusCode = &v
  return s
}

func (s *GetBotAttackIncidentLogDataRequest) SetUserAgent(v string) *GetBotAttackIncidentLogDataRequest {
  s.UserAgent = &v
  return s
}

func (s *GetBotAttackIncidentLogDataRequest) SetUuid(v string) *GetBotAttackIncidentLogDataRequest {
  s.Uuid = &v
  return s
}

func (s *GetBotAttackIncidentLogDataRequest) SetPageSize(v int) *GetBotAttackIncidentLogDataRequest {
  s.PageSize = &v
  return s
}

func (s *GetBotAttackIncidentLogDataRequest) SetCurrentPage(v int) *GetBotAttackIncidentLogDataRequest {
  s.CurrentPage = &v
  return s
}

func (s *GetBotAttackIncidentLogDataRequest) SetClientId(v string) *GetBotAttackIncidentLogDataRequest {
  s.ClientId = &v
  return s
}

func (s *GetBotAttackIncidentLogDataRequest) SetBrowserFp(v string) *GetBotAttackIncidentLogDataRequest {
  s.BrowserFp = &v
  return s
}

func (s *GetBotAttackIncidentLogDataRequest) SetRuleName(v string) *GetBotAttackIncidentLogDataRequest {
  s.RuleName = &v
  return s
}

func (s *GetBotAttackIncidentLogDataRequest) SetBotRuleName(v string) *GetBotAttackIncidentLogDataRequest {
  s.BotRuleName = &v
  return s
}

func (s *GetBotAttackIncidentLogDataRequest) SetIpCondition(v int) *GetBotAttackIncidentLogDataRequest {
  s.IpCondition = &v
  return s
}

func (s *GetBotAttackIncidentLogDataRequest) SetUrlCondition(v int) *GetBotAttackIncidentLogDataRequest {
  s.UrlCondition = &v
  return s
}

func (s *GetBotAttackIncidentLogDataRequest) SetRefererCondition(v int) *GetBotAttackIncidentLogDataRequest {
  s.RefererCondition = &v
  return s
}

func (s *GetBotAttackIncidentLogDataRequest) SetStatusCodeConditon(v int) *GetBotAttackIncidentLogDataRequest {
  s.StatusCodeConditon = &v
  return s
}

func (s *GetBotAttackIncidentLogDataRequest) SetUserAgentCondition(v int) *GetBotAttackIncidentLogDataRequest {
  s.UserAgentCondition = &v
  return s
}

type GetBotAttackIncidentLogDataResponse struct {
  // {"en":"Status code, success is '200'.", "zh_CN":"状态码，成功为“200”。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Return message, success is 'Success'.", "zh_CN":"返回信息，成功为“Success”。"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Data returned.", "zh_CN":"返回数据"}
  Data *GetBotAttackIncidentLogDataResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
}

func (s GetBotAttackIncidentLogDataResponse) String() string {
  return tea.Prettify(s)
}

func (s GetBotAttackIncidentLogDataResponse) GoString() string {
  return s.String()
}

func (s *GetBotAttackIncidentLogDataResponse) SetCode(v string) *GetBotAttackIncidentLogDataResponse {
  s.Code = &v
  return s
}

func (s *GetBotAttackIncidentLogDataResponse) SetMessage(v string) *GetBotAttackIncidentLogDataResponse {
  s.Message = &v
  return s
}

func (s *GetBotAttackIncidentLogDataResponse) SetData(v *GetBotAttackIncidentLogDataResponseData) *GetBotAttackIncidentLogDataResponse {
  s.Data = v
  return s
}

type GetBotAttackIncidentLogDataResponseData struct {
  // {"en":"Rule name.", "zh_CN":"当前页码。"}
  CurrentPage *int `json:"currentPage,omitempty" xml:"currentPage,omitempty" require:"true"`
  // {"en":"Current page number.", "zh_CN":"首页页码。"}
  FirstPage *int `json:"firstPage,omitempty" xml:"firstPage,omitempty" require:"true"`
  // {"en":"last page number.", "zh_CN":"末页页码。"}
  LastPage *int `json:"lastPage,omitempty" xml:"lastPage,omitempty" require:"true"`
  // {"en":"The number of entries displayed per page.Maximum limit 10,000.", "zh_CN":"每页显示的条目数。最大限制10,000。"}
  PageSize *int `json:"pageSize,omitempty" xml:"pageSize,omitempty" require:"true"`
  // {"en":"Total entries.Maximum limit 1,000,000.", "zh_CN":"总条目数。最大限制1,000,000。"}
  TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty" require:"true"`
  // {"en":"Total page count.", "zh_CN":"总页数。"}
  TotalPageCount *int `json:"totalPageCount,omitempty" xml:"totalPageCount,omitempty" require:"true"`
  // {"en":"Data List", "zh_CN":"数据列表"}
  List []*GetBotAttackIncidentLogDataResponseDataList `json:"list,omitempty" xml:"list,omitempty" require:"true" type:"Repeated"`
}

func (s GetBotAttackIncidentLogDataResponseData) String() string {
  return tea.Prettify(s)
}

func (s GetBotAttackIncidentLogDataResponseData) GoString() string {
  return s.String()
}

func (s *GetBotAttackIncidentLogDataResponseData) SetCurrentPage(v int) *GetBotAttackIncidentLogDataResponseData {
  s.CurrentPage = &v
  return s
}

func (s *GetBotAttackIncidentLogDataResponseData) SetFirstPage(v int) *GetBotAttackIncidentLogDataResponseData {
  s.FirstPage = &v
  return s
}

func (s *GetBotAttackIncidentLogDataResponseData) SetLastPage(v int) *GetBotAttackIncidentLogDataResponseData {
  s.LastPage = &v
  return s
}

func (s *GetBotAttackIncidentLogDataResponseData) SetPageSize(v int) *GetBotAttackIncidentLogDataResponseData {
  s.PageSize = &v
  return s
}

func (s *GetBotAttackIncidentLogDataResponseData) SetTotalCount(v int64) *GetBotAttackIncidentLogDataResponseData {
  s.TotalCount = &v
  return s
}

func (s *GetBotAttackIncidentLogDataResponseData) SetTotalPageCount(v int) *GetBotAttackIncidentLogDataResponseData {
  s.TotalPageCount = &v
  return s
}

func (s *GetBotAttackIncidentLogDataResponseData) SetList(v []*GetBotAttackIncidentLogDataResponseDataList) *GetBotAttackIncidentLogDataResponseData {
  s.List = v
  return s
}

type GetBotAttackIncidentLogDataResponseDataList struct     {
  // {"en":"Referer.", "zh_CN":"Referer。"}
  Referer *string `json:"referer,omitempty" xml:"referer,omitempty" require:"true"`
  // {"en":"Browser Fingerprint.", "zh_CN":"浏览器指纹。"}
  Browser_fp *string `json:"browser_fp,omitempty" xml:"browser_fp,omitempty" require:"true"`
  // {"en":"Attack type.", "zh_CN":"攻击类型。"}
  Attack_type *string `json:"attack_type,omitempty" xml:"attack_type,omitempty" require:"true"`
  // {"en":"Rule name.", "zh_CN":"规则名称。"}
  Rule_name *string `json:"rule_name,omitempty" xml:"rule_name,omitempty" require:"true"`
  // {"en":"IP.", "zh_CN":"IP。"}
  Ip *string `json:"ip,omitempty" xml:"ip,omitempty" require:"true"`
  // {"en":"Event id.", "zh_CN":"事件ID。"}
  Uuid *string `json:"uuid,omitempty" xml:"uuid,omitempty" require:"true"`
  // {"en":"Version.", "zh_CN":"版本号。"}
  Version *string `json:"version,omitempty" xml:"version,omitempty" require:"true"`
  // {"en":"Client id.", "zh_CN":"客户端ID。"}
  Client_id *string `json:"client_id,omitempty" xml:"client_id,omitempty" require:"true"`
  // {"en":"URI.", "zh_CN":"URI。"}
  Url *string `json:"url,omitempty" xml:"url,omitempty" require:"true"`
  // {"en":"Block id.", "zh_CN":"Block id。"}
  Block_id *string `json:"block_id,omitempty" xml:"block_id,omitempty" require:"true"`
  // {"en":"Content.", "zh_CN":"内容。"}
  Content *string `json:"content,omitempty" xml:"content,omitempty" require:"true"`
  // {"en":"Request method.", "zh_CN":"请求方法。"}
  Mode *string `json:"mode,omitempty" xml:"mode,omitempty" require:"true"`
  // {"en":"Rule id.", "zh_CN":"规则id。"}
  Final_rule_id *int `json:"final_rule_id,omitempty" xml:"final_rule_id,omitempty" require:"true"`
  // {"en":"Event type.", "zh_CN":"事件类型。"}
  Event_type *string `json:"event_type,omitempty" xml:"event_type,omitempty" require:"true"`
  // {"en":"Processing action.", "zh_CN":"处理动作。"}
  Act *string `json:"act,omitempty" xml:"act,omitempty" require:"true"`
  // {"en":"Zone.", "zh_CN":"时区。"}
  Zone *string `json:"zone,omitempty" xml:"zone,omitempty" require:"true"`
  // {"en":"Attack time.", "zh_CN":"攻击时间"}
  Attack_time *string `json:"attack_time,omitempty" xml:"attack_time,omitempty" require:"true"`
  // {"en":"Strategy description.", "zh_CN":"策略描述。。"}
  Strategy_desc *string `json:"strategy_desc,omitempty" xml:"strategy_desc,omitempty" require:"true"`
  // {"en":"Domain.", "zh_CN":"域名。"}
  Host *string `json:"host,omitempty" xml:"host,omitempty" require:"true"`
  // {"en":"IP geographical location.", "zh_CN":"IP地理位置。"}
  Location *string `json:"location,omitempty" xml:"location,omitempty" require:"true"`
  // {"en":"Strategy name.", "zh_CN":"策略名称。"}
  Strategy_name *string `json:"strategy_name,omitempty" xml:"strategy_name,omitempty" require:"true"`
  // {"en":"User-Agent.", "zh_CN":"User-Agent"}
  User_agent *string `json:"user_agent,omitempty" xml:"user_agent,omitempty" require:"true"`
  // {"en":"Domain detail.", "zh_CN":"域名详情。"}
  Detail_host *string `json:"detail_host,omitempty" xml:"detail_host,omitempty" require:"true"`
  // {"en":"Status code.", "zh_CN":"状态码。"}
  StatusCode *int `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s GetBotAttackIncidentLogDataResponseDataList) String() string {
  return tea.Prettify(s)
}

func (s GetBotAttackIncidentLogDataResponseDataList) GoString() string {
  return s.String()
}

func (s *GetBotAttackIncidentLogDataResponseDataList) SetReferer(v string) *GetBotAttackIncidentLogDataResponseDataList {
  s.Referer = &v
  return s
}

func (s *GetBotAttackIncidentLogDataResponseDataList) SetBrowser_fp(v string) *GetBotAttackIncidentLogDataResponseDataList {
  s.Browser_fp = &v
  return s
}

func (s *GetBotAttackIncidentLogDataResponseDataList) SetAttack_type(v string) *GetBotAttackIncidentLogDataResponseDataList {
  s.Attack_type = &v
  return s
}

func (s *GetBotAttackIncidentLogDataResponseDataList) SetRule_name(v string) *GetBotAttackIncidentLogDataResponseDataList {
  s.Rule_name = &v
  return s
}

func (s *GetBotAttackIncidentLogDataResponseDataList) SetIp(v string) *GetBotAttackIncidentLogDataResponseDataList {
  s.Ip = &v
  return s
}

func (s *GetBotAttackIncidentLogDataResponseDataList) SetUuid(v string) *GetBotAttackIncidentLogDataResponseDataList {
  s.Uuid = &v
  return s
}

func (s *GetBotAttackIncidentLogDataResponseDataList) SetVersion(v string) *GetBotAttackIncidentLogDataResponseDataList {
  s.Version = &v
  return s
}

func (s *GetBotAttackIncidentLogDataResponseDataList) SetClient_id(v string) *GetBotAttackIncidentLogDataResponseDataList {
  s.Client_id = &v
  return s
}

func (s *GetBotAttackIncidentLogDataResponseDataList) SetUrl(v string) *GetBotAttackIncidentLogDataResponseDataList {
  s.Url = &v
  return s
}

func (s *GetBotAttackIncidentLogDataResponseDataList) SetBlock_id(v string) *GetBotAttackIncidentLogDataResponseDataList {
  s.Block_id = &v
  return s
}

func (s *GetBotAttackIncidentLogDataResponseDataList) SetContent(v string) *GetBotAttackIncidentLogDataResponseDataList {
  s.Content = &v
  return s
}

func (s *GetBotAttackIncidentLogDataResponseDataList) SetMode(v string) *GetBotAttackIncidentLogDataResponseDataList {
  s.Mode = &v
  return s
}

func (s *GetBotAttackIncidentLogDataResponseDataList) SetFinal_rule_id(v int) *GetBotAttackIncidentLogDataResponseDataList {
  s.Final_rule_id = &v
  return s
}

func (s *GetBotAttackIncidentLogDataResponseDataList) SetEvent_type(v string) *GetBotAttackIncidentLogDataResponseDataList {
  s.Event_type = &v
  return s
}

func (s *GetBotAttackIncidentLogDataResponseDataList) SetAct(v string) *GetBotAttackIncidentLogDataResponseDataList {
  s.Act = &v
  return s
}

func (s *GetBotAttackIncidentLogDataResponseDataList) SetZone(v string) *GetBotAttackIncidentLogDataResponseDataList {
  s.Zone = &v
  return s
}

func (s *GetBotAttackIncidentLogDataResponseDataList) SetAttack_time(v string) *GetBotAttackIncidentLogDataResponseDataList {
  s.Attack_time = &v
  return s
}

func (s *GetBotAttackIncidentLogDataResponseDataList) SetStrategy_desc(v string) *GetBotAttackIncidentLogDataResponseDataList {
  s.Strategy_desc = &v
  return s
}

func (s *GetBotAttackIncidentLogDataResponseDataList) SetHost(v string) *GetBotAttackIncidentLogDataResponseDataList {
  s.Host = &v
  return s
}

func (s *GetBotAttackIncidentLogDataResponseDataList) SetLocation(v string) *GetBotAttackIncidentLogDataResponseDataList {
  s.Location = &v
  return s
}

func (s *GetBotAttackIncidentLogDataResponseDataList) SetStrategy_name(v string) *GetBotAttackIncidentLogDataResponseDataList {
  s.Strategy_name = &v
  return s
}

func (s *GetBotAttackIncidentLogDataResponseDataList) SetUser_agent(v string) *GetBotAttackIncidentLogDataResponseDataList {
  s.User_agent = &v
  return s
}

func (s *GetBotAttackIncidentLogDataResponseDataList) SetDetail_host(v string) *GetBotAttackIncidentLogDataResponseDataList {
  s.Detail_host = &v
  return s
}

func (s *GetBotAttackIncidentLogDataResponseDataList) SetStatusCode(v int) *GetBotAttackIncidentLogDataResponseDataList {
  s.StatusCode = &v
  return s
}

type GetBotAttackIncidentLogDataPaths struct {
}

func (s GetBotAttackIncidentLogDataPaths) String() string {
  return tea.Prettify(s)
}

func (s GetBotAttackIncidentLogDataPaths) GoString() string {
  return s.String()
}

type GetBotAttackIncidentLogDataParameters struct {
}

func (s GetBotAttackIncidentLogDataParameters) String() string {
  return tea.Prettify(s)
}

func (s GetBotAttackIncidentLogDataParameters) GoString() string {
  return s.String()
}

type GetBotAttackIncidentLogDataRequestHeader struct {
}

func (s GetBotAttackIncidentLogDataRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetBotAttackIncidentLogDataRequestHeader) GoString() string {
  return s.String()
}

type GetBotAttackIncidentLogDataResponseHeader struct {
}

func (s GetBotAttackIncidentLogDataResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetBotAttackIncidentLogDataResponseHeader) GoString() string {
  return s.String()
}




type GetRequestsByRecordTypeRequest struct {
}

func (s GetRequestsByRecordTypeRequest) String() string {
  return tea.Prettify(s)
}

func (s GetRequestsByRecordTypeRequest) GoString() string {
  return s.String()
}

type GetRequestsByRecordTypeResponse struct {
  // {"en":"return data", "zh_CN":"返回值"}
  Data []*GetRequestsByRecordTypeResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
  // {"en":"code,success is 0", "zh_CN":"状态码,成功为0"}
  Code *int `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"error message or Success", "zh_CN":"错误信息或Success"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s GetRequestsByRecordTypeResponse) String() string {
  return tea.Prettify(s)
}

func (s GetRequestsByRecordTypeResponse) GoString() string {
  return s.String()
}

func (s *GetRequestsByRecordTypeResponse) SetData(v []*GetRequestsByRecordTypeResponseData) *GetRequestsByRecordTypeResponse {
  s.Data = v
  return s
}

func (s *GetRequestsByRecordTypeResponse) SetCode(v int) *GetRequestsByRecordTypeResponse {
  s.Code = &v
  return s
}

func (s *GetRequestsByRecordTypeResponse) SetMessage(v string) *GetRequestsByRecordTypeResponse {
  s.Message = &v
  return s
}

type GetRequestsByRecordTypeResponseData struct     {
  // {"en":"timestamp (accurate to the second)", "zh_CN":"时间戳，精确到秒"}
  Ts *int `json:"ts,omitempty" xml:"ts,omitempty" require:"true"`
  // {"en":"data", "zh_CN":"值"}
  Results []*GetRequestsByRecordTypeResponseDataResults `json:"results,omitempty" xml:"results,omitempty" require:"true" type:"Repeated"`
}

func (s GetRequestsByRecordTypeResponseData) String() string {
  return tea.Prettify(s)
}

func (s GetRequestsByRecordTypeResponseData) GoString() string {
  return s.String()
}

func (s *GetRequestsByRecordTypeResponseData) SetTs(v int) *GetRequestsByRecordTypeResponseData {
  s.Ts = &v
  return s
}

func (s *GetRequestsByRecordTypeResponseData) SetResults(v []*GetRequestsByRecordTypeResponseDataResults) *GetRequestsByRecordTypeResponseData {
  s.Results = v
  return s
}

type GetRequestsByRecordTypeResponseDataResults struct     {
  // {"en":"record type(A,AAAA,CNAME,TXT,MX,SRV,RP,SPF,RP,PTR,NS,SOA,UNKNOWN)", "zh_CN":"记录类型(A,AAAA,CNAME,TXT,MX,SRV,RP,SPF,RP,PTR,NS,SOA,UNKNOWN)"}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
  // {"en":"count", "zh_CN":"数量"}
  Count *int `json:"count,omitempty" xml:"count,omitempty" require:"true"`
}

func (s GetRequestsByRecordTypeResponseDataResults) String() string {
  return tea.Prettify(s)
}

func (s GetRequestsByRecordTypeResponseDataResults) GoString() string {
  return s.String()
}

func (s *GetRequestsByRecordTypeResponseDataResults) SetType(v string) *GetRequestsByRecordTypeResponseDataResults {
  s.Type = &v
  return s
}

func (s *GetRequestsByRecordTypeResponseDataResults) SetCount(v int) *GetRequestsByRecordTypeResponseDataResults {
  s.Count = &v
  return s
}

type GetRequestsByRecordTypePaths struct {
}

func (s GetRequestsByRecordTypePaths) String() string {
  return tea.Prettify(s)
}

func (s GetRequestsByRecordTypePaths) GoString() string {
  return s.String()
}

type GetRequestsByRecordTypeParameters struct {
  // {"en":"from time(timestamp) (accurate to the second)", "zh_CN":"起始时间，timestamp类型，精确到秒"}
  From *int `json:"from,omitempty" xml:"from,omitempty" require:"true"`
  // {"en":"to time(timestamp) (accurate to the second)", "zh_CN":"结束时间，timestamp类型，精确到秒"}
  To *int `json:"to,omitempty" xml:"to,omitempty" require:"true"`
  // {"en":"interval time, default is 'hourly'. Support interval:oneminute: 1 min.(Range <= 1 day);fiveminutes: 5 min(Range <= 7 days); hourly: 1 hr; daily: 1 day; monthly: 1 month; all: Without interval(Combine data into one, Range>=1hr)", "zh_CN":"间隔时间，默认为hourly。支持的类型有:oneminute: 1 min.(Range <= 1 day);fiveminutes: 5 min(Range <= 7 days); hourly: 1 hr; daily: 1 day; monthly: 1 month; all: Without interval(合并数据成一条返回，Range>=1hr)"}
  Interval *string `json:"interval,omitempty" xml:"interval,omitempty"`
  // {"en":"zone names, separeated by ',', case insensitive", "zh_CN":"zone名称字符串，英文逗号分隔，不区分大小写"}
  Zone *string `json:"zone,omitempty" xml:"zone,omitempty"`
  // {"en":"timezone is necessary when interval=daily/monthly/all, the default timezone is 0 (UTC+0)", "zh_CN":"时区，当interval为daily/monthly/all时需要判断时区。"}
  Timezone *int `json:"timezone,omitempty" xml:"timezone,omitempty"`
}

func (s GetRequestsByRecordTypeParameters) String() string {
  return tea.Prettify(s)
}

func (s GetRequestsByRecordTypeParameters) GoString() string {
  return s.String()
}

func (s *GetRequestsByRecordTypeParameters) SetFrom(v int) *GetRequestsByRecordTypeParameters {
  s.From = &v
  return s
}

func (s *GetRequestsByRecordTypeParameters) SetTo(v int) *GetRequestsByRecordTypeParameters {
  s.To = &v
  return s
}

func (s *GetRequestsByRecordTypeParameters) SetInterval(v string) *GetRequestsByRecordTypeParameters {
  s.Interval = &v
  return s
}

func (s *GetRequestsByRecordTypeParameters) SetZone(v string) *GetRequestsByRecordTypeParameters {
  s.Zone = &v
  return s
}

func (s *GetRequestsByRecordTypeParameters) SetTimezone(v int) *GetRequestsByRecordTypeParameters {
  s.Timezone = &v
  return s
}

type GetRequestsByRecordTypeRequestHeader struct {
}

func (s GetRequestsByRecordTypeRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetRequestsByRecordTypeRequestHeader) GoString() string {
  return s.String()
}

type GetRequestsByRecordTypeResponseHeader struct {
}

func (s GetRequestsByRecordTypeResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetRequestsByRecordTypeResponseHeader) GoString() string {
  return s.String()
}




type GetOverallRequestsRequest struct {
}

func (s GetOverallRequestsRequest) String() string {
  return tea.Prettify(s)
}

func (s GetOverallRequestsRequest) GoString() string {
  return s.String()
}

type GetOverallRequestsResponse struct {
  // {"en":"return data", "zh_CN":"返回值"}
  Data []*GetOverallRequestsResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
  // {"en":"code,success is 0", "zh_CN":"状态码,成功为0"}
  Code *int `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"error message or Success", "zh_CN":"错误信息或Success"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s GetOverallRequestsResponse) String() string {
  return tea.Prettify(s)
}

func (s GetOverallRequestsResponse) GoString() string {
  return s.String()
}

func (s *GetOverallRequestsResponse) SetData(v []*GetOverallRequestsResponseData) *GetOverallRequestsResponse {
  s.Data = v
  return s
}

func (s *GetOverallRequestsResponse) SetCode(v int) *GetOverallRequestsResponse {
  s.Code = &v
  return s
}

func (s *GetOverallRequestsResponse) SetMessage(v string) *GetOverallRequestsResponse {
  s.Message = &v
  return s
}

type GetOverallRequestsResponseData struct     {
  // {"en":"timestamp (accurate to the second)", "zh_CN":"时间戳，精确到秒"}
  Ts *int `json:"ts,omitempty" xml:"ts,omitempty" require:"true"`
  // {"en":"count(Returns when groupBy has no value)", "zh_CN":"数量(groupBy无值时返回)"}
  Count *int `json:"count,omitempty" xml:"count,omitempty"`
  // {"en":"Returns only if groupBy has a value", "zh_CN":"只有在groupBy有值时返回"}
  Results []*GetOverallRequestsResponseDataResults `json:"results,omitempty" xml:"results,omitempty" type:"Repeated"`
}

func (s GetOverallRequestsResponseData) String() string {
  return tea.Prettify(s)
}

func (s GetOverallRequestsResponseData) GoString() string {
  return s.String()
}

func (s *GetOverallRequestsResponseData) SetTs(v int) *GetOverallRequestsResponseData {
  s.Ts = &v
  return s
}

func (s *GetOverallRequestsResponseData) SetCount(v int) *GetOverallRequestsResponseData {
  s.Count = &v
  return s
}

func (s *GetOverallRequestsResponseData) SetResults(v []*GetOverallRequestsResponseDataResults) *GetOverallRequestsResponseData {
  s.Results = v
  return s
}

type GetOverallRequestsResponseDataResults struct     {
  // {"en":"Use when groupBy is equal to zone", "zh_CN":"groupBy等于zone时使用"}
  Zone *string `json:"zone,omitempty" xml:"zone,omitempty"`
  // {"en":"Use when groupBy is equal to clb_domain", "zh_CN":"groupBy等于clb_domain时使用"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty"`
  // {"en":"count", "zh_CN":"数量"}
  Count *int `json:"count,omitempty" xml:"count,omitempty"`
}

func (s GetOverallRequestsResponseDataResults) String() string {
  return tea.Prettify(s)
}

func (s GetOverallRequestsResponseDataResults) GoString() string {
  return s.String()
}

func (s *GetOverallRequestsResponseDataResults) SetZone(v string) *GetOverallRequestsResponseDataResults {
  s.Zone = &v
  return s
}

func (s *GetOverallRequestsResponseDataResults) SetDomain(v string) *GetOverallRequestsResponseDataResults {
  s.Domain = &v
  return s
}

func (s *GetOverallRequestsResponseDataResults) SetCount(v int) *GetOverallRequestsResponseDataResults {
  s.Count = &v
  return s
}

type GetOverallRequestsPaths struct {
}

func (s GetOverallRequestsPaths) String() string {
  return tea.Prettify(s)
}

func (s GetOverallRequestsPaths) GoString() string {
  return s.String()
}

type GetOverallRequestsParameters struct {
  // {"en":"from time(timestamp) (accurate to the second)", "zh_CN":"起始时间，timestamp类型，精确到秒"}
  From *int `json:"from,omitempty" xml:"from,omitempty" require:"true"`
  // {"en":"to time(timestamp) (accurate to the second)", "zh_CN":"结束时间，timestamp类型，精确到秒"}
  To *int `json:"to,omitempty" xml:"to,omitempty" require:"true"`
  // {"en":"interval time, default is 'hourly'. Support interval:oneminute: 1 min.(Range <= 1 day);fiveminutes: 5 min(Range <= 7 days); hourly: 1 hr; daily: 1 day; monthly: 1 month; all: Without interval(Combine data into one, Range>=1hr)", "zh_CN":"间隔时间，默认为hourly。支持的类型有:oneminute: 1 min.(Range <= 1 day);fiveminutes: 5 min(Range <= 7 days); hourly: 1 hr; daily: 1 day; monthly: 1 month; all: Without interval(合并数据成一条返回，Range>=1hr)"}
  Interval *string `json:"interval,omitempty" xml:"interval,omitempty"`
  // {"en":"zone names, separeated by ',', case insensitive", "zh_CN":"zone名称字符串，英文逗号分隔，不区分大小写"}
  Zone *string `json:"zone,omitempty" xml:"zone,omitempty"`
  // {"en":"search group by, value can be 'zone' or 'clb_domain'. If it is 'clb_domain', only the statistics of Clb Domain will be queried", "zh_CN":"按选定类型分组，值可以为'zone'或'clb_domain'。如果为'clb_domain'，只会查询Clb Domain的统计数据"}
  GroupBy *string `json:"groupBy,omitempty" xml:"groupBy,omitempty"`
  // {"en":"timezone is necessary when interval=daily/monthly/all, the default timezone is 0 (UTC+0)", "zh_CN":"时区，当interval为daily/monthly/all时需要判断时区。"}
  Timezone *int `json:"timezone,omitempty" xml:"timezone,omitempty"`
}

func (s GetOverallRequestsParameters) String() string {
  return tea.Prettify(s)
}

func (s GetOverallRequestsParameters) GoString() string {
  return s.String()
}

func (s *GetOverallRequestsParameters) SetFrom(v int) *GetOverallRequestsParameters {
  s.From = &v
  return s
}

func (s *GetOverallRequestsParameters) SetTo(v int) *GetOverallRequestsParameters {
  s.To = &v
  return s
}

func (s *GetOverallRequestsParameters) SetInterval(v string) *GetOverallRequestsParameters {
  s.Interval = &v
  return s
}

func (s *GetOverallRequestsParameters) SetZone(v string) *GetOverallRequestsParameters {
  s.Zone = &v
  return s
}

func (s *GetOverallRequestsParameters) SetGroupBy(v string) *GetOverallRequestsParameters {
  s.GroupBy = &v
  return s
}

func (s *GetOverallRequestsParameters) SetTimezone(v int) *GetOverallRequestsParameters {
  s.Timezone = &v
  return s
}

type GetOverallRequestsRequestHeader struct {
}

func (s GetOverallRequestsRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetOverallRequestsRequestHeader) GoString() string {
  return s.String()
}

type GetOverallRequestsResponseHeader struct {
}

func (s GetOverallRequestsResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetOverallRequestsResponseHeader) GoString() string {
  return s.String()
}




