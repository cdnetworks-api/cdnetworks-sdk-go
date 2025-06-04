package cgmanage

// This file is auto-generated, don't edit it. Thanks.
import (
	"github.com/alibabacloud-go/tea/tea"
)

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

type CreateControlGroupRequestAccountList struct {
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
	// {"en":"Request ID","zh_CN":"策略ID"}
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
	// {"en":"Control Group Code, If this parameter is empty, all authorized control groups will be returned", "zh_CN":"服务组编号，不传则返回全部有权限的服务组"}
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

type GetDomainListOfControlGroupResponse struct {
	GetDomainListOfControlGroupControlGroupDetail []*GetDomainListOfControlGroupControlGroupDetail `json:"controlGroupDetail,omitempty" xml:"controlGroupDetail,omitempty" require:"true" type:"Repeated"`
}

func (s GetDomainListOfControlGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDomainListOfControlGroupResponse) GoString() string {
	return s.String()
}

func (s *GetDomainListOfControlGroupResponse) SetControlGroupDetail(v []*GetDomainListOfControlGroupControlGroupDetail) *GetDomainListOfControlGroupResponse {
	s.GetDomainListOfControlGroupControlGroupDetail = v
	return s
}

type GetDomainListOfControlGroupControlGroupDetail struct {
	// {"en":"Control Group Code", "zh_CN":"服务组编号"}
	ControlGroupCode *string `json:"controlGroupCode,omitempty" xml:"controlGroupCode,omitempty" require:"true"`
	// {"en":"Control Group Name", "zh_CN":"服务组名称"}
	ControlGroupName *string `json:"controlGroupName,omitempty" xml:"controlGroupName,omitempty" require:"true"`
	// {"en":"domain list", "zh_CN":"域名列表"}
	DomainList []*string `json:"domainList,omitempty" xml:"domainList,omitempty" require:"true" type:"Repeated"`
}

func (s GetDomainListOfControlGroupControlGroupDetail) String() string {
	return tea.Prettify(s)
}

func (s GetDomainListOfControlGroupControlGroupDetail) GoString() string {
	return s.String()
}

func (s *GetDomainListOfControlGroupControlGroupDetail) SetControlGroupCode(v string) *GetDomainListOfControlGroupControlGroupDetail {
	s.ControlGroupCode = &v
	return s
}

func (s *GetDomainListOfControlGroupControlGroupDetail) SetControlGroupName(v string) *GetDomainListOfControlGroupControlGroupDetail {
	s.ControlGroupName = &v
	return s
}

func (s *GetDomainListOfControlGroupControlGroupDetail) SetDomainList(v []*string) *GetDomainListOfControlGroupControlGroupDetail {
	s.DomainList = v
	return s
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

type GetDomainListOfControlGroupRequestHeader struct {
}

func (s GetDomainListOfControlGroupRequestHeader) String() string {
	return tea.Prettify(s)
}

func (s GetDomainListOfControlGroupRequestHeader) GoString() string {
	return s.String()
}

type GetDomainListOfControlGroupResponseHeader struct {
}

func (s GetDomainListOfControlGroupResponseHeader) String() string {
	return tea.Prettify(s)
}

func (s GetDomainListOfControlGroupResponseHeader) GoString() string {
	return s.String()
}

type EditControlGroupRequest struct {
	// {"en":"Control Group name, which only the User Customized type Control Group can be modified, customer type Control Group and product type Control Group can not be modified. User Customized type Control Group keeps the original Control Group name if no value is passed", "zh_CN":"Control Group名称，只有自定义类型的Control Group可做修改，若是客户类型与合同类型Control Group则不做修改。自定义类型Control Group若不传值则保持原来的Control Group名称"}
	ControlGroupName *string `json:"controlGroupName,omitempty" xml:"controlGroupName,omitempty"`
	// {"en":"Account object array,Used to specify accounts with permission.  all types of Control Group can be modified, if no value is passed, the original accountList will be emptied", "zh_CN":"账号对象数组, 用来指定有权限访问的账号。客户类型，合同类型与自定义类型的Control Group都可以做修改，若不传值则将原accountList清空"}
	AccountList []*EditControlGroupLoginName `json:"accountList,omitempty" xml:"accountList,omitempty" type:"Repeated"`
	// {"en":"Domain array, which only the User Customized type Control Group can be modified, customer type Control Group and product type Control Group can not be modified.User Customized type Control Group empties the original domainList if no value is passed", "zh_CN":"域名字符串数组，只有自定义类型的Control Group可做修改，若是客户类型与合同类型Control Group则不做修改。自定义类型Control Group若不传值则将原domainList清空"}
	DomainList []*string `json:"domainList,omitempty" xml:"domainList,omitempty" type:"Repeated"`
	// {"en":"Whether to add:
	// 1. Do not pass or pass false: rewrite method;
	// 2. Pass true: append method.", "zh_CN":"是否追加:
	// 1.不传或false：覆盖方式;
	// 2.传true：追加方式."}
	IsAdd *bool `json:"isAdd,omitempty" xml:"isAdd,omitempty"`
}

func (s EditControlGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s EditControlGroupRequest) GoString() string {
	return s.String()
}

func (s *EditControlGroupRequest) SetControlGroupName(v string) *EditControlGroupRequest {
	s.ControlGroupName = &v
	return s
}

func (s *EditControlGroupRequest) SetAccountList(v []*EditControlGroupLoginName) *EditControlGroupRequest {
	s.AccountList = v
	return s
}

func (s *EditControlGroupRequest) SetDomainList(v []*string) *EditControlGroupRequest {
	s.DomainList = v
	return s
}

func (s *EditControlGroupRequest) SetIsAdd(v bool) *EditControlGroupRequest {
	s.IsAdd = &v
	return s
}

type EditControlGroupLoginName struct {
	// {"en":"Account", "zh_CN":"账号"}
	EditControlGroupLoginName *string `json:"loginName,omitempty" xml:"loginName,omitempty"`
}

func (s EditControlGroupLoginName) String() string {
	return tea.Prettify(s)
}

func (s EditControlGroupLoginName) GoString() string {
	return s.String()
}

func (s *EditControlGroupLoginName) SetEditControlGroupLoginName(v string) *EditControlGroupLoginName {
	s.EditControlGroupLoginName = &v
	return s
}

type EditControlGroupResponse struct {
	// {"en":"Status Code", "zh_CN":"错误具体状态码"}
	Code *int `json:"code,omitempty" xml:"code,omitempty" require:"true"`
	// {"en":"Message", "zh_CN":"消息提示"}
	Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
	// {"en":"Request ID","zh_CN":"请求ID"}
	RequestId *string                       `json:"requestId,omitempty" xml:"requestId,omitempty" require:"true"`
	Data      *EditControlGroupResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
}

func (s EditControlGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s EditControlGroupResponse) GoString() string {
	return s.String()
}

func (s *EditControlGroupResponse) SetCode(v int) *EditControlGroupResponse {
	s.Code = &v
	return s
}

func (s *EditControlGroupResponse) SetMsg(v string) *EditControlGroupResponse {
	s.Msg = &v
	return s
}

func (s *EditControlGroupResponse) SetRequestId(v string) *EditControlGroupResponse {
	s.RequestId = &v
	return s
}

func (s *EditControlGroupResponse) SetData(v *EditControlGroupResponseData) *EditControlGroupResponse {
	s.Data = v
	return s
}

type EditControlGroupResponseData struct {
	// {"en":"Control Group Code", "zh_CN":"Control Group Code"}
	ControlGroupCode *string `json:"controlGroupCode,omitempty" xml:"controlGroupCode,omitempty" require:"true"`
	// {"en":"Control Group Name", "zh_CN":"Control Group名称"}
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

type EditControlGroupPaths struct {
	// {"en":"Control Group Code, Can be obtained through the API interface [QueryControlGroupList]", "zh_CN":"Control Group 编号，可通过API接口 【查询ControlGroupList接口】 获取"}
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

type EditControlGroupRequestHeader struct {
}

func (s EditControlGroupRequestHeader) String() string {
	return tea.Prettify(s)
}

func (s EditControlGroupRequestHeader) GoString() string {
	return s.String()
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
	Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
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

func (s *BatchAssociateOrDetachControlGroupWithSubAccountResponse) SetMessage(v string) *BatchAssociateOrDetachControlGroupWithSubAccountResponse {
	s.Message = &v
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

type ApiIamQuerysubaccountrelatedcontrolgroupDnaRequest struct {
}

func (s ApiIamQuerysubaccountrelatedcontrolgroupDnaRequest) String() string {
	return tea.Prettify(s)
}

func (s ApiIamQuerysubaccountrelatedcontrolgroupDnaRequest) GoString() string {
	return s.String()
}

type ApiIamQuerysubaccountrelatedcontrolgroupDnaResponse struct {
	// {"en":"Request result status code", "zh_CN":"请求结果状态码"}
	Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
	// {"en":"Request result information", "zh_CN":"请求结果信息"}
	Message *string                                                        `json:"message,omitempty" xml:"message,omitempty" require:"true"`
	Data    []*ApiIamQuerysubaccountrelatedcontrolgroupDnaControlGroupInfo `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s ApiIamQuerysubaccountrelatedcontrolgroupDnaResponse) String() string {
	return tea.Prettify(s)
}

func (s ApiIamQuerysubaccountrelatedcontrolgroupDnaResponse) GoString() string {
	return s.String()
}

func (s *ApiIamQuerysubaccountrelatedcontrolgroupDnaResponse) SetCode(v string) *ApiIamQuerysubaccountrelatedcontrolgroupDnaResponse {
	s.Code = &v
	return s
}

func (s *ApiIamQuerysubaccountrelatedcontrolgroupDnaResponse) SetMessage(v string) *ApiIamQuerysubaccountrelatedcontrolgroupDnaResponse {
	s.Message = &v
	return s
}

func (s *ApiIamQuerysubaccountrelatedcontrolgroupDnaResponse) SetData(v []*ApiIamQuerysubaccountrelatedcontrolgroupDnaControlGroupInfo) *ApiIamQuerysubaccountrelatedcontrolgroupDnaResponse {
	s.Data = v
	return s
}

type ApiIamQuerysubaccountrelatedcontrolgroupDnaControlGroupInfo struct {
	// {"en":"Control Group Name", "zh_CN":"Control Group名称"}
	ControlGroupName *string `json:"controlGroupName,omitempty" xml:"controlGroupName,omitempty" require:"true"`
	// {"en":"Control Group Code", "zh_CN":"Control Group编号"}
	ControlGroupCode *string `json:"controlGroupCode,omitempty" xml:"controlGroupCode,omitempty" require:"true"`
}

func (s ApiIamQuerysubaccountrelatedcontrolgroupDnaControlGroupInfo) String() string {
	return tea.Prettify(s)
}

func (s ApiIamQuerysubaccountrelatedcontrolgroupDnaControlGroupInfo) GoString() string {
	return s.String()
}

func (s *ApiIamQuerysubaccountrelatedcontrolgroupDnaControlGroupInfo) SetControlGroupName(v string) *ApiIamQuerysubaccountrelatedcontrolgroupDnaControlGroupInfo {
	s.ControlGroupName = &v
	return s
}

func (s *ApiIamQuerysubaccountrelatedcontrolgroupDnaControlGroupInfo) SetControlGroupCode(v string) *ApiIamQuerysubaccountrelatedcontrolgroupDnaControlGroupInfo {
	s.ControlGroupCode = &v
	return s
}

type ApiIamQuerysubaccountrelatedcontrolgroupDnaPaths struct {
	// {"en":"LoginName of ordinary subAccount", "zh_CN":"子用户登录名"}
	LoginName *string `json:"loginName,omitempty" xml:"loginName,omitempty" require:"true"`
}

func (s ApiIamQuerysubaccountrelatedcontrolgroupDnaPaths) String() string {
	return tea.Prettify(s)
}

func (s ApiIamQuerysubaccountrelatedcontrolgroupDnaPaths) GoString() string {
	return s.String()
}

func (s *ApiIamQuerysubaccountrelatedcontrolgroupDnaPaths) SetLoginName(v string) *ApiIamQuerysubaccountrelatedcontrolgroupDnaPaths {
	s.LoginName = &v
	return s
}

type ApiIamQuerysubaccountrelatedcontrolgroupDnaParameters struct {
}

func (s ApiIamQuerysubaccountrelatedcontrolgroupDnaParameters) String() string {
	return tea.Prettify(s)
}

func (s ApiIamQuerysubaccountrelatedcontrolgroupDnaParameters) GoString() string {
	return s.String()
}

type ApiIamQuerysubaccountrelatedcontrolgroupDnaRequestHeader struct {
}

func (s ApiIamQuerysubaccountrelatedcontrolgroupDnaRequestHeader) String() string {
	return tea.Prettify(s)
}

func (s ApiIamQuerysubaccountrelatedcontrolgroupDnaRequestHeader) GoString() string {
	return s.String()
}

type ApiIamQuerysubaccountrelatedcontrolgroupDnaResponseHeader struct {
}

func (s ApiIamQuerysubaccountrelatedcontrolgroupDnaResponseHeader) String() string {
	return tea.Prettify(s)
}

func (s ApiIamQuerysubaccountrelatedcontrolgroupDnaResponseHeader) GoString() string {
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

type QueryControlGroupListRequestHeader struct {
}

func (s QueryControlGroupListRequestHeader) String() string {
	return tea.Prettify(s)
}

func (s QueryControlGroupListRequestHeader) GoString() string {
	return s.String()
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

type QueryControlGroupListResponse struct {
	// {"en":"Message","zh_CN":"消息提示"}
	Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
	// {"en":"Status Code 200:sueecss, 500:failed","zh_CN":"错误具体状态码 # 200:sueecss， 500: failed"}
	Code *int `json:"code,omitempty" xml:"code,omitempty" require:"true"`
	// {"en":"","zh_CN":""}
	Data []*QueryControlGroupListResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
	// {"en":"Request ID","zh_CN":"请求ID"}
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty" require:"true"`
}

func (s QueryControlGroupListResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryControlGroupListResponse) GoString() string {
	return s.String()
}

func (s *QueryControlGroupListResponse) SetMsg(v string) *QueryControlGroupListResponse {
	s.Msg = &v
	return s
}

func (s *QueryControlGroupListResponse) SetCode(v int) *QueryControlGroupListResponse {
	s.Code = &v
	return s
}

func (s *QueryControlGroupListResponse) SetData(v []*QueryControlGroupListResponseData) *QueryControlGroupListResponse {
	s.Data = v
	return s
}

func (s *QueryControlGroupListResponse) SetRequestId(v string) *QueryControlGroupListResponse {
	s.RequestId = &v
	return s
}

type QueryControlGroupListResponseData struct {
	// {"en":"Control Group Type","zh_CN":"Control Group 类型"}
	CONTROLGROUPTYPE *string `json:"CONTROL_GROUP_TYPE,omitempty" xml:"CONTROL_GROUP_TYPE,omitempty" require:"true"`
	// {"en":"Control Group Name","zh_CN":"Control Group名称"}
	CONTROLGROUPNAME *string `json:"CONTROL_GROUP_NAME,omitempty" xml:"CONTROL_GROUP_NAME,omitempty" require:"true"`
	// {"en":"Control Group Code","zh_CN":"Control Group编号"}
	CONTROLGROUPCODE *string `json:"CONTROL_GROUP_CODE,omitempty" xml:"CONTROL_GROUP_CODE,omitempty" require:"true"`
	// {"en":"Control Group ID","zh_CN":"Control GroupID"}
	CONTROLGROUPID *string `json:"CONTROL_GROUP_ID,omitempty" xml:"CONTROL_GROUP_ID,omitempty" require:"true"`
}

func (s QueryControlGroupListResponseData) String() string {
	return tea.Prettify(s)
}

func (s QueryControlGroupListResponseData) GoString() string {
	return s.String()
}

func (s *QueryControlGroupListResponseData) SetCONTROLGROUPTYPE(v string) *QueryControlGroupListResponseData {
	s.CONTROLGROUPTYPE = &v
	return s
}

func (s *QueryControlGroupListResponseData) SetCONTROLGROUPNAME(v string) *QueryControlGroupListResponseData {
	s.CONTROLGROUPNAME = &v
	return s
}

func (s *QueryControlGroupListResponseData) SetCONTROLGROUPCODE(v string) *QueryControlGroupListResponseData {
	s.CONTROLGROUPCODE = &v
	return s
}

func (s *QueryControlGroupListResponseData) SetCONTROLGROUPID(v string) *QueryControlGroupListResponseData {
	s.CONTROLGROUPID = &v
	return s
}

type QueryControlGroupListResponseHeader struct {
}

func (s QueryControlGroupListResponseHeader) String() string {
	return tea.Prettify(s)
}

func (s QueryControlGroupListResponseHeader) GoString() string {
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

type EditControlGroupByCoverRequestAccountList struct {
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
	// {"en":"Request ID","zh_CN":"策略ID"}
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
