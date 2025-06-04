package productmanagement

// This file is auto-generated, don't edit it. Thanks.
import (
  "github.com/alibabacloud-go/tea/tea"
)

type GetAListOfOriginShieldsPaths struct {
}

func (s GetAListOfOriginShieldsPaths) String() string {
  return tea.Prettify(s)
}

func (s GetAListOfOriginShieldsPaths) GoString() string {
  return s.String()
}

type GetAListOfOriginShieldsParameters struct {
  // {"en" : "The value of the search parameter is matched on the shields' name, id, regions, ipV4 , and ipV6 fields. Only those shields that match will be returned in the response.", "zh_CN": "输入搜索关键字查询shield。接口将返回名称，id，区域或ip地址中含有关键字的shield。"}
  Search *string `json:"search,omitempty" xml:"search,omitempty"`
  // {"en" : "Enum: lastUpdated,effectiveDate 
  // Default: lastUpdated 
  // Returns shields sorted by when they were last updated or when they became active.", "zh_CN": "取值范围: lastUpdated,effectiveDate 
  // 默认值: lastUpdated 
  // 按最近更新时间或生效时间对shield进行排序。"}
  SortBy *string `json:"sortBy,omitempty" xml:"sortBy,omitempty"`
  // {"en" : "Enum: asc,desc 
  // Default: desc 
  // Returns shields in sorted order, either 'asc' for ascending or 'desc' for descending. The default is to return the last shield updated first.", "zh_CN": "取值范围: asc,desc 
  // 默认值: desc 
  // 按升序或降序对shield进行排序。"}
  SortOrder *string `json:"sortOrder,omitempty" xml:"sortOrder,omitempty"`
  // {"en" : "Enum: false,true 
  // Filter results based on whether shields are used in properties deployed to staging or production. If unspecified, all available shields are returned.", "zh_CN": "取值范围: false,true 
  // 按照shield是否在已部署的加速项目中使用来查询。默认查询所有shield。"}
  UsedInProperties *string `json:"usedInProperties,omitempty" xml:"usedInProperties,omitempty"`
}

func (s GetAListOfOriginShieldsParameters) String() string {
  return tea.Prettify(s)
}

func (s GetAListOfOriginShieldsParameters) GoString() string {
  return s.String()
}

func (s *GetAListOfOriginShieldsParameters) SetSearch(v string) *GetAListOfOriginShieldsParameters {
  s.Search = &v
  return s
}

func (s *GetAListOfOriginShieldsParameters) SetSortBy(v string) *GetAListOfOriginShieldsParameters {
  s.SortBy = &v
  return s
}

func (s *GetAListOfOriginShieldsParameters) SetSortOrder(v string) *GetAListOfOriginShieldsParameters {
  s.SortOrder = &v
  return s
}

func (s *GetAListOfOriginShieldsParameters) SetUsedInProperties(v string) *GetAListOfOriginShieldsParameters {
  s.UsedInProperties = &v
  return s
}

type GetAListOfOriginShieldsRequestHeader struct {
}

func (s GetAListOfOriginShieldsRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetAListOfOriginShieldsRequestHeader) GoString() string {
  return s.String()
}

type GetAListOfOriginShieldsRequest struct {
}

func (s GetAListOfOriginShieldsRequest) String() string {
  return tea.Prettify(s)
}

func (s GetAListOfOriginShieldsRequest) GoString() string {
  return s.String()
}

type GetAListOfOriginShieldsResponseHeader struct {
}

func (s GetAListOfOriginShieldsResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetAListOfOriginShieldsResponseHeader) GoString() string {
  return s.String()
}

type GetAListOfOriginShieldsResponse struct {
  // {"en" : "A list of shields that can be used in a property.", "zh_CN": "shield列表。"}
  Shields []*GetAListOfOriginShieldsResponseShields `json:"shields,omitempty" xml:"shields,omitempty" require:"true" type:"Repeated"`
}

func (s GetAListOfOriginShieldsResponse) String() string {
  return tea.Prettify(s)
}

func (s GetAListOfOriginShieldsResponse) GoString() string {
  return s.String()
}

func (s *GetAListOfOriginShieldsResponse) SetShields(v []*GetAListOfOriginShieldsResponseShields) *GetAListOfOriginShieldsResponse {
  s.Shields = v
  return s
}

type GetAListOfOriginShieldsResponseShields struct     {
  // {"en" : "Name of the shield.", "zh_CN": "shield名称。"}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // {"en" : "A unique identifier representing the shield. This ID is referenced in properties using the shield.", "zh_CN": "shield唯一标识的ID。在加速项目中以该ID引用shield。"}
  Id *string `json:"id,omitempty" xml:"id,omitempty"`
  // {"en" : "Indicates the region of the origin shield. It is typically an ISO-3166-1 code representing a country. When choosing an origin shield for your property, it is best to select one closer to your origin servers.
  // Use of a shield located in the 'cn' region (China) requires your property to have 'hasBeian' set to true which means all of its hostnames have an ICP Beian license required by the Chinese government.", "zh_CN": "shield所在区域。通常以ISO-3166-1的国家码表示。当您选择shield时，建议选择离源站最近的shield，以达到最佳效果。如您需要使用位于中国区域的shield，则您的加速项目配置中是否备案需要设为'是'，您相关的加速域名必须已取得备案。"}
  Region *string `json:"region,omitempty" xml:"region,omitempty"`
  // {"en" : "A list of IPv4 addresses used by the shield. If your origin server has an access control list, include these IP addresses.", "zh_CN": "shield所使用到的IPv4地址。如果您的源站开启了访问控制，需要对此处的IP列表进行加白。"}
  IpV4 []*string `json:"ipV4,omitempty" xml:"ipV4,omitempty" type:"Repeated"`
  // {"en" : "A list of IPv6 addresses used by the shield. If your origin server has an access control list, include these IP addresses.", "zh_CN": "shield所使用到的IPv6地址。如果您的源站开启了访问控制，需要对此处的IP列表进行加白。"}
  IpV6 []*string `json:"ipV6,omitempty" xml:"ipV6,omitempty" type:"Repeated"`
  // {"en" : "An RFC 3339 date indicating when the shield was last updated.", "zh_CN": "shield最近一次更新的时间，以RFC 3339日期格式表示。"}
  LastUpdated *string `json:"lastUpdated,omitempty" xml:"lastUpdated,omitempty"`
  // {"en" : "An RFC 3339 date indicating when the shield became active.", "zh_CN": "shield的生效时间，以RFC 3339日期格式表示。"}
  EffectiveDate *string `json:"effectiveDate,omitempty" xml:"effectiveDate,omitempty"`
  // {"en" : "Whether the shield is used in any properties deployed to staging or production.", "zh_CN": "该shield是否在已部署的加速项目中使用。"}
  UsedInProperties *bool `json:"usedInProperties,omitempty" xml:"usedInProperties,omitempty"`
}

func (s GetAListOfOriginShieldsResponseShields) String() string {
  return tea.Prettify(s)
}

func (s GetAListOfOriginShieldsResponseShields) GoString() string {
  return s.String()
}

func (s *GetAListOfOriginShieldsResponseShields) SetName(v string) *GetAListOfOriginShieldsResponseShields {
  s.Name = &v
  return s
}

func (s *GetAListOfOriginShieldsResponseShields) SetId(v string) *GetAListOfOriginShieldsResponseShields {
  s.Id = &v
  return s
}

func (s *GetAListOfOriginShieldsResponseShields) SetRegion(v string) *GetAListOfOriginShieldsResponseShields {
  s.Region = &v
  return s
}

func (s *GetAListOfOriginShieldsResponseShields) SetIpV4(v []*string) *GetAListOfOriginShieldsResponseShields {
  s.IpV4 = v
  return s
}

func (s *GetAListOfOriginShieldsResponseShields) SetIpV6(v []*string) *GetAListOfOriginShieldsResponseShields {
  s.IpV6 = v
  return s
}

func (s *GetAListOfOriginShieldsResponseShields) SetLastUpdated(v string) *GetAListOfOriginShieldsResponseShields {
  s.LastUpdated = &v
  return s
}

func (s *GetAListOfOriginShieldsResponseShields) SetEffectiveDate(v string) *GetAListOfOriginShieldsResponseShields {
  s.EffectiveDate = &v
  return s
}

func (s *GetAListOfOriginShieldsResponseShields) SetUsedInProperties(v bool) *GetAListOfOriginShieldsResponseShields {
  s.UsedInProperties = &v
  return s
}




type GetSystemConfigurationPaths struct {
}

func (s GetSystemConfigurationPaths) String() string {
  return tea.Prettify(s)
}

func (s GetSystemConfigurationPaths) GoString() string {
  return s.String()
}

type GetSystemConfigurationParameters struct {
}

func (s GetSystemConfigurationParameters) String() string {
  return tea.Prettify(s)
}

func (s GetSystemConfigurationParameters) GoString() string {
  return s.String()
}

type GetSystemConfigurationRequestHeader struct {
}

func (s GetSystemConfigurationRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetSystemConfigurationRequestHeader) GoString() string {
  return s.String()
}

type GetSystemConfigurationRequest struct {
}

func (s GetSystemConfigurationRequest) String() string {
  return tea.Prettify(s)
}

func (s GetSystemConfigurationRequest) GoString() string {
  return s.String()
}

type GetSystemConfigurationResponseHeader struct {
}

func (s GetSystemConfigurationResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetSystemConfigurationResponseHeader) GoString() string {
  return s.String()
}

type GetSystemConfigurationResponse struct {
  // {"en" : "The minimum version of the cache server required to validate and deploy properties.", "zh_CN": "Cache服务器最低版本。低于该版本，加速项目的验证和部署可能无法正常工作。"}
  MinimumSupportedCacheVersion *string `json:"minimumSupportedCacheVersion,omitempty" xml:"minimumSupportedCacheVersion,omitempty" require:"true"`
  // {"en" : "A list of Edge Logic directives that are available to everyone.", "zh_CN": "边缘逻辑基础指令列表。这些指令默认对所有客户开放。"}
  BaseDirectives []*string `json:"baseDirectives,omitempty" xml:"baseDirectives,omitempty" require:"true" type:"Repeated"`
  // {"en" : "A list of advanced Edge Logic directives that can be enabled by contacting our support team.", "zh_CN": "边缘逻辑高级指令列表。这些指令需联系我们的技术支持团队开通。"}
  AdvancedDirectives []*string `json:"advancedDirectives,omitempty" xml:"advancedDirectives,omitempty" require:"true" type:"Repeated"`
  // {"en" : "This field lists nonstandard ports that can be used to handle HTTP and HTTPS requests. If your website uses a nonstandard port which is not listed, please contact our support team. You can specify one of these port numbers in the extraServicePorts field of your property configuration.", "zh_CN": "该字段返回CDN Pro支持的HTTP, HTTPS非标准服务端口。这些端口可以用在加速项目的extraServicePorts配置项中。如果你所使用的非标准端口不在该列表中，可联系我们的技术支持团队开通。"}
  ExtraServicePorts *GetSystemConfigurationResponseExtraServicePorts `json:"extraServicePorts,omitempty" xml:"extraServicePorts,omitempty" require:"true" type:"Struct"`
  // {"en" : "List of advanced features that can be enabled for customers via the service quota API. 'realTimeLog' is an advanced feature allowing a customer to monitor requests from visitors to the content as they are received. 'nearChina' is an advanced feature allowing use of a special server group to provide better performance to visitors in China for domains without ICP Beian. 'originShield' allows use of an extra layer of servers in front of a property's origin servers. Please contact our support team if you require an advanced feature.", "zh_CN": "CDN Pro支持的高级功能列表。例如，'realTimeLog'日志可用于实时回传日志。如果需要使用这些高级功能，请联系我们的技术支持团队开通。"}
  AdvancedFeaures []*string `json:"advancedFeaures,omitempty" xml:"advancedFeaures,omitempty" require:"true" type:"Repeated"`
  // {"en" : "RFC 3339 date indicating the oldest validated property that can be deployed. Properties validated before this date must be re-validated before you can deploy them. Example: '2021-02-05T00:00:00Z'", "zh_CN": "加速项目验证通过的有效起始时间。如果加速项目的验证早于该时间，则必须重新验证通过后才能部署。采用RFC 3339日期格式，例如'2021-02-05T00:00:00Z'。"}
  OldestPropertyValidationTime *string `json:"oldestPropertyValidationTime,omitempty" xml:"oldestPropertyValidationTime,omitempty" require:"true"`
  // {"en" : "List of Edge Logic directives that can be used in the loadBalancerLogic field of a property by everyone.", "zh_CN": "可用于加速项目中loadBalancerLogic字段的边缘逻辑指令列表。这些指令默认对所有客户开放。"}
  BaseLbDirectives []*string `json:"baseLbDirectives,omitempty" xml:"baseLbDirectives,omitempty" require:"true" type:"Repeated"`
  // {"en" : "A list of advanced Edge Logic directives that can be enabled for use in the loadBalancerLogic field of a property. If you require one that is not enabled for your account, please contact our support team.", "zh_CN": "可用于加速项目中loadBalancerLogic字段的边缘逻辑高级指令列表。如果需要使用这些指令，请联系我们的技术支持团队开通。"}
  AdvancedLbDirectives []*string `json:"advancedLbDirectives,omitempty" xml:"advancedLbDirectives,omitempty" require:"true" type:"Repeated"`
}

func (s GetSystemConfigurationResponse) String() string {
  return tea.Prettify(s)
}

func (s GetSystemConfigurationResponse) GoString() string {
  return s.String()
}

func (s *GetSystemConfigurationResponse) SetMinimumSupportedCacheVersion(v string) *GetSystemConfigurationResponse {
  s.MinimumSupportedCacheVersion = &v
  return s
}

func (s *GetSystemConfigurationResponse) SetBaseDirectives(v []*string) *GetSystemConfigurationResponse {
  s.BaseDirectives = v
  return s
}

func (s *GetSystemConfigurationResponse) SetAdvancedDirectives(v []*string) *GetSystemConfigurationResponse {
  s.AdvancedDirectives = v
  return s
}

func (s *GetSystemConfigurationResponse) SetExtraServicePorts(v *GetSystemConfigurationResponseExtraServicePorts) *GetSystemConfigurationResponse {
  s.ExtraServicePorts = v
  return s
}

func (s *GetSystemConfigurationResponse) SetAdvancedFeaures(v []*string) *GetSystemConfigurationResponse {
  s.AdvancedFeaures = v
  return s
}

func (s *GetSystemConfigurationResponse) SetOldestPropertyValidationTime(v string) *GetSystemConfigurationResponse {
  s.OldestPropertyValidationTime = &v
  return s
}

func (s *GetSystemConfigurationResponse) SetBaseLbDirectives(v []*string) *GetSystemConfigurationResponse {
  s.BaseLbDirectives = v
  return s
}

func (s *GetSystemConfigurationResponse) SetAdvancedLbDirectives(v []*string) *GetSystemConfigurationResponse {
  s.AdvancedLbDirectives = v
  return s
}

type GetSystemConfigurationResponseExtraServicePorts struct {
  // {"en" : "Specifies ports other than the default 80 which can be used to handle HTTP requests.", "zh_CN": "可用于监听HTTP请求的非标准端口。80端口默认支持。"}
  Http []*string `json:"http,omitempty" xml:"http,omitempty" type:"Repeated"`
  // {"en" : "Specify ports other than the default 443 which can be used to handle HTTPS requests.", "zh_CN": "可用于监听HTTPS请求的非标准端口。443端口默认支持。"}
  Https []*string `json:"https,omitempty" xml:"https,omitempty" type:"Repeated"`
}

func (s GetSystemConfigurationResponseExtraServicePorts) String() string {
  return tea.Prettify(s)
}

func (s GetSystemConfigurationResponseExtraServicePorts) GoString() string {
  return s.String()
}

func (s *GetSystemConfigurationResponseExtraServicePorts) SetHttp(v []*string) *GetSystemConfigurationResponseExtraServicePorts {
  s.Http = v
  return s
}

func (s *GetSystemConfigurationResponseExtraServicePorts) SetHttps(v []*string) *GetSystemConfigurationResponseExtraServicePorts {
  s.Https = v
  return s
}




type GetOriginFastRouteIpListPaths struct {
}

func (s GetOriginFastRouteIpListPaths) String() string {
  return tea.Prettify(s)
}

func (s GetOriginFastRouteIpListPaths) GoString() string {
  return s.String()
}

type GetOriginFastRouteIpListParameters struct {
}

func (s GetOriginFastRouteIpListParameters) String() string {
  return tea.Prettify(s)
}

func (s GetOriginFastRouteIpListParameters) GoString() string {
  return s.String()
}

type GetOriginFastRouteIpListRequestHeader struct {
}

func (s GetOriginFastRouteIpListRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetOriginFastRouteIpListRequestHeader) GoString() string {
  return s.String()
}

type GetOriginFastRouteIpListRequest struct {
}

func (s GetOriginFastRouteIpListRequest) String() string {
  return tea.Prettify(s)
}

func (s GetOriginFastRouteIpListRequest) GoString() string {
  return s.String()
}

type GetOriginFastRouteIpListResponseHeader struct {
}

func (s GetOriginFastRouteIpListResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetOriginFastRouteIpListResponseHeader) GoString() string {
  return s.String()
}

type GetOriginFastRouteIpListResponse struct {
  // {"en" : "RFC 3339 date indicating when the list of IP addresses was last updated.", "zh_CN": "IP列表最后一次更新的时间，以RFC 3339日期格式表示。"}
  LastUpdated *string `json:"lastUpdated,omitempty" xml:"lastUpdated,omitempty" require:"true"`
  // {"en" : "RFC 3339 date indicating when the IP addresses became effective.", "zh_CN": "IP列表的生效时间，以RFC 3339日期格式表示。"}
  EffectiveDate *string `json:"effectiveDate,omitempty" xml:"effectiveDate,omitempty" require:"true"`
  // {"en" : "IPv4 addresses used by origin fast route servers.", "zh_CN": "快速回源服务器的IPv4地址。"}
  IpV4 []*string `json:"ipV4,omitempty" xml:"ipV4,omitempty" require:"true" type:"Repeated"`
  // {"en" : "IPv6 addresses used by origin fast route servers.", "zh_CN": "快速回源服务器的IPv6地址。"}
  IpV6 []*string `json:"ipV6,omitempty" xml:"ipV6,omitempty" require:"true" type:"Repeated"`
}

func (s GetOriginFastRouteIpListResponse) String() string {
  return tea.Prettify(s)
}

func (s GetOriginFastRouteIpListResponse) GoString() string {
  return s.String()
}

func (s *GetOriginFastRouteIpListResponse) SetLastUpdated(v string) *GetOriginFastRouteIpListResponse {
  s.LastUpdated = &v
  return s
}

func (s *GetOriginFastRouteIpListResponse) SetEffectiveDate(v string) *GetOriginFastRouteIpListResponse {
  s.EffectiveDate = &v
  return s
}

func (s *GetOriginFastRouteIpListResponse) SetIpV4(v []*string) *GetOriginFastRouteIpListResponse {
  s.IpV4 = v
  return s
}

func (s *GetOriginFastRouteIpListResponse) SetIpV6(v []*string) *GetOriginFastRouteIpListResponse {
  s.IpV6 = v
  return s
}




type GetStagingServersListPaths struct {
}

func (s GetStagingServersListPaths) String() string {
  return tea.Prettify(s)
}

func (s GetStagingServersListPaths) GoString() string {
  return s.String()
}

type GetStagingServersListParameters struct {
}

func (s GetStagingServersListParameters) String() string {
  return tea.Prettify(s)
}

func (s GetStagingServersListParameters) GoString() string {
  return s.String()
}

type GetStagingServersListRequestHeader struct {
}

func (s GetStagingServersListRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetStagingServersListRequestHeader) GoString() string {
  return s.String()
}

type GetStagingServersListRequest struct {
}

func (s GetStagingServersListRequest) String() string {
  return tea.Prettify(s)
}

func (s GetStagingServersListRequest) GoString() string {
  return s.String()
}

type GetStagingServersListResponseHeader struct {
}

func (s GetStagingServersListResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetStagingServersListResponseHeader) GoString() string {
  return s.String()
}

type GetStagingServersListResponse struct {
  // {"en" : "Each entry describe a staging server.", "zh_CN": "每个对象代表一台服务器。"}
  StagingServers []*GetStagingServersListResponseStagingServers `json:"stagingServers,omitempty" xml:"stagingServers,omitempty" require:"true" type:"Repeated"`
}

func (s GetStagingServersListResponse) String() string {
  return tea.Prettify(s)
}

func (s GetStagingServersListResponse) GoString() string {
  return s.String()
}

func (s *GetStagingServersListResponse) SetStagingServers(v []*GetStagingServersListResponseStagingServers) *GetStagingServersListResponse {
  s.StagingServers = v
  return s
}

type GetStagingServersListResponseStagingServers struct     {
  // {"en" : "IP address. It can be IPv4 or IPv6 format.", "zh_CN": "服务器的IP地址，IPv4或IPv6格式。"}
  Ip *string `json:"ip,omitempty" xml:"ip,omitempty"`
  // {"en" : "A code representing the staging server.", "zh_CN": "服务器所在的区域的代码。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty"`
  // {"en" : "Location of the staging server.", "zh_CN": "服务器所在的区域。"}
  Location *string `json:"location,omitempty" xml:"location,omitempty"`
  // {"en" : "Enum: 4,6 
  // Indicates the format of the 'ip' field.", "zh_CN": "取值范围: 4,6 
  // IP协议版本。"}
  IpVersion *int `json:"ipVersion,omitempty" xml:"ipVersion,omitempty"`
}

func (s GetStagingServersListResponseStagingServers) String() string {
  return tea.Prettify(s)
}

func (s GetStagingServersListResponseStagingServers) GoString() string {
  return s.String()
}

func (s *GetStagingServersListResponseStagingServers) SetIp(v string) *GetStagingServersListResponseStagingServers {
  s.Ip = &v
  return s
}

func (s *GetStagingServersListResponseStagingServers) SetCode(v string) *GetStagingServersListResponseStagingServers {
  s.Code = &v
  return s
}

func (s *GetStagingServersListResponseStagingServers) SetLocation(v string) *GetStagingServersListResponseStagingServers {
  s.Location = &v
  return s
}

func (s *GetStagingServersListResponseStagingServers) SetIpVersion(v int) *GetStagingServersListResponseStagingServers {
  s.IpVersion = &v
  return s
}




type GetIpAddressesOfTheCdnPaths struct {
}

func (s GetIpAddressesOfTheCdnPaths) String() string {
  return tea.Prettify(s)
}

func (s GetIpAddressesOfTheCdnPaths) GoString() string {
  return s.String()
}

type GetIpAddressesOfTheCdnParameters struct {
}

func (s GetIpAddressesOfTheCdnParameters) String() string {
  return tea.Prettify(s)
}

func (s GetIpAddressesOfTheCdnParameters) GoString() string {
  return s.String()
}

type GetIpAddressesOfTheCdnRequestHeader struct {
}

func (s GetIpAddressesOfTheCdnRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetIpAddressesOfTheCdnRequestHeader) GoString() string {
  return s.String()
}

type GetIpAddressesOfTheCdnRequest struct {
}

func (s GetIpAddressesOfTheCdnRequest) String() string {
  return tea.Prettify(s)
}

func (s GetIpAddressesOfTheCdnRequest) GoString() string {
  return s.String()
}

type GetIpAddressesOfTheCdnResponseHeader struct {
}

func (s GetIpAddressesOfTheCdnResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetIpAddressesOfTheCdnResponseHeader) GoString() string {
  return s.String()
}

type GetIpAddressesOfTheCdnResponse struct {
  // {"en" : "RFC 3339 date indicating when the list of IP addresses was last updated.", "zh_CN": "IP地址列表最近一次更新时间，以RFC 3339格式展示。"}
  LastUpdated *string `json:"lastUpdated,omitempty" xml:"lastUpdated,omitempty" require:"true"`
  // {"en" : "RFC 3339 date indicating when the IP address list became effective.", "zh_CN": "IP地址列表生效时间，以RFC 3339格式展示。"}
  EffectiveDate *string `json:"effectiveDate,omitempty" xml:"effectiveDate,omitempty" require:"true"`
  // {"en" : "IPv4 addresses used by CDN Pro.", "zh_CN": "CDN Pro所使用的IPv4地址。"}
  IpV4 []*string `json:"ipV4,omitempty" xml:"ipV4,omitempty" require:"true" type:"Repeated"`
  // {"en" : "IPv6 addresses used by CDN Pro.", "zh_CN": "CDN Pro所使用的IPv6地址。"}
  IpV6 []*string `json:"ipV6,omitempty" xml:"ipV6,omitempty" require:"true" type:"Repeated"`
}

func (s GetIpAddressesOfTheCdnResponse) String() string {
  return tea.Prettify(s)
}

func (s GetIpAddressesOfTheCdnResponse) GoString() string {
  return s.String()
}

func (s *GetIpAddressesOfTheCdnResponse) SetLastUpdated(v string) *GetIpAddressesOfTheCdnResponse {
  s.LastUpdated = &v
  return s
}

func (s *GetIpAddressesOfTheCdnResponse) SetEffectiveDate(v string) *GetIpAddressesOfTheCdnResponse {
  s.EffectiveDate = &v
  return s
}

func (s *GetIpAddressesOfTheCdnResponse) SetIpV4(v []*string) *GetIpAddressesOfTheCdnResponse {
  s.IpV4 = v
  return s
}

func (s *GetIpAddressesOfTheCdnResponse) SetIpV6(v []*string) *GetIpAddressesOfTheCdnResponse {
  s.IpV6 = v
  return s
}




type GetAShieldByItsIdPaths struct {
  // {"en" : "ID of a shield", "zh_CN": "shield id。"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
}

func (s GetAShieldByItsIdPaths) String() string {
  return tea.Prettify(s)
}

func (s GetAShieldByItsIdPaths) GoString() string {
  return s.String()
}

func (s *GetAShieldByItsIdPaths) SetId(v string) *GetAShieldByItsIdPaths {
  s.Id = &v
  return s
}

type GetAShieldByItsIdParameters struct {
}

func (s GetAShieldByItsIdParameters) String() string {
  return tea.Prettify(s)
}

func (s GetAShieldByItsIdParameters) GoString() string {
  return s.String()
}

type GetAShieldByItsIdRequestHeader struct {
}

func (s GetAShieldByItsIdRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetAShieldByItsIdRequestHeader) GoString() string {
  return s.String()
}

type GetAShieldByItsIdRequest struct {
}

func (s GetAShieldByItsIdRequest) String() string {
  return tea.Prettify(s)
}

func (s GetAShieldByItsIdRequest) GoString() string {
  return s.String()
}

type GetAShieldByItsIdResponseHeader struct {
}

func (s GetAShieldByItsIdResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetAShieldByItsIdResponseHeader) GoString() string {
  return s.String()
}

type GetAShieldByItsIdResponse struct {
  // {"en" : "Name of the shield.", "zh_CN": "shield名称。"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en" : "A unique identifier representing the shield. This ID is referenced in properties using the shield.", "zh_CN": "shield唯一标识的ID。在加速项目中以该ID引用shield。"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {"en" : "Indicates the region of the origin shield. It is typically an ISO-3166-1 code representing a country. When choosing an origin shield for your property, it is best to select one closer to your origin servers.
  // Use of a shield located in the 'cn' region (China) requires your property to have 'hasBeian' set to true which means all of its hostnames have an ICP Beian license required by the Chinese government.", "zh_CN": "shield所在区域。通常以ISO-3166-1的国家码表示。当您选择shield时，建议选择离源站最近的shield，以达到最佳效果。如您需要使用位于中国区域的shield，则您的加速项目配置中是否备案需要设为'是'，您相关的加速域名必须已取得备案。"}
  Region *string `json:"region,omitempty" xml:"region,omitempty" require:"true"`
  // {"en" : "A list of IPv4 addresses used by the shield. If your origin server has an access control list, include these IP addresses.", "zh_CN": "shield所使用到的IPv4地址。如果您的源站开启了访问控制，需要对此处的IP列表进行加白。"}
  IpV4 []*string `json:"ipV4,omitempty" xml:"ipV4,omitempty" require:"true" type:"Repeated"`
  // {"en" : "A list of IPv6 addresses used by the shield. If your origin server has an access control list, include these IP addresses.", "zh_CN": "shield所使用到的IPv6地址。如果您的源站开启了访问控制，需要对此处的IP列表进行加白。"}
  IpV6 []*string `json:"ipV6,omitempty" xml:"ipV6,omitempty" require:"true" type:"Repeated"`
  // {"en" : "An RFC 3339 date indicating when the shield was last updated.", "zh_CN": "shield最近一次更新的时间，以RFC 3339日期格式表示。"}
  LastUpdated *string `json:"lastUpdated,omitempty" xml:"lastUpdated,omitempty" require:"true"`
  // {"en" : "An RFC 3339 date indicating when the shield became active.", "zh_CN": "shield的生效时间，以RFC 3339日期格式表示。"}
  EffectiveDate *string `json:"effectiveDate,omitempty" xml:"effectiveDate,omitempty" require:"true"`
  // {"en" : "Whether the shield is used in any properties deployed to staging or production.", "zh_CN": "该shield是否在已部署的加速项目中使用。"}
  UsedInProperties *bool `json:"usedInProperties,omitempty" xml:"usedInProperties,omitempty" require:"true"`
}

func (s GetAShieldByItsIdResponse) String() string {
  return tea.Prettify(s)
}

func (s GetAShieldByItsIdResponse) GoString() string {
  return s.String()
}

func (s *GetAShieldByItsIdResponse) SetName(v string) *GetAShieldByItsIdResponse {
  s.Name = &v
  return s
}

func (s *GetAShieldByItsIdResponse) SetId(v string) *GetAShieldByItsIdResponse {
  s.Id = &v
  return s
}

func (s *GetAShieldByItsIdResponse) SetRegion(v string) *GetAShieldByItsIdResponse {
  s.Region = &v
  return s
}

func (s *GetAShieldByItsIdResponse) SetIpV4(v []*string) *GetAShieldByItsIdResponse {
  s.IpV4 = v
  return s
}

func (s *GetAShieldByItsIdResponse) SetIpV6(v []*string) *GetAShieldByItsIdResponse {
  s.IpV6 = v
  return s
}

func (s *GetAShieldByItsIdResponse) SetLastUpdated(v string) *GetAShieldByItsIdResponse {
  s.LastUpdated = &v
  return s
}

func (s *GetAShieldByItsIdResponse) SetEffectiveDate(v string) *GetAShieldByItsIdResponse {
  s.EffectiveDate = &v
  return s
}

func (s *GetAShieldByItsIdResponse) SetUsedInProperties(v bool) *GetAShieldByItsIdResponse {
  s.UsedInProperties = &v
  return s
}




type CheckIfIpAddressesBelongToTheCdnProPlatformPaths struct {
}

func (s CheckIfIpAddressesBelongToTheCdnProPlatformPaths) String() string {
  return tea.Prettify(s)
}

func (s CheckIfIpAddressesBelongToTheCdnProPlatformPaths) GoString() string {
  return s.String()
}

type CheckIfIpAddressesBelongToTheCdnProPlatformParameters struct {
}

func (s CheckIfIpAddressesBelongToTheCdnProPlatformParameters) String() string {
  return tea.Prettify(s)
}

func (s CheckIfIpAddressesBelongToTheCdnProPlatformParameters) GoString() string {
  return s.String()
}

type CheckIfIpAddressesBelongToTheCdnProPlatformRequestHeader struct {
}

func (s CheckIfIpAddressesBelongToTheCdnProPlatformRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s CheckIfIpAddressesBelongToTheCdnProPlatformRequestHeader) GoString() string {
  return s.String()
}

type CheckIfIpAddressesBelongToTheCdnProPlatformRequest struct {
  // {"en" : "Range: <= 128 items 
  // Array of IPv4 or IPv6 addresses to check.", "zh_CN": "取值范围: <= 128 条目 
  // 需要查询的IP地址列表。支持IPv4和IPv6。"}
  IpList []*string `json:"ipList,omitempty" xml:"ipList,omitempty" type:"Repeated"`
}

func (s CheckIfIpAddressesBelongToTheCdnProPlatformRequest) String() string {
  return tea.Prettify(s)
}

func (s CheckIfIpAddressesBelongToTheCdnProPlatformRequest) GoString() string {
  return s.String()
}

func (s *CheckIfIpAddressesBelongToTheCdnProPlatformRequest) SetIpList(v []*string) *CheckIfIpAddressesBelongToTheCdnProPlatformRequest {
  s.IpList = v
  return s
}

type CheckIfIpAddressesBelongToTheCdnProPlatformResponseHeader struct {
}

func (s CheckIfIpAddressesBelongToTheCdnProPlatformResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s CheckIfIpAddressesBelongToTheCdnProPlatformResponseHeader) GoString() string {
  return s.String()
}

type CheckIfIpAddressesBelongToTheCdnProPlatformResponse struct {
  // {"en" : "Describes an IP address.", "zh_CN": "IP地址列表。"}
  IpDetails []*CheckIfIpAddressesBelongToTheCdnProPlatformResponseIpDetails `json:"ipDetails,omitempty" xml:"ipDetails,omitempty" require:"true" type:"Repeated"`
}

func (s CheckIfIpAddressesBelongToTheCdnProPlatformResponse) String() string {
  return tea.Prettify(s)
}

func (s CheckIfIpAddressesBelongToTheCdnProPlatformResponse) GoString() string {
  return s.String()
}

func (s *CheckIfIpAddressesBelongToTheCdnProPlatformResponse) SetIpDetails(v []*CheckIfIpAddressesBelongToTheCdnProPlatformResponseIpDetails) *CheckIfIpAddressesBelongToTheCdnProPlatformResponse {
  s.IpDetails = v
  return s
}

type CheckIfIpAddressesBelongToTheCdnProPlatformResponseIpDetails struct     {
  // {"en" : "An IPv4 or IPv6 address.", "zh_CN": "IPv4或IPv6地址。"}
  Ip *string `json:"ip,omitempty" xml:"ip,omitempty"`
  // {"en" : "True if the IP address belongs to the CDN Pro platform.", "zh_CN": "IP地址是否来自CDN Pro。"}
  IsCdnProIp *bool `json:"isCdnProIp,omitempty" xml:"isCdnProIp,omitempty"`
  // {"en" : "Enum: IPv4,IPv6 
  // Indicates the protocol of the IP address.", "zh_CN": "取值范围: IPv4,IPv6 
  // IP协议。"}
  Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
}

func (s CheckIfIpAddressesBelongToTheCdnProPlatformResponseIpDetails) String() string {
  return tea.Prettify(s)
}

func (s CheckIfIpAddressesBelongToTheCdnProPlatformResponseIpDetails) GoString() string {
  return s.String()
}

func (s *CheckIfIpAddressesBelongToTheCdnProPlatformResponseIpDetails) SetIp(v string) *CheckIfIpAddressesBelongToTheCdnProPlatformResponseIpDetails {
  s.Ip = &v
  return s
}

func (s *CheckIfIpAddressesBelongToTheCdnProPlatformResponseIpDetails) SetIsCdnProIp(v bool) *CheckIfIpAddressesBelongToTheCdnProPlatformResponseIpDetails {
  s.IsCdnProIp = &v
  return s
}

func (s *CheckIfIpAddressesBelongToTheCdnProPlatformResponseIpDetails) SetProtocol(v string) *CheckIfIpAddressesBelongToTheCdnProPlatformResponseIpDetails {
  s.Protocol = &v
  return s
}




