package domainconfigurtion

// This file is auto-generated, don't edit it. Thanks.
import (
  "github.com/alibabacloud-go/tea/tea"
)

type EditAccessSpeedLimitRequest struct {
  // {"en":"Access speed limit configuration
  // 1. This item is required when access speed limit is required
  // 2. Clear the configuration for blank access-speed-rules", "zh_CN":"访问限速配置
  // 1.需要访问限速时，此项必填
  // 2.只有空标签access-speed-rules时清空访问限速配置"}
  AccessSpeedRules []*EditAccessSpeedLimitRequestAccessSpeedRules `json:"access-speed-rules,omitempty" xml:"access-speed-rules,omitempty" require:"true" type:"Repeated"`
}

func (s EditAccessSpeedLimitRequest) String() string {
  return tea.Prettify(s)
}

func (s EditAccessSpeedLimitRequest) GoString() string {
  return s.String()
}

func (s *EditAccessSpeedLimitRequest) SetAccessSpeedRules(v []*EditAccessSpeedLimitRequestAccessSpeedRules) *EditAccessSpeedLimitRequest {
  s.AccessSpeedRules = v
  return s
}

type EditAccessSpeedLimitRequestAccessSpeedRules struct     {
  // {"en":"The url matching mode supports fuzzy regularization. If all matches, the input parameters can be configured as: *.", "zh_CN":"url匹配模式，支持正则，如果是全部匹配，入参可以配置为：.*"}
  PathPattern *string `json:"path-pattern,omitempty" xml:"path-pattern,omitempty" require:"true"`
  // {"en":"The speed limit method can be limited by the start size (unit: KB) or the start time (unit: s), both of which are mutually exclusive
  // Allowed values: size/time.", "zh_CN":"限速方式，可按开始大小（单位KB）或开始时间（单位：s）限速，两者互斥
  // 允许值：size/time"}
  LimitMode *string `json:"limit-mode,omitempty" xml:"limit-mode,omitempty" require:"true"`
  // {"en":"Specify the number of bytes to start limiting, in KB, which means that the limit starts after a certain number of bytes, and 0 means no limit. When limit-mode is size, this item is required, and start-time is cleared.", "zh_CN":"开始限制字节，单位KB，从多少字节后开始限制，0表示不限制
  // limit-mode为size时，此项必填，start-time清空"}
  StartSize *string `json:"start-size,omitempty" xml:"start-size,omitempty"`
  // {"en":"Specifies how long the speed limit starts after the connection is established, the unit is S, 0 means no limit.
  // When the limit-mode is time, this item is required, and the start-size is cleared.", "zh_CN":"开始限制时间，单位S，从建连后多长时间后开始限制，0表示不限制
  // limit-mode为time时，此项必填，start-size清空"}
  StartTime *string `json:"start-time,omitempty" xml:"start-time,omitempty"`
  // {"en":"The rate before starting to limit, unit KB/S.", "zh_CN":"开始限制之前的速率，单位KB/S"}
  StartSpeed *string `json:"start-speed,omitempty" xml:"start-speed,omitempty"`
  // {"en":"fill in the integer, if match multiple entries, Large number has a high priority", "zh_CN":"优先级，填写整数，当配置多条时数字大的优先"}
  Priority *string `json:"priority,omitempty" xml:"priority,omitempty"`
  // {"en":"Speed limit value, refers to the speed after the limit, stable speed, unit KB/S, -1 means no limit.", "zh_CN":"限速值，指限制之后的速率，稳定速率，单位KB/S"}
  Speed *string `json:"speed,omitempty" xml:"speed,omitempty" require:"true"`
  // {"en":"DataId is to indicate a specific group configuration when the client has multiple groups of configurations. dataId can be retrieved through a query interface. Note: A. If dataId is passed, it means that one group of configuration items is specified to be modified, and no other group configuration items need to be modified. B. If multiple groups of configurations are included, some of them are configured with dataId and others are not, then the expression of dataId is used to modify a specific group of configurations, and a new group of configurations is added on the original basis without the expression of dataId. C. If the dataId is not transmitted, it means that the original configuration will be fully covered by this configuration. D. If no configuration parameter is passed, only domain name and secondary label are passed, which means that all configuration of domain name secondary service corresponding to this interface is cleared. E. If there is no specific configuration item in a set of configurations, the dataId must be filled in, and the value is the actual dataId, which means clearing the value of the corresponding dataId configuration item; it is not allowed that there is no specific configuration item or dataId in a set of configurations.", "zh_CN":"配置多组配置时，具体某组配置的id。本功能只支持一组配置。dataId可以通过查询接口获取。 注意： a、如果有传dataId，说明指定修改该组配置项内容； b、如果入参都没有传dataId,表示用本次的配置全量覆盖原先配置； c、如果一组配置没有具体的配置项，则dataId必填，且值为实际存在的dataId，表示清空这个dataId对应配置项的值；不允许一组配置没有具体的配置项也没有dataId。"}
  DataId *int64 `json:"data-id,omitempty" xml:"data-id,omitempty"`
}

func (s EditAccessSpeedLimitRequestAccessSpeedRules) String() string {
  return tea.Prettify(s)
}

func (s EditAccessSpeedLimitRequestAccessSpeedRules) GoString() string {
  return s.String()
}

func (s *EditAccessSpeedLimitRequestAccessSpeedRules) SetPathPattern(v string) *EditAccessSpeedLimitRequestAccessSpeedRules {
  s.PathPattern = &v
  return s
}

func (s *EditAccessSpeedLimitRequestAccessSpeedRules) SetLimitMode(v string) *EditAccessSpeedLimitRequestAccessSpeedRules {
  s.LimitMode = &v
  return s
}

func (s *EditAccessSpeedLimitRequestAccessSpeedRules) SetStartSize(v string) *EditAccessSpeedLimitRequestAccessSpeedRules {
  s.StartSize = &v
  return s
}

func (s *EditAccessSpeedLimitRequestAccessSpeedRules) SetStartTime(v string) *EditAccessSpeedLimitRequestAccessSpeedRules {
  s.StartTime = &v
  return s
}

func (s *EditAccessSpeedLimitRequestAccessSpeedRules) SetStartSpeed(v string) *EditAccessSpeedLimitRequestAccessSpeedRules {
  s.StartSpeed = &v
  return s
}

func (s *EditAccessSpeedLimitRequestAccessSpeedRules) SetPriority(v string) *EditAccessSpeedLimitRequestAccessSpeedRules {
  s.Priority = &v
  return s
}

func (s *EditAccessSpeedLimitRequestAccessSpeedRules) SetSpeed(v string) *EditAccessSpeedLimitRequestAccessSpeedRules {
  s.Speed = &v
  return s
}

func (s *EditAccessSpeedLimitRequestAccessSpeedRules) SetDataId(v int64) *EditAccessSpeedLimitRequestAccessSpeedRules {
  s.DataId = &v
  return s
}

type EditAccessSpeedLimitResponse struct {
  // {"en":"Error code, which appears when HTTPStatus is not 202, represents the error type of the current request call", "zh_CN":"错误代码，当HTTPStatus不为202时出现，表示当前请求调用的错误类型"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response message, and shows as success when it is successful.", "zh_CN":"响应信息，成功时为success"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s EditAccessSpeedLimitResponse) String() string {
  return tea.Prettify(s)
}

func (s EditAccessSpeedLimitResponse) GoString() string {
  return s.String()
}

func (s *EditAccessSpeedLimitResponse) SetCode(v string) *EditAccessSpeedLimitResponse {
  s.Code = &v
  return s
}

func (s *EditAccessSpeedLimitResponse) SetMessage(v string) *EditAccessSpeedLimitResponse {
  s.Message = &v
  return s
}

type EditAccessSpeedLimitPaths struct {
  // {"en":"The domain you want to update, support domain id and domain name.", "zh_CN":"需要查询配置的域名或域名id"}
  DomainName *string `json:"domain-name,omitempty" xml:"domain-name,omitempty" require:"true"`
}

func (s EditAccessSpeedLimitPaths) String() string {
  return tea.Prettify(s)
}

func (s EditAccessSpeedLimitPaths) GoString() string {
  return s.String()
}

func (s *EditAccessSpeedLimitPaths) SetDomainName(v string) *EditAccessSpeedLimitPaths {
  s.DomainName = &v
  return s
}

type EditAccessSpeedLimitParameters struct {
}

func (s EditAccessSpeedLimitParameters) String() string {
  return tea.Prettify(s)
}

func (s EditAccessSpeedLimitParameters) GoString() string {
  return s.String()
}

type EditAccessSpeedLimitRequestHeader struct {
}

func (s EditAccessSpeedLimitRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s EditAccessSpeedLimitRequestHeader) GoString() string {
  return s.String()
}

type EditAccessSpeedLimitResponseHeader struct {
  // {"en":"httpstatus=202; Indicates that the new domain API was successfully invoked, and the current deployment of the new domain can be viewed using x-cnc-request-id in the header", "zh_CN":"httpstatus=202;   表示成功调用新增域名接口，可使用header中的x-cnc-request-id查看当前新增域名的部署情况"}
  HttpStatus *int `json:"http status code,omitempty" xml:"http status code,omitempty" require:"true"`
  // {"en":"Uniquely identified id for querying tasks per request (for all API)", "zh_CN":"唯一标示的id，用于查询每次请求的任务 （适用全部接口）"}
  XCncRequestId *string `json:"x-cnc-request-id,omitempty" xml:"x-cnc-request-id,omitempty" require:"true"`
}

func (s EditAccessSpeedLimitResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s EditAccessSpeedLimitResponseHeader) GoString() string {
  return s.String()
}

func (s *EditAccessSpeedLimitResponseHeader) SetHttpStatus(v int) *EditAccessSpeedLimitResponseHeader {
  s.HttpStatus = &v
  return s
}

func (s *EditAccessSpeedLimitResponseHeader) SetXCncRequestId(v string) *EditAccessSpeedLimitResponseHeader {
  s.XCncRequestId = &v
  return s
}




type EnableOrDisableDMSprotectionRequest struct {
  // {"en":"Domain names list, the parent tag.", "zh_CN":"开启/关闭Dms防护的域名列表， 父标签"}
  DomainNames []*string `json:"domainNames,omitempty" xml:"domainNames,omitempty" require:"true" type:"Repeated"`
  // {"en":"1: Enable Dms protection; 0: Disable Dms  protection", "zh_CN":"1：开启Dms防护，0：关闭Dms防护"}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
}

func (s EnableOrDisableDMSprotectionRequest) String() string {
  return tea.Prettify(s)
}

func (s EnableOrDisableDMSprotectionRequest) GoString() string {
  return s.String()
}

func (s *EnableOrDisableDMSprotectionRequest) SetDomainNames(v []*string) *EnableOrDisableDMSprotectionRequest {
  s.DomainNames = v
  return s
}

func (s *EnableOrDisableDMSprotectionRequest) SetType(v string) *EnableOrDisableDMSprotectionRequest {
  s.Type = &v
  return s
}

type EnableOrDisableDMSprotectionResponse struct {
  // {"en":"httpstatus=202; Indicates that the new domain API was successfully invoked, and the current deployment of the new domain can be viewed using x-cnc-request-id in the header", "zh_CN":"httpstatus=202;   表示成功调用新增域名接口，可使用header中的x-cnc-request-id查看当前新增域名的部署情况"}
  HttpStatus *int `json:"http status code,omitempty" xml:"http status code,omitempty" require:"true"`
  // {"en":"Uniquely identified id for querying tasks per request (for all API)", "zh_CN":"唯一标示的id，用于查询每次请求的任务 （适用全部接口）"}
  XCncRequestId *string `json:"x-cnc-request-id,omitempty" xml:"x-cnc-request-id,omitempty" require:"true"`
  // {"en":"Error code. 0 message successful", "zh_CN":"错误代码。 0：成功"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"The response message. Response success when calling API successful", "zh_CN":"响应信息，成功时为success"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"The body of return data.", "zh_CN":"返回体"}
  Data *string `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s EnableOrDisableDMSprotectionResponse) String() string {
  return tea.Prettify(s)
}

func (s EnableOrDisableDMSprotectionResponse) GoString() string {
  return s.String()
}

func (s *EnableOrDisableDMSprotectionResponse) SetHttpStatus(v int) *EnableOrDisableDMSprotectionResponse {
  s.HttpStatus = &v
  return s
}

func (s *EnableOrDisableDMSprotectionResponse) SetXCncRequestId(v string) *EnableOrDisableDMSprotectionResponse {
  s.XCncRequestId = &v
  return s
}

func (s *EnableOrDisableDMSprotectionResponse) SetCode(v string) *EnableOrDisableDMSprotectionResponse {
  s.Code = &v
  return s
}

func (s *EnableOrDisableDMSprotectionResponse) SetMessage(v string) *EnableOrDisableDMSprotectionResponse {
  s.Message = &v
  return s
}

func (s *EnableOrDisableDMSprotectionResponse) SetData(v string) *EnableOrDisableDMSprotectionResponse {
  s.Data = &v
  return s
}

type EnableOrDisableDMSprotectionPaths struct {
}

func (s EnableOrDisableDMSprotectionPaths) String() string {
  return tea.Prettify(s)
}

func (s EnableOrDisableDMSprotectionPaths) GoString() string {
  return s.String()
}

type EnableOrDisableDMSprotectionParameters struct {
}

func (s EnableOrDisableDMSprotectionParameters) String() string {
  return tea.Prettify(s)
}

func (s EnableOrDisableDMSprotectionParameters) GoString() string {
  return s.String()
}

type EnableOrDisableDMSprotectionRequestHeader struct {
}

func (s EnableOrDisableDMSprotectionRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s EnableOrDisableDMSprotectionRequestHeader) GoString() string {
  return s.String()
}

type EnableOrDisableDMSprotectionResponseHeader struct {
}

func (s EnableOrDisableDMSprotectionResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s EnableOrDisableDMSprotectionResponseHeader) GoString() string {
  return s.String()
}




type EditHttpHeaderConfigRequest struct {
  // {"en":"Http header settings
  // note:
  // 1. When you need to cancel the http header setting, you can pass in the empty node <header-modify-rules></header-modify-rules>.
  // 2. indicating that you need to set the http header, this field is required", "zh_CN":"http头设置
  // 注意：
  // 1. 需要取消http头设置时，可以传入空节点<header-modify-rules></header-modify-rules>。
  // 2. 表示需要设置http头，此项必填"}
  HeaderModifyRules []*EditHttpHeaderConfigEditHttpHeaderConfigRequestHeaderModifyRules `json:"header-modify-rules,omitempty" xml:"header-modify-rules,omitempty" require:"true" type:"Repeated"`
}

func (s EditHttpHeaderConfigRequest) String() string {
  return tea.Prettify(s)
}

func (s EditHttpHeaderConfigRequest) GoString() string {
  return s.String()
}

func (s *EditHttpHeaderConfigRequest) SetHeaderModifyRules(v []*EditHttpHeaderConfigEditHttpHeaderConfigRequestHeaderModifyRules) *EditHttpHeaderConfigRequest {
  s.HeaderModifyRules = v
  return s
}

type EditHttpHeaderConfigEditHttpHeaderConfigRequestHeaderModifyRules struct     {
  // {"en":"Add a grid type identifier to indicate a specific group configuration when the client has multiple groups of configurations.", "zh_CN":"添加grid类型标识，表示客户多组配置时，具体某组配置。
  // data-id可以通过查询接口获取。
  // 注意：添加grid类型标识：data-id，每一组配置对应一个data-id：
  // a、如果客户有传data-id，说明指定修改其中一组配置项内容，不需求修改其他组配置内容不需要入参；
  // b、如果客户入参多组配置，其中有些组配置有传data-id，有些没有传，则有传data-id的表示修改具体某组配置，没有传data-id的表示在原来基础上新增一组配置；
  // c、如果客户入参都没有传data-id,表示用本次的配置全量覆盖原先配置；
  // d、如果客户入参没有传任何配置项参数，只传了域名和二级标签，表示清空这个接口对应域名二级服务所有配置。（c、d内容和当前方案实现一致）；
  // e、一个gird标签下的入参不能为空，如果，没有具体的配置项，则data-id必填，且值为实际存在的data-id,表示清空这个data-id对应配置项的值；"}
  DataId *int64 `json:"data-id,omitempty" xml:"data-id,omitempty"`
  // {"en":"The url matching mode supports fuzzy regularization. If all matches, the input parameters can be configured as: *", "zh_CN":"url匹配模式，支持正则，如果是全部匹配，入参可以配置为：.*"}
  PathPattern *string `json:"path-pattern,omitempty" xml:"path-pattern,omitempty"`
  // {"en":"Exception url matching pattern, support regular. Example: ", "zh_CN":"例外的url匹配模式，支持正则。 入参参考："}
  ExceptPathPattern *string `json:"except-path-pattern,omitempty" xml:"except-path-pattern,omitempty"`
  // {"en":"Matching conditions: specify common types, optional values are all or homepage. 1. all: all files 2. homepage: home page", "zh_CN":"匹配条件：指定常用类型，可选值为all或homepage 1. all：全部文件 2. homepage：首页"}
  CustomPattern *string `json:"custom-pattern,omitempty" xml:"custom-pattern,omitempty"`
  // {"en":"Matching conditions: file type, please separate by semicolon, optional values: gif png bmp jpeg jpg html htm shtml mp3 wma flv mp4 wmv zip exe rar css txt ico js swf m3u8 xml f4m bootstarp ts.", "zh_CN":"匹配条件：文件类型，多个请以英文;分隔，可选值：gif png bmp jpeg jpg html htm shtml mp3 wma flv mp4 wmv zip exe rar css txt ico js swf m3u8 xml f4m bootstarp ts"}
  FileType *string `json:"file-type,omitempty" xml:"file-type,omitempty"`
  // {"en":"Matching condition: Custom file type, separate by semicolon.", "zh_CN":"匹配条件：自定义文件类型，多个请以英文分号分隔。"}
  CustomFileType *string `json:"custom-file-type,omitempty" xml:"custom-file-type,omitempty"`
  // {"en":"Directory", "zh_CN":"目录"}
  Directory *string `json:"directory,omitempty" xml:"directory,omitempty"`
  // {"en":"Matching Condition: Specify URL.
  // The input parameter does not support the URI format starting with http(s)://", "zh_CN":"匹配条件：指定URL
  // 入参不支持含http(s):// 开头的URI格式"}
  SpecifyUrl *string `json:"specify-url,omitempty" xml:"specify-url,omitempty"`
  // {"en":"The matching request method, the optional values are: GET, POST, PUT, HEAD, DELETE, OPTIONS, separate by semicolons.", "zh_CN":"匹配的请求方式，可选值为：GET、POST、PUT、HEAD、DELETE、OPTIONS，多个请以英文分号分隔"}
  RequestMethod *string `json:"request-method,omitempty" xml:"request-method,omitempty"`
  // {"en":"The control direction of the http header, the optional value is cache2visitor/cache2origin/visitor2cache/origin2cache, single-select.
  // Cache2origin refers to the source direction---corresponding to the configuration item return source request;
  // Cache2visitor refers to the direction of the client back - the corresponding configuration item returns to the client response;
  // Visitor2cache refers to receiving client requests
  // Origin2cache refers to the receiving source response", "zh_CN":"http头的控制方向，可选值为cache2visitor/cache2origin/visitor2cache/origin2cache，单选。
  // cache2origin是指回源方向---对应配置项回源请求；
  // cache2visitor是指回客户端方向—对应配置项回客户端应答；
  // visitor2cache是指接收客户端请求
  // origin2cache是指接收源应答
  // 配置接收源应答方向，添加非CACHE control头，无法传递给客户端"}
  HeaderDirection *string `json:"header-direction,omitempty" xml:"header-direction,omitempty"`
  // {"en":"The control type of the http header supports the addition and deletion of the http header value. The optional value is add|set|delete, which is single-selected. Corresponding to the header-name and header-value parameters.
  // 1. Add: add a header
  // 2. Set: modify the header value
  // 3. Delete: delete the header
  // Note: priority is delete > set > add", "zh_CN":"http头的控制类型，支持http头部的增删改，可选值为add|set|delete，单选。对应header-name、header-value参数
  // 1. add：表示新增一个头部，头部名称为header-name，头部值为header-value
  // 2. set：表示修改指定头部header-name的值为header-value
  // 3. delete：表示删除头部，header-name可同时配置多个
  // 注意：优先级delete>set>add。当源站有对应响应头，则按源站响应的头部响应给客户端，此处新增的无效。"}
  Action *string `json:"action,omitempty" xml:"action,omitempty"`
  // {"en":"Http header regular match, optional value: true / false.
  // True: indicates that the value of the header-name is handled as a regular match.
  // False: indicates that the value of the header-name is processed according to the actual parameters, and no regular match is made.
  // Do not pass the default is false", "zh_CN":"http头正则匹配，可选值：true/false。
  // true：表示对header-name的值按正则匹配方式处理
  // false:表示对header-name的值按实际入参处理，不做正则匹配。
  // 不传默认是false"}
  AllowRegexp *string `json:"allow-regexp,omitempty" xml:"allow-regexp,omitempty"`
  // {"en":"Http header name, add or modify the http header, only one is allowed; delete the http header to allow multiple entries, separated by a semicolon ';'.
  // Note: The operation of the special http header is limited, and the http header and operation type of the operation are allowed.
  // This item is required and cannot be empty
  // When the action is add: indicates that the header-name header is added.
  // When the action is set: modify the header-name header
  // When the action is delete: delete the header-name header", "zh_CN":"http头名称，新增或修改http头，只允许输入一个；删除http头允许输入多个，以分号“;”隔开。
  // 1.当action为add：表示新增这个header-name头部
  // 2.当action为set：修改这个header-name头部的值
  // 3.当action为delete：删除这个header-name头部
  //   
  // 注意：对特殊http头的操作是受限的，允许操作的http头及操作类型请参看【概览】-【附件2： header操作】"}
  HeaderName *string `json:"header-name,omitempty" xml:"header-name,omitempty"`
  // {"en":"The value corresponding to the HTTP header field, for example: mytest.example.com
  // Note:
  // 1. When the action is add or set, the input parameter must be passed a value
  // 2. When the action is delete, the input parameter is not passed
  // Support to get the value of specified variable by keyword, such as client IP, including:
  // Key words: meaning
  // #timestamp: current time, timestamp as 1559124945
  // #request-host: host in the request header
  // #request-url: request url, which contains the full path of the protocol domain name, etc., such as http://aaa.aa.com/a.html
  // #request-uri: request uri, relative path format, such as /index.html
  // #origin- IP: return source IP
  // #cache-ip: edge node IP
  // #server-ip: external service IP
  // #client-ip: client IP, or visitor IP
  // #response-header{XXX} : get the value in the response header, such as #response-header{etag}, get the etag value in response-header 
  // #header{XXX} : to get the value in the HTTP header of the request, such as #header{user-agent}, is to get the user-agent value in the header
  // #cookie{XXX} : get the value in the cookie, such as #cookie{account}, is to get the value of the account set in the cookie", "zh_CN":"http头域对应的值，例如：mytest.example.com 
  // 注意：
  // 1. 当action为add或set时，该入参必须传值
  // 2. 当action为delete时，该入参不用传
  // 支持通过关键字获取指定变量值，如客户端ip，包含如下：
  // 关键字：含义
  // #timestamp：当前时间，时间戳如1559124945
  // #request-host：请求头中的HOST
  // #request-url：请求url，包含协议域名等的全路径，如http://aaa.aa.com/a.html
  // #request-uri：请求uri，相对路径格式，如/index.html
  // #origin-ip：回源IP
  // #cache-ip：边缘节点IP
  // #server-ip：对外服务IP
  // #client-ip：客户端IP，即访客IP
  // #response-header{xxx}：获取响应头中的值，如#response-header{etag}，获取response-header中的etag值
  // #header{xxx}：获取请求的http header中的值，如#header{User-Agent}，是获取header中的User-Agent值
  // #cookie{xxx}：获取cookie中的值，如#cookie{account}，是获取cookie中设置的account的值  "}
  HeaderValue *string `json:"header-value,omitempty" xml:"header-value,omitempty"`
  // {"en":"The original value corresponding to the HTTP header field","zh_CN":"http头域对应的原始值"}
  HeaderValueOld *string `json:"header-value-old,omitempty" xml:"header-value-old,omitempty"`
  // {"en":"Match request header, header values support regular, header and header values separated by Spaces, e.g. : Range bytes=[0-9]{9,}", "zh_CN":"2匹配请求头，头部值支持正则，头和头部值用空格隔开，如：Range bytes=[0-9]{9,}"}
  EditHttpHeaderConfigRequestHeader *string `json:"request-header,omitempty" xml:"request-header,omitempty"`
  // {"en":"Indicates the priority of execution order for multiple sets of configurations. A higher number indicates higher priority. If no parameters are passed, the default value is 10 and cannot be cleared.", "zh_CN":"表示客户多组配置的优先执行顺序。数字越大，优先级越高。 不传参默认为10，不可清空"}
  Priority *string `json:"priority,omitempty" xml:"priority,omitempty"`
  // {"en":"Exception file type.", "zh_CN":"例外的文件类型"}
  ExceptFileType *string `json:"except-file-type,omitempty" xml:"except-file-type,omitempty"`
  // {"en":"Exception directory.", "zh_CN":"例外的目录"}
  ExceptDirectory *string `json:"except-directory,omitempty" xml:"except-directory,omitempty"`
  // {"en":"Exception request method.", "zh_CN":"例外的请求方式"}
  ExceptRequestMethod *string `json:"except-request-method,omitempty" xml:"except-request-method,omitempty"`
  // {"en":"Exception request header.", "zh_CN":"例外的请求头"}
  ExceptEditHttpHeaderConfigRequestHeader *string `json:"except-request-header,omitempty" xml:"except-request-header,omitempty"`
}

func (s EditHttpHeaderConfigEditHttpHeaderConfigRequestHeaderModifyRules) String() string {
  return tea.Prettify(s)
}

func (s EditHttpHeaderConfigEditHttpHeaderConfigRequestHeaderModifyRules) GoString() string {
  return s.String()
}

func (s *EditHttpHeaderConfigEditHttpHeaderConfigRequestHeaderModifyRules) SetDataId(v int64) *EditHttpHeaderConfigEditHttpHeaderConfigRequestHeaderModifyRules {
  s.DataId = &v
  return s
}

func (s *EditHttpHeaderConfigEditHttpHeaderConfigRequestHeaderModifyRules) SetPathPattern(v string) *EditHttpHeaderConfigEditHttpHeaderConfigRequestHeaderModifyRules {
  s.PathPattern = &v
  return s
}

func (s *EditHttpHeaderConfigEditHttpHeaderConfigRequestHeaderModifyRules) SetExceptPathPattern(v string) *EditHttpHeaderConfigEditHttpHeaderConfigRequestHeaderModifyRules {
  s.ExceptPathPattern = &v
  return s
}

func (s *EditHttpHeaderConfigEditHttpHeaderConfigRequestHeaderModifyRules) SetCustomPattern(v string) *EditHttpHeaderConfigEditHttpHeaderConfigRequestHeaderModifyRules {
  s.CustomPattern = &v
  return s
}

func (s *EditHttpHeaderConfigEditHttpHeaderConfigRequestHeaderModifyRules) SetFileType(v string) *EditHttpHeaderConfigEditHttpHeaderConfigRequestHeaderModifyRules {
  s.FileType = &v
  return s
}

func (s *EditHttpHeaderConfigEditHttpHeaderConfigRequestHeaderModifyRules) SetCustomFileType(v string) *EditHttpHeaderConfigEditHttpHeaderConfigRequestHeaderModifyRules {
  s.CustomFileType = &v
  return s
}

func (s *EditHttpHeaderConfigEditHttpHeaderConfigRequestHeaderModifyRules) SetDirectory(v string) *EditHttpHeaderConfigEditHttpHeaderConfigRequestHeaderModifyRules {
  s.Directory = &v
  return s
}

func (s *EditHttpHeaderConfigEditHttpHeaderConfigRequestHeaderModifyRules) SetSpecifyUrl(v string) *EditHttpHeaderConfigEditHttpHeaderConfigRequestHeaderModifyRules {
  s.SpecifyUrl = &v
  return s
}

func (s *EditHttpHeaderConfigEditHttpHeaderConfigRequestHeaderModifyRules) SetRequestMethod(v string) *EditHttpHeaderConfigEditHttpHeaderConfigRequestHeaderModifyRules {
  s.RequestMethod = &v
  return s
}

func (s *EditHttpHeaderConfigEditHttpHeaderConfigRequestHeaderModifyRules) SetHeaderDirection(v string) *EditHttpHeaderConfigEditHttpHeaderConfigRequestHeaderModifyRules {
  s.HeaderDirection = &v
  return s
}

func (s *EditHttpHeaderConfigEditHttpHeaderConfigRequestHeaderModifyRules) SetAction(v string) *EditHttpHeaderConfigEditHttpHeaderConfigRequestHeaderModifyRules {
  s.Action = &v
  return s
}

func (s *EditHttpHeaderConfigEditHttpHeaderConfigRequestHeaderModifyRules) SetAllowRegexp(v string) *EditHttpHeaderConfigEditHttpHeaderConfigRequestHeaderModifyRules {
  s.AllowRegexp = &v
  return s
}

func (s *EditHttpHeaderConfigEditHttpHeaderConfigRequestHeaderModifyRules) SetHeaderName(v string) *EditHttpHeaderConfigEditHttpHeaderConfigRequestHeaderModifyRules {
  s.HeaderName = &v
  return s
}

func (s *EditHttpHeaderConfigEditHttpHeaderConfigRequestHeaderModifyRules) SetHeaderValue(v string) *EditHttpHeaderConfigEditHttpHeaderConfigRequestHeaderModifyRules {
  s.HeaderValue = &v
  return s
}

func (s *EditHttpHeaderConfigEditHttpHeaderConfigRequestHeaderModifyRules) SetHeaderValueOld(v string) *EditHttpHeaderConfigEditHttpHeaderConfigRequestHeaderModifyRules {
  s.HeaderValueOld = &v
  return s
}

func (s *EditHttpHeaderConfigEditHttpHeaderConfigRequestHeaderModifyRules) SetEditHttpHeaderConfigRequestHeader(v string) *EditHttpHeaderConfigEditHttpHeaderConfigRequestHeaderModifyRules {
  s.EditHttpHeaderConfigRequestHeader = &v
  return s
}

func (s *EditHttpHeaderConfigEditHttpHeaderConfigRequestHeaderModifyRules) SetPriority(v string) *EditHttpHeaderConfigEditHttpHeaderConfigRequestHeaderModifyRules {
  s.Priority = &v
  return s
}

func (s *EditHttpHeaderConfigEditHttpHeaderConfigRequestHeaderModifyRules) SetExceptFileType(v string) *EditHttpHeaderConfigEditHttpHeaderConfigRequestHeaderModifyRules {
  s.ExceptFileType = &v
  return s
}

func (s *EditHttpHeaderConfigEditHttpHeaderConfigRequestHeaderModifyRules) SetExceptDirectory(v string) *EditHttpHeaderConfigEditHttpHeaderConfigRequestHeaderModifyRules {
  s.ExceptDirectory = &v
  return s
}

func (s *EditHttpHeaderConfigEditHttpHeaderConfigRequestHeaderModifyRules) SetExceptRequestMethod(v string) *EditHttpHeaderConfigEditHttpHeaderConfigRequestHeaderModifyRules {
  s.ExceptRequestMethod = &v
  return s
}

func (s *EditHttpHeaderConfigEditHttpHeaderConfigRequestHeaderModifyRules) SetExceptEditHttpHeaderConfigRequestHeader(v string) *EditHttpHeaderConfigEditHttpHeaderConfigRequestHeaderModifyRules {
  s.ExceptEditHttpHeaderConfigRequestHeader = &v
  return s
}

type EditHttpHeaderConfigResponse struct {
  // {"en":"httpstatus=202; Indicates that the new domain API was successfully invoked, and the current deployment of the new domain can be viewed using x-cnc-request-id in the header", "zh_CN":"httpstatus=202;   表示成功调用新增域名接口，可使用header中的x-cnc-request-id查看当前新增域名的部署情况"}
  HttpStatus *int `json:"http status code,omitempty" xml:"http status code,omitempty" require:"true"`
  // {"en":"Uniquely identified id for querying tasks per request (for all API)", "zh_CN":"唯一标示的id，用于查询每次请求的任务 （适用全部接口）"}
  XCncRequestId *string `json:"x-cnc-request-id,omitempty" xml:"x-cnc-request-id,omitempty" require:"true"`
  // {"en":"The URL used to access the domain information, where domain-id is the unique token generated by our cloud platform for the domain name and whose value is a string.", "zh_CN":"用于访问该域名信息的URL，其中domain-id为我司云平台为该域名生成的唯一标示，其值为字符串。"}
  Location *string `json:"location,omitempty" xml:"location,omitempty" require:"true"`
  // {"en":"Error code, which appears when HTTPStatus is not 202, represents the error type of the current request call", "zh_CN":"错误代码，当HTTPStatus不为202时出现，表示当前请求调用的错误类型"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response information, success when successful", "zh_CN":"响应信息，成功时为success"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s EditHttpHeaderConfigResponse) String() string {
  return tea.Prettify(s)
}

func (s EditHttpHeaderConfigResponse) GoString() string {
  return s.String()
}

func (s *EditHttpHeaderConfigResponse) SetHttpStatus(v int) *EditHttpHeaderConfigResponse {
  s.HttpStatus = &v
  return s
}

func (s *EditHttpHeaderConfigResponse) SetXCncRequestId(v string) *EditHttpHeaderConfigResponse {
  s.XCncRequestId = &v
  return s
}

func (s *EditHttpHeaderConfigResponse) SetLocation(v string) *EditHttpHeaderConfigResponse {
  s.Location = &v
  return s
}

func (s *EditHttpHeaderConfigResponse) SetCode(v string) *EditHttpHeaderConfigResponse {
  s.Code = &v
  return s
}

func (s *EditHttpHeaderConfigResponse) SetMessage(v string) *EditHttpHeaderConfigResponse {
  s.Message = &v
  return s
}

type EditHttpHeaderConfigPaths struct {
  // {"en":"the domain whoes need query config", "zh_CN":"需要查询配置的域名或域名id"}
  DomainName *string `json:"domain-name,omitempty" xml:"domain-name,omitempty" require:"true"`
}

func (s EditHttpHeaderConfigPaths) String() string {
  return tea.Prettify(s)
}

func (s EditHttpHeaderConfigPaths) GoString() string {
  return s.String()
}

func (s *EditHttpHeaderConfigPaths) SetDomainName(v string) *EditHttpHeaderConfigPaths {
  s.DomainName = &v
  return s
}

type EditHttpHeaderConfigParameters struct {
}

func (s EditHttpHeaderConfigParameters) String() string {
  return tea.Prettify(s)
}

func (s EditHttpHeaderConfigParameters) GoString() string {
  return s.String()
}

type EditHttpHeaderConfigRequestHeader struct {
}

func (s EditHttpHeaderConfigRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s EditHttpHeaderConfigRequestHeader) GoString() string {
  return s.String()
}

type EditHttpHeaderConfigResponseHeader struct {
}

func (s EditHttpHeaderConfigResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s EditHttpHeaderConfigResponseHeader) GoString() string {
  return s.String()
}




type UpdateIpVersionConfigRequest struct {
  // {"en":"1: Update DNS region IP version, optional values: V4, V6", "zh_CN":"可选值为V4，V6。新增域名时，默认是使用IPV4。可以开启V6协议。
  // 示例：
  // \"ipVersion\":[\"V4\",\"V6\"]，表示全球使用V4和V6资源。
  // \"ipVersion\":[\"V4\"]，表示全球使用V4资源。
  // \"ipVersion\":[\"V6\"]，表示全球使用V6资源。"}
  IpVersion []*string `json:"ipVersion,omitempty" xml:"ipVersion,omitempty" type:"Repeated"`
}

func (s UpdateIpVersionConfigRequest) String() string {
  return tea.Prettify(s)
}

func (s UpdateIpVersionConfigRequest) GoString() string {
  return s.String()
}

func (s *UpdateIpVersionConfigRequest) SetIpVersion(v []*string) *UpdateIpVersionConfigRequest {
  s.IpVersion = v
  return s
}

type UpdateIpVersionConfigResponse struct {
  // {"en":"Error code. 0: successful", "zh_CN":"错误代码。 0：成功"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"The response message. Response 'Success' when calling API successful.", "zh_CN":"响应信息，成功时为success"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"httpstatus=202; Indicates that the new domain API was successfully invoked, and the current deployment of the new domain can be viewed using x-cnc-request-id in the header", "zh_CN":"httpstatus=202;   表示成功调用新增域名接口，可使用header中的x-cnc-request-id查看当前新增域名的部署情况"}
  HttpStatus *int `json:"httpStatus,omitempty" xml:"httpStatus,omitempty" require:"true"`
  // {"en":"The body of return data.", "zh_CN":"返回体数据"}
  Data *string `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s UpdateIpVersionConfigResponse) String() string {
  return tea.Prettify(s)
}

func (s UpdateIpVersionConfigResponse) GoString() string {
  return s.String()
}

func (s *UpdateIpVersionConfigResponse) SetCode(v string) *UpdateIpVersionConfigResponse {
  s.Code = &v
  return s
}

func (s *UpdateIpVersionConfigResponse) SetMessage(v string) *UpdateIpVersionConfigResponse {
  s.Message = &v
  return s
}

func (s *UpdateIpVersionConfigResponse) SetHttpStatus(v int) *UpdateIpVersionConfigResponse {
  s.HttpStatus = &v
  return s
}

func (s *UpdateIpVersionConfigResponse) SetData(v string) *UpdateIpVersionConfigResponse {
  s.Data = &v
  return s
}

type UpdateIpVersionConfigPaths struct {
  // {"en":"Domain name or domain id to query configuration", "zh_CN":"需要查询配置的域名（domainName）或域名id（domainId）"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
}

func (s UpdateIpVersionConfigPaths) String() string {
  return tea.Prettify(s)
}

func (s UpdateIpVersionConfigPaths) GoString() string {
  return s.String()
}

func (s *UpdateIpVersionConfigPaths) SetDomain(v string) *UpdateIpVersionConfigPaths {
  s.Domain = &v
  return s
}

type UpdateIpVersionConfigParameters struct {
}

func (s UpdateIpVersionConfigParameters) String() string {
  return tea.Prettify(s)
}

func (s UpdateIpVersionConfigParameters) GoString() string {
  return s.String()
}

type UpdateIpVersionConfigRequestHeader struct {
}

func (s UpdateIpVersionConfigRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s UpdateIpVersionConfigRequestHeader) GoString() string {
  return s.String()
}

type UpdateIpVersionConfigResponseHeader struct {
}

func (s UpdateIpVersionConfigResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s UpdateIpVersionConfigResponseHeader) GoString() string {
  return s.String()
}




type QueryHttpHeaderConfigRequest struct {
}

func (s QueryHttpHeaderConfigRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryHttpHeaderConfigRequest) GoString() string {
  return s.String()
}

type QueryHttpHeaderConfigResponse struct {
  // {"en":"Domain name or domain name id to query configuration", "zh_CN":"需要查询配置的域名或域名id"}
  DomainName *string `json:"domain-name,omitempty" xml:"domain-name,omitempty" require:"true"`
  // {"en":"Domain name or domain name id to query configuration", "zh_CN":"需要查询配置的域名或域名id"}
  DomainId *string `json:"domain-id,omitempty" xml:"domain-id,omitempty" require:"true"`
  HeaderModifyRules []*QueryHttpHeaderConfigQueryHttpHeaderConfigResponseHeaderModifyRules `json:"header-modify-rules,omitempty" xml:"header-modify-rules,omitempty" require:"true" type:"Repeated"`
}

func (s QueryHttpHeaderConfigResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryHttpHeaderConfigResponse) GoString() string {
  return s.String()
}

func (s *QueryHttpHeaderConfigResponse) SetDomainName(v string) *QueryHttpHeaderConfigResponse {
  s.DomainName = &v
  return s
}

func (s *QueryHttpHeaderConfigResponse) SetDomainId(v string) *QueryHttpHeaderConfigResponse {
  s.DomainId = &v
  return s
}

func (s *QueryHttpHeaderConfigResponse) SetHeaderModifyRules(v []*QueryHttpHeaderConfigQueryHttpHeaderConfigResponseHeaderModifyRules) *QueryHttpHeaderConfigResponse {
  s.HeaderModifyRules = v
  return s
}

type QueryHttpHeaderConfigQueryHttpHeaderConfigResponseHeaderModifyRules struct     {
  // {"en":"Exception url matching pattern, support regular. Example: ", "zh_CN":"例外的url匹配模式，支持正则。 入参参考："}
  ExceptPathPattern *string `json:"except-path-pattern,omitempty" xml:"except-path-pattern,omitempty" require:"true"`
  // {"en":"Matching conditions: specify common types, optional values are all or homepage. 1. all: all files 2. homepage: home page", "zh_CN":"匹配条件：指定常用类型，可选值为all或homepage 1、all：全部文件 2、homepage：首页"}
  CustomPattern *string `json:"custom-pattern,omitempty" xml:"custom-pattern,omitempty" require:"true"`
  // {"en":"Matching conditions: file type, please separate by semicolon, optional values: gif png bmp jpeg jpg html htm shtml mp3 wma flv mp4 wmv zip exe rar css txt ico js swf m3u8 xml f4m bootstarp ts.", "zh_CN":"匹配条件：文件类型，多个请以英文;分隔，可选值：gif png bmp jpeg jpg html htm shtml mp3 wma flv mp4 wmv zip exe rar css txt ico js swf m3u8 xml f4m bootstarp ts"}
  FileType *string `json:"file-type,omitempty" xml:"file-type,omitempty" require:"true"`
  // {"en":"Matching condition: Custom file type, separate by semicolon.", "zh_CN":"匹配条件：自定义文件类型，多个请以英文分号分隔。"}
  CustomFileType *string `json:"custom-file-type,omitempty" xml:"custom-file-type,omitempty" require:"true"`
  // {"en":"directory", "zh_CN":"目录"}
  Directory *string `json:"directory,omitempty" xml:"directory,omitempty" require:"true"`
  // {"en":"Matching Condition: Specify URL. The input parameter does not support the URI format starting with http(s)://", "zh_CN":"匹配条件：指定URL 入参不支持含http(s):// 开头的URI格式"}
  SpecifyUrl *string `json:"specify-url,omitempty" xml:"specify-url,omitempty" require:"true"`
  // {"en":"The matching request method, the optional values are: GET, POST, PUT, HEAD, DELETE, OPTIONS, separate by semicolons.", "zh_CN":"匹配的请求方式，可选值为：GET、POST、PUT、HEAD、DELETE、OPTIONS，多个请以英文分号分隔"}
  RequestMethod *string `json:"request-method,omitempty" xml:"request-method,omitempty" require:"true"`
  // {"en":"Add a grid type identifier to indicate a specific group configuration when the client has multiple groups of configurations.", "zh_CN":"添加grid类型标识，表示客户多组配置时，具体某组配置"}
  DataId *int `json:"data-id,omitempty" xml:"data-id,omitempty" require:"true"`
  // {"en":"The url matching mode supports fuzzy regularization. If all matches, the input parameters can be configured as: *", "zh_CN":"url匹配模式，支持正则，如果是全部匹配，入参可以配置为：.*"}
  PathPattern *string `json:" path-pattern,omitempty" xml:" path-pattern,omitempty" require:"true"`
  // {"en":"The control direction of the http header, the optional value is cache2visitor/cache2origin/visitor2cache/origin2cache, single-select.
  // Cache2origin refers to the source direction---corresponding to the configuration item return source request;
  // Cache2visitor refers to the direction of the client back - the corresponding configuration item returns to the client response;
  // Visitor2cache refers to receiving client requests
  // Origin2cache refers to the receiving source response", "zh_CN":"http头的控制方向，可选值为cache2visitor/cache2origin/visitor2cache/origin2cache，单选。&nbsp;
  // cache2origin是指回源方向---对应配置项回源请求；&nbsp;
  // cache2visitor是指回客户端方向&mdash;对应配置项回客户端应答；
  // visitor2cache是指接收客户端请求
  // origin2cache是指接收源应答"}
  HeaderDirection *string `json:"header-direction,omitempty" xml:"header-direction,omitempty" require:"true"`
  // {"en":"The control type of the http header supports the addition and deletion of the http header value. The optional value is add|set|delete, which is single-selected. Corresponding to the header-name and header-value parameters
  // Add: add a header
  // Set: modify the header
  // Delete: delete the header
  // Note: priority is delete>set>add", "zh_CN":"http头的控制类型，支持http头部值的增改删，可选值为add|delete|set，单选。对应header-name、header-value参数
  // add：新增头部
  // set：修改头部
  // delete：删除头部
  // 注意：优先级顺序为delete>set>add"}
  Action *string `json:"action,omitempty" xml:"action,omitempty" require:"true"`
  // {"en":"Http header regular match, optional value: true / false.
  // True: indicates that the value of the header-name is handled as a regular match.
  // False: indicates that the value of the header-name is processed according to the actual parameters, and no regular match is made.
  // Do not pass the default is false", "zh_CN":"http头正则匹配，可选值：true/false。
  // true：表示对header-name的值按正则匹配方式处理
  // false:表示对header-name的值按实际入参处理，不做正则匹配。
  // 不传默认是false"}
  AllowRegexp *string `json:" allow-regexp,omitempty" xml:" allow-regexp,omitempty" require:"true"`
  // {"en":"Http header name, add or modify the http header, only one is allowed; delete the http header to allow multiple entries, separated by a semicolon ';'.
  // Note: The operation of the special http header is limited, and the http header and operation type of the operation are allowed.
  // This item is required and cannot be empty
  // When the action is add: indicates that the header-name header is added.
  // When the action is set: modify the header-name header
  // When the action is delete: delete the header-name header", "zh_CN":"http头名称，新增或修改http头，只允许输入一个；删除http头允许输入多个，以分号&ldquo;;&rdquo;隔开。
  // 注意：对特殊http头的操作是受限的，允许操作的http头及操作类型参看
  // 此项为必填项，不能为空
  // 当action为add：表示新增这个header-name头部
  // 当action为set：修改这个header-name头部
  // 当action为delete：删除这个header-name头部"}
  HeaderName *string `json:"header-name,omitempty" xml:"header-name,omitempty" require:"true"`
  // {"en":"The value corresponding to the HTTP header field, for example: mytest.example.com
  // 
  // Note:
  // 
  // 1. When the action is add or set, the input parameter must be passed a value
  // 
  // 2. When the action is delete, the input parameter is not passed
  // 
  // Support to get the value of specified variable by keyword, such as client IP, including:
  // 
  // Key words: meaning
  // 
  // #timestamp: current time, timestamp as 1559124945
  // #request-host: host in the request header
  // #request-url: request url, which contains the full path of the protocol domain name, etc., such as http://aaa.aa.com/a.html
  // #request-uri: request uri, relative path format, such as /index.html
  // #origin- IP: return source IP
  // #cache-ip: edge node IP
  // #server-ip: external service IP
  // #client-ip: client IP, or visitor IP
  // #response-header{XXX} : get the value in the response header, such as #response-header{etag}, get the etag value in response-header
  // 
  // #header{XXX} : to get the value in the HTTP header of the request, such as #header{user-agent}, is to get the user-agent value in the header
  // 
  // #cookie{XXX} : get the value in the cookie, such as #cookie{account}, is to get the value of the account set in the cookie", "zh_CN":"http头域对应的值，例如：mytest.example.com   
  // 注意：
  // 1、当action为add或set时，该入参必须传值
  // 2、当action为delete时，该入参不用传
  // 支持通过关键字获取指定变量值，如客户端ip，包含如下：
  // 关键字：含义
  // #timestamp ：当前时间，时间戳如1559124945
  // #request-host：请求头中的HOST
  // #request-url ：请求url，包含协议域名等的全路径，如http://aaa.aa.com/a.html
  // #request-uri ：请求uri，相对路径格式，如/index.html
  // #origin-ip ：回源IP
  // #cache-ip ：边缘节点IP
  // #server-ip ：对外服务IP
  // #client-ip ：客户端IP，即访客IP
  // #response-header{xxx}：获取响应头中的值，如#response-header{etag}，获取response-header中的etag值
  // #header{xxx}：获取请求的http header中的值，如#header{User-Agent}，是获取header中的User-Agent值
  // #cookie{xxx}：获取cookie中的值，如#cookie{account}，是获取cookie中设置的account的值&nbsp;&nbsp;"}
  HeaderValue *string `json:"header-value,omitempty" xml:"header-value,omitempty" require:"true"`
  // {"en":"The original value corresponding to the HTTP header field","zh_CN":"http头域对应的原始值"}
  HeaderValueOld *string `json:"header-value-old,omitempty" xml:"header-value-old,omitempty" require:"true"`
  // {"en":"Match request header, header values support regular, header and header values separated by Spaces, e.g. : Range bytes=[0-9]{9,}", "zh_CN":"匹配请求头，头部值支持正则，头和头部值用空格隔开，如：Range bytes=[0-9]{9,}"}
  RequesHeader *string `json:"request-header,omitempty" xml:"request-header,omitempty" require:"true"`
  // {"en":"Indicates the priority of execution order for multiple sets of configurations. A higher number indicates higher priority. If no parameters are passed, the default value is 10 and cannot be cleared.", "zh_CN":"表示客户多组配置的优先执行顺序。数字越大，优先级越高。 不传参默认为10，不可清空"}
  Priority *string `json:"priority,omitempty" xml:"priority,omitempty" require:"true"`
  // {"en":"Exception file type.", "zh_CN":"例外的文件类型	"}
  ExceptFileType *string `json:"except-file-type,omitempty" xml:"except-file-type,omitempty" require:"true"`
  // {"en":"Exception directory.", "zh_CN":"例外的目录"}
  ExceptDirectory *string `json:"except-directory,omitempty" xml:"except-directory,omitempty" require:"true"`
  // {"en":"Exception request method.", "zh_CN":"例外的请求方式"}
  ExceptRequestMethod *string `json:"except-request-method,omitempty" xml:"except-request-method,omitempty" require:"true"`
  // {"en":"Exception request header.", "zh_CN":"例外的请求头"}
  ExceptQueryHttpHeaderConfigRequestHeader *string `json:"except-request-header,omitempty" xml:"except-request-header,omitempty" require:"true"`
}

func (s QueryHttpHeaderConfigQueryHttpHeaderConfigResponseHeaderModifyRules) String() string {
  return tea.Prettify(s)
}

func (s QueryHttpHeaderConfigQueryHttpHeaderConfigResponseHeaderModifyRules) GoString() string {
  return s.String()
}

func (s *QueryHttpHeaderConfigQueryHttpHeaderConfigResponseHeaderModifyRules) SetExceptPathPattern(v string) *QueryHttpHeaderConfigQueryHttpHeaderConfigResponseHeaderModifyRules {
  s.ExceptPathPattern = &v
  return s
}

func (s *QueryHttpHeaderConfigQueryHttpHeaderConfigResponseHeaderModifyRules) SetCustomPattern(v string) *QueryHttpHeaderConfigQueryHttpHeaderConfigResponseHeaderModifyRules {
  s.CustomPattern = &v
  return s
}

func (s *QueryHttpHeaderConfigQueryHttpHeaderConfigResponseHeaderModifyRules) SetFileType(v string) *QueryHttpHeaderConfigQueryHttpHeaderConfigResponseHeaderModifyRules {
  s.FileType = &v
  return s
}

func (s *QueryHttpHeaderConfigQueryHttpHeaderConfigResponseHeaderModifyRules) SetCustomFileType(v string) *QueryHttpHeaderConfigQueryHttpHeaderConfigResponseHeaderModifyRules {
  s.CustomFileType = &v
  return s
}

func (s *QueryHttpHeaderConfigQueryHttpHeaderConfigResponseHeaderModifyRules) SetDirectory(v string) *QueryHttpHeaderConfigQueryHttpHeaderConfigResponseHeaderModifyRules {
  s.Directory = &v
  return s
}

func (s *QueryHttpHeaderConfigQueryHttpHeaderConfigResponseHeaderModifyRules) SetSpecifyUrl(v string) *QueryHttpHeaderConfigQueryHttpHeaderConfigResponseHeaderModifyRules {
  s.SpecifyUrl = &v
  return s
}

func (s *QueryHttpHeaderConfigQueryHttpHeaderConfigResponseHeaderModifyRules) SetRequestMethod(v string) *QueryHttpHeaderConfigQueryHttpHeaderConfigResponseHeaderModifyRules {
  s.RequestMethod = &v
  return s
}

func (s *QueryHttpHeaderConfigQueryHttpHeaderConfigResponseHeaderModifyRules) SetDataId(v int) *QueryHttpHeaderConfigQueryHttpHeaderConfigResponseHeaderModifyRules {
  s.DataId = &v
  return s
}

func (s *QueryHttpHeaderConfigQueryHttpHeaderConfigResponseHeaderModifyRules) SetPathPattern(v string) *QueryHttpHeaderConfigQueryHttpHeaderConfigResponseHeaderModifyRules {
  s.PathPattern = &v
  return s
}

func (s *QueryHttpHeaderConfigQueryHttpHeaderConfigResponseHeaderModifyRules) SetHeaderDirection(v string) *QueryHttpHeaderConfigQueryHttpHeaderConfigResponseHeaderModifyRules {
  s.HeaderDirection = &v
  return s
}

func (s *QueryHttpHeaderConfigQueryHttpHeaderConfigResponseHeaderModifyRules) SetAction(v string) *QueryHttpHeaderConfigQueryHttpHeaderConfigResponseHeaderModifyRules {
  s.Action = &v
  return s
}

func (s *QueryHttpHeaderConfigQueryHttpHeaderConfigResponseHeaderModifyRules) SetAllowRegexp(v string) *QueryHttpHeaderConfigQueryHttpHeaderConfigResponseHeaderModifyRules {
  s.AllowRegexp = &v
  return s
}

func (s *QueryHttpHeaderConfigQueryHttpHeaderConfigResponseHeaderModifyRules) SetHeaderName(v string) *QueryHttpHeaderConfigQueryHttpHeaderConfigResponseHeaderModifyRules {
  s.HeaderName = &v
  return s
}

func (s *QueryHttpHeaderConfigQueryHttpHeaderConfigResponseHeaderModifyRules) SetHeaderValue(v string) *QueryHttpHeaderConfigQueryHttpHeaderConfigResponseHeaderModifyRules {
  s.HeaderValue = &v
  return s
}

func (s *QueryHttpHeaderConfigQueryHttpHeaderConfigResponseHeaderModifyRules) SetHeaderValueOld(v string) *QueryHttpHeaderConfigQueryHttpHeaderConfigResponseHeaderModifyRules {
  s.HeaderValueOld = &v
  return s
}

func (s *QueryHttpHeaderConfigQueryHttpHeaderConfigResponseHeaderModifyRules) SetRequesHeader(v string) *QueryHttpHeaderConfigQueryHttpHeaderConfigResponseHeaderModifyRules {
  s.RequesHeader = &v
  return s
}

func (s *QueryHttpHeaderConfigQueryHttpHeaderConfigResponseHeaderModifyRules) SetPriority(v string) *QueryHttpHeaderConfigQueryHttpHeaderConfigResponseHeaderModifyRules {
  s.Priority = &v
  return s
}

func (s *QueryHttpHeaderConfigQueryHttpHeaderConfigResponseHeaderModifyRules) SetExceptFileType(v string) *QueryHttpHeaderConfigQueryHttpHeaderConfigResponseHeaderModifyRules {
  s.ExceptFileType = &v
  return s
}

func (s *QueryHttpHeaderConfigQueryHttpHeaderConfigResponseHeaderModifyRules) SetExceptDirectory(v string) *QueryHttpHeaderConfigQueryHttpHeaderConfigResponseHeaderModifyRules {
  s.ExceptDirectory = &v
  return s
}

func (s *QueryHttpHeaderConfigQueryHttpHeaderConfigResponseHeaderModifyRules) SetExceptRequestMethod(v string) *QueryHttpHeaderConfigQueryHttpHeaderConfigResponseHeaderModifyRules {
  s.ExceptRequestMethod = &v
  return s
}

func (s *QueryHttpHeaderConfigQueryHttpHeaderConfigResponseHeaderModifyRules) SetExceptQueryHttpHeaderConfigRequestHeader(v string) *QueryHttpHeaderConfigQueryHttpHeaderConfigResponseHeaderModifyRules {
  s.ExceptQueryHttpHeaderConfigRequestHeader = &v
  return s
}

type QueryHttpHeaderConfigPaths struct {
  // {"en":"the domain whoes need query config", "zh_CN":"需要查询配置的域名或域名id"}
  DomainName *string `json:"domain-name,omitempty" xml:"domain-name,omitempty" require:"true"`
}

func (s QueryHttpHeaderConfigPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryHttpHeaderConfigPaths) GoString() string {
  return s.String()
}

func (s *QueryHttpHeaderConfigPaths) SetDomainName(v string) *QueryHttpHeaderConfigPaths {
  s.DomainName = &v
  return s
}

type QueryHttpHeaderConfigParameters struct {
}

func (s QueryHttpHeaderConfigParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryHttpHeaderConfigParameters) GoString() string {
  return s.String()
}

type QueryHttpHeaderConfigRequestHeader struct {
}

func (s QueryHttpHeaderConfigRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryHttpHeaderConfigRequestHeader) GoString() string {
  return s.String()
}

type QueryHttpHeaderConfigResponseHeader struct {
}

func (s QueryHttpHeaderConfigResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryHttpHeaderConfigResponseHeader) GoString() string {
  return s.String()
}




type UpdateCacheKeyConfigurationRequest struct {
  // {"en":"Custom Cachekey Configuration, parent node
  // 1. When you need to configure the cachekey rules,this must be filled in.
  // 2. Configuration of clearing for <cacheKeyRules/>.", "zh_CN":"配置自定义缓存key功能。
  // 1. 需要设置自定义缓存key配置时，此项必填
  // 2. 为<cacheKeyRules/>时清空自定义缓存key配置"}
  CacheKeyRules []*UpdateCacheKeyConfigurationRequestCacheKeyRules `json:"cacheKeyRules,omitempty" xml:"cacheKeyRules,omitempty" require:"true" type:"Repeated"`
}

func (s UpdateCacheKeyConfigurationRequest) String() string {
  return tea.Prettify(s)
}

func (s UpdateCacheKeyConfigurationRequest) GoString() string {
  return s.String()
}

func (s *UpdateCacheKeyConfigurationRequest) SetCacheKeyRules(v []*UpdateCacheKeyConfigurationRequestCacheKeyRules) *UpdateCacheKeyConfigurationRequest {
  s.CacheKeyRules = v
  return s
}

type UpdateCacheKeyConfigurationRequestCacheKeyRules struct     {
  // {"en":"The url matching mode supports fuzzy regularization. If all matches, the input parameters can be configured as: *", "zh_CN":"url匹配模式，支持正则，如果是全部匹配，入参可以配置为：.*"}
  PathPattern *string `json:"pathPattern,omitempty" xml:"pathPattern,omitempty"`
  // {"en":"Specify a uri, such as /test/specifyurl", "zh_CN":"指定具体的uri，如/test/specifyurl"}
  SpecifyUrl *string `json:"specifyUrl,omitempty" xml:"specifyUrl,omitempty"`
  // {"en":"Whether to match specifyUrl exactly or not, you can select true and false.
  // True:means match exactly. False: means fuzzy match, such as specifyUrl='/test/uri', and  request for /test/uri?p=1 will be matched.", "zh_CN":"是否完全匹配specifyUrl，可选择为true和false。
  // 为true则完全匹配；为false则模糊匹配，如指定/test/uri，请求/test/uri?p=1也会匹配"}
  FullMatch4SpecifyUrl *string `json:"fullMatch4SpecifyUrl,omitempty" xml:"fullMatch4SpecifyUrl,omitempty"`
  // {"en":"Specify common types: Select the domain name that requires the cache to be all files or the home page. :
  // E.g:
  // All: all files
  // Homepage: homepage", "zh_CN":"指定常用类型：选择缓存域名的是全部文件还是首页。入参参考值： all：全部文件 homepage：首页"}
  CustomPattern *string `json:"customPattern,omitempty" xml:"customPattern,omitempty"`
  // {"en":"File Type: Specify the file type for cache settings.
  // File types include: gif png bmp jpeg jpg html htm shtml mp3 wma flv mp4 wmv zip exe rar css txt ico js swf
  // If you need all types, pass all directly. Multiples are separated by semicolons, and all and specific file types cannot be configured at the same time.", "zh_CN":"文件类型：指定需要缓存的文件类型。 文件类型包括：gif png bmp jpeg jpg html htm shtml mp3 wma flv mp4 wmv zip exe rar css txt ico js swf 如果需要全部类型，则直接传all。多个以分号隔开，all和具体文件类型不能同时配置。"}
  FileType *string `json:"fileType,omitempty" xml:"fileType,omitempty"`
  // {"en":"Custom file type: Fill in the appropriate identifiable file type according to your needs outside of the specified file type. Can be used with file-type. If the file-type is also configured, the actual file type is the sum of the two parameters.", "zh_CN":"自定义文件类型：在指定文件类型外根据自身需求，填写适当的可识别文件类型。可以搭配file-type使用。如果file-type也有配置，实际生效的文件类型是两个入参的总和"}
  CustomFileType *string `json:"customFileType,omitempty" xml:"customFileType,omitempty"`
  // {"en":"Directory: Specify the directory cache.
  // Enter a legal directory format. Multiple separated by semicolons", "zh_CN":"目录：指定目录缓存。 输入合法的目录格式。多个以英文分号隔开"}
  Directory *string `json:"directory,omitempty" xml:"directory,omitempty"`
  // {"en":"Ignore case, the optional value is true or false, true means to ignore case; false means not to ignore case;
  // When adding a new configuration item, the default is true", "zh_CN":"是否忽略大小写：允许值为true和false，默认为忽略"}
  IgnoreCase *string `json:"ignoreCase,omitempty" xml:"ignoreCase,omitempty"`
  // {"en":"Header name.
  // Example: If you specify a header as &lsquo;lang', Then, if the value of Lang is consistent, one copy will be cached", "zh_CN":"头部名称
  // 例如：指定头部lang，lang的值一致则缓存一份"}
  HeaderName *string `json:"headerName,omitempty" xml:"headerName,omitempty"`
  // {"en":"Parameter Of the specified Header，
  // Example: Specifies the header as 'cookie', parameterOfHeader as 'name'. Then, if the value of name is consistent, one copy will be cached.", "zh_CN":"头部值的参数名，
  // 例如：指定头部Cookie，头部值的参数名为name。则name的值一致则缓存一份。"}
  ParameterOfHeader *string `json:"parameterOfHeader,omitempty" xml:"parameterOfHeader,omitempty"`
  // {"en":"Indicates the priority execution order of multiple sets of redirected content by the customer. The higher the number, the higher the priority.
  // When adding a new configuration item, the default is 10", "zh_CN":"优先级，表示客户多组配置的优先执行顺序。数字越大，优先级越高。不传默认为10，不可清空。"}
  Priority *string `json:"priority,omitempty" xml:"priority,omitempty"`
  // {"en":"DataId is to indicate a specific group configuration when the client has multiple groups of configurations. dataId can be retrieved through a query interface. Note: A. If dataId is passed, it means that one group of configuration items is specified to be modified, and no other group configuration items need to be modified. B. If multiple groups of configurations are included, some of them are configured with dataId and others are not, then the expression of dataId is used to modify a specific group of configurations, and a new group of configurations is added on the original basis without the expression of dataId. C. If the dataId is not transmitted, it means that the original configuration will be fully covered by this configuration. D. If no configuration parameter is passed, only domain name and secondary label are passed, which means that all configuration of domain name secondary service corresponding to this interface is cleared. E. If there is no specific configuration item in a set of configurations, the dataId must be filled in, and the value is the actual dataId, which means clearing the value of the corresponding dataId configuration item; it is not allowed that there is no specific configuration item or dataId in a set of configurations.", "zh_CN":"配置多组配置时，具体某组配置的id。本功能只支持一组配置。dataId可以通过查询接口获取。 注意： a、如果有传dataId，说明指定修改该组配置项内容； b、如果入参都没有传dataId,表示用本次的配置全量覆盖原先配置； c、如果一组配置没有具体的配置项，则dataId必填，且值为实际存在的dataId，表示清空这个dataId对应配置项的值；不允许一组配置没有具体的配置项也没有dataId。"}
  DataId *int64 `json:"dataId,omitempty" xml:"dataId,omitempty"`
}

func (s UpdateCacheKeyConfigurationRequestCacheKeyRules) String() string {
  return tea.Prettify(s)
}

func (s UpdateCacheKeyConfigurationRequestCacheKeyRules) GoString() string {
  return s.String()
}

func (s *UpdateCacheKeyConfigurationRequestCacheKeyRules) SetPathPattern(v string) *UpdateCacheKeyConfigurationRequestCacheKeyRules {
  s.PathPattern = &v
  return s
}

func (s *UpdateCacheKeyConfigurationRequestCacheKeyRules) SetSpecifyUrl(v string) *UpdateCacheKeyConfigurationRequestCacheKeyRules {
  s.SpecifyUrl = &v
  return s
}

func (s *UpdateCacheKeyConfigurationRequestCacheKeyRules) SetFullMatch4SpecifyUrl(v string) *UpdateCacheKeyConfigurationRequestCacheKeyRules {
  s.FullMatch4SpecifyUrl = &v
  return s
}

func (s *UpdateCacheKeyConfigurationRequestCacheKeyRules) SetCustomPattern(v string) *UpdateCacheKeyConfigurationRequestCacheKeyRules {
  s.CustomPattern = &v
  return s
}

func (s *UpdateCacheKeyConfigurationRequestCacheKeyRules) SetFileType(v string) *UpdateCacheKeyConfigurationRequestCacheKeyRules {
  s.FileType = &v
  return s
}

func (s *UpdateCacheKeyConfigurationRequestCacheKeyRules) SetCustomFileType(v string) *UpdateCacheKeyConfigurationRequestCacheKeyRules {
  s.CustomFileType = &v
  return s
}

func (s *UpdateCacheKeyConfigurationRequestCacheKeyRules) SetDirectory(v string) *UpdateCacheKeyConfigurationRequestCacheKeyRules {
  s.Directory = &v
  return s
}

func (s *UpdateCacheKeyConfigurationRequestCacheKeyRules) SetIgnoreCase(v string) *UpdateCacheKeyConfigurationRequestCacheKeyRules {
  s.IgnoreCase = &v
  return s
}

func (s *UpdateCacheKeyConfigurationRequestCacheKeyRules) SetHeaderName(v string) *UpdateCacheKeyConfigurationRequestCacheKeyRules {
  s.HeaderName = &v
  return s
}

func (s *UpdateCacheKeyConfigurationRequestCacheKeyRules) SetParameterOfHeader(v string) *UpdateCacheKeyConfigurationRequestCacheKeyRules {
  s.ParameterOfHeader = &v
  return s
}

func (s *UpdateCacheKeyConfigurationRequestCacheKeyRules) SetPriority(v string) *UpdateCacheKeyConfigurationRequestCacheKeyRules {
  s.Priority = &v
  return s
}

func (s *UpdateCacheKeyConfigurationRequestCacheKeyRules) SetDataId(v int64) *UpdateCacheKeyConfigurationRequestCacheKeyRules {
  s.DataId = &v
  return s
}

type UpdateCacheKeyConfigurationResponse struct {
  // {"en":"Error code. Appeared when http status is not 200 or 202.", "zh_CN":"错误代码，当HTTPStatus不为202时出现，表示当前请求调用的错误类型"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Reponse message.", "zh_CN":"响应信息，成功时为success"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Response data", "zh_CN":"响应数据"}
  Data *string `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s UpdateCacheKeyConfigurationResponse) String() string {
  return tea.Prettify(s)
}

func (s UpdateCacheKeyConfigurationResponse) GoString() string {
  return s.String()
}

func (s *UpdateCacheKeyConfigurationResponse) SetCode(v string) *UpdateCacheKeyConfigurationResponse {
  s.Code = &v
  return s
}

func (s *UpdateCacheKeyConfigurationResponse) SetMessage(v string) *UpdateCacheKeyConfigurationResponse {
  s.Message = &v
  return s
}

func (s *UpdateCacheKeyConfigurationResponse) SetData(v string) *UpdateCacheKeyConfigurationResponse {
  s.Data = &v
  return s
}

type UpdateCacheKeyConfigurationPaths struct {
  // {"en":"the domain whoes need query config", "zh_CN":"需要查询配置的域名或域名id"}
  DomainName *string `json:"domain-name,omitempty" xml:"domain-name,omitempty" require:"true"`
}

func (s UpdateCacheKeyConfigurationPaths) String() string {
  return tea.Prettify(s)
}

func (s UpdateCacheKeyConfigurationPaths) GoString() string {
  return s.String()
}

func (s *UpdateCacheKeyConfigurationPaths) SetDomainName(v string) *UpdateCacheKeyConfigurationPaths {
  s.DomainName = &v
  return s
}

type UpdateCacheKeyConfigurationParameters struct {
}

func (s UpdateCacheKeyConfigurationParameters) String() string {
  return tea.Prettify(s)
}

func (s UpdateCacheKeyConfigurationParameters) GoString() string {
  return s.String()
}

type UpdateCacheKeyConfigurationRequestHeader struct {
}

func (s UpdateCacheKeyConfigurationRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s UpdateCacheKeyConfigurationRequestHeader) GoString() string {
  return s.String()
}

type UpdateCacheKeyConfigurationResponseHeader struct {
}

func (s UpdateCacheKeyConfigurationResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s UpdateCacheKeyConfigurationResponseHeader) GoString() string {
  return s.String()
}




type QueryWebsocketConfigRequest struct {
}

func (s QueryWebsocketConfigRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryWebsocketConfigRequest) GoString() string {
  return s.String()
}

type QueryWebsocketConfigResponse struct {
  // {"en":"Error code, 0 is success.", "zh_CN":"错误代码，成功为0"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response information, success when successful", "zh_CN":"响应信息，成功时为success"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"data content, parent node.", "zh_CN":"数据内容，父标签"}
  Data *QueryWebsocketConfigResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
}

func (s QueryWebsocketConfigResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryWebsocketConfigResponse) GoString() string {
  return s.String()
}

func (s *QueryWebsocketConfigResponse) SetCode(v string) *QueryWebsocketConfigResponse {
  s.Code = &v
  return s
}

func (s *QueryWebsocketConfigResponse) SetMessage(v string) *QueryWebsocketConfigResponse {
  s.Message = &v
  return s
}

func (s *QueryWebsocketConfigResponse) SetData(v *QueryWebsocketConfigResponseData) *QueryWebsocketConfigResponse {
  s.Data = v
  return s
}

type QueryWebsocketConfigResponseData struct {
  // {"en":"domain id", "zh_CN":"域名id"}
  DomainId *int `json:"domainId,omitempty" xml:"domainId,omitempty" require:"true"`
  // {"en":"domain name", "zh_CN":"域名名称"}
  DomainName *string `json:"domainName,omitempty" xml:"domainName,omitempty" require:"true"`
  // {"en":"Open or close websocket function, parent node, you can set <websocketSwitch/> to clear this configuration.
  // Scope of application: wsa, web pages", "zh_CN":"开启或关闭websocket功能，父标签，为<websocketSwitch/>则清空websocket开关配置
  // 适用范围：wsa、网页"}
  WebsocketSwitch *QueryWebsocketConfigResponseDataWebsocketSwitch `json:"websocketSwitch,omitempty" xml:"websocketSwitch,omitempty" require:"true" type:"Struct"`
}

func (s QueryWebsocketConfigResponseData) String() string {
  return tea.Prettify(s)
}

func (s QueryWebsocketConfigResponseData) GoString() string {
  return s.String()
}

func (s *QueryWebsocketConfigResponseData) SetDomainId(v int) *QueryWebsocketConfigResponseData {
  s.DomainId = &v
  return s
}

func (s *QueryWebsocketConfigResponseData) SetDomainName(v string) *QueryWebsocketConfigResponseData {
  s.DomainName = &v
  return s
}

func (s *QueryWebsocketConfigResponseData) SetWebsocketSwitch(v *QueryWebsocketConfigResponseDataWebsocketSwitch) *QueryWebsocketConfigResponseData {
  s.WebsocketSwitch = v
  return s
}

type QueryWebsocketConfigResponseDataWebsocketSwitch struct {
  // {"en":"Whether to turn on the websocket function, the allowable values are true and false, default false", "zh_CN":"是否开启websocket功能,允许值为true和false，默认为否"}
  EnableWebsocket *bool `json:"enableWebsocket,omitempty" xml:"enableWebsocket,omitempty" require:"true"`
}

func (s QueryWebsocketConfigResponseDataWebsocketSwitch) String() string {
  return tea.Prettify(s)
}

func (s QueryWebsocketConfigResponseDataWebsocketSwitch) GoString() string {
  return s.String()
}

func (s *QueryWebsocketConfigResponseDataWebsocketSwitch) SetEnableWebsocket(v bool) *QueryWebsocketConfigResponseDataWebsocketSwitch {
  s.EnableWebsocket = &v
  return s
}

type QueryWebsocketConfigPaths struct {
  // {"en":"Domain name or domain name id to query configuration", "zh_CN":"需要查询配置的域名（domainName）或域名id（domainId）"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
}

func (s QueryWebsocketConfigPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryWebsocketConfigPaths) GoString() string {
  return s.String()
}

func (s *QueryWebsocketConfigPaths) SetDomain(v string) *QueryWebsocketConfigPaths {
  s.Domain = &v
  return s
}

type QueryWebsocketConfigParameters struct {
}

func (s QueryWebsocketConfigParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryWebsocketConfigParameters) GoString() string {
  return s.String()
}

type QueryWebsocketConfigRequestHeader struct {
}

func (s QueryWebsocketConfigRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryWebsocketConfigRequestHeader) GoString() string {
  return s.String()
}

type QueryWebsocketConfigResponseHeader struct {
}

func (s QueryWebsocketConfigResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryWebsocketConfigResponseHeader) GoString() string {
  return s.String()
}




type EditOriginUriAndHostRequest struct {
  // {"en":"Return path rewrite configuration
  // 
  // 1. When you need to set the rewrite configuration of the backsource path, this must be filled in
  // 
  // 2. Rewrite configuration for clearing the return path for <origin-rules-rewrites/>.", "zh_CN":"回源路径改写配置
  // 1.需要设置回源路径改写配置时，此项必填
  // 2.为<origin-rules-rewrites/>时清空回源路径改写配置"}
  OriginRulesRewrites []*EditOriginUriAndHostRequestOriginRulesRewrites `json:"originRulesRewrites,omitempty" xml:"originRulesRewrites,omitempty" require:"true" type:"Repeated"`
}

func (s EditOriginUriAndHostRequest) String() string {
  return tea.Prettify(s)
}

func (s EditOriginUriAndHostRequest) GoString() string {
  return s.String()
}

func (s *EditOriginUriAndHostRequest) SetOriginRulesRewrites(v []*EditOriginUriAndHostRequestOriginRulesRewrites) *EditOriginUriAndHostRequest {
  s.OriginRulesRewrites = v
  return s
}

type EditOriginUriAndHostRequestOriginRulesRewrites struct     {
  // {"en":"Add a grid type identifier to represent a specific group of configurations when a customer has multiple groups of configurations.
  // Note: Add grid type identifier: data-id, each group configuration corresponds to a data-id: a. If the customer has passed data-id, specify that modifying one group of configuration items content does not require modifying other group configuration content does not need to be included; B. If the customer enters multiple groups of configuration, some of them have data-id, some have not. If there is transmission, the expression of data-id is used to modify a specific group of configurations, but no expression of data-id is used to add a new group of configurations on the original basis; C. If no data-id is transmitted to the customer, it means that the original configuration is completely covered by this configuration; D. If no configuration parameters are transmitted to the customer, only the domain name and the second level are transmitted. Label, which indicates that clearing this interface corresponds to all configuration of domain name secondary service. (c, D content is consistent with the current solution); e, a gird tag can not be empty, if there is no specific configuration item, then data-id must be filled in, and the value is the actual data-id, indicating the value of clearing this data-id corresponding configuration item;", "zh_CN":"添加grid类型标识，表示客户多组配置时，具体某组配置
  // 注意：添加grid类型标识：data-id，每一组配置对应一个data-id：a、如果客户有传data-id，说明指定修改其中一组配置项内容，不需求修改其他组配置内容不需要入参；b、如果客户入参多组配置，其中有些组配置有传data-id，有些没有传，则有传data-id的表示修改具体某组配置，没有传data-id的表示在原来基础上新增一组配置；c、如果客户入参都没有传data-id,表示用本次的配置全量覆盖原先配置；d、如果客户入参没有传任何配置项参数，只传了域名和二级标签，表示清空这个接口对应域名二级服务所有配置。（c、d内容和当前方案实现一致）；e、一个gird标签下的入参不能为空，如果，没有具体的配置项，则data-id必填，且值为实际存在的data-id,表示清空这个data-id对应配置项的值；"}
  DataId *int64 `json:"dataId,omitempty" xml:"dataId,omitempty"`
  // {"en":"The URL matching mode supports regularization. If all matches are made, the input can be configured as:.*", "zh_CN":"url匹配模式，支持正则，如果是全部匹配，入参可以配置为：.*"}
  PathPattern *string `json:"pathPattern,omitempty" xml:"pathPattern,omitempty" require:"true"`
  // {"en":"The protocol of URL matching mode, which is used with path-pattern, is supported by adding parameters: default is empty, blank is default is to support HTTP and HTTPS protocol before URL matching mode path is needed at the same time;
  // 
  // Http: URL matches pattern path with HTTP protocol
  // 
  // Https: URL matches pattern path with HTTPS protocol
  // 
  // Ignore: URL matching mode path without protocol", "zh_CN":"url匹配模式的协议，该配置项与path-pattern搭配使用；入参支持：默认为空，为空则默认为需要同时支持url匹配模式路径前支持http和https协议；
  // http：url匹配模式路径前加上http协议
  // https:url 匹配模式路径前加上HTTPS协议
  // ignore:url匹配模式路径前不加协议"}
  PathPatternHttp *string `json:"pathPatternHttp,omitempty" xml:"pathPatternHttp,omitempty"`
  // {"en":"Exceptional URL matching pattern in the same format as path pattern", "zh_CN":"例外的url匹配模式，格式同pathPattern"}
  ExceptPathPattern *string `json:"exceptPathPattern,omitempty" xml:"exceptPathPattern,omitempty"`
  // {"en":"Exceptional URL matching mode protocol, which is used in conjunction with except-path-pattern; participation support: default is empty, blank is default is required to support both HTTP and HTTPS protocol before URL matching mode path;
  // 
  // Http: URL matches pattern path with HTTP protocol.
  // Https: URL matches pattern path with HTTPS protocol.
  // Ignore: URL matching mode path without protocol.", "zh_CN":"例外的url匹配模式的协议，该配置项与except-path-pattern搭配使用；入参支持：默认为空，为空则默认为需要同时支持url匹配模式路径前支持http和https协议；
  // http：url匹配模式路径前加上http协议
  // https:url 匹配模式路径前加上HTTPS协议
  // ignore:url匹配模式路径前不加协议"}
  ExceptPathPatternHttp *string `json:"exceptPathPatternHttp,omitempty" xml:"exceptPathPatternHttp,omitempty"`
  // {"en":"Ignore case or not: the allowable values are true and false, and the default is Ignore", "zh_CN":"是否忽略大小写：允许值为true和false，默认为忽略"}
  IgnoreLetterCase *string `json:"ignoreLetterCase,omitempty" xml:"ignoreLetterCase,omitempty"`
  // {"en":"Back-source information, you can enter IP or domain name.
  // That is, customer source IP or domain name", "zh_CN":"回源信息，可以输入ip或者域名
  // 即客户源站IP或域名"}
  OriginInfo *string `json:"originInfo,omitempty" xml:"originInfo,omitempty"`
  // {"en":"Represents the priority execution order of the customer's multi-group redirected content. The bigger the number, the higher the priority.", "zh_CN":"表示客户多组重定向内容的优先执行顺序。数字越大，优先级越高。
  // 新增配置项时，不传默认为 10
  // 如果传了值，不能为空"}
  Priority *string `json:"priority,omitempty" xml:"priority,omitempty" require:"true"`
  // {"en":"Back source host,that support to enter domain names;
  // Example: The backsource host of domain name A configures domain name B. When the A domain name requests the source, the requested URL uses the B domain name instead of the A domain name.", "zh_CN":"回源host,支持入参域名；
  // 示例：域名A的回源host配置了域名B。当A域名请求的回源的时候，请求的url上用B域名代替A域名&nbsp;"}
  OriginHost *string `json:"originHost,omitempty" xml:"originHost,omitempty"`
  // {"en":"Pre-rewrite uri. That is, the original request URI for user access. Support regular configuration", "zh_CN":"改写前的uri.&nbsp;即用户访问的原始请求uri&nbsp;。支持正则配置"}
  BeforeRewritedUri *string `json:"beforeRewritedUri,omitempty" xml:"beforeRewritedUri,omitempty"`
  // {"en":"The rewritten uri, the request URI configured before-rewrited-uri, is retrieved with the rewritten uri. Rewrite the source path. Support regular configuration", "zh_CN":"改写后的uri,即将&nbsp;before-rewrited-uri配置的请求uri，用改写后的uri回源。实现回源路径改写。支持正则配置"}
  AfterRewritedUri *string `json:"afterRewritedUri,omitempty" xml:"afterRewritedUri,omitempty"`
}

func (s EditOriginUriAndHostRequestOriginRulesRewrites) String() string {
  return tea.Prettify(s)
}

func (s EditOriginUriAndHostRequestOriginRulesRewrites) GoString() string {
  return s.String()
}

func (s *EditOriginUriAndHostRequestOriginRulesRewrites) SetDataId(v int64) *EditOriginUriAndHostRequestOriginRulesRewrites {
  s.DataId = &v
  return s
}

func (s *EditOriginUriAndHostRequestOriginRulesRewrites) SetPathPattern(v string) *EditOriginUriAndHostRequestOriginRulesRewrites {
  s.PathPattern = &v
  return s
}

func (s *EditOriginUriAndHostRequestOriginRulesRewrites) SetPathPatternHttp(v string) *EditOriginUriAndHostRequestOriginRulesRewrites {
  s.PathPatternHttp = &v
  return s
}

func (s *EditOriginUriAndHostRequestOriginRulesRewrites) SetExceptPathPattern(v string) *EditOriginUriAndHostRequestOriginRulesRewrites {
  s.ExceptPathPattern = &v
  return s
}

func (s *EditOriginUriAndHostRequestOriginRulesRewrites) SetExceptPathPatternHttp(v string) *EditOriginUriAndHostRequestOriginRulesRewrites {
  s.ExceptPathPatternHttp = &v
  return s
}

func (s *EditOriginUriAndHostRequestOriginRulesRewrites) SetIgnoreLetterCase(v string) *EditOriginUriAndHostRequestOriginRulesRewrites {
  s.IgnoreLetterCase = &v
  return s
}

func (s *EditOriginUriAndHostRequestOriginRulesRewrites) SetOriginInfo(v string) *EditOriginUriAndHostRequestOriginRulesRewrites {
  s.OriginInfo = &v
  return s
}

func (s *EditOriginUriAndHostRequestOriginRulesRewrites) SetPriority(v string) *EditOriginUriAndHostRequestOriginRulesRewrites {
  s.Priority = &v
  return s
}

func (s *EditOriginUriAndHostRequestOriginRulesRewrites) SetOriginHost(v string) *EditOriginUriAndHostRequestOriginRulesRewrites {
  s.OriginHost = &v
  return s
}

func (s *EditOriginUriAndHostRequestOriginRulesRewrites) SetBeforeRewritedUri(v string) *EditOriginUriAndHostRequestOriginRulesRewrites {
  s.BeforeRewritedUri = &v
  return s
}

func (s *EditOriginUriAndHostRequestOriginRulesRewrites) SetAfterRewritedUri(v string) *EditOriginUriAndHostRequestOriginRulesRewrites {
  s.AfterRewritedUri = &v
  return s
}

type EditOriginUriAndHostResponse struct {
  // {"en":"Error code. It pops up when the HTTPStatus is not 202, and shows the revoking error type of the current request.", "zh_CN":"错误代码，当HTTPStatus不为202时出现，表示当前请求调用的错误类型"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response message, and shows as success when it succeeds.", "zh_CN":"响应信息，成功时为success"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Response data.", "zh_CN":"返回数据主体"}
  Data *string `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s EditOriginUriAndHostResponse) String() string {
  return tea.Prettify(s)
}

func (s EditOriginUriAndHostResponse) GoString() string {
  return s.String()
}

func (s *EditOriginUriAndHostResponse) SetCode(v string) *EditOriginUriAndHostResponse {
  s.Code = &v
  return s
}

func (s *EditOriginUriAndHostResponse) SetMessage(v string) *EditOriginUriAndHostResponse {
  s.Message = &v
  return s
}

func (s *EditOriginUriAndHostResponse) SetData(v string) *EditOriginUriAndHostResponse {
  s.Data = &v
  return s
}

type EditOriginUriAndHostPaths struct {
  // {"en":"the domain whoes need query config", "zh_CN":"需要查询配置的域名或域名id"}
  DomainName *string `json:"domain-name,omitempty" xml:"domain-name,omitempty" require:"true"`
}

func (s EditOriginUriAndHostPaths) String() string {
  return tea.Prettify(s)
}

func (s EditOriginUriAndHostPaths) GoString() string {
  return s.String()
}

func (s *EditOriginUriAndHostPaths) SetDomainName(v string) *EditOriginUriAndHostPaths {
  s.DomainName = &v
  return s
}

type EditOriginUriAndHostParameters struct {
}

func (s EditOriginUriAndHostParameters) String() string {
  return tea.Prettify(s)
}

func (s EditOriginUriAndHostParameters) GoString() string {
  return s.String()
}

type EditOriginUriAndHostRequestHeader struct {
}

func (s EditOriginUriAndHostRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s EditOriginUriAndHostRequestHeader) GoString() string {
  return s.String()
}

type EditOriginUriAndHostResponseHeader struct {
}

func (s EditOriginUriAndHostResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s EditOriginUriAndHostResponseHeader) GoString() string {
  return s.String()
}




type QuerysourceverificationconfigRequest struct {
}

func (s QuerysourceverificationconfigRequest) String() string {
  return tea.Prettify(s)
}

func (s QuerysourceverificationconfigRequest) GoString() string {
  return s.String()
}

type QuerysourceverificationconfigResponse struct {
  // {"en":"", "zh_CN":"转推回源带参数鉴权配置
  // 注意：
  // 1、需要取消花椒转推回源带参数鉴权配置时，可以传入空节点<source-verification></source-verification>。
  // 2、表示需要设置花椒转推回源带参数鉴权配置时，此项必填"}
  SourceVerification *QuerysourceverificationconfigResponseSourceVerification `json:"source-verification,omitempty" xml:"source-verification,omitempty" require:"true" type:"Struct"`
}

func (s QuerysourceverificationconfigResponse) String() string {
  return tea.Prettify(s)
}

func (s QuerysourceverificationconfigResponse) GoString() string {
  return s.String()
}

func (s *QuerysourceverificationconfigResponse) SetSourceVerification(v *QuerysourceverificationconfigResponseSourceVerification) *QuerysourceverificationconfigResponse {
  s.SourceVerification = v
  return s
}

type QuerysourceverificationconfigResponseSourceVerification struct {
  // {"en":"", "zh_CN":"生效位置，代表配置生效的位置。可选值，多个分号隔开
  // Cache:边缘
  // stfuCache：静态中转
  // dyfuCache：动态中转"}
  SwitchPosition *string `json:"switch-position,omitempty" xml:"switch-position,omitempty" require:"true"`
  // {"en":"", "zh_CN":"是否开启回源校验。
  // 可选值：
  // close：关闭
  // push:转推使用
  // origin：回源使用
  // push-origin：即转推又回源使用"}
  RefSwitch *string `json:"ref-switch,omitempty" xml:"ref-switch,omitempty" require:"true"`
  // {"en":"", "zh_CN":"md5加密方式
  // 标识客户md5加密涉及的内容。参考入参
  // ||%S||BASE10||12345678||"}
  Md5Path *string `json:"md5-path,omitempty" xml:"md5-path,omitempty" require:"true"`
  // {"en":"", "zh_CN":"时间偏移量:未配置时以当前时间为准；配了时移就以当前时间+时移秒数的那个时间为准。不传默认为0，即无时间偏移"}
  TimeOffset *int `json:"time-offset,omitempty" xml:"time-offset,omitempty" require:"true"`
  // {"en":"", "zh_CN":"sign字段名称：问号后需要鉴权的参数名
  // 参考入参：sign"}
  SignName *string `json:"sign-name,omitempty" xml:"sign-name,omitempty" require:"true"`
  // {"en":"", "zh_CN":"time字段名称：问号后需要鉴权的参数名
  // 参考入参：time"}
  TimeName *string `json:"time-name,omitempty" xml:"time-name,omitempty" require:"true"`
  // {"en":"", "zh_CN":"额外固定参数：需要带额外的参数回客户源站鉴权
  // 参考入参：trans=live.a.com"}
  OtherQuerysourceverificationconfigParameters *string `json:"other-parameters,omitempty" xml:"other-parameters,omitempty" require:"true"`
}

func (s QuerysourceverificationconfigResponseSourceVerification) String() string {
  return tea.Prettify(s)
}

func (s QuerysourceverificationconfigResponseSourceVerification) GoString() string {
  return s.String()
}

func (s *QuerysourceverificationconfigResponseSourceVerification) SetSwitchPosition(v string) *QuerysourceverificationconfigResponseSourceVerification {
  s.SwitchPosition = &v
  return s
}

func (s *QuerysourceverificationconfigResponseSourceVerification) SetRefSwitch(v string) *QuerysourceverificationconfigResponseSourceVerification {
  s.RefSwitch = &v
  return s
}

func (s *QuerysourceverificationconfigResponseSourceVerification) SetMd5Path(v string) *QuerysourceverificationconfigResponseSourceVerification {
  s.Md5Path = &v
  return s
}

func (s *QuerysourceverificationconfigResponseSourceVerification) SetTimeOffset(v int) *QuerysourceverificationconfigResponseSourceVerification {
  s.TimeOffset = &v
  return s
}

func (s *QuerysourceverificationconfigResponseSourceVerification) SetSignName(v string) *QuerysourceverificationconfigResponseSourceVerification {
  s.SignName = &v
  return s
}

func (s *QuerysourceverificationconfigResponseSourceVerification) SetTimeName(v string) *QuerysourceverificationconfigResponseSourceVerification {
  s.TimeName = &v
  return s
}

func (s *QuerysourceverificationconfigResponseSourceVerification) SetOtherQuerysourceverificationconfigParameters(v string) *QuerysourceverificationconfigResponseSourceVerification {
  s.OtherQuerysourceverificationconfigParameters = &v
  return s
}

type QuerysourceverificationconfigPaths struct {
  // {"en":"The domain whoes need query config.", "zh_CN":"需要查询配置的域名或域名id"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
}

func (s QuerysourceverificationconfigPaths) String() string {
  return tea.Prettify(s)
}

func (s QuerysourceverificationconfigPaths) GoString() string {
  return s.String()
}

func (s *QuerysourceverificationconfigPaths) SetDomain(v string) *QuerysourceverificationconfigPaths {
  s.Domain = &v
  return s
}

type QuerysourceverificationconfigParameters struct {
}

func (s QuerysourceverificationconfigParameters) String() string {
  return tea.Prettify(s)
}

func (s QuerysourceverificationconfigParameters) GoString() string {
  return s.String()
}

type QuerysourceverificationconfigRequestHeader struct {
}

func (s QuerysourceverificationconfigRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QuerysourceverificationconfigRequestHeader) GoString() string {
  return s.String()
}

type QuerysourceverificationconfigResponseHeader struct {
}

func (s QuerysourceverificationconfigResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QuerysourceverificationconfigResponseHeader) GoString() string {
  return s.String()
}




type QueryIgnoreProtocolRequest struct {
}

func (s QueryIgnoreProtocolRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryIgnoreProtocolRequest) GoString() string {
  return s.String()
}

type QueryIgnoreProtocolResponse struct {
  // {"en":"Domain name or domain ID that needs to be queried for configuration.", "zh_CN":"需要查询配置的域名或域名id"}
  DomainName *string `json:"domain-name,omitempty" xml:"domain-name,omitempty" require:"true"`
  // {"en":"Domain name or domain ID that needs to be queried for configuration.", "zh_CN":"需要查询配置的域名或域名id"}
  DomainId *string `json:"domain-id,omitempty" xml:"domain-id,omitempty" require:"true"`
  // {"en":"Ignore protocol caching and push configuration, parent tags
  // 
  // 1. This must be filled when protocol cache and push configuration need to be ignored
  // 2.<ignore-protocol-rules/>:Clear the configuration ignore about protocol cache and pushing", "zh_CN":"忽略协议缓存和推送配置，父标签
  // 1.需要设置忽略协议缓存和推送配置时，此项必填
  // 2.为<ignore-protocol-rules/>时清空忽略协议缓存和推送的配置"}
  IgnoreProtocolRules []*QueryIgnoreProtocolResponseIgnoreProtocolRules `json:"ignore-protocol-rules,omitempty" xml:"ignore-protocol-rules,omitempty" require:"true" type:"Repeated"`
}

func (s QueryIgnoreProtocolResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryIgnoreProtocolResponse) GoString() string {
  return s.String()
}

func (s *QueryIgnoreProtocolResponse) SetDomainName(v string) *QueryIgnoreProtocolResponse {
  s.DomainName = &v
  return s
}

func (s *QueryIgnoreProtocolResponse) SetDomainId(v string) *QueryIgnoreProtocolResponse {
  s.DomainId = &v
  return s
}

func (s *QueryIgnoreProtocolResponse) SetIgnoreProtocolRules(v []*QueryIgnoreProtocolResponseIgnoreProtocolRules) *QueryIgnoreProtocolResponse {
  s.IgnoreProtocolRules = v
  return s
}

type QueryIgnoreProtocolResponseIgnoreProtocolRules struct     {
  // {"en":"Url matching pattern, support regular, if all matches, input parameters can be configured as:.*", "zh_CN":"url匹配模式，支持正则，如果是全部匹配，入参可以配置为：.*"}
  PathPattern *string `json:"path-pattern,omitempty" xml:"path-pattern,omitempty" require:"true"`
  // {"en":"The exception url matches the pattern in the same format as the path-pattern", "zh_CN":"例外的url匹配模式，格式同path-pattern"}
  ExceptPathPattern *string `json:"except-path-pattern,omitempty" xml:"except-path-pattern,omitempty" require:"true"`
  // {"en":"Whether to ignore the protocol cache, with allowable values of true and false. True turns on the HTTP/HTTPS Shared cache. Not on by default.", "zh_CN":"是否忽略协议缓存，允许值为true和false。为true则开启http/https共用缓存。默认不开启。"}
  CacheIgnoreProtocol *string `json:"cache-ignore-protocol,omitempty" xml:"cache-ignore-protocol,omitempty" require:"true"`
  // {"en":"It is recommended to use with cache-ignore protocol to avoid push failure.
  // 
  // Note:
  // 
  // 1. Once configured, the global effect is not applied to the matched path-pattern.
  // 
  // 2. Directory push does not distinguish protocols, while url push can distinguish protocols", "zh_CN":"是否忽略协议推送，允许值为true和false。为true则推送时忽略协议；为false则区分协议推送。
  // 建议和cache-ignore-protocol配套使用，避免推送失效。
  // 注意：
  // 1.一旦配置，则全局生效，不针对匹配的path-pattern生效。
  // 2.目录推送不区分协议，url推送可区分协议"}
  PurgeIgnoreProtocol *string `json:"purge-ignore-protocol,omitempty" xml:"purge-ignore-protocol,omitempty" require:"true"`
  // {"en":"When configuring multiple groups of configurations, specify the id of a particular group of configurations", "zh_CN":"配置多组配置时，具体某组配置的id"}
  DataId *int64 `json:"data-id,omitempty" xml:"data-id,omitempty" require:"true"`
}

func (s QueryIgnoreProtocolResponseIgnoreProtocolRules) String() string {
  return tea.Prettify(s)
}

func (s QueryIgnoreProtocolResponseIgnoreProtocolRules) GoString() string {
  return s.String()
}

func (s *QueryIgnoreProtocolResponseIgnoreProtocolRules) SetPathPattern(v string) *QueryIgnoreProtocolResponseIgnoreProtocolRules {
  s.PathPattern = &v
  return s
}

func (s *QueryIgnoreProtocolResponseIgnoreProtocolRules) SetExceptPathPattern(v string) *QueryIgnoreProtocolResponseIgnoreProtocolRules {
  s.ExceptPathPattern = &v
  return s
}

func (s *QueryIgnoreProtocolResponseIgnoreProtocolRules) SetCacheIgnoreProtocol(v string) *QueryIgnoreProtocolResponseIgnoreProtocolRules {
  s.CacheIgnoreProtocol = &v
  return s
}

func (s *QueryIgnoreProtocolResponseIgnoreProtocolRules) SetPurgeIgnoreProtocol(v string) *QueryIgnoreProtocolResponseIgnoreProtocolRules {
  s.PurgeIgnoreProtocol = &v
  return s
}

func (s *QueryIgnoreProtocolResponseIgnoreProtocolRules) SetDataId(v int64) *QueryIgnoreProtocolResponseIgnoreProtocolRules {
  s.DataId = &v
  return s
}

type QueryIgnoreProtocolPaths struct {
  // {"en":"Domain name or domain ID that needs to be queried for configuration.", "zh_CN":"需要查询配置的域名或域名id"}
  DomainName *string `json:"domain-name,omitempty" xml:"domain-name,omitempty" require:"true"`
}

func (s QueryIgnoreProtocolPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryIgnoreProtocolPaths) GoString() string {
  return s.String()
}

func (s *QueryIgnoreProtocolPaths) SetDomainName(v string) *QueryIgnoreProtocolPaths {
  s.DomainName = &v
  return s
}

type QueryIgnoreProtocolParameters struct {
}

func (s QueryIgnoreProtocolParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryIgnoreProtocolParameters) GoString() string {
  return s.String()
}

type QueryIgnoreProtocolRequestHeader struct {
}

func (s QueryIgnoreProtocolRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryIgnoreProtocolRequestHeader) GoString() string {
  return s.String()
}

type QueryIgnoreProtocolResponseHeader struct {
}

func (s QueryIgnoreProtocolResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryIgnoreProtocolResponseHeader) GoString() string {
  return s.String()
}




type EnableOrDisableWAFProtectionRequest struct {
  // {"en":"Domain names list, the parent tag.", "zh_CN":"开启/关闭WAF防护的域名列表， 父标签"}
  DomainNames []*string `json:"domainNames,omitempty" xml:"domainNames,omitempty" require:"true" type:"Repeated"`
  // {"en":"The reference  domain name. the reference domain cannot be null while enabling WAF protection", "zh_CN":"参考域名， 开启waf防护时，参考域名不能为空"}
  ReferDomain *string `json:"referDomain,omitempty" xml:"referDomain,omitempty" require:"true"`
  // {"en":"1: Enable WAF protection; 
  // 0: Disable WAF protection", "zh_CN":"1：开启WAF防护，0：关闭WAF防护"}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
}

func (s EnableOrDisableWAFProtectionRequest) String() string {
  return tea.Prettify(s)
}

func (s EnableOrDisableWAFProtectionRequest) GoString() string {
  return s.String()
}

func (s *EnableOrDisableWAFProtectionRequest) SetDomainNames(v []*string) *EnableOrDisableWAFProtectionRequest {
  s.DomainNames = v
  return s
}

func (s *EnableOrDisableWAFProtectionRequest) SetReferDomain(v string) *EnableOrDisableWAFProtectionRequest {
  s.ReferDomain = &v
  return s
}

func (s *EnableOrDisableWAFProtectionRequest) SetType(v string) *EnableOrDisableWAFProtectionRequest {
  s.Type = &v
  return s
}

type EnableOrDisableWAFProtectionResponse struct {
  // {"en":"Error code.
  //  0：successful", "zh_CN":"错误代码。 0：成功"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"The response message.
  // Response "Success" when calling API successful", "zh_CN":"响应信息，成功时为success"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"httpstatus=202; Indicates that the new domain API was successfully invoked, and the current deployment of the new domain can be viewed using x-cnc-request-id in the header", "zh_CN":"httpstatus=202;   表示成功调用新增域名接口，可使用header中的x-cnc-request-id查看当前新增域名的部署情况"}
  HttpStatus *int `json:"http status code,omitempty" xml:"http status code,omitempty" require:"true"`
  // {"en":"Uniquely identified id for querying tasks per request (for all API)", "zh_CN":"唯一标示的id，用于查询每次请求的任务 （适用全部接口）"}
  XCncRequestId *string `json:"x-cnc-request-id,omitempty" xml:"x-cnc-request-id,omitempty" require:"true"`
  // {"en":"The body of return data.", "zh_CN":"返回体数据"}
  Data *string `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s EnableOrDisableWAFProtectionResponse) String() string {
  return tea.Prettify(s)
}

func (s EnableOrDisableWAFProtectionResponse) GoString() string {
  return s.String()
}

func (s *EnableOrDisableWAFProtectionResponse) SetCode(v string) *EnableOrDisableWAFProtectionResponse {
  s.Code = &v
  return s
}

func (s *EnableOrDisableWAFProtectionResponse) SetMessage(v string) *EnableOrDisableWAFProtectionResponse {
  s.Message = &v
  return s
}

func (s *EnableOrDisableWAFProtectionResponse) SetHttpStatus(v int) *EnableOrDisableWAFProtectionResponse {
  s.HttpStatus = &v
  return s
}

func (s *EnableOrDisableWAFProtectionResponse) SetXCncRequestId(v string) *EnableOrDisableWAFProtectionResponse {
  s.XCncRequestId = &v
  return s
}

func (s *EnableOrDisableWAFProtectionResponse) SetData(v string) *EnableOrDisableWAFProtectionResponse {
  s.Data = &v
  return s
}

type EnableOrDisableWAFProtectionPaths struct {
}

func (s EnableOrDisableWAFProtectionPaths) String() string {
  return tea.Prettify(s)
}

func (s EnableOrDisableWAFProtectionPaths) GoString() string {
  return s.String()
}

type EnableOrDisableWAFProtectionParameters struct {
}

func (s EnableOrDisableWAFProtectionParameters) String() string {
  return tea.Prettify(s)
}

func (s EnableOrDisableWAFProtectionParameters) GoString() string {
  return s.String()
}

type EnableOrDisableWAFProtectionRequestHeader struct {
}

func (s EnableOrDisableWAFProtectionRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s EnableOrDisableWAFProtectionRequestHeader) GoString() string {
  return s.String()
}

type EnableOrDisableWAFProtectionResponseHeader struct {
}

func (s EnableOrDisableWAFProtectionResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s EnableOrDisableWAFProtectionResponseHeader) GoString() string {
  return s.String()
}




type UpdatesourceverificationconfigRequest struct {
  // {"en":"", "zh_CN":"转推回源带参数鉴权配置
  // 注意：
  // 1、需要取消花椒转推回源带参数鉴权配置时，可以传入空节点<source-verification></source-verification>。
  // 2、表示需要设置花椒转推回源带参数鉴权配置时，此项必填"}
  SourceVerification *UpdatesourceverificationconfigRequestSourceVerification `json:"source-verification,omitempty" xml:"source-verification,omitempty" require:"true" type:"Struct"`
}

func (s UpdatesourceverificationconfigRequest) String() string {
  return tea.Prettify(s)
}

func (s UpdatesourceverificationconfigRequest) GoString() string {
  return s.String()
}

func (s *UpdatesourceverificationconfigRequest) SetSourceVerification(v *UpdatesourceverificationconfigRequestSourceVerification) *UpdatesourceverificationconfigRequest {
  s.SourceVerification = v
  return s
}

type UpdatesourceverificationconfigRequestSourceVerification struct {
  // {"en":"", "zh_CN":"生效位置，代表配置生效的位置。可选值，多个分号隔开
  // Cache:边缘
  // stfuCache：静态中转
  // dyfuCache：动态中转"}
  SwitchPosition *string `json:"switch-position,omitempty" xml:"switch-position,omitempty" require:"true"`
  // {"en":"", "zh_CN":"是否开启回源校验。
  // 可选值：
  // close：关闭
  // push:转推使用
  // origin：回源使用
  // push-origin：即转推又回源使用"}
  RefSwitch *string `json:"ref-switch,omitempty" xml:"ref-switch,omitempty" require:"true"`
  // {"en":"", "zh_CN":"md5加密方式
  // 标识客户md5加密涉及的内容。参考入参
  // ||%S||BASE10||12345678||"}
  Md5Path *string `json:"md5-path,omitempty" xml:"md5-path,omitempty" require:"true"`
  // {"en":"", "zh_CN":"时间偏移量:未配置时以当前时间为准；配了时移就以当前时间+时移秒数的那个时间为准。不传默认为0，即无时间偏移"}
  TimeOffset *int `json:"time-offset,omitempty" xml:"time-offset,omitempty"`
  // {"en":"", "zh_CN":"sign字段名称：问号后需要鉴权的参数名
  // 参考入参：sign"}
  SignName *string `json:"sign-name,omitempty" xml:"sign-name,omitempty"`
  // {"en":"", "zh_CN":"time字段名称：问号后需要鉴权的参数名
  // 参考入参：time"}
  TimeName *string `json:"time-name,omitempty" xml:"time-name,omitempty"`
  // {"en":"", "zh_CN":"额外固定参数：需要带额外的参数回客户源站鉴权
  // 参考入参：trans=live.a.com"}
  OtherUpdatesourceverificationconfigParameters *string `json:"other-parameters,omitempty" xml:"other-parameters,omitempty"`
}

func (s UpdatesourceverificationconfigRequestSourceVerification) String() string {
  return tea.Prettify(s)
}

func (s UpdatesourceverificationconfigRequestSourceVerification) GoString() string {
  return s.String()
}

func (s *UpdatesourceverificationconfigRequestSourceVerification) SetSwitchPosition(v string) *UpdatesourceverificationconfigRequestSourceVerification {
  s.SwitchPosition = &v
  return s
}

func (s *UpdatesourceverificationconfigRequestSourceVerification) SetRefSwitch(v string) *UpdatesourceverificationconfigRequestSourceVerification {
  s.RefSwitch = &v
  return s
}

func (s *UpdatesourceverificationconfigRequestSourceVerification) SetMd5Path(v string) *UpdatesourceverificationconfigRequestSourceVerification {
  s.Md5Path = &v
  return s
}

func (s *UpdatesourceverificationconfigRequestSourceVerification) SetTimeOffset(v int) *UpdatesourceverificationconfigRequestSourceVerification {
  s.TimeOffset = &v
  return s
}

func (s *UpdatesourceverificationconfigRequestSourceVerification) SetSignName(v string) *UpdatesourceverificationconfigRequestSourceVerification {
  s.SignName = &v
  return s
}

func (s *UpdatesourceverificationconfigRequestSourceVerification) SetTimeName(v string) *UpdatesourceverificationconfigRequestSourceVerification {
  s.TimeName = &v
  return s
}

func (s *UpdatesourceverificationconfigRequestSourceVerification) SetOtherUpdatesourceverificationconfigParameters(v string) *UpdatesourceverificationconfigRequestSourceVerification {
  s.OtherUpdatesourceverificationconfigParameters = &v
  return s
}

type UpdatesourceverificationconfigResponse struct {
  // {"en":"httpstatus=202; Indicates that the new domain API was successfully invoked, and the current deployment of the new domain can be viewed using x-cnc-request-id in the header", "zh_CN":"httpstatus=202;   表示成功调用新增域名接口，可使用header中的x-cnc-request-id查看当前新增域名的部署情况"}
  HttpStatus *int `json:"http status code,omitempty" xml:"http status code,omitempty" require:"true"`
  // {"en":"Uniquely identified id for querying tasks per request (for all API)", "zh_CN":"唯一标示的id，用于查询每次请求的任务 （适用全部接口）"}
  XCncRequestId *string `json:"x-cnc-request-id,omitempty" xml:"x-cnc-request-id,omitempty" require:"true"`
  // {"en":"The error code, when HTTPStatus is not 202, indicates the type of error the current request is calling.", "zh_CN":"错误代码，当HTTPStatus不为202时出现，表示当前请求调用的错误类型
  // 错误代码，当HTTPStatus不为202时出现，表示当前请求调用的错误类型"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response information, when success is successful", "zh_CN":"响应信息，成功时为success
  // 响应信息，成功时为success"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s UpdatesourceverificationconfigResponse) String() string {
  return tea.Prettify(s)
}

func (s UpdatesourceverificationconfigResponse) GoString() string {
  return s.String()
}

func (s *UpdatesourceverificationconfigResponse) SetHttpStatus(v int) *UpdatesourceverificationconfigResponse {
  s.HttpStatus = &v
  return s
}

func (s *UpdatesourceverificationconfigResponse) SetXCncRequestId(v string) *UpdatesourceverificationconfigResponse {
  s.XCncRequestId = &v
  return s
}

func (s *UpdatesourceverificationconfigResponse) SetCode(v string) *UpdatesourceverificationconfigResponse {
  s.Code = &v
  return s
}

func (s *UpdatesourceverificationconfigResponse) SetMessage(v string) *UpdatesourceverificationconfigResponse {
  s.Message = &v
  return s
}

type UpdatesourceverificationconfigPaths struct {
  // {"en":"The domain whoes need query config.", "zh_CN":"需要查询配置的域名或域名id"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
}

func (s UpdatesourceverificationconfigPaths) String() string {
  return tea.Prettify(s)
}

func (s UpdatesourceverificationconfigPaths) GoString() string {
  return s.String()
}

func (s *UpdatesourceverificationconfigPaths) SetDomain(v string) *UpdatesourceverificationconfigPaths {
  s.Domain = &v
  return s
}

type UpdatesourceverificationconfigParameters struct {
}

func (s UpdatesourceverificationconfigParameters) String() string {
  return tea.Prettify(s)
}

func (s UpdatesourceverificationconfigParameters) GoString() string {
  return s.String()
}

type UpdatesourceverificationconfigRequestHeader struct {
}

func (s UpdatesourceverificationconfigRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s UpdatesourceverificationconfigRequestHeader) GoString() string {
  return s.String()
}

type UpdatesourceverificationconfigResponseHeader struct {
}

func (s UpdatesourceverificationconfigResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s UpdatesourceverificationconfigResponseHeader) GoString() string {
  return s.String()
}




type QueryCompressionConfigRequest struct {
}

func (s QueryCompressionConfigRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryCompressionConfigRequest) GoString() string {
  return s.String()
}

type QueryCompressionConfigResponse struct {
  // {"en":"The domain id you are query.", "zh_CN":"需要查询配置的域名id"}
  DomainId *string `json:"domain-id,omitempty" xml:"domain-id,omitempty" require:"true"`
  // {"en":"The domain name you are query.", "zh_CN":"需要查询配置的域名"}
  DomainName *string `json:"domain-name,omitempty" xml:"domain-name,omitempty" require:"true"`
  // {"en":"Compress setting config", "zh_CN":"压缩响应功能配置
  // 1.需要设置压缩响应配置时，此项必填
  // 2.为<compression-settings/>空时清空压缩响应配置"}
  CompressionSettings *QueryCompressionConfigResponseCompressionSettings `json:"compression-settings,omitempty" xml:"compression-settings,omitempty" require:"true" type:"Struct"`
}

func (s QueryCompressionConfigResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryCompressionConfigResponse) GoString() string {
  return s.String()
}

func (s *QueryCompressionConfigResponse) SetDomainId(v string) *QueryCompressionConfigResponse {
  s.DomainId = &v
  return s
}

func (s *QueryCompressionConfigResponse) SetDomainName(v string) *QueryCompressionConfigResponse {
  s.DomainName = &v
  return s
}

func (s *QueryCompressionConfigResponse) SetCompressionSettings(v *QueryCompressionConfigResponseCompressionSettings) *QueryCompressionConfigResponse {
  s.CompressionSettings = v
  return s
}

type QueryCompressionConfigResponseCompressionSettings struct {
  // {"en":"To enable compress setting, allowed true or false.", "zh_CN":"开启压缩响应功能：允许值为true和false"}
  CompressionEnabled *string `json:"compression-enabled,omitempty" xml:"compression-enabled,omitempty" require:"true"`
  // {"en":"The url matching mode. If all matches, the input parameters can be configured as: .*", "zh_CN":"url匹配模式，支持正则，如果是全部匹配，入参可以配置为：.*"}
  PathPattern *string `json:"path-pattern,omitempty" xml:"path-pattern,omitempty" require:"true"`
  // {"en":"Whether to ignore letter case.", "zh_CN":"是否忽略大小写：允许值为true和false"}
  IgnoreLetterCase *string `json:"ignore-letter-case,omitempty" xml:"ignore-letter-case,omitempty" require:"true"`
  // {"en":"Define the file types to be compressed. 'text/' will be compressed by default.", "zh_CN":"配置需要压缩的文件类型，默认只对'text'文件类型压缩，配置为*时压缩任意文件类型"}
  FileTypes []*string `json:"file-types,omitempty" xml:"file-types,omitempty" require:"true" type:"Repeated"`
  // {"en":"Another way to specify the file type to open the compressed response, <file-types-other/> can be cleared and the configuration is the parent tag of file-type-other", "zh_CN":"指定文件类型开启压缩响应的另一种方式，<file-types-other/>可清空配置，是file-type-other的父标签"}
  FileTypeOthers []*string `json:"file-type-others,omitempty" xml:"file-type-others,omitempty" require:"true" type:"Repeated"`
  // {"en":"Use br compression.The allowed values are true and false.", "zh_CN":"是否使用br压缩：允许值为true和false"}
  BrTypes *string `json:"br-types,omitempty" xml:"br-types,omitempty"`
}

func (s QueryCompressionConfigResponseCompressionSettings) String() string {
  return tea.Prettify(s)
}

func (s QueryCompressionConfigResponseCompressionSettings) GoString() string {
  return s.String()
}

func (s *QueryCompressionConfigResponseCompressionSettings) SetCompressionEnabled(v string) *QueryCompressionConfigResponseCompressionSettings {
  s.CompressionEnabled = &v
  return s
}

func (s *QueryCompressionConfigResponseCompressionSettings) SetPathPattern(v string) *QueryCompressionConfigResponseCompressionSettings {
  s.PathPattern = &v
  return s
}

func (s *QueryCompressionConfigResponseCompressionSettings) SetIgnoreLetterCase(v string) *QueryCompressionConfigResponseCompressionSettings {
  s.IgnoreLetterCase = &v
  return s
}

func (s *QueryCompressionConfigResponseCompressionSettings) SetFileTypes(v []*string) *QueryCompressionConfigResponseCompressionSettings {
  s.FileTypes = v
  return s
}

func (s *QueryCompressionConfigResponseCompressionSettings) SetFileTypeOthers(v []*string) *QueryCompressionConfigResponseCompressionSettings {
  s.FileTypeOthers = v
  return s
}

func (s *QueryCompressionConfigResponseCompressionSettings) SetBrTypes(v string) *QueryCompressionConfigResponseCompressionSettings {
  s.BrTypes = &v
  return s
}

type QueryCompressionConfigPaths struct {
  // {"en":"The domain you want to query, support domain id and domain name.", "zh_CN":"需要查询配置的域名或域名id"}
  DomainName *string `json:"domain-name,omitempty" xml:"domain-name,omitempty"`
}

func (s QueryCompressionConfigPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryCompressionConfigPaths) GoString() string {
  return s.String()
}

func (s *QueryCompressionConfigPaths) SetDomainName(v string) *QueryCompressionConfigPaths {
  s.DomainName = &v
  return s
}

type QueryCompressionConfigParameters struct {
}

func (s QueryCompressionConfigParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryCompressionConfigParameters) GoString() string {
  return s.String()
}

type QueryCompressionConfigRequestHeader struct {
}

func (s QueryCompressionConfigRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryCompressionConfigRequestHeader) GoString() string {
  return s.String()
}

type QueryCompressionConfigResponseHeader struct {
}

func (s QueryCompressionConfigResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryCompressionConfigResponseHeader) GoString() string {
  return s.String()
}




type QueryHttpCodeCasheConfigRequest struct {
}

func (s QueryHttpCodeCasheConfigRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryHttpCodeCasheConfigRequest) GoString() string {
  return s.String()
}

type QueryHttpCodeCasheConfigResponse struct {
  // {"en":"Domain name or domain ID that needs to be queried for configuration.", "zh_CN":"需要查询配置的域名"}
  DomainName *string `json:"domain-name,omitempty" xml:"domain-name,omitempty" require:"true"`
  // {"en":"Domain name or domain ID that needs to be queried for configuration.", "zh_CN":"需要查询配置域名id"}
  DomainId *string `json:"domain-id,omitempty" xml:"domain-id,omitempty" require:"true"`
  // {"en":"State Code Caching Rule Configuration, parent node
  // 1. When you need to set state code caching rules, this must be filled in.
  // 2. Configuration of Clear State Code Caching Rules for <http-code-cache-rules/>.", "zh_CN":"状态码缓存规则配置，父标签
  // 1.需要设置状态码缓存规则时，此项必填
  // 2.为<http-code-cache-rules/>时清空状态码缓存规则配置"}
  HttpCodeCacheRules []*QueryHttpCodeCasheConfigResponseHttpCodeCacheRules `json:"http-code-cache-rules,omitempty" xml:"http-code-cache-rules,omitempty" require:"true" type:"Repeated"`
}

func (s QueryHttpCodeCasheConfigResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryHttpCodeCasheConfigResponse) GoString() string {
  return s.String()
}

func (s *QueryHttpCodeCasheConfigResponse) SetDomainName(v string) *QueryHttpCodeCasheConfigResponse {
  s.DomainName = &v
  return s
}

func (s *QueryHttpCodeCasheConfigResponse) SetDomainId(v string) *QueryHttpCodeCasheConfigResponse {
  s.DomainId = &v
  return s
}

func (s *QueryHttpCodeCasheConfigResponse) SetHttpCodeCacheRules(v []*QueryHttpCodeCasheConfigResponseHttpCodeCacheRules) *QueryHttpCodeCasheConfigResponse {
  s.HttpCodeCacheRules = v
  return s
}

type QueryHttpCodeCasheConfigResponseHttpCodeCacheRules struct     {
  // {"en":"Configure HTTP status code, parent node", "zh_CN":"配置http状态码，父标签"}
  HttpCodes []*int32 `json:"http-codes,omitempty" xml:"http-codes,omitempty" require:"true" type:"Repeated"`
  // {"en":"Define the caching time of the specified status code in units s, 0 to indicate no caching", "zh_CN":"配置指定的状态码的缓存时间，单位s，0表示不缓存"}
  CacheTtl *int `json:"cache-ttl,omitempty" xml:"cache-ttl,omitempty" require:"true"`
  // {"en":"Add a grid type identifier to indicate a specific group configuration when the client has multiple groups of configurations.", "zh_CN":"配置多组配置时，具体某组配置的id"}
  DataId *int `json:"data-id,omitempty" xml:"data-id,omitempty" require:"true"`
}

func (s QueryHttpCodeCasheConfigResponseHttpCodeCacheRules) String() string {
  return tea.Prettify(s)
}

func (s QueryHttpCodeCasheConfigResponseHttpCodeCacheRules) GoString() string {
  return s.String()
}

func (s *QueryHttpCodeCasheConfigResponseHttpCodeCacheRules) SetHttpCodes(v []*int32) *QueryHttpCodeCasheConfigResponseHttpCodeCacheRules {
  s.HttpCodes = v
  return s
}

func (s *QueryHttpCodeCasheConfigResponseHttpCodeCacheRules) SetCacheTtl(v int) *QueryHttpCodeCasheConfigResponseHttpCodeCacheRules {
  s.CacheTtl = &v
  return s
}

func (s *QueryHttpCodeCasheConfigResponseHttpCodeCacheRules) SetDataId(v int) *QueryHttpCodeCasheConfigResponseHttpCodeCacheRules {
  s.DataId = &v
  return s
}

type QueryHttpCodeCasheConfigPaths struct {
  // {"en":"Domain name or domain ID that needs to be queried for configuration.", "zh_CN":"需要查询配置的域名或域名id"}
  DomainName *string `json:"domain-name,omitempty" xml:"domain-name,omitempty" require:"true"`
}

func (s QueryHttpCodeCasheConfigPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryHttpCodeCasheConfigPaths) GoString() string {
  return s.String()
}

func (s *QueryHttpCodeCasheConfigPaths) SetDomainName(v string) *QueryHttpCodeCasheConfigPaths {
  s.DomainName = &v
  return s
}

type QueryHttpCodeCasheConfigParameters struct {
}

func (s QueryHttpCodeCasheConfigParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryHttpCodeCasheConfigParameters) GoString() string {
  return s.String()
}

type QueryHttpCodeCasheConfigRequestHeader struct {
}

func (s QueryHttpCodeCasheConfigRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryHttpCodeCasheConfigRequestHeader) GoString() string {
  return s.String()
}

type QueryHttpCodeCasheConfigResponseHeader struct {
}

func (s QueryHttpCodeCasheConfigResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryHttpCodeCasheConfigResponseHeader) GoString() string {
  return s.String()
}




type UpdateStreamNotificationConfigRequest struct {
  // {"en":"Streaming notify Configuration
  // 1. It is requied when you need to set streaming notify rules.
  // 2. Use \"streamNotifications\":[] to clear Streaming notify Configuration.", "zh_CN":"推流状态反馈配置
  // 1.需要设置推流状态反馈配置时，此项必填
  // 2.为\"streamNotifications\":[]时清空配置"}
  StreamNotifications []*UpdateStreamNotificationConfigRequestStreamNotifications `json:"streamNotifications,omitempty" xml:"streamNotifications,omitempty" require:"true" type:"Repeated"`
}

func (s UpdateStreamNotificationConfigRequest) String() string {
  return tea.Prettify(s)
}

func (s UpdateStreamNotificationConfigRequest) GoString() string {
  return s.String()
}

func (s *UpdateStreamNotificationConfigRequest) SetStreamNotifications(v []*UpdateStreamNotificationConfigRequestStreamNotifications) *UpdateStreamNotificationConfigRequest {
  s.StreamNotifications = v
  return s
}

type UpdateStreamNotificationConfigRequestStreamNotifications struct     {
  // {"en":"Switch for streaming notification feature. The optional values are true and false.", "zh_CN":"推流状态反馈开关配置，可选值为true和false。
  // 为true则开启，且需完整设置配置信息
  // 为false则其他入参无效。"}
  EnableStreamNotification *bool `json:"enableStreamNotification,omitempty" xml:"enableStreamNotification,omitempty" require:"true"`
  // {"en":"Notify address, format is protocol://{host:port}, the protocol is HTTP or HTTPS, for example:
  // http://www.exam.com 
  // http://www.exam.com:81
  // http://12.12.12.12
  // http://12.12.12.12:81", "zh_CN":"推流状态汇报地址，格式为：protocol://{host:port}，例如：
  // http://www.exam.com 
  // http://www.exam.com:81
  // http://12.12.12.12
  // http://12.12.12.12:81"}
  NotifyAddress *string `json:"notifyAddress,omitempty" xml:"notifyAddress,omitempty" require:"true"`
  // {"en":"Request method, support POST and GET.", "zh_CN":"推流汇报请求方式，支持POST和GET"}
  NotifyMethod *string `json:"notifyMethod,omitempty" xml:"notifyMethod,omitempty"`
  // {"en":"notifyParams", "zh_CN":"设置推流开始的汇报参数，推流结束的汇报参数。最多只能有一个开始，一个结束。"}
  NotifyParams []*UpdateStreamNotificationConfigRequestStreamNotificationsNotifyParams `json:"notifyParams,omitempty" xml:"notifyParams,omitempty" require:"true" type:"Repeated"`
  // {"en":"Data-id is to indicate a specific group configuration when the client has multiple groups of configurations. Data-id can be retrieved through a query interface. Note: A. If data-id is passed, it means that one group of configuration items is specified to be modified, and no other group configuration items need to be modified. B. If multiple groups of configurations are included, some of them are configured with data-id and others are not, then the expression of data-id is used to modify a specific group of configurations, and a new group of configurations is added on the original basis without the expression of data-id. C. If the data-id is not transmitted, it means that the original configuration will be fully covered by this configuration. D. If no configuration parameter is passed, only domain name and secondary label are passed, which means that all configuration of domain name secondary service corresponding to this interface is cleared. E. If there is no specific configuration item in a set of configurations, the data-id must be filled in, and the value is the actual data-id, which means clearing the value of the corresponding data-id configuration item; it is not allowed that there is no specific configuration item or data-id in a set of configurations.", "zh_CN":"配置多组配置时，具体某组配置的id。dataId可以通过查询接口获取。 注意： a、如果有传dataId，说明指定修改其中一组配置项内容，不需求修改其他组配置内容不需要入参； b、如果入参多组配置，其中有些组配置有传dataId，有些没有传，则有传dataId的表示修改具体某组配置，没有传dataId的表示在原来基础上新增一组配置； c、如果入参都没有传dataId,表示用本次的配置全量覆盖原先配置； d、如果入参没有传任何配置项参数，只传了域名和二级标签，表示清空这个接口对应域名二级服务所有配置； e、如果一组配置没有具体的配置项，则dataId必填，且值为实际存在的dataId，表示清空这个dataId对应配置项的值；不允许一组配置没有具体的配置项也没有dataId。"}
  DataId *int64 `json:"dataId,omitempty" xml:"dataId,omitempty"`
}

func (s UpdateStreamNotificationConfigRequestStreamNotifications) String() string {
  return tea.Prettify(s)
}

func (s UpdateStreamNotificationConfigRequestStreamNotifications) GoString() string {
  return s.String()
}

func (s *UpdateStreamNotificationConfigRequestStreamNotifications) SetEnableStreamNotification(v bool) *UpdateStreamNotificationConfigRequestStreamNotifications {
  s.EnableStreamNotification = &v
  return s
}

func (s *UpdateStreamNotificationConfigRequestStreamNotifications) SetNotifyAddress(v string) *UpdateStreamNotificationConfigRequestStreamNotifications {
  s.NotifyAddress = &v
  return s
}

func (s *UpdateStreamNotificationConfigRequestStreamNotifications) SetNotifyMethod(v string) *UpdateStreamNotificationConfigRequestStreamNotifications {
  s.NotifyMethod = &v
  return s
}

func (s *UpdateStreamNotificationConfigRequestStreamNotifications) SetNotifyParams(v []*UpdateStreamNotificationConfigRequestStreamNotificationsNotifyParams) *UpdateStreamNotificationConfigRequestStreamNotifications {
  s.NotifyParams = v
  return s
}

func (s *UpdateStreamNotificationConfigRequestStreamNotifications) SetDataId(v int64) *UpdateStreamNotificationConfigRequestStreamNotifications {
  s.DataId = &v
  return s
}

type UpdateStreamNotificationConfigRequestStreamNotificationsNotifyParams struct     {
  // {"en":"notifyType", "zh_CN":"可选值：publish_start、publish_end，分别表示推流开始和推流结束(断流)"}
  NotifyType *string `json:"notifyType,omitempty" xml:"notifyType,omitempty" require:"true"`
  // {"en":"Notify path, it should start with a slash.", "zh_CN":"推流汇报路径，/ 开头"}
  Path *string `json:"path,omitempty" xml:"path,omitempty" require:"true"`
  // {"en":"Stream name, the default is 'id', cannot be cleared.", "zh_CN":"流名，不传默认为id，不可清空"}
  Id *string `json:"id,omitempty" xml:"id,omitempty"`
  // {"en":"Streaming client IP, anchor IP, the default is 'ip', cannot be cleared.", "zh_CN":"推流的客户端IP，主播IP，不传默认为ip，不可清空"}
  ClientIp *string `json:"clientIp,omitempty" xml:"clientIp,omitempty"`
  // {"en":"The IP of the node where the CDN accepts the streaming, the default is 'node', cannot be cleared.", "zh_CN":"CDN接受推流的节点IP，不传默认为node，不可清空"}
  NodeIp *string `json:"nodeIp,omitempty" xml:"nodeIp,omitempty"`
  // {"en":"Push domain, the default is 'app', cannot be cleared.", "zh_CN":"推流域名，不传默认为app，不可清空"}
  App *string `json:"app,omitempty" xml:"app,omitempty"`
  // {"en":"AppName, the default is 'appname', cannot be cleared.", "zh_CN":"发布点，不传默认为appname，不可清空"}
  AppName *string `json:"appName,omitempty" xml:"appName,omitempty"`
  // {"en":"Port, the default is 'port', cannot be cleared.", "zh_CN":"端口，不传默认为port，不可清空"}
  Port *string `json:"port,omitempty" xml:"port,omitempty"`
  // {"en":"Report time (milliseconds), the default is 'milltime', cannot be cleared.", "zh_CN":"汇报时间（毫秒级），不传默认为milltime，不可清空"}
  MillTime *string `json:"millTime,omitempty" xml:"millTime,omitempty"`
  // {"en":"Specify the parameters to be carried in the push report, which come from the user push parameters, and these parameters are not carried by default. Separated by semicolons.", "zh_CN":"指定推流汇报要携带的参数，来自用户推流参数，默认不携带这些参数。多个按分号分隔。"}
  CustomParams *string `json:"customParams,omitempty" xml:"customParams,omitempty"`
}

func (s UpdateStreamNotificationConfigRequestStreamNotificationsNotifyParams) String() string {
  return tea.Prettify(s)
}

func (s UpdateStreamNotificationConfigRequestStreamNotificationsNotifyParams) GoString() string {
  return s.String()
}

func (s *UpdateStreamNotificationConfigRequestStreamNotificationsNotifyParams) SetNotifyType(v string) *UpdateStreamNotificationConfigRequestStreamNotificationsNotifyParams {
  s.NotifyType = &v
  return s
}

func (s *UpdateStreamNotificationConfigRequestStreamNotificationsNotifyParams) SetPath(v string) *UpdateStreamNotificationConfigRequestStreamNotificationsNotifyParams {
  s.Path = &v
  return s
}

func (s *UpdateStreamNotificationConfigRequestStreamNotificationsNotifyParams) SetId(v string) *UpdateStreamNotificationConfigRequestStreamNotificationsNotifyParams {
  s.Id = &v
  return s
}

func (s *UpdateStreamNotificationConfigRequestStreamNotificationsNotifyParams) SetClientIp(v string) *UpdateStreamNotificationConfigRequestStreamNotificationsNotifyParams {
  s.ClientIp = &v
  return s
}

func (s *UpdateStreamNotificationConfigRequestStreamNotificationsNotifyParams) SetNodeIp(v string) *UpdateStreamNotificationConfigRequestStreamNotificationsNotifyParams {
  s.NodeIp = &v
  return s
}

func (s *UpdateStreamNotificationConfigRequestStreamNotificationsNotifyParams) SetApp(v string) *UpdateStreamNotificationConfigRequestStreamNotificationsNotifyParams {
  s.App = &v
  return s
}

func (s *UpdateStreamNotificationConfigRequestStreamNotificationsNotifyParams) SetAppName(v string) *UpdateStreamNotificationConfigRequestStreamNotificationsNotifyParams {
  s.AppName = &v
  return s
}

func (s *UpdateStreamNotificationConfigRequestStreamNotificationsNotifyParams) SetPort(v string) *UpdateStreamNotificationConfigRequestStreamNotificationsNotifyParams {
  s.Port = &v
  return s
}

func (s *UpdateStreamNotificationConfigRequestStreamNotificationsNotifyParams) SetMillTime(v string) *UpdateStreamNotificationConfigRequestStreamNotificationsNotifyParams {
  s.MillTime = &v
  return s
}

func (s *UpdateStreamNotificationConfigRequestStreamNotificationsNotifyParams) SetCustomParams(v string) *UpdateStreamNotificationConfigRequestStreamNotificationsNotifyParams {
  s.CustomParams = &v
  return s
}

type UpdateStreamNotificationConfigResponse struct {
  // {"en":"code", "zh_CN":"错误代码，当HTTPStatus不为202时出现，表示当前请求调用的错误类型"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"message", "zh_CN":"响应信息，成功时为success"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s UpdateStreamNotificationConfigResponse) String() string {
  return tea.Prettify(s)
}

func (s UpdateStreamNotificationConfigResponse) GoString() string {
  return s.String()
}

func (s *UpdateStreamNotificationConfigResponse) SetCode(v string) *UpdateStreamNotificationConfigResponse {
  s.Code = &v
  return s
}

func (s *UpdateStreamNotificationConfigResponse) SetMessage(v string) *UpdateStreamNotificationConfigResponse {
  s.Message = &v
  return s
}

type UpdateStreamNotificationConfigPaths struct {
  // {"en":"The domain whoes need to update.", "zh_CN":"需要修改的域名或域名id"}
  DomainName *string `json:"domain-name,omitempty" xml:"domain-name,omitempty" require:"true"`
}

func (s UpdateStreamNotificationConfigPaths) String() string {
  return tea.Prettify(s)
}

func (s UpdateStreamNotificationConfigPaths) GoString() string {
  return s.String()
}

func (s *UpdateStreamNotificationConfigPaths) SetDomainName(v string) *UpdateStreamNotificationConfigPaths {
  s.DomainName = &v
  return s
}

type UpdateStreamNotificationConfigParameters struct {
}

func (s UpdateStreamNotificationConfigParameters) String() string {
  return tea.Prettify(s)
}

func (s UpdateStreamNotificationConfigParameters) GoString() string {
  return s.String()
}

type UpdateStreamNotificationConfigRequestHeader struct {
}

func (s UpdateStreamNotificationConfigRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s UpdateStreamNotificationConfigRequestHeader) GoString() string {
  return s.String()
}

type UpdateStreamNotificationConfigResponseHeader struct {
}

func (s UpdateStreamNotificationConfigResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s UpdateStreamNotificationConfigResponseHeader) GoString() string {
  return s.String()
}




type QueryHttp2SettingsConfigForWplusRequest struct {
}

func (s QueryHttp2SettingsConfigForWplusRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryHttp2SettingsConfigForWplusRequest) GoString() string {
  return s.String()
}

type QueryHttp2SettingsConfigForWplusResponse struct {
  // {"en":"Domain name or domain name id to query configuration", "zh_CN":"需要查询配置的域名（domainName）"}
  DomainName *string `json:"domainName,omitempty" xml:"domainName,omitempty" require:"true"`
  // {"en":"Domain name or domain name id to query configuration", "zh_CN":"需要查询配置的域名id（domainId）"}
  DomainId *string `json:"domainId,omitempty" xml:"domainId,omitempty" require:"true"`
  // {"en":"Http2.0 settings, used to enable or disable http2.0, parent node.", "zh_CN":"http2.0设置，用于设置http2.0的开启或关闭，父标签"}
  Http2Settings *QueryHttp2SettingsConfigForWplusResponseHttp2Settings `json:"http2Settings,omitempty" xml:"http2Settings,omitempty" require:"true" type:"Struct"`
}

func (s QueryHttp2SettingsConfigForWplusResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryHttp2SettingsConfigForWplusResponse) GoString() string {
  return s.String()
}

func (s *QueryHttp2SettingsConfigForWplusResponse) SetDomainName(v string) *QueryHttp2SettingsConfigForWplusResponse {
  s.DomainName = &v
  return s
}

func (s *QueryHttp2SettingsConfigForWplusResponse) SetDomainId(v string) *QueryHttp2SettingsConfigForWplusResponse {
  s.DomainId = &v
  return s
}

func (s *QueryHttp2SettingsConfigForWplusResponse) SetHttp2Settings(v *QueryHttp2SettingsConfigForWplusResponseHttp2Settings) *QueryHttp2SettingsConfigForWplusResponse {
  s.Http2Settings = v
  return s
}

type QueryHttp2SettingsConfigForWplusResponseHttp2Settings struct {
  // {"en":"Enable http2.0. The optional values are true and false. If it is empty, the default value is false. True means http2.0 is on; false means http2.0 is off.", "zh_CN":"开启http2.0，可选值为true和false，为空时默认为false。true表示开启http2.0；false表示关闭http2.0"}
  EnableHttp2 *bool `json:"enableHttp2,omitempty" xml:"enableHttp2,omitempty" require:"true"`
  // {"en":"Back-to-origin protocol, the optional value is
  // http1.1: Use the HTTP1.1 protocol version to back to source. if not filled, use it as default.
  // follow-request: Same as client request protocol
  // http2.0: Use the HTTP2.0 protocol. version to back to source.", "zh_CN":"回源协议，可选值为
  // http1.1：使用HTTP1.1协议版本回源，不填时默认该协议
  // follow-request：跟随客户端请求协议
  // http2.0：使用HTTP2.0协议版本回源"}
  BackToOriginProtocol *string `json:"backToOriginProtocol,omitempty" xml:"backToOriginProtocol,omitempty" require:"true"`
}

func (s QueryHttp2SettingsConfigForWplusResponseHttp2Settings) String() string {
  return tea.Prettify(s)
}

func (s QueryHttp2SettingsConfigForWplusResponseHttp2Settings) GoString() string {
  return s.String()
}

func (s *QueryHttp2SettingsConfigForWplusResponseHttp2Settings) SetEnableHttp2(v bool) *QueryHttp2SettingsConfigForWplusResponseHttp2Settings {
  s.EnableHttp2 = &v
  return s
}

func (s *QueryHttp2SettingsConfigForWplusResponseHttp2Settings) SetBackToOriginProtocol(v string) *QueryHttp2SettingsConfigForWplusResponseHttp2Settings {
  s.BackToOriginProtocol = &v
  return s
}

type QueryHttp2SettingsConfigForWplusPaths struct {
  // {"en":"Domain name or domain name id to query configuration", "zh_CN":"需要查询配置的域名（domainName）或域名id（domainId）"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
}

func (s QueryHttp2SettingsConfigForWplusPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryHttp2SettingsConfigForWplusPaths) GoString() string {
  return s.String()
}

func (s *QueryHttp2SettingsConfigForWplusPaths) SetDomain(v string) *QueryHttp2SettingsConfigForWplusPaths {
  s.Domain = &v
  return s
}

type QueryHttp2SettingsConfigForWplusParameters struct {
}

func (s QueryHttp2SettingsConfigForWplusParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryHttp2SettingsConfigForWplusParameters) GoString() string {
  return s.String()
}

type QueryHttp2SettingsConfigForWplusRequestHeader struct {
}

func (s QueryHttp2SettingsConfigForWplusRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryHttp2SettingsConfigForWplusRequestHeader) GoString() string {
  return s.String()
}

type QueryHttp2SettingsConfigForWplusResponseHeader struct {
}

func (s QueryHttp2SettingsConfigForWplusResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryHttp2SettingsConfigForWplusResponseHeader) GoString() string {
  return s.String()
}




type EditIgnoreProtocolRequest struct {
  // {"en":"Ignore protocol caching and push configuration, parent tags
  // 
  // 1. This must be filled when protocol cache and push configuration need to be ignored
  // 2.<ignore-protocol-rules/>:Clear the configuration ignore about protocol cache and pushing", "zh_CN":"忽略协议缓存和推送配置，父标签
  // 1.需要设置忽略协议缓存和推送配置时，此项必填
  // 2.为<ignore-protocol-rules/>时清空忽略协议缓存和推送的配置"}
  IgnoreProtocolRules []*EditIgnoreProtocolRequestIgnoreProtocolRules `json:"ignore-protocol-rules,omitempty" xml:"ignore-protocol-rules,omitempty" require:"true" type:"Repeated"`
}

func (s EditIgnoreProtocolRequest) String() string {
  return tea.Prettify(s)
}

func (s EditIgnoreProtocolRequest) GoString() string {
  return s.String()
}

func (s *EditIgnoreProtocolRequest) SetIgnoreProtocolRules(v []*EditIgnoreProtocolRequestIgnoreProtocolRules) *EditIgnoreProtocolRequest {
  s.IgnoreProtocolRules = v
  return s
}

type EditIgnoreProtocolRequestIgnoreProtocolRules struct     {
  // {"en":"Url matching pattern, support regular, if all matches, input parameters can be configured as:.*", "zh_CN":"url匹配模式，支持正则，如果是全部匹配，入参可以配置为：.*"}
  PathPattern *string `json:"path-pattern,omitempty" xml:"path-pattern,omitempty" require:"true"`
  // {"en":"The exception url matches the pattern in the same format as the path-pattern", "zh_CN":"例外的url匹配模式，格式同path-pattern"}
  ExceptPathPattern *string `json:"except-path-pattern,omitempty" xml:"except-path-pattern,omitempty"`
  // {"en":"Whether to ignore the protocol cache, with allowable values of true and false. True turns on the HTTP/HTTPS Shared cache. Not on by default.", "zh_CN":"是否忽略协议缓存，允许值为true和false。为true则开启http/https共用缓存。默认不开启。"}
  CacheIgnoreProtocol *string `json:"cache-ignore-protocol,omitempty" xml:"cache-ignore-protocol,omitempty"`
  // {"en":"It is recommended to use with cache-ignore protocol to avoid push failure.
  // 
  // Note:
  // 
  // 1. Once configured, the global effect is not applied to the matched path-pattern.
  // 
  // 2. Directory push does not distinguish protocols, while url push can distinguish protocols", "zh_CN":"是否忽略协议推送，允许值为true和false。为true则推送时忽略协议；为false则区分协议推送。
  // 建议和cache-ignore-protocol配套使用，避免推送失效。
  // 注意：
  // 1.一旦配置，则全局生效，不针对匹配的path-pattern生效。
  // 2.目录推送不区分协议，url推送可区分协议"}
  PurgeIgnoreProtocol *string `json:"purge-ignore-protocol,omitempty" xml:"purge-ignore-protocol,omitempty"`
  // {"en":"When configuring multiple groups of configurations, specify the id of a particular group of configurations", "zh_CN":"配置多组配置时，具体某组配置的id"}
  DataId *int64 `json:"data-id,omitempty" xml:"data-id,omitempty"`
}

func (s EditIgnoreProtocolRequestIgnoreProtocolRules) String() string {
  return tea.Prettify(s)
}

func (s EditIgnoreProtocolRequestIgnoreProtocolRules) GoString() string {
  return s.String()
}

func (s *EditIgnoreProtocolRequestIgnoreProtocolRules) SetPathPattern(v string) *EditIgnoreProtocolRequestIgnoreProtocolRules {
  s.PathPattern = &v
  return s
}

func (s *EditIgnoreProtocolRequestIgnoreProtocolRules) SetExceptPathPattern(v string) *EditIgnoreProtocolRequestIgnoreProtocolRules {
  s.ExceptPathPattern = &v
  return s
}

func (s *EditIgnoreProtocolRequestIgnoreProtocolRules) SetCacheIgnoreProtocol(v string) *EditIgnoreProtocolRequestIgnoreProtocolRules {
  s.CacheIgnoreProtocol = &v
  return s
}

func (s *EditIgnoreProtocolRequestIgnoreProtocolRules) SetPurgeIgnoreProtocol(v string) *EditIgnoreProtocolRequestIgnoreProtocolRules {
  s.PurgeIgnoreProtocol = &v
  return s
}

func (s *EditIgnoreProtocolRequestIgnoreProtocolRules) SetDataId(v int64) *EditIgnoreProtocolRequestIgnoreProtocolRules {
  s.DataId = &v
  return s
}

type EditIgnoreProtocolResponse struct {
  // {"en":"httpstatus=202; Indicates that the new domain API was successfully invoked, and the current deployment of the new domain can be viewed using x-cnc-request-id in the header", "zh_CN":"httpstatus=202;   表示成功调用新增域名接口，可使用header中的x-cnc-request-id查看当前新增域名的部署情况"}
  HttpStatus *int `json:"http status code,omitempty" xml:"http status code,omitempty" require:"true"`
  // {"en":"Uniquely identified id for querying tasks per request (for all API)", "zh_CN":"唯一标示的id，用于查询每次请求的任务 （适用全部接口）"}
  XCncRequestId *string `json:"x-cnc-request-id,omitempty" xml:"x-cnc-request-id,omitempty" require:"true"`
  // {"en":"The URL used to access the domain information, where domain-id is the unique token generated by our cloud platform for the domain name and whose value is a string.", "zh_CN":"响应信用于访问该域名信息的URL，其中domain-id为
  //      我司云平台为该域名生成的唯一标示，其值为字符串。"}
  Location *string `json:"location,omitempty" xml:"location,omitempty" require:"true"`
  // {"en":"Error code, which appears when HTTPStatus is not 202, represents the error type of the current request call", "zh_CN":"错误代码，当HTTPStatus不为202时出现，表示当前请求调用的错误类型"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response information, success when successful", "zh_CN":"响应信息，成功时为success"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s EditIgnoreProtocolResponse) String() string {
  return tea.Prettify(s)
}

func (s EditIgnoreProtocolResponse) GoString() string {
  return s.String()
}

func (s *EditIgnoreProtocolResponse) SetHttpStatus(v int) *EditIgnoreProtocolResponse {
  s.HttpStatus = &v
  return s
}

func (s *EditIgnoreProtocolResponse) SetXCncRequestId(v string) *EditIgnoreProtocolResponse {
  s.XCncRequestId = &v
  return s
}

func (s *EditIgnoreProtocolResponse) SetLocation(v string) *EditIgnoreProtocolResponse {
  s.Location = &v
  return s
}

func (s *EditIgnoreProtocolResponse) SetCode(v string) *EditIgnoreProtocolResponse {
  s.Code = &v
  return s
}

func (s *EditIgnoreProtocolResponse) SetMessage(v string) *EditIgnoreProtocolResponse {
  s.Message = &v
  return s
}

type EditIgnoreProtocolPaths struct {
  // {"en":"the domain whoes need query config", "zh_CN":"需要查询配置的域名或域名id"}
  DomainName *string `json:"domain-name,omitempty" xml:"domain-name,omitempty" require:"true"`
}

func (s EditIgnoreProtocolPaths) String() string {
  return tea.Prettify(s)
}

func (s EditIgnoreProtocolPaths) GoString() string {
  return s.String()
}

func (s *EditIgnoreProtocolPaths) SetDomainName(v string) *EditIgnoreProtocolPaths {
  s.DomainName = &v
  return s
}

type EditIgnoreProtocolParameters struct {
}

func (s EditIgnoreProtocolParameters) String() string {
  return tea.Prettify(s)
}

func (s EditIgnoreProtocolParameters) GoString() string {
  return s.String()
}

type EditIgnoreProtocolRequestHeader struct {
}

func (s EditIgnoreProtocolRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s EditIgnoreProtocolRequestHeader) GoString() string {
  return s.String()
}

type EditIgnoreProtocolResponseHeader struct {
}

func (s EditIgnoreProtocolResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s EditIgnoreProtocolResponseHeader) GoString() string {
  return s.String()
}




type EditHttpCodeCacheRequest struct {
  // {"en":"Status Code Caching Rule Configuration, parent node
  // 1. When you need to set status code caching rules, this must be filled in.
  // 2. Configuration of Clear Status Code Caching Rules for <http-code-cache-rules/>.", "zh_CN":"状态码缓存规则配置，父标签
  // 1.需要设置状态码缓存规则时，此项必填
  // 2.为<http-code-cache-rules/>时清空状态码缓存规则配置"}
  HttpCodeCacheRules []*EditHttpCodeCacheRequestHttpCodeCacheRules `json:"http-code-cache-rules,omitempty" xml:"http-code-cache-rules,omitempty" require:"true" type:"Repeated"`
}

func (s EditHttpCodeCacheRequest) String() string {
  return tea.Prettify(s)
}

func (s EditHttpCodeCacheRequest) GoString() string {
  return s.String()
}

func (s *EditHttpCodeCacheRequest) SetHttpCodeCacheRules(v []*EditHttpCodeCacheRequestHttpCodeCacheRules) *EditHttpCodeCacheRequest {
  s.HttpCodeCacheRules = v
  return s
}

type EditHttpCodeCacheRequestHttpCodeCacheRules struct     {
  // {"en":"Configure HTTP status code, parent node", "zh_CN":"配置http状态码，父标签"}
  HttpCodes []*string `json:"http-codes,omitempty" xml:"http-codes,omitempty" require:"true" type:"Repeated"`
  // {"en":"Define the caching time of the specified status code in units s, 0 to indicate no caching", "zh_CN":"配置指定的状态码的缓存时间，单位s，0表示不缓存"}
  CacheTtl *string `json:"cache-ttl,omitempty" xml:"cache-ttl,omitempty" require:"true"`
  // {"en":"Data-id is to indicate a specific group configuration when the client has multiple groups of configurations. Data-id can be retrieved through a query interface. Note:
  // A. If data-id is passed, it means that one group of configuration items is specified to be modified, and no other group configuration items need to be modified.
  // B. If multiple groups of configurations are included, some of them are configured with data-id and others are not, then the expression of data-id is used to modify a specific group of configurations, and a new group of configurations is added on the original basis without the expression of data-id.
  // C. If the data-id is not transmitted, it means that the original configuration will be fully covered by this configuration.
  // D. If no configuration parameter is passed, only domain name and secondary label are passed, which means that all configuration of domain name secondary service corresponding to this interface is cleared.
  // E. If there is no specific configuration item in a set of configurations, the data-id must be filled in, and the value is the actual data-id, which means clearing the value of the corresponding data-id configuration item; it is not allowed that there is no specific configuration item or data-id in a set of configurations.", "zh_CN":"配置多组配置时，具体某组配置的id。data-id可以通过查询接口获取。 注意：
  // a、如果有传data-id，说明指定修改其中一组配置项内容，不需求修改其他组配置内容不需要入参； 
  // b、如果入参多组配置，其中有些组配置有传data-id，有些没有传，则有传data-id的表示修改具体某组配置，没有传data-id的表示在原来基础上新增一组配置； 
  // c、如果入参都没有传data-id,表示用本次的配置全量覆盖原先配置； 
  // d、如果入参没有传任何配置项参数，只传了域名和二级标签，表示清空这个接口对应域名二级服务所有配置； 
  // e、如果一组配置没有具体的配置项，则data-id必填，且值为实际存在的data-id，表示清空这个data-id对应配置项的值；不允许一组配置没有具体的配置项也没有data-id。"}
  DataId *int64 `json:"data-id,omitempty" xml:"data-id,omitempty"`
}

func (s EditHttpCodeCacheRequestHttpCodeCacheRules) String() string {
  return tea.Prettify(s)
}

func (s EditHttpCodeCacheRequestHttpCodeCacheRules) GoString() string {
  return s.String()
}

func (s *EditHttpCodeCacheRequestHttpCodeCacheRules) SetHttpCodes(v []*string) *EditHttpCodeCacheRequestHttpCodeCacheRules {
  s.HttpCodes = v
  return s
}

func (s *EditHttpCodeCacheRequestHttpCodeCacheRules) SetCacheTtl(v string) *EditHttpCodeCacheRequestHttpCodeCacheRules {
  s.CacheTtl = &v
  return s
}

func (s *EditHttpCodeCacheRequestHttpCodeCacheRules) SetDataId(v int64) *EditHttpCodeCacheRequestHttpCodeCacheRules {
  s.DataId = &v
  return s
}

type EditHttpCodeCacheResponse struct {
  // {"en":"The error code, when HTTPStatus is not 202, indicates the type of error the current request is calling.", "zh_CN":"错误代码，当HTTPStatus不为202时出现，表示当前请求调用的错误类型"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response information, when success is successful", "zh_CN":"响应信息，成功时为success"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s EditHttpCodeCacheResponse) String() string {
  return tea.Prettify(s)
}

func (s EditHttpCodeCacheResponse) GoString() string {
  return s.String()
}

func (s *EditHttpCodeCacheResponse) SetCode(v string) *EditHttpCodeCacheResponse {
  s.Code = &v
  return s
}

func (s *EditHttpCodeCacheResponse) SetMessage(v string) *EditHttpCodeCacheResponse {
  s.Message = &v
  return s
}

type EditHttpCodeCachePaths struct {
  // {"en":"the domain whoes need query config", "zh_CN":"需要查询配置的域名或域名id"}
  DomainName *string `json:"domain-name,omitempty" xml:"domain-name,omitempty" require:"true"`
}

func (s EditHttpCodeCachePaths) String() string {
  return tea.Prettify(s)
}

func (s EditHttpCodeCachePaths) GoString() string {
  return s.String()
}

func (s *EditHttpCodeCachePaths) SetDomainName(v string) *EditHttpCodeCachePaths {
  s.DomainName = &v
  return s
}

type EditHttpCodeCacheParameters struct {
}

func (s EditHttpCodeCacheParameters) String() string {
  return tea.Prettify(s)
}

func (s EditHttpCodeCacheParameters) GoString() string {
  return s.String()
}

type EditHttpCodeCacheRequestHeader struct {
}

func (s EditHttpCodeCacheRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s EditHttpCodeCacheRequestHeader) GoString() string {
  return s.String()
}

type EditHttpCodeCacheResponseHeader struct {
}

func (s EditHttpCodeCacheResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s EditHttpCodeCacheResponseHeader) GoString() string {
  return s.String()
}




type EditWebsocketConfigRequest struct {
  // {"en":"Open or close websocket function, parent node, you can set <websocketSwitch/> to clear this configuration.
  // Scope of application: wsa, web pages", "zh_CN":"开启或关闭websocket功能，父标签，为<websocketSwitch/>则清空websocket开关配置
  // 适用范围：wsa、网页"}
  WebsocketSwitch *EditWebsocketConfigRequestWebsocketSwitch `json:"websocketSwitch,omitempty" xml:"websocketSwitch,omitempty" require:"true" type:"Struct"`
}

func (s EditWebsocketConfigRequest) String() string {
  return tea.Prettify(s)
}

func (s EditWebsocketConfigRequest) GoString() string {
  return s.String()
}

func (s *EditWebsocketConfigRequest) SetWebsocketSwitch(v *EditWebsocketConfigRequestWebsocketSwitch) *EditWebsocketConfigRequest {
  s.WebsocketSwitch = v
  return s
}

type EditWebsocketConfigRequestWebsocketSwitch struct {
  // {"en":"Whether to turn on the websocket function, the allowable values are true and false, default false", "zh_CN":"是否开启websocket功能,允许值为true和false，默认为否"}
  EnableWebsocket *bool `json:"enableWebsocket,omitempty" xml:"enableWebsocket,omitempty"`
}

func (s EditWebsocketConfigRequestWebsocketSwitch) String() string {
  return tea.Prettify(s)
}

func (s EditWebsocketConfigRequestWebsocketSwitch) GoString() string {
  return s.String()
}

func (s *EditWebsocketConfigRequestWebsocketSwitch) SetEnableWebsocket(v bool) *EditWebsocketConfigRequestWebsocketSwitch {
  s.EnableWebsocket = &v
  return s
}

type EditWebsocketConfigResponse struct {
  // {"en":"Uniquely labeled id for querying each requested task (for all interfaces)", "zh_CN":"唯一标示的id，用于查询每次请求的任务 （适用全部接口）"}
  XCncRequestId *string `json:"x-cnc-request-id,omitempty" xml:"x-cnc-request-id,omitempty" require:"true"`
  // {"en":"The URL used to access the domain information, where domain-id is the unique token generated by our cloud platform for the domain name and whose value is a string.", "zh_CN":"用于访问该域名信息的URL，其中domain-id为我司云平台为该域名生成的唯一标示，其值为字符串。"}
  Location *string `json:"Location,omitempty" xml:"Location,omitempty" require:"true"`
  // {"en":"Error code, 0 is success.", "zh_CN":"错误代码，成功为0"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response information, success when successful", "zh_CN":"响应信息，成功时为success"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s EditWebsocketConfigResponse) String() string {
  return tea.Prettify(s)
}

func (s EditWebsocketConfigResponse) GoString() string {
  return s.String()
}

func (s *EditWebsocketConfigResponse) SetXCncRequestId(v string) *EditWebsocketConfigResponse {
  s.XCncRequestId = &v
  return s
}

func (s *EditWebsocketConfigResponse) SetLocation(v string) *EditWebsocketConfigResponse {
  s.Location = &v
  return s
}

func (s *EditWebsocketConfigResponse) SetCode(v string) *EditWebsocketConfigResponse {
  s.Code = &v
  return s
}

func (s *EditWebsocketConfigResponse) SetMessage(v string) *EditWebsocketConfigResponse {
  s.Message = &v
  return s
}

type EditWebsocketConfigPaths struct {
  // {"en":"The domain whoes need query config.", "zh_CN":"需要查询配置的域名或域名id"}
  DomainName *string `json:"domain-name,omitempty" xml:"domain-name,omitempty"`
}

func (s EditWebsocketConfigPaths) String() string {
  return tea.Prettify(s)
}

func (s EditWebsocketConfigPaths) GoString() string {
  return s.String()
}

func (s *EditWebsocketConfigPaths) SetDomainName(v string) *EditWebsocketConfigPaths {
  s.DomainName = &v
  return s
}

type EditWebsocketConfigParameters struct {
}

func (s EditWebsocketConfigParameters) String() string {
  return tea.Prettify(s)
}

func (s EditWebsocketConfigParameters) GoString() string {
  return s.String()
}

type EditWebsocketConfigRequestHeader struct {
}

func (s EditWebsocketConfigRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s EditWebsocketConfigRequestHeader) GoString() string {
  return s.String()
}

type EditWebsocketConfigResponseHeader struct {
}

func (s EditWebsocketConfigResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s EditWebsocketConfigResponseHeader) GoString() string {
  return s.String()
}




type EditBack2originProtocolRewriteConfigRequest struct {
  // {"en":"Back to origin rewrite rule.", "zh_CN":"修改回源协议和端口；若要按原始请求回源，则可清空该对象，示例\"backToOriginRewriteRule\":{}"}
  BackToOriginRewriteRule *EditBack2originProtocolRewriteConfigRequestBackToOriginRewriteRule `json:"backToOriginRewriteRule,omitempty" xml:"backToOriginRewriteRule,omitempty" require:"true" type:"Struct"`
}

func (s EditBack2originProtocolRewriteConfigRequest) String() string {
  return tea.Prettify(s)
}

func (s EditBack2originProtocolRewriteConfigRequest) GoString() string {
  return s.String()
}

func (s *EditBack2originProtocolRewriteConfigRequest) SetBackToOriginRewriteRule(v *EditBack2originProtocolRewriteConfigRequestBackToOriginRewriteRule) *EditBack2originProtocolRewriteConfigRequest {
  s.BackToOriginRewriteRule = v
  return s
}

type EditBack2originProtocolRewriteConfigRequestBackToOriginRewriteRule struct {
  // {"en":"The specified protocol is either http or https.", "zh_CN":"改写后的回源协议，可选值：http、https"}
  Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
  // {"en":"If the protocol is http, the default is 80. If the protocol is https, the default is 443", "zh_CN":"改写后的回源端口，若protocol为http时，默认为80，若protocol为https时，默认为443"}
  Port *string `json:"port,omitempty" xml:"port,omitempty"`
}

func (s EditBack2originProtocolRewriteConfigRequestBackToOriginRewriteRule) String() string {
  return tea.Prettify(s)
}

func (s EditBack2originProtocolRewriteConfigRequestBackToOriginRewriteRule) GoString() string {
  return s.String()
}

func (s *EditBack2originProtocolRewriteConfigRequestBackToOriginRewriteRule) SetProtocol(v string) *EditBack2originProtocolRewriteConfigRequestBackToOriginRewriteRule {
  s.Protocol = &v
  return s
}

func (s *EditBack2originProtocolRewriteConfigRequestBackToOriginRewriteRule) SetPort(v string) *EditBack2originProtocolRewriteConfigRequestBackToOriginRewriteRule {
  s.Port = &v
  return s
}

type EditBack2originProtocolRewriteConfigResponse struct {
  // {"en":"Error code. Appeared when http status is not 200 or 202.", "zh_CN":"错误代码，当HTTPStatus不为202时出现，表示当前请求调用的错误类型"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Reponse message.", "zh_CN":"响应信息，成功时为success"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Response data", "zh_CN":"响应数据"}
  Data *string `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s EditBack2originProtocolRewriteConfigResponse) String() string {
  return tea.Prettify(s)
}

func (s EditBack2originProtocolRewriteConfigResponse) GoString() string {
  return s.String()
}

func (s *EditBack2originProtocolRewriteConfigResponse) SetCode(v string) *EditBack2originProtocolRewriteConfigResponse {
  s.Code = &v
  return s
}

func (s *EditBack2originProtocolRewriteConfigResponse) SetMessage(v string) *EditBack2originProtocolRewriteConfigResponse {
  s.Message = &v
  return s
}

func (s *EditBack2originProtocolRewriteConfigResponse) SetData(v string) *EditBack2originProtocolRewriteConfigResponse {
  s.Data = &v
  return s
}

type EditBack2originProtocolRewriteConfigPaths struct {
  // {"en":"the domain whoes need query config", "zh_CN":"需要查询配置的域名或域名id"}
  DomainName *string `json:"domain-name,omitempty" xml:"domain-name,omitempty" require:"true"`
}

func (s EditBack2originProtocolRewriteConfigPaths) String() string {
  return tea.Prettify(s)
}

func (s EditBack2originProtocolRewriteConfigPaths) GoString() string {
  return s.String()
}

func (s *EditBack2originProtocolRewriteConfigPaths) SetDomainName(v string) *EditBack2originProtocolRewriteConfigPaths {
  s.DomainName = &v
  return s
}

type EditBack2originProtocolRewriteConfigParameters struct {
}

func (s EditBack2originProtocolRewriteConfigParameters) String() string {
  return tea.Prettify(s)
}

func (s EditBack2originProtocolRewriteConfigParameters) GoString() string {
  return s.String()
}

type EditBack2originProtocolRewriteConfigRequestHeader struct {
}

func (s EditBack2originProtocolRewriteConfigRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s EditBack2originProtocolRewriteConfigRequestHeader) GoString() string {
  return s.String()
}

type EditBack2originProtocolRewriteConfigResponseHeader struct {
}

func (s EditBack2originProtocolRewriteConfigResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s EditBack2originProtocolRewriteConfigResponseHeader) GoString() string {
  return s.String()
}




type UpdateDomainMultiCertConfigForWplusRequest struct {
  // {"en":"Signature certificate info", "zh_CN":"签名证书信息"}
  UpdateDomainMultiCertConfigForWplusDomainCertConfig *UpdateDomainMultiCertConfigForWplusDomainCertConfig `json:"domainCertConfig,omitempty" xml:"domainCertConfig,omitempty" require:"true"`
  // {"en":"Encryption certificate info", "zh_CN":"加密证书信息"}
  UpdateDomainMultiCertConfigForWplusOtherConfig *UpdateDomainMultiCertConfigForWplusOtherConfig `json:"otherConfig,omitempty" xml:"otherConfig,omitempty" require:"true"`
}

func (s UpdateDomainMultiCertConfigForWplusRequest) String() string {
  return tea.Prettify(s)
}

func (s UpdateDomainMultiCertConfigForWplusRequest) GoString() string {
  return s.String()
}

func (s *UpdateDomainMultiCertConfigForWplusRequest) SetDomainCertConfig(v *UpdateDomainMultiCertConfigForWplusDomainCertConfig) *UpdateDomainMultiCertConfigForWplusRequest {
  s.UpdateDomainMultiCertConfigForWplusDomainCertConfig = v
  return s
}

func (s *UpdateDomainMultiCertConfigForWplusRequest) SetOtherConfig(v *UpdateDomainMultiCertConfigForWplusOtherConfig) *UpdateDomainMultiCertConfigForWplusRequest {
  s.UpdateDomainMultiCertConfigForWplusOtherConfig = v
  return s
}

type UpdateDomainMultiCertConfigForWplusDomainCertConfig struct {
  // {"en":"operateType", "zh_CN":"CANCEL表示删除，ADD表示关联"}
  OperateType *string `json:"operateType,omitempty" xml:"operateType,omitempty" require:"true"`
  // {"en":"cert config", "zh_CN":"域名证书配置"}
  UpdateDomainMultiCertConfigForWplusCertConfigs []*UpdateDomainMultiCertConfigForWplusCertConfigs `json:"certConfigs,omitempty" xml:"certConfigs,omitempty" require:"true" type:"Repeated"`
}

func (s UpdateDomainMultiCertConfigForWplusDomainCertConfig) String() string {
  return tea.Prettify(s)
}

func (s UpdateDomainMultiCertConfigForWplusDomainCertConfig) GoString() string {
  return s.String()
}

func (s *UpdateDomainMultiCertConfigForWplusDomainCertConfig) SetOperateType(v string) *UpdateDomainMultiCertConfigForWplusDomainCertConfig {
  s.OperateType = &v
  return s
}

func (s *UpdateDomainMultiCertConfigForWplusDomainCertConfig) SetCertConfigs(v []*UpdateDomainMultiCertConfigForWplusCertConfigs) *UpdateDomainMultiCertConfigForWplusDomainCertConfig {
  s.UpdateDomainMultiCertConfigForWplusCertConfigs = v
  return s
}

type UpdateDomainMultiCertConfigForWplusCertConfigs struct {
  // {"en":"certUsage", "zh_CN":"证书用途"}
  CertUsage *string `json:"certUsage,omitempty" xml:"certUsage,omitempty" require:"true"`
  // {"en":"certificate id list", "zh_CN":"证书id列表"}
  CertificateIds []*string `json:"certificateIds,omitempty" xml:"certificateIds,omitempty" require:"true" type:"Repeated"`
}

func (s UpdateDomainMultiCertConfigForWplusCertConfigs) String() string {
  return tea.Prettify(s)
}

func (s UpdateDomainMultiCertConfigForWplusCertConfigs) GoString() string {
  return s.String()
}

func (s *UpdateDomainMultiCertConfigForWplusCertConfigs) SetCertUsage(v string) *UpdateDomainMultiCertConfigForWplusCertConfigs {
  s.CertUsage = &v
  return s
}

func (s *UpdateDomainMultiCertConfigForWplusCertConfigs) SetCertificateIds(v []*string) *UpdateDomainMultiCertConfigForWplusCertConfigs {
  s.CertificateIds = v
  return s
}

type UpdateDomainMultiCertConfigForWplusOtherConfig struct {
  // {"en":"tlsVersion", "zh_CN":"TLS协议"}
  TlsVersion *string `json:"tlsVersion,omitempty" xml:"tlsVersion,omitempty" require:"true"`
  // {"en":"enableOCSP", "zh_CN":"支持OCSP"}
  EnableOCSP *string `json:"enableOCSP,omitempty" xml:"enableOCSP,omitempty" require:"true"`
  // {"en":"cipherSuites", "zh_CN":"cipher套件"}
  CipherSuites *string `json:"cipherSuites,omitempty" xml:"cipherSuites,omitempty" require:"true"`
}

func (s UpdateDomainMultiCertConfigForWplusOtherConfig) String() string {
  return tea.Prettify(s)
}

func (s UpdateDomainMultiCertConfigForWplusOtherConfig) GoString() string {
  return s.String()
}

func (s *UpdateDomainMultiCertConfigForWplusOtherConfig) SetTlsVersion(v string) *UpdateDomainMultiCertConfigForWplusOtherConfig {
  s.TlsVersion = &v
  return s
}

func (s *UpdateDomainMultiCertConfigForWplusOtherConfig) SetEnableOCSP(v string) *UpdateDomainMultiCertConfigForWplusOtherConfig {
  s.EnableOCSP = &v
  return s
}

func (s *UpdateDomainMultiCertConfigForWplusOtherConfig) SetCipherSuites(v string) *UpdateDomainMultiCertConfigForWplusOtherConfig {
  s.CipherSuites = &v
  return s
}

type UpdateDomainMultiCertConfigForWplusResponse struct {
  // {"en":"code", "zh_CN":"错误码，成功为0"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"error message", "zh_CN":"错误信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s UpdateDomainMultiCertConfigForWplusResponse) String() string {
  return tea.Prettify(s)
}

func (s UpdateDomainMultiCertConfigForWplusResponse) GoString() string {
  return s.String()
}

func (s *UpdateDomainMultiCertConfigForWplusResponse) SetCode(v string) *UpdateDomainMultiCertConfigForWplusResponse {
  s.Code = &v
  return s
}

func (s *UpdateDomainMultiCertConfigForWplusResponse) SetMessage(v string) *UpdateDomainMultiCertConfigForWplusResponse {
  s.Message = &v
  return s
}

type UpdateDomainMultiCertConfigForWplusPaths struct {
}

func (s UpdateDomainMultiCertConfigForWplusPaths) String() string {
  return tea.Prettify(s)
}

func (s UpdateDomainMultiCertConfigForWplusPaths) GoString() string {
  return s.String()
}

type UpdateDomainMultiCertConfigForWplusParameters struct {
}

func (s UpdateDomainMultiCertConfigForWplusParameters) String() string {
  return tea.Prettify(s)
}

func (s UpdateDomainMultiCertConfigForWplusParameters) GoString() string {
  return s.String()
}

type UpdateDomainMultiCertConfigForWplusRequestHeader struct {
}

func (s UpdateDomainMultiCertConfigForWplusRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s UpdateDomainMultiCertConfigForWplusRequestHeader) GoString() string {
  return s.String()
}

type UpdateDomainMultiCertConfigForWplusResponseHeader struct {
}

func (s UpdateDomainMultiCertConfigForWplusResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s UpdateDomainMultiCertConfigForWplusResponseHeader) GoString() string {
  return s.String()
}




type DeleteBanURLsRequest struct {
  // {"en":"The domain name", "zh_CN":"域名名称"}
  DomainName *string `json:"domainName,omitempty" xml:"domainName,omitempty" require:"true"`
  // {"en":"List of urls to mask.Can contain multiple <url> nodes.<banUrls> cannot be specified as blank when it exists.When the value of the <deleteAll> node is true, the <banUrls> node and its children cannot exist.When the <deleteAll> node value is false or does not exist, the <banUrls> node and its children must exist and have a value.", "zh_CN":"要屏蔽的url列表。可包含多个<url>节点。当<banUrls>存在时，不能被指定为空白。当<deleteAll>节点值为true时，<banUrls>节点及其子节点都不能存在。当<deleteAll>节点值为false或不存在时，<banUrls>节点及其子节点必须存在，且有值。"}
  BanUrls []*string `json:"banUrls,omitempty" xml:"banUrls,omitempty" type:"Repeated"`
  // {"en":"Whether to remove all ban urls under the domain name.Optional value [true: delete all, false: do not delete all] (ignore case).Cannot be specified as blank when existing.When the value of the <deleteAll> node is true, the <banUrls> node and its children cannot exist.When the <deleteAll> node value is false or does not exist, the <banUrls> node and its children must exist and have a value.", "zh_CN":"是否删除域名下所有ban url。可选值为【true：删除所有，false：不删除所有】（忽略大小写）。当存在时，不能被指定为空白。当<deleteAll>节点值为true时，<banUrls>节点及其子节点都不能存在。当<deleteAll>节点值为false或不存在时，<banUrls>节点及其子节点必须存在，且有值。"}
  DeleteAll *bool `json:"deleteAll,omitempty" xml:"deleteAll,omitempty" require:"true"`
}

func (s DeleteBanURLsRequest) String() string {
  return tea.Prettify(s)
}

func (s DeleteBanURLsRequest) GoString() string {
  return s.String()
}

func (s *DeleteBanURLsRequest) SetDomainName(v string) *DeleteBanURLsRequest {
  s.DomainName = &v
  return s
}

func (s *DeleteBanURLsRequest) SetBanUrls(v []*string) *DeleteBanURLsRequest {
  s.BanUrls = v
  return s
}

func (s *DeleteBanURLsRequest) SetDeleteAll(v bool) *DeleteBanURLsRequest {
  s.DeleteAll = &v
  return s
}

type DeleteBanURLsResponse struct {
  // {"en":"Error code", "zh_CN":"错误码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"The message body", "zh_CN":"消息体"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Returns the body of the data. The <date> returned by this interface is always empty", "zh_CN":"返回数据体，此接口返回的<date>恒为空"}
  Data *string `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s DeleteBanURLsResponse) String() string {
  return tea.Prettify(s)
}

func (s DeleteBanURLsResponse) GoString() string {
  return s.String()
}

func (s *DeleteBanURLsResponse) SetCode(v string) *DeleteBanURLsResponse {
  s.Code = &v
  return s
}

func (s *DeleteBanURLsResponse) SetMessage(v string) *DeleteBanURLsResponse {
  s.Message = &v
  return s
}

func (s *DeleteBanURLsResponse) SetData(v string) *DeleteBanURLsResponse {
  s.Data = &v
  return s
}

type DeleteBanURLsPaths struct {
}

func (s DeleteBanURLsPaths) String() string {
  return tea.Prettify(s)
}

func (s DeleteBanURLsPaths) GoString() string {
  return s.String()
}

type DeleteBanURLsParameters struct {
}

func (s DeleteBanURLsParameters) String() string {
  return tea.Prettify(s)
}

func (s DeleteBanURLsParameters) GoString() string {
  return s.String()
}

type DeleteBanURLsRequestHeader struct {
}

func (s DeleteBanURLsRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteBanURLsRequestHeader) GoString() string {
  return s.String()
}

type DeleteBanURLsResponseHeader struct {
}

func (s DeleteBanURLsResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteBanURLsResponseHeader) GoString() string {
  return s.String()
}




type EditQueryStringUrlConfigRequest struct {
  // {"en":"Query String Settings Configuration, parent node
  // 1. When you need to configure the query string, this must be filled in.
  // 2. Configuration of clearing query string settings for <query-string-settings/>.", "zh_CN":"查询串设置配置，父标签
  // 1.需要设置查询串配置时，此项必填
  // 2.为<query-string-settings/>时清空查询串设置的配置"}
  QueryStringSettings []*EditQueryStringUrlConfigRequestQueryStringSettings `json:"query-string-settings,omitempty" xml:"query-string-settings,omitempty" require:"true" type:"Repeated"`
}

func (s EditQueryStringUrlConfigRequest) String() string {
  return tea.Prettify(s)
}

func (s EditQueryStringUrlConfigRequest) GoString() string {
  return s.String()
}

func (s *EditQueryStringUrlConfigRequest) SetQueryStringSettings(v []*EditQueryStringUrlConfigRequestQueryStringSettings) *EditQueryStringUrlConfigRequest {
  s.QueryStringSettings = v
  return s
}

type EditQueryStringUrlConfigRequestQueryStringSettings struct     {
  // {"en":"The url matching mode. If all matches, the input parameters can be configured as: .*", "zh_CN":"url匹配模式，支持正则，如果是全部匹配，入参可以配置为：.*
  // 注：url匹配模式、文件类型（自定义文件类型）、常用类型、指定url、目录，有且仅有一项必填"}
  PathPattern *string `json:"path-pattern,omitempty" xml:"path-pattern,omitempty"`
  // {"en":"File Type: Specify the file type for anti-theft chain settings.
  // File types include: gif png bmp jpeg jpg html htm shtml mp3 wma flv mp4 wmv zip exe rar css txt ico js swf
  // If you need all types, pass all directly. Multiples are separated by semicolons, and all and specific file types cannot be configured at the same time.", "zh_CN":"文件类型：指定文件类型进行防盗链设置。可选文件类型包括：gif png bmp jpeg jpg html htm shtml mp3 wma flv mp4 wmv zip exe rar css txt ico js swf m3u8 xml f4m bootstarp ts 如果需要全部类型，则直接传all。多个以分号隔开，all和具体文件类型不能同时配置"}
  FileTypes *string `json:"file-types,omitempty" xml:"file-types,omitempty"`
  // {"en":"Custom file type: Fill in the appropriate identifiable file type according to your needs outside of the specified file type. Can be used with file-type. If the file-type is also configured, the actual file type is the sum of the two parameters.", "zh_CN":"自定义文件类型：在指定文件类型外根据自身需求，填写适当的可识别文件类型。可以搭配file-type使用。如果file-type也有配置，实际生效的文件类型是两个入参的总和"}
  CustomFileTypes *string `json:"custom-file-types,omitempty" xml:"custom-file-types,omitempty"`
  // {"en":"Specify common types: Select the domain name that requires the anti-theft chain to be all files or the home page. :
  // E.g:
  // All: all files
  // Homepage: homepage", "zh_CN":"常用类型：可选值为homepage和all
  // all：全部文件
  // homepage：首页"}
  CustomPattern *string `json:"custom-pattern,omitempty" xml:"custom-pattern,omitempty"`
  // {"en":"Specify URL cache: Specify url according to requirements for anti-theft chain setting
  // INS format does not support URI format with http(s)://", "zh_CN":"指定url，非http(s):// 开头，多个以换行分隔"}
  SpecifyUrlPattern *string `json:"specify-url-pattern,omitempty" xml:"specify-url-pattern,omitempty"`
  // {"en":"Directory: Specify the directory for anti-theft chain settings
  // Enter a legal directory format. Multiple separated by semicolons", "zh_CN":"目录，可输入合法的目录格式。以/开头和结尾，多个以分号隔开。"}
  Directories *string `json:"directories,omitempty" xml:"directories,omitempty"`
  // {"en":"Indicates the priority execution order of multiple sets of redirected content by the customer. The higher the number, the higher the priority.
  // When adding a new configuration item, the default is 10", "zh_CN":"优先级，表示客户多组配置的优先执行顺序。数字越大，优先级越高。不传默认为10，不可清空。"}
  Priority *string `json:"priority,omitempty" xml:"priority,omitempty"`
  // {"en":"Whether to ignore letter case.", "zh_CN":"是否忽略大小写：允许值为true和false，默认为忽略"}
  IgnoreLetterCase *string `json:"ignore-letter-case,omitempty" xml:"ignore-letter-case,omitempty"`
  // {"en":"Define whether to cache with query strings, 'true' means ignore query strings, while 'false' means cache with all query strings.", "zh_CN":"缓存是否忽略查询串，允许值为true和false。
  // true表示忽略查询串，相同拷贝；false表示不忽略，分别缓存，即带全部参数缓存。"}
  IgnoreQueryString *string `json:"ignore-query-string,omitempty" xml:"ignore-query-string,omitempty"`
  // {"en":"Cache with the specified query string parameters. If the kept parameter values are the same, one copy will be cached.
  // Note:
  // 1. query-string-kept and query-string-removed are mutually exclusive, and only one of them has a value.
  // 2. query-string-kept and ignore-query-string are mutually exclusive, and only one has a value.", "zh_CN":"缓存保留参数，指定保留的参数值相同，则缓存一份。
  // 注：
  // 1.query-string-kept和query-string-removed两者互斥，只能一个有值。
  // 2.query-string-kept和ignore-query-string两者互斥，只能一个有值。"}
  QueryStringKept *string `json:"query-string-kept,omitempty" xml:"query-string-kept,omitempty"`
  // {"en":"Cache without the specified query string parameters. After deleting the specified parameter, if the other parameter values are the same, one copy will be cached.
  // 1. query-string-kept and query string removed are mutually exclusive, and only one has a value.
  // 2. query-string-removed and ignore-query-string are mutually exclusive.", "zh_CN":"缓存删除参数，删除指定的参数后，其余参数值相同，则缓存一份。
  // 1.query-string-kept和query-string-removed两者互斥，只能一个有值。
  // 2.query-string-removed和ignore-query-string两者互斥，只能一个有值。"}
  QueryStringRemoved *string `json:"query-string-removed,omitempty" xml:"query-string-removed,omitempty"`
  // {"en":"Whether to use the original URL back source, the allowable values are true and false.
  // When ignore-query-string is true or not set, source-with-query is true to indicate that the source is returned according to the original request, and false to indicate that the question mark is returned.
  // When ignore-query-string is false, this default setting is empty (input is invalid)", "zh_CN":"是否用原始url回源，允许值为true和false。
  // ignore-query-string为true或未设置时，source-with-query为true表示按原始请求回源；为false表示去问号回源。
  // ignore-query-string为false时，此项默认设置为空（输入无效）。"}
  SourceWithQuery *string `json:"source-with-query,omitempty" xml:"source-with-query,omitempty"`
  // {"en":"Return to the source after specifying the reserved parameter value. Please separate them with semicolons, if no parameters reserved, please fill in:- . 1. Source-key-kept and ignore-query-string are mutually exclusive, and only one of them has a value. 2. Source-key-kept and source-key-removed are mutually exclusive, and only one of them has a value.
  // ", "zh_CN":"回源保留参数，指定保留的参数值后回源。多个请以英文分号分隔，任何参数都不保留请填：- 1、source-key-kept和ignore-query-string两者互斥，只能一个有值。 2、source-key-kept和source-key-removed两者互斥，只能一个有值。
  // "}
  SourceKeyKept *string `json:"source-key-kept,omitempty" xml:"source-key-kept,omitempty"`
  // {"en":"Return to the source after specifying the deleted parameter value. Please separate them with semicolons, and if you do not delete any parameters, please fill in:- . 1. Source-key-removed and ignore-query-string are mutually exclusive, and only one of them has a value. 2. Source-key-kept and source-key-removed are mutually exclusive, and only one of them has a value.
  // ", "zh_CN":"回源删除参数，指定删除的参数值后回源。多个请以英文分号分隔，任何参数都不删除请填：- 1、source-key-removed和ignore-query-string两者互斥，只能一个有值。 2、source-key-kept和source-key-removed两者互斥，只能一个有值。
  // "}
  SourceKeyRemoved *string `json:"source-key-removed,omitempty" xml:"source-key-removed,omitempty"`
  // {"en":"Data-id is to indicate a specific group configuration when the client has multiple groups of configurations. Data-id can be retrieved through a query interface. Note: A. If data-id is passed, it means that one group of configuration items is specified to be modified, and no other group configuration items need to be modified. B. If multiple groups of configurations are included, some of them are configured with data-id and others are not, then the expression of data-id is used to modify a specific group of configurations, and a new group of configurations is added on the original basis without the expression of data-id. C. If the data-id is not transmitted, it means that the original configuration will be fully covered by this configuration. D. If no configuration parameter is passed, only domain name and secondary label are passed, which means that all configuration of domain name secondary service corresponding to this interface is cleared. E. If there is no specific configuration item in a set of configurations, the data-id must be filled in, and the value is the actual data-id, which means clearing the value of the corresponding data-id configuration item; it is not allowed that there is no specific configuration item or data-id in a set of configurations.", "zh_CN":"配置多组配置时，具体某组配置的id。data-id可以通过查询接口获取。 注意： a、如果有传data-id，说明指定修改其中一组配置项内容，不需求修改其他组配置内容不需要入参；  b、如果入参多组配置，其中有些组配置有传data-id，有些没有传，则有传data-id的表示修改具体某组配置，没有传data-id的表示在原来基础上新增一组配置；  c、如果入参都没有传data-id,表示用本次的配置全量覆盖原先配置；  d、如果入参没有传任何配置项参数，只传了域名和二级标签，表示清空这个接口对应域名二级服务所有配置；  e、如果一组配置没有具体的配置项，则data-id必填，且值为实际存在的data-id，表示清空这个data-id对应配置项的值；不允许一组配置没有具体的配置项也没有data-id。"}
  DataId *int64 `json:"data-id,omitempty" xml:"data-id,omitempty"`
}

func (s EditQueryStringUrlConfigRequestQueryStringSettings) String() string {
  return tea.Prettify(s)
}

func (s EditQueryStringUrlConfigRequestQueryStringSettings) GoString() string {
  return s.String()
}

func (s *EditQueryStringUrlConfigRequestQueryStringSettings) SetPathPattern(v string) *EditQueryStringUrlConfigRequestQueryStringSettings {
  s.PathPattern = &v
  return s
}

func (s *EditQueryStringUrlConfigRequestQueryStringSettings) SetFileTypes(v string) *EditQueryStringUrlConfigRequestQueryStringSettings {
  s.FileTypes = &v
  return s
}

func (s *EditQueryStringUrlConfigRequestQueryStringSettings) SetCustomFileTypes(v string) *EditQueryStringUrlConfigRequestQueryStringSettings {
  s.CustomFileTypes = &v
  return s
}

func (s *EditQueryStringUrlConfigRequestQueryStringSettings) SetCustomPattern(v string) *EditQueryStringUrlConfigRequestQueryStringSettings {
  s.CustomPattern = &v
  return s
}

func (s *EditQueryStringUrlConfigRequestQueryStringSettings) SetSpecifyUrlPattern(v string) *EditQueryStringUrlConfigRequestQueryStringSettings {
  s.SpecifyUrlPattern = &v
  return s
}

func (s *EditQueryStringUrlConfigRequestQueryStringSettings) SetDirectories(v string) *EditQueryStringUrlConfigRequestQueryStringSettings {
  s.Directories = &v
  return s
}

func (s *EditQueryStringUrlConfigRequestQueryStringSettings) SetPriority(v string) *EditQueryStringUrlConfigRequestQueryStringSettings {
  s.Priority = &v
  return s
}

func (s *EditQueryStringUrlConfigRequestQueryStringSettings) SetIgnoreLetterCase(v string) *EditQueryStringUrlConfigRequestQueryStringSettings {
  s.IgnoreLetterCase = &v
  return s
}

func (s *EditQueryStringUrlConfigRequestQueryStringSettings) SetIgnoreQueryString(v string) *EditQueryStringUrlConfigRequestQueryStringSettings {
  s.IgnoreQueryString = &v
  return s
}

func (s *EditQueryStringUrlConfigRequestQueryStringSettings) SetQueryStringKept(v string) *EditQueryStringUrlConfigRequestQueryStringSettings {
  s.QueryStringKept = &v
  return s
}

func (s *EditQueryStringUrlConfigRequestQueryStringSettings) SetQueryStringRemoved(v string) *EditQueryStringUrlConfigRequestQueryStringSettings {
  s.QueryStringRemoved = &v
  return s
}

func (s *EditQueryStringUrlConfigRequestQueryStringSettings) SetSourceWithQuery(v string) *EditQueryStringUrlConfigRequestQueryStringSettings {
  s.SourceWithQuery = &v
  return s
}

func (s *EditQueryStringUrlConfigRequestQueryStringSettings) SetSourceKeyKept(v string) *EditQueryStringUrlConfigRequestQueryStringSettings {
  s.SourceKeyKept = &v
  return s
}

func (s *EditQueryStringUrlConfigRequestQueryStringSettings) SetSourceKeyRemoved(v string) *EditQueryStringUrlConfigRequestQueryStringSettings {
  s.SourceKeyRemoved = &v
  return s
}

func (s *EditQueryStringUrlConfigRequestQueryStringSettings) SetDataId(v int64) *EditQueryStringUrlConfigRequestQueryStringSettings {
  s.DataId = &v
  return s
}

type EditQueryStringUrlConfigResponse struct {
  // {"en":"The error code, when HTTPStatus is not 202, indicates the type of error the current request is calling.", "zh_CN":"错误代码，当HTTPStatus不为202时出现，表示当前请求调用的错误类型"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response information, when success is successful", "zh_CN":"响应信息，成功时为success"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s EditQueryStringUrlConfigResponse) String() string {
  return tea.Prettify(s)
}

func (s EditQueryStringUrlConfigResponse) GoString() string {
  return s.String()
}

func (s *EditQueryStringUrlConfigResponse) SetCode(v string) *EditQueryStringUrlConfigResponse {
  s.Code = &v
  return s
}

func (s *EditQueryStringUrlConfigResponse) SetMessage(v string) *EditQueryStringUrlConfigResponse {
  s.Message = &v
  return s
}

type EditQueryStringUrlConfigPaths struct {
  // {"en":"the domain whoes need query config", "zh_CN":"需要查询配置的域名或域名id"}
  DomainName *string `json:"domain-name,omitempty" xml:"domain-name,omitempty" require:"true"`
}

func (s EditQueryStringUrlConfigPaths) String() string {
  return tea.Prettify(s)
}

func (s EditQueryStringUrlConfigPaths) GoString() string {
  return s.String()
}

func (s *EditQueryStringUrlConfigPaths) SetDomainName(v string) *EditQueryStringUrlConfigPaths {
  s.DomainName = &v
  return s
}

type EditQueryStringUrlConfigParameters struct {
}

func (s EditQueryStringUrlConfigParameters) String() string {
  return tea.Prettify(s)
}

func (s EditQueryStringUrlConfigParameters) GoString() string {
  return s.String()
}

type EditQueryStringUrlConfigRequestHeader struct {
}

func (s EditQueryStringUrlConfigRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s EditQueryStringUrlConfigRequestHeader) GoString() string {
  return s.String()
}

type EditQueryStringUrlConfigResponseHeader struct {
}

func (s EditQueryStringUrlConfigResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s EditQueryStringUrlConfigResponseHeader) GoString() string {
  return s.String()
}




type GetBasicConfigurationOfDomainRequest struct {
}

func (s GetBasicConfigurationOfDomainRequest) String() string {
  return tea.Prettify(s)
}

func (s GetBasicConfigurationOfDomainRequest) GoString() string {
  return s.String()
}

type GetBasicConfigurationOfDomainResponse struct {
  // {"en":"Acceleration domain ID returned by the system.", "zh_CN":"系统返回的加速域名ID"}
  DomainId *int `json:"domain-id,omitempty" xml:"domain-id,omitempty" require:"true"`
  // {"en":"Accelerated domain name.", "zh_CN":"加速域名的名称"}
  DomainName *string `json:"domain-name,omitempty" xml:"domain-name,omitempty" require:"true"`
  // {"en":"Domain name creation time,
  // Format: week, dd month yyyy hh:mm:ss GMT +8:00", "zh_CN":"域名创建时间，格式: week, dd month yyyy hh:mm:ss GMT +8:00
  // 如：Mon, 18 Feb 2019 02:54:19 GMT +8:00"}
  CreatedDate *string `json:"created-date,omitempty" xml:"created-date,omitempty" require:"true"`
  // {"en":"Domain name last modified time,
  // Format: week, dd month yyyy hh:mm:ss GMT +8:00
  // Example: Mon, 18 Feb 2019 02:54:19 GMT +8 p.m", "zh_CN":"域名最近修改时间，格式: week, dd month yyyy hh:mm:ss GMT +8:00
  // 如：Mon, 18 Feb 2019 02:54:19 GMT +8:00"}
  LastModified *string `json:"last-modified,omitempty" xml:"last-modified,omitempty" require:"true"`
  // {"en":"Speed up domain name service types, including the following:
  // Web/web-https: web acceleration / web acceleration - https
  // Wsa/wsa-https: Total Station Acceleration / Total Station Acceleration - https
  // Vodstream/vod-https: on-demand acceleration/on-demand acceleration-https
  // Download/dl-https: Download acceleration/download acceleration - https", "zh_CN":"加速域名的服务类型，包括如下：
  // web/web-https：网页加速/网页加速-https
  // wsa/wsa-https：全站加速/全站加速-https
  // vodstream/vod-https：点播加速/点播加速-https
  // download/dl-https：下载加速/下载加速-https"}
  ServiceType *string `json:"service-type,omitempty" xml:"service-type,omitempty" require:"true"`
  // {"en":"Remarks, up to 1000 characters", "zh_CN":"备注信息，最大限制1000个字符"}
  Comment *string `json:"comment,omitempty" xml:"comment,omitempty" require:"true"`
  // {"en":"Service areas of the domain name", "zh_CN":"域名的加速区域."}
  ServiceAreas *string `json:"service-areas,omitempty" xml:"service-areas,omitempty" require:"true"`
  // {"en":"CNAME of the domain you queried, for example: 7nt6mrh7sdkslj.cdn30.com.", 
  //     "zh_CN":"加速域名对应的CNAME域名，例如：7nt6mrh7sdkslj.cdn30.com。"}
  Cname *string `json:"cname,omitempty" xml:"cname,omitempty" require:"true"`
  // {"en":"The deployment status of the domain name. Deployed indicates that the domain configuration is distributed successfully. InProgress indicates that the deployment task of the domain configuration is still in progress, and may be in queue, or failed.", 
  //     "zh_CN":"加速域名的部署状态，Deployed表示该加速域名配置完成部署；InProgress表示该加速域名配置的部署任务还在进行中，可能处于排队、部署中或失败任意一种状态。"}
  Status *string `json:"status,omitempty" xml:"status,omitempty" require:"true"`
  // {"en":"Accelerate status of the domain in our CDN, true means the CDN acceleration is normal; false means all request will back to origin directly in DNS.", "zh_CN":"加速域名的CDN服务状态，true表示启用CDN加速服务；false表示取消CDN加速服务。"}
  CdnServiceStatus *string `json:"cdn-service-status,omitempty" xml:"cdn-service-status,omitempty" require:"true"`
  // {"en":"Activation of the domain. It is false when the domain service is disabled, and true when the domain service is enabled.","zh_CN":"加速域名的启用状态，当禁用加速域名服务后，此项为false；当启用加速域名服务后，此项为true"}
  Enabled *string `json:"enabled,omitempty" xml:"enabled,omitempty" require:"true"`
  // {"en":"Shared first level alias", "zh_CN":"共用一级别名"}
  CnameLabel *string `json:"cname-label,omitempty" xml:"cname-label,omitempty" require:"true"`
  // {"en":"Return source policy setting, used to set the source station information and return source policy of the accelerated domain name", "zh_CN":"回源策略设置，用于设置加速域名的源站信息和回源策略"}
  OriginConfig *GetBasicConfigurationOfDomainResponseOriginConfig `json:"origin-config,omitempty" xml:"origin-config,omitempty" require:"true" type:"Struct"`
  // {"en":"Ssl certificate settings, used to set the ssl certificate configuration for the accelerated domain name", "zh_CN":"ssl证书设置，用于设置加速域名的ssl证书配置"}
  Ssl *GetBasicConfigurationOfDomainResponseSsl `json:"ssl,omitempty" xml:"ssl,omitempty" require:"true" type:"Struct"`
  // {"en":"Cache rule settings for setting cache rules for accelerated domain names", "zh_CN":"查询缓存时间配置，请使用新接口：【查询缓存时间配置】接口"}
  CacheBehaviors map[string]interface{} `json:"cache-behaviors,omitempty" xml:"cache-behaviors,omitempty" require:"true"`
  // {"en":"Cache file HOST (not return by default, application is required to use)", "zh_CN":"缓存文件HOST（默认不返回，使用需申请）
  // 缓存HOST域名和加速域名的”缓存规则”必须一致
  // 注意：该节点下的相关参数配置，除开通API调用权限外，还需要联系专属客服申请开通对应的API客户模板"}
  CacheHost *string `json:"cache-host,omitempty" xml:"cache-host,omitempty" require:"true"`
  // {"en":"Enable httpdns settings (not return by default, application is required to use)", "zh_CN":"启用httpdns设置（默认不返回，使用需申请）
  // 可选值为true和false，true表示启用；false表示关闭
  // 注意：该节点下的相关参数配置，除开通API调用权限外，还需要联系专属客服申请开通对应的API客户模板"}
  EnableHttpdns *string `json:"enable-httpdns,omitempty" xml:"enable-httpdns,omitempty" require:"true"`
  // {"en":"Pass the response header of client IP. The optional values are Cdn-Src-Ip and X-Forwarded-For. The default value is Cdn-Src-Ip", "zh_CN":"传递客户端ip的响应头部，可选值为Cdn-Src-Ip和X-Forwarded-For，默认值为Cdn-Src-Ip"}
  HeaderOfClientip *string `json:"header-of-clientip,omitempty" xml:"header-of-clientip,omitempty" require:"true"`
  // {"en":"The live push-pull stream type, the optional values are pull and push, pull means pull flow; push means push flow.", "zh_CN":"直播推拉流类型，可选值为pull和push，pull表示拉流；   push表示推流。"}
  DomainStreamType *string `json:"domain-stream-type,omitempty" xml:"domain-stream-type,omitempty" require:"true"`
  // {"en":"Live domain name configuration, rtmp live acceleration domain name push-pull flow", "zh_CN":"直播域名配置，rtmp直播加速域名的推拉流"}
  LiveConfig *GetBasicConfigurationOfDomainResponseLiveConfig `json:"live-config,omitempty" xml:"live-config,omitempty" require:"true" type:"Struct"`
  // {"en":"The launch point of the live push-pull domain name", "zh_CN":"直播推拉流域名的发布点
  // 注意：拉流和对应的推流域名，发布点是相同的"}
  PublishPoints *GetBasicConfigurationOfDomainResponsePublishPoints `json:"publish-points,omitempty" xml:"publish-points,omitempty" require:"true" type:"Struct"`
}

func (s GetBasicConfigurationOfDomainResponse) String() string {
  return tea.Prettify(s)
}

func (s GetBasicConfigurationOfDomainResponse) GoString() string {
  return s.String()
}

func (s *GetBasicConfigurationOfDomainResponse) SetDomainId(v int) *GetBasicConfigurationOfDomainResponse {
  s.DomainId = &v
  return s
}

func (s *GetBasicConfigurationOfDomainResponse) SetDomainName(v string) *GetBasicConfigurationOfDomainResponse {
  s.DomainName = &v
  return s
}

func (s *GetBasicConfigurationOfDomainResponse) SetCreatedDate(v string) *GetBasicConfigurationOfDomainResponse {
  s.CreatedDate = &v
  return s
}

func (s *GetBasicConfigurationOfDomainResponse) SetLastModified(v string) *GetBasicConfigurationOfDomainResponse {
  s.LastModified = &v
  return s
}

func (s *GetBasicConfigurationOfDomainResponse) SetServiceType(v string) *GetBasicConfigurationOfDomainResponse {
  s.ServiceType = &v
  return s
}

func (s *GetBasicConfigurationOfDomainResponse) SetComment(v string) *GetBasicConfigurationOfDomainResponse {
  s.Comment = &v
  return s
}

func (s *GetBasicConfigurationOfDomainResponse) SetServiceAreas(v string) *GetBasicConfigurationOfDomainResponse {
  s.ServiceAreas = &v
  return s
}

func (s *GetBasicConfigurationOfDomainResponse) SetCname(v string) *GetBasicConfigurationOfDomainResponse {
  s.Cname = &v
  return s
}

func (s *GetBasicConfigurationOfDomainResponse) SetStatus(v string) *GetBasicConfigurationOfDomainResponse {
  s.Status = &v
  return s
}

func (s *GetBasicConfigurationOfDomainResponse) SetCdnServiceStatus(v string) *GetBasicConfigurationOfDomainResponse {
  s.CdnServiceStatus = &v
  return s
}

func (s *GetBasicConfigurationOfDomainResponse) SetEnabled(v string) *GetBasicConfigurationOfDomainResponse {
  s.Enabled = &v
  return s
}

func (s *GetBasicConfigurationOfDomainResponse) SetCnameLabel(v string) *GetBasicConfigurationOfDomainResponse {
  s.CnameLabel = &v
  return s
}

func (s *GetBasicConfigurationOfDomainResponse) SetOriginConfig(v *GetBasicConfigurationOfDomainResponseOriginConfig) *GetBasicConfigurationOfDomainResponse {
  s.OriginConfig = v
  return s
}

func (s *GetBasicConfigurationOfDomainResponse) SetSsl(v *GetBasicConfigurationOfDomainResponseSsl) *GetBasicConfigurationOfDomainResponse {
  s.Ssl = v
  return s
}

func (s *GetBasicConfigurationOfDomainResponse) SetCacheBehaviors(v map[string]interface{}) *GetBasicConfigurationOfDomainResponse {
  s.CacheBehaviors = v
  return s
}

func (s *GetBasicConfigurationOfDomainResponse) SetCacheHost(v string) *GetBasicConfigurationOfDomainResponse {
  s.CacheHost = &v
  return s
}

func (s *GetBasicConfigurationOfDomainResponse) SetEnableHttpdns(v string) *GetBasicConfigurationOfDomainResponse {
  s.EnableHttpdns = &v
  return s
}

func (s *GetBasicConfigurationOfDomainResponse) SetHeaderOfClientip(v string) *GetBasicConfigurationOfDomainResponse {
  s.HeaderOfClientip = &v
  return s
}

func (s *GetBasicConfigurationOfDomainResponse) SetDomainStreamType(v string) *GetBasicConfigurationOfDomainResponse {
  s.DomainStreamType = &v
  return s
}

func (s *GetBasicConfigurationOfDomainResponse) SetLiveConfig(v *GetBasicConfigurationOfDomainResponseLiveConfig) *GetBasicConfigurationOfDomainResponse {
  s.LiveConfig = v
  return s
}

func (s *GetBasicConfigurationOfDomainResponse) SetPublishPoints(v *GetBasicConfigurationOfDomainResponsePublishPoints) *GetBasicConfigurationOfDomainResponse {
  s.PublishPoints = v
  return s
}

type GetBasicConfigurationOfDomainResponseOriginConfig struct {
  // {"en":"Return source address, which can be IP or domain name.", "zh_CN":"回源地址，可以是IP或域名。
  // 1、IP以分号分隔，支持多个。
  // 2、域名只能一个。
  // 3、限制最大不能超过500个字符长度。"}
  OriginIps *string `json:"originIps,omitempty" xml:"originIps,omitempty" require:"true"`
  // {"en":"Back to the source HOST, used to change the HOST field in the source HTTP request header.", "zh_CN":"回源HOST，用于更改回源HTTP请求头中的HOST字段。"}
  DefaultOriginHostHeader *string `json:"default-origin-host-header,omitempty" xml:"default-origin-host-header,omitempty" require:"true"`
  // {"en":"advance origin config", "zh_CN":"高级源配置"}
  AdvOriginConfigs *GetBasicConfigurationOfDomainResponseOriginConfigAdvOriginConfigs `json:"adv-origin-configs,omitempty" xml:"adv-origin-configs,omitempty" require:"true" type:"Struct"`
}

func (s GetBasicConfigurationOfDomainResponseOriginConfig) String() string {
  return tea.Prettify(s)
}

func (s GetBasicConfigurationOfDomainResponseOriginConfig) GoString() string {
  return s.String()
}

func (s *GetBasicConfigurationOfDomainResponseOriginConfig) SetOriginIps(v string) *GetBasicConfigurationOfDomainResponseOriginConfig {
  s.OriginIps = &v
  return s
}

func (s *GetBasicConfigurationOfDomainResponseOriginConfig) SetDefaultOriginHostHeader(v string) *GetBasicConfigurationOfDomainResponseOriginConfig {
  s.DefaultOriginHostHeader = &v
  return s
}

func (s *GetBasicConfigurationOfDomainResponseOriginConfig) SetAdvOriginConfigs(v *GetBasicConfigurationOfDomainResponseOriginConfigAdvOriginConfigs) *GetBasicConfigurationOfDomainResponseOriginConfig {
  s.AdvOriginConfigs = v
  return s
}

type GetBasicConfigurationOfDomainResponseOriginConfigAdvOriginConfigs struct {
  // {"en":"The advanced source monitors the url, and requests <master-ips> through the url. If the response is not 2**, 3** response, it is considered that the primary source ip is faulty, and <backup-ips> is used at this time.", "zh_CN":"高级源监控url，通过该url请求<master-ips>，如果返回非2**，3**响应时，认为主要回源ip故障，此时使用<backup-ips>。
  // 完整的url，例如：http://a.example.com/test.html"}
  DetectUrl *string `json:"detect-url,omitempty" xml:"detect-url,omitempty" require:"true"`
  // {"en":"Advanced source monitoring period, in seconds, optional as an integer greater than or equal to 0, 0 means no monitoring", "zh_CN":"高级源监控周期，单位秒，可选值为大于等于0的整数，0表示不监控"}
  DetectPeriod *int `json:"detect-period,omitempty" xml:"detect-period,omitempty" require:"true"`
  // {"en":"advance origin config", "zh_CN":"高级源配置"}
  AdvOriginConfig *GetBasicConfigurationOfDomainResponseOriginConfigAdvOriginConfigsAdvOriginConfig `json:"adv-origin-config,omitempty" xml:"adv-origin-config,omitempty" require:"true" type:"Struct"`
}

func (s GetBasicConfigurationOfDomainResponseOriginConfigAdvOriginConfigs) String() string {
  return tea.Prettify(s)
}

func (s GetBasicConfigurationOfDomainResponseOriginConfigAdvOriginConfigs) GoString() string {
  return s.String()
}

func (s *GetBasicConfigurationOfDomainResponseOriginConfigAdvOriginConfigs) SetDetectUrl(v string) *GetBasicConfigurationOfDomainResponseOriginConfigAdvOriginConfigs {
  s.DetectUrl = &v
  return s
}

func (s *GetBasicConfigurationOfDomainResponseOriginConfigAdvOriginConfigs) SetDetectPeriod(v int) *GetBasicConfigurationOfDomainResponseOriginConfigAdvOriginConfigs {
  s.DetectPeriod = &v
  return s
}

func (s *GetBasicConfigurationOfDomainResponseOriginConfigAdvOriginConfigs) SetAdvOriginConfig(v *GetBasicConfigurationOfDomainResponseOriginConfigAdvOriginConfigsAdvOriginConfig) *GetBasicConfigurationOfDomainResponseOriginConfigAdvOriginConfigs {
  s.AdvOriginConfig = v
  return s
}

type GetBasicConfigurationOfDomainResponseOriginConfigAdvOriginConfigsAdvOriginConfig struct {
  // {"en":"The advanced source mainly returns the source IP. Multiple IPs are separated by a semicolon \";\", and the return source IP cannot be repeated.", "zh_CN":"高级源主要回源IP，多个IP用分号“;”分隔，回源IP不能重复"}
  MasterIps *string `json:"master-ips,omitempty" xml:"master-ips,omitempty" require:"true"`
  // {"en":"Advanced source backup source IP, multiple IPs are separated by semicolon \";\", and the return source IP cannot be duplicated.", "zh_CN":"高级源备用回源IP，多个IP用分号“;”分隔，回源IP不能重复"}
  BackupIps *string `json:"backup-ips,omitempty" xml:"backup-ips,omitempty" require:"true"`
}

func (s GetBasicConfigurationOfDomainResponseOriginConfigAdvOriginConfigsAdvOriginConfig) String() string {
  return tea.Prettify(s)
}

func (s GetBasicConfigurationOfDomainResponseOriginConfigAdvOriginConfigsAdvOriginConfig) GoString() string {
  return s.String()
}

func (s *GetBasicConfigurationOfDomainResponseOriginConfigAdvOriginConfigsAdvOriginConfig) SetMasterIps(v string) *GetBasicConfigurationOfDomainResponseOriginConfigAdvOriginConfigsAdvOriginConfig {
  s.MasterIps = &v
  return s
}

func (s *GetBasicConfigurationOfDomainResponseOriginConfigAdvOriginConfigsAdvOriginConfig) SetBackupIps(v string) *GetBasicConfigurationOfDomainResponseOriginConfigAdvOriginConfigsAdvOriginConfig {
  s.BackupIps = &v
  return s
}

type GetBasicConfigurationOfDomainResponseSsl struct {
  // {"en":"Use a certificate, the optional values are true and false, true means to use the certificate, false means not to use the certificate", "zh_CN":"使用证书，可选值为true和false，true表示使用证书，false表示不使用证书"}
  UseSsl *string `json:"use-ssl,omitempty" xml:"use-ssl,omitempty" require:"true"`
  // {"en":"Use sni certificate, the optional values are true and false, true means use sni certificate, false means use non-sni traditional certificate", "zh_CN":"使用sni证书，可选值为true和false，true表示使用sni证书，false表示使用非sni的传统证书"}
  UseForSni *string `json:"use-for-sni,omitempty" xml:"use-for-sni,omitempty" require:"true"`
  // {"en":"Certificate ID, the certificate ID returned by the system after the new certificate is successfully added.", "zh_CN":"证书ID，新增证书成功后，系统返回的证书ID"}
  SslCertificateId *int `json:"ssl-certificate-id,omitempty" xml:"ssl-certificate-id,omitempty" require:"true"`
}

func (s GetBasicConfigurationOfDomainResponseSsl) String() string {
  return tea.Prettify(s)
}

func (s GetBasicConfigurationOfDomainResponseSsl) GoString() string {
  return s.String()
}

func (s *GetBasicConfigurationOfDomainResponseSsl) SetUseSsl(v string) *GetBasicConfigurationOfDomainResponseSsl {
  s.UseSsl = &v
  return s
}

func (s *GetBasicConfigurationOfDomainResponseSsl) SetUseForSni(v string) *GetBasicConfigurationOfDomainResponseSsl {
  s.UseForSni = &v
  return s
}

func (s *GetBasicConfigurationOfDomainResponseSsl) SetSslCertificateId(v int) *GetBasicConfigurationOfDomainResponseSsl {
  s.SslCertificateId = &v
  return s
}

type GetBasicConfigurationOfDomainResponseLiveConfig struct {
  // {"en":"The live push-pull stream type, the optional values are pull and push.", "zh_CN":"直播推拉流类型，可选值为pull和push，pull表示拉流；   push表示推流。"}
  StreamType *string `json:"stream-type,omitempty" xml:"stream-type,omitempty" require:"true"`
  // {"en":"Source station IP. When the stream-type is pull, at least one of the source station IP and the companion push stream domain name is not empty.", "zh_CN":"源站IP，当stream-type为pull时，源站IP和配套推流域名至少一个不为空。
  // 1、如果是推拉流配套，则返回127.0.0.1
  // 2、如果是直接回源拉流，则返回源站IP"}
  OriginIps *string `json:"origin-ips,omitempty" xml:"origin-ips,omitempty" require:"true"`
  // {"en":"The corresponding push domain name, the rtmp live streaming domain name corresponding to the push domain name", "zh_CN":"配套推流域名，rtmp直播拉流域名对应的推流域名"}
  OriginPushHost *string `json:"origin-push-host,omitempty" xml:"origin-push-host,omitempty" require:"true"`
}

func (s GetBasicConfigurationOfDomainResponseLiveConfig) String() string {
  return tea.Prettify(s)
}

func (s GetBasicConfigurationOfDomainResponseLiveConfig) GoString() string {
  return s.String()
}

func (s *GetBasicConfigurationOfDomainResponseLiveConfig) SetStreamType(v string) *GetBasicConfigurationOfDomainResponseLiveConfig {
  s.StreamType = &v
  return s
}

func (s *GetBasicConfigurationOfDomainResponseLiveConfig) SetOriginIps(v string) *GetBasicConfigurationOfDomainResponseLiveConfig {
  s.OriginIps = &v
  return s
}

func (s *GetBasicConfigurationOfDomainResponseLiveConfig) SetOriginPushHost(v string) *GetBasicConfigurationOfDomainResponseLiveConfig {
  s.OriginPushHost = &v
  return s
}

type GetBasicConfigurationOfDomainResponsePublishPoints struct {
  // {"en":"Release point, support multiple, the system default is \"/\"", "zh_CN":"发布点，支持多个，系统默认值为“/”"}
  Uri *string `json:"uri,omitempty" xml:"uri,omitempty" require:"true"`
}

func (s GetBasicConfigurationOfDomainResponsePublishPoints) String() string {
  return tea.Prettify(s)
}

func (s GetBasicConfigurationOfDomainResponsePublishPoints) GoString() string {
  return s.String()
}

func (s *GetBasicConfigurationOfDomainResponsePublishPoints) SetUri(v string) *GetBasicConfigurationOfDomainResponsePublishPoints {
  s.Uri = &v
  return s
}

type GetBasicConfigurationOfDomainPaths struct {
  // {"en":"Accelerated domain name or domain ID", "zh_CN":"域名名称或域名id"}
  Domain *string `json:"domain-name,omitempty" xml:"domain-name,omitempty" require:"true"`
}

func (s GetBasicConfigurationOfDomainPaths) String() string {
  return tea.Prettify(s)
}

func (s GetBasicConfigurationOfDomainPaths) GoString() string {
  return s.String()
}

func (s *GetBasicConfigurationOfDomainPaths) SetDomain(v string) *GetBasicConfigurationOfDomainPaths {
  s.Domain = &v
  return s
}

type GetBasicConfigurationOfDomainParameters struct {
}

func (s GetBasicConfigurationOfDomainParameters) String() string {
  return tea.Prettify(s)
}

func (s GetBasicConfigurationOfDomainParameters) GoString() string {
  return s.String()
}

type GetBasicConfigurationOfDomainRequestHeader struct {
}

func (s GetBasicConfigurationOfDomainRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetBasicConfigurationOfDomainRequestHeader) GoString() string {
  return s.String()
}

type GetBasicConfigurationOfDomainResponseHeader struct {
  // {"en":"Httpstatus=200 indicates that the interface is successfully invoked.", "zh_CN":"httpstatus=200;   表示成功调用接口"}
  HttpStatusCode *int `json:"http status code,omitempty" xml:"http status code,omitempty" require:"true"`
  // {"en":"Uniquely labeled id for querying each requested task (for all interfaces)", "zh_CN":"唯一标示的id，用于查询每次请求的任务 （适用全部接口）"}
  XCncRequestId *string `json:"x-cnc-request-id,omitempty" xml:"x-cnc-request-id,omitempty" require:"true"`
  // {"en":"The corresponding deployment version number of this modification", "zh_CN":"本次修改对应的部署版本号"}
  XCncDeployVersion *string `json:"x-cnc-deploy-version,omitempty" xml:"x-cnc-deploy-version,omitempty" require:"true"`
}

func (s GetBasicConfigurationOfDomainResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetBasicConfigurationOfDomainResponseHeader) GoString() string {
  return s.String()
}

func (s *GetBasicConfigurationOfDomainResponseHeader) SetHttpStatusCode(v int) *GetBasicConfigurationOfDomainResponseHeader {
  s.HttpStatusCode = &v
  return s
}

func (s *GetBasicConfigurationOfDomainResponseHeader) SetXCncRequestId(v string) *GetBasicConfigurationOfDomainResponseHeader {
  s.XCncRequestId = &v
  return s
}

func (s *GetBasicConfigurationOfDomainResponseHeader) SetXCncDeployVersion(v string) *GetBasicConfigurationOfDomainResponseHeader {
  s.XCncDeployVersion = &v
  return s
}




type EditDomainConfigRequest struct {
  // {"en":"Version, the current version is 1.0.0", "zh_CN":"版本号，当前版本号1.0.0"}
  Version *string `json:"version,omitempty" xml:"version,omitempty"`
  // {"en":"Remarks, up to 1000 characters.", "zh_CN":"备注信息，最大限制1000个字符"}
  Comment *string `json:"comment,omitempty" xml:"comment,omitempty"`
  // {"en":"The acceleration area of the acceleration domain, if the resource coverage needs to be limited according to the area, the acceleration area needs to be specified. When no acceleration area is specified, we will provide acceleration services with optimal resource coverage according to the service area opened by the customer. Multiple regions are separated by semicolons, and the supported regions are as follows: cn (Mainland China), am (Americas), emea (Europe, Middle East, Africa), apac (Asia-Pacific region).", 
  //     "zh_CN":"加速域名的加速区域，如果有需要根据区域限定资源覆盖时，才需要指定加速区域。未指定加速区域时，我们将按照客户开通的服务区域，以最优的资源覆盖提供加速服务。多个区域以分号分隔，支持配置的区域如下：cn（中国大陆）、am（美洲）、emea（欧洲、中东、非洲）、apac（亚太地区）"}
  ServiceAreas *string `json:"service-areas,omitempty" xml:"service-areas,omitempty"`
  // {"en":"If you need to share a CNAME between domains, you can use this parameter. This parameter is a unique label for a public CNAME. Domains with the same cname-label will have the same CNAME. 
  // Note:
  // 1. Domains with the same cname-label have the same coverage.
  // 2. Constraints of sharing a CNAME: consistent service-type, consistent certificate-id (if there is a certificate), consistent service-areas
  // 3. Multiple http domains can share a CNAME, multiple sni https domains can share a CNAME too.
  // 4. When a cname-label is used by a single domain, then the domain can be canceled acceleration. While a cname-label using by more then one domains, they can not be canceled acceleration.
  // 5. Support the purpose of modifying cname by modifying cname-label. )", 
  // "zh_CN":"共用一级标签，若有多个加速域名需要共用一级域名，则可以使用该参数。即拥有相同cname-label的一组域名，共用一级cname。
  // 注意：
  // 1、拥有相同cname-label的域名共用一级cname，且有完全一致的dns覆盖
  // 2、共用一级的约束：加速类型一致(service-type)、证书id一致（certificate-id,如果有证书）、加速区域一致(service-areas)
  // 3、多个http域名可共用一级，多个sni https域名可共用一级
  // 4、单个域名使用cname-label时，域名可cancel；多个域名共用一级时，不允许cancel这些域名
  // 5、支持通过修改cname-label达到修改cname的目的。"}
  CnameLabel *string `json:"cname-label,omitempty" xml:"cname-label,omitempty"`
  // {"en":"Back to origin policy settings for setting source site information and return source policies for accelerated domain names", "zh_CN":"回源策略设置，用于设置加速域名的源站信息和回源策略"}
  OriginConfig *EditDomainConfigRequestOriginConfig `json:"origin-config,omitempty" xml:"origin-config,omitempty" type:"Struct"`
  // {"en":"SSL settings, to bind a certificate with the accelerated domain. You can use the interface [AddCertificate] to upload your  certificates. If you want to modify a certificate, please use the interface: [UpdateCertificate]", "zh_CN":"ssl证书设置，用于设置加速域名的ssl证书配置。上传证书请使用接口：【新增证书V2】；若要修改证书，请使用接口：【修改证书V2】"}
  Ssl *EditDomainConfigRequestSsl `json:"ssl,omitempty" xml:"ssl,omitempty" type:"Struct"`
  // {"en":"Cache file HOST.
  // Cache rules for caching HOST domain names and accelerated domain names must be consistent.", "zh_CN":"缓存文件HOST。缓存HOST域名和加速域名的缓存规则必须一致。"}
  CacheHost *string `json:"cache-host,omitempty" xml:"cache-host,omitempty"`
  // {"en":"Enable httpdns settings.
  // The optional values are true and false, true means enabled; false means off. This function is not enabled by default. If you need it, please contact the technical support to apply for this feature.", "zh_CN":"启用httpdns设置（使用需申请）
  // 可选值为true和false，true表示启用；false表示关闭
  // 注意：该功能默认不开启，若您有需要，请联系专属客服申请开通。"}
  EnableHttpdns *string `json:"enable-httpdns,omitempty" xml:"enable-httpdns,omitempty"`
  // {"en":"Pass the response header of client IP. The optional values are Cdn-Src-Ip and X-Forwarded-For. The default value is Cdn-Src-Ip.", "zh_CN":"传递客户端ip的响应头部，可选值为Cdn-Src-Ip和X-Forwarded-For，默认值为Cdn-Src-Ip"}
  HeaderOfClientip *string `json:"header-of-clientip,omitempty" xml:"header-of-clientip,omitempty"`
  // {"en":"Live domain name configuration, used to set the push flow of live acceleration domain name", "zh_CN":"直播域名配置，用于设置直播加速域名的推拉流"}
  LiveConfig *EditDomainConfigRequestLiveConfig `json:"live-config,omitempty" xml:"live-config,omitempty" type:"Struct"`
  // {"en":"Set the publishing point of the live push-pull domain.
  // Note:
  // 1. The pull stream and the corresponding push stream domain must be configured with the same publishing point.
  // 2. If you are not going to modify the publishing point, please do not pass this param.
  // 3. The publishing point adopts the overlay update. Each time you modify, you need to submit all the publishing points. You cannot submit only the parts that need to be modified.", "zh_CN":"设置直播推拉流域名的发布点
  // 注意：
  // 1、拉流和对应的推流域名，必须配置相同的发布点；
  // 2、不想修改发布点时，不要传入该节点及以下入参；
  // 3、发布点采用覆盖式更新，每次修改时，需要提交全部发布点，不能仅提交需要修改的部分。"}
  PublishPoints []*EditDomainConfigRequestPublishPoints `json:"publish-points,omitempty" xml:"publish-points,omitempty" type:"Repeated"`
}

func (s EditDomainConfigRequest) String() string {
  return tea.Prettify(s)
}

func (s EditDomainConfigRequest) GoString() string {
  return s.String()
}

func (s *EditDomainConfigRequest) SetVersion(v string) *EditDomainConfigRequest {
  s.Version = &v
  return s
}

func (s *EditDomainConfigRequest) SetComment(v string) *EditDomainConfigRequest {
  s.Comment = &v
  return s
}

func (s *EditDomainConfigRequest) SetServiceAreas(v string) *EditDomainConfigRequest {
  s.ServiceAreas = &v
  return s
}

func (s *EditDomainConfigRequest) SetCnameLabel(v string) *EditDomainConfigRequest {
  s.CnameLabel = &v
  return s
}

func (s *EditDomainConfigRequest) SetOriginConfig(v *EditDomainConfigRequestOriginConfig) *EditDomainConfigRequest {
  s.OriginConfig = v
  return s
}

func (s *EditDomainConfigRequest) SetSsl(v *EditDomainConfigRequestSsl) *EditDomainConfigRequest {
  s.Ssl = v
  return s
}

func (s *EditDomainConfigRequest) SetCacheHost(v string) *EditDomainConfigRequest {
  s.CacheHost = &v
  return s
}

func (s *EditDomainConfigRequest) SetEnableHttpdns(v string) *EditDomainConfigRequest {
  s.EnableHttpdns = &v
  return s
}

func (s *EditDomainConfigRequest) SetHeaderOfClientip(v string) *EditDomainConfigRequest {
  s.HeaderOfClientip = &v
  return s
}

func (s *EditDomainConfigRequest) SetLiveConfig(v *EditDomainConfigRequestLiveConfig) *EditDomainConfigRequest {
  s.LiveConfig = v
  return s
}

func (s *EditDomainConfigRequest) SetPublishPoints(v []*EditDomainConfigRequestPublishPoints) *EditDomainConfigRequest {
  s.PublishPoints = v
  return s
}

type EditDomainConfigRequestOriginConfig struct {
  // {"en":"Origin address, which can be an IP or domain name.
  // 1. Multiple IPs are supported, separated by semicolons.
  // 2. Only one domain name is allowed. IP and domain name cannot exist at the same time.
  // 3. The length cannot exceed 500 characters.
  // 4. The number of IPs cannot exceed 15.
  // ", "zh_CN":"回源地址，可以是IP或域名。
  // 1、IP以分号分隔，支持多个。
  // 2、域名只能输入一个。IP与域名不能同时输入。
  // 3、限制最大不能超过500个字符长度。
  // 4、IP最多15个。"}
  OriginIps *string `json:"origin-ips,omitempty" xml:"origin-ips,omitempty"`
  // {"en":"Back-to-origin HOST, used to change the HOST field in the back-to-origin HTTP request header. The supported formats are: ① domain name ③ ip 
  // Note:
  // 1. Must comply with the ip/domain name format specification. If it is a domain name, the length of the domain name must be less than or equal to 128 characters.", "zh_CN":"回源HOST，用于更改回源HTTP请求头中的HOST字段。支持格式为: ①域名；②ip；
  // 注意：
  // 1、必须符合ip/域名格式规范。如果是域名，则域名长度小于等于128。"}
  DefaultOriginHostHeader *string `json:"default-origin-host-header,omitempty" xml:"default-origin-host-header,omitempty"`
}

func (s EditDomainConfigRequestOriginConfig) String() string {
  return tea.Prettify(s)
}

func (s EditDomainConfigRequestOriginConfig) GoString() string {
  return s.String()
}

func (s *EditDomainConfigRequestOriginConfig) SetOriginIps(v string) *EditDomainConfigRequestOriginConfig {
  s.OriginIps = &v
  return s
}

func (s *EditDomainConfigRequestOriginConfig) SetDefaultOriginHostHeader(v string) *EditDomainConfigRequestOriginConfig {
  s.DefaultOriginHostHeader = &v
  return s
}

type EditDomainConfigRequestSsl struct {
  // {"en":"Use a certificate, the optional values are true and false, true means to use the certificate, false means not to use the certificate", "zh_CN":"使用证书，可选值为true和false，true表示使用证书，false表示不使用证书"}
  UseSsl *string `json:"use-ssl,omitempty" xml:"use-ssl,omitempty"`
  // {"en":"Use sni certificate, the optional values are true and false, true means use sni certificate, false means use shared certificate (not supported)", "zh_CN":"使用sni证书，可选值为true和false，true表示使用sni证书，false表示使用合用证书（暂不支持）"}
  UseForSni *string `json:"use-for-sni,omitempty" xml:"use-for-sni,omitempty"`
  // {"en":"Use sni certificate, the optional values are true and false, true means use sni certificate, false means use shared certificate (not supported)", "zh_CN":"证书ID，新增证书成功后，系统返回的证书ID
  // use-ssl为true时，才能传ssl-certificate-id。"}
  SslCertificateId *int `json:"ssl-certificate-id,omitempty" xml:"ssl-certificate-id,omitempty"`
}

func (s EditDomainConfigRequestSsl) String() string {
  return tea.Prettify(s)
}

func (s EditDomainConfigRequestSsl) GoString() string {
  return s.String()
}

func (s *EditDomainConfigRequestSsl) SetUseSsl(v string) *EditDomainConfigRequestSsl {
  s.UseSsl = &v
  return s
}

func (s *EditDomainConfigRequestSsl) SetUseForSni(v string) *EditDomainConfigRequestSsl {
  s.UseForSni = &v
  return s
}

func (s *EditDomainConfigRequestSsl) SetSslCertificateId(v int) *EditDomainConfigRequestSsl {
  s.SslCertificateId = &v
  return s
}

type EditDomainConfigRequestLiveConfig struct {
  // {"en":"Source station IP. When the stream-type is pull, at least one of the source station IP and the companion push stream domain name is not empty.
  // 1. If it is a push-pull flow package, fill in 127.0.0.1, and the system will also default to 127.0.0.1.
  // 2. If it is directly returning to the source, fill in the source IP of the source pull stream.
  // ", "zh_CN":"源站IP，当stream-type为pull时，源站IP和配套推流域名至少一个不为空。
  // 1、如果是推拉流配套，则填写127.0.0.1，不传系统也默认为127.0.0.1
  // 2、如果是直接回源拉流，则填写回源拉流的源站IP"}
  LiveConfigOriginIps *string `json:"origin-ips,omitempty" xml:"origin-ips,omitempty"`
  // {"en":"A matching push domain name is used to set up a current domain name corresponding to the live streaming domain name. When the stream-type is pull, the source station IP and the supporting current domain name are at least one empty; when stream-type is push, it does not need to be introduced.", "zh_CN":"配套推流域名，用于设置直播拉流域名对应的推流域名，当stream-type为pull时，源站IP和配套推流域名至少一个不为空；当stream-type为push时，无需传入。"}
  OriginPushHost *string `json:"origin-push-host,omitempty" xml:"origin-push-host,omitempty"`
}

func (s EditDomainConfigRequestLiveConfig) String() string {
  return tea.Prettify(s)
}

func (s EditDomainConfigRequestLiveConfig) GoString() string {
  return s.String()
}

func (s *EditDomainConfigRequestLiveConfig) SetLiveConfigOriginIps(v string) *EditDomainConfigRequestLiveConfig {
  s.LiveConfigOriginIps = &v
  return s
}

func (s *EditDomainConfigRequestLiveConfig) SetOriginPushHost(v string) *EditDomainConfigRequestLiveConfig {
  s.OriginPushHost = &v
  return s
}

type EditDomainConfigRequestPublishPoints struct     {
  // {"en":"Livestream domain settings. Publish point, support multiple, do not pass the system by default to generate a publishing point uri for [/]", "zh_CN":"发布点，支持多个，不传系统默认生成一条发布点uri为“/”"}
  Uri *string `json:"uri,omitempty" xml:"uri,omitempty"`
}

func (s EditDomainConfigRequestPublishPoints) String() string {
  return tea.Prettify(s)
}

func (s EditDomainConfigRequestPublishPoints) GoString() string {
  return s.String()
}

func (s *EditDomainConfigRequestPublishPoints) SetUri(v string) *EditDomainConfigRequestPublishPoints {
  s.Uri = &v
  return s
}

type EditDomainConfigResponse struct {
  // {"en":"If httpstatus=202, the interface is successfully invoked. You can use the x-cnc-request-id in the header to check the deployment of domain.", "zh_CN":"httpstatus=202;   表示成功调用接口，可使用header中的x-cnc-request-id查看域名的部署情况"}
  HttpStatusCode *int `json:"http status code,omitempty" xml:"http status code,omitempty" require:"true"`
  // {"en":"Uniquely labeled id for querying each requested task (for all interfaces)", "zh_CN":"唯一标示的id，用于查询每次请求的任务 （适用全部接口）"}
  XCncRequestId *string `json:"x-cnc-request-id,omitempty" xml:"x-cnc-request-id,omitempty" require:"true"`
  // {"en":"The corresponding deployment version number of this modification", "zh_CN":"本次修改对应的部署版本号"}
  XCncDeployVersion *string `json:"x-cnc-deploy-version,omitempty" xml:"x-cnc-deploy-version,omitempty" require:"true"`
  // {"en":"Error code, which appears when HTTPStatus is not 202, represents the error type of the current request call", "zh_CN":"错误代码，当HTTPStatus不为202时出现，表示当前请求调用的错误类型"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response information, success when successful", "zh_CN":"响应信息，成功时为success"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s EditDomainConfigResponse) String() string {
  return tea.Prettify(s)
}

func (s EditDomainConfigResponse) GoString() string {
  return s.String()
}

func (s *EditDomainConfigResponse) SetHttpStatusCode(v int) *EditDomainConfigResponse {
  s.HttpStatusCode = &v
  return s
}

func (s *EditDomainConfigResponse) SetXCncRequestId(v string) *EditDomainConfigResponse {
  s.XCncRequestId = &v
  return s
}

func (s *EditDomainConfigResponse) SetXCncDeployVersion(v string) *EditDomainConfigResponse {
  s.XCncDeployVersion = &v
  return s
}

func (s *EditDomainConfigResponse) SetCode(v string) *EditDomainConfigResponse {
  s.Code = &v
  return s
}

func (s *EditDomainConfigResponse) SetMessage(v string) *EditDomainConfigResponse {
  s.Message = &v
  return s
}

type EditDomainConfigPaths struct {
  // {"en":"The domain you are going to modify, it can be domain id or domain name.", "zh_CN":"需要修改的域名，可以是域名名称或域名id。"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
}

func (s EditDomainConfigPaths) String() string {
  return tea.Prettify(s)
}

func (s EditDomainConfigPaths) GoString() string {
  return s.String()
}

func (s *EditDomainConfigPaths) SetDomain(v string) *EditDomainConfigPaths {
  s.Domain = &v
  return s
}

type EditDomainConfigParameters struct {
}

func (s EditDomainConfigParameters) String() string {
  return tea.Prettify(s)
}

func (s EditDomainConfigParameters) GoString() string {
  return s.String()
}

type EditDomainConfigRequestHeader struct {
}

func (s EditDomainConfigRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s EditDomainConfigRequestHeader) GoString() string {
  return s.String()
}

type EditDomainConfigResponseHeader struct {
}

func (s EditDomainConfigResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s EditDomainConfigResponseHeader) GoString() string {
  return s.String()
}




type QueryInnerRedirectRequest struct {
}

func (s QueryInnerRedirectRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryInnerRedirectRequest) GoString() string {
  return s.String()
}

type QueryInnerRedirectResponse struct {
  // {"en":"the domain whoes need query config", "zh_CN":"需要查询配置的域名id"}
  DomainId *string `json:"domain-id,omitempty" xml:"domain-id,omitempty" require:"true"`
  // {"en":"the domain whoes need query config", "zh_CN":"需要查询配置的域名"}
  DomainName *string `json:"domain-name,omitempty" xml:"domain-name,omitempty" require:"true"`
  RewriteRuleSettings []*QueryInnerRedirectResponseRewriteRuleSettings `json:"rewrite-rule-settings,omitempty" xml:"rewrite-rule-settings,omitempty" require:"true" type:"Repeated"`
}

func (s QueryInnerRedirectResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryInnerRedirectResponse) GoString() string {
  return s.String()
}

func (s *QueryInnerRedirectResponse) SetDomainId(v string) *QueryInnerRedirectResponse {
  s.DomainId = &v
  return s
}

func (s *QueryInnerRedirectResponse) SetDomainName(v string) *QueryInnerRedirectResponse {
  s.DomainName = &v
  return s
}

func (s *QueryInnerRedirectResponse) SetRewriteRuleSettings(v []*QueryInnerRedirectResponseRewriteRuleSettings) *QueryInnerRedirectResponse {
  s.RewriteRuleSettings = v
  return s
}

type QueryInnerRedirectResponseRewriteRuleSettings struct     {
  // {"en":"Add a grid type identifier to indicate a specific group configuration when the client has multiple groups of configurations.", "zh_CN":"添加grid类型标识，表示客户多组配置时，具体某组配置；data-id重复，已入参同个id最后一组为准生效"}
  DataId *string `json:"data-id,omitempty" xml:"data-id,omitempty" require:"true"`
  // {"en":"The url matching mode supports fuzzy regularization. If all matches, the input parameters can be configured as: *", "zh_CN":"url匹配模式，支持正则,  .*：匹配所有文件
  // 客户入参参考：.*
  // 对于匹配到的URL进行内容重定向"}
  PathPattern *string `json:"path-pattern,omitempty" xml:"path-pattern,omitempty" require:"true"`
  // {"en":"Exceptional url matching mode, except for certain URLs: such as abc.jpg, no content redirection
  // Customer reference: ^https?://[^/]+/.*\.m3u8", "zh_CN":"例外的url匹配模式，某些URL除外：如abc.jpg，不做内容重定向
  // 客户入参参考：^https?://[^/]+/.*\.m3u8"}
  ExceptPathPattern *string `json:"except-path-pattern,omitempty" xml:"except-path-pattern,omitempty" require:"true"`
  // {"en":"Matching conditions: specify common types, optional values are all or homepage 1. all: all files 2. homepage: home page", "zh_CN":"匹配条件：指定常用类型，可选值为all或homepage 1、all：全部文件 2、homepage：首页"}
  CustomPattern *string `json:"custom-pattern,omitempty" xml:"custom-pattern,omitempty" require:"true"`
  // {"en":"directory", "zh_CN":"目录"}
  Directory *string `json:"directory,omitempty" xml:"directory,omitempty" require:"true"`
  // {"en":"gif png bmp jpeg jpg html htm shtml mp3 wma flv mp4 wmv zip exe rar css txt ico js swf m3u8 xml f4m bootstarp ts", "zh_CN":"匹配条件：文件类型，多个请以英文;分隔，可选值：gif png bmp jpeg jpg html htm shtml mp3 wma flv mp4 wmv zip exe rar css txt ico js swf m3u8 xml f4m bootstarp ts"}
  FileType *string `json:"file-type,omitempty" xml:"file-type,omitempty" require:"true"`
  // {"en":"Ignore case, the optional value is true or false, true means to ignore case; false means not to ignore case;
  // When adding a new configuration item, the default is not true.
  // If the client passes a null value: such as <ignore-letter-case></ignore-letter-case>, the configuration is cleared.", "zh_CN":"忽略大小写，可选值为true或false，true表示忽略大小写；false表示不忽略大小写；
  // 新增配置项时，不传默认为 true
  // 如果客户传了空值：如<ignore-letter-case></ignore-letter-case>，则表示清空配置"}
  IgnoreLetterCase *string `json:"ignore-letter-case,omitempty" xml:"ignore-letter-case,omitempty" require:"true"`
  // {"en":"Rewrite the location where the content is generated. The input value is: Cache indicates the node;
  // Other input formats are not supported at this time", "zh_CN":"改写内容的生成位置。可输入值为：Cache表示节点；
  // 暂不支持其他入参格式"}
  PublishType *string `json:"publish-type,omitempty" xml:"publish-type,omitempty" require:"true"`
  // {"en":"Indicates the priority execution order of multiple sets of redirected content by the customer. The higher the number, the higher the priority.
  // When adding a new configuration item, the default is 10", "zh_CN":"表示客户多组重定向内容的优先执行顺序。数字越大，优先级越高。
  // 新增配置项时，不传默认为 10"}
  Priority *int `json:"priority,omitempty" xml:"priority,omitempty" require:"true"`
  // {"en":"Configuration item: old url
  // Indicates the protocol mode before rewriting (that is, the object that needs to be rewritten), such as: ^https://([^/]+/.*)", "zh_CN":"配置项：旧url
  // 表示改写前的协议方式（即需要改写的对象），如：^https://([^/]+/.*)"}
  BeforeValue *string `json:"before-value,omitempty" xml:"before-value,omitempty" require:"true"`
  // {"en":"Configuration item: new url
  // Indicates the protocol method after rewriting, such as: http://$1", "zh_CN":"配置项：新url
  // 表示改写后的协议方式，如：http://$1"}
  AfterValue *string `json:"after-value,omitempty" xml:"after-value,omitempty" require:"true"`
  // {"en":"Redirection type; support for input:
  // before: before the anti-theft chain
  // after: after the anti-theft chain", "zh_CN":"重定向类型；支持入参：
  // before：防盗链之前
  // after：防盗链之后"}
  RewriteType *string `json:"rewrite-type,omitempty" xml:"rewrite-type,omitempty" require:"true"`
  // {"en":"Matching condition: Request header", "zh_CN":"匹配条件：请求头"}
  QueryInnerRedirectRequestHeader *string `json:"request-header,omitempty" xml:"request-header,omitempty" require:"true"`
  // {"en":"Matching condition: Exception request header", "zh_CN":"匹配条件：例外的请求头"}
  ExceptionQueryInnerRedirectRequestHeader *string `json:"exception-request-header,omitempty" xml:"exception-request-header,omitempty" require:"true"`
}

func (s QueryInnerRedirectResponseRewriteRuleSettings) String() string {
  return tea.Prettify(s)
}

func (s QueryInnerRedirectResponseRewriteRuleSettings) GoString() string {
  return s.String()
}

func (s *QueryInnerRedirectResponseRewriteRuleSettings) SetDataId(v string) *QueryInnerRedirectResponseRewriteRuleSettings {
  s.DataId = &v
  return s
}

func (s *QueryInnerRedirectResponseRewriteRuleSettings) SetPathPattern(v string) *QueryInnerRedirectResponseRewriteRuleSettings {
  s.PathPattern = &v
  return s
}

func (s *QueryInnerRedirectResponseRewriteRuleSettings) SetExceptPathPattern(v string) *QueryInnerRedirectResponseRewriteRuleSettings {
  s.ExceptPathPattern = &v
  return s
}

func (s *QueryInnerRedirectResponseRewriteRuleSettings) SetCustomPattern(v string) *QueryInnerRedirectResponseRewriteRuleSettings {
  s.CustomPattern = &v
  return s
}

func (s *QueryInnerRedirectResponseRewriteRuleSettings) SetDirectory(v string) *QueryInnerRedirectResponseRewriteRuleSettings {
  s.Directory = &v
  return s
}

func (s *QueryInnerRedirectResponseRewriteRuleSettings) SetFileType(v string) *QueryInnerRedirectResponseRewriteRuleSettings {
  s.FileType = &v
  return s
}

func (s *QueryInnerRedirectResponseRewriteRuleSettings) SetIgnoreLetterCase(v string) *QueryInnerRedirectResponseRewriteRuleSettings {
  s.IgnoreLetterCase = &v
  return s
}

func (s *QueryInnerRedirectResponseRewriteRuleSettings) SetPublishType(v string) *QueryInnerRedirectResponseRewriteRuleSettings {
  s.PublishType = &v
  return s
}

func (s *QueryInnerRedirectResponseRewriteRuleSettings) SetPriority(v int) *QueryInnerRedirectResponseRewriteRuleSettings {
  s.Priority = &v
  return s
}

func (s *QueryInnerRedirectResponseRewriteRuleSettings) SetBeforeValue(v string) *QueryInnerRedirectResponseRewriteRuleSettings {
  s.BeforeValue = &v
  return s
}

func (s *QueryInnerRedirectResponseRewriteRuleSettings) SetAfterValue(v string) *QueryInnerRedirectResponseRewriteRuleSettings {
  s.AfterValue = &v
  return s
}

func (s *QueryInnerRedirectResponseRewriteRuleSettings) SetRewriteType(v string) *QueryInnerRedirectResponseRewriteRuleSettings {
  s.RewriteType = &v
  return s
}

func (s *QueryInnerRedirectResponseRewriteRuleSettings) SetQueryInnerRedirectRequestHeader(v string) *QueryInnerRedirectResponseRewriteRuleSettings {
  s.QueryInnerRedirectRequestHeader = &v
  return s
}

func (s *QueryInnerRedirectResponseRewriteRuleSettings) SetExceptionQueryInnerRedirectRequestHeader(v string) *QueryInnerRedirectResponseRewriteRuleSettings {
  s.ExceptionQueryInnerRedirectRequestHeader = &v
  return s
}

type QueryInnerRedirectPaths struct {
  // {"en":"the domain whoes need query config", "zh_CN":"需要查询配置的域名或域名id"}
  DomainName *string `json:"domain-name,omitempty" xml:"domain-name,omitempty" require:"true"`
}

func (s QueryInnerRedirectPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryInnerRedirectPaths) GoString() string {
  return s.String()
}

func (s *QueryInnerRedirectPaths) SetDomainName(v string) *QueryInnerRedirectPaths {
  s.DomainName = &v
  return s
}

type QueryInnerRedirectParameters struct {
}

func (s QueryInnerRedirectParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryInnerRedirectParameters) GoString() string {
  return s.String()
}

type QueryInnerRedirectRequestHeader struct {
}

func (s QueryInnerRedirectRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryInnerRedirectRequestHeader) GoString() string {
  return s.String()
}

type QueryInnerRedirectResponseHeader struct {
}

func (s QueryInnerRedirectResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryInnerRedirectResponseHeader) GoString() string {
  return s.String()
}




type EditDomainRedirectConfigRequest struct {
  // {"en":"redirection function
  // note:
  // 1. Define a set of internal redirected content. If there is internal redirected content, this field is required.
  // 2. need to clear the content redirection content under the domain name, you can pass the empty node <rewrite-rule-settings></rewrite-rule-settings>", "zh_CN":"一级服务改写替换&mdash;二级服务 内部重定向
  // 注意：
  // 1. 定义一组内部重定向内容，，如果有使用内部重定向内容，此项必填
  // 2. 需要清空域名下的内容重定向内容，可以传入空节点<rewrite-rule-settings></rewrite-rule-settings>
  // 3. 如果有开启其他高级配置（如防盗链配置），有些配置可能会有配置冲突，建议先与技术支持人员确认"}
  RewriteRuleSettings []*EditDomainRedirectConfigRequestRewriteRuleSettings `json:"rewrite-rule-settings,omitempty" xml:"rewrite-rule-settings,omitempty" require:"true" type:"Repeated"`
}

func (s EditDomainRedirectConfigRequest) String() string {
  return tea.Prettify(s)
}

func (s EditDomainRedirectConfigRequest) GoString() string {
  return s.String()
}

func (s *EditDomainRedirectConfigRequest) SetRewriteRuleSettings(v []*EditDomainRedirectConfigRequestRewriteRuleSettings) *EditDomainRedirectConfigRequest {
  s.RewriteRuleSettings = v
  return s
}

type EditDomainRedirectConfigRequestRewriteRuleSettings struct     {
  // {"en":"Add a grid type identifier to indicate a specific group configuration when the client has multiple groups of configurations.", "zh_CN":"添加grid类型标识，表示客户多组配置时，具体某组配置；data-id重复，以入参同个id最后一组为准生效
  // data-id可以通过查询接口获取。
  // 注意：添加grid类型标识：data-id，每一组配置对应一个data-id：
  // a、如果客户有传data-id，说明指定修改其中一组配置项内容，不需求修改其他组配置内容不需要入参；
  // b、如果客户入参多组配置，其中有些组配置有传data-id，有些没有传，则有传data-id的表示修改具体某组配置，没有传data-id的表示在原来基础上新增一组配置；
  // c、如果客户入参都没有传data-id,表示用本次的配置全量覆盖原先配置；
  // d、如果客户入参没有传任何配置项参数，只传了域名和二级标签，表示清空这个接口对应域名二级服务所有配置。（c、d内容和当前方案实现一致）；
  // e、一个gird标签下的入参不能为空，如果，没有具体的配置项，则data-id必填，且值为实际存在的data-id,表示清空这个data-id对应配置项的值；"}
  DataId *int64 `json:"data-id,omitempty" xml:"data-id,omitempty"`
  // {"en":"The url matching mode supports fuzzy regularization. If all matches, the input parameters can be configured as: *", "zh_CN":"url匹配模式，支持正则，客户入参参考：.*
  // 对于匹配到的URL进行内容重定向"}
  PathPattern *string `json:"path-pattern,omitempty" xml:"path-pattern,omitempty"`
  // {"en":"Matching conditions: specify common types, optional values are all or homepage 1. all: all files 2. homepage: home page", "zh_CN":"匹配条件：指定常用类型，可选值为all或homepage 1. all：全部文件 2. homepage：首页"}
  CustomPattern *string `json:"custom-pattern,omitempty" xml:"custom-pattern,omitempty"`
  // {"en":"directory", "zh_CN":"目录"}
  Directory *string `json:"directory,omitempty" xml:"directory,omitempty"`
  // {"en":"gif png bmp jpeg jpg html htm shtml mp3 wma flv mp4 wmv zip exe rar css txt ico js swf m3u8 xml f4m bootstarp ts", "zh_CN":"匹配条件：文件类型，多个请以英文;分隔，可选值：gif png bmp jpeg jpg html htm shtml mp3 wma flv mp4 wmv zip exe rar css txt ico js swf m3u8 xml f4m bootstarp ts"}
  FileType *string `json:"file-type,omitempty" xml:"file-type,omitempty"`
  // {"en":"Matching condition: Custom file type, please separate them by semicolon.", "zh_CN":"匹配条件：自定义文件类型，多个请以英文;分隔。"}
  CustomFileType *string `json:"custom-file-type,omitempty" xml:"custom-file-type,omitempty"`
  // {"en":"Exceptional url matching mode, except for certain URLs: such as abc.jpg, no content redirection
  // Customer reference: ^https?://[^/]+/.*\.m3u8", "zh_CN":"例外的url匹配模式，某些URL除外：如abc.jpg，不做内容重定向
  // 客户入参参考：^https?://[^/]+/.*\.m3u8"}
  ExceptPathPattern *string `json:"except-path-pattern,omitempty" xml:"except-path-pattern,omitempty"`
  // {"en":"Ignore case, the optional value is true or false, true means to ignore case; false means not to ignore case;
  // When adding a new configuration item, the default is not true.
  // If the client passes a null value: such as <ignore-letter-case></ignore-letter-case>, the configuration is cleared.", "zh_CN":"忽略大小写，可选值为true或false，true表示忽略大小写；false表示不忽略大小写；
  // 新增配置项时，不传默认为 true
  // 如果客户传了空值：如<ignore-letter-case></ignore-letter-case>，则表示清空配置"}
  IgnoreLetterCase *string `json:"ignore-letter-case,omitempty" xml:"ignore-letter-case,omitempty"`
  // {"en":"Rewrite the location where the content is generated. The input value is: Cache indicates the node;
  // Other input formats are not supported at this time", "zh_CN":"改写内容的生成位置。可输入值为：Cache表示节点；
  // 暂不支持其他入参格式"}
  PublishType *string `json:"publish-type,omitempty" xml:"publish-type,omitempty" require:"true"`
  // {"en":"Indicates the priority execution order of multiple sets of redirected content by the customer. The higher the number, the higher the priority.
  // When adding a new configuration item, the default is 10", "zh_CN":"表示客户多组重定向内容的优先执行顺序。数字越大，优先级越高。
  // 新增配置项时，不传默认为 10"}
  Priority *string `json:"priority,omitempty" xml:"priority,omitempty"`
  // {"en":"Configuration item: old url
  // Indicates the protocol mode before rewriting (that is, the object that needs to be rewritten), such as: ^https://([^/]+/.*)", "zh_CN":"配置项：旧url
  // 表示改写前的协议方式（即需要改写的对象），如：^https://([^/]+/.*)
  // 如果是回源协议改写，则表示客户请求的原始url，配套的参数after-value，表示客户请求需要转换的回源请求。"}
  BeforeValue *string `json:"before-value,omitempty" xml:"before-value,omitempty" require:"true"`
  // {"en":"Configuration item: new url
  // Indicates the protocol method after rewriting, such as: http://$1", "zh_CN":"配置项：新url
  // 表示改写后的协议方式，如：http://$1
  // 如果请求重定向带状态码则参考入参：301:https://$1
  // 注：如果url含域名，则域名需要是本身。"}
  AfterValue *string `json:"after-value,omitempty" xml:"after-value,omitempty" require:"true"`
  // {"en":"Redirection type; support for input:
  // before: before the anti-theft chain
  // after: after the anti-theft chain", "zh_CN":"重定向类型；支持入参：
  // before：防盗链之前
  // after：防盗链之后"}
  RewriteType *string `json:"rewrite-type,omitempty" xml:"rewrite-type,omitempty" require:"true"`
  // {"en":"Matching condition: Request header", "zh_CN":"匹配条件：请求头"}
  EditDomainRedirectConfigRequestHeader *string `json:"request-header,omitempty" xml:"request-header,omitempty"`
  // {"en":"Matching condition: Exception request header", "zh_CN":"匹配条件：例外的请求头"}
  ExceptionEditDomainRedirectConfigRequestHeader *string `json:"exception-request-header,omitempty" xml:"exception-request-header,omitempty"`
}

func (s EditDomainRedirectConfigRequestRewriteRuleSettings) String() string {
  return tea.Prettify(s)
}

func (s EditDomainRedirectConfigRequestRewriteRuleSettings) GoString() string {
  return s.String()
}

func (s *EditDomainRedirectConfigRequestRewriteRuleSettings) SetDataId(v int64) *EditDomainRedirectConfigRequestRewriteRuleSettings {
  s.DataId = &v
  return s
}

func (s *EditDomainRedirectConfigRequestRewriteRuleSettings) SetPathPattern(v string) *EditDomainRedirectConfigRequestRewriteRuleSettings {
  s.PathPattern = &v
  return s
}

func (s *EditDomainRedirectConfigRequestRewriteRuleSettings) SetCustomPattern(v string) *EditDomainRedirectConfigRequestRewriteRuleSettings {
  s.CustomPattern = &v
  return s
}

func (s *EditDomainRedirectConfigRequestRewriteRuleSettings) SetDirectory(v string) *EditDomainRedirectConfigRequestRewriteRuleSettings {
  s.Directory = &v
  return s
}

func (s *EditDomainRedirectConfigRequestRewriteRuleSettings) SetFileType(v string) *EditDomainRedirectConfigRequestRewriteRuleSettings {
  s.FileType = &v
  return s
}

func (s *EditDomainRedirectConfigRequestRewriteRuleSettings) SetCustomFileType(v string) *EditDomainRedirectConfigRequestRewriteRuleSettings {
  s.CustomFileType = &v
  return s
}

func (s *EditDomainRedirectConfigRequestRewriteRuleSettings) SetExceptPathPattern(v string) *EditDomainRedirectConfigRequestRewriteRuleSettings {
  s.ExceptPathPattern = &v
  return s
}

func (s *EditDomainRedirectConfigRequestRewriteRuleSettings) SetIgnoreLetterCase(v string) *EditDomainRedirectConfigRequestRewriteRuleSettings {
  s.IgnoreLetterCase = &v
  return s
}

func (s *EditDomainRedirectConfigRequestRewriteRuleSettings) SetPublishType(v string) *EditDomainRedirectConfigRequestRewriteRuleSettings {
  s.PublishType = &v
  return s
}

func (s *EditDomainRedirectConfigRequestRewriteRuleSettings) SetPriority(v string) *EditDomainRedirectConfigRequestRewriteRuleSettings {
  s.Priority = &v
  return s
}

func (s *EditDomainRedirectConfigRequestRewriteRuleSettings) SetBeforeValue(v string) *EditDomainRedirectConfigRequestRewriteRuleSettings {
  s.BeforeValue = &v
  return s
}

func (s *EditDomainRedirectConfigRequestRewriteRuleSettings) SetAfterValue(v string) *EditDomainRedirectConfigRequestRewriteRuleSettings {
  s.AfterValue = &v
  return s
}

func (s *EditDomainRedirectConfigRequestRewriteRuleSettings) SetRewriteType(v string) *EditDomainRedirectConfigRequestRewriteRuleSettings {
  s.RewriteType = &v
  return s
}

func (s *EditDomainRedirectConfigRequestRewriteRuleSettings) SetEditDomainRedirectConfigRequestHeader(v string) *EditDomainRedirectConfigRequestRewriteRuleSettings {
  s.EditDomainRedirectConfigRequestHeader = &v
  return s
}

func (s *EditDomainRedirectConfigRequestRewriteRuleSettings) SetExceptionEditDomainRedirectConfigRequestHeader(v string) *EditDomainRedirectConfigRequestRewriteRuleSettings {
  s.ExceptionEditDomainRedirectConfigRequestHeader = &v
  return s
}

type EditDomainRedirectConfigResponse struct {
  // {"en":"httpstatus=202; Indicates that the new domain API was successfully invoked, and the current deployment of the new domain can be viewed using x-cnc-request-id in the header", "zh_CN":"httpstatus=202;   表示成功调用新增域名接口，可使用header中的x-cnc-request-id查看当前新增域名的部署情况"}
  HttpStatus *int `json:"http status code,omitempty" xml:"http status code,omitempty" require:"true"`
  // {"en":"Uniquely identified id for querying tasks per request (for all API)", "zh_CN":"唯一标示的id，用于查询每次请求的任务 （适用全部接口）"}
  XCncRequestId *string `json:"x-cnc-request-id,omitempty" xml:"x-cnc-request-id,omitempty" require:"true"`
  // {"en":"The error code, when HTTPStatus is not 202, indicates the type of error the current request is calling.", "zh_CN":"错误代码，当HTTPStatus不为202时出现，表示当前请求调用的错误类型"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response information, when success is successful", "zh_CN":"响应信息，成功时为success"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s EditDomainRedirectConfigResponse) String() string {
  return tea.Prettify(s)
}

func (s EditDomainRedirectConfigResponse) GoString() string {
  return s.String()
}

func (s *EditDomainRedirectConfigResponse) SetHttpStatus(v int) *EditDomainRedirectConfigResponse {
  s.HttpStatus = &v
  return s
}

func (s *EditDomainRedirectConfigResponse) SetXCncRequestId(v string) *EditDomainRedirectConfigResponse {
  s.XCncRequestId = &v
  return s
}

func (s *EditDomainRedirectConfigResponse) SetCode(v string) *EditDomainRedirectConfigResponse {
  s.Code = &v
  return s
}

func (s *EditDomainRedirectConfigResponse) SetMessage(v string) *EditDomainRedirectConfigResponse {
  s.Message = &v
  return s
}

type EditDomainRedirectConfigPaths struct {
  // {"en":"the domain whoes need query config", "zh_CN":"需要查询配置的域名或域名id"}
  DomainName *string `json:"domain-name,omitempty" xml:"domain-name,omitempty" require:"true"`
}

func (s EditDomainRedirectConfigPaths) String() string {
  return tea.Prettify(s)
}

func (s EditDomainRedirectConfigPaths) GoString() string {
  return s.String()
}

func (s *EditDomainRedirectConfigPaths) SetDomainName(v string) *EditDomainRedirectConfigPaths {
  s.DomainName = &v
  return s
}

type EditDomainRedirectConfigParameters struct {
}

func (s EditDomainRedirectConfigParameters) String() string {
  return tea.Prettify(s)
}

func (s EditDomainRedirectConfigParameters) GoString() string {
  return s.String()
}

type EditDomainRedirectConfigRequestHeader struct {
}

func (s EditDomainRedirectConfigRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s EditDomainRedirectConfigRequestHeader) GoString() string {
  return s.String()
}

type EditDomainRedirectConfigResponseHeader struct {
}

func (s EditDomainRedirectConfigResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s EditDomainRedirectConfigResponseHeader) GoString() string {
  return s.String()
}




type UpdateHttp2SettingsConfigForWplusRequest struct {
  // {"en":"Http2.0 settings, used to enable or disable http2.0, parent node.", "zh_CN":"http2.0设置，用于设置http2.0的开启或关闭，父标签"}
  Http2Settings *UpdateHttp2SettingsConfigForWplusRequestHttp2Settings `json:"http2Settings,omitempty" xml:"http2Settings,omitempty" require:"true" type:"Struct"`
}

func (s UpdateHttp2SettingsConfigForWplusRequest) String() string {
  return tea.Prettify(s)
}

func (s UpdateHttp2SettingsConfigForWplusRequest) GoString() string {
  return s.String()
}

func (s *UpdateHttp2SettingsConfigForWplusRequest) SetHttp2Settings(v *UpdateHttp2SettingsConfigForWplusRequestHttp2Settings) *UpdateHttp2SettingsConfigForWplusRequest {
  s.Http2Settings = v
  return s
}

type UpdateHttp2SettingsConfigForWplusRequestHttp2Settings struct {
  // {"en":"Enable http2.0. The optional values are true and false. If it is empty, the default value is false. True means http2.0 is on; false means http2.0 is off.", "zh_CN":"开启http2.0，可选值为true和false，为空时默认为false。true表示开启http2.0；false表示关闭http2.0"}
  EnableHttp2 *string `json:"enableHttp2,omitempty" xml:"enableHttp2,omitempty"`
  // {"en":"Back-to-origin protocol, the optional value is
  // http1.1: Use the HTTP1.1 protocol version to back to source. if not filled, use it as default.
  // follow-request: Same as client request protocol
  // http2.0: Use the HTTP2.0 protocol. version to back to source.", "zh_CN":"回源协议，可选值为
  // http1.1：使用HTTP1.1协议版本回源，不填时默认该协议
  // follow-request：跟随客户端请求协议
  // http2.0：使用HTTP2.0协议版本回源（点播下载直播产品不支持H2回源）"}
  BackToOriginProtocol *string `json:"backToOriginProtocol,omitempty" xml:"backToOriginProtocol,omitempty"`
}

func (s UpdateHttp2SettingsConfigForWplusRequestHttp2Settings) String() string {
  return tea.Prettify(s)
}

func (s UpdateHttp2SettingsConfigForWplusRequestHttp2Settings) GoString() string {
  return s.String()
}

func (s *UpdateHttp2SettingsConfigForWplusRequestHttp2Settings) SetEnableHttp2(v string) *UpdateHttp2SettingsConfigForWplusRequestHttp2Settings {
  s.EnableHttp2 = &v
  return s
}

func (s *UpdateHttp2SettingsConfigForWplusRequestHttp2Settings) SetBackToOriginProtocol(v string) *UpdateHttp2SettingsConfigForWplusRequestHttp2Settings {
  s.BackToOriginProtocol = &v
  return s
}

type UpdateHttp2SettingsConfigForWplusResponse struct {
  // {"en":"Error code, which appears when HTTPStatus is not 202, represents the error type of the current request call", "zh_CN":"错误代码，当HTTPStatus不为202时出现，表示当前请求调用的错误类型"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response information, success when successful", "zh_CN":"响应信息，成功时为success"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s UpdateHttp2SettingsConfigForWplusResponse) String() string {
  return tea.Prettify(s)
}

func (s UpdateHttp2SettingsConfigForWplusResponse) GoString() string {
  return s.String()
}

func (s *UpdateHttp2SettingsConfigForWplusResponse) SetCode(v string) *UpdateHttp2SettingsConfigForWplusResponse {
  s.Code = &v
  return s
}

func (s *UpdateHttp2SettingsConfigForWplusResponse) SetMessage(v string) *UpdateHttp2SettingsConfigForWplusResponse {
  s.Message = &v
  return s
}

type UpdateHttp2SettingsConfigForWplusPaths struct {
  // {"en":"the domain whoes need query config", "zh_CN":"需要查询配置的域名或域名id"}
  DomainName *string `json:"domain-name,omitempty" xml:"domain-name,omitempty" require:"true"`
}

func (s UpdateHttp2SettingsConfigForWplusPaths) String() string {
  return tea.Prettify(s)
}

func (s UpdateHttp2SettingsConfigForWplusPaths) GoString() string {
  return s.String()
}

func (s *UpdateHttp2SettingsConfigForWplusPaths) SetDomainName(v string) *UpdateHttp2SettingsConfigForWplusPaths {
  s.DomainName = &v
  return s
}

type UpdateHttp2SettingsConfigForWplusParameters struct {
}

func (s UpdateHttp2SettingsConfigForWplusParameters) String() string {
  return tea.Prettify(s)
}

func (s UpdateHttp2SettingsConfigForWplusParameters) GoString() string {
  return s.String()
}

type UpdateHttp2SettingsConfigForWplusRequestHeader struct {
}

func (s UpdateHttp2SettingsConfigForWplusRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s UpdateHttp2SettingsConfigForWplusRequestHeader) GoString() string {
  return s.String()
}

type UpdateHttp2SettingsConfigForWplusResponseHeader struct {
}

func (s UpdateHttp2SettingsConfigForWplusResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s UpdateHttp2SettingsConfigForWplusResponseHeader) GoString() string {
  return s.String()
}




type QueryAmazonS3AuthorizationConfigRequest struct {
}

func (s QueryAmazonS3AuthorizationConfigRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryAmazonS3AuthorizationConfigRequest) GoString() string {
  return s.String()
}

type QueryAmazonS3AuthorizationConfigResponse struct {
  // {"en":"Domain name or domain ID that needs to be queried for configuration.", "zh_CN":"需要查询配置的域名或域名id"}
  DomainName *string `json:"domain-name,omitempty" xml:"domain-name,omitempty" require:"true"`
  // {"en":"Domain name or domain ID that needs to be queried for configuration.", "zh_CN":"需要查询配置的域名或域名id"}
  DomainId *string `json:"domain-id,omitempty" xml:"domain-id,omitempty" require:"true"`
  // {"en":"Amazon S3 Access Authorization Configuration, parent node
  // 1. When you need to configure the Amazon S3 Access Authorization rules, this must be filled in.
  // 2. Configuration of clearing for <amazon-s3-access-authorization-rules/>.
  // 3.vodstream/download support, web/wsa does not support.
  // 4.Amason S3 and Aliyun OSS cannot be configured simultaneously.", "zh_CN":"Amazon S3鉴权配置，父标签
  // 1.需要设置Amazon S3鉴权时，此项必填
  // 2.为<amazon-s3-access-authorization-rules/>时清空Amazon S3鉴权配置的配置
  // 3.点播下载支持，网页wsa不支持
  // 4.Amason S3和Aliyun OSS不可同时配置"}
  AmazonS3AccessAuthorizationRules []*QueryAmazonS3AuthorizationConfigResponseAmazonS3AccessAuthorizationRules `json:"amazon-s3-access-authorization-rules,omitempty" xml:"amazon-s3-access-authorization-rules,omitempty" require:"true" type:"Repeated"`
}

func (s QueryAmazonS3AuthorizationConfigResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryAmazonS3AuthorizationConfigResponse) GoString() string {
  return s.String()
}

func (s *QueryAmazonS3AuthorizationConfigResponse) SetDomainName(v string) *QueryAmazonS3AuthorizationConfigResponse {
  s.DomainName = &v
  return s
}

func (s *QueryAmazonS3AuthorizationConfigResponse) SetDomainId(v string) *QueryAmazonS3AuthorizationConfigResponse {
  s.DomainId = &v
  return s
}

func (s *QueryAmazonS3AuthorizationConfigResponse) SetAmazonS3AccessAuthorizationRules(v []*QueryAmazonS3AuthorizationConfigResponseAmazonS3AccessAuthorizationRules) *QueryAmazonS3AuthorizationConfigResponse {
  s.AmazonS3AccessAuthorizationRules = v
  return s
}

type QueryAmazonS3AuthorizationConfigResponseAmazonS3AccessAuthorizationRules struct     {
  // {"en":"The url matching mode. If all matches, the input parameters can be configured as: .*", "zh_CN":"url匹配模式，支持正则，如果是全部匹配，入参可以配置为：.*"}
  PathPattern *string `json:"path-pattern,omitempty" xml:"path-pattern,omitempty" require:"true"`
  // {"en":"The exception url matching mode.", "zh_CN":"例外的url匹配模式，格式同path-pattern"}
  ExceptPathPattern *string `json:"except-path-pattern,omitempty" xml:"except-path-pattern,omitempty" require:"true"`
  // {"en":"Define whether to add Authorization header when back to Amazon S3 source. Allowed  true and false.", "zh_CN":"是否添加鉴权头部，为true，则回源按照Amazon S3的算法添加 添加Authorization头部。
  // 允许值为true和false，默认为false。"}
  AddAuthorizationHeader *string `json:"add-authorization-header,omitempty" xml:"add-authorization-header,omitempty" require:"true"`
  // {"en":"access key", "zh_CN":"校验所需的密钥"}
  AccessKey *string `json:"access-key,omitempty" xml:"access-key,omitempty" require:"true"`
  // {"en":"access key id", "zh_CN":"校验所需的密钥ID"}
  AccessKeyId *string `json:"access-key-id,omitempty" xml:"access-key-id,omitempty" require:"true"`
  // {"en":"Signature version, default value is 2. Support 2 and 4.", "zh_CN":"算法版本，默认值2，目前仅支持2和4"}
  SignatureVersion *string `json:"signature-version,omitempty" xml:"signature-version,omitempty" require:"true"`
  // {"en":"Authorization region. Please refer to the AWS official site for detail.", "zh_CN":"认证地区，见AWS官网区域说明"}
  Region *string `json:"region,omitempty" xml:"region,omitempty" require:"true"`
  // {"en":"When configuring multiple configurations, the ID of a specific group of configurations", "zh_CN":"配置多组配置时，具体某组配置的id"}
  DataId *int64 `json:"data-id,omitempty" xml:"data-id,omitempty" require:"true"`
}

func (s QueryAmazonS3AuthorizationConfigResponseAmazonS3AccessAuthorizationRules) String() string {
  return tea.Prettify(s)
}

func (s QueryAmazonS3AuthorizationConfigResponseAmazonS3AccessAuthorizationRules) GoString() string {
  return s.String()
}

func (s *QueryAmazonS3AuthorizationConfigResponseAmazonS3AccessAuthorizationRules) SetPathPattern(v string) *QueryAmazonS3AuthorizationConfigResponseAmazonS3AccessAuthorizationRules {
  s.PathPattern = &v
  return s
}

func (s *QueryAmazonS3AuthorizationConfigResponseAmazonS3AccessAuthorizationRules) SetExceptPathPattern(v string) *QueryAmazonS3AuthorizationConfigResponseAmazonS3AccessAuthorizationRules {
  s.ExceptPathPattern = &v
  return s
}

func (s *QueryAmazonS3AuthorizationConfigResponseAmazonS3AccessAuthorizationRules) SetAddAuthorizationHeader(v string) *QueryAmazonS3AuthorizationConfigResponseAmazonS3AccessAuthorizationRules {
  s.AddAuthorizationHeader = &v
  return s
}

func (s *QueryAmazonS3AuthorizationConfigResponseAmazonS3AccessAuthorizationRules) SetAccessKey(v string) *QueryAmazonS3AuthorizationConfigResponseAmazonS3AccessAuthorizationRules {
  s.AccessKey = &v
  return s
}

func (s *QueryAmazonS3AuthorizationConfigResponseAmazonS3AccessAuthorizationRules) SetAccessKeyId(v string) *QueryAmazonS3AuthorizationConfigResponseAmazonS3AccessAuthorizationRules {
  s.AccessKeyId = &v
  return s
}

func (s *QueryAmazonS3AuthorizationConfigResponseAmazonS3AccessAuthorizationRules) SetSignatureVersion(v string) *QueryAmazonS3AuthorizationConfigResponseAmazonS3AccessAuthorizationRules {
  s.SignatureVersion = &v
  return s
}

func (s *QueryAmazonS3AuthorizationConfigResponseAmazonS3AccessAuthorizationRules) SetRegion(v string) *QueryAmazonS3AuthorizationConfigResponseAmazonS3AccessAuthorizationRules {
  s.Region = &v
  return s
}

func (s *QueryAmazonS3AuthorizationConfigResponseAmazonS3AccessAuthorizationRules) SetDataId(v int64) *QueryAmazonS3AuthorizationConfigResponseAmazonS3AccessAuthorizationRules {
  s.DataId = &v
  return s
}

type QueryAmazonS3AuthorizationConfigPaths struct {
  // {"en":"Domain name or domain ID that needs to be queried for configuration.", "zh_CN":"需要查询配置的域名或域名id"}
  DomainName *string `json:"domain-name,omitempty" xml:"domain-name,omitempty"`
}

func (s QueryAmazonS3AuthorizationConfigPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryAmazonS3AuthorizationConfigPaths) GoString() string {
  return s.String()
}

func (s *QueryAmazonS3AuthorizationConfigPaths) SetDomainName(v string) *QueryAmazonS3AuthorizationConfigPaths {
  s.DomainName = &v
  return s
}

type QueryAmazonS3AuthorizationConfigParameters struct {
}

func (s QueryAmazonS3AuthorizationConfigParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryAmazonS3AuthorizationConfigParameters) GoString() string {
  return s.String()
}

type QueryAmazonS3AuthorizationConfigRequestHeader struct {
}

func (s QueryAmazonS3AuthorizationConfigRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryAmazonS3AuthorizationConfigRequestHeader) GoString() string {
  return s.String()
}

type QueryAmazonS3AuthorizationConfigResponseHeader struct {
}

func (s QueryAmazonS3AuthorizationConfigResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryAmazonS3AuthorizationConfigResponseHeader) GoString() string {
  return s.String()
}




type QueryLivestreamingAntihotlinkingConfigRequest struct {
}

func (s QueryLivestreamingAntihotlinkingConfigRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryLivestreamingAntihotlinkingConfigRequest) GoString() string {
  return s.String()
}

type QueryLivestreamingAntihotlinkingConfigResponse struct {
  // {"en":"The error code, when HTTPStatus is not 200, indicates the type of error the current request is calling.", "zh_CN":"错误代码，当HTTPStatus不为200时出现，表示当前请求调用的错误类型"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response information, when success is successful", "zh_CN":"响应信息，成功时为success"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"response data", "zh_CN":"响应数据"}
  Data *QueryLivestreamingAntihotlinkingConfigResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
}

func (s QueryLivestreamingAntihotlinkingConfigResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryLivestreamingAntihotlinkingConfigResponse) GoString() string {
  return s.String()
}

func (s *QueryLivestreamingAntihotlinkingConfigResponse) SetCode(v string) *QueryLivestreamingAntihotlinkingConfigResponse {
  s.Code = &v
  return s
}

func (s *QueryLivestreamingAntihotlinkingConfigResponse) SetMessage(v string) *QueryLivestreamingAntihotlinkingConfigResponse {
  s.Message = &v
  return s
}

func (s *QueryLivestreamingAntihotlinkingConfigResponse) SetData(v *QueryLivestreamingAntihotlinkingConfigResponseData) *QueryLivestreamingAntihotlinkingConfigResponse {
  s.Data = v
  return s
}

type QueryLivestreamingAntihotlinkingConfigResponseData struct {
  // {"en":"domain id", "zh_CN":"域名ID"}
  DomainId *string `json:"domainId,omitempty" xml:"domainId,omitempty" require:"true"`
  // {"en":"domain name", "zh_CN":"域名名称"}
  DomainName *string `json:"domainName,omitempty" xml:"domainName,omitempty" require:"true"`
  // {"en":"Streaming visit control configuration, parent node
  // 1. This item is required when the function of streaming visit control configuration needs to be set
  // 2. Clear the Streaming visit control configuration when <visitControlRules/>", "zh_CN":"流媒体防盗链配置，父标签
  // 1.需要设置流媒体防盗链配置时，此项必填
  // 2.为<visitControlRules/>时清空流媒体防盗链配置"}
  VisitControlRules []*QueryLivestreamingAntihotlinkingConfigResponseDataVisitControlRules `json:"visitControlRules,omitempty" xml:"visitControlRules,omitempty" require:"true" type:"Repeated"`
}

func (s QueryLivestreamingAntihotlinkingConfigResponseData) String() string {
  return tea.Prettify(s)
}

func (s QueryLivestreamingAntihotlinkingConfigResponseData) GoString() string {
  return s.String()
}

func (s *QueryLivestreamingAntihotlinkingConfigResponseData) SetDomainId(v string) *QueryLivestreamingAntihotlinkingConfigResponseData {
  s.DomainId = &v
  return s
}

func (s *QueryLivestreamingAntihotlinkingConfigResponseData) SetDomainName(v string) *QueryLivestreamingAntihotlinkingConfigResponseData {
  s.DomainName = &v
  return s
}

func (s *QueryLivestreamingAntihotlinkingConfigResponseData) SetVisitControlRules(v []*QueryLivestreamingAntihotlinkingConfigResponseDataVisitControlRules) *QueryLivestreamingAntihotlinkingConfigResponseData {
  s.VisitControlRules = v
  return s
}

type QueryLivestreamingAntihotlinkingConfigResponseDataVisitControlRules struct     {
  // {"en":"Control action, optional values: allow, forbid
  // Note: when add or modify a rule, you must configure both controlaction and IP, referer. At least one of IP and referer must be configured, otherwise the function will not work.", "zh_CN":"控制动作，允许或禁止，可选值：allow，forbid
  // 注意：配置或修改流媒体防盗链时，必须同时配置controlAction和ip、referer，ip和referer至少配置一项，否则此防盗链功能无效。"}
  ControlAction *string `json:"controlAction,omitempty" xml:"controlAction,omitempty" require:"true"`
  // {"en":"Allowed or forbidden IP, support IP and IP segment, parent node. Controlaction must be configured at the same time. Such as:
  // <controlAction>forbid</controlAction>
  // <ips>
  // <ip>1.1.1.1</ip>
  // <ip>2.2.2.0/24</ip>
  // </ips>", "zh_CN":"允许或禁止的IP，支持IP或IP段，父标签。必须同时配置controlAction。
  // 格式如
  // <controlAction>forbid</controlAction>
  // <ips>
  // <ip>1.1.1.1</ip> 
  // <ip>2.2.2.0/24</ip>
  // </ips>"}
  Ips []*string `json:"ips,omitempty" xml:"ips,omitempty" require:"true" type:"Repeated"`
  // {"en":"The allowed or forbidden referer, support domain name and URL format. Parent node. It must start with the protocol header, such as http://, https://. Controlaction must be configured at the same time. Such as:
  // <controlAction>forbid</controlAction>
  // <referers>
  // <referer>http://www.referer1.com</referer>
  // <referer>http://www.referer2.com</referer>
  // </referers>", "zh_CN":"允许或禁止的referer，支持输入域名或url格式，父标签，必须以协议头开头，如http://、https://。必须同时配置controlAction。格式如：
  // <controlAction>forbid</controlAction>
  // <referers>
  // <referer>http://www.referer1.com</referer>
  // <referer>http://www.referer2.com</referer>
  // </referers>"}
  Referers []*string `json:"referers,omitempty" xml:"referers,omitempty" require:"true" type:"Repeated"`
  // {"en":"Allow or forbid null referer. Optional value is 'true' or 'false',default value is 'false'.
  // When controlAction equals forbid, it means whether null referer is forbidden. If it is true, it means null referer is prohibited.
  // When controlAction  equals  allow, it means whether null referer is allowed. If it is true, it means null referer is allowed.", "zh_CN":"是否允许/禁止空referer，可选值为true、false，默认为否。
  // 当控制动作=forbid，表示是否禁止空referer，为true，表示禁止空referer
  // 当控制动作=allow，表示是否允许空referer，为true，表示允许空referer"}
  AllowNullReferer *bool `json:"allowNullReferer,omitempty" xml:"allowNullReferer,omitempty" require:"true"`
  // {"en":"IP and refer condition relationship, optional values: and, or, default to and.", "zh_CN":"ip和refer条件关系，可选值：and，or，默认为and
  // 为and，表示ip和refer都满足，才允许/禁止。为or，表示ip和refer满足一个，就允许/禁止"}
  ControlRelation *string `json:"controlRelation,omitempty" xml:"controlRelation,omitempty" require:"true"`
  // {"en":"Data-id is to indicate a specific group configuration when the client has multiple groups of configurations. Data-id can be retrieved through a query interface. Note: A. If data-id is passed, it means that one group of configuration items is specified to be modified, and no other group configuration items need to be modified. B. If multiple groups of configurations are included, some of them are configured with data-id and others are not, then the expression of data-id is used to modify a specific group of configurations, and a new group of configurations is added on the original basis without the expression of data-id. C. If the data-id is not transmitted, it means that the original configuration will be fully covered by this configuration. D. If no configuration parameter is passed, only domain name and secondary label are passed, which means that all configuration of domain name secondary service corresponding to this interface is cleared. E. If there is no specific configuration item in a set of configurations, the data-id must be filled in, and the value is the actual data-id, which means clearing the value of the corresponding data-id configuration item; it is not allowed that there is no specific configuration item or data-id in a set of configurations.", "zh_CN":"配置多组配置时，具体某组配置的id。dataId可以通过查询接口获取。 注意： a、如果有传dataId，说明指定修改其中一组配置项内容，不需求修改其他组配置内容不需要入参； b、如果入参多组配置，其中有些组配置有传dataId，有些没有传，则有传dataId的表示修改具体某组配置，没有传dataId的表示在原来基础上新增一组配置； c、如果入参都没有传dataId,表示用本次的配置全量覆盖原先配置； d、如果入参没有传任何配置项参数，只传了域名和二级标签，表示清空这个接口对应域名二级服务所有配置； e、如果一组配置没有具体的配置项，则dataId必填，且值为实际存在的dataId，表示清空这个dataId对应配置项的值；不允许一组配置没有具体的配置项也没有dataId。"}
  DataId *int `json:"dataId,omitempty" xml:"dataId,omitempty" require:"true"`
}

func (s QueryLivestreamingAntihotlinkingConfigResponseDataVisitControlRules) String() string {
  return tea.Prettify(s)
}

func (s QueryLivestreamingAntihotlinkingConfigResponseDataVisitControlRules) GoString() string {
  return s.String()
}

func (s *QueryLivestreamingAntihotlinkingConfigResponseDataVisitControlRules) SetControlAction(v string) *QueryLivestreamingAntihotlinkingConfigResponseDataVisitControlRules {
  s.ControlAction = &v
  return s
}

func (s *QueryLivestreamingAntihotlinkingConfigResponseDataVisitControlRules) SetIps(v []*string) *QueryLivestreamingAntihotlinkingConfigResponseDataVisitControlRules {
  s.Ips = v
  return s
}

func (s *QueryLivestreamingAntihotlinkingConfigResponseDataVisitControlRules) SetReferers(v []*string) *QueryLivestreamingAntihotlinkingConfigResponseDataVisitControlRules {
  s.Referers = v
  return s
}

func (s *QueryLivestreamingAntihotlinkingConfigResponseDataVisitControlRules) SetAllowNullReferer(v bool) *QueryLivestreamingAntihotlinkingConfigResponseDataVisitControlRules {
  s.AllowNullReferer = &v
  return s
}

func (s *QueryLivestreamingAntihotlinkingConfigResponseDataVisitControlRules) SetControlRelation(v string) *QueryLivestreamingAntihotlinkingConfigResponseDataVisitControlRules {
  s.ControlRelation = &v
  return s
}

func (s *QueryLivestreamingAntihotlinkingConfigResponseDataVisitControlRules) SetDataId(v int) *QueryLivestreamingAntihotlinkingConfigResponseDataVisitControlRules {
  s.DataId = &v
  return s
}

type QueryLivestreamingAntihotlinkingConfigPaths struct {
  // {"en":"Domain name or domain name id to query configuration", "zh_CN":"需要查询配置的域名（domainName）或域名id（domainId）"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
}

func (s QueryLivestreamingAntihotlinkingConfigPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryLivestreamingAntihotlinkingConfigPaths) GoString() string {
  return s.String()
}

func (s *QueryLivestreamingAntihotlinkingConfigPaths) SetDomain(v string) *QueryLivestreamingAntihotlinkingConfigPaths {
  s.Domain = &v
  return s
}

type QueryLivestreamingAntihotlinkingConfigParameters struct {
}

func (s QueryLivestreamingAntihotlinkingConfigParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryLivestreamingAntihotlinkingConfigParameters) GoString() string {
  return s.String()
}

type QueryLivestreamingAntihotlinkingConfigRequestHeader struct {
}

func (s QueryLivestreamingAntihotlinkingConfigRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryLivestreamingAntihotlinkingConfigRequestHeader) GoString() string {
  return s.String()
}

type QueryLivestreamingAntihotlinkingConfigResponseHeader struct {
}

func (s QueryLivestreamingAntihotlinkingConfigResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryLivestreamingAntihotlinkingConfigResponseHeader) GoString() string {
  return s.String()
}




type QueryAntiHotlinkingConfigRequest struct {
}

func (s QueryAntiHotlinkingConfigRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryAntiHotlinkingConfigRequest) GoString() string {
  return s.String()
}

type QueryAntiHotlinkingConfigResponse struct {
  // {"en":"the domain whoes need query config", "zh_CN":"需要查询配置的域名"}
  DomainName *string `json:"domain-name,omitempty" xml:"domain-name,omitempty" require:"true"`
  // {"en":"the domain whoes need query config", "zh_CN":"需要查询配置的域名id"}
  DomainId *string `json:"domain-id,omitempty" xml:"domain-id,omitempty" require:"true"`
  // {"en":"Anti-theft chain configuration
  // note:
  // 1. When you need to cancel the anti-theft chain configuration settings, you can pass in the empty node <cache-time-behaviors></cache-time-behaviors>.
  // 2. When it is necessary to set the anti-theft chain configuration, this item is required.", "zh_CN":"防盗链配置
  // 注意：
  // 1、需要取消防盗链配置设置时，可以传入空节点<cache-time-behaviors></cache-time-behaviors>。
  // 2、表示需要设置防盗链配置时，此项必填"}
  VisitControlRules []*QueryAntiHotlinkingConfigResponseVisitControlRules `json:"visit-control-rules,omitempty" xml:"visit-control-rules,omitempty" require:"true" type:"Repeated"`
}

func (s QueryAntiHotlinkingConfigResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryAntiHotlinkingConfigResponse) GoString() string {
  return s.String()
}

func (s *QueryAntiHotlinkingConfigResponse) SetDomainName(v string) *QueryAntiHotlinkingConfigResponse {
  s.DomainName = &v
  return s
}

func (s *QueryAntiHotlinkingConfigResponse) SetDomainId(v string) *QueryAntiHotlinkingConfigResponse {
  s.DomainId = &v
  return s
}

func (s *QueryAntiHotlinkingConfigResponse) SetVisitControlRules(v []*QueryAntiHotlinkingConfigResponseVisitControlRules) *QueryAntiHotlinkingConfigResponse {
  s.VisitControlRules = v
  return s
}

type QueryAntiHotlinkingConfigResponseVisitControlRules struct     {
  // {"en":"Add a grid type identifier to indicate a specific group configuration when the client has multiple groups of configurations.", "zh_CN":"添加grid类型标识，表示客户多组配置时，具体某组配置"}
  DataId *int32 `json:"data-id,omitempty" xml:"data-id,omitempty" require:"true"`
  // {"en":"The url matching mode supports regularization. If all matches, the input parameters can be configured as: .*", "zh_CN":"url匹配模式，支持正则，如果是全部匹配，入参可以配置为：.*"}
  PathPattern *string `json:"path-pattern,omitempty" xml:"path-pattern,omitempty" require:"true"`
  // {"en":"Exceptional url matching mode, except for some URLs: such as abc.jpg, do not do anti-theft chain function
  // E.g: ^https?://[^/]+/.*\.m3u8", "zh_CN":"例外的url匹配模式，某些URL除外：如abc.jpg，不做防盗链功能
  // 客户入参参考：^https?://[^/]+/.*\.m3u8"}
  ExceptPathPattern *string `json:"except-path-pattern,omitempty" xml:"except-path-pattern,omitempty" require:"true"`
  // {"en":"Specify common types: Select the domain name that requires the anti-theft chain to be all files or the home page. :
  // E.g:
  // All: all files
  // Homepage: homepage", "zh_CN":"指定常用类型：选择需要防盗链的域名是全部文件还是首页。入参参考值：
  // all：全部文件
  // homepage：首页"}
  CustomPattern *string `json:"custom-pattern,omitempty" xml:"custom-pattern,omitempty" require:"true"`
  // {"en":"File Type: Specify the file type for anti-theft chain settings.
  // File types include: gif png bmp jpeg jpg html htm shtml mp3 wma flv mp4 wmv zip exe rar css txt ico js swf
  // If you need all types, pass all directly. Multiples are separated by semicolons, and all and specific file types cannot be configured at the same time.", "zh_CN":"文件类型：指定文件类型进行防盗链设置。
  // 文件类型包括：gif png bmp jpeg jpg html htm shtml mp3 wma flv mp4 wmv zip exe rar css txt ico js swf
  // 如果需要全部类型，则直接传all。多个以分号隔开，all和具体文件类型不能同时配置。"}
  FileType *string `json:"file-type,omitempty" xml:"file-type,omitempty" require:"true"`
  // {"en":"Custom file type: Fill in the appropriate identifiable file type according to your needs outside of the specified file type. Can be used with file-type. If the file-type is also configured, the actual file type is the sum of the two parameters.", "zh_CN":"自定义文件类型：在指定文件类型外根据自身需求，填写适当的可识别文件类型。可以搭配file-type使用。如果file-type也有配置，实际生效的文件类型是两个入参的总和"}
  CustomFileType *string `json:"custom-file-type,omitempty" xml:"custom-file-type,omitempty" require:"true"`
  // {"en":"Specify URL cache: Specify url according to requirements for anti-theft chain setting
  // INS format does not support URI format with http(s)://", "zh_CN":"指定URL缓存：根据需求指定url进行防盗链设置
  // 入参不支持含http(s):// 开头的URI格式"}
  SpecifyUrlPattern *string `json:"specify-url-pattern,omitempty" xml:"specify-url-pattern,omitempty" require:"true"`
  // {"en":"Directory: Specify the directory for anti-theft chain settings
  // Enter a legal directory format. Multiple separated by semicolons", "zh_CN":"目录：指定目录进行防盗链设置
  // 输入合法的目录格式。多个以英文分号隔开"}
  Directory *string `json:"directory,omitempty" xml:"directory,omitempty" require:"true"`
  // {"en":"Exception file type: Specify the file type that does not require anti-theft chain function
  // File types include: gif png bmp jpeg jpg html htm shtml mp3 wma flv mp4 wmv zip exe rar css txt ico js swf
  // If you need all types, pass all directly. Multiple separated by semicolons, all and specific file types cannot be configured at the same time
  // If file-type=all, except-file-type=all means that the task file type is not matched.", "zh_CN":"例外的文件类型：指定不需要进行防盗链功能的文件类型
  // 文件类型包括：gif png bmp jpeg jpg html htm shtml mp3 wma flv mp4 wmv zip exe rar css txt ico js swf
  // 如果需要全部类型，则直接传all。多个以分号隔开，all和具体文件类型不能同时配置
  // 如果file-type=all,except-file-type=all 则表示不匹配任务文件类型"}
  ExceptFileType *string `json:"except-file-type,omitempty" xml:"except-file-type,omitempty" require:"true"`
  // {"en":"Exceptional custom file types: Fill in the appropriate identifiable file types based on your needs, outside of the specified file type. Can be used with the exception-file-type. If the except-file-type is also configured, the actual file type is the sum of the two parameters.", "zh_CN":"例外的自定义文件类型：在指定文件类型外根据自身需求，填写适当的可识别文件类型。可以搭配except-file-type使用。如果except-file-type也有配置，实际生效的文件类型是两个入参的总和"}
  ExceptCustomFileType *string `json:"except-custom-file-type,omitempty" xml:"except-custom-file-type,omitempty" require:"true"`
  // {"en":"Exceptional directory: Specify a directory that does not require anti-theft chain settings
  // Enter a legal directory format. Multiple separated by semicolons", "zh_CN":"例外的目录：指定不需要进行进行防盗链设置的目录
  // 输入合法的目录格式。多个以英文分号隔开"}
  ExceptDirectory *string `json:"except-directory,omitempty" xml:"except-directory,omitempty" require:"true"`
  // {"en":"control direction. Available values: 403 and 302
  // 1) 403 means to return a specific error status code to reject the service (the default mode, the status code can be specified, generally 403).
  // 2) 302 means to return 302 the redirect url of the Found, the redirected url can be specified. If pass 302, rewrite-to is required", "zh_CN":"控制方向。可选值：403和302
  // 1） 403表示返回特定的错误状态码来拒绝服务（默认方式，状态码可以指定，一般为403）。
  // 2） 302表示返回302 Found的重定向url，重定向的url可以指定。如果传302，rewrite-to必填"}
  ControlAction *string `json:"control-action,omitempty" xml:"control-action,omitempty" require:"true"`
  // {"en":"Specify the url after the 302 jump. This field is required if the control-action value is 302.", "zh_CN":"指定302跳转后的url。如果control-action值为302，此项必填"}
  RewriteTo *string `json:"rewrite-to,omitempty" xml:"rewrite-to,omitempty" require:"true"`
  // {"en":"Indicates the priority execution order of multiple sets of redirected content by the customer. The higher the number, the higher the priority.
  // When adding a new configuration item, the default is 10", "zh_CN":"表示客户多组重定向内容的优先执行顺序。数字越大，优先级越高。
  // 新增配置项时，不传默认为 10"}
  Priority *int32 `json:"priority,omitempty" xml:"priority,omitempty" require:"true"`
  // {"en":"Exceptional Request Method.Multiple entires separated by ;", "zh_CN":"例外的请求方法。多个以;隔开"}
  ExceptionalRequest *string `json:"exceptional-request,omitempty" xml:"exceptional-request,omitempty" require:"true"`
  // {"en":"Only true and false are allowed to be filled in ignoredcase.", "zh_CN":"是否忽略大小写。只允许填写true或false"}
  IgnoredCase *string `json:"ignored-case,omitempty" xml:"ignored-case,omitempty" require:"true"`
  // {"en":"You can set other visit control rules here.", "zh_CN":"配置其他访问控制策略，比如禁止访问的时间"}
  AdvanceControlRules *QueryAntiHotlinkingConfigResponseVisitControlRulesAdvanceControlRules `json:"advance-control-rules,omitempty" xml:"advance-control-rules,omitempty" require:"true" type:"Struct"`
  // {"en":"Identify IP black and white list anti-theft chain note: 1, a set of black and white list anti-theft chain, only one set under a data-id 2. When the air interface label indicates the exception of the IP segment configuration and the forbidden IP segment configuration.", "zh_CN":"标识IP黑白名单防盗链 注意： 1、表示一组黑白名单防盗链，一个data-id下只能一组 2、当传空标签表示清楚例外的IP段配置和禁止的IP段配置。"}
  IpControlRule *QueryAntiHotlinkingConfigResponseVisitControlRulesIpControlRule `json:"ip-control-rule,omitempty" xml:"ip-control-rule,omitempty" require:"true" type:"Struct"`
  // {"en":"The exception IP segment supports input IP or IP segment, and the IP segments are separated by a semicolon (;), such as 1.1.1.0/24; 2.2.2.2, some IP exceptions, no anti-theft chain", "zh_CN":"例外的IP段，支持输入IP或IP段，IP段之间用分号(;)隔开，如1.1.1.0/24;2.2.2.2，某些IP例外，不做防盗链"}
  AllowedIps *string `json:"allowed-ips,omitempty" xml:"allowed-ips,omitempty" require:"true"`
  // {"en":"Identify referer anti-theft chain
  // Note:
  // 1. Represents a set of referer security chains, and a single data-id can only have one set under one
  // 2. when the empty tag means to clear referer security chain
  // 3. legal refer, ( legal domain name, legal URL), illegal refer, ( illegal domain name, illegal URL) these four, a data-id can only configure one or all empty under one data-id", "zh_CN":"标识referer防盗链
  // 注意：
  // 1、表示一组referer防盗链，一个data-id下只能一组
  // 2、当传空标签表示清除referer防盗链
  // 3、合法refer、（合法域名、合法URL）、非法refer、（非法域名、非法URL）这四项，一个data-id下只能配置一个或者都为空"}
  RefererControlRule *QueryAntiHotlinkingConfigResponseVisitControlRulesRefererControlRule `json:"referer-control-rule,omitempty" xml:"referer-control-rule,omitempty" require:"true" type:"Struct"`
  // {"en":"Configuration cookie control rules.Allow-cookie and forbidden-cookie are not allowed to be configured together.", "zh_CN":"配置Cookie防盗链策略。【允许的cookie】和【禁止的cookie】只允许配置一个"}
  CookieControlRules *QueryAntiHotlinkingConfigResponseVisitControlRulesCookieControlRules `json:"cookie-control-rules,omitempty" xml:"cookie-control-rules,omitempty" require:"true" type:"Struct"`
  // {"en":"Configuration custom header control rules.Header-whitelist and header-blacklist are not allowed to be configured together.", "zh_CN":"配置自定义头部防盗链。【头域黑名单】和【头域白名单】只允许配置一个"}
  CustomHeaderControlRules *QueryAntiHotlinkingConfigResponseVisitControlRulesCustomHeaderControlRules `json:"custom-header-control-rules,omitempty" xml:"custom-header-control-rules,omitempty" require:"true" type:"Struct"`
}

func (s QueryAntiHotlinkingConfigResponseVisitControlRules) String() string {
  return tea.Prettify(s)
}

func (s QueryAntiHotlinkingConfigResponseVisitControlRules) GoString() string {
  return s.String()
}

func (s *QueryAntiHotlinkingConfigResponseVisitControlRules) SetDataId(v int32) *QueryAntiHotlinkingConfigResponseVisitControlRules {
  s.DataId = &v
  return s
}

func (s *QueryAntiHotlinkingConfigResponseVisitControlRules) SetPathPattern(v string) *QueryAntiHotlinkingConfigResponseVisitControlRules {
  s.PathPattern = &v
  return s
}

func (s *QueryAntiHotlinkingConfigResponseVisitControlRules) SetExceptPathPattern(v string) *QueryAntiHotlinkingConfigResponseVisitControlRules {
  s.ExceptPathPattern = &v
  return s
}

func (s *QueryAntiHotlinkingConfigResponseVisitControlRules) SetCustomPattern(v string) *QueryAntiHotlinkingConfigResponseVisitControlRules {
  s.CustomPattern = &v
  return s
}

func (s *QueryAntiHotlinkingConfigResponseVisitControlRules) SetFileType(v string) *QueryAntiHotlinkingConfigResponseVisitControlRules {
  s.FileType = &v
  return s
}

func (s *QueryAntiHotlinkingConfigResponseVisitControlRules) SetCustomFileType(v string) *QueryAntiHotlinkingConfigResponseVisitControlRules {
  s.CustomFileType = &v
  return s
}

func (s *QueryAntiHotlinkingConfigResponseVisitControlRules) SetSpecifyUrlPattern(v string) *QueryAntiHotlinkingConfigResponseVisitControlRules {
  s.SpecifyUrlPattern = &v
  return s
}

func (s *QueryAntiHotlinkingConfigResponseVisitControlRules) SetDirectory(v string) *QueryAntiHotlinkingConfigResponseVisitControlRules {
  s.Directory = &v
  return s
}

func (s *QueryAntiHotlinkingConfigResponseVisitControlRules) SetExceptFileType(v string) *QueryAntiHotlinkingConfigResponseVisitControlRules {
  s.ExceptFileType = &v
  return s
}

func (s *QueryAntiHotlinkingConfigResponseVisitControlRules) SetExceptCustomFileType(v string) *QueryAntiHotlinkingConfigResponseVisitControlRules {
  s.ExceptCustomFileType = &v
  return s
}

func (s *QueryAntiHotlinkingConfigResponseVisitControlRules) SetExceptDirectory(v string) *QueryAntiHotlinkingConfigResponseVisitControlRules {
  s.ExceptDirectory = &v
  return s
}

func (s *QueryAntiHotlinkingConfigResponseVisitControlRules) SetControlAction(v string) *QueryAntiHotlinkingConfigResponseVisitControlRules {
  s.ControlAction = &v
  return s
}

func (s *QueryAntiHotlinkingConfigResponseVisitControlRules) SetRewriteTo(v string) *QueryAntiHotlinkingConfigResponseVisitControlRules {
  s.RewriteTo = &v
  return s
}

func (s *QueryAntiHotlinkingConfigResponseVisitControlRules) SetPriority(v int32) *QueryAntiHotlinkingConfigResponseVisitControlRules {
  s.Priority = &v
  return s
}

func (s *QueryAntiHotlinkingConfigResponseVisitControlRules) SetExceptionalRequest(v string) *QueryAntiHotlinkingConfigResponseVisitControlRules {
  s.ExceptionalRequest = &v
  return s
}

func (s *QueryAntiHotlinkingConfigResponseVisitControlRules) SetIgnoredCase(v string) *QueryAntiHotlinkingConfigResponseVisitControlRules {
  s.IgnoredCase = &v
  return s
}

func (s *QueryAntiHotlinkingConfigResponseVisitControlRules) SetAdvanceControlRules(v *QueryAntiHotlinkingConfigResponseVisitControlRulesAdvanceControlRules) *QueryAntiHotlinkingConfigResponseVisitControlRules {
  s.AdvanceControlRules = v
  return s
}

func (s *QueryAntiHotlinkingConfigResponseVisitControlRules) SetIpControlRule(v *QueryAntiHotlinkingConfigResponseVisitControlRulesIpControlRule) *QueryAntiHotlinkingConfigResponseVisitControlRules {
  s.IpControlRule = v
  return s
}

func (s *QueryAntiHotlinkingConfigResponseVisitControlRules) SetAllowedIps(v string) *QueryAntiHotlinkingConfigResponseVisitControlRules {
  s.AllowedIps = &v
  return s
}

func (s *QueryAntiHotlinkingConfigResponseVisitControlRules) SetRefererControlRule(v *QueryAntiHotlinkingConfigResponseVisitControlRulesRefererControlRule) *QueryAntiHotlinkingConfigResponseVisitControlRules {
  s.RefererControlRule = v
  return s
}

func (s *QueryAntiHotlinkingConfigResponseVisitControlRules) SetCookieControlRules(v *QueryAntiHotlinkingConfigResponseVisitControlRulesCookieControlRules) *QueryAntiHotlinkingConfigResponseVisitControlRules {
  s.CookieControlRules = v
  return s
}

func (s *QueryAntiHotlinkingConfigResponseVisitControlRules) SetCustomHeaderControlRules(v *QueryAntiHotlinkingConfigResponseVisitControlRulesCustomHeaderControlRules) *QueryAntiHotlinkingConfigResponseVisitControlRules {
  s.CustomHeaderControlRules = v
  return s
}

type QueryAntiHotlinkingConfigResponseVisitControlRulesAdvanceControlRules struct {
  // {"en":"The format is: YYYY start date - end date weekday start time - end date, Beijing time. For example: 2022 1/1-6/30 Mon 00:00-16:00, which means: 0:00-16:00 every Monday from January 1st to June 30th, 2017. You can only configure the start time, such as 00:00-16:00, which means 0:00 to 16:00 every day.", "zh_CN":"禁止访问的时间"}
  InvalidTime *string `json:"invalid-time,omitempty" xml:"invalid-time,omitempty" require:"true"`
}

func (s QueryAntiHotlinkingConfigResponseVisitControlRulesAdvanceControlRules) String() string {
  return tea.Prettify(s)
}

func (s QueryAntiHotlinkingConfigResponseVisitControlRulesAdvanceControlRules) GoString() string {
  return s.String()
}

func (s *QueryAntiHotlinkingConfigResponseVisitControlRulesAdvanceControlRules) SetInvalidTime(v string) *QueryAntiHotlinkingConfigResponseVisitControlRulesAdvanceControlRules {
  s.InvalidTime = &v
  return s
}

type QueryAntiHotlinkingConfigResponseVisitControlRulesIpControlRule struct {
  // {"en":"Prohibited IP segment
  // Input parameter limit reference interface limit
  // Forbidden IP and exceptional IP cannot be configured at the same time", "zh_CN":"禁止的IP段
  // 禁止的IP和例外的IP只能一个有值"}
  ForbiddenIps *string `json:"forbidden-ips,omitempty" xml:"forbidden-ips,omitempty" require:"true"`
}

func (s QueryAntiHotlinkingConfigResponseVisitControlRulesIpControlRule) String() string {
  return tea.Prettify(s)
}

func (s QueryAntiHotlinkingConfigResponseVisitControlRulesIpControlRule) GoString() string {
  return s.String()
}

func (s *QueryAntiHotlinkingConfigResponseVisitControlRulesIpControlRule) SetForbiddenIps(v string) *QueryAntiHotlinkingConfigResponseVisitControlRulesIpControlRule {
  s.ForbiddenIps = &v
  return s
}

type QueryAntiHotlinkingConfigResponseVisitControlRulesRefererControlRule struct {
  // {"en":"If any of the four terms 'nullreferer: legal referer, (legal domain name, legal URL), illegal referer, (illegal domain name, illegal URL)' is allowed, then 'nullreferer' cannot be null.If the four terms 'legal refer', 'legal domain name, legal URL', 'illegal refer', 'illegal domain name, illegal URL' are all null values, then 'whether to allow a null referer' must be null", "zh_CN":"是否允许空referer：合法refer、（合法域名、合法URL）、非法refer、（非法域名、非法URL）这四项任意一项有值，则&ldquo;是否允许空referer&rdquo;不能为空；合法refer、（合法域名、合法URL）、非法refer、（非法域名、非法URL）这四项都为空值，则&ldquo;是否允许空referer&rdquo;必须为空"}
  AllowNullReferer *string `json:"allow-null-referer,omitempty" xml:"allow-null-referer,omitempty" require:"true"`
  // {"en":"Legal referer.", "zh_CN":"合法referer.可以输入url或域名"}
  ValidReferer *string `json:"valid-referer,omitempty" xml:"valid-referer,omitempty" require:"true"`
  // {"en":"Legal url, enter the correct url format", "zh_CN":"合法url，输入正确的url格式"}
  ValidUrl *string `json:"valid-url,omitempty" xml:"valid-url,omitempty" require:"true"`
  // {"en":"Legal domain name", "zh_CN":"合法域名"}
  ValidDomain *string `json:"valid-domain,omitempty" xml:"valid-domain,omitempty" require:"true"`
  // {"en":"Illegal referer", "zh_CN":"非法referer，可以输入url或域名"}
  InvalidReferer *string `json:"invalid-referer,omitempty" xml:"invalid-referer,omitempty" require:"true"`
  // {"en":"Invalid url, enter the correct url format", "zh_CN":"非法url，输入正确的url格式"}
  InvalidUrl *string `json:"invalid-url,omitempty" xml:"invalid-url,omitempty" require:"true"`
  // {"en":"Illegal domain name", "zh_CN":"非法域名"}
  InvalidDomain *string `json:"invalid-domain,omitempty" xml:"invalid-domain,omitempty" require:"true"`
  // {"en":"UA head protection against hotlinking,
  // Note:
  // 1. Represents a group of UA head defense hotlinking, and only one group under a data-id
  // 2. when empty label means clear UA head protection hotlinking", "zh_CN":"标识UA头防盗链，
  // 注意：
  // 1、表示一组UA头防盗链，一个data-id下只能一组
  // 2、当传空标签表示清除UA头防盗链"}
  UaControlRule *QueryAntiHotlinkingConfigResponseVisitControlRulesRefererControlRuleUaControlRule `json:"ua-control-rule,omitempty" xml:"ua-control-rule,omitempty" require:"true" type:"Struct"`
  // {"en":"Configure other access control rules, such as invalid visitor regions, example:
  // advance-control-rules:{invalid-visitor-region:CN;JP;K}", "zh_CN":"配置其他访问控制策略，比如禁止的访客区域，JSON示例：
  // advance-control-rules:{invalid-visitor-region:CN;JP;KR}"}
  AdvanceControlRules *QueryAntiHotlinkingConfigResponseVisitControlRulesRefererControlRuleAdvanceControlRules `json:"advance-control-rules,omitempty" xml:"advance-control-rules,omitempty" require:"true" type:"Struct"`
}

func (s QueryAntiHotlinkingConfigResponseVisitControlRulesRefererControlRule) String() string {
  return tea.Prettify(s)
}

func (s QueryAntiHotlinkingConfigResponseVisitControlRulesRefererControlRule) GoString() string {
  return s.String()
}

func (s *QueryAntiHotlinkingConfigResponseVisitControlRulesRefererControlRule) SetAllowNullReferer(v string) *QueryAntiHotlinkingConfigResponseVisitControlRulesRefererControlRule {
  s.AllowNullReferer = &v
  return s
}

func (s *QueryAntiHotlinkingConfigResponseVisitControlRulesRefererControlRule) SetValidReferer(v string) *QueryAntiHotlinkingConfigResponseVisitControlRulesRefererControlRule {
  s.ValidReferer = &v
  return s
}

func (s *QueryAntiHotlinkingConfigResponseVisitControlRulesRefererControlRule) SetValidUrl(v string) *QueryAntiHotlinkingConfigResponseVisitControlRulesRefererControlRule {
  s.ValidUrl = &v
  return s
}

func (s *QueryAntiHotlinkingConfigResponseVisitControlRulesRefererControlRule) SetValidDomain(v string) *QueryAntiHotlinkingConfigResponseVisitControlRulesRefererControlRule {
  s.ValidDomain = &v
  return s
}

func (s *QueryAntiHotlinkingConfigResponseVisitControlRulesRefererControlRule) SetInvalidReferer(v string) *QueryAntiHotlinkingConfigResponseVisitControlRulesRefererControlRule {
  s.InvalidReferer = &v
  return s
}

func (s *QueryAntiHotlinkingConfigResponseVisitControlRulesRefererControlRule) SetInvalidUrl(v string) *QueryAntiHotlinkingConfigResponseVisitControlRulesRefererControlRule {
  s.InvalidUrl = &v
  return s
}

func (s *QueryAntiHotlinkingConfigResponseVisitControlRulesRefererControlRule) SetInvalidDomain(v string) *QueryAntiHotlinkingConfigResponseVisitControlRulesRefererControlRule {
  s.InvalidDomain = &v
  return s
}

func (s *QueryAntiHotlinkingConfigResponseVisitControlRulesRefererControlRule) SetUaControlRule(v *QueryAntiHotlinkingConfigResponseVisitControlRulesRefererControlRuleUaControlRule) *QueryAntiHotlinkingConfigResponseVisitControlRulesRefererControlRule {
  s.UaControlRule = v
  return s
}

func (s *QueryAntiHotlinkingConfigResponseVisitControlRulesRefererControlRule) SetAdvanceControlRules(v *QueryAntiHotlinkingConfigResponseVisitControlRulesRefererControlRuleAdvanceControlRules) *QueryAntiHotlinkingConfigResponseVisitControlRulesRefererControlRule {
  s.AdvanceControlRules = v
  return s
}

type QueryAntiHotlinkingConfigResponseVisitControlRulesRefererControlRuleUaControlRule struct {
  // {"en":"Allows clients, regular matching, to configure multiple UA such as:
  // < valid user - agents > Android | iPhone < / valid - the user - agents >", "zh_CN":"允许的客户端，正则匹配，配置多个UA如：<valid-user-agents>Android|iPhone</valid-user-agents>"}
  ValidUserAgents *string `json:"valid-user-agents,omitempty" xml:"valid-user-agents,omitempty" require:"true"`
  // {"en":"Forbidden client, regular match, configure multiple UA such as:
  // <invalid-user-agents>Android|iPhone</invalid-user-agents>", "zh_CN":"禁止的客户端，正则匹配，配置多个UA如：<invalid-user-agents>Android|iPhone</invalid-user-agents>"}
  InvalidUserAgents *string `json:"invalid-user-agents,omitempty" xml:"invalid-user-agents,omitempty" require:"true"`
}

func (s QueryAntiHotlinkingConfigResponseVisitControlRulesRefererControlRuleUaControlRule) String() string {
  return tea.Prettify(s)
}

func (s QueryAntiHotlinkingConfigResponseVisitControlRulesRefererControlRuleUaControlRule) GoString() string {
  return s.String()
}

func (s *QueryAntiHotlinkingConfigResponseVisitControlRulesRefererControlRuleUaControlRule) SetValidUserAgents(v string) *QueryAntiHotlinkingConfigResponseVisitControlRulesRefererControlRuleUaControlRule {
  s.ValidUserAgents = &v
  return s
}

func (s *QueryAntiHotlinkingConfigResponseVisitControlRulesRefererControlRuleUaControlRule) SetInvalidUserAgents(v string) *QueryAntiHotlinkingConfigResponseVisitControlRulesRefererControlRuleUaControlRule {
  s.InvalidUserAgents = &v
  return s
}

type QueryAntiHotlinkingConfigResponseVisitControlRulesRefererControlRuleAdvanceControlRules struct {
  // {"en":"Forbidden visitor regions, separate with semicolons. Note:
  // 1. Only support ISO 3166-1-alpha-2 two-letter country codes.
  // 2. If you have special regional configuration requirements, please contact your technical support.
  // 3. In the same set of rules, forbidden and allowed visitor regions cannot be configured at the same time.", "zh_CN":"禁止的访客区域，多个请用英文分号分隔。注意
  // 1、仅支持iso 3166-1国家二字简称
  // 2、如果有特殊区域配置需求，请联系您的专属。
  // 3、同一组规则里，禁止的访客区域、允许的访客区域，不能同时配"}
  InvalidVisitorRegion *string `json:"invalid-visitor-region,omitempty" xml:"invalid-visitor-region,omitempty" require:"true"`
  // {"en":"Allowed  visitor regions, separate with semicolons. Note:
  // 1. Only support ISO 3166-1-alpha-2 two-letter country codes.
  // 2. If you have special regional configuration requirements, please contact your technical support.
  // 3. In the same set of rules, forbidden and allowed visitor regions cannot be configured at the same time.", "zh_CN":"允许的访客区域，多个请用英文分号分隔。注意
  // 1、仅支持iso 3166-1国家二字简称
  // 2、如果有特殊区域配置需求，请联系您的专属。
  // 3、同一组规则里，禁止的访客区域、允许的访客区域，不能同时配"}
  ValidVisitorRegion *string `json:"valid-visitor-region,omitempty" xml:"valid-visitor-region,omitempty" require:"true"`
}

func (s QueryAntiHotlinkingConfigResponseVisitControlRulesRefererControlRuleAdvanceControlRules) String() string {
  return tea.Prettify(s)
}

func (s QueryAntiHotlinkingConfigResponseVisitControlRulesRefererControlRuleAdvanceControlRules) GoString() string {
  return s.String()
}

func (s *QueryAntiHotlinkingConfigResponseVisitControlRulesRefererControlRuleAdvanceControlRules) SetInvalidVisitorRegion(v string) *QueryAntiHotlinkingConfigResponseVisitControlRulesRefererControlRuleAdvanceControlRules {
  s.InvalidVisitorRegion = &v
  return s
}

func (s *QueryAntiHotlinkingConfigResponseVisitControlRulesRefererControlRuleAdvanceControlRules) SetValidVisitorRegion(v string) *QueryAntiHotlinkingConfigResponseVisitControlRulesRefererControlRuleAdvanceControlRules {
  s.ValidVisitorRegion = &v
  return s
}

type QueryAntiHotlinkingConfigResponseVisitControlRulesCookieControlRules struct {
  // {"en":"Allow Cookie.Fill in regular format, e.g(. *) (range1 | range2) (. *)", "zh_CN":"允许的cookie。填写正则格式，比如(.*)(range1|range2)(.*)。"}
  AllowCookie *string `json:"allow-cookie,omitempty" xml:"allow-cookie,omitempty" require:"true"`
  // {"en":"Allow Null Cookie.Only true and false are allowed to be filled in allow null cookie", "zh_CN":"是否允许空cookie。只允许填写true或false。"}
  AllowNullCookie *string `json:"allow-null-cookie,omitempty" xml:"allow-null-cookie,omitempty" require:"true"`
  // {"en":"Forbidden Cookie.Fill in regular format, e.g(. *) (range1 | range2) (. *)", "zh_CN":"禁止的cookie。填写正则格式，比如(.*)(range1|range2)(.*)。"}
  ForbiddenCookie *string `json:"forbidden-cookie,omitempty" xml:"forbidden-cookie,omitempty" require:"true"`
}

func (s QueryAntiHotlinkingConfigResponseVisitControlRulesCookieControlRules) String() string {
  return tea.Prettify(s)
}

func (s QueryAntiHotlinkingConfigResponseVisitControlRulesCookieControlRules) GoString() string {
  return s.String()
}

func (s *QueryAntiHotlinkingConfigResponseVisitControlRulesCookieControlRules) SetAllowCookie(v string) *QueryAntiHotlinkingConfigResponseVisitControlRulesCookieControlRules {
  s.AllowCookie = &v
  return s
}

func (s *QueryAntiHotlinkingConfigResponseVisitControlRulesCookieControlRules) SetAllowNullCookie(v string) *QueryAntiHotlinkingConfigResponseVisitControlRulesCookieControlRules {
  s.AllowNullCookie = &v
  return s
}

func (s *QueryAntiHotlinkingConfigResponseVisitControlRulesCookieControlRules) SetForbiddenCookie(v string) *QueryAntiHotlinkingConfigResponseVisitControlRulesCookieControlRules {
  s.ForbiddenCookie = &v
  return s
}

type QueryAntiHotlinkingConfigResponseVisitControlRulesCustomHeaderControlRules struct {
  // {"en":"Header Direction.Can choose from the client or the server.Only allowed to fill client or server", "zh_CN":"来源。可选择来源于客户端还是服务端。客户端填写client，服务端填写server"}
  HeaderDirection *string `json:"header-direction,omitempty" xml:"header-direction,omitempty" require:"true"`
  // {"en":"Header Whitelist", "zh_CN":"头域白名单"}
  HeaderWhitelist *string `json:"header-whitelist,omitempty" xml:"header-whitelist,omitempty" require:"true"`
  // {"en":"Header Value Whitelist", "zh_CN":"头域值白名单"}
  HeaderValueWhitelist *string `json:"header-value-whitelist,omitempty" xml:"header-value-whitelist,omitempty" require:"true"`
  // {"en":"Header Blacklist", "zh_CN":"头域黑名单"}
  HeaderBlacklist *string `json:"header-blacklist,omitempty" xml:"header-blacklist,omitempty" require:"true"`
  // {"en":"Header Value Blacklist", "zh_CN":"头域值黑名单"}
  HeaderValueBlacklist *string `json:"header-value-blacklist,omitempty" xml:"header-value-blacklist,omitempty" require:"true"`
}

func (s QueryAntiHotlinkingConfigResponseVisitControlRulesCustomHeaderControlRules) String() string {
  return tea.Prettify(s)
}

func (s QueryAntiHotlinkingConfigResponseVisitControlRulesCustomHeaderControlRules) GoString() string {
  return s.String()
}

func (s *QueryAntiHotlinkingConfigResponseVisitControlRulesCustomHeaderControlRules) SetHeaderDirection(v string) *QueryAntiHotlinkingConfigResponseVisitControlRulesCustomHeaderControlRules {
  s.HeaderDirection = &v
  return s
}

func (s *QueryAntiHotlinkingConfigResponseVisitControlRulesCustomHeaderControlRules) SetHeaderWhitelist(v string) *QueryAntiHotlinkingConfigResponseVisitControlRulesCustomHeaderControlRules {
  s.HeaderWhitelist = &v
  return s
}

func (s *QueryAntiHotlinkingConfigResponseVisitControlRulesCustomHeaderControlRules) SetHeaderValueWhitelist(v string) *QueryAntiHotlinkingConfigResponseVisitControlRulesCustomHeaderControlRules {
  s.HeaderValueWhitelist = &v
  return s
}

func (s *QueryAntiHotlinkingConfigResponseVisitControlRulesCustomHeaderControlRules) SetHeaderBlacklist(v string) *QueryAntiHotlinkingConfigResponseVisitControlRulesCustomHeaderControlRules {
  s.HeaderBlacklist = &v
  return s
}

func (s *QueryAntiHotlinkingConfigResponseVisitControlRulesCustomHeaderControlRules) SetHeaderValueBlacklist(v string) *QueryAntiHotlinkingConfigResponseVisitControlRulesCustomHeaderControlRules {
  s.HeaderValueBlacklist = &v
  return s
}

type QueryAntiHotlinkingConfigPaths struct {
  // {"en":"the domain whoes need query config", "zh_CN":"需要查询配置的域名或域名id"}
  DomainName *string `json:"domain-name,omitempty" xml:"domain-name,omitempty" require:"true"`
}

func (s QueryAntiHotlinkingConfigPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryAntiHotlinkingConfigPaths) GoString() string {
  return s.String()
}

func (s *QueryAntiHotlinkingConfigPaths) SetDomainName(v string) *QueryAntiHotlinkingConfigPaths {
  s.DomainName = &v
  return s
}

type QueryAntiHotlinkingConfigParameters struct {
}

func (s QueryAntiHotlinkingConfigParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryAntiHotlinkingConfigParameters) GoString() string {
  return s.String()
}

type QueryAntiHotlinkingConfigRequestHeader struct {
}

func (s QueryAntiHotlinkingConfigRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryAntiHotlinkingConfigRequestHeader) GoString() string {
  return s.String()
}

type QueryAntiHotlinkingConfigResponseHeader struct {
}

func (s QueryAntiHotlinkingConfigResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryAntiHotlinkingConfigResponseHeader) GoString() string {
  return s.String()
}




type QueryCacheTimeRequest struct {
}

func (s QueryCacheTimeRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryCacheTimeRequest) GoString() string {
  return s.String()
}

type QueryCacheTimeResponse struct {
  // {"en":"Domain name or domain name id to query configuration", "zh_CN":"需要查询配置的域名"}
  DomainName *string `json:"domain-name,omitempty" xml:"domain-name,omitempty" require:"true"`
  // {"en":"Domain name or domain name id to query configuration", "zh_CN":"需要查询配置的域名id"}
  DomainId *string `json:"domain-id,omitempty" xml:"domain-id,omitempty" require:"true"`
  CacheTimeBehaviors []*QueryCacheTimeResponseCacheTimeBehaviors `json:"cache-time-behaviors,omitempty" xml:"cache-time-behaviors,omitempty" require:"true" type:"Repeated"`
}

func (s QueryCacheTimeResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryCacheTimeResponse) GoString() string {
  return s.String()
}

func (s *QueryCacheTimeResponse) SetDomainName(v string) *QueryCacheTimeResponse {
  s.DomainName = &v
  return s
}

func (s *QueryCacheTimeResponse) SetDomainId(v string) *QueryCacheTimeResponse {
  s.DomainId = &v
  return s
}

func (s *QueryCacheTimeResponse) SetCacheTimeBehaviors(v []*QueryCacheTimeResponseCacheTimeBehaviors) *QueryCacheTimeResponse {
  s.CacheTimeBehaviors = v
  return s
}

type QueryCacheTimeResponseCacheTimeBehaviors struct     {
  // {"en":"Add a grid type identifier to indicate a specific group configuration when the client has multiple groups of configurations.", "zh_CN":"添加grid类型标识，表示客户多组配置时，具体某组配置"}
  DataId *int32 `json:"data-id,omitempty" xml:"data-id,omitempty" require:"true"`
  // {"en":"The url matching mode supports fuzzy regularization. If all matches, the input parameters can be configured as: *", "zh_CN":"url匹配模式，支持正则，如果是全部匹配，入参可以配置为：.*"}
  PathPattern *string `json:"path-pattern,omitempty" xml:"path-pattern,omitempty" require:"true"`
  // {"en":"Exceptional url matching mode, except for some URLs: such as abc.jpg, do not do anti-theft chain function
  // E.g: ^https?://[^/]+/.*\.m3u8", "zh_CN":"例外的url匹配模式，某些URL除外：如abc.jpg，不做内容重定向
  // 客户入参参考：^https?://[^/]+/.*\.m3u8"}
  ExceptPathPattern *string `json:"except-path-pattern,omitempty" xml:"except-path-pattern,omitempty" require:"true"`
  // {"en":"Specify common types: Select the domain name that requires the cache  to be all files or the home page. :
  // E.g:
  // All: all files
  // Homepage: homepage", "zh_CN":"指定常用类型：选择缓存域名的是全部文件还是首页。入参参考值：
  // all：全部文件
  // homepage：首页"}
  CustomPattern *string `json:"custom-pattern,omitempty" xml:"custom-pattern,omitempty" require:"true"`
  // {"en":"File Type: Specify the file type for cache settings.
  // File types include: gif png bmp jpeg jpg html htm shtml mp3 wma flv mp4 wmv zip exe rar css txt ico js swf
  // If you need all types, pass all directly. Multiples are separated by semicolons, and all and specific file types cannot be configured at the same time.", "zh_CN":"文件类型：指定需要缓存的文件类型。
  // 文件类型包括：gif png bmp jpeg jpg html htm shtml mp3 wma flv mp4 wmv zip exe rar css txt ico js swf
  // 如果需要全部类型，则直接传all。多个以分号隔开，all和具体文件类型不能同时配置。"}
  FileType *string `json:"file-type,omitempty" xml:"file-type,omitempty" require:"true"`
  // {"en":"Custom file type: Fill in the appropriate identifiable file type according to your needs outside of the specified file type. Can be used with file-type. If the file-type is also configured, the actual file type is the sum of the two parameters.", "zh_CN":"自定义文件类型：在指定文件类型外根据自身需求，填写适当的可识别文件类型。可以搭配file-type使用。如果file-type也有配置，实际生效的文件类型是两个入参的总和"}
  CustomFileType *string `json:"custom-file-type,omitempty" xml:"custom-file-type,omitempty" require:"true"`
  // {"en":"Specify URL cache: Specify url according to requirements for cache
  // INS format does not support URI format with http(s)://", "zh_CN":"指定URL缓存：根据需求指定url进行缓存
  // 入参不支持含http(s):// 开头的URI格式"}
  SpecifyUrlPattern *string `json:"specify-url-pattern,omitempty" xml:"specify-url-pattern,omitempty" require:"true"`
  // {"en":"Directory: Specify the directory cache.
  // Enter a legal directory format. Multiple separated by semicolons", "zh_CN":"目录：指定目录缓存。
  // 输入合法的目录格式。多个以英文分号隔开"}
  Directory *string `json:"directory,omitempty" xml:"directory,omitempty" require:"true"`
  // {"en":"Cache time: set the time corresponding to the cache object
  // Input format: integer plus unit, such as 20s, 30m, 1h, 2d, no cache is set to 0. Do not enter the unit default is seconds
  // There is no upper limit on the cache time theory. This time is set according to the customer's own needs. If the customer feels that some of the files are not changed frequently, then the setting is longer. For example, the text class js, css, html, etc. can be set shorter, the picture, video and audio classes can be set longer (because the cache time will be replaced by the new file due to the file heat algorithm, the longest suggestion Do not exceed one month)", "zh_CN":"缓存时间：设置缓存对象对应的时间
  // 入参格式：整数加单位，比如20s、30m、1h、2d，不缓存设置为0。不输入单位默认是秒
  // 缓存时间理论上没有上限限制，这个时间根据客户自身的需求设定，如果客户觉得其中一些文件，变更不频繁，那么就设置长一点。例如，文本类的js，css，html等可以设置得短一些，图片、视频音频类的可以设置的长一点（因为缓存时间会因文件热度算法，旧文件会被新文件替换掉，最长建议不要超过一个月）"}
  CacheTtl *int32 `json:"cache-ttl,omitempty" xml:"cache-ttl,omitempty" require:"true"`
  // {"en":"Ignore the source station does not cache the header. The optional values are true and false, which are used to ignore the two configurations of cache-control in the request header (private, no-cache) and the Authorization set by the client.
  // The ture indicates that the source station's settings for the three are ignored. Enables resources to be cached on the service node in the form of cache-control: public, and then our nodes can cache this type of resource and provide acceleration services.
  // False means that when the source station sets cache-control: private, cache-control: no-cache for a resource or specifies to cache according to authorization, our service node will not cache such files.", "zh_CN":"忽略源站不缓存头。可选值为true和false，用于忽略请求头中cache-control的两种配置（private，no-cache）和客户端设置的Authorization。
  // ture表示会忽略掉源站对于这三者的设定。使得资源能够以cache-control: public的方式缓存在服务节点上，然后我们的节点才能缓存这种类型的资源，提供加速服务。
  // false表示当源站对某种资源设定了cache-control: private,cache-control:no-cache或指定根据authorization进行缓存时，我们的服务节点将不会对此类文件进行缓存。"}
  IgnoreCacheControl *string `json:"ignore-cache-control,omitempty" xml:"ignore-cache-control,omitempty" require:"true"`
  // {"en":"Respect the server: Accelerate whether to prioritize the source cache time.
  // Optional values: true and false
  // True: indicates that the server is time-first
  // False: The cache time of the CDN configuration takes precedence.", "zh_CN":"尊重服务端：加速是否要按源站缓存时间优先。
  // 可选值：true和false
  // true：表示重服务端时间优先
  // false:CDN配置的缓存时间优先"}
  IsRespectServer *string `json:"is-respect-server,omitempty" xml:"is-respect-server,omitempty" require:"true"`
  // {"en":"Ignore case, the optional value is true or false, true means to ignore case; false means not to ignore case;
  // When adding a new configuration item, the default is not true.", "zh_CN":"忽略大小写，可选值为true或false，true表示忽略大小写；false表示不忽略大小写；
  // 新增配置项时，不传默认为 true"}
  IgnoreLetterCase *string `json:"ignore-letter-case,omitempty" xml:"ignore-letter-case,omitempty" require:"true"`
  // {"en":"Reload processing rules, optional: ignore or if-modified-since
  // If-modified-since: indicates that you want to convert to if-modified-since
  // Ignore: means to ignore client refresh", "zh_CN":"reload处理规则，可选项：ignore或者if-modified-since
  // if-modified-since：表示要转成if-modified-since
  // ignore:表示忽略客户端刷新"}
  ReloadManage *string `json:"reload-manage,omitempty" xml:"reload-manage,omitempty" require:"true"`
  // {"en":"Indicates the priority execution order of multiple sets of redirected content by the customer. The higher the number, the higher the priority.
  // When adding a new configuration item, the default is 10", "zh_CN":"表示客户多组重定向内容的优先执行顺序。数字越大，优先级越高。
  // 新增配置项时，不传默认为 10
  // 如果传了值，不能为空"}
  Priority *int32 `json:"priority,omitempty" xml:"priority,omitempty" require:"true"`
  // {"en":"Request Header", "zh_CN":"请求头域"}
  QueryCacheTimeRequestHeaderField *string `json:"request-header-field,omitempty" xml:"request-header-field,omitempty" require:"true"`
  // {"en":"Request Header Value", "zh_CN":"请求头的值"}
  ValueQueryCacheTimeRequestHeader *string `json:"value-request-header,omitempty" xml:"value-request-header,omitempty" require:"true"`
  // {"en":"You can set it 'true' to cache
  // ignoring the http header 'Authentication'.  If it is empty, the header is not ignored by default.", "zh_CN":"忽略鉴权头部Authentication，可选值为true和false。为空默认为不忽略。"}
  IgnoreAuthenticationHeader *bool `json:"ignore-authentication-header,omitempty" xml:"ignore-authentication-header,omitempty" require:"true"`
}

func (s QueryCacheTimeResponseCacheTimeBehaviors) String() string {
  return tea.Prettify(s)
}

func (s QueryCacheTimeResponseCacheTimeBehaviors) GoString() string {
  return s.String()
}

func (s *QueryCacheTimeResponseCacheTimeBehaviors) SetDataId(v int32) *QueryCacheTimeResponseCacheTimeBehaviors {
  s.DataId = &v
  return s
}

func (s *QueryCacheTimeResponseCacheTimeBehaviors) SetPathPattern(v string) *QueryCacheTimeResponseCacheTimeBehaviors {
  s.PathPattern = &v
  return s
}

func (s *QueryCacheTimeResponseCacheTimeBehaviors) SetExceptPathPattern(v string) *QueryCacheTimeResponseCacheTimeBehaviors {
  s.ExceptPathPattern = &v
  return s
}

func (s *QueryCacheTimeResponseCacheTimeBehaviors) SetCustomPattern(v string) *QueryCacheTimeResponseCacheTimeBehaviors {
  s.CustomPattern = &v
  return s
}

func (s *QueryCacheTimeResponseCacheTimeBehaviors) SetFileType(v string) *QueryCacheTimeResponseCacheTimeBehaviors {
  s.FileType = &v
  return s
}

func (s *QueryCacheTimeResponseCacheTimeBehaviors) SetCustomFileType(v string) *QueryCacheTimeResponseCacheTimeBehaviors {
  s.CustomFileType = &v
  return s
}

func (s *QueryCacheTimeResponseCacheTimeBehaviors) SetSpecifyUrlPattern(v string) *QueryCacheTimeResponseCacheTimeBehaviors {
  s.SpecifyUrlPattern = &v
  return s
}

func (s *QueryCacheTimeResponseCacheTimeBehaviors) SetDirectory(v string) *QueryCacheTimeResponseCacheTimeBehaviors {
  s.Directory = &v
  return s
}

func (s *QueryCacheTimeResponseCacheTimeBehaviors) SetCacheTtl(v int32) *QueryCacheTimeResponseCacheTimeBehaviors {
  s.CacheTtl = &v
  return s
}

func (s *QueryCacheTimeResponseCacheTimeBehaviors) SetIgnoreCacheControl(v string) *QueryCacheTimeResponseCacheTimeBehaviors {
  s.IgnoreCacheControl = &v
  return s
}

func (s *QueryCacheTimeResponseCacheTimeBehaviors) SetIsRespectServer(v string) *QueryCacheTimeResponseCacheTimeBehaviors {
  s.IsRespectServer = &v
  return s
}

func (s *QueryCacheTimeResponseCacheTimeBehaviors) SetIgnoreLetterCase(v string) *QueryCacheTimeResponseCacheTimeBehaviors {
  s.IgnoreLetterCase = &v
  return s
}

func (s *QueryCacheTimeResponseCacheTimeBehaviors) SetReloadManage(v string) *QueryCacheTimeResponseCacheTimeBehaviors {
  s.ReloadManage = &v
  return s
}

func (s *QueryCacheTimeResponseCacheTimeBehaviors) SetPriority(v int32) *QueryCacheTimeResponseCacheTimeBehaviors {
  s.Priority = &v
  return s
}

func (s *QueryCacheTimeResponseCacheTimeBehaviors) SetQueryCacheTimeRequestHeaderField(v string) *QueryCacheTimeResponseCacheTimeBehaviors {
  s.QueryCacheTimeRequestHeaderField = &v
  return s
}

func (s *QueryCacheTimeResponseCacheTimeBehaviors) SetValueQueryCacheTimeRequestHeader(v string) *QueryCacheTimeResponseCacheTimeBehaviors {
  s.ValueQueryCacheTimeRequestHeader = &v
  return s
}

func (s *QueryCacheTimeResponseCacheTimeBehaviors) SetIgnoreAuthenticationHeader(v bool) *QueryCacheTimeResponseCacheTimeBehaviors {
  s.IgnoreAuthenticationHeader = &v
  return s
}

type QueryCacheTimePaths struct {
  // {"en":"Domain name or domain name id to query configuration", "zh_CN":"需要查询配置的域名或域名id"}
  DomainName *string `json:"domain-name,omitempty" xml:"domain-name,omitempty" require:"true"`
}

func (s QueryCacheTimePaths) String() string {
  return tea.Prettify(s)
}

func (s QueryCacheTimePaths) GoString() string {
  return s.String()
}

func (s *QueryCacheTimePaths) SetDomainName(v string) *QueryCacheTimePaths {
  s.DomainName = &v
  return s
}

type QueryCacheTimeParameters struct {
}

func (s QueryCacheTimeParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryCacheTimeParameters) GoString() string {
  return s.String()
}

type QueryCacheTimeRequestHeader struct {
}

func (s QueryCacheTimeRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryCacheTimeRequestHeader) GoString() string {
  return s.String()
}

type QueryCacheTimeResponseHeader struct {
}

func (s QueryCacheTimeResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryCacheTimeResponseHeader) GoString() string {
  return s.String()
}




type QueryDomainCertConfigRequest struct {
}

func (s QueryDomainCertConfigRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryDomainCertConfigRequest) GoString() string {
  return s.String()
}

type QueryDomainCertConfigResponse struct {
  // {"en":"Error code, which appears when HTTPStatus is not 202, represents the error type of the current request call", "zh_CN":"错误代码，当HTTPStatus不为202时出现，表示当前请求调用的错误类型"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response information, success when successful", "zh_CN":"响应信息，成功时为success"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  Data *QueryDomainCertConfigResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
}

func (s QueryDomainCertConfigResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryDomainCertConfigResponse) GoString() string {
  return s.String()
}

func (s *QueryDomainCertConfigResponse) SetCode(v string) *QueryDomainCertConfigResponse {
  s.Code = &v
  return s
}

func (s *QueryDomainCertConfigResponse) SetMessage(v string) *QueryDomainCertConfigResponse {
  s.Message = &v
  return s
}

func (s *QueryDomainCertConfigResponse) SetData(v *QueryDomainCertConfigResponseData) *QueryDomainCertConfigResponse {
  s.Data = v
  return s
}

type QueryDomainCertConfigResponseData struct {
  // {"en":"domain id", "zh_CN":"域名id"}
  DomainId *string `json:"domainId,omitempty" xml:"domainId,omitempty" require:"true"`
  // {"en":"domain name", "zh_CN":"域名名称"}
  DomainName *string `json:"domainName,omitempty" xml:"domainName,omitempty" require:"true"`
  // {"en":"The certificate ID you want to bind.", "zh_CN":"证书ID，为域名关联证书或换证书。请注意：置空时为取消关联证书。"}
  CertificateId *int `json:"certificateId,omitempty" xml:"certificateId,omitempty" require:"true"`
  // {"en":"TLS version. Optional values: SSLv3,TLSv1,TLSv1.1,TLSv1.2,TLSv1.3", "zh_CN":"TLS协议（TLSVersion），支持配置多个，英文分号分隔。可选值为: SSLv3、TLSv1、TLSv1.1、TLSv1.2、TLSv1.3。域名有配置证书时，该配置才有效。"}
  TLSVersion *string `json:"TLSVersion,omitempty" xml:"TLSVersion,omitempty" require:"true"`
  // {"en":"Enable OCSP(Online Certificate Status Protocol).", "zh_CN":"支持OCSP，默认不启用，可选值true、false。域名有配置证书时，该配置才有效。"}
  EnableOCSP *string `json:"enableOCSP,omitempty" xml:"enableOCSP,omitempty" require:"true"`
  // {"en":"This optional object is used to specify a colon separated list of cipher suites which are permitted when clients negotiate security settings to access your content. Cipher suites which you can specify are: LOW, ALL:!LOW, HIGH, !EXPORT, !aNULL, !RC4, !DH, !SHA, !MD5, @STRENGTH,  AES128-SHA, AES256-SHA, AES128-SHA256, AES256-SHA256, AES128-GCM-SHA256, AES256-GCM-SHA384, ECDHE-RSA-AES128-SHA, ECDHE-RSA-AES256-SHA, ECDHE-RSA-AES128-SHA256, ECDHE-RSA-AES256-SHA384, ECDHE-RSA-AES128-GCM-SHA256, and ECDHE-RSA-AES256-GCM-SHA384. These cipher suites are a subset of those supported by OpenSSL, https://www.openssl.org/docs/man1.0.2/man1/ciphers.html. Please note that !MD5 or !SHA must appear after HIGH.", "zh_CN":"设置cipher加密套件，冒号分隔。请注意：系统已有默认的算法，如无特殊修改需要，无需调整。"}
  CipherSuites *string `json:"cipherSuites,omitempty" xml:"cipherSuites,omitempty" require:"true"`
}

func (s QueryDomainCertConfigResponseData) String() string {
  return tea.Prettify(s)
}

func (s QueryDomainCertConfigResponseData) GoString() string {
  return s.String()
}

func (s *QueryDomainCertConfigResponseData) SetDomainId(v string) *QueryDomainCertConfigResponseData {
  s.DomainId = &v
  return s
}

func (s *QueryDomainCertConfigResponseData) SetDomainName(v string) *QueryDomainCertConfigResponseData {
  s.DomainName = &v
  return s
}

func (s *QueryDomainCertConfigResponseData) SetCertificateId(v int) *QueryDomainCertConfigResponseData {
  s.CertificateId = &v
  return s
}

func (s *QueryDomainCertConfigResponseData) SetTLSVersion(v string) *QueryDomainCertConfigResponseData {
  s.TLSVersion = &v
  return s
}

func (s *QueryDomainCertConfigResponseData) SetEnableOCSP(v string) *QueryDomainCertConfigResponseData {
  s.EnableOCSP = &v
  return s
}

func (s *QueryDomainCertConfigResponseData) SetCipherSuites(v string) *QueryDomainCertConfigResponseData {
  s.CipherSuites = &v
  return s
}

type QueryDomainCertConfigPaths struct {
  // {"en":"The domain whoes need query config.", "zh_CN":"需要查询配置的域名或域名id"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
}

func (s QueryDomainCertConfigPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryDomainCertConfigPaths) GoString() string {
  return s.String()
}

func (s *QueryDomainCertConfigPaths) SetDomain(v string) *QueryDomainCertConfigPaths {
  s.Domain = &v
  return s
}

type QueryDomainCertConfigParameters struct {
}

func (s QueryDomainCertConfigParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryDomainCertConfigParameters) GoString() string {
  return s.String()
}

type QueryDomainCertConfigRequestHeader struct {
}

func (s QueryDomainCertConfigRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryDomainCertConfigRequestHeader) GoString() string {
  return s.String()
}

type QueryDomainCertConfigResponseHeader struct {
}

func (s QueryDomainCertConfigResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryDomainCertConfigResponseHeader) GoString() string {
  return s.String()
}




type QueryLivestreamingTimestampAntihotlinkingConfigRequest struct {
}

func (s QueryLivestreamingTimestampAntihotlinkingConfigRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryLivestreamingTimestampAntihotlinkingConfigRequest) GoString() string {
  return s.String()
}

type QueryLivestreamingTimestampAntihotlinkingConfigResponse struct {
  // {"en":"The error code, when HTTPStatus is not 200, indicates the type of error the current request is calling.", "zh_CN":"错误代码，当HTTPStatus不为200时出现，表示当前请求调用的错误类型"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response information, when success is successful", "zh_CN":"响应信息，成功时为success"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"The response data", "zh_CN":"响应数据"}
  Data *QueryLivestreamingTimestampAntihotlinkingConfigResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
  // {"en":"Uniquely labeled id for querying each requested task (for all interfaces)", "zh_CN":"唯一标示的id，用于查询每次请求的任务 （适用全部接口）"}
  XCncRequestId *string `json:"x-cnc-request-id,omitempty" xml:"x-cnc-request-id,omitempty" require:"true"`
}

func (s QueryLivestreamingTimestampAntihotlinkingConfigResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryLivestreamingTimestampAntihotlinkingConfigResponse) GoString() string {
  return s.String()
}

func (s *QueryLivestreamingTimestampAntihotlinkingConfigResponse) SetCode(v string) *QueryLivestreamingTimestampAntihotlinkingConfigResponse {
  s.Code = &v
  return s
}

func (s *QueryLivestreamingTimestampAntihotlinkingConfigResponse) SetMessage(v string) *QueryLivestreamingTimestampAntihotlinkingConfigResponse {
  s.Message = &v
  return s
}

func (s *QueryLivestreamingTimestampAntihotlinkingConfigResponse) SetData(v *QueryLivestreamingTimestampAntihotlinkingConfigResponseData) *QueryLivestreamingTimestampAntihotlinkingConfigResponse {
  s.Data = v
  return s
}

func (s *QueryLivestreamingTimestampAntihotlinkingConfigResponse) SetXCncRequestId(v string) *QueryLivestreamingTimestampAntihotlinkingConfigResponse {
  s.XCncRequestId = &v
  return s
}

type QueryLivestreamingTimestampAntihotlinkingConfigResponseData struct {
  DomainId *string `json:"domainId,omitempty" xml:"domainId,omitempty" require:"true"`
  DomainName *string `json:"domainName,omitempty" xml:"domainName,omitempty" require:"true"`
  // {"en":"Streaming media anti-hotlinking configuration, parent tag
  // 1. This must be filled when the hotlinking configuration of streaming media needs to be set
  // 2. Empty the configuration for <timestampVisitControls/>", "zh_CN":"流媒体防盗链配置，父标签
  // 1.需要设置流媒体防盗链配置时，此项必填
  // 2.为<timestampVisitControls/>时清空配置"}
  TimestampVisitControlRules []*QueryLivestreamingTimestampAntihotlinkingConfigResponseDataTimestampVisitControlRules `json:"timestampVisitControlRules,omitempty" xml:"timestampVisitControlRules,omitempty" require:"true" type:"Repeated"`
}

func (s QueryLivestreamingTimestampAntihotlinkingConfigResponseData) String() string {
  return tea.Prettify(s)
}

func (s QueryLivestreamingTimestampAntihotlinkingConfigResponseData) GoString() string {
  return s.String()
}

func (s *QueryLivestreamingTimestampAntihotlinkingConfigResponseData) SetDomainId(v string) *QueryLivestreamingTimestampAntihotlinkingConfigResponseData {
  s.DomainId = &v
  return s
}

func (s *QueryLivestreamingTimestampAntihotlinkingConfigResponseData) SetDomainName(v string) *QueryLivestreamingTimestampAntihotlinkingConfigResponseData {
  s.DomainName = &v
  return s
}

func (s *QueryLivestreamingTimestampAntihotlinkingConfigResponseData) SetTimestampVisitControlRules(v []*QueryLivestreamingTimestampAntihotlinkingConfigResponseDataTimestampVisitControlRules) *QueryLivestreamingTimestampAntihotlinkingConfigResponseData {
  s.TimestampVisitControlRules = v
  return s
}

type QueryLivestreamingTimestampAntihotlinkingConfigResponseDataTimestampVisitControlRules struct     {
  // {"en":"The name of the time parameter, and the default value is wsTime if no parameters are entered.", "zh_CN":"密文参数名称，未入参则默认值为wsSecret"}
  CipherParam *string `json:"cipherParam,omitempty" xml:"cipherParam,omitempty" require:"true"`
  // {"en":"The name of the time parameter, and the default value is wsTime if no parameters are entered.", "zh_CN":"时间参数名称，未入参则默认值为wsTime"}
  TimeParam *string `json:"timeParam,omitempty" xml:"timeParam,omitempty" require:"true"`
  // {"en":"The secret key", "zh_CN":"密钥"}
  SecretKey *string `json:"secretKey,omitempty" xml:"secretKey,omitempty" require:"true"`
  // {"en":"The encrypted time format, is required， with values of [UNIX timestamp] and [hex], corresponding to [UNIX timestamp] and [hex], respectively.", "zh_CN":"加密时间格式，必填，可选值为【UNIX timestamp】和【hex】，分别为【UNIX时间戳】、【16进制】。"}
  TimeFormat *string `json:"timeFormat,omitempty" xml:"timeFormat,omitempty" require:"true"`
  // {"en":"Effective time method, optional values: duration, abstime, no verification.
  // 1. duration, means get effective time by duration. When this mode is set, effectiveTime is required.
  // 2. abstime, means get effective time by absolute time. Support a URL to be valid for a certain future time (the maximum 3652 days)
  // 3. No verification, means it will not verify the time. When this mode is set, effectiveTime will be cleared cleaned.", "zh_CN":"有效时间计算方式，可选值：duration、abstime、no verification
  // 1、duration，表示按时长。设置为“按时长”，则effectiveTime必填。
  // 2、abstime，表示按绝对时间。支持某个url到某个未来时间有效（最大支持3652天）
  // 3、no verification，表示不校验时间。则自动清空effectiveTime。"}
  EffectiveTimeMode *string `json:"effectiveTimeMode,omitempty" xml:"effectiveTimeMode,omitempty" require:"true"`
  // {"en":"The effective length, unit second (s), positive integer,and value >=1", "zh_CN":"有效时长，单位秒（s），正整数，值要>=1。"}
  EffectiveTime *int `json:"effectiveTime,omitempty" xml:"effectiveTime,omitempty" require:"true"`
  // {"en":"The tolerance value, error time, unit second (s), positive integer and value >=1", "zh_CN":"容忍值，误差时间，单位秒（s），正整数，值要>=1"}
  TolerantTime *int `json:"tolerantTime,omitempty" xml:"tolerantTime,omitempty" require:"true"`
  // {"en":"Ciphertext combination mode, with the following six optional values:
  // 
  // key+path+time
  // key+time+path
  // path+key+time
  // path+time+key
  // time+path+key
  // time+key+path", "zh_CN":"密文组合方式，可选值为以下六项：
  // 
  // key+path+time
  // key+time+path
  // path+key+time
  // path+time+key
  // time+path+key
  // time+key+path"}
  CipheCombination *string `json:"cipheCombination,omitempty" xml:"cipheCombination,omitempty" require:"true"`
  // {"en":"Data-id is to indicate a specific group configuration when the client has multiple groups of configurations. Data-id can be retrieved through a query interface. Note: A. If data-id is passed, it means that one group of configuration items is specified to be modified, and no other group configuration items need to be modified. B. If multiple groups of configurations are included, some of them are configured with data-id and others are not, then the expression of data-id is used to modify a specific group of configurations, and a new group of configurations is added on the original basis without the expression of data-id. C. If the data-id is not transmitted, it means that the original configuration will be fully covered by this configuration. D. If no configuration parameter is passed, only domain name and secondary label are passed, which means that all configuration of domain name secondary service corresponding to this interface is cleared. E. If there is no specific configuration item in a set of configurations, the data-id must be filled in, and the value is the actual data-id, which means clearing the value of the corresponding data-id configuration item; it is not allowed that there is no specific configuration item or data-id in a set of configurations.", "zh_CN":"配置多组配置时，具体某组配置的id。dataId可以通过查询接口获取。 注意： a、如果有传dataId，说明指定修改其中一组配置项内容，不需求修改其他组配置内容不需要入参； b、如果入参多组配置，其中有些组配置有传dataId，有些没有传，则有传dataId的表示修改具体某组配置，没有传dataId的表示在原来基础上新增一组配置； c、如果入参都没有传dataId,表示用本次的配置全量覆盖原先配置； d、如果入参没有传任何配置项参数，只传了域名和二级标签，表示清空这个接口对应域名二级服务所有配置； e、如果一组配置没有具体的配置项，则dataId必填，且值为实际存在的dataId，表示清空这个dataId对应配置项的值；不允许一组配置没有具体的配置项也没有dataId。"}
  DataId *int `json:"dataId,omitempty" xml:"dataId,omitempty" require:"true"`
}

func (s QueryLivestreamingTimestampAntihotlinkingConfigResponseDataTimestampVisitControlRules) String() string {
  return tea.Prettify(s)
}

func (s QueryLivestreamingTimestampAntihotlinkingConfigResponseDataTimestampVisitControlRules) GoString() string {
  return s.String()
}

func (s *QueryLivestreamingTimestampAntihotlinkingConfigResponseDataTimestampVisitControlRules) SetCipherParam(v string) *QueryLivestreamingTimestampAntihotlinkingConfigResponseDataTimestampVisitControlRules {
  s.CipherParam = &v
  return s
}

func (s *QueryLivestreamingTimestampAntihotlinkingConfigResponseDataTimestampVisitControlRules) SetTimeParam(v string) *QueryLivestreamingTimestampAntihotlinkingConfigResponseDataTimestampVisitControlRules {
  s.TimeParam = &v
  return s
}

func (s *QueryLivestreamingTimestampAntihotlinkingConfigResponseDataTimestampVisitControlRules) SetSecretKey(v string) *QueryLivestreamingTimestampAntihotlinkingConfigResponseDataTimestampVisitControlRules {
  s.SecretKey = &v
  return s
}

func (s *QueryLivestreamingTimestampAntihotlinkingConfigResponseDataTimestampVisitControlRules) SetTimeFormat(v string) *QueryLivestreamingTimestampAntihotlinkingConfigResponseDataTimestampVisitControlRules {
  s.TimeFormat = &v
  return s
}

func (s *QueryLivestreamingTimestampAntihotlinkingConfigResponseDataTimestampVisitControlRules) SetEffectiveTimeMode(v string) *QueryLivestreamingTimestampAntihotlinkingConfigResponseDataTimestampVisitControlRules {
  s.EffectiveTimeMode = &v
  return s
}

func (s *QueryLivestreamingTimestampAntihotlinkingConfigResponseDataTimestampVisitControlRules) SetEffectiveTime(v int) *QueryLivestreamingTimestampAntihotlinkingConfigResponseDataTimestampVisitControlRules {
  s.EffectiveTime = &v
  return s
}

func (s *QueryLivestreamingTimestampAntihotlinkingConfigResponseDataTimestampVisitControlRules) SetTolerantTime(v int) *QueryLivestreamingTimestampAntihotlinkingConfigResponseDataTimestampVisitControlRules {
  s.TolerantTime = &v
  return s
}

func (s *QueryLivestreamingTimestampAntihotlinkingConfigResponseDataTimestampVisitControlRules) SetCipheCombination(v string) *QueryLivestreamingTimestampAntihotlinkingConfigResponseDataTimestampVisitControlRules {
  s.CipheCombination = &v
  return s
}

func (s *QueryLivestreamingTimestampAntihotlinkingConfigResponseDataTimestampVisitControlRules) SetDataId(v int) *QueryLivestreamingTimestampAntihotlinkingConfigResponseDataTimestampVisitControlRules {
  s.DataId = &v
  return s
}

type QueryLivestreamingTimestampAntihotlinkingConfigPaths struct {
  // {"en":"Domain name or domain name id to query configuration", "zh_CN":"需要查询配置的域名（domainName）或域名id（domainId）"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
}

func (s QueryLivestreamingTimestampAntihotlinkingConfigPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryLivestreamingTimestampAntihotlinkingConfigPaths) GoString() string {
  return s.String()
}

func (s *QueryLivestreamingTimestampAntihotlinkingConfigPaths) SetDomain(v string) *QueryLivestreamingTimestampAntihotlinkingConfigPaths {
  s.Domain = &v
  return s
}

type QueryLivestreamingTimestampAntihotlinkingConfigParameters struct {
}

func (s QueryLivestreamingTimestampAntihotlinkingConfigParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryLivestreamingTimestampAntihotlinkingConfigParameters) GoString() string {
  return s.String()
}

type QueryLivestreamingTimestampAntihotlinkingConfigRequestHeader struct {
}

func (s QueryLivestreamingTimestampAntihotlinkingConfigRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryLivestreamingTimestampAntihotlinkingConfigRequestHeader) GoString() string {
  return s.String()
}

type QueryLivestreamingTimestampAntihotlinkingConfigResponseHeader struct {
}

func (s QueryLivestreamingTimestampAntihotlinkingConfigResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryLivestreamingTimestampAntihotlinkingConfigResponseHeader) GoString() string {
  return s.String()
}




type AddBanUrltoDomianRequest struct {
  // {"en":"Customer code", "zh_CN":"需要非法屏蔽的客户在我司设置的英文名"}
  CustomerCode *string `json:"customer-code,omitempty" xml:"customer-code,omitempty"`
  // {"en":"illegal information", "zh_CN":"非法信息屏蔽配置标签
  // 注意：表示需要设置非法信息屏蔽配置时，此项必填"}
  IllegalInformation []*AddBanUrltoDomianRequestIllegalInformation `json:"illegal-information,omitempty" xml:"illegal-information,omitempty" require:"true" type:"Repeated"`
}

func (s AddBanUrltoDomianRequest) String() string {
  return tea.Prettify(s)
}

func (s AddBanUrltoDomianRequest) GoString() string {
  return s.String()
}

func (s *AddBanUrltoDomianRequest) SetCustomerCode(v string) *AddBanUrltoDomianRequest {
  s.CustomerCode = &v
  return s
}

func (s *AddBanUrltoDomianRequest) SetIllegalInformation(v []*AddBanUrltoDomianRequestIllegalInformation) *AddBanUrltoDomianRequest {
  s.IllegalInformation = v
  return s
}

type AddBanUrltoDomianRequestIllegalInformation struct     {
  // {"en":"Area code, mandatory field, like Global.", "zh_CN":"需要设置非法信息屏蔽的区域。则此项必填.支持多个。如设置全球，则输入 Global"}
  Areas []*string `json:"areas,omitempty" xml:"areas,omitempty" require:"true" type:"Repeated"`
  // {"en":"URL", "zh_CN":"一个区域下配置需要非法信息屏蔽的url，如果某些区域有多个屏蔽url则有多组"}
  BanUrls []*string `json:"ban-urls,omitempty" xml:"ban-urls,omitempty" require:"true" type:"Repeated"`
  // {"en":"Matching method, support fuzzy and precise, fuzzy match: fuzzy, precise match: exact.", "zh_CN":"匹配方式，支持模糊和精确；模糊入参为：fuzzy;精确匹配：exact
  // 如果没有传默认是exact"}
  Method *string `json:"method,omitempty" xml:"method,omitempty"`
}

func (s AddBanUrltoDomianRequestIllegalInformation) String() string {
  return tea.Prettify(s)
}

func (s AddBanUrltoDomianRequestIllegalInformation) GoString() string {
  return s.String()
}

func (s *AddBanUrltoDomianRequestIllegalInformation) SetAreas(v []*string) *AddBanUrltoDomianRequestIllegalInformation {
  s.Areas = v
  return s
}

func (s *AddBanUrltoDomianRequestIllegalInformation) SetBanUrls(v []*string) *AddBanUrltoDomianRequestIllegalInformation {
  s.BanUrls = v
  return s
}

func (s *AddBanUrltoDomianRequestIllegalInformation) SetMethod(v string) *AddBanUrltoDomianRequestIllegalInformation {
  s.Method = &v
  return s
}

type AddBanUrltoDomianResponse struct {
  // {"en":"httpstatus=202; Indicates that the new domain API was successfully invoked, and the current deployment of the new domain can be viewed using x-cnc-request-id in the header", "zh_CN":"httpstatus=202;   表示成功调用新增域名接口，可使用header中的x-cnc-request-id查看当前新增域名的部署情况"}
  HttpStatus *int `json:"http status code,omitempty" xml:"http status code,omitempty" require:"true"`
  // {"en":"Uniquely identified id for querying tasks per request (for all API)", "zh_CN":"唯一标示的id，用于查询每次请求的任务 （适用全部接口）"}
  XCncRequestId *string `json:"x-cnc-request-id,omitempty" xml:"x-cnc-request-id,omitempty" require:"true"`
  // {"en":"The URL used to access the domain information, where domain-id is the unique token generated by our cloud platform for the domain name and whose value is a string.", "zh_CN":"响应信用于访问该域名信息的URL，其中domain-id为我司云平台为该域名生成的唯一标示，其值为字符串。"}
  Location *string `json:"location,omitempty" xml:"location,omitempty" require:"true"`
  // {"en":"Service Domain assigned by the system, e.g. xxxx.cdn30.com", "zh_CN":"由自动生成的服务域名名称，例如：xxxx.cdn30.com"}
  Cname *string `json:"cname,omitempty" xml:"cname,omitempty" require:"true"`
  // {"en":"Error code. It will only have a value when the HTTP Status of response is not 202. ", "zh_CN":"错误代码，当HTTPStatus不为202时出现，表示当前请求调用的错误类型"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response message. When the request is successful, the response message is success.", "zh_CN":"响应信息，成功时为success"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s AddBanUrltoDomianResponse) String() string {
  return tea.Prettify(s)
}

func (s AddBanUrltoDomianResponse) GoString() string {
  return s.String()
}

func (s *AddBanUrltoDomianResponse) SetHttpStatus(v int) *AddBanUrltoDomianResponse {
  s.HttpStatus = &v
  return s
}

func (s *AddBanUrltoDomianResponse) SetXCncRequestId(v string) *AddBanUrltoDomianResponse {
  s.XCncRequestId = &v
  return s
}

func (s *AddBanUrltoDomianResponse) SetLocation(v string) *AddBanUrltoDomianResponse {
  s.Location = &v
  return s
}

func (s *AddBanUrltoDomianResponse) SetCname(v string) *AddBanUrltoDomianResponse {
  s.Cname = &v
  return s
}

func (s *AddBanUrltoDomianResponse) SetCode(v string) *AddBanUrltoDomianResponse {
  s.Code = &v
  return s
}

func (s *AddBanUrltoDomianResponse) SetMessage(v string) *AddBanUrltoDomianResponse {
  s.Message = &v
  return s
}

type AddBanUrltoDomianPaths struct {
}

func (s AddBanUrltoDomianPaths) String() string {
  return tea.Prettify(s)
}

func (s AddBanUrltoDomianPaths) GoString() string {
  return s.String()
}

type AddBanUrltoDomianParameters struct {
}

func (s AddBanUrltoDomianParameters) String() string {
  return tea.Prettify(s)
}

func (s AddBanUrltoDomianParameters) GoString() string {
  return s.String()
}

type AddBanUrltoDomianRequestHeader struct {
}

func (s AddBanUrltoDomianRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s AddBanUrltoDomianRequestHeader) GoString() string {
  return s.String()
}

type AddBanUrltoDomianResponseHeader struct {
}

func (s AddBanUrltoDomianResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s AddBanUrltoDomianResponseHeader) GoString() string {
  return s.String()
}




type EditDomainPropertyRequest struct {
  // {"en":"Origin config.", "zh_CN":"回源配置"}
  OriginConfig *EditDomainPropertyRequestOriginConfig `json:"origin-config,omitempty" xml:"origin-config,omitempty" require:"true" type:"Struct"`
}

func (s EditDomainPropertyRequest) String() string {
  return tea.Prettify(s)
}

func (s EditDomainPropertyRequest) GoString() string {
  return s.String()
}

func (s *EditDomainPropertyRequest) SetOriginConfig(v *EditDomainPropertyRequestOriginConfig) *EditDomainPropertyRequest {
  s.OriginConfig = v
  return s
}

type EditDomainPropertyRequestOriginConfig struct {
  // {"en":"Back to source host, not required", "zh_CN":"回源host，非必填"}
  OriginHost *string `json:"origin-host,omitempty" xml:"origin-host,omitempty"`
  // {"en":"Source IP or domain name. 
  // 1. Support IPV4,IPV6.
  // 2. Multiple IPs are separated by semicolons, such as:1.1.1.1;2.2.2.2.
  // 3. Support single domain name.", "zh_CN":"源站Ip或者域名，非必填。
  // 1.支持IPV4,IPV6
  // 2.多个IP用分号隔开，如1.1.1.1;2.2.2.2
  // 3.支持单个域名"}
  OriginIps *string `json:"origin-ips,omitempty" xml:"origin-ips,omitempty"`
  // {"en":"Origin port. Port range:(0-65535].", "zh_CN":"回源端口，非必填
  // 端口范围：(0-65535]"}
  OriginPort *string `json:"origin-port,omitempty" xml:"origin-port,omitempty"`
}

func (s EditDomainPropertyRequestOriginConfig) String() string {
  return tea.Prettify(s)
}

func (s EditDomainPropertyRequestOriginConfig) GoString() string {
  return s.String()
}

func (s *EditDomainPropertyRequestOriginConfig) SetOriginHost(v string) *EditDomainPropertyRequestOriginConfig {
  s.OriginHost = &v
  return s
}

func (s *EditDomainPropertyRequestOriginConfig) SetOriginIps(v string) *EditDomainPropertyRequestOriginConfig {
  s.OriginIps = &v
  return s
}

func (s *EditDomainPropertyRequestOriginConfig) SetOriginPort(v string) *EditDomainPropertyRequestOriginConfig {
  s.OriginPort = &v
  return s
}

type EditDomainPropertyResponse struct {
  // {"en":"The error code, which appears when HTTPStatus is not 202, indicates the error type of the current request invocation", "zh_CN":"错误代码，当HTTPStatus不为202时出现，表示当前请求调用的错误类型"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response information, when success is successful", "zh_CN":"响应信息，成功时为success"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s EditDomainPropertyResponse) String() string {
  return tea.Prettify(s)
}

func (s EditDomainPropertyResponse) GoString() string {
  return s.String()
}

func (s *EditDomainPropertyResponse) SetCode(v string) *EditDomainPropertyResponse {
  s.Code = &v
  return s
}

func (s *EditDomainPropertyResponse) SetMessage(v string) *EditDomainPropertyResponse {
  s.Message = &v
  return s
}

type EditDomainPropertyPaths struct {
  // {"en":"the domain whoes need query config", "zh_CN":"需要查询配置的域名或域名id"}
  DomainName *string `json:"domain-name,omitempty" xml:"domain-name,omitempty" require:"true"`
}

func (s EditDomainPropertyPaths) String() string {
  return tea.Prettify(s)
}

func (s EditDomainPropertyPaths) GoString() string {
  return s.String()
}

func (s *EditDomainPropertyPaths) SetDomainName(v string) *EditDomainPropertyPaths {
  s.DomainName = &v
  return s
}

type EditDomainPropertyParameters struct {
}

func (s EditDomainPropertyParameters) String() string {
  return tea.Prettify(s)
}

func (s EditDomainPropertyParameters) GoString() string {
  return s.String()
}

type EditDomainPropertyRequestHeader struct {
}

func (s EditDomainPropertyRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s EditDomainPropertyRequestHeader) GoString() string {
  return s.String()
}

type EditDomainPropertyResponseHeader struct {
}

func (s EditDomainPropertyResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s EditDomainPropertyResponseHeader) GoString() string {
  return s.String()
}




type QueryCacheKeyConfigurationRequest struct {
}

func (s QueryCacheKeyConfigurationRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryCacheKeyConfigurationRequest) GoString() string {
  return s.String()
}

type QueryCacheKeyConfigurationResponse struct {
  // {"en":"Domain id", "zh_CN":"域名ID"}
  DomainId *string `json:"domainId,omitempty" xml:"domainId,omitempty" require:"true"`
  // {"en":"Domain name", "zh_CN":"域名"}
  DomainName *string `json:"domainName,omitempty" xml:"domainName,omitempty" require:"true"`
  // {"en":"Custom Cachekey Configuration, parent node
  // 1. When you need to configure the cachekey rules, this must be filled in.
  // 2. Configuration of clearing for <cacheKeyRules/>.", "zh_CN":"配置自定义缓存key功能。
  // 1.需要设置自定义缓存key配置时，此项必填
  // 2.为<cacheKeyRules/>时清空自定义缓存key配置"}
  CacheKeyRules []*QueryCacheKeyConfigurationResponseCacheKeyRules `json:"cacheKeyRules,omitempty" xml:"cacheKeyRules,omitempty" require:"true" type:"Repeated"`
}

func (s QueryCacheKeyConfigurationResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryCacheKeyConfigurationResponse) GoString() string {
  return s.String()
}

func (s *QueryCacheKeyConfigurationResponse) SetDomainId(v string) *QueryCacheKeyConfigurationResponse {
  s.DomainId = &v
  return s
}

func (s *QueryCacheKeyConfigurationResponse) SetDomainName(v string) *QueryCacheKeyConfigurationResponse {
  s.DomainName = &v
  return s
}

func (s *QueryCacheKeyConfigurationResponse) SetCacheKeyRules(v []*QueryCacheKeyConfigurationResponseCacheKeyRules) *QueryCacheKeyConfigurationResponse {
  s.CacheKeyRules = v
  return s
}

type QueryCacheKeyConfigurationResponseCacheKeyRules struct     {
  // {"en":"The url matching mode supports fuzzy regularization. If all matches, the input parameters can be configured as: *", "zh_CN":"url匹配模式，支持正则，如果是全部匹配，入参可以配置为：.*"}
  PathPattern *string `json:"pathPattern,omitempty" xml:"pathPattern,omitempty" require:"true"`
  // {"en":"Specify a uri, such as /test/specifyurl", "zh_CN":"指定具体的url，如/test/specifyurl"}
  SpecifyUrl *string `json:"specifyUrl,omitempty" xml:"specifyUrl,omitempty" require:"true"`
  // {"en":"Whether to match specifyUrl exactly or not, you can select true and false.
  // True:means match exactly. False: means fuzzy match, such as specifyUrl='/test/uri', and  request for /test/uri?p=1 will be matched.", "zh_CN":"是否完全匹配specifyUrl，可选择为true和false。
  // 为true则完全匹配；为false则模糊匹配，如指定/test/uri，请求/test/uri?p=1也会匹配"}
  FullMatch4SpecifyUrl *string `json:"fullMatch4SpecifyUrl,omitempty" xml:"fullMatch4SpecifyUrl,omitempty" require:"true"`
  // {"en":"Specify common types: Select the domain name that requires the cache to be all files or the home page. :
  // E.g:
  // All: all files
  // Homepage: homepage", "zh_CN":"指定常用类型：选择缓存域名的是全部文件还是首页。入参参考值： all：全部文件 homepage：首页"}
  CustomPattern *string `json:"customPattern,omitempty" xml:"customPattern,omitempty" require:"true"`
  // {"en":"File Type: Specify the file type for cache settings.
  // File types include: gif png bmp jpeg jpg html htm shtml mp3 wma flv mp4 wmv zip exe rar css txt ico js swf
  // If you need all types, pass all directly. Multiples are separated by semicolons, and all and specific file types cannot be configured at the same time.", "zh_CN":"文件类型：指定需要缓存的文件类型。 文件类型包括：gif png bmp jpeg jpg html htm shtml mp3 wma flv mp4 wmv zip exe rar css txt ico js swf 如果需要全部类型，则直接传all。多个以分号隔开，all和具体文件类型不能同时配置。"}
  FileType *string `json:"fileType,omitempty" xml:"fileType,omitempty" require:"true"`
  // {"en":"Custom file type: Fill in the appropriate identifiable file type according to your needs outside of the specified file type. Can be used with file-type. If the file-type is also configured, the actual file type is the sum of the two parameters.", "zh_CN":"自定义文件类型：在指定文件类型外根据自身需求，填写适当的可识别文件类型。可以搭配file-type使用。如果file-type也有配置，实际生效的文件类型是两个入参的总和"}
  CustomFileType *string `json:"customFileType,omitempty" xml:"customFileType,omitempty" require:"true"`
  // {"en":"Directory: Specify the directory cache.
  // Enter a legal directory format. Multiple separated by semicolons", "zh_CN":"目录：指定目录缓存。 输入合法的目录格式。多个以英文分号隔开"}
  Directory *string `json:"directory,omitempty" xml:"directory,omitempty" require:"true"`
  // {"en":"Ignore case, the optional value is true or false, true means to ignore case; false means not to ignore case;
  // When adding a new configuration item, the default is true", "zh_CN":"是否忽略大小写：允许值为true和false，默认为忽略"}
  IgnoreCase *string `json:"ignoreCase,omitempty" xml:"ignoreCase,omitempty" require:"true"`
  // {"en":"Header name.
  // Example: If you specify a header as &lsquo;lang', Then, if the value of Lang is consistent, one copy will be cached", "zh_CN":"头部名称
  // 例如：指定头部lang，lang的值一致则缓存一份"}
  HeaderName *string `json:"headerName,omitempty" xml:"headerName,omitempty" require:"true"`
  // {"en":"Parameter Of the specified Header，
  // Example: Specifies the header as 'cookie', parameterOfHeader as 'name'. Then, if the value of name is consistent, one copy will be cached.", "zh_CN":"头部值的参数名，
  // 例如：指定头部Cookie，头部值的参数名为name。则name的值一致则缓存一份。"}
  ParameterOfHeader *string `json:"parameterOfHeader,omitempty" xml:"parameterOfHeader,omitempty" require:"true"`
  // {"en":"Indicates the priority execution order of multiple sets of redirected content by the customer. The higher the number, the higher the priority.
  // When adding a new configuration item, the default is 10", "zh_CN":"优先级，表示客户多组配置的优先执行顺序。数字越大，优先级越高。不传默认为10，不可清空。"}
  Priority *string `json:"priority,omitempty" xml:"priority,omitempty" require:"true"`
  // {"en":"DataId is to indicate a specific group configuration when the client has multiple groups of configurations. dataId can be retrieved through a query interface. Note: A. If dataId is passed, it means that one group of configuration items is specified to be modified, and no other group configuration items need to be modified. B. If multiple groups of configurations are included, some of them are configured with dataId and others are not, then the expression of dataId is used to modify a specific group of configurations, and a new group of configurations is added on the original basis without the expression of dataId. C. If the dataId is not transmitted, it means that the original configuration will be fully covered by this configuration. D. If no configuration parameter is passed, only domain name and secondary label are passed, which means that all configuration of domain name secondary service corresponding to this interface is cleared. E. If there is no specific configuration item in a set of configurations, the dataId must be filled in, and the value is the actual dataId, which means clearing the value of the corresponding dataId configuration item; it is not allowed that there is no specific configuration item or dataId in a set of configurations.", "zh_CN":"配置多组配置时，具体某组配置的id。本功能只支持一组配置。dataId可以通过查询接口获取。 注意： a、如果有传dataId，说明指定修改该组配置项内容； b、如果入参都没有传dataId,表示用本次的配置全量覆盖原先配置； c、如果一组配置没有具体的配置项，则dataId必填，且值为实际存在的dataId，表示清空这个dataId对应配置项的值；不允许一组配置没有具体的配置项也没有dataId。"}
  DataId *int64 `json:"dataId,omitempty" xml:"dataId,omitempty" require:"true"`
}

func (s QueryCacheKeyConfigurationResponseCacheKeyRules) String() string {
  return tea.Prettify(s)
}

func (s QueryCacheKeyConfigurationResponseCacheKeyRules) GoString() string {
  return s.String()
}

func (s *QueryCacheKeyConfigurationResponseCacheKeyRules) SetPathPattern(v string) *QueryCacheKeyConfigurationResponseCacheKeyRules {
  s.PathPattern = &v
  return s
}

func (s *QueryCacheKeyConfigurationResponseCacheKeyRules) SetSpecifyUrl(v string) *QueryCacheKeyConfigurationResponseCacheKeyRules {
  s.SpecifyUrl = &v
  return s
}

func (s *QueryCacheKeyConfigurationResponseCacheKeyRules) SetFullMatch4SpecifyUrl(v string) *QueryCacheKeyConfigurationResponseCacheKeyRules {
  s.FullMatch4SpecifyUrl = &v
  return s
}

func (s *QueryCacheKeyConfigurationResponseCacheKeyRules) SetCustomPattern(v string) *QueryCacheKeyConfigurationResponseCacheKeyRules {
  s.CustomPattern = &v
  return s
}

func (s *QueryCacheKeyConfigurationResponseCacheKeyRules) SetFileType(v string) *QueryCacheKeyConfigurationResponseCacheKeyRules {
  s.FileType = &v
  return s
}

func (s *QueryCacheKeyConfigurationResponseCacheKeyRules) SetCustomFileType(v string) *QueryCacheKeyConfigurationResponseCacheKeyRules {
  s.CustomFileType = &v
  return s
}

func (s *QueryCacheKeyConfigurationResponseCacheKeyRules) SetDirectory(v string) *QueryCacheKeyConfigurationResponseCacheKeyRules {
  s.Directory = &v
  return s
}

func (s *QueryCacheKeyConfigurationResponseCacheKeyRules) SetIgnoreCase(v string) *QueryCacheKeyConfigurationResponseCacheKeyRules {
  s.IgnoreCase = &v
  return s
}

func (s *QueryCacheKeyConfigurationResponseCacheKeyRules) SetHeaderName(v string) *QueryCacheKeyConfigurationResponseCacheKeyRules {
  s.HeaderName = &v
  return s
}

func (s *QueryCacheKeyConfigurationResponseCacheKeyRules) SetParameterOfHeader(v string) *QueryCacheKeyConfigurationResponseCacheKeyRules {
  s.ParameterOfHeader = &v
  return s
}

func (s *QueryCacheKeyConfigurationResponseCacheKeyRules) SetPriority(v string) *QueryCacheKeyConfigurationResponseCacheKeyRules {
  s.Priority = &v
  return s
}

func (s *QueryCacheKeyConfigurationResponseCacheKeyRules) SetDataId(v int64) *QueryCacheKeyConfigurationResponseCacheKeyRules {
  s.DataId = &v
  return s
}

type QueryCacheKeyConfigurationPaths struct {
  // {"en":"Domain name or domain id to query configuration", "zh_CN":"需要查询配置的域名（domainName）或域名id（domainId）"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
}

func (s QueryCacheKeyConfigurationPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryCacheKeyConfigurationPaths) GoString() string {
  return s.String()
}

func (s *QueryCacheKeyConfigurationPaths) SetDomain(v string) *QueryCacheKeyConfigurationPaths {
  s.Domain = &v
  return s
}

type QueryCacheKeyConfigurationParameters struct {
}

func (s QueryCacheKeyConfigurationParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryCacheKeyConfigurationParameters) GoString() string {
  return s.String()
}

type QueryCacheKeyConfigurationRequestHeader struct {
}

func (s QueryCacheKeyConfigurationRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryCacheKeyConfigurationRequestHeader) GoString() string {
  return s.String()
}

type QueryCacheKeyConfigurationResponseHeader struct {
}

func (s QueryCacheKeyConfigurationResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryCacheKeyConfigurationResponseHeader) GoString() string {
  return s.String()
}




type UpdatetimecontrolServiceRequest struct {
  // {"en":"", "zh_CN":"时间戳防盗链设置
  // 注意：
  // 1、时间戳防盗链分为两部分，一部分是防盗链校验，一部分是时间有效性校验。二者都有效，则防盗链通过，否则不通过。
  // 2、防盗链校验：加密算法为md5sum，按照参与MD5计算的参数及组合顺序进行防盗链加密串的计算，对匹配目录下所有文件的url进行防盗链校验，未匹配到的url，则拒绝访问。
  // 3、时间有效性检验：按照年月日时分秒换算的当前时间，与请求url中所带的名文时间相减，判断是否超过设置的上下限（即前后60s内），时间差小于设置上下限的，系统才会给予正常的响应，否则拒绝请求，返回403
  // 4、日志记录没有带加密串的url
  // 6、需要清空时间戳防盗链规则时，可以只传入节点<timestamp-visit-control-rule></timestamp-visit-control-rule>"}
  TimestampVisitControlRule *UpdatetimecontrolServiceRequestTimestampVisitControlRule `json:"timestamp-visit-control-rule,omitempty" xml:"timestamp-visit-control-rule,omitempty" require:"true" type:"Struct"`
}

func (s UpdatetimecontrolServiceRequest) String() string {
  return tea.Prettify(s)
}

func (s UpdatetimecontrolServiceRequest) GoString() string {
  return s.String()
}

func (s *UpdatetimecontrolServiceRequest) SetTimestampVisitControlRule(v *UpdatetimecontrolServiceRequestTimestampVisitControlRule) *UpdatetimecontrolServiceRequest {
  s.TimestampVisitControlRule = v
  return s
}

type UpdatetimecontrolServiceRequestTimestampVisitControlRule struct {
  // {"en":"The url matching mode supports fuzzy regularization. If all matches, the input parameters can be configured as: *.
  // Verify the time stamp of the matched URL for anti-leeching; reject URLs that are not matched.", "zh_CN":"url匹配模式，支持正则，如果是全部匹配，入参可以配置为：.*
  // 对匹配到的URL进行时间戳防盗链验证；未匹配到的URL，则拒绝。"}
  PathPattern *string `json:"path-pattern,omitempty" xml:"path-pattern,omitempty"`
  // {"en":"Exceptional url matching mode, except for some URLs: such as abc.jpg, do not do anti-theft chain function
  // ", "zh_CN":"例外的url匹配模式，某些URL除外：如abc.jpg，不做防盗链"}
  ExceptPathPattern *string `json:"except-path-pattern,omitempty" xml:"except-path-pattern,omitempty"`
  // {"en":"Optional values are: http, https, http;https, noprefix, empty. If it is empty, it defaults to \"http;https\"; if it is noprefix, it means that the protocol prefix of url is not specified, and it only matches according to the regularity of path-pattern. This configuration item only matches with path-pattern. example: 1. Specify protocol-of-path-pattern=https and path-pattern=.* to match all https requests, but not http requests. 2. Specify protocol-of-path-pattern=http;https, path-pattern=.* to match all http and https requests. 3. Specify protocol-of-path-pattern=noprefix and path-pattern=^http://[^/]+/.* to match all http requests but not https requests.",
  //                               "zh_CN":"可选值为: http、https、http;https、noprefix、空。为空默认为\"http;https\"；为noprefix表示不指定url的协议前缀，仅按path-pattern的正则匹配。本配置项只与path-pattern（url匹配模式）结合匹配。 例子： 1、指定protocol-of-path-pattern=https，path-pattern=.*，则匹配所有https的请求，不匹配http的请求。 2、指定protocol-of-path-pattern=http;https，path-pattern=.*，则匹配所有http和https的请求。 3、指定protocol-of-path-pattern=noprefix，path-pattern=^http://[^/]+/.*，则匹配所有http的请求，不匹配https的请求。"}
  ProtocolOfPathPattern *string `json:"protocol-of-path-pattern,omitempty" xml:"protocol-of-path-pattern,omitempty"`
  // {"en":"Directory, multiple separated by English semicolons. Perform timestamp anti-leech verification for the matched directory; reject the unmatched directory. mutually exclusive with path-pattern.", "zh_CN":"目录，多个以英文分号隔开。对于匹配到的目录进行时间戳防盗链验证；未匹配到的则拒绝。和path-pattern互斥。"}
  Directory *string `json:"directory,omitempty" xml:"directory,omitempty"`
  // {"en":"Exceptional IP, supports input of IP or IP range, separate IP ranges with semicolons (;), such as 1.1.1.0/24;2.2.2.2, some IP exceptions, no anti-leeching.
  // ", "zh_CN":"例外的IP，支持输入IP或IP段，IP段之间用分号(;)隔开，如1.1.1.0/24;2.2.2.2，某些IP例外，不做防盗链"}
  AllowedIps *string `json:"allowed-ips,omitempty" xml:"allowed-ips,omitempty"`
  // {"en":"Whether to remove / from $uri in anti-leech, the optional values are true or false, and the default is false, that is, it contains /. For example: http://www.test.com/1.flv, then $uri is /1.flv by default, if ignore-uri-slash is true, then $uri is 1.flv", "zh_CN":"防盗链中的$uri是否去掉/，可选值为true、false，默认为false，即包含/。 例如： http://www.test.com/1.flv，则$uri默认为/1.flv,若ignore-uri-slash为true，则$uri为1.flv"}
  IgnoreUriSlash *string `json:"ignore-uri-slash,omitempty" xml:"ignore-uri-slash,omitempty"`
  // {"en":"Whether key and time are allowed to be interchanged, the optional values are true and false, true is allowed, false is not allowed. By default, the order of key parameters and time parameters must strictly follow the order required by the authentication mode, that is, the default key and time cannot be interchanged. If true is selected, the key and time positions can be interchanged and the authentication succeeds.", "zh_CN":"key与time是否允许互换，可选值为true和false，true则允许，false则不允许。默认情况下，密钥参数和时间参数的顺序，必须严格参照鉴权模式要求的顺序，即，默认key和time不能互换位置。如选择true，key和time位置可以互换并鉴权成功。"}
  IgnoreKeyAndTimePosition *string `json:"ignore-key-and-time-position,omitempty" xml:"ignore-key-and-time-position,omitempty"`
  // {"en":"Encryption Algorithm. Support value: md5sum
  // ", "zh_CN":"加密算法
  // 当前支持入参：md5sum"}
  EncryptMethod *string `json:"encrypt-method,omitempty" xml:"encrypt-method,omitempty"`
  // {"en":"The anti-leech chain generation method, the parameters involved in the MD5 calculation and the combination sequence only support the following parameters: $uri: A string between domain and question mark, special values are configured in the input parameter <uri-select> For example http://cdn.example.com/v0/test.dat?k=v, the URI is /v0/test.dat $ourkey: secret key, the actual secret key is configured in the input parameter <secret-key> $time: time string $spec_name: file name For example: http://cdn.example.com/v0/test.dat?k=v, the file name is test.dat $args: the value of a specific key in QUERY_STRING Example: <cipher-combination>$uri$ourkey$time$args{k}</cipher-combination> Notice: 1. K in $args{k} only allows A-Z uppercase and lowercase letters, numbers, underscores, and bars 2. Except for $args, other parameters are only allowed to appear once 3. Each parameter can be spliced according to any combination order
  // ", "zh_CN":"防盗链生成方式，参与MD5计算的参数及组合顺序，仅支持传入以下参数：
  // $uri：介于domain和问号之间的字符串
  // 例如http://cdn.example.com/v0/test.dat?k=v，则URI为/v0/test.dat
  // $ourkey：秘钥，实际秘钥在入参<secret-key>中配置
  // $time：时间串
  // $spec_name：文件名
  // 例如：http://cdn.example.com/v0/test.dat?k=v，文件名为 test.dat
  // $args：QUERY_STRING中的某个具体key的值
  // 示例：<cipher-combination>$uri$ourkey$time$args{k}</cipher-combination>
  // 注意：
  // 1、$args{k}中的K只允许A-Z大小写字母、数字、下划线、横杠
  // 2、除了$args外，其它参数只允许出现1次
  // 3、可以按照任意组合顺序，拼接各个参数"}
  CipherCombination *string `json:"cipher-combination,omitempty" xml:"cipher-combination,omitempty"`
  // {"en":"The key of the anti-leech encryption string, only one key is allowed to be passed in Example: <secret-key>abcdef</secret-key> Notice: 1. For the secret key agreed with the customer, a value in <multiple-secret-keys> can be equal to the corresponding value of <secret-key> 2. The value of $ourkey mainly comes from the configuration of <secret-key> a) If <secret-key> is not a value in multiple-secret-keys>, the value of $ourkey is <multiple-secret-keys> b) <secret-key> is a value in multiple-secret-keys>, then the value of $ourkey is <secret-key> c) If <secret-key> is not passed or is empty, then the value of $ourkey is <multiple-secret-keys>
  // ", "zh_CN":"防盗链加密串的秘钥，只允许传入一个秘钥
  // 示例：<secret-key>abcdef</secret-key>
  // 注意：
  // 1、与客户约定好的秘钥，入参<multiple-secret-keys>中的某个值可以等于<secret-key>对应的值
  // 2、$ourkey值主要是来自<secret-key>的配置
  // a)如果<secret-key>不是multiple-secret-keys>中的某个值，则$ourkey值取<multiple-secret-keys>
  // b）<secret-key>是multiple-secret-keys>中的某个值，则$ourkey值取<secret-key>
  // c）如果<secret-key>没有传或者空值，则$ourkey值取<multiple-secret-keys>"}
  SecretKey *string `json:"secret-key,omitempty" xml:"secret-key,omitempty"`
  // {"en":"Parameter name of the anti-leech string Example: <cipher-param>keyname</cipher-param> Notice: 1. If the anti-leech encryption string is in the parameter after the question mark in the url, the parameter name of the anti-leech encryption string is determined by the configuration of <cipher-param>; 2. If <cipher-param> is empty, the key is used as the parameter name in the URL of the default request
  // ", "zh_CN":"防盗链串的参数名称
  // 示例：<cipher-param>keyname</cipher-param>
  // 注意：
  // 1、如果防盗链加密串是在url问号后的参数中，则防盗链加密串的参数名由<cipher-param>的配置决定；
  // 2、如果<cipher-param>为空，则默认请求的url中使用key作为参数名"}
  CipherParam *string `json:"cipher-param,omitempty" xml:"cipher-param,omitempty"`
  // {"en":"Parameter name for time string Example: <time-param>tname</time-param> Notice: 1. If the anti-leech time string is placed in the parameter after the url question mark, the parameter name of the anti-leech time string is determined by the configuration of <time-param>; 2. If <time-param> is empty, time is used as the parameter name in the URL of the default request
  // ", "zh_CN":"时间串的参数名称
  // 示例：<time-param>tname</time-param>
  // 注意：
  // 1、如果防盗链时间串是放在url问号后面的参数中，则防盗链时间串的参数名由<time-param>的配置决定；
  // 2、如果<time-param>为空，则默认请求的url中使用time作为参数名"}
  TimeParam *string `json:"time-param,omitempty" xml:"time-param,omitempty"`
  // {"en":"The lower limit of the expiration time of the anti-leech chain Example: <lower-limit-expiry-time>200</lower-limit-expiry-time> The configuration methods corresponding to various scenarios are as follows: Notice: 1. If the timestamp carried in the request URL is the generation time of the URL, the expiration time needs to be added with the effective duration, that is, <lower-limit-expiry-time> and <upper-limit-expiry-time> are configured as the effective duration . 2. If the timestamp carried in the URL is the expiration time, it can be set to zero.
  // ", "zh_CN":"防盗链串的过期时间下限
  // 示例：
  // <lower-limit-expiry-time>200</lower-limit-expiry-time>
  // 对应各种场景的配置方式如下：
  // 注意：
  // 1、请求URL中携带的时间戳如果是URL的生成时间，需要加上有效时长才是过期时间，即<lower-limit-expiry-time>和<upper-limit-expiry-time>配置为有效时长。
  // 2、如果URL携带的时间戳是过期时间，则可以配成零。"}
  LowerLimitExpiryTime *int `json:"lower-limit-expiry-time,omitempty" xml:"lower-limit-expiry-time,omitempty"`
  // {"en":"The upper limit of the expiration time of the anti-leech chain Example: <upper-limit-expiry-time>5000</upper-limit-expiry-time> The configuration methods corresponding to various scenarios are as follows: Notice: 1. If the timestamp carried in the request URL is the generation time of the URL, the expiration time needs to be added with the effective duration, that is, <lower-limit-expiry-time> and <upper-limit-expiry-time> are configured as the effective duration . 2. If the timestamp carried in the URL is the expiration time, it can be set to zero.
  // ", "zh_CN":"防盗链串的过期时间上限
  // 示例：
  // <upper-limit-expiry-time>5000</upper-limit-expiry-time>
  // 对应各种场景的配置方式如下：
  // 注意：
  // 1、请求URL中携带的时间戳如果是URL的生成时间，需要加上有效时长才是过期时间，即<lower-limit-expiry-time>和<upper-limit-expiry-time>配置为有效时长。
  // 2、如果URL携带的时间戳是过期时间，则可以配成零。"}
  UpperLimitExpiryTime *int `json:"upper-limit-expiry-time,omitempty" xml:"upper-limit-expiry-time,omitempty"`
  // {"en":"Anti-leech encrypted string, multiple encrypted strings are supported, and multiple encrypted strings are separated by semicolons (;) Example: <multiple-secret-keys>abcdef;uvwxyz</multiple-secret-keys> Notice: 1. Support setting multiple keys for url anti-leech. Support customers to modify the key at will, and achieve seamless switching. The anti-theft chain level is higher. 2. As long as the key of the requested url is consistent with the key calculated from any of the encrypted strings, the verification will pass
  // ", "zh_CN":"防盗链加密串，支持多个加密串，多个加密串以分号(;)隔开
  // 示例：<multiple-secret-keys>abcdef;uvwxyz</multiple-secret-keys>
  // 注意：
  // 1、支持对url 防盗链设置多个密钥。支持客户任意修改密钥，并做到无缝切换。防盗链等级更高。
  // 2、请求url的key只要跟其中任意一个加密串算出来的key一致就验证通过"}
  MultipleSecretKeys *string `json:"multiple-secret-keys,omitempty" xml:"multiple-secret-keys,omitempty"`
  // {"en":"The time format of the anti-leech encryption string, multiple choices can be selected, separated by semicolons (;) Year|Month|Day|Hour|Minute|Second|UNIX timestamp|Hexadecimal timestamp|Timestamp in milliseconds: 1Y;2m;3d;4H;5M;6S;7s;8x Example: <time-format>1Y;2m;3d;4H;5M;6S;7s;8x</time-format> or <time-format>1Y;2m;3d;4H;5M;7s</time-format > Notice: 1. Must be English letters and numbers 2. Each value can only appear once
  // ", "zh_CN":"防盗链加密串时间格式，可多选，以分号(;)分隔
  // 年|月|日|时|分|秒|UNIX时间戳|16进制时间戳：1Y;2m;3d;4H;5M;6S;7s;8x
  // 示例：<time-format>1Y;2m;3d;4H;5M;6S;7s;8x</time-format> 或 <time-format>1Y;2m;3d;4H;5M;7s</time-format>
  // 注意：
  // 1、必须是英文字母和数字
  // 2、每个值只能出现1次
  // 3、如果配置的是16进制时间戳，则需要同时入参：UNIX时间戳|16进制时间戳"}
  TimeFormat *string `json:"time-format,omitempty" xml:"time-format,omitempty"`
  // {"en":"The anti-leech request url format supports two anti-leech methods, that is, the encrypted string and timestamp are placed after \"?\" or the encrypted string and timestamp are placed after \"host\". The parameters supported by the url format are as follows: $domain: domain name $uri: the url part that does not contain the domain name $key: MD5 value of anti-leech encrypted string $time: Anti-leech time string $args: the QUERY_STRING parameter after the question mark Example: The following request url format is supported, which can be replaced with https://. The url request protocol is based on actual use. If you don't know how to configure it correctly, please ask customer technical support for assistance; the parameter name carrying two values of encrypted string and time string\" keyname\" and \"tname\", which can be replaced by the actual parameter names used <request-url-style>http://$domain/$key/$time/$uri?$args</request-url-style> <request-url-style>http://$domain/$time/$key/$uri?$args</request-url-style> http://$domain/$uri?auth_key=$key <request-url-style>http://$domain/$uri?keyname=$key&tname=$time</request-url-style> <request-url-style>http://$domain/$uri?$args&keyname=$key&tname=$time</request-url-style> <request-url-style>http://$domain/$uri?keyname=$key&tname=$time&$args</request-url-style> <request-url-style>http://$domain/$uri?$args&keyname=$key&tname=$time&$args</request-url-style> <request-url-style>http://$domain/$uri?tname=$time&keyname=$key</request-url-style> <request-url-style>http://$domain/$uri?$args&tname=$time&keyname=$key</request-url-style> <request-url-style>http://$domain/$uri?tname=$time&keyname=$key&$args</request-url-style> <request-url-style>http://$domain/$uri?$args&tname=$time&keyname=$key&$args</request-url-style> Notice: 1. The input url must start with \"http/https\" 2. If the encrypted string and timestamp are placed after \"?\", the keyname and tname must be consistent with the configured values of <cipher-param> and <time-param> 3. If there is no configuration value for <cipher-param> and <time-param>, the parameter name corresponding to $key defaults to key, and the parameter name corresponding to $time defaults to time 4. If the anti-leech encryption string and time string are in the parameters behind the question mark in the url, the \"keyname\" and \"tname\" in the url correspond to the anti-leech string and time string parameter names configured in cipher-param and time-param.", "zh_CN":"防盗链请求url格式，支持两种防盗链方式，即加密串和时间戳放到“?”后面或者是加密串和时间戳放到“host”后面，url格式支持的参数如下：
  // $domain：域名
  // $uri：不包含域名的url部分
  // $key：防盗链加密串的MD5值
  // $time：防盗链时间串
  // $args：问号后的QUERY_STRING参数
  // 示例：支持以下请求url格式，可替换为https://，url请求协议根据实际使用，如不知道如何正确配置，请找客户技术支持协助；携带加密串和时间串两个值的参数名“keyname”和“tname”，可替换为实际使用的参数名
  // <request-url-style>http://$domain/$key/$time/$uri?$args</request-url-style>
  // <request-url-style>http://$domain/$time/$key/$uri?$args</request-url-style>
  //  http://$domain/$uri?auth_key=$key 
  // <request-url-style>http://$domain/$uri?keyname=$key&tname=$time</request-url-style>
  // <request-url-style>http://$domain/$uri?$args&keyname=$key&tname=$time</request-url-style>
  // <request-url-style>http://$domain/$uri?keyname=$key&tname=$time&$args</request-url-style>
  // <request-url-style>http://$domain/$uri?$args&keyname=$key&tname=$time&$args</request-url-style>
  // <request-url-style>http://$domain/$uri?tname=$time&keyname=$key</request-url-style>
  // <request-url-style>http://$domain/$uri?$args&tname=$time&keyname=$key</request-url-style>
  // <request-url-style>http://$domain/$uri?tname=$time&keyname=$key&$args</request-url-style>
  // <request-url-style>http://$domain/$uri?$args&tname=$time&keyname=$key&$args</request-url-style>
  // 注意：
  // 1、输入的url必须以“http/https”开头
  // 2、如果加密串和时间戳是放到“?”后面时，keyname和tname必须跟<cipher-param>和<time-param>配置的值一致
  // 3、如果<cipher-param>和<time-param>没有配置值，则$key对应的参数名默认为key，$time对应的参数名默认为time
  // 4、如果防盗链加密串和时间串在url问号后面的参数中，url中的“keyname”和“tname”，对应的是cipher-param和 time-param配置的防盗链串和时间串参数名称。"}
  RequestUrlStyle *string `json:"request-url-style,omitempty" xml:"request-url-style,omitempty"`
  // {"en":"Anti-leech back-to-source method, optional values: 1 (use unencrypted url to go back to the source), 2 (use the customer request to return to the source with encrypted string url) Example: <dst-style>1<dst-style> Notice: 1. If the URL format is: http://www.xxx.com/md5/time/uri? parameter, please contact your technical support.
  // ", "zh_CN":"防盗链回源方式，可选值：1（使用未加密url回源）、2（使用客户请求带加密串url回源）
  // 示例：<dst-style>1<dst-style>
  // 注意：
  // 1、如果URL格式是：http://www.xxx.com/md5/time/uri?参数，则需要下工单给对应客服，让客服在父配置去掉时间戳格式再缓存。"}
  DstStyle *int `json:"dst-style,omitempty" xml:"dst-style,omitempty"`
  // {"en":"Logging original url, optional values: true (logging original url), false (do not enable logging original url)
  // ", "zh_CN":" 日志记录原始url，可选值：true（日志记录原始url）、false（不开启日志记录原始url）"}
  LogFormat *string `json:"log-format,omitempty" xml:"log-format,omitempty"`
  // {"en":"Used to configure the name of the key in the url Example: <url-key>auth_key</url-key>
  // ", "zh_CN":"用于配置获取url中的key的名称
  // 示例： <url-key>auth_key</url-key>"}
  UrlKey *string `json:"url-key,omitempty" xml:"url-key,omitempty"`
}

func (s UpdatetimecontrolServiceRequestTimestampVisitControlRule) String() string {
  return tea.Prettify(s)
}

func (s UpdatetimecontrolServiceRequestTimestampVisitControlRule) GoString() string {
  return s.String()
}

func (s *UpdatetimecontrolServiceRequestTimestampVisitControlRule) SetPathPattern(v string) *UpdatetimecontrolServiceRequestTimestampVisitControlRule {
  s.PathPattern = &v
  return s
}

func (s *UpdatetimecontrolServiceRequestTimestampVisitControlRule) SetExceptPathPattern(v string) *UpdatetimecontrolServiceRequestTimestampVisitControlRule {
  s.ExceptPathPattern = &v
  return s
}

func (s *UpdatetimecontrolServiceRequestTimestampVisitControlRule) SetProtocolOfPathPattern(v string) *UpdatetimecontrolServiceRequestTimestampVisitControlRule {
  s.ProtocolOfPathPattern = &v
  return s
}

func (s *UpdatetimecontrolServiceRequestTimestampVisitControlRule) SetDirectory(v string) *UpdatetimecontrolServiceRequestTimestampVisitControlRule {
  s.Directory = &v
  return s
}

func (s *UpdatetimecontrolServiceRequestTimestampVisitControlRule) SetAllowedIps(v string) *UpdatetimecontrolServiceRequestTimestampVisitControlRule {
  s.AllowedIps = &v
  return s
}

func (s *UpdatetimecontrolServiceRequestTimestampVisitControlRule) SetIgnoreUriSlash(v string) *UpdatetimecontrolServiceRequestTimestampVisitControlRule {
  s.IgnoreUriSlash = &v
  return s
}

func (s *UpdatetimecontrolServiceRequestTimestampVisitControlRule) SetIgnoreKeyAndTimePosition(v string) *UpdatetimecontrolServiceRequestTimestampVisitControlRule {
  s.IgnoreKeyAndTimePosition = &v
  return s
}

func (s *UpdatetimecontrolServiceRequestTimestampVisitControlRule) SetEncryptMethod(v string) *UpdatetimecontrolServiceRequestTimestampVisitControlRule {
  s.EncryptMethod = &v
  return s
}

func (s *UpdatetimecontrolServiceRequestTimestampVisitControlRule) SetCipherCombination(v string) *UpdatetimecontrolServiceRequestTimestampVisitControlRule {
  s.CipherCombination = &v
  return s
}

func (s *UpdatetimecontrolServiceRequestTimestampVisitControlRule) SetSecretKey(v string) *UpdatetimecontrolServiceRequestTimestampVisitControlRule {
  s.SecretKey = &v
  return s
}

func (s *UpdatetimecontrolServiceRequestTimestampVisitControlRule) SetCipherParam(v string) *UpdatetimecontrolServiceRequestTimestampVisitControlRule {
  s.CipherParam = &v
  return s
}

func (s *UpdatetimecontrolServiceRequestTimestampVisitControlRule) SetTimeParam(v string) *UpdatetimecontrolServiceRequestTimestampVisitControlRule {
  s.TimeParam = &v
  return s
}

func (s *UpdatetimecontrolServiceRequestTimestampVisitControlRule) SetLowerLimitExpiryTime(v int) *UpdatetimecontrolServiceRequestTimestampVisitControlRule {
  s.LowerLimitExpiryTime = &v
  return s
}

func (s *UpdatetimecontrolServiceRequestTimestampVisitControlRule) SetUpperLimitExpiryTime(v int) *UpdatetimecontrolServiceRequestTimestampVisitControlRule {
  s.UpperLimitExpiryTime = &v
  return s
}

func (s *UpdatetimecontrolServiceRequestTimestampVisitControlRule) SetMultipleSecretKeys(v string) *UpdatetimecontrolServiceRequestTimestampVisitControlRule {
  s.MultipleSecretKeys = &v
  return s
}

func (s *UpdatetimecontrolServiceRequestTimestampVisitControlRule) SetTimeFormat(v string) *UpdatetimecontrolServiceRequestTimestampVisitControlRule {
  s.TimeFormat = &v
  return s
}

func (s *UpdatetimecontrolServiceRequestTimestampVisitControlRule) SetRequestUrlStyle(v string) *UpdatetimecontrolServiceRequestTimestampVisitControlRule {
  s.RequestUrlStyle = &v
  return s
}

func (s *UpdatetimecontrolServiceRequestTimestampVisitControlRule) SetDstStyle(v int) *UpdatetimecontrolServiceRequestTimestampVisitControlRule {
  s.DstStyle = &v
  return s
}

func (s *UpdatetimecontrolServiceRequestTimestampVisitControlRule) SetLogFormat(v string) *UpdatetimecontrolServiceRequestTimestampVisitControlRule {
  s.LogFormat = &v
  return s
}

func (s *UpdatetimecontrolServiceRequestTimestampVisitControlRule) SetUrlKey(v string) *UpdatetimecontrolServiceRequestTimestampVisitControlRule {
  s.UrlKey = &v
  return s
}

type UpdatetimecontrolServiceResponse struct {
  // {"en":"httpstatus=202; Indicates that the new domain API was successfully invoked, and the current deployment of the new domain can be viewed using x-cnc-request-id in the header", "zh_CN":"httpstatus=202;   表示成功调用新增域名接口，可使用header中的x-cnc-request-id查看当前新增域名的部署情况"}
  HttpStatus *int `json:"http status code,omitempty" xml:"http status code,omitempty" require:"true"`
  // {"en":"Uniquely identified id for querying tasks per request (for all API)", "zh_CN":"唯一标示的id，用于查询每次请求的任务 （适用全部接口）"}
  XCncRequestId *string `json:"x-cnc-request-id,omitempty" xml:"x-cnc-request-id,omitempty" require:"true"`
  // {"en":"Error code. It pops up when the HTTPStatus is not 202, and shows the revoking error type of the current request.", "zh_CN":"错误代码，当HTTPStatus不为202时出现，表示当前请求调用的错误类型"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response message, and shows as success when it succeeds.", "zh_CN":"响应信息，成功时为success"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s UpdatetimecontrolServiceResponse) String() string {
  return tea.Prettify(s)
}

func (s UpdatetimecontrolServiceResponse) GoString() string {
  return s.String()
}

func (s *UpdatetimecontrolServiceResponse) SetHttpStatus(v int) *UpdatetimecontrolServiceResponse {
  s.HttpStatus = &v
  return s
}

func (s *UpdatetimecontrolServiceResponse) SetXCncRequestId(v string) *UpdatetimecontrolServiceResponse {
  s.XCncRequestId = &v
  return s
}

func (s *UpdatetimecontrolServiceResponse) SetCode(v string) *UpdatetimecontrolServiceResponse {
  s.Code = &v
  return s
}

func (s *UpdatetimecontrolServiceResponse) SetMessage(v string) *UpdatetimecontrolServiceResponse {
  s.Message = &v
  return s
}

type UpdatetimecontrolServicePaths struct {
  // {"en":"The domain whoes need query config.", "zh_CN":"需要查询配置的域名或域名id"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
}

func (s UpdatetimecontrolServicePaths) String() string {
  return tea.Prettify(s)
}

func (s UpdatetimecontrolServicePaths) GoString() string {
  return s.String()
}

func (s *UpdatetimecontrolServicePaths) SetDomain(v string) *UpdatetimecontrolServicePaths {
  s.Domain = &v
  return s
}

type UpdatetimecontrolServiceParameters struct {
}

func (s UpdatetimecontrolServiceParameters) String() string {
  return tea.Prettify(s)
}

func (s UpdatetimecontrolServiceParameters) GoString() string {
  return s.String()
}

type UpdatetimecontrolServiceRequestHeader struct {
}

func (s UpdatetimecontrolServiceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s UpdatetimecontrolServiceRequestHeader) GoString() string {
  return s.String()
}

type UpdatetimecontrolServiceResponseHeader struct {
}

func (s UpdatetimecontrolServiceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s UpdatetimecontrolServiceResponseHeader) GoString() string {
  return s.String()
}




type EditAntiHotlinkingConfigRequest struct {
  // {"en":"Anti-theft chain configuration
  // note:
  // 1. When you need to cancel the anti-theft chain configuration settings, you can pass in the empty node <cache-time-behaviors></cache-time-behaviors>.
  // 2. When it is necessary to set the anti-theft chain configuration, this item is required.", "zh_CN":"防盗链配置
  // 注意：
  // 1. 需要取消防盗链配置设置时，可以传入空节点<cache-time-behaviors></cache-time-behaviors>。
  // 2. 表示需要设置防盗链配置时，此项必填"}
  VisitControlRules []*EditAntiHotlinkingConfigRequestVisitControlRules `json:"visit-control-rules,omitempty" xml:"visit-control-rules,omitempty" require:"true" type:"Repeated"`
}

func (s EditAntiHotlinkingConfigRequest) String() string {
  return tea.Prettify(s)
}

func (s EditAntiHotlinkingConfigRequest) GoString() string {
  return s.String()
}

func (s *EditAntiHotlinkingConfigRequest) SetVisitControlRules(v []*EditAntiHotlinkingConfigRequestVisitControlRules) *EditAntiHotlinkingConfigRequest {
  s.VisitControlRules = v
  return s
}

type EditAntiHotlinkingConfigRequestVisitControlRules struct     {
  // {"en":"When configuring multiple configuration sets, the specific configuration set's ID. The data-id can be obtained through the query interface. Note: a. If data-id is provided, it indicates the modification of a specific Configuration Item in one of the configuration sets. No modification is needed for other configuration sets. b. If multiple configuration sets are provided as input, and some have data-id while others do not, then those with data-id represent modifications to specific configuration sets, whereas those without data-id represent new configurations added on top of existing ones. c. If none of the inputs have data-id, it means the current configuration completely overrides the previous configuration. d. If no configuration parameters are provided and only the domain and secondary tag are transmitted, it indicates clearing all configurations corresponding to the domain's secondary service for this interface. e. If a configuration set has no specific Configuration Item, then data-id is required with an actual existing data-id value, indicating the clearing of the Configuration Item corresponding to this data-id. A configuration set with no specific Configuration Item and no data-id is not allowed.", "zh_CN":"配置多组配置时，具体某组配置的id。data-id可以通过查询接口获取。 注意： a、如果有传data-id，说明指定修改其中一组配置项内容，不需求修改其他组配置内容不需要入参；  b、如果入参多组配置，其中有些组配置有传data-id，有些没有传，则有传data-id的表示修改具体某组配置，没有传data-id的表示在原来基础上新增一组配置；  c、如果入参都没有传data-id,表示用本次的配置全量覆盖原先配置；  d、如果入参没有传任何配置项参数，只传了域名和二级标签，表示清空这个接口对应域名二级服务所有配置；  e、如果一组配置没有具体的配置项，则data-id必填，且值为实际存在的data-id，表示清空这个data-id对应配置项的值；不允许一组配置没有具体的配置项也没有data-id。"}
  DataId *int64 `json:"data-id,omitempty" xml:"data-id,omitempty"`
  // {"en":"The url matching mode supports regularization. If all matches, the input parameters can be configured as: .*", "zh_CN":"url匹配模式，支持正则，如果是全部匹配，入参可以配置为：.*"}
  PathPattern *string `json:"path-pattern,omitempty" xml:"path-pattern,omitempty"`
  // {"en":"Exceptional url matching mode, except for some URLs: such as abc.jpg, do not do anti-theft chain function
  // E.g: ^https?://[^/]+/.*\.m3u8", "zh_CN":"例外的url匹配模式，某些URL除外：如abc.jpg，不做防盗链功能
  // 客户入参参考：^https?://[^/]+/.*\.m3u8"}
  ExceptPathPattern *string `json:"except-path-pattern,omitempty" xml:"except-path-pattern,omitempty"`
  // {"en":"Specify common types: Select the domain name that requires the anti-theft chain to be all files or the home page. :
  // E.g:
  // All: all files
  // Homepage: homepage", "zh_CN":"指定常用类型：选择需要防盗链的域名是全部文件还是首页。入参参考值：
  // all：全部文件
  // homepage：首页"}
  CustomPattern *string `json:"custom-pattern,omitempty" xml:"custom-pattern,omitempty"`
  // {"en":"File Type: Specify the file type for anti-theft chain settings.
  // File types include: gif png bmp jpeg jpg html htm shtml mp3 wma flv mp4 wmv zip exe rar css txt ico js swf
  // If you need all types, pass all directly. Multiples are separated by semicolons, and all and specific file types cannot be configured at the same time.", "zh_CN":"文件类型：指定文件类型进行防盗链设置。
  // 文件类型包括：gif png bmp jpeg jpg html htm shtml mp3 wma flv mp4 wmv zip exe rar css txt ico js swf
  // 如果需要全部类型，则直接传all。多个以分号隔开，all和具体文件类型不能同时配置。"}
  FileType *string `json:"file-type,omitempty" xml:"file-type,omitempty"`
  // {"en":"Custom file type: Fill in the appropriate identifiable file type according to your needs outside of the specified file type. Can be used with file-type. If the file-type is also configured, the actual file type is the sum of the two parameters.", "zh_CN":"自定义文件类型：在指定文件类型外根据自身需求，填写适当的可识别文件类型。可以搭配file-type使用。如果file-type也有配置，实际生效的文件类型是两个入参的总和"}
  CustomFileType *string `json:"custom-file-type,omitempty" xml:"custom-file-type,omitempty"`
  // {"en":"Specify URL cache: Specify url according to requirements for anti-theft chain setting
  // INS format does not support URI format with http(s)://", "zh_CN":"指定URL缓存：根据需求指定url进行防盗链设置
  // 入参不支持含http(s):// 开头的URI格式"}
  SpecifyUrlPattern *string `json:"specify-url-pattern,omitempty" xml:"specify-url-pattern,omitempty"`
  // {"en":"Directory: Specify the directory for anti-theft chain settings
  // Enter a legal directory format. Multiple separated by semicolons", "zh_CN":"目录：指定目录进行防盗链设置
  // 输入合法的目录格式。多个以英文分号隔开"}
  Directory *string `json:"directory,omitempty" xml:"directory,omitempty"`
  // {"en":"Exception file type: Specify the file type that does not require anti-theft chain function
  // File types include: gif png bmp jpeg jpg html htm shtml mp3 wma flv mp4 wmv zip exe rar css txt ico js swf
  // If you need all types, pass all directly. Multiple separated by semicolons, all and specific file types cannot be configured at the same time
  // If file-type=all, except-file-type=all means that the task file type is not matched.", "zh_CN":"例外的文件类型：指定不需要进行防盗链功能的文件类型
  // 文件类型包括：gif png bmp jpeg jpg html htm shtml mp3 wma flv mp4 wmv zip exe rar css txt ico js swf
  // 如果需要全部类型，则直接传all。多个以分号隔开，all和具体文件类型不能同时配置
  // 如果file-type=all,except-file-type=all 则表示不匹配任务文件类型"}
  ExceptFileType *string `json:"except-file-type,omitempty" xml:"except-file-type,omitempty"`
  // {"en":"Exceptional custom file types: Fill in the appropriate identifiable file types based on your needs, outside of the specified file type. Can be used with the exception-file-type. If the except-file-type is also configured, the actual file type is the sum of the two parameters.", "zh_CN":"例外的自定义文件类型：在指定文件类型外根据自身需求，填写适当的可识别文件类型。可以搭配except-file-type使用。如果except-file-type也有配置，实际生效的文件类型是两个入参的总和"}
  ExceptCustomFileType *string `json:"except-custom-file-type,omitempty" xml:"except-custom-file-type,omitempty"`
  // {"en":"Exceptional directory: Specify a directory that does not require anti-theft chain settings
  // Enter a legal directory format. Multiple separated by semicolons", "zh_CN":"例外的目录：指定不需要进行进行防盗链设置的目录
  // 输入合法的目录格式。多个以英文分号隔开"}
  ExceptDirectory *string `json:"except-directory,omitempty" xml:"except-directory,omitempty"`
  // {"en":"control direction. Available values: 403 and 302
  // 1) 403 means to return a specific error status code to reject the service (the default mode, the status code can be specified, generally 403).
  // 2) 302 means to return 302 the redirect url of the Found, the redirected url can be specified. If pass 302, rewrite-to is required", "zh_CN":"控制方向。可选值：403和302
  // 1） 403表示返回特定的错误状态码来拒绝服务（默认方式，状态码可以指定，一般为403）。
  // 2） 302表示返回302 Found的重定向url，重定向的url可以指定。如果传302，rewrite-to必填"}
  ControlAction *string `json:"control-action,omitempty" xml:"control-action,omitempty"`
  // {"en":"Specify the url after the 302 jump. This field is required if the control-action value is 302.", "zh_CN":"指定302跳转后的url。如果control-action值为302，此项必填，值需为具体调整的url，不支持正则。如果control-action值为403，此项填不需要输入，值无效"}
  RewriteTo *string `json:"rewrite-to,omitempty" xml:"rewrite-to,omitempty"`
  // {"en":"Indicates the priority execution order of multiple sets of redirected content by the customer. The higher the number, the higher the priority.
  // When adding a new configuration item, the default is 10", "zh_CN":"表示客户多组重定向内容的优先执行顺序。数字越大，优先级越高。
  // 新增配置项时，不传默认为 10"}
  Priority *string `json:"priority,omitempty" xml:"priority,omitempty"`
  // {"en":"Exceptional Request Method.Multiple entires separated by ;", "zh_CN":"例外的请求方法。多个以;隔开"}
  ExceptionalRequest *string `json:"exceptional-request,omitempty" xml:"exceptional-request,omitempty"`
  // {"en":"Only true and false are allowed to be filled in ignoredcase.", "zh_CN":"是否忽略大小写。只允许填写true或false"}
  IgnoredCase *string `json:"ignored-case,omitempty" xml:"ignored-case,omitempty"`
  // {"en":"Identify IP black and white list anti-theft chain
  // note:
  // 1. a set of black and white list anti-theft chain, only one set under a data-id
  // 2. When the air interface label indicates the exception of the IP segment configuration and the forbidden IP segment configuration.", "zh_CN":"标识IP黑白名单防盗链
  // 注意：
  // 1. 表示一组黑白名单防盗链，一个data-id下只能一组
  // 2. 当传空标签表示清楚例外的IP段配置和禁止的IP段配置。"}
  IpControlRule *EditAntiHotlinkingConfigRequestVisitControlRulesIpControlRule `json:"ip-control-rule,omitempty" xml:"ip-control-rule,omitempty" type:"Struct"`
  // {"en":"Identify referer anti-theft chain
  // Note:
  // 1. Represents a set of referer security chains, and a single data-id can only have one set under one
  // 2. when the empty tag means to clear referer security chain
  // 3. legal refer, ( legal domain name, legal URL), illegal refer, ( illegal domain name, illegal URL) these four, a data-id can only configure one or all empty under one data-id", "zh_CN":"标识referer防盗链
  // 注意：
  // 1. 表示一组referer防盗链，一个data-id下只能一组
  // 2. 当传空标签表示清除referer防盗链
  // 3. 合法refer、（合法域名、合法URL）、非法refer、（非法域名、非法URL）这四项，一个data-id下只能配置一个或者都为空
  // 4. 匹配条件一致或者有存在交集的情况下（匹配条件包括URL匹配模式；文件类型；自定义文件类型；目录；指定常用类型；指定url），且控制动作均为禁止时，多条配置不能同时配置<合法refer>或者<合法域名>或者<合法URL>或者（<合法域名>和<合法URL>）"}
  RefererControlRule *EditAntiHotlinkingConfigRequestVisitControlRulesRefererControlRule `json:"referer-control-rule,omitempty" xml:"referer-control-rule,omitempty" type:"Struct"`
  // {"en":"UA head protection against hotlinking,
  // Note:
  // 1. Represents a group of UA head defense hotlinking, and only one group under a data-id
  // 2. when empty label means clear UA head protection hotlinking", "zh_CN":"标识UA头防盗链，
  // 注意：
  // 1. 表示一组UA头防盗链，一个data-id下只能一组
  // 2. 当传空标签表示清除UA头防盗链"}
  UaControlRule *EditAntiHotlinkingConfigRequestVisitControlRulesUaControlRule `json:"ua-control-rule,omitempty" xml:"ua-control-rule,omitempty" type:"Struct"`
  // {"en":"Configure other access control rules, such as invalid visitor regions, example:
  // advance-control-rules:{invalid-visitor-region:CN;JP;K}", "zh_CN":"配置其他访问控制策略，比如禁止的访客区域，JSON示例：
  // advance-control-rules:{invalid-visitor-region:CN;JP;KR}"}
  AdvanceControlRules *EditAntiHotlinkingConfigRequestVisitControlRulesAdvanceControlRules `json:"advance-control-rules,omitempty" xml:"advance-control-rules,omitempty" type:"Struct"`
  // {"en":"Configuration cookie control rules.Allow-cookie and forbidden-cookie are not allowed to be configured together.", "zh_CN":"配置Cookie防盗链策略。【允许的cookie】和【禁止的cookie】只允许配置一个"}
  CookieControlRules *EditAntiHotlinkingConfigRequestVisitControlRulesCookieControlRules `json:"cookie-control-rules,omitempty" xml:"cookie-control-rules,omitempty" type:"Struct"`
  // {"en":"Configuration custom header control rules.Header-whitelist and header-blacklist are not allowed to be configured together.", "zh_CN":"配置自定义头部防盗链。【头域黑名单】和【头域白名单】只允许配置一个"}
  CustomHeaderControlRules *EditAntiHotlinkingConfigRequestVisitControlRulesCustomHeaderControlRules `json:"custom-header-control-rules,omitempty" xml:"custom-header-control-rules,omitempty" type:"Struct"`
}

func (s EditAntiHotlinkingConfigRequestVisitControlRules) String() string {
  return tea.Prettify(s)
}

func (s EditAntiHotlinkingConfigRequestVisitControlRules) GoString() string {
  return s.String()
}

func (s *EditAntiHotlinkingConfigRequestVisitControlRules) SetDataId(v int64) *EditAntiHotlinkingConfigRequestVisitControlRules {
  s.DataId = &v
  return s
}

func (s *EditAntiHotlinkingConfigRequestVisitControlRules) SetPathPattern(v string) *EditAntiHotlinkingConfigRequestVisitControlRules {
  s.PathPattern = &v
  return s
}

func (s *EditAntiHotlinkingConfigRequestVisitControlRules) SetExceptPathPattern(v string) *EditAntiHotlinkingConfigRequestVisitControlRules {
  s.ExceptPathPattern = &v
  return s
}

func (s *EditAntiHotlinkingConfigRequestVisitControlRules) SetCustomPattern(v string) *EditAntiHotlinkingConfigRequestVisitControlRules {
  s.CustomPattern = &v
  return s
}

func (s *EditAntiHotlinkingConfigRequestVisitControlRules) SetFileType(v string) *EditAntiHotlinkingConfigRequestVisitControlRules {
  s.FileType = &v
  return s
}

func (s *EditAntiHotlinkingConfigRequestVisitControlRules) SetCustomFileType(v string) *EditAntiHotlinkingConfigRequestVisitControlRules {
  s.CustomFileType = &v
  return s
}

func (s *EditAntiHotlinkingConfigRequestVisitControlRules) SetSpecifyUrlPattern(v string) *EditAntiHotlinkingConfigRequestVisitControlRules {
  s.SpecifyUrlPattern = &v
  return s
}

func (s *EditAntiHotlinkingConfigRequestVisitControlRules) SetDirectory(v string) *EditAntiHotlinkingConfigRequestVisitControlRules {
  s.Directory = &v
  return s
}

func (s *EditAntiHotlinkingConfigRequestVisitControlRules) SetExceptFileType(v string) *EditAntiHotlinkingConfigRequestVisitControlRules {
  s.ExceptFileType = &v
  return s
}

func (s *EditAntiHotlinkingConfigRequestVisitControlRules) SetExceptCustomFileType(v string) *EditAntiHotlinkingConfigRequestVisitControlRules {
  s.ExceptCustomFileType = &v
  return s
}

func (s *EditAntiHotlinkingConfigRequestVisitControlRules) SetExceptDirectory(v string) *EditAntiHotlinkingConfigRequestVisitControlRules {
  s.ExceptDirectory = &v
  return s
}

func (s *EditAntiHotlinkingConfigRequestVisitControlRules) SetControlAction(v string) *EditAntiHotlinkingConfigRequestVisitControlRules {
  s.ControlAction = &v
  return s
}

func (s *EditAntiHotlinkingConfigRequestVisitControlRules) SetRewriteTo(v string) *EditAntiHotlinkingConfigRequestVisitControlRules {
  s.RewriteTo = &v
  return s
}

func (s *EditAntiHotlinkingConfigRequestVisitControlRules) SetPriority(v string) *EditAntiHotlinkingConfigRequestVisitControlRules {
  s.Priority = &v
  return s
}

func (s *EditAntiHotlinkingConfigRequestVisitControlRules) SetExceptionalRequest(v string) *EditAntiHotlinkingConfigRequestVisitControlRules {
  s.ExceptionalRequest = &v
  return s
}

func (s *EditAntiHotlinkingConfigRequestVisitControlRules) SetIgnoredCase(v string) *EditAntiHotlinkingConfigRequestVisitControlRules {
  s.IgnoredCase = &v
  return s
}

func (s *EditAntiHotlinkingConfigRequestVisitControlRules) SetIpControlRule(v *EditAntiHotlinkingConfigRequestVisitControlRulesIpControlRule) *EditAntiHotlinkingConfigRequestVisitControlRules {
  s.IpControlRule = v
  return s
}

func (s *EditAntiHotlinkingConfigRequestVisitControlRules) SetRefererControlRule(v *EditAntiHotlinkingConfigRequestVisitControlRulesRefererControlRule) *EditAntiHotlinkingConfigRequestVisitControlRules {
  s.RefererControlRule = v
  return s
}

func (s *EditAntiHotlinkingConfigRequestVisitControlRules) SetUaControlRule(v *EditAntiHotlinkingConfigRequestVisitControlRulesUaControlRule) *EditAntiHotlinkingConfigRequestVisitControlRules {
  s.UaControlRule = v
  return s
}

func (s *EditAntiHotlinkingConfigRequestVisitControlRules) SetAdvanceControlRules(v *EditAntiHotlinkingConfigRequestVisitControlRulesAdvanceControlRules) *EditAntiHotlinkingConfigRequestVisitControlRules {
  s.AdvanceControlRules = v
  return s
}

func (s *EditAntiHotlinkingConfigRequestVisitControlRules) SetCookieControlRules(v *EditAntiHotlinkingConfigRequestVisitControlRulesCookieControlRules) *EditAntiHotlinkingConfigRequestVisitControlRules {
  s.CookieControlRules = v
  return s
}

func (s *EditAntiHotlinkingConfigRequestVisitControlRules) SetCustomHeaderControlRules(v *EditAntiHotlinkingConfigRequestVisitControlRulesCustomHeaderControlRules) *EditAntiHotlinkingConfigRequestVisitControlRules {
  s.CustomHeaderControlRules = v
  return s
}

type EditAntiHotlinkingConfigRequestVisitControlRulesIpControlRule struct {
  // {"en":"Prohibited IP segment
  // Input parameter limit reference interface limit
  // Forbidden IP and exceptional IP cannot be configured at the same time", "zh_CN":"禁止的IP段
  // 支持输入IP或IP段，IP段之间用分号(;)隔开，如1.1.1.0/24;2.2.2.2
  // 禁止的IP和例外的IP，只能一个有值"}
  ForbiddenIps *string `json:"forbidden-ips,omitempty" xml:"forbidden-ips,omitempty"`
  // {"en":"The exception IP segment supports input IP or IP segment, and the IP segments are separated by a semicolon (;), such as 1.1.1.0/24; 2.2.2.2, some IP exceptions, no anti-theft chain", "zh_CN":"例外的IP段，支持输入IP或IP段，IP段之间用分号(;)隔开，如1.1.1.0/24;2.2.2.2，某些IP例外，不做防盗链"}
  AllowedIps *string `json:"allowed-ips,omitempty" xml:"allowed-ips,omitempty"`
}

func (s EditAntiHotlinkingConfigRequestVisitControlRulesIpControlRule) String() string {
  return tea.Prettify(s)
}

func (s EditAntiHotlinkingConfigRequestVisitControlRulesIpControlRule) GoString() string {
  return s.String()
}

func (s *EditAntiHotlinkingConfigRequestVisitControlRulesIpControlRule) SetForbiddenIps(v string) *EditAntiHotlinkingConfigRequestVisitControlRulesIpControlRule {
  s.ForbiddenIps = &v
  return s
}

func (s *EditAntiHotlinkingConfigRequestVisitControlRulesIpControlRule) SetAllowedIps(v string) *EditAntiHotlinkingConfigRequestVisitControlRulesIpControlRule {
  s.AllowedIps = &v
  return s
}

type EditAntiHotlinkingConfigRequestVisitControlRulesRefererControlRule struct {
  // {"en":"If any of the four terms 'nullreferer: legal referer, (legal domain name, legal URL), illegal referer, (illegal domain name, illegal URL)' is allowed, then 'nullreferer' cannot be null.If the four terms 'legal refer', 'legal domain name, legal URL', 'illegal refer', 'illegal domain name, illegal URL' are all null values, then 'whether to allow a null referer' must be null", "zh_CN":"是否允许空referer：合法refer、（合法域名、合法URL）、非法refer、（非法域名、非法URL）这六项任意一项有值，则&ldquo;是否允许空referer&rdquo;不能为空；合法refer、（合法域名、合法URL）、非法refer、（非法域名、非法URL）这四项都为空值，则&ldquo;是否允许空referer&rdquo;必须为空"}
  AllowNullReferer *string `json:"allow-null-referer,omitempty" xml:"allow-null-referer,omitempty"`
  // {"en":"Legal referer.", "zh_CN":"合法referer.可以输入url或域名，支持正则，可以多个，多个以空格隔开"}
  ValidReferer *string `json:"valid-referer,omitempty" xml:"valid-referer,omitempty"`
  // {"en":"Legal url, enter the correct url format", "zh_CN":"合法url，输入正确的url格式，不支持正则，可以多个，多个以分号分割。"}
  ValidUrl *string `json:"valid-url,omitempty" xml:"valid-url,omitempty"`
  // {"en":"Legal domain name", "zh_CN":"合法域名，不支持正则，可以多个，多个以分号分割"}
  ValidDomain *string `json:"valid-domain,omitempty" xml:"valid-domain,omitempty"`
  // {"en":"Illegal referer", "zh_CN":"非法referer，可以输入url或域名，支持正则，可以多个，多个以空格隔开"}
  InvalidReferer *string `json:"invalid-referer,omitempty" xml:"invalid-referer,omitempty"`
  // {"en":"Invalid url, enter the correct url format", "zh_CN":"非法url，输入正确的url格式，不支持正则，可以多个，多个以分号分割"}
  InvalidUrl *string `json:"invalid-url,omitempty" xml:"invalid-url,omitempty"`
  // {"en":"Illegal domain name", "zh_CN":"非法域名，不支持正则，可以多个，多个以分号分割"}
  InvalidDomain *string `json:"invalid-domain,omitempty" xml:"invalid-domain,omitempty"`
}

func (s EditAntiHotlinkingConfigRequestVisitControlRulesRefererControlRule) String() string {
  return tea.Prettify(s)
}

func (s EditAntiHotlinkingConfigRequestVisitControlRulesRefererControlRule) GoString() string {
  return s.String()
}

func (s *EditAntiHotlinkingConfigRequestVisitControlRulesRefererControlRule) SetAllowNullReferer(v string) *EditAntiHotlinkingConfigRequestVisitControlRulesRefererControlRule {
  s.AllowNullReferer = &v
  return s
}

func (s *EditAntiHotlinkingConfigRequestVisitControlRulesRefererControlRule) SetValidReferer(v string) *EditAntiHotlinkingConfigRequestVisitControlRulesRefererControlRule {
  s.ValidReferer = &v
  return s
}

func (s *EditAntiHotlinkingConfigRequestVisitControlRulesRefererControlRule) SetValidUrl(v string) *EditAntiHotlinkingConfigRequestVisitControlRulesRefererControlRule {
  s.ValidUrl = &v
  return s
}

func (s *EditAntiHotlinkingConfigRequestVisitControlRulesRefererControlRule) SetValidDomain(v string) *EditAntiHotlinkingConfigRequestVisitControlRulesRefererControlRule {
  s.ValidDomain = &v
  return s
}

func (s *EditAntiHotlinkingConfigRequestVisitControlRulesRefererControlRule) SetInvalidReferer(v string) *EditAntiHotlinkingConfigRequestVisitControlRulesRefererControlRule {
  s.InvalidReferer = &v
  return s
}

func (s *EditAntiHotlinkingConfigRequestVisitControlRulesRefererControlRule) SetInvalidUrl(v string) *EditAntiHotlinkingConfigRequestVisitControlRulesRefererControlRule {
  s.InvalidUrl = &v
  return s
}

func (s *EditAntiHotlinkingConfigRequestVisitControlRulesRefererControlRule) SetInvalidDomain(v string) *EditAntiHotlinkingConfigRequestVisitControlRulesRefererControlRule {
  s.InvalidDomain = &v
  return s
}

type EditAntiHotlinkingConfigRequestVisitControlRulesUaControlRule struct {
  // {"en":"Allowed clients, regular matching, no spaces allowed, to configure multiple UA such as:
  // <valid-user-agents>Android|iPhone</valid-user-agents>", "zh_CN":"允许的客户端，不允许空格，配置多个UA如：<valid-user-agents>Android|iPhone</valid-user-agents>"}
  ValidUserAgents *string `json:"valid-user-agents,omitempty" xml:"valid-user-agents,omitempty"`
  // {"en":"Forbidden client, regular match, no spaces allowed, configure multiple UA such as:
  // <invalid-user-agents>Android|iPhone</invalid-user-agents>", "zh_CN":"禁止的客户端，不允许空格，配置多个UA如：<invalid-user-agents>Android|iPhone</invalid-user-agents>"}
  InvalidUserAgents *string `json:"invalid-user-agents,omitempty" xml:"invalid-user-agents,omitempty"`
}

func (s EditAntiHotlinkingConfigRequestVisitControlRulesUaControlRule) String() string {
  return tea.Prettify(s)
}

func (s EditAntiHotlinkingConfigRequestVisitControlRulesUaControlRule) GoString() string {
  return s.String()
}

func (s *EditAntiHotlinkingConfigRequestVisitControlRulesUaControlRule) SetValidUserAgents(v string) *EditAntiHotlinkingConfigRequestVisitControlRulesUaControlRule {
  s.ValidUserAgents = &v
  return s
}

func (s *EditAntiHotlinkingConfigRequestVisitControlRulesUaControlRule) SetInvalidUserAgents(v string) *EditAntiHotlinkingConfigRequestVisitControlRulesUaControlRule {
  s.InvalidUserAgents = &v
  return s
}

type EditAntiHotlinkingConfigRequestVisitControlRulesAdvanceControlRules struct {
  // {"en":"Forbidden visitor regions, separate with semicolons. Note:
  // 1. Only support ISO 3166-1-alpha-2 two-letter country codes.
  // 2. If you have special regional configuration requirements, please contact your technical support.
  // 3. In the same set of rules, forbidden and allowed visitor regions cannot be configured at the same time.", "zh_CN":"禁止的访客区域，多个请用英文分号分隔。注意
  // 1、仅支持iso 3166-1国家二字简称
  // 2、如果有特殊区域配置需求，请联系您的专属。
  // 3、同一组规则里，禁止的访客区域、允许的访客区域，不能同时配"}
  InvalidVisitorRegion *string `json:"invalid-visitor-region,omitempty" xml:"invalid-visitor-region,omitempty"`
  // {"en":"Allowed  visitor regions, separate with semicolons. Note:
  // 1. Only support ISO 3166-1-alpha-2 two-letter country codes.
  // 2. If you have special regional configuration requirements, please contact your technical support.
  // 3. In the same set of rules, forbidden and allowed visitor regions cannot be configured at the same time.", "zh_CN":"允许的访客区域，多个请用英文分号分隔。注意
  // 1、仅支持iso 3166-1国家二字简称
  // 2、如果有特殊区域配置需求，请联系您的专属。
  // 3、同一组规则里，禁止的访客区域、允许的访客区域，不能同时配"}
  ValidVisitorRegion *string `json:"valid-visitor-region,omitempty" xml:"valid-visitor-region,omitempty"`
}

func (s EditAntiHotlinkingConfigRequestVisitControlRulesAdvanceControlRules) String() string {
  return tea.Prettify(s)
}

func (s EditAntiHotlinkingConfigRequestVisitControlRulesAdvanceControlRules) GoString() string {
  return s.String()
}

func (s *EditAntiHotlinkingConfigRequestVisitControlRulesAdvanceControlRules) SetInvalidVisitorRegion(v string) *EditAntiHotlinkingConfigRequestVisitControlRulesAdvanceControlRules {
  s.InvalidVisitorRegion = &v
  return s
}

func (s *EditAntiHotlinkingConfigRequestVisitControlRulesAdvanceControlRules) SetValidVisitorRegion(v string) *EditAntiHotlinkingConfigRequestVisitControlRulesAdvanceControlRules {
  s.ValidVisitorRegion = &v
  return s
}

type EditAntiHotlinkingConfigRequestVisitControlRulesCookieControlRules struct {
  // {"en":"Allow Cookie.Fill in regular format, e.g(. *) (range1 | range2) (. *)", "zh_CN":"允许的cookie。填写正则格式，比如(.*)(range1|range2)(.*)。"}
  AllowCookie *string `json:"allow-cookie,omitempty" xml:"allow-cookie,omitempty"`
  // {"en":"Allow Null Cookie.Only true and false are allowed to be filled in allow null cookie", "zh_CN":"是否允许空cookie。只允许填写true或false。"}
  AllowNullCookie *string `json:"allow-null-cookie,omitempty" xml:"allow-null-cookie,omitempty"`
  // {"en":"Forbidden Cookie.Fill in regular format, e.g(. *) (range1 | range2) (. *)", "zh_CN":"禁止的cookie。填写正则格式，比如(.*)(range1|range2)(.*)。"}
  ForbiddenCookie *string `json:"forbidden-cookie,omitempty" xml:"forbidden-cookie,omitempty"`
}

func (s EditAntiHotlinkingConfigRequestVisitControlRulesCookieControlRules) String() string {
  return tea.Prettify(s)
}

func (s EditAntiHotlinkingConfigRequestVisitControlRulesCookieControlRules) GoString() string {
  return s.String()
}

func (s *EditAntiHotlinkingConfigRequestVisitControlRulesCookieControlRules) SetAllowCookie(v string) *EditAntiHotlinkingConfigRequestVisitControlRulesCookieControlRules {
  s.AllowCookie = &v
  return s
}

func (s *EditAntiHotlinkingConfigRequestVisitControlRulesCookieControlRules) SetAllowNullCookie(v string) *EditAntiHotlinkingConfigRequestVisitControlRulesCookieControlRules {
  s.AllowNullCookie = &v
  return s
}

func (s *EditAntiHotlinkingConfigRequestVisitControlRulesCookieControlRules) SetForbiddenCookie(v string) *EditAntiHotlinkingConfigRequestVisitControlRulesCookieControlRules {
  s.ForbiddenCookie = &v
  return s
}

type EditAntiHotlinkingConfigRequestVisitControlRulesCustomHeaderControlRules struct {
  // {"en":"Header Direction.Can choose from the client or the server.Only allowed to fill client or server", "zh_CN":"来源。可选择来源于客户端还是服务端。客户端填写client，服务端填写server"}
  HeaderDirection *string `json:"header-direction,omitempty" xml:"header-direction,omitempty"`
  // {"en":"Header Whitelist", "zh_CN":"头域白名单"}
  HeaderWhitelist *string `json:"header-whitelist,omitempty" xml:"header-whitelist,omitempty"`
  // {"en":"Header Value Whitelist", "zh_CN":"头域值白名单"}
  HeaderValueWhitelist *string `json:"header-value-whitelist,omitempty" xml:"header-value-whitelist,omitempty"`
  // {"en":"Header Blacklist", "zh_CN":"头域黑名单"}
  HeaderBlacklist *string `json:"header-blacklist,omitempty" xml:"header-blacklist,omitempty"`
  // {"en":"Header Value Blacklist", "zh_CN":"头域值黑名单"}
  HeaderValueBlacklist *string `json:"header-value-blacklist,omitempty" xml:"header-value-blacklist,omitempty"`
}

func (s EditAntiHotlinkingConfigRequestVisitControlRulesCustomHeaderControlRules) String() string {
  return tea.Prettify(s)
}

func (s EditAntiHotlinkingConfigRequestVisitControlRulesCustomHeaderControlRules) GoString() string {
  return s.String()
}

func (s *EditAntiHotlinkingConfigRequestVisitControlRulesCustomHeaderControlRules) SetHeaderDirection(v string) *EditAntiHotlinkingConfigRequestVisitControlRulesCustomHeaderControlRules {
  s.HeaderDirection = &v
  return s
}

func (s *EditAntiHotlinkingConfigRequestVisitControlRulesCustomHeaderControlRules) SetHeaderWhitelist(v string) *EditAntiHotlinkingConfigRequestVisitControlRulesCustomHeaderControlRules {
  s.HeaderWhitelist = &v
  return s
}

func (s *EditAntiHotlinkingConfigRequestVisitControlRulesCustomHeaderControlRules) SetHeaderValueWhitelist(v string) *EditAntiHotlinkingConfigRequestVisitControlRulesCustomHeaderControlRules {
  s.HeaderValueWhitelist = &v
  return s
}

func (s *EditAntiHotlinkingConfigRequestVisitControlRulesCustomHeaderControlRules) SetHeaderBlacklist(v string) *EditAntiHotlinkingConfigRequestVisitControlRulesCustomHeaderControlRules {
  s.HeaderBlacklist = &v
  return s
}

func (s *EditAntiHotlinkingConfigRequestVisitControlRulesCustomHeaderControlRules) SetHeaderValueBlacklist(v string) *EditAntiHotlinkingConfigRequestVisitControlRulesCustomHeaderControlRules {
  s.HeaderValueBlacklist = &v
  return s
}

type EditAntiHotlinkingConfigResponse struct {
  // {"en":"httpstatus=202; Indicates that the new domain API was successfully invoked, and the current deployment of the new domain can be viewed using x-cnc-request-id in the header", "zh_CN":"httpstatus=202;   表示成功调用新增域名接口，可使用header中的x-cnc-request-id查看当前新增域名的部署情况"}
  HttpStatus *int `json:"http status code,omitempty" xml:"http status code,omitempty" require:"true"`
  // {"en":"Uniquely identified id for querying tasks per request (for all API)", "zh_CN":"唯一标示的id，用于查询每次请求的任务 （适用全部接口）"}
  XCncRequestId *string `json:"x-cnc-request-id,omitempty" xml:"x-cnc-request-id,omitempty" require:"true"`
  // {"en":"The URL used to access the domain information, where domain-id is the unique token generated by our cloud platform for the domain name and whose value is a string.", "zh_CN":"响应信用于访问该域名信息的URL，其中domain-id为我司云平台为该域名生成的唯一标示，其值为字符串。"}
  Location *string `json:"location,omitempty" xml:"location,omitempty" require:"true"`
  // {"en":"Error code, which appears when HTTPStatus is not 202, represents the error type of the current request call", "zh_CN":"错误代码，当HTTPStatus不为202时出现，表示当前请求调用的错误类型"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response information, success when successful", "zh_CN":"响应信息，成功时为success"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s EditAntiHotlinkingConfigResponse) String() string {
  return tea.Prettify(s)
}

func (s EditAntiHotlinkingConfigResponse) GoString() string {
  return s.String()
}

func (s *EditAntiHotlinkingConfigResponse) SetHttpStatus(v int) *EditAntiHotlinkingConfigResponse {
  s.HttpStatus = &v
  return s
}

func (s *EditAntiHotlinkingConfigResponse) SetXCncRequestId(v string) *EditAntiHotlinkingConfigResponse {
  s.XCncRequestId = &v
  return s
}

func (s *EditAntiHotlinkingConfigResponse) SetLocation(v string) *EditAntiHotlinkingConfigResponse {
  s.Location = &v
  return s
}

func (s *EditAntiHotlinkingConfigResponse) SetCode(v string) *EditAntiHotlinkingConfigResponse {
  s.Code = &v
  return s
}

func (s *EditAntiHotlinkingConfigResponse) SetMessage(v string) *EditAntiHotlinkingConfigResponse {
  s.Message = &v
  return s
}

type EditAntiHotlinkingConfigPaths struct {
  // {"en":"the domain whoes need query config", "zh_CN":"需要查询配置的域名或域名id"}
  DomainName *string `json:"domain-name,omitempty" xml:"domain-name,omitempty" require:"true"`
}

func (s EditAntiHotlinkingConfigPaths) String() string {
  return tea.Prettify(s)
}

func (s EditAntiHotlinkingConfigPaths) GoString() string {
  return s.String()
}

func (s *EditAntiHotlinkingConfigPaths) SetDomainName(v string) *EditAntiHotlinkingConfigPaths {
  s.DomainName = &v
  return s
}

type EditAntiHotlinkingConfigParameters struct {
}

func (s EditAntiHotlinkingConfigParameters) String() string {
  return tea.Prettify(s)
}

func (s EditAntiHotlinkingConfigParameters) GoString() string {
  return s.String()
}

type EditAntiHotlinkingConfigRequestHeader struct {
}

func (s EditAntiHotlinkingConfigRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s EditAntiHotlinkingConfigRequestHeader) GoString() string {
  return s.String()
}

type EditAntiHotlinkingConfigResponseHeader struct {
}

func (s EditAntiHotlinkingConfigResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s EditAntiHotlinkingConfigResponseHeader) GoString() string {
  return s.String()
}




type QueryBack2OriginProtocolRewriteConfigRequest struct {
}

func (s QueryBack2OriginProtocolRewriteConfigRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryBack2OriginProtocolRewriteConfigRequest) GoString() string {
  return s.String()
}

type QueryBack2OriginProtocolRewriteConfigResponse struct {
  // {"en":"Error code. Appeared when http status is not 200 or 202.", "zh_CN":"错误代码，当HTTPStatus不为202时出现，表示当前请求调用的错误类型"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Reponse message.", "zh_CN":"响应信息，成功时为success"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Response data", "zh_CN":"响应数据"}
  Data *QueryBack2OriginProtocolRewriteConfigResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
}

func (s QueryBack2OriginProtocolRewriteConfigResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryBack2OriginProtocolRewriteConfigResponse) GoString() string {
  return s.String()
}

func (s *QueryBack2OriginProtocolRewriteConfigResponse) SetCode(v string) *QueryBack2OriginProtocolRewriteConfigResponse {
  s.Code = &v
  return s
}

func (s *QueryBack2OriginProtocolRewriteConfigResponse) SetMessage(v string) *QueryBack2OriginProtocolRewriteConfigResponse {
  s.Message = &v
  return s
}

func (s *QueryBack2OriginProtocolRewriteConfigResponse) SetData(v *QueryBack2OriginProtocolRewriteConfigResponseData) *QueryBack2OriginProtocolRewriteConfigResponse {
  s.Data = v
  return s
}

type QueryBack2OriginProtocolRewriteConfigResponseData struct {
  // {"en":"Domain id", "zh_CN":"域名ID"}
  DomainId *int `json:"domainId,omitempty" xml:"domainId,omitempty" require:"true"`
  // {"en":"Domain name", "zh_CN":"域名"}
  DomainName *string `json:"domainName,omitempty" xml:"domainName,omitempty" require:"true"`
  // {"en":"Back To Origin Rewrite Rule Configuration, parent node 1. When you need to configure the Back To Origin Rewrite Rule, this must be filled in. 2. Configuration of clearing for <backToOriginRewriteRule/>.", "zh_CN":"回源协议配置，父标签 1.需要设置回源协议改写配置时，此项必填 2.为<backToOriginRewriteRule/>时清空配置"}
  BackToOriginRewriteRule *QueryBack2OriginProtocolRewriteConfigResponseDataBackToOriginRewriteRule `json:"backToOriginRewriteRule,omitempty" xml:"backToOriginRewriteRule,omitempty" require:"true" type:"Struct"`
}

func (s QueryBack2OriginProtocolRewriteConfigResponseData) String() string {
  return tea.Prettify(s)
}

func (s QueryBack2OriginProtocolRewriteConfigResponseData) GoString() string {
  return s.String()
}

func (s *QueryBack2OriginProtocolRewriteConfigResponseData) SetDomainId(v int) *QueryBack2OriginProtocolRewriteConfigResponseData {
  s.DomainId = &v
  return s
}

func (s *QueryBack2OriginProtocolRewriteConfigResponseData) SetDomainName(v string) *QueryBack2OriginProtocolRewriteConfigResponseData {
  s.DomainName = &v
  return s
}

func (s *QueryBack2OriginProtocolRewriteConfigResponseData) SetBackToOriginRewriteRule(v *QueryBack2OriginProtocolRewriteConfigResponseDataBackToOriginRewriteRule) *QueryBack2OriginProtocolRewriteConfigResponseData {
  s.BackToOriginRewriteRule = v
  return s
}

type QueryBack2OriginProtocolRewriteConfigResponseDataBackToOriginRewriteRule struct {
  // {"en":"The specified protocol is either http or https.", "zh_CN":"改写后的回源协议，可选值：http、https"}
  Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty" require:"true"`
  // {"en":"If the protocol is http, the default is 80. If the protocol is https, the default is 443", "zh_CN":"改写后的回源端口，若protocol为http时，默认为80，若protocol为https时，默认为443"}
  Port *string `json:"port,omitempty" xml:"port,omitempty" require:"true"`
}

func (s QueryBack2OriginProtocolRewriteConfigResponseDataBackToOriginRewriteRule) String() string {
  return tea.Prettify(s)
}

func (s QueryBack2OriginProtocolRewriteConfigResponseDataBackToOriginRewriteRule) GoString() string {
  return s.String()
}

func (s *QueryBack2OriginProtocolRewriteConfigResponseDataBackToOriginRewriteRule) SetProtocol(v string) *QueryBack2OriginProtocolRewriteConfigResponseDataBackToOriginRewriteRule {
  s.Protocol = &v
  return s
}

func (s *QueryBack2OriginProtocolRewriteConfigResponseDataBackToOriginRewriteRule) SetPort(v string) *QueryBack2OriginProtocolRewriteConfigResponseDataBackToOriginRewriteRule {
  s.Port = &v
  return s
}

type QueryBack2OriginProtocolRewriteConfigPaths struct {
  // {"en":"Domain name or domain id to query configuration", "zh_CN":"需要查询配置的域名（domainName）或域名id（domainId）"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
}

func (s QueryBack2OriginProtocolRewriteConfigPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryBack2OriginProtocolRewriteConfigPaths) GoString() string {
  return s.String()
}

func (s *QueryBack2OriginProtocolRewriteConfigPaths) SetDomain(v string) *QueryBack2OriginProtocolRewriteConfigPaths {
  s.Domain = &v
  return s
}

type QueryBack2OriginProtocolRewriteConfigParameters struct {
}

func (s QueryBack2OriginProtocolRewriteConfigParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryBack2OriginProtocolRewriteConfigParameters) GoString() string {
  return s.String()
}

type QueryBack2OriginProtocolRewriteConfigRequestHeader struct {
}

func (s QueryBack2OriginProtocolRewriteConfigRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryBack2OriginProtocolRewriteConfigRequestHeader) GoString() string {
  return s.String()
}

type QueryBack2OriginProtocolRewriteConfigResponseHeader struct {
}

func (s QueryBack2OriginProtocolRewriteConfigResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryBack2OriginProtocolRewriteConfigResponseHeader) GoString() string {
  return s.String()
}




type UpdateLiveVisitControlConfigRequest struct {
  // {"en":"Streaming visit control configuration, parent node
  // 1. This item is required when the function of streaming visit control configuration needs to be set
  // 2. Clear the Streaming visit control configuration when <visitControlRules/>", "zh_CN":"流媒体防盗链配置，父标签
  // 1.需要设置流媒体防盗链配置时，此项必填
  // 2.为<visitControlRules/>时清空流媒体防盗链配置"}
  VisitControlRules []*UpdateLiveVisitControlConfigRequestVisitControlRules `json:"visitControlRules,omitempty" xml:"visitControlRules,omitempty" require:"true" type:"Repeated"`
}

func (s UpdateLiveVisitControlConfigRequest) String() string {
  return tea.Prettify(s)
}

func (s UpdateLiveVisitControlConfigRequest) GoString() string {
  return s.String()
}

func (s *UpdateLiveVisitControlConfigRequest) SetVisitControlRules(v []*UpdateLiveVisitControlConfigRequestVisitControlRules) *UpdateLiveVisitControlConfigRequest {
  s.VisitControlRules = v
  return s
}

type UpdateLiveVisitControlConfigRequestVisitControlRules struct     {
  // {"en":"Control action, optional values: allow, forbid
  // Note: when add or modify a rule, you must configure both controlaction and IP, referer. At least one of IP and referer must be configured, otherwise the function will not work.", "zh_CN":"控制动作，允许或禁止，可选值：allow，forbid
  // 注意：配置或修改流媒体防盗链时，必须同时配置controlAction和ip、referer，ip和referer至少配置一项，否则此防盗链功能无效。"}
  ControlAction *string `json:"controlAction,omitempty" xml:"controlAction,omitempty" require:"true"`
  // {"en":"Allowed or forbidden IP, support IP and IP segment, parent node. Controlaction must be configured at the same time. Such as:
  // <controlAction>forbid</controlAction>
  // <ips>
  // <ip>1.1.1.1</ip>
  // <ip>2.2.2.0/24</ip>
  // </ips>", "zh_CN":"允许或禁止的IP，支持IP或IP段，父标签。必须同时配置controlAction。
  // 格式如
  // <controlAction>forbid</controlAction>
  // <ips>
  // <ip>1.1.1.1</ip> 
  // <ip>2.2.2.0/24</ip>
  // </ips>"}
  Ips []*string `json:"ips,omitempty" xml:"ips,omitempty" type:"Repeated"`
  // {"en":"The allowed or forbidden referer, support domain name and URL format. Parent node. It must start with the protocol header, such as http://, https://. Controlaction must be configured at the same time. Such as:
  // <controlAction>forbid</controlAction>
  // <referers>
  // <referer>http://www.referer1.com</referer>
  // <referer>http://www.referer2.com</referer>
  // </referers>", "zh_CN":"允许或禁止的referer，支持输入域名或url格式，父标签，必须以协议头开头，如http://、https://。必须同时配置controlAction。格式如：
  // <controlAction>forbid</controlAction>
  // <referers>
  // <referer>http://www.referer1.com</referer>
  // <referer>http://www.referer2.com</referer>
  // </referers>"}
  Referers []*string `json:"referers,omitempty" xml:"referers,omitempty" type:"Repeated"`
  // {"en":"Allow or forbid null referer. Optional value is 'true' or 'false',default value is 'false'.
  // When controlAction equals forbid, it means whether null referer is forbidden. If it is true, it means null referer is prohibited.
  // When controlAction  equals  allow, it means whether null referer is allowed. If it is true, it means null referer is allowed.", "zh_CN":"是否允许/禁止空referer，可选值为true、false，默认为否。
  // 当控制动作=forbid，表示是否禁止空referer，为true，表示禁止空referer
  // 当控制动作=allow，表示是否允许空referer，为true，表示允许空referer"}
  AllowNullReferer *bool `json:"allowNullReferer,omitempty" xml:"allowNullReferer,omitempty"`
  // {"en":"IP and refer condition relationship, optional values: and, or, default to and.", "zh_CN":"ip和refer条件关系，可选值：and，or，默认为and
  // 为and，表示ip和refer都满足，才允许/禁止。为or，表示ip和refer满足一个，就允许/禁止"}
  ControlRelation *string `json:"controlRelation,omitempty" xml:"controlRelation,omitempty"`
  // {"en":"Data-id is to indicate a specific group configuration when the client has multiple groups of configurations. Data-id can be retrieved through a query interface. Note: A. If data-id is passed, it means that one group of configuration items is specified to be modified, and no other group configuration items need to be modified. B. If multiple groups of configurations are included, some of them are configured with data-id and others are not, then the expression of data-id is used to modify a specific group of configurations, and a new group of configurations is added on the original basis without the expression of data-id. C. If the data-id is not transmitted, it means that the original configuration will be fully covered by this configuration. D. If no configuration parameter is passed, only domain name and secondary label are passed, which means that all configuration of domain name secondary service corresponding to this interface is cleared. E. If there is no specific configuration item in a set of configurations, the data-id must be filled in, and the value is the actual data-id, which means clearing the value of the corresponding data-id configuration item; it is not allowed that there is no specific configuration item or data-id in a set of configurations.", "zh_CN":"配置多组配置时，具体某组配置的id。dataId可以通过查询接口获取。 注意： a、如果有传dataId，说明指定修改其中一组配置项内容，不需求修改其他组配置内容不需要入参； b、如果入参多组配置，其中有些组配置有传dataId，有些没有传，则有传dataId的表示修改具体某组配置，没有传dataId的表示在原来基础上新增一组配置； c、如果入参都没有传dataId,表示用本次的配置全量覆盖原先配置； d、如果入参没有传任何配置项参数，只传了域名和二级标签，表示清空这个接口对应域名二级服务所有配置； e、如果一组配置没有具体的配置项，则dataId必填，且值为实际存在的dataId，表示清空这个dataId对应配置项的值；不允许一组配置没有具体的配置项也没有dataId。"}
  DataId *int `json:"dataId,omitempty" xml:"dataId,omitempty"`
}

func (s UpdateLiveVisitControlConfigRequestVisitControlRules) String() string {
  return tea.Prettify(s)
}

func (s UpdateLiveVisitControlConfigRequestVisitControlRules) GoString() string {
  return s.String()
}

func (s *UpdateLiveVisitControlConfigRequestVisitControlRules) SetControlAction(v string) *UpdateLiveVisitControlConfigRequestVisitControlRules {
  s.ControlAction = &v
  return s
}

func (s *UpdateLiveVisitControlConfigRequestVisitControlRules) SetIps(v []*string) *UpdateLiveVisitControlConfigRequestVisitControlRules {
  s.Ips = v
  return s
}

func (s *UpdateLiveVisitControlConfigRequestVisitControlRules) SetReferers(v []*string) *UpdateLiveVisitControlConfigRequestVisitControlRules {
  s.Referers = v
  return s
}

func (s *UpdateLiveVisitControlConfigRequestVisitControlRules) SetAllowNullReferer(v bool) *UpdateLiveVisitControlConfigRequestVisitControlRules {
  s.AllowNullReferer = &v
  return s
}

func (s *UpdateLiveVisitControlConfigRequestVisitControlRules) SetControlRelation(v string) *UpdateLiveVisitControlConfigRequestVisitControlRules {
  s.ControlRelation = &v
  return s
}

func (s *UpdateLiveVisitControlConfigRequestVisitControlRules) SetDataId(v int) *UpdateLiveVisitControlConfigRequestVisitControlRules {
  s.DataId = &v
  return s
}

type UpdateLiveVisitControlConfigResponse struct {
  // {"en":"The error code, when HTTPStatus is not 202, indicates the type of error the current request is calling.", "zh_CN":"错误代码，当HTTPStatus不为202时出现，表示当前请求调用的错误类型"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response information, when success is successful", "zh_CN":"响应信息，成功时为success"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"response data", "zh_CN":"响应数据"}
  Data *string `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s UpdateLiveVisitControlConfigResponse) String() string {
  return tea.Prettify(s)
}

func (s UpdateLiveVisitControlConfigResponse) GoString() string {
  return s.String()
}

func (s *UpdateLiveVisitControlConfigResponse) SetCode(v string) *UpdateLiveVisitControlConfigResponse {
  s.Code = &v
  return s
}

func (s *UpdateLiveVisitControlConfigResponse) SetMessage(v string) *UpdateLiveVisitControlConfigResponse {
  s.Message = &v
  return s
}

func (s *UpdateLiveVisitControlConfigResponse) SetData(v string) *UpdateLiveVisitControlConfigResponse {
  s.Data = &v
  return s
}

type UpdateLiveVisitControlConfigPaths struct {
  // {"en":"The domain whoes need query config.", "zh_CN":"需要查询配置的域名或域名id"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
}

func (s UpdateLiveVisitControlConfigPaths) String() string {
  return tea.Prettify(s)
}

func (s UpdateLiveVisitControlConfigPaths) GoString() string {
  return s.String()
}

func (s *UpdateLiveVisitControlConfigPaths) SetDomain(v string) *UpdateLiveVisitControlConfigPaths {
  s.Domain = &v
  return s
}

type UpdateLiveVisitControlConfigParameters struct {
}

func (s UpdateLiveVisitControlConfigParameters) String() string {
  return tea.Prettify(s)
}

func (s UpdateLiveVisitControlConfigParameters) GoString() string {
  return s.String()
}

type UpdateLiveVisitControlConfigRequestHeader struct {
}

func (s UpdateLiveVisitControlConfigRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s UpdateLiveVisitControlConfigRequestHeader) GoString() string {
  return s.String()
}

type UpdateLiveVisitControlConfigResponseHeader struct {
}

func (s UpdateLiveVisitControlConfigResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s UpdateLiveVisitControlConfigResponseHeader) GoString() string {
  return s.String()
}




type QueryQueryStringUrlConfigRequest struct {
}

func (s QueryQueryStringUrlConfigRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryQueryStringUrlConfigRequest) GoString() string {
  return s.String()
}

type QueryQueryStringUrlConfigResponse struct {
  // {"en":"Domain name or domain ID that needs to be queried for configuration.", "zh_CN":"需要查询配置的域名"}
  DomainName *string `json:"domain-name,omitempty" xml:"domain-name,omitempty" require:"true"`
  // {"en":"Domain name or domain ID that needs to be queried for configuration.", "zh_CN":"需要查询配置的域名id"}
  DomainId *string `json:"domain-id,omitempty" xml:"domain-id,omitempty" require:"true"`
  // {"en":"Query String Settings Configuration, parent node
  // 1. When you need to configure the query string, this must be filled in.
  // 2. Configuration of clearing query string settings for <query-string-settings/>.", "zh_CN":"查询串设置配置，父标签
  // 1.需要设置查询串配置时，此项必填
  // 2.为<query-string-settings/>时清空查询串设置的配置"}
  QueryStringSettings []*QueryQueryStringUrlConfigResponseQueryStringSettings `json:"query-string-settings,omitempty" xml:"query-string-settings,omitempty" require:"true" type:"Repeated"`
}

func (s QueryQueryStringUrlConfigResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryQueryStringUrlConfigResponse) GoString() string {
  return s.String()
}

func (s *QueryQueryStringUrlConfigResponse) SetDomainName(v string) *QueryQueryStringUrlConfigResponse {
  s.DomainName = &v
  return s
}

func (s *QueryQueryStringUrlConfigResponse) SetDomainId(v string) *QueryQueryStringUrlConfigResponse {
  s.DomainId = &v
  return s
}

func (s *QueryQueryStringUrlConfigResponse) SetQueryStringSettings(v []*QueryQueryStringUrlConfigResponseQueryStringSettings) *QueryQueryStringUrlConfigResponse {
  s.QueryStringSettings = v
  return s
}

type QueryQueryStringUrlConfigResponseQueryStringSettings struct     {
  // {"en":"The url matching mode. If all matches, the input parameters can be configured as: .*
  // Note: URL matching mode, file type (custom file type), commonly used type, specified url, directory, with and only one required", "zh_CN":"url匹配模式，支持正则，如果是全部匹配，入参可以配置为：.*
  // 注：url匹配模式、文件类型（自定义文件类型）、常用类型、指定url、目录，有且仅有一项必填"}
  PathPattern *string `json:"path-pattern,omitempty" xml:"path-pattern,omitempty" require:"true"`
  // {"en":"File Type: Specify the file types to set the anti-theft chain. Optional file types include: GIF PNG BMP JPEG JPG HTML HTM shtml MP3 WMA flv MP4 WMV zip exe rar CSS TXT ICO JS SWF m3u8 XML f4m bootstarps if all types are required. Multiple separated by semicolons, all and specific file types cannot be configured at the same time", "zh_CN":"文件类型：指定文件类型进行防盗链设置。可选文件类型包括：gif png bmp jpeg jpg html htm shtml mp3 wma flv mp4 wmv zip exe rar css txt ico js swf m3u8 xml f4m bootstarp ts 如果需要全部类型，则直接传all。多个以分号隔开，all和具体文件类型不能同时配置"}
  FileTypes *string `json:"file-types,omitempty" xml:"file-types,omitempty" require:"true"`
  // {"en":"Custom file type: Fill in the appropriate identifiable file type according to your own needs outside the specified file type. It can be used with file-type. If file-type is also configured, the actual file type in effect is the sum of the two entries.", "zh_CN":"自定义文件类型：在指定文件类型外根据自身需求，填写适当的可识别文件类型。可以搭配file-type使用。如果file-type也有配置，实际生效的文件类型是两个入参的总和"}
  CustomFileTypes *string `json:"custom-file-types,omitempty" xml:"custom-file-types,omitempty" require:"true"`
  // {"en":"Common types: optional values are home page and all
  // All: All files types
  // Homepage: home page", "zh_CN":"常用类型：可选值为homepage和all
  // all：全部文件
  // homepage：首页"}
  CustomPattern *string `json:"custom-pattern,omitempty" xml:"custom-pattern,omitempty" require:"true"`
  // {"en":"Specify url, not http (s): // start, multiple separated by newline", "zh_CN":"指定url，非http(s):// 开头，多个以换行分隔"}
  SpecifyUrlPattern *string `json:"specify-url-pattern,omitempty" xml:"specify-url-pattern,omitempty" require:"true"`
  // {"en":"directories, you can enter a legitimate directory format. Start and end with / and multiple separated by semicolons.", "zh_CN":"目录，可输入合法的目录格式。以/开头和结尾，多个以分号隔开。"}
  Directories *string `json:"directories,omitempty" xml:"directories,omitempty" require:"true"`
  // {"en":"Priority, which represents the priority execution order of the customer's multi-group configuration. The bigger the number, the higher the priority. No transmission defaults to 10, not empty.", "zh_CN":"优先级，表示客户多组配置的优先执行顺序。数字越大，优先级越高。不传默认为10，不可清空。"}
  Priority *int `json:"priority,omitempty" xml:"priority,omitempty" require:"true"`
  // {"en":"Whether to ignore letter case.", "zh_CN":"是否忽略大小写：允许值为true和false，默认为忽略"}
  IgnoreLetterCase *string `json:"ignore-letter-case,omitempty" xml:"ignore-letter-case,omitempty" require:"true"`
  // {"en":"Define the file types to be compressed.  'text/' will be compressed by default.", "zh_CN":"缓存是否忽略查询串，允许值为true和false。
  // true表示忽略查询串，相同拷贝；false表示不忽略，分别缓存；当存在缓存保留参数或缓存删除参数时返回空。"}
  IgnoreQueryString *string `json:"ignore-query-string,omitempty" xml:"ignore-query-string,omitempty" require:"true"`
  // {"en":"Cache with the specified query string parameters. If the kept parameter values are the same, one copy will be cached.
  // Note:
  // 1. query-string-kept and query-string-removed are mutually exclusive, and only one of them has a value.
  // 2. query-string-kept and ignore-query-string are mutually exclusive, and only one has a value.", "zh_CN":"缓存保留参数，指定保留的参数值相同，则缓存一份。
  // 注：
  // 1.query-string-kept和query-string-removed两者互斥，只能一个有值。
  // 2.query-string-kept和ignore-query-string两者互斥，只能一个有值。"}
  QueryStringKept *string `json:"query-string-kept,omitempty" xml:"query-string-kept,omitempty" require:"true"`
  // {"en":"Cache without the specified query string parameters. After deleting the specified parameter, if the other parameter values are the same, one copy will be cached.
  // 1. query-string-kept and query string removed are mutually exclusive, and only one has a value.
  // 2. query-string-removed and ignore-query-string are mutually exclusive.", "zh_CN":"缓存删除参数，删除指定的参数后，其余参数值相同，则缓存一份。
  // 1.query-string-kept和query-string-removed两者互斥，只能一个有值。
  // 2.query-string-removed和ignore-query-string两者互斥，只能一个有值。"}
  QueryStringRemoved *string `json:"query-string-removed,omitempty" xml:"query-string-removed,omitempty" require:"true"`
  // {"en":"Whether to use the original URL back source, the allowable values are true and false.
  // When ignore-query-string is true or not set, source-with-query is true to indicate that the source is returned according to the original request, and false to indicate that the question mark is returned.
  // When ignore-query-string is false, this default setting is empty (input is invalid)", "zh_CN":"是否用原始url回源，允许值为true和false。
  // ignore-query-string为true或未设置时，source-with-query为true表示按原始请求回源；为false表示去问号回源。
  // ignore-query-string为false时，此项默认设置为空（输入无效）。"}
  SourceWithQuery *string `json:"source-with-query,omitempty" xml:"source-with-query,omitempty" require:"true"`
  // {"en":"Return to the source after specifying the reserved parameter value. Please separate them with semicolons, if no parameters reserved, please fill in:- . 1. Source-key-kept and ignore-query-string are mutually exclusive, and only one of them has a value. 2. Source-key-kept and source-key-removed are mutually exclusive, and only one of them has a value.
  // ", "zh_CN":"回源保留参数，指定保留的参数值后回源。多个请以英文分号分隔，任何参数都不保留请填：- 1、source-key-kept和ignore-query-string两者互斥，只能一个有值。 2、source-key-kept和source-key-removed两者互斥，只能一个有值。
  // "}
  SourceKeyKept *string `json:"source-key-kept,omitempty" xml:"source-key-kept,omitempty" require:"true"`
  // {"en":"Return to the source after specifying the deleted parameter value. Please separate them with semicolons, and if you do not delete any parameters, please fill in:- . 1. Source-key-removed and ignore-query-string are mutually exclusive, and only one of them has a value. 2. Source-key-kept and source-key-removed are mutually exclusive, and only one of them has a value.
  // ", "zh_CN":"回源删除参数，指定删除的参数值后回源。多个请以英文分号分隔，任何参数都不删除请填：- 1、source-key-removed和ignore-query-string两者互斥，只能一个有值。 2、source-key-kept和source-key-removed两者互斥，只能一个有值。
  // "}
  SourceKeyRemoved *string `json:"source-key-removed,omitempty" xml:"source-key-removed,omitempty" require:"true"`
  // {"en":"When configuring multiple configurations, the ID of a specific group of configurations", "zh_CN":"配置多组配置时，具体某组配置的id"}
  DataId *int32 `json:"data-id,omitempty" xml:"data-id,omitempty" require:"true"`
}

func (s QueryQueryStringUrlConfigResponseQueryStringSettings) String() string {
  return tea.Prettify(s)
}

func (s QueryQueryStringUrlConfigResponseQueryStringSettings) GoString() string {
  return s.String()
}

func (s *QueryQueryStringUrlConfigResponseQueryStringSettings) SetPathPattern(v string) *QueryQueryStringUrlConfigResponseQueryStringSettings {
  s.PathPattern = &v
  return s
}

func (s *QueryQueryStringUrlConfigResponseQueryStringSettings) SetFileTypes(v string) *QueryQueryStringUrlConfigResponseQueryStringSettings {
  s.FileTypes = &v
  return s
}

func (s *QueryQueryStringUrlConfigResponseQueryStringSettings) SetCustomFileTypes(v string) *QueryQueryStringUrlConfigResponseQueryStringSettings {
  s.CustomFileTypes = &v
  return s
}

func (s *QueryQueryStringUrlConfigResponseQueryStringSettings) SetCustomPattern(v string) *QueryQueryStringUrlConfigResponseQueryStringSettings {
  s.CustomPattern = &v
  return s
}

func (s *QueryQueryStringUrlConfigResponseQueryStringSettings) SetSpecifyUrlPattern(v string) *QueryQueryStringUrlConfigResponseQueryStringSettings {
  s.SpecifyUrlPattern = &v
  return s
}

func (s *QueryQueryStringUrlConfigResponseQueryStringSettings) SetDirectories(v string) *QueryQueryStringUrlConfigResponseQueryStringSettings {
  s.Directories = &v
  return s
}

func (s *QueryQueryStringUrlConfigResponseQueryStringSettings) SetPriority(v int) *QueryQueryStringUrlConfigResponseQueryStringSettings {
  s.Priority = &v
  return s
}

func (s *QueryQueryStringUrlConfigResponseQueryStringSettings) SetIgnoreLetterCase(v string) *QueryQueryStringUrlConfigResponseQueryStringSettings {
  s.IgnoreLetterCase = &v
  return s
}

func (s *QueryQueryStringUrlConfigResponseQueryStringSettings) SetIgnoreQueryString(v string) *QueryQueryStringUrlConfigResponseQueryStringSettings {
  s.IgnoreQueryString = &v
  return s
}

func (s *QueryQueryStringUrlConfigResponseQueryStringSettings) SetQueryStringKept(v string) *QueryQueryStringUrlConfigResponseQueryStringSettings {
  s.QueryStringKept = &v
  return s
}

func (s *QueryQueryStringUrlConfigResponseQueryStringSettings) SetQueryStringRemoved(v string) *QueryQueryStringUrlConfigResponseQueryStringSettings {
  s.QueryStringRemoved = &v
  return s
}

func (s *QueryQueryStringUrlConfigResponseQueryStringSettings) SetSourceWithQuery(v string) *QueryQueryStringUrlConfigResponseQueryStringSettings {
  s.SourceWithQuery = &v
  return s
}

func (s *QueryQueryStringUrlConfigResponseQueryStringSettings) SetSourceKeyKept(v string) *QueryQueryStringUrlConfigResponseQueryStringSettings {
  s.SourceKeyKept = &v
  return s
}

func (s *QueryQueryStringUrlConfigResponseQueryStringSettings) SetSourceKeyRemoved(v string) *QueryQueryStringUrlConfigResponseQueryStringSettings {
  s.SourceKeyRemoved = &v
  return s
}

func (s *QueryQueryStringUrlConfigResponseQueryStringSettings) SetDataId(v int32) *QueryQueryStringUrlConfigResponseQueryStringSettings {
  s.DataId = &v
  return s
}

type QueryQueryStringUrlConfigPaths struct {
  // {"en":"Domain name or domain ID that needs to be queried for configuration.", "zh_CN":"需要查询配置的域名或域名id"}
  DomainName *string `json:"domain-name,omitempty" xml:"domain-name,omitempty"`
}

func (s QueryQueryStringUrlConfigPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryQueryStringUrlConfigPaths) GoString() string {
  return s.String()
}

func (s *QueryQueryStringUrlConfigPaths) SetDomainName(v string) *QueryQueryStringUrlConfigPaths {
  s.DomainName = &v
  return s
}

type QueryQueryStringUrlConfigParameters struct {
}

func (s QueryQueryStringUrlConfigParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryQueryStringUrlConfigParameters) GoString() string {
  return s.String()
}

type QueryQueryStringUrlConfigRequestHeader struct {
}

func (s QueryQueryStringUrlConfigRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryQueryStringUrlConfigRequestHeader) GoString() string {
  return s.String()
}

type QueryQueryStringUrlConfigResponseHeader struct {
}

func (s QueryQueryStringUrlConfigResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryQueryStringUrlConfigResponseHeader) GoString() string {
  return s.String()
}




type EnableOrDisableBotProtectionRequest struct {
  // {"en":"Domain names list, the parent tag.", "zh_CN":"开启/关闭Bot防护的域名列表， 父标签"}
  DomainNames []*string `json:"domainNames,omitempty" xml:"domainNames,omitempty" require:"true" type:"Repeated"`
  // {"en":"1: Enable Bot protection; 
  // 0: Disable Bot protection;", "zh_CN":"1：开启BOT防护，0：关闭BOT防护"}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
}

func (s EnableOrDisableBotProtectionRequest) String() string {
  return tea.Prettify(s)
}

func (s EnableOrDisableBotProtectionRequest) GoString() string {
  return s.String()
}

func (s *EnableOrDisableBotProtectionRequest) SetDomainNames(v []*string) *EnableOrDisableBotProtectionRequest {
  s.DomainNames = v
  return s
}

func (s *EnableOrDisableBotProtectionRequest) SetType(v string) *EnableOrDisableBotProtectionRequest {
  s.Type = &v
  return s
}

type EnableOrDisableBotProtectionResponse struct {
  // {"en":"Error code. 0 means successful", "zh_CN":"错误代码。 0：成功"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"The response message. Response success when calling API successful", "zh_CN":"响应信息，成功时为success"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"The body of return data.", "zh_CN":"返回体"}
  Data *string `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s EnableOrDisableBotProtectionResponse) String() string {
  return tea.Prettify(s)
}

func (s EnableOrDisableBotProtectionResponse) GoString() string {
  return s.String()
}

func (s *EnableOrDisableBotProtectionResponse) SetCode(v string) *EnableOrDisableBotProtectionResponse {
  s.Code = &v
  return s
}

func (s *EnableOrDisableBotProtectionResponse) SetMessage(v string) *EnableOrDisableBotProtectionResponse {
  s.Message = &v
  return s
}

func (s *EnableOrDisableBotProtectionResponse) SetData(v string) *EnableOrDisableBotProtectionResponse {
  s.Data = &v
  return s
}

type EnableOrDisableBotProtectionPaths struct {
}

func (s EnableOrDisableBotProtectionPaths) String() string {
  return tea.Prettify(s)
}

func (s EnableOrDisableBotProtectionPaths) GoString() string {
  return s.String()
}

type EnableOrDisableBotProtectionParameters struct {
}

func (s EnableOrDisableBotProtectionParameters) String() string {
  return tea.Prettify(s)
}

func (s EnableOrDisableBotProtectionParameters) GoString() string {
  return s.String()
}

type EnableOrDisableBotProtectionRequestHeader struct {
}

func (s EnableOrDisableBotProtectionRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s EnableOrDisableBotProtectionRequestHeader) GoString() string {
  return s.String()
}

type EnableOrDisableBotProtectionResponseHeader struct {
}

func (s EnableOrDisableBotProtectionResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s EnableOrDisableBotProtectionResponseHeader) GoString() string {
  return s.String()
}




type QueryDomainBillingareaRequest struct {
}

func (s QueryDomainBillingareaRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryDomainBillingareaRequest) GoString() string {
  return s.String()
}

type QueryDomainBillingareaResponse struct {
  // {"en":"Domain ID.", "zh_CN":"域名ID"}
  DomainId *int `json:"domainId,omitempty" xml:"domainId,omitempty" require:"true"`
  // {"en":"Domain name.", "zh_CN":"域名名称"}
  DomainName *string `json:"domainName,omitempty" xml:"domainName,omitempty" require:"true"`
  // {"en":"Billing areas.", "zh_CN":"计费区域"}
  BillingArea *string `json:"billingArea,omitempty" xml:"billingArea,omitempty" require:"true"`
}

func (s QueryDomainBillingareaResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryDomainBillingareaResponse) GoString() string {
  return s.String()
}

func (s *QueryDomainBillingareaResponse) SetDomainId(v int) *QueryDomainBillingareaResponse {
  s.DomainId = &v
  return s
}

func (s *QueryDomainBillingareaResponse) SetDomainName(v string) *QueryDomainBillingareaResponse {
  s.DomainName = &v
  return s
}

func (s *QueryDomainBillingareaResponse) SetBillingArea(v string) *QueryDomainBillingareaResponse {
  s.BillingArea = &v
  return s
}

type QueryDomainBillingareaPaths struct {
  // {"en":"The domain whoes need query config.", "zh_CN":"需要查询配置的域名或域名id"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
}

func (s QueryDomainBillingareaPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryDomainBillingareaPaths) GoString() string {
  return s.String()
}

func (s *QueryDomainBillingareaPaths) SetDomain(v string) *QueryDomainBillingareaPaths {
  s.Domain = &v
  return s
}

type QueryDomainBillingareaParameters struct {
}

func (s QueryDomainBillingareaParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryDomainBillingareaParameters) GoString() string {
  return s.String()
}

type QueryDomainBillingareaRequestHeader struct {
}

func (s QueryDomainBillingareaRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryDomainBillingareaRequestHeader) GoString() string {
  return s.String()
}

type QueryDomainBillingareaResponseHeader struct {
}

func (s QueryDomainBillingareaResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryDomainBillingareaResponseHeader) GoString() string {
  return s.String()
}




type QueryStreamNotificationConfigRequest struct {
}

func (s QueryStreamNotificationConfigRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryStreamNotificationConfigRequest) GoString() string {
  return s.String()
}

type QueryStreamNotificationConfigResponse struct {
  // {"en":"code", "zh_CN":"错误代码，当HTTPStatus不为202时出现，表示当前请求调用的错误类型"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"message", "zh_CN":"响应信息，成功时为success"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"data", "zh_CN":"返回结果数据"}
  Data *QueryStreamNotificationConfigResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
}

func (s QueryStreamNotificationConfigResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryStreamNotificationConfigResponse) GoString() string {
  return s.String()
}

func (s *QueryStreamNotificationConfigResponse) SetCode(v string) *QueryStreamNotificationConfigResponse {
  s.Code = &v
  return s
}

func (s *QueryStreamNotificationConfigResponse) SetMessage(v string) *QueryStreamNotificationConfigResponse {
  s.Message = &v
  return s
}

func (s *QueryStreamNotificationConfigResponse) SetData(v *QueryStreamNotificationConfigResponseData) *QueryStreamNotificationConfigResponse {
  s.Data = v
  return s
}

type QueryStreamNotificationConfigResponseData struct {
  // {"en":"domainId", "zh_CN":"域名ID"}
  DomainId *string `json:"domainId,omitempty" xml:"domainId,omitempty" require:"true"`
  // {"en":"domainName", "zh_CN":"域名名称"}
  DomainName *string `json:"domainName,omitempty" xml:"domainName,omitempty" require:"true"`
  // {"en":"Streaming notify Configuration, parent node", "zh_CN":"推流状态反馈配置，父标签"}
  StreamNotifications []*QueryStreamNotificationConfigResponseDataStreamNotifications `json:"streamNotifications,omitempty" xml:"streamNotifications,omitempty" require:"true" type:"Repeated"`
}

func (s QueryStreamNotificationConfigResponseData) String() string {
  return tea.Prettify(s)
}

func (s QueryStreamNotificationConfigResponseData) GoString() string {
  return s.String()
}

func (s *QueryStreamNotificationConfigResponseData) SetDomainId(v string) *QueryStreamNotificationConfigResponseData {
  s.DomainId = &v
  return s
}

func (s *QueryStreamNotificationConfigResponseData) SetDomainName(v string) *QueryStreamNotificationConfigResponseData {
  s.DomainName = &v
  return s
}

func (s *QueryStreamNotificationConfigResponseData) SetStreamNotifications(v []*QueryStreamNotificationConfigResponseDataStreamNotifications) *QueryStreamNotificationConfigResponseData {
  s.StreamNotifications = v
  return s
}

type QueryStreamNotificationConfigResponseDataStreamNotifications struct     {
  // {"en":"Switch for streaming notification feature. The optional values are true and false.", "zh_CN":"推流状态反馈开关配置，可选值为true和false。
  // 为true则开启，且需完整设置配置信息
  // 为false则其他入参无效。"}
  EnableStreamNotification *bool `json:"enableStreamNotification,omitempty" xml:"enableStreamNotification,omitempty" require:"true"`
  // {"en":"Notify address, format is protocol://{host:port}, the protocol is HTTP or HTTPS, for example:
  // http://www.exam.com 
  // http://www.exam.com:81
  // http://12.12.12.12
  // http://12.12.12.12:81", "zh_CN":"推流状态汇报地址，格式为：protocol://{host:port}，例如：
  // http://www.exam.com 
  // http://www.exam.com:81
  // http://12.12.12.12
  // http://12.12.12.12:81"}
  NotifyAddress *string `json:"notifyAddress,omitempty" xml:"notifyAddress,omitempty" require:"true"`
  // {"en":"Request method, support POST and GET.", "zh_CN":"推流汇报请求方式，支持POST和GET"}
  NotifyMethod *string `json:"notifyMethod,omitempty" xml:"notifyMethod,omitempty" require:"true"`
  // {"en":"notifyParams", "zh_CN":"设置推流开始的汇报参数，推流结束的汇报参数。最多只能有一个开始，一个结束。"}
  NotifyParams []*QueryStreamNotificationConfigResponseDataStreamNotificationsNotifyParams `json:"notifyParams,omitempty" xml:"notifyParams,omitempty" require:"true" type:"Repeated"`
  // {"en":"Data-id is to indicate a specific group configuration when the client has multiple groups of configurations.", "zh_CN":"配置多组配置时，具体某组配置的id。"}
  DataId *int64 `json:"dataId,omitempty" xml:"dataId,omitempty" require:"true"`
}

func (s QueryStreamNotificationConfigResponseDataStreamNotifications) String() string {
  return tea.Prettify(s)
}

func (s QueryStreamNotificationConfigResponseDataStreamNotifications) GoString() string {
  return s.String()
}

func (s *QueryStreamNotificationConfigResponseDataStreamNotifications) SetEnableStreamNotification(v bool) *QueryStreamNotificationConfigResponseDataStreamNotifications {
  s.EnableStreamNotification = &v
  return s
}

func (s *QueryStreamNotificationConfigResponseDataStreamNotifications) SetNotifyAddress(v string) *QueryStreamNotificationConfigResponseDataStreamNotifications {
  s.NotifyAddress = &v
  return s
}

func (s *QueryStreamNotificationConfigResponseDataStreamNotifications) SetNotifyMethod(v string) *QueryStreamNotificationConfigResponseDataStreamNotifications {
  s.NotifyMethod = &v
  return s
}

func (s *QueryStreamNotificationConfigResponseDataStreamNotifications) SetNotifyParams(v []*QueryStreamNotificationConfigResponseDataStreamNotificationsNotifyParams) *QueryStreamNotificationConfigResponseDataStreamNotifications {
  s.NotifyParams = v
  return s
}

func (s *QueryStreamNotificationConfigResponseDataStreamNotifications) SetDataId(v int64) *QueryStreamNotificationConfigResponseDataStreamNotifications {
  s.DataId = &v
  return s
}

type QueryStreamNotificationConfigResponseDataStreamNotificationsNotifyParams struct     {
  // {"en":"notifyType", "zh_CN":"可选值：publish_start、publish_end，分别表示推流开始和推流结束(断流)"}
  NotifyType *string `json:"notifyType,omitempty" xml:"notifyType,omitempty" require:"true"`
  // {"en":"Notify path, it should start with a slash", "zh_CN":"推流汇报路径，/ 开头"}
  Path *string `json:"path,omitempty" xml:"path,omitempty" require:"true"`
  // {"en":"id", "zh_CN":"流名"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {"en":"clientIp", "zh_CN":"推流的客户端IP，主播IP"}
  ClientIp *string `json:"clientIp,omitempty" xml:"clientIp,omitempty" require:"true"`
  // {"en":"nodeIp", "zh_CN":"CDN接受推流的节点IP"}
  NodeIp *string `json:"nodeIp,omitempty" xml:"nodeIp,omitempty" require:"true"`
  // {"en":"app", "zh_CN":"推流域名"}
  App *string `json:"app,omitempty" xml:"app,omitempty" require:"true"`
  // {"en":"appName", "zh_CN":"发布点"}
  AppName *string `json:"appName,omitempty" xml:"appName,omitempty" require:"true"`
  // {"en":"port", "zh_CN":"端口"}
  Port *string `json:"port,omitempty" xml:"port,omitempty" require:"true"`
  // {"en":"millTime", "zh_CN":"汇报时间（毫秒级）"}
  MillTime *string `json:"millTime,omitempty" xml:"millTime,omitempty" require:"true"`
  // {"en":"customParams", "zh_CN":"指定推流汇报要携带的参数，来自用户推流参数"}
  CustomParams *string `json:"customParams,omitempty" xml:"customParams,omitempty" require:"true"`
}

func (s QueryStreamNotificationConfigResponseDataStreamNotificationsNotifyParams) String() string {
  return tea.Prettify(s)
}

func (s QueryStreamNotificationConfigResponseDataStreamNotificationsNotifyParams) GoString() string {
  return s.String()
}

func (s *QueryStreamNotificationConfigResponseDataStreamNotificationsNotifyParams) SetNotifyType(v string) *QueryStreamNotificationConfigResponseDataStreamNotificationsNotifyParams {
  s.NotifyType = &v
  return s
}

func (s *QueryStreamNotificationConfigResponseDataStreamNotificationsNotifyParams) SetPath(v string) *QueryStreamNotificationConfigResponseDataStreamNotificationsNotifyParams {
  s.Path = &v
  return s
}

func (s *QueryStreamNotificationConfigResponseDataStreamNotificationsNotifyParams) SetId(v string) *QueryStreamNotificationConfigResponseDataStreamNotificationsNotifyParams {
  s.Id = &v
  return s
}

func (s *QueryStreamNotificationConfigResponseDataStreamNotificationsNotifyParams) SetClientIp(v string) *QueryStreamNotificationConfigResponseDataStreamNotificationsNotifyParams {
  s.ClientIp = &v
  return s
}

func (s *QueryStreamNotificationConfigResponseDataStreamNotificationsNotifyParams) SetNodeIp(v string) *QueryStreamNotificationConfigResponseDataStreamNotificationsNotifyParams {
  s.NodeIp = &v
  return s
}

func (s *QueryStreamNotificationConfigResponseDataStreamNotificationsNotifyParams) SetApp(v string) *QueryStreamNotificationConfigResponseDataStreamNotificationsNotifyParams {
  s.App = &v
  return s
}

func (s *QueryStreamNotificationConfigResponseDataStreamNotificationsNotifyParams) SetAppName(v string) *QueryStreamNotificationConfigResponseDataStreamNotificationsNotifyParams {
  s.AppName = &v
  return s
}

func (s *QueryStreamNotificationConfigResponseDataStreamNotificationsNotifyParams) SetPort(v string) *QueryStreamNotificationConfigResponseDataStreamNotificationsNotifyParams {
  s.Port = &v
  return s
}

func (s *QueryStreamNotificationConfigResponseDataStreamNotificationsNotifyParams) SetMillTime(v string) *QueryStreamNotificationConfigResponseDataStreamNotificationsNotifyParams {
  s.MillTime = &v
  return s
}

func (s *QueryStreamNotificationConfigResponseDataStreamNotificationsNotifyParams) SetCustomParams(v string) *QueryStreamNotificationConfigResponseDataStreamNotificationsNotifyParams {
  s.CustomParams = &v
  return s
}

type QueryStreamNotificationConfigPaths struct {
  // {"en":"Domain name or domain name id to query configuration", "zh_CN":"需要查询配置的域名（domainName）或域名id（domainId）"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
}

func (s QueryStreamNotificationConfigPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryStreamNotificationConfigPaths) GoString() string {
  return s.String()
}

func (s *QueryStreamNotificationConfigPaths) SetDomain(v string) *QueryStreamNotificationConfigPaths {
  s.Domain = &v
  return s
}

type QueryStreamNotificationConfigParameters struct {
}

func (s QueryStreamNotificationConfigParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryStreamNotificationConfigParameters) GoString() string {
  return s.String()
}

type QueryStreamNotificationConfigRequestHeader struct {
}

func (s QueryStreamNotificationConfigRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryStreamNotificationConfigRequestHeader) GoString() string {
  return s.String()
}

type QueryStreamNotificationConfigResponseHeader struct {
}

func (s QueryStreamNotificationConfigResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryStreamNotificationConfigResponseHeader) GoString() string {
  return s.String()
}




type UpdateDomainCertConfigRequest struct {
  // {"en":"The certificate ID you want to bind.", "zh_CN":"证书ID，为域名关联证书或换证书。请注意：置空时为取消关联证书。"}
  CertificateId *string `json:"certificateId,omitempty" xml:"certificateId,omitempty"`
  // {"en":"TLS version. Optional values: SSLv3,TLSv1,TLSv1.1,TLSv1.2,TLSv1.3", "zh_CN":"TLS协议（TLSVersion），支持配置多个，英文分号分隔。可选值为: SSLv3、TLSv1、TLSv1.1、TLSv1.2、TLSv1.3。域名有配置证书时，该配置才有效。"}
  TLSVersion *string `json:"TLSVersion,omitempty" xml:"TLSVersion,omitempty"`
  // {"en":"Enable OCSP(Online Certificate Status Protocol).", "zh_CN":"支持OCSP，默认不启用，可选值true、false。域名有配置证书时，该配置才有效。"}
  EnableOCSP *string `json:"enableOCSP,omitempty" xml:"enableOCSP,omitempty"`
  // {"en":"This optional object is used to specify a colon separated list of cipher suites which are permitted when clients negotiate security settings to access your content. Cipher suites which you can specify are: LOW, ALL:!LOW, HIGH, !EXPORT, !aNULL, !RC4, !DH, !SHA, !MD5, @STRENGTH,  AES128-SHA, AES256-SHA, AES128-SHA256, AES256-SHA256, AES128-GCM-SHA256, AES256-GCM-SHA384, ECDHE-RSA-AES128-SHA, ECDHE-RSA-AES256-SHA, ECDHE-RSA-AES128-SHA256, ECDHE-RSA-AES256-SHA384, ECDHE-RSA-AES128-GCM-SHA256, and ECDHE-RSA-AES256-GCM-SHA384. These cipher suites are a subset of those supported by OpenSSL, https://www.openssl.org/docs/man1.0.2/man1/ciphers.html. Please note that !MD5 or !SHA must appear after HIGH..", "zh_CN":"设置cipher加密套件，冒号分隔。请注意：系统已有默认的算法，如无特殊修改需要，无需调整。"}
  CipherSuites *string `json:"cipherSuites,omitempty" xml:"cipherSuites,omitempty"`
}

func (s UpdateDomainCertConfigRequest) String() string {
  return tea.Prettify(s)
}

func (s UpdateDomainCertConfigRequest) GoString() string {
  return s.String()
}

func (s *UpdateDomainCertConfigRequest) SetCertificateId(v string) *UpdateDomainCertConfigRequest {
  s.CertificateId = &v
  return s
}

func (s *UpdateDomainCertConfigRequest) SetTLSVersion(v string) *UpdateDomainCertConfigRequest {
  s.TLSVersion = &v
  return s
}

func (s *UpdateDomainCertConfigRequest) SetEnableOCSP(v string) *UpdateDomainCertConfigRequest {
  s.EnableOCSP = &v
  return s
}

func (s *UpdateDomainCertConfigRequest) SetCipherSuites(v string) *UpdateDomainCertConfigRequest {
  s.CipherSuites = &v
  return s
}

type UpdateDomainCertConfigResponse struct {
  // {"en":"Error code, which appears when HTTPStatus is not 202, represents the error type of the current request call", "zh_CN":"错误代码，当HTTPStatus不为202时出现，表示当前请求调用的错误类型"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response information, success when successful", "zh_CN":"响应信息，成功时为success"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s UpdateDomainCertConfigResponse) String() string {
  return tea.Prettify(s)
}

func (s UpdateDomainCertConfigResponse) GoString() string {
  return s.String()
}

func (s *UpdateDomainCertConfigResponse) SetCode(v string) *UpdateDomainCertConfigResponse {
  s.Code = &v
  return s
}

func (s *UpdateDomainCertConfigResponse) SetMessage(v string) *UpdateDomainCertConfigResponse {
  s.Message = &v
  return s
}

type UpdateDomainCertConfigPaths struct {
  // {"en":"The domain whoes need query config.", "zh_CN":"需要查询配置的域名或域名id"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
}

func (s UpdateDomainCertConfigPaths) String() string {
  return tea.Prettify(s)
}

func (s UpdateDomainCertConfigPaths) GoString() string {
  return s.String()
}

func (s *UpdateDomainCertConfigPaths) SetDomain(v string) *UpdateDomainCertConfigPaths {
  s.Domain = &v
  return s
}

type UpdateDomainCertConfigParameters struct {
}

func (s UpdateDomainCertConfigParameters) String() string {
  return tea.Prettify(s)
}

func (s UpdateDomainCertConfigParameters) GoString() string {
  return s.String()
}

type UpdateDomainCertConfigRequestHeader struct {
}

func (s UpdateDomainCertConfigRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s UpdateDomainCertConfigRequestHeader) GoString() string {
  return s.String()
}

type UpdateDomainCertConfigResponseHeader struct {
}

func (s UpdateDomainCertConfigResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s UpdateDomainCertConfigResponseHeader) GoString() string {
  return s.String()
}




type EditLivestreamingTimestampAntihotlinkingConfigRequest struct {
  // {"en":"Streaming timestamp visit control configuration, parent tag
  // 1. This must be filled when the hotlinking configuration of streaming media needs to be set
  // 2. Empty the configuration for <timestampVisitControls/>", "zh_CN":"流媒体时间戳防盗链配置，父标签
  // 1.需要设置流媒体时间戳防盗链配置时，此项必填
  // 2.为<timestampVisitControls/>时清空配置"}
  TimestampVisitControlRules []*EditLivestreamingTimestampAntihotlinkingConfigRequestTimestampVisitControlRules `json:"timestampVisitControlRules,omitempty" xml:"timestampVisitControlRules,omitempty" require:"true" type:"Repeated"`
}

func (s EditLivestreamingTimestampAntihotlinkingConfigRequest) String() string {
  return tea.Prettify(s)
}

func (s EditLivestreamingTimestampAntihotlinkingConfigRequest) GoString() string {
  return s.String()
}

func (s *EditLivestreamingTimestampAntihotlinkingConfigRequest) SetTimestampVisitControlRules(v []*EditLivestreamingTimestampAntihotlinkingConfigRequestTimestampVisitControlRules) *EditLivestreamingTimestampAntihotlinkingConfigRequest {
  s.TimestampVisitControlRules = v
  return s
}

type EditLivestreamingTimestampAntihotlinkingConfigRequestTimestampVisitControlRules struct     {
  // {"en":"The name of the time parameter, and the default value is wsTime if no parameters are entered.", "zh_CN":"密文参数名称，未入参则默认值为wsSecret"}
  CipherParam *string `json:"cipherParam,omitempty" xml:"cipherParam,omitempty" require:"true"`
  // {"en":"The name of the time parameter, and the default value is wsTime if no parameters are entered.", "zh_CN":"时间参数名称，未入参则默认值为wsTime"}
  TimeParam *string `json:"timeParam,omitempty" xml:"timeParam,omitempty" require:"true"`
  // {"en":"The secret key", "zh_CN":"密钥"}
  SecretKey *string `json:"secretKey,omitempty" xml:"secretKey,omitempty" require:"true"`
  // {"en":"The encrypted time format, required. Opional values: [UNIX timestamp] and [hex], corresponding to [UNIX timestamp] and [hex], respectively.", "zh_CN":"加密时间格式，必填，可选值为【UNIX timestamp】和【hex】，分别为【UNIX时间戳】、【16进制】。"}
  TimeFormat *string `json:"timeFormat,omitempty" xml:"timeFormat,omitempty" require:"true"`
  // {"en":"Effective time method, optional values: duration, abstime, no verification.
  // 1. duration, means get effective time by duration. When this mode is set, effectiveTime is required.
  // 2. abstime, means get effective time by absolute time. Support a URL to be valid for a certain future time (the maximum 3652 days)
  // 3. No verification, means it will not verify the time. When this mode is set, effectiveTime will be cleared cleaned.", "zh_CN":"有效时间计算方式，可选值：duration、abstime、no verification
  // 1. duration，表示按时长。设置为“按时长”，则effectiveTime必填。
  // 2. abstime，表示按绝对时间。支持某个url到某个未来时间有效（最大支持3652天）
  // 3 .no verification，表示不校验时间。则自动清空effectiveTime。"}
  EffectiveTimeMode *string `json:"effectiveTimeMode,omitempty" xml:"effectiveTimeMode,omitempty"`
  // {"en":"The effective length, unit second (s), positive integer,and value >=1.
  // When effectiveTimeMode equals 'duration', this item is required.", "zh_CN":"有效时长，单位秒（s），正整数，值要>=1。
  // 有效时间计算方式为“按时长”时，该字段必填。"}
  EffectiveTime *int `json:"effectiveTime,omitempty" xml:"effectiveTime,omitempty"`
  // {"en":"The tolerance value, error time, unit second (s), positive integer and value >=1", "zh_CN":"容忍值，误差时间，单位秒（s），正整数，值要>=1
  // 1. 当客户时间比我司的时间慢10s时，无需配置容忍，但是这样的话用户就会少看10秒
  // 2. 当客户时间比我司的时间快10s时，需要配置容忍为10 ，但这样的话这用户也会少看了10秒"}
  TolerantTime *int `json:"tolerantTime,omitempty" xml:"tolerantTime,omitempty"`
  // {"en":"Ciphertext combination mode, with the following six optional values:
  // 
  // key+path+time
  // key+time+path
  // path+key+time
  // path+time+key
  // time+path+key
  // time+key+path", "zh_CN":"密文组合方式，可选值为以下六项：
  // 
  // key+path+time
  // key+time+path
  // path+key+time
  // path+time+key
  // time+path+key
  // time+key+path"}
  CipheCombination *string `json:"cipheCombination,omitempty" xml:"cipheCombination,omitempty" require:"true"`
  // {"en":"Data-id is to indicate a specific group configuration when the client has multiple groups of configurations. Data-id can be retrieved through a query interface. Note: A. If data-id is passed, it means that one group of configuration items is specified to be modified, and no other group configuration items need to be modified. B. If multiple groups of configurations are included, some of them are configured with data-id and others are not, then the expression of data-id is used to modify a specific group of configurations, and a new group of configurations is added on the original basis without the expression of data-id. C. If the data-id is not transmitted, it means that the original configuration will be fully covered by this configuration. D. If no configuration parameter is passed, only domain name and secondary label are passed, which means that all configuration of domain name secondary service corresponding to this interface is cleared. E. If there is no specific configuration item in a set of configurations, the data-id must be filled in, and the value is the actual data-id, which means clearing the value of the corresponding data-id configuration item; it is not allowed that there is no specific configuration item or data-id in a set of configurations.", "zh_CN":"配置多组配置时，具体某组配置的id。dataId可以通过查询接口获取。 注意： a、如果有传dataId，说明指定修改其中一组配置项内容，不需求修改其他组配置内容不需要入参； b、如果入参多组配置，其中有些组配置有传dataId，有些没有传，则有传dataId的表示修改具体某组配置，没有传dataId的表示在原来基础上新增一组配置； c、如果入参都没有传dataId,表示用本次的配置全量覆盖原先配置； d、如果入参没有传任何配置项参数，只传了域名和二级标签，表示清空这个接口对应域名二级服务所有配置； e、如果一组配置没有具体的配置项，则dataId必填，且值为实际存在的dataId，表示清空这个dataId对应配置项的值；不允许一组配置没有具体的配置项也没有dataId。"}
  DataId *int `json:"dataId,omitempty" xml:"dataId,omitempty"`
}

func (s EditLivestreamingTimestampAntihotlinkingConfigRequestTimestampVisitControlRules) String() string {
  return tea.Prettify(s)
}

func (s EditLivestreamingTimestampAntihotlinkingConfigRequestTimestampVisitControlRules) GoString() string {
  return s.String()
}

func (s *EditLivestreamingTimestampAntihotlinkingConfigRequestTimestampVisitControlRules) SetCipherParam(v string) *EditLivestreamingTimestampAntihotlinkingConfigRequestTimestampVisitControlRules {
  s.CipherParam = &v
  return s
}

func (s *EditLivestreamingTimestampAntihotlinkingConfigRequestTimestampVisitControlRules) SetTimeParam(v string) *EditLivestreamingTimestampAntihotlinkingConfigRequestTimestampVisitControlRules {
  s.TimeParam = &v
  return s
}

func (s *EditLivestreamingTimestampAntihotlinkingConfigRequestTimestampVisitControlRules) SetSecretKey(v string) *EditLivestreamingTimestampAntihotlinkingConfigRequestTimestampVisitControlRules {
  s.SecretKey = &v
  return s
}

func (s *EditLivestreamingTimestampAntihotlinkingConfigRequestTimestampVisitControlRules) SetTimeFormat(v string) *EditLivestreamingTimestampAntihotlinkingConfigRequestTimestampVisitControlRules {
  s.TimeFormat = &v
  return s
}

func (s *EditLivestreamingTimestampAntihotlinkingConfigRequestTimestampVisitControlRules) SetEffectiveTimeMode(v string) *EditLivestreamingTimestampAntihotlinkingConfigRequestTimestampVisitControlRules {
  s.EffectiveTimeMode = &v
  return s
}

func (s *EditLivestreamingTimestampAntihotlinkingConfigRequestTimestampVisitControlRules) SetEffectiveTime(v int) *EditLivestreamingTimestampAntihotlinkingConfigRequestTimestampVisitControlRules {
  s.EffectiveTime = &v
  return s
}

func (s *EditLivestreamingTimestampAntihotlinkingConfigRequestTimestampVisitControlRules) SetTolerantTime(v int) *EditLivestreamingTimestampAntihotlinkingConfigRequestTimestampVisitControlRules {
  s.TolerantTime = &v
  return s
}

func (s *EditLivestreamingTimestampAntihotlinkingConfigRequestTimestampVisitControlRules) SetCipheCombination(v string) *EditLivestreamingTimestampAntihotlinkingConfigRequestTimestampVisitControlRules {
  s.CipheCombination = &v
  return s
}

func (s *EditLivestreamingTimestampAntihotlinkingConfigRequestTimestampVisitControlRules) SetDataId(v int) *EditLivestreamingTimestampAntihotlinkingConfigRequestTimestampVisitControlRules {
  s.DataId = &v
  return s
}

type EditLivestreamingTimestampAntihotlinkingConfigResponse struct {
  // {"en":"The error code", "zh_CN":"错误码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"The message body", "zh_CN":"消息体"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Returns the body of the data.", "zh_CN":"返回数据体。"}
  Data *string `json:"data,omitempty" xml:"data,omitempty" require:"true"`
  // {"en":"Uniquely labeled id for querying each requested task (for all interfaces)", "zh_CN":"唯一标示的id，用于查询每次请求的任务 （适用全部接口）"}
  XCncRequestId *string `json:"x-cnc-request-id,omitempty" xml:"x-cnc-request-id,omitempty" require:"true"`
}

func (s EditLivestreamingTimestampAntihotlinkingConfigResponse) String() string {
  return tea.Prettify(s)
}

func (s EditLivestreamingTimestampAntihotlinkingConfigResponse) GoString() string {
  return s.String()
}

func (s *EditLivestreamingTimestampAntihotlinkingConfigResponse) SetCode(v string) *EditLivestreamingTimestampAntihotlinkingConfigResponse {
  s.Code = &v
  return s
}

func (s *EditLivestreamingTimestampAntihotlinkingConfigResponse) SetMessage(v string) *EditLivestreamingTimestampAntihotlinkingConfigResponse {
  s.Message = &v
  return s
}

func (s *EditLivestreamingTimestampAntihotlinkingConfigResponse) SetData(v string) *EditLivestreamingTimestampAntihotlinkingConfigResponse {
  s.Data = &v
  return s
}

func (s *EditLivestreamingTimestampAntihotlinkingConfigResponse) SetXCncRequestId(v string) *EditLivestreamingTimestampAntihotlinkingConfigResponse {
  s.XCncRequestId = &v
  return s
}

type EditLivestreamingTimestampAntihotlinkingConfigPaths struct {
  // {"en":"The domain whoes need query config.", "zh_CN":"需要查询配置的域名或域名id"}
  DomainName *string `json:"domain-name,omitempty" xml:"domain-name,omitempty"`
}

func (s EditLivestreamingTimestampAntihotlinkingConfigPaths) String() string {
  return tea.Prettify(s)
}

func (s EditLivestreamingTimestampAntihotlinkingConfigPaths) GoString() string {
  return s.String()
}

func (s *EditLivestreamingTimestampAntihotlinkingConfigPaths) SetDomainName(v string) *EditLivestreamingTimestampAntihotlinkingConfigPaths {
  s.DomainName = &v
  return s
}

type EditLivestreamingTimestampAntihotlinkingConfigParameters struct {
}

func (s EditLivestreamingTimestampAntihotlinkingConfigParameters) String() string {
  return tea.Prettify(s)
}

func (s EditLivestreamingTimestampAntihotlinkingConfigParameters) GoString() string {
  return s.String()
}

type EditLivestreamingTimestampAntihotlinkingConfigRequestHeader struct {
}

func (s EditLivestreamingTimestampAntihotlinkingConfigRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s EditLivestreamingTimestampAntihotlinkingConfigRequestHeader) GoString() string {
  return s.String()
}

type EditLivestreamingTimestampAntihotlinkingConfigResponseHeader struct {
}

func (s EditLivestreamingTimestampAntihotlinkingConfigResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s EditLivestreamingTimestampAntihotlinkingConfigResponseHeader) GoString() string {
  return s.String()
}




type QueryOriginUriAndHostRequest struct {
}

func (s QueryOriginUriAndHostRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryOriginUriAndHostRequest) GoString() string {
  return s.String()
}

type QueryOriginUriAndHostResponse struct {
  // {"en":"domainName.", "zh_CN":"域名"}
  DomainName *string `json:"domain-name,omitempty" xml:"domain-name,omitempty" require:"true"`
  // {"en":"domainId.", "zh_CN":"域名ID"}
  DomainId *string `json:"domain-id,omitempty" xml:"domain-id,omitempty" require:"true"`
  // {"en":"Return path rewrite configuration
  // 
  // 1. When you need to set the rewrite configuration of the backsource path, this must be filled in
  // 
  // 2. Rewrite configuration for clearing the return path for <origin-rules-rewrites/>.", "zh_CN":"回源路径改写配置
  // 1.需要设置回源路径改写配置时，此项必填
  // 2.为<origin-rules-rewrites/>时清空回源路径改写配置"}
  OriginRulesRewrites []*QueryOriginUriAndHostResponseOriginRulesRewrites `json:"originRulesRewrites,omitempty" xml:"originRulesRewrites,omitempty" require:"true" type:"Repeated"`
}

func (s QueryOriginUriAndHostResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryOriginUriAndHostResponse) GoString() string {
  return s.String()
}

func (s *QueryOriginUriAndHostResponse) SetDomainName(v string) *QueryOriginUriAndHostResponse {
  s.DomainName = &v
  return s
}

func (s *QueryOriginUriAndHostResponse) SetDomainId(v string) *QueryOriginUriAndHostResponse {
  s.DomainId = &v
  return s
}

func (s *QueryOriginUriAndHostResponse) SetOriginRulesRewrites(v []*QueryOriginUriAndHostResponseOriginRulesRewrites) *QueryOriginUriAndHostResponse {
  s.OriginRulesRewrites = v
  return s
}

type QueryOriginUriAndHostResponseOriginRulesRewrites struct     {
  // {"en":"Add a grid type identifier to represent a specific group of configurations when a customer has multiple groups of configurations
  // 
  // Note: Add grid type identifier: data-id, each group configuration corresponds to a data-id: a. If the customer has passed data-id, specify that modifying one group of configuration items content does not require modifying other group configuration content does not need to be included; B. If the customer enters multiple groups of configuration, some of them have data-id, some have not. If there is transmission, the expression of data-id is used to modify a specific group of configurations, but no expression of data-id is used to add a new group of configurations on the original basis; C. If no data-id is transmitted to the customer, it means that the original configuration is completely covered by this configuration; D. If no configuration parameters are transmitted to the customer, only the domain name and the second level are transmitted. Label, which indicates that clearing this interface corresponds to all configuration of domain name secondary service. (c, D content is consistent with the current solution); e, a gird tag can not be empty, if there is no specific configuration item, then data-id must be filled in, and the value is the actual data-id, indicating the value of clearing this data-id corresponding configuration item;", "zh_CN":"添加grid类型标识，表示客户多组配置时，具体某组配置
  // 注意：添加grid类型标识：data-id，每一组配置对应一个data-id：a、如果客户有传data-id，说明指定修改其中一组配置项内容，不需求修改其他组配置内容不需要入参；b、如果客户入参多组配置，其中有些组配置有传data-id，有些没有传，则有传data-id的表示修改具体某组配置，没有传data-id的表示在原来基础上新增一组配置；c、如果客户入参都没有传data-id,表示用本次的配置全量覆盖原先配置；d、如果客户入参没有传任何配置项参数，只传了域名和二级标签，表示清空这个接口对应域名二级服务所有配置。（c、d内容和当前方案实现一致）；e、一个gird标签下的入参不能为空，如果，没有具体的配置项，则data-id必填，且值为实际存在的data-id,表示清空这个data-id对应配置项的值；"}
  DataId *int64 `json:"dataId,omitempty" xml:"dataId,omitempty" require:"true"`
  // {"en":"The URL matching mode supports regularization. If all matches are made, the input can be configured as:.*", "zh_CN":"url匹配模式，支持正则，如果是全部匹配，入参可以配置为：.*"}
  PathPattern *string `json:"pathPattern,omitempty" xml:"pathPattern,omitempty" require:"true"`
  // {"en":"The protocol of URL matching mode, which is used with path-pattern, is supported by adding parameters: default is empty, blank is default is to support HTTP and HTTPS protocol before URL matching mode path is needed at the same time;
  // 
  // Http: URL matches pattern path with HTTP protocol
  // 
  // Https: URL matches pattern path with HTTPS protocol
  // 
  // Ignore: URL matching mode path without protocol", "zh_CN":"url匹配模式的协议，该配置项与path-pattern搭配使用；入参支持：默认为空，为空则默认为需要同时支持url匹配模式路径前支持http和https协议；
  // http：url匹配模式路径前加上http协议
  // https:url 匹配模式路径前加上HTTPS协议
  // ignore:url匹配模式路径前不加协议"}
  PathPatternHttp *string `json:"pathPatternHttp,omitempty" xml:"pathPatternHttp,omitempty" require:"true"`
  // {"en":"Exceptional URL matching pattern in the same format as path pattern", "zh_CN":"例外的url匹配模式，格式同pathPattern"}
  ExceptPathPattern *string `json:"exceptPathPattern,omitempty" xml:"exceptPathPattern,omitempty" require:"true"`
  // {"en":"Exceptional URL matching mode protocol, which is used in conjunction with except-path-pattern; participation support: default is empty, blank is default is required to support both HTTP and HTTPS protocol before URL matching mode path;
  // 
  // Http: URL matches pattern path with HTTP protocol.
  // Https: URL matches pattern path with HTTPS protocol.
  // Ignore: URL matching mode path without protocol.", "zh_CN":"例外的url匹配模式的协议，该配置项与except-path-pattern搭配使用；入参支持：默认为空，为空则默认为需要同时支持url匹配模式路径前支持http和https协议；
  // http：url匹配模式路径前加上http协议
  // https:url 匹配模式路径前加上HTTPS协议
  // ignore:url匹配模式路径前不加协议"}
  ExceptPathPatternHttp *string `json:"exceptPathPatternHttp,omitempty" xml:"exceptPathPatternHttp,omitempty" require:"true"`
  // {"en":"Ignore case: Allowed values are true and false, the default is ignored.", "zh_CN":"是否忽略大小写：允许值为true和false，默认为忽略"}
  IgnoreLetterCase *bool `json:"ignoreLetterCase,omitempty" xml:"ignoreLetterCase,omitempty" require:"true"`
  // {"en":"Back-source information, you can enter IP or domain name.
  // That is, customer source IP or domain name", "zh_CN":"回源信息，可以输入ip或者域名
  // 即客户源站IP或域名"}
  OriginInfo *string `json:"originInfo,omitempty" xml:"originInfo,omitempty" require:"true"`
  // {"en":"Represents the priority execution order of the customer's multi-group redirected content. The bigger the number, the higher the priority.", "zh_CN":"表示客户多组重定向内容的优先执行顺序。数字越大，优先级越高。
  // 新增配置项时，不传默认为 10
  // 如果传了值，不能为空"}
  Priority *int32 `json:"priority,omitempty" xml:"priority,omitempty" require:"true"`
  // {"en":"Back source host,that support to enter domain names;
  // Example: The backsource host of domain name A configures domain name B. When the A domain name requests the source, the requested URL uses the B domain name instead of the A domain name.", "zh_CN":"回源host,支持入参域名；
  // 示例：域名A的回源host配置了域名B。当A域名请求的回源的时候，请求的url上用B域名代替A域名&nbsp;"}
  OriginHost *string `json:"originHost,omitempty" xml:"originHost,omitempty" require:"true"`
  // {"en":"Pre-rewrite uri. That is, the original request URI for user access. Support regular configuration", "zh_CN":"改写前的uri.&nbsp;即用户访问的原始请求uri&nbsp;。支持正则配置"}
  BeforeRewritedUri *string `json:"beforeRewritedUri,omitempty" xml:"beforeRewritedUri,omitempty" require:"true"`
  // {"en":"The rewritten uri, that is, the request uri configured in before-rewritten-uri, uses the rewritten uri to return to the source. Realize the rewriting of the back-to-source path. Support regular configuration.", "zh_CN":"改写后的uri,即将before-rewrited-uri配置的请求uri，用改写后的uri回源。实现回源路径改写。支持正则配置"}
  AfterRewritedUri *string `json:"afterRewritedUri,omitempty" xml:"afterRewritedUri,omitempty" require:"true"`
}

func (s QueryOriginUriAndHostResponseOriginRulesRewrites) String() string {
  return tea.Prettify(s)
}

func (s QueryOriginUriAndHostResponseOriginRulesRewrites) GoString() string {
  return s.String()
}

func (s *QueryOriginUriAndHostResponseOriginRulesRewrites) SetDataId(v int64) *QueryOriginUriAndHostResponseOriginRulesRewrites {
  s.DataId = &v
  return s
}

func (s *QueryOriginUriAndHostResponseOriginRulesRewrites) SetPathPattern(v string) *QueryOriginUriAndHostResponseOriginRulesRewrites {
  s.PathPattern = &v
  return s
}

func (s *QueryOriginUriAndHostResponseOriginRulesRewrites) SetPathPatternHttp(v string) *QueryOriginUriAndHostResponseOriginRulesRewrites {
  s.PathPatternHttp = &v
  return s
}

func (s *QueryOriginUriAndHostResponseOriginRulesRewrites) SetExceptPathPattern(v string) *QueryOriginUriAndHostResponseOriginRulesRewrites {
  s.ExceptPathPattern = &v
  return s
}

func (s *QueryOriginUriAndHostResponseOriginRulesRewrites) SetExceptPathPatternHttp(v string) *QueryOriginUriAndHostResponseOriginRulesRewrites {
  s.ExceptPathPatternHttp = &v
  return s
}

func (s *QueryOriginUriAndHostResponseOriginRulesRewrites) SetIgnoreLetterCase(v bool) *QueryOriginUriAndHostResponseOriginRulesRewrites {
  s.IgnoreLetterCase = &v
  return s
}

func (s *QueryOriginUriAndHostResponseOriginRulesRewrites) SetOriginInfo(v string) *QueryOriginUriAndHostResponseOriginRulesRewrites {
  s.OriginInfo = &v
  return s
}

func (s *QueryOriginUriAndHostResponseOriginRulesRewrites) SetPriority(v int32) *QueryOriginUriAndHostResponseOriginRulesRewrites {
  s.Priority = &v
  return s
}

func (s *QueryOriginUriAndHostResponseOriginRulesRewrites) SetOriginHost(v string) *QueryOriginUriAndHostResponseOriginRulesRewrites {
  s.OriginHost = &v
  return s
}

func (s *QueryOriginUriAndHostResponseOriginRulesRewrites) SetBeforeRewritedUri(v string) *QueryOriginUriAndHostResponseOriginRulesRewrites {
  s.BeforeRewritedUri = &v
  return s
}

func (s *QueryOriginUriAndHostResponseOriginRulesRewrites) SetAfterRewritedUri(v string) *QueryOriginUriAndHostResponseOriginRulesRewrites {
  s.AfterRewritedUri = &v
  return s
}

type QueryOriginUriAndHostPaths struct {
  // {"en":"domainName or domainId.", "zh_CN":"域名或者域名ID"}
  DomainName *string `json:"domainName,omitempty" xml:"domainName,omitempty" require:"true"`
}

func (s QueryOriginUriAndHostPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryOriginUriAndHostPaths) GoString() string {
  return s.String()
}

func (s *QueryOriginUriAndHostPaths) SetDomainName(v string) *QueryOriginUriAndHostPaths {
  s.DomainName = &v
  return s
}

type QueryOriginUriAndHostParameters struct {
}

func (s QueryOriginUriAndHostParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryOriginUriAndHostParameters) GoString() string {
  return s.String()
}

type QueryOriginUriAndHostRequestHeader struct {
}

func (s QueryOriginUriAndHostRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryOriginUriAndHostRequestHeader) GoString() string {
  return s.String()
}

type QueryOriginUriAndHostResponseHeader struct {
}

func (s QueryOriginUriAndHostResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryOriginUriAndHostResponseHeader) GoString() string {
  return s.String()
}




type UpdateCompressionConfigRequest struct {
  // {"en":"Compress setting config", "zh_CN":"压缩响应功能配置
  // 1.需要设置压缩响应配置时，此项必填
  // 2.为空<compression-settings/>时清空压缩响应配置"}
  CompressionSettings *UpdateCompressionConfigRequestCompressionSettings `json:"compression-settings,omitempty" xml:"compression-settings,omitempty" require:"true" type:"Struct"`
}

func (s UpdateCompressionConfigRequest) String() string {
  return tea.Prettify(s)
}

func (s UpdateCompressionConfigRequest) GoString() string {
  return s.String()
}

func (s *UpdateCompressionConfigRequest) SetCompressionSettings(v *UpdateCompressionConfigRequestCompressionSettings) *UpdateCompressionConfigRequest {
  s.CompressionSettings = v
  return s
}

type UpdateCompressionConfigRequestCompressionSettings struct {
  // {"en":"To enable compress setting, allowed true or false.", "zh_CN":"开启压缩响应功能：允许值为true和false"}
  CompressionEnabled *string `json:"compression-enabled,omitempty" xml:"compression-enabled,omitempty" require:"true"`
  // {"en":"The url matching mode. If all matches, the input parameters can be configured as: .*", "zh_CN":"url匹配模式，支持正则，如果是全部匹配，入参可以配置为：.*"}
  PathPattern *string `json:"path-pattern,omitempty" xml:"path-pattern,omitempty" require:"true"`
  // {"en":"Whether to ignore letter case.", "zh_CN":"是否忽略大小写：允许值为true和false"}
  IgnoreLetterCase *string `json:"ignore-letter-case,omitempty" xml:"ignore-letter-case,omitempty"`
  // {"en":"Define the file types to be compressed. 'text/' will be compressed by default.", "zh_CN":"配置需要压缩的文件类型，默认只对'text'文件类型压缩，配置为*时压缩任意文件类型"}
  FileTypes []*string `json:"file-types,omitempty" xml:"file-types,omitempty" type:"Repeated"`
  // {"en":"Use br compression.The allowed values are true and false.", "zh_CN":"是否使用br压缩：允许值为true和false"}
  BrTypes *string `json:"br-types,omitempty" xml:"br-types,omitempty"`
}

func (s UpdateCompressionConfigRequestCompressionSettings) String() string {
  return tea.Prettify(s)
}

func (s UpdateCompressionConfigRequestCompressionSettings) GoString() string {
  return s.String()
}

func (s *UpdateCompressionConfigRequestCompressionSettings) SetCompressionEnabled(v string) *UpdateCompressionConfigRequestCompressionSettings {
  s.CompressionEnabled = &v
  return s
}

func (s *UpdateCompressionConfigRequestCompressionSettings) SetPathPattern(v string) *UpdateCompressionConfigRequestCompressionSettings {
  s.PathPattern = &v
  return s
}

func (s *UpdateCompressionConfigRequestCompressionSettings) SetIgnoreLetterCase(v string) *UpdateCompressionConfigRequestCompressionSettings {
  s.IgnoreLetterCase = &v
  return s
}

func (s *UpdateCompressionConfigRequestCompressionSettings) SetFileTypes(v []*string) *UpdateCompressionConfigRequestCompressionSettings {
  s.FileTypes = v
  return s
}

func (s *UpdateCompressionConfigRequestCompressionSettings) SetBrTypes(v string) *UpdateCompressionConfigRequestCompressionSettings {
  s.BrTypes = &v
  return s
}

type UpdateCompressionConfigResponse struct {
  // {"en":"httpstatus=202; Indicates that the new domain API was successfully invoked, and the current deployment of the new domain can be viewed using x-cnc-request-id in the header", "zh_CN":"httpstatus=202;   表示成功调用新增域名接口，可使用header中的x-cnc-request-id查看当前新增域名的部署情况"}
  HttpStatus *int `json:"http status code,omitempty" xml:"http status code,omitempty" require:"true"`
  // {"en":"Uniquely labeled id for querying each requested task (for all interfaces)", "zh_CN":"唯一标示的id，用于查询每次请求的任务 （适用全部接口）"}
  XCncRequestId *string `json:"x-cnc-request-id,omitempty" xml:"x-cnc-request-id,omitempty" require:"true"`
  // {"en":"The error code, when HTTPStatus is not 202, indicates the type of error the current request is calling.", "zh_CN":"错误代码，当HTTPStatus不为202时出现，表示当前请求调用的错误类型"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response information, when success is successful", "zh_CN":"响应信息，成功时为success"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s UpdateCompressionConfigResponse) String() string {
  return tea.Prettify(s)
}

func (s UpdateCompressionConfigResponse) GoString() string {
  return s.String()
}

func (s *UpdateCompressionConfigResponse) SetHttpStatus(v int) *UpdateCompressionConfigResponse {
  s.HttpStatus = &v
  return s
}

func (s *UpdateCompressionConfigResponse) SetXCncRequestId(v string) *UpdateCompressionConfigResponse {
  s.XCncRequestId = &v
  return s
}

func (s *UpdateCompressionConfigResponse) SetCode(v string) *UpdateCompressionConfigResponse {
  s.Code = &v
  return s
}

func (s *UpdateCompressionConfigResponse) SetMessage(v string) *UpdateCompressionConfigResponse {
  s.Message = &v
  return s
}

type UpdateCompressionConfigPaths struct {
  // {"en":"The domain whoes need query config.", "zh_CN":"需要查询配置的域名或域名id"}
  DomainName *string `json:"domain-name,omitempty" xml:"domain-name,omitempty"`
}

func (s UpdateCompressionConfigPaths) String() string {
  return tea.Prettify(s)
}

func (s UpdateCompressionConfigPaths) GoString() string {
  return s.String()
}

func (s *UpdateCompressionConfigPaths) SetDomainName(v string) *UpdateCompressionConfigPaths {
  s.DomainName = &v
  return s
}

type UpdateCompressionConfigParameters struct {
}

func (s UpdateCompressionConfigParameters) String() string {
  return tea.Prettify(s)
}

func (s UpdateCompressionConfigParameters) GoString() string {
  return s.String()
}

type UpdateCompressionConfigRequestHeader struct {
}

func (s UpdateCompressionConfigRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s UpdateCompressionConfigRequestHeader) GoString() string {
  return s.String()
}

type UpdateCompressionConfigResponseHeader struct {
}

func (s UpdateCompressionConfigResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s UpdateCompressionConfigResponseHeader) GoString() string {
  return s.String()
}




type QuerytimecontrolServiceRequest struct {
}

func (s QuerytimecontrolServiceRequest) String() string {
  return tea.Prettify(s)
}

func (s QuerytimecontrolServiceRequest) GoString() string {
  return s.String()
}

type QuerytimecontrolServiceResponse struct {
  // {"en":"httpstatus=202; Indicates that the new domain API was successfully invoked, and the current deployment of the new domain can be viewed using x-cnc-request-id in the header", "zh_CN":"httpstatus=202;   表示成功调用新增域名接口，可使用header中的x-cnc-request-id查看当前新增域名的部署情况"}
  HttpStatus *int `json:"http status code,omitempty" xml:"http status code,omitempty" require:"true"`
  // {"en":"Uniquely identified id for querying tasks per request (for all API)", "zh_CN":"唯一标示的id，用于查询每次请求的任务 （适用全部接口）"}
  XCncRequestId *string `json:"x-cnc-request-id,omitempty" xml:"x-cnc-request-id,omitempty" require:"true"`
  // {"en":"Accelerated domain name ID", "zh_CN":"加速域名ID"}
  DomainId *string `json:"domain-id,omitempty" xml:"domain-id,omitempty" require:"true"`
  // {"en":"Name of accelerated domain name", "zh_CN":"加速域名的名称"}
  DomainName *string `json:"domain-name,omitempty" xml:"domain-name,omitempty" require:"true"`
  // {"en":"", "zh_CN":"时间戳防盗链设置
  // 注意：
  // 1、时间戳防盗链分为两部分，一部分是防盗链校验，一部分是时间有效性校验。二者都有效，则防盗链通过，否则不通过。
  // 2、防盗链校验：加密算法为md5sum，按照参与MD5计算的参数及组合顺序进行防盗链加密串的计算，对匹配目录下所有文件的url进行防盗链校验，未匹配到的url，则拒绝访问。
  // 3、时间有效性检验：按照年月日时分秒换算的当前时间，与请求url中所带的名文时间相减，判断是否超过设置的上下限（即前后60s内），时间差小于设置上下限的，系统才会给予正常的响应，否则拒绝请求，返回403
  // 4、日志记录没有带加密串的url
  // 6、需要清空时间戳防盗链规则时，可以只传入节点<timestamp-visit-control-rule></timestamp-visit-control-rule>"}
  TimestampVisitControlRule *QuerytimecontrolServiceResponseTimestampVisitControlRule `json:"timestamp-visit-control-rule,omitempty" xml:"timestamp-visit-control-rule,omitempty" require:"true" type:"Struct"`
}

func (s QuerytimecontrolServiceResponse) String() string {
  return tea.Prettify(s)
}

func (s QuerytimecontrolServiceResponse) GoString() string {
  return s.String()
}

func (s *QuerytimecontrolServiceResponse) SetHttpStatus(v int) *QuerytimecontrolServiceResponse {
  s.HttpStatus = &v
  return s
}

func (s *QuerytimecontrolServiceResponse) SetXCncRequestId(v string) *QuerytimecontrolServiceResponse {
  s.XCncRequestId = &v
  return s
}

func (s *QuerytimecontrolServiceResponse) SetDomainId(v string) *QuerytimecontrolServiceResponse {
  s.DomainId = &v
  return s
}

func (s *QuerytimecontrolServiceResponse) SetDomainName(v string) *QuerytimecontrolServiceResponse {
  s.DomainName = &v
  return s
}

func (s *QuerytimecontrolServiceResponse) SetTimestampVisitControlRule(v *QuerytimecontrolServiceResponseTimestampVisitControlRule) *QuerytimecontrolServiceResponse {
  s.TimestampVisitControlRule = v
  return s
}

type QuerytimecontrolServiceResponseTimestampVisitControlRule struct {
  // {"en":"", "zh_CN":"同缓存规则设置中的&ldquo;path-pattern&rdquo;，用于URL匹配。对于匹配到的URL进行时间戳防盗链验证；未匹配到的URL，则拒绝。
  // 同缓存规则设置中的&ldquo;path-pattern&rdquo;，用于URL匹配。对于匹配到的URL进行时间戳防盗链验证；未匹配到的URL，则拒绝。"}
  PathPattern *string `json:"path-pattern,omitempty" xml:"path-pattern,omitempty" require:"true"`
  // {"en":"Optional values are: http, https, http;https, noprefix, empty. If it is empty, it defaults to "http;https"; if it is noprefix, it means that the protocol prefix of url is not specified, and it only matches according to the regularity of path-pattern. This configuration item only matches with path-pattern. example: 1. Specify protocol-of-path-pattern=https and path-pattern=.* to match all https requests, but not http requests. 2. Specify protocol-of-path-pattern=http;https, path-pattern=.* to match all http and https requests. 3. Specify protocol-of-path-pattern=noprefix and path-pattern=^http://[^/]+/.* to match all http requests but not https requests.",
  // "zh_CN":"可选值为: http、https、http;https、noprefix、空。为空默认为“http;https”；为noprefix表示不指定url的协议前缀，仅按path-pattern的正则匹配。本配置项只与path-pattern（url匹配模式）结合匹配。 例子： 1、指定protocol-of-path-pattern=https，path-pattern=.*，则匹配所有https的请求，不匹配http的请求。 2、指定protocol-of-path-pattern=http;https，path-pattern=.*，则匹配所有http和https的请求。 3、指定protocol-of-path-pattern=noprefix，path-pattern=^http://[^/]+/.*，则匹配所有http的请求，不匹配https的请求。"}
  ProtocolOfPathPattern *string `json:"protocol-of-path-pattern,omitempty" xml:"protocol-of-path-pattern,omitempty"`
  // {"en":"Directory, multiple separated by English semicolons. Perform timestamp anti-leech verification for the matched directory; reject the unmatched directory. mutually exclusive with path-pattern.", "zh_CN":"目录，多个以英文分号隔开。对于匹配到的目录进行时间戳防盗链验证；未匹配到的则拒绝。和path-pattern互斥。"}
  Directory *string `json:"directory,omitempty" xml:"directory,omitempty"`
  // {"en":"Whether to remove / from $uri in anti-leech, the optional values are true or false, and the default is false, that is, it contains /. For example: http://www.test.com/1.flv, then $uri is /1.flv by default, if ignore-uri-slash is true, then $uri is 1.flv", "zh_CN":"防盗链中的$uri是否去掉/，可选值为true、false，默认为false，即包含/。 例如： http://www.test.com/1.flv，则$uri默认为/1.flv,若ignore-uri-slash为true，则$uri为1.flv"}
  IgnoreUriSlash *string `json:"ignore-uri-slash,omitempty" xml:"ignore-uri-slash,omitempty"`
  // {"en":"Whether key and time are allowed to be interchanged, the optional values are true and false, true is allowed, false is not allowed. By default, the order of key parameters and time parameters must strictly follow the order required by the authentication mode, that is, the default key and time cannot be interchanged. If "true" is selected, the key and time positions can be interchanged and the authentication succeeds.", "zh_CN":"key与time是否允许互换，可选值为true和false，true则允许，false则不允许。默认情况下，密钥参数和时间参数的顺序，必须严格参照鉴权模式要求的顺序，即，默认key和time不能互换位置。如选择“true”，key和time位置可以互换并鉴权成功。"}
  IgnoreKeyAndTimePosition *string `json:"ignore-key-and-time-position,omitempty" xml:"ignore-key-and-time-position,omitempty"`
  // {"en":"", "zh_CN":"例外的url匹配模式，某些URL除外：如abc.jpg，不做防盗链
  // 例外的url匹配模式，某些URL除外：如abc.jpg，不做防盗链"}
  ExceptPathPattern *string `json:"except-path-pattern,omitempty" xml:"except-path-pattern,omitempty" require:"true"`
  // {"en":"", "zh_CN":"例外的IP，支持输入IP或IP段，IP段之间用分号(;)隔开，如1.1.1.0/24;2.2.2.2，某些IP例外，不做防盗链"}
  AllowedIps *string `json:"allowed-ips,omitempty" xml:"allowed-ips,omitempty" require:"true"`
  // {"en":"", "zh_CN":"加密算法
  // 当前支持入参：md5sum"}
  EncryptMethod *string `json:"encrypt-method,omitempty" xml:"encrypt-method,omitempty" require:"true"`
  // {"en":"", "zh_CN":"防盗链加密串，支持多个加密串，多个加密串以分号(;)隔开
  // 示例：<multiple-secret-keys>abcdef;uvwxyz</multiple-secret-keys>
  // 注意：
  // 1、支持对url 防盗链设置多个密钥。支持客户任意修改密钥，并做到无缝切换。防盗链等级更高。
  // 2、请求url的key只要跟其中任意一个加密串算出来的key一致就验证通过"}
  MultipleSecretKeys *string `json:"multiple-secret-keys,omitempty" xml:"multiple-secret-keys,omitempty" require:"true"`
  // {"en":"", "zh_CN":"防盗链加密串时间格式，可多选，以分号(;)分隔
  // 年|月|日|时|分|秒|UNIX时间戳|16进制时间戳：1Y;2m;3d;4H;5M;6S;7s;8x
  // 示例：<time-format>1Y;2m;3d;4H;5M;6S;7s;8x</time-format> 或 <time-format>1Y;2m;3d;4H;5M;7s</time-format>
  // 注意：
  // 1、必须是英文字母和数字
  // 2、每个值只能出现1次"}
  TimeFormat *string `json:"time-format,omitempty" xml:"time-format,omitempty" require:"true"`
  // {"en":"", "zh_CN":"防盗链回源方式，可选值：1（使用未加密url回源）、2（使用客户请求带加密串url回源）
  // 示例：<dst-style>1<dst-style>
  // 注意：
  // 1、如果URL格式是：http://www.xxx.com/md5/time/uri?参数，则需要下工单给对应客服，让客服在父配置去掉时间戳格式再缓存。"}
  DstStyle *string `json:"dst-style,omitempty" xml:"dst-style,omitempty" require:"true"`
  // {"en":"", "zh_CN":"&nbsp;日志记录原始url，可选值：true（日志记录原始url）、false（不开启日志记录原始url）"}
  LogFormat *string `json:"log-format,omitempty" xml:"log-format,omitempty" require:"true"`
  // {"en":"", "zh_CN":"该配置项用于通过匹配的才进行m3u8改写，可选值：false（不进行m3u8改写）、true(进行m3u8改写）
  // 示例：<m3u8>true</m3u8>"}
  M3u8 *string `json:"m3u8,omitempty" xml:"m3u8,omitempty" require:"true"`
  // {"en":"", "zh_CN":"用于配置获取url中的key的名称
  // 示例： <url-key>auth_key</url-key>"}
  UrlKey *string `json:"url-key,omitempty" xml:"url-key,omitempty" require:"true"`
  // {"en":"", "zh_CN":"时间戳防盗链多条匹配模式
  // 注意：多条配置项如下：
  // 1、url匹配模式的协议：
  // 2、url匹配模式
  // 3、例外url匹配模式的协议
  // 4、例外url匹配模式
  // 5、防盗链生成方式
  // 6、请求url格式
  // 7、通用防盗链取uri对应第几个&ldquo;/&rdquo;
  // 8、清空多条配置只传入节点：<timestamp-control-rules></timestamp-control-rules>"}
  TimestampControlRules []*QuerytimecontrolServiceResponseTimestampVisitControlRuleTimestampControlRules `json:"timestamp-control-rules,omitempty" xml:"timestamp-control-rules,omitempty" require:"true" type:"Repeated"`
}

func (s QuerytimecontrolServiceResponseTimestampVisitControlRule) String() string {
  return tea.Prettify(s)
}

func (s QuerytimecontrolServiceResponseTimestampVisitControlRule) GoString() string {
  return s.String()
}

func (s *QuerytimecontrolServiceResponseTimestampVisitControlRule) SetPathPattern(v string) *QuerytimecontrolServiceResponseTimestampVisitControlRule {
  s.PathPattern = &v
  return s
}

func (s *QuerytimecontrolServiceResponseTimestampVisitControlRule) SetProtocolOfPathPattern(v string) *QuerytimecontrolServiceResponseTimestampVisitControlRule {
  s.ProtocolOfPathPattern = &v
  return s
}

func (s *QuerytimecontrolServiceResponseTimestampVisitControlRule) SetDirectory(v string) *QuerytimecontrolServiceResponseTimestampVisitControlRule {
  s.Directory = &v
  return s
}

func (s *QuerytimecontrolServiceResponseTimestampVisitControlRule) SetIgnoreUriSlash(v string) *QuerytimecontrolServiceResponseTimestampVisitControlRule {
  s.IgnoreUriSlash = &v
  return s
}

func (s *QuerytimecontrolServiceResponseTimestampVisitControlRule) SetIgnoreKeyAndTimePosition(v string) *QuerytimecontrolServiceResponseTimestampVisitControlRule {
  s.IgnoreKeyAndTimePosition = &v
  return s
}

func (s *QuerytimecontrolServiceResponseTimestampVisitControlRule) SetExceptPathPattern(v string) *QuerytimecontrolServiceResponseTimestampVisitControlRule {
  s.ExceptPathPattern = &v
  return s
}

func (s *QuerytimecontrolServiceResponseTimestampVisitControlRule) SetAllowedIps(v string) *QuerytimecontrolServiceResponseTimestampVisitControlRule {
  s.AllowedIps = &v
  return s
}

func (s *QuerytimecontrolServiceResponseTimestampVisitControlRule) SetEncryptMethod(v string) *QuerytimecontrolServiceResponseTimestampVisitControlRule {
  s.EncryptMethod = &v
  return s
}

func (s *QuerytimecontrolServiceResponseTimestampVisitControlRule) SetMultipleSecretKeys(v string) *QuerytimecontrolServiceResponseTimestampVisitControlRule {
  s.MultipleSecretKeys = &v
  return s
}

func (s *QuerytimecontrolServiceResponseTimestampVisitControlRule) SetTimeFormat(v string) *QuerytimecontrolServiceResponseTimestampVisitControlRule {
  s.TimeFormat = &v
  return s
}

func (s *QuerytimecontrolServiceResponseTimestampVisitControlRule) SetDstStyle(v string) *QuerytimecontrolServiceResponseTimestampVisitControlRule {
  s.DstStyle = &v
  return s
}

func (s *QuerytimecontrolServiceResponseTimestampVisitControlRule) SetLogFormat(v string) *QuerytimecontrolServiceResponseTimestampVisitControlRule {
  s.LogFormat = &v
  return s
}

func (s *QuerytimecontrolServiceResponseTimestampVisitControlRule) SetM3u8(v string) *QuerytimecontrolServiceResponseTimestampVisitControlRule {
  s.M3u8 = &v
  return s
}

func (s *QuerytimecontrolServiceResponseTimestampVisitControlRule) SetUrlKey(v string) *QuerytimecontrolServiceResponseTimestampVisitControlRule {
  s.UrlKey = &v
  return s
}

func (s *QuerytimecontrolServiceResponseTimestampVisitControlRule) SetTimestampControlRules(v []*QuerytimecontrolServiceResponseTimestampVisitControlRuleTimestampControlRules) *QuerytimecontrolServiceResponseTimestampVisitControlRule {
  s.TimestampControlRules = v
  return s
}

type QuerytimecontrolServiceResponseTimestampVisitControlRuleTimestampControlRules struct     {
  // {"en":"", "zh_CN":"添加grid类型标识，表示客户多组配置时，具体某组配置
  // 注意：添加grid类型标识：data-id，每一组配置对应一个data-id：a、如果客户有传data-id，说明指定修改其中一组配置项内容，不需求修改其他组配置内容不需要入参；b、如果客户入参多组配置，其中有些组配置有传data-id，有些没有传，则有传data-id的表示修改具体某组配置，没有传data-id的表示在原来基础上新增一组配置；c、如果客户入参都没有传data-id,表示用本次的配置全量覆盖原先配置；d、如果客户入参没有传任何配置项参数，只传了域名和二级标签，表示清空这个接口对应域名二级服务所有配置。（c、d内容和当前方案实现一致）；e、一个gird标签下的入参不能为空，如果，没有具体的配置项，则data-id必填，且值为实际存在的data-id,表示清空这个data-id对应配置项的值；"}
  DataId *int64 `json:"data-id,omitempty" xml:"data-id,omitempty" require:"true"`
  PathPattern *string `json:"path-pattern,omitempty" xml:"path-pattern,omitempty" require:"true"`
  ExceptPathPattern *string `json:"except-path-pattern,omitempty" xml:"except-path-pattern,omitempty" require:"true"`
  // {"en":"", "zh_CN":"协议，可选值：http|https|http;https|noprefix
  // http:http协议
  // https:http协议
  // http;https:http和https协议
  // noprefix:不加前缀，按path-pattern设置的规则
  // 用法示例：
  // path-pattern：输入.*\jpg$     
  // path-pattern-protocol：选择http，则表示^http://[^/]+/.*\jpg$  ；
  // 选择https，则表示 ^https://[^/]+/ .*\jpg$；
  // 选择http;https，则表示^https?://[^/]+/.*\jpg$；
  // 选择noprefix，则表示.*\jpg$
  // 注意：
  // 1、为空时，则默认&ldquo;http和https协议&rdquo;
  // 2、只与url匹配模式匹配（path-pattern）
  // 3、协议必须跟模式一起，但是模式可以单独存在"}
  PathPatternProtocol *string `json:"path-pattern-protocol,omitempty" xml:"path-pattern-protocol,omitempty" require:"true"`
  // {"en":"", "zh_CN":"协议，可选值：http|https|http;https|noprefix
  // http:http协议
  // https:http协议
  // http;https:http和https协议
  // noprefix:不加前缀
  // 用法示例：
  // except-path-pattern：输入.*\jpg$     
  // except-path-pattern-protocol：选择http，则表示^http://[^/]+/.*\jpg$  ；
  // 选择https，则表示 ^https://[^/]+/ .*\jpg$；
  // 选择http;https，则表示^https?://[^/]+/.*\jpg$；
  // 选择noprefix，则表示.*\jpg$
  // 注意：
  // 1、为空时，则默认&ldquo;http和https协议&rdquo;
  // 2、只与例外url匹配模式匹配（except-path-pattern）
  // 3、协议必须跟模式一起，但是模式可以单独存在"}
  ExceptPathPatternProtocol *string `json:"except-path-pattern-protocol,omitempty" xml:"except-path-pattern-protocol,omitempty" require:"true"`
  // {"en":"", "zh_CN":"防盗链生成方式，参与MD5计算的参数及组合顺序，仅支持传入以下参数：
  // $uri：介于domain和问号之间的字符串，特殊取值在入参<uri-select>中配置
  // 例如http://cdn.example.com/v0/test.dat?k=v，则URI为/v0/test.dat
  // $ourkey：秘钥，实际秘钥在入参<secret-key>中配置
  // $time：时间串
  // $spec_name：文件名
  // 例如：http://cdn.example.com/v0/test.dat?k=v，文件名为 test.dat
  // $args：QUERY_STRING中的某个具体key的值
  // 示例：<cipher-combination>$uri$ourkey$time$args{k}</cipher-combination>
  // 注意：
  // 1、$args{k}中的K只允许A-Z大小写字母、数字、下划线、横杠
  // 2、除了$args外，其它参数只允许出现1次
  // 3、可以按照任意组合顺序，拼接各个参数"}
  CipherCombination *string `json:"cipher-combination,omitempty" xml:"cipher-combination,omitempty" require:"true"`
  // {"en":"", "zh_CN":"防盗链加密串的秘钥，只允许传入一个秘钥
  // 示例：<secret-key>abcdef</secret-key>
  // 注意：
  // 1、与客户约定好的秘钥，入参<multiple-secret-keys>中的某个值可以等于<secret-key>对应的值
  // 2、$ourkey值主要是来自<secret-key>的配置"}
  SecretKey *string `json:"secret-key,omitempty" xml:"secret-key,omitempty" require:"true"`
  // {"en":"", "zh_CN":"防盗链串的参数名称
  // 示例：<cipher-param>keyname</cipher-param>
  // 注意：
  // 1、如果防盗链加密串是在url问号后的参数中，则防盗链加密串的参数名由<cipher-param>的配置决定；
  // 2、如果<cipher-param>为空，则默认请求的url中使用key作为参数名"}
  CipherParam *string `json:"cipher-param,omitempty" xml:"cipher-param,omitempty" require:"true"`
  // {"en":"", "zh_CN":"时间串的参数名称
  // 示例：<time-param>tname</time-param>
  // 注意：
  // 1、如果防盗链时间串是放在url问号后面的参数中，则防盗链时间串的参数名由<time-param>的配置决定；
  // 2、如果<time-param>为空，则默认请求的url中使用time作为参数名"}
  TimeParam *string `json:"time-param,omitempty" xml:"time-param,omitempty" require:"true"`
  // {"en":"", "zh_CN":"防盗链串的过期时间下限
  // 示例：
  // <lower-limit-expiry-time>200</lower-limit-expiry-time>
  // 对应各种场景的配置方式如下：
  // 注意：
  // 1、请求URL中携带的时间戳如果是URL的生成时间，需要加上有效时长才是过期时间，即<lower-limit-expiry-time>和<upper-limit-expiry-time>配置为有效时长。
  // 2、如果URL携带的时间戳是过期时间，则可以配成零。"}
  LowerLimitExpiryTime *string `json:"lower-limit-expiry-time,omitempty" xml:"lower-limit-expiry-time,omitempty" require:"true"`
  // {"en":"", "zh_CN":"防盗链串的过期时间上限
  // 示例：
  // <upper-limit-expiry-time>5000</upper-limit-expiry-time>
  // 对应各种场景的配置方式如下：
  // 注意：
  // 1、请求URL中携带的时间戳如果是URL的生成时间，需要加上有效时长才是过期时间，即<lower-limit-expiry-time>和<upper-limit-expiry-time>配置为有效时长。
  // 2、如果URL携带的时间戳是过期时间，则可以配成零。"}
  UpperLimitExpiryTime *string `json:"upper-limit-expiry-time,omitempty" xml:"upper-limit-expiry-time,omitempty" require:"true"`
  // {"en":"", "zh_CN":"防盗链请求url格式，支持两种防盗链方式，即加密串和时间戳放到&ldquo;?&rdquo;后面或者是加密串和时间戳放到&ldquo;host&rdquo;后面，url格式支持的参数如下：
  // $domain：域名
  // $uri：不包含域名的url部分
  // $key：防盗链加密串的MD5值
  // $time：防盗链时间串
  // $args：问号后的QUERY_STRING参数
  // 示例：支持以下请求url格式，可替换为https://，url请求协议根据实际使用，如不知道如何正确配置，请找客户技术支持协助；携带加密串和时间串两个值的参数名&ldquo;keyname&rdquo;和&ldquo;tname&rdquo;，可替换为实际使用的参数名
  // <request-url-style>http://$domain/$key/$time/$uri?$args</request-url-style>
  // <request-url-style>http://$domain/$time/$key/$uri?$args</request-url-style>
  //  http://$domain/$uri?auth_key=$key 
  // <request-url-style>http://$domain/$uri?keyname=$key&tname=$time</request-url-style>
  // <request-url-style>http://$domain/$uri?$args&keyname=$key&tname=$time</request-url-style>
  // <request-url-style>http://$domain/$uri?keyname=$key&tname=$time&$args</request-url-style>
  // <request-url-style>http://$domain/$uri?$args&keyname=$key&tname=$time&$args</request-url-style>
  // <request-url-style>http://$domain/$uri?tname=$time&keyname=$key</request-url-style>
  // <request-url-style>http://$domain/$uri?$args&tname=$time&keyname=$key</request-url-style>
  // <request-url-style>http://$domain/$uri?tname=$time&keyname=$key&$args</request-url-style>
  // <request-url-style>http://$domain/$uri?$args&tname=$time&keyname=$key&$args</request-url-style>
  // 注意：
  // 1、输入的url必须以&ldquo;http/https&rdquo;开头
  // 2、如果加密串和时间戳是放到&ldquo;?&rdquo;后面时，keyname和tname必须跟<cipher-param>和<time-param>配置的值一致
  // 3、如果<cipher-param>和<time-param>没有配置值，则$key对应的参数名默认为key，$time对应的参数名默认为time
  // 4、如果防盗链加密串和时间串在url问号后面的参数中，url中的&ldquo;keyname&rdquo;和&ldquo;tname&rdquo;，对应的是cipher-param和 time-param配置的防盗链串和时间串参数名称。"}
  RequestUrlStyle *string `json:"request-url-style,omitempty" xml:"request-url-style,omitempty" require:"true"`
  // {"en":"", "zh_CN":"通用防盗链取uri对应第几个&ldquo;/&rdquo;，值为数字，多个值，以分号隔开。
  // 配置通用防盗链取uri对应第几个&ldquo;/&rdquo;,从配置为0则为正向第一个，-1为逆向第一个，可分频道配置，默认空值。
  // 比如http://a.com/b/c/d/e/f/1.html
  // 那么0;3; 4;-1 就是取/b/e/f/1.html
  // 注意：
  // 1、取值范围：[-100,+100]
  // 2、重复配置生效规则（建议不要配重复）
  // 重复的将会跳过解析：
  // 如http://cdn.example.com/e/0/e2/test.dat?k=v
  // 可以配<uri-select> 2;2;3;3</uri-select>，但只取2;3
  // URI则为/e2/test.dat；"}
  UriSelect *string `json:"uri-select,omitempty" xml:"uri-select,omitempty" require:"true"`
}

func (s QuerytimecontrolServiceResponseTimestampVisitControlRuleTimestampControlRules) String() string {
  return tea.Prettify(s)
}

func (s QuerytimecontrolServiceResponseTimestampVisitControlRuleTimestampControlRules) GoString() string {
  return s.String()
}

func (s *QuerytimecontrolServiceResponseTimestampVisitControlRuleTimestampControlRules) SetDataId(v int64) *QuerytimecontrolServiceResponseTimestampVisitControlRuleTimestampControlRules {
  s.DataId = &v
  return s
}

func (s *QuerytimecontrolServiceResponseTimestampVisitControlRuleTimestampControlRules) SetPathPattern(v string) *QuerytimecontrolServiceResponseTimestampVisitControlRuleTimestampControlRules {
  s.PathPattern = &v
  return s
}

func (s *QuerytimecontrolServiceResponseTimestampVisitControlRuleTimestampControlRules) SetExceptPathPattern(v string) *QuerytimecontrolServiceResponseTimestampVisitControlRuleTimestampControlRules {
  s.ExceptPathPattern = &v
  return s
}

func (s *QuerytimecontrolServiceResponseTimestampVisitControlRuleTimestampControlRules) SetPathPatternProtocol(v string) *QuerytimecontrolServiceResponseTimestampVisitControlRuleTimestampControlRules {
  s.PathPatternProtocol = &v
  return s
}

func (s *QuerytimecontrolServiceResponseTimestampVisitControlRuleTimestampControlRules) SetExceptPathPatternProtocol(v string) *QuerytimecontrolServiceResponseTimestampVisitControlRuleTimestampControlRules {
  s.ExceptPathPatternProtocol = &v
  return s
}

func (s *QuerytimecontrolServiceResponseTimestampVisitControlRuleTimestampControlRules) SetCipherCombination(v string) *QuerytimecontrolServiceResponseTimestampVisitControlRuleTimestampControlRules {
  s.CipherCombination = &v
  return s
}

func (s *QuerytimecontrolServiceResponseTimestampVisitControlRuleTimestampControlRules) SetSecretKey(v string) *QuerytimecontrolServiceResponseTimestampVisitControlRuleTimestampControlRules {
  s.SecretKey = &v
  return s
}

func (s *QuerytimecontrolServiceResponseTimestampVisitControlRuleTimestampControlRules) SetCipherParam(v string) *QuerytimecontrolServiceResponseTimestampVisitControlRuleTimestampControlRules {
  s.CipherParam = &v
  return s
}

func (s *QuerytimecontrolServiceResponseTimestampVisitControlRuleTimestampControlRules) SetTimeParam(v string) *QuerytimecontrolServiceResponseTimestampVisitControlRuleTimestampControlRules {
  s.TimeParam = &v
  return s
}

func (s *QuerytimecontrolServiceResponseTimestampVisitControlRuleTimestampControlRules) SetLowerLimitExpiryTime(v string) *QuerytimecontrolServiceResponseTimestampVisitControlRuleTimestampControlRules {
  s.LowerLimitExpiryTime = &v
  return s
}

func (s *QuerytimecontrolServiceResponseTimestampVisitControlRuleTimestampControlRules) SetUpperLimitExpiryTime(v string) *QuerytimecontrolServiceResponseTimestampVisitControlRuleTimestampControlRules {
  s.UpperLimitExpiryTime = &v
  return s
}

func (s *QuerytimecontrolServiceResponseTimestampVisitControlRuleTimestampControlRules) SetRequestUrlStyle(v string) *QuerytimecontrolServiceResponseTimestampVisitControlRuleTimestampControlRules {
  s.RequestUrlStyle = &v
  return s
}

func (s *QuerytimecontrolServiceResponseTimestampVisitControlRuleTimestampControlRules) SetUriSelect(v string) *QuerytimecontrolServiceResponseTimestampVisitControlRuleTimestampControlRules {
  s.UriSelect = &v
  return s
}

type QuerytimecontrolServicePaths struct {
  // {"en":"", "zh_CN":"域名名称或域名id，在请求的url后面"}
  DomainName *string `json:"domain-name,omitempty" xml:"domain-name,omitempty" require:"true"`
}

func (s QuerytimecontrolServicePaths) String() string {
  return tea.Prettify(s)
}

func (s QuerytimecontrolServicePaths) GoString() string {
  return s.String()
}

func (s *QuerytimecontrolServicePaths) SetDomainName(v string) *QuerytimecontrolServicePaths {
  s.DomainName = &v
  return s
}

type QuerytimecontrolServiceParameters struct {
}

func (s QuerytimecontrolServiceParameters) String() string {
  return tea.Prettify(s)
}

func (s QuerytimecontrolServiceParameters) GoString() string {
  return s.String()
}

type QuerytimecontrolServiceRequestHeader struct {
}

func (s QuerytimecontrolServiceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QuerytimecontrolServiceRequestHeader) GoString() string {
  return s.String()
}

type QuerytimecontrolServiceResponseHeader struct {
}

func (s QuerytimecontrolServiceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QuerytimecontrolServiceResponseHeader) GoString() string {
  return s.String()
}




type EditCacheTimeConfigRequest struct {
  // {"en":"Cache time configuration
  // note:
  // 1. When you need to cancel the cache time configuration setting, you can pass in the empty node <cache-time-behaviors></cache-time-behaviors>.
  // 2. When it is required to set the cache time configuration, this item is required.", "zh_CN":"缓存时间配置
  // 注意：
  // 1. 需要取消缓存时间配置设置时，可以传入空节点<cache-time-behaviors></cache-time-behaviors>。
  // 2. 表示需要设置缓存时间配置时，此项必填"}
  CacheTimeBehaviors []*EditCacheTimeConfigRequestCacheTimeBehaviors `json:"cache-time-behaviors,omitempty" xml:"cache-time-behaviors,omitempty" require:"true" type:"Repeated"`
}

func (s EditCacheTimeConfigRequest) String() string {
  return tea.Prettify(s)
}

func (s EditCacheTimeConfigRequest) GoString() string {
  return s.String()
}

func (s *EditCacheTimeConfigRequest) SetCacheTimeBehaviors(v []*EditCacheTimeConfigRequestCacheTimeBehaviors) *EditCacheTimeConfigRequest {
  s.CacheTimeBehaviors = v
  return s
}

type EditCacheTimeConfigRequestCacheTimeBehaviors struct     {
  // {"en":"dataId is to indicate a specific group configuration when the client has multiple groups of configurations. dataId can be retrieved through a query interface.
  // Note:
  // 1. If dataId is passed, it means that one group of configuration items is specified to be modified, and no other group configuration items need to be modified.
  // 2. If multiple groups of configurations are included, some of them are configured with dataId and others are not, then the expression of dataId is used to modify a specific group of configurations, and a new group of configurations is added on the original basis without the expression of dataId.
  // 3. If the dataId is not passed, it means that the original configuration will be fully covered by this configuration.
  // 4. If no configuration parameter is passed, only domain name and secondary label are passed. It means that all configuration corresponding to this interface is cleared.
  // 5. If there is no specific configuration item in a set of configurations, the dataId must be filled in, and the value should be the actual dataId, which means clearing the value of the corresponding dataId configuration item.
  // 6. Tt is not allowed when neither configuration item nor dataId is specified in a set of configurations.", 
  //     "zh_CN":"中文：配置多组配置时，具体某组配置的id。dataId可以通过查询接口获取。
  // 注意：
  // 1、如果有传dataId，说明指定修改其中一组配置项内容，不需求修改其他组配置内容不需要入参；
  // 2、如果入参多组配置，其中有些组配置有传dataId，有些没有传，则有传dataId的表示修改具体某组配置，没有传dataId的表示在原来基础上新增一组配置；
  // 3、如果入参都没有传dataId,表示用本次的配置全量覆盖原先配置；
  // 4、如果入参没有传任何配置项参数，只传了域名和二级标签，表示清空这个该接口对应的域名配置；
  // 5、如果一组配置没有具体的配置项，则dataId必填，且值要为实际存在的dataId，此时表示清空这个dataId对应的一组配置；
  // 6、不允许一组配置没有指定具体的配置项也没有dataId。"}
  DataId *int64 `json:"data-id,omitempty" xml:"data-id,omitempty"`
  // {"en":"The url matching mode supports fuzzy regularization. If all matches, the input parameters can be configured as: *", "zh_CN":"url匹配模式，支持正则，如果是全部匹配，入参可以配置为：*"}
  PathPattern *string `json:"path-pattern,omitempty" xml:"path-pattern,omitempty"`
  // {"en":"Exceptional url matching mode, except for some URLs: such as abc.jpg, do not do anti-theft chain function
  // E.g: ^https?://[^/]+/.*\.m3u8", "zh_CN":"例外的url匹配模式，某些URL除外：如abc.jpg，不做内容重定向
  // 客户入参参考：^https?://[^/]+/.*\.m3u8"}
  ExceptPathPattern *string `json:"except-path-pattern,omitempty" xml:"except-path-pattern,omitempty"`
  // {"en":"Specify common types: Select the domain name that requires the cache  to be all files or the home page. :
  // E.g:
  // All: all files
  // Homepage: homepage", "zh_CN":"指定常用类型：选择缓存域名的是全部文件还是首页。入参参考值：
  // all：全部文件
  // homepage：首页"}
  CustomPattern *string `json:"custom-pattern,omitempty" xml:"custom-pattern,omitempty"`
  // {"en":"File Type: Specify the file type for cache settings.
  // File types include: gif png bmp jpeg jpg html htm shtml mp3 wma flv mp4 wmv zip exe rar css txt ico js swf
  // If you need all types, pass all directly. Multiples are separated by semicolons, and all and specific file types cannot be configured at the same time.", "zh_CN":"文件类型：指定需要缓存的文件类型。
  // 文件类型包括：gif png bmp jpeg jpg html htm shtml mp3 wma flv mp4 wmv zip exe rar css txt ico js swf
  // 如果需要全部类型，则直接传all。多个以分号隔开，all和具体文件类型不能同时配置。"}
  FileType *string `json:"file-type,omitempty" xml:"file-type,omitempty"`
  // {"en":"Custom file type: Fill in the appropriate identifiable file type according to your needs outside of the specified file type. Can be used with file-type. If the file-type is also configured, the actual file type is the sum of the two parameters.", "zh_CN":"自定义文件类型：在指定文件类型外根据自身需求，填写适当的可识别文件类型。可以搭配file-type使用。如果file-type也有配置，实际生效的文件类型是两个入参的总和"}
  CustomFileType *string `json:"custom-file-type,omitempty" xml:"custom-file-type,omitempty"`
  // {"en":"Specify URL cache: Specify url according to requirements for cache
  // INS format does not support URI format with http(s)://", "zh_CN":"指定URL缓存：根据需求指定url进行缓存
  // 入参不支持含http(s):// 开头的URI格式"}
  SpecifyUrlPattern *string `json:"specify-url-pattern,omitempty" xml:"specify-url-pattern,omitempty"`
  // {"en":"Directory: Specify the directory cache.
  // Enter a legal directory format. Multiple separated by semicolons", "zh_CN":"目录：指定目录缓存。
  // 输入合法的目录格式。多个以英文分号隔开"}
  Directory *string `json:"directory,omitempty" xml:"directory,omitempty"`
  // {"en":"Cache time: set the time corresponding to the cache object
  // Input format: integer plus unit, such as 20s, 30m, 1h, 2d, no cache is set to 0. Do not enter the unit default is seconds
  // There is no upper limit on the cache time theory. This time is set according to the customer's own needs. If the customer feels that some of the files are not changed frequently, then the setting is longer. For example, the text class js, css, html, etc. can be set shorter, the picture, video and audio classes can be set longer (because the cache time will be replaced by the new file due to the file heat algorithm, the longest suggestion Do not exceed one month)", "zh_CN":"缓存时间：设置缓存对象对应的时间
  // 入参格式：整数加单位，比如20s、30m、1h、2d，不缓存设置为0。不输入单位默认是秒
  // 缓存时间理论上没有上限限制，这个时间根据客户自身的需求设定，如果客户觉得其中一些文件，变更不频繁，那么就设置长一点。例如，文本类的js，css，html等可以设置得短一些，图片、视频音频类的可以设置的长一点（因为缓存时间会因文件热度算法，旧文件会被新文件替换掉，最长建议不要超过一个月）"}
  CacheTtl *string `json:"cache-ttl,omitempty" xml:"cache-ttl,omitempty"`
  // {"en":"Ignore the source station does not cache the header. The optional values are true and false, which are used to ignore the two configurations of cache-control in the request header (private, no-cache) and the Authorization set by the client.
  // The ture indicates that the source station's settings for the three are ignored. Enables resources to be cached on the service node in the form of cache-control: public, and then our nodes can cache this type of resource and provide acceleration services.
  // False means that when the source station sets cache-control: private, cache-control: no-cache for a resource or specifies to cache according to authorization, our service node will not cache such files.", "zh_CN":"忽略源站不缓存头。可选值为true和false，用于忽略请求头中cache-control的两种配置（private，no-cache）和客户端设置的Authorization。
  // ture表示会忽略掉源站对于这三者的设定。使得资源能够以cache-control: public的方式缓存在服务节点上，然后我们的节点才能缓存这种类型的资源，提供加速服务。
  // false表示当源站对某种资源设定了cache-control: private,cache-control:no-cache或指定根据authorization进行缓存时，我们的服务节点将不会对此类文件进行缓存。"}
  IgnoreCacheControl *string `json:"ignore-cache-control,omitempty" xml:"ignore-cache-control,omitempty"`
  // {"en":"Respect the server: Accelerate whether to prioritize the source cache time.
  // Optional values: true and false
  // True: indicates that the server is time-first
  // False: The cache time of the CDN configuration takes precedence.", "zh_CN":"尊重服务端：加速是否要按源站缓存时间优先。
  // 可选值：true和false
  // true：表示重服务端时间优先
  // false:CDN配置的缓存时间优先"}
  IsRespectServer *string `json:"is-respect-server,omitempty" xml:"is-respect-server,omitempty"`
  // {"en":"Ignore case, the optional value is true or false, true means to ignore case; false means not to ignore case;
  // When adding a new configuration item, the default is not true.", "zh_CN":"忽略大小写，可选值为true或false，true表示忽略大小写；false表示不忽略大小写；
  // 新增配置项时，不传默认为 true"}
  IgnoreLetterCase *string `json:"ignore-letter-case,omitempty" xml:"ignore-letter-case,omitempty"`
  // {"en":"Reload processing rules, optional: ignore or if-modified-since
  // If-modified-since: indicates that you want to convert to if-modified-since
  // Ignore: means to ignore client refresh", "zh_CN":"reload处理规则，可选项：ignore或者if-modified-since
  // if-modified-since：表示要转成if-modified-since
  // ignore:表示忽略客户端刷新"}
  ReloadManage *string `json:"reload-manage,omitempty" xml:"reload-manage,omitempty"`
  // {"en":"Indicates the priority execution order of multiple sets of redirected content by the customer. The higher the number, the higher the priority.
  // When adding a new configuration item, the default is 10", "zh_CN":"表示客户多组重定向内容的优先执行顺序。数字越大，优先级越高。
  // 新增配置项时，不传默认为 10
  // 如果传了值，不能为空"}
  Priority *string `json:"priority,omitempty" xml:"priority,omitempty"`
  // {"en":"Request Header", "zh_CN":"请求头域"}
  EditCacheTimeConfigRequestHeaderField *string `json:"request-header-field,omitempty" xml:"request-header-field,omitempty"`
  // {"en":"Request Header Value", "zh_CN":"请求头的值"}
  ValueEditCacheTimeConfigRequestHeader *string `json:"value-request-header,omitempty" xml:"value-request-header,omitempty"`
  // {"en":"You can set it 'true' to cache
  // ignoring the http header 'Authentication'.  If it is empty, the header is not ignored by default.", "zh_CN":"忽略鉴权头部Authentication，可选值为true和false。默认为不忽略。"}
  IgnoreAuthenticationHeader *string `json:"ignore-authentication-header,omitempty" xml:"ignore-authentication-header,omitempty"`
}

func (s EditCacheTimeConfigRequestCacheTimeBehaviors) String() string {
  return tea.Prettify(s)
}

func (s EditCacheTimeConfigRequestCacheTimeBehaviors) GoString() string {
  return s.String()
}

func (s *EditCacheTimeConfigRequestCacheTimeBehaviors) SetDataId(v int64) *EditCacheTimeConfigRequestCacheTimeBehaviors {
  s.DataId = &v
  return s
}

func (s *EditCacheTimeConfigRequestCacheTimeBehaviors) SetPathPattern(v string) *EditCacheTimeConfigRequestCacheTimeBehaviors {
  s.PathPattern = &v
  return s
}

func (s *EditCacheTimeConfigRequestCacheTimeBehaviors) SetExceptPathPattern(v string) *EditCacheTimeConfigRequestCacheTimeBehaviors {
  s.ExceptPathPattern = &v
  return s
}

func (s *EditCacheTimeConfigRequestCacheTimeBehaviors) SetCustomPattern(v string) *EditCacheTimeConfigRequestCacheTimeBehaviors {
  s.CustomPattern = &v
  return s
}

func (s *EditCacheTimeConfigRequestCacheTimeBehaviors) SetFileType(v string) *EditCacheTimeConfigRequestCacheTimeBehaviors {
  s.FileType = &v
  return s
}

func (s *EditCacheTimeConfigRequestCacheTimeBehaviors) SetCustomFileType(v string) *EditCacheTimeConfigRequestCacheTimeBehaviors {
  s.CustomFileType = &v
  return s
}

func (s *EditCacheTimeConfigRequestCacheTimeBehaviors) SetSpecifyUrlPattern(v string) *EditCacheTimeConfigRequestCacheTimeBehaviors {
  s.SpecifyUrlPattern = &v
  return s
}

func (s *EditCacheTimeConfigRequestCacheTimeBehaviors) SetDirectory(v string) *EditCacheTimeConfigRequestCacheTimeBehaviors {
  s.Directory = &v
  return s
}

func (s *EditCacheTimeConfigRequestCacheTimeBehaviors) SetCacheTtl(v string) *EditCacheTimeConfigRequestCacheTimeBehaviors {
  s.CacheTtl = &v
  return s
}

func (s *EditCacheTimeConfigRequestCacheTimeBehaviors) SetIgnoreCacheControl(v string) *EditCacheTimeConfigRequestCacheTimeBehaviors {
  s.IgnoreCacheControl = &v
  return s
}

func (s *EditCacheTimeConfigRequestCacheTimeBehaviors) SetIsRespectServer(v string) *EditCacheTimeConfigRequestCacheTimeBehaviors {
  s.IsRespectServer = &v
  return s
}

func (s *EditCacheTimeConfigRequestCacheTimeBehaviors) SetIgnoreLetterCase(v string) *EditCacheTimeConfigRequestCacheTimeBehaviors {
  s.IgnoreLetterCase = &v
  return s
}

func (s *EditCacheTimeConfigRequestCacheTimeBehaviors) SetReloadManage(v string) *EditCacheTimeConfigRequestCacheTimeBehaviors {
  s.ReloadManage = &v
  return s
}

func (s *EditCacheTimeConfigRequestCacheTimeBehaviors) SetPriority(v string) *EditCacheTimeConfigRequestCacheTimeBehaviors {
  s.Priority = &v
  return s
}

func (s *EditCacheTimeConfigRequestCacheTimeBehaviors) SetEditCacheTimeConfigRequestHeaderField(v string) *EditCacheTimeConfigRequestCacheTimeBehaviors {
  s.EditCacheTimeConfigRequestHeaderField = &v
  return s
}

func (s *EditCacheTimeConfigRequestCacheTimeBehaviors) SetValueEditCacheTimeConfigRequestHeader(v string) *EditCacheTimeConfigRequestCacheTimeBehaviors {
  s.ValueEditCacheTimeConfigRequestHeader = &v
  return s
}

func (s *EditCacheTimeConfigRequestCacheTimeBehaviors) SetIgnoreAuthenticationHeader(v string) *EditCacheTimeConfigRequestCacheTimeBehaviors {
  s.IgnoreAuthenticationHeader = &v
  return s
}

type EditCacheTimeConfigResponse struct {
  // {"en":"Error code, which appears when HTTPStatus is not 202, represents the error type of the current request call", "zh_CN":"错误代码，当HTTPStatus不为202时出现，表示当前请求调用的错误类型"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response message, and shows as success when it is successful.", "zh_CN":"响应信息，成功时为success"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s EditCacheTimeConfigResponse) String() string {
  return tea.Prettify(s)
}

func (s EditCacheTimeConfigResponse) GoString() string {
  return s.String()
}

func (s *EditCacheTimeConfigResponse) SetCode(v string) *EditCacheTimeConfigResponse {
  s.Code = &v
  return s
}

func (s *EditCacheTimeConfigResponse) SetMessage(v string) *EditCacheTimeConfigResponse {
  s.Message = &v
  return s
}

type EditCacheTimeConfigPaths struct {
  // {"en":"the domain whoes need query config", "zh_CN":"需要查询配置的域名或域名id"}
  DomainName *string `json:"domain-name,omitempty" xml:"domain-name,omitempty" require:"true"`
}

func (s EditCacheTimeConfigPaths) String() string {
  return tea.Prettify(s)
}

func (s EditCacheTimeConfigPaths) GoString() string {
  return s.String()
}

func (s *EditCacheTimeConfigPaths) SetDomainName(v string) *EditCacheTimeConfigPaths {
  s.DomainName = &v
  return s
}

type EditCacheTimeConfigParameters struct {
}

func (s EditCacheTimeConfigParameters) String() string {
  return tea.Prettify(s)
}

func (s EditCacheTimeConfigParameters) GoString() string {
  return s.String()
}

type EditCacheTimeConfigRequestHeader struct {
}

func (s EditCacheTimeConfigRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s EditCacheTimeConfigRequestHeader) GoString() string {
  return s.String()
}

type EditCacheTimeConfigResponseHeader struct {
}

func (s EditCacheTimeConfigResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s EditCacheTimeConfigResponseHeader) GoString() string {
  return s.String()
}




type QueryApiDeployServiceRequest struct {
  // {"en":"or each account request record, a unique cnc-request-id (for all API) is generated through which the execution results of each asynchronous request task can be queried", "zh_CN":"对于账号每一次请求记录，都会生成唯一的cnc-request-id（适用全部接口），通过该id可以查询每次异步请求任务的执行结果"}
  CncCequestId *string `json:"cncCequestId,omitempty" xml:"cncCequestId,omitempty" require:"true"`
}

func (s QueryApiDeployServiceRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryApiDeployServiceRequest) GoString() string {
  return s.String()
}

func (s *QueryApiDeployServiceRequest) SetCncCequestId(v string) *QueryApiDeployServiceRequest {
  s.CncCequestId = &v
  return s
}

type QueryApiDeployServiceResponse struct {
  // {"en":"httpstatus=202; Indicates that the new domain API was successfully invoked, and the current deployment of the new domain can be viewed using x-cnc-request-id in the header", "zh_CN":"httpstatus=202;   表示成功调用新增域名接口，可使用header中的x-cnc-request-id查看当前新增域名的部署情况"}
  HttpStatus *int `json:"http-status-code,omitempty" xml:"http-status-code,omitempty" require:"true"`
  // {"en":"Uniquely identified id for querying tasks per request (for all API)", "zh_CN":"唯一标示的id，用于查询每次请求的任务 （适用全部接口）"}
  XCncRequestId *string `json:"x-cnc-request-id,omitempty" xml:"x-cnc-request-id,omitempty" require:"true"`
  // {"en":"The cnc-request-id corresponding to the request record you want to query this time", "zh_CN":"本次想要查询的请求记录对应的cnc-request-id"}
  CncRequestId *string `json:"cnc-request-id,omitempty" xml:"cnc-request-id,omitempty" require:"true"`
  // {"en":"The submission time for the request record that you want to query, for example: Thu, 09 Nov 2017 22:37:53 CST", "zh_CN":"本次想要查询的请求记录的提交时间，例如：Thu,09 Nov 2017 22:37:53 CST"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {"en":"Task execution results for asynchronous requests, including the following four states: WAIT: Indicates that the request is pending execution INPROGRESS: indicates that the request is in execution SUCCESS: Indicates that the request has been successfully executed FAIL: Indicates that the request execution failed Note: 1. If the query result is a failure of execution, there will be a mechanism or manual intervention in the backend of the system until the deployment is successful.If the result is the asynchronous request task of the new domain name, the system will have a resubmission mechanism or human intervention. 2. Modify the asynchronous request task for the domain name configuration, as well as add the domain name, there will be the mechanism of reintroduction and human intervention.", "zh_CN":"异步请求的任务执行结果，包括以下4种状态： WAIT：表示该请求等待执行 INPROGRESS：表示该请求执行中 SUCCESS：表示该请求已经执行成功 FAIL：表示该请求执行失败 注意： 1.新建域名的异步请求任务，通过本接口如果查询到结果是执行失败，系统后台会会有重提机制或者人工干预，直到部署成功。 2.修改域名配置的异步请求任务，和新增域名一样，也会有重提机制和人工干预。"}
  AsyncResult *string `json:"async-result,omitempty" xml:"async-result,omitempty" require:"true"`
  // {"en":"More information about task execution results for asynchronous requests", "zh_CN":"异步请求的任务执行结果的更多相关信息"}
  AsyncMessage *string `json:"async-message,omitempty" xml:"async-message,omitempty" require:"true"`
}

func (s QueryApiDeployServiceResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryApiDeployServiceResponse) GoString() string {
  return s.String()
}

func (s *QueryApiDeployServiceResponse) SetHttpStatus(v int) *QueryApiDeployServiceResponse {
  s.HttpStatus = &v
  return s
}

func (s *QueryApiDeployServiceResponse) SetXCncRequestId(v string) *QueryApiDeployServiceResponse {
  s.XCncRequestId = &v
  return s
}

func (s *QueryApiDeployServiceResponse) SetCncRequestId(v string) *QueryApiDeployServiceResponse {
  s.CncRequestId = &v
  return s
}

func (s *QueryApiDeployServiceResponse) SetTimestamp(v string) *QueryApiDeployServiceResponse {
  s.Timestamp = &v
  return s
}

func (s *QueryApiDeployServiceResponse) SetAsyncResult(v string) *QueryApiDeployServiceResponse {
  s.AsyncResult = &v
  return s
}

func (s *QueryApiDeployServiceResponse) SetAsyncMessage(v string) *QueryApiDeployServiceResponse {
  s.AsyncMessage = &v
  return s
}

type QueryApiDeployServicePaths struct {
  // {"en":"or each account request record, a unique cnc-request-id (for all API) is generated through which the execution results of each asynchronous request task can be queried", "zh_CN":"对于账号每一次请求记录，都会生成唯一的cnc-request-id（适用全部接口），通过该id可以查询每次异步请求任务的执行结果"}
  CncCequestId *string `json:"cnc-request-id,omitempty" xml:"cnc-request-id,omitempty" require:"true"`
}

func (s QueryApiDeployServicePaths) String() string {
  return tea.Prettify(s)
}

func (s QueryApiDeployServicePaths) GoString() string {
  return s.String()
}

func (s *QueryApiDeployServicePaths) SetCncCequestId(v string) *QueryApiDeployServicePaths {
  s.CncCequestId = &v
  return s
}

type QueryApiDeployServiceParameters struct {
}

func (s QueryApiDeployServiceParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryApiDeployServiceParameters) GoString() string {
  return s.String()
}

type QueryApiDeployServiceRequestHeader struct {
}

func (s QueryApiDeployServiceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryApiDeployServiceRequestHeader) GoString() string {
  return s.String()
}

type QueryApiDeployServiceResponseHeader struct {
}

func (s QueryApiDeployServiceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryApiDeployServiceResponseHeader) GoString() string {
  return s.String()
}




type QueryIpv6ConfigRequest struct {
}

func (s QueryIpv6ConfigRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryIpv6ConfigRequest) GoString() string {
  return s.String()
}

type QueryIpv6ConfigResponse struct {
  // {"en":"The domain whoes need query config.", "zh_CN":"需要查询配置的域名或域名id"}
  DomainName *string `json:"domain-name,omitempty" xml:"domain-name,omitempty" require:"true"`
  // {"en":"The domain whoes need query config.", "zh_CN":"需要查询配置的域名或域名id"}
  DomainId *string `json:"domain-id,omitempty" xml:"domain-id,omitempty" require:"true"`
  // {"en":"Whether a domain uses ipv6.
  // the value is true or false.", "zh_CN":"域名是否使用ipv6资源
  // 值为true或false"}
  UseIpv6 *string `json:"useIpv6,omitempty" xml:"useIpv6,omitempty" require:"true"`
}

func (s QueryIpv6ConfigResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryIpv6ConfigResponse) GoString() string {
  return s.String()
}

func (s *QueryIpv6ConfigResponse) SetDomainName(v string) *QueryIpv6ConfigResponse {
  s.DomainName = &v
  return s
}

func (s *QueryIpv6ConfigResponse) SetDomainId(v string) *QueryIpv6ConfigResponse {
  s.DomainId = &v
  return s
}

func (s *QueryIpv6ConfigResponse) SetUseIpv6(v string) *QueryIpv6ConfigResponse {
  s.UseIpv6 = &v
  return s
}

type QueryIpv6ConfigPaths struct {
  // {"en":"The domain whoes need query config.", "zh_CN":"需要查询配置的域名或域名id"}
  DomainName *string `json:"domain-name,omitempty" xml:"domain-name,omitempty" require:"true"`
}

func (s QueryIpv6ConfigPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryIpv6ConfigPaths) GoString() string {
  return s.String()
}

func (s *QueryIpv6ConfigPaths) SetDomainName(v string) *QueryIpv6ConfigPaths {
  s.DomainName = &v
  return s
}

type QueryIpv6ConfigParameters struct {
}

func (s QueryIpv6ConfigParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryIpv6ConfigParameters) GoString() string {
  return s.String()
}

type QueryIpv6ConfigRequestHeader struct {
}

func (s QueryIpv6ConfigRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryIpv6ConfigRequestHeader) GoString() string {
  return s.String()
}

type QueryIpv6ConfigResponseHeader struct {
}

func (s QueryIpv6ConfigResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryIpv6ConfigResponseHeader) GoString() string {
  return s.String()
}




type UpdateAmazonS3AuthorizationConfigRequest struct {
  // {"en":"Amazon S3 Access Authorization Configuration, parent node
  // 1. When you need to configure the Amazon S3 Access Authorization rules, this must be filled in.
  // 2. Configuration of clearing for <amazon-s3-access-authorization-rules/>.
  // 3.vodstream/download support, web/wsa does not support.
  // 4.Amason S3 and Aliyun OSS cannot be configured simultaneously.", "zh_CN":"Amazon S3鉴权配置，父标签
  // 1.需要设置Amazon S3鉴权时，此项必填
  // 2.为<amazon-s3-access-authorization-rules/>时清空Amazon S3鉴权配置的配置
  // 3.点播下载支持，网页wsa不支持
  // 4.Amason S3和Aliyun OSS不可同时配置"}
  AmazonS3AccessAuthorizationRules []*UpdateAmazonS3AuthorizationConfigRequestAmazonS3AccessAuthorizationRules `json:"amazon-s3-access-authorization-rules,omitempty" xml:"amazon-s3-access-authorization-rules,omitempty" require:"true" type:"Repeated"`
}

func (s UpdateAmazonS3AuthorizationConfigRequest) String() string {
  return tea.Prettify(s)
}

func (s UpdateAmazonS3AuthorizationConfigRequest) GoString() string {
  return s.String()
}

func (s *UpdateAmazonS3AuthorizationConfigRequest) SetAmazonS3AccessAuthorizationRules(v []*UpdateAmazonS3AuthorizationConfigRequestAmazonS3AccessAuthorizationRules) *UpdateAmazonS3AuthorizationConfigRequest {
  s.AmazonS3AccessAuthorizationRules = v
  return s
}

type UpdateAmazonS3AuthorizationConfigRequestAmazonS3AccessAuthorizationRules struct     {
  // {"en":"The url matching mode. If all matches, the input parameters can be configured as: .*", "zh_CN":"url匹配模式，支持正则，如果是全部匹配，入参可以配置为：.*"}
  PathPattern *string `json:"path-pattern,omitempty" xml:"path-pattern,omitempty" require:"true"`
  // {"en":"The exception url matching mode.", "zh_CN":"例外的url匹配模式，格式同path-pattern"}
  ExceptPathPattern *string `json:"except-path-pattern,omitempty" xml:"except-path-pattern,omitempty"`
  // {"en":"Define whether to add Authorization header when back to Amazon S3 source. Allowed  true and false.", "zh_CN":"是否添加鉴权头部，为true，则回源按照Amazon S3的算法添加 添加Authorization头部。
  // 允许值为true和false，默认为false。"}
  AddAuthorizationHeader *string `json:"add-authorization-header,omitempty" xml:"add-authorization-header,omitempty" require:"true"`
  // {"en":"access key", "zh_CN":"校验所需的密钥"}
  AccessKey *string `json:"access-key,omitempty" xml:"access-key,omitempty" require:"true"`
  // {"en":"access key id", "zh_CN":"校验所需的密钥ID"}
  AccessKeyId *string `json:"access-key-id,omitempty" xml:"access-key-id,omitempty" require:"true"`
  // {"en":"Signature version, default value is 2. Support 2 and 4.", "zh_CN":"算法版本，默认值2，目前仅支持2和4"}
  SignatureVersion *string `json:"signature-version,omitempty" xml:"signature-version,omitempty"`
  // {"en":"Authorization region. Please refer to the AWS official site for detail.", "zh_CN":"认证地区，见AWS官网区域说明"}
  Region *string `json:"region,omitempty" xml:"region,omitempty"`
  // {"en":"Data-id is to indicate a specific group configuration when the client has multiple groups of configurations. Data-id can be retrieved through a query interface. Note:
  // A. If data-id is passed, it means that one group of configuration items is specified to be modified, and no other group configuration items need to be modified.
  // B. If multiple groups of configurations are included, some of them are configured with data-id and others are not, then the expression of data-id is used to modify a specific group of configurations, and a new group of configurations is added on the original basis without the expression of data-id.
  // C. If the data-id is not transmitted, it means that the original configuration will be fully covered by this configuration.
  // D. If no configuration parameter is passed, only domain name and secondary label are passed, which means that all configuration of domain name secondary service corresponding to this interface is cleared.
  // E. If there is no specific configuration item in a set of configurations, the data-id must be filled in, and the value is the actual data-id, which means clearing the value of the corresponding data-id configuration item; it is not allowed that there is no specific configuration item or data-id in a set of configurations.", "zh_CN":"配置多组配置时，具体某组配置的id。data-id可以通过查询接口获取。 注意：
  // a、如果有传data-id，说明指定修改其中一组配置项内容，不需求修改其他组配置内容不需要入参； 
  // b、如果入参多组配置，其中有些组配置有传data-id，有些没有传，则有传data-id的表示修改具体某组配置，没有传data-id的表示在原来基础上新增一组配置； 
  // c、如果入参都没有传data-id,表示用本次的配置全量覆盖原先配置； 
  // d、如果入参没有传任何配置项参数，只传了域名和二级标签，表示清空这个接口对应域名二级服务所有配置； 
  // e、如果一组配置没有具体的配置项，则data-id必填，且值为实际存在的data-id，表示清空这个data-id对应配置项的值；不允许一组配置没有具体的配置项也没有data-id。"}
  DataId *int64 `json:"data-id,omitempty" xml:"data-id,omitempty"`
}

func (s UpdateAmazonS3AuthorizationConfigRequestAmazonS3AccessAuthorizationRules) String() string {
  return tea.Prettify(s)
}

func (s UpdateAmazonS3AuthorizationConfigRequestAmazonS3AccessAuthorizationRules) GoString() string {
  return s.String()
}

func (s *UpdateAmazonS3AuthorizationConfigRequestAmazonS3AccessAuthorizationRules) SetPathPattern(v string) *UpdateAmazonS3AuthorizationConfigRequestAmazonS3AccessAuthorizationRules {
  s.PathPattern = &v
  return s
}

func (s *UpdateAmazonS3AuthorizationConfigRequestAmazonS3AccessAuthorizationRules) SetExceptPathPattern(v string) *UpdateAmazonS3AuthorizationConfigRequestAmazonS3AccessAuthorizationRules {
  s.ExceptPathPattern = &v
  return s
}

func (s *UpdateAmazonS3AuthorizationConfigRequestAmazonS3AccessAuthorizationRules) SetAddAuthorizationHeader(v string) *UpdateAmazonS3AuthorizationConfigRequestAmazonS3AccessAuthorizationRules {
  s.AddAuthorizationHeader = &v
  return s
}

func (s *UpdateAmazonS3AuthorizationConfigRequestAmazonS3AccessAuthorizationRules) SetAccessKey(v string) *UpdateAmazonS3AuthorizationConfigRequestAmazonS3AccessAuthorizationRules {
  s.AccessKey = &v
  return s
}

func (s *UpdateAmazonS3AuthorizationConfigRequestAmazonS3AccessAuthorizationRules) SetAccessKeyId(v string) *UpdateAmazonS3AuthorizationConfigRequestAmazonS3AccessAuthorizationRules {
  s.AccessKeyId = &v
  return s
}

func (s *UpdateAmazonS3AuthorizationConfigRequestAmazonS3AccessAuthorizationRules) SetSignatureVersion(v string) *UpdateAmazonS3AuthorizationConfigRequestAmazonS3AccessAuthorizationRules {
  s.SignatureVersion = &v
  return s
}

func (s *UpdateAmazonS3AuthorizationConfigRequestAmazonS3AccessAuthorizationRules) SetRegion(v string) *UpdateAmazonS3AuthorizationConfigRequestAmazonS3AccessAuthorizationRules {
  s.Region = &v
  return s
}

func (s *UpdateAmazonS3AuthorizationConfigRequestAmazonS3AccessAuthorizationRules) SetDataId(v int64) *UpdateAmazonS3AuthorizationConfigRequestAmazonS3AccessAuthorizationRules {
  s.DataId = &v
  return s
}

type UpdateAmazonS3AuthorizationConfigResponse struct {
  // {"en":"The error code, when HTTPStatus is not 202, indicates the type of error the current request is calling.", "zh_CN":"错误代码，当HTTPStatus不为202时出现，表示当前请求调用的错误类型"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response information, when success is successful", "zh_CN":"响应信息，成功时为success"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s UpdateAmazonS3AuthorizationConfigResponse) String() string {
  return tea.Prettify(s)
}

func (s UpdateAmazonS3AuthorizationConfigResponse) GoString() string {
  return s.String()
}

func (s *UpdateAmazonS3AuthorizationConfigResponse) SetCode(v string) *UpdateAmazonS3AuthorizationConfigResponse {
  s.Code = &v
  return s
}

func (s *UpdateAmazonS3AuthorizationConfigResponse) SetMessage(v string) *UpdateAmazonS3AuthorizationConfigResponse {
  s.Message = &v
  return s
}

type UpdateAmazonS3AuthorizationConfigPaths struct {
  // {"en":"the domain whoes need query config", "zh_CN":"需要查询配置的域名或域名id"}
  DomainName *string `json:"domain-name,omitempty" xml:"domain-name,omitempty" require:"true"`
}

func (s UpdateAmazonS3AuthorizationConfigPaths) String() string {
  return tea.Prettify(s)
}

func (s UpdateAmazonS3AuthorizationConfigPaths) GoString() string {
  return s.String()
}

func (s *UpdateAmazonS3AuthorizationConfigPaths) SetDomainName(v string) *UpdateAmazonS3AuthorizationConfigPaths {
  s.DomainName = &v
  return s
}

type UpdateAmazonS3AuthorizationConfigParameters struct {
}

func (s UpdateAmazonS3AuthorizationConfigParameters) String() string {
  return tea.Prettify(s)
}

func (s UpdateAmazonS3AuthorizationConfigParameters) GoString() string {
  return s.String()
}

type UpdateAmazonS3AuthorizationConfigRequestHeader struct {
}

func (s UpdateAmazonS3AuthorizationConfigRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s UpdateAmazonS3AuthorizationConfigRequestHeader) GoString() string {
  return s.String()
}

type UpdateAmazonS3AuthorizationConfigResponseHeader struct {
}

func (s UpdateAmazonS3AuthorizationConfigResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s UpdateAmazonS3AuthorizationConfigResponseHeader) GoString() string {
  return s.String()
}




