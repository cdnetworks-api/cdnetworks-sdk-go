package cgmanage

// This file is auto-generated, don't edit it. Thanks.
import (
  "github.com/alibabacloud-go/tea/tea"
)

type QueryCustomizedControlGroupByNameRequest struct {
  // {"en":"ControlGroup Code","zh_CN":"ControlGroup 编码"}
  ControlGroupCode *string `json:"controlGroupCode,omitempty" xml:"controlGroupCode,omitempty"`
  // {"en":"controlGroup Name","zh_CN":"ControlGroup 名称"}
  ControlGroupName *string `json:"controlGroupName,omitempty" xml:"controlGroupName,omitempty"`
}

func (s QueryCustomizedControlGroupByNameRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryCustomizedControlGroupByNameRequest) GoString() string {
  return s.String()
}

func (s *QueryCustomizedControlGroupByNameRequest) SetControlGroupCode(v string) *QueryCustomizedControlGroupByNameRequest {
  s.ControlGroupCode = &v
  return s
}

func (s *QueryCustomizedControlGroupByNameRequest) SetControlGroupName(v string) *QueryCustomizedControlGroupByNameRequest {
  s.ControlGroupName = &v
  return s
}

type QueryCustomizedControlGroupByNameRequestHeader struct {
}

func (s QueryCustomizedControlGroupByNameRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryCustomizedControlGroupByNameRequestHeader) GoString() string {
  return s.String()
}

type QueryCustomizedControlGroupByNamePaths struct {
}

func (s QueryCustomizedControlGroupByNamePaths) String() string {
  return tea.Prettify(s)
}

func (s QueryCustomizedControlGroupByNamePaths) GoString() string {
  return s.String()
}

type QueryCustomizedControlGroupByNameParameters struct {
}

func (s QueryCustomizedControlGroupByNameParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryCustomizedControlGroupByNameParameters) GoString() string {
  return s.String()
}

type QueryCustomizedControlGroupByNameResponse struct {
  // {"en":"Message","zh_CN":"消息提示"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"Status Code, 200:success,  500: failed","zh_CN":"错误具体状态码，200:success， 500: failed"}
  Code *int `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"","zh_CN":""}
  Data *QueryCustomizedControlGroupByNameResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
  // {"en":"Request ID","zh_CN":"请求ID"}
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty" require:"true"`
}

func (s QueryCustomizedControlGroupByNameResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryCustomizedControlGroupByNameResponse) GoString() string {
  return s.String()
}

func (s *QueryCustomizedControlGroupByNameResponse) SetMsg(v string) *QueryCustomizedControlGroupByNameResponse {
  s.Msg = &v
  return s
}

func (s *QueryCustomizedControlGroupByNameResponse) SetCode(v int) *QueryCustomizedControlGroupByNameResponse {
  s.Code = &v
  return s
}

func (s *QueryCustomizedControlGroupByNameResponse) SetData(v *QueryCustomizedControlGroupByNameResponseData) *QueryCustomizedControlGroupByNameResponse {
  s.Data = v
  return s
}

func (s *QueryCustomizedControlGroupByNameResponse) SetRequestId(v string) *QueryCustomizedControlGroupByNameResponse {
  s.RequestId = &v
  return s
}

type QueryCustomizedControlGroupByNameResponseData struct {
  // {"en":"Control Group Code","zh_CN":"Control Group编号"}
  ControlGroupCode *string `json:"controlGroupCode,omitempty" xml:"controlGroupCode,omitempty" require:"true"`
  // {"en":"Domain array, Used to specify the domain contained in the Control Group","zh_CN":"域名字符串数组，用来指定该Control Group所包含的域名"}
  DomainList []*string `json:"domainList,omitempty" xml:"domainList,omitempty" require:"true" type:"Repeated"`
  // {"en":"Account object array, Used to specify accounts with permission.","zh_CN":"账号对象数组，用来指定有权限访问的账号。"}
  AccountNameList []*string `json:"accountNameList,omitempty" xml:"accountNameList,omitempty" require:"true" type:"Repeated"`
  // {"en":"Control Group Name","zh_CN":"Control Group名称"}
  ControlGroupName *string `json:"controlGroupName,omitempty" xml:"controlGroupName,omitempty" require:"true"`
}

func (s QueryCustomizedControlGroupByNameResponseData) String() string {
  return tea.Prettify(s)
}

func (s QueryCustomizedControlGroupByNameResponseData) GoString() string {
  return s.String()
}

func (s *QueryCustomizedControlGroupByNameResponseData) SetControlGroupCode(v string) *QueryCustomizedControlGroupByNameResponseData {
  s.ControlGroupCode = &v
  return s
}

func (s *QueryCustomizedControlGroupByNameResponseData) SetDomainList(v []*string) *QueryCustomizedControlGroupByNameResponseData {
  s.DomainList = v
  return s
}

func (s *QueryCustomizedControlGroupByNameResponseData) SetAccountNameList(v []*string) *QueryCustomizedControlGroupByNameResponseData {
  s.AccountNameList = v
  return s
}

func (s *QueryCustomizedControlGroupByNameResponseData) SetControlGroupName(v string) *QueryCustomizedControlGroupByNameResponseData {
  s.ControlGroupName = &v
  return s
}

type QueryCustomizedControlGroupByNameResponseHeader struct {
}

func (s QueryCustomizedControlGroupByNameResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryCustomizedControlGroupByNameResponseHeader) GoString() string {
  return s.String()
}




type CreateControlGroupRequest struct {
  // {"en":"Domain array, Used to specify the domain contained in the Control Group","zh_CN":"域名字符串数组，用来指定该Control Group所包含的域名"}
  DomainList []*string `json:"domainList,omitempty" xml:"domainList,omitempty" require:"true" type:"Repeated"`
  // {"en":"Account object array, Used to specify accounts with permission. (If the account created this Control Group also needs the permission, please remember to add this account into accountList)","zh_CN":"账号对象数组，用来指定有权限访问的账号。（若创建这个Control Group的账号也需要权限记得把自己的账号也加进去）"}
  AccountList []*CreateControlGroupRequestAccountList `json:"accountList,omitempty" xml:"accountList,omitempty" type:"Repeated"`
  // {"en":"Control Group name","zh_CN":"Control Group名称"}
  ControlGroupName *string `json:"controlGroupName,omitempty" xml:"controlGroupName,omitempty" require:"true"`
}

func (s CreateControlGroupRequest) String() string {
  return tea.Prettify(s)
}

func (s CreateControlGroupRequest) GoString() string {
  return s.String()
}

func (s *CreateControlGroupRequest) SetDomainList(v []*string) *CreateControlGroupRequest {
  s.DomainList = v
  return s
}

func (s *CreateControlGroupRequest) SetAccountList(v []*CreateControlGroupRequestAccountList) *CreateControlGroupRequest {
  s.AccountList = v
  return s
}

func (s *CreateControlGroupRequest) SetControlGroupName(v string) *CreateControlGroupRequest {
  s.ControlGroupName = &v
  return s
}

type CreateControlGroupRequestAccountList struct     {
  // {"en":"Account","zh_CN":"账号"}
  LoginName *string `json:"loginName,omitempty" xml:"loginName,omitempty"`
}

func (s CreateControlGroupRequestAccountList) String() string {
  return tea.Prettify(s)
}

func (s CreateControlGroupRequestAccountList) GoString() string {
  return s.String()
}

func (s *CreateControlGroupRequestAccountList) SetLoginName(v string) *CreateControlGroupRequestAccountList {
  s.LoginName = &v
  return s
}

type CreateControlGroupRequestHeader struct {
}

func (s CreateControlGroupRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s CreateControlGroupRequestHeader) GoString() string {
  return s.String()
}

type CreateControlGroupPaths struct {
}

func (s CreateControlGroupPaths) String() string {
  return tea.Prettify(s)
}

func (s CreateControlGroupPaths) GoString() string {
  return s.String()
}

type CreateControlGroupParameters struct {
}

func (s CreateControlGroupParameters) String() string {
  return tea.Prettify(s)
}

func (s CreateControlGroupParameters) GoString() string {
  return s.String()
}

type CreateControlGroupResponse struct {
  // {"en":"Message","zh_CN":"消息提示"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"Status Code","zh_CN":"错误具体状态码"}
  Code *int `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"","zh_CN":""}
  Data *CreateControlGroupResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
  // {"en":"Request ID","zh_CN":"请求ID"}
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty" require:"true"`
}

func (s CreateControlGroupResponse) String() string {
  return tea.Prettify(s)
}

func (s CreateControlGroupResponse) GoString() string {
  return s.String()
}

func (s *CreateControlGroupResponse) SetMsg(v string) *CreateControlGroupResponse {
  s.Msg = &v
  return s
}

func (s *CreateControlGroupResponse) SetCode(v int) *CreateControlGroupResponse {
  s.Code = &v
  return s
}

func (s *CreateControlGroupResponse) SetData(v *CreateControlGroupResponseData) *CreateControlGroupResponse {
  s.Data = v
  return s
}

func (s *CreateControlGroupResponse) SetRequestId(v string) *CreateControlGroupResponse {
  s.RequestId = &v
  return s
}

type CreateControlGroupResponseData struct {
  // {"en":"Control Group Code","zh_CN":"Control Group编号"}
  ControlGroupCode *string `json:"controlGroupCode,omitempty" xml:"controlGroupCode,omitempty" require:"true"`
  // {"en":"Control Group Name","zh_CN":"Control Group名称"}
  ControlGroupName *string `json:"controlGroupName,omitempty" xml:"controlGroupName,omitempty" require:"true"`
}

func (s CreateControlGroupResponseData) String() string {
  return tea.Prettify(s)
}

func (s CreateControlGroupResponseData) GoString() string {
  return s.String()
}

func (s *CreateControlGroupResponseData) SetControlGroupCode(v string) *CreateControlGroupResponseData {
  s.ControlGroupCode = &v
  return s
}

func (s *CreateControlGroupResponseData) SetControlGroupName(v string) *CreateControlGroupResponseData {
  s.ControlGroupName = &v
  return s
}

type CreateControlGroupResponseHeader struct {
}

func (s CreateControlGroupResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s CreateControlGroupResponseHeader) GoString() string {
  return s.String()
}




type GetDomainListOfControlGroupRequest struct {
  // {"en":"Control Group Code, If this parameter is empty, all authorized control groups will be returned","zh_CN":"服务组编号，不传则返回全部有权限的服务组"}
  ControlGroupCode []*string `json:"controlGroupCode,omitempty" xml:"controlGroupCode,omitempty" type:"Repeated"`
}

func (s GetDomainListOfControlGroupRequest) String() string {
  return tea.Prettify(s)
}

func (s GetDomainListOfControlGroupRequest) GoString() string {
  return s.String()
}

func (s *GetDomainListOfControlGroupRequest) SetControlGroupCode(v []*string) *GetDomainListOfControlGroupRequest {
  s.ControlGroupCode = v
  return s
}

type GetDomainListOfControlGroupRequestHeader struct {
}

func (s GetDomainListOfControlGroupRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetDomainListOfControlGroupRequestHeader) GoString() string {
  return s.String()
}

type GetDomainListOfControlGroupPaths struct {
}

func (s GetDomainListOfControlGroupPaths) String() string {
  return tea.Prettify(s)
}

func (s GetDomainListOfControlGroupPaths) GoString() string {
  return s.String()
}

type GetDomainListOfControlGroupParameters struct {
}

func (s GetDomainListOfControlGroupParameters) String() string {
  return tea.Prettify(s)
}

func (s GetDomainListOfControlGroupParameters) GoString() string {
  return s.String()
}

type GetDomainListOfControlGroupResponse struct {
  // {"en":"","zh_CN":""}
  ControlGroupDetail []*GetDomainListOfControlGroupResponseControlGroupDetail `json:"controlGroupDetail,omitempty" xml:"controlGroupDetail,omitempty" require:"true" type:"Repeated"`
}

func (s GetDomainListOfControlGroupResponse) String() string {
  return tea.Prettify(s)
}

func (s GetDomainListOfControlGroupResponse) GoString() string {
  return s.String()
}

func (s *GetDomainListOfControlGroupResponse) SetControlGroupDetail(v []*GetDomainListOfControlGroupResponseControlGroupDetail) *GetDomainListOfControlGroupResponse {
  s.ControlGroupDetail = v
  return s
}

type GetDomainListOfControlGroupResponseControlGroupDetail struct     {
  // {"en":"Control Group Code","zh_CN":"服务组编号"}
  ControlGroupCode *string `json:"controlGroupCode,omitempty" xml:"controlGroupCode,omitempty" require:"true"`
  // {"en":"domain list","zh_CN":"域名列表"}
  DomainList []*string `json:"domainList,omitempty" xml:"domainList,omitempty" require:"true" type:"Repeated"`
  // {"en":"Control Group Name","zh_CN":"服务组名称"}
  ControlGroupName *string `json:"controlGroupName,omitempty" xml:"controlGroupName,omitempty" require:"true"`
}

func (s GetDomainListOfControlGroupResponseControlGroupDetail) String() string {
  return tea.Prettify(s)
}

func (s GetDomainListOfControlGroupResponseControlGroupDetail) GoString() string {
  return s.String()
}

func (s *GetDomainListOfControlGroupResponseControlGroupDetail) SetControlGroupCode(v string) *GetDomainListOfControlGroupResponseControlGroupDetail {
  s.ControlGroupCode = &v
  return s
}

func (s *GetDomainListOfControlGroupResponseControlGroupDetail) SetDomainList(v []*string) *GetDomainListOfControlGroupResponseControlGroupDetail {
  s.DomainList = v
  return s
}

func (s *GetDomainListOfControlGroupResponseControlGroupDetail) SetControlGroupName(v string) *GetDomainListOfControlGroupResponseControlGroupDetail {
  s.ControlGroupName = &v
  return s
}

type GetDomainListOfControlGroupResponseHeader struct {
}

func (s GetDomainListOfControlGroupResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetDomainListOfControlGroupResponseHeader) GoString() string {
  return s.String()
}




type QueryCustomizedControlGroupRequest struct {
}

func (s QueryCustomizedControlGroupRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryCustomizedControlGroupRequest) GoString() string {
  return s.String()
}

type QueryCustomizedControlGroupResponse struct {
  // {"en":"Status Code, 200:success,  500: failed", "zh_CN":"错误具体状态码，200:success， 500: failed"}
  Code *int `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Message", "zh_CN":"消息提示"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"Request ID","zh_CN":"请求ID"}
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty" require:"true"`
  Data *QueryCustomizedControlGroupResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
}

func (s QueryCustomizedControlGroupResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryCustomizedControlGroupResponse) GoString() string {
  return s.String()
}

func (s *QueryCustomizedControlGroupResponse) SetCode(v int) *QueryCustomizedControlGroupResponse {
  s.Code = &v
  return s
}

func (s *QueryCustomizedControlGroupResponse) SetMsg(v string) *QueryCustomizedControlGroupResponse {
  s.Msg = &v
  return s
}

func (s *QueryCustomizedControlGroupResponse) SetRequestId(v string) *QueryCustomizedControlGroupResponse {
  s.RequestId = &v
  return s
}

func (s *QueryCustomizedControlGroupResponse) SetData(v *QueryCustomizedControlGroupResponseData) *QueryCustomizedControlGroupResponse {
  s.Data = v
  return s
}

type QueryCustomizedControlGroupResponseData struct {
  // {"en":"Control Group Name", "zh_CN":"Control Group名称"}
  ControlGroupName *string `json:"controlGroupName,omitempty" xml:"controlGroupName,omitempty" require:"true"`
  // {"en":"Control Group Code", "zh_CN":"Control Group编号"}
  ControlGroupCode *string `json:"controlGroupCode,omitempty" xml:"controlGroupCode,omitempty" require:"true"`
  // {"en":"Account object array, Used to specify accounts with permission.", "zh_CN":"账号对象数组，用来指定有权限访问的账号。"}
  AccountNameList []*string `json:"accountNameList,omitempty" xml:"accountNameList,omitempty" require:"true" type:"Repeated"`
  // {"en":"Domain array, Used to specify the domain contained in the Control Group", "zh_CN":"域名字符串数组，用来指定该Control Group所包含的域名"}
  DomainList []*string `json:"domainList,omitempty" xml:"domainList,omitempty" require:"true" type:"Repeated"`
}

func (s QueryCustomizedControlGroupResponseData) String() string {
  return tea.Prettify(s)
}

func (s QueryCustomizedControlGroupResponseData) GoString() string {
  return s.String()
}

func (s *QueryCustomizedControlGroupResponseData) SetControlGroupName(v string) *QueryCustomizedControlGroupResponseData {
  s.ControlGroupName = &v
  return s
}

func (s *QueryCustomizedControlGroupResponseData) SetControlGroupCode(v string) *QueryCustomizedControlGroupResponseData {
  s.ControlGroupCode = &v
  return s
}

func (s *QueryCustomizedControlGroupResponseData) SetAccountNameList(v []*string) *QueryCustomizedControlGroupResponseData {
  s.AccountNameList = v
  return s
}

func (s *QueryCustomizedControlGroupResponseData) SetDomainList(v []*string) *QueryCustomizedControlGroupResponseData {
  s.DomainList = v
  return s
}

type QueryCustomizedControlGroupPaths struct {
  // {"en":"Control Group Code", "zh_CN":"Control Group 编号"}
  ControlGroupCode *string `json:"ControlGroupCode,omitempty" xml:"ControlGroupCode,omitempty" require:"true"`
}

func (s QueryCustomizedControlGroupPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryCustomizedControlGroupPaths) GoString() string {
  return s.String()
}

func (s *QueryCustomizedControlGroupPaths) SetControlGroupCode(v string) *QueryCustomizedControlGroupPaths {
  s.ControlGroupCode = &v
  return s
}

type QueryCustomizedControlGroupParameters struct {
}

func (s QueryCustomizedControlGroupParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryCustomizedControlGroupParameters) GoString() string {
  return s.String()
}

type QueryCustomizedControlGroupRequestHeader struct {
}

func (s QueryCustomizedControlGroupRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryCustomizedControlGroupRequestHeader) GoString() string {
  return s.String()
}

type QueryCustomizedControlGroupResponseHeader struct {
}

func (s QueryCustomizedControlGroupResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryCustomizedControlGroupResponseHeader) GoString() string {
  return s.String()
}




type EditControlGroupRequest struct {
  // {"en":"Domain array, which only the User Customized type Control Group can be modified, customer type Control Group and product type Control Group can not be modified.User Customized type Control Group empties the original domainList if no value is passed","zh_CN":"域名字符串数组，只有自定义类型的Control Group可做修改，若是客户类型与合同类型Control Group则不做修改。自定义类型Control Group若不传值则将原domainList清空"}
  DomainList []*string `json:"domainList,omitempty" xml:"domainList,omitempty" type:"Repeated"`
  // {"en":"Account object array,Used to specify accounts with permission.  all types of Control Group can be modified, if no value is passed, the original accountList will be emptied","zh_CN":"账号对象数组, 用来指定有权限访问的账号。客户类型，合同类型与自定义类型的Control Group都可以做修改，若不传值则将原accountList清空"}
  AccountList []*EditControlGroupRequestAccountList `json:"accountList,omitempty" xml:"accountList,omitempty" type:"Repeated"`
  // {"en":"Whether to add:\n1. Do not pass or pass false: rewrite method;\n2. Pass true: append method.","zh_CN":"是否追加:\n1.不传或false：覆盖方式;\n2.传true：追加方式."}
  IsAdd *bool `json:"isAdd,omitempty" xml:"isAdd,omitempty"`
  // {"en":"Control Group name, which only the User Customized type Control Group can be modified, customer type Control Group and product type Control Group can not be modified. User Customized type Control Group keeps the original Control Group name if no value is passed","zh_CN":"Control Group名称，只有自定义类型的Control Group可做修改，若是客户类型与合同类型Control Group则不做修改。自定义类型Control Group若不传值则保持原来的Control Group名称"}
  ControlGroupName *string `json:"controlGroupName,omitempty" xml:"controlGroupName,omitempty"`
}

func (s EditControlGroupRequest) String() string {
  return tea.Prettify(s)
}

func (s EditControlGroupRequest) GoString() string {
  return s.String()
}

func (s *EditControlGroupRequest) SetDomainList(v []*string) *EditControlGroupRequest {
  s.DomainList = v
  return s
}

func (s *EditControlGroupRequest) SetAccountList(v []*EditControlGroupRequestAccountList) *EditControlGroupRequest {
  s.AccountList = v
  return s
}

func (s *EditControlGroupRequest) SetIsAdd(v bool) *EditControlGroupRequest {
  s.IsAdd = &v
  return s
}

func (s *EditControlGroupRequest) SetControlGroupName(v string) *EditControlGroupRequest {
  s.ControlGroupName = &v
  return s
}

type EditControlGroupRequestAccountList struct     {
  // {"en":"Account","zh_CN":"账号"}
  LoginName *string `json:"loginName,omitempty" xml:"loginName,omitempty"`
}

func (s EditControlGroupRequestAccountList) String() string {
  return tea.Prettify(s)
}

func (s EditControlGroupRequestAccountList) GoString() string {
  return s.String()
}

func (s *EditControlGroupRequestAccountList) SetLoginName(v string) *EditControlGroupRequestAccountList {
  s.LoginName = &v
  return s
}

type EditControlGroupRequestHeader struct {
}

func (s EditControlGroupRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s EditControlGroupRequestHeader) GoString() string {
  return s.String()
}

type EditControlGroupPaths struct {
  // {"en":"Control Group Code, Can be obtained through the API interface [QueryControlGroupList]","zh_CN":"Control Group 编号，可通过API接口 【查询ControlGroupList接口】 获取"}
  ControlGroupCode *string `json:"ControlGroupCode,omitempty" xml:"ControlGroupCode,omitempty" require:"true"`
}

func (s EditControlGroupPaths) String() string {
  return tea.Prettify(s)
}

func (s EditControlGroupPaths) GoString() string {
  return s.String()
}

func (s *EditControlGroupPaths) SetControlGroupCode(v string) *EditControlGroupPaths {
  s.ControlGroupCode = &v
  return s
}

type EditControlGroupParameters struct {
}

func (s EditControlGroupParameters) String() string {
  return tea.Prettify(s)
}

func (s EditControlGroupParameters) GoString() string {
  return s.String()
}

type EditControlGroupResponse struct {
  // {"en":"Message","zh_CN":"消息提示"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"Status Code","zh_CN":"错误具体状态码"}
  Code *int `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"","zh_CN":""}
  Data *EditControlGroupResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
  // {"en":"Request ID","zh_CN":"请求ID"}
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty" require:"true"`
}

func (s EditControlGroupResponse) String() string {
  return tea.Prettify(s)
}

func (s EditControlGroupResponse) GoString() string {
  return s.String()
}

func (s *EditControlGroupResponse) SetMsg(v string) *EditControlGroupResponse {
  s.Msg = &v
  return s
}

func (s *EditControlGroupResponse) SetCode(v int) *EditControlGroupResponse {
  s.Code = &v
  return s
}

func (s *EditControlGroupResponse) SetData(v *EditControlGroupResponseData) *EditControlGroupResponse {
  s.Data = v
  return s
}

func (s *EditControlGroupResponse) SetRequestId(v string) *EditControlGroupResponse {
  s.RequestId = &v
  return s
}

type EditControlGroupResponseData struct {
  // {"en":"Control Group Code","zh_CN":"Control Group Code"}
  ControlGroupCode *string `json:"controlGroupCode,omitempty" xml:"controlGroupCode,omitempty" require:"true"`
  // {"en":"Control Group Name","zh_CN":"Control Group名称"}
  ControlGroupName *string `json:"controlGroupName,omitempty" xml:"controlGroupName,omitempty" require:"true"`
}

func (s EditControlGroupResponseData) String() string {
  return tea.Prettify(s)
}

func (s EditControlGroupResponseData) GoString() string {
  return s.String()
}

func (s *EditControlGroupResponseData) SetControlGroupCode(v string) *EditControlGroupResponseData {
  s.ControlGroupCode = &v
  return s
}

func (s *EditControlGroupResponseData) SetControlGroupName(v string) *EditControlGroupResponseData {
  s.ControlGroupName = &v
  return s
}

type EditControlGroupResponseHeader struct {
}

func (s EditControlGroupResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s EditControlGroupResponseHeader) GoString() string {
  return s.String()
}




type BatchAssociateOrDetachControlGroupWithSubAccountRequest struct {
  // {"en":"Sub account login name", "zh_CN":"子用户登录名"}
  LoginName *string `json:"loginName,omitempty" xml:"loginName,omitempty" require:"true"`
  // {"en":"The list of control group code", "zh_CN":"Control Group编号列表"}
  ControlGroupCode []*string `json:"controlGroupCode,omitempty" xml:"controlGroupCode,omitempty" require:"true" type:"Repeated"`
  // {"en":"Select you want to add or revoke control group for sub account.
  // 
  // 0:add 
  // 
  // 1:revoke", "zh_CN":"选择需要为子用户添加或撤销关联CG
  // 
  // 0：添加关联
  // 
  // 1：撤销关联"}
  Type *int `json:"type,omitempty" xml:"type,omitempty" require:"true"`
}

func (s BatchAssociateOrDetachControlGroupWithSubAccountRequest) String() string {
  return tea.Prettify(s)
}

func (s BatchAssociateOrDetachControlGroupWithSubAccountRequest) GoString() string {
  return s.String()
}

func (s *BatchAssociateOrDetachControlGroupWithSubAccountRequest) SetLoginName(v string) *BatchAssociateOrDetachControlGroupWithSubAccountRequest {
  s.LoginName = &v
  return s
}

func (s *BatchAssociateOrDetachControlGroupWithSubAccountRequest) SetControlGroupCode(v []*string) *BatchAssociateOrDetachControlGroupWithSubAccountRequest {
  s.ControlGroupCode = v
  return s
}

func (s *BatchAssociateOrDetachControlGroupWithSubAccountRequest) SetType(v int) *BatchAssociateOrDetachControlGroupWithSubAccountRequest {
  s.Type = &v
  return s
}

type BatchAssociateOrDetachControlGroupWithSubAccountResponse struct {
  // {"en":"Request result status code", "zh_CN":"请求结果状态码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Request result information
  // ", "zh_CN":"请求结果信息"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
}

func (s BatchAssociateOrDetachControlGroupWithSubAccountResponse) String() string {
  return tea.Prettify(s)
}

func (s BatchAssociateOrDetachControlGroupWithSubAccountResponse) GoString() string {
  return s.String()
}

func (s *BatchAssociateOrDetachControlGroupWithSubAccountResponse) SetCode(v string) *BatchAssociateOrDetachControlGroupWithSubAccountResponse {
  s.Code = &v
  return s
}

func (s *BatchAssociateOrDetachControlGroupWithSubAccountResponse) SetMsg(v string) *BatchAssociateOrDetachControlGroupWithSubAccountResponse {
  s.Msg = &v
  return s
}

type BatchAssociateOrDetachControlGroupWithSubAccountPaths struct {
}

func (s BatchAssociateOrDetachControlGroupWithSubAccountPaths) String() string {
  return tea.Prettify(s)
}

func (s BatchAssociateOrDetachControlGroupWithSubAccountPaths) GoString() string {
  return s.String()
}

type BatchAssociateOrDetachControlGroupWithSubAccountParameters struct {
}

func (s BatchAssociateOrDetachControlGroupWithSubAccountParameters) String() string {
  return tea.Prettify(s)
}

func (s BatchAssociateOrDetachControlGroupWithSubAccountParameters) GoString() string {
  return s.String()
}

type BatchAssociateOrDetachControlGroupWithSubAccountRequestHeader struct {
}

func (s BatchAssociateOrDetachControlGroupWithSubAccountRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s BatchAssociateOrDetachControlGroupWithSubAccountRequestHeader) GoString() string {
  return s.String()
}

type BatchAssociateOrDetachControlGroupWithSubAccountResponseHeader struct {
}

func (s BatchAssociateOrDetachControlGroupWithSubAccountResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s BatchAssociateOrDetachControlGroupWithSubAccountResponseHeader) GoString() string {
  return s.String()
}




type QuerySubAccountRelatedControlGroupRequest struct {
}

func (s QuerySubAccountRelatedControlGroupRequest) String() string {
  return tea.Prettify(s)
}

func (s QuerySubAccountRelatedControlGroupRequest) GoString() string {
  return s.String()
}

type QuerySubAccountRelatedControlGroupRequestHeader struct {
}

func (s QuerySubAccountRelatedControlGroupRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QuerySubAccountRelatedControlGroupRequestHeader) GoString() string {
  return s.String()
}

type QuerySubAccountRelatedControlGroupPaths struct {
  // {"en":"LoginName of ordinary subAccount","zh_CN":"子用户登录名"}
  LoginName *string `json:"loginName,omitempty" xml:"loginName,omitempty" require:"true"`
}

func (s QuerySubAccountRelatedControlGroupPaths) String() string {
  return tea.Prettify(s)
}

func (s QuerySubAccountRelatedControlGroupPaths) GoString() string {
  return s.String()
}

func (s *QuerySubAccountRelatedControlGroupPaths) SetLoginName(v string) *QuerySubAccountRelatedControlGroupPaths {
  s.LoginName = &v
  return s
}

type QuerySubAccountRelatedControlGroupParameters struct {
}

func (s QuerySubAccountRelatedControlGroupParameters) String() string {
  return tea.Prettify(s)
}

func (s QuerySubAccountRelatedControlGroupParameters) GoString() string {
  return s.String()
}

type QuerySubAccountRelatedControlGroupResponse struct {
  // {"en":"Request result information","zh_CN":"请求结果信息"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"Request result status code","zh_CN":"请求结果状态码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"","zh_CN":""}
  Data []*QuerySubAccountRelatedControlGroupResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s QuerySubAccountRelatedControlGroupResponse) String() string {
  return tea.Prettify(s)
}

func (s QuerySubAccountRelatedControlGroupResponse) GoString() string {
  return s.String()
}

func (s *QuerySubAccountRelatedControlGroupResponse) SetMsg(v string) *QuerySubAccountRelatedControlGroupResponse {
  s.Msg = &v
  return s
}

func (s *QuerySubAccountRelatedControlGroupResponse) SetCode(v string) *QuerySubAccountRelatedControlGroupResponse {
  s.Code = &v
  return s
}

func (s *QuerySubAccountRelatedControlGroupResponse) SetData(v []*QuerySubAccountRelatedControlGroupResponseData) *QuerySubAccountRelatedControlGroupResponse {
  s.Data = v
  return s
}

type QuerySubAccountRelatedControlGroupResponseData struct     {
  // {"en":"Control Group Code","zh_CN":"Control Group编号"}
  ControlGroupCode *string `json:"controlGroupCode,omitempty" xml:"controlGroupCode,omitempty" require:"true"`
  // {"en":"Control Group Name","zh_CN":"Control Group名称"}
  ControlGroupName *string `json:"controlGroupName,omitempty" xml:"controlGroupName,omitempty" require:"true"`
}

func (s QuerySubAccountRelatedControlGroupResponseData) String() string {
  return tea.Prettify(s)
}

func (s QuerySubAccountRelatedControlGroupResponseData) GoString() string {
  return s.String()
}

func (s *QuerySubAccountRelatedControlGroupResponseData) SetControlGroupCode(v string) *QuerySubAccountRelatedControlGroupResponseData {
  s.ControlGroupCode = &v
  return s
}

func (s *QuerySubAccountRelatedControlGroupResponseData) SetControlGroupName(v string) *QuerySubAccountRelatedControlGroupResponseData {
  s.ControlGroupName = &v
  return s
}

type QuerySubAccountRelatedControlGroupResponseHeader struct {
}

func (s QuerySubAccountRelatedControlGroupResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QuerySubAccountRelatedControlGroupResponseHeader) GoString() string {
  return s.String()
}




type QueryControlGroupListRequest struct {
}

func (s QueryControlGroupListRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryControlGroupListRequest) GoString() string {
  return s.String()
}

type QueryControlGroupListResponse struct {
  // {"en":"Status Code 200:sueecss, 500:failed", "zh_CN":"错误具体状态码 # 200:sueecss， 500: failed"}
  Code *int `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Message", "zh_CN":"消息提示"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"Request ID","zh_CN":"请求ID"}
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty" require:"true"`
  Data []*QueryControlGroupListResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s QueryControlGroupListResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryControlGroupListResponse) GoString() string {
  return s.String()
}

func (s *QueryControlGroupListResponse) SetCode(v int) *QueryControlGroupListResponse {
  s.Code = &v
  return s
}

func (s *QueryControlGroupListResponse) SetMsg(v string) *QueryControlGroupListResponse {
  s.Msg = &v
  return s
}

func (s *QueryControlGroupListResponse) SetRequestId(v string) *QueryControlGroupListResponse {
  s.RequestId = &v
  return s
}

func (s *QueryControlGroupListResponse) SetData(v []*QueryControlGroupListResponseData) *QueryControlGroupListResponse {
  s.Data = v
  return s
}

type QueryControlGroupListResponseData struct     {
  // {"en":"Control Group Name", "zh_CN":"Control Group名称"}
  CONTROL_GROUP_NAME *string `json:"CONTROL_GROUP_NAME,omitempty" xml:"CONTROL_GROUP_NAME,omitempty" require:"true"`
  // {"en":"Control Group Code", "zh_CN":"Control Group编号"}
  CONTROL_GROUP_CODE *string `json:"CONTROL_GROUP_CODE,omitempty" xml:"CONTROL_GROUP_CODE,omitempty" require:"true"`
  // {"en":"Control Group ID", "zh_CN":"Control GroupID"}
  CONTROL_GROUP_ID *string `json:"CONTROL_GROUP_ID,omitempty" xml:"CONTROL_GROUP_ID,omitempty" require:"true"`
  // {"en":"Control Group Type", "zh_CN":"Control Group 类型"}
  CONTROL_GROUP_TYPE *string `json:"CONTROL_GROUP_TYPE,omitempty" xml:"CONTROL_GROUP_TYPE,omitempty" require:"true"`
}

func (s QueryControlGroupListResponseData) String() string {
  return tea.Prettify(s)
}

func (s QueryControlGroupListResponseData) GoString() string {
  return s.String()
}

func (s *QueryControlGroupListResponseData) SetCONTROL_GROUP_NAME(v string) *QueryControlGroupListResponseData {
  s.CONTROL_GROUP_NAME = &v
  return s
}

func (s *QueryControlGroupListResponseData) SetCONTROL_GROUP_CODE(v string) *QueryControlGroupListResponseData {
  s.CONTROL_GROUP_CODE = &v
  return s
}

func (s *QueryControlGroupListResponseData) SetCONTROL_GROUP_ID(v string) *QueryControlGroupListResponseData {
  s.CONTROL_GROUP_ID = &v
  return s
}

func (s *QueryControlGroupListResponseData) SetCONTROL_GROUP_TYPE(v string) *QueryControlGroupListResponseData {
  s.CONTROL_GROUP_TYPE = &v
  return s
}

type QueryControlGroupListPaths struct {
}

func (s QueryControlGroupListPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryControlGroupListPaths) GoString() string {
  return s.String()
}

type QueryControlGroupListParameters struct {
}

func (s QueryControlGroupListParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryControlGroupListParameters) GoString() string {
  return s.String()
}

type QueryControlGroupListRequestHeader struct {
}

func (s QueryControlGroupListRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryControlGroupListRequestHeader) GoString() string {
  return s.String()
}

type QueryControlGroupListResponseHeader struct {
}

func (s QueryControlGroupListResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryControlGroupListResponseHeader) GoString() string {
  return s.String()
}




type EditControlGroupByCoverRequest struct {
  // {"en":"Control Group Code","zh_CN":"Control Group 编码，可通过API接口 【查询ControlGroupList接口】 获取"}
  ControlGroupCode *string `json:"controlGroupCode,omitempty" xml:"controlGroupCode,omitempty" require:"true"`
  // {"en":"Domain array, which only the User Customized type Control Group can be modified, customer type Control Group and product type Control Group can not be modified.User Customized type Control Group empties the original domainList if no value is passed","zh_CN":"域名字符串数组，只有自定义类型的Control Group可做修改，若是客户类型与合同类型Control Group则不做修改。自定义类型Control Group若不传值则将原domainList清空"}
  DomainList []*string `json:"domainList,omitempty" xml:"domainList,omitempty" type:"Repeated"`
  // {"en":"Account object array,Used to specify accounts with permission.  all types of Control Group can be modified, if no value is passed, the original accountList will be emptied","zh_CN":"账号对象数组, 用来指定有权限访问的账号。客户类型，合同类型与自定义类型的Control Group都可以做修改，若不传值则将原accountList清空"}
  AccountList []*EditControlGroupByCoverRequestAccountList `json:"accountList,omitempty" xml:"accountList,omitempty" type:"Repeated"`
  // {"en":"Control Group name, which only the User Customized type Control Group can be modified, customer type Control Group and product type Control Group can not be modified. User Customized type Control Group keeps the original Control Group name if no value is passed","zh_CN":"Control Group名称，只有自定义类型的Control Group可做修改，若是客户类型与合同类型Control Group则不做修改。自定义类型Control Group若不传值则保持原来的Control Group名称"}
  ControlGroupName *string `json:"controlGroupName,omitempty" xml:"controlGroupName,omitempty"`
}

func (s EditControlGroupByCoverRequest) String() string {
  return tea.Prettify(s)
}

func (s EditControlGroupByCoverRequest) GoString() string {
  return s.String()
}

func (s *EditControlGroupByCoverRequest) SetControlGroupCode(v string) *EditControlGroupByCoverRequest {
  s.ControlGroupCode = &v
  return s
}

func (s *EditControlGroupByCoverRequest) SetDomainList(v []*string) *EditControlGroupByCoverRequest {
  s.DomainList = v
  return s
}

func (s *EditControlGroupByCoverRequest) SetAccountList(v []*EditControlGroupByCoverRequestAccountList) *EditControlGroupByCoverRequest {
  s.AccountList = v
  return s
}

func (s *EditControlGroupByCoverRequest) SetControlGroupName(v string) *EditControlGroupByCoverRequest {
  s.ControlGroupName = &v
  return s
}

type EditControlGroupByCoverRequestAccountList struct     {
  // {"en":"Account","zh_CN":"账号"}
  LoginName *string `json:"loginName,omitempty" xml:"loginName,omitempty" require:"true"`
}

func (s EditControlGroupByCoverRequestAccountList) String() string {
  return tea.Prettify(s)
}

func (s EditControlGroupByCoverRequestAccountList) GoString() string {
  return s.String()
}

func (s *EditControlGroupByCoverRequestAccountList) SetLoginName(v string) *EditControlGroupByCoverRequestAccountList {
  s.LoginName = &v
  return s
}

type EditControlGroupByCoverRequestHeader struct {
}

func (s EditControlGroupByCoverRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s EditControlGroupByCoverRequestHeader) GoString() string {
  return s.String()
}

type EditControlGroupByCoverPaths struct {
}

func (s EditControlGroupByCoverPaths) String() string {
  return tea.Prettify(s)
}

func (s EditControlGroupByCoverPaths) GoString() string {
  return s.String()
}

type EditControlGroupByCoverParameters struct {
}

func (s EditControlGroupByCoverParameters) String() string {
  return tea.Prettify(s)
}

func (s EditControlGroupByCoverParameters) GoString() string {
  return s.String()
}

type EditControlGroupByCoverResponse struct {
  // {"en":"Message","zh_CN":"消息提示"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"Status Code","zh_CN":"错误具体状态码"}
  Code *int `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"","zh_CN":""}
  Data *EditControlGroupByCoverResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
  // {"en":"Request ID","zh_CN":"请求ID"}
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty" require:"true"`
}

func (s EditControlGroupByCoverResponse) String() string {
  return tea.Prettify(s)
}

func (s EditControlGroupByCoverResponse) GoString() string {
  return s.String()
}

func (s *EditControlGroupByCoverResponse) SetMsg(v string) *EditControlGroupByCoverResponse {
  s.Msg = &v
  return s
}

func (s *EditControlGroupByCoverResponse) SetCode(v int) *EditControlGroupByCoverResponse {
  s.Code = &v
  return s
}

func (s *EditControlGroupByCoverResponse) SetData(v *EditControlGroupByCoverResponseData) *EditControlGroupByCoverResponse {
  s.Data = v
  return s
}

func (s *EditControlGroupByCoverResponse) SetRequestId(v string) *EditControlGroupByCoverResponse {
  s.RequestId = &v
  return s
}

type EditControlGroupByCoverResponseData struct {
  // {"en":"Control Group Code","zh_CN":"Control Group Code"}
  ControlGroupCode *string `json:"controlGroupCode,omitempty" xml:"controlGroupCode,omitempty" require:"true"`
  // {"en":"Control Group Name","zh_CN":"Control Group名称"}
  ControlGroupName *string `json:"controlGroupName,omitempty" xml:"controlGroupName,omitempty" require:"true"`
}

func (s EditControlGroupByCoverResponseData) String() string {
  return tea.Prettify(s)
}

func (s EditControlGroupByCoverResponseData) GoString() string {
  return s.String()
}

func (s *EditControlGroupByCoverResponseData) SetControlGroupCode(v string) *EditControlGroupByCoverResponseData {
  s.ControlGroupCode = &v
  return s
}

func (s *EditControlGroupByCoverResponseData) SetControlGroupName(v string) *EditControlGroupByCoverResponseData {
  s.ControlGroupName = &v
  return s
}

type EditControlGroupByCoverResponseHeader struct {
}

func (s EditControlGroupByCoverResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s EditControlGroupByCoverResponseHeader) GoString() string {
  return s.String()
}




type DeleteControlGroupRequest struct {
}

func (s DeleteControlGroupRequest) String() string {
  return tea.Prettify(s)
}

func (s DeleteControlGroupRequest) GoString() string {
  return s.String()
}

type DeleteControlGroupResponse struct {
  // {"en":"Status Code", "zh_CN":"错误具体状态码"}
  Code *int `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Message", "zh_CN":"消息提示"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"Success Mark", "zh_CN":"成功标记"}
  Success *bool `json:"success,omitempty" xml:"success,omitempty" require:"true"`
}

func (s DeleteControlGroupResponse) String() string {
  return tea.Prettify(s)
}

func (s DeleteControlGroupResponse) GoString() string {
  return s.String()
}

func (s *DeleteControlGroupResponse) SetCode(v int) *DeleteControlGroupResponse {
  s.Code = &v
  return s
}

func (s *DeleteControlGroupResponse) SetMsg(v string) *DeleteControlGroupResponse {
  s.Msg = &v
  return s
}

func (s *DeleteControlGroupResponse) SetSuccess(v bool) *DeleteControlGroupResponse {
  s.Success = &v
  return s
}

type DeleteControlGroupPaths struct {
  // {"en":"Control Group Code", "zh_CN":"Control Group 编号"}
  ControlGroupCode *string `json:"ControlGroupCode,omitempty" xml:"ControlGroupCode,omitempty" require:"true"`
}

func (s DeleteControlGroupPaths) String() string {
  return tea.Prettify(s)
}

func (s DeleteControlGroupPaths) GoString() string {
  return s.String()
}

func (s *DeleteControlGroupPaths) SetControlGroupCode(v string) *DeleteControlGroupPaths {
  s.ControlGroupCode = &v
  return s
}

type DeleteControlGroupParameters struct {
}

func (s DeleteControlGroupParameters) String() string {
  return tea.Prettify(s)
}

func (s DeleteControlGroupParameters) GoString() string {
  return s.String()
}

type DeleteControlGroupRequestHeader struct {
}

func (s DeleteControlGroupRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteControlGroupRequestHeader) GoString() string {
  return s.String()
}

type DeleteControlGroupResponseHeader struct {
}

func (s DeleteControlGroupResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteControlGroupResponseHeader) GoString() string {
  return s.String()
}




