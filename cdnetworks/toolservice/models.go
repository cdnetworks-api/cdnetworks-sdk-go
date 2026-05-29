package toolservice

// This file is auto-generated, don't edit it. Thanks.
import (
  "github.com/alibabacloud-go/tea/tea"
)

type IcpQueryServiceRequest struct {
}

func (s IcpQueryServiceRequest) String() string {
  return tea.Prettify(s)
}

func (s IcpQueryServiceRequest) GoString() string {
  return s.String()
}

type IcpQueryServiceResponse struct {
  // {'en':'domainIcpData', 'zh_CN':'域名备案信息'}
  DomainIcpDataList []*IcpQueryServiceResponseDomainIcpDataList `json:"domain-icp-data,omitempty" xml:"domain-icp-data,omitempty" require:"true" type:"Repeated"`
}

func (s IcpQueryServiceResponse) String() string {
  return tea.Prettify(s)
}

func (s IcpQueryServiceResponse) GoString() string {
  return s.String()
}

func (s *IcpQueryServiceResponse) SetDomainIcpDataList(v []*IcpQueryServiceResponseDomainIcpDataList) *IcpQueryServiceResponse {
  s.DomainIcpDataList = v
  return s
}

type IcpQueryServiceResponseDomainIcpDataList struct     {
  // {"en":"Domain", "zh_CN":"域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {"en":"Registration No.", "zh_CN":"备案号"}
  IcpNumber *string `json:"icp-number,omitempty" xml:"icp-number,omitempty" require:"true"`
}

func (s IcpQueryServiceResponseDomainIcpDataList) String() string {
  return tea.Prettify(s)
}

func (s IcpQueryServiceResponseDomainIcpDataList) GoString() string {
  return s.String()
}

func (s *IcpQueryServiceResponseDomainIcpDataList) SetDomain(v string) *IcpQueryServiceResponseDomainIcpDataList {
  s.Domain = &v
  return s
}

func (s *IcpQueryServiceResponseDomainIcpDataList) SetIcpNumber(v string) *IcpQueryServiceResponseDomainIcpDataList {
  s.IcpNumber = &v
  return s
}

type IcpQueryServicePaths struct {
}

func (s IcpQueryServicePaths) String() string {
  return tea.Prettify(s)
}

func (s IcpQueryServicePaths) GoString() string {
  return s.String()
}

type IcpQueryServiceParameters struct {
  // {"en":"Domain names, multiple domain names are separated by English semicolons. The maximum number of domain names is 20.", "zh_CN":"域名，多个以英文分号分隔。域名数上限为20个。"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
}

func (s IcpQueryServiceParameters) String() string {
  return tea.Prettify(s)
}

func (s IcpQueryServiceParameters) GoString() string {
  return s.String()
}

func (s *IcpQueryServiceParameters) SetDomain(v string) *IcpQueryServiceParameters {
  s.Domain = &v
  return s
}

type IcpQueryServiceRequestHeader struct {
}

func (s IcpQueryServiceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s IcpQueryServiceRequestHeader) GoString() string {
  return s.String()
}

type IcpQueryServiceResponseHeader struct {
}

func (s IcpQueryServiceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s IcpQueryServiceResponseHeader) GoString() string {
  return s.String()
}




type QueryBandwidthLimitTaskListServiceRequest struct {
}

func (s QueryBandwidthLimitTaskListServiceRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryBandwidthLimitTaskListServiceRequest) GoString() string {
  return s.String()
}

type QueryBandwidthLimitTaskListServiceRequestHeader struct {
}

func (s QueryBandwidthLimitTaskListServiceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryBandwidthLimitTaskListServiceRequestHeader) GoString() string {
  return s.String()
}

type QueryBandwidthLimitTaskListServicePaths struct {
}

func (s QueryBandwidthLimitTaskListServicePaths) String() string {
  return tea.Prettify(s)
}

func (s QueryBandwidthLimitTaskListServicePaths) GoString() string {
  return s.String()
}

type QueryBandwidthLimitTaskListServiceParameters struct {
}

func (s QueryBandwidthLimitTaskListServiceParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryBandwidthLimitTaskListServiceParameters) GoString() string {
  return s.String()
}

type QueryBandwidthLimitTaskListServiceResponse struct {
  // {"en":"result","zh_CN":"结果"}
  Result []*QueryBandwidthLimitTaskListServiceResponseResult `json:"result,omitempty" xml:"result,omitempty" require:"true" type:"Repeated"`
}

func (s QueryBandwidthLimitTaskListServiceResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryBandwidthLimitTaskListServiceResponse) GoString() string {
  return s.String()
}

func (s *QueryBandwidthLimitTaskListServiceResponse) SetResult(v []*QueryBandwidthLimitTaskListServiceResponseResult) *QueryBandwidthLimitTaskListServiceResponse {
  s.Result = v
  return s
}

type QueryBandwidthLimitTaskListServiceResponseResult struct     {
  // {"en":"Domain","zh_CN":"域名"}
  DomainName *string `json:"domainName,omitempty" xml:"domainName,omitempty" require:"true"`
  // {"en":"Maximum bandwidth set","zh_CN":"设置的最大带宽值"}
  BandwidthLimit *int `json:"bandwidthLimit,omitempty" xml:"bandwidthLimit,omitempty" require:"true"`
  // {"en":"Task name","zh_CN":"任务名称"}
  TaskName *string `json:"taskName,omitempty" xml:"taskName,omitempty" require:"true"`
}

func (s QueryBandwidthLimitTaskListServiceResponseResult) String() string {
  return tea.Prettify(s)
}

func (s QueryBandwidthLimitTaskListServiceResponseResult) GoString() string {
  return s.String()
}

func (s *QueryBandwidthLimitTaskListServiceResponseResult) SetDomainName(v string) *QueryBandwidthLimitTaskListServiceResponseResult {
  s.DomainName = &v
  return s
}

func (s *QueryBandwidthLimitTaskListServiceResponseResult) SetBandwidthLimit(v int) *QueryBandwidthLimitTaskListServiceResponseResult {
  s.BandwidthLimit = &v
  return s
}

func (s *QueryBandwidthLimitTaskListServiceResponseResult) SetTaskName(v string) *QueryBandwidthLimitTaskListServiceResponseResult {
  s.TaskName = &v
  return s
}

type QueryBandwidthLimitTaskListServiceResponseHeader struct {
}

func (s QueryBandwidthLimitTaskListServiceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryBandwidthLimitTaskListServiceResponseHeader) GoString() string {
  return s.String()
}




type IpDomainServiceRequest struct {
  // {"en":"IP", "zh_CN":"IP"}
  Ip []*string `json:"ip,omitempty" xml:"ip,omitempty" require:"true" type:"Repeated"`
}

func (s IpDomainServiceRequest) String() string {
  return tea.Prettify(s)
}

func (s IpDomainServiceRequest) GoString() string {
  return s.String()
}

func (s *IpDomainServiceRequest) SetIp(v []*string) *IpDomainServiceRequest {
  s.Ip = v
  return s
}

type IpDomainServiceResponse struct {
  // {"en":"request result status code", "zh_CN":"请求结果状态码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Request result information", "zh_CN":"请求结果信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Detailed data on the results of the request", "zh_CN":"请求结果的详细数据"}
  Data []*IpDomainServiceResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s IpDomainServiceResponse) String() string {
  return tea.Prettify(s)
}

func (s IpDomainServiceResponse) GoString() string {
  return s.String()
}

func (s *IpDomainServiceResponse) SetCode(v string) *IpDomainServiceResponse {
  s.Code = &v
  return s
}

func (s *IpDomainServiceResponse) SetMessage(v string) *IpDomainServiceResponse {
  s.Message = &v
  return s
}

func (s *IpDomainServiceResponse) SetData(v []*IpDomainServiceResponseData) *IpDomainServiceResponse {
  s.Data = v
  return s
}

type IpDomainServiceResponseData struct     {
  // {"en":"ip", "zh_CN":"IP名称"}
  Ip *string `json:"ip,omitempty" xml:"ip,omitempty" require:"true"`
  // {"en":"Whether to use:
  // 
  //   idle --IP not used yet;
  //   runing -- IP in use;
  //   out of range -- IP is not in a queryable range", "zh_CN":"是否使用:
  //   idle -- 暂未使用;
  //   runing -- 使用中;
  //   out of range -- 不在查询范围内的ip"}
  Status *string `json:"status,omitempty" xml:"status,omitempty" require:"true"`
  // {"en":"List of domain using this IP.The domain list of the IP that was idle or out of range was empty", "zh_CN":"用该IP的域名列表,未使用的ip/不在查询范围内的ip,域名列表为空"}
  DomainList []*string `json:"domainList,omitempty" xml:"domainList,omitempty" require:"true" type:"Repeated"`
}

func (s IpDomainServiceResponseData) String() string {
  return tea.Prettify(s)
}

func (s IpDomainServiceResponseData) GoString() string {
  return s.String()
}

func (s *IpDomainServiceResponseData) SetIp(v string) *IpDomainServiceResponseData {
  s.Ip = &v
  return s
}

func (s *IpDomainServiceResponseData) SetStatus(v string) *IpDomainServiceResponseData {
  s.Status = &v
  return s
}

func (s *IpDomainServiceResponseData) SetDomainList(v []*string) *IpDomainServiceResponseData {
  s.DomainList = v
  return s
}

type IpDomainServicePaths struct {
}

func (s IpDomainServicePaths) String() string {
  return tea.Prettify(s)
}

func (s IpDomainServicePaths) GoString() string {
  return s.String()
}

type IpDomainServiceParameters struct {
}

func (s IpDomainServiceParameters) String() string {
  return tea.Prettify(s)
}

func (s IpDomainServiceParameters) GoString() string {
  return s.String()
}

type IpDomainServiceRequestHeader struct {
}

func (s IpDomainServiceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s IpDomainServiceRequestHeader) GoString() string {
  return s.String()
}

type IpDomainServiceResponseHeader struct {
}

func (s IpDomainServiceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s IpDomainServiceResponseHeader) GoString() string {
  return s.String()
}




type BandwidthLimitServiceRequest struct {
  // {"en":"Task: Limits to the number of tasks can be adjusted depending on different accounts. The default value is 3","zh_CN":"任务：任务个数限制根据账号可调，默认为3个"}
  Task *string `json:"task,omitempty" xml:"task,omitempty"`
  // {"en":"Operation type: \nenable: set, update or enable bandwidth limit; \ndisable: to disable bandwidth limit","zh_CN":"操作类型：\nenable 设置、更新或开启带宽限制，\ndisable 关闭带宽限制"}
  Action *string `json:"action,omitempty" xml:"action,omitempty"`
  // {"en":"Domain list.","zh_CN":"域名列表"}
  DomainList *string `json:"domain-list,omitempty" xml:"domain-list,omitempty"`
  // {"en":"Domain, must follow regular expression rule of (([\w-]{1,62})?(\.[\w-]{1,62})+)","zh_CN":"域名，必须符合正则(([\w-]{1,62})?(\.[\w-]{1,62})+)"}
  DomainName *string `json:"domain-name,omitempty" xml:"domain-name,omitempty"`
  // {"en":"1. Cancel the control ratio (the proportion to which the threshold is reduced, and cancel the control bandwidth), calculated as a percentage;\n\n2. The input value is a positive integer. When no parameters are passed, the default is 60.","zh_CN":"1.取消控制比例（阈值降到多少比例，取消控制带宽），按百分比计算；\n\n2.输入值为正整数。未传参时，默认为40。"}
  CtrlMinRatio *int `json:"ctrlMinRatio,omitempty" xml:"ctrlMinRatio,omitempty"`
  // {"en":"1. Control the effective ratio (what ratio the threshold reaches and start to control the bandwidth), calculated as a percentage;\n2. The input value is a positive integer. When no parameter is passed, the default is 60;\n3. The effective ratio of control is greater than the ratio of canceled control;\n4. The effective control ratio and the cancel control ratio need to be paired and configured.","zh_CN":"1.控制生效比例（ 阈值达到多少比例，开始控制带宽），按百分比计算；\n\n2.输入值为正整数。未传参时，默认为60；\n\n3.控制生效比例 要大于 取消控制比例；\n\n4.控制生效比例 和 取消控制比例 需配对配置。"}
  CtrlInitRatio *int `json:"ctrlInitRatio,omitempty" xml:"ctrlInitRatio,omitempty"`
  // {"en":"Bandwidth limit, positive integer, unit is Mbps. This filed is required when action is enable","zh_CN":"带宽限制值，为正整数，单位为Mbps，当action为enable时为必选项"}
  BandwidthLimit *int `json:"bandwidth-limit,omitempty" xml:"bandwidth-limit,omitempty"`
}

func (s BandwidthLimitServiceRequest) String() string {
  return tea.Prettify(s)
}

func (s BandwidthLimitServiceRequest) GoString() string {
  return s.String()
}

func (s *BandwidthLimitServiceRequest) SetTask(v string) *BandwidthLimitServiceRequest {
  s.Task = &v
  return s
}

func (s *BandwidthLimitServiceRequest) SetAction(v string) *BandwidthLimitServiceRequest {
  s.Action = &v
  return s
}

func (s *BandwidthLimitServiceRequest) SetDomainList(v string) *BandwidthLimitServiceRequest {
  s.DomainList = &v
  return s
}

func (s *BandwidthLimitServiceRequest) SetDomainName(v string) *BandwidthLimitServiceRequest {
  s.DomainName = &v
  return s
}

func (s *BandwidthLimitServiceRequest) SetCtrlMinRatio(v int) *BandwidthLimitServiceRequest {
  s.CtrlMinRatio = &v
  return s
}

func (s *BandwidthLimitServiceRequest) SetCtrlInitRatio(v int) *BandwidthLimitServiceRequest {
  s.CtrlInitRatio = &v
  return s
}

func (s *BandwidthLimitServiceRequest) SetBandwidthLimit(v int) *BandwidthLimitServiceRequest {
  s.BandwidthLimit = &v
  return s
}

type BandwidthLimitServiceRequestHeader struct {
}

func (s BandwidthLimitServiceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s BandwidthLimitServiceRequestHeader) GoString() string {
  return s.String()
}

type BandwidthLimitServicePaths struct {
}

func (s BandwidthLimitServicePaths) String() string {
  return tea.Prettify(s)
}

func (s BandwidthLimitServicePaths) GoString() string {
  return s.String()
}

type BandwidthLimitServiceParameters struct {
}

func (s BandwidthLimitServiceParameters) String() string {
  return tea.Prettify(s)
}

func (s BandwidthLimitServiceParameters) GoString() string {
  return s.String()
}

type BandwidthLimitServiceResponse struct {
  // {"en":"Task ID","zh_CN":"任务ID"}
  TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty" require:"true"`
  // {"en":"Domain","zh_CN":"域名"}
  DomainName *string `json:"domainName,omitempty" xml:"domainName,omitempty" require:"true"`
  // {"en":"Status codes, 1 for success and 0   means failed","zh_CN":"状态码，1：成功，0：失败"}
  Code *int `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Description of results","zh_CN":"结果描述信息"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
}

func (s BandwidthLimitServiceResponse) String() string {
  return tea.Prettify(s)
}

func (s BandwidthLimitServiceResponse) GoString() string {
  return s.String()
}

func (s *BandwidthLimitServiceResponse) SetTaskId(v string) *BandwidthLimitServiceResponse {
  s.TaskId = &v
  return s
}

func (s *BandwidthLimitServiceResponse) SetDomainName(v string) *BandwidthLimitServiceResponse {
  s.DomainName = &v
  return s
}

func (s *BandwidthLimitServiceResponse) SetCode(v int) *BandwidthLimitServiceResponse {
  s.Code = &v
  return s
}

func (s *BandwidthLimitServiceResponse) SetMsg(v string) *BandwidthLimitServiceResponse {
  s.Msg = &v
  return s
}

type BandwidthLimitServiceResponseHeader struct {
}

func (s BandwidthLimitServiceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s BandwidthLimitServiceResponseHeader) GoString() string {
  return s.String()
}




type QueryAllBandwidthLimitTaskListServiceRequest struct {
  // {"en":"Domain:\n1. The maximum number of domain is 100 by default (you can contact technical support for adjustment);\n2. Automatically filter out invalid domain (an illegal domain will be filtered, and the query result will only return the data of valid domains).","zh_CN":"域名:\n1.可传递域名数量上限默认为100个(可联系技术支持调整);\n2.自动过滤掉无效域名(如传递非法域名,会被过滤,查询结果只返回有效域名的数据)。"}
  Domain []*string `json:"domain,omitempty" xml:"domain,omitempty" require:"true" type:"Repeated"`
  // {"en":"Whether to ignore task status.\n0: No, only return the configuration whose task status is enabled; default: 0;\n1: Yes, return whether the task status is open or closed.","zh_CN":"是否忽略任务状态。0:否,只返回任务状态为开启的配置;默认:0; 1:是,不论任务状态是开启还是关闭都返回。"}
  IgnoreTaskStatus *string `json:"ignoreTaskStatus,omitempty" xml:"ignoreTaskStatus,omitempty"`
  // {"en":"Whether the returned data contains all customer domain names involved in the task, the default is 0;\n0: no;\n1: Yes.","zh_CN":"返回数据是否包含任务涉及的所有客户域名。0:否;默认:0;1:是。"}
  ContainDomain *string `json:"containDomain,omitempty" xml:"containDomain,omitempty"`
}

func (s QueryAllBandwidthLimitTaskListServiceRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryAllBandwidthLimitTaskListServiceRequest) GoString() string {
  return s.String()
}

func (s *QueryAllBandwidthLimitTaskListServiceRequest) SetDomain(v []*string) *QueryAllBandwidthLimitTaskListServiceRequest {
  s.Domain = v
  return s
}

func (s *QueryAllBandwidthLimitTaskListServiceRequest) SetIgnoreTaskStatus(v string) *QueryAllBandwidthLimitTaskListServiceRequest {
  s.IgnoreTaskStatus = &v
  return s
}

func (s *QueryAllBandwidthLimitTaskListServiceRequest) SetContainDomain(v string) *QueryAllBandwidthLimitTaskListServiceRequest {
  s.ContainDomain = &v
  return s
}

type QueryAllBandwidthLimitTaskListServiceRequestHeader struct {
}

func (s QueryAllBandwidthLimitTaskListServiceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryAllBandwidthLimitTaskListServiceRequestHeader) GoString() string {
  return s.String()
}

type QueryAllBandwidthLimitTaskListServicePaths struct {
}

func (s QueryAllBandwidthLimitTaskListServicePaths) String() string {
  return tea.Prettify(s)
}

func (s QueryAllBandwidthLimitTaskListServicePaths) GoString() string {
  return s.String()
}

type QueryAllBandwidthLimitTaskListServiceParameters struct {
}

func (s QueryAllBandwidthLimitTaskListServiceParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryAllBandwidthLimitTaskListServiceParameters) GoString() string {
  return s.String()
}

type QueryAllBandwidthLimitTaskListServiceResponse struct {
  // {"en":"request result status code","zh_CN":"请求结果状态码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Request result information","zh_CN":"请求结果信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"","zh_CN":""}
  Data []*QueryAllBandwidthLimitTaskListServiceResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s QueryAllBandwidthLimitTaskListServiceResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryAllBandwidthLimitTaskListServiceResponse) GoString() string {
  return s.String()
}

func (s *QueryAllBandwidthLimitTaskListServiceResponse) SetCode(v string) *QueryAllBandwidthLimitTaskListServiceResponse {
  s.Code = &v
  return s
}

func (s *QueryAllBandwidthLimitTaskListServiceResponse) SetMessage(v string) *QueryAllBandwidthLimitTaskListServiceResponse {
  s.Message = &v
  return s
}

func (s *QueryAllBandwidthLimitTaskListServiceResponse) SetData(v []*QueryAllBandwidthLimitTaskListServiceResponseData) *QueryAllBandwidthLimitTaskListServiceResponse {
  s.Data = v
  return s
}

type QueryAllBandwidthLimitTaskListServiceResponseData struct     {
  // {"en":"domain","zh_CN":"客户域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {"en":"Whether there is a configuration control task, 0: no, 1: yes","zh_CN":"是否有配置控制任务,0表示没有,1表示有"}
  IsExist *string `json:"isExist,omitempty" xml:"isExist,omitempty" require:"true"`
  // {"en":"","zh_CN":""}
  Content []*QueryAllBandwidthLimitTaskListServiceResponseDataContent `json:"content,omitempty" xml:"content,omitempty" require:"true" type:"Repeated"`
}

func (s QueryAllBandwidthLimitTaskListServiceResponseData) String() string {
  return tea.Prettify(s)
}

func (s QueryAllBandwidthLimitTaskListServiceResponseData) GoString() string {
  return s.String()
}

func (s *QueryAllBandwidthLimitTaskListServiceResponseData) SetDomain(v string) *QueryAllBandwidthLimitTaskListServiceResponseData {
  s.Domain = &v
  return s
}

func (s *QueryAllBandwidthLimitTaskListServiceResponseData) SetIsExist(v string) *QueryAllBandwidthLimitTaskListServiceResponseData {
  s.IsExist = &v
  return s
}

func (s *QueryAllBandwidthLimitTaskListServiceResponseData) SetContent(v []*QueryAllBandwidthLimitTaskListServiceResponseDataContent) *QueryAllBandwidthLimitTaskListServiceResponseData {
  s.Content = v
  return s
}

type QueryAllBandwidthLimitTaskListServiceResponseDataContent struct     {
  // {"en":"taskName","zh_CN":"任务名称"}
  TaskName *string `json:"taskName,omitempty" xml:"taskName,omitempty" require:"true"`
  // {"en":"Domain name configuration task types: 1. Static bandwidth control task, 2. Bandwidth buyout task, 3. Flow buyout task, 4. Request number buyout task, 5. Redundant pool speed limit task, 6. Back-to-source task, 7. POP running high scheduling task (this kind of task is quite special, the domain name is a global quantity, so as long as there is configuration, it will be enabled by default), 8. IP ban task","zh_CN":"域名配置的任务类型:1静态带宽控制任务,2带宽买断任务,3流量买断任务,4请求数买断任务,5冗余池限速任务,6回源任务,7POP跑高调度任务(此种任务比较特殊,域名为全局量,所以只要有配置,就默认开启),8IP封禁任务"}
  TaskType *string `json:"taskType,omitempty" xml:"taskType,omitempty" require:"true"`
  // {"en":"Task status, 0: task is off, 1: task is on.","zh_CN":"任务状态,0表示任务关闭,1表示任务开启。"}
  TaskStatus *string `json:"taskStatus,omitempty" xml:"taskStatus,omitempty" require:"true"`
  // {"en":"Control strategy. 0 means squid default, 1 rejects, 2 when taskType=8, is ladder blocked, the rest is redirected ip, 3 when taskType=6, is backup source, the rest is speed limit = rejection + maximum download rate, 4 speed limit = rejection + timeout to disconnect, 5 does not process, controls each connection at minimum speed, does not process the excess part, 6 backups, 7 redirects domain name, 8 redirects URL","zh_CN":"控制策略。0表示squid默认,1拒绝,2当taskType=8时,为阶梯封禁,其余为重定向ip,3当taskType=6时,为主备回源,其余为限速=拒绝+最大下载速率,4限速=拒绝+超时断开连接,5不处理,按最小速率控制每个连接,超出部分不处理,6回源,7重定向域名,8重定向URL"}
  CtrlMode *string `json:"ctrlMode,omitempty" xml:"ctrlMode,omitempty" require:"true"`
  // {"en":"Bandwidth limit value, when taskType=1,2,5,6, in Mbps, when taskType=3, in G, when taskType=4, in MH (millions), when taskType=7, in -1, there is no bandwidth limit value, when taskType=8, in seconds.","zh_CN":"带宽限制值,当taskType=1,2,5,6时,单位为Mbps,当taskType=3,单位为G,当taskType=4,单位为MH(百万个),当taskType=7,为-1,表示没有带宽限制值,当taskType=8,单位为次。"}
  CtrlValue *string `json:"ctrlValue,omitempty" xml:"ctrlValue,omitempty" require:"true"`
  // {"en":"List of customer domains involved under the task. Values are only available when entering ContainOrichannelName=1.","zh_CN":"任务下涉及的客户域名列表。只有在入参containOrichannelName=1的时候有值。"}
  DomainList []*string `json:"domainList,omitempty" xml:"domainList,omitempty" require:"true" type:"Repeated"`
}

func (s QueryAllBandwidthLimitTaskListServiceResponseDataContent) String() string {
  return tea.Prettify(s)
}

func (s QueryAllBandwidthLimitTaskListServiceResponseDataContent) GoString() string {
  return s.String()
}

func (s *QueryAllBandwidthLimitTaskListServiceResponseDataContent) SetTaskName(v string) *QueryAllBandwidthLimitTaskListServiceResponseDataContent {
  s.TaskName = &v
  return s
}

func (s *QueryAllBandwidthLimitTaskListServiceResponseDataContent) SetTaskType(v string) *QueryAllBandwidthLimitTaskListServiceResponseDataContent {
  s.TaskType = &v
  return s
}

func (s *QueryAllBandwidthLimitTaskListServiceResponseDataContent) SetTaskStatus(v string) *QueryAllBandwidthLimitTaskListServiceResponseDataContent {
  s.TaskStatus = &v
  return s
}

func (s *QueryAllBandwidthLimitTaskListServiceResponseDataContent) SetCtrlMode(v string) *QueryAllBandwidthLimitTaskListServiceResponseDataContent {
  s.CtrlMode = &v
  return s
}

func (s *QueryAllBandwidthLimitTaskListServiceResponseDataContent) SetCtrlValue(v string) *QueryAllBandwidthLimitTaskListServiceResponseDataContent {
  s.CtrlValue = &v
  return s
}

func (s *QueryAllBandwidthLimitTaskListServiceResponseDataContent) SetDomainList(v []*string) *QueryAllBandwidthLimitTaskListServiceResponseDataContent {
  s.DomainList = v
  return s
}

type QueryAllBandwidthLimitTaskListServiceResponseHeader struct {
}

func (s QueryAllBandwidthLimitTaskListServiceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryAllBandwidthLimitTaskListServiceResponseHeader) GoString() string {
  return s.String()
}




