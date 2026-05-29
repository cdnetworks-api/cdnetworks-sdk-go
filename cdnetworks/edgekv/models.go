package edgekv

// This file is auto-generated, don't edit it. Thanks.
import (
  "github.com/alibabacloud-go/tea/tea"
)

type GetShortUrlRequest struct {
  // {"en":"ShortUrl","zh_CN":"短网址"}
  ShortUrl *string `json:"ShortUrl,omitempty" xml:"ShortUrl,omitempty" require:"true"`
}

func (s GetShortUrlRequest) String() string {
  return tea.Prettify(s)
}

func (s GetShortUrlRequest) GoString() string {
  return s.String()
}

func (s *GetShortUrlRequest) SetShortUrl(v string) *GetShortUrlRequest {
  s.ShortUrl = &v
  return s
}

type GetShortUrlRequestHeader struct {
}

func (s GetShortUrlRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetShortUrlRequestHeader) GoString() string {
  return s.String()
}

type GetShortUrlPaths struct {
}

func (s GetShortUrlPaths) String() string {
  return tea.Prettify(s)
}

func (s GetShortUrlPaths) GoString() string {
  return s.String()
}

type GetShortUrlParameters struct {
}

func (s GetShortUrlParameters) String() string {
  return tea.Prettify(s)
}

func (s GetShortUrlParameters) GoString() string {
  return s.String()
}

type GetShortUrlResponse struct {
  // {"en":"Code","zh_CN":"状态码"}
  Code *int `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
  // {"en":"Msg","zh_CN":"返回信息"}
  Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty" require:"true"`
  // {"en":"LongUrl","zh_CN":"长网址"}
  LongUrl *string `json:"LongUrl,omitempty" xml:"LongUrl,omitempty" require:"true"`
}

func (s GetShortUrlResponse) String() string {
  return tea.Prettify(s)
}

func (s GetShortUrlResponse) GoString() string {
  return s.String()
}

func (s *GetShortUrlResponse) SetCode(v int) *GetShortUrlResponse {
  s.Code = &v
  return s
}

func (s *GetShortUrlResponse) SetMsg(v string) *GetShortUrlResponse {
  s.Msg = &v
  return s
}

func (s *GetShortUrlResponse) SetLongUrl(v string) *GetShortUrlResponse {
  s.LongUrl = &v
  return s
}

type GetShortUrlResponseHeader struct {
}

func (s GetShortUrlResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetShortUrlResponseHeader) GoString() string {
  return s.String()
}




type GetKeyValueKeyValue struct {
  // {'en':'key', 'zh_CN':'key'}
  Key *string `json:"key,omitempty" xml:"key,omitempty" require:"true"`
  // {'en':'value,base64 encoded', 'zh_CN':'base64 encode过的value'}
  Value *string `json:"value,omitempty" xml:"value,omitempty" require:"true"`
  // {'en':'expiration', 'zh_CN':'过期时间，unix 时间戳'}
  Exp *int32 `json:"exp,omitempty" xml:"exp,omitempty" require:"true"`
}

func (s GetKeyValueKeyValue) String() string {
  return tea.Prettify(s)
}

func (s GetKeyValueKeyValue) GoString() string {
  return s.String()
}

func (s *GetKeyValueKeyValue) SetKey(v string) *GetKeyValueKeyValue {
  s.Key = &v
  return s
}

func (s *GetKeyValueKeyValue) SetValue(v string) *GetKeyValueKeyValue {
  s.Value = &v
  return s
}

func (s *GetKeyValueKeyValue) SetExp(v int32) *GetKeyValueKeyValue {
  s.Exp = &v
  return s
}

type GetKeyValueRequest struct {
  // {'en':'spaceName', 'zh_CN':'空间名称'}
  SpaceName *string `json:"spaceName,omitempty" xml:"spaceName,omitempty" require:"true"`
  // {'en':'key to query', 'zh_CN':'要查询的key'}
  Key *string `json:"key,omitempty" xml:"key,omitempty" require:"true"`
}

func (s GetKeyValueRequest) String() string {
  return tea.Prettify(s)
}

func (s GetKeyValueRequest) GoString() string {
  return s.String()
}

func (s *GetKeyValueRequest) SetSpaceName(v string) *GetKeyValueRequest {
  s.SpaceName = &v
  return s
}

func (s *GetKeyValueRequest) SetKey(v string) *GetKeyValueRequest {
  s.Key = &v
  return s
}

type GetKeyValueResponse struct {
  // {'en':'code', 'zh_CN':'状态码'}
  Code *int32 `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {'en':'message', 'zh_CN':'返回信息'}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {'en':'result', 'zh_CN':'结果'}
  Result *GetKeyValueKeyValue `json:"result,omitempty" xml:"result,omitempty" require:"true"`
}

func (s GetKeyValueResponse) String() string {
  return tea.Prettify(s)
}

func (s GetKeyValueResponse) GoString() string {
  return s.String()
}

func (s *GetKeyValueResponse) SetCode(v int32) *GetKeyValueResponse {
  s.Code = &v
  return s
}

func (s *GetKeyValueResponse) SetMsg(v string) *GetKeyValueResponse {
  s.Msg = &v
  return s
}

func (s *GetKeyValueResponse) SetResult(v *GetKeyValueKeyValue) *GetKeyValueResponse {
  s.Result = v
  return s
}

type GetKeyValuePaths struct {
}

func (s GetKeyValuePaths) String() string {
  return tea.Prettify(s)
}

func (s GetKeyValuePaths) GoString() string {
  return s.String()
}

type GetKeyValueParameters struct {
}

func (s GetKeyValueParameters) String() string {
  return tea.Prettify(s)
}

func (s GetKeyValueParameters) GoString() string {
  return s.String()
}

type GetKeyValueRequestHeader struct {
}

func (s GetKeyValueRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetKeyValueRequestHeader) GoString() string {
  return s.String()
}

type GetKeyValueResponseHeader struct {
}

func (s GetKeyValueResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetKeyValueResponseHeader) GoString() string {
  return s.String()
}




type DeleteKeyValueRequest struct {
  // {'en':'spaceName', 'zh_CN':'空间名称'}
  SpaceName *string `json:"spaceName,omitempty" xml:"spaceName,omitempty" require:"true"`
  // {'en':'keys to delete', 'zh_CN':'要删除的key'}
  Keys []*string `json:"keys,omitempty" xml:"keys,omitempty" require:"true" type:"Repeated"`
}

func (s DeleteKeyValueRequest) String() string {
  return tea.Prettify(s)
}

func (s DeleteKeyValueRequest) GoString() string {
  return s.String()
}

func (s *DeleteKeyValueRequest) SetSpaceName(v string) *DeleteKeyValueRequest {
  s.SpaceName = &v
  return s
}

func (s *DeleteKeyValueRequest) SetKeys(v []*string) *DeleteKeyValueRequest {
  s.Keys = v
  return s
}

type DeleteKeyValueResponse struct {
  // {'en':'code', 'zh_CN':'状态码'}
  Code *int32 `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {'en':'message', 'zh_CN':'返回信息'}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
}

func (s DeleteKeyValueResponse) String() string {
  return tea.Prettify(s)
}

func (s DeleteKeyValueResponse) GoString() string {
  return s.String()
}

func (s *DeleteKeyValueResponse) SetCode(v int32) *DeleteKeyValueResponse {
  s.Code = &v
  return s
}

func (s *DeleteKeyValueResponse) SetMsg(v string) *DeleteKeyValueResponse {
  s.Msg = &v
  return s
}

type DeleteKeyValuePaths struct {
}

func (s DeleteKeyValuePaths) String() string {
  return tea.Prettify(s)
}

func (s DeleteKeyValuePaths) GoString() string {
  return s.String()
}

type DeleteKeyValueParameters struct {
}

func (s DeleteKeyValueParameters) String() string {
  return tea.Prettify(s)
}

func (s DeleteKeyValueParameters) GoString() string {
  return s.String()
}

type DeleteKeyValueRequestHeader struct {
}

func (s DeleteKeyValueRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteKeyValueRequestHeader) GoString() string {
  return s.String()
}

type DeleteKeyValueResponseHeader struct {
}

func (s DeleteKeyValueResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteKeyValueResponseHeader) GoString() string {
  return s.String()
}




type EcaKvInfoRequest struct {
  // {"en":"cust_en_name of sub-client.\nWhen a merged-account wants to  view the information of the subclient,the cust_en_name is required.","zh_CN":"合并账号下的某个客户的英文名，当合并账号要查看子客户的信息时，必须填写子客户的英文名"}
  Cust *string `json:"cust,omitempty" xml:"cust,omitempty"`
  // {"en":"Specifies the query date:\n1.With format yyyy-mm-dd.\n2.If not Specifies,it means today as default.","zh_CN":"查询的日期，日期格式为yyyy-mm-dd,不选或者为空时默认为当天；"}
  Date *string `json:"date,omitempty" xml:"date,omitempty"`
  // {"en":"1.Must work with 'enddate' and they  specify the query date scope.\n2.With format yyyy-mm-dd hh:MM.If 'hh:MM' not specified,it means '00:01'.\n3.If there is a 'date' parameter,this parameter will be omitted.","zh_CN":"查询的起始日期,精确到分钟,日期格式为yyyy-mm-dd hh:MM若没有输入时、分，则时分默认为00:01；此参数需与enddate参数配合,若存在date参数,则该参数无效"}
  Startdate *string `json:"startdate,omitempty" xml:"startdate,omitempty"`
  // {"en":"1.Must work with 'startdate' and they  specify the query date scope.\n2.With format yyyy-mm-dd hh:MM.If 'hh:MM' not specified,it means '24:00'.\n3.If there is a 'date' parameter,this parameter will be omitted.","zh_CN":"查询的结束日期,精确到分钟,日期格式为yyyy-mm-dd hh:MM,若没有输入时、分，则时分默认为24:00；此参数需与startdate参数配合,若存在date参数,则该参数无效。"}
  Enddate *string `json:"enddate,omitempty" xml:"enddate,omitempty"`
  // {"en":"1.If there are multiple inputs,use ';' as separator.For example,u can use 'region=cn;apac' to query data of cn and apac region.\n2.If not specified, it means all the regions.","zh_CN":"查询的加速区域的缩写，多个区域请用英文分号';'分隔开，如查询大陆及亚太区域，参数填写为：'region=cn;apac'。不选或者为空时默认为全部区域。"}
  Region *string `json:"region,omitempty" xml:"region,omitempty"`
  // {"en":"Space names, multiple names should be separated by a semicolon ';'.","zh_CN":"空间名称，多个请用英文分号“;”分隔开。"}
  Space *string `json:"space,omitempty" xml:"space,omitempty"`
  // {"en":"Greenwich Mean Time (GMT) zone, the parameter format is GMT+09:00 to indicate East 9th Zone, and GMT-09:00 to indicate West 9th Zone. If not provided, the default is the local time zone (East 8th Zone).","zh_CN":"格林尼治时区，参数格式 GMT+09:00 表示东九区，GMT-09:00 表示西9区，不传则默认为本地时区（东八区）"}
  Timezone *string `json:"timezone,omitempty" xml:"timezone,omitempty"`
  // {"en":"Return result format, supported formats are XML and JSON, default is XML.","zh_CN":"返回结果格式,支持格式为xml和json,默认为xml"}
  Dataformat *string `json:"dataformat,omitempty" xml:"dataformat,omitempty"`
  // {"en":"4 types. 0: Storage capacity; 1: Read request count; 2: Write request count; 3: Delete request count. Default is: 0.","zh_CN":"4种类型。0：存储量；1：读请求数；2：写请求数；3：删请求数。默认取：0"}
  Datatype *string `json:"datatype,omitempty" xml:"datatype,omitempty"`
  // {"en":"Whether to aggregate in a specific manner, format: number_day|hour. For example, 3_hour means to aggregate by 3 hours; 2_day means to aggregate by 2 days.","zh_CN":"是否按照特定方式聚合,格式 :  数字_day|hour .例如 3_hour表示按照3小时聚合; 2_day表示按照2天聚合"}
  ReturnType *string `json:"returnType,omitempty" xml:"returnType,omitempty"`
}

func (s EcaKvInfoRequest) String() string {
  return tea.Prettify(s)
}

func (s EcaKvInfoRequest) GoString() string {
  return s.String()
}

func (s *EcaKvInfoRequest) SetCust(v string) *EcaKvInfoRequest {
  s.Cust = &v
  return s
}

func (s *EcaKvInfoRequest) SetDate(v string) *EcaKvInfoRequest {
  s.Date = &v
  return s
}

func (s *EcaKvInfoRequest) SetStartdate(v string) *EcaKvInfoRequest {
  s.Startdate = &v
  return s
}

func (s *EcaKvInfoRequest) SetEnddate(v string) *EcaKvInfoRequest {
  s.Enddate = &v
  return s
}

func (s *EcaKvInfoRequest) SetRegion(v string) *EcaKvInfoRequest {
  s.Region = &v
  return s
}

func (s *EcaKvInfoRequest) SetSpace(v string) *EcaKvInfoRequest {
  s.Space = &v
  return s
}

func (s *EcaKvInfoRequest) SetTimezone(v string) *EcaKvInfoRequest {
  s.Timezone = &v
  return s
}

func (s *EcaKvInfoRequest) SetDataformat(v string) *EcaKvInfoRequest {
  s.Dataformat = &v
  return s
}

func (s *EcaKvInfoRequest) SetDatatype(v string) *EcaKvInfoRequest {
  s.Datatype = &v
  return s
}

func (s *EcaKvInfoRequest) SetReturnType(v string) *EcaKvInfoRequest {
  s.ReturnType = &v
  return s
}

type EcaKvInfoRequestHeader struct {
}

func (s EcaKvInfoRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s EcaKvInfoRequestHeader) GoString() string {
  return s.String()
}

type EcaKvInfoPaths struct {
}

func (s EcaKvInfoPaths) String() string {
  return tea.Prettify(s)
}

func (s EcaKvInfoPaths) GoString() string {
  return s.String()
}

type EcaKvInfoParameters struct {
}

func (s EcaKvInfoParameters) String() string {
  return tea.Prettify(s)
}

func (s EcaKvInfoParameters) GoString() string {
  return s.String()
}

type EcaKvInfoResponse struct {
  // {"en":"provider","zh_CN":"结果"}
  Provider *string `json:"provider,omitempty" xml:"provider,omitempty" require:"true"`
  // {"en":"Peak time","zh_CN":"峰值时间"}
  PeakTime *string `json:"peakTime,omitempty" xml:"peakTime,omitempty" require:"true"`
  // {"en":"Peak, unit GB","zh_CN":"峰值，单位GB"}
  PeakValue *string `json:"peakValue,omitempty" xml:"peakValue,omitempty" require:"true"`
  // {"en":"Peak average, unit GB","zh_CN":"峰值平均，单位GB"}
  PeakAvgValue *string `json:"peakAvgValue,omitempty" xml:"peakAvgValue,omitempty" require:"true"`
  // {"en":"total hits","zh_CN":"总请求数"}
  TotalHit *string `json:"totalHit,omitempty" xml:"totalHit,omitempty" require:"true"`
  // {"en":"Time point, format yyyy-mm-dd hh:MM:ss","zh_CN":"时间点，格式yyyy-mm-dd hh:MM:ss"}
  Time *string `json:"time,omitempty" xml:"time,omitempty" require:"true"`
  // {"en":"Value corresponding to the time point","zh_CN":"时间点对应的值"}
  Text *string `json:"text,omitempty" xml:"text,omitempty" require:"true"`
}

func (s EcaKvInfoResponse) String() string {
  return tea.Prettify(s)
}

func (s EcaKvInfoResponse) GoString() string {
  return s.String()
}

func (s *EcaKvInfoResponse) SetProvider(v string) *EcaKvInfoResponse {
  s.Provider = &v
  return s
}

func (s *EcaKvInfoResponse) SetPeakTime(v string) *EcaKvInfoResponse {
  s.PeakTime = &v
  return s
}

func (s *EcaKvInfoResponse) SetPeakValue(v string) *EcaKvInfoResponse {
  s.PeakValue = &v
  return s
}

func (s *EcaKvInfoResponse) SetPeakAvgValue(v string) *EcaKvInfoResponse {
  s.PeakAvgValue = &v
  return s
}

func (s *EcaKvInfoResponse) SetTotalHit(v string) *EcaKvInfoResponse {
  s.TotalHit = &v
  return s
}

func (s *EcaKvInfoResponse) SetTime(v string) *EcaKvInfoResponse {
  s.Time = &v
  return s
}

func (s *EcaKvInfoResponse) SetText(v string) *EcaKvInfoResponse {
  s.Text = &v
  return s
}

type EcaKvInfoResponseHeader struct {
}

func (s EcaKvInfoResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s EcaKvInfoResponseHeader) GoString() string {
  return s.String()
}




type CreateShortUrlRequest struct {
  // {"en":"LongUrls","zh_CN":"长网址（原网址）"}
  LongUrls []*string `json:"LongUrls,omitempty" xml:"LongUrls,omitempty" require:"true" type:"Repeated"`
  // {"en":"Ttl","zh_CN":"短网址有效时长(单位秒)"}
  Ttl *int `json:"Ttl,omitempty" xml:"Ttl,omitempty" require:"true"`
}

func (s CreateShortUrlRequest) String() string {
  return tea.Prettify(s)
}

func (s CreateShortUrlRequest) GoString() string {
  return s.String()
}

func (s *CreateShortUrlRequest) SetLongUrls(v []*string) *CreateShortUrlRequest {
  s.LongUrls = v
  return s
}

func (s *CreateShortUrlRequest) SetTtl(v int) *CreateShortUrlRequest {
  s.Ttl = &v
  return s
}

type CreateShortUrlRequestHeader struct {
}

func (s CreateShortUrlRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s CreateShortUrlRequestHeader) GoString() string {
  return s.String()
}

type CreateShortUrlPaths struct {
}

func (s CreateShortUrlPaths) String() string {
  return tea.Prettify(s)
}

func (s CreateShortUrlPaths) GoString() string {
  return s.String()
}

type CreateShortUrlParameters struct {
}

func (s CreateShortUrlParameters) String() string {
  return tea.Prettify(s)
}

func (s CreateShortUrlParameters) GoString() string {
  return s.String()
}

type CreateShortUrlResponse struct {
  // {"en":"Code","zh_CN":"状态码"}
  Code *int `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
  // {"en":"Msg","zh_CN":"返回信息"}
  Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty" require:"true"`
  // {"en":"ShortUrls","zh_CN":"短网址"}
  ShortUrls []*CreateShortUrlResponseShortUrls `json:"ShortUrls,omitempty" xml:"ShortUrls,omitempty" require:"true" type:"Repeated"`
}

func (s CreateShortUrlResponse) String() string {
  return tea.Prettify(s)
}

func (s CreateShortUrlResponse) GoString() string {
  return s.String()
}

func (s *CreateShortUrlResponse) SetCode(v int) *CreateShortUrlResponse {
  s.Code = &v
  return s
}

func (s *CreateShortUrlResponse) SetMsg(v string) *CreateShortUrlResponse {
  s.Msg = &v
  return s
}

func (s *CreateShortUrlResponse) SetShortUrls(v []*CreateShortUrlResponseShortUrls) *CreateShortUrlResponse {
  s.ShortUrls = v
  return s
}

type CreateShortUrlResponseShortUrls struct     {
  // {"en":"Code","zh_CN":"状态码"}
  Code *int `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
  // {"en":"Msg","zh_CN":"返回信息"}
  Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty" require:"true"`
  // {"en":"ShortUrl","zh_CN":"短网址"}
  ShortUrl *string `json:"ShortUrl,omitempty" xml:"ShortUrl,omitempty" require:"true"`
  // {"en":"LongUrl","zh_CN":"长网址（原网址）"}
  LongUrl *string `json:"LongUrl,omitempty" xml:"LongUrl,omitempty" require:"true"`
}

func (s CreateShortUrlResponseShortUrls) String() string {
  return tea.Prettify(s)
}

func (s CreateShortUrlResponseShortUrls) GoString() string {
  return s.String()
}

func (s *CreateShortUrlResponseShortUrls) SetCode(v int) *CreateShortUrlResponseShortUrls {
  s.Code = &v
  return s
}

func (s *CreateShortUrlResponseShortUrls) SetMsg(v string) *CreateShortUrlResponseShortUrls {
  s.Msg = &v
  return s
}

func (s *CreateShortUrlResponseShortUrls) SetShortUrl(v string) *CreateShortUrlResponseShortUrls {
  s.ShortUrl = &v
  return s
}

func (s *CreateShortUrlResponseShortUrls) SetLongUrl(v string) *CreateShortUrlResponseShortUrls {
  s.LongUrl = &v
  return s
}

type CreateShortUrlResponseHeader struct {
}

func (s CreateShortUrlResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s CreateShortUrlResponseHeader) GoString() string {
  return s.String()
}




type DelShortUrlRequest struct {
  // {"en":"ShortUrl","zh_CN":"短网址"}
  ShortUrl *string `json:"ShortUrl,omitempty" xml:"ShortUrl,omitempty" require:"true"`
}

func (s DelShortUrlRequest) String() string {
  return tea.Prettify(s)
}

func (s DelShortUrlRequest) GoString() string {
  return s.String()
}

func (s *DelShortUrlRequest) SetShortUrl(v string) *DelShortUrlRequest {
  s.ShortUrl = &v
  return s
}

type DelShortUrlRequestHeader struct {
}

func (s DelShortUrlRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s DelShortUrlRequestHeader) GoString() string {
  return s.String()
}

type DelShortUrlPaths struct {
}

func (s DelShortUrlPaths) String() string {
  return tea.Prettify(s)
}

func (s DelShortUrlPaths) GoString() string {
  return s.String()
}

type DelShortUrlParameters struct {
}

func (s DelShortUrlParameters) String() string {
  return tea.Prettify(s)
}

func (s DelShortUrlParameters) GoString() string {
  return s.String()
}

type DelShortUrlResponse struct {
  // {"en":"Code","zh_CN":"状态码"}
  Code *int `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
  // {"en":"Msg","zh_CN":"返回信息"}
  Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty" require:"true"`
}

func (s DelShortUrlResponse) String() string {
  return tea.Prettify(s)
}

func (s DelShortUrlResponse) GoString() string {
  return s.String()
}

func (s *DelShortUrlResponse) SetCode(v int) *DelShortUrlResponse {
  s.Code = &v
  return s
}

func (s *DelShortUrlResponse) SetMsg(v string) *DelShortUrlResponse {
  s.Msg = &v
  return s
}

type DelShortUrlResponseHeader struct {
}

func (s DelShortUrlResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s DelShortUrlResponseHeader) GoString() string {
  return s.String()
}




type SetKeyValueKeyValue struct {
  // {'en':'key', 'zh_CN':'key'}
  Key *string `json:"key,omitempty" xml:"key,omitempty" require:"true"`
  // {'en':'value', 'zh_CN':'value,如果是图片等可以先进行bease64 encode。'}
  Value *string `json:"value,omitempty" xml:"value,omitempty" require:"true"`
  // {'en':'Whether or not the server should base64 decode the value before storing it.', 'zh_CN':'是否需要对value进行base64 decode后再存储'}
  Base64 *bool `json:"base64,omitempty" xml:"base64,omitempty"`
  // {'en':'The time, measured in number of seconds since the UNIX epoch, at which the key should expire.', 'zh_CN':'过期时间戳，unix时间戳格式'}
  Exp *int32 `json:"exp,omitempty" xml:"exp,omitempty"`
}

func (s SetKeyValueKeyValue) String() string {
  return tea.Prettify(s)
}

func (s SetKeyValueKeyValue) GoString() string {
  return s.String()
}

func (s *SetKeyValueKeyValue) SetKey(v string) *SetKeyValueKeyValue {
  s.Key = &v
  return s
}

func (s *SetKeyValueKeyValue) SetValue(v string) *SetKeyValueKeyValue {
  s.Value = &v
  return s
}

func (s *SetKeyValueKeyValue) SetBase64(v bool) *SetKeyValueKeyValue {
  s.Base64 = &v
  return s
}

func (s *SetKeyValueKeyValue) SetExp(v int32) *SetKeyValueKeyValue {
  s.Exp = &v
  return s
}

type SetKeyValueRequest struct {
  // {'en':'spaceName', 'zh_CN':'空间名称'}
  SpaceName *string `json:"spaceName,omitempty" xml:"spaceName,omitempty" require:"true"`
  // {'en':'key value pairs', 'zh_CN':'key value 数据'}
  Data []*SetKeyValueKeyValue `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s SetKeyValueRequest) String() string {
  return tea.Prettify(s)
}

func (s SetKeyValueRequest) GoString() string {
  return s.String()
}

func (s *SetKeyValueRequest) SetSpaceName(v string) *SetKeyValueRequest {
  s.SpaceName = &v
  return s
}

func (s *SetKeyValueRequest) SetData(v []*SetKeyValueKeyValue) *SetKeyValueRequest {
  s.Data = v
  return s
}

type SetKeyValueResponse struct {
  // {'en':'code', 'zh_CN':'状态码'}
  Code *int32 `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {'en':'message', 'zh_CN':'返回信息'}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
}

func (s SetKeyValueResponse) String() string {
  return tea.Prettify(s)
}

func (s SetKeyValueResponse) GoString() string {
  return s.String()
}

func (s *SetKeyValueResponse) SetCode(v int32) *SetKeyValueResponse {
  s.Code = &v
  return s
}

func (s *SetKeyValueResponse) SetMsg(v string) *SetKeyValueResponse {
  s.Msg = &v
  return s
}

type SetKeyValuePaths struct {
}

func (s SetKeyValuePaths) String() string {
  return tea.Prettify(s)
}

func (s SetKeyValuePaths) GoString() string {
  return s.String()
}

type SetKeyValueParameters struct {
}

func (s SetKeyValueParameters) String() string {
  return tea.Prettify(s)
}

func (s SetKeyValueParameters) GoString() string {
  return s.String()
}

type SetKeyValueRequestHeader struct {
}

func (s SetKeyValueRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s SetKeyValueRequestHeader) GoString() string {
  return s.String()
}

type SetKeyValueResponseHeader struct {
}

func (s SetKeyValueResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s SetKeyValueResponseHeader) GoString() string {
  return s.String()
}




