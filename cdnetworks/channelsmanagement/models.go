package channelsmanagement

// This file is auto-generated, don't edit it. Thanks.
import (
  "github.com/alibabacloud-go/tea/tea"
)

type GetChannelShareUrlRequest struct {
  // {"en":"Channel push ID", "zh_CN":"频道推流 ID"}
  PullId *string `json:"pullId,omitempty" xml:"pullId,omitempty"`
}

func (s GetChannelShareUrlRequest) String() string {
  return tea.Prettify(s)
}

func (s GetChannelShareUrlRequest) GoString() string {
  return s.String()
}

func (s *GetChannelShareUrlRequest) SetPullId(v string) *GetChannelShareUrlRequest {
  s.PullId = &v
  return s
}

type GetChannelShareUrlResponse struct {
  // {"en":"200 success", "zh_CN":"200，操作成功"}
  Code *int32 `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Successful operation", "zh_CN":"操作成功"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Return data", "zh_CN":"分享页url"}
  Data *string `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s GetChannelShareUrlResponse) String() string {
  return tea.Prettify(s)
}

func (s GetChannelShareUrlResponse) GoString() string {
  return s.String()
}

func (s *GetChannelShareUrlResponse) SetCode(v int32) *GetChannelShareUrlResponse {
  s.Code = &v
  return s
}

func (s *GetChannelShareUrlResponse) SetMessage(v string) *GetChannelShareUrlResponse {
  s.Message = &v
  return s
}

func (s *GetChannelShareUrlResponse) SetData(v string) *GetChannelShareUrlResponse {
  s.Data = &v
  return s
}

type GetChannelShareUrlPaths struct {
}

func (s GetChannelShareUrlPaths) String() string {
  return tea.Prettify(s)
}

func (s GetChannelShareUrlPaths) GoString() string {
  return s.String()
}

type GetChannelShareUrlParameters struct {
}

func (s GetChannelShareUrlParameters) String() string {
  return tea.Prettify(s)
}

func (s GetChannelShareUrlParameters) GoString() string {
  return s.String()
}

type GetChannelShareUrlRequestHeader struct {
}

func (s GetChannelShareUrlRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetChannelShareUrlRequestHeader) GoString() string {
  return s.String()
}

type GetChannelShareUrlResponseHeader struct {
}

func (s GetChannelShareUrlResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetChannelShareUrlResponseHeader) GoString() string {
  return s.String()
}




type GetChannelListRequest struct {
  // {"en":"Gets the sorting of the channel list, default is 0;
  // 0. In descending order by creation time; 1. In ascending order by creation time", "zh_CN":"获取频道列表的排序，默认为 0； 
  // 0、按创建时间降序，1、按创建时间升序 "}
  ListOrder *int32 `json:"listOrder,omitempty" xml:"listOrder,omitempty"`
  // {"en":"The number of pages in the channel page list, starting from 1; Default 1", "zh_CN":"频道分页列表中第几页，从 1 开始取值；默认 1"}
  PageIndex *int32 `json:"pageIndex,omitempty" xml:"pageIndex,omitempty"`
  // {"en":"Average number of channels per page. The default value is 10. The value ranges from 1 to 50", "zh_CN":"平均每页频道数量，默认为 10，取值在 1-50 之间"}
  PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
  // {"en":"Application scenario", "zh_CN":"应用场景 "}
  AppScenario *string `json:"appScenario,omitempty" xml:"appScenario,omitempty"`
}

func (s GetChannelListRequest) String() string {
  return tea.Prettify(s)
}

func (s GetChannelListRequest) GoString() string {
  return s.String()
}

func (s *GetChannelListRequest) SetListOrder(v int32) *GetChannelListRequest {
  s.ListOrder = &v
  return s
}

func (s *GetChannelListRequest) SetPageIndex(v int32) *GetChannelListRequest {
  s.PageIndex = &v
  return s
}

func (s *GetChannelListRequest) SetPageSize(v int32) *GetChannelListRequest {
  s.PageSize = &v
  return s
}

func (s *GetChannelListRequest) SetAppScenario(v string) *GetChannelListRequest {
  s.AppScenario = &v
  return s
}

type GetChannelListResponse struct {
  // {'en':'code', 'zh_CN':'返回状态码'}
  Code *int32 `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  Data *GetChannelListDataList `json:"data,omitempty" xml:"data,omitempty" require:"true"`
  // {'en':'message', 'zh_CN':'返回消息'}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s GetChannelListResponse) String() string {
  return tea.Prettify(s)
}

func (s GetChannelListResponse) GoString() string {
  return s.String()
}

func (s *GetChannelListResponse) SetCode(v int32) *GetChannelListResponse {
  s.Code = &v
  return s
}

func (s *GetChannelListResponse) SetData(v *GetChannelListDataList) *GetChannelListResponse {
  s.Data = v
  return s
}

func (s *GetChannelListResponse) SetMessage(v string) *GetChannelListResponse {
  s.Message = &v
  return s
}

type GetChannelListDataList struct {
  GetChannelListChannelList []*GetChannelListChannelList `json:"channelList,omitempty" xml:"channelList,omitempty" require:"true" type:"Repeated"`
  // {'en':'The number of records currently returned for the channel list information. Note that the number of records returned here is only the number of records for the current page.', 'zh_CN':'当前返回的频道列表信息的记录数，注意这里返回的记录数只是当前页的记录数。'}
  ChannelToal *int32 `json:"channelToal,omitempty" xml:"channelToal,omitempty" require:"true"`
}

func (s GetChannelListDataList) String() string {
  return tea.Prettify(s)
}

func (s GetChannelListDataList) GoString() string {
  return s.String()
}

func (s *GetChannelListDataList) SetChannelList(v []*GetChannelListChannelList) *GetChannelListDataList {
  s.GetChannelListChannelList = v
  return s
}

func (s *GetChannelListDataList) SetChannelToal(v int32) *GetChannelListDataList {
  s.ChannelToal = &v
  return s
}

type GetChannelListChannelList struct {
  // {"en":"Application scenario", "zh_CN":"应用场景"}
  AppScenario *string `json:"appScenario,omitempty" xml:"appScenario,omitempty" require:"true"`
  // {"en":"Channel name", "zh_CN":"频道名称"}
  ChannelName *string `json:"channelName,omitempty" xml:"channelName,omitempty" require:"true"`
  // {"en":"Channel current live status;
  // 1:live broadcast 2:not broadcast 3:banned", "zh_CN":"频道当前直播状态； 
  // 1、直播中   2、未开播  3、禁播 "}
  ChannelStatus *int32 `json:"channelStatus,omitempty" xml:"channelStatus,omitempty" require:"true"`
  // {"en":"Channel type", "zh_CN":"频道类型"}
  ChannelType *string `json:"channelType,omitempty" xml:"channelType,omitempty" require:"true"`
  // {"en":"Channel creation time, for example, 2016-05-08 12:00:00", "zh_CN":"频道创建时间，例如：2016-05-08 12:00:00"}
  CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty" require:"true"`
  // {"en":"Recorded directory ID. The value is returned only when the value-added service of recorded directory management is enabled and the recorded directory is configured for the channel.", "zh_CN":"录制目录 ID，有开通录制目录管理增值服务且有给频道配置过录制目录才会有返回值。"}
  DirectoryId *string `json:"directoryId,omitempty" xml:"directoryId,omitempty" require:"true"`
  // {"en":"If it is enabled, push it and record it. If 1 is enabled, 0 is disabled.", "zh_CN":"是否开启即推即录，1 开启，0 未开启。"}
  Live2vod *string `json:"live2vod,omitempty" xml:"live2vod,omitempty" require:"true"`
  // {"en":"Pull flow ID", "zh_CN":"拉流 ID"}
  PullId *string `json:"pullId,omitempty" xml:"pullId,omitempty" require:"true"`
  // {"en":"Push URL", "zh_CN":"推流 URL"}
  PushUrl *string `json:"pushUrl,omitempty" xml:"pushUrl,omitempty" require:"true"`
  // {"en":"The URL is in SRT format", "zh_CN":"SRT 格式的推流 URL"}
  SrtPushUrl *string `json:"srtPushUrl,omitempty" xml:"srtPushUrl,omitempty" require:"true"`
  // {"en":"User ID for creating a channel", "zh_CN":"创建频道的用户 ID"}
  UserId *string `json:"userId,omitempty" xml:"userId,omitempty" require:"true"`
}

func (s GetChannelListChannelList) String() string {
  return tea.Prettify(s)
}

func (s GetChannelListChannelList) GoString() string {
  return s.String()
}

func (s *GetChannelListChannelList) SetAppScenario(v string) *GetChannelListChannelList {
  s.AppScenario = &v
  return s
}

func (s *GetChannelListChannelList) SetChannelName(v string) *GetChannelListChannelList {
  s.ChannelName = &v
  return s
}

func (s *GetChannelListChannelList) SetChannelStatus(v int32) *GetChannelListChannelList {
  s.ChannelStatus = &v
  return s
}

func (s *GetChannelListChannelList) SetChannelType(v string) *GetChannelListChannelList {
  s.ChannelType = &v
  return s
}

func (s *GetChannelListChannelList) SetCreateTime(v string) *GetChannelListChannelList {
  s.CreateTime = &v
  return s
}

func (s *GetChannelListChannelList) SetDirectoryId(v string) *GetChannelListChannelList {
  s.DirectoryId = &v
  return s
}

func (s *GetChannelListChannelList) SetLive2vod(v string) *GetChannelListChannelList {
  s.Live2vod = &v
  return s
}

func (s *GetChannelListChannelList) SetPullId(v string) *GetChannelListChannelList {
  s.PullId = &v
  return s
}

func (s *GetChannelListChannelList) SetPushUrl(v string) *GetChannelListChannelList {
  s.PushUrl = &v
  return s
}

func (s *GetChannelListChannelList) SetSrtPushUrl(v string) *GetChannelListChannelList {
  s.SrtPushUrl = &v
  return s
}

func (s *GetChannelListChannelList) SetUserId(v string) *GetChannelListChannelList {
  s.UserId = &v
  return s
}

type GetChannelListPaths struct {
}

func (s GetChannelListPaths) String() string {
  return tea.Prettify(s)
}

func (s GetChannelListPaths) GoString() string {
  return s.String()
}

type GetChannelListParameters struct {
}

func (s GetChannelListParameters) String() string {
  return tea.Prettify(s)
}

func (s GetChannelListParameters) GoString() string {
  return s.String()
}

type GetChannelListRequestHeader struct {
}

func (s GetChannelListRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetChannelListRequestHeader) GoString() string {
  return s.String()
}

type GetChannelListResponseHeader struct {
}

func (s GetChannelListResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetChannelListResponseHeader) GoString() string {
  return s.String()
}




type GetChannelConfRequest struct {
  // {"en":"Pull stream id, unique ID of the channel", "zh_CN":"拉流 id，频道的唯一 ID"}
  PullId *string `json:"pullId,omitempty" xml:"pullId,omitempty" require:"true"`
}

func (s GetChannelConfRequest) String() string {
  return tea.Prettify(s)
}

func (s GetChannelConfRequest) GoString() string {
  return s.String()
}

func (s *GetChannelConfRequest) SetPullId(v string) *GetChannelConfRequest {
  s.PullId = &v
  return s
}

type GetChannelConfResponse struct {
  // {'en':'Status code', 'zh_CN':'返回状态码'}
  Code *int32 `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  Data *GetChannelConfDataList `json:"data,omitempty" xml:"data,omitempty" require:"true"`
  // {'en':'Operational information', 'zh_CN':'操作信息'}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s GetChannelConfResponse) String() string {
  return tea.Prettify(s)
}

func (s GetChannelConfResponse) GoString() string {
  return s.String()
}

func (s *GetChannelConfResponse) SetCode(v int32) *GetChannelConfResponse {
  s.Code = &v
  return s
}

func (s *GetChannelConfResponse) SetData(v *GetChannelConfDataList) *GetChannelConfResponse {
  s.Data = v
  return s
}

func (s *GetChannelConfResponse) SetMessage(v string) *GetChannelConfResponse {
  s.Message = &v
  return s
}

type GetChannelConfDataList struct {
  // {"en":"Application scenario", "zh_CN":"应用场景"}
  AppScenario *string `json:"appScenario,omitempty" xml:"appScenario,omitempty" require:"true"`
  // {"en":"Channel name", "zh_CN":"频道名称"}
  ChannelName *string `json:"channelName,omitempty" xml:"channelName,omitempty" require:"true"`
  // {"en":"Live channel current status;
  // 1:live broadcast 2:not broadcast 3:banned", "zh_CN":"频道当前直播状态； 
  // 1、直播中   2、未开播  3、禁播 "}
  ChannelStatus *int32 `json:"channelStatus,omitempty" xml:"channelStatus,omitempty" require:"true"`
  // {"en":"Channel creation time, for example, 2016-05-08 12:00:00", "zh_CN":"频道创建时间，例如：2016-05-08 12:00:00"}
  CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty" require:"true"`
  // {"en":"The user name who created the channel", "zh_CN":"创建频道的 userName"}
  CreateUser *string `json:"createUser,omitempty" xml:"createUser,omitempty" require:"true"`
  // {"en":"Recorded directory ID. The value is returned only when the value-added service of recorded directory management is enabled and the recorded directory is configured for the channel.", "zh_CN":"录制目录 ID，有开通录制目录管理增值服务且有给频道配置过录制目录才会有返回值。"}
  DirectoryId *string `json:"directoryId,omitempty" xml:"directoryId,omitempty" require:"true"`
  // {"en":"If it is enabled, push it and record it. If 1 is enabled, 0 is disabled.", "zh_CN":"是否开启即推即录，1 开启，0 未开启。"}
  Live2vod *string `json:"live2vod,omitempty" xml:"live2vod,omitempty" require:"true"`
  // {"en":"The flash player ID used by the channel", "zh_CN":"频道使用的 flash 播放器 ID "}
  PlayerId *string `json:"playerId,omitempty" xml:"playerId,omitempty" require:"true"`
  // {"en":"Video stream encryption. 0: No encryption is required 1:HLS universal encryption 2:DRM encryption", "zh_CN":"视频流加密。0:不做任何加密 1:HLS通用加密 2:DRM加密"}
  VideoEncrypted *int32 `json:"videoEncrypted,omitempty" xml:"videoEncrypted,omitempty" require:"true"`
  // {"en":"HLS universal encryption validity time (unit: second)", "zh_CN":"HLS通用加密有效时间（单位：秒）"}
  HlsEncryptDuration *int64 `json:"hlsEncryptDuration,omitempty" xml:"hlsEncryptDuration,omitempty" require:"true"`
  GetChannelConfPullConfigureList []*GetChannelConfPullConfigureList `json:"pullConfigureList,omitempty" xml:"pullConfigureList,omitempty" require:"true" type:"Repeated"`
  GetChannelConfPushConfigureInfo []*GetChannelConfPushConfigureInfo `json:"pushConfigureInfo,omitempty" xml:"pushConfigureInfo,omitempty" require:"true" type:"Repeated"`
  GetChannelConfSrtPullConfigureList []*GetChannelConfSrtPullConfigureList `json:"srtPullConfigureList,omitempty" xml:"srtPullConfigureList,omitempty" require:"true" type:"Repeated"`
}

func (s GetChannelConfDataList) String() string {
  return tea.Prettify(s)
}

func (s GetChannelConfDataList) GoString() string {
  return s.String()
}

func (s *GetChannelConfDataList) SetAppScenario(v string) *GetChannelConfDataList {
  s.AppScenario = &v
  return s
}

func (s *GetChannelConfDataList) SetChannelName(v string) *GetChannelConfDataList {
  s.ChannelName = &v
  return s
}

func (s *GetChannelConfDataList) SetChannelStatus(v int32) *GetChannelConfDataList {
  s.ChannelStatus = &v
  return s
}

func (s *GetChannelConfDataList) SetCreateTime(v string) *GetChannelConfDataList {
  s.CreateTime = &v
  return s
}

func (s *GetChannelConfDataList) SetCreateUser(v string) *GetChannelConfDataList {
  s.CreateUser = &v
  return s
}

func (s *GetChannelConfDataList) SetDirectoryId(v string) *GetChannelConfDataList {
  s.DirectoryId = &v
  return s
}

func (s *GetChannelConfDataList) SetLive2vod(v string) *GetChannelConfDataList {
  s.Live2vod = &v
  return s
}

func (s *GetChannelConfDataList) SetPlayerId(v string) *GetChannelConfDataList {
  s.PlayerId = &v
  return s
}

func (s *GetChannelConfDataList) SetVideoEncrypted(v int32) *GetChannelConfDataList {
  s.VideoEncrypted = &v
  return s
}

func (s *GetChannelConfDataList) SetHlsEncryptDuration(v int64) *GetChannelConfDataList {
  s.HlsEncryptDuration = &v
  return s
}

func (s *GetChannelConfDataList) SetPullConfigureList(v []*GetChannelConfPullConfigureList) *GetChannelConfDataList {
  s.GetChannelConfPullConfigureList = v
  return s
}

func (s *GetChannelConfDataList) SetPushConfigureInfo(v []*GetChannelConfPushConfigureInfo) *GetChannelConfDataList {
  s.GetChannelConfPushConfigureInfo = v
  return s
}

func (s *GetChannelConfDataList) SetSrtPullConfigureList(v []*GetChannelConfSrtPullConfigureList) *GetChannelConfDataList {
  s.GetChannelConfSrtPullConfigureList = v
  return s
}

type GetChannelConfPullConfigureList struct {
  // {"en":"Pull the live stream domain name", "zh_CN":"拉流域名"}
  PullDomain *string `json:"pullDomain,omitempty" xml:"pullDomain,omitempty" require:"true"`
  // {"en":"Pull-flow protocol", "zh_CN":"拉流协议"}
  PullProtocol *string `json:"pullProtocol,omitempty" xml:"pullProtocol,omitempty" require:"true"`
  GetChannelConfPullUrlList []*GetChannelConfPullUrlList `json:"pullUrlList,omitempty" xml:"pullUrlList,omitempty" require:"true" type:"Repeated"`
}

func (s GetChannelConfPullConfigureList) String() string {
  return tea.Prettify(s)
}

func (s GetChannelConfPullConfigureList) GoString() string {
  return s.String()
}

func (s *GetChannelConfPullConfigureList) SetPullDomain(v string) *GetChannelConfPullConfigureList {
  s.PullDomain = &v
  return s
}

func (s *GetChannelConfPullConfigureList) SetPullProtocol(v string) *GetChannelConfPullConfigureList {
  s.PullProtocol = &v
  return s
}

func (s *GetChannelConfPullConfigureList) SetPullUrlList(v []*GetChannelConfPullUrlList) *GetChannelConfPullConfigureList {
  s.GetChannelConfPullUrlList = v
  return s
}

type GetChannelConfPullUrlList struct {
  // {"en":"Audio Url", "zh_CN":"音频Url"}
  AudioUrl *string `json:"audioUrl,omitempty" xml:"audioUrl,omitempty" require:"true"`
  // {"en":"Smooth bit rate pull stream url", "zh_CN":"流畅码率拉流 url"}
  FluentPullUrl *string `json:"fluentPullUrl,omitempty" xml:"fluentPullUrl,omitempty" require:"true"`
  // {"en":"Smooth bit rate (smart HD) pull stream url", "zh_CN":"流畅码率（智控高清）拉流 url"}
  FluentZkgqPullUrl *string `json:"fluentZkgqPullUrl,omitempty" xml:"fluentZkgqPullUrl,omitempty" require:"true"`
  // {"en":"Ultra clear bit rate pull url", "zh_CN":"超清码率拉流 url "}
  HdPullUrl *string `json:"hdPullUrl,omitempty" xml:"hdPullUrl,omitempty" require:"true"`
  // {"en":"Ultra - clear bit rate (intelligent control HD) pull - stream url", "zh_CN":"超清码率（智控高清）拉流 url"}
  HdZkgqPullUrl *string `json:"hdZkgqPullUrl,omitempty" xml:"hdZkgqPullUrl,omitempty" require:"true"`
  // {"en":"Hd bit rate pull stream url", "zh_CN":"高清码率拉流 url"}
  HighPullUrl *string `json:"highPullUrl,omitempty" xml:"highPullUrl,omitempty" require:"true"`
  // {"en":"Hd bit rate (intelligent control HD) pull url", "zh_CN":"高清码率（智控高清）拉流 url"}
  HighZkgqPullUrl *string `json:"highZkgqPullUrl,omitempty" xml:"highZkgqPullUrl,omitempty" require:"true"`
  // {"en":"Source rate pull url", "zh_CN":"源码率拉流 url"}
  OriginPullUrl *string `json:"originPullUrl,omitempty" xml:"originPullUrl,omitempty" require:"true"`
  // {"en":"Source rate (intelligent control HD) pull url", "zh_CN":"源码率（智控高清）拉流 url"}
  OriginZkgqPullUrl *string `json:"originZkgqPullUrl,omitempty" xml:"originZkgqPullUrl,omitempty" require:"true"`
  // {"en":"Standard definition bit rate pull url", "zh_CN":"标清码率拉流 url"}
  SdPullUrl *string `json:"sdPullUrl,omitempty" xml:"sdPullUrl,omitempty" require:"true"`
  // {"en":"Standard definition code rate (intelligent control HD) pull url", "zh_CN":"标清码率（智控高清）拉流 url"}
  SdZkgqPullUrl *string `json:"sdZkgqPullUrl,omitempty" xml:"sdZkgqPullUrl,omitempty" require:"true"`
}

func (s GetChannelConfPullUrlList) String() string {
  return tea.Prettify(s)
}

func (s GetChannelConfPullUrlList) GoString() string {
  return s.String()
}

func (s *GetChannelConfPullUrlList) SetAudioUrl(v string) *GetChannelConfPullUrlList {
  s.AudioUrl = &v
  return s
}

func (s *GetChannelConfPullUrlList) SetFluentPullUrl(v string) *GetChannelConfPullUrlList {
  s.FluentPullUrl = &v
  return s
}

func (s *GetChannelConfPullUrlList) SetFluentZkgqPullUrl(v string) *GetChannelConfPullUrlList {
  s.FluentZkgqPullUrl = &v
  return s
}

func (s *GetChannelConfPullUrlList) SetHdPullUrl(v string) *GetChannelConfPullUrlList {
  s.HdPullUrl = &v
  return s
}

func (s *GetChannelConfPullUrlList) SetHdZkgqPullUrl(v string) *GetChannelConfPullUrlList {
  s.HdZkgqPullUrl = &v
  return s
}

func (s *GetChannelConfPullUrlList) SetHighPullUrl(v string) *GetChannelConfPullUrlList {
  s.HighPullUrl = &v
  return s
}

func (s *GetChannelConfPullUrlList) SetHighZkgqPullUrl(v string) *GetChannelConfPullUrlList {
  s.HighZkgqPullUrl = &v
  return s
}

func (s *GetChannelConfPullUrlList) SetOriginPullUrl(v string) *GetChannelConfPullUrlList {
  s.OriginPullUrl = &v
  return s
}

func (s *GetChannelConfPullUrlList) SetOriginZkgqPullUrl(v string) *GetChannelConfPullUrlList {
  s.OriginZkgqPullUrl = &v
  return s
}

func (s *GetChannelConfPullUrlList) SetSdPullUrl(v string) *GetChannelConfPullUrlList {
  s.SdPullUrl = &v
  return s
}

func (s *GetChannelConfPullUrlList) SetSdZkgqPullUrl(v string) *GetChannelConfPullUrlList {
  s.SdZkgqPullUrl = &v
  return s
}

type GetChannelConfPushConfigureInfo struct {
  // {"en":"Push-flow protocol", "zh_CN":"推流协议"}
  PushProtocol *string `json:"pushProtocol,omitempty" xml:"pushProtocol,omitempty" require:"true"`
  // {"en":"Push URL", "zh_CN":"推流 URL"}
  PushUrl *string `json:"pushUrl,omitempty" xml:"pushUrl,omitempty" require:"true"`
}

func (s GetChannelConfPushConfigureInfo) String() string {
  return tea.Prettify(s)
}

func (s GetChannelConfPushConfigureInfo) GoString() string {
  return s.String()
}

func (s *GetChannelConfPushConfigureInfo) SetPushProtocol(v string) *GetChannelConfPushConfigureInfo {
  s.PushProtocol = &v
  return s
}

func (s *GetChannelConfPushConfigureInfo) SetPushUrl(v string) *GetChannelConfPushConfigureInfo {
  s.PushUrl = &v
  return s
}

type GetChannelConfSrtPullConfigureList struct {
  // {"en":"Pull the live stream domain name", "zh_CN":"拉流域名"}
  PullDomain *string `json:"pullDomain,omitempty" xml:"pullDomain,omitempty" require:"true"`
  // {"en":"Pull-flow protocol", "zh_CN":"拉流协议"}
  PullProtocol *string `json:"pullProtocol,omitempty" xml:"pullProtocol,omitempty" require:"true"`
  GetChannelConfPullUrlList []*GetChannelConfPullUrlList `json:"pullUrlList,omitempty" xml:"pullUrlList,omitempty" require:"true" type:"Repeated"`
}

func (s GetChannelConfSrtPullConfigureList) String() string {
  return tea.Prettify(s)
}

func (s GetChannelConfSrtPullConfigureList) GoString() string {
  return s.String()
}

func (s *GetChannelConfSrtPullConfigureList) SetPullDomain(v string) *GetChannelConfSrtPullConfigureList {
  s.PullDomain = &v
  return s
}

func (s *GetChannelConfSrtPullConfigureList) SetPullProtocol(v string) *GetChannelConfSrtPullConfigureList {
  s.PullProtocol = &v
  return s
}

func (s *GetChannelConfSrtPullConfigureList) SetPullUrlList(v []*GetChannelConfPullUrlList) *GetChannelConfSrtPullConfigureList {
  s.GetChannelConfPullUrlList = v
  return s
}

type GetChannelConfPaths struct {
}

func (s GetChannelConfPaths) String() string {
  return tea.Prettify(s)
}

func (s GetChannelConfPaths) GoString() string {
  return s.String()
}

type GetChannelConfParameters struct {
}

func (s GetChannelConfParameters) String() string {
  return tea.Prettify(s)
}

func (s GetChannelConfParameters) GoString() string {
  return s.String()
}

type GetChannelConfRequestHeader struct {
}

func (s GetChannelConfRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetChannelConfRequestHeader) GoString() string {
  return s.String()
}

type GetChannelConfResponseHeader struct {
}

func (s GetChannelConfResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetChannelConfResponseHeader) GoString() string {
  return s.String()
}




type CreateRecordTaskRequest struct {
  // {"en":"Channel pull ID. If an unstarted, started task already exists for the pullId, you cannot add another task for the same pullId.", "zh_CN":"频道拉流 ID。如果该 pullId 已存在未开始、已开始的任务，则不能再添加同一个 pullId 的任务。"}
  PullId *string `json:"pullId,omitempty" xml:"pullId,omitempty" require:"true"`
  // {"en":"The service ID must be unique. You are advised to use a 32-bit UUID and the value can be a string of up to 32 characters", "zh_CN":"业务 ID，需用户自己控制唯一性，建议使用 32位 UUID，并且最长为 32 位字符串"}
  TransNo *string `json:"transNo,omitempty" xml:"transNo,omitempty"`
  // {"en":"Whether to start recording immediately: true: start recording immediately; false: start recording periodically. When true, startRecordTime must not be passed or be an empty string.", "zh_CN":"是否立即启动录制，true 表示立即启动，false为定时启动。当为 true 时，startRecordTime必须不传或为空字符串。"}
  QuickStart *bool `json:"quickStart,omitempty" xml:"quickStart,omitempty" require:"true"`
  // {"en":"Schedule a recording start time. The format is a timestamp in seconds, for example, 1502438820. Actual participation in the specific program calculation, will automatically remove the number of extra seconds accurate to minutes, the actual accuracy is only minutes. The recording start time accurate to one minute later must be later than the current time. This parameter is mandatory when quickStart is false.", "zh_CN":"计划录制开始时间。格式为秒级时间戳，如1502438820。实际参与到具体程序运算中，会自动去除多余秒数精确到分钟，实际精度只到分钟。精确到分钟后的录制开始时间必须是大于当前时间。当 quickStart 为 false 时必填。"}
  StartTime *int32 `json:"startTime,omitempty" xml:"startTime,omitempty"`
  // {"en":"Scheduled recording end time. This parameter is mandatory. The format is a timestamp in seconds, for example, 1502438880. Actual participation in the specific program calculation, will automatically remove the number of extra seconds accurate to minutes, the actual accuracy is only minutes. The recording end time accurate to minute must be later than the current time and later than the start time accurate to minute. The interval between the end time and the start time, which is accurate to minutes by seconds, cannot exceed 72 hours. There will be a small error between the recorded video duration and the planned recording time", "zh_CN":"计划录制结束时间，必填。格式为秒级时间戳，如 1502438880。 实际参与到具体程序运算中，会自动去除多余秒数精确到分钟，实际精度只到分钟。精确到分钟后的录制结束时间必须大于当前时间且大于精确到分钟后的开始时间。结束时间与开始时间去除秒数精确到分钟的时间，最大间隔不能超过 72 小时。录制的视频时长与计划录制时间会有少量误差。"}
  EndTime *int32 `json:"endTime,omitempty" xml:"endTime,omitempty" require:"true"`
  // {"en":"User-defined notification address", "zh_CN":"用户自定义的通知地址，必须以 http://开头或 https://开头。"}
  NotifyUrl *string `json:"notifyUrl,omitempty" xml:"notifyUrl,omitempty"`
  // {"en":"Recorded file format, default is mp4, optional flv, or m3u8, value is lowercase", "zh_CN":"录制文件的格式，默认是 mp4,可选 flv，或m3u8，值为小写"}
  FileType *string `json:"fileType,omitempty" xml:"fileType,omitempty"`
  // {"en":"The default value is flase. If true, the fileType must be in m3u8 format. Otherwise, a parameter error will be reported. The recording function of third-party storage does not support splicing", "zh_CN":"是否合并成一个视频，默认是 flase，如果选true,则 fileType 必须是 m3u8 格式才支持，否则会报参数错误。开启第三方存储录制时不支持拼接"}
  IsConcat *bool `json:"isConcat,omitempty" xml:"isConcat,omitempty"`
  // {"en":"The default value is false. When set to true, http pull streams are invoked to record. The specific decision logic is as follows: 1. There are two types of http headers: http and https. 2. Encapsulation protocol determination. The default is HDL. 1) If the corresponding pull-flow protocol includes HDL and HLS. HDL is preferred. 2) If the corresponding pull-flow protocol only has HLS, set the value to HLS. 3) If the corresponding pull-flow protocol does not have HDL and HLS. An error message is displayed. If the parameter is true, the domain name corresponding to the flow needs to be configured with HDL or HLS", "zh_CN":"默认值 false，当设置成 true 时，则调用 http拉流去做录制。具体判定逻辑：1、http 协议头有 http 和 https 两种，内部通过查询拉流配置确定。2、封装协议判定。默认取 HDL。1）如对应的拉流协议有 HDL,HLS 两种。优先取HDL。2）如对应拉流协议只有 HLS，取 HLS。3）如对应拉流协议无 HDL 和 HLS。报错，返回提示当该入参为“true”时，该流对应的域名需要配置 HDL 或者 HLS 协议的拉流。"}
  HttpPullRecord *bool `json:"httpPullRecord,omitempty" xml:"httpPullRecord,omitempty"`
  // {"en":"Third-party storage Specifies the storage path for recording (without file suffixes).", "zh_CN":"第三方存储录制时指定存储路径（不含文件后缀）"}
  Key *string `json:"key,omitempty" xml:"key,omitempty"`
}

func (s CreateRecordTaskRequest) String() string {
  return tea.Prettify(s)
}

func (s CreateRecordTaskRequest) GoString() string {
  return s.String()
}

func (s *CreateRecordTaskRequest) SetPullId(v string) *CreateRecordTaskRequest {
  s.PullId = &v
  return s
}

func (s *CreateRecordTaskRequest) SetTransNo(v string) *CreateRecordTaskRequest {
  s.TransNo = &v
  return s
}

func (s *CreateRecordTaskRequest) SetQuickStart(v bool) *CreateRecordTaskRequest {
  s.QuickStart = &v
  return s
}

func (s *CreateRecordTaskRequest) SetStartTime(v int32) *CreateRecordTaskRequest {
  s.StartTime = &v
  return s
}

func (s *CreateRecordTaskRequest) SetEndTime(v int32) *CreateRecordTaskRequest {
  s.EndTime = &v
  return s
}

func (s *CreateRecordTaskRequest) SetNotifyUrl(v string) *CreateRecordTaskRequest {
  s.NotifyUrl = &v
  return s
}

func (s *CreateRecordTaskRequest) SetFileType(v string) *CreateRecordTaskRequest {
  s.FileType = &v
  return s
}

func (s *CreateRecordTaskRequest) SetIsConcat(v bool) *CreateRecordTaskRequest {
  s.IsConcat = &v
  return s
}

func (s *CreateRecordTaskRequest) SetHttpPullRecord(v bool) *CreateRecordTaskRequest {
  s.HttpPullRecord = &v
  return s
}

func (s *CreateRecordTaskRequest) SetKey(v string) *CreateRecordTaskRequest {
  s.Key = &v
  return s
}

type CreateRecordTaskResponse struct {
  // {"en":"200 success", "zh_CN":"200为成功"}
  Code *int32 `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Successful operation", "zh_CN":"操作成功"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Return data", "zh_CN":"返回数据"}
  CreateRecordTaskData *CreateRecordTaskData `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s CreateRecordTaskResponse) String() string {
  return tea.Prettify(s)
}

func (s CreateRecordTaskResponse) GoString() string {
  return s.String()
}

func (s *CreateRecordTaskResponse) SetCode(v int32) *CreateRecordTaskResponse {
  s.Code = &v
  return s
}

func (s *CreateRecordTaskResponse) SetMessage(v string) *CreateRecordTaskResponse {
  s.Message = &v
  return s
}

func (s *CreateRecordTaskResponse) SetData(v *CreateRecordTaskData) *CreateRecordTaskResponse {
  s.CreateRecordTaskData = v
  return s
}

type CreateRecordTaskData struct {
  // {"en":"taskId", "zh_CN":"直播录制任务id"}
  TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty" require:"true"`
}

func (s CreateRecordTaskData) String() string {
  return tea.Prettify(s)
}

func (s CreateRecordTaskData) GoString() string {
  return s.String()
}

func (s *CreateRecordTaskData) SetTaskId(v string) *CreateRecordTaskData {
  s.TaskId = &v
  return s
}

type CreateRecordTaskPaths struct {
}

func (s CreateRecordTaskPaths) String() string {
  return tea.Prettify(s)
}

func (s CreateRecordTaskPaths) GoString() string {
  return s.String()
}

type CreateRecordTaskParameters struct {
}

func (s CreateRecordTaskParameters) String() string {
  return tea.Prettify(s)
}

func (s CreateRecordTaskParameters) GoString() string {
  return s.String()
}

type CreateRecordTaskRequestHeader struct {
}

func (s CreateRecordTaskRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s CreateRecordTaskRequestHeader) GoString() string {
  return s.String()
}

type CreateRecordTaskResponseHeader struct {
}

func (s CreateRecordTaskResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s CreateRecordTaskResponseHeader) GoString() string {
  return s.String()
}




type CreateChannelRequest struct {
  // {"en":"Channel name,the value contains a maximum of 255 characters", "zh_CN":"频道名称，最长 255 个字符"}
  ChannelName *string `json:"channelName,omitempty" xml:"channelName,omitempty" require:"true"`
  // {"en":"Channel type,It must be the channel type configured under the name of the push basin.
  // It can only be a combination of English digits and underscores.
  // The system automatically generates a push
  // Use this domain when streaming urls,
  // For example:
  // rtmp:// push channel name/channel type/pushId
  // pushId is automatically assigned by the system after the channel is successfully created", "zh_CN":"频道类型，必须是推流域名下配置的频道类型，
  // 只能是英文数字下划线组合；系统自动生成推
  // 流 URL 时使用该域名，例如： 
  // rtmp://推流域名／频道类型／pushId 
  // pushId 在频道创建成功后由系统自动分配"}
  ChannelType *string `json:"channelType,omitempty" xml:"channelType,omitempty" require:"true"`
  // {"en":"Source push domain name,This domain name is used when the system automatically generates a URL for pushing streams.
  // For example, rtmp:// push the channel name/channel type/pushId
  // pushId is automatically allocated by the system after the channel is created successfully.", "zh_CN":"源端推流域名；  
  // 创建成功后系统自动生成推流 URL 时使用该域
  // 名，例如：rtmp://推流域名／频道类型／
  // pushId 
  // pushId 在频道创建成功后由系统自动分配；"}
  PushDomain *string `json:"pushDomain,omitempty" xml:"pushDomain,omitempty" require:"true"`
  // {"en":"Play the pull stream protocol,The default value is 0. Multiple options can be selected.
  // 2 and 3 cannot be selected at the same time. The protocol must be included in the basin name
  // The actual configuration protocol type is as follows:
  // 0:RTMP, 1:HDL, 2:HLS, 3:HDS
  // After the URL is created successfully, the system automatically generates the URL for playing the pull stream:
  // For example, rtmp:// pull watershed name/channel type/pushId.
  // If 2.19 Creating a Live Recording Task is required,
  // This parameter must contain at least one of RTMP or HDL protocols when creating a channel.", "zh_CN":"播放拉流协议，默认为 0，可多选； 
  // 2 和 3 不可同时选择，协议需包含在推流域名
  // 实际配置的转出协议类型中： 
  // 0、 RTMP，1、HDL，2、HLS，3、HDS 
  // 创建成功后，系统会自动生成播放拉流 URL：例
  // 如：rtmp://拉流域名／频道类型／pushId。 
  // 如果需要使用 2.19 创建直播录制任务接口进
  // 行录制，则创建频道时该参数必须至少包含
  // RTMP 协议或 HDL 协议中的一个。"}
  PullProtocol *string `json:"pullProtocol,omitempty" xml:"pullProtocol,omitempty"`
  // {"en":"Player pull bit rate,Multiple options can be selected to ensure that the selected bit rate is included in the range of the transcoding bit rate configured in the push basin name;
  // 0:source rate, 1:smooth, 2:standard definition, 3:HD, 4:super clear", "zh_CN":"播放端拉流码率，可多选，确保选择的码率包含在
  // 推流域名配置的转码码率范围内； 
  // 0、源码率，1、流畅，2、标清，3、高清，4、超清 "}
  PullRate *string `json:"pullRate,omitempty" xml:"pullRate,omitempty" require:"true"`
  // {"en":"The player ID used by the channel. If not set, the default player is used.", "zh_CN":"频道使用的播放器 ID，未设置则使用默认播放器。"}
  PlayerId *string `json:"playerId,omitempty" xml:"playerId,omitempty"`
  // {"en":"Record directory ID. For details, see Creating a Record Directory. If the record directory management function is not enabled, you do not need to enter this parameter.", "zh_CN":"录制目录 ID，详见 新建录制目录 如果未开通录制目录管理功能，无需填写。"}
  DirectoryId *string `json:"directoryId,omitempty" xml:"directoryId,omitempty"`
  // {"en":"Stream check string", "zh_CN":"流校验串"}
  Check *string `json:"check,omitempty" xml:"check,omitempty"`
  // {"en":"If this parameter is not specified, the whiteboard function remains unchanged. urlsafe Base64-encoded JSON string,
  // Example: eyJvcGVuIjogMSwid2hpdGVCb2FyZElkIjogInh4eCJ9
  // The format of the json string before encoding is as follows:
  // {\"open\": 1,\"whiteBoardId\": \"xxx\"}
  // There are two attributes:
  // open, mandatory. Whether to enable the whiteboard. 0 is disabled and 1 is enabled.
  // whiteBoardId is mandatory. This parameter is mandatory only when the whiteboard id and open are 1.
  // This interface does not support changing whiteboard ids", "zh_CN":"如果不传此参数，则白板功能保持原样，不修改。经过 urlsafe base64 编码的 JSON 字符串, 
  // 例：
  // eyJvcGVuIjogMSwid2hpdGVCb2FyZElkIjogInh4eCJ
  // 9 
  // 编码前的 json 串格式如下： 
  // {\"open\": 1,\"whiteBoardId\": \"xxx\"} 
  // 有两个属性： 
  // open，必填，是否开启白板，0 不开启，1 开启； 
  // whiteBoardId，非必填，白板 id，open 为 1 时才
  // 需要设值。 
  // 此接口不支持修改白板 id "}
  WhiteBoardConfig *string `json:"whiteBoardConfig,omitempty" xml:"whiteBoardConfig,omitempty"`
  // {"en":"Function configuration parameter. Currently, you can only set whether to display the share button on the web page. If the parameter is not transmitted, no modification is performed.
  // Parameter values are urlsafe base64 encoded JSON strings,
  // For example, eyJzaGFyZSI6IDF9,
  // The json string before encoding is {\"share\": 1}
  // For example: eyJzaGFyZSI6IDB9,
  // The json string before encoding is {\"share\": 0}
  // The funcConfig parameter must be valid if it is passed and is not empty, and the value of the share entry can only be one of 0 or 1.", "zh_CN":"功能配置参数，目前只支持设置是否在网页推流
  // 页面显示分享按钮。如果参数不传，则不做任何
  // 修改。 
  // 参数值为经过 urlsafe base64 编码的 JSON 字符
  // 串，例如：eyJzaGFyZSI6IDF9，编码前的 json 串
  // 为{\"share\": 1} 
  // 如：eyJzaGFyZSI6IDB9，编码前的 json 串为
  // {\"share\": 0} 
  // funcConfig 参数如果有传且不为空，所传值必须
  // 为合法，且 share 项的值只能是 0 或 1 中的某一
  // 个。 "}
  FuncConfig *string `json:"funcConfig,omitempty" xml:"funcConfig,omitempty"`
  // {"en":"If it is enabled, push it and record it. 1 is enabled, 0 is not enabled. If the customer does not provide the required value-added service, this parameter is ignored.", "zh_CN":"是否开启即推即录，1 开启，0 不开启。如果该客户
  // 未开通应对的增值服务，接口中忽略此参数。"}
  Live2vod *string `json:"live2vod,omitempty" xml:"live2vod,omitempty"`
  // {"en":"Application scenario: The specified application scenario must exist", "zh_CN":"应用场景，指定应用场景时应用场景需已存在"}
  AppScenario *string `json:"appScenario,omitempty" xml:"appScenario,omitempty"`
  // {"en":"Video stream encryption. 0: No encryption is required 1:HLS universal encryption 2:DRM encryption. The default value is 0", "zh_CN":"视频流加密。0:不做任何加密 1:HLS通用加密 2:DRM加密，默认为0"}
  VideoEncrypted *int32 `json:"videoEncrypted,omitempty" xml:"videoEncrypted,omitempty"`
  // {"en":"HLS Universal encryption validity time (unit: second). If videoEncryted encryption type is HLS Universal Encryption and the current field is not filled in, the default value is 604800 (7 days)", "zh_CN":"HLS通用加密有效时间（单位：秒）。如果videoEncryted的加密类型为HLS通用加密，而当前字段未填，则默认值为604800（7天）"}
  HlsEncryptDuration *int64 `json:"hlsEncryptDuration,omitempty" xml:"hlsEncryptDuration,omitempty"`
}

func (s CreateChannelRequest) String() string {
  return tea.Prettify(s)
}

func (s CreateChannelRequest) GoString() string {
  return s.String()
}

func (s *CreateChannelRequest) SetChannelName(v string) *CreateChannelRequest {
  s.ChannelName = &v
  return s
}

func (s *CreateChannelRequest) SetChannelType(v string) *CreateChannelRequest {
  s.ChannelType = &v
  return s
}

func (s *CreateChannelRequest) SetPushDomain(v string) *CreateChannelRequest {
  s.PushDomain = &v
  return s
}

func (s *CreateChannelRequest) SetPullProtocol(v string) *CreateChannelRequest {
  s.PullProtocol = &v
  return s
}

func (s *CreateChannelRequest) SetPullRate(v string) *CreateChannelRequest {
  s.PullRate = &v
  return s
}

func (s *CreateChannelRequest) SetPlayerId(v string) *CreateChannelRequest {
  s.PlayerId = &v
  return s
}

func (s *CreateChannelRequest) SetDirectoryId(v string) *CreateChannelRequest {
  s.DirectoryId = &v
  return s
}

func (s *CreateChannelRequest) SetCheck(v string) *CreateChannelRequest {
  s.Check = &v
  return s
}

func (s *CreateChannelRequest) SetWhiteBoardConfig(v string) *CreateChannelRequest {
  s.WhiteBoardConfig = &v
  return s
}

func (s *CreateChannelRequest) SetFuncConfig(v string) *CreateChannelRequest {
  s.FuncConfig = &v
  return s
}

func (s *CreateChannelRequest) SetLive2vod(v string) *CreateChannelRequest {
  s.Live2vod = &v
  return s
}

func (s *CreateChannelRequest) SetAppScenario(v string) *CreateChannelRequest {
  s.AppScenario = &v
  return s
}

func (s *CreateChannelRequest) SetVideoEncrypted(v int32) *CreateChannelRequest {
  s.VideoEncrypted = &v
  return s
}

func (s *CreateChannelRequest) SetHlsEncryptDuration(v int64) *CreateChannelRequest {
  s.HlsEncryptDuration = &v
  return s
}

type CreateChannelResponse struct {
  // {"en":"Status code", "zh_CN":"状态码"}
  Code *int32 `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Operational information", "zh_CN":"操作信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Return data", "zh_CN":"返回数据"}
  CreateChannelData *CreateChannelData `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s CreateChannelResponse) String() string {
  return tea.Prettify(s)
}

func (s CreateChannelResponse) GoString() string {
  return s.String()
}

func (s *CreateChannelResponse) SetCode(v int32) *CreateChannelResponse {
  s.Code = &v
  return s
}

func (s *CreateChannelResponse) SetMessage(v string) *CreateChannelResponse {
  s.Message = &v
  return s
}

func (s *CreateChannelResponse) SetData(v *CreateChannelData) *CreateChannelResponse {
  s.CreateChannelData = v
  return s
}

type CreateChannelData struct {
  // {"en":"Create timestamp", "zh_CN":"创建时间戳"}
  CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty" require:"true"`
  // {"en":"Pull stream ID, a globally unique value.", "zh_CN":"拉流 ID，全局唯一值。"}
  PullId *string `json:"pullId,omitempty" xml:"pullId,omitempty" require:"true"`
}

func (s CreateChannelData) String() string {
  return tea.Prettify(s)
}

func (s CreateChannelData) GoString() string {
  return s.String()
}

func (s *CreateChannelData) SetCreateTime(v string) *CreateChannelData {
  s.CreateTime = &v
  return s
}

func (s *CreateChannelData) SetPullId(v string) *CreateChannelData {
  s.PullId = &v
  return s
}

type CreateChannelPaths struct {
}

func (s CreateChannelPaths) String() string {
  return tea.Prettify(s)
}

func (s CreateChannelPaths) GoString() string {
  return s.String()
}

type CreateChannelParameters struct {
}

func (s CreateChannelParameters) String() string {
  return tea.Prettify(s)
}

func (s CreateChannelParameters) GoString() string {
  return s.String()
}

type CreateChannelRequestHeader struct {
}

func (s CreateChannelRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s CreateChannelRequestHeader) GoString() string {
  return s.String()
}

type CreateChannelResponseHeader struct {
}

func (s CreateChannelResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s CreateChannelResponseHeader) GoString() string {
  return s.String()
}




type SetChannelStateRealtimeBackRequest struct {
  // {"en":"Pushname, you need to enable the channel status callback function to push the basin name", "zh_CN":"推流域名，需要开启频道状态回调功能的推流域名"}
  DomainName *string `json:"domainName,omitempty" xml:"domainName,omitempty" require:"true"`
  // {"en":"The callback address must start with http:// or https:// in lower case.", "zh_CN":"回调地址，必须以小写的 http://开头或https://开头。"}
  NotifyUrl *string `json:"notifyUrl,omitempty" xml:"notifyUrl,omitempty" require:"true"`
  // {"en":"1: enable, 0: disable. The default value is 1", "zh_CN":"1：开启，0：关闭，新增时默认 1"}
  Open *int32 `json:"open,omitempty" xml:"open,omitempty" require:"true"`
}

func (s SetChannelStateRealtimeBackRequest) String() string {
  return tea.Prettify(s)
}

func (s SetChannelStateRealtimeBackRequest) GoString() string {
  return s.String()
}

func (s *SetChannelStateRealtimeBackRequest) SetDomainName(v string) *SetChannelStateRealtimeBackRequest {
  s.DomainName = &v
  return s
}

func (s *SetChannelStateRealtimeBackRequest) SetNotifyUrl(v string) *SetChannelStateRealtimeBackRequest {
  s.NotifyUrl = &v
  return s
}

func (s *SetChannelStateRealtimeBackRequest) SetOpen(v int32) *SetChannelStateRealtimeBackRequest {
  s.Open = &v
  return s
}

type SetChannelStateRealtimeBackResponse struct {
  // {"en":"200 success", "zh_CN":"200为成功"}
  Code *int32 `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Successful operation", "zh_CN":"操作成功"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Return data", "zh_CN":"返回数据"}
  SetChannelStateRealtimeBackData *SetChannelStateRealtimeBackData `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s SetChannelStateRealtimeBackResponse) String() string {
  return tea.Prettify(s)
}

func (s SetChannelStateRealtimeBackResponse) GoString() string {
  return s.String()
}

func (s *SetChannelStateRealtimeBackResponse) SetCode(v int32) *SetChannelStateRealtimeBackResponse {
  s.Code = &v
  return s
}

func (s *SetChannelStateRealtimeBackResponse) SetMessage(v string) *SetChannelStateRealtimeBackResponse {
  s.Message = &v
  return s
}

func (s *SetChannelStateRealtimeBackResponse) SetData(v *SetChannelStateRealtimeBackData) *SetChannelStateRealtimeBackResponse {
  s.SetChannelStateRealtimeBackData = v
  return s
}

type SetChannelStateRealtimeBackData struct {
}

func (s SetChannelStateRealtimeBackData) String() string {
  return tea.Prettify(s)
}

func (s SetChannelStateRealtimeBackData) GoString() string {
  return s.String()
}

type SetChannelStateRealtimeBackPaths struct {
}

func (s SetChannelStateRealtimeBackPaths) String() string {
  return tea.Prettify(s)
}

func (s SetChannelStateRealtimeBackPaths) GoString() string {
  return s.String()
}

type SetChannelStateRealtimeBackParameters struct {
}

func (s SetChannelStateRealtimeBackParameters) String() string {
  return tea.Prettify(s)
}

func (s SetChannelStateRealtimeBackParameters) GoString() string {
  return s.String()
}

type SetChannelStateRealtimeBackRequestHeader struct {
}

func (s SetChannelStateRealtimeBackRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s SetChannelStateRealtimeBackRequestHeader) GoString() string {
  return s.String()
}

type SetChannelStateRealtimeBackResponseHeader struct {
}

func (s SetChannelStateRealtimeBackResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s SetChannelStateRealtimeBackResponseHeader) GoString() string {
  return s.String()
}




type DeleteChannelRequest struct {
  // {"en":"Pull stream id, channel unique identification, multiple ids separated by \",\";", "zh_CN":"拉流 id，频道唯一标识，多个 ID 用\",\"隔开；"}
  PullId *string `json:"pullId,omitempty" xml:"pullId,omitempty" require:"true"`
}

func (s DeleteChannelRequest) String() string {
  return tea.Prettify(s)
}

func (s DeleteChannelRequest) GoString() string {
  return s.String()
}

func (s *DeleteChannelRequest) SetPullId(v string) *DeleteChannelRequest {
  s.PullId = &v
  return s
}

type DeleteChannelResponse struct {
  // {"en":"code", "zh_CN":"状态码"}
  Code *int32 `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"message", "zh_CN":"操作信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Return data", "zh_CN":"返回数据"}
  Data *string `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s DeleteChannelResponse) String() string {
  return tea.Prettify(s)
}

func (s DeleteChannelResponse) GoString() string {
  return s.String()
}

func (s *DeleteChannelResponse) SetCode(v int32) *DeleteChannelResponse {
  s.Code = &v
  return s
}

func (s *DeleteChannelResponse) SetMessage(v string) *DeleteChannelResponse {
  s.Message = &v
  return s
}

func (s *DeleteChannelResponse) SetData(v string) *DeleteChannelResponse {
  s.Data = &v
  return s
}

type DeleteChannelPaths struct {
}

func (s DeleteChannelPaths) String() string {
  return tea.Prettify(s)
}

func (s DeleteChannelPaths) GoString() string {
  return s.String()
}

type DeleteChannelParameters struct {
}

func (s DeleteChannelParameters) String() string {
  return tea.Prettify(s)
}

func (s DeleteChannelParameters) GoString() string {
  return s.String()
}

type DeleteChannelRequestHeader struct {
}

func (s DeleteChannelRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteChannelRequestHeader) GoString() string {
  return s.String()
}

type DeleteChannelResponseHeader struct {
}

func (s DeleteChannelResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteChannelResponseHeader) GoString() string {
  return s.String()
}




type ChannelEditRequest struct {
  // {"en":"Channel pull id", "zh_CN":"频道拉流 id"}
  PullId *string `json:"pullId,omitempty" xml:"pullId,omitempty" require:"true"`
  // {"en":"Channel name,The value contains a maximum of 255 characters", "zh_CN":"频道名称，最长 255 个字符"}
  ChannelName *string `json:"channelName,omitempty" xml:"channelName,omitempty"`
  // {"en":"Reset the push stream ID. The default is 0,1 reset and 0 does not reset", "zh_CN":"重设推流 ID，默认为 0，1 为重设，0 为不重设"}
  ResetPushId *string `json:"resetPushId,omitempty" xml:"resetPushId,omitempty"`
  // {"en":"The player pulls the stream bit rate. Multiple options can be selected to ensure that the selected bit rate is included in the range of the transcoding bit rate configured in the push basin name.
  // 0:source rate, 1:smooth, 2:standard definition, 3:HD, 4:super clear", "zh_CN":"播放端拉流码率，可多选，确保选择的码率包含在
  // 推流域名配置的转码码率范围内； 
  // 0、源码率，1、流畅，2、标清，3、高清，4、超清 "}
  PullRate *string `json:"pullRate,omitempty" xml:"pullRate,omitempty"`
  // {"en":"ID of the recording directory. For details, see Creating a Recording Directory
  // If the recorded directory management function is not enabled, this parameter is not required.", "zh_CN":"录制目录 ID，详见 新建录制目录 
  // 如果未开通录制目录管理功能，无需填写。"}
  DirectoryId *string `json:"directoryId,omitempty" xml:"directoryId,omitempty"`
  // {"en":"Stream check string", "zh_CN":"流校验串"}
  Check *string `json:"check,omitempty" xml:"check,omitempty"`
  // {"en":"If this parameter is not specified, the whiteboard function remains unchanged. urlsafe Base64-encoded JSON string,
  // Example: eyJvcGVuIjogMSwid2hpdGVCb2FyZElkIjogInh4eCJ9
  // The format of the json string before encoding is as follows:
  // {\"open\": 1,\"whiteBoardId\": \"xxx\"}
  // There are two attributes:
  // open, mandatory. Whether to enable the whiteboard. 0 is disabled and 1 is enabled.
  // whiteBoardId is mandatory. This parameter is mandatory only when the whiteboard id and open are 1.
  // This interface does not support changing whiteboard ids", "zh_CN":"如果不传此参数，则白板功能保持原样，不修改。经过 urlsafe base64 编码的 JSON 字符串, 
  // 例：
  // eyJvcGVuIjogMSwid2hpdGVCb2FyZElkIjogInh4eCJ9 
  // 编码前的 json 串格式如下： 
  // {\"open\": 1,\"whiteBoardId\": \"xxx\"} 
  // 有两个属性： 
  // open，必填，是否开启白板，0 不开启，1 开启； 
  // whiteBoardId，非必填，白板 id，open 为 1 时才
  // 需要设值。 
  // 此接口不支持修改白板 id "}
  WhiteBoardConfig *string `json:"whiteBoardConfig,omitempty" xml:"whiteBoardConfig,omitempty"`
  // {"en":"Function configuration parameter. Currently, you can only set whether to display the share button on the web page. If the parameter is not transmitted, no modification is performed.
  // Parameter values are urlsafe base64 encoded JSON strings,
  // For example, eyJzaGFyZSI6IDF9,
  // The json string before encoding is {\"share\": 1}
  // For example: eyJzaGFyZSI6IDB9,
  // The json string before encoding is {\"share\": 0}
  // The funcConfig parameter must be valid if it is passed and is not empty, and the value of the share entry can only be one of 0 or 1.", "zh_CN":"功能配置参数，目前只支持设置是否在网页推流
  // 页面显示分享按钮。如果参数不传，则不做任何
  // 修改。 
  // 参数值为经过 urlsafe base64 编码的 JSON 字符
  // 串，例如：eyJzaGFyZSI6IDF9，编码前的 json 串
  // 为{\"share\": 1} 
  // 如：eyJzaGFyZSI6IDB9，编码前的 json 串为
  // {\"share\": 0} 
  // funcConfig 参数如果有传且不为空，所传值必须
  // 为合法，且 share 项的值只能是 0 或 1 中的某一
  // 个。 "}
  FuncConfig *string `json:"funcConfig,omitempty" xml:"funcConfig,omitempty"`
  // {"en":"If it is enabled, push it and record it. 1 is enabled, 0 is not enabled. If the customer
  // This parameter is ignored because the value-added service is not enabled.", "zh_CN":"是否开启即推即录，1 开启，0 不开启。如果该客户
  // 未开通应对的增值服务，接口中忽略此参数。"}
  Live2vod *string `json:"live2vod,omitempty" xml:"live2vod,omitempty"`
  // {"en":"Application scenario: The specified application scenario must exist", "zh_CN":"应用场景，指定应用场景时应用场景需已存在"}
  AppScenario *string `json:"appScenario,omitempty" xml:"appScenario,omitempty"`
  // {"en":"Video stream encryption. 0: No encryption is required 1:HLS universal encryption 2:DRM encryption", "zh_CN":"视频流加密。0:不做任何加密 1:HLS通用加密 2:DRM加密"}
  VideoEncrypted *int32 `json:"videoEncrypted,omitempty" xml:"videoEncrypted,omitempty"`
  // {"en":"HLS Universal encryption validity time (unit: second).
  // 1. videoEncryted encryption type has been HLS universal encryption
  // (1) If the current field has been filled in, the valid duration will be updated according to the current time and the valid time.
  // (2) If the current field is not filled in, the validity period will not be updated.
  // 2. videoEncryted encryption type changed from other to HLS universal encryption,
  // (1) If the current field is not filled in, the default value is 604800 (7 days).
  // (2) If the current field has been filled in, the valid duration will be updated according to the current time and the valid time.", "zh_CN":"HLS通用加密有效时间（单位：秒）。
  // 1. videoEncryted的加密类型已经是HLS通用加密的情况
  //     （1）当前字段已填写，则根据当前时间和填入有效时间重新更新有效时长。
  //     （2）当前字段未填写，则不更新有效时长。
  // 2. videoEncryted的加密类型由其他改为HLS通用加密的情况，
  //     （1）当前字段未填写，则默认值为604800（7天）。
  //     （2）当前字段已填写，则根据当前时间和填入有效时间重新更新有效时长。"}
  HlsEncryptDuration *int64 `json:"hlsEncryptDuration,omitempty" xml:"hlsEncryptDuration,omitempty"`
}

func (s ChannelEditRequest) String() string {
  return tea.Prettify(s)
}

func (s ChannelEditRequest) GoString() string {
  return s.String()
}

func (s *ChannelEditRequest) SetPullId(v string) *ChannelEditRequest {
  s.PullId = &v
  return s
}

func (s *ChannelEditRequest) SetChannelName(v string) *ChannelEditRequest {
  s.ChannelName = &v
  return s
}

func (s *ChannelEditRequest) SetResetPushId(v string) *ChannelEditRequest {
  s.ResetPushId = &v
  return s
}

func (s *ChannelEditRequest) SetPullRate(v string) *ChannelEditRequest {
  s.PullRate = &v
  return s
}

func (s *ChannelEditRequest) SetDirectoryId(v string) *ChannelEditRequest {
  s.DirectoryId = &v
  return s
}

func (s *ChannelEditRequest) SetCheck(v string) *ChannelEditRequest {
  s.Check = &v
  return s
}

func (s *ChannelEditRequest) SetWhiteBoardConfig(v string) *ChannelEditRequest {
  s.WhiteBoardConfig = &v
  return s
}

func (s *ChannelEditRequest) SetFuncConfig(v string) *ChannelEditRequest {
  s.FuncConfig = &v
  return s
}

func (s *ChannelEditRequest) SetLive2vod(v string) *ChannelEditRequest {
  s.Live2vod = &v
  return s
}

func (s *ChannelEditRequest) SetAppScenario(v string) *ChannelEditRequest {
  s.AppScenario = &v
  return s
}

func (s *ChannelEditRequest) SetVideoEncrypted(v int32) *ChannelEditRequest {
  s.VideoEncrypted = &v
  return s
}

func (s *ChannelEditRequest) SetHlsEncryptDuration(v int64) *ChannelEditRequest {
  s.HlsEncryptDuration = &v
  return s
}

type ChannelEditResponse struct {
  // {"en":"status code", "zh_CN":"状态码"}
  Code *int32 `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"message", "zh_CN":"操作信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"return data", "zh_CN":"返回数据"}
  ChannelEditData *ChannelEditData `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s ChannelEditResponse) String() string {
  return tea.Prettify(s)
}

func (s ChannelEditResponse) GoString() string {
  return s.String()
}

func (s *ChannelEditResponse) SetCode(v int32) *ChannelEditResponse {
  s.Code = &v
  return s
}

func (s *ChannelEditResponse) SetMessage(v string) *ChannelEditResponse {
  s.Message = &v
  return s
}

func (s *ChannelEditResponse) SetData(v *ChannelEditData) *ChannelEditResponse {
  s.ChannelEditData = v
  return s
}

type ChannelEditData struct {
  // {"en":"Channel name,The value contains a maximum of 255 characters", "zh_CN":"频道名称，最长 255 个字符"}
  ChannelName *string `json:"channelName,omitempty" xml:"channelName,omitempty" require:"true"`
  // {"en":"Push URL", "zh_CN":"推流 URL"}
  PushUrl *string `json:"pushUrl,omitempty" xml:"pushUrl,omitempty" require:"true"`
  // {"en":"The URL is in SRT format", "zh_CN":"SRT 格式的推流 URL"}
  SrtPushUrl *string `json:"srtPushUrl,omitempty" xml:"srtPushUrl,omitempty" require:"true"`
}

func (s ChannelEditData) String() string {
  return tea.Prettify(s)
}

func (s ChannelEditData) GoString() string {
  return s.String()
}

func (s *ChannelEditData) SetChannelName(v string) *ChannelEditData {
  s.ChannelName = &v
  return s
}

func (s *ChannelEditData) SetPushUrl(v string) *ChannelEditData {
  s.PushUrl = &v
  return s
}

func (s *ChannelEditData) SetSrtPushUrl(v string) *ChannelEditData {
  s.SrtPushUrl = &v
  return s
}

type ChannelEditPaths struct {
}

func (s ChannelEditPaths) String() string {
  return tea.Prettify(s)
}

func (s ChannelEditPaths) GoString() string {
  return s.String()
}

type ChannelEditParameters struct {
}

func (s ChannelEditParameters) String() string {
  return tea.Prettify(s)
}

func (s ChannelEditParameters) GoString() string {
  return s.String()
}

type ChannelEditRequestHeader struct {
}

func (s ChannelEditRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ChannelEditRequestHeader) GoString() string {
  return s.String()
}

type ChannelEditResponseHeader struct {
}

func (s ChannelEditResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ChannelEditResponseHeader) GoString() string {
  return s.String()
}




type ChannelReBrodcastRequest struct {
  // {"en":"Pull stream id, channel unique identification, multiple ids separated by \",\";", "zh_CN":"拉流 id，频道唯一标识，多个 ID 用\",\"隔开；"}
  PullId *string `json:"pullId,omitempty" xml:"pullId,omitempty" require:"true"`
}

func (s ChannelReBrodcastRequest) String() string {
  return tea.Prettify(s)
}

func (s ChannelReBrodcastRequest) GoString() string {
  return s.String()
}

func (s *ChannelReBrodcastRequest) SetPullId(v string) *ChannelReBrodcastRequest {
  s.PullId = &v
  return s
}

type ChannelReBrodcastResponse struct {
  // {"en":"status code", "zh_CN":"状态码"}
  Code *int32 `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"message", "zh_CN":"操作信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"return data", "zh_CN":"返回数据"}
  Data *string `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s ChannelReBrodcastResponse) String() string {
  return tea.Prettify(s)
}

func (s ChannelReBrodcastResponse) GoString() string {
  return s.String()
}

func (s *ChannelReBrodcastResponse) SetCode(v int32) *ChannelReBrodcastResponse {
  s.Code = &v
  return s
}

func (s *ChannelReBrodcastResponse) SetMessage(v string) *ChannelReBrodcastResponse {
  s.Message = &v
  return s
}

func (s *ChannelReBrodcastResponse) SetData(v string) *ChannelReBrodcastResponse {
  s.Data = &v
  return s
}

type ChannelReBrodcastPaths struct {
}

func (s ChannelReBrodcastPaths) String() string {
  return tea.Prettify(s)
}

func (s ChannelReBrodcastPaths) GoString() string {
  return s.String()
}

type ChannelReBrodcastParameters struct {
}

func (s ChannelReBrodcastParameters) String() string {
  return tea.Prettify(s)
}

func (s ChannelReBrodcastParameters) GoString() string {
  return s.String()
}

type ChannelReBrodcastRequestHeader struct {
}

func (s ChannelReBrodcastRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ChannelReBrodcastRequestHeader) GoString() string {
  return s.String()
}

type ChannelReBrodcastResponseHeader struct {
}

func (s ChannelReBrodcastResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ChannelReBrodcastResponseHeader) GoString() string {
  return s.String()
}




type GetDirectoryRequest struct {
  // {"en":"id of the recording directory whose information you want to obtain", "zh_CN":"需要获取录制目录信息的录制目录id"}
  DirectoryId *string `json:"directoryId,omitempty" xml:"directoryId,omitempty" require:"true"`
}

func (s GetDirectoryRequest) String() string {
  return tea.Prettify(s)
}

func (s GetDirectoryRequest) GoString() string {
  return s.String()
}

func (s *GetDirectoryRequest) SetDirectoryId(v string) *GetDirectoryRequest {
  s.DirectoryId = &v
  return s
}

type GetDirectoryResponse struct {
  // {"en":"200 success", "zh_CN":"200为成功"}
  Code *int32 `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Successful operation", "zh_CN":"操作成功"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Return data", "zh_CN":"返回数据"}
  GetDirectoryData *GetDirectoryData `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s GetDirectoryResponse) String() string {
  return tea.Prettify(s)
}

func (s GetDirectoryResponse) GoString() string {
  return s.String()
}

func (s *GetDirectoryResponse) SetCode(v int32) *GetDirectoryResponse {
  s.Code = &v
  return s
}

func (s *GetDirectoryResponse) SetMessage(v string) *GetDirectoryResponse {
  s.Message = &v
  return s
}

func (s *GetDirectoryResponse) SetData(v *GetDirectoryData) *GetDirectoryResponse {
  s.GetDirectoryData = v
  return s
}

type GetDirectoryData struct {
  // {"en":"Record directory creation time", "zh_CN":"录制目录创建时间"}
  CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty" require:"true"`
  // {"en":"Recorded directory ID", "zh_CN":"录制目录 ID"}
  DirectoryId *string `json:"directoryId,omitempty" xml:"directoryId,omitempty" require:"true"`
  // {"en":"Directory type", "zh_CN":"目录类型"}
  DirectoryType *string `json:"directoryType,omitempty" xml:"directoryType,omitempty" require:"true"`
  // {"en":"Primary directory", "zh_CN":"一级目录"}
  FirstLevelDirectory *string `json:"firstLevelDirectory,omitempty" xml:"firstLevelDirectory,omitempty" require:"true"`
  // {"en":"Secondary directory", "zh_CN":"二级目录"}
  SecondLevelDirectory *string `json:"secondLevelDirectory,omitempty" xml:"secondLevelDirectory,omitempty" require:"true"`
  // {"en":"Three-level directory", "zh_CN":"三级目录"}
  ThirdLevelDirectory *string `json:"thirdLevelDirectory,omitempty" xml:"thirdLevelDirectory,omitempty" require:"true"`
  // {"en":"Record directory modification time", "zh_CN":"录制目录修改时间"}
  UpdateTime *string `json:"updateTime,omitempty" xml:"updateTime,omitempty" require:"true"`
}

func (s GetDirectoryData) String() string {
  return tea.Prettify(s)
}

func (s GetDirectoryData) GoString() string {
  return s.String()
}

func (s *GetDirectoryData) SetCreateTime(v string) *GetDirectoryData {
  s.CreateTime = &v
  return s
}

func (s *GetDirectoryData) SetDirectoryId(v string) *GetDirectoryData {
  s.DirectoryId = &v
  return s
}

func (s *GetDirectoryData) SetDirectoryType(v string) *GetDirectoryData {
  s.DirectoryType = &v
  return s
}

func (s *GetDirectoryData) SetFirstLevelDirectory(v string) *GetDirectoryData {
  s.FirstLevelDirectory = &v
  return s
}

func (s *GetDirectoryData) SetSecondLevelDirectory(v string) *GetDirectoryData {
  s.SecondLevelDirectory = &v
  return s
}

func (s *GetDirectoryData) SetThirdLevelDirectory(v string) *GetDirectoryData {
  s.ThirdLevelDirectory = &v
  return s
}

func (s *GetDirectoryData) SetUpdateTime(v string) *GetDirectoryData {
  s.UpdateTime = &v
  return s
}

type GetDirectoryPaths struct {
}

func (s GetDirectoryPaths) String() string {
  return tea.Prettify(s)
}

func (s GetDirectoryPaths) GoString() string {
  return s.String()
}

type GetDirectoryParameters struct {
}

func (s GetDirectoryParameters) String() string {
  return tea.Prettify(s)
}

func (s GetDirectoryParameters) GoString() string {
  return s.String()
}

type GetDirectoryRequestHeader struct {
}

func (s GetDirectoryRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetDirectoryRequestHeader) GoString() string {
  return s.String()
}

type GetDirectoryResponseHeader struct {
}

func (s GetDirectoryResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetDirectoryResponseHeader) GoString() string {
  return s.String()
}




type GetChannelLiveStateRequest struct {
  // {"en":"pullid", "zh_CN":"拉流 id"}
  PullId *string `json:"pullId,omitempty" xml:"pullId,omitempty" require:"true"`
}

func (s GetChannelLiveStateRequest) String() string {
  return tea.Prettify(s)
}

func (s GetChannelLiveStateRequest) GoString() string {
  return s.String()
}

func (s *GetChannelLiveStateRequest) SetPullId(v string) *GetChannelLiveStateRequest {
  s.PullId = &v
  return s
}

type GetChannelLiveStateResponse struct {
  // {"en":"200 success", "zh_CN":"状态码"}
  Code *int32 `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Successful operation", "zh_CN":"操作信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Return data", "zh_CN":"返回数据"}
  GetChannelLiveStateData *GetChannelLiveStateData `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s GetChannelLiveStateResponse) String() string {
  return tea.Prettify(s)
}

func (s GetChannelLiveStateResponse) GoString() string {
  return s.String()
}

func (s *GetChannelLiveStateResponse) SetCode(v int32) *GetChannelLiveStateResponse {
  s.Code = &v
  return s
}

func (s *GetChannelLiveStateResponse) SetMessage(v string) *GetChannelLiveStateResponse {
  s.Message = &v
  return s
}

func (s *GetChannelLiveStateResponse) SetData(v *GetChannelLiveStateData) *GetChannelLiveStateResponse {
  s.GetChannelLiveStateData = v
  return s
}

type GetChannelLiveStateData struct {
  // {"en":"Return the live stream status corresponding to each pull stream ID:
  // 1:live broadcast 2:not broadcast 3:banned", "zh_CN":"返回每个拉流 ID 对应的直播流实时状态： 
  // 1、直播中   2、未开播  3、禁播 "}
  ChannelState *string `json:"channelState,omitempty" xml:"channelState,omitempty" require:"true"`
}

func (s GetChannelLiveStateData) String() string {
  return tea.Prettify(s)
}

func (s GetChannelLiveStateData) GoString() string {
  return s.String()
}

func (s *GetChannelLiveStateData) SetChannelState(v string) *GetChannelLiveStateData {
  s.ChannelState = &v
  return s
}

type GetChannelLiveStatePaths struct {
}

func (s GetChannelLiveStatePaths) String() string {
  return tea.Prettify(s)
}

func (s GetChannelLiveStatePaths) GoString() string {
  return s.String()
}

type GetChannelLiveStateParameters struct {
}

func (s GetChannelLiveStateParameters) String() string {
  return tea.Prettify(s)
}

func (s GetChannelLiveStateParameters) GoString() string {
  return s.String()
}

type GetChannelLiveStateRequestHeader struct {
}

func (s GetChannelLiveStateRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetChannelLiveStateRequestHeader) GoString() string {
  return s.String()
}

type GetChannelLiveStateResponseHeader struct {
}

func (s GetChannelLiveStateResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetChannelLiveStateResponseHeader) GoString() string {
  return s.String()
}




type GetRecordFileInfoRequest struct {
  // {"en":"pullIds", "zh_CN":"1至多个频道ID，用半角逗号（,）隔开。要求不能输入数字、大小写英文、半角逗号以外的其他字符，否则会认为是非法字符。但是如果输入的是空字符串则不会认为是非法字符，而是认为该参数为空。pullIds 不能全部是逗号。pullIds 中的频道 ID 个数不能超过 50 个。pullIds 和videoIds 两个参数必须有且只有一个不为空。"}
  PullIds *string `json:"pullIds,omitempty" xml:"pullIds,omitempty"`
  // {"en":"videoIds", "zh_CN":"1至多个视频ID，用半角逗号（,）隔开。要求不能输入数字、大小写英文、半角逗号以外的其他字符，否则会认为是非法字符。但是如果输入的是空字符串则不会认为是非法字符，而是认为该参数为空。videoIds 不能全部是逗号。videoIds 中的频道 ID 个数不能超过 50 个。pullIds 和videoIds 两个参数必须有且只有一个不为空。"}
  VideoIds *string `json:"videoIds,omitempty" xml:"videoIds,omitempty"`
}

func (s GetRecordFileInfoRequest) String() string {
  return tea.Prettify(s)
}

func (s GetRecordFileInfoRequest) GoString() string {
  return s.String()
}

func (s *GetRecordFileInfoRequest) SetPullIds(v string) *GetRecordFileInfoRequest {
  s.PullIds = &v
  return s
}

func (s *GetRecordFileInfoRequest) SetVideoIds(v string) *GetRecordFileInfoRequest {
  s.VideoIds = &v
  return s
}

type GetRecordFileInfoResponse struct {
  // {"en":"200 success", "zh_CN":"200为成功"}
  Code *int32 `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Successful operation", "zh_CN":"操作成功"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Return data", "zh_CN":"返回数据"}
  Data []*GetRecordFileInfoList `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s GetRecordFileInfoResponse) String() string {
  return tea.Prettify(s)
}

func (s GetRecordFileInfoResponse) GoString() string {
  return s.String()
}

func (s *GetRecordFileInfoResponse) SetCode(v int32) *GetRecordFileInfoResponse {
  s.Code = &v
  return s
}

func (s *GetRecordFileInfoResponse) SetMessage(v string) *GetRecordFileInfoResponse {
  s.Message = &v
  return s
}

func (s *GetRecordFileInfoResponse) SetData(v []*GetRecordFileInfoList) *GetRecordFileInfoResponse {
  s.Data = v
  return s
}

type GetRecordFileInfoList struct {
  // {"en":"Watch Password", "zh_CN":"观看密码"}
  Password *string `json:"password,omitempty" xml:"password,omitempty" require:"true"`
  // {"en":"Channel pull ID", "zh_CN":"频道拉流 ID"}
  PullId *string `json:"pullId,omitempty" xml:"pullId,omitempty" require:"true"`
  // {"en":"Share page url, if the encrypted transcoding video sharing page url is empty.", "zh_CN":"分享页 url，如果加密转码的视频分享页 url 为空。"}
  SharePageUrl *string `json:"sharePageUrl,omitempty" xml:"sharePageUrl,omitempty" require:"true"`
  // {"en":"Recording start time", "zh_CN":"录制开始时间"}
  StartTime *int32 `json:"startTime,omitempty" xml:"startTime,omitempty" require:"true"`
  // {"en":"Recording duration", "zh_CN":"录制时长"}
  VideoDuration *int32 `json:"videoDuration,omitempty" xml:"videoDuration,omitempty" require:"true"`
  // {"en":"videoId", "zh_CN":"视频 ID"}
  VideoId *string `json:"videoId,omitempty" xml:"videoId,omitempty" require:"true"`
  // {"en":"videoName", "zh_CN":"视频名称"}
  VideoName *string `json:"videoName,omitempty" xml:"videoName,omitempty" require:"true"`
}

func (s GetRecordFileInfoList) String() string {
  return tea.Prettify(s)
}

func (s GetRecordFileInfoList) GoString() string {
  return s.String()
}

func (s *GetRecordFileInfoList) SetPassword(v string) *GetRecordFileInfoList {
  s.Password = &v
  return s
}

func (s *GetRecordFileInfoList) SetPullId(v string) *GetRecordFileInfoList {
  s.PullId = &v
  return s
}

func (s *GetRecordFileInfoList) SetSharePageUrl(v string) *GetRecordFileInfoList {
  s.SharePageUrl = &v
  return s
}

func (s *GetRecordFileInfoList) SetStartTime(v int32) *GetRecordFileInfoList {
  s.StartTime = &v
  return s
}

func (s *GetRecordFileInfoList) SetVideoDuration(v int32) *GetRecordFileInfoList {
  s.VideoDuration = &v
  return s
}

func (s *GetRecordFileInfoList) SetVideoId(v string) *GetRecordFileInfoList {
  s.VideoId = &v
  return s
}

func (s *GetRecordFileInfoList) SetVideoName(v string) *GetRecordFileInfoList {
  s.VideoName = &v
  return s
}

type GetRecordFileInfoPaths struct {
}

func (s GetRecordFileInfoPaths) String() string {
  return tea.Prettify(s)
}

func (s GetRecordFileInfoPaths) GoString() string {
  return s.String()
}

type GetRecordFileInfoParameters struct {
}

func (s GetRecordFileInfoParameters) String() string {
  return tea.Prettify(s)
}

func (s GetRecordFileInfoParameters) GoString() string {
  return s.String()
}

type GetRecordFileInfoRequestHeader struct {
}

func (s GetRecordFileInfoRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetRecordFileInfoRequestHeader) GoString() string {
  return s.String()
}

type GetRecordFileInfoResponseHeader struct {
}

func (s GetRecordFileInfoResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetRecordFileInfoResponseHeader) GoString() string {
  return s.String()
}




type EndForwardTaskRequest struct {
  // {"en":"taskId", "zh_CN":"转推任务 ID"}
  TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty" require:"true"`
}

func (s EndForwardTaskRequest) String() string {
  return tea.Prettify(s)
}

func (s EndForwardTaskRequest) GoString() string {
  return s.String()
}

func (s *EndForwardTaskRequest) SetTaskId(v string) *EndForwardTaskRequest {
  s.TaskId = &v
  return s
}

type EndForwardTaskResponse struct {
  // {"en":"200 success", "zh_CN":"200为成功"}
  Code *int32 `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Successful operation", "zh_CN":"操作成功"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Return data", "zh_CN":"返回数据"}
  Data *string `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s EndForwardTaskResponse) String() string {
  return tea.Prettify(s)
}

func (s EndForwardTaskResponse) GoString() string {
  return s.String()
}

func (s *EndForwardTaskResponse) SetCode(v int32) *EndForwardTaskResponse {
  s.Code = &v
  return s
}

func (s *EndForwardTaskResponse) SetMessage(v string) *EndForwardTaskResponse {
  s.Message = &v
  return s
}

func (s *EndForwardTaskResponse) SetData(v string) *EndForwardTaskResponse {
  s.Data = &v
  return s
}

type EndForwardTaskPaths struct {
}

func (s EndForwardTaskPaths) String() string {
  return tea.Prettify(s)
}

func (s EndForwardTaskPaths) GoString() string {
  return s.String()
}

type EndForwardTaskParameters struct {
}

func (s EndForwardTaskParameters) String() string {
  return tea.Prettify(s)
}

func (s EndForwardTaskParameters) GoString() string {
  return s.String()
}

type EndForwardTaskRequestHeader struct {
}

func (s EndForwardTaskRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s EndForwardTaskRequestHeader) GoString() string {
  return s.String()
}

type EndForwardTaskResponseHeader struct {
}

func (s EndForwardTaskResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s EndForwardTaskResponseHeader) GoString() string {
  return s.String()
}




type ChannelForbiddenRequest struct {
  // {"en":"Pull stream id, channel unique identification, multiple ids separated by \",\";", "zh_CN":"拉流 id，频道唯一标识，多个 ID 用\",\"隔开；"}
  PullId *string `json:"pullId,omitempty" xml:"pullId,omitempty" require:"true"`
}

func (s ChannelForbiddenRequest) String() string {
  return tea.Prettify(s)
}

func (s ChannelForbiddenRequest) GoString() string {
  return s.String()
}

func (s *ChannelForbiddenRequest) SetPullId(v string) *ChannelForbiddenRequest {
  s.PullId = &v
  return s
}

type ChannelForbiddenResponse struct {
  // {"en":"status code", "zh_CN":"状态码"}
  Code *int32 `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"message", "zh_CN":"操作信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"return data", "zh_CN":"返回数据"}
  Data *string `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s ChannelForbiddenResponse) String() string {
  return tea.Prettify(s)
}

func (s ChannelForbiddenResponse) GoString() string {
  return s.String()
}

func (s *ChannelForbiddenResponse) SetCode(v int32) *ChannelForbiddenResponse {
  s.Code = &v
  return s
}

func (s *ChannelForbiddenResponse) SetMessage(v string) *ChannelForbiddenResponse {
  s.Message = &v
  return s
}

func (s *ChannelForbiddenResponse) SetData(v string) *ChannelForbiddenResponse {
  s.Data = &v
  return s
}

type ChannelForbiddenPaths struct {
}

func (s ChannelForbiddenPaths) String() string {
  return tea.Prettify(s)
}

func (s ChannelForbiddenPaths) GoString() string {
  return s.String()
}

type ChannelForbiddenParameters struct {
}

func (s ChannelForbiddenParameters) String() string {
  return tea.Prettify(s)
}

func (s ChannelForbiddenParameters) GoString() string {
  return s.String()
}

type ChannelForbiddenRequestHeader struct {
}

func (s ChannelForbiddenRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ChannelForbiddenRequestHeader) GoString() string {
  return s.String()
}

type ChannelForbiddenResponseHeader struct {
}

func (s ChannelForbiddenResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ChannelForbiddenResponseHeader) GoString() string {
  return s.String()
}




type GetPullChannelListRequest struct {
  // {"en":"Gets the sorting of the channel list; 0. In descending order by creation time; 1. In ascending order by creation time. Default is 0", "zh_CN":"获取频道列表的排序；0、按创建时间降序，1、按创建时间升序。默认为 0"}
  ListOrder *int32 `json:"listOrder,omitempty" xml:"listOrder,omitempty"`
  // {"en":"The number of pages in the channel page list, starting from 1; Default is 1", "zh_CN":"频道分页列表中第几页，从 1 开始取值；默认为 1"}
  PageIndex *int32 `json:"pageIndex,omitempty" xml:"pageIndex,omitempty"`
  // {"en":"Average number of channels per page, ranging from 1 to 50; Default is 10", "zh_CN":"平均每页频道数量，取值在 1-50 之间；默认为 10"}
  PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s GetPullChannelListRequest) String() string {
  return tea.Prettify(s)
}

func (s GetPullChannelListRequest) GoString() string {
  return s.String()
}

func (s *GetPullChannelListRequest) SetListOrder(v int32) *GetPullChannelListRequest {
  s.ListOrder = &v
  return s
}

func (s *GetPullChannelListRequest) SetPageIndex(v int32) *GetPullChannelListRequest {
  s.PageIndex = &v
  return s
}

func (s *GetPullChannelListRequest) SetPageSize(v int32) *GetPullChannelListRequest {
  s.PageSize = &v
  return s
}

type GetPullChannelListResponse struct {
  // {"en":"200 success", "zh_CN":"200为成功"}
  Code *int32 `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Successful operation", "zh_CN":"操作成功"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Return data", "zh_CN":"返回数据"}
  GetPullChannelListData *GetPullChannelListData `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s GetPullChannelListResponse) String() string {
  return tea.Prettify(s)
}

func (s GetPullChannelListResponse) GoString() string {
  return s.String()
}

func (s *GetPullChannelListResponse) SetCode(v int32) *GetPullChannelListResponse {
  s.Code = &v
  return s
}

func (s *GetPullChannelListResponse) SetMessage(v string) *GetPullChannelListResponse {
  s.Message = &v
  return s
}

func (s *GetPullChannelListResponse) SetData(v *GetPullChannelListData) *GetPullChannelListResponse {
  s.GetPullChannelListData = v
  return s
}

type GetPullChannelListData struct {
  // {"en":"Number of channels", "zh_CN":"频道总数"}
  ChannelToal *string `json:"channelToal,omitempty" xml:"channelToal,omitempty" require:"true"`
  // {"en":"Channel list array", "zh_CN":"频道列表数组"}
  ChannelList []*GetPullChannelListChannelListItem `json:"channelList,omitempty" xml:"channelList,omitempty" require:"true" type:"Repeated"`
}

func (s GetPullChannelListData) String() string {
  return tea.Prettify(s)
}

func (s GetPullChannelListData) GoString() string {
  return s.String()
}

func (s *GetPullChannelListData) SetChannelToal(v string) *GetPullChannelListData {
  s.ChannelToal = &v
  return s
}

func (s *GetPullChannelListData) SetChannelList(v []*GetPullChannelListChannelListItem) *GetPullChannelListData {
  s.ChannelList = v
  return s
}

type GetPullChannelListChannelListItem struct {
  // {"en":"channelName", "zh_CN":"频道名称"}
  ChannelName *string `json:"channelName,omitempty" xml:"channelName,omitempty" require:"true"`
  // {"en":"channelType", "zh_CN":"频道类型"}
  ChannelType *string `json:"channelType,omitempty" xml:"channelType,omitempty" require:"true"`
  // {"en":"Sources-end pull-through protocols: RTMP, Http flv, RTMP, Http flv, HLS, HDS, TS", "zh_CN":"源端拉流协议，RTMP、Http flv、RTMP、Http flv、HLS、HDS、TS"}
  SourcePullProtocol *string `json:"sourcePullProtocol,omitempty" xml:"sourcePullProtocol,omitempty" require:"true"`
  // {"en":"sourcePullUrl", "zh_CN":"源端拉流 url"}
  SourcePullUrl *string `json:"sourcePullUrl,omitempty" xml:"sourcePullUrl,omitempty" require:"true"`
  // {"en":"createTime", "zh_CN":"频道创建时间，例如 2015-05-07 12:00:00"}
  CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty" require:"true"`
  // {"en":"userId", "zh_CN":"创建频道的用户id"}
  UserId *string `json:"userId,omitempty" xml:"userId,omitempty" require:"true"`
  // {"en":"pullId", "zh_CN":"拉流id"}
  PullId *string `json:"pullId,omitempty" xml:"pullId,omitempty" require:"true"`
}

func (s GetPullChannelListChannelListItem) String() string {
  return tea.Prettify(s)
}

func (s GetPullChannelListChannelListItem) GoString() string {
  return s.String()
}

func (s *GetPullChannelListChannelListItem) SetChannelName(v string) *GetPullChannelListChannelListItem {
  s.ChannelName = &v
  return s
}

func (s *GetPullChannelListChannelListItem) SetChannelType(v string) *GetPullChannelListChannelListItem {
  s.ChannelType = &v
  return s
}

func (s *GetPullChannelListChannelListItem) SetSourcePullProtocol(v string) *GetPullChannelListChannelListItem {
  s.SourcePullProtocol = &v
  return s
}

func (s *GetPullChannelListChannelListItem) SetSourcePullUrl(v string) *GetPullChannelListChannelListItem {
  s.SourcePullUrl = &v
  return s
}

func (s *GetPullChannelListChannelListItem) SetCreateTime(v string) *GetPullChannelListChannelListItem {
  s.CreateTime = &v
  return s
}

func (s *GetPullChannelListChannelListItem) SetUserId(v string) *GetPullChannelListChannelListItem {
  s.UserId = &v
  return s
}

func (s *GetPullChannelListChannelListItem) SetPullId(v string) *GetPullChannelListChannelListItem {
  s.PullId = &v
  return s
}

type GetPullChannelListPaths struct {
}

func (s GetPullChannelListPaths) String() string {
  return tea.Prettify(s)
}

func (s GetPullChannelListPaths) GoString() string {
  return s.String()
}

type GetPullChannelListParameters struct {
}

func (s GetPullChannelListParameters) String() string {
  return tea.Prettify(s)
}

func (s GetPullChannelListParameters) GoString() string {
  return s.String()
}

type GetPullChannelListRequestHeader struct {
}

func (s GetPullChannelListRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetPullChannelListRequestHeader) GoString() string {
  return s.String()
}

type GetPullChannelListResponseHeader struct {
}

func (s GetPullChannelListResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetPullChannelListResponseHeader) GoString() string {
  return s.String()
}




type GetPullChannelConfigureRequest struct {
  // {"en":"Pull stream id, unique ID of the channel", "zh_CN":"拉流 id，频道的唯一 ID"}
  PullId *string `json:"pullId,omitempty" xml:"pullId,omitempty" require:"true"`
}

func (s GetPullChannelConfigureRequest) String() string {
  return tea.Prettify(s)
}

func (s GetPullChannelConfigureRequest) GoString() string {
  return s.String()
}

func (s *GetPullChannelConfigureRequest) SetPullId(v string) *GetPullChannelConfigureRequest {
  s.PullId = &v
  return s
}

type GetPullChannelConfigureResponse struct {
  // {"en":"200 success", "zh_CN":"200为成功"}
  Code *int32 `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Successful operation", "zh_CN":"操作成功"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Return data", "zh_CN":"返回数据"}
  GetPullChannelConfigureData *GetPullChannelConfigureData `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s GetPullChannelConfigureResponse) String() string {
  return tea.Prettify(s)
}

func (s GetPullChannelConfigureResponse) GoString() string {
  return s.String()
}

func (s *GetPullChannelConfigureResponse) SetCode(v int32) *GetPullChannelConfigureResponse {
  s.Code = &v
  return s
}

func (s *GetPullChannelConfigureResponse) SetMessage(v string) *GetPullChannelConfigureResponse {
  s.Message = &v
  return s
}

func (s *GetPullChannelConfigureResponse) SetData(v *GetPullChannelConfigureData) *GetPullChannelConfigureResponse {
  s.GetPullChannelConfigureData = v
  return s
}

type GetPullChannelConfigureData struct {
  // {"en":"Channel name, maximum 40 Chinese characters", "zh_CN":"频道名称，最长 40 中文"}
  ChannelName *string `json:"channelName,omitempty" xml:"channelName,omitempty" require:"true"`
  // {"en":"Keep the field and always return it", "zh_CN":"保留字段，始终返回"}
  ChannelStatus *string `json:"channelStatus,omitempty" xml:"channelStatus,omitempty" require:"true"`
  // {"en":"Source channel creation time, for example, 2016-05-08 12:00:00", "zh_CN":"源频道创建时间，例如：2016-05-08 12:00:00"}
  CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty" require:"true"`
  // {"en":"createUser", "zh_CN":"创建频道的 userName"}
  CreateUser *string `json:"createUser,omitempty" xml:"createUser,omitempty" require:"true"`
  // {"en":"flash player ID used by the channel. If not set, the default player is used", "zh_CN":"频道使用的 flash 播放器 ID，未设置则使用默认播放器"}
  PlayerId *string `json:"playerId,omitempty" xml:"playerId,omitempty" require:"true"`
  // {"en":"Source pull channel configuration information", "zh_CN":"源拉流频道配置信息"}
  SourcePullConfigureInfoList []*GetPullChannelConfigureSourcePullConfigureInfoListItem `json:"sourcePullConfigureInfoList,omitempty" xml:"sourcePullConfigureInfoList,omitempty" require:"true" type:"Repeated"`
}

func (s GetPullChannelConfigureData) String() string {
  return tea.Prettify(s)
}

func (s GetPullChannelConfigureData) GoString() string {
  return s.String()
}

func (s *GetPullChannelConfigureData) SetChannelName(v string) *GetPullChannelConfigureData {
  s.ChannelName = &v
  return s
}

func (s *GetPullChannelConfigureData) SetChannelStatus(v string) *GetPullChannelConfigureData {
  s.ChannelStatus = &v
  return s
}

func (s *GetPullChannelConfigureData) SetCreateTime(v string) *GetPullChannelConfigureData {
  s.CreateTime = &v
  return s
}

func (s *GetPullChannelConfigureData) SetCreateUser(v string) *GetPullChannelConfigureData {
  s.CreateUser = &v
  return s
}

func (s *GetPullChannelConfigureData) SetPlayerId(v string) *GetPullChannelConfigureData {
  s.PlayerId = &v
  return s
}

func (s *GetPullChannelConfigureData) SetSourcePullConfigureInfoList(v []*GetPullChannelConfigureSourcePullConfigureInfoListItem) *GetPullChannelConfigureData {
  s.SourcePullConfigureInfoList = v
  return s
}

type GetPullChannelConfigureSourcePullConfigureInfoListItem struct {
  // {"en":"Source pull url The system pulls the live stream as the live stream source", "zh_CN":"源端拉流 url系统会拉取该直播流作为直播源"}
  SourcePullUrl *string `json:"sourcePullUrl,omitempty" xml:"sourcePullUrl,omitempty" require:"true"`
  // {"en":"Source pull protocol: A channel can have only one back pull protocol, and a source pull basin name can be used for only one source pull protocol channel. 0, RTMP, 1, Http flv", "zh_CN":"源端拉流协议，一个频道只能有一种回源拉流协议，且一个源拉流域名只能用于一种源拉流协议频道；0、RTMP，1、Http flv"}
  SourcePullProtocol *string `json:"sourcePullProtocol,omitempty" xml:"sourcePullProtocol,omitempty" require:"true"`
  // {"en":"Player pull-down configuration information list (pull-down protocol, pull-down url)", "zh_CN":"播放端拉流配置信息列表（拉流协议，拉流 url）"}
  PullConfigureList []*GetPullChannelConfigurePullConfigureListItem `json:"pullConfigureList,omitempty" xml:"pullConfigureList,omitempty" require:"true" type:"Repeated"`
}

func (s GetPullChannelConfigureSourcePullConfigureInfoListItem) String() string {
  return tea.Prettify(s)
}

func (s GetPullChannelConfigureSourcePullConfigureInfoListItem) GoString() string {
  return s.String()
}

func (s *GetPullChannelConfigureSourcePullConfigureInfoListItem) SetSourcePullUrl(v string) *GetPullChannelConfigureSourcePullConfigureInfoListItem {
  s.SourcePullUrl = &v
  return s
}

func (s *GetPullChannelConfigureSourcePullConfigureInfoListItem) SetSourcePullProtocol(v string) *GetPullChannelConfigureSourcePullConfigureInfoListItem {
  s.SourcePullProtocol = &v
  return s
}

func (s *GetPullChannelConfigureSourcePullConfigureInfoListItem) SetPullConfigureList(v []*GetPullChannelConfigurePullConfigureListItem) *GetPullChannelConfigureSourcePullConfigureInfoListItem {
  s.PullConfigureList = v
  return s
}

type GetPullChannelConfigurePullConfigureListItem struct {
  // {"en":"Player pull stream protocol", "zh_CN":"播放端拉流协议"}
  PullProtocol *string `json:"pullProtocol,omitempty" xml:"pullProtocol,omitempty" require:"true"`
  // {"en":"pullUrlList", "zh_CN":"拉流 url 列表"}
  PullUrlList []*GetPullChannelConfigurePullUrlListItem `json:"pullUrlList,omitempty" xml:"pullUrlList,omitempty" require:"true" type:"Repeated"`
}

func (s GetPullChannelConfigurePullConfigureListItem) String() string {
  return tea.Prettify(s)
}

func (s GetPullChannelConfigurePullConfigureListItem) GoString() string {
  return s.String()
}

func (s *GetPullChannelConfigurePullConfigureListItem) SetPullProtocol(v string) *GetPullChannelConfigurePullConfigureListItem {
  s.PullProtocol = &v
  return s
}

func (s *GetPullChannelConfigurePullConfigureListItem) SetPullUrlList(v []*GetPullChannelConfigurePullUrlListItem) *GetPullChannelConfigurePullConfigureListItem {
  s.PullUrlList = v
  return s
}

type GetPullChannelConfigurePullUrlListItem struct {
  // {"en":"originPullUrl", "zh_CN":"源码率拉流url"}
  OriginPullUrl *string `json:"originPullUrl,omitempty" xml:"originPullUrl,omitempty" require:"true"`
  // {"en":"fluentPullUrl", "zh_CN":"流畅码率拉流url"}
  FluentPullUrl *string `json:"fluentPullUrl,omitempty" xml:"fluentPullUrl,omitempty" require:"true"`
  // {"en":"sdPullUrl", "zh_CN":"标清码率拉流url"}
  SdPullUrl *string `json:"sdPullUrl,omitempty" xml:"sdPullUrl,omitempty" require:"true"`
  // {"en":"highPullUrl", "zh_CN":"高清码率拉流url"}
  HighPullUrl *string `json:"highPullUrl,omitempty" xml:"highPullUrl,omitempty" require:"true"`
  // {"en":"hdPullUrl", "zh_CN":"超清码率拉流url"}
  HdPullUrl *string `json:"hdPullUrl,omitempty" xml:"hdPullUrl,omitempty" require:"true"`
}

func (s GetPullChannelConfigurePullUrlListItem) String() string {
  return tea.Prettify(s)
}

func (s GetPullChannelConfigurePullUrlListItem) GoString() string {
  return s.String()
}

func (s *GetPullChannelConfigurePullUrlListItem) SetOriginPullUrl(v string) *GetPullChannelConfigurePullUrlListItem {
  s.OriginPullUrl = &v
  return s
}

func (s *GetPullChannelConfigurePullUrlListItem) SetFluentPullUrl(v string) *GetPullChannelConfigurePullUrlListItem {
  s.FluentPullUrl = &v
  return s
}

func (s *GetPullChannelConfigurePullUrlListItem) SetSdPullUrl(v string) *GetPullChannelConfigurePullUrlListItem {
  s.SdPullUrl = &v
  return s
}

func (s *GetPullChannelConfigurePullUrlListItem) SetHighPullUrl(v string) *GetPullChannelConfigurePullUrlListItem {
  s.HighPullUrl = &v
  return s
}

func (s *GetPullChannelConfigurePullUrlListItem) SetHdPullUrl(v string) *GetPullChannelConfigurePullUrlListItem {
  s.HdPullUrl = &v
  return s
}

type GetPullChannelConfigurePaths struct {
}

func (s GetPullChannelConfigurePaths) String() string {
  return tea.Prettify(s)
}

func (s GetPullChannelConfigurePaths) GoString() string {
  return s.String()
}

type GetPullChannelConfigureParameters struct {
}

func (s GetPullChannelConfigureParameters) String() string {
  return tea.Prettify(s)
}

func (s GetPullChannelConfigureParameters) GoString() string {
  return s.String()
}

type GetPullChannelConfigureRequestHeader struct {
}

func (s GetPullChannelConfigureRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetPullChannelConfigureRequestHeader) GoString() string {
  return s.String()
}

type GetPullChannelConfigureResponseHeader struct {
}

func (s GetPullChannelConfigureResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetPullChannelConfigureResponseHeader) GoString() string {
  return s.String()
}




type EditPullChannelRequest struct {
  // {"en":"Channel pull id", "zh_CN":"频道拉流 id"}
  PullId *string `json:"pullId,omitempty" xml:"pullId,omitempty" require:"true"`
  // {"en":"Channel name, maximum 255 characters; Unfilled default unchanged", "zh_CN":"频道名称，最长255字符；不填默认不变"}
  ChannelName *string `json:"channelName,omitempty" xml:"channelName,omitempty"`
  // {"en":"Source pull url", "zh_CN":"源端拉流 url"}
  SourcePullUrl *string `json:"sourcePullUrl,omitempty" xml:"sourcePullUrl,omitempty"`
  // {"en":"Source pull protocol: A channel can have only one back pull protocol, and a source pull basin name can be used for only one source pull protocol channel. 0, RTMP, 1, Http flv, 2, HLS, 3, HDS", "zh_CN":"源端拉流协议，一个频道只能有一种回源拉流协议，且一个源拉流域名只能用于一种源拉流协议频道；0、RTMP，1、Http flv，2、HLS，3、HDS"}
  SourcePullProtocol *string `json:"sourcePullProtocol,omitempty" xml:"sourcePullProtocol,omitempty"`
  // {"en":"The player can select multiple bit rates, separated by commas (,), and ensure that the selected bit rates are within the range of the transcoding bit rates configured by the basin name. 0, source rate, 1, smooth, 2, standard definition, 3, HD, 4, super clear;", "zh_CN":"播放端拉流码率，可多选，用“，”隔开，确保选择的码率包含在推流域名配置的转码码率范围内；0、源码率，1、流畅，2、标清，3、高清，4、超清；"}
  PullRate *string `json:"pullRate,omitempty" xml:"pullRate,omitempty"`
}

func (s EditPullChannelRequest) String() string {
  return tea.Prettify(s)
}

func (s EditPullChannelRequest) GoString() string {
  return s.String()
}

func (s *EditPullChannelRequest) SetPullId(v string) *EditPullChannelRequest {
  s.PullId = &v
  return s
}

func (s *EditPullChannelRequest) SetChannelName(v string) *EditPullChannelRequest {
  s.ChannelName = &v
  return s
}

func (s *EditPullChannelRequest) SetSourcePullUrl(v string) *EditPullChannelRequest {
  s.SourcePullUrl = &v
  return s
}

func (s *EditPullChannelRequest) SetSourcePullProtocol(v string) *EditPullChannelRequest {
  s.SourcePullProtocol = &v
  return s
}

func (s *EditPullChannelRequest) SetPullRate(v string) *EditPullChannelRequest {
  s.PullRate = &v
  return s
}

type EditPullChannelResponse struct {
  // {"en":"status code", "zh_CN":"200为成功"}
  Code *int32 `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"message", "zh_CN":"操作成功"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"return data", "zh_CN":"返回数据"}
  EditPullChannelData *EditPullChannelData `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s EditPullChannelResponse) String() string {
  return tea.Prettify(s)
}

func (s EditPullChannelResponse) GoString() string {
  return s.String()
}

func (s *EditPullChannelResponse) SetCode(v int32) *EditPullChannelResponse {
  s.Code = &v
  return s
}

func (s *EditPullChannelResponse) SetMessage(v string) *EditPullChannelResponse {
  s.Message = &v
  return s
}

func (s *EditPullChannelResponse) SetData(v *EditPullChannelData) *EditPullChannelResponse {
  s.EditPullChannelData = v
  return s
}

type EditPullChannelData struct {
}

func (s EditPullChannelData) String() string {
  return tea.Prettify(s)
}

func (s EditPullChannelData) GoString() string {
  return s.String()
}

type EditPullChannelPaths struct {
}

func (s EditPullChannelPaths) String() string {
  return tea.Prettify(s)
}

func (s EditPullChannelPaths) GoString() string {
  return s.String()
}

type EditPullChannelParameters struct {
}

func (s EditPullChannelParameters) String() string {
  return tea.Prettify(s)
}

func (s EditPullChannelParameters) GoString() string {
  return s.String()
}

type EditPullChannelRequestHeader struct {
}

func (s EditPullChannelRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s EditPullChannelRequestHeader) GoString() string {
  return s.String()
}

type EditPullChannelResponseHeader struct {
}

func (s EditPullChannelResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s EditPullChannelResponseHeader) GoString() string {
  return s.String()
}




type GetSinglePullCodeRequest struct {
  // {"en":"Channel pull id", "zh_CN":"频道拉流 id"}
  PullId *string `json:"pullId,omitempty" xml:"pullId,omitempty" require:"true"`
  // {"en":"Enter the player ID configured for the channel, which must be consistent with that maintained in 3. Player Management. Do not enter the default player.", "zh_CN":"填写为该频道配置的播放器 ID，必须与三、播放器管理中维护的一致。不填写就采用默认播放器。"}
  PlayerId *string `json:"playerId,omitempty" xml:"playerId,omitempty"`
  // {"en":"To match the aspect ratio of the amplifier, five fixed options are provided
  // (1280*720, 1024*768, 1024*576, 640*360,
  // 480*360), and a self-defining lattice with a length-to-width ratio greater than
  // 480 times 360, which is less than 1920 times 1080. The default value is 480 x 360", "zh_CN":"配置播放器的长宽比，提供五个固定选项
  // （ 1280*720 、 1024*768 、 1024*576 、 640*360 、
  // 480*360 ），及一个自定义格式 ，长宽比大于
  // 480*360，小于 1920*1080。不填写默认 480*360"}
  PlayerSize *string `json:"playerSize,omitempty" xml:"playerSize,omitempty"`
  // {"en":"Whether the video plays automatically. The default value is 0.
  // 0 indicates non-autoplay, 1 indicates autoplay", "zh_CN":"视频是否自动播放，默认为 0； 
  // 0 为非自动播放，1 为自动播放"}
  AutoPlay *int32 `json:"autoPlay,omitempty" xml:"autoPlay,omitempty"`
  // {"en":"Select the type of playback code you want to get. The default is all code:
  // 0:all, 2:swf code, 4:video url, 5:adaptive code;", "zh_CN":"选择需要获取的播放代码类型，默认为全部代码： 
  // 0、全部，2、swf 代码， 4、视频 url，5、自适应代码；  "}
  CodeType *string `json:"codeType,omitempty" xml:"codeType,omitempty"`
}

func (s GetSinglePullCodeRequest) String() string {
  return tea.Prettify(s)
}

func (s GetSinglePullCodeRequest) GoString() string {
  return s.String()
}

func (s *GetSinglePullCodeRequest) SetPullId(v string) *GetSinglePullCodeRequest {
  s.PullId = &v
  return s
}

func (s *GetSinglePullCodeRequest) SetPlayerId(v string) *GetSinglePullCodeRequest {
  s.PlayerId = &v
  return s
}

func (s *GetSinglePullCodeRequest) SetPlayerSize(v string) *GetSinglePullCodeRequest {
  s.PlayerSize = &v
  return s
}

func (s *GetSinglePullCodeRequest) SetAutoPlay(v int32) *GetSinglePullCodeRequest {
  s.AutoPlay = &v
  return s
}

func (s *GetSinglePullCodeRequest) SetCodeType(v string) *GetSinglePullCodeRequest {
  s.CodeType = &v
  return s
}

type GetSinglePullCodeResponse struct {
  // {'en':'status code', 'zh_CN':'返回状态码'}
  Code *int32 `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  Data *GetSinglePullCodePullCodeInfo `json:"data,omitempty" xml:"data,omitempty" require:"true"`
  // {'en':'message', 'zh_CN':'返回消息'}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s GetSinglePullCodeResponse) String() string {
  return tea.Prettify(s)
}

func (s GetSinglePullCodeResponse) GoString() string {
  return s.String()
}

func (s *GetSinglePullCodeResponse) SetCode(v int32) *GetSinglePullCodeResponse {
  s.Code = &v
  return s
}

func (s *GetSinglePullCodeResponse) SetData(v *GetSinglePullCodePullCodeInfo) *GetSinglePullCodeResponse {
  s.Data = v
  return s
}

func (s *GetSinglePullCodeResponse) SetMessage(v string) *GetSinglePullCodeResponse {
  s.Message = &v
  return s
}

type GetSinglePullCodePullCodeInfo struct {
  // {"en":"Pull flow ID", "zh_CN":"拉流 ID"}
  PullId *string `json:"pullId,omitempty" xml:"pullId,omitempty" require:"true"`
  // {"en":"List of channel pull urls", "zh_CN":"频道拉流 url 列表"}
  GetSinglePullCodeVideoUrl []*GetSinglePullCodeVideoUrl `json:"videoUrl,omitempty" xml:"videoUrl,omitempty" require:"true" type:"Repeated"`
  // {"en":"Video adaptive code, encrypted video is empty", "zh_CN":"视频自适应代码，加密视频为空"}
  AutoCode *string `json:"autoCode,omitempty" xml:"autoCode,omitempty" require:"true"`
  // {"en":"Channel swf code", "zh_CN":"频道 swf 代码"}
  SwfCode *string `json:"swfCode,omitempty" xml:"swfCode,omitempty" require:"true"`
  // {"en":"A list of pull-down urls in channel SRT format", "zh_CN":"频道 SRT 格式的拉流 url 列表"}
  GetSinglePullCodeSrtVideoUrl []*GetSinglePullCodeSrtVideoUrl `json:"srtVideoUrl,omitempty" xml:"srtVideoUrl,omitempty" require:"true" type:"Repeated"`
  // {"en":"Audio adaptive code, encrypted video is empty", "zh_CN":"音频自适应代码，加密视频为空"}
  AudioAutoCode *string `json:"audioAutoCode,omitempty" xml:"audioAutoCode,omitempty" require:"true"`
}

func (s GetSinglePullCodePullCodeInfo) String() string {
  return tea.Prettify(s)
}

func (s GetSinglePullCodePullCodeInfo) GoString() string {
  return s.String()
}

func (s *GetSinglePullCodePullCodeInfo) SetPullId(v string) *GetSinglePullCodePullCodeInfo {
  s.PullId = &v
  return s
}

func (s *GetSinglePullCodePullCodeInfo) SetVideoUrl(v []*GetSinglePullCodeVideoUrl) *GetSinglePullCodePullCodeInfo {
  s.GetSinglePullCodeVideoUrl = v
  return s
}

func (s *GetSinglePullCodePullCodeInfo) SetAutoCode(v string) *GetSinglePullCodePullCodeInfo {
  s.AutoCode = &v
  return s
}

func (s *GetSinglePullCodePullCodeInfo) SetSwfCode(v string) *GetSinglePullCodePullCodeInfo {
  s.SwfCode = &v
  return s
}

func (s *GetSinglePullCodePullCodeInfo) SetSrtVideoUrl(v []*GetSinglePullCodeSrtVideoUrl) *GetSinglePullCodePullCodeInfo {
  s.GetSinglePullCodeSrtVideoUrl = v
  return s
}

func (s *GetSinglePullCodePullCodeInfo) SetAudioAutoCode(v string) *GetSinglePullCodePullCodeInfo {
  s.AudioAutoCode = &v
  return s
}

type GetSinglePullCodeVideoUrl struct {
  // {"en":"Pull-flow protocol, 0, RTMP, 1, Http flv, 2, HLS, 3, HDS, 4, Http TS", "zh_CN":"拉流协议，0、RTMP，1、Http flv，2、HLS，3、HDS，4、Http TS"}
  PullProtocol *string `json:"pullProtocol,omitempty" xml:"pullProtocol,omitempty" require:"true"`
  GetSinglePullCodeUrlInfo *GetSinglePullCodeUrlInfo `json:"urlInfo,omitempty" xml:"urlInfo,omitempty" require:"true"`
}

func (s GetSinglePullCodeVideoUrl) String() string {
  return tea.Prettify(s)
}

func (s GetSinglePullCodeVideoUrl) GoString() string {
  return s.String()
}

func (s *GetSinglePullCodeVideoUrl) SetPullProtocol(v string) *GetSinglePullCodeVideoUrl {
  s.PullProtocol = &v
  return s
}

func (s *GetSinglePullCodeVideoUrl) SetUrlInfo(v *GetSinglePullCodeUrlInfo) *GetSinglePullCodeVideoUrl {
  s.GetSinglePullCodeUrlInfo = v
  return s
}

type GetSinglePullCodeSrtVideoUrl struct {
  // {"en":"Pull-flow protocol, 0, RTMP, 1, Http flv, 2, HLS, 3, HDS, 4, Http TS", "zh_CN":"拉流协议，0、RTMP，1、Http flv，2、HLS，3、HDS，4、Http TS"}
  PullProtocol *string `json:"pullProtocol,omitempty" xml:"pullProtocol,omitempty" require:"true"`
  GetSinglePullCodeUrlInfo *GetSinglePullCodeUrlInfo `json:"urlInfo,omitempty" xml:"urlInfo,omitempty" require:"true"`
}

func (s GetSinglePullCodeSrtVideoUrl) String() string {
  return tea.Prettify(s)
}

func (s GetSinglePullCodeSrtVideoUrl) GoString() string {
  return s.String()
}

func (s *GetSinglePullCodeSrtVideoUrl) SetPullProtocol(v string) *GetSinglePullCodeSrtVideoUrl {
  s.PullProtocol = &v
  return s
}

func (s *GetSinglePullCodeSrtVideoUrl) SetUrlInfo(v *GetSinglePullCodeUrlInfo) *GetSinglePullCodeSrtVideoUrl {
  s.GetSinglePullCodeUrlInfo = v
  return s
}

type GetSinglePullCodeUrlInfo struct {
  // {"en":"Smooth bit rate pull stream url", "zh_CN":"流畅码率拉流 url"}
  FluentPullUrl *string `json:"fluentPullUrl,omitempty" xml:"fluentPullUrl,omitempty" require:"true"`
  // {"en":"Smooth bit rate (smart HD) pull stream url", "zh_CN":"流畅码率（智控高清）拉流 url"}
  FluentZkgqPullUrl *string `json:"fluentZkgqPullUrl,omitempty" xml:"fluentZkgqPullUrl,omitempty" require:"true"`
  // {"en":"Ultra clear bit rate pull url", "zh_CN":"超清码率拉流 url "}
  HdPullUrl *string `json:"hdPullUrl,omitempty" xml:"hdPullUrl,omitempty" require:"true"`
  // {"en":"Ultra - clear bit rate (intelligent control HD) pull - stream url", "zh_CN":"超清码率（智控高清）拉流 url"}
  HdZkgqPullUrl *string `json:"hdZkgqPullUrl,omitempty" xml:"hdZkgqPullUrl,omitempty" require:"true"`
  // {"en":"Hd bit rate pull stream url", "zh_CN":"高清码率拉流 url"}
  HighPullUrl *string `json:"highPullUrl,omitempty" xml:"highPullUrl,omitempty" require:"true"`
  // {"en":"Hd bit rate (intelligent control HD) pull url", "zh_CN":"高清码率（智控高清）拉流 url"}
  HighZkgqPullUrl *string `json:"highZkgqPullUrl,omitempty" xml:"highZkgqPullUrl,omitempty" require:"true"`
  // {"en":"Source rate pull url", "zh_CN":"源码率拉流 url"}
  OriginPullUrl *string `json:"originPullUrl,omitempty" xml:"originPullUrl,omitempty" require:"true"`
  // {"en":"Source rate (intelligent control HD) pull url", "zh_CN":"源码率（智控高清）拉流 url"}
  OriginZkgqPullUrl *string `json:"originZkgqPullUrl,omitempty" xml:"originZkgqPullUrl,omitempty" require:"true"`
  // {"en":"Standard definition bit rate pull url", "zh_CN":"标清码率拉流 url"}
  SdPullUrl *string `json:"sdPullUrl,omitempty" xml:"sdPullUrl,omitempty" require:"true"`
  // {"en":"Standard definition code rate (intelligent control HD) pull url", "zh_CN":"标清码率（智控高清）拉流 url"}
  SdZkgqPullUrl *string `json:"sdZkgqPullUrl,omitempty" xml:"sdZkgqPullUrl,omitempty" require:"true"`
}

func (s GetSinglePullCodeUrlInfo) String() string {
  return tea.Prettify(s)
}

func (s GetSinglePullCodeUrlInfo) GoString() string {
  return s.String()
}

func (s *GetSinglePullCodeUrlInfo) SetFluentPullUrl(v string) *GetSinglePullCodeUrlInfo {
  s.FluentPullUrl = &v
  return s
}

func (s *GetSinglePullCodeUrlInfo) SetFluentZkgqPullUrl(v string) *GetSinglePullCodeUrlInfo {
  s.FluentZkgqPullUrl = &v
  return s
}

func (s *GetSinglePullCodeUrlInfo) SetHdPullUrl(v string) *GetSinglePullCodeUrlInfo {
  s.HdPullUrl = &v
  return s
}

func (s *GetSinglePullCodeUrlInfo) SetHdZkgqPullUrl(v string) *GetSinglePullCodeUrlInfo {
  s.HdZkgqPullUrl = &v
  return s
}

func (s *GetSinglePullCodeUrlInfo) SetHighPullUrl(v string) *GetSinglePullCodeUrlInfo {
  s.HighPullUrl = &v
  return s
}

func (s *GetSinglePullCodeUrlInfo) SetHighZkgqPullUrl(v string) *GetSinglePullCodeUrlInfo {
  s.HighZkgqPullUrl = &v
  return s
}

func (s *GetSinglePullCodeUrlInfo) SetOriginPullUrl(v string) *GetSinglePullCodeUrlInfo {
  s.OriginPullUrl = &v
  return s
}

func (s *GetSinglePullCodeUrlInfo) SetOriginZkgqPullUrl(v string) *GetSinglePullCodeUrlInfo {
  s.OriginZkgqPullUrl = &v
  return s
}

func (s *GetSinglePullCodeUrlInfo) SetSdPullUrl(v string) *GetSinglePullCodeUrlInfo {
  s.SdPullUrl = &v
  return s
}

func (s *GetSinglePullCodeUrlInfo) SetSdZkgqPullUrl(v string) *GetSinglePullCodeUrlInfo {
  s.SdZkgqPullUrl = &v
  return s
}

type GetSinglePullCodePaths struct {
}

func (s GetSinglePullCodePaths) String() string {
  return tea.Prettify(s)
}

func (s GetSinglePullCodePaths) GoString() string {
  return s.String()
}

type GetSinglePullCodeParameters struct {
}

func (s GetSinglePullCodeParameters) String() string {
  return tea.Prettify(s)
}

func (s GetSinglePullCodeParameters) GoString() string {
  return s.String()
}

type GetSinglePullCodeRequestHeader struct {
}

func (s GetSinglePullCodeRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetSinglePullCodeRequestHeader) GoString() string {
  return s.String()
}

type GetSinglePullCodeResponseHeader struct {
}

func (s GetSinglePullCodeResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetSinglePullCodeResponseHeader) GoString() string {
  return s.String()
}




type GetForwardTaskListRequest struct {
  // {"en":"pullId", "zh_CN":"频道 ID"}
  PullId *string `json:"pullId,omitempty" xml:"pullId,omitempty"`
  // {"en":"taskId", "zh_CN":"转推任务 ID"}
  TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
  // {"en":"Task forwarding status, used to filter tasks. Value range: 0: not started. 1: transfer process; 2: has ended; 3: abnormal; Multiple options are supported, separated by hyphens (ids) in English", "zh_CN":"转推任务状态，用于筛选任务，取值范围：0：未开始；1：转推中；2：已结束；3：异常；支持多选，使用英文,号隔开"}
  Status *string `json:"status,omitempty" xml:"status,omitempty"`
  // {"en":"On which page, the value starts from 1. The default value is 1.", "zh_CN":"取第几页，从 1 开始取值,默认为 1。"}
  PageIndex *int32 `json:"pageIndex,omitempty" xml:"pageIndex,omitempty"`
  // {"en":"Number of forwarding tasks per page. The value ranges from 1 to 50. The default value is 10.", "zh_CN":"每页转推任务数量，取值范围 1-50，默认为10。"}
  PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s GetForwardTaskListRequest) String() string {
  return tea.Prettify(s)
}

func (s GetForwardTaskListRequest) GoString() string {
  return s.String()
}

func (s *GetForwardTaskListRequest) SetPullId(v string) *GetForwardTaskListRequest {
  s.PullId = &v
  return s
}

func (s *GetForwardTaskListRequest) SetTaskId(v string) *GetForwardTaskListRequest {
  s.TaskId = &v
  return s
}

func (s *GetForwardTaskListRequest) SetStatus(v string) *GetForwardTaskListRequest {
  s.Status = &v
  return s
}

func (s *GetForwardTaskListRequest) SetPageIndex(v int32) *GetForwardTaskListRequest {
  s.PageIndex = &v
  return s
}

func (s *GetForwardTaskListRequest) SetPageSize(v int32) *GetForwardTaskListRequest {
  s.PageSize = &v
  return s
}

type GetForwardTaskListResponse struct {
  // {"en":"200 success", "zh_CN":"200为成功"}
  Code *int32 `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Successful operation", "zh_CN":"操作成功"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Return data", "zh_CN":"返回数据"}
  GetForwardTaskListData *GetForwardTaskListData `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s GetForwardTaskListResponse) String() string {
  return tea.Prettify(s)
}

func (s GetForwardTaskListResponse) GoString() string {
  return s.String()
}

func (s *GetForwardTaskListResponse) SetCode(v int32) *GetForwardTaskListResponse {
  s.Code = &v
  return s
}

func (s *GetForwardTaskListResponse) SetMessage(v string) *GetForwardTaskListResponse {
  s.Message = &v
  return s
}

func (s *GetForwardTaskListResponse) SetData(v *GetForwardTaskListData) *GetForwardTaskListResponse {
  s.GetForwardTaskListData = v
  return s
}

type GetForwardTaskListData struct {
  // {"en":"The number of task list records that are currently returned", "zh_CN":"当前返回的任务列表信息的记录数"}
  Total *int32 `json:"total,omitempty" xml:"total,omitempty" require:"true"`
  // {"en":"ForwardTaskList information", "zh_CN":"转推任务列表信息"}
  ForwardTastList *GetForwardTaskListForwardTastListItem `json:"forwardTastList,omitempty" xml:"forwardTastList,omitempty" require:"true"`
}

func (s GetForwardTaskListData) String() string {
  return tea.Prettify(s)
}

func (s GetForwardTaskListData) GoString() string {
  return s.String()
}

func (s *GetForwardTaskListData) SetTotal(v int32) *GetForwardTaskListData {
  s.Total = &v
  return s
}

func (s *GetForwardTaskListData) SetForwardTastList(v *GetForwardTaskListForwardTastListItem) *GetForwardTaskListData {
  s.ForwardTastList = v
  return s
}

type GetForwardTaskListForwardTastListItem struct {
  // {"en":"taskId", "zh_CN":"转推任务 ID"}
  TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty" require:"true"`
  // {"en":"pullId", "zh_CN":"频道 ID"}
  PullId *string `json:"pullId,omitempty" xml:"pullId,omitempty" require:"true"`
  // {"en":"Task status, which is used to filter tasks. Value range: 0: not started. 1: transfer process; 2: has ended; 3: Exception", "zh_CN":"任务状态，用于筛选任务，取值范围：0：未开始；1：转推中；2：已结束；3：异常"}
  Status *int32 `json:"status,omitempty" xml:"status,omitempty" require:"true"`
  // {"en":"forwardUrl", "zh_CN":"转推地址"}
  ForwardUrl *string `json:"forwardUrl,omitempty" xml:"forwardUrl,omitempty" require:"true"`
  // {"en":"Start push time, second timestamp", "zh_CN":"开始转推时间,秒级时间戳"}
  StartTime *int32 `json:"startTime,omitempty" xml:"startTime,omitempty" require:"true"`
  // {"en":"End of push time, second timestamp", "zh_CN":"结束转推时间,秒级时间戳"}
  EndTime *int32 `json:"endTime,omitempty" xml:"endTime,omitempty" require:"true"`
  // {"en":"User-defined notification address", "zh_CN":"用户自定义的通知地址，以 http://开头或https://开头。"}
  NotifyUrl *string `json:"notifyUrl,omitempty" xml:"notifyUrl,omitempty" require:"true"`
  // {"en":"createUser", "zh_CN":"创建人"}
  CreateUser *string `json:"createUser,omitempty" xml:"createUser,omitempty" require:"true"`
  // {"en":"crreateTime", "zh_CN":"创建时间"}
  CrreateTime *int32 `json:"crreateTime,omitempty" xml:"crreateTime,omitempty" require:"true"`
  // {"en":"Set the push encoding parameter", "zh_CN":"转推编码参数设置"}
  GetForwardTaskListCodingParams *GetForwardTaskListCodingParams `json:"codingParams,omitempty" xml:"codingParams,omitempty" require:"true"`
}

func (s GetForwardTaskListForwardTastListItem) String() string {
  return tea.Prettify(s)
}

func (s GetForwardTaskListForwardTastListItem) GoString() string {
  return s.String()
}

func (s *GetForwardTaskListForwardTastListItem) SetTaskId(v string) *GetForwardTaskListForwardTastListItem {
  s.TaskId = &v
  return s
}

func (s *GetForwardTaskListForwardTastListItem) SetPullId(v string) *GetForwardTaskListForwardTastListItem {
  s.PullId = &v
  return s
}

func (s *GetForwardTaskListForwardTastListItem) SetStatus(v int32) *GetForwardTaskListForwardTastListItem {
  s.Status = &v
  return s
}

func (s *GetForwardTaskListForwardTastListItem) SetForwardUrl(v string) *GetForwardTaskListForwardTastListItem {
  s.ForwardUrl = &v
  return s
}

func (s *GetForwardTaskListForwardTastListItem) SetStartTime(v int32) *GetForwardTaskListForwardTastListItem {
  s.StartTime = &v
  return s
}

func (s *GetForwardTaskListForwardTastListItem) SetEndTime(v int32) *GetForwardTaskListForwardTastListItem {
  s.EndTime = &v
  return s
}

func (s *GetForwardTaskListForwardTastListItem) SetNotifyUrl(v string) *GetForwardTaskListForwardTastListItem {
  s.NotifyUrl = &v
  return s
}

func (s *GetForwardTaskListForwardTastListItem) SetCreateUser(v string) *GetForwardTaskListForwardTastListItem {
  s.CreateUser = &v
  return s
}

func (s *GetForwardTaskListForwardTastListItem) SetCrreateTime(v int32) *GetForwardTaskListForwardTastListItem {
  s.CrreateTime = &v
  return s
}

func (s *GetForwardTaskListForwardTastListItem) SetCodingParams(v *GetForwardTaskListCodingParams) *GetForwardTaskListForwardTastListItem {
  s.GetForwardTaskListCodingParams = v
  return s
}

type GetForwardTaskListCodingParams struct {
  // {"en":"Bit rate, in kbps, for example, 1200", "zh_CN":"码率，单位 kbps，例如 1200"}
  Bitrate *int32 `json:"bitrate,omitempty" xml:"bitrate,omitempty" require:"true"`
  // {"en":"Resolution, for example: 420x720 (with a lowercase x in the middle)", "zh_CN":"分辨率，例如：420x720 (中间是小写字母的 x)"}
  Resolution *string `json:"resolution,omitempty" xml:"resolution,omitempty" require:"true"`
  // {"en":"Frame rate, for example, 25", "zh_CN":"帧率，例如：25"}
  FrameRate *int32 `json:"frameRate,omitempty" xml:"frameRate,omitempty" require:"true"`
  // {"en":"libx264, supported schemes: libx264, libx265, libvpx, etc. Also need to add watermark fixed fill in libx264", "zh_CN":"libx264，支持方案：libx264，libx265，libvpx 等。同时需添加水印时固定填写为libx264。"}
  Vcodec *string `json:"vcodec,omitempty" xml:"vcodec,omitempty" require:"true"`
  // {"en":"libfaac, supported schemes: libmp3lame, libfaac, libvorbis, etc", "zh_CN":"libfaac，支持方案：libmp3lame，libfaac，libvorbis 等。"}
  Acodec *string `json:"acodec,omitempty" xml:"acodec,omitempty" require:"true"`
  // {"en":"Watermark image URL address: If the url contains &, url escape is required", "zh_CN":"水印图片的 URL 地址：如果 url 中带&，则需要进行 url 转义"}
  Wmimage *string `json:"wmimage,omitempty" xml:"wmimage,omitempty" require:"true"`
  // {"en":"Watermark position", "zh_CN":"水印位置，TOP_LEFT 左上角；TOP_CENTER 上部居中；TOP_RIGHT 右上角；CENTER_LEFT 中部靠左；CENTER 居中；CENTER_RIGHT 中部靠右；BOTTOM_LEFT 左下角；BOTTOM_CENTER 下居中；BOTTOM_RIGHT 右下角。"}
  Wmposition *string `json:"wmposition,omitempty" xml:"wmposition,omitempty" require:"true"`
}

func (s GetForwardTaskListCodingParams) String() string {
  return tea.Prettify(s)
}

func (s GetForwardTaskListCodingParams) GoString() string {
  return s.String()
}

func (s *GetForwardTaskListCodingParams) SetBitrate(v int32) *GetForwardTaskListCodingParams {
  s.Bitrate = &v
  return s
}

func (s *GetForwardTaskListCodingParams) SetResolution(v string) *GetForwardTaskListCodingParams {
  s.Resolution = &v
  return s
}

func (s *GetForwardTaskListCodingParams) SetFrameRate(v int32) *GetForwardTaskListCodingParams {
  s.FrameRate = &v
  return s
}

func (s *GetForwardTaskListCodingParams) SetVcodec(v string) *GetForwardTaskListCodingParams {
  s.Vcodec = &v
  return s
}

func (s *GetForwardTaskListCodingParams) SetAcodec(v string) *GetForwardTaskListCodingParams {
  s.Acodec = &v
  return s
}

func (s *GetForwardTaskListCodingParams) SetWmimage(v string) *GetForwardTaskListCodingParams {
  s.Wmimage = &v
  return s
}

func (s *GetForwardTaskListCodingParams) SetWmposition(v string) *GetForwardTaskListCodingParams {
  s.Wmposition = &v
  return s
}

type GetForwardTaskListPaths struct {
}

func (s GetForwardTaskListPaths) String() string {
  return tea.Prettify(s)
}

func (s GetForwardTaskListPaths) GoString() string {
  return s.String()
}

type GetForwardTaskListParameters struct {
}

func (s GetForwardTaskListParameters) String() string {
  return tea.Prettify(s)
}

func (s GetForwardTaskListParameters) GoString() string {
  return s.String()
}

type GetForwardTaskListRequestHeader struct {
}

func (s GetForwardTaskListRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetForwardTaskListRequestHeader) GoString() string {
  return s.String()
}

type GetForwardTaskListResponseHeader struct {
}

func (s GetForwardTaskListResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetForwardTaskListResponseHeader) GoString() string {
  return s.String()
}




type DisconnectStreamRequest struct {
  // {"en":"Channel pull ID, unique identifier of the channel", "zh_CN":"频道拉流 ID，频道唯一标识"}
  PullId *string `json:"pullId,omitempty" xml:"pullId,omitempty" require:"true"`
}

func (s DisconnectStreamRequest) String() string {
  return tea.Prettify(s)
}

func (s DisconnectStreamRequest) GoString() string {
  return s.String()
}

func (s *DisconnectStreamRequest) SetPullId(v string) *DisconnectStreamRequest {
  s.PullId = &v
  return s
}

type DisconnectStreamResponse struct {
  // {"en":"200 success", "zh_CN":"200为成功"}
  Code *int32 `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Successful operation", "zh_CN":"操作成功"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Return data", "zh_CN":"返回数据"}
  Data *string `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s DisconnectStreamResponse) String() string {
  return tea.Prettify(s)
}

func (s DisconnectStreamResponse) GoString() string {
  return s.String()
}

func (s *DisconnectStreamResponse) SetCode(v int32) *DisconnectStreamResponse {
  s.Code = &v
  return s
}

func (s *DisconnectStreamResponse) SetMessage(v string) *DisconnectStreamResponse {
  s.Message = &v
  return s
}

func (s *DisconnectStreamResponse) SetData(v string) *DisconnectStreamResponse {
  s.Data = &v
  return s
}

type DisconnectStreamPaths struct {
}

func (s DisconnectStreamPaths) String() string {
  return tea.Prettify(s)
}

func (s DisconnectStreamPaths) GoString() string {
  return s.String()
}

type DisconnectStreamParameters struct {
}

func (s DisconnectStreamParameters) String() string {
  return tea.Prettify(s)
}

func (s DisconnectStreamParameters) GoString() string {
  return s.String()
}

type DisconnectStreamRequestHeader struct {
}

func (s DisconnectStreamRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s DisconnectStreamRequestHeader) GoString() string {
  return s.String()
}

type DisconnectStreamResponseHeader struct {
}

func (s DisconnectStreamResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s DisconnectStreamResponseHeader) GoString() string {
  return s.String()
}




type EndRecordTaskRequest struct {
  // {"en":"id of the live recording task", "zh_CN":"直播录制任务id"}
  TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty" require:"true"`
}

func (s EndRecordTaskRequest) String() string {
  return tea.Prettify(s)
}

func (s EndRecordTaskRequest) GoString() string {
  return s.String()
}

func (s *EndRecordTaskRequest) SetTaskId(v string) *EndRecordTaskRequest {
  s.TaskId = &v
  return s
}

type EndRecordTaskResponse struct {
  // {"en":"200 success", "zh_CN":"200为成功"}
  Code *int32 `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Successful operation", "zh_CN":"操作成功"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Return data", "zh_CN":"返回数据"}
  EndRecordTaskData *EndRecordTaskData `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s EndRecordTaskResponse) String() string {
  return tea.Prettify(s)
}

func (s EndRecordTaskResponse) GoString() string {
  return s.String()
}

func (s *EndRecordTaskResponse) SetCode(v int32) *EndRecordTaskResponse {
  s.Code = &v
  return s
}

func (s *EndRecordTaskResponse) SetMessage(v string) *EndRecordTaskResponse {
  s.Message = &v
  return s
}

func (s *EndRecordTaskResponse) SetData(v *EndRecordTaskData) *EndRecordTaskResponse {
  s.EndRecordTaskData = v
  return s
}

type EndRecordTaskData struct {
}

func (s EndRecordTaskData) String() string {
  return tea.Prettify(s)
}

func (s EndRecordTaskData) GoString() string {
  return s.String()
}

type EndRecordTaskPaths struct {
}

func (s EndRecordTaskPaths) String() string {
  return tea.Prettify(s)
}

func (s EndRecordTaskPaths) GoString() string {
  return s.String()
}

type EndRecordTaskParameters struct {
}

func (s EndRecordTaskParameters) String() string {
  return tea.Prettify(s)
}

func (s EndRecordTaskParameters) GoString() string {
  return s.String()
}

type EndRecordTaskRequestHeader struct {
}

func (s EndRecordTaskRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s EndRecordTaskRequestHeader) GoString() string {
  return s.String()
}

type EndRecordTaskResponseHeader struct {
}

func (s EndRecordTaskResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s EndRecordTaskResponseHeader) GoString() string {
  return s.String()
}




type StartForwardTaskRequest struct {
  // {"en":"pullId", "zh_CN":"频道 ID"}
  PullId *string `json:"pullId,omitempty" xml:"pullId,omitempty" require:"true"`
  // {"en":"Indicates the address to be forwarded. Only the rtmp format is supported.", "zh_CN":"需要转推的地址，只支持 rtmp 格式推流。"}
  ForwardUrl *string `json:"forwardUrl,omitempty" xml:"forwardUrl,omitempty" require:"true"`
  // {"en":"Scheduled start time, timestamp in seconds. If the start time is less than the current time, push immediately from the current time.", "zh_CN":"计划开始时间，秒级时间戳，开始小于当前时间时，则从当前时间立即开始转推。"}
  StartTime *int32 `json:"startTime,omitempty" xml:"startTime,omitempty"`
  // {"en":"The scheduled end time is a timestamp in seconds. The end time must be more than 5 minutes later than the start time and more than 5 minutes later than the current time.", "zh_CN":"计划结束时间，秒级时间戳，结束时间必须晚于开始时间 5 分钟以上，并且晚于当前时间 5 分钟以上。"}
  EndTime *int32 `json:"endTime,omitempty" xml:"endTime,omitempty"`
  // {"en":"jsonType", "zh_CN":"此字段为json类型，字段如下：
  //   bitrate int 非必填 码率，单位 kbps，例如 800
  //   resolution string 非必填 分辨率，例如：420x720 (中间是小写字母的“x”)
  //   frameRate int 非必填 帧率，例如：25
  //   vcodec string 非必填 libx264，支持方案：libx264，libx265，libvpx 等。同时需添加水印时固定填写为libx264
  //   acodec string 非必填 libfaac，支持方案：libmp3lame，libfaac，libvorbis 等
  //   wmimage string 非必填 水印图片的 URL 地址：如果 url 中带&，则需要进行 url 转义。wmimage 参数和wmposition 参数必须同时都传或同时不传
  //   wmposition string 非必填 水印位置。wmimage 参数和 wmposition 参数必须同时都传或同时不传。TOP_LEFT 左上角；TOP_CENTER 上部居中；TOP_RIGHT 右上角；CENTER_LEFT 中部靠左；CENTER 居中；CENTER_RIGHT 中部靠右；BOTTOM_LEFT 左下角；BOTTOM_CENTER 下居中；BOTTOM_RIGHT 右下角 
  //   "}
  CodingParams *string `json:"codingParams,omitempty" xml:"codingParams,omitempty"`
  // {"en":"User-defined notification address", "zh_CN":"用户自定义的通知地址，必须以 http://开头或 https://开头。"}
  NotifyUrl *string `json:"notifyUrl,omitempty" xml:"notifyUrl,omitempty"`
}

func (s StartForwardTaskRequest) String() string {
  return tea.Prettify(s)
}

func (s StartForwardTaskRequest) GoString() string {
  return s.String()
}

func (s *StartForwardTaskRequest) SetPullId(v string) *StartForwardTaskRequest {
  s.PullId = &v
  return s
}

func (s *StartForwardTaskRequest) SetForwardUrl(v string) *StartForwardTaskRequest {
  s.ForwardUrl = &v
  return s
}

func (s *StartForwardTaskRequest) SetStartTime(v int32) *StartForwardTaskRequest {
  s.StartTime = &v
  return s
}

func (s *StartForwardTaskRequest) SetEndTime(v int32) *StartForwardTaskRequest {
  s.EndTime = &v
  return s
}

func (s *StartForwardTaskRequest) SetCodingParams(v string) *StartForwardTaskRequest {
  s.CodingParams = &v
  return s
}

func (s *StartForwardTaskRequest) SetNotifyUrl(v string) *StartForwardTaskRequest {
  s.NotifyUrl = &v
  return s
}

type StartForwardTaskResponse struct {
  // {"en":"200 success", "zh_CN":"200为成功"}
  Code *int32 `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Successful operation", "zh_CN":"操作成功"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Return data", "zh_CN":"返回数据"}
  StartForwardTaskData *StartForwardTaskData `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s StartForwardTaskResponse) String() string {
  return tea.Prettify(s)
}

func (s StartForwardTaskResponse) GoString() string {
  return s.String()
}

func (s *StartForwardTaskResponse) SetCode(v int32) *StartForwardTaskResponse {
  s.Code = &v
  return s
}

func (s *StartForwardTaskResponse) SetMessage(v string) *StartForwardTaskResponse {
  s.Message = &v
  return s
}

func (s *StartForwardTaskResponse) SetData(v *StartForwardTaskData) *StartForwardTaskResponse {
  s.StartForwardTaskData = v
  return s
}

type StartForwardTaskData struct {
  // {"en":"taskId", "zh_CN":"转推任务 ID"}
  TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty" require:"true"`
}

func (s StartForwardTaskData) String() string {
  return tea.Prettify(s)
}

func (s StartForwardTaskData) GoString() string {
  return s.String()
}

func (s *StartForwardTaskData) SetTaskId(v string) *StartForwardTaskData {
  s.TaskId = &v
  return s
}

type StartForwardTaskPaths struct {
}

func (s StartForwardTaskPaths) String() string {
  return tea.Prettify(s)
}

func (s StartForwardTaskPaths) GoString() string {
  return s.String()
}

type StartForwardTaskParameters struct {
}

func (s StartForwardTaskParameters) String() string {
  return tea.Prettify(s)
}

func (s StartForwardTaskParameters) GoString() string {
  return s.String()
}

type StartForwardTaskRequestHeader struct {
}

func (s StartForwardTaskRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s StartForwardTaskRequestHeader) GoString() string {
  return s.String()
}

type StartForwardTaskResponseHeader struct {
}

func (s StartForwardTaskResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s StartForwardTaskResponseHeader) GoString() string {
  return s.String()
}




type CreatePullChannelRequest struct {
  // {"en":"Channel name, The value contains a maximum of 255 characters", "zh_CN":"频道名称，最长255个字符"}
  ChannelName *string `json:"channelName,omitempty" xml:"channelName,omitempty" require:"true"`
  // {"en":"Source pull url The system pulls the live stream as the live stream source", "zh_CN":"源端拉流 url系统会拉取该直播流作为直播源"}
  SourcePullUrl *string `json:"sourcePullUrl,omitempty" xml:"sourcePullUrl,omitempty" require:"true"`
  // {"en":"Source pull protocol: A channel can have only one back pull protocol, and a source pull basin name can be used for only one source pull protocol channel. 0, RTMP, 1, Http flv", "zh_CN":"源端拉流协议，一个频道只能有一种回源拉流协议，且一个源拉流域名只能用于一种源拉流协议频道；0、RTMP，1、Http flv"}
  SourcePullProtocol *string `json:"sourcePullProtocol,omitempty" xml:"sourcePullProtocol,omitempty" require:"true"`
  // {"en":"Player end pull basin name; The domain name must exist in the background domain name list of the cloud live streaming service", "zh_CN":"播放端拉流域名；域名必须存在于云直播服务后台域名列表中"}
  PullDomain *string `json:"pullDomain,omitempty" xml:"pullDomain,omitempty" require:"true"`
  // {"en":"The value can be multiple, separated by commas (,). The default value is 0 and 2/3. The value cannot be both. 0, RTMP, 1, Http flv, 2, HLS, 3, HDS", "zh_CN":"播放端拉流协议，可多选，用“，”隔开，默认选择 0，2/3 不可同时选择，协议需包含在推流域名实际配置的转出协议类型中；0、RTMP，1、Http flv，2、HLS，3、HDS"}
  PullProtocol *string `json:"pullProtocol,omitempty" xml:"pullProtocol,omitempty"`
  // {"en":"The player can select multiple bit rates, separated by commas (,), and ensure that the selected bit rates are within the range of the transcoding bit rates configured by the basin name. 0, source rate, 1, smooth, 2, standard definition, 3, HD, 4, super clear;", "zh_CN":"播放端拉流码率，可多选，用“，”隔开，确保选择的码率包含在推流域名配置的转码码率范围内；0、源码率，1、流畅，2、标清，3、高清，4、超清；"}
  PullRate *string `json:"pullRate,omitempty" xml:"pullRate,omitempty" require:"true"`
  // {"en":"flash player ID used by the channel. If not set, the default player is used", "zh_CN":"频道使用的 flash 播放器 ID，未设置则使用默认播放器"}
  PlayerId *string `json:"playerId,omitempty" xml:"playerId,omitempty"`
}

func (s CreatePullChannelRequest) String() string {
  return tea.Prettify(s)
}

func (s CreatePullChannelRequest) GoString() string {
  return s.String()
}

func (s *CreatePullChannelRequest) SetChannelName(v string) *CreatePullChannelRequest {
  s.ChannelName = &v
  return s
}

func (s *CreatePullChannelRequest) SetSourcePullUrl(v string) *CreatePullChannelRequest {
  s.SourcePullUrl = &v
  return s
}

func (s *CreatePullChannelRequest) SetSourcePullProtocol(v string) *CreatePullChannelRequest {
  s.SourcePullProtocol = &v
  return s
}

func (s *CreatePullChannelRequest) SetPullDomain(v string) *CreatePullChannelRequest {
  s.PullDomain = &v
  return s
}

func (s *CreatePullChannelRequest) SetPullProtocol(v string) *CreatePullChannelRequest {
  s.PullProtocol = &v
  return s
}

func (s *CreatePullChannelRequest) SetPullRate(v string) *CreatePullChannelRequest {
  s.PullRate = &v
  return s
}

func (s *CreatePullChannelRequest) SetPlayerId(v string) *CreatePullChannelRequest {
  s.PlayerId = &v
  return s
}

type CreatePullChannelResponse struct {
  // {"en":"status code", "zh_CN":"200为成功"}
  Code *int32 `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"message", "zh_CN":"操作成功"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"return data", "zh_CN":"返回数据"}
  CreatePullChannelData *CreatePullChannelData `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s CreatePullChannelResponse) String() string {
  return tea.Prettify(s)
}

func (s CreatePullChannelResponse) GoString() string {
  return s.String()
}

func (s *CreatePullChannelResponse) SetCode(v int32) *CreatePullChannelResponse {
  s.Code = &v
  return s
}

func (s *CreatePullChannelResponse) SetMessage(v string) *CreatePullChannelResponse {
  s.Message = &v
  return s
}

func (s *CreatePullChannelResponse) SetData(v *CreatePullChannelData) *CreatePullChannelResponse {
  s.CreatePullChannelData = v
  return s
}

type CreatePullChannelData struct {
  // {"en":"Pull stream ID, globally unique, immutable", "zh_CN":"拉流 ID，全局唯一值，不可变"}
  PullId *string `json:"pullId,omitempty" xml:"pullId,omitempty" require:"true"`
}

func (s CreatePullChannelData) String() string {
  return tea.Prettify(s)
}

func (s CreatePullChannelData) GoString() string {
  return s.String()
}

func (s *CreatePullChannelData) SetPullId(v string) *CreatePullChannelData {
  s.PullId = &v
  return s
}

type CreatePullChannelPaths struct {
}

func (s CreatePullChannelPaths) String() string {
  return tea.Prettify(s)
}

func (s CreatePullChannelPaths) GoString() string {
  return s.String()
}

type CreatePullChannelParameters struct {
}

func (s CreatePullChannelParameters) String() string {
  return tea.Prettify(s)
}

func (s CreatePullChannelParameters) GoString() string {
  return s.String()
}

type CreatePullChannelRequestHeader struct {
}

func (s CreatePullChannelRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s CreatePullChannelRequestHeader) GoString() string {
  return s.String()
}

type CreatePullChannelResponseHeader struct {
}

func (s CreatePullChannelResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s CreatePullChannelResponseHeader) GoString() string {
  return s.String()
}




type GetDirectoryListRequest struct {
  // {"en":"Page number, default is 1", "zh_CN":"第几页，默认为 1"}
  PageIndex *int32 `json:"pageIndex,omitempty" xml:"pageIndex,omitempty"`
  // {"en":"The value ranges from 1 to 50. The default value is 10", "zh_CN":"每页数量，默认为 10，取值范围 1-50"}
  PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s GetDirectoryListRequest) String() string {
  return tea.Prettify(s)
}

func (s GetDirectoryListRequest) GoString() string {
  return s.String()
}

func (s *GetDirectoryListRequest) SetPageIndex(v int32) *GetDirectoryListRequest {
  s.PageIndex = &v
  return s
}

func (s *GetDirectoryListRequest) SetPageSize(v int32) *GetDirectoryListRequest {
  s.PageSize = &v
  return s
}

type GetDirectoryListResponse struct {
  // {"en":"200 success", "zh_CN":"200为成功"}
  Code *int32 `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Successful operation", "zh_CN":"操作成功"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Return data", "zh_CN":"返回数据"}
  GetDirectoryListData []*GetDirectoryListData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s GetDirectoryListResponse) String() string {
  return tea.Prettify(s)
}

func (s GetDirectoryListResponse) GoString() string {
  return s.String()
}

func (s *GetDirectoryListResponse) SetCode(v int32) *GetDirectoryListResponse {
  s.Code = &v
  return s
}

func (s *GetDirectoryListResponse) SetMessage(v string) *GetDirectoryListResponse {
  s.Message = &v
  return s
}

func (s *GetDirectoryListResponse) SetData(v []*GetDirectoryListData) *GetDirectoryListResponse {
  s.GetDirectoryListData = v
  return s
}

type GetDirectoryListData struct {
  // {"en":"Total number of recorded directories", "zh_CN":"录制目录总数量"}
  TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty" require:"true"`
  // {"en":"The recorded directory list is returned", "zh_CN":"返回的录制目录列表信息"}
  DirectoryList []*GetDirectoryListDirectoryListItem `json:"directoryList,omitempty" xml:"directoryList,omitempty" require:"true" type:"Repeated"`
}

func (s GetDirectoryListData) String() string {
  return tea.Prettify(s)
}

func (s GetDirectoryListData) GoString() string {
  return s.String()
}

func (s *GetDirectoryListData) SetTotalCount(v int32) *GetDirectoryListData {
  s.TotalCount = &v
  return s
}

func (s *GetDirectoryListData) SetDirectoryList(v []*GetDirectoryListDirectoryListItem) *GetDirectoryListData {
  s.DirectoryList = v
  return s
}

type GetDirectoryListDirectoryListItem struct {
  // {"en":"Record directory creation time", "zh_CN":"录制目录创建时间"}
  CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty" require:"true"`
  // {"en":"Record directory id", "zh_CN":"录制目录 ID"}
  DirectoryId *string `json:"directoryId,omitempty" xml:"directoryId,omitempty" require:"true"`
  // {"en":"Record directory type", "zh_CN":"录制目录 type"}
  DirectoryType *string `json:"directoryType,omitempty" xml:"directoryType,omitempty" require:"true"`
  // {"en":"Primary directory", "zh_CN":"一级目录"}
  FirstLevelDirectory *string `json:"firstLevelDirectory,omitempty" xml:"firstLevelDirectory,omitempty" require:"true"`
  // {"en":"Secondary directory", "zh_CN":"二级目录"}
  SecondLevelDirectory *string `json:"secondLevelDirectory,omitempty" xml:"secondLevelDirectory,omitempty" require:"true"`
  // {"en":"Three-level directory", "zh_CN":"三级目录"}
  ThirdLevelDirectory *string `json:"thirdLevelDirectory,omitempty" xml:"thirdLevelDirectory,omitempty" require:"true"`
  // {"en":"Record directory modification time", "zh_CN":"录制目录修改时间"}
  UpdateTime *string `json:"updateTime,omitempty" xml:"updateTime,omitempty" require:"true"`
}

func (s GetDirectoryListDirectoryListItem) String() string {
  return tea.Prettify(s)
}

func (s GetDirectoryListDirectoryListItem) GoString() string {
  return s.String()
}

func (s *GetDirectoryListDirectoryListItem) SetCreateTime(v string) *GetDirectoryListDirectoryListItem {
  s.CreateTime = &v
  return s
}

func (s *GetDirectoryListDirectoryListItem) SetDirectoryId(v string) *GetDirectoryListDirectoryListItem {
  s.DirectoryId = &v
  return s
}

func (s *GetDirectoryListDirectoryListItem) SetDirectoryType(v string) *GetDirectoryListDirectoryListItem {
  s.DirectoryType = &v
  return s
}

func (s *GetDirectoryListDirectoryListItem) SetFirstLevelDirectory(v string) *GetDirectoryListDirectoryListItem {
  s.FirstLevelDirectory = &v
  return s
}

func (s *GetDirectoryListDirectoryListItem) SetSecondLevelDirectory(v string) *GetDirectoryListDirectoryListItem {
  s.SecondLevelDirectory = &v
  return s
}

func (s *GetDirectoryListDirectoryListItem) SetThirdLevelDirectory(v string) *GetDirectoryListDirectoryListItem {
  s.ThirdLevelDirectory = &v
  return s
}

func (s *GetDirectoryListDirectoryListItem) SetUpdateTime(v string) *GetDirectoryListDirectoryListItem {
  s.UpdateTime = &v
  return s
}

type GetDirectoryListPaths struct {
}

func (s GetDirectoryListPaths) String() string {
  return tea.Prettify(s)
}

func (s GetDirectoryListPaths) GoString() string {
  return s.String()
}

type GetDirectoryListParameters struct {
}

func (s GetDirectoryListParameters) String() string {
  return tea.Prettify(s)
}

func (s GetDirectoryListParameters) GoString() string {
  return s.String()
}

type GetDirectoryListRequestHeader struct {
}

func (s GetDirectoryListRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetDirectoryListRequestHeader) GoString() string {
  return s.String()
}

type GetDirectoryListResponseHeader struct {
}

func (s GetDirectoryListResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetDirectoryListResponseHeader) GoString() string {
  return s.String()
}




type DirectoryCreateRequest struct {
  // {"en":"Directory type. The value contains a maximum of 40 fields, including letters, digits, and underscores (_). The name of a user must be unique.", "zh_CN":"目录类型, 只支持中文、大小写字母、数字和下划线，不超过 40 个字段，同个用户下名称不能重复。"}
  DirectoryType *string `json:"directoryType,omitempty" xml:"directoryType,omitempty" require:"true"`
  // {"en":"The level-1 directory cannot start with cloudv or cloudv-, contains a maximum of 128 fields, supports only lowercase letters, digits, and hyphens (-), and cannot be the same as the reserved character. Reserved words have video, live, pullvideo, clip, concat, watermarker, logo, buffer, record, static, AD", "zh_CN":"一级目录，不能以 cloudv 或 cloudv-开头,长度不超过 128 个字段，只支持小写字母、数字和中划线 ， 不 能 与 保 留 字 同 名 ， 保 留 字 有video,live,pullvideo,clip,concat,watermarker,logo,buffer,record,static,ad"}
  FirstLevelDirectory *string `json:"firstLevelDirectory,omitempty" xml:"firstLevelDirectory,omitempty" require:"true"`
  // {"en":"The second-level directory cannot start with cloudv or cloudv- and contains a maximum of 128 fields. Only lowercase letters, digits, and hyphens (-) are supported", "zh_CN":"二级目录，不能以 cloudv 或 cloudv-开头,长度不超过 128 个字段，只支持小写字母、数字和中划线"}
  SecondLevelDirectory *string `json:"secondLevelDirectory,omitempty" xml:"secondLevelDirectory,omitempty"`
  // {"en":"The three-level directory cannot start with cloudv or cloudv-, contains a maximum of 128 fields, and supports only lowercase letters, digits, and hyphens (-)", "zh_CN":"三级目录，不能以 cloudv 或 cloudv-开头,长度不超过 128 个字段，只支持小写字母、数字和中划线"}
  ThirdLevelDirectory *string `json:"thirdLevelDirectory,omitempty" xml:"thirdLevelDirectory,omitempty"`
}

func (s DirectoryCreateRequest) String() string {
  return tea.Prettify(s)
}

func (s DirectoryCreateRequest) GoString() string {
  return s.String()
}

func (s *DirectoryCreateRequest) SetDirectoryType(v string) *DirectoryCreateRequest {
  s.DirectoryType = &v
  return s
}

func (s *DirectoryCreateRequest) SetFirstLevelDirectory(v string) *DirectoryCreateRequest {
  s.FirstLevelDirectory = &v
  return s
}

func (s *DirectoryCreateRequest) SetSecondLevelDirectory(v string) *DirectoryCreateRequest {
  s.SecondLevelDirectory = &v
  return s
}

func (s *DirectoryCreateRequest) SetThirdLevelDirectory(v string) *DirectoryCreateRequest {
  s.ThirdLevelDirectory = &v
  return s
}

type DirectoryCreateResponse struct {
  // {"en":"200 success", "zh_CN":"200为成功"}
  Code *int32 `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Successful operation", "zh_CN":"操作成功"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Return data", "zh_CN":"返回数据"}
  DirectoryCreateData *DirectoryCreateData `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s DirectoryCreateResponse) String() string {
  return tea.Prettify(s)
}

func (s DirectoryCreateResponse) GoString() string {
  return s.String()
}

func (s *DirectoryCreateResponse) SetCode(v int32) *DirectoryCreateResponse {
  s.Code = &v
  return s
}

func (s *DirectoryCreateResponse) SetMessage(v string) *DirectoryCreateResponse {
  s.Message = &v
  return s
}

func (s *DirectoryCreateResponse) SetData(v *DirectoryCreateData) *DirectoryCreateResponse {
  s.DirectoryCreateData = v
  return s
}

type DirectoryCreateData struct {
  // {"en":"Record directory creation time", "zh_CN":"录制目录创建时间"}
  CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty" require:"true"`
  // {"en":"Recorded directory ID", "zh_CN":"录制目录 ID"}
  DirectoryId *string `json:"directoryId,omitempty" xml:"directoryId,omitempty" require:"true"`
  // {"en":"Directory type", "zh_CN":"目录类型"}
  DirectoryType *string `json:"directoryType,omitempty" xml:"directoryType,omitempty" require:"true"`
  // {"en":"Primary directory", "zh_CN":"一级目录"}
  FirstLevelDirectory *string `json:"firstLevelDirectory,omitempty" xml:"firstLevelDirectory,omitempty" require:"true"`
  // {"en":"Secondary directory", "zh_CN":"二级目录"}
  SecondLevelDirectory *string `json:"secondLevelDirectory,omitempty" xml:"secondLevelDirectory,omitempty" require:"true"`
  // {"en":"Three-level directory", "zh_CN":"三级目录"}
  ThirdLevelDirectory *string `json:"thirdLevelDirectory,omitempty" xml:"thirdLevelDirectory,omitempty" require:"true"`
  // {"en":"Record directory modification time", "zh_CN":"录制目录修改时间"}
  UpdateTime *string `json:"updateTime,omitempty" xml:"updateTime,omitempty" require:"true"`
}

func (s DirectoryCreateData) String() string {
  return tea.Prettify(s)
}

func (s DirectoryCreateData) GoString() string {
  return s.String()
}

func (s *DirectoryCreateData) SetCreateTime(v string) *DirectoryCreateData {
  s.CreateTime = &v
  return s
}

func (s *DirectoryCreateData) SetDirectoryId(v string) *DirectoryCreateData {
  s.DirectoryId = &v
  return s
}

func (s *DirectoryCreateData) SetDirectoryType(v string) *DirectoryCreateData {
  s.DirectoryType = &v
  return s
}

func (s *DirectoryCreateData) SetFirstLevelDirectory(v string) *DirectoryCreateData {
  s.FirstLevelDirectory = &v
  return s
}

func (s *DirectoryCreateData) SetSecondLevelDirectory(v string) *DirectoryCreateData {
  s.SecondLevelDirectory = &v
  return s
}

func (s *DirectoryCreateData) SetThirdLevelDirectory(v string) *DirectoryCreateData {
  s.ThirdLevelDirectory = &v
  return s
}

func (s *DirectoryCreateData) SetUpdateTime(v string) *DirectoryCreateData {
  s.UpdateTime = &v
  return s
}

type DirectoryCreatePaths struct {
}

func (s DirectoryCreatePaths) String() string {
  return tea.Prettify(s)
}

func (s DirectoryCreatePaths) GoString() string {
  return s.String()
}

type DirectoryCreateParameters struct {
}

func (s DirectoryCreateParameters) String() string {
  return tea.Prettify(s)
}

func (s DirectoryCreateParameters) GoString() string {
  return s.String()
}

type DirectoryCreateRequestHeader struct {
}

func (s DirectoryCreateRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s DirectoryCreateRequestHeader) GoString() string {
  return s.String()
}

type DirectoryCreateResponseHeader struct {
}

func (s DirectoryCreateResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s DirectoryCreateResponseHeader) GoString() string {
  return s.String()
}




type SetPushTsatcRequest struct {
  // {"en":"Channel pull id", "zh_CN":"频道拉流 id"}
  PullId *string `json:"pullId,omitempty" xml:"pullId,omitempty" require:"true"`
  // {"en":"Validity period: Set the validity period of anti-leaver link parameters, accurate to second.", "zh_CN":"有效时长，设置防盗链参数有效期，精确到秒。"}
  EffectiveHourLong *string `json:"effectiveHourLong,omitempty" xml:"effectiveHourLong,omitempty" require:"true"`
}

func (s SetPushTsatcRequest) String() string {
  return tea.Prettify(s)
}

func (s SetPushTsatcRequest) GoString() string {
  return s.String()
}

func (s *SetPushTsatcRequest) SetPullId(v string) *SetPushTsatcRequest {
  s.PullId = &v
  return s
}

func (s *SetPushTsatcRequest) SetEffectiveHourLong(v string) *SetPushTsatcRequest {
  s.EffectiveHourLong = &v
  return s
}

type SetPushTsatcResponse struct {
  // {"en":"status code", "zh_CN":"状态码"}
  Code *int32 `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"message", "zh_CN":"操作信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"return data", "zh_CN":"返回数据"}
  SetPushTsatcData *SetPushTsatcData `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s SetPushTsatcResponse) String() string {
  return tea.Prettify(s)
}

func (s SetPushTsatcResponse) GoString() string {
  return s.String()
}

func (s *SetPushTsatcResponse) SetCode(v int32) *SetPushTsatcResponse {
  s.Code = &v
  return s
}

func (s *SetPushTsatcResponse) SetMessage(v string) *SetPushTsatcResponse {
  s.Message = &v
  return s
}

func (s *SetPushTsatcResponse) SetData(v *SetPushTsatcData) *SetPushTsatcResponse {
  s.SetPushTsatcData = v
  return s
}

type SetPushTsatcData struct {
  // {"en":"Push URL", "zh_CN":"推流 URL"}
  PushUrl *string `json:"pushUrl,omitempty" xml:"pushUrl,omitempty" require:"true"`
}

func (s SetPushTsatcData) String() string {
  return tea.Prettify(s)
}

func (s SetPushTsatcData) GoString() string {
  return s.String()
}

func (s *SetPushTsatcData) SetPushUrl(v string) *SetPushTsatcData {
  s.PushUrl = &v
  return s
}

type SetPushTsatcPaths struct {
}

func (s SetPushTsatcPaths) String() string {
  return tea.Prettify(s)
}

func (s SetPushTsatcPaths) GoString() string {
  return s.String()
}

type SetPushTsatcParameters struct {
}

func (s SetPushTsatcParameters) String() string {
  return tea.Prettify(s)
}

func (s SetPushTsatcParameters) GoString() string {
  return s.String()
}

type SetPushTsatcRequestHeader struct {
}

func (s SetPushTsatcRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s SetPushTsatcRequestHeader) GoString() string {
  return s.String()
}

type SetPushTsatcResponseHeader struct {
}

func (s SetPushTsatcResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s SetPushTsatcResponseHeader) GoString() string {
  return s.String()
}




type BatchChannelLiveStateRequest struct {
  // {"en":"Pull stream id. Multiple pull stream ids are separated by \",\"; Left blank Default return all channel status", "zh_CN":"拉流 id，多个拉流 id 用\",\"隔开；未填写默认返回所有频道的状态"}
  PullIds *string `json:"pullIds,omitempty" xml:"pullIds,omitempty" require:"true"`
}

func (s BatchChannelLiveStateRequest) String() string {
  return tea.Prettify(s)
}

func (s BatchChannelLiveStateRequest) GoString() string {
  return s.String()
}

func (s *BatchChannelLiveStateRequest) SetPullIds(v string) *BatchChannelLiveStateRequest {
  s.PullIds = &v
  return s
}

type BatchChannelLiveStateResponse struct {
  // {"en":"status code", "zh_CN":"状态码"}
  Code *int32 `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"message", "zh_CN":"操作信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"return data", "zh_CN":"返回数据"}
  Data []*BatchChannelLiveStateChannelStatusList `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s BatchChannelLiveStateResponse) String() string {
  return tea.Prettify(s)
}

func (s BatchChannelLiveStateResponse) GoString() string {
  return s.String()
}

func (s *BatchChannelLiveStateResponse) SetCode(v int32) *BatchChannelLiveStateResponse {
  s.Code = &v
  return s
}

func (s *BatchChannelLiveStateResponse) SetMessage(v string) *BatchChannelLiveStateResponse {
  s.Message = &v
  return s
}

func (s *BatchChannelLiveStateResponse) SetData(v []*BatchChannelLiveStateChannelStatusList) *BatchChannelLiveStateResponse {
  s.Data = v
  return s
}

type BatchChannelLiveStateChannelStatusList struct {
  // {"en":"Return the live stream status corresponding to each pull stream ID:
  // 1. Live broadcast 2. Not live broadcast 3. Banned", "zh_CN":"返回每个拉流 ID 对应的直播流实时状态： 
  // 1、直播中   2、未直播  3、禁播"}
  ChannelStatus *string `json:"channelStatus,omitempty" xml:"channelStatus,omitempty" require:"true"`
  // {"en":"The pull stream id of the channel to be queried must correspond to ChannelState", "zh_CN":"待查询频道的拉流 id，与 ChannelState 要一一对应 "}
  PullId *string `json:"pullId,omitempty" xml:"pullId,omitempty" require:"true"`
}

func (s BatchChannelLiveStateChannelStatusList) String() string {
  return tea.Prettify(s)
}

func (s BatchChannelLiveStateChannelStatusList) GoString() string {
  return s.String()
}

func (s *BatchChannelLiveStateChannelStatusList) SetChannelStatus(v string) *BatchChannelLiveStateChannelStatusList {
  s.ChannelStatus = &v
  return s
}

func (s *BatchChannelLiveStateChannelStatusList) SetPullId(v string) *BatchChannelLiveStateChannelStatusList {
  s.PullId = &v
  return s
}

type BatchChannelLiveStatePaths struct {
}

func (s BatchChannelLiveStatePaths) String() string {
  return tea.Prettify(s)
}

func (s BatchChannelLiveStatePaths) GoString() string {
  return s.String()
}

type BatchChannelLiveStateParameters struct {
}

func (s BatchChannelLiveStateParameters) String() string {
  return tea.Prettify(s)
}

func (s BatchChannelLiveStateParameters) GoString() string {
  return s.String()
}

type BatchChannelLiveStateRequestHeader struct {
}

func (s BatchChannelLiveStateRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s BatchChannelLiveStateRequestHeader) GoString() string {
  return s.String()
}

type BatchChannelLiveStateResponseHeader struct {
}

func (s BatchChannelLiveStateResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s BatchChannelLiveStateResponseHeader) GoString() string {
  return s.String()
}




type RecordTaskQueryRequest struct {
  // {"en":"id of the live recording task", "zh_CN":"直播录制任务id"}
  TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
  // {"en":"The custom task number passed by the customer,transNo and taskId have at least one value that is not empty or an empty string. If both values are passed, only taskId is used as the query condition.", "zh_CN":"客户传入的自定义任务编号,transNo 和 taskId 至少有一个值不为空或空字符串，如果两个都有传的情况下，只使用 taskId 作为查询条件。"}
  TransNo *string `json:"transNo,omitempty" xml:"transNo,omitempty"`
}

func (s RecordTaskQueryRequest) String() string {
  return tea.Prettify(s)
}

func (s RecordTaskQueryRequest) GoString() string {
  return s.String()
}

func (s *RecordTaskQueryRequest) SetTaskId(v string) *RecordTaskQueryRequest {
  s.TaskId = &v
  return s
}

func (s *RecordTaskQueryRequest) SetTransNo(v string) *RecordTaskQueryRequest {
  s.TransNo = &v
  return s
}

type RecordTaskQueryResponse struct {
  // {"en":"200 success", "zh_CN":"200为成功"}
  Code *int32 `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Successful operation", "zh_CN":"操作成功"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Return data", "zh_CN":"返回数据"}
  RecordTaskQueryData *RecordTaskQueryData `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s RecordTaskQueryResponse) String() string {
  return tea.Prettify(s)
}

func (s RecordTaskQueryResponse) GoString() string {
  return s.String()
}

func (s *RecordTaskQueryResponse) SetCode(v int32) *RecordTaskQueryResponse {
  s.Code = &v
  return s
}

func (s *RecordTaskQueryResponse) SetMessage(v string) *RecordTaskQueryResponse {
  s.Message = &v
  return s
}

func (s *RecordTaskQueryResponse) SetData(v *RecordTaskQueryData) *RecordTaskQueryResponse {
  s.RecordTaskQueryData = v
  return s
}

type RecordTaskQueryData struct {
  // {"en":"Status of the recording task: 0 is not started, 1 is started, 2 is normally finished, and 3 is terminated", "zh_CN":"录制任务的状态， 0 未开始，1 已开始，2 正常结束，3 终止"}
  Status *int32 `json:"status,omitempty" xml:"status,omitempty" require:"true"`
  // {"en":"The custom task number passed in by the customer", "zh_CN":"客户传入的自定义任务编号"}
  TransNo *string `json:"transNo,omitempty" xml:"transNo,omitempty" require:"true"`
  // {"en":"Id of the live recording task", "zh_CN":"直播录制任务 ID"}
  TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty" require:"true"`
  // {"en":"List of files that have completed recording", "zh_CN":"录制完成的文件列表"}
  FileList []*RecordTaskQueryFileListItem `json:"fileList,omitempty" xml:"fileList,omitempty" require:"true" type:"Repeated"`
}

func (s RecordTaskQueryData) String() string {
  return tea.Prettify(s)
}

func (s RecordTaskQueryData) GoString() string {
  return s.String()
}

func (s *RecordTaskQueryData) SetStatus(v int32) *RecordTaskQueryData {
  s.Status = &v
  return s
}

func (s *RecordTaskQueryData) SetTransNo(v string) *RecordTaskQueryData {
  s.TransNo = &v
  return s
}

func (s *RecordTaskQueryData) SetTaskId(v string) *RecordTaskQueryData {
  s.TaskId = &v
  return s
}

func (s *RecordTaskQueryData) SetFileList(v []*RecordTaskQueryFileListItem) *RecordTaskQueryData {
  s.FileList = v
  return s
}

type RecordTaskQueryFileListItem struct {
  // {"en":"Video name", "zh_CN":"视频名称"}
  FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty" require:"true"`
  // {"en":"Video id", "zh_CN":"视频id"}
  VideoId *string `json:"videoId,omitempty" xml:"videoId,omitempty" require:"true"`
  // {"en":"Video size", "zh_CN":"视频大小"}
  FileSize *string `json:"fileSize,omitempty" xml:"fileSize,omitempty" require:"true"`
  // {"en":"Video duration", "zh_CN":"视频时长"}
  VideoDuration *string `json:"videoDuration,omitempty" xml:"videoDuration,omitempty" require:"true"`
  // {"en":"Task start time", "zh_CN":"任务开始时间"}
  StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty" require:"true"`
  // {"en":"Task end time", "zh_CN":"任务结束时间"}
  EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty" require:"true"`
  // {"en":"URL of the recording file. If the value-added recording directory management service is enabled, the recorded files are saved to the corresponding recording directory of the channel.", "zh_CN":"录制文件的 URL。如果有开通录制目录管理增值服务，则录制文件会保存到频道对应的录制目录下。"}
  FileUrl *string `json:"fileUrl,omitempty" xml:"fileUrl,omitempty" require:"true"`
}

func (s RecordTaskQueryFileListItem) String() string {
  return tea.Prettify(s)
}

func (s RecordTaskQueryFileListItem) GoString() string {
  return s.String()
}

func (s *RecordTaskQueryFileListItem) SetFileName(v string) *RecordTaskQueryFileListItem {
  s.FileName = &v
  return s
}

func (s *RecordTaskQueryFileListItem) SetVideoId(v string) *RecordTaskQueryFileListItem {
  s.VideoId = &v
  return s
}

func (s *RecordTaskQueryFileListItem) SetFileSize(v string) *RecordTaskQueryFileListItem {
  s.FileSize = &v
  return s
}

func (s *RecordTaskQueryFileListItem) SetVideoDuration(v string) *RecordTaskQueryFileListItem {
  s.VideoDuration = &v
  return s
}

func (s *RecordTaskQueryFileListItem) SetStartTime(v string) *RecordTaskQueryFileListItem {
  s.StartTime = &v
  return s
}

func (s *RecordTaskQueryFileListItem) SetEndTime(v string) *RecordTaskQueryFileListItem {
  s.EndTime = &v
  return s
}

func (s *RecordTaskQueryFileListItem) SetFileUrl(v string) *RecordTaskQueryFileListItem {
  s.FileUrl = &v
  return s
}

type RecordTaskQueryPaths struct {
}

func (s RecordTaskQueryPaths) String() string {
  return tea.Prettify(s)
}

func (s RecordTaskQueryPaths) GoString() string {
  return s.String()
}

type RecordTaskQueryParameters struct {
}

func (s RecordTaskQueryParameters) String() string {
  return tea.Prettify(s)
}

func (s RecordTaskQueryParameters) GoString() string {
  return s.String()
}

type RecordTaskQueryRequestHeader struct {
}

func (s RecordTaskQueryRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s RecordTaskQueryRequestHeader) GoString() string {
  return s.String()
}

type RecordTaskQueryResponseHeader struct {
}

func (s RecordTaskQueryResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s RecordTaskQueryResponseHeader) GoString() string {
  return s.String()
}




type SetPullTsatcRequest struct {
  // {"en":"Channel pull id", "zh_CN":"频道拉流 id"}
  PullId *string `json:"pullId,omitempty" xml:"pullId,omitempty" require:"true"`
  // {"en":"Validity period: Set the validity period of anti-leaver link parameters, accurate to second.", "zh_CN":"有效时长，设置防盗链参数有效期，精确到秒。"}
  EffectiveHourLong *string `json:"effectiveHourLong,omitempty" xml:"effectiveHourLong,omitempty" require:"true"`
}

func (s SetPullTsatcRequest) String() string {
  return tea.Prettify(s)
}

func (s SetPullTsatcRequest) GoString() string {
  return s.String()
}

func (s *SetPullTsatcRequest) SetPullId(v string) *SetPullTsatcRequest {
  s.PullId = &v
  return s
}

func (s *SetPullTsatcRequest) SetEffectiveHourLong(v string) *SetPullTsatcRequest {
  s.EffectiveHourLong = &v
  return s
}

type SetPullTsatcResponse struct {
  // {"en":"Status code", "zh_CN":"状态码"}
  Code *int32 `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Operational infomation", "zh_CN":"操作信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Return data", "zh_CN":"返回数据"}
  Data []*SetPullTsatcPullUrlInfoList `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s SetPullTsatcResponse) String() string {
  return tea.Prettify(s)
}

func (s SetPullTsatcResponse) GoString() string {
  return s.String()
}

func (s *SetPullTsatcResponse) SetCode(v int32) *SetPullTsatcResponse {
  s.Code = &v
  return s
}

func (s *SetPullTsatcResponse) SetMessage(v string) *SetPullTsatcResponse {
  s.Message = &v
  return s
}

func (s *SetPullTsatcResponse) SetData(v []*SetPullTsatcPullUrlInfoList) *SetPullTsatcResponse {
  s.Data = v
  return s
}

type SetPullTsatcPullUrlInfoList struct {
  // {"en":"Pull domain", "zh_CN":"拉流域名"}
  PullDomain *string `json:"pullDomain,omitempty" xml:"pullDomain,omitempty" require:"true"`
  // {"en":"Pull protocal", "zh_CN":"拉流协议"}
  PullProtocol *string `json:"pullProtocol,omitempty" xml:"pullProtocol,omitempty" require:"true"`
  // {"en":"Audio Url", "zh_CN":"音频Url"}
  AudioUrl *string `json:"audioUrl,omitempty" xml:"audioUrl,omitempty" require:"true"`
  // {"en":"Fluent code rate of flow url", "zh_CN":"流畅码率拉流 url"}
  FluentPullUrl *string `json:"fluentPullUrl,omitempty" xml:"fluentPullUrl,omitempty" require:"true"`
  // {"en":"Fluent code rate of flow url(smart HD)", "zh_CN":"流畅码率（智控高清）拉流 url"}
  FluentZkgqPullUrl *string `json:"fluentZkgqPullUrl,omitempty" xml:"fluentZkgqPullUrl,omitempty" require:"true"`
  // {"en":"Super clear bit rate of flow url", "zh_CN":"超清码率拉流 url "}
  HdPullUrl *string `json:"hdPullUrl,omitempty" xml:"hdPullUrl,omitempty" require:"true"`
  // {"en":"Super clear bit rate of flow url(smart HD)", "zh_CN":"超清码率（智控高清）拉流 url"}
  HdZkgqPullUrl *string `json:"hdZkgqPullUrl,omitempty" xml:"hdZkgqPullUrl,omitempty" require:"true"`
  // {"en":"High clear bit of flow url", "zh_CN":"高清码率拉流 url"}
  HighPullUrl *string `json:"highPullUrl,omitempty" xml:"highPullUrl,omitempty" require:"true"`
  // {"en":"High clear bit of flow url(smart HD)", "zh_CN":"高清码率（智控高清）拉流 url"}
  HighZkgqPullUrl *string `json:"highZkgqPullUrl,omitempty" xml:"highZkgqPullUrl,omitempty" require:"true"`
  // {"en":"Origin bit of flow url", "zh_CN":"源码率拉流 url"}
  OriginPullUrl *string `json:"originPullUrl,omitempty" xml:"originPullUrl,omitempty" require:"true"`
  // {"en":"Origin bit of flow url(smart HD)", "zh_CN":"源码率（智控高清）拉流 url"}
  OriginZkgqPullUrl *string `json:"originZkgqPullUrl,omitempty" xml:"originZkgqPullUrl,omitempty" require:"true"`
  // {"en":"Standard bit of flow url", "zh_CN":"标清码率拉流 url"}
  SdPullUrl *string `json:"sdPullUrl,omitempty" xml:"sdPullUrl,omitempty" require:"true"`
  // {"en":"Standard bit of flow url(smart HD)", "zh_CN":"标清码率（智控高清）拉流 url"}
  SdZkgqPullUrl *string `json:"sdZkgqPullUrl,omitempty" xml:"sdZkgqPullUrl,omitempty" require:"true"`
}

func (s SetPullTsatcPullUrlInfoList) String() string {
  return tea.Prettify(s)
}

func (s SetPullTsatcPullUrlInfoList) GoString() string {
  return s.String()
}

func (s *SetPullTsatcPullUrlInfoList) SetPullDomain(v string) *SetPullTsatcPullUrlInfoList {
  s.PullDomain = &v
  return s
}

func (s *SetPullTsatcPullUrlInfoList) SetPullProtocol(v string) *SetPullTsatcPullUrlInfoList {
  s.PullProtocol = &v
  return s
}

func (s *SetPullTsatcPullUrlInfoList) SetAudioUrl(v string) *SetPullTsatcPullUrlInfoList {
  s.AudioUrl = &v
  return s
}

func (s *SetPullTsatcPullUrlInfoList) SetFluentPullUrl(v string) *SetPullTsatcPullUrlInfoList {
  s.FluentPullUrl = &v
  return s
}

func (s *SetPullTsatcPullUrlInfoList) SetFluentZkgqPullUrl(v string) *SetPullTsatcPullUrlInfoList {
  s.FluentZkgqPullUrl = &v
  return s
}

func (s *SetPullTsatcPullUrlInfoList) SetHdPullUrl(v string) *SetPullTsatcPullUrlInfoList {
  s.HdPullUrl = &v
  return s
}

func (s *SetPullTsatcPullUrlInfoList) SetHdZkgqPullUrl(v string) *SetPullTsatcPullUrlInfoList {
  s.HdZkgqPullUrl = &v
  return s
}

func (s *SetPullTsatcPullUrlInfoList) SetHighPullUrl(v string) *SetPullTsatcPullUrlInfoList {
  s.HighPullUrl = &v
  return s
}

func (s *SetPullTsatcPullUrlInfoList) SetHighZkgqPullUrl(v string) *SetPullTsatcPullUrlInfoList {
  s.HighZkgqPullUrl = &v
  return s
}

func (s *SetPullTsatcPullUrlInfoList) SetOriginPullUrl(v string) *SetPullTsatcPullUrlInfoList {
  s.OriginPullUrl = &v
  return s
}

func (s *SetPullTsatcPullUrlInfoList) SetOriginZkgqPullUrl(v string) *SetPullTsatcPullUrlInfoList {
  s.OriginZkgqPullUrl = &v
  return s
}

func (s *SetPullTsatcPullUrlInfoList) SetSdPullUrl(v string) *SetPullTsatcPullUrlInfoList {
  s.SdPullUrl = &v
  return s
}

func (s *SetPullTsatcPullUrlInfoList) SetSdZkgqPullUrl(v string) *SetPullTsatcPullUrlInfoList {
  s.SdZkgqPullUrl = &v
  return s
}

type SetPullTsatcPaths struct {
}

func (s SetPullTsatcPaths) String() string {
  return tea.Prettify(s)
}

func (s SetPullTsatcPaths) GoString() string {
  return s.String()
}

type SetPullTsatcParameters struct {
}

func (s SetPullTsatcParameters) String() string {
  return tea.Prettify(s)
}

func (s SetPullTsatcParameters) GoString() string {
  return s.String()
}

type SetPullTsatcRequestHeader struct {
}

func (s SetPullTsatcRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s SetPullTsatcRequestHeader) GoString() string {
  return s.String()
}

type SetPullTsatcResponseHeader struct {
}

func (s SetPullTsatcResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s SetPullTsatcResponseHeader) GoString() string {
  return s.String()
}




