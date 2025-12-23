package other

// This file is auto-generated, don't edit it. Thanks.
import (
  "github.com/alibabacloud-go/tea/tea"
)

type GetLiveStreamPushingStatusRequest struct {
}

func (s GetLiveStreamPushingStatusRequest) String() string {
  return tea.Prettify(s)
}

func (s GetLiveStreamPushingStatusRequest) GoString() string {
  return s.String()
}

type GetLiveStreamPushingStatusRequestHeader struct {
}

func (s GetLiveStreamPushingStatusRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetLiveStreamPushingStatusRequestHeader) GoString() string {
  return s.String()
}

type GetLiveStreamPushingStatusPaths struct {
}

func (s GetLiveStreamPushingStatusPaths) String() string {
  return tea.Prettify(s)
}

func (s GetLiveStreamPushingStatusPaths) GoString() string {
  return s.String()
}

type GetLiveStreamPushingStatusParameters struct {
  // {"en":"Push domain(multiple domains supported, separated by commas)(    All parameters are passed via HTTP GET requests.)","zh_CN":"推流域名（所有参数以HTTP GET请求方式传参)"}
  U *string `json:"u,omitempty" xml:"u,omitempty" require:"true"`
  // {"en":"Time, eg: 20160527152300, if not filled in, the current time -5 minutes","zh_CN":"时间，eg：20160527152300，不填时为当前时间-5分钟"}
  T *int `json:"t,omitempty" xml:"t,omitempty"`
  // {"en":"Push channel URL (multiple push channel URLs are supported, separated by commas)","zh_CN":"推流流名(支持多个推流流名，以逗号分隔)"}
  Channel *string `json:"channel,omitempty" xml:"channel,omitempty"`
  // {"en":"Query interval. The value can be 10 or 60. The default value is 60\nWhen g is 10, query the data of the nearest whole 10 seconds to time t\nWhen g is 60, query the data of the nearest whole minute to time t","zh_CN":"查询间隔，可选值10、60，默认60\n当g为10时，查询距离时间t最近的整10秒点数据\n当g为60时，查询距离时间t最近的整分钟点数据"}
  G *string `json:"g,omitempty" xml:"g,omitempty"`
}

func (s GetLiveStreamPushingStatusParameters) String() string {
  return tea.Prettify(s)
}

func (s GetLiveStreamPushingStatusParameters) GoString() string {
  return s.String()
}

func (s *GetLiveStreamPushingStatusParameters) SetU(v string) *GetLiveStreamPushingStatusParameters {
  s.U = &v
  return s
}

func (s *GetLiveStreamPushingStatusParameters) SetT(v int) *GetLiveStreamPushingStatusParameters {
  s.T = &v
  return s
}

func (s *GetLiveStreamPushingStatusParameters) SetChannel(v string) *GetLiveStreamPushingStatusParameters {
  s.Channel = &v
  return s
}

func (s *GetLiveStreamPushingStatusParameters) SetG(v string) *GetLiveStreamPushingStatusParameters {
  s.G = &v
  return s
}

type GetLiveStreamPushingStatusResponse struct {
  // {"en":"The time of the data returned","zh_CN":"返回的数据的时间"}
  Rettime *string `json:"rettime,omitempty" xml:"rettime,omitempty" require:"true"`
  // {"en":"Number of data items. 0 is returned if there is no data","zh_CN":"数据条数，无数据返回0"}
  Retcode *int64 `json:"retcode,omitempty" xml:"retcode,omitempty" require:"true"`
  // {"en":"","zh_CN":"推流信息数据集合"}
  DataValue []*GetLiveStreamPushingStatusResponseDataValue `json:"dataValue,omitempty" xml:"dataValue,omitempty" require:"true" type:"Repeated"`
}

func (s GetLiveStreamPushingStatusResponse) String() string {
  return tea.Prettify(s)
}

func (s GetLiveStreamPushingStatusResponse) GoString() string {
  return s.String()
}

func (s *GetLiveStreamPushingStatusResponse) SetRettime(v string) *GetLiveStreamPushingStatusResponse {
  s.Rettime = &v
  return s
}

func (s *GetLiveStreamPushingStatusResponse) SetRetcode(v int64) *GetLiveStreamPushingStatusResponse {
  s.Retcode = &v
  return s
}

func (s *GetLiveStreamPushingStatusResponse) SetDataValue(v []*GetLiveStreamPushingStatusResponseDataValue) *GetLiveStreamPushingStatusResponse {
  s.DataValue = v
  return s
}

type GetLiveStreamPushingStatusResponseDataValue struct     {
  // {"en":"Push stream CDN node, that is, the IP address of edge node of the received data source, separated by multiple commas","zh_CN":"推流cdn节点，即收到数据源的edge节点IP地址，多个逗号分隔"}
  Deployaddress *string `json:"deployaddress,omitempty" xml:"deployaddress,omitempty" require:"true"`
  // {"en":"Push user IP (data source) address, separated by multiple commas","zh_CN":"推流用户ip（数据源）地址，多个逗号分隔"}
  Inaddress *string `json:"inaddress,omitempty" xml:"inaddress,omitempty" require:"true"`
  // {"en":"The channel of anchor","zh_CN":"主播流名"}
  Streamname *string `json:"streamname,omitempty" xml:"streamname,omitempty" require:"true"`
  // {"en":"Anchor Current encoding frame rate","zh_CN":"主播当前编码帧率"}
  Fps *int64 `json:"fps,omitempty" xml:"fps,omitempty" require:"true"`
  // {"en":"Current frame loss rate of anchor","zh_CN":"主播当前丢帧率"}
  Lfr *GetLiveStreamPushingStatusResponseDataValueLfr `json:"lfr,omitempty" xml:"lfr,omitempty" require:"true" type:"Struct"`
  // {"en":"Anchor Current bit rate","zh_CN":"主播当前码率"}
  Inbandwidth *int64 `json:"inbandwidth,omitempty" xml:"inbandwidth,omitempty" require:"true"`
  // {"en":"Video timestamps are separated by multiple commas","zh_CN":"视频时间戳 多个逗号分隔"}
  Videotmstmp *string `json:"videotmstmp,omitempty" xml:"videotmstmp,omitempty" require:"true"`
  // {"en":"Audio timestamp, separated by multiple commas","zh_CN":"音频时间戳，多个逗号分隔"}
  Audiotmstmp *string `json:"audiotmstmp,omitempty" xml:"audiotmstmp,omitempty" require:"true"`
}

func (s GetLiveStreamPushingStatusResponseDataValue) String() string {
  return tea.Prettify(s)
}

func (s GetLiveStreamPushingStatusResponseDataValue) GoString() string {
  return s.String()
}

func (s *GetLiveStreamPushingStatusResponseDataValue) SetDeployaddress(v string) *GetLiveStreamPushingStatusResponseDataValue {
  s.Deployaddress = &v
  return s
}

func (s *GetLiveStreamPushingStatusResponseDataValue) SetInaddress(v string) *GetLiveStreamPushingStatusResponseDataValue {
  s.Inaddress = &v
  return s
}

func (s *GetLiveStreamPushingStatusResponseDataValue) SetStreamname(v string) *GetLiveStreamPushingStatusResponseDataValue {
  s.Streamname = &v
  return s
}

func (s *GetLiveStreamPushingStatusResponseDataValue) SetFps(v int64) *GetLiveStreamPushingStatusResponseDataValue {
  s.Fps = &v
  return s
}

func (s *GetLiveStreamPushingStatusResponseDataValue) SetLfr(v *GetLiveStreamPushingStatusResponseDataValueLfr) *GetLiveStreamPushingStatusResponseDataValue {
  s.Lfr = v
  return s
}

func (s *GetLiveStreamPushingStatusResponseDataValue) SetInbandwidth(v int64) *GetLiveStreamPushingStatusResponseDataValue {
  s.Inbandwidth = &v
  return s
}

func (s *GetLiveStreamPushingStatusResponseDataValue) SetVideotmstmp(v string) *GetLiveStreamPushingStatusResponseDataValue {
  s.Videotmstmp = &v
  return s
}

func (s *GetLiveStreamPushingStatusResponseDataValue) SetAudiotmstmp(v string) *GetLiveStreamPushingStatusResponseDataValue {
  s.Audiotmstmp = &v
  return s
}

type GetLiveStreamPushingStatusResponseDataValueLfr struct {
}

func (s GetLiveStreamPushingStatusResponseDataValueLfr) String() string {
  return tea.Prettify(s)
}

func (s GetLiveStreamPushingStatusResponseDataValueLfr) GoString() string {
  return s.String()
}

type GetLiveStreamPushingStatusResponseHeader struct {
}

func (s GetLiveStreamPushingStatusResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetLiveStreamPushingStatusResponseHeader) GoString() string {
  return s.String()
}




type ResourceGroupDomainChangeRequest struct {
  // {"en":"Domains, separate by ;, no more than 100 domains.", "zh_CN":"域名，多个使用;分隔，最多不超过100个域名。"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {"en":"The resource group for domains want to changed.", "zh_CN":"域名需要变更的资源组。"}
  Resource *string `json:"resource,omitempty" xml:"resource,omitempty" require:"true"`
}

func (s ResourceGroupDomainChangeRequest) String() string {
  return tea.Prettify(s)
}

func (s ResourceGroupDomainChangeRequest) GoString() string {
  return s.String()
}

func (s *ResourceGroupDomainChangeRequest) SetDomain(v string) *ResourceGroupDomainChangeRequest {
  s.Domain = &v
  return s
}

func (s *ResourceGroupDomainChangeRequest) SetResource(v string) *ResourceGroupDomainChangeRequest {
  s.Resource = &v
  return s
}

type ResourceGroupDomainChangeResponse struct {
  // {"en":"code", "zh_CN":"返回code"}
  Code *int `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"msg", "zh_CN":"返回描述信息"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
}

func (s ResourceGroupDomainChangeResponse) String() string {
  return tea.Prettify(s)
}

func (s ResourceGroupDomainChangeResponse) GoString() string {
  return s.String()
}

func (s *ResourceGroupDomainChangeResponse) SetCode(v int) *ResourceGroupDomainChangeResponse {
  s.Code = &v
  return s
}

func (s *ResourceGroupDomainChangeResponse) SetMsg(v string) *ResourceGroupDomainChangeResponse {
  s.Msg = &v
  return s
}

type ResourceGroupDomainChangePaths struct {
}

func (s ResourceGroupDomainChangePaths) String() string {
  return tea.Prettify(s)
}

func (s ResourceGroupDomainChangePaths) GoString() string {
  return s.String()
}

type ResourceGroupDomainChangeParameters struct {
}

func (s ResourceGroupDomainChangeParameters) String() string {
  return tea.Prettify(s)
}

func (s ResourceGroupDomainChangeParameters) GoString() string {
  return s.String()
}

type ResourceGroupDomainChangeRequestHeader struct {
}

func (s ResourceGroupDomainChangeRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ResourceGroupDomainChangeRequestHeader) GoString() string {
  return s.String()
}

type ResourceGroupDomainChangeResponseHeader struct {
}

func (s ResourceGroupDomainChangeResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ResourceGroupDomainChangeResponseHeader) GoString() string {
  return s.String()
}




type QueryDailyLiveTranscodingDurationRequest struct {
  // {"en":"cust_en_name of sub-client.
  // When a merged-account wants to  view the information of the subclient,the cust_en_name is required.", "zh_CN":"合并账号下的某个客户的英文名，当合并账号要查看子客户的信息时，必须填写子客户的英文名"}
  Cust *string `json:"cust,omitempty" xml:"cust,omitempty"`
  // {"en":"Specifies the query date:
  // 1.With format yyyy-mm-dd.
  // 2.If not specified,it means today as default.", "zh_CN":"查询的日期，日期格式为yyyy-mm-dd,不选或者为空时默认为当天；"}
  Date *string `json:"date,omitempty" xml:"date,omitempty"`
  // {"en":"1.Must work with 'enddate' and they  specify the query date scope. 
  // 2.With format yyyy-mm-dd.
  // 3.If there is a 'date' parameter,this parameter will be omitted.", "zh_CN":"查询的起始日期,日期格式为yyyy-mm-dd；此参数需与enddate参数配合,若存在date参数,则该参数无效"}
  Startdate *string `json:"startdate,omitempty" xml:"startdate,omitempty"`
  // {"en":"1.Must work with 'startdate' and they  specify the query date scope. 
  // 2.With format yyyy-mm-dd.
  // 3.If there is a 'date' parameter,this parameter will be omitted.", "zh_CN":"查询的结束日期,日期格式为yyyy-mm-dd；此参数需与startdate参数配合,若存在date参数,则该参数无效。"}
  Enddate *string `json:"enddate,omitempty" xml:"enddate,omitempty"`
  // {"en":"domains that been queried:
  // 1.If there are multiple inputs,use  ';' as separator.
  // 2.If not specified, it means all the domains of the account .", "zh_CN":"查询的频道，多个频道值请用英文分号';'，不选或者为空时默认为所查询客户的所有频道"}
  Channel *string `json:"channel,omitempty" xml:"channel,omitempty"`
  // {"en":"This parameter specifies if the 'channel' parameter should be exactly matched:
  // 1.'true' as default.
  // 2. If not 'true',it will query data of channels that ends with any item of input 'channel's.", "zh_CN":"&nbsp;频道是否完全匹配,为true时，必须填写完整的域名(此时会过滤用户输入的无效或重复频道,所有输入频道都无效时返403.。不为true时，显示以用户输入的频道为结尾的所有频道。默认为true"}
  IsExactMatch *string `json:"isExactMatch,omitempty" xml:"isExactMatch,omitempty"`
  // {"en":"acceleration type.
  // 1.If there are multiple inputs,use ';' as separator.
  // 2.If not specified or specified as 'all', it means all the accetypes.", "zh_CN":"加速类型参数，如accetype=web。多个请用英文分号';'分隔开，不填或值为all表示所有类型"}
  Accetype *string `json:"accetype,omitempty" xml:"accetype,omitempty"`
  // {"en":"The response format:
  // 1.optional values:xml, json.
  // 2.'xml' as default.", "zh_CN":"返回结果格式,支持格式为xml和json,默认为xml"}
  Dataformat *string `json:"dataformat,omitempty" xml:"dataformat,omitempty"`
  // {"en":"Display statistic result in merged or separate way:
  // 1.If specified 1,get the merged result.
  // 2.If specified 2,get the separate result.
  // 3.If specified 3,get both merged result and separate result.
  // 4.If not specified,means '1'.", "zh_CN":"&nbsp;结果的显示是否提供合并值。填写1时：只提供合并结果；填写2时：只提供拆分值；填写3时：既提供合并值，又提供拆分值。不选或者为空时默认为'1'。"}
  ResultType *string `json:"resultType,omitempty" xml:"resultType,omitempty"`
  // {"en":"Transcoding type, values can be h264, h265, zdgq_264, zdgq_265, cf_264, cf_265, or other. Multiple transcoding types should be separated by a semicolon. If some of the transcoding types are incorrect, the system will return data for the correct types; if all transcoding types are incorrect, it will return an error 'invalid transcodeType.' If not provided or left empty, it defaults to all types.", "zh_CN":"转码类型,值为h264、h265、zdgq_264、zdgq_265、cf_264、cf_265，other 多个转码类型用英文分号;分隔开。当传入转码类型部分错误时，返回正确的类型的数据；当传入转码类型全部错误时，返回错误invalid transcodeType. 不填或为空，默认为所有类型."}
  TranscodeType *string `json:"transcodeType,omitempty" xml:"transcodeType,omitempty"`
  // {"en":"GMT time zone, parameter format: GMT+09:00 means east 9th zone, GMT-09:00 means west 9th zone, if not transmitted, the default is local time zone (east 8th zone).", "zh_CN":"格林尼治时区，参数格式 GMT+09:00 表示东九区，GMT-09:00 表示西9区，不传则默认为本地时区（东八区）"}
  Timezone *string `json:"timezone,omitempty" xml:"timezone,omitempty"`
  // {"en":"Resolution types include LD480,SD720,HD1080,2K,4K,8K,SD576,SD540,LD360,LD240. Multiple resolutions are separated by a semicolon. When isAudio=1, this parameter is invalid and will return an error. Param definition must be empty when querying audio data.", "zh_CN":"清晰度类型,值为LD480,SD720,HD1080,2K,4K,8K,SD576,SD540,LD360,LD240，多个清晰度用英文分号;分隔开, 当isAudio=1时，此入参无效返回错误 param definition must be empty when query audio data."}
  Definition *string `json:"definition,omitempty" xml:"definition,omitempty"`
  // {"en":"Audio/Video Type, 1: Audio 2: Video. Defaults to 2 if not selected or empty. Only a single value is allowed.", "zh_CN":"音视频类型, 1:音频   2:视频. 不选或者为空时默认为2. 只能输入单个值."}
  IsAudio *string `json:"isAudio,omitempty" xml:"isAudio,omitempty"`
}

func (s QueryDailyLiveTranscodingDurationRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryDailyLiveTranscodingDurationRequest) GoString() string {
  return s.String()
}

func (s *QueryDailyLiveTranscodingDurationRequest) SetCust(v string) *QueryDailyLiveTranscodingDurationRequest {
  s.Cust = &v
  return s
}

func (s *QueryDailyLiveTranscodingDurationRequest) SetDate(v string) *QueryDailyLiveTranscodingDurationRequest {
  s.Date = &v
  return s
}

func (s *QueryDailyLiveTranscodingDurationRequest) SetStartdate(v string) *QueryDailyLiveTranscodingDurationRequest {
  s.Startdate = &v
  return s
}

func (s *QueryDailyLiveTranscodingDurationRequest) SetEnddate(v string) *QueryDailyLiveTranscodingDurationRequest {
  s.Enddate = &v
  return s
}

func (s *QueryDailyLiveTranscodingDurationRequest) SetChannel(v string) *QueryDailyLiveTranscodingDurationRequest {
  s.Channel = &v
  return s
}

func (s *QueryDailyLiveTranscodingDurationRequest) SetIsExactMatch(v string) *QueryDailyLiveTranscodingDurationRequest {
  s.IsExactMatch = &v
  return s
}

func (s *QueryDailyLiveTranscodingDurationRequest) SetAccetype(v string) *QueryDailyLiveTranscodingDurationRequest {
  s.Accetype = &v
  return s
}

func (s *QueryDailyLiveTranscodingDurationRequest) SetDataformat(v string) *QueryDailyLiveTranscodingDurationRequest {
  s.Dataformat = &v
  return s
}

func (s *QueryDailyLiveTranscodingDurationRequest) SetResultType(v string) *QueryDailyLiveTranscodingDurationRequest {
  s.ResultType = &v
  return s
}

func (s *QueryDailyLiveTranscodingDurationRequest) SetTranscodeType(v string) *QueryDailyLiveTranscodingDurationRequest {
  s.TranscodeType = &v
  return s
}

func (s *QueryDailyLiveTranscodingDurationRequest) SetTimezone(v string) *QueryDailyLiveTranscodingDurationRequest {
  s.Timezone = &v
  return s
}

func (s *QueryDailyLiveTranscodingDurationRequest) SetDefinition(v string) *QueryDailyLiveTranscodingDurationRequest {
  s.Definition = &v
  return s
}

func (s *QueryDailyLiveTranscodingDurationRequest) SetIsAudio(v string) *QueryDailyLiveTranscodingDurationRequest {
  s.IsAudio = &v
  return s
}

type QueryDailyLiveTranscodingDurationResponse struct {
  // {'en':'provider', 'zh_CN':'结果'}
  Provider *QueryDailyLiveTranscodingDurationResponseProvider `json:"provider,omitempty" xml:"provider,omitempty" require:"true" type:"Struct"`
}

func (s QueryDailyLiveTranscodingDurationResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryDailyLiveTranscodingDurationResponse) GoString() string {
  return s.String()
}

func (s *QueryDailyLiveTranscodingDurationResponse) SetProvider(v *QueryDailyLiveTranscodingDurationResponseProvider) *QueryDailyLiveTranscodingDurationResponse {
  s.Provider = v
  return s
}

type QueryDailyLiveTranscodingDurationResponseProvider struct {
  // {'en':'tenant', 'zh_CN':'租户'}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {'en':'type', 'zh_CN':'接口类型'}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
  // {'en':'data', 'zh_CN':'直播转码时长每日统计数据'}
  Date *QueryDailyLiveTranscodingDurationResponseProviderDate `json:"date,omitempty" xml:"date,omitempty" require:"true" type:"Struct"`
}

func (s QueryDailyLiveTranscodingDurationResponseProvider) String() string {
  return tea.Prettify(s)
}

func (s QueryDailyLiveTranscodingDurationResponseProvider) GoString() string {
  return s.String()
}

func (s *QueryDailyLiveTranscodingDurationResponseProvider) SetName(v string) *QueryDailyLiveTranscodingDurationResponseProvider {
  s.Name = &v
  return s
}

func (s *QueryDailyLiveTranscodingDurationResponseProvider) SetType(v string) *QueryDailyLiveTranscodingDurationResponseProvider {
  s.Type = &v
  return s
}

func (s *QueryDailyLiveTranscodingDurationResponseProvider) SetDate(v *QueryDailyLiveTranscodingDurationResponseProviderDate) *QueryDailyLiveTranscodingDurationResponseProvider {
  s.Date = v
  return s
}

type QueryDailyLiveTranscodingDurationResponseProviderDate struct {
  // {'en':'startdate', 'zh_CN':'开始时间'}
  Startdate *string `json:"startdate,omitempty" xml:"startdate,omitempty" require:"true"`
  // {'en':'enddate', 'zh_CN':'结束时间'}
  Enddate *string `json:"enddate,omitempty" xml:"enddate,omitempty" require:"true"`
  // {'en':'transcoding', 'zh_CN':'转码类型'}
  Transcoding *QueryDailyLiveTranscodingDurationResponseProviderDateTranscoding `json:"transcoding,omitempty" xml:"transcoding,omitempty" require:"true" type:"Struct"`
}

func (s QueryDailyLiveTranscodingDurationResponseProviderDate) String() string {
  return tea.Prettify(s)
}

func (s QueryDailyLiveTranscodingDurationResponseProviderDate) GoString() string {
  return s.String()
}

func (s *QueryDailyLiveTranscodingDurationResponseProviderDate) SetStartdate(v string) *QueryDailyLiveTranscodingDurationResponseProviderDate {
  s.Startdate = &v
  return s
}

func (s *QueryDailyLiveTranscodingDurationResponseProviderDate) SetEnddate(v string) *QueryDailyLiveTranscodingDurationResponseProviderDate {
  s.Enddate = &v
  return s
}

func (s *QueryDailyLiveTranscodingDurationResponseProviderDate) SetTranscoding(v *QueryDailyLiveTranscodingDurationResponseProviderDateTranscoding) *QueryDailyLiveTranscodingDurationResponseProviderDate {
  s.Transcoding = v
  return s
}

type QueryDailyLiveTranscodingDurationResponseProviderDateTranscoding struct {
  // {'en':'channel', 'zh_CN':'频道'}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {'en':'live', 'zh_CN':'直播转码时长数据'}
  Live []*QueryDailyLiveTranscodingDurationResponseProviderDateTranscodingLive `json:"live,omitempty" xml:"live,omitempty" require:"true" type:"Repeated"`
}

func (s QueryDailyLiveTranscodingDurationResponseProviderDateTranscoding) String() string {
  return tea.Prettify(s)
}

func (s QueryDailyLiveTranscodingDurationResponseProviderDateTranscoding) GoString() string {
  return s.String()
}

func (s *QueryDailyLiveTranscodingDurationResponseProviderDateTranscoding) SetName(v string) *QueryDailyLiveTranscodingDurationResponseProviderDateTranscoding {
  s.Name = &v
  return s
}

func (s *QueryDailyLiveTranscodingDurationResponseProviderDateTranscoding) SetLive(v []*QueryDailyLiveTranscodingDurationResponseProviderDateTranscodingLive) *QueryDailyLiveTranscodingDurationResponseProviderDateTranscoding {
  s.Live = v
  return s
}

type QueryDailyLiveTranscodingDurationResponseProviderDateTranscodingLive struct     {
  // {'en':'timestamp', 'zh_CN':'时间点'}
  Time *string `json:"time,omitempty" xml:"time,omitempty" require:"true"`
  // {'en':'h264 transcoding time', 'zh_CN':'h264转码类型的转码时长(单位分钟，固定保留2位小数)'}
  H264 *string `json:"h264,omitempty" xml:"h264,omitempty" require:"true"`
  // {'en':'h265 transcoding time', 'zh_CN':'h265转码类型的转码时长(单位分钟，固定保留2位小数)'}
  H265 *string `json:"h265,omitempty" xml:"h265,omitempty" require:"true"`
  // {'en':'zdgq_264 transcoding time', 'zh_CN':'zdgq_264转码类型的转码时长(单位分钟，固定保留2位小数)'}
  Zdgq_264 *string `json:"zdgq_264,omitempty" xml:"zdgq_264,omitempty" require:"true"`
  // {'en':'zdgq_265 transcoding time', 'zh_CN':'zdgq_265转码类型的转码时长(单位分钟，固定保留2位小数)'}
  Zdgq_265 *string `json:"zdgq_265,omitempty" xml:"zdgq_265,omitempty" require:"true"`
  // {'en':'voice transcoding time', 'zh_CN':'voice转码类型的转码时长(单位分钟，固定保留2位小数)'}
  Voice *string `json:"voice,omitempty" xml:"voice,omitempty" require:"true"`
  // {'en':'total transcoding time', 'zh_CN':'总转码类型的转码时长(单位分钟，固定保留2位小数)'}
  Total *string `json:"total,omitempty" xml:"total,omitempty" require:"true"`
}

func (s QueryDailyLiveTranscodingDurationResponseProviderDateTranscodingLive) String() string {
  return tea.Prettify(s)
}

func (s QueryDailyLiveTranscodingDurationResponseProviderDateTranscodingLive) GoString() string {
  return s.String()
}

func (s *QueryDailyLiveTranscodingDurationResponseProviderDateTranscodingLive) SetTime(v string) *QueryDailyLiveTranscodingDurationResponseProviderDateTranscodingLive {
  s.Time = &v
  return s
}

func (s *QueryDailyLiveTranscodingDurationResponseProviderDateTranscodingLive) SetH264(v string) *QueryDailyLiveTranscodingDurationResponseProviderDateTranscodingLive {
  s.H264 = &v
  return s
}

func (s *QueryDailyLiveTranscodingDurationResponseProviderDateTranscodingLive) SetH265(v string) *QueryDailyLiveTranscodingDurationResponseProviderDateTranscodingLive {
  s.H265 = &v
  return s
}

func (s *QueryDailyLiveTranscodingDurationResponseProviderDateTranscodingLive) SetZdgq_264(v string) *QueryDailyLiveTranscodingDurationResponseProviderDateTranscodingLive {
  s.Zdgq_264 = &v
  return s
}

func (s *QueryDailyLiveTranscodingDurationResponseProviderDateTranscodingLive) SetZdgq_265(v string) *QueryDailyLiveTranscodingDurationResponseProviderDateTranscodingLive {
  s.Zdgq_265 = &v
  return s
}

func (s *QueryDailyLiveTranscodingDurationResponseProviderDateTranscodingLive) SetVoice(v string) *QueryDailyLiveTranscodingDurationResponseProviderDateTranscodingLive {
  s.Voice = &v
  return s
}

func (s *QueryDailyLiveTranscodingDurationResponseProviderDateTranscodingLive) SetTotal(v string) *QueryDailyLiveTranscodingDurationResponseProviderDateTranscodingLive {
  s.Total = &v
  return s
}

type QueryDailyLiveTranscodingDurationPaths struct {
}

func (s QueryDailyLiveTranscodingDurationPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryDailyLiveTranscodingDurationPaths) GoString() string {
  return s.String()
}

type QueryDailyLiveTranscodingDurationParameters struct {
}

func (s QueryDailyLiveTranscodingDurationParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryDailyLiveTranscodingDurationParameters) GoString() string {
  return s.String()
}

type QueryDailyLiveTranscodingDurationRequestHeader struct {
}

func (s QueryDailyLiveTranscodingDurationRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryDailyLiveTranscodingDurationRequestHeader) GoString() string {
  return s.String()
}

type QueryDailyLiveTranscodingDurationResponseHeader struct {
}

func (s QueryDailyLiveTranscodingDurationResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryDailyLiveTranscodingDurationResponseHeader) GoString() string {
  return s.String()
}




type ReportDomainStreamDurationServiceRequest struct {
  // {"en":"Start time:\n1.Format is yyyyMMdd;\n2.Must be smaller than the current system time;\n3.Default value of current time is used if the field is not specified;\n4.You can only query data for the last 6 months.","zh_CN":"开始时间:\n1.时间格式为yyyy-MM-dd,例如,2021-08-10\n2.不能大于当前时间\n3.只能查询最近半年内数据。"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty"`
  // {"en":"End Time:\n1. The time format is yyyy-MM-dd\n2. The end time must be greater than the start time. If the end time is greater than the current time, take the current time\n3. Both dateFrom and dateTo have not been passed, and the past 7 days are queried by default; if only one has not been passed, an exception will be thrown\n4. The maximum time interval allowed for query: 31 days, that is, the difference between dateFrom and dateTo cannot exceed 31 days","zh_CN":"结束时间:\n1.时间格式为yyyy-MM-dd\n2.结束时间需大于开始时间,结束时间如果大于当前时间,取当前时间\n3.dateFrom,dateTo二者都未传,默认查询过去的7天;如仅有一个未传,抛异常\n4.允许查询最大时间间隔:31天,即dateFrom和dateTo相差不能超过31天。"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty"`
  // {"en":"-","zh_CN":"-"}
  DomainStream []*ReportDomainStreamDurationServiceRequestDomainStream `json:"domainStream,omitempty" xml:"domainStream,omitempty" type:"Repeated"`
}

func (s ReportDomainStreamDurationServiceRequest) String() string {
  return tea.Prettify(s)
}

func (s ReportDomainStreamDurationServiceRequest) GoString() string {
  return s.String()
}

func (s *ReportDomainStreamDurationServiceRequest) SetDateFrom(v string) *ReportDomainStreamDurationServiceRequest {
  s.DateFrom = &v
  return s
}

func (s *ReportDomainStreamDurationServiceRequest) SetDateTo(v string) *ReportDomainStreamDurationServiceRequest {
  s.DateTo = &v
  return s
}

func (s *ReportDomainStreamDurationServiceRequest) SetDomainStream(v []*ReportDomainStreamDurationServiceRequestDomainStream) *ReportDomainStreamDurationServiceRequest {
  s.DomainStream = v
  return s
}

type ReportDomainStreamDurationServiceRequestDomainStream struct     {
  // {"en":"Domain name:\n1. The maximum number of domain names that can be transferred is 1 by default;\n2. Automatically filter out invalid domain names (if an illegal domain name is transferred, it will be filtered out, and the query results will only return the data of valid domain names).","zh_CN":"域名:\n1、可传递域名数量上限默认为1个;\n2、自动过滤掉无效域名(如传递非法域名,会被过滤掉,查询结果只返回有效域名的数据)。"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty"`
  // {"en":"stream name:\nJust pass the publishing point + stream name, Example: live/test-20180101-test where live is a publishing point and test-20180101-test is a stream name","zh_CN":"流名:\n只需要传发布点+流名,例如:live/test-20180101-test ,其中live是发布点，test-20180101-test是流名"}
  Stream []*string `json:"stream,omitempty" xml:"stream,omitempty" type:"Repeated"`
}

func (s ReportDomainStreamDurationServiceRequestDomainStream) String() string {
  return tea.Prettify(s)
}

func (s ReportDomainStreamDurationServiceRequestDomainStream) GoString() string {
  return s.String()
}

func (s *ReportDomainStreamDurationServiceRequestDomainStream) SetDomain(v string) *ReportDomainStreamDurationServiceRequestDomainStream {
  s.Domain = &v
  return s
}

func (s *ReportDomainStreamDurationServiceRequestDomainStream) SetStream(v []*string) *ReportDomainStreamDurationServiceRequestDomainStream {
  s.Stream = v
  return s
}

type ReportDomainStreamDurationServiceRequestHeader struct {
}

func (s ReportDomainStreamDurationServiceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportDomainStreamDurationServiceRequestHeader) GoString() string {
  return s.String()
}

type ReportDomainStreamDurationServicePaths struct {
}

func (s ReportDomainStreamDurationServicePaths) String() string {
  return tea.Prettify(s)
}

func (s ReportDomainStreamDurationServicePaths) GoString() string {
  return s.String()
}

type ReportDomainStreamDurationServiceParameters struct {
}

func (s ReportDomainStreamDurationServiceParameters) String() string {
  return tea.Prettify(s)
}

func (s ReportDomainStreamDurationServiceParameters) GoString() string {
  return s.String()
}

type ReportDomainStreamDurationServiceResponse struct {
  // {"en":"request result status code","zh_CN":"请求结果状态码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Request result information","zh_CN":"请求结果信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Detailed data on the result of the request","zh_CN":"请求结果的详细数据"}
  Data []*ReportDomainStreamDurationServiceResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s ReportDomainStreamDurationServiceResponse) String() string {
  return tea.Prettify(s)
}

func (s ReportDomainStreamDurationServiceResponse) GoString() string {
  return s.String()
}

func (s *ReportDomainStreamDurationServiceResponse) SetCode(v string) *ReportDomainStreamDurationServiceResponse {
  s.Code = &v
  return s
}

func (s *ReportDomainStreamDurationServiceResponse) SetMessage(v string) *ReportDomainStreamDurationServiceResponse {
  s.Message = &v
  return s
}

func (s *ReportDomainStreamDurationServiceResponse) SetData(v []*ReportDomainStreamDurationServiceResponseData) *ReportDomainStreamDurationServiceResponse {
  s.Data = v
  return s
}

type ReportDomainStreamDurationServiceResponseData struct     {
  // {"en":"Domain","zh_CN":"域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {"en":"-","zh_CN":"-"}
  StreamList []*ReportDomainStreamDurationServiceResponseDataStreamList `json:"streamList,omitempty" xml:"streamList,omitempty" require:"true" type:"Repeated"`
}

func (s ReportDomainStreamDurationServiceResponseData) String() string {
  return tea.Prettify(s)
}

func (s ReportDomainStreamDurationServiceResponseData) GoString() string {
  return s.String()
}

func (s *ReportDomainStreamDurationServiceResponseData) SetDomain(v string) *ReportDomainStreamDurationServiceResponseData {
  s.Domain = &v
  return s
}

func (s *ReportDomainStreamDurationServiceResponseData) SetStreamList(v []*ReportDomainStreamDurationServiceResponseDataStreamList) *ReportDomainStreamDurationServiceResponseData {
  s.StreamList = v
  return s
}

type ReportDomainStreamDurationServiceResponseDataStreamList struct     {
  // {"en":"domain + publishing point + stream name","zh_CN":"流名(域名+发布点+流名)"}
  Stream *string `json:"stream,omitempty" xml:"stream,omitempty" require:"true"`
  // {"en":"The sum of the flow duration of the flow name in the corresponding time period, in milliseconds","zh_CN":"对应时间段内流名推流时长之和,单位为毫秒"}
  SumTime *ReportDomainStreamDurationServiceResponseDataStreamListSumTime `json:"sumTime,omitempty" xml:"sumTime,omitempty" require:"true" type:"Struct"`
  // {"en":"-","zh_CN":"-"}
  DurationDetailList []*ReportDomainStreamDurationServiceResponseDataStreamListDurationDetailList `json:"durationDetailList,omitempty" xml:"durationDetailList,omitempty" require:"true" type:"Repeated"`
}

func (s ReportDomainStreamDurationServiceResponseDataStreamList) String() string {
  return tea.Prettify(s)
}

func (s ReportDomainStreamDurationServiceResponseDataStreamList) GoString() string {
  return s.String()
}

func (s *ReportDomainStreamDurationServiceResponseDataStreamList) SetStream(v string) *ReportDomainStreamDurationServiceResponseDataStreamList {
  s.Stream = &v
  return s
}

func (s *ReportDomainStreamDurationServiceResponseDataStreamList) SetSumTime(v *ReportDomainStreamDurationServiceResponseDataStreamListSumTime) *ReportDomainStreamDurationServiceResponseDataStreamList {
  s.SumTime = v
  return s
}

func (s *ReportDomainStreamDurationServiceResponseDataStreamList) SetDurationDetailList(v []*ReportDomainStreamDurationServiceResponseDataStreamListDurationDetailList) *ReportDomainStreamDurationServiceResponseDataStreamList {
  s.DurationDetailList = v
  return s
}

type ReportDomainStreamDurationServiceResponseDataStreamListSumTime struct {
}

func (s ReportDomainStreamDurationServiceResponseDataStreamListSumTime) String() string {
  return tea.Prettify(s)
}

func (s ReportDomainStreamDurationServiceResponseDataStreamListSumTime) GoString() string {
  return s.String()
}

type ReportDomainStreamDurationServiceResponseDataStreamListDurationDetailList struct     {
  // {"en":"Stream start time","zh_CN":"推流起始时间"}
  StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty" require:"true"`
  // {"en":"Stream end time","zh_CN":"推流终止时间"}
  EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty" require:"true"`
  // {"en":"Streaming duration, in milliseconds","zh_CN":"推流时长,单位为毫秒"}
  Duration *ReportDomainStreamDurationServiceResponseDataStreamListDurationDetailListDuration `json:"duration,omitempty" xml:"duration,omitempty" require:"true" type:"Struct"`
}

func (s ReportDomainStreamDurationServiceResponseDataStreamListDurationDetailList) String() string {
  return tea.Prettify(s)
}

func (s ReportDomainStreamDurationServiceResponseDataStreamListDurationDetailList) GoString() string {
  return s.String()
}

func (s *ReportDomainStreamDurationServiceResponseDataStreamListDurationDetailList) SetStartTime(v string) *ReportDomainStreamDurationServiceResponseDataStreamListDurationDetailList {
  s.StartTime = &v
  return s
}

func (s *ReportDomainStreamDurationServiceResponseDataStreamListDurationDetailList) SetEndTime(v string) *ReportDomainStreamDurationServiceResponseDataStreamListDurationDetailList {
  s.EndTime = &v
  return s
}

func (s *ReportDomainStreamDurationServiceResponseDataStreamListDurationDetailList) SetDuration(v *ReportDomainStreamDurationServiceResponseDataStreamListDurationDetailListDuration) *ReportDomainStreamDurationServiceResponseDataStreamListDurationDetailList {
  s.Duration = v
  return s
}

type ReportDomainStreamDurationServiceResponseDataStreamListDurationDetailListDuration struct {
}

func (s ReportDomainStreamDurationServiceResponseDataStreamListDurationDetailListDuration) String() string {
  return tea.Prettify(s)
}

func (s ReportDomainStreamDurationServiceResponseDataStreamListDurationDetailListDuration) GoString() string {
  return s.String()
}

type ReportDomainStreamDurationServiceResponseHeader struct {
}

func (s ReportDomainStreamDurationServiceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ReportDomainStreamDurationServiceResponseHeader) GoString() string {
  return s.String()
}




type QueryDirectoryRankbyFailedRequestHourlyRequest struct {
  // {"en":"Start time:
  // 1. Time format is yyyy-MM-ddTHH:mm:ss+08:00,
  // 2. No bigger than the current time.
  // 3. Data in the last 183 days at most can be queried.", "zh_CN":"开始时间:
  // 1.时间格式为yyyy-MM-ddTHH:mm:ss+08:00
  // 2.不能大于当前时间
  // 3.最多可获取最近半年(183天)的数据"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty"`
  // {"en":"End time:
  // 1. the time format is yyyy-MM-ddTHH:mm:ss+08:00
  // 2. End time should be greater than start time. If the end time is greater than current time, current time will be used.
  // 3. If both fields of dataFrom and dateTo are left empty, then data in the last 1day will be queried by default; if only one field is filled in and one is left empty, then exception will be occur.
  // 4. Allowable maximum time range for query: 31days, means the period between dateFrom to dateTo should not exceed 31days (can be adjusted by contacting technical support).", "zh_CN":"结束时间:
  // 
  // 1.时间格式yyyy-MM-ddTHH:mm:ss+08:00
  // 2.结束时间需大于开始时间,结束时间如果大于当前时间,取当前时间
  // 3.dateFrom,dateTo二者都未传,默认查询过去的1天;如仅有一个未传,抛异常
  // 4.允许查询最大时间间隔31天:,即dateFrom和dateTo相差不能超过31天。(可联系技术支持调整)"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty"`
  // {"en":"Domain:
  // 1. The maximum number of domain is 100 by default (you can contact technical support for adjustment);
  // 2. Automatically filter out invalid domain (an illegal domain will be filtered, and the query result will only return the data of valid domains).
  // 3. If the input parameter is not inputed, all domains and multipaths under the account will be queried by default, but an error will be prompted when the number of domains under the account exceeds the upper limit", "zh_CN":"域名:
  // 1.可传递域名数量上限默认为100个(可联系技术支持调整);
  // 2.自动过滤掉无效域名(如传递非法域名,会被过滤,查询结果只返回有效域名的数据)。
  // 3.未传递该入参时,默认查询账号下所有域名和multipath,但当账号下域名数量超过上限时提示错误"}
  Domain []*string `json:"domain,omitempty" xml:"domain,omitempty" type:"Repeated"`
}

func (s QueryDirectoryRankbyFailedRequestHourlyRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryDirectoryRankbyFailedRequestHourlyRequest) GoString() string {
  return s.String()
}

func (s *QueryDirectoryRankbyFailedRequestHourlyRequest) SetDateFrom(v string) *QueryDirectoryRankbyFailedRequestHourlyRequest {
  s.DateFrom = &v
  return s
}

func (s *QueryDirectoryRankbyFailedRequestHourlyRequest) SetDateTo(v string) *QueryDirectoryRankbyFailedRequestHourlyRequest {
  s.DateTo = &v
  return s
}

func (s *QueryDirectoryRankbyFailedRequestHourlyRequest) SetDomain(v []*string) *QueryDirectoryRankbyFailedRequestHourlyRequest {
  s.Domain = v
  return s
}

type QueryDirectoryRankbyFailedRequestHourlyResponse struct {
  // {"en":"Top ranking", "zh_CN":"top排名"}
  Top *string `json:"top,omitempty" xml:"top,omitempty" require:"true"`
  // {"en":"Directory", "zh_CN":"目录"}
  Directory *string `json:"directory,omitempty" xml:"directory,omitempty" require:"true"`
  // {"en":"Number of failed requests", "zh_CN":"失败请求数"}
  FailRequests *string `json:"failRequests,omitempty" xml:"failRequests,omitempty" require:"true"`
  // {"en":"Total successful flow: the unit of measurement is MB, with 2 decimal places reserved", "zh_CN":"成功总流量:计量单位MB,保留2位小数"}
  SuccessTotalFlow *string `json:"successTotalFlow,omitempty" xml:"successTotalFlow,omitempty" require:"true"`
  // {"en":"Number of successful requests", "zh_CN":"成功请求数"}
  SuccessRequests *string `json:"successRequests,omitempty" xml:"successRequests,omitempty" require:"true"`
}

func (s QueryDirectoryRankbyFailedRequestHourlyResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryDirectoryRankbyFailedRequestHourlyResponse) GoString() string {
  return s.String()
}

func (s *QueryDirectoryRankbyFailedRequestHourlyResponse) SetTop(v string) *QueryDirectoryRankbyFailedRequestHourlyResponse {
  s.Top = &v
  return s
}

func (s *QueryDirectoryRankbyFailedRequestHourlyResponse) SetDirectory(v string) *QueryDirectoryRankbyFailedRequestHourlyResponse {
  s.Directory = &v
  return s
}

func (s *QueryDirectoryRankbyFailedRequestHourlyResponse) SetFailRequests(v string) *QueryDirectoryRankbyFailedRequestHourlyResponse {
  s.FailRequests = &v
  return s
}

func (s *QueryDirectoryRankbyFailedRequestHourlyResponse) SetSuccessTotalFlow(v string) *QueryDirectoryRankbyFailedRequestHourlyResponse {
  s.SuccessTotalFlow = &v
  return s
}

func (s *QueryDirectoryRankbyFailedRequestHourlyResponse) SetSuccessRequests(v string) *QueryDirectoryRankbyFailedRequestHourlyResponse {
  s.SuccessRequests = &v
  return s
}

type QueryDirectoryRankbyFailedRequestHourlyPaths struct {
}

func (s QueryDirectoryRankbyFailedRequestHourlyPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryDirectoryRankbyFailedRequestHourlyPaths) GoString() string {
  return s.String()
}

type QueryDirectoryRankbyFailedRequestHourlyParameters struct {
}

func (s QueryDirectoryRankbyFailedRequestHourlyParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryDirectoryRankbyFailedRequestHourlyParameters) GoString() string {
  return s.String()
}

type QueryDirectoryRankbyFailedRequestHourlyRequestHeader struct {
}

func (s QueryDirectoryRankbyFailedRequestHourlyRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryDirectoryRankbyFailedRequestHourlyRequestHeader) GoString() string {
  return s.String()
}

type QueryDirectoryRankbyFailedRequestHourlyResponseHeader struct {
}

func (s QueryDirectoryRankbyFailedRequestHourlyResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryDirectoryRankbyFailedRequestHourlyResponseHeader) GoString() string {
  return s.String()
}




type QueryDDoSMitigatedBandwidthBySuiteOrProductRequest struct {
  // {"en":"Start time, yyyy-MM-dd HH:mm:ss.", "zh_CN":"开始时间，yyyy-MM-dd HH:mm:ss。"}
  Startdate *string `json:"startdate,omitempty" xml:"startdate,omitempty" require:"true"`
  // {"en":"PackageId or acctype should be selected at least one.", "zh_CN":"套餐ID: packageId和acctype至少传一个,但不能同时传。"}
  PackageId *string `json:"packageId,omitempty" xml:"packageId,omitempty"`
  // {"en":"need Detail : 1
  // no need Detail: 0
  // default : 1.", "zh_CN":"是否需要查看域名或是转发规则带宽的详细信息：0：不需要；1：需要，默认需要。"}
  NeedDetail *int `json:"needDetail,omitempty" xml:"needDetail,omitempty"`
  // {"en":"End time, yyyy-MM-dd HH:mm:ss.", "zh_CN":"结束时间，yyyy-MM-dd HH:mm:ss。"}
  Enddate *string `json:"enddate,omitempty" xml:"enddate,omitempty" require:"true"`
  // {"en":"Customer English Name.", "zh_CN":"客户英文名。"}
  CustomCode *string `json:"customCode,omitempty" xml:"customCode,omitempty" require:"true"`
  // {"en":"PackageId or acctype should be selected at least one.
  // acctype( Only One can be selected): gess, fsa, app-s, dms-https, wss, dms, wss-https, s-appa, esa, wsa-https, 1551.", "zh_CN":"PackageId和acctype不能同时传且至少传一个；产品外部服务类型,只支持传1个:gess，fsa，app-s，dms-https，wss, dms， wss-https，s-appa，esa，wsa-https，1551。"}
  Acctype *string `json:"acctype,omitempty" xml:"acctype,omitempty"`
  // {"en":"ContractId", "zh_CN":"合同号  
  // 支持按Contract# item#粒度查询
  // 当ContractId不为空时，会替换掉packageId"}
  ContractId *string `json:"ContractId,omitempty" xml:"ContractId,omitempty"`
}

func (s QueryDDoSMitigatedBandwidthBySuiteOrProductRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryDDoSMitigatedBandwidthBySuiteOrProductRequest) GoString() string {
  return s.String()
}

func (s *QueryDDoSMitigatedBandwidthBySuiteOrProductRequest) SetStartdate(v string) *QueryDDoSMitigatedBandwidthBySuiteOrProductRequest {
  s.Startdate = &v
  return s
}

func (s *QueryDDoSMitigatedBandwidthBySuiteOrProductRequest) SetPackageId(v string) *QueryDDoSMitigatedBandwidthBySuiteOrProductRequest {
  s.PackageId = &v
  return s
}

func (s *QueryDDoSMitigatedBandwidthBySuiteOrProductRequest) SetNeedDetail(v int) *QueryDDoSMitigatedBandwidthBySuiteOrProductRequest {
  s.NeedDetail = &v
  return s
}

func (s *QueryDDoSMitigatedBandwidthBySuiteOrProductRequest) SetEnddate(v string) *QueryDDoSMitigatedBandwidthBySuiteOrProductRequest {
  s.Enddate = &v
  return s
}

func (s *QueryDDoSMitigatedBandwidthBySuiteOrProductRequest) SetCustomCode(v string) *QueryDDoSMitigatedBandwidthBySuiteOrProductRequest {
  s.CustomCode = &v
  return s
}

func (s *QueryDDoSMitigatedBandwidthBySuiteOrProductRequest) SetAcctype(v string) *QueryDDoSMitigatedBandwidthBySuiteOrProductRequest {
  s.Acctype = &v
  return s
}

func (s *QueryDDoSMitigatedBandwidthBySuiteOrProductRequest) SetContractId(v string) *QueryDDoSMitigatedBandwidthBySuiteOrProductRequest {
  s.ContractId = &v
  return s
}

type QueryDDoSMitigatedBandwidthBySuiteOrProductResponse struct {
  // {"en":"错误信息或Success。", "zh_CN":"Error message or Success."}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"CleanBandwidth:
  //     time: "2018-08-21 00:00:00"
  //    atFlow: cleaned bandwidth Mbps", "zh_CN":"带宽信息： 
  // Time："2018-08-21 00:00:00"
  // atFlow：已清洗的带宽　Mbps"}
  At_bw *string `json:"at_bw,omitempty" xml:"at_bw,omitempty" require:"true"`
  // {"en":"peak information:
  //    peakTime: "2018-08-21 00:00:00"
  //    peakValue:  peak of Clean Bandwidth (Mbps)
  //    mitigatedTraffic: Cleaned flow (GB)", "zh_CN":"峰值统计信息：
  // peakTime：峰值时间 "2018-08-21 00:00:00"
  // peakValue：DDoS攻击峰值带宽Mbps
  // mitigatedTraffic：已清洗的流量 GB"}
  PeakStat *string `json:"peakStat,omitempty" xml:"peakStat,omitempty" require:"true"`
  // {"en":"Return 200 means success, please see <Error code> to check other status code.", "zh_CN":"200状态码表示请求成功，其他状态码说明请参见《错误码》。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
}

func (s QueryDDoSMitigatedBandwidthBySuiteOrProductResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryDDoSMitigatedBandwidthBySuiteOrProductResponse) GoString() string {
  return s.String()
}

func (s *QueryDDoSMitigatedBandwidthBySuiteOrProductResponse) SetMsg(v string) *QueryDDoSMitigatedBandwidthBySuiteOrProductResponse {
  s.Msg = &v
  return s
}

func (s *QueryDDoSMitigatedBandwidthBySuiteOrProductResponse) SetAt_bw(v string) *QueryDDoSMitigatedBandwidthBySuiteOrProductResponse {
  s.At_bw = &v
  return s
}

func (s *QueryDDoSMitigatedBandwidthBySuiteOrProductResponse) SetPeakStat(v string) *QueryDDoSMitigatedBandwidthBySuiteOrProductResponse {
  s.PeakStat = &v
  return s
}

func (s *QueryDDoSMitigatedBandwidthBySuiteOrProductResponse) SetCode(v string) *QueryDDoSMitigatedBandwidthBySuiteOrProductResponse {
  s.Code = &v
  return s
}

type QueryDDoSMitigatedBandwidthBySuiteOrProductPaths struct {
}

func (s QueryDDoSMitigatedBandwidthBySuiteOrProductPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryDDoSMitigatedBandwidthBySuiteOrProductPaths) GoString() string {
  return s.String()
}

type QueryDDoSMitigatedBandwidthBySuiteOrProductParameters struct {
}

func (s QueryDDoSMitigatedBandwidthBySuiteOrProductParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryDDoSMitigatedBandwidthBySuiteOrProductParameters) GoString() string {
  return s.String()
}

type QueryDDoSMitigatedBandwidthBySuiteOrProductRequestHeader struct {
}

func (s QueryDDoSMitigatedBandwidthBySuiteOrProductRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryDDoSMitigatedBandwidthBySuiteOrProductRequestHeader) GoString() string {
  return s.String()
}

type QueryDDoSMitigatedBandwidthBySuiteOrProductResponseHeader struct {
}

func (s QueryDDoSMitigatedBandwidthBySuiteOrProductResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryDDoSMitigatedBandwidthBySuiteOrProductResponseHeader) GoString() string {
  return s.String()
}




type QueryDirectoryRankByTotalTrafficHourlyRequest struct {
  // {"en":"Start time:
  // 1. The format is yyyyy-MM-ddTHH: mm: SS + 08:00, for example, 2016-12-02T10:00 + 08:00 (10:0:00 Beijing time on December 2, 2016);
  // 2. Can not exceed the current time;
  // 3. The latest half year (183 days) data can be obtained at most.", "zh_CN":"开始时间：
  // 1.格式为yyyy-MM-ddTHH:mm:ss+08:00,例如,2016-12-02T10:00:00+08:00（为北京时间2016年12月2日10点0分0秒）;
  // 2.不能大于当前时间;
  // 3.最多可获取最近半年（183天）的数据。"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty"`
  // {"en":"End time:
  // 1. The format is yyyy-MM-ddTHH:mm:ss+08:00;
  // 2. The end time is greater than the start time.
  // 3. If the end time is greater than the current time, the current time is taken.
  // 4. DateFrom and dateTo are not uploaded, defaulting to query the past 24 hours; if only one is not uploaded, throw an exception;
  // 5. Maximum query interval allowed:31 days, that is, the difference between dateFrom and dateTo can not exceed 31 days. ", "zh_CN":"结束时间：
  // 1.格式为yyyy-MM-ddTHH:mm:ss+08:00;
  // 2.结束时间需大于开始时间;
  // 3.结束时间如果大于当前时间,取当前时间;
  // 4.dateFrom, dateTo二者都未传,默认查询过去的24小时;如仅有一个未传,抛异常;
  // 5.允许查询最大间隔：31天,即dateFrom和dateTo相差不能超过31天。"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty"`
  // {"en":"Domain:
  // 1. The maximum number of domain is 100 by default (you can contact technical support for adjustment);
  // 2. Automatically filter out invalid domain (an illegal domain will be filtered, and the query result will only return the data of valid domains).
  // 3. If the input parameter is not inputed, all domains and multipaths under the account will be queried by default, but an error will be prompted when the number of domains under the account exceeds the upper limit", "zh_CN":"域名:
  // 1.可传递域名数量上限默认为100个（可联系技术支持调整）;
  // 2.自动过滤掉无效域名（如传递非法域名,会被过滤,查询结果只返回有效域名的数据）。
  // 3.未传递该入参时,默认查询账号下所有域名和multipath,但当账号下域名数量超过上限时提示错误"}
  Domain []*string `json:"domain,omitempty" xml:"domain,omitempty" type:"Repeated"`
}

func (s QueryDirectoryRankByTotalTrafficHourlyRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryDirectoryRankByTotalTrafficHourlyRequest) GoString() string {
  return s.String()
}

func (s *QueryDirectoryRankByTotalTrafficHourlyRequest) SetDateFrom(v string) *QueryDirectoryRankByTotalTrafficHourlyRequest {
  s.DateFrom = &v
  return s
}

func (s *QueryDirectoryRankByTotalTrafficHourlyRequest) SetDateTo(v string) *QueryDirectoryRankByTotalTrafficHourlyRequest {
  s.DateTo = &v
  return s
}

func (s *QueryDirectoryRankByTotalTrafficHourlyRequest) SetDomain(v []*string) *QueryDirectoryRankByTotalTrafficHourlyRequest {
  s.Domain = v
  return s
}

type QueryDirectoryRankByTotalTrafficHourlyResponse struct {
  // {"en":"Top ranking", "zh_CN":"top排行"}
  Top *string `json:"top,omitempty" xml:"top,omitempty" require:"true"`
  // {"en":"Directory", "zh_CN":"目录"}
  Directory *string `json:"directory,omitempty" xml:"directory,omitempty" require:"true"`
  // {"en":"Total traffic, unit is MB, keep two decimal places", "zh_CN":"总流量：计量单位MB,保留2位小数"}
  TotalFlow *string `json:"totalFlow,omitempty" xml:"totalFlow,omitempty" require:"true"`
  // {"en":"Requests", "zh_CN":"请求数"}
  Requests *string `json:"requests,omitempty" xml:"requests,omitempty" require:"true"`
}

func (s QueryDirectoryRankByTotalTrafficHourlyResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryDirectoryRankByTotalTrafficHourlyResponse) GoString() string {
  return s.String()
}

func (s *QueryDirectoryRankByTotalTrafficHourlyResponse) SetTop(v string) *QueryDirectoryRankByTotalTrafficHourlyResponse {
  s.Top = &v
  return s
}

func (s *QueryDirectoryRankByTotalTrafficHourlyResponse) SetDirectory(v string) *QueryDirectoryRankByTotalTrafficHourlyResponse {
  s.Directory = &v
  return s
}

func (s *QueryDirectoryRankByTotalTrafficHourlyResponse) SetTotalFlow(v string) *QueryDirectoryRankByTotalTrafficHourlyResponse {
  s.TotalFlow = &v
  return s
}

func (s *QueryDirectoryRankByTotalTrafficHourlyResponse) SetRequests(v string) *QueryDirectoryRankByTotalTrafficHourlyResponse {
  s.Requests = &v
  return s
}

type QueryDirectoryRankByTotalTrafficHourlyPaths struct {
}

func (s QueryDirectoryRankByTotalTrafficHourlyPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryDirectoryRankByTotalTrafficHourlyPaths) GoString() string {
  return s.String()
}

type QueryDirectoryRankByTotalTrafficHourlyParameters struct {
}

func (s QueryDirectoryRankByTotalTrafficHourlyParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryDirectoryRankByTotalTrafficHourlyParameters) GoString() string {
  return s.String()
}

type QueryDirectoryRankByTotalTrafficHourlyRequestHeader struct {
}

func (s QueryDirectoryRankByTotalTrafficHourlyRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryDirectoryRankByTotalTrafficHourlyRequestHeader) GoString() string {
  return s.String()
}

type QueryDirectoryRankByTotalTrafficHourlyResponseHeader struct {
}

func (s QueryDirectoryRankByTotalTrafficHourlyResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryDirectoryRankByTotalTrafficHourlyResponseHeader) GoString() string {
  return s.String()
}




type Query5minLiveTranscodingDurationRequest struct {
  // {"en":"cust_en_name of sub-client.
  // When a merged-account wants to  view the information of the subclient,the cust_en_name is required.", "zh_CN":"合并账号下的某个客户的英文名，当合并账号要查看子客户的信息时，必须填写子客户的英文名"}
  Cust *string `json:"cust,omitempty" xml:"cust,omitempty"`
  // {"en":"Specifies the query date:
  // 1.With format yyyy-mm-dd.
  // 2.If not specified,it means today as default.", "zh_CN":"查询的日期，日期格式为yyyy-mm-dd,不选或者为空时默认为当天；"}
  Date *string `json:"date,omitempty" xml:"date,omitempty"`
  // {"en":"1.Must work with 'enddate' and they  specify the query date scope.
  // 2.With format yyyy-mm-dd.
  // 3.If there is a 'date' parameter,this parameter will be omitted.", "zh_CN":"查询的起始日期,日期格式为yyyy-mm-dd；此参数需与enddate参数配合,若存在date参数,则该参数无效"}
  Startdate *string `json:"startdate,omitempty" xml:"startdate,omitempty"`
  // {"en":"1.Must work with 'startdate' and they  specify the query date scope.
  // 2.With format yyyy-mm-dd.
  // 3.If there is a 'date' parameter,this parameter will be omitted.", "zh_CN":"查询的结束日期,日期格式为yyyy-mm-dd；此参数需与startdate参数配合,若存在date参数,则该参数无效。"}
  Enddate *string `json:"enddate,omitempty" xml:"enddate,omitempty"`
  // {"en":"GMT time zone, parameter format: GMT+09:00 means east 9th zone, GMT-09:00 means west 9th zone, if not transmitted, the default is local time zone (east 8th zone).", "zh_CN":"格林尼治时区，参数格式 GMT+09:00 表示东九区，GMT-09:00 表示西9区，不传则默认为本地时区（东八区）"}
  Timezone *string `json:"timezone,omitempty" xml:"timezone,omitempty"`
  // {"en":"domains that been queried:
  // 1.If there are multiple inputs,use  ';' as separator.
  // 2.If not specified, it means all the domains of the account .", "zh_CN":"查询的频道，多个频道值请用英文分号';'，不选或者为空时默认为所查询客户的所有频道"}
  Channel *string `json:"channel,omitempty" xml:"channel,omitempty"`
  // {"en":"This parameter specifies if the 'channel' parameter should be exactly matched:
  // 1.'true' as default.
  // 2. If not 'true',it will query data of channels that ends with any item of input 'channel's.", "zh_CN":"&nbsp;频道是否完全匹配,为true时，必须填写完整的域名(此时会过滤用户输入的无效或重复频道,所有输入频道都无效时返403.。不为true时，显示以用户输入的频道为结尾的所有频道。默认为true"}
  IsExactMatch *string `json:"isExactMatch,omitempty" xml:"isExactMatch,omitempty"`
  // {"en":"acceleration type:
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
  // {"en":"Transcoding type, values can be h264, h265, zdgq_264, zdgq_265, cf_264, cf_265, or other. Multiple transcoding types should be separated by a semicolon. If some of the transcoding types are incorrect, the system will return data for the correct types; if all transcoding types are incorrect, it will return an error 'invalid transcodeType.' If not provided or left empty, it defaults to all types.", "zh_CN":"转码类型,值为h264、h265、zdgq_264、zdgq_265、cf_264、cf_265，other 多个转码类型用英文分号;分隔开。当传入转码类型部分错误时，返回正确的类型的数据；当传入转码类型全部错误时，返回错误invalid transcodeType. 不填或为空，默认为所有类型."}
  TranscodeType *string `json:"transcodeType,omitempty" xml:"transcodeType,omitempty"`
  // {"en":"Resolution types include LD480,SD720,HD1080,2K,4K,8K,SD576,SD540,LD360,LD240. Multiple resolutions are separated by a semicolon. When isAudio=1, this parameter is invalid and will return an error. Param definition must be empty when querying audio data.", "zh_CN":"清晰度类型,值为LD480,SD720,HD1080,2K,4K,8K,SD576,SD540,LD360,LD240，多个清晰度用英文分号;分隔开, 当isAudio=1时，此入参无效返回错误 param definition must be empty when query audio data."}
  Definition *string `json:"definition,omitempty" xml:"definition,omitempty"`
  // {"en":"Audio/Video Type, 1: Audio 2: Video. Defaults to 2 if not selected or empty. Only a single value is allowed.", "zh_CN":"音视频类型, 1:音频   2:视频. 不选或者为空时默认为2. 只能输入单个值."}
  IsAudio *string `json:"isAudio,omitempty" xml:"isAudio,omitempty"`
}

func (s Query5minLiveTranscodingDurationRequest) String() string {
  return tea.Prettify(s)
}

func (s Query5minLiveTranscodingDurationRequest) GoString() string {
  return s.String()
}

func (s *Query5minLiveTranscodingDurationRequest) SetCust(v string) *Query5minLiveTranscodingDurationRequest {
  s.Cust = &v
  return s
}

func (s *Query5minLiveTranscodingDurationRequest) SetDate(v string) *Query5minLiveTranscodingDurationRequest {
  s.Date = &v
  return s
}

func (s *Query5minLiveTranscodingDurationRequest) SetStartdate(v string) *Query5minLiveTranscodingDurationRequest {
  s.Startdate = &v
  return s
}

func (s *Query5minLiveTranscodingDurationRequest) SetEnddate(v string) *Query5minLiveTranscodingDurationRequest {
  s.Enddate = &v
  return s
}

func (s *Query5minLiveTranscodingDurationRequest) SetTimezone(v string) *Query5minLiveTranscodingDurationRequest {
  s.Timezone = &v
  return s
}

func (s *Query5minLiveTranscodingDurationRequest) SetChannel(v string) *Query5minLiveTranscodingDurationRequest {
  s.Channel = &v
  return s
}

func (s *Query5minLiveTranscodingDurationRequest) SetIsExactMatch(v string) *Query5minLiveTranscodingDurationRequest {
  s.IsExactMatch = &v
  return s
}

func (s *Query5minLiveTranscodingDurationRequest) SetAccetype(v string) *Query5minLiveTranscodingDurationRequest {
  s.Accetype = &v
  return s
}

func (s *Query5minLiveTranscodingDurationRequest) SetDataformat(v string) *Query5minLiveTranscodingDurationRequest {
  s.Dataformat = &v
  return s
}

func (s *Query5minLiveTranscodingDurationRequest) SetResultType(v string) *Query5minLiveTranscodingDurationRequest {
  s.ResultType = &v
  return s
}

func (s *Query5minLiveTranscodingDurationRequest) SetTranscodeType(v string) *Query5minLiveTranscodingDurationRequest {
  s.TranscodeType = &v
  return s
}

func (s *Query5minLiveTranscodingDurationRequest) SetDefinition(v string) *Query5minLiveTranscodingDurationRequest {
  s.Definition = &v
  return s
}

func (s *Query5minLiveTranscodingDurationRequest) SetIsAudio(v string) *Query5minLiveTranscodingDurationRequest {
  s.IsAudio = &v
  return s
}

type Query5minLiveTranscodingDurationResponse struct {
  // {'en':'provider', 'zh_CN':'结果'}
  Provider *Query5minLiveTranscodingDurationResponseProvider `json:"provider,omitempty" xml:"provider,omitempty" require:"true" type:"Struct"`
}

func (s Query5minLiveTranscodingDurationResponse) String() string {
  return tea.Prettify(s)
}

func (s Query5minLiveTranscodingDurationResponse) GoString() string {
  return s.String()
}

func (s *Query5minLiveTranscodingDurationResponse) SetProvider(v *Query5minLiveTranscodingDurationResponseProvider) *Query5minLiveTranscodingDurationResponse {
  s.Provider = v
  return s
}

type Query5minLiveTranscodingDurationResponseProvider struct {
  // {'en':'tenant', 'zh_CN':'租户'}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {'en':'type', 'zh_CN':'接口类型'}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
  // {'en':'data', 'zh_CN':'直播转码时长数据'}
  Date *Query5minLiveTranscodingDurationResponseProviderDate `json:"date,omitempty" xml:"date,omitempty" require:"true" type:"Struct"`
}

func (s Query5minLiveTranscodingDurationResponseProvider) String() string {
  return tea.Prettify(s)
}

func (s Query5minLiveTranscodingDurationResponseProvider) GoString() string {
  return s.String()
}

func (s *Query5minLiveTranscodingDurationResponseProvider) SetName(v string) *Query5minLiveTranscodingDurationResponseProvider {
  s.Name = &v
  return s
}

func (s *Query5minLiveTranscodingDurationResponseProvider) SetType(v string) *Query5minLiveTranscodingDurationResponseProvider {
  s.Type = &v
  return s
}

func (s *Query5minLiveTranscodingDurationResponseProvider) SetDate(v *Query5minLiveTranscodingDurationResponseProviderDate) *Query5minLiveTranscodingDurationResponseProvider {
  s.Date = v
  return s
}

type Query5minLiveTranscodingDurationResponseProviderDate struct {
  // {'en':'startdate', 'zh_CN':'开始时间'}
  Startdate *string `json:"startdate,omitempty" xml:"startdate,omitempty" require:"true"`
  // {'en':'enddate', 'zh_CN':'结束时间'}
  Enddate *string `json:"enddate,omitempty" xml:"enddate,omitempty" require:"true"`
  // {'en':'transcoding', 'zh_CN':'转码类型'}
  Transcoding *Query5minLiveTranscodingDurationResponseProviderDateTranscoding `json:"transcoding,omitempty" xml:"transcoding,omitempty" require:"true" type:"Struct"`
}

func (s Query5minLiveTranscodingDurationResponseProviderDate) String() string {
  return tea.Prettify(s)
}

func (s Query5minLiveTranscodingDurationResponseProviderDate) GoString() string {
  return s.String()
}

func (s *Query5minLiveTranscodingDurationResponseProviderDate) SetStartdate(v string) *Query5minLiveTranscodingDurationResponseProviderDate {
  s.Startdate = &v
  return s
}

func (s *Query5minLiveTranscodingDurationResponseProviderDate) SetEnddate(v string) *Query5minLiveTranscodingDurationResponseProviderDate {
  s.Enddate = &v
  return s
}

func (s *Query5minLiveTranscodingDurationResponseProviderDate) SetTranscoding(v *Query5minLiveTranscodingDurationResponseProviderDateTranscoding) *Query5minLiveTranscodingDurationResponseProviderDate {
  s.Transcoding = v
  return s
}

type Query5minLiveTranscodingDurationResponseProviderDateTranscoding struct {
  // {'en':'channel', 'zh_CN':'频道'}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {'en':'live', 'zh_CN':'直播转码时长数据'}
  Live []*Query5minLiveTranscodingDurationResponseProviderDateTranscodingLive `json:"live,omitempty" xml:"live,omitempty" require:"true" type:"Repeated"`
}

func (s Query5minLiveTranscodingDurationResponseProviderDateTranscoding) String() string {
  return tea.Prettify(s)
}

func (s Query5minLiveTranscodingDurationResponseProviderDateTranscoding) GoString() string {
  return s.String()
}

func (s *Query5minLiveTranscodingDurationResponseProviderDateTranscoding) SetName(v string) *Query5minLiveTranscodingDurationResponseProviderDateTranscoding {
  s.Name = &v
  return s
}

func (s *Query5minLiveTranscodingDurationResponseProviderDateTranscoding) SetLive(v []*Query5minLiveTranscodingDurationResponseProviderDateTranscodingLive) *Query5minLiveTranscodingDurationResponseProviderDateTranscoding {
  s.Live = v
  return s
}

type Query5minLiveTranscodingDurationResponseProviderDateTranscodingLive struct     {
  // {'en':'timestamp', 'zh_CN':'时间点'}
  Time *string `json:"time,omitempty" xml:"time,omitempty" require:"true"`
  // {'en':'h264 transcoding time', 'zh_CN':'h264转码类型的转码时长(单位分钟，固定保留2位小数)'}
  H264 *string `json:"h264,omitempty" xml:"h264,omitempty" require:"true"`
  // {'en':'h265 transcoding time', 'zh_CN':'h265转码类型的转码时长(单位分钟，固定保留2位小数)'}
  H265 *string `json:"h265,omitempty" xml:"h265,omitempty" require:"true"`
  // {'en':'zdgq_264 transcoding time', 'zh_CN':'zdgq_264转码类型的转码时长(单位分钟，固定保留2位小数)'}
  Zdgq_264 *string `json:"zdgq_264,omitempty" xml:"zdgq_264,omitempty" require:"true"`
  // {'en':'zdgq_265 transcoding time', 'zh_CN':'zdgq_265转码类型的转码时长(单位分钟，固定保留2位小数)'}
  Zdgq_265 *string `json:"zdgq_265,omitempty" xml:"zdgq_265,omitempty" require:"true"`
  // {'en':'voice transcoding time', 'zh_CN':'voice转码类型的转码时长(单位分钟，固定保留2位小数)'}
  Voice *string `json:"voice,omitempty" xml:"voice,omitempty" require:"true"`
  // {'en':'cf_264 transcoding time', 'zh_CN':'cf_264转码类型的转码时长(单位分钟，固定保留2位小数)'}
  Cf_264 *string `json:"cf_264,omitempty" xml:"cf_264,omitempty" require:"true"`
  // {'en':'cf_265 transcoding time', 'zh_CN':'cf_265转码类型的转码时长(单位分钟，固定保留2位小数)'}
  Cf_265 *string `json:"cf_265,omitempty" xml:"cf_265,omitempty" require:"true"`
  // {'en':'definition_2K transcoding time', 'zh_CN':'definition_2K转码类型的转码时长(单位分钟，固定保留2位小数)'}
  Definition_2K *string `json:"definition_2K,omitempty" xml:"definition_2K,omitempty" require:"true"`
  // {'en':'definition_4K transcoding time', 'zh_CN':'definition_4K转码类型的转码时长(单位分钟，固定保留2位小数)'}
  Definition_4K *string `json:"definition_4K,omitempty" xml:"definition_4K,omitempty" require:"true"`
  // {'en':'definition_8K transcoding time', 'zh_CN':'definition_8K转码类型的转码时长(单位分钟，固定保留2位小数)'}
  Definition_8K *string `json:"definition_8K,omitempty" xml:"definition_8K,omitempty" require:"true"`
  // {'en':'definition_LD480 transcoding time', 'zh_CN':'definition_LD480转码类型的转码时长(单位分钟，固定保留2位小数)'}
  Definition_LD480 *string `json:"definition_LD480,omitempty" xml:"definition_LD480,omitempty" require:"true"`
  // {'en':'definition_SD720 transcoding time', 'zh_CN':'definition_SD720转码类型的转码时长(单位分钟，固定保留2位小数)'}
  Definition_SD720 *string `json:"definition_SD720,omitempty" xml:"definition_SD720,omitempty" require:"true"`
  // {'en':'definition_HD1080 transcoding time', 'zh_CN':'definition_HD1080转码类型的转码时长(单位分钟，固定保留2位小数)'}
  Definition_HD1080 *string `json:"definition_HD1080,omitempty" xml:"definition_HD1080,omitempty" require:"true"`
  // {'en':'definition_SD576 transcoding time', 'zh_CN':'definition_SD576转码类型的转码时长(单位分钟，固定保留2位小数)'}
  Definition_SD576 *string `json:"definition_SD576,omitempty" xml:"definition_SD576,omitempty" require:"true"`
  // {'en':'total transcoding time', 'zh_CN':'总转码类型的转码时长(单位分钟，固定保留2位小数)'}
  Total *string `json:"total,omitempty" xml:"total,omitempty" require:"true"`
}

func (s Query5minLiveTranscodingDurationResponseProviderDateTranscodingLive) String() string {
  return tea.Prettify(s)
}

func (s Query5minLiveTranscodingDurationResponseProviderDateTranscodingLive) GoString() string {
  return s.String()
}

func (s *Query5minLiveTranscodingDurationResponseProviderDateTranscodingLive) SetTime(v string) *Query5minLiveTranscodingDurationResponseProviderDateTranscodingLive {
  s.Time = &v
  return s
}

func (s *Query5minLiveTranscodingDurationResponseProviderDateTranscodingLive) SetH264(v string) *Query5minLiveTranscodingDurationResponseProviderDateTranscodingLive {
  s.H264 = &v
  return s
}

func (s *Query5minLiveTranscodingDurationResponseProviderDateTranscodingLive) SetH265(v string) *Query5minLiveTranscodingDurationResponseProviderDateTranscodingLive {
  s.H265 = &v
  return s
}

func (s *Query5minLiveTranscodingDurationResponseProviderDateTranscodingLive) SetZdgq_264(v string) *Query5minLiveTranscodingDurationResponseProviderDateTranscodingLive {
  s.Zdgq_264 = &v
  return s
}

func (s *Query5minLiveTranscodingDurationResponseProviderDateTranscodingLive) SetZdgq_265(v string) *Query5minLiveTranscodingDurationResponseProviderDateTranscodingLive {
  s.Zdgq_265 = &v
  return s
}

func (s *Query5minLiveTranscodingDurationResponseProviderDateTranscodingLive) SetVoice(v string) *Query5minLiveTranscodingDurationResponseProviderDateTranscodingLive {
  s.Voice = &v
  return s
}

func (s *Query5minLiveTranscodingDurationResponseProviderDateTranscodingLive) SetCf_264(v string) *Query5minLiveTranscodingDurationResponseProviderDateTranscodingLive {
  s.Cf_264 = &v
  return s
}

func (s *Query5minLiveTranscodingDurationResponseProviderDateTranscodingLive) SetCf_265(v string) *Query5minLiveTranscodingDurationResponseProviderDateTranscodingLive {
  s.Cf_265 = &v
  return s
}

func (s *Query5minLiveTranscodingDurationResponseProviderDateTranscodingLive) SetDefinition_2K(v string) *Query5minLiveTranscodingDurationResponseProviderDateTranscodingLive {
  s.Definition_2K = &v
  return s
}

func (s *Query5minLiveTranscodingDurationResponseProviderDateTranscodingLive) SetDefinition_4K(v string) *Query5minLiveTranscodingDurationResponseProviderDateTranscodingLive {
  s.Definition_4K = &v
  return s
}

func (s *Query5minLiveTranscodingDurationResponseProviderDateTranscodingLive) SetDefinition_8K(v string) *Query5minLiveTranscodingDurationResponseProviderDateTranscodingLive {
  s.Definition_8K = &v
  return s
}

func (s *Query5minLiveTranscodingDurationResponseProviderDateTranscodingLive) SetDefinition_LD480(v string) *Query5minLiveTranscodingDurationResponseProviderDateTranscodingLive {
  s.Definition_LD480 = &v
  return s
}

func (s *Query5minLiveTranscodingDurationResponseProviderDateTranscodingLive) SetDefinition_SD720(v string) *Query5minLiveTranscodingDurationResponseProviderDateTranscodingLive {
  s.Definition_SD720 = &v
  return s
}

func (s *Query5minLiveTranscodingDurationResponseProviderDateTranscodingLive) SetDefinition_HD1080(v string) *Query5minLiveTranscodingDurationResponseProviderDateTranscodingLive {
  s.Definition_HD1080 = &v
  return s
}

func (s *Query5minLiveTranscodingDurationResponseProviderDateTranscodingLive) SetDefinition_SD576(v string) *Query5minLiveTranscodingDurationResponseProviderDateTranscodingLive {
  s.Definition_SD576 = &v
  return s
}

func (s *Query5minLiveTranscodingDurationResponseProviderDateTranscodingLive) SetTotal(v string) *Query5minLiveTranscodingDurationResponseProviderDateTranscodingLive {
  s.Total = &v
  return s
}

type Query5minLiveTranscodingDurationPaths struct {
}

func (s Query5minLiveTranscodingDurationPaths) String() string {
  return tea.Prettify(s)
}

func (s Query5minLiveTranscodingDurationPaths) GoString() string {
  return s.String()
}

type Query5minLiveTranscodingDurationParameters struct {
}

func (s Query5minLiveTranscodingDurationParameters) String() string {
  return tea.Prettify(s)
}

func (s Query5minLiveTranscodingDurationParameters) GoString() string {
  return s.String()
}

type Query5minLiveTranscodingDurationRequestHeader struct {
}

func (s Query5minLiveTranscodingDurationRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s Query5minLiveTranscodingDurationRequestHeader) GoString() string {
  return s.String()
}

type Query5minLiveTranscodingDurationResponseHeader struct {
}

func (s Query5minLiveTranscodingDurationResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s Query5minLiveTranscodingDurationResponseHeader) GoString() string {
  return s.String()
}




type PicProcessStatisticsRequest struct {
  // {"en":"cust_en_name of sub-client.
  // When a merged-account wants to  view the information of the subclient,the cust_en_name is required.", "zh_CN":"合并账号下的某个客户的英文名，当合并账号要查看子客户的信息时，必须填写子客户的英文名"}
  Cust *string `json:"cust,omitempty" xml:"cust,omitempty"`
  // {"en":"Specifies the query date:
  // 1.With format yyyy-mm-dd.
  // 2.If not specified,it means today as default.", "zh_CN":"查询的日期，日期格式为yyyy-mm-dd,不选或者为空时默认为当天；"}
  Date *string `json:"date,omitempty" xml:"date,omitempty"`
  // {"en":"1.Must work with 'enddate' and they  specify the query date scope.
  // 2.With format yyyy-mm-dd.
  // 3.If there is a 'date' parameter,this parameter will be omitted.", "zh_CN":"查询的起始日期 ,日期格式为yyyy-mm-dd；此参数需与enddate参数配合,若存在date参数,则该参数无效"}
  Startdate *string `json:"startdate,omitempty" xml:"startdate,omitempty"`
  // {"en":"1.Must work with 'startdate' and they  specify the query date scope.
  // 2.With format yyyy-mm-dd.
  // 3.If there is a 'date' parameter,this parameter will be omitted.", "zh_CN":"查询的结束日期 ,日期格式为yyyy-mm-dd；此参数需与startdate参数配合,若存在date参数,则该参数无效。"}
  Enddate *string `json:"enddate,omitempty" xml:"enddate,omitempty"`
  // {"en":"domains that been queried:
  // 1.If there are multiple inputs,use  ';' as separator.
  // 2.If not specified, it means all the domains of the account .", "zh_CN":"查询的频道，多个频道值请用英文分号';'，不选或者为空时默认为所查询客户的所有频道"}
  Channel *string `json:"channel,omitempty" xml:"channel,omitempty"`
  // {"en":"1.If there are multiple inputs,use ';' as separator.For example,u can use 'region=cn;apac' to query data of cn and apac region.
  // 2.If not specified, it means all the regions.", "zh_CN":"查询的加速区域的缩写，多个区域请用英文分号';'分隔开，如查询大陆及亚太区域，参数填写为：'region=cn;apac'。不选或者为空时默认为全部区域。"}
  Region *string `json:"region,omitempty" xml:"region,omitempty"`
  // {"en":"The response format:
  // 1.optional values:xml, json.
  // 2.'xml' as default.", "zh_CN":"返回结果格式,支持格式为xml和json,默认为xml"}
  Dataformat *string `json:"dataformat,omitempty" xml:"dataformat,omitempty"`
  // {"en":"Display statistic result in merged or separate way:
  // 1.If specified 1,get the merged result.
  // 2.If specified 2,get the separate result.
  // 3.If specified 3,get both merged result and separate result.
  // 4.If not specified,means '1'.", "zh_CN":"&nbsp;结果的显示是否提供合并值。填写1时：只提供合并结果；填写2时：只提供拆分值；填写3时：既提供合并值，又提供拆分值。不选或者为空时默认为'1'。"}
  ResultType *string `json:"resultType,omitempty" xml:"resultType,omitempty"`
}

func (s PicProcessStatisticsRequest) String() string {
  return tea.Prettify(s)
}

func (s PicProcessStatisticsRequest) GoString() string {
  return s.String()
}

func (s *PicProcessStatisticsRequest) SetCust(v string) *PicProcessStatisticsRequest {
  s.Cust = &v
  return s
}

func (s *PicProcessStatisticsRequest) SetDate(v string) *PicProcessStatisticsRequest {
  s.Date = &v
  return s
}

func (s *PicProcessStatisticsRequest) SetStartdate(v string) *PicProcessStatisticsRequest {
  s.Startdate = &v
  return s
}

func (s *PicProcessStatisticsRequest) SetEnddate(v string) *PicProcessStatisticsRequest {
  s.Enddate = &v
  return s
}

func (s *PicProcessStatisticsRequest) SetChannel(v string) *PicProcessStatisticsRequest {
  s.Channel = &v
  return s
}

func (s *PicProcessStatisticsRequest) SetRegion(v string) *PicProcessStatisticsRequest {
  s.Region = &v
  return s
}

func (s *PicProcessStatisticsRequest) SetDataformat(v string) *PicProcessStatisticsRequest {
  s.Dataformat = &v
  return s
}

func (s *PicProcessStatisticsRequest) SetResultType(v string) *PicProcessStatisticsRequest {
  s.ResultType = &v
  return s
}

type PicProcessStatisticsResponse struct {
  // {'en':'provider', 'zh_CN':'结果'}
  Provider *PicProcessStatisticsResponseProvider `json:"provider,omitempty" xml:"provider,omitempty" require:"true" type:"Struct"`
}

func (s PicProcessStatisticsResponse) String() string {
  return tea.Prettify(s)
}

func (s PicProcessStatisticsResponse) GoString() string {
  return s.String()
}

func (s *PicProcessStatisticsResponse) SetProvider(v *PicProcessStatisticsResponseProvider) *PicProcessStatisticsResponse {
  s.Provider = v
  return s
}

type PicProcessStatisticsResponseProvider struct {
  // {'en':'tenant', 'zh_CN':'租户'}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {'en':'type', 'zh_CN':'接口类型'}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
  // {'en':'resultType', 'zh_CN':'统计类型'}
  ResultType *string `json:"resultType,omitempty" xml:"resultType,omitempty" require:"true"`
  // {'en':'data', 'zh_CN':'明细数据'}
  Date *PicProcessStatisticsResponseProviderDate `json:"date,omitempty" xml:"date,omitempty" require:"true" type:"Struct"`
}

func (s PicProcessStatisticsResponseProvider) String() string {
  return tea.Prettify(s)
}

func (s PicProcessStatisticsResponseProvider) GoString() string {
  return s.String()
}

func (s *PicProcessStatisticsResponseProvider) SetName(v string) *PicProcessStatisticsResponseProvider {
  s.Name = &v
  return s
}

func (s *PicProcessStatisticsResponseProvider) SetType(v string) *PicProcessStatisticsResponseProvider {
  s.Type = &v
  return s
}

func (s *PicProcessStatisticsResponseProvider) SetResultType(v string) *PicProcessStatisticsResponseProvider {
  s.ResultType = &v
  return s
}

func (s *PicProcessStatisticsResponseProvider) SetDate(v *PicProcessStatisticsResponseProviderDate) *PicProcessStatisticsResponseProvider {
  s.Date = v
  return s
}

type PicProcessStatisticsResponseProviderDate struct {
  // {'en':'startdate', 'zh_CN':'开始时间'}
  Startdate *string `json:"startdate,omitempty" xml:"startdate,omitempty" require:"true"`
  // {'en':'enddate', 'zh_CN':'结束时间'}
  Enddate *string `json:"enddate,omitempty" xml:"enddate,omitempty" require:"true"`
  // {'en':'channel', 'zh_CN':'频道'}
  Channel *PicProcessStatisticsResponseProviderDateChannel `json:"channel,omitempty" xml:"channel,omitempty" require:"true" type:"Struct"`
}

func (s PicProcessStatisticsResponseProviderDate) String() string {
  return tea.Prettify(s)
}

func (s PicProcessStatisticsResponseProviderDate) GoString() string {
  return s.String()
}

func (s *PicProcessStatisticsResponseProviderDate) SetStartdate(v string) *PicProcessStatisticsResponseProviderDate {
  s.Startdate = &v
  return s
}

func (s *PicProcessStatisticsResponseProviderDate) SetEnddate(v string) *PicProcessStatisticsResponseProviderDate {
  s.Enddate = &v
  return s
}

func (s *PicProcessStatisticsResponseProviderDate) SetChannel(v *PicProcessStatisticsResponseProviderDateChannel) *PicProcessStatisticsResponseProviderDate {
  s.Channel = v
  return s
}

type PicProcessStatisticsResponseProviderDateChannel struct {
  // {'en':'channel', 'zh_CN':'频道'}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {'en':'picflowhit', 'zh_CN':'请求数数据'}
  Picflowhit []*PicProcessStatisticsResponseProviderDateChannelPicflowhit `json:"picflowhit,omitempty" xml:"picflowhit,omitempty" require:"true" type:"Repeated"`
}

func (s PicProcessStatisticsResponseProviderDateChannel) String() string {
  return tea.Prettify(s)
}

func (s PicProcessStatisticsResponseProviderDateChannel) GoString() string {
  return s.String()
}

func (s *PicProcessStatisticsResponseProviderDateChannel) SetName(v string) *PicProcessStatisticsResponseProviderDateChannel {
  s.Name = &v
  return s
}

func (s *PicProcessStatisticsResponseProviderDateChannel) SetPicflowhit(v []*PicProcessStatisticsResponseProviderDateChannelPicflowhit) *PicProcessStatisticsResponseProviderDateChannel {
  s.Picflowhit = v
  return s
}

type PicProcessStatisticsResponseProviderDateChannelPicflowhit struct     {
  // {'en':'timestamp', 'zh_CN':'时间点'}
  Time *string `json:"time,omitempty" xml:"time,omitempty" require:"true"`
  // {'en':'hit count', 'zh_CN':'请求数'}
  Text *string `json:"text,omitempty" xml:"text,omitempty" require:"true"`
}

func (s PicProcessStatisticsResponseProviderDateChannelPicflowhit) String() string {
  return tea.Prettify(s)
}

func (s PicProcessStatisticsResponseProviderDateChannelPicflowhit) GoString() string {
  return s.String()
}

func (s *PicProcessStatisticsResponseProviderDateChannelPicflowhit) SetTime(v string) *PicProcessStatisticsResponseProviderDateChannelPicflowhit {
  s.Time = &v
  return s
}

func (s *PicProcessStatisticsResponseProviderDateChannelPicflowhit) SetText(v string) *PicProcessStatisticsResponseProviderDateChannelPicflowhit {
  s.Text = &v
  return s
}

type PicProcessStatisticsPaths struct {
}

func (s PicProcessStatisticsPaths) String() string {
  return tea.Prettify(s)
}

func (s PicProcessStatisticsPaths) GoString() string {
  return s.String()
}

type PicProcessStatisticsParameters struct {
}

func (s PicProcessStatisticsParameters) String() string {
  return tea.Prettify(s)
}

func (s PicProcessStatisticsParameters) GoString() string {
  return s.String()
}

type PicProcessStatisticsRequestHeader struct {
}

func (s PicProcessStatisticsRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s PicProcessStatisticsRequestHeader) GoString() string {
  return s.String()
}

type PicProcessStatisticsResponseHeader struct {
}

func (s PicProcessStatisticsResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s PicProcessStatisticsResponseHeader) GoString() string {
  return s.String()
}




type QueryDomainsForResourceGroupRequest struct {
  // {"en":"Resource group."}
  Resource *string `json:"resource,omitempty" xml:"resource,omitempty" require:"true"`
  // {"en":"domain."}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
}

func (s QueryDomainsForResourceGroupRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryDomainsForResourceGroupRequest) GoString() string {
  return s.String()
}

func (s *QueryDomainsForResourceGroupRequest) SetResource(v string) *QueryDomainsForResourceGroupRequest {
  s.Resource = &v
  return s
}

func (s *QueryDomainsForResourceGroupRequest) SetDomain(v string) *QueryDomainsForResourceGroupRequest {
  s.Domain = &v
  return s
}

type QueryDomainsForResourceGroupResponse struct {
  // {'en':'Please refer to the error code for exceptions.'}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {'en':'Description.'}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {'en':'Data.'}
  Data []*QueryDomainsForResourceGroupDomainBean `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s QueryDomainsForResourceGroupResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryDomainsForResourceGroupResponse) GoString() string {
  return s.String()
}

func (s *QueryDomainsForResourceGroupResponse) SetCode(v string) *QueryDomainsForResourceGroupResponse {
  s.Code = &v
  return s
}

func (s *QueryDomainsForResourceGroupResponse) SetMsg(v string) *QueryDomainsForResourceGroupResponse {
  s.Msg = &v
  return s
}

func (s *QueryDomainsForResourceGroupResponse) SetData(v []*QueryDomainsForResourceGroupDomainBean) *QueryDomainsForResourceGroupResponse {
  s.Data = v
  return s
}

type QueryDomainsForResourceGroupDomainBean struct {
  // {"en":"Domain name."}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {"en":"Deploy status, 1: Deployment success 2: In deployment"}
  DeployStatus *string `json:"deployStatus,omitempty" xml:"deployStatus,omitempty" require:"true"`
}

func (s QueryDomainsForResourceGroupDomainBean) String() string {
  return tea.Prettify(s)
}

func (s QueryDomainsForResourceGroupDomainBean) GoString() string {
  return s.String()
}

func (s *QueryDomainsForResourceGroupDomainBean) SetDomain(v string) *QueryDomainsForResourceGroupDomainBean {
  s.Domain = &v
  return s
}

func (s *QueryDomainsForResourceGroupDomainBean) SetDeployStatus(v string) *QueryDomainsForResourceGroupDomainBean {
  s.DeployStatus = &v
  return s
}

type QueryDomainsForResourceGroupPaths struct {
}

func (s QueryDomainsForResourceGroupPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryDomainsForResourceGroupPaths) GoString() string {
  return s.String()
}

type QueryDomainsForResourceGroupParameters struct {
}

func (s QueryDomainsForResourceGroupParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryDomainsForResourceGroupParameters) GoString() string {
  return s.String()
}

type QueryDomainsForResourceGroupRequestHeader struct {
}

func (s QueryDomainsForResourceGroupRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryDomainsForResourceGroupRequestHeader) GoString() string {
  return s.String()
}

type QueryDomainsForResourceGroupResponseHeader struct {
}

func (s QueryDomainsForResourceGroupResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryDomainsForResourceGroupResponseHeader) GoString() string {
  return s.String()
}




type StreamTrafficServiceRequest struct {
  // {"en":"Domain(s),  multiple domains are separated by commas,up to a maximum of 5 supported","zh_CN":"域名，可多个，英文逗号分隔，最多可支持5个"}
  DomainList *string `json:"domainList,omitempty" xml:"domainList,omitempty" require:"true"`
  // {"en":"Region(s), multiple regions are separated by commas","zh_CN":"区域，可多个，英文逗号分隔"}
  RegionList *string `json:"regionList,omitempty" xml:"regionList,omitempty"`
  // {"en":"Stream(s), multiple allowed, separated by commas","zh_CN":"流名，可多个，英文逗号分隔"}
  StreamNameList *string `json:"streamNameList,omitempty" xml:"streamNameList,omitempty"`
  // {"en":"App(s), multiple, separated by commas","zh_CN":"发布点，可多个，英文逗号分隔"}
  AppList *string `json:"appList,omitempty" xml:"appList,omitempty"`
  // {"en":"Query start time. The format is yyyy-MM-ddTHH:mm:ss+08:00; for example, 2024-12-12T10:00:00+08:00 (which is 10:00 AM Beijing time on December 12, 2024). You can query data for up to 1 day.","zh_CN":"查询开始时间。格式为yyyy-MM-ddTHH:mm:ss+08:00；例如，2024-12-12T10:00:00+08:00（为北京时间2024年12月12日10点0分0秒）;最多查1天数据"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty" require:"true"`
  // {"en":"Query end time. The format is yyyy-MM-ddTHH:mm:ss+08:00","zh_CN":"查询结束时间。格式为yyyy-MM-ddTHH:mm:ss+08:00"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty" require:"true"`
  // {"en":"Value: rtmp, hdl, hls, rtc, srt, other; multiple values allowed, separated by commas.","zh_CN":"值：rtmp，hdl，hls，rtc，srt，other  可多个，英文逗号分隔"}
  ProtocolList *string `json:"protocolList,omitempty" xml:"protocolList,omitempty"`
}

func (s StreamTrafficServiceRequest) String() string {
  return tea.Prettify(s)
}

func (s StreamTrafficServiceRequest) GoString() string {
  return s.String()
}

func (s *StreamTrafficServiceRequest) SetDomainList(v string) *StreamTrafficServiceRequest {
  s.DomainList = &v
  return s
}

func (s *StreamTrafficServiceRequest) SetRegionList(v string) *StreamTrafficServiceRequest {
  s.RegionList = &v
  return s
}

func (s *StreamTrafficServiceRequest) SetStreamNameList(v string) *StreamTrafficServiceRequest {
  s.StreamNameList = &v
  return s
}

func (s *StreamTrafficServiceRequest) SetAppList(v string) *StreamTrafficServiceRequest {
  s.AppList = &v
  return s
}

func (s *StreamTrafficServiceRequest) SetDateFrom(v string) *StreamTrafficServiceRequest {
  s.DateFrom = &v
  return s
}

func (s *StreamTrafficServiceRequest) SetDateTo(v string) *StreamTrafficServiceRequest {
  s.DateTo = &v
  return s
}

func (s *StreamTrafficServiceRequest) SetProtocolList(v string) *StreamTrafficServiceRequest {
  s.ProtocolList = &v
  return s
}

type StreamTrafficServiceRequestHeader struct {
}

func (s StreamTrafficServiceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s StreamTrafficServiceRequestHeader) GoString() string {
  return s.String()
}

type StreamTrafficServicePaths struct {
}

func (s StreamTrafficServicePaths) String() string {
  return tea.Prettify(s)
}

func (s StreamTrafficServicePaths) GoString() string {
  return s.String()
}

type StreamTrafficServiceParameters struct {
}

func (s StreamTrafficServiceParameters) String() string {
  return tea.Prettify(s)
}

func (s StreamTrafficServiceParameters) GoString() string {
  return s.String()
}

type StreamTrafficServiceResponse struct {
  // {"en":"requestId","zh_CN":"请求id"}
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty" require:"true"`
  // {"en":"Query start time","zh_CN":"查询开始时间"}
  StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty" require:"true"`
  // {"en":"Query end time","zh_CN":"查询结束时间"}
  EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty" require:"true"`
  // {"en":"regions","zh_CN":"区域"}
  Regions []*StreamTrafficServiceResponseRegions `json:"regions,omitempty" xml:"regions,omitempty" require:"true" type:"Repeated"`
}

func (s StreamTrafficServiceResponse) String() string {
  return tea.Prettify(s)
}

func (s StreamTrafficServiceResponse) GoString() string {
  return s.String()
}

func (s *StreamTrafficServiceResponse) SetRequestId(v string) *StreamTrafficServiceResponse {
  s.RequestId = &v
  return s
}

func (s *StreamTrafficServiceResponse) SetStartTime(v string) *StreamTrafficServiceResponse {
  s.StartTime = &v
  return s
}

func (s *StreamTrafficServiceResponse) SetEndTime(v string) *StreamTrafficServiceResponse {
  s.EndTime = &v
  return s
}

func (s *StreamTrafficServiceResponse) SetRegions(v []*StreamTrafficServiceResponseRegions) *StreamTrafficServiceResponse {
  s.Regions = v
  return s
}

type StreamTrafficServiceResponseRegions struct     {
  // {"en":"regionCode","zh_CN":"区域编码"}
  RegionCode *string `json:"regionCode,omitempty" xml:"regionCode,omitempty" require:"true"`
  // {"en":"","zh_CN":""}
  Domains []*StreamTrafficServiceResponseRegionsDomains `json:"domains,omitempty" xml:"domains,omitempty" require:"true" type:"Repeated"`
}

func (s StreamTrafficServiceResponseRegions) String() string {
  return tea.Prettify(s)
}

func (s StreamTrafficServiceResponseRegions) GoString() string {
  return s.String()
}

func (s *StreamTrafficServiceResponseRegions) SetRegionCode(v string) *StreamTrafficServiceResponseRegions {
  s.RegionCode = &v
  return s
}

func (s *StreamTrafficServiceResponseRegions) SetDomains(v []*StreamTrafficServiceResponseRegionsDomains) *StreamTrafficServiceResponseRegions {
  s.Domains = v
  return s
}

type StreamTrafficServiceResponseRegionsDomains struct     {
  // {"en":"domain","zh_CN":"域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {"en":"","zh_CN":""}
  Streams []*StreamTrafficServiceResponseRegionsDomainsStreams `json:"streams,omitempty" xml:"streams,omitempty" require:"true" type:"Repeated"`
}

func (s StreamTrafficServiceResponseRegionsDomains) String() string {
  return tea.Prettify(s)
}

func (s StreamTrafficServiceResponseRegionsDomains) GoString() string {
  return s.String()
}

func (s *StreamTrafficServiceResponseRegionsDomains) SetDomain(v string) *StreamTrafficServiceResponseRegionsDomains {
  s.Domain = &v
  return s
}

func (s *StreamTrafficServiceResponseRegionsDomains) SetStreams(v []*StreamTrafficServiceResponseRegionsDomainsStreams) *StreamTrafficServiceResponseRegionsDomains {
  s.Streams = v
  return s
}

type StreamTrafficServiceResponseRegionsDomainsStreams struct     {
  // {"en":"streamName","zh_CN":"流名"}
  StreamName *string `json:"streamName,omitempty" xml:"streamName,omitempty" require:"true"`
  // {"en":"Application Name","zh_CN":"发布点"}
  App *string `json:"app,omitempty" xml:"app,omitempty" require:"true"`
  // {"en":"protocol","zh_CN":"协议"}
  Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty" require:"true"`
  // {"en":"flow","zh_CN":"流量"}
  Traffic *int64 `json:"traffic,omitempty" xml:"traffic,omitempty" require:"true"`
}

func (s StreamTrafficServiceResponseRegionsDomainsStreams) String() string {
  return tea.Prettify(s)
}

func (s StreamTrafficServiceResponseRegionsDomainsStreams) GoString() string {
  return s.String()
}

func (s *StreamTrafficServiceResponseRegionsDomainsStreams) SetStreamName(v string) *StreamTrafficServiceResponseRegionsDomainsStreams {
  s.StreamName = &v
  return s
}

func (s *StreamTrafficServiceResponseRegionsDomainsStreams) SetApp(v string) *StreamTrafficServiceResponseRegionsDomainsStreams {
  s.App = &v
  return s
}

func (s *StreamTrafficServiceResponseRegionsDomainsStreams) SetProtocol(v string) *StreamTrafficServiceResponseRegionsDomainsStreams {
  s.Protocol = &v
  return s
}

func (s *StreamTrafficServiceResponseRegionsDomainsStreams) SetTraffic(v int64) *StreamTrafficServiceResponseRegionsDomainsStreams {
  s.Traffic = &v
  return s
}

type StreamTrafficServiceResponseHeader struct {
}

func (s StreamTrafficServiceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s StreamTrafficServiceResponseHeader) GoString() string {
  return s.String()
}




type QueryDirectoryRankbySuccessRequestHourlyRequest struct {
  // {"en":"Start time:
  // 1. Time format is yyyy-MM-ddTHH:mm:ss+08:00,
  // 2. No bigger than the current time.
  // 3. Data in the last 183 days at most can be queried.", "zh_CN":"开始时间:
  // 
  // 1.时间格式为yyyy-MM-ddTHH:mm:ss+08:00
  // 2.不能大于当前时间
  // 3.最多可获取最近半年(183天)的数据"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty"`
  // {"en":"End time:
  // 1. The time format is yyyy-MM-ddTHH:mm:ss+08:00
  // 2. End time should be greater than start time. If the end time is greater than current time, current time will be used.
  // 3. If both fields of dataFrom and dateTo are left empty, then data in the last 1day will be queried by default; if only one field is filled in and one is left empty, then exception will be occur.
  // 4. Allowable maximum time range for query: 31days, means the period between dateFrom to dateTo should not exceed 31days (can be adjusted by contacting technical support).", "zh_CN":"结束时间:
  // 
  // 1.时间格式yyyy-MM-ddTHH:mm:ss+08:00
  // 2.结束时间需大于开始时间,结束时间如果大于当前时间,取当前时间
  // 3.dateFrom,dateTo二者都未传,默认查询过去的1天;如仅有一个未传,抛异常
  // 4.允许查询最大时间间隔31天:,即dateFrom和dateTo相差不能超过31天。(可联系技术支持调整)"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty"`
  // {"en":"Domain:
  // 1. The maximum number of domain is 100 by default (the upper limit is 500 according to technical support adjustment);
  // 2. Automatically filter out invalid domain (an illegal domain will be filtered, and the query result will only return the data of valid domains).
  // 3. If the input parameter is not inputed, all domains and multipaths under the account will be queried by default, but an error will be prompted when the number of domains under the account exceeds the upper limit", "zh_CN":"域名:
  // 1、可传递域名数量上限默认为100个(可联系技术支持调整,最高上限500);
  // 2、自动过滤掉无效域名(如传递非法域名,会被过滤,查询结果只返回有效域名的数据)。
  // 3、未传递该入参时,默认查询账号下所有域名和multipath,但当账号下域名数量超过上限时提示错误"}
  Domain []*string `json:"domain,omitempty" xml:"domain,omitempty" type:"Repeated"`
}

func (s QueryDirectoryRankbySuccessRequestHourlyRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryDirectoryRankbySuccessRequestHourlyRequest) GoString() string {
  return s.String()
}

func (s *QueryDirectoryRankbySuccessRequestHourlyRequest) SetDateFrom(v string) *QueryDirectoryRankbySuccessRequestHourlyRequest {
  s.DateFrom = &v
  return s
}

func (s *QueryDirectoryRankbySuccessRequestHourlyRequest) SetDateTo(v string) *QueryDirectoryRankbySuccessRequestHourlyRequest {
  s.DateTo = &v
  return s
}

func (s *QueryDirectoryRankbySuccessRequestHourlyRequest) SetDomain(v []*string) *QueryDirectoryRankbySuccessRequestHourlyRequest {
  s.Domain = v
  return s
}

type QueryDirectoryRankbySuccessRequestHourlyResponse struct {
  // {"en":"Top ranking", "zh_CN":"top排名"}
  Top *string `json:"top,omitempty" xml:"top,omitempty" require:"true"`
  // {"en":"Directory", "zh_CN":"目录"}
  Directory *string `json:"directory,omitempty" xml:"directory,omitempty" require:"true"`
  // {"en":"Total successful flow: the unit of measurement is MB, with 2 decimal places reserved", "zh_CN":"成功总流量:计量单位MB,保留2位小数"}
  SuccessTotalFlow *string `json:"successTotalFlow,omitempty" xml:"successTotalFlow,omitempty" require:"true"`
  // {"en":"Number of successful requests", "zh_CN":"成功请求数"}
  SuccessRequests *string `json:"successRequests,omitempty" xml:"successRequests,omitempty" require:"true"`
}

func (s QueryDirectoryRankbySuccessRequestHourlyResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryDirectoryRankbySuccessRequestHourlyResponse) GoString() string {
  return s.String()
}

func (s *QueryDirectoryRankbySuccessRequestHourlyResponse) SetTop(v string) *QueryDirectoryRankbySuccessRequestHourlyResponse {
  s.Top = &v
  return s
}

func (s *QueryDirectoryRankbySuccessRequestHourlyResponse) SetDirectory(v string) *QueryDirectoryRankbySuccessRequestHourlyResponse {
  s.Directory = &v
  return s
}

func (s *QueryDirectoryRankbySuccessRequestHourlyResponse) SetSuccessTotalFlow(v string) *QueryDirectoryRankbySuccessRequestHourlyResponse {
  s.SuccessTotalFlow = &v
  return s
}

func (s *QueryDirectoryRankbySuccessRequestHourlyResponse) SetSuccessRequests(v string) *QueryDirectoryRankbySuccessRequestHourlyResponse {
  s.SuccessRequests = &v
  return s
}

type QueryDirectoryRankbySuccessRequestHourlyPaths struct {
}

func (s QueryDirectoryRankbySuccessRequestHourlyPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryDirectoryRankbySuccessRequestHourlyPaths) GoString() string {
  return s.String()
}

type QueryDirectoryRankbySuccessRequestHourlyParameters struct {
}

func (s QueryDirectoryRankbySuccessRequestHourlyParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryDirectoryRankbySuccessRequestHourlyParameters) GoString() string {
  return s.String()
}

type QueryDirectoryRankbySuccessRequestHourlyRequestHeader struct {
}

func (s QueryDirectoryRankbySuccessRequestHourlyRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryDirectoryRankbySuccessRequestHourlyRequestHeader) GoString() string {
  return s.String()
}

type QueryDirectoryRankbySuccessRequestHourlyResponseHeader struct {
}

func (s QueryDirectoryRankbySuccessRequestHourlyResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryDirectoryRankbySuccessRequestHourlyResponseHeader) GoString() string {
  return s.String()
}




type AntiHijackIPListRequest struct {
  // {"en":"Status: 0: All, 1: Effective, 2: Deploying, 3: Deployment Exception, default is All","zh_CN":"状态：0：全部，1：已生效，2：部署中，3：部署异常，默认为全部"}
  Status *int `json:"status,omitempty" xml:"status,omitempty"`
  // {"en":"User access area 0: All 1: Mainland China 2: Global, default is All","zh_CN":"用户访问区域 0:全部 1:中国大陆 2:全球 默认为全部"}
  Area *int `json:"area,omitempty" xml:"area,omitempty"`
  // {"en":"Anti-hijack IP","zh_CN":"劫持缓解IP"}
  AntiHijackIp *string `json:"antiHijackIp,omitempty" xml:"antiHijackIp,omitempty"`
  // {"en":"Related business domain name, note that using this condition defaults useStatus to USED, only accessed mitigation IPs will have this value","zh_CN":"关联业务域名，注意，使用此条件默认useStatus为USED，只有已接入的缓解IP才会有此值"}
  RelateDomain *string `json:"relateDomain,omitempty" xml:"relateDomain,omitempty"`
  // {"en":"Usage status, USED: Used, NO_USED: Not used","zh_CN":"使用状态,USED:已使用，NO_USED: 未使用","exampleValue":"USED,NO_USE"}
  UseStatus *string `json:"useStatus,omitempty" xml:"useStatus,omitempty"`
}

func (s AntiHijackIPListRequest) String() string {
  return tea.Prettify(s)
}

func (s AntiHijackIPListRequest) GoString() string {
  return s.String()
}

func (s *AntiHijackIPListRequest) SetStatus(v int) *AntiHijackIPListRequest {
  s.Status = &v
  return s
}

func (s *AntiHijackIPListRequest) SetArea(v int) *AntiHijackIPListRequest {
  s.Area = &v
  return s
}

func (s *AntiHijackIPListRequest) SetAntiHijackIp(v string) *AntiHijackIPListRequest {
  s.AntiHijackIp = &v
  return s
}

func (s *AntiHijackIPListRequest) SetRelateDomain(v string) *AntiHijackIPListRequest {
  s.RelateDomain = &v
  return s
}

func (s *AntiHijackIPListRequest) SetUseStatus(v string) *AntiHijackIPListRequest {
  s.UseStatus = &v
  return s
}

type AntiHijackIPListRequestHeader struct {
}

func (s AntiHijackIPListRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s AntiHijackIPListRequestHeader) GoString() string {
  return s.String()
}

type AntiHijackIPListPaths struct {
}

func (s AntiHijackIPListPaths) String() string {
  return tea.Prettify(s)
}

func (s AntiHijackIPListPaths) GoString() string {
  return s.String()
}

type AntiHijackIPListParameters struct {
}

func (s AntiHijackIPListParameters) String() string {
  return tea.Prettify(s)
}

func (s AntiHijackIPListParameters) GoString() string {
  return s.String()
}

type AntiHijackIPListResponse struct {
  // {"en":"200: Success, others: Failure","zh_CN":"200:成功，其他失败"}
  Code *int `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Return description message","zh_CN":"返回描述信息"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"Data.","zh_CN":""}
  Data []*AntiHijackIPListResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s AntiHijackIPListResponse) String() string {
  return tea.Prettify(s)
}

func (s AntiHijackIPListResponse) GoString() string {
  return s.String()
}

func (s *AntiHijackIPListResponse) SetCode(v int) *AntiHijackIPListResponse {
  s.Code = &v
  return s
}

func (s *AntiHijackIPListResponse) SetMsg(v string) *AntiHijackIPListResponse {
  s.Msg = &v
  return s
}

func (s *AntiHijackIPListResponse) SetData(v []*AntiHijackIPListResponseData) *AntiHijackIPListResponse {
  s.Data = v
  return s
}

type AntiHijackIPListResponseData struct     {
  // {"en":"ID","zh_CN":"id"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {"en":"Anti-hijack IP","zh_CN":"防劫持ip"}
  AntiHijackIp *string `json:"antiHijackIp,omitempty" xml:"antiHijackIp,omitempty" require:"true"`
  // {"en":"Area, 1: Mainland, 2: Global","zh_CN":"区域，1：大陆，2：全球"}
  Area *int `json:"area,omitempty" xml:"area,omitempty" require:"true"`
  // {"en":"Related business domain","zh_CN":"关联业务域名"}
  RelateDomain *string `json:"relateDomain,omitempty" xml:"relateDomain,omitempty" require:"true"`
  // {"en":"Status","zh_CN":"状态"}
  Status *int `json:"status,omitempty" xml:"status,omitempty" require:"true"`
  // {"en":"Usage status","zh_CN":"使用状态","exampleValue":"USED,NO_USE"}
  UseStatus *string `json:"useStatus,omitempty" xml:"useStatus,omitempty" require:"true"`
  // {"en":"remark","zh_CN":"备注"}
  Remark *string `json:"remark,omitempty" xml:"remark,omitempty" require:"true"`
}

func (s AntiHijackIPListResponseData) String() string {
  return tea.Prettify(s)
}

func (s AntiHijackIPListResponseData) GoString() string {
  return s.String()
}

func (s *AntiHijackIPListResponseData) SetId(v string) *AntiHijackIPListResponseData {
  s.Id = &v
  return s
}

func (s *AntiHijackIPListResponseData) SetAntiHijackIp(v string) *AntiHijackIPListResponseData {
  s.AntiHijackIp = &v
  return s
}

func (s *AntiHijackIPListResponseData) SetArea(v int) *AntiHijackIPListResponseData {
  s.Area = &v
  return s
}

func (s *AntiHijackIPListResponseData) SetRelateDomain(v string) *AntiHijackIPListResponseData {
  s.RelateDomain = &v
  return s
}

func (s *AntiHijackIPListResponseData) SetStatus(v int) *AntiHijackIPListResponseData {
  s.Status = &v
  return s
}

func (s *AntiHijackIPListResponseData) SetUseStatus(v string) *AntiHijackIPListResponseData {
  s.UseStatus = &v
  return s
}

func (s *AntiHijackIPListResponseData) SetRemark(v string) *AntiHijackIPListResponseData {
  s.Remark = &v
  return s
}

type AntiHijackIPListResponseHeader struct {
}

func (s AntiHijackIPListResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s AntiHijackIPListResponseHeader) GoString() string {
  return s.String()
}




type QueryDomainResourceGroupRequest struct {
  // {"en":"Domain list, up to 100 domains","zh_CN":"域名列表，最多100个域名"}
  DomainList []*string `json:"domainList,omitempty" xml:"domainList,omitempty" require:"true" type:"Repeated"`
}

func (s QueryDomainResourceGroupRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryDomainResourceGroupRequest) GoString() string {
  return s.String()
}

func (s *QueryDomainResourceGroupRequest) SetDomainList(v []*string) *QueryDomainResourceGroupRequest {
  s.DomainList = v
  return s
}

type QueryDomainResourceGroupRequestHeader struct {
}

func (s QueryDomainResourceGroupRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryDomainResourceGroupRequestHeader) GoString() string {
  return s.String()
}

type QueryDomainResourceGroupPaths struct {
}

func (s QueryDomainResourceGroupPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryDomainResourceGroupPaths) GoString() string {
  return s.String()
}

type QueryDomainResourceGroupParameters struct {
}

func (s QueryDomainResourceGroupParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryDomainResourceGroupParameters) GoString() string {
  return s.String()
}

type QueryDomainResourceGroupResponse struct {
  // {"en":"Data","zh_CN":"数据"}
  Data *QueryDomainResourceGroupResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
  // {"en":"Status code","zh_CN":"状态码"}
  Code *int `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Message","zh_CN":"信息"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
}

func (s QueryDomainResourceGroupResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryDomainResourceGroupResponse) GoString() string {
  return s.String()
}

func (s *QueryDomainResourceGroupResponse) SetData(v *QueryDomainResourceGroupResponseData) *QueryDomainResourceGroupResponse {
  s.Data = v
  return s
}

func (s *QueryDomainResourceGroupResponse) SetCode(v int) *QueryDomainResourceGroupResponse {
  s.Code = &v
  return s
}

func (s *QueryDomainResourceGroupResponse) SetMsg(v string) *QueryDomainResourceGroupResponse {
  s.Msg = &v
  return s
}

type QueryDomainResourceGroupResponseData struct {
  // {"en":"List of domain associated resource group information","zh_CN":"域名关联资源组信息列表"}
  InfoList []*QueryDomainResourceGroupResponseDataInfoList `json:"infoList,omitempty" xml:"infoList,omitempty" require:"true" type:"Repeated"`
}

func (s QueryDomainResourceGroupResponseData) String() string {
  return tea.Prettify(s)
}

func (s QueryDomainResourceGroupResponseData) GoString() string {
  return s.String()
}

func (s *QueryDomainResourceGroupResponseData) SetInfoList(v []*QueryDomainResourceGroupResponseDataInfoList) *QueryDomainResourceGroupResponseData {
  s.InfoList = v
  return s
}

type QueryDomainResourceGroupResponseDataInfoList struct     {
  // {"en":"Domain","zh_CN":"域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {"en":"Resource group name","zh_CN":"资源组名称"}
  Resource *string `json:"resource,omitempty" xml:"resource,omitempty" require:"true"`
  // {"en":"Resource group ID","zh_CN":"资源组Id"}
  ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty" require:"true"`
}

func (s QueryDomainResourceGroupResponseDataInfoList) String() string {
  return tea.Prettify(s)
}

func (s QueryDomainResourceGroupResponseDataInfoList) GoString() string {
  return s.String()
}

func (s *QueryDomainResourceGroupResponseDataInfoList) SetDomain(v string) *QueryDomainResourceGroupResponseDataInfoList {
  s.Domain = &v
  return s
}

func (s *QueryDomainResourceGroupResponseDataInfoList) SetResource(v string) *QueryDomainResourceGroupResponseDataInfoList {
  s.Resource = &v
  return s
}

func (s *QueryDomainResourceGroupResponseDataInfoList) SetResourceGroupId(v string) *QueryDomainResourceGroupResponseDataInfoList {
  s.ResourceGroupId = &v
  return s
}

type QueryDomainResourceGroupResponseHeader struct {
}

func (s QueryDomainResourceGroupResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryDomainResourceGroupResponseHeader) GoString() string {
  return s.String()
}




type QueryLiveRecordingDurationRequest struct {
  // {"en":"cust_en_name of sub-client.
  // When a merged-account wants to  view the information of the subclient,the cust_en_name is required.", "zh_CN":"合并账号下的某个客户的英文名，当合并账号要查看子客户的信息时，必须填写子客户的英文名"}
  Cust *string `json:"cust,omitempty" xml:"cust,omitempty"`
  // {"en":"Specifies the query date:
  // 1.With format yyyy-mm-dd.
  // 2.If not specified,it means today as default.", "zh_CN":"查询的日期，日期格式为yyyy-mm-dd,不选或者为空时默认为当天；"}
  Date *string `json:"date,omitempty" xml:"date,omitempty"`
  // {"en":"1.Must work with 'enddate' and they  specify the query date scope. 
  // 2.With format yyyy-mm-dd.
  // 3.If there is a 'date' parameter,this parameter will be omitted.", "zh_CN":"查询的起始日期 ,日期格式为yyyy-mm-dd；此参数需与enddate参数配合,若存在date参数,则该参数无效"}
  Startdate *string `json:"startdate,omitempty" xml:"startdate,omitempty"`
  // {"en":"1.Must work with 'startdate' and they  specify the query date scope. 
  // 2.With format yyyy-mm-dd
  // 3.If there is a 'date' parameter,this parameter will be omitted.", "zh_CN":"查询的结束日期 ,日期格式为yyyy-mm-dd；此参数需与startdate参数配合,若存在date参数,则该参数无效。"}
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
  // {"en":"Display statistic result in merged or separate way
  // 1.If specified 1, get the merged result.
  // 2.If  specified 2,get the separate result.
  // 3.If specifed 3,get both merged result and separate result.
  // 4.If not specified,means '1'.", "zh_CN":"&nbsp;结果的显示是否提供合并值。填写1时：只提供合并结果；填写2时：只提供拆分值；填写3时：既提供合并值，又提供拆分值。不选或者为空时默认为'1'。"}
  ResultType *string `json:"resultType,omitempty" xml:"resultType,omitempty"`
}

func (s QueryLiveRecordingDurationRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryLiveRecordingDurationRequest) GoString() string {
  return s.String()
}

func (s *QueryLiveRecordingDurationRequest) SetCust(v string) *QueryLiveRecordingDurationRequest {
  s.Cust = &v
  return s
}

func (s *QueryLiveRecordingDurationRequest) SetDate(v string) *QueryLiveRecordingDurationRequest {
  s.Date = &v
  return s
}

func (s *QueryLiveRecordingDurationRequest) SetStartdate(v string) *QueryLiveRecordingDurationRequest {
  s.Startdate = &v
  return s
}

func (s *QueryLiveRecordingDurationRequest) SetEnddate(v string) *QueryLiveRecordingDurationRequest {
  s.Enddate = &v
  return s
}

func (s *QueryLiveRecordingDurationRequest) SetTimezone(v string) *QueryLiveRecordingDurationRequest {
  s.Timezone = &v
  return s
}

func (s *QueryLiveRecordingDurationRequest) SetChannel(v string) *QueryLiveRecordingDurationRequest {
  s.Channel = &v
  return s
}

func (s *QueryLiveRecordingDurationRequest) SetRegion(v string) *QueryLiveRecordingDurationRequest {
  s.Region = &v
  return s
}

func (s *QueryLiveRecordingDurationRequest) SetAccetype(v string) *QueryLiveRecordingDurationRequest {
  s.Accetype = &v
  return s
}

func (s *QueryLiveRecordingDurationRequest) SetDataformat(v string) *QueryLiveRecordingDurationRequest {
  s.Dataformat = &v
  return s
}

func (s *QueryLiveRecordingDurationRequest) SetResultType(v string) *QueryLiveRecordingDurationRequest {
  s.ResultType = &v
  return s
}

type QueryLiveRecordingDurationResponse struct {
  // {'en':'provider', 'zh_CN':'结果'}
  Provider *QueryLiveRecordingDurationResponseProvider `json:"provider,omitempty" xml:"provider,omitempty" require:"true" type:"Struct"`
}

func (s QueryLiveRecordingDurationResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryLiveRecordingDurationResponse) GoString() string {
  return s.String()
}

func (s *QueryLiveRecordingDurationResponse) SetProvider(v *QueryLiveRecordingDurationResponseProvider) *QueryLiveRecordingDurationResponse {
  s.Provider = v
  return s
}

type QueryLiveRecordingDurationResponseProvider struct {
  // {'en':'tenant', 'zh_CN':'租户'}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {'en':'type', 'zh_CN':'接口类型'}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
  // {'en':'resultType', 'zh_CN':'统计类型'}
  ResultType *string `json:"resultType,omitempty" xml:"resultType,omitempty" require:"true"`
  // {'en':'data', 'zh_CN':'请明细数据'}
  Date *QueryLiveRecordingDurationResponseProviderDate `json:"date,omitempty" xml:"date,omitempty" require:"true" type:"Struct"`
}

func (s QueryLiveRecordingDurationResponseProvider) String() string {
  return tea.Prettify(s)
}

func (s QueryLiveRecordingDurationResponseProvider) GoString() string {
  return s.String()
}

func (s *QueryLiveRecordingDurationResponseProvider) SetName(v string) *QueryLiveRecordingDurationResponseProvider {
  s.Name = &v
  return s
}

func (s *QueryLiveRecordingDurationResponseProvider) SetType(v string) *QueryLiveRecordingDurationResponseProvider {
  s.Type = &v
  return s
}

func (s *QueryLiveRecordingDurationResponseProvider) SetResultType(v string) *QueryLiveRecordingDurationResponseProvider {
  s.ResultType = &v
  return s
}

func (s *QueryLiveRecordingDurationResponseProvider) SetDate(v *QueryLiveRecordingDurationResponseProviderDate) *QueryLiveRecordingDurationResponseProvider {
  s.Date = v
  return s
}

type QueryLiveRecordingDurationResponseProviderDate struct {
  // {'en':'start', 'zh_CN':'开始时间'}
  Start *string `json:"start,omitempty" xml:"start,omitempty" require:"true"`
  // {'en':'end', 'zh_CN':'结束时间'}
  End *string `json:"end,omitempty" xml:"end,omitempty" require:"true"`
  // {'en':'channel', 'zh_CN':'频道'}
  RecordingTime *QueryLiveRecordingDurationResponseProviderDateRecordingTime `json:"recordingTime,omitempty" xml:"recordingTime,omitempty" require:"true" type:"Struct"`
}

func (s QueryLiveRecordingDurationResponseProviderDate) String() string {
  return tea.Prettify(s)
}

func (s QueryLiveRecordingDurationResponseProviderDate) GoString() string {
  return s.String()
}

func (s *QueryLiveRecordingDurationResponseProviderDate) SetStart(v string) *QueryLiveRecordingDurationResponseProviderDate {
  s.Start = &v
  return s
}

func (s *QueryLiveRecordingDurationResponseProviderDate) SetEnd(v string) *QueryLiveRecordingDurationResponseProviderDate {
  s.End = &v
  return s
}

func (s *QueryLiveRecordingDurationResponseProviderDate) SetRecordingTime(v *QueryLiveRecordingDurationResponseProviderDateRecordingTime) *QueryLiveRecordingDurationResponseProviderDate {
  s.RecordingTime = v
  return s
}

type QueryLiveRecordingDurationResponseProviderDateRecordingTime struct {
  // {'en':'channel', 'zh_CN':'频道'}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  Channel *QueryLiveRecordingDurationResponseProviderDateRecordingTimeChannel `json:"channel,omitempty" xml:"channel,omitempty" require:"true" type:"Struct"`
}

func (s QueryLiveRecordingDurationResponseProviderDateRecordingTime) String() string {
  return tea.Prettify(s)
}

func (s QueryLiveRecordingDurationResponseProviderDateRecordingTime) GoString() string {
  return s.String()
}

func (s *QueryLiveRecordingDurationResponseProviderDateRecordingTime) SetName(v string) *QueryLiveRecordingDurationResponseProviderDateRecordingTime {
  s.Name = &v
  return s
}

func (s *QueryLiveRecordingDurationResponseProviderDateRecordingTime) SetChannel(v *QueryLiveRecordingDurationResponseProviderDateRecordingTimeChannel) *QueryLiveRecordingDurationResponseProviderDateRecordingTime {
  s.Channel = v
  return s
}

type QueryLiveRecordingDurationResponseProviderDateRecordingTimeChannel struct {
  // {'en':'channel', 'zh_CN':'频道'}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {'en':'bandwidth', 'zh_CN':'明细数据'}
  Live []*QueryLiveRecordingDurationResponseProviderDateRecordingTimeChannelLive `json:"live,omitempty" xml:"live,omitempty" require:"true" type:"Repeated"`
}

func (s QueryLiveRecordingDurationResponseProviderDateRecordingTimeChannel) String() string {
  return tea.Prettify(s)
}

func (s QueryLiveRecordingDurationResponseProviderDateRecordingTimeChannel) GoString() string {
  return s.String()
}

func (s *QueryLiveRecordingDurationResponseProviderDateRecordingTimeChannel) SetName(v string) *QueryLiveRecordingDurationResponseProviderDateRecordingTimeChannel {
  s.Name = &v
  return s
}

func (s *QueryLiveRecordingDurationResponseProviderDateRecordingTimeChannel) SetLive(v []*QueryLiveRecordingDurationResponseProviderDateRecordingTimeChannelLive) *QueryLiveRecordingDurationResponseProviderDateRecordingTimeChannel {
  s.Live = v
  return s
}

type QueryLiveRecordingDurationResponseProviderDateRecordingTimeChannelLive struct     {
  // {'en':'timestamp', 'zh_CN':'时间点'}
  Time *string `json:"time,omitempty" xml:"time,omitempty" require:"true"`
  // {'en':'hit count', 'zh_CN':'明细数据'}
  Text *string `json:"text,omitempty" xml:"text,omitempty" require:"true"`
}

func (s QueryLiveRecordingDurationResponseProviderDateRecordingTimeChannelLive) String() string {
  return tea.Prettify(s)
}

func (s QueryLiveRecordingDurationResponseProviderDateRecordingTimeChannelLive) GoString() string {
  return s.String()
}

func (s *QueryLiveRecordingDurationResponseProviderDateRecordingTimeChannelLive) SetTime(v string) *QueryLiveRecordingDurationResponseProviderDateRecordingTimeChannelLive {
  s.Time = &v
  return s
}

func (s *QueryLiveRecordingDurationResponseProviderDateRecordingTimeChannelLive) SetText(v string) *QueryLiveRecordingDurationResponseProviderDateRecordingTimeChannelLive {
  s.Text = &v
  return s
}

type QueryLiveRecordingDurationPaths struct {
}

func (s QueryLiveRecordingDurationPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryLiveRecordingDurationPaths) GoString() string {
  return s.String()
}

type QueryLiveRecordingDurationParameters struct {
}

func (s QueryLiveRecordingDurationParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryLiveRecordingDurationParameters) GoString() string {
  return s.String()
}

type QueryLiveRecordingDurationRequestHeader struct {
}

func (s QueryLiveRecordingDurationRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryLiveRecordingDurationRequestHeader) GoString() string {
  return s.String()
}

type QueryLiveRecordingDurationResponseHeader struct {
}

func (s QueryLiveRecordingDurationResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryLiveRecordingDurationResponseHeader) GoString() string {
  return s.String()
}




type HttpDnsStatisticsRequest struct {
  // {"en":"cust_en_name of sub-client.
  // When a merged-account wants to  view the information of the subclient,the cust_en_name is required.", "zh_CN":"合并账号下的某个客户的英文名，当合并账号要查看子客户的信息时，必须填写子客户的英文名"}
  Cust *string `json:"cust,omitempty" xml:"cust,omitempty"`
  // {"en":"Querying date, the date format is yyyy-mm-dd, and the current date will be adopted by default if the field is not selected or left empty;", "zh_CN":"查询的日期，日期格式为yyyy-mm-dd,不选或者为空时默认为当天；"}
  Date *string `json:"date,omitempty" xml:"date,omitempty"`
  // {"en":"Querying start date, and the date format is yyyy-mm-dd.
  // The parameter needs to work in concert with enddate, and it will be invalid if there exists date parameter", "zh_CN":"查询的起始日期,日期格式为yyyy-mm-dd；此参数需与enddate参数配合,若存在date参数,则该参数无效"}
  Startdate *string `json:"startdate,omitempty" xml:"startdate,omitempty"`
  // {"en":"Querying end date, and the date format is yyyy-mm-dd.
  // The parameter should work in concert with the parameter startdate. It will be invalid if there exists date parameter", "zh_CN":"查询的结束日期,日期格式为yyyy-mm-dd；此参数需与startdate参数配合,若存在date参数,则该参数无效。"}
  Enddate *string `json:"enddate,omitempty" xml:"enddate,omitempty"`
  // {"en":"Querying domain, and use English semicolon ;  as separator if there are multiple domains; all domains of the user will be queried if the field is not selected or left empty.", "zh_CN":"查询的频道，多个频道值请用英文分号;，不选或者为空时默认为所查询客户的所有频道"}
  Channel *string `json:"channel,omitempty" xml:"channel,omitempty"`
  // {"en":"acceleration type.
  // 1.If there are multiple inputs,use ';' as separator.
  // 2.If not specified or specified as 'all', it means all the accetypes.", "zh_CN":"查询域名所属的加速类型，如accetype=web。多个请用英文分号“;”分隔开，不填或值为all表示所有类型"}
  Accetype *string `json:"accetype,omitempty" xml:"accetype,omitempty"`
  // {"en":"Response result format, and the supporting format is xml and json. Default as xml.", "zh_CN":"返回结果格式,支持格式为xml和json,默认为xml"}
  Dataformat *string `json:"dataformat,omitempty" xml:"dataformat,omitempty"`
  // {"en":"Whether the channel matches exactly: when true, the full domain name must be provided (at this point, any invalid or duplicate channels entered by the user will be filtered out, and a 403 error will be returned if all input channels are invalid). When not true, all channels ending with the user-entered channel are displayed. The default is true.", "zh_CN":" 频道是否完全匹配,为true时，必须填写完整的域名(此时会过滤用户输入的无效或重复频道,所有输入频道都无效时返403)。不为true时，显示以用户输入的频道为结尾的所有频道。默认为true"}
  IsExactMatch *string `json:"isExactMatch,omitempty" xml:"isExactMatch,omitempty"`
  // {"en":"Whether the display of results includes merged values: Enter 1 to provide merged results; enter 2 to provide only split values; if not selected or left empty, the default is 1.", "zh_CN":"结果的显示是否提供合并值。填写1时：提供合并结果；填写2时：只提供拆分值；不选或者为空时默认为1。"}
  ResultType *string `json:"resultType,omitempty" xml:"resultType,omitempty"`
}

func (s HttpDnsStatisticsRequest) String() string {
  return tea.Prettify(s)
}

func (s HttpDnsStatisticsRequest) GoString() string {
  return s.String()
}

func (s *HttpDnsStatisticsRequest) SetCust(v string) *HttpDnsStatisticsRequest {
  s.Cust = &v
  return s
}

func (s *HttpDnsStatisticsRequest) SetDate(v string) *HttpDnsStatisticsRequest {
  s.Date = &v
  return s
}

func (s *HttpDnsStatisticsRequest) SetStartdate(v string) *HttpDnsStatisticsRequest {
  s.Startdate = &v
  return s
}

func (s *HttpDnsStatisticsRequest) SetEnddate(v string) *HttpDnsStatisticsRequest {
  s.Enddate = &v
  return s
}

func (s *HttpDnsStatisticsRequest) SetChannel(v string) *HttpDnsStatisticsRequest {
  s.Channel = &v
  return s
}

func (s *HttpDnsStatisticsRequest) SetAccetype(v string) *HttpDnsStatisticsRequest {
  s.Accetype = &v
  return s
}

func (s *HttpDnsStatisticsRequest) SetDataformat(v string) *HttpDnsStatisticsRequest {
  s.Dataformat = &v
  return s
}

func (s *HttpDnsStatisticsRequest) SetIsExactMatch(v string) *HttpDnsStatisticsRequest {
  s.IsExactMatch = &v
  return s
}

func (s *HttpDnsStatisticsRequest) SetResultType(v string) *HttpDnsStatisticsRequest {
  s.ResultType = &v
  return s
}

type HttpDnsStatisticsResponse struct {
  // {'en':'provider', 'zh_CN':'结果'}
  Provider *HttpDnsStatisticsResponseProvider `json:"provider,omitempty" xml:"provider,omitempty" require:"true" type:"Struct"`
}

func (s HttpDnsStatisticsResponse) String() string {
  return tea.Prettify(s)
}

func (s HttpDnsStatisticsResponse) GoString() string {
  return s.String()
}

func (s *HttpDnsStatisticsResponse) SetProvider(v *HttpDnsStatisticsResponseProvider) *HttpDnsStatisticsResponse {
  s.Provider = v
  return s
}

type HttpDnsStatisticsResponseProvider struct {
  // {'en':'tenant', 'zh_CN':'租户'}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {'en':'type', 'zh_CN':'接口类型'}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
  // {'en':'resultType', 'zh_CN':'统计类型'}
  ResultType *string `json:"resultType,omitempty" xml:"resultType,omitempty" require:"true"`
  // {'en':'data', 'zh_CN':'解析量数据'}
  Date *HttpDnsStatisticsResponseProviderDate `json:"date,omitempty" xml:"date,omitempty" require:"true" type:"Struct"`
}

func (s HttpDnsStatisticsResponseProvider) String() string {
  return tea.Prettify(s)
}

func (s HttpDnsStatisticsResponseProvider) GoString() string {
  return s.String()
}

func (s *HttpDnsStatisticsResponseProvider) SetName(v string) *HttpDnsStatisticsResponseProvider {
  s.Name = &v
  return s
}

func (s *HttpDnsStatisticsResponseProvider) SetType(v string) *HttpDnsStatisticsResponseProvider {
  s.Type = &v
  return s
}

func (s *HttpDnsStatisticsResponseProvider) SetResultType(v string) *HttpDnsStatisticsResponseProvider {
  s.ResultType = &v
  return s
}

func (s *HttpDnsStatisticsResponseProvider) SetDate(v *HttpDnsStatisticsResponseProviderDate) *HttpDnsStatisticsResponseProvider {
  s.Date = v
  return s
}

type HttpDnsStatisticsResponseProviderDate struct {
  // {'en':'startdate', 'zh_CN':'开始时间'}
  Startdate *string `json:"startdate,omitempty" xml:"startdate,omitempty" require:"true"`
  // {'en':'enddate', 'zh_CN':'结束时间'}
  Enddate *string `json:"enddate,omitempty" xml:"enddate,omitempty" require:"true"`
  // {'en':'channel', 'zh_CN':'频道'}
  Channel *HttpDnsStatisticsResponseProviderDateChannel `json:"channel,omitempty" xml:"channel,omitempty" require:"true" type:"Struct"`
}

func (s HttpDnsStatisticsResponseProviderDate) String() string {
  return tea.Prettify(s)
}

func (s HttpDnsStatisticsResponseProviderDate) GoString() string {
  return s.String()
}

func (s *HttpDnsStatisticsResponseProviderDate) SetStartdate(v string) *HttpDnsStatisticsResponseProviderDate {
  s.Startdate = &v
  return s
}

func (s *HttpDnsStatisticsResponseProviderDate) SetEnddate(v string) *HttpDnsStatisticsResponseProviderDate {
  s.Enddate = &v
  return s
}

func (s *HttpDnsStatisticsResponseProviderDate) SetChannel(v *HttpDnsStatisticsResponseProviderDateChannel) *HttpDnsStatisticsResponseProviderDate {
  s.Channel = v
  return s
}

type HttpDnsStatisticsResponseProviderDateChannel struct {
  // {'en':'channel', 'zh_CN':'频道'}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {'en':'total', 'zh_CN':'总解析量'}
  Total *string `json:"total,omitempty" xml:"total,omitempty" require:"true"`
  // {'en':'httpdns', 'zh_CN':'解析量数据'}
  Httpdns []*HttpDnsStatisticsResponseProviderDateChannelHttpdns `json:"httpdns,omitempty" xml:"httpdns,omitempty" require:"true" type:"Repeated"`
}

func (s HttpDnsStatisticsResponseProviderDateChannel) String() string {
  return tea.Prettify(s)
}

func (s HttpDnsStatisticsResponseProviderDateChannel) GoString() string {
  return s.String()
}

func (s *HttpDnsStatisticsResponseProviderDateChannel) SetName(v string) *HttpDnsStatisticsResponseProviderDateChannel {
  s.Name = &v
  return s
}

func (s *HttpDnsStatisticsResponseProviderDateChannel) SetTotal(v string) *HttpDnsStatisticsResponseProviderDateChannel {
  s.Total = &v
  return s
}

func (s *HttpDnsStatisticsResponseProviderDateChannel) SetHttpdns(v []*HttpDnsStatisticsResponseProviderDateChannelHttpdns) *HttpDnsStatisticsResponseProviderDateChannel {
  s.Httpdns = v
  return s
}

type HttpDnsStatisticsResponseProviderDateChannelHttpdns struct     {
  // {'en':'timestamp', 'zh_CN':'时间点'}
  Time *string `json:"time,omitempty" xml:"time,omitempty" require:"true"`
  // {'en':'bandwidth', 'zh_CN':'解析量'}
  Text *string `json:"text,omitempty" xml:"text,omitempty" require:"true"`
}

func (s HttpDnsStatisticsResponseProviderDateChannelHttpdns) String() string {
  return tea.Prettify(s)
}

func (s HttpDnsStatisticsResponseProviderDateChannelHttpdns) GoString() string {
  return s.String()
}

func (s *HttpDnsStatisticsResponseProviderDateChannelHttpdns) SetTime(v string) *HttpDnsStatisticsResponseProviderDateChannelHttpdns {
  s.Time = &v
  return s
}

func (s *HttpDnsStatisticsResponseProviderDateChannelHttpdns) SetText(v string) *HttpDnsStatisticsResponseProviderDateChannelHttpdns {
  s.Text = &v
  return s
}

type HttpDnsStatisticsPaths struct {
}

func (s HttpDnsStatisticsPaths) String() string {
  return tea.Prettify(s)
}

func (s HttpDnsStatisticsPaths) GoString() string {
  return s.String()
}

type HttpDnsStatisticsParameters struct {
}

func (s HttpDnsStatisticsParameters) String() string {
  return tea.Prettify(s)
}

func (s HttpDnsStatisticsParameters) GoString() string {
  return s.String()
}

type HttpDnsStatisticsRequestHeader struct {
}

func (s HttpDnsStatisticsRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s HttpDnsStatisticsRequestHeader) GoString() string {
  return s.String()
}

type HttpDnsStatisticsResponseHeader struct {
}

func (s HttpDnsStatisticsResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s HttpDnsStatisticsResponseHeader) GoString() string {
  return s.String()
}




type QueryLiveStreamStatusRequest struct {
}

func (s QueryLiveStreamStatusRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryLiveStreamStatusRequest) GoString() string {
  return s.String()
}

type QueryLiveStreamStatusRequestHeader struct {
}

func (s QueryLiveStreamStatusRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryLiveStreamStatusRequestHeader) GoString() string {
  return s.String()
}

type QueryLiveStreamStatusPaths struct {
}

func (s QueryLiveStreamStatusPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryLiveStreamStatusPaths) GoString() string {
  return s.String()
}

type QueryLiveStreamStatusParameters struct {
  // {"en":"Domain (multiple domains supported, separated by commas)(    All parameters are passed via HTTP GET requests.)","zh_CN":"域名（支持多个域名，以逗号分隔）(所有参数以HTTP GET请求方式传参)"}
  U *string `json:"u,omitempty" xml:"u,omitempty" require:"true"`
  // {"en":"20160527152300, non-real-time indicates the current time -5 minutes, real-time indicates the current time -30 seconds\nA:\n(t is not transmitted, the current system time is obtained and rounded to second (for example, 2017/3/28 14:38:55 rounded to second 2017/3/28 14:38:50 rounded to second);) If the query interval g=10, the final time (dateFrom) is rounded in seconds (e.g. 2017/3/28 14:38:55 after rounded in seconds 2017/3/28 14:38:50); If the g! =10, rounded minutes: (e.g. 2017/3/28 14:38:55 rounded seconds 2017/3/28 14:38:00); DateTo = dataFrom + 9 (9 seconds)","zh_CN":"20160527152300，不填为当前时间-5分钟\n详解：\n（t不传，获取当前系统时间，对秒取整(如：2017/3/28 14:38:55 对秒取整后2017/3/28 14:38:50)；）如果查询间隔g=10，最后获得的时间对秒取整(如：2017/3/28 14:38:55 对秒取整后2017/3/28 14:38:50)；如果g!=10，对分钟取整：(如：2017/3/28 14:38:55 对秒取整后2017/3/28 14:38:00);"}
  T *string `json:"t,omitempty" xml:"t,omitempty"`
  // {"en":"Channel URL (simple channel url are supported,eg:push1.test.com/test/test1)","zh_CN":"流名(支持单流名，如：push1.test.com/test/test1)"}
  Channel *string `json:"channel,omitempty" xml:"channel,omitempty"`
  // {"en":"Optional value: push or pull. The default value is push\nPush stands for watershed name\nPull indicates the name of a pull basin","zh_CN":"域名类型，可选值：push、pull，默认push\npush代表推流域名\npull代表拉流域名"}
  D *string `json:"d,omitempty" xml:"d,omitempty"`
  // {"en":"The value can be true or false. The default value is false\nOnly realtime data is returned when realtime=true","zh_CN":"是否返回端口，可选值：true、false，默认false"}
  Showport *string `json:"showport,omitempty" xml:"showport,omitempty"`
  // {"en":"Query interval. The value can be 10 or 60. The default value is 60\nWhen g is 10, query the data of the nearest whole 10 seconds to time t\nWhen g is 60, query the data of the nearest whole minute to time t","zh_CN":"查询间隔，可选值10、60，默认60\n当g为10时，查询距离时间t最近的整10秒点数据\n当g为60时，查询距离时间t最近的整分钟点数据"}
  G *string `json:"g,omitempty" xml:"g,omitempty"`
  // {"en":"it is query realtime datas,default false","zh_CN":"是否返回实时数据，默认false"}
  Realtime *string `json:"realtime,omitempty" xml:"realtime,omitempty"`
  // {"en":"Data extension fields(Only supports non-real-time query), providing optional return fields:gop","zh_CN":"数据扩展字段(仅支持按非实时查)，提供可选返回字段：gop"}
  Expand *string `json:"expand,omitempty" xml:"expand,omitempty"`
}

func (s QueryLiveStreamStatusParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryLiveStreamStatusParameters) GoString() string {
  return s.String()
}

func (s *QueryLiveStreamStatusParameters) SetU(v string) *QueryLiveStreamStatusParameters {
  s.U = &v
  return s
}

func (s *QueryLiveStreamStatusParameters) SetT(v string) *QueryLiveStreamStatusParameters {
  s.T = &v
  return s
}

func (s *QueryLiveStreamStatusParameters) SetChannel(v string) *QueryLiveStreamStatusParameters {
  s.Channel = &v
  return s
}

func (s *QueryLiveStreamStatusParameters) SetD(v string) *QueryLiveStreamStatusParameters {
  s.D = &v
  return s
}

func (s *QueryLiveStreamStatusParameters) SetShowport(v string) *QueryLiveStreamStatusParameters {
  s.Showport = &v
  return s
}

func (s *QueryLiveStreamStatusParameters) SetG(v string) *QueryLiveStreamStatusParameters {
  s.G = &v
  return s
}

func (s *QueryLiveStreamStatusParameters) SetRealtime(v string) *QueryLiveStreamStatusParameters {
  s.Realtime = &v
  return s
}

func (s *QueryLiveStreamStatusParameters) SetExpand(v string) *QueryLiveStreamStatusParameters {
  s.Expand = &v
  return s
}

type QueryLiveStreamStatusResponse struct {
  // {"en":"The time of the data returned","zh_CN":"返回的数据的时间"}
  Rettime *string `json:"rettime,omitempty" xml:"rettime,omitempty" require:"true"`
  // {"en":"Number of data items. 0 is returned if there is no data","zh_CN":"数据条数，无数据返回0"}
  Retcode *int64 `json:"retcode,omitempty" xml:"retcode,omitempty" require:"true"`
  // {"en":"Total onlines","zh_CN":"总在线人数"}
  Histscount *int64 `json:"histscount,omitempty" xml:"histscount,omitempty" require:"true"`
  // {"en":"Total channel bandwidth","zh_CN":"总频道带宽"}
  Bandwidthcount *int64 `json:"bandwidthcount,omitempty" xml:"bandwidthcount,omitempty" require:"true"`
  // {"en":"Timestamp","zh_CN":"数据的时间戳,如果有传t,则等于t"}
  Datetime *int64 `json:"datetime,omitempty" xml:"datetime,omitempty" require:"true"`
  // {"en":"","zh_CN":"数据集合"}
  DataValue []*QueryLiveStreamStatusResponseDataValue `json:"dataValue,omitempty" xml:"dataValue,omitempty" require:"true" type:"Repeated"`
}

func (s QueryLiveStreamStatusResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryLiveStreamStatusResponse) GoString() string {
  return s.String()
}

func (s *QueryLiveStreamStatusResponse) SetRettime(v string) *QueryLiveStreamStatusResponse {
  s.Rettime = &v
  return s
}

func (s *QueryLiveStreamStatusResponse) SetRetcode(v int64) *QueryLiveStreamStatusResponse {
  s.Retcode = &v
  return s
}

func (s *QueryLiveStreamStatusResponse) SetHistscount(v int64) *QueryLiveStreamStatusResponse {
  s.Histscount = &v
  return s
}

func (s *QueryLiveStreamStatusResponse) SetBandwidthcount(v int64) *QueryLiveStreamStatusResponse {
  s.Bandwidthcount = &v
  return s
}

func (s *QueryLiveStreamStatusResponse) SetDatetime(v int64) *QueryLiveStreamStatusResponse {
  s.Datetime = &v
  return s
}

func (s *QueryLiveStreamStatusResponse) SetDataValue(v []*QueryLiveStreamStatusResponseDataValue) *QueryLiveStreamStatusResponse {
  s.DataValue = v
  return s
}

type QueryLiveStreamStatusResponseDataValue struct     {
  // {"en":"The channel of anchor","zh_CN":"主播流名"}
  Streamname *string `json:"streamname,omitempty" xml:"streamname,omitempty" require:"true"`
  // {"en":"IP address of the CDN node","zh_CN":"推流cdn节点IP"}
  Deployaddress *string `json:"deployaddress,omitempty" xml:"deployaddress,omitempty" require:"true"`
  // {"en":"Anchor exit Address","zh_CN":"主播出口地址"}
  Inaddress *string `json:"inaddress,omitempty" xml:"inaddress,omitempty" require:"true"`
  // {"en":"The number of online","zh_CN":"在线人数"}
  Hists *int64 `json:"hists,omitempty" xml:"hists,omitempty" require:"true"`
  // {"en":"Anchor Current bit rate (transcoding stream has no bit rate data) unit: BPS","zh_CN":"主播当前码率(转码流没有码率数据) 单位：bps"}
  Inbandwidth *int64 `json:"inbandwidth,omitempty" xml:"inbandwidth,omitempty" require:"true"`
  // {"en":"Channel current viewing bandwidth unit: BPS","zh_CN":"频道当前观看带宽 单位：bps"}
  Bandwidth *int64 `json:"bandwidth,omitempty" xml:"bandwidth,omitempty" require:"true"`
  // {"en":"Anchor delay (MS)","zh_CN":"主播延迟(ms)"}
  Delay *int64 `json:"delay,omitempty" xml:"delay,omitempty" require:"true"`
  // {"en":"Anchor Current encoding frame rate","zh_CN":"主播当前编码帧率"}
  Fps *int64 `json:"fps,omitempty" xml:"fps,omitempty" require:"true"`
  // {"en":"Current frame loss rate of anchor","zh_CN":"主播当前丢帧率"}
  Lfr *QueryLiveStreamStatusResponseDataValueLfr `json:"lfr,omitempty" xml:"lfr,omitempty" require:"true" type:"Struct"`
  // {"en":"Anchor raw frame rate","zh_CN":"主播原始帧率"}
  Ofr *int64 `json:"ofr,omitempty" xml:"ofr,omitempty" require:"true"`
  // {"en":"The resolution of the","zh_CN":"分辨率"}
  Resolution *string `json:"resolution,omitempty" xml:"resolution,omitempty" require:"true"`
  // {"en":"Video coding","zh_CN":"视频编码"}
  VideoCodec *string `json:"video_codec,omitempty" xml:"video_codec,omitempty" require:"true"`
  // {"en":"Audio coding","zh_CN":"音频编码"}
  AudioCodec *string `json:"audio_codec,omitempty" xml:"audio_codec,omitempty" require:"true"`
  // {"en":"Keyframe interval","zh_CN":"关键帧间隔,有传expand才会返回"}
  Gop *int64 `json:"gop,omitempty" xml:"gop,omitempty" require:"true"`
}

func (s QueryLiveStreamStatusResponseDataValue) String() string {
  return tea.Prettify(s)
}

func (s QueryLiveStreamStatusResponseDataValue) GoString() string {
  return s.String()
}

func (s *QueryLiveStreamStatusResponseDataValue) SetStreamname(v string) *QueryLiveStreamStatusResponseDataValue {
  s.Streamname = &v
  return s
}

func (s *QueryLiveStreamStatusResponseDataValue) SetDeployaddress(v string) *QueryLiveStreamStatusResponseDataValue {
  s.Deployaddress = &v
  return s
}

func (s *QueryLiveStreamStatusResponseDataValue) SetInaddress(v string) *QueryLiveStreamStatusResponseDataValue {
  s.Inaddress = &v
  return s
}

func (s *QueryLiveStreamStatusResponseDataValue) SetHists(v int64) *QueryLiveStreamStatusResponseDataValue {
  s.Hists = &v
  return s
}

func (s *QueryLiveStreamStatusResponseDataValue) SetInbandwidth(v int64) *QueryLiveStreamStatusResponseDataValue {
  s.Inbandwidth = &v
  return s
}

func (s *QueryLiveStreamStatusResponseDataValue) SetBandwidth(v int64) *QueryLiveStreamStatusResponseDataValue {
  s.Bandwidth = &v
  return s
}

func (s *QueryLiveStreamStatusResponseDataValue) SetDelay(v int64) *QueryLiveStreamStatusResponseDataValue {
  s.Delay = &v
  return s
}

func (s *QueryLiveStreamStatusResponseDataValue) SetFps(v int64) *QueryLiveStreamStatusResponseDataValue {
  s.Fps = &v
  return s
}

func (s *QueryLiveStreamStatusResponseDataValue) SetLfr(v *QueryLiveStreamStatusResponseDataValueLfr) *QueryLiveStreamStatusResponseDataValue {
  s.Lfr = v
  return s
}

func (s *QueryLiveStreamStatusResponseDataValue) SetOfr(v int64) *QueryLiveStreamStatusResponseDataValue {
  s.Ofr = &v
  return s
}

func (s *QueryLiveStreamStatusResponseDataValue) SetResolution(v string) *QueryLiveStreamStatusResponseDataValue {
  s.Resolution = &v
  return s
}

func (s *QueryLiveStreamStatusResponseDataValue) SetVideoCodec(v string) *QueryLiveStreamStatusResponseDataValue {
  s.VideoCodec = &v
  return s
}

func (s *QueryLiveStreamStatusResponseDataValue) SetAudioCodec(v string) *QueryLiveStreamStatusResponseDataValue {
  s.AudioCodec = &v
  return s
}

func (s *QueryLiveStreamStatusResponseDataValue) SetGop(v int64) *QueryLiveStreamStatusResponseDataValue {
  s.Gop = &v
  return s
}

type QueryLiveStreamStatusResponseDataValueLfr struct {
}

func (s QueryLiveStreamStatusResponseDataValueLfr) String() string {
  return tea.Prettify(s)
}

func (s QueryLiveStreamStatusResponseDataValueLfr) GoString() string {
  return s.String()
}

type QueryLiveStreamStatusResponseHeader struct {
}

func (s QueryLiveStreamStatusResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryLiveStreamStatusResponseHeader) GoString() string {
  return s.String()
}




type ResourceGroupListRequest struct {
  // {"en":"Resource group, multiple separated by semicolon. Do not query all resource groups","zh_CN":"资源组，多个使用分号分隔。不传查询所有"}
  Resources *string `json:"resources,omitempty" xml:"resources,omitempty"`
}

func (s ResourceGroupListRequest) String() string {
  return tea.Prettify(s)
}

func (s ResourceGroupListRequest) GoString() string {
  return s.String()
}

func (s *ResourceGroupListRequest) SetResources(v string) *ResourceGroupListRequest {
  s.Resources = &v
  return s
}

type ResourceGroupListRequestHeader struct {
}

func (s ResourceGroupListRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ResourceGroupListRequestHeader) GoString() string {
  return s.String()
}

type ResourceGroupListPaths struct {
}

func (s ResourceGroupListPaths) String() string {
  return tea.Prettify(s)
}

func (s ResourceGroupListPaths) GoString() string {
  return s.String()
}

type ResourceGroupListParameters struct {
}

func (s ResourceGroupListParameters) String() string {
  return tea.Prettify(s)
}

func (s ResourceGroupListParameters) GoString() string {
  return s.String()
}

type ResourceGroupListResponse struct {
  // {"en":"msg","zh_CN":"返回描述信息"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"code","zh_CN":"返回code"}
  Code *int `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"data","zh_CN":"返回数据"}
  Data []*ResourceGroupListResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s ResourceGroupListResponse) String() string {
  return tea.Prettify(s)
}

func (s ResourceGroupListResponse) GoString() string {
  return s.String()
}

func (s *ResourceGroupListResponse) SetMsg(v string) *ResourceGroupListResponse {
  s.Msg = &v
  return s
}

func (s *ResourceGroupListResponse) SetCode(v int) *ResourceGroupListResponse {
  s.Code = &v
  return s
}

func (s *ResourceGroupListResponse) SetData(v []*ResourceGroupListResponseData) *ResourceGroupListResponse {
  s.Data = v
  return s
}

type ResourceGroupListResponseData struct     {
  // {"en":"Resource Group.","zh_CN":"资源组名称"}
  Resource *string `json:"resource,omitempty" xml:"resource,omitempty" require:"true"`
  // {"en":"Number of domains.","zh_CN":"域名数量."}
  DomainNum *string `json:"domainNum,omitempty" xml:"domainNum,omitempty" require:"true"`
  // {"en":"Resource Group ID.","zh_CN":"资源组ID"}
  ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty" require:"true"`
}

func (s ResourceGroupListResponseData) String() string {
  return tea.Prettify(s)
}

func (s ResourceGroupListResponseData) GoString() string {
  return s.String()
}

func (s *ResourceGroupListResponseData) SetResource(v string) *ResourceGroupListResponseData {
  s.Resource = &v
  return s
}

func (s *ResourceGroupListResponseData) SetDomainNum(v string) *ResourceGroupListResponseData {
  s.DomainNum = &v
  return s
}

func (s *ResourceGroupListResponseData) SetResourceGroupId(v string) *ResourceGroupListResponseData {
  s.ResourceGroupId = &v
  return s
}

type ResourceGroupListResponseHeader struct {
}

func (s ResourceGroupListResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ResourceGroupListResponseHeader) GoString() string {
  return s.String()
}




type QueryOnlineViewerCountRequest struct {
}

func (s QueryOnlineViewerCountRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryOnlineViewerCountRequest) GoString() string {
  return s.String()
}

type QueryOnlineViewerCountRequestHeader struct {
}

func (s QueryOnlineViewerCountRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryOnlineViewerCountRequestHeader) GoString() string {
  return s.String()
}

type QueryOnlineViewerCountPaths struct {
}

func (s QueryOnlineViewerCountPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryOnlineViewerCountPaths) GoString() string {
  return s.String()
}

type QueryOnlineViewerCountParameters struct {
  // {"en":"Domain (multiple domains supported, separated by commas)(    All parameters are passed via HTTP GET requests.)","zh_CN":"域名（支持多个域名，以逗号分隔）(所有参数以HTTP GET请求方式传参)"}
  U *string `json:"u,omitempty" xml:"u,omitempty" require:"true"`
  // {"en":"Time,eg:20160527152300.If the parameter is not specified, the value is 3 minutes","zh_CN":"时间，eg：20160527152300，不填为当前时间-3分钟"}
  T *string `json:"t,omitempty" xml:"t,omitempty"`
  // {"en":"The domain type can be pull or push. If this parameter is not specified, the default value is pull","zh_CN":"域名类型，pull或push，不填时默认为pull"}
  D *string `json:"d,omitempty" xml:"d,omitempty"`
  // {"en":"Channel URL(single channel query only),It is not recommended to query with this parameter, and the performance of range query is poor.","zh_CN":"频道URL(仅支持单频道查询),不建议带该参数查询，范围查询性能较差"}
  Channel *string `json:"channel,omitempty" xml:"channel,omitempty"`
  // {"en":"Start time, eg: 20160803103500, when both from and to are filled or left blank, channel parameter is mandatory. The query time span is two hours at most. If the query time exceeds two hours, the system queries the data within two hours from the start time and the number of online users within the last 7 days","zh_CN":"开始时间，eg: 20160803103500，from和to都填或都不填，都填时channel参数必填，查询时间跨度最大为两个小时，如果超过两个小时，将查询开始时间两个小时内的数据，可查近7天内在线人数数据"}
  From *string `json:"from,omitempty" xml:"from,omitempty"`
  // {"en":"End time, eg: 20160803103900, when both from and to are filled or left blank, channel parameter is mandatory. The query time span is two hours at most. If the query time exceeds two hours, the system queries the data within two hours from the start time and the number of online users in the last 7 days","zh_CN":"结束时间，eg: 20160803103900，from和to都填或都不填，都填时channel参数必填，查询时间跨度最大为两个小时，如果超过两个小时，将查询开始时间两个小时内的数据，可查近7天内在线人数数据"}
  To *string `json:"to,omitempty" xml:"to,omitempty"`
  // {"en":"Query interval, optional value: 10, 60s.\nWhen g is 10, the number of online users every 10 seconds is queried;\nWhen g is 60, the number of online users at the whole minute within the time range is queried.","zh_CN":"查询间隔，可选值10、60s，\n当g为10时，查询时间范围内每10秒的在线人数\n当g为60时，查询时间范围内整分钟点对应的在线人数"}
  G *string `json:"g,omitempty" xml:"g,omitempty"`
  // {"en":"it is query realtime datas,default false","zh_CN":"是否返回实时数据，默认false"}
  Realtime *string `json:"realtime,omitempty" xml:"realtime,omitempty"`
  // {"en":"The default value is false. If the value is true, the domain data is split. If the value is false, the domain data is merged","zh_CN":"域名拆分控制，默认为false，为true时，拆分域名数据，为false时，合并域名数据"}
  Unpack *string `json:"unpack,omitempty" xml:"unpack,omitempty"`
}

func (s QueryOnlineViewerCountParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryOnlineViewerCountParameters) GoString() string {
  return s.String()
}

func (s *QueryOnlineViewerCountParameters) SetU(v string) *QueryOnlineViewerCountParameters {
  s.U = &v
  return s
}

func (s *QueryOnlineViewerCountParameters) SetT(v string) *QueryOnlineViewerCountParameters {
  s.T = &v
  return s
}

func (s *QueryOnlineViewerCountParameters) SetD(v string) *QueryOnlineViewerCountParameters {
  s.D = &v
  return s
}

func (s *QueryOnlineViewerCountParameters) SetChannel(v string) *QueryOnlineViewerCountParameters {
  s.Channel = &v
  return s
}

func (s *QueryOnlineViewerCountParameters) SetFrom(v string) *QueryOnlineViewerCountParameters {
  s.From = &v
  return s
}

func (s *QueryOnlineViewerCountParameters) SetTo(v string) *QueryOnlineViewerCountParameters {
  s.To = &v
  return s
}

func (s *QueryOnlineViewerCountParameters) SetG(v string) *QueryOnlineViewerCountParameters {
  s.G = &v
  return s
}

func (s *QueryOnlineViewerCountParameters) SetRealtime(v string) *QueryOnlineViewerCountParameters {
  s.Realtime = &v
  return s
}

func (s *QueryOnlineViewerCountParameters) SetUnpack(v string) *QueryOnlineViewerCountParameters {
  s.Unpack = &v
  return s
}

type QueryOnlineViewerCountResponse struct {
  // {"en":"Total number of online users. This parameter is displayed only when from and to are empty","zh_CN":"在线总人数，仅当查询时间点，即from和to为空时才显示"}
  Count *int64 `json:"count,omitempty" xml:"count,omitempty" require:"true"`
  // {"en":"Total abnormal online number, only need customers return","zh_CN":"异常总在线人数，只对有需要客户进行返回"}
  ErrorCount *int64 `json:"errorCount,omitempty" xml:"errorCount,omitempty" require:"true"`
  // {"en":"The number of data items is displayed only when from and to are empty.","zh_CN":"数据条数，仅当查询时间点，即from和to为空时才显示"}
  Retcode *int64 `json:"retcode,omitempty" xml:"retcode,omitempty" require:"true"`
  // {"en":"The time of the data returned","zh_CN":"返回的数据的时间"}
  Rettime *string `json:"rettime,omitempty" xml:"rettime,omitempty" require:"true"`
  // {"en":"","zh_CN":"数据集合"}
  DataValue []*QueryOnlineViewerCountResponseDataValue `json:"dataValue,omitempty" xml:"dataValue,omitempty" require:"true" type:"Repeated"`
}

func (s QueryOnlineViewerCountResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryOnlineViewerCountResponse) GoString() string {
  return s.String()
}

func (s *QueryOnlineViewerCountResponse) SetCount(v int64) *QueryOnlineViewerCountResponse {
  s.Count = &v
  return s
}

func (s *QueryOnlineViewerCountResponse) SetErrorCount(v int64) *QueryOnlineViewerCountResponse {
  s.ErrorCount = &v
  return s
}

func (s *QueryOnlineViewerCountResponse) SetRetcode(v int64) *QueryOnlineViewerCountResponse {
  s.Retcode = &v
  return s
}

func (s *QueryOnlineViewerCountResponse) SetRettime(v string) *QueryOnlineViewerCountResponse {
  s.Rettime = &v
  return s
}

func (s *QueryOnlineViewerCountResponse) SetDataValue(v []*QueryOnlineViewerCountResponseDataValue) *QueryOnlineViewerCountResponse {
  s.DataValue = v
  return s
}

type QueryOnlineViewerCountResponseDataValue struct     {
  // {"en":"Stream name","zh_CN":"流名"}
  Prog *string `json:"prog,omitempty" xml:"prog,omitempty" require:"true"`
  // {"en":"Time is displayed only if the query time range (channel, FROM, and to parameters) is not empty","zh_CN":"时间，仅当查询时间范围，即channel，from和to参数不为空时才显示"}
  Time *string `json:"time,omitempty" xml:"time,omitempty" require:"true"`
  // {"en":"The number of online","zh_CN":"在线人数"}
  Value *int64 `json:"value,omitempty" xml:"value,omitempty" require:"true"`
  // {"en":"Abnormal online number, only need customers return","zh_CN":"异常在线人数，只对需要客户进行返回"}
  ErrorValue *int64 `json:"errorValue,omitempty" xml:"errorValue,omitempty" require:"true"`
}

func (s QueryOnlineViewerCountResponseDataValue) String() string {
  return tea.Prettify(s)
}

func (s QueryOnlineViewerCountResponseDataValue) GoString() string {
  return s.String()
}

func (s *QueryOnlineViewerCountResponseDataValue) SetProg(v string) *QueryOnlineViewerCountResponseDataValue {
  s.Prog = &v
  return s
}

func (s *QueryOnlineViewerCountResponseDataValue) SetTime(v string) *QueryOnlineViewerCountResponseDataValue {
  s.Time = &v
  return s
}

func (s *QueryOnlineViewerCountResponseDataValue) SetValue(v int64) *QueryOnlineViewerCountResponseDataValue {
  s.Value = &v
  return s
}

func (s *QueryOnlineViewerCountResponseDataValue) SetErrorValue(v int64) *QueryOnlineViewerCountResponseDataValue {
  s.ErrorValue = &v
  return s
}

type QueryOnlineViewerCountResponseHeader struct {
}

func (s QueryOnlineViewerCountResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryOnlineViewerCountResponseHeader) GoString() string {
  return s.String()
}




type QueryEdgeActiveConnectionsForMultiDomainsRequest struct {
  // {"en":"Start time: 
  // 	1. The format is yyyyy-MM-ddTHH: mm: SS + 08:00, for example, 2019-12-02T10:00:00+08:00 (10:00:00 Beijing time on December 2, 2019); 
  // 	2. Can not exceed the current time; 
  //     3. The latest half year (183 days) data can be obtained at most.
  // ", "zh_CN":"开始时间:
  // 1.时间格式为yyyy-MM-ddTHH:mm:ss+08:00,例如,2019-12-02T10:00:00+08:00(为北京时间2019年12月2日10点0分0秒);
  // 2.不能大于当前时间
  // 3.最多可获取最近半年(183天)的数据。"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty"`
  // {"en":"End time: 
  // 	1. The 1format is yyyy-MM-ddTHH:mm:ss+08:00; 	
  // 	2. The end time is greater than the start time. If the end time is greater than the current time, the current time is taken. 
  // 	3. DateFrom and dateTo are not uploaded, defaulting to query the past 24 hours; if only one is not uploaded, throw an exception; 
  // 	4. Maximum query interval allowed: 42 days, that is, the difference between dateFrom and dateTo can not exceed 31 days.
  // ", "zh_CN":"结束时间:
  // 1.时间格式为yyyy-MM-ddTHH:mm:ss+08:00,例如,2019-01-02T10:00:00+08:00
  // 2.结束时间需大于开始时间,结束时间如果大于当前时间,取当前时间。
  // 3.dateFrom,dateTo二者都未传,默认查询过去的24小时;如仅有一个未传,抛异常
  // 4.允许查询最大时间间隔:31天,即dateFrom和dateTo相差不能超过31天。"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty"`
  // {"en":"domain name: 
  // 	1. The maximum number of deliverable domain names is 20 by default (can be contacted by technical support); 
  // 	2. Automatically filter out illegal domain names (such as passing illegal domain names, they will be filtered out, and the query results only return data of legitimate domain names);
  // 	3. Domain is not uploaded, Query all domain names of the account.
  // ", "zh_CN":"域名:
  // 
  // 可传递域名数量上限默认为20个(可联系技术支持调整)
  // 自动过滤掉非法域名(如传递非法域名,会被过滤掉,查询结果只返回合法域名的数据)
  // 未传递该入参时,默认查询账号下所有域名,但当账号下域名数量超过上限时提示错误"}
  Domain []*string `json:"domain,omitempty" xml:"domain,omitempty" type:"Repeated"`
}

func (s QueryEdgeActiveConnectionsForMultiDomainsRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryEdgeActiveConnectionsForMultiDomainsRequest) GoString() string {
  return s.String()
}

func (s *QueryEdgeActiveConnectionsForMultiDomainsRequest) SetDateFrom(v string) *QueryEdgeActiveConnectionsForMultiDomainsRequest {
  s.DateFrom = &v
  return s
}

func (s *QueryEdgeActiveConnectionsForMultiDomainsRequest) SetDateTo(v string) *QueryEdgeActiveConnectionsForMultiDomainsRequest {
  s.DateTo = &v
  return s
}

func (s *QueryEdgeActiveConnectionsForMultiDomainsRequest) SetDomain(v []*string) *QueryEdgeActiveConnectionsForMultiDomainsRequest {
  s.Domain = v
  return s
}

type QueryEdgeActiveConnectionsForMultiDomainsResponse struct {
  // {"en":"request result status code", "zh_CN":"请求结果状态码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Request result information", "zh_CN":"请求结果信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  QueryEdgeActiveConnectionsForMultiDomainsData []*QueryEdgeActiveConnectionsForMultiDomainsData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s QueryEdgeActiveConnectionsForMultiDomainsResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryEdgeActiveConnectionsForMultiDomainsResponse) GoString() string {
  return s.String()
}

func (s *QueryEdgeActiveConnectionsForMultiDomainsResponse) SetCode(v string) *QueryEdgeActiveConnectionsForMultiDomainsResponse {
  s.Code = &v
  return s
}

func (s *QueryEdgeActiveConnectionsForMultiDomainsResponse) SetMessage(v string) *QueryEdgeActiveConnectionsForMultiDomainsResponse {
  s.Message = &v
  return s
}

func (s *QueryEdgeActiveConnectionsForMultiDomainsResponse) SetData(v []*QueryEdgeActiveConnectionsForMultiDomainsData) *QueryEdgeActiveConnectionsForMultiDomainsResponse {
  s.QueryEdgeActiveConnectionsForMultiDomainsData = v
  return s
}

type QueryEdgeActiveConnectionsForMultiDomainsData struct {
  // {"en":"Time slice, the format is yyyy-MM-dd HH:mm", "zh_CN":"时间片,格式为yyyy-MM-dd HH:mm"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {"en":"The number of instantaneous active connections at the moment of 5 minutes", "zh_CN":"整5分钟的时刻的瞬间活跃连接数"}
  ActiveConn *string `json:"activeConn,omitempty" xml:"activeConn,omitempty" require:"true"`
}

func (s QueryEdgeActiveConnectionsForMultiDomainsData) String() string {
  return tea.Prettify(s)
}

func (s QueryEdgeActiveConnectionsForMultiDomainsData) GoString() string {
  return s.String()
}

func (s *QueryEdgeActiveConnectionsForMultiDomainsData) SetTimestamp(v string) *QueryEdgeActiveConnectionsForMultiDomainsData {
  s.Timestamp = &v
  return s
}

func (s *QueryEdgeActiveConnectionsForMultiDomainsData) SetActiveConn(v string) *QueryEdgeActiveConnectionsForMultiDomainsData {
  s.ActiveConn = &v
  return s
}

type QueryEdgeActiveConnectionsForMultiDomainsPaths struct {
}

func (s QueryEdgeActiveConnectionsForMultiDomainsPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryEdgeActiveConnectionsForMultiDomainsPaths) GoString() string {
  return s.String()
}

type QueryEdgeActiveConnectionsForMultiDomainsParameters struct {
}

func (s QueryEdgeActiveConnectionsForMultiDomainsParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryEdgeActiveConnectionsForMultiDomainsParameters) GoString() string {
  return s.String()
}

type QueryEdgeActiveConnectionsForMultiDomainsRequestHeader struct {
}

func (s QueryEdgeActiveConnectionsForMultiDomainsRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryEdgeActiveConnectionsForMultiDomainsRequestHeader) GoString() string {
  return s.String()
}

type QueryEdgeActiveConnectionsForMultiDomainsResponseHeader struct {
}

func (s QueryEdgeActiveConnectionsForMultiDomainsResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryEdgeActiveConnectionsForMultiDomainsResponseHeader) GoString() string {
  return s.String()
}




