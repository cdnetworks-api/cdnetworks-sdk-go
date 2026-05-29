package uploadassets

// This file is auto-generated, don't edit it. Thanks.
import (
  "github.com/alibabacloud-go/tea/tea"
)

type GetAudioUploadTokenRequest struct {
}

func (s GetAudioUploadTokenRequest) String() string {
  return tea.Prettify(s)
}

func (s GetAudioUploadTokenRequest) GoString() string {
  return s.String()
}

type GetAudioUploadTokenResponse struct {
  // {"en":"Result status code, 200 indicates success", "zh_CN":"结果状态码，200为成功"}
  Code *int32 `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"message", "zh_CN":"返回消息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Return data", "zh_CN":"返回数据"}
  GetAudioUploadTokenData *GetAudioUploadTokenData `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s GetAudioUploadTokenResponse) String() string {
  return tea.Prettify(s)
}

func (s GetAudioUploadTokenResponse) GoString() string {
  return s.String()
}

func (s *GetAudioUploadTokenResponse) SetCode(v int32) *GetAudioUploadTokenResponse {
  s.Code = &v
  return s
}

func (s *GetAudioUploadTokenResponse) SetMessage(v string) *GetAudioUploadTokenResponse {
  s.Message = &v
  return s
}

func (s *GetAudioUploadTokenResponse) SetData(v *GetAudioUploadTokenData) *GetAudioUploadTokenResponse {
  s.GetAudioUploadTokenData = v
  return s
}

type GetAudioUploadTokenData struct {
  // {"en":"Upload url address", "zh_CN":"上传url地址"}
  UploadUrl *string `json:"uploadUrl,omitempty" xml:"uploadUrl,omitempty" require:"true"`
  // {"en":"The bucket name of the WCS to upload", "zh_CN":"要上传的WCS的bucket名称"}
  BucketName *string `json:"bucketName,omitempty" xml:"bucketName,omitempty" require:"true"`
  // {"en":"HttpDns server address", "zh_CN":"HttpDns服务器地址"}
  HttpDnsServer *string `json:"httpDnsServer,omitempty" xml:"httpDnsServer,omitempty" require:"true"`
  // {"en":"Specific token information", "zh_CN":"具体token信息"}
  Items []*GetAudioUploadTokenItem `json:"items,omitempty" xml:"items,omitempty" require:"true" type:"Repeated"`
}

func (s GetAudioUploadTokenData) String() string {
  return tea.Prettify(s)
}

func (s GetAudioUploadTokenData) GoString() string {
  return s.String()
}

func (s *GetAudioUploadTokenData) SetUploadUrl(v string) *GetAudioUploadTokenData {
  s.UploadUrl = &v
  return s
}

func (s *GetAudioUploadTokenData) SetBucketName(v string) *GetAudioUploadTokenData {
  s.BucketName = &v
  return s
}

func (s *GetAudioUploadTokenData) SetHttpDnsServer(v string) *GetAudioUploadTokenData {
  s.HttpDnsServer = &v
  return s
}

func (s *GetAudioUploadTokenData) SetItems(v []*GetAudioUploadTokenItem) *GetAudioUploadTokenData {
  s.Items = v
  return s
}

type GetAudioUploadTokenItem struct {
  // {"en":"File name", "zh_CN":"文件名"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"File suffix", "zh_CN":"文件后缀"}
  Suffix *string `json:"suffix,omitempty" xml:"suffix,omitempty" require:"true"`
  // {"en":"Audio id", "zh_CN":"音频id"}
  AudioId *string `json:"audioId,omitempty" xml:"audioId,omitempty" require:"true"`
  // {"en":"Upload token", "zh_CN":"上传token"}
  UploadToken *string `json:"uploadToken,omitempty" xml:"uploadToken,omitempty" require:"true"`
  // {"en":"The path of the token file was uploaded", "zh_CN":"文件路径"}
  FileFullUrl *string `json:"fileFullUrl,omitempty" xml:"fileFullUrl,omitempty" require:"true"`
}

func (s GetAudioUploadTokenItem) String() string {
  return tea.Prettify(s)
}

func (s GetAudioUploadTokenItem) GoString() string {
  return s.String()
}

func (s *GetAudioUploadTokenItem) SetName(v string) *GetAudioUploadTokenItem {
  s.Name = &v
  return s
}

func (s *GetAudioUploadTokenItem) SetSuffix(v string) *GetAudioUploadTokenItem {
  s.Suffix = &v
  return s
}

func (s *GetAudioUploadTokenItem) SetAudioId(v string) *GetAudioUploadTokenItem {
  s.AudioId = &v
  return s
}

func (s *GetAudioUploadTokenItem) SetUploadToken(v string) *GetAudioUploadTokenItem {
  s.UploadToken = &v
  return s
}

func (s *GetAudioUploadTokenItem) SetFileFullUrl(v string) *GetAudioUploadTokenItem {
  s.FileFullUrl = &v
  return s
}

type GetAudioUploadTokenPaths struct {
}

func (s GetAudioUploadTokenPaths) String() string {
  return tea.Prettify(s)
}

func (s GetAudioUploadTokenPaths) GoString() string {
  return s.String()
}

type GetAudioUploadTokenParameters struct {
  // {"en":"A list of documents that need to get the up-token, expressed as a json string, with url_safe_base64 encoding, up to 50 at a time
  // The parameters are as follows:
  // 1) name: Mandatory. The name of the upload file can contain a maximum of 200 characters
  // 2) suffix: Mandatory, file suffix, such as mp3. Currently, only mp3 format is supported
  // For example: Made from the following string url_safe_base64 coding [{\"name\":\"fileName1\",\"suffix\":\"mp3\"},{\"name\":\"fileName2\",\"suffix\":\"mp3\"}]", "zh_CN":"需要获取上传令牌的文档列表，用json字符串表示，并做url_safe_base64编码，最多一次性获取50个
  // 参数如下：
  // 1）name：必填，上传文件名， 长度最多不能超过200个字符
  // 2）suffix：必填，文件后缀，如mp3，目前只支持mp3格式
  // 例：用以下字符串做url_safe_base64编码[{\"name\":\"fileName1\",\"suffix\":\"mp3\"},{\"name\":\"fileName2\",\"suffix\":\"mp3\"}]"}
  FileList *string `json:"fileList,omitempty" xml:"fileList,omitempty" require:"true"`
  // {"en":"Audio domain name: If this parameter is left blank, set it to the default audio domain name. If the domain name does not exist, an error is returned. Without http:// or https://, for example, xxx.com", "zh_CN":"音频域名，如果不填或为空，则设为默认音频域名。如果域名不存在，返回错误。不带http://或https://，例：xxx.com"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty"`
  // {"en":"Upload policy, whether to overwrite. Default to true", "zh_CN":"上传策略，是否覆盖。默认为true"}
  Overwrite *bool `json:"overwrite,omitempty" xml:"overwrite,omitempty"`
}

func (s GetAudioUploadTokenParameters) String() string {
  return tea.Prettify(s)
}

func (s GetAudioUploadTokenParameters) GoString() string {
  return s.String()
}

func (s *GetAudioUploadTokenParameters) SetFileList(v string) *GetAudioUploadTokenParameters {
  s.FileList = &v
  return s
}

func (s *GetAudioUploadTokenParameters) SetDomain(v string) *GetAudioUploadTokenParameters {
  s.Domain = &v
  return s
}

func (s *GetAudioUploadTokenParameters) SetOverwrite(v bool) *GetAudioUploadTokenParameters {
  s.Overwrite = &v
  return s
}

type GetAudioUploadTokenRequestHeader struct {
}

func (s GetAudioUploadTokenRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetAudioUploadTokenRequestHeader) GoString() string {
  return s.String()
}

type GetAudioUploadTokenResponseHeader struct {
}

func (s GetAudioUploadTokenResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetAudioUploadTokenResponseHeader) GoString() string {
  return s.String()
}




type GetUploadTokenRequest struct {
  // {"en":"Upload the file name, including the file format", "zh_CN":"上传文件名，包含文件格式"}
  OriginFileName *string `json:"originFileName,omitempty" xml:"originFileName,omitempty" require:"true"`
  // {"en":"File ID. The value is a string of up to 32 characters. It is used for resumable data from breakpoints. If the data is not transmitted, the resumable data from breakpoints is not supported.", "zh_CN":"文件ID，最长32位的任意字符串。用于断点续传，如果不传则不支持断点续传功能。"}
  FileId *string `json:"fileId,omitempty" xml:"fileId,omitempty"`
  // {"en":"md5 value of the file, to be deprecated and replaced by fileId", "zh_CN":"文件的md5值，即将弃用，使用fileId代替"}
  FileMd5 *string `json:"fileMd5,omitempty" xml:"fileMd5,omitempty"`
  // {"en":"If the video distribution domain name is not transmitted, the default video management domain name is used as the video distribution domain name. You can set the default domain name on the VOD Console > Global Configuration > Default Domain Name", "zh_CN":"视频发布域名，若不传，则以视频管理默认域名作为本视频的发布域名，可通过云点播控制台>全局配置>默认域名设置默认域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty"`
  // {"en":"Workflow ID, workflowId overrides cmd, waterMarkName, transCodeCombineName, and subtitleId", "zh_CN":"工作流ID，workflowId会覆盖cmd、waterMarkName、transCodeCombineName、subtitleId"}
  WorkflowId *string `json:"workflowId,omitempty" xml:"workflowId,omitempty"`
  // {"en":"Pass policy, whether to overwrite. Value range:
  // 0(no)
  // 1(Yes)", "zh_CN":"上传策略，是否覆盖。取值范围：
  // 0(否)
  // 1(是)"}
  Overwrite *int32 `json:"overwrite,omitempty" xml:"overwrite,omitempty"`
  // {"en":"Video classification. You can specify parent and child categories.Such as: [{\"childName\":\"sub category 1\",\"parentName\":\"parent category 1\"},{\"childName\":\"subclassification 2\",\"parentName\":\"Parent category 2\"}]", "zh_CN":"视频分类，可指定父分类和子分类。如：[{\"childName\":\"子分类1\",\"parentName\":\"父分类1\"},{\"childName\":\"子分类2\",\"parentName\":\"父分类2\"}]"}
  CategoryNames *string `json:"categoryNames,omitempty" xml:"categoryNames,omitempty"`
  // {"en":"Watermark name. After successful upload, the watermark will be automatically transcoded", "zh_CN":"水印名,上传成功后会自动转码增加水印"}
  WaterMarkName *string `json:"waterMarkName,omitempty" xml:"waterMarkName,omitempty"`
  // {"en":"Transcoding combination name, after successful upload will automatically transcoding to increase clarity", "zh_CN":"转码组合名,上传成功后会自动转码增加清晰度"}
  TransCodeCombineName *string `json:"transCodeCombineName,omitempty" xml:"transCodeCombineName,omitempty"`
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

func (s GetUploadTokenRequest) String() string {
  return tea.Prettify(s)
}

func (s GetUploadTokenRequest) GoString() string {
  return s.String()
}

func (s *GetUploadTokenRequest) SetOriginFileName(v string) *GetUploadTokenRequest {
  s.OriginFileName = &v
  return s
}

func (s *GetUploadTokenRequest) SetFileId(v string) *GetUploadTokenRequest {
  s.FileId = &v
  return s
}

func (s *GetUploadTokenRequest) SetFileMd5(v string) *GetUploadTokenRequest {
  s.FileMd5 = &v
  return s
}

func (s *GetUploadTokenRequest) SetDomain(v string) *GetUploadTokenRequest {
  s.Domain = &v
  return s
}

func (s *GetUploadTokenRequest) SetWorkflowId(v string) *GetUploadTokenRequest {
  s.WorkflowId = &v
  return s
}

func (s *GetUploadTokenRequest) SetOverwrite(v int32) *GetUploadTokenRequest {
  s.Overwrite = &v
  return s
}

func (s *GetUploadTokenRequest) SetCategoryNames(v string) *GetUploadTokenRequest {
  s.CategoryNames = &v
  return s
}

func (s *GetUploadTokenRequest) SetWaterMarkName(v string) *GetUploadTokenRequest {
  s.WaterMarkName = &v
  return s
}

func (s *GetUploadTokenRequest) SetTransCodeCombineName(v string) *GetUploadTokenRequest {
  s.TransCodeCombineName = &v
  return s
}

func (s *GetUploadTokenRequest) SetSubtitleId(v string) *GetUploadTokenRequest {
  s.SubtitleId = &v
  return s
}

func (s *GetUploadTokenRequest) SetSubtitle(v string) *GetUploadTokenRequest {
  s.Subtitle = &v
  return s
}

type GetUploadTokenResponse struct {
  // {"en":"Create the result status code. 200 indicates success", "zh_CN":"创建结果状态码，200为成功"}
  Code *int32 `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"message", "zh_CN":"返回消息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"return data", "zh_CN":"返回数据"}
  GetUploadTokenData *GetUploadTokenData `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s GetUploadTokenResponse) String() string {
  return tea.Prettify(s)
}

func (s GetUploadTokenResponse) GoString() string {
  return s.String()
}

func (s *GetUploadTokenResponse) SetCode(v int32) *GetUploadTokenResponse {
  s.Code = &v
  return s
}

func (s *GetUploadTokenResponse) SetMessage(v string) *GetUploadTokenResponse {
  s.Message = &v
  return s
}

func (s *GetUploadTokenResponse) SetData(v *GetUploadTokenData) *GetUploadTokenResponse {
  s.GetUploadTokenData = v
  return s
}

type GetUploadTokenData struct {
  // {"en":"Upload domain address", "zh_CN":"上传域名地址"}
  UploadUrl *string `json:"uploadUrl,omitempty" xml:"uploadUrl,omitempty" require:"true"`
  // {"en":"The relative path to the uploaded file, without the domain name or the leading slash", "zh_CN":"上传文件的相对路径，不带域名和最前面的斜杠"}
  FileKey *string `json:"fileKey,omitempty" xml:"fileKey,omitempty" require:"true"`
  // {"en":"HttpDns server address", "zh_CN":"HttpDns服务器地址"}
  HttpDnsServer *string `json:"httpDnsServer,omitempty" xml:"httpDnsServer,omitempty" require:"true"`
  // {"en":"Video ID. You can use this ID to query information about the video after it is uploaded", "zh_CN":"视频ID，上传完成后可通过该ID查询该视频相关信息"}
  VideoId *string `json:"videoId,omitempty" xml:"videoId,omitempty" require:"true"`
  // {"en":"Upload token", "zh_CN":"上传token"}
  UploadToken *string `json:"uploadToken,omitempty" xml:"uploadToken,omitempty" require:"true"`
  // {"en":"[Planned Deprecation] Upload Accelerated domain address", "zh_CN":"【计划弃用】上传加速域名地址"}
  SpeedDomainlUrl *string `json:"speedDomainlUrl,omitempty" xml:"speedDomainlUrl,omitempty" require:"true"`
  // {"en":"Spatial name", "zh_CN":"空间名称"}
  BucketName *string `json:"bucketName,omitempty" xml:"bucketName,omitempty" require:"true"`
}

func (s GetUploadTokenData) String() string {
  return tea.Prettify(s)
}

func (s GetUploadTokenData) GoString() string {
  return s.String()
}

func (s *GetUploadTokenData) SetUploadUrl(v string) *GetUploadTokenData {
  s.UploadUrl = &v
  return s
}

func (s *GetUploadTokenData) SetFileKey(v string) *GetUploadTokenData {
  s.FileKey = &v
  return s
}

func (s *GetUploadTokenData) SetHttpDnsServer(v string) *GetUploadTokenData {
  s.HttpDnsServer = &v
  return s
}

func (s *GetUploadTokenData) SetVideoId(v string) *GetUploadTokenData {
  s.VideoId = &v
  return s
}

func (s *GetUploadTokenData) SetUploadToken(v string) *GetUploadTokenData {
  s.UploadToken = &v
  return s
}

func (s *GetUploadTokenData) SetSpeedDomainlUrl(v string) *GetUploadTokenData {
  s.SpeedDomainlUrl = &v
  return s
}

func (s *GetUploadTokenData) SetBucketName(v string) *GetUploadTokenData {
  s.BucketName = &v
  return s
}

type GetUploadTokenPaths struct {
}

func (s GetUploadTokenPaths) String() string {
  return tea.Prettify(s)
}

func (s GetUploadTokenPaths) GoString() string {
  return s.String()
}

type GetUploadTokenParameters struct {
}

func (s GetUploadTokenParameters) String() string {
  return tea.Prettify(s)
}

func (s GetUploadTokenParameters) GoString() string {
  return s.String()
}

type GetUploadTokenRequestHeader struct {
}

func (s GetUploadTokenRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetUploadTokenRequestHeader) GoString() string {
  return s.String()
}

type GetUploadTokenResponseHeader struct {
}

func (s GetUploadTokenResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetUploadTokenResponseHeader) GoString() string {
  return s.String()
}




type PullVideoQueryRequest struct {
  // {"en":"Task ID
  // taskId and transNo must pass at least one packet, with taskId taking precedence", "zh_CN":"任务ID
  // taskId和transNo至少传一个，taskId优先使用"}
  TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
  // {"en":"Service ID
  // taskId and transNo must pass at least one packet, with taskId taking precedence", "zh_CN":"业务ID
  // taskId和transNo至少传一个，taskId优先使用"}
  TransNo *string `json:"transNo,omitempty" xml:"transNo,omitempty"`
}

func (s PullVideoQueryRequest) String() string {
  return tea.Prettify(s)
}

func (s PullVideoQueryRequest) GoString() string {
  return s.String()
}

func (s *PullVideoQueryRequest) SetTaskId(v string) *PullVideoQueryRequest {
  s.TaskId = &v
  return s
}

func (s *PullVideoQueryRequest) SetTransNo(v string) *PullVideoQueryRequest {
  s.TransNo = &v
  return s
}

type PullVideoQueryResponse struct {
  // {"en":"Result status code, 200 indicates success", "zh_CN":"结果状态码，200为成功"}
  Code *int32 `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"message", "zh_CN":"返回消息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Return data", "zh_CN":"返回数据"}
  PullVideoQueryData *PullVideoQueryData `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s PullVideoQueryResponse) String() string {
  return tea.Prettify(s)
}

func (s PullVideoQueryResponse) GoString() string {
  return s.String()
}

func (s *PullVideoQueryResponse) SetCode(v int32) *PullVideoQueryResponse {
  s.Code = &v
  return s
}

func (s *PullVideoQueryResponse) SetMessage(v string) *PullVideoQueryResponse {
  s.Message = &v
  return s
}

func (s *PullVideoQueryResponse) SetData(v *PullVideoQueryData) *PullVideoQueryResponse {
  s.PullVideoQueryData = v
  return s
}

type PullVideoQueryData struct {
  // {"en":"Task ID", "zh_CN":"任务ID"}
  TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty" require:"true"`
  // {"en":"trans NO", "zh_CN":"业务ID"}
  TransNo *string `json:"transNo,omitempty" xml:"transNo,omitempty" require:"true"`
  // {"en":"Millisecond timestamp", "zh_CN":"毫秒级时间戳"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {"en":"Task status
  // Value range:
  // 1(in process)
  // 2(Completed)", "zh_CN":"任务状态
  // 取值范围 ：
  // 1(处理中)
  // 2(已完成)"}
  Status *int32 `json:"status,omitempty" xml:"status,omitempty" require:"true"`
  // {"en":"Pull the result status information for each video.
  // If the processing request includes multiple videos, items contain multiple pieces of information", "zh_CN":"每个视频拉取结果状态信息。
  // 如果处理请求包括多个视频，则items包含多条信息"}
  Items []*PullVideoQueryItem `json:"items,omitempty" xml:"items,omitempty" require:"true" type:"Repeated"`
}

func (s PullVideoQueryData) String() string {
  return tea.Prettify(s)
}

func (s PullVideoQueryData) GoString() string {
  return s.String()
}

func (s *PullVideoQueryData) SetTaskId(v string) *PullVideoQueryData {
  s.TaskId = &v
  return s
}

func (s *PullVideoQueryData) SetTransNo(v string) *PullVideoQueryData {
  s.TransNo = &v
  return s
}

func (s *PullVideoQueryData) SetTimestamp(v string) *PullVideoQueryData {
  s.Timestamp = &v
  return s
}

func (s *PullVideoQueryData) SetStatus(v int32) *PullVideoQueryData {
  s.Status = &v
  return s
}

func (s *PullVideoQueryData) SetItems(v []*PullVideoQueryItem) *PullVideoQueryData {
  s.Items = v
  return s
}

type PullVideoQueryItem struct {
  // {"en":"File name", "zh_CN":"文件名"}
  FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty" require:"true"`
  // {"en":"Pull URL", "zh_CN":"拉取URL"}
  FetchUrl *string `json:"fetchUrl,omitempty" xml:"fetchUrl,omitempty" require:"true"`
  // {"en":"Video md5", "zh_CN":"视频md5"}
  Md5 *string `json:"md5,omitempty" xml:"md5,omitempty" require:"true"`
  // {"en":"Pull the task execution status
  // Value range:
  // 1(in process)
  // 2(Successful)
  // 3(Failure)", "zh_CN":"拉取任务执行状态
  // 取值范围 ：
  // 1(处理中)
  // 2(成功)
  // 3(失败)"}
  PullStatus *int32 `json:"pullStatus,omitempty" xml:"pullStatus,omitempty" require:"true"`
  // {"en":"Execution status of the integration command
  // Value range:
  // 1(in process)
  // 2(Successful)
  // 3(Failure)", "zh_CN":"一体化命令执行状态
  // 取值范围 ：
  // 1(处理中)
  // 2(成功)
  // 3(失败)"}
  CmdStatus *int32 `json:"cmdStatus,omitempty" xml:"cmdStatus,omitempty" require:"true"`
  // {"en":"Contains video id, video file list.
  // Video information includes bit rate, sharpness, resolution, terminal type, video URL", "zh_CN":"包含视频id,视频文件列表。
  // 视频信息包含码率，清晰度，分辨率，终端类型，视频URL"}
  PullVideoQueryVideoInfo *PullVideoQueryVideoInfo `json:"videoInfo,omitempty" xml:"videoInfo,omitempty" require:"true"`
}

func (s PullVideoQueryItem) String() string {
  return tea.Prettify(s)
}

func (s PullVideoQueryItem) GoString() string {
  return s.String()
}

func (s *PullVideoQueryItem) SetFileName(v string) *PullVideoQueryItem {
  s.FileName = &v
  return s
}

func (s *PullVideoQueryItem) SetFetchUrl(v string) *PullVideoQueryItem {
  s.FetchUrl = &v
  return s
}

func (s *PullVideoQueryItem) SetMd5(v string) *PullVideoQueryItem {
  s.Md5 = &v
  return s
}

func (s *PullVideoQueryItem) SetPullStatus(v int32) *PullVideoQueryItem {
  s.PullStatus = &v
  return s
}

func (s *PullVideoQueryItem) SetCmdStatus(v int32) *PullVideoQueryItem {
  s.CmdStatus = &v
  return s
}

func (s *PullVideoQueryItem) SetVideoInfo(v *PullVideoQueryVideoInfo) *PullVideoQueryItem {
  s.PullVideoQueryVideoInfo = v
  return s
}

type PullVideoQueryVideoInfo struct {
  // {"en":"Video id", "zh_CN":"视频id"}
  VideoId *string `json:"videoId,omitempty" xml:"videoId,omitempty" require:"true"`
  // {"en":"duration", "zh_CN":"时长"}
  Duration *int64 `json:"duration,omitempty" xml:"duration,omitempty" require:"true"`
  // {"en":"Video file list", "zh_CN":"视频文件列表"}
  VideoFileList []*PullVideoQueryVideoFile `json:"videoFileList,omitempty" xml:"videoFileList,omitempty" require:"true" type:"Repeated"`
}

func (s PullVideoQueryVideoInfo) String() string {
  return tea.Prettify(s)
}

func (s PullVideoQueryVideoInfo) GoString() string {
  return s.String()
}

func (s *PullVideoQueryVideoInfo) SetVideoId(v string) *PullVideoQueryVideoInfo {
  s.VideoId = &v
  return s
}

func (s *PullVideoQueryVideoInfo) SetDuration(v int64) *PullVideoQueryVideoInfo {
  s.Duration = &v
  return s
}

func (s *PullVideoQueryVideoInfo) SetVideoFileList(v []*PullVideoQueryVideoFile) *PullVideoQueryVideoInfo {
  s.VideoFileList = v
  return s
}

type PullVideoQueryVideoFile struct {
  // {"en":"Clarity. Value range:
  // 1(original painting)
  // 2(fluency)
  // 3(standard definition)
  // 4(HD)
  // 5(Super clear)", "zh_CN":"清晰度。取值范围 ：
  // 1(原画)
  // 2(流畅)
  // 3(标清)
  // 4(高清)
  // 5(超清)"}
  Clarity *int32 `json:"clarity,omitempty" xml:"clarity,omitempty" require:"true"`
  // {"en":"Terminal type. Value range:
  // 0(PC)
  // 1(original video)", "zh_CN":"终端类型。取值范围 ：
  // 0(PC)
  // 1(原视频)"}
  ServerType *int32 `json:"serverType,omitempty" xml:"serverType,omitempty" require:"true"`
  // {"en":"Bit rate", "zh_CN":"码率"}
  Bitrate *int32 `json:"bitrate,omitempty" xml:"bitrate,omitempty" require:"true"`
  // {"en":"resolution", "zh_CN":"分辨率"}
  Resolution *string `json:"resolution,omitempty" xml:"resolution,omitempty" require:"true"`
  // {"en":"File size", "zh_CN":"文件大小"}
  FileSize *int64 `json:"fileSize,omitempty" xml:"fileSize,omitempty" require:"true"`
  // {"en":"Video url", "zh_CN":"视频url"}
  FileUrl *string `json:"fileUrl,omitempty" xml:"fileUrl,omitempty" require:"true"`
}

func (s PullVideoQueryVideoFile) String() string {
  return tea.Prettify(s)
}

func (s PullVideoQueryVideoFile) GoString() string {
  return s.String()
}

func (s *PullVideoQueryVideoFile) SetClarity(v int32) *PullVideoQueryVideoFile {
  s.Clarity = &v
  return s
}

func (s *PullVideoQueryVideoFile) SetServerType(v int32) *PullVideoQueryVideoFile {
  s.ServerType = &v
  return s
}

func (s *PullVideoQueryVideoFile) SetBitrate(v int32) *PullVideoQueryVideoFile {
  s.Bitrate = &v
  return s
}

func (s *PullVideoQueryVideoFile) SetResolution(v string) *PullVideoQueryVideoFile {
  s.Resolution = &v
  return s
}

func (s *PullVideoQueryVideoFile) SetFileSize(v int64) *PullVideoQueryVideoFile {
  s.FileSize = &v
  return s
}

func (s *PullVideoQueryVideoFile) SetFileUrl(v string) *PullVideoQueryVideoFile {
  s.FileUrl = &v
  return s
}

type PullVideoQueryPaths struct {
}

func (s PullVideoQueryPaths) String() string {
  return tea.Prettify(s)
}

func (s PullVideoQueryPaths) GoString() string {
  return s.String()
}

type PullVideoQueryParameters struct {
}

func (s PullVideoQueryParameters) String() string {
  return tea.Prettify(s)
}

func (s PullVideoQueryParameters) GoString() string {
  return s.String()
}

type PullVideoQueryRequestHeader struct {
}

func (s PullVideoQueryRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s PullVideoQueryRequestHeader) GoString() string {
  return s.String()
}

type PullVideoQueryResponseHeader struct {
}

func (s PullVideoQueryResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s PullVideoQueryResponseHeader) GoString() string {
  return s.String()
}




type GetMaterialUploadTokenRequest struct {
  // {"en":"A list of documents (parameters including name and suffix) that need to be passed up tokens. A maximum of 50 documents need to be obtained at one time
  // For example: [{"name":"fileName1","suffix":"jpg"},{"name":"fileName2","suffix":"jpg"}]", "zh_CN":"需要获取上传令牌的文档列表（参数包含name、suffix），最多一次性获取50个
  // 例：[{"name":"fileName1","suffix":"jpg"},{"name":"fileName2","suffix":"jpg"}]"}
  FileList []*GetMaterialUploadTokenFile `json:"fileList,omitempty" xml:"fileList,omitempty" require:"true" type:"Repeated"`
  // {"en":"Material domain name. If this parameter is left blank, set it to the default material domain name. If the domain name does not exist, an error is returned. Without http:// or https://, for example, xxx.com", "zh_CN":"素材域名，如果不填或为空，则设为默认素材域名。如果域名不存在，返回错误。不带http://或https://，例：xxx.com"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty"`
  // {"en":"Upload policy, whether to overwrite. Default is false", "zh_CN":"上传策略，是否覆盖。默认为false"}
  Overwrite *bool `json:"overwrite,omitempty" xml:"overwrite,omitempty"`
}

func (s GetMaterialUploadTokenRequest) String() string {
  return tea.Prettify(s)
}

func (s GetMaterialUploadTokenRequest) GoString() string {
  return s.String()
}

func (s *GetMaterialUploadTokenRequest) SetFileList(v []*GetMaterialUploadTokenFile) *GetMaterialUploadTokenRequest {
  s.FileList = v
  return s
}

func (s *GetMaterialUploadTokenRequest) SetDomain(v string) *GetMaterialUploadTokenRequest {
  s.Domain = &v
  return s
}

func (s *GetMaterialUploadTokenRequest) SetOverwrite(v bool) *GetMaterialUploadTokenRequest {
  s.Overwrite = &v
  return s
}

type GetMaterialUploadTokenResponse struct {
  // {"en":"Result status code, 200 indicates success", "zh_CN":"结果状态码，200为成功"}
  Code *int32 `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"message", "zh_CN":"返回消息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Return data", "zh_CN":"返回数据"}
  GetMaterialUploadTokenData *GetMaterialUploadTokenData `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s GetMaterialUploadTokenResponse) String() string {
  return tea.Prettify(s)
}

func (s GetMaterialUploadTokenResponse) GoString() string {
  return s.String()
}

func (s *GetMaterialUploadTokenResponse) SetCode(v int32) *GetMaterialUploadTokenResponse {
  s.Code = &v
  return s
}

func (s *GetMaterialUploadTokenResponse) SetMessage(v string) *GetMaterialUploadTokenResponse {
  s.Message = &v
  return s
}

func (s *GetMaterialUploadTokenResponse) SetData(v *GetMaterialUploadTokenData) *GetMaterialUploadTokenResponse {
  s.GetMaterialUploadTokenData = v
  return s
}

type GetMaterialUploadTokenFile struct {
  // {"en":"The upload file name contains a maximum of 40 characters", "zh_CN":"上传文件名， 长度最多不能超过40个字符"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"GetMaterialUploadTokenFile suffix", "zh_CN":"文件后缀"}
  Suffix *string `json:"suffix,omitempty" xml:"suffix,omitempty"`
}

func (s GetMaterialUploadTokenFile) String() string {
  return tea.Prettify(s)
}

func (s GetMaterialUploadTokenFile) GoString() string {
  return s.String()
}

func (s *GetMaterialUploadTokenFile) SetName(v string) *GetMaterialUploadTokenFile {
  s.Name = &v
  return s
}

func (s *GetMaterialUploadTokenFile) SetSuffix(v string) *GetMaterialUploadTokenFile {
  s.Suffix = &v
  return s
}

type GetMaterialUploadTokenData struct {
  // {"en":"Upload url address", "zh_CN":"上传url地址"}
  UploadUrl *string `json:"uploadUrl,omitempty" xml:"uploadUrl,omitempty" require:"true"`
  // {"en":"The bucket name of the WCS to upload", "zh_CN":"要上传的WCS的bucket名称"}
  BucketName *string `json:"bucketName,omitempty" xml:"bucketName,omitempty" require:"true"`
  // {"en":"HttpDns server address", "zh_CN":"HttpDns服务器地址"}
  HttpDnsServer *string `json:"httpDnsServer,omitempty" xml:"httpDnsServer,omitempty" require:"true"`
  // {"en":"Specific token information", "zh_CN":"具体token信息"}
  Items []*GetMaterialUploadTokenItem `json:"items,omitempty" xml:"items,omitempty" require:"true" type:"Repeated"`
}

func (s GetMaterialUploadTokenData) String() string {
  return tea.Prettify(s)
}

func (s GetMaterialUploadTokenData) GoString() string {
  return s.String()
}

func (s *GetMaterialUploadTokenData) SetUploadUrl(v string) *GetMaterialUploadTokenData {
  s.UploadUrl = &v
  return s
}

func (s *GetMaterialUploadTokenData) SetBucketName(v string) *GetMaterialUploadTokenData {
  s.BucketName = &v
  return s
}

func (s *GetMaterialUploadTokenData) SetHttpDnsServer(v string) *GetMaterialUploadTokenData {
  s.HttpDnsServer = &v
  return s
}

func (s *GetMaterialUploadTokenData) SetItems(v []*GetMaterialUploadTokenItem) *GetMaterialUploadTokenData {
  s.Items = v
  return s
}

type GetMaterialUploadTokenItem struct {
  // {"en":"GetMaterialUploadTokenFile name", "zh_CN":"文件名"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"GetMaterialUploadTokenFile suffix", "zh_CN":"文件后缀"}
  Suffix *string `json:"suffix,omitempty" xml:"suffix,omitempty" require:"true"`
  // {"en":"Material id", "zh_CN":"素材id"}
  MaterialId *string `json:"materialId,omitempty" xml:"materialId,omitempty" require:"true"`
  // {"en":"Upload token", "zh_CN":"上传token"}
  UploadToken *string `json:"uploadToken,omitempty" xml:"uploadToken,omitempty" require:"true"`
  // {"en":"GetMaterialUploadTokenFile path", "zh_CN":"文件路径"}
  FileFullUrl *string `json:"fileFullUrl,omitempty" xml:"fileFullUrl,omitempty" require:"true"`
}

func (s GetMaterialUploadTokenItem) String() string {
  return tea.Prettify(s)
}

func (s GetMaterialUploadTokenItem) GoString() string {
  return s.String()
}

func (s *GetMaterialUploadTokenItem) SetName(v string) *GetMaterialUploadTokenItem {
  s.Name = &v
  return s
}

func (s *GetMaterialUploadTokenItem) SetSuffix(v string) *GetMaterialUploadTokenItem {
  s.Suffix = &v
  return s
}

func (s *GetMaterialUploadTokenItem) SetMaterialId(v string) *GetMaterialUploadTokenItem {
  s.MaterialId = &v
  return s
}

func (s *GetMaterialUploadTokenItem) SetUploadToken(v string) *GetMaterialUploadTokenItem {
  s.UploadToken = &v
  return s
}

func (s *GetMaterialUploadTokenItem) SetFileFullUrl(v string) *GetMaterialUploadTokenItem {
  s.FileFullUrl = &v
  return s
}

type GetMaterialUploadTokenPaths struct {
}

func (s GetMaterialUploadTokenPaths) String() string {
  return tea.Prettify(s)
}

func (s GetMaterialUploadTokenPaths) GoString() string {
  return s.String()
}

type GetMaterialUploadTokenParameters struct {
}

func (s GetMaterialUploadTokenParameters) String() string {
  return tea.Prettify(s)
}

func (s GetMaterialUploadTokenParameters) GoString() string {
  return s.String()
}

type GetMaterialUploadTokenRequestHeader struct {
}

func (s GetMaterialUploadTokenRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetMaterialUploadTokenRequestHeader) GoString() string {
  return s.String()
}

type GetMaterialUploadTokenResponseHeader struct {
}

func (s GetMaterialUploadTokenResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetMaterialUploadTokenResponseHeader) GoString() string {
  return s.String()
}




type PullVideoRequest struct {
  // {"en":"The service ID must be unique and controlled by users
  // It is recommended that a 32-bit UUID be a string of up to 32 characters", "zh_CN":"业务ID，需用户自己控制唯一性
  // 建议使用32位UUID，并且最长为32位字符串"}
  TransNo *string `json:"transNo,omitempty" xml:"transNo,omitempty"`
  // {"en":"Command in the following format
  // {tcTemplateName:xxxx,workflowCode:xxx}", "zh_CN":"命令，格式如下
  // {tcTemplateName:xxxx,workflowCode:xxx}"}
  Cmd *string `json:"cmd,omitempty" xml:"cmd,omitempty"`
  // {"en":"The url and corresponding parameters of the video to be pulled are expressed in json strings and encoded in url_safe_base64. A maximum of 50 urls and parameters can be pulled at a time.
  // The parameters are as follows:
  // 1) filName: Specifies the name of the pulled file. It is recommended to include a format suffix. If no format suffix is included, the pulled video will not have a format suffix. If not, the file name is the URI at the end of the URL
  // 2) fetchUrl: Required, url of the file to be pulled
  // 3) md5: indicates the md5 value of the file to be pulled. Used to verify whether a file is damaged after being pulled. Do not fill, do not check
  // Example: Use the following string for url_safe_base64 encoding [{\"fileName\":\"fileName1\",\"fetchUrl\":\"fetchUrl1\",\"md5\":\"md51\"},{\"fileName\":\"fileName2\",\"fetchUrl\":\"fetchUrl2\",\"md5\":\"md52\"}]", "zh_CN":"待拉取视频的url及对应参数，用json字符串表示，并做url_safe_base64编码，最多一次性拉取50个。
  // 参数如下：
  // 1）fileName:指定拉取文件的文件名，建议包含格式后缀，如果不包含格式后缀，拉取后的视频也会没有格式后缀。如果不传，以URL最后一段URI为文件名
  // 2）fetchUrl:必填，待拉取文件url
  // 3）md5:待拉取文件md5值；用于验证拉取后文件是否有损坏。不填就不校验
  // 例：用以下字符串做url_safe_base64编码[{\"fileName\":\"fileName1\",\"fetchUrl\":\"fetchUrl1\",\"md5\":\"md51\"},{\"fileName\":\"fileName2\",\"fetchUrl\":\"fetchUrl2\",\"md5\":\"md52\"}]"}
  FileList *string `json:"fileList,omitempty" xml:"fileList,omitempty" require:"true"`
  // {"en":"Workflow ID", "zh_CN":"工作流ID"}
  WorkflowId *string `json:"workflowId,omitempty" xml:"workflowId,omitempty"`
  // {"en":"url to receive processing results. This url does not need to be encrypted.
  // For details about the notification content, see the interface document. The notification content is encoded in url_safe_base64", "zh_CN":"接收处理结果的url，该url不需要做加密操作。
  // 通知内容详见接口文档的说明，通知的内容会做url_safe_base64编码"}
  NotifyUrl *string `json:"notifyUrl,omitempty" xml:"notifyUrl,omitempty"`
  // {"en":"Whether the processing instructions for pulling multiple videos are notified separately.
  // Value range:
  // 0(entire batch merge notification)
  // 1(each video is notified independently)
  // Default is 0", "zh_CN":"拉取多个视频的处理指令是否分开通知。
  // 取值范围 ：
  // 0(整个批次合并通知)
  // 1(每个视频独立通知)
  // 默认为0"}
  Separate *int32 `json:"separate,omitempty" xml:"separate,omitempty"`
  // {"en":"The default value is 1.
  // 0: indicates not to capture TS files 
  // 1: Indicates that TS files need to be captured
  // Note: Encrypted hls files cannot be pulled", "zh_CN":"默认为1；
  // 0：表示不抓取TS文件 
  // 1：表示需要抓取TS文件
  // 备注：不能拉取加密的hls文件"}
  FetchTs *int32 `json:"fetchTs,omitempty" xml:"fetchTs,omitempty"`
}

func (s PullVideoRequest) String() string {
  return tea.Prettify(s)
}

func (s PullVideoRequest) GoString() string {
  return s.String()
}

func (s *PullVideoRequest) SetTransNo(v string) *PullVideoRequest {
  s.TransNo = &v
  return s
}

func (s *PullVideoRequest) SetCmd(v string) *PullVideoRequest {
  s.Cmd = &v
  return s
}

func (s *PullVideoRequest) SetFileList(v string) *PullVideoRequest {
  s.FileList = &v
  return s
}

func (s *PullVideoRequest) SetWorkflowId(v string) *PullVideoRequest {
  s.WorkflowId = &v
  return s
}

func (s *PullVideoRequest) SetNotifyUrl(v string) *PullVideoRequest {
  s.NotifyUrl = &v
  return s
}

func (s *PullVideoRequest) SetSeparate(v int32) *PullVideoRequest {
  s.Separate = &v
  return s
}

func (s *PullVideoRequest) SetFetchTs(v int32) *PullVideoRequest {
  s.FetchTs = &v
  return s
}

type PullVideoFile struct {
  // {"en":"You are advised to include a format suffix. If no format suffix is included, the video that is pulled does not have a format suffix. If not, the file name is the URI at the end of the URL", "zh_CN":"指定拉取文件的文件名，建议包含格式后缀，如果不包含格式后缀，拉取后的视频也会没有格式后缀。如果不传，以URL最后一段URI为文件名"}
  FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
  // {"en":"url of the file to be pulled", "zh_CN":"待拉取文件url"}
  FetchUrl *string `json:"fetchUrl,omitempty" xml:"fetchUrl,omitempty" require:"true"`
  // {"en":"md5 value of the file to be pulled; Used to verify whether a file is damaged after being pulled. Do not fill, do not check", "zh_CN":"待拉取文件md5值；用于验证拉取后文件是否有损坏。不填就不校验"}
  Md5 *string `json:"md5,omitempty" xml:"md5,omitempty"`
}

func (s PullVideoFile) String() string {
  return tea.Prettify(s)
}

func (s PullVideoFile) GoString() string {
  return s.String()
}

func (s *PullVideoFile) SetFileName(v string) *PullVideoFile {
  s.FileName = &v
  return s
}

func (s *PullVideoFile) SetFetchUrl(v string) *PullVideoFile {
  s.FetchUrl = &v
  return s
}

func (s *PullVideoFile) SetMd5(v string) *PullVideoFile {
  s.Md5 = &v
  return s
}

type PullVideoResponse struct {
  // {"en":"Result status code, 200 indicates success", "zh_CN":"结果状态码，200为成功"}
  Code *int32 `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"message", "zh_CN":"返回消息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Return data", "zh_CN":"返回数据"}
  PullVideoData *PullVideoData `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s PullVideoResponse) String() string {
  return tea.Prettify(s)
}

func (s PullVideoResponse) GoString() string {
  return s.String()
}

func (s *PullVideoResponse) SetCode(v int32) *PullVideoResponse {
  s.Code = &v
  return s
}

func (s *PullVideoResponse) SetMessage(v string) *PullVideoResponse {
  s.Message = &v
  return s
}

func (s *PullVideoResponse) SetData(v *PullVideoData) *PullVideoResponse {
  s.PullVideoData = v
  return s
}

type PullVideoData struct {
  // {"en":"trans NO", "zh_CN":"业务ID"}
  TranNo *string `json:"tranNo,omitempty" xml:"tranNo,omitempty" require:"true"`
  // {"en":"Task ID", "zh_CN":"任务ID"}
  TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty" require:"true"`
}

func (s PullVideoData) String() string {
  return tea.Prettify(s)
}

func (s PullVideoData) GoString() string {
  return s.String()
}

func (s *PullVideoData) SetTranNo(v string) *PullVideoData {
  s.TranNo = &v
  return s
}

func (s *PullVideoData) SetTaskId(v string) *PullVideoData {
  s.TaskId = &v
  return s
}

type PullVideoPaths struct {
}

func (s PullVideoPaths) String() string {
  return tea.Prettify(s)
}

func (s PullVideoPaths) GoString() string {
  return s.String()
}

type PullVideoParameters struct {
}

func (s PullVideoParameters) String() string {
  return tea.Prettify(s)
}

func (s PullVideoParameters) GoString() string {
  return s.String()
}

type PullVideoRequestHeader struct {
}

func (s PullVideoRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s PullVideoRequestHeader) GoString() string {
  return s.String()
}

type PullVideoResponseHeader struct {
}

func (s PullVideoResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s PullVideoResponseHeader) GoString() string {
  return s.String()
}




