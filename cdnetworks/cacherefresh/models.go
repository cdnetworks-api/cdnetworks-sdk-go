package cacherefresh

// This file is auto-generated, don't edit it. Thanks.
import (
  "github.com/alibabacloud-go/tea/tea"
)

type GetPurgeQuotaRequest struct {
}

func (s GetPurgeQuotaRequest) String() string {
  return tea.Prettify(s)
}

func (s GetPurgeQuotaRequest) GoString() string {
  return s.String()
}

type GetPurgeQuotaRequestHeader struct {
}

func (s GetPurgeQuotaRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetPurgeQuotaRequestHeader) GoString() string {
  return s.String()
}

type GetPurgeQuotaPaths struct {
}

func (s GetPurgeQuotaPaths) String() string {
  return tea.Prettify(s)
}

func (s GetPurgeQuotaPaths) GoString() string {
  return s.String()
}

type GetPurgeQuotaParameters struct {
}

func (s GetPurgeQuotaParameters) String() string {
  return tea.Prettify(s)
}

func (s GetPurgeQuotaParameters) GoString() string {
  return s.String()
}

type GetPurgeQuotaResponse struct {
  // {"en":"status code, 0 indicates success and 1 indicates failure.","zh_CN":"状态码，0表示成功，1表示失败"}
  Code *int `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"response message","zh_CN":"响应消息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"purge file quota","zh_CN":"推送文件日上限"}
  UrlQuota *int `json:"urlQuota,omitempty" xml:"urlQuota,omitempty" require:"true"`
  // {"en":"purge file remaining quota","zh_CN":"推送文件日剩余量"}
  RemainingUrlQuota *int `json:"remainingUrlQuota,omitempty" xml:"remainingUrlQuota,omitempty" require:"true"`
  // {"en":"purge dir quota","zh_CN":"推送目录日上限"}
  DirQuota *int `json:"dirQuota,omitempty" xml:"dirQuota,omitempty" require:"true"`
  // {"en":"purge dir remaining quota","zh_CN":"推送目录剩余量"}
  RemainingDirQuota *int `json:"remainingDirQuota,omitempty" xml:"remainingDirQuota,omitempty" require:"true"`
  // {"en":"purge regex quota","zh_CN":"推送正则日上限"}
  RegexQuota *int `json:"regexQuota,omitempty" xml:"regexQuota,omitempty" require:"true"`
  // {"en":"purge regex remaining quota","zh_CN":"推送正则剩余量"}
  RemainingRegexQuota *int `json:"remainingRegexQuota,omitempty" xml:"remainingRegexQuota,omitempty" require:"true"`
  // {"en":"purge tag quota","zh_CN":"推送tag日上限"}
  TagQuota *int `json:"tagQuota,omitempty" xml:"tagQuota,omitempty" require:"true"`
  // {"en":"purge tag remaining quota","zh_CN":"推送tag剩余量"}
  RemainingTagQuota *int `json:"remainingTagQuota,omitempty" xml:"remainingTagQuota,omitempty" require:"true"`
}

func (s GetPurgeQuotaResponse) String() string {
  return tea.Prettify(s)
}

func (s GetPurgeQuotaResponse) GoString() string {
  return s.String()
}

func (s *GetPurgeQuotaResponse) SetCode(v int) *GetPurgeQuotaResponse {
  s.Code = &v
  return s
}

func (s *GetPurgeQuotaResponse) SetMessage(v string) *GetPurgeQuotaResponse {
  s.Message = &v
  return s
}

func (s *GetPurgeQuotaResponse) SetUrlQuota(v int) *GetPurgeQuotaResponse {
  s.UrlQuota = &v
  return s
}

func (s *GetPurgeQuotaResponse) SetRemainingUrlQuota(v int) *GetPurgeQuotaResponse {
  s.RemainingUrlQuota = &v
  return s
}

func (s *GetPurgeQuotaResponse) SetDirQuota(v int) *GetPurgeQuotaResponse {
  s.DirQuota = &v
  return s
}

func (s *GetPurgeQuotaResponse) SetRemainingDirQuota(v int) *GetPurgeQuotaResponse {
  s.RemainingDirQuota = &v
  return s
}

func (s *GetPurgeQuotaResponse) SetRegexQuota(v int) *GetPurgeQuotaResponse {
  s.RegexQuota = &v
  return s
}

func (s *GetPurgeQuotaResponse) SetRemainingRegexQuota(v int) *GetPurgeQuotaResponse {
  s.RemainingRegexQuota = &v
  return s
}

func (s *GetPurgeQuotaResponse) SetTagQuota(v int) *GetPurgeQuotaResponse {
  s.TagQuota = &v
  return s
}

func (s *GetPurgeQuotaResponse) SetRemainingTagQuota(v int) *GetPurgeQuotaResponse {
  s.RemainingTagQuota = &v
  return s
}

type GetPurgeQuotaResponseHeader struct {
}

func (s GetPurgeQuotaResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetPurgeQuotaResponseHeader) GoString() string {
  return s.String()
}




type GetPrefetchQuotaRequest struct {
}

func (s GetPrefetchQuotaRequest) String() string {
  return tea.Prettify(s)
}

func (s GetPrefetchQuotaRequest) GoString() string {
  return s.String()
}

type GetPrefetchQuotaRequestHeader struct {
}

func (s GetPrefetchQuotaRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetPrefetchQuotaRequestHeader) GoString() string {
  return s.String()
}

type GetPrefetchQuotaPaths struct {
}

func (s GetPrefetchQuotaPaths) String() string {
  return tea.Prettify(s)
}

func (s GetPrefetchQuotaPaths) GoString() string {
  return s.String()
}

type GetPrefetchQuotaParameters struct {
}

func (s GetPrefetchQuotaParameters) String() string {
  return tea.Prettify(s)
}

func (s GetPrefetchQuotaParameters) GoString() string {
  return s.String()
}

type GetPrefetchQuotaResponse struct {
  // {"en":"status code, 0 indicates success and 1 indicates failure","zh_CN":"状态码，0表示成功，1表示失败"}
  Code *int `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"response message","zh_CN":"响应消息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"prefetch quota","zh_CN":"预取条数日上限"}
  PrefetchQuota *int `json:"prefetchQuota,omitempty" xml:"prefetchQuota,omitempty" require:"true"`
  // {"en":"prefetch remaining quota","zh_CN":"预取条数剩余量"}
  RemainingPrefetchQuota *int `json:"remainingPrefetchQuota,omitempty" xml:"remainingPrefetchQuota,omitempty" require:"true"`
  // {"en":"prefetch volume quota","zh_CN":"预取大小日上限"}
  PrefetchVolumeQuota *int `json:"prefetchVolumeQuota,omitempty" xml:"prefetchVolumeQuota,omitempty" require:"true"`
  // {"en":"prefetch volume volume quota","zh_CN":"预取大小剩余量"}
  RemainingPrefetchVolumeQuota *int `json:"remainingPrefetchVolumeQuota,omitempty" xml:"remainingPrefetchVolumeQuota,omitempty" require:"true"`
}

func (s GetPrefetchQuotaResponse) String() string {
  return tea.Prettify(s)
}

func (s GetPrefetchQuotaResponse) GoString() string {
  return s.String()
}

func (s *GetPrefetchQuotaResponse) SetCode(v int) *GetPrefetchQuotaResponse {
  s.Code = &v
  return s
}

func (s *GetPrefetchQuotaResponse) SetMessage(v string) *GetPrefetchQuotaResponse {
  s.Message = &v
  return s
}

func (s *GetPrefetchQuotaResponse) SetPrefetchQuota(v int) *GetPrefetchQuotaResponse {
  s.PrefetchQuota = &v
  return s
}

func (s *GetPrefetchQuotaResponse) SetRemainingPrefetchQuota(v int) *GetPrefetchQuotaResponse {
  s.RemainingPrefetchQuota = &v
  return s
}

func (s *GetPrefetchQuotaResponse) SetPrefetchVolumeQuota(v int) *GetPrefetchQuotaResponse {
  s.PrefetchVolumeQuota = &v
  return s
}

func (s *GetPrefetchQuotaResponse) SetRemainingPrefetchVolumeQuota(v int) *GetPrefetchQuotaResponse {
  s.RemainingPrefetchVolumeQuota = &v
  return s
}

type GetPrefetchQuotaResponseHeader struct {
}

func (s GetPrefetchQuotaResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetPrefetchQuotaResponseHeader) GoString() string {
  return s.String()
}




type PurgeUrlsMatchingRegexRequest struct {
  // {"en":"The set of regular URLs for which the cache needs to be cleared. Requirements for the format of regular URLs:<br>The URL should conform to a regular expression. Example input: http://www.a.com/(.).png. If you want to push the regular expression http://www.abc.com/test/..txt, when using the interface, you need to escape the backslash, that is, http://www.abc.com/test/.*\.txt.<br>The domain name of each regular URL must be a domain name accelerated by our company.<br>The maximum length of each regular URL is 2000 characters.<br>If special characters are included in the regular URL, they need to be escaped in UTF - 8 format.<br>For the same regular URL, the system will remove duplicates before submission.<br>The default limit is 500 per day.","zh_CN":"要清理缓存的正则url集合，正则url的格式要求：<br>1）URL 符合正则表达式，输入示例：http://www.a.com/(.*).png。如要推送&nbsp;http://www.abc.com/test/.*\.txt&nbsp;正则表达式，在使用接口需要对反斜杠进行转义，即&nbsp;http://www.abc.com/test/.*\.txt。<br>2）每个正则url所在的域名必须是在我司加速的域名。<br>3）每个正则url最大长度 2000 字符。<br>4）正则url中如果包含特殊字符，需要进行转义，采用utf-8方式转义。<br>5) 相同的正则URL，系统会去重后提交。<br>6）默认500条/天。"}
  UrlRegulars []*string `json:"urlRegulars,omitempty" xml:"urlRegulars,omitempty" type:"Repeated"`
}

func (s PurgeUrlsMatchingRegexRequest) String() string {
  return tea.Prettify(s)
}

func (s PurgeUrlsMatchingRegexRequest) GoString() string {
  return s.String()
}

func (s *PurgeUrlsMatchingRegexRequest) SetUrlRegulars(v []*string) *PurgeUrlsMatchingRegexRequest {
  s.UrlRegulars = v
  return s
}

type PurgeUrlsMatchingRegexRequestHeader struct {
}

func (s PurgeUrlsMatchingRegexRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s PurgeUrlsMatchingRegexRequestHeader) GoString() string {
  return s.String()
}

type PurgeUrlsMatchingRegexPaths struct {
}

func (s PurgeUrlsMatchingRegexPaths) String() string {
  return tea.Prettify(s)
}

func (s PurgeUrlsMatchingRegexPaths) GoString() string {
  return s.String()
}

type PurgeUrlsMatchingRegexParameters struct {
}

func (s PurgeUrlsMatchingRegexParameters) String() string {
  return tea.Prettify(s)
}

func (s PurgeUrlsMatchingRegexParameters) GoString() string {
  return s.String()
}

type PurgeUrlsMatchingRegexResponse struct {
  // {"en":"Status code indicating the result of the task creation","zh_CN":"表示任务创建结果的状态码"}
  Code *string `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
  // {"en":"Indicates the response message of the system after the task is submitted.","zh_CN":"表示任务提交后，系统的响应消息"}
  Message *string `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
  // {"en":"After the interface is called once and the task is successfully submitted, it will return an itemamId, which is the unique identifier of the task submitted at that time. The itemId can be used to query the status (success/failure) of the task in batches.","zh_CN":"调用一次接口并提交任务成功后，将返回一个iteamId，是当次提交任务的唯一标识，通过itemId可批量查询任务的状态（成功/失败）。"}
  ItemId *string `json:"itemId,omitempty" xml:"itemId,omitempty" require:"true"`
}

func (s PurgeUrlsMatchingRegexResponse) String() string {
  return tea.Prettify(s)
}

func (s PurgeUrlsMatchingRegexResponse) GoString() string {
  return s.String()
}

func (s *PurgeUrlsMatchingRegexResponse) SetCode(v string) *PurgeUrlsMatchingRegexResponse {
  s.Code = &v
  return s
}

func (s *PurgeUrlsMatchingRegexResponse) SetMessage(v string) *PurgeUrlsMatchingRegexResponse {
  s.Message = &v
  return s
}

func (s *PurgeUrlsMatchingRegexResponse) SetItemId(v string) *PurgeUrlsMatchingRegexResponse {
  s.ItemId = &v
  return s
}

type PurgeUrlsMatchingRegexResponseHeader struct {
}

func (s PurgeUrlsMatchingRegexResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s PurgeUrlsMatchingRegexResponseHeader) GoString() string {
  return s.String()
}




type PurgeRequest struct {
  // {"en":"The set of URLs for which the cache needs to be cleared. Requirements for the URL format:\nThe URL must start with 'http://' or 'https://'. Example input: http://www.a.com/image/test.png.\nThe maximum length of each URL is 2000 characters.\nThe domain name of each URL must be a domain name accelerated by our company.\nThe number of refreshed URLs should not exceed 5000 per day (adjustable at the account level, contact technical support for adjustment).\nThe total number of URLs and dirs in each API call should not exceed 500.\nNote: URLs and dirs cannot be both empty.","zh_CN":"要清理缓存的url集合，url的格式要求：\n1）URL 必须以'http://' 或 'https://'开头，输入示例：http://www.a.com/image/test.png。\n2）每个url最大长度 2000字符。\n3）每个url所在的域名必须是在我司加速的域名。\n4）刷新url每日不超过5000条（账号粒度可调，可联系技术支持调整）。\n5）每次接口调用urls和dirs的总数不超过500条。\n注意：urls和dirs不能同时为空。"}
  Urls []*string `json:"urls,omitempty" xml:"urls,omitempty" type:"Repeated"`
  // {"en":"The set of directories for which the cache needs to be cleared. Requirements for the dir format:\nIt must start with 'http://domain name/' and end with '/'. For example, to refresh all files in the directory \"test\", the input format should be: http://www.a.com/test/.\nThe maximum length of each directory is 1000 characters.\nThe domain name of each directory must be a domain name accelerated by our company.\nDirectory refresh has a default expiration, and directory deletion is not supported.\nThe number of refreshed directories should not exceed 500 per day (adjustable at the account level, contact technical support for adjustment).\nNote: URLs and dirs cannot be both empty.","zh_CN":"指要清理缓存的目录集合，dir的格式要求：\n1）必须以'http://域名/'开始，以'/'结尾， 如刷新目录test下所有文件，输入格式为：http://www.a.com/test/；\n2）每个目录最大长度 1000 字符。\n3）每个目录所在的域名必须是在我司加速的域名。\n4）目录刷新默认过期，不支持目录删除。\n5）刷新目录每日不超过500条（账号粒度可调，可联系技术支持调整）\n注意：urls和dirs不能同时为空。"}
  Dirs []*string `json:"dirs,omitempty" xml:"dirs,omitempty" type:"Repeated"`
  // {"defaultValue":"default","en":"1) default: default value, the url cache is processed using the pre-configured operation type of domain. When no value is assigned to this element or the element is not submitted, the configuration of domain is read by default.\n2) delete: ignore the operation type in the domain configuration, directly delete the cache file of the submitted url.\n3) expire: Ignore the operation type in the domain  configuration, and set the file with the cached commit url to expire. When there is a http access for the first time, a request is made to the origin server to check if the file is up-to-date. If the origin has a new file, the new version is directly pulled back from the source station and returned to the user. If there is no update, then the source station responds with http code 304, and the CDN provides the cache file in the edge node to the user.","zh_CN":"仅对入参'urls'有效，指清理url缓存要以哪种类型操作：\n1）default：默认值，以频道预先配置好的操作类型处理url缓存，当不赋值或未传参时，默认取频道配置。\n2）delete：忽略频道配置里的操作类型，当前提交urls里的所有url，直接删除节点的缓存文件\n3）expire：忽略频道配置里的操作类型，当前提交urls里的所有url，将有缓存的节点置为过期，当第一次有访问时，回源站校验文件是否更新，有更新时从源站重新拉取新版本返回给客户，未有更新时则源站响应304，提供节点缓存文件给客户。","exampleValue":"default,delete,expire"}
  UrlAction *string `json:"urlAction,omitempty" xml:"urlAction,omitempty"`
  // {"defaultValue":"default","en":"It refers to the type of operation to clean up the dir cache, which is only valid for dirs elements. When no assignment or parameters are passed, the default channel configuration is selected. The optional values of this parameter are as follows:<br>\n1) default: default value, the dir cache is processed using the pre-configured operation type of domain. When no value is assigned to this element or the element is not submitted, the configuration of domain is read by default.<br>\n2) delete: ignore the type of operation in the channel configuration, delete the cache directory of the node directly.<br>\n3) expire: ignore the type of operation in the channel configuration, set the cached node to expire, when the first visit, check whether the directory of the source station is updated, when there is an update, retrieve the new version from the source station to return to the customer, and when there is no update, the source station responds 304, providing the cached directory of the node to the customer.<br>\nNote: The use of directory delete function will result in all files cached in the directory empty, all files need to be retrieved, resulting in increased backsource bandwidth. Because of the high risk, the permission does not open by default. If you need to open this function, please contact the corresponding customer service technical support to apply for opening. Only the customer\\'s account can be transferred to delete after opening.","zh_CN":"指清理dir缓存要以哪种类型操作，仅对dirs元素有效，当不赋值或未传参时，默认取频道配置。该参数可选值如下：\n1）default：默认值，以频道预先配置好的操作类型处理目录缓存，当不赋值或未传参时，默认取频道配置。\n2）delete：忽略频道配置里的操作类型，直接删除节点的缓存目录。\n3）expire：忽略频道配置里的操作类型，将有缓存的节点置为过期，当第一次有访问时，回源站校验目录是否更新，有更新时从源站重新拉取新版本返回给客户，未有更新时则源站响应304，提供节点缓存目录给客户。\n注：使用目录删除（delete）功能，会导致该目录下所有文件缓存全部清空，所有文件需要重新回源拉取，造成回源带宽增加。由于风险较大，该权限默认不开启，如需开启该功能，请联系对应的客服技术支持进行申请开启，开通后只有该客户的账号才能传入delete。","exampleValue":"default,delete,expire"}
  DirAction *string `json:"dirAction,omitempty" xml:"dirAction,omitempty"`
}

func (s PurgeRequest) String() string {
  return tea.Prettify(s)
}

func (s PurgeRequest) GoString() string {
  return s.String()
}

func (s *PurgeRequest) SetUrls(v []*string) *PurgeRequest {
  s.Urls = v
  return s
}

func (s *PurgeRequest) SetDirs(v []*string) *PurgeRequest {
  s.Dirs = v
  return s
}

func (s *PurgeRequest) SetUrlAction(v string) *PurgeRequest {
  s.UrlAction = &v
  return s
}

func (s *PurgeRequest) SetDirAction(v string) *PurgeRequest {
  s.DirAction = &v
  return s
}

type PurgeRequestHeader struct {
}

func (s PurgeRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s PurgeRequestHeader) GoString() string {
  return s.String()
}

type PurgePaths struct {
}

func (s PurgePaths) String() string {
  return tea.Prettify(s)
}

func (s PurgePaths) GoString() string {
  return s.String()
}

type PurgeParameters struct {
}

func (s PurgeParameters) String() string {
  return tea.Prettify(s)
}

func (s PurgeParameters) GoString() string {
  return s.String()
}

type PurgeResponse struct {
  // {"en":"The status code of the task creation result:\n1) 1 means success,\n2) 0 means failure.","zh_CN":"表示任务创建结果的状态码:\n1) 1表示成功\n2) 0表示失败","exampleValue":"0,1"}
  Code *int `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
  // {"en":"Content system response message after submitting the task.","zh_CN":"表示任务提交后，系统的响应消息"}
  Message *string `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
  // {"en":"After calling the API once and submitting the task successfully, the content system will return an itemId. This ID is the unique identifier for each submission. You can use itemId to batch query the status (success/failure) of the task.","zh_CN":"调用一次接口并提交任务成功后，将返回一个iteamId，是当次提交任务的唯一标识，通过itemId可批量查询任务的状态（成功/失败）。"}
  ItemId *string `json:"itemId,omitempty" xml:"itemId,omitempty" require:"true"`
}

func (s PurgeResponse) String() string {
  return tea.Prettify(s)
}

func (s PurgeResponse) GoString() string {
  return s.String()
}

func (s *PurgeResponse) SetCode(v int) *PurgeResponse {
  s.Code = &v
  return s
}

func (s *PurgeResponse) SetMessage(v string) *PurgeResponse {
  s.Message = &v
  return s
}

func (s *PurgeResponse) SetItemId(v string) *PurgeResponse {
  s.ItemId = &v
  return s
}

type PurgeResponseHeader struct {
}

func (s PurgeResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s PurgeResponseHeader) GoString() string {
  return s.String()
}




type QueryPurgeResidualsRequest struct {
}

func (s QueryPurgeResidualsRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryPurgeResidualsRequest) GoString() string {
  return s.String()
}

type QueryPurgeResidualsRequestHeader struct {
}

func (s QueryPurgeResidualsRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryPurgeResidualsRequestHeader) GoString() string {
  return s.String()
}

type QueryPurgeResidualsPaths struct {
}

func (s QueryPurgeResidualsPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryPurgeResidualsPaths) GoString() string {
  return s.String()
}

type QueryPurgeResidualsParameters struct {
  // {"en":"Query the feature's daily limit:\nPurge: Query the number of daily purge\nFetch: Query the number and size of daily prefetch","zh_CN":"用于指定查询哪种业务类型的每日资源上限，可选值有：\npurge：表示查询每日刷新缓存的数量限制\nfetch：表示查询每日预取文件的数量、大小限制","exampleValue":"purge,fetch"}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
}

func (s QueryPurgeResidualsParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryPurgeResidualsParameters) GoString() string {
  return s.String()
}

func (s *QueryPurgeResidualsParameters) SetType(v string) *QueryPurgeResidualsParameters {
  s.Type = &v
  return s
}

type QueryPurgeResidualsResponse struct {
  // {"en":"It is the logo of our company.","zh_CN":"数据提供方"}
  Supplier *string `json:"supplier,omitempty" xml:"supplier,omitempty" require:"true"`
  // {"en":"The status code of the query result:\n1) 0 means success,\n2) 1 means failure.","zh_CN":"查询结果:\n1) 0表示成功\n2) 1表示失败","exampleValue":"0,1"}
  Code *int `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Content system response message after query.","zh_CN":"查询的响应消息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"The maximum number of the purge url today.","zh_CN":"当天url刷新任务允许提交的最大数量"}
  UrlUpper *int `json:"urlUpper,omitempty" xml:"urlUpper,omitempty" require:"true"`
  // {"en":"The maximum number of the purge directory today.","zh_CN":"当天目录刷新任务允许提交的最大数量"}
  DirUpper *int `json:"dirUpper,omitempty" xml:"dirUpper,omitempty" require:"true"`
  // {"en":"The number of urls that can be purged today.","zh_CN":"当天url刷新任务允许提交的剩余数量"}
  UrlRemain *int `json:"urlRemain,omitempty" xml:"urlRemain,omitempty" require:"true"`
  // {"en":"The number of directorys that can be purged today.","zh_CN":"当天目录刷新任务允许提交的剩余数量"}
  DirRemain *int `json:"dirRemain,omitempty" xml:"dirRemain,omitempty" require:"true"`
  // {"en":"The maximum number of the purge tag today.","zh_CN":"当天tag任务允许提交的最大数量"}
  TagUpper *int `json:"tagUpper,omitempty" xml:"tagUpper,omitempty" require:"true"`
  // {"en":"The number of urls that can be purged tag today.","zh_CN":"当天tag任务允许提交的剩余数量"}
  TagRemain *int `json:"tagRemain,omitempty" xml:"tagRemain,omitempty" require:"true"`
}

func (s QueryPurgeResidualsResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryPurgeResidualsResponse) GoString() string {
  return s.String()
}

func (s *QueryPurgeResidualsResponse) SetSupplier(v string) *QueryPurgeResidualsResponse {
  s.Supplier = &v
  return s
}

func (s *QueryPurgeResidualsResponse) SetCode(v int) *QueryPurgeResidualsResponse {
  s.Code = &v
  return s
}

func (s *QueryPurgeResidualsResponse) SetMessage(v string) *QueryPurgeResidualsResponse {
  s.Message = &v
  return s
}

func (s *QueryPurgeResidualsResponse) SetUrlUpper(v int) *QueryPurgeResidualsResponse {
  s.UrlUpper = &v
  return s
}

func (s *QueryPurgeResidualsResponse) SetDirUpper(v int) *QueryPurgeResidualsResponse {
  s.DirUpper = &v
  return s
}

func (s *QueryPurgeResidualsResponse) SetUrlRemain(v int) *QueryPurgeResidualsResponse {
  s.UrlRemain = &v
  return s
}

func (s *QueryPurgeResidualsResponse) SetDirRemain(v int) *QueryPurgeResidualsResponse {
  s.DirRemain = &v
  return s
}

func (s *QueryPurgeResidualsResponse) SetTagUpper(v int) *QueryPurgeResidualsResponse {
  s.TagUpper = &v
  return s
}

func (s *QueryPurgeResidualsResponse) SetTagRemain(v int) *QueryPurgeResidualsResponse {
  s.TagRemain = &v
  return s
}

type QueryPurgeResidualsResponseHeader struct {
}

func (s QueryPurgeResidualsResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryPurgeResidualsResponseHeader) GoString() string {
  return s.String()
}




type PurgeFileWithTagRequest struct {
  // {"en":"Tag value:
  // 1. Separate multiple tag values with commas, allowing a maximum of 30 values.
  // 2. Each tag value can have a maximum length of 128 characters.
  // 3. Special characters are not allowed. These include Chinese characters, spaces, and *''(),:;<=>?@[]{}<>.", "zh_CN":"tag的值：
  // 1.多个tag值用,隔开,最大允许30个值
  // 2.每个tag值最大长度为128字符
  // 3.不允许特殊字符，特殊字符是指：中文，空格符及*''(),:;&lt;=&gt;?@[]{}<>"}
  Tag *string `json:"tag,omitempty" xml:"tag,omitempty" require:"true"`
  // {'en':'1. If left blank, the default value is 0.
  // 2. 0 represents deletion; 1 represents expiration.', 'zh_CN':'刷新动作
  // 1、不填的话默认是删除
  // 2、0代表删除；1代表过期'}
  Action *string `json:"action,omitempty" xml:"action,omitempty"`
}

func (s PurgeFileWithTagRequest) String() string {
  return tea.Prettify(s)
}

func (s PurgeFileWithTagRequest) GoString() string {
  return s.String()
}

func (s *PurgeFileWithTagRequest) SetTag(v string) *PurgeFileWithTagRequest {
  s.Tag = &v
  return s
}

func (s *PurgeFileWithTagRequest) SetAction(v string) *PurgeFileWithTagRequest {
  s.Action = &v
  return s
}

type PurgeFileWithTagResponse struct {
  // {'en':'code', 'zh_CN':'表示任务创建结果的状态码'}
  Code *string `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
  // {'en':'message', 'zh_CN':'表示任务提交后，系统的响应消息'}
  Message *string `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
  // {'en':'itemId', 'zh_CN':'调用一次接口并提交任务成功后，将返回一个iteamId，是当次提交任务的唯一标识，通过itemId可批量查询任务的状态（成功/失败）。'}
  ItemId *string `json:"itemId,omitempty" xml:"itemId,omitempty" require:"true"`
}

func (s PurgeFileWithTagResponse) String() string {
  return tea.Prettify(s)
}

func (s PurgeFileWithTagResponse) GoString() string {
  return s.String()
}

func (s *PurgeFileWithTagResponse) SetCode(v string) *PurgeFileWithTagResponse {
  s.Code = &v
  return s
}

func (s *PurgeFileWithTagResponse) SetMessage(v string) *PurgeFileWithTagResponse {
  s.Message = &v
  return s
}

func (s *PurgeFileWithTagResponse) SetItemId(v string) *PurgeFileWithTagResponse {
  s.ItemId = &v
  return s
}

type PurgeFileWithTagPaths struct {
}

func (s PurgeFileWithTagPaths) String() string {
  return tea.Prettify(s)
}

func (s PurgeFileWithTagPaths) GoString() string {
  return s.String()
}

type PurgeFileWithTagParameters struct {
}

func (s PurgeFileWithTagParameters) String() string {
  return tea.Prettify(s)
}

func (s PurgeFileWithTagParameters) GoString() string {
  return s.String()
}

type PurgeFileWithTagRequestHeader struct {
}

func (s PurgeFileWithTagRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s PurgeFileWithTagRequestHeader) GoString() string {
  return s.String()
}

type PurgeFileWithTagResponseHeader struct {
}

func (s PurgeFileWithTagResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s PurgeFileWithTagResponseHeader) GoString() string {
  return s.String()
}




type QueryPurgeLimitRequest struct {
}

func (s QueryPurgeLimitRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryPurgeLimitRequest) GoString() string {
  return s.String()
}

type QueryPurgeLimitResponse struct {
  // {'en':'url purge daily limit','zh_CN':'url 每日上限'}
  UrlUpper *int32 `json:"urlUpper,omitempty" xml:"urlUpper,omitempty" require:"true"`
  // {'en':'url purge remain','zh_CN':'url每日剩余量'}
  UrlRemain *int32 `json:"urlRemain,omitempty" xml:"urlRemain,omitempty" require:"true"`
  // {'en':'dir purge daily limit','zh_CN':'目录每日上限'}
  DirUpper *int32 `json:"dirUpper,omitempty" xml:"dirUpper,omitempty" require:"true"`
  // {'en':'dir purge remain','zh_CN':'目录每日剩余量'}
  DirRemain *int32 `json:"dirRemain,omitempty" xml:"dirRemain,omitempty" require:"true"`
  // {'en':'regex purge daily limit','zh_CN':'正则每日上限'}
  RegexUpper *int32 `json:"regexUpper,omitempty" xml:"regexUpper,omitempty" require:"true"`
  // {'en':'regex purge remain','zh_CN':'正则每日剩余量'}
  RegexRemain *int32 `json:"regexRemain,omitempty" xml:"regexRemain,omitempty" require:"true"`
  // {'en':'tag purge daily limit','zh_CN':'tag每日上限'}
  TagUpper *int32 `json:"tagUpper,omitempty" xml:"tagUpper,omitempty" require:"true"`
  // {'en':'tag purge remain','zh_CN':'tag每日剩余量'}
  TagRemain *int32 `json:"tagRemain,omitempty" xml:"tagRemain,omitempty" require:"true"`
  // {'en':'error code','zh_CN':'错误码'}
  Code *int32 `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {'en':'message','zh_CN':'错误信息'}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {'en':'supplier','zh_CN':'供应方'}
  Supplier *string `json:"supplier,omitempty" xml:"supplier,omitempty" require:"true"`
}

func (s QueryPurgeLimitResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryPurgeLimitResponse) GoString() string {
  return s.String()
}

func (s *QueryPurgeLimitResponse) SetUrlUpper(v int32) *QueryPurgeLimitResponse {
  s.UrlUpper = &v
  return s
}

func (s *QueryPurgeLimitResponse) SetUrlRemain(v int32) *QueryPurgeLimitResponse {
  s.UrlRemain = &v
  return s
}

func (s *QueryPurgeLimitResponse) SetDirUpper(v int32) *QueryPurgeLimitResponse {
  s.DirUpper = &v
  return s
}

func (s *QueryPurgeLimitResponse) SetDirRemain(v int32) *QueryPurgeLimitResponse {
  s.DirRemain = &v
  return s
}

func (s *QueryPurgeLimitResponse) SetRegexUpper(v int32) *QueryPurgeLimitResponse {
  s.RegexUpper = &v
  return s
}

func (s *QueryPurgeLimitResponse) SetRegexRemain(v int32) *QueryPurgeLimitResponse {
  s.RegexRemain = &v
  return s
}

func (s *QueryPurgeLimitResponse) SetTagUpper(v int32) *QueryPurgeLimitResponse {
  s.TagUpper = &v
  return s
}

func (s *QueryPurgeLimitResponse) SetTagRemain(v int32) *QueryPurgeLimitResponse {
  s.TagRemain = &v
  return s
}

func (s *QueryPurgeLimitResponse) SetCode(v int32) *QueryPurgeLimitResponse {
  s.Code = &v
  return s
}

func (s *QueryPurgeLimitResponse) SetMessage(v string) *QueryPurgeLimitResponse {
  s.Message = &v
  return s
}

func (s *QueryPurgeLimitResponse) SetSupplier(v string) *QueryPurgeLimitResponse {
  s.Supplier = &v
  return s
}

type QueryPurgeLimitPaths struct {
}

func (s QueryPurgeLimitPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryPurgeLimitPaths) GoString() string {
  return s.String()
}

type QueryPurgeLimitParameters struct {
}

func (s QueryPurgeLimitParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryPurgeLimitParameters) GoString() string {
  return s.String()
}

type QueryPurgeLimitRequestHeader struct {
}

func (s QueryPurgeLimitRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryPurgeLimitRequestHeader) GoString() string {
  return s.String()
}

type QueryPurgeLimitResponseHeader struct {
}

func (s QueryPurgeLimitResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryPurgeLimitResponseHeader) GoString() string {
  return s.String()
}




type GetPurgeStatusRequest struct {
  // {"en":"Start time of task creation,The format is yyyy-MM-dd HH:MM:ss; For example 2017-01-10 06:33:26,Can only query refresh tasks within 3 days.","zh_CN":"任务创建开始时间，格式为yyyy-MM-dd HH:mm:ss；例如 2017-01-10 06:33:26，只能查询3天之内的刷新任务。"}
  StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
  // {"en":"End time of task creation,The format is yyyy-MM-dd HH:MM:ss; For example 2017-01-10 06:33:26","zh_CN":"任务创建结束时间，格式为yyyy-MM-dd HH:mm:ss；例如 2017-01-10 06:33:26，"}
  EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
  // {"en":"A unique identifier for the same batch of tasks. If you submit multiple URLs from an API request, the ID is a unique number for these tasks.Query tasks by batch, such as submitting 10 url refreshes at a time. After successful submission, the content management system will return an itemId in the response message.","zh_CN":"表示任务单次提交多个url时任务的唯一标识。按批次查询任务，如单次提交10条url刷新，提交成功后内容管理系统将返回一个itemId在响应消息里。itemId 和 查询起止时间不能同时为空。"}
  ItemId *string `json:"itemId,omitempty" xml:"itemId,omitempty"`
  // {"en":"The cached url needs to be refreshed (url type could be directory, regular file, tag, and common file), and a single call allows only one url to be entered","zh_CN":"需要刷新缓存的url(支持目录、正则、tag、文件)，单次调用只允许输入一条url"}
  Url *string `json:"url,omitempty" xml:"url,omitempty"`
  // {"en":"Task status. The system allows you to select a task status query. These states can be queried:\nsuccess: Purge success.\nfailure: Purge failed.\nwait: The purge task is waiting to be processed.\nrun: The purge task is being executed.","zh_CN":"任务状态，允许指定任务状态过滤，支持查询的状态有：\n1）success\n2）failure\n3）wait\n4）run","exampleValue":"success,failure,wait,run"}
  Status *string `json:"status,omitempty" xml:"status,omitempty"`
  // {"defaultValue":"1","en":"Request page number. The default is 1.","zh_CN":"请求页数，缺省情况下，默认为1"}
  PageNo *int `json:"pageNo,omitempty" xml:"pageNo,omitempty"`
  // {"defaultValue":"20","en":"The number of pages displayed. The default is 20.","zh_CN":"每页显示的条数，缺省情况下，默认值为20"}
  PageSize *int `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s GetPurgeStatusRequest) String() string {
  return tea.Prettify(s)
}

func (s GetPurgeStatusRequest) GoString() string {
  return s.String()
}

func (s *GetPurgeStatusRequest) SetStartTime(v string) *GetPurgeStatusRequest {
  s.StartTime = &v
  return s
}

func (s *GetPurgeStatusRequest) SetEndTime(v string) *GetPurgeStatusRequest {
  s.EndTime = &v
  return s
}

func (s *GetPurgeStatusRequest) SetItemId(v string) *GetPurgeStatusRequest {
  s.ItemId = &v
  return s
}

func (s *GetPurgeStatusRequest) SetUrl(v string) *GetPurgeStatusRequest {
  s.Url = &v
  return s
}

func (s *GetPurgeStatusRequest) SetStatus(v string) *GetPurgeStatusRequest {
  s.Status = &v
  return s
}

func (s *GetPurgeStatusRequest) SetPageNo(v int) *GetPurgeStatusRequest {
  s.PageNo = &v
  return s
}

func (s *GetPurgeStatusRequest) SetPageSize(v int) *GetPurgeStatusRequest {
  s.PageSize = &v
  return s
}

type GetPurgeStatusRequestHeader struct {
}

func (s GetPurgeStatusRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetPurgeStatusRequestHeader) GoString() string {
  return s.String()
}

type GetPurgeStatusPaths struct {
}

func (s GetPurgeStatusPaths) String() string {
  return tea.Prettify(s)
}

func (s GetPurgeStatusPaths) GoString() string {
  return s.String()
}

type GetPurgeStatusParameters struct {
}

func (s GetPurgeStatusParameters) String() string {
  return tea.Prettify(s)
}

func (s GetPurgeStatusParameters) GoString() string {
  return s.String()
}

type GetPurgeStatusResponse struct {
  // {"en":"The number of results for this query, count for 10 if 10 tasks are eligible for a query","zh_CN":"本次查询结果的数量，如有10个任务符合查询条件，则count的值为10"}
  Count *int `json:"count,omitempty" xml:"count,omitempty" require:"true"`
  // {"en":"After the task commits, the system's response code:\n1) 0 means failure \n2) 1 means success","zh_CN":"任务提交后，系统的响应:\n1) 0表示失败\n2) 1表示成功","exampleValue":"0,1"}
  Code *int `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
  // {"en":"System response message after task submit.","zh_CN":"任务提交后，系统的响应消息"}
  Message *string `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
  // {"en":"Total page count of task query results","zh_CN":"任务查询结果的总页数"}
  PageNo *int `json:"pageNo,omitempty" xml:"pageNo,omitempty" require:"true"`
  // {"en":"How many refresh task data per page","zh_CN":"每页显示多少条刷新任务数据"}
  PageSize *int `json:"pageSize,omitempty" xml:"pageSize,omitempty" require:"true"`
  // {"en":"Set of task results","zh_CN":"任务结果的集合"}
  ResultDetail []*GetPurgeStatusResponseResultDetail `json:"resultDetail,omitempty" xml:"resultDetail,omitempty" require:"true" type:"Repeated"`
}

func (s GetPurgeStatusResponse) String() string {
  return tea.Prettify(s)
}

func (s GetPurgeStatusResponse) GoString() string {
  return s.String()
}

func (s *GetPurgeStatusResponse) SetCount(v int) *GetPurgeStatusResponse {
  s.Count = &v
  return s
}

func (s *GetPurgeStatusResponse) SetCode(v int) *GetPurgeStatusResponse {
  s.Code = &v
  return s
}

func (s *GetPurgeStatusResponse) SetMessage(v string) *GetPurgeStatusResponse {
  s.Message = &v
  return s
}

func (s *GetPurgeStatusResponse) SetPageNo(v int) *GetPurgeStatusResponse {
  s.PageNo = &v
  return s
}

func (s *GetPurgeStatusResponse) SetPageSize(v int) *GetPurgeStatusResponse {
  s.PageSize = &v
  return s
}

func (s *GetPurgeStatusResponse) SetResultDetail(v []*GetPurgeStatusResponseResultDetail) *GetPurgeStatusResponse {
  s.ResultDetail = v
  return s
}

type GetPurgeStatusResponseResultDetail struct     {
  // {"en":"Time for content management system to start cache refresh task","zh_CN":"内容管理系统开始执行缓存刷新任务的时间"}
  BeginTime *string `json:"beginTime,omitempty" xml:"beginTime,omitempty" require:"true"`
  // {"en":"Time the content management system creates a cache refresh task","zh_CN":"内容管理系统创建缓存刷新任务的时间"}
  CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty" require:"true"`
  // {"en":"Content management system executes and summarizes the completion time of cache refresh tasks","zh_CN":"内容管理系统执行并汇总缓存刷新任务的完成时间"}
  FinishTime *string `json:"finishTime,omitempty" xml:"finishTime,omitempty" require:"true"`
  // {"en":"Success rate for performing cache refresh tasks, such as 98%, the value is 98","zh_CN":"执行缓存刷新任务的成功率，如98%，则值为98"}
  Rate *string `json:"rate,omitempty" xml:"rate,omitempty" require:"true"`
  // {"en":"Cache refresh task execution status, which has the following states:\n1) success: Indicates that the task that refreshes the file cache has performed successfully \n2) failure:  Indicates that the task that refreshes the file cache has performed failed \n3) wait: Indicates that the task of refreshing the cache is in the queue \n4) run: Indicates that the task of refreshing the cache is in progress","zh_CN":"缓存刷新任务执行的状态，有以下几种状态：success：表示刷新文件缓存的任务执行成功 failure：表示刷新文件缓存的任务执行失败 wait：表示刷新缓存的任务正在排队中 \nrun：表示刷新缓存的任务正在执行中","exampleValue":"success,failure,wait,run"}
  Status *string `json:"status,omitempty" xml:"status,omitempty" require:"true"`
  // {"en":"Refresh the cache of a file or directory.","zh_CN":"执行缓存刷新的具体文件或目录"}
  Url *string `json:"url,omitempty" xml:"url,omitempty" require:"true"`
  // {"en":"The operation type of the purge task:\n1) 0: Refresh a specific file \n2) 1:  regular file \n3) 2: Refresh all files in a directory \n4) 3: Refresh  a tag","zh_CN":"刷新任务的操作类型：\n1) 0：刷新某个具体文件 \n2) 1：正则文件 \n3) 2：刷新某个目录下的所有文件 \n4) 3：刷新某个tag","exampleValue":"0,1,2,3"}
  IsDir *int `json:"isDir,omitempty" xml:"isDir,omitempty" require:"true"`
}

func (s GetPurgeStatusResponseResultDetail) String() string {
  return tea.Prettify(s)
}

func (s GetPurgeStatusResponseResultDetail) GoString() string {
  return s.String()
}

func (s *GetPurgeStatusResponseResultDetail) SetBeginTime(v string) *GetPurgeStatusResponseResultDetail {
  s.BeginTime = &v
  return s
}

func (s *GetPurgeStatusResponseResultDetail) SetCreateTime(v string) *GetPurgeStatusResponseResultDetail {
  s.CreateTime = &v
  return s
}

func (s *GetPurgeStatusResponseResultDetail) SetFinishTime(v string) *GetPurgeStatusResponseResultDetail {
  s.FinishTime = &v
  return s
}

func (s *GetPurgeStatusResponseResultDetail) SetRate(v string) *GetPurgeStatusResponseResultDetail {
  s.Rate = &v
  return s
}

func (s *GetPurgeStatusResponseResultDetail) SetStatus(v string) *GetPurgeStatusResponseResultDetail {
  s.Status = &v
  return s
}

func (s *GetPurgeStatusResponseResultDetail) SetUrl(v string) *GetPurgeStatusResponseResultDetail {
  s.Url = &v
  return s
}

func (s *GetPurgeStatusResponseResultDetail) SetIsDir(v int) *GetPurgeStatusResponseResultDetail {
  s.IsDir = &v
  return s
}

type GetPurgeStatusResponseHeader struct {
}

func (s GetPurgeStatusResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetPurgeStatusResponseHeader) GoString() string {
  return s.String()
}




