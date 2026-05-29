package streamforbid

// This file is auto-generated, don't edit it. Thanks.
import (
  "github.com/alibabacloud-go/tea/tea"
)

type StopLivestreamingRequest struct {
  // {"en":"Push  or pull  address complete url, URL format requirements:
  // 1) The URL must start with'http:/'or'https:/' or'rtmp', and enter an example: http://www.a.com/live/201802221082.
  // 2) The maximum length of each URL is 1000 characters.
  // 3) The domain name of each URL must be the domain name of our company which is speeded up by live broadcasting.
  // 4) If a question mark is included in the url, only the URL before the question mark is submitted.", "zh_CN":"推流或拉流的地址完整的url，url的格式要求：
  // 1）URL 必须以'http://' 或 'https://' 或&lsquo;rtmp&rsquo;开头，输入示例：http://www.a.com/live/201802221082。
  // 2）每个url最大长度 1000 字符。
  // 3）每个url所在的域名必须是在我司直播加速的域名。
  // 4）url中如果包含问号，则只需提交问号前面的url。"}
  LiveUrl *string `json:"liveUrl,omitempty" xml:"liveUrl,omitempty" require:"true"`
  // {'en':'The values are: play (back-source pull), publish (live push).', 'zh_CN':'值为：play (回源拉流)，publish(直播推流)。'}
  Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s StopLivestreamingRequest) String() string {
  return tea.Prettify(s)
}

func (s StopLivestreamingRequest) GoString() string {
  return s.String()
}

func (s *StopLivestreamingRequest) SetLiveUrl(v string) *StopLivestreamingRequest {
  s.LiveUrl = &v
  return s
}

func (s *StopLivestreamingRequest) SetType(v string) *StopLivestreamingRequest {
  s.Type = &v
  return s
}

type StopLivestreamingResponse struct {
  // {'en':'message', 'zh_CN':'错误信息'}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {'en':'code', 'zh_CN':'错误码'}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
}

func (s StopLivestreamingResponse) String() string {
  return tea.Prettify(s)
}

func (s StopLivestreamingResponse) GoString() string {
  return s.String()
}

func (s *StopLivestreamingResponse) SetMessage(v string) *StopLivestreamingResponse {
  s.Message = &v
  return s
}

func (s *StopLivestreamingResponse) SetCode(v string) *StopLivestreamingResponse {
  s.Code = &v
  return s
}

type StopLivestreamingPaths struct {
}

func (s StopLivestreamingPaths) String() string {
  return tea.Prettify(s)
}

func (s StopLivestreamingPaths) GoString() string {
  return s.String()
}

type StopLivestreamingParameters struct {
}

func (s StopLivestreamingParameters) String() string {
  return tea.Prettify(s)
}

func (s StopLivestreamingParameters) GoString() string {
  return s.String()
}

type StopLivestreamingRequestHeader struct {
}

func (s StopLivestreamingRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s StopLivestreamingRequestHeader) GoString() string {
  return s.String()
}

type StopLivestreamingResponseHeader struct {
}

func (s StopLivestreamingResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s StopLivestreamingResponseHeader) GoString() string {
  return s.String()
}




type QueryForbidLivestreamRecordRequest struct {
  // {'en':'channel', 'zh_CN':'加速域名，如果账号没有域名权限，调用成功但接口会返回错误的提示信息。注：一次只能查询一个域名'}
  Channel *string `json:"channel,omitempty" xml:"channel,omitempty" require:"true"`
  // {'en':'url,Complete push stream or pull stream url', 'zh_CN':'推流或拉流的地址完整的url，即具体某个直播的频道url'}
  Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s QueryForbidLivestreamRecordRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryForbidLivestreamRecordRequest) GoString() string {
  return s.String()
}

func (s *QueryForbidLivestreamRecordRequest) SetChannel(v string) *QueryForbidLivestreamRecordRequest {
  s.Channel = &v
  return s
}

func (s *QueryForbidLivestreamRecordRequest) SetUrl(v string) *QueryForbidLivestreamRecordRequest {
  s.Url = &v
  return s
}

type QueryForbidLivestreamRecordResponse struct {
  // {'en':'result code,00 means success', 'zh_CN':'表示提交结果，00表示成功'}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {'en':'result message', 'zh_CN':'表示任务提交后，系统的响应消息'}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {'en':'Set of task results', 'zh_CN':'任务结果的集合'}
  QueryForbidLivestreamRecordResultDetail []*QueryForbidLivestreamRecordResultDetail `json:"resultDetail,omitempty" xml:"resultDetail,omitempty" require:"true" type:"Repeated"`
}

func (s QueryForbidLivestreamRecordResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryForbidLivestreamRecordResponse) GoString() string {
  return s.String()
}

func (s *QueryForbidLivestreamRecordResponse) SetCode(v string) *QueryForbidLivestreamRecordResponse {
  s.Code = &v
  return s
}

func (s *QueryForbidLivestreamRecordResponse) SetMessage(v string) *QueryForbidLivestreamRecordResponse {
  s.Message = &v
  return s
}

func (s *QueryForbidLivestreamRecordResponse) SetResultDetail(v []*QueryForbidLivestreamRecordResultDetail) *QueryForbidLivestreamRecordResponse {
  s.QueryForbidLivestreamRecordResultDetail = v
  return s
}

type QueryForbidLivestreamRecordResultDetail struct {
  // {'en':'channel', 'zh_CN':'禁止推流或拉流的域名'}
  Channel *string `json:"channel,omitempty" xml:"channel,omitempty" require:"true"`
  // {'en':'url', 'zh_CN':'禁止推流或拉流的地址完整的url'}
  Url *string `json:"url,omitempty" xml:"url,omitempty" require:"true"`
  // {'en':'start time, format:1532413615', 'zh_CN':'开始禁止播放流的时间，格式为 1532413615'}
  StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty" require:"true"`
  // {'en':'end time,format:1532413615', 'zh_CN':'流恢复播放的时间，例如 1532413615'}
  EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty" require:"true"`
  // {'en':'ip of forbidden user', 'zh_CN':'禁止某个客户观看直播的用户IP'}
  Ip *string `json:"ip,omitempty" xml:"ip,omitempty" require:"true"`
}

func (s QueryForbidLivestreamRecordResultDetail) String() string {
  return tea.Prettify(s)
}

func (s QueryForbidLivestreamRecordResultDetail) GoString() string {
  return s.String()
}

func (s *QueryForbidLivestreamRecordResultDetail) SetChannel(v string) *QueryForbidLivestreamRecordResultDetail {
  s.Channel = &v
  return s
}

func (s *QueryForbidLivestreamRecordResultDetail) SetUrl(v string) *QueryForbidLivestreamRecordResultDetail {
  s.Url = &v
  return s
}

func (s *QueryForbidLivestreamRecordResultDetail) SetStartTime(v string) *QueryForbidLivestreamRecordResultDetail {
  s.StartTime = &v
  return s
}

func (s *QueryForbidLivestreamRecordResultDetail) SetEndTime(v string) *QueryForbidLivestreamRecordResultDetail {
  s.EndTime = &v
  return s
}

func (s *QueryForbidLivestreamRecordResultDetail) SetIp(v string) *QueryForbidLivestreamRecordResultDetail {
  s.Ip = &v
  return s
}

type QueryForbidLivestreamRecordPaths struct {
}

func (s QueryForbidLivestreamRecordPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryForbidLivestreamRecordPaths) GoString() string {
  return s.String()
}

type QueryForbidLivestreamRecordParameters struct {
}

func (s QueryForbidLivestreamRecordParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryForbidLivestreamRecordParameters) GoString() string {
  return s.String()
}

type QueryForbidLivestreamRecordRequestHeader struct {
}

func (s QueryForbidLivestreamRecordRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryForbidLivestreamRecordRequestHeader) GoString() string {
  return s.String()
}

type QueryForbidLivestreamRecordResponseHeader struct {
}

func (s QueryForbidLivestreamRecordResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryForbidLivestreamRecordResponseHeader) GoString() string {
  return s.String()
}




type ForbidLivestreamingRequest struct {
  // {"en":"'Only one URL can be submitted at a time, the address of push or pull flow is complete, URL format requirements:
  // 1) The URL must start with'http:/'or'https:/' or'rtmp', and enter an example: http://www.a.com/live/201802221082.
  // 2) The maximum length of each URL is 1000 characters.
  // 3) If a question mark is included in the url, only the URL before the question mark is submitted.
  // 4) Maximum 500 times a day. '", "zh_CN":"一次只能提交一个url，推流或拉流的地址完整的url，url的格式要求：
  // 
  // 1）URL 必须以'http://' 或 'https://' 或&lsquo;rtmp&rsquo;开头，输入示例：http://www.a.com/live/201802221082。
  // 
  // 2）每个url最大长度 1000 字符。
  // 
  // 3）url中如果包含问号，则只需提交问号前面的url。
  // 
  // 4）每天最大500次，禁止，恢复，停止3个操作共用上限。"}
  LiveUrl *string `json:"liveUrl,omitempty" xml:"liveUrl,omitempty" require:"true"`
  // {"en":"Relative time, the time interval in which the operation takes effect, is calculated from the current time.
  // Unit: seconds
  // For example: reltime = 300, from the current time to ban 300 seconds
  // Note:
  // 1) When both abstime_end and reltime parameters are passed, only reltime parameters are valid.
  // 2) When reltime, abstime_end and reltime parameters are not set, the default one-day ban is applied.
  // 3) When reltime is less than or equal to 0, it is set to 1 day.
  // 4) When reltime is greater than 30 days, it is set to 30 days.'", "zh_CN":"相对时间，表示操作生效的时间区间，从当前时间开始计算。
  // 单位为：秒
  // 如：reltime=300，从当前时间开始禁播300s
  // 
  // 注：
  // 
  // 1)abstime_end和reltime参数都传时，仅reltime参数有效
  // 
  // 2)reltime、abstime_end和reltime参数都没有设置时，默认禁播1天
  // 
  // 3)reltime小于等于0时，会设置为1天
  // 
  // 4)reltime大于30天时，会设置为30天"}
  Reltime *string `json:"reltime,omitempty" xml:"reltime,omitempty"`
  // {'en':'Start time, only support current time,not future time; use with abstime_end parameter in seconds;
  // The time format is yyyymmddHMMSS or 10-digit UNIX timestamp, such as 20140306121035, 1511256503.
  // Abstime_start < current time or abstime_start is empty, default starts at current time.
  // Note:
  // 1) When both abstime_end and reltime parameters are passed, only reltime parameters are valid.
  // 2) When reltime, abstime_end and reltime parameters are not set, the default one-day ban is applied.
  // 3) abstime_start is larger than the current time of 30 days and will report errors.', 'zh_CN':'开始时间，只支持从当前时间开始，不支持未来时间；与abstime_end参数配合使用，单位为：秒； 
  // 时间格式为：yyyymmddHHMMSS或10进制unix时间戳，如：20140306121035、1511256503
  // abstime_start < 当前时间 或 abstime_start 为空 时，默认从当前时间开始。
  // 
  // 注：
  // 
  // 1)abstime_end和reltime参数都传时，仅reltime参数有效
  // 
  // 2)reltime、abstime_end和reltime参数都没有设置时，默认禁播1天
  // 
  // 3)abstime_start 大于当前时间30天，会报错'}
  AbstimeStart *string `json:"abstimeStart,omitempty" xml:"abstimeStart,omitempty"`
  // {'en':'End time, specify the end time point of the operation, and use it with abstime_start parameter in seconds.
  // The time format is yyyymmddHMMSS or 10-digit UNIX timestamp, such as 20140306121535, 1511256503.
  // When abstime_end < current time, the parameter is invalid.
  // Note:
  // 1) When both abstime_end and reltime parameters are passed, only reltime parameters are valid.
  // 2) When reltime, abstime_end and reltime parameters are not set, the default one-day ban is applied.
  // 3) When abstime_end is free, if abstime_start is empty, the duration of forbidden sowing is one day longer than the current time; if abstime_start is valuable, the duration of forbidden sowing is one day longer than abstime_start.
  // 4) When abstime_end is less than the current time, an error is reported.
  // 5) When abstime_end is greater than the current time of 30 days, the duration of banning broadcasting is set as: current time + 30 days.', 'zh_CN':'结束时间，指定操作生效结束时间点；与abstime_start参数配合使用，单位为：秒；
  // 时间格式为：yyyymmddHHMMSS或10进制unix时间戳，如：20140306121535、1511256503
  // abstime_end < 当前时间 时，参数无效。
  // 
  // 注：
  // 
  // 1)abstime_end和reltime参数都传时，仅reltime参数有效
  // 2)reltime、abstime_end和reltime参数都没有设置时，默认禁播1天
  // 3)当abstime_end有空时，如果abstime_start为空，则禁播时长是当前时间加1天；若abstime_start有值，则禁播时长是abstime_start加1天
  // 4)当abstime_end小于当前时间，报错。
  // 5)当abstime_end 大于当前时间30天，则禁播时长会设置为：当前时间+30天'}
  AbstimeEnd *string `json:"abstimeEnd,omitempty" xml:"abstimeEnd,omitempty"`
  // {'en':'The values are: play (back-source pull), publish (live push).', 'zh_CN':'值为：play (回源拉流)，publish(直播推流)。'}
  Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ForbidLivestreamingRequest) String() string {
  return tea.Prettify(s)
}

func (s ForbidLivestreamingRequest) GoString() string {
  return s.String()
}

func (s *ForbidLivestreamingRequest) SetLiveUrl(v string) *ForbidLivestreamingRequest {
  s.LiveUrl = &v
  return s
}

func (s *ForbidLivestreamingRequest) SetReltime(v string) *ForbidLivestreamingRequest {
  s.Reltime = &v
  return s
}

func (s *ForbidLivestreamingRequest) SetAbstimeStart(v string) *ForbidLivestreamingRequest {
  s.AbstimeStart = &v
  return s
}

func (s *ForbidLivestreamingRequest) SetAbstimeEnd(v string) *ForbidLivestreamingRequest {
  s.AbstimeEnd = &v
  return s
}

func (s *ForbidLivestreamingRequest) SetType(v string) *ForbidLivestreamingRequest {
  s.Type = &v
  return s
}

type ForbidLivestreamingResponse struct {
  // {'en':'message', 'zh_CN':'错误信息'}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {'en':'code', 'zh_CN':'错误码'}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
}

func (s ForbidLivestreamingResponse) String() string {
  return tea.Prettify(s)
}

func (s ForbidLivestreamingResponse) GoString() string {
  return s.String()
}

func (s *ForbidLivestreamingResponse) SetMessage(v string) *ForbidLivestreamingResponse {
  s.Message = &v
  return s
}

func (s *ForbidLivestreamingResponse) SetCode(v string) *ForbidLivestreamingResponse {
  s.Code = &v
  return s
}

type ForbidLivestreamingPaths struct {
}

func (s ForbidLivestreamingPaths) String() string {
  return tea.Prettify(s)
}

func (s ForbidLivestreamingPaths) GoString() string {
  return s.String()
}

type ForbidLivestreamingParameters struct {
}

func (s ForbidLivestreamingParameters) String() string {
  return tea.Prettify(s)
}

func (s ForbidLivestreamingParameters) GoString() string {
  return s.String()
}

type ForbidLivestreamingRequestHeader struct {
}

func (s ForbidLivestreamingRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ForbidLivestreamingRequestHeader) GoString() string {
  return s.String()
}

type ForbidLivestreamingResponseHeader struct {
}

func (s ForbidLivestreamingResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ForbidLivestreamingResponseHeader) GoString() string {
  return s.String()
}




type ResumeLivestreamingRequest struct {
  // {"en":"Push  or pull  address complete url, URL format requirements:
  // 1) The URL must start with'http:/'or'https:/' or'rtmp', and enter an example: http://www.a.com/live/201802221082.
  // 2) The maximum length of each URL is 1000 characters.
  // 3) The domain name of each URL must be the domain name of our company which is speeded up by live broadcasting.
  // 4) If a question mark is included in the url, only the URL before the question mark is submitted.", "zh_CN":"推流或拉流的地址完整的url，url的格式要求：
  // 1）URL 必须以'http://' 或 'https://' 或&lsquo;rtmp&rsquo;开头，输入示例：http://www.a.com/live/201802221082。
  // 2）每个url最大长度 1000 字符。
  // 3）每个url所在的域名必须是在我司直播加速的域名。
  // 4）url中如果包含问号，则只需提交问号前面的url。"}
  LiveUrl *string `json:"liveUrl,omitempty" xml:"liveUrl,omitempty" require:"true"`
  // {"en":"The values are: play (back-source pull), publish (live push).", "zh_CN":"值为：play (回源拉流)，publish(直播推流)。"}
  Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ResumeLivestreamingRequest) String() string {
  return tea.Prettify(s)
}

func (s ResumeLivestreamingRequest) GoString() string {
  return s.String()
}

func (s *ResumeLivestreamingRequest) SetLiveUrl(v string) *ResumeLivestreamingRequest {
  s.LiveUrl = &v
  return s
}

func (s *ResumeLivestreamingRequest) SetType(v string) *ResumeLivestreamingRequest {
  s.Type = &v
  return s
}

type ResumeLivestreamingResponse struct {
  // {'en':'message', 'zh_CN':'错误信息'}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {'en':'error code', 'zh_CN':'错误码'}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
}

func (s ResumeLivestreamingResponse) String() string {
  return tea.Prettify(s)
}

func (s ResumeLivestreamingResponse) GoString() string {
  return s.String()
}

func (s *ResumeLivestreamingResponse) SetMessage(v string) *ResumeLivestreamingResponse {
  s.Message = &v
  return s
}

func (s *ResumeLivestreamingResponse) SetCode(v string) *ResumeLivestreamingResponse {
  s.Code = &v
  return s
}

type ResumeLivestreamingPaths struct {
}

func (s ResumeLivestreamingPaths) String() string {
  return tea.Prettify(s)
}

func (s ResumeLivestreamingPaths) GoString() string {
  return s.String()
}

type ResumeLivestreamingParameters struct {
}

func (s ResumeLivestreamingParameters) String() string {
  return tea.Prettify(s)
}

func (s ResumeLivestreamingParameters) GoString() string {
  return s.String()
}

type ResumeLivestreamingRequestHeader struct {
}

func (s ResumeLivestreamingRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ResumeLivestreamingRequestHeader) GoString() string {
  return s.String()
}

type ResumeLivestreamingResponseHeader struct {
}

func (s ResumeLivestreamingResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ResumeLivestreamingResponseHeader) GoString() string {
  return s.String()
}




