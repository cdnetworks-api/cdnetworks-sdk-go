package ngcontentmanagement

// This file is auto-generated, don't edit it. Thanks.
import (
  "github.com/alibabacloud-go/tea/tea"
)

type GetPurgeRequestStatusPaths struct {
  // {"en" : "ID of a purge request.", "zh_CN": "刷新任务的ID。"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
}

func (s GetPurgeRequestStatusPaths) String() string {
  return tea.Prettify(s)
}

func (s GetPurgeRequestStatusPaths) GoString() string {
  return s.String()
}

func (s *GetPurgeRequestStatusPaths) SetId(v string) *GetPurgeRequestStatusPaths {
  s.Id = &v
  return s
}

type GetPurgeRequestStatusParameters struct {
}

func (s GetPurgeRequestStatusParameters) String() string {
  return tea.Prettify(s)
}

func (s GetPurgeRequestStatusParameters) GoString() string {
  return s.String()
}

type GetPurgeRequestStatusRequestHeader struct {
}

func (s GetPurgeRequestStatusRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetPurgeRequestStatusRequestHeader) GoString() string {
  return s.String()
}

type GetPurgeRequestStatusRequest struct {
}

func (s GetPurgeRequestStatusRequest) String() string {
  return tea.Prettify(s)
}

func (s GetPurgeRequestStatusRequest) GoString() string {
  return s.String()
}

type GetPurgeRequestStatusResponseHeader struct {
}

func (s GetPurgeRequestStatusResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetPurgeRequestStatusResponseHeader) GoString() string {
  return s.String()
}

type GetPurgeRequestStatusResponse struct {
  // {"en" : "Name of the purge request.", "zh_CN": "刷新请求的名称。"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en" : "The files that were purged.", "zh_CN": "被刷新的文件。"}
  FileUrls []*string `json:"fileUrls,omitempty" xml:"fileUrls,omitempty" require:"true" type:"Repeated"`
  // {"en" : "If a file's cache key depends on request headers, you can specify the header values that are applicable to purge one version of the cached file. The same set of header values will apply to all entries in fileUrls.  ", "zh_CN": "如果文件的缓存键与请求头相关，则可以指定请求头和值来刷新相应的缓存文件。此处指定的请求头和值将应用于fileUrls中的所有条目。"}
  FileHeaders []*GetPurgeRequestStatusResponseFileHeaders `json:"fileHeaders,omitempty" xml:"fileHeaders,omitempty" require:"true" type:"Repeated"`
  // {"en" : "<= 20 items 
  // The directories that were purged.", "zh_CN": "<= 20 条目 
  // 被刷新的目录。"}
  DirUrls []*string `json:"dirUrls,omitempty" xml:"dirUrls,omitempty" require:"true" type:"Repeated"`
  // {"en" : "If a directory's cache key depends on request headers, you can specify the header values that are applicable to purge one version of the cached directory. The same set of header values will apply to all entries in dirUrls.", "zh_CN": "如果目录的缓存键与请求头相关，则可以指定请求头和值来刷新相应的缓存目录。此处指定的请求头和值将应用于dirUrls中的所有条目。"}
  DirHeaders []*GetPurgeRequestStatusResponseDirHeaders `json:"dirHeaders,omitempty" xml:"dirHeaders,omitempty" require:"true" type:"Repeated"`
  // {"en" : "<= 2 items 
  // Regular expression patterns used to match the cache key. Each must begin with the following format: 
  //  {scheme}://{hostname}/. {scheme} can be http, https, or any which matches any scheme.
  // Example: 
  // https://test.domain.com/my.*\.(jpg|png)\?q=
  // <br/>
  // For performance considerations, the following restrictions apply:
  // The regular expression pattern following the hostname can be up to 126 characters.
  // 
  // It can consist of up to two unlimited quantifiers ('*', '+', or ',}').
  // The upper limit on a quantifier cannot be more than 59, for example, {1,59}", "zh_CN": "<= 2 条目 
  // 用于匹配缓存键的正则表达式。每一个表达式必须以以下格式开始:
  // {协议}://{域名}/. {协议}可以是http, https或any（表示不限协议）。
  // 示例:
  // https://test.domain.com/my.*\.(jpg|png)\?q=
  // <br/>
  // 出于性能考虑，使用正则表达式有以下限制：
  // 
  // 在域名后面的正则表达式最多只能包含126个字符。
  // 最多只能包含两个限定符('*'、'+'或',}')。
  // 限定符的上限不能超过59，例如{1,59}"}
  RegexPatterns []*string `json:"regexPatterns,omitempty" xml:"regexPatterns,omitempty" require:"true" type:"Repeated"`
  // {"en" : "Enum: staging production 
  // Indicates whether the purge is in the staging or production environment.", "zh_CN": "取值范围: staging, production 
  // 刷新任务对应的环境，即演练环境或生产环境。"}
  Target *string `json:"target,omitempty" xml:"target,omitempty" require:"true"`
  // {"en" : "An RFC 3339 date indicating when the purge request was created.", "zh_CN": "RFC 3339格式的日期，表示刷新请求的创建时间。"}
  SubmissionTime *string `json:"submissionTime,omitempty" xml:"submissionTime,omitempty" require:"true"`
  // {"en" : "[ 0 .. 100 ] 
  // A percentage that indicates the completion of the purge request.", "zh_CN": "取值范围: [ 0 .. 100 ] 
  // 刷新请求任务的成功率。"}
  SuccessRate *int `json:"successRate,omitempty" xml:"successRate,omitempty" require:"true"`
  // {"en" : "Enum: waiting, inprogress, finished 
  // Indicates the status of the purge request.", "zh_CN": "取值范围: waiting, inprogress, finished 
  // 刷新请求的任务执行状态，包括等待中，执行中，已完成等状态。"}
  Status *string `json:"status,omitempty" xml:"status,omitempty" require:"true"`
  // {"en" : "Enum: delete, invalidate 
  // This controls whether cached files and directories should be removed altogether from the CDN Pro servers (delete) or flagged as invalid (invalidate).", "zh_CN": "取值范围: delete, invalidate 
  // 刷新请求的类型，包括完全删除(delete)和标记为无效(invalidate)。"}
  Action *string `json:"action,omitempty" xml:"action,omitempty" require:"true"`
  // {"en" : "RFC 3339 date indicating when the purge was completed. It can be empty if the purge is still in progress.", "zh_CN": "RFC 3339格式的日期，表示刷新任务完成的时间。如果刷新还在进行中，则返回空值。"}
  FinishTime *string `json:"finishTime,omitempty" xml:"finishTime,omitempty" require:"true"`
  // {"en" : "Indicates the corresponding API request ID.", "zh_CN": "API请求ID。"}
  ApiRequestId *string `json:"apiRequestId,omitempty" xml:"apiRequestId,omitempty" require:"true"`
  // {"en" : "ID of a webhook to call when the purge task completes.", "zh_CN": "刷新任务完成时要调用的webhook的ID。"}
  Webhook *string `json:"webhook,omitempty" xml:"webhook,omitempty" require:"true"`
}

func (s GetPurgeRequestStatusResponse) String() string {
  return tea.Prettify(s)
}

func (s GetPurgeRequestStatusResponse) GoString() string {
  return s.String()
}

func (s *GetPurgeRequestStatusResponse) SetName(v string) *GetPurgeRequestStatusResponse {
  s.Name = &v
  return s
}

func (s *GetPurgeRequestStatusResponse) SetFileUrls(v []*string) *GetPurgeRequestStatusResponse {
  s.FileUrls = v
  return s
}

func (s *GetPurgeRequestStatusResponse) SetFileHeaders(v []*GetPurgeRequestStatusResponseFileHeaders) *GetPurgeRequestStatusResponse {
  s.FileHeaders = v
  return s
}

func (s *GetPurgeRequestStatusResponse) SetDirUrls(v []*string) *GetPurgeRequestStatusResponse {
  s.DirUrls = v
  return s
}

func (s *GetPurgeRequestStatusResponse) SetDirHeaders(v []*GetPurgeRequestStatusResponseDirHeaders) *GetPurgeRequestStatusResponse {
  s.DirHeaders = v
  return s
}

func (s *GetPurgeRequestStatusResponse) SetRegexPatterns(v []*string) *GetPurgeRequestStatusResponse {
  s.RegexPatterns = v
  return s
}

func (s *GetPurgeRequestStatusResponse) SetTarget(v string) *GetPurgeRequestStatusResponse {
  s.Target = &v
  return s
}

func (s *GetPurgeRequestStatusResponse) SetSubmissionTime(v string) *GetPurgeRequestStatusResponse {
  s.SubmissionTime = &v
  return s
}

func (s *GetPurgeRequestStatusResponse) SetSuccessRate(v int) *GetPurgeRequestStatusResponse {
  s.SuccessRate = &v
  return s
}

func (s *GetPurgeRequestStatusResponse) SetStatus(v string) *GetPurgeRequestStatusResponse {
  s.Status = &v
  return s
}

func (s *GetPurgeRequestStatusResponse) SetAction(v string) *GetPurgeRequestStatusResponse {
  s.Action = &v
  return s
}

func (s *GetPurgeRequestStatusResponse) SetFinishTime(v string) *GetPurgeRequestStatusResponse {
  s.FinishTime = &v
  return s
}

func (s *GetPurgeRequestStatusResponse) SetApiRequestId(v string) *GetPurgeRequestStatusResponse {
  s.ApiRequestId = &v
  return s
}

func (s *GetPurgeRequestStatusResponse) SetWebhook(v string) *GetPurgeRequestStatusResponse {
  s.Webhook = &v
  return s
}

type GetPurgeRequestStatusResponseFileHeaders struct     {
  // {"en" : "HTTP header name.", "zh_CN": "HTTP头部名称。"}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // {"en" : "Value of an HTTP header.", "zh_CN": "HTTP头部值。"}
  Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s GetPurgeRequestStatusResponseFileHeaders) String() string {
  return tea.Prettify(s)
}

func (s GetPurgeRequestStatusResponseFileHeaders) GoString() string {
  return s.String()
}

func (s *GetPurgeRequestStatusResponseFileHeaders) SetName(v string) *GetPurgeRequestStatusResponseFileHeaders {
  s.Name = &v
  return s
}

func (s *GetPurgeRequestStatusResponseFileHeaders) SetValue(v string) *GetPurgeRequestStatusResponseFileHeaders {
  s.Value = &v
  return s
}

type GetPurgeRequestStatusResponseDirHeaders struct     {
  // {"en" : "HTTP header name.", "zh_CN": "HTTP头部名称。"}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // {"en" : "alue of an HTTP header.", "zh_CN": "HTTP头部值。"}
  Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s GetPurgeRequestStatusResponseDirHeaders) String() string {
  return tea.Prettify(s)
}

func (s GetPurgeRequestStatusResponseDirHeaders) GoString() string {
  return s.String()
}

func (s *GetPurgeRequestStatusResponseDirHeaders) SetName(v string) *GetPurgeRequestStatusResponseDirHeaders {
  s.Name = &v
  return s
}

func (s *GetPurgeRequestStatusResponseDirHeaders) SetValue(v string) *GetPurgeRequestStatusResponseDirHeaders {
  s.Value = &v
  return s
}




type CreateAPurgeRequestPaths struct {
}

func (s CreateAPurgeRequestPaths) String() string {
  return tea.Prettify(s)
}

func (s CreateAPurgeRequestPaths) GoString() string {
  return s.String()
}

type CreateAPurgeRequestParameters struct {
}

func (s CreateAPurgeRequestParameters) String() string {
  return tea.Prettify(s)
}

func (s CreateAPurgeRequestParameters) GoString() string {
  return s.String()
}

type CreateAPurgeRequestRequestHeader struct {
}

func (s CreateAPurgeRequestRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s CreateAPurgeRequestRequestHeader) GoString() string {
  return s.String()
}

type CreateAPurgeRequestRequest struct {
  // {"en" : "A description of the purge request.", "zh_CN": "刷新请求的说明。"}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // {"en" : "URLs of files to purge.  File URLs should not contain the asterisk character, '*'.   If a directory or filename in a URL includes a percent character, '%', be sure to encode it. A URL can be up to 2048 characters.", "zh_CN": "要刷新的文件的URL。URL不能包含星号字符'*'。如果URL中的目录或文件名包含'%'等特殊符号，需要先进行URL编码。每个URL长度不能超过2048个字符。"}
  FileUrls []*string `json:"fileUrls,omitempty" xml:"fileUrls,omitempty" type:"Repeated"`
  // {"en" : "If a file's cache key depends on request headers, you can specify the header values that are applicable to purge one version of the cached file. The same set of header values will apply to all entries in fileUrls.", "zh_CN": "如果文件的缓存键与请求头相关，则可以指定请求头和值来刷新相应的缓存文件。此处指定的请求头和值将应用于fileUrls中的所有条目。"}
  FileHeaders []*CreateAPurgeRequestRequestFileHeaders `json:"fileHeaders,omitempty" xml:"fileHeaders,omitempty" type:"Repeated"`
  // {"en" : "<= 20 items 
  // URLs to purge. URLs must begin with http:// or https:// and can be up to 2048 characters. Use the '*' character to purge multiple files or directories. If a URL has multiple sets of asterisk characters, only the last '*' or '**' will be treated as a wildcard. Other instances of '*' earlier in the URL will be treated as the literal character '*'.
  // <table><tr><th>Example</th><th>Description</th></tr><tr><td>http://test.domain2.com/mydir</td><td>Purge all variations of a single directory, but not its subdirectories or files. Variations may exist if custom cache keys are used.</td></tr><tr><td>http://test.domain2.com/mydir/**</td><td>Purge all files and subdirectories whose cache key begins with http://test.domain2.com/mydir/.</td></tr><tr><td>http://test.domain2.com/mydir/*</td><td>Purge all files, but not subdirectories, within a directory.</td></tr><tr><td>http://test.domain2.com/mydir/*.jpg</td><td>Purge all cache entries ending with the .jpg file extension. Subdirectories of http://test.domain2.com/mydir/ are not purged. </td></tr><tr><td>http://test.domain2.com/mydir/a*</td><td>Purge all files, but not subdirectories, that start with the letter 'a'.</td></tr><tr><td>http://test.domain2.com/mydir/a**</td><td>Purge all files and subdirectories that start with the letter 'a'.</td></tr><tr><td>http://test.domain2.com/mydir/a.jpg</td><td>Purge all variations of 'a.jpg'. Variations may exist if custom cache keys are used.</td></tr><tr><td>http://test.domain2.com/my**jpg</td><td>Purge all entries whose cache key begins with http://test.domain2.com/my and ends with the suffix jpg. The '**' can match anything in the path including additional subdirectories. For example, http://test.domain2.com/mydirectory/picture.jpg would be purged.</td></tr></table>
  // If a directory or filename in a URL includes a percent character, '%', be sure to encode it.", "zh_CN": "<= 20 条目 
  // 要刷新的目录的URL。URL必须以http:// 或者 https://开头，每条URL最多只能包含2048个字符。 在URL中使用'*'字符可以匹配多个文件或目录。如果一条URL中带有多组'*'，则只有最后一个'*'或'**'会被当成通配符来进行匹配，其它的'*'只会被当成普通字符。
  // <table><tr><th>示例</th><th>描述</th></tr><tr><td>http://test.domain2.com/mydir</td><td>刷新单个目录的所有变体，但不包括其子目录或文件。当您自定义了缓存键时，则可能存在变体。</td></tr><tr><td>http://test.domain2.com/mydir/**</td><td>刷新缓存键以http://test.domain2.com/mydir/开头的所有文件和子目录。</td></tr><tr><td>http://test.domain2.com/mydir/*</td><td>刷新目录中的所有文件，但不包括子目录。</td></tr><tr><td>http://test.domain2.com/mydir/*.jpg</td><td>刷新所有以.jpg文件扩展名结尾的缓存，但不会刷新http://test.domain2.com/mydir/ 的子目录。 </td></tr><tr><td>http://test.domain2.com/mydir/a*</td><td>刷新以字母'a'开头的所有文件，但不包括子目录。</td></tr><tr><td>http://test.domain2.com/mydir/a**</td><td>刷新以字母'a'开头的所有文件和子目录。</td></tr><tr><td>http://test.domain2.com/mydir/a.jpg</td><td>刷新'a.jpg'文件的所有变体。当您自定义了缓存键时，则可能存在变体。</td></tr><tr><td>http://test.domain2.com/my**jpg</td><td>刷新缓存键以 http://test.domain2.com/my 开头并以后缀 jpg 结尾的所有条目。'**'可以匹配路径中的任何内容，包括其他子目录。例如，http://test.domain2.com/mydirectory/picture.jpg 将被刷新。</td></tr></table>
  // 如果URL中的目录或文件名包含百分号'%'等特殊符号时，请确保先进行URL编码。"}
  DirUrls []*string `json:"dirUrls,omitempty" xml:"dirUrls,omitempty" type:"Repeated"`
  // {"en" : "<= 2 items 
  // Regular expression patterns used to match the cache key. Each must begin with the following format: 
  //  {scheme}://{hostname}/. {scheme} can be http, https, or any, which matches any scheme.
  // Example: 
  // https://test.domain.com/my.*\.(jpg|png)\?q=
  // <br/>
  // For performance considerations, the following restrictions apply:
  // The regular expression pattern following the hostname can be up to 126 characters.
  // 
  // It can consist of up to two unlimited quantifiers ('*', '+', or ',}').
  // The upper limit on a quantifier cannot be more than 59, for example, {1,59}", "zh_CN": "<= 2 条目 
  // 用于匹配缓存键的正则表达式。
  // 每个表达式必须以
  // {协议}://{域名}/ 格式开头。其中，{协议} 可以是 http, https，或any（表示不限协议）。
  // 示例：
  // https://test.domain.com/my.*\.(jpg|png)\?q=
  // <br/>
  // 出于性能考虑，使用正则表达式有以下限制：
  // 
  // 在域名后面的正则表达式最多只能包含126个字符。
  // 最多只能包含两个限定符('*'、'+'或',}')。
  // 限定符的上限不能超过59，例如{1,59}"}
  RegexPatterns []*string `json:"regexPatterns,omitempty" xml:"regexPatterns,omitempty" type:"Repeated"`
  // {"en" : "Enum: delete invalidate 
  // Default: invalidate 
  // This controls whether cached files and directories should be removed altogether from the CDN Pro servers (delete) or flagged as invalid (invalidate).", "zh_CN": "取值范围: delete, invalidate 
  // 默认值: invalidate 
  // 指定刷新类型，包括完全删除(delete)和标记为无效(invalidate)。"}
  Action *string `json:"action,omitempty" xml:"action,omitempty"`
  // {"en" : "Enum: staging production 
  // Specify if the purge request applies to the staging or production environment.", "zh_CN": "取值范围: staging, production 
  // 指定刷新请求应用于演练环境还是生产环境。"}
  Target *string `json:"target,omitempty" xml:"target,omitempty" require:"true"`
  // {"en" : "ID of a webhook to call when the purge task completes.", "zh_CN": "刷新任务完成时要调用的webhook的ID。"}
  Webhook *string `json:"webhook,omitempty" xml:"webhook,omitempty"`
}

func (s CreateAPurgeRequestRequest) String() string {
  return tea.Prettify(s)
}

func (s CreateAPurgeRequestRequest) GoString() string {
  return s.String()
}

func (s *CreateAPurgeRequestRequest) SetName(v string) *CreateAPurgeRequestRequest {
  s.Name = &v
  return s
}

func (s *CreateAPurgeRequestRequest) SetFileUrls(v []*string) *CreateAPurgeRequestRequest {
  s.FileUrls = v
  return s
}

func (s *CreateAPurgeRequestRequest) SetFileHeaders(v []*CreateAPurgeRequestRequestFileHeaders) *CreateAPurgeRequestRequest {
  s.FileHeaders = v
  return s
}

func (s *CreateAPurgeRequestRequest) SetDirUrls(v []*string) *CreateAPurgeRequestRequest {
  s.DirUrls = v
  return s
}

func (s *CreateAPurgeRequestRequest) SetRegexPatterns(v []*string) *CreateAPurgeRequestRequest {
  s.RegexPatterns = v
  return s
}

func (s *CreateAPurgeRequestRequest) SetAction(v string) *CreateAPurgeRequestRequest {
  s.Action = &v
  return s
}

func (s *CreateAPurgeRequestRequest) SetTarget(v string) *CreateAPurgeRequestRequest {
  s.Target = &v
  return s
}

func (s *CreateAPurgeRequestRequest) SetWebhook(v string) *CreateAPurgeRequestRequest {
  s.Webhook = &v
  return s
}

type CreateAPurgeRequestRequestFileHeaders struct     {
  // {"en" : "HTTP header name.", "zh_CN": "HTTP 头部名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // {"en" : "Value of an HTTP header.", "zh_CN": "HTTP 头部的值"}
  Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s CreateAPurgeRequestRequestFileHeaders) String() string {
  return tea.Prettify(s)
}

func (s CreateAPurgeRequestRequestFileHeaders) GoString() string {
  return s.String()
}

func (s *CreateAPurgeRequestRequestFileHeaders) SetName(v string) *CreateAPurgeRequestRequestFileHeaders {
  s.Name = &v
  return s
}

func (s *CreateAPurgeRequestRequestFileHeaders) SetValue(v string) *CreateAPurgeRequestRequestFileHeaders {
  s.Value = &v
  return s
}

type CreateAPurgeRequestResponse struct {
}

func (s CreateAPurgeRequestResponse) String() string {
  return tea.Prettify(s)
}

func (s CreateAPurgeRequestResponse) GoString() string {
  return s.String()
}

type CreateAPurgeRequestResponseHeader struct {
  // {"en":"The Location header is a URL representing the new purge request, for example, <code>Location: https://{domain}/cdn/purges/e91e8674-c2c5-4440-a1de-8b2ea99293dd</code>.", "zh_CN":"通过Location响应头返回新建的刷新任务的URL。URL中包含刷新任务的ID，可使用该ID调用'查询刷新任务详情'接口来查看刷新任务详情。URL示例：<code>Location: https://{domain}/cdn/purges/5dca2205f9e9cc0001df7b33"}
  Location *string `json:"Location,omitempty" xml:"Location,omitempty" require:"true"`
}

func (s CreateAPurgeRequestResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s CreateAPurgeRequestResponseHeader) GoString() string {
  return s.String()
}

func (s *CreateAPurgeRequestResponseHeader) SetLocation(v string) *CreateAPurgeRequestResponseHeader {
  s.Location = &v
  return s
}




type GetPrefetchRequestStatusPaths struct {
  // {"en" : "ID of a prefetch request.", "zh_CN": "预取请求的 id。"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
}

func (s GetPrefetchRequestStatusPaths) String() string {
  return tea.Prettify(s)
}

func (s GetPrefetchRequestStatusPaths) GoString() string {
  return s.String()
}

func (s *GetPrefetchRequestStatusPaths) SetId(v string) *GetPrefetchRequestStatusPaths {
  s.Id = &v
  return s
}

type GetPrefetchRequestStatusParameters struct {
}

func (s GetPrefetchRequestStatusParameters) String() string {
  return tea.Prettify(s)
}

func (s GetPrefetchRequestStatusParameters) GoString() string {
  return s.String()
}

type GetPrefetchRequestStatusRequestHeader struct {
}

func (s GetPrefetchRequestStatusRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetPrefetchRequestStatusRequestHeader) GoString() string {
  return s.String()
}

type GetPrefetchRequestStatusRequest struct {
}

func (s GetPrefetchRequestStatusRequest) String() string {
  return tea.Prettify(s)
}

func (s GetPrefetchRequestStatusRequest) GoString() string {
  return s.String()
}

type GetPrefetchRequestStatusResponseHeader struct {
}

func (s GetPrefetchRequestStatusResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetPrefetchRequestStatusResponseHeader) GoString() string {
  return s.String()
}

type GetPrefetchRequestStatusResponse struct {
  // {"en" : "Range: <= 1000 characters 
  // A short description of the prefetch task.", "zh_CN": "取值范围: <= 1000 字符 
  // 预取请求的简短描述。"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en" : "", "zh_CN": ""}
  FileList []*GetPrefetchRequestStatusResponseFileList `json:"fileList,omitempty" xml:"fileList,omitempty" require:"true" type:"Repeated"`
  // {"en" : "Details about the prefetch request's status.", "zh_CN": "预取请求的状态信息。"}
  Metadata *GetPrefetchRequestStatusResponseMetadata `json:"metadata,omitempty" xml:"metadata,omitempty" require:"true" type:"Struct"`
  // {"en" : "A list of continents representing the regions in which to perform the prefetch. Omitting the field means the prefetch will be done by all regions' servers.", "zh_CN": "需要预取内容的大洲，以大洲英文名表示。未指定时表示预取内容到所有大洲的服务器。"}
  Regions []*string `json:"regions,omitempty" xml:"regions,omitempty" require:"true" type:"Repeated"`
  // {"en" : "RFC 3339 date indicating when the prefetch should begin. This must be in UTC time, for example, '2021-03-06T00:00:00Z'.", "zh_CN": "RFC 3339格式的日期，表示开始预取的时间。必须使用UTC时间，例如'2021-03-06T00:00:00Z'。"}
  StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty" require:"true"`
  // {"en" : "ID of a webhook to call when the prefetch task completes.", "zh_CN": "预取任务完成时调用的webhook ID。"}
  Webhook *string `json:"webhook,omitempty" xml:"webhook,omitempty" require:"true"`
}

func (s GetPrefetchRequestStatusResponse) String() string {
  return tea.Prettify(s)
}

func (s GetPrefetchRequestStatusResponse) GoString() string {
  return s.String()
}

func (s *GetPrefetchRequestStatusResponse) SetName(v string) *GetPrefetchRequestStatusResponse {
  s.Name = &v
  return s
}

func (s *GetPrefetchRequestStatusResponse) SetFileList(v []*GetPrefetchRequestStatusResponseFileList) *GetPrefetchRequestStatusResponse {
  s.FileList = v
  return s
}

func (s *GetPrefetchRequestStatusResponse) SetMetadata(v *GetPrefetchRequestStatusResponseMetadata) *GetPrefetchRequestStatusResponse {
  s.Metadata = v
  return s
}

func (s *GetPrefetchRequestStatusResponse) SetRegions(v []*string) *GetPrefetchRequestStatusResponse {
  s.Regions = v
  return s
}

func (s *GetPrefetchRequestStatusResponse) SetStartTime(v string) *GetPrefetchRequestStatusResponse {
  s.StartTime = &v
  return s
}

func (s *GetPrefetchRequestStatusResponse) SetWebhook(v string) *GetPrefetchRequestStatusResponse {
  s.Webhook = &v
  return s
}

type GetPrefetchRequestStatusResponseFileList struct     {
  // {"en" : "Range: [ 10 .. 2048 ] characters 
  // A URL to prefetch. It must begin with 'http' or 'https' and can be up to 2048 characters.", "zh_CN": "取值范围: [ 10 .. 2048 ] 字符 
  // 预取的URL。必须以'http'或'https'开头，长度不超过2048个字符。"}
  Url *string `json:"url,omitempty" xml:"url,omitempty"`
  // {"en" : "If a URL's cache key depends on request headers, you can specify the header values that are applicable to prefetch one version of the URL.", "zh_CN": "如果需要在缓存键中加入请求头，可用该字段指定请求头。"}
  Headers []*GetPrefetchRequestStatusResponseFileListHeaders `json:"headers,omitempty" xml:"headers,omitempty" type:"Repeated"`
}

func (s GetPrefetchRequestStatusResponseFileList) String() string {
  return tea.Prettify(s)
}

func (s GetPrefetchRequestStatusResponseFileList) GoString() string {
  return s.String()
}

func (s *GetPrefetchRequestStatusResponseFileList) SetUrl(v string) *GetPrefetchRequestStatusResponseFileList {
  s.Url = &v
  return s
}

func (s *GetPrefetchRequestStatusResponseFileList) SetHeaders(v []*GetPrefetchRequestStatusResponseFileListHeaders) *GetPrefetchRequestStatusResponseFileList {
  s.Headers = v
  return s
}

type GetPrefetchRequestStatusResponseFileListHeaders struct     {
  // {"en" : "HTTP header name.", "zh_CN": "HTTP头部名称。"}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // {"en" : "HTTP header value.", "zh_CN": "HTTP头部值。"}
  Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s GetPrefetchRequestStatusResponseFileListHeaders) String() string {
  return tea.Prettify(s)
}

func (s GetPrefetchRequestStatusResponseFileListHeaders) GoString() string {
  return s.String()
}

func (s *GetPrefetchRequestStatusResponseFileListHeaders) SetName(v string) *GetPrefetchRequestStatusResponseFileListHeaders {
  s.Name = &v
  return s
}

func (s *GetPrefetchRequestStatusResponseFileListHeaders) SetValue(v string) *GetPrefetchRequestStatusResponseFileListHeaders {
  s.Value = &v
  return s
}

type GetPrefetchRequestStatusResponseMetadata struct {
  // {"en" : "ID of the prefetch request.", "zh_CN": "预取请求ID。"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {"en" : "RFC 3339 date indicating when the prefetch request was submitted.", "zh_CN": "RFC 3339格式的日期，表示预取请求的提交时间。"}
  SubmissionTime *string `json:"submissionTime,omitempty" xml:"submissionTime,omitempty" require:"true"`
  // {"en" : "Range: [ 0 .. 100 ] 
  // A percentage that indicates the completion of the prefetch request.", "zh_CN": "取值范围: [ 0 .. 100 ] 
  // 预取任务的成功率。"}
  SuccessRate *int `json:"successRate,omitempty" xml:"successRate,omitempty" require:"true"`
  // {"en" : "Enum: waiting,inprogress,finished 
  // Indicates the status of the prefetch request.", "zh_CN": "取值范围: waiting,inprogress,finished 
  // 预取请求的任务执行状态，包括等待中，进行中，已完成等状态。"}
  Status *string `json:"status,omitempty" xml:"status,omitempty" require:"true"`
  // {"en" : "RFC 3339 date indicating when the prefetch completed.", "zh_CN": "RFC 3339格式的日期，表示预取完成的时间。"}
  FinishTime *string `json:"finishTime,omitempty" xml:"finishTime,omitempty"`
  // {"en" : "Indicates the corresponding API request ID.", "zh_CN": "API请求ID。"}
  ApiRequestId *string `json:"apiRequestId,omitempty" xml:"apiRequestId,omitempty" require:"true"`
}

func (s GetPrefetchRequestStatusResponseMetadata) String() string {
  return tea.Prettify(s)
}

func (s GetPrefetchRequestStatusResponseMetadata) GoString() string {
  return s.String()
}

func (s *GetPrefetchRequestStatusResponseMetadata) SetId(v string) *GetPrefetchRequestStatusResponseMetadata {
  s.Id = &v
  return s
}

func (s *GetPrefetchRequestStatusResponseMetadata) SetSubmissionTime(v string) *GetPrefetchRequestStatusResponseMetadata {
  s.SubmissionTime = &v
  return s
}

func (s *GetPrefetchRequestStatusResponseMetadata) SetSuccessRate(v int) *GetPrefetchRequestStatusResponseMetadata {
  s.SuccessRate = &v
  return s
}

func (s *GetPrefetchRequestStatusResponseMetadata) SetStatus(v string) *GetPrefetchRequestStatusResponseMetadata {
  s.Status = &v
  return s
}

func (s *GetPrefetchRequestStatusResponseMetadata) SetFinishTime(v string) *GetPrefetchRequestStatusResponseMetadata {
  s.FinishTime = &v
  return s
}

func (s *GetPrefetchRequestStatusResponseMetadata) SetApiRequestId(v string) *GetPrefetchRequestStatusResponseMetadata {
  s.ApiRequestId = &v
  return s
}




type GetPurgeSummaryPaths struct {
}

func (s GetPurgeSummaryPaths) String() string {
  return tea.Prettify(s)
}

func (s GetPurgeSummaryPaths) GoString() string {
  return s.String()
}

type GetPurgeSummaryParameters struct {
  // {"en" : "RFC 3339 date indicating the beginning of the time period. The time must be specified using the UTC timezone; it cannot be an offset. Example: startdate=2019-10-30T00:00:00Z ", "zh_CN": "RFC 3339格式的日期，表示查询的开始时间。必须使用UTC时区，不能是其它时区。例如：startdate=2019-10-30T00:00:00Z"}
  Startdate *string `json:"startdate,omitempty" xml:"startdate,omitempty" require:"true"`
  // {"en" : "RFC 3339 date indicating the end of the timeframe. The time must be specified using the UTC timezone; it cannot be an offset. Example: enddate=2019-10-30T00:00:00Z ", "zh_CN": "RFC 3339格式的日期，表示查询的结束时间。必须使用UTC时区，不能是其它时区。例如：enddate=2019-10-30T00:00:00Z "}
  Enddate *string `json:"enddate,omitempty" xml:"enddate,omitempty" require:"true"`
  // {"en" : "Enum: staging production 
  // Default: production 
  // Indicates the environment.", "zh_CN": "取值范围: staging, production 
  // 默认值: production 
  // 刷新请求对应的环境。"}
  Target *string `json:"target,omitempty" xml:"target,omitempty"`
}

func (s GetPurgeSummaryParameters) String() string {
  return tea.Prettify(s)
}

func (s GetPurgeSummaryParameters) GoString() string {
  return s.String()
}

func (s *GetPurgeSummaryParameters) SetStartdate(v string) *GetPurgeSummaryParameters {
  s.Startdate = &v
  return s
}

func (s *GetPurgeSummaryParameters) SetEnddate(v string) *GetPurgeSummaryParameters {
  s.Enddate = &v
  return s
}

func (s *GetPurgeSummaryParameters) SetTarget(v string) *GetPurgeSummaryParameters {
  s.Target = &v
  return s
}

type GetPurgeSummaryRequestHeader struct {
}

func (s GetPurgeSummaryRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetPurgeSummaryRequestHeader) GoString() string {
  return s.String()
}

type GetPurgeSummaryRequest struct {
}

func (s GetPurgeSummaryRequest) String() string {
  return tea.Prettify(s)
}

func (s GetPurgeSummaryRequest) GoString() string {
  return s.String()
}

type GetPurgeSummaryResponseHeader struct {
}

func (s GetPurgeSummaryResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetPurgeSummaryResponseHeader) GoString() string {
  return s.String()
}

type GetPurgeSummaryResponse struct {
  // {"en" : ">= 0 
  // Number of purge requests.", "zh_CN": "取值范围: >= 0 
  // 刷新请求的数量。"}
  Requests *int `json:"requests,omitempty" xml:"requests,omitempty" require:"true"`
  // {"en" : ">= 0 
  // Number of purge API requests with non-empty fileURLs.
  // ", "zh_CN": "取值范围: >= 0 
  // 包含了文件刷新类型的刷新请求的数量。"}
  FileRequests *int `json:"fileRequests,omitempty" xml:"fileRequests,omitempty" require:"true"`
  // {"en" : ">= 0 
  // Number of purge API requests with non-empty dirUrls.", "zh_CN": "取值范围: >= 0 
  // 包含了目录刷新类型的刷新请求的数量。"}
  DirRequests *int `json:"dirRequests,omitempty" xml:"dirRequests,omitempty" require:"true"`
  // {"en" : ">= 0 
  // Total number of fileUrls' entries in purge requests.", "zh_CN": "取值范围: >= 0 
  // 刷新请求中文件（fileURL）条目的总数。"}
  FileEntries *int `json:"fileEntries,omitempty" xml:"fileEntries,omitempty" require:"true"`
  // {"en" : ">= 0 
  // Total number of dirUrls' entries in purge requests.", "zh_CN": "取值范围: >= 0 
  // 刷新请求中目录（dirURL）条目的总数。"}
  DirEntries *int `json:"dirEntries,omitempty" xml:"dirEntries,omitempty" require:"true"`
  // {"en" : ">= 0 
  // Number of purge API requests with non-empty regexPatterns.", "zh_CN": "取值范围: >= 0 
  // 包含了正则表达式刷新类型的刷新请求的数量。"}
  RegexRequests *int `json:"regexRequests,omitempty" xml:"regexRequests,omitempty" require:"true"`
  // {"en" : ">= 0 
  // Total number of regexPatterns' entries in purge requests.", "zh_CN": "取值范围: >= 0 
  // 刷新请求中正则表达式条目的总数。"}
  RegexEntries *int `json:"regexEntries,omitempty" xml:"regexEntries,omitempty" require:"true"`
}

func (s GetPurgeSummaryResponse) String() string {
  return tea.Prettify(s)
}

func (s GetPurgeSummaryResponse) GoString() string {
  return s.String()
}

func (s *GetPurgeSummaryResponse) SetRequests(v int) *GetPurgeSummaryResponse {
  s.Requests = &v
  return s
}

func (s *GetPurgeSummaryResponse) SetFileRequests(v int) *GetPurgeSummaryResponse {
  s.FileRequests = &v
  return s
}

func (s *GetPurgeSummaryResponse) SetDirRequests(v int) *GetPurgeSummaryResponse {
  s.DirRequests = &v
  return s
}

func (s *GetPurgeSummaryResponse) SetFileEntries(v int) *GetPurgeSummaryResponse {
  s.FileEntries = &v
  return s
}

func (s *GetPurgeSummaryResponse) SetDirEntries(v int) *GetPurgeSummaryResponse {
  s.DirEntries = &v
  return s
}

func (s *GetPurgeSummaryResponse) SetRegexRequests(v int) *GetPurgeSummaryResponse {
  s.RegexRequests = &v
  return s
}

func (s *GetPurgeSummaryResponse) SetRegexEntries(v int) *GetPurgeSummaryResponse {
  s.RegexEntries = &v
  return s
}




type GetListOfPrefetchRequestsPaths struct {
}

func (s GetListOfPrefetchRequestsPaths) String() string {
  return tea.Prettify(s)
}

func (s GetListOfPrefetchRequestsPaths) GoString() string {
  return s.String()
}

type GetListOfPrefetchRequestsParameters struct {
  // {"en" : "RFC 3339 date indicating the beginning of the time period. It must be in UTC time, for example, '2021-03-01T01:00:00Z'.", "zh_CN": "RFC 3339格式的日期，表示查询的开始时间。必须为UTC时间，如'2021-03-01T01:00:00Z'。"}
  Startdate *string `json:"startdate,omitempty" xml:"startdate,omitempty"`
  // {"en" : "RFC 3339 date indicating the end of the time period. It must be in UTC time, for example, '2021-03-01T01:00:00Z'.", "zh_CN": "RFC 3339格式的日期，表示查询的结束时间。必须为UTC时间，如'2021-03-01T01:00:00Z'。"}
  Enddate *string `json:"enddate,omitempty" xml:"enddate,omitempty"`
  // {"en" : "Default: 0 Range: >= 0 
  // Index of the first prefetch request to return.", "zh_CN": "默认值: 0 取值范围: >= 0 
  // 查询起始位置。"}
  Offset *int `json:"offset,omitempty" xml:"offset,omitempty"`
  // {"en" : "Range: [ 1 .. 200 ] 
  // Maximum number of prefetch requests to return.", "zh_CN": "默认值: 100 取值范围: <= 200 
  // 每次查询的最大条数。"}
  Limit *int `json:"limit,omitempty" xml:"limit,omitempty"`
  // {"en" : "Enum: submissionTime,finishTime 
  // Controls the order in which prefetch requests are returned. The default is to return the most recently submitted request first.", "zh_CN": "取值范围: submissionTime,finishTime 
  // 返回查询结果的排序依据。默认按预取请求的创建时间降序排序。"}
  SortBy *string `json:"sortBy,omitempty" xml:"sortBy,omitempty"`
  // {"en" : "Enum: asc,desc 
  // Order of prefetch requests. The default is to return the most recent prefetch request first. Enter 'asc' for ascending order or 'desc' for descending order.", "zh_CN": "取值范围: asc,desc 
  // 返回查询结果的排序顺序。默认按预取请求的创建时间降序排序。"}
  SortOrder *string `json:"sortOrder,omitempty" xml:"sortOrder,omitempty"`
  // {"en" : "The value will be used to match on prefetch request names, hostnames, and the task IDs to limit the prefetch requests that are returned.", "zh_CN": "根据搜索关键字匹配预取请求的名称、ID和相关的加速域名，过滤预取请求。"}
  Search *string `json:"search,omitempty" xml:"search,omitempty"`
}

func (s GetListOfPrefetchRequestsParameters) String() string {
  return tea.Prettify(s)
}

func (s GetListOfPrefetchRequestsParameters) GoString() string {
  return s.String()
}

func (s *GetListOfPrefetchRequestsParameters) SetStartdate(v string) *GetListOfPrefetchRequestsParameters {
  s.Startdate = &v
  return s
}

func (s *GetListOfPrefetchRequestsParameters) SetEnddate(v string) *GetListOfPrefetchRequestsParameters {
  s.Enddate = &v
  return s
}

func (s *GetListOfPrefetchRequestsParameters) SetOffset(v int) *GetListOfPrefetchRequestsParameters {
  s.Offset = &v
  return s
}

func (s *GetListOfPrefetchRequestsParameters) SetLimit(v int) *GetListOfPrefetchRequestsParameters {
  s.Limit = &v
  return s
}

func (s *GetListOfPrefetchRequestsParameters) SetSortBy(v string) *GetListOfPrefetchRequestsParameters {
  s.SortBy = &v
  return s
}

func (s *GetListOfPrefetchRequestsParameters) SetSortOrder(v string) *GetListOfPrefetchRequestsParameters {
  s.SortOrder = &v
  return s
}

func (s *GetListOfPrefetchRequestsParameters) SetSearch(v string) *GetListOfPrefetchRequestsParameters {
  s.Search = &v
  return s
}

type GetListOfPrefetchRequestsRequestHeader struct {
}

func (s GetListOfPrefetchRequestsRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetListOfPrefetchRequestsRequestHeader) GoString() string {
  return s.String()
}

type GetListOfPrefetchRequestsRequest struct {
}

func (s GetListOfPrefetchRequestsRequest) String() string {
  return tea.Prettify(s)
}

func (s GetListOfPrefetchRequestsRequest) GoString() string {
  return s.String()
}

type GetListOfPrefetchRequestsResponseHeader struct {
}

func (s GetListOfPrefetchRequestsResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetListOfPrefetchRequestsResponseHeader) GoString() string {
  return s.String()
}

type GetListOfPrefetchRequestsResponse struct {
  // {"en" : "Range: >= 0 
  // Total number of prefetch requests matching the criteria specified in the query parameters. The actual number of entries in prefetchRequests may be fewer if you specified the parameter.", "zh_CN": "取值范围: >= 0 
  // 预取请求的总数。该数量取决于查询参数。"}
  Count *int `json:"count,omitempty" xml:"count,omitempty" require:"true"`
  // {"en" : "List of prefetch requests.", "zh_CN": "预取请求列表。"}
  PrefetchRequests []*GetListOfPrefetchRequestsResponsePrefetchRequests `json:"prefetchRequests,omitempty" xml:"prefetchRequests,omitempty" require:"true" type:"Repeated"`
}

func (s GetListOfPrefetchRequestsResponse) String() string {
  return tea.Prettify(s)
}

func (s GetListOfPrefetchRequestsResponse) GoString() string {
  return s.String()
}

func (s *GetListOfPrefetchRequestsResponse) SetCount(v int) *GetListOfPrefetchRequestsResponse {
  s.Count = &v
  return s
}

func (s *GetListOfPrefetchRequestsResponse) SetPrefetchRequests(v []*GetListOfPrefetchRequestsResponsePrefetchRequests) *GetListOfPrefetchRequestsResponse {
  s.PrefetchRequests = v
  return s
}

type GetListOfPrefetchRequestsResponsePrefetchRequests struct     {
  // {"en" : "prefetch request task ID.
  // ", "zh_CN": "预取请求的ID。
  // "}
  Id *string `json:"id,omitempty" xml:"id,omitempty"`
  // {"en" : " a short description of the prefetch task.", "zh_CN": "预取请求的简短描述。"}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // {"en" : "Enum: waiting,inprogress,finished 
  // Indicates the status of the prefetch request.", "zh_CN": "取值范围: waiting,inprogress,finished 
  // 预取请求的任务执行状态，包括等待中，进行中，已完成等状态。"}
  Status *string `json:"status,omitempty" xml:"status,omitempty"`
  // {"en" : "RFC 3339 date indicating when the prefetch request was submitted.", "zh_CN": "RFC 3339格式的日期，表示预取请求的提交时间。"}
  SubmissionTime *string `json:"submissionTime,omitempty" xml:"submissionTime,omitempty"`
  // {"en" : "RFC 3339 date indicating when the prefetch completed.", "zh_CN": "RFC 3339格式的日期，表示预取任务的完成时间。"}
  FinishTime *string `json:"finishTime,omitempty" xml:"finishTime,omitempty"`
  // {"en" : "Range: [ 0 .. 100 ] 
  // A percentage that indicates the completion of the prefetch request.", "zh_CN": "取值范围: [ 0 .. 100 ] 
  // 预取任务的成功率。"}
  SuccessRate *int `json:"successRate,omitempty" xml:"successRate,omitempty"`
  // {"en" : "Indicates the corresponding API request ID.", "zh_CN": "API请求ID。"}
  ApiRequestId *string `json:"apiRequestId,omitempty" xml:"apiRequestId,omitempty"`
  // {"en" : "Range: >= 1 
  // Number of URLs in the prefetch request.", "zh_CN": "取值范围: >= 1 
  // 预取请求中的URL数量。"}
  FileEntries *int `json:"fileEntries,omitempty" xml:"fileEntries,omitempty"`
  // {"en" : "List of hostnames that were prefetched.", "zh_CN": "预取请求涉及的加速域名。"}
  Hostnames []*string `json:"hostnames,omitempty" xml:"hostnames,omitempty" type:"Repeated"`
}

func (s GetListOfPrefetchRequestsResponsePrefetchRequests) String() string {
  return tea.Prettify(s)
}

func (s GetListOfPrefetchRequestsResponsePrefetchRequests) GoString() string {
  return s.String()
}

func (s *GetListOfPrefetchRequestsResponsePrefetchRequests) SetId(v string) *GetListOfPrefetchRequestsResponsePrefetchRequests {
  s.Id = &v
  return s
}

func (s *GetListOfPrefetchRequestsResponsePrefetchRequests) SetName(v string) *GetListOfPrefetchRequestsResponsePrefetchRequests {
  s.Name = &v
  return s
}

func (s *GetListOfPrefetchRequestsResponsePrefetchRequests) SetStatus(v string) *GetListOfPrefetchRequestsResponsePrefetchRequests {
  s.Status = &v
  return s
}

func (s *GetListOfPrefetchRequestsResponsePrefetchRequests) SetSubmissionTime(v string) *GetListOfPrefetchRequestsResponsePrefetchRequests {
  s.SubmissionTime = &v
  return s
}

func (s *GetListOfPrefetchRequestsResponsePrefetchRequests) SetFinishTime(v string) *GetListOfPrefetchRequestsResponsePrefetchRequests {
  s.FinishTime = &v
  return s
}

func (s *GetListOfPrefetchRequestsResponsePrefetchRequests) SetSuccessRate(v int) *GetListOfPrefetchRequestsResponsePrefetchRequests {
  s.SuccessRate = &v
  return s
}

func (s *GetListOfPrefetchRequestsResponsePrefetchRequests) SetApiRequestId(v string) *GetListOfPrefetchRequestsResponsePrefetchRequests {
  s.ApiRequestId = &v
  return s
}

func (s *GetListOfPrefetchRequestsResponsePrefetchRequests) SetFileEntries(v int) *GetListOfPrefetchRequestsResponsePrefetchRequests {
  s.FileEntries = &v
  return s
}

func (s *GetListOfPrefetchRequestsResponsePrefetchRequests) SetHostnames(v []*string) *GetListOfPrefetchRequestsResponsePrefetchRequests {
  s.Hostnames = v
  return s
}




type GetListOfPurgeRequestsPaths struct {
}

func (s GetListOfPurgeRequestsPaths) String() string {
  return tea.Prettify(s)
}

func (s GetListOfPurgeRequestsPaths) GoString() string {
  return s.String()
}

type GetListOfPurgeRequestsParameters struct {
  // {"en" : "RFC 3339 date indicating the beginning of the time period. It must be in UTC time, for example, '2019-11-01T01:00:00Z'.", "zh_CN": "RFC 3339格式的日期，表示查询的开始时间。必须为UTC时间，如'2019-11-01T01:00:00Z'。"}
  Startdate *string `json:"startdate,omitempty" xml:"startdate,omitempty"`
  // {"en" : "[ 0 .. 0 ] characters 
  // RFC 3339 date indicating the end of the time period. It must be in UTC time, for example, '2019-11-01T01:00:00Z'.", "zh_CN": "RFC 3339格式的日期，表示查询的结束时间。必须为UTC时间，如'2019-11-01T01:00:00Z'。"}
  Enddate *string `json:"enddate,omitempty" xml:"enddate,omitempty"`
  // {"en" : "Default: 0 >= 0 
  // Index of the first purge request to return. ", "zh_CN": "默认值: 0 取值范围: >= 0 
  // 查询起始位置。"}
  Offset *int `json:"offset,omitempty" xml:"offset,omitempty"`
  // {"en" : "Maximum number of purge requests to return.", "zh_CN": "默认值: 100 取值范围: <= 200 
  // 每次查询的最大条数。"}
  Limit *int `json:"limit,omitempty" xml:"limit,omitempty"`
  // {"en" : "Enum: submissionTime finishTime 
  // Default: submissionTime 
  // Controls the order in which purge requests are returned. The default is to return the most recently submitted request first.", "zh_CN": "取值范围: submissionTime, finishTime 
  // 默认值: submissionTime 
  // 返回结果的排序依据。默认按刷新请求的创建时间降序排序。"}
  SortBy *string `json:"sortBy,omitempty" xml:"sortBy,omitempty"`
  // {"en" : "Enum: asc desc 
  // Default: desc 
  // Order of purge requests. The default is to return the most recent purge request first.", "zh_CN": "取值范围: asc, desc 
  // 默认值: desc 
  // 返回结果的排序顺序。"}
  SortOrder *string `json:"sortOrder,omitempty" xml:"sortOrder,omitempty"`
  // {"en" : "The value will be used to match on hostnames and the purge task ID to limit the purge requests that are returned.", "zh_CN": "根据搜索关键字匹配加速域名和刷新任务的ID进行过滤。"}
  Search *string `json:"search,omitempty" xml:"search,omitempty"`
  // {"en" : "Enum: production staging 
  // The target environment of the purge request. By default, all are returned.", "zh_CN": "取值范围: production, staging 
  // 根据部署环境来查询刷新请求。默认返回所有环境的刷新请求。"}
  Target *string `json:"target,omitempty" xml:"target,omitempty"`
  // {"en" : "Enum: invalidate delete 
  // The type of purge request. By default, all are returned.", "zh_CN": "取值范围: invalidate, delete 
  // 刷新请求的类型。默认返回所有类型的刷新请求。"}
  Action *string `json:"action,omitempty" xml:"action,omitempty"`
  // {"en" : "[ 0 .. 100 ] 
  // A decimal value indicating the maximum success rate of purge requests to return. By default, all are returned.", "zh_CN": "取值范围: [ 0 .. 100 ] 
  // 指定最大成功率来查询刷新任务，以十进制值表示。默认返回所有刷新请求。"}
  MaxSuccessRate *int `json:"maxSuccessRate,omitempty" xml:"maxSuccessRate,omitempty"`
}

func (s GetListOfPurgeRequestsParameters) String() string {
  return tea.Prettify(s)
}

func (s GetListOfPurgeRequestsParameters) GoString() string {
  return s.String()
}

func (s *GetListOfPurgeRequestsParameters) SetStartdate(v string) *GetListOfPurgeRequestsParameters {
  s.Startdate = &v
  return s
}

func (s *GetListOfPurgeRequestsParameters) SetEnddate(v string) *GetListOfPurgeRequestsParameters {
  s.Enddate = &v
  return s
}

func (s *GetListOfPurgeRequestsParameters) SetOffset(v int) *GetListOfPurgeRequestsParameters {
  s.Offset = &v
  return s
}

func (s *GetListOfPurgeRequestsParameters) SetLimit(v int) *GetListOfPurgeRequestsParameters {
  s.Limit = &v
  return s
}

func (s *GetListOfPurgeRequestsParameters) SetSortBy(v string) *GetListOfPurgeRequestsParameters {
  s.SortBy = &v
  return s
}

func (s *GetListOfPurgeRequestsParameters) SetSortOrder(v string) *GetListOfPurgeRequestsParameters {
  s.SortOrder = &v
  return s
}

func (s *GetListOfPurgeRequestsParameters) SetSearch(v string) *GetListOfPurgeRequestsParameters {
  s.Search = &v
  return s
}

func (s *GetListOfPurgeRequestsParameters) SetTarget(v string) *GetListOfPurgeRequestsParameters {
  s.Target = &v
  return s
}

func (s *GetListOfPurgeRequestsParameters) SetAction(v string) *GetListOfPurgeRequestsParameters {
  s.Action = &v
  return s
}

func (s *GetListOfPurgeRequestsParameters) SetMaxSuccessRate(v int) *GetListOfPurgeRequestsParameters {
  s.MaxSuccessRate = &v
  return s
}

type GetListOfPurgeRequestsRequestHeader struct {
}

func (s GetListOfPurgeRequestsRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetListOfPurgeRequestsRequestHeader) GoString() string {
  return s.String()
}

type GetListOfPurgeRequestsRequest struct {
}

func (s GetListOfPurgeRequestsRequest) String() string {
  return tea.Prettify(s)
}

func (s GetListOfPurgeRequestsRequest) GoString() string {
  return s.String()
}

type GetListOfPurgeRequestsResponseHeader struct {
}

func (s GetListOfPurgeRequestsResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetListOfPurgeRequestsResponseHeader) GoString() string {
  return s.String()
}

type GetListOfPurgeRequestsResponse struct {
  // {"en" : ">= 0 
  // Total number of purge requests matching the criteria specified in the query parameters. The actual number of entries in purgeRequests may be fewer if you specified the <limit> parameter.", "zh_CN": "取值范围: >= 0 
  // 刷新请求的总数。该数值取决于查询参数。"}
  Count *int `json:"count,omitempty" xml:"count,omitempty" require:"true"`
  // {"en" : "List of purge requests.", "zh_CN": "刷新请求列表。"}
  PurgeRequests []*GetListOfPurgeRequestsResponsePurgeRequests `json:"purgeRequests,omitempty" xml:"purgeRequests,omitempty" require:"true" type:"Repeated"`
}

func (s GetListOfPurgeRequestsResponse) String() string {
  return tea.Prettify(s)
}

func (s GetListOfPurgeRequestsResponse) GoString() string {
  return s.String()
}

func (s *GetListOfPurgeRequestsResponse) SetCount(v int) *GetListOfPurgeRequestsResponse {
  s.Count = &v
  return s
}

func (s *GetListOfPurgeRequestsResponse) SetPurgeRequests(v []*GetListOfPurgeRequestsResponsePurgeRequests) *GetListOfPurgeRequestsResponse {
  s.PurgeRequests = v
  return s
}

type GetListOfPurgeRequestsResponsePurgeRequests struct     {
  // {"en" : "ID associated with the purge request. You can call the Query purge request status API to get further information about it.", "zh_CN": "刷新请求的ID。您可以调用'查询刷新任务详情'接口来获得更多信息。"}
  Id *string `json:"id,omitempty" xml:"id,omitempty"`
  // {"en" : "An RFC 3339 date indicating when the purge request was created.", "zh_CN": "RFC 3339格式的日期，表示刷新请求的创建时间。"}
  SubmissionTime *string `json:"submissionTime,omitempty" xml:"submissionTime,omitempty"`
  // {"en" : "Hostnames affected by the purge request.", "zh_CN": "刷新请求涉及的加速域名。"}
  Hostnames []*string `json:"hostnames,omitempty" xml:"hostnames,omitempty" type:"Repeated"`
  // {"en" : ">= 0 
  // Number of file URLs that were part of the purge request.", "zh_CN": "取值范围: >= 0 
  // 刷新请求中文件URL的数量。"}
  FileEntries *int `json:"fileEntries,omitempty" xml:"fileEntries,omitempty"`
  // {"en" : "[ 0 .. 20 ] 
  // Number of directory URLs that were part of the purge request.", "zh_CN": "取值范围: [ 0 .. 20 ] 
  // 刷新请求中目录URL的数量。"}
  DirEntries *int `json:"dirEntries,omitempty" xml:"dirEntries,omitempty"`
  // {"en" : "Enum: staging production 
  // Indicates whether the purge is in the staging or production environment.
  // ", "zh_CN": "取值范围: staging, production 
  // 刷新请求对应的环境，即演练环境或生产环境。"}
  Target *string `json:"target,omitempty" xml:"target,omitempty"`
  // {"en" : "[ 0 .. 100 ] 
  // A percentage that indicates the completion of the purge request. ", "zh_CN": "取值范围: [ 0 .. 100 ] 
  // 刷新请求任务的成功率。"}
  SuccessRate *int `json:"successRate,omitempty" xml:"successRate,omitempty"`
  // {"en" : "Enum: waiting inprogress finished 
  // Indicates the status of the purge request.", "zh_CN": "取值范围: waiting, inprogress, finished 
  // 刷新请求的执行状态，包括等待中，进行中，已完成等状态。"}
  Status *string `json:"status,omitempty" xml:"status,omitempty"`
  // {"en" : "RFC 3339 date indicating when the purge was completed. It can be empty if the purge is still in progress.", "zh_CN": "RFC 3339格式的日期，表示刷新完成的时间。如果刷新还在进行中，则返回空值。"}
  FinishTime *string `json:"finishTime,omitempty" xml:"finishTime,omitempty"`
  // {"en" : "Indicates the corresponding API request ID.", "zh_CN": "API请求ID。"}
  ApiRequestId *string `json:"apiRequestId,omitempty" xml:"apiRequestId,omitempty"`
  // {"en" : ">= 0 
  // Number of regex patterns that were part of the purge request.", "zh_CN": "取值范围: >= 0 
  // 刷新请求中正则表达式的条数。"}
  RegexEntries *int `json:"regexEntries,omitempty" xml:"regexEntries,omitempty"`
}

func (s GetListOfPurgeRequestsResponsePurgeRequests) String() string {
  return tea.Prettify(s)
}

func (s GetListOfPurgeRequestsResponsePurgeRequests) GoString() string {
  return s.String()
}

func (s *GetListOfPurgeRequestsResponsePurgeRequests) SetId(v string) *GetListOfPurgeRequestsResponsePurgeRequests {
  s.Id = &v
  return s
}

func (s *GetListOfPurgeRequestsResponsePurgeRequests) SetSubmissionTime(v string) *GetListOfPurgeRequestsResponsePurgeRequests {
  s.SubmissionTime = &v
  return s
}

func (s *GetListOfPurgeRequestsResponsePurgeRequests) SetHostnames(v []*string) *GetListOfPurgeRequestsResponsePurgeRequests {
  s.Hostnames = v
  return s
}

func (s *GetListOfPurgeRequestsResponsePurgeRequests) SetFileEntries(v int) *GetListOfPurgeRequestsResponsePurgeRequests {
  s.FileEntries = &v
  return s
}

func (s *GetListOfPurgeRequestsResponsePurgeRequests) SetDirEntries(v int) *GetListOfPurgeRequestsResponsePurgeRequests {
  s.DirEntries = &v
  return s
}

func (s *GetListOfPurgeRequestsResponsePurgeRequests) SetTarget(v string) *GetListOfPurgeRequestsResponsePurgeRequests {
  s.Target = &v
  return s
}

func (s *GetListOfPurgeRequestsResponsePurgeRequests) SetSuccessRate(v int) *GetListOfPurgeRequestsResponsePurgeRequests {
  s.SuccessRate = &v
  return s
}

func (s *GetListOfPurgeRequestsResponsePurgeRequests) SetStatus(v string) *GetListOfPurgeRequestsResponsePurgeRequests {
  s.Status = &v
  return s
}

func (s *GetListOfPurgeRequestsResponsePurgeRequests) SetFinishTime(v string) *GetListOfPurgeRequestsResponsePurgeRequests {
  s.FinishTime = &v
  return s
}

func (s *GetListOfPurgeRequestsResponsePurgeRequests) SetApiRequestId(v string) *GetListOfPurgeRequestsResponsePurgeRequests {
  s.ApiRequestId = &v
  return s
}

func (s *GetListOfPurgeRequestsResponsePurgeRequests) SetRegexEntries(v int) *GetListOfPurgeRequestsResponsePurgeRequests {
  s.RegexEntries = &v
  return s
}




type CreateAPrefetchRequestPaths struct {
}

func (s CreateAPrefetchRequestPaths) String() string {
  return tea.Prettify(s)
}

func (s CreateAPrefetchRequestPaths) GoString() string {
  return s.String()
}

type CreateAPrefetchRequestParameters struct {
}

func (s CreateAPrefetchRequestParameters) String() string {
  return tea.Prettify(s)
}

func (s CreateAPrefetchRequestParameters) GoString() string {
  return s.String()
}

type CreateAPrefetchRequestRequestHeader struct {
}

func (s CreateAPrefetchRequestRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s CreateAPrefetchRequestRequestHeader) GoString() string {
  return s.String()
}

type CreateAPrefetchRequestRequest struct {
  // {"en" : "Range: <= 1000 characters 
  // Enter a short description of the prefetch task.", "zh_CN": "取值范围: <= 1000 字符 
  // 预取请求的简短描述。"}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // {"en" : "", "zh_CN": ""}
  FileList []*CreateAPrefetchRequestRequestFileList `json:"fileList,omitempty" xml:"fileList,omitempty" require:"true" type:"Repeated"`
  // {"en" : "A list of continents representing the regions in which to perform the prefetch. Specify 'Mainland China' as the region, if prefetch by servers in mainland China only is desired. Omitting the field means the prefetch will be done by all regions' servers.", "zh_CN": "指定需要预取内容的大洲，以大洲英文全名表示，例如Asia, Europe。支持仅预取到中国大陆的服务器，区域名称以Mainland China表示。未指定区域时，表示预取内容到所有大洲的服务器。"}
  Regions []*string `json:"regions,omitempty" xml:"regions,omitempty" type:"Repeated"`
  // {"en" : "RFC 3339 date indicating when the prefetch should begin. This must be in UTC time, for example, '2021-03-06T00:00:00Z'.", "zh_CN": "RFC 3339格式的日期，表示开始预取的时间。必须使用UTC时间，例如'2021-03-06T00:00:00Z'。"}
  StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
  // {"en" : "ID of a webhook to call when the purge task completes. Wehook menas the callback endpoint as created via the 'Create a webhook' API.", "zh_CN": "刷新任务完成时要调用的webhook的ID。webhook是指通过“创建webhook接口”创建的回调接口。"}
  Webhook *string `json:"webhook,omitempty" xml:"webhook,omitempty"`
}

func (s CreateAPrefetchRequestRequest) String() string {
  return tea.Prettify(s)
}

func (s CreateAPrefetchRequestRequest) GoString() string {
  return s.String()
}

func (s *CreateAPrefetchRequestRequest) SetName(v string) *CreateAPrefetchRequestRequest {
  s.Name = &v
  return s
}

func (s *CreateAPrefetchRequestRequest) SetFileList(v []*CreateAPrefetchRequestRequestFileList) *CreateAPrefetchRequestRequest {
  s.FileList = v
  return s
}

func (s *CreateAPrefetchRequestRequest) SetRegions(v []*string) *CreateAPrefetchRequestRequest {
  s.Regions = v
  return s
}

func (s *CreateAPrefetchRequestRequest) SetStartTime(v string) *CreateAPrefetchRequestRequest {
  s.StartTime = &v
  return s
}

func (s *CreateAPrefetchRequestRequest) SetWebhook(v string) *CreateAPrefetchRequestRequest {
  s.Webhook = &v
  return s
}

type CreateAPrefetchRequestRequestFileList struct     {
  // {"en" : "Range: [ 10 .. 2048 ] characters 
  // A URL to prefetch. It must begin with 'http' or 'https' and can be up to 2048 characters.
  // ", "zh_CN": "取值范围: [ 10 .. 2048 ] 字符 
  // 预取的URL。必须以'http'或'https'开头，长度不超过2048个字符。"}
  Url *string `json:"url,omitempty" xml:"url,omitempty"`
  // {"en" : "If a URL's cache key depends on request headers, you can specify the header values that are applicable to prefetch one version of the URL.", "zh_CN": "如果需要在缓存键中加入请求头，可用该字段指定请求头。"}
  Headers []*CreateAPrefetchRequestRequestFileListHeaders `json:"headers,omitempty" xml:"headers,omitempty" type:"Repeated"`
}

func (s CreateAPrefetchRequestRequestFileList) String() string {
  return tea.Prettify(s)
}

func (s CreateAPrefetchRequestRequestFileList) GoString() string {
  return s.String()
}

func (s *CreateAPrefetchRequestRequestFileList) SetUrl(v string) *CreateAPrefetchRequestRequestFileList {
  s.Url = &v
  return s
}

func (s *CreateAPrefetchRequestRequestFileList) SetHeaders(v []*CreateAPrefetchRequestRequestFileListHeaders) *CreateAPrefetchRequestRequestFileList {
  s.Headers = v
  return s
}

type CreateAPrefetchRequestRequestFileListHeaders struct     {
  // {"en" : "HTTP header name.", "zh_CN": "HTTP 头部名称。"}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // {"en" : "HTTP header value.", "zh_CN": "HTTP 头部值。"}
  Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s CreateAPrefetchRequestRequestFileListHeaders) String() string {
  return tea.Prettify(s)
}

func (s CreateAPrefetchRequestRequestFileListHeaders) GoString() string {
  return s.String()
}

func (s *CreateAPrefetchRequestRequestFileListHeaders) SetName(v string) *CreateAPrefetchRequestRequestFileListHeaders {
  s.Name = &v
  return s
}

func (s *CreateAPrefetchRequestRequestFileListHeaders) SetValue(v string) *CreateAPrefetchRequestRequestFileListHeaders {
  s.Value = &v
  return s
}

type CreateAPrefetchRequestResponse struct {
}

func (s CreateAPrefetchRequestResponse) String() string {
  return tea.Prettify(s)
}

func (s CreateAPrefetchRequestResponse) GoString() string {
  return s.String()
}

type CreateAPrefetchRequestResponseHeader struct {
  // {"en":"The Location header is a URL representing the new prefetch request, for example, <code>Location: https://{domain}/cdn/prefetches/e91e8674-c2c5-4440-a1de-8b2ea99293dd</code>.", "zh_CN":"通过Location响应头返回新建的预取任务的URL。URL中包含预取任务的ID，可使用该ID调用'查询预取任务详情'接口来查看预取任务详情。URL示例：<code>Location: https://{domain}/cdn/prefetches/5dca2205f9e9cc0001df7b33"}
  Location *string `json:"Location,omitempty" xml:"Location,omitempty" require:"true"`
}

func (s CreateAPrefetchRequestResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s CreateAPrefetchRequestResponseHeader) GoString() string {
  return s.String()
}

func (s *CreateAPrefetchRequestResponseHeader) SetLocation(v string) *CreateAPrefetchRequestResponseHeader {
  s.Location = &v
  return s
}




type GetAvailablePurgeRequestsPaths struct {
}

func (s GetAvailablePurgeRequestsPaths) String() string {
  return tea.Prettify(s)
}

func (s GetAvailablePurgeRequestsPaths) GoString() string {
  return s.String()
}

type GetAvailablePurgeRequestsParameters struct {
  // {"en" : "Enum: staging,production 
  // Default: production 
  // Indicates the environment.", "zh_CN": "取值范围: staging,production 
  // 默认值: production 
  // 指定环境查询可用额度，即演练环境或生产环境。"}
  Target *string `json:"target,omitempty" xml:"target,omitempty" require:"true"`
}

func (s GetAvailablePurgeRequestsParameters) String() string {
  return tea.Prettify(s)
}

func (s GetAvailablePurgeRequestsParameters) GoString() string {
  return s.String()
}

func (s *GetAvailablePurgeRequestsParameters) SetTarget(v string) *GetAvailablePurgeRequestsParameters {
  s.Target = &v
  return s
}

type GetAvailablePurgeRequestsRequestHeader struct {
}

func (s GetAvailablePurgeRequestsRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetAvailablePurgeRequestsRequestHeader) GoString() string {
  return s.String()
}

type GetAvailablePurgeRequestsRequest struct {
}

func (s GetAvailablePurgeRequestsRequest) String() string {
  return tea.Prettify(s)
}

func (s GetAvailablePurgeRequestsRequest) GoString() string {
  return s.String()
}

type GetAvailablePurgeRequestsResponseHeader struct {
}

func (s GetAvailablePurgeRequestsResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetAvailablePurgeRequestsResponseHeader) GoString() string {
  return s.String()
}

type GetAvailablePurgeRequestsResponse struct {
  // {"en" : "Range: >= 0 
  // Number of file URL purge requests available.", "zh_CN": "取值范围: >= 0 
  // 文件刷新当前的可用额度。"}
  FilePurgeTokens *int `json:"filePurgeTokens,omitempty" xml:"filePurgeTokens,omitempty" require:"true"`
  // {"en" : "Range: >= 0 
  //  Number of directory URL purge requests available.", "zh_CN": "取值范围: >= 0 
  // 目录刷新当前的可用额度。"}
  DirPurgeTokens *int `json:"dirPurgeTokens,omitempty" xml:"dirPurgeTokens,omitempty" require:"true"`
}

func (s GetAvailablePurgeRequestsResponse) String() string {
  return tea.Prettify(s)
}

func (s GetAvailablePurgeRequestsResponse) GoString() string {
  return s.String()
}

func (s *GetAvailablePurgeRequestsResponse) SetFilePurgeTokens(v int) *GetAvailablePurgeRequestsResponse {
  s.FilePurgeTokens = &v
  return s
}

func (s *GetAvailablePurgeRequestsResponse) SetDirPurgeTokens(v int) *GetAvailablePurgeRequestsResponse {
  s.DirPurgeTokens = &v
  return s
}




