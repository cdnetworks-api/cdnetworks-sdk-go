package secretmanagement

// This file is auto-generated, don't edit it. Thanks.
import (
  "github.com/alibabacloud-go/tea/tea"
)

type UpdateASecretPaths struct {
  // {"en" : "ID of a secret.", "zh_CN": "保密信息id。"}
  SecretId *string `json:"secretId,omitempty" xml:"secretId,omitempty" require:"true"`
}

func (s UpdateASecretPaths) String() string {
  return tea.Prettify(s)
}

func (s UpdateASecretPaths) GoString() string {
  return s.String()
}

func (s *UpdateASecretPaths) SetSecretId(v string) *UpdateASecretPaths {
  s.SecretId = &v
  return s
}

type UpdateASecretParameters struct {
}

func (s UpdateASecretParameters) String() string {
  return tea.Prettify(s)
}

func (s UpdateASecretParameters) GoString() string {
  return s.String()
}

type UpdateASecretRequestHeader struct {
}

func (s UpdateASecretRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s UpdateASecretRequestHeader) GoString() string {
  return s.String()
}

type UpdateASecretRequest struct {
  // {"en" : "Range: [ 1 .. 30 ] characters ^[A-Za-z0-9_.-]+ 
  // Name of the secret. It can consist of letters, numbers, hyphens, underscores, and periods. Refer to this in your Edge Logic using the $SECRET(secretName) syntax. If you change the name, property versions referring to the old name will not pass future validation.", "zh_CN": "取值范围: [ 1 .. 30 ] 字符 ^[A-Za-z0-9_.-]+ 
  // 保密信息的名称。可以由字母、数字、连字符、下划线和句点组成。在边缘逻辑中使用$SECRET(secretName) 语法引用。如果您修改了保密信息的名称，那么以原有名称引用保密信息的加速项目版本将无法通过验证。"}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // {"en" : "Range: <= 250 characters 
  // A description of the secret. This may be useful for recording the purpose of the secret.", "zh_CN": "取值范围: <= 250 字符 
  // 保密信息的描述。"}
  Description *string `json:"description,omitempty" xml:"description,omitempty"`
  // {"en" : "Range: [ 8 .. 3599 ] characters 
  // The sensitive text you want to protect. The secret can consist of printable ASCII characters along with the tab, newline, and return characters. When you deploy your property, references to $SECRET(secretName) in your Edge Logic will be replaced by this text. Once the property is deployed, it will use the value of the secret at that time. If you change the secret's value and want your property to use the new value, you must revalidate and redeploy the property.", "zh_CN": "取值范围: [ 8 .. 3599 ] 字符 
  // 保密信息存放的敏感内容。可以包含可打印的ASCII字符、制表符(0x09)、换行符(0x0a) 和回车符(0x0d)。保密信息更新后，在已部署加速项目中已引用的保密信息不会自动被更新。如果您修改了保密信息存放的内容，必须重新验证并部署，加速项目才能使用到最新的值。"}
  Secret *string `json:"secret,omitempty" xml:"secret,omitempty"`
}

func (s UpdateASecretRequest) String() string {
  return tea.Prettify(s)
}

func (s UpdateASecretRequest) GoString() string {
  return s.String()
}

func (s *UpdateASecretRequest) SetName(v string) *UpdateASecretRequest {
  s.Name = &v
  return s
}

func (s *UpdateASecretRequest) SetDescription(v string) *UpdateASecretRequest {
  s.Description = &v
  return s
}

func (s *UpdateASecretRequest) SetSecret(v string) *UpdateASecretRequest {
  s.Secret = &v
  return s
}

type UpdateASecretResponseHeader struct {
}

func (s UpdateASecretResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s UpdateASecretResponseHeader) GoString() string {
  return s.String()
}

type UpdateASecretResponse struct {
}

func (s UpdateASecretResponse) String() string {
  return tea.Prettify(s)
}

func (s UpdateASecretResponse) GoString() string {
  return s.String()
}




type GetASecretPaths struct {
  // {"en" : "ID of a secret.", "zh_CN": "保密信息id。"}
  SecretId *string `json:"secretId,omitempty" xml:"secretId,omitempty" require:"true"`
}

func (s GetASecretPaths) String() string {
  return tea.Prettify(s)
}

func (s GetASecretPaths) GoString() string {
  return s.String()
}

func (s *GetASecretPaths) SetSecretId(v string) *GetASecretPaths {
  s.SecretId = &v
  return s
}

type GetASecretParameters struct {
}

func (s GetASecretParameters) String() string {
  return tea.Prettify(s)
}

func (s GetASecretParameters) GoString() string {
  return s.String()
}

type GetASecretRequestHeader struct {
}

func (s GetASecretRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetASecretRequestHeader) GoString() string {
  return s.String()
}

type GetASecretRequest struct {
}

func (s GetASecretRequest) String() string {
  return tea.Prettify(s)
}

func (s GetASecretRequest) GoString() string {
  return s.String()
}

type GetASecretResponseHeader struct {
}

func (s GetASecretResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetASecretResponseHeader) GoString() string {
  return s.String()
}

type GetASecretResponse struct {
  // {"en" : "ID of the secret.", "zh_CN": "保密信息ID。"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {"en" : "Range: [ 1 .. 30 ] characters ^[A-Za-z0-9_.-]+ 
  // Name of the secret. It can consist of letters, numbers, hyphens, underscores, and periods. Refer to this in your Edge Logic using the $SECRET(secretName) syntax.", "zh_CN": "取值范围: [ 1 .. 30 ] 字符 ^[A-Za-z0-9_.-]+ 
  // 保密信息的名称。可以由字母、数字、连字符、下划线和句点组成。在边缘逻辑中使用$SECRET(secretName) 语法引用。"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en" : "Range: <= 250 characters 
  // Description of the secret.", "zh_CN": "取值范围: <= 250 字符 
  // 保密信息的描述。"}
  Description *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
  // {"en" : "Range: [ 8 .. 3599 ] characters 
  // The value of the secret. The secret can consist of printable ASCII characters along with the tab, newline, and return characters. For security, we will not show all the characters. The middle of the secret will be masked out by asterisk characters. References to $SECRET(secretName) in your Edge Logic will be replaced by the value.", "zh_CN": "取值范围: [ 8 .. 3599 ] 字符 
  // 保密信息存放的敏感内容。可以包含可打印的ASCII字符、制表符(0x09)、换行符(0x0a) 和回车符(0x0d)。出于安全考虑，此处不显示所有字符，部分内容将用星号标记。"}
  Secret *string `json:"secret,omitempty" xml:"secret,omitempty" require:"true"`
  // {"en" : "An RFC 3339 date indicating when the secret was last updated, for example, '2021-07-06T00:00:00Z'.", "zh_CN": "RFC 3339格式的日期，表示保密信息的最后一次更新时间。例如，'2021-07-06T00:00:00Z'。"}
  LastUpdateTime *string `json:"lastUpdateTime,omitempty" xml:"lastUpdateTime,omitempty" require:"true"`
  // {"en" : "An RFC 3339 date indicating when the secret was created. This will be in UTC time, for example, '2021-07-06T00:00:00Z'.", "zh_CN": "RFC 3339格式的日期，表示保密信息的创建时间，使用UTC时区。例如，'2021-07-06T00:00:00Z'。"}
  CreationTime *string `json:"creationTime,omitempty" xml:"creationTime,omitempty" require:"true"`
  // {"en" : "Properties deployed to production that refer to the secret.", "zh_CN": "保密信息在生产环境中的使用情况。"}
  UsageInProduction []*GetASecretResponseUsageInProduction `json:"usageInProduction,omitempty" xml:"usageInProduction,omitempty" require:"true" type:"Repeated"`
  // {"en" : "Properties deployed to staging that refer to the secret.", "zh_CN": "保密信息在演练环境中的使用情况。"}
  UsageInStaging []*GetASecretResponseUsageInStaging `json:"usageInStaging,omitempty" xml:"usageInStaging,omitempty" require:"true" type:"Repeated"`
}

func (s GetASecretResponse) String() string {
  return tea.Prettify(s)
}

func (s GetASecretResponse) GoString() string {
  return s.String()
}

func (s *GetASecretResponse) SetId(v string) *GetASecretResponse {
  s.Id = &v
  return s
}

func (s *GetASecretResponse) SetName(v string) *GetASecretResponse {
  s.Name = &v
  return s
}

func (s *GetASecretResponse) SetDescription(v string) *GetASecretResponse {
  s.Description = &v
  return s
}

func (s *GetASecretResponse) SetSecret(v string) *GetASecretResponse {
  s.Secret = &v
  return s
}

func (s *GetASecretResponse) SetLastUpdateTime(v string) *GetASecretResponse {
  s.LastUpdateTime = &v
  return s
}

func (s *GetASecretResponse) SetCreationTime(v string) *GetASecretResponse {
  s.CreationTime = &v
  return s
}

func (s *GetASecretResponse) SetUsageInProduction(v []*GetASecretResponseUsageInProduction) *GetASecretResponse {
  s.UsageInProduction = v
  return s
}

func (s *GetASecretResponse) SetUsageInStaging(v []*GetASecretResponseUsageInStaging) *GetASecretResponse {
  s.UsageInStaging = v
  return s
}

type GetASecretResponseUsageInProduction struct     {
  // {"en" : "ID of a property using the secret.", "zh_CN": "引用此保密信息的加速项目ID。"}
  PropertyId *string `json:"propertyId,omitempty" xml:"propertyId,omitempty"`
}

func (s GetASecretResponseUsageInProduction) String() string {
  return tea.Prettify(s)
}

func (s GetASecretResponseUsageInProduction) GoString() string {
  return s.String()
}

func (s *GetASecretResponseUsageInProduction) SetPropertyId(v string) *GetASecretResponseUsageInProduction {
  s.PropertyId = &v
  return s
}

type GetASecretResponseUsageInStaging struct     {
  // {"en" : "ID of a property using the secret.", "zh_CN": "引用此保密信息的加速项目ID。"}
  PropertyId *string `json:"propertyId,omitempty" xml:"propertyId,omitempty"`
}

func (s GetASecretResponseUsageInStaging) String() string {
  return tea.Prettify(s)
}

func (s GetASecretResponseUsageInStaging) GoString() string {
  return s.String()
}

func (s *GetASecretResponseUsageInStaging) SetPropertyId(v string) *GetASecretResponseUsageInStaging {
  s.PropertyId = &v
  return s
}




type DeleteASecretPaths struct {
  // {"en" : "ID of a secret.", "zh_CN": "保密信息id。"}
  SecretId *string `json:"secretId,omitempty" xml:"secretId,omitempty" require:"true"`
}

func (s DeleteASecretPaths) String() string {
  return tea.Prettify(s)
}

func (s DeleteASecretPaths) GoString() string {
  return s.String()
}

func (s *DeleteASecretPaths) SetSecretId(v string) *DeleteASecretPaths {
  s.SecretId = &v
  return s
}

type DeleteASecretParameters struct {
}

func (s DeleteASecretParameters) String() string {
  return tea.Prettify(s)
}

func (s DeleteASecretParameters) GoString() string {
  return s.String()
}

type DeleteASecretRequestHeader struct {
}

func (s DeleteASecretRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteASecretRequestHeader) GoString() string {
  return s.String()
}

type DeleteASecretRequest struct {
}

func (s DeleteASecretRequest) String() string {
  return tea.Prettify(s)
}

func (s DeleteASecretRequest) GoString() string {
  return s.String()
}

type DeleteASecretResponseHeader struct {
}

func (s DeleteASecretResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteASecretResponseHeader) GoString() string {
  return s.String()
}

type DeleteASecretResponse struct {
}

func (s DeleteASecretResponse) String() string {
  return tea.Prettify(s)
}

func (s DeleteASecretResponse) GoString() string {
  return s.String()
}




type GetAListOfSecretsPaths struct {
}

func (s GetAListOfSecretsPaths) String() string {
  return tea.Prettify(s)
}

func (s GetAListOfSecretsPaths) GoString() string {
  return s.String()
}

type GetAListOfSecretsParameters struct {
  // {"en" : "Default: 0 Range: >= 0 
  // Index of the first secret to return.", "zh_CN": "默认值: 0 取值范围: >= 0 
  // 查询起始位置。"}
  Offset *int `json:"offset,omitempty" xml:"offset,omitempty"`
  // {"en" : "Default: 100 Range: [ 1 .. 200 ] 
  // Maximum number of secrets to return.", "zh_CN": "默认值: 100 取值范围: <= 200 
  // 每次查询的最大条数。"}
  Limit *int `json:"limit,omitempty" xml:"limit,omitempty"`
  // {"en" : "The value of the search parameter is matched on the id, name, and description fields of the secrets. Partial matches are supported.", "zh_CN": "根据搜索关键字匹配保密信息的ID，名称和描述进行查询。"}
  Search *string `json:"search,omitempty" xml:"search,omitempty"`
}

func (s GetAListOfSecretsParameters) String() string {
  return tea.Prettify(s)
}

func (s GetAListOfSecretsParameters) GoString() string {
  return s.String()
}

func (s *GetAListOfSecretsParameters) SetOffset(v int) *GetAListOfSecretsParameters {
  s.Offset = &v
  return s
}

func (s *GetAListOfSecretsParameters) SetLimit(v int) *GetAListOfSecretsParameters {
  s.Limit = &v
  return s
}

func (s *GetAListOfSecretsParameters) SetSearch(v string) *GetAListOfSecretsParameters {
  s.Search = &v
  return s
}

type GetAListOfSecretsRequestHeader struct {
}

func (s GetAListOfSecretsRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetAListOfSecretsRequestHeader) GoString() string {
  return s.String()
}

type GetAListOfSecretsRequest struct {
}

func (s GetAListOfSecretsRequest) String() string {
  return tea.Prettify(s)
}

func (s GetAListOfSecretsRequest) GoString() string {
  return s.String()
}

type GetAListOfSecretsResponseHeader struct {
}

func (s GetAListOfSecretsResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetAListOfSecretsResponseHeader) GoString() string {
  return s.String()
}

type GetAListOfSecretsResponse struct {
  // {"en" : "Range: >= 0 
  // Total number of secrets. The actual number returned in the secrets field of the response will depend on the limit and offset parameters.", "zh_CN": "取值范围: >= 0 
  // 保密信息的总数。返回的实际数量取决于查询参数。"}
  Count *int `json:"count,omitempty" xml:"count,omitempty" require:"true"`
  // {"en" : "A list of secrets.", "zh_CN": "保密信息列表。"}
  Secrets []*GetAListOfSecretsResponseSecrets `json:"secrets,omitempty" xml:"secrets,omitempty" require:"true" type:"Repeated"`
}

func (s GetAListOfSecretsResponse) String() string {
  return tea.Prettify(s)
}

func (s GetAListOfSecretsResponse) GoString() string {
  return s.String()
}

func (s *GetAListOfSecretsResponse) SetCount(v int) *GetAListOfSecretsResponse {
  s.Count = &v
  return s
}

func (s *GetAListOfSecretsResponse) SetSecrets(v []*GetAListOfSecretsResponseSecrets) *GetAListOfSecretsResponse {
  s.Secrets = v
  return s
}

type GetAListOfSecretsResponseSecrets struct     {
  // {"en" : "ID of a secret.", "zh_CN": "保密信息ID。"}
  Id *string `json:"id,omitempty" xml:"id,omitempty"`
  // {"en" : "Range: [ 1 .. 30 ] characters ^[A-Za-z0-9_.-]+ 
  // Name of the secret. It can consist of letters, numbers, hyphens, underscores, and periods. Refer to this in your Edge Logic as $SECRET(secretName).", "zh_CN": "取值范围: [ 1 .. 30 ] 字符 ^[A-Za-z0-9_.-]+ 
  // 保密信息的名称。可以由字母、数字、连字符、下划线和句点组成。在边缘逻辑中使用$SECRET(secretName) 语法引用。"}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // {"en" : "Range: <= 250 characters 
  // A description of the secret.", "zh_CN": "取值范围: <= 250 字符 
  // 保密信息的描述。"}
  Description *string `json:"description,omitempty" xml:"description,omitempty"`
  // {"en" : "An RFC 3339 date indicating when the secret was created. This will be in UTC time, for example, '2021-07-06T00:00:00Z'.
  // ", "zh_CN": "RFC 3339格式的日期，表示保密信息的创建时间，使用UTC时区。例如，'2021-07-06T00:00:00Z'。"}
  CreationTime *string `json:"creationTime,omitempty" xml:"creationTime,omitempty"`
  // {"en" : "An RFC 3339 date indicating when the secret was last updated, for example, '2021-07-06T00:00:00Z'.
  // ", "zh_CN": "RFC 3339格式的日期，表示保密信息的最后一次更新时间。例如，'2021-07-06T00:00:00Z'。"}
  LastUpdateTime *string `json:"lastUpdateTime,omitempty" xml:"lastUpdateTime,omitempty"`
}

func (s GetAListOfSecretsResponseSecrets) String() string {
  return tea.Prettify(s)
}

func (s GetAListOfSecretsResponseSecrets) GoString() string {
  return s.String()
}

func (s *GetAListOfSecretsResponseSecrets) SetId(v string) *GetAListOfSecretsResponseSecrets {
  s.Id = &v
  return s
}

func (s *GetAListOfSecretsResponseSecrets) SetName(v string) *GetAListOfSecretsResponseSecrets {
  s.Name = &v
  return s
}

func (s *GetAListOfSecretsResponseSecrets) SetDescription(v string) *GetAListOfSecretsResponseSecrets {
  s.Description = &v
  return s
}

func (s *GetAListOfSecretsResponseSecrets) SetCreationTime(v string) *GetAListOfSecretsResponseSecrets {
  s.CreationTime = &v
  return s
}

func (s *GetAListOfSecretsResponseSecrets) SetLastUpdateTime(v string) *GetAListOfSecretsResponseSecrets {
  s.LastUpdateTime = &v
  return s
}




type CreateASecretPaths struct {
}

func (s CreateASecretPaths) String() string {
  return tea.Prettify(s)
}

func (s CreateASecretPaths) GoString() string {
  return s.String()
}

type CreateASecretParameters struct {
}

func (s CreateASecretParameters) String() string {
  return tea.Prettify(s)
}

func (s CreateASecretParameters) GoString() string {
  return s.String()
}

type CreateASecretRequestHeader struct {
}

func (s CreateASecretRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s CreateASecretRequestHeader) GoString() string {
  return s.String()
}

type CreateASecretRequest struct {
  // {"en" : "Range: [ 1 .. 30 ] characters ^[A-Za-z0-9_.-]+ 
  // The name of the secret. It can consist of letters, numbers, hyphens, underscores, and periods. Refer to this in your Edge Logic using the $SECRET(secretName) syntax.", "zh_CN": "取值范围: [ 1 .. 30 ] 字符 ^[A-Za-z0-9_.-]+ 
  // 保密信息的名称。可以由字母、数字、连字符、下划线和句点组成。在边缘逻辑中使用$SECRET(secretName) 语法来引用。"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en" : "Range: <= 250 characters 
  // A description of the secret. This may be useful for recording the purpose of the secret.", "zh_CN": "取值范围: <= 250 字符 
  // 保密信息的描述。"}
  Description *string `json:"description,omitempty" xml:"description,omitempty"`
  // {"en" : "Range: [ 8 .. 3599 ] characters 
  // The sensitive text you want to protect. When you deploy your property, references to $SECRET(secretName) in your Edge Logic will be replaced by this text. The secret can consist of printable ASCII characters along with the tab, newline, and return characters.", "zh_CN": "取值范围: [ 8 .. 3599 ] 字符 
  // 需要被保护的敏感内容。当部署加速项目时，将对边缘逻辑中的 $SECRET(secretName)进行解析，提取出敏感内容。支持可打印的ASCII字符、制表符(0x09)、换行符(0x0a) 和回车符(0x0d)。"}
  Secret *string `json:"secret,omitempty" xml:"secret,omitempty" require:"true"`
}

func (s CreateASecretRequest) String() string {
  return tea.Prettify(s)
}

func (s CreateASecretRequest) GoString() string {
  return s.String()
}

func (s *CreateASecretRequest) SetName(v string) *CreateASecretRequest {
  s.Name = &v
  return s
}

func (s *CreateASecretRequest) SetDescription(v string) *CreateASecretRequest {
  s.Description = &v
  return s
}

func (s *CreateASecretRequest) SetSecret(v string) *CreateASecretRequest {
  s.Secret = &v
  return s
}

type CreateASecretResponse struct {
}

func (s CreateASecretResponse) String() string {
  return tea.Prettify(s)
}

func (s CreateASecretResponse) GoString() string {
  return s.String()
}

type CreateASecretResponseHeader struct {
  // {"en":"The location header contains a reference to the new secret's ID.", "zh_CN":"通过Location响应头返回新建的保密信息的URL。URL中包含保密信息的ID，可使用该ID调用‘获取保密信息详请’接口来查看保密信息的详请。URL示例：<code>Location: https://openapi.chinanetcenter.com/cdn/secrets/60d6707cca3e387d2a28fc9e</code>。"}
  Location *string `json:"Location,omitempty" xml:"Location,omitempty" require:"true"`
}

func (s CreateASecretResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s CreateASecretResponseHeader) GoString() string {
  return s.String()
}

func (s *CreateASecretResponseHeader) SetLocation(v string) *CreateASecretResponseHeader {
  s.Location = &v
  return s
}




