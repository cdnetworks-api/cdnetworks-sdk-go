package protectassets

// This file is auto-generated, don't edit it. Thanks.
import (
  "github.com/alibabacloud-go/tea/tea"
)

type BatchDeleteAPIDefinitionRequest struct {
  // {"en":"API definition ID array.","zh_CN":"API定义ID数组。"}
  IdList []*string `json:"idList,omitempty" xml:"idList,omitempty" require:"true" type:"Repeated"`
}

func (s BatchDeleteAPIDefinitionRequest) String() string {
  return tea.Prettify(s)
}

func (s BatchDeleteAPIDefinitionRequest) GoString() string {
  return s.String()
}

func (s *BatchDeleteAPIDefinitionRequest) SetIdList(v []*string) *BatchDeleteAPIDefinitionRequest {
  s.IdList = v
  return s
}

type BatchDeleteAPIDefinitionRequestHeader struct {
  // {"en":"Security service type. Please enter a specific service type, if you purchase multiple security services.","zh_CN":"安全服务类型。有使用多个不同的安全服务时，需要填写具体的服务类型。"}
  ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty"`
}

func (s BatchDeleteAPIDefinitionRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s BatchDeleteAPIDefinitionRequestHeader) GoString() string {
  return s.String()
}

func (s *BatchDeleteAPIDefinitionRequestHeader) SetServiceType(v string) *BatchDeleteAPIDefinitionRequestHeader {
  s.ServiceType = &v
  return s
}

type BatchDeleteAPIDefinitionPaths struct {
}

func (s BatchDeleteAPIDefinitionPaths) String() string {
  return tea.Prettify(s)
}

func (s BatchDeleteAPIDefinitionPaths) GoString() string {
  return s.String()
}

type BatchDeleteAPIDefinitionParameters struct {
}

func (s BatchDeleteAPIDefinitionParameters) String() string {
  return tea.Prettify(s)
}

func (s BatchDeleteAPIDefinitionParameters) GoString() string {
  return s.String()
}

type BatchDeleteAPIDefinitionResponse struct {
  // {"dictionary":"belong=WAAP-MS-Ext|dict=waap_retCodeEnum","en":"Please refer to the error code for exceptions.","zh_CN":"请参照错误码。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Description.","zh_CN":"描述信息。"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"Data.","zh_CN":"出参数据。"}
  Data *bool `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s BatchDeleteAPIDefinitionResponse) String() string {
  return tea.Prettify(s)
}

func (s BatchDeleteAPIDefinitionResponse) GoString() string {
  return s.String()
}

func (s *BatchDeleteAPIDefinitionResponse) SetCode(v string) *BatchDeleteAPIDefinitionResponse {
  s.Code = &v
  return s
}

func (s *BatchDeleteAPIDefinitionResponse) SetMsg(v string) *BatchDeleteAPIDefinitionResponse {
  s.Msg = &v
  return s
}

func (s *BatchDeleteAPIDefinitionResponse) SetData(v bool) *BatchDeleteAPIDefinitionResponse {
  s.Data = &v
  return s
}

type BatchDeleteAPIDefinitionResponseHeader struct {
}

func (s BatchDeleteAPIDefinitionResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s BatchDeleteAPIDefinitionResponseHeader) GoString() string {
  return s.String()
}




type QueryAPIDefinitionDetailRequest struct {
  // {"en":"API definition ID.","zh_CN":"API定义ID。"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
}

func (s QueryAPIDefinitionDetailRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryAPIDefinitionDetailRequest) GoString() string {
  return s.String()
}

func (s *QueryAPIDefinitionDetailRequest) SetId(v string) *QueryAPIDefinitionDetailRequest {
  s.Id = &v
  return s
}

type QueryAPIDefinitionDetailRequestHeader struct {
  // {"en":"Security service type. Please enter a specific service type, if you purchase multiple security services.","zh_CN":"安全服务类型。有使用多个不同的安全服务时，需要填写具体的服务类型。"}
  ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty"`
}

func (s QueryAPIDefinitionDetailRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryAPIDefinitionDetailRequestHeader) GoString() string {
  return s.String()
}

func (s *QueryAPIDefinitionDetailRequestHeader) SetServiceType(v string) *QueryAPIDefinitionDetailRequestHeader {
  s.ServiceType = &v
  return s
}

type QueryAPIDefinitionDetailPaths struct {
}

func (s QueryAPIDefinitionDetailPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryAPIDefinitionDetailPaths) GoString() string {
  return s.String()
}

type QueryAPIDefinitionDetailParameters struct {
}

func (s QueryAPIDefinitionDetailParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryAPIDefinitionDetailParameters) GoString() string {
  return s.String()
}

type QueryAPIDefinitionDetailResponse struct {
  // {"dictionary":"belong=WAAP-MS-Ext|dict=waap_retCodeEnum","en":"Please refer to the error code for exceptions.","zh_CN":"请参照错误码。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Description.","zh_CN":"描述信息。"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"Data.","zh_CN":"出参数据。"}
  Data *QueryAPIDefinitionDetailResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
}

func (s QueryAPIDefinitionDetailResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryAPIDefinitionDetailResponse) GoString() string {
  return s.String()
}

func (s *QueryAPIDefinitionDetailResponse) SetCode(v string) *QueryAPIDefinitionDetailResponse {
  s.Code = &v
  return s
}

func (s *QueryAPIDefinitionDetailResponse) SetMsg(v string) *QueryAPIDefinitionDetailResponse {
  s.Msg = &v
  return s
}

func (s *QueryAPIDefinitionDetailResponse) SetData(v *QueryAPIDefinitionDetailResponseData) *QueryAPIDefinitionDetailResponse {
  s.Data = v
  return s
}

type QueryAPIDefinitionDetailResponseData struct {
  // {"en":"Basic information.","zh_CN":"基础信息。"}
  Basic *QueryAPIDefinitionDetailResponseDataBasic `json:"basic,omitempty" xml:"basic,omitempty" require:"true" type:"Struct"`
  // {"en":"Endpoint information.","zh_CN":"端点信息。"}
  Endpoint *QueryAPIDefinitionDetailResponseDataEndpoint `json:"endpoint,omitempty" xml:"endpoint,omitempty" require:"true" type:"Struct"`
  // {"en":"Authentication configuration.","zh_CN":"鉴权配置。"}
  Auth *QueryAPIDefinitionDetailResponseDataAuth `json:"auth,omitempty" xml:"auth,omitempty" require:"true" type:"Struct"`
  // {"en":"Body restrictions.","zh_CN":"body限制。"}
  BodyLimit *QueryAPIDefinitionDetailResponseDataBodyLimit `json:"bodyLimit,omitempty" xml:"bodyLimit,omitempty" require:"true" type:"Struct"`
  // {"en":"Parameter limit.","zh_CN":"参数限制。"}
  ParamLimit *QueryAPIDefinitionDetailResponseDataParamLimit `json:"paramLimit,omitempty" xml:"paramLimit,omitempty" require:"true" type:"Struct"`
}

func (s QueryAPIDefinitionDetailResponseData) String() string {
  return tea.Prettify(s)
}

func (s QueryAPIDefinitionDetailResponseData) GoString() string {
  return s.String()
}

func (s *QueryAPIDefinitionDetailResponseData) SetBasic(v *QueryAPIDefinitionDetailResponseDataBasic) *QueryAPIDefinitionDetailResponseData {
  s.Basic = v
  return s
}

func (s *QueryAPIDefinitionDetailResponseData) SetEndpoint(v *QueryAPIDefinitionDetailResponseDataEndpoint) *QueryAPIDefinitionDetailResponseData {
  s.Endpoint = v
  return s
}

func (s *QueryAPIDefinitionDetailResponseData) SetAuth(v *QueryAPIDefinitionDetailResponseDataAuth) *QueryAPIDefinitionDetailResponseData {
  s.Auth = v
  return s
}

func (s *QueryAPIDefinitionDetailResponseData) SetBodyLimit(v *QueryAPIDefinitionDetailResponseDataBodyLimit) *QueryAPIDefinitionDetailResponseData {
  s.BodyLimit = v
  return s
}

func (s *QueryAPIDefinitionDetailResponseData) SetParamLimit(v *QueryAPIDefinitionDetailResponseDataParamLimit) *QueryAPIDefinitionDetailResponseData {
  s.ParamLimit = v
  return s
}

type QueryAPIDefinitionDetailResponseDataBasic struct {
  // {"en":"API definition ID.","zh_CN":"API定义ID。"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {"en":"API name.","zh_CN":"API名称。"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"API groups.","zh_CN":"API分组。"}
  GroupName *string `json:"groupName,omitempty" xml:"groupName,omitempty" require:"true"`
  // {"en":"Description.","zh_CN":"备注。"}
  Remark *string `json:"remark,omitempty" xml:"remark,omitempty" require:"true"`
  // {"en":"Attributed domain.","zh_CN":"归属域名。"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {"en":"API base path.","zh_CN":"端点路径。"}
  Path *string `json:"path,omitempty" xml:"path,omitempty" require:"true"`
  // {"en":"Creation time, format: yyyy-MM-dd HH:mm:ss.","zh_CN":"创建时间，格式：yyyy-MM-dd HH:mm:ss。"}
  CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty" require:"true"`
  // {"en":"Update time, format: yyyy-MM-dd HH:mm:ss.","zh_CN":"更新时间，格式：yyyy-MM-dd HH:mm:ss。"}
  UpdateTime *string `json:"updateTime,omitempty" xml:"updateTime,omitempty" require:"true"`
}

func (s QueryAPIDefinitionDetailResponseDataBasic) String() string {
  return tea.Prettify(s)
}

func (s QueryAPIDefinitionDetailResponseDataBasic) GoString() string {
  return s.String()
}

func (s *QueryAPIDefinitionDetailResponseDataBasic) SetId(v string) *QueryAPIDefinitionDetailResponseDataBasic {
  s.Id = &v
  return s
}

func (s *QueryAPIDefinitionDetailResponseDataBasic) SetName(v string) *QueryAPIDefinitionDetailResponseDataBasic {
  s.Name = &v
  return s
}

func (s *QueryAPIDefinitionDetailResponseDataBasic) SetGroupName(v string) *QueryAPIDefinitionDetailResponseDataBasic {
  s.GroupName = &v
  return s
}

func (s *QueryAPIDefinitionDetailResponseDataBasic) SetRemark(v string) *QueryAPIDefinitionDetailResponseDataBasic {
  s.Remark = &v
  return s
}

func (s *QueryAPIDefinitionDetailResponseDataBasic) SetDomain(v string) *QueryAPIDefinitionDetailResponseDataBasic {
  s.Domain = &v
  return s
}

func (s *QueryAPIDefinitionDetailResponseDataBasic) SetPath(v string) *QueryAPIDefinitionDetailResponseDataBasic {
  s.Path = &v
  return s
}

func (s *QueryAPIDefinitionDetailResponseDataBasic) SetCreateTime(v string) *QueryAPIDefinitionDetailResponseDataBasic {
  s.CreateTime = &v
  return s
}

func (s *QueryAPIDefinitionDetailResponseDataBasic) SetUpdateTime(v string) *QueryAPIDefinitionDetailResponseDataBasic {
  s.UpdateTime = &v
  return s
}

type QueryAPIDefinitionDetailResponseDataEndpoint struct {
  // {"en":"Request methods.\nGET:GET, configurable parameter limits\nPOST:POST, configurable parameter limits\nDELETE:DELETE, configurable parameter limits\nUPDATE:UPDATE\nPUT:PUT, configurable parameter limits\nHEAD:HEAD, configurable parameter limits\nCONNECT:CONNECT\nOPTIONS:OPTIONS, configurable parameter limits\nCOPY:COPY\nLOCK:LOCK\nUNLOCK:UNLOCK\nTRACE:TRACE\nPATCH:PATCH, configurable parameter limits\nPROPFIND:PROPFIND\nMKCOL:MKCOL\nMOVE:MOVE","zh_CN":"请求方法。\nGET：GET，可配置参数限制\nPOST：POST，可配置参数限制\nDELETE：DELETE，可配置参数限制\nUPDATE：UPDATE\nPUT：PUT，可配置参数限制\nHEAD：HEAD，可配置参数限制\nCONNECT：CONNECT\nOPTIONS：OPTIONS，可配置参数限制\nCOPY：COPY\nLOCK：LOCK\nUNLOCK：UNLOCK\nTRACE：TRACE\nPATCH：PATCH，可配置参数限制\nPROPFIND：PROPFIND\nMKCOL：MKCOL\nMOVE：MOVE","exampleValue":"GET,POST,DELETE,UPDATE,PUT,HEAD,CONNECT,OPTIONS,COPY,LOCK,UNLOCK,TRACE,PATCH,PROPFIND,MKCOL,MOVE"}
  RequestMethod *string `json:"requestMethod,omitempty" xml:"requestMethod,omitempty" require:"true"`
  // {"en":"API type.\nNORMAL:Common API\nWHEN_CASE:Common API","zh_CN":"API类型。\nNORMAL：普通接口\nWHEN_CASE：when-case接口","exampleValue":"NORMAL,WHEN_CASE"}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
  // {"en":"Path matching type.\nEQUAL: Complete match\nREGEX: Regular matching","zh_CN":"路径匹配类型。\nEQUAL：完整匹配\nREGEX：正则匹配","exampleValue":"EQUAL,REGEX"}
  MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
  // {"en":"API base path.","zh_CN":"端点路径。"}
  Path *string `json:"path,omitempty" xml:"path,omitempty" require:"true"`
  // {"en":"Case sensitive.\nTRUE:Yes\nFALSE:No","zh_CN":"大小写是否敏感。\nTRUE：是\nFALSE：否","exampleValue":"TRUE,FALSE"}
  CaseSensitive *string `json:"caseSensitive,omitempty" xml:"caseSensitive,omitempty" require:"true"`
  // {"en":"Match QueryAPIDefinitionDetailParameters in the Path.\nTRUE:Yes\nFALSE:No","zh_CN":"匹配路径参数。\nTRUE：是\nFALSE：否","exampleValue":"TRUE,FALSE"}
  MatchPathVar *string `json:"matchPathVar,omitempty" xml:"matchPathVar,omitempty" require:"true"`
}

func (s QueryAPIDefinitionDetailResponseDataEndpoint) String() string {
  return tea.Prettify(s)
}

func (s QueryAPIDefinitionDetailResponseDataEndpoint) GoString() string {
  return s.String()
}

func (s *QueryAPIDefinitionDetailResponseDataEndpoint) SetRequestMethod(v string) *QueryAPIDefinitionDetailResponseDataEndpoint {
  s.RequestMethod = &v
  return s
}

func (s *QueryAPIDefinitionDetailResponseDataEndpoint) SetType(v string) *QueryAPIDefinitionDetailResponseDataEndpoint {
  s.Type = &v
  return s
}

func (s *QueryAPIDefinitionDetailResponseDataEndpoint) SetMatchType(v string) *QueryAPIDefinitionDetailResponseDataEndpoint {
  s.MatchType = &v
  return s
}

func (s *QueryAPIDefinitionDetailResponseDataEndpoint) SetPath(v string) *QueryAPIDefinitionDetailResponseDataEndpoint {
  s.Path = &v
  return s
}

func (s *QueryAPIDefinitionDetailResponseDataEndpoint) SetCaseSensitive(v string) *QueryAPIDefinitionDetailResponseDataEndpoint {
  s.CaseSensitive = &v
  return s
}

func (s *QueryAPIDefinitionDetailResponseDataEndpoint) SetMatchPathVar(v string) *QueryAPIDefinitionDetailResponseDataEndpoint {
  s.MatchPathVar = &v
  return s
}

type QueryAPIDefinitionDetailResponseDataAuth struct {
  // {"en":"Authentication method.\nNO_NEED:No authentication required\nSIGN:Key authentication","zh_CN":"鉴权方法。\nNO_NEED：免鉴权\nSIGN：秘钥对鉴权","exampleValue":"NO_NEED,SIGN"}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
  // {"en":"Authentication key.","zh_CN":"鉴权秘钥。"}
  AuthKey *string `json:"authKey,omitempty" xml:"authKey,omitempty" require:"true"`
  // {"en":"Authentication parameter location.\nHTTP_HEADER:Http Header\nQUERY_STRING:Query String\nCOOKIE:Cookie","zh_CN":"鉴权参数位置。\nHTTP_HEADER：Http Header\nQUERY_STRING：Query String\nCOOKIE：Cookie","exampleValue":"HTTP_HEADER,QUERY_STRING,COOKIE"}
  ParamPosition *string `json:"paramPosition,omitempty" xml:"paramPosition,omitempty" require:"true"`
  // {"en":"Authentication parameter name.","zh_CN":"鉴权参数名称。"}
  ParamName *string `json:"paramName,omitempty" xml:"paramName,omitempty" require:"true"`
  // {"en":"Authentication token validity period, in seconds.","zh_CN":"鉴权有效期，单位秒。"}
  ValidityTime *int64 `json:"validityTime,omitempty" xml:"validityTime,omitempty" require:"true"`
}

func (s QueryAPIDefinitionDetailResponseDataAuth) String() string {
  return tea.Prettify(s)
}

func (s QueryAPIDefinitionDetailResponseDataAuth) GoString() string {
  return s.String()
}

func (s *QueryAPIDefinitionDetailResponseDataAuth) SetType(v string) *QueryAPIDefinitionDetailResponseDataAuth {
  s.Type = &v
  return s
}

func (s *QueryAPIDefinitionDetailResponseDataAuth) SetAuthKey(v string) *QueryAPIDefinitionDetailResponseDataAuth {
  s.AuthKey = &v
  return s
}

func (s *QueryAPIDefinitionDetailResponseDataAuth) SetParamPosition(v string) *QueryAPIDefinitionDetailResponseDataAuth {
  s.ParamPosition = &v
  return s
}

func (s *QueryAPIDefinitionDetailResponseDataAuth) SetParamName(v string) *QueryAPIDefinitionDetailResponseDataAuth {
  s.ParamName = &v
  return s
}

func (s *QueryAPIDefinitionDetailResponseDataAuth) SetValidityTime(v int64) *QueryAPIDefinitionDetailResponseDataAuth {
  s.ValidityTime = &v
  return s
}

type QueryAPIDefinitionDetailResponseDataBodyLimit struct {
  // {"en":"Request body restriction switch.\nON:ON\nOFF:OFF","zh_CN":"请求body限制开关。\nON：开启\nOFF：关闭","exampleValue":"ON,OFF"}
  DefendSwitch *string `json:"defendSwitch,omitempty" xml:"defendSwitch,omitempty" require:"true"`
  // {"en":"Content-Type.\nFORM:FORM\nJSON:JSON\nANY:ANY\nEMPTY:EMPTY or NON-EXISTENT","zh_CN":"Content-Type。\nFORM：FORM表单\nJSON：JSON\nANY：任意\nEMPTY：为空或不存在","exampleValue":"FORM,JSON,ANY,EMPTY"}
  ContentType *string `json:"contentType,omitempty" xml:"contentType,omitempty" require:"true"`
  // {"en":"Maximum body limit(bytes).","zh_CN":"Body最大限制(bytes)。"}
  BodyLimitMax *int64 `json:"bodyLimitMax,omitempty" xml:"bodyLimitMax,omitempty" require:"true"`
  // {"en":"Maximum nesting depth.","zh_CN":"最大嵌套层数。"}
  NestMax *int64 `json:"nestMax,omitempty" xml:"nestMax,omitempty" require:"true"`
  // {"en":"Maximum number of parameters for JSON.","zh_CN":"JSON最大参数个数。"}
  ParamCountMax *int64 `json:"paramCountMax,omitempty" xml:"paramCountMax,omitempty" require:"true"`
}

func (s QueryAPIDefinitionDetailResponseDataBodyLimit) String() string {
  return tea.Prettify(s)
}

func (s QueryAPIDefinitionDetailResponseDataBodyLimit) GoString() string {
  return s.String()
}

func (s *QueryAPIDefinitionDetailResponseDataBodyLimit) SetDefendSwitch(v string) *QueryAPIDefinitionDetailResponseDataBodyLimit {
  s.DefendSwitch = &v
  return s
}

func (s *QueryAPIDefinitionDetailResponseDataBodyLimit) SetContentType(v string) *QueryAPIDefinitionDetailResponseDataBodyLimit {
  s.ContentType = &v
  return s
}

func (s *QueryAPIDefinitionDetailResponseDataBodyLimit) SetBodyLimitMax(v int64) *QueryAPIDefinitionDetailResponseDataBodyLimit {
  s.BodyLimitMax = &v
  return s
}

func (s *QueryAPIDefinitionDetailResponseDataBodyLimit) SetNestMax(v int64) *QueryAPIDefinitionDetailResponseDataBodyLimit {
  s.NestMax = &v
  return s
}

func (s *QueryAPIDefinitionDetailResponseDataBodyLimit) SetParamCountMax(v int64) *QueryAPIDefinitionDetailResponseDataBodyLimit {
  s.ParamCountMax = &v
  return s
}

type QueryAPIDefinitionDetailResponseDataParamLimit struct {
  // {"en":"Basic information.","zh_CN":"基础信息。"}
  Basic *QueryAPIDefinitionDetailResponseDataParamLimitBasic `json:"basic,omitempty" xml:"basic,omitempty" require:"true" type:"Struct"`
  // {"en":"Method attributes.","zh_CN":"方法属性。"}
  MethodList []*QueryAPIDefinitionDetailResponseDataParamLimitMethodList `json:"methodList,omitempty" xml:"methodList,omitempty" require:"true" type:"Repeated"`
}

func (s QueryAPIDefinitionDetailResponseDataParamLimit) String() string {
  return tea.Prettify(s)
}

func (s QueryAPIDefinitionDetailResponseDataParamLimit) GoString() string {
  return s.String()
}

func (s *QueryAPIDefinitionDetailResponseDataParamLimit) SetBasic(v *QueryAPIDefinitionDetailResponseDataParamLimitBasic) *QueryAPIDefinitionDetailResponseDataParamLimit {
  s.Basic = v
  return s
}

func (s *QueryAPIDefinitionDetailResponseDataParamLimit) SetMethodList(v []*QueryAPIDefinitionDetailResponseDataParamLimitMethodList) *QueryAPIDefinitionDetailResponseDataParamLimit {
  s.MethodList = v
  return s
}

type QueryAPIDefinitionDetailResponseDataParamLimitBasic struct {
  // {"en":"Parameter limit.\nON:ON\nOFF:OFF","zh_CN":"参数限制。\nON：ON\nOFF：OFF","exampleValue":"ON,OFF"}
  DefendSwitch *string `json:"defendSwitch,omitempty" xml:"defendSwitch,omitempty" require:"true"`
  // {"en":"Query String Parameter Detection Mode.\nLOOSE:Loose mode - detect some parameters\nSTRICT:Strict mode - checks all parameters","zh_CN":"Query String参数检测模式。\nLOOSE：宽松模式-检测部分参数\nSTRICT：严格模式-检测所有参数","exampleValue":"LOOSE,STRICT"}
  ParamCheckMode *string `json:"paramCheckMode,omitempty" xml:"paramCheckMode,omitempty" require:"true"`
}

func (s QueryAPIDefinitionDetailResponseDataParamLimitBasic) String() string {
  return tea.Prettify(s)
}

func (s QueryAPIDefinitionDetailResponseDataParamLimitBasic) GoString() string {
  return s.String()
}

func (s *QueryAPIDefinitionDetailResponseDataParamLimitBasic) SetDefendSwitch(v string) *QueryAPIDefinitionDetailResponseDataParamLimitBasic {
  s.DefendSwitch = &v
  return s
}

func (s *QueryAPIDefinitionDetailResponseDataParamLimitBasic) SetParamCheckMode(v string) *QueryAPIDefinitionDetailResponseDataParamLimitBasic {
  s.ParamCheckMode = &v
  return s
}

type QueryAPIDefinitionDetailResponseDataParamLimitMethodList struct     {
  // {"en":"Request methods.\nGET:GET\nPOST:POST\nDELETE:DELETE\nPUT:PUT\nHEAD:HEAD\nOPTIONS:OPTIONS\nPATCH:PATCH","zh_CN":"请求方法。\nGET：GET\nPOST：POST\nDELETE：DELETE\nPUT：PUT\nHEAD：HEAD\nOPTIONS：OPTIONS\nPATCH：PATCH","exampleValue":"GET,POST,DELETE,PUT,HEAD,OPTIONS,PATCH"}
  Method *string `json:"method,omitempty" xml:"method,omitempty" require:"true"`
  // {"en":"Whether to define body parameters.\nTRUE:Yes\nFALSE:No","zh_CN":"是否定义body参数。\nTRUE：是\nFALSE：否","exampleValue":"TRUE,FALSE"}
  BodyFlag *string `json:"bodyFlag,omitempty" xml:"bodyFlag,omitempty" require:"true"`
  // {"en":"Body parameter type.\nJSON:JSON\nFROM:FROM","zh_CN":"body参数类型。\nJSON：JSON\nFROM：FROM表单","exampleValue":"FORM,JSON"}
  BodyType *string `json:"bodyType,omitempty" xml:"bodyType,omitempty" require:"true"`
  // {"en":"Normal parameter list.","zh_CN":"普通参数数组。"}
  NormalParamList []*QueryAPIDefinitionDetailResponseDataParamLimitMethodListNormalParamList `json:"normalParamList,omitempty" xml:"normalParamList,omitempty" require:"true" type:"Repeated"`
  // {"en":"FROM form parameter array.","zh_CN":"FROM 表单参数数组。"}
  FormParamList []*QueryAPIDefinitionDetailResponseDataParamLimitMethodListFormParamList `json:"formParamList,omitempty" xml:"formParamList,omitempty" require:"true" type:"Repeated"`
  // {"en":"JSON parameter array.","zh_CN":"JSON参数数组。"}
  JsonParam *QueryAPIDefinitionDetailResponseDataParamLimitMethodListJsonParam `json:"jsonParam,omitempty" xml:"jsonParam,omitempty" require:"true" type:"Struct"`
}

func (s QueryAPIDefinitionDetailResponseDataParamLimitMethodList) String() string {
  return tea.Prettify(s)
}

func (s QueryAPIDefinitionDetailResponseDataParamLimitMethodList) GoString() string {
  return s.String()
}

func (s *QueryAPIDefinitionDetailResponseDataParamLimitMethodList) SetMethod(v string) *QueryAPIDefinitionDetailResponseDataParamLimitMethodList {
  s.Method = &v
  return s
}

func (s *QueryAPIDefinitionDetailResponseDataParamLimitMethodList) SetBodyFlag(v string) *QueryAPIDefinitionDetailResponseDataParamLimitMethodList {
  s.BodyFlag = &v
  return s
}

func (s *QueryAPIDefinitionDetailResponseDataParamLimitMethodList) SetBodyType(v string) *QueryAPIDefinitionDetailResponseDataParamLimitMethodList {
  s.BodyType = &v
  return s
}

func (s *QueryAPIDefinitionDetailResponseDataParamLimitMethodList) SetNormalParamList(v []*QueryAPIDefinitionDetailResponseDataParamLimitMethodListNormalParamList) *QueryAPIDefinitionDetailResponseDataParamLimitMethodList {
  s.NormalParamList = v
  return s
}

func (s *QueryAPIDefinitionDetailResponseDataParamLimitMethodList) SetFormParamList(v []*QueryAPIDefinitionDetailResponseDataParamLimitMethodListFormParamList) *QueryAPIDefinitionDetailResponseDataParamLimitMethodList {
  s.FormParamList = v
  return s
}

func (s *QueryAPIDefinitionDetailResponseDataParamLimitMethodList) SetJsonParam(v *QueryAPIDefinitionDetailResponseDataParamLimitMethodListJsonParam) *QueryAPIDefinitionDetailResponseDataParamLimitMethodList {
  s.JsonParam = v
  return s
}

type QueryAPIDefinitionDetailResponseDataParamLimitMethodListNormalParamList struct     {
  // {"en":"Request methods.\nGET:GET\nPOST:POST\nDELETE:DELETE\nPUT:PUT\nHEAD:HEAD\nOPTIONS:OPTIONS\nPATCH:PATCH","zh_CN":"请求方法。\nGET：GET\nPOST：POST\nDELETE：DELETE\nPUT：PUT\nHEAD：HEAD\nOPTIONS：OPTIONS\nPATCH：PATCH","exampleValue":"GET,POST,DELETE,PUT,HEAD,OPTIONS,PATCH"}
  Method *string `json:"method,omitempty" xml:"method,omitempty" require:"true"`
  // {"en":"Parameter name.","zh_CN":"参数名称。"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"Parameter position.\nHTTP_HEADER:Http Header\nQUERY_STRING:Query String\nCOOKIE:Cookie\nPATH_PARAMS:Path","zh_CN":"参数变量。\nHTTP_HEADER：Http Header\nQUERY_STRING：Query String\nCOOKIE：Cookie\nPATH_PARAMS：路径变量","exampleValue":"HTTP_HEADER,QUERY_STRING,COOKIE,PATH_PARAMS"}
  ParamPosition *string `json:"paramPosition,omitempty" xml:"paramPosition,omitempty" require:"true"`
  // {"en":"Parameter type.\nInteger:Integer\nNumber:Number\nString:String\nBoolean:Boolean\nEnumeration:Enumeration","zh_CN":"参数类型。\nInteger：整数\nNumber：数字\nString：字符串\nBoolean：布尔\nEnumeration：枚举","exampleValue":"Integer,Number,String,Boolean,Enumeration"}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
  // {"en":"Minimum value/Minimum length.","zh_CN":"最小值、最小长度。"}
  MinVal *string `json:"minVal,omitempty" xml:"minVal,omitempty" require:"true"`
  // {"en":"Maximum value/Maximum length.","zh_CN":"最大值、最大长度。"}
  MaxVal *string `json:"maxVal,omitempty" xml:"maxVal,omitempty" require:"true"`
  // {"en":"Required.\nTRUE:Yes\nFALSE:No","zh_CN":"必带参数。\nTRUE：是\nFALSE：否","exampleValue":"TRUE,FALSE"}
  Required *string `json:"required,omitempty" xml:"required,omitempty" require:"true"`
  // {"en":"Content.","zh_CN":"内容。"}
  Content *string `json:"content,omitempty" xml:"content,omitempty" require:"true"`
  // {"en":"Description.","zh_CN":"备注。"}
  Remark *string `json:"remark,omitempty" xml:"remark,omitempty" require:"true"`
}

func (s QueryAPIDefinitionDetailResponseDataParamLimitMethodListNormalParamList) String() string {
  return tea.Prettify(s)
}

func (s QueryAPIDefinitionDetailResponseDataParamLimitMethodListNormalParamList) GoString() string {
  return s.String()
}

func (s *QueryAPIDefinitionDetailResponseDataParamLimitMethodListNormalParamList) SetMethod(v string) *QueryAPIDefinitionDetailResponseDataParamLimitMethodListNormalParamList {
  s.Method = &v
  return s
}

func (s *QueryAPIDefinitionDetailResponseDataParamLimitMethodListNormalParamList) SetName(v string) *QueryAPIDefinitionDetailResponseDataParamLimitMethodListNormalParamList {
  s.Name = &v
  return s
}

func (s *QueryAPIDefinitionDetailResponseDataParamLimitMethodListNormalParamList) SetParamPosition(v string) *QueryAPIDefinitionDetailResponseDataParamLimitMethodListNormalParamList {
  s.ParamPosition = &v
  return s
}

func (s *QueryAPIDefinitionDetailResponseDataParamLimitMethodListNormalParamList) SetType(v string) *QueryAPIDefinitionDetailResponseDataParamLimitMethodListNormalParamList {
  s.Type = &v
  return s
}

func (s *QueryAPIDefinitionDetailResponseDataParamLimitMethodListNormalParamList) SetMinVal(v string) *QueryAPIDefinitionDetailResponseDataParamLimitMethodListNormalParamList {
  s.MinVal = &v
  return s
}

func (s *QueryAPIDefinitionDetailResponseDataParamLimitMethodListNormalParamList) SetMaxVal(v string) *QueryAPIDefinitionDetailResponseDataParamLimitMethodListNormalParamList {
  s.MaxVal = &v
  return s
}

func (s *QueryAPIDefinitionDetailResponseDataParamLimitMethodListNormalParamList) SetRequired(v string) *QueryAPIDefinitionDetailResponseDataParamLimitMethodListNormalParamList {
  s.Required = &v
  return s
}

func (s *QueryAPIDefinitionDetailResponseDataParamLimitMethodListNormalParamList) SetContent(v string) *QueryAPIDefinitionDetailResponseDataParamLimitMethodListNormalParamList {
  s.Content = &v
  return s
}

func (s *QueryAPIDefinitionDetailResponseDataParamLimitMethodListNormalParamList) SetRemark(v string) *QueryAPIDefinitionDetailResponseDataParamLimitMethodListNormalParamList {
  s.Remark = &v
  return s
}

type QueryAPIDefinitionDetailResponseDataParamLimitMethodListFormParamList struct     {
  // {"en":"Request methods.\nPOST:POST\nPUT:PUT\nPATCH:PATCH","zh_CN":"请求方法。\nPOST：POST\nPUT：PUT\nPATCH：PATCH","exampleValue":"POST,PUT,PATCH"}
  Method *string `json:"method,omitempty" xml:"method,omitempty" require:"true"`
  // {"en":"Parameter name.","zh_CN":"参数名称。"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"Parameter type.\nInteger:Integer\nNumber:Number\nString:String\nBoolean:Boolean\nEnumeration:Enumeration","zh_CN":"参数类型。\nInteger：整数\nNumber：数字\nString：字符串\nBoolean：布尔\nEnumeration：枚举","exampleValue":"Integer,Number,String,Boolean,Enumeration"}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
  // {"en":"Minimum value/Minimum length.","zh_CN":"最小值、最小长度。"}
  MinVal *string `json:"minVal,omitempty" xml:"minVal,omitempty" require:"true"`
  // {"en":"Maximum value/Maximum length.","zh_CN":"最大值、最大长度。"}
  MaxVal *string `json:"maxVal,omitempty" xml:"maxVal,omitempty" require:"true"`
  // {"en":"Required.\nTRUE:Yes\nFALSE:No","zh_CN":"必带参数。\nTRUE：是\nFALSE：否","exampleValue":"TRUE,FALSE"}
  Required *string `json:"required,omitempty" xml:"required,omitempty" require:"true"`
  // {"en":"Content.","zh_CN":"内容。"}
  Content *string `json:"content,omitempty" xml:"content,omitempty" require:"true"`
  // {"en":"Description.","zh_CN":"备注。"}
  Remark *string `json:"remark,omitempty" xml:"remark,omitempty" require:"true"`
}

func (s QueryAPIDefinitionDetailResponseDataParamLimitMethodListFormParamList) String() string {
  return tea.Prettify(s)
}

func (s QueryAPIDefinitionDetailResponseDataParamLimitMethodListFormParamList) GoString() string {
  return s.String()
}

func (s *QueryAPIDefinitionDetailResponseDataParamLimitMethodListFormParamList) SetMethod(v string) *QueryAPIDefinitionDetailResponseDataParamLimitMethodListFormParamList {
  s.Method = &v
  return s
}

func (s *QueryAPIDefinitionDetailResponseDataParamLimitMethodListFormParamList) SetName(v string) *QueryAPIDefinitionDetailResponseDataParamLimitMethodListFormParamList {
  s.Name = &v
  return s
}

func (s *QueryAPIDefinitionDetailResponseDataParamLimitMethodListFormParamList) SetType(v string) *QueryAPIDefinitionDetailResponseDataParamLimitMethodListFormParamList {
  s.Type = &v
  return s
}

func (s *QueryAPIDefinitionDetailResponseDataParamLimitMethodListFormParamList) SetMinVal(v string) *QueryAPIDefinitionDetailResponseDataParamLimitMethodListFormParamList {
  s.MinVal = &v
  return s
}

func (s *QueryAPIDefinitionDetailResponseDataParamLimitMethodListFormParamList) SetMaxVal(v string) *QueryAPIDefinitionDetailResponseDataParamLimitMethodListFormParamList {
  s.MaxVal = &v
  return s
}

func (s *QueryAPIDefinitionDetailResponseDataParamLimitMethodListFormParamList) SetRequired(v string) *QueryAPIDefinitionDetailResponseDataParamLimitMethodListFormParamList {
  s.Required = &v
  return s
}

func (s *QueryAPIDefinitionDetailResponseDataParamLimitMethodListFormParamList) SetContent(v string) *QueryAPIDefinitionDetailResponseDataParamLimitMethodListFormParamList {
  s.Content = &v
  return s
}

func (s *QueryAPIDefinitionDetailResponseDataParamLimitMethodListFormParamList) SetRemark(v string) *QueryAPIDefinitionDetailResponseDataParamLimitMethodListFormParamList {
  s.Remark = &v
  return s
}

type QueryAPIDefinitionDetailResponseDataParamLimitMethodListJsonParam struct {
  // {"en":"Request methods.\nPOST:POST\nPUT:PUT\nPATCH:PATCH","zh_CN":"请求方法。\nPOST：POST\nPUT：PUT\nPATCH：PATCH","exampleValue":"POST,PUT,PATCH"}
  Method *string `json:"method,omitempty" xml:"method,omitempty" require:"true"`
  // {"en":"Parameter Name.","zh_CN":"参数名称。"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"Level.","zh_CN":"层级。"}
  Level *int `json:"level,omitempty" xml:"level,omitempty" require:"true"`
  // {"en":"Parameter Type.\nInteger:Integer\nNumber:Number\nString:String\nBoolean:Boolean\nEnumeration:Enumeration\nArray:Array\nJson:JSON Object","zh_CN":"参数类型。\nInteger：整数\nNumber：数字\nString：字符串\nBoolean：布尔\nEnumeration：枚举\nArray：数组\nJson：JSON对象","exampleValue":"Integer,Number,String,Boolean,Enumeration,Array,Json"}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
  // {"en":"Minimum value/Minimum length.","zh_CN":"最小值、最小长度。"}
  MinVal *string `json:"minVal,omitempty" xml:"minVal,omitempty" require:"true"`
  // {"en":"Maximum value/Maximum length.","zh_CN":"最大值、最大长度。"}
  MaxVal *string `json:"maxVal,omitempty" xml:"maxVal,omitempty" require:"true"`
  // {"en":"Required.\nTRUE:Yes\nFALSE:No","zh_CN":"必带参数。\nTRUE：是\nFALSE：否","exampleValue":"TRUE,FALSE"}
  Required *string `json:"required,omitempty" xml:"required,omitempty" require:"true"`
  // {"en":"Content.","zh_CN":"内容。"}
  Content *string `json:"content,omitempty" xml:"content,omitempty" require:"true"`
  // {"en":"Child node array, required when type= Array.","zh_CN":"子节点数组，type = Array 时必填。"}
  Children []*QueryAPIDefinitionDetailResponseDataParamLimitMethodListJsonParamChildren `json:"children,omitempty" xml:"children,omitempty" require:"true" type:"Repeated"`
}

func (s QueryAPIDefinitionDetailResponseDataParamLimitMethodListJsonParam) String() string {
  return tea.Prettify(s)
}

func (s QueryAPIDefinitionDetailResponseDataParamLimitMethodListJsonParam) GoString() string {
  return s.String()
}

func (s *QueryAPIDefinitionDetailResponseDataParamLimitMethodListJsonParam) SetMethod(v string) *QueryAPIDefinitionDetailResponseDataParamLimitMethodListJsonParam {
  s.Method = &v
  return s
}

func (s *QueryAPIDefinitionDetailResponseDataParamLimitMethodListJsonParam) SetName(v string) *QueryAPIDefinitionDetailResponseDataParamLimitMethodListJsonParam {
  s.Name = &v
  return s
}

func (s *QueryAPIDefinitionDetailResponseDataParamLimitMethodListJsonParam) SetLevel(v int) *QueryAPIDefinitionDetailResponseDataParamLimitMethodListJsonParam {
  s.Level = &v
  return s
}

func (s *QueryAPIDefinitionDetailResponseDataParamLimitMethodListJsonParam) SetType(v string) *QueryAPIDefinitionDetailResponseDataParamLimitMethodListJsonParam {
  s.Type = &v
  return s
}

func (s *QueryAPIDefinitionDetailResponseDataParamLimitMethodListJsonParam) SetMinVal(v string) *QueryAPIDefinitionDetailResponseDataParamLimitMethodListJsonParam {
  s.MinVal = &v
  return s
}

func (s *QueryAPIDefinitionDetailResponseDataParamLimitMethodListJsonParam) SetMaxVal(v string) *QueryAPIDefinitionDetailResponseDataParamLimitMethodListJsonParam {
  s.MaxVal = &v
  return s
}

func (s *QueryAPIDefinitionDetailResponseDataParamLimitMethodListJsonParam) SetRequired(v string) *QueryAPIDefinitionDetailResponseDataParamLimitMethodListJsonParam {
  s.Required = &v
  return s
}

func (s *QueryAPIDefinitionDetailResponseDataParamLimitMethodListJsonParam) SetContent(v string) *QueryAPIDefinitionDetailResponseDataParamLimitMethodListJsonParam {
  s.Content = &v
  return s
}

func (s *QueryAPIDefinitionDetailResponseDataParamLimitMethodListJsonParam) SetChildren(v []*QueryAPIDefinitionDetailResponseDataParamLimitMethodListJsonParamChildren) *QueryAPIDefinitionDetailResponseDataParamLimitMethodListJsonParam {
  s.Children = v
  return s
}

type QueryAPIDefinitionDetailResponseDataParamLimitMethodListJsonParamChildren struct     {
  // {"en":"Request methods.\nPOST:POST\nPUT:PUT\nPATCH:PATCH","zh_CN":"请求方法。\nPOST：POST\nPUT：PUT\nPATCH：PATCH","exampleValue":"POST,PUT,PATCH"}
  Method *string `json:"method,omitempty" xml:"method,omitempty" require:"true"`
  // {"en":"Parameter Name.","zh_CN":"参数名称。"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"Level.","zh_CN":"层级。"}
  Level *int `json:"level,omitempty" xml:"level,omitempty" require:"true"`
  // {"en":"Parameter Type.\nInteger:Integer\nNumber:Number\nString:String\nBoolean:Boolean\nEnumeration:Enumeration\nArray:Array\nJson:JSON Object","zh_CN":"参数类型。\nInteger：整数\nNumber：数字\nString：字符串\nBoolean：布尔\nEnumeration：枚举\nArray：数组\nJson：JSON对象","exampleValue":"Integer,Number,String,Boolean,Enumeration,Array,Json"}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
  // {"en":"Minimum value/Minimum length.","zh_CN":"最小值、最小长度。"}
  MinVal *string `json:"minVal,omitempty" xml:"minVal,omitempty" require:"true"`
  // {"en":"Maximum value/Maximum length.","zh_CN":"最大值、最大长度。"}
  MaxVal *string `json:"maxVal,omitempty" xml:"maxVal,omitempty" require:"true"`
  // {"en":"Required.\nTRUE:Yes\nFALSE:No","zh_CN":"必带参数。\nTRUE：是\nFALSE：否","exampleValue":"TRUE,FALSE"}
  Required *string `json:"required,omitempty" xml:"required,omitempty" require:"true"`
  // {"en":"Content.","zh_CN":"内容。"}
  Content *string `json:"content,omitempty" xml:"content,omitempty" require:"true"`
  // {"en":"Child node array (JSON string), the structure is consistent with the parent node, required when type= Array.","zh_CN":"子节点数组（JSON字符串），结构与父结点一致，type = Array 时必填。"}
  Children []*string `json:"children,omitempty" xml:"children,omitempty" require:"true" type:"Repeated"`
}

func (s QueryAPIDefinitionDetailResponseDataParamLimitMethodListJsonParamChildren) String() string {
  return tea.Prettify(s)
}

func (s QueryAPIDefinitionDetailResponseDataParamLimitMethodListJsonParamChildren) GoString() string {
  return s.String()
}

func (s *QueryAPIDefinitionDetailResponseDataParamLimitMethodListJsonParamChildren) SetMethod(v string) *QueryAPIDefinitionDetailResponseDataParamLimitMethodListJsonParamChildren {
  s.Method = &v
  return s
}

func (s *QueryAPIDefinitionDetailResponseDataParamLimitMethodListJsonParamChildren) SetName(v string) *QueryAPIDefinitionDetailResponseDataParamLimitMethodListJsonParamChildren {
  s.Name = &v
  return s
}

func (s *QueryAPIDefinitionDetailResponseDataParamLimitMethodListJsonParamChildren) SetLevel(v int) *QueryAPIDefinitionDetailResponseDataParamLimitMethodListJsonParamChildren {
  s.Level = &v
  return s
}

func (s *QueryAPIDefinitionDetailResponseDataParamLimitMethodListJsonParamChildren) SetType(v string) *QueryAPIDefinitionDetailResponseDataParamLimitMethodListJsonParamChildren {
  s.Type = &v
  return s
}

func (s *QueryAPIDefinitionDetailResponseDataParamLimitMethodListJsonParamChildren) SetMinVal(v string) *QueryAPIDefinitionDetailResponseDataParamLimitMethodListJsonParamChildren {
  s.MinVal = &v
  return s
}

func (s *QueryAPIDefinitionDetailResponseDataParamLimitMethodListJsonParamChildren) SetMaxVal(v string) *QueryAPIDefinitionDetailResponseDataParamLimitMethodListJsonParamChildren {
  s.MaxVal = &v
  return s
}

func (s *QueryAPIDefinitionDetailResponseDataParamLimitMethodListJsonParamChildren) SetRequired(v string) *QueryAPIDefinitionDetailResponseDataParamLimitMethodListJsonParamChildren {
  s.Required = &v
  return s
}

func (s *QueryAPIDefinitionDetailResponseDataParamLimitMethodListJsonParamChildren) SetContent(v string) *QueryAPIDefinitionDetailResponseDataParamLimitMethodListJsonParamChildren {
  s.Content = &v
  return s
}

func (s *QueryAPIDefinitionDetailResponseDataParamLimitMethodListJsonParamChildren) SetChildren(v []*string) *QueryAPIDefinitionDetailResponseDataParamLimitMethodListJsonParamChildren {
  s.Children = v
  return s
}

type QueryAPIDefinitionDetailResponseHeader struct {
}

func (s QueryAPIDefinitionDetailResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryAPIDefinitionDetailResponseHeader) GoString() string {
  return s.String()
}




type FeedbackWrongAPIAssetDiscoveryRequest struct {
  // {"en":"API discovery ID.","zh_CN":"API发现ID。"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {"en":"Feedback, maximum 200 characters.","zh_CN":"反馈意见，最多200个字符。"}
  Feedback *string `json:"feedback,omitempty" xml:"feedback,omitempty"`
}

func (s FeedbackWrongAPIAssetDiscoveryRequest) String() string {
  return tea.Prettify(s)
}

func (s FeedbackWrongAPIAssetDiscoveryRequest) GoString() string {
  return s.String()
}

func (s *FeedbackWrongAPIAssetDiscoveryRequest) SetId(v string) *FeedbackWrongAPIAssetDiscoveryRequest {
  s.Id = &v
  return s
}

func (s *FeedbackWrongAPIAssetDiscoveryRequest) SetFeedback(v string) *FeedbackWrongAPIAssetDiscoveryRequest {
  s.Feedback = &v
  return s
}

type FeedbackWrongAPIAssetDiscoveryRequestHeader struct {
  // {"en":"Security service type. Please enter a specific service type, if you purchase multiple security services.","zh_CN":"安全服务类型。有使用多个不同的安全服务时，需要填写具体的服务类型。"}
  ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty"`
}

func (s FeedbackWrongAPIAssetDiscoveryRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s FeedbackWrongAPIAssetDiscoveryRequestHeader) GoString() string {
  return s.String()
}

func (s *FeedbackWrongAPIAssetDiscoveryRequestHeader) SetServiceType(v string) *FeedbackWrongAPIAssetDiscoveryRequestHeader {
  s.ServiceType = &v
  return s
}

type FeedbackWrongAPIAssetDiscoveryPaths struct {
}

func (s FeedbackWrongAPIAssetDiscoveryPaths) String() string {
  return tea.Prettify(s)
}

func (s FeedbackWrongAPIAssetDiscoveryPaths) GoString() string {
  return s.String()
}

type FeedbackWrongAPIAssetDiscoveryParameters struct {
}

func (s FeedbackWrongAPIAssetDiscoveryParameters) String() string {
  return tea.Prettify(s)
}

func (s FeedbackWrongAPIAssetDiscoveryParameters) GoString() string {
  return s.String()
}

type FeedbackWrongAPIAssetDiscoveryResponse struct {
  // {"dictionary":"belong=WAAP-MS-Ext|dict=waap_retCodeEnum","en":"Please refer to the error code for exceptions.","zh_CN":"请参照错误码。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Description.","zh_CN":"描述信息。"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
}

func (s FeedbackWrongAPIAssetDiscoveryResponse) String() string {
  return tea.Prettify(s)
}

func (s FeedbackWrongAPIAssetDiscoveryResponse) GoString() string {
  return s.String()
}

func (s *FeedbackWrongAPIAssetDiscoveryResponse) SetCode(v string) *FeedbackWrongAPIAssetDiscoveryResponse {
  s.Code = &v
  return s
}

func (s *FeedbackWrongAPIAssetDiscoveryResponse) SetMsg(v string) *FeedbackWrongAPIAssetDiscoveryResponse {
  s.Msg = &v
  return s
}

type FeedbackWrongAPIAssetDiscoveryResponseHeader struct {
}

func (s FeedbackWrongAPIAssetDiscoveryResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s FeedbackWrongAPIAssetDiscoveryResponseHeader) GoString() string {
  return s.String()
}




type UpdateAPIDefinitionRequest struct {
  // {"en":"Basic information.","zh_CN":"基础信息。"}
  Basic *UpdateAPIDefinitionRequestBasic `json:"basic,omitempty" xml:"basic,omitempty" require:"true" type:"Struct"`
  // {"en":"Endpoint information.","zh_CN":"端点信息。"}
  Endpoint *UpdateAPIDefinitionRequestEndpoint `json:"endpoint,omitempty" xml:"endpoint,omitempty" type:"Struct"`
  // {"en":"Authentication configuration.","zh_CN":"鉴权配置。"}
  Auth *UpdateAPIDefinitionRequestAuth `json:"auth,omitempty" xml:"auth,omitempty" type:"Struct"`
  // {"en":"Body restrictions.","zh_CN":"body限制。"}
  BodyLimit *UpdateAPIDefinitionRequestBodyLimit `json:"bodyLimit,omitempty" xml:"bodyLimit,omitempty" type:"Struct"`
  // {"en":"Parameter limit.","zh_CN":"参数限制。"}
  ParamLimit *UpdateAPIDefinitionRequestParamLimit `json:"paramLimit,omitempty" xml:"paramLimit,omitempty" type:"Struct"`
}

func (s UpdateAPIDefinitionRequest) String() string {
  return tea.Prettify(s)
}

func (s UpdateAPIDefinitionRequest) GoString() string {
  return s.String()
}

func (s *UpdateAPIDefinitionRequest) SetBasic(v *UpdateAPIDefinitionRequestBasic) *UpdateAPIDefinitionRequest {
  s.Basic = v
  return s
}

func (s *UpdateAPIDefinitionRequest) SetEndpoint(v *UpdateAPIDefinitionRequestEndpoint) *UpdateAPIDefinitionRequest {
  s.Endpoint = v
  return s
}

func (s *UpdateAPIDefinitionRequest) SetAuth(v *UpdateAPIDefinitionRequestAuth) *UpdateAPIDefinitionRequest {
  s.Auth = v
  return s
}

func (s *UpdateAPIDefinitionRequest) SetBodyLimit(v *UpdateAPIDefinitionRequestBodyLimit) *UpdateAPIDefinitionRequest {
  s.BodyLimit = v
  return s
}

func (s *UpdateAPIDefinitionRequest) SetParamLimit(v *UpdateAPIDefinitionRequestParamLimit) *UpdateAPIDefinitionRequest {
  s.ParamLimit = v
  return s
}

type UpdateAPIDefinitionRequestBasic struct {
  // {"en":"API define ID.","zh_CN":"API定义ID。"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {"en":"API name, maximum 200 characters.","zh_CN":"API名称，最多200个字符。"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"API groups, maximum 64 characters.","zh_CN":"API分组，最多64个字符。"}
  GroupName *string `json:"groupName,omitempty" xml:"groupName,omitempty" require:"true"`
  // {"en":"Description, maximum 200 characters.","zh_CN":"备注，最多200个字符。"}
  Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
}

func (s UpdateAPIDefinitionRequestBasic) String() string {
  return tea.Prettify(s)
}

func (s UpdateAPIDefinitionRequestBasic) GoString() string {
  return s.String()
}

func (s *UpdateAPIDefinitionRequestBasic) SetId(v string) *UpdateAPIDefinitionRequestBasic {
  s.Id = &v
  return s
}

func (s *UpdateAPIDefinitionRequestBasic) SetName(v string) *UpdateAPIDefinitionRequestBasic {
  s.Name = &v
  return s
}

func (s *UpdateAPIDefinitionRequestBasic) SetGroupName(v string) *UpdateAPIDefinitionRequestBasic {
  s.GroupName = &v
  return s
}

func (s *UpdateAPIDefinitionRequestBasic) SetRemark(v string) *UpdateAPIDefinitionRequestBasic {
  s.Remark = &v
  return s
}

type UpdateAPIDefinitionRequestEndpoint struct {
  // {"en":"Request methods,Multiple separated by ; sign.\nGET:GET, configurable parameter limits\nPOST:POST, configurable parameter limits\nDELETE:DELETE, configurable parameter limits\nUPDATE:UPDATE\nPUT:PUT, configurable parameter limits\nHEAD:HEAD, configurable parameter limits\nCONNECT:CONNECT\nOPTIONS:OPTIONS, configurable parameter limits\nCOPY:COPY\nLOCK:LOCK\nUNLOCK:UNLOCK\nTRACE:TRACE\nPATCH:PATCH, configurable parameter limits\nPROPFIND:PROPFIND\nMKCOL:MKCOL\nMOVE:MOVE","zh_CN":"请求方法。多个以 ; 号分隔。\nGET：GET，可配置参数限制\nPOST：POST，可配置参数限制\nDELETE：DELETE，可配置参数限制\nUPDATE：UPDATE\nPUT：PUT，可配置参数限制\nHEAD：HEAD，可配置参数限制\nCONNECT：CONNECT\nOPTIONS：OPTIONS，可配置参数限制\nCOPY：COPY\nLOCK：LOCK\nUNLOCK：UNLOCK\nTRACE：TRACE\nPATCH：PATCH，可配置参数限制\nPROPFIND：PROPFIND\nMKCOL：MKCOL\nMOVE：MOVE","exampleValue":"GET,POST,DELETE,UPDATE,PUT,HEAD,CONNECT,OPTIONS,COPY,LOCK,UNLOCK,TRACE,PATCH,PROPFIND,MKCOL,MOVE"}
  RequestMethod *string `json:"requestMethod,omitempty" xml:"requestMethod,omitempty" require:"true"`
  // {"en":"Case sensitive.\nTRUE:Yes\nFALSE:No","zh_CN":"大小写是否敏感。\nTRUE：是\nFALSE：否","exampleValue":"TRUE,FALSE"}
  CaseSensitive *string `json:"caseSensitive,omitempty" xml:"caseSensitive,omitempty" require:"true"`
}

func (s UpdateAPIDefinitionRequestEndpoint) String() string {
  return tea.Prettify(s)
}

func (s UpdateAPIDefinitionRequestEndpoint) GoString() string {
  return s.String()
}

func (s *UpdateAPIDefinitionRequestEndpoint) SetRequestMethod(v string) *UpdateAPIDefinitionRequestEndpoint {
  s.RequestMethod = &v
  return s
}

func (s *UpdateAPIDefinitionRequestEndpoint) SetCaseSensitive(v string) *UpdateAPIDefinitionRequestEndpoint {
  s.CaseSensitive = &v
  return s
}

type UpdateAPIDefinitionRequestAuth struct {
  // {"en":"Authentication method.\nNO_NEED:No authentication required\nSIGN:Key authentication","zh_CN":"鉴权方法。\nNO_NEED：免鉴权\nSIGN：秘钥对鉴权","exampleValue":"NO_NEED,SIGN"}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
  // {"en":"Authentication key, type = SIGN is required. If it has been set, it will be ignored. The format is a 16-digit string containing uppercase and lowercase letters and numbers. Example: gjZkg2E1uNkXBDxj.","zh_CN":"鉴权秘钥，type = SIGN 是必填，如已设置则忽略，格式为16位含大小写字母与数字字符串，示例：gjZkg2E1uNkXBDxj。"}
  AuthKey *string `json:"authKey,omitempty" xml:"authKey,omitempty"`
  // {"en":"Authentication parameter location, type = SIGN is required.\nHTTP_HEADER:Http Header\nQUERY_STRING:Query String\nCOOKIE:Cookie","zh_CN":"鉴权参数位置，type = SIGN 时必填。\nHTTP_HEADER：Http Header\nQUERY_STRING：Query String\nCOOKIE：Cookie","exampleValue":"HTTP_HEADER,QUERY_STRING,COOKIE"}
  ParamPosition *string `json:"paramPosition,omitempty" xml:"paramPosition,omitempty"`
  // {"en":"Authentication parameter name, type = SIGN is required.","zh_CN":"鉴权参数名称，type = SIGN 时必填。"}
  ParamName *string `json:"paramName,omitempty" xml:"paramName,omitempty"`
  // {"en":"Authentication token validity period, in seconds, type = SIGN is required.","zh_CN":"鉴权有效期，单位秒，type = SIGN 时必填。"}
  ValidityTime *int64 `json:"validityTime,omitempty" xml:"validityTime,omitempty"`
}

func (s UpdateAPIDefinitionRequestAuth) String() string {
  return tea.Prettify(s)
}

func (s UpdateAPIDefinitionRequestAuth) GoString() string {
  return s.String()
}

func (s *UpdateAPIDefinitionRequestAuth) SetType(v string) *UpdateAPIDefinitionRequestAuth {
  s.Type = &v
  return s
}

func (s *UpdateAPIDefinitionRequestAuth) SetAuthKey(v string) *UpdateAPIDefinitionRequestAuth {
  s.AuthKey = &v
  return s
}

func (s *UpdateAPIDefinitionRequestAuth) SetParamPosition(v string) *UpdateAPIDefinitionRequestAuth {
  s.ParamPosition = &v
  return s
}

func (s *UpdateAPIDefinitionRequestAuth) SetParamName(v string) *UpdateAPIDefinitionRequestAuth {
  s.ParamName = &v
  return s
}

func (s *UpdateAPIDefinitionRequestAuth) SetValidityTime(v int64) *UpdateAPIDefinitionRequestAuth {
  s.ValidityTime = &v
  return s
}

type UpdateAPIDefinitionRequestBodyLimit struct {
  // {"defaultValue":"OFF","en":"Request body restriction switch. default value:OFF.\nON:On\nOFF:Off","zh_CN":"请求body限制开关。默认值：关。\nON：开启\nOFF：关闭","exampleValue":"ON,OFF"}
  DefendSwitch *string `json:"defendSwitch,omitempty" xml:"defendSwitch,omitempty"`
  // {"en":"Content-Type, required when defendSwitch = ON.\nFORM:FORM\nJSON:JSON\nANY:ANY\nEMPTY:EMPTY or NON-EXISTENT","zh_CN":"Content-Type，defendSwitch = ON 时必填。\nFORM：FORM表单\nJSON：JSON\nANY：任意\nEMPTY：为空或不存在","exampleValue":"FORM,JSON,ANY,EMPTY"}
  ContentType *string `json:"contentType,omitempty" xml:"contentType,omitempty"`
  // {"en":"Maximum body limit(bytes),does not exceed 16,384.","zh_CN":"Body最大限制(bytes),最大不超过16384。"}
  BodyLimitMax *int64 `json:"bodyLimitMax,omitempty" xml:"bodyLimitMax,omitempty"`
  // {"en":"Maximum nesting depth, enter the maximum allowable JSON nesting depth in the request body.","zh_CN":"最大嵌套层数，输入允许的请求body中JSON嵌套层数最大值。"}
  NestMax *int64 `json:"nestMax,omitempty" xml:"nestMax,omitempty"`
  // {"en":"Maximum number of parameters for JSON, enter the maximum number of JSON parameters allowed in the request body.","zh_CN":"JSON最大参数个数，输入允许的请求body中JSON参数个数的最大值。"}
  ParamCountMax *int64 `json:"paramCountMax,omitempty" xml:"paramCountMax,omitempty"`
}

func (s UpdateAPIDefinitionRequestBodyLimit) String() string {
  return tea.Prettify(s)
}

func (s UpdateAPIDefinitionRequestBodyLimit) GoString() string {
  return s.String()
}

func (s *UpdateAPIDefinitionRequestBodyLimit) SetDefendSwitch(v string) *UpdateAPIDefinitionRequestBodyLimit {
  s.DefendSwitch = &v
  return s
}

func (s *UpdateAPIDefinitionRequestBodyLimit) SetContentType(v string) *UpdateAPIDefinitionRequestBodyLimit {
  s.ContentType = &v
  return s
}

func (s *UpdateAPIDefinitionRequestBodyLimit) SetBodyLimitMax(v int64) *UpdateAPIDefinitionRequestBodyLimit {
  s.BodyLimitMax = &v
  return s
}

func (s *UpdateAPIDefinitionRequestBodyLimit) SetNestMax(v int64) *UpdateAPIDefinitionRequestBodyLimit {
  s.NestMax = &v
  return s
}

func (s *UpdateAPIDefinitionRequestBodyLimit) SetParamCountMax(v int64) *UpdateAPIDefinitionRequestBodyLimit {
  s.ParamCountMax = &v
  return s
}

type UpdateAPIDefinitionRequestParamLimit struct {
  // {"en":"Basic information.","zh_CN":"基础信息。"}
  Basic *UpdateAPIDefinitionRequestParamLimitBasic `json:"basic,omitempty" xml:"basic,omitempty" require:"true" type:"Struct"`
  // {"en":"There are request methods for changing configurations.\nGET:GET\nPOST:POST\nDELETE:DELETE\nPUT:PUT\nHEAD:HEAD\nOPTIONS:OPTIONS\nPATCH:PATCH","zh_CN":"有变更配置的请求方法。\nGET：GET\nPOST：POST\nDELETE：DELETE\nPUT：PUT\nHEAD：HEAD\nOPTIONS：OPTIONS\nPATCH：PATCH","exampleValue":"GET,POST,DELETE,PUT,HEAD,OPTIONS,PATCH"}
  ChangeParamLimitMethodList []*string `json:"changeParamLimitMethodList,omitempty" xml:"changeParamLimitMethodList,omitempty" type:"Repeated"`
  // {"en":"Method attributes.","zh_CN":"方法属性。"}
  MethodList []*UpdateAPIDefinitionRequestParamLimitMethodList `json:"methodList,omitempty" xml:"methodList,omitempty" type:"Repeated"`
}

func (s UpdateAPIDefinitionRequestParamLimit) String() string {
  return tea.Prettify(s)
}

func (s UpdateAPIDefinitionRequestParamLimit) GoString() string {
  return s.String()
}

func (s *UpdateAPIDefinitionRequestParamLimit) SetBasic(v *UpdateAPIDefinitionRequestParamLimitBasic) *UpdateAPIDefinitionRequestParamLimit {
  s.Basic = v
  return s
}

func (s *UpdateAPIDefinitionRequestParamLimit) SetChangeParamLimitMethodList(v []*string) *UpdateAPIDefinitionRequestParamLimit {
  s.ChangeParamLimitMethodList = v
  return s
}

func (s *UpdateAPIDefinitionRequestParamLimit) SetMethodList(v []*UpdateAPIDefinitionRequestParamLimitMethodList) *UpdateAPIDefinitionRequestParamLimit {
  s.MethodList = v
  return s
}

type UpdateAPIDefinitionRequestParamLimitBasic struct {
  // {"en":"Parameter limit.\nON:On\nOFF:Off","zh_CN":"参数限制。\nON：开启\nOFF：关闭","exampleValue":"ON,OFF"}
  DefendSwitch *string `json:"defendSwitch,omitempty" xml:"defendSwitch,omitempty"`
  // {"en":"Query String Parameter Detection Mode, required when defendSwitch = ON.\nLOOSE:Loose mode - detect some parameters\nSTRICT:Strict mode - checks all parameters","zh_CN":"Query String参数检测模式，defendSwitch = ON 时必填。\nLOOSE：宽松模式-检测部分参数\nSTRICT：严格模式-检测所有参数","exampleValue":"LOOSE,STRICT"}
  ParamCheckMode *string `json:"paramCheckMode,omitempty" xml:"paramCheckMode,omitempty"`
}

func (s UpdateAPIDefinitionRequestParamLimitBasic) String() string {
  return tea.Prettify(s)
}

func (s UpdateAPIDefinitionRequestParamLimitBasic) GoString() string {
  return s.String()
}

func (s *UpdateAPIDefinitionRequestParamLimitBasic) SetDefendSwitch(v string) *UpdateAPIDefinitionRequestParamLimitBasic {
  s.DefendSwitch = &v
  return s
}

func (s *UpdateAPIDefinitionRequestParamLimitBasic) SetParamCheckMode(v string) *UpdateAPIDefinitionRequestParamLimitBasic {
  s.ParamCheckMode = &v
  return s
}

type UpdateAPIDefinitionRequestParamLimitMethodList struct     {
  // {"en":"Request methods.\nGET:GET\nPOST:POST\nDELETE:DELETE\nPUT:PUT\nHEAD:HEAD\nOPTIONS:OPTIONS\nPATCH:PATCH","zh_CN":"请求方法。\nGET：GET\nPOST：POST\nDELETE：DELETE\nPUT：PUT\nHEAD：HEAD\nOPTIONS：OPTIONS\nPATCH：PATCH","exampleValue":"GET,POST,DELETE,PUT,HEAD,OPTIONS,PATCH"}
  Method *string `json:"method,omitempty" xml:"method,omitempty"`
  // {"en":"Whether to define body parameters, required when method = POST/PUT/PATCH.\nTRUE:Yes\nFALSE:No","zh_CN":"是否定义body参数，method = POST/PUT/PATCH时必填。\nTRUE：是\nFALSE：否","exampleValue":"TRUE,FALSE"}
  BodyFlag *string `json:"bodyFlag,omitempty" xml:"bodyFlag,omitempty"`
  // {"en":"Body parameter type, required when bodyFlag = TRUE.\nJSON:JSON\nFROM:FROM","zh_CN":"body参数类型，bodyFlag = TRUE 时必填。\nJSON：JSON\nFROM：FROM表单","exampleValue":"FORM,JSON"}
  BodyType *string `json:"bodyType,omitempty" xml:"bodyType,omitempty"`
  // {"en":"Normal parameter list.","zh_CN":"普通参数数组。"}
  NormalParamList []*UpdateAPIDefinitionRequestParamLimitMethodListNormalParamList `json:"normalParamList,omitempty" xml:"normalParamList,omitempty" type:"Repeated"`
  // {"en":"FROM form parameter array, Optional when bodyType = FROM.","zh_CN":"FROM 表单参数数组，bodyType = FROM 时选填。"}
  FormParamList []*UpdateAPIDefinitionRequestParamLimitMethodListFormParamList `json:"formParamList,omitempty" xml:"formParamList,omitempty" type:"Repeated"`
  // {"en":"JSON parameter array, Optional when bodyType = JSON.","zh_CN":"JSON参数数组，bodyType = JSON 时选填。"}
  JsonParam *UpdateAPIDefinitionRequestParamLimitMethodListJsonParam `json:"jsonParam,omitempty" xml:"jsonParam,omitempty" type:"Struct"`
}

func (s UpdateAPIDefinitionRequestParamLimitMethodList) String() string {
  return tea.Prettify(s)
}

func (s UpdateAPIDefinitionRequestParamLimitMethodList) GoString() string {
  return s.String()
}

func (s *UpdateAPIDefinitionRequestParamLimitMethodList) SetMethod(v string) *UpdateAPIDefinitionRequestParamLimitMethodList {
  s.Method = &v
  return s
}

func (s *UpdateAPIDefinitionRequestParamLimitMethodList) SetBodyFlag(v string) *UpdateAPIDefinitionRequestParamLimitMethodList {
  s.BodyFlag = &v
  return s
}

func (s *UpdateAPIDefinitionRequestParamLimitMethodList) SetBodyType(v string) *UpdateAPIDefinitionRequestParamLimitMethodList {
  s.BodyType = &v
  return s
}

func (s *UpdateAPIDefinitionRequestParamLimitMethodList) SetNormalParamList(v []*UpdateAPIDefinitionRequestParamLimitMethodListNormalParamList) *UpdateAPIDefinitionRequestParamLimitMethodList {
  s.NormalParamList = v
  return s
}

func (s *UpdateAPIDefinitionRequestParamLimitMethodList) SetFormParamList(v []*UpdateAPIDefinitionRequestParamLimitMethodListFormParamList) *UpdateAPIDefinitionRequestParamLimitMethodList {
  s.FormParamList = v
  return s
}

func (s *UpdateAPIDefinitionRequestParamLimitMethodList) SetJsonParam(v *UpdateAPIDefinitionRequestParamLimitMethodListJsonParam) *UpdateAPIDefinitionRequestParamLimitMethodList {
  s.JsonParam = v
  return s
}

type UpdateAPIDefinitionRequestParamLimitMethodListNormalParamList struct     {
  // {"en":"Parameter name, when paramPosition = PATH_PARAMS, the path need to be matched, for example: /basePath/{name}/.","zh_CN":"参数名称，paramPosition = PATH_PARAMS 时需匹配路径变量，例如：/basePath/{name}/。"}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // {"en":"Parameter position.\nHTTP_HEADER:Http Header\nQUERY_STRING:Query String\nCOOKIE:Cookie\nPATH_PARAMS:Path, endpoint.matchPathVar needs to equal TRUE","zh_CN":"参数位置。\nHTTP_HEADER：Http Header\nQUERY_STRING：Query String\nCOOKIE：Cookie\nPATH_PARAMS：路径变量，endpoint.matchPathVar需等于TRUE","exampleValue":"HTTP_HEADER,QUERY_STRING,COOKIE,PATH_PARAMS"}
  ParamPosition *string `json:"paramPosition,omitempty" xml:"paramPosition,omitempty"`
  // {"en":"Parameter type.\nInteger:Integer\nNumber:Number\nString:String\nBoolean:Boolean\nEnumeration:Enumeration","zh_CN":"参数类型。\nInteger：整数\nNumber：数字\nString：字符串\nBoolean：布尔\nEnumeration：枚举","exampleValue":"Integer,Number,String,Boolean,Enumeration"}
  Type *string `json:"type,omitempty" xml:"type,omitempty"`
  // {"en":"type = Integer/Number:minimum value, type = String:minimum length, type = Boolean/Enumeration:leave blank.","zh_CN":"type = Integer/Number：最小值，type = String：最小长度，type = Boolean/Enumeration：置空。"}
  MinVal *string `json:"minVal,omitempty" xml:"minVal,omitempty"`
  // {"en":"type = Integer/Number:maximum value, type = String:maximum length, type = Boolean/Enumeration:leave blank.","zh_CN":"type = Integer/Number：最大值，type = String：最大长度，type = Boolean/Enumeration：置空。"}
  MaxVal *string `json:"maxVal,omitempty" xml:"maxVal,omitempty"`
  // {"en":"Required.\nTRUE:Yes\nFALSE:No","zh_CN":"必带参数。\nTRUE：是\nFALSE：否","exampleValue":"TRUE,FALSE"}
  Required *string `json:"required,omitempty" xml:"required,omitempty"`
  // {"en":"Content (maximum 2000 characters), required when type = Enumeration, Multiple separated by ; sign.","zh_CN":"内容（最多2000个字符），type = Enumeration 时必填，多个以 ; 号分隔。"}
  Content *string `json:"content,omitempty" xml:"content,omitempty"`
  // {"en":"Description.","zh_CN":"备注。"}
  Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
}

func (s UpdateAPIDefinitionRequestParamLimitMethodListNormalParamList) String() string {
  return tea.Prettify(s)
}

func (s UpdateAPIDefinitionRequestParamLimitMethodListNormalParamList) GoString() string {
  return s.String()
}

func (s *UpdateAPIDefinitionRequestParamLimitMethodListNormalParamList) SetName(v string) *UpdateAPIDefinitionRequestParamLimitMethodListNormalParamList {
  s.Name = &v
  return s
}

func (s *UpdateAPIDefinitionRequestParamLimitMethodListNormalParamList) SetParamPosition(v string) *UpdateAPIDefinitionRequestParamLimitMethodListNormalParamList {
  s.ParamPosition = &v
  return s
}

func (s *UpdateAPIDefinitionRequestParamLimitMethodListNormalParamList) SetType(v string) *UpdateAPIDefinitionRequestParamLimitMethodListNormalParamList {
  s.Type = &v
  return s
}

func (s *UpdateAPIDefinitionRequestParamLimitMethodListNormalParamList) SetMinVal(v string) *UpdateAPIDefinitionRequestParamLimitMethodListNormalParamList {
  s.MinVal = &v
  return s
}

func (s *UpdateAPIDefinitionRequestParamLimitMethodListNormalParamList) SetMaxVal(v string) *UpdateAPIDefinitionRequestParamLimitMethodListNormalParamList {
  s.MaxVal = &v
  return s
}

func (s *UpdateAPIDefinitionRequestParamLimitMethodListNormalParamList) SetRequired(v string) *UpdateAPIDefinitionRequestParamLimitMethodListNormalParamList {
  s.Required = &v
  return s
}

func (s *UpdateAPIDefinitionRequestParamLimitMethodListNormalParamList) SetContent(v string) *UpdateAPIDefinitionRequestParamLimitMethodListNormalParamList {
  s.Content = &v
  return s
}

func (s *UpdateAPIDefinitionRequestParamLimitMethodListNormalParamList) SetRemark(v string) *UpdateAPIDefinitionRequestParamLimitMethodListNormalParamList {
  s.Remark = &v
  return s
}

type UpdateAPIDefinitionRequestParamLimitMethodListFormParamList struct     {
  // {"en":"Parameter name.","zh_CN":"参数名称。"}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // {"en":"Parameter type.\nInteger:Integer\nNumber:Number\nString:String\nBoolean:Boolean\nEnumeration:Enumeration","zh_CN":"参数类型。\nInteger：整数\nNumber：数字\nString：字符串\nBoolean：布尔\nEnumeration：枚举","exampleValue":"Integer,Number,String,Boolean,Enumeration"}
  Type *string `json:"type,omitempty" xml:"type,omitempty"`
  // {"en":"type = Integer/Number:minimum value, type = String:minimum length, type = Boolean/Enumeration:leave blank.","zh_CN":"type = Integer/Number：最小值，type = String：最小长度，type = Boolean/Enumeration：置空。"}
  MinVal *string `json:"minVal,omitempty" xml:"minVal,omitempty"`
  // {"en":"type = Integer/Number:maximum value, type = String:maximum length, type = Boolean/Enumeration:leave blank.","zh_CN":"type = Integer/Number：最大值，type = String：最大长度，type = Boolean/Enumeration：置空。"}
  MaxVal *string `json:"maxVal,omitempty" xml:"maxVal,omitempty"`
  // {"en":"Required.\nTRUE:Yes\nFALSE:No","zh_CN":"必带参数。\nTRUE：是\nFALSE：否","exampleValue":"TRUE,FALSE"}
  Required *string `json:"required,omitempty" xml:"required,omitempty"`
  // {"en":"Content (maximum 2000 characters), required when type = Enumeration, multiple separated by ; sign.","zh_CN":"内容（最多2000个字符），type = Enumeration 时必填，多个以 ; 号分隔。"}
  Content *string `json:"content,omitempty" xml:"content,omitempty"`
  // {"en":"Description.","zh_CN":"备注。"}
  Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
}

func (s UpdateAPIDefinitionRequestParamLimitMethodListFormParamList) String() string {
  return tea.Prettify(s)
}

func (s UpdateAPIDefinitionRequestParamLimitMethodListFormParamList) GoString() string {
  return s.String()
}

func (s *UpdateAPIDefinitionRequestParamLimitMethodListFormParamList) SetName(v string) *UpdateAPIDefinitionRequestParamLimitMethodListFormParamList {
  s.Name = &v
  return s
}

func (s *UpdateAPIDefinitionRequestParamLimitMethodListFormParamList) SetType(v string) *UpdateAPIDefinitionRequestParamLimitMethodListFormParamList {
  s.Type = &v
  return s
}

func (s *UpdateAPIDefinitionRequestParamLimitMethodListFormParamList) SetMinVal(v string) *UpdateAPIDefinitionRequestParamLimitMethodListFormParamList {
  s.MinVal = &v
  return s
}

func (s *UpdateAPIDefinitionRequestParamLimitMethodListFormParamList) SetMaxVal(v string) *UpdateAPIDefinitionRequestParamLimitMethodListFormParamList {
  s.MaxVal = &v
  return s
}

func (s *UpdateAPIDefinitionRequestParamLimitMethodListFormParamList) SetRequired(v string) *UpdateAPIDefinitionRequestParamLimitMethodListFormParamList {
  s.Required = &v
  return s
}

func (s *UpdateAPIDefinitionRequestParamLimitMethodListFormParamList) SetContent(v string) *UpdateAPIDefinitionRequestParamLimitMethodListFormParamList {
  s.Content = &v
  return s
}

func (s *UpdateAPIDefinitionRequestParamLimitMethodListFormParamList) SetRemark(v string) *UpdateAPIDefinitionRequestParamLimitMethodListFormParamList {
  s.Remark = &v
  return s
}

type UpdateAPIDefinitionRequestParamLimitMethodListJsonParam struct {
  // {"en":"Parameter Name.","zh_CN":"参数名称。"}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // {"en":"Parameter Type.\nArray:Array\nJson:JSON Object","zh_CN":"参数类型。\nArray：数组\nJson：JSON对象","exampleValue":"Array,Json"}
  Type *string `json:"type,omitempty" xml:"type,omitempty"`
  // {"en":"type = Integer/Number:minimum value, type = String:minimum length, type = Boolean/Enumeration:leave blank.","zh_CN":"type = Integer/Number：最小值，type = String：最小长度，type = Boolean/Enumeration：置空。"}
  MinVal *string `json:"minVal,omitempty" xml:"minVal,omitempty"`
  // {"en":"type = Integer/Number:maximum value, type = String:maximum length, type = Boolean/Enumeration:leave blank.","zh_CN":"type = Integer/Number：最大值，type = String：最大长度，type = Boolean/Enumeration：置空。"}
  MaxVal *string `json:"maxVal,omitempty" xml:"maxVal,omitempty"`
  // {"en":"Required.\nTRUE:Yes\nFALSE:No","zh_CN":"必带参数。\nTRUE：是\nFALSE：否","exampleValue":"TRUE,FALSE"}
  Required *string `json:"required,omitempty" xml:"required,omitempty"`
  // {"en":"Content (maximum 2000 characters), required when type = Enumeration, multiple separated by ; sign.","zh_CN":"内容（最多2000个字符），type = Enumeration 时必填，多个以 ; 号分隔。"}
  Content *string `json:"content,omitempty" xml:"content,omitempty"`
  // {"en":"Array of child nodes, required when type= Array.","zh_CN":"子节点数组，type = Array 时必填。"}
  Children []*UpdateAPIDefinitionRequestParamLimitMethodListJsonParamChildren `json:"children,omitempty" xml:"children,omitempty" type:"Repeated"`
}

func (s UpdateAPIDefinitionRequestParamLimitMethodListJsonParam) String() string {
  return tea.Prettify(s)
}

func (s UpdateAPIDefinitionRequestParamLimitMethodListJsonParam) GoString() string {
  return s.String()
}

func (s *UpdateAPIDefinitionRequestParamLimitMethodListJsonParam) SetName(v string) *UpdateAPIDefinitionRequestParamLimitMethodListJsonParam {
  s.Name = &v
  return s
}

func (s *UpdateAPIDefinitionRequestParamLimitMethodListJsonParam) SetType(v string) *UpdateAPIDefinitionRequestParamLimitMethodListJsonParam {
  s.Type = &v
  return s
}

func (s *UpdateAPIDefinitionRequestParamLimitMethodListJsonParam) SetMinVal(v string) *UpdateAPIDefinitionRequestParamLimitMethodListJsonParam {
  s.MinVal = &v
  return s
}

func (s *UpdateAPIDefinitionRequestParamLimitMethodListJsonParam) SetMaxVal(v string) *UpdateAPIDefinitionRequestParamLimitMethodListJsonParam {
  s.MaxVal = &v
  return s
}

func (s *UpdateAPIDefinitionRequestParamLimitMethodListJsonParam) SetRequired(v string) *UpdateAPIDefinitionRequestParamLimitMethodListJsonParam {
  s.Required = &v
  return s
}

func (s *UpdateAPIDefinitionRequestParamLimitMethodListJsonParam) SetContent(v string) *UpdateAPIDefinitionRequestParamLimitMethodListJsonParam {
  s.Content = &v
  return s
}

func (s *UpdateAPIDefinitionRequestParamLimitMethodListJsonParam) SetChildren(v []*UpdateAPIDefinitionRequestParamLimitMethodListJsonParamChildren) *UpdateAPIDefinitionRequestParamLimitMethodListJsonParam {
  s.Children = v
  return s
}

type UpdateAPIDefinitionRequestParamLimitMethodListJsonParamChildren struct     {
  // {"en":"Parameter Name.","zh_CN":"参数名称。"}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // {"en":"Parameter Type.\nInteger:Integer\nNumber:Number\nString:String\nBoolean:Boolean\nEnumeration:Enumeration\nArray:Array\nJson:JSON Object","zh_CN":"参数类型。\nInteger：整数\nNumber：数字\nString：字符串\nBoolean：布尔\nEnumeration：枚举\nArray：数组\nJson：JSON对象","exampleValue":"Integer,Number,String,Boolean,Enumeration,Array,Json"}
  Type *string `json:"type,omitempty" xml:"type,omitempty"`
  // {"en":"type = Integer/Number:minimum value, type = String:minimum length, type = Boolean/Enumeration:leave blank.","zh_CN":"type = Integer/Number：最小值，type = String：最小长度，type = Boolean/Enumeration：置空。"}
  MinVal *string `json:"minVal,omitempty" xml:"minVal,omitempty"`
  // {"en":"type = Integer/Number:maximum value, type = String:maximum length, type = Boolean/Enumeration:leave blank.","zh_CN":"type = Integer/Number：最大值，type = String：最大长度，type = Boolean/Enumeration：置空。"}
  MaxVal *string `json:"maxVal,omitempty" xml:"maxVal,omitempty"`
  // {"en":"Required.\nTRUE:Yes\nFALSE:No","zh_CN":"必带参数。\nTRUE：是\nFALSE：否","exampleValue":"TRUE,FALSE"}
  Required *string `json:"required,omitempty" xml:"required,omitempty"`
  // {"en":"Content (maximum 2000 characters), required when type = Enumeration, multiple separated by; sign.","zh_CN":"内容（最多2000个字符），type = Enumeration 时必填，多个以 ; 号分隔。"}
  Content *string `json:"content,omitempty" xml:"content,omitempty"`
  // {"en":"Child node array (JSON string), the structure is consistent with the parent node, required when type= Array.","zh_CN":"子节点数组（JSON字符串），结构与父结点一致，type = Array 时必填。"}
  Children []*string `json:"children,omitempty" xml:"children,omitempty" type:"Repeated"`
}

func (s UpdateAPIDefinitionRequestParamLimitMethodListJsonParamChildren) String() string {
  return tea.Prettify(s)
}

func (s UpdateAPIDefinitionRequestParamLimitMethodListJsonParamChildren) GoString() string {
  return s.String()
}

func (s *UpdateAPIDefinitionRequestParamLimitMethodListJsonParamChildren) SetName(v string) *UpdateAPIDefinitionRequestParamLimitMethodListJsonParamChildren {
  s.Name = &v
  return s
}

func (s *UpdateAPIDefinitionRequestParamLimitMethodListJsonParamChildren) SetType(v string) *UpdateAPIDefinitionRequestParamLimitMethodListJsonParamChildren {
  s.Type = &v
  return s
}

func (s *UpdateAPIDefinitionRequestParamLimitMethodListJsonParamChildren) SetMinVal(v string) *UpdateAPIDefinitionRequestParamLimitMethodListJsonParamChildren {
  s.MinVal = &v
  return s
}

func (s *UpdateAPIDefinitionRequestParamLimitMethodListJsonParamChildren) SetMaxVal(v string) *UpdateAPIDefinitionRequestParamLimitMethodListJsonParamChildren {
  s.MaxVal = &v
  return s
}

func (s *UpdateAPIDefinitionRequestParamLimitMethodListJsonParamChildren) SetRequired(v string) *UpdateAPIDefinitionRequestParamLimitMethodListJsonParamChildren {
  s.Required = &v
  return s
}

func (s *UpdateAPIDefinitionRequestParamLimitMethodListJsonParamChildren) SetContent(v string) *UpdateAPIDefinitionRequestParamLimitMethodListJsonParamChildren {
  s.Content = &v
  return s
}

func (s *UpdateAPIDefinitionRequestParamLimitMethodListJsonParamChildren) SetChildren(v []*string) *UpdateAPIDefinitionRequestParamLimitMethodListJsonParamChildren {
  s.Children = v
  return s
}

type UpdateAPIDefinitionRequestHeader struct {
  // {"en":"Security service type. Please enter a specific service type, if you purchase multiple security services.","zh_CN":"安全服务类型。有使用多个不同的安全服务时，需要填写具体的服务类型。"}
  ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty"`
}

func (s UpdateAPIDefinitionRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s UpdateAPIDefinitionRequestHeader) GoString() string {
  return s.String()
}

func (s *UpdateAPIDefinitionRequestHeader) SetServiceType(v string) *UpdateAPIDefinitionRequestHeader {
  s.ServiceType = &v
  return s
}

type UpdateAPIDefinitionPaths struct {
}

func (s UpdateAPIDefinitionPaths) String() string {
  return tea.Prettify(s)
}

func (s UpdateAPIDefinitionPaths) GoString() string {
  return s.String()
}

type UpdateAPIDefinitionParameters struct {
}

func (s UpdateAPIDefinitionParameters) String() string {
  return tea.Prettify(s)
}

func (s UpdateAPIDefinitionParameters) GoString() string {
  return s.String()
}

type UpdateAPIDefinitionResponse struct {
  // {"dictionary":"belong=WAAP-MS-Ext|dict=waap_retCodeEnum","en":"Please refer to the error code for exceptions.","zh_CN":"请参照错误码。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Description.","zh_CN":"描述信息。"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
}

func (s UpdateAPIDefinitionResponse) String() string {
  return tea.Prettify(s)
}

func (s UpdateAPIDefinitionResponse) GoString() string {
  return s.String()
}

func (s *UpdateAPIDefinitionResponse) SetCode(v string) *UpdateAPIDefinitionResponse {
  s.Code = &v
  return s
}

func (s *UpdateAPIDefinitionResponse) SetMsg(v string) *UpdateAPIDefinitionResponse {
  s.Msg = &v
  return s
}

type UpdateAPIDefinitionResponseHeader struct {
}

func (s UpdateAPIDefinitionResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s UpdateAPIDefinitionResponseHeader) GoString() string {
  return s.String()
}




type ListAPIAssetDiscoveryRequest struct {
  // {"en":"Domain list.","zh_CN":"域名列表。"}
  DomainList []*string `json:"domainList,omitempty" xml:"domainList,omitempty" require:"true" type:"Repeated"`
  // {"en":"API base path.","zh_CN":"端点路径。"}
  PathList []*string `json:"pathList,omitempty" xml:"pathList,omitempty" type:"Repeated"`
  // {"en":"Definition state.\nDEFINED: Defined\nUNDEFINED:Undefined","zh_CN":"定义状态。\nDEFINED: 已定义\nUNDEFINED:未定义","exampleValue":"DEFINED,UNDEFINED"}
  DefineStatus *string `json:"defineStatus,omitempty" xml:"defineStatus,omitempty"`
  // {"en":"Sort, format: field1,sort1;field2,sort2.\nOptional field:\nlastDiscoveryTime: Last discovery time\nfirstDiscoveryTime: First discovery time\nreqCount24h: 24h Requests\nOptional sort:\nascending:Ascend\ndescending:Descend.","zh_CN":"排序，格式：字段1,排序1;字段2,排序2。\n可选字段：\nlastDiscoveryTime：最新发现时间\nfirstDiscoveryTime：首次发现时间\nreqCount24h：24h调用量\n可选排序：\nascending：升序\ndescending：降序"}
  OrderBy *string `json:"orderBy,omitempty" xml:"orderBy,omitempty"`
}

func (s ListAPIAssetDiscoveryRequest) String() string {
  return tea.Prettify(s)
}

func (s ListAPIAssetDiscoveryRequest) GoString() string {
  return s.String()
}

func (s *ListAPIAssetDiscoveryRequest) SetDomainList(v []*string) *ListAPIAssetDiscoveryRequest {
  s.DomainList = v
  return s
}

func (s *ListAPIAssetDiscoveryRequest) SetPathList(v []*string) *ListAPIAssetDiscoveryRequest {
  s.PathList = v
  return s
}

func (s *ListAPIAssetDiscoveryRequest) SetDefineStatus(v string) *ListAPIAssetDiscoveryRequest {
  s.DefineStatus = &v
  return s
}

func (s *ListAPIAssetDiscoveryRequest) SetOrderBy(v string) *ListAPIAssetDiscoveryRequest {
  s.OrderBy = &v
  return s
}

type ListAPIAssetDiscoveryRequestHeader struct {
  // {"en":"Security service type. Please enter a specific service type, if you purchase multiple security services.","zh_CN":"安全服务类型。有使用多个不同的安全服务时，需要填写具体的服务类型。"}
  ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty"`
}

func (s ListAPIAssetDiscoveryRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ListAPIAssetDiscoveryRequestHeader) GoString() string {
  return s.String()
}

func (s *ListAPIAssetDiscoveryRequestHeader) SetServiceType(v string) *ListAPIAssetDiscoveryRequestHeader {
  s.ServiceType = &v
  return s
}

type ListAPIAssetDiscoveryPaths struct {
}

func (s ListAPIAssetDiscoveryPaths) String() string {
  return tea.Prettify(s)
}

func (s ListAPIAssetDiscoveryPaths) GoString() string {
  return s.String()
}

type ListAPIAssetDiscoveryParameters struct {
}

func (s ListAPIAssetDiscoveryParameters) String() string {
  return tea.Prettify(s)
}

func (s ListAPIAssetDiscoveryParameters) GoString() string {
  return s.String()
}

type ListAPIAssetDiscoveryResponse struct {
  // {"dictionary":"belong=WAAP-MS-Ext|dict=waap_retCodeEnum","en":"Please refer to the error code for exceptions.","zh_CN":"请参照错误码。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Description.","zh_CN":"描述信息。"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"Data.","zh_CN":"出参数据。"}
  Data []*ListAPIAssetDiscoveryResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s ListAPIAssetDiscoveryResponse) String() string {
  return tea.Prettify(s)
}

func (s ListAPIAssetDiscoveryResponse) GoString() string {
  return s.String()
}

func (s *ListAPIAssetDiscoveryResponse) SetCode(v string) *ListAPIAssetDiscoveryResponse {
  s.Code = &v
  return s
}

func (s *ListAPIAssetDiscoveryResponse) SetMsg(v string) *ListAPIAssetDiscoveryResponse {
  s.Msg = &v
  return s
}

func (s *ListAPIAssetDiscoveryResponse) SetData(v []*ListAPIAssetDiscoveryResponseData) *ListAPIAssetDiscoveryResponse {
  s.Data = v
  return s
}

type ListAPIAssetDiscoveryResponseData struct     {
  // {"en":"API discovery ID.","zh_CN":"API发现ID。"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {"en":"Domain.","zh_CN":"域名。"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {"en":"API base path.","zh_CN":"端点路径。"}
  Path *string `json:"path,omitempty" xml:"path,omitempty" require:"true"`
  // {"en":"Last Discovery Time, format: yyyy-MM-dd HH:mm:ss.","zh_CN":"最新发现时间，格式：yyyy-MM-dd HH:mm:ss。"}
  LastDiscoveryTime *string `json:"lastDiscoveryTime,omitempty" xml:"lastDiscoveryTime,omitempty" require:"true"`
  // {"en":"First Discovery Time, format: yyyy-MM-dd HH:mm:ss.","zh_CN":"首次发现时间，格式：yyyy-MM-dd HH:mm:ss。"}
  FirstDiscoveryTime *string `json:"firstDiscoveryTime,omitempty" xml:"firstDiscoveryTime,omitempty" require:"true"`
  // {"en":"24h Requests.","zh_CN":"24h调用量。"}
  ReqCount24h *int64 `json:"reqCount24h,omitempty" xml:"reqCount24h,omitempty" require:"true"`
  // {"en":"Status.\nDEFINED: Defined\nUNDEFINED:Pending","zh_CN":"定义状态。\nDEFINED: 已定义\nUNDEFINED: 待确认","exampleValue":"DEFINED,UNDEFINED"}
  DefineStatus *string `json:"defineStatus,omitempty" xml:"defineStatus,omitempty" require:"true"`
}

func (s ListAPIAssetDiscoveryResponseData) String() string {
  return tea.Prettify(s)
}

func (s ListAPIAssetDiscoveryResponseData) GoString() string {
  return s.String()
}

func (s *ListAPIAssetDiscoveryResponseData) SetId(v string) *ListAPIAssetDiscoveryResponseData {
  s.Id = &v
  return s
}

func (s *ListAPIAssetDiscoveryResponseData) SetDomain(v string) *ListAPIAssetDiscoveryResponseData {
  s.Domain = &v
  return s
}

func (s *ListAPIAssetDiscoveryResponseData) SetPath(v string) *ListAPIAssetDiscoveryResponseData {
  s.Path = &v
  return s
}

func (s *ListAPIAssetDiscoveryResponseData) SetLastDiscoveryTime(v string) *ListAPIAssetDiscoveryResponseData {
  s.LastDiscoveryTime = &v
  return s
}

func (s *ListAPIAssetDiscoveryResponseData) SetFirstDiscoveryTime(v string) *ListAPIAssetDiscoveryResponseData {
  s.FirstDiscoveryTime = &v
  return s
}

func (s *ListAPIAssetDiscoveryResponseData) SetReqCount24h(v int64) *ListAPIAssetDiscoveryResponseData {
  s.ReqCount24h = &v
  return s
}

func (s *ListAPIAssetDiscoveryResponseData) SetDefineStatus(v string) *ListAPIAssetDiscoveryResponseData {
  s.DefineStatus = &v
  return s
}

type ListAPIAssetDiscoveryResponseHeader struct {
}

func (s ListAPIAssetDiscoveryResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ListAPIAssetDiscoveryResponseHeader) GoString() string {
  return s.String()
}




type CreateAPIDefinitionRequest struct {
  // {"en":"Basic information.","zh_CN":"基础信息。"}
  Basic *CreateAPIDefinitionRequestBasic `json:"basic,omitempty" xml:"basic,omitempty" require:"true" type:"Struct"`
  // {"en":"Endpoint information.","zh_CN":"端点信息。"}
  Endpoint *CreateAPIDefinitionRequestEndpoint `json:"endpoint,omitempty" xml:"endpoint,omitempty" require:"true" type:"Struct"`
  // {"en":"Authentication configuration.","zh_CN":"鉴权配置。"}
  Auth *CreateAPIDefinitionRequestAuth `json:"auth,omitempty" xml:"auth,omitempty" require:"true" type:"Struct"`
  // {"en":"Body restrictions.","zh_CN":"body限制。"}
  BodyLimit *CreateAPIDefinitionRequestBodyLimit `json:"bodyLimit,omitempty" xml:"bodyLimit,omitempty" type:"Struct"`
  // {"en":"Parameter limit.","zh_CN":"参数限制。"}
  ParamLimit *CreateAPIDefinitionRequestParamLimit `json:"paramLimit,omitempty" xml:"paramLimit,omitempty" type:"Struct"`
}

func (s CreateAPIDefinitionRequest) String() string {
  return tea.Prettify(s)
}

func (s CreateAPIDefinitionRequest) GoString() string {
  return s.String()
}

func (s *CreateAPIDefinitionRequest) SetBasic(v *CreateAPIDefinitionRequestBasic) *CreateAPIDefinitionRequest {
  s.Basic = v
  return s
}

func (s *CreateAPIDefinitionRequest) SetEndpoint(v *CreateAPIDefinitionRequestEndpoint) *CreateAPIDefinitionRequest {
  s.Endpoint = v
  return s
}

func (s *CreateAPIDefinitionRequest) SetAuth(v *CreateAPIDefinitionRequestAuth) *CreateAPIDefinitionRequest {
  s.Auth = v
  return s
}

func (s *CreateAPIDefinitionRequest) SetBodyLimit(v *CreateAPIDefinitionRequestBodyLimit) *CreateAPIDefinitionRequest {
  s.BodyLimit = v
  return s
}

func (s *CreateAPIDefinitionRequest) SetParamLimit(v *CreateAPIDefinitionRequestParamLimit) *CreateAPIDefinitionRequest {
  s.ParamLimit = v
  return s
}

type CreateAPIDefinitionRequestBasic struct {
  // {"en":"API name, maximum 200 characters.","zh_CN":"API名称，最多200个字符。"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"API groups, maximum 64 characters.","zh_CN":"API分组，最多64个字符。"}
  GroupName *string `json:"groupName,omitempty" xml:"groupName,omitempty" require:"true"`
  // {"en":"Domain.","zh_CN":"所属域名。"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {"en":"Description, maximum 200 characters.","zh_CN":"备注，最多200个字符。"}
  Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
}

func (s CreateAPIDefinitionRequestBasic) String() string {
  return tea.Prettify(s)
}

func (s CreateAPIDefinitionRequestBasic) GoString() string {
  return s.String()
}

func (s *CreateAPIDefinitionRequestBasic) SetName(v string) *CreateAPIDefinitionRequestBasic {
  s.Name = &v
  return s
}

func (s *CreateAPIDefinitionRequestBasic) SetGroupName(v string) *CreateAPIDefinitionRequestBasic {
  s.GroupName = &v
  return s
}

func (s *CreateAPIDefinitionRequestBasic) SetDomain(v string) *CreateAPIDefinitionRequestBasic {
  s.Domain = &v
  return s
}

func (s *CreateAPIDefinitionRequestBasic) SetRemark(v string) *CreateAPIDefinitionRequestBasic {
  s.Remark = &v
  return s
}

type CreateAPIDefinitionRequestEndpoint struct {
  // {"en":"Request methods.Multiple separated by ; sign.\nGET:GET, configurable parameter limits\nPOST:POST, configurable parameter limits\nDELETE:DELETE, configurable parameter limits\nUPDATE:UPDATE\nPUT:PUT, configurable parameter limits\nHEAD:HEAD, configurable parameter limits\nCONNECT:CONNECT\nOPTIONS:OPTIONS, configurable parameter limits\nCOPY:COPY\nLOCK:LOCK\nUNLOCK:UNLOCK\nTRACE:TRACE\nPATCH:PATCH, configurable parameter limits\nPROPFIND:PROPFIND\nMKCOL:MKCOL\nMOVE:MOVE","zh_CN":"请求方法。多个以 ; 号分隔。\nGET：GET，可配置参数限制\nPOST：POST，可配置参数限制\nDELETE：DELETE，可配置参数限制\nUPDATE：UPDATE\nPUT：PUT，可配置参数限制\nHEAD：HEAD，可配置参数限制\nCONNECT：CONNECT\nOPTIONS：OPTIONS，可配置参数限制\nCOPY：COPY\nLOCK：LOCK\nUNLOCK：UNLOCK\nTRACE：TRACE\nPATCH：PATCH，可配置参数限制\nPROPFIND：PROPFIND\nMKCOL：MKCOL\nMOVE：MOVE","exampleValue":"GET,POST,DELETE,UPDATE,PUT,HEAD,CONNECT,OPTIONS,COPY,LOCK,UNLOCK,TRACE,PATCH,PROPFIND,MKCOL,MOVE"}
  RequestMethod *string `json:"requestMethod,omitempty" xml:"requestMethod,omitempty" require:"true"`
  // {"en":"API type.\nNORMAL:Common API, the path does not contain query string parameters, such as http://www.test.com/api.\nWHEN_CASE:Common API, the path contains query string parameters, such as http://www.test.com/api?action=1 and http://www.test.com/api?action=2 are two different APIs.","zh_CN":"API类型。\nNORMAL：普通接口，路径中不包含query string参数的普通接口，如http://www.test.com/api。\nWHEN_CASE：when-case接口，路径中包含query string参数，如http://www.test.com/api?action=1与http://www.test.com/api?action=2 是两个不同的接口。","exampleValue":"NORMAL,WHEN_CASE"}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
  // {"en":"Path matching type, EQUAL is passed when type = WHEM_CASE.\nEQUAL: Complete match\nREGEX: Regular matching","zh_CN":"路径匹配类型，type = WHEN_CASE 时传 EQUAL。\nEQUAL：完整匹配\nREGEX：正则匹配","exampleValue":"EQUAL,REGEX"}
  MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
  // {"en":"API base Path, only English characters, underscores, hyphens, or numbers are supported, maximum 200 characters.","zh_CN":"端点路径，只支持英文字符、下划线、短横线或数字，最多200个字符。"}
  Path *string `json:"path,omitempty" xml:"path,omitempty" require:"true"`
  // {"en":"Case sensitive, whether the endpoint path is case-sensitive. Example for case sensitive: /Order and /order are two different API paths.\nTRUE:Yes\nFALSE:No","zh_CN":"大小写是否敏感，端点路径是否区分大小写。若大小写敏感，示例：/Order 和 /order为两不同的API路径。\nTRUE：是\nFALSE：否","exampleValue":"TRUE,FALSE"}
  CaseSensitive *string `json:"caseSensitive,omitempty" xml:"caseSensitive,omitempty" require:"true"`
  // {"en":"Match CreateAPIDefinitionParameters in the Path(Effective when type = NORMAL && matchType = EQUAL), with matching path parameters turned on, you can use curly braces \"{}\" to define path parameters. Example: /basePath/{pathParam1}/{pathParam2}/{pathParam3}.\nTRUE:Yes\nFALSE:No","zh_CN":"匹配路径参数（type = NORMAL && matchType = EQUAL 时生效），开启匹配路径参数后，您可以使用花括号({})来定义路径参数。例：/basePath/{pathParam1}/{pathParam2}/{pathParam3}。\nTRUE：是\nFALSE：否","exampleValue":"TRUE,FALSE"}
  MatchPathVar *string `json:"matchPathVar,omitempty" xml:"matchPathVar,omitempty"`
}

func (s CreateAPIDefinitionRequestEndpoint) String() string {
  return tea.Prettify(s)
}

func (s CreateAPIDefinitionRequestEndpoint) GoString() string {
  return s.String()
}

func (s *CreateAPIDefinitionRequestEndpoint) SetRequestMethod(v string) *CreateAPIDefinitionRequestEndpoint {
  s.RequestMethod = &v
  return s
}

func (s *CreateAPIDefinitionRequestEndpoint) SetType(v string) *CreateAPIDefinitionRequestEndpoint {
  s.Type = &v
  return s
}

func (s *CreateAPIDefinitionRequestEndpoint) SetMatchType(v string) *CreateAPIDefinitionRequestEndpoint {
  s.MatchType = &v
  return s
}

func (s *CreateAPIDefinitionRequestEndpoint) SetPath(v string) *CreateAPIDefinitionRequestEndpoint {
  s.Path = &v
  return s
}

func (s *CreateAPIDefinitionRequestEndpoint) SetCaseSensitive(v string) *CreateAPIDefinitionRequestEndpoint {
  s.CaseSensitive = &v
  return s
}

func (s *CreateAPIDefinitionRequestEndpoint) SetMatchPathVar(v string) *CreateAPIDefinitionRequestEndpoint {
  s.MatchPathVar = &v
  return s
}

type CreateAPIDefinitionRequestAuth struct {
  // {"en":"Authentication method.\nNO_NEED:No authentication required\nSIGN:Key authentication","zh_CN":"鉴权方法。\nNO_NEED：免鉴权\nSIGN：秘钥对鉴权","exampleValue":"NO_NEED,SIGN"}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
  // {"en":"Authentication key, type = SIGN is required, the format is a 16-digit string containing uppercase and lowercase letters and numbers, example: gjZkg2E1uNkXBDxj.","zh_CN":"鉴权秘钥，type = SIGN 时必填，格式为16位含大小写字母与数字字符串，示例：gjZkg2E1uNkXBDxj。"}
  AuthKey *string `json:"authKey,omitempty" xml:"authKey,omitempty"`
  // {"en":"Authentication parameter location, type = SIGN is required.\nHTTP_HEADER:Http Header\nQUERY_STRING:Query String\nCOOKIE:Cookie","zh_CN":"鉴权参数位置，type = SIGN 时必填。\nHTTP_HEADER：Http Header\nQUERY_STRING：Query String\nCOOKIE：Cookie","exampleValue":"HTTP_HEADER,QUERY_STRING,COOKIE"}
  ParamPosition *string `json:"paramPosition,omitempty" xml:"paramPosition,omitempty"`
  // {"en":"Authentication parameter name, type = SIGN is required.","zh_CN":"鉴权参数名称，type = SIGN 时必填。"}
  ParamName *string `json:"paramName,omitempty" xml:"paramName,omitempty"`
  // {"en":"Authentication token validity period, in seconds, type = SIGN is required.","zh_CN":"鉴权有效期，单位秒，type = SIGN 时必填。"}
  ValidityTime *int64 `json:"validityTime,omitempty" xml:"validityTime,omitempty"`
}

func (s CreateAPIDefinitionRequestAuth) String() string {
  return tea.Prettify(s)
}

func (s CreateAPIDefinitionRequestAuth) GoString() string {
  return s.String()
}

func (s *CreateAPIDefinitionRequestAuth) SetType(v string) *CreateAPIDefinitionRequestAuth {
  s.Type = &v
  return s
}

func (s *CreateAPIDefinitionRequestAuth) SetAuthKey(v string) *CreateAPIDefinitionRequestAuth {
  s.AuthKey = &v
  return s
}

func (s *CreateAPIDefinitionRequestAuth) SetParamPosition(v string) *CreateAPIDefinitionRequestAuth {
  s.ParamPosition = &v
  return s
}

func (s *CreateAPIDefinitionRequestAuth) SetParamName(v string) *CreateAPIDefinitionRequestAuth {
  s.ParamName = &v
  return s
}

func (s *CreateAPIDefinitionRequestAuth) SetValidityTime(v int64) *CreateAPIDefinitionRequestAuth {
  s.ValidityTime = &v
  return s
}

type CreateAPIDefinitionRequestBodyLimit struct {
  // {"defaultValue":"OFF","en":"Request body restriction switch.default value:OFF.\nON:On\nOFF:Off","zh_CN":"请求body限制开关。默认值：关。\nON：开启\nOFF：关闭","exampleValue":"ON,OFF"}
  DefendSwitch *string `json:"defendSwitch,omitempty" xml:"defendSwitch,omitempty"`
  // {"en":"Content-Type, required when defendSwitch = ON.\nFORM:FORM\nJSON:JSON\nANY:ANY\nEMPTY:EMPTY or NON-EXISTENT","zh_CN":"Content-Type，defendSwitch = ON 时必填。\nFORM：FORM表单\nJSON：JSON\nANY：任意\nEMPTY：为空或不存在","exampleValue":"FORM,JSON,ANY,EMPTY"}
  ContentType *string `json:"contentType,omitempty" xml:"contentType,omitempty"`
  // {"en":"Maximum body limit(bytes), does not exceed 16,384.","zh_CN":"Body最大限制(bytes)，最大不超过16384。"}
  BodyLimitMax *int64 `json:"bodyLimitMax,omitempty" xml:"bodyLimitMax,omitempty"`
  // {"en":"Maximum nesting depth, enter the maximum allowable JSON nesting depth in the request body.","zh_CN":"最大嵌套层数，输入允许的请求body中JSON嵌套层数最大值。"}
  NestMax *int64 `json:"nestMax,omitempty" xml:"nestMax,omitempty"`
  // {"en":"Maximum number of parameters for JSON, enter the maximum number of JSON parameters allowed in the request body.","zh_CN":"JSON最大参数个数，输入允许的请求body中JSON参数个数的最大值。"}
  ParamCountMax *int64 `json:"paramCountMax,omitempty" xml:"paramCountMax,omitempty"`
}

func (s CreateAPIDefinitionRequestBodyLimit) String() string {
  return tea.Prettify(s)
}

func (s CreateAPIDefinitionRequestBodyLimit) GoString() string {
  return s.String()
}

func (s *CreateAPIDefinitionRequestBodyLimit) SetDefendSwitch(v string) *CreateAPIDefinitionRequestBodyLimit {
  s.DefendSwitch = &v
  return s
}

func (s *CreateAPIDefinitionRequestBodyLimit) SetContentType(v string) *CreateAPIDefinitionRequestBodyLimit {
  s.ContentType = &v
  return s
}

func (s *CreateAPIDefinitionRequestBodyLimit) SetBodyLimitMax(v int64) *CreateAPIDefinitionRequestBodyLimit {
  s.BodyLimitMax = &v
  return s
}

func (s *CreateAPIDefinitionRequestBodyLimit) SetNestMax(v int64) *CreateAPIDefinitionRequestBodyLimit {
  s.NestMax = &v
  return s
}

func (s *CreateAPIDefinitionRequestBodyLimit) SetParamCountMax(v int64) *CreateAPIDefinitionRequestBodyLimit {
  s.ParamCountMax = &v
  return s
}

type CreateAPIDefinitionRequestParamLimit struct {
  // {"en":"Basic information.","zh_CN":"基础信息。"}
  Basic *CreateAPIDefinitionRequestParamLimitBasic `json:"basic,omitempty" xml:"basic,omitempty" require:"true" type:"Struct"`
  // {"en":"Method attributes.","zh_CN":"方法属性。"}
  MethodList []*CreateAPIDefinitionRequestParamLimitMethodList `json:"methodList,omitempty" xml:"methodList,omitempty" type:"Repeated"`
}

func (s CreateAPIDefinitionRequestParamLimit) String() string {
  return tea.Prettify(s)
}

func (s CreateAPIDefinitionRequestParamLimit) GoString() string {
  return s.String()
}

func (s *CreateAPIDefinitionRequestParamLimit) SetBasic(v *CreateAPIDefinitionRequestParamLimitBasic) *CreateAPIDefinitionRequestParamLimit {
  s.Basic = v
  return s
}

func (s *CreateAPIDefinitionRequestParamLimit) SetMethodList(v []*CreateAPIDefinitionRequestParamLimitMethodList) *CreateAPIDefinitionRequestParamLimit {
  s.MethodList = v
  return s
}

type CreateAPIDefinitionRequestParamLimitBasic struct {
  // {"en":"Parameter limit.\nON:On\nOFF:Off","zh_CN":"参数限制。\nON：开启\nOFF：关闭","exampleValue":"ON,OFF"}
  DefendSwitch *string `json:"defendSwitch,omitempty" xml:"defendSwitch,omitempty"`
  // {"en":"Query String Parameter Detection Mode, required when defendSwitch = ON.\nLOOSE:Loose mode - detect some parameters\nSTRICT:Strict mode - checks all parameters","zh_CN":"Query String参数检测模式，defendSwitch = ON 时必填。\nLOOSE：宽松模式-检测部分参数\nSTRICT：严格模式-检测所有参数","exampleValue":"LOOSE,STRICT"}
  ParamCheckMode *string `json:"paramCheckMode,omitempty" xml:"paramCheckMode,omitempty"`
}

func (s CreateAPIDefinitionRequestParamLimitBasic) String() string {
  return tea.Prettify(s)
}

func (s CreateAPIDefinitionRequestParamLimitBasic) GoString() string {
  return s.String()
}

func (s *CreateAPIDefinitionRequestParamLimitBasic) SetDefendSwitch(v string) *CreateAPIDefinitionRequestParamLimitBasic {
  s.DefendSwitch = &v
  return s
}

func (s *CreateAPIDefinitionRequestParamLimitBasic) SetParamCheckMode(v string) *CreateAPIDefinitionRequestParamLimitBasic {
  s.ParamCheckMode = &v
  return s
}

type CreateAPIDefinitionRequestParamLimitMethodList struct     {
  // {"en":"Request methods.\nGET:GET\nPOST:POST\nDELETE:DELETE\nPUT:PUT\nHEAD:HEAD\nOPTIONS:OPTIONS\nPATCH:PATCH","zh_CN":"请求方法。\nGET：GET\nPOST：POST\nDELETE：DELETE\nPUT：PUT\nHEAD：HEAD\nOPTIONS：OPTIONS\nPATCH：PATCH","exampleValue":"GET,POST,DELETE,PUT,HEAD,OPTIONS,PATCH"}
  Method *string `json:"method,omitempty" xml:"method,omitempty"`
  // {"en":"Whether to define body parameters, required when method = POST/PUT/PATCH.\nTRUE:Yes\nFALSE:No","zh_CN":"是否定义body参数，method = POST/PUT/PATCH时必填。\nTRUE：是\nFALSE：否","exampleValue":"TRUE,FALSE"}
  BodyFlag *string `json:"bodyFlag,omitempty" xml:"bodyFlag,omitempty"`
  // {"en":"Body parameter type, required when bodyFlag = TRUE.\nJSON:JSON\nFROM:FROM","zh_CN":"body参数类型，bodyFlag = TRUE 时必填。\nJSON：JSON\nFROM：FROM表单","exampleValue":"FORM,JSON"}
  BodyType *string `json:"bodyType,omitempty" xml:"bodyType,omitempty"`
  // {"en":"Normal parameter list.","zh_CN":"普通参数数组。"}
  NormalParamList []*CreateAPIDefinitionRequestParamLimitMethodListNormalParamList `json:"normalParamList,omitempty" xml:"normalParamList,omitempty" type:"Repeated"`
  // {"en":"FROM form parameter array, Optional when bodyType = FROM.","zh_CN":"FROM 表单参数数组，bodyType = FROM 时选填。"}
  FormParamList []*CreateAPIDefinitionRequestParamLimitMethodListFormParamList `json:"formParamList,omitempty" xml:"formParamList,omitempty" type:"Repeated"`
  // {"en":"JSON parameter array, Optional when bodyType = JSON.","zh_CN":"JSON参数数组，bodyType = JSON 时选填。"}
  JsonParam *CreateAPIDefinitionRequestParamLimitMethodListJsonParam `json:"jsonParam,omitempty" xml:"jsonParam,omitempty" type:"Struct"`
}

func (s CreateAPIDefinitionRequestParamLimitMethodList) String() string {
  return tea.Prettify(s)
}

func (s CreateAPIDefinitionRequestParamLimitMethodList) GoString() string {
  return s.String()
}

func (s *CreateAPIDefinitionRequestParamLimitMethodList) SetMethod(v string) *CreateAPIDefinitionRequestParamLimitMethodList {
  s.Method = &v
  return s
}

func (s *CreateAPIDefinitionRequestParamLimitMethodList) SetBodyFlag(v string) *CreateAPIDefinitionRequestParamLimitMethodList {
  s.BodyFlag = &v
  return s
}

func (s *CreateAPIDefinitionRequestParamLimitMethodList) SetBodyType(v string) *CreateAPIDefinitionRequestParamLimitMethodList {
  s.BodyType = &v
  return s
}

func (s *CreateAPIDefinitionRequestParamLimitMethodList) SetNormalParamList(v []*CreateAPIDefinitionRequestParamLimitMethodListNormalParamList) *CreateAPIDefinitionRequestParamLimitMethodList {
  s.NormalParamList = v
  return s
}

func (s *CreateAPIDefinitionRequestParamLimitMethodList) SetFormParamList(v []*CreateAPIDefinitionRequestParamLimitMethodListFormParamList) *CreateAPIDefinitionRequestParamLimitMethodList {
  s.FormParamList = v
  return s
}

func (s *CreateAPIDefinitionRequestParamLimitMethodList) SetJsonParam(v *CreateAPIDefinitionRequestParamLimitMethodListJsonParam) *CreateAPIDefinitionRequestParamLimitMethodList {
  s.JsonParam = v
  return s
}

type CreateAPIDefinitionRequestParamLimitMethodListNormalParamList struct     {
  // {"en":"Parameter name, when paramPosition = PATH_PARAMS, the path need to be matched, for example: /basePath/{name}/.","zh_CN":"参数名称，paramPosition = PATH_PARAMS 时需匹配路径变量，例如：/basePath/{name}/。"}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // {"en":"Parameter position.\nHTTP_HEADER:Http Header\nQUERY_STRING:Query String\nCOOKIE:Cookie\nPATH_PARAMS:Path, endpoint.matchPathVar needs to equal TRUE","zh_CN":"参数位置。\nHTTP_HEADER：Http Header\nQUERY_STRING：Query String\nCOOKIE：Cookie\nPATH_PARAMS：路径变量，endpoint.matchPathVar需等于TRUE","exampleValue":"HTTP_HEADER,QUERY_STRING,COOKIE,PATH_PARAMS"}
  ParamPosition *string `json:"paramPosition,omitempty" xml:"paramPosition,omitempty"`
  // {"en":"Parameter type.\nInteger:Integer\nNumber:Number\nString:String\nBoolean:Boolean\nEnumeration:Enumeration","zh_CN":"参数类型。\nInteger：整数\nNumber：数字\nString：字符串\nBoolean：布尔\nEnumeration：枚举","exampleValue":"Integer,Number,String,Boolean,Enumeration"}
  Type *string `json:"type,omitempty" xml:"type,omitempty"`
  // {"en":"type = Integer/Number:minimum value, type = String:minimum length, type = Boolean/Enumeration:leave blank.","zh_CN":"type = Integer/Number：最小值，type = String：最小长度，type = Boolean/Enumeration：置空。"}
  MinVal *string `json:"minVal,omitempty" xml:"minVal,omitempty"`
  // {"en":"type = Integer/Number:maximum value, type = String:maximum length, type = Boolean/Enumeration:leave blank.","zh_CN":"type = Integer/Number：最大值，type = String：最大长度，type = Boolean/Enumeration：置空。"}
  MaxVal *string `json:"maxVal,omitempty" xml:"maxVal,omitempty"`
  // {"en":"Required.\nTRUE:Yes\nFALSE:No","zh_CN":"必带参数。\nTRUE：是\nFALSE：否","exampleValue":"TRUE,FALSE"}
  Required *string `json:"required,omitempty" xml:"required,omitempty"`
  // {"en":"Content (maximum 2000 characters), required when type = Enumeration,multiple separated by  ; sign.","zh_CN":"内容（最多2000个字符），type = Enumeration 时必填，多个以 ; 号分隔。"}
  Content *string `json:"content,omitempty" xml:"content,omitempty"`
  // {"en":"Description.","zh_CN":"备注。"}
  Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
}

func (s CreateAPIDefinitionRequestParamLimitMethodListNormalParamList) String() string {
  return tea.Prettify(s)
}

func (s CreateAPIDefinitionRequestParamLimitMethodListNormalParamList) GoString() string {
  return s.String()
}

func (s *CreateAPIDefinitionRequestParamLimitMethodListNormalParamList) SetName(v string) *CreateAPIDefinitionRequestParamLimitMethodListNormalParamList {
  s.Name = &v
  return s
}

func (s *CreateAPIDefinitionRequestParamLimitMethodListNormalParamList) SetParamPosition(v string) *CreateAPIDefinitionRequestParamLimitMethodListNormalParamList {
  s.ParamPosition = &v
  return s
}

func (s *CreateAPIDefinitionRequestParamLimitMethodListNormalParamList) SetType(v string) *CreateAPIDefinitionRequestParamLimitMethodListNormalParamList {
  s.Type = &v
  return s
}

func (s *CreateAPIDefinitionRequestParamLimitMethodListNormalParamList) SetMinVal(v string) *CreateAPIDefinitionRequestParamLimitMethodListNormalParamList {
  s.MinVal = &v
  return s
}

func (s *CreateAPIDefinitionRequestParamLimitMethodListNormalParamList) SetMaxVal(v string) *CreateAPIDefinitionRequestParamLimitMethodListNormalParamList {
  s.MaxVal = &v
  return s
}

func (s *CreateAPIDefinitionRequestParamLimitMethodListNormalParamList) SetRequired(v string) *CreateAPIDefinitionRequestParamLimitMethodListNormalParamList {
  s.Required = &v
  return s
}

func (s *CreateAPIDefinitionRequestParamLimitMethodListNormalParamList) SetContent(v string) *CreateAPIDefinitionRequestParamLimitMethodListNormalParamList {
  s.Content = &v
  return s
}

func (s *CreateAPIDefinitionRequestParamLimitMethodListNormalParamList) SetRemark(v string) *CreateAPIDefinitionRequestParamLimitMethodListNormalParamList {
  s.Remark = &v
  return s
}

type CreateAPIDefinitionRequestParamLimitMethodListFormParamList struct     {
  // {"en":"Parameter name.","zh_CN":"参数名称。"}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // {"en":"Parameter type.\nInteger:Integer\nNumber:Number\nString:String\nBoolean:Boolean\nEnumeration:Enumeration","zh_CN":"参数类型。\nInteger：整数\nNumber：数字\nString：字符串\nBoolean：布尔\nEnumeration：枚举","exampleValue":"Integer,Number,String,Boolean,Enumeration"}
  Type *string `json:"type,omitempty" xml:"type,omitempty"`
  // {"en":"type = Integer/Number:minimum value, type = String:minimum length, type = Boolean/Enumeration:leave blank.","zh_CN":"type = Integer/Number：最小值，type = String：最小长度，type = Boolean/Enumeration：置空。"}
  MinVal *string `json:"minVal,omitempty" xml:"minVal,omitempty"`
  // {"en":"type = Integer/Number:maximum value, type = String:maximum length, type = Boolean/Enumeration:leave blank.","zh_CN":"type = Integer/Number：最大值，type = String：最大长度，type = Boolean/Enumeration：置空。"}
  MaxVal *string `json:"maxVal,omitempty" xml:"maxVal,omitempty"`
  // {"en":"Required.\nTRUE:Yes\nFALSE:No","zh_CN":"必带参数。\nTRUE：是\nFALSE：否","exampleValue":"TRUE,FALSE"}
  Required *string `json:"required,omitempty" xml:"required,omitempty"`
  // {"en":"Content (maximum 2000 characters), required when type = Enumeration, multiple separated by ; sign.","zh_CN":"内容（最多2000个字符），type = Enumeration 时必填，多个以 ; 号分隔。"}
  Content *string `json:"content,omitempty" xml:"content,omitempty"`
  // {"en":"Description.","zh_CN":"备注。"}
  Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
}

func (s CreateAPIDefinitionRequestParamLimitMethodListFormParamList) String() string {
  return tea.Prettify(s)
}

func (s CreateAPIDefinitionRequestParamLimitMethodListFormParamList) GoString() string {
  return s.String()
}

func (s *CreateAPIDefinitionRequestParamLimitMethodListFormParamList) SetName(v string) *CreateAPIDefinitionRequestParamLimitMethodListFormParamList {
  s.Name = &v
  return s
}

func (s *CreateAPIDefinitionRequestParamLimitMethodListFormParamList) SetType(v string) *CreateAPIDefinitionRequestParamLimitMethodListFormParamList {
  s.Type = &v
  return s
}

func (s *CreateAPIDefinitionRequestParamLimitMethodListFormParamList) SetMinVal(v string) *CreateAPIDefinitionRequestParamLimitMethodListFormParamList {
  s.MinVal = &v
  return s
}

func (s *CreateAPIDefinitionRequestParamLimitMethodListFormParamList) SetMaxVal(v string) *CreateAPIDefinitionRequestParamLimitMethodListFormParamList {
  s.MaxVal = &v
  return s
}

func (s *CreateAPIDefinitionRequestParamLimitMethodListFormParamList) SetRequired(v string) *CreateAPIDefinitionRequestParamLimitMethodListFormParamList {
  s.Required = &v
  return s
}

func (s *CreateAPIDefinitionRequestParamLimitMethodListFormParamList) SetContent(v string) *CreateAPIDefinitionRequestParamLimitMethodListFormParamList {
  s.Content = &v
  return s
}

func (s *CreateAPIDefinitionRequestParamLimitMethodListFormParamList) SetRemark(v string) *CreateAPIDefinitionRequestParamLimitMethodListFormParamList {
  s.Remark = &v
  return s
}

type CreateAPIDefinitionRequestParamLimitMethodListJsonParam struct {
  // {"en":"Parameter Name.","zh_CN":"参数名称。"}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // {"en":"Parameter Type.\nArray:Array\nJson:JSON Object","zh_CN":"参数类型。\nArray：数组\nJson：JSON对象","exampleValue":"Array,Json"}
  Type *string `json:"type,omitempty" xml:"type,omitempty"`
  // {"en":"type = Integer/Number:minimum value, type = String:minimum length, type = Boolean/Enumeration:leave blank.","zh_CN":"type = Integer/Number：最小值，type = String：最小长度，type = Boolean/Enumeration：置空。"}
  MinVal *string `json:"minVal,omitempty" xml:"minVal,omitempty"`
  // {"en":"type = Integer/Number:maximum value, type = String:maximum length, type = Boolean/Enumeration:leave blank.","zh_CN":"type = Integer/Number：最大值，type = String：最大长度，type = Boolean/Enumeration：置空。"}
  MaxVal *string `json:"maxVal,omitempty" xml:"maxVal,omitempty"`
  // {"en":"Required.\nTRUE:Yes\nFALSE:No","zh_CN":"必带参数。\nTRUE：是\nFALSE：否","exampleValue":"TRUE,FALSE"}
  Required *string `json:"required,omitempty" xml:"required,omitempty"`
  // {"en":"Content (maximum 2000 characters), required when type = Enumeration, multiple separated by ; sign.","zh_CN":"内容（最多2000个字符），type = Enumeration 时必填，多个以 ; 号分隔。"}
  Content *string `json:"content,omitempty" xml:"content,omitempty"`
  // {"en":"Child node array, required when type= Array.","zh_CN":"子节点数组，type = Array 时必填。"}
  Children []*CreateAPIDefinitionRequestParamLimitMethodListJsonParamChildren `json:"children,omitempty" xml:"children,omitempty" type:"Repeated"`
}

func (s CreateAPIDefinitionRequestParamLimitMethodListJsonParam) String() string {
  return tea.Prettify(s)
}

func (s CreateAPIDefinitionRequestParamLimitMethodListJsonParam) GoString() string {
  return s.String()
}

func (s *CreateAPIDefinitionRequestParamLimitMethodListJsonParam) SetName(v string) *CreateAPIDefinitionRequestParamLimitMethodListJsonParam {
  s.Name = &v
  return s
}

func (s *CreateAPIDefinitionRequestParamLimitMethodListJsonParam) SetType(v string) *CreateAPIDefinitionRequestParamLimitMethodListJsonParam {
  s.Type = &v
  return s
}

func (s *CreateAPIDefinitionRequestParamLimitMethodListJsonParam) SetMinVal(v string) *CreateAPIDefinitionRequestParamLimitMethodListJsonParam {
  s.MinVal = &v
  return s
}

func (s *CreateAPIDefinitionRequestParamLimitMethodListJsonParam) SetMaxVal(v string) *CreateAPIDefinitionRequestParamLimitMethodListJsonParam {
  s.MaxVal = &v
  return s
}

func (s *CreateAPIDefinitionRequestParamLimitMethodListJsonParam) SetRequired(v string) *CreateAPIDefinitionRequestParamLimitMethodListJsonParam {
  s.Required = &v
  return s
}

func (s *CreateAPIDefinitionRequestParamLimitMethodListJsonParam) SetContent(v string) *CreateAPIDefinitionRequestParamLimitMethodListJsonParam {
  s.Content = &v
  return s
}

func (s *CreateAPIDefinitionRequestParamLimitMethodListJsonParam) SetChildren(v []*CreateAPIDefinitionRequestParamLimitMethodListJsonParamChildren) *CreateAPIDefinitionRequestParamLimitMethodListJsonParam {
  s.Children = v
  return s
}

type CreateAPIDefinitionRequestParamLimitMethodListJsonParamChildren struct     {
  // {"en":"Parameter Name.","zh_CN":"参数名称。"}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // {"en":"Parameter Type.\nInteger:Integer\nNumber:Number\nString:String\nBoolean:Boolean\nEnumeration:Enumeration\nArray:Array\nJson:JSON Object","zh_CN":"参数类型。\nInteger：整数\nNumber：数字\nString：字符串\nBoolean：布尔\nEnumeration：枚举\nArray：数组\nJson：JSON对象","exampleValue":"Integer,Number,String,Boolean,Enumeration,Array,Json"}
  Type *string `json:"type,omitempty" xml:"type,omitempty"`
  // {"en":"type = Integer/Number:minimum value, type = String:minimum length, type = Boolean/Enumeration:leave blank.","zh_CN":"type = Integer/Number：最小值，type = String：最小长度，type = Boolean/Enumeration：置空。"}
  MinVal *string `json:"minVal,omitempty" xml:"minVal,omitempty"`
  // {"en":"type = Integer/Number:maximum value, type = String:maximum length, type = Boolean/Enumeration:leave blank.","zh_CN":"type = Integer/Number：最大值，type = String：最大长度，type = Boolean/Enumeration：置空。"}
  MaxVal *string `json:"maxVal,omitempty" xml:"maxVal,omitempty"`
  // {"en":"Required.\nTRUE:Yes\nFALSE:No","zh_CN":"必带参数。\nTRUE：是\nFALSE：否","exampleValue":"TRUE,FALSE"}
  Required *string `json:"required,omitempty" xml:"required,omitempty"`
  // {"en":"Content (maximum 2000 characters), required when type = Enumeration, multiple separated by ; sign.","zh_CN":"内容（最多2000个字符），type = Enumeration 时必填，多个以 ; 号分隔。"}
  Content *string `json:"content,omitempty" xml:"content,omitempty"`
  // {"en":"Child node array (JSON string), the structure is consistent with the parent node, required when type= Array.","zh_CN":"子节点数组（JSON字符串），结构与父结点一致，type = Array 时必填。"}
  Children []*string `json:"children,omitempty" xml:"children,omitempty" type:"Repeated"`
}

func (s CreateAPIDefinitionRequestParamLimitMethodListJsonParamChildren) String() string {
  return tea.Prettify(s)
}

func (s CreateAPIDefinitionRequestParamLimitMethodListJsonParamChildren) GoString() string {
  return s.String()
}

func (s *CreateAPIDefinitionRequestParamLimitMethodListJsonParamChildren) SetName(v string) *CreateAPIDefinitionRequestParamLimitMethodListJsonParamChildren {
  s.Name = &v
  return s
}

func (s *CreateAPIDefinitionRequestParamLimitMethodListJsonParamChildren) SetType(v string) *CreateAPIDefinitionRequestParamLimitMethodListJsonParamChildren {
  s.Type = &v
  return s
}

func (s *CreateAPIDefinitionRequestParamLimitMethodListJsonParamChildren) SetMinVal(v string) *CreateAPIDefinitionRequestParamLimitMethodListJsonParamChildren {
  s.MinVal = &v
  return s
}

func (s *CreateAPIDefinitionRequestParamLimitMethodListJsonParamChildren) SetMaxVal(v string) *CreateAPIDefinitionRequestParamLimitMethodListJsonParamChildren {
  s.MaxVal = &v
  return s
}

func (s *CreateAPIDefinitionRequestParamLimitMethodListJsonParamChildren) SetRequired(v string) *CreateAPIDefinitionRequestParamLimitMethodListJsonParamChildren {
  s.Required = &v
  return s
}

func (s *CreateAPIDefinitionRequestParamLimitMethodListJsonParamChildren) SetContent(v string) *CreateAPIDefinitionRequestParamLimitMethodListJsonParamChildren {
  s.Content = &v
  return s
}

func (s *CreateAPIDefinitionRequestParamLimitMethodListJsonParamChildren) SetChildren(v []*string) *CreateAPIDefinitionRequestParamLimitMethodListJsonParamChildren {
  s.Children = v
  return s
}

type CreateAPIDefinitionRequestHeader struct {
  // {"en":"Security service type. Please enter a specific service type, if you purchase multiple security services.","zh_CN":"安全服务类型。有使用多个不同的安全服务时，需要填写具体的服务类型。"}
  ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty"`
}

func (s CreateAPIDefinitionRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s CreateAPIDefinitionRequestHeader) GoString() string {
  return s.String()
}

func (s *CreateAPIDefinitionRequestHeader) SetServiceType(v string) *CreateAPIDefinitionRequestHeader {
  s.ServiceType = &v
  return s
}

type CreateAPIDefinitionPaths struct {
}

func (s CreateAPIDefinitionPaths) String() string {
  return tea.Prettify(s)
}

func (s CreateAPIDefinitionPaths) GoString() string {
  return s.String()
}

type CreateAPIDefinitionParameters struct {
}

func (s CreateAPIDefinitionParameters) String() string {
  return tea.Prettify(s)
}

func (s CreateAPIDefinitionParameters) GoString() string {
  return s.String()
}

type CreateAPIDefinitionResponse struct {
  // {"dictionary":"belong=WAAP-MS-Ext|dict=waap_retCodeEnum","en":"Please refer to the error code for exceptions.","zh_CN":"请参照错误码。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Description.","zh_CN":"描述信息。"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"API definition ID.","zh_CN":"API定义ID。"}
  Data *string `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s CreateAPIDefinitionResponse) String() string {
  return tea.Prettify(s)
}

func (s CreateAPIDefinitionResponse) GoString() string {
  return s.String()
}

func (s *CreateAPIDefinitionResponse) SetCode(v string) *CreateAPIDefinitionResponse {
  s.Code = &v
  return s
}

func (s *CreateAPIDefinitionResponse) SetMsg(v string) *CreateAPIDefinitionResponse {
  s.Msg = &v
  return s
}

func (s *CreateAPIDefinitionResponse) SetData(v string) *CreateAPIDefinitionResponse {
  s.Data = &v
  return s
}

type CreateAPIDefinitionResponseHeader struct {
}

func (s CreateAPIDefinitionResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s CreateAPIDefinitionResponseHeader) GoString() string {
  return s.String()
}




type ListAPIDefinitionsRequest struct {
  // {"en":"Domain array, by default all valid domains.","zh_CN":"域名数组，默认所有生效域名。"}
  DomainList []*string `json:"domainList,omitempty" xml:"domainList,omitempty" type:"Repeated"`
  // {"en":"API name array.","zh_CN":"API名称数组。"}
  NameList []*string `json:"nameList,omitempty" xml:"nameList,omitempty" type:"Repeated"`
  // {"en":"API grouped array.","zh_CN":"API分组数组。"}
  GroupNameList []*string `json:"groupNameList,omitempty" xml:"groupNameList,omitempty" type:"Repeated"`
  // {"en":"Sorting method, reverse order by default.\ntrue: Modify time in reverse order\nfalse: Modify time sequence","zh_CN":"排序方式，默认倒序。\ntrue：修改时间倒序\nfalse：修改时间正序","exampleValue":"true,false"}
  Desc *bool `json:"desc,omitempty" xml:"desc,omitempty"`
}

func (s ListAPIDefinitionsRequest) String() string {
  return tea.Prettify(s)
}

func (s ListAPIDefinitionsRequest) GoString() string {
  return s.String()
}

func (s *ListAPIDefinitionsRequest) SetDomainList(v []*string) *ListAPIDefinitionsRequest {
  s.DomainList = v
  return s
}

func (s *ListAPIDefinitionsRequest) SetNameList(v []*string) *ListAPIDefinitionsRequest {
  s.NameList = v
  return s
}

func (s *ListAPIDefinitionsRequest) SetGroupNameList(v []*string) *ListAPIDefinitionsRequest {
  s.GroupNameList = v
  return s
}

func (s *ListAPIDefinitionsRequest) SetDesc(v bool) *ListAPIDefinitionsRequest {
  s.Desc = &v
  return s
}

type ListAPIDefinitionsRequestHeader struct {
  // {"en":"Security service type. Please enter a specific service type, if you purchase multiple security services.","zh_CN":"安全服务类型。有使用多个不同的安全服务时，需要填写具体的服务类型。"}
  ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty"`
}

func (s ListAPIDefinitionsRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ListAPIDefinitionsRequestHeader) GoString() string {
  return s.String()
}

func (s *ListAPIDefinitionsRequestHeader) SetServiceType(v string) *ListAPIDefinitionsRequestHeader {
  s.ServiceType = &v
  return s
}

type ListAPIDefinitionsPaths struct {
}

func (s ListAPIDefinitionsPaths) String() string {
  return tea.Prettify(s)
}

func (s ListAPIDefinitionsPaths) GoString() string {
  return s.String()
}

type ListAPIDefinitionsParameters struct {
}

func (s ListAPIDefinitionsParameters) String() string {
  return tea.Prettify(s)
}

func (s ListAPIDefinitionsParameters) GoString() string {
  return s.String()
}

type ListAPIDefinitionsResponse struct {
  // {"dictionary":"belong=WAAP-MS-Ext|dict=waap_retCodeEnum","en":"Please refer to the error code for exceptions.","zh_CN":"请参照错误码。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Description.","zh_CN":"描述信息。"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"Data.","zh_CN":"出参数据。"}
  Data []*ListAPIDefinitionsResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s ListAPIDefinitionsResponse) String() string {
  return tea.Prettify(s)
}

func (s ListAPIDefinitionsResponse) GoString() string {
  return s.String()
}

func (s *ListAPIDefinitionsResponse) SetCode(v string) *ListAPIDefinitionsResponse {
  s.Code = &v
  return s
}

func (s *ListAPIDefinitionsResponse) SetMsg(v string) *ListAPIDefinitionsResponse {
  s.Msg = &v
  return s
}

func (s *ListAPIDefinitionsResponse) SetData(v []*ListAPIDefinitionsResponseData) *ListAPIDefinitionsResponse {
  s.Data = v
  return s
}

type ListAPIDefinitionsResponseData struct     {
  // {"en":"API definition ID.","zh_CN":"API定义ID。"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {"en":"API name.","zh_CN":"API名称。"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"API groups.","zh_CN":"API分组。"}
  GroupName *string `json:"groupName,omitempty" xml:"groupName,omitempty" require:"true"`
  // {"en":"Description.","zh_CN":"备注。"}
  Remark *string `json:"remark,omitempty" xml:"remark,omitempty" require:"true"`
  // {"en":"Attributed domain.","zh_CN":"归属域名。"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {"en":"API base path.","zh_CN":"端点路径。"}
  Path *string `json:"path,omitempty" xml:"path,omitempty" require:"true"`
  // {"en":"Creation time, format: yyyy-MM-dd HH:mm:ss.","zh_CN":"创建时间，格式：yyyy-MM-dd HH:mm:ss。"}
  CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty" require:"true"`
  // {"en":"Update time, format: yyyy-MM-dd HH:mm:ss.","zh_CN":"更新时间，格式：yyyy-MM-dd HH:mm:ss。"}
  UpdateTime *string `json:"updateTime,omitempty" xml:"updateTime,omitempty" require:"true"`
}

func (s ListAPIDefinitionsResponseData) String() string {
  return tea.Prettify(s)
}

func (s ListAPIDefinitionsResponseData) GoString() string {
  return s.String()
}

func (s *ListAPIDefinitionsResponseData) SetId(v string) *ListAPIDefinitionsResponseData {
  s.Id = &v
  return s
}

func (s *ListAPIDefinitionsResponseData) SetName(v string) *ListAPIDefinitionsResponseData {
  s.Name = &v
  return s
}

func (s *ListAPIDefinitionsResponseData) SetGroupName(v string) *ListAPIDefinitionsResponseData {
  s.GroupName = &v
  return s
}

func (s *ListAPIDefinitionsResponseData) SetRemark(v string) *ListAPIDefinitionsResponseData {
  s.Remark = &v
  return s
}

func (s *ListAPIDefinitionsResponseData) SetDomain(v string) *ListAPIDefinitionsResponseData {
  s.Domain = &v
  return s
}

func (s *ListAPIDefinitionsResponseData) SetPath(v string) *ListAPIDefinitionsResponseData {
  s.Path = &v
  return s
}

func (s *ListAPIDefinitionsResponseData) SetCreateTime(v string) *ListAPIDefinitionsResponseData {
  s.CreateTime = &v
  return s
}

func (s *ListAPIDefinitionsResponseData) SetUpdateTime(v string) *ListAPIDefinitionsResponseData {
  s.UpdateTime = &v
  return s
}

type ListAPIDefinitionsResponseHeader struct {
}

func (s ListAPIDefinitionsResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ListAPIDefinitionsResponseHeader) GoString() string {
  return s.String()
}




