package policy

// This file is auto-generated, don't edit it. Thanks.
import (
  "github.com/alibabacloud-go/tea/tea"
)

type DeletePolicyRequest struct {
  // {"en":"Policy ID. Only supports deleting custom permission policies.", "zh_CN":"策略ID。仅支持删除自定义权限策略"}
  PolicyId *string `json:"policyId,omitempty" xml:"policyId,omitempty" require:"true"`
}

func (s DeletePolicyRequest) String() string {
  return tea.Prettify(s)
}

func (s DeletePolicyRequest) GoString() string {
  return s.String()
}

func (s *DeletePolicyRequest) SetPolicyId(v string) *DeletePolicyRequest {
  s.PolicyId = &v
  return s
}

type DeletePolicyResponse struct {
  // {"en":"request result status code", "zh_CN":"请求结果状态码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Request result information", "zh_CN":"请求结果信息"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
}

func (s DeletePolicyResponse) String() string {
  return tea.Prettify(s)
}

func (s DeletePolicyResponse) GoString() string {
  return s.String()
}

func (s *DeletePolicyResponse) SetCode(v string) *DeletePolicyResponse {
  s.Code = &v
  return s
}

func (s *DeletePolicyResponse) SetMsg(v string) *DeletePolicyResponse {
  s.Msg = &v
  return s
}

type DeletePolicyPaths struct {
}

func (s DeletePolicyPaths) String() string {
  return tea.Prettify(s)
}

func (s DeletePolicyPaths) GoString() string {
  return s.String()
}

type DeletePolicyParameters struct {
}

func (s DeletePolicyParameters) String() string {
  return tea.Prettify(s)
}

func (s DeletePolicyParameters) GoString() string {
  return s.String()
}

type DeletePolicyRequestHeader struct {
}

func (s DeletePolicyRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s DeletePolicyRequestHeader) GoString() string {
  return s.String()
}

type DeletePolicyResponseHeader struct {
}

func (s DeletePolicyResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s DeletePolicyResponseHeader) GoString() string {
  return s.String()
}




type GetAccountSummaryRequest struct {
}

func (s GetAccountSummaryRequest) String() string {
  return tea.Prettify(s)
}

func (s GetAccountSummaryRequest) GoString() string {
  return s.String()
}

type GetAccountSummaryRequestHeader struct {
}

func (s GetAccountSummaryRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetAccountSummaryRequestHeader) GoString() string {
  return s.String()
}

type GetAccountSummaryPaths struct {
}

func (s GetAccountSummaryPaths) String() string {
  return tea.Prettify(s)
}

func (s GetAccountSummaryPaths) GoString() string {
  return s.String()
}

type GetAccountSummaryParameters struct {
}

func (s GetAccountSummaryParameters) String() string {
  return tea.Prettify(s)
}

func (s GetAccountSummaryParameters) GoString() string {
  return s.String()
}

type GetAccountSummaryResponse struct {
  // {"en":"code","zh_CN":"请求结果状态码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"message","zh_CN":"请求结果信息"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"","zh_CN":"请求结果的详细数据"}
  Data *GetAccountSummaryResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
}

func (s GetAccountSummaryResponse) String() string {
  return tea.Prettify(s)
}

func (s GetAccountSummaryResponse) GoString() string {
  return s.String()
}

func (s *GetAccountSummaryResponse) SetCode(v string) *GetAccountSummaryResponse {
  s.Code = &v
  return s
}

func (s *GetAccountSummaryResponse) SetMsg(v string) *GetAccountSummaryResponse {
  s.Msg = &v
  return s
}

func (s *GetAccountSummaryResponse) SetData(v *GetAccountSummaryResponseData) *GetAccountSummaryResponse {
  s.Data = v
  return s
}

type GetAccountSummaryResponseData struct {
  // {"en":"Maximum number of user groups to be created","zh_CN":"允许创建用户组最大数量"}
  GroupsQuota *int64 `json:"groupsQuota,omitempty" xml:"groupsQuota,omitempty" require:"true"`
  // {"en":"Number of user groups","zh_CN":"用户组数量"}
  Groups *int64 `json:"groups,omitempty" xml:"groups,omitempty" require:"true"`
  // {"en":"Maximum number of custom Policies allowed to be created","zh_CN":"允许创建自定义策略的最大数量"}
  PoliciesQuota *int64 `json:"policiesQuota,omitempty" xml:"policiesQuota,omitempty" require:"true"`
  // {"en":"Number of custom policies","zh_CN":"自定义策略数量"}
  Policies *int64 `json:"policies,omitempty" xml:"policies,omitempty" require:"true"`
  // {"en":"Maximum number of permission policy versions","zh_CN":"权限策略版本的最大数量"}
  VersionPolicyQuota *int64 `json:"versionPolicyQuota,omitempty" xml:"versionPolicyQuota,omitempty" require:"true"`
  // {"en":"Maximum number of sub-users allowed to create","zh_CN":"允许创建的子用户最大数量"}
  UsersQuota *int64 `json:"usersQuota,omitempty" xml:"usersQuota,omitempty" require:"true"`
  // {"en":"Number of sub-users","zh_CN":"子用户数量"}
  Users *int64 `json:"users,omitempty" xml:"users,omitempty" require:"true"`
}

func (s GetAccountSummaryResponseData) String() string {
  return tea.Prettify(s)
}

func (s GetAccountSummaryResponseData) GoString() string {
  return s.String()
}

func (s *GetAccountSummaryResponseData) SetGroupsQuota(v int64) *GetAccountSummaryResponseData {
  s.GroupsQuota = &v
  return s
}

func (s *GetAccountSummaryResponseData) SetGroups(v int64) *GetAccountSummaryResponseData {
  s.Groups = &v
  return s
}

func (s *GetAccountSummaryResponseData) SetPoliciesQuota(v int64) *GetAccountSummaryResponseData {
  s.PoliciesQuota = &v
  return s
}

func (s *GetAccountSummaryResponseData) SetPolicies(v int64) *GetAccountSummaryResponseData {
  s.Policies = &v
  return s
}

func (s *GetAccountSummaryResponseData) SetVersionPolicyQuota(v int64) *GetAccountSummaryResponseData {
  s.VersionPolicyQuota = &v
  return s
}

func (s *GetAccountSummaryResponseData) SetUsersQuota(v int64) *GetAccountSummaryResponseData {
  s.UsersQuota = &v
  return s
}

func (s *GetAccountSummaryResponseData) SetUsers(v int64) *GetAccountSummaryResponseData {
  s.Users = &v
  return s
}

type GetAccountSummaryResponseHeader struct {
}

func (s GetAccountSummaryResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetAccountSummaryResponseHeader) GoString() string {
  return s.String()
}




type EditPolicyRequest struct {
  // {"en":"The ID of the policy to be modified.","zh_CN":"要修改的策略id"}
  PolicyId *string `json:"policyId,omitempty" xml:"policyId,omitempty" require:"true"`
  // {"en":"The name of the policy to be modified..Not passing means no modification.","zh_CN":"要修改的策略名称。不传代表不修改。"}
  PolicyName *string `json:"policyName,omitempty" xml:"policyName,omitempty"`
  // {"en":"Policy description. You may describe the strategy here, including permission details, limited to a maximum of 250 characters.Not passing means no modification.","zh_CN":"修改后策略描述。可在此用文字简单描述策略包含权限内容，限制最多250字符。不传代表不修改。"}
  Description *string `json:"description,omitempty" xml:"description,omitempty"`
  // {"en":"Permission policy content, not passing means no modification. please fill in the permission policy script as follows:\n\n\n[{\n\"effect\": \"allow\",\n\"action\": [\n\"productCode:actionCode\"\n],\n\"resource\": [\n\"*\"\n]\n}]\n\n\nA single permission policy can include permissions for multiple products, but CDN and non-CDN product permissions cannot be added to the same policy simultaneously.\n\nField descriptions:\n\n-effect: The authorization effect includes two types: allow and deny.\n\n-action: Describes the specific operations allowed or denied,  format: productCode:actionCode.\n\n-resource: The specific resources authorized. For all resources use, for specific resources refer to the format: wsc:&lt;service-name&gt;:&lt;region&gt;:&lt;account&gt;:&lt;relatice-id&gt;. Note: CDN products do not support specifying resources.","zh_CN":"修改后权限策略内容，不传代表不修改。请填写权限策略脚本，示例如下：\n\n[{\n\"effect\": \"allow\",\n\"action\": [\n\"productCode:actionCode\"\n],\n\"resource\": [\n\"*\"\n]\n}]\n\n一个权限策略内可以包含多个产品的权限策略，但CDN和非CDN类产品权限不能同时添加到同一个策略里。\n\n各字段说明：\n\n-effect：授权效果包括两种允许（allow）和拒绝（deny）\n\n-action：描述允许或拒绝的特定操作，格式：productCode:actionCode\n\n-resource：授权的具体资源对象。若是全部资源用*表示，若是指定资源参考格式：wsc:&lt;service-name&gt;:&lt;region&gt;:&lt;account&gt;:&lt;relatice-id&gt;。注意：CDN类产品不支持指定资源"}
  PolicyDocument *string `json:"policyDocument,omitempty" xml:"policyDocument,omitempty"`
}

func (s EditPolicyRequest) String() string {
  return tea.Prettify(s)
}

func (s EditPolicyRequest) GoString() string {
  return s.String()
}

func (s *EditPolicyRequest) SetPolicyId(v string) *EditPolicyRequest {
  s.PolicyId = &v
  return s
}

func (s *EditPolicyRequest) SetPolicyName(v string) *EditPolicyRequest {
  s.PolicyName = &v
  return s
}

func (s *EditPolicyRequest) SetDescription(v string) *EditPolicyRequest {
  s.Description = &v
  return s
}

func (s *EditPolicyRequest) SetPolicyDocument(v string) *EditPolicyRequest {
  s.PolicyDocument = &v
  return s
}

type EditPolicyRequestHeader struct {
}

func (s EditPolicyRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s EditPolicyRequestHeader) GoString() string {
  return s.String()
}

type EditPolicyPaths struct {
}

func (s EditPolicyPaths) String() string {
  return tea.Prettify(s)
}

func (s EditPolicyPaths) GoString() string {
  return s.String()
}

type EditPolicyParameters struct {
}

func (s EditPolicyParameters) String() string {
  return tea.Prettify(s)
}

func (s EditPolicyParameters) GoString() string {
  return s.String()
}

type EditPolicyResponse struct {
  // {"en":"request result status code","zh_CN":"请求结果状态码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Request result information","zh_CN":"请求结果信息"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
}

func (s EditPolicyResponse) String() string {
  return tea.Prettify(s)
}

func (s EditPolicyResponse) GoString() string {
  return s.String()
}

func (s *EditPolicyResponse) SetCode(v string) *EditPolicyResponse {
  s.Code = &v
  return s
}

func (s *EditPolicyResponse) SetMsg(v string) *EditPolicyResponse {
  s.Msg = &v
  return s
}

type EditPolicyResponseHeader struct {
}

func (s EditPolicyResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s EditPolicyResponseHeader) GoString() string {
  return s.String()
}




type GetPolicyRequest struct {
  // {"en":"Policy ID","zh_CN":"策略ID"}
  PolicyId *int64 `json:"policyId,omitempty" xml:"policyId,omitempty"`
  // {"en":"Policy name.Either Policy ID or Policy Name must be provided. If Policy ID is entered, it will be prioritized.","zh_CN":"策略名称，策略ID和策略名称取任一必填，若填写了策略ID将优先使用策略ID"}
  PolicyName *string `json:"policyName,omitempty" xml:"policyName,omitempty"`
}

func (s GetPolicyRequest) String() string {
  return tea.Prettify(s)
}

func (s GetPolicyRequest) GoString() string {
  return s.String()
}

func (s *GetPolicyRequest) SetPolicyId(v int64) *GetPolicyRequest {
  s.PolicyId = &v
  return s
}

func (s *GetPolicyRequest) SetPolicyName(v string) *GetPolicyRequest {
  s.PolicyName = &v
  return s
}

type GetPolicyRequestHeader struct {
  // {"en":"Select the specified language and return the policy description of the corresponding language. Optional values:zh_CN，en，ko_KR，ja_JP；Default language is en. Custom policy descriptions are not language","zh_CN":"选择指定语言返回对应语言的策略描述，可选值：zh_CN，en，ko_KR，ja_JP；未选择默认en。  自定义策略描述不区分语言，语言参数不生效。"}
  AcceptLanguage *string `json:"Accept-Language,omitempty" xml:"Accept-Language,omitempty" require:"true"`
}

func (s GetPolicyRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetPolicyRequestHeader) GoString() string {
  return s.String()
}

func (s *GetPolicyRequestHeader) SetAcceptLanguage(v string) *GetPolicyRequestHeader {
  s.AcceptLanguage = &v
  return s
}

type GetPolicyPaths struct {
}

func (s GetPolicyPaths) String() string {
  return tea.Prettify(s)
}

func (s GetPolicyPaths) GoString() string {
  return s.String()
}

type GetPolicyParameters struct {
}

func (s GetPolicyParameters) String() string {
  return tea.Prettify(s)
}

func (s GetPolicyParameters) GoString() string {
  return s.String()
}

type GetPolicyResponse struct {
  // {"en":"request result status code","zh_CN":"请求结果状态码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Request result information","zh_CN":"请求结果信息"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"Detailed data on the results of the request","zh_CN":"请求结果的详细数据"}
  Data *GetPolicyResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
  // {"en":"Request ID","zh_CN":"请求ID"}
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty" require:"true"`
}

func (s GetPolicyResponse) String() string {
  return tea.Prettify(s)
}

func (s GetPolicyResponse) GoString() string {
  return s.String()
}

func (s *GetPolicyResponse) SetCode(v string) *GetPolicyResponse {
  s.Code = &v
  return s
}

func (s *GetPolicyResponse) SetMsg(v string) *GetPolicyResponse {
  s.Msg = &v
  return s
}

func (s *GetPolicyResponse) SetData(v *GetPolicyResponseData) *GetPolicyResponse {
  s.Data = v
  return s
}

func (s *GetPolicyResponse) SetRequestId(v string) *GetPolicyResponse {
  s.RequestId = &v
  return s
}

type GetPolicyResponseData struct {
  // {"en":"Policy ID","zh_CN":"策略ID"}
  PolicyId *int64 `json:"policyId,omitempty" xml:"policyId,omitempty" require:"true"`
  // {"en":"Policy name","zh_CN":"策略名称"}
  PolicyName *string `json:"policyName,omitempty" xml:"policyName,omitempty" require:"true"`
  // {"en":"Policy description","zh_CN":"策略描述"}
  Description *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
  // {"en":"Policy type: custom represents a custom Policy , system represents a system global Policy","zh_CN":"策略类型：custom代表自定义策略  system代表系统策略"}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
  // {"en":"Permission policy content, as follows:\n\n[{\n\"effect\": \"allow\",\n\"action\": [\n\"productCode:actionCode\"\n],\n\"resource\": [\n\"*\"\n]\n}]\n\nField descriptions:\n\n-effect: The authorization effect includes two types: allow and deny.\n\n-action: Describes the specific operations allowed or denied,  format: productCode:actionCode.\n\n-resource: The specific resources authorized. For all resources use, for specific resources refer to the format: wsc:&lt;service-name&gt;:&lt;region&gt;:&lt;account&gt;:&lt;relatice-id&gt;. Note: CDN products do not support specifying resources.","zh_CN":"策略权限描述，示例如下：\n\n[{\n\"effect\": \"allow\",\n\"action\": [\n\"productCode:actionCode\"\n],\n\"resource\": [\n\"*\"\n]\n}]\n\n各字段说明：\n\n-effect：授权效果包括两种允许（allow）和拒绝（deny）\n\n-action：描述允许或拒绝的特定操作，格式：productCode:actionCode\n\n-resource：授权的具体资源对象。若是全部资源用*表示，若是指定资源参考格式：wsc:&lt;service-name&gt;:&lt;region&gt;:&lt;account&gt;:&lt;relatice-id&gt;。注意：CDN类产品不支持指定资源"}
  PolicyDocument *string `json:"policyDocument,omitempty" xml:"policyDocument,omitempty" require:"true"`
}

func (s GetPolicyResponseData) String() string {
  return tea.Prettify(s)
}

func (s GetPolicyResponseData) GoString() string {
  return s.String()
}

func (s *GetPolicyResponseData) SetPolicyId(v int64) *GetPolicyResponseData {
  s.PolicyId = &v
  return s
}

func (s *GetPolicyResponseData) SetPolicyName(v string) *GetPolicyResponseData {
  s.PolicyName = &v
  return s
}

func (s *GetPolicyResponseData) SetDescription(v string) *GetPolicyResponseData {
  s.Description = &v
  return s
}

func (s *GetPolicyResponseData) SetType(v string) *GetPolicyResponseData {
  s.Type = &v
  return s
}

func (s *GetPolicyResponseData) SetPolicyDocument(v string) *GetPolicyResponseData {
  s.PolicyDocument = &v
  return s
}

type GetPolicyResponseHeader struct {
}

func (s GetPolicyResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetPolicyResponseHeader) GoString() string {
  return s.String()
}




type CreatePolicyRequest struct {
  // {"en":"Policy Name. Supports Chinese, English, and underline, with no more than 150 characters","zh_CN":"策略名称，支持中文、英文和下划线，不超过150字符"}
  PolicyName *string `json:"policyName,omitempty" xml:"policyName,omitempty" require:"true"`
  // {"en":"Policy description. You may describe the strategy here, including permission details, limited to a maximum of 250 characters.","zh_CN":"策略描述。可在此用文字简单描述策略包含权限内容，限制最多250字符"}
  Description *string `json:"description,omitempty" xml:"description,omitempty"`
  // {"en":"Permission policy content, please fill in the permission policy script as follows:\n\n\n[{\n\"effect\": \"allow\",\n\"action\": [\n\"productCode:actionCode\"\n],\n\"resource\": [\n\"*\"\n]\n}]\n\n\nA single permission policy can include permissions for multiple products, but CDN and non-CDN product permissions cannot be added to the same policy simultaneously.\n\nField descriptions:\n\n-effect: The authorization effect includes two types: allow and deny.\n\n-action: Describes the specific operations allowed or denied,  format: productCode:actionCode.\n\n-resource: The specific resources authorized. For all resources use, for specific resources refer to the format: wsc:&lt;service-name&gt;:&lt;region&gt;:&lt;account&gt;:&lt;relatice-id&gt;. Note: CDN products do not support specifying resources.","zh_CN":"权限策略内容，请填写权限策略脚本，示例如下：\n\n[{\n\"effect\": \"allow\",\n\"action\": [\n\"productCode:actionCode\"\n],\n\"resource\": [\n\"*\"\n]\n}]\n\n一个权限策略内可以包含多个产品的权限策略，但CDN和非CDN类产品权限不能同时添加到同一个策略里。\n\n各字段说明：\n\n-effect：授权效果包括两种允许（allow）和拒绝（deny）\n\n-action：描述允许或拒绝的特定操作，格式：productCode:actionCode\n\n-resource：授权的具体资源对象。若是全部资源用*表示，若是指定资源参考格式：wsc:&lt;service-name&gt;:&lt;region&gt;:&lt;account&gt;:&lt;relatice-id&gt;。注意：CDN类产品不支持指定资源"}
  PolicyDocument *string `json:"policyDocument,omitempty" xml:"policyDocument,omitempty" require:"true"`
}

func (s CreatePolicyRequest) String() string {
  return tea.Prettify(s)
}

func (s CreatePolicyRequest) GoString() string {
  return s.String()
}

func (s *CreatePolicyRequest) SetPolicyName(v string) *CreatePolicyRequest {
  s.PolicyName = &v
  return s
}

func (s *CreatePolicyRequest) SetDescription(v string) *CreatePolicyRequest {
  s.Description = &v
  return s
}

func (s *CreatePolicyRequest) SetPolicyDocument(v string) *CreatePolicyRequest {
  s.PolicyDocument = &v
  return s
}

type CreatePolicyRequestHeader struct {
}

func (s CreatePolicyRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s CreatePolicyRequestHeader) GoString() string {
  return s.String()
}

type CreatePolicyPaths struct {
}

func (s CreatePolicyPaths) String() string {
  return tea.Prettify(s)
}

func (s CreatePolicyPaths) GoString() string {
  return s.String()
}

type CreatePolicyParameters struct {
}

func (s CreatePolicyParameters) String() string {
  return tea.Prettify(s)
}

func (s CreatePolicyParameters) GoString() string {
  return s.String()
}

type CreatePolicyResponse struct {
  // {"en":"request result status code","zh_CN":"请求结果状态码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Request result information","zh_CN":"请求结果信息"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"Detailed data on the results of the request","zh_CN":"请求结果的详细数据"}
  Data *CreatePolicyResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
  // {"en":"Request ID","zh_CN":"请求ID"}
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty" require:"true"`
}

func (s CreatePolicyResponse) String() string {
  return tea.Prettify(s)
}

func (s CreatePolicyResponse) GoString() string {
  return s.String()
}

func (s *CreatePolicyResponse) SetCode(v string) *CreatePolicyResponse {
  s.Code = &v
  return s
}

func (s *CreatePolicyResponse) SetMsg(v string) *CreatePolicyResponse {
  s.Msg = &v
  return s
}

func (s *CreatePolicyResponse) SetData(v *CreatePolicyResponseData) *CreatePolicyResponse {
  s.Data = v
  return s
}

func (s *CreatePolicyResponse) SetRequestId(v string) *CreatePolicyResponse {
  s.RequestId = &v
  return s
}

type CreatePolicyResponseData struct {
  // {"en":"Policy ID","zh_CN":"策略ID"}
  PolicyId *int64 `json:"policyId,omitempty" xml:"policyId,omitempty" require:"true"`
}

func (s CreatePolicyResponseData) String() string {
  return tea.Prettify(s)
}

func (s CreatePolicyResponseData) GoString() string {
  return s.String()
}

func (s *CreatePolicyResponseData) SetPolicyId(v int64) *CreatePolicyResponseData {
  s.PolicyId = &v
  return s
}

type CreatePolicyResponseHeader struct {
}

func (s CreatePolicyResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s CreatePolicyResponseHeader) GoString() string {
  return s.String()
}




