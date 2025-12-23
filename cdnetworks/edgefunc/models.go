package edgefunc

// This file is auto-generated, don't edit it. Thanks.
import (
  "github.com/alibabacloud-go/tea/tea"
)

type DeleteFuncDomainTriggerRequest struct {
}

func (s DeleteFuncDomainTriggerRequest) String() string {
  return tea.Prettify(s)
}

func (s DeleteFuncDomainTriggerRequest) GoString() string {
  return s.String()
}

type DeleteFuncDomainTriggerRequestHeader struct {
}

func (s DeleteFuncDomainTriggerRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteFuncDomainTriggerRequestHeader) GoString() string {
  return s.String()
}

type DeleteFuncDomainTriggerPaths struct {
  // {"en":"Domain ID","zh_CN":"域名的 ID"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
}

func (s DeleteFuncDomainTriggerPaths) String() string {
  return tea.Prettify(s)
}

func (s DeleteFuncDomainTriggerPaths) GoString() string {
  return s.String()
}

func (s *DeleteFuncDomainTriggerPaths) SetId(v string) *DeleteFuncDomainTriggerPaths {
  s.Id = &v
  return s
}

type DeleteFuncDomainTriggerParameters struct {
}

func (s DeleteFuncDomainTriggerParameters) String() string {
  return tea.Prettify(s)
}

func (s DeleteFuncDomainTriggerParameters) GoString() string {
  return s.String()
}

type DeleteFuncDomainTriggerResponse struct {
  // {"en":"Status code","zh_CN":"响应码"}
  Code *int `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Status description","zh_CN":"响应信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s DeleteFuncDomainTriggerResponse) String() string {
  return tea.Prettify(s)
}

func (s DeleteFuncDomainTriggerResponse) GoString() string {
  return s.String()
}

func (s *DeleteFuncDomainTriggerResponse) SetCode(v int) *DeleteFuncDomainTriggerResponse {
  s.Code = &v
  return s
}

func (s *DeleteFuncDomainTriggerResponse) SetMessage(v string) *DeleteFuncDomainTriggerResponse {
  s.Message = &v
  return s
}

type DeleteFuncDomainTriggerResponseHeader struct {
}

func (s DeleteFuncDomainTriggerResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteFuncDomainTriggerResponseHeader) GoString() string {
  return s.String()
}




type QueryEdgeFuncTriggerRequest struct {
}

func (s QueryEdgeFuncTriggerRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryEdgeFuncTriggerRequest) GoString() string {
  return s.String()
}

type QueryEdgeFuncTriggerRequestHeader struct {
}

func (s QueryEdgeFuncTriggerRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryEdgeFuncTriggerRequestHeader) GoString() string {
  return s.String()
}

type QueryEdgeFuncTriggerPaths struct {
}

func (s QueryEdgeFuncTriggerPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryEdgeFuncTriggerPaths) GoString() string {
  return s.String()
}

type QueryEdgeFuncTriggerParameters struct {
  // {"en":"Domain, Either domain or function name is required.","zh_CN":"域名，与函数任填其一"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty"`
  // {"en":"Function name, Either domain or function name is required.","zh_CN":"函数名称，与域名任填其一"}
  FuncName *string `json:"funcName,omitempty" xml:"funcName,omitempty"`
  // {"en":"The page number for the query, starting from 1.","zh_CN":"查询的页码，从1开始"}
  PageNo *int `json:"pageNo,omitempty" xml:"pageNo,omitempty"`
  // {"en":"The number of items to return per page.","zh_CN":"每页返回的条目数量"}
  PageSize *int `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s QueryEdgeFuncTriggerParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryEdgeFuncTriggerParameters) GoString() string {
  return s.String()
}

func (s *QueryEdgeFuncTriggerParameters) SetDomain(v string) *QueryEdgeFuncTriggerParameters {
  s.Domain = &v
  return s
}

func (s *QueryEdgeFuncTriggerParameters) SetFuncName(v string) *QueryEdgeFuncTriggerParameters {
  s.FuncName = &v
  return s
}

func (s *QueryEdgeFuncTriggerParameters) SetPageNo(v int) *QueryEdgeFuncTriggerParameters {
  s.PageNo = &v
  return s
}

func (s *QueryEdgeFuncTriggerParameters) SetPageSize(v int) *QueryEdgeFuncTriggerParameters {
  s.PageSize = &v
  return s
}

type QueryEdgeFuncTriggerResponse struct {
  // {"en":"Status code","zh_CN":"响应码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Status description","zh_CN":"响应信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Response Data","zh_CN":"响应数据"}
  Data *QueryEdgeFuncTriggerResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
}

func (s QueryEdgeFuncTriggerResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryEdgeFuncTriggerResponse) GoString() string {
  return s.String()
}

func (s *QueryEdgeFuncTriggerResponse) SetCode(v string) *QueryEdgeFuncTriggerResponse {
  s.Code = &v
  return s
}

func (s *QueryEdgeFuncTriggerResponse) SetMessage(v string) *QueryEdgeFuncTriggerResponse {
  s.Message = &v
  return s
}

func (s *QueryEdgeFuncTriggerResponse) SetData(v *QueryEdgeFuncTriggerResponseData) *QueryEdgeFuncTriggerResponse {
  s.Data = v
  return s
}

type QueryEdgeFuncTriggerResponseData struct {
  // {"en":"Trigger List","zh_CN":"触发器列表"}
  Result []*QueryEdgeFuncTriggerResponseDataResult `json:"result,omitempty" xml:"result,omitempty" require:"true" type:"Repeated"`
}

func (s QueryEdgeFuncTriggerResponseData) String() string {
  return tea.Prettify(s)
}

func (s QueryEdgeFuncTriggerResponseData) GoString() string {
  return s.String()
}

func (s *QueryEdgeFuncTriggerResponseData) SetResult(v []*QueryEdgeFuncTriggerResponseDataResult) *QueryEdgeFuncTriggerResponseData {
  s.Result = v
  return s
}

type QueryEdgeFuncTriggerResponseDataResult struct     {
  // {"en":"Domain ID","zh_CN":"域名 iD"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {"en":"Domain","zh_CN":"域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {"en":"Router List","zh_CN":"路由列表"}
  Routers *QueryEdgeFuncTriggerResponseDataResultRouters `json:"routers,omitempty" xml:"routers,omitempty" require:"true" type:"Struct"`
}

func (s QueryEdgeFuncTriggerResponseDataResult) String() string {
  return tea.Prettify(s)
}

func (s QueryEdgeFuncTriggerResponseDataResult) GoString() string {
  return s.String()
}

func (s *QueryEdgeFuncTriggerResponseDataResult) SetId(v string) *QueryEdgeFuncTriggerResponseDataResult {
  s.Id = &v
  return s
}

func (s *QueryEdgeFuncTriggerResponseDataResult) SetDomain(v string) *QueryEdgeFuncTriggerResponseDataResult {
  s.Domain = &v
  return s
}

func (s *QueryEdgeFuncTriggerResponseDataResult) SetRouters(v *QueryEdgeFuncTriggerResponseDataResultRouters) *QueryEdgeFuncTriggerResponseDataResult {
  s.Routers = v
  return s
}

type QueryEdgeFuncTriggerResponseDataResultRouters struct {
  // {"en":"Func name","zh_CN":"函数名称"}
  FuncName *string `json:"funcName,omitempty" xml:"funcName,omitempty" require:"true"`
  // {"en":"Match path","zh_CN":"匹配路径"}
  Path *string `json:"path,omitempty" xml:"path,omitempty" require:"true"`
}

func (s QueryEdgeFuncTriggerResponseDataResultRouters) String() string {
  return tea.Prettify(s)
}

func (s QueryEdgeFuncTriggerResponseDataResultRouters) GoString() string {
  return s.String()
}

func (s *QueryEdgeFuncTriggerResponseDataResultRouters) SetFuncName(v string) *QueryEdgeFuncTriggerResponseDataResultRouters {
  s.FuncName = &v
  return s
}

func (s *QueryEdgeFuncTriggerResponseDataResultRouters) SetPath(v string) *QueryEdgeFuncTriggerResponseDataResultRouters {
  s.Path = &v
  return s
}

type QueryEdgeFuncTriggerResponseHeader struct {
}

func (s QueryEdgeFuncTriggerResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryEdgeFuncTriggerResponseHeader) GoString() string {
  return s.String()
}




type CreateEdgeFuncTriggerRequest struct {
  // {"en":"Domain","zh_CN":"域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {"en":"Router list","zh_CN":"路由列表"}
  Routers []*CreateEdgeFuncTriggerRequestRouters `json:"routers,omitempty" xml:"routers,omitempty" require:"true" type:"Repeated"`
}

func (s CreateEdgeFuncTriggerRequest) String() string {
  return tea.Prettify(s)
}

func (s CreateEdgeFuncTriggerRequest) GoString() string {
  return s.String()
}

func (s *CreateEdgeFuncTriggerRequest) SetDomain(v string) *CreateEdgeFuncTriggerRequest {
  s.Domain = &v
  return s
}

func (s *CreateEdgeFuncTriggerRequest) SetRouters(v []*CreateEdgeFuncTriggerRequestRouters) *CreateEdgeFuncTriggerRequest {
  s.Routers = v
  return s
}

type CreateEdgeFuncTriggerRequestRouters struct     {
  // {"en":"Func name","zh_CN":"函数名称"}
  FuncName *string `json:"funcName,omitempty" xml:"funcName,omitempty" require:"true"`
  // {"en":"Match path","zh_CN":"匹配路径"}
  Path *string `json:"path,omitempty" xml:"path,omitempty"`
}

func (s CreateEdgeFuncTriggerRequestRouters) String() string {
  return tea.Prettify(s)
}

func (s CreateEdgeFuncTriggerRequestRouters) GoString() string {
  return s.String()
}

func (s *CreateEdgeFuncTriggerRequestRouters) SetFuncName(v string) *CreateEdgeFuncTriggerRequestRouters {
  s.FuncName = &v
  return s
}

func (s *CreateEdgeFuncTriggerRequestRouters) SetPath(v string) *CreateEdgeFuncTriggerRequestRouters {
  s.Path = &v
  return s
}

type CreateEdgeFuncTriggerRequestHeader struct {
}

func (s CreateEdgeFuncTriggerRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s CreateEdgeFuncTriggerRequestHeader) GoString() string {
  return s.String()
}

type CreateEdgeFuncTriggerPaths struct {
}

func (s CreateEdgeFuncTriggerPaths) String() string {
  return tea.Prettify(s)
}

func (s CreateEdgeFuncTriggerPaths) GoString() string {
  return s.String()
}

type CreateEdgeFuncTriggerParameters struct {
}

func (s CreateEdgeFuncTriggerParameters) String() string {
  return tea.Prettify(s)
}

func (s CreateEdgeFuncTriggerParameters) GoString() string {
  return s.String()
}

type CreateEdgeFuncTriggerResponse struct {
  // {"en":"Status code","zh_CN":"响应码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Status description","zh_CN":"响应信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Response Data","zh_CN":"响应数据"}
  Data *CreateEdgeFuncTriggerResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
}

func (s CreateEdgeFuncTriggerResponse) String() string {
  return tea.Prettify(s)
}

func (s CreateEdgeFuncTriggerResponse) GoString() string {
  return s.String()
}

func (s *CreateEdgeFuncTriggerResponse) SetCode(v string) *CreateEdgeFuncTriggerResponse {
  s.Code = &v
  return s
}

func (s *CreateEdgeFuncTriggerResponse) SetMessage(v string) *CreateEdgeFuncTriggerResponse {
  s.Message = &v
  return s
}

func (s *CreateEdgeFuncTriggerResponse) SetData(v *CreateEdgeFuncTriggerResponseData) *CreateEdgeFuncTriggerResponse {
  s.Data = v
  return s
}

type CreateEdgeFuncTriggerResponseData struct {
  // {"en":"Domain ID","zh_CN":"域名 iD"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {"en":"Domain","zh_CN":"域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {"en":"Router List","zh_CN":"路由列表"}
  Routers *CreateEdgeFuncTriggerResponseDataRouters `json:"routers,omitempty" xml:"routers,omitempty" require:"true" type:"Struct"`
}

func (s CreateEdgeFuncTriggerResponseData) String() string {
  return tea.Prettify(s)
}

func (s CreateEdgeFuncTriggerResponseData) GoString() string {
  return s.String()
}

func (s *CreateEdgeFuncTriggerResponseData) SetId(v string) *CreateEdgeFuncTriggerResponseData {
  s.Id = &v
  return s
}

func (s *CreateEdgeFuncTriggerResponseData) SetDomain(v string) *CreateEdgeFuncTriggerResponseData {
  s.Domain = &v
  return s
}

func (s *CreateEdgeFuncTriggerResponseData) SetRouters(v *CreateEdgeFuncTriggerResponseDataRouters) *CreateEdgeFuncTriggerResponseData {
  s.Routers = v
  return s
}

type CreateEdgeFuncTriggerResponseDataRouters struct {
  // {"en":"Func name","zh_CN":"函数名称"}
  FuncName *string `json:"funcName,omitempty" xml:"funcName,omitempty" require:"true"`
  // {"en":"Match path","zh_CN":"匹配路径"}
  Path *string `json:"path,omitempty" xml:"path,omitempty" require:"true"`
}

func (s CreateEdgeFuncTriggerResponseDataRouters) String() string {
  return tea.Prettify(s)
}

func (s CreateEdgeFuncTriggerResponseDataRouters) GoString() string {
  return s.String()
}

func (s *CreateEdgeFuncTriggerResponseDataRouters) SetFuncName(v string) *CreateEdgeFuncTriggerResponseDataRouters {
  s.FuncName = &v
  return s
}

func (s *CreateEdgeFuncTriggerResponseDataRouters) SetPath(v string) *CreateEdgeFuncTriggerResponseDataRouters {
  s.Path = &v
  return s
}

type CreateEdgeFuncTriggerResponseHeader struct {
}

func (s CreateEdgeFuncTriggerResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s CreateEdgeFuncTriggerResponseHeader) GoString() string {
  return s.String()
}




type UploadFuncCodeRequest struct {
  // {"en":"func name","zh_CN":"函数名称"}
  FuncName *string `json:"FuncName,omitempty" xml:"FuncName,omitempty"`
  // {"en":"commit message","zh_CN":"提交信息"}
  CommitMessage *string `json:"CommitMessage,omitempty" xml:"CommitMessage,omitempty"`
  // {"en":"func file list","zh_CN":"函数文件列表"}
  Files []*UploadFuncCodeRequestFiles `json:"Files,omitempty" xml:"Files,omitempty" type:"Repeated"`
}

func (s UploadFuncCodeRequest) String() string {
  return tea.Prettify(s)
}

func (s UploadFuncCodeRequest) GoString() string {
  return s.String()
}

func (s *UploadFuncCodeRequest) SetFuncName(v string) *UploadFuncCodeRequest {
  s.FuncName = &v
  return s
}

func (s *UploadFuncCodeRequest) SetCommitMessage(v string) *UploadFuncCodeRequest {
  s.CommitMessage = &v
  return s
}

func (s *UploadFuncCodeRequest) SetFiles(v []*UploadFuncCodeRequestFiles) *UploadFuncCodeRequest {
  s.Files = v
  return s
}

type UploadFuncCodeRequestFiles struct     {
  // {"en":"FileName","zh_CN":"文件名称"}
  FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
  // {"en":"File code (base64 encoding)","zh_CN":"文件所含代码（Base64 编码）"}
  Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s UploadFuncCodeRequestFiles) String() string {
  return tea.Prettify(s)
}

func (s UploadFuncCodeRequestFiles) GoString() string {
  return s.String()
}

func (s *UploadFuncCodeRequestFiles) SetFileName(v string) *UploadFuncCodeRequestFiles {
  s.FileName = &v
  return s
}

func (s *UploadFuncCodeRequestFiles) SetCode(v string) *UploadFuncCodeRequestFiles {
  s.Code = &v
  return s
}

type UploadFuncCodeRequestHeader struct {
}

func (s UploadFuncCodeRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s UploadFuncCodeRequestHeader) GoString() string {
  return s.String()
}

type UploadFuncCodePaths struct {
}

func (s UploadFuncCodePaths) String() string {
  return tea.Prettify(s)
}

func (s UploadFuncCodePaths) GoString() string {
  return s.String()
}

type UploadFuncCodeParameters struct {
}

func (s UploadFuncCodeParameters) String() string {
  return tea.Prettify(s)
}

func (s UploadFuncCodeParameters) GoString() string {
  return s.String()
}

type UploadFuncCodeResponse struct {
  // {"en":"code","zh_CN":"状态码"}
  Code *string `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
  // {"en":"message","zh_CN":"状态码描述"}
  Message *string `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
}

func (s UploadFuncCodeResponse) String() string {
  return tea.Prettify(s)
}

func (s UploadFuncCodeResponse) GoString() string {
  return s.String()
}

func (s *UploadFuncCodeResponse) SetCode(v string) *UploadFuncCodeResponse {
  s.Code = &v
  return s
}

func (s *UploadFuncCodeResponse) SetMessage(v string) *UploadFuncCodeResponse {
  s.Message = &v
  return s
}

type UploadFuncCodeResponseHeader struct {
}

func (s UploadFuncCodeResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s UploadFuncCodeResponseHeader) GoString() string {
  return s.String()
}




