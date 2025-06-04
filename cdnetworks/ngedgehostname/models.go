package ngedgehostname

// This file is auto-generated, don't edit it. Thanks.
import (
  "github.com/alibabacloud-go/tea/tea"
)

type GetAListOfEdgeHostnamesPaths struct {
}

func (s GetAListOfEdgeHostnamesPaths) String() string {
  return tea.Prettify(s)
}

func (s GetAListOfEdgeHostnamesPaths) GoString() string {
  return s.String()
}

type GetAListOfEdgeHostnamesParameters struct {
  // {"en" : "The value is matched against the 'description' and 'edgeHostname' fields of each edge hostname.", "zh_CN": "通过搜索关键字匹配调度域名的前缀和描述进行过滤。"}
  Search *string `json:"search,omitempty" xml:"search,omitempty"`
  // {"en" : "Default: 0 Range: >= 0 
  // Indicates the first edge hostname to return. The first has an offset of 0.", "zh_CN": "默认值: 0 取值范围: >= 0 
  // 查询起始位置。"}
  Offset *int `json:"offset,omitempty" xml:"offset,omitempty"`
  // {"en" : "Default: 100 Range: <= 200 
  // Maximum number of edge hostnames to return.", "zh_CN": "默认值: 100 取值范围: <= 200 
  // 每次查询的最大条数。"}
  Limit *int `json:"limit,omitempty" xml:"limit,omitempty"`
  // {"en" : "Enum: creationTime,lastUpdateTime 
  // Default: lastUpdateTime 
  // Order the list of edge hostnames.", "zh_CN": "取值范围: creationTime,lastUpdateTime 
  // 默认值: lastUpdateTime 
  // 返回结果的排序依据。支持按创建时间或最后一次更新时间进行排序。"}
  SortBy *string `json:"sortBy,omitempty" xml:"sortBy,omitempty"`
  // {"en" : "Range: >= 0 characters 
  // Enum: asc,desc 
  // Default: desc 
  // Indicates the order of edge hostnames returned. Use 'asc' for ascending order or 'desc' for descending order.", "zh_CN": "取值范围: >= 0 字符 
  // 取值范围: asc,desc 
  // 默认值: desc 
  // 返回结果的排序顺序。'asc'表示升序，'desc'表示降序。"}
  SortOrder *string `json:"sortOrder,omitempty" xml:"sortOrder,omitempty"`
  // {"en" : "Specify true to obtain a list of edge hostnames flagged for use with hostnames that have a Beian license, allowing them to be served by servers in mainland China. Specify false to retrieve only those edge hostnames without the Beian flag. By default, all edge hostnames are returned.", "zh_CN": "根据是否备案查询调度域名。"}
  HasBeian *bool `json:"hasBeian,omitempty" xml:"hasBeian,omitempty"`
}

func (s GetAListOfEdgeHostnamesParameters) String() string {
  return tea.Prettify(s)
}

func (s GetAListOfEdgeHostnamesParameters) GoString() string {
  return s.String()
}

func (s *GetAListOfEdgeHostnamesParameters) SetSearch(v string) *GetAListOfEdgeHostnamesParameters {
  s.Search = &v
  return s
}

func (s *GetAListOfEdgeHostnamesParameters) SetOffset(v int) *GetAListOfEdgeHostnamesParameters {
  s.Offset = &v
  return s
}

func (s *GetAListOfEdgeHostnamesParameters) SetLimit(v int) *GetAListOfEdgeHostnamesParameters {
  s.Limit = &v
  return s
}

func (s *GetAListOfEdgeHostnamesParameters) SetSortBy(v string) *GetAListOfEdgeHostnamesParameters {
  s.SortBy = &v
  return s
}

func (s *GetAListOfEdgeHostnamesParameters) SetSortOrder(v string) *GetAListOfEdgeHostnamesParameters {
  s.SortOrder = &v
  return s
}

func (s *GetAListOfEdgeHostnamesParameters) SetHasBeian(v bool) *GetAListOfEdgeHostnamesParameters {
  s.HasBeian = &v
  return s
}

type GetAListOfEdgeHostnamesRequestHeader struct {
}

func (s GetAListOfEdgeHostnamesRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetAListOfEdgeHostnamesRequestHeader) GoString() string {
  return s.String()
}

type GetAListOfEdgeHostnamesRequest struct {
}

func (s GetAListOfEdgeHostnamesRequest) String() string {
  return tea.Prettify(s)
}

func (s GetAListOfEdgeHostnamesRequest) GoString() string {
  return s.String()
}

type GetAListOfEdgeHostnamesResponseHeader struct {
}

func (s GetAListOfEdgeHostnamesResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetAListOfEdgeHostnamesResponseHeader) GoString() string {
  return s.String()
}

type GetAListOfEdgeHostnamesResponse struct {
  // {"en" : "List of edge hostnames.", "zh_CN": "调度域名列表。"}
  EdgeHostnames []*GetAListOfEdgeHostnamesResponseEdgeHostnames `json:"edgeHostnames,omitempty" xml:"edgeHostnames,omitempty" require:"true" type:"Repeated"`
  // {"en" : "Range: >= 0 
  // Total number of edge hostnames in the account. The actual number of edge hostnames returned in the edgehostnames field may be smaller if query parameters are specified.", "zh_CN": "取值范围: >= 0 
  // 调度域名的总数。"}
  Count *int `json:"count,omitempty" xml:"count,omitempty" require:"true"`
}

func (s GetAListOfEdgeHostnamesResponse) String() string {
  return tea.Prettify(s)
}

func (s GetAListOfEdgeHostnamesResponse) GoString() string {
  return s.String()
}

func (s *GetAListOfEdgeHostnamesResponse) SetEdgeHostnames(v []*GetAListOfEdgeHostnamesResponseEdgeHostnames) *GetAListOfEdgeHostnamesResponse {
  s.EdgeHostnames = v
  return s
}

func (s *GetAListOfEdgeHostnamesResponse) SetCount(v int) *GetAListOfEdgeHostnamesResponse {
  s.Count = &v
  return s
}

type GetAListOfEdgeHostnamesResponseEdgeHostnames struct     {
  // {"en" : "An edge hostname.", "zh_CN": "调度域名。"}
  EdgeHostname *string `json:"edgeHostname,omitempty" xml:"edgeHostname,omitempty"`
  // {"en" : "Description of the edge hostname.", "zh_CN": "调度域名的描述。"}
  Description *string `json:"description,omitempty" xml:"description,omitempty"`
  // {"en" : "RFC3339 format string indicating when the edge hostname was last updated.", "zh_CN": "RFC 3339格式的日期，表示调度域名最后一次更新的时间。"}
  LastUpdateTime *string `json:"lastUpdateTime,omitempty" xml:"lastUpdateTime,omitempty"`
  // {"en" : "RFC3339 format string indicating when the edge hostname was created.", "zh_CN": "RFC 3339格式的日期，表示调度域名的创建时间。"}
  CreationTime *string `json:"creationTime,omitempty" xml:"creationTime,omitempty"`
  // {"en" : "true indicates that domains being accelerated have a Beian license allowing content to be served by servers in mainland China.", "zh_CN": "是否备案的标记。当值为true时，将使用包括中国大陆的服务器提供内容加速服务。"}
  HasBeian *bool `json:"hasBeian,omitempty" xml:"hasBeian,omitempty"`
  // {"en" : "Enum: standard,premium,deluxe,ultra 
  // Indicates the highest server group that can be used to serve this edge hostname. The server groups are specified when you create a client zone rule for a new edge hostname. ", "zh_CN": "取值范围: standard,premium,deluxe,ultra 
  // 指该调度域名可使用的最高等级的节点组。当您在'创建调度域名'接口中定义访客分区规则时，可指定节点组。 "}
  HighestServerGroup *string `json:"highestServerGroup,omitempty" xml:"highestServerGroup,omitempty"`
}

func (s GetAListOfEdgeHostnamesResponseEdgeHostnames) String() string {
  return tea.Prettify(s)
}

func (s GetAListOfEdgeHostnamesResponseEdgeHostnames) GoString() string {
  return s.String()
}

func (s *GetAListOfEdgeHostnamesResponseEdgeHostnames) SetEdgeHostname(v string) *GetAListOfEdgeHostnamesResponseEdgeHostnames {
  s.EdgeHostname = &v
  return s
}

func (s *GetAListOfEdgeHostnamesResponseEdgeHostnames) SetDescription(v string) *GetAListOfEdgeHostnamesResponseEdgeHostnames {
  s.Description = &v
  return s
}

func (s *GetAListOfEdgeHostnamesResponseEdgeHostnames) SetLastUpdateTime(v string) *GetAListOfEdgeHostnamesResponseEdgeHostnames {
  s.LastUpdateTime = &v
  return s
}

func (s *GetAListOfEdgeHostnamesResponseEdgeHostnames) SetCreationTime(v string) *GetAListOfEdgeHostnamesResponseEdgeHostnames {
  s.CreationTime = &v
  return s
}

func (s *GetAListOfEdgeHostnamesResponseEdgeHostnames) SetHasBeian(v bool) *GetAListOfEdgeHostnamesResponseEdgeHostnames {
  s.HasBeian = &v
  return s
}

func (s *GetAListOfEdgeHostnamesResponseEdgeHostnames) SetHighestServerGroup(v string) *GetAListOfEdgeHostnamesResponseEdgeHostnames {
  s.HighestServerGroup = &v
  return s
}




type GetClientRegionsPaths struct {
}

func (s GetClientRegionsPaths) String() string {
  return tea.Prettify(s)
}

func (s GetClientRegionsPaths) GoString() string {
  return s.String()
}

type GetClientRegionsParameters struct {
  // {"en" : "Return regions whose code or name contains the specified value.
  // ", "zh_CN": "通过搜索关键字匹配访客区域的代码或名称进行过滤。"}
  Search *string `json:"search,omitempty" xml:"search,omitempty"`
  // {"en" : "Enum: asc,desc 
  // Default: asc 
  // Order of client regions returned. Use 'asc' for ascending order or 'desc' for descending order.", "zh_CN": "取值范围: asc,desc 
  // 默认值: asc 
  // 返回结果的排序顺序。使用'asc'表示升序，'desc'表示降序。"}
  SortOrder *string `json:"sortOrder,omitempty" xml:"sortOrder,omitempty"`
  // {"en" : "Limit the returned regions to those where an ISP operates. The value should be one of the ISP codes returned by the ISPs API.", "zh_CN": "仅返回某个运营商有提供服务的访客区域。运营商是指'查询支持的ISP运营商列表'接口所支持的运营商。"}
  IspCode *string `json:"ispCode,omitempty" xml:"ispCode,omitempty"`
}

func (s GetClientRegionsParameters) String() string {
  return tea.Prettify(s)
}

func (s GetClientRegionsParameters) GoString() string {
  return s.String()
}

func (s *GetClientRegionsParameters) SetSearch(v string) *GetClientRegionsParameters {
  s.Search = &v
  return s
}

func (s *GetClientRegionsParameters) SetSortOrder(v string) *GetClientRegionsParameters {
  s.SortOrder = &v
  return s
}

func (s *GetClientRegionsParameters) SetIspCode(v string) *GetClientRegionsParameters {
  s.IspCode = &v
  return s
}

type GetClientRegionsRequestHeader struct {
}

func (s GetClientRegionsRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetClientRegionsRequestHeader) GoString() string {
  return s.String()
}

type GetClientRegionsRequest struct {
}

func (s GetClientRegionsRequest) String() string {
  return tea.Prettify(s)
}

func (s GetClientRegionsRequest) GoString() string {
  return s.String()
}

type GetClientRegionsResponseHeader struct {
}

func (s GetClientRegionsResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetClientRegionsResponseHeader) GoString() string {
  return s.String()
}

type GetClientRegionsResponse struct {
  // {"en" : "Regions list.", "zh_CN": "访客区域的代码和名称的集合。"}
  Data []*GetClientRegionsResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s GetClientRegionsResponse) String() string {
  return tea.Prettify(s)
}

func (s GetClientRegionsResponse) GoString() string {
  return s.String()
}

func (s *GetClientRegionsResponse) SetData(v []*GetClientRegionsResponseData) *GetClientRegionsResponse {
  s.Data = v
  return s
}

type GetClientRegionsResponseData struct     {
  // {"en" : "Regions code.", "zh_CN": "访客区域的代码。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en" : "Regions name.", "zh_CN": "访客区域的名称。"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
}

func (s GetClientRegionsResponseData) String() string {
  return tea.Prettify(s)
}

func (s GetClientRegionsResponseData) GoString() string {
  return s.String()
}

func (s *GetClientRegionsResponseData) SetCode(v string) *GetClientRegionsResponseData {
  s.Code = &v
  return s
}

func (s *GetClientRegionsResponseData) SetName(v string) *GetClientRegionsResponseData {
  s.Name = &v
  return s
}




type GetAListOfIspsPaths struct {
}

func (s GetAListOfIspsPaths) String() string {
  return tea.Prettify(s)
}

func (s GetAListOfIspsPaths) GoString() string {
  return s.String()
}

type GetAListOfIspsParameters struct {
  // {"en" : "Return ISPs whose code or name contains this value.", "zh_CN": "通过搜索关键字匹配运营商的代码或名称进行过滤。"}
  Search *string `json:"search,omitempty" xml:"search,omitempty"`
  // {"en" : "Enum: asc,desc 
  // Default: asc 
  // Order of ISPs returned. Use 'asc' for ascending order or 'desc' for descending order.", "zh_CN": "取值范围: asc,desc 
  // 默认值: asc 
  // 返回结果的排序顺序。使用'asc'表示升序，'desc'表示降序。"}
  SortOrder *string `json:"sortOrder,omitempty" xml:"sortOrder,omitempty"`
  // {"en" : "Limit returned ISPs to those operating in a particular region. The value of regionCode should be one of the codes representing a region returned by the client regions API.", "zh_CN": "仅返回在某个访客区域有提供服务的运营商。访客区域是指'查询支持的区域列表'接口中支持的访客区域。"}
  RegionCode *string `json:"regionCode,omitempty" xml:"regionCode,omitempty"`
}

func (s GetAListOfIspsParameters) String() string {
  return tea.Prettify(s)
}

func (s GetAListOfIspsParameters) GoString() string {
  return s.String()
}

func (s *GetAListOfIspsParameters) SetSearch(v string) *GetAListOfIspsParameters {
  s.Search = &v
  return s
}

func (s *GetAListOfIspsParameters) SetSortOrder(v string) *GetAListOfIspsParameters {
  s.SortOrder = &v
  return s
}

func (s *GetAListOfIspsParameters) SetRegionCode(v string) *GetAListOfIspsParameters {
  s.RegionCode = &v
  return s
}

type GetAListOfIspsRequestHeader struct {
}

func (s GetAListOfIspsRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetAListOfIspsRequestHeader) GoString() string {
  return s.String()
}

type GetAListOfIspsRequest struct {
}

func (s GetAListOfIspsRequest) String() string {
  return tea.Prettify(s)
}

func (s GetAListOfIspsRequest) GoString() string {
  return s.String()
}

type GetAListOfIspsResponseHeader struct {
}

func (s GetAListOfIspsResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetAListOfIspsResponseHeader) GoString() string {
  return s.String()
}

type GetAListOfIspsResponse struct {
  // {"en" : "ISPs list.", "zh_CN": "运营商的代码和名称的集合。"}
  Data []*GetAListOfIspsResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s GetAListOfIspsResponse) String() string {
  return tea.Prettify(s)
}

func (s GetAListOfIspsResponse) GoString() string {
  return s.String()
}

func (s *GetAListOfIspsResponse) SetData(v []*GetAListOfIspsResponseData) *GetAListOfIspsResponse {
  s.Data = v
  return s
}

type GetAListOfIspsResponseData struct     {
  // {"en" : "ISPs code.", "zh_CN": "运营商的代码。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en" : "ISPs name.", "zh_CN": "运营商的名称。"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
}

func (s GetAListOfIspsResponseData) String() string {
  return tea.Prettify(s)
}

func (s GetAListOfIspsResponseData) GoString() string {
  return s.String()
}

func (s *GetAListOfIspsResponseData) SetCode(v string) *GetAListOfIspsResponseData {
  s.Code = &v
  return s
}

func (s *GetAListOfIspsResponseData) SetName(v string) *GetAListOfIspsResponseData {
  s.Name = &v
  return s
}




type UpdateAnEdgeHostnamePartPaths struct {
  // {"en" : "an edge hostname", "zh_CN": "调度域名。"}
  EdgeHostname *string `json:"edgeHostname,omitempty" xml:"edgeHostname,omitempty" require:"true"`
}

func (s UpdateAnEdgeHostnamePartPaths) String() string {
  return tea.Prettify(s)
}

func (s UpdateAnEdgeHostnamePartPaths) GoString() string {
  return s.String()
}

func (s *UpdateAnEdgeHostnamePartPaths) SetEdgeHostname(v string) *UpdateAnEdgeHostnamePartPaths {
  s.EdgeHostname = &v
  return s
}

type UpdateAnEdgeHostnamePartParameters struct {
}

func (s UpdateAnEdgeHostnamePartParameters) String() string {
  return tea.Prettify(s)
}

func (s UpdateAnEdgeHostnamePartParameters) GoString() string {
  return s.String()
}

type UpdateAnEdgeHostnamePartRequestHeader struct {
}

func (s UpdateAnEdgeHostnamePartRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s UpdateAnEdgeHostnamePartRequestHeader) GoString() string {
  return s.String()
}

type UpdateAnEdgeHostnamePartRequest struct {
  // {"en" : "A description of the edge hostname.", "zh_CN": "调度域名的描述。"}
  Description *string `json:"description,omitempty" xml:"description,omitempty"`
  // {"en" : "Whether PCI compliance is required.  true means content will only be served by PCI certified PoPs.", "zh_CN": "表示流量调度是否需要遵循PCI规范。当值为true时，表示只能使用已通过PCI认证的节点提供内容分发服务。"}
  PciRequired *bool `json:"pciRequired,omitempty" xml:"pciRequired,omitempty"`
  // {"en" : "If set to 'true', clients from European Economic Area (EEA) countries will only be served by IP addresses in EEA countries.", "zh_CN": "表示流量调度是否需要遵循GDPR的规定。当值为'true'时，对于来自欧洲经济区(EEA)国家的请求，将仅使用归属EEA国家的IP地址提供服务。"}
  GrprCompliant *bool `json:"grprCompliant,omitempty" xml:"grprCompliant,omitempty"`
  // {"en" : "Specify rules to control how requests from client zones are handled. There must be a rule that covers all regions and all ISPs.", "zh_CN": "自定义规则来控制如何处理不同访客分区的请求。您必须至少创建一条覆盖所有区域和所有运营商的规则。"}
  ClientZones []*UpdateAnEdgeHostnamePartRequestClientZones `json:"clientZones,omitempty" xml:"clientZones,omitempty" type:"Repeated"`
  // {"en" : "An estimate of the bandwidth required to serve content using this edge hostname. Units of measurement should be in Tbps, Gbps, Mbps, or kbps. Example: 100Gbps", "zh_CN": "通过该调度域名进行CDN加速预计需要的带宽。单位应为Tbps、Gbps、Mbps或kbps。示例：100 Gbps。"}
  EstimatedBandwidth *string `json:"estimatedBandwidth,omitempty" xml:"estimatedBandwidth,omitempty"`
}

func (s UpdateAnEdgeHostnamePartRequest) String() string {
  return tea.Prettify(s)
}

func (s UpdateAnEdgeHostnamePartRequest) GoString() string {
  return s.String()
}

func (s *UpdateAnEdgeHostnamePartRequest) SetDescription(v string) *UpdateAnEdgeHostnamePartRequest {
  s.Description = &v
  return s
}

func (s *UpdateAnEdgeHostnamePartRequest) SetPciRequired(v bool) *UpdateAnEdgeHostnamePartRequest {
  s.PciRequired = &v
  return s
}

func (s *UpdateAnEdgeHostnamePartRequest) SetGrprCompliant(v bool) *UpdateAnEdgeHostnamePartRequest {
  s.GrprCompliant = &v
  return s
}

func (s *UpdateAnEdgeHostnamePartRequest) SetClientZones(v []*UpdateAnEdgeHostnamePartRequestClientZones) *UpdateAnEdgeHostnamePartRequest {
  s.ClientZones = v
  return s
}

func (s *UpdateAnEdgeHostnamePartRequest) SetEstimatedBandwidth(v string) *UpdateAnEdgeHostnamePartRequest {
  s.EstimatedBandwidth = &v
  return s
}

type UpdateAnEdgeHostnamePartRequestClientZones struct     {
  // {"en" : "This field indicates the region in which the rule applies. Refer to our API to get client regions to get valid region codes. For example, if you wish to create a rule that covers all of Europe, simply specify 'eu' as the region. You can indicate specific countries. For example, 'na.us' represents the 'United States of America', and 'eu.fr' represents 'France'.  
  // 
  // A special client region 'all' can be used to specify that the rule applies to the entire world. If overlapping regions are specified, the more specific one takes precedence. For example, if you specify 'as' in one rule and 'as.cn' in another, a request from China will follow the rule for 'as.cn'.", "zh_CN": "该规则适用的区域。可调用'查询支持的区域列表'接口来查看区域信息。例如，如果您希望创建规则覆盖整个欧洲，则指定'eu'为区域。 您可以指定具体的国家。例如，'na.us'代表'美国'，而'eu.fr'代表'法国'。
  // 
  // 'all'是一个特殊的区域，可用于指定适用于全球的规则。如果不同规则指定的区域存在重叠，则以更细粒度的区域优先。例如，如果您在一条规则中指定'as'，在另一条规则中指定'as.cn'，则来自中国的请求将优先匹配'as.cn'的规则。"}
  Region *string `json:"region,omitempty" xml:"region,omitempty"`
  // {"en" : "Default: 100 Range: [ 0 .. 100 ] 
  // When multiple actions are specified for a client zone, the weight is used to adjust the probability of a client zone rule applying to a client request.
  // Consider two rules that apply to 'as.cn': {'region':'as.cn', 'isp':'all', 'action':{'type':'deliver', 'by':['standard', 'premium', 'deluxe']}},
  //     {'region':'as.cn', 'isp':'all', 'action':{'type':'redirect', 'to':['alternate.cname.com']}, 'weight':10}
  // 
  //  The probability of 'deliver' will be 100/(100+10) or 0.909 while the probability of 'redirect' will be 10/(100+10) or 0.091.
  //  
  // ", "zh_CN": "默认值: 100 取值范围: [ 0 .. 100 ] 
  // 可以为同个访客分区指定多条规则。通过在规则中指定weight字段，可控制规则匹配的权重。
  // 以'as.cn'区域的2条规则为例: {'region':'as.cn', 'isp':'all', 'action':{'type':'deliver', 'by':['standard', 'premium', 'deluxe']}},
  //     {'region':'as.cn', 'isp':'all', 'action':{'type':'redirect', 'to':['alternate.cname.com']}, 'weight':10}
  // 
  // 按照以上规则，客户端请求匹配规则1进行'分发'的比例为100/(100+10)，即0.909，匹配规则2进行'跳转'的比例为10/(100+10)，即0.091。
  //  
  // "}
  Weight *int `json:"weight,omitempty" xml:"weight,omitempty"`
  // {"en" : "This object describes the action to take for requests matching the rule.", "zh_CN": "当规则匹配时执行的动作。"}
  Action *UpdateAnEdgeHostnamePartRequestClientZonesAction `json:"action,omitempty" xml:"action,omitempty" type:"Struct"`
  // {"en" : "Specify the code representing an ISP (Internet Service Provider) if the rule only applies to requests from a particular ISP. Call our API to get a list of supported ISPs. Specify 'all' to indicate all ISPs rather than a particular one. Specify a comma-separated list of up to 10 ISP codes if you want your rule to apply to more than one ISP.", "zh_CN": "该规则适用的运营商。可调用我们的'查询支持的ISP运营商列表'接口查看运营商信息。指定'all'表示所有运营商。如果希望该规则应用于多个运营商，则可指定多个运营商，用逗号分隔，但最多只能包含10个运营商。"}
  Isp *string `json:"isp,omitempty" xml:"isp,omitempty"`
}

func (s UpdateAnEdgeHostnamePartRequestClientZones) String() string {
  return tea.Prettify(s)
}

func (s UpdateAnEdgeHostnamePartRequestClientZones) GoString() string {
  return s.String()
}

func (s *UpdateAnEdgeHostnamePartRequestClientZones) SetRegion(v string) *UpdateAnEdgeHostnamePartRequestClientZones {
  s.Region = &v
  return s
}

func (s *UpdateAnEdgeHostnamePartRequestClientZones) SetWeight(v int) *UpdateAnEdgeHostnamePartRequestClientZones {
  s.Weight = &v
  return s
}

func (s *UpdateAnEdgeHostnamePartRequestClientZones) SetAction(v *UpdateAnEdgeHostnamePartRequestClientZonesAction) *UpdateAnEdgeHostnamePartRequestClientZones {
  s.Action = v
  return s
}

func (s *UpdateAnEdgeHostnamePartRequestClientZones) SetIsp(v string) *UpdateAnEdgeHostnamePartRequestClientZones {
  s.Isp = &v
  return s
}

type UpdateAnEdgeHostnamePartRequestClientZonesAction struct {
  // {"en" : "Enum: deliver,redirect,reject 
  // Defines the action to take for requests to the zone. Options are to deliver using one or more server groups, to reject the request altogether, or to redirect to another domain. If 'reject' is specified, the client request will be redirected to a server that always responds with HTTP response code 403 representing 'Forbidden'. Up to one 'reject' action is allowed for each client zone.
  // ", "zh_CN": "取值范围: deliver,redirect,reject 
  // 当规则匹配时，对客户端请求执行的动作的类型。包括分发，拒绝和跳转3个类型。如果指定了'拒绝'，则客户端请求将被调度到一台服务器，该服务器总是响应403状态码，表示'Forbidden'拒绝访问。每个访客分区最多只允许一个'拒绝'动作。"}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
  // {"en" : "This field is used if the action is of type 'deliver'. Specify one or more of the server groups (standard, premium, deluxe,  ultra, nearChina, ChinaStandard, ChinaPremium) to control the servers used to deliver content. If unspecified, 'standard' is used. <br/><br/>'nearChina' is a special server group that can be enabled for you if your domains lack ICP Beian, but you want optimal performance serving Chinese visitors. A client zone rule using nearChina cannot be configured to simultaneously deliver to other server groups, though you may create separate client zone rules to use other server groups.<br/><br/>If you have an ICP Beian license and want your content to be delivered by servers within China, you can opt to use 'ChinaStandard' and 'ChinaPremium.' Please contact our support team if you require nearChina, ChinaStandard, or ChinaPremium.<br/><br/>If an edge hostname is not initially configured to use ChinaStandard or ChinaPremium, you will not be able to modify it later to use these two server groups. Instead, you will need to create a new edge hostname with client zone rules delivering to those server groups.", "zh_CN": "如果动作类型为'分发'，则使用此字段指定一个或多个节点组(standard, premium, deluxe,  ultra, nearChina, ChinaStandard, ChinaPremium) 来选择提供内容分发服务的缓存节点。如果未指定，则使用'standard'。<br/><br/>'nearChina'  是一个特殊的节点组。如果您需要使用nearChina节点组，请联系我们的技术支持开通。不能在同一条规则中同时指定nearChina节点组和其他节点组。如果要使用其他节点组，需要创建单独的访客分区规则。<br/><br/>如果您的加速域名有ICP备案，希望由中国大陆的服务器提供内容分发服务，您可以选择使用'ChinaStandard'和'ChinaPremium'节点组。
  // <br/><br/>如果调度域名创建时没有指定使用ChinaStandard或ChinaPremium节点组，则无法通过更新调度域名来使用这两个节点组。您需要创建一个新的调度域名，在新的调度域名中指定使用ChinaStandard或ChinaPremium，才能使用这2个节点组。"}
  By []*string `json:"by,omitempty" xml:"by,omitempty" type:"Repeated"`
  // {"en" : "This field is used if the action is of type 'redirect'. It indicates the hostname or IP address to redirect to. This is typically an origin server or another CDN provider.", "zh_CN": "如果动作类型为'跳转'，则通过该字段指定跳转的目标域名或IP地址。'跳转'目标通常是源站服务器或其它CDN厂商。"}
  To []*string `json:"to,omitempty" xml:"to,omitempty" type:"Repeated"`
  // {"en" : "Default: True 
  // Indicates whether an IPv6 address can be used. This value is used only if the client zone rule's action type is 'deliver'.", "zh_CN": "默认值: True 
  // 指定是否允许使用IPv6地址进行内容分发。仅当动作类型为'分发'时该值才有效。"}
  EnableIPv6 *bool `json:"enableIPv6,omitempty" xml:"enableIPv6,omitempty"`
}

func (s UpdateAnEdgeHostnamePartRequestClientZonesAction) String() string {
  return tea.Prettify(s)
}

func (s UpdateAnEdgeHostnamePartRequestClientZonesAction) GoString() string {
  return s.String()
}

func (s *UpdateAnEdgeHostnamePartRequestClientZonesAction) SetType(v string) *UpdateAnEdgeHostnamePartRequestClientZonesAction {
  s.Type = &v
  return s
}

func (s *UpdateAnEdgeHostnamePartRequestClientZonesAction) SetBy(v []*string) *UpdateAnEdgeHostnamePartRequestClientZonesAction {
  s.By = v
  return s
}

func (s *UpdateAnEdgeHostnamePartRequestClientZonesAction) SetTo(v []*string) *UpdateAnEdgeHostnamePartRequestClientZonesAction {
  s.To = v
  return s
}

func (s *UpdateAnEdgeHostnamePartRequestClientZonesAction) SetEnableIPv6(v bool) *UpdateAnEdgeHostnamePartRequestClientZonesAction {
  s.EnableIPv6 = &v
  return s
}

type UpdateAnEdgeHostnamePartResponseHeader struct {
}

func (s UpdateAnEdgeHostnamePartResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s UpdateAnEdgeHostnamePartResponseHeader) GoString() string {
  return s.String()
}

type UpdateAnEdgeHostnamePartResponse struct {
}

func (s UpdateAnEdgeHostnamePartResponse) String() string {
  return tea.Prettify(s)
}

func (s UpdateAnEdgeHostnamePartResponse) GoString() string {
  return s.String()
}




type UpdateAnEdgeHostnameAllPaths struct {
  // {"en" : "an edge hostname", "zh_CN": "调度域名。"}
  EdgeHostname *string `json:"edgeHostname,omitempty" xml:"edgeHostname,omitempty" require:"true"`
}

func (s UpdateAnEdgeHostnameAllPaths) String() string {
  return tea.Prettify(s)
}

func (s UpdateAnEdgeHostnameAllPaths) GoString() string {
  return s.String()
}

func (s *UpdateAnEdgeHostnameAllPaths) SetEdgeHostname(v string) *UpdateAnEdgeHostnameAllPaths {
  s.EdgeHostname = &v
  return s
}

type UpdateAnEdgeHostnameAllParameters struct {
}

func (s UpdateAnEdgeHostnameAllParameters) String() string {
  return tea.Prettify(s)
}

func (s UpdateAnEdgeHostnameAllParameters) GoString() string {
  return s.String()
}

type UpdateAnEdgeHostnameAllRequestHeader struct {
}

func (s UpdateAnEdgeHostnameAllRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s UpdateAnEdgeHostnameAllRequestHeader) GoString() string {
  return s.String()
}

type UpdateAnEdgeHostnameAllRequest struct {
  // {"en" : "The edge hostname. It cannot be modified.", "zh_CN": "调度域名。该字段不可修改。"}
  EdgeHostname *string `json:"edgeHostname,omitempty" xml:"edgeHostname,omitempty" require:"true"`
  // {"en" : "A description of the edge hostname.", "zh_CN": "调度域名的描述。"}
  Description *string `json:"description,omitempty" xml:"description,omitempty"`
  // {"en" : "Whether PCI compliance is required.  true means content will only be served by PCI certified PoPs.", "zh_CN": "表示流量调度是否需要遵循PCI规范。当值为true时，表示只能使用已通过PCI认证的节点提供内容分发服务。"}
  PciRequired *bool `json:"pciRequired,omitempty" xml:"pciRequired,omitempty"`
  // {"en" : "If set to 'true', clients from European Economic Area (EEA) countries will only be served by IP addresses in EEA countries.", "zh_CN": "表示流量调度是否需要遵循GDPR的规定。当值为'true'时，对于来自欧洲经济区(EEA)国家的请求，将仅使用归属EEA国家的IP地址提供服务。"}
  GdprCompliant *bool `json:"gdprCompliant,omitempty" xml:"gdprCompliant,omitempty"`
  // {"en" : "Specify rules to control how requests from client zones are handled. There must be a rule that covers all regions and all ISPs.", "zh_CN": "自定义规则来控制如何处理不同访客分区的请求。您必须至少创建一条覆盖所有区域和所有运营商的规则。"}
  ClientZones []*UpdateAnEdgeHostnameAllRequestClientZones `json:"clientZones,omitempty" xml:"clientZones,omitempty" type:"Repeated"`
  // {"en" : "An estimate of the bandwidth required to serve content using this edge hostname. Units of measurement should be in Tbps, Gbps, Mbps, or kbps. Example: 100Gbps", "zh_CN": "通过该调度域名进行CDN加速预计需要的带宽。单位应为Tbps、Gbps、Mbps或kbps。示例：100 Gbps。"}
  EstimatedBandwidth *string `json:"estimatedBandwidth,omitempty" xml:"estimatedBandwidth,omitempty"`
}

func (s UpdateAnEdgeHostnameAllRequest) String() string {
  return tea.Prettify(s)
}

func (s UpdateAnEdgeHostnameAllRequest) GoString() string {
  return s.String()
}

func (s *UpdateAnEdgeHostnameAllRequest) SetEdgeHostname(v string) *UpdateAnEdgeHostnameAllRequest {
  s.EdgeHostname = &v
  return s
}

func (s *UpdateAnEdgeHostnameAllRequest) SetDescription(v string) *UpdateAnEdgeHostnameAllRequest {
  s.Description = &v
  return s
}

func (s *UpdateAnEdgeHostnameAllRequest) SetPciRequired(v bool) *UpdateAnEdgeHostnameAllRequest {
  s.PciRequired = &v
  return s
}

func (s *UpdateAnEdgeHostnameAllRequest) SetGdprCompliant(v bool) *UpdateAnEdgeHostnameAllRequest {
  s.GdprCompliant = &v
  return s
}

func (s *UpdateAnEdgeHostnameAllRequest) SetClientZones(v []*UpdateAnEdgeHostnameAllRequestClientZones) *UpdateAnEdgeHostnameAllRequest {
  s.ClientZones = v
  return s
}

func (s *UpdateAnEdgeHostnameAllRequest) SetEstimatedBandwidth(v string) *UpdateAnEdgeHostnameAllRequest {
  s.EstimatedBandwidth = &v
  return s
}

type UpdateAnEdgeHostnameAllRequestClientZones struct     {
  // {"en" : "This field indicates the region in which the rule applies. Refer to our API to get client regions to get valid region codes. For example, if you wish to create a rule that covers all of Europe, simply specify 'eu' as the region. You can indicate specific countries. For example, 'na.us' represents the 'United States of America', and 'eu.fr' represents 'France'.  
  // 
  // A special client region 'all' can be used to specify that the rule applies to the entire world. If overlapping regions are specified, the more specific one takes precedence. For example, if you specify 'as' in one rule and 'as.cn' in another, a request from China will follow the rule for 'as.cn'.", "zh_CN": "该规则适用的区域。可调用'查询支持的区域列表'接口来查看区域信息。例如，如果您希望创建规则覆盖整个欧洲，则指定'eu'为区域。 您可以指定具体的国家。例如，'na.us'代表'美国'，而'eu.fr'代表'法国'。
  // 
  // 'all'是一个特殊的区域，可用于指定适用于全球的规则。如果不同规则指定的区域存在重叠，则以更细粒度的区域优先。例如，如果您在一条规则中指定'as'，在另一条规则中指定'as.cn'，则来自中国的请求将优先匹配'as.cn'的规则。"}
  Region *string `json:"region,omitempty" xml:"region,omitempty"`
  // {"en" : "Default: 100 Range: [ 0 .. 100 ] 
  // When multiple actions are specified for a client zone, the weight is used to adjust the probability of a client zone rule applying to a client request.
  // Consider two rules that apply to 'as.cn': {'region':'as.cn', 'isp':'all', 'action':{'type':'deliver', 'by':['standard', 'premium', 'deluxe']}},
  //     {'region':'as.cn', 'isp':'all', 'action':{'type':'redirect', 'to':['alternate.cname.com']}, 'weight':10}
  // 
  //  The probability of 'deliver' will be 100/(100+10) or 0.909 while the probability of 'redirect' will be 10/(100+10) or 0.091.
  //  
  // ", "zh_CN": "默认值: 100 取值范围: [ 0 .. 100 ] 
  // 可以为同个访客分区指定多条规则。通过在规则中指定weight字段，可控制规则匹配的权重。
  // 以'as.cn'区域的2条规则为例: {'region':'as.cn', 'isp':'all', 'action':{'type':'deliver', 'by':['standard', 'premium', 'deluxe']}},
  //     {'region':'as.cn', 'isp':'all', 'action':{'type':'redirect', 'to':['alternate.cname.com']}, 'weight':10}
  // 
  // 按照以上规则，客户端请求匹配规则1进行'分发'的比例为100/(100+10)，即0.909，匹配规则2进行'跳转'的比例为10/(100+10)，即0.091。
  //  
  // "}
  Weight *int `json:"weight,omitempty" xml:"weight,omitempty"`
  // {"en" : "This object describes the action to take for requests matching the rule.", "zh_CN": "当规则匹配时执行的动作。"}
  Action *UpdateAnEdgeHostnameAllRequestClientZonesAction `json:"action,omitempty" xml:"action,omitempty" type:"Struct"`
  // {"en" : "Specify the code representing an ISP (Internet Service Provider) if the rule only applies to requests from a particular ISP. Call our API to get a list of supported ISPs. Specify 'all' to indicate all ISPs rather than a particular one. Specify a comma-separated list of up to 10 ISP codes if you want your rule to apply to more than one ISP.", "zh_CN": "该规则适用的运营商。可调用我们的'查询支持的ISP运营商列表'接口查看运营商信息。指定'all'表示所有运营商。如果希望该规则应用于多个运营商，则可指定多个运营商，用逗号分隔，但最多只能包含10个运营商。"}
  Isp *string `json:"isp,omitempty" xml:"isp,omitempty"`
}

func (s UpdateAnEdgeHostnameAllRequestClientZones) String() string {
  return tea.Prettify(s)
}

func (s UpdateAnEdgeHostnameAllRequestClientZones) GoString() string {
  return s.String()
}

func (s *UpdateAnEdgeHostnameAllRequestClientZones) SetRegion(v string) *UpdateAnEdgeHostnameAllRequestClientZones {
  s.Region = &v
  return s
}

func (s *UpdateAnEdgeHostnameAllRequestClientZones) SetWeight(v int) *UpdateAnEdgeHostnameAllRequestClientZones {
  s.Weight = &v
  return s
}

func (s *UpdateAnEdgeHostnameAllRequestClientZones) SetAction(v *UpdateAnEdgeHostnameAllRequestClientZonesAction) *UpdateAnEdgeHostnameAllRequestClientZones {
  s.Action = v
  return s
}

func (s *UpdateAnEdgeHostnameAllRequestClientZones) SetIsp(v string) *UpdateAnEdgeHostnameAllRequestClientZones {
  s.Isp = &v
  return s
}

type UpdateAnEdgeHostnameAllRequestClientZonesAction struct {
  // {"en" : "Enum: deliver,redirect,reject 
  // Defines the action to take for requests to the zone. Options are to deliver using one or more server groups, to reject the request altogether, or to redirect to another domain. If 'reject' is specified, the client request will be redirected to a server that always responds with HTTP response code 403 representing 'Forbidden'. Up to one 'reject' action is allowed for each client zone.
  // ", "zh_CN": "取值范围: deliver,redirect,reject 
  // 当规则匹配时，对客户端请求执行的动作的类型。包括分发，拒绝和跳转3个类型。如果指定了'拒绝'，则客户端请求将被调度到一台服务器，该服务器总是响应403状态码，表示'Forbidden'拒绝访问。每个访客分区最多只允许一个'拒绝'动作。"}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
  // {"en" : "This field is used if the action is of type 'deliver'. Specify one or more of the server groups (standard, premium, deluxe,  ultra, nearChina, ChinaStandard, ChinaPremium) to control the servers used to deliver content. If unspecified, 'standard' is used. <br/><br/>'nearChina' is a special server group that can be enabled for you if your domains lack ICP Beian, but you want optimal performance serving Chinese visitors. A client zone rule using nearChina cannot be configured to simultaneously deliver to other server groups, though you may create separate client zone rules to use other server groups.<br/><br/>If you have an ICP Beian license and want your content to be delivered by servers within China, you can opt to use 'ChinaStandard' and 'ChinaPremium.' Please contact our support team if you require nearChina, ChinaStandard, or ChinaPremium.<br/><br/>If an edge hostname is not initially configured to use ChinaStandard or ChinaPremium, you will not be able to modify it later to use these two server groups. Instead, you will need to create a new edge hostname with client zone rules delivering to those server groups.", "zh_CN": "如果动作类型为'分发'，则使用此字段指定一个或多个节点组(standard, premium, deluxe,  ultra, nearChina, ChinaStandard, ChinaPremium) 来选择提供内容分发服务的缓存节点。如果未指定，则使用'standard'。<br/><br/>'nearChina'  是一个特殊的节点组。如果您需要使用nearChina节点组，请联系我们的技术支持开通。不能在同一条规则中同时指定nearChina节点组和其他节点组。如果要使用其他节点组，需要创建单独的访客分区规则。<br/><br/>如果您的加速域名有ICP备案，希望由中国大陆的服务器提供内容分发服务，您可以选择使用'ChinaStandard'和'ChinaPremium'节点组。
  // <br/><br/>如果调度域名创建时没有指定使用ChinaStandard或ChinaPremium节点组，则无法通过更新调度域名来使用这两个节点组。您需要创建一个新的调度域名，在新的调度域名中指定使用ChinaStandard或ChinaPremium，才能使用这2个节点组。"}
  By []*string `json:"by,omitempty" xml:"by,omitempty" type:"Repeated"`
  // {"en" : "This field is used if the action is of type 'redirect'. It indicates the hostname or IP address to redirect to. This is typically an origin server or another CDN provider.", "zh_CN": "如果动作类型为'跳转'，则通过该字段指定跳转的目标域名或IP地址。'跳转'目标通常是源站服务器或其它CDN厂商。"}
  To []*string `json:"to,omitempty" xml:"to,omitempty" type:"Repeated"`
  // {"en" : "Default: True 
  // Indicates whether an IPv6 address can be used. This value is used only if the client zone rule's action type is 'deliver'.", "zh_CN": "默认值: True 
  // 指定是否允许使用IPv6地址进行内容分发。仅当动作类型为'分发'时该值才有效。"}
  EnableIPv6 *bool `json:"enableIPv6,omitempty" xml:"enableIPv6,omitempty"`
}

func (s UpdateAnEdgeHostnameAllRequestClientZonesAction) String() string {
  return tea.Prettify(s)
}

func (s UpdateAnEdgeHostnameAllRequestClientZonesAction) GoString() string {
  return s.String()
}

func (s *UpdateAnEdgeHostnameAllRequestClientZonesAction) SetType(v string) *UpdateAnEdgeHostnameAllRequestClientZonesAction {
  s.Type = &v
  return s
}

func (s *UpdateAnEdgeHostnameAllRequestClientZonesAction) SetBy(v []*string) *UpdateAnEdgeHostnameAllRequestClientZonesAction {
  s.By = v
  return s
}

func (s *UpdateAnEdgeHostnameAllRequestClientZonesAction) SetTo(v []*string) *UpdateAnEdgeHostnameAllRequestClientZonesAction {
  s.To = v
  return s
}

func (s *UpdateAnEdgeHostnameAllRequestClientZonesAction) SetEnableIPv6(v bool) *UpdateAnEdgeHostnameAllRequestClientZonesAction {
  s.EnableIPv6 = &v
  return s
}

type UpdateAnEdgeHostnameAllResponseHeader struct {
}

func (s UpdateAnEdgeHostnameAllResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s UpdateAnEdgeHostnameAllResponseHeader) GoString() string {
  return s.String()
}

type UpdateAnEdgeHostnameAllResponse struct {
}

func (s UpdateAnEdgeHostnameAllResponse) String() string {
  return tea.Prettify(s)
}

func (s UpdateAnEdgeHostnameAllResponse) GoString() string {
  return s.String()
}




type DeleteAnEdgeHostnamePaths struct {
  // {"en" : "an edge hostname", "zh_CN": "调度域名。"}
  EdgeHostname *string `json:"edgeHostname,omitempty" xml:"edgeHostname,omitempty" require:"true"`
}

func (s DeleteAnEdgeHostnamePaths) String() string {
  return tea.Prettify(s)
}

func (s DeleteAnEdgeHostnamePaths) GoString() string {
  return s.String()
}

func (s *DeleteAnEdgeHostnamePaths) SetEdgeHostname(v string) *DeleteAnEdgeHostnamePaths {
  s.EdgeHostname = &v
  return s
}

type DeleteAnEdgeHostnameParameters struct {
}

func (s DeleteAnEdgeHostnameParameters) String() string {
  return tea.Prettify(s)
}

func (s DeleteAnEdgeHostnameParameters) GoString() string {
  return s.String()
}

type DeleteAnEdgeHostnameRequestHeader struct {
  // {"en":"When you request to delete an edgeHostname, we check if the edgeHostname is still active. Deletion of the edgeHostname will be rejected if there are DNS requests to the edgeHostname in the past 24 hours. This is to prevent accidental deletion. If you are sure that the edgeHostname can be deleted, you can bypass this check by passing the <i>Check-Usage: no</i> header.", "zh_CN":"当您删除调度域名时，后台会校验调度域名是否处于活跃状态。如果调度域名在过去24小时有DNS解析请求，则会拒绝删除调度域名，避免误操作。如果您确定要删除调度域名，可以在调用接口时传入请求头<i>Check-Usage: no</i>来跳过校验。"}
  CheckUsage *string `json:"Check-Usage,omitempty" xml:"Check-Usage,omitempty"`
}

func (s DeleteAnEdgeHostnameRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteAnEdgeHostnameRequestHeader) GoString() string {
  return s.String()
}

func (s *DeleteAnEdgeHostnameRequestHeader) SetCheckUsage(v string) *DeleteAnEdgeHostnameRequestHeader {
  s.CheckUsage = &v
  return s
}

type DeleteAnEdgeHostnameRequest struct {
}

func (s DeleteAnEdgeHostnameRequest) String() string {
  return tea.Prettify(s)
}

func (s DeleteAnEdgeHostnameRequest) GoString() string {
  return s.String()
}

type DeleteAnEdgeHostnameResponseHeader struct {
}

func (s DeleteAnEdgeHostnameResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteAnEdgeHostnameResponseHeader) GoString() string {
  return s.String()
}

type DeleteAnEdgeHostnameResponse struct {
}

func (s DeleteAnEdgeHostnameResponse) String() string {
  return tea.Prettify(s)
}

func (s DeleteAnEdgeHostnameResponse) GoString() string {
  return s.String()
}




type CreateAnEdgeHostnamePaths struct {
}

func (s CreateAnEdgeHostnamePaths) String() string {
  return tea.Prettify(s)
}

func (s CreateAnEdgeHostnamePaths) GoString() string {
  return s.String()
}

type CreateAnEdgeHostnameParameters struct {
}

func (s CreateAnEdgeHostnameParameters) String() string {
  return tea.Prettify(s)
}

func (s CreateAnEdgeHostnameParameters) GoString() string {
  return s.String()
}

type CreateAnEdgeHostnameRequestHeader struct {
}

func (s CreateAnEdgeHostnameRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s CreateAnEdgeHostnameRequestHeader) GoString() string {
  return s.String()
}

type CreateAnEdgeHostnameRequest struct {
  // {"en" : "A hostname to use in your DNS server settings. When your property has been deployed to production, you should modify your DNS server so that the CNAME records of the property's hostnames refer to this value. The edge hostname should consist of at least ten characters followed by a permitted suffix such as '.qtlcdn.com' (for example, abcdefg123.qtlcdn.com). Your service quota may include an edgeHostnameZones field. If it does, the suffix used for your edge hostname should be one of the listed values. ", "zh_CN": "当您部署加速项目到生产环境后，您应该修改DNS解析，添加一条CNAME记录，将您的加速域名指向该调度域名。每个调度域名由2部组成，即前缀和后缀（例如， abcdefg123.qtlcdn.com）。前缀可自定义，后缀为既定的DNS zone，如'.qtlcdn.com'。如果您需要使用自定义的DNS zone，请联系我们的技术支持。"}
  EdgeHostname *string `json:"edgeHostname,omitempty" xml:"edgeHostname,omitempty"`
  // {"en" : "A description of the edge hostname.", "zh_CN": "调度域名的描述。"}
  Description *string `json:"description,omitempty" xml:"description,omitempty"`
  // {"en" : "Default: False 
  // Indicates whether PCI compliance is required.  true means content will only be served by PCI certified PoPs.", "zh_CN": "默认值: False 
  // 表示流量调度是否需要遵循PCI规范。当值为true时，表示只能使用已通过PCI认证的节点提供内容分发服务。"}
  PciRequired *bool `json:"pciRequired,omitempty" xml:"pciRequired,omitempty"`
  // {"en" : "Default: False 
  // If set to 'true', clients from European Economic Area (EEA) countries will only be served by IP addresses in EEA countries.", "zh_CN": "默认值: False 
  // 表示流量调度是否需要遵循GDPR的规定。当值为'true'时，对于来自欧洲经济区(EEA)国家的请求，将仅使用归属EEA国家的IP地址提供服务。"}
  GdprCompliant *bool `json:"gdprCompliant,omitempty" xml:"gdprCompliant,omitempty"`
  // {"en" : "Specify rules to control how requests from client zones are handled. There must be a rule that covers all regions and all ISPs.", "zh_CN": "自定义规则来控制如何处理不同访客分区的请求。您必须至少创建一条覆盖所有区域和所有运营商的规则。"}
  ClientZones []*CreateAnEdgeHostnameRequestClientZones `json:"clientZones,omitempty" xml:"clientZones,omitempty" type:"Repeated"`
  // {"en" : "An estimate of the bandwidth required to serve content using this edge hostname. Units of measurement should be in Tbps, Gbps, Mbps, or kbps. Example: 100Gbps", "zh_CN": "通过该调度域名进行CDN加速预计需要的带宽。单位应为Tbps、Gbps、Mbps或kbps。示例：100 Gbps。"}
  EstimatedBandwidth *string `json:"estimatedBandwidth,omitempty" xml:"estimatedBandwidth,omitempty"`
}

func (s CreateAnEdgeHostnameRequest) String() string {
  return tea.Prettify(s)
}

func (s CreateAnEdgeHostnameRequest) GoString() string {
  return s.String()
}

func (s *CreateAnEdgeHostnameRequest) SetEdgeHostname(v string) *CreateAnEdgeHostnameRequest {
  s.EdgeHostname = &v
  return s
}

func (s *CreateAnEdgeHostnameRequest) SetDescription(v string) *CreateAnEdgeHostnameRequest {
  s.Description = &v
  return s
}

func (s *CreateAnEdgeHostnameRequest) SetPciRequired(v bool) *CreateAnEdgeHostnameRequest {
  s.PciRequired = &v
  return s
}

func (s *CreateAnEdgeHostnameRequest) SetGdprCompliant(v bool) *CreateAnEdgeHostnameRequest {
  s.GdprCompliant = &v
  return s
}

func (s *CreateAnEdgeHostnameRequest) SetClientZones(v []*CreateAnEdgeHostnameRequestClientZones) *CreateAnEdgeHostnameRequest {
  s.ClientZones = v
  return s
}

func (s *CreateAnEdgeHostnameRequest) SetEstimatedBandwidth(v string) *CreateAnEdgeHostnameRequest {
  s.EstimatedBandwidth = &v
  return s
}

type CreateAnEdgeHostnameRequestClientZones struct     {
  // {"en" : "This field indicates the region in which the rule applies. Refer to our API to get client regions to get valid region codes. For example, if you wish to create a rule that covers all of Europe, simply specify 'eu' as the region. You can indicate specific countries. For example, 'na.us' represents the 'United States of America', and 'eu.fr' represents 'France'.  
  // 
  // A special client region 'all' can be used to specify that the rule applies to the entire world. If overlapping regions are specified, the more specific one takes precedence. For example, if you specify 'as' in one rule and 'as.cn' in another, a request from China will follow the rule for 'as.cn'.", "zh_CN": "该规则适用的区域。可调用'查询支持的区域列表'接口来查看区域信息。例如，如果您希望创建规则覆盖整个欧洲，则指定'eu'为区域。 您可以指定具体的国家。例如，'na.us'代表'美国'，而'eu.fr'代表'法国'。
  // 
  // 'all'是一个特殊的区域，可用于指定适用于全球的规则。如果不同规则指定的区域存在重叠，则以更细粒度的区域优先。例如，如果您在一条规则中指定'as'，在另一条规则中指定'as.cn'，则来自中国的请求将优先匹配'as.cn'的规则。"}
  Region *string `json:"region,omitempty" xml:"region,omitempty"`
  // {"en" : "Default: 100 Range: [ 0 .. 100 ] 
  // When multiple actions are specified for a client zone, the weight is used to adjust the probability of a client zone rule applying to a client request.
  // Consider two rules that apply to 'as.cn': {'region':'as.cn', 'isp':'all', 'action':{'type':'deliver', 'by':['standard', 'premium', 'deluxe']}},
  //     {'region':'as.cn', 'isp':'all', 'action':{'type':'redirect', 'to':['alternate.cname.com']}, 'weight':10}
  // 
  //  The probability of 'deliver' will be 100/(100+10) or 0.909 while the probability of 'redirect' will be 10/(100+10) or 0.091.
  //  
  // ", "zh_CN": "默认值: 100 取值范围: [ 0 .. 100 ] 
  // 可以为同个访客分区指定多条规则。通过在规则中指定weight字段，可控制规则匹配的权重。
  // 以'as.cn'区域的2条规则为例: {'region':'as.cn', 'isp':'all', 'action':{'type':'deliver', 'by':['standard', 'premium', 'deluxe']}},
  //     {'region':'as.cn', 'isp':'all', 'action':{'type':'redirect', 'to':['alternate.cname.com']}, 'weight':10}
  // 
  // 按照以上规则，客户端请求匹配规则1进行'分发'的比例为100/(100+10)，即0.909，匹配规则2进行'跳转'的比例为10/(100+10)，即0.091。
  //  
  // "}
  Weight *int `json:"weight,omitempty" xml:"weight,omitempty"`
  // {"en" : "This object describes the action to take for requests matching the rule.", "zh_CN": "当规则匹配时执行的动作。"}
  Action *CreateAnEdgeHostnameRequestClientZonesAction `json:"action,omitempty" xml:"action,omitempty" type:"Struct"`
  // {"en" : "Specify the code representing an ISP (Internet Service Provider) if the rule only applies to requests from a particular ISP. Call our API to get a list of supported ISPs. Specify 'all' to indicate all ISPs rather than a particular one. Specify a comma-separated list of up to 10 ISP codes if you want your rule to apply to more than one ISP.", "zh_CN": "该规则适用的运营商。可调用我们的'查询支持的ISP运营商列表'接口查看运营商信息。指定'all'表示所有运营商。如果希望该规则应用于多个运营商，则可指定多个运营商，用逗号分隔，但最多只能包含10个运营商。"}
  Isp *string `json:"isp,omitempty" xml:"isp,omitempty"`
}

func (s CreateAnEdgeHostnameRequestClientZones) String() string {
  return tea.Prettify(s)
}

func (s CreateAnEdgeHostnameRequestClientZones) GoString() string {
  return s.String()
}

func (s *CreateAnEdgeHostnameRequestClientZones) SetRegion(v string) *CreateAnEdgeHostnameRequestClientZones {
  s.Region = &v
  return s
}

func (s *CreateAnEdgeHostnameRequestClientZones) SetWeight(v int) *CreateAnEdgeHostnameRequestClientZones {
  s.Weight = &v
  return s
}

func (s *CreateAnEdgeHostnameRequestClientZones) SetAction(v *CreateAnEdgeHostnameRequestClientZonesAction) *CreateAnEdgeHostnameRequestClientZones {
  s.Action = v
  return s
}

func (s *CreateAnEdgeHostnameRequestClientZones) SetIsp(v string) *CreateAnEdgeHostnameRequestClientZones {
  s.Isp = &v
  return s
}

type CreateAnEdgeHostnameRequestClientZonesAction struct {
  // {"en" : "Enum: deliver,redirect,reject 
  // Defines the action to take for requests to the zone. Options are to deliver using one or more server groups, to reject the request altogether, or to redirect to another domain. If 'reject' is specified, the client request will be redirected to a server that always responds with HTTP response code 403 representing 'Forbidden'. Up to one 'reject' action is allowed for each client zone.
  // ", "zh_CN": "取值范围: deliver,redirect,reject 
  // 当规则匹配时，对客户端请求执行的动作的类型。包括分发，拒绝和跳转3个类型。如果指定了'拒绝'，则客户端请求将被调度到一台服务器，该服务器总是响应403状态码，表示'Forbidden'拒绝访问。每个访客分区最多只允许一个'拒绝'动作。"}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
  // {"en" : "This field is used if the action is of type 'deliver'. Specify one or more of the server groups (standard, premium, deluxe,  ultra, nearChina, ChinaStandard, ChinaPremium) to control the servers used to deliver content. If unspecified, 'standard' is used. <br/><br/>'nearChina' is a special server group that can be enabled for you if your domains lack ICP Beian, but you want optimal performance serving Chinese visitors. A client zone rule using nearChina cannot be configured to simultaneously deliver to other server groups, though you may create separate client zone rules to use other server groups.<br/><br/>If you have an ICP Beian license and want your content to be delivered by servers within China, you can opt to use 'ChinaStandard' and 'ChinaPremium.' Please contact our support team if you require nearChina, ChinaStandard, or ChinaPremium.<br/><br/>If an edge hostname is not initially configured to use ChinaStandard or ChinaPremium, you will not be able to modify it later to use these two server groups. Instead, you will need to create a new edge hostname with client zone rules delivering to those server groups.", "zh_CN": "如果动作类型为'分发'，则使用此字段指定一个或多个节点组(standard, premium, deluxe,  ultra, nearChina, ChinaStandard, ChinaPremium) 来选择提供内容分发服务的缓存节点。如果未指定，则使用'standard'。<br/><br/>'nearChina'  是一个特殊的节点组。如果您需要使用nearChina节点组，请联系我们的技术支持开通。不能在同一条规则中同时指定nearChina节点组和其他节点组。如果要使用其他节点组，需要创建单独的访客分区规则。<br/><br/>如果您的加速域名有ICP备案，希望由中国大陆的服务器提供内容分发服务，您可以选择使用'ChinaStandard'和'ChinaPremium'节点组。
  // <br/><br/>如果调度域名创建时没有指定使用ChinaStandard或ChinaPremium节点组，则无法通过更新调度域名来使用这两个节点组。您需要创建一个新的调度域名，在新的调度域名中指定使用ChinaStandard或ChinaPremium，才能使用这2个节点组。"}
  By []*string `json:"by,omitempty" xml:"by,omitempty" type:"Repeated"`
  // {"en" : "This field is used if the action is of type 'redirect'. It indicates the hostname or IP address to redirect to. This is typically an origin server or another CDN provider.", "zh_CN": "如果动作类型为'跳转'，则通过该字段指定跳转的目标域名或IP地址。'跳转'目标通常是源站服务器或其它CDN厂商。"}
  To []*string `json:"to,omitempty" xml:"to,omitempty" type:"Repeated"`
  // {"en" : "Default: True 
  // Indicates whether an IPv6 address can be used. This value is used only if the client zone rule's action type is 'deliver'.", "zh_CN": "默认值: True 
  // 指定是否允许使用IPv6地址进行内容分发。仅当动作类型为'分发'时该值才有效。"}
  EnableIPv6 *bool `json:"enableIPv6,omitempty" xml:"enableIPv6,omitempty"`
}

func (s CreateAnEdgeHostnameRequestClientZonesAction) String() string {
  return tea.Prettify(s)
}

func (s CreateAnEdgeHostnameRequestClientZonesAction) GoString() string {
  return s.String()
}

func (s *CreateAnEdgeHostnameRequestClientZonesAction) SetType(v string) *CreateAnEdgeHostnameRequestClientZonesAction {
  s.Type = &v
  return s
}

func (s *CreateAnEdgeHostnameRequestClientZonesAction) SetBy(v []*string) *CreateAnEdgeHostnameRequestClientZonesAction {
  s.By = v
  return s
}

func (s *CreateAnEdgeHostnameRequestClientZonesAction) SetTo(v []*string) *CreateAnEdgeHostnameRequestClientZonesAction {
  s.To = v
  return s
}

func (s *CreateAnEdgeHostnameRequestClientZonesAction) SetEnableIPv6(v bool) *CreateAnEdgeHostnameRequestClientZonesAction {
  s.EnableIPv6 = &v
  return s
}

type CreateAnEdgeHostnameResponse struct {
}

func (s CreateAnEdgeHostnameResponse) String() string {
  return tea.Prettify(s)
}

func (s CreateAnEdgeHostnameResponse) GoString() string {
  return s.String()
}

type CreateAnEdgeHostnameResponseHeader struct {
  // {"en":"The value is a URL to the new edge hostname.", "zh_CN":"通过Location响应头返回新建的调度域名的URL。可使用该URL调用'查询调度域名详情'接口来查看调度域名的详细信息。URL示例：<code>Location: http://open.chinanetcenter.com/cdn/edgeHostnames/abcde12345.qtlcdn.com"}
  Location *string `json:"Location,omitempty" xml:"Location,omitempty" require:"true"`
}

func (s CreateAnEdgeHostnameResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s CreateAnEdgeHostnameResponseHeader) GoString() string {
  return s.String()
}

func (s *CreateAnEdgeHostnameResponseHeader) SetLocation(v string) *CreateAnEdgeHostnameResponseHeader {
  s.Location = &v
  return s
}




type GetAnEdgeHostnamePaths struct {
  // {"en" : "an edge hostname", "zh_CN": "调度域名。"}
  EdgeHostname *string `json:"edgeHostname,omitempty" xml:"edgeHostname,omitempty" require:"true"`
}

func (s GetAnEdgeHostnamePaths) String() string {
  return tea.Prettify(s)
}

func (s GetAnEdgeHostnamePaths) GoString() string {
  return s.String()
}

func (s *GetAnEdgeHostnamePaths) SetEdgeHostname(v string) *GetAnEdgeHostnamePaths {
  s.EdgeHostname = &v
  return s
}

type GetAnEdgeHostnameParameters struct {
}

func (s GetAnEdgeHostnameParameters) String() string {
  return tea.Prettify(s)
}

func (s GetAnEdgeHostnameParameters) GoString() string {
  return s.String()
}

type GetAnEdgeHostnameRequestHeader struct {
}

func (s GetAnEdgeHostnameRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetAnEdgeHostnameRequestHeader) GoString() string {
  return s.String()
}

type GetAnEdgeHostnameRequest struct {
}

func (s GetAnEdgeHostnameRequest) String() string {
  return tea.Prettify(s)
}

func (s GetAnEdgeHostnameRequest) GoString() string {
  return s.String()
}

type GetAnEdgeHostnameResponseHeader struct {
}

func (s GetAnEdgeHostnameResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetAnEdgeHostnameResponseHeader) GoString() string {
  return s.String()
}

type GetAnEdgeHostnameResponse struct {
  // {"en" : "Operation history", "zh_CN": "调度域名的操作记录。"}
  History []*GetAnEdgeHostnameResponseHistory `json:"history,omitempty" xml:"history,omitempty" require:"true" type:"Repeated"`
  // {"en" : "RFC 3339 format date indicating when the edge hostname was last updated.", "zh_CN": "RFC 3339格式的日期，表示调度域名最后一次更新的时间。"}
  LastUpdateTime *string `json:"lastUpdateTime,omitempty" xml:"lastUpdateTime,omitempty" require:"true"`
  // {"en" : "RFC 3339 format date indicating when the edge hostname was created.", "zh_CN": "RFC 3339格式的日期，表示调度域名的创建时间。"}
  CreationTime *string `json:"creationTime,omitempty" xml:"creationTime,omitempty" require:"true"`
  // {"en" : "Contains settings of the edge hostname.", "zh_CN": "调度域名的具体配置。此处展示的内容是调度域名当前已部署生效的配置。如果您提交了更新调度域名的请求，但配置仍在部署中还未生效，此处不会展示您修改后的配置。如果调度域名暂时还没有部署生效的配置，此处展示的是最近一次提交的配置。"}
  Configs *GetAnEdgeHostnameResponseConfigs `json:"configs,omitempty" xml:"configs,omitempty" require:"true" type:"Struct"`
}

func (s GetAnEdgeHostnameResponse) String() string {
  return tea.Prettify(s)
}

func (s GetAnEdgeHostnameResponse) GoString() string {
  return s.String()
}

func (s *GetAnEdgeHostnameResponse) SetHistory(v []*GetAnEdgeHostnameResponseHistory) *GetAnEdgeHostnameResponse {
  s.History = v
  return s
}

func (s *GetAnEdgeHostnameResponse) SetLastUpdateTime(v string) *GetAnEdgeHostnameResponse {
  s.LastUpdateTime = &v
  return s
}

func (s *GetAnEdgeHostnameResponse) SetCreationTime(v string) *GetAnEdgeHostnameResponse {
  s.CreationTime = &v
  return s
}

func (s *GetAnEdgeHostnameResponse) SetConfigs(v *GetAnEdgeHostnameResponseConfigs) *GetAnEdgeHostnameResponse {
  s.Configs = v
  return s
}

type GetAnEdgeHostnameResponseHistory struct     {
  // {"en" : "Enum: creation,update,deletion 
  // Indicates the action taken.", "zh_CN": "取值范围: creation,update,deletion 
  // 操作类型，即创建，更新或删除。"}
  Operation *string `json:"operation,omitempty" xml:"operation,omitempty"`
  // {"en" : "RFC 3339 format date of the event. Same as submissionTime. ", "zh_CN": "操作时间，以RFC 3339日期格式表示。该字段的值等同于submissionTime。"}
  Time *string `json:"time,omitempty" xml:"time,omitempty"`
  // {"en" : "An ID representing the API call that made the change.", "zh_CN": "API请求的ID。"}
  ApiRequestId *string `json:"apiRequestId,omitempty" xml:"apiRequestId,omitempty"`
  // {"en" : "API account that made the request.", "zh_CN": "调用接口的API账号的名称。"}
  ApiName *string `json:"apiName,omitempty" xml:"apiName,omitempty"`
  // {"en" : "RFC 3339 format date of the event.", "zh_CN": "操作时间，以RFC 3339日期格式表示。"}
  SubmissionTime *string `json:"submissionTime,omitempty" xml:"submissionTime,omitempty"`
  // {"en" : "RFC 3339 format date indicating when the deployment completes.", "zh_CN": "配置异步部署的结束时间，以RFC 3339日期格式表示。"}
  FinishTime *string `json:"finishTime,omitempty" xml:"finishTime,omitempty"`
  // {"en" : "Enum: waiting,in-progress,succeeded,failed
  // Deployment status.", "zh_CN": "取值范围: waiting,in-progress,succeeded,failed
  // 配置部署的状态。"}
  Status *string `json:"status,omitempty" xml:"status,omitempty"`
  // {"en" : "Details describing the status.", "zh_CN": "部署状态的描述信息。"}
  StatusDetails *string `json:"statusDetails,omitempty" xml:"statusDetails,omitempty"`
  // {"en" : "Contains settings that can be modified.", "zh_CN": "调度域名的相关设置。"}
  Configuration *GetAnEdgeHostnameResponseHistoryConfiguration `json:"configuration,omitempty" xml:"configuration,omitempty" require:"true" type:"Struct"`
}

func (s GetAnEdgeHostnameResponseHistory) String() string {
  return tea.Prettify(s)
}

func (s GetAnEdgeHostnameResponseHistory) GoString() string {
  return s.String()
}

func (s *GetAnEdgeHostnameResponseHistory) SetOperation(v string) *GetAnEdgeHostnameResponseHistory {
  s.Operation = &v
  return s
}

func (s *GetAnEdgeHostnameResponseHistory) SetTime(v string) *GetAnEdgeHostnameResponseHistory {
  s.Time = &v
  return s
}

func (s *GetAnEdgeHostnameResponseHistory) SetApiRequestId(v string) *GetAnEdgeHostnameResponseHistory {
  s.ApiRequestId = &v
  return s
}

func (s *GetAnEdgeHostnameResponseHistory) SetApiName(v string) *GetAnEdgeHostnameResponseHistory {
  s.ApiName = &v
  return s
}

func (s *GetAnEdgeHostnameResponseHistory) SetSubmissionTime(v string) *GetAnEdgeHostnameResponseHistory {
  s.SubmissionTime = &v
  return s
}

func (s *GetAnEdgeHostnameResponseHistory) SetFinishTime(v string) *GetAnEdgeHostnameResponseHistory {
  s.FinishTime = &v
  return s
}

func (s *GetAnEdgeHostnameResponseHistory) SetStatus(v string) *GetAnEdgeHostnameResponseHistory {
  s.Status = &v
  return s
}

func (s *GetAnEdgeHostnameResponseHistory) SetStatusDetails(v string) *GetAnEdgeHostnameResponseHistory {
  s.StatusDetails = &v
  return s
}

func (s *GetAnEdgeHostnameResponseHistory) SetConfiguration(v *GetAnEdgeHostnameResponseHistoryConfiguration) *GetAnEdgeHostnameResponseHistory {
  s.Configuration = v
  return s
}

type GetAnEdgeHostnameResponseHistoryConfiguration struct {
  // {"en" : "An edge hostname.", "zh_CN": "调度域名。"}
  EdgeHostname *string `json:"edgeHostname,omitempty" xml:"edgeHostname,omitempty"`
  // {"en" : "Default: False 
  // Indicates whether you have a Beian license for service in China. If you do, content can be served by servers located within mainland China using the ChinaStandard or ChinaPremium server group.
  // ", "zh_CN": "默认值: False 
  // 是否备案的标记。当值为true时，将使用包括中国大陆的服务器提供内容分发服务。"}
  HasBeian *bool `json:"hasBeian,omitempty" xml:"hasBeian,omitempty"`
  // {"en" : "Rules control how requests from client zones are handled. There must be a rule that covers all regions and all ISPs. ", "zh_CN": "自定义规则来控制如何处理不同访客分区的请求。您必须至少创建一条覆盖所有区域和所有运营商的规则。"}
  ClientZones []*GetAnEdgeHostnameResponseHistoryConfigurationClientZones `json:"clientZones,omitempty" xml:"clientZones,omitempty" type:"Repeated"`
  // {"en" : "An estimate of the bandwidth required to serve content using this edge hostname. Units of measurement should be in Tbps, Gbps, Mbps, or kbps. Example: 100Gbps", "zh_CN": "通过该调度域名进行CDN加速预计需要的带宽。单位应为Tbps、Gbps、Mbps或kbps。示例：100 Gbps。"}
  EstimatedBandwidth *string `json:"estimatedBandwidth,omitempty" xml:"estimatedBandwidth,omitempty"`
}

func (s GetAnEdgeHostnameResponseHistoryConfiguration) String() string {
  return tea.Prettify(s)
}

func (s GetAnEdgeHostnameResponseHistoryConfiguration) GoString() string {
  return s.String()
}

func (s *GetAnEdgeHostnameResponseHistoryConfiguration) SetEdgeHostname(v string) *GetAnEdgeHostnameResponseHistoryConfiguration {
  s.EdgeHostname = &v
  return s
}

func (s *GetAnEdgeHostnameResponseHistoryConfiguration) SetHasBeian(v bool) *GetAnEdgeHostnameResponseHistoryConfiguration {
  s.HasBeian = &v
  return s
}

func (s *GetAnEdgeHostnameResponseHistoryConfiguration) SetClientZones(v []*GetAnEdgeHostnameResponseHistoryConfigurationClientZones) *GetAnEdgeHostnameResponseHistoryConfiguration {
  s.ClientZones = v
  return s
}

func (s *GetAnEdgeHostnameResponseHistoryConfiguration) SetEstimatedBandwidth(v string) *GetAnEdgeHostnameResponseHistoryConfiguration {
  s.EstimatedBandwidth = &v
  return s
}

type GetAnEdgeHostnameResponseHistoryConfigurationClientZones struct     {
  // {"en" : "This field indicates the region in which the rule applies. Refer to our API to get client regions to get valid region codes. For example, if you wish to create a rule that covers all of Europe, simply specify 'eu' as the region. You can indicate specific countries. For example, 'na.us' represents the 'United States of America', and 'eu.fr' represents 'France'.  
  // 
  // A special client region 'all' can be used to specify that the rule applies to the entire world. If overlapping regions are specified, the more specific one takes precedence. For example, if you specify 'as' in one rule and 'as.cn' in another, a request from China will follow the rule for 'as.cn'.", "zh_CN": "该规则适用的区域。可调用'查询支持的区域列表'接口来查看区域信息。例如，如果您希望创建规则覆盖整个欧洲，则指定'eu'为区域。 您可以指定具体的国家。例如，'na.us'代表'美国'，而'eu.fr'代表'法国'。
  // 
  // 'all'是一个特殊的区域，可用于指定适用于全球的规则。如果不同规则指定的区域存在重叠，则以更细粒度的区域优先。例如，如果您在一条规则中指定'as'，在另一条规则中指定'as.cn'，则来自中国的请求将优先匹配'as.cn'的规则。"}
  Region *string `json:"region,omitempty" xml:"region,omitempty"`
  // {"en" : "Default: 100 Range: [ 0 .. 100 ] 
  // When multiple actions are specified for a client zone, the weight is used to adjust the probability of a client zone rule applying to a client request.
  // Consider two rules that apply to 'as.cn': {'region':'as.cn', 'isp':'all', 'action':{'type':'deliver', 'by':['standard', 'premium', 'deluxe']}},
  //     {'region':'as.cn', 'isp':'all', 'action':{'type':'redirect', 'to':['alternate.cname.com']}, 'weight':10}
  // 
  //  The probability of 'deliver' will be 100/(100+10) or 0.909 while the probability of 'redirect' will be 10/(100+10) or 0.091.
  //  
  // ", "zh_CN": "默认值: 100 取值范围: [ 0 .. 100 ] 
  // 可以为同个访客分区指定多条规则。通过在规则中指定weight字段，可控制规则匹配的权重。
  // 以'as.cn'区域的2条规则为例: {'region':'as.cn', 'isp':'all', 'action':{'type':'deliver', 'by':['standard', 'premium', 'deluxe']}},
  //     {'region':'as.cn', 'isp':'all', 'action':{'type':'redirect', 'to':['alternate.cname.com']}, 'weight':10}
  // 
  // 按照以上规则，客户端请求匹配规则1进行'分发'的比例为100/(100+10)，即0.909，匹配规则2进行'跳转'的比例为10/(100+10)，即0.091。
  //  
  // "}
  Weight *int `json:"weight,omitempty" xml:"weight,omitempty"`
  // {"en" : "This object describes the action to take for requests matching the rule.", "zh_CN": "当规则匹配时执行的动作。"}
  Action *GetAnEdgeHostnameResponseHistoryConfigurationClientZonesAction `json:"action,omitempty" xml:"action,omitempty" type:"Struct"`
  // {"en" : "Specify the code representing an ISP (Internet Service Provider) if the rule only applies to requests from a particular ISP. Call our API to get a list of supported ISPs. Specify 'all' to indicate all ISPs rather than a particular one. Specify a comma-separated list of up to 10 ISP codes if you want your rule to apply to more than one ISP.", "zh_CN": "该规则适用的运营商。可调用我们的'查询支持的ISP运营商列表'接口查看运营商信息。指定'all'表示所有运营商。如果希望该规则应用于多个运营商，则可指定多个运营商，用逗号分隔，但最多只能包含10个运营商。"}
  Isp *string `json:"isp,omitempty" xml:"isp,omitempty"`
}

func (s GetAnEdgeHostnameResponseHistoryConfigurationClientZones) String() string {
  return tea.Prettify(s)
}

func (s GetAnEdgeHostnameResponseHistoryConfigurationClientZones) GoString() string {
  return s.String()
}

func (s *GetAnEdgeHostnameResponseHistoryConfigurationClientZones) SetRegion(v string) *GetAnEdgeHostnameResponseHistoryConfigurationClientZones {
  s.Region = &v
  return s
}

func (s *GetAnEdgeHostnameResponseHistoryConfigurationClientZones) SetWeight(v int) *GetAnEdgeHostnameResponseHistoryConfigurationClientZones {
  s.Weight = &v
  return s
}

func (s *GetAnEdgeHostnameResponseHistoryConfigurationClientZones) SetAction(v *GetAnEdgeHostnameResponseHistoryConfigurationClientZonesAction) *GetAnEdgeHostnameResponseHistoryConfigurationClientZones {
  s.Action = v
  return s
}

func (s *GetAnEdgeHostnameResponseHistoryConfigurationClientZones) SetIsp(v string) *GetAnEdgeHostnameResponseHistoryConfigurationClientZones {
  s.Isp = &v
  return s
}

type GetAnEdgeHostnameResponseHistoryConfigurationClientZonesAction struct {
  // {"en" : "Enum: deliver,redirect,reject 
  // Defines the action to take for requests to the zone. Options are to deliver using one or more server groups, to reject the request altogether, or to redirect to another domain. If 'reject' is specified, the client request will be redirected to a server that always responds with HTTP response code 403 representing 'Forbidden'. Up to one 'reject' action is allowed for each client zone.
  // ", "zh_CN": "取值范围: deliver,redirect,reject 
  // 当规则匹配时，对客户端请求执行的动作的类型。包括分发，拒绝和跳转3个类型。如果指定了'拒绝'，则客户端请求将被调度到一台服务器，该服务器总是响应403状态码，表示'Forbidden'拒绝访问。每个访客分区最多只允许一个'拒绝'动作。"}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
  // {"en" : "This field is used if the action is of type 'deliver'. Specify one or more of the server groups (standard, premium, deluxe,  ultra, nearChina, ChinaStandard, ChinaPremium) to control the servers used to deliver content. If unspecified, 'standard' is used. <br/><br/>'nearChina' is a special server group that can be enabled for you if your domains lack ICP Beian, but you want optimal performance serving Chinese visitors. A client zone rule using nearChina cannot be configured to simultaneously deliver to other server groups, though you may create separate client zone rules to use other server groups.<br/><br/>If you have an ICP Beian license and want your content to be delivered by servers within China, you can opt to use 'ChinaStandard' and 'ChinaPremium.' Please contact our support team if you require nearChina, ChinaStandard, or ChinaPremium.<br/><br/>If an edge hostname is not initially configured to use ChinaStandard or ChinaPremium, you will not be able to modify it later to use these two server groups. Instead, you will need to create a new edge hostname with client zone rules delivering to those server groups.", "zh_CN": "如果动作类型为'分发'，则使用此字段指定一个或多个节点组(standard, premium, deluxe,  ultra, nearChina, ChinaStandard, ChinaPremium) 来选择提供内容分发服务的缓存节点。如果未指定，则使用'standard'。<br/><br/>'nearChina'  是一个特殊的节点组。如果您需要使用nearChina节点组，请联系我们的技术支持开通。不能在同一条规则中同时指定nearChina节点组和其他节点组。如果要使用其他节点组，需要创建单独的访客分区规则。<br/><br/>如果您的加速域名有ICP备案，希望由中国大陆的服务器提供内容分发服务，您可以选择使用'ChinaStandard'和'ChinaPremium'节点组。
  // <br/><br/>如果调度域名创建时没有指定使用ChinaStandard或ChinaPremium节点组，则无法通过更新调度域名来使用这两个节点组。您需要创建一个新的调度域名，在新的调度域名中指定使用ChinaStandard或ChinaPremium，才能使用这2个节点组。"}
  By []*string `json:"by,omitempty" xml:"by,omitempty" type:"Repeated"`
  // {"en" : "This field is used if the action is of type 'redirect'. It indicates the hostname or IP address to redirect to. This is typically an origin server or another CDN provider.", "zh_CN": "如果动作类型为'跳转'，则通过该字段指定跳转的目标域名或IP地址。'跳转'目标通常是源站服务器或其它CDN厂商。"}
  To []*string `json:"to,omitempty" xml:"to,omitempty" type:"Repeated"`
  // {"en" : "Default: True 
  // Indicates whether an IPv6 address can be used. This value is used only if the client zone rule's action type is 'deliver'.", "zh_CN": "默认值: True 
  // 指定是否允许使用IPv6地址进行内容分发。仅当动作类型为'分发'时该值才有效。"}
  EnableIPv6 *bool `json:"enableIPv6,omitempty" xml:"enableIPv6,omitempty"`
}

func (s GetAnEdgeHostnameResponseHistoryConfigurationClientZonesAction) String() string {
  return tea.Prettify(s)
}

func (s GetAnEdgeHostnameResponseHistoryConfigurationClientZonesAction) GoString() string {
  return s.String()
}

func (s *GetAnEdgeHostnameResponseHistoryConfigurationClientZonesAction) SetType(v string) *GetAnEdgeHostnameResponseHistoryConfigurationClientZonesAction {
  s.Type = &v
  return s
}

func (s *GetAnEdgeHostnameResponseHistoryConfigurationClientZonesAction) SetBy(v []*string) *GetAnEdgeHostnameResponseHistoryConfigurationClientZonesAction {
  s.By = v
  return s
}

func (s *GetAnEdgeHostnameResponseHistoryConfigurationClientZonesAction) SetTo(v []*string) *GetAnEdgeHostnameResponseHistoryConfigurationClientZonesAction {
  s.To = v
  return s
}

func (s *GetAnEdgeHostnameResponseHistoryConfigurationClientZonesAction) SetEnableIPv6(v bool) *GetAnEdgeHostnameResponseHistoryConfigurationClientZonesAction {
  s.EnableIPv6 = &v
  return s
}

type GetAnEdgeHostnameResponseConfigs struct {
  // {"en" : "An edge hostname.", "zh_CN": "调度域名。"}
  EdgeHostname *string `json:"edgeHostname,omitempty" xml:"edgeHostname,omitempty"`
  // {"en" : "Default: False 
  // Indicates whether PCI compliance is required.  true means content will only be served by PCI certified PoPs.", "zh_CN": "默认值: False 
  // 表示流量调度是否需要遵循PCI规范。当值为true时，表示只能使用已通过PCI认证的节点提供内容分发服务。"}
  PciRequired *bool `json:"pciRequired,omitempty" xml:"pciRequired,omitempty"`
  // {"en" : "A description of the edge hostname.", "zh_CN": "调度域名的描述。"}
  Description *string `json:"description,omitempty" xml:"description,omitempty"`
  // {"en" : "Default: False 
  // Indicates whether you have a Beian license for service in China. If you do, content can be served by servers located within mainland China using the ChinaStandard or ChinaPremium server group.
  // ", "zh_CN": "默认值: False 
  // 是否备案的标记。当值为true时，将使用包括中国大陆的服务器提供内容分发服务。"}
  HasBeian *bool `json:"hasBeian,omitempty" xml:"hasBeian,omitempty"`
  // {"en" : "Rules control how requests from client zones are handled. There must be a rule that covers all regions and all ISPs. ", "zh_CN": "自定义规则来控制如何处理不同访客分区的请求。您必须至少创建一条覆盖所有区域和所有运营商的规则。"}
  ClientZones []*GetAnEdgeHostnameResponseConfigsClientZones `json:"clientZones,omitempty" xml:"clientZones,omitempty" type:"Repeated"`
  // {"en" : "An estimate of the bandwidth required to serve content using this edge hostname. Units of measurement should be in Tbps, Gbps, Mbps, or kbps. Example: 100Gbps", "zh_CN": "通过该调度域名进行CDN加速预计需要的带宽。单位应为Tbps、Gbps、Mbps或kbps。示例：100 Gbps。"}
  EstimatedBandwidth *string `json:"estimatedBandwidth,omitempty" xml:"estimatedBandwidth,omitempty"`
}

func (s GetAnEdgeHostnameResponseConfigs) String() string {
  return tea.Prettify(s)
}

func (s GetAnEdgeHostnameResponseConfigs) GoString() string {
  return s.String()
}

func (s *GetAnEdgeHostnameResponseConfigs) SetEdgeHostname(v string) *GetAnEdgeHostnameResponseConfigs {
  s.EdgeHostname = &v
  return s
}

func (s *GetAnEdgeHostnameResponseConfigs) SetPciRequired(v bool) *GetAnEdgeHostnameResponseConfigs {
  s.PciRequired = &v
  return s
}

func (s *GetAnEdgeHostnameResponseConfigs) SetDescription(v string) *GetAnEdgeHostnameResponseConfigs {
  s.Description = &v
  return s
}

func (s *GetAnEdgeHostnameResponseConfigs) SetHasBeian(v bool) *GetAnEdgeHostnameResponseConfigs {
  s.HasBeian = &v
  return s
}

func (s *GetAnEdgeHostnameResponseConfigs) SetClientZones(v []*GetAnEdgeHostnameResponseConfigsClientZones) *GetAnEdgeHostnameResponseConfigs {
  s.ClientZones = v
  return s
}

func (s *GetAnEdgeHostnameResponseConfigs) SetEstimatedBandwidth(v string) *GetAnEdgeHostnameResponseConfigs {
  s.EstimatedBandwidth = &v
  return s
}

type GetAnEdgeHostnameResponseConfigsClientZones struct     {
  // {"en" : "This field indicates the region in which the rule applies. Refer to our API to get client regions to get valid region codes. For example, if you wish to create a rule that covers all of Europe, simply specify 'eu' as the region. You can indicate specific countries. For example, 'na.us' represents the 'United States of America', and 'eu.fr' represents 'France'.  
  // 
  // A special client region 'all' can be used to specify that the rule applies to the entire world. If overlapping regions are specified, the more specific one takes precedence. For example, if you specify 'as' in one rule and 'as.cn' in another, a request from China will follow the rule for 'as.cn'.", "zh_CN": "该规则适用的区域。可调用'查询支持的区域列表'接口来查看区域信息。例如，如果您希望创建规则覆盖整个欧洲，则指定'eu'为区域。 您可以指定具体的国家。例如，'na.us'代表'美国'，而'eu.fr'代表'法国'。
  // 
  // 'all'是一个特殊的区域，可用于指定适用于全球的规则。如果不同规则指定的区域存在重叠，则以更细粒度的区域优先。例如，如果您在一条规则中指定'as'，在另一条规则中指定'as.cn'，则来自中国的请求将优先匹配'as.cn'的规则。"}
  Region *string `json:"region,omitempty" xml:"region,omitempty"`
  // {"en" : "Default: 100 Range: [ 0 .. 100 ] 
  // When multiple actions are specified for a client zone, the weight is used to adjust the probability of a client zone rule applying to a client request.
  // Consider two rules that apply to 'as.cn': {'region':'as.cn', 'isp':'all', 'action':{'type':'deliver', 'by':['standard', 'premium', 'deluxe']}},
  //     {'region':'as.cn', 'isp':'all', 'action':{'type':'redirect', 'to':['alternate.cname.com']}, 'weight':10}
  // 
  //  The probability of 'deliver' will be 100/(100+10) or 0.909 while the probability of 'redirect' will be 10/(100+10) or 0.091.
  //  
  // ", "zh_CN": "默认值: 100 取值范围: [ 0 .. 100 ] 
  // 可以为同个访客分区指定多条规则。通过在规则中指定weight字段，可控制规则匹配的权重。
  // 以'as.cn'区域的2条规则为例: {'region':'as.cn', 'isp':'all', 'action':{'type':'deliver', 'by':['standard', 'premium', 'deluxe']}},
  //     {'region':'as.cn', 'isp':'all', 'action':{'type':'redirect', 'to':['alternate.cname.com']}, 'weight':10}
  // 
  // 按照以上规则，客户端请求匹配规则1进行'分发'的比例为100/(100+10)，即0.909，匹配规则2进行'跳转'的比例为10/(100+10)，即0.091。
  //  
  // "}
  Weight *int `json:"weight,omitempty" xml:"weight,omitempty"`
  // {"en" : "This object describes the action to take for requests matching the rule.", "zh_CN": "当规则匹配时执行的动作。"}
  Action *GetAnEdgeHostnameResponseConfigsClientZonesAction `json:"action,omitempty" xml:"action,omitempty" type:"Struct"`
  // {"en" : "Specify the code representing an ISP (Internet Service Provider) if the rule only applies to requests from a particular ISP. Call our API to get a list of supported ISPs. Specify 'all' to indicate all ISPs rather than a particular one. Specify a comma-separated list of up to 10 ISP codes if you want your rule to apply to more than one ISP.", "zh_CN": "该规则适用的运营商。可调用我们的'查询支持的ISP运营商列表'接口查看运营商信息。指定'all'表示所有运营商。如果希望该规则应用于多个运营商，则可指定多个运营商，用逗号分隔，但最多只能包含10个运营商。"}
  Isp *string `json:"isp,omitempty" xml:"isp,omitempty"`
}

func (s GetAnEdgeHostnameResponseConfigsClientZones) String() string {
  return tea.Prettify(s)
}

func (s GetAnEdgeHostnameResponseConfigsClientZones) GoString() string {
  return s.String()
}

func (s *GetAnEdgeHostnameResponseConfigsClientZones) SetRegion(v string) *GetAnEdgeHostnameResponseConfigsClientZones {
  s.Region = &v
  return s
}

func (s *GetAnEdgeHostnameResponseConfigsClientZones) SetWeight(v int) *GetAnEdgeHostnameResponseConfigsClientZones {
  s.Weight = &v
  return s
}

func (s *GetAnEdgeHostnameResponseConfigsClientZones) SetAction(v *GetAnEdgeHostnameResponseConfigsClientZonesAction) *GetAnEdgeHostnameResponseConfigsClientZones {
  s.Action = v
  return s
}

func (s *GetAnEdgeHostnameResponseConfigsClientZones) SetIsp(v string) *GetAnEdgeHostnameResponseConfigsClientZones {
  s.Isp = &v
  return s
}

type GetAnEdgeHostnameResponseConfigsClientZonesAction struct {
  // {"en" : "Enum: deliver,redirect,reject 
  // Defines the action to take for requests to the zone. Options are to deliver using one or more server groups, to reject the request altogether, or to redirect to another domain. If 'reject' is specified, the client request will be redirected to a server that always responds with HTTP response code 403 representing 'Forbidden'. Up to one 'reject' action is allowed for each client zone.
  // ", "zh_CN": "取值范围: deliver,redirect,reject 
  // 当规则匹配时，对客户端请求执行的动作的类型。包括分发，拒绝和跳转3个类型。如果指定了'拒绝'，则客户端请求将被调度到一台服务器，该服务器总是响应403状态码，表示'Forbidden'拒绝访问。每个访客分区最多只允许一个'拒绝'动作。"}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
  // {"en" : "This field is used if the action is of type 'deliver'. Specify one or more of the server groups (standard, premium, deluxe,  ultra, nearChina, ChinaStandard, ChinaPremium) to control the servers used to deliver content. If unspecified, 'standard' is used. <br/><br/>'nearChina' is a special server group that can be enabled for you if your domains lack ICP Beian, but you want optimal performance serving Chinese visitors. A client zone rule using nearChina cannot be configured to simultaneously deliver to other server groups, though you may create separate client zone rules to use other server groups.<br/><br/>If you have an ICP Beian license and want your content to be delivered by servers within China, you can opt to use 'ChinaStandard' and 'ChinaPremium.' Please contact our support team if you require nearChina, ChinaStandard, or ChinaPremium.<br/><br/>If an edge hostname is not initially configured to use ChinaStandard or ChinaPremium, you will not be able to modify it later to use these two server groups. Instead, you will need to create a new edge hostname with client zone rules delivering to those server groups.", "zh_CN": "如果动作类型为'分发'，则使用此字段指定一个或多个节点组(standard, premium, deluxe,  ultra, nearChina, ChinaStandard, ChinaPremium) 来选择提供内容分发服务的缓存节点。如果未指定，则使用'standard'。<br/><br/>'nearChina'  是一个特殊的节点组。如果您需要使用nearChina节点组，请联系我们的技术支持开通。不能在同一条规则中同时指定nearChina节点组和其他节点组。如果要使用其他节点组，需要创建单独的访客分区规则。<br/><br/>如果您的加速域名有ICP备案，希望由中国大陆的服务器提供内容分发服务，您可以选择使用'ChinaStandard'和'ChinaPremium'节点组。
  // <br/><br/>如果调度域名创建时没有指定使用ChinaStandard或ChinaPremium节点组，则无法通过更新调度域名来使用这两个节点组。您需要创建一个新的调度域名，在新的调度域名中指定使用ChinaStandard或ChinaPremium，才能使用这2个节点组。"}
  By []*string `json:"by,omitempty" xml:"by,omitempty" type:"Repeated"`
  // {"en" : "This field is used if the action is of type 'redirect'. It indicates the hostname or IP address to redirect to. This is typically an origin server or another CDN provider.", "zh_CN": "如果动作类型为'跳转'，则通过该字段指定跳转的目标域名或IP地址。'跳转'目标通常是源站服务器或其它CDN厂商。"}
  To []*string `json:"to,omitempty" xml:"to,omitempty" type:"Repeated"`
  // {"en" : "Default: True 
  // Indicates whether an IPv6 address can be used. This value is used only if the client zone rule's action type is 'deliver'.", "zh_CN": "默认值: True 
  // 指定是否允许使用IPv6地址进行内容分发。仅当动作类型为'分发'时该值才有效。"}
  EnableIPv6 *bool `json:"enableIPv6,omitempty" xml:"enableIPv6,omitempty"`
}

func (s GetAnEdgeHostnameResponseConfigsClientZonesAction) String() string {
  return tea.Prettify(s)
}

func (s GetAnEdgeHostnameResponseConfigsClientZonesAction) GoString() string {
  return s.String()
}

func (s *GetAnEdgeHostnameResponseConfigsClientZonesAction) SetType(v string) *GetAnEdgeHostnameResponseConfigsClientZonesAction {
  s.Type = &v
  return s
}

func (s *GetAnEdgeHostnameResponseConfigsClientZonesAction) SetBy(v []*string) *GetAnEdgeHostnameResponseConfigsClientZonesAction {
  s.By = v
  return s
}

func (s *GetAnEdgeHostnameResponseConfigsClientZonesAction) SetTo(v []*string) *GetAnEdgeHostnameResponseConfigsClientZonesAction {
  s.To = v
  return s
}

func (s *GetAnEdgeHostnameResponseConfigsClientZonesAction) SetEnableIPv6(v bool) *GetAnEdgeHostnameResponseConfigsClientZonesAction {
  s.EnableIPv6 = &v
  return s
}




