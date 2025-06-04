package recording

// This file is auto-generated, don't edit it. Thanks.
import (
  "github.com/alibabacloud-go/tea/tea"
)

type LiveVideoConcatRequest struct {
  // {"en":"Channel pull id", "zh_CN":"频道拉流id"}
  PullId *string `json:"pullId,omitempty" xml:"pullId,omitempty" require:"true"`
  // {"en":"The service ID must be unique. You are advised to use a 32-bit UUID and the value can be a string of up to 32 characters", "zh_CN":"业务ID，需用户自己控制唯一性，建议使用32位UUID，并且最长为32位字符串"}
  TransNo *string `json:"transNo,omitempty" xml:"transNo,omitempty"`
  // {"en":"Start time, unix time stamp. Default is the first broadcast start time", "zh_CN":"开始时间，unix时间戳，默认为第一次直播开始时间"}
  Start *int64 `json:"start,omitempty" xml:"start,omitempty"`
  // {"en":"End time, unix timestamp, default current time. End time Future time is prohibited", "zh_CN":"结束时间，unix时间戳，默认当前时间。结束时间禁止填未来时间"}
  End *int64 `json:"end,omitempty" xml:"end,omitempty"`
  // {"en":"File name. If it is empty, the system automatically generates a file name (stream name _ start time _ end time).", "zh_CN":"文件名。若为空则系统自动生成一个文件名（流名_开始时间_结束时间）"}
  Fname *string `json:"fname,omitempty" xml:"fname,omitempty"`
  // {"en":"The optional file format is:
  // 
  // flv: FLV format, which combines multiple recorded videos into a single flv file. Default format
  // 
  // mp4: MP4 format, which combines multiple recorded videos into a single mp4 file.
  // 
  // Format not supported:
  // 
  // m3u8: indicates the HLS format", "zh_CN":"文件格式，可选文件格式为：
  // 
  // flv：FLV格式，将多个录制视频合并成单个flv文件。默认格式
  // 
  // mp4：MP4格式，将多个录制视频合并成单个mp4文件。
  // 
  // 不支持格式：
  // 
  // m3u8：HLS格式"}
  Suffix *string `json:"suffix,omitempty" xml:"suffix,omitempty"`
  // {"en":"Callback address. When the task is complete, the callback address is not specified. If the address is not specified, the callback is not performed", "zh_CN":"回调地址。完成任务后回调通知地址，不指定表示不做回调"}
  Notify *string `json:"notify,omitempty" xml:"notify,omitempty"`
  // {"en":"Start forced transcoding, 1: start. 0: Do not start, the default is 0", "zh_CN":"启动强制转码，1:启动。0:不启动 默认是0"}
  EnableVideoTranscode *int32 `json:"enableVideoTranscode,omitempty" xml:"enableVideoTranscode,omitempty"`
}

func (s LiveVideoConcatRequest) String() string {
  return tea.Prettify(s)
}

func (s LiveVideoConcatRequest) GoString() string {
  return s.String()
}

func (s *LiveVideoConcatRequest) SetPullId(v string) *LiveVideoConcatRequest {
  s.PullId = &v
  return s
}

func (s *LiveVideoConcatRequest) SetTransNo(v string) *LiveVideoConcatRequest {
  s.TransNo = &v
  return s
}

func (s *LiveVideoConcatRequest) SetStart(v int64) *LiveVideoConcatRequest {
  s.Start = &v
  return s
}

func (s *LiveVideoConcatRequest) SetEnd(v int64) *LiveVideoConcatRequest {
  s.End = &v
  return s
}

func (s *LiveVideoConcatRequest) SetFname(v string) *LiveVideoConcatRequest {
  s.Fname = &v
  return s
}

func (s *LiveVideoConcatRequest) SetSuffix(v string) *LiveVideoConcatRequest {
  s.Suffix = &v
  return s
}

func (s *LiveVideoConcatRequest) SetNotify(v string) *LiveVideoConcatRequest {
  s.Notify = &v
  return s
}

func (s *LiveVideoConcatRequest) SetEnableVideoTranscode(v int32) *LiveVideoConcatRequest {
  s.EnableVideoTranscode = &v
  return s
}

type LiveVideoConcatResponse struct {
  // {"en":"Status code", "zh_CN":"状态码"}
  Code *int32 `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Operational infomation", "zh_CN":"操作信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Return data", "zh_CN":"返回数据"}
  LiveVideoConcatData *LiveVideoConcatData `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s LiveVideoConcatResponse) String() string {
  return tea.Prettify(s)
}

func (s LiveVideoConcatResponse) GoString() string {
  return s.String()
}

func (s *LiveVideoConcatResponse) SetCode(v int32) *LiveVideoConcatResponse {
  s.Code = &v
  return s
}

func (s *LiveVideoConcatResponse) SetMessage(v string) *LiveVideoConcatResponse {
  s.Message = &v
  return s
}

func (s *LiveVideoConcatResponse) SetData(v *LiveVideoConcatData) *LiveVideoConcatResponse {
  s.LiveVideoConcatData = v
  return s
}

type LiveVideoConcatData struct {
  // {"en":"transNo", "zh_CN":"业务ID"}
  TranNo *string `json:"tranNo,omitempty" xml:"tranNo,omitempty" require:"true"`
}

func (s LiveVideoConcatData) String() string {
  return tea.Prettify(s)
}

func (s LiveVideoConcatData) GoString() string {
  return s.String()
}

func (s *LiveVideoConcatData) SetTranNo(v string) *LiveVideoConcatData {
  s.TranNo = &v
  return s
}

type LiveVideoConcatPaths struct {
}

func (s LiveVideoConcatPaths) String() string {
  return tea.Prettify(s)
}

func (s LiveVideoConcatPaths) GoString() string {
  return s.String()
}

type LiveVideoConcatParameters struct {
}

func (s LiveVideoConcatParameters) String() string {
  return tea.Prettify(s)
}

func (s LiveVideoConcatParameters) GoString() string {
  return s.String()
}

type LiveVideoConcatRequestHeader struct {
}

func (s LiveVideoConcatRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s LiveVideoConcatRequestHeader) GoString() string {
  return s.String()
}

type LiveVideoConcatResponseHeader struct {
}

func (s LiveVideoConcatResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s LiveVideoConcatResponseHeader) GoString() string {
  return s.String()
}




type GetRecordTaskListQueryRequest struct {
  // {"en":"Channel pull ID. Multiple streaming IDs are separated by ,; if not filled in, recording tasks for all channels will be returned by default", "zh_CN":"频道拉流ID。多个拉流id用,隔开；未填写默认返回所有频道的录制任务"}
  PullId *string `json:"pullId,omitempty" xml:"pullId,omitempty"`
  // {"en":"Business ID needs to be uniquely controlled by the user. Use , to separate multiple ones.", "zh_CN":"业务ID，需用户自己控制唯一性。多个用,隔开。"}
  TransNo *string `json:"transNo,omitempty" xml:"transNo,omitempty"`
  // {"en":"The status of the recording task. If not filled in, all will be queried. 0 has not started, 1 has started, 2 has ended normally, and 3 has terminated.", "zh_CN":"录制任务的状态， 不填则查询所有，0未开始，1已开始，2正常结束，3终止"}
  Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
  // {"en":"List order, value range: <br>
  // 0(in descending order by creation time)<br>
  // 1(in ascending order of creation time) The default value is 0", "zh_CN":"列表排列顺序，取值范围 ：<br>
  // 0(按创建时间降序排列)<br>
  // 1(按创建时间升序排列)默认为0"}
  ListOrder *int32 `json:"listOrder,omitempty" xml:"listOrder,omitempty"`
  // {"en":"The page number in the task paging list starts from 1; the default is 1", "zh_CN":"task分页列表中第几页，从1开始取值；默认1"}
  PageIndex *int32 `json:"pageIndex,omitempty" xml:"pageIndex,omitempty"`
  // {"en":"The average number of channels per page, the default is 10, the value is between 1-50.", "zh_CN":"平均每页频道数量，默认为10，取值在1-50之间"}
  PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
  // {"en":"Query starting time, the time format is, 2016-01-01 12:00:00; used to query recording tasks according to the creation time period; if not filled in, query all queries all", "zh_CN":"查询起始时间，时间格式为，2016-01-01 12:00:00；用于按创建时间段查询录制任务；如果不填则查询所有查询所有"}
  StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
  // {"en":"Query deadline, the time format is, 2016-01-01 12:00:00; used to query recording tasks according to the creation time period, which is less than the current query time;. If not filled in, query all query all", "zh_CN":"查询截止时间，时间格式为，2016-01-01 12:00:00；用于按创建时间段查询录制任务，小于当前查询时间；。如果不填则查询所有查询所有"}
  EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
}

func (s GetRecordTaskListQueryRequest) String() string {
  return tea.Prettify(s)
}

func (s GetRecordTaskListQueryRequest) GoString() string {
  return s.String()
}

func (s *GetRecordTaskListQueryRequest) SetPullId(v string) *GetRecordTaskListQueryRequest {
  s.PullId = &v
  return s
}

func (s *GetRecordTaskListQueryRequest) SetTransNo(v string) *GetRecordTaskListQueryRequest {
  s.TransNo = &v
  return s
}

func (s *GetRecordTaskListQueryRequest) SetStatus(v int32) *GetRecordTaskListQueryRequest {
  s.Status = &v
  return s
}

func (s *GetRecordTaskListQueryRequest) SetListOrder(v int32) *GetRecordTaskListQueryRequest {
  s.ListOrder = &v
  return s
}

func (s *GetRecordTaskListQueryRequest) SetPageIndex(v int32) *GetRecordTaskListQueryRequest {
  s.PageIndex = &v
  return s
}

func (s *GetRecordTaskListQueryRequest) SetPageSize(v int32) *GetRecordTaskListQueryRequest {
  s.PageSize = &v
  return s
}

func (s *GetRecordTaskListQueryRequest) SetStartTime(v string) *GetRecordTaskListQueryRequest {
  s.StartTime = &v
  return s
}

func (s *GetRecordTaskListQueryRequest) SetEndTime(v string) *GetRecordTaskListQueryRequest {
  s.EndTime = &v
  return s
}

type GetRecordTaskListQueryResponse struct {
  // {"en":"Status code", "zh_CN":"状态码"}
  Code *int32 `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Operational information", "zh_CN":"操作信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Return data", "zh_CN":"返回数据"}
  Data *GetRecordTaskListQueryRecordTask `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s GetRecordTaskListQueryResponse) String() string {
  return tea.Prettify(s)
}

func (s GetRecordTaskListQueryResponse) GoString() string {
  return s.String()
}

func (s *GetRecordTaskListQueryResponse) SetCode(v int32) *GetRecordTaskListQueryResponse {
  s.Code = &v
  return s
}

func (s *GetRecordTaskListQueryResponse) SetMessage(v string) *GetRecordTaskListQueryResponse {
  s.Message = &v
  return s
}

func (s *GetRecordTaskListQueryResponse) SetData(v *GetRecordTaskListQueryRecordTask) *GetRecordTaskListQueryResponse {
  s.Data = v
  return s
}

type GetRecordTaskListQueryRecordTask struct {
  // {"en":"The number of records of the recording task list information currently returned. Note that the number of records returned here is only the number of records on the current page.", "zh_CN":"当前返回的录制任务列表信息的记录数，注意这里返回的记录数只是当前页的记录数"}
  RecordTaskTotal *int32 `json:"recordTaskTotal,omitempty" xml:"recordTaskTotal,omitempty" require:"true"`
  // {"en":"Recording task list", "zh_CN":"录制任务列表"}
  RecordTaskQueryResponseList []*GetRecordTaskListQueryRecordTaskQuery `json:"recordTaskQueryResponseList,omitempty" xml:"recordTaskQueryResponseList,omitempty" require:"true" type:"Repeated"`
}

func (s GetRecordTaskListQueryRecordTask) String() string {
  return tea.Prettify(s)
}

func (s GetRecordTaskListQueryRecordTask) GoString() string {
  return s.String()
}

func (s *GetRecordTaskListQueryRecordTask) SetRecordTaskTotal(v int32) *GetRecordTaskListQueryRecordTask {
  s.RecordTaskTotal = &v
  return s
}

func (s *GetRecordTaskListQueryRecordTask) SetRecordTaskQueryResponseList(v []*GetRecordTaskListQueryRecordTaskQuery) *GetRecordTaskListQueryRecordTask {
  s.RecordTaskQueryResponseList = v
  return s
}

type GetRecordTaskListQueryRecordTaskQuery struct {
  // {"en":"Task Id", "zh_CN":"直播录制任务ID"}
  TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty" require:"true"`
  // {"en":"Channel streaming ID", "zh_CN":"频道拉流ID"}
  PullId *string `json:"pullId,omitempty" xml:"pullId,omitempty" require:"true"`
  // {"en":"The status of the recording task, 0 has not started, 1 has started, 2 has ended normally, and 3 has terminated", "zh_CN":"录制任务的状态， 0未开始，1已开始，2正常结束，3终止"}
  Status *int32 `json:"status,omitempty" xml:"status,omitempty" require:"true"`
  // {"en":"Custom task number passed in by the customer", "zh_CN":"客户传入的自定义任务编号"}
  TransNo *string `json:"transNo,omitempty" xml:"transNo,omitempty" require:"true"`
  // {"en":"Recording file format", "zh_CN":"录制文件的格式"}
  FileType *string `json:"fileType,omitempty" xml:"fileType,omitempty" require:"true"`
  // {"en":"Whether to merge into one video", "zh_CN":"是否合并成一个视频"}
  IsConcat *bool `json:"isConcat,omitempty" xml:"isConcat,omitempty" require:"true"`
  // {"en":"Task start time", "zh_CN":"任务开始时间"}
  StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty" require:"true"`
  // {"en":"Task end time", "zh_CN":"任务结束时间"}
  EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty" require:"true"`
}

func (s GetRecordTaskListQueryRecordTaskQuery) String() string {
  return tea.Prettify(s)
}

func (s GetRecordTaskListQueryRecordTaskQuery) GoString() string {
  return s.String()
}

func (s *GetRecordTaskListQueryRecordTaskQuery) SetTaskId(v string) *GetRecordTaskListQueryRecordTaskQuery {
  s.TaskId = &v
  return s
}

func (s *GetRecordTaskListQueryRecordTaskQuery) SetPullId(v string) *GetRecordTaskListQueryRecordTaskQuery {
  s.PullId = &v
  return s
}

func (s *GetRecordTaskListQueryRecordTaskQuery) SetStatus(v int32) *GetRecordTaskListQueryRecordTaskQuery {
  s.Status = &v
  return s
}

func (s *GetRecordTaskListQueryRecordTaskQuery) SetTransNo(v string) *GetRecordTaskListQueryRecordTaskQuery {
  s.TransNo = &v
  return s
}

func (s *GetRecordTaskListQueryRecordTaskQuery) SetFileType(v string) *GetRecordTaskListQueryRecordTaskQuery {
  s.FileType = &v
  return s
}

func (s *GetRecordTaskListQueryRecordTaskQuery) SetIsConcat(v bool) *GetRecordTaskListQueryRecordTaskQuery {
  s.IsConcat = &v
  return s
}

func (s *GetRecordTaskListQueryRecordTaskQuery) SetStartTime(v string) *GetRecordTaskListQueryRecordTaskQuery {
  s.StartTime = &v
  return s
}

func (s *GetRecordTaskListQueryRecordTaskQuery) SetEndTime(v string) *GetRecordTaskListQueryRecordTaskQuery {
  s.EndTime = &v
  return s
}

type GetRecordTaskListQueryPaths struct {
}

func (s GetRecordTaskListQueryPaths) String() string {
  return tea.Prettify(s)
}

func (s GetRecordTaskListQueryPaths) GoString() string {
  return s.String()
}

type GetRecordTaskListQueryParameters struct {
}

func (s GetRecordTaskListQueryParameters) String() string {
  return tea.Prettify(s)
}

func (s GetRecordTaskListQueryParameters) GoString() string {
  return s.String()
}

type GetRecordTaskListQueryRequestHeader struct {
}

func (s GetRecordTaskListQueryRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetRecordTaskListQueryRequestHeader) GoString() string {
  return s.String()
}

type GetRecordTaskListQueryResponseHeader struct {
}

func (s GetRecordTaskListQueryResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetRecordTaskListQueryResponseHeader) GoString() string {
  return s.String()
}




type LiveVideoConcatQueryRequest struct {
  // {"en":"Service ID, service ID and channel pull id cannot be empty at the same time", "zh_CN":"业务ID，业务ID和频道拉流id不可以同时为空"}
  TransNo *string `json:"transNo,omitempty" xml:"transNo,omitempty"`
  // {"en":"Channel pull id, service ID, and channel pull id cannot be empty at the same time", "zh_CN":"频道拉流id，业务ID和频道拉流id不可以同时为空"}
  PullId *string `json:"pullId,omitempty" xml:"pullId,omitempty"`
  // {"en":"List order, value range:
  // 0(in descending order of creation time)
  // 1(in ascending order of creation time)
  // Default is 0", "zh_CN":"列表排列顺序，取值范围 ：
  // 0(按创建时间降序排列)
  // 1(按创建时间升序排列)
  // 默认为0"}
  ListOrder *int32 `json:"listOrder,omitempty" xml:"listOrder,omitempty"`
  // {"en":"On the page of the scenario list, the value starts from 1. The default value is 1. The product of pageIndex and pageSize must be less than 100000.", "zh_CN":"取场景列表第几页，从1开始取值,默认为1。入参pageIndex和pageSize的乘积必须不大于100000。"}
  PageIndex *int32 `json:"pageIndex,omitempty" xml:"pageIndex,omitempty"`
  // {"en":"Average Number of scenarios per page. The value ranges from 1 to 50. The default value is 10. The product of pageIndex and pageSize must be less than 100000.", "zh_CN":"平均每页场景数量，取值范围1-50，默认为10。入参pageIndex和pageSize的乘积必须不大于100000。"}
  PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s LiveVideoConcatQueryRequest) String() string {
  return tea.Prettify(s)
}

func (s LiveVideoConcatQueryRequest) GoString() string {
  return s.String()
}

func (s *LiveVideoConcatQueryRequest) SetTransNo(v string) *LiveVideoConcatQueryRequest {
  s.TransNo = &v
  return s
}

func (s *LiveVideoConcatQueryRequest) SetPullId(v string) *LiveVideoConcatQueryRequest {
  s.PullId = &v
  return s
}

func (s *LiveVideoConcatQueryRequest) SetListOrder(v int32) *LiveVideoConcatQueryRequest {
  s.ListOrder = &v
  return s
}

func (s *LiveVideoConcatQueryRequest) SetPageIndex(v int32) *LiveVideoConcatQueryRequest {
  s.PageIndex = &v
  return s
}

func (s *LiveVideoConcatQueryRequest) SetPageSize(v int32) *LiveVideoConcatQueryRequest {
  s.PageSize = &v
  return s
}

type LiveVideoConcatQueryResponse struct {
  // {"en":"Result status code, 200 indicates success", "zh_CN":"结果状态码，200为成功"}
  Code *int32 `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Return message", "zh_CN":"返回消息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Return data", "zh_CN":"返回数据"}
  LiveVideoConcatQueryData []*LiveVideoConcatQueryData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s LiveVideoConcatQueryResponse) String() string {
  return tea.Prettify(s)
}

func (s LiveVideoConcatQueryResponse) GoString() string {
  return s.String()
}

func (s *LiveVideoConcatQueryResponse) SetCode(v int32) *LiveVideoConcatQueryResponse {
  s.Code = &v
  return s
}

func (s *LiveVideoConcatQueryResponse) SetMessage(v string) *LiveVideoConcatQueryResponse {
  s.Message = &v
  return s
}

func (s *LiveVideoConcatQueryResponse) SetData(v []*LiveVideoConcatQueryData) *LiveVideoConcatQueryResponse {
  s.LiveVideoConcatQueryData = v
  return s
}

type LiveVideoConcatQueryData struct {
  // {"en":"Task id", "zh_CN":"任务ID"}
  TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty" require:"true"`
  // {"en":"transNo", "zh_CN":"业务ID"}
  TranNo *string `json:"tranNo,omitempty" xml:"tranNo,omitempty" require:"true"`
  // {"en":"status", "zh_CN":"状态
  // 取值范围：
  // 1：处理中
  // 2：成功
  // 3：失败"}
  Status *int32 `json:"status,omitempty" xml:"status,omitempty" require:"true"`
  // {"en":"Video details; This is only returned if statu is equal to 2", "zh_CN":"视频详情；只要statu = 2的时候才会返回这个值"}
  LiveVideoConcatQueryVideoInfo *LiveVideoConcatQueryVideoInfo `json:"videoInfo,omitempty" xml:"videoInfo,omitempty" require:"true"`
}

func (s LiveVideoConcatQueryData) String() string {
  return tea.Prettify(s)
}

func (s LiveVideoConcatQueryData) GoString() string {
  return s.String()
}

func (s *LiveVideoConcatQueryData) SetTaskId(v string) *LiveVideoConcatQueryData {
  s.TaskId = &v
  return s
}

func (s *LiveVideoConcatQueryData) SetTranNo(v string) *LiveVideoConcatQueryData {
  s.TranNo = &v
  return s
}

func (s *LiveVideoConcatQueryData) SetStatus(v int32) *LiveVideoConcatQueryData {
  s.Status = &v
  return s
}

func (s *LiveVideoConcatQueryData) SetVideoInfo(v *LiveVideoConcatQueryVideoInfo) *LiveVideoConcatQueryData {
  s.LiveVideoConcatQueryVideoInfo = v
  return s
}

type LiveVideoConcatQueryVideoInfo struct {
  // {"en":"Video name", "zh_CN":"视频名称"}
  VideoName *string `json:"videoName,omitempty" xml:"videoName,omitempty" require:"true"`
  // {"en":"Video id", "zh_CN":"视频ID"}
  VideoId *string `json:"videoId,omitempty" xml:"videoId,omitempty" require:"true"`
  // {"en":"Whether to encrypt transcoding files
  // Value range: 0(unencrypted), 1(encrypted)", "zh_CN":"是否加密转码文件
  // 取值范围 ：0(不加密)，1(加密)"}
  Encrypt *int32 `json:"encrypt,omitempty" xml:"encrypt,omitempty" require:"true"`
  // {"en":"The space occupied by the video, and the total space used by the video and its transcoding", "zh_CN":"视频占用空间大小，视频及其转码后视频的总空间使用量"}
  VideoSize *string `json:"videoSize,omitempty" xml:"videoSize,omitempty" require:"true"`
  // {"en":"Video duration", "zh_CN":"视频时长"}
  VideoDuration *string `json:"videoDuration,omitempty" xml:"videoDuration,omitempty" require:"true"`
  // {"en":"Video upload time", "zh_CN":"视频上传时间"}
  UploadTime *string `json:"uploadTime,omitempty" xml:"uploadTime,omitempty" require:"true"`
  // {"en":"Video modification time", "zh_CN":"视频修改时间"}
  UpdateTime *string `json:"updateTime,omitempty" xml:"updateTime,omitempty" require:"true"`
  // {"en":"Video description", "zh_CN":"视频描述"}
  VideoDescription *string `json:"videoDescription,omitempty" xml:"videoDescription,omitempty" require:"true"`
  // {"en":"Video classification", "zh_CN":"视频分类"}
  VideoClassification *string `json:"videoClassification,omitempty" xml:"videoClassification,omitempty" require:"true"`
  // {"en":"Video cover URL", "zh_CN":"视频封面URL"}
  ImageUrl *string `json:"imageUrl,omitempty" xml:"imageUrl,omitempty" require:"true"`
  // {"en":"The domain name of the video", "zh_CN":"视频的发布域名"}
  PublishDomain *string `json:"publishDomain,omitempty" xml:"publishDomain,omitempty" require:"true"`
  // {"en":"Name of the player used by the video", "zh_CN":"视频使用的播放器名称"}
  PlayerName *string `json:"playerName,omitempty" xml:"playerName,omitempty" require:"true"`
  // {"en":"The player ID used by the video", "zh_CN":"视频使用的播放器ID"}
  PlayerId *string `json:"playerId,omitempty" xml:"playerId,omitempty" require:"true"`
  // {"en":"Video state
  // Value range: 0(normal), 1(masked)", "zh_CN":"视频状态
  // 取值范围：0(正常)，1(屏蔽)"}
  VideoState *string `json:"videoState,omitempty" xml:"videoState,omitempty" require:"true"`
  // {"en":"If authorized play is not enabled, the video transcoding status ranges from:
  // 1(transcoding), 2(not transcoding), 3(transcoding), 4(transcoding failed)
  // Value range of transcoding status when the Authorized Play (video encryption) function is enabled:
  // 1(encrypted transcoding), 2(unencrypted transcoding), 3(in transcoding), 4(transcoding failed), 5(not transcoding)", "zh_CN":"未开启授权播放，视频的转码状态的取值范围 ：
  // 1(已转码)，2(未转码)，3(转码中)，4(转码失败)
  // 开启授权播放（视频加密）功能时的转码状态的取值范围 ：
  // 1(已加密转码)，2(非加密转码)，3(转码中)，4(转码失败)，5(未转码)"}
  TranscodeState *string `json:"transcodeState,omitempty" xml:"transcodeState,omitempty" require:"true"`
  // {"en":"Video source
  // Value range:
  // 0(other), 1(upload), 2 (live to record), 3 (video pull), 4 (video cut), 5 (video splicing), 6 (edge pull to record), 10 (general version live to record), 11 (upload SDK), 12 (upload tool)", "zh_CN":"视频来源
  // 取值范围：
  // 0(其他)，1(上传)，2（直播转录制），3（视频拉取），4（视频剪切），5（视频拼接），6（边缘拉流录制），10（通用版直播转录制），11（上传SDK），12（上传工具）"}
  VideoSourceCode *int32 `json:"videoSourceCode,omitempty" xml:"videoSourceCode,omitempty" require:"true"`
  // {"en":"Video resolution and other information", "zh_CN":"视频分辨率等信息"}
  LiveVideoConcatQueryVideoResolutions []*LiveVideoConcatQueryVideoResolutions `json:"videoResolutions,omitempty" xml:"videoResolutions,omitempty" require:"true" type:"Repeated"`
}

func (s LiveVideoConcatQueryVideoInfo) String() string {
  return tea.Prettify(s)
}

func (s LiveVideoConcatQueryVideoInfo) GoString() string {
  return s.String()
}

func (s *LiveVideoConcatQueryVideoInfo) SetVideoName(v string) *LiveVideoConcatQueryVideoInfo {
  s.VideoName = &v
  return s
}

func (s *LiveVideoConcatQueryVideoInfo) SetVideoId(v string) *LiveVideoConcatQueryVideoInfo {
  s.VideoId = &v
  return s
}

func (s *LiveVideoConcatQueryVideoInfo) SetEncrypt(v int32) *LiveVideoConcatQueryVideoInfo {
  s.Encrypt = &v
  return s
}

func (s *LiveVideoConcatQueryVideoInfo) SetVideoSize(v string) *LiveVideoConcatQueryVideoInfo {
  s.VideoSize = &v
  return s
}

func (s *LiveVideoConcatQueryVideoInfo) SetVideoDuration(v string) *LiveVideoConcatQueryVideoInfo {
  s.VideoDuration = &v
  return s
}

func (s *LiveVideoConcatQueryVideoInfo) SetUploadTime(v string) *LiveVideoConcatQueryVideoInfo {
  s.UploadTime = &v
  return s
}

func (s *LiveVideoConcatQueryVideoInfo) SetUpdateTime(v string) *LiveVideoConcatQueryVideoInfo {
  s.UpdateTime = &v
  return s
}

func (s *LiveVideoConcatQueryVideoInfo) SetVideoDescription(v string) *LiveVideoConcatQueryVideoInfo {
  s.VideoDescription = &v
  return s
}

func (s *LiveVideoConcatQueryVideoInfo) SetVideoClassification(v string) *LiveVideoConcatQueryVideoInfo {
  s.VideoClassification = &v
  return s
}

func (s *LiveVideoConcatQueryVideoInfo) SetImageUrl(v string) *LiveVideoConcatQueryVideoInfo {
  s.ImageUrl = &v
  return s
}

func (s *LiveVideoConcatQueryVideoInfo) SetPublishDomain(v string) *LiveVideoConcatQueryVideoInfo {
  s.PublishDomain = &v
  return s
}

func (s *LiveVideoConcatQueryVideoInfo) SetPlayerName(v string) *LiveVideoConcatQueryVideoInfo {
  s.PlayerName = &v
  return s
}

func (s *LiveVideoConcatQueryVideoInfo) SetPlayerId(v string) *LiveVideoConcatQueryVideoInfo {
  s.PlayerId = &v
  return s
}

func (s *LiveVideoConcatQueryVideoInfo) SetVideoState(v string) *LiveVideoConcatQueryVideoInfo {
  s.VideoState = &v
  return s
}

func (s *LiveVideoConcatQueryVideoInfo) SetTranscodeState(v string) *LiveVideoConcatQueryVideoInfo {
  s.TranscodeState = &v
  return s
}

func (s *LiveVideoConcatQueryVideoInfo) SetVideoSourceCode(v int32) *LiveVideoConcatQueryVideoInfo {
  s.VideoSourceCode = &v
  return s
}

func (s *LiveVideoConcatQueryVideoInfo) SetVideoResolutions(v []*LiveVideoConcatQueryVideoResolutions) *LiveVideoConcatQueryVideoInfo {
  s.LiveVideoConcatQueryVideoResolutions = v
  return s
}

type LiveVideoConcatQueryVideoResolutions struct {
  // {"en":"Clarity. Value range: 1(original painting), 2(smooth), 3(standard definition), 4(HD), 5(ultra HD), -99(recorded file)", "zh_CN":"清晰度。取值范围 ：1(原画)，2(流畅)，3(标清)，4(高清)，5(超清)，-99(录制文件)"}
  Clarity *int32 `json:"clarity,omitempty" xml:"clarity,omitempty" require:"true"`
  // {"en":"Terminal type. Value range: 0(PC), 1(original video)", "zh_CN":"终端类型。取值范围 ：0(PC)，1(原视频)"}
  ServerType *int32 `json:"serverType,omitempty" xml:"serverType,omitempty" require:"true"`
  // {"en":"Height", "zh_CN":"高度"}
  Height *int32 `json:"height,omitempty" xml:"height,omitempty" require:"true"`
  // {"en":"Width", "zh_CN":"宽度"}
  Width *int32 `json:"width,omitempty" xml:"width,omitempty" require:"true"`
  // {"en":"File size(bit)", "zh_CN":"文件大小(单位为bit)"}
  FileSize *int64 `json:"fileSize,omitempty" xml:"fileSize,omitempty" require:"true"`
}

func (s LiveVideoConcatQueryVideoResolutions) String() string {
  return tea.Prettify(s)
}

func (s LiveVideoConcatQueryVideoResolutions) GoString() string {
  return s.String()
}

func (s *LiveVideoConcatQueryVideoResolutions) SetClarity(v int32) *LiveVideoConcatQueryVideoResolutions {
  s.Clarity = &v
  return s
}

func (s *LiveVideoConcatQueryVideoResolutions) SetServerType(v int32) *LiveVideoConcatQueryVideoResolutions {
  s.ServerType = &v
  return s
}

func (s *LiveVideoConcatQueryVideoResolutions) SetHeight(v int32) *LiveVideoConcatQueryVideoResolutions {
  s.Height = &v
  return s
}

func (s *LiveVideoConcatQueryVideoResolutions) SetWidth(v int32) *LiveVideoConcatQueryVideoResolutions {
  s.Width = &v
  return s
}

func (s *LiveVideoConcatQueryVideoResolutions) SetFileSize(v int64) *LiveVideoConcatQueryVideoResolutions {
  s.FileSize = &v
  return s
}

type LiveVideoConcatQueryPaths struct {
}

func (s LiveVideoConcatQueryPaths) String() string {
  return tea.Prettify(s)
}

func (s LiveVideoConcatQueryPaths) GoString() string {
  return s.String()
}

type LiveVideoConcatQueryParameters struct {
}

func (s LiveVideoConcatQueryParameters) String() string {
  return tea.Prettify(s)
}

func (s LiveVideoConcatQueryParameters) GoString() string {
  return s.String()
}

type LiveVideoConcatQueryRequestHeader struct {
}

func (s LiveVideoConcatQueryRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s LiveVideoConcatQueryRequestHeader) GoString() string {
  return s.String()
}

type LiveVideoConcatQueryResponseHeader struct {
}

func (s LiveVideoConcatQueryResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s LiveVideoConcatQueryResponseHeader) GoString() string {
  return s.String()
}




