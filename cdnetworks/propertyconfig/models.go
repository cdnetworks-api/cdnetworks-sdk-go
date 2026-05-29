package propertyconfig

// This file is auto-generated, don't edit it. Thanks.
import (
  "github.com/alibabacloud-go/tea/tea"
)

type CreatePropertyForAkamaiMigrationRequest struct {
  // {"en":"Product Service Type related to your contract.","zh_CN":"产品服务类型。请根据您的合同产品服务类型填写。","exampleValue":"wsa,wsa-https"}
  ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty"`
  // {"en":"Property description.The length must not exceed 256 characters.","zh_CN":"项目的描述。长度不超过256个字符。"}
  PropertyComment *string `json:"propertyComment,omitempty" xml:"propertyComment,omitempty"`
  // {"en":"Property name. The length must not exceed 128 characters.","zh_CN":"项目的名称。长度不超过128个字符。"}
  PropertyName *string `json:"propertyName,omitempty" xml:"propertyName,omitempty" require:"true"`
  // {"en":"Property version description.The length must not exceed 256 characters.","zh_CN":"项目版本的描述。长度不超过256个字符。"}
  VersionComment *string `json:"versionComment,omitempty" xml:"versionComment,omitempty"`
  // {"en":"Akamai Property's RuleTree JSON.","zh_CN":"Akamai项目规则树JSON配置。"}
  AkProperty *string `json:"akProperty,omitempty" xml:"akProperty,omitempty" require:"true"`
  // {"en":"hostnames","zh_CN":"域名列表"}
  Hostnames []*CreatePropertyForAkamaiMigrationRequestHostnames `json:"hostnames,omitempty" xml:"hostnames,omitempty" type:"Repeated"`
  // {"en":"The id of contract, such as 40015677","zh_CN":"合同号，如40015677"}
  ContractId *string `json:"contractId,omitempty" xml:"contractId,omitempty"`
  // {"en":"The id of product, such as 10","zh_CN":"产品号，如10"}
  ItemId *string `json:"itemId,omitempty" xml:"itemId,omitempty"`
  // {"en":"Add new version to this propertyId","zh_CN":"新增版本到已有的propertyId"}
  PropertyId *int `json:"propertyId,omitempty" xml:"propertyId,omitempty"`
  // {"en":"Is filter config by hostname. Default value is '0'.","zh_CN":"是否按域名过滤配置。默认值为'0'","exampleValue":"0,1"}
  FilterConfigByHostname *string `json:"filterConfigByHostname,omitempty" xml:"filterConfigByHostname,omitempty"`
}

func (s CreatePropertyForAkamaiMigrationRequest) String() string {
  return tea.Prettify(s)
}

func (s CreatePropertyForAkamaiMigrationRequest) GoString() string {
  return s.String()
}

func (s *CreatePropertyForAkamaiMigrationRequest) SetServiceType(v string) *CreatePropertyForAkamaiMigrationRequest {
  s.ServiceType = &v
  return s
}

func (s *CreatePropertyForAkamaiMigrationRequest) SetPropertyComment(v string) *CreatePropertyForAkamaiMigrationRequest {
  s.PropertyComment = &v
  return s
}

func (s *CreatePropertyForAkamaiMigrationRequest) SetPropertyName(v string) *CreatePropertyForAkamaiMigrationRequest {
  s.PropertyName = &v
  return s
}

func (s *CreatePropertyForAkamaiMigrationRequest) SetVersionComment(v string) *CreatePropertyForAkamaiMigrationRequest {
  s.VersionComment = &v
  return s
}

func (s *CreatePropertyForAkamaiMigrationRequest) SetAkProperty(v string) *CreatePropertyForAkamaiMigrationRequest {
  s.AkProperty = &v
  return s
}

func (s *CreatePropertyForAkamaiMigrationRequest) SetHostnames(v []*CreatePropertyForAkamaiMigrationRequestHostnames) *CreatePropertyForAkamaiMigrationRequest {
  s.Hostnames = v
  return s
}

func (s *CreatePropertyForAkamaiMigrationRequest) SetContractId(v string) *CreatePropertyForAkamaiMigrationRequest {
  s.ContractId = &v
  return s
}

func (s *CreatePropertyForAkamaiMigrationRequest) SetItemId(v string) *CreatePropertyForAkamaiMigrationRequest {
  s.ItemId = &v
  return s
}

func (s *CreatePropertyForAkamaiMigrationRequest) SetPropertyId(v int) *CreatePropertyForAkamaiMigrationRequest {
  s.PropertyId = &v
  return s
}

func (s *CreatePropertyForAkamaiMigrationRequest) SetFilterConfigByHostname(v string) *CreatePropertyForAkamaiMigrationRequest {
  s.FilterConfigByHostname = &v
  return s
}

type CreatePropertyForAkamaiMigrationRequestHostnames struct     {
  // {"en":"hostnames. The length must not exceed 128 characters. A wildcard hostname must start with an asterisk (*). Multiple hostnames are separated by English semicolons(;).","zh_CN":"加速域名。长度不超过128个字符。泛域名需要以“*”开头，例如：`*.example.com`。多个域名以英文分号(;)分隔。"}
  Hostname *string `json:"hostname,omitempty" xml:"hostname,omitempty" require:"true"`
  // {"en":"Associated edge certificate configurations.","zh_CN":"关联的边缘证书"}
  Certificates []*CreatePropertyForAkamaiMigrationRequestHostnamesCertificates `json:"certificates,omitempty" xml:"certificates,omitempty" type:"Repeated"`
  // {"en":"Edge Hostname","zh_CN":"调度域名"}
  EdgeHostname *CreatePropertyForAkamaiMigrationRequestHostnamesEdgeHostname `json:"edgeHostname,omitempty" xml:"edgeHostname,omitempty" type:"Struct"`
  // {"en":"default origin","zh_CN":"默认源站"}
  DefaultOrigin *CreatePropertyForAkamaiMigrationRequestHostnamesDefaultOrigin `json:"defaultOrigin,omitempty" xml:"defaultOrigin,omitempty" type:"Struct"`
}

func (s CreatePropertyForAkamaiMigrationRequestHostnames) String() string {
  return tea.Prettify(s)
}

func (s CreatePropertyForAkamaiMigrationRequestHostnames) GoString() string {
  return s.String()
}

func (s *CreatePropertyForAkamaiMigrationRequestHostnames) SetHostname(v string) *CreatePropertyForAkamaiMigrationRequestHostnames {
  s.Hostname = &v
  return s
}

func (s *CreatePropertyForAkamaiMigrationRequestHostnames) SetCertificates(v []*CreatePropertyForAkamaiMigrationRequestHostnamesCertificates) *CreatePropertyForAkamaiMigrationRequestHostnames {
  s.Certificates = v
  return s
}

func (s *CreatePropertyForAkamaiMigrationRequestHostnames) SetEdgeHostname(v *CreatePropertyForAkamaiMigrationRequestHostnamesEdgeHostname) *CreatePropertyForAkamaiMigrationRequestHostnames {
  s.EdgeHostname = v
  return s
}

func (s *CreatePropertyForAkamaiMigrationRequestHostnames) SetDefaultOrigin(v *CreatePropertyForAkamaiMigrationRequestHostnamesDefaultOrigin) *CreatePropertyForAkamaiMigrationRequestHostnames {
  s.DefaultOrigin = v
  return s
}

type CreatePropertyForAkamaiMigrationRequestHostnamesCertificates struct     {
  // {"en":"Certificate ID","zh_CN":"证书ID"}
  CertificateId *int `json:"certificateId,omitempty" xml:"certificateId,omitempty" require:"true"`
  // {"en":"certificate usage. \n- default_sni: SNI Edge Certificate;\n- dual_sni: Additional SNI Edge Certificate;\n- gm_sm2_enc: Additional SM2 Encryption Certificate;\n- gm_sm2_sign: Additional SM2 Signature Certificate;\n- client_mtls: mTLS Client CA Certificate;","zh_CN":"证书用途。\n- default_sni：表示SNI边缘证书；\n- dual_sni：表示多栈SNI边缘证书；\n- gm_sm2_enc：表示国密加密证书；\n- gm_sm2_sign：表示国密签名证书；\n- client_mtls：表示mTLS客户端CA证书。","exampleValue":"default_sni,dual_sni,gm_sm2_enc,gm_sm2_sign,client_mtls"}
  CertificateUsage *string `json:"certificateUsage,omitempty" xml:"certificateUsage,omitempty" require:"true"`
}

func (s CreatePropertyForAkamaiMigrationRequestHostnamesCertificates) String() string {
  return tea.Prettify(s)
}

func (s CreatePropertyForAkamaiMigrationRequestHostnamesCertificates) GoString() string {
  return s.String()
}

func (s *CreatePropertyForAkamaiMigrationRequestHostnamesCertificates) SetCertificateId(v int) *CreatePropertyForAkamaiMigrationRequestHostnamesCertificates {
  s.CertificateId = &v
  return s
}

func (s *CreatePropertyForAkamaiMigrationRequestHostnamesCertificates) SetCertificateUsage(v string) *CreatePropertyForAkamaiMigrationRequestHostnamesCertificates {
  s.CertificateUsage = &v
  return s
}

type CreatePropertyForAkamaiMigrationRequestHostnamesEdgeHostname struct {
  // {"en":"Edge Hostname prefix","zh_CN":"调度域名前缀"}
  EdgeHostnamePrefix *string `json:"edgeHostnamePrefix,omitempty" xml:"edgeHostnamePrefix,omitempty" require:"true"`
  // {"en":"Edge Hostname suffix","zh_CN":"调度域名后缀"}
  EdgeHostnameSuffix *string `json:"edgeHostnameSuffix,omitempty" xml:"edgeHostnameSuffix,omitempty" require:"true"`
  // {"en":"Edge Hostname description","zh_CN":"调度域名描述"}
  Comment *string `json:"comment,omitempty" xml:"comment,omitempty"`
}

func (s CreatePropertyForAkamaiMigrationRequestHostnamesEdgeHostname) String() string {
  return tea.Prettify(s)
}

func (s CreatePropertyForAkamaiMigrationRequestHostnamesEdgeHostname) GoString() string {
  return s.String()
}

func (s *CreatePropertyForAkamaiMigrationRequestHostnamesEdgeHostname) SetEdgeHostnamePrefix(v string) *CreatePropertyForAkamaiMigrationRequestHostnamesEdgeHostname {
  s.EdgeHostnamePrefix = &v
  return s
}

func (s *CreatePropertyForAkamaiMigrationRequestHostnamesEdgeHostname) SetEdgeHostnameSuffix(v string) *CreatePropertyForAkamaiMigrationRequestHostnamesEdgeHostname {
  s.EdgeHostnameSuffix = &v
  return s
}

func (s *CreatePropertyForAkamaiMigrationRequestHostnamesEdgeHostname) SetComment(v string) *CreatePropertyForAkamaiMigrationRequestHostnamesEdgeHostname {
  s.Comment = &v
  return s
}

type CreatePropertyForAkamaiMigrationRequestHostnamesDefaultOrigin struct {
  // {"en":"origin servers","zh_CN":"源站服务器"}
  Servers []*string `json:"servers,omitempty" xml:"servers,omitempty" type:"Repeated"`
  // {"en":"origin host","zh_CN":"回源HOST"}
  Host *string `json:"host,omitempty" xml:"host,omitempty"`
  // {"en":"http port","zh_CN":"http端口"}
  HttpPort *int `json:"httpPort,omitempty" xml:"httpPort,omitempty"`
  // {"en":"https port","zh_CN":"https端口"}
  HttpsPort *int `json:"httpsPort,omitempty" xml:"httpsPort,omitempty"`
  // {"en":"IP version","zh_CN":"IP版本","exampleValue":"dual,ipv4,ipv6"}
  IpVersion *string `json:"ipVersion,omitempty" xml:"ipVersion,omitempty"`
  // {"en":"Scheme used for origin requests. Options include:- http: Fixedly use HTTP.- https: Fixedly use HTTPS.- honor: default option.Follow the scheme of the client request (default).","zh_CN":"回源请求HTTP/HTTPS协议。可选项如下： - http：固定使用HTTP协议回源。- https：固定使用HTTPS协议回源。- honor：默认选项。根据客户端请求的协议类型回源。即客户端请求使用HTTP则用HTTP回源；客户端请求使用HTTPS则用HTTPS回源。","exampleValue":"http,https,honor"}
  Scheme *string `json:"scheme,omitempty" xml:"scheme,omitempty"`
  // {"en":"Whether to enable SNI. Options include:- true: The origin TLS handshake will carry SNI (Server Name Indication).Note: Before enabling SNI, please upload the certificate to the origin and set the back-to-origin scheme to https.- false: The origin TLS handshake will not carry SNI (Server Name Indication)","zh_CN":"是否启用SNI。可选项如下：- true：回源TLS握手将携带SNI（Server Name Indication）。注意：启用SNI前，请为源站配置证书并设置回源scheme为https。- false：回源TLS握手不携带SNI（Server Name Indication）。"}
  ProxySSLSNIEnabled *bool `json:"proxySSLSNIEnabled,omitempty" xml:"proxySSLSNIEnabled,omitempty"`
  // {"en":"After enabling SNI, you need to configure the specified SNI server to include the SNI information carried in the origin TLS handshake. If left empty or don't carry the parameter, the default value will follows the origin host header.","zh_CN":"启用SNI后需同时配置指定SNI服务器，即回源TLS握手将携带的SNI信息。支持自定义填写回源SNI服务器。参数值置空或者不传参数，表示回源SNI与回源HOST值一致。"}
  ProxySSLSNIServer *string `json:"proxySSLSNIServer,omitempty" xml:"proxySSLSNIServer,omitempty"`
  // {"en":"List of origin TLS protocol versions, supporting configuring multiple versions.- Options: TLSv1, TLSv1.1, TLSv1.2, TLSv1.3- Default: TLSv1, TLSv1.1, TLSv1.2, TLSv1.3","zh_CN":"回源TLS协议版本列表，支持配置多个版本。- 可选项：TLSv1、TLSv1.1、TLSv1.2、TLSv1.3- 默认值：TLSv1、TLSv1.1、TLSv1.2、TLSv1.3","exampleValue":"TLSv1,TLSv1.1,TLSv1.2,TLSv1.3"}
  ProxySSLVersion []*string `json:"proxySSLVersion,omitempty" xml:"proxySSLVersion,omitempty" type:"Repeated"`
}

func (s CreatePropertyForAkamaiMigrationRequestHostnamesDefaultOrigin) String() string {
  return tea.Prettify(s)
}

func (s CreatePropertyForAkamaiMigrationRequestHostnamesDefaultOrigin) GoString() string {
  return s.String()
}

func (s *CreatePropertyForAkamaiMigrationRequestHostnamesDefaultOrigin) SetServers(v []*string) *CreatePropertyForAkamaiMigrationRequestHostnamesDefaultOrigin {
  s.Servers = v
  return s
}

func (s *CreatePropertyForAkamaiMigrationRequestHostnamesDefaultOrigin) SetHost(v string) *CreatePropertyForAkamaiMigrationRequestHostnamesDefaultOrigin {
  s.Host = &v
  return s
}

func (s *CreatePropertyForAkamaiMigrationRequestHostnamesDefaultOrigin) SetHttpPort(v int) *CreatePropertyForAkamaiMigrationRequestHostnamesDefaultOrigin {
  s.HttpPort = &v
  return s
}

func (s *CreatePropertyForAkamaiMigrationRequestHostnamesDefaultOrigin) SetHttpsPort(v int) *CreatePropertyForAkamaiMigrationRequestHostnamesDefaultOrigin {
  s.HttpsPort = &v
  return s
}

func (s *CreatePropertyForAkamaiMigrationRequestHostnamesDefaultOrigin) SetIpVersion(v string) *CreatePropertyForAkamaiMigrationRequestHostnamesDefaultOrigin {
  s.IpVersion = &v
  return s
}

func (s *CreatePropertyForAkamaiMigrationRequestHostnamesDefaultOrigin) SetScheme(v string) *CreatePropertyForAkamaiMigrationRequestHostnamesDefaultOrigin {
  s.Scheme = &v
  return s
}

func (s *CreatePropertyForAkamaiMigrationRequestHostnamesDefaultOrigin) SetProxySSLSNIEnabled(v bool) *CreatePropertyForAkamaiMigrationRequestHostnamesDefaultOrigin {
  s.ProxySSLSNIEnabled = &v
  return s
}

func (s *CreatePropertyForAkamaiMigrationRequestHostnamesDefaultOrigin) SetProxySSLSNIServer(v string) *CreatePropertyForAkamaiMigrationRequestHostnamesDefaultOrigin {
  s.ProxySSLSNIServer = &v
  return s
}

func (s *CreatePropertyForAkamaiMigrationRequestHostnamesDefaultOrigin) SetProxySSLVersion(v []*string) *CreatePropertyForAkamaiMigrationRequestHostnamesDefaultOrigin {
  s.ProxySSLVersion = v
  return s
}

type CreatePropertyForAkamaiMigrationRequestHeader struct {
  // {"en":"Property's Contract ID in Akamai.","zh_CN":"Akamai项目所属合同编号(Contract ID)。"}
  XAkAkamaiContractId *string `json:"x-ak-akamai-contract-id,omitempty" xml:"x-ak-akamai-contract-id,omitempty"`
  // {"en":"Customer account in Akamai.","zh_CN":"Akamai账号(Customer account)。"}
  XAkCustomerMainAccount *string `json:"x-ak-customer-main-account,omitempty" xml:"x-ak-customer-main-account,omitempty"`
  // {"en":"Property's Parent Group ID in Akamai.","zh_CN":"Akamai项目所属用户组的父组编号(Parent Group ID)。"}
  XAkParentGroupId *string `json:"x-ak-parent-group-id,omitempty" xml:"x-ak-parent-group-id,omitempty"`
  // {"en":"Property's Parent Group name in Akamai.","zh_CN":"Akamai项目所属用户组的父组名称(Parent Group name)。"}
  XAkParentGroupName *string `json:"x-ak-parent-group-name,omitempty" xml:"x-ak-parent-group-name,omitempty"`
  // {"en":"Property's group ID of the group.","zh_CN":"Akamai项目所属用户组编号(Group ID)。"}
  XAkGroupId *string `json:"x-ak-group-id,omitempty" xml:"x-ak-group-id,omitempty"`
  // {"en":"Property's Group name in Akamai.","zh_CN":"Akamai项目所属用户组名称(Group name)。"}
  XAkGroupName *string `json:"x-ak-group-name,omitempty" xml:"x-ak-group-name,omitempty"`
}

func (s CreatePropertyForAkamaiMigrationRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s CreatePropertyForAkamaiMigrationRequestHeader) GoString() string {
  return s.String()
}

func (s *CreatePropertyForAkamaiMigrationRequestHeader) SetXAkAkamaiContractId(v string) *CreatePropertyForAkamaiMigrationRequestHeader {
  s.XAkAkamaiContractId = &v
  return s
}

func (s *CreatePropertyForAkamaiMigrationRequestHeader) SetXAkCustomerMainAccount(v string) *CreatePropertyForAkamaiMigrationRequestHeader {
  s.XAkCustomerMainAccount = &v
  return s
}

func (s *CreatePropertyForAkamaiMigrationRequestHeader) SetXAkParentGroupId(v string) *CreatePropertyForAkamaiMigrationRequestHeader {
  s.XAkParentGroupId = &v
  return s
}

func (s *CreatePropertyForAkamaiMigrationRequestHeader) SetXAkParentGroupName(v string) *CreatePropertyForAkamaiMigrationRequestHeader {
  s.XAkParentGroupName = &v
  return s
}

func (s *CreatePropertyForAkamaiMigrationRequestHeader) SetXAkGroupId(v string) *CreatePropertyForAkamaiMigrationRequestHeader {
  s.XAkGroupId = &v
  return s
}

func (s *CreatePropertyForAkamaiMigrationRequestHeader) SetXAkGroupName(v string) *CreatePropertyForAkamaiMigrationRequestHeader {
  s.XAkGroupName = &v
  return s
}

type CreatePropertyForAkamaiMigrationPaths struct {
}

func (s CreatePropertyForAkamaiMigrationPaths) String() string {
  return tea.Prettify(s)
}

func (s CreatePropertyForAkamaiMigrationPaths) GoString() string {
  return s.String()
}

type CreatePropertyForAkamaiMigrationParameters struct {
}

func (s CreatePropertyForAkamaiMigrationParameters) String() string {
  return tea.Prettify(s)
}

func (s CreatePropertyForAkamaiMigrationParameters) GoString() string {
  return s.String()
}

type CreatePropertyForAkamaiMigrationResponse struct {
  // {"en":"Response code, 0 means successful.","zh_CN":"接口响应code，0代表成功。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response data.","zh_CN":"接口响应数据"}
  Data *CreatePropertyForAkamaiMigrationResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
  // {"en":"Response error message if failed.","zh_CN":"接口响应信息，success代表成功，失败则提供失败信息。"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Metadate transformation details. Metadate include Akamai Behavior and Criteria directives.","zh_CN":"来源元数据转换详情。元数据包含了 Akamai 的 Behaviors 和 Criterias 指令。"}
  MetaDataMsg *CreatePropertyForAkamaiMigrationResponseMetaDataMsg `json:"metaDataMsg,omitempty" xml:"metaDataMsg,omitempty" require:"true" type:"Struct"`
}

func (s CreatePropertyForAkamaiMigrationResponse) String() string {
  return tea.Prettify(s)
}

func (s CreatePropertyForAkamaiMigrationResponse) GoString() string {
  return s.String()
}

func (s *CreatePropertyForAkamaiMigrationResponse) SetCode(v string) *CreatePropertyForAkamaiMigrationResponse {
  s.Code = &v
  return s
}

func (s *CreatePropertyForAkamaiMigrationResponse) SetData(v *CreatePropertyForAkamaiMigrationResponseData) *CreatePropertyForAkamaiMigrationResponse {
  s.Data = v
  return s
}

func (s *CreatePropertyForAkamaiMigrationResponse) SetMessage(v string) *CreatePropertyForAkamaiMigrationResponse {
  s.Message = &v
  return s
}

func (s *CreatePropertyForAkamaiMigrationResponse) SetMetaDataMsg(v *CreatePropertyForAkamaiMigrationResponseMetaDataMsg) *CreatePropertyForAkamaiMigrationResponse {
  s.MetaDataMsg = v
  return s
}

type CreatePropertyForAkamaiMigrationResponseData struct {
  // {"en":"Property version.","zh_CN":"项目的版本"}
  PropertyVersion *int `json:"propertyVersion,omitempty" xml:"propertyVersion,omitempty" require:"true"`
  // {"en":"Property name.","zh_CN":"项目的名称"}
  PropertyName *string `json:"propertyName,omitempty" xml:"propertyName,omitempty" require:"true"`
  // {"en":"Property ID","zh_CN":"项目标识"}
  PropertyId *int64 `json:"propertyId,omitempty" xml:"propertyId,omitempty" require:"true"`
}

func (s CreatePropertyForAkamaiMigrationResponseData) String() string {
  return tea.Prettify(s)
}

func (s CreatePropertyForAkamaiMigrationResponseData) GoString() string {
  return s.String()
}

func (s *CreatePropertyForAkamaiMigrationResponseData) SetPropertyVersion(v int) *CreatePropertyForAkamaiMigrationResponseData {
  s.PropertyVersion = &v
  return s
}

func (s *CreatePropertyForAkamaiMigrationResponseData) SetPropertyName(v string) *CreatePropertyForAkamaiMigrationResponseData {
  s.PropertyName = &v
  return s
}

func (s *CreatePropertyForAkamaiMigrationResponseData) SetPropertyId(v int64) *CreatePropertyForAkamaiMigrationResponseData {
  s.PropertyId = &v
  return s
}

type CreatePropertyForAkamaiMigrationResponseMetaDataMsg struct {
  // {"en":"Metadate transformation details. For Behavior/Criteria metadata in Akamai JSON, classifies transformation log level into 4 message types (INFO, NOTICE, WARN, ERR ) and 9 error code (100001, 100002, 200001 etc.)","zh_CN":"元数据转换结果详情。对Akamai JSON中的Behavior/Criteria元数据，打印了配置转换反馈日志，划分了4类消息类型（INFO, NOTICE, WARN, ERR）和9类错误码（100001，100002，200001等）。"}
  MsgData *string `json:"msgData,omitempty" xml:"msgData,omitempty" require:"true"`
  // {"en":"Statistical information of metadate transformation.  The Transformation success rate can be calculated based on statistical information. Calculation formula:success rate = (1- ERR/COUNT)100%.","zh_CN":"元数据转换结果统计信息。可根据统计信息计算转换成功率。计算公式：转换成功率 = （1- ERR/COUNT）*100%。"}
  MsgStat *int64 `json:"msgStat,omitempty" xml:"msgStat,omitempty" require:"true"`
}

func (s CreatePropertyForAkamaiMigrationResponseMetaDataMsg) String() string {
  return tea.Prettify(s)
}

func (s CreatePropertyForAkamaiMigrationResponseMetaDataMsg) GoString() string {
  return s.String()
}

func (s *CreatePropertyForAkamaiMigrationResponseMetaDataMsg) SetMsgData(v string) *CreatePropertyForAkamaiMigrationResponseMetaDataMsg {
  s.MsgData = &v
  return s
}

func (s *CreatePropertyForAkamaiMigrationResponseMetaDataMsg) SetMsgStat(v int64) *CreatePropertyForAkamaiMigrationResponseMetaDataMsg {
  s.MsgStat = &v
  return s
}

type CreatePropertyForAkamaiMigrationResponseHeader struct {
}

func (s CreatePropertyForAkamaiMigrationResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s CreatePropertyForAkamaiMigrationResponseHeader) GoString() string {
  return s.String()
}




type CreatePropertyRequest struct {
  // {"en":"Product Service Type related to your contract.","zh_CN":"产品服务类型。请根据您的合同产品服务类型填写。","exampleValue":"wsa,wsa-https"}
  ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty"`
  // {"en":"Property description.The length must not exceed 256 characters.","zh_CN":"项目的描述。长度不超过256个字符。"}
  PropertyComment *string `json:"propertyComment,omitempty" xml:"propertyComment,omitempty"`
  // {"en":"The Variables feature allows you to define variables, assign values to them, and reuse them in functional test cases. A variable consists of a name and a value.","zh_CN":"变量配置，详情查看schema接口。"}
  Variables []*CreatePropertyRequestVariables `json:"variables,omitempty" xml:"variables,omitempty" type:"Repeated"`
  // {"en":"Clone param","zh_CN":"克隆版本参数"}
  CloneParam *CreatePropertyRequestCloneParam `json:"cloneParam,omitempty" xml:"cloneParam,omitempty" type:"Struct"`
  // {"en":"The name of the property. The length must not exceed 128 characters.","zh_CN":"项目的名称。长度不超过128个字符。"}
  PropertyName *string `json:"propertyName,omitempty" xml:"propertyName,omitempty"`
  // {"en":"Property version description.The length must not exceed 256 characters.","zh_CN":"项目版本的描述。长度不超过256个字符。"}
  VersionComment *string `json:"versionComment,omitempty" xml:"versionComment,omitempty"`
  // {"en":"Specify the hostname and settings used to contact the origin once service begins.","zh_CN":"回源配置，详情查看schema接口。"}
  Origins []*CreatePropertyRequestOrigins `json:"origins,omitempty" xml:"origins,omitempty" type:"Repeated"`
  // {"en":"hostnames","zh_CN":"域名列表"}
  Hostnames []*CreatePropertyRequestHostnames `json:"hostnames,omitempty" xml:"hostnames,omitempty" require:"true" type:"Repeated"`
  // {"en":"Rules","zh_CN":"规则"}
  Rules *CreatePropertyRequestRules `json:"rules,omitempty" xml:"rules,omitempty" type:"Struct"`
  // {"en":"The id of contract, such as 40015677","zh_CN":"合同号，如40015677"}
  ContractId *string `json:"contractId,omitempty" xml:"contractId,omitempty"`
  // {"en":"The id of product, such as 10","zh_CN":"产品号，如10"}
  ItemId *string `json:"itemId,omitempty" xml:"itemId,omitempty"`
}

func (s CreatePropertyRequest) String() string {
  return tea.Prettify(s)
}

func (s CreatePropertyRequest) GoString() string {
  return s.String()
}

func (s *CreatePropertyRequest) SetServiceType(v string) *CreatePropertyRequest {
  s.ServiceType = &v
  return s
}

func (s *CreatePropertyRequest) SetPropertyComment(v string) *CreatePropertyRequest {
  s.PropertyComment = &v
  return s
}

func (s *CreatePropertyRequest) SetVariables(v []*CreatePropertyRequestVariables) *CreatePropertyRequest {
  s.Variables = v
  return s
}

func (s *CreatePropertyRequest) SetCloneParam(v *CreatePropertyRequestCloneParam) *CreatePropertyRequest {
  s.CloneParam = v
  return s
}

func (s *CreatePropertyRequest) SetPropertyName(v string) *CreatePropertyRequest {
  s.PropertyName = &v
  return s
}

func (s *CreatePropertyRequest) SetVersionComment(v string) *CreatePropertyRequest {
  s.VersionComment = &v
  return s
}

func (s *CreatePropertyRequest) SetOrigins(v []*CreatePropertyRequestOrigins) *CreatePropertyRequest {
  s.Origins = v
  return s
}

func (s *CreatePropertyRequest) SetHostnames(v []*CreatePropertyRequestHostnames) *CreatePropertyRequest {
  s.Hostnames = v
  return s
}

func (s *CreatePropertyRequest) SetRules(v *CreatePropertyRequestRules) *CreatePropertyRequest {
  s.Rules = v
  return s
}

func (s *CreatePropertyRequest) SetContractId(v string) *CreatePropertyRequest {
  s.ContractId = &v
  return s
}

func (s *CreatePropertyRequest) SetItemId(v string) *CreatePropertyRequest {
  s.ItemId = &v
  return s
}

type CreatePropertyRequestVariables struct     {
  // {"en":"condition expression","zh_CN":"条件表达式"}
  Condition *string `json:"condition,omitempty" xml:"condition,omitempty"`
  // {"en":"action","zh_CN":"动作"}
  Action *CreatePropertyRequestVariablesAction `json:"action,omitempty" xml:"action,omitempty" type:"Struct"`
  // {"en":"The name of variable","zh_CN":"变量名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
}

func (s CreatePropertyRequestVariables) String() string {
  return tea.Prettify(s)
}

func (s CreatePropertyRequestVariables) GoString() string {
  return s.String()
}

func (s *CreatePropertyRequestVariables) SetCondition(v string) *CreatePropertyRequestVariables {
  s.Condition = &v
  return s
}

func (s *CreatePropertyRequestVariables) SetAction(v *CreatePropertyRequestVariablesAction) *CreatePropertyRequestVariables {
  s.Action = v
  return s
}

func (s *CreatePropertyRequestVariables) SetName(v string) *CreatePropertyRequestVariables {
  s.Name = &v
  return s
}

type CreatePropertyRequestVariablesAction struct {
  // {"en":"name","zh_CN":"名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // {"en":"parameter","zh_CN":"参数"}
  Options *CreatePropertyRequestVariablesActionOptions `json:"options,omitempty" xml:"options,omitempty" type:"Struct"`
}

func (s CreatePropertyRequestVariablesAction) String() string {
  return tea.Prettify(s)
}

func (s CreatePropertyRequestVariablesAction) GoString() string {
  return s.String()
}

func (s *CreatePropertyRequestVariablesAction) SetName(v string) *CreatePropertyRequestVariablesAction {
  s.Name = &v
  return s
}

func (s *CreatePropertyRequestVariablesAction) SetOptions(v *CreatePropertyRequestVariablesActionOptions) *CreatePropertyRequestVariablesAction {
  s.Options = v
  return s
}

type CreatePropertyRequestVariablesActionOptions struct {
  // {"en":"variate name","zh_CN":"变量名称"}
  VarName *string `json:"varName,omitempty" xml:"varName,omitempty"`
  // {"en":"variate value","zh_CN":"变量值"}
  Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s CreatePropertyRequestVariablesActionOptions) String() string {
  return tea.Prettify(s)
}

func (s CreatePropertyRequestVariablesActionOptions) GoString() string {
  return s.String()
}

func (s *CreatePropertyRequestVariablesActionOptions) SetVarName(v string) *CreatePropertyRequestVariablesActionOptions {
  s.VarName = &v
  return s
}

func (s *CreatePropertyRequestVariablesActionOptions) SetValue(v string) *CreatePropertyRequestVariablesActionOptions {
  s.Value = &v
  return s
}

type CreatePropertyRequestCloneParam struct {
  // {"en":"The property version you need to clone","zh_CN":"克隆项目版本"}
  PropertyVersion *int `json:"propertyVersion,omitempty" xml:"propertyVersion,omitempty" require:"true"`
  // {"en":"The property Id you need to clone","zh_CN":"克隆项目id"}
  PropertyId *int `json:"propertyId,omitempty" xml:"propertyId,omitempty" require:"true"`
}

func (s CreatePropertyRequestCloneParam) String() string {
  return tea.Prettify(s)
}

func (s CreatePropertyRequestCloneParam) GoString() string {
  return s.String()
}

func (s *CreatePropertyRequestCloneParam) SetPropertyVersion(v int) *CreatePropertyRequestCloneParam {
  s.PropertyVersion = &v
  return s
}

func (s *CreatePropertyRequestCloneParam) SetPropertyId(v int) *CreatePropertyRequestCloneParam {
  s.PropertyId = &v
  return s
}

type CreatePropertyRequestOrigins struct     {
  // {"en":"Origin servers list","zh_CN":"源站服务器列表"}
  Servers []*CreatePropertyRequestOriginsServers `json:"servers,omitempty" xml:"servers,omitempty" type:"Repeated"`
  // {"en":"Http port","zh_CN":"http端口"}
  HttpPort *int `json:"httpPort,omitempty" xml:"httpPort,omitempty"`
  // {"en":"Origin name. Supports upper and lower case of English characters(a-z, A-Z), and dots(.) Underline(_), horizontal bar(-).","zh_CN":"源站名称。支持英文字符大小写（a-z，A-Z），点号(.)、下划线（_）、横杠（-）。"}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // {"en":"Https port","zh_CN":"https端口"}
  HttpsPort *int `json:"httpsPort,omitempty" xml:"httpsPort,omitempty"`
  // {"en":"Enable SNI. Whether to enable SNI. If enabled, the origin TLS handshake will carry SNI (Server Name Indication).","zh_CN":"是否启用SNI。开启后，回源TLS握手将携带SNI（Server Name Indication）。","exampleValue":"true,false"}
  ProxySSLSNIEnabled *bool `json:"proxySSLSNIEnabled,omitempty" xml:"proxySSLSNIEnabled,omitempty"`
  // {"en":"SNI Server Name. After enabling SNI, you need to configure the specified SNI server to include the SNI information carried in the origin TLS handshake. Default: ${http.request.host}, which follows the client request host. If left empty, the default value will be used.","zh_CN":"SNI服务器。启用SNI后需同时配置指定SNI服务器，即回源TLS握手将携带的SNI信息。默认值：${http.request.host}，即跟随客户端请求host。空值：置空则与默认值保持一致。"}
  ProxySSLSNIServer *string `json:"proxySSLSNIServer,omitempty" xml:"proxySSLSNIServer,omitempty"`
  // {"defaultValue":"TLSv1,TLSv1.1,TLSv1.2,TLSv1.3","en":"TLS Versions. List of origin TLS protocol versions, supporting configuring multiple versions.","zh_CN":"TLS版本。回源TLS协议版本列表，支持配置多个版本。","exampleValue":"SSLv3,TLSv1,TLSv1.1,TLSv1.2,TLSv1.3"}
  ProxySSLVersion []*string `json:"proxySSLVersion,omitempty" xml:"proxySSLVersion,omitempty" type:"Repeated"`
  // {"en":"Verify Origin Certificate. Whether to enable origin certificate verification. If enabled, the back to origin node will verify the certificate offer by origin.","zh_CN":"是否开启节点验证源站证书。开启后，节点访问源站将会验证源站发过来的证书。","exampleValue":"true,false"}
  ProxySSLVerifyEnabled *bool `json:"proxySSLVerifyEnabled,omitempty" xml:"proxySSLVerifyEnabled,omitempty"`
  // {"en":"Trust CA Certificate ID. Trusted CA Certificate for origin certificate verification.","zh_CN":"信任CA证书ID。配置节点验证源站时使用的信任证书。"}
  ProxySSLTrustedCertificate *string `json:"proxySSLTrustedCertificate,omitempty" xml:"proxySSLTrustedCertificate,omitempty"`
}

func (s CreatePropertyRequestOrigins) String() string {
  return tea.Prettify(s)
}

func (s CreatePropertyRequestOrigins) GoString() string {
  return s.String()
}

func (s *CreatePropertyRequestOrigins) SetServers(v []*CreatePropertyRequestOriginsServers) *CreatePropertyRequestOrigins {
  s.Servers = v
  return s
}

func (s *CreatePropertyRequestOrigins) SetHttpPort(v int) *CreatePropertyRequestOrigins {
  s.HttpPort = &v
  return s
}

func (s *CreatePropertyRequestOrigins) SetName(v string) *CreatePropertyRequestOrigins {
  s.Name = &v
  return s
}

func (s *CreatePropertyRequestOrigins) SetHttpsPort(v int) *CreatePropertyRequestOrigins {
  s.HttpsPort = &v
  return s
}

func (s *CreatePropertyRequestOrigins) SetProxySSLSNIEnabled(v bool) *CreatePropertyRequestOrigins {
  s.ProxySSLSNIEnabled = &v
  return s
}

func (s *CreatePropertyRequestOrigins) SetProxySSLSNIServer(v string) *CreatePropertyRequestOrigins {
  s.ProxySSLSNIServer = &v
  return s
}

func (s *CreatePropertyRequestOrigins) SetProxySSLVersion(v []*string) *CreatePropertyRequestOrigins {
  s.ProxySSLVersion = v
  return s
}

func (s *CreatePropertyRequestOrigins) SetProxySSLVerifyEnabled(v bool) *CreatePropertyRequestOrigins {
  s.ProxySSLVerifyEnabled = &v
  return s
}

func (s *CreatePropertyRequestOrigins) SetProxySSLTrustedCertificate(v string) *CreatePropertyRequestOrigins {
  s.ProxySSLTrustedCertificate = &v
  return s
}

type CreatePropertyRequestOriginsServers struct     {
  // {"en":"Origin server","zh_CN":"源站服务器"}
  Server *string `json:"server,omitempty" xml:"server,omitempty" require:"true"`
  // {"en":"weight","zh_CN":"权重"}
  Weight *int `json:"weight,omitempty" xml:"weight,omitempty"`
  // {"en":"priority","zh_CN":"优先级"}
  Priority *int `json:"priority,omitempty" xml:"priority,omitempty"`
}

func (s CreatePropertyRequestOriginsServers) String() string {
  return tea.Prettify(s)
}

func (s CreatePropertyRequestOriginsServers) GoString() string {
  return s.String()
}

func (s *CreatePropertyRequestOriginsServers) SetServer(v string) *CreatePropertyRequestOriginsServers {
  s.Server = &v
  return s
}

func (s *CreatePropertyRequestOriginsServers) SetWeight(v int) *CreatePropertyRequestOriginsServers {
  s.Weight = &v
  return s
}

func (s *CreatePropertyRequestOriginsServers) SetPriority(v int) *CreatePropertyRequestOriginsServers {
  s.Priority = &v
  return s
}

type CreatePropertyRequestHostnames struct     {
  // {"en":"default origin","zh_CN":"默认源站"}
  DefaultOrigin *CreatePropertyRequestHostnamesDefaultOrigin `json:"defaultOrigin,omitempty" xml:"defaultOrigin,omitempty" type:"Struct"`
  // {"en":"hostname, the length must not exceed 128 characters. A wildcard hostname must start with an asterisk (*).","zh_CN":"域名，长度不超过128个字符。泛域名需要以“*”开头。"}
  Hostname *string `json:"hostname,omitempty" xml:"hostname,omitempty"`
  // {"en":"hostname association ssl configuration","zh_CN":"关联证书配置"}
  Certificates []*CreatePropertyRequestHostnamesCertificates `json:"certificates,omitempty" xml:"certificates,omitempty" type:"Repeated"`
  // {"en":"hostname edge-hostname config","zh_CN":"调度域名配置"}
  EdgeHostname *CreatePropertyRequestHostnamesEdgeHostname `json:"edgeHostname,omitempty" xml:"edgeHostname,omitempty" type:"Struct"`
}

func (s CreatePropertyRequestHostnames) String() string {
  return tea.Prettify(s)
}

func (s CreatePropertyRequestHostnames) GoString() string {
  return s.String()
}

func (s *CreatePropertyRequestHostnames) SetDefaultOrigin(v *CreatePropertyRequestHostnamesDefaultOrigin) *CreatePropertyRequestHostnames {
  s.DefaultOrigin = v
  return s
}

func (s *CreatePropertyRequestHostnames) SetHostname(v string) *CreatePropertyRequestHostnames {
  s.Hostname = &v
  return s
}

func (s *CreatePropertyRequestHostnames) SetCertificates(v []*CreatePropertyRequestHostnamesCertificates) *CreatePropertyRequestHostnames {
  s.Certificates = v
  return s
}

func (s *CreatePropertyRequestHostnames) SetEdgeHostname(v *CreatePropertyRequestHostnamesEdgeHostname) *CreatePropertyRequestHostnames {
  s.EdgeHostname = v
  return s
}

type CreatePropertyRequestHostnamesDefaultOrigin struct {
  // {"en":"origin servers","zh_CN":"源站服务器"}
  Servers []*string `json:"servers,omitempty" xml:"servers,omitempty" type:"Repeated"`
  // {"en":"IP version","zh_CN":"IP版本","exampleValue":"dual,ipv4,ipv6"}
  IpVersion *string `json:"ipVersion,omitempty" xml:"ipVersion,omitempty"`
  // {"en":"http port","zh_CN":"http端口"}
  HttpPort *int `json:"httpPort,omitempty" xml:"httpPort,omitempty"`
  // {"en":"origin host","zh_CN":"回源HOST"}
  Host *string `json:"host,omitempty" xml:"host,omitempty"`
  // {"en":"https port","zh_CN":"https端口"}
  HttpsPort *int `json:"httpsPort,omitempty" xml:"httpsPort,omitempty"`
  // {"en":"Scheme used for origin requests. Options include:\n- http: Fixedly use HTTP.\n- https: Fixedly use HTTPS.\n- honor: default option.Follow the scheme of the client request (default).","zh_CN":"回源请求HTTP/HTTPS协议。可选项如下： \n- http：固定使用HTTP协议回源。\n- https：固定使用HTTPS协议回源。\n- honor：默认选项。根据客户端请求的协议类型回源。即客户端请求使用HTTP则用HTTP回源；客户端请求使用HTTPS则用HTTPS回源。","exampleValue":"http,https,honor"}
  Scheme *string `json:"scheme,omitempty" xml:"scheme,omitempty"`
  // {"en":"Whether to enable SNI. Options include:\n- true: The origin TLS handshake will carry SNI (Server Name Indication).Note: Before enabling SNI, please upload the certificate to the origin and set the back-to-origin scheme to https.\n- false: The origin TLS handshake will not carry SNI (Server Name Indication)","zh_CN":"是否启用SNI。可选项如下：\n- true：回源TLS握手将携带SNI（Server Name Indication）。注意：启用SNI前，请为源站配置证书并设置回源scheme为https。\n- false：回源TLS握手不携带SNI（Server Name Indication）。","exampleValue":"true,false"}
  ProxySSLSNIEnabled *bool `json:"proxySSLSNIEnabled,omitempty" xml:"proxySSLSNIEnabled,omitempty"`
  // {"en":"After enabling SNI, you need to configure the specified SNI server to include the SNI information carried in the origin TLS handshake. If left empty or don't carry the parameter, the default value will follows the origin host header.","zh_CN":"启用SNI后需同时配置指定SNI服务器，即回源TLS握手将携带的SNI信息。支持自定义填写回源SNI服务器。参数值置空或者不传参数，表示回源SNI与回源HOST值一致。"}
  ProxySSLSNIServer *string `json:"proxySSLSNIServer,omitempty" xml:"proxySSLSNIServer,omitempty"`
  // {"en":"List of origin TLS protocol versions, supporting configuring multiple versions.\n- Options: TLSv1, TLSv1.1, TLSv1.2, TLSv1.3\n- Default: TLSv1, TLSv1.1, TLSv1.2, TLSv1.3","zh_CN":"回源TLS协议版本列表，支持配置多个版本。\n- 可选项：TLSv1、TLSv1.1、TLSv1.2、TLSv1.3\n- 默认值：TLSv1、TLSv1.1、TLSv1.2、TLSv1.3","exampleValue":"TLSv1,TLSv1.1,TLSv1.2,TLSv1.3"}
  ProxySSLVersion []*string `json:"proxySSLVersion,omitempty" xml:"proxySSLVersion,omitempty" type:"Repeated"`
}

func (s CreatePropertyRequestHostnamesDefaultOrigin) String() string {
  return tea.Prettify(s)
}

func (s CreatePropertyRequestHostnamesDefaultOrigin) GoString() string {
  return s.String()
}

func (s *CreatePropertyRequestHostnamesDefaultOrigin) SetServers(v []*string) *CreatePropertyRequestHostnamesDefaultOrigin {
  s.Servers = v
  return s
}

func (s *CreatePropertyRequestHostnamesDefaultOrigin) SetIpVersion(v string) *CreatePropertyRequestHostnamesDefaultOrigin {
  s.IpVersion = &v
  return s
}

func (s *CreatePropertyRequestHostnamesDefaultOrigin) SetHttpPort(v int) *CreatePropertyRequestHostnamesDefaultOrigin {
  s.HttpPort = &v
  return s
}

func (s *CreatePropertyRequestHostnamesDefaultOrigin) SetHost(v string) *CreatePropertyRequestHostnamesDefaultOrigin {
  s.Host = &v
  return s
}

func (s *CreatePropertyRequestHostnamesDefaultOrigin) SetHttpsPort(v int) *CreatePropertyRequestHostnamesDefaultOrigin {
  s.HttpsPort = &v
  return s
}

func (s *CreatePropertyRequestHostnamesDefaultOrigin) SetScheme(v string) *CreatePropertyRequestHostnamesDefaultOrigin {
  s.Scheme = &v
  return s
}

func (s *CreatePropertyRequestHostnamesDefaultOrigin) SetProxySSLSNIEnabled(v bool) *CreatePropertyRequestHostnamesDefaultOrigin {
  s.ProxySSLSNIEnabled = &v
  return s
}

func (s *CreatePropertyRequestHostnamesDefaultOrigin) SetProxySSLSNIServer(v string) *CreatePropertyRequestHostnamesDefaultOrigin {
  s.ProxySSLSNIServer = &v
  return s
}

func (s *CreatePropertyRequestHostnamesDefaultOrigin) SetProxySSLVersion(v []*string) *CreatePropertyRequestHostnamesDefaultOrigin {
  s.ProxySSLVersion = v
  return s
}

type CreatePropertyRequestHostnamesCertificates struct     {
  // {"en":"Certificate ID","zh_CN":"证书ID"}
  CertificateId *int `json:"certificateId,omitempty" xml:"certificateId,omitempty"`
  // {"en":"certificate usage. \n- default_sni: SNI Edge Certificate;\n- dual_sni: Additional SNI Edge Certificate;\n- gm_sm2_enc: Additional SM2 Encryption Certificate;\n- gm_sm2_sign: Additional SM2 Signature Certificate;\n- client_mtls: mTLS Client CA Certificate;\n- origin_mtls: Origin mTLS Certificate;","zh_CN":"证书用途。\n- default_sni：表示SNI边缘证书；\n- dual_sni：表示多栈SNI边缘证书；\n- gm_sm2_enc：表示国密加密证书；\n- gm_sm2_sign：表示国密签名证书；\n- client_mtls：表示mTLS客户端CA证书；\n- origin_mtls：表示回源mTLS证书。","exampleValue":"default_sni,dual_sni,gm_sm2_enc,gm_sm2_sign,client_mtls"}
  CertificateUsage *string `json:"certificateUsage,omitempty" xml:"certificateUsage,omitempty"`
}

func (s CreatePropertyRequestHostnamesCertificates) String() string {
  return tea.Prettify(s)
}

func (s CreatePropertyRequestHostnamesCertificates) GoString() string {
  return s.String()
}

func (s *CreatePropertyRequestHostnamesCertificates) SetCertificateId(v int) *CreatePropertyRequestHostnamesCertificates {
  s.CertificateId = &v
  return s
}

func (s *CreatePropertyRequestHostnamesCertificates) SetCertificateUsage(v string) *CreatePropertyRequestHostnamesCertificates {
  s.CertificateUsage = &v
  return s
}

type CreatePropertyRequestHostnamesEdgeHostname struct {
  // {"en":"edge-hostname prefix","zh_CN":"调度域名前缀"}
  EdgeHostnamePrefix *string `json:"edgeHostnamePrefix,omitempty" xml:"edgeHostnamePrefix,omitempty" require:"true"`
  // {"en":"edge-hostname comment","zh_CN":"调度域名描述"}
  Comment *string `json:"comment,omitempty" xml:"comment,omitempty"`
  // {"en":"edge-hostname suffix","zh_CN":"调度域名后缀"}
  EdgeHostnameSuffix *string `json:"edgeHostnameSuffix,omitempty" xml:"edgeHostnameSuffix,omitempty" require:"true"`
}

func (s CreatePropertyRequestHostnamesEdgeHostname) String() string {
  return tea.Prettify(s)
}

func (s CreatePropertyRequestHostnamesEdgeHostname) GoString() string {
  return s.String()
}

func (s *CreatePropertyRequestHostnamesEdgeHostname) SetEdgeHostnamePrefix(v string) *CreatePropertyRequestHostnamesEdgeHostname {
  s.EdgeHostnamePrefix = &v
  return s
}

func (s *CreatePropertyRequestHostnamesEdgeHostname) SetComment(v string) *CreatePropertyRequestHostnamesEdgeHostname {
  s.Comment = &v
  return s
}

func (s *CreatePropertyRequestHostnamesEdgeHostname) SetEdgeHostnameSuffix(v string) *CreatePropertyRequestHostnamesEdgeHostname {
  s.EdgeHostnameSuffix = &v
  return s
}

type CreatePropertyRequestRules struct {
  // {"en":"response phase","zh_CN":"响应阶段"}
  ResponsePhase []*CreatePropertyRequestRulesResponsePhase `json:"responsePhase,omitempty" xml:"responsePhase,omitempty" type:"Repeated"`
  // {"en":"origin phase","zh_CN":"回源阶段"}
  OriginPhase []*CreatePropertyRequestRulesOriginPhase `json:"originPhase,omitempty" xml:"originPhase,omitempty" type:"Repeated"`
  // {"en":"cache phase","zh_CN":"缓存阶段"}
  CachePhase []*CreatePropertyRequestRulesCachePhase `json:"cachePhase,omitempty" xml:"cachePhase,omitempty" type:"Repeated"`
  // {"en":"connect phase","zh_CN":"连接阶段"}
  ConnectPhase []*CreatePropertyRequestRulesConnectPhase `json:"connectPhase,omitempty" xml:"connectPhase,omitempty" type:"Repeated"`
  // {"en":"request phase","zh_CN":"请求阶段"}
  ReqeustPhase []*CreatePropertyRequestRulesReqeustPhase `json:"reqeustPhase,omitempty" xml:"reqeustPhase,omitempty" type:"Repeated"`
}

func (s CreatePropertyRequestRules) String() string {
  return tea.Prettify(s)
}

func (s CreatePropertyRequestRules) GoString() string {
  return s.String()
}

func (s *CreatePropertyRequestRules) SetResponsePhase(v []*CreatePropertyRequestRulesResponsePhase) *CreatePropertyRequestRules {
  s.ResponsePhase = v
  return s
}

func (s *CreatePropertyRequestRules) SetOriginPhase(v []*CreatePropertyRequestRulesOriginPhase) *CreatePropertyRequestRules {
  s.OriginPhase = v
  return s
}

func (s *CreatePropertyRequestRules) SetCachePhase(v []*CreatePropertyRequestRulesCachePhase) *CreatePropertyRequestRules {
  s.CachePhase = v
  return s
}

func (s *CreatePropertyRequestRules) SetConnectPhase(v []*CreatePropertyRequestRulesConnectPhase) *CreatePropertyRequestRules {
  s.ConnectPhase = v
  return s
}

func (s *CreatePropertyRequestRules) SetReqeustPhase(v []*CreatePropertyRequestRulesReqeustPhase) *CreatePropertyRequestRules {
  s.ReqeustPhase = v
  return s
}

type CreatePropertyRequestRulesResponsePhase struct     {
  // {"en":"The name of rule.","zh_CN":"规则名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // {"en":"condition expression","zh_CN":"条件表达式"}
  Condition *string `json:"condition,omitempty" xml:"condition,omitempty"`
  // {"en":"action","zh_CN":"动作"}
  Action *CreatePropertyRequestRulesResponsePhaseAction `json:"action,omitempty" xml:"action,omitempty" require:"true" type:"Struct"`
  // {"en":"comment","zh_CN":"备注"}
  Description *string `json:"description,omitempty" xml:"description,omitempty"`
  // {"en":"priority","zh_CN":"排序"}
  Priority *int `json:"priority,omitempty" xml:"priority,omitempty"`
  // {"defaultValue":"true","en":"state","zh_CN":"状态","exampleValue":"true,false"}
  Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
}

func (s CreatePropertyRequestRulesResponsePhase) String() string {
  return tea.Prettify(s)
}

func (s CreatePropertyRequestRulesResponsePhase) GoString() string {
  return s.String()
}

func (s *CreatePropertyRequestRulesResponsePhase) SetName(v string) *CreatePropertyRequestRulesResponsePhase {
  s.Name = &v
  return s
}

func (s *CreatePropertyRequestRulesResponsePhase) SetCondition(v string) *CreatePropertyRequestRulesResponsePhase {
  s.Condition = &v
  return s
}

func (s *CreatePropertyRequestRulesResponsePhase) SetAction(v *CreatePropertyRequestRulesResponsePhaseAction) *CreatePropertyRequestRulesResponsePhase {
  s.Action = v
  return s
}

func (s *CreatePropertyRequestRulesResponsePhase) SetDescription(v string) *CreatePropertyRequestRulesResponsePhase {
  s.Description = &v
  return s
}

func (s *CreatePropertyRequestRulesResponsePhase) SetPriority(v int) *CreatePropertyRequestRulesResponsePhase {
  s.Priority = &v
  return s
}

func (s *CreatePropertyRequestRulesResponsePhase) SetEnabled(v bool) *CreatePropertyRequestRulesResponsePhase {
  s.Enabled = &v
  return s
}

type CreatePropertyRequestRulesResponsePhaseAction struct {
  // {"en":"name","zh_CN":"名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // {"en":"parameter","zh_CN":"参数"}
  Options *CreatePropertyRequestRulesResponsePhaseActionOptions `json:"options,omitempty" xml:"options,omitempty" type:"Struct"`
}

func (s CreatePropertyRequestRulesResponsePhaseAction) String() string {
  return tea.Prettify(s)
}

func (s CreatePropertyRequestRulesResponsePhaseAction) GoString() string {
  return s.String()
}

func (s *CreatePropertyRequestRulesResponsePhaseAction) SetName(v string) *CreatePropertyRequestRulesResponsePhaseAction {
  s.Name = &v
  return s
}

func (s *CreatePropertyRequestRulesResponsePhaseAction) SetOptions(v *CreatePropertyRequestRulesResponsePhaseActionOptions) *CreatePropertyRequestRulesResponsePhaseAction {
  s.Options = v
  return s
}

type CreatePropertyRequestRulesResponsePhaseActionOptions struct {
}

func (s CreatePropertyRequestRulesResponsePhaseActionOptions) String() string {
  return tea.Prettify(s)
}

func (s CreatePropertyRequestRulesResponsePhaseActionOptions) GoString() string {
  return s.String()
}

type CreatePropertyRequestRulesOriginPhase struct     {
  // {"en":"The name of rule.","zh_CN":"规则名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // {"en":"condition expression","zh_CN":"条件表达式"}
  Condition *string `json:"condition,omitempty" xml:"condition,omitempty"`
  // {"en":"action","zh_CN":"动作"}
  Action *CreatePropertyRequestRulesOriginPhaseAction `json:"action,omitempty" xml:"action,omitempty" require:"true" type:"Struct"`
  // {"en":"comment","zh_CN":"备注"}
  Description *string `json:"description,omitempty" xml:"description,omitempty"`
  // {"en":"priority","zh_CN":"排序"}
  Priority *int `json:"priority,omitempty" xml:"priority,omitempty"`
  // {"defaultValue":"true","en":"state","zh_CN":"状态","exampleValue":"true,false"}
  Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
}

func (s CreatePropertyRequestRulesOriginPhase) String() string {
  return tea.Prettify(s)
}

func (s CreatePropertyRequestRulesOriginPhase) GoString() string {
  return s.String()
}

func (s *CreatePropertyRequestRulesOriginPhase) SetName(v string) *CreatePropertyRequestRulesOriginPhase {
  s.Name = &v
  return s
}

func (s *CreatePropertyRequestRulesOriginPhase) SetCondition(v string) *CreatePropertyRequestRulesOriginPhase {
  s.Condition = &v
  return s
}

func (s *CreatePropertyRequestRulesOriginPhase) SetAction(v *CreatePropertyRequestRulesOriginPhaseAction) *CreatePropertyRequestRulesOriginPhase {
  s.Action = v
  return s
}

func (s *CreatePropertyRequestRulesOriginPhase) SetDescription(v string) *CreatePropertyRequestRulesOriginPhase {
  s.Description = &v
  return s
}

func (s *CreatePropertyRequestRulesOriginPhase) SetPriority(v int) *CreatePropertyRequestRulesOriginPhase {
  s.Priority = &v
  return s
}

func (s *CreatePropertyRequestRulesOriginPhase) SetEnabled(v bool) *CreatePropertyRequestRulesOriginPhase {
  s.Enabled = &v
  return s
}

type CreatePropertyRequestRulesOriginPhaseAction struct {
  // {"en":"name","zh_CN":"名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // {"en":"parameter","zh_CN":"参数"}
  Options *CreatePropertyRequestRulesOriginPhaseActionOptions `json:"options,omitempty" xml:"options,omitempty" type:"Struct"`
}

func (s CreatePropertyRequestRulesOriginPhaseAction) String() string {
  return tea.Prettify(s)
}

func (s CreatePropertyRequestRulesOriginPhaseAction) GoString() string {
  return s.String()
}

func (s *CreatePropertyRequestRulesOriginPhaseAction) SetName(v string) *CreatePropertyRequestRulesOriginPhaseAction {
  s.Name = &v
  return s
}

func (s *CreatePropertyRequestRulesOriginPhaseAction) SetOptions(v *CreatePropertyRequestRulesOriginPhaseActionOptions) *CreatePropertyRequestRulesOriginPhaseAction {
  s.Options = v
  return s
}

type CreatePropertyRequestRulesOriginPhaseActionOptions struct {
}

func (s CreatePropertyRequestRulesOriginPhaseActionOptions) String() string {
  return tea.Prettify(s)
}

func (s CreatePropertyRequestRulesOriginPhaseActionOptions) GoString() string {
  return s.String()
}

type CreatePropertyRequestRulesCachePhase struct     {
  // {"en":"The name of rule.","zh_CN":"规则名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // {"en":"condition expression","zh_CN":"条件表达式"}
  Condition *string `json:"condition,omitempty" xml:"condition,omitempty"`
  // {"en":"action","zh_CN":"动作"}
  Action *CreatePropertyRequestRulesCachePhaseAction `json:"action,omitempty" xml:"action,omitempty" require:"true" type:"Struct"`
  // {"en":"comment","zh_CN":"备注"}
  Description *string `json:"description,omitempty" xml:"description,omitempty"`
  // {"en":"priority","zh_CN":"排序"}
  Priority *int `json:"priority,omitempty" xml:"priority,omitempty"`
  // {"defaultValue":"true","en":"state","zh_CN":"状态","exampleValue":"true,false"}
  Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
}

func (s CreatePropertyRequestRulesCachePhase) String() string {
  return tea.Prettify(s)
}

func (s CreatePropertyRequestRulesCachePhase) GoString() string {
  return s.String()
}

func (s *CreatePropertyRequestRulesCachePhase) SetName(v string) *CreatePropertyRequestRulesCachePhase {
  s.Name = &v
  return s
}

func (s *CreatePropertyRequestRulesCachePhase) SetCondition(v string) *CreatePropertyRequestRulesCachePhase {
  s.Condition = &v
  return s
}

func (s *CreatePropertyRequestRulesCachePhase) SetAction(v *CreatePropertyRequestRulesCachePhaseAction) *CreatePropertyRequestRulesCachePhase {
  s.Action = v
  return s
}

func (s *CreatePropertyRequestRulesCachePhase) SetDescription(v string) *CreatePropertyRequestRulesCachePhase {
  s.Description = &v
  return s
}

func (s *CreatePropertyRequestRulesCachePhase) SetPriority(v int) *CreatePropertyRequestRulesCachePhase {
  s.Priority = &v
  return s
}

func (s *CreatePropertyRequestRulesCachePhase) SetEnabled(v bool) *CreatePropertyRequestRulesCachePhase {
  s.Enabled = &v
  return s
}

type CreatePropertyRequestRulesCachePhaseAction struct {
  // {"en":"name","zh_CN":"名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // {"en":"parameter","zh_CN":"参数"}
  Options *CreatePropertyRequestRulesCachePhaseActionOptions `json:"options,omitempty" xml:"options,omitempty" type:"Struct"`
}

func (s CreatePropertyRequestRulesCachePhaseAction) String() string {
  return tea.Prettify(s)
}

func (s CreatePropertyRequestRulesCachePhaseAction) GoString() string {
  return s.String()
}

func (s *CreatePropertyRequestRulesCachePhaseAction) SetName(v string) *CreatePropertyRequestRulesCachePhaseAction {
  s.Name = &v
  return s
}

func (s *CreatePropertyRequestRulesCachePhaseAction) SetOptions(v *CreatePropertyRequestRulesCachePhaseActionOptions) *CreatePropertyRequestRulesCachePhaseAction {
  s.Options = v
  return s
}

type CreatePropertyRequestRulesCachePhaseActionOptions struct {
}

func (s CreatePropertyRequestRulesCachePhaseActionOptions) String() string {
  return tea.Prettify(s)
}

func (s CreatePropertyRequestRulesCachePhaseActionOptions) GoString() string {
  return s.String()
}

type CreatePropertyRequestRulesConnectPhase struct     {
  // {"en":"The name of rule.","zh_CN":"规则名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // {"en":"condition expression","zh_CN":"条件表达式"}
  Condition *string `json:"condition,omitempty" xml:"condition,omitempty"`
  // {"en":"action","zh_CN":"动作"}
  Action *CreatePropertyRequestRulesConnectPhaseAction `json:"action,omitempty" xml:"action,omitempty" require:"true" type:"Struct"`
  // {"en":"comment","zh_CN":"备注"}
  Description *string `json:"description,omitempty" xml:"description,omitempty"`
  // {"en":"priority","zh_CN":"排序"}
  Priority *int `json:"priority,omitempty" xml:"priority,omitempty"`
  // {"defaultValue":"true","en":"state","zh_CN":"状态","exampleValue":"true,false"}
  Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
}

func (s CreatePropertyRequestRulesConnectPhase) String() string {
  return tea.Prettify(s)
}

func (s CreatePropertyRequestRulesConnectPhase) GoString() string {
  return s.String()
}

func (s *CreatePropertyRequestRulesConnectPhase) SetName(v string) *CreatePropertyRequestRulesConnectPhase {
  s.Name = &v
  return s
}

func (s *CreatePropertyRequestRulesConnectPhase) SetCondition(v string) *CreatePropertyRequestRulesConnectPhase {
  s.Condition = &v
  return s
}

func (s *CreatePropertyRequestRulesConnectPhase) SetAction(v *CreatePropertyRequestRulesConnectPhaseAction) *CreatePropertyRequestRulesConnectPhase {
  s.Action = v
  return s
}

func (s *CreatePropertyRequestRulesConnectPhase) SetDescription(v string) *CreatePropertyRequestRulesConnectPhase {
  s.Description = &v
  return s
}

func (s *CreatePropertyRequestRulesConnectPhase) SetPriority(v int) *CreatePropertyRequestRulesConnectPhase {
  s.Priority = &v
  return s
}

func (s *CreatePropertyRequestRulesConnectPhase) SetEnabled(v bool) *CreatePropertyRequestRulesConnectPhase {
  s.Enabled = &v
  return s
}

type CreatePropertyRequestRulesConnectPhaseAction struct {
  // {"en":"name","zh_CN":"名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // {"en":"parameter","zh_CN":"参数"}
  Options *CreatePropertyRequestRulesConnectPhaseActionOptions `json:"options,omitempty" xml:"options,omitempty" type:"Struct"`
}

func (s CreatePropertyRequestRulesConnectPhaseAction) String() string {
  return tea.Prettify(s)
}

func (s CreatePropertyRequestRulesConnectPhaseAction) GoString() string {
  return s.String()
}

func (s *CreatePropertyRequestRulesConnectPhaseAction) SetName(v string) *CreatePropertyRequestRulesConnectPhaseAction {
  s.Name = &v
  return s
}

func (s *CreatePropertyRequestRulesConnectPhaseAction) SetOptions(v *CreatePropertyRequestRulesConnectPhaseActionOptions) *CreatePropertyRequestRulesConnectPhaseAction {
  s.Options = v
  return s
}

type CreatePropertyRequestRulesConnectPhaseActionOptions struct {
}

func (s CreatePropertyRequestRulesConnectPhaseActionOptions) String() string {
  return tea.Prettify(s)
}

func (s CreatePropertyRequestRulesConnectPhaseActionOptions) GoString() string {
  return s.String()
}

type CreatePropertyRequestRulesReqeustPhase struct     {
  // {"en":"The name of rule.","zh_CN":"规则名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // {"en":"condition expression","zh_CN":"条件表达式"}
  Condition *string `json:"condition,omitempty" xml:"condition,omitempty"`
  // {"en":"action","zh_CN":"动作"}
  Action *CreatePropertyRequestRulesReqeustPhaseAction `json:"action,omitempty" xml:"action,omitempty" require:"true" type:"Struct"`
  // {"en":"comment","zh_CN":"备注"}
  Description *string `json:"description,omitempty" xml:"description,omitempty"`
  // {"en":"priority","zh_CN":"排序"}
  Priority *int `json:"priority,omitempty" xml:"priority,omitempty"`
  // {"defaultValue":"true","en":"state","zh_CN":"状态","exampleValue":"true,false"}
  Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
}

func (s CreatePropertyRequestRulesReqeustPhase) String() string {
  return tea.Prettify(s)
}

func (s CreatePropertyRequestRulesReqeustPhase) GoString() string {
  return s.String()
}

func (s *CreatePropertyRequestRulesReqeustPhase) SetName(v string) *CreatePropertyRequestRulesReqeustPhase {
  s.Name = &v
  return s
}

func (s *CreatePropertyRequestRulesReqeustPhase) SetCondition(v string) *CreatePropertyRequestRulesReqeustPhase {
  s.Condition = &v
  return s
}

func (s *CreatePropertyRequestRulesReqeustPhase) SetAction(v *CreatePropertyRequestRulesReqeustPhaseAction) *CreatePropertyRequestRulesReqeustPhase {
  s.Action = v
  return s
}

func (s *CreatePropertyRequestRulesReqeustPhase) SetDescription(v string) *CreatePropertyRequestRulesReqeustPhase {
  s.Description = &v
  return s
}

func (s *CreatePropertyRequestRulesReqeustPhase) SetPriority(v int) *CreatePropertyRequestRulesReqeustPhase {
  s.Priority = &v
  return s
}

func (s *CreatePropertyRequestRulesReqeustPhase) SetEnabled(v bool) *CreatePropertyRequestRulesReqeustPhase {
  s.Enabled = &v
  return s
}

type CreatePropertyRequestRulesReqeustPhaseAction struct {
  // {"en":"name","zh_CN":"名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // {"en":"parameter","zh_CN":"参数"}
  Options *CreatePropertyRequestRulesReqeustPhaseActionOptions `json:"options,omitempty" xml:"options,omitempty" type:"Struct"`
}

func (s CreatePropertyRequestRulesReqeustPhaseAction) String() string {
  return tea.Prettify(s)
}

func (s CreatePropertyRequestRulesReqeustPhaseAction) GoString() string {
  return s.String()
}

func (s *CreatePropertyRequestRulesReqeustPhaseAction) SetName(v string) *CreatePropertyRequestRulesReqeustPhaseAction {
  s.Name = &v
  return s
}

func (s *CreatePropertyRequestRulesReqeustPhaseAction) SetOptions(v *CreatePropertyRequestRulesReqeustPhaseActionOptions) *CreatePropertyRequestRulesReqeustPhaseAction {
  s.Options = v
  return s
}

type CreatePropertyRequestRulesReqeustPhaseActionOptions struct {
}

func (s CreatePropertyRequestRulesReqeustPhaseActionOptions) String() string {
  return tea.Prettify(s)
}

func (s CreatePropertyRequestRulesReqeustPhaseActionOptions) GoString() string {
  return s.String()
}

type CreatePropertyRequestHeader struct {
}

func (s CreatePropertyRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s CreatePropertyRequestHeader) GoString() string {
  return s.String()
}

type CreatePropertyPaths struct {
}

func (s CreatePropertyPaths) String() string {
  return tea.Prettify(s)
}

func (s CreatePropertyPaths) GoString() string {
  return s.String()
}

type CreatePropertyParameters struct {
}

func (s CreatePropertyParameters) String() string {
  return tea.Prettify(s)
}

func (s CreatePropertyParameters) GoString() string {
  return s.String()
}

type CreatePropertyResponse struct {
  // {"en":"Response code, 0 means successful.","zh_CN":"接口响应code，0代表成功。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response data.","zh_CN":"接口响应数据"}
  Data *CreatePropertyResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
  // {"en":"Response error message if failed.","zh_CN":"接口响应信息，success代表成功，失败则提供失败信息。"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s CreatePropertyResponse) String() string {
  return tea.Prettify(s)
}

func (s CreatePropertyResponse) GoString() string {
  return s.String()
}

func (s *CreatePropertyResponse) SetCode(v string) *CreatePropertyResponse {
  s.Code = &v
  return s
}

func (s *CreatePropertyResponse) SetData(v *CreatePropertyResponseData) *CreatePropertyResponse {
  s.Data = v
  return s
}

func (s *CreatePropertyResponse) SetMessage(v string) *CreatePropertyResponse {
  s.Message = &v
  return s
}

type CreatePropertyResponseData struct {
  // {"en":"Property version.","zh_CN":"项目的版本"}
  PropertyVersion *int `json:"propertyVersion,omitempty" xml:"propertyVersion,omitempty" require:"true"`
  // {"en":"Property name.","zh_CN":"项目的名称"}
  PropertyName *string `json:"propertyName,omitempty" xml:"propertyName,omitempty" require:"true"`
  // {"en":"Property ID","zh_CN":"项目标识"}
  PropertyId *int64 `json:"propertyId,omitempty" xml:"propertyId,omitempty" require:"true"`
}

func (s CreatePropertyResponseData) String() string {
  return tea.Prettify(s)
}

func (s CreatePropertyResponseData) GoString() string {
  return s.String()
}

func (s *CreatePropertyResponseData) SetPropertyVersion(v int) *CreatePropertyResponseData {
  s.PropertyVersion = &v
  return s
}

func (s *CreatePropertyResponseData) SetPropertyName(v string) *CreatePropertyResponseData {
  s.PropertyName = &v
  return s
}

func (s *CreatePropertyResponseData) SetPropertyId(v int64) *CreatePropertyResponseData {
  s.PropertyId = &v
  return s
}

type CreatePropertyResponseHeader struct {
}

func (s CreatePropertyResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s CreatePropertyResponseHeader) GoString() string {
  return s.String()
}




type ListPropertiesRequest struct {
}

func (s ListPropertiesRequest) String() string {
  return tea.Prettify(s)
}

func (s ListPropertiesRequest) GoString() string {
  return s.String()
}

type ListPropertiesRequestHeader struct {
}

func (s ListPropertiesRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ListPropertiesRequestHeader) GoString() string {
  return s.String()
}

type ListPropertiesPaths struct {
}

func (s ListPropertiesPaths) String() string {
  return tea.Prettify(s)
}

func (s ListPropertiesPaths) GoString() string {
  return s.String()
}

type ListPropertiesParameters struct {
  // {"en":"Unique identifier for the product","zh_CN":"服务类型","exampleValue":"wsa,wsa-https"}
  ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty"`
  // {"en":"Deployment Environment","zh_CN":"加速项目的部署环境","exampleValue":"staging,production"}
  Target *string `json:"target,omitempty" xml:"target,omitempty"`
  // {"defaultValue":"0","en":"Indicates the first item to return.","zh_CN":"查询起始位置，取值范围：>= 0"}
  Offset *int `json:"offset,omitempty" xml:"offset,omitempty"`
  // {"defaultValue":"100","en":"Maximum number of properties to return.  Range: <= 200","zh_CN":"每次查询的最大条数。取值范围: <= 200"}
  Limit *int `json:"limit,omitempty" xml:"limit,omitempty"`
  // {"defaultValue":"desc","en":"Order of properties to return.","zh_CN":"返回结果的顺序。默认按最后更新时间降序。","exampleValue":"asc,desc"}
  SortOrder *string `json:"sortOrder,omitempty" xml:"sortOrder,omitempty"`
  // {"defaultValue":"lastUpdateTime","en":"Returns results in sorted order.","zh_CN":"返回结果的排序依据。","exampleValue":"creationTime,lastUpdateTime"}
  SortBy *string `json:"sortBy,omitempty" xml:"sortBy,omitempty"`
  // {"en":"The id of contract, such as 40015677","zh_CN":"合同号，如40015677"}
  ContractId *string `json:"contractId,omitempty" xml:"contractId,omitempty"`
  // {"en":"The id of product, such as 10","zh_CN":"产品号，如10"}
  ItemId *string `json:"itemId,omitempty" xml:"itemId,omitempty"`
}

func (s ListPropertiesParameters) String() string {
  return tea.Prettify(s)
}

func (s ListPropertiesParameters) GoString() string {
  return s.String()
}

func (s *ListPropertiesParameters) SetServiceType(v string) *ListPropertiesParameters {
  s.ServiceType = &v
  return s
}

func (s *ListPropertiesParameters) SetTarget(v string) *ListPropertiesParameters {
  s.Target = &v
  return s
}

func (s *ListPropertiesParameters) SetOffset(v int) *ListPropertiesParameters {
  s.Offset = &v
  return s
}

func (s *ListPropertiesParameters) SetLimit(v int) *ListPropertiesParameters {
  s.Limit = &v
  return s
}

func (s *ListPropertiesParameters) SetSortOrder(v string) *ListPropertiesParameters {
  s.SortOrder = &v
  return s
}

func (s *ListPropertiesParameters) SetSortBy(v string) *ListPropertiesParameters {
  s.SortBy = &v
  return s
}

func (s *ListPropertiesParameters) SetContractId(v string) *ListPropertiesParameters {
  s.ContractId = &v
  return s
}

func (s *ListPropertiesParameters) SetItemId(v string) *ListPropertiesParameters {
  s.ItemId = &v
  return s
}

type ListPropertiesResponse struct {
  // {"en":"Response code, 0 means successful.","zh_CN":"接口响应code，0代表成功。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response error message if failed.","zh_CN":"接口响应信息，success代表成功，失败则提供失败信息。"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Response data.","zh_CN":"接口响应数据"}
  Data *ListPropertiesResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
}

func (s ListPropertiesResponse) String() string {
  return tea.Prettify(s)
}

func (s ListPropertiesResponse) GoString() string {
  return s.String()
}

func (s *ListPropertiesResponse) SetCode(v string) *ListPropertiesResponse {
  s.Code = &v
  return s
}

func (s *ListPropertiesResponse) SetMessage(v string) *ListPropertiesResponse {
  s.Message = &v
  return s
}

func (s *ListPropertiesResponse) SetData(v *ListPropertiesResponseData) *ListPropertiesResponse {
  s.Data = v
  return s
}

type ListPropertiesResponseData struct {
  // {"en":"Number of properties.","zh_CN":"项目数量。"}
  Count *int `json:"count,omitempty" xml:"count,omitempty" require:"true"`
  // {"en":"List of properties.","zh_CN":"项目列表。"}
  Properties []*ListPropertiesResponseDataProperties `json:"properties,omitempty" xml:"properties,omitempty" require:"true" type:"Repeated"`
}

func (s ListPropertiesResponseData) String() string {
  return tea.Prettify(s)
}

func (s ListPropertiesResponseData) GoString() string {
  return s.String()
}

func (s *ListPropertiesResponseData) SetCount(v int) *ListPropertiesResponseData {
  s.Count = &v
  return s
}

func (s *ListPropertiesResponseData) SetProperties(v []*ListPropertiesResponseDataProperties) *ListPropertiesResponseData {
  s.Properties = v
  return s
}

type ListPropertiesResponseDataProperties struct     {
  // {"en":"Property ID","zh_CN":"项目标识"}
  PropertyId *int64 `json:"propertyId,omitempty" xml:"propertyId,omitempty" require:"true"`
  // {"en":"Name of the property.","zh_CN":"项目的名称"}
  PropertyName *string `json:"propertyName,omitempty" xml:"propertyName,omitempty" require:"true"`
  // {"en":"A description of the property.","zh_CN":"项目的描述。"}
  PropertyComment *string `json:"propertyComment,omitempty" xml:"propertyComment,omitempty" require:"true"`
  // {"en":"Unique identifier for the product.","zh_CN":"服务类型","exampleValue":"wsa,wsa-https"}
  ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty" require:"true"`
  // {"en":"RFC3339 format date indicating when the property was created.","zh_CN":"RFC 3339格式的日期，表示创建项目的时间。"}
  CreationTime *string `json:"creationTime,omitempty" xml:"creationTime,omitempty" require:"true"`
  // {"en":"RFC3339 date indicating when the property was last updated.","zh_CN":"RFC 3339格式的日期，表示项目的最近更新时间。"}
  LastUpdateTime *string `json:"lastUpdateTime,omitempty" xml:"lastUpdateTime,omitempty" require:"true"`
  // {"en":"Latest version of the property.","zh_CN":"项目的最新版本。"}
  LatestVersion *int `json:"latestVersion,omitempty" xml:"latestVersion,omitempty" require:"true"`
  // {"en":"Describes the version of the property deployed to staging.","zh_CN":"描述已部署到演练环境的项目版本。"}
  StagingVersion *ListPropertiesResponseDataPropertiesStagingVersion `json:"stagingVersion,omitempty" xml:"stagingVersion,omitempty" require:"true" type:"Struct"`
  // {"en":"Describes the version of the property deployed to production.","zh_CN":"描述已部署到生产环境的项目版本。"}
  ProductionVersion *ListPropertiesResponseDataPropertiesProductionVersion `json:"productionVersion,omitempty" xml:"productionVersion,omitempty" require:"true" type:"Struct"`
  // {"en":"Describes the version of the property deploying to staging.","zh_CN":"描述正在部署到演练环境的项目版本。"}
  StagingDeployingVersion *ListPropertiesResponseDataPropertiesStagingDeployingVersion `json:"stagingDeployingVersion,omitempty" xml:"stagingDeployingVersion,omitempty" require:"true" type:"Struct"`
  // {"en":"Describes the version of the property deploying to production.","zh_CN":"描述正在部署到生产环境的项目版本。"}
  ProductionDeployingVersion *ListPropertiesResponseDataPropertiesProductionDeployingVersion `json:"productionDeployingVersion,omitempty" xml:"productionDeployingVersion,omitempty" require:"true" type:"Struct"`
  // {"en":"The id of contract","zh_CN":"合同号"}
  ContractId *string `json:"contractId,omitempty" xml:"contractId,omitempty" require:"true"`
  // {"en":"The id of product","zh_CN":"产品号"}
  ItemId *string `json:"itemId,omitempty" xml:"itemId,omitempty" require:"true"`
}

func (s ListPropertiesResponseDataProperties) String() string {
  return tea.Prettify(s)
}

func (s ListPropertiesResponseDataProperties) GoString() string {
  return s.String()
}

func (s *ListPropertiesResponseDataProperties) SetPropertyId(v int64) *ListPropertiesResponseDataProperties {
  s.PropertyId = &v
  return s
}

func (s *ListPropertiesResponseDataProperties) SetPropertyName(v string) *ListPropertiesResponseDataProperties {
  s.PropertyName = &v
  return s
}

func (s *ListPropertiesResponseDataProperties) SetPropertyComment(v string) *ListPropertiesResponseDataProperties {
  s.PropertyComment = &v
  return s
}

func (s *ListPropertiesResponseDataProperties) SetServiceType(v string) *ListPropertiesResponseDataProperties {
  s.ServiceType = &v
  return s
}

func (s *ListPropertiesResponseDataProperties) SetCreationTime(v string) *ListPropertiesResponseDataProperties {
  s.CreationTime = &v
  return s
}

func (s *ListPropertiesResponseDataProperties) SetLastUpdateTime(v string) *ListPropertiesResponseDataProperties {
  s.LastUpdateTime = &v
  return s
}

func (s *ListPropertiesResponseDataProperties) SetLatestVersion(v int) *ListPropertiesResponseDataProperties {
  s.LatestVersion = &v
  return s
}

func (s *ListPropertiesResponseDataProperties) SetStagingVersion(v *ListPropertiesResponseDataPropertiesStagingVersion) *ListPropertiesResponseDataProperties {
  s.StagingVersion = v
  return s
}

func (s *ListPropertiesResponseDataProperties) SetProductionVersion(v *ListPropertiesResponseDataPropertiesProductionVersion) *ListPropertiesResponseDataProperties {
  s.ProductionVersion = v
  return s
}

func (s *ListPropertiesResponseDataProperties) SetStagingDeployingVersion(v *ListPropertiesResponseDataPropertiesStagingDeployingVersion) *ListPropertiesResponseDataProperties {
  s.StagingDeployingVersion = v
  return s
}

func (s *ListPropertiesResponseDataProperties) SetProductionDeployingVersion(v *ListPropertiesResponseDataPropertiesProductionDeployingVersion) *ListPropertiesResponseDataProperties {
  s.ProductionDeployingVersion = v
  return s
}

func (s *ListPropertiesResponseDataProperties) SetContractId(v string) *ListPropertiesResponseDataProperties {
  s.ContractId = &v
  return s
}

func (s *ListPropertiesResponseDataProperties) SetItemId(v string) *ListPropertiesResponseDataProperties {
  s.ItemId = &v
  return s
}

type ListPropertiesResponseDataPropertiesStagingVersion struct {
  // {"en":"Version of the property.","zh_CN":"项目的版本"}
  Version *int `json:"version,omitempty" xml:"version,omitempty" require:"true"`
  // {"en":"hostnames.","zh_CN":"域名列表"}
  Hostnames []*string `json:"hostnames,omitempty" xml:"hostnames,omitempty" require:"true" type:"Repeated"`
}

func (s ListPropertiesResponseDataPropertiesStagingVersion) String() string {
  return tea.Prettify(s)
}

func (s ListPropertiesResponseDataPropertiesStagingVersion) GoString() string {
  return s.String()
}

func (s *ListPropertiesResponseDataPropertiesStagingVersion) SetVersion(v int) *ListPropertiesResponseDataPropertiesStagingVersion {
  s.Version = &v
  return s
}

func (s *ListPropertiesResponseDataPropertiesStagingVersion) SetHostnames(v []*string) *ListPropertiesResponseDataPropertiesStagingVersion {
  s.Hostnames = v
  return s
}

type ListPropertiesResponseDataPropertiesProductionVersion struct {
  // {"en":"Version of the property.","zh_CN":"项目的版本"}
  Version *int `json:"version,omitempty" xml:"version,omitempty" require:"true"`
  // {"en":"hostnames.","zh_CN":"域名列表"}
  Hostnames []*string `json:"hostnames,omitempty" xml:"hostnames,omitempty" require:"true" type:"Repeated"`
}

func (s ListPropertiesResponseDataPropertiesProductionVersion) String() string {
  return tea.Prettify(s)
}

func (s ListPropertiesResponseDataPropertiesProductionVersion) GoString() string {
  return s.String()
}

func (s *ListPropertiesResponseDataPropertiesProductionVersion) SetVersion(v int) *ListPropertiesResponseDataPropertiesProductionVersion {
  s.Version = &v
  return s
}

func (s *ListPropertiesResponseDataPropertiesProductionVersion) SetHostnames(v []*string) *ListPropertiesResponseDataPropertiesProductionVersion {
  s.Hostnames = v
  return s
}

type ListPropertiesResponseDataPropertiesStagingDeployingVersion struct {
  // {"en":"Version of the property.","zh_CN":"项目的版本"}
  Version *int `json:"version,omitempty" xml:"version,omitempty" require:"true"`
  // {"en":"hostnames.","zh_CN":"域名列表"}
  Hostnames []*string `json:"hostnames,omitempty" xml:"hostnames,omitempty" require:"true" type:"Repeated"`
}

func (s ListPropertiesResponseDataPropertiesStagingDeployingVersion) String() string {
  return tea.Prettify(s)
}

func (s ListPropertiesResponseDataPropertiesStagingDeployingVersion) GoString() string {
  return s.String()
}

func (s *ListPropertiesResponseDataPropertiesStagingDeployingVersion) SetVersion(v int) *ListPropertiesResponseDataPropertiesStagingDeployingVersion {
  s.Version = &v
  return s
}

func (s *ListPropertiesResponseDataPropertiesStagingDeployingVersion) SetHostnames(v []*string) *ListPropertiesResponseDataPropertiesStagingDeployingVersion {
  s.Hostnames = v
  return s
}

type ListPropertiesResponseDataPropertiesProductionDeployingVersion struct {
  // {"en":"Version of the property.","zh_CN":"项目的版本"}
  Version *int `json:"version,omitempty" xml:"version,omitempty" require:"true"`
  // {"en":"hostnames.","zh_CN":"域名列表"}
  Hostnames []*string `json:"hostnames,omitempty" xml:"hostnames,omitempty" require:"true" type:"Repeated"`
}

func (s ListPropertiesResponseDataPropertiesProductionDeployingVersion) String() string {
  return tea.Prettify(s)
}

func (s ListPropertiesResponseDataPropertiesProductionDeployingVersion) GoString() string {
  return s.String()
}

func (s *ListPropertiesResponseDataPropertiesProductionDeployingVersion) SetVersion(v int) *ListPropertiesResponseDataPropertiesProductionDeployingVersion {
  s.Version = &v
  return s
}

func (s *ListPropertiesResponseDataPropertiesProductionDeployingVersion) SetHostnames(v []*string) *ListPropertiesResponseDataPropertiesProductionDeployingVersion {
  s.Hostnames = v
  return s
}

type ListPropertiesResponseHeader struct {
}

func (s ListPropertiesResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ListPropertiesResponseHeader) GoString() string {
  return s.String()
}




type GetPropertyRequest struct {
}

func (s GetPropertyRequest) String() string {
  return tea.Prettify(s)
}

func (s GetPropertyRequest) GoString() string {
  return s.String()
}

type GetPropertyRequestHeader struct {
}

func (s GetPropertyRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetPropertyRequestHeader) GoString() string {
  return s.String()
}

type GetPropertyPaths struct {
  // {"en":"Property ID","zh_CN":"项目标识"}
  PropertyId *int `json:"propertyId,omitempty" xml:"propertyId,omitempty" require:"true"`
}

func (s GetPropertyPaths) String() string {
  return tea.Prettify(s)
}

func (s GetPropertyPaths) GoString() string {
  return s.String()
}

func (s *GetPropertyPaths) SetPropertyId(v int) *GetPropertyPaths {
  s.PropertyId = &v
  return s
}

type GetPropertyParameters struct {
}

func (s GetPropertyParameters) String() string {
  return tea.Prettify(s)
}

func (s GetPropertyParameters) GoString() string {
  return s.String()
}

type GetPropertyResponse struct {
  // {"en":"Response code, 0 means successful.","zh_CN":"接口响应code，0代表成功。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response error message if failed.","zh_CN":"接口响应信息，success代表成功，失败则提供失败信息。"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Response data.","zh_CN":"接口响应数据"}
  Data *GetPropertyResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
}

func (s GetPropertyResponse) String() string {
  return tea.Prettify(s)
}

func (s GetPropertyResponse) GoString() string {
  return s.String()
}

func (s *GetPropertyResponse) SetCode(v string) *GetPropertyResponse {
  s.Code = &v
  return s
}

func (s *GetPropertyResponse) SetMessage(v string) *GetPropertyResponse {
  s.Message = &v
  return s
}

func (s *GetPropertyResponse) SetData(v *GetPropertyResponseData) *GetPropertyResponse {
  s.Data = v
  return s
}

type GetPropertyResponseData struct {
  // {"en":"Property ID","zh_CN":"项目标识"}
  PropertyId *int64 `json:"propertyId,omitempty" xml:"propertyId,omitempty" require:"true"`
  // {"en":"Name of the property.","zh_CN":"项目的名称"}
  PropertyName *string `json:"propertyName,omitempty" xml:"propertyName,omitempty" require:"true"`
  // {"en":"A description of the property.","zh_CN":"项目的描述。"}
  PropertyComment *string `json:"propertyComment,omitempty" xml:"propertyComment,omitempty" require:"true"`
  // {"en":"Unique identifier for the product.","zh_CN":"服务类型","exampleValue":"wsa,wsa-https"}
  ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty" require:"true"`
  // {"en":"RFC3339 format date indicating when the property was created.","zh_CN":"RFC 3339格式的日期，表示创建项目的时间。"}
  CreationTime *string `json:"creationTime,omitempty" xml:"creationTime,omitempty" require:"true"`
  // {"en":"RFC3339 date indicating when the property was last updated.","zh_CN":"RFC 3339格式的日期，表示项目的最近更新时间。"}
  LastUpdateTime *string `json:"lastUpdateTime,omitempty" xml:"lastUpdateTime,omitempty" require:"true"`
  // {"en":"Latest version of the property.","zh_CN":"项目的最新版本。"}
  LatestVersion *int `json:"latestVersion,omitempty" xml:"latestVersion,omitempty" require:"true"`
  // {"en":"Describes the version of the property deployed to staging.","zh_CN":"描述已部署到演练环境的项目版本。"}
  StagingVersion *GetPropertyResponseDataStagingVersion `json:"stagingVersion,omitempty" xml:"stagingVersion,omitempty" require:"true" type:"Struct"`
  // {"en":"Describes the version of the property deployed to production.","zh_CN":"描述已部署到生产环境的项目版本。"}
  ProductionVersion *GetPropertyResponseDataProductionVersion `json:"productionVersion,omitempty" xml:"productionVersion,omitempty" require:"true" type:"Struct"`
  // {"en":"Describes the version of the property deploying to staging.","zh_CN":"描述正在部署到演练环境的项目版本。"}
  StagingDeployingVersion *GetPropertyResponseDataStagingDeployingVersion `json:"stagingDeployingVersion,omitempty" xml:"stagingDeployingVersion,omitempty" require:"true" type:"Struct"`
  // {"en":"Describes the version of the property deploying to production.","zh_CN":"描述正在部署到生产环境的项目版本。"}
  ProductionDeployingVersion *GetPropertyResponseDataProductionDeployingVersion `json:"productionDeployingVersion,omitempty" xml:"productionDeployingVersion,omitempty" require:"true" type:"Struct"`
  // {"en":"The id of contract, such as 40015677","zh_CN":"合同号，如40015677"}
  ContractId *string `json:"contractId,omitempty" xml:"contractId,omitempty" require:"true"`
  // {"en":"The id of product, such as 10","zh_CN":"产品号，如10"}
  ItemId *string `json:"itemId,omitempty" xml:"itemId,omitempty" require:"true"`
}

func (s GetPropertyResponseData) String() string {
  return tea.Prettify(s)
}

func (s GetPropertyResponseData) GoString() string {
  return s.String()
}

func (s *GetPropertyResponseData) SetPropertyId(v int64) *GetPropertyResponseData {
  s.PropertyId = &v
  return s
}

func (s *GetPropertyResponseData) SetPropertyName(v string) *GetPropertyResponseData {
  s.PropertyName = &v
  return s
}

func (s *GetPropertyResponseData) SetPropertyComment(v string) *GetPropertyResponseData {
  s.PropertyComment = &v
  return s
}

func (s *GetPropertyResponseData) SetServiceType(v string) *GetPropertyResponseData {
  s.ServiceType = &v
  return s
}

func (s *GetPropertyResponseData) SetCreationTime(v string) *GetPropertyResponseData {
  s.CreationTime = &v
  return s
}

func (s *GetPropertyResponseData) SetLastUpdateTime(v string) *GetPropertyResponseData {
  s.LastUpdateTime = &v
  return s
}

func (s *GetPropertyResponseData) SetLatestVersion(v int) *GetPropertyResponseData {
  s.LatestVersion = &v
  return s
}

func (s *GetPropertyResponseData) SetStagingVersion(v *GetPropertyResponseDataStagingVersion) *GetPropertyResponseData {
  s.StagingVersion = v
  return s
}

func (s *GetPropertyResponseData) SetProductionVersion(v *GetPropertyResponseDataProductionVersion) *GetPropertyResponseData {
  s.ProductionVersion = v
  return s
}

func (s *GetPropertyResponseData) SetStagingDeployingVersion(v *GetPropertyResponseDataStagingDeployingVersion) *GetPropertyResponseData {
  s.StagingDeployingVersion = v
  return s
}

func (s *GetPropertyResponseData) SetProductionDeployingVersion(v *GetPropertyResponseDataProductionDeployingVersion) *GetPropertyResponseData {
  s.ProductionDeployingVersion = v
  return s
}

func (s *GetPropertyResponseData) SetContractId(v string) *GetPropertyResponseData {
  s.ContractId = &v
  return s
}

func (s *GetPropertyResponseData) SetItemId(v string) *GetPropertyResponseData {
  s.ItemId = &v
  return s
}

type GetPropertyResponseDataStagingVersion struct {
  // {"en":"Version of the property.","zh_CN":"项目的版本"}
  Version *int `json:"version,omitempty" xml:"version,omitempty" require:"true"`
  // {"en":"hostnames.","zh_CN":"域名列表"}
  Hostnames []*string `json:"hostnames,omitempty" xml:"hostnames,omitempty" require:"true" type:"Repeated"`
}

func (s GetPropertyResponseDataStagingVersion) String() string {
  return tea.Prettify(s)
}

func (s GetPropertyResponseDataStagingVersion) GoString() string {
  return s.String()
}

func (s *GetPropertyResponseDataStagingVersion) SetVersion(v int) *GetPropertyResponseDataStagingVersion {
  s.Version = &v
  return s
}

func (s *GetPropertyResponseDataStagingVersion) SetHostnames(v []*string) *GetPropertyResponseDataStagingVersion {
  s.Hostnames = v
  return s
}

type GetPropertyResponseDataProductionVersion struct {
  // {"en":"Version of the property.","zh_CN":"项目的版本"}
  Version *int `json:"version,omitempty" xml:"version,omitempty" require:"true"`
  // {"en":"hostnames.","zh_CN":"域名列表"}
  Hostnames []*string `json:"hostnames,omitempty" xml:"hostnames,omitempty" require:"true" type:"Repeated"`
}

func (s GetPropertyResponseDataProductionVersion) String() string {
  return tea.Prettify(s)
}

func (s GetPropertyResponseDataProductionVersion) GoString() string {
  return s.String()
}

func (s *GetPropertyResponseDataProductionVersion) SetVersion(v int) *GetPropertyResponseDataProductionVersion {
  s.Version = &v
  return s
}

func (s *GetPropertyResponseDataProductionVersion) SetHostnames(v []*string) *GetPropertyResponseDataProductionVersion {
  s.Hostnames = v
  return s
}

type GetPropertyResponseDataStagingDeployingVersion struct {
  // {"en":"Version of the property.","zh_CN":"项目的版本"}
  Version *int `json:"version,omitempty" xml:"version,omitempty" require:"true"`
  // {"en":"hostnames.","zh_CN":"域名列表"}
  Hostnames []*string `json:"hostnames,omitempty" xml:"hostnames,omitempty" require:"true" type:"Repeated"`
}

func (s GetPropertyResponseDataStagingDeployingVersion) String() string {
  return tea.Prettify(s)
}

func (s GetPropertyResponseDataStagingDeployingVersion) GoString() string {
  return s.String()
}

func (s *GetPropertyResponseDataStagingDeployingVersion) SetVersion(v int) *GetPropertyResponseDataStagingDeployingVersion {
  s.Version = &v
  return s
}

func (s *GetPropertyResponseDataStagingDeployingVersion) SetHostnames(v []*string) *GetPropertyResponseDataStagingDeployingVersion {
  s.Hostnames = v
  return s
}

type GetPropertyResponseDataProductionDeployingVersion struct {
  // {"en":"Version of the property.","zh_CN":"项目的版本"}
  Version *int `json:"version,omitempty" xml:"version,omitempty" require:"true"`
  // {"en":"hostnames.","zh_CN":"域名列表"}
  Hostnames []*string `json:"hostnames,omitempty" xml:"hostnames,omitempty" require:"true" type:"Repeated"`
}

func (s GetPropertyResponseDataProductionDeployingVersion) String() string {
  return tea.Prettify(s)
}

func (s GetPropertyResponseDataProductionDeployingVersion) GoString() string {
  return s.String()
}

func (s *GetPropertyResponseDataProductionDeployingVersion) SetVersion(v int) *GetPropertyResponseDataProductionDeployingVersion {
  s.Version = &v
  return s
}

func (s *GetPropertyResponseDataProductionDeployingVersion) SetHostnames(v []*string) *GetPropertyResponseDataProductionDeployingVersion {
  s.Hostnames = v
  return s
}

type GetPropertyResponseHeader struct {
}

func (s GetPropertyResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetPropertyResponseHeader) GoString() string {
  return s.String()
}




