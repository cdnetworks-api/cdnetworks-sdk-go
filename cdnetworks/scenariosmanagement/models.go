package scenariosmanagement

// This file is auto-generated, don't edit it. Thanks.
import (
  "github.com/alibabacloud-go/tea/tea"
)

type GetSceneListRequest struct {
  // {"en":"Create a user. The value is blank by default. Multiple entries are allowed. They are separated by commas (,) and cannot start or end with a comma. This parameter is restricted by permissions. Only sub-accounts or users with special permissions can be queried.", "zh_CN":"创建用户。默认为空，允许传多个，以半角逗号隔开，不能以逗号开头或结尾，两个逗号之间的内容不为能为空或空白字符。该参数受权限限制，只能查询子账户或权限特殊配置的用户。"}
  CreateUser *string `json:"createUser,omitempty" xml:"createUser,omitempty"`
  // {"en":"The start time is in the format of 2016-01-01 12:00:00. Query by creation time range.", "zh_CN":"查询起始时间，时间格式为，2016-01-01 12:00:00；用于按创建时间段查询；"}
  StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
  // {"en":"The query end time is in the format of 2016-01-01 12:00:00. This parameter is used for query by creation time range, which is smaller than the current query time.", "zh_CN":"查询截止时间，时间格式为，2016-01-01 12:00:00；用于按创建时间段查询，小于当前查询时间；"}
  EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
  // {"en":"List order, value range: <br>
  // 0(in descending order by creation time)<br>
  // 1(in ascending order of creation time) The default value is 0", "zh_CN":"列表排列顺序，取值范围 ：<br>
  // 0(按创建时间降序排列)<br>
  // 1(按创建时间升序排列)默认为0"}
  ListOrder *int32 `json:"listOrder,omitempty" xml:"listOrder,omitempty"`
  // {"en":"On the page of the scenario list, the value starts from 1. The default value is 1. The product of pageIndex and pageSize must be less than 100000.", "zh_CN":"取场景列表第几页，从1开始取值,默认为1。入参pageIndex和pageSize的乘积必须不大于100000。"}
  PageIndex *int32 `json:"pageIndex,omitempty" xml:"pageIndex,omitempty"`
  // {"en":"Average Number of scenarios per page. The value ranges from 1 to 50. The default value is 10. The product of pageIndex and pageSize must be less than 100000.", "zh_CN":"平均每页场景数量，取值范围1-50，默认为10。入参pageIndex和pageSize的乘积必须不大于100000。"}
  PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s GetSceneListRequest) String() string {
  return tea.Prettify(s)
}

func (s GetSceneListRequest) GoString() string {
  return s.String()
}

func (s *GetSceneListRequest) SetCreateUser(v string) *GetSceneListRequest {
  s.CreateUser = &v
  return s
}

func (s *GetSceneListRequest) SetStartTime(v string) *GetSceneListRequest {
  s.StartTime = &v
  return s
}

func (s *GetSceneListRequest) SetEndTime(v string) *GetSceneListRequest {
  s.EndTime = &v
  return s
}

func (s *GetSceneListRequest) SetListOrder(v int32) *GetSceneListRequest {
  s.ListOrder = &v
  return s
}

func (s *GetSceneListRequest) SetPageIndex(v int32) *GetSceneListRequest {
  s.PageIndex = &v
  return s
}

func (s *GetSceneListRequest) SetPageSize(v int32) *GetSceneListRequest {
  s.PageSize = &v
  return s
}

type GetSceneListResponse struct {
  // {"en":"Status code", "zh_CN":"状态码"}
  Code *int32 `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Operational information", "zh_CN":"操作信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Return data", "zh_CN":"返回数据"}
  Data []*GetSceneListScene `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s GetSceneListResponse) String() string {
  return tea.Prettify(s)
}

func (s GetSceneListResponse) GoString() string {
  return s.String()
}

func (s *GetSceneListResponse) SetCode(v int32) *GetSceneListResponse {
  s.Code = &v
  return s
}

func (s *GetSceneListResponse) SetMessage(v string) *GetSceneListResponse {
  s.Message = &v
  return s
}

func (s *GetSceneListResponse) SetData(v []*GetSceneListScene) *GetSceneListResponse {
  s.Data = v
  return s
}

type GetSceneListScene struct {
  // {"en":"GetSceneListScene Id", "zh_CN":"场景Id"}
  SceneId *string `json:"sceneId,omitempty" xml:"sceneId,omitempty" require:"true"`
  // {"en":"GetSceneListScene name", "zh_CN":"场景名称"}
  SceneName *string `json:"sceneName,omitempty" xml:"sceneName,omitempty" require:"true"`
  // {"en":"Create time", "zh_CN":"场景创建时间"}
  CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty" require:"true"`
  // {"en":"Create user", "zh_CN":"场景创建的用户"}
  CreateUser *string `json:"createUser,omitempty" xml:"createUser,omitempty" require:"true"`
}

func (s GetSceneListScene) String() string {
  return tea.Prettify(s)
}

func (s GetSceneListScene) GoString() string {
  return s.String()
}

func (s *GetSceneListScene) SetSceneId(v string) *GetSceneListScene {
  s.SceneId = &v
  return s
}

func (s *GetSceneListScene) SetSceneName(v string) *GetSceneListScene {
  s.SceneName = &v
  return s
}

func (s *GetSceneListScene) SetCreateTime(v string) *GetSceneListScene {
  s.CreateTime = &v
  return s
}

func (s *GetSceneListScene) SetCreateUser(v string) *GetSceneListScene {
  s.CreateUser = &v
  return s
}

type GetSceneListPaths struct {
}

func (s GetSceneListPaths) String() string {
  return tea.Prettify(s)
}

func (s GetSceneListPaths) GoString() string {
  return s.String()
}

type GetSceneListParameters struct {
}

func (s GetSceneListParameters) String() string {
  return tea.Prettify(s)
}

func (s GetSceneListParameters) GoString() string {
  return s.String()
}

type GetSceneListRequestHeader struct {
}

func (s GetSceneListRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetSceneListRequestHeader) GoString() string {
  return s.String()
}

type GetSceneListResponseHeader struct {
}

func (s GetSceneListResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetSceneListResponseHeader) GoString() string {
  return s.String()
}




type SceneEditRequest struct {
  // {"en":"Scene name", "zh_CN":"场景名称"}
  SceneName *string `json:"sceneName,omitempty" xml:"sceneName,omitempty" require:"true"`
  // {"en":"The scene Id to edit", "zh_CN":"需要编辑的场景Id"}
  SceneId *string `json:"sceneId,omitempty" xml:"sceneId,omitempty" require:"true"`
}

func (s SceneEditRequest) String() string {
  return tea.Prettify(s)
}

func (s SceneEditRequest) GoString() string {
  return s.String()
}

func (s *SceneEditRequest) SetSceneName(v string) *SceneEditRequest {
  s.SceneName = &v
  return s
}

func (s *SceneEditRequest) SetSceneId(v string) *SceneEditRequest {
  s.SceneId = &v
  return s
}

type SceneEditResponse struct {
  // {"en":"Status code", "zh_CN":"状态码"}
  Code *int `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Operational information", "zh_CN":"操作信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Return data", "zh_CN":"返回数据"}
  Data *string `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s SceneEditResponse) String() string {
  return tea.Prettify(s)
}

func (s SceneEditResponse) GoString() string {
  return s.String()
}

func (s *SceneEditResponse) SetCode(v int) *SceneEditResponse {
  s.Code = &v
  return s
}

func (s *SceneEditResponse) SetMessage(v string) *SceneEditResponse {
  s.Message = &v
  return s
}

func (s *SceneEditResponse) SetData(v string) *SceneEditResponse {
  s.Data = &v
  return s
}

type SceneEditPaths struct {
}

func (s SceneEditPaths) String() string {
  return tea.Prettify(s)
}

func (s SceneEditPaths) GoString() string {
  return s.String()
}

type SceneEditParameters struct {
}

func (s SceneEditParameters) String() string {
  return tea.Prettify(s)
}

func (s SceneEditParameters) GoString() string {
  return s.String()
}

type SceneEditRequestHeader struct {
}

func (s SceneEditRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s SceneEditRequestHeader) GoString() string {
  return s.String()
}

type SceneEditResponseHeader struct {
}

func (s SceneEditResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s SceneEditResponseHeader) GoString() string {
  return s.String()
}




type CreateSceneRequest struct {
  // {"en":"Scene name", "zh_CN":"场景名称"}
  SceneName *string `json:"sceneName,omitempty" xml:"sceneName,omitempty" require:"true"`
}

func (s CreateSceneRequest) String() string {
  return tea.Prettify(s)
}

func (s CreateSceneRequest) GoString() string {
  return s.String()
}

func (s *CreateSceneRequest) SetSceneName(v string) *CreateSceneRequest {
  s.SceneName = &v
  return s
}

type CreateSceneResponse struct {
  // {"en":"Status code", "zh_CN":"状态码"}
  Code *int32 `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Operational information", "zh_CN":"操作信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Return data", "zh_CN":"返回数据"}
  CreateSceneData *CreateSceneData `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s CreateSceneResponse) String() string {
  return tea.Prettify(s)
}

func (s CreateSceneResponse) GoString() string {
  return s.String()
}

func (s *CreateSceneResponse) SetCode(v int32) *CreateSceneResponse {
  s.Code = &v
  return s
}

func (s *CreateSceneResponse) SetMessage(v string) *CreateSceneResponse {
  s.Message = &v
  return s
}

func (s *CreateSceneResponse) SetData(v *CreateSceneData) *CreateSceneResponse {
  s.CreateSceneData = v
  return s
}

type CreateSceneData struct {
  // {"en":"Scene ID", "zh_CN":"场景ID"}
  SceneId *string `json:"sceneId,omitempty" xml:"sceneId,omitempty" require:"true"`
}

func (s CreateSceneData) String() string {
  return tea.Prettify(s)
}

func (s CreateSceneData) GoString() string {
  return s.String()
}

func (s *CreateSceneData) SetSceneId(v string) *CreateSceneData {
  s.SceneId = &v
  return s
}

type CreateScenePaths struct {
}

func (s CreateScenePaths) String() string {
  return tea.Prettify(s)
}

func (s CreateScenePaths) GoString() string {
  return s.String()
}

type CreateSceneParameters struct {
}

func (s CreateSceneParameters) String() string {
  return tea.Prettify(s)
}

func (s CreateSceneParameters) GoString() string {
  return s.String()
}

type CreateSceneRequestHeader struct {
}

func (s CreateSceneRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s CreateSceneRequestHeader) GoString() string {
  return s.String()
}

type CreateSceneResponseHeader struct {
}

func (s CreateSceneResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s CreateSceneResponseHeader) GoString() string {
  return s.String()
}




type DeleteSceneRequest struct {
  // {"en":"id of the scene you want to delete. Multiple users can be deleted at the same time", "zh_CN":"需要删除的场景的id;支持同时删除多个"}
  SceneIds []*string `json:"sceneIds,omitempty" xml:"sceneIds,omitempty" require:"true" type:"Repeated"`
}

func (s DeleteSceneRequest) String() string {
  return tea.Prettify(s)
}

func (s DeleteSceneRequest) GoString() string {
  return s.String()
}

func (s *DeleteSceneRequest) SetSceneIds(v []*string) *DeleteSceneRequest {
  s.SceneIds = v
  return s
}

type DeleteSceneResponse struct {
  // {"en":"Status code", "zh_CN":"状态码"}
  Code *int `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Operational information", "zh_CN":"操作信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Return data", "zh_CN":"返回数据"}
  Data *string `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s DeleteSceneResponse) String() string {
  return tea.Prettify(s)
}

func (s DeleteSceneResponse) GoString() string {
  return s.String()
}

func (s *DeleteSceneResponse) SetCode(v int) *DeleteSceneResponse {
  s.Code = &v
  return s
}

func (s *DeleteSceneResponse) SetMessage(v string) *DeleteSceneResponse {
  s.Message = &v
  return s
}

func (s *DeleteSceneResponse) SetData(v string) *DeleteSceneResponse {
  s.Data = &v
  return s
}

type DeleteScenePaths struct {
}

func (s DeleteScenePaths) String() string {
  return tea.Prettify(s)
}

func (s DeleteScenePaths) GoString() string {
  return s.String()
}

type DeleteSceneParameters struct {
}

func (s DeleteSceneParameters) String() string {
  return tea.Prettify(s)
}

func (s DeleteSceneParameters) GoString() string {
  return s.String()
}

type DeleteSceneRequestHeader struct {
}

func (s DeleteSceneRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteSceneRequestHeader) GoString() string {
  return s.String()
}

type DeleteSceneResponseHeader struct {
}

func (s DeleteSceneResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteSceneResponseHeader) GoString() string {
  return s.String()
}




