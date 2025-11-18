package usermanage

// This file is auto-generated, don't edit it. Thanks.
import (
  "github.com/alibabacloud-go/tea/tea"
)

type DeleteSubAccountRequest struct {
}

func (s DeleteSubAccountRequest) String() string {
  return tea.Prettify(s)
}

func (s DeleteSubAccountRequest) GoString() string {
  return s.String()
}

type DeleteSubAccountRequestHeader struct {
}

func (s DeleteSubAccountRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteSubAccountRequestHeader) GoString() string {
  return s.String()
}

type DeleteSubAccountPaths struct {
  // {"en":"Sub account login name","zh_CN":"子用户登录名"}
  LoginName *string `json:"loginName,omitempty" xml:"loginName,omitempty" require:"true"`
}

func (s DeleteSubAccountPaths) String() string {
  return tea.Prettify(s)
}

func (s DeleteSubAccountPaths) GoString() string {
  return s.String()
}

func (s *DeleteSubAccountPaths) SetLoginName(v string) *DeleteSubAccountPaths {
  s.LoginName = &v
  return s
}

type DeleteSubAccountParameters struct {
}

func (s DeleteSubAccountParameters) String() string {
  return tea.Prettify(s)
}

func (s DeleteSubAccountParameters) GoString() string {
  return s.String()
}

type DeleteSubAccountResponse struct {
  // {"en":"Status Code","zh_CN":"请求结果状态码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"message","zh_CN":"请求结果信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s DeleteSubAccountResponse) String() string {
  return tea.Prettify(s)
}

func (s DeleteSubAccountResponse) GoString() string {
  return s.String()
}

func (s *DeleteSubAccountResponse) SetCode(v string) *DeleteSubAccountResponse {
  s.Code = &v
  return s
}

func (s *DeleteSubAccountResponse) SetMessage(v string) *DeleteSubAccountResponse {
  s.Message = &v
  return s
}

type DeleteSubAccountResponseHeader struct {
}

func (s DeleteSubAccountResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteSubAccountResponseHeader) GoString() string {
  return s.String()
}




type QuerySubAccountInfoRequest struct {
}

func (s QuerySubAccountInfoRequest) String() string {
  return tea.Prettify(s)
}

func (s QuerySubAccountInfoRequest) GoString() string {
  return s.String()
}

type QuerySubAccountInfoRequestHeader struct {
}

func (s QuerySubAccountInfoRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QuerySubAccountInfoRequestHeader) GoString() string {
  return s.String()
}

type QuerySubAccountInfoPaths struct {
  // {"en":"Login Name","zh_CN":"子账号登录名"}
  LoginName *string `json:"loginName,omitempty" xml:"loginName,omitempty" require:"true"`
}

func (s QuerySubAccountInfoPaths) String() string {
  return tea.Prettify(s)
}

func (s QuerySubAccountInfoPaths) GoString() string {
  return s.String()
}

func (s *QuerySubAccountInfoPaths) SetLoginName(v string) *QuerySubAccountInfoPaths {
  s.LoginName = &v
  return s
}

type QuerySubAccountInfoParameters struct {
}

func (s QuerySubAccountInfoParameters) String() string {
  return tea.Prettify(s)
}

func (s QuerySubAccountInfoParameters) GoString() string {
  return s.String()
}

type QuerySubAccountInfoResponse struct {
  // {"en":"Status Code","zh_CN":"错误具体状态码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"data","zh_CN":"返回数据"}
  Data *QuerySubAccountInfoResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
  // {"en":"Message","zh_CN":"消息提示"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s QuerySubAccountInfoResponse) String() string {
  return tea.Prettify(s)
}

func (s QuerySubAccountInfoResponse) GoString() string {
  return s.String()
}

func (s *QuerySubAccountInfoResponse) SetCode(v string) *QuerySubAccountInfoResponse {
  s.Code = &v
  return s
}

func (s *QuerySubAccountInfoResponse) SetData(v *QuerySubAccountInfoResponseData) *QuerySubAccountInfoResponse {
  s.Data = v
  return s
}

func (s *QuerySubAccountInfoResponse) SetMessage(v string) *QuerySubAccountInfoResponse {
  s.Message = &v
  return s
}

type QuerySubAccountInfoResponseData struct {
  // {"en":"create time","zh_CN":"创建时间"}
  CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty" require:"true"`
  // {"en":"consoleEnable","zh_CN":"是否允许登录控制台：1是 0 否"}
  ConsoleEnable *int `json:"consoleEnable,omitempty" xml:"consoleEnable,omitempty" require:"true"`
  // {"en":"display name","zh_CN":"称呼"}
  DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty" require:"true"`
  // {"en":"login name","zh_CN":"登录名"}
  LoginName *string `json:"loginName,omitempty" xml:"loginName,omitempty" require:"true"`
  // {"en":"mobile","zh_CN":"手机"}
  Mobile *string `json:"mobile,omitempty" xml:"mobile,omitempty" require:"true"`
  // {"en":"email","zh_CN":"邮箱"}
  Email *string `json:"email,omitempty" xml:"email,omitempty" require:"true"`
  // {"en":"status","zh_CN":"状态： 1 启用 0 停用"}
  Status *int `json:"status,omitempty" xml:"status,omitempty" require:"true"`
}

func (s QuerySubAccountInfoResponseData) String() string {
  return tea.Prettify(s)
}

func (s QuerySubAccountInfoResponseData) GoString() string {
  return s.String()
}

func (s *QuerySubAccountInfoResponseData) SetCreateTime(v string) *QuerySubAccountInfoResponseData {
  s.CreateTime = &v
  return s
}

func (s *QuerySubAccountInfoResponseData) SetConsoleEnable(v int) *QuerySubAccountInfoResponseData {
  s.ConsoleEnable = &v
  return s
}

func (s *QuerySubAccountInfoResponseData) SetDisplayName(v string) *QuerySubAccountInfoResponseData {
  s.DisplayName = &v
  return s
}

func (s *QuerySubAccountInfoResponseData) SetLoginName(v string) *QuerySubAccountInfoResponseData {
  s.LoginName = &v
  return s
}

func (s *QuerySubAccountInfoResponseData) SetMobile(v string) *QuerySubAccountInfoResponseData {
  s.Mobile = &v
  return s
}

func (s *QuerySubAccountInfoResponseData) SetEmail(v string) *QuerySubAccountInfoResponseData {
  s.Email = &v
  return s
}

func (s *QuerySubAccountInfoResponseData) SetStatus(v int) *QuerySubAccountInfoResponseData {
  s.Status = &v
  return s
}

type QuerySubAccountInfoResponseHeader struct {
}

func (s QuerySubAccountInfoResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QuerySubAccountInfoResponseHeader) GoString() string {
  return s.String()
}




type BatchAddOrRevokePolicyToSubAccountRequest struct {
  // {"en":"Specify policy ID","zh_CN":"指定权限策略id"}
  PolicyId []*int64 `json:"policyId,omitempty" xml:"policyId,omitempty" type:"Repeated"`
  // {"en":"Policy name","zh_CN":"策略名称"}
  PolicyName []*string `json:"policyName,omitempty" xml:"policyName,omitempty" type:"Repeated"`
  // {"en":"Sub account login name","zh_CN":"子用户登录名"}
  LoginName *string `json:"loginName,omitempty" xml:"loginName,omitempty" require:"true"`
  // {"en":"Select you want to add or revoke policy for sub account.\n\n0:add policy\n\n1:revoke policy","zh_CN":"选择需要为子用户添加或撤销权限策略\n\n0：添加权限\n\n1：撤销权限"}
  Type *int `json:"type,omitempty" xml:"type,omitempty" require:"true"`
}

func (s BatchAddOrRevokePolicyToSubAccountRequest) String() string {
  return tea.Prettify(s)
}

func (s BatchAddOrRevokePolicyToSubAccountRequest) GoString() string {
  return s.String()
}

func (s *BatchAddOrRevokePolicyToSubAccountRequest) SetPolicyId(v []*int64) *BatchAddOrRevokePolicyToSubAccountRequest {
  s.PolicyId = v
  return s
}

func (s *BatchAddOrRevokePolicyToSubAccountRequest) SetPolicyName(v []*string) *BatchAddOrRevokePolicyToSubAccountRequest {
  s.PolicyName = v
  return s
}

func (s *BatchAddOrRevokePolicyToSubAccountRequest) SetLoginName(v string) *BatchAddOrRevokePolicyToSubAccountRequest {
  s.LoginName = &v
  return s
}

func (s *BatchAddOrRevokePolicyToSubAccountRequest) SetType(v int) *BatchAddOrRevokePolicyToSubAccountRequest {
  s.Type = &v
  return s
}

type BatchAddOrRevokePolicyToSubAccountRequestHeader struct {
}

func (s BatchAddOrRevokePolicyToSubAccountRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s BatchAddOrRevokePolicyToSubAccountRequestHeader) GoString() string {
  return s.String()
}

type BatchAddOrRevokePolicyToSubAccountPaths struct {
}

func (s BatchAddOrRevokePolicyToSubAccountPaths) String() string {
  return tea.Prettify(s)
}

func (s BatchAddOrRevokePolicyToSubAccountPaths) GoString() string {
  return s.String()
}

type BatchAddOrRevokePolicyToSubAccountParameters struct {
}

func (s BatchAddOrRevokePolicyToSubAccountParameters) String() string {
  return tea.Prettify(s)
}

func (s BatchAddOrRevokePolicyToSubAccountParameters) GoString() string {
  return s.String()
}

type BatchAddOrRevokePolicyToSubAccountResponse struct {
  // {"en":"Request result status code","zh_CN":"请求结果状态码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Request result information","zh_CN":"请求结果信息"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
}

func (s BatchAddOrRevokePolicyToSubAccountResponse) String() string {
  return tea.Prettify(s)
}

func (s BatchAddOrRevokePolicyToSubAccountResponse) GoString() string {
  return s.String()
}

func (s *BatchAddOrRevokePolicyToSubAccountResponse) SetCode(v string) *BatchAddOrRevokePolicyToSubAccountResponse {
  s.Code = &v
  return s
}

func (s *BatchAddOrRevokePolicyToSubAccountResponse) SetMsg(v string) *BatchAddOrRevokePolicyToSubAccountResponse {
  s.Msg = &v
  return s
}

type BatchAddOrRevokePolicyToSubAccountResponseHeader struct {
}

func (s BatchAddOrRevokePolicyToSubAccountResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s BatchAddOrRevokePolicyToSubAccountResponseHeader) GoString() string {
  return s.String()
}




type AddSubAccountRequest struct {
  // {"en":"password","zh_CN":"密码"}
  Password *string `json:"password,omitempty" xml:"password,omitempty" require:"true"`
  // {"en":"parentLoginName","zh_CN":"父账号登录名"}
  ParentLoginName *string `json:"parentLoginName,omitempty" xml:"parentLoginName,omitempty" require:"true"`
  // {"en":"apiKey","zh_CN":"apiKey"}
  ApiKey *string `json:"apiKey,omitempty" xml:"apiKey,omitempty"`
  // {"en":"phone","zh_CN":"电话"}
  Phone *string `json:"phone,omitempty" xml:"phone,omitempty"`
  // {"en":"consoleEnable","zh_CN":"是否允许登录控制台：1是 0 否"}
  ConsoleEnable *int `json:"consoleEnable,omitempty" xml:"consoleEnable,omitempty"`
  // {"en":"display name","zh_CN":"称呼"}
  DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty" require:"true"`
  // {"en":"login name","zh_CN":"登录名"}
  LoginName *string `json:"loginName,omitempty" xml:"loginName,omitempty" require:"true"`
  // {"en":"mobile","zh_CN":"手机"}
  Mobile *string `json:"mobile,omitempty" xml:"mobile,omitempty"`
  // {"en":"programmaticEnable","zh_CN":"是否允许编程访问：1是 0 否"}
  ProgrammaticEnable *int `json:"programmaticEnable,omitempty" xml:"programmaticEnable,omitempty"`
  // {"en":"loginResetPassword","zh_CN":"登录是否需重置密码：1是 0 否"}
  LoginResetPassword *int `json:"loginResetPassword,omitempty" xml:"loginResetPassword,omitempty"`
  // {"en":"email","zh_CN":"邮箱"}
  Email *string `json:"email,omitempty" xml:"email,omitempty"`
  // {"en":"status","zh_CN":"状态： 1 启用 0 停用"}
  Status *int `json:"status,omitempty" xml:"status,omitempty" require:"true"`
}

func (s AddSubAccountRequest) String() string {
  return tea.Prettify(s)
}

func (s AddSubAccountRequest) GoString() string {
  return s.String()
}

func (s *AddSubAccountRequest) SetPassword(v string) *AddSubAccountRequest {
  s.Password = &v
  return s
}

func (s *AddSubAccountRequest) SetParentLoginName(v string) *AddSubAccountRequest {
  s.ParentLoginName = &v
  return s
}

func (s *AddSubAccountRequest) SetApiKey(v string) *AddSubAccountRequest {
  s.ApiKey = &v
  return s
}

func (s *AddSubAccountRequest) SetPhone(v string) *AddSubAccountRequest {
  s.Phone = &v
  return s
}

func (s *AddSubAccountRequest) SetConsoleEnable(v int) *AddSubAccountRequest {
  s.ConsoleEnable = &v
  return s
}

func (s *AddSubAccountRequest) SetDisplayName(v string) *AddSubAccountRequest {
  s.DisplayName = &v
  return s
}

func (s *AddSubAccountRequest) SetLoginName(v string) *AddSubAccountRequest {
  s.LoginName = &v
  return s
}

func (s *AddSubAccountRequest) SetMobile(v string) *AddSubAccountRequest {
  s.Mobile = &v
  return s
}

func (s *AddSubAccountRequest) SetProgrammaticEnable(v int) *AddSubAccountRequest {
  s.ProgrammaticEnable = &v
  return s
}

func (s *AddSubAccountRequest) SetLoginResetPassword(v int) *AddSubAccountRequest {
  s.LoginResetPassword = &v
  return s
}

func (s *AddSubAccountRequest) SetEmail(v string) *AddSubAccountRequest {
  s.Email = &v
  return s
}

func (s *AddSubAccountRequest) SetStatus(v int) *AddSubAccountRequest {
  s.Status = &v
  return s
}

type AddSubAccountRequestHeader struct {
}

func (s AddSubAccountRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s AddSubAccountRequestHeader) GoString() string {
  return s.String()
}

type AddSubAccountPaths struct {
}

func (s AddSubAccountPaths) String() string {
  return tea.Prettify(s)
}

func (s AddSubAccountPaths) GoString() string {
  return s.String()
}

type AddSubAccountParameters struct {
}

func (s AddSubAccountParameters) String() string {
  return tea.Prettify(s)
}

func (s AddSubAccountParameters) GoString() string {
  return s.String()
}

type AddSubAccountResponse struct {
  // {"en":"Status Code","zh_CN":"错误具体状态码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Message","zh_CN":"消息提示"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s AddSubAccountResponse) String() string {
  return tea.Prettify(s)
}

func (s AddSubAccountResponse) GoString() string {
  return s.String()
}

func (s *AddSubAccountResponse) SetCode(v string) *AddSubAccountResponse {
  s.Code = &v
  return s
}

func (s *AddSubAccountResponse) SetMessage(v string) *AddSubAccountResponse {
  s.Message = &v
  return s
}

type AddSubAccountResponseHeader struct {
}

func (s AddSubAccountResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s AddSubAccountResponseHeader) GoString() string {
  return s.String()
}




type GetSubAccountListRequest struct {
  // {"en":"Get a list of sub accounts under the main account","zh_CN":"主用户登录名"}
  LoginName *string `json:"loginName,omitempty" xml:"loginName,omitempty" require:"true"`
  // {"en":"Sub account loginName or displayName fuzzy query","zh_CN":"子账号loginName或displayName模糊查询"}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // {"en":"Page Number of Current Page.If it was empty, it would be treated as not being divided into pages. The contents filled in the pageSize field would not be effective","zh_CN":"指定分页查询时,当前页的页码。为空则不分页处理全部返回,pageSize字段填写内容不生效"}
  PageIndex *int `json:"pageIndex,omitempty" xml:"pageIndex,omitempty"`
  // {"en":"The maximum number of data displayed on each page.\nThe maximum value of the PageSize is 100. The number of data bar displayed on each page was 20 by default. When the PageSize value is empty, and pageIndex is not empty,20 data would be returned by default.","zh_CN":"指定分页查询时,每页显示的数据最大条数。\nPageSize参数最大取值为100。每页默认显示的数据条数为20条,PageSize参数值为空时,将默认返回20条数据。"}
  PageSize *int `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s GetSubAccountListRequest) String() string {
  return tea.Prettify(s)
}

func (s GetSubAccountListRequest) GoString() string {
  return s.String()
}

func (s *GetSubAccountListRequest) SetLoginName(v string) *GetSubAccountListRequest {
  s.LoginName = &v
  return s
}

func (s *GetSubAccountListRequest) SetName(v string) *GetSubAccountListRequest {
  s.Name = &v
  return s
}

func (s *GetSubAccountListRequest) SetPageIndex(v int) *GetSubAccountListRequest {
  s.PageIndex = &v
  return s
}

func (s *GetSubAccountListRequest) SetPageSize(v int) *GetSubAccountListRequest {
  s.PageSize = &v
  return s
}

type GetSubAccountListRequestHeader struct {
}

func (s GetSubAccountListRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetSubAccountListRequestHeader) GoString() string {
  return s.String()
}

type GetSubAccountListPaths struct {
}

func (s GetSubAccountListPaths) String() string {
  return tea.Prettify(s)
}

func (s GetSubAccountListPaths) GoString() string {
  return s.String()
}

type GetSubAccountListParameters struct {
}

func (s GetSubAccountListParameters) String() string {
  return tea.Prettify(s)
}

func (s GetSubAccountListParameters) GoString() string {
  return s.String()
}

type GetSubAccountListResponse struct {
  // {"en":"code","zh_CN":"请求结果状态码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"message","zh_CN":"请求结果信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"data","zh_CN":"返回值"}
  Data *GetSubAccountListResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
}

func (s GetSubAccountListResponse) String() string {
  return tea.Prettify(s)
}

func (s GetSubAccountListResponse) GoString() string {
  return s.String()
}

func (s *GetSubAccountListResponse) SetCode(v string) *GetSubAccountListResponse {
  s.Code = &v
  return s
}

func (s *GetSubAccountListResponse) SetMessage(v string) *GetSubAccountListResponse {
  s.Message = &v
  return s
}

func (s *GetSubAccountListResponse) SetData(v *GetSubAccountListResponseData) *GetSubAccountListResponse {
  s.Data = v
  return s
}

type GetSubAccountListResponseData struct {
  // {"en":"page Index","zh_CN":"页码"}
  PageIndex *int `json:"pageIndex,omitempty" xml:"pageIndex,omitempty" require:"true"`
  // {"en":"page Size","zh_CN":"每页个数"}
  PageSize *int `json:"pageSize,omitempty" xml:"pageSize,omitempty" require:"true"`
  // {"en":"total","zh_CN":"总条数"}
  Total *int64 `json:"total,omitempty" xml:"total,omitempty" require:"true"`
  // {"en":"rows","zh_CN":"每页数据"}
  Rows []*GetSubAccountListResponseDataRows `json:"rows,omitempty" xml:"rows,omitempty" require:"true" type:"Repeated"`
}

func (s GetSubAccountListResponseData) String() string {
  return tea.Prettify(s)
}

func (s GetSubAccountListResponseData) GoString() string {
  return s.String()
}

func (s *GetSubAccountListResponseData) SetPageIndex(v int) *GetSubAccountListResponseData {
  s.PageIndex = &v
  return s
}

func (s *GetSubAccountListResponseData) SetPageSize(v int) *GetSubAccountListResponseData {
  s.PageSize = &v
  return s
}

func (s *GetSubAccountListResponseData) SetTotal(v int64) *GetSubAccountListResponseData {
  s.Total = &v
  return s
}

func (s *GetSubAccountListResponseData) SetRows(v []*GetSubAccountListResponseDataRows) *GetSubAccountListResponseData {
  s.Rows = v
  return s
}

type GetSubAccountListResponseDataRows struct     {
  // {"en":"Sub account login name","zh_CN":"子用户登录名"}
  LoginName *string `json:"loginName,omitempty" xml:"loginName,omitempty" require:"true"`
  // {"en":"Sub account display name","zh_CN":"子用户显示名"}
  DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty" require:"true"`
  // {"en":"Sub accout's E-mail","zh_CN":"绑定的邮箱"}
  Email *string `json:"email,omitempty" xml:"email,omitempty" require:"true"`
  // {"en":"Sub accout's mobile phone","zh_CN":"绑定的手机"}
  Mobile *string `json:"mobile,omitempty" xml:"mobile,omitempty" require:"true"`
  // {"en":"CreateTime","zh_CN":"创建时间"}
  CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty" require:"true"`
  // {"en":"The status of accout1:Activate0:Disable","zh_CN":"账号启用/禁用状态,1代表启用,0代表禁用"}
  Status *int `json:"status,omitempty" xml:"status,omitempty" require:"true"`
}

func (s GetSubAccountListResponseDataRows) String() string {
  return tea.Prettify(s)
}

func (s GetSubAccountListResponseDataRows) GoString() string {
  return s.String()
}

func (s *GetSubAccountListResponseDataRows) SetLoginName(v string) *GetSubAccountListResponseDataRows {
  s.LoginName = &v
  return s
}

func (s *GetSubAccountListResponseDataRows) SetDisplayName(v string) *GetSubAccountListResponseDataRows {
  s.DisplayName = &v
  return s
}

func (s *GetSubAccountListResponseDataRows) SetEmail(v string) *GetSubAccountListResponseDataRows {
  s.Email = &v
  return s
}

func (s *GetSubAccountListResponseDataRows) SetMobile(v string) *GetSubAccountListResponseDataRows {
  s.Mobile = &v
  return s
}

func (s *GetSubAccountListResponseDataRows) SetCreateTime(v string) *GetSubAccountListResponseDataRows {
  s.CreateTime = &v
  return s
}

func (s *GetSubAccountListResponseDataRows) SetStatus(v int) *GetSubAccountListResponseDataRows {
  s.Status = &v
  return s
}

type GetSubAccountListResponseHeader struct {
}

func (s GetSubAccountListResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetSubAccountListResponseHeader) GoString() string {
  return s.String()
}




type QueryPolicyAttachedMainAccountOrSubAccountRequest struct {
}

func (s QueryPolicyAttachedMainAccountOrSubAccountRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryPolicyAttachedMainAccountOrSubAccountRequest) GoString() string {
  return s.String()
}

type QueryPolicyAttachedMainAccountOrSubAccountRequestHeader struct {
  // {"en":"Select the specified language and return the policy description of the corresponding language. Optional values:zh_CN，en，ko_KR，ja_JP；Default language is en","zh_CN":"选择指定语言返回对应语言的策略描述，可选值：zh_CN，en，ko_KR，ja_JP；未选择默认en"}
  AcceptLanguage *string `json:"Accept-Language,omitempty" xml:"Accept-Language,omitempty" require:"true"`
}

func (s QueryPolicyAttachedMainAccountOrSubAccountRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryPolicyAttachedMainAccountOrSubAccountRequestHeader) GoString() string {
  return s.String()
}

func (s *QueryPolicyAttachedMainAccountOrSubAccountRequestHeader) SetAcceptLanguage(v string) *QueryPolicyAttachedMainAccountOrSubAccountRequestHeader {
  s.AcceptLanguage = &v
  return s
}

type QueryPolicyAttachedMainAccountOrSubAccountPaths struct {
  // {"en":"loginName(Main or ordinary subAccount are available)","zh_CN":"用户登录名（可传主子用户）"}
  LoginName *string `json:"loginName,omitempty" xml:"loginName,omitempty" require:"true"`
}

func (s QueryPolicyAttachedMainAccountOrSubAccountPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryPolicyAttachedMainAccountOrSubAccountPaths) GoString() string {
  return s.String()
}

func (s *QueryPolicyAttachedMainAccountOrSubAccountPaths) SetLoginName(v string) *QueryPolicyAttachedMainAccountOrSubAccountPaths {
  s.LoginName = &v
  return s
}

type QueryPolicyAttachedMainAccountOrSubAccountParameters struct {
}

func (s QueryPolicyAttachedMainAccountOrSubAccountParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryPolicyAttachedMainAccountOrSubAccountParameters) GoString() string {
  return s.String()
}

type QueryPolicyAttachedMainAccountOrSubAccountResponse struct {
  // {"en":"Status Code","zh_CN":"请求结果状态码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Message","zh_CN":"请求结果信息"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"","zh_CN":""}
  Data []*QueryPolicyAttachedMainAccountOrSubAccountResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
  // {"en":"Request ID","zh_CN":"请求ID"}
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty" require:"true"`
}

func (s QueryPolicyAttachedMainAccountOrSubAccountResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryPolicyAttachedMainAccountOrSubAccountResponse) GoString() string {
  return s.String()
}

func (s *QueryPolicyAttachedMainAccountOrSubAccountResponse) SetCode(v string) *QueryPolicyAttachedMainAccountOrSubAccountResponse {
  s.Code = &v
  return s
}

func (s *QueryPolicyAttachedMainAccountOrSubAccountResponse) SetMsg(v string) *QueryPolicyAttachedMainAccountOrSubAccountResponse {
  s.Msg = &v
  return s
}

func (s *QueryPolicyAttachedMainAccountOrSubAccountResponse) SetData(v []*QueryPolicyAttachedMainAccountOrSubAccountResponseData) *QueryPolicyAttachedMainAccountOrSubAccountResponse {
  s.Data = v
  return s
}

func (s *QueryPolicyAttachedMainAccountOrSubAccountResponse) SetRequestId(v string) *QueryPolicyAttachedMainAccountOrSubAccountResponse {
  s.RequestId = &v
  return s
}

type QueryPolicyAttachedMainAccountOrSubAccountResponseData struct     {
  // {"en":"policyId","zh_CN":"策略id"}
  PolicyId *int64 `json:"policyId,omitempty" xml:"policyId,omitempty" require:"true"`
  // {"en":"policy name","zh_CN":"策略名称"}
  PolicyName *string `json:"policyName,omitempty" xml:"policyName,omitempty" require:"true"`
  // {"en":"Policy description","zh_CN":"策略描述内容"}
  PolicyDescribe *string `json:"policyDescribe,omitempty" xml:"policyDescribe,omitempty" require:"true"`
  // {"en":"policy type","zh_CN":"策略类型：system：系统策略、custom自定义策略"}
  PolicyType *string `json:"policyType,omitempty" xml:"policyType,omitempty" require:"true"`
}

func (s QueryPolicyAttachedMainAccountOrSubAccountResponseData) String() string {
  return tea.Prettify(s)
}

func (s QueryPolicyAttachedMainAccountOrSubAccountResponseData) GoString() string {
  return s.String()
}

func (s *QueryPolicyAttachedMainAccountOrSubAccountResponseData) SetPolicyId(v int64) *QueryPolicyAttachedMainAccountOrSubAccountResponseData {
  s.PolicyId = &v
  return s
}

func (s *QueryPolicyAttachedMainAccountOrSubAccountResponseData) SetPolicyName(v string) *QueryPolicyAttachedMainAccountOrSubAccountResponseData {
  s.PolicyName = &v
  return s
}

func (s *QueryPolicyAttachedMainAccountOrSubAccountResponseData) SetPolicyDescribe(v string) *QueryPolicyAttachedMainAccountOrSubAccountResponseData {
  s.PolicyDescribe = &v
  return s
}

func (s *QueryPolicyAttachedMainAccountOrSubAccountResponseData) SetPolicyType(v string) *QueryPolicyAttachedMainAccountOrSubAccountResponseData {
  s.PolicyType = &v
  return s
}

type QueryPolicyAttachedMainAccountOrSubAccountResponseHeader struct {
}

func (s QueryPolicyAttachedMainAccountOrSubAccountResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryPolicyAttachedMainAccountOrSubAccountResponseHeader) GoString() string {
  return s.String()
}




type UpdateSubAccountRequest struct {
  // {"en":"area code","zh_CN":"手机号区号"}
  AreaCode *string `json:"areaCode,omitempty" xml:"areaCode,omitempty"`
  // {"en":"password","zh_CN":"密码"}
  Password *string `json:"password,omitempty" xml:"password,omitempty"`
  // {"en":"openApiStatus","zh_CN":"是否开启OpenAPI 0否，1是"}
  OpenApiStatus *int `json:"openApiStatus,omitempty" xml:"openApiStatus,omitempty"`
  // {"en":"Programmatic Access or not 1:Yes0:No","zh_CN":"子用户是否允许登录控制台：1是 0 否"}
  ConsoleEnable *int `json:"consoleEnable,omitempty" xml:"consoleEnable,omitempty"`
  // {"en":"display name","zh_CN":"子用户显示名称"}
  DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
  // {"en":"Sub account login name","zh_CN":"子用户登录名"}
  LoginName *string `json:"loginName,omitempty" xml:"loginName,omitempty" require:"true"`
  // {"en":"mobile","zh_CN":"手机"}
  Mobile *string `json:"mobile,omitempty" xml:"mobile,omitempty"`
  // {"en":"loginResetPassword","zh_CN":"登录是否重置密码 1 是 0  否"}
  LoginResetPassword *int `json:"loginResetPassword,omitempty" xml:"loginResetPassword,omitempty"`
  // {"en":"email","zh_CN":"邮箱"}
  Email *string `json:"email,omitempty" xml:"email,omitempty"`
  // {"en":"","zh_CN":"账号状态：0：禁用，1：启用"}
  Status *int `json:"status,omitempty" xml:"status,omitempty"`
}

func (s UpdateSubAccountRequest) String() string {
  return tea.Prettify(s)
}

func (s UpdateSubAccountRequest) GoString() string {
  return s.String()
}

func (s *UpdateSubAccountRequest) SetAreaCode(v string) *UpdateSubAccountRequest {
  s.AreaCode = &v
  return s
}

func (s *UpdateSubAccountRequest) SetPassword(v string) *UpdateSubAccountRequest {
  s.Password = &v
  return s
}

func (s *UpdateSubAccountRequest) SetOpenApiStatus(v int) *UpdateSubAccountRequest {
  s.OpenApiStatus = &v
  return s
}

func (s *UpdateSubAccountRequest) SetConsoleEnable(v int) *UpdateSubAccountRequest {
  s.ConsoleEnable = &v
  return s
}

func (s *UpdateSubAccountRequest) SetDisplayName(v string) *UpdateSubAccountRequest {
  s.DisplayName = &v
  return s
}

func (s *UpdateSubAccountRequest) SetLoginName(v string) *UpdateSubAccountRequest {
  s.LoginName = &v
  return s
}

func (s *UpdateSubAccountRequest) SetMobile(v string) *UpdateSubAccountRequest {
  s.Mobile = &v
  return s
}

func (s *UpdateSubAccountRequest) SetLoginResetPassword(v int) *UpdateSubAccountRequest {
  s.LoginResetPassword = &v
  return s
}

func (s *UpdateSubAccountRequest) SetEmail(v string) *UpdateSubAccountRequest {
  s.Email = &v
  return s
}

func (s *UpdateSubAccountRequest) SetStatus(v int) *UpdateSubAccountRequest {
  s.Status = &v
  return s
}

type UpdateSubAccountRequestHeader struct {
}

func (s UpdateSubAccountRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s UpdateSubAccountRequestHeader) GoString() string {
  return s.String()
}

type UpdateSubAccountPaths struct {
}

func (s UpdateSubAccountPaths) String() string {
  return tea.Prettify(s)
}

func (s UpdateSubAccountPaths) GoString() string {
  return s.String()
}

type UpdateSubAccountParameters struct {
}

func (s UpdateSubAccountParameters) String() string {
  return tea.Prettify(s)
}

func (s UpdateSubAccountParameters) GoString() string {
  return s.String()
}

type UpdateSubAccountResponse struct {
  // {"en":"Status Code","zh_CN":"请求结果状态码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"message","zh_CN":"请求结果信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s UpdateSubAccountResponse) String() string {
  return tea.Prettify(s)
}

func (s UpdateSubAccountResponse) GoString() string {
  return s.String()
}

func (s *UpdateSubAccountResponse) SetCode(v string) *UpdateSubAccountResponse {
  s.Code = &v
  return s
}

func (s *UpdateSubAccountResponse) SetMessage(v string) *UpdateSubAccountResponse {
  s.Message = &v
  return s
}

type UpdateSubAccountResponseHeader struct {
}

func (s UpdateSubAccountResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s UpdateSubAccountResponseHeader) GoString() string {
  return s.String()
}




