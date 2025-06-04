package domainmanagement

// This file is auto-generated, don't edit it. Thanks.
import (
  "github.com/alibabacloud-go/tea/tea"
)

type UpdateCustomerAnycastIPRecordStatusRequest struct {
  // {"en":"Ip list", "zh_CN":"ip 列表"}
  Ips []*string `json:"ips,omitempty" xml:"ips,omitempty" require:"true" type:"Repeated"`
  // {"en":"Record Status, For example, lock or unlock; data of length 1 or 2 can be passed.", "zh_CN":"记录状态，例如锁定、非锁定等，可以传长度为1或2的数据"}
  RecordStatus *string `json:"recordStatus,omitempty" xml:"recordStatus,omitempty" require:"true"`
}

func (s UpdateCustomerAnycastIPRecordStatusRequest) String() string {
  return tea.Prettify(s)
}

func (s UpdateCustomerAnycastIPRecordStatusRequest) GoString() string {
  return s.String()
}

func (s *UpdateCustomerAnycastIPRecordStatusRequest) SetIps(v []*string) *UpdateCustomerAnycastIPRecordStatusRequest {
  s.Ips = v
  return s
}

func (s *UpdateCustomerAnycastIPRecordStatusRequest) SetRecordStatus(v string) *UpdateCustomerAnycastIPRecordStatusRequest {
  s.RecordStatus = &v
  return s
}

type UpdateCustomerAnycastIPRecordStatusResponse struct {
  // {"en":"The error code that appears when the HTTP status is not 202, indicating the type of error for the current request.", "zh_CN":"错误代码，当HTTPStatus不为202时出现，表示当前请求调用的错误类型"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response information, when success is successful", "zh_CN":"响应信息，成功时为success"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s UpdateCustomerAnycastIPRecordStatusResponse) String() string {
  return tea.Prettify(s)
}

func (s UpdateCustomerAnycastIPRecordStatusResponse) GoString() string {
  return s.String()
}

func (s *UpdateCustomerAnycastIPRecordStatusResponse) SetCode(v string) *UpdateCustomerAnycastIPRecordStatusResponse {
  s.Code = &v
  return s
}

func (s *UpdateCustomerAnycastIPRecordStatusResponse) SetMessage(v string) *UpdateCustomerAnycastIPRecordStatusResponse {
  s.Message = &v
  return s
}

type UpdateCustomerAnycastIPRecordStatusPaths struct {
}

func (s UpdateCustomerAnycastIPRecordStatusPaths) String() string {
  return tea.Prettify(s)
}

func (s UpdateCustomerAnycastIPRecordStatusPaths) GoString() string {
  return s.String()
}

type UpdateCustomerAnycastIPRecordStatusParameters struct {
}

func (s UpdateCustomerAnycastIPRecordStatusParameters) String() string {
  return tea.Prettify(s)
}

func (s UpdateCustomerAnycastIPRecordStatusParameters) GoString() string {
  return s.String()
}

type UpdateCustomerAnycastIPRecordStatusRequestHeader struct {
}

func (s UpdateCustomerAnycastIPRecordStatusRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s UpdateCustomerAnycastIPRecordStatusRequestHeader) GoString() string {
  return s.String()
}

type UpdateCustomerAnycastIPRecordStatusResponseHeader struct {
}

func (s UpdateCustomerAnycastIPRecordStatusResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s UpdateCustomerAnycastIPRecordStatusResponseHeader) GoString() string {
  return s.String()
}




type PreDeployChangeServerConfigRequest struct {
  // {"en":"Change servers configuration, parent tag
  // 1. This must be filled when the hotlinking configuration of streaming media needs to be set
  // 2. Empty the configuration for <change-servers/>", "zh_CN":"【接入域名跳转】
  // 注意：
  // 1、需要取消【接入域名跳转】时，可以传入空节点<change-servers></change-servers>。
  // 2、表示需要设置【接入域名跳转】，此项必填"}
  ChangeServers []*PreDeployChangeServerConfigRequestChangeServers `json:"change-servers,omitempty" xml:"change-servers,omitempty" require:"true" type:"Repeated"`
}

func (s PreDeployChangeServerConfigRequest) String() string {
  return tea.Prettify(s)
}

func (s PreDeployChangeServerConfigRequest) GoString() string {
  return s.String()
}

func (s *PreDeployChangeServerConfigRequest) SetChangeServers(v []*PreDeployChangeServerConfigRequestChangeServers) *PreDeployChangeServerConfigRequest {
  s.ChangeServers = v
  return s
}

type PreDeployChangeServerConfigRequestChangeServers struct     {
  // {"en":"If it is a universal domain name, set it to a universal domain name, for example, *.56.com.", "zh_CN":"如果是泛域名，需要填写为泛域名，例如：*.56.com"}
  TargetServer *string `json:"target-server,omitempty" xml:"target-server,omitempty"`
  // {"en":"Data-id is to indicate a specific group configuration when the client has multiple groups of configurations. Data-id can be retrieved through a query interface. Note: 
  // A. If data-id is passed, it means that one group of configuration items is specified to be modified, and no other group configuration items need to be modified. 
  // B. If multiple groups of configurations are included, some of them are configured with data-id and others are not, then the expression of data-id is used to modify a specific group of configurations, and a new group of configurations is added on the original basis without the expression of data-id. 
  // C. If the data-id is not transmitted, it means that the original configuration will be fully covered by this configuration. 
  // D. If no configuration parameter is passed, only domain name and secondary label are passed, which means that all configuration of domain name secondary service corresponding to this interface is cleared. 
  // E. If there is no specific configuration item in a set of configurations, the data-id must be filled in, and the value is the actual data-id, which means clearing the value of the corresponding data-id configuration item; it is not allowed that there is no specific configuration item or data-id in a set of configurations.", 
  //       "zh_CN":"配置多组配置时，具体某组配置的id。dataId可以通过查询接口获取。 注意： 
  // a、如果有传dataId，说明指定修改其中一组配置项内容，不需求修改其他组配置内容不需要入参； 
  // b、如果入参多组配置，其中有些组配置有传dataId，有些没有传，则有传dataId的表示修改具体某组配置，没有传dataId的表示在原来基础上新增一组配置； 
  // c、如果入参都没有传dataId,表示用本次的配置全量覆盖原先配置； d、如果入参没有传任何配置项参数，只传了域名和二级标签，表示清空这个接口对应域名二级服务所有配置； 
  // e、如果一组配置没有具体的配置项，则dataId必填，且值为实际存在的dataId，表示清空这个dataId对应配置项的值；不允许一组配置没有具体的配置项也没有dataId。
  // "}
  DataId *int `json:"dataId,omitempty" xml:"dataId,omitempty"`
}

func (s PreDeployChangeServerConfigRequestChangeServers) String() string {
  return tea.Prettify(s)
}

func (s PreDeployChangeServerConfigRequestChangeServers) GoString() string {
  return s.String()
}

func (s *PreDeployChangeServerConfigRequestChangeServers) SetTargetServer(v string) *PreDeployChangeServerConfigRequestChangeServers {
  s.TargetServer = &v
  return s
}

func (s *PreDeployChangeServerConfigRequestChangeServers) SetDataId(v int) *PreDeployChangeServerConfigRequestChangeServers {
  s.DataId = &v
  return s
}

type PreDeployChangeServerConfigResponse struct {
  // {"en":"The error code", "zh_CN":"错误码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"The message body", "zh_CN":"消息体"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Returns the body of the data.", "zh_CN":"返回数据体"}
  Data *PreDeployChangeServerConfigResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
  // {"en":"Uniquely labeled id for querying each requested task (for all interfaces)", "zh_CN":"唯一标示的id，用于查询每次请求的任务 （适用全部接口）"}
  XCncRequestId *string `json:"x-cnc-request-id,omitempty" xml:"x-cnc-request-id,omitempty" require:"true"`
}

func (s PreDeployChangeServerConfigResponse) String() string {
  return tea.Prettify(s)
}

func (s PreDeployChangeServerConfigResponse) GoString() string {
  return s.String()
}

func (s *PreDeployChangeServerConfigResponse) SetCode(v string) *PreDeployChangeServerConfigResponse {
  s.Code = &v
  return s
}

func (s *PreDeployChangeServerConfigResponse) SetMessage(v string) *PreDeployChangeServerConfigResponse {
  s.Message = &v
  return s
}

func (s *PreDeployChangeServerConfigResponse) SetData(v *PreDeployChangeServerConfigResponseData) *PreDeployChangeServerConfigResponse {
  s.Data = v
  return s
}

func (s *PreDeployChangeServerConfigResponse) SetXCncRequestId(v string) *PreDeployChangeServerConfigResponse {
  s.XCncRequestId = &v
  return s
}

type PreDeployChangeServerConfigResponseData struct {
  // {"en":"The preliminary deployment id", "zh_CN":"预部署id"}
  PreDeployId *string `json:"preDeployId,omitempty" xml:"preDeployId,omitempty" require:"true"`
}

func (s PreDeployChangeServerConfigResponseData) String() string {
  return tea.Prettify(s)
}

func (s PreDeployChangeServerConfigResponseData) GoString() string {
  return s.String()
}

func (s *PreDeployChangeServerConfigResponseData) SetPreDeployId(v string) *PreDeployChangeServerConfigResponseData {
  s.PreDeployId = &v
  return s
}

type PreDeployChangeServerConfigPaths struct {
  // {"en":"The domain whoes need query config.", "zh_CN":"需要查询配置的域名或域名id"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
}

func (s PreDeployChangeServerConfigPaths) String() string {
  return tea.Prettify(s)
}

func (s PreDeployChangeServerConfigPaths) GoString() string {
  return s.String()
}

func (s *PreDeployChangeServerConfigPaths) SetDomain(v string) *PreDeployChangeServerConfigPaths {
  s.Domain = &v
  return s
}

type PreDeployChangeServerConfigParameters struct {
}

func (s PreDeployChangeServerConfigParameters) String() string {
  return tea.Prettify(s)
}

func (s PreDeployChangeServerConfigParameters) GoString() string {
  return s.String()
}

type PreDeployChangeServerConfigRequestHeader struct {
}

func (s PreDeployChangeServerConfigRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s PreDeployChangeServerConfigRequestHeader) GoString() string {
  return s.String()
}

type PreDeployChangeServerConfigResponseHeader struct {
}

func (s PreDeployChangeServerConfigResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s PreDeployChangeServerConfigResponseHeader) GoString() string {
  return s.String()
}




type QueryDomainByOriginIPRequest struct {
}

func (s QueryDomainByOriginIPRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryDomainByOriginIPRequest) GoString() string {
  return s.String()
}

type QueryDomainByOriginIPResponse struct {
  // {"en":"httpstatus=202; Indicates that the new domain API was successfully invoked, and the current deployment of the new domain can be viewed using x-cnc-request-id in the header", "zh_CN":"httpstatus=202;   表示成功调用新增域名接口，可使用header中的x-cnc-request-id查看当前新增域名的部署情况"}
  HttpStatus *int `json:"http status code,omitempty" xml:"http status code,omitempty" require:"true"`
  // {"en":"Uniquely identified id for querying tasks per request (for all API)", "zh_CN":"唯一标示的id，用于查询每次请求的任务 （适用全部接口）"}
  XCncRequestId *string `json:"x-cnc-request-id,omitempty" xml:"x-cnc-request-id,omitempty" require:"true"`
  // {"en":"Code =200 indicates that relevant data was returned successfully", "zh_CN":"code=200，表示成功返回相关数据"}
  Code *int `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Query the source station IP", "zh_CN":"查询的源站IP"}
  Originip *string `json:"originip,omitempty" xml:"originip,omitempty" require:"true"`
  // {"en":"Returns a list of domain name names corresponding to each source station IP", "zh_CN":"返回各个源站IP对应的域名名称列表"}
  DomainList []*string `json:"domainList,omitempty" xml:"domainList,omitempty" require:"true" type:"Repeated"`
}

func (s QueryDomainByOriginIPResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryDomainByOriginIPResponse) GoString() string {
  return s.String()
}

func (s *QueryDomainByOriginIPResponse) SetHttpStatus(v int) *QueryDomainByOriginIPResponse {
  s.HttpStatus = &v
  return s
}

func (s *QueryDomainByOriginIPResponse) SetXCncRequestId(v string) *QueryDomainByOriginIPResponse {
  s.XCncRequestId = &v
  return s
}

func (s *QueryDomainByOriginIPResponse) SetCode(v int) *QueryDomainByOriginIPResponse {
  s.Code = &v
  return s
}

func (s *QueryDomainByOriginIPResponse) SetOriginip(v string) *QueryDomainByOriginIPResponse {
  s.Originip = &v
  return s
}

func (s *QueryDomainByOriginIPResponse) SetDomainList(v []*string) *QueryDomainByOriginIPResponse {
  s.DomainList = v
  return s
}

type QueryDomainByOriginIPPaths struct {
}

func (s QueryDomainByOriginIPPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryDomainByOriginIPPaths) GoString() string {
  return s.String()
}

type QueryDomainByOriginIPParameters struct {
  // {"en":"Source station IP, multiple IPs separated by semicolons", "zh_CN":"源站IP，多个IP用分号隔开
  // 注意：
  // 1、每次查询最多只能传入10个源站IP
  // 2、不支持源站域名的查询
  // 3、高级源匹配到对应IP时也能查到对应域名"}
  Originip *string `json:"originip,omitempty" xml:"originip,omitempty" require:"true"`
}

func (s QueryDomainByOriginIPParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryDomainByOriginIPParameters) GoString() string {
  return s.String()
}

func (s *QueryDomainByOriginIPParameters) SetOriginip(v string) *QueryDomainByOriginIPParameters {
  s.Originip = &v
  return s
}

type QueryDomainByOriginIPRequestHeader struct {
}

func (s QueryDomainByOriginIPRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryDomainByOriginIPRequestHeader) GoString() string {
  return s.String()
}

type QueryDomainByOriginIPResponseHeader struct {
}

func (s QueryDomainByOriginIPResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryDomainByOriginIPResponseHeader) GoString() string {
  return s.String()
}




type QueryCdnwAliasDomainsForWplusRequest struct {
}

func (s QueryCdnwAliasDomainsForWplusRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryCdnwAliasDomainsForWplusRequest) GoString() string {
  return s.String()
}

type QueryCdnwAliasDomainsForWplusResponse struct {
  // {"en":"Master domain Associated by Aliases.", "zh_CN":"关联的主域名"}
  DomainName *string `json:"domainName,omitempty" xml:"domainName,omitempty" require:"true"`
  // {"en":"Aiases, returning all aliases under the master domain.", "zh_CN":"别名，返回多个别名域名"}
  DomainAlias []*string `json:"domainAlias,omitempty" xml:"domainAlias,omitempty" require:"true" type:"Repeated"`
}

func (s QueryCdnwAliasDomainsForWplusResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryCdnwAliasDomainsForWplusResponse) GoString() string {
  return s.String()
}

func (s *QueryCdnwAliasDomainsForWplusResponse) SetDomainName(v string) *QueryCdnwAliasDomainsForWplusResponse {
  s.DomainName = &v
  return s
}

func (s *QueryCdnwAliasDomainsForWplusResponse) SetDomainAlias(v []*string) *QueryCdnwAliasDomainsForWplusResponse {
  s.DomainAlias = v
  return s
}

type QueryCdnwAliasDomainsForWplusPaths struct {
  // {"en":"Request result code", "zh_CN":"请求结果状态码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Request result information", "zh_CN":"请求结果信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Accelerated domain name or domain ID.
  // Query the master domain associated with the domain and all aliases of the master domain.", "zh_CN":"域名或者域名id:
  // 根据入参域名查询其主域名及主域名下的所有别名。"}
  DomainName *string `json:"domain-name,omitempty" xml:"domain-name,omitempty" require:"true"`
}

func (s QueryCdnwAliasDomainsForWplusPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryCdnwAliasDomainsForWplusPaths) GoString() string {
  return s.String()
}

func (s *QueryCdnwAliasDomainsForWplusPaths) SetCode(v string) *QueryCdnwAliasDomainsForWplusPaths {
  s.Code = &v
  return s
}

func (s *QueryCdnwAliasDomainsForWplusPaths) SetMessage(v string) *QueryCdnwAliasDomainsForWplusPaths {
  s.Message = &v
  return s
}

func (s *QueryCdnwAliasDomainsForWplusPaths) SetDomainName(v string) *QueryCdnwAliasDomainsForWplusPaths {
  s.DomainName = &v
  return s
}

type QueryCdnwAliasDomainsForWplusParameters struct {
}

func (s QueryCdnwAliasDomainsForWplusParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryCdnwAliasDomainsForWplusParameters) GoString() string {
  return s.String()
}

type QueryCdnwAliasDomainsForWplusRequestHeader struct {
}

func (s QueryCdnwAliasDomainsForWplusRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryCdnwAliasDomainsForWplusRequestHeader) GoString() string {
  return s.String()
}

type QueryCdnwAliasDomainsForWplusResponseHeader struct {
}

func (s QueryCdnwAliasDomainsForWplusResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryCdnwAliasDomainsForWplusResponseHeader) GoString() string {
  return s.String()
}




type EnableSingleDomainServiceRequest struct {
}

func (s EnableSingleDomainServiceRequest) String() string {
  return tea.Prettify(s)
}

func (s EnableSingleDomainServiceRequest) GoString() string {
  return s.String()
}

type EnableSingleDomainServiceResponse struct {
  // {"en":"Error code, which appears when HTTPStatus is not 202, represents the error type of the current request call", "zh_CN":"错误代码，当HTTPStatus不为202时出现，表示当前请求调用的错误类型"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response information, success when successful", "zh_CN":"响应信息，成功时为success"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"httpstatus=202; Indicates that the new domain API was successfully invoked, and the current deployment of the new domain can be viewed using x-cnc-request-id in the header", "zh_CN":"httpstatus=202;   表示成功调用新增域名接口，可使用header中的x-cnc-request-id查看当前新增域名的部署情况"}
  HttpStatus *int `json:"http status code,omitempty" xml:"http status code,omitempty" require:"true"`
  // {"en":"Uniquely identified id for querying tasks per request (for all API)", "zh_CN":"唯一标示的id，用于查询每次请求的任务 （适用全部接口）"}
  XCncRequestId *string `json:"x-cnc-request-id,omitempty" xml:"x-cnc-request-id,omitempty" require:"true"`
}

func (s EnableSingleDomainServiceResponse) String() string {
  return tea.Prettify(s)
}

func (s EnableSingleDomainServiceResponse) GoString() string {
  return s.String()
}

func (s *EnableSingleDomainServiceResponse) SetCode(v string) *EnableSingleDomainServiceResponse {
  s.Code = &v
  return s
}

func (s *EnableSingleDomainServiceResponse) SetMessage(v string) *EnableSingleDomainServiceResponse {
  s.Message = &v
  return s
}

func (s *EnableSingleDomainServiceResponse) SetHttpStatus(v int) *EnableSingleDomainServiceResponse {
  s.HttpStatus = &v
  return s
}

func (s *EnableSingleDomainServiceResponse) SetXCncRequestId(v string) *EnableSingleDomainServiceResponse {
  s.XCncRequestId = &v
  return s
}

type EnableSingleDomainServicePaths struct {
  // {"en":"Accelerate the ID of the domain name in the system
  // Note:
  // 1. See the url in the request example, 123344 for domain-id
  // 2. After the domain name is successfully submitted, the location access url in the return parameter can be queried to the domain-id of the domain name; You can also query domain-id through the Get domain Configuration and Get domain List interfaces", "zh_CN":"加速域名在系统中对应的ID
  // 注意：
  // 1. 参看请求示例中的url，123344对应的就是domain-id
  // 2. 可以通过【获取域名配置】和【获取域名列表】接口查询到domain-id"}
  DomainId *int `json:"domainId,omitempty" xml:"domainId,omitempty" require:"true"`
}

func (s EnableSingleDomainServicePaths) String() string {
  return tea.Prettify(s)
}

func (s EnableSingleDomainServicePaths) GoString() string {
  return s.String()
}

func (s *EnableSingleDomainServicePaths) SetDomainId(v int) *EnableSingleDomainServicePaths {
  s.DomainId = &v
  return s
}

type EnableSingleDomainServiceParameters struct {
}

func (s EnableSingleDomainServiceParameters) String() string {
  return tea.Prettify(s)
}

func (s EnableSingleDomainServiceParameters) GoString() string {
  return s.String()
}

type EnableSingleDomainServiceRequestHeader struct {
}

func (s EnableSingleDomainServiceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s EnableSingleDomainServiceRequestHeader) GoString() string {
  return s.String()
}

type EnableSingleDomainServiceResponseHeader struct {
}

func (s EnableSingleDomainServiceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s EnableSingleDomainServiceResponseHeader) GoString() string {
  return s.String()
}




type DeleteDomainAliasRequest struct {
  // {"en":"The domain name", "zh_CN":"域名名称"}
  DomainName *string `json:"domainName,omitempty" xml:"domainName,omitempty" require:"true"`
  // {"en":"List of domain alias to be deleted.Can contain multiple <aliasDomainName> nodes.", "zh_CN":"要删除的别名列表。可包含多个<aliasDomainName>节点。"}
  Aliases []*string `json:"aliases,omitempty" xml:"aliases,omitempty" require:"true" type:"Repeated"`
}

func (s DeleteDomainAliasRequest) String() string {
  return tea.Prettify(s)
}

func (s DeleteDomainAliasRequest) GoString() string {
  return s.String()
}

func (s *DeleteDomainAliasRequest) SetDomainName(v string) *DeleteDomainAliasRequest {
  s.DomainName = &v
  return s
}

func (s *DeleteDomainAliasRequest) SetAliases(v []*string) *DeleteDomainAliasRequest {
  s.Aliases = v
  return s
}

type DeleteDomainAliasResponse struct {
  // {"en":"Error code", "zh_CN":"错误码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"The message body", "zh_CN":"消息体"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Returns the body of the data. The <date> returned by this interface is always empty", "zh_CN":"返回数据体，此接口返回的<date>恒为空"}
  Data *string `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s DeleteDomainAliasResponse) String() string {
  return tea.Prettify(s)
}

func (s DeleteDomainAliasResponse) GoString() string {
  return s.String()
}

func (s *DeleteDomainAliasResponse) SetCode(v string) *DeleteDomainAliasResponse {
  s.Code = &v
  return s
}

func (s *DeleteDomainAliasResponse) SetMessage(v string) *DeleteDomainAliasResponse {
  s.Message = &v
  return s
}

func (s *DeleteDomainAliasResponse) SetData(v string) *DeleteDomainAliasResponse {
  s.Data = &v
  return s
}

type DeleteDomainAliasPaths struct {
}

func (s DeleteDomainAliasPaths) String() string {
  return tea.Prettify(s)
}

func (s DeleteDomainAliasPaths) GoString() string {
  return s.String()
}

type DeleteDomainAliasParameters struct {
}

func (s DeleteDomainAliasParameters) String() string {
  return tea.Prettify(s)
}

func (s DeleteDomainAliasParameters) GoString() string {
  return s.String()
}

type DeleteDomainAliasRequestHeader struct {
}

func (s DeleteDomainAliasRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteDomainAliasRequestHeader) GoString() string {
  return s.String()
}

type DeleteDomainAliasResponseHeader struct {
}

func (s DeleteDomainAliasResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteDomainAliasResponseHeader) GoString() string {
  return s.String()
}




type AddDomainAliasRequest struct {
  // {"en":"Domain Name, that Associated by domain-aliases.", "zh_CN":"域名，被别名关联的域名"}
  DomainName *string `json:"domain-name,omitempty" xml:"domain-name,omitempty" require:"true"`
  // {"en":"Domain name alias, support multiple.", "zh_CN":"域名别名，支持多个"}
  Aliases []*string `json:"aliases,omitempty" xml:"aliases,omitempty" require:"true" type:"Repeated"`
}

func (s AddDomainAliasRequest) String() string {
  return tea.Prettify(s)
}

func (s AddDomainAliasRequest) GoString() string {
  return s.String()
}

func (s *AddDomainAliasRequest) SetDomainName(v string) *AddDomainAliasRequest {
  s.DomainName = &v
  return s
}

func (s *AddDomainAliasRequest) SetAliases(v []*string) *AddDomainAliasRequest {
  s.Aliases = v
  return s
}

type AddDomainAliasResponse struct {
  // {"en":"httpstatus=202; Indicates that the new domain API was successfully invoked, and the current deployment of the new domain can be viewed using x-cnc-request-id in the header", "zh_CN":"httpstatus=202;   表示成功调用新增域名接口，可使用header中的x-cnc-request-id查看当前新增域名的部署情况"}
  HttpStatus *int `json:"http status code,omitempty" xml:"http status code,omitempty" require:"true"`
  // {"en":"Uniquely identified id for querying tasks per request (for all API)", "zh_CN":"唯一标示的id，用于查询每次请求的任务 （适用全部接口）"}
  XCncRequestId *string `json:"x-cnc-request-id,omitempty" xml:"x-cnc-request-id,omitempty" require:"true"`
  // {"en":"The error code, which appears when HTTPStatus is not 202, indicates the error type of the current request invocation", "zh_CN":"错误代码，当HTTPStatus不为202时出现，表示当前请求调用的错误类型"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response information, when success is successful", "zh_CN":"响应信息，成功时为success"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"Response data array.", "zh_CN":"接口响应数据"}
  Data *AddDomainAliasResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
}

func (s AddDomainAliasResponse) String() string {
  return tea.Prettify(s)
}

func (s AddDomainAliasResponse) GoString() string {
  return s.String()
}

func (s *AddDomainAliasResponse) SetHttpStatus(v int) *AddDomainAliasResponse {
  s.HttpStatus = &v
  return s
}

func (s *AddDomainAliasResponse) SetXCncRequestId(v string) *AddDomainAliasResponse {
  s.XCncRequestId = &v
  return s
}

func (s *AddDomainAliasResponse) SetCode(v string) *AddDomainAliasResponse {
  s.Code = &v
  return s
}

func (s *AddDomainAliasResponse) SetMsg(v string) *AddDomainAliasResponse {
  s.Msg = &v
  return s
}

func (s *AddDomainAliasResponse) SetData(v *AddDomainAliasResponseData) *AddDomainAliasResponse {
  s.Data = v
  return s
}

type AddDomainAliasResponseData struct {
  // {"en":"Domain name, return multiple alias domain names.", "zh_CN":"域名，返回多个别名域名"}
  Domains []*AddDomainAliasResponseDataDomains `json:"domains,omitempty" xml:"domains,omitempty" require:"true" type:"Repeated"`
}

func (s AddDomainAliasResponseData) String() string {
  return tea.Prettify(s)
}

func (s AddDomainAliasResponseData) GoString() string {
  return s.String()
}

func (s *AddDomainAliasResponseData) SetDomains(v []*AddDomainAliasResponseDataDomains) *AddDomainAliasResponseData {
  s.Domains = v
  return s
}

type AddDomainAliasResponseDataDomains struct     {
  // {"en":"Alias domain name.", "zh_CN":"别名名称"}
  AliasDomainName *string `json:"alias-domain-name,omitempty" xml:"alias-domain-name,omitempty" require:"true"`
}

func (s AddDomainAliasResponseDataDomains) String() string {
  return tea.Prettify(s)
}

func (s AddDomainAliasResponseDataDomains) GoString() string {
  return s.String()
}

func (s *AddDomainAliasResponseDataDomains) SetAliasDomainName(v string) *AddDomainAliasResponseDataDomains {
  s.AliasDomainName = &v
  return s
}

type AddDomainAliasPaths struct {
}

func (s AddDomainAliasPaths) String() string {
  return tea.Prettify(s)
}

func (s AddDomainAliasPaths) GoString() string {
  return s.String()
}

type AddDomainAliasParameters struct {
}

func (s AddDomainAliasParameters) String() string {
  return tea.Prettify(s)
}

func (s AddDomainAliasParameters) GoString() string {
  return s.String()
}

type AddDomainAliasRequestHeader struct {
}

func (s AddDomainAliasRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s AddDomainAliasRequestHeader) GoString() string {
  return s.String()
}

type AddDomainAliasResponseHeader struct {
}

func (s AddDomainAliasResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s AddDomainAliasResponseHeader) GoString() string {
  return s.String()
}




type DisableSingleDomainServiceRequest struct {
}

func (s DisableSingleDomainServiceRequest) String() string {
  return tea.Prettify(s)
}

func (s DisableSingleDomainServiceRequest) GoString() string {
  return s.String()
}

type DisableSingleDomainServiceResponse struct {
  // {"en":"Error code, which appears when HTTPStatus is not 202, represents the error type of the current request call", "zh_CN":"错误代码，当HTTPStatus不为202时出现，表示当前请求调用的错误类型"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response information, success when successful", "zh_CN":"响应信息，成功时为success"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"httpstatus=202; Indicates that the new domain API was successfully invoked, and the current deployment of the new domain can be viewed using x-cnc-request-id in the header", "zh_CN":"httpstatus=202;   表示成功调用新增域名接口，可使用header中的x-cnc-request-id查看当前新增域名的部署情况"}
  HttpStatus *int `json:"http status code,omitempty" xml:"http status code,omitempty" require:"true"`
  // {"en":"Uniquely identified id for querying tasks per request (for all API)", "zh_CN":"唯一标示的id，用于查询每次请求的任务 （适用全部接口）"}
  XCncRequestId *string `json:"x-cnc-request-id,omitempty" xml:"x-cnc-request-id,omitempty" require:"true"`
}

func (s DisableSingleDomainServiceResponse) String() string {
  return tea.Prettify(s)
}

func (s DisableSingleDomainServiceResponse) GoString() string {
  return s.String()
}

func (s *DisableSingleDomainServiceResponse) SetCode(v string) *DisableSingleDomainServiceResponse {
  s.Code = &v
  return s
}

func (s *DisableSingleDomainServiceResponse) SetMessage(v string) *DisableSingleDomainServiceResponse {
  s.Message = &v
  return s
}

func (s *DisableSingleDomainServiceResponse) SetHttpStatus(v int) *DisableSingleDomainServiceResponse {
  s.HttpStatus = &v
  return s
}

func (s *DisableSingleDomainServiceResponse) SetXCncRequestId(v string) *DisableSingleDomainServiceResponse {
  s.XCncRequestId = &v
  return s
}

type DisableSingleDomainServicePaths struct {
  // {"en":"", "zh_CN":"加速域名在系统中对应的ID
  // 1. 参看请求示例中的url，123344对应的就是domain-id
  // 2. 可以通过【获取域名配置】和【获取域名列表】接口查询到domain-id"}
  DomainId *int `json:"domainId,omitempty" xml:"domainId,omitempty" require:"true"`
}

func (s DisableSingleDomainServicePaths) String() string {
  return tea.Prettify(s)
}

func (s DisableSingleDomainServicePaths) GoString() string {
  return s.String()
}

func (s *DisableSingleDomainServicePaths) SetDomainId(v int) *DisableSingleDomainServicePaths {
  s.DomainId = &v
  return s
}

type DisableSingleDomainServiceParameters struct {
}

func (s DisableSingleDomainServiceParameters) String() string {
  return tea.Prettify(s)
}

func (s DisableSingleDomainServiceParameters) GoString() string {
  return s.String()
}

type DisableSingleDomainServiceRequestHeader struct {
}

func (s DisableSingleDomainServiceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s DisableSingleDomainServiceRequestHeader) GoString() string {
  return s.String()
}

type DisableSingleDomainServiceResponseHeader struct {
}

func (s DisableSingleDomainServiceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s DisableSingleDomainServiceResponseHeader) GoString() string {
  return s.String()
}




type QueryCdnDomainServiceRequest struct {
}

func (s QueryCdnDomainServiceRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryCdnDomainServiceRequest) GoString() string {
  return s.String()
}

type QueryCdnDomainServiceResponse struct {
  // {"en":"The deployment status of the accelerate domain name. Deployed indicates that the accelerated domain name configuration is complete. InProgress indicates that the deployment task of the accelerated domain name configuration is still in progress, and may be in queue, deployed, or failed.", "zh_CN":"加速域名的部署状态。部署表示加速域名配置已完成。InProgress表示加速域名配置的部署任务仍在进行中，可能正在排队、部署或失败。"}
  Status *string `json:"status,omitempty" xml:"status,omitempty" require:"true"`
  // {"en":"Ssl certificate settings, used to set the ssl certificate configuration for the accelerated domain name", "zh_CN":"Ssl证书设置，用于设置加速域名的Ssl证书配置"}
  Ssl *QueryCdnDomainServiceResponseSsl `json:"ssl,omitempty" xml:"ssl,omitempty" require:"true" type:"Struct"`
  // {"en":"Accelerated domain name service types, including the following: 1028 : Content Acceleration; 1115 : Dynamic Web Acceleration; 1369 : Media Acceleration - RTMP 1391 : Download Acceleration 1348 : Media Acceleration Live Broadcast 1551 : Flood Shield", "zh_CN":"加速的域名服务类型，包括： 1028：内容加速； 1115：动态Web加速； 1369：媒体加速； 1391：下载加速； 1348：媒体加速直播 1551 : Flood Shield"}
  ServiceType *string `json:"service-type,omitempty" xml:"service-type,omitempty" require:"true"`
  // {"en":"back-to-origin policy setting, used to set the source station information and back-to-origin policy of the accelerated domain name", "zh_CN":"回源策略设置，用于设置加速域名的源站信息和回源策略"}
  OriginConfig *QueryCdnDomainServiceResponseOriginConfig `json:"origin-config,omitempty" xml:"origin-config,omitempty" require:"true" type:"Struct"`
  // {"en":"Speed up the activation of the domain name. This is false when the accelerated domain name service is disabled; true when the accelerated domain name service is enabled.", "zh_CN":"加快域名的状态。 如果禁用了加速域名服务，则为false；否则为true。"}
  Enabled *string `json:"enabled,omitempty" xml:"enabled,omitempty" require:"true"`
  // {"en":"Accelerated domain name", "zh_CN":"加速域名名称"}
  DomainName *string `json:"domain-name,omitempty" xml:"domain-name,omitempty" require:"true"`
  // {"en":"Accelerated domain ID returned by the system", "zh_CN":"系统返回的加速域名ID"}
  DomainId *int `json:"domain-id,omitempty" xml:"domain-id,omitempty" require:"true"`
  // {"en":"Remarks, up to 1000 characters", "zh_CN":"备注，最多1000个字符"}
  Comment *string `json:"comment,omitempty" xml:"comment,omitempty" require:"true"`
  // {"en":"Accelerate the CDN service status of the domain name, true means to enable the CDN acceleration service; false means to cancel the CDN acceleration service.", "zh_CN":"加速域名的CDN服务状态，true表示启用CDN加速服务； false表示取消CDN加速服务。"}
  CdnServiceStatus *string `json:"cdn-service-status,omitempty" xml:"cdn-service-status,omitempty" require:"true"`
  // {"en":"It is cache policy settings for setting cache rules of accelerated domain names Note: The default is that the higher ranked rules have higher priority.", "zh_CN":"它是用于设置加速域名的缓存规则的缓存策略设置 注意：默认情况是排名较高的规则具有更高的优先级。"}
  CacheBehavior []*QueryCdnDomainServiceResponseCacheBehavior `json:"cache-behavior,omitempty" xml:"cache-behavior,omitempty" require:"true" type:"Repeated"`
  // {"en":"Pass the response header of client IP.", "zh_CN":"传递客户端IP的响应头。"}
  HeaderOfClientip *string `json:"header-of-clientip,omitempty" xml:"header-of-clientip,omitempty" require:"true"`
  // {"en":"The live push-pull stream type, the optional values are pull and push, pull means pull flow; push means push flow.", "zh_CN":"直播推拉流类型，可选值为pull和push，pull表示拉流；   push表示推流。"}
  DomainStreamType *string `json:"domain-stream-type,omitempty" xml:"domain-stream-type,omitempty" require:"true"`
  // {"en":"contract-id", "zh_CN":"合同号"}
  ContractId *string `json:"contract-id,omitempty" xml:"contract-id,omitempty" require:"true"`
  // {"en":"item-id", "zh_CN":"产品号"}
  ItemId *string `json:"item-id,omitempty" xml:"item-id,omitempty" require:"true"`
  // {"en":"CNAME of accelerate domain,for example:7nt6mrh7sdkslj.cdn30.com", "zh_CN":"加速域名对应的CNAME域名，例如：7nt6mrh7sdkslj.cdn30.com"}
  Cname *string `json:"cname,omitempty" xml:"cname,omitempty" require:"true"`
  // {"en":"The WAF protection status, true or false.", "zh_CN":"是否开始WAF防护"}
  WafEnable *string `json:"waf-enable,omitempty" xml:"waf-enable,omitempty" require:"true"`
  // {"en":"Cache file HOST (use required).", "zh_CN":"缓存文件HOST（使用需申请） 缓存HOST域名和加速域名的”缓存规则”必须一致 注意：该节点下的相关参数配置，除开通API调用权限外，还需要联系专属客服申请开通对应的API客户模板"}
  CacheHost *string `json:"cache-host,omitempty" xml:"cache-host,omitempty" require:"true"`
}

func (s QueryCdnDomainServiceResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryCdnDomainServiceResponse) GoString() string {
  return s.String()
}

func (s *QueryCdnDomainServiceResponse) SetStatus(v string) *QueryCdnDomainServiceResponse {
  s.Status = &v
  return s
}

func (s *QueryCdnDomainServiceResponse) SetSsl(v *QueryCdnDomainServiceResponseSsl) *QueryCdnDomainServiceResponse {
  s.Ssl = v
  return s
}

func (s *QueryCdnDomainServiceResponse) SetServiceType(v string) *QueryCdnDomainServiceResponse {
  s.ServiceType = &v
  return s
}

func (s *QueryCdnDomainServiceResponse) SetOriginConfig(v *QueryCdnDomainServiceResponseOriginConfig) *QueryCdnDomainServiceResponse {
  s.OriginConfig = v
  return s
}

func (s *QueryCdnDomainServiceResponse) SetEnabled(v string) *QueryCdnDomainServiceResponse {
  s.Enabled = &v
  return s
}

func (s *QueryCdnDomainServiceResponse) SetDomainName(v string) *QueryCdnDomainServiceResponse {
  s.DomainName = &v
  return s
}

func (s *QueryCdnDomainServiceResponse) SetDomainId(v int) *QueryCdnDomainServiceResponse {
  s.DomainId = &v
  return s
}

func (s *QueryCdnDomainServiceResponse) SetComment(v string) *QueryCdnDomainServiceResponse {
  s.Comment = &v
  return s
}

func (s *QueryCdnDomainServiceResponse) SetCdnServiceStatus(v string) *QueryCdnDomainServiceResponse {
  s.CdnServiceStatus = &v
  return s
}

func (s *QueryCdnDomainServiceResponse) SetCacheBehavior(v []*QueryCdnDomainServiceResponseCacheBehavior) *QueryCdnDomainServiceResponse {
  s.CacheBehavior = v
  return s
}

func (s *QueryCdnDomainServiceResponse) SetHeaderOfClientip(v string) *QueryCdnDomainServiceResponse {
  s.HeaderOfClientip = &v
  return s
}

func (s *QueryCdnDomainServiceResponse) SetDomainStreamType(v string) *QueryCdnDomainServiceResponse {
  s.DomainStreamType = &v
  return s
}

func (s *QueryCdnDomainServiceResponse) SetContractId(v string) *QueryCdnDomainServiceResponse {
  s.ContractId = &v
  return s
}

func (s *QueryCdnDomainServiceResponse) SetItemId(v string) *QueryCdnDomainServiceResponse {
  s.ItemId = &v
  return s
}

func (s *QueryCdnDomainServiceResponse) SetCname(v string) *QueryCdnDomainServiceResponse {
  s.Cname = &v
  return s
}

func (s *QueryCdnDomainServiceResponse) SetWafEnable(v string) *QueryCdnDomainServiceResponse {
  s.WafEnable = &v
  return s
}

func (s *QueryCdnDomainServiceResponse) SetCacheHost(v string) *QueryCdnDomainServiceResponse {
  s.CacheHost = &v
  return s
}

type QueryCdnDomainServiceResponseSsl struct {
  // {"en":"Use a certificate, the optional values are true and false, true means using the certificate, false means not using the certificate", "zh_CN":"使用证书，可选值为true和false，true表示使用证书，false表示不使用证书"}
  UseSsl *string `json:"use-ssl,omitempty" xml:"use-ssl,omitempty" require:"true"`
  // {"en":"Use sni certificate, the optional values are true and false, true means use sni certificate, false means use non-sni certificate", "zh_CN":"使用sni证书，可选值为true和false，true表示使用sni证书，false表示使用非sni证书"}
  UseForSni *string `json:"use-for-sni,omitempty" xml:"use-for-sni,omitempty" require:"true"`
  // {"en":"Certificate ID. The certificate ID returned by the system after the new certificate is successfully added.", "zh_CN":"证书ID。成功添加新证书后系统返回的证书ID。"}
  SslCertificateId *int `json:"ssl-certificate-id,omitempty" xml:"ssl-certificate-id,omitempty" require:"true"`
}

func (s QueryCdnDomainServiceResponseSsl) String() string {
  return tea.Prettify(s)
}

func (s QueryCdnDomainServiceResponseSsl) GoString() string {
  return s.String()
}

func (s *QueryCdnDomainServiceResponseSsl) SetUseSsl(v string) *QueryCdnDomainServiceResponseSsl {
  s.UseSsl = &v
  return s
}

func (s *QueryCdnDomainServiceResponseSsl) SetUseForSni(v string) *QueryCdnDomainServiceResponseSsl {
  s.UseForSni = &v
  return s
}

func (s *QueryCdnDomainServiceResponseSsl) SetSslCertificateId(v int) *QueryCdnDomainServiceResponseSsl {
  s.SslCertificateId = &v
  return s
}

type QueryCdnDomainServiceResponseOriginConfig struct {
  // {"en":"Return source address, which can be IP or domain name.", "zh_CN":"回源地址，可以是IP或域名。"}
  OriginIps *string `json:"origin-ips,omitempty" xml:"origin-ips,omitempty" require:"true"`
  // {"en":"Back-to-origin HOST is used to change the HOST field in the source HTTP request header.", "zh_CN":"Back-to-origin HOST用于更改源HTTP请求标头中的HOST字段。"}
  DefaultOriginHostHeader *string `json:"default-origin-host-header,omitempty" xml:"default-origin-host-header,omitempty" require:"true"`
  AdvOriginConfig []*QueryCdnDomainServiceResponseOriginConfigAdvOriginConfig `json:"advOriginConfig,omitempty" xml:"advOriginConfig,omitempty" require:"true" type:"Repeated"`
}

func (s QueryCdnDomainServiceResponseOriginConfig) String() string {
  return tea.Prettify(s)
}

func (s QueryCdnDomainServiceResponseOriginConfig) GoString() string {
  return s.String()
}

func (s *QueryCdnDomainServiceResponseOriginConfig) SetOriginIps(v string) *QueryCdnDomainServiceResponseOriginConfig {
  s.OriginIps = &v
  return s
}

func (s *QueryCdnDomainServiceResponseOriginConfig) SetDefaultOriginHostHeader(v string) *QueryCdnDomainServiceResponseOriginConfig {
  s.DefaultOriginHostHeader = &v
  return s
}

func (s *QueryCdnDomainServiceResponseOriginConfig) SetAdvOriginConfig(v []*QueryCdnDomainServiceResponseOriginConfigAdvOriginConfig) *QueryCdnDomainServiceResponseOriginConfig {
  s.AdvOriginConfig = v
  return s
}

type QueryCdnDomainServiceResponseOriginConfigAdvOriginConfig struct     {
  // {"en":"The advanced source monitors the url, and requests <master-ips> through the url. If the response is neither 2xx nor 3xx response, it is considered that the primary source ip is faulty, and <backup-ips> is used at this time. The full url, for example: http://a.example.com/test.html", "zh_CN":"高级源监视URL，并通过URL请求<master-ips>。 如果响应既不是2xx也不是3xx响应，则认为主要源ip错误，并且此时使用<backup-ips>。 例如：http://a.example.com/test.html"}
  DetectUrl *string `json:"detect-url,omitempty" xml:"detect-url,omitempty" require:"true"`
  // {"en":"Advanced source monitoring period, and the unit is in seconds, optional as an integer greater than or equal to 0, 0 means no monitoring", "zh_CN":"高级源监视周期，单位为秒，可选为大于或等于0的整数，0表示不监视"}
  DetectPeriod *int `json:"detect-period,omitempty" xml:"detect-period,omitempty" require:"true"`
  AdvOriginConfig *QueryCdnDomainServiceResponseOriginConfigAdvOriginConfigAdvOriginConfig `json:"advOriginConfig,omitempty" xml:"advOriginConfig,omitempty" require:"true" type:"Struct"`
}

func (s QueryCdnDomainServiceResponseOriginConfigAdvOriginConfig) String() string {
  return tea.Prettify(s)
}

func (s QueryCdnDomainServiceResponseOriginConfigAdvOriginConfig) GoString() string {
  return s.String()
}

func (s *QueryCdnDomainServiceResponseOriginConfigAdvOriginConfig) SetDetectUrl(v string) *QueryCdnDomainServiceResponseOriginConfigAdvOriginConfig {
  s.DetectUrl = &v
  return s
}

func (s *QueryCdnDomainServiceResponseOriginConfigAdvOriginConfig) SetDetectPeriod(v int) *QueryCdnDomainServiceResponseOriginConfigAdvOriginConfig {
  s.DetectPeriod = &v
  return s
}

func (s *QueryCdnDomainServiceResponseOriginConfigAdvOriginConfig) SetAdvOriginConfig(v *QueryCdnDomainServiceResponseOriginConfigAdvOriginConfigAdvOriginConfig) *QueryCdnDomainServiceResponseOriginConfigAdvOriginConfig {
  s.AdvOriginConfig = v
  return s
}

type QueryCdnDomainServiceResponseOriginConfigAdvOriginConfigAdvOriginConfig struct {
  // {"en":"The advanced source mainly returns the source IP. Multiple IPs are separated by a semicolon \";\", and the return source IP cannot be repeated.", "zh_CN":"高级源返回源IP。 多个IP以分号“;”分隔，并且返回源IP无法重复。"}
  MasterIps *string `json:"master-ips,omitempty" xml:"master-ips,omitempty" require:"true"`
  // {"en":"Advanced source backup source IP, multiple IPs are separated by semicolon \";\", and the return source IP cannot be duplicated.", "zh_CN":"高级源备份源IP，多个IP用分号“;”分隔，并且返回源IP无法重复。"}
  BackupIps *string `json:"backup-ips,omitempty" xml:"backup-ips,omitempty" require:"true"`
}

func (s QueryCdnDomainServiceResponseOriginConfigAdvOriginConfigAdvOriginConfig) String() string {
  return tea.Prettify(s)
}

func (s QueryCdnDomainServiceResponseOriginConfigAdvOriginConfigAdvOriginConfig) GoString() string {
  return s.String()
}

func (s *QueryCdnDomainServiceResponseOriginConfigAdvOriginConfigAdvOriginConfig) SetMasterIps(v string) *QueryCdnDomainServiceResponseOriginConfigAdvOriginConfigAdvOriginConfig {
  s.MasterIps = &v
  return s
}

func (s *QueryCdnDomainServiceResponseOriginConfigAdvOriginConfigAdvOriginConfig) SetBackupIps(v string) *QueryCdnDomainServiceResponseOriginConfigAdvOriginConfigAdvOriginConfig {
  s.BackupIps = &v
  return s
}

type QueryCdnDomainServiceResponseCacheBehavior struct     {
  // {"en":"The url matching mode supports fuzzy regularization. The optional values are: /(a|b)/*.(jpg|bmp|png|gif) and other regular contents. The fuzzy regularization description is as follows: /*: Match all files /*.jpg: All jpg files also contain all jpg files in subdirectories such as /xx/. /*.(jpg|gif): All jpg or gif files, including all jpg or gif files in subdirectories such as /xx/. /a/*: All files in directory a, including all files in subdirectories such as /a/xx/. /(a|b)/*.jpg: All jpg files in directory a or directory b also contain all jpg files under /(a/b)/xx/.", "zh_CN":"网址匹配模式支持模糊正则化。 可选值是：/(a|b)/*.(jpg|bmp|png|gif）和其他常规内容。 模糊正则化描述如下： / *：匹配所有文件 /*.jpg：所有jpg文件还包含子目录（例如/ xx /）中的所有jpg文件。 /*.(jpg|gif）：所有jpg或gif文件，包括/ xx /等子目录中的所有jpg或gif文件。 / a / *：目录a中的所有文件，包括/ a / xx /等子目录中的所有文件。 /(a|b)/*.jpg：目录a或目录b中的所有jpg文件也包含/（a / b）/ xx /下的所有jpg文件。"}
  PathPattern *string `json:"path-pattern,omitempty" xml:"path-pattern,omitempty" require:"true"`
  // {"en":"Cache time, and its unit is in seconds. There is no upper limit on the cache time rule. This time is set attuned to the customer is own needs. If the customer feels that some of the files do not change frequently, then the setting is longer. For example, the text class js, css, html, etc. can be set shorter, the picture, video and audio classes can be set longer (because file will be replaced by the new file due to the file heat algorithm, the suggestion that the length should not exceed one month)", "zh_CN":"缓存时间，单位为秒。 缓存时间规则没有上限。 此时间设置为适合客户自己的需求。 如果客户认为某些文件不经常更改，则设置会更长。 例如，可以将文本类js，css，html等设置为较短，将图片，视频和音频类设置为较长（由于文件加热算法，文件将被新文件替换，因此建议 长度不得超过一个月）"}
  CacheTtl *string `json:"cache-ttl,omitempty" xml:"cache-ttl,omitempty" require:"true"`
  // {"en":"Ignore that the source station does not cache the header. The optional values are true and false, which are used to ignore the two configurations of cache-control in the request header (private, no-cache) and the authorization set by the client. The true indicates that the source station is settings for the three are ignored. Enables resources to be cached on the service node in the form of cache-control: public, and then our nodes can cache this type of resource and provide acceleration services. False means that when the source station sets cache-control: private, cache-control: no-cache for a resource or specifies to cache according to authorization, our service node will not cache such files.", "zh_CN":"忽略源站不缓存头。 可选值true和false，用于忽略请求标头中的缓存控制的两种配置（专用，无缓存）和客户端设置的授权。 true表示忽略源站的三个设置。 使资源能够以cache-control：public的形式缓存在服务节点上，然后我们的节点可以缓存这种类型的资源并提供加速服务。 False表示当源站点为资源设置缓存控制：私有，缓存控制：无缓存或指定根据授权进行缓存时，我们的服务节点将不会缓存此类文件。"}
  IgnoreCacheControl *string `json:"ignore-cache-control,omitempty" xml:"ignore-cache-control,omitempty" require:"true"`
}

func (s QueryCdnDomainServiceResponseCacheBehavior) String() string {
  return tea.Prettify(s)
}

func (s QueryCdnDomainServiceResponseCacheBehavior) GoString() string {
  return s.String()
}

func (s *QueryCdnDomainServiceResponseCacheBehavior) SetPathPattern(v string) *QueryCdnDomainServiceResponseCacheBehavior {
  s.PathPattern = &v
  return s
}

func (s *QueryCdnDomainServiceResponseCacheBehavior) SetCacheTtl(v string) *QueryCdnDomainServiceResponseCacheBehavior {
  s.CacheTtl = &v
  return s
}

func (s *QueryCdnDomainServiceResponseCacheBehavior) SetIgnoreCacheControl(v string) *QueryCdnDomainServiceResponseCacheBehavior {
  s.IgnoreCacheControl = &v
  return s
}

type QueryCdnDomainServicePaths struct {
  // {"en":"Accelerated domain name or domain ID", "zh_CN":"域名名称或域名id，在请求的url后面"}
  DomainName *string `json:"domain-name,omitempty" xml:"domain-name,omitempty" require:"true"`
}

func (s QueryCdnDomainServicePaths) String() string {
  return tea.Prettify(s)
}

func (s QueryCdnDomainServicePaths) GoString() string {
  return s.String()
}

func (s *QueryCdnDomainServicePaths) SetDomainName(v string) *QueryCdnDomainServicePaths {
  s.DomainName = &v
  return s
}

type QueryCdnDomainServiceParameters struct {
}

func (s QueryCdnDomainServiceParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryCdnDomainServiceParameters) GoString() string {
  return s.String()
}

type QueryCdnDomainServiceRequestHeader struct {
}

func (s QueryCdnDomainServiceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryCdnDomainServiceRequestHeader) GoString() string {
  return s.String()
}

type QueryCdnDomainServiceResponseHeader struct {
}

func (s QueryCdnDomainServiceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryCdnDomainServiceResponseHeader) GoString() string {
  return s.String()
}




type DeleteApiDomainServiceRequest struct {
  // {"en":"Accelerated domain name, choose one from domain-id. Accelerate the ID of the domain name in the system
  // Note:
  // 1. See the url in the request example, 123344 for domain-id
  // 2、After the domain name is successfully submitted, the location access url in the return parameter can be queried to the domain-id of the domain name; You can also query domain-id through the Get domain Configuration and Get domain List interfaces", "zh_CN":"加速域名与domain-id二选一。
  // domain-id：加速域名在系统中对应的ID
  // domain-name：加速的域名
  // 注意：
  // 1、参看请求示例中的url，123344对应的就是domain-id
  // 2、创建域名成功提交后，返回参数中的location访问url中，能够查询到域名对应的domain-id；也可以通过【获取域名配置】和【获取域名列表】接口查询到domain-id"}
  DomainName *string `json:"domainName,omitempty" xml:"domainName,omitempty" require:"true"`
  // {"en":"Accelerated domain name, choose one from domain-id. Accelerate the ID of the domain name in the system
  // Note:
  // 1. See the url in the request example, 123344 for domain-id
  // 2、After the domain name is successfully submitted, the location access url in the return parameter can be queried to the domain-id of the domain name; You can also query domain-id through the Get domain Configuration and Get domain List interfaces", "zh_CN":"加速域名与domain-id二选一。
  // domain-id：加速域名在系统中对应的ID
  // domain-name：加速的域名
  // 注意：
  // 1、参看请求示例中的url，123344对应的就是domain-id
  // 2、创建域名成功提交后，返回参数中的location访问url中，能够查询到域名对应的domain-id；也可以通过【获取域名配置】和【获取域名列表】接口查询到domain-id"}
  DomainId *string `json:"domainId,omitempty" xml:"domainId,omitempty" require:"true"`
}

func (s DeleteApiDomainServiceRequest) String() string {
  return tea.Prettify(s)
}

func (s DeleteApiDomainServiceRequest) GoString() string {
  return s.String()
}

func (s *DeleteApiDomainServiceRequest) SetDomainName(v string) *DeleteApiDomainServiceRequest {
  s.DomainName = &v
  return s
}

func (s *DeleteApiDomainServiceRequest) SetDomainId(v string) *DeleteApiDomainServiceRequest {
  s.DomainId = &v
  return s
}

type DeleteApiDomainServiceResponse struct {
  // {"en":"Error code, which appears when HTTPStatus is not 202, represents the error type of the current request call", "zh_CN":"错误代码，当HTTPStatus不为202时出现，表示当前请求调用的错误类型"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response information, success when successful", "zh_CN":"响应信息，成功时为success"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"httpstatus=202; Indicates that the new domain API was successfully invoked, and the current deployment of the new domain can be viewed using x-cnc-request-id in the header", "zh_CN":"httpstatus=202;   表示成功调用新增域名接口，可使用header中的x-cnc-request-id查看当前新增域名的部署情况"}
  HttpStatus *int `json:"http status code,omitempty" xml:"http status code,omitempty" require:"true"`
  // {"en":"Uniquely identified id for querying tasks per request (for all API)", "zh_CN":"唯一标示的id，用于查询每次请求的任务 （适用全部接口）"}
  XCncRequestId *string `json:"x-cnc-request-id,omitempty" xml:"x-cnc-request-id,omitempty" require:"true"`
}

func (s DeleteApiDomainServiceResponse) String() string {
  return tea.Prettify(s)
}

func (s DeleteApiDomainServiceResponse) GoString() string {
  return s.String()
}

func (s *DeleteApiDomainServiceResponse) SetCode(v string) *DeleteApiDomainServiceResponse {
  s.Code = &v
  return s
}

func (s *DeleteApiDomainServiceResponse) SetMessage(v string) *DeleteApiDomainServiceResponse {
  s.Message = &v
  return s
}

func (s *DeleteApiDomainServiceResponse) SetHttpStatus(v int) *DeleteApiDomainServiceResponse {
  s.HttpStatus = &v
  return s
}

func (s *DeleteApiDomainServiceResponse) SetXCncRequestId(v string) *DeleteApiDomainServiceResponse {
  s.XCncRequestId = &v
  return s
}

type DeleteApiDomainServicePaths struct {
  // {"en":"", "zh_CN":"域名名称或域名id，在请求的url后面"}
  DomainName *string `json:"domain-name,omitempty" xml:"domain-name,omitempty" require:"true"`
}

func (s DeleteApiDomainServicePaths) String() string {
  return tea.Prettify(s)
}

func (s DeleteApiDomainServicePaths) GoString() string {
  return s.String()
}

func (s *DeleteApiDomainServicePaths) SetDomainName(v string) *DeleteApiDomainServicePaths {
  s.DomainName = &v
  return s
}

type DeleteApiDomainServiceParameters struct {
}

func (s DeleteApiDomainServiceParameters) String() string {
  return tea.Prettify(s)
}

func (s DeleteApiDomainServiceParameters) GoString() string {
  return s.String()
}

type DeleteApiDomainServiceRequestHeader struct {
}

func (s DeleteApiDomainServiceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteApiDomainServiceRequestHeader) GoString() string {
  return s.String()
}

type DeleteApiDomainServiceResponseHeader struct {
}

func (s DeleteApiDomainServiceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteApiDomainServiceResponseHeader) GoString() string {
  return s.String()
}




type AddCdnDomainRequest struct {
  // {"en":"The current version is 1.0.0", "zh_CN":"版本号，当前版本号1.0.0"}
  Version *string `json:"version,omitempty" xml:"version,omitempty" require:"true"`
  // {"en":"CDN accelerated domain name. Support generic domain name, witch starts with a dot, such as: .example.com, If the domain example.com have China ICP, then the domain name xx.example.com have ICP too.", "zh_CN":"需要接入CDN的域名。支持泛域名，以符号“.”开头，如：.example.com，
  //   泛域名也包含多级 a.b.example.com。 如果example.com已备案，那么域名xx.example.com则不需要备案。"}
  DomainName *string `json:"domain-name,omitempty" xml:"domain-name,omitempty" require:"true"`
  // {"en":"The value should be an effective domain. The added domain will copy the config of this domain.", 
  //   "zh_CN":"指定参考域名。当使用参考域名来新增域名是，新域名将复用参考域名的配置"}
  ReferencedDomainName *string `json:"referenced-domain-name,omitempty" xml:"referenced-domain-name,omitempty"`
  // {"en":"Configure single template. If you want to add a new accelerated domain name with a specified configuration, you can use a specified configured template. For details, please consult the corresponding customer support staff.", "zh_CN":"配置单模板，特定的使用场景下，如果希望新增的加速域名参照某些指定配置时，可以指定配置单模板，具体使用请咨询对应的客户负责人。"}
  ConfigFormId *string `json:"config-form-id,omitempty" xml:"config-form-id,omitempty"`
  // {"en":"The optional value is: true or false.", "zh_CN":"是否纯海外加速，入参范围：true、false"}
  AccelerateNoChina *string `json:"accelerate-no-china,omitempty" xml:"accelerate-no-china,omitempty"`
  // {"en":"The id of contract, such as 40015677", "zh_CN":"合同号，如40015677"}
  ContractId *string `json:"contract-id,omitempty" xml:"contract-id,omitempty" require:"true"`
  // {"en":"The id of product, such as 10", "zh_CN":"产品号，如10"}
  ItemId *string `json:"item-id,omitempty" xml:"item-id,omitempty" require:"true"`
  // {"en":"An alias for public service cname. When you have multiple domains that need to share a cname, then you may specify a cname-label here. If domains use the same cname-label, they will share the same cname and has the same dns coverage. Note: 1. Restrictions on domains with the same cname-label: they should have the same acceleration type,  certificate (if use ssl), service areas. 2. Multiple http domains can share the same cname; multiple sni https domains can share the same cname as wll. 3. When a single domain uses a cname-label, it can be cancelled acceleration; while multiple domains are not allowed to cancel acceleration for part of them.", "zh_CN":"共用一级别名，客户存在较多一级域名共用的需求，因此在接口中引入cname-label标识，即拥有相同cname-label的一组域名，共用一级cname。当加速域名的加速区域、加速类型、资源链是一致的时候，建议使用共用一级别名，相同的业务使用相同一级别名cname。 注意： 1、拥有相同cname-label的域名共用一级cname，且有完全一致的dns覆盖 2、共用一级的约束：配置一致、加速类型一致、证书id一致（如果有证书）、加速区域一致、平台套餐一致 3、多个http域名可共用一级，多个sni https域名可共用一级，多个共享证书域名可共用一级 4、单个域名使用cname-label时，域名可取消加速；多个域名共用一级时，不允许取消加速这些域名 5、支持通过修改cname-label达到修改cname的目的"}
  CnameLabel *string `json:"cname-label,omitempty" xml:"cname-label,omitempty"`
  // {"en":"Remarks, up to 1000 characters", "zh_CN":"备注信息，最大限制1000个字符"}
  Comment *string `json:"comment,omitempty" xml:"comment,omitempty"`
  // {"en":"Pass the response header of client IP. The optional values are Cdn-Src-Ip and X-Forwarded-For. The default value is Cdn-Src-Ip.", "zh_CN":"传递客户端ip的响应头部，可选值为Cdn-Src-Ip和X-Forwarded-For，默认值为Cdn-Src-Ip"}
  HeaderOfClientIp *string `json:"header-of-clientip,omitempty" xml:"header-of-clientip,omitempty"`
  // {"en":"Back-to-origin policy setting, which is used to set the origin site information and the back-to-origin policy of the accelerated domain name", "zh_CN":"回源策略设置，用于设置加速域名的源站信息和回源策略"}
  OriginConfig *AddCdnDomainRequestOriginConfig `json:"origin-config,omitempty" xml:"origin-config,omitempty" require:"true" type:"Struct"`
  // {"en":"Live domain configuration, used to set the livestream acceleration domain origin config.
  // Note: In addition to the API call permission, you need to contact the dedicated customer service to apply for the corresponding API client template.", "zh_CN":"直播域名配置，用于设置直播加速域名的推拉流（使用需申请）
  // 注意：该节点下的相关参数配置，除开通API调用权限外，还需要联系专属客服申请开通对应的API客户模板"}
  LiveConfig *AddCdnDomainRequestLiveConfig `json:"live-config,omitempty" xml:"live-config,omitempty" type:"Struct"`
  // {"en":"Livestream domain settings. Set the publishing point of the live push-pull domain.
  // note:
  // 1. The pull stream and the corresponding push stream domain must be configured with the same publishing point.
  // 2. If you are not going to modify the publishing point, please do not pass this param.
  // 3. The publishing point adopts the overlay update. Each time you modify, you need to submit all the publishing points. You cannot submit only the parts that need to be modified.", "zh_CN":"设置直播推拉流域名的发布点
  // 注意：
  // 1、拉流和对应的推流域名，必须配置相同的发布点；
  // 2、不想修改发布点时，不要传入该节点及以下入参；
  // 3、发布点采用覆盖式更新，每次修改时，需要提交全部发布点，不能仅提交需要修改的部分。"}
  PublishPoints []*AddCdnDomainRequestPublishPoints `json:"publish-points,omitempty" xml:"publish-points,omitempty" type:"Repeated"`
  // {"en":"SSL settings, to bind a certificate with the accelerated domain. You can use the interface [AddCertificate] to upload your  certificates. If you want to modify a certificate, please use the interface: [UpdateCertificate]", "zh_CN":"ssl证书设置，用于设置加速域名的ssl证书配置。上传证书请使用接口：【新增证书V2】；若要修改证书，请使用接口：【修改证书V2】"}
  Ssl *AddCdnDomainRequestSsl `json:"ssl,omitempty" xml:"ssl,omitempty" type:"Struct"`
}

func (s AddCdnDomainRequest) String() string {
  return tea.Prettify(s)
}

func (s AddCdnDomainRequest) GoString() string {
  return s.String()
}

func (s *AddCdnDomainRequest) SetVersion(v string) *AddCdnDomainRequest {
  s.Version = &v
  return s
}

func (s *AddCdnDomainRequest) SetDomainName(v string) *AddCdnDomainRequest {
  s.DomainName = &v
  return s
}

func (s *AddCdnDomainRequest) SetReferencedDomainName(v string) *AddCdnDomainRequest {
  s.ReferencedDomainName = &v
  return s
}

func (s *AddCdnDomainRequest) SetConfigFormId(v string) *AddCdnDomainRequest {
  s.ConfigFormId = &v
  return s
}

func (s *AddCdnDomainRequest) SetAccelerateNoChina(v string) *AddCdnDomainRequest {
  s.AccelerateNoChina = &v
  return s
}

func (s *AddCdnDomainRequest) SetContractId(v string) *AddCdnDomainRequest {
  s.ContractId = &v
  return s
}

func (s *AddCdnDomainRequest) SetItemId(v string) *AddCdnDomainRequest {
  s.ItemId = &v
  return s
}

func (s *AddCdnDomainRequest) SetCnameLabel(v string) *AddCdnDomainRequest {
  s.CnameLabel = &v
  return s
}

func (s *AddCdnDomainRequest) SetComment(v string) *AddCdnDomainRequest {
  s.Comment = &v
  return s
}

func (s *AddCdnDomainRequest) SetHeaderOfClientIp(v string) *AddCdnDomainRequest {
  s.HeaderOfClientIp = &v
  return s
}

func (s *AddCdnDomainRequest) SetOriginConfig(v *AddCdnDomainRequestOriginConfig) *AddCdnDomainRequest {
  s.OriginConfig = v
  return s
}

func (s *AddCdnDomainRequest) SetLiveConfig(v *AddCdnDomainRequestLiveConfig) *AddCdnDomainRequest {
  s.LiveConfig = v
  return s
}

func (s *AddCdnDomainRequest) SetPublishPoints(v []*AddCdnDomainRequestPublishPoints) *AddCdnDomainRequest {
  s.PublishPoints = v
  return s
}

func (s *AddCdnDomainRequest) SetSsl(v *AddCdnDomainRequestSsl) *AddCdnDomainRequest {
  s.Ssl = v
  return s
}

type AddCdnDomainRequestOriginConfig struct {
  // {"en":"Origin site address, which can be an IP or a domain name.
  // 1. IP is separated by semicolons and multiple IPs are supported.
  // 2. Only one domain name can be entered. IP and domain names cannot be entered at the same time.
  // 3. Maximum character limit is 500.", "zh_CN":"回源地址，可以是IP或域名。 1. IP以分号分隔，支持多个。 2. 域名只能输入一个。IP与域名不能同时输入。 3.限制最大不能超过500个字符长度。"}
  OriginIps *string `json:"origin-ips,omitempty" xml:"origin-ips,omitempty"`
  // {"en":"The Origin HOST for changing the HOST field in the return source HTTP request header.
  // Note: It should be domain or IP format. For domain name format, each segement separated by a dot, does not exceed 62 characters, the total length should not exceed 128 characters.", "zh_CN":"回源HOST，用于更改回源HTTP请求头中的HOST字段。支持格式为: 域名或ip。
  // 注意：必须符合ip/域名格式规范。如果是域名，则域名每段（点号分隔）长度小于等于62，域名总长度小于等于128。"}
  DefaultOriginHostHeader *string `json:"default-origin-host-header,omitempty" xml:"default-origin-host-header,omitempty"`
  // {"en":"Origin port, Value cannot be zero", "zh_CN":"回源请求端口，不能为0"}
  OriginPort *string `json:"origin-port,omitempty" xml:"origin-port,omitempty"`
}

func (s AddCdnDomainRequestOriginConfig) String() string {
  return tea.Prettify(s)
}

func (s AddCdnDomainRequestOriginConfig) GoString() string {
  return s.String()
}

func (s *AddCdnDomainRequestOriginConfig) SetOriginIps(v string) *AddCdnDomainRequestOriginConfig {
  s.OriginIps = &v
  return s
}

func (s *AddCdnDomainRequestOriginConfig) SetDefaultOriginHostHeader(v string) *AddCdnDomainRequestOriginConfig {
  s.DefaultOriginHostHeader = &v
  return s
}

func (s *AddCdnDomainRequestOriginConfig) SetOriginPort(v string) *AddCdnDomainRequestOriginConfig {
  s.OriginPort = &v
  return s
}

type AddCdnDomainRequestLiveConfig struct {
  // {"en":"The live push-pull stream type, the optional values are pull and push, pull means pull flow; push means push flow.", "zh_CN":"直播推拉流类型，可选值为pull和push，pull表示拉流；   push表示推流。"}
  StreamType *string `json:"stream-type,omitempty" xml:"stream-type,omitempty"`
  // {"en":"The push-pull domain name is used to set the push-flow domain name corresponding to the live streaming domain name. When the stream-type is pull, at least one of the source IP address and the corresponding push-stream domain name is not empty. When the stream-type is push, Incoming.", "zh_CN":"配套推流域名，用于设置直播拉流域名对应的推流域名，当stream-type为pull时，源站IP和配套推流域名至少一个不为空；当stream-type为push时，无需传入。"}
  OriginPushHost *string `json:"origin-push-host,omitempty" xml:"origin-push-host,omitempty"`
  // {"en":"Source station IP. When the stream-type is pull, at least one of the source station IP and the companion push stream domain name is not empty.
  // 1. If it is a push-pull flow package, fill in 127.0.0.1, and the system will also default to 127.0.0.1.
  // 2. If it is directly returning to the source, fill in the source IP of the source pull stream.
  // ", "zh_CN":"源站IP，当stream-type为pull时，源站IP和配套推流域名至少一个不为空。
  // 1、如果是推拉流配套，则填写127.0.0.1，不传系统也默认为127.0.0.1
  // 2、如果是直接回源拉流，则填写回源拉流的源站IP"}
  LiveConfigOriginIps *string `json:"origin-ips,omitempty" xml:"origin-ips,omitempty"`
}

func (s AddCdnDomainRequestLiveConfig) String() string {
  return tea.Prettify(s)
}

func (s AddCdnDomainRequestLiveConfig) GoString() string {
  return s.String()
}

func (s *AddCdnDomainRequestLiveConfig) SetStreamType(v string) *AddCdnDomainRequestLiveConfig {
  s.StreamType = &v
  return s
}

func (s *AddCdnDomainRequestLiveConfig) SetOriginPushHost(v string) *AddCdnDomainRequestLiveConfig {
  s.OriginPushHost = &v
  return s
}

func (s *AddCdnDomainRequestLiveConfig) SetLiveConfigOriginIps(v string) *AddCdnDomainRequestLiveConfig {
  s.LiveConfigOriginIps = &v
  return s
}

type AddCdnDomainRequestPublishPoints struct     {
  // {"en":"Publish point, support multiple, multiple values separated by ||,do not pass the system by default to generate a publishing point uri for slash /", "zh_CN":"发布点，支持多个，多个值用||隔开，不传系统默认生成一条发布点uri为“/”"}
  Uri *string `json:"uri,omitempty" xml:"uri,omitempty"`
}

func (s AddCdnDomainRequestPublishPoints) String() string {
  return tea.Prettify(s)
}

func (s AddCdnDomainRequestPublishPoints) GoString() string {
  return s.String()
}

func (s *AddCdnDomainRequestPublishPoints) SetUri(v string) *AddCdnDomainRequestPublishPoints {
  s.Uri = &v
  return s
}

type AddCdnDomainRequestSsl struct {
  // {"en":"Use a certificate, the optional values are true and false, true means to use the certificate, false means not to use the certificate", "zh_CN":"使用证书，可选值为true和false，true表示使用证书，false表示不使用证书"}
  UseSsl *string `json:"use-ssl,omitempty" xml:"use-ssl,omitempty"`
  // {"en":"Use sni certificate, the optional values are true and false, true means use sni certificate, false means use shared certificate (not supported)", "zh_CN":"使用sni证书，可选值为true和false，true表示使用sni证书，false表示使用合用证书（暂不支持）"}
  UseForSni *string `json:"use-for-sni,omitempty" xml:"use-for-sni,omitempty"`
  // {"en":"Use sni certificate, the optional values are true and false, true means use sni certificate, false means use shared certificate (not supported)", "zh_CN":"证书ID，新增证书成功后，系统返回的证书ID，use-ssl为true时，才能传ssl-certificate-id。"}
  SslCertificateId *int `json:"ssl-certificate-id,omitempty" xml:"ssl-certificate-id,omitempty"`
}

func (s AddCdnDomainRequestSsl) String() string {
  return tea.Prettify(s)
}

func (s AddCdnDomainRequestSsl) GoString() string {
  return s.String()
}

func (s *AddCdnDomainRequestSsl) SetUseSsl(v string) *AddCdnDomainRequestSsl {
  s.UseSsl = &v
  return s
}

func (s *AddCdnDomainRequestSsl) SetUseForSni(v string) *AddCdnDomainRequestSsl {
  s.UseForSni = &v
  return s
}

func (s *AddCdnDomainRequestSsl) SetSslCertificateId(v int) *AddCdnDomainRequestSsl {
  s.SslCertificateId = &v
  return s
}

type AddCdnDomainResponse struct {
  // {"en":"httpstatus=202; Indicates that the new domain API was successfully invoked, and the current deployment of the new domain can be viewed using x-cnc-request-id in the header", "zh_CN":"httpstatus=202;   表示成功调用新增域名接口，可使用header中的x-cnc-request-id查看当前新增域名的部署情况"}
  HttpStatus *int `json:"http status code,omitempty" xml:"http status code,omitempty" require:"true"`
  // {"en":"Uniquely identified id for querying tasks per request (for all API)", "zh_CN":"唯一标示的id，用于查询每次请求的任务 （适用全部接口）"}
  XCncRequestId *string `json:"x-cnc-request-id,omitempty" xml:"x-cnc-request-id,omitempty" require:"true"`
  // {"en":"The URL used to access the domain information, where domain-id is the unique token generated by our cloud platform for the domain name and whose value is a string.", "zh_CN":"响应信用于访问该域名信息的URL，其中domain-id为我司云平台为该域名生成的唯一标示，其值为字符串。"}
  Location *string `json:"location,omitempty" xml:"location,omitempty" require:"true"`
  // {"en":"The name of the service domain automatically generated by the My company, for example: xxxx.cdn30.com", "zh_CN":"由我司自动生成的服务域名名称，例如：xxxx.cdn30.com"}
  Cname *string `json:"cname,omitempty" xml:"cname,omitempty" require:"true"`
  // {"en":"Request result code", "zh_CN":"请求结果状态码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response information, when success is successful", "zh_CN":"响应信息，成功时为success"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s AddCdnDomainResponse) String() string {
  return tea.Prettify(s)
}

func (s AddCdnDomainResponse) GoString() string {
  return s.String()
}

func (s *AddCdnDomainResponse) SetHttpStatus(v int) *AddCdnDomainResponse {
  s.HttpStatus = &v
  return s
}

func (s *AddCdnDomainResponse) SetXCncRequestId(v string) *AddCdnDomainResponse {
  s.XCncRequestId = &v
  return s
}

func (s *AddCdnDomainResponse) SetLocation(v string) *AddCdnDomainResponse {
  s.Location = &v
  return s
}

func (s *AddCdnDomainResponse) SetCname(v string) *AddCdnDomainResponse {
  s.Cname = &v
  return s
}

func (s *AddCdnDomainResponse) SetCode(v string) *AddCdnDomainResponse {
  s.Code = &v
  return s
}

func (s *AddCdnDomainResponse) SetMessage(v string) *AddCdnDomainResponse {
  s.Message = &v
  return s
}

type AddCdnDomainPaths struct {
}

func (s AddCdnDomainPaths) String() string {
  return tea.Prettify(s)
}

func (s AddCdnDomainPaths) GoString() string {
  return s.String()
}

type AddCdnDomainParameters struct {
}

func (s AddCdnDomainParameters) String() string {
  return tea.Prettify(s)
}

func (s AddCdnDomainParameters) GoString() string {
  return s.String()
}

type AddCdnDomainRequestHeader struct {
}

func (s AddCdnDomainRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s AddCdnDomainRequestHeader) GoString() string {
  return s.String()
}

type AddCdnDomainResponseHeader struct {
}

func (s AddCdnDomainResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s AddCdnDomainResponseHeader) GoString() string {
  return s.String()
}




type QueryCdnwDomainAliasRelationForWplusRequest struct {
  // {"en":"Specifies the domainAlias for the query, allows multiple domain-aliases. For example:
  // <domainAlias> <aliasDomainName>alias1.example.com</aliasDomainName> <aliasDomainName>alias2.example.com</aliasDomainName> <aliasDomainName>alias2.example.com</aliasDomainName> </domainAlias>", "zh_CN":"指定要查询的别名列表"}
  QueryCdnwDomainAliasRelationForWplusDomainAlias []*string `json:"domainAlias,omitempty" xml:"domainAlias,omitempty" require:"true" type:"Repeated"`
}

func (s QueryCdnwDomainAliasRelationForWplusRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryCdnwDomainAliasRelationForWplusRequest) GoString() string {
  return s.String()
}

func (s *QueryCdnwDomainAliasRelationForWplusRequest) SetDomainAlias(v []*string) *QueryCdnwDomainAliasRelationForWplusRequest {
  s.QueryCdnwDomainAliasRelationForWplusDomainAlias = v
  return s
}

type QueryCdnwDomainAliasRelationForWplusResponse struct {
  // {"en":"Response code, 0 means successful.", "zh_CN":"接口响应code，0代表成功。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"esponse error message if failed.", "zh_CN":"接口响应信息，success代表成功，失败则提供失败信息。"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Response data array.", "zh_CN":"接口响应数据"}
  Data []*QueryCdnwDomainAliasRelationForWplusDomainAlias `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s QueryCdnwDomainAliasRelationForWplusResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryCdnwDomainAliasRelationForWplusResponse) GoString() string {
  return s.String()
}

func (s *QueryCdnwDomainAliasRelationForWplusResponse) SetCode(v string) *QueryCdnwDomainAliasRelationForWplusResponse {
  s.Code = &v
  return s
}

func (s *QueryCdnwDomainAliasRelationForWplusResponse) SetMessage(v string) *QueryCdnwDomainAliasRelationForWplusResponse {
  s.Message = &v
  return s
}

func (s *QueryCdnwDomainAliasRelationForWplusResponse) SetData(v []*QueryCdnwDomainAliasRelationForWplusDomainAlias) *QueryCdnwDomainAliasRelationForWplusResponse {
  s.Data = v
  return s
}

type QueryCdnwDomainAliasRelationForWplusDomainAlias struct {
  // {"en":"Domain ID that related to queried domainAlias .", "zh_CN":"别名关联的域名id"}
  DomainId *int `json:"domainId,omitempty" xml:"domainId,omitempty" require:"true"`
  // {"en":"Domain name that related to queried domainAlias .", "zh_CN":"别名关联的域名名称。"}
  DomainName *string `json:"domainName,omitempty" xml:"domainName,omitempty" require:"true"`
  // {"en":"Domain alias list.", "zh_CN":"别名列表"}
  QueryCdnwDomainAliasRelationForWplusDomainAlias []*string `json:"domainAlias,omitempty" xml:"domainAlias,omitempty" require:"true" type:"Repeated"`
}

func (s QueryCdnwDomainAliasRelationForWplusDomainAlias) String() string {
  return tea.Prettify(s)
}

func (s QueryCdnwDomainAliasRelationForWplusDomainAlias) GoString() string {
  return s.String()
}

func (s *QueryCdnwDomainAliasRelationForWplusDomainAlias) SetDomainId(v int) *QueryCdnwDomainAliasRelationForWplusDomainAlias {
  s.DomainId = &v
  return s
}

func (s *QueryCdnwDomainAliasRelationForWplusDomainAlias) SetDomainName(v string) *QueryCdnwDomainAliasRelationForWplusDomainAlias {
  s.DomainName = &v
  return s
}

func (s *QueryCdnwDomainAliasRelationForWplusDomainAlias) SetDomainAlias(v []*string) *QueryCdnwDomainAliasRelationForWplusDomainAlias {
  s.QueryCdnwDomainAliasRelationForWplusDomainAlias = v
  return s
}

type QueryCdnwDomainAliasRelationForWplusPaths struct {
}

func (s QueryCdnwDomainAliasRelationForWplusPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryCdnwDomainAliasRelationForWplusPaths) GoString() string {
  return s.String()
}

type QueryCdnwDomainAliasRelationForWplusParameters struct {
}

func (s QueryCdnwDomainAliasRelationForWplusParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryCdnwDomainAliasRelationForWplusParameters) GoString() string {
  return s.String()
}

type QueryCdnwDomainAliasRelationForWplusRequestHeader struct {
}

func (s QueryCdnwDomainAliasRelationForWplusRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryCdnwDomainAliasRelationForWplusRequestHeader) GoString() string {
  return s.String()
}

type QueryCdnwDomainAliasRelationForWplusResponseHeader struct {
}

func (s QueryCdnwDomainAliasRelationForWplusResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryCdnwDomainAliasRelationForWplusResponseHeader) GoString() string {
  return s.String()
}




type UpdateChangeServerRequest struct {
  // {"en":"Change servers configuration, parent tag
  // 1. This must be filled when the hotlinking configuration of streaming media needs to be set
  // 2. Empty the configuration for <change-servers/>", "zh_CN":"【接入域名跳转】
  // 注意：
  // 1、需要取消【接入域名跳转】时，可以传入空节点<change-servers></change-servers>。
  // 2、表示需要设置【接入域名跳转】，此项必填"}
  ChangeServers []*UpdateChangeServerRequestChangeServers `json:"change-servers,omitempty" xml:"change-servers,omitempty" require:"true" type:"Repeated"`
}

func (s UpdateChangeServerRequest) String() string {
  return tea.Prettify(s)
}

func (s UpdateChangeServerRequest) GoString() string {
  return s.String()
}

func (s *UpdateChangeServerRequest) SetChangeServers(v []*UpdateChangeServerRequestChangeServers) *UpdateChangeServerRequest {
  s.ChangeServers = v
  return s
}

type UpdateChangeServerRequestChangeServers struct     {
  // {"en":"If it is a universal domain name, set it to a universal domain name, for example, *.56.com.", "zh_CN":"如果是泛域名，需要填写为泛域名，例如：*.56.com"}
  TargetServer *string `json:"target-server,omitempty" xml:"target-server,omitempty"`
  // {"en":"Data-id is to indicate a specific group configuration when the client has multiple groups of configurations. Data-id can be retrieved through a query interface. Note: 
  // A. If data-id is passed, it means that one group of configuration items is specified to be modified, and no other group configuration items need to be modified. 
  // B. If multiple groups of configurations are included, some of them are configured with data-id and others are not, then the expression of data-id is used to modify a specific group of configurations, and a new group of configurations is added on the original basis without the expression of data-id. 
  // C. If the data-id is not transmitted, it means that the original configuration will be fully covered by this configuration. 
  // D. If no configuration parameter is passed, only domain name and secondary label are passed, which means that all configuration of domain name secondary service corresponding to this interface is cleared. 
  // E. If there is no specific configuration item in a set of configurations, the data-id must be filled in, and the value is the actual data-id, which means clearing the value of the corresponding data-id configuration item; it is not allowed that there is no specific configuration item or data-id in a set of configurations.", 
  //       "zh_CN":"配置多组配置时，具体某组配置的id。dataId可以通过查询接口获取。 注意： 
  // a、如果有传dataId，说明指定修改其中一组配置项内容，不需求修改其他组配置内容不需要入参； 
  // b、如果入参多组配置，其中有些组配置有传dataId，有些没有传，则有传dataId的表示修改具体某组配置，没有传dataId的表示在原来基础上新增一组配置； 
  // c、如果入参都没有传dataId,表示用本次的配置全量覆盖原先配置； 
  // d、如果入参没有传任何配置项参数，只传了域名和二级标签，表示清空这个接口对应域名二级服务所有配置； 
  // e、如果一组配置没有具体的配置项，则dataId必填，且值为实际存在的dataId，表示清空这个dataId对应配置项的值；不允许一组配置没有具体的配置项也没有dataId。
  // "}
  DataId *int `json:"dataId,omitempty" xml:"dataId,omitempty"`
}

func (s UpdateChangeServerRequestChangeServers) String() string {
  return tea.Prettify(s)
}

func (s UpdateChangeServerRequestChangeServers) GoString() string {
  return s.String()
}

func (s *UpdateChangeServerRequestChangeServers) SetTargetServer(v string) *UpdateChangeServerRequestChangeServers {
  s.TargetServer = &v
  return s
}

func (s *UpdateChangeServerRequestChangeServers) SetDataId(v int) *UpdateChangeServerRequestChangeServers {
  s.DataId = &v
  return s
}

type UpdateChangeServerResponse struct {
  // {"en":"The error code", "zh_CN":"错误码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"The message body", "zh_CN":"消息体"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Returns the body of the data.", "zh_CN":"返回数据体。"}
  Data *string `json:"data,omitempty" xml:"data,omitempty" require:"true"`
  // {"en":"Uniquely labeled id for querying each requested task (for all interfaces)", "zh_CN":"唯一标示的id，用于查询每次请求的任务 （适用全部接口）"}
  XCncRequestId *string `json:"x-cnc-request-id,omitempty" xml:"x-cnc-request-id,omitempty" require:"true"`
}

func (s UpdateChangeServerResponse) String() string {
  return tea.Prettify(s)
}

func (s UpdateChangeServerResponse) GoString() string {
  return s.String()
}

func (s *UpdateChangeServerResponse) SetCode(v string) *UpdateChangeServerResponse {
  s.Code = &v
  return s
}

func (s *UpdateChangeServerResponse) SetMessage(v string) *UpdateChangeServerResponse {
  s.Message = &v
  return s
}

func (s *UpdateChangeServerResponse) SetData(v string) *UpdateChangeServerResponse {
  s.Data = &v
  return s
}

func (s *UpdateChangeServerResponse) SetXCncRequestId(v string) *UpdateChangeServerResponse {
  s.XCncRequestId = &v
  return s
}

type UpdateChangeServerPaths struct {
  // {"en":"The domain whoes need query config.", "zh_CN":"需要查询配置的域名或域名id"}
  DomainName *string `json:"domain-name,omitempty" xml:"domain-name,omitempty"`
}

func (s UpdateChangeServerPaths) String() string {
  return tea.Prettify(s)
}

func (s UpdateChangeServerPaths) GoString() string {
  return s.String()
}

func (s *UpdateChangeServerPaths) SetDomainName(v string) *UpdateChangeServerPaths {
  s.DomainName = &v
  return s
}

type UpdateChangeServerParameters struct {
}

func (s UpdateChangeServerParameters) String() string {
  return tea.Prettify(s)
}

func (s UpdateChangeServerParameters) GoString() string {
  return s.String()
}

type UpdateChangeServerRequestHeader struct {
}

func (s UpdateChangeServerRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s UpdateChangeServerRequestHeader) GoString() string {
  return s.String()
}

type UpdateChangeServerResponseHeader struct {
}

func (s UpdateChangeServerResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s UpdateChangeServerResponseHeader) GoString() string {
  return s.String()
}




type GetFuzzyPagingDomainListRequest struct {
  // {"en":"Page number must be a positive integer greater than 0", "zh_CN":"分页的页码，必须为大于0的正整数"}
  PageNumber *int `json:"pageNumber,omitempty" xml:"pageNumber,omitempty" require:"true"`
  // {"en":"Number of domain name data items for paging, must be a positive integer greater than 0", "zh_CN":"分页的域名数据条数，必须大于0的正整数"}
  PageSize *int `json:"pageSize,omitempty" xml:"pageSize,omitempty" require:"true"`
  // {"en":"Specifies the service type of the query, only one type per query, and no default lookup for all types", "zh_CN":"指定查询的服务类型，每次查询只能传一个类型，不传默认查全部类型"}
  ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty"`
  // {"en":"Specifies the accelerated domain name for the query, allows multiple domains, commas delimited, and no default lookup of all domain names", "zh_CN":"指定查询的加速域名，允许多个域名，逗号分隔，不传默认查全部域名"}
  DomainName []*string `json:"domainName,omitempty" xml:"domainName,omitempty" type:"Repeated"`
  // {"en":"Query to accelerated domain name, optional value is: fuzzy_match for fuzzy query; Full_match represents an exact query
  // No fuzzy_match by default, for accelerated domain name only", "zh_CN":"查询加速域名的方式，可选值为：fuzzy_match表示模糊查询；full_match表示精确查询
  // 不传默认为fuzzy_match，仅针对加速域名"}
  QueryType *string `json:"queryType,omitempty" xml:"queryType,omitempty"`
  // {"en":"Query start time, support for years, months, days, hours, minutes, and seconds, for example: 20170101.09 million. Time equals", "zh_CN":"查询开始时间，支持范围为年月日时分秒，例如：20170101090000。时间含等于"}
  StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
  // {"en":"Query end time, query time within the existence of the accelerated domain name, time is equal to, do not pass the default query all", "zh_CN":"查询结束时间，查询时间段内存在的加速域名，时间含等于，不传默认查询所有"}
  EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
  // {"en":"Accelerate the status of the domain name, enabled indicates that it is in effect; Disabled indicates that it is Disabled; Deploying means in the process of deployment; Checking indicates that the audit is in progress; Disabling: Indicates disabled, no default lookup for all", "zh_CN":"加速域名的状态，enabled表示已生效；disabled表示已禁用；deploying表示部署中；checking表示审核中；disabling:表示禁用中，不传默认查全部"}
  DomainStatus *string `json:"domainStatus,omitempty" xml:"domainStatus,omitempty"`
}

func (s GetFuzzyPagingDomainListRequest) String() string {
  return tea.Prettify(s)
}

func (s GetFuzzyPagingDomainListRequest) GoString() string {
  return s.String()
}

func (s *GetFuzzyPagingDomainListRequest) SetPageNumber(v int) *GetFuzzyPagingDomainListRequest {
  s.PageNumber = &v
  return s
}

func (s *GetFuzzyPagingDomainListRequest) SetPageSize(v int) *GetFuzzyPagingDomainListRequest {
  s.PageSize = &v
  return s
}

func (s *GetFuzzyPagingDomainListRequest) SetServiceType(v string) *GetFuzzyPagingDomainListRequest {
  s.ServiceType = &v
  return s
}

func (s *GetFuzzyPagingDomainListRequest) SetDomainName(v []*string) *GetFuzzyPagingDomainListRequest {
  s.DomainName = v
  return s
}

func (s *GetFuzzyPagingDomainListRequest) SetQueryType(v string) *GetFuzzyPagingDomainListRequest {
  s.QueryType = &v
  return s
}

func (s *GetFuzzyPagingDomainListRequest) SetStartTime(v string) *GetFuzzyPagingDomainListRequest {
  s.StartTime = &v
  return s
}

func (s *GetFuzzyPagingDomainListRequest) SetEndTime(v string) *GetFuzzyPagingDomainListRequest {
  s.EndTime = &v
  return s
}

func (s *GetFuzzyPagingDomainListRequest) SetDomainStatus(v string) *GetFuzzyPagingDomainListRequest {
  s.DomainStatus = &v
  return s
}

type GetFuzzyPagingDomainListResponse struct {
  // {"en":"httpstatus=202; Indicates that the new domain API was successfully invoked, and the current deployment of the new domain can be viewed using x-cnc-request-id in the header", "zh_CN":"httpstatus=202;   表示成功调用新增域名接口，可使用header中的x-cnc-request-id查看当前新增域名的部署情况"}
  Code *int `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Uniquely identified id for querying tasks per request (for all API)", "zh_CN":"唯一标示的id，用于查询每次请求的任务 （适用全部接口）"}
  XCncRequestId *string `json:"x-cnc-request-id,omitempty" xml:"x-cnc-request-id,omitempty" require:"true"`
  // {"en":"Responses the page number of the data", "zh_CN":"所有满足条件的数据总条数"}
  TotalCount *int `json:"totalCount,omitempty" xml:"totalCount,omitempty" require:"true"`
  // {"en":"total pages", "zh_CN":"总页数"}
  TotalPageNumber *int `json:"totalPageNumber,omitempty" xml:"totalPageNumber,omitempty" require:"true"`
  // {"en":"Responses the page number of the data", "zh_CN":"返回数据的页码"}
  PageNumber *int `json:"pageNumber,omitempty" xml:"pageNumber,omitempty" require:"true"`
  // {"en":"Number of data page", "zh_CN":"每个页面的数据条数"}
  PageSize *int `json:"pageSize,omitempty" xml:"pageSize,omitempty" require:"true"`
  // {"en":"Responses status information for the accelerated domain name", "zh_CN":"返回加速域名的状态信息"}
  ResultList []*GetFuzzyPagingDomainListResponseResultList `json:"resultList,omitempty" xml:"resultList,omitempty" require:"true" type:"Repeated"`
}

func (s GetFuzzyPagingDomainListResponse) String() string {
  return tea.Prettify(s)
}

func (s GetFuzzyPagingDomainListResponse) GoString() string {
  return s.String()
}

func (s *GetFuzzyPagingDomainListResponse) SetCode(v int) *GetFuzzyPagingDomainListResponse {
  s.Code = &v
  return s
}

func (s *GetFuzzyPagingDomainListResponse) SetXCncRequestId(v string) *GetFuzzyPagingDomainListResponse {
  s.XCncRequestId = &v
  return s
}

func (s *GetFuzzyPagingDomainListResponse) SetTotalCount(v int) *GetFuzzyPagingDomainListResponse {
  s.TotalCount = &v
  return s
}

func (s *GetFuzzyPagingDomainListResponse) SetTotalPageNumber(v int) *GetFuzzyPagingDomainListResponse {
  s.TotalPageNumber = &v
  return s
}

func (s *GetFuzzyPagingDomainListResponse) SetPageNumber(v int) *GetFuzzyPagingDomainListResponse {
  s.PageNumber = &v
  return s
}

func (s *GetFuzzyPagingDomainListResponse) SetPageSize(v int) *GetFuzzyPagingDomainListResponse {
  s.PageSize = &v
  return s
}

func (s *GetFuzzyPagingDomainListResponse) SetResultList(v []*GetFuzzyPagingDomainListResponseResultList) *GetFuzzyPagingDomainListResponse {
  s.ResultList = v
  return s
}

type GetFuzzyPagingDomainListResponseResultList struct     {
  // {"en":"Accelerated domain CNAME corresponding to CNAME, for example: 7nt6mrh7sdkslj.cdn30.com", "zh_CN":"加速域名对应的CNAME域名，例如：7nt6mrh7sdkslj.cdn30.com"}
  Cname *string `json:"cname,omitempty" xml:"cname,omitempty" require:"true"`
  // {"en":"Configuration name", "zh_CN":"配置单名称"}
  ConfigFormName *string `json:"configFormName,omitempty" xml:"configFormName,omitempty" require:"true"`
  // {"en":"The time format is: 20160323112310", "zh_CN":"时间格式为：20160323112310"}
  CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty" require:"true"`
  // {"en":"Corresponding domain ID", "zh_CN":"对应的域名ID"}
  DomainId *string `json:"domainId,omitempty" xml:"domainId,omitempty" require:"true"`
  // {"en":"Accelerated domain name", "zh_CN":"加速域名名称"}
  DomainName *string `json:"domainName,omitempty" xml:"domainName,omitempty" require:"true"`
  // {"en":"Operator of this query", "zh_CN":"本次查询的操作者"}
  Operator *string `json:"operator,omitempty" xml:"operator,omitempty" require:"true"`
  // {"en":"Accelerate the origin IP of a domain name", "zh_CN":"加速域名的回源IP"}
  OriginIps *string `json:"originIps,omitempty" xml:"originIps,omitempty" require:"true"`
  // {"en":"Service type for accelerated domain name", "zh_CN":"加速域名的服务类型"}
  ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty" require:"true"`
  // {"en":"Status of accelerated domain name.", "zh_CN":"加速域名的状态"}
  DomainStatus *string `json:"domainStatus,omitempty" xml:"domainStatus,omitempty" require:"true"`
  // {"en":"Deployment version code", "zh_CN":"部署版本号"}
  DeployVersion *string `json:"deployVersion,omitempty" xml:"deployVersion,omitempty" require:"true"`
  // {"en":"Does the domain name enable CDN acceleration services, Y and N?", "zh_CN":"域名是否启用CDN加速服务，Y和N"}
  CdnServiceStatus *string `json:"cdnServiceStatus,omitempty" xml:"cdnServiceStatus,omitempty" require:"true"`
  // {"en":"Whether the accelerated domain name is enabled, Y and N?", "zh_CN":"加速域名是否启用，Y和N"}
  IsEnabled *string `json:"isEnabled,omitempty" xml:"isEnabled,omitempty" require:"true"`
}

func (s GetFuzzyPagingDomainListResponseResultList) String() string {
  return tea.Prettify(s)
}

func (s GetFuzzyPagingDomainListResponseResultList) GoString() string {
  return s.String()
}

func (s *GetFuzzyPagingDomainListResponseResultList) SetCname(v string) *GetFuzzyPagingDomainListResponseResultList {
  s.Cname = &v
  return s
}

func (s *GetFuzzyPagingDomainListResponseResultList) SetConfigFormName(v string) *GetFuzzyPagingDomainListResponseResultList {
  s.ConfigFormName = &v
  return s
}

func (s *GetFuzzyPagingDomainListResponseResultList) SetCreateTime(v string) *GetFuzzyPagingDomainListResponseResultList {
  s.CreateTime = &v
  return s
}

func (s *GetFuzzyPagingDomainListResponseResultList) SetDomainId(v string) *GetFuzzyPagingDomainListResponseResultList {
  s.DomainId = &v
  return s
}

func (s *GetFuzzyPagingDomainListResponseResultList) SetDomainName(v string) *GetFuzzyPagingDomainListResponseResultList {
  s.DomainName = &v
  return s
}

func (s *GetFuzzyPagingDomainListResponseResultList) SetOperator(v string) *GetFuzzyPagingDomainListResponseResultList {
  s.Operator = &v
  return s
}

func (s *GetFuzzyPagingDomainListResponseResultList) SetOriginIps(v string) *GetFuzzyPagingDomainListResponseResultList {
  s.OriginIps = &v
  return s
}

func (s *GetFuzzyPagingDomainListResponseResultList) SetServiceType(v string) *GetFuzzyPagingDomainListResponseResultList {
  s.ServiceType = &v
  return s
}

func (s *GetFuzzyPagingDomainListResponseResultList) SetDomainStatus(v string) *GetFuzzyPagingDomainListResponseResultList {
  s.DomainStatus = &v
  return s
}

func (s *GetFuzzyPagingDomainListResponseResultList) SetDeployVersion(v string) *GetFuzzyPagingDomainListResponseResultList {
  s.DeployVersion = &v
  return s
}

func (s *GetFuzzyPagingDomainListResponseResultList) SetCdnServiceStatus(v string) *GetFuzzyPagingDomainListResponseResultList {
  s.CdnServiceStatus = &v
  return s
}

func (s *GetFuzzyPagingDomainListResponseResultList) SetIsEnabled(v string) *GetFuzzyPagingDomainListResponseResultList {
  s.IsEnabled = &v
  return s
}

type GetFuzzyPagingDomainListPaths struct {
}

func (s GetFuzzyPagingDomainListPaths) String() string {
  return tea.Prettify(s)
}

func (s GetFuzzyPagingDomainListPaths) GoString() string {
  return s.String()
}

type GetFuzzyPagingDomainListParameters struct {
}

func (s GetFuzzyPagingDomainListParameters) String() string {
  return tea.Prettify(s)
}

func (s GetFuzzyPagingDomainListParameters) GoString() string {
  return s.String()
}

type GetFuzzyPagingDomainListRequestHeader struct {
}

func (s GetFuzzyPagingDomainListRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetFuzzyPagingDomainListRequestHeader) GoString() string {
  return s.String()
}

type GetFuzzyPagingDomainListResponseHeader struct {
}

func (s GetFuzzyPagingDomainListResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetFuzzyPagingDomainListResponseHeader) GoString() string {
  return s.String()
}




type QueryCustomerAnycastIPForWplusRequest struct {
}

func (s QueryCustomerAnycastIPForWplusRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryCustomerAnycastIPForWplusRequest) GoString() string {
  return s.String()
}

type QueryCustomerAnycastIPForWplusResponse struct {
  // {"en":"code", "zh_CN":"错误码，成功为0"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"error message", "zh_CN":"错误信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"data", "zh_CN":"anycastIp详细"}
  Data []*QueryCustomerAnycastIPForWplusAnycastIPDetail `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s QueryCustomerAnycastIPForWplusResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryCustomerAnycastIPForWplusResponse) GoString() string {
  return s.String()
}

func (s *QueryCustomerAnycastIPForWplusResponse) SetCode(v string) *QueryCustomerAnycastIPForWplusResponse {
  s.Code = &v
  return s
}

func (s *QueryCustomerAnycastIPForWplusResponse) SetMessage(v string) *QueryCustomerAnycastIPForWplusResponse {
  s.Message = &v
  return s
}

func (s *QueryCustomerAnycastIPForWplusResponse) SetData(v []*QueryCustomerAnycastIPForWplusAnycastIPDetail) *QueryCustomerAnycastIPForWplusResponse {
  s.Data = v
  return s
}

type QueryCustomerAnycastIPForWplusAnycastIPDetail struct {
  // {"en":"IP", "zh_CN":"IP"}
  Ip *string `json:"ip,omitempty" xml:"ip,omitempty" require:"true"`
  // {"en":"status", "zh_CN":"域名加速状态: USED, AVAILABLE"}
  Status *string `json:"status,omitempty" xml:"status,omitempty" require:"true"`
  // {"en":"Record Status, For example, lock or unlock; data of length 1 or 2 can be passed.", "zh_CN":"记录状态，例如锁定、非锁定等，可以传长度为1或2的数据"}
  RecordStatus *string `json:"recordStatus,omitempty" xml:"recordStatus,omitempty" require:"true"`
  // {"en":"Is china mainland, 0:NO,1: YES", "zh_CN":"是否属于大陆，0：否，1：是"}
  IsChinaMainland *string `json:"isChinaMainland,omitempty" xml:"isChinaMainland,omitempty" require:"true"`
}

func (s QueryCustomerAnycastIPForWplusAnycastIPDetail) String() string {
  return tea.Prettify(s)
}

func (s QueryCustomerAnycastIPForWplusAnycastIPDetail) GoString() string {
  return s.String()
}

func (s *QueryCustomerAnycastIPForWplusAnycastIPDetail) SetIp(v string) *QueryCustomerAnycastIPForWplusAnycastIPDetail {
  s.Ip = &v
  return s
}

func (s *QueryCustomerAnycastIPForWplusAnycastIPDetail) SetStatus(v string) *QueryCustomerAnycastIPForWplusAnycastIPDetail {
  s.Status = &v
  return s
}

func (s *QueryCustomerAnycastIPForWplusAnycastIPDetail) SetRecordStatus(v string) *QueryCustomerAnycastIPForWplusAnycastIPDetail {
  s.RecordStatus = &v
  return s
}

func (s *QueryCustomerAnycastIPForWplusAnycastIPDetail) SetIsChinaMainland(v string) *QueryCustomerAnycastIPForWplusAnycastIPDetail {
  s.IsChinaMainland = &v
  return s
}

type QueryCustomerAnycastIPForWplusPaths struct {
}

func (s QueryCustomerAnycastIPForWplusPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryCustomerAnycastIPForWplusPaths) GoString() string {
  return s.String()
}

type QueryCustomerAnycastIPForWplusParameters struct {
}

func (s QueryCustomerAnycastIPForWplusParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryCustomerAnycastIPForWplusParameters) GoString() string {
  return s.String()
}

type QueryCustomerAnycastIPForWplusRequestHeader struct {
}

func (s QueryCustomerAnycastIPForWplusRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryCustomerAnycastIPForWplusRequestHeader) GoString() string {
  return s.String()
}

type QueryCustomerAnycastIPForWplusResponseHeader struct {
}

func (s QueryCustomerAnycastIPForWplusResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryCustomerAnycastIPForWplusResponseHeader) GoString() string {
  return s.String()
}




type QueryApiDomainListServiceRequest struct {
}

func (s QueryApiDomainListServiceRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryApiDomainListServiceRequest) GoString() string {
  return s.String()
}

type QueryApiDomainListServiceResponse struct {
  // {"en":"domain list", "zh_CN":"域名列表"}
  DomainList []*QueryApiDomainListServiceResponseDomainList `json:"domain-list,omitempty" xml:"domain-list,omitempty" require:"true" type:"Repeated"`
}

func (s QueryApiDomainListServiceResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryApiDomainListServiceResponse) GoString() string {
  return s.String()
}

func (s *QueryApiDomainListServiceResponse) SetDomainList(v []*QueryApiDomainListServiceResponseDomainList) *QueryApiDomainListServiceResponse {
  s.DomainList = v
  return s
}

type QueryApiDomainListServiceResponseDomainList struct     {
  // {"en":"Name of accelerated domain name", "zh_CN":"加速域名的名称"}
  DomainName *string `json:"domain-name,omitempty" xml:"domain-name,omitempty" require:"true"`
  // {"en":"The corresponding domain name ID: the domain name ID, used to perform the query and modification operations of the related domain name.", "zh_CN":"对应的域名ID：域名ID，用于执行相关域名的查询、修改操作等。"}
  DomainId *int `json:"domain-id,omitempty" xml:"domain-id,omitempty" require:"true"`
  // {"en":"Accelerated domain CNAME corresponding to CNAME, for example: 7nt6mrh7sdkslj.cdn30.com", "zh_CN":"加速域名对应的CNAME域名，例如：7nt6mrh7sdkslj.cdn30.com"}
  Cname *string `json:"cname,omitempty" xml:"cname,omitempty" require:"true"`
  // {"en":"Speed up the service type of the domain name, the value is:
  // Web/web-https: web acceleration / web acceleration - https
  // Wsa/wsa-https: Total Station Acceleration / Total Station Acceleration - https
  // Vodstream/vod-https: on-demand acceleration/on-demand acceleration - https
  // Download/dl-https: Download acceleration/download acceleration - https
  // livestream/live-https/cloudv-live: Live acceleration/Live acceleration - https/Cloud vedio for live
  // 1028: Content Acceleration;
  // 1115: Dynamic Web Acceleration;
  // 1369: Media Acceleration - RTMP
  // 1391: Download Acceleration
  // 1348: Media Acceleration Live Broadcast
  // 1551: Floodshield", "zh_CN":"加速域名的服务类型，取值：
  // web/web-https：网页加速/网页加速-https
  // wsa/wsa-https：全站加速/全站加速-https
  // vodstream/vod-https：点播加速/点播加速-https
  // download/dl-https：下载加速/下载加速-https
  // livestream/live-https/cloudv-live：直播加速/直播加速-https/云直播
  // appa/s-appa：应用加速/应用安全加速解决方案
  // 1028 : Content Acceleration;
  // 1115 : Dynamic Web Acceleration;
  // 1369 : Media Acceleration - RTMP
  // 1391 : Download Acceleration
  // 1348 : Media Acceleration Live Broadcast
  // 1551 : Floodshield"}
  ServiceType *string `json:"service-type,omitempty" xml:"service-type,omitempty" require:"true"`
  // {"en":"The deployment status of the accelerated domain name: Deployed indicates that the accelerated domain name configuration is complete; InProgress indicates that the deployment task for this accelerated domain name configuration is still in InProgress and may be in a queue, deploy, or fail in any one of the states", "zh_CN":"加速域名的部署状态：Deployed表示该加速域名配置完成部署；InProgress表示该加速域名配置的部署任务还在进行中，可能处于排队、部署中或失败任意一种状态"}
  Status *string `json:"status,omitempty" xml:"status,omitempty" require:"true"`
  // {"en":"Accelerate the CDN service status of the domain name: This is false when the accelerated domain name CDN service is canceled; this is true when the accelerated domain name CDN service is restored.", "zh_CN":"加速域名的CDN服务状态：当取消加速域名CDN服务后，此项为false；当恢复加速域名CDN服务后，此项为true"}
  CdnServiceStatus *string `json:"cdn-service-status,omitempty" xml:"cdn-service-status,omitempty" require:"true"`
  // {"en":"Accelerated domain activation: This is false when the accelerated domain name service is disabled; true when the accelerated domain name service is enabled", "zh_CN":"加速域名的启用状态：当禁用加速域名服务后，此项为false；当启用加速域名服务后，此项为true"}
  Enabled *string `json:"enabled,omitempty" xml:"enabled,omitempty" require:"true"`
  // {"en":"Domain name last modified time,
  // Format: 2024-01-01T22:30:00+08:00", "zh_CN":"域名最近修改时间，格式: 2024-01-01T22:30:00+08:00"}
  LastModified *string `json:"last-modified,omitempty" xml:"last-modified,omitempty" require:"true"`
  // {"en":"Billing areas.", "zh_CN":"计费区域"}
  BillingArea *string `json:"billing-areas,omitempty" xml:"billing-areas,omitempty" require:"true"`
}

func (s QueryApiDomainListServiceResponseDomainList) String() string {
  return tea.Prettify(s)
}

func (s QueryApiDomainListServiceResponseDomainList) GoString() string {
  return s.String()
}

func (s *QueryApiDomainListServiceResponseDomainList) SetDomainName(v string) *QueryApiDomainListServiceResponseDomainList {
  s.DomainName = &v
  return s
}

func (s *QueryApiDomainListServiceResponseDomainList) SetDomainId(v int) *QueryApiDomainListServiceResponseDomainList {
  s.DomainId = &v
  return s
}

func (s *QueryApiDomainListServiceResponseDomainList) SetCname(v string) *QueryApiDomainListServiceResponseDomainList {
  s.Cname = &v
  return s
}

func (s *QueryApiDomainListServiceResponseDomainList) SetServiceType(v string) *QueryApiDomainListServiceResponseDomainList {
  s.ServiceType = &v
  return s
}

func (s *QueryApiDomainListServiceResponseDomainList) SetStatus(v string) *QueryApiDomainListServiceResponseDomainList {
  s.Status = &v
  return s
}

func (s *QueryApiDomainListServiceResponseDomainList) SetCdnServiceStatus(v string) *QueryApiDomainListServiceResponseDomainList {
  s.CdnServiceStatus = &v
  return s
}

func (s *QueryApiDomainListServiceResponseDomainList) SetEnabled(v string) *QueryApiDomainListServiceResponseDomainList {
  s.Enabled = &v
  return s
}

func (s *QueryApiDomainListServiceResponseDomainList) SetLastModified(v string) *QueryApiDomainListServiceResponseDomainList {
  s.LastModified = &v
  return s
}

func (s *QueryApiDomainListServiceResponseDomainList) SetBillingArea(v string) *QueryApiDomainListServiceResponseDomainList {
  s.BillingArea = &v
  return s
}

type QueryApiDomainListServicePaths struct {
}

func (s QueryApiDomainListServicePaths) String() string {
  return tea.Prettify(s)
}

func (s QueryApiDomainListServicePaths) GoString() string {
  return s.String()
}

type QueryApiDomainListServiceParameters struct {
  // {"en":"Public CNAME alias, optional entry, does not indicate all domain names under the query account number
  // The customer has the demand that the domain cname share more than one level, so we introduce the cname-label identifier in the interface, which is a set of domain cname with the same cname-label, and share the first level of cname.", "zh_CN":"共用一级别名标示，可选入参，不选表示查询账号下所有域名
  // 客户存在较多一级域名共用的需求，因此在接口中引入cname-label标识，即拥有相同cname-label的一组域名，共用一级cname。关于cname-label的具体使用方式和注意事项，请参看【创建加速域名】和【修改域名配置】接口"}
  CnameLabel *string `json:"cname-label,omitempty" xml:"cname-label,omitempty"`
}

func (s QueryApiDomainListServiceParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryApiDomainListServiceParameters) GoString() string {
  return s.String()
}

func (s *QueryApiDomainListServiceParameters) SetCnameLabel(v string) *QueryApiDomainListServiceParameters {
  s.CnameLabel = &v
  return s
}

type QueryApiDomainListServiceRequestHeader struct {
}

func (s QueryApiDomainListServiceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryApiDomainListServiceRequestHeader) GoString() string {
  return s.String()
}

type QueryApiDomainListServiceResponseHeader struct {
  // {"en":"httpstatus=202; Indicates that the new domain API was successfully invoked, and the current deployment of the new domain can be viewed using x-cnc-request-id in the header", "zh_CN":"httpstatus=202;   表示成功调用新增域名接口，可使用header中的x-cnc-request-id查看当前新增域名的部署情况"}
  HttpStatus *int `json:"http-status-code,omitempty" xml:"http-status-code,omitempty" require:"true"`
  // {"en":"Uniquely identified id for querying tasks per request (for all API)", "zh_CN":"唯一标示的id，用于查询每次请求的任务 （适用全部接口）"}
  XCncRequestId *string `json:"x-cnc-request-id,omitempty" xml:"x-cnc-request-id,omitempty" require:"true"`
}

func (s QueryApiDomainListServiceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryApiDomainListServiceResponseHeader) GoString() string {
  return s.String()
}

func (s *QueryApiDomainListServiceResponseHeader) SetHttpStatus(v int) *QueryApiDomainListServiceResponseHeader {
  s.HttpStatus = &v
  return s
}

func (s *QueryApiDomainListServiceResponseHeader) SetXCncRequestId(v string) *QueryApiDomainListServiceResponseHeader {
  s.XCncRequestId = &v
  return s
}




type UpdateCdnDomainRequest struct {
  // {"en":"The current version is 1.0.0.", "zh_CN":"版本号"}
  Version *string `json:"version,omitempty" xml:"version,omitempty"`
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
  // {"en":"Remarks, up to 1000 characters", "zh_CN":"备注信息，最大限制1000个字符"}
  Comment *string `json:"comment,omitempty" xml:"comment,omitempty"`
  // {"en":"Cache file HOST.
  // Cache rules for caching HOST domain names and accelerated domain names must be consistent.", "zh_CN":"缓存文件HOST。缓存HOST域名和加速域名的缓存规则必须一致。"}
  CacheHost *string `json:"cache-host,omitempty" xml:"cache-host,omitempty"`
  // {"en":"Pass the response header of client IP. The optional values areCdn-Src-Ip and X-Forwarded-For. The default value isCdn-Src-Ip.", "zh_CN":"传递客户端ip的响应头部，可选值为Cdn-Src-Ip和X-Forwarded-For，默认值为Cdn-Src-Ip"}
  HeaderOfClientIp *string `json:"header-of-clientip,omitempty" xml:"header-of-clientip,omitempty"`
  // {"en":"Back-to-origin policy setting, which is used to set the origin site information and the back-to-origin policy of the accelerated domain name", "zh_CN":"回源策略设置，用于设置加速域名的源站信息和回源策略"}
  OriginConfig *UpdateCdnDomainRequestOriginConfig `json:"origin-config,omitempty" xml:"origin-config,omitempty" type:"Struct"`
  // {"en":"Live domain configuration, used to set the push and pull domain of live streaming.", "zh_CN":"直播域名配置，用于设置直播加速域名的推拉流"}
  LiveConfigView *UpdateCdnDomainRequestLiveConfigView `json:"live-config,omitempty" xml:"live-config,omitempty" type:"Struct"`
  // {"en":"SSL settings, to bind a certificate with the accelerated domain. You can use the interface [AddCertificate] to upload your certificates. If you want to modify a certificate, please use the interface: [UpdateCertificate].", "zh_CN":"ssl证书设置，用于设置加速域名的ssl证书配置。上传证书请使用接口：【新增证书V2】；若要修改证书，请使用接口：【修改证书V2】。"}
  Ssl *UpdateCdnDomainRequestSsl `json:"ssl,omitempty" xml:"ssl,omitempty" type:"Struct"`
  // {"en":"Cache policy settings are for setting cache rules for accelerated domain names", "zh_CN":"缓存规则设置，用于设置加速域名的缓存规则,请使用新接口：修改缓存时间配置接口"}
  CacheBehaviors []*UpdateCdnDomainRequestCacheBehaviors `json:"cache-behaviors,omitempty" xml:"cache-behaviors,omitempty" type:"Repeated"`
  // {"en":"Livestream domain settings. Set the publishing point of the live push-pull domain.
  // note:
  // 1. The pull stream and the corresponding push stream domain must be configured with the same publishing point.
  // 2. If you are not going to modify the publishing point, please do not pass this param.
  // 3. The publishing point adopts the overlay update. Each time you modify, you need to submit all the publishing points. You cannot submit only the parts that need to be modified.", "zh_CN":"设置直播推拉流域名的发布点
  // 注意：
  // 1、拉流和对应的推流域名，必须配置相同的发布点；
  // 2、不想修改发布点时，不要传入该节点及以下入参；
  // 3、发布点采用覆盖式更新，每次修改时，需要提交全部发布点，不能仅提交需要修改的部分。"}
  PublishPoints []*UpdateCdnDomainRequestPublishPoints `json:"publish-points,omitempty" xml:"publish-points,omitempty" type:"Repeated"`
}

func (s UpdateCdnDomainRequest) String() string {
  return tea.Prettify(s)
}

func (s UpdateCdnDomainRequest) GoString() string {
  return s.String()
}

func (s *UpdateCdnDomainRequest) SetVersion(v string) *UpdateCdnDomainRequest {
  s.Version = &v
  return s
}

func (s *UpdateCdnDomainRequest) SetServiceAreas(v string) *UpdateCdnDomainRequest {
  s.ServiceAreas = &v
  return s
}

func (s *UpdateCdnDomainRequest) SetCnameLabel(v string) *UpdateCdnDomainRequest {
  s.CnameLabel = &v
  return s
}

func (s *UpdateCdnDomainRequest) SetComment(v string) *UpdateCdnDomainRequest {
  s.Comment = &v
  return s
}

func (s *UpdateCdnDomainRequest) SetCacheHost(v string) *UpdateCdnDomainRequest {
  s.CacheHost = &v
  return s
}

func (s *UpdateCdnDomainRequest) SetHeaderOfClientIp(v string) *UpdateCdnDomainRequest {
  s.HeaderOfClientIp = &v
  return s
}

func (s *UpdateCdnDomainRequest) SetOriginConfig(v *UpdateCdnDomainRequestOriginConfig) *UpdateCdnDomainRequest {
  s.OriginConfig = v
  return s
}

func (s *UpdateCdnDomainRequest) SetLiveConfigView(v *UpdateCdnDomainRequestLiveConfigView) *UpdateCdnDomainRequest {
  s.LiveConfigView = v
  return s
}

func (s *UpdateCdnDomainRequest) SetSsl(v *UpdateCdnDomainRequestSsl) *UpdateCdnDomainRequest {
  s.Ssl = v
  return s
}

func (s *UpdateCdnDomainRequest) SetCacheBehaviors(v []*UpdateCdnDomainRequestCacheBehaviors) *UpdateCdnDomainRequest {
  s.CacheBehaviors = v
  return s
}

func (s *UpdateCdnDomainRequest) SetPublishPoints(v []*UpdateCdnDomainRequestPublishPoints) *UpdateCdnDomainRequest {
  s.PublishPoints = v
  return s
}

type UpdateCdnDomainRequestOriginConfig struct {
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
  AdvOriginConfigs *UpdateCdnDomainRequestOriginConfigAdvOriginConfigs `json:"adv-origin-configs,omitempty" xml:"adv-origin-configs,omitempty" type:"Struct"`
  // {"en":"The Origin HOST for changing the HOST field in the return source HTTP request header. It should be domain name or IP format. For domain name, each segement separated by a dot, does not exceed 62 characters, the total length should not exceed 128 characters.", "zh_CN":"回源HOST，用于更改回源HTTP请求头中的HOST字段。支持格式为域名或ip。
  // 注意：必须符合ip/域名格式规范。如果是域名，则域名每段（点号分隔）长度小于等于62，域名总长度小于等于128。"}
  DefaultOriginHostHeader *string `json:"default-origin-host-header,omitempty" xml:"default-origin-host-header,omitempty"`
  // {"en":"Origin port.", "zh_CN":"回源请求端口。"}
  OriginPort *string `json:"origin-port,omitempty" xml:"origin-port,omitempty"`
}

func (s UpdateCdnDomainRequestOriginConfig) String() string {
  return tea.Prettify(s)
}

func (s UpdateCdnDomainRequestOriginConfig) GoString() string {
  return s.String()
}

func (s *UpdateCdnDomainRequestOriginConfig) SetOriginIps(v string) *UpdateCdnDomainRequestOriginConfig {
  s.OriginIps = &v
  return s
}

func (s *UpdateCdnDomainRequestOriginConfig) SetAdvOriginConfigs(v *UpdateCdnDomainRequestOriginConfigAdvOriginConfigs) *UpdateCdnDomainRequestOriginConfig {
  s.AdvOriginConfigs = v
  return s
}

func (s *UpdateCdnDomainRequestOriginConfig) SetDefaultOriginHostHeader(v string) *UpdateCdnDomainRequestOriginConfig {
  s.DefaultOriginHostHeader = &v
  return s
}

func (s *UpdateCdnDomainRequestOriginConfig) SetOriginPort(v string) *UpdateCdnDomainRequestOriginConfig {
  s.OriginPort = &v
  return s
}

type UpdateCdnDomainRequestOriginConfigAdvOriginConfigs struct {
  // {"en":"The advanced origin monitoring url, and requests <master-ips> through the url. If the response is neither 2** nor 3** response, it is considered that the primary source ip is faulty, and <backup-ips> is used at this time.", "zh_CN":"高级源监控url，通过该url请求<master-ips>，如果返回非2**，3**响应时，认为主要回源ip故障，此时使用<backup-ips>。 需要输入完整的url，例如：http://a.example.com/test.html"}
  DetectUrl *string `json:"detect-url,omitempty" xml:"detect-url,omitempty"`
  // {"en":"Advanced source monitoring period, and the unit is in seconds, optional as an integer greater than or equal to 0, 0 means no monitoring", "zh_CN":"高级源监控周期，单位秒，可选值为大于等于0的整数，0表示不监控"}
  DetectPeriod *string `json:"detect-period,omitempty" xml:"detect-period,omitempty"`
  // {"en":"adv-origin-config", "zh_CN":"高级源配置"}
  AdvOriginConfigList []*UpdateCdnDomainRequestOriginConfigAdvOriginConfigsAdvOriginConfigList `json:"adv-origin-config,omitempty" xml:"adv-origin-config,omitempty" type:"Repeated"`
}

func (s UpdateCdnDomainRequestOriginConfigAdvOriginConfigs) String() string {
  return tea.Prettify(s)
}

func (s UpdateCdnDomainRequestOriginConfigAdvOriginConfigs) GoString() string {
  return s.String()
}

func (s *UpdateCdnDomainRequestOriginConfigAdvOriginConfigs) SetDetectUrl(v string) *UpdateCdnDomainRequestOriginConfigAdvOriginConfigs {
  s.DetectUrl = &v
  return s
}

func (s *UpdateCdnDomainRequestOriginConfigAdvOriginConfigs) SetDetectPeriod(v string) *UpdateCdnDomainRequestOriginConfigAdvOriginConfigs {
  s.DetectPeriod = &v
  return s
}

func (s *UpdateCdnDomainRequestOriginConfigAdvOriginConfigs) SetAdvOriginConfigList(v []*UpdateCdnDomainRequestOriginConfigAdvOriginConfigsAdvOriginConfigList) *UpdateCdnDomainRequestOriginConfigAdvOriginConfigs {
  s.AdvOriginConfigList = v
  return s
}

type UpdateCdnDomainRequestOriginConfigAdvOriginConfigsAdvOriginConfigList struct     {
  // {"en":"The advanced source mainly returns the source IP. Multiple IPs are separated by a semicolon ';', and the returned source IP cannot be repeated.", "zh_CN":"高级源主要回源IP，多个IP用分号&ldquo;;&rdquo;分隔，回源IP不能重复"}
  MasterIps *string `json:"master-ips,omitempty" xml:"master-ips,omitempty"`
  // {"en":"Advanced source backup source IP, multiple IPs are separated by semicolon ';', and the returned source IP cannot be duplicated.", "zh_CN":"高级源备用回源IP，多个IP用分号&ldquo;;&rdquo;分隔，回源IP不能重复"}
  BackupIps *string `json:"backup-ips,omitempty" xml:"backup-ips,omitempty"`
}

func (s UpdateCdnDomainRequestOriginConfigAdvOriginConfigsAdvOriginConfigList) String() string {
  return tea.Prettify(s)
}

func (s UpdateCdnDomainRequestOriginConfigAdvOriginConfigsAdvOriginConfigList) GoString() string {
  return s.String()
}

func (s *UpdateCdnDomainRequestOriginConfigAdvOriginConfigsAdvOriginConfigList) SetMasterIps(v string) *UpdateCdnDomainRequestOriginConfigAdvOriginConfigsAdvOriginConfigList {
  s.MasterIps = &v
  return s
}

func (s *UpdateCdnDomainRequestOriginConfigAdvOriginConfigsAdvOriginConfigList) SetBackupIps(v string) *UpdateCdnDomainRequestOriginConfigAdvOriginConfigsAdvOriginConfigList {
  s.BackupIps = &v
  return s
}

type UpdateCdnDomainRequestLiveConfigView struct {
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

func (s UpdateCdnDomainRequestLiveConfigView) String() string {
  return tea.Prettify(s)
}

func (s UpdateCdnDomainRequestLiveConfigView) GoString() string {
  return s.String()
}

func (s *UpdateCdnDomainRequestLiveConfigView) SetLiveConfigOriginIps(v string) *UpdateCdnDomainRequestLiveConfigView {
  s.LiveConfigOriginIps = &v
  return s
}

func (s *UpdateCdnDomainRequestLiveConfigView) SetOriginPushHost(v string) *UpdateCdnDomainRequestLiveConfigView {
  s.OriginPushHost = &v
  return s
}

type UpdateCdnDomainRequestSsl struct {
  // {"en":"Use a certificate, and the optional values are true and false, true means using this certificate, false means not using the certificate.", "zh_CN":"使用证书，可选值为true和false，true表示使用证书，false表示不使用证书"}
  UseSsl *string `json:"use-ssl,omitempty" xml:"use-ssl,omitempty"`
  // {"en":"Use sni certificate, and the optional values are true and false, true means using sni certificate, false means using shared certificate (not supported)", "zh_CN":"使用sni证书，可选值为true和false，true表示使用sni证书，false表示使用合用证书（暂不支持）"}
  UseForSni *string `json:"use-for-sni,omitempty" xml:"use-for-sni,omitempty"`
  // {"en":"Certificate ID, and it is the certificate ID returned by the system after the new certificate is successfully added.", "zh_CN":"证书ID，新增证书成功后，系统返回的证书ID
  // use-ssl为true时，才能传ssl-certificate-id。（点播和下载加速类型的证书需要下单给运维单独配置）"}
  SslCertificateId *string `json:"ssl-certificate-id,omitempty" xml:"ssl-certificate-id,omitempty"`
}

func (s UpdateCdnDomainRequestSsl) String() string {
  return tea.Prettify(s)
}

func (s UpdateCdnDomainRequestSsl) GoString() string {
  return s.String()
}

func (s *UpdateCdnDomainRequestSsl) SetUseSsl(v string) *UpdateCdnDomainRequestSsl {
  s.UseSsl = &v
  return s
}

func (s *UpdateCdnDomainRequestSsl) SetUseForSni(v string) *UpdateCdnDomainRequestSsl {
  s.UseForSni = &v
  return s
}

func (s *UpdateCdnDomainRequestSsl) SetSslCertificateId(v string) *UpdateCdnDomainRequestSsl {
  s.SslCertificateId = &v
  return s
}

type UpdateCdnDomainRequestCacheBehaviors struct     {
  // {"en":"The url matching mode supports fuzzy regularization. The optional values are?: /(a|b)/*.(jpg|bmp|png|gif) and other regular contents. If you want to modify the caching rules, this item is required.", "zh_CN":"url匹配模式，支持正则，如果是全部匹配，入参可以配置为：* 
  // 如果要修改缓存规则，此项必填"}
  PathPattern *string `json:"path-pattern,omitempty" xml:"path-pattern,omitempty"`
  // {"en":"Indicates the priority of a configuration, the larger the value, the higher the priority.", "zh_CN":"表示一条配置的优先级，值越大优先级越高"}
  Priority *string `json:"priority,omitempty" xml:"priority,omitempty"`
  // {"en":"Cache Time: Set the time corresponding to the cache object. Input UpdateCdnDomainParameters format: integer, set to 0 if no cache is used. There is no upper limit on the cache time rule. The time is set attuned to the customer's own needs. If the customer feels that some of the files does not change frequently, then the setting is longer. For example, the text class js, css, html, etc. can be set shorter, the picture, video and audio classes can be set longer (because the cache time will be replaced by the new file due to the file heat algorithm, the suggestion is that the length should not exceed one month)
  // If you want to modify the caching rules, this item is required.", "zh_CN":"缓存时间：设置缓存对象对应的时间。入参格式：整数，不缓存设置为0。
  // 缓存时间理论上没有上限限制，这个时间根据客户自身的需求设定，如果客户觉得其中一些文件，变更不频繁，那么就设置长一点。例如，文本类的js，css，html等可以设置得短一些，图片、视频音频类的可以设置的长一点（因为缓存时间会因文件热度算法，旧文件会被新文件替换掉，最长建议不要超过一个月）
  // 如果要修改缓存规则，此项必填"}
  CacheTtl *string `json:"cache-ttl,omitempty" xml:"cache-ttl,omitempty"`
  // {"en":"Cache Time unit: Set the time unit corresponding to the cache object: such as s, m, h, d. If no unit is entered, the default is seconds.", "zh_CN":"缓存时间单位：设置缓存对象对应的时间单位：比如s、m、h、d。不输入单位默认是秒"}
  CacheUnit *string `json:"cache-unit,omitempty" xml:"cache-unit,omitempty"`
  // {"en":"Ignore that the source station does not cache the header. The optional values are true and false, which are used to ignore the two configurations of cache-control in the request header (private, no-cache) and the authorization is set by the client.", "zh_CN":"忽略源站不缓存头。可选值为true和false，用于忽略请求头中cache-control的两种配置（private，no-cache）和客户端设置的Authorization。
  // ture表示会忽略掉源站对于这三者的设定。使得资源能够以cache-control: public的方式缓存在服务节点上，然后我们的节点才能缓存这种类型的资源，提供加速服务。
  // false表示当源站对某种资源设定了cache-control: private,cache-control:no-cache或指定根据authorization进行缓存时，我们的服务节点将不会对此类文件进行缓存。"}
  IgnoreCacheControl *string `json:"ignore-cache-control,omitempty" xml:"ignore-cache-control,omitempty"`
}

func (s UpdateCdnDomainRequestCacheBehaviors) String() string {
  return tea.Prettify(s)
}

func (s UpdateCdnDomainRequestCacheBehaviors) GoString() string {
  return s.String()
}

func (s *UpdateCdnDomainRequestCacheBehaviors) SetPathPattern(v string) *UpdateCdnDomainRequestCacheBehaviors {
  s.PathPattern = &v
  return s
}

func (s *UpdateCdnDomainRequestCacheBehaviors) SetPriority(v string) *UpdateCdnDomainRequestCacheBehaviors {
  s.Priority = &v
  return s
}

func (s *UpdateCdnDomainRequestCacheBehaviors) SetCacheTtl(v string) *UpdateCdnDomainRequestCacheBehaviors {
  s.CacheTtl = &v
  return s
}

func (s *UpdateCdnDomainRequestCacheBehaviors) SetCacheUnit(v string) *UpdateCdnDomainRequestCacheBehaviors {
  s.CacheUnit = &v
  return s
}

func (s *UpdateCdnDomainRequestCacheBehaviors) SetIgnoreCacheControl(v string) *UpdateCdnDomainRequestCacheBehaviors {
  s.IgnoreCacheControl = &v
  return s
}

type UpdateCdnDomainRequestPublishPoints struct     {
  // {"en":"Livestream domain settings. Publish point, support multiple, do not pass the system by default to generate a publishing point uri for [/]", "zh_CN":"发布点，支持多个，不传系统默认生成一条发布点uri为“/”"}
  Uri *string `json:"uri,omitempty" xml:"uri,omitempty"`
}

func (s UpdateCdnDomainRequestPublishPoints) String() string {
  return tea.Prettify(s)
}

func (s UpdateCdnDomainRequestPublishPoints) GoString() string {
  return s.String()
}

func (s *UpdateCdnDomainRequestPublishPoints) SetUri(v string) *UpdateCdnDomainRequestPublishPoints {
  s.Uri = &v
  return s
}

type UpdateCdnDomainResponse struct {
  // {"en":"Error code occurs when HTTP Status is not 202, indicating the error type of the current request calls.", "zh_CN":"Error code occurs when HTTPStatus is not 202, indicating the error type of the current request calls."}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response information, demonstrates success when it is successful.", "zh_CN":"Response information, demonstrates success when it is successful."}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s UpdateCdnDomainResponse) String() string {
  return tea.Prettify(s)
}

func (s UpdateCdnDomainResponse) GoString() string {
  return s.String()
}

func (s *UpdateCdnDomainResponse) SetCode(v string) *UpdateCdnDomainResponse {
  s.Code = &v
  return s
}

func (s *UpdateCdnDomainResponse) SetMessage(v string) *UpdateCdnDomainResponse {
  s.Message = &v
  return s
}

type UpdateCdnDomainPaths struct {
  // {"en":"The domain you are going to modify, it can be domain id or domain name.", "zh_CN":"需要修改的域名，可以是域名名称或域名id。"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
}

func (s UpdateCdnDomainPaths) String() string {
  return tea.Prettify(s)
}

func (s UpdateCdnDomainPaths) GoString() string {
  return s.String()
}

func (s *UpdateCdnDomainPaths) SetDomain(v string) *UpdateCdnDomainPaths {
  s.Domain = &v
  return s
}

type UpdateCdnDomainParameters struct {
}

func (s UpdateCdnDomainParameters) String() string {
  return tea.Prettify(s)
}

func (s UpdateCdnDomainParameters) GoString() string {
  return s.String()
}

type UpdateCdnDomainRequestHeader struct {
}

func (s UpdateCdnDomainRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s UpdateCdnDomainRequestHeader) GoString() string {
  return s.String()
}

type UpdateCdnDomainResponseHeader struct {
}

func (s UpdateCdnDomainResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s UpdateCdnDomainResponseHeader) GoString() string {
  return s.String()
}




type GetPagingDomainListRequest struct {
  // {"en":"Page number must be a positive integer greater than 0.If not passed, then no paging. If it is passed, pageSize is required.", "zh_CN":"分页的页码，必须为大于0的正整数。不传默认不分页，若传参则pageSize必填.。"}
  PageNumber *int `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
  // {"en":"Number of domain name data items for paging, must be a positive integer greater than 0.If not passed, then no paging. If it is passed, pageSize is required.", "zh_CN":"分页的域名数据条数，必须大于0的正整数。不传默认不分页，若传参则pageSize必填.。"}
  PageSize *int `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
  // {"en":"Specify the service type to be queried. Multiple services are allowed. Data will be returned if any one service is satisfied. If not passed, all services will be checked by default. For example: [wsa,waf], returns all domains whose services include wsa or include waf.", "zh_CN":"指定查询的服务，允许多个服务，任意一个服务满足就返回数据，不传默认查全部服务。如：[wsa,waf], 则返回服务包含wsa或包含waf的所有域名。"}
  ServiceTypes []*string `json:"serviceTypes,omitempty" xml:"serviceTypes,omitempty" type:"Repeated"`
  // {"en":"Specify the accelerated domain name for the query. Multiple domain names are allowed. If not specified, all domain names will be searched by default.", "zh_CN":"指定查询的加速域名，允许多个域名，不传默认查全部域名。"}
  DomainNames []*string `json:"domainNames,omitempty" xml:"domainNames,omitempty" type:"Repeated"`
  // {"en":"RFC3339 formatted date indicating the starting date. Example: 2024-01-01T22:30:00+08:00", "zh_CN":"	查询开始时间，支持时间格式如：2024-01-01T22:30:00+08:00"}
  StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
  // {"en":"RFC3339 formatted date indicating the ending date. Example: 2024-01-01T22:30:00+08:00", "zh_CN":"查询结束时间，支持时间格式如：2024-01-01T22:30:00+08:00"}
  EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
  // {"en":" Status of the accelerated domain. Optional value: enabled, disabled, deploying, checking, disabling, deployFailed, disableFailed.", "zh_CN":"加速域名的状态：enabled表示已启用；disabled表示已禁用；deploying表示配置部署中；checking表示审核中；disabling表示禁用中；deployFailed表示配置部署失败；disableFailed表示禁用失败。不传默认查全部"}
  Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s GetPagingDomainListRequest) String() string {
  return tea.Prettify(s)
}

func (s GetPagingDomainListRequest) GoString() string {
  return s.String()
}

func (s *GetPagingDomainListRequest) SetPageNumber(v int) *GetPagingDomainListRequest {
  s.PageNumber = &v
  return s
}

func (s *GetPagingDomainListRequest) SetPageSize(v int) *GetPagingDomainListRequest {
  s.PageSize = &v
  return s
}

func (s *GetPagingDomainListRequest) SetServiceTypes(v []*string) *GetPagingDomainListRequest {
  s.ServiceTypes = v
  return s
}

func (s *GetPagingDomainListRequest) SetDomainNames(v []*string) *GetPagingDomainListRequest {
  s.DomainNames = v
  return s
}

func (s *GetPagingDomainListRequest) SetStartTime(v string) *GetPagingDomainListRequest {
  s.StartTime = &v
  return s
}

func (s *GetPagingDomainListRequest) SetEndTime(v string) *GetPagingDomainListRequest {
  s.EndTime = &v
  return s
}

func (s *GetPagingDomainListRequest) SetStatus(v string) *GetPagingDomainListRequest {
  s.Status = &v
  return s
}

type GetPagingDomainListResponse struct {
  // {"en":"Responses the page number of the data", "zh_CN":"所有满足条件的数据总条数"}
  TotalCount *int `json:"totalCount,omitempty" xml:"totalCount,omitempty" require:"true"`
  // {"en":"total pages", "zh_CN":"总页数"}
  TotalPageNumber *int `json:"totalPageNumber,omitempty" xml:"totalPageNumber,omitempty" require:"true"`
  // {"en":"Responses the page number of the data", "zh_CN":"返回数据的页码"}
  PageNumber *int `json:"pageNumber,omitempty" xml:"pageNumber,omitempty" require:"true"`
  // {"en":"Number of data page", "zh_CN":"每个页面的数据条数"}
  PageSize *int `json:"pageSize,omitempty" xml:"pageSize,omitempty" require:"true"`
  // {"en":"Domain list.", "zh_CN":"域名列表"}
  ResultList []*GetPagingDomainListResponseResultList `json:"resultList,omitempty" xml:"resultList,omitempty" require:"true" type:"Repeated"`
}

func (s GetPagingDomainListResponse) String() string {
  return tea.Prettify(s)
}

func (s GetPagingDomainListResponse) GoString() string {
  return s.String()
}

func (s *GetPagingDomainListResponse) SetTotalCount(v int) *GetPagingDomainListResponse {
  s.TotalCount = &v
  return s
}

func (s *GetPagingDomainListResponse) SetTotalPageNumber(v int) *GetPagingDomainListResponse {
  s.TotalPageNumber = &v
  return s
}

func (s *GetPagingDomainListResponse) SetPageNumber(v int) *GetPagingDomainListResponse {
  s.PageNumber = &v
  return s
}

func (s *GetPagingDomainListResponse) SetPageSize(v int) *GetPagingDomainListResponse {
  s.PageSize = &v
  return s
}

func (s *GetPagingDomainListResponse) SetResultList(v []*GetPagingDomainListResponseResultList) *GetPagingDomainListResponse {
  s.ResultList = v
  return s
}

type GetPagingDomainListResponseResultList struct     {
  // {"en":"Cname of the accelerated domain", "zh_CN":"加速域名cname，如：a1.example.com.wscdns.com"}
  Cname *string `json:"cname,omitempty" xml:"cname,omitempty" require:"true"`
  // {"en":"Create time of the accelerated domain. Example: 2024-01-01T22:30:00+08:00", "zh_CN":"	域名创建时间，时间格式如：2024-01-01T22:30:00+08:00"}
  CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty" require:"true"`
  // {"en":"Corresponding domain ID", "zh_CN":"对应的域名ID"}
  DomainId *string `json:"domainId,omitempty" xml:"domainId,omitempty" require:"true"`
  // {"en":"Accelerated domain name", "zh_CN":"加速域名名称"}
  DomainName *string `json:"domainName,omitempty" xml:"domainName,omitempty" require:"true"`
  // {"en":"Service type for accelerated domain name", "zh_CN":"加速域名的服务，如[wsa,waf]。"}
  ServiceTypes []*string `json:"serviceTypes,omitempty" xml:"serviceTypes,omitempty" require:"true" type:"Repeated"`
  // {"en":"Status of the accelerated domain. Optional value: enabled, disabled, deploying, checking, disabling.", "zh_CN":"	加速域名的状态：enabled表示已启用；disabled表示已禁用；deploying表示部署中；checking表示审核中；disabling表示禁用中；deployFailed表示配置部署失败；disableFailed表示禁用失败"}
  Status *string `json:"status,omitempty" xml:"status,omitempty" require:"true"`
  // {"en":"Accelerated domain enabling status: true indicates that it is enabled, false indicates that it is disabled.", "zh_CN":"加速域名启用状态：true为启用，false为禁用。"}
  Enabled *string `json:"enabled,omitempty" xml:"enabled,omitempty" require:"true"`
}

func (s GetPagingDomainListResponseResultList) String() string {
  return tea.Prettify(s)
}

func (s GetPagingDomainListResponseResultList) GoString() string {
  return s.String()
}

func (s *GetPagingDomainListResponseResultList) SetCname(v string) *GetPagingDomainListResponseResultList {
  s.Cname = &v
  return s
}

func (s *GetPagingDomainListResponseResultList) SetCreateTime(v string) *GetPagingDomainListResponseResultList {
  s.CreateTime = &v
  return s
}

func (s *GetPagingDomainListResponseResultList) SetDomainId(v string) *GetPagingDomainListResponseResultList {
  s.DomainId = &v
  return s
}

func (s *GetPagingDomainListResponseResultList) SetDomainName(v string) *GetPagingDomainListResponseResultList {
  s.DomainName = &v
  return s
}

func (s *GetPagingDomainListResponseResultList) SetServiceTypes(v []*string) *GetPagingDomainListResponseResultList {
  s.ServiceTypes = v
  return s
}

func (s *GetPagingDomainListResponseResultList) SetStatus(v string) *GetPagingDomainListResponseResultList {
  s.Status = &v
  return s
}

func (s *GetPagingDomainListResponseResultList) SetEnabled(v string) *GetPagingDomainListResponseResultList {
  s.Enabled = &v
  return s
}

type GetPagingDomainListPaths struct {
}

func (s GetPagingDomainListPaths) String() string {
  return tea.Prettify(s)
}

func (s GetPagingDomainListPaths) GoString() string {
  return s.String()
}

type GetPagingDomainListParameters struct {
}

func (s GetPagingDomainListParameters) String() string {
  return tea.Prettify(s)
}

func (s GetPagingDomainListParameters) GoString() string {
  return s.String()
}

type GetPagingDomainListRequestHeader struct {
}

func (s GetPagingDomainListRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetPagingDomainListRequestHeader) GoString() string {
  return s.String()
}

type GetPagingDomainListResponseHeader struct {
  // {"en":"httpstatus=202; Indicates that the new domain API was successfully invoked, and the current deployment of the new domain can be viewed using x-cnc-request-id in the header", "zh_CN":"httpstatus=202;   表示成功调用新增域名接口，可使用header中的x-cnc-request-id查看当前新增域名的部署情况"}
  Code *int `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Uniquely identified id for querying tasks per request (for all API)", "zh_CN":"唯一标示的id，用于查询每次请求的任务 （适用全部接口）"}
  XCncRequestId *string `json:"x-cnc-request-id,omitempty" xml:"x-cnc-request-id,omitempty" require:"true"`
}

func (s GetPagingDomainListResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetPagingDomainListResponseHeader) GoString() string {
  return s.String()
}

func (s *GetPagingDomainListResponseHeader) SetCode(v int) *GetPagingDomainListResponseHeader {
  s.Code = &v
  return s
}

func (s *GetPagingDomainListResponseHeader) SetXCncRequestId(v string) *GetPagingDomainListResponseHeader {
  s.XCncRequestId = &v
  return s
}




type QueryApiDomainListRequest struct {
}

func (s QueryApiDomainListRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryApiDomainListRequest) GoString() string {
  return s.String()
}

type QueryApiDomainListResponse struct {
  // {"en":"", "zh_CN":"httpstatus=202; 表示成功调用新增域名接口"}
  Code *int `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"", "zh_CN":"唯一标示的id，用于查询每次请求的任务 （适用全部接口）"}
  XCncRequestId *string `json:"x-cnc-request-id,omitempty" xml:"x-cnc-request-id,omitempty" require:"true"`
  DomainSummary []*QueryApiDomainListResponseDomainSummary `json:"domain-summary,omitempty" xml:"domain-summary,omitempty" require:"true" type:"Repeated"`
}

func (s QueryApiDomainListResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryApiDomainListResponse) GoString() string {
  return s.String()
}

func (s *QueryApiDomainListResponse) SetCode(v int) *QueryApiDomainListResponse {
  s.Code = &v
  return s
}

func (s *QueryApiDomainListResponse) SetXCncRequestId(v string) *QueryApiDomainListResponse {
  s.XCncRequestId = &v
  return s
}

func (s *QueryApiDomainListResponse) SetDomainSummary(v []*QueryApiDomainListResponseDomainSummary) *QueryApiDomainListResponse {
  s.DomainSummary = v
  return s
}

type QueryApiDomainListResponseDomainSummary struct     {
  // {"en":"", "zh_CN":"加速域名对应的CNAME域名，例如：7nt6mrh7sdkslj.cdn30.com"}
  Cname *string `json:"cname,omitempty" xml:"cname,omitempty" require:"true"`
  // {"en":"", "zh_CN":"对应的域名ID"}
  DomainId *int `json:"domain-id,omitempty" xml:"domain-id,omitempty" require:"true"`
  // {"en":"", "zh_CN":"加速域名名称"}
  DomainName *string `json:"domain-name,omitempty" xml:"domain-name,omitempty" require:"true"`
  // {"en":"", "zh_CN":"加速域名的回源IP"}
  OriginIps *string `json:"origin-ips,omitempty" xml:"origin-ips,omitempty" require:"true"`
  // {"en":"", "zh_CN":"加速域名的服务类型"}
  ServiceType *string `json:"service-type,omitempty" xml:"service-type,omitempty" require:"true"`
  // {"en":"", "zh_CN":"加速域名的部署状态"}
  Status *string `json:"status,omitempty" xml:"status,omitempty" require:"true"`
  // {"en":"", "zh_CN":"加速域名是否启用，true和false"}
  Enabled *string `json:"enabled,omitempty" xml:"enabled,omitempty" require:"true"`
}

func (s QueryApiDomainListResponseDomainSummary) String() string {
  return tea.Prettify(s)
}

func (s QueryApiDomainListResponseDomainSummary) GoString() string {
  return s.String()
}

func (s *QueryApiDomainListResponseDomainSummary) SetCname(v string) *QueryApiDomainListResponseDomainSummary {
  s.Cname = &v
  return s
}

func (s *QueryApiDomainListResponseDomainSummary) SetDomainId(v int) *QueryApiDomainListResponseDomainSummary {
  s.DomainId = &v
  return s
}

func (s *QueryApiDomainListResponseDomainSummary) SetDomainName(v string) *QueryApiDomainListResponseDomainSummary {
  s.DomainName = &v
  return s
}

func (s *QueryApiDomainListResponseDomainSummary) SetOriginIps(v string) *QueryApiDomainListResponseDomainSummary {
  s.OriginIps = &v
  return s
}

func (s *QueryApiDomainListResponseDomainSummary) SetServiceType(v string) *QueryApiDomainListResponseDomainSummary {
  s.ServiceType = &v
  return s
}

func (s *QueryApiDomainListResponseDomainSummary) SetStatus(v string) *QueryApiDomainListResponseDomainSummary {
  s.Status = &v
  return s
}

func (s *QueryApiDomainListResponseDomainSummary) SetEnabled(v string) *QueryApiDomainListResponseDomainSummary {
  s.Enabled = &v
  return s
}

type QueryApiDomainListPaths struct {
}

func (s QueryApiDomainListPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryApiDomainListPaths) GoString() string {
  return s.String()
}

type QueryApiDomainListParameters struct {
}

func (s QueryApiDomainListParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryApiDomainListParameters) GoString() string {
  return s.String()
}

type QueryApiDomainListRequestHeader struct {
}

func (s QueryApiDomainListRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryApiDomainListRequestHeader) GoString() string {
  return s.String()
}

type QueryApiDomainListResponseHeader struct {
}

func (s QueryApiDomainListResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryApiDomainListResponseHeader) GoString() string {
  return s.String()
}




type QueryChangeServerRequest struct {
}

func (s QueryChangeServerRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryChangeServerRequest) GoString() string {
  return s.String()
}

type QueryChangeServerResponse struct {
  // {"en":"The error code, when HTTPStatus is not 200, indicates the type of error the current request is calling.", "zh_CN":"错误代码，当HTTPStatus不为200时出现，表示当前请求调用的错误类型"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response information, when success is successful", "zh_CN":"响应信息，成功时为success"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"The response data", "zh_CN":"响应数据"}
  Data *QueryChangeServerResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
  // {"en":"Uniquely labeled id for querying each requested task (for all interfaces)", "zh_CN":"唯一标示的id，用于查询每次请求的任务 （适用全部接口）"}
  XCncRequestId *string `json:"x-cnc-request-id,omitempty" xml:"x-cnc-request-id,omitempty" require:"true"`
}

func (s QueryChangeServerResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryChangeServerResponse) GoString() string {
  return s.String()
}

func (s *QueryChangeServerResponse) SetCode(v string) *QueryChangeServerResponse {
  s.Code = &v
  return s
}

func (s *QueryChangeServerResponse) SetMessage(v string) *QueryChangeServerResponse {
  s.Message = &v
  return s
}

func (s *QueryChangeServerResponse) SetData(v *QueryChangeServerResponseData) *QueryChangeServerResponse {
  s.Data = v
  return s
}

func (s *QueryChangeServerResponse) SetXCncRequestId(v string) *QueryChangeServerResponse {
  s.XCncRequestId = &v
  return s
}

type QueryChangeServerResponseData struct {
  // {"en":"domain id.", "zh_CN":"域名id"}
  DomainId *string `json:"domainId,omitempty" xml:"domainId,omitempty" require:"true"`
  // {"en":"domain name.", "zh_CN":"域名名称"}
  DomainName *string `json:"domainName,omitempty" xml:"domainName,omitempty" require:"true"`
  // {"en":"Change servers configuration, parent tag
  // 1. This must be filled when the hotlinking configuration of streaming media needs to be set
  // 2. Empty the configuration for <change-servers/>", "zh_CN":"【接入域名跳转】
  // 注意：
  // 1、需要取消【接入域名跳转】时，可以传入空节点<change-servers></change-servers>。
  // 2、表示需要设置【接入域名跳转】，此项必填"}
  ChangeServers []*QueryChangeServerResponseDataChangeServers `json:"change-servers,omitempty" xml:"change-servers,omitempty" require:"true" type:"Repeated"`
}

func (s QueryChangeServerResponseData) String() string {
  return tea.Prettify(s)
}

func (s QueryChangeServerResponseData) GoString() string {
  return s.String()
}

func (s *QueryChangeServerResponseData) SetDomainId(v string) *QueryChangeServerResponseData {
  s.DomainId = &v
  return s
}

func (s *QueryChangeServerResponseData) SetDomainName(v string) *QueryChangeServerResponseData {
  s.DomainName = &v
  return s
}

func (s *QueryChangeServerResponseData) SetChangeServers(v []*QueryChangeServerResponseDataChangeServers) *QueryChangeServerResponseData {
  s.ChangeServers = v
  return s
}

type QueryChangeServerResponseDataChangeServers struct     {
  // {"en":"If it is a universal domain name, set it to a universal domain name, for example, *.56.com.", "zh_CN":"如果是泛域名，需要填写为泛域名，例如：*.56.com"}
  TargetServer *string `json:"change-server,omitempty" xml:"change-server,omitempty" require:"true"`
  // {"en":"Data-id is to indicate a specific group configuration when the client has multiple groups of configurations. Data-id can be retrieved through a query interface. Note: A. If data-id is passed, it means that one group of configuration items is specified to be modified, and no other group configuration items need to be modified. B. If multiple groups of configurations are included, some of them are configured with data-id and others are not, then the expression of data-id is used to modify a specific group of configurations, and a new group of configurations is added on the original basis without the expression of data-id. C. If the data-id is not transmitted, it means that the original configuration will be fully covered by this configuration. D. If no configuration parameter is passed, only domain name and secondary label are passed, which means that all configuration of domain name secondary service corresponding to this interface is cleared. E. If there is no specific configuration item in a set of configurations, the data-id must be filled in, and the value is the actual data-id, which means clearing the value of the corresponding data-id configuration item; it is not allowed that there is no specific configuration item or data-id in a set of configurations.", "zh_CN":"配置多组配置时，具体某组配置的id。dataId可以通过查询接口获取。 注意： a、如果有传dataId，说明指定修改其中一组配置项内容，不需求修改其他组配置内容不需要入参； b、如果入参多组配置，其中有些组配置有传dataId，有些没有传，则有传dataId的表示修改具体某组配置，没有传dataId的表示在原来基础上新增一组配置； c、如果入参都没有传dataId,表示用本次的配置全量覆盖原先配置； d、如果入参没有传任何配置项参数，只传了域名和二级标签，表示清空这个接口对应域名二级服务所有配置； e、如果一组配置没有具体的配置项，则dataId必填，且值为实际存在的dataId，表示清空这个dataId对应配置项的值；不允许一组配置没有具体的配置项也没有dataId。"}
  DataId *int `json:"dataId,omitempty" xml:"dataId,omitempty" require:"true"`
}

func (s QueryChangeServerResponseDataChangeServers) String() string {
  return tea.Prettify(s)
}

func (s QueryChangeServerResponseDataChangeServers) GoString() string {
  return s.String()
}

func (s *QueryChangeServerResponseDataChangeServers) SetTargetServer(v string) *QueryChangeServerResponseDataChangeServers {
  s.TargetServer = &v
  return s
}

func (s *QueryChangeServerResponseDataChangeServers) SetDataId(v int) *QueryChangeServerResponseDataChangeServers {
  s.DataId = &v
  return s
}

type QueryChangeServerPaths struct {
  // {"en":"Domain name or domain name id to query configuration", "zh_CN":"需要查询配置的域名（domainName）或域名id（domainId）"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
}

func (s QueryChangeServerPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryChangeServerPaths) GoString() string {
  return s.String()
}

func (s *QueryChangeServerPaths) SetDomain(v string) *QueryChangeServerPaths {
  s.Domain = &v
  return s
}

type QueryChangeServerParameters struct {
}

func (s QueryChangeServerParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryChangeServerParameters) GoString() string {
  return s.String()
}

type QueryChangeServerRequestHeader struct {
}

func (s QueryChangeServerRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryChangeServerRequestHeader) GoString() string {
  return s.String()
}

type QueryChangeServerResponseHeader struct {
}

func (s QueryChangeServerResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryChangeServerResponseHeader) GoString() string {
  return s.String()
}




