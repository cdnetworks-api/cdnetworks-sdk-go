package fileprefetch

// This file is auto-generated, don't edit it. Thanks.
import (
  "github.com/alibabacloud-go/tea/tea"
)

type QueryPrefetchStatusRequest struct {
  // {'en':'Query the start time of the task creation time, such as 2017-01-10 23:33:26. It is not allowed to query tasks before 7 days ago.', 'zh_CN':'查询的任务创建开始时间，如 2017-01-10 06:33:26，不允许查询3天之前的任务'}
  StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
  // {'en':'Query the end time of the task creation time, such as 2017-01-10 23:33:26,. The query time is no more than 1 days from the start time.', 'zh_CN':'查询的任务创建结束时间，如 2017-01-10 23:33:26
  // 1）查询跨度不能超过1天；'}
  EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
  // {'en':'A unique identifier for the same batch of tasks. If you submit multiple urls from an API request, the ID is a unique number for these tasks.
  // Query tasks by batch, such as submitting 10 url refreshes at a time. After the submission is successful, the content management system will return an itemId in the response message.', 'zh_CN':'表示任务单次提交多个url时任务的唯一标识。
  // 按批次查询任务，如单次提交10条url文件预取，提交成功后内容管理系统将返回一个itemId在响应消息里。
  // itemId 和 查询起止时间不能同时为空。'}
  ItemId *string `json:"itemId,omitempty" xml:"itemId,omitempty"`
  // {'en':'It is the url that you want to prefetch. This element only allows one url to be submitted per query.', 'zh_CN':'需要预取的文件完整访问路径（url），单次查询只允许输入一条url'}
  Url *string `json:"url,omitempty" xml:"url,omitempty"`
  // {'en':'Task status. The system allows you to select a task status query. These states can be queried:
  // 1.success
  // 2.failure', 'zh_CN':'任务状态，允许执行任务状态过滤查询结果，支持查询的状态有：
  // 1、success
  // 2、failure'}
  Status *string `json:"status,omitempty" xml:"status,omitempty"`
  // {'en':'Request page number. The default is 1.', 'zh_CN':'请求页数，缺省情况下，默认为1'}
  PageNo *string `json:"pageNo,omitempty" xml:"pageNo,omitempty"`
  // {'en':'The number of pages displayed. The default is 20.', 'zh_CN':'每页显示的条数，缺省情况下，默认值为20'}
  PageSize *string `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s QueryPrefetchStatusRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryPrefetchStatusRequest) GoString() string {
  return s.String()
}

func (s *QueryPrefetchStatusRequest) SetStartTime(v string) *QueryPrefetchStatusRequest {
  s.StartTime = &v
  return s
}

func (s *QueryPrefetchStatusRequest) SetEndTime(v string) *QueryPrefetchStatusRequest {
  s.EndTime = &v
  return s
}

func (s *QueryPrefetchStatusRequest) SetItemId(v string) *QueryPrefetchStatusRequest {
  s.ItemId = &v
  return s
}

func (s *QueryPrefetchStatusRequest) SetUrl(v string) *QueryPrefetchStatusRequest {
  s.Url = &v
  return s
}

func (s *QueryPrefetchStatusRequest) SetStatus(v string) *QueryPrefetchStatusRequest {
  s.Status = &v
  return s
}

func (s *QueryPrefetchStatusRequest) SetPageNo(v string) *QueryPrefetchStatusRequest {
  s.PageNo = &v
  return s
}

func (s *QueryPrefetchStatusRequest) SetPageSize(v string) *QueryPrefetchStatusRequest {
  s.PageSize = &v
  return s
}

type QueryPrefetchStatusResponse struct {
  // {'en':'The number of tasks that match the query criteria. If 10 tasks meet the query criteria, the value of count is 10.', 'zh_CN':'表示本次查询结果的数量，如有10个任务符合查询条件，则count的值为10'}
  Count *int32 `json:"count,omitempty" xml:"count,omitempty" require:"true"`
  // {'en':'The status code of the task creation result. 1 means success, 0 means failure.', 'zh_CN':'任务提交后，系统的响应码，0表示失败，1表示成功'}
  Code *int32 `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
  // {'en':'Content system response message after submitting the task.', 'zh_CN':'表示任务提交后，系统的响应消息'}
  Message *string `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
  // {'en':'The total number of pages for task query results.', 'zh_CN':'任务查询结果的总页数'}
  PageNo *int32 `json:"pageNo,omitempty" xml:"pageNo,omitempty" require:"true"`
  // {'en':'How many purge task data is displayed per page.', 'zh_CN':'每页显示多少条预取文件的任务数据'}
  PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty" require:"true"`
  // {'en':'Collection of task results.', 'zh_CN':'任务结果的集合'}
  ResultDetail []*QueryPrefetchStatusResponseResultDetail `json:"resultDetail,omitempty" xml:"resultDetail,omitempty" require:"true" type:"Repeated"`
}

func (s QueryPrefetchStatusResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryPrefetchStatusResponse) GoString() string {
  return s.String()
}

func (s *QueryPrefetchStatusResponse) SetCount(v int32) *QueryPrefetchStatusResponse {
  s.Count = &v
  return s
}

func (s *QueryPrefetchStatusResponse) SetCode(v int32) *QueryPrefetchStatusResponse {
  s.Code = &v
  return s
}

func (s *QueryPrefetchStatusResponse) SetMessage(v string) *QueryPrefetchStatusResponse {
  s.Message = &v
  return s
}

func (s *QueryPrefetchStatusResponse) SetPageNo(v int32) *QueryPrefetchStatusResponse {
  s.PageNo = &v
  return s
}

func (s *QueryPrefetchStatusResponse) SetPageSize(v int32) *QueryPrefetchStatusResponse {
  s.PageSize = &v
  return s
}

func (s *QueryPrefetchStatusResponse) SetResultDetail(v []*QueryPrefetchStatusResponseResultDetail) *QueryPrefetchStatusResponse {
  s.ResultDetail = v
  return s
}

type QueryPrefetchStatusResponseResultDetail struct     {
  // {'en':'The time at which the content management system begins to get the file.', 'zh_CN':'内容管理系统开始同步文件的时间'}
  BeginTime *string `json:"beginTime,omitempty" xml:"beginTime,omitempty" require:"true"`
  // {'en':'The time at which the content management system receive the request and creates a prefetch task.', 'zh_CN':'内容管理系统接收预取任务成功并创建预取任务的时间'}
  CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty" require:"true"`
  // {'en':'The time at which cdn cache the file and the content management system completes the summary task results.', 'zh_CN':'内容管理系统执行预取完成的时间'}
  FinishTime *string `json:"finishTime,omitempty" xml:"finishTime,omitempty" require:"true"`
  // {"en":"The content management system handles the prefetch tasks's success rate. If the success rate is 98%, the value is 98.", "zh_CN":"执行文件预取任务的成功率，如98%，则值为98"}
  Rate *string `json:"rate,omitempty" xml:"rate,omitempty" require:"true"`
  // {'en':'The status of the prefetch task. There are several states:
  // Success: Prefetch success.
  // Failure: Prefetch failed.
  // Wait: The prefetch task is waiting to be processed.
  // Run: The prefetch task is being executed.', 'zh_CN':'预取文件的任务状态：
  // success：表示文件预取的任务执行成功
  // failure：表示文件预取的任务执行失败
  // wait：表示文件预取的任务正在排队中
  // run：表示文件预取的任务正在执行中'}
  Status *string `json:"status,omitempty" xml:"status,omitempty" require:"true"`
  // {'en':'Prefetched file URL.', 'zh_CN':'要预取的文件url'}
  Url *string `json:"url,omitempty" xml:"url,omitempty" require:"true"`
}

func (s QueryPrefetchStatusResponseResultDetail) String() string {
  return tea.Prettify(s)
}

func (s QueryPrefetchStatusResponseResultDetail) GoString() string {
  return s.String()
}

func (s *QueryPrefetchStatusResponseResultDetail) SetBeginTime(v string) *QueryPrefetchStatusResponseResultDetail {
  s.BeginTime = &v
  return s
}

func (s *QueryPrefetchStatusResponseResultDetail) SetCreateTime(v string) *QueryPrefetchStatusResponseResultDetail {
  s.CreateTime = &v
  return s
}

func (s *QueryPrefetchStatusResponseResultDetail) SetFinishTime(v string) *QueryPrefetchStatusResponseResultDetail {
  s.FinishTime = &v
  return s
}

func (s *QueryPrefetchStatusResponseResultDetail) SetRate(v string) *QueryPrefetchStatusResponseResultDetail {
  s.Rate = &v
  return s
}

func (s *QueryPrefetchStatusResponseResultDetail) SetStatus(v string) *QueryPrefetchStatusResponseResultDetail {
  s.Status = &v
  return s
}

func (s *QueryPrefetchStatusResponseResultDetail) SetUrl(v string) *QueryPrefetchStatusResponseResultDetail {
  s.Url = &v
  return s
}

type QueryPrefetchStatusPaths struct {
}

func (s QueryPrefetchStatusPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryPrefetchStatusPaths) GoString() string {
  return s.String()
}

type QueryPrefetchStatusParameters struct {
}

func (s QueryPrefetchStatusParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryPrefetchStatusParameters) GoString() string {
  return s.String()
}

type QueryPrefetchStatusRequestHeader struct {
}

func (s QueryPrefetchStatusRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryPrefetchStatusRequestHeader) GoString() string {
  return s.String()
}

type QueryPrefetchStatusResponseHeader struct {
}

func (s QueryPrefetchStatusResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryPrefetchStatusResponseHeader) GoString() string {
  return s.String()
}




type PrefetchRequest struct {
  // {'en':'If you need to prefetch several cached URLs. The submitted URL should meet the following format requirements:
  // 1.The URL must start with http:// or https://, url example: http://www.a.com/image/test.png.
  // 2.Each url has a maximum length of 2000 characters.
  // 3.The domain in the URL is must be the domain of the CDN service.
  // 4.If the url contains special characters such as Chinese characters and spaces, our system will generate multiple push tasks. In addition to pushing the original URL, these special characters will be converted int32o ASCII codes and pushed. If you only want to clean up the transcoded URL, you need to use UTF-8 to complete the transcoding before submitting the URL, and then submit the escaped url to our system.
  // 5.No more than 20000 urls per day, and no more than 200G file size (it can be adjusted according to account, contact your technical support).
  // 6.The total number of URLs called by each interface shall not exceed 400'
  // , 'zh_CN':'要预取到CDN节点的url集合，url格式说明：
  // 1、URL 必须以 http:// 或 https:// 开头，输入示例：http://www.a.com/image/test.png。
  // 2、每个url最大长度 2000 字符。
  // 3、每个url所在的域名必须是在我司加速的域名且有预取权限。
  // 4、url中如果包含中文字符，则提交的url需要是中文转义后的url，采用utf-8方式转义。
  // 5、每日不超过20000条，不超过200G文件大小（账号粒度可调，联系技术支持人员调整）。
  // 6、每次接口调用url的总数不超过400条。'}
  Urls []*string `json:"urls,omitempty" xml:"urls,omitempty" require:"true" type:"Repeated"`
  // {"en":"Only prefetch a range segment of the file header. The user get the file from the beginning, and they will select quickly their int32erested. If the file header is cached, the first pack time of the user's http request will be short.This feature allows users to filter content faster. For example, if a file has 200MB, only the size of the file 0~range is prefetched, instead of prefetching the entire file. Each account can be configured with a size of the range. If you need to modify the size, please contact us. If this element is assigned a value of 1, the default prefetch is 0~512KB.", "zh_CN":"是否需要预取range段。
  // 
  // 1、默认为0，表示预取完整的文件；
  // 2、1表示预取文件0~512KB的range段（账号粒度可调，联系技术支持人员调整）。"}
  IsRange *int32 `json:"isRange,omitempty" xml:"isRange,omitempty"`
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

func (s *PrefetchRequest) SetIsRange(v int32) *PrefetchRequest {
  s.IsRange = &v
  return s
}

type PrefetchResponse struct {
  // {'en':'The status code of the task creation result, 1 means success, 0 means failure.', 'zh_CN':'表示任务创建结果的状态码，1表示任务提交成功，0表示任务提交失败'}
  Code *int32 `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
  // {'en':'Content system response message after submitting the task.', 'zh_CN':'表示任务提交后，系统的响应消息'}
  Message *string `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
  // {'en':'After calling the API once and submitting the task successfully, the content system will return an itemId. This ID is the unique identifier for each submission. You can use itemId to batch query the status (success/failure) of the task.', 'zh_CN':'调用一次接口并提交任务成功后，将返回一个itemId，是当次提交任务的唯一标识，通过itemId可批量查询任务的状态（成功/失败）。'}
  ItemId *string `json:"itemId,omitempty" xml:"itemId,omitempty" require:"true"`
}

func (s PrefetchResponse) String() string {
  return tea.Prettify(s)
}

func (s PrefetchResponse) GoString() string {
  return s.String()
}

func (s *PrefetchResponse) SetCode(v int32) *PrefetchResponse {
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

type PrefetchRequestHeader struct {
}

func (s PrefetchRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s PrefetchRequestHeader) GoString() string {
  return s.String()
}

type PrefetchResponseHeader struct {
}

func (s PrefetchResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s PrefetchResponseHeader) GoString() string {
  return s.String()
}




