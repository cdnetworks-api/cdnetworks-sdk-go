package nghostnames

// This file is auto-generated, don't edit it. Thanks.
import (
  "github.com/alibabacloud-go/tea/tea"
)

type GetListOfHostnamesThatHaveBeenDeployedPaths struct {
}

func (s GetListOfHostnamesThatHaveBeenDeployedPaths) String() string {
  return tea.Prettify(s)
}

func (s GetListOfHostnamesThatHaveBeenDeployedPaths) GoString() string {
  return s.String()
}

type GetListOfHostnamesThatHaveBeenDeployedParameters struct {
  // {"en" : "Default: 0 >= 0 
  // Index of the first hostname to return. Defaults to 0.", "zh_CN": "默认值: 0 取值范围：>= 0 
  // 查询起始位置。"}
  Offset *int `json:"offset,omitempty" xml:"offset,omitempty"`
  // {"en" : "Default: 100 [ 1 .. 200 ] 
  // Maximum number of hostnames to return.", "zh_CN": "默认值: 100 取值范围：<= 200 
  // 每次查询的最大条数。"}
  Limit *int `json:"limit,omitempty" xml:"limit,omitempty"`
  // {"en" : "Enum: all staging production 
  // Default: all 
  // Limit the response to hostnames from properties deployed to staging only or production only. By default, all hostnames from properties deployed to either production or staging are returned.", "zh_CN": "取值范围: all, staging, production 
  // 默认值: all 
  // 根据部署环境来查询加速域名。默认情况下，查询部署到演练和生产环境下的所有域名。"}
  Target *string `json:"target,omitempty" xml:"target,omitempty"`
}

func (s GetListOfHostnamesThatHaveBeenDeployedParameters) String() string {
  return tea.Prettify(s)
}

func (s GetListOfHostnamesThatHaveBeenDeployedParameters) GoString() string {
  return s.String()
}

func (s *GetListOfHostnamesThatHaveBeenDeployedParameters) SetOffset(v int) *GetListOfHostnamesThatHaveBeenDeployedParameters {
  s.Offset = &v
  return s
}

func (s *GetListOfHostnamesThatHaveBeenDeployedParameters) SetLimit(v int) *GetListOfHostnamesThatHaveBeenDeployedParameters {
  s.Limit = &v
  return s
}

func (s *GetListOfHostnamesThatHaveBeenDeployedParameters) SetTarget(v string) *GetListOfHostnamesThatHaveBeenDeployedParameters {
  s.Target = &v
  return s
}

type GetListOfHostnamesThatHaveBeenDeployedRequestHeader struct {
}

func (s GetListOfHostnamesThatHaveBeenDeployedRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetListOfHostnamesThatHaveBeenDeployedRequestHeader) GoString() string {
  return s.String()
}

type GetListOfHostnamesThatHaveBeenDeployedRequest struct {
}

func (s GetListOfHostnamesThatHaveBeenDeployedRequest) String() string {
  return tea.Prettify(s)
}

func (s GetListOfHostnamesThatHaveBeenDeployedRequest) GoString() string {
  return s.String()
}

type GetListOfHostnamesThatHaveBeenDeployedResponseHeader struct {
}

func (s GetListOfHostnamesThatHaveBeenDeployedResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetListOfHostnamesThatHaveBeenDeployedResponseHeader) GoString() string {
  return s.String()
}

type GetListOfHostnamesThatHaveBeenDeployedResponse struct {
  // {"en" : "List of hostnames.", "zh_CN": "加速域名列表。"}
  Hostnames []*GetListOfHostnamesThatHaveBeenDeployedResponseHostnames `json:"hostnames,omitempty" xml:"hostnames,omitempty" require:"true" type:"Repeated"`
  // {"en" : ">= 0 
  // Number of hostnames.", "zh_CN": "取值范围: >= 0 
  // 加速域名总数。"}
  Count *int `json:"count,omitempty" xml:"count,omitempty" require:"true"`
}

func (s GetListOfHostnamesThatHaveBeenDeployedResponse) String() string {
  return tea.Prettify(s)
}

func (s GetListOfHostnamesThatHaveBeenDeployedResponse) GoString() string {
  return s.String()
}

func (s *GetListOfHostnamesThatHaveBeenDeployedResponse) SetHostnames(v []*GetListOfHostnamesThatHaveBeenDeployedResponseHostnames) *GetListOfHostnamesThatHaveBeenDeployedResponse {
  s.Hostnames = v
  return s
}

func (s *GetListOfHostnamesThatHaveBeenDeployedResponse) SetCount(v int) *GetListOfHostnamesThatHaveBeenDeployedResponse {
  s.Count = &v
  return s
}

type GetListOfHostnamesThatHaveBeenDeployedResponseHostnames struct     {
  // {"en" : "A hostname whose property has been deployed to production or staging.", "zh_CN": "加速域名。"}
  Hostname *string `json:"hostname,omitempty" xml:"hostname,omitempty"`
  // {"en" : "Describes the property deployed to production. It can be 'null' if the property is not deployed to production.", "zh_CN": "如果加速域名对应的加速项目已部署到生产环境，则该字段返回该加速项目的相关信息。如果加速项目未部署到生产环境，则该字段返回'null'。"}
  PropertyInProduction *string `json:"propertyInProduction,omitempty" xml:"propertyInProduction,omitempty"`
  // {"en" : "Describes the property deployed to staging. It can be 'null' if the property is not deployed to staging.", "zh_CN": "如果加速域名对应的加速项目已部署到演练环境，则该字段返回该加速项目的相关信息。如果加速项目未部署到演练环境，则该字段返回'null'。"}
  PropertyInStaging *string `json:"propertyInStaging,omitempty" xml:"propertyInStaging,omitempty"`
}

func (s GetListOfHostnamesThatHaveBeenDeployedResponseHostnames) String() string {
  return tea.Prettify(s)
}

func (s GetListOfHostnamesThatHaveBeenDeployedResponseHostnames) GoString() string {
  return s.String()
}

func (s *GetListOfHostnamesThatHaveBeenDeployedResponseHostnames) SetHostname(v string) *GetListOfHostnamesThatHaveBeenDeployedResponseHostnames {
  s.Hostname = &v
  return s
}

func (s *GetListOfHostnamesThatHaveBeenDeployedResponseHostnames) SetPropertyInProduction(v string) *GetListOfHostnamesThatHaveBeenDeployedResponseHostnames {
  s.PropertyInProduction = &v
  return s
}

func (s *GetListOfHostnamesThatHaveBeenDeployedResponseHostnames) SetPropertyInStaging(v string) *GetListOfHostnamesThatHaveBeenDeployedResponseHostnames {
  s.PropertyInStaging = &v
  return s
}




type GetHistoricalInformationAboutOneHostnamePaths struct {
  // {"en" : "A hostname that was defined in a property.", "zh_CN": "加速域名。"}
  Hostname *string `json:"hostname,omitempty" xml:"hostname,omitempty" require:"true"`
}

func (s GetHistoricalInformationAboutOneHostnamePaths) String() string {
  return tea.Prettify(s)
}

func (s GetHistoricalInformationAboutOneHostnamePaths) GoString() string {
  return s.String()
}

func (s *GetHistoricalInformationAboutOneHostnamePaths) SetHostname(v string) *GetHistoricalInformationAboutOneHostnamePaths {
  s.Hostname = &v
  return s
}

type GetHistoricalInformationAboutOneHostnameParameters struct {
  // {"en" : "Beginning of the time period in RFC 3339 format. The time must be specified using the UTC timezone; it cannot be an offset. Example: startdate=2019-11-14T00:00:00Z By default, the value is when you began using CDN Pro.", "zh_CN": "指定查询开始时间，以RFC 3339日期格式表示。必须使用UTC时区，不支持指定其它时区。示例：startdate=2019-11-14T00:00:00Z。如果开始时间未指定，则默认为您开通CDN Pro服务的时间。"}
  Startdate *string `json:"startdate,omitempty" xml:"startdate,omitempty"`
  // {"en" : "End of the time period in RFC 3339 format. The time must be specified using the UTC timezone; it cannot be an offset. Example: enddate=2019-11-14T00:00:00Z The default is the current time.", "zh_CN": "指定查询结束时间，以RFC 3339日期格式表示。必须使用UTC时区，不支持指定其它时区。示例：enddate=2019-12-14T00:00:00Z。如果结束时间未指定，则默认为当前时间。"}
  Enddate *string `json:"enddate,omitempty" xml:"enddate,omitempty"`
  // {"en" : "Enum: all,staging,production 
  // Default: all 
  // The value can be 'all', 'staging', or 'production' to filter the results based on where the hostnames have been deployed. 
  // ", "zh_CN": "取值范围: all,staging,production 
  // 默认值: all 
  // 根据加速域名的部署环境过滤。该值可以是'all', 'staging', 或'production'，分别表示所有环境，演练环境和生产环境。"}
  Target *string `json:"target,omitempty" xml:"target,omitempty"`
}

func (s GetHistoricalInformationAboutOneHostnameParameters) String() string {
  return tea.Prettify(s)
}

func (s GetHistoricalInformationAboutOneHostnameParameters) GoString() string {
  return s.String()
}

func (s *GetHistoricalInformationAboutOneHostnameParameters) SetStartdate(v string) *GetHistoricalInformationAboutOneHostnameParameters {
  s.Startdate = &v
  return s
}

func (s *GetHistoricalInformationAboutOneHostnameParameters) SetEnddate(v string) *GetHistoricalInformationAboutOneHostnameParameters {
  s.Enddate = &v
  return s
}

func (s *GetHistoricalInformationAboutOneHostnameParameters) SetTarget(v string) *GetHistoricalInformationAboutOneHostnameParameters {
  s.Target = &v
  return s
}

type GetHistoricalInformationAboutOneHostnameRequestHeader struct {
}

func (s GetHistoricalInformationAboutOneHostnameRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetHistoricalInformationAboutOneHostnameRequestHeader) GoString() string {
  return s.String()
}

type GetHistoricalInformationAboutOneHostnameRequest struct {
}

func (s GetHistoricalInformationAboutOneHostnameRequest) String() string {
  return tea.Prettify(s)
}

func (s GetHistoricalInformationAboutOneHostnameRequest) GoString() string {
  return s.String()
}

type GetHistoricalInformationAboutOneHostnameResponseHeader struct {
}

func (s GetHistoricalInformationAboutOneHostnameResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetHistoricalInformationAboutOneHostnameResponseHeader) GoString() string {
  return s.String()
}

type GetHistoricalInformationAboutOneHostnameResponse struct {
  // {"en" : "A hostname that was defined in a property.", "zh_CN": "加速域名。"}
  Hostname *string `json:"hostname,omitempty" xml:"hostname,omitempty" require:"true"`
  // {"en" : "The history contains deployment and undeployment dates. It is empty if the hostname has never been deployed to production.", "zh_CN": "加速域名部署到生产环境或从生产环境卸载的历史信息。如果加速域名从未部署到生产环境，则返回空对象。"}
  History []*GetHistoricalInformationAboutOneHostnameResponseHistory `json:"history,omitempty" xml:"history,omitempty" require:"true" type:"Repeated"`
}

func (s GetHistoricalInformationAboutOneHostnameResponse) String() string {
  return tea.Prettify(s)
}

func (s GetHistoricalInformationAboutOneHostnameResponse) GoString() string {
  return s.String()
}

func (s *GetHistoricalInformationAboutOneHostnameResponse) SetHostname(v string) *GetHistoricalInformationAboutOneHostnameResponse {
  s.Hostname = &v
  return s
}

func (s *GetHistoricalInformationAboutOneHostnameResponse) SetHistory(v []*GetHistoricalInformationAboutOneHostnameResponseHistory) *GetHistoricalInformationAboutOneHostnameResponse {
  s.History = v
  return s
}

type GetHistoricalInformationAboutOneHostnameResponseHistory struct     {
  // {"en" : "The environment where the hostname is deployed.", "zh_CN": "加速域名所部署的环境。"}
  Target *string `json:"target,omitempty" xml:"target,omitempty"`
  // {"en" : "ID of the property that included the hostname.", "zh_CN": "加速域名对应的加速项目的ID。"}
  PropertyId *string `json:"propertyId,omitempty" xml:"propertyId,omitempty"`
  // {"en" : "RFC 3339 date with UTC time that indicates when the property with the hostname was deployed. Example: '2020-04-24T20:09:15Z'", "zh_CN": "RFC 3339格式的日期，表示加速项目的部署时间，采用UTC时区。示例：'2020-04-24T20:09:15Z'。"}
  DeploymentDate *string `json:"deploymentDate,omitempty" xml:"deploymentDate,omitempty"`
  // {"en" : "RFC 3339 date with UTC time that indicates when the property with the hostname was undeployed. Example: '2020-04-24T20:09:15Z'", "zh_CN": "RFC 3339格式的日期，表示加速项目的卸载时间，采用UTC时区。示例：'2020-04-24T20:09:15Z'。"}
  UndeploymentDate *string `json:"undeploymentDate,omitempty" xml:"undeploymentDate,omitempty"`
}

func (s GetHistoricalInformationAboutOneHostnameResponseHistory) String() string {
  return tea.Prettify(s)
}

func (s GetHistoricalInformationAboutOneHostnameResponseHistory) GoString() string {
  return s.String()
}

func (s *GetHistoricalInformationAboutOneHostnameResponseHistory) SetTarget(v string) *GetHistoricalInformationAboutOneHostnameResponseHistory {
  s.Target = &v
  return s
}

func (s *GetHistoricalInformationAboutOneHostnameResponseHistory) SetPropertyId(v string) *GetHistoricalInformationAboutOneHostnameResponseHistory {
  s.PropertyId = &v
  return s
}

func (s *GetHistoricalInformationAboutOneHostnameResponseHistory) SetDeploymentDate(v string) *GetHistoricalInformationAboutOneHostnameResponseHistory {
  s.DeploymentDate = &v
  return s
}

func (s *GetHistoricalInformationAboutOneHostnameResponseHistory) SetUndeploymentDate(v string) *GetHistoricalInformationAboutOneHostnameResponseHistory {
  s.UndeploymentDate = &v
  return s
}




type GetHistoricalInformationAboutHostnamesPaths struct {
}

func (s GetHistoricalInformationAboutHostnamesPaths) String() string {
  return tea.Prettify(s)
}

func (s GetHistoricalInformationAboutHostnamesPaths) GoString() string {
  return s.String()
}

type GetHistoricalInformationAboutHostnamesParameters struct {
  // {"en" : "Beginning of the time period in RFC 3339 format. The time must be specified using the UTC timezone; it cannot be an offset. Example: startdate=2019-11-14T00:00:00Z By default, the value is when you began using CDN Pro.", "zh_CN": "指定查询开始时间，以RFC 3339日期格式表示。必须使用UTC时区，不支持指定其它时区。示例：startdate=2019-11-14T00:00:00Z。如果开始时间未指定，则默认为您开通CDN Pro服务的时间。"}
  Startdate *string `json:"startdate,omitempty" xml:"startdate,omitempty"`
  // {"en" : "End of the time period in RFC 3339 format. The time must be specified using the UTC timezone; it cannot be an offset. Example: enddate=2019-11-14T00:00:00Z The default is the current time.", "zh_CN": "指定查询结束时间，以RFC 3339日期格式表示。必须使用UTC时区，不支持指定其它时区。示例：enddate=2019-12-14T00:00:00Z。如果结束时间未指定，默认为当前时间。"}
  Enddate *string `json:"enddate,omitempty" xml:"enddate,omitempty"`
  // {"en" : "Filter by looking up the beginning portion of a hostname.", "zh_CN": "通过搜索关键字匹配加速域名进行过滤。"}
  Search *string `json:"search,omitempty" xml:"search,omitempty"`
  // {"en" : "Enum: all,staging,production 
  // Default: all 
  // The value can be 'all', 'staging', or 'production' to filter the results based on where the hostnames have been deployed. 
  // ", "zh_CN": "取值范围: all,staging,production 
  // 默认值: all 
  // 根据加速域名的部署环境过滤。该值可以是'all', 'staging', 或'production'，分别表示所有环境，演练环境和生产环境。"}
  Target *string `json:"target,omitempty" xml:"target,omitempty"`
  // {"en" : "Default: 0 Range: >= 0 
  // Index of the first hostname to return.", "zh_CN": "默认值: 0 取值范围: >= 0 
  // 查询起始位置。"}
  Offset *int `json:"offset,omitempty" xml:"offset,omitempty"`
  // {"en" : "Default: 100 Range: [ 1 .. 10000 ] 
  // Maximum number of hostnames to return.", "zh_CN": "默认值: 100 取值范围: <= 10000 
  // 每次查询的最大条数。"}
  Limit *int `json:"limit,omitempty" xml:"limit,omitempty"`
}

func (s GetHistoricalInformationAboutHostnamesParameters) String() string {
  return tea.Prettify(s)
}

func (s GetHistoricalInformationAboutHostnamesParameters) GoString() string {
  return s.String()
}

func (s *GetHistoricalInformationAboutHostnamesParameters) SetStartdate(v string) *GetHistoricalInformationAboutHostnamesParameters {
  s.Startdate = &v
  return s
}

func (s *GetHistoricalInformationAboutHostnamesParameters) SetEnddate(v string) *GetHistoricalInformationAboutHostnamesParameters {
  s.Enddate = &v
  return s
}

func (s *GetHistoricalInformationAboutHostnamesParameters) SetSearch(v string) *GetHistoricalInformationAboutHostnamesParameters {
  s.Search = &v
  return s
}

func (s *GetHistoricalInformationAboutHostnamesParameters) SetTarget(v string) *GetHistoricalInformationAboutHostnamesParameters {
  s.Target = &v
  return s
}

func (s *GetHistoricalInformationAboutHostnamesParameters) SetOffset(v int) *GetHistoricalInformationAboutHostnamesParameters {
  s.Offset = &v
  return s
}

func (s *GetHistoricalInformationAboutHostnamesParameters) SetLimit(v int) *GetHistoricalInformationAboutHostnamesParameters {
  s.Limit = &v
  return s
}

type GetHistoricalInformationAboutHostnamesRequestHeader struct {
}

func (s GetHistoricalInformationAboutHostnamesRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetHistoricalInformationAboutHostnamesRequestHeader) GoString() string {
  return s.String()
}

type GetHistoricalInformationAboutHostnamesRequest struct {
}

func (s GetHistoricalInformationAboutHostnamesRequest) String() string {
  return tea.Prettify(s)
}

func (s GetHistoricalInformationAboutHostnamesRequest) GoString() string {
  return s.String()
}

type GetHistoricalInformationAboutHostnamesResponseHeader struct {
}

func (s GetHistoricalInformationAboutHostnamesResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetHistoricalInformationAboutHostnamesResponseHeader) GoString() string {
  return s.String()
}

type GetHistoricalInformationAboutHostnamesResponse struct {
  // {"en" : "Range: >= 0 
  // Total number of historical hostnames. This may be larger than number actually returned.", "zh_CN": "取值范围: >= 0 
  // 历史加速域名的总数。"}
  Count *int `json:"count,omitempty" xml:"count,omitempty" require:"true"`
  // {"en" : "List of hostnames that were deployed during the timeframe.", "zh_CN": "在指定时间段内部署到生产环境过的加速域名列表。"}
  Hostnames []*string `json:"hostnames,omitempty" xml:"hostnames,omitempty" require:"true" type:"Repeated"`
}

func (s GetHistoricalInformationAboutHostnamesResponse) String() string {
  return tea.Prettify(s)
}

func (s GetHistoricalInformationAboutHostnamesResponse) GoString() string {
  return s.String()
}

func (s *GetHistoricalInformationAboutHostnamesResponse) SetCount(v int) *GetHistoricalInformationAboutHostnamesResponse {
  s.Count = &v
  return s
}

func (s *GetHistoricalInformationAboutHostnamesResponse) SetHostnames(v []*string) *GetHistoricalInformationAboutHostnamesResponse {
  s.Hostnames = v
  return s
}




type GetInformationAboutASpecificHostnamePaths struct {
  // {"en" : "Hostname to fetch information about.", "zh_CN": "加速域名。"}
  Hostname *string `json:"hostname,omitempty" xml:"hostname,omitempty" require:"true"`
}

func (s GetInformationAboutASpecificHostnamePaths) String() string {
  return tea.Prettify(s)
}

func (s GetInformationAboutASpecificHostnamePaths) GoString() string {
  return s.String()
}

func (s *GetInformationAboutASpecificHostnamePaths) SetHostname(v string) *GetInformationAboutASpecificHostnamePaths {
  s.Hostname = &v
  return s
}

type GetInformationAboutASpecificHostnameParameters struct {
}

func (s GetInformationAboutASpecificHostnameParameters) String() string {
  return tea.Prettify(s)
}

func (s GetInformationAboutASpecificHostnameParameters) GoString() string {
  return s.String()
}

type GetInformationAboutASpecificHostnameRequestHeader struct {
}

func (s GetInformationAboutASpecificHostnameRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetInformationAboutASpecificHostnameRequestHeader) GoString() string {
  return s.String()
}

type GetInformationAboutASpecificHostnameRequest struct {
}

func (s GetInformationAboutASpecificHostnameRequest) String() string {
  return tea.Prettify(s)
}

func (s GetInformationAboutASpecificHostnameRequest) GoString() string {
  return s.String()
}

type GetInformationAboutASpecificHostnameResponseHeader struct {
}

func (s GetInformationAboutASpecificHostnameResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetInformationAboutASpecificHostnameResponseHeader) GoString() string {
  return s.String()
}

type GetInformationAboutASpecificHostnameResponse struct {
  // {"en" : "A hostname whose property has been deployed to production or staging.", "zh_CN": "加速域名。"}
  Hostname *string `json:"hostname,omitempty" xml:"hostname,omitempty" require:"true"`
  // {"en" : "Describes the property deployed to production. It can be 'null' if the property is not deployed to production.", "zh_CN": "如果加速域名对应的加速项目已部署到生产环境，则该字段返回该加速项目的相关信息。如果加速项目未部署到生产环境，则该字段返回'null'。"}
  PropertyInProduction *string `json:"propertyInProduction,omitempty" xml:"propertyInProduction,omitempty" require:"true"`
  // {"en" : "Describes the property deployed to staging. It can be 'null' if the property is not deployed to staging.", "zh_CN": "如果加速域名对应的加速项目已部署到演练环境，则该字段返回该加速项目的相关信息。如果加速项目未部署到演练环境，则该字段返回'null'。"}
  PropertyInStaging *string `json:"propertyInStaging,omitempty" xml:"propertyInStaging,omitempty" require:"true"`
}

func (s GetInformationAboutASpecificHostnameResponse) String() string {
  return tea.Prettify(s)
}

func (s GetInformationAboutASpecificHostnameResponse) GoString() string {
  return s.String()
}

func (s *GetInformationAboutASpecificHostnameResponse) SetHostname(v string) *GetInformationAboutASpecificHostnameResponse {
  s.Hostname = &v
  return s
}

func (s *GetInformationAboutASpecificHostnameResponse) SetPropertyInProduction(v string) *GetInformationAboutASpecificHostnameResponse {
  s.PropertyInProduction = &v
  return s
}

func (s *GetInformationAboutASpecificHostnameResponse) SetPropertyInStaging(v string) *GetInformationAboutASpecificHostnameResponse {
  s.PropertyInStaging = &v
  return s
}




