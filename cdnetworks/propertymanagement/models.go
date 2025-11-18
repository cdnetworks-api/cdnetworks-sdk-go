package propertymanagement

// This file is auto-generated, don't edit it. Thanks.
import (
  "github.com/alibabacloud-go/tea/tea"
)

type GetListOfPropertyVersionsPaths struct {
  // {"en" : "ID of a property", "zh_CN": "加速项目ID。"}
  PropertyId *string `json:"propertyId,omitempty" xml:"propertyId,omitempty" require:"true"`
}

func (s GetListOfPropertyVersionsPaths) String() string {
  return tea.Prettify(s)
}

func (s GetListOfPropertyVersionsPaths) GoString() string {
  return s.String()
}

func (s *GetListOfPropertyVersionsPaths) SetPropertyId(v string) *GetListOfPropertyVersionsPaths {
  s.PropertyId = &v
  return s
}

type GetListOfPropertyVersionsParameters struct {
  // {"en" : "The search parameter is used to match against the description field to filter the results.", "zh_CN": "根据搜索关键字匹配加速项目版本的描述字段来过滤结果。"}
  Search *string `json:"search,omitempty" xml:"search,omitempty"`
  // {"en" : "Default: 100 Range: [ 1 .. 200 ] 
  // Controls the number of versions to return.", "zh_CN": "默认值: 100 取值范围: <= 200 
  // 每次查询的最大条数。"}
  Limit *int `json:"limit,omitempty" xml:"limit,omitempty"`
  // {"en" : "Default: 0 Range: >= 0 
  // Returns property versions beginning with this index instead of the default first one.", "zh_CN": "默认值: 0 取值范围: >= 0 
  // 查询起始位置。"}
  Offset *int `json:"offset,omitempty" xml:"offset,omitempty"`
  // {"en" : "Enum: asc,desc 
  // Default: asc 
  // Order of property versions to return.", "zh_CN": "取值范围: asc,desc 
  // 默认值: asc 
  // 返回结果的顺序。默认按版本号升序排列。"}
  SortOrder *string `json:"sortOrder,omitempty" xml:"sortOrder,omitempty"`
  // {"en" : "Enum: version,lastUpdateTime 
  // Default: version 
  // Returns results in sorted order.", "zh_CN": "取值范围: version,lastUpdateTime 
  // 默认值: version 
  // 返回结果的排序依据。支持按版本号和版本最后更新时间进行排序。"}
  SortBy *string `json:"sortBy,omitempty" xml:"sortBy,omitempty"`
}

func (s GetListOfPropertyVersionsParameters) String() string {
  return tea.Prettify(s)
}

func (s GetListOfPropertyVersionsParameters) GoString() string {
  return s.String()
}

func (s *GetListOfPropertyVersionsParameters) SetSearch(v string) *GetListOfPropertyVersionsParameters {
  s.Search = &v
  return s
}

func (s *GetListOfPropertyVersionsParameters) SetLimit(v int) *GetListOfPropertyVersionsParameters {
  s.Limit = &v
  return s
}

func (s *GetListOfPropertyVersionsParameters) SetOffset(v int) *GetListOfPropertyVersionsParameters {
  s.Offset = &v
  return s
}

func (s *GetListOfPropertyVersionsParameters) SetSortOrder(v string) *GetListOfPropertyVersionsParameters {
  s.SortOrder = &v
  return s
}

func (s *GetListOfPropertyVersionsParameters) SetSortBy(v string) *GetListOfPropertyVersionsParameters {
  s.SortBy = &v
  return s
}

type GetListOfPropertyVersionsRequestHeader struct {
}

func (s GetListOfPropertyVersionsRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetListOfPropertyVersionsRequestHeader) GoString() string {
  return s.String()
}

type GetListOfPropertyVersionsRequest struct {
}

func (s GetListOfPropertyVersionsRequest) String() string {
  return tea.Prettify(s)
}

func (s GetListOfPropertyVersionsRequest) GoString() string {
  return s.String()
}

type GetListOfPropertyVersionsResponseHeader struct {
}

func (s GetListOfPropertyVersionsResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetListOfPropertyVersionsResponseHeader) GoString() string {
  return s.String()
}

type GetListOfPropertyVersionsResponse struct {
  // {"en" : "A summary of each version of the property.", "zh_CN": "加速项目每个版本的摘要。"}
  PropertyVersions []*GetListOfPropertyVersionsResponsePropertyVersions `json:"propertyVersions,omitempty" xml:"propertyVersions,omitempty" require:"true" type:"Repeated"`
  // {"en" : "Range: >= 0 
  // Indicates the number of versions of the property.", "zh_CN": "取值范围: >= 0 
  // 加速项目的版本数量。"}
  Count *int `json:"count,omitempty" xml:"count,omitempty" require:"true"`
}

func (s GetListOfPropertyVersionsResponse) String() string {
  return tea.Prettify(s)
}

func (s GetListOfPropertyVersionsResponse) GoString() string {
  return s.String()
}

func (s *GetListOfPropertyVersionsResponse) SetPropertyVersions(v []*GetListOfPropertyVersionsResponsePropertyVersions) *GetListOfPropertyVersionsResponse {
  s.PropertyVersions = v
  return s
}

func (s *GetListOfPropertyVersionsResponse) SetCount(v int) *GetListOfPropertyVersionsResponse {
  s.Count = &v
  return s
}

type GetListOfPropertyVersionsResponsePropertyVersions struct     {
  // {"en" : "Range: >= 1 
  // A version number.", "zh_CN": "取值范围: >= 1 
  // 版本号。"}
  Version *int `json:"version,omitempty" xml:"version,omitempty"`
  // {"en" : "A description of the property version.", "zh_CN": "加速项目版本描述。"}
  Description *string `json:"description,omitempty" xml:"description,omitempty"`
  // {"en" : "RFC 3339 format date indicating when the property version was created. Example: 2016-12-28T23:10:42Z", "zh_CN": "RFC 3339格式的日期，表示创建加速项目版本的时间。示例：2016-12-28T23:10:42Z"}
  CreationTime *string `json:"creationTime,omitempty" xml:"creationTime,omitempty"`
  // {"en" : "Hostnames associated with the property.", "zh_CN": "加速项目下的加速域名。"}
  Hostnames []*string `json:"hostnames,omitempty" xml:"hostnames,omitempty" type:"Repeated"`
  // {"en" : "Default: False 
  // Indicates whether the property version is frozen or can still be updated. A property version is frozen once it has been deployed.", "zh_CN": "默认值: False 
  // 该加速项目版本是否处于冻结状态。当加速项目版本部署后即进入冻结状态，不可再更新该版本。"}
  Frozen *bool `json:"frozen,omitempty" xml:"frozen,omitempty"`
  // {"en" : "RFC 3339 format date indicating when the property version was last updated. Example: 2016-12-28T23:10:42Z", "zh_CN": "RFC 3339格式的日期，表示加速项目版本的最近更新时间。示例：2016-12-28T23:10:42Z"}
  LastUpdateTime *string `json:"lastUpdateTime,omitempty" xml:"lastUpdateTime,omitempty"`
  // {"en" : "Enum: NotValidated,InProgress,Validated,Failed 
  // Indicates the status of the most recent validation of the property version.", "zh_CN": "取值范围: NotValidated,InProgress,Validated,Failed 
  // 加速项目版本最近一次验证的状态，包括未验证，验证中，已验证，验证失败等状态。"}
  LastValidationStatus *string `json:"lastValidationStatus,omitempty" xml:"lastValidationStatus,omitempty"`
}

func (s GetListOfPropertyVersionsResponsePropertyVersions) String() string {
  return tea.Prettify(s)
}

func (s GetListOfPropertyVersionsResponsePropertyVersions) GoString() string {
  return s.String()
}

func (s *GetListOfPropertyVersionsResponsePropertyVersions) SetVersion(v int) *GetListOfPropertyVersionsResponsePropertyVersions {
  s.Version = &v
  return s
}

func (s *GetListOfPropertyVersionsResponsePropertyVersions) SetDescription(v string) *GetListOfPropertyVersionsResponsePropertyVersions {
  s.Description = &v
  return s
}

func (s *GetListOfPropertyVersionsResponsePropertyVersions) SetCreationTime(v string) *GetListOfPropertyVersionsResponsePropertyVersions {
  s.CreationTime = &v
  return s
}

func (s *GetListOfPropertyVersionsResponsePropertyVersions) SetHostnames(v []*string) *GetListOfPropertyVersionsResponsePropertyVersions {
  s.Hostnames = v
  return s
}

func (s *GetListOfPropertyVersionsResponsePropertyVersions) SetFrozen(v bool) *GetListOfPropertyVersionsResponsePropertyVersions {
  s.Frozen = &v
  return s
}

func (s *GetListOfPropertyVersionsResponsePropertyVersions) SetLastUpdateTime(v string) *GetListOfPropertyVersionsResponsePropertyVersions {
  s.LastUpdateTime = &v
  return s
}

func (s *GetListOfPropertyVersionsResponsePropertyVersions) SetLastValidationStatus(v string) *GetListOfPropertyVersionsResponsePropertyVersions {
  s.LastValidationStatus = &v
  return s
}




type DeleteAPropertyPaths struct {
  // {"en" : "ID of a property
  // ", "zh_CN": ""}
  PropertyId *string `json:"propertyId,omitempty" xml:"propertyId,omitempty" require:"true"`
}

func (s DeleteAPropertyPaths) String() string {
  return tea.Prettify(s)
}

func (s DeleteAPropertyPaths) GoString() string {
  return s.String()
}

func (s *DeleteAPropertyPaths) SetPropertyId(v string) *DeleteAPropertyPaths {
  s.PropertyId = &v
  return s
}

type DeleteAPropertyParameters struct {
}

func (s DeleteAPropertyParameters) String() string {
  return tea.Prettify(s)
}

func (s DeleteAPropertyParameters) GoString() string {
  return s.String()
}

type DeleteAPropertyRequestHeader struct {
}

func (s DeleteAPropertyRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteAPropertyRequestHeader) GoString() string {
  return s.String()
}

type DeleteAPropertyRequest struct {
}

func (s DeleteAPropertyRequest) String() string {
  return tea.Prettify(s)
}

func (s DeleteAPropertyRequest) GoString() string {
  return s.String()
}

type DeleteAPropertyResponseHeader struct {
}

func (s DeleteAPropertyResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteAPropertyResponseHeader) GoString() string {
  return s.String()
}

type DeleteAPropertyResponse struct {
}

func (s DeleteAPropertyResponse) String() string {
  return tea.Prettify(s)
}

func (s DeleteAPropertyResponse) GoString() string {
  return s.String()
}




type DeleteAPropertyVersionPaths struct {
  // {"en" : "ID of a property.", "zh_CN": "加速项目ID。"}
  PropertyId *string `json:"propertyId,omitempty" xml:"propertyId,omitempty" require:"true"`
  // {"en" : "A property version. It must be an integer value >=1. The Get a property version API also permits you to specify the word 'latest' to obtain the most recent version.", "zh_CN": "加速项目版本。必须是大于等于1的整数值。"}
  Version *string `json:"version,omitempty" xml:"version,omitempty" require:"true"`
}

func (s DeleteAPropertyVersionPaths) String() string {
  return tea.Prettify(s)
}

func (s DeleteAPropertyVersionPaths) GoString() string {
  return s.String()
}

func (s *DeleteAPropertyVersionPaths) SetPropertyId(v string) *DeleteAPropertyVersionPaths {
  s.PropertyId = &v
  return s
}

func (s *DeleteAPropertyVersionPaths) SetVersion(v string) *DeleteAPropertyVersionPaths {
  s.Version = &v
  return s
}

type DeleteAPropertyVersionParameters struct {
}

func (s DeleteAPropertyVersionParameters) String() string {
  return tea.Prettify(s)
}

func (s DeleteAPropertyVersionParameters) GoString() string {
  return s.String()
}

type DeleteAPropertyVersionRequestHeader struct {
}

func (s DeleteAPropertyVersionRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteAPropertyVersionRequestHeader) GoString() string {
  return s.String()
}

type DeleteAPropertyVersionRequest struct {
}

func (s DeleteAPropertyVersionRequest) String() string {
  return tea.Prettify(s)
}

func (s DeleteAPropertyVersionRequest) GoString() string {
  return s.String()
}

type DeleteAPropertyVersionResponseHeader struct {
}

func (s DeleteAPropertyVersionResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteAPropertyVersionResponseHeader) GoString() string {
  return s.String()
}

type DeleteAPropertyVersionResponse struct {
}

func (s DeleteAPropertyVersionResponse) String() string {
  return tea.Prettify(s)
}

func (s DeleteAPropertyVersionResponse) GoString() string {
  return s.String()
}




type GetAPropertyRequest struct {
}

func (s GetAPropertyRequest) String() string {
  return tea.Prettify(s)
}

func (s GetAPropertyRequest) GoString() string {
  return s.String()
}

type GetAPropertyRequestHeader struct {
}

func (s GetAPropertyRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetAPropertyRequestHeader) GoString() string {
  return s.String()
}

type GetAPropertyPaths struct {
  // {"en":"ID of a property","zh_CN":""}
  PropertyId *string `json:"propertyId,omitempty" xml:"propertyId,omitempty" require:"true"`
}

func (s GetAPropertyPaths) String() string {
  return tea.Prettify(s)
}

func (s GetAPropertyPaths) GoString() string {
  return s.String()
}

func (s *GetAPropertyPaths) SetPropertyId(v string) *GetAPropertyPaths {
  s.PropertyId = &v
  return s
}

type GetAPropertyParameters struct {
}

func (s GetAPropertyParameters) String() string {
  return tea.Prettify(s)
}

func (s GetAPropertyParameters) GoString() string {
  return s.String()
}

type GetAPropertyResponse struct {
  // {"en":"Name of the property.","zh_CN":"加速项目的名称。"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"A description of the property.","zh_CN":"加速项目的描述。"}
  Description *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
  // {"en":"Indicates the version of the property deployed to production. The field will be omitted if the property was not yet deployed.","zh_CN":"部署到生产环境的加速项目版本。如果尚未部署加速项目，则将省略该字段。"}
  VersionInProduction *int `json:"versionInProduction,omitempty" xml:"versionInProduction,omitempty" require:"true"`
  // {"en":"Indicates the version of the property deployed to staging. The field will be omitted if the property was not yet deployed.","zh_CN":"部署到演练环境的加速项目版本。如果尚未部署加速项目，则将省略该字段。"}
  VersionInStaging *int `json:"versionInStaging,omitempty" xml:"versionInStaging,omitempty" require:"true"`
  // {"en":"An RFC 3339 date indicating when the property was created.","zh_CN":"RFC 3339格式的日期，表示创建加速项目的时间。"}
  CreationTime *string `json:"creationTime,omitempty" xml:"creationTime,omitempty" require:"true"`
  // {"en":"An RFC 3339 date indicating when the property was last updated.","zh_CN":"RFC 3339格式的日期，表示加速项目的最近更新时间。"}
  LastUpdateTime *string `json:"lastUpdateTime,omitempty" xml:"lastUpdateTime,omitempty" require:"true"`
  // {"en":"Number of versions of the property.","zh_CN":"加速项目的版本数量。"}
  Versions *int `json:"versions,omitempty" xml:"versions,omitempty" require:"true"`
  // {"en":"Enum: wsapro,webpro,vodpro,downloadpro\nService type, one of wsapro, webpro, vodpro , downloadpro , 1523.","zh_CN":"取值范围: wsapro,webpro,vodpro,downloadpro,MALL-cdnpro\n加速域名关联的商品（服务类型），即全站加速，网页加速，点播加速，下载加速，CDN Pro，以及自助CDN。"}
  LegacyType *string `json:"legacyType,omitempty" xml:"legacyType,omitempty" require:"true"`
}

func (s GetAPropertyResponse) String() string {
  return tea.Prettify(s)
}

func (s GetAPropertyResponse) GoString() string {
  return s.String()
}

func (s *GetAPropertyResponse) SetName(v string) *GetAPropertyResponse {
  s.Name = &v
  return s
}

func (s *GetAPropertyResponse) SetDescription(v string) *GetAPropertyResponse {
  s.Description = &v
  return s
}

func (s *GetAPropertyResponse) SetVersionInProduction(v int) *GetAPropertyResponse {
  s.VersionInProduction = &v
  return s
}

func (s *GetAPropertyResponse) SetVersionInStaging(v int) *GetAPropertyResponse {
  s.VersionInStaging = &v
  return s
}

func (s *GetAPropertyResponse) SetCreationTime(v string) *GetAPropertyResponse {
  s.CreationTime = &v
  return s
}

func (s *GetAPropertyResponse) SetLastUpdateTime(v string) *GetAPropertyResponse {
  s.LastUpdateTime = &v
  return s
}

func (s *GetAPropertyResponse) SetVersions(v int) *GetAPropertyResponse {
  s.Versions = &v
  return s
}

func (s *GetAPropertyResponse) SetLegacyType(v string) *GetAPropertyResponse {
  s.LegacyType = &v
  return s
}

type GetAPropertyResponseHeader struct {
}

func (s GetAPropertyResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetAPropertyResponseHeader) GoString() string {
  return s.String()
}




type UpdateAPropertyPaths struct {
  // {"en" : "ID of a property
  // ", "zh_CN": "加速项目ID。"}
  PropertyId *string `json:"propertyId,omitempty" xml:"propertyId,omitempty" require:"true"`
}

func (s UpdateAPropertyPaths) String() string {
  return tea.Prettify(s)
}

func (s UpdateAPropertyPaths) GoString() string {
  return s.String()
}

func (s *UpdateAPropertyPaths) SetPropertyId(v string) *UpdateAPropertyPaths {
  s.PropertyId = &v
  return s
}

type UpdateAPropertyParameters struct {
}

func (s UpdateAPropertyParameters) String() string {
  return tea.Prettify(s)
}

func (s UpdateAPropertyParameters) GoString() string {
  return s.String()
}

type UpdateAPropertyRequestHeader struct {
}

func (s UpdateAPropertyRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s UpdateAPropertyRequestHeader) GoString() string {
  return s.String()
}

type UpdateAPropertyRequest struct {
  // {"en" : "Name of the property.", "zh_CN": "加速项目的名称。"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en" : "Description of the property. If unspecified, the description will not be updated.", "zh_CN": "加速项目的描述。"}
  Description *string `json:"description,omitempty" xml:"description,omitempty"`
}

func (s UpdateAPropertyRequest) String() string {
  return tea.Prettify(s)
}

func (s UpdateAPropertyRequest) GoString() string {
  return s.String()
}

func (s *UpdateAPropertyRequest) SetName(v string) *UpdateAPropertyRequest {
  s.Name = &v
  return s
}

func (s *UpdateAPropertyRequest) SetDescription(v string) *UpdateAPropertyRequest {
  s.Description = &v
  return s
}

type UpdateAPropertyResponseHeader struct {
}

func (s UpdateAPropertyResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s UpdateAPropertyResponseHeader) GoString() string {
  return s.String()
}

type UpdateAPropertyResponse struct {
}

func (s UpdateAPropertyResponse) String() string {
  return tea.Prettify(s)
}

func (s UpdateAPropertyResponse) GoString() string {
  return s.String()
}




type GetAPropertyVersionRequest struct {
}

func (s GetAPropertyVersionRequest) String() string {
  return tea.Prettify(s)
}

func (s GetAPropertyVersionRequest) GoString() string {
  return s.String()
}

type GetAPropertyVersionRequestHeader struct {
}

func (s GetAPropertyVersionRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetAPropertyVersionRequestHeader) GoString() string {
  return s.String()
}

type GetAPropertyVersionPaths struct {
  // {"en":"ID of a property.","zh_CN":"加速项目ID。"}
  PropertyId *string `json:"propertyId,omitempty" xml:"propertyId,omitempty" require:"true"`
  // {"en":"A property version. It must be an integer value >=1. The Get a property version API also permits you to specify the word 'latest' to obtain the most recent version.","zh_CN":"加速项目版本。必须是大于等于1的整数值，但可以指定'latest'来查询最新版本。"}
  Version *string `json:"version,omitempty" xml:"version,omitempty" require:"true"`
}

func (s GetAPropertyVersionPaths) String() string {
  return tea.Prettify(s)
}

func (s GetAPropertyVersionPaths) GoString() string {
  return s.String()
}

func (s *GetAPropertyVersionPaths) SetPropertyId(v string) *GetAPropertyVersionPaths {
  s.PropertyId = &v
  return s
}

func (s *GetAPropertyVersionPaths) SetVersion(v string) *GetAPropertyVersionPaths {
  s.Version = &v
  return s
}

type GetAPropertyVersionParameters struct {
}

func (s GetAPropertyVersionParameters) String() string {
  return tea.Prettify(s)
}

func (s GetAPropertyVersionParameters) GoString() string {
  return s.String()
}

type GetAPropertyVersionResponse struct {
  // {"en":"Version number.","zh_CN":"版本号。"}
  Version *int `json:"version,omitempty" xml:"version,omitempty" require:"true"`
  // {"en":"Describes a property configuration. This contains all the settings.","zh_CN":"描述加速项目的所有配置信息。"}
  Configs *GetAPropertyVersionResponseConfigs `json:"configs,omitempty" xml:"configs,omitempty" require:"true" type:"Struct"`
  // {"en":"Status of the property version.","zh_CN":"状态信息。"}
  Status *GetAPropertyVersionResponseStatus `json:"status,omitempty" xml:"status,omitempty" require:"true" type:"Struct"`
}

func (s GetAPropertyVersionResponse) String() string {
  return tea.Prettify(s)
}

func (s GetAPropertyVersionResponse) GoString() string {
  return s.String()
}

func (s *GetAPropertyVersionResponse) SetVersion(v int) *GetAPropertyVersionResponse {
  s.Version = &v
  return s
}

func (s *GetAPropertyVersionResponse) SetConfigs(v *GetAPropertyVersionResponseConfigs) *GetAPropertyVersionResponse {
  s.Configs = v
  return s
}

func (s *GetAPropertyVersionResponse) SetStatus(v *GetAPropertyVersionResponseStatus) *GetAPropertyVersionResponse {
  s.Status = v
  return s
}

type GetAPropertyVersionResponseConfigs struct {
  // {"en":"A description of the version.","zh_CN":"版本描述。"}
  Description *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
  // {"en":"Hostnames associated with the property. A valid value is a fully qualified hostname such as www.domain.com or a wildcard hostname such as.domain.com. Any given hostname can only be in one deployed property at a particular time. However, a wildcard hostname is permitted to overlap other hostnames you own.","zh_CN":"与加速项目关联的加速域名。必须是FQDN完全限定域名（如 www.domain.com），或泛域名（如*.domain.com）。\n同一个加速域名在同一时间只能存在于一个已部署的加速项目中，但泛域名可以与关联的完全限定域名一同部署。"}
  Hostnames []*string `json:"hostnames,omitempty" xml:"hostnames,omitempty" require:"true" type:"Repeated"`
  // {"en":"Default: False\nThe value indicates whether all the hostnames in this property have Beian license on file with the Chinese government. This is required to serve this property from servers in mainland China. A value of false means servers outside of mainland China will be used to distribute content to visitors in China. If set to true you must also specify the content type in the beianContentType field.","zh_CN":"默认值: False\n此加速项目中的所有加速域名是否都已获得ICP备案。只有获取备案的域名才能使用中国大陆节点来提供服务。如果设置为true，还必须在beianContentType字段中指定内容类型。"}
  HasBeian *bool `json:"hasBeian,omitempty" xml:"hasBeian,omitempty" require:"true"`
  // {"en":"Range: [ 1 .. 24 ]\nIf you are planning to distribute content through servers in mainland China, you will need to set the hasBeian field to true and also indicate the type of content you are distributing. Enter the value that best corresponds to your content:\n1. Instant Messaging\n2. Search Engine\n3. Web Portal\n4. Online Postal Service\n5. Online News\n6. Blog\n7. Advertisement\n8. Organization Web Portal\n9. Shopping\n10. Online Payment\n11. Online Bank\n12. Online Stock Market\n13. Online Gaming\n14. Online Music\n15. Online Movie\n16. Online Picture\n17. Software Download\n18. Job Hunting\n19. Online Dating\n20. Online Real Property\n21. Online Education\n22. Web Design\n23. WAP\n24. Others","zh_CN":"取值范围: [ 1 .. 24 ]\n如果您的域名已获得备案，希望通过中国大陆的服务器分发内容，则需要将hasBeian字段设置为true，并说明您分发的内容类型。请选择与您的内容最匹配的值：\n1. 即时通信\n2. 搜索引擎\n3. 综合门户\n4. 网上邮局\n5. 网络新闻\n6. 博客/个人空间\n7. 网络广告\n8. 单位门户网站\n9. 网络购物\n10. 网上支付\n11. 网上银行\n12. 网上炒股\n13. 网络游戏\n14. 网络音乐\n15. 网络影视\n16. 网络图片\n17. 软件下载\n18. 网上求职\n19. 在线交友\n20. 网上房产\n21. 网络教育\n22. 网站建设\n23. WAP\n24. 其他"}
  BeianContentType *int `json:"beianContentType,omitempty" xml:"beianContentType,omitempty" require:"true"`
  // {"en":"Describes the origin servers for the property's content.","zh_CN":"描述加速项目对应的源站。"}
  Origins []*GetAPropertyVersionResponseConfigsOrigins `json:"origins,omitempty" xml:"origins,omitempty" require:"true" type:"Repeated"`
  // {"en":"Default: False\nThis setting applies only if the property has an attached certificate allowing client requests to use HTTPS. If the value of allowProtocolDowngrade is false, we require all the origin servers to support HTTPS. If the value is true, we allow origins that support only HTTP, which reduces security.","zh_CN":"默认值: False\n是否允许回源时从HTTPS降级为HTTP。仅当加速项目中指定了用于HTTPS请求的证书时，此设置才适用。如果allowProtocolDowngrade的值为false，且加速项目开启了HTTPS，则要求所有源服务器支持HTTPS。如果值为true，则允许仅支持HTTP的源，但这会降低安全性。"}
  AllowProtocolDowngrade *bool `json:"allowProtocolDowngrade,omitempty" xml:"allowProtocolDowngrade,omitempty" require:"true"`
  // {"en":"Default: False\nEnables or disables using IPv6 addresses as origin.","zh_CN":"默认值: False\n是否开启IPv6回源"}
  EnableIpv6Origin *bool `json:"enableIpv6Origin,omitempty" xml:"enableIpv6Origin,omitempty" require:"true"`
  // {"en":"Enum: auto, strict, off\nDefault: off\n\nWhen enableIpv6Origin is allowed, this setting will control whether to follow the client IP protocol version for back-to-origin requests.\nAuto: Attempt to reach the origin server using the same IP addressing as the client, but use the other version as backup if needed.\nStrict: Use only the same addressing method as the client when accessing the origin for content.\nOff: Use both IPv4 and IPv6 addressing to reach the origin, regardless of the IP addressing method used by the client.","zh_CN":"取值范围：auto, strict, off\n默认值：off\n当enableIpv6Origin为true时，该设置将控制是否跟随客户端IP协议版本回源。\nauto：表示根据客户端请求的IP协议版本来选择回源IP地址，当源站服务器不可用时自动切换成其他IP协议地址回源\nstrict：严格根据客户端请求的IP协议版本来选择回源IP地址，当源站服务器不可用时不允许切换成其他协议地址回源\noff：忽略客户端请求的IP协议版本，随机选择可用的IPv4或IPv6地址回源"}
  FollowClientIpVersion *string `json:"followClientIpVersion,omitempty" xml:"followClientIpVersion,omitempty" require:"true"`
  // {"en":"Range: <= 65530 characters\nRefer to Edge Logic Introduction.","zh_CN":"取值范围: <= 65530 字符\n自定义边缘逻辑。参考边缘逻辑介绍。"}
  EdgeLogic *string `json:"edgeLogic,omitempty" xml:"edgeLogic,omitempty" require:"true"`
  // {"en":"Default: 1\nThe value must be set to 1 at this time.","zh_CN":"默认值: 1\n当前仅允许值为1。"}
  SyntaxVersion *int `json:"syntaxVersion,omitempty" xml:"syntaxVersion,omitempty" require:"true"`
  // {"en":"Range: <= 80 characters\nThe cachekey hostname must be a string made of lowercase letters, digits, dot, dash, and underscore. If not specified, the incoming Host header value will be used and the cache will not be shared across different hostnames in this property.","zh_CN":"取值范围: <= 80\n用于缓存键（Cache Key）的主机名。必须是由小写字母、数字、点、连字符和下划线组成的字符串。如果未指定，默认将使用传入的Host请求头值，并且此加速项目中的不同加速域名之间不共享缓存。"}
  CacheKeyHostname *string `json:"cacheKeyHostname,omitempty" xml:"cacheKeyHostname,omitempty" require:"true"`
  // {"en":"Enum: preRewrite,postRewrite\nDefault: preRewrite\nThis setting controls how the URI of the incoming request is incorporated into the cache key if you use a 'rewrite' directive in the property's Edge Logic. The default value, 'preRewrite', puts the unmodified URI into the cache key while 'postRewrite' uses the modified URI. If your rewrite directive converts multiple different URIs into the same value, using 'postRewrite' may result in a higher cache hit ratio.","zh_CN":"取值范围: preRewrite,postRewrite\n默认值: preRewrite\n如果在加速项目的边缘逻辑中使用了'rewrite'指令，可使用该字段来控制客户端请求的URI如何合并到缓存键（Cache Key）中 。默认值'preRewrite'指将改写前的URI放入缓存键，而'postRewrite'则使用改写后的URI。如果您的'rewrite'指令将多个不同的URI改写为相同的值，则使用'postRewrite'可以提高缓存命中率。"}
  CacheKeyUri *string `json:"cacheKeyUri,omitempty" xml:"cacheKeyUri,omitempty" require:"true"`
  // {"en":"Default: False\nSpecifies whether the scheme ('http' or 'https') should be included in the cache key. Default behavior is false. Set to true if the content is known to be different for different schemes.","zh_CN":"默认值: False\n指定缓存键（Cache Key）是否区分协议（'http'或'https'）。默认为false。如果不同协议响应的内容不同，则设置为true。"}
  SchemeInCacheKey *bool `json:"schemeInCacheKey,omitempty" xml:"schemeInCacheKey,omitempty" require:"true"`
  // {"en":"Default: False\nThis field can be set to either a boolean value or a string. If the value is set to true, the server will redirect all HTTP requests to HTTPS using status code 301. You can also specify string values '302', '307', or '308' instead if you wish to return a different status code when a request is redirected.","zh_CN":"默认值: False\n此字段可以设置为布尔值或字符串。如果设置为true，则CDN Pro服务器会将所有HTTP请求重定向到HTTPS，并返回301状态码。如果您希望在重定向请求时返回不同的状态码，可在此处指定需要的状态码，如'302'、'307'或'308'。"}
  RedirectHttpToHttps *string `json:"redirectHttpToHttps,omitempty" xml:"redirectHttpToHttps,omitempty" require:"true"`
  // {"en":"Default: False\nSet to true to disable support for HTTP/2 and support only HTTP 1.1.","zh_CN":"默认值: False\n当值为true时，禁用对HTTP/2的支持，仅支持HTTP/1.1。"}
  DisableHttp2 *bool `json:"disableHttp2,omitempty" xml:"disableHttp2,omitempty" require:"true"`
  // {"en":"Default: False\nSet to true to enable support for HTTP/3. HTTP/3 requires TLS to work, so this field will be ignored unless you specify a certificate in the tlsCertificateId field.","zh_CN":"默认值: False\n是否开启HTTP/3。需要配置TLS才生效，若您未在tlsCertificateId中指定证书，此字段将被忽略。"}
  EnableHttp3 *bool `json:"enableHttp3,omitempty" xml:"enableHttp3,omitempty" require:"true"`
  // {"en":"This field lists ports other than the default 80 used to handle HTTP requests and ports other than the default 443 used to handle HTTPS requests.","zh_CN":"除标准的80，443端口外，我们还支持一些扩展端口。可用该字段指定用于处理HTTP和HTTPS请求的扩展端口。"}
  ExtraServicePorts *GetAPropertyVersionResponseConfigsExtraServicePorts `json:"extraServicePorts,omitempty" xml:"extraServicePorts,omitempty" require:"true" type:"Struct"`
  // {"en":"Specify a certificate to enable HTTPS for your hostnames.","zh_CN":"指定用于HTTPS请求的证书的ID。如果未指定证书，则不会为此加速项目启用HTTPS。"}
  TlsCertificateId *string `json:"tlsCertificateId,omitempty" xml:"tlsCertificateId,omitempty" require:"true"`
  // {"en":"Specify a second certificate to be used for HTTPS. With tlsCertificateId and tlsCertificateId1, you can specify two certificates of different types, i.e. one RSA and one EC. When there are two certificates available for your hostnames, the client and server will negotiate at the time of requests to pick one of the certificates for SSL. This could mean better performance and fault tolerance.\n\nIt is invalid to only set the tlsCertificateId1 field without setting the tlsCertificateId field. The certificates set for tlsCertificateId and tlsCertificateId1 must be of different type.","zh_CN":"指定用于HTTPS请求的第二本证书。必须先指定tlsCertificateId，才能指定tlsCertificateId1。tlsCertificateId和tlsCertificateId1所指定的证书必须是不同证书类型。\n\ntlsCertificateId和tlsCertificateId1允许您指定两本不同类型的证书，即一本RSA证书，一本EC证书。如果您指定了多本证书，当客户端发起请求时，客户端和服务端可以协商选择使用其中一本证书。"}
  TlsCertificateId1 *int `json:"tlsCertificateId1,omitempty" xml:"tlsCertificateId1,omitempty" require:"true"`
  // {"en":"Enum: 1.1,1.2,1.3\nDefault: 1.3\nMaximum supported TLS version.","zh_CN":"取值范围: 1.1,1.2,1.3\n默认值: 1.3\n支持的TLS最高版本。"}
  TlsMaxVersion *string `json:"tlsMaxVersion,omitempty" xml:"tlsMaxVersion,omitempty" require:"true"`
  // {"en":"Enum: 1.1,1.2,1.3,1\nDefault: 1\nMinimum required TLS version.","zh_CN":"取值范围: 1.1,1.2,1.3,1\n默认值: 1\n所需的TLS最低版本。"}
  TlsMinVersion *string `json:"tlsMinVersion,omitempty" xml:"tlsMinVersion,omitempty" require:"true"`
  // {"en":"Range: <= 2040 characters\nAny cipher list supported by 'openssl ciphers'. See https://www.openssl.org/docs/manmaster/man1/ciphers.html","zh_CN":"取值范围: <= 2040 字符\n'openssl ciphers'支持的任何加密算法套件。参考：https://www.openssl.org/docs/manmaster/man1/ciphers.html"}
  TlsCiphers *string `json:"tlsCiphers,omitempty" xml:"tlsCiphers,omitempty" require:"true"`
  // {"en":"Default: 1800 Range: [ 300 .. 86400 ]\nLifespan of TLS session ticket in seconds.","zh_CN":"默认值: 1800 取值范围: [ 300 .. 86400 ]\nTLS会话ticket的有效期（秒）。"}
  TlsSessionTimeout *int `json:"tlsSessionTimeout,omitempty" xml:"tlsSessionTimeout,omitempty" require:"true"`
  // {"en":"Default: False\nEnable TLS zero round-trip time, a feature of TLS 1.3 that can improve performance for repeat visits to a site. If enabling it though, be sure your site is not vulnerable to the HTTP replay attack.","zh_CN":"默认值: False\n是否开启TLS 0-RTT功能。这是TLS 1.3版本支持的功能。启用该功能后，当用户频繁访问您的站点时，可提高访问响应速度。需要注意的是，启用该功能可能会加大站点遭受HTTP replay攻击的风险。"}
  Tls0Rtt *bool `json:"tls0Rtt,omitempty" xml:"tls0Rtt,omitempty" require:"true"`
  // {"en":"Default: False\nEnables Online Certificate Status Protocol (OCSP) stapling to check the revocation status of digital certificates. (Refer to https://en.wikipedia.org/wiki/OCSP_stapling)","zh_CN":"默认值: False\n是否启用在线证书状态协议（OCSP）装订以检查数字证书的吊销状态。参考：https://en.wikipedia.org/wiki/OCSP_stapling"}
  EnableOcspStapling *bool `json:"enableOcspStapling,omitempty" xml:"enableOcspStapling,omitempty" require:"true"`
  // {"en":"Default: False\nBy default, CDN Pro takes control of the contents under the /.well-known/{acme-challenge, pki-validation} directories to support certificate auto-renew for properties. If for any reason you need to manage these two directories by yourself on the origin, for example, to implement your own certificate auto-renew mechanism, you can use this configuration option to disable the default behavior by setting its value to true. For more information about our support for certificate auto-renewal, refer to the description of the autoRenew field in the Create a certificate API.","zh_CN":"默认值: False\n默认情况下，CDN Pro控制/.well-known/{acme-challenge, pki-validation}目录下的内容，以支持加速项目的证书自动更新功能。如果您需要自己在源站管理这两个目录，例如，为了实现您自己的证书自动更新机制，您可以将此字段设置为true来禁用默认行为。关于证书自动更新的更多信息，请参考'创建证书'接口中autoRenew字段的说明。"}
  DisableCertAutomation *bool `json:"disableCertAutomation,omitempty" xml:"disableCertAutomation,omitempty" require:"true"`
  // {"en":"Specify one or more access control rules to restrict access to your content. More advanced configuration can be done using Edge Logic. These access control rules take precedence over Edge Logic if both are defined.","zh_CN":"指定一个或多个访问控制规则以限制对内容的访问。可以使用边缘逻辑进行更高级的配置。此处定义的访问控制规则，优先级高于边缘逻辑。"}
  AccessControlRules []*GetAPropertyVersionResponseConfigsAccessControlRules `json:"accessControlRules,omitempty" xml:"accessControlRules,omitempty" require:"true" type:"Repeated"`
  // {"en":"This optional field allows you to configure notifications about client requests to be sent to a remote server. It can be used only if you have access to our realtime_log_switch directive. Please contact our support team if you require this feature.","zh_CN":"此可选字段用来配置发送消息通知（即实时日志）到您的远程服务器。当有客户端请求访问您的加速域名时，将触发通知。这是高级功能，如果您需要此功能，请联系我们的技术支持开通。"}
  RealTimeLog *GetAPropertyVersionResponseConfigsRealTimeLog `json:"realTimeLog,omitempty" xml:"realTimeLog,omitempty" require:"true" type:"Struct"`
  // {"en":"This object allows you to support video players requesting partial content through query string parameters. If you specify videoSeek, you must enter a value for startParameter.","zh_CN":"此对象用来支持视频播放器通过指定查询参数来请求部分内容。当videoSeek对象存在时，必须为startParameter设置一个值。"}
  VideoSeek *GetAPropertyVersionResponseConfigsVideoSeek `json:"videoSeek,omitempty" xml:"videoSeek,omitempty" require:"true" type:"Struct"`
  // {"en":"Range: <= 65530 characters\nThis field allows you to customize load balancing. The field will be deprecated soon. Please use the field edgeLogic.","zh_CN":"取值范围: <= 65530 字符\n用于控制负载均衡器的行为。该字段即将废弃，请使用edgeLogic字段进行配置。"}
  LoadBalancerLogic *string `json:"loadBalancerLogic,omitempty" xml:"loadBalancerLogic,omitempty" require:"true"`
  // {"en":"Range: <= 100 characters\nUsed to customize hash key for load balancer. The field will soon be deprecated. Please stop using it.","zh_CN":"取值范围: <= 100 字符\n用于自定义负载均衡器的哈希key。该字段即将废弃，请勿使用该字段。"}
  LoadBalancerHashKey *string `json:"loadBalancerHashKey,omitempty" xml:"loadBalancerHashKey,omitempty" require:"true"`
}

func (s GetAPropertyVersionResponseConfigs) String() string {
  return tea.Prettify(s)
}

func (s GetAPropertyVersionResponseConfigs) GoString() string {
  return s.String()
}

func (s *GetAPropertyVersionResponseConfigs) SetDescription(v string) *GetAPropertyVersionResponseConfigs {
  s.Description = &v
  return s
}

func (s *GetAPropertyVersionResponseConfigs) SetHostnames(v []*string) *GetAPropertyVersionResponseConfigs {
  s.Hostnames = v
  return s
}

func (s *GetAPropertyVersionResponseConfigs) SetHasBeian(v bool) *GetAPropertyVersionResponseConfigs {
  s.HasBeian = &v
  return s
}

func (s *GetAPropertyVersionResponseConfigs) SetBeianContentType(v int) *GetAPropertyVersionResponseConfigs {
  s.BeianContentType = &v
  return s
}

func (s *GetAPropertyVersionResponseConfigs) SetOrigins(v []*GetAPropertyVersionResponseConfigsOrigins) *GetAPropertyVersionResponseConfigs {
  s.Origins = v
  return s
}

func (s *GetAPropertyVersionResponseConfigs) SetAllowProtocolDowngrade(v bool) *GetAPropertyVersionResponseConfigs {
  s.AllowProtocolDowngrade = &v
  return s
}

func (s *GetAPropertyVersionResponseConfigs) SetEnableIpv6Origin(v bool) *GetAPropertyVersionResponseConfigs {
  s.EnableIpv6Origin = &v
  return s
}

func (s *GetAPropertyVersionResponseConfigs) SetFollowClientIpVersion(v string) *GetAPropertyVersionResponseConfigs {
  s.FollowClientIpVersion = &v
  return s
}

func (s *GetAPropertyVersionResponseConfigs) SetEdgeLogic(v string) *GetAPropertyVersionResponseConfigs {
  s.EdgeLogic = &v
  return s
}

func (s *GetAPropertyVersionResponseConfigs) SetSyntaxVersion(v int) *GetAPropertyVersionResponseConfigs {
  s.SyntaxVersion = &v
  return s
}

func (s *GetAPropertyVersionResponseConfigs) SetCacheKeyHostname(v string) *GetAPropertyVersionResponseConfigs {
  s.CacheKeyHostname = &v
  return s
}

func (s *GetAPropertyVersionResponseConfigs) SetCacheKeyUri(v string) *GetAPropertyVersionResponseConfigs {
  s.CacheKeyUri = &v
  return s
}

func (s *GetAPropertyVersionResponseConfigs) SetSchemeInCacheKey(v bool) *GetAPropertyVersionResponseConfigs {
  s.SchemeInCacheKey = &v
  return s
}

func (s *GetAPropertyVersionResponseConfigs) SetRedirectHttpToHttps(v string) *GetAPropertyVersionResponseConfigs {
  s.RedirectHttpToHttps = &v
  return s
}

func (s *GetAPropertyVersionResponseConfigs) SetDisableHttp2(v bool) *GetAPropertyVersionResponseConfigs {
  s.DisableHttp2 = &v
  return s
}

func (s *GetAPropertyVersionResponseConfigs) SetEnableHttp3(v bool) *GetAPropertyVersionResponseConfigs {
  s.EnableHttp3 = &v
  return s
}

func (s *GetAPropertyVersionResponseConfigs) SetExtraServicePorts(v *GetAPropertyVersionResponseConfigsExtraServicePorts) *GetAPropertyVersionResponseConfigs {
  s.ExtraServicePorts = v
  return s
}

func (s *GetAPropertyVersionResponseConfigs) SetTlsCertificateId(v string) *GetAPropertyVersionResponseConfigs {
  s.TlsCertificateId = &v
  return s
}

func (s *GetAPropertyVersionResponseConfigs) SetTlsCertificateId1(v int) *GetAPropertyVersionResponseConfigs {
  s.TlsCertificateId1 = &v
  return s
}

func (s *GetAPropertyVersionResponseConfigs) SetTlsMaxVersion(v string) *GetAPropertyVersionResponseConfigs {
  s.TlsMaxVersion = &v
  return s
}

func (s *GetAPropertyVersionResponseConfigs) SetTlsMinVersion(v string) *GetAPropertyVersionResponseConfigs {
  s.TlsMinVersion = &v
  return s
}

func (s *GetAPropertyVersionResponseConfigs) SetTlsCiphers(v string) *GetAPropertyVersionResponseConfigs {
  s.TlsCiphers = &v
  return s
}

func (s *GetAPropertyVersionResponseConfigs) SetTlsSessionTimeout(v int) *GetAPropertyVersionResponseConfigs {
  s.TlsSessionTimeout = &v
  return s
}

func (s *GetAPropertyVersionResponseConfigs) SetTls0Rtt(v bool) *GetAPropertyVersionResponseConfigs {
  s.Tls0Rtt = &v
  return s
}

func (s *GetAPropertyVersionResponseConfigs) SetEnableOcspStapling(v bool) *GetAPropertyVersionResponseConfigs {
  s.EnableOcspStapling = &v
  return s
}

func (s *GetAPropertyVersionResponseConfigs) SetDisableCertAutomation(v bool) *GetAPropertyVersionResponseConfigs {
  s.DisableCertAutomation = &v
  return s
}

func (s *GetAPropertyVersionResponseConfigs) SetAccessControlRules(v []*GetAPropertyVersionResponseConfigsAccessControlRules) *GetAPropertyVersionResponseConfigs {
  s.AccessControlRules = v
  return s
}

func (s *GetAPropertyVersionResponseConfigs) SetRealTimeLog(v *GetAPropertyVersionResponseConfigsRealTimeLog) *GetAPropertyVersionResponseConfigs {
  s.RealTimeLog = v
  return s
}

func (s *GetAPropertyVersionResponseConfigs) SetVideoSeek(v *GetAPropertyVersionResponseConfigsVideoSeek) *GetAPropertyVersionResponseConfigs {
  s.VideoSeek = v
  return s
}

func (s *GetAPropertyVersionResponseConfigs) SetLoadBalancerLogic(v string) *GetAPropertyVersionResponseConfigs {
  s.LoadBalancerLogic = &v
  return s
}

func (s *GetAPropertyVersionResponseConfigs) SetLoadBalancerHashKey(v string) *GetAPropertyVersionResponseConfigs {
  s.LoadBalancerHashKey = &v
  return s
}

type GetAPropertyVersionResponseConfigsOrigins struct     {
  // {"en":"^[a-zA-z0-9_]\nName of an origin. It must be unique within this property.","zh_CN":"^[a-zA-z0-9_]\n源站名称，在此加速项目中必须唯一。"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"Each entry should be a string consisting of an address optionally followed by one or more parameters, separated by space characters. The address can take one of the following three forms:\n{hostname or IP address}\n{hostname or IP address}:{http port}\n{hostname or IP address}:{http port},{https port}\nValues for the HTTP and HTTPS ports should be integers in the range [1,65535].\nIn the third form, one of the two ports can be empty, but the HTTPS port must always be specified after the comma.\n\nAccepts ipv6 addresses if 'enableIpv6Origin' is set to true, but an ipv6 address must be enclosed in [] so as to tell the port apart.\n\nSupported parameters are 'backup' and 'weight'.\n\n'backup' is used to indicate the server is a backup server. It will be used when the primary servers are unavailable.\n\n'weight' is a value in the range [1, 100]. The default value is 1. It lets you control the likelihood that one origin server will be used relative to another.\n\nThere should be at least one primary server defined; it does not have the 'backup' parameter.\n\nExample:\n['www.abc.com weight=1', 'www1.abc.com:8080 weight=4', 'www2.abc.com:880,4443 backup']\n\nThis example specifies two primary servers which are used in a 1:4 ratio meaning one gets about 20% of the requests to origin while the other gets 80% of the requests. In addition, a backup server is specified.","zh_CN":"源站服务器列表。每个条目为一个字符串，由一个地址与一个或多个参数组成，参数之间以空格分隔。地址可以采用以下三种形式之一：\n{域名或IP地址}\n{域名或IP地址}:{HTTP端口}\n{域名或IP地址}:{HTTP端口},{HTTPS端口}\nHTTP和HTTPS端口的值应为[1,65535]范围内的整数。\n在第三种形式中，HTTPS端口必须固定写在逗号后。如果HTTP端口未指定，仍需保留逗号，写为“,{HTTPS端口}”。\n\n当enableIpv6Origin设为true时，此处支持指定ipv6地址。ipv6地址应包裹在[]中，以便和端口区分开，例如[2001:0db8:85a3:0000:0000:8a2e:0370:7334]:80\n\n支持的参数包括'backup'和'weight'。\n'backup'用于标识备份服务器。它将在主服务器不可用时使用。\n'weight'是范围[1,100]内的值，默认值为1，用来设置服务器权重，控制一台源站服务器相对于另一台被使用的可能性。\n每个源站应至少定义一个主服务器（即不带'backup'参数的服务器）。\n\n示例：\n['www.abc.com weight=1', 'www1.abc.com:8080 weight=4', 'www2.abc.com:880,4443 backup']\n\n此示例指定了两个主服务器和一个备份服务器，其中，两个主服务器的权重为1:4，表示第一个服务器将获得约20%的回源请求，而另一个将获得约80%。"}
  Servers []*string `json:"servers,omitempty" xml:"servers,omitempty" require:"true" type:"Repeated"`
  // {"en":"Enum: http,https,both\nThis optional field indicates the protocol supported by the origin server. If this property has a certificate attached to it, the value can be set to http if the allowProtocolDowngrade field is also set to true.","zh_CN":"取值范围: http,https,both\n此可选字段表示源站服务器支持的协议。如果此加速项目附加了证书，且allowProtocolDowngrade字段也设置为true，则可以将该值设置为http。"}
  SupportedProtocol *string `json:"supportedProtocol,omitempty" xml:"supportedProtocol,omitempty" require:"true"`
  // {"en":"Default: auto\nOptional. This parameter tells us how important it is to directly go to this origin without going through any intermediate cache to fetch content. The values can be 'noDirect', 'auto', 'alwaysDirect'.\n\n'noDirect' means we always go through at least one intermediate cache. Customers who care particularly about the origin's workload can use this option. Our cache scheduler will dynamically pick the intermediate cache based on performance and resource availability.\n\n'auto' means the cache scheduler will make the choice dynamically based on cacheability of content. This is the default behavior if not specified.\n\n'alwaysDirect' means we always directly connect to this origin.","zh_CN":"默认值: auto\n此可选字段用来指定回源方式，可以是'noDirect'、'auto'、'alwaysDirect'。\n\n'noDirect'指总是通过至少一个中间缓存节点回源。对于特别关心源站负载的客户可以使用此选项。我们的调度程序将根据性能和资源可用性动态选择中间缓存节点。\n'alwaysDirect'指总是直接回源。\n'auto'指自动根据内容是否可缓存选择直接回源或通过中间缓存节点回源。对于可缓存的内容，将通过中间缓存节点回源；对于不可缓存的内容，将直接回源。该字段未指定时，采用此默认行为。"}
  DirectConnection *string `json:"directConnection,omitempty" xml:"directConnection,omitempty" require:"true"`
  // {"en":"Optional. It should be a valid hostname. It will also be used as the SNI servername when contacting the HTTPS origin. We also allow it to be an nginx variable that begins with $ followed by [a-z_]{3,60}. Nginx variable names are case insensitive, so we only allow lowercase characters.\nIf not specified, the value of the HOST header in the request will be used.\n\nOne constraint is that if 'aswS3' is specified for origin authentication, the value of hostHeader cannot be a variable or empty. It has to be a valid Amazon S3 hostname. Refer to https://docs.aws.amazon.com/AmazonS3/latest/dev/VirtualHosting.html","zh_CN":"可选，用来指定回源HOST头部。必须是有效的域名。当连接HTTPS源站时，该值也作为SNI域名。可以用nginx变量来指定，变量以'$'开头，后跟[a-z_]{3,60}。Nginx变量名不区分大小写，因此我们只允许小写字符。\n如果未指定，将使用请求中的HOST请求头的值。\n\n注意：如果指定了'awsS3'作为回源鉴权参数，则hostHeader的值不能为变量或为空，而必须是有效的Amazon S3主机名。参考：https://docs.aws.amazon.com/AmazonS3/latest/dev/VirtualHosting.html"}
  HostHeader *string `json:"hostHeader,omitempty" xml:"hostHeader,omitempty" require:"true"`
  // {"en":"Default: False\nOptional. It controls whether the cache will validate the certificate of the origin.","zh_CN":"默认值: False\n可选。用来设定CDN缓存节点是否验证源站证书。"}
  VerifyOrigin *bool `json:"verifyOrigin,omitempty" xml:"verifyOrigin,omitempty" require:"true"`
  // {"en":"Optional. It is a structure to specify the parameters, such as password, for authentication with the origin servers. It should have at least one field: 'methodName', which should be a string indicating one of the predefined authentication methods. The only valid value at this time is 'awsS3'. When 'awsS3' is used, the following parameters are required:\nawsS3.region\nawsS3.accessKey\nawsS3.secretKey\nAlso, the value of the hostHeader field cannot be a variable or empty. It must be a valid Amazon S3 hostname.\n\nExample:\n<code>{\"methodName\":\"awsS3\",\n\"awsS3\":{\n\"region\":\"us-west-1\",\n\"accessKey\":\"sdnu2932\",\n\"secretKey\":\"d12345678abcdefghi\"\n}}</code>","zh_CN":"可选。用于指定与源服务器进行身份验证（回源鉴权）的相关参数（如密码）。必须至少有'methodName'（字符串）字段，用来指定预定义的鉴权方法。目前仅支持源站为AWS S3的回源鉴权，因此唯一有效的值是'awsS3'。使用'awsS3'时，需要提供以下参数：\nawsS3.region\nawsS3.accessKey\nawsS3.secretKey\n此外，hostHeader字段的值不能为变量或为空，必须是有效的Amazon S3主机名。\n\n示例：\n<code>{\"methodName\":\"awsS3\",\n\"awsS3\":{\n\"region\":\"us-west-1\",\n\"accessKey\":\"sdnu2932\",\n\"secretKey\":\"d12345678abcdefghi\"\n}}</code>"}
  Authentication *GetAPropertyVersionResponseConfigsOriginsAuthentication `json:"authentication,omitempty" xml:"authentication,omitempty" require:"true" type:"Struct"`
  // {"en":"Default: 60 Range: [ 5 .. 600 ]\nTimeout in seconds during which an idle keepalive connection to an upstream server will stay open. A service quota setting of maxUpstreamKeepaliveTimeOut can change the maximum permitted value.","zh_CN":"默认值: 60 取值范围: [ 5 .. 600 ]\n该字段用于指定CDN Pro服务器和源站建连的Keep-Alive超时时间，单位为秒。通过maxUpstreamKeepaliveTimeOut 该服务设置项可以更改允许的最大值。如果需要调整最大值，请联系我们的技术支持。"}
  KeepAliveTimeout *int `json:"keepAliveTimeout,omitempty" xml:"keepAliveTimeout,omitempty" require:"true"`
  // {"en":"This setting allows you to specify the number of unsuccessful attempts to reach one of the origin's IP addresses or peers before it is marked as unavailable for a period of time designated by the timeout. If all peers of all servers are unavailable, requests for content from the origin will receive an immediate 502 HTTP response. By default, six attempts to a peer are made, after which a one-second timeout applies to that peer. To disable this feature, set peerFailureTimeout to 'off'.\nExample: <code>{\"peerFailureTimeout\":{\"failureThreshold\":10,\"timeout\":3}}</code>","zh_CN":"该字段用于设置源站故障剔除。开启此功能后，当连接某个源站服务器的失败次数达到设定阈值，该源站服务器将被标记为不可用，并保持该状态直到设定的超时时长。失败次数阈值和超时时长分别通过failureThreshold参数和timeout参数设置。CDN Pro回源时将剔除不可用的源站服务器。如果所有源站服务器都被标记为不可用，则对应的请求将立即响应502状态码。默认情况下，当连接某个源站服务器连续失败6次之后，该服务器将被标记为不可用，超时时长为1秒。如果要禁用此功能，请将peerFailureTimeout设置为'off'。开启源站故障剔除配置示例：<code>{\"peerFailureTimeout\":{\"failureThreshold\":10,\"timeout\":3}}</code>"}
  PeerFailureTimeout *string `json:"peerFailureTimeout,omitempty" xml:"peerFailureTimeout,omitempty" require:"true"`
  // {"en":"Refers to a certificate used to authenticate with the origin server. This certificate must also be deployed.","zh_CN":"用于源站验证CDN Pro服务器的证书，该证书同样必须被部署。"}
  TlsCertificateId *string `json:"tlsCertificateId,omitempty" xml:"tlsCertificateId,omitempty" require:"true"`
  // {"en":"Refers to the ID of an origin shield representing a set of servers that will make requests to your origin servers rather than the CDN Pro edge servers. This can help further reduce traffic to your origin. This is an advanced feature that can be enabled by contacting our support team. The origin shields API provides a list of available shields along with their IP addresses. It is best to select a shield whose region is closest to your origin servers. Use of a shield in China requires your property have 'hasBeian' set to true. If your origin servers have an access control list, add the IPs of the shield you choose to use.","zh_CN":"指定某个源站盾（origin shield）的ID。使用源站盾功能后，所有回源请求都会通过源站盾中转回源，这可以帮助收敛回源流量。源站盾是高级功能，如需使用请联系我们的技术支持开通。可通过调用‘获取源站盾列表’接口查询可用的源站盾及其对应的IP地址。您可根据源站的位置，选择合适的源站盾。如果您需使用中国大陆的源站盾，则该加速项目的所有域名必须已取得备案。如果您的源站开启了访问控制，请将所选择源站盾的IP地址加入白名单。"}
  Shield *string `json:"shield,omitempty" xml:"shield,omitempty" require:"true"`
}

func (s GetAPropertyVersionResponseConfigsOrigins) String() string {
  return tea.Prettify(s)
}

func (s GetAPropertyVersionResponseConfigsOrigins) GoString() string {
  return s.String()
}

func (s *GetAPropertyVersionResponseConfigsOrigins) SetName(v string) *GetAPropertyVersionResponseConfigsOrigins {
  s.Name = &v
  return s
}

func (s *GetAPropertyVersionResponseConfigsOrigins) SetServers(v []*string) *GetAPropertyVersionResponseConfigsOrigins {
  s.Servers = v
  return s
}

func (s *GetAPropertyVersionResponseConfigsOrigins) SetSupportedProtocol(v string) *GetAPropertyVersionResponseConfigsOrigins {
  s.SupportedProtocol = &v
  return s
}

func (s *GetAPropertyVersionResponseConfigsOrigins) SetDirectConnection(v string) *GetAPropertyVersionResponseConfigsOrigins {
  s.DirectConnection = &v
  return s
}

func (s *GetAPropertyVersionResponseConfigsOrigins) SetHostHeader(v string) *GetAPropertyVersionResponseConfigsOrigins {
  s.HostHeader = &v
  return s
}

func (s *GetAPropertyVersionResponseConfigsOrigins) SetVerifyOrigin(v bool) *GetAPropertyVersionResponseConfigsOrigins {
  s.VerifyOrigin = &v
  return s
}

func (s *GetAPropertyVersionResponseConfigsOrigins) SetAuthentication(v *GetAPropertyVersionResponseConfigsOriginsAuthentication) *GetAPropertyVersionResponseConfigsOrigins {
  s.Authentication = v
  return s
}

func (s *GetAPropertyVersionResponseConfigsOrigins) SetKeepAliveTimeout(v int) *GetAPropertyVersionResponseConfigsOrigins {
  s.KeepAliveTimeout = &v
  return s
}

func (s *GetAPropertyVersionResponseConfigsOrigins) SetPeerFailureTimeout(v string) *GetAPropertyVersionResponseConfigsOrigins {
  s.PeerFailureTimeout = &v
  return s
}

func (s *GetAPropertyVersionResponseConfigsOrigins) SetTlsCertificateId(v string) *GetAPropertyVersionResponseConfigsOrigins {
  s.TlsCertificateId = &v
  return s
}

func (s *GetAPropertyVersionResponseConfigsOrigins) SetShield(v string) *GetAPropertyVersionResponseConfigsOrigins {
  s.Shield = &v
  return s
}

type GetAPropertyVersionResponseConfigsOriginsAuthentication struct {
  // {"en":"Authentication method.","zh_CN":"鉴权方法。"}
  MethodName *string `json:"methodName,omitempty" xml:"methodName,omitempty" require:"true"`
}

func (s GetAPropertyVersionResponseConfigsOriginsAuthentication) String() string {
  return tea.Prettify(s)
}

func (s GetAPropertyVersionResponseConfigsOriginsAuthentication) GoString() string {
  return s.String()
}

func (s *GetAPropertyVersionResponseConfigsOriginsAuthentication) SetMethodName(v string) *GetAPropertyVersionResponseConfigsOriginsAuthentication {
  s.MethodName = &v
  return s
}

type GetAPropertyVersionResponseConfigsExtraServicePorts struct {
  // {"en":"This is a list of ports other than 80 which are used to handle HTTP requests. The available values can be found in the systemConfigs API's response. If you need another port, please contact our support team.","zh_CN":"指定用于处理HTTP请求的端口列表（80端口除外）。可通过调用'获取系统配置'接口来查询系统支持的端口。如果您需要开通其他端口，请联系技术支持。"}
  Http []*string `json:"http,omitempty" xml:"http,omitempty" require:"true" type:"Repeated"`
  // {"en":"This is a list of ports other than 443 which are used to handle HTTPS requests. The available values can be found in the systemConfigs API's response. If you need another port, please contact our support team.","zh_CN":"指定用于处理HTTPS请求的端口列表（443端口除外）。可通过调用'获取系统配置'接口来查询系统支持的端口。如果您需要开通其他端口，请联系技术支持。"}
  Https []*string `json:"https,omitempty" xml:"https,omitempty" require:"true" type:"Repeated"`
}

func (s GetAPropertyVersionResponseConfigsExtraServicePorts) String() string {
  return tea.Prettify(s)
}

func (s GetAPropertyVersionResponseConfigsExtraServicePorts) GoString() string {
  return s.String()
}

func (s *GetAPropertyVersionResponseConfigsExtraServicePorts) SetHttp(v []*string) *GetAPropertyVersionResponseConfigsExtraServicePorts {
  s.Http = v
  return s
}

func (s *GetAPropertyVersionResponseConfigsExtraServicePorts) SetHttps(v []*string) *GetAPropertyVersionResponseConfigsExtraServicePorts {
  s.Https = v
  return s
}

type GetAPropertyVersionResponseConfigsAccessControlRules struct     {
  // {"en":"Range: [ 0 .. 60 ] characters\nAn optional ID for the access control rule.","zh_CN":"取值范围: [ 0 .. 60 ] 字符\n访问控制规则ID。"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {"en":"Specify the conditions that the incoming request must match. At least one condition must be specified. If multiple are specified, all must match.","zh_CN":"指定客户端请求必须匹配的条件。必须至少指定一个条件。如果指定了多个条件，则必须全部匹配。"}
  Conditions *GetAPropertyVersionResponseConfigsAccessControlRulesConditions `json:"conditions,omitempty" xml:"conditions,omitempty" require:"true" type:"Struct"`
  // {"en":"Indicates the action to take in response to a request that matches the conditions of the access control rule.","zh_CN":"对于匹配到以上条件的请求所采取的相应操作。"}
  Action *GetAPropertyVersionResponseConfigsAccessControlRulesAction `json:"action,omitempty" xml:"action,omitempty" require:"true" type:"Struct"`
}

func (s GetAPropertyVersionResponseConfigsAccessControlRules) String() string {
  return tea.Prettify(s)
}

func (s GetAPropertyVersionResponseConfigsAccessControlRules) GoString() string {
  return s.String()
}

func (s *GetAPropertyVersionResponseConfigsAccessControlRules) SetId(v string) *GetAPropertyVersionResponseConfigsAccessControlRules {
  s.Id = &v
  return s
}

func (s *GetAPropertyVersionResponseConfigsAccessControlRules) SetConditions(v *GetAPropertyVersionResponseConfigsAccessControlRulesConditions) *GetAPropertyVersionResponseConfigsAccessControlRules {
  s.Conditions = v
  return s
}

func (s *GetAPropertyVersionResponseConfigsAccessControlRules) SetAction(v *GetAPropertyVersionResponseConfigsAccessControlRulesAction) *GetAPropertyVersionResponseConfigsAccessControlRules {
  s.Action = v
  return s
}

type GetAPropertyVersionResponseConfigsAccessControlRulesConditions struct {
  // {"en":"Enum: https,http\nIndicates whether the incoming request uses HTTP or HTTPS.","zh_CN":"取值范围: https,http\n客户端请求的协议，HTTP或HTTPS。"}
  Scheme *string `json:"scheme,omitempty" xml:"scheme,omitempty" require:"true"`
  // {"en":"Indicates the hostname requested. It must be one of the hostnames defined in the property. If a wildcard hostname is configured in the property, you can specify either the wildcard hostname or their spcific hostnames here.","zh_CN":"客户端请求对应的域名。必须是加速项目中定义的加速域名之一。如果加速域名按泛域名配置，此处可以指定泛域名或者泛域名下的某个具体域名。"}
  Hostname *string `json:"hostname,omitempty" xml:"hostname,omitempty" require:"true"`
  // {"en":"Range: <= 500 characters\nThe URI begins with '/' and can end with '*' for prefix matching.","zh_CN":"取值范围: <= 500 字符\n用于前缀匹配的URI，以'/'开头，可以以'*'结尾，表示模糊匹配。"}
  Uri *string `json:"uri,omitempty" xml:"uri,omitempty" require:"true"`
  // {"en":"An array of ISO-3166-1 alpha-2 codes indicating the countries of the servers, for example, 'US' to represent 'United States of America'.","zh_CN":"服务器所在区域，以ISO-3166-1 两位国家代码表示。例如，'US'代表'美国'。"}
  ServerRegions []*string `json:"serverRegions,omitempty" xml:"serverRegions,omitempty" require:"true" type:"Repeated"`
  // {"en":"An array of ISO-3166-1 alpha-2 codes indicating one or more countries which are the source of the client requests, for example, 'CN' to represent 'China'.","zh_CN":"客户端请求来源区域，以ISO-3166-1 两位国家代码表示。例如，'CN'代表'中国'。"}
  ClientRegions []*string `json:"clientRegions,omitempty" xml:"clientRegions,omitempty" require:"true" type:"Repeated"`
  // {"en":"Indicates the starting and ending IP addresses of the client requests to match against. They must be in IPv4 or IPv6 format.","zh_CN":"用于指定要匹配的客户端请求的开始和结束IP地址。必须是IPv4或IPv6格式。"}
  ClientIpRange []*string `json:"clientIpRange,omitempty" xml:"clientIpRange,omitempty" require:"true" type:"Repeated"`
}

func (s GetAPropertyVersionResponseConfigsAccessControlRulesConditions) String() string {
  return tea.Prettify(s)
}

func (s GetAPropertyVersionResponseConfigsAccessControlRulesConditions) GoString() string {
  return s.String()
}

func (s *GetAPropertyVersionResponseConfigsAccessControlRulesConditions) SetScheme(v string) *GetAPropertyVersionResponseConfigsAccessControlRulesConditions {
  s.Scheme = &v
  return s
}

func (s *GetAPropertyVersionResponseConfigsAccessControlRulesConditions) SetHostname(v string) *GetAPropertyVersionResponseConfigsAccessControlRulesConditions {
  s.Hostname = &v
  return s
}

func (s *GetAPropertyVersionResponseConfigsAccessControlRulesConditions) SetUri(v string) *GetAPropertyVersionResponseConfigsAccessControlRulesConditions {
  s.Uri = &v
  return s
}

func (s *GetAPropertyVersionResponseConfigsAccessControlRulesConditions) SetServerRegions(v []*string) *GetAPropertyVersionResponseConfigsAccessControlRulesConditions {
  s.ServerRegions = v
  return s
}

func (s *GetAPropertyVersionResponseConfigsAccessControlRulesConditions) SetClientRegions(v []*string) *GetAPropertyVersionResponseConfigsAccessControlRulesConditions {
  s.ClientRegions = v
  return s
}

func (s *GetAPropertyVersionResponseConfigsAccessControlRulesConditions) SetClientIpRange(v []*string) *GetAPropertyVersionResponseConfigsAccessControlRulesConditions {
  s.ClientIpRange = v
  return s
}

type GetAPropertyVersionResponseConfigsAccessControlRulesAction struct {
  // {"en":"Indicates the HTTP status code to respond with. It must be in the range 300-309, 400-409, or 500-509 to indicate a redirection or error.","zh_CN":"响应的HTTP状态码，范围必须在300-309、400-409或500-509之间，分别表示重定向或错误。"}
  Status *string `json:"status,omitempty" xml:"status,omitempty" require:"true"`
  // {"en":"Range: <= 200 characters\nIf the value of the status field is in the range 300-309, the message field's value will be placed in a Location HTTP header returned to the client. If the status is in the range 400-409 or 500-509, the value will be returned in the response body to the client.","zh_CN":"取值范围: <= 200 字符\n如果status字段的值在300-309范围内，message字段的值将在返回给客户端的Location响应头中。如果status字段的值在400-409或500-509范围内，则message字段的值将在返回给客户端的响应体中。"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s GetAPropertyVersionResponseConfigsAccessControlRulesAction) String() string {
  return tea.Prettify(s)
}

func (s GetAPropertyVersionResponseConfigsAccessControlRulesAction) GoString() string {
  return s.String()
}

func (s *GetAPropertyVersionResponseConfigsAccessControlRulesAction) SetStatus(v string) *GetAPropertyVersionResponseConfigsAccessControlRulesAction {
  s.Status = &v
  return s
}

func (s *GetAPropertyVersionResponseConfigsAccessControlRulesAction) SetMessage(v string) *GetAPropertyVersionResponseConfigsAccessControlRulesAction {
  s.Message = &v
  return s
}

type GetAPropertyVersionResponseConfigsRealTimeLog struct {
  // {"en":"The URL that receives the notifications. It must begin with 'http' or 'https'. The server should support the POST method. This is a required field.","zh_CN":"接收通知的服务器URL地址。必须以'http'或'https'开头。服务器须支持POST方法。这是必填字段。"}
  LogUrl *string `json:"logUrl,omitempty" xml:"logUrl,omitempty" require:"true"`
  // {"en":"Default: 1 Range: [ 1 .. 65536 ]\nAn integer between [1, 65536]. n means one notification for every n edge requests. 1 means every edge request will generate a notification. This is optional. Default is 1.","zh_CN":"默认值: 1 取值范围: [ 1 .. 65536 ]\n采样率。介于[1, 65536]之间的整数。n表示每n个边缘请求发送1条通知。1表示每个边缘请求都会发送。这是可选字段。默认值为1。"}
  SampleRate *int `json:"sampleRate,omitempty" xml:"sampleRate,omitempty" require:"true"`
  // {"en":"Enum: json,none,compactJson\nSpecify json to enable JSON character escaping in variables, compactJson to enable JSON character escaping with whitespace in the body removed, or none to disable escaping.","zh_CN":"取值范围: json,none,compactJson\n指定json以开启变量中的json字符转义；指定compactJson以开启变量中的json字符转义，并移除变量中的空白行。如果要禁用转义则指定none。"}
  Escape *string `json:"escape,omitempty" xml:"escape,omitempty" require:"true"`
  // {"en":"This is a required field to specify the notification body to be sent to the remote server. It can be any printable text that can be sent in the body of an HTTP POST request.\n\nThe CDN Pro built-in variables can be used within the format field. They will be replaced with the actual values in the notifications.","zh_CN":"该字段用来定义要发送到远程服务器的通知的格式。通知正文可以包括任何能在HTTP POST请求体中发送的可打印文本。这是必填字段。\n\n可以在格式定义中使用CDN Pro支持的内置变量。发送通知时，变量将被替换为实际值。"}
  Format *string `json:"format,omitempty" xml:"format,omitempty" require:"true"`
  // {"en":"HTTP header names and values to be sent to the notification server. A header name can contain any alphanumeric character or hyphen, '-'. A header value can contain any printable characters. It can also include any of the built-in variables supported in the format field of the realTimeLog object.","zh_CN":"需要发送到远程服务器的HTTP请求头名称和值。请求头名称可以包含任何字母，数字或连字符'-'。值可以包含任何可打印字符，也可以使用realTimeLog对象format字段中支持的任何内置变量。"}
  Headers []*GetAPropertyVersionResponseConfigsRealTimeLogHeaders `json:"headers,omitempty" xml:"headers,omitempty" require:"true" type:"Repeated"`
}

func (s GetAPropertyVersionResponseConfigsRealTimeLog) String() string {
  return tea.Prettify(s)
}

func (s GetAPropertyVersionResponseConfigsRealTimeLog) GoString() string {
  return s.String()
}

func (s *GetAPropertyVersionResponseConfigsRealTimeLog) SetLogUrl(v string) *GetAPropertyVersionResponseConfigsRealTimeLog {
  s.LogUrl = &v
  return s
}

func (s *GetAPropertyVersionResponseConfigsRealTimeLog) SetSampleRate(v int) *GetAPropertyVersionResponseConfigsRealTimeLog {
  s.SampleRate = &v
  return s
}

func (s *GetAPropertyVersionResponseConfigsRealTimeLog) SetEscape(v string) *GetAPropertyVersionResponseConfigsRealTimeLog {
  s.Escape = &v
  return s
}

func (s *GetAPropertyVersionResponseConfigsRealTimeLog) SetFormat(v string) *GetAPropertyVersionResponseConfigsRealTimeLog {
  s.Format = &v
  return s
}

func (s *GetAPropertyVersionResponseConfigsRealTimeLog) SetHeaders(v []*GetAPropertyVersionResponseConfigsRealTimeLogHeaders) *GetAPropertyVersionResponseConfigsRealTimeLog {
  s.Headers = v
  return s
}

type GetAPropertyVersionResponseConfigsRealTimeLogHeaders struct     {
  // {"en":"Name of an HTTP header.","zh_CN":"HTTP标头名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"Value of an HTTP header.","zh_CN":"HTTP标头值"}
  Value *string `json:"value,omitempty" xml:"value,omitempty" require:"true"`
}

func (s GetAPropertyVersionResponseConfigsRealTimeLogHeaders) String() string {
  return tea.Prettify(s)
}

func (s GetAPropertyVersionResponseConfigsRealTimeLogHeaders) GoString() string {
  return s.String()
}

func (s *GetAPropertyVersionResponseConfigsRealTimeLogHeaders) SetName(v string) *GetAPropertyVersionResponseConfigsRealTimeLogHeaders {
  s.Name = &v
  return s
}

func (s *GetAPropertyVersionResponseConfigsRealTimeLogHeaders) SetValue(v string) *GetAPropertyVersionResponseConfigsRealTimeLogHeaders {
  s.Value = &v
  return s
}

type GetAPropertyVersionResponseConfigsVideoSeek struct {
  // {"en":"Range: [ 1 .. 31 ] characters\nName of the query parameter indicating the starting offset in bytes of the content to fetch. The parameter name should begin with a letter (a-z, A-Z) and may be followed by up to 30 letters and numbers.","zh_CN":"取值范围: [ 1 .. 31 ] 字符\n查询参数的名称，用来指定要获取的内容的起始位置（以字节计算）。参数名称应以字母（a-z，A-Z）开头，后面最多可以有30个字母和数字。"}
  StartParameter *string `json:"startParameter,omitempty" xml:"startParameter,omitempty" require:"true"`
  // {"en":"Range: [ 0 .. 31 ] characters\nName of the query parameter indicating the ending offset of the content to fetch. The parameter name should begin with a letter (a-z, A-Z) and may be followed by up to 30 letters and numbers.","zh_CN":"取值范围: [ 0 .. 31 ] 字符\n查询参数的名称，用来指定要获取的内容的结束位置（以字节计算）。参数名称应以字母（a-z，A-Z）开头，后面最多可以有30个字母和数字。"}
  EndParameter *string `json:"endParameter,omitempty" xml:"endParameter,omitempty" require:"true"`
}

func (s GetAPropertyVersionResponseConfigsVideoSeek) String() string {
  return tea.Prettify(s)
}

func (s GetAPropertyVersionResponseConfigsVideoSeek) GoString() string {
  return s.String()
}

func (s *GetAPropertyVersionResponseConfigsVideoSeek) SetStartParameter(v string) *GetAPropertyVersionResponseConfigsVideoSeek {
  s.StartParameter = &v
  return s
}

func (s *GetAPropertyVersionResponseConfigsVideoSeek) SetEndParameter(v string) *GetAPropertyVersionResponseConfigsVideoSeek {
  s.EndParameter = &v
  return s
}

type GetAPropertyVersionResponseStatus struct {
  // {"en":"Result of the last validation of the property version.","zh_CN":"该版本最近一次配置验证的结果。"}
  LastValidationStatus *string `json:"lastValidationStatus,omitempty" xml:"lastValidationStatus,omitempty" require:"true"`
  // {"en":"Flag indicating if the version is frozen. A property version will be freezed after it is deployed to either staging or production environment.","zh_CN":"该版本是否被冻结。当一个版本部署到演练或生产环境后，该版本即被冻结，不可更改。"}
  Frozen *bool `json:"frozen,omitempty" xml:"frozen,omitempty" require:"true"`
  // {"en":"The datetime when the version is created.","zh_CN":"该版本创建时间。"}
  CreationTime *string `json:"creationTime,omitempty" xml:"creationTime,omitempty" require:"true"`
  // {"en":"The datetime when the version is last updated.","zh_CN":"该版本最近一次更新时间。"}
  LastUpdateTime *string `json:"lastUpdateTime,omitempty" xml:"lastUpdateTime,omitempty" require:"true"`
  // {"en":"Flag indicating if the version is deployed to production.","zh_CN":"该版本是否部署到生产环境。"}
  InProduction *bool `json:"inProduction,omitempty" xml:"inProduction,omitempty" require:"true"`
  // {"en":"Flag indicating if the version is deployed to staging.","zh_CN":"该版本是否部署到演练环境。"}
  InStaging *bool `json:"inStaging,omitempty" xml:"inStaging,omitempty" require:"true"`
}

func (s GetAPropertyVersionResponseStatus) String() string {
  return tea.Prettify(s)
}

func (s GetAPropertyVersionResponseStatus) GoString() string {
  return s.String()
}

func (s *GetAPropertyVersionResponseStatus) SetLastValidationStatus(v string) *GetAPropertyVersionResponseStatus {
  s.LastValidationStatus = &v
  return s
}

func (s *GetAPropertyVersionResponseStatus) SetFrozen(v bool) *GetAPropertyVersionResponseStatus {
  s.Frozen = &v
  return s
}

func (s *GetAPropertyVersionResponseStatus) SetCreationTime(v string) *GetAPropertyVersionResponseStatus {
  s.CreationTime = &v
  return s
}

func (s *GetAPropertyVersionResponseStatus) SetLastUpdateTime(v string) *GetAPropertyVersionResponseStatus {
  s.LastUpdateTime = &v
  return s
}

func (s *GetAPropertyVersionResponseStatus) SetInProduction(v bool) *GetAPropertyVersionResponseStatus {
  s.InProduction = &v
  return s
}

func (s *GetAPropertyVersionResponseStatus) SetInStaging(v bool) *GetAPropertyVersionResponseStatus {
  s.InStaging = &v
  return s
}

type GetAPropertyVersionResponseHeader struct {
}

func (s GetAPropertyVersionResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetAPropertyVersionResponseHeader) GoString() string {
  return s.String()
}




type GetListOfPropertiesRequest struct {
}

func (s GetListOfPropertiesRequest) String() string {
  return tea.Prettify(s)
}

func (s GetListOfPropertiesRequest) GoString() string {
  return s.String()
}

type GetListOfPropertiesRequestHeader struct {
}

func (s GetListOfPropertiesRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetListOfPropertiesRequestHeader) GoString() string {
  return s.String()
}

type GetListOfPropertiesPaths struct {
}

func (s GetListOfPropertiesPaths) String() string {
  return tea.Prettify(s)
}

func (s GetListOfPropertiesPaths) GoString() string {
  return s.String()
}

type GetListOfPropertiesParameters struct {
  // {"en":"The value of the search parameter is matched on the id, name, description, and hostnames fields of each property. A value like 'domain.com' would match on hostnames domain.com and abc.123domain.com.\n\nUse a value beginning with the '^' character to narrow the matches to properties with name, description, or hostnames fields beginning with the text following the '^'. For example, '^domain' would match on hostnames domain.com and domain123.com but not abc.123domain.com.\n\nSince a limited number of characters are permitted in URIs, you must encode '^' as '%5E' when specifying a search parameter containing it in the URI, for example, curl -i --url 'https://ngapi.cdnetworks.com/cdn/properties?search=%5Edomain' ... .\n\nNote that a property will be returned if the search matches on any version of the property which has not been deleted. Also, when searching for a property by its ID, you must specify the entire ID as partial matches are not supported.","zh_CN":"根据搜索关键字匹配每个加速项目的id、name、description以及hostnames字段进行过滤。例如，'domain.com'将匹配到加速域名'domain.com'和'abc.123domain.com'。可以使用'^'字符限定只匹配以搜索关键字开头的值。例如，'^domain'将匹配到加速域名'domain.com'和'domain123.com'，但不会匹配到'abc.123domain.com'。由于URI仅允许有限数量的字符，当search参数包含^符号时，必须将'^'编码为'%5E'。例如，curl -i --url 'https://openapi.wangsu.com/cdn/properties?search=%5Edomain' ... 。注意，该搜索关键字将会对加速项目的所有未删除的版本进行匹配。当通过ID搜索加速项目时，必须指定完整ID，不支持部分匹配。"}
  Search *string `json:"search,omitempty" xml:"search,omitempty"`
  // {"en":"Service type, one of wsapro, webpro, vodpro , downloadpro , 1523.","zh_CN":"服务类型，根据实际服务类型传wsapro，webpro，vodpro，downloadpro，cdnpro，或MALL-cdnpro，分别指全站加速，网页加速，点播加速，下载加速，CDN Pro，以及自助CDN。。"}
  LegacyType *string `json:"legacyType,omitempty" xml:"legacyType,omitempty"`
  // {"en":"The value of hasConfig is used to filter the results. It can be any field name of a property version, optionally followed by a colon and a value. For example:\n\nhasConfig=hasBeian:true\n\nWhen a colon and value are specified, only properties whose configurations have that value for the field will be returned. The value can be preceded by an exclamation mark, '!', to invert the search.  For example:\n\nhasConfig=edgeLogic!sorted. In this case, we check that the edgeLogic field does not include the value 'sorted'.\n\nSearches for values of numeric or boolean fields require an exact match. However, partial matches are supported for string fields. For example, hasConfig=hostnames:domain returns all properties with a hostname containing the string 'domain' such as 'mydomain.com' and 'thedomains.com'.\n\nIf a colon and value are omitted, all properties with a non-empty value for the field will be returned.\n\nSpecify multiple hasConfig parameters to filter on multiple conditions. Example:\n\nhasConfig=hasBeian:true&hasConfig=realTimeLog Only properties matching all the parameters will be returned.\n\nIf a property has multiple versions, the property will be included in the API response if any of its versions matches the hasConfig parameter.\n\nSubfields of a property version can be specified by using the period symbol as a separator. Refer to additional examples below.\n\nIf matching against values with spaces or special characters, remember to encode them in the URL.\n\nAdditional examples:\n\n<table><tr><th>Value</th><th>Description</th></tr><tr><td>hasConfig=disableHttp2:true</td><td>Return properties that don't support HTTP2</td></tr><tr><td>hasConfig=extraServicePorts.http:85</td><td>Return properties supporting port 85 for HTTP requests.</td></tr><tr><td>hasConfig=origins.servers:myorigin.com</td><td>Return properties using myorigin.com as an origin server.</td></tr></table>","zh_CN":"通过hasConfig参数指定字段来过滤加速项目。支持使用加速项目版本的任何字段名，在字段名后跟上冒号和值（可选）进行过滤。例如：hasConfig=hasBeian:true。当指定冒号和值时，只返回匹配该字段值的加速项目。如果需要反向查询，可以在字段值前面加上感叹号'!'。例如：hasConfig=edgeLogic!sorted，这个查询将筛选出edgeLogic字段不包含值'sorted'的加速项目。查询数字或布尔类型字段的值时需要精确匹配，而字符串字段支持部分匹配。例如，hasConfig=hostnames:domain将返回加速域名包含字符串'domain'的所有加速项目，例如'mydomain.com'和'thedomains.com'。如果省略冒号和值，则将返回指定字段值为非空的所有加速项目。可以指定多个hasConfig参数使用多个条件进行查询。例如：hasConfig=hasBeian:true&hasConfig=realTimeLog，这个查询只返回与所有参数匹配的加速项目。如果一个加速项目有多个版本，只要该加速项目的任何版本与hasConfig参数匹配，则该加速项目将包含在该接口的响应中。可以使用点号作为分隔符来指定某个字段的子字段。如果要匹配的值带有空格或特殊字符，请在URL中对其进行编码。以下是一些示例： <table><tr><th>值</th><th>描述</th></tr><tr><td>hasConfig=disableHttp2:true</td><td>返回禁用HTTP2的加速项目。</td></tr><tr><td>hasConfig=extraServicePorts.http:85</td><td>返回支持HTTP请求端口85的加速项目。</td></tr><tr><td>hasConfig=origins.servers:myorigin.com</td><td>返回使用myorigin.com作为源站的加速项目。</td></tr></table>"}
  HasConfig *string `json:"hasConfig,omitempty" xml:"hasConfig,omitempty"`
  // {"en":"Enum: all,staging,production\nDefault: all\nThe value can be 'all', 'staging', or 'production' to filter the results based on where the property has been deployed.","zh_CN":"取值范围: all,staging,production\n默认值: all\n根据加速项目的部署环境过滤。该值可以是'all', 'staging', 或'production'，分别表示所有环境，演练环境和生产环境。"}
  Target *string `json:"target,omitempty" xml:"target,omitempty"`
  // {"en":"Default: 0\nIndicates the first item to return. The default is '0'.","zh_CN":"默认值: 0 取值范围: >= 0\n查询起始位置。"}
  Offset *int `json:"offset,omitempty" xml:"offset,omitempty"`
  // {"en":"Default: 100 Range: <= 200\nMaximum number of properties to return. The default is to return a summary of all properties.","zh_CN":"默认值: 100 取值范围: <= 200\n每次查询的最大条数。"}
  Limit *int `json:"limit,omitempty" xml:"limit,omitempty"`
  // {"en":"Enum: asc,desc\nDefault: desc\nOrder of properties to return. The default is to return the last one updated first.","zh_CN":"取值范围: asc,desc\n默认值: desc\n返回结果的顺序。默认按最后更新时间降序。"}
  SortOrder *string `json:"sortOrder,omitempty" xml:"sortOrder,omitempty"`
  // {"en":"Enum: creationTime,lastUpdateTime\nDefault: lastUpdateTime\nReturns results in sorted order.","zh_CN":"取值范围: creationTime,lastUpdateTime\n默认值: lastUpdateTime\n返回结果的排序依据。"}
  SortBy *string `json:"sortBy,omitempty" xml:"sortBy,omitempty"`
}

func (s GetListOfPropertiesParameters) String() string {
  return tea.Prettify(s)
}

func (s GetListOfPropertiesParameters) GoString() string {
  return s.String()
}

func (s *GetListOfPropertiesParameters) SetSearch(v string) *GetListOfPropertiesParameters {
  s.Search = &v
  return s
}

func (s *GetListOfPropertiesParameters) SetLegacyType(v string) *GetListOfPropertiesParameters {
  s.LegacyType = &v
  return s
}

func (s *GetListOfPropertiesParameters) SetHasConfig(v string) *GetListOfPropertiesParameters {
  s.HasConfig = &v
  return s
}

func (s *GetListOfPropertiesParameters) SetTarget(v string) *GetListOfPropertiesParameters {
  s.Target = &v
  return s
}

func (s *GetListOfPropertiesParameters) SetOffset(v int) *GetListOfPropertiesParameters {
  s.Offset = &v
  return s
}

func (s *GetListOfPropertiesParameters) SetLimit(v int) *GetListOfPropertiesParameters {
  s.Limit = &v
  return s
}

func (s *GetListOfPropertiesParameters) SetSortOrder(v string) *GetListOfPropertiesParameters {
  s.SortOrder = &v
  return s
}

func (s *GetListOfPropertiesParameters) SetSortBy(v string) *GetListOfPropertiesParameters {
  s.SortBy = &v
  return s
}

type GetListOfPropertiesResponse struct {
  // {"en":"List of properties.","zh_CN":"加速项目列表。"}
  Properties []*GetListOfPropertiesResponseProperties `json:"properties,omitempty" xml:"properties,omitempty" require:"true" type:"Repeated"`
  // {"en":"Range: >= 0\nNumber of properties.","zh_CN":"取值范围: >= 0\n加速项目数量。"}
  Count *int `json:"count,omitempty" xml:"count,omitempty" require:"true"`
}

func (s GetListOfPropertiesResponse) String() string {
  return tea.Prettify(s)
}

func (s GetListOfPropertiesResponse) GoString() string {
  return s.String()
}

func (s *GetListOfPropertiesResponse) SetProperties(v []*GetListOfPropertiesResponseProperties) *GetListOfPropertiesResponse {
  s.Properties = v
  return s
}

func (s *GetListOfPropertiesResponse) SetCount(v int) *GetListOfPropertiesResponse {
  s.Count = &v
  return s
}

type GetListOfPropertiesResponseProperties struct     {
  // {"en":"ID of the property.","zh_CN":"加速项目ID。"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {"en":"A description of the property.","zh_CN":"加速项目的描述。"}
  Description *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
  // {"en":"Property name.","zh_CN":"加速项目名称。"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"RFC3339 format date indicating when the property was created.","zh_CN":"RFC 3339格式的日期，表示创建加速项目的时间。"}
  CreationTime *string `json:"creationTime,omitempty" xml:"creationTime,omitempty" require:"true"`
  // {"en":"RFC3339 date indicating when the property was last updated.","zh_CN":"RFC 3339格式的日期，表示加速项目的最近更新时间。"}
  LastUpdateTime *string `json:"lastUpdateTime,omitempty" xml:"lastUpdateTime,omitempty" require:"true"`
  // {"en":"Latest version of the property.","zh_CN":"加速项目的最新版本。"}
  LatestVersion *int `json:"latestVersion,omitempty" xml:"latestVersion,omitempty" require:"true"`
  // {"en":"Enum: wsapro,webpro,vodpro,downloadpro\nService type, one of wsapro, webpro, vodpro , downloadpro , 1523.","zh_CN":"取值范围: wsapro,webpro,vodpro,downloadpro,cdnpro,MALL-cdnpro\n加速项目关联的服务类型，即全站加速，网页加速，点播加速，下载加速，CDN Pro，以及自助CDN。"}
  LegacyType *string `json:"legacyType,omitempty" xml:"legacyType,omitempty" require:"true"`
  // {"en":"Describes the version of the property deployed to staging.","zh_CN":"描述部署到演练环境的加速项目版本。"}
  StagingVersion *GetListOfPropertiesResponsePropertiesStagingVersion `json:"stagingVersion,omitempty" xml:"stagingVersion,omitempty" require:"true" type:"Struct"`
  // {"en":"Describes the version of the property deployed to production.","zh_CN":"描述部署到生产环境的加速项目版本。"}
  ProductionVersion *GetListOfPropertiesResponsePropertiesProductionVersion `json:"productionVersion,omitempty" xml:"productionVersion,omitempty" require:"true" type:"Struct"`
}

func (s GetListOfPropertiesResponseProperties) String() string {
  return tea.Prettify(s)
}

func (s GetListOfPropertiesResponseProperties) GoString() string {
  return s.String()
}

func (s *GetListOfPropertiesResponseProperties) SetId(v string) *GetListOfPropertiesResponseProperties {
  s.Id = &v
  return s
}

func (s *GetListOfPropertiesResponseProperties) SetDescription(v string) *GetListOfPropertiesResponseProperties {
  s.Description = &v
  return s
}

func (s *GetListOfPropertiesResponseProperties) SetName(v string) *GetListOfPropertiesResponseProperties {
  s.Name = &v
  return s
}

func (s *GetListOfPropertiesResponseProperties) SetCreationTime(v string) *GetListOfPropertiesResponseProperties {
  s.CreationTime = &v
  return s
}

func (s *GetListOfPropertiesResponseProperties) SetLastUpdateTime(v string) *GetListOfPropertiesResponseProperties {
  s.LastUpdateTime = &v
  return s
}

func (s *GetListOfPropertiesResponseProperties) SetLatestVersion(v int) *GetListOfPropertiesResponseProperties {
  s.LatestVersion = &v
  return s
}

func (s *GetListOfPropertiesResponseProperties) SetLegacyType(v string) *GetListOfPropertiesResponseProperties {
  s.LegacyType = &v
  return s
}

func (s *GetListOfPropertiesResponseProperties) SetStagingVersion(v *GetListOfPropertiesResponsePropertiesStagingVersion) *GetListOfPropertiesResponseProperties {
  s.StagingVersion = v
  return s
}

func (s *GetListOfPropertiesResponseProperties) SetProductionVersion(v *GetListOfPropertiesResponsePropertiesProductionVersion) *GetListOfPropertiesResponseProperties {
  s.ProductionVersion = v
  return s
}

type GetListOfPropertiesResponsePropertiesStagingVersion struct {
  // {"en":"Property version deployed to staging.","zh_CN":"部署到演练环境的加速项目版本号。"}
  Version *int `json:"version,omitempty" xml:"version,omitempty" require:"true"`
  // {"en":"Hostnames of the property deployed to staging.","zh_CN":"加速域名。"}
  Hostnames []*string `json:"hostnames,omitempty" xml:"hostnames,omitempty" require:"true" type:"Repeated"`
}

func (s GetListOfPropertiesResponsePropertiesStagingVersion) String() string {
  return tea.Prettify(s)
}

func (s GetListOfPropertiesResponsePropertiesStagingVersion) GoString() string {
  return s.String()
}

func (s *GetListOfPropertiesResponsePropertiesStagingVersion) SetVersion(v int) *GetListOfPropertiesResponsePropertiesStagingVersion {
  s.Version = &v
  return s
}

func (s *GetListOfPropertiesResponsePropertiesStagingVersion) SetHostnames(v []*string) *GetListOfPropertiesResponsePropertiesStagingVersion {
  s.Hostnames = v
  return s
}

type GetListOfPropertiesResponsePropertiesProductionVersion struct {
  // {"en":"Property version deployed to production.","zh_CN":"部署到生产环境的加速项目版本号。"}
  Version *int `json:"version,omitempty" xml:"version,omitempty" require:"true"`
  // {"en":"Hostnames of the property deployed to production.","zh_CN":"加速域名。"}
  Hostnames []*string `json:"hostnames,omitempty" xml:"hostnames,omitempty" require:"true" type:"Repeated"`
}

func (s GetListOfPropertiesResponsePropertiesProductionVersion) String() string {
  return tea.Prettify(s)
}

func (s GetListOfPropertiesResponsePropertiesProductionVersion) GoString() string {
  return s.String()
}

func (s *GetListOfPropertiesResponsePropertiesProductionVersion) SetVersion(v int) *GetListOfPropertiesResponsePropertiesProductionVersion {
  s.Version = &v
  return s
}

func (s *GetListOfPropertiesResponsePropertiesProductionVersion) SetHostnames(v []*string) *GetListOfPropertiesResponsePropertiesProductionVersion {
  s.Hostnames = v
  return s
}

type GetListOfPropertiesResponseHeader struct {
}

func (s GetListOfPropertiesResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetListOfPropertiesResponseHeader) GoString() string {
  return s.String()
}




type CreateAPropertyVersionRequest struct {
  // {"en":"A description of the version.","zh_CN":"版本描述。"}
  Description *string `json:"description,omitempty" xml:"description,omitempty"`
  // {"en":"Hostnames associated with the property. A valid value is a fully qualified hostname such as www.domain.com or a wildcard hostname such as.domain.com. Any given hostname can only be in one deployed property at a particular time. However, a wildcard hostname is permitted to overlap other hostnames you own.","zh_CN":"与加速项目关联的加速域名。必须是FQDN完全限定域名（如 www.domain.com），或泛域名（如*.domain.com）。\n同一个加速域名在同一时间只能存在于一个已部署的加速项目中，但泛域名可以与关联的完全限定域名一同部署。"}
  Hostnames []*string `json:"hostnames,omitempty" xml:"hostnames,omitempty" require:"true" type:"Repeated"`
  // {"en":"Default: False\nThe value indicates whether all the hostnames in this property have Beian license on file with the Chinese government. This is required to serve this property from servers in mainland China. A value of false means servers outside of mainland China will be used to distribute content to visitors in China. If set to true you must also specify the content type in the beianContentType field.","zh_CN":"默认值: False\n此加速项目中的所有加速域名是否都已获得ICP备案。只有获取备案的域名才能使用中国大陆节点来提供服务。如果设置为true，还必须在beianContentType字段中指定内容类型。"}
  HasBeian *bool `json:"hasBeian,omitempty" xml:"hasBeian,omitempty"`
  // {"en":"Range: [ 1 .. 24 ]\nIf you are planning to distribute content through servers in mainland China, you will need to set the hasBeian field to true and also indicate the type of content you are distributing. Enter the value that best corresponds to your content:\n1. Instant Messaging\n2. Search Engine\n3. Web Portal\n4. Online Postal Service\n5. Online News\n6. Blog\n7. Advertisement\n8. Organization Web Portal\n9. Shopping\n10. Online Payment\n11. Online Bank\n12. Online Stock Market\n13. Online Gaming\n14. Online Music\n15. Online Movie\n16. Online Picture\n17. Software Download\n18. Job Hunting\n19. Online Dating\n20. Online Real Property\n21. Online Education\n22. Web Design\n23. WAP\n24. Others","zh_CN":"取值范围: [ 1 .. 24 ]\n如果您的域名已获得备案，希望通过中国大陆的服务器分发内容，则需要将hasBeian字段设置为true，并说明您分发的内容类型。请选择与您的内容最匹配的值：\n1. 即时通信\n2. 搜索引擎\n3. 综合门户\n4. 网上邮局\n5. 网络新闻\n6. 博客/个人空间\n7. 网络广告\n8. 单位门户网站\n9. 网络购物\n10. 网上支付\n11. 网上银行\n12. 网上炒股\n13. 网络游戏\n14. 网络音乐\n15. 网络影视\n16. 网络图片\n17. 软件下载\n18. 网上求职\n19. 在线交友\n20. 网上房产\n21. 网络教育\n22. 网站建设\n23. WAP\n24. 其他"}
  BeianContentType *int `json:"beianContentType,omitempty" xml:"beianContentType,omitempty"`
  // {"en":"Describes the origin servers for the property's content.","zh_CN":"描述加速项目对应的源站。"}
  Origins []*CreateAPropertyVersionRequestOrigins `json:"origins,omitempty" xml:"origins,omitempty" type:"Repeated"`
  // {"en":"Default: False\nWhether to allow downgrading from HTTPS to HTTP when fetching content from origin. This setting applies only if the property has an attached certificate allowing client requests to use HTTPS. If the value of allowProtocolDowngrade is false and HTTPS is enabled, we require all the origin servers to support HTTPS. If the value is true, we allow origins that support only HTTP, but this reduces security level.","zh_CN":"默认值: False\n是否允许回源时从HTTPS降级为HTTP。仅当加速项目中指定了用于HTTPS请求的证书时，此设置才适用。如果allowProtocolDowngrade的值为false，且加速项目开启了HTTPS，则要求所有源服务器支持HTTPS。如果值为true，则允许仅支持HTTP的源，但这会降低安全性。"}
  AllowProtocolDowngrade *bool `json:"allowProtocolDowngrade,omitempty" xml:"allowProtocolDowngrade,omitempty"`
  // {"en":"Default: False\nEnables or disables using IPv6 addresses as origin.","zh_CN":"默认值: False\n是否开启IPv6回源"}
  EnableIpv6Origin *bool `json:"enableIpv6Origin,omitempty" xml:"enableIpv6Origin,omitempty"`
  // {"en":"Enum: auto, strict, off\nDefault: off\n\nWhen enableIpv6Origin is allowed, this setting will control whether to follow the client IP protocol version for back-to-origin requests.\nAuto: Attempt to reach the origin server using the same IP addressing as the client, but use the other version as backup if needed.\nStrict: Use only the same addressing method as the client when accessing the origin for content.\nOff: Use both IPv4 and IPv6 addressing to reach the origin, regardless of the IP addressing method used by the client.","zh_CN":"取值范围：auto, strict, off\n默认值：off\n当enableIpv6Origin为true时，该设置将控制是否跟随客户端IP协议版本回源。\nauto：表示根据客户端请求的IP协议版本来选择回源IP地址，当源站服务器不可用时自动切换成其他IP协议地址回源\nstrict：严格根据客户端请求的IP协议版本来选择回源IP地址，当源站服务器不可用时不允许切换成其他协议地址回源\noff：忽略客户端请求的IP协议版本，随机选择可用的IPv4或IPv6地址回源"}
  FollowClientIpVersion *string `json:"followClientIpVersion,omitempty" xml:"followClientIpVersion,omitempty"`
  // {"en":"Range: <= 65530 characters\nRefer to Edge Logic Introduction.","zh_CN":"取值范围: <= 65530 字符\n自定义边缘逻辑。参考边缘逻辑介绍。"}
  EdgeLogic *string `json:"edgeLogic,omitempty" xml:"edgeLogic,omitempty" require:"true"`
  // {"en":"Specify the syntax version that should be used when parsing the configuration. The value must be set to 1 at this time.","zh_CN":"指定配置解析适用的语法版本。当前仅允许值为1。"}
  SyntaxVersion *int `json:"syntaxVersion,omitempty" xml:"syntaxVersion,omitempty" require:"true"`
  // {"en":"Range: <= 80\nThe cachekey hostname must be a string made of lowercase letters, digits, dot, dash, and underscore. If not specified, the incoming Host header value will be used and the cache will not be shared across different hostnames in this property.","zh_CN":"取值范围: <= 80\n用于缓存键（Cache Key）的主机名。必须是由小写字母、数字、点、连字符和下划线组成的字符串。如果未指定，默认将使用客户端传入的Host请求头值，并且此加速项目中的不同加速域名之间不共享缓存。"}
  CacheKeyHostname *string `json:"cacheKeyHostname,omitempty" xml:"cacheKeyHostname,omitempty"`
  // {"en":"Enum: preRewrite,postRewrite\nDefault: preRewrite\nThis setting controls how the URI of the incoming request is incorporated into the cache key if you use a 'rewrite' directive in the property's Edge Logic. The default value, 'preRewrite', puts the unmodified URI into the cache key while 'postRewrite' uses the modified URI. If your rewrite directive converts multiple different URIs into the same value, using 'postRewrite' may result in a higher cache hit ratio.","zh_CN":"取值范围: preRewrite,postRewrite\n默认值: preRewrite\n如果在加速项目的边缘逻辑中使用了'rewrite'指令，可使用该字段来控制客户端请求的URI如何合并到缓存键（Cache Key）中 。默认值'preRewrite'指将改写前的URI放入缓存键，而'postRewrite'则使用改写后的URI。如果您的'rewrite'指令将多个不同的URI改写为相同的值，则使用'postRewrite'可以提高缓存命中率。"}
  CacheKeyUri *string `json:"cacheKeyUri,omitempty" xml:"cacheKeyUri,omitempty"`
  // {"en":"Default: False\nSpecifies whether the scheme ('http' or 'https') should be included in the cache key. Default behavior is false. Set to true if the content is known to be different for different schemes.","zh_CN":"默认值: False\n指定缓存键（Cache Key）是否区分协议（'http'或'https'）。默认为false。如果不同协议响应的内容不同，则设置为true。"}
  SchemeInCacheKey *bool `json:"schemeInCacheKey,omitempty" xml:"schemeInCacheKey,omitempty"`
  // {"en":"Default: False\nThis field can be set to either a boolean value or a string. If the value is set to true, the server will redirect all HTTP requests to HTTPS using status code 301. You can also specify string values '302', '307', or '308' instead if you wish to return a different status code when a request is redirected.","zh_CN":"默认值: False\n此字段可以设置为布尔值或字符串。如果设置为true，则CDN Pro服务器会将所有HTTP请求重定向到HTTPS，并返回301状态码。如果您希望在重定向请求时返回不同的状态码，可在此处指定需要的状态码，如'302'、'307'或'308'。"}
  RedirectHttpToHttps *string `json:"redirectHttpToHttps,omitempty" xml:"redirectHttpToHttps,omitempty"`
  // {"en":"Default: False\nSet to true to disable support for HTTP/2 and support only HTTP 1.1.","zh_CN":"默认值: False\n当值为true时，禁用对HTTP/2的支持，仅支持HTTP/1.1。"}
  DisableHttp2 *bool `json:"disableHttp2,omitempty" xml:"disableHttp2,omitempty"`
  // {"en":"Default: False\nSet to true to enable support for HTTP/3. HTTP/3 requires TLS to work, so this field will be ignored unless you specify a certificate in the tlsCertificateId field.","zh_CN":"默认值: False\n是否开启HTTP/3。需要配置TLS才生效，若您未在 tlsCertificateId中指定证书，此字段将被忽略。"}
  EnableHttp3 *bool `json:"enableHttp3,omitempty" xml:"enableHttp3,omitempty"`
  // {"en":"This field lists ports other than the default 80 used to handle HTTP requests and ports other than the default 443 used to handle HTTPS requests.","zh_CN":"除标准的80，443端口外，我们还支持一些扩展端口。可用该字段指定用于处理HTTP和HTTPS请求的扩展端口。"}
  ExtraServicePorts *CreateAPropertyVersionRequestExtraServicePorts `json:"extraServicePorts,omitempty" xml:"extraServicePorts,omitempty" type:"Struct"`
  // {"en":"Specify a certificate to enable HTTPS for your hostnames.","zh_CN":"指定用于HTTPS请求的证书的ID。如果未指定证书，则不会为此加速项目启用HTTPS。"}
  TlsCertificateId *string `json:"tlsCertificateId,omitempty" xml:"tlsCertificateId,omitempty"`
  // {"en":"Specify a second certificate to be used for HTTPS. With tlsCertificateId and tlsCertificateId1, you can specify two certificates of different types, i.e. one RSA and one EC. When there are two certificates available for your hostnames, the client and server will negotiate at the time of requests to pick one of the certificates for SSL. This could mean better performance and fault tolerance.\n\nIt is invalid to only set the tlsCertificateId1 field without setting the tlsCertificateId field. The certificates set for tlsCertificateId and tlsCertificateId1 must be of different type.","zh_CN":"指定用于HTTPS请求的第二本证书。必须先指定tlsCertificateId，才能指定tlsCertificateId1。tlsCertificateId和tlsCertificateId1所指定的证书必须是不同证书类型。\n\ntlsCertificateId和tlsCertificateId1允许您指定两本不同类型的证书，即一本RSA证书，一本EC证书。如果您指定了多本证书，当客户端发起请求时，客户端和服务端可以协商选择使用其中一本证书。"}
  TlsCertificateId1 *string `json:"tlsCertificateId1,omitempty" xml:"tlsCertificateId1,omitempty"`
  // {"en":"Enum: 1.1,1.2,1.3\nDefault: 1.3\nMaximum supported TLS version.","zh_CN":"取值范围: 1.1,1.2,1.3\n默认值: 1.3\n支持的TLS最高版本。"}
  TlsMaxVersion *string `json:"tlsMaxVersion,omitempty" xml:"tlsMaxVersion,omitempty"`
  // {"en":"Enum: 1.1,1.2,1.3,1\nDefault: 1\nMinimum required TLS version.","zh_CN":"取值范围: 1.1,1.2,1.3,1\n默认值: 1\n所需的TLS最低版本。"}
  TlsMinVersion *string `json:"tlsMinVersion,omitempty" xml:"tlsMinVersion,omitempty"`
  // {"en":"Range: <= 2040 characters\nAny cipher list supported by 'openssl ciphers'. See https://www.openssl.org/docs/manmaster/man1/ciphers.html","zh_CN":"取值范围: <= 2040 字符\n'openssl ciphers'支持的任何加密算法套件。参考：https://www.openssl.org/docs/manmaster/man1/ciphers.html"}
  TlsCiphers *string `json:"tlsCiphers,omitempty" xml:"tlsCiphers,omitempty"`
  // {"en":"Default: 1800 Range: [ 300 .. 86400 ]\nLifespan of TLS session ticket in seconds.","zh_CN":"默认值: 1800 取值范围: [ 300 .. 86400 ]\nTLS会话ticket的有效期（秒）。"}
  TlsSessionTimeout *int `json:"tlsSessionTimeout,omitempty" xml:"tlsSessionTimeout,omitempty"`
  // {"en":"Default: False\nEnable TLS zero round-trip time, a feature of TLS 1.3 that can improve performance for repeat visits to a site. If enabling it though, be sure your site is not vulnerable to the HTTP replay attack.","zh_CN":"默认值: False\n是否开启TLS 0-RTT功能。这是TLS 1.3版本支持的功能。启用该功能后，当用户频繁访问您的站点时，可提高访问响应速度。需要注意的是，启用该功能可能会加大站点遭受HTTP replay攻击的风险。"}
  Tls0Rtt *bool `json:"tls0Rtt,omitempty" xml:"tls0Rtt,omitempty"`
  // {"en":"Default: False\nEnables Online Certificate Status Protocol (OCSP) stapling to check the revocation status of digital certificates. (Refer to https://en.wikipedia.org/wiki/OCSP_stapling)","zh_CN":"默认值: False\n是否启用在线证书状态协议（OCSP）装订以检查数字证书的吊销状态。参考：https://en.wikipedia.org/wiki/OCSP_stapling"}
  EnableOcspStapling *bool `json:"enableOcspStapling,omitempty" xml:"enableOcspStapling,omitempty"`
  // {"en":"Default: False\nBy default, CDN Pro takes control of the contents under the /.well-known/{acme-challenge, pki-validation} directories to support certificate auto-renew for properties. If for any reason you need to manage these two directories by yourself on the origin, for example, to implement your own certificate auto-renew mechanism, you can use this configuration option to disable the default behavior by setting its value to true. For more information about our support for certificate auto-renewal, refer to the description of the autoRenew field in the Create a certificate API.","zh_CN":"默认值: False\n默认情况下，CDN Pro控制/.well-known/{acme-challenge, pki-validation}目录下的内容，以支持加速项目的证书自动更新功能。如果您需要自己在源站管理这两个目录，例如，为了实现您自己的证书自动更新机制，您可以将此字段设置为true来禁用默认行为。关于证书自动更新的更多信息，请参考'创建证书'接口中autoRenew字段的说明。"}
  DisableCertAutomation *bool `json:"disableCertAutomation,omitempty" xml:"disableCertAutomation,omitempty"`
  // {"en":"Specify one or more access control rules to restrict access to your content. More advanced configuration can be done using Edge Logic. These access control rules take precedence over Edge Logic if both are defined.","zh_CN":"指定一个或多个访问控制规则以限制对内容的访问。可以使用边缘逻辑进行更高级的配置。此处定义的访问控制规则，优先级高于边缘逻辑。"}
  AccessControlRules []*CreateAPropertyVersionRequestAccessControlRules `json:"accessControlRules,omitempty" xml:"accessControlRules,omitempty" type:"Repeated"`
  // {"en":"This optional field allows you to configure notifications about client requests to be sent to a remote server. It can be used only if you have access to our realtime_log_switch directive. Please contact our support team if you require this feature.","zh_CN":"此可选字段用来配置发送消息通知（即实时日志）到您的远程服务器。当有客户端请求访问您的加速域名时，将触发通知。这是高级功能，如果您需要此功能，请联系我们的技术支持开通。"}
  RealTimeLog *CreateAPropertyVersionRequestRealTimeLog `json:"realTimeLog,omitempty" xml:"realTimeLog,omitempty" type:"Struct"`
  // {"en":"This object allows you to support video players requesting partial content through query string parameters. If you specify videoSeek, you must enter a value for startParameter.","zh_CN":"此对象用来支持视频播放器通过指定查询参数来请求部分内容。当videoSeek对象存在时，必须为startParameter设置一个值。"}
  VideoSeek *CreateAPropertyVersionRequestVideoSeek `json:"videoSeek,omitempty" xml:"videoSeek,omitempty" type:"Struct"`
  // {"en":"Range: <= 65530 characters\nThis field allows you to customize load balancing. The field will be deprecated soon. Please use the field edgeLogic.","zh_CN":"取值范围: <= 65530 字符\n用于控制负载均衡器的行为。该字段即将废弃，请使用edgeLogic字段进行配置。"}
  LoadBalancerLogic *string `json:"loadBalancerLogic,omitempty" xml:"loadBalancerLogic,omitempty"`
  // {"en":"Range: <= 100 characters\nUsed to customize hash key for load balancer. The field will soon be deprecated. Please stop using it.","zh_CN":"取值范围: <= 100 字符\n用于自定义负载均衡器的哈希key。该字段即将废弃，请勿使用该字段。"}
  LoadBalancerHashKey *string `json:"loadBalancerHashKey,omitempty" xml:"loadBalancerHashKey,omitempty"`
}

func (s CreateAPropertyVersionRequest) String() string {
  return tea.Prettify(s)
}

func (s CreateAPropertyVersionRequest) GoString() string {
  return s.String()
}

func (s *CreateAPropertyVersionRequest) SetDescription(v string) *CreateAPropertyVersionRequest {
  s.Description = &v
  return s
}

func (s *CreateAPropertyVersionRequest) SetHostnames(v []*string) *CreateAPropertyVersionRequest {
  s.Hostnames = v
  return s
}

func (s *CreateAPropertyVersionRequest) SetHasBeian(v bool) *CreateAPropertyVersionRequest {
  s.HasBeian = &v
  return s
}

func (s *CreateAPropertyVersionRequest) SetBeianContentType(v int) *CreateAPropertyVersionRequest {
  s.BeianContentType = &v
  return s
}

func (s *CreateAPropertyVersionRequest) SetOrigins(v []*CreateAPropertyVersionRequestOrigins) *CreateAPropertyVersionRequest {
  s.Origins = v
  return s
}

func (s *CreateAPropertyVersionRequest) SetAllowProtocolDowngrade(v bool) *CreateAPropertyVersionRequest {
  s.AllowProtocolDowngrade = &v
  return s
}

func (s *CreateAPropertyVersionRequest) SetEnableIpv6Origin(v bool) *CreateAPropertyVersionRequest {
  s.EnableIpv6Origin = &v
  return s
}

func (s *CreateAPropertyVersionRequest) SetFollowClientIpVersion(v string) *CreateAPropertyVersionRequest {
  s.FollowClientIpVersion = &v
  return s
}

func (s *CreateAPropertyVersionRequest) SetEdgeLogic(v string) *CreateAPropertyVersionRequest {
  s.EdgeLogic = &v
  return s
}

func (s *CreateAPropertyVersionRequest) SetSyntaxVersion(v int) *CreateAPropertyVersionRequest {
  s.SyntaxVersion = &v
  return s
}

func (s *CreateAPropertyVersionRequest) SetCacheKeyHostname(v string) *CreateAPropertyVersionRequest {
  s.CacheKeyHostname = &v
  return s
}

func (s *CreateAPropertyVersionRequest) SetCacheKeyUri(v string) *CreateAPropertyVersionRequest {
  s.CacheKeyUri = &v
  return s
}

func (s *CreateAPropertyVersionRequest) SetSchemeInCacheKey(v bool) *CreateAPropertyVersionRequest {
  s.SchemeInCacheKey = &v
  return s
}

func (s *CreateAPropertyVersionRequest) SetRedirectHttpToHttps(v string) *CreateAPropertyVersionRequest {
  s.RedirectHttpToHttps = &v
  return s
}

func (s *CreateAPropertyVersionRequest) SetDisableHttp2(v bool) *CreateAPropertyVersionRequest {
  s.DisableHttp2 = &v
  return s
}

func (s *CreateAPropertyVersionRequest) SetEnableHttp3(v bool) *CreateAPropertyVersionRequest {
  s.EnableHttp3 = &v
  return s
}

func (s *CreateAPropertyVersionRequest) SetExtraServicePorts(v *CreateAPropertyVersionRequestExtraServicePorts) *CreateAPropertyVersionRequest {
  s.ExtraServicePorts = v
  return s
}

func (s *CreateAPropertyVersionRequest) SetTlsCertificateId(v string) *CreateAPropertyVersionRequest {
  s.TlsCertificateId = &v
  return s
}

func (s *CreateAPropertyVersionRequest) SetTlsCertificateId1(v string) *CreateAPropertyVersionRequest {
  s.TlsCertificateId1 = &v
  return s
}

func (s *CreateAPropertyVersionRequest) SetTlsMaxVersion(v string) *CreateAPropertyVersionRequest {
  s.TlsMaxVersion = &v
  return s
}

func (s *CreateAPropertyVersionRequest) SetTlsMinVersion(v string) *CreateAPropertyVersionRequest {
  s.TlsMinVersion = &v
  return s
}

func (s *CreateAPropertyVersionRequest) SetTlsCiphers(v string) *CreateAPropertyVersionRequest {
  s.TlsCiphers = &v
  return s
}

func (s *CreateAPropertyVersionRequest) SetTlsSessionTimeout(v int) *CreateAPropertyVersionRequest {
  s.TlsSessionTimeout = &v
  return s
}

func (s *CreateAPropertyVersionRequest) SetTls0Rtt(v bool) *CreateAPropertyVersionRequest {
  s.Tls0Rtt = &v
  return s
}

func (s *CreateAPropertyVersionRequest) SetEnableOcspStapling(v bool) *CreateAPropertyVersionRequest {
  s.EnableOcspStapling = &v
  return s
}

func (s *CreateAPropertyVersionRequest) SetDisableCertAutomation(v bool) *CreateAPropertyVersionRequest {
  s.DisableCertAutomation = &v
  return s
}

func (s *CreateAPropertyVersionRequest) SetAccessControlRules(v []*CreateAPropertyVersionRequestAccessControlRules) *CreateAPropertyVersionRequest {
  s.AccessControlRules = v
  return s
}

func (s *CreateAPropertyVersionRequest) SetRealTimeLog(v *CreateAPropertyVersionRequestRealTimeLog) *CreateAPropertyVersionRequest {
  s.RealTimeLog = v
  return s
}

func (s *CreateAPropertyVersionRequest) SetVideoSeek(v *CreateAPropertyVersionRequestVideoSeek) *CreateAPropertyVersionRequest {
  s.VideoSeek = v
  return s
}

func (s *CreateAPropertyVersionRequest) SetLoadBalancerLogic(v string) *CreateAPropertyVersionRequest {
  s.LoadBalancerLogic = &v
  return s
}

func (s *CreateAPropertyVersionRequest) SetLoadBalancerHashKey(v string) *CreateAPropertyVersionRequest {
  s.LoadBalancerHashKey = &v
  return s
}

type CreateAPropertyVersionRequestOrigins struct     {
  // {"en":"^[a-zA-z0-9_]\nName of an origin. It must be unique within this property.","zh_CN":"^[a-zA-z0-9_]\n源站名称，在此加速项目中必须唯一。"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"Each entry should be a string consisting of an address optionally followed by one or more parameters, separated by space characters. The address can take one of the following three forms:\n{hostname or IP address}\n{hostname or IP address}:{http port}\n{hostname or IP address}:{http port},{https port}\nValues for the HTTP and HTTPS ports should be integers in the range [1,65535].\nIn the third form, one of the two ports can be empty, but the HTTPS port must always be specified after the comma.\n\nAccepts ipv6 addresses if 'enableIpv6Origin' is set to true, but an ipv6 address must be enclosed in [] so as to tell the port apart.\n\nSupported parameters are 'backup' and 'weight'.\n\n'backup' is used to indicate the server is a backup server. It will be used when the primary servers are unavailable.\n\n'weight' is a value in the range [1, 100]. The default value is 1. It lets you control the likelihood that one origin server will be used relative to another.\n\nThere should be at least one primary server defined; it does not have the 'backup' parameter.\n\nExample:\n['www.abc.com weight=1', 'www1.abc.com:8080 weight=4', 'www2.abc.com:880,4443 backup']\n\nThis example specifies two primary servers which are used in a 1:4 ratio meaning one gets about 20% of the requests to origin while the other gets 80% of the requests. In addition, a backup server is specified.","zh_CN":"源站服务器列表。每个条目为一个字符串，由一个地址与一个或多个参数组成，参数之间以空格分隔。地址可以采用以下三种形式之一：\n{域名或IP地址}\n{域名或IP地址}:{HTTP端口}\n{域名或IP地址}:{HTTP端口},{HTTPS端口}\nHTTP和HTTPS端口的值应为[1,65535]范围内的整数。\n在第三种形式中，HTTPS端口必须固定写在逗号后。如果HTTP端口未指定，仍需保留逗号，写为“,{HTTPS端口}”。\n\n当enableIpv6Origin设为true时，此处支持指定ipv6地址。ipv6地址应包裹在[]中，以便和端口区分开，例如[2001:0db8:85a3:0000:0000:8a2e:0370:7334]:80\n\n支持的参数包括'backup'和'weight'。\n'backup'用于标识备份服务器。它将在主服务器不可用时使用。\n'weight'是范围[1,100]内的值，默认值为1，用来设置服务器权重，控制一台源站服务器相对于另一台被使用的可能性。\n每个源站应至少定义一个主服务器（即不带'backup'参数的服务器）。\n\n示例：\n['www.abc.com weight=1', 'www1.abc.com:8080 weight=4', 'www2.abc.com:880,4443 backup']\n\n此示例指定了两个主服务器和一个备份服务器，其中，两个主服务器的权重为1:4，表示第一个服务器将获得约20%的回源请求，而另一个将获得约80%。"}
  Servers []*string `json:"servers,omitempty" xml:"servers,omitempty" require:"true" type:"Repeated"`
  // {"en":"Enum: http,https,both\nThis optional field indicates the protocol supported by the origin server. If this property has a certificate attached to it, the value can be set to http if the allowProtocolDowngrade field is also set to true.","zh_CN":"取值范围: http,https,both\n此可选字段表示源站服务器支持的协议。如果此加速项目附加了证书，且allowProtocolDowngrade字段也设置为true，则可以将该值设置为http。"}
  SupportedProtocol *string `json:"supportedProtocol,omitempty" xml:"supportedProtocol,omitempty"`
  // {"en":"Default: auto\nOptional. This parameter tells us how important it is to directly go to this origin without going through any intermediate cache to fetch content. The values can be 'noDirect', 'auto', 'alwaysDirect'.\n\n'noDirect' means we always go through at least one intermediate cache. Customers who care particularly about the origin's workload can use this option. Our cache scheduler will dynamically pick the intermediate cache based on performance and resource availability.\n\n'auto' means the cache scheduler will make the choice dynamically based on cacheability of content. This is the default behavior if not specified.\n\n'alwaysDirect' means we always directly connect to this origin.","zh_CN":"默认值: auto\n此可选字段用来指定回源方式，可以是'noDirect'、'auto'、'alwaysDirect'。\n\n'noDirect'指总是通过至少一个中间缓存节点回源。对于特别关心源站负载的客户可以使用此选项。我们的调度程序将根据性能和资源可用性动态选择中间缓存节点。\n'alwaysDirect'指总是直接回源。\n'auto'指自动根据内容是否可缓存选择直接回源或通过中间缓存节点回源。对于可缓存的内容，将通过中间缓存节点回源；对于不可缓存的内容，将直接回源。该字段未指定时，采用此默认行为。"}
  DirectConnection *string `json:"directConnection,omitempty" xml:"directConnection,omitempty"`
  // {"en":"Optional. It should be a valid hostname. It will also be used as the SNI servername when contacting the HTTPS origin. We also allow it to be an nginx variable that begins with '$' followed by [a-z_]{3,60}. Nginx variable names are case insensitive, so we only allow lowercase characters.\nIf not specified, the value of the HOST header in the request will be used.\n\nOne constraint is that if 'aswS3' is specified for origin authentication, the value of hostHeader cannot be a variable or empty. It has to be a valid Amazon S3 hostname. Refer to https://docs.aws.amazon.com/AmazonS3/latest/dev/VirtualHosting.html","zh_CN":"可选，用来指定回源HOST头部。必须是有效的域名。当连接HTTPS源站时，该值也作为SNI域名。可以用nginx变量来指定，变量以'$'开头，后跟[a-z_]{3,60}。Nginx变量名不区分大小写，因此我们只允许小写字符。\n如果未指定，将使用请求中的HOST请求头的值。\n\n注意：如果指定了'awsS3'作为回源鉴权参数，则hostHeader的值不能为变量或为空，而必须是有效的Amazon S3主机名。参考：https://docs.aws.amazon.com/AmazonS3/latest/dev/VirtualHosting.html"}
  HostHeader *string `json:"hostHeader,omitempty" xml:"hostHeader,omitempty"`
  // {"en":"Default: False\nOptional. It controls whether the cache will validate the certificate of the origin.","zh_CN":"默认值: False\n可选。用来设定CDN缓存节点是否验证源站证书。"}
  VerifyOrigin *bool `json:"verifyOrigin,omitempty" xml:"verifyOrigin,omitempty"`
  // {"en":"Optional. It is a structure to specify the parameters, such as password, for authentication with the origin servers. It should have at least one field: 'methodName', which should be a string indicating one of the predefined authentication methods. The only valid value at this time is 'awsS3'. When 'awsS3' is used, the following parameters are required:\nawsS3.region:\nThe AWS S3 region where your resources are hosted, e.g. 'us-west-1', 'ap-northeast-1', 'eu-north-1', and 'cn-north-1'.\nawsS3.accessKey:\nYour AWS access key that CDN Pro will use to access your resources stored on AWS S3.\nawsS3.secretKey:\nYour AWS secret key that CDN Pro will use to access your resources stored on AWS S3.\nAlso, the value of the hostHeader field cannot be a variable or empty. It must be a valid Amazon S3 hostname.\n\nExample:\n<code>{\"methodName\":\"awsS3\",\n\"awsS3\":{\n\"region\":\"us-west-1\",\n\"accessKey\":\"sdnu2932\",\n\"secretKey\":\"d12345678abcdefghi\"\n}}</code>","zh_CN":"可选。用于指定与源服务器进行身份验证（回源鉴权）的相关参数（如密码）。必须至少有'methodName'（字符串）字段，用来指定预定义的鉴权方法。目前仅支持源站为AWS S3的回源鉴权，因此唯一有效的值是'awsS3'。使用'awsS3'时，需要提供以下参数：\nawsS3.region:\n您存储在S3上的资源所在的AWS区域。例如，'us-west-1'，'ap-northeast-1'，'eu-north-1'，'cn-north-1'。\nawsS3.accessKey:\n您的 AWS 访问密钥（access key），CDN Pro 将使用它来访问您存储在 AWS S3 上的资源。\nawsS3.secretKey:\n您的 AWS 密钥（secret key），CDN Pro 将使用它来访问您存储在 AWS S3 上的资源。\n此外，hostHeader字段的值不能为变量或为空，必须是有效的Amazon S3主机名。\n\n示例：\n<code>{\"methodName\":\"awsS3\",\n\"awsS3\":{\n\"region\":\"us-west-1\",\n\"accessKey\":\"sdnu2932\",\n\"secretKey\":\"d12345678abcdefghi\"\n}}</code>"}
  Authentication *CreateAPropertyVersionRequestOriginsAuthentication `json:"authentication,omitempty" xml:"authentication,omitempty" type:"Struct"`
  // {"en":"Default: 60 Range: [ 5 .. 600 ]\nTimeout in seconds during which an idle keepalive connection to an upstream server will stay open. A service quota setting of maxUpstreamKeepaliveTimeOut can change the maximum permitted value.","zh_CN":"默认值: 60 取值范围: [ 5 .. 600 ]\n该字段用于指定CDN Pro服务器和源站建连的Keep-Alive超时时间，单位为秒。通过maxUpstreamKeepaliveTimeOut 该服务设置项可以更改允许的最大值。如果需要调整最大值，请联系我们的技术支持。"}
  KeepAliveTimeout *int `json:"keepAliveTimeout,omitempty" xml:"keepAliveTimeout,omitempty"`
  // {"en":"This setting allows you to specify the number of unsuccessful attempts to reach one of the origin's IP addresses or peers before it is marked as unavailable for a period of time designated by the timeout. If all peers of all servers are unavailable, requests for content from the origin will receive an immediate 502 HTTP response. By default, six attempts to a peer are made, after which a one-second timeout applies to that peer. To disable this feature, set peerFailureTimeout to 'off'.\nExample: <code>{\"peerFailureTimeout\":{\"failureThreshold\":10,\"timeout\":3}}</code>","zh_CN":"该字段用于设置源站故障剔除。开启此功能后，当连接某个源站服务器的失败次数达到设定阈值，该源站服务器将被标记为不可用，并保持该状态直到设定的超时时长。失败次数阈值和超时时长分别通过failureThreshold参数和timeout参数设置。CDN Pro回源时将剔除不可用的源站服务器。如果所有源站服务器都被标记为不可用，则对应的请求将立即响应502状态码。默认情况下，当连接某个源站服务器连续失败6次之后，该服务器将被标记为不可用，超时时长为1秒。如果要禁用此功能，请将peerFailureTimeout设置为'off'。开启源站故障剔除配置示例：<code>{\"peerFailureTimeout\":{\"failureThreshold\":10,\"timeout\":3}}</code>"}
  PeerFailureTimeout *string `json:"peerFailureTimeout,omitempty" xml:"peerFailureTimeout,omitempty"`
  // {"en":"Refers to a certificate used to authenticate with the origin server. This certificate must also be deployed.","zh_CN":"用于源站验证CDN Pro服务器的证书，该证书同样必须被部署。"}
  TlsCertificateId *string `json:"tlsCertificateId,omitempty" xml:"tlsCertificateId,omitempty"`
  // {"en":"Refers to the ID of an origin shield representing a set of servers that will make requests to your origin servers rather than the CDN Pro edge servers. This can help further reduce traffic to your origin. Origin shield is allowed only when directConnection is set to noDirect. This is an advanced feature that can be enabled by contacting our support team. The origin shields API provides a list of available shields along with their IP addresses. It is best to select a shield whose region is closest to your origin servers. Use of a shield in China requires your property have 'hasBeian' set to true. If your origin servers have an access control list, add the IPs of the shield you choose to use.","zh_CN":"指定某个源站盾（origin shield）的ID。使用源站盾功能后，所有回源请求都会通过源站盾中转回源，这可以帮助收敛回源流量。需要把directConnection设置为noDirect，才允许开启源站盾。源站盾是高级功能，如需使用请联系我们的技术支持开通。可通过调用‘获取源站盾列表’接口查询可用的源站盾及其对应的IP地址。您可根据源站的位置，选择合适的源站盾。如果您需使用中国大陆的源站盾，则该加速项目的所有域名必须已取得备案。如果您的源站开启了访问控制，请将所选择源站盾的IP地址加入白名单。"}
  Shield *string `json:"shield,omitempty" xml:"shield,omitempty"`
}

func (s CreateAPropertyVersionRequestOrigins) String() string {
  return tea.Prettify(s)
}

func (s CreateAPropertyVersionRequestOrigins) GoString() string {
  return s.String()
}

func (s *CreateAPropertyVersionRequestOrigins) SetName(v string) *CreateAPropertyVersionRequestOrigins {
  s.Name = &v
  return s
}

func (s *CreateAPropertyVersionRequestOrigins) SetServers(v []*string) *CreateAPropertyVersionRequestOrigins {
  s.Servers = v
  return s
}

func (s *CreateAPropertyVersionRequestOrigins) SetSupportedProtocol(v string) *CreateAPropertyVersionRequestOrigins {
  s.SupportedProtocol = &v
  return s
}

func (s *CreateAPropertyVersionRequestOrigins) SetDirectConnection(v string) *CreateAPropertyVersionRequestOrigins {
  s.DirectConnection = &v
  return s
}

func (s *CreateAPropertyVersionRequestOrigins) SetHostHeader(v string) *CreateAPropertyVersionRequestOrigins {
  s.HostHeader = &v
  return s
}

func (s *CreateAPropertyVersionRequestOrigins) SetVerifyOrigin(v bool) *CreateAPropertyVersionRequestOrigins {
  s.VerifyOrigin = &v
  return s
}

func (s *CreateAPropertyVersionRequestOrigins) SetAuthentication(v *CreateAPropertyVersionRequestOriginsAuthentication) *CreateAPropertyVersionRequestOrigins {
  s.Authentication = v
  return s
}

func (s *CreateAPropertyVersionRequestOrigins) SetKeepAliveTimeout(v int) *CreateAPropertyVersionRequestOrigins {
  s.KeepAliveTimeout = &v
  return s
}

func (s *CreateAPropertyVersionRequestOrigins) SetPeerFailureTimeout(v string) *CreateAPropertyVersionRequestOrigins {
  s.PeerFailureTimeout = &v
  return s
}

func (s *CreateAPropertyVersionRequestOrigins) SetTlsCertificateId(v string) *CreateAPropertyVersionRequestOrigins {
  s.TlsCertificateId = &v
  return s
}

func (s *CreateAPropertyVersionRequestOrigins) SetShield(v string) *CreateAPropertyVersionRequestOrigins {
  s.Shield = &v
  return s
}

type CreateAPropertyVersionRequestOriginsAuthentication struct {
  // {"en":"Authentication method.","zh_CN":"鉴权方法。"}
  MethodName *string `json:"methodName,omitempty" xml:"methodName,omitempty"`
}

func (s CreateAPropertyVersionRequestOriginsAuthentication) String() string {
  return tea.Prettify(s)
}

func (s CreateAPropertyVersionRequestOriginsAuthentication) GoString() string {
  return s.String()
}

func (s *CreateAPropertyVersionRequestOriginsAuthentication) SetMethodName(v string) *CreateAPropertyVersionRequestOriginsAuthentication {
  s.MethodName = &v
  return s
}

type CreateAPropertyVersionRequestExtraServicePorts struct {
  // {"en":"This is a list of ports other than 80 which are used to handle HTTP requests. The available values can be found in the systemConfigs API's response. If you need another port, please contact our support team.","zh_CN":"指定用于处理HTTP请求的端口列表（80端口除外）。可通过调用'获取系统配置'接口来查询系统支持的端口。如果您需要开通其他端口，请联系技术支持。"}
  Http []*string `json:"http,omitempty" xml:"http,omitempty" type:"Repeated"`
  // {"en":"This is a list of ports other than 443 which are used to handle HTTPS requests. The available values can be found in the systemConfigs API's response. If you need another port, please contact our support team.","zh_CN":"指定用于处理HTTPS请求的端口列表（443端口除外）。可通过调用'获取系统配置'接口来查询系统支持的端口。如果您需要开通其他端口，请联系技术支持。"}
  Https []*string `json:"https,omitempty" xml:"https,omitempty" type:"Repeated"`
}

func (s CreateAPropertyVersionRequestExtraServicePorts) String() string {
  return tea.Prettify(s)
}

func (s CreateAPropertyVersionRequestExtraServicePorts) GoString() string {
  return s.String()
}

func (s *CreateAPropertyVersionRequestExtraServicePorts) SetHttp(v []*string) *CreateAPropertyVersionRequestExtraServicePorts {
  s.Http = v
  return s
}

func (s *CreateAPropertyVersionRequestExtraServicePorts) SetHttps(v []*string) *CreateAPropertyVersionRequestExtraServicePorts {
  s.Https = v
  return s
}

type CreateAPropertyVersionRequestAccessControlRules struct     {
  // {"en":"Range: [ 0 .. 60 ] characters\nAn optional ID for the access control rule.","zh_CN":"取值范围: [ 0 .. 60 ] 字符\n访问控制规则ID。"}
  Id *string `json:"id,omitempty" xml:"id,omitempty"`
  // {"en":"Specify the conditions that the incoming request must match. At least one condition must be specified. If multiple are specified, all must match.","zh_CN":"指定客户端请求必须匹配的条件。必须至少指定一个条件。如果指定了多个条件，则必须全部匹配。"}
  Conditions *CreateAPropertyVersionRequestAccessControlRulesConditions `json:"conditions,omitempty" xml:"conditions,omitempty" require:"true" type:"Struct"`
  // {"en":"Indicates the action to take in response to a request that matches the conditions of the access control rule.","zh_CN":"对于匹配到以上条件的请求所采取的相应操作。"}
  Action *CreateAPropertyVersionRequestAccessControlRulesAction `json:"action,omitempty" xml:"action,omitempty" type:"Struct"`
}

func (s CreateAPropertyVersionRequestAccessControlRules) String() string {
  return tea.Prettify(s)
}

func (s CreateAPropertyVersionRequestAccessControlRules) GoString() string {
  return s.String()
}

func (s *CreateAPropertyVersionRequestAccessControlRules) SetId(v string) *CreateAPropertyVersionRequestAccessControlRules {
  s.Id = &v
  return s
}

func (s *CreateAPropertyVersionRequestAccessControlRules) SetConditions(v *CreateAPropertyVersionRequestAccessControlRulesConditions) *CreateAPropertyVersionRequestAccessControlRules {
  s.Conditions = v
  return s
}

func (s *CreateAPropertyVersionRequestAccessControlRules) SetAction(v *CreateAPropertyVersionRequestAccessControlRulesAction) *CreateAPropertyVersionRequestAccessControlRules {
  s.Action = v
  return s
}

type CreateAPropertyVersionRequestAccessControlRulesConditions struct {
  // {"en":"Enum: https,http\nIndicates whether the incoming request uses HTTP or HTTPS.","zh_CN":"取值范围: https,http\n客户端请求的协议，HTTP或HTTPS。"}
  Scheme *string `json:"scheme,omitempty" xml:"scheme,omitempty"`
  // {"en":"Indicates the hostname requested. It must be one of the hostnames defined in the property. If a wildcard hostname is configured in the property, you can specify either the wildcard hostname or their spcific hostnames here.","zh_CN":"客户端请求对应的域名。必须是加速项目中定义的加速域名之一。如果加速域名按泛域名配置，此处可以指定泛域名或者泛域名下的某个具体域名。"}
  Hostname *string `json:"hostname,omitempty" xml:"hostname,omitempty"`
  // {"en":"Range: <= 500 characters\nThe URI begins with '/' and can end with '*' for prefix matching.","zh_CN":"取值范围: <= 500 字符\n用于前缀匹配的URI，以'/'开头，可以以'*'结尾，表示模糊匹配。"}
  Uri *string `json:"uri,omitempty" xml:"uri,omitempty"`
  // {"en":"An array of ISO-3166-1 alpha-2 codes indicating the countries of the servers, for example, 'US' to represent 'United States of America'.","zh_CN":"服务器所在区域，以ISO-3166-1 两位国家代码表示。例如，'US'代表'美国'。"}
  ServerRegions []*string `json:"serverRegions,omitempty" xml:"serverRegions,omitempty" type:"Repeated"`
  // {"en":"An array of ISO-3166-1 alpha-2 codes indicating one or more countries which are the source of the client requests, for example, 'CN' to represent 'China'.","zh_CN":"客户端请求来源区域，以ISO-3166-1 两位国家代码表示。例如，'CN'代表'中国'。"}
  ClientRegions []*string `json:"clientRegions,omitempty" xml:"clientRegions,omitempty" type:"Repeated"`
  // {"en":"Indicates the starting and ending IP addresses of the client requests to match against. They must be in IPv4 or IPv6 format.","zh_CN":"用于指定要匹配的客户端请求的开始和结束IP地址。必须是IPv4或IPv6格式。"}
  ClientIpRange []*string `json:"clientIpRange,omitempty" xml:"clientIpRange,omitempty" type:"Repeated"`
}

func (s CreateAPropertyVersionRequestAccessControlRulesConditions) String() string {
  return tea.Prettify(s)
}

func (s CreateAPropertyVersionRequestAccessControlRulesConditions) GoString() string {
  return s.String()
}

func (s *CreateAPropertyVersionRequestAccessControlRulesConditions) SetScheme(v string) *CreateAPropertyVersionRequestAccessControlRulesConditions {
  s.Scheme = &v
  return s
}

func (s *CreateAPropertyVersionRequestAccessControlRulesConditions) SetHostname(v string) *CreateAPropertyVersionRequestAccessControlRulesConditions {
  s.Hostname = &v
  return s
}

func (s *CreateAPropertyVersionRequestAccessControlRulesConditions) SetUri(v string) *CreateAPropertyVersionRequestAccessControlRulesConditions {
  s.Uri = &v
  return s
}

func (s *CreateAPropertyVersionRequestAccessControlRulesConditions) SetServerRegions(v []*string) *CreateAPropertyVersionRequestAccessControlRulesConditions {
  s.ServerRegions = v
  return s
}

func (s *CreateAPropertyVersionRequestAccessControlRulesConditions) SetClientRegions(v []*string) *CreateAPropertyVersionRequestAccessControlRulesConditions {
  s.ClientRegions = v
  return s
}

func (s *CreateAPropertyVersionRequestAccessControlRulesConditions) SetClientIpRange(v []*string) *CreateAPropertyVersionRequestAccessControlRulesConditions {
  s.ClientIpRange = v
  return s
}

type CreateAPropertyVersionRequestAccessControlRulesAction struct {
  // {"en":"Indicates the HTTP status code to respond with. It must be in the range 300-309, 400-409, or 500-509 to indicate a redirection or error.","zh_CN":"响应的HTTP状态码，范围必须在300-309、400-409或500-509之间，分别表示重定向或错误。"}
  StatusCode *string `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
  // {"en":"Range: <= 200 characters\nIf the value of the status field is in the range 300-309, the message field's value will be placed in a Location HTTP header returned to the client. If the status is in the range 400-409 or 500-509, the value will be returned in the response body to the client.","zh_CN":"取值范围: <= 200 字符\n如果status字段的值在300-309范围内，message字段的值将在返回给客户端的Location响应头中。如果status字段的值在400-409或500-509范围内，则message字段的值将在返回给客户端的响应体中。"}
  Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s CreateAPropertyVersionRequestAccessControlRulesAction) String() string {
  return tea.Prettify(s)
}

func (s CreateAPropertyVersionRequestAccessControlRulesAction) GoString() string {
  return s.String()
}

func (s *CreateAPropertyVersionRequestAccessControlRulesAction) SetStatusCode(v string) *CreateAPropertyVersionRequestAccessControlRulesAction {
  s.StatusCode = &v
  return s
}

func (s *CreateAPropertyVersionRequestAccessControlRulesAction) SetMessage(v string) *CreateAPropertyVersionRequestAccessControlRulesAction {
  s.Message = &v
  return s
}

type CreateAPropertyVersionRequestRealTimeLog struct {
  // {"en":"The URL that receives the notifications. It must begin with 'http' or 'https'. The server should support the POST method. This is a required field.","zh_CN":"接收通知的服务器URL地址。必须以'http'或'https'开头。服务器须支持POST方法。这是必填字段。"}
  LogUrl *string `json:"logUrl,omitempty" xml:"logUrl,omitempty" require:"true"`
  // {"en":"Default: 1 Range: [ 1 .. 65536 ]\nAn integer between [1, 65536]. n means one notification for every n edge requests. 1 means every edge request will generate a notification. This is optional. Default is 1.","zh_CN":"默认值: 1 取值范围: [ 1 .. 65536 ]\n采样率。介于[1, 65536]之间的整数。n表示每n个边缘请求发送1条通知。1表示每个边缘请求都会发送。这是可选字段。默认值为1。"}
  SampleRate *int `json:"sampleRate,omitempty" xml:"sampleRate,omitempty"`
  // {"en":"Enum: json,none,compactJson\nSpecify json to enable JSON character escaping in variables, compactJson to enable JSON character escaping with whitespace in the body removed, or none to disable escaping.","zh_CN":"取值范围: json,none,compactJson\n指定json以开启变量中的json字符转义；指定compactJson以开启变量中的json字符转义，并移除变量中的空白行。如果要禁用转义则指定none。"}
  Escape *string `json:"escape,omitempty" xml:"escape,omitempty"`
  // {"en":"This is a required field to specify the notification body to be sent to the remote server. It can be any printable text that can be sent in the body of an HTTP POST request.\n\nThe CDN Pro built-in variables can be used within the format field. They will be replaced with the actual values in the notifications.","zh_CN":"该字段用来定义要发送到远程服务器的通知的格式。通知正文可以包括任何能在HTTP POST请求体中发送的可打印文本。这是必填字段。\n\n可以在格式定义中使用CDN Pro支持的内置变量。发送通知时，变量将被替换为实际值。"}
  Format *string `json:"format,omitempty" xml:"format,omitempty" require:"true"`
  // {"en":"HTTP header names and values to be sent to the notification server. A header name can contain any alphanumeric character or hyphen, '-'. A header value can contain any printable characters. It can also include any of the built-in variables supported in the format field of the realTimeLog object.","zh_CN":"需要发送到远程服务器的HTTP请求头名称和值。请求头名称可以包含任何字母，数字或连字符'-'。值可以包含任何可打印字符，也可以使用realTimeLog对象format字段中支持的任何内置变量。"}
  Headers []*CreateAPropertyVersionRequestRealTimeLogHeaders `json:"headers,omitempty" xml:"headers,omitempty" type:"Repeated"`
}

func (s CreateAPropertyVersionRequestRealTimeLog) String() string {
  return tea.Prettify(s)
}

func (s CreateAPropertyVersionRequestRealTimeLog) GoString() string {
  return s.String()
}

func (s *CreateAPropertyVersionRequestRealTimeLog) SetLogUrl(v string) *CreateAPropertyVersionRequestRealTimeLog {
  s.LogUrl = &v
  return s
}

func (s *CreateAPropertyVersionRequestRealTimeLog) SetSampleRate(v int) *CreateAPropertyVersionRequestRealTimeLog {
  s.SampleRate = &v
  return s
}

func (s *CreateAPropertyVersionRequestRealTimeLog) SetEscape(v string) *CreateAPropertyVersionRequestRealTimeLog {
  s.Escape = &v
  return s
}

func (s *CreateAPropertyVersionRequestRealTimeLog) SetFormat(v string) *CreateAPropertyVersionRequestRealTimeLog {
  s.Format = &v
  return s
}

func (s *CreateAPropertyVersionRequestRealTimeLog) SetHeaders(v []*CreateAPropertyVersionRequestRealTimeLogHeaders) *CreateAPropertyVersionRequestRealTimeLog {
  s.Headers = v
  return s
}

type CreateAPropertyVersionRequestRealTimeLogHeaders struct     {
  // {"en":"Name of an HTTP header.","zh_CN":"HTTP标头名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // {"en":"Value of an HTTP header.","zh_CN":"HTTP标头值"}
  Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s CreateAPropertyVersionRequestRealTimeLogHeaders) String() string {
  return tea.Prettify(s)
}

func (s CreateAPropertyVersionRequestRealTimeLogHeaders) GoString() string {
  return s.String()
}

func (s *CreateAPropertyVersionRequestRealTimeLogHeaders) SetName(v string) *CreateAPropertyVersionRequestRealTimeLogHeaders {
  s.Name = &v
  return s
}

func (s *CreateAPropertyVersionRequestRealTimeLogHeaders) SetValue(v string) *CreateAPropertyVersionRequestRealTimeLogHeaders {
  s.Value = &v
  return s
}

type CreateAPropertyVersionRequestVideoSeek struct {
  // {"en":"Range: [ 1 .. 31 ] characters\nName of the query parameter indicating the starting offset in bytes of the content to fetch. The parameter name should begin with a letter (a-z, A-Z) and may be followed by up to 30 letters and numbers.","zh_CN":"取值范围: [ 1 .. 31 ] 字符\n查询参数的名称，用来指定要获取的内容的起始位置（以字节计算）。参数名称应以字母（a-z，A-Z）开头，后面最多可以有30个字母和数字。"}
  StartParameter *string `json:"startParameter,omitempty" xml:"startParameter,omitempty" require:"true"`
  // {"en":"Range: [ 0 .. 31 ] characters\nName of the query parameter indicating the ending offset of the content to fetch. The parameter name should begin with a letter (a-z, A-Z) and may be followed by up to 30 letters and numbers.","zh_CN":"取值范围: [ 0 .. 31 ] 字符\n查询参数的名称，用来指定要获取的内容的结束位置（以字节计算）。参数名称应以字母（a-z，A-Z）开头，后面最多可以有30个字母和数字。"}
  EndParameter *string `json:"endParameter,omitempty" xml:"endParameter,omitempty"`
}

func (s CreateAPropertyVersionRequestVideoSeek) String() string {
  return tea.Prettify(s)
}

func (s CreateAPropertyVersionRequestVideoSeek) GoString() string {
  return s.String()
}

func (s *CreateAPropertyVersionRequestVideoSeek) SetStartParameter(v string) *CreateAPropertyVersionRequestVideoSeek {
  s.StartParameter = &v
  return s
}

func (s *CreateAPropertyVersionRequestVideoSeek) SetEndParameter(v string) *CreateAPropertyVersionRequestVideoSeek {
  s.EndParameter = &v
  return s
}

type CreateAPropertyVersionRequestHeader struct {
}

func (s CreateAPropertyVersionRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s CreateAPropertyVersionRequestHeader) GoString() string {
  return s.String()
}

type CreateAPropertyVersionPaths struct {
  // {"en":"ID of a property","zh_CN":"加速项目ID。"}
  PropertyId *string `json:"propertyId,omitempty" xml:"propertyId,omitempty" require:"true"`
}

func (s CreateAPropertyVersionPaths) String() string {
  return tea.Prettify(s)
}

func (s CreateAPropertyVersionPaths) GoString() string {
  return s.String()
}

func (s *CreateAPropertyVersionPaths) SetPropertyId(v string) *CreateAPropertyVersionPaths {
  s.PropertyId = &v
  return s
}

type CreateAPropertyVersionParameters struct {
}

func (s CreateAPropertyVersionParameters) String() string {
  return tea.Prettify(s)
}

func (s CreateAPropertyVersionParameters) GoString() string {
  return s.String()
}

type CreateAPropertyVersionResponse struct {
}

func (s CreateAPropertyVersionResponse) String() string {
  return tea.Prettify(s)
}

func (s CreateAPropertyVersionResponse) GoString() string {
  return s.String()
}

type CreateAPropertyVersionResponseHeader struct {
  // {"en":"Returns a URL pointing to the new property version created, if the request is accepted. The URL contains the property ID and the number of the new property version. </br> URL format: <code>{scheme}://{domain}/cdn/properties/{propertyId}/versions/{versionNumber}</code> Example URL: <code>https://api.example.com/cdn/properties/5dwa2205f9e9cc0001df7b24/verions/2</code>","zh_CN":"当接口调用成功时，通过Location响应头返回新建的加速项目版本的URL。URL中包含了加速项目的ID以及版本号。</br> URL格式：<code>{协议}://{域名}/cdn/properties/{加速项目ID}/versions/{版本号}</code> URL示例： <code>https://open.chinanetcenter.com/cdn/properties/5dca2205f9e9cc0001df7b33/versions/2</code>"}
  Location *string `json:"Location,omitempty" xml:"Location,omitempty" require:"true"`
}

func (s CreateAPropertyVersionResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s CreateAPropertyVersionResponseHeader) GoString() string {
  return s.String()
}

func (s *CreateAPropertyVersionResponseHeader) SetLocation(v string) *CreateAPropertyVersionResponseHeader {
  s.Location = &v
  return s
}




type UpdateAPropertyVersionPaths struct {
  // {"en" : "ID of a property.", "zh_CN": "加速项目ID。"}
  PropertyId *string `json:"propertyId,omitempty" xml:"propertyId,omitempty" require:"true"`
  // {"en" : "A property version. It must be an integer value >=1. The Get a property version API also permits you to specify the word 'latest' to obtain the most recent version.", "zh_CN": "加速项目版本。必须是大于等于1的整数值。"}
  Version *string `json:"version,omitempty" xml:"version,omitempty" require:"true"`
}

func (s UpdateAPropertyVersionPaths) String() string {
  return tea.Prettify(s)
}

func (s UpdateAPropertyVersionPaths) GoString() string {
  return s.String()
}

func (s *UpdateAPropertyVersionPaths) SetPropertyId(v string) *UpdateAPropertyVersionPaths {
  s.PropertyId = &v
  return s
}

func (s *UpdateAPropertyVersionPaths) SetVersion(v string) *UpdateAPropertyVersionPaths {
  s.Version = &v
  return s
}

type UpdateAPropertyVersionParameters struct {
}

func (s UpdateAPropertyVersionParameters) String() string {
  return tea.Prettify(s)
}

func (s UpdateAPropertyVersionParameters) GoString() string {
  return s.String()
}

type UpdateAPropertyVersionRequestHeader struct {
}

func (s UpdateAPropertyVersionRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s UpdateAPropertyVersionRequestHeader) GoString() string {
  return s.String()
}

type UpdateAPropertyVersionRequest struct {
  // {"en" : "A description of the version.
  // ", "zh_CN": "版本描述。"}
  Description *string `json:"description,omitempty" xml:"description,omitempty"`
  // {"en" : "Range: <= 80 characters ^[a-z.-_]+ 
  // The cachekey hostname must be a string made of lowercase letters, digits, dot, dash, and underscore. If not specified, the incoming Host header value will be used and the cache will not be shared across different hostnames in this property.
  // ", "zh_CN": "取值范围: <= 80 字符 ^[a-z.-_]+ 
  // 用于缓存键（Cache Key）的主机名必须是由小写字母、数字、点、连字符和下划线组成的字符串。如果未指定，默认将使用传入的Host请求头值，并且此加速项目中的不同加速域名之间不共享缓存。"}
  CacheKeyHostname *string `json:"cacheKeyHostname,omitempty" xml:"cacheKeyHostname,omitempty"`
  // {"en" : "Hostnames associated with the property. A valid value is a fully qualified hostname such as www.domain.com or a wildcard hostname such as *.domain.com. Any given hostname can only be in one deployed property at a particular time. However, a wildcard hostname is permitted to overlap other hostnames you own.", "zh_CN": "与加速项目关联的加速域名。必须是FQDN完全限定域名（如 www.domain.com），或泛域名（如*.domain.com）。
  // 同一个加速域名在同一时间只能存在于一个已部署的加速项目中，但泛域名可以与关联的完全限定域名一同部署。"}
  Hostnames []*string `json:"hostnames,omitempty" xml:"hostnames,omitempty" type:"Repeated"`
  // {"en" : "This optional field allows you to configure notifications about client requests to be sent to a remote server. It can be used only if you have access to our realtime_log_switch directive. Please contact our support team if you require this feature.", "zh_CN": "此可选字段用来配置发送消息通知（即实时日志）到您的远程服务器。当有客户端请求访问您的加速域名时，将触发通知。这是高级功能，如果您需要此功能，请联系我们的技术支持开通。"}
  RealTimeLog *UpdateAPropertyVersionRequestRealTimeLog `json:"realTimeLog,omitempty" xml:"realTimeLog,omitempty" type:"Struct"`
  // {"en" : "Range: <= 65530 characters 
  // Refer to Edge Logic Introduction.", "zh_CN": "取值范围: <= 65530 字符 
  // 自定义边缘逻辑。参考边缘逻辑介绍。"}
  EdgeLogic *string `json:"edgeLogic,omitempty" xml:"edgeLogic,omitempty"`
  // {"en" : "Default: False 
  // The value indicates whether all the hostnames in this property have Beian license on file with the Chinese government. This is required to serve this property from servers in mainland China. A value of false means servers outside of mainland China will be used to distribute content to visitors in China. If set to true you must also specify the content type in the beianContentType field.", "zh_CN": "默认值: False 
  // 此加速项目中的所有加速域名是否都已获得中国政府的ICP备案。只有获取备案的域名才能使用中国大陆节点来提供服务。如果设置为true，还必须在beianContentType字段中指定内容类型。"}
  HasBeian *bool `json:"hasBeian,omitempty" xml:"hasBeian,omitempty"`
  // {"en" : "Default: False 
  // This field can be set to either a boolean value or a string. If the value is set to true, the server will redirect all HTTP requests to HTTPS using status code 301. You can also specify string values '302', '307', or '308' instead if you wish to return a different status code when a request is redirected.", "zh_CN": "默认值: False 
  // 此字段可以设置为布尔值或字符串。如果设置为true，则CDN Pro服务器会将所有HTTP请求重定向到HTTPS，并返回301状态码。如果您希望在重定向请求时返回不同的状态码，可在此处指定需要的状态码，如'302'、'307'或'308'。"}
  RedirectHttpToHttps *string `json:"redirectHttpToHttps,omitempty" xml:"redirectHttpToHttps,omitempty"`
  // {"en" : "Describes the origin servers for the property's content.", "zh_CN": "描述加速项目对应的源站。"}
  Origins []*UpdateAPropertyVersionRequestOrigins `json:"origins,omitempty" xml:"origins,omitempty" type:"Repeated"`
  // {"en" : "Default: 1 
  // The value must be set to 1 at this time.", "zh_CN": "默认值: 1 
  // 当前仅允许值为1。"}
  SyntaxVersion *int `json:"syntaxVersion,omitempty" xml:"syntaxVersion,omitempty"`
  // {"en" : "Default: False 
  // Set to true to disable support for HTTP/2 and support only HTTP 1.1.", "zh_CN": "默认值: False 
  // 当值为true时，禁用对HTTP/2的支持，仅支持HTTP/1.1。"}
  DisableHttp2 *bool `json:"disableHttp2,omitempty" xml:"disableHttp2,omitempty"`
  // {"en" : "Default: False 
  // Set to true to enable support for HTTP/3. HTTP/3 requires TLS to work, so this field will be ignored unless you specify a certificate in the tlsCertificateId field.", "zh_CN": "默认值: False 
  // 是否开启HTTP/3，HTTP/3 需要 需要配置TLS 才生效，若您未在 tlsCertificateId 中指定证书，此字段将被忽略。"}
  EnableHttp3 *bool `json:"enableHttp3,omitempty" xml:"enableHttp3,omitempty"`
  // {"en" : "Default: False 
  // Specifies whether the scheme ('http' or 'https') should be included in the cache key. Default behavior is false. Set to true if the content is known to be different for different schemes.
  // ", "zh_CN": "默认值: False 
  // 指定缓存键（Cache Key）是否区分协议（'http'或'https'）。默认为false。如果不同协议响应的内容不同，则设置为true。"}
  SchemeInCacheKey *bool `json:"schemeInCacheKey,omitempty" xml:"schemeInCacheKey,omitempty"`
  // {"en" : "Enum: 1.1,1.2,1.3 
  // Default: 1.3 
  // Maximum supported TLS version.", "zh_CN": "取值范围: 1.1,1.2,1.3 
  // 默认值: 1.3 
  // 支持的TLS最高版本。"}
  TlsMaxVersion *string `json:"tlsMaxVersion,omitempty" xml:"tlsMaxVersion,omitempty"`
  // {"en" : "Range: <= 100 characters 
  // Used to customize hash key for load balancer. The field will soon be deprecated. Please stop using it.", "zh_CN": "取值范围: <= 100 字符 
  // 用于自定义负载均衡器的哈希key。该字段即将废弃，请勿使用该字段。"}
  LoadBalancerHashKey *string `json:"loadBalancerHashKey,omitempty" xml:"loadBalancerHashKey,omitempty"`
  // {"en" : "Refers to a certificate. It is invalid to only set the tlsCertificateId1 field without setting the tlsCertificateId field. If tlsCertificateId is not set, HTTPS will not be enabled for this property. This is a feature to enable you to specify two certificates of different types, i.e. one RSA, one EC. If two certificates of the same type are specified, the one specified by tlsCertificateId will be ignored. ", "zh_CN": "该加速项目用到的证书ID。仅设置tlsCertificateId1字段而不设置tlsCertificateId字段是无效的。如果未设置tlsCertificateId，则不会为此加速项目启用HTTPS。此功能允许您指定两个不同类型的证书，即一个RSA，一个EC。如果指定了两个相同类型的证书，则将忽略tlsCertificateId指定的证书。 "}
  TlsCertificateId *string `json:"tlsCertificateId,omitempty" xml:"tlsCertificateId,omitempty"`
  // {"en" : "Enum: 1.1,1.2,1.3,1 
  // Default: 1 
  // Minimum required TLS version.", "zh_CN": "取值范围: 1.1,1.2,1.3,1 
  // 默认值: 1 
  // 所需的TLS最低版本。"}
  TlsMinVersion *string `json:"tlsMinVersion,omitempty" xml:"tlsMinVersion,omitempty"`
  // {"en" : "Range: <= 2040 characters 
  // Any cipher list supported by 'openssl ciphers'. See https://www.openssl.org/docs/manmaster/man1/ciphers.html
  // ", "zh_CN": "取值范围: <= 2040 字符 
  // 'openssl ciphers'支持的任何加密算法套件。参考：https://www.openssl.org/docs/manmaster/man1/ciphers.html"}
  TlsCiphers *string `json:"tlsCiphers,omitempty" xml:"tlsCiphers,omitempty"`
  // {"en" : "Default: False 
  // This setting applies only if the property has an attached certificate allowing client requests to use HTTPS. If the value of allowProtocolDowngrade is false, we require all the origin servers to support HTTPS. If the value is true, we allow origins that support only HTTP, which reduces security.
  // ", "zh_CN": "默认值: False 
  // 是否允许协议降级。仅当加速项目中存在允许客户端使用HTTPS请求的证书时，此设置才适用。如果allowProtocolDowngrade的值为false，则要求所有源服务器支持HTTPS。如果值为true，则允许仅支持HTTP的源，但这会降低安全性。"}
  AllowProtocolDowngrade *bool `json:"allowProtocolDowngrade,omitempty" xml:"allowProtocolDowngrade,omitempty"`
  // {"en" : "Refers to a certificate. It is invalid to only set the tlsCertificateId1 field without setting the tlsCertificateId field. If tlsCertificateId is not set, HTTPS will not be enabled for this property. This is a feature to enable you to specify two certificates of different types, i.e. one RSA, one EC. If two certificates of the same type are specified, the one specified by tlsCertificateId will be ignored. 
  // ", "zh_CN": "指该加速项目用到的证书ID。仅设置tlsCertificateId1字段而不设置tlsCertificateId字段是无效的。如果未设置tlsCertificateId，则不会为此加速项目启用HTTPS。此功能允许您指定两个不同类型的证书，即一个RSA，一个EC。如果指定了两个相同类型的证书，则将忽略tlsCertificateId指定的证书。"}
  TlsCertificateId1 *string `json:"tlsCertificateId1,omitempty" xml:"tlsCertificateId1,omitempty"`
  // {"en" : "This object allows you to support video players requesting partial content through query string parameters. If you specify videoSeek, you must enter a value for startParameter.", "zh_CN": "此对象用来支持视频播放器通过指定查询参数来请求部分内容。当videoSeek对象存在时，必须为startParameter设置一个值。"}
  VideoSeek *UpdateAPropertyVersionRequestVideoSeek `json:"videoSeek,omitempty" xml:"videoSeek,omitempty" type:"Struct"`
  // {"en" : "Default: 1800 Range: [ 300 .. 86400 ] 
  // Lifespan of TLS session ticket in seconds.", "zh_CN": "默认值: 1800 取值范围: [ 300 .. 86400 ] 
  // TLS会话ticket的有效期（秒）。"}
  TlsSessionTimeout *int `json:"tlsSessionTimeout,omitempty" xml:"tlsSessionTimeout,omitempty"`
  // {"en" : "Default: False 
  // Enables Online Certificate Status Protocol (OCSP) stapling to check the revocation status of digital certificates. (Refer to https://en.wikipedia.org/wiki/OCSP_stapling)", "zh_CN": "默认值: False 
  // 是否启用在线证书状态协议（OCSP）装订以检查数字证书的吊销状态。参考：https://en.wikipedia.org/wiki/OCSP_stapling"}
  EnableOcspStapling *bool `json:"enableOcspStapling,omitempty" xml:"enableOcspStapling,omitempty"`
  // {"en" : "Range: [ 1 .. 24 ] 
  // If you are planning to distribute content through servers in mainland China, you will need to set the hasBeian field to true and also indicate the type of content you are distributing. Enter the value that best corresponds to your content:
  // 
  // <table><tr><th>Value</th><th>Content Type</th></tr><tr><td>1</td><td>Instant Messaging</td></tr><tr><td>2</td><td>Search Engine</td></tr><tr><td>3</td><td>Web Portal</td></tr><tr><td>4</td><td>Online Postal Service</td></tr><tr><td>5</td><td>Online News</td></tr><tr><td>6</td><td>Blog</td></tr><tr><td>7</td><td>Advertisement</td></tr><tr><td>8</td><td>Organization Web Portal</td></tr><tr><td>9</td><td>Shopping</td></tr><tr><td>10</td><td>Online Payment</td></tr><tr><td>11</td><td>Online Bank</td></tr><tr><td>12</td><td>Online Stock Market</td></tr><tr><td>13</td><td>Online Gaming</td></tr><tr><td>14</td><td>Online Music</td></tr><tr><td>15</td><td>Online Movie</td></tr><tr><td>16</td><td>Online Picture</td></tr><tr><td>17</td><td>Software Download</td></tr><tr><td>18</td><td>Job Hunting</td></tr><tr><td>19</td><td>Online Dating</td></tr><tr><td>20</td><td>Online Real Property</td></tr><tr><td>21</td><td>Online Education</td></tr><tr><td>22</td><td>Web Design</td></tr><tr><td>23</td><td>WAP</td></tr><tr><td>24</td><td>Others</td></tr></table>", "zh_CN": "取值范围: [ 1 .. 24 ] 
  // 如果您的域名已获得备案，希望通过中国大陆的服务器分发内容，则需要将hasBeian字段设置为true，并说明您分发的内容类型。请选择与您的内容最匹配的值：
  // 
  // <table><tr><th>值</th><th>内容类型</th></tr><tr><td>1</td><td>即时通信</td></tr><tr><td>2</td><td>搜索引擎</td></tr><tr><td>3</td><td>综合门户</td></tr><tr><td>4</td><td>网上邮局</td></tr><tr><td>5</td><td>网络新闻</td></tr><tr><td>6</td><td>博客/个人空间</td></tr><tr><td>7</td><td>网络广告</td></tr><tr><td>8</td><td>单位门户网站</td></tr><tr><td>9</td><td>网络购物</td></tr><tr><td>10</td><td>网上支付</td></tr><tr><td>11</td><td>网上银行</td></tr><tr><td>12</td><td>网上炒股</td></tr><tr><td>13</td><td>网络游戏</td></tr><tr><td>14</td><td>网络音乐</td></tr><tr><td>15</td><td>网络影视</td></tr><tr><td>16</td><td>网络图片</td></tr><tr><td>17</td><td>软件下载</td></tr><tr><td>18</td><td>网上求职</td></tr><tr><td>19</td><td>在线交友</td></tr><tr><td>20</td><td>网上房产</td></tr><tr><td>21</td><td>网络教育</td></tr><tr><td>22</td><td>网站建设</td></tr><tr><td>23</td><td>WAP</td></tr><tr><td>24</td><td>其他</td></tr></table>"}
  BeianContentType *int `json:"beianContentType,omitempty" xml:"beianContentType,omitempty"`
  // {"en" : "Default: False 
  // Is it allowed to use IPv6 addresses as the source server for this project.", "zh_CN": "默认值: False 
  // 是否允许使用 IPv6 地址作为该项目的源站服务器"}
  EnableIpv6Origin *bool `json:"enableIpv6Origin,omitempty" xml:"enableIpv6Origin,omitempty"`
  // {"en" : "Default value: off Optional values: auto, strict, off
  // When enableIpv6Origin is allowed, this setting will control whether to follow the client IP protocol version back to the source.
  // Auto: Refers to selecting the IP address for returning to the source based on the IP protocol version requested by the client. When the source server is unavailable, it automatically switches to another IP protocol address for returning to the source
  // Strict: Strictly select the IP address for returning to the source based on the IP protocol version requested by the client. When the source server is unavailable, it is not allowed to switch to another protocol address for returning to the source
  // Off: Ignore the IP protocol version requested by the client and randomly select an available IPv4 or IPv6 address to return to the source.", "zh_CN": "默认值：off 可选值：auto、strict、off 当enableIpv6Origin为允许时，该设置将控制是否跟随客户端IP协议版本回源。
  // auto：表示根据客户端请求的IP协议版本来选择回源IP地址，当源站服务器不可用时自动切换成其他IP协议地址回源
  // strict：严格根据客户端请求的IP协议版本来选择回源IP地址，当源站服务器不可用时不允许切换成其他协议地址回源
  // off：忽略客户端请求的IP协议版本，随机选择可用的IPv4或IPv6地址回源"}
  FollowClientIpVersion *string `json:"followClientIpVersion,omitempty" xml:"followClientIpVersion,omitempty"`
  // {"en" : "Specify one or more access control rules to restrict access to your content. More advanced configuration can be done using Edge Logic. These access control rules take precedence over Edge Logic if both are defined.", "zh_CN": "指定一个或多个访问控制规则以限制对内容的访问。可以使用边缘逻辑进行更高级的配置。此处定义的访问控制规则，优先级高于边缘逻辑。"}
  AccessControlRules []*UpdateAPropertyVersionRequestAccessControlRules `json:"accessControlRules,omitempty" xml:"accessControlRules,omitempty" type:"Repeated"`
  // {"en" : "Default: False 
  // By default, CDN Pro takes control of the contents under the /.well-known/{acme-challenge, pki-validation} directories to support certificate auto-renew for properties. If for any reason you need to manage these two directories by yourself on the origin, for example, to implement your own certificate auto-renew mechanism, you can use this configuration option to disable the default behavior by setting its value to true. For more information about our support for certificate auto-renewal, refer to the description of the autoRenew field in the Create a certificate API.", "zh_CN": "默认值: False 
  // 默认情况下，CDN Pro控制/.well-known/{acme-challenge, pki-validation}目录下的内容，以支持加速项目的证书自动更新功能。如果您需要自己在源站管理这两个目录，例如，为了实现您自己的证书自动更新机制，您可以将此字段设置为true来禁用默认行为。关于证书自动更新的更多信息，请参考'创建证书'接口中autoRenew字段的说明。"}
  DisableCertAutomation *bool `json:"disableCertAutomation,omitempty" xml:"disableCertAutomation,omitempty"`
  // {"en" : "Enum: preRewrite,postRewrite 
  // Default: preRewrite 
  // This setting controls how the URI of the incoming request is incorporated into the cache key if you use a 'rewrite' directive in the property's Edge Logic. The default value, 'preRewrite', puts the unmodified URI into the cache key while 'postRewrite' uses the modified URI. If your rewrite directive converts multiple different URIs into the same value, using 'postRewrite' may result in a higher cache hit ratio.", "zh_CN": "取值范围: preRewrite,postRewrite 
  // 默认值: preRewrite 
  // 如果在加速项目的边缘逻辑中使用了'rewrite'指令，可使用该字段来控制客户端请求的URI如何合并到缓存键（Cache Key）中 。默认值'preRewrite'指将改写前的URI放入缓存键，而'postRewrite'则使用改写后的URI。如果您的'rewrite'指令将多个不同的URI改写为相同的值，则使用'postRewrite'可以提高缓存命中率。"}
  CacheKeyUri *string `json:"cacheKeyUri,omitempty" xml:"cacheKeyUri,omitempty"`
  // {"en" : "This field lists ports other than the default 80 used to handle HTTP requests and ports other than the default 443 used to handle HTTPS requests. ", "zh_CN": "除标准的80，443端口外，我们还支持一些扩展端口。可用该字段指定用于处理HTTP和HTTPS请求的扩展端口。"}
  ExtraServicePorts *UpdateAPropertyVersionRequestExtraServicePorts `json:"extraServicePorts,omitempty" xml:"extraServicePorts,omitempty" type:"Struct"`
  // {"en" : "Range: <= 65530 characters 
  // This field allows you to customize load balancing. The field will be deprecated soon. Please use the field edgeLogic.", "zh_CN": "取值范围: <= 65530 字符 
  // 用于控制负载均衡器的行为。该字段即将废弃，请使用edgeLogic字段进行配置。"}
  LoadBalancerLogic *string `json:"loadBalancerLogic,omitempty" xml:"loadBalancerLogic,omitempty"`
  // {"en" : "Default: False 
  // Enable TLS zero round-trip time, a feature of TLS 1.3 that can improve performance for repeat visits to a site. If enabling it though, be sure your site is not vulnerable to the HTTP replay attack.", "zh_CN": "默认值: False 
  // 是否开启TLS 0-RTT功能。这是TLS 1.3版本支持的功能。启用该功能后，当用户频繁访问您的站点时，可提高访问响应速度。需要注意的是，启用该功能可能会加大站点遭受HTTP replay攻击的风险。"}
  Tls0Rtt *bool `json:"tls0Rtt,omitempty" xml:"tls0Rtt,omitempty"`
}

func (s UpdateAPropertyVersionRequest) String() string {
  return tea.Prettify(s)
}

func (s UpdateAPropertyVersionRequest) GoString() string {
  return s.String()
}

func (s *UpdateAPropertyVersionRequest) SetDescription(v string) *UpdateAPropertyVersionRequest {
  s.Description = &v
  return s
}

func (s *UpdateAPropertyVersionRequest) SetCacheKeyHostname(v string) *UpdateAPropertyVersionRequest {
  s.CacheKeyHostname = &v
  return s
}

func (s *UpdateAPropertyVersionRequest) SetHostnames(v []*string) *UpdateAPropertyVersionRequest {
  s.Hostnames = v
  return s
}

func (s *UpdateAPropertyVersionRequest) SetRealTimeLog(v *UpdateAPropertyVersionRequestRealTimeLog) *UpdateAPropertyVersionRequest {
  s.RealTimeLog = v
  return s
}

func (s *UpdateAPropertyVersionRequest) SetEdgeLogic(v string) *UpdateAPropertyVersionRequest {
  s.EdgeLogic = &v
  return s
}

func (s *UpdateAPropertyVersionRequest) SetHasBeian(v bool) *UpdateAPropertyVersionRequest {
  s.HasBeian = &v
  return s
}

func (s *UpdateAPropertyVersionRequest) SetRedirectHttpToHttps(v string) *UpdateAPropertyVersionRequest {
  s.RedirectHttpToHttps = &v
  return s
}

func (s *UpdateAPropertyVersionRequest) SetOrigins(v []*UpdateAPropertyVersionRequestOrigins) *UpdateAPropertyVersionRequest {
  s.Origins = v
  return s
}

func (s *UpdateAPropertyVersionRequest) SetSyntaxVersion(v int) *UpdateAPropertyVersionRequest {
  s.SyntaxVersion = &v
  return s
}

func (s *UpdateAPropertyVersionRequest) SetDisableHttp2(v bool) *UpdateAPropertyVersionRequest {
  s.DisableHttp2 = &v
  return s
}

func (s *UpdateAPropertyVersionRequest) SetEnableHttp3(v bool) *UpdateAPropertyVersionRequest {
  s.EnableHttp3 = &v
  return s
}

func (s *UpdateAPropertyVersionRequest) SetSchemeInCacheKey(v bool) *UpdateAPropertyVersionRequest {
  s.SchemeInCacheKey = &v
  return s
}

func (s *UpdateAPropertyVersionRequest) SetTlsMaxVersion(v string) *UpdateAPropertyVersionRequest {
  s.TlsMaxVersion = &v
  return s
}

func (s *UpdateAPropertyVersionRequest) SetLoadBalancerHashKey(v string) *UpdateAPropertyVersionRequest {
  s.LoadBalancerHashKey = &v
  return s
}

func (s *UpdateAPropertyVersionRequest) SetTlsCertificateId(v string) *UpdateAPropertyVersionRequest {
  s.TlsCertificateId = &v
  return s
}

func (s *UpdateAPropertyVersionRequest) SetTlsMinVersion(v string) *UpdateAPropertyVersionRequest {
  s.TlsMinVersion = &v
  return s
}

func (s *UpdateAPropertyVersionRequest) SetTlsCiphers(v string) *UpdateAPropertyVersionRequest {
  s.TlsCiphers = &v
  return s
}

func (s *UpdateAPropertyVersionRequest) SetAllowProtocolDowngrade(v bool) *UpdateAPropertyVersionRequest {
  s.AllowProtocolDowngrade = &v
  return s
}

func (s *UpdateAPropertyVersionRequest) SetTlsCertificateId1(v string) *UpdateAPropertyVersionRequest {
  s.TlsCertificateId1 = &v
  return s
}

func (s *UpdateAPropertyVersionRequest) SetVideoSeek(v *UpdateAPropertyVersionRequestVideoSeek) *UpdateAPropertyVersionRequest {
  s.VideoSeek = v
  return s
}

func (s *UpdateAPropertyVersionRequest) SetTlsSessionTimeout(v int) *UpdateAPropertyVersionRequest {
  s.TlsSessionTimeout = &v
  return s
}

func (s *UpdateAPropertyVersionRequest) SetEnableOcspStapling(v bool) *UpdateAPropertyVersionRequest {
  s.EnableOcspStapling = &v
  return s
}

func (s *UpdateAPropertyVersionRequest) SetBeianContentType(v int) *UpdateAPropertyVersionRequest {
  s.BeianContentType = &v
  return s
}

func (s *UpdateAPropertyVersionRequest) SetEnableIpv6Origin(v bool) *UpdateAPropertyVersionRequest {
  s.EnableIpv6Origin = &v
  return s
}

func (s *UpdateAPropertyVersionRequest) SetFollowClientIpVersion(v string) *UpdateAPropertyVersionRequest {
  s.FollowClientIpVersion = &v
  return s
}

func (s *UpdateAPropertyVersionRequest) SetAccessControlRules(v []*UpdateAPropertyVersionRequestAccessControlRules) *UpdateAPropertyVersionRequest {
  s.AccessControlRules = v
  return s
}

func (s *UpdateAPropertyVersionRequest) SetDisableCertAutomation(v bool) *UpdateAPropertyVersionRequest {
  s.DisableCertAutomation = &v
  return s
}

func (s *UpdateAPropertyVersionRequest) SetCacheKeyUri(v string) *UpdateAPropertyVersionRequest {
  s.CacheKeyUri = &v
  return s
}

func (s *UpdateAPropertyVersionRequest) SetExtraServicePorts(v *UpdateAPropertyVersionRequestExtraServicePorts) *UpdateAPropertyVersionRequest {
  s.ExtraServicePorts = v
  return s
}

func (s *UpdateAPropertyVersionRequest) SetLoadBalancerLogic(v string) *UpdateAPropertyVersionRequest {
  s.LoadBalancerLogic = &v
  return s
}

func (s *UpdateAPropertyVersionRequest) SetTls0Rtt(v bool) *UpdateAPropertyVersionRequest {
  s.Tls0Rtt = &v
  return s
}

type UpdateAPropertyVersionRequestRealTimeLog struct {
  // {"en" : "The URL that receives the notifications. It must begin with 'http' or 'https'. The server should support the POST method. This is a required field.", "zh_CN": "接收通知的服务器URL地址。必须以'http'或'https'开头。服务器须支持POST方法。这是必填字段。"}
  LogUrl *string `json:"logUrl,omitempty" xml:"logUrl,omitempty" require:"true"`
  // {"en" : "Default: 1 Range: [ 1 .. 65536 ] 
  // An integer between [1, 65536]. n means one notification for every n edge requests. 1 means every edge request will generate a notification. This is optional. Default is 1.", "zh_CN": "默认值: 1 取值范围: [ 1 .. 65536 ] 
  // 采样率。介于[1, 65536]之间的整数。n表示每n个边缘请求发送1条通知。1表示每个边缘请求都会发送。这是可选字段。默认值为1。"}
  SampleRate *int `json:"sampleRate,omitempty" xml:"sampleRate,omitempty"`
  // {"en" : "Enum: json,none 
  // Specify json to enable JSON character escaping in variables or none to disable escaping.", "zh_CN": "取值范围: json,none 
  // 指定json以开启变量中的json字符转义。如果要禁用转义则指定none。"}
  Escape *string `json:"escape,omitempty" xml:"escape,omitempty"`
  // {"en" : "This is a required field to specify the notification body to be sent to the remote server. It can be any printable text that can be sent in the body of an HTTP POST request. 
  // 
  // The following built-in variables can be used within the format field. They will be replaced with the actual values in the notifications.
  // <table><tr><th>Name</th><th>Description</th></tr><tr><td>$body_bytes_sent</td><td>Size of the response body.</td></tr><tr><td>$bytes_sent</td><td>Size of the response, including body, headers and response line.</td></tr><tr><td>$client_country_code</td><td>An ISO 3166-1 country code representing the country of the client request, for example, 'US'. 'ZZ' is returned if the country is unknown.</td></tr><tr><td>$client_real_ip</td><td>IP address of the client request.</td></tr><tr><td>$cookie_x</td><td>This lets you obtain any cookie named x.  For example, $cookie_account would retrieve the value of a cookie named 'account'.</td></tr><tr><td>$http_x</td><td>Obtain any HTTP header named x from the original request. The header name is converted to lower case with dashes replaced by underscores. For example, specify $http_user_agent to get the value of User-Agent.</td></tr><tr><td>$msec</td><td>Current unix time in seconds with millisecond precision.</td></tr><tr><td>$qtl_req_id</td><td>Unique identifier representing the request.</td></tr><tr><td>$request_uri</td><td>HTTP request URI</td></tr><tr><td>$request_method</td><td>The HTTP request method used to access the origin.</td></tr><tr><td>$request_time</td><td>Response time in milliseconds. It is the time between receiving the request's  first byte and serving the last byte of the response.</td></tr><tr><td>$sc_completed</td><td>1 to indicate the last byte of the object was served to the user, 0 otherwise.</td></tr><tr><td>$sc_initial</td><td>1 to indicate the first byte of the object was served to the user, 0 otherwise.</td></tr><tr><td>$scheme</td><td>Indicates the protocol of the user's request ('http' or https').</td></tr><tr><td>$sent_http_content_length</td><td>the original file size.</td></tr><tr><td>$sent_http_x</td><td>Obtain the value of an HTTP header named x that is returned in the response to the client. The header name should be converted to lower case with dashes replaced by underscores. For example, $sent_http_etag would give you the value of the ETag header.</td></tr><tr><td>$server_addr</td><td>IP address of the edge node serving the user's request.</td></tr><tr><td>$server_protocol</td><td>Indicates the version of HTTP used in the user's request, either 'HTTP/1.0', 'HTTP/1.1', or 'HTTP/2.0'.</td></tr><tr><td>$ssl_cipher</td><td>Indicates the cipher suite used for the TLS (SSL) connection.</td></tr><tr><td>$ssl_server_name</td><td>The hostname that a client initiating a TLS (SSL) connection is attempting to connect to. It is only sent by clients supporting SNI (Server Name Indication).</td></tr><tr><td>$ssl_protocol</td><td>Indicates the TLS version used for the TLS (SSL) connection. Example versions include 'SSLv3', 'TLSv1', 'TLSv1.1', 'TLSv1.2', and 'unknown'.</td></tr><tr><td>$status</td><td>An HTTP response code for the user's request.</td></tr><tr><td>$tcpinfo_rtt</td><td>The time in microseconds taken by a packet to travel to the destination and back.</td></tr></table>", "zh_CN": "该字段用来定义要发送到远程服务器的通知的格式。通知正文可以包括任何能在HTTP POST请求体中发送的可打印文本。这是必填字段。
  // 
  // 可以使用以下内置变量定义通知的格式。发送通知时，它们将被替换为实际值。
  // <table><tr><th>变量名称</th><th>描述</th></tr><tr><td>$body_bytes_sent</td><td>响应体大小。</td></tr><tr><td>$bytes_sent</td><td>响应的大小，包括响应体、响应头和响应行。</td></tr><tr><td>$client_country_code</td><td>客户端请求来源国家，以ISO 3166-1国家代码表示。例如'US'。如果国家/地区未知，则返回'ZZ'。</td></tr><tr><td>$client_real_ip</td><td>客户端请求的IP地址。</td></tr><tr><td>$cookie_x</td><td>获取某个cookie。例如，指定$cookie_account可获取名为'account'的cookie值。</td></tr><tr><td>$http_x</td><td>从原始请求中获取某个HTTP请求头。请求头名称需转换为小写，并用下划线替换连字符。例如，指定$http_user_agent来获取User-Agent的值。</td></tr><tr><td>$msec</td><td>当前unix时间，以毫秒为单位。</td></tr><tr><td>$qtl_req_id</td><td>请求的唯一标识符。</td></tr><tr><td>$request_uri</td><td>HTTP请求URI。</td></tr><tr><td>$request_method</td><td>用于访问源站的HTTP请求方法。</td></tr><tr><td>$request_time</td><td>响应时间，以毫秒为单位。这是从接收到请求的第一个字节到服务端响应最后一个字节之间的时间。</td></tr><tr><td>$sc_completed</td><td>1表示对象的最后一个字节已返回给用户，否则为0。</td></tr><tr><td>$sc_initial</td><td>1表示对象的第一个字节已返回给用户，否则为0。</td></tr><tr><td>$scheme</td><td>表示用户请求的协议（'http'或'https'）。</td></tr><tr><td>$sent_http_content_length</td><td>原始文件大小。</td></tr><tr><td>$sent_http_x</td><td>获取在对客户端响应中某个HTTP响应头的值。响应头名称需转换为小写，并用下划线替换连字符。例如，$sent_http_etag可获取ETag头的值。</td></tr><tr><td>$server_addr</td><td>为用户请求提供服务的边缘节点的IP地址。</td></tr><tr><td>$server_protocol</td><td>表示用户请求中使用的HTTP版本，可以是'HTTP/1.0'、'HTTP/1.1'或'HTTP/2.0'。</td></tr><tr><td>$ssl_cipher</td><td>表示用于TLS（SSL）连接的加密算法套件。</td></tr><tr><td>$ssl_server_name</td><td>客户端发起TLS（SSL）连接所要连接的域名。仅由支持SNI（Server Name Indication）的客户端发送。</td></tr><tr><td>$ssl_protocol</td><td>表示用于TLS（SSL）连接的TLS版本。例如，'SSLv3'、'TLSv1'、'TLSv1.1'、'TLSv1.2'和'unknown'。</td></tr><tr><td>$status</td><td>用户请求的HTTP状态码。</td></tr><tr><td>$tcpinfo_rtt</td><td>数据包往返目的地所用的时间，以微秒为单位。</td></tr></table>"}
  Format *string `json:"format,omitempty" xml:"format,omitempty"`
  // {"en" : "HTTP header names and values to be sent to the notification server. A header name can contain any alphanumeric character or hyphen, '-'. A header value can contain any printable characters. It can also include any of the built-in variables supported in the format field of the realTimeLog object.", "zh_CN": "需要发送到远程服务器的HTTP请求头名称和值。请求头名称可以包含任何字母，数字或连字符'-'。值可以包含任何可打印字符，也可以使用realTimeLog对象format字段中支持的任何内置变量。"}
  Headers []*UpdateAPropertyVersionRequestRealTimeLogHeaders `json:"headers,omitempty" xml:"headers,omitempty" type:"Repeated"`
}

func (s UpdateAPropertyVersionRequestRealTimeLog) String() string {
  return tea.Prettify(s)
}

func (s UpdateAPropertyVersionRequestRealTimeLog) GoString() string {
  return s.String()
}

func (s *UpdateAPropertyVersionRequestRealTimeLog) SetLogUrl(v string) *UpdateAPropertyVersionRequestRealTimeLog {
  s.LogUrl = &v
  return s
}

func (s *UpdateAPropertyVersionRequestRealTimeLog) SetSampleRate(v int) *UpdateAPropertyVersionRequestRealTimeLog {
  s.SampleRate = &v
  return s
}

func (s *UpdateAPropertyVersionRequestRealTimeLog) SetEscape(v string) *UpdateAPropertyVersionRequestRealTimeLog {
  s.Escape = &v
  return s
}

func (s *UpdateAPropertyVersionRequestRealTimeLog) SetFormat(v string) *UpdateAPropertyVersionRequestRealTimeLog {
  s.Format = &v
  return s
}

func (s *UpdateAPropertyVersionRequestRealTimeLog) SetHeaders(v []*UpdateAPropertyVersionRequestRealTimeLogHeaders) *UpdateAPropertyVersionRequestRealTimeLog {
  s.Headers = v
  return s
}

type UpdateAPropertyVersionRequestRealTimeLogHeaders struct     {
  // {"en" : "Name of an HTTP header.", "zh_CN": "HTTP标头名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // {"en" : "Value of an HTTP header.", "zh_CN": "HTTP标头值"}
  Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s UpdateAPropertyVersionRequestRealTimeLogHeaders) String() string {
  return tea.Prettify(s)
}

func (s UpdateAPropertyVersionRequestRealTimeLogHeaders) GoString() string {
  return s.String()
}

func (s *UpdateAPropertyVersionRequestRealTimeLogHeaders) SetName(v string) *UpdateAPropertyVersionRequestRealTimeLogHeaders {
  s.Name = &v
  return s
}

func (s *UpdateAPropertyVersionRequestRealTimeLogHeaders) SetValue(v string) *UpdateAPropertyVersionRequestRealTimeLogHeaders {
  s.Value = &v
  return s
}

type UpdateAPropertyVersionRequestOrigins struct     {
  // {"en" : "^[a-zA-z0-9_] 
  // Name of an origin. It must be unique within this property.
  // ", "zh_CN": "^[a-zA-z0-9_] 
  // 源站名称，在此加速项目中必须唯一。"}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // {"en" : "Each entry should be a string consisting of an address optionally followed by one or more parameters, separated by space characters. The address can take one of the following three forms:
  // {hostname or IP address}
  // {hostname or IP address}:{http port}
  // {hostname or IP address}:{http port},{https port}
  // Values for the HTTP and HTTPS ports should be integers in the range [1,65535]. 
  // Even if the value of supportedProtocol is 'https', the HTTPS port has to be specified after the comma.
  // In the third form, one of the two ports can be empty.
  // 
  // Supported parameters are 'backup' and 'weight'.
  // 
  // 'backup' is used to indicate the server is a backup server. It will be used when the primary servers are unavailable.
  // 
  // 'weight' is a value in the range [1, 100]. The default value is 1. It lets you control the likelihood that one origin server will be used relative to another.
  // 
  // There should be at least one primary server defined; it does not have the 'backup' parameter.
  // 
  // Example:
  // ['www.abc.com weight=1', 'www1.abc.com:8080 weight=4', 'www2.abc.com:880,4443 backup']
  // 
  // This example specifies two primary servers which are used in a 1:4 ratio meaning one gets about 20% of the requests to origin while the other gets 80% of the requests. In addition, a backup server is specified.
  // ", "zh_CN": "源站服务器列表。每个条目为一个字符串，由一个地址与一个或多个参数组成，参数之间以空格分隔。地址可以采用以下三种形式之一：
  // {域名或IP地址}
  // {域名或IP地址}:{HTTP端口}
  // {域名或IP地址}:{HTTP端口},{HTTPS端口}
  // HTTP和HTTPS端口的值应为[1,65535]范围内的整数。
  // 即使supportedProtocol的值设为'https'，也必须在此处指定HTTPS端口。
  // 在第三种形式中，两个端口中的一个可以为空。
  // 支持ipv6 当enableIpv6Origin为true时，ipv6传入{域名或IP地址}:{HTTP端口},{HTTPS端口}时，域名或IP地址需要放入[]，如[2001:0db8:85a3:0000:0000:8a2e:0370:7334]:80
  // 
  // 支持的参数包括'backup'和'weight'。
  // 'backup'用于标识备份服务器。它将在主服务器不可用时使用。
  // 'weight'是范围[1,100]内的值，默认值为1，用来设置服务器权重，控制一台源站服务器相对于另一台被使用的可能性。
  // 每个源站应至少定义一个主服务器（即不带'backup'参数的服务器）。
  // 
  // 示例：
  // ['www.abc.com weight=1', 'www1.abc.com:8080 weight=4', 'www2.abc.com:880,4443 backup']
  // 
  // 此示例指定了两个主服务器和一个备份服务器，其中，主服务器的权重为1:4，表示第一个服务器将获得约20%的回源请求，而另一个将获得约80%。"}
  Servers []*string `json:"servers,omitempty" xml:"servers,omitempty" type:"Repeated"`
  // {"en" : "Enum: http,https,both 
  // This optional field indicates the protocol supported by the origin server. If this property has a certificate attached to it, the value can be set to http if the allowProtocolDowngrade field is also set to true.", "zh_CN": "取值范围: http,https,both 
  // 此可选字段表示源站服务器支持的协议。如果此加速项目附加了证书，且allowProtocolDowngrade字段也设置为true，则可以将该值设置为http。"}
  SupportedProtocol *string `json:"supportedProtocol,omitempty" xml:"supportedProtocol,omitempty"`
  // {"en" : "Default: auto 
  // Optional. This parameter tells us how important it is to directly go to this origin without going through any intermediate cache to fetch content. The values can be 'noDirect', 'auto', 'alwaysDirect'. 
  // 
  // 'noDirect' means we always go through at least one intermediate cache. Customers who care particularly about the origin's workload can use this option. Our cache scheduler will dynamically pick the intermediate cache based on performance and resource availability.
  // 
  // 'auto' means the cache scheduler will make the choice dynamically based on performance and resource availability. This is the default behavior if not specified.
  // 
  // 'alwaysDirect' means we always directly connect to this origin.
  // ", "zh_CN": "默认值: auto 
  // 此可选字段用来指定回源方式，可以是'noDirect'、'auto'、'alwaysDirect'。
  // 
  // 'noDirect'指总是通过至少一个中间缓存节点回源。对于特别关心源站负载的客户可以使用此选项。我们的调度程序将根据性能和资源可用性动态选择中间缓存节点。
  // 'alwaysDirect'指总是直接回源。
  // 'auto'指自动选择直接回源或通过中间缓存节点回源。调度程序将根据性能和资源可用性动态做出选择。该字段未指定时，采用此默认行为。"}
  DirectConnection *string `json:"directConnection,omitempty" xml:"directConnection,omitempty"`
  // {"en" : "Optional. It should be a valid hostname. It will also be used as the SNI servername when contacting the HTTPS origin. We also allow it to be an nginx variable that begins with '$' followed by [a-z_]{3,60}. Nginx variable names are case insensitive, so we only allow lowercase characters.
  // If not specified, the value of the HOST header in the request will be used. 
  // 
  // One constraint is that if 'aswS3' is specified for origin authentication, the value of hostHeader cannot be a variable or empty. It has to be a valid Amazon S3 hostname. Refer to https://docs.aws.amazon.com/AmazonS3/latest/dev/VirtualHosting.html", "zh_CN": "可选，用来指定回源HOST头部。必须是有效的域名。当连接HTTPS源站时，该值也作为SNI域名。可以用nginx变量来指定，变量以'$'开头，后跟[a-z_]{3,60}。Nginx变量名不区分大小写，因此我们只允许小写字符。
  // 如果未指定，将使用请求中的HOST请求头的值。
  // 
  // 注意：如果指定了'awsS3'作为回源鉴权参数，则hostHeader的值不能为变量或为空，而必须是有效的Amazon S3主机名。参考：https://docs.aws.amazon.com/AmazonS3/latest/dev/VirtualHosting.html"}
  HostHeader *string `json:"hostHeader,omitempty" xml:"hostHeader,omitempty"`
  // {"en" : "Default: False 
  // Optional. It controls whether the cache will validate the certificate of the origin.", "zh_CN": "默认值: False 
  // 可选。用来设定CDN缓存节点是否验证源站证书。"}
  VerifyOrigin *bool `json:"verifyOrigin,omitempty" xml:"verifyOrigin,omitempty"`
  // {"en" : "Optional. It is a structure to specify the parameters, such as password, for authentication with the origin servers. It should have at least one field: 'methodName', which should be a string indicating one of the predefined authentication methods. The only valid value at this time is 'awsS3'. When 'awsS3' is used, the following parameters are required:
  // awsS3.region:
  // The AWS S3 region where your resources are hosted, e.g. 'us-west-1', 'ap-northeast-1', 'eu-north-1', and 'cn-north-1'.
  // awsS3.accessKey:
  // Your AWS access key that CDN Pro will use to access your resources stored on AWS S3.
  // awsS3.secretKey:
  // Your AWS secret key that CDN Pro will use to access your resources stored on AWS S3.
  // Also, the value of the hostHeader field cannot be a variable or empty. It must be a valid Amazon S3 hostname.
  // 
  // Example:
  // 
  // {'methodName':'awsS3',
  // 'awsS3':{
  // 'region':'us-west-1',
  // 'accessKey':'sdnu2932',
  // 'secretKey':'d12345678abcdefghi'
  // }}
  // ", "zh_CN": "可选。用于指定与源服务器进行身份验证（回源鉴权）的相关参数（如密码）。必须至少有'methodName'（字符串）字段，用来指定预定义的鉴权方法。目前仅支持源站为AWS S3的回源鉴权，因此唯一有效的值是'awsS3'。使用'awsS3'时，需要提供以下参数：
  // awsS3.region:
  // 您存储在S3上的资源所在的AWS区域。例如，'us-west-1'，'ap-northeast-1'，'eu-north-1'，'cn-north-1'。
  // awsS3.accessKey:
  // 您的 AWS 访问密钥（access key），CDN Pro 将使用它来访问您存储在 AWS S3 上的资源。
  // awsS3.secretKey:
  // 您的 AWS 密钥（secret key），CDN Pro 将使用它来访问您存储在 AWS S3 上的资源。
  // 此外，hostHeader字段的值不能为变量或为空，必须是有效的Amazon S3主机名。
  // 
  // 示例：
  // 
  // {'methodName':'awsS3',
  // 'awsS3':{
  // 'region':'us-west-1',
  // 'accessKey':'sdnu2932',
  // 'secretKey':'d12345678abcdefghi'
  // }}
  // "}
  Authentication *UpdateAPropertyVersionRequestOriginsAuthentication `json:"authentication,omitempty" xml:"authentication,omitempty" type:"Struct"`
  // {"en" : "Default: 60 Range: [ 5 .. 600 ] 
  // Timeout in seconds during which an idle keepalive connection to an upstream server will stay open. A service quota setting of maxUpstreamKeepaliveTimeOut can change the maximum permitted value.", "zh_CN": "默认值: 60 取值范围: [ 5 .. 600 ] 
  // 该字段用于指定CDN Pro服务器和源站建连的Keep-Alive超时时间，单位为秒。通过maxUpstreamKeepaliveTimeOut 该服务设置项可以更改允许的最大值。如果需要调整最大值，请联系我们的技术支持。"}
  KeepAliveTimeout *int `json:"keepAliveTimeout,omitempty" xml:"keepAliveTimeout,omitempty"`
  // {"en" : "
  // This setting allows you to specify the number of unsuccessful attempts to reach one of the origin's IP addresses or peers before it is marked as unavailable for a period of time designated by the timeout. If all peers of all servers are unavailable, requests for content from the origin will receive an immediate 502 HTTP response. By default, six attempts to a peer are made, after which a one-second timeout applies to that peer. To disable this feature, set peerFailureTimeout to 'off'.", "zh_CN": "
  // 该字段用于设置源站故障剔除。开启此功能后，当连接某个源站服务器的失败次数达到设定阈值，该源站服务器将被标记为不可用，并保持该状态直到设定的超时时长。失败次数阈值和超时时长分别通过failureThreshold参数和timeout参数设置。CDN Pro回源时将剔除不可用的源站服务器。如果所有源站服务器都被标记为不可用，则对应的请求将立即响应502状态码。默认情况下，当连接某个源站服务器连续失败6次之后，该服务器将被标记为不可用，超时时长为1秒。如果要禁用此功能，请将peerFailureTimeout设置为'off'。开启源站故障剔除配置示例：<code>{'peerFailureTimeout':{'failureThreshold':10,'timeout':3}}</code>"}
  PeerFailureTimeout *string `json:"peerFailureTimeout,omitempty" xml:"peerFailureTimeout,omitempty"`
  // {"en" : "Refers to a certificate used to authenticate with the origin server. This certificate must also be deployed.", "zh_CN": "用于验证源站服务器的证书，该证书同样必须被部署。"}
  TlsCertificateId *string `json:"tlsCertificateId,omitempty" xml:"tlsCertificateId,omitempty"`
  // {"en" : "Refers to the ID of an origin shield representing a set of servers that will make requests to your origin servers rather than the CDN Pro edge servers. This can help further reduce traffic to your origin. Origin shield is allowed only when directConnection is set to noDirect. This is an advanced feature that can be enabled by contacting our support team. The origin shields API provides a list of available shields along with their IP addresses. It is best to select a shield whose region is closest to your origin servers. Use of a shield in China requires your property have 'hasBeian' set to true. If your origin servers have an access control list, add the IPs of the shield you choose to use.", "zh_CN": "指定某个源站盾（origin shield）的ID。使用源站盾功能后，所有回源请求都会通过源站盾中转回源，这可以帮助收敛回源流量。需要把directConnection设置为noDirect，才允许开启源站盾。源站盾是高级功能，如需使用请联系我们的技术支持开通。可通过调用‘获取源站盾列表’接口查询可用的源站盾及其对应的IP地址。您可根据源站的位置，选择合适的源站盾。如果您需使用中国大陆的源站盾，则该加速项目的所有域名必须已取得备案。如果您的源站开启了访问控制，请将所选择源站盾的IP地址加入白名单。"}
  Shield *string `json:"shield,omitempty" xml:"shield,omitempty"`
}

func (s UpdateAPropertyVersionRequestOrigins) String() string {
  return tea.Prettify(s)
}

func (s UpdateAPropertyVersionRequestOrigins) GoString() string {
  return s.String()
}

func (s *UpdateAPropertyVersionRequestOrigins) SetName(v string) *UpdateAPropertyVersionRequestOrigins {
  s.Name = &v
  return s
}

func (s *UpdateAPropertyVersionRequestOrigins) SetServers(v []*string) *UpdateAPropertyVersionRequestOrigins {
  s.Servers = v
  return s
}

func (s *UpdateAPropertyVersionRequestOrigins) SetSupportedProtocol(v string) *UpdateAPropertyVersionRequestOrigins {
  s.SupportedProtocol = &v
  return s
}

func (s *UpdateAPropertyVersionRequestOrigins) SetDirectConnection(v string) *UpdateAPropertyVersionRequestOrigins {
  s.DirectConnection = &v
  return s
}

func (s *UpdateAPropertyVersionRequestOrigins) SetHostHeader(v string) *UpdateAPropertyVersionRequestOrigins {
  s.HostHeader = &v
  return s
}

func (s *UpdateAPropertyVersionRequestOrigins) SetVerifyOrigin(v bool) *UpdateAPropertyVersionRequestOrigins {
  s.VerifyOrigin = &v
  return s
}

func (s *UpdateAPropertyVersionRequestOrigins) SetAuthentication(v *UpdateAPropertyVersionRequestOriginsAuthentication) *UpdateAPropertyVersionRequestOrigins {
  s.Authentication = v
  return s
}

func (s *UpdateAPropertyVersionRequestOrigins) SetKeepAliveTimeout(v int) *UpdateAPropertyVersionRequestOrigins {
  s.KeepAliveTimeout = &v
  return s
}

func (s *UpdateAPropertyVersionRequestOrigins) SetPeerFailureTimeout(v string) *UpdateAPropertyVersionRequestOrigins {
  s.PeerFailureTimeout = &v
  return s
}

func (s *UpdateAPropertyVersionRequestOrigins) SetTlsCertificateId(v string) *UpdateAPropertyVersionRequestOrigins {
  s.TlsCertificateId = &v
  return s
}

func (s *UpdateAPropertyVersionRequestOrigins) SetShield(v string) *UpdateAPropertyVersionRequestOrigins {
  s.Shield = &v
  return s
}

type UpdateAPropertyVersionRequestOriginsAuthentication struct {
  // {"en" : "Authentication method.", "zh_CN": "鉴权方法。"}
  MethodName *string `json:"methodName,omitempty" xml:"methodName,omitempty"`
}

func (s UpdateAPropertyVersionRequestOriginsAuthentication) String() string {
  return tea.Prettify(s)
}

func (s UpdateAPropertyVersionRequestOriginsAuthentication) GoString() string {
  return s.String()
}

func (s *UpdateAPropertyVersionRequestOriginsAuthentication) SetMethodName(v string) *UpdateAPropertyVersionRequestOriginsAuthentication {
  s.MethodName = &v
  return s
}

type UpdateAPropertyVersionRequestVideoSeek struct {
  // {"en" : "Range: [ 1 .. 31 ] characters 
  // Name of the query parameter indicating the starting offset in bytes of the content to fetch. The parameter name should begin with a letter (a-z, A-Z) and may be followed by up to 30 letters and numbers.", "zh_CN": "取值范围: [ 1 .. 31 ] 字符 
  // 查询参数的名称，用来指定要获取的内容的起始位置（以字节计算）。参数名称应以字母（a-z，A-Z）开头，后面最多可以有30个字母和数字。"}
  StartParameter *string `json:"startParameter,omitempty" xml:"startParameter,omitempty" require:"true"`
  // {"en" : "Range: [ 0 .. 31 ] characters 
  // Name of the query parameter indicating the ending offset of the content to fetch. The parameter name should begin with a letter (a-z, A-Z) and may be followed by up to 30 letters and numbers.", "zh_CN": "取值范围: [ 0 .. 31 ] 字符 
  // 查询参数的名称，用来指定要获取的内容的结束位置（以字节计算）。参数名称应以字母（a-z，A-Z）开头，后面最多可以有30个字母和数字。"}
  EndParameter *string `json:"endParameter,omitempty" xml:"endParameter,omitempty"`
}

func (s UpdateAPropertyVersionRequestVideoSeek) String() string {
  return tea.Prettify(s)
}

func (s UpdateAPropertyVersionRequestVideoSeek) GoString() string {
  return s.String()
}

func (s *UpdateAPropertyVersionRequestVideoSeek) SetStartParameter(v string) *UpdateAPropertyVersionRequestVideoSeek {
  s.StartParameter = &v
  return s
}

func (s *UpdateAPropertyVersionRequestVideoSeek) SetEndParameter(v string) *UpdateAPropertyVersionRequestVideoSeek {
  s.EndParameter = &v
  return s
}

type UpdateAPropertyVersionRequestAccessControlRules struct     {
  // {"en" : "Range: [ 0 .. 60 ] characters 
  // An optional ID for the access control rule.", "zh_CN": "取值范围: [ 0 .. 60 ] 字符 
  // 访问控制规则ID。"}
  Id *string `json:"id,omitempty" xml:"id,omitempty"`
  // {"en" : "Specify the conditions that the incoming request must match. At least one condition must be specified. If multiple are specified, all must match.", "zh_CN": "指定客户端请求必须匹配的条件。必须至少指定一个条件。如果指定了多个条件，则必须全部匹配。"}
  Conditions *UpdateAPropertyVersionRequestAccessControlRulesConditions `json:"conditions,omitempty" xml:"conditions,omitempty" type:"Struct"`
  // {"en" : "Indicates the action to take in response to a request that matches the conditions of the access control rule.", "zh_CN": "对于匹配到以上条件的请求所采取的相应操作。"}
  Action *UpdateAPropertyVersionRequestAccessControlRulesAction `json:"action,omitempty" xml:"action,omitempty" type:"Struct"`
}

func (s UpdateAPropertyVersionRequestAccessControlRules) String() string {
  return tea.Prettify(s)
}

func (s UpdateAPropertyVersionRequestAccessControlRules) GoString() string {
  return s.String()
}

func (s *UpdateAPropertyVersionRequestAccessControlRules) SetId(v string) *UpdateAPropertyVersionRequestAccessControlRules {
  s.Id = &v
  return s
}

func (s *UpdateAPropertyVersionRequestAccessControlRules) SetConditions(v *UpdateAPropertyVersionRequestAccessControlRulesConditions) *UpdateAPropertyVersionRequestAccessControlRules {
  s.Conditions = v
  return s
}

func (s *UpdateAPropertyVersionRequestAccessControlRules) SetAction(v *UpdateAPropertyVersionRequestAccessControlRulesAction) *UpdateAPropertyVersionRequestAccessControlRules {
  s.Action = v
  return s
}

type UpdateAPropertyVersionRequestAccessControlRulesConditions struct {
  // {"en" : "Enum: https,http 
  // Indicates whether the incoming request uses HTTP or HTTPS.", "zh_CN": "取值范围: https,http 
  // 客户端请求的协议，HTTP或HTTPS。"}
  Scheme *string `json:"scheme,omitempty" xml:"scheme,omitempty"`
  // {"en" : "Indicates the hostname requested. It must be one of the hostnames defined in the property.", "zh_CN": "请求对应的加速域名。必须是加速项目中定义的加速域名之一。"}
  Hostname *string `json:"hostname,omitempty" xml:"hostname,omitempty"`
  // {"en" : "Range: <= 500 characters 
  // The URI begins with '/' and can end with '*' for prefix matching. ", "zh_CN": "取值范围: <= 500 字符 
  // 用于前缀匹配的URI，以'/'开头，可以以'*'结尾。"}
  Uri *string `json:"uri,omitempty" xml:"uri,omitempty"`
  // {"en" : "An array of ISO-3166-1 alpha-2 codes indicating the countries of the servers, for example, 'US' to represent 'United States of America'.", "zh_CN": "服务器所在区域，以ISO-3166-1 两位国家代码表示。例如，'US'代表'美国'。"}
  ServerRegions []*string `json:"serverRegions,omitempty" xml:"serverRegions,omitempty" type:"Repeated"`
  // {"en" : "An array of ISO-3166-1 alpha-2 codes indicating one or more countries which are the source of the client requests, for example, 'CN' to represent 'China'.", "zh_CN": "客户端请求来源区域，以ISO-3166-1 两位国家代码表示。例如，'CN'代表'中国'。"}
  ClientRegions []*string `json:"clientRegions,omitempty" xml:"clientRegions,omitempty" type:"Repeated"`
  // {"en" : "Indicates the starting and ending IP addresses of the client requests to match against. They must be in IPv4 or IPv6 format.", "zh_CN": "用于指定要匹配的客户端请求的开始和结束IP地址。必须是IPv4或IPv6格式。"}
  ClientIpRange []*string `json:"clientIpRange,omitempty" xml:"clientIpRange,omitempty" type:"Repeated"`
}

func (s UpdateAPropertyVersionRequestAccessControlRulesConditions) String() string {
  return tea.Prettify(s)
}

func (s UpdateAPropertyVersionRequestAccessControlRulesConditions) GoString() string {
  return s.String()
}

func (s *UpdateAPropertyVersionRequestAccessControlRulesConditions) SetScheme(v string) *UpdateAPropertyVersionRequestAccessControlRulesConditions {
  s.Scheme = &v
  return s
}

func (s *UpdateAPropertyVersionRequestAccessControlRulesConditions) SetHostname(v string) *UpdateAPropertyVersionRequestAccessControlRulesConditions {
  s.Hostname = &v
  return s
}

func (s *UpdateAPropertyVersionRequestAccessControlRulesConditions) SetUri(v string) *UpdateAPropertyVersionRequestAccessControlRulesConditions {
  s.Uri = &v
  return s
}

func (s *UpdateAPropertyVersionRequestAccessControlRulesConditions) SetServerRegions(v []*string) *UpdateAPropertyVersionRequestAccessControlRulesConditions {
  s.ServerRegions = v
  return s
}

func (s *UpdateAPropertyVersionRequestAccessControlRulesConditions) SetClientRegions(v []*string) *UpdateAPropertyVersionRequestAccessControlRulesConditions {
  s.ClientRegions = v
  return s
}

func (s *UpdateAPropertyVersionRequestAccessControlRulesConditions) SetClientIpRange(v []*string) *UpdateAPropertyVersionRequestAccessControlRulesConditions {
  s.ClientIpRange = v
  return s
}

type UpdateAPropertyVersionRequestAccessControlRulesAction struct {
  // {"en" : "Indicates the HTTP status code to respond with. It must be in the range 300-309, 400-409, or 500-509 to indicate a redirection or error.", "zh_CN": "响应的HTTP状态码，范围必须在300-309、400-409或500-509之间，分别表示重定向或错误。"}
  StatusCode *string `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
  // {"en" : "Range: <= 200 characters 
  // If the value of the status field is in the range 300-309, the message field's value will be placed in a Location HTTP header returned to the client. If the status is in the range 400-409 or 500-509, the value will be returned in the response body to the client.", "zh_CN": "取值范围: <= 200 字符 
  // 如果status字段的值在300-309范围内，message字段的值将在返回给客户端的Location响应头中。如果status字段的值在400-409或500-509范围内，则message字段的值将在返回给客户端的响应体中。"}
  Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s UpdateAPropertyVersionRequestAccessControlRulesAction) String() string {
  return tea.Prettify(s)
}

func (s UpdateAPropertyVersionRequestAccessControlRulesAction) GoString() string {
  return s.String()
}

func (s *UpdateAPropertyVersionRequestAccessControlRulesAction) SetStatusCode(v string) *UpdateAPropertyVersionRequestAccessControlRulesAction {
  s.StatusCode = &v
  return s
}

func (s *UpdateAPropertyVersionRequestAccessControlRulesAction) SetMessage(v string) *UpdateAPropertyVersionRequestAccessControlRulesAction {
  s.Message = &v
  return s
}

type UpdateAPropertyVersionRequestExtraServicePorts struct {
  // {"en" : "This is a list of ports other than 80 which are used to handle HTTP requests. The available values can be found in the systemConfigs API's response. If you need another port, please contact our support team.", "zh_CN": "指定用于处理HTTP请求的端口列表（80端口除外）。可通过调用'获取系统配置'接口来查询系统支持的端口。如果您需要开通其他端口，请联系技术支持。"}
  Http []*string `json:"http,omitempty" xml:"http,omitempty" type:"Repeated"`
  // {"en" : "This is a list of ports other than 443 which are used to handle HTTPS requests. The available values can be found in the systemConfigs API's response. If you need another port, please contact our support team.", "zh_CN": "指定用于处理HTTPS请求的端口列表（443端口除外）。可通过调用'获取系统配置'接口来查询系统支持的端口。如果您需要开通其他端口，请联系技术支持。"}
  Https []*string `json:"https,omitempty" xml:"https,omitempty" type:"Repeated"`
}

func (s UpdateAPropertyVersionRequestExtraServicePorts) String() string {
  return tea.Prettify(s)
}

func (s UpdateAPropertyVersionRequestExtraServicePorts) GoString() string {
  return s.String()
}

func (s *UpdateAPropertyVersionRequestExtraServicePorts) SetHttp(v []*string) *UpdateAPropertyVersionRequestExtraServicePorts {
  s.Http = v
  return s
}

func (s *UpdateAPropertyVersionRequestExtraServicePorts) SetHttps(v []*string) *UpdateAPropertyVersionRequestExtraServicePorts {
  s.Https = v
  return s
}

type UpdateAPropertyVersionResponseHeader struct {
}

func (s UpdateAPropertyVersionResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s UpdateAPropertyVersionResponseHeader) GoString() string {
  return s.String()
}

type UpdateAPropertyVersionResponse struct {
}

func (s UpdateAPropertyVersionResponse) String() string {
  return tea.Prettify(s)
}

func (s UpdateAPropertyVersionResponse) GoString() string {
  return s.String()
}




type CreateAPropertyRequest struct {
  // {"en":"Name of the property.","zh_CN":"加速项目的名称。"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"Service type that hostnames in this property should associate with. One of wsapro, webpro, vodpro, downloadpro, cdnpro, 1523. Please select the service type that you've subscribed to.","zh_CN":"取值范围: wsapro,webpro,vodpro,downloadpro,cdnpro,MALL-cdnpro\n加速域名关联的商品（服务类型），即全站加速，网页加速，点播加速，下载加速，CDN Pro，以及自助CDN。请根据您已购买的商品进行选择。"}
  LegacyType *string `json:"legacyType,omitempty" xml:"legacyType,omitempty" require:"true"`
  // {"en":"Range: <= 250 characters\nDescription of the property.","zh_CN":"取值范围: <= 250 字符\n加速项目的描述。"}
  Description *string `json:"description,omitempty" xml:"description,omitempty"`
  // {"en":"Describes a property configuration. This contains all the settings.","zh_CN":"描述加速项目的所有配置信息。"}
  Version *CreateAPropertyRequestVersion `json:"version,omitempty" xml:"version,omitempty" require:"true" type:"Struct"`
}

func (s CreateAPropertyRequest) String() string {
  return tea.Prettify(s)
}

func (s CreateAPropertyRequest) GoString() string {
  return s.String()
}

func (s *CreateAPropertyRequest) SetName(v string) *CreateAPropertyRequest {
  s.Name = &v
  return s
}

func (s *CreateAPropertyRequest) SetLegacyType(v string) *CreateAPropertyRequest {
  s.LegacyType = &v
  return s
}

func (s *CreateAPropertyRequest) SetDescription(v string) *CreateAPropertyRequest {
  s.Description = &v
  return s
}

func (s *CreateAPropertyRequest) SetVersion(v *CreateAPropertyRequestVersion) *CreateAPropertyRequest {
  s.Version = v
  return s
}

type CreateAPropertyRequestVersion struct {
  // {"en":"A description of the version.","zh_CN":"版本描述。"}
  Description *string `json:"description,omitempty" xml:"description,omitempty"`
  // {"en":"Hostnames associated with the property. A valid value is a fully qualified hostname such as www.domain.com or a wildcard hostname such as.domain.com. Any given hostname can only be in one deployed property at a particular time. However, a wildcard hostname is permitted to overlap other hostnames you own.","zh_CN":"与加速项目关联的加速域名。必须是FQDN完全限定域名（如 www.domain.com），或泛域名（如*.domain.com）。\n同一个加速域名在同一时间只能存在于一个已部署的加速项目中，但泛域名可以与关联的完全限定域名一同部署。"}
  Hostnames []*string `json:"hostnames,omitempty" xml:"hostnames,omitempty" require:"true" type:"Repeated"`
  // {"en":"Default: False\nThe value indicates whether all the hostnames in this property have Beian license on file with the Chinese government. This is required to serve this property from servers in mainland China. A value of false means servers outside of mainland China will be used to distribute content to visitors in China. If set to true you must also specify the content type in the beianContentType field.","zh_CN":"默认值: False\n此加速项目中的所有加速域名是否都已获得ICP备案。只有获取备案的域名才能使用中国大陆节点来提供服务。如果设置为true，还必须在beianContentType字段中指定内容类型。"}
  HasBeian *bool `json:"hasBeian,omitempty" xml:"hasBeian,omitempty"`
  // {"en":"Range: [ 1 .. 24 ]\nIf you are planning to distribute content through servers in mainland China, you will need to set the hasBeian field to true and also indicate the type of content you are distributing. Enter the value that best corresponds to your content:\n1. Instant Messaging\n2. Search Engine\n3. Web Portal\n4. Online Postal Service\n5. Online News\n6. Blog\n7. Advertisement\n8. Organization Web Portal\n9. Shopping\n10. Online Payment\n11. Online Bank\n12. Online Stock Market\n13. Online Gaming\n14. Online Music\n15. Online Movie\n16. Online Picture\n17. Software Download\n18. Job Hunting\n19. Online Dating\n20. Online Real Property\n21. Online Education\n22. Web Design\n23. WAP\n24. Others","zh_CN":"取值范围: [ 1 .. 24 ]\n如果您的域名已获得备案，希望通过中国大陆的服务器分发内容，则需要将hasBeian字段设置为true，并说明您分发的内容类型。请选择与您的内容最匹配的值：\n1. 即时通信\n2. 搜索引擎\n3. 综合门户\n4. 网上邮局\n5. 网络新闻\n6. 博客/个人空间\n7. 网络广告\n8. 单位门户网站\n9. 网络购物\n10. 网上支付\n11. 网上银行\n12. 网上炒股\n13. 网络游戏\n14. 网络音乐\n15. 网络影视\n16. 网络图片\n17. 软件下载\n18. 网上求职\n19. 在线交友\n20. 网上房产\n21. 网络教育\n22. 网站建设\n23. WAP\n24. 其他"}
  BeianContentType *int `json:"beianContentType,omitempty" xml:"beianContentType,omitempty"`
  // {"en":"Describes the origin servers for the property's content.","zh_CN":"描述加速项目对应的源站。"}
  Origins []*CreateAPropertyRequestVersionOrigins `json:"origins,omitempty" xml:"origins,omitempty" type:"Repeated"`
  // {"en":"Default: False\nWhether to allow downgrading from HTTPS to HTTP when fetching content from origin. This setting applies only if the property has an attached certificate allowing client requests to use HTTPS. If the value of allowProtocolDowngrade is false and HTTPS is enabled, we require all the origin servers to support HTTPS. If the value is true, we allow origins that support only HTTP, but this reduces security level.","zh_CN":"默认值: False\n是否允许回源时从HTTPS降级为HTTP。仅当加速项目中指定了用于HTTPS请求的证书时，此设置才适用。如果allowProtocolDowngrade的值为false，且加速项目开启了HTTPS，则要求所有源服务器支持HTTPS。如果值为true，则允许仅支持HTTP的源，但这会降低安全性。"}
  AllowProtocolDowngrade *bool `json:"allowProtocolDowngrade,omitempty" xml:"allowProtocolDowngrade,omitempty"`
  // {"en":"Default: False\nEnables or disables using IPv6 addresses as origin.","zh_CN":"默认值: False\n是否开启IPv6回源"}
  EnableIpv6Origin *bool `json:"enableIpv6Origin,omitempty" xml:"enableIpv6Origin,omitempty"`
  // {"en":"Enum: auto, strict, off\nDefault: off\n\nWhen enableIpv6Origin is allowed, this setting will control whether to follow the client IP protocol version for back-to-origin requests.\nAuto: Attempt to reach the origin server using the same IP addressing as the client, but use the other version as backup if needed.\nStrict: Use only the same addressing method as the client when accessing the origin for content.\nOff: Use both IPv4 and IPv6 addressing to reach the origin, regardless of the IP addressing method used by the client.","zh_CN":"取值范围：auto, strict, off \n默认值：off \n当enableIpv6Origin为true时，该设置将控制是否跟随客户端IP协议版本回源。\nauto：表示根据客户端请求的IP协议版本来选择回源IP地址，当源站服务器不可用时自动切换成其他IP协议地址回源\nstrict：严格根据客户端请求的IP协议版本来选择回源IP地址，当源站服务器不可用时不允许切换成其他协议地址回源\noff：忽略客户端请求的IP协议版本，随机选择可用的IPv4或IPv6地址回源"}
  FollowClientIpVersion *string `json:"followClientIpVersion,omitempty" xml:"followClientIpVersion,omitempty"`
  // {"en":"Range: <= 65530 characters\nRefer to Edge Logic Introduction.","zh_CN":"取值范围: <= 65530 字符\n自定义边缘逻辑。参考边缘逻辑介绍。"}
  EdgeLogic *string `json:"edgeLogic,omitempty" xml:"edgeLogic,omitempty" require:"true"`
  // {"en":"Specify the syntax version that should be used when parsing the configuration. The value must be set to 1 at this time.","zh_CN":"指定配置解析适用的语法版本。当前仅允许值为1。"}
  SyntaxVersion *int `json:"syntaxVersion,omitempty" xml:"syntaxVersion,omitempty" require:"true"`
  // {"en":"Range: <= 80 characters\nThe cachekey hostname must be a string made of lowercase letters, digits, dot, dash, and underscore. If not specified, the incoming Host header value will be used and the cache will not be shared across different hostnames in this property.","zh_CN":"取值范围: <= 80\n用于缓存键（Cache Key）的主机名。必须是由小写字母、数字、点、连字符和下划线组成的字符串。如果未指定，默认将使用客户端传入的Host请求头值，并且此加速项目中的不同加速域名之间不共享缓存。"}
  CacheKeyHostname *string `json:"cacheKeyHostname,omitempty" xml:"cacheKeyHostname,omitempty"`
  // {"en":"Enum: preRewrite,postRewrite\nDefault: preRewrite\nThis setting controls how the URI of the incoming request is incorporated into the cache key if you use a 'rewrite' directive in the property's Edge Logic. The default value, 'preRewrite', puts the unmodified URI into the cache key while 'postRewrite' uses the modified URI. If your rewrite directive converts multiple different URIs into the same value, using 'postRewrite' may result in a higher cache hit ratio.","zh_CN":"取值范围: preRewrite,postRewrite\n默认值: preRewrite\n如果在加速项目的边缘逻辑中使用了'rewrite'指令，可使用该字段来控制客户端请求的URI如何合并到缓存键（Cache Key）中 。默认值'preRewrite'指将改写前的URI放入缓存键，而'postRewrite'则使用改写后的URI。如果您的'rewrite'指令将多个不同的URI改写为相同的值，则使用'postRewrite'可以提高缓存命中率。"}
  CacheKeyUri *string `json:"cacheKeyUri,omitempty" xml:"cacheKeyUri,omitempty"`
  // {"en":"Default: False\nSpecifies whether the scheme ('http' or 'https') should be included in the cache key. Default behavior is false. Set to true if the content is known to be different for different schemes.","zh_CN":"默认值: False\n指定缓存键（Cache Key）是否区分协议（'http'或'https'）。默认为false。如果不同协议响应的内容不同，则设置为true。"}
  SchemeInCacheKey *bool `json:"schemeInCacheKey,omitempty" xml:"schemeInCacheKey,omitempty"`
  // {"en":"Default: False\nThis field can be set to either a boolean value or a string. If the value is set to true, the server will redirect all HTTP requests to HTTPS using status code 301. You can also specify string values '302', '307', or '308' instead if you wish to return a different status code when a request is redirected.","zh_CN":"默认值: False\n此字段可以设置为布尔值或字符串。如果设置为true，则CDN Pro服务器会将所有HTTP请求重定向到HTTPS，并返回301状态码。如果您希望在重定向请求时返回不同的状态码，可在此处指定需要的状态码，如'302'、'307'或'308'。"}
  RedirectHttpToHttps *string `json:"redirectHttpToHttps,omitempty" xml:"redirectHttpToHttps,omitempty"`
  // {"en":"Default: False\nSet to true to disable support for HTTP/2 and support only HTTP 1.1.","zh_CN":"默认值: False\n当值为true时，禁用对HTTP/2的支持，仅支持HTTP/1.1。"}
  DisableHttp2 *bool `json:"disableHttp2,omitempty" xml:"disableHttp2,omitempty"`
  // {"en":"Default: False\nSet to true to enable support for HTTP/3. HTTP/3 requires TLS to work, so this field will be ignored unless you specify a certificate in the tlsCertificateId field.","zh_CN":"默认值: False\n是否开启HTTP/3。需要配置TLS才生效，若您未在 tlsCertificateId 中指定证书，此字段将被忽略。"}
  EnableHttp3 *bool `json:"enableHttp3,omitempty" xml:"enableHttp3,omitempty"`
  // {"en":"This field lists ports other than the default 80 used to handle HTTP requests and ports other than the default 443 used to handle HTTPS requests.","zh_CN":"除标准的80，443端口外，我们还支持一些扩展端口。可用该字段指定用于处理HTTP和HTTPS请求的扩展端口。"}
  ExtraServicePorts *CreateAPropertyRequestVersionExtraServicePorts `json:"extraServicePorts,omitempty" xml:"extraServicePorts,omitempty" type:"Struct"`
  // {"en":"Specify a certificate to enable HTTPS for your hostnames.","zh_CN":"指定用于HTTPS请求的证书的ID。如果未指定证书，则不会为此加速项目启用HTTPS。"}
  TlsCertificateId *string `json:"tlsCertificateId,omitempty" xml:"tlsCertificateId,omitempty"`
  // {"en":"Specify a second certificate to be used for HTTPS. With tlsCertificateId and tlsCertificateId1, you can specify two certificates of different types, i.e. one RSA and one EC. When there are two certificates available for your hostnames, the client and server will negotiate at the time of requests to pick one of the certificates for SSL. This could mean better performance and fault tolerance.\n\nIt is invalid to only set the tlsCertificateId1 field without setting the tlsCertificateId field. The certificates set for tlsCertificateId and tlsCertificateId1 must be of different type.","zh_CN":"指定用于HTTPS请求的第二本证书。必须先指定tlsCertificateId，才能指定tlsCertificateId1。tlsCertificateId和tlsCertificateId1所指定的证书必须是不同证书类型。\n\ntlsCertificateId和tlsCertificateId1允许您指定两本不同类型的证书，即一本RSA证书，一本EC证书。如果您指定了多本证书，当客户端发起请求时，客户端和服务端可以协商选择使用其中一本证书。"}
  TlsCertificateId1 *string `json:"tlsCertificateId1,omitempty" xml:"tlsCertificateId1,omitempty"`
  // {"en":"Enum: 1.1,1.2,1.3\nDefault: 1.3\nMaximum supported TLS version.","zh_CN":"取值范围: 1.1,1.2,1.3\n默认值: 1.3\n支持的TLS最高版本。"}
  TlsMaxVersion *string `json:"tlsMaxVersion,omitempty" xml:"tlsMaxVersion,omitempty"`
  // {"en":"Enum: 1.1,1.2,1.3,1\nDefault: 1\nMinimum required TLS version.","zh_CN":"取值范围: 1.1,1.2,1.3,1\n默认值: 1\n所需的TLS最低版本。"}
  TlsMinVersion *string `json:"tlsMinVersion,omitempty" xml:"tlsMinVersion,omitempty"`
  // {"en":"Range: <= 2040 characters\nAny cipher list supported by 'openssl ciphers'. See https://www.openssl.org/docs/manmaster/man1/ciphers.html","zh_CN":"取值范围: <= 2040 字符\n'openssl ciphers'支持的任何加密算法套件。参考：https://www.openssl.org/docs/manmaster/man1/ciphers.html"}
  TlsCiphers *string `json:"tlsCiphers,omitempty" xml:"tlsCiphers,omitempty"`
  // {"en":"Default: 1800 Range: [ 300 .. 86400 ]\nLifespan of TLS session ticket in seconds.","zh_CN":"默认值: 1800 取值范围: [ 300 .. 86400 ]\nTLS会话ticket的有效期（秒）。"}
  TlsSessionTimeout *int `json:"tlsSessionTimeout,omitempty" xml:"tlsSessionTimeout,omitempty"`
  // {"en":"Default: False\nEnable TLS zero round-trip time, a feature of TLS 1.3 that can improve performance for repeat visits to a site. If enabling it though, be sure your site is not vulnerable to the HTTP replay attack.","zh_CN":"默认值: False\n是否开启TLS 0-RTT功能。这是TLS 1.3版本支持的功能。启用该功能后，当用户频繁访问您的站点时，可提高访问响应速度。需要注意的是，启用该功能可能会加大站点遭受HTTP replay攻击的风险。"}
  Tls0Rtt *bool `json:"tls0Rtt,omitempty" xml:"tls0Rtt,omitempty"`
  // {"en":"Default: False\nEnables Online Certificate Status Protocol (OCSP) stapling to check the revocation status of digital certificates. (Refer to https://en.wikipedia.org/wiki/OCSP_stapling)","zh_CN":"默认值: False\n是否启用在线证书状态协议（OCSP）装订以检查数字证书的吊销状态。参考：https://en.wikipedia.org/wiki/OCSP_stapling"}
  EnableOcspStapling *bool `json:"enableOcspStapling,omitempty" xml:"enableOcspStapling,omitempty"`
  // {"en":"Default: False\nBy default, CDN Pro takes control of the contents under the /.well-known/{acme-challenge, pki-validation} directories to support certificate auto-renew for properties. If for any reason you need to manage these two directories by yourself on the origin, for example, to implement your own certificate auto-renew mechanism, you can use this configuration option to disable the default behavior by setting its value to true. For more information about our support for certificate auto-renewal, refer to the description of the autoRenew field in the Create a certificate API.","zh_CN":"默认值: False\n默认情况下，CDN Pro控制/.well-known/{acme-challenge, pki-validation}目录下的内容，以支持加速项目的证书自动更新功能。如果您需要自己在源站管理这两个目录，例如，为了实现您自己的证书自动更新机制，您可以将此字段设置为true来禁用默认行为。关于证书自动更新的更多信息，请参考'创建证书'接口中autoRenew字段的说明。"}
  DisableCertAutomation *bool `json:"disableCertAutomation,omitempty" xml:"disableCertAutomation,omitempty"`
  // {"en":"Specify one or more access control rules to restrict access to your content. More advanced configuration can be done using Edge Logic. These access control rules take precedence over Edge Logic if both are defined.","zh_CN":"指定一个或多个访问控制规则以限制对内容的访问。可以使用边缘逻辑进行更高级的配置。此处定义的访问控制规则，优先级高于边缘逻辑。"}
  AccessControlRules []*CreateAPropertyRequestVersionAccessControlRules `json:"accessControlRules,omitempty" xml:"accessControlRules,omitempty" type:"Repeated"`
  // {"en":"This optional field allows you to configure notifications about client requests to be sent to a remote server. This is an advanced feature. Please contact our support team if you require this feature.","zh_CN":"此可选字段用来配置发送消息通知（即实时日志）到您的远程服务器。当有客户端请求访问您的加速域名时，将触发通知。这是高级功能，如果您需要此功能，请联系我们的技术支持开通。"}
  RealTimeLog *CreateAPropertyRequestVersionRealTimeLog `json:"realTimeLog,omitempty" xml:"realTimeLog,omitempty" type:"Struct"`
  // {"en":"This object allows you to support video players requesting partial content through query string parameters. If you specify videoSeek, you must enter a value for startParameter.","zh_CN":"此对象用来支持视频播放器通过指定查询参数来请求部分内容。当videoSeek对象存在时，必须为startParameter设置一个值。"}
  VideoSeek *CreateAPropertyRequestVersionVideoSeek `json:"videoSeek,omitempty" xml:"videoSeek,omitempty" type:"Struct"`
  // {"en":"Range: <= 65530 characters\nThis field allows you to customize load balancing. The field will be deprecated soon. Please use the field edgeLogic.","zh_CN":"取值范围: <= 65530 字符\n用于控制负载均衡器的行为。该字段即将废弃，请使用edgeLogic字段进行配置。"}
  LoadBalancerLogic *string `json:"loadBalancerLogic,omitempty" xml:"loadBalancerLogic,omitempty"`
  // {"en":"Range: <= 100 characters\nUsed to customize hash key for load balancer. The field will soon be deprecated. Please stop using it.","zh_CN":"取值范围: <= 100 字符\n用于自定义负载均衡器的哈希key。该字段即将废弃，请勿使用该字段。"}
  LoadBalancerHashKey *string `json:"loadBalancerHashKey,omitempty" xml:"loadBalancerHashKey,omitempty"`
}

func (s CreateAPropertyRequestVersion) String() string {
  return tea.Prettify(s)
}

func (s CreateAPropertyRequestVersion) GoString() string {
  return s.String()
}

func (s *CreateAPropertyRequestVersion) SetDescription(v string) *CreateAPropertyRequestVersion {
  s.Description = &v
  return s
}

func (s *CreateAPropertyRequestVersion) SetHostnames(v []*string) *CreateAPropertyRequestVersion {
  s.Hostnames = v
  return s
}

func (s *CreateAPropertyRequestVersion) SetHasBeian(v bool) *CreateAPropertyRequestVersion {
  s.HasBeian = &v
  return s
}

func (s *CreateAPropertyRequestVersion) SetBeianContentType(v int) *CreateAPropertyRequestVersion {
  s.BeianContentType = &v
  return s
}

func (s *CreateAPropertyRequestVersion) SetOrigins(v []*CreateAPropertyRequestVersionOrigins) *CreateAPropertyRequestVersion {
  s.Origins = v
  return s
}

func (s *CreateAPropertyRequestVersion) SetAllowProtocolDowngrade(v bool) *CreateAPropertyRequestVersion {
  s.AllowProtocolDowngrade = &v
  return s
}

func (s *CreateAPropertyRequestVersion) SetEnableIpv6Origin(v bool) *CreateAPropertyRequestVersion {
  s.EnableIpv6Origin = &v
  return s
}

func (s *CreateAPropertyRequestVersion) SetFollowClientIpVersion(v string) *CreateAPropertyRequestVersion {
  s.FollowClientIpVersion = &v
  return s
}

func (s *CreateAPropertyRequestVersion) SetEdgeLogic(v string) *CreateAPropertyRequestVersion {
  s.EdgeLogic = &v
  return s
}

func (s *CreateAPropertyRequestVersion) SetSyntaxVersion(v int) *CreateAPropertyRequestVersion {
  s.SyntaxVersion = &v
  return s
}

func (s *CreateAPropertyRequestVersion) SetCacheKeyHostname(v string) *CreateAPropertyRequestVersion {
  s.CacheKeyHostname = &v
  return s
}

func (s *CreateAPropertyRequestVersion) SetCacheKeyUri(v string) *CreateAPropertyRequestVersion {
  s.CacheKeyUri = &v
  return s
}

func (s *CreateAPropertyRequestVersion) SetSchemeInCacheKey(v bool) *CreateAPropertyRequestVersion {
  s.SchemeInCacheKey = &v
  return s
}

func (s *CreateAPropertyRequestVersion) SetRedirectHttpToHttps(v string) *CreateAPropertyRequestVersion {
  s.RedirectHttpToHttps = &v
  return s
}

func (s *CreateAPropertyRequestVersion) SetDisableHttp2(v bool) *CreateAPropertyRequestVersion {
  s.DisableHttp2 = &v
  return s
}

func (s *CreateAPropertyRequestVersion) SetEnableHttp3(v bool) *CreateAPropertyRequestVersion {
  s.EnableHttp3 = &v
  return s
}

func (s *CreateAPropertyRequestVersion) SetExtraServicePorts(v *CreateAPropertyRequestVersionExtraServicePorts) *CreateAPropertyRequestVersion {
  s.ExtraServicePorts = v
  return s
}

func (s *CreateAPropertyRequestVersion) SetTlsCertificateId(v string) *CreateAPropertyRequestVersion {
  s.TlsCertificateId = &v
  return s
}

func (s *CreateAPropertyRequestVersion) SetTlsCertificateId1(v string) *CreateAPropertyRequestVersion {
  s.TlsCertificateId1 = &v
  return s
}

func (s *CreateAPropertyRequestVersion) SetTlsMaxVersion(v string) *CreateAPropertyRequestVersion {
  s.TlsMaxVersion = &v
  return s
}

func (s *CreateAPropertyRequestVersion) SetTlsMinVersion(v string) *CreateAPropertyRequestVersion {
  s.TlsMinVersion = &v
  return s
}

func (s *CreateAPropertyRequestVersion) SetTlsCiphers(v string) *CreateAPropertyRequestVersion {
  s.TlsCiphers = &v
  return s
}

func (s *CreateAPropertyRequestVersion) SetTlsSessionTimeout(v int) *CreateAPropertyRequestVersion {
  s.TlsSessionTimeout = &v
  return s
}

func (s *CreateAPropertyRequestVersion) SetTls0Rtt(v bool) *CreateAPropertyRequestVersion {
  s.Tls0Rtt = &v
  return s
}

func (s *CreateAPropertyRequestVersion) SetEnableOcspStapling(v bool) *CreateAPropertyRequestVersion {
  s.EnableOcspStapling = &v
  return s
}

func (s *CreateAPropertyRequestVersion) SetDisableCertAutomation(v bool) *CreateAPropertyRequestVersion {
  s.DisableCertAutomation = &v
  return s
}

func (s *CreateAPropertyRequestVersion) SetAccessControlRules(v []*CreateAPropertyRequestVersionAccessControlRules) *CreateAPropertyRequestVersion {
  s.AccessControlRules = v
  return s
}

func (s *CreateAPropertyRequestVersion) SetRealTimeLog(v *CreateAPropertyRequestVersionRealTimeLog) *CreateAPropertyRequestVersion {
  s.RealTimeLog = v
  return s
}

func (s *CreateAPropertyRequestVersion) SetVideoSeek(v *CreateAPropertyRequestVersionVideoSeek) *CreateAPropertyRequestVersion {
  s.VideoSeek = v
  return s
}

func (s *CreateAPropertyRequestVersion) SetLoadBalancerLogic(v string) *CreateAPropertyRequestVersion {
  s.LoadBalancerLogic = &v
  return s
}

func (s *CreateAPropertyRequestVersion) SetLoadBalancerHashKey(v string) *CreateAPropertyRequestVersion {
  s.LoadBalancerHashKey = &v
  return s
}

type CreateAPropertyRequestVersionOrigins struct     {
  // {"en":"^[a-zA-z0-9_]\nName of an origin. It must be unique within this property.","zh_CN":"^[a-zA-z0-9_]\n源站名称，在此加速项目中必须唯一。"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"Each entry should be a string consisting of an address optionally followed by one or more parameters, separated by space characters. The address can take one of the following three forms:\n{hostname or IP address}\n{hostname or IP address}:{http port}\n{hostname or IP address}:{http port},{https port}\nValues for the HTTP and HTTPS ports should be integers in the range [1,65535].\nIn the third form, one of the two ports can be empty, but the HTTPS port must always be specified after the comma.\n\nAccepts ipv6 addresses if 'enableIpv6Origin' is set to true, but an ipv6 address must be enclosed in [] so as to tell the port apart.\n\nSupported parameters are 'backup' and 'weight'.\n\n'backup' is used to indicate the server is a backup server. It will be used when the primary servers are unavailable.\n\n'weight' is a value in the range [1, 100]. The default value is 1. It lets you control the likelihood that one origin server will be used relative to another.\n\nThere should be at least one primary server defined; it does not have the 'backup' parameter.\n\nExample:\n['www.abc.com weight=1', 'www1.abc.com:8080 weight=4', 'www2.abc.com:880,4443 backup']\n\nThis example specifies two primary servers which are used in a 1:4 ratio meaning one gets about 20% of the requests to origin while the other gets 80% of the requests. In addition, a backup server is specified.","zh_CN":"源站服务器列表。每个条目为一个字符串，由一个地址与一个或多个参数组成，参数之间以空格分隔。地址可以采用以下三种形式之一：\n{域名或IP地址}\n{域名或IP地址}:{HTTP端口}\n{域名或IP地址}:{HTTP端口},{HTTPS端口}\nHTTP和HTTPS端口的值应为[1,65535]范围内的整数。\n在第三种形式中，HTTPS端口必须固定写在逗号后。如果HTTP端口未指定，仍需保留逗号，写为“,{HTTPS端口}”。\n\n当enableIpv6Origin设为true时，此处支持指定ipv6地址。ipv6地址应包裹在[]中，以便和端口区分开，例如[2001:0db8:85a3:0000:0000:8a2e:0370:7334]:80\n\n支持的参数包括'backup'和'weight'。\n'backup'用于标识备份服务器。它将在主服务器不可用时使用。\n'weight'是范围[1,100]内的值，默认值为1，用来设置服务器权重，控制一台源站服务器相对于另一台被使用的可能性。\n每个源站应至少定义一个主服务器（即不带'backup'参数的服务器）。\n\n示例：\n['www.abc.com weight=1', 'www1.abc.com:8080 weight=4', 'www2.abc.com:880,4443 backup']\n\n此示例指定了两个主服务器和一个备份服务器，其中，两个主服务器的权重为1:4，表示第一个服务器将获得约20%的回源请求，而另一个将获得约80%。"}
  Servers []*string `json:"servers,omitempty" xml:"servers,omitempty" require:"true" type:"Repeated"`
  // {"en":"Enum: http,https,both\nThis optional field indicates the protocol supported by the origin server. If this property has a certificate attached to it, the value can be set to http if the allowProtocolDowngrade field is also set to true.","zh_CN":"取值范围: http,https,both\n此可选字段表示源站服务器支持的协议。如果此加速项目附加了证书，且allowProtocolDowngrade字段也设置为true，则可以将该值设置为http。"}
  SupportedProtocol *string `json:"supportedProtocol,omitempty" xml:"supportedProtocol,omitempty"`
  // {"en":"Default: auto\nOptional. This parameter tells us how important it is to directly go to this origin without going through any intermediate cache to fetch content. The values can be 'noDirect', 'auto', 'alwaysDirect'.\n\n'noDirect' means we always go through at least one intermediate cache. Customers who care particularly about the origin's workload can use this option. Our cache scheduler will dynamically pick the intermediate cache based on performance and resource availability.\n\n'auto' means the cache scheduler will make the choice dynamically based on cacheability of content. This is the default behavior if not specified.\n\n'alwaysDirect' means we always directly connect to this origin.","zh_CN":"默认值: auto\n此可选字段用来指定回源方式，可以是'noDirect'、'auto'、'alwaysDirect'。\n\n'noDirect'指总是通过至少一个中间缓存节点回源。对于特别关心源站负载的客户可以使用此选项。我们的调度程序将根据性能和资源可用性动态选择中间缓存节点。\n'alwaysDirect'指总是直接回源。\n'auto'指自动根据内容是否可缓存选择直接回源或通过中间缓存节点回源。对于可缓存的内容，将通过中间缓存节点回源；对于不可缓存的内容，将直接回源。该字段未指定时，采用此默认行为。"}
  DirectConnection *string `json:"directConnection,omitempty" xml:"directConnection,omitempty"`
  // {"en":"Optional. It should be a valid hostname. It will also be used as the SNI servername when contacting the HTTPS origin. We also allow it to be an nginx variable that begins with '$' followed by [a-z_]{3,60}. Nginx variable names are case insensitive, so we only allow lowercase characters.\nIf not specified, the value of the HOST header in the request will be used.\n\nOne constraint is that if 'aswS3' is specified for origin authentication, the value of hostHeader cannot be a variable or empty. It has to be a valid Amazon S3 hostname. Refer to https://docs.aws.amazon.com/AmazonS3/latest/dev/VirtualHosting.html","zh_CN":"可选，用来指定回源HOST头部。必须是有效的域名。当连接HTTPS源站时，该值也作为SNI域名。可以用nginx变量来指定，变量以'$'开头，后跟[a-z_]{3,60}。Nginx变量名不区分大小写，因此我们只允许小写字符。\n如果未指定，将使用请求中的HOST请求头的值。\n\n注意：如果指定了'awsS3'作为回源鉴权参数，则hostHeader的值不能为变量或为空，而必须是有效的Amazon S3主机名。参考：https://docs.aws.amazon.com/AmazonS3/latest/dev/VirtualHosting.html"}
  HostHeader *string `json:"hostHeader,omitempty" xml:"hostHeader,omitempty"`
  // {"en":"Default: False\nOptional. It controls whether the cache will validate the certificate of the origin.","zh_CN":"默认值: False\n可选。用来设定CDN缓存节点是否验证源站证书。"}
  VerifyOrigin *bool `json:"verifyOrigin,omitempty" xml:"verifyOrigin,omitempty"`
  // {"en":"Optional. It is a structure to specify the parameters, such as password, for authentication with the origin servers. It should have at least one field: 'methodName', which should be a string indicating one of the predefined authentication methods. The only valid value at this time is 'awsS3'. When 'awsS3' is used, the following parameters are required:\nawsS3.region:\nThe AWS S3 region where your resources are hosted, e.g. 'us-west-1', 'ap-northeast-1', 'eu-north-1', and 'cn-north-1'.\nawsS3.accessKey:\nYour AWS access key that CDN Pro will use to access your resources stored on AWS S3.\nawsS3.secretKey:\nYour AWS secret key that CDN Pro will use to access your resources stored on AWS S3.\nAlso, the value of the hostHeader field cannot be a variable or empty. It must be a valid Amazon S3 hostname.\n\nExample:\n<code>{\"methodName\":\"awsS3\",\n\"awsS3\":{\n\"region\":\"us-west-1\",\n\"accessKey\":\"sdnu2932\",\n\"secretKey\":\"d12345678abcdefghi\"\n}}</code>","zh_CN":"可选。用于指定与源服务器进行身份验证（回源鉴权）的相关参数（如密码）。必须至少有'methodName'（字符串）字段，用来指定预定义的鉴权方法。目前仅支持源站为AWS S3的回源鉴权，因此唯一有效的值是'awsS3'。使用'awsS3'时，需要提供以下参数：\nawsS3.region:\n您存储在S3上的资源所在的AWS区域。例如，'us-west-1'，'ap-northeast-1'，'eu-north-1'，'cn-north-1'。\nawsS3.accessKey:\n您的 AWS 访问密钥（access key），CDN Pro 将使用它来访问您存储在 AWS S3 上的资源。\nawsS3.secretKey:\n您的 AWS 密钥（secret key），CDN Pro 将使用它来访问您存储在 AWS S3 上的资源。\n此外，hostHeader字段的值不能为变量或为空，必须是有效的Amazon S3主机名。\n\n示例：\n<code>{\"methodName\":\"awsS3\",\n\"awsS3\":{\n\"region\":\"us-west-1\",\n\"accessKey\":\"sdnu2932\",\n\"secretKey\":\"d12345678abcdefghi\"\n}}</code>"}
  Authentication *CreateAPropertyRequestVersionOriginsAuthentication `json:"authentication,omitempty" xml:"authentication,omitempty" type:"Struct"`
  // {"en":"Default: 60 Range: [ 5 .. 600 ]\nTimeout in seconds during which an idle keepalive connection to an upstream server will stay open. A service quota setting of maxUpstreamKeepaliveTimeOut can change the maximum permitted value.","zh_CN":"默认值: 60 取值范围: [ 5 .. 600 ]\n该字段用于指定CDN Pro服务器和源站建连的Keep-Alive超时时间，单位为秒。通过maxUpstreamKeepaliveTimeOut 该服务设置项可以更改允许的最大值。如果需要调整最大值，请联系我们的技术支持。"}
  KeepAliveTimeout *int `json:"keepAliveTimeout,omitempty" xml:"keepAliveTimeout,omitempty"`
  // {"en":"This setting allows you to specify the number of unsuccessful attempts to reach one of the origin's IP addresses or peers before it is marked as unavailable for a period of time designated by the timeout. If all peers of all servers are unavailable, requests for content from the origin will receive an immediate 502 HTTP response. By default, six attempts to a peer are made, after which a one-second timeout applies to that peer. To disable this feature, set peerFailureTimeout to 'off'.\nExample:\n<code>{\"peerFailureTimeout\":{\"failureThreshold\":10,\"timeout\":3}}</code>","zh_CN":"该字段用于设置源站故障剔除。开启此功能后，当连接某个源站服务器的失败次数达到设定阈值，该源站服务器将被标记为不可用，并保持该状态直到设定的超时时长。失败次数阈值和超时时长分别通过failureThreshold参数和timeout参数设置。CDN Pro回源时将剔除不可用的源站服务器。如果所有源站服务器都被标记为不可用，则对应的请求将立即响应502状态码。默认情况下，当连接某个源站服务器连续失败6次之后，该服务器将被标记为不可用，超时时长为1秒。如果要禁用此功能，请将peerFailureTimeout设置为'off'。开启源站故障剔除配置示例：<code>{\"peerFailureTimeout\":{\"failureThreshold\":10,\"timeout\":3}}</code>"}
  PeerFailureTimeout *string `json:"peerFailureTimeout,omitempty" xml:"peerFailureTimeout,omitempty"`
  // {"en":"Refers to a certificate used to authenticate with the origin server. This certificate must also be deployed.","zh_CN":"用于源站验证CDN Pro服务器的证书，该证书同样必须被部署。"}
  TlsCertificateId *string `json:"tlsCertificateId,omitempty" xml:"tlsCertificateId,omitempty"`
  // {"en":"Refers to the ID of an origin shield representing a set of servers that will make requests to your origin servers rather than the CDN Pro edge servers. This can help further reduce traffic to your origin. Origin shield is allowed only when direcConnection is set to noDirect. This is an advanced feature that can be enabled by contacting our support team. The origin shields API provides a list of available shields along with their IP addresses. It is best to select a shield whose region is closest to your origin servers. Use of a shield in China requires your property have 'hasBeian' set to true. If your origin servers have an access control list, add the IPs of the shield you choose to use.","zh_CN":"指定某个源站盾（origin shield）的ID。使用源站盾功能后，所有回源请求都会通过源站盾中转回源，这可以帮助收敛回源流量。需要把directConnection设置为noDirect，才允许开启源站盾。源站盾是高级功能，如需使用请联系我们的技术支持开通。可通过调用‘获取源站盾列表’接口查询可用的源站盾及其对应的IP地址。您可根据源站的位置，选择合适的源站盾。如果您需使用中国大陆的源站盾，则该加速项目的所有域名必须已取得备案。如果您的源站开启了访问控制，请将所选择源站盾的IP地址加入白名单。"}
  Shield *string `json:"shield,omitempty" xml:"shield,omitempty"`
}

func (s CreateAPropertyRequestVersionOrigins) String() string {
  return tea.Prettify(s)
}

func (s CreateAPropertyRequestVersionOrigins) GoString() string {
  return s.String()
}

func (s *CreateAPropertyRequestVersionOrigins) SetName(v string) *CreateAPropertyRequestVersionOrigins {
  s.Name = &v
  return s
}

func (s *CreateAPropertyRequestVersionOrigins) SetServers(v []*string) *CreateAPropertyRequestVersionOrigins {
  s.Servers = v
  return s
}

func (s *CreateAPropertyRequestVersionOrigins) SetSupportedProtocol(v string) *CreateAPropertyRequestVersionOrigins {
  s.SupportedProtocol = &v
  return s
}

func (s *CreateAPropertyRequestVersionOrigins) SetDirectConnection(v string) *CreateAPropertyRequestVersionOrigins {
  s.DirectConnection = &v
  return s
}

func (s *CreateAPropertyRequestVersionOrigins) SetHostHeader(v string) *CreateAPropertyRequestVersionOrigins {
  s.HostHeader = &v
  return s
}

func (s *CreateAPropertyRequestVersionOrigins) SetVerifyOrigin(v bool) *CreateAPropertyRequestVersionOrigins {
  s.VerifyOrigin = &v
  return s
}

func (s *CreateAPropertyRequestVersionOrigins) SetAuthentication(v *CreateAPropertyRequestVersionOriginsAuthentication) *CreateAPropertyRequestVersionOrigins {
  s.Authentication = v
  return s
}

func (s *CreateAPropertyRequestVersionOrigins) SetKeepAliveTimeout(v int) *CreateAPropertyRequestVersionOrigins {
  s.KeepAliveTimeout = &v
  return s
}

func (s *CreateAPropertyRequestVersionOrigins) SetPeerFailureTimeout(v string) *CreateAPropertyRequestVersionOrigins {
  s.PeerFailureTimeout = &v
  return s
}

func (s *CreateAPropertyRequestVersionOrigins) SetTlsCertificateId(v string) *CreateAPropertyRequestVersionOrigins {
  s.TlsCertificateId = &v
  return s
}

func (s *CreateAPropertyRequestVersionOrigins) SetShield(v string) *CreateAPropertyRequestVersionOrigins {
  s.Shield = &v
  return s
}

type CreateAPropertyRequestVersionOriginsAuthentication struct {
  // {"en":"Authentication method.","zh_CN":"鉴权方法。"}
  MethodName *string `json:"methodName,omitempty" xml:"methodName,omitempty"`
}

func (s CreateAPropertyRequestVersionOriginsAuthentication) String() string {
  return tea.Prettify(s)
}

func (s CreateAPropertyRequestVersionOriginsAuthentication) GoString() string {
  return s.String()
}

func (s *CreateAPropertyRequestVersionOriginsAuthentication) SetMethodName(v string) *CreateAPropertyRequestVersionOriginsAuthentication {
  s.MethodName = &v
  return s
}

type CreateAPropertyRequestVersionExtraServicePorts struct {
  // {"en":"This is a list of ports other than 80 which are used to handle HTTP requests. The available values can be found in the systemConfigs API's response. If you need another port, please contact our support team.","zh_CN":"指定用于处理HTTP请求的端口列表（80端口除外）。可通过调用'获取系统配置'接口来查询系统支持的端口。如果您需要开通其他端口，请联系技术支持。"}
  Http []*string `json:"http,omitempty" xml:"http,omitempty" type:"Repeated"`
  // {"en":"This is a list of ports other than 443 which are used to handle HTTPS requests. The available values can be found in the systemConfigs API's response. If you need another port, please contact our support team.","zh_CN":"指定用于处理HTTPS请求的端口列表（443端口除外）。可通过调用'获取系统配置'接口来查询系统支持的端口。如果您需要开通其他端口，请联系技术支持。"}
  Https []*string `json:"https,omitempty" xml:"https,omitempty" type:"Repeated"`
}

func (s CreateAPropertyRequestVersionExtraServicePorts) String() string {
  return tea.Prettify(s)
}

func (s CreateAPropertyRequestVersionExtraServicePorts) GoString() string {
  return s.String()
}

func (s *CreateAPropertyRequestVersionExtraServicePorts) SetHttp(v []*string) *CreateAPropertyRequestVersionExtraServicePorts {
  s.Http = v
  return s
}

func (s *CreateAPropertyRequestVersionExtraServicePorts) SetHttps(v []*string) *CreateAPropertyRequestVersionExtraServicePorts {
  s.Https = v
  return s
}

type CreateAPropertyRequestVersionAccessControlRules struct     {
  // {"en":"Range: [ 0 .. 60 ] characters\nAn optional ID for the access control rule.","zh_CN":"取值范围: [ 0 .. 60 ] 字符\n访问控制规则ID。"}
  Id *string `json:"id,omitempty" xml:"id,omitempty"`
  // {"en":"Specify the conditions that the incoming request must match. At least one condition must be specified. If multiple are specified, all must match.","zh_CN":"指定客户端请求必须匹配的条件。必须至少指定一个条件。如果指定了多个条件，则必须全部匹配。"}
  Conditions *CreateAPropertyRequestVersionAccessControlRulesConditions `json:"conditions,omitempty" xml:"conditions,omitempty" require:"true" type:"Struct"`
  // {"en":"Indicates the action to take in response to a request that matches the conditions of the access control rule.","zh_CN":"对于匹配到以上条件的请求所采取的相应操作。"}
  Action *CreateAPropertyRequestVersionAccessControlRulesAction `json:"action,omitempty" xml:"action,omitempty" require:"true" type:"Struct"`
}

func (s CreateAPropertyRequestVersionAccessControlRules) String() string {
  return tea.Prettify(s)
}

func (s CreateAPropertyRequestVersionAccessControlRules) GoString() string {
  return s.String()
}

func (s *CreateAPropertyRequestVersionAccessControlRules) SetId(v string) *CreateAPropertyRequestVersionAccessControlRules {
  s.Id = &v
  return s
}

func (s *CreateAPropertyRequestVersionAccessControlRules) SetConditions(v *CreateAPropertyRequestVersionAccessControlRulesConditions) *CreateAPropertyRequestVersionAccessControlRules {
  s.Conditions = v
  return s
}

func (s *CreateAPropertyRequestVersionAccessControlRules) SetAction(v *CreateAPropertyRequestVersionAccessControlRulesAction) *CreateAPropertyRequestVersionAccessControlRules {
  s.Action = v
  return s
}

type CreateAPropertyRequestVersionAccessControlRulesConditions struct {
  // {"en":"Enum: https,http\nIndicates whether the incoming request uses HTTP or HTTPS.","zh_CN":"取值范围: https,http\n客户端请求的协议，HTTP或HTTPS。"}
  Scheme *string `json:"scheme,omitempty" xml:"scheme,omitempty"`
  // {"en":"Indicates the hostname requested. It must be one of the hostnames defined in the property. If a wildcard hostname is configured in the property, you can specify either the wildcard hostname or their spcific hostnames here.","zh_CN":"客户端请求对应的域名。必须是加速项目中定义的加速域名之一。如果加速域名按泛域名配置，此处可以指定泛域名或者泛域名下的某个具体域名。"}
  Hostname *string `json:"hostname,omitempty" xml:"hostname,omitempty"`
  // {"en":"Range: <= 500 characters\nThe URI begins with '/' and can end with '*' for prefix matching.","zh_CN":"取值范围: <= 500 字符\n用于前缀匹配的URI，以'/'开头，可以以'*'结尾，表示模糊匹配。"}
  Uri *string `json:"uri,omitempty" xml:"uri,omitempty"`
  // {"en":"An array of ISO-3166-1 alpha-2 codes indicating the countries of the servers, for example, 'US' to represent 'United States of America'.","zh_CN":"服务器所在区域，以ISO-3166-1 两位国家代码表示。例如，'US'代表'美国'。"}
  ServerRegions []*string `json:"serverRegions,omitempty" xml:"serverRegions,omitempty" type:"Repeated"`
  // {"en":"An array of ISO-3166-1 alpha-2 codes indicating one or more countries which are the source of the client requests, for example, 'CN' to represent 'China'.","zh_CN":"客户端请求来源区域，以ISO-3166-1 两位国家代码表示。例如，'CN'代表'中国'。"}
  ClientRegions []*string `json:"clientRegions,omitempty" xml:"clientRegions,omitempty" type:"Repeated"`
  // {"en":"Indicates the starting and ending IP addresses of the client requests to match against. They must be in IPv4 or IPv6 format.","zh_CN":"用于指定要匹配的客户端请求的开始和结束IP地址。必须是IPv4或IPv6格式。"}
  ClientIpRange []*string `json:"clientIpRange,omitempty" xml:"clientIpRange,omitempty" type:"Repeated"`
}

func (s CreateAPropertyRequestVersionAccessControlRulesConditions) String() string {
  return tea.Prettify(s)
}

func (s CreateAPropertyRequestVersionAccessControlRulesConditions) GoString() string {
  return s.String()
}

func (s *CreateAPropertyRequestVersionAccessControlRulesConditions) SetScheme(v string) *CreateAPropertyRequestVersionAccessControlRulesConditions {
  s.Scheme = &v
  return s
}

func (s *CreateAPropertyRequestVersionAccessControlRulesConditions) SetHostname(v string) *CreateAPropertyRequestVersionAccessControlRulesConditions {
  s.Hostname = &v
  return s
}

func (s *CreateAPropertyRequestVersionAccessControlRulesConditions) SetUri(v string) *CreateAPropertyRequestVersionAccessControlRulesConditions {
  s.Uri = &v
  return s
}

func (s *CreateAPropertyRequestVersionAccessControlRulesConditions) SetServerRegions(v []*string) *CreateAPropertyRequestVersionAccessControlRulesConditions {
  s.ServerRegions = v
  return s
}

func (s *CreateAPropertyRequestVersionAccessControlRulesConditions) SetClientRegions(v []*string) *CreateAPropertyRequestVersionAccessControlRulesConditions {
  s.ClientRegions = v
  return s
}

func (s *CreateAPropertyRequestVersionAccessControlRulesConditions) SetClientIpRange(v []*string) *CreateAPropertyRequestVersionAccessControlRulesConditions {
  s.ClientIpRange = v
  return s
}

type CreateAPropertyRequestVersionAccessControlRulesAction struct {
  // {"en":"Indicates the HTTP status code to respond with. It must be in the range 300-309, 400-409, or 500-509 to indicate a redirection or error.","zh_CN":"响应的HTTP状态码，范围必须在300-309、400-409或500-509之间，分别表示重定向或错误。"}
  StatusCode *string `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
  // {"en":"Range: <= 200 characters\nIf the value of the status field is in the range 300-309, the message field's value will be placed in a Location HTTP header returned to the client. If the status is in the range 400-409 or 500-509, the value will be returned in the response body to the client.","zh_CN":"取值范围: <= 200 字符\n如果status字段的值在300-309范围内，message字段的值将在返回给客户端的Location响应头中。如果status字段的值在400-409或500-509范围内，则message字段的值将在返回给客户端的响应体中。"}
  Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s CreateAPropertyRequestVersionAccessControlRulesAction) String() string {
  return tea.Prettify(s)
}

func (s CreateAPropertyRequestVersionAccessControlRulesAction) GoString() string {
  return s.String()
}

func (s *CreateAPropertyRequestVersionAccessControlRulesAction) SetStatusCode(v string) *CreateAPropertyRequestVersionAccessControlRulesAction {
  s.StatusCode = &v
  return s
}

func (s *CreateAPropertyRequestVersionAccessControlRulesAction) SetMessage(v string) *CreateAPropertyRequestVersionAccessControlRulesAction {
  s.Message = &v
  return s
}

type CreateAPropertyRequestVersionRealTimeLog struct {
  // {"en":"The URL that receives the notifications. It must begin with 'http' or 'https'. The server should support the POST method. This is a required field.","zh_CN":"接收通知的服务器URL地址。必须以'http'或'https'开头。服务器须支持POST方法。这是必填字段。"}
  LogUrl *string `json:"logUrl,omitempty" xml:"logUrl,omitempty" require:"true"`
  // {"en":"Default: 1 Range: [ 1 .. 65536 ]\nAn integer between [1, 65536]. n means one notification for every n edge requests. 1 means every edge request will generate a notification. This is optional. Default is 1.","zh_CN":"默认值: 1 取值范围: [ 1 .. 65536 ]\n采样率。介于[1, 65536]之间的整数。n表示每n个边缘请求发送1条通知。1表示每个边缘请求都会发送。这是可选字段。默认值为1。"}
  SampleRate *int `json:"sampleRate,omitempty" xml:"sampleRate,omitempty"`
  // {"en":"Enum: json,none,compactJson\nSpecify json to enable JSON character escaping in variables, compactJson to enable JSON character escaping with whitespace in the body removed, or none to disable escaping.","zh_CN":"取值范围: json,none,compactJson\n指定json以开启变量中的json字符转义；指定compactJson以开启变量中的json字符转义，并移除变量中的空白行。如果要禁用转义则指定none。"}
  Escape *string `json:"escape,omitempty" xml:"escape,omitempty"`
  // {"en":"This is a required field to specify the notification body to be sent to the remote server. It can be any printable text that can be sent in the body of an HTTP POST request.\n\nThe CDN Pro built-in variables can be used within the format field. They will be replaced with the actual values in the notifications.","zh_CN":"该字段用来定义要发送到远程服务器的通知的格式。通知正文可以包括任何能在HTTP POST请求体中发送的可打印文本。这是必填字段。\n\n可以在格式定义中使用CDN Pro支持的内置变量。发送通知时，变量将被替换为实际值。"}
  Format *string `json:"format,omitempty" xml:"format,omitempty" require:"true"`
  // {"en":"HTTP header names and values to be sent to the notification server. A header name can contain any alphanumeric character or hyphen, '-'. A header value can contain any printable characters. It can also include any of the built-in variables supported in the format field of the realTimeLog object.","zh_CN":"需要发送到远程服务器的HTTP请求头名称和值。请求头名称可以包含任何字母，数字或连字符'-'。值可以包含任何可打印字符，也可以使用realTimeLog对象format字段中支持的任何内置变量。"}
  Headers []*CreateAPropertyRequestVersionRealTimeLogHeaders `json:"headers,omitempty" xml:"headers,omitempty" type:"Repeated"`
}

func (s CreateAPropertyRequestVersionRealTimeLog) String() string {
  return tea.Prettify(s)
}

func (s CreateAPropertyRequestVersionRealTimeLog) GoString() string {
  return s.String()
}

func (s *CreateAPropertyRequestVersionRealTimeLog) SetLogUrl(v string) *CreateAPropertyRequestVersionRealTimeLog {
  s.LogUrl = &v
  return s
}

func (s *CreateAPropertyRequestVersionRealTimeLog) SetSampleRate(v int) *CreateAPropertyRequestVersionRealTimeLog {
  s.SampleRate = &v
  return s
}

func (s *CreateAPropertyRequestVersionRealTimeLog) SetEscape(v string) *CreateAPropertyRequestVersionRealTimeLog {
  s.Escape = &v
  return s
}

func (s *CreateAPropertyRequestVersionRealTimeLog) SetFormat(v string) *CreateAPropertyRequestVersionRealTimeLog {
  s.Format = &v
  return s
}

func (s *CreateAPropertyRequestVersionRealTimeLog) SetHeaders(v []*CreateAPropertyRequestVersionRealTimeLogHeaders) *CreateAPropertyRequestVersionRealTimeLog {
  s.Headers = v
  return s
}

type CreateAPropertyRequestVersionRealTimeLogHeaders struct     {
  // {"en":"Name of an HTTP header.","zh_CN":"HTTP标头名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // {"en":"Value of an HTTP header.","zh_CN":"HTTP标头值"}
  Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s CreateAPropertyRequestVersionRealTimeLogHeaders) String() string {
  return tea.Prettify(s)
}

func (s CreateAPropertyRequestVersionRealTimeLogHeaders) GoString() string {
  return s.String()
}

func (s *CreateAPropertyRequestVersionRealTimeLogHeaders) SetName(v string) *CreateAPropertyRequestVersionRealTimeLogHeaders {
  s.Name = &v
  return s
}

func (s *CreateAPropertyRequestVersionRealTimeLogHeaders) SetValue(v string) *CreateAPropertyRequestVersionRealTimeLogHeaders {
  s.Value = &v
  return s
}

type CreateAPropertyRequestVersionVideoSeek struct {
  // {"en":"Range: [ 1 .. 31 ] characters\nName of the query parameter indicating the starting offset in bytes of the content to fetch. The parameter name should begin with a letter (a-z, A-Z) and may be followed by up to 30 letters and numbers.","zh_CN":"取值范围: [ 1 .. 31 ] 字符\n查询参数的名称，用来指定要获取的内容的起始位置（以字节计算）。参数名称应以字母（a-z，A-Z）开头，后面最多可以有30个字母和数字。"}
  StartParameter *string `json:"startParameter,omitempty" xml:"startParameter,omitempty" require:"true"`
  // {"en":"Range: [ 0 .. 31 ] characters\nName of the query parameter indicating the ending offset of the content to fetch. The parameter name should begin with a letter (a-z, A-Z) and may be followed by up to 30 letters and numbers.","zh_CN":"取值范围: [ 0 .. 31 ] 字符\n查询参数的名称，用来指定要获取的内容的结束位置（以字节计算）。参数名称应以字母（a-z，A-Z）开头，后面最多可以有30个字母和数字。"}
  EndParameter *string `json:"endParameter,omitempty" xml:"endParameter,omitempty"`
}

func (s CreateAPropertyRequestVersionVideoSeek) String() string {
  return tea.Prettify(s)
}

func (s CreateAPropertyRequestVersionVideoSeek) GoString() string {
  return s.String()
}

func (s *CreateAPropertyRequestVersionVideoSeek) SetStartParameter(v string) *CreateAPropertyRequestVersionVideoSeek {
  s.StartParameter = &v
  return s
}

func (s *CreateAPropertyRequestVersionVideoSeek) SetEndParameter(v string) *CreateAPropertyRequestVersionVideoSeek {
  s.EndParameter = &v
  return s
}

type CreateAPropertyRequestHeader struct {
}

func (s CreateAPropertyRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s CreateAPropertyRequestHeader) GoString() string {
  return s.String()
}

type CreateAPropertyPaths struct {
}

func (s CreateAPropertyPaths) String() string {
  return tea.Prettify(s)
}

func (s CreateAPropertyPaths) GoString() string {
  return s.String()
}

type CreateAPropertyParameters struct {
}

func (s CreateAPropertyParameters) String() string {
  return tea.Prettify(s)
}

func (s CreateAPropertyParameters) GoString() string {
  return s.String()
}

type CreateAPropertyResponse struct {
}

func (s CreateAPropertyResponse) String() string {
  return tea.Prettify(s)
}

func (s CreateAPropertyResponse) GoString() string {
  return s.String()
}

type CreateAPropertyResponseHeader struct {
  // {"en":"Returns a URL pointing to the new property created, if the request is accepted. The URL contains the ID of the new property. </br> URL format: <code>{scheme}://{domain}/cdn/properties/{propertyId}</code> Example URL: <code>https://api.example.com/cdn/properties/5dwa2205f9e9cc0001df7b24</code>","zh_CN":"当接口调用成功时，通过Location响应头返回新建的加速项目的URL。URL中包含加速项目ID。可使用该ID调用'查询加速项目的基础信息及版本信息'接口来查看加速项目详情。</br> URL格式：<code>{协议}://{域名}/cdn/properties/{加速项目ID}</code> URL示例：<code>https://open.chinanetcenter.com/cdn/properties/5dwa2205f9e9cc0001df7b24</code>"}
  Location *string `json:"Location,omitempty" xml:"Location,omitempty" require:"true"`
}

func (s CreateAPropertyResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s CreateAPropertyResponseHeader) GoString() string {
  return s.String()
}

func (s *CreateAPropertyResponseHeader) SetLocation(v string) *CreateAPropertyResponseHeader {
  s.Location = &v
  return s
}




