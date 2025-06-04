package propertyvalidation

// This file is auto-generated, don't edit it. Thanks.
import (
  "github.com/alibabacloud-go/tea/tea"
)

type GetPropertyValidationTaskPaths struct {
  // {"en" : "ID of a validation task", "zh_CN": "验证任务的id。"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
}

func (s GetPropertyValidationTaskPaths) String() string {
  return tea.Prettify(s)
}

func (s GetPropertyValidationTaskPaths) GoString() string {
  return s.String()
}

func (s *GetPropertyValidationTaskPaths) SetId(v string) *GetPropertyValidationTaskPaths {
  s.Id = &v
  return s
}

type GetPropertyValidationTaskParameters struct {
}

func (s GetPropertyValidationTaskParameters) String() string {
  return tea.Prettify(s)
}

func (s GetPropertyValidationTaskParameters) GoString() string {
  return s.String()
}

type GetPropertyValidationTaskRequestHeader struct {
}

func (s GetPropertyValidationTaskRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetPropertyValidationTaskRequestHeader) GoString() string {
  return s.String()
}

type GetPropertyValidationTaskRequest struct {
}

func (s GetPropertyValidationTaskRequest) String() string {
  return tea.Prettify(s)
}

func (s GetPropertyValidationTaskRequest) GoString() string {
  return s.String()
}

type GetPropertyValidationTaskResponseHeader struct {
}

func (s GetPropertyValidationTaskResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetPropertyValidationTaskResponseHeader) GoString() string {
  return s.String()
}

type GetPropertyValidationTaskResponse struct {
  // {"en" : "ID of the property being validated.", "zh_CN": "提交验证的加速项目的ID。"}
  PropertyId *string `json:"propertyId,omitempty" xml:"propertyId,omitempty" require:"true"`
  // {"en" : "Version of the property being validated.", "zh_CN": "提交验证的加速项目的版本号。"}
  Version *int `json:"version,omitempty" xml:"version,omitempty" require:"true"`
  // {"en" : "An RFC 3339 format string indicating when the task was created. Example: '2019-11-12T03:06:16Z'", "zh_CN": "RFC 3339格式的日期，表示验证任务的提交时间。如：'2019-11-12T03:06:16Z'。"}
  SubmissionTime *string `json:"submissionTime,omitempty" xml:"submissionTime,omitempty" require:"true"`
  // {"en" : "Enum: waiting,failed,succeeded,in progress 
  // Indicates the validation task's status.", "zh_CN": "取值范围: waiting,failed,succeeded,in progress 
  // 验证任务的执行状态，包括等待中，进行中，验证成功，验证失败等状态。"}
  Status *string `json:"status,omitempty" xml:"status,omitempty" require:"true"`
  // {"en" : "Further information about the task's current status.
  // ", "zh_CN": "当前任务状态的描述信息。"}
  StatusDetails *string `json:"statusDetails,omitempty" xml:"statusDetails,omitempty" require:"true"`
  // {"en" : "Refers to the ID of the associated API request.", "zh_CN": "API请求的ID。"}
  ApiRequestId *string `json:"apiRequestId,omitempty" xml:"apiRequestId,omitempty" require:"true"`
  // {"en" : "An RFC 3339 date indicating when the task finished.", "zh_CN": "RFC 3339格式的日期，表示验证任务的完成时间。"}
  FinishTime *string `json:"finishTime,omitempty" xml:"finishTime,omitempty" require:"true"`
  // {"en" : "Name of the property validation task.", "zh_CN": "验证任务的名称。"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en" : "CDN cache version used when validating the property.", "zh_CN": "验证加速项目时使用的CDN Pro cache服务器的版本。"}
  CacheVersion *string `json:"cacheVersion,omitempty" xml:"cacheVersion,omitempty" require:"true"`
  // {"en" : "ID of a webhook to call when the validation task completes.", "zh_CN": "验证任务完成时需要调用的webhook ID。"}
  Webhook *string `json:"webhook,omitempty" xml:"webhook,omitempty" require:"true"`
}

func (s GetPropertyValidationTaskResponse) String() string {
  return tea.Prettify(s)
}

func (s GetPropertyValidationTaskResponse) GoString() string {
  return s.String()
}

func (s *GetPropertyValidationTaskResponse) SetPropertyId(v string) *GetPropertyValidationTaskResponse {
  s.PropertyId = &v
  return s
}

func (s *GetPropertyValidationTaskResponse) SetVersion(v int) *GetPropertyValidationTaskResponse {
  s.Version = &v
  return s
}

func (s *GetPropertyValidationTaskResponse) SetSubmissionTime(v string) *GetPropertyValidationTaskResponse {
  s.SubmissionTime = &v
  return s
}

func (s *GetPropertyValidationTaskResponse) SetStatus(v string) *GetPropertyValidationTaskResponse {
  s.Status = &v
  return s
}

func (s *GetPropertyValidationTaskResponse) SetStatusDetails(v string) *GetPropertyValidationTaskResponse {
  s.StatusDetails = &v
  return s
}

func (s *GetPropertyValidationTaskResponse) SetApiRequestId(v string) *GetPropertyValidationTaskResponse {
  s.ApiRequestId = &v
  return s
}

func (s *GetPropertyValidationTaskResponse) SetFinishTime(v string) *GetPropertyValidationTaskResponse {
  s.FinishTime = &v
  return s
}

func (s *GetPropertyValidationTaskResponse) SetName(v string) *GetPropertyValidationTaskResponse {
  s.Name = &v
  return s
}

func (s *GetPropertyValidationTaskResponse) SetCacheVersion(v string) *GetPropertyValidationTaskResponse {
  s.CacheVersion = &v
  return s
}

func (s *GetPropertyValidationTaskResponse) SetWebhook(v string) *GetPropertyValidationTaskResponse {
  s.Webhook = &v
  return s
}




type GetListOfPropertyValidationTasksPaths struct {
}

func (s GetListOfPropertyValidationTasksPaths) String() string {
  return tea.Prettify(s)
}

func (s GetListOfPropertyValidationTasksPaths) GoString() string {
  return s.String()
}

type GetListOfPropertyValidationTasksParameters struct {
  // {"en" : "Return validation tasks associated only with a specific property ID.", "zh_CN": "指定加速项目ID，查询该加速项目相关的验证任务。"}
  PropertyId *string `json:"propertyId,omitempty" xml:"propertyId,omitempty"`
  // {"en" : "Return validation tasks only for a particular property version.
  // ", "zh_CN": "指定加速项目版本，查询该加速项目版本相关的验证任务。"}
  PropertyVersion *string `json:"propertyVersion,omitempty" xml:"propertyVersion,omitempty"`
  // {"en" : "Enum: waiting,in progress,succeeded,failed 
  // Filter validation tasks per status. By default, all tasks are returned regardless of their status. ", "zh_CN": "取值范围: waiting,in progress,succeeded,failed 
  // 根据状态查询验证任务。默认情况下，查询所有状态下的验证任务。"}
  Status *string `json:"status,omitempty" xml:"status,omitempty"`
  // {"en" : "Default: 100 Range: [ 1 .. 200 ] 
  // Indicates the maximum number of validation tasks to return.", "zh_CN": "默认值: 100 取值范围: <= 200 
  // 每次查询的最大条数。"}
  Limit *int `json:"limit,omitempty" xml:"limit,omitempty"`
  // {"en" : "Default: 0 Range: >= 0 
  // Indicates the index of the first validation task to return.", "zh_CN": "默认值: 0 取值范围: >= 0 
  // 查询起始位置。"}
  Offset *int `json:"offset,omitempty" xml:"offset,omitempty"`
  // {"en" : "The results will be filtered based on the presence of the value as a validation task ID or in a task name.", "zh_CN": "根据该关键字匹配验证任务的ID或者名称进行查询。"}
  Search *string `json:"search,omitempty" xml:"search,omitempty"`
  // {"en" : "Enum: submissionTime,finishTime 
  // Default: finishTime 
  // Return results sorted by when the task was submitted (submissionTime) or when the task was completed (finishTime). ", "zh_CN": "取值范围: submissionTime,finishTime 
  // 默认值: finishTime 
  // 查询结果根据验证任务的提交时间（submissionTime）或者完成时间（finishTime）进行排序。"}
  SortBy *string `json:"sortBy,omitempty" xml:"sortBy,omitempty"`
  // {"en" : "Enum: asc,desc 
  // Default: desc 
  // Return results sorted in ascending (asc) or descending (desc) order.
  // ", "zh_CN": "取值范围: asc,desc 
  // 默认值: desc 
  // 查询结果按照升序（asc）或者降序（desc）返回。"}
  SortOrder *string `json:"sortOrder,omitempty" xml:"sortOrder,omitempty"`
  // {"en" : "Enter an RFC 3339 format date to return validation tasks which were submitted by this date.", "zh_CN": "指定RFC 3339格式的日期，查询提交时间晚于该日期的验证任务。"}
  Startdate *string `json:"startdate,omitempty" xml:"startdate,omitempty"`
  // {"en" : "Enter an RFC 3339 format date to return validation tasks which were submitted no later than this date.", "zh_CN": "指定RFC 3339格式的日期，查询提交时间早于该日期的验证任务。"}
  Enddate *string `json:"enddate,omitempty" xml:"enddate,omitempty"`
}

func (s GetListOfPropertyValidationTasksParameters) String() string {
  return tea.Prettify(s)
}

func (s GetListOfPropertyValidationTasksParameters) GoString() string {
  return s.String()
}

func (s *GetListOfPropertyValidationTasksParameters) SetPropertyId(v string) *GetListOfPropertyValidationTasksParameters {
  s.PropertyId = &v
  return s
}

func (s *GetListOfPropertyValidationTasksParameters) SetPropertyVersion(v string) *GetListOfPropertyValidationTasksParameters {
  s.PropertyVersion = &v
  return s
}

func (s *GetListOfPropertyValidationTasksParameters) SetStatus(v string) *GetListOfPropertyValidationTasksParameters {
  s.Status = &v
  return s
}

func (s *GetListOfPropertyValidationTasksParameters) SetLimit(v int) *GetListOfPropertyValidationTasksParameters {
  s.Limit = &v
  return s
}

func (s *GetListOfPropertyValidationTasksParameters) SetOffset(v int) *GetListOfPropertyValidationTasksParameters {
  s.Offset = &v
  return s
}

func (s *GetListOfPropertyValidationTasksParameters) SetSearch(v string) *GetListOfPropertyValidationTasksParameters {
  s.Search = &v
  return s
}

func (s *GetListOfPropertyValidationTasksParameters) SetSortBy(v string) *GetListOfPropertyValidationTasksParameters {
  s.SortBy = &v
  return s
}

func (s *GetListOfPropertyValidationTasksParameters) SetSortOrder(v string) *GetListOfPropertyValidationTasksParameters {
  s.SortOrder = &v
  return s
}

func (s *GetListOfPropertyValidationTasksParameters) SetStartdate(v string) *GetListOfPropertyValidationTasksParameters {
  s.Startdate = &v
  return s
}

func (s *GetListOfPropertyValidationTasksParameters) SetEnddate(v string) *GetListOfPropertyValidationTasksParameters {
  s.Enddate = &v
  return s
}

type GetListOfPropertyValidationTasksRequestHeader struct {
}

func (s GetListOfPropertyValidationTasksRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetListOfPropertyValidationTasksRequestHeader) GoString() string {
  return s.String()
}

type GetListOfPropertyValidationTasksRequest struct {
}

func (s GetListOfPropertyValidationTasksRequest) String() string {
  return tea.Prettify(s)
}

func (s GetListOfPropertyValidationTasksRequest) GoString() string {
  return s.String()
}

type GetListOfPropertyValidationTasksResponseHeader struct {
}

func (s GetListOfPropertyValidationTasksResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetListOfPropertyValidationTasksResponseHeader) GoString() string {
  return s.String()
}

type GetListOfPropertyValidationTasksResponse struct {
  // {"en" : "Range: >= 0 
  // Total number of validation tasks", "zh_CN": "取值范围: >= 0 
  // 验证任务的总数。"}
  Count *int `json:"count,omitempty" xml:"count,omitempty" require:"true"`
  // {"en" : "Summarizes validation tasks. Further details about a validation task can be obtained by calling the Query a property validation task API.", "zh_CN": "验证任务摘要信息。通过调用'查询验证任务详情'接口，可获得有关验证任务的详细信息。"}
  Validations []*GetListOfPropertyValidationTasksResponseValidations `json:"validations,omitempty" xml:"validations,omitempty" require:"true" type:"Repeated"`
}

func (s GetListOfPropertyValidationTasksResponse) String() string {
  return tea.Prettify(s)
}

func (s GetListOfPropertyValidationTasksResponse) GoString() string {
  return s.String()
}

func (s *GetListOfPropertyValidationTasksResponse) SetCount(v int) *GetListOfPropertyValidationTasksResponse {
  s.Count = &v
  return s
}

func (s *GetListOfPropertyValidationTasksResponse) SetValidations(v []*GetListOfPropertyValidationTasksResponseValidations) *GetListOfPropertyValidationTasksResponse {
  s.Validations = v
  return s
}

type GetListOfPropertyValidationTasksResponseValidations struct     {
  // {"en" : "ID of the validation task.", "zh_CN": "验证任务ID。"}
  Id *string `json:"id,omitempty" xml:"id,omitempty"`
  // {"en" : "ID of the property that was validated.", "zh_CN": "提交验证的加速项目的ID。"}
  PropertyId *string `json:"propertyId,omitempty" xml:"propertyId,omitempty"`
  // {"en" : "Range: >= 1 
  // Version of the property that was validated.", "zh_CN": "取值范围: >= 1 
  // 提交验证的加速项目的版本。"}
  Version *int `json:"version,omitempty" xml:"version,omitempty"`
  // {"en" : "An RFC 3339 format string indicating when the validation task was submitted. Example: '2019-11-12T03:09:26Z'", "zh_CN": "RFC 3339格式的日期，表示验证任务的提交时间。如：'2019-11-12T03:09:26Z'。"}
  SubmissionTime *string `json:"submissionTime,omitempty" xml:"submissionTime,omitempty"`
  // {"en" : "Enum: waiting,in progress,succeeded,failed 
  // Status of the validation task.", "zh_CN": "取值范围: waiting,in progress,succeeded,failed 
  // 验证任务的执行状态，包括等待中，进行中，验证成功以及验证失败等状态。"}
  Status *string `json:"status,omitempty" xml:"status,omitempty"`
  // {"en" : "Identifies the internal cache version on which the property was validated.", "zh_CN": "验证加速项目时使用的CDN Pro cache服务器的版本。"}
  CacheVersion *string `json:"cacheVersion,omitempty" xml:"cacheVersion,omitempty"`
  // {"en" : "Indicates the associated API request.", "zh_CN": "API请求ID。"}
  ApiRequestId *string `json:"apiRequestId,omitempty" xml:"apiRequestId,omitempty"`
}

func (s GetListOfPropertyValidationTasksResponseValidations) String() string {
  return tea.Prettify(s)
}

func (s GetListOfPropertyValidationTasksResponseValidations) GoString() string {
  return s.String()
}

func (s *GetListOfPropertyValidationTasksResponseValidations) SetId(v string) *GetListOfPropertyValidationTasksResponseValidations {
  s.Id = &v
  return s
}

func (s *GetListOfPropertyValidationTasksResponseValidations) SetPropertyId(v string) *GetListOfPropertyValidationTasksResponseValidations {
  s.PropertyId = &v
  return s
}

func (s *GetListOfPropertyValidationTasksResponseValidations) SetVersion(v int) *GetListOfPropertyValidationTasksResponseValidations {
  s.Version = &v
  return s
}

func (s *GetListOfPropertyValidationTasksResponseValidations) SetSubmissionTime(v string) *GetListOfPropertyValidationTasksResponseValidations {
  s.SubmissionTime = &v
  return s
}

func (s *GetListOfPropertyValidationTasksResponseValidations) SetStatus(v string) *GetListOfPropertyValidationTasksResponseValidations {
  s.Status = &v
  return s
}

func (s *GetListOfPropertyValidationTasksResponseValidations) SetCacheVersion(v string) *GetListOfPropertyValidationTasksResponseValidations {
  s.CacheVersion = &v
  return s
}

func (s *GetListOfPropertyValidationTasksResponseValidations) SetApiRequestId(v string) *GetListOfPropertyValidationTasksResponseValidations {
  s.ApiRequestId = &v
  return s
}




type CreateAValidationTaskPaths struct {
}

func (s CreateAValidationTaskPaths) String() string {
  return tea.Prettify(s)
}

func (s CreateAValidationTaskPaths) GoString() string {
  return s.String()
}

type CreateAValidationTaskParameters struct {
}

func (s CreateAValidationTaskParameters) String() string {
  return tea.Prettify(s)
}

func (s CreateAValidationTaskParameters) GoString() string {
  return s.String()
}

type CreateAValidationTaskRequestHeader struct {
}

func (s CreateAValidationTaskRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s CreateAValidationTaskRequestHeader) GoString() string {
  return s.String()
}

type CreateAValidationTaskRequest struct {
  // {"en" : "Name of the validation task.", "zh_CN": "验证任务的名称。"}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // {"en" : "Indicates the property to validate.", "zh_CN": "需要验证的加速项目的ID。"}
  PropertyId *string `json:"propertyId,omitempty" xml:"propertyId,omitempty" require:"true"`
  // {"en" : "Indicates the version of the property to validate.", "zh_CN": "需要验证的加速项目的版本号。"}
  Version *int `json:"version,omitempty" xml:"version,omitempty" require:"true"`
  // {"en" : "ID of a webhook to call when the validation task completes.", "zh_CN": "验证任务完成时需要调用的webhook ID。webhook是指通过“创建webhook接口”创建的回调接口。"}
  Webhook *string `json:"webhook,omitempty" xml:"webhook,omitempty"`
}

func (s CreateAValidationTaskRequest) String() string {
  return tea.Prettify(s)
}

func (s CreateAValidationTaskRequest) GoString() string {
  return s.String()
}

func (s *CreateAValidationTaskRequest) SetName(v string) *CreateAValidationTaskRequest {
  s.Name = &v
  return s
}

func (s *CreateAValidationTaskRequest) SetPropertyId(v string) *CreateAValidationTaskRequest {
  s.PropertyId = &v
  return s
}

func (s *CreateAValidationTaskRequest) SetVersion(v int) *CreateAValidationTaskRequest {
  s.Version = &v
  return s
}

func (s *CreateAValidationTaskRequest) SetWebhook(v string) *CreateAValidationTaskRequest {
  s.Webhook = &v
  return s
}

type CreateAValidationTaskResponse struct {
}

func (s CreateAValidationTaskResponse) String() string {
  return tea.Prettify(s)
}

func (s CreateAValidationTaskResponse) GoString() string {
  return s.String()
}

type CreateAValidationTaskResponseHeader struct {
  // {"en":"The Location header contains a URL to the validation task. Example: <code>Location: https://{domain}/cdn/validations/5dca2205f9e9cc0001df7b24</code>", "zh_CN":"通过Location响应头返回新建的验证任务的URL。URL中包含了验证任务的ID，可使用该ID调用'查询验证任务详情'接口来查看任务详情。URL示例：<code>Location: https://{domain}/cdn/validations/5dca2205f9e9cc0001df7b24</code>"}
  Location *string `json:"Location,omitempty" xml:"Location,omitempty" require:"true"`
}

func (s CreateAValidationTaskResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s CreateAValidationTaskResponseHeader) GoString() string {
  return s.String()
}

func (s *CreateAValidationTaskResponseHeader) SetLocation(v string) *CreateAValidationTaskResponseHeader {
  s.Location = &v
  return s
}




