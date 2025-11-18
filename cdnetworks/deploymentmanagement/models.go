package deploymentmanagement

// This file is auto-generated, don't edit it. Thanks.
import (
  "github.com/alibabacloud-go/tea/tea"
)

type GetDeploymentTaskRequest struct {
}

func (s GetDeploymentTaskRequest) String() string {
  return tea.Prettify(s)
}

func (s GetDeploymentTaskRequest) GoString() string {
  return s.String()
}

type GetDeploymentTaskRequestHeader struct {
}

func (s GetDeploymentTaskRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetDeploymentTaskRequestHeader) GoString() string {
  return s.String()
}

type GetDeploymentTaskPaths struct {
  // {"en":"ID of the deployment task.","zh_CN":"部署任务id。"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
}

func (s GetDeploymentTaskPaths) String() string {
  return tea.Prettify(s)
}

func (s GetDeploymentTaskPaths) GoString() string {
  return s.String()
}

func (s *GetDeploymentTaskPaths) SetId(v string) *GetDeploymentTaskPaths {
  s.Id = &v
  return s
}

type GetDeploymentTaskParameters struct {
}

func (s GetDeploymentTaskParameters) String() string {
  return tea.Prettify(s)
}

func (s GetDeploymentTaskParameters) GoString() string {
  return s.String()
}

type GetDeploymentTaskResponse struct {
  // {"en":"Name of the deployment task.","zh_CN":"部署任务的名称。"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"RFC 3339 date indicating when the task was created.","zh_CN":"RFC 3339格式的日期，表示任务的创建时间。"}
  SubmissionTime *string `json:"submissionTime,omitempty" xml:"submissionTime,omitempty" require:"true"`
  // {"en":"This array contains all the actions related to a deployment. They can include deployment and removal of properties and certificates to the staging or production environments.","zh_CN":"部署任务所要执行的操作，可以包括加速项目和证书的部署或卸载操作。"}
  Actions []*GetDeploymentTaskResponseActions `json:"actions,omitempty" xml:"actions,omitempty" require:"true" type:"Repeated"`
  // {"en":"Enum: staging, production\nIndicates the environment affected by the deployment.","zh_CN":"取值范围: staging, production\n部署任务对应的环境。"}
  Target *string `json:"target,omitempty" xml:"target,omitempty" require:"true"`
  // {"en":"RFC 3339 date indicating when the task completed.","zh_CN":"RFC 3339格式的日期，表示任务完成的时间。"}
  FinishTime *string `json:"finishTime,omitempty" xml:"finishTime,omitempty" require:"true"`
  // {"en":"Enum: waiting, in progress, failed, succeeded\nIndicates the status of the deployment task.","zh_CN":"取值范围: waiting, in progress, failed, succeeded\n部署任务的执行状态，包括等待中，进行中，部署失败，部署成功等状态。"}
  Status *string `json:"status,omitempty" xml:"status,omitempty" require:"true"`
  // {"en":"Additional information about the status of the deployment task.","zh_CN":"部署任务执行状态的描述信息。"}
  StatusDetails *string `json:"statusDetails,omitempty" xml:"statusDetails,omitempty" require:"true"`
  // {"en":"ID of the API call to create the deployment task.","zh_CN":"创建部署任务时的API请求ID。"}
  ApiRequestId *string `json:"apiRequestId,omitempty" xml:"apiRequestId,omitempty" require:"true"`
  // {"en":"ID of a webhook to call when the deployment task completes.","zh_CN":"部署任务完成时要调用的webhook的ID。"}
  Webhook *string `json:"webhook,omitempty" xml:"webhook,omitempty" require:"true"`
}

func (s GetDeploymentTaskResponse) String() string {
  return tea.Prettify(s)
}

func (s GetDeploymentTaskResponse) GoString() string {
  return s.String()
}

func (s *GetDeploymentTaskResponse) SetName(v string) *GetDeploymentTaskResponse {
  s.Name = &v
  return s
}

func (s *GetDeploymentTaskResponse) SetSubmissionTime(v string) *GetDeploymentTaskResponse {
  s.SubmissionTime = &v
  return s
}

func (s *GetDeploymentTaskResponse) SetActions(v []*GetDeploymentTaskResponseActions) *GetDeploymentTaskResponse {
  s.Actions = v
  return s
}

func (s *GetDeploymentTaskResponse) SetTarget(v string) *GetDeploymentTaskResponse {
  s.Target = &v
  return s
}

func (s *GetDeploymentTaskResponse) SetFinishTime(v string) *GetDeploymentTaskResponse {
  s.FinishTime = &v
  return s
}

func (s *GetDeploymentTaskResponse) SetStatus(v string) *GetDeploymentTaskResponse {
  s.Status = &v
  return s
}

func (s *GetDeploymentTaskResponse) SetStatusDetails(v string) *GetDeploymentTaskResponse {
  s.StatusDetails = &v
  return s
}

func (s *GetDeploymentTaskResponse) SetApiRequestId(v string) *GetDeploymentTaskResponse {
  s.ApiRequestId = &v
  return s
}

func (s *GetDeploymentTaskResponse) SetWebhook(v string) *GetDeploymentTaskResponse {
  s.Webhook = &v
  return s
}

type GetDeploymentTaskResponseActions struct     {
  // {"en":"Enum: deploy_property, remove_property, deploy_cert, remove_cert\nDescribe an action to take. You can deploy a property, remove a property, deploy a certificate, or remove a certificate.","zh_CN":"取值范围: deploy_property, remove_property, deploy_cert, remove_cert\n指定操作类型，包括部署加速项目、卸载加速项目、部署证书以及卸载证书。"}
  Action *string `json:"action,omitempty" xml:"action,omitempty" require:"true"`
  // {"en":"ID of the property to deploy or remove from the staging or production environment.","zh_CN":"指定要部署或卸载的加速项目ID。"}
  PropertyId *string `json:"propertyId,omitempty" xml:"propertyId,omitempty" require:"true"`
  // {"en":"Indicates the certificate to deploy or remove from the staging or production environment.","zh_CN":"指定要部署或卸载的证书ID。"}
  CertificateId *string `json:"certificateId,omitempty" xml:"certificateId,omitempty" require:"true"`
  // {"en":"Indicates the version of the certificate or property to deploy or remove.","zh_CN":"指定要部署或卸载的证书或加速项目的版本。"}
  Version *string `json:"version,omitempty" xml:"version,omitempty" require:"true"`
}

func (s GetDeploymentTaskResponseActions) String() string {
  return tea.Prettify(s)
}

func (s GetDeploymentTaskResponseActions) GoString() string {
  return s.String()
}

func (s *GetDeploymentTaskResponseActions) SetAction(v string) *GetDeploymentTaskResponseActions {
  s.Action = &v
  return s
}

func (s *GetDeploymentTaskResponseActions) SetPropertyId(v string) *GetDeploymentTaskResponseActions {
  s.PropertyId = &v
  return s
}

func (s *GetDeploymentTaskResponseActions) SetCertificateId(v string) *GetDeploymentTaskResponseActions {
  s.CertificateId = &v
  return s
}

func (s *GetDeploymentTaskResponseActions) SetVersion(v string) *GetDeploymentTaskResponseActions {
  s.Version = &v
  return s
}

type GetDeploymentTaskResponseHeader struct {
}

func (s GetDeploymentTaskResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetDeploymentTaskResponseHeader) GoString() string {
  return s.String()
}




type CreateADeploymentTaskRequest struct {
  // {"en":"Name of the deployment task.","zh_CN":"部署任务的名称。"}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // {"en":"Enum: staging, production\nIndicates whether to deploy to staging or production.","zh_CN":"取值范围: staging, production\n指定部署任务的目标环境，即演练或生产环境。"}
  Target *string `json:"target,omitempty" xml:"target,omitempty" require:"true"`
  // {"en":"This array contains all the actions related to a deployment. They can include deployment and removal of properties and certificates to the staging or production environments.","zh_CN":"部署任务所要执行的操作，可以包括加速项目和证书的部署或卸载操作。"}
  Actions []*CreateADeploymentTaskRequestActions `json:"actions,omitempty" xml:"actions,omitempty" require:"true" type:"Repeated"`
  // {"en":"ID of a webhook to call when the deployment task completes.","zh_CN":"异步部署任务完成时需要回调的webhook的ID。"}
  Webhook *string `json:"webhook,omitempty" xml:"webhook,omitempty"`
}

func (s CreateADeploymentTaskRequest) String() string {
  return tea.Prettify(s)
}

func (s CreateADeploymentTaskRequest) GoString() string {
  return s.String()
}

func (s *CreateADeploymentTaskRequest) SetName(v string) *CreateADeploymentTaskRequest {
  s.Name = &v
  return s
}

func (s *CreateADeploymentTaskRequest) SetTarget(v string) *CreateADeploymentTaskRequest {
  s.Target = &v
  return s
}

func (s *CreateADeploymentTaskRequest) SetActions(v []*CreateADeploymentTaskRequestActions) *CreateADeploymentTaskRequest {
  s.Actions = v
  return s
}

func (s *CreateADeploymentTaskRequest) SetWebhook(v string) *CreateADeploymentTaskRequest {
  s.Webhook = &v
  return s
}

type CreateADeploymentTaskRequestActions struct     {
  // {"en":"Enum: deploy_property, remove_property, deploy_cert, remove_cert\nDescribe an action to take. You can deploy a property, remove a property, deploy a certificate, or remove a certificate.","zh_CN":"取值范围: deploy_property, remove_property, deploy_cert, remove_cert\n指定操作类型，包括部署加速项目、卸载加速项目、部署证书以及卸载证书。"}
  Action *string `json:"action,omitempty" xml:"action,omitempty" require:"true"`
  // {"en":"ID of the property to deploy or remove from the staging or production environment. Required when the action is deploy_property or remove_property.","zh_CN":"指定要部署或卸载的加速项目ID。当部署或卸载加速项目时，必填。"}
  PropertyId *string `json:"propertyId,omitempty" xml:"propertyId,omitempty"`
  // {"en":"Indicates the certificate to deploy or remove from the staging or production environment. Required when the action is deploy_cert or remove_cert.","zh_CN":"指定要部署或卸载的证书ID。当部署或卸载证书时，必填。"}
  CertificateId *string `json:"certificateId,omitempty" xml:"certificateId,omitempty"`
  // {"en":"Indicates the version of the certificate or property to deploy. Required when the action is deploy_property or deploy_cert.","zh_CN":"指定证书或加速项目的版本。当部署证书或部署加速项目时，必填。"}
  Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s CreateADeploymentTaskRequestActions) String() string {
  return tea.Prettify(s)
}

func (s CreateADeploymentTaskRequestActions) GoString() string {
  return s.String()
}

func (s *CreateADeploymentTaskRequestActions) SetAction(v string) *CreateADeploymentTaskRequestActions {
  s.Action = &v
  return s
}

func (s *CreateADeploymentTaskRequestActions) SetPropertyId(v string) *CreateADeploymentTaskRequestActions {
  s.PropertyId = &v
  return s
}

func (s *CreateADeploymentTaskRequestActions) SetCertificateId(v string) *CreateADeploymentTaskRequestActions {
  s.CertificateId = &v
  return s
}

func (s *CreateADeploymentTaskRequestActions) SetVersion(v string) *CreateADeploymentTaskRequestActions {
  s.Version = &v
  return s
}

type CreateADeploymentTaskRequestHeader struct {
  // {"en":"By default, if you deploy a property using a certificate, we check that the hostnames in the property are covered by the subject alternative name (SAN) of the certificate and that the certificate has not expired. For test purposes, you can bypass this check by passing the <i>Check-Certificate: no</i> header when creating the deployment task. You should not use this header when deploying for real users of your content since their browsers will warn about the invalid certificate and discourage them from visiting.","zh_CN":"当您提交部署任务时，后台会校验加速项目所引用的证书是否已过期，以及加速域名是否与证书的授权域名匹配。如果您的部署操作是出于测试目的，您可以在调用接口时传入请求头<i>Check-Certificate: no</i>来跳过校验。非测试目的，不建议跳过校验。"}
  CheckCertificate *string `json:"Check-Certificate,omitempty" xml:"Check-Certificate,omitempty"`
  // {"en":"When you request to undeploy a property, we check if the hostnames in the property are still active. Undeployment of the property will be rejected if there are requests to the hostnames in the past 24 hours. This is to prevent accidental undeployment. If you are sure that the property can be undeployed, you can bypass this check by passing the <i>Check-Usage: no</i> header.","zh_CN":"当您卸载加速项目时，后台会校验加速项目下的域名是否处于活跃状态。如果域名在过去24小时仍有产生请求，则会拒绝卸载加速项目，避免误操作。如果您确定要卸载加速项目，可以在调用接口时传入请求头<i>Check-Usage: no</i>来跳过校验。"}
  CheckUsage *string `json:"Check-Usage,omitempty" xml:"Check-Usage,omitempty"`
}

func (s CreateADeploymentTaskRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s CreateADeploymentTaskRequestHeader) GoString() string {
  return s.String()
}

func (s *CreateADeploymentTaskRequestHeader) SetCheckCertificate(v string) *CreateADeploymentTaskRequestHeader {
  s.CheckCertificate = &v
  return s
}

func (s *CreateADeploymentTaskRequestHeader) SetCheckUsage(v string) *CreateADeploymentTaskRequestHeader {
  s.CheckUsage = &v
  return s
}

type CreateADeploymentTaskPaths struct {
}

func (s CreateADeploymentTaskPaths) String() string {
  return tea.Prettify(s)
}

func (s CreateADeploymentTaskPaths) GoString() string {
  return s.String()
}

type CreateADeploymentTaskParameters struct {
}

func (s CreateADeploymentTaskParameters) String() string {
  return tea.Prettify(s)
}

func (s CreateADeploymentTaskParameters) GoString() string {
  return s.String()
}

type CreateADeploymentTaskResponse struct {
}

func (s CreateADeploymentTaskResponse) String() string {
  return tea.Prettify(s)
}

func (s CreateADeploymentTaskResponse) GoString() string {
  return s.String()
}

type CreateADeploymentTaskResponseHeader struct {
  // {"en":"Returns a URL pointing to the new deployment task created, if the request is accepted. The URL contains the ID of the new deployment task. </br> URL format: <code>{scheme}://{domain}/cdn/deploymentTasks/{deploymentTaskId}</code> Example URL: <code>https://api.example.com/cdn/deploymentTasks/ac8e3085-ef92-4c12-ab1b-19b18ac9383c</code>","zh_CN":"当接口调用成功时，通过Location响应头返回新建的部署任务的URL。URL中包含部署任务的ID，可使用该URL调用'获取部署任务详细信息'接口来查看部署任务的详细信息。</br> URL格式：<code>{协议}://{域名}/cdn/deploymentTasks/{部署任务ID}</code> URL示例： <code>https://open.chinanetcenter.com/cdn/deploymentTasks/ac8e3085-ef92-4c12-ab1b-19b18ac9383c</code>"}
  Location *string `json:"Location,omitempty" xml:"Location,omitempty" require:"true"`
}

func (s CreateADeploymentTaskResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s CreateADeploymentTaskResponseHeader) GoString() string {
  return s.String()
}

func (s *CreateADeploymentTaskResponseHeader) SetLocation(v string) *CreateADeploymentTaskResponseHeader {
  s.Location = &v
  return s
}




type GetListOfDeploymentTasksRequest struct {
}

func (s GetListOfDeploymentTasksRequest) String() string {
  return tea.Prettify(s)
}

func (s GetListOfDeploymentTasksRequest) GoString() string {
  return s.String()
}

type GetListOfDeploymentTasksRequestHeader struct {
}

func (s GetListOfDeploymentTasksRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetListOfDeploymentTasksRequestHeader) GoString() string {
  return s.String()
}

type GetListOfDeploymentTasksPaths struct {
}

func (s GetListOfDeploymentTasksPaths) String() string {
  return tea.Prettify(s)
}

func (s GetListOfDeploymentTasksPaths) GoString() string {
  return s.String()
}

type GetListOfDeploymentTasksParameters struct {
  // {"en":"Default: 0 Range: >= 0\nIndex of the first task to return.","zh_CN":"默认值: 0 取值范围: >= 0\n查询起始位置。"}
  Offset *int `json:"offset,omitempty" xml:"offset,omitempty"`
  // {"en":"Default: 100 Range: <= 200\nMaximum number of tasks to return.","zh_CN":"默认值: 100 取值范围: <= 200\n每次查询的最大条数。"}
  Limit *int `json:"limit,omitempty" xml:"limit,omitempty"`
  // {"en":"Only return deployment tasks related to this property.","zh_CN":"指定加速项目ID，查询该加速项目相关的部署任务。"}
  PropertyId *string `json:"propertyId,omitempty" xml:"propertyId,omitempty"`
  // {"en":"Only return deployment tasks related to this certificate.","zh_CN":"指定证书ID，查询该证书相关的部署任务。"}
  CertificateId *string `json:"certificateId,omitempty" xml:"certificateId,omitempty"`
  // {"en":"Enum: staging, production </br>Return deployment tasks to a particular environment. By default, all tasks are returned.","zh_CN":"取值范围: staging, production </br>根据部署环境查询。默认情况下，查询所有环境的部署任务。"}
  Target *string `json:"target,omitempty" xml:"target,omitempty"`
  // {"en":"Enum: waiting, inprogress, succeeded, failed </br>Filter deployment tasks per status. By default, all tasks are returned regardless of their status.","zh_CN":"取值范围: waiting, inprogress, succeeded, failed </br>根据状态查询部署任务。默认情况下，查询所有状态下的部署任务。"}
  Status *string `json:"status,omitempty" xml:"status,omitempty"`
  // {"en":"Filter based on text in the deployment name or in a property's hostname.","zh_CN":"根据搜索关键字匹配部署任务名称及加速项目下的域名进行查询。"}
  Search *string `json:"search,omitempty" xml:"search,omitempty"`
  // {"en":"A comma-separated list of task IDs indicating which ones to return.","zh_CN":"根据部署任务ID查询。可指定多个ID，以逗号分隔。"}
  TaskIds *string `json:"taskIds,omitempty" xml:"taskIds,omitempty"`
  // {"en":"Enum: submissionTime, finishTime </br>Default: finishTimeReturn results sorted by this value.","zh_CN":"取值范围: submissionTime, finishTime </br>默认值: finishTime根据任务提交时间或完成时间进行排序。"}
  SortBy *string `json:"sortBy,omitempty" xml:"sortBy,omitempty"`
  // {"en":"Enum: asc, descDefault: descReturn results in this order.","zh_CN":"取值范围: asc, desc默认值: desc指定排序顺序。"}
  SortOrder *string `json:"sortOrder,omitempty" xml:"sortOrder,omitempty"`
  // {"en":"Enter an RFC 3339 format date to return deployments which were submitted by this date.","zh_CN":"指定RFC 3339格式的日期，查询提交时间晚于该日期的部署任务。"}
  Startdate *string `json:"startdate,omitempty" xml:"startdate,omitempty"`
  // {"en":"Enter an RFC 3339 format date to return deployments which were submitted no later than this date.","zh_CN":"指定RFC 3339格式的日期，查询提交时间早于该日期的部署任务。"}
  Enddate *string `json:"enddate,omitempty" xml:"enddate,omitempty"`
}

func (s GetListOfDeploymentTasksParameters) String() string {
  return tea.Prettify(s)
}

func (s GetListOfDeploymentTasksParameters) GoString() string {
  return s.String()
}

func (s *GetListOfDeploymentTasksParameters) SetOffset(v int) *GetListOfDeploymentTasksParameters {
  s.Offset = &v
  return s
}

func (s *GetListOfDeploymentTasksParameters) SetLimit(v int) *GetListOfDeploymentTasksParameters {
  s.Limit = &v
  return s
}

func (s *GetListOfDeploymentTasksParameters) SetPropertyId(v string) *GetListOfDeploymentTasksParameters {
  s.PropertyId = &v
  return s
}

func (s *GetListOfDeploymentTasksParameters) SetCertificateId(v string) *GetListOfDeploymentTasksParameters {
  s.CertificateId = &v
  return s
}

func (s *GetListOfDeploymentTasksParameters) SetTarget(v string) *GetListOfDeploymentTasksParameters {
  s.Target = &v
  return s
}

func (s *GetListOfDeploymentTasksParameters) SetStatus(v string) *GetListOfDeploymentTasksParameters {
  s.Status = &v
  return s
}

func (s *GetListOfDeploymentTasksParameters) SetSearch(v string) *GetListOfDeploymentTasksParameters {
  s.Search = &v
  return s
}

func (s *GetListOfDeploymentTasksParameters) SetTaskIds(v string) *GetListOfDeploymentTasksParameters {
  s.TaskIds = &v
  return s
}

func (s *GetListOfDeploymentTasksParameters) SetSortBy(v string) *GetListOfDeploymentTasksParameters {
  s.SortBy = &v
  return s
}

func (s *GetListOfDeploymentTasksParameters) SetSortOrder(v string) *GetListOfDeploymentTasksParameters {
  s.SortOrder = &v
  return s
}

func (s *GetListOfDeploymentTasksParameters) SetStartdate(v string) *GetListOfDeploymentTasksParameters {
  s.Startdate = &v
  return s
}

func (s *GetListOfDeploymentTasksParameters) SetEnddate(v string) *GetListOfDeploymentTasksParameters {
  s.Enddate = &v
  return s
}

type GetListOfDeploymentTasksResponse struct {
  // {"en":"List of deployment task summaries.","zh_CN":"部署任务列表。"}
  Deploy []*GetListOfDeploymentTasksResponseDeploy `json:"deploy,omitempty" xml:"deploy,omitempty" require:"true" type:"Repeated"`
  // {"en":"Range: >= 0\nTotal number of deployment tasks.","zh_CN":"取值范围: >= 0\n部署任务的总数。"}
  Count *int `json:"count,omitempty" xml:"count,omitempty" require:"true"`
}

func (s GetListOfDeploymentTasksResponse) String() string {
  return tea.Prettify(s)
}

func (s GetListOfDeploymentTasksResponse) GoString() string {
  return s.String()
}

func (s *GetListOfDeploymentTasksResponse) SetDeploy(v []*GetListOfDeploymentTasksResponseDeploy) *GetListOfDeploymentTasksResponse {
  s.Deploy = v
  return s
}

func (s *GetListOfDeploymentTasksResponse) SetCount(v int) *GetListOfDeploymentTasksResponse {
  s.Count = &v
  return s
}

type GetListOfDeploymentTasksResponseDeploy struct     {
  // {"en":"ID representing the deployment task. You can obtain more information about a task by calling the Query deployment task API.","zh_CN":"部署任务的ID。您可以通过调用'获取部署任务的详细信息'接口来获取部署任务的更多信息。"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {"en":"Name of the deployment task.","zh_CN":"部署任务的名称。"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"An RFC 3339 format date indicates when the task was submitted.","zh_CN":"RFC 3339格式的日期，表示任务的提交时间。"}
  SubmissionTime *string `json:"submissionTime,omitempty" xml:"submissionTime,omitempty" require:"true"`
  // {"en":"An RFC 3339 date indicates when the task completed.","zh_CN":"RFC 3339格式的日期，表示任务的完成时间。"}
  FinishTime *string `json:"finishTime,omitempty" xml:"finishTime,omitempty" require:"true"`
  // {"en":"Enum: waiting, inprogress, failed, succeeded\nIndicates the status of the deployment task.","zh_CN":"取值范围: waiting, inprogress, failed, succeeded\n部署任务的执行状态，包括等待中，执行中，部署失败，部署成功等状态。"}
  Status *string `json:"status,omitempty" xml:"status,omitempty" require:"true"`
  // {"en":"Enum: staging, production\nIndicates where the deployment will go to.","zh_CN":"取值范围: staging, production\n部署任务对应的环境，即演练环境或生产环境。"}
  Target *string `json:"target,omitempty" xml:"target,omitempty" require:"true"`
  // {"en":"An internal ID indicating the associated API call.","zh_CN":"API请求的ID。"}
  ApiRequestId *string `json:"apiRequestId,omitempty" xml:"apiRequestId,omitempty" require:"true"`
}

func (s GetListOfDeploymentTasksResponseDeploy) String() string {
  return tea.Prettify(s)
}

func (s GetListOfDeploymentTasksResponseDeploy) GoString() string {
  return s.String()
}

func (s *GetListOfDeploymentTasksResponseDeploy) SetId(v string) *GetListOfDeploymentTasksResponseDeploy {
  s.Id = &v
  return s
}

func (s *GetListOfDeploymentTasksResponseDeploy) SetName(v string) *GetListOfDeploymentTasksResponseDeploy {
  s.Name = &v
  return s
}

func (s *GetListOfDeploymentTasksResponseDeploy) SetSubmissionTime(v string) *GetListOfDeploymentTasksResponseDeploy {
  s.SubmissionTime = &v
  return s
}

func (s *GetListOfDeploymentTasksResponseDeploy) SetFinishTime(v string) *GetListOfDeploymentTasksResponseDeploy {
  s.FinishTime = &v
  return s
}

func (s *GetListOfDeploymentTasksResponseDeploy) SetStatus(v string) *GetListOfDeploymentTasksResponseDeploy {
  s.Status = &v
  return s
}

func (s *GetListOfDeploymentTasksResponseDeploy) SetTarget(v string) *GetListOfDeploymentTasksResponseDeploy {
  s.Target = &v
  return s
}

func (s *GetListOfDeploymentTasksResponseDeploy) SetApiRequestId(v string) *GetListOfDeploymentTasksResponseDeploy {
  s.ApiRequestId = &v
  return s
}

type GetListOfDeploymentTasksResponseHeader struct {
}

func (s GetListOfDeploymentTasksResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetListOfDeploymentTasksResponseHeader) GoString() string {
  return s.String()
}




