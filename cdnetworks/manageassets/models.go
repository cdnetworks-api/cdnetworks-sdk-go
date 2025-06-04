package manageassets

// This file is auto-generated, don't edit it. Thanks.
import (
  "github.com/alibabacloud-go/tea/tea"
)

type VideoEnableRequest struct {
  // {"en":"The video id to enable", "zh_CN":"要启用的视频id"}
  VideoId *string `json:"videoId,omitempty" xml:"videoId,omitempty" require:"true"`
}

func (s VideoEnableRequest) String() string {
  return tea.Prettify(s)
}

func (s VideoEnableRequest) GoString() string {
  return s.String()
}

func (s *VideoEnableRequest) SetVideoId(v string) *VideoEnableRequest {
  s.VideoId = &v
  return s
}

type VideoEnableResponse struct {
  // {"en":"Result status code, 200 indicates success", "zh_CN":"结果状态码，200为成功"}
  Code *int32 `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Return message", "zh_CN":"返回消息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Return data", "zh_CN":"返回数据"}
  Data *string `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s VideoEnableResponse) String() string {
  return tea.Prettify(s)
}

func (s VideoEnableResponse) GoString() string {
  return s.String()
}

func (s *VideoEnableResponse) SetCode(v int32) *VideoEnableResponse {
  s.Code = &v
  return s
}

func (s *VideoEnableResponse) SetMessage(v string) *VideoEnableResponse {
  s.Message = &v
  return s
}

func (s *VideoEnableResponse) SetData(v string) *VideoEnableResponse {
  s.Data = &v
  return s
}

type VideoEnablePaths struct {
}

func (s VideoEnablePaths) String() string {
  return tea.Prettify(s)
}

func (s VideoEnablePaths) GoString() string {
  return s.String()
}

type VideoEnableParameters struct {
}

func (s VideoEnableParameters) String() string {
  return tea.Prettify(s)
}

func (s VideoEnableParameters) GoString() string {
  return s.String()
}

type VideoEnableRequestHeader struct {
}

func (s VideoEnableRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s VideoEnableRequestHeader) GoString() string {
  return s.String()
}

type VideoEnableResponseHeader struct {
}

func (s VideoEnableResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s VideoEnableResponseHeader) GoString() string {
  return s.String()
}




type DeleteCategoryRequest struct {
  // {"en":"Node id", "zh_CN":"节点Id"}
  NodeId *int32 `json:"nodeId,omitempty" xml:"nodeId,omitempty" require:"true"`
}

func (s DeleteCategoryRequest) String() string {
  return tea.Prettify(s)
}

func (s DeleteCategoryRequest) GoString() string {
  return s.String()
}

func (s *DeleteCategoryRequest) SetNodeId(v int32) *DeleteCategoryRequest {
  s.NodeId = &v
  return s
}

type DeleteCategoryResponse struct {
  // {"en":"Result status code, 200 indicates success", "zh_CN":"结果状态码，200为成功"}
  Code *int32 `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Return message", "zh_CN":"返回消息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s DeleteCategoryResponse) String() string {
  return tea.Prettify(s)
}

func (s DeleteCategoryResponse) GoString() string {
  return s.String()
}

func (s *DeleteCategoryResponse) SetCode(v int32) *DeleteCategoryResponse {
  s.Code = &v
  return s
}

func (s *DeleteCategoryResponse) SetMessage(v string) *DeleteCategoryResponse {
  s.Message = &v
  return s
}

type DeleteCategoryPaths struct {
}

func (s DeleteCategoryPaths) String() string {
  return tea.Prettify(s)
}

func (s DeleteCategoryPaths) GoString() string {
  return s.String()
}

type DeleteCategoryParameters struct {
}

func (s DeleteCategoryParameters) String() string {
  return tea.Prettify(s)
}

func (s DeleteCategoryParameters) GoString() string {
  return s.String()
}

type DeleteCategoryRequestHeader struct {
}

func (s DeleteCategoryRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteCategoryRequestHeader) GoString() string {
  return s.String()
}

type DeleteCategoryResponseHeader struct {
}

func (s DeleteCategoryResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteCategoryResponseHeader) GoString() string {
  return s.String()
}




type GetCategoryListRequest struct {
  // {"en":"Id of the parent node. If empty, finds all first-level category tags", "zh_CN":"父节点Id。如果为空则查找所有一级分类标签"}
  ParentNodeId *string `json:"parentNodeId,omitempty" xml:"parentNodeId,omitempty"`
  // {"en":"Number of items per page The value ranges from 1 to 100. The default value is 10", "zh_CN":"每页数量条数数，取值范围1-50，默认为10"}
  PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
  // {"en":"Index of page The value starts from 1. The default value is 1", "zh_CN":"取数据第几页，从1开始取值,默认为1"}
  PageIndex *int32 `json:"pageIndex,omitempty" xml:"pageIndex,omitempty"`
}

func (s GetCategoryListRequest) String() string {
  return tea.Prettify(s)
}

func (s GetCategoryListRequest) GoString() string {
  return s.String()
}

func (s *GetCategoryListRequest) SetParentNodeId(v string) *GetCategoryListRequest {
  s.ParentNodeId = &v
  return s
}

func (s *GetCategoryListRequest) SetPageSize(v int32) *GetCategoryListRequest {
  s.PageSize = &v
  return s
}

func (s *GetCategoryListRequest) SetPageIndex(v int32) *GetCategoryListRequest {
  s.PageIndex = &v
  return s
}

type GetCategoryListResponse struct {
  // {"en":"Result status code, 200 indicates success", "zh_CN":"结果状态码，200为成功"}
  Code *int32 `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Return message", "zh_CN":"返回消息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Return data", "zh_CN":"返回数据"}
  GetCategoryListData []*GetCategoryListData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s GetCategoryListResponse) String() string {
  return tea.Prettify(s)
}

func (s GetCategoryListResponse) GoString() string {
  return s.String()
}

func (s *GetCategoryListResponse) SetCode(v int32) *GetCategoryListResponse {
  s.Code = &v
  return s
}

func (s *GetCategoryListResponse) SetMessage(v string) *GetCategoryListResponse {
  s.Message = &v
  return s
}

func (s *GetCategoryListResponse) SetData(v []*GetCategoryListData) *GetCategoryListResponse {
  s.GetCategoryListData = v
  return s
}

type GetCategoryListData struct {
  // {"en":"Total number of data items", "zh_CN":"总数据条数"}
  Total *int64 `json:"total,omitempty" xml:"total,omitempty" require:"true"`
  // {"en":"Returns the number of data bars", "zh_CN":"返回数据条数"}
  PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty" require:"true"`
  // {"en":"Return to page number", "zh_CN":"返回第几页"}
  PageIndex *int64 `json:"pageIndex,omitempty" xml:"pageIndex,omitempty" require:"true"`
  // {"en":"GetCategoryListData detail list", "zh_CN":"数据明细列表"}
  CategoryList []*GetCategoryListCategory `json:"categoryList,omitempty" xml:"categoryList,omitempty" require:"true" type:"Repeated"`
}

func (s GetCategoryListData) String() string {
  return tea.Prettify(s)
}

func (s GetCategoryListData) GoString() string {
  return s.String()
}

func (s *GetCategoryListData) SetTotal(v int64) *GetCategoryListData {
  s.Total = &v
  return s
}

func (s *GetCategoryListData) SetPageSize(v int64) *GetCategoryListData {
  s.PageSize = &v
  return s
}

func (s *GetCategoryListData) SetPageIndex(v int64) *GetCategoryListData {
  s.PageIndex = &v
  return s
}

func (s *GetCategoryListData) SetCategoryList(v []*GetCategoryListCategory) *GetCategoryListData {
  s.CategoryList = v
  return s
}

type GetCategoryListCategory struct {
  // {"en":"Label name", "zh_CN":"标签名称"}
  LabelName *string `json:"labelName,omitempty" xml:"labelName,omitempty" require:"true"`
  // {"en":"ID of the upper-layer node. A Tier 1 node has no upper-level node. Only secondary nodes have parent nodes", "zh_CN":"上级节点ID。一级节点无上级节点。只有二级节点才有上级节点"}
  ParentNodeId *int32 `json:"parentNodeId,omitempty" xml:"parentNodeId,omitempty" require:"true"`
  // {"en":"Tag node ID", "zh_CN":"标签节点ID"}
  NodeId *int32 `json:"nodeId,omitempty" xml:"nodeId,omitempty" require:"true"`
  // {"en":"Create a user", "zh_CN":"创建用户"}
  CreateUser *string `json:"createUser,omitempty" xml:"createUser,omitempty" require:"true"`
}

func (s GetCategoryListCategory) String() string {
  return tea.Prettify(s)
}

func (s GetCategoryListCategory) GoString() string {
  return s.String()
}

func (s *GetCategoryListCategory) SetLabelName(v string) *GetCategoryListCategory {
  s.LabelName = &v
  return s
}

func (s *GetCategoryListCategory) SetParentNodeId(v int32) *GetCategoryListCategory {
  s.ParentNodeId = &v
  return s
}

func (s *GetCategoryListCategory) SetNodeId(v int32) *GetCategoryListCategory {
  s.NodeId = &v
  return s
}

func (s *GetCategoryListCategory) SetCreateUser(v string) *GetCategoryListCategory {
  s.CreateUser = &v
  return s
}

type GetCategoryListPaths struct {
}

func (s GetCategoryListPaths) String() string {
  return tea.Prettify(s)
}

func (s GetCategoryListPaths) GoString() string {
  return s.String()
}

type GetCategoryListParameters struct {
}

func (s GetCategoryListParameters) String() string {
  return tea.Prettify(s)
}

func (s GetCategoryListParameters) GoString() string {
  return s.String()
}

type GetCategoryListRequestHeader struct {
}

func (s GetCategoryListRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetCategoryListRequestHeader) GoString() string {
  return s.String()
}

type GetCategoryListResponseHeader struct {
}

func (s GetCategoryListResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetCategoryListResponseHeader) GoString() string {
  return s.String()
}




type DeleteVideoByCategoryRequest struct {
  // {"en":"Node id", "zh_CN":"节点Id"}
  NodeId *int32 `json:"nodeId,omitempty" xml:"nodeId,omitempty" require:"true"`
  // {"en":"Delete the result callback notification address", "zh_CN":"删除结果回调通知地址"}
  NotifyUrl *string `json:"notifyUrl,omitempty" xml:"notifyUrl,omitempty"`
}

func (s DeleteVideoByCategoryRequest) String() string {
  return tea.Prettify(s)
}

func (s DeleteVideoByCategoryRequest) GoString() string {
  return s.String()
}

func (s *DeleteVideoByCategoryRequest) SetNodeId(v int32) *DeleteVideoByCategoryRequest {
  s.NodeId = &v
  return s
}

func (s *DeleteVideoByCategoryRequest) SetNotifyUrl(v string) *DeleteVideoByCategoryRequest {
  s.NotifyUrl = &v
  return s
}

type DeleteVideoByCategoryResponse struct {
  // {"en":"Result status code, 200 indicates success", "zh_CN":"结果状态码，200为成功"}
  Code *int32 `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Return message", "zh_CN":"返回消息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Return data", "zh_CN":"返回数据"}
  DeleteVideoByCategoryData []*DeleteVideoByCategoryData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s DeleteVideoByCategoryResponse) String() string {
  return tea.Prettify(s)
}

func (s DeleteVideoByCategoryResponse) GoString() string {
  return s.String()
}

func (s *DeleteVideoByCategoryResponse) SetCode(v int32) *DeleteVideoByCategoryResponse {
  s.Code = &v
  return s
}

func (s *DeleteVideoByCategoryResponse) SetMessage(v string) *DeleteVideoByCategoryResponse {
  s.Message = &v
  return s
}

func (s *DeleteVideoByCategoryResponse) SetData(v []*DeleteVideoByCategoryData) *DeleteVideoByCategoryResponse {
  s.DeleteVideoByCategoryData = v
  return s
}

type DeleteVideoByCategoryData struct {
  // {"en":"Delete task ID", "zh_CN":"删除任务ID"}
  TaskId *int64 `json:"taskId,omitempty" xml:"taskId,omitempty" require:"true"`
}

func (s DeleteVideoByCategoryData) String() string {
  return tea.Prettify(s)
}

func (s DeleteVideoByCategoryData) GoString() string {
  return s.String()
}

func (s *DeleteVideoByCategoryData) SetTaskId(v int64) *DeleteVideoByCategoryData {
  s.TaskId = &v
  return s
}

type DeleteVideoByCategoryPaths struct {
}

func (s DeleteVideoByCategoryPaths) String() string {
  return tea.Prettify(s)
}

func (s DeleteVideoByCategoryPaths) GoString() string {
  return s.String()
}

type DeleteVideoByCategoryParameters struct {
}

func (s DeleteVideoByCategoryParameters) String() string {
  return tea.Prettify(s)
}

func (s DeleteVideoByCategoryParameters) GoString() string {
  return s.String()
}

type DeleteVideoByCategoryRequestHeader struct {
}

func (s DeleteVideoByCategoryRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteVideoByCategoryRequestHeader) GoString() string {
  return s.String()
}

type DeleteVideoByCategoryResponseHeader struct {
}

func (s DeleteVideoByCategoryResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteVideoByCategoryResponseHeader) GoString() string {
  return s.String()
}




type GetMaterialListRequest struct {
  // {"en":"GetMaterialListMaterial ID", "zh_CN":"素材ID"}
  MaterialId *string `json:"materialId,omitempty" xml:"materialId,omitempty"`
  // {"en":"GetMaterialListMaterial name (fuzzy query)", "zh_CN":"素材名称（模糊查询）"}
  MaterialName *string `json:"materialName,omitempty" xml:"materialName,omitempty"`
  // {"en":"GetMaterialListMaterial suffix", "zh_CN":"素材后缀"}
  Suffix *string `json:"suffix,omitempty" xml:"suffix,omitempty"`
  // {"en":"Create a user", "zh_CN":"创建用户"}
  CreateUser *string `json:"createUser,omitempty" xml:"createUser,omitempty"`
  // {"en":"Publish domain name", "zh_CN":"发布域名"}
  PublishDomain *string `json:"publishDomain,omitempty" xml:"publishDomain,omitempty"`
  // {"en":"The query start time is based on the upload time. The query end time cannot be later than the timestamp in seconds", "zh_CN":"查询开始时间,根据上传时间查询，秒级时间戳，不得晚查询结束时间"}
  UploadTimeStart *int32 `json:"uploadTimeStart,omitempty" xml:"uploadTimeStart,omitempty"`
  // {"en":"The query end time is based on the upload time. The timestamp is in seconds and cannot be earlier than the query start time", "zh_CN":"查询结束时间,根据上传时间查询，秒级时间戳，不得早于查询开始时间"}
  UploadTimeEnd *int32 `json:"uploadTimeEnd,omitempty" xml:"uploadTimeEnd,omitempty"`
  // {"en":"The page in the material list starts with 1. The default value is 1", "zh_CN":"取素材列表第几页，从1开始取值,默认为1"}
  PageIndex *string `json:"pageIndex,omitempty" xml:"pageIndex,omitempty"`
  // {"en":"Average Number of materials per page. The value ranges from 1 to 50. The default value is 10", "zh_CN":"平均每页素材数量，取值范围1-50，默认为10"}
  PageSize *string `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s GetMaterialListRequest) String() string {
  return tea.Prettify(s)
}

func (s GetMaterialListRequest) GoString() string {
  return s.String()
}

func (s *GetMaterialListRequest) SetMaterialId(v string) *GetMaterialListRequest {
  s.MaterialId = &v
  return s
}

func (s *GetMaterialListRequest) SetMaterialName(v string) *GetMaterialListRequest {
  s.MaterialName = &v
  return s
}

func (s *GetMaterialListRequest) SetSuffix(v string) *GetMaterialListRequest {
  s.Suffix = &v
  return s
}

func (s *GetMaterialListRequest) SetCreateUser(v string) *GetMaterialListRequest {
  s.CreateUser = &v
  return s
}

func (s *GetMaterialListRequest) SetPublishDomain(v string) *GetMaterialListRequest {
  s.PublishDomain = &v
  return s
}

func (s *GetMaterialListRequest) SetUploadTimeStart(v int32) *GetMaterialListRequest {
  s.UploadTimeStart = &v
  return s
}

func (s *GetMaterialListRequest) SetUploadTimeEnd(v int32) *GetMaterialListRequest {
  s.UploadTimeEnd = &v
  return s
}

func (s *GetMaterialListRequest) SetPageIndex(v string) *GetMaterialListRequest {
  s.PageIndex = &v
  return s
}

func (s *GetMaterialListRequest) SetPageSize(v string) *GetMaterialListRequest {
  s.PageSize = &v
  return s
}

type GetMaterialListResponse struct {
  // {"en":"Result status code, 200 indicates success", "zh_CN":"结果状态码，200为成功"}
  Code *int32 `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Return message", "zh_CN":"返回消息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Return data", "zh_CN":"返回数据"}
  GetMaterialListData *GetMaterialListData `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s GetMaterialListResponse) String() string {
  return tea.Prettify(s)
}

func (s GetMaterialListResponse) GoString() string {
  return s.String()
}

func (s *GetMaterialListResponse) SetCode(v int32) *GetMaterialListResponse {
  s.Code = &v
  return s
}

func (s *GetMaterialListResponse) SetMessage(v string) *GetMaterialListResponse {
  s.Message = &v
  return s
}

func (s *GetMaterialListResponse) SetData(v *GetMaterialListData) *GetMaterialListResponse {
  s.GetMaterialListData = v
  return s
}

type GetMaterialListData struct {
  // {"en":"Total number of materials that match search conditions", "zh_CN":"符合查询条件的素材总个数"}
  Total *int32 `json:"total,omitempty" xml:"total,omitempty" require:"true"`
  // {"en":"GetMaterialListMaterial list", "zh_CN":"素材列表"}
  MaterialList []*GetMaterialListMaterial `json:"materialList,omitempty" xml:"materialList,omitempty" require:"true" type:"Repeated"`
}

func (s GetMaterialListData) String() string {
  return tea.Prettify(s)
}

func (s GetMaterialListData) GoString() string {
  return s.String()
}

func (s *GetMaterialListData) SetTotal(v int32) *GetMaterialListData {
  s.Total = &v
  return s
}

func (s *GetMaterialListData) SetMaterialList(v []*GetMaterialListMaterial) *GetMaterialListData {
  s.MaterialList = v
  return s
}

type GetMaterialListMaterial struct {
  // {"en":"GetMaterialListMaterial ID", "zh_CN":"素材ID"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {"en":"GetMaterialListMaterial name", "zh_CN":"素材名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"File suffix", "zh_CN":"文件后缀"}
  Suffix *string `json:"suffix,omitempty" xml:"suffix,omitempty" require:"true"`
  // {"en":"File size (unit: bit)", "zh_CN":"文件大小(单位为bit)"}
  FileSize *int64 `json:"fileSize,omitempty" xml:"fileSize,omitempty" require:"true"`
  // {"en":"File url", "zh_CN":"文件url"}
  Url *string `json:"url,omitempty" xml:"url,omitempty" require:"true"`
  // {"en":"Create user", "zh_CN":"创建用户"}
  CreateUser *string `json:"createUser,omitempty" xml:"createUser,omitempty" require:"true"`
  // {"en":"Create time", "zh_CN":"创建时间"}
  CreateTime *int32 `json:"createTime,omitempty" xml:"createTime,omitempty" require:"true"`
}

func (s GetMaterialListMaterial) String() string {
  return tea.Prettify(s)
}

func (s GetMaterialListMaterial) GoString() string {
  return s.String()
}

func (s *GetMaterialListMaterial) SetId(v string) *GetMaterialListMaterial {
  s.Id = &v
  return s
}

func (s *GetMaterialListMaterial) SetName(v string) *GetMaterialListMaterial {
  s.Name = &v
  return s
}

func (s *GetMaterialListMaterial) SetSuffix(v string) *GetMaterialListMaterial {
  s.Suffix = &v
  return s
}

func (s *GetMaterialListMaterial) SetFileSize(v int64) *GetMaterialListMaterial {
  s.FileSize = &v
  return s
}

func (s *GetMaterialListMaterial) SetUrl(v string) *GetMaterialListMaterial {
  s.Url = &v
  return s
}

func (s *GetMaterialListMaterial) SetCreateUser(v string) *GetMaterialListMaterial {
  s.CreateUser = &v
  return s
}

func (s *GetMaterialListMaterial) SetCreateTime(v int32) *GetMaterialListMaterial {
  s.CreateTime = &v
  return s
}

type GetMaterialListPaths struct {
}

func (s GetMaterialListPaths) String() string {
  return tea.Prettify(s)
}

func (s GetMaterialListPaths) GoString() string {
  return s.String()
}

type GetMaterialListParameters struct {
}

func (s GetMaterialListParameters) String() string {
  return tea.Prettify(s)
}

func (s GetMaterialListParameters) GoString() string {
  return s.String()
}

type GetMaterialListRequestHeader struct {
}

func (s GetMaterialListRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetMaterialListRequestHeader) GoString() string {
  return s.String()
}

type GetMaterialListResponseHeader struct {
}

func (s GetMaterialListResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetMaterialListResponseHeader) GoString() string {
  return s.String()
}




type DeleteAudioRequest struct {
  // {"en":"id of the audio you want to delete", "zh_CN":"需要删除的音频的id"}
  AudioId *string `json:"audioId,omitempty" xml:"audioId,omitempty" require:"true"`
  // {"en":"Occupancy check; 0 not check, 1 check, default check", "zh_CN":"占用校验；0 不校验，1 校验, 默认校验"}
  ValidateOccupy *int32 `json:"validateOccupy,omitempty" xml:"validateOccupy,omitempty"`
}

func (s DeleteAudioRequest) String() string {
  return tea.Prettify(s)
}

func (s DeleteAudioRequest) GoString() string {
  return s.String()
}

func (s *DeleteAudioRequest) SetAudioId(v string) *DeleteAudioRequest {
  s.AudioId = &v
  return s
}

func (s *DeleteAudioRequest) SetValidateOccupy(v int32) *DeleteAudioRequest {
  s.ValidateOccupy = &v
  return s
}

type DeleteAudioResponse struct {
  // {"en":"Create the result status code. 200 indicates success", "zh_CN":"创建结果状态码，200为成功"}
  Code *int32 `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Successful operation", "zh_CN":"操作成功"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Return data", "zh_CN":"返回数据"}
  Data *string `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s DeleteAudioResponse) String() string {
  return tea.Prettify(s)
}

func (s DeleteAudioResponse) GoString() string {
  return s.String()
}

func (s *DeleteAudioResponse) SetCode(v int32) *DeleteAudioResponse {
  s.Code = &v
  return s
}

func (s *DeleteAudioResponse) SetMessage(v string) *DeleteAudioResponse {
  s.Message = &v
  return s
}

func (s *DeleteAudioResponse) SetData(v string) *DeleteAudioResponse {
  s.Data = &v
  return s
}

type DeleteAudioPaths struct {
}

func (s DeleteAudioPaths) String() string {
  return tea.Prettify(s)
}

func (s DeleteAudioPaths) GoString() string {
  return s.String()
}

type DeleteAudioParameters struct {
}

func (s DeleteAudioParameters) String() string {
  return tea.Prettify(s)
}

func (s DeleteAudioParameters) GoString() string {
  return s.String()
}

type DeleteAudioRequestHeader struct {
}

func (s DeleteAudioRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteAudioRequestHeader) GoString() string {
  return s.String()
}

type DeleteAudioResponseHeader struct {
}

func (s DeleteAudioResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteAudioResponseHeader) GoString() string {
  return s.String()
}




type EditAudioRequest struct {
  // {"en":"Audio id to edit", "zh_CN":"需编辑的音频id"}
  AudioId *string `json:"audioId,omitempty" xml:"audioId,omitempty" require:"true"`
  // {"en":"The value is a new audio name, with a maximum of 40 Chinese characters.", "zh_CN":"修改后的音频名称，最大40中文。"}
  AudioName *string `json:"audioName,omitempty" xml:"audioName,omitempty"`
  // {"en":"Adjusted published domain name.", "zh_CN":"调整后的发布域名。"}
  PublishDomain *string `json:"publishDomain,omitempty" xml:"publishDomain,omitempty"`
}

func (s EditAudioRequest) String() string {
  return tea.Prettify(s)
}

func (s EditAudioRequest) GoString() string {
  return s.String()
}

func (s *EditAudioRequest) SetAudioId(v string) *EditAudioRequest {
  s.AudioId = &v
  return s
}

func (s *EditAudioRequest) SetAudioName(v string) *EditAudioRequest {
  s.AudioName = &v
  return s
}

func (s *EditAudioRequest) SetPublishDomain(v string) *EditAudioRequest {
  s.PublishDomain = &v
  return s
}

type EditAudioResponse struct {
  // {"en":"Status code", "zh_CN":"状态码"}
  Code *int32 `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Operational information", "zh_CN":"操作信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Return data", "zh_CN":"返回数据"}
  Data *string `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s EditAudioResponse) String() string {
  return tea.Prettify(s)
}

func (s EditAudioResponse) GoString() string {
  return s.String()
}

func (s *EditAudioResponse) SetCode(v int32) *EditAudioResponse {
  s.Code = &v
  return s
}

func (s *EditAudioResponse) SetMessage(v string) *EditAudioResponse {
  s.Message = &v
  return s
}

func (s *EditAudioResponse) SetData(v string) *EditAudioResponse {
  s.Data = &v
  return s
}

type EditAudioPaths struct {
}

func (s EditAudioPaths) String() string {
  return tea.Prettify(s)
}

func (s EditAudioPaths) GoString() string {
  return s.String()
}

type EditAudioParameters struct {
}

func (s EditAudioParameters) String() string {
  return tea.Prettify(s)
}

func (s EditAudioParameters) GoString() string {
  return s.String()
}

type EditAudioRequestHeader struct {
}

func (s EditAudioRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s EditAudioRequestHeader) GoString() string {
  return s.String()
}

type EditAudioResponseHeader struct {
}

func (s EditAudioResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s EditAudioResponseHeader) GoString() string {
  return s.String()
}




type GetAudioListRequest struct {
  // {"en":"Create a user. The value is blank by default. Multiple entries are allowed. They are separated by commas (,) and cannot start or end with a comma. This parameter is restricted by permissions. Only sub-accounts or users with special permissions can be queried.", "zh_CN":"创建用户。默认为空，允许传多个，以半角逗号隔开，不能以逗号开头或结尾，两个逗号之间的内容不为能为空或空白字符。该参数受权限限制，只能查询子账户或权限特殊配置的用户。"}
  CreateUser *string `json:"createUser,omitempty" xml:"createUser,omitempty"`
  // {"en":"The start time is in the format of 2016-01-01 12:00:00. It is used to query the audio according to the creation time range.", "zh_CN":"查询起始时间，时间格式为，2016-01-01 12:00:00；用于按创建时间段查询音频；"}
  StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
  // {"en":"The query end time is in the format of 2016-01-01 12:00:00. This parameter is used to query audio files within the creation period, which is smaller than the current query time.", "zh_CN":"查询截止时间，时间格式为，2016-01-01 12:00:00；用于按创建时间段查询音频，小于当前查询时间；"}
  EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
  // {"en":"Audio name, used to filter audio, support fuzzy matching;", "zh_CN":"音频名称，用于筛选音频，支持模糊匹配；"}
  AudioName *string `json:"audioName,omitempty" xml:"audioName,omitempty"`
  // {"en":"Audio ID, used to filter audio;", "zh_CN":"音频ID，用于筛选音频；"}
  AudioId *string `json:"audioId,omitempty" xml:"audioId,omitempty"`
  // {"en":"The value range is · 0(in descending order by creation time)· 1(in ascending order by creation time) The default value is 0", "zh_CN":"列表排列顺序，取值范围 ： · 0(按创建时间降序排列)· 1(按创建时间升序排列)默认为0"}
  ListOrder *string `json:"listOrder,omitempty" xml:"listOrder,omitempty"`
  // {"en":"The page of the audio list starts with 1. The default value is 1. The product of pageIndex and pageSize must be less than 100000.", "zh_CN":"取音频列表第几页，从1开始取值,默认为1。入参pageIndex和pageSize的乘积必须不大于100000。"}
  PageIndex *string `json:"pageIndex,omitempty" xml:"pageIndex,omitempty"`
  // {"en":"Average number of audio files per page. The value ranges from 1 to 50. The default value is 10. The product of pageIndex and pageSize must be less than 100000", "zh_CN":"平均每页音频数量，取值范围1-50，默认为10。入参pageIndex和pageSize的乘积必须不大于100000。"}
  PageSize *string `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s GetAudioListRequest) String() string {
  return tea.Prettify(s)
}

func (s GetAudioListRequest) GoString() string {
  return s.String()
}

func (s *GetAudioListRequest) SetCreateUser(v string) *GetAudioListRequest {
  s.CreateUser = &v
  return s
}

func (s *GetAudioListRequest) SetStartTime(v string) *GetAudioListRequest {
  s.StartTime = &v
  return s
}

func (s *GetAudioListRequest) SetEndTime(v string) *GetAudioListRequest {
  s.EndTime = &v
  return s
}

func (s *GetAudioListRequest) SetAudioName(v string) *GetAudioListRequest {
  s.AudioName = &v
  return s
}

func (s *GetAudioListRequest) SetAudioId(v string) *GetAudioListRequest {
  s.AudioId = &v
  return s
}

func (s *GetAudioListRequest) SetListOrder(v string) *GetAudioListRequest {
  s.ListOrder = &v
  return s
}

func (s *GetAudioListRequest) SetPageIndex(v string) *GetAudioListRequest {
  s.PageIndex = &v
  return s
}

func (s *GetAudioListRequest) SetPageSize(v string) *GetAudioListRequest {
  s.PageSize = &v
  return s
}

type GetAudioListResponse struct {
  // {'en':"Return status code", "zh_CN":"返回状态码"}
  Code *int32 `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  Data *GetAudioListDataList `json:"data,omitempty" xml:"data,omitempty" require:"true"`
  // {'en':"Return message","zh_CN":"返回消息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s GetAudioListResponse) String() string {
  return tea.Prettify(s)
}

func (s GetAudioListResponse) GoString() string {
  return s.String()
}

func (s *GetAudioListResponse) SetCode(v int32) *GetAudioListResponse {
  s.Code = &v
  return s
}

func (s *GetAudioListResponse) SetData(v *GetAudioListDataList) *GetAudioListResponse {
  s.Data = v
  return s
}

func (s *GetAudioListResponse) SetMessage(v string) *GetAudioListResponse {
  s.Message = &v
  return s
}

type GetAudioListDataList struct {
  GetAudioListAudioGetListResponseList []*GetAudioListAudioGetListResponseList `json:"audioGetListResponseList,omitempty" xml:"audioGetListResponseList,omitempty" require:"true" type:"Repeated"`
  // {'en':'Audio upload time', 'zh_CN':'音频上传时间'}
  AudioTotal *int32 `json:"audioTotal,omitempty" xml:"audioTotal,omitempty" require:"true"`
}

func (s GetAudioListDataList) String() string {
  return tea.Prettify(s)
}

func (s GetAudioListDataList) GoString() string {
  return s.String()
}

func (s *GetAudioListDataList) SetAudioGetListResponseList(v []*GetAudioListAudioGetListResponseList) *GetAudioListDataList {
  s.GetAudioListAudioGetListResponseList = v
  return s
}

func (s *GetAudioListDataList) SetAudioTotal(v int32) *GetAudioListDataList {
  s.AudioTotal = &v
  return s
}

type GetAudioListAudioGetListResponseList struct {
  // {'en':'Audio name', 'zh_CN':'音频名称'}
  AudioName *string `json:"audioName,omitempty" xml:"audioName,omitempty" require:"true"`
  // {'en':'Audio ID', 'zh_CN':'音频ID'}
  AudioId *string `json:"audioId,omitempty" xml:"audioId,omitempty" require:"true"`
  // {'en':'Create user', 'zh_CN':'创建人'}
  CreateUser *string `json:"createUser,omitempty" xml:"createUser,omitempty" require:"true"`
  // {'en':'The space occupied by audio, and the total space used by video and video after transcoding', 'zh_CN':'音频占用空间大小，视频及其转码后视频的总空间使用量'}
  AudioSize *int32 `json:"audioSize,omitempty" xml:"audioSize,omitempty" require:"true"`
  // {'en':'Audio duration', 'zh_CN':'音频时长'}
  AudioDuration *int32 `json:"audioDuration,omitempty" xml:"audioDuration,omitempty" require:"true"`
  // {'en':'File url', 'zh_CN':'文件url'}
  Url *string `json:"url,omitempty" xml:"url,omitempty" require:"true"`
  // {'en':'File suffix', 'zh_CN':'文件后缀'}
  Suffix *string `json:"suffix,omitempty" xml:"suffix,omitempty" require:"true"`
  // {'en':'Create time', 'zh_CN':'音频创建时间'}
  CreateTime *int32 `json:"createTime,omitempty" xml:"createTime,omitempty" require:"true"`
  // {'en':'Upload time', 'zh_CN':'音频上传时间'}
  UploadTime *int32 `json:"uploadTime,omitempty" xml:"uploadTime,omitempty" require:"true"`
}

func (s GetAudioListAudioGetListResponseList) String() string {
  return tea.Prettify(s)
}

func (s GetAudioListAudioGetListResponseList) GoString() string {
  return s.String()
}

func (s *GetAudioListAudioGetListResponseList) SetAudioName(v string) *GetAudioListAudioGetListResponseList {
  s.AudioName = &v
  return s
}

func (s *GetAudioListAudioGetListResponseList) SetAudioId(v string) *GetAudioListAudioGetListResponseList {
  s.AudioId = &v
  return s
}

func (s *GetAudioListAudioGetListResponseList) SetCreateUser(v string) *GetAudioListAudioGetListResponseList {
  s.CreateUser = &v
  return s
}

func (s *GetAudioListAudioGetListResponseList) SetAudioSize(v int32) *GetAudioListAudioGetListResponseList {
  s.AudioSize = &v
  return s
}

func (s *GetAudioListAudioGetListResponseList) SetAudioDuration(v int32) *GetAudioListAudioGetListResponseList {
  s.AudioDuration = &v
  return s
}

func (s *GetAudioListAudioGetListResponseList) SetUrl(v string) *GetAudioListAudioGetListResponseList {
  s.Url = &v
  return s
}

func (s *GetAudioListAudioGetListResponseList) SetSuffix(v string) *GetAudioListAudioGetListResponseList {
  s.Suffix = &v
  return s
}

func (s *GetAudioListAudioGetListResponseList) SetCreateTime(v int32) *GetAudioListAudioGetListResponseList {
  s.CreateTime = &v
  return s
}

func (s *GetAudioListAudioGetListResponseList) SetUploadTime(v int32) *GetAudioListAudioGetListResponseList {
  s.UploadTime = &v
  return s
}

type GetAudioListPaths struct {
}

func (s GetAudioListPaths) String() string {
  return tea.Prettify(s)
}

func (s GetAudioListPaths) GoString() string {
  return s.String()
}

type GetAudioListParameters struct {
}

func (s GetAudioListParameters) String() string {
  return tea.Prettify(s)
}

func (s GetAudioListParameters) GoString() string {
  return s.String()
}

type GetAudioListRequestHeader struct {
}

func (s GetAudioListRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetAudioListRequestHeader) GoString() string {
  return s.String()
}

type GetAudioListResponseHeader struct {
}

func (s GetAudioListResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetAudioListResponseHeader) GoString() string {
  return s.String()
}




type MaterialEditRequest struct {
  // {"en":"id of the material to edit", "zh_CN":"需编辑的素材id"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {"en":"The value can contain a maximum of 40 characters.", "zh_CN":"修改后的素材名称，最大40字符。"}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // {"en":"Adjusted published domain name.", "zh_CN":"调整后的发布域名。"}
  PublishDomain *string `json:"publishDomain,omitempty" xml:"publishDomain,omitempty"`
}

func (s MaterialEditRequest) String() string {
  return tea.Prettify(s)
}

func (s MaterialEditRequest) GoString() string {
  return s.String()
}

func (s *MaterialEditRequest) SetId(v string) *MaterialEditRequest {
  s.Id = &v
  return s
}

func (s *MaterialEditRequest) SetName(v string) *MaterialEditRequest {
  s.Name = &v
  return s
}

func (s *MaterialEditRequest) SetPublishDomain(v string) *MaterialEditRequest {
  s.PublishDomain = &v
  return s
}

type MaterialEditResponse struct {
  // {"en":"Result status code, 200 indicates success", "zh_CN":"结果状态码，200为成功"}
  Code *int32 `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Return message", "zh_CN":"返回消息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Return data", "zh_CN":"返回数据"}
  Data *string `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s MaterialEditResponse) String() string {
  return tea.Prettify(s)
}

func (s MaterialEditResponse) GoString() string {
  return s.String()
}

func (s *MaterialEditResponse) SetCode(v int32) *MaterialEditResponse {
  s.Code = &v
  return s
}

func (s *MaterialEditResponse) SetMessage(v string) *MaterialEditResponse {
  s.Message = &v
  return s
}

func (s *MaterialEditResponse) SetData(v string) *MaterialEditResponse {
  s.Data = &v
  return s
}

type MaterialEditPaths struct {
}

func (s MaterialEditPaths) String() string {
  return tea.Prettify(s)
}

func (s MaterialEditPaths) GoString() string {
  return s.String()
}

type MaterialEditParameters struct {
}

func (s MaterialEditParameters) String() string {
  return tea.Prettify(s)
}

func (s MaterialEditParameters) GoString() string {
  return s.String()
}

type MaterialEditRequestHeader struct {
}

func (s MaterialEditRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s MaterialEditRequestHeader) GoString() string {
  return s.String()
}

type MaterialEditResponseHeader struct {
}

func (s MaterialEditResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s MaterialEditResponseHeader) GoString() string {
  return s.String()
}




type DeleteMaterialRequest struct {
  // {"en":"ID of the material that you want to delete", "zh_CN":"需要删除的素材ID"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {"en":"Occupied check, default check
  // 0 uncheck
  // 1 check", "zh_CN":"占用校验，默认校验
  // 0 不校验
  // 1 校验"}
  ValidateOccupy *string `json:"validateOccupy,omitempty" xml:"validateOccupy,omitempty"`
}

func (s DeleteMaterialRequest) String() string {
  return tea.Prettify(s)
}

func (s DeleteMaterialRequest) GoString() string {
  return s.String()
}

func (s *DeleteMaterialRequest) SetId(v string) *DeleteMaterialRequest {
  s.Id = &v
  return s
}

func (s *DeleteMaterialRequest) SetValidateOccupy(v string) *DeleteMaterialRequest {
  s.ValidateOccupy = &v
  return s
}

type DeleteMaterialResponse struct {
  // {"en":"Result status code, 200 indicates success", "zh_CN":"结果状态码，200为成功"}
  Code *int32 `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Return message", "zh_CN":"返回消息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Return data", "zh_CN":"返回数据"}
  Data *string `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s DeleteMaterialResponse) String() string {
  return tea.Prettify(s)
}

func (s DeleteMaterialResponse) GoString() string {
  return s.String()
}

func (s *DeleteMaterialResponse) SetCode(v int32) *DeleteMaterialResponse {
  s.Code = &v
  return s
}

func (s *DeleteMaterialResponse) SetMessage(v string) *DeleteMaterialResponse {
  s.Message = &v
  return s
}

func (s *DeleteMaterialResponse) SetData(v string) *DeleteMaterialResponse {
  s.Data = &v
  return s
}

type DeleteMaterialPaths struct {
}

func (s DeleteMaterialPaths) String() string {
  return tea.Prettify(s)
}

func (s DeleteMaterialPaths) GoString() string {
  return s.String()
}

type DeleteMaterialParameters struct {
}

func (s DeleteMaterialParameters) String() string {
  return tea.Prettify(s)
}

func (s DeleteMaterialParameters) GoString() string {
  return s.String()
}

type DeleteMaterialRequestHeader struct {
}

func (s DeleteMaterialRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteMaterialRequestHeader) GoString() string {
  return s.String()
}

type DeleteMaterialResponseHeader struct {
}

func (s DeleteMaterialResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteMaterialResponseHeader) GoString() string {
  return s.String()
}




type DeleteVideoRequest struct {
  // {"en":"id of the video you want to delete", "zh_CN":"需要删除的视频的id"}
  VideoId *string `json:"videoId,omitempty" xml:"videoId,omitempty" require:"true"`
  // {"en":"Occupied check, default check
  // 0 uncheck
  // 1 check", "zh_CN":"占用校验，默认校验
  // 0 不校验
  // 1 校验"}
  ValidateOccupy *int32 `json:"validateOccupy,omitempty" xml:"validateOccupy,omitempty"`
  // {"en":"Delete the result callback notification address", "zh_CN":"删除结果回调通知地址"}
  NotifyUrl *string `json:"notifyUrl,omitempty" xml:"notifyUrl,omitempty"`
  // {"en":"Whether to delete asynchronously, the default value is 1;
  // 0: delete files synchronously, 1: delete files asynchronously", "zh_CN":"是否异步删除，默认为1；
  // 0:同步删除文件，1：异步删除文件"}
  IsAsync *int32 `json:"isAsync,omitempty" xml:"isAsync,omitempty"`
  // {"en":"The default value is 0; 0: delete all files (including original files and Transcoding files), 1: delete only the original file", "zh_CN":"只删除原文件
  // 默认为0；
  // 0:删除所有文件（包含原文件和转码文件），1：只删除原文件"}
  OnlyDeleteSource *int32 `json:"onlyDeleteSource,omitempty" xml:"onlyDeleteSource,omitempty"`
}

func (s DeleteVideoRequest) String() string {
  return tea.Prettify(s)
}

func (s DeleteVideoRequest) GoString() string {
  return s.String()
}

func (s *DeleteVideoRequest) SetVideoId(v string) *DeleteVideoRequest {
  s.VideoId = &v
  return s
}

func (s *DeleteVideoRequest) SetValidateOccupy(v int32) *DeleteVideoRequest {
  s.ValidateOccupy = &v
  return s
}

func (s *DeleteVideoRequest) SetNotifyUrl(v string) *DeleteVideoRequest {
  s.NotifyUrl = &v
  return s
}

func (s *DeleteVideoRequest) SetIsAsync(v int32) *DeleteVideoRequest {
  s.IsAsync = &v
  return s
}

func (s *DeleteVideoRequest) SetOnlyDeleteSource(v int32) *DeleteVideoRequest {
  s.OnlyDeleteSource = &v
  return s
}

type DeleteVideoResponse struct {
  // {"en":"Result status code, 200 indicates success", "zh_CN":"结果状态码，200为成功"}
  Code *int32 `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Return message", "zh_CN":"返回消息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Return data", "zh_CN":"返回数据"}
  Data *string `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s DeleteVideoResponse) String() string {
  return tea.Prettify(s)
}

func (s DeleteVideoResponse) GoString() string {
  return s.String()
}

func (s *DeleteVideoResponse) SetCode(v int32) *DeleteVideoResponse {
  s.Code = &v
  return s
}

func (s *DeleteVideoResponse) SetMessage(v string) *DeleteVideoResponse {
  s.Message = &v
  return s
}

func (s *DeleteVideoResponse) SetData(v string) *DeleteVideoResponse {
  s.Data = &v
  return s
}

type DeleteVideoPaths struct {
}

func (s DeleteVideoPaths) String() string {
  return tea.Prettify(s)
}

func (s DeleteVideoPaths) GoString() string {
  return s.String()
}

type DeleteVideoParameters struct {
}

func (s DeleteVideoParameters) String() string {
  return tea.Prettify(s)
}

func (s DeleteVideoParameters) GoString() string {
  return s.String()
}

type DeleteVideoRequestHeader struct {
}

func (s DeleteVideoRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteVideoRequestHeader) GoString() string {
  return s.String()
}

type DeleteVideoResponseHeader struct {
}

func (s DeleteVideoResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteVideoResponseHeader) GoString() string {
  return s.String()
}




type VideoEditRequest struct {
  // {"en":"id of the video to edit", "zh_CN":"需编辑的视频id"}
  VideoId *string `json:"videoId,omitempty" xml:"videoId,omitempty" require:"true"`
  // {"en":"The modified video name contains a maximum of 40 Chinese characters", "zh_CN":"修改后的视频名称，最大40中文"}
  VideoName *string `json:"videoName,omitempty" xml:"videoName,omitempty"`
  // {"en":"Video description, a maximum of 200 Chinese characters", "zh_CN":"视频的描述信息，最大200中文"}
  VideoDescription *string `json:"videoDescription,omitempty" xml:"videoDescription,omitempty"`
  // {"en":"Modified video subcategories. This field is valid only when categoryNames is not entered. If there are multiple subcategories that are the same, the video sets all of them. videoClassification can only contain Chinese characters, uppercase and lowercase letters, numbers, underscores, and spaces. It cannot start with a space. Multiple characters can be separated by commas. (categoryNames is recommended). ", "zh_CN":"修改后的视频子分类。categoryNames不填时，此字段才有效。 如果有多个相同的子分类，则视频会设置所有的子分类。videoClassification只能包含中文、大小写字母、数字、下划线、空格，不能空格开头，多个支持逗号分隔。 （建议使用categoryNames）"}
  VideoClassification *string `json:"videoClassification,omitempty" xml:"videoClassification,omitempty"`
  // {"en":"After the video category is modified, you can specify parent category and child category. The format is an urlsafe base64 encoded JSON string array. Transfer an empty array, indicating that the video classification Settings are cleared.
  // The parameters are as follows:
  // 1) parentName: indicates the name of the parent class.
  // 2) childName: indicates the name of the subcategory
  // }", "zh_CN":"修改后的视频分类，可指定父分类和子分类。 格式为经过urlsafe base64编码的JSON字符串数组。传输空数组，则表示清空视频的分类设置。 
  // 参数如下：
  // 1）parentName: 父分类名称；
  // 2）childName: 子分类名称
  // 例：W3siY2hpbGROYW1lIjoi5a2Q5YiG57G7MSIsInBhcmVudE5hbWUiOiLniLbliIbnsbsxIn0seyJjaGlsZE5hbWUiOiLlrZDliIbnsbsyIiwicGFyZW50TmFtZSI6IueItuWIhuexuzIifV0= 编码前的json串格式如下：
  // [{\"childName\":\"子分类1\",\"parentName\":\"父分类1\"},{\"childName\":\"子分类2\",\"parentName\":\"父分类2\"}]"}
  CategoryNames *string `json:"categoryNames,omitempty" xml:"categoryNames,omitempty"`
  // {"en":"Adjusted published domain name", "zh_CN":"调整后的发布域名"}
  PublishDomain *string `json:"publishDomain,omitempty" xml:"publishDomain,omitempty"`
  // {"en":"The adjusted player ID", "zh_CN":"调整后的播放器ID"}
  PlayerId *string `json:"playerId,omitempty" xml:"playerId,omitempty"`
}

func (s VideoEditRequest) String() string {
  return tea.Prettify(s)
}

func (s VideoEditRequest) GoString() string {
  return s.String()
}

func (s *VideoEditRequest) SetVideoId(v string) *VideoEditRequest {
  s.VideoId = &v
  return s
}

func (s *VideoEditRequest) SetVideoName(v string) *VideoEditRequest {
  s.VideoName = &v
  return s
}

func (s *VideoEditRequest) SetVideoDescription(v string) *VideoEditRequest {
  s.VideoDescription = &v
  return s
}

func (s *VideoEditRequest) SetVideoClassification(v string) *VideoEditRequest {
  s.VideoClassification = &v
  return s
}

func (s *VideoEditRequest) SetCategoryNames(v string) *VideoEditRequest {
  s.CategoryNames = &v
  return s
}

func (s *VideoEditRequest) SetPublishDomain(v string) *VideoEditRequest {
  s.PublishDomain = &v
  return s
}

func (s *VideoEditRequest) SetPlayerId(v string) *VideoEditRequest {
  s.PlayerId = &v
  return s
}

type VideoEditResponse struct {
  // {"en":"Result status code, 200 indicates success", "zh_CN":"结果状态码，200为成功"}
  Code *int32 `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Return message", "zh_CN":"返回消息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s VideoEditResponse) String() string {
  return tea.Prettify(s)
}

func (s VideoEditResponse) GoString() string {
  return s.String()
}

func (s *VideoEditResponse) SetCode(v int32) *VideoEditResponse {
  s.Code = &v
  return s
}

func (s *VideoEditResponse) SetMessage(v string) *VideoEditResponse {
  s.Message = &v
  return s
}

type VideoEditPaths struct {
}

func (s VideoEditPaths) String() string {
  return tea.Prettify(s)
}

func (s VideoEditPaths) GoString() string {
  return s.String()
}

type VideoEditParameters struct {
}

func (s VideoEditParameters) String() string {
  return tea.Prettify(s)
}

func (s VideoEditParameters) GoString() string {
  return s.String()
}

type VideoEditRequestHeader struct {
}

func (s VideoEditRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s VideoEditRequestHeader) GoString() string {
  return s.String()
}

type VideoEditResponseHeader struct {
}

func (s VideoEditResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s VideoEditResponseHeader) GoString() string {
  return s.String()
}




type VideoBlockRequest struct {
  // {"en":"id of the video to be masked", "zh_CN":"要屏蔽的视频id"}
  VideoId *string `json:"videoId,omitempty" xml:"videoId,omitempty" require:"true"`
}

func (s VideoBlockRequest) String() string {
  return tea.Prettify(s)
}

func (s VideoBlockRequest) GoString() string {
  return s.String()
}

func (s *VideoBlockRequest) SetVideoId(v string) *VideoBlockRequest {
  s.VideoId = &v
  return s
}

type VideoBlockResponse struct {
  // {"en":"Result status code, 200 indicates success", "zh_CN":"结果状态码，200为成功"}
  Code *int32 `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Return message", "zh_CN":"返回消息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Return data", "zh_CN":"返回数据"}
  Data *string `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s VideoBlockResponse) String() string {
  return tea.Prettify(s)
}

func (s VideoBlockResponse) GoString() string {
  return s.String()
}

func (s *VideoBlockResponse) SetCode(v int32) *VideoBlockResponse {
  s.Code = &v
  return s
}

func (s *VideoBlockResponse) SetMessage(v string) *VideoBlockResponse {
  s.Message = &v
  return s
}

func (s *VideoBlockResponse) SetData(v string) *VideoBlockResponse {
  s.Data = &v
  return s
}

type VideoBlockPaths struct {
}

func (s VideoBlockPaths) String() string {
  return tea.Prettify(s)
}

func (s VideoBlockPaths) GoString() string {
  return s.String()
}

type VideoBlockParameters struct {
}

func (s VideoBlockParameters) String() string {
  return tea.Prettify(s)
}

func (s VideoBlockParameters) GoString() string {
  return s.String()
}

type VideoBlockRequestHeader struct {
}

func (s VideoBlockRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s VideoBlockRequestHeader) GoString() string {
  return s.String()
}

type VideoBlockResponseHeader struct {
}

func (s VideoBlockResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s VideoBlockResponseHeader) GoString() string {
  return s.String()
}




type CreateCategoryRequest struct {
  // {"en":"Id of the parent node. If it is empty, create the parent node", "zh_CN":"父节点Id。如果为空则创建父级节点"}
  ParentNodeId *int32 `json:"parentNodeId,omitempty" xml:"parentNodeId,omitempty"`
  // {"en":"Label name (32 characters supported)", "zh_CN":"标签名称（支持32个字符）"}
  LabelName *string `json:"labelName,omitempty" xml:"labelName,omitempty" require:"true"`
}

func (s CreateCategoryRequest) String() string {
  return tea.Prettify(s)
}

func (s CreateCategoryRequest) GoString() string {
  return s.String()
}

func (s *CreateCategoryRequest) SetParentNodeId(v int32) *CreateCategoryRequest {
  s.ParentNodeId = &v
  return s
}

func (s *CreateCategoryRequest) SetLabelName(v string) *CreateCategoryRequest {
  s.LabelName = &v
  return s
}

type CreateCategoryResponse struct {
  // {"en":"Result status code, 200 indicates success", "zh_CN":"结果状态码，200为成功"}
  Code *int32 `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Return message", "zh_CN":"返回消息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Return data", "zh_CN":"返回数据"}
  CreateCategoryData []*CreateCategoryData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s CreateCategoryResponse) String() string {
  return tea.Prettify(s)
}

func (s CreateCategoryResponse) GoString() string {
  return s.String()
}

func (s *CreateCategoryResponse) SetCode(v int32) *CreateCategoryResponse {
  s.Code = &v
  return s
}

func (s *CreateCategoryResponse) SetMessage(v string) *CreateCategoryResponse {
  s.Message = &v
  return s
}

func (s *CreateCategoryResponse) SetData(v []*CreateCategoryData) *CreateCategoryResponse {
  s.CreateCategoryData = v
  return s
}

type CreateCategoryData struct {
  // {"en":"Tag node ID", "zh_CN":"标签节点ID"}
  NodeId *int64 `json:"nodeId,omitempty" xml:"nodeId,omitempty" require:"true"`
}

func (s CreateCategoryData) String() string {
  return tea.Prettify(s)
}

func (s CreateCategoryData) GoString() string {
  return s.String()
}

func (s *CreateCategoryData) SetNodeId(v int64) *CreateCategoryData {
  s.NodeId = &v
  return s
}

type CreateCategoryPaths struct {
}

func (s CreateCategoryPaths) String() string {
  return tea.Prettify(s)
}

func (s CreateCategoryPaths) GoString() string {
  return s.String()
}

type CreateCategoryParameters struct {
}

func (s CreateCategoryParameters) String() string {
  return tea.Prettify(s)
}

func (s CreateCategoryParameters) GoString() string {
  return s.String()
}

type CreateCategoryRequestHeader struct {
}

func (s CreateCategoryRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s CreateCategoryRequestHeader) GoString() string {
  return s.String()
}

type CreateCategoryResponseHeader struct {
}

func (s CreateCategoryResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s CreateCategoryResponseHeader) GoString() string {
  return s.String()
}




type GetVideoListRequest struct {
  // {"en":"Create a user. The value is blank by default. Multiple entries are allowed. They are separated by commas (,) and cannot start or end with a comma. This parameter is restricted by permissions. Only sub-accounts or users with special permissions can be queried.", "zh_CN":"创建用户。默认为空，允许传多个，以半角逗号隔开，不能以逗号开头或结尾，两个逗号之间的内容不为能为空或空白字符。该参数受权限限制，只能查询子账户或权限特殊配置的用户。"}
  CreateUser *string `json:"createUser,omitempty" xml:"createUser,omitempty"`
  // {"en":"The start time is in the format of 2016-01-01 12:00:00. Query videos by creation time range.", "zh_CN":"查询起始时间，时间格式为，2016-01-01 12:00:00；用于按创建时间段查询视频；"}
  StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
  // {"en":"The query end time is in the format of 2016-01-01 12:00:00. This parameter is used to query videos in the creation period, which is smaller than the current query time.", "zh_CN":"查询截止时间，时间格式为，2016-01-01 12:00:00；用于按创建时间段查询视频，小于当前查询时间；"}
  EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
  // {"en":"Video name, used to filter videos, support fuzzy matching;", "zh_CN":"视频名称，用于筛选视频，支持模糊匹配；"}
  VideoName *string `json:"videoName,omitempty" xml:"videoName,omitempty"`
  // {"en":"Video ID, used to filter videos;", "zh_CN":"视频ID，用于筛选视频；"}
  VideoId *string `json:"videoId,omitempty" xml:"videoId,omitempty"`
  // {"en":"Video status, used to filter videos", "zh_CN":"视频状态，用于筛选视频。取值范围 ：
  // 0(启用)
  // 1(屏蔽)"}
  VideoStatus *int32 `json:"videoStatus,omitempty" xml:"videoStatus,omitempty"`
  // {"en":"Authorized play is not enabled. Procedure", "zh_CN":"未开启授权播放，视频的转码状态的取值范围 ：
  // 1(已转码)
  // 2(未转码)
  // 3(转码中)
  // 4(转码失败)
  // 
  // 开启授权播放（视频加密）功能时的转码状态的取值范围 ：
  // 1(已加密转码)
  // 2(非加密转码)
  // 3(转码中)
  // 4(转码失败)
  // 5(未转码)"}
  TranscodeState *int32 `json:"transcodeState,omitempty" xml:"transcodeState,omitempty"`
  // {"en":"Subcategory names for video classification, supporting multiple subcategory names simultaneously, separated by',' for each subcategory name. videoClassification can only contain Chinese characters, uppercase and lowercase letters, numbers, underscores, and spaces. It cannot start with a space. Multiple characters can be separated by commas.", "zh_CN":"视频分类子分类名称，支持同时输入多个子分类名称，每个子分类名称用“,”隔开。只能包含中文、大小写字母、数字、下划线、空格，不能空格开头，多个支持逗号分隔"}
  VideoClassification *string `json:"videoClassification,omitempty" xml:"videoClassification,omitempty"`
  // {"en":"List order", "zh_CN":"列表排列顺序，取值范围 ：
  // 0(按创建时间降序排列)
  // 1(按创建时间升序排列)
  // 默认为0"}
  ListOrder *int32 `json:"listOrder,omitempty" xml:"listOrder,omitempty"`
  // {"en":"Page of the video list. The value starts from 1. The default value is 1. The product of pageIndex and pageSize must be less than 100000.", "zh_CN":"取视频列表第几页，从1开始取值,默认为1。入参pageIndex和pageSize的乘积必须不大于100000。"}
  PageIndex *int32 `json:"pageIndex,omitempty" xml:"pageIndex,omitempty"`
  // {"en":"Average number of videos per page. The value ranges from 1 to 50. The default value is 10. The product of pageIndex and pageSize must be less than 100000.", "zh_CN":"平均每页视频数量，取值范围1-50，默认为10。入参pageIndex和pageSize的乘积必须不大于100000。"}
  PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s GetVideoListRequest) String() string {
  return tea.Prettify(s)
}

func (s GetVideoListRequest) GoString() string {
  return s.String()
}

func (s *GetVideoListRequest) SetCreateUser(v string) *GetVideoListRequest {
  s.CreateUser = &v
  return s
}

func (s *GetVideoListRequest) SetStartTime(v string) *GetVideoListRequest {
  s.StartTime = &v
  return s
}

func (s *GetVideoListRequest) SetEndTime(v string) *GetVideoListRequest {
  s.EndTime = &v
  return s
}

func (s *GetVideoListRequest) SetVideoName(v string) *GetVideoListRequest {
  s.VideoName = &v
  return s
}

func (s *GetVideoListRequest) SetVideoId(v string) *GetVideoListRequest {
  s.VideoId = &v
  return s
}

func (s *GetVideoListRequest) SetVideoStatus(v int32) *GetVideoListRequest {
  s.VideoStatus = &v
  return s
}

func (s *GetVideoListRequest) SetTranscodeState(v int32) *GetVideoListRequest {
  s.TranscodeState = &v
  return s
}

func (s *GetVideoListRequest) SetVideoClassification(v string) *GetVideoListRequest {
  s.VideoClassification = &v
  return s
}

func (s *GetVideoListRequest) SetListOrder(v int32) *GetVideoListRequest {
  s.ListOrder = &v
  return s
}

func (s *GetVideoListRequest) SetPageIndex(v int32) *GetVideoListRequest {
  s.PageIndex = &v
  return s
}

func (s *GetVideoListRequest) SetPageSize(v int32) *GetVideoListRequest {
  s.PageSize = &v
  return s
}

type GetVideoListResponse struct {
  // {"en":"200 success", "zh_CN":"结果状态码，200为成功"}
  Code *int32 `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Successful operation", "zh_CN":"返回消息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Return data", "zh_CN":"返回数据"}
  GetVideoListData *GetVideoListData `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s GetVideoListResponse) String() string {
  return tea.Prettify(s)
}

func (s GetVideoListResponse) GoString() string {
  return s.String()
}

func (s *GetVideoListResponse) SetCode(v int32) *GetVideoListResponse {
  s.Code = &v
  return s
}

func (s *GetVideoListResponse) SetMessage(v string) *GetVideoListResponse {
  s.Message = &v
  return s
}

func (s *GetVideoListResponse) SetData(v *GetVideoListData) *GetVideoListResponse {
  s.GetVideoListData = v
  return s
}

type GetVideoListData struct {
  // {"en":"The number of records of the video list information currently returned. Note that the number of records returned here is only the number of records of the current page.", "zh_CN":"当前返回的视频列表信息的记录数，注意这里返回的记录数只是当前页的记录数。"}
  VideoTotal *int32 `json:"videoTotal,omitempty" xml:"videoTotal,omitempty" require:"true"`
  // {"en":"videoListInfo", "zh_CN":"视频列表信息"}
  VideoListInfo []*GetVideoListVideoInfo `json:"videoListInfo,omitempty" xml:"videoListInfo,omitempty" require:"true" type:"Repeated"`
}

func (s GetVideoListData) String() string {
  return tea.Prettify(s)
}

func (s GetVideoListData) GoString() string {
  return s.String()
}

func (s *GetVideoListData) SetVideoTotal(v int32) *GetVideoListData {
  s.VideoTotal = &v
  return s
}

func (s *GetVideoListData) SetVideoListInfo(v []*GetVideoListVideoInfo) *GetVideoListData {
  s.VideoListInfo = v
  return s
}

type GetVideoListVideoInfo struct {
  // {"en":"videoName", "zh_CN":"视频名称"}
  VideoName *string `json:"videoName,omitempty" xml:"videoName,omitempty" require:"true"`
  // {"en":"videoId", "zh_CN":"视频ID"}
  VideoId *string `json:"videoId,omitempty" xml:"videoId,omitempty" require:"true"`
  // {"en":"createUser", "zh_CN":"创建人"}
  CreateUser *string `json:"createUser,omitempty" xml:"createUser,omitempty" require:"true"`
  // {"en":"Whether to encrypt transcoding files. Value range:
  // 0(unencrypted)
  // 1(encrypted)", "zh_CN":"是否加密转码文件。取值范围 ：
  // 0(不加密)
  // 1(加密)"}
  Encrypt *int32 `json:"encrypt,omitempty" xml:"encrypt,omitempty" require:"true"`
  // {"en":"The space occupied by the video, and the total space used by the video and its transcoding", "zh_CN":"视频占用空间大小，视频及其转码后视频的总空间使用量"}
  VideoSize *string `json:"videoSize,omitempty" xml:"videoSize,omitempty" require:"true"`
  // {"en":"videoDuration", "zh_CN":"视频时长"}
  VideoDuration *string `json:"videoDuration,omitempty" xml:"videoDuration,omitempty" require:"true"`
  // {"en":"createTime", "zh_CN":"视频创建时间"}
  CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty" require:"true"`
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
  // {"en":"Video state
  // Value range:
  // 0(enable)
  // 1(mask)", "zh_CN":"视频状态
  // 取值范围：
  // 0(启用)
  // 1(屏蔽)"}
  VideoState *int32 `json:"videoState,omitempty" xml:"videoState,omitempty" require:"true"`
  // {"en":"If authorized play is not enabled, the video transcoding status ranges from:
  // 1(transcoded)
  // 2(no transcoding)
  // 3(in transcoding)
  // 4(Transcoding fails)
  // 
  // Value range of transcoding status when the Authorized Play (video encryption) function is enabled:
  // 1(encrypted transcoding)
  // 2(non-encrypted transcoding)
  // 3(in transcoding)
  // 4(Transcoding fails)
  // 5(not transcoded)", "zh_CN":"未开启授权播放，视频的转码状态的取值范围 ：
  // 1(已转码)
  // 2(未转码)
  // 3(转码中)
  // 4(转码失败)
  // 
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
  VideoResolutions []*GetVideoListVideoResolution `json:"videoResolutions,omitempty" xml:"videoResolutions,omitempty" require:"true" type:"Repeated"`
}

func (s GetVideoListVideoInfo) String() string {
  return tea.Prettify(s)
}

func (s GetVideoListVideoInfo) GoString() string {
  return s.String()
}

func (s *GetVideoListVideoInfo) SetVideoName(v string) *GetVideoListVideoInfo {
  s.VideoName = &v
  return s
}

func (s *GetVideoListVideoInfo) SetVideoId(v string) *GetVideoListVideoInfo {
  s.VideoId = &v
  return s
}

func (s *GetVideoListVideoInfo) SetCreateUser(v string) *GetVideoListVideoInfo {
  s.CreateUser = &v
  return s
}

func (s *GetVideoListVideoInfo) SetEncrypt(v int32) *GetVideoListVideoInfo {
  s.Encrypt = &v
  return s
}

func (s *GetVideoListVideoInfo) SetVideoSize(v string) *GetVideoListVideoInfo {
  s.VideoSize = &v
  return s
}

func (s *GetVideoListVideoInfo) SetVideoDuration(v string) *GetVideoListVideoInfo {
  s.VideoDuration = &v
  return s
}

func (s *GetVideoListVideoInfo) SetCreateTime(v string) *GetVideoListVideoInfo {
  s.CreateTime = &v
  return s
}

func (s *GetVideoListVideoInfo) SetUploadTime(v string) *GetVideoListVideoInfo {
  s.UploadTime = &v
  return s
}

func (s *GetVideoListVideoInfo) SetUpdateTime(v string) *GetVideoListVideoInfo {
  s.UpdateTime = &v
  return s
}

func (s *GetVideoListVideoInfo) SetVideoDescription(v string) *GetVideoListVideoInfo {
  s.VideoDescription = &v
  return s
}

func (s *GetVideoListVideoInfo) SetVideoClassification(v string) *GetVideoListVideoInfo {
  s.VideoClassification = &v
  return s
}

func (s *GetVideoListVideoInfo) SetImageUrl(v string) *GetVideoListVideoInfo {
  s.ImageUrl = &v
  return s
}

func (s *GetVideoListVideoInfo) SetPublishDomain(v string) *GetVideoListVideoInfo {
  s.PublishDomain = &v
  return s
}

func (s *GetVideoListVideoInfo) SetPlayerName(v string) *GetVideoListVideoInfo {
  s.PlayerName = &v
  return s
}

func (s *GetVideoListVideoInfo) SetPlayerId(v string) *GetVideoListVideoInfo {
  s.PlayerId = &v
  return s
}

func (s *GetVideoListVideoInfo) SetVideoState(v int32) *GetVideoListVideoInfo {
  s.VideoState = &v
  return s
}

func (s *GetVideoListVideoInfo) SetTranscodeState(v int32) *GetVideoListVideoInfo {
  s.TranscodeState = &v
  return s
}

func (s *GetVideoListVideoInfo) SetVideoSourceCode(v int32) *GetVideoListVideoInfo {
  s.VideoSourceCode = &v
  return s
}

func (s *GetVideoListVideoInfo) SetVideoResolutions(v []*GetVideoListVideoResolution) *GetVideoListVideoInfo {
  s.VideoResolutions = v
  return s
}

type GetVideoListVideoResolution struct {
  // {"en":"Clarity. Value range:
  // 1(original painting)
  // 2(fluency)
  // 3(standard definition)
  // 4(HD)
  // 5(Super clear)
  // -99(record file)", "zh_CN":"	清晰度。取值范围 ：
  // 1(原画)
  // 2(流畅)
  // 3(标清)
  // 4(高清)
  // 5(超清)
  // -99(录制文件)"}
  Clarity *int32 `json:"clarity,omitempty" xml:"clarity,omitempty" require:"true"`
  // {"en":"Terminal type. Value range:
  // -1(Source video)
  // 0(PC)
  // 1(mobile terminal)", "zh_CN":"终端类型。取值范围 ：
  // -1(源视频)
  // 0(PC端)
  // 1(移动端)"}
  ServerType *int32 `json:"serverType,omitempty" xml:"serverType,omitempty" require:"true"`
  // {"en":"height", "zh_CN":"高度"}
  Height *int32 `json:"height,omitempty" xml:"height,omitempty" require:"true"`
  // {"en":"width", "zh_CN":"宽度"}
  Width *int32 `json:"width,omitempty" xml:"width,omitempty" require:"true"`
  // {"en":"fileSize", "zh_CN":"文件大小(单位为bit)"}
  FileSize *int64 `json:"fileSize,omitempty" xml:"fileSize,omitempty" require:"true"`
}

func (s GetVideoListVideoResolution) String() string {
  return tea.Prettify(s)
}

func (s GetVideoListVideoResolution) GoString() string {
  return s.String()
}

func (s *GetVideoListVideoResolution) SetClarity(v int32) *GetVideoListVideoResolution {
  s.Clarity = &v
  return s
}

func (s *GetVideoListVideoResolution) SetServerType(v int32) *GetVideoListVideoResolution {
  s.ServerType = &v
  return s
}

func (s *GetVideoListVideoResolution) SetHeight(v int32) *GetVideoListVideoResolution {
  s.Height = &v
  return s
}

func (s *GetVideoListVideoResolution) SetWidth(v int32) *GetVideoListVideoResolution {
  s.Width = &v
  return s
}

func (s *GetVideoListVideoResolution) SetFileSize(v int64) *GetVideoListVideoResolution {
  s.FileSize = &v
  return s
}

type GetVideoListPaths struct {
}

func (s GetVideoListPaths) String() string {
  return tea.Prettify(s)
}

func (s GetVideoListPaths) GoString() string {
  return s.String()
}

type GetVideoListParameters struct {
}

func (s GetVideoListParameters) String() string {
  return tea.Prettify(s)
}

func (s GetVideoListParameters) GoString() string {
  return s.String()
}

type GetVideoListRequestHeader struct {
}

func (s GetVideoListRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetVideoListRequestHeader) GoString() string {
  return s.String()
}

type GetVideoListResponseHeader struct {
}

func (s GetVideoListResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetVideoListResponseHeader) GoString() string {
  return s.String()
}




