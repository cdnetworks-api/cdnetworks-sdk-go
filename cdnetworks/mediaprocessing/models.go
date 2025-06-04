package mediaprocessing

// This file is auto-generated, don't edit it. Thanks.
import (
  "github.com/alibabacloud-go/tea/tea"
)

type VideoClipRequest struct {
  // {"en":"The length of the customer video clipping task must be less than or equal to 32 bits. The customer must ensure that it is globally unique on the customer platform", "zh_CN":"客户视频剪切任务编码，长度小于等于32位，客户需要保证在客户平台侧全局唯一"}
  TransNo *string `json:"transNo,omitempty" xml:"transNo,omitempty" require:"true"`
  // {"en":"Video that needs to be edited", "zh_CN":"需要剪辑的视频"}
  VideoId *string `json:"videoId,omitempty" xml:"videoId,omitempty" require:"true"`
  // {"en":"Specifies the start time for video interception, in seconds", "zh_CN":"指定视频截取的开始时间，单位：秒"}
  SeekStart *int32 `json:"seekStart,omitempty" xml:"seekStart,omitempty" require:"true"`
  // {"en":"Specifies the length of the video capture, in seconds", "zh_CN":"指定视频截取的长度，单位：秒"}
  Duration *int32 `json:"duration,omitempty" xml:"duration,omitempty" require:"true"`
  // {"en":"Output video file name", "zh_CN":"输出视频文件名"}
  Filename *string `json:"filename,omitempty" xml:"filename,omitempty" require:"true"`
  // {"en":"Output video resolution in wxh format, for example, 640x480", "zh_CN":"输出视频分辨率，格式为 wxh，例如:640x480"}
  Resolution *string `json:"resolution,omitempty" xml:"resolution,omitempty"`
  // {"en":"Video frame rate: The number of frames displayed per second (unit: Hertz). Commonly used frame rates: 24,25,30, etc", "zh_CN":"视频帧率，每秒显示的帧数，单位：赫兹（Hz）。常用帧率：24，25，30等"}
  FrameRate *int32 `json:"frameRate,omitempty" xml:"frameRate,omitempty"`
  // {"en":"Video bit rate, unit: kbit/s. Common video bit rate: 128,1250,5000, etc", "zh_CN":"视频比特率，单位： kbit/s。常用视频比特率：128，1250，5000等"}
  VideoBitRate *int32 `json:"videoBitRate,omitempty" xml:"videoBitRate,omitempty"`
  // {"en":"Audio bit rate, unit: kbit/s. Common bit rate: 64,128,192,256,320, etc", "zh_CN":"音频码率，单位： kbit/s。常用码率：64，128，192，256，320等"}
  AudioBitRate *int32 `json:"audioBitRate,omitempty" xml:"audioBitRate,omitempty"`
  // {"en":"Audio sampling frequency (unit: Hertz). Common sampling frequency: 8000,12050,22050,44100, etc. Note: flv only supports 4410220511025", "zh_CN":"音频采样频率，单位：赫兹（Hz）。常用采样频率：8000，12050，22050，44100等。注：flv只支持4410220511025"}
  TassamplingRatekId *int32 `json:"tassamplingRatekId,omitempty" xml:"tassamplingRatekId,omitempty"`
  // {"en":"The completion address used to receive callback information", "zh_CN":"用于接收回调信息的完成地址"}
  NotifyUrl *string `json:"notifyUrl,omitempty" xml:"notifyUrl,omitempty"`
  // {"en":"Output format
  // Value range:
  // flv
  // mp4
  // Default flv", "zh_CN":"输出格式
  // 取值范围 ：
  // flv
  // mp4
  // 默认flv"}
  Suffix *string `json:"suffix,omitempty" xml:"suffix,omitempty"`
  // {"en":"Start forced transcoding, 1: start. 0: Do not start, the default is 0", "zh_CN":"启动强制转码，1:启动。0:不启动 默认是0"}
  EnableVideoTranscode *int32 `json:"enableVideoTranscode,omitempty" xml:"enableVideoTranscode,omitempty"`
}

func (s VideoClipRequest) String() string {
  return tea.Prettify(s)
}

func (s VideoClipRequest) GoString() string {
  return s.String()
}

func (s *VideoClipRequest) SetTransNo(v string) *VideoClipRequest {
  s.TransNo = &v
  return s
}

func (s *VideoClipRequest) SetVideoId(v string) *VideoClipRequest {
  s.VideoId = &v
  return s
}

func (s *VideoClipRequest) SetSeekStart(v int32) *VideoClipRequest {
  s.SeekStart = &v
  return s
}

func (s *VideoClipRequest) SetDuration(v int32) *VideoClipRequest {
  s.Duration = &v
  return s
}

func (s *VideoClipRequest) SetFilename(v string) *VideoClipRequest {
  s.Filename = &v
  return s
}

func (s *VideoClipRequest) SetResolution(v string) *VideoClipRequest {
  s.Resolution = &v
  return s
}

func (s *VideoClipRequest) SetFrameRate(v int32) *VideoClipRequest {
  s.FrameRate = &v
  return s
}

func (s *VideoClipRequest) SetVideoBitRate(v int32) *VideoClipRequest {
  s.VideoBitRate = &v
  return s
}

func (s *VideoClipRequest) SetAudioBitRate(v int32) *VideoClipRequest {
  s.AudioBitRate = &v
  return s
}

func (s *VideoClipRequest) SetTassamplingRatekId(v int32) *VideoClipRequest {
  s.TassamplingRatekId = &v
  return s
}

func (s *VideoClipRequest) SetNotifyUrl(v string) *VideoClipRequest {
  s.NotifyUrl = &v
  return s
}

func (s *VideoClipRequest) SetSuffix(v string) *VideoClipRequest {
  s.Suffix = &v
  return s
}

func (s *VideoClipRequest) SetEnableVideoTranscode(v int32) *VideoClipRequest {
  s.EnableVideoTranscode = &v
  return s
}

type VideoClipResponse struct {
  // {"en":"code", "zh_CN":"结果状态码，200为成功"}
  Code *int32 `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"message", "zh_CN":"返回消息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"data", "zh_CN":"返回数据"}
  Data *string `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s VideoClipResponse) String() string {
  return tea.Prettify(s)
}

func (s VideoClipResponse) GoString() string {
  return s.String()
}

func (s *VideoClipResponse) SetCode(v int32) *VideoClipResponse {
  s.Code = &v
  return s
}

func (s *VideoClipResponse) SetMessage(v string) *VideoClipResponse {
  s.Message = &v
  return s
}

func (s *VideoClipResponse) SetData(v string) *VideoClipResponse {
  s.Data = &v
  return s
}

type VideoClipPaths struct {
}

func (s VideoClipPaths) String() string {
  return tea.Prettify(s)
}

func (s VideoClipPaths) GoString() string {
  return s.String()
}

type VideoClipParameters struct {
}

func (s VideoClipParameters) String() string {
  return tea.Prettify(s)
}

func (s VideoClipParameters) GoString() string {
  return s.String()
}

type VideoClipRequestHeader struct {
}

func (s VideoClipRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s VideoClipRequestHeader) GoString() string {
  return s.String()
}

type VideoClipResponseHeader struct {
}

func (s VideoClipResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s VideoClipResponseHeader) GoString() string {
  return s.String()
}




type ClearAdTaskQueryRequest struct {
  // {"en":"AI clear AD task ID", "zh_CN":"AI清除广告任务ID"}
  TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty" require:"true"`
}

func (s ClearAdTaskQueryRequest) String() string {
  return tea.Prettify(s)
}

func (s ClearAdTaskQueryRequest) GoString() string {
  return s.String()
}

func (s *ClearAdTaskQueryRequest) SetTaskId(v string) *ClearAdTaskQueryRequest {
  s.TaskId = &v
  return s
}

type ClearAdTaskQueryResponse struct {
  // {"en":"Status code", "zh_CN":"状态码"}
  Code *int32 `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Operational information", "zh_CN":"操作信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Return data", "zh_CN":"返回数据"}
  ClearAdTaskQueryData *ClearAdTaskQueryData `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s ClearAdTaskQueryResponse) String() string {
  return tea.Prettify(s)
}

func (s ClearAdTaskQueryResponse) GoString() string {
  return s.String()
}

func (s *ClearAdTaskQueryResponse) SetCode(v int32) *ClearAdTaskQueryResponse {
  s.Code = &v
  return s
}

func (s *ClearAdTaskQueryResponse) SetMessage(v string) *ClearAdTaskQueryResponse {
  s.Message = &v
  return s
}

func (s *ClearAdTaskQueryResponse) SetData(v *ClearAdTaskQueryData) *ClearAdTaskQueryResponse {
  s.ClearAdTaskQueryData = v
  return s
}

type ClearAdTaskQueryData struct {
  // {"en":"AI clear AD task ID", "zh_CN":"AI清除广告任务ID"}
  TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty" require:"true"`
  // {"en":"Status:
  // 0(not started)
  // 1(in progress)
  // 2(Successful)
  // 3(failure)", "zh_CN":"状态,取值范围:
  // 0(未开始)
  // 1(进行中)
  // 2(成功)
  // 3(失败)"}
  Status *int32 `json:"status,omitempty" xml:"status,omitempty" require:"true"`
  // {"en":"New video ID", "zh_CN":"生成的新视频ID"}
  Output *string `json:"output,omitempty" xml:"output,omitempty" require:"true"`
}

func (s ClearAdTaskQueryData) String() string {
  return tea.Prettify(s)
}

func (s ClearAdTaskQueryData) GoString() string {
  return s.String()
}

func (s *ClearAdTaskQueryData) SetTaskId(v string) *ClearAdTaskQueryData {
  s.TaskId = &v
  return s
}

func (s *ClearAdTaskQueryData) SetStatus(v int32) *ClearAdTaskQueryData {
  s.Status = &v
  return s
}

func (s *ClearAdTaskQueryData) SetOutput(v string) *ClearAdTaskQueryData {
  s.Output = &v
  return s
}

type ClearAdTaskQueryPaths struct {
}

func (s ClearAdTaskQueryPaths) String() string {
  return tea.Prettify(s)
}

func (s ClearAdTaskQueryPaths) GoString() string {
  return s.String()
}

type ClearAdTaskQueryParameters struct {
}

func (s ClearAdTaskQueryParameters) String() string {
  return tea.Prettify(s)
}

func (s ClearAdTaskQueryParameters) GoString() string {
  return s.String()
}

type ClearAdTaskQueryRequestHeader struct {
}

func (s ClearAdTaskQueryRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ClearAdTaskQueryRequestHeader) GoString() string {
  return s.String()
}

type ClearAdTaskQueryResponseHeader struct {
}

func (s ClearAdTaskQueryResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ClearAdTaskQueryResponseHeader) GoString() string {
  return s.String()
}




type VideoConcatRequest struct {
  // {"en":"The length of the customer video clipping task must be less than or equal to 32 bits. The customer must ensure that it is globally unique on the customer platform", "zh_CN":"客户视频剪切任务编码，长度小于等于32位，客户需要保证在客户平台侧全局唯一"}
  TransNo *string `json:"transNo,omitempty" xml:"transNo,omitempty" require:"true"`
  // {"en":"Separate the videos to be spliced using commas (,) in the splicing order", "zh_CN":"需要拼接的视频，用英文逗号按拼接顺序分隔"}
  VideoIds *string `json:"videoIds,omitempty" xml:"videoIds,omitempty" require:"true"`
  // {"en":"Output video file name", "zh_CN":"输出视频文件名"}
  Filename *string `json:"filename,omitempty" xml:"filename,omitempty" require:"true"`
  // {"en":"ID of the directory to be stored", "zh_CN":"拼接需要存储的目录ID"}
  DirId *string `json:"dirId,omitempty" xml:"dirId,omitempty"`
  // {"en":"The completion address used to receive callback information", "zh_CN":"用于接收回调信息的完成地址"}
  NotifyUrl *string `json:"notifyUrl,omitempty" xml:"notifyUrl,omitempty"`
  // {"en":"Output format
  // Value range:
  // flv
  // mp4
  // m3u8
  // Default flv", "zh_CN":"输出格式
  // 取值范围 ：
  // flv
  // mp4
  // m3u8
  // 默认flv"}
  Suffix *string `json:"suffix,omitempty" xml:"suffix,omitempty"`
}

func (s VideoConcatRequest) String() string {
  return tea.Prettify(s)
}

func (s VideoConcatRequest) GoString() string {
  return s.String()
}

func (s *VideoConcatRequest) SetTransNo(v string) *VideoConcatRequest {
  s.TransNo = &v
  return s
}

func (s *VideoConcatRequest) SetVideoIds(v string) *VideoConcatRequest {
  s.VideoIds = &v
  return s
}

func (s *VideoConcatRequest) SetFilename(v string) *VideoConcatRequest {
  s.Filename = &v
  return s
}

func (s *VideoConcatRequest) SetDirId(v string) *VideoConcatRequest {
  s.DirId = &v
  return s
}

func (s *VideoConcatRequest) SetNotifyUrl(v string) *VideoConcatRequest {
  s.NotifyUrl = &v
  return s
}

func (s *VideoConcatRequest) SetSuffix(v string) *VideoConcatRequest {
  s.Suffix = &v
  return s
}

type VideoConcatResponse struct {
  // {"en":"200 success", "zh_CN":"结果状态码，200为成功"}
  Code *int32 `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Successful operation", "zh_CN":"返回消息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Return data", "zh_CN":"返回数据"}
  Data *string `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s VideoConcatResponse) String() string {
  return tea.Prettify(s)
}

func (s VideoConcatResponse) GoString() string {
  return s.String()
}

func (s *VideoConcatResponse) SetCode(v int32) *VideoConcatResponse {
  s.Code = &v
  return s
}

func (s *VideoConcatResponse) SetMessage(v string) *VideoConcatResponse {
  s.Message = &v
  return s
}

func (s *VideoConcatResponse) SetData(v string) *VideoConcatResponse {
  s.Data = &v
  return s
}

type VideoConcatPaths struct {
}

func (s VideoConcatPaths) String() string {
  return tea.Prettify(s)
}

func (s VideoConcatPaths) GoString() string {
  return s.String()
}

type VideoConcatParameters struct {
}

func (s VideoConcatParameters) String() string {
  return tea.Prettify(s)
}

func (s VideoConcatParameters) GoString() string {
  return s.String()
}

type VideoConcatRequestHeader struct {
}

func (s VideoConcatRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s VideoConcatRequestHeader) GoString() string {
  return s.String()
}

type VideoConcatResponseHeader struct {
}

func (s VideoConcatResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s VideoConcatResponseHeader) GoString() string {
  return s.String()
}




type VideoConcatQueryRequest struct {
  // {"en":"Customer video cut task code, up to 32 bits", "zh_CN":"客户视频剪切任务编码，32位以下"}
  TransNo *string `json:"transNo,omitempty" xml:"transNo,omitempty" require:"true"`
}

func (s VideoConcatQueryRequest) String() string {
  return tea.Prettify(s)
}

func (s VideoConcatQueryRequest) GoString() string {
  return s.String()
}

func (s *VideoConcatQueryRequest) SetTransNo(v string) *VideoConcatQueryRequest {
  s.TransNo = &v
  return s
}

type VideoConcatQueryResponse struct {
  // {"en":"200 success", "zh_CN":"状态码"}
  Code *int32 `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Successful operation", "zh_CN":"操作信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Return data", "zh_CN":"返回数据"}
  VideoConcatQueryData *VideoConcatQueryData `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s VideoConcatQueryResponse) String() string {
  return tea.Prettify(s)
}

func (s VideoConcatQueryResponse) GoString() string {
  return s.String()
}

func (s *VideoConcatQueryResponse) SetCode(v int32) *VideoConcatQueryResponse {
  s.Code = &v
  return s
}

func (s *VideoConcatQueryResponse) SetMessage(v string) *VideoConcatQueryResponse {
  s.Message = &v
  return s
}

func (s *VideoConcatQueryResponse) SetData(v *VideoConcatQueryData) *VideoConcatQueryResponse {
  s.VideoConcatQueryData = v
  return s
}

type VideoConcatQueryData struct {
  // {"en":"taskId", "zh_CN":"任务ID。"}
  TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty" require:"true"`
  // {"en":"transNo", "zh_CN":"业务ID"}
  TransNo *string `json:"transNo,omitempty" xml:"transNo,omitempty" require:"true"`
  // {"en":"status", "zh_CN":"状态
  // 取值范围 ：
  // 1(处理中)
  // 2(成功)
  // 3(失败)"}
  Status *int32 `json:"status,omitempty" xml:"status,omitempty" require:"true"`
  VideoConcatQueryVideoInfo *VideoConcatQueryVideoInfo `json:"videoInfo,omitempty" xml:"videoInfo,omitempty" require:"true"`
}

func (s VideoConcatQueryData) String() string {
  return tea.Prettify(s)
}

func (s VideoConcatQueryData) GoString() string {
  return s.String()
}

func (s *VideoConcatQueryData) SetTaskId(v string) *VideoConcatQueryData {
  s.TaskId = &v
  return s
}

func (s *VideoConcatQueryData) SetTransNo(v string) *VideoConcatQueryData {
  s.TransNo = &v
  return s
}

func (s *VideoConcatQueryData) SetStatus(v int32) *VideoConcatQueryData {
  s.Status = &v
  return s
}

func (s *VideoConcatQueryData) SetVideoInfo(v *VideoConcatQueryVideoInfo) *VideoConcatQueryData {
  s.VideoConcatQueryVideoInfo = v
  return s
}

type VideoConcatQueryVideoInfo struct {
  // {"en":"videoName", "zh_CN":"视频名称"}
  VideoName *string `json:"videoName,omitempty" xml:"videoName,omitempty" require:"true"`
  // {"en":"videoId", "zh_CN":"视频ID"}
  VideoId *string `json:"videoId,omitempty" xml:"videoId,omitempty" require:"true"`
  // {"en":"Whether to encrypt transcoding files
  // Value range:
  // 0(unencrypted)
  // 1(encryption)", "zh_CN":"是否加密转码文件
  // 取值范围 ：
  // 0(不加密)
  // 1(加密)"}
  Encrypt *int32 `json:"encrypt,omitempty" xml:"encrypt,omitempty" require:"true"`
  // {"en":"The space occupied by the video, and the total space used by the video and its transcoding", "zh_CN":"视频占用空间大小，视频及其转码后视频的总空间使用量"}
  VideoSize *string `json:"videoSize,omitempty" xml:"videoSize,omitempty" require:"true"`
  // {"en":"videoDuration", "zh_CN":"视频时长"}
  VideoDuration *string `json:"videoDuration,omitempty" xml:"videoDuration,omitempty" require:"true"`
  // {"en":"uploadTime", "zh_CN":"视频上传时间"}
  UploadTime *string `json:"uploadTime,omitempty" xml:"uploadTime,omitempty" require:"true"`
  // {"en":"updateTime", "zh_CN":"视频修改时间"}
  UpdateTime *string `json:"updateTime,omitempty" xml:"updateTime,omitempty" require:"true"`
  // {"en":"videoDescription", "zh_CN":"视频描述"}
  VideoDescription *string `json:"videoDescription,omitempty" xml:"videoDescription,omitempty" require:"true"`
  // {"en":"videoClassification", "zh_CN":"视频分类"}
  VideoClassification *string `json:"videoClassification,omitempty" xml:"videoClassification,omitempty" require:"true"`
  // {"en":"imageUrl", "zh_CN":"视频封面URL"}
  ImageUrl *string `json:"imageUrl,omitempty" xml:"imageUrl,omitempty" require:"true"`
  // {"en":"publishDomain", "zh_CN":"视频的发布域名"}
  PublishDomain *string `json:"publishDomain,omitempty" xml:"publishDomain,omitempty" require:"true"`
  // {"en":"playerName", "zh_CN":"视频使用的播放器名称"}
  PlayerName *string `json:"playerName,omitempty" xml:"playerName,omitempty" require:"true"`
  // {"en":"playerId", "zh_CN":"视频使用的播放器ID"}
  PlayerId *string `json:"playerId,omitempty" xml:"playerId,omitempty" require:"true"`
  // {"en":"videoState", "zh_CN":"视频状态
  // 取值范围：
  // 0(正常)
  // 1(屏蔽)"}
  VideoState *int32 `json:"videoState,omitempty" xml:"videoState,omitempty" require:"true"`
  // {"en":"If authorized play is not enabled, the video transcoding status ranges from:
  // 1(transcoded)
  // 2(no transcoding)
  // 3(in transcoding)
  // 4(Transcoding fails)
  // Value range of transcoding status when the Authorized Play (video encryption) function is enabled:
  // 1(encrypted transcoding)
  // 2(non-encrypted transcoding)
  // 3(in transcoding)
  // 4(Transcoding fails)
  // 5(no transcoding)", "zh_CN":"未开启授权播放，视频的转码状态的取值范围 ：
  // 1(已转码)
  // 2(未转码)
  // 3(转码中)
  // 4(转码失败)
  // 开启授权播放（视频加密）功能时的转码状态的取值范围 ：
  // 1(已加密转码)
  // 2(非加密转码)
  // 3(转码中)
  // 4(转码失败)
  // 5(未转码)"}
  TranscodeState *int32 `json:"transcodeState,omitempty" xml:"transcodeState,omitempty" require:"true"`
  // {"en":"Video source
  // Value range:
  // 0(other)
  // 1(Upload)
  // 2 (Live streaming to recording)
  // 3 (Video pull)
  // 4 (Video cutting)
  // 5 (Video Stitching)
  // 6 (edge pull recording)
  // 10 (universal version of live broadcasting to recording)
  // 11 (Uploading SDK)
  // 12 (Upload tool)", "zh_CN":"	视频来源
  // 取值范围：
  // 0(其他)
  // 1(上传)
  // 2（直播转录制）
  // 3（视频拉取）
  // 4（视频剪切）
  // 5（视频拼接）
  // 6（边缘拉流录制）
  // 10（通用版直播转录制）
  // 11（上传SDK）
  // 12（上传工具）"}
  VideoSourceCode *int32 `json:"videoSourceCode,omitempty" xml:"videoSourceCode,omitempty" require:"true"`
  // {"en":"Video resolution and other information", "zh_CN":"视频分辨率等信息"}
  VideoConcatQueryVideoResolutions []*VideoConcatQueryVideoResolutions `json:"videoResolutions,omitempty" xml:"videoResolutions,omitempty" require:"true" type:"Repeated"`
}

func (s VideoConcatQueryVideoInfo) String() string {
  return tea.Prettify(s)
}

func (s VideoConcatQueryVideoInfo) GoString() string {
  return s.String()
}

func (s *VideoConcatQueryVideoInfo) SetVideoName(v string) *VideoConcatQueryVideoInfo {
  s.VideoName = &v
  return s
}

func (s *VideoConcatQueryVideoInfo) SetVideoId(v string) *VideoConcatQueryVideoInfo {
  s.VideoId = &v
  return s
}

func (s *VideoConcatQueryVideoInfo) SetEncrypt(v int32) *VideoConcatQueryVideoInfo {
  s.Encrypt = &v
  return s
}

func (s *VideoConcatQueryVideoInfo) SetVideoSize(v string) *VideoConcatQueryVideoInfo {
  s.VideoSize = &v
  return s
}

func (s *VideoConcatQueryVideoInfo) SetVideoDuration(v string) *VideoConcatQueryVideoInfo {
  s.VideoDuration = &v
  return s
}

func (s *VideoConcatQueryVideoInfo) SetUploadTime(v string) *VideoConcatQueryVideoInfo {
  s.UploadTime = &v
  return s
}

func (s *VideoConcatQueryVideoInfo) SetUpdateTime(v string) *VideoConcatQueryVideoInfo {
  s.UpdateTime = &v
  return s
}

func (s *VideoConcatQueryVideoInfo) SetVideoDescription(v string) *VideoConcatQueryVideoInfo {
  s.VideoDescription = &v
  return s
}

func (s *VideoConcatQueryVideoInfo) SetVideoClassification(v string) *VideoConcatQueryVideoInfo {
  s.VideoClassification = &v
  return s
}

func (s *VideoConcatQueryVideoInfo) SetImageUrl(v string) *VideoConcatQueryVideoInfo {
  s.ImageUrl = &v
  return s
}

func (s *VideoConcatQueryVideoInfo) SetPublishDomain(v string) *VideoConcatQueryVideoInfo {
  s.PublishDomain = &v
  return s
}

func (s *VideoConcatQueryVideoInfo) SetPlayerName(v string) *VideoConcatQueryVideoInfo {
  s.PlayerName = &v
  return s
}

func (s *VideoConcatQueryVideoInfo) SetPlayerId(v string) *VideoConcatQueryVideoInfo {
  s.PlayerId = &v
  return s
}

func (s *VideoConcatQueryVideoInfo) SetVideoState(v int32) *VideoConcatQueryVideoInfo {
  s.VideoState = &v
  return s
}

func (s *VideoConcatQueryVideoInfo) SetTranscodeState(v int32) *VideoConcatQueryVideoInfo {
  s.TranscodeState = &v
  return s
}

func (s *VideoConcatQueryVideoInfo) SetVideoSourceCode(v int32) *VideoConcatQueryVideoInfo {
  s.VideoSourceCode = &v
  return s
}

func (s *VideoConcatQueryVideoInfo) SetVideoResolutions(v []*VideoConcatQueryVideoResolutions) *VideoConcatQueryVideoInfo {
  s.VideoConcatQueryVideoResolutions = v
  return s
}

type VideoConcatQueryVideoResolutions struct {
  // {"en":"Clarity. Value range:
  // 1(original painting)
  // 2(fluency)
  // 3(standard definition)
  // 4(HD)
  // 5(Super clear)
  // -99(record file)", "zh_CN":"清晰度。取值范围 ：
  // 1(原画)
  // 2(流畅)
  // 3(标清)
  // 4(高清)
  // 5(超清)
  // -99(录制文件)"}
  Clarity *int32 `json:"clarity,omitempty" xml:"clarity,omitempty" require:"true"`
  // {"en":"Terminal type. Value range:
  // 0(PC)
  // 1(original video)", "zh_CN":"	终端类型。取值范围 ：
  // 0(PC)
  // 1(原视频)"}
  ServerType *int32 `json:"serverType,omitempty" xml:"serverType,omitempty" require:"true"`
  // {"en":"height", "zh_CN":"高度"}
  Height *int32 `json:"height,omitempty" xml:"height,omitempty" require:"true"`
  // {"en":"width", "zh_CN":"宽度"}
  Width *int32 `json:"width,omitempty" xml:"width,omitempty" require:"true"`
  // {"en":"fileSize", "zh_CN":"文件大小(单位为bit)"}
  FileSize *int64 `json:"fileSize,omitempty" xml:"fileSize,omitempty" require:"true"`
}

func (s VideoConcatQueryVideoResolutions) String() string {
  return tea.Prettify(s)
}

func (s VideoConcatQueryVideoResolutions) GoString() string {
  return s.String()
}

func (s *VideoConcatQueryVideoResolutions) SetClarity(v int32) *VideoConcatQueryVideoResolutions {
  s.Clarity = &v
  return s
}

func (s *VideoConcatQueryVideoResolutions) SetServerType(v int32) *VideoConcatQueryVideoResolutions {
  s.ServerType = &v
  return s
}

func (s *VideoConcatQueryVideoResolutions) SetHeight(v int32) *VideoConcatQueryVideoResolutions {
  s.Height = &v
  return s
}

func (s *VideoConcatQueryVideoResolutions) SetWidth(v int32) *VideoConcatQueryVideoResolutions {
  s.Width = &v
  return s
}

func (s *VideoConcatQueryVideoResolutions) SetFileSize(v int64) *VideoConcatQueryVideoResolutions {
  s.FileSize = &v
  return s
}

type VideoConcatQueryPaths struct {
}

func (s VideoConcatQueryPaths) String() string {
  return tea.Prettify(s)
}

func (s VideoConcatQueryPaths) GoString() string {
  return s.String()
}

type VideoConcatQueryParameters struct {
}

func (s VideoConcatQueryParameters) String() string {
  return tea.Prettify(s)
}

func (s VideoConcatQueryParameters) GoString() string {
  return s.String()
}

type VideoConcatQueryRequestHeader struct {
}

func (s VideoConcatQueryRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s VideoConcatQueryRequestHeader) GoString() string {
  return s.String()
}

type VideoConcatQueryResponseHeader struct {
}

func (s VideoConcatQueryResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s VideoConcatQueryResponseHeader) GoString() string {
  return s.String()
}




type CreateClearAdTaskRequest struct {
  // {"en":"videoId", "zh_CN":"视频ID"}
  VideoId *string `json:"videoId,omitempty" xml:"videoId,omitempty" require:"true"`
  // {"en":"Whether to clear video ads. Choose at least one type of AD to remove. If not, the default value is false.", "zh_CN":"是否清除视频广告。至少选择一种要清除的广告。如果不填，默认为false。"}
  ClearVideoAd *bool `json:"clearVideoAd,omitempty" xml:"clearVideoAd,omitempty"`
  // {"en":"Whether to clear watermark ads. Choose at least one type of AD to remove. If not, the default value is false.", "zh_CN":"是否清除水印广告。至少选择一种要清除的广告。如果不填，默认为false。"}
  ClearWatermarkAd *bool `json:"clearWatermarkAd,omitempty" xml:"clearWatermarkAd,omitempty"`
  // {"en":"Whether to clear text ads. Choose at least one type of AD to remove. If not, the default value is false.", "zh_CN":"是否清除文字广告。至少选择一种要清除的广告。如果不填，默认为false。"}
  ClearTextAd *bool `json:"clearTextAd,omitempty" xml:"clearTextAd,omitempty"`
  // {"en":"Type of video ads processed", "zh_CN":"处理的视频广告类型,取值范围:
  // “0”(清除片头广告)
  // “0,1,2”(清除全部广告，包含片头+片中+片尾)
  // 目前只支持“0”和“0,1,2”两种模式，不填则默认为“0,1,2”清楚全部广告"}
  VideoAdTypes *string `json:"videoAdTypes,omitempty" xml:"videoAdTypes,omitempty"`
  // {"en":"Video AD clearance level, value range:
  // 0(default clearance level, may cause missing deletion)
  // 1(Strong clear level, may cause incorrect deletion)
  // If this parameter is not specified, the default value is 0.", "zh_CN":"视频广告清除等级,取值范围:
  // 0(默认清除等级，可能造成漏删)
  // 1(强力清除等级，可能造成误删)
  // 不填则默认为0。"}
  VideoAdLevel *int32 `json:"videoAdLevel,omitempty" xml:"videoAdLevel,omitempty"`
  // {"en":"Processing watermark advertising types, value range:
  // 0(stretch blur)
  // 1(Gaussian blur)
  // The default value is 0.
  // After api-1.240.5 (inclusive), the watermarkAdType and textAdType parameters must be consistent. If they are inconsistent, the system automatically uses 0 by default", "zh_CN":"处理水印广告类型,取值范围:
  // 0(拉伸模糊)
  // 1(高斯模糊)
  // 不填默认为0。
  // 版本api-1.240.5（含）之后，watermarkAdType和textAdType两个参数内容必须一致，若不一致系统自动使用默认0 "}
  WatermarkAdType *int32 `json:"watermarkAdType,omitempty" xml:"watermarkAdType,omitempty"`
  // {"en":"Clear handling of running horse lamp text advertising", "zh_CN":"处理清楚跑马灯文字广告,取值范围:
  // 0(拉伸模糊)
  // 1(高斯模糊)
  // 不填默认为0。"}
  TextAdType *int32 `json:"textAdType,omitempty" xml:"textAdType,omitempty"`
  // {"en":"notifyUrl", "zh_CN":"通知地址，必须以 http://或https:// 开头。长度不能超过255个字符。通知内容详见 [AI清除广告-任务完成通知](https://www.wangsu.com/document/cate/70/17168)"}
  NotifyUrl *string `json:"notifyUrl,omitempty" xml:"notifyUrl,omitempty" require:"true"`
}

func (s CreateClearAdTaskRequest) String() string {
  return tea.Prettify(s)
}

func (s CreateClearAdTaskRequest) GoString() string {
  return s.String()
}

func (s *CreateClearAdTaskRequest) SetVideoId(v string) *CreateClearAdTaskRequest {
  s.VideoId = &v
  return s
}

func (s *CreateClearAdTaskRequest) SetClearVideoAd(v bool) *CreateClearAdTaskRequest {
  s.ClearVideoAd = &v
  return s
}

func (s *CreateClearAdTaskRequest) SetClearWatermarkAd(v bool) *CreateClearAdTaskRequest {
  s.ClearWatermarkAd = &v
  return s
}

func (s *CreateClearAdTaskRequest) SetClearTextAd(v bool) *CreateClearAdTaskRequest {
  s.ClearTextAd = &v
  return s
}

func (s *CreateClearAdTaskRequest) SetVideoAdTypes(v string) *CreateClearAdTaskRequest {
  s.VideoAdTypes = &v
  return s
}

func (s *CreateClearAdTaskRequest) SetVideoAdLevel(v int32) *CreateClearAdTaskRequest {
  s.VideoAdLevel = &v
  return s
}

func (s *CreateClearAdTaskRequest) SetWatermarkAdType(v int32) *CreateClearAdTaskRequest {
  s.WatermarkAdType = &v
  return s
}

func (s *CreateClearAdTaskRequest) SetTextAdType(v int32) *CreateClearAdTaskRequest {
  s.TextAdType = &v
  return s
}

func (s *CreateClearAdTaskRequest) SetNotifyUrl(v string) *CreateClearAdTaskRequest {
  s.NotifyUrl = &v
  return s
}

type CreateClearAdTaskResponse struct {
  // {"en":"200 success", "zh_CN":"状态码"}
  Code *int32 `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Successful operation", "zh_CN":"操作信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Return data", "zh_CN":"返回数据"}
  CreateClearAdTaskData *CreateClearAdTaskData `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s CreateClearAdTaskResponse) String() string {
  return tea.Prettify(s)
}

func (s CreateClearAdTaskResponse) GoString() string {
  return s.String()
}

func (s *CreateClearAdTaskResponse) SetCode(v int32) *CreateClearAdTaskResponse {
  s.Code = &v
  return s
}

func (s *CreateClearAdTaskResponse) SetMessage(v string) *CreateClearAdTaskResponse {
  s.Message = &v
  return s
}

func (s *CreateClearAdTaskResponse) SetData(v *CreateClearAdTaskData) *CreateClearAdTaskResponse {
  s.CreateClearAdTaskData = v
  return s
}

type CreateClearAdTaskData struct {
  // {"en":"taskId", "zh_CN":"AI清除广告任务ID"}
  TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty" require:"true"`
}

func (s CreateClearAdTaskData) String() string {
  return tea.Prettify(s)
}

func (s CreateClearAdTaskData) GoString() string {
  return s.String()
}

func (s *CreateClearAdTaskData) SetTaskId(v string) *CreateClearAdTaskData {
  s.TaskId = &v
  return s
}

type CreateClearAdTaskPaths struct {
}

func (s CreateClearAdTaskPaths) String() string {
  return tea.Prettify(s)
}

func (s CreateClearAdTaskPaths) GoString() string {
  return s.String()
}

type CreateClearAdTaskParameters struct {
}

func (s CreateClearAdTaskParameters) String() string {
  return tea.Prettify(s)
}

func (s CreateClearAdTaskParameters) GoString() string {
  return s.String()
}

type CreateClearAdTaskRequestHeader struct {
}

func (s CreateClearAdTaskRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s CreateClearAdTaskRequestHeader) GoString() string {
  return s.String()
}

type CreateClearAdTaskResponseHeader struct {
}

func (s CreateClearAdTaskResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s CreateClearAdTaskResponseHeader) GoString() string {
  return s.String()
}




type TransCodeRequest struct {
  // {"en":"Please provide the ID of the video(s) you want to view. If there are multiple videos, separate each ID with a comma ,", "zh_CN":"视频ID，多个视频通过“,“隔开；"}
  VideoIds *string `json:"videoIds,omitempty" xml:"videoIds,omitempty" require:"true"`
  // {"en":"By default, there is no watermark added to the video. However, you have the option to select a watermark template from our cloud-based management platform and add it to your video on-demand.", "zh_CN":"水印模板名，默认不添加水印，可选择云点播管理平台中的模板为视频添加水印；"}
  WaterMarkTemplateName *string `json:"waterMarkTemplateName,omitempty" xml:"waterMarkTemplateName,omitempty"`
  // {"en":"The transcoding combination template name, by default, is set to the template provided by the cloud VOD management platform.", "zh_CN":"转码组合模板名，默认使用云点播管理平台设置的默认模板，"}
  TransCodeTemplateName *string `json:"transCodeTemplateName,omitempty" xml:"transCodeTemplateName,omitempty"`
  // {"en":"The Watermark template Id is optional and by default no watermark is added to the video. However, you can conveniently choose a template from the cloud-based management platform to easily add a watermark to your videos", "zh_CN":"水印模板Id，默认不添加水印，可选择云点播管理平台中的模板为视频添加水印；"}
  WatermarkTemplateId *string `json:"watermarkTemplateId,omitempty" xml:"watermarkTemplateId,omitempty"`
  // {"en":"Transcoding combination template ID, the default template set by the cloud on-demand management platform is used by default", "zh_CN":"转码组合模板ID，默认使用云点播管理平台设置的默认模板"}
  TransCodeTemplateId *string `json:"transCodeTemplateId,omitempty" xml:"transCodeTemplateId,omitempty"`
  // {"en":"Subtitle ID, corresponding to the material ID of Cloud VOD material management. After successful upload, subtitles will be automatically Transcoding and added; only ass or srt subtitle formats are supported.", "zh_CN":"字幕ID ,对应云点播素材管理的素材ID。上传成功后会自动转码增加字幕；仅支持ass或srt字幕格式。"}
  SubtitleId *string `json:"subtitleId,omitempty" xml:"subtitleId,omitempty"`
  // {"en":"Supports multiple subtitles, up to 13 subtitles can be added. Only supports multi- Bitrate adaptive Transcoding. Only supports vtt subtitle format.
  // The format content is:
  // lang: subtitle code, which can be defined according to your needs
  // subtitleId: subtitle ID, corresponding to the material ID of Cloud VOD material management
  // code[{\"lang\":\"cn\",\"subtitleId\":\"8a36dfe101921000368ac14400000000\"},{\"lang\":\"en-US\",\"subtitleId\":\"8a38e428019210004d56ef8c00000000\"},{\"lang\":\"ko\",\"subtitleId\":\"8a36dfe101921000368ac14400000000\"}] base64 encryption
  // The subtitle language corresponding code of the console player.
  // Language code
  // Chinese: cn
  // English: en-US
  // Japanese:ja
  // Traditional Chinese:zh-tw
  // French:fr
  // German: de
  // Spanish: es
  // Portuguese:pt
  // Russian:ru
  // Korean:ko
  // Thai:th
  // Vietnamese:vt
  // Indonesian:id"
  //   , "zh_CN":"支持多个字幕，最多可以添加13个字幕。只支持多码率自适应转码。仅支持vtt字幕格式。
  // 格式内容为：
  // lang：字幕code，可以根据自己需求定义
  // subtitleId：字幕ID，对应云点播素材管理的素材ID
  // code[{\"lang\":\"cn\",\"subtitleId\":\"8a36dfe101921000368ac14400000000\"},{\"lang\":\"en-US\",\"subtitleId\":\"8a38e428019210004d56ef8c00000000\"},{\"lang\":\"ko\",\"subtitleId\":\"8a36dfe101921000368ac14400000000\"}] 的base64加密
  // 控制台播放器字幕语言对应code。
  // 语言 code
  // 中文：cn
  // 英文：en-US
  // 日文:ja
  // 繁体中文:zh-tw
  // 法语:fr
  // 德语:de
  // 西班牙语:es
  // 葡萄牙语:pt
  // 俄语:ru
  // 韩语:ko
  // 泰语:th
  // 越南语:vt
  // 印尼语:id"}
  Subtitle *string `json:"subtitle,omitempty" xml:"subtitle,omitempty"`
}

func (s TransCodeRequest) String() string {
  return tea.Prettify(s)
}

func (s TransCodeRequest) GoString() string {
  return s.String()
}

func (s *TransCodeRequest) SetVideoIds(v string) *TransCodeRequest {
  s.VideoIds = &v
  return s
}

func (s *TransCodeRequest) SetWaterMarkTemplateName(v string) *TransCodeRequest {
  s.WaterMarkTemplateName = &v
  return s
}

func (s *TransCodeRequest) SetTransCodeTemplateName(v string) *TransCodeRequest {
  s.TransCodeTemplateName = &v
  return s
}

func (s *TransCodeRequest) SetWatermarkTemplateId(v string) *TransCodeRequest {
  s.WatermarkTemplateId = &v
  return s
}

func (s *TransCodeRequest) SetTransCodeTemplateId(v string) *TransCodeRequest {
  s.TransCodeTemplateId = &v
  return s
}

func (s *TransCodeRequest) SetSubtitleId(v string) *TransCodeRequest {
  s.SubtitleId = &v
  return s
}

func (s *TransCodeRequest) SetSubtitle(v string) *TransCodeRequest {
  s.Subtitle = &v
  return s
}

type TransCodeResponse struct {
  // {"en":"success", "zh_CN":"结果状态码，200为成功"}
  Code *int32 `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"message", "zh_CN":"返回消息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"data", "zh_CN":"返回数据"}
  Data *string `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s TransCodeResponse) String() string {
  return tea.Prettify(s)
}

func (s TransCodeResponse) GoString() string {
  return s.String()
}

func (s *TransCodeResponse) SetCode(v int32) *TransCodeResponse {
  s.Code = &v
  return s
}

func (s *TransCodeResponse) SetMessage(v string) *TransCodeResponse {
  s.Message = &v
  return s
}

func (s *TransCodeResponse) SetData(v string) *TransCodeResponse {
  s.Data = &v
  return s
}

type TransCodePaths struct {
}

func (s TransCodePaths) String() string {
  return tea.Prettify(s)
}

func (s TransCodePaths) GoString() string {
  return s.String()
}

type TransCodeParameters struct {
}

func (s TransCodeParameters) String() string {
  return tea.Prettify(s)
}

func (s TransCodeParameters) GoString() string {
  return s.String()
}

type TransCodeRequestHeader struct {
}

func (s TransCodeRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s TransCodeRequestHeader) GoString() string {
  return s.String()
}

type TransCodeResponseHeader struct {
}

func (s TransCodeResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s TransCodeResponseHeader) GoString() string {
  return s.String()
}




type VideoClipQueryRequest struct {
  // {"en":"Customer video cut task code, up to 32 bits", "zh_CN":"客户视频剪切任务编码，32位以下"}
  TransNo *string `json:"transNo,omitempty" xml:"transNo,omitempty" require:"true"`
}

func (s VideoClipQueryRequest) String() string {
  return tea.Prettify(s)
}

func (s VideoClipQueryRequest) GoString() string {
  return s.String()
}

func (s *VideoClipQueryRequest) SetTransNo(v string) *VideoClipQueryRequest {
  s.TransNo = &v
  return s
}

type VideoClipQueryResponse struct {
  // {"en":"code", "zh_CN":"状态码"}
  Code *int32 `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"message", "zh_CN":"操作信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"data", "zh_CN":"返回数据"}
  VideoClipQueryData *VideoClipQueryData `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s VideoClipQueryResponse) String() string {
  return tea.Prettify(s)
}

func (s VideoClipQueryResponse) GoString() string {
  return s.String()
}

func (s *VideoClipQueryResponse) SetCode(v int32) *VideoClipQueryResponse {
  s.Code = &v
  return s
}

func (s *VideoClipQueryResponse) SetMessage(v string) *VideoClipQueryResponse {
  s.Message = &v
  return s
}

func (s *VideoClipQueryResponse) SetData(v *VideoClipQueryData) *VideoClipQueryResponse {
  s.VideoClipQueryData = v
  return s
}

type VideoClipQueryData struct {
  // {"en":"taskId", "zh_CN":"任务ID。"}
  TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty" require:"true"`
  // {"en":"transNo", "zh_CN":"业务ID"}
  TransNo *string `json:"transNo,omitempty" xml:"transNo,omitempty" require:"true"`
  // {"en":"
  // state
  // Value range:
  // 1(in process)
  // 2(Successful)
  // 3(Failure)", "zh_CN":"状态
  // 取值范围 ：
  // 1(处理中)
  // 2(成功)
  // 3(失败)"}
  Status *int32 `json:"status,omitempty" xml:"status,omitempty" require:"true"`
  VideoClipQueryVideoInfo *VideoClipQueryVideoInfo `json:"videoInfo,omitempty" xml:"videoInfo,omitempty" require:"true"`
}

func (s VideoClipQueryData) String() string {
  return tea.Prettify(s)
}

func (s VideoClipQueryData) GoString() string {
  return s.String()
}

func (s *VideoClipQueryData) SetTaskId(v string) *VideoClipQueryData {
  s.TaskId = &v
  return s
}

func (s *VideoClipQueryData) SetTransNo(v string) *VideoClipQueryData {
  s.TransNo = &v
  return s
}

func (s *VideoClipQueryData) SetStatus(v int32) *VideoClipQueryData {
  s.Status = &v
  return s
}

func (s *VideoClipQueryData) SetVideoInfo(v *VideoClipQueryVideoInfo) *VideoClipQueryData {
  s.VideoClipQueryVideoInfo = v
  return s
}

type VideoClipQueryVideoInfo struct {
  // {"en":"videoName", "zh_CN":"视频名称"}
  VideoName *string `json:"videoName,omitempty" xml:"videoName,omitempty" require:"true"`
  // {"en":"videoId", "zh_CN":"视频ID"}
  VideoId *string `json:"videoId,omitempty" xml:"videoId,omitempty" require:"true"`
  // {"en":"Whether to encrypt transcoding files
  // Value range:
  // 0(unencrypted)
  // 1(encryption)", "zh_CN":"是否加密转码文件
  // 取值范围 ：
  // 0(不加密)
  // 1(加密)"}
  Encrypt *int32 `json:"encrypt,omitempty" xml:"encrypt,omitempty" require:"true"`
  // {"en":"The space occupied by the video, and the total space used by the video and its transcoding", "zh_CN":"视频占用空间大小，视频及其转码后视频的总空间使用量"}
  VideoSize *string `json:"videoSize,omitempty" xml:"videoSize,omitempty" require:"true"`
  // {"en":"videoDuration", "zh_CN":"视频时长"}
  VideoDuration *string `json:"videoDuration,omitempty" xml:"videoDuration,omitempty" require:"true"`
  // {"en":"uploadTime", "zh_CN":"视频上传时间"}
  UploadTime *string `json:"uploadTime,omitempty" xml:"uploadTime,omitempty" require:"true"`
  // {"en":"updateTime", "zh_CN":"视频修改时间"}
  UpdateTime *string `json:"updateTime,omitempty" xml:"updateTime,omitempty" require:"true"`
  // {"en":"videoDescription", "zh_CN":"视频描述"}
  VideoDescription *string `json:"videoDescription,omitempty" xml:"videoDescription,omitempty" require:"true"`
  // {"en":"videoClassification", "zh_CN":"视频分类"}
  VideoClassification *string `json:"videoClassification,omitempty" xml:"videoClassification,omitempty" require:"true"`
  // {"en":"Video cover URL", "zh_CN":"视频封面URL"}
  ImageUrl *string `json:"imageUrl,omitempty" xml:"imageUrl,omitempty" require:"true"`
  // {"en":"publishDomain", "zh_CN":"视频的发布域名"}
  PublishDomain *string `json:"publishDomain,omitempty" xml:"publishDomain,omitempty" require:"true"`
  // {"en":"playerName", "zh_CN":"视频使用的播放器名称"}
  PlayerName *string `json:"playerName,omitempty" xml:"playerName,omitempty" require:"true"`
  // {"en":"playerId", "zh_CN":"视频使用的播放器ID"}
  PlayerId *string `json:"playerId,omitempty" xml:"playerId,omitempty" require:"true"`
  // {"en":"Video state
  // Value range:
  // 0(normal)
  // 1(mask)", "zh_CN":"视频状态
  // 取值范围：
  // 0(正常)
  // 1(屏蔽)"}
  VideoState *int32 `json:"videoState,omitempty" xml:"videoState,omitempty" require:"true"`
  // {"en":"If authorized play is not enabled, the video transcoding status ranges from:
  // 1(transcoded)
  // 2(no transcoding)
  // 3(in transcoding)
  // 4(Transcoding fails)
  // Value range of transcoding status when the Authorized Play (video encryption) function is enabled:
  // 1(encrypted transcoding)
  // 2(non-encrypted transcoding)
  // 3(in transcoding)
  // 4(Transcoding fails)
  // 5(no transcoding)", "zh_CN":"未开启授权播放，视频的转码状态的取值范围 ：
  // 1(已转码)
  // 2(未转码)
  // 3(转码中)
  // 4(转码失败)
  // 开启授权播放（视频加密）功能时的转码状态的取值范围 ：
  // 1(已加密转码)
  // 2(非加密转码)
  // 3(转码中)
  // 4(转码失败)
  // 5(未转码)"}
  TranscodeState *int32 `json:"transcodeState,omitempty" xml:"transcodeState,omitempty" require:"true"`
  // {"en":"Video source
  // Value range:
  // 0(other)
  // 1(Upload)
  // 2 (Live streaming to recording)
  // 3 (Video pull)
  // 4 (Video cutting)
  // 5 (Video Stitching)
  // 6 (edge pull recording)
  // 10 (universal version of live broadcasting to recording)
  // 11 (Uploading SDK)
  // 12 (Upload tool)", "zh_CN":"	视频来源
  // 取值范围：
  // 0(其他)
  // 1(上传)
  // 2（直播转录制）
  // 3（视频拉取）
  // 4（视频剪切）
  // 5（视频拼接）
  // 6（边缘拉流录制）
  // 10（通用版直播转录制）
  // 11（上传SDK）
  // 12（上传工具）"}
  VideoSourceCode *int32 `json:"videoSourceCode,omitempty" xml:"videoSourceCode,omitempty" require:"true"`
  // {"en":"Video resolution and other information", "zh_CN":"视频分辨率等信息"}
  VideoResolutions []*VideoClipQueryVideoResolution `json:"videoResolutions,omitempty" xml:"videoResolutions,omitempty" require:"true" type:"Repeated"`
}

func (s VideoClipQueryVideoInfo) String() string {
  return tea.Prettify(s)
}

func (s VideoClipQueryVideoInfo) GoString() string {
  return s.String()
}

func (s *VideoClipQueryVideoInfo) SetVideoName(v string) *VideoClipQueryVideoInfo {
  s.VideoName = &v
  return s
}

func (s *VideoClipQueryVideoInfo) SetVideoId(v string) *VideoClipQueryVideoInfo {
  s.VideoId = &v
  return s
}

func (s *VideoClipQueryVideoInfo) SetEncrypt(v int32) *VideoClipQueryVideoInfo {
  s.Encrypt = &v
  return s
}

func (s *VideoClipQueryVideoInfo) SetVideoSize(v string) *VideoClipQueryVideoInfo {
  s.VideoSize = &v
  return s
}

func (s *VideoClipQueryVideoInfo) SetVideoDuration(v string) *VideoClipQueryVideoInfo {
  s.VideoDuration = &v
  return s
}

func (s *VideoClipQueryVideoInfo) SetUploadTime(v string) *VideoClipQueryVideoInfo {
  s.UploadTime = &v
  return s
}

func (s *VideoClipQueryVideoInfo) SetUpdateTime(v string) *VideoClipQueryVideoInfo {
  s.UpdateTime = &v
  return s
}

func (s *VideoClipQueryVideoInfo) SetVideoDescription(v string) *VideoClipQueryVideoInfo {
  s.VideoDescription = &v
  return s
}

func (s *VideoClipQueryVideoInfo) SetVideoClassification(v string) *VideoClipQueryVideoInfo {
  s.VideoClassification = &v
  return s
}

func (s *VideoClipQueryVideoInfo) SetImageUrl(v string) *VideoClipQueryVideoInfo {
  s.ImageUrl = &v
  return s
}

func (s *VideoClipQueryVideoInfo) SetPublishDomain(v string) *VideoClipQueryVideoInfo {
  s.PublishDomain = &v
  return s
}

func (s *VideoClipQueryVideoInfo) SetPlayerName(v string) *VideoClipQueryVideoInfo {
  s.PlayerName = &v
  return s
}

func (s *VideoClipQueryVideoInfo) SetPlayerId(v string) *VideoClipQueryVideoInfo {
  s.PlayerId = &v
  return s
}

func (s *VideoClipQueryVideoInfo) SetVideoState(v int32) *VideoClipQueryVideoInfo {
  s.VideoState = &v
  return s
}

func (s *VideoClipQueryVideoInfo) SetTranscodeState(v int32) *VideoClipQueryVideoInfo {
  s.TranscodeState = &v
  return s
}

func (s *VideoClipQueryVideoInfo) SetVideoSourceCode(v int32) *VideoClipQueryVideoInfo {
  s.VideoSourceCode = &v
  return s
}

func (s *VideoClipQueryVideoInfo) SetVideoResolutions(v []*VideoClipQueryVideoResolution) *VideoClipQueryVideoInfo {
  s.VideoResolutions = v
  return s
}

type VideoClipQueryVideoResolution struct {
  // {"en":"Clarity. Value range:
  // 1(original painting)
  // 2(fluency)
  // 3(standard definition)
  // 4(HD)
  // 5(Super clear)
  // -99(record file)", "zh_CN":"清晰度。取值范围 ：
  // 1(原画)
  // 2(流畅)
  // 3(标清)
  // 4(高清)
  // 5(超清)
  // -99(录制文件)"}
  Clarity *int32 `json:"clarity,omitempty" xml:"clarity,omitempty" require:"true"`
  // {"en":"Terminal type. Value range:
  // 0(PC)
  // 1(original video)", "zh_CN":"	终端类型。取值范围 ：
  // 0(PC)
  // 1(原视频)"}
  ServerType *int32 `json:"serverType,omitempty" xml:"serverType,omitempty" require:"true"`
  // {"en":"height", "zh_CN":"高度"}
  Height *int32 `json:"height,omitempty" xml:"height,omitempty" require:"true"`
  // {"en":"width", "zh_CN":"宽度"}
  Width *int32 `json:"width,omitempty" xml:"width,omitempty" require:"true"`
  // {"en":"fileSize", "zh_CN":"文件大小(单位为bit)"}
  FileSize *int64 `json:"fileSize,omitempty" xml:"fileSize,omitempty" require:"true"`
}

func (s VideoClipQueryVideoResolution) String() string {
  return tea.Prettify(s)
}

func (s VideoClipQueryVideoResolution) GoString() string {
  return s.String()
}

func (s *VideoClipQueryVideoResolution) SetClarity(v int32) *VideoClipQueryVideoResolution {
  s.Clarity = &v
  return s
}

func (s *VideoClipQueryVideoResolution) SetServerType(v int32) *VideoClipQueryVideoResolution {
  s.ServerType = &v
  return s
}

func (s *VideoClipQueryVideoResolution) SetHeight(v int32) *VideoClipQueryVideoResolution {
  s.Height = &v
  return s
}

func (s *VideoClipQueryVideoResolution) SetWidth(v int32) *VideoClipQueryVideoResolution {
  s.Width = &v
  return s
}

func (s *VideoClipQueryVideoResolution) SetFileSize(v int64) *VideoClipQueryVideoResolution {
  s.FileSize = &v
  return s
}

type VideoClipQueryPaths struct {
}

func (s VideoClipQueryPaths) String() string {
  return tea.Prettify(s)
}

func (s VideoClipQueryPaths) GoString() string {
  return s.String()
}

type VideoClipQueryParameters struct {
}

func (s VideoClipQueryParameters) String() string {
  return tea.Prettify(s)
}

func (s VideoClipQueryParameters) GoString() string {
  return s.String()
}

type VideoClipQueryRequestHeader struct {
}

func (s VideoClipQueryRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s VideoClipQueryRequestHeader) GoString() string {
  return s.String()
}

type VideoClipQueryResponseHeader struct {
}

func (s VideoClipQueryResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s VideoClipQueryResponseHeader) GoString() string {
  return s.String()
}




