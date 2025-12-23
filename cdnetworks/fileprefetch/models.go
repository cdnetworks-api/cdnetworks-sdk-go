package fileprefetch

// This file is auto-generated, don't edit it. Thanks.
import (
  "github.com/alibabacloud-go/tea/tea"
)

type GetPrefetchStatusRequest struct {
  // {"en":"Query the start time of the task creation time, such as 2017-01-10 23:33:26. It is not allowed to query tasks before 7 days ago.","zh_CN":"查询的任务创建开始时间，如 2017-01-10 06:33:26，不允许查询3天之前的任务"}
  StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
  // {"en":"Query the end time of the task creation time, such as 2017-01-10 23:33:26,. The query time is no more than 1 days from the start time.","zh_CN":"查询的任务创建结束时间，如 2017-01-10 23:33:26\n1）查询跨度不能超过1天；"}
  EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
  // {"en":"A unique identifier for the same batch of tasks. If you submit multiple urls from an API request, the ID is a unique number for these tasks.\nQuery tasks by batch, such as submitting 10 url refreshes at a time. After the submission is successful, the content management system will return an itemId in the response message.","zh_CN":"表示任务单次提交多个url时任务的唯一标识。\n按批次查询任务，如单次提交10条url文件预取，提交成功后内容管理系统将返回一个itemId在响应消息里。\nitemId 和 查询起止时间不能同时为空。"}
  ItemId *string `json:"itemId,omitempty" xml:"itemId,omitempty"`
  // {"en":"It is the url that you want to prefetch. This element only allows one url to be submitted per query.","zh_CN":"需要预取的文件完整访问路径（url），单次查询只允许输入一条url"}
  Url *string `json:"url,omitempty" xml:"url,omitempty"`
  // {"en":"Task status. The system allows you to select a task status query. These states can be queried:\n1) success\n2) failure\n3) wait\n4) run","zh_CN":"任务状态，允许执行任务状态过滤查询结果，支持查询的状态有：\n1) success\n2) failure\n3) wait\n4) run","exampleValue":"success,failure,wait,run"}
  Status *string `json:"status,omitempty" xml:"status,omitempty"`
  // {"defaultValue":"1","en":"Request page number. The default is 1.","zh_CN":"请求页数，缺省情况下，默认为1"}
  PageNo *string `json:"pageNo,omitempty" xml:"pageNo,omitempty"`
  // {"defaultValue":"20","en":"The number of pages displayed. The default is 20.","zh_CN":"每页显示的条数，缺省情况下，默认值为20"}
  PageSize *string `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s GetPrefetchStatusRequest) String() string {
  return tea.Prettify(s)
}

func (s GetPrefetchStatusRequest) GoString() string {
  return s.String()
}

func (s *GetPrefetchStatusRequest) SetStartTime(v string) *GetPrefetchStatusRequest {
  s.StartTime = &v
  return s
}

func (s *GetPrefetchStatusRequest) SetEndTime(v string) *GetPrefetchStatusRequest {
  s.EndTime = &v
  return s
}

func (s *GetPrefetchStatusRequest) SetItemId(v string) *GetPrefetchStatusRequest {
  s.ItemId = &v
  return s
}

func (s *GetPrefetchStatusRequest) SetUrl(v string) *GetPrefetchStatusRequest {
  s.Url = &v
  return s
}

func (s *GetPrefetchStatusRequest) SetStatus(v string) *GetPrefetchStatusRequest {
  s.Status = &v
  return s
}

func (s *GetPrefetchStatusRequest) SetPageNo(v string) *GetPrefetchStatusRequest {
  s.PageNo = &v
  return s
}

func (s *GetPrefetchStatusRequest) SetPageSize(v string) *GetPrefetchStatusRequest {
  s.PageSize = &v
  return s
}

type GetPrefetchStatusRequestHeader struct {
}

func (s GetPrefetchStatusRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetPrefetchStatusRequestHeader) GoString() string {
  return s.String()
}

type GetPrefetchStatusPaths struct {
}

func (s GetPrefetchStatusPaths) String() string {
  return tea.Prettify(s)
}

func (s GetPrefetchStatusPaths) GoString() string {
  return s.String()
}

type GetPrefetchStatusParameters struct {
}

func (s GetPrefetchStatusParameters) String() string {
  return tea.Prettify(s)
}

func (s GetPrefetchStatusParameters) GoString() string {
  return s.String()
}

type GetPrefetchStatusResponse struct {
  // {"en":"The number of tasks that match the query criteria. If 10 tasks meet the query criteria, the value of count is 10.","zh_CN":"表示本次查询结果的数量，如有10个任务符合查询条件，则count的值为10"}
  Count *int `json:"count,omitempty" xml:"count,omitempty" require:"true"`
  // {"en":"The status code of the task creation result:\n1) 1 means success\n2) 0 means failure.","zh_CN":"任务提交后，系统的响应码:\n1) 0表示失败\n2) 1表示成功","exampleValue":"0,1"}
  Code *int `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
  // {"en":"Content system response message after submitting the task.","zh_CN":"表示任务提交后，系统的响应消息"}
  Message *string `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
  // {"en":"The total number of pages for task query results.","zh_CN":"任务查询结果的总页数"}
  PageNo *int `json:"pageNo,omitempty" xml:"pageNo,omitempty" require:"true"`
  // {"en":"How many purge task data is displayed per page.","zh_CN":"每页显示多少条预取文件的任务数据"}
  PageSize *int `json:"pageSize,omitempty" xml:"pageSize,omitempty" require:"true"`
  // {"en":"Collection of task results.","zh_CN":"任务结果的集合"}
  ResultDetail []*GetPrefetchStatusResponseResultDetail `json:"resultDetail,omitempty" xml:"resultDetail,omitempty" require:"true" type:"Repeated"`
}

func (s GetPrefetchStatusResponse) String() string {
  return tea.Prettify(s)
}

func (s GetPrefetchStatusResponse) GoString() string {
  return s.String()
}

func (s *GetPrefetchStatusResponse) SetCount(v int) *GetPrefetchStatusResponse {
  s.Count = &v
  return s
}

func (s *GetPrefetchStatusResponse) SetCode(v int) *GetPrefetchStatusResponse {
  s.Code = &v
  return s
}

func (s *GetPrefetchStatusResponse) SetMessage(v string) *GetPrefetchStatusResponse {
  s.Message = &v
  return s
}

func (s *GetPrefetchStatusResponse) SetPageNo(v int) *GetPrefetchStatusResponse {
  s.PageNo = &v
  return s
}

func (s *GetPrefetchStatusResponse) SetPageSize(v int) *GetPrefetchStatusResponse {
  s.PageSize = &v
  return s
}

func (s *GetPrefetchStatusResponse) SetResultDetail(v []*GetPrefetchStatusResponseResultDetail) *GetPrefetchStatusResponse {
  s.ResultDetail = v
  return s
}

type GetPrefetchStatusResponseResultDetail struct     {
  // {"en":"The time at which the content management system begins to get the file.","zh_CN":"内容管理系统开始同步文件的时间"}
  BeginTime *string `json:"beginTime,omitempty" xml:"beginTime,omitempty" require:"true"`
  // {"en":"The time at which the content management system receive the request and creates a prefetch task.","zh_CN":"内容管理系统接收预取任务成功并创建预取任务的时间"}
  CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty" require:"true"`
  // {"en":"The time at which cdn cache the file and the content management system completes the summary task results.","zh_CN":"内容管理系统执行预取完成的时间"}
  FinishTime *string `json:"finishTime,omitempty" xml:"finishTime,omitempty" require:"true"`
  // {"en":"The content management system handles the prefetch tasks's success rate. If the success rate is 98%, the value is 98.","zh_CN":"执行文件预取任务的成功率，如98%，则值为98"}
  Rate *string `json:"rate,omitempty" xml:"rate,omitempty" require:"true"`
  // {"en":"The status of the prefetch task. There are several states:\n1) success: Prefetch success.\n2) failure: Prefetch failed.\n3) wait: The prefetch task is waiting to be processed.\n4) run: The prefetch task is being executed.","zh_CN":"预取文件的任务状态：\n1) success：表示文件预取的任务执行成功\n2) failure：表示文件预取的任务执行失败\n3) wait：表示文件预取的任务正在排队中\n4) run：表示文件预取的任务正在执行中","exampleValue":"success,failure,wait,run"}
  Status *string `json:"status,omitempty" xml:"status,omitempty" require:"true"`
  // {"en":"Prefetched file URL.","zh_CN":"要预取的文件url"}
  Url *string `json:"url,omitempty" xml:"url,omitempty" require:"true"`
}

func (s GetPrefetchStatusResponseResultDetail) String() string {
  return tea.Prettify(s)
}

func (s GetPrefetchStatusResponseResultDetail) GoString() string {
  return s.String()
}

func (s *GetPrefetchStatusResponseResultDetail) SetBeginTime(v string) *GetPrefetchStatusResponseResultDetail {
  s.BeginTime = &v
  return s
}

func (s *GetPrefetchStatusResponseResultDetail) SetCreateTime(v string) *GetPrefetchStatusResponseResultDetail {
  s.CreateTime = &v
  return s
}

func (s *GetPrefetchStatusResponseResultDetail) SetFinishTime(v string) *GetPrefetchStatusResponseResultDetail {
  s.FinishTime = &v
  return s
}

func (s *GetPrefetchStatusResponseResultDetail) SetRate(v string) *GetPrefetchStatusResponseResultDetail {
  s.Rate = &v
  return s
}

func (s *GetPrefetchStatusResponseResultDetail) SetStatus(v string) *GetPrefetchStatusResponseResultDetail {
  s.Status = &v
  return s
}

func (s *GetPrefetchStatusResponseResultDetail) SetUrl(v string) *GetPrefetchStatusResponseResultDetail {
  s.Url = &v
  return s
}

type GetPrefetchStatusResponseHeader struct {
}

func (s GetPrefetchStatusResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetPrefetchStatusResponseHeader) GoString() string {
  return s.String()
}




type PrefetchRequest struct {
  // {"en":"To submit URLs for prefetch (pre-loading) onto CDN edge nodes, the following requirements must be met:\n\nThe URL must begin with http:// or https://.\n\nExample: http://www.a.com/image/test.png.\n\nEach individual URL has a maximum length of 2,000 characters.\n\nThe domain specified in the URL must be an active domain currently under our CDN acceleration service and authorized for prefetch.\n\nIf the URL contains non-ASCII characters (e.g., Chinese characters), the submitted URL must be percent-encoded (URL-escaped) using UTF-8 encoding.\n\nDaily limits (per account, adjustable via technical support):\n\n- Maximum 20,000 URLs.\n- Maximum 200 GB of total file size.\n\nThe total number of URLs submitted per single API call must not exceed 400.","zh_CN":"要预取到CDN节点的url集合，url格式说明：\n1、URL 必须以 http:// 或 https:// 开头，输入示例：http://www.a.com/image/test.png。\n2、每个url最大长度 2000 字符。\n3、每个url所在的域名必须是在我司加速的域名且有预取权限。\n4、url中如果包含中文字符，则提交的url需要是中文转义后的url，采用utf-8方式转义。\n5、每日不超过20000条，不超过200G文件大小（账号粒度可调，联系技术支持人员调整）。\n6、每次接口调用url的总数不超过400条。"}
  Urls []*string `json:"urls,omitempty" xml:"urls,omitempty" require:"true" type:"Repeated"`
  // {"defaultValue":"0","en":"Whether to prefetch a range segment:\nThe default value is 0, which means the entire file will be prefetched.\nA value of 1 means only a specific range from the beginning of the file (the file header) is prefetched.\nFor example, if a file is 200 MB in size, only the bytes from 0 up to the configured limit are prefetched-instead of downloading the entire file.\nWhen this option is set to 1, the default prefetch range is 0 to 512 KB.\nEach account can have its own custom prefetch range size. If you need to adjust this setting, please contact us.","zh_CN":"是否预取文件的部分范围（Range）：\n\n默认值为 0，表示预取整个文件。\n设置为 1 时，仅预取文件开头的一段数据（即文件头部分）。\n\n例如，对于一个 200 MB 的文件，系统将只预取从第 0 字节开始、到配置上限为止的数据，而不会下载整个文件。\n\n当该选项设为 1 时，默认预取范围为 0 到 512 KB。\n\n每个账号均可单独配置自己的预取范围大小。如需调整此设置，请联系我们。","exampleValue":"0,1"}
  IsRange *int `json:"isRange,omitempty" xml:"isRange,omitempty"`
}

func (s PrefetchRequest) String() string {
  return tea.Prettify(s)
}

func (s PrefetchRequest) GoString() string {
  return s.String()
}

func (s *PrefetchRequest) SetUrls(v []*string) *PrefetchRequest {
  s.Urls = v
  return s
}

func (s *PrefetchRequest) SetIsRange(v int) *PrefetchRequest {
  s.IsRange = &v
  return s
}

type PrefetchRequestHeader struct {
}

func (s PrefetchRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s PrefetchRequestHeader) GoString() string {
  return s.String()
}

type PrefetchPaths struct {
}

func (s PrefetchPaths) String() string {
  return tea.Prettify(s)
}

func (s PrefetchPaths) GoString() string {
  return s.String()
}

type PrefetchParameters struct {
}

func (s PrefetchParameters) String() string {
  return tea.Prettify(s)
}

func (s PrefetchParameters) GoString() string {
  return s.String()
}

type PrefetchResponse struct {
  // {"en":"The status code of the task creation result:\n1) 1 means success,\n2) 0 means failure.","zh_CN":"表示任务创建结果的状态码:\n1) 1表示任务提交成功\n2) 0表示任务提交失败","exampleValue":"0,1"}
  Code *int `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
  // {"en":"Content system response message after submitting the task.","zh_CN":"表示任务提交后，系统的响应消息"}
  Message *string `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
  // {"en":"After calling the API once and submitting the task successfully, the content system will return an itemId. This ID is the unique identifier for each submission. You can use itemId to batch query the status (success/failure) of the task.","zh_CN":"调用一次接口并提交任务成功后，将返回一个itemId，是当次提交任务的唯一标识，通过itemId可批量查询任务的状态（成功/失败）。"}
  ItemId *string `json:"itemId,omitempty" xml:"itemId,omitempty" require:"true"`
}

func (s PrefetchResponse) String() string {
  return tea.Prettify(s)
}

func (s PrefetchResponse) GoString() string {
  return s.String()
}

func (s *PrefetchResponse) SetCode(v int) *PrefetchResponse {
  s.Code = &v
  return s
}

func (s *PrefetchResponse) SetMessage(v string) *PrefetchResponse {
  s.Message = &v
  return s
}

func (s *PrefetchResponse) SetItemId(v string) *PrefetchResponse {
  s.ItemId = &v
  return s
}

type PrefetchResponseHeader struct {
}

func (s PrefetchResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s PrefetchResponseHeader) GoString() string {
  return s.String()
}




