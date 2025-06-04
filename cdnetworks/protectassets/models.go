package protectassets

// This file is auto-generated, don't edit it. Thanks.
import (
  "github.com/alibabacloud-go/tea/tea"
)

type BatchDeleteAPIDefinitionRequest struct {
  // {'en':'API definition ID array.', 'zh_CN':'API定义ID数组。'}
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

type BatchDeleteAPIDefinitionResponse struct {
  // {'en':'Please refer to the error code for exceptions.', 'zh_CN':'请参照错误码。','dictionary':'belong=WAAP-MS-Ext|dict=waap_retCodeEnum'}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {'en':'Description.', 'zh_CN':'描述信息。'}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {'en':'Data.', 'zh_CN':'出参数据。'}
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

type BatchDeleteAPIDefinitionRequestHeader struct {
  // {"zh_CN":"安全服务类型。有使用多个不同的安全服务时，需要填写具体的服务类型。","en":"Security service type. Please enter a specific service type, if you purchase multiple security services.","dictionary":"belong=WAAP-MS-Ext|dict=waap_serviceType"}
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

type BatchDeleteAPIDefinitionResponseHeader struct {
}

func (s BatchDeleteAPIDefinitionResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s BatchDeleteAPIDefinitionResponseHeader) GoString() string {
  return s.String()
}




type QueryAPIDefinitionDetailRequest struct {
  // {'en':'API definition ID.', 'zh_CN':'API定义ID。'}
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

type QueryAPIDefinitionDetailResponse struct {
  // {'en':'Please refer to the error code for exceptions.', 'zh_CN':'请参照错误码。','dictionary':'belong=WAAP-MS-Ext|dict=waap_retCodeEnum'}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {'en':'Description.', 'zh_CN':'描述信息。'}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {'en':'Data.', 'zh_CN':'出参数据。'}
  Data *QueryAPIDefinitionDetailApiDefineVO `json:"data,omitempty" xml:"data,omitempty" require:"true"`
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

func (s *QueryAPIDefinitionDetailResponse) SetData(v *QueryAPIDefinitionDetailApiDefineVO) *QueryAPIDefinitionDetailResponse {
  s.Data = v
  return s
}

type QueryAPIDefinitionDetailApiDefineVO struct {
  // {'en':'Basic information.', 'zh_CN':'基础信息。'}
  Basic *QueryAPIDefinitionDetailApiDefineBasicVO `json:"basic,omitempty" xml:"basic,omitempty" require:"true"`
  // {'en':'Endpoint information.', 'zh_CN':'端点信息。'}
  Endpoint *QueryAPIDefinitionDetailApiDefineEndpointVO `json:"endpoint,omitempty" xml:"endpoint,omitempty" require:"true"`
  // {'en':'Authentication configuration.', 'zh_CN':'鉴权配置。'}
  Auth *QueryAPIDefinitionDetailApiDefineAuthVO `json:"auth,omitempty" xml:"auth,omitempty" require:"true"`
  // {'en':'Body restrictions.', 'zh_CN':'body限制。'}
  BodyLimit *QueryAPIDefinitionDetailApiDefineBodyLimitVO `json:"bodyLimit,omitempty" xml:"bodyLimit,omitempty" require:"true"`
  // {'en':'Parameter limit.', 'zh_CN':'参数限制。'}
  ParamLimit *QueryAPIDefinitionDetailApiDefineParamLimitVO `json:"paramLimit,omitempty" xml:"paramLimit,omitempty" require:"true"`
}

func (s QueryAPIDefinitionDetailApiDefineVO) String() string {
  return tea.Prettify(s)
}

func (s QueryAPIDefinitionDetailApiDefineVO) GoString() string {
  return s.String()
}

func (s *QueryAPIDefinitionDetailApiDefineVO) SetBasic(v *QueryAPIDefinitionDetailApiDefineBasicVO) *QueryAPIDefinitionDetailApiDefineVO {
  s.Basic = v
  return s
}

func (s *QueryAPIDefinitionDetailApiDefineVO) SetEndpoint(v *QueryAPIDefinitionDetailApiDefineEndpointVO) *QueryAPIDefinitionDetailApiDefineVO {
  s.Endpoint = v
  return s
}

func (s *QueryAPIDefinitionDetailApiDefineVO) SetAuth(v *QueryAPIDefinitionDetailApiDefineAuthVO) *QueryAPIDefinitionDetailApiDefineVO {
  s.Auth = v
  return s
}

func (s *QueryAPIDefinitionDetailApiDefineVO) SetBodyLimit(v *QueryAPIDefinitionDetailApiDefineBodyLimitVO) *QueryAPIDefinitionDetailApiDefineVO {
  s.BodyLimit = v
  return s
}

func (s *QueryAPIDefinitionDetailApiDefineVO) SetParamLimit(v *QueryAPIDefinitionDetailApiDefineParamLimitVO) *QueryAPIDefinitionDetailApiDefineVO {
  s.ParamLimit = v
  return s
}

type QueryAPIDefinitionDetailApiDefineBasicVO struct {
  // {'en':'API definition ID.', 'zh_CN':'API定义ID。'}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {'en':'API name.', 'zh_CN':'API名称。'}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {'en':'API groups.', 'zh_CN':'API分组。'}
  GroupName *string `json:"groupName,omitempty" xml:"groupName,omitempty" require:"true"`
  // {'en':'Description.', 'zh_CN':'备注。'}
  Remark *string `json:"remark,omitempty" xml:"remark,omitempty" require:"true"`
  // {'en':'Attributed hostname.', 'zh_CN':'归属域名。'}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {'en':'API base path.', 'zh_CN':'端点路径。'}
  Path *string `json:"path,omitempty" xml:"path,omitempty" require:"true"`
  // {'en':'Creation time, format: yyyy-MM-dd HH:mm:ss.', 'zh_CN':'创建时间，格式：yyyy-MM-dd HH:mm:ss。'}
  CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty" require:"true"`
  // {'en':'Update time, format: yyyy-MM-dd HH:mm:ss.', 'zh_CN':'更新时间，格式：yyyy-MM-dd HH:mm:ss。'}
  UpdateTime *string `json:"updateTime,omitempty" xml:"updateTime,omitempty" require:"true"`
}

func (s QueryAPIDefinitionDetailApiDefineBasicVO) String() string {
  return tea.Prettify(s)
}

func (s QueryAPIDefinitionDetailApiDefineBasicVO) GoString() string {
  return s.String()
}

func (s *QueryAPIDefinitionDetailApiDefineBasicVO) SetId(v string) *QueryAPIDefinitionDetailApiDefineBasicVO {
  s.Id = &v
  return s
}

func (s *QueryAPIDefinitionDetailApiDefineBasicVO) SetName(v string) *QueryAPIDefinitionDetailApiDefineBasicVO {
  s.Name = &v
  return s
}

func (s *QueryAPIDefinitionDetailApiDefineBasicVO) SetGroupName(v string) *QueryAPIDefinitionDetailApiDefineBasicVO {
  s.GroupName = &v
  return s
}

func (s *QueryAPIDefinitionDetailApiDefineBasicVO) SetRemark(v string) *QueryAPIDefinitionDetailApiDefineBasicVO {
  s.Remark = &v
  return s
}

func (s *QueryAPIDefinitionDetailApiDefineBasicVO) SetDomain(v string) *QueryAPIDefinitionDetailApiDefineBasicVO {
  s.Domain = &v
  return s
}

func (s *QueryAPIDefinitionDetailApiDefineBasicVO) SetPath(v string) *QueryAPIDefinitionDetailApiDefineBasicVO {
  s.Path = &v
  return s
}

func (s *QueryAPIDefinitionDetailApiDefineBasicVO) SetCreateTime(v string) *QueryAPIDefinitionDetailApiDefineBasicVO {
  s.CreateTime = &v
  return s
}

func (s *QueryAPIDefinitionDetailApiDefineBasicVO) SetUpdateTime(v string) *QueryAPIDefinitionDetailApiDefineBasicVO {
  s.UpdateTime = &v
  return s
}

type QueryAPIDefinitionDetailApiDefineEndpointVO struct {
  // {'en':'Request methods.
  // GET:GET, configurable parameter limits
  // POST:POST, configurable parameter limits
  // DELETE:DELETE, configurable parameter limits
  // UPDATE:UPDATE
  // PUT:PUT, configurable parameter limits
  // HEAD:HEAD, configurable parameter limits
  // CONNECT:CONNECT
  // OPTIONS:OPTIONS, configurable parameter limits
  // COPY:COPY
  // LOCK:LOCK
  // UNLOCK:UNLOCK
  // TRACE:TRACE
  // PATCH:PATCH, configurable parameter limits
  // PROPFIND:PROPFIND
  // MKCOL:MKCOL
  // MOVE:MOVE', 'zh_CN':'请求方法。
  // GET：GET，可配置参数限制
  // POST：POST，可配置参数限制
  // DELETE：DELETE，可配置参数限制
  // UPDATE：UPDATE
  // PUT：PUT，可配置参数限制
  // HEAD：HEAD，可配置参数限制
  // CONNECT：CONNECT
  // OPTIONS：OPTIONS，可配置参数限制
  // COPY：COPY
  // LOCK：LOCK
  // UNLOCK：UNLOCK
  // TRACE：TRACE
  // PATCH：PATCH，可配置参数限制
  // PROPFIND：PROPFIND
  // MKCOL：MKCOL
  // MOVE：MOVE'}
  RequestMethod *string `json:"requestMethod,omitempty" xml:"requestMethod,omitempty" require:"true"`
  // {'en':'API type.
  // NORMAL:Common API
  // WHEN_CASE:Common API', 'zh_CN':'API类型。
  // NORMAL：普通接口
  // WHEN_CASE：when-case接口'}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
  // {'en':'Path matching type.
  // EQUAL: Complete match
  // REGEX: Regular matching', 'zh_CN':'路径匹配类型。
  // EQUAL：完整匹配
  // REGEX：正则匹配'}
  MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
  // {'en':'API base path.', 'zh_CN':'端点路径。'}
  Path *string `json:"path,omitempty" xml:"path,omitempty" require:"true"`
  // {'en':'Case sensitive.
  // TRUE:Yes
  // FALSE:No', 'zh_CN':'大小写是否敏感。
  // TRUE：是
  // FALSE：否'}
  CaseSensitive *string `json:"caseSensitive,omitempty" xml:"caseSensitive,omitempty" require:"true"`
  // {'en':'Match QueryAPIDefinitionDetailParameters in the Path.
  // TRUE:Yes
  // FALSE:No', 'zh_CN':'匹配路径参数。
  // TRUE：是
  // FALSE：否'}
  MatchPathVar *string `json:"matchPathVar,omitempty" xml:"matchPathVar,omitempty" require:"true"`
}

func (s QueryAPIDefinitionDetailApiDefineEndpointVO) String() string {
  return tea.Prettify(s)
}

func (s QueryAPIDefinitionDetailApiDefineEndpointVO) GoString() string {
  return s.String()
}

func (s *QueryAPIDefinitionDetailApiDefineEndpointVO) SetRequestMethod(v string) *QueryAPIDefinitionDetailApiDefineEndpointVO {
  s.RequestMethod = &v
  return s
}

func (s *QueryAPIDefinitionDetailApiDefineEndpointVO) SetType(v string) *QueryAPIDefinitionDetailApiDefineEndpointVO {
  s.Type = &v
  return s
}

func (s *QueryAPIDefinitionDetailApiDefineEndpointVO) SetMatchType(v string) *QueryAPIDefinitionDetailApiDefineEndpointVO {
  s.MatchType = &v
  return s
}

func (s *QueryAPIDefinitionDetailApiDefineEndpointVO) SetPath(v string) *QueryAPIDefinitionDetailApiDefineEndpointVO {
  s.Path = &v
  return s
}

func (s *QueryAPIDefinitionDetailApiDefineEndpointVO) SetCaseSensitive(v string) *QueryAPIDefinitionDetailApiDefineEndpointVO {
  s.CaseSensitive = &v
  return s
}

func (s *QueryAPIDefinitionDetailApiDefineEndpointVO) SetMatchPathVar(v string) *QueryAPIDefinitionDetailApiDefineEndpointVO {
  s.MatchPathVar = &v
  return s
}

type QueryAPIDefinitionDetailApiDefineAuthVO struct {
  // {'en':'Authentication method.
  // NO_NEED:No authentication required
  // SIGN:Key authentication', 'zh_CN':'鉴权方法。
  // NO_NEED：免鉴权
  // SIGN：秘钥对鉴权'}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
  // {'en':'Authentication key.', 'zh_CN':'鉴权秘钥。'}
  AuthKey *string `json:"authKey,omitempty" xml:"authKey,omitempty" require:"true"`
  // {'en':'Authentication parameter location.
  // HTTP_HEADER:Http Header
  // QUERY_STRING:Query String
  // COOKIE:Cookie', 'zh_CN':'鉴权参数位置。
  // HTTP_HEADER：Http Header
  // QUERY_STRING：Query String
  // COOKIE：Cookie'}
  ParamPosition *string `json:"paramPosition,omitempty" xml:"paramPosition,omitempty" require:"true"`
  // {'en':'Authentication parameter name.', 'zh_CN':'鉴权参数名称。'}
  ParamName *string `json:"paramName,omitempty" xml:"paramName,omitempty" require:"true"`
  // {'en':'Authentication token validity period, in seconds.', 'zh_CN':'鉴权有效期，单位秒。'}
  ValidityTime *int64 `json:"validityTime,omitempty" xml:"validityTime,omitempty" require:"true"`
}

func (s QueryAPIDefinitionDetailApiDefineAuthVO) String() string {
  return tea.Prettify(s)
}

func (s QueryAPIDefinitionDetailApiDefineAuthVO) GoString() string {
  return s.String()
}

func (s *QueryAPIDefinitionDetailApiDefineAuthVO) SetType(v string) *QueryAPIDefinitionDetailApiDefineAuthVO {
  s.Type = &v
  return s
}

func (s *QueryAPIDefinitionDetailApiDefineAuthVO) SetAuthKey(v string) *QueryAPIDefinitionDetailApiDefineAuthVO {
  s.AuthKey = &v
  return s
}

func (s *QueryAPIDefinitionDetailApiDefineAuthVO) SetParamPosition(v string) *QueryAPIDefinitionDetailApiDefineAuthVO {
  s.ParamPosition = &v
  return s
}

func (s *QueryAPIDefinitionDetailApiDefineAuthVO) SetParamName(v string) *QueryAPIDefinitionDetailApiDefineAuthVO {
  s.ParamName = &v
  return s
}

func (s *QueryAPIDefinitionDetailApiDefineAuthVO) SetValidityTime(v int64) *QueryAPIDefinitionDetailApiDefineAuthVO {
  s.ValidityTime = &v
  return s
}

type QueryAPIDefinitionDetailApiDefineBodyLimitVO struct {
  // {'en':'Request body restriction switch.
  // ON:ON
  // OFF:OFF', 'zh_CN':'请求body限制开关。
  // ON：开启
  // OFF：关闭'}
  DefendSwitch *string `json:"defendSwitch,omitempty" xml:"defendSwitch,omitempty" require:"true"`
  // {'en':'Content-Type.
  // FORM:FORM
  // JSON:JSON
  // ANY:ANY
  // EMPTY:EMPTY or NON-EXISTENT', 'zh_CN':'Content-Type。
  // FORM：FORM表单
  // JSON：JSON
  // ANY：任意
  // EMPTY：为空或不存在'}
  ContentType *string `json:"contentType,omitempty" xml:"contentType,omitempty" require:"true"`
  // {'en':'Maximum body limit(bytes).', 'zh_CN':'Body最大限制(bytes)。'}
  BodyLimitMax *int64 `json:"bodyLimitMax,omitempty" xml:"bodyLimitMax,omitempty" require:"true"`
  // {'en':'Maximum nesting depth.', 'zh_CN':'最大嵌套层数。'}
  NestMax *int64 `json:"nestMax,omitempty" xml:"nestMax,omitempty" require:"true"`
  // {'en':'Maximum number of parameters for JSON.', 'zh_CN':'JSON最大参数个数。'}
  ParamCountMax *int64 `json:"paramCountMax,omitempty" xml:"paramCountMax,omitempty" require:"true"`
}

func (s QueryAPIDefinitionDetailApiDefineBodyLimitVO) String() string {
  return tea.Prettify(s)
}

func (s QueryAPIDefinitionDetailApiDefineBodyLimitVO) GoString() string {
  return s.String()
}

func (s *QueryAPIDefinitionDetailApiDefineBodyLimitVO) SetDefendSwitch(v string) *QueryAPIDefinitionDetailApiDefineBodyLimitVO {
  s.DefendSwitch = &v
  return s
}

func (s *QueryAPIDefinitionDetailApiDefineBodyLimitVO) SetContentType(v string) *QueryAPIDefinitionDetailApiDefineBodyLimitVO {
  s.ContentType = &v
  return s
}

func (s *QueryAPIDefinitionDetailApiDefineBodyLimitVO) SetBodyLimitMax(v int64) *QueryAPIDefinitionDetailApiDefineBodyLimitVO {
  s.BodyLimitMax = &v
  return s
}

func (s *QueryAPIDefinitionDetailApiDefineBodyLimitVO) SetNestMax(v int64) *QueryAPIDefinitionDetailApiDefineBodyLimitVO {
  s.NestMax = &v
  return s
}

func (s *QueryAPIDefinitionDetailApiDefineBodyLimitVO) SetParamCountMax(v int64) *QueryAPIDefinitionDetailApiDefineBodyLimitVO {
  s.ParamCountMax = &v
  return s
}

type QueryAPIDefinitionDetailApiDefineParamLimitVO struct {
  // {'en':'Basic information.', 'zh_CN':'基础信息。'}
  Basic *QueryAPIDefinitionDetailApiDefineParamLimitBasicVO `json:"basic,omitempty" xml:"basic,omitempty" require:"true"`
  // {'en':'Method attributes.', 'zh_CN':'方法属性。'}
  MethodList []*QueryAPIDefinitionDetailApiDefineParamLimitMethodVO `json:"methodList,omitempty" xml:"methodList,omitempty" require:"true" type:"Repeated"`
}

func (s QueryAPIDefinitionDetailApiDefineParamLimitVO) String() string {
  return tea.Prettify(s)
}

func (s QueryAPIDefinitionDetailApiDefineParamLimitVO) GoString() string {
  return s.String()
}

func (s *QueryAPIDefinitionDetailApiDefineParamLimitVO) SetBasic(v *QueryAPIDefinitionDetailApiDefineParamLimitBasicVO) *QueryAPIDefinitionDetailApiDefineParamLimitVO {
  s.Basic = v
  return s
}

func (s *QueryAPIDefinitionDetailApiDefineParamLimitVO) SetMethodList(v []*QueryAPIDefinitionDetailApiDefineParamLimitMethodVO) *QueryAPIDefinitionDetailApiDefineParamLimitVO {
  s.MethodList = v
  return s
}

type QueryAPIDefinitionDetailApiDefineParamLimitBasicVO struct {
  // {'en':'Parameter limit.
  // ON:ON
  // OFF:OFF', 'zh_CN':'参数限制。
  // ON：ON
  // OFF：OFF'}
  DefendSwitch *string `json:"defendSwitch,omitempty" xml:"defendSwitch,omitempty" require:"true"`
  // {'en':'Query String Parameter Detection Mode.
  // LOOSE:Loose mode - detect some parameters
  // STRICT:Strict mode - checks all parameters', 'zh_CN':'Query String参数检测模式。
  // LOOSE：宽松模式-检测部分参数
  // STRICT：严格模式-检测所有参数'}
  ParamCheckMode *string `json:"paramCheckMode,omitempty" xml:"paramCheckMode,omitempty" require:"true"`
}

func (s QueryAPIDefinitionDetailApiDefineParamLimitBasicVO) String() string {
  return tea.Prettify(s)
}

func (s QueryAPIDefinitionDetailApiDefineParamLimitBasicVO) GoString() string {
  return s.String()
}

func (s *QueryAPIDefinitionDetailApiDefineParamLimitBasicVO) SetDefendSwitch(v string) *QueryAPIDefinitionDetailApiDefineParamLimitBasicVO {
  s.DefendSwitch = &v
  return s
}

func (s *QueryAPIDefinitionDetailApiDefineParamLimitBasicVO) SetParamCheckMode(v string) *QueryAPIDefinitionDetailApiDefineParamLimitBasicVO {
  s.ParamCheckMode = &v
  return s
}

type QueryAPIDefinitionDetailApiDefineParamLimitMethodVO struct {
  // {'en':'Request methods.
  // GET:GET
  // POST:POST
  // DELETE:DELETE
  // PUT:PUT
  // HEAD:HEAD
  // OPTIONS:OPTIONS
  // PATCH:PATCH', 'zh_CN':'请求方法。
  // GET：GET
  // POST：POST
  // DELETE：DELETE
  // PUT：PUT
  // HEAD：HEAD
  // OPTIONS：OPTIONS
  // PATCH：PATCH'}
  Method *string `json:"method,omitempty" xml:"method,omitempty" require:"true"`
  // {'en':'Whether to define body parameters.
  // TRUE:Yes
  // FALSE:No', 'zh_CN':'是否定义body参数。
  // TRUE：是
  // FALSE：否'}
  BodyFlag *string `json:"bodyFlag,omitempty" xml:"bodyFlag,omitempty" require:"true"`
  // {'en':'Body parameter type.
  // JSON:JSON
  // FROM:FROM', 'zh_CN':'body参数类型。
  // JSON：JSON
  // FROM：FROM表单'}
  BodyType *string `json:"bodyType,omitempty" xml:"bodyType,omitempty" require:"true"`
  // {'en':'Normal parameter list.', 'zh_CN':'普通参数数组。'}
  NormalParamList []*QueryAPIDefinitionDetailApiDefineParamLimitNormalVO `json:"normalParamList,omitempty" xml:"normalParamList,omitempty" require:"true" type:"Repeated"`
  // {'en':'FROM form parameter array.', 'zh_CN':'FROM 表单参数数组。'}
  FormParamList []*QueryAPIDefinitionDetailApiDefineParamLimitFormVO `json:"formParamList,omitempty" xml:"formParamList,omitempty" require:"true" type:"Repeated"`
  // {'en':'JSON parameter array.', 'zh_CN':'JSON参数数组。'}
  JsonParam *QueryAPIDefinitionDetailApiDefineParamLimitJsonVO `json:"jsonParam,omitempty" xml:"jsonParam,omitempty" require:"true"`
}

func (s QueryAPIDefinitionDetailApiDefineParamLimitMethodVO) String() string {
  return tea.Prettify(s)
}

func (s QueryAPIDefinitionDetailApiDefineParamLimitMethodVO) GoString() string {
  return s.String()
}

func (s *QueryAPIDefinitionDetailApiDefineParamLimitMethodVO) SetMethod(v string) *QueryAPIDefinitionDetailApiDefineParamLimitMethodVO {
  s.Method = &v
  return s
}

func (s *QueryAPIDefinitionDetailApiDefineParamLimitMethodVO) SetBodyFlag(v string) *QueryAPIDefinitionDetailApiDefineParamLimitMethodVO {
  s.BodyFlag = &v
  return s
}

func (s *QueryAPIDefinitionDetailApiDefineParamLimitMethodVO) SetBodyType(v string) *QueryAPIDefinitionDetailApiDefineParamLimitMethodVO {
  s.BodyType = &v
  return s
}

func (s *QueryAPIDefinitionDetailApiDefineParamLimitMethodVO) SetNormalParamList(v []*QueryAPIDefinitionDetailApiDefineParamLimitNormalVO) *QueryAPIDefinitionDetailApiDefineParamLimitMethodVO {
  s.NormalParamList = v
  return s
}

func (s *QueryAPIDefinitionDetailApiDefineParamLimitMethodVO) SetFormParamList(v []*QueryAPIDefinitionDetailApiDefineParamLimitFormVO) *QueryAPIDefinitionDetailApiDefineParamLimitMethodVO {
  s.FormParamList = v
  return s
}

func (s *QueryAPIDefinitionDetailApiDefineParamLimitMethodVO) SetJsonParam(v *QueryAPIDefinitionDetailApiDefineParamLimitJsonVO) *QueryAPIDefinitionDetailApiDefineParamLimitMethodVO {
  s.JsonParam = v
  return s
}

type QueryAPIDefinitionDetailApiDefineParamLimitNormalVO struct {
  // {'en':'Request methods.
  // GET:GET
  // POST:POST
  // DELETE:DELETE
  // PUT:PUT
  // HEAD:HEAD
  // OPTIONS:OPTIONS
  // PATCH:PATCH', 'zh_CN':'请求方法。
  // GET：GET
  // POST：POST
  // DELETE：DELETE
  // PUT：PUT
  // HEAD：HEAD
  // OPTIONS：OPTIONS
  // PATCH：PATCH'}
  Method *string `json:"method,omitempty" xml:"method,omitempty" require:"true"`
  // {'en':'Parameter name.', 'zh_CN':'参数名称。'}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {'en':'Parameter position.
  // HTTP_HEADER:Http Header
  // QUERY_STRING:Query String
  // COOKIE:Cookie
  // PATH_PARAMS:Path', 'zh_CN':'参数变量。
  // HTTP_HEADER：Http Header
  // QUERY_STRING：Query String
  // COOKIE：Cookie
  // PATH_PARAMS：路径变量'}
  ParamPosition *string `json:"paramPosition,omitempty" xml:"paramPosition,omitempty" require:"true"`
  // {'en':'Parameter type.
  // Integer:Integer
  // Number:Number
  // String:String
  // Boolean:Boolean
  // Enumeration:Enumeration', 'zh_CN':'参数类型。
  // Integer：整数
  // Number：数字
  // String：字符串
  // Boolean：布尔
  // Enumeration：枚举'}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
  // {'en':'Minimum value/Minimum length.', 'zh_CN':'最小值、最小长度。'}
  MinVal *string `json:"minVal,omitempty" xml:"minVal,omitempty" require:"true"`
  // {'en':'Maximum value/Maximum length.', 'zh_CN':'最大值、最大长度。'}
  MaxVal *string `json:"maxVal,omitempty" xml:"maxVal,omitempty" require:"true"`
  // {'en':'Required.
  // TRUE:Yes
  // FALSE:No', 'zh_CN':'必带参数。
  // TRUE：是
  // FALSE：否'}
  Required *string `json:"required,omitempty" xml:"required,omitempty" require:"true"`
  // {'en':'Content.', 'zh_CN':'内容。'}
  Content *string `json:"content,omitempty" xml:"content,omitempty" require:"true"`
  // {'en':'Description.', 'zh_CN':'备注。'}
  Remark *string `json:"remark,omitempty" xml:"remark,omitempty" require:"true"`
}

func (s QueryAPIDefinitionDetailApiDefineParamLimitNormalVO) String() string {
  return tea.Prettify(s)
}

func (s QueryAPIDefinitionDetailApiDefineParamLimitNormalVO) GoString() string {
  return s.String()
}

func (s *QueryAPIDefinitionDetailApiDefineParamLimitNormalVO) SetMethod(v string) *QueryAPIDefinitionDetailApiDefineParamLimitNormalVO {
  s.Method = &v
  return s
}

func (s *QueryAPIDefinitionDetailApiDefineParamLimitNormalVO) SetName(v string) *QueryAPIDefinitionDetailApiDefineParamLimitNormalVO {
  s.Name = &v
  return s
}

func (s *QueryAPIDefinitionDetailApiDefineParamLimitNormalVO) SetParamPosition(v string) *QueryAPIDefinitionDetailApiDefineParamLimitNormalVO {
  s.ParamPosition = &v
  return s
}

func (s *QueryAPIDefinitionDetailApiDefineParamLimitNormalVO) SetType(v string) *QueryAPIDefinitionDetailApiDefineParamLimitNormalVO {
  s.Type = &v
  return s
}

func (s *QueryAPIDefinitionDetailApiDefineParamLimitNormalVO) SetMinVal(v string) *QueryAPIDefinitionDetailApiDefineParamLimitNormalVO {
  s.MinVal = &v
  return s
}

func (s *QueryAPIDefinitionDetailApiDefineParamLimitNormalVO) SetMaxVal(v string) *QueryAPIDefinitionDetailApiDefineParamLimitNormalVO {
  s.MaxVal = &v
  return s
}

func (s *QueryAPIDefinitionDetailApiDefineParamLimitNormalVO) SetRequired(v string) *QueryAPIDefinitionDetailApiDefineParamLimitNormalVO {
  s.Required = &v
  return s
}

func (s *QueryAPIDefinitionDetailApiDefineParamLimitNormalVO) SetContent(v string) *QueryAPIDefinitionDetailApiDefineParamLimitNormalVO {
  s.Content = &v
  return s
}

func (s *QueryAPIDefinitionDetailApiDefineParamLimitNormalVO) SetRemark(v string) *QueryAPIDefinitionDetailApiDefineParamLimitNormalVO {
  s.Remark = &v
  return s
}

type QueryAPIDefinitionDetailApiDefineParamLimitFormVO struct {
  // {'en':'Request methods.
  // POST:POST
  // PUT:PUT
  // PATCH:PATCH', 'zh_CN':'请求方法。
  // POST：POST
  // PUT：PUT
  // PATCH：PATCH'}
  Method *string `json:"method,omitempty" xml:"method,omitempty" require:"true"`
  // {'en':'Parameter name.', 'zh_CN':'参数名称。'}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {'en':'Parameter type.
  // Integer:Integer
  // Number:Number
  // String:String
  // Boolean:Boolean
  // Enumeration:Enumeration', 'zh_CN':'参数类型。
  // Integer：整数
  // Number：数字
  // String：字符串
  // Boolean：布尔
  // Enumeration：枚举'}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
  // {'en':'Minimum value/Minimum length.', 'zh_CN':'最小值、最小长度。'}
  MinVal *string `json:"minVal,omitempty" xml:"minVal,omitempty" require:"true"`
  // {'en':'Maximum value/Maximum length.', 'zh_CN':'最大值、最大长度。'}
  MaxVal *string `json:"maxVal,omitempty" xml:"maxVal,omitempty" require:"true"`
  // {'en':'Required.
  // TRUE:Yes
  // FALSE:No', 'zh_CN':'必带参数。
  // TRUE：是
  // FALSE：否'}
  Required *string `json:"required,omitempty" xml:"required,omitempty" require:"true"`
  // {'en':'Content.', 'zh_CN':'内容。'}
  Content *string `json:"content,omitempty" xml:"content,omitempty" require:"true"`
  // {'en':'Description.', 'zh_CN':'备注。'}
  Remark *string `json:"remark,omitempty" xml:"remark,omitempty" require:"true"`
}

func (s QueryAPIDefinitionDetailApiDefineParamLimitFormVO) String() string {
  return tea.Prettify(s)
}

func (s QueryAPIDefinitionDetailApiDefineParamLimitFormVO) GoString() string {
  return s.String()
}

func (s *QueryAPIDefinitionDetailApiDefineParamLimitFormVO) SetMethod(v string) *QueryAPIDefinitionDetailApiDefineParamLimitFormVO {
  s.Method = &v
  return s
}

func (s *QueryAPIDefinitionDetailApiDefineParamLimitFormVO) SetName(v string) *QueryAPIDefinitionDetailApiDefineParamLimitFormVO {
  s.Name = &v
  return s
}

func (s *QueryAPIDefinitionDetailApiDefineParamLimitFormVO) SetType(v string) *QueryAPIDefinitionDetailApiDefineParamLimitFormVO {
  s.Type = &v
  return s
}

func (s *QueryAPIDefinitionDetailApiDefineParamLimitFormVO) SetMinVal(v string) *QueryAPIDefinitionDetailApiDefineParamLimitFormVO {
  s.MinVal = &v
  return s
}

func (s *QueryAPIDefinitionDetailApiDefineParamLimitFormVO) SetMaxVal(v string) *QueryAPIDefinitionDetailApiDefineParamLimitFormVO {
  s.MaxVal = &v
  return s
}

func (s *QueryAPIDefinitionDetailApiDefineParamLimitFormVO) SetRequired(v string) *QueryAPIDefinitionDetailApiDefineParamLimitFormVO {
  s.Required = &v
  return s
}

func (s *QueryAPIDefinitionDetailApiDefineParamLimitFormVO) SetContent(v string) *QueryAPIDefinitionDetailApiDefineParamLimitFormVO {
  s.Content = &v
  return s
}

func (s *QueryAPIDefinitionDetailApiDefineParamLimitFormVO) SetRemark(v string) *QueryAPIDefinitionDetailApiDefineParamLimitFormVO {
  s.Remark = &v
  return s
}

type QueryAPIDefinitionDetailApiDefineParamLimitJsonVO struct {
  // {'en':'Request methods.
  // POST:POST
  // PUT:PUT
  // PATCH:PATCH', 'zh_CN':'请求方法。
  // POST：POST
  // PUT：PUT
  // PATCH：PATCH'}
  Method *string `json:"method,omitempty" xml:"method,omitempty" require:"true"`
  // {'en':'Parameter Name.', 'zh_CN':'参数名称。'}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {'en':'Level.', 'zh_CN':'层级。'}
  Level *int32 `json:"level,omitempty" xml:"level,omitempty" require:"true"`
  // {'en':'Parameter Type.
  // Integer:Integer
  // Number:Number
  // String:String
  // Boolean:Boolean
  // Enumeration:Enumeration
  // Array:Array
  // Json:JSON Object', 'zh_CN':'参数类型。
  // Integer：整数
  // Number：数字
  // String：字符串
  // Boolean：布尔
  // Enumeration：枚举
  // Array：数组
  // Json：JSON对象'}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
  // {'en':'Minimum value/Minimum length.', 'zh_CN':'最小值、最小长度。'}
  MinVal *string `json:"minVal,omitempty" xml:"minVal,omitempty" require:"true"`
  // {'en':'Maximum value/Maximum length.', 'zh_CN':'最大值、最大长度。'}
  MaxVal *string `json:"maxVal,omitempty" xml:"maxVal,omitempty" require:"true"`
  // {'en':'Required.
  // TRUE:Yes
  // FALSE:No', 'zh_CN':'必带参数。
  // TRUE：是
  // FALSE：否'}
  Required *string `json:"required,omitempty" xml:"required,omitempty" require:"true"`
  // {'en':'Content.', 'zh_CN':'内容。'}
  Content *string `json:"content,omitempty" xml:"content,omitempty" require:"true"`
  // {'en':'Child node array, required when type= Array.', 'zh_CN':'子节点数组，type = Array 时必填。'}
  Children []*QueryAPIDefinitionDetailApiDefineParamLimitJsonChildrenVO `json:"children,omitempty" xml:"children,omitempty" require:"true" type:"Repeated"`
}

func (s QueryAPIDefinitionDetailApiDefineParamLimitJsonVO) String() string {
  return tea.Prettify(s)
}

func (s QueryAPIDefinitionDetailApiDefineParamLimitJsonVO) GoString() string {
  return s.String()
}

func (s *QueryAPIDefinitionDetailApiDefineParamLimitJsonVO) SetMethod(v string) *QueryAPIDefinitionDetailApiDefineParamLimitJsonVO {
  s.Method = &v
  return s
}

func (s *QueryAPIDefinitionDetailApiDefineParamLimitJsonVO) SetName(v string) *QueryAPIDefinitionDetailApiDefineParamLimitJsonVO {
  s.Name = &v
  return s
}

func (s *QueryAPIDefinitionDetailApiDefineParamLimitJsonVO) SetLevel(v int32) *QueryAPIDefinitionDetailApiDefineParamLimitJsonVO {
  s.Level = &v
  return s
}

func (s *QueryAPIDefinitionDetailApiDefineParamLimitJsonVO) SetType(v string) *QueryAPIDefinitionDetailApiDefineParamLimitJsonVO {
  s.Type = &v
  return s
}

func (s *QueryAPIDefinitionDetailApiDefineParamLimitJsonVO) SetMinVal(v string) *QueryAPIDefinitionDetailApiDefineParamLimitJsonVO {
  s.MinVal = &v
  return s
}

func (s *QueryAPIDefinitionDetailApiDefineParamLimitJsonVO) SetMaxVal(v string) *QueryAPIDefinitionDetailApiDefineParamLimitJsonVO {
  s.MaxVal = &v
  return s
}

func (s *QueryAPIDefinitionDetailApiDefineParamLimitJsonVO) SetRequired(v string) *QueryAPIDefinitionDetailApiDefineParamLimitJsonVO {
  s.Required = &v
  return s
}

func (s *QueryAPIDefinitionDetailApiDefineParamLimitJsonVO) SetContent(v string) *QueryAPIDefinitionDetailApiDefineParamLimitJsonVO {
  s.Content = &v
  return s
}

func (s *QueryAPIDefinitionDetailApiDefineParamLimitJsonVO) SetChildren(v []*QueryAPIDefinitionDetailApiDefineParamLimitJsonChildrenVO) *QueryAPIDefinitionDetailApiDefineParamLimitJsonVO {
  s.Children = v
  return s
}

type QueryAPIDefinitionDetailApiDefineParamLimitJsonChildrenVO struct {
  // {'en':'Request methods.
  // POST:POST
  // PUT:PUT
  // PATCH:PATCH', 'zh_CN':'请求方法。
  // POST：POST
  // PUT：PUT
  // PATCH：PATCH'}
  Method *string `json:"method,omitempty" xml:"method,omitempty" require:"true"`
  // {'en':'Parameter Name.', 'zh_CN':'参数名称。'}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {'en':'Level.', 'zh_CN':'层级。'}
  Level *int32 `json:"level,omitempty" xml:"level,omitempty" require:"true"`
  // {'en':'Parameter Type.
  // Integer:Integer
  // Number:Number
  // String:String
  // Boolean:Boolean
  // Enumeration:Enumeration
  // Array:Array
  // Json:JSON Object', 'zh_CN':'参数类型。
  // Integer：整数
  // Number：数字
  // String：字符串
  // Boolean：布尔
  // Enumeration：枚举
  // Array：数组
  // Json：JSON对象'}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
  // {'en':'Minimum value/Minimum length.', 'zh_CN':'最小值、最小长度。'}
  MinVal *string `json:"minVal,omitempty" xml:"minVal,omitempty" require:"true"`
  // {'en':'Maximum value/Maximum length.', 'zh_CN':'最大值、最大长度。'}
  MaxVal *string `json:"maxVal,omitempty" xml:"maxVal,omitempty" require:"true"`
  // {'en':'Required.
  // TRUE:Yes
  // FALSE:No', 'zh_CN':'必带参数。
  // TRUE：是
  // FALSE：否'}
  Required *string `json:"required,omitempty" xml:"required,omitempty" require:"true"`
  // {'en':'Content.', 'zh_CN':'内容。'}
  Content *string `json:"content,omitempty" xml:"content,omitempty" require:"true"`
  // {'en':'Child node array (JSON string), the structure is consistent with the parent node, required when type= Array.', 'zh_CN':'子节点数组（JSON字符串），结构与父结点一致，type = Array 时必填。'}
  Children []*string `json:"children,omitempty" xml:"children,omitempty" require:"true" type:"Repeated"`
}

func (s QueryAPIDefinitionDetailApiDefineParamLimitJsonChildrenVO) String() string {
  return tea.Prettify(s)
}

func (s QueryAPIDefinitionDetailApiDefineParamLimitJsonChildrenVO) GoString() string {
  return s.String()
}

func (s *QueryAPIDefinitionDetailApiDefineParamLimitJsonChildrenVO) SetMethod(v string) *QueryAPIDefinitionDetailApiDefineParamLimitJsonChildrenVO {
  s.Method = &v
  return s
}

func (s *QueryAPIDefinitionDetailApiDefineParamLimitJsonChildrenVO) SetName(v string) *QueryAPIDefinitionDetailApiDefineParamLimitJsonChildrenVO {
  s.Name = &v
  return s
}

func (s *QueryAPIDefinitionDetailApiDefineParamLimitJsonChildrenVO) SetLevel(v int32) *QueryAPIDefinitionDetailApiDefineParamLimitJsonChildrenVO {
  s.Level = &v
  return s
}

func (s *QueryAPIDefinitionDetailApiDefineParamLimitJsonChildrenVO) SetType(v string) *QueryAPIDefinitionDetailApiDefineParamLimitJsonChildrenVO {
  s.Type = &v
  return s
}

func (s *QueryAPIDefinitionDetailApiDefineParamLimitJsonChildrenVO) SetMinVal(v string) *QueryAPIDefinitionDetailApiDefineParamLimitJsonChildrenVO {
  s.MinVal = &v
  return s
}

func (s *QueryAPIDefinitionDetailApiDefineParamLimitJsonChildrenVO) SetMaxVal(v string) *QueryAPIDefinitionDetailApiDefineParamLimitJsonChildrenVO {
  s.MaxVal = &v
  return s
}

func (s *QueryAPIDefinitionDetailApiDefineParamLimitJsonChildrenVO) SetRequired(v string) *QueryAPIDefinitionDetailApiDefineParamLimitJsonChildrenVO {
  s.Required = &v
  return s
}

func (s *QueryAPIDefinitionDetailApiDefineParamLimitJsonChildrenVO) SetContent(v string) *QueryAPIDefinitionDetailApiDefineParamLimitJsonChildrenVO {
  s.Content = &v
  return s
}

func (s *QueryAPIDefinitionDetailApiDefineParamLimitJsonChildrenVO) SetChildren(v []*string) *QueryAPIDefinitionDetailApiDefineParamLimitJsonChildrenVO {
  s.Children = v
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

type QueryAPIDefinitionDetailRequestHeader struct {
  // {"zh_CN":"安全服务类型。有使用多个不同的安全服务时，需要填写具体的服务类型。","en":"Security service type. Please enter a specific service type, if you purchase multiple security services.","dictionary":"belong=WAAP-MS-Ext|dict=waap_serviceType"}
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

type QueryAPIDefinitionDetailResponseHeader struct {
}

func (s QueryAPIDefinitionDetailResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryAPIDefinitionDetailResponseHeader) GoString() string {
  return s.String()
}




type FeedbackWrongAPIAssetDiscoveryRequest struct {
  // {"en":"API discovery ID.", "zh_CN":"API发现ID。"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {"en":"Feedback, maximum 200 characters.", "zh_CN":"反馈意见，最多200个字符。"}
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

type FeedbackWrongAPIAssetDiscoveryResponse struct {
  // {'en':'Please refer to the error code for exceptions.', 'zh_CN':'请参照错误码。','dictionary':'belong=WAAP-MS-Ext|dict=waap_retCodeEnum'}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Description.", "zh_CN":"描述信息。"}
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

type FeedbackWrongAPIAssetDiscoveryRequestHeader struct {
  // {"zh_CN":"安全服务类型。有使用多个不同的安全服务时，需要填写具体的服务类型。","en":"Security service type. Please enter a specific service type, if you purchase multiple security services.","dictionary":"belong=WAAP-MS-Ext|dict=waap_serviceType"}
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

type FeedbackWrongAPIAssetDiscoveryResponseHeader struct {
}

func (s FeedbackWrongAPIAssetDiscoveryResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s FeedbackWrongAPIAssetDiscoveryResponseHeader) GoString() string {
  return s.String()
}




type UpdateAPIDefinitionRequest struct {
  // {'en':'Basic information.', 'zh_CN':'基础信息。'}
  Basic *UpdateAPIDefinitionApiDefineBasicEditDTO `json:"basic,omitempty" xml:"basic,omitempty" require:"true"`
  // {'en':'Endpoint information.', 'zh_CN':'端点信息。'}
  Endpoint *UpdateAPIDefinitionApiDefineEndpointEditDTO `json:"endpoint,omitempty" xml:"endpoint,omitempty"`
  // {'en':'Authentication configuration.', 'zh_CN':'鉴权配置。'}
  Auth *UpdateAPIDefinitionApiDefineAuthEditDTO `json:"auth,omitempty" xml:"auth,omitempty"`
  // {'en':'Body restrictions.', 'zh_CN':'body限制。'}
  BodyLimit *UpdateAPIDefinitionApiDefineBodyLimitEditDTO `json:"bodyLimit,omitempty" xml:"bodyLimit,omitempty"`
  // {'en':'Parameter limit.', 'zh_CN':'参数限制。'}
  ParamLimit *UpdateAPIDefinitionApiDefineParamLimitEditDTO `json:"paramLimit,omitempty" xml:"paramLimit,omitempty"`
}

func (s UpdateAPIDefinitionRequest) String() string {
  return tea.Prettify(s)
}

func (s UpdateAPIDefinitionRequest) GoString() string {
  return s.String()
}

func (s *UpdateAPIDefinitionRequest) SetBasic(v *UpdateAPIDefinitionApiDefineBasicEditDTO) *UpdateAPIDefinitionRequest {
  s.Basic = v
  return s
}

func (s *UpdateAPIDefinitionRequest) SetEndpoint(v *UpdateAPIDefinitionApiDefineEndpointEditDTO) *UpdateAPIDefinitionRequest {
  s.Endpoint = v
  return s
}

func (s *UpdateAPIDefinitionRequest) SetAuth(v *UpdateAPIDefinitionApiDefineAuthEditDTO) *UpdateAPIDefinitionRequest {
  s.Auth = v
  return s
}

func (s *UpdateAPIDefinitionRequest) SetBodyLimit(v *UpdateAPIDefinitionApiDefineBodyLimitEditDTO) *UpdateAPIDefinitionRequest {
  s.BodyLimit = v
  return s
}

func (s *UpdateAPIDefinitionRequest) SetParamLimit(v *UpdateAPIDefinitionApiDefineParamLimitEditDTO) *UpdateAPIDefinitionRequest {
  s.ParamLimit = v
  return s
}

type UpdateAPIDefinitionApiDefineBasicEditDTO struct {
  // {'en':'API define ID.', 'zh_CN':'API定义ID。'}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {'en':'API name, maximum 200 characters.', 'zh_CN':'API名称，最多200个字符。'}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {'en':'API groups, maximum 64 characters.', 'zh_CN':'API分组，最多64个字符。'}
  GroupName *string `json:"groupName,omitempty" xml:"groupName,omitempty" require:"true"`
  // {'en':'Description, maximum 200 characters.', 'zh_CN':'备注，最多200个字符。'}
  Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
}

func (s UpdateAPIDefinitionApiDefineBasicEditDTO) String() string {
  return tea.Prettify(s)
}

func (s UpdateAPIDefinitionApiDefineBasicEditDTO) GoString() string {
  return s.String()
}

func (s *UpdateAPIDefinitionApiDefineBasicEditDTO) SetId(v string) *UpdateAPIDefinitionApiDefineBasicEditDTO {
  s.Id = &v
  return s
}

func (s *UpdateAPIDefinitionApiDefineBasicEditDTO) SetName(v string) *UpdateAPIDefinitionApiDefineBasicEditDTO {
  s.Name = &v
  return s
}

func (s *UpdateAPIDefinitionApiDefineBasicEditDTO) SetGroupName(v string) *UpdateAPIDefinitionApiDefineBasicEditDTO {
  s.GroupName = &v
  return s
}

func (s *UpdateAPIDefinitionApiDefineBasicEditDTO) SetRemark(v string) *UpdateAPIDefinitionApiDefineBasicEditDTO {
  s.Remark = &v
  return s
}

type UpdateAPIDefinitionApiDefineEndpointEditDTO struct {
  // {'en':'Request methods,Multiple separated by ; sign.
  // GET:GET, configurable parameter limits
  // POST:POST, configurable parameter limits
  // DELETE:DELETE, configurable parameter limits
  // UPDATE:UPDATE
  // PUT:PUT, configurable parameter limits
  // HEAD:HEAD, configurable parameter limits
  // CONNECT:CONNECT
  // OPTIONS:OPTIONS, configurable parameter limits
  // COPY:COPY
  // LOCK:LOCK
  // UNLOCK:UNLOCK
  // TRACE:TRACE
  // PATCH:PATCH, configurable parameter limits
  // PROPFIND:PROPFIND
  // MKCOL:MKCOL
  // MOVE:MOVE', 'zh_CN':'请求方法。多个以 ; 号分隔。
  // GET：GET，可配置参数限制
  // POST：POST，可配置参数限制
  // DELETE：DELETE，可配置参数限制
  // UPDATE：UPDATE
  // PUT：PUT，可配置参数限制
  // HEAD：HEAD，可配置参数限制
  // CONNECT：CONNECT
  // OPTIONS：OPTIONS，可配置参数限制
  // COPY：COPY
  // LOCK：LOCK
  // UNLOCK：UNLOCK
  // TRACE：TRACE
  // PATCH：PATCH，可配置参数限制
  // PROPFIND：PROPFIND
  // MKCOL：MKCOL
  // MOVE：MOVE'}
  RequestMethod *string `json:"requestMethod,omitempty" xml:"requestMethod,omitempty" require:"true"`
  // {'en':'Case sensitive.
  // TRUE:Yes
  // FALSE:No', 'zh_CN':'大小写是否敏感。
  // TRUE：是
  // FALSE：否'}
  CaseSensitive *string `json:"caseSensitive,omitempty" xml:"caseSensitive,omitempty" require:"true"`
}

func (s UpdateAPIDefinitionApiDefineEndpointEditDTO) String() string {
  return tea.Prettify(s)
}

func (s UpdateAPIDefinitionApiDefineEndpointEditDTO) GoString() string {
  return s.String()
}

func (s *UpdateAPIDefinitionApiDefineEndpointEditDTO) SetRequestMethod(v string) *UpdateAPIDefinitionApiDefineEndpointEditDTO {
  s.RequestMethod = &v
  return s
}

func (s *UpdateAPIDefinitionApiDefineEndpointEditDTO) SetCaseSensitive(v string) *UpdateAPIDefinitionApiDefineEndpointEditDTO {
  s.CaseSensitive = &v
  return s
}

type UpdateAPIDefinitionApiDefineAuthEditDTO struct {
  // {'en':'Authentication method.
  // NO_NEED:No authentication required
  // SIGN:Key authentication', 'zh_CN':'鉴权方法。
  // NO_NEED：免鉴权
  // SIGN：秘钥对鉴权'}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
  // {'en':'Authentication key, type = SIGN is required. If it has been set, it will be ignored. The format is a 16-digit string containing uppercase and lowercase letters and numbers. Example: gjZkg2E1uNkXBDxj.', 'zh_CN':'鉴权秘钥，type = SIGN 是必填，如已设置则忽略，格式为16位含大小写字母与数字字符串，示例：gjZkg2E1uNkXBDxj。'}
  AuthKey *string `json:"authKey,omitempty" xml:"authKey,omitempty"`
  // {'en':'Authentication parameter location, type = SIGN is required.
  // HTTP_HEADER:Http Header
  // QUERY_STRING:Query String
  // COOKIE:Cookie', 'zh_CN':'鉴权参数位置，type = SIGN 时必填。
  // HTTP_HEADER：Http Header
  // QUERY_STRING：Query String
  // COOKIE：Cookie'}
  ParamPosition *string `json:"paramPosition,omitempty" xml:"paramPosition,omitempty"`
  // {'en':'Authentication parameter name, type = SIGN is required.', 'zh_CN':'鉴权参数名称，type = SIGN 时必填。'}
  ParamName *string `json:"paramName,omitempty" xml:"paramName,omitempty"`
  // {'en':'Authentication token validity period, in seconds, type = SIGN is required.', 'zh_CN':'鉴权有效期，单位秒，type = SIGN 时必填。'}
  ValidityTime *int64 `json:"validityTime,omitempty" xml:"validityTime,omitempty"`
}

func (s UpdateAPIDefinitionApiDefineAuthEditDTO) String() string {
  return tea.Prettify(s)
}

func (s UpdateAPIDefinitionApiDefineAuthEditDTO) GoString() string {
  return s.String()
}

func (s *UpdateAPIDefinitionApiDefineAuthEditDTO) SetType(v string) *UpdateAPIDefinitionApiDefineAuthEditDTO {
  s.Type = &v
  return s
}

func (s *UpdateAPIDefinitionApiDefineAuthEditDTO) SetAuthKey(v string) *UpdateAPIDefinitionApiDefineAuthEditDTO {
  s.AuthKey = &v
  return s
}

func (s *UpdateAPIDefinitionApiDefineAuthEditDTO) SetParamPosition(v string) *UpdateAPIDefinitionApiDefineAuthEditDTO {
  s.ParamPosition = &v
  return s
}

func (s *UpdateAPIDefinitionApiDefineAuthEditDTO) SetParamName(v string) *UpdateAPIDefinitionApiDefineAuthEditDTO {
  s.ParamName = &v
  return s
}

func (s *UpdateAPIDefinitionApiDefineAuthEditDTO) SetValidityTime(v int64) *UpdateAPIDefinitionApiDefineAuthEditDTO {
  s.ValidityTime = &v
  return s
}

type UpdateAPIDefinitionApiDefineBodyLimitEditDTO struct {
  // {'en':'Request body restriction switch. default value:OFF.
  // ON:On
  // OFF:Off', 'zh_CN':'请求body限制开关。默认值：关。
  // ON：开启
  // OFF：关闭'}
  DefendSwitch *string `json:"defendSwitch,omitempty" xml:"defendSwitch,omitempty"`
  // {'en':'Content-Type, required when defendSwitch = ON.
  // FORM:FORM
  // JSON:JSON
  // ANY:ANY
  // EMPTY:EMPTY or NON-EXISTENT', 'zh_CN':'Content-Type，defendSwitch = ON 时必填。
  // FORM：FORM表单
  // JSON：JSON
  // ANY：任意
  // EMPTY：为空或不存在'}
  ContentType *string `json:"contentType,omitempty" xml:"contentType,omitempty"`
  // {'en':'Maximum body limit(bytes).', 'zh_CN':'Body最大限制(bytes)。'}
  BodyLimitMax *int64 `json:"bodyLimitMax,omitempty" xml:"bodyLimitMax,omitempty"`
  // {'en':'Maximum nesting depth, enter the maximum allowable JSON nesting depth in the request body.', 'zh_CN':'最大嵌套层数，输入允许的请求body中JSON嵌套层数最大值。'}
  NestMax *int64 `json:"nestMax,omitempty" xml:"nestMax,omitempty"`
  // {'en':'Maximum number of parameters for JSON, enter the maximum number of JSON parameters allowed in the request body.', 'zh_CN':'JSON最大参数个数，输入允许的请求body中JSON参数个数的最大值。'}
  ParamCountMax *int64 `json:"paramCountMax,omitempty" xml:"paramCountMax,omitempty"`
}

func (s UpdateAPIDefinitionApiDefineBodyLimitEditDTO) String() string {
  return tea.Prettify(s)
}

func (s UpdateAPIDefinitionApiDefineBodyLimitEditDTO) GoString() string {
  return s.String()
}

func (s *UpdateAPIDefinitionApiDefineBodyLimitEditDTO) SetDefendSwitch(v string) *UpdateAPIDefinitionApiDefineBodyLimitEditDTO {
  s.DefendSwitch = &v
  return s
}

func (s *UpdateAPIDefinitionApiDefineBodyLimitEditDTO) SetContentType(v string) *UpdateAPIDefinitionApiDefineBodyLimitEditDTO {
  s.ContentType = &v
  return s
}

func (s *UpdateAPIDefinitionApiDefineBodyLimitEditDTO) SetBodyLimitMax(v int64) *UpdateAPIDefinitionApiDefineBodyLimitEditDTO {
  s.BodyLimitMax = &v
  return s
}

func (s *UpdateAPIDefinitionApiDefineBodyLimitEditDTO) SetNestMax(v int64) *UpdateAPIDefinitionApiDefineBodyLimitEditDTO {
  s.NestMax = &v
  return s
}

func (s *UpdateAPIDefinitionApiDefineBodyLimitEditDTO) SetParamCountMax(v int64) *UpdateAPIDefinitionApiDefineBodyLimitEditDTO {
  s.ParamCountMax = &v
  return s
}

type UpdateAPIDefinitionApiDefineParamLimitEditDTO struct {
  // {'en':'Basic information.', 'zh_CN':'基础信息。'}
  Basic *UpdateAPIDefinitionApiDefineParamLimitBasicEditDTO `json:"basic,omitempty" xml:"basic,omitempty" require:"true"`
  // {'en':'There are request methods for changing configurations.
  // GET:GET
  // POST:POST
  // DELETE:DELETE
  // PUT:PUT
  // HEAD:HEAD
  // OPTIONS:OPTIONS
  // PATCH:PATCH', 'zh_CN':'有变更配置的请求方法。
  // GET：GET
  // POST：POST
  // DELETE：DELETE
  // PUT：PUT
  // HEAD：HEAD
  // OPTIONS：OPTIONS
  // PATCH：PATCH'}
  ChangeParamLimitMethodList []*string `json:"changeParamLimitMethodList,omitempty" xml:"changeParamLimitMethodList,omitempty" type:"Repeated"`
  // {'en':'Method attributes.', 'zh_CN':'方法属性。'}
  MethodList []*UpdateAPIDefinitionApiDefineParamLimitMethodEditDTO `json:"methodList,omitempty" xml:"methodList,omitempty" type:"Repeated"`
}

func (s UpdateAPIDefinitionApiDefineParamLimitEditDTO) String() string {
  return tea.Prettify(s)
}

func (s UpdateAPIDefinitionApiDefineParamLimitEditDTO) GoString() string {
  return s.String()
}

func (s *UpdateAPIDefinitionApiDefineParamLimitEditDTO) SetBasic(v *UpdateAPIDefinitionApiDefineParamLimitBasicEditDTO) *UpdateAPIDefinitionApiDefineParamLimitEditDTO {
  s.Basic = v
  return s
}

func (s *UpdateAPIDefinitionApiDefineParamLimitEditDTO) SetChangeParamLimitMethodList(v []*string) *UpdateAPIDefinitionApiDefineParamLimitEditDTO {
  s.ChangeParamLimitMethodList = v
  return s
}

func (s *UpdateAPIDefinitionApiDefineParamLimitEditDTO) SetMethodList(v []*UpdateAPIDefinitionApiDefineParamLimitMethodEditDTO) *UpdateAPIDefinitionApiDefineParamLimitEditDTO {
  s.MethodList = v
  return s
}

type UpdateAPIDefinitionApiDefineParamLimitBasicEditDTO struct {
  // {'en':'Parameter limit.
  // ON:On
  // OFF:Off', 'zh_CN':'参数限制。
  // ON：开启
  // OFF：关闭'}
  DefendSwitch *string `json:"defendSwitch,omitempty" xml:"defendSwitch,omitempty"`
  // {'en':'Query String Parameter Detection Mode, required when defendSwitch = ON.
  // LOOSE:Loose mode - detect some parameters
  // STRICT:Strict mode - checks all parameters', 'zh_CN':'Query String参数检测模式，defendSwitch = ON 时必填。
  // LOOSE：宽松模式-检测部分参数
  // STRICT：严格模式-检测所有参数'}
  ParamCheckMode *string `json:"paramCheckMode,omitempty" xml:"paramCheckMode,omitempty"`
}

func (s UpdateAPIDefinitionApiDefineParamLimitBasicEditDTO) String() string {
  return tea.Prettify(s)
}

func (s UpdateAPIDefinitionApiDefineParamLimitBasicEditDTO) GoString() string {
  return s.String()
}

func (s *UpdateAPIDefinitionApiDefineParamLimitBasicEditDTO) SetDefendSwitch(v string) *UpdateAPIDefinitionApiDefineParamLimitBasicEditDTO {
  s.DefendSwitch = &v
  return s
}

func (s *UpdateAPIDefinitionApiDefineParamLimitBasicEditDTO) SetParamCheckMode(v string) *UpdateAPIDefinitionApiDefineParamLimitBasicEditDTO {
  s.ParamCheckMode = &v
  return s
}

type UpdateAPIDefinitionApiDefineParamLimitMethodEditDTO struct {
  // {'en':'Request methods.
  // GET:GET
  // POST:POST
  // DELETE:DELETE
  // PUT:PUT
  // HEAD:HEAD
  // OPTIONS:OPTIONS
  // PATCH:PATCH', 'zh_CN':'请求方法。
  // GET：GET
  // POST：POST
  // DELETE：DELETE
  // PUT：PUT
  // HEAD：HEAD
  // OPTIONS：OPTIONS
  // PATCH：PATCH'}
  Method *string `json:"method,omitempty" xml:"method,omitempty"`
  // {'en':'Whether to define body parameters, required when method = POST/PUT/PATCH.
  // TRUE:Yes
  // FALSE:No', 'zh_CN':'是否定义body参数，method = POST/PUT/PATCH时必填。
  // TRUE：是
  // FALSE：否'}
  BodyFlag *string `json:"bodyFlag,omitempty" xml:"bodyFlag,omitempty"`
  // {'en':'Body parameter type, required when bodyFlag = TRUE.
  // JSON:JSON
  // FROM:FROM', 'zh_CN':'body参数类型，bodyFlag = TRUE 时必填。
  // JSON：JSON
  // FROM：FROM表单'}
  BodyType *string `json:"bodyType,omitempty" xml:"bodyType,omitempty"`
  // {'en':'Normal parameter list.', 'zh_CN':'普通参数数组。'}
  NormalParamList []*UpdateAPIDefinitionApiDefineParamLimitNormalEditDTO `json:"normalParamList,omitempty" xml:"normalParamList,omitempty" type:"Repeated"`
  // {'en':'FROM form parameter array, Optional when bodyType = FROM.', 'zh_CN':'FROM 表单参数数组，bodyType = FROM 时选填。'}
  FormParamList []*UpdateAPIDefinitionApiDefineParamLimitFormEditDTO `json:"formParamList,omitempty" xml:"formParamList,omitempty" type:"Repeated"`
  // {'en':'JSON parameter array, Optional when bodyType = JSON.', 'zh_CN':'JSON参数数组，bodyType = JSON 时选填。'}
  JsonParam *UpdateAPIDefinitionApiDefineParamLimitJsonEditDTO `json:"jsonParam,omitempty" xml:"jsonParam,omitempty"`
}

func (s UpdateAPIDefinitionApiDefineParamLimitMethodEditDTO) String() string {
  return tea.Prettify(s)
}

func (s UpdateAPIDefinitionApiDefineParamLimitMethodEditDTO) GoString() string {
  return s.String()
}

func (s *UpdateAPIDefinitionApiDefineParamLimitMethodEditDTO) SetMethod(v string) *UpdateAPIDefinitionApiDefineParamLimitMethodEditDTO {
  s.Method = &v
  return s
}

func (s *UpdateAPIDefinitionApiDefineParamLimitMethodEditDTO) SetBodyFlag(v string) *UpdateAPIDefinitionApiDefineParamLimitMethodEditDTO {
  s.BodyFlag = &v
  return s
}

func (s *UpdateAPIDefinitionApiDefineParamLimitMethodEditDTO) SetBodyType(v string) *UpdateAPIDefinitionApiDefineParamLimitMethodEditDTO {
  s.BodyType = &v
  return s
}

func (s *UpdateAPIDefinitionApiDefineParamLimitMethodEditDTO) SetNormalParamList(v []*UpdateAPIDefinitionApiDefineParamLimitNormalEditDTO) *UpdateAPIDefinitionApiDefineParamLimitMethodEditDTO {
  s.NormalParamList = v
  return s
}

func (s *UpdateAPIDefinitionApiDefineParamLimitMethodEditDTO) SetFormParamList(v []*UpdateAPIDefinitionApiDefineParamLimitFormEditDTO) *UpdateAPIDefinitionApiDefineParamLimitMethodEditDTO {
  s.FormParamList = v
  return s
}

func (s *UpdateAPIDefinitionApiDefineParamLimitMethodEditDTO) SetJsonParam(v *UpdateAPIDefinitionApiDefineParamLimitJsonEditDTO) *UpdateAPIDefinitionApiDefineParamLimitMethodEditDTO {
  s.JsonParam = v
  return s
}

type UpdateAPIDefinitionApiDefineParamLimitNormalEditDTO struct {
  // {'en':'Parameter name, when paramPosition = PATH_PARAMS, the path need to be matched, for example: /basePath/{name}/.', 'zh_CN':'参数名称，paramPosition = PATH_PARAMS 时需匹配路径变量，例如：/basePath/{name}/。'}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // {'en':'Parameter position.
  // HTTP_HEADER:Http Header
  // QUERY_STRING:Query String
  // COOKIE:Cookie
  // PATH_PARAMS:Path, endpoint.matchPathVar needs to equal TRUE', 'zh_CN':'参数位置。
  // HTTP_HEADER：Http Header
  // QUERY_STRING：Query String
  // COOKIE：Cookie
  // PATH_PARAMS：路径变量，endpoint.matchPathVar需等于TRUE'}
  ParamPosition *string `json:"paramPosition,omitempty" xml:"paramPosition,omitempty"`
  // {'en':'Parameter type.
  // Integer:Integer
  // Number:Number
  // String:String
  // Boolean:Boolean
  // Enumeration:Enumeration', 'zh_CN':'参数类型。
  // Integer：整数
  // Number：数字
  // String：字符串
  // Boolean：布尔
  // Enumeration：枚举'}
  Type *string `json:"type,omitempty" xml:"type,omitempty"`
  // {'en':'type = Integer/Number:minimum value, type = String:minimum length, type = Boolean/Enumeration:leave blank.', 'zh_CN':'type = Integer/Number：最小值，type = String：最小长度，type = Boolean/Enumeration：置空。'}
  MinVal *string `json:"minVal,omitempty" xml:"minVal,omitempty"`
  // {'en':'type = Integer/Number:maximum value, type = String:maximum length, type = Boolean/Enumeration:leave blank.', 'zh_CN':'type = Integer/Number：最大值，type = String：最大长度，type = Boolean/Enumeration：置空。'}
  MaxVal *string `json:"maxVal,omitempty" xml:"maxVal,omitempty"`
  // {'en':'Required.
  // TRUE:Yes
  // FALSE:No', 'zh_CN':'必带参数。
  // TRUE：是
  // FALSE：否'}
  Required *string `json:"required,omitempty" xml:"required,omitempty"`
  // {'en':'Content (maximum 2000 characters), required when type = Enumeration, Multiple separated by ; sign.', 'zh_CN':'内容（最多2000个字符），type = Enumeration 时必填，多个以 ; 号分隔。'}
  Content *string `json:"content,omitempty" xml:"content,omitempty"`
  // {'en':'Description.', 'zh_CN':'备注。'}
  Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
}

func (s UpdateAPIDefinitionApiDefineParamLimitNormalEditDTO) String() string {
  return tea.Prettify(s)
}

func (s UpdateAPIDefinitionApiDefineParamLimitNormalEditDTO) GoString() string {
  return s.String()
}

func (s *UpdateAPIDefinitionApiDefineParamLimitNormalEditDTO) SetName(v string) *UpdateAPIDefinitionApiDefineParamLimitNormalEditDTO {
  s.Name = &v
  return s
}

func (s *UpdateAPIDefinitionApiDefineParamLimitNormalEditDTO) SetParamPosition(v string) *UpdateAPIDefinitionApiDefineParamLimitNormalEditDTO {
  s.ParamPosition = &v
  return s
}

func (s *UpdateAPIDefinitionApiDefineParamLimitNormalEditDTO) SetType(v string) *UpdateAPIDefinitionApiDefineParamLimitNormalEditDTO {
  s.Type = &v
  return s
}

func (s *UpdateAPIDefinitionApiDefineParamLimitNormalEditDTO) SetMinVal(v string) *UpdateAPIDefinitionApiDefineParamLimitNormalEditDTO {
  s.MinVal = &v
  return s
}

func (s *UpdateAPIDefinitionApiDefineParamLimitNormalEditDTO) SetMaxVal(v string) *UpdateAPIDefinitionApiDefineParamLimitNormalEditDTO {
  s.MaxVal = &v
  return s
}

func (s *UpdateAPIDefinitionApiDefineParamLimitNormalEditDTO) SetRequired(v string) *UpdateAPIDefinitionApiDefineParamLimitNormalEditDTO {
  s.Required = &v
  return s
}

func (s *UpdateAPIDefinitionApiDefineParamLimitNormalEditDTO) SetContent(v string) *UpdateAPIDefinitionApiDefineParamLimitNormalEditDTO {
  s.Content = &v
  return s
}

func (s *UpdateAPIDefinitionApiDefineParamLimitNormalEditDTO) SetRemark(v string) *UpdateAPIDefinitionApiDefineParamLimitNormalEditDTO {
  s.Remark = &v
  return s
}

type UpdateAPIDefinitionApiDefineParamLimitFormEditDTO struct {
  // {'en':'Parameter name.', 'zh_CN':'参数名称。'}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // {'en':'Parameter type.
  // Integer:Integer
  // Number:Number
  // String:String
  // Boolean:Boolean
  // Enumeration:Enumeration', 'zh_CN':'参数类型。
  // Integer：整数
  // Number：数字
  // String：字符串
  // Boolean：布尔
  // Enumeration：枚举'}
  Type *string `json:"type,omitempty" xml:"type,omitempty"`
  // {'en':'type = Integer/Number:minimum value, type = String:minimum length, type = Boolean/Enumeration:leave blank.', 'zh_CN':'type = Integer/Number：最小值，type = String：最小长度，type = Boolean/Enumeration：置空。'}
  MinVal *string `json:"minVal,omitempty" xml:"minVal,omitempty"`
  // {'en':'type = Integer/Number:maximum value, type = String:maximum length, type = Boolean/Enumeration:leave blank.', 'zh_CN':'type = Integer/Number：最大值，type = String：最大长度，type = Boolean/Enumeration：置空。'}
  MaxVal *string `json:"maxVal,omitempty" xml:"maxVal,omitempty"`
  // {'en':'Required.
  // TRUE:Yes
  // FALSE:No', 'zh_CN':'必带参数。
  // TRUE：是
  // FALSE：否'}
  Required *string `json:"required,omitempty" xml:"required,omitempty"`
  // {'en':'Content (maximum 2000 characters), required when type = Enumeration, multiple separated by ; sign.', 'zh_CN':'内容（最多2000个字符），type = Enumeration 时必填，多个以 ; 号分隔。'}
  Content *string `json:"content,omitempty" xml:"content,omitempty"`
  // {'en':'Description.', 'zh_CN':'备注。'}
  Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
}

func (s UpdateAPIDefinitionApiDefineParamLimitFormEditDTO) String() string {
  return tea.Prettify(s)
}

func (s UpdateAPIDefinitionApiDefineParamLimitFormEditDTO) GoString() string {
  return s.String()
}

func (s *UpdateAPIDefinitionApiDefineParamLimitFormEditDTO) SetName(v string) *UpdateAPIDefinitionApiDefineParamLimitFormEditDTO {
  s.Name = &v
  return s
}

func (s *UpdateAPIDefinitionApiDefineParamLimitFormEditDTO) SetType(v string) *UpdateAPIDefinitionApiDefineParamLimitFormEditDTO {
  s.Type = &v
  return s
}

func (s *UpdateAPIDefinitionApiDefineParamLimitFormEditDTO) SetMinVal(v string) *UpdateAPIDefinitionApiDefineParamLimitFormEditDTO {
  s.MinVal = &v
  return s
}

func (s *UpdateAPIDefinitionApiDefineParamLimitFormEditDTO) SetMaxVal(v string) *UpdateAPIDefinitionApiDefineParamLimitFormEditDTO {
  s.MaxVal = &v
  return s
}

func (s *UpdateAPIDefinitionApiDefineParamLimitFormEditDTO) SetRequired(v string) *UpdateAPIDefinitionApiDefineParamLimitFormEditDTO {
  s.Required = &v
  return s
}

func (s *UpdateAPIDefinitionApiDefineParamLimitFormEditDTO) SetContent(v string) *UpdateAPIDefinitionApiDefineParamLimitFormEditDTO {
  s.Content = &v
  return s
}

func (s *UpdateAPIDefinitionApiDefineParamLimitFormEditDTO) SetRemark(v string) *UpdateAPIDefinitionApiDefineParamLimitFormEditDTO {
  s.Remark = &v
  return s
}

type UpdateAPIDefinitionApiDefineParamLimitJsonEditDTO struct {
  // {'en':'Parameter Name.', 'zh_CN':'参数名称。'}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // {'en':'Parameter Type.
  // Array:Array
  // Json:JSON Object', 'zh_CN':'参数类型。
  // Array：数组
  // Json：JSON对象'}
  Type *string `json:"type,omitempty" xml:"type,omitempty"`
  // {'en':'type = Integer/Number:minimum value, type = String:minimum length, type = Boolean/Enumeration:leave blank.', 'zh_CN':'type = Integer/Number：最小值，type = String：最小长度，type = Boolean/Enumeration：置空。'}
  MinVal *string `json:"minVal,omitempty" xml:"minVal,omitempty"`
  // {'en':'type = Integer/Number:maximum value, type = String:maximum length, type = Boolean/Enumeration:leave blank.', 'zh_CN':'type = Integer/Number：最大值，type = String：最大长度，type = Boolean/Enumeration：置空。'}
  MaxVal *string `json:"maxVal,omitempty" xml:"maxVal,omitempty"`
  // {'en':'Required.
  // TRUE:Yes
  // FALSE:No', 'zh_CN':'必带参数。
  // TRUE：是
  // FALSE：否'}
  Required *string `json:"required,omitempty" xml:"required,omitempty"`
  // {'en':'Content (maximum 2000 characters), required when type = Enumeration, multiple separated by ; sign.', 'zh_CN':'内容（最多2000个字符），type = Enumeration 时必填，多个以 ; 号分隔。'}
  Content *string `json:"content,omitempty" xml:"content,omitempty"`
  // {'en':'Array of child nodes, required when type= Array.', 'zh_CN':'子节点数组，type = Array 时必填。'}
  Children []*UpdateAPIDefinitionApiDefineParamLimitJsonChildrenDTO `json:"children,omitempty" xml:"children,omitempty" type:"Repeated"`
}

func (s UpdateAPIDefinitionApiDefineParamLimitJsonEditDTO) String() string {
  return tea.Prettify(s)
}

func (s UpdateAPIDefinitionApiDefineParamLimitJsonEditDTO) GoString() string {
  return s.String()
}

func (s *UpdateAPIDefinitionApiDefineParamLimitJsonEditDTO) SetName(v string) *UpdateAPIDefinitionApiDefineParamLimitJsonEditDTO {
  s.Name = &v
  return s
}

func (s *UpdateAPIDefinitionApiDefineParamLimitJsonEditDTO) SetType(v string) *UpdateAPIDefinitionApiDefineParamLimitJsonEditDTO {
  s.Type = &v
  return s
}

func (s *UpdateAPIDefinitionApiDefineParamLimitJsonEditDTO) SetMinVal(v string) *UpdateAPIDefinitionApiDefineParamLimitJsonEditDTO {
  s.MinVal = &v
  return s
}

func (s *UpdateAPIDefinitionApiDefineParamLimitJsonEditDTO) SetMaxVal(v string) *UpdateAPIDefinitionApiDefineParamLimitJsonEditDTO {
  s.MaxVal = &v
  return s
}

func (s *UpdateAPIDefinitionApiDefineParamLimitJsonEditDTO) SetRequired(v string) *UpdateAPIDefinitionApiDefineParamLimitJsonEditDTO {
  s.Required = &v
  return s
}

func (s *UpdateAPIDefinitionApiDefineParamLimitJsonEditDTO) SetContent(v string) *UpdateAPIDefinitionApiDefineParamLimitJsonEditDTO {
  s.Content = &v
  return s
}

func (s *UpdateAPIDefinitionApiDefineParamLimitJsonEditDTO) SetChildren(v []*UpdateAPIDefinitionApiDefineParamLimitJsonChildrenDTO) *UpdateAPIDefinitionApiDefineParamLimitJsonEditDTO {
  s.Children = v
  return s
}

type UpdateAPIDefinitionApiDefineParamLimitJsonChildrenDTO struct {
  // {'en':'Parameter Name.', 'zh_CN':'参数名称。'}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // {'en':'Parameter Type.
  // Integer:Integer
  // Number:Number
  // String:String
  // Boolean:Boolean
  // Enumeration:Enumeration
  // Array:Array
  // Json:JSON Object', 'zh_CN':'参数类型。
  // Integer：整数
  // Number：数字
  // String：字符串
  // Boolean：布尔
  // Enumeration：枚举
  // Array：数组
  // Json：JSON对象'}
  Type *string `json:"type,omitempty" xml:"type,omitempty"`
  // {'en':'type = Integer/Number:minimum value, type = String:minimum length, type = Boolean/Enumeration:leave blank.', 'zh_CN':'type = Integer/Number：最小值，type = String：最小长度，type = Boolean/Enumeration：置空。'}
  MinVal *string `json:"minVal,omitempty" xml:"minVal,omitempty"`
  // {'en':'type = Integer/Number:maximum value, type = String:maximum length, type = Boolean/Enumeration:leave blank.', 'zh_CN':'type = Integer/Number：最大值，type = String：最大长度，type = Boolean/Enumeration：置空。'}
  MaxVal *string `json:"maxVal,omitempty" xml:"maxVal,omitempty"`
  // {'en':'Required.
  // TRUE:Yes
  // FALSE:No', 'zh_CN':'必带参数。
  // TRUE：是
  // FALSE：否'}
  Required *string `json:"required,omitempty" xml:"required,omitempty"`
  // {'en':'Content (maximum 2000 characters), required when type = Enumeration, multiple separated by; sign.', 'zh_CN':'内容（最多2000个字符），type = Enumeration 时必填，多个以 ; 号分隔。'}
  Content *string `json:"content,omitempty" xml:"content,omitempty"`
  // {'en':'Child node array (JSON string), the structure is consistent with the parent node, required when type= Array.', 'zh_CN':'子节点数组（JSON字符串），结构与父结点一致，type = Array 时必填。'}
  Children []*string `json:"children,omitempty" xml:"children,omitempty" type:"Repeated"`
}

func (s UpdateAPIDefinitionApiDefineParamLimitJsonChildrenDTO) String() string {
  return tea.Prettify(s)
}

func (s UpdateAPIDefinitionApiDefineParamLimitJsonChildrenDTO) GoString() string {
  return s.String()
}

func (s *UpdateAPIDefinitionApiDefineParamLimitJsonChildrenDTO) SetName(v string) *UpdateAPIDefinitionApiDefineParamLimitJsonChildrenDTO {
  s.Name = &v
  return s
}

func (s *UpdateAPIDefinitionApiDefineParamLimitJsonChildrenDTO) SetType(v string) *UpdateAPIDefinitionApiDefineParamLimitJsonChildrenDTO {
  s.Type = &v
  return s
}

func (s *UpdateAPIDefinitionApiDefineParamLimitJsonChildrenDTO) SetMinVal(v string) *UpdateAPIDefinitionApiDefineParamLimitJsonChildrenDTO {
  s.MinVal = &v
  return s
}

func (s *UpdateAPIDefinitionApiDefineParamLimitJsonChildrenDTO) SetMaxVal(v string) *UpdateAPIDefinitionApiDefineParamLimitJsonChildrenDTO {
  s.MaxVal = &v
  return s
}

func (s *UpdateAPIDefinitionApiDefineParamLimitJsonChildrenDTO) SetRequired(v string) *UpdateAPIDefinitionApiDefineParamLimitJsonChildrenDTO {
  s.Required = &v
  return s
}

func (s *UpdateAPIDefinitionApiDefineParamLimitJsonChildrenDTO) SetContent(v string) *UpdateAPIDefinitionApiDefineParamLimitJsonChildrenDTO {
  s.Content = &v
  return s
}

func (s *UpdateAPIDefinitionApiDefineParamLimitJsonChildrenDTO) SetChildren(v []*string) *UpdateAPIDefinitionApiDefineParamLimitJsonChildrenDTO {
  s.Children = v
  return s
}

type UpdateAPIDefinitionResponse struct {
  // {'en':'Please refer to the error code for exceptions.', 'zh_CN':'请参照错误码。','dictionary':'belong=WAAP-MS-Ext|dict=waap_retCodeEnum'}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {'en':'Description.', 'zh_CN':'描述信息。'}
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

type UpdateAPIDefinitionRequestHeader struct {
  // {"zh_CN":"安全服务类型。有使用多个不同的安全服务时，需要填写具体的服务类型。","en":"Security service type. Please enter a specific service type, if you purchase multiple security services.","dictionary":"belong=WAAP-MS-Ext|dict=waap_serviceType"}
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

type UpdateAPIDefinitionResponseHeader struct {
}

func (s UpdateAPIDefinitionResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s UpdateAPIDefinitionResponseHeader) GoString() string {
  return s.String()
}




type ListAPIAssetDiscoveryRequest struct {
  // {"en":"Hostname list.", "zh_CN":"域名列表。"}
  DomainList []*string `json:"domainList,omitempty" xml:"domainList,omitempty" require:"true" type:"Repeated"`
  // {"en":"API base path.", "zh_CN":"端点路径。"}
  PathList []*string `json:"pathList,omitempty" xml:"pathList,omitempty" type:"Repeated"`
  // {"en":"Definition state.
  // DEFINED: Defined
  // UNDEFINED:Undefined", "zh_CN":"定义状态。
  // DEFINED: 已定义
  // UNDEFINED:未定义"}
  DefineStatus *string `json:"defineStatus,omitempty" xml:"defineStatus,omitempty"`
  // {"en":"Sort, format: field1,sort1;field2,sort2.
  // Optional field:
  // lastDiscoveryTime: Last discovery time
  // firstDiscoveryTime: First discovery time
  // reqCount24h: 24h Requests
  // Optional sort:
  // ascending:Ascend
  // descending:Descend.", "zh_CN":"排序，格式：字段1,排序1;字段2,排序2。
  // 可选字段：
  // lastDiscoveryTime：最新发现时间
  // firstDiscoveryTime：首次发现时间
  // reqCount24h：24h调用量
  // 可选排序：
  // ascending：升序
  // descending：降序"}
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

type ListAPIAssetDiscoveryResponse struct {
  // {'en':'Please refer to the error code for exceptions.', 'zh_CN':'请参照错误码。','dictionary':'belong=WAAP-MS-Ext|dict=waap_retCodeEnum'}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Description.", "zh_CN":"描述信息。"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"Data.", "zh_CN":"出参数据。"}
  Data []*ListAPIAssetDiscoveryApiDiscoveryLogVO `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
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

func (s *ListAPIAssetDiscoveryResponse) SetData(v []*ListAPIAssetDiscoveryApiDiscoveryLogVO) *ListAPIAssetDiscoveryResponse {
  s.Data = v
  return s
}

type ListAPIAssetDiscoveryApiDiscoveryLogVO struct {
  // {"en":"API discovery ID.", "zh_CN":"API发现ID。"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {"en":"Hostname.", "zh_CN":"域名。"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {"en":"API base path.", "zh_CN":"端点路径。"}
  Path *string `json:"path,omitempty" xml:"path,omitempty" require:"true"`
  // {"en":"Last Discovery Time, format: yyyy-MM-dd HH:mm:ss.", "zh_CN":"最新发现时间，格式：yyyy-MM-dd HH:mm:ss。"}
  LastDiscoveryTime *string `json:"lastDiscoveryTime,omitempty" xml:"lastDiscoveryTime,omitempty" require:"true"`
  // {"en":"First Discovery Time, format: yyyy-MM-dd HH:mm:ss.", "zh_CN":"首次发现时间，格式：yyyy-MM-dd HH:mm:ss。"}
  FirstDiscoveryTime *string `json:"firstDiscoveryTime,omitempty" xml:"firstDiscoveryTime,omitempty" require:"true"`
  // {"en":"24h Requests.", "zh_CN":"24h调用量。"}
  ReqCount24h *int64 `json:"reqCount24h,omitempty" xml:"reqCount24h,omitempty" require:"true"`
  // {"en":"Status.
  // DEFINED: Defined
  // UNDEFINED:Pending", "zh_CN":"定义状态。
  // DEFINED: 已定义
  // UNDEFINED: 待确认"}
  DefineStatus *string `json:"defineStatus,omitempty" xml:"defineStatus,omitempty" require:"true"`
}

func (s ListAPIAssetDiscoveryApiDiscoveryLogVO) String() string {
  return tea.Prettify(s)
}

func (s ListAPIAssetDiscoveryApiDiscoveryLogVO) GoString() string {
  return s.String()
}

func (s *ListAPIAssetDiscoveryApiDiscoveryLogVO) SetId(v string) *ListAPIAssetDiscoveryApiDiscoveryLogVO {
  s.Id = &v
  return s
}

func (s *ListAPIAssetDiscoveryApiDiscoveryLogVO) SetDomain(v string) *ListAPIAssetDiscoveryApiDiscoveryLogVO {
  s.Domain = &v
  return s
}

func (s *ListAPIAssetDiscoveryApiDiscoveryLogVO) SetPath(v string) *ListAPIAssetDiscoveryApiDiscoveryLogVO {
  s.Path = &v
  return s
}

func (s *ListAPIAssetDiscoveryApiDiscoveryLogVO) SetLastDiscoveryTime(v string) *ListAPIAssetDiscoveryApiDiscoveryLogVO {
  s.LastDiscoveryTime = &v
  return s
}

func (s *ListAPIAssetDiscoveryApiDiscoveryLogVO) SetFirstDiscoveryTime(v string) *ListAPIAssetDiscoveryApiDiscoveryLogVO {
  s.FirstDiscoveryTime = &v
  return s
}

func (s *ListAPIAssetDiscoveryApiDiscoveryLogVO) SetReqCount24h(v int64) *ListAPIAssetDiscoveryApiDiscoveryLogVO {
  s.ReqCount24h = &v
  return s
}

func (s *ListAPIAssetDiscoveryApiDiscoveryLogVO) SetDefineStatus(v string) *ListAPIAssetDiscoveryApiDiscoveryLogVO {
  s.DefineStatus = &v
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

type ListAPIAssetDiscoveryRequestHeader struct {
  // {"zh_CN":"安全服务类型。有使用多个不同的安全服务时，需要填写具体的服务类型。","en":"Security service type. Please enter a specific service type, if you purchase multiple security services.","dictionary":"belong=WAAP-MS-Ext|dict=waap_serviceType"}
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

type ListAPIAssetDiscoveryResponseHeader struct {
}

func (s ListAPIAssetDiscoveryResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ListAPIAssetDiscoveryResponseHeader) GoString() string {
  return s.String()
}




type CreateAPIDefinitionRequest struct {
  // {'en':'Basic information.', 'zh_CN':'基础信息。'}
  Basic *CreateAPIDefinitionApiDefineBasicAddDTO `json:"basic,omitempty" xml:"basic,omitempty" require:"true"`
  // {'en':'Endpoint information.', 'zh_CN':'端点信息。'}
  Endpoint *CreateAPIDefinitionApiDefineEndpointAddDTO `json:"endpoint,omitempty" xml:"endpoint,omitempty" require:"true"`
  // {'en':'Authentication configuration.', 'zh_CN':'鉴权配置。'}
  Auth *CreateAPIDefinitionApiDefineAuthAddDTO `json:"auth,omitempty" xml:"auth,omitempty" require:"true"`
  // {'en':'Body restrictions.', 'zh_CN':'body限制。'}
  BodyLimit *CreateAPIDefinitionApiDefineBodyLimitAddDTO `json:"bodyLimit,omitempty" xml:"bodyLimit,omitempty"`
  // {'en':'Parameter limit.', 'zh_CN':'参数限制。'}
  ParamLimit *CreateAPIDefinitionApiDefineParamLimitAddDTO `json:"paramLimit,omitempty" xml:"paramLimit,omitempty"`
}

func (s CreateAPIDefinitionRequest) String() string {
  return tea.Prettify(s)
}

func (s CreateAPIDefinitionRequest) GoString() string {
  return s.String()
}

func (s *CreateAPIDefinitionRequest) SetBasic(v *CreateAPIDefinitionApiDefineBasicAddDTO) *CreateAPIDefinitionRequest {
  s.Basic = v
  return s
}

func (s *CreateAPIDefinitionRequest) SetEndpoint(v *CreateAPIDefinitionApiDefineEndpointAddDTO) *CreateAPIDefinitionRequest {
  s.Endpoint = v
  return s
}

func (s *CreateAPIDefinitionRequest) SetAuth(v *CreateAPIDefinitionApiDefineAuthAddDTO) *CreateAPIDefinitionRequest {
  s.Auth = v
  return s
}

func (s *CreateAPIDefinitionRequest) SetBodyLimit(v *CreateAPIDefinitionApiDefineBodyLimitAddDTO) *CreateAPIDefinitionRequest {
  s.BodyLimit = v
  return s
}

func (s *CreateAPIDefinitionRequest) SetParamLimit(v *CreateAPIDefinitionApiDefineParamLimitAddDTO) *CreateAPIDefinitionRequest {
  s.ParamLimit = v
  return s
}

type CreateAPIDefinitionApiDefineBasicAddDTO struct {
  // {'en':'API name, maximum 200 characters.', 'zh_CN':'API名称，最多200个字符。'}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {'en':'API groups, maximum 64 characters.', 'zh_CN':'API分组，最多64个字符。'}
  GroupName *string `json:"groupName,omitempty" xml:"groupName,omitempty" require:"true"`
  // {'en':'Hostname.', 'zh_CN':'所属域名。'}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {'en':'Description, maximum 200 characters.', 'zh_CN':'备注，最多200个字符。'}
  Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
}

func (s CreateAPIDefinitionApiDefineBasicAddDTO) String() string {
  return tea.Prettify(s)
}

func (s CreateAPIDefinitionApiDefineBasicAddDTO) GoString() string {
  return s.String()
}

func (s *CreateAPIDefinitionApiDefineBasicAddDTO) SetName(v string) *CreateAPIDefinitionApiDefineBasicAddDTO {
  s.Name = &v
  return s
}

func (s *CreateAPIDefinitionApiDefineBasicAddDTO) SetGroupName(v string) *CreateAPIDefinitionApiDefineBasicAddDTO {
  s.GroupName = &v
  return s
}

func (s *CreateAPIDefinitionApiDefineBasicAddDTO) SetDomain(v string) *CreateAPIDefinitionApiDefineBasicAddDTO {
  s.Domain = &v
  return s
}

func (s *CreateAPIDefinitionApiDefineBasicAddDTO) SetRemark(v string) *CreateAPIDefinitionApiDefineBasicAddDTO {
  s.Remark = &v
  return s
}

type CreateAPIDefinitionApiDefineEndpointAddDTO struct {
  // {'en':'Request methods.Multiple separated by ; sign.
  // GET:GET, configurable parameter limits
  // POST:POST, configurable parameter limits
  // DELETE:DELETE, configurable parameter limits
  // UPDATE:UPDATE
  // PUT:PUT, configurable parameter limits
  // HEAD:HEAD, configurable parameter limits
  // CONNECT:CONNECT
  // OPTIONS:OPTIONS, configurable parameter limits
  // COPY:COPY
  // LOCK:LOCK
  // UNLOCK:UNLOCK
  // TRACE:TRACE
  // PATCH:PATCH, configurable parameter limits
  // PROPFIND:PROPFIND
  // MKCOL:MKCOL
  // MOVE:MOVE', 'zh_CN':'请求方法。多个以 ; 号分隔。
  // GET：GET，可配置参数限制
  // POST：POST，可配置参数限制
  // DELETE：DELETE，可配置参数限制
  // UPDATE：UPDATE
  // PUT：PUT，可配置参数限制
  // HEAD：HEAD，可配置参数限制
  // CONNECT：CONNECT
  // OPTIONS：OPTIONS，可配置参数限制
  // COPY：COPY
  // LOCK：LOCK
  // UNLOCK：UNLOCK
  // TRACE：TRACE
  // PATCH：PATCH，可配置参数限制
  // PROPFIND：PROPFIND
  // MKCOL：MKCOL
  // MOVE：MOVE'}
  RequestMethod *string `json:"requestMethod,omitempty" xml:"requestMethod,omitempty" require:"true"`
  // {'en':'API type.
  // NORMAL:Common API, the path does not contain query string parameters, such as http://www.test.com/api.
  // WHEN_CASE:Common API, the path contains query string parameters, such as http://www.test.com/api?action=1 and http://www.test.com/api?action=2 are two different APIs.', 'zh_CN':'API类型。
  // NORMAL：普通接口，路径中不包含query string参数的普通接口，如http://www.test.com/api。
  // WHEN_CASE：when-case接口，路径中包含query string参数，如http://www.test.com/api?action=1与http://www.test.com/api?action=2 是两个不同的接口。'}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
  // {'en':'Path matching type, EQUAL is passed when type = WHEM_CASE.
  // EQUAL: Complete match
  // REGEX: Regular matching', 'zh_CN':'路径匹配类型，type = WHEN_CASE 时传 EQUAL。
  // EQUAL：完整匹配
  // REGEX：正则匹配'}
  MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
  // {'en':'API base Path, only English characters, underscores, hyphens, or numbers are supported, maximum 200 characters.', 'zh_CN':'端点路径，只支持英文字符、下划线、短横线或数字，最多200个字符。'}
  Path *string `json:"path,omitempty" xml:"path,omitempty" require:"true"`
  // {'en':'Case sensitive, whether the endpoint path is case-sensitive. Example for case sensitive: /Order and /order are two different API paths.
  // TRUE:Yes
  // FALSE:No', 'zh_CN':'大小写是否敏感，端点路径是否区分大小写。若大小写敏感，示例：/Order 和 /order为两不同的API路径。
  // TRUE：是
  // FALSE：否'}
  CaseSensitive *string `json:"caseSensitive,omitempty" xml:"caseSensitive,omitempty" require:"true"`
  // {'en':'Match CreateAPIDefinitionParameters in the Path(Effective when type = NORMAL && matchType = EQUAL), with matching path parameters turned on, you can use curly braces "{}" to define path parameters. Example: /basePath/{pathParam1}/{pathParam2}/{pathParam3}.
  // TRUE:Yes
  // FALSE:No', 'zh_CN':'匹配路径参数（type = NORMAL && matchType = EQUAL 时生效），开启匹配路径参数后，您可以使用花括号({})来定义路径参数。例：/basePath/{pathParam1}/{pathParam2}/{pathParam3}。
  // TRUE：是
  // FALSE：否'}
  MatchPathVar *string `json:"matchPathVar,omitempty" xml:"matchPathVar,omitempty"`
}

func (s CreateAPIDefinitionApiDefineEndpointAddDTO) String() string {
  return tea.Prettify(s)
}

func (s CreateAPIDefinitionApiDefineEndpointAddDTO) GoString() string {
  return s.String()
}

func (s *CreateAPIDefinitionApiDefineEndpointAddDTO) SetRequestMethod(v string) *CreateAPIDefinitionApiDefineEndpointAddDTO {
  s.RequestMethod = &v
  return s
}

func (s *CreateAPIDefinitionApiDefineEndpointAddDTO) SetType(v string) *CreateAPIDefinitionApiDefineEndpointAddDTO {
  s.Type = &v
  return s
}

func (s *CreateAPIDefinitionApiDefineEndpointAddDTO) SetMatchType(v string) *CreateAPIDefinitionApiDefineEndpointAddDTO {
  s.MatchType = &v
  return s
}

func (s *CreateAPIDefinitionApiDefineEndpointAddDTO) SetPath(v string) *CreateAPIDefinitionApiDefineEndpointAddDTO {
  s.Path = &v
  return s
}

func (s *CreateAPIDefinitionApiDefineEndpointAddDTO) SetCaseSensitive(v string) *CreateAPIDefinitionApiDefineEndpointAddDTO {
  s.CaseSensitive = &v
  return s
}

func (s *CreateAPIDefinitionApiDefineEndpointAddDTO) SetMatchPathVar(v string) *CreateAPIDefinitionApiDefineEndpointAddDTO {
  s.MatchPathVar = &v
  return s
}

type CreateAPIDefinitionApiDefineAuthAddDTO struct {
  // {'en':'Authentication method.
  // NO_NEED:No authentication required
  // SIGN:Key authentication', 'zh_CN':'鉴权方法。
  // NO_NEED：免鉴权
  // SIGN：秘钥对鉴权'}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
  // {'en':'Authentication key, type = SIGN is required, the format is a 16-digit string containing uppercase and lowercase letters and numbers, example: gjZkg2E1uNkXBDxj.', 'zh_CN':'鉴权秘钥，type = SIGN 时必填，格式为16位含大小写字母与数字字符串，示例：gjZkg2E1uNkXBDxj。'}
  AuthKey *string `json:"authKey,omitempty" xml:"authKey,omitempty"`
  // {'en':'Authentication parameter location, type = SIGN is required.
  // HTTP_HEADER:Http Header
  // QUERY_STRING:Query String
  // COOKIE:Cookie', 'zh_CN':'鉴权参数位置，type = SIGN 时必填。
  // HTTP_HEADER：Http Header
  // QUERY_STRING：Query String
  // COOKIE：Cookie'}
  ParamPosition *string `json:"paramPosition,omitempty" xml:"paramPosition,omitempty"`
  // {'en':'Authentication parameter name, type = SIGN is required.', 'zh_CN':'鉴权参数名称，type = SIGN 时必填。'}
  ParamName *string `json:"paramName,omitempty" xml:"paramName,omitempty"`
  // {'en':'Authentication token validity period, in seconds, type = SIGN is required.', 'zh_CN':'鉴权有效期，单位秒，type = SIGN 时必填。'}
  ValidityTime *int64 `json:"validityTime,omitempty" xml:"validityTime,omitempty"`
}

func (s CreateAPIDefinitionApiDefineAuthAddDTO) String() string {
  return tea.Prettify(s)
}

func (s CreateAPIDefinitionApiDefineAuthAddDTO) GoString() string {
  return s.String()
}

func (s *CreateAPIDefinitionApiDefineAuthAddDTO) SetType(v string) *CreateAPIDefinitionApiDefineAuthAddDTO {
  s.Type = &v
  return s
}

func (s *CreateAPIDefinitionApiDefineAuthAddDTO) SetAuthKey(v string) *CreateAPIDefinitionApiDefineAuthAddDTO {
  s.AuthKey = &v
  return s
}

func (s *CreateAPIDefinitionApiDefineAuthAddDTO) SetParamPosition(v string) *CreateAPIDefinitionApiDefineAuthAddDTO {
  s.ParamPosition = &v
  return s
}

func (s *CreateAPIDefinitionApiDefineAuthAddDTO) SetParamName(v string) *CreateAPIDefinitionApiDefineAuthAddDTO {
  s.ParamName = &v
  return s
}

func (s *CreateAPIDefinitionApiDefineAuthAddDTO) SetValidityTime(v int64) *CreateAPIDefinitionApiDefineAuthAddDTO {
  s.ValidityTime = &v
  return s
}

type CreateAPIDefinitionApiDefineBodyLimitAddDTO struct {
  // {'en':'Request body restriction switch.default value:OFF.
  // ON:On
  // OFF:Off', 'zh_CN':'请求body限制开关。默认值：关。
  // ON：开启
  // OFF：关闭'}
  DefendSwitch *string `json:"defendSwitch,omitempty" xml:"defendSwitch,omitempty"`
  // {'en':'Content-Type, required when defendSwitch = ON.
  // FORM:FORM
  // JSON:JSON
  // ANY:ANY
  // EMPTY:EMPTY or NON-EXISTENT', 'zh_CN':'Content-Type，defendSwitch = ON 时必填。
  // FORM：FORM表单
  // JSON：JSON
  // ANY：任意
  // EMPTY：为空或不存在'}
  ContentType *string `json:"contentType,omitempty" xml:"contentType,omitempty"`
  // {'en':'Maximum body limit(bytes).', 'zh_CN':'Body最大限制(bytes)。'}
  BodyLimitMax *int64 `json:"bodyLimitMax,omitempty" xml:"bodyLimitMax,omitempty"`
  // {'en':'Maximum nesting depth, enter the maximum allowable JSON nesting depth in the request body.', 'zh_CN':'最大嵌套层数，输入允许的请求body中JSON嵌套层数最大值。'}
  NestMax *int64 `json:"nestMax,omitempty" xml:"nestMax,omitempty"`
  // {'en':'Maximum number of parameters for JSON, enter the maximum number of JSON parameters allowed in the request body.', 'zh_CN':'JSON最大参数个数，输入允许的请求body中JSON参数个数的最大值。'}
  ParamCountMax *int64 `json:"paramCountMax,omitempty" xml:"paramCountMax,omitempty"`
}

func (s CreateAPIDefinitionApiDefineBodyLimitAddDTO) String() string {
  return tea.Prettify(s)
}

func (s CreateAPIDefinitionApiDefineBodyLimitAddDTO) GoString() string {
  return s.String()
}

func (s *CreateAPIDefinitionApiDefineBodyLimitAddDTO) SetDefendSwitch(v string) *CreateAPIDefinitionApiDefineBodyLimitAddDTO {
  s.DefendSwitch = &v
  return s
}

func (s *CreateAPIDefinitionApiDefineBodyLimitAddDTO) SetContentType(v string) *CreateAPIDefinitionApiDefineBodyLimitAddDTO {
  s.ContentType = &v
  return s
}

func (s *CreateAPIDefinitionApiDefineBodyLimitAddDTO) SetBodyLimitMax(v int64) *CreateAPIDefinitionApiDefineBodyLimitAddDTO {
  s.BodyLimitMax = &v
  return s
}

func (s *CreateAPIDefinitionApiDefineBodyLimitAddDTO) SetNestMax(v int64) *CreateAPIDefinitionApiDefineBodyLimitAddDTO {
  s.NestMax = &v
  return s
}

func (s *CreateAPIDefinitionApiDefineBodyLimitAddDTO) SetParamCountMax(v int64) *CreateAPIDefinitionApiDefineBodyLimitAddDTO {
  s.ParamCountMax = &v
  return s
}

type CreateAPIDefinitionApiDefineParamLimitAddDTO struct {
  // {'en':'Basic information.', 'zh_CN':'基础信息。'}
  Basic *CreateAPIDefinitionApiDefineParamLimitBasicAddDTO `json:"basic,omitempty" xml:"basic,omitempty" require:"true"`
  // {'en':'Method attributes.', 'zh_CN':'方法属性。'}
  MethodList []*CreateAPIDefinitionApiDefineParamLimitMethodAddDTO `json:"methodList,omitempty" xml:"methodList,omitempty" type:"Repeated"`
}

func (s CreateAPIDefinitionApiDefineParamLimitAddDTO) String() string {
  return tea.Prettify(s)
}

func (s CreateAPIDefinitionApiDefineParamLimitAddDTO) GoString() string {
  return s.String()
}

func (s *CreateAPIDefinitionApiDefineParamLimitAddDTO) SetBasic(v *CreateAPIDefinitionApiDefineParamLimitBasicAddDTO) *CreateAPIDefinitionApiDefineParamLimitAddDTO {
  s.Basic = v
  return s
}

func (s *CreateAPIDefinitionApiDefineParamLimitAddDTO) SetMethodList(v []*CreateAPIDefinitionApiDefineParamLimitMethodAddDTO) *CreateAPIDefinitionApiDefineParamLimitAddDTO {
  s.MethodList = v
  return s
}

type CreateAPIDefinitionApiDefineParamLimitBasicAddDTO struct {
  // {'en':'Parameter limit.
  // ON:On
  // OFF:Off', 'zh_CN':'参数限制。
  // ON：开启
  // OFF：关闭'}
  DefendSwitch *string `json:"defendSwitch,omitempty" xml:"defendSwitch,omitempty"`
  // {'en':'Query String Parameter Detection Mode, required when defendSwitch = ON.
  // LOOSE:Loose mode - detect some parameters
  // STRICT:Strict mode - checks all parameters', 'zh_CN':'Query String参数检测模式，defendSwitch = ON 时必填。
  // LOOSE：宽松模式-检测部分参数
  // STRICT：严格模式-检测所有参数'}
  ParamCheckMode *string `json:"paramCheckMode,omitempty" xml:"paramCheckMode,omitempty"`
}

func (s CreateAPIDefinitionApiDefineParamLimitBasicAddDTO) String() string {
  return tea.Prettify(s)
}

func (s CreateAPIDefinitionApiDefineParamLimitBasicAddDTO) GoString() string {
  return s.String()
}

func (s *CreateAPIDefinitionApiDefineParamLimitBasicAddDTO) SetDefendSwitch(v string) *CreateAPIDefinitionApiDefineParamLimitBasicAddDTO {
  s.DefendSwitch = &v
  return s
}

func (s *CreateAPIDefinitionApiDefineParamLimitBasicAddDTO) SetParamCheckMode(v string) *CreateAPIDefinitionApiDefineParamLimitBasicAddDTO {
  s.ParamCheckMode = &v
  return s
}

type CreateAPIDefinitionApiDefineParamLimitMethodAddDTO struct {
  // {'en':'Request methods.
  // GET:GET
  // POST:POST
  // DELETE:DELETE
  // PUT:PUT
  // HEAD:HEAD
  // OPTIONS:OPTIONS
  // PATCH:PATCH', 'zh_CN':'请求方法。
  // GET：GET
  // POST：POST
  // DELETE：DELETE
  // PUT：PUT
  // HEAD：HEAD
  // OPTIONS：OPTIONS
  // PATCH：PATCH'}
  Method *string `json:"method,omitempty" xml:"method,omitempty"`
  // {'en':'Whether to define body parameters, required when method = POST/PUT/PATCH.
  // TRUE:Yes
  // FALSE:No', 'zh_CN':'是否定义body参数，method = POST/PUT/PATCH时必填。
  // TRUE：是
  // FALSE：否'}
  BodyFlag *string `json:"bodyFlag,omitempty" xml:"bodyFlag,omitempty"`
  // {'en':'Body parameter type, required when bodyFlag = TRUE.
  // JSON:JSON
  // FROM:FROM', 'zh_CN':'body参数类型，bodyFlag = TRUE 时必填。
  // JSON：JSON
  // FROM：FROM表单'}
  BodyType *string `json:"bodyType,omitempty" xml:"bodyType,omitempty"`
  // {'en':'Normal parameter list.', 'zh_CN':'普通参数数组。'}
  NormalParamList []*CreateAPIDefinitionApiDefineParamLimitNormalAddDTO `json:"normalParamList,omitempty" xml:"normalParamList,omitempty" type:"Repeated"`
  // {'en':'FROM form parameter array, Optional when bodyType = FROM.', 'zh_CN':'FROM 表单参数数组，bodyType = FROM 时选填。'}
  FormParamList []*CreateAPIDefinitionApiDefineParamLimitFormAddDTO `json:"formParamList,omitempty" xml:"formParamList,omitempty" type:"Repeated"`
  // {'en':'JSON parameter array, Optional when bodyType = JSON.', 'zh_CN':'JSON参数数组，bodyType = JSON 时选填。'}
  JsonParam *CreateAPIDefinitionApiDefineParamLimitJsonAddDTO `json:"jsonParam,omitempty" xml:"jsonParam,omitempty"`
}

func (s CreateAPIDefinitionApiDefineParamLimitMethodAddDTO) String() string {
  return tea.Prettify(s)
}

func (s CreateAPIDefinitionApiDefineParamLimitMethodAddDTO) GoString() string {
  return s.String()
}

func (s *CreateAPIDefinitionApiDefineParamLimitMethodAddDTO) SetMethod(v string) *CreateAPIDefinitionApiDefineParamLimitMethodAddDTO {
  s.Method = &v
  return s
}

func (s *CreateAPIDefinitionApiDefineParamLimitMethodAddDTO) SetBodyFlag(v string) *CreateAPIDefinitionApiDefineParamLimitMethodAddDTO {
  s.BodyFlag = &v
  return s
}

func (s *CreateAPIDefinitionApiDefineParamLimitMethodAddDTO) SetBodyType(v string) *CreateAPIDefinitionApiDefineParamLimitMethodAddDTO {
  s.BodyType = &v
  return s
}

func (s *CreateAPIDefinitionApiDefineParamLimitMethodAddDTO) SetNormalParamList(v []*CreateAPIDefinitionApiDefineParamLimitNormalAddDTO) *CreateAPIDefinitionApiDefineParamLimitMethodAddDTO {
  s.NormalParamList = v
  return s
}

func (s *CreateAPIDefinitionApiDefineParamLimitMethodAddDTO) SetFormParamList(v []*CreateAPIDefinitionApiDefineParamLimitFormAddDTO) *CreateAPIDefinitionApiDefineParamLimitMethodAddDTO {
  s.FormParamList = v
  return s
}

func (s *CreateAPIDefinitionApiDefineParamLimitMethodAddDTO) SetJsonParam(v *CreateAPIDefinitionApiDefineParamLimitJsonAddDTO) *CreateAPIDefinitionApiDefineParamLimitMethodAddDTO {
  s.JsonParam = v
  return s
}

type CreateAPIDefinitionApiDefineParamLimitNormalAddDTO struct {
  // {'en':'Parameter name, when paramPosition = PATH_PARAMS, the path need to be matched, for example: /basePath/{name}/.', 'zh_CN':'参数名称，paramPosition = PATH_PARAMS 时需匹配路径变量，例如：/basePath/{name}/。'}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // {'en':'Parameter position.
  // HTTP_HEADER:Http Header
  // QUERY_STRING:Query String
  // COOKIE:Cookie
  // PATH_PARAMS:Path, endpoint.matchPathVar needs to equal TRUE', 'zh_CN':'参数位置。
  // HTTP_HEADER：Http Header
  // QUERY_STRING：Query String
  // COOKIE：Cookie
  // PATH_PARAMS：路径变量，endpoint.matchPathVar需等于TRUE'}
  ParamPosition *string `json:"paramPosition,omitempty" xml:"paramPosition,omitempty"`
  // {'en':'Parameter type.
  // Integer:Integer
  // Number:Number
  // String:String
  // Boolean:Boolean
  // Enumeration:Enumeration', 'zh_CN':'参数类型。
  // Integer：整数
  // Number：数字
  // String：字符串
  // Boolean：布尔
  // Enumeration：枚举'}
  Type *string `json:"type,omitempty" xml:"type,omitempty"`
  // {'en':'type = Integer/Number:minimum value, type = String:minimum length, type = Boolean/Enumeration:leave blank.', 'zh_CN':'type = Integer/Number：最小值，type = String：最小长度，type = Boolean/Enumeration：置空。'}
  MinVal *string `json:"minVal,omitempty" xml:"minVal,omitempty"`
  // {'en':'type = Integer/Number:maximum value, type = String:maximum length, type = Boolean/Enumeration:leave blank.', 'zh_CN':'type = Integer/Number：最大值，type = String：最大长度，type = Boolean/Enumeration：置空。'}
  MaxVal *string `json:"maxVal,omitempty" xml:"maxVal,omitempty"`
  // {'en':'Required.
  // TRUE:Yes
  // FALSE:No', 'zh_CN':'必带参数。
  // TRUE：是
  // FALSE：否'}
  Required *string `json:"required,omitempty" xml:"required,omitempty"`
  // {'en':'Content (maximum 2000 characters), required when type = Enumeration,multiple separated by  ; sign.', 'zh_CN':'内容（最多2000个字符），type = Enumeration 时必填，多个以 ; 号分隔。'}
  Content *string `json:"content,omitempty" xml:"content,omitempty"`
  // {'en':'Description.', 'zh_CN':'备注。'}
  Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
}

func (s CreateAPIDefinitionApiDefineParamLimitNormalAddDTO) String() string {
  return tea.Prettify(s)
}

func (s CreateAPIDefinitionApiDefineParamLimitNormalAddDTO) GoString() string {
  return s.String()
}

func (s *CreateAPIDefinitionApiDefineParamLimitNormalAddDTO) SetName(v string) *CreateAPIDefinitionApiDefineParamLimitNormalAddDTO {
  s.Name = &v
  return s
}

func (s *CreateAPIDefinitionApiDefineParamLimitNormalAddDTO) SetParamPosition(v string) *CreateAPIDefinitionApiDefineParamLimitNormalAddDTO {
  s.ParamPosition = &v
  return s
}

func (s *CreateAPIDefinitionApiDefineParamLimitNormalAddDTO) SetType(v string) *CreateAPIDefinitionApiDefineParamLimitNormalAddDTO {
  s.Type = &v
  return s
}

func (s *CreateAPIDefinitionApiDefineParamLimitNormalAddDTO) SetMinVal(v string) *CreateAPIDefinitionApiDefineParamLimitNormalAddDTO {
  s.MinVal = &v
  return s
}

func (s *CreateAPIDefinitionApiDefineParamLimitNormalAddDTO) SetMaxVal(v string) *CreateAPIDefinitionApiDefineParamLimitNormalAddDTO {
  s.MaxVal = &v
  return s
}

func (s *CreateAPIDefinitionApiDefineParamLimitNormalAddDTO) SetRequired(v string) *CreateAPIDefinitionApiDefineParamLimitNormalAddDTO {
  s.Required = &v
  return s
}

func (s *CreateAPIDefinitionApiDefineParamLimitNormalAddDTO) SetContent(v string) *CreateAPIDefinitionApiDefineParamLimitNormalAddDTO {
  s.Content = &v
  return s
}

func (s *CreateAPIDefinitionApiDefineParamLimitNormalAddDTO) SetRemark(v string) *CreateAPIDefinitionApiDefineParamLimitNormalAddDTO {
  s.Remark = &v
  return s
}

type CreateAPIDefinitionApiDefineParamLimitFormAddDTO struct {
  // {'en':'Parameter name.', 'zh_CN':'参数名称。'}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // {'en':'Parameter type.
  // Integer:Integer
  // Number:Number
  // String:String
  // Boolean:Boolean
  // Enumeration:Enumeration', 'zh_CN':'参数类型。
  // Integer：整数
  // Number：数字
  // String：字符串
  // Boolean：布尔
  // Enumeration：枚举'}
  Type *string `json:"type,omitempty" xml:"type,omitempty"`
  // {'en':'type = Integer/Number:minimum value, type = String:minimum length, type = Boolean/Enumeration:leave blank.', 'zh_CN':'type = Integer/Number：最小值，type = String：最小长度，type = Boolean/Enumeration：置空。'}
  MinVal *string `json:"minVal,omitempty" xml:"minVal,omitempty"`
  // {'en':'type = Integer/Number:maximum value, type = String:maximum length, type = Boolean/Enumeration:leave blank.', 'zh_CN':'type = Integer/Number：最大值，type = String：最大长度，type = Boolean/Enumeration：置空。'}
  MaxVal *string `json:"maxVal,omitempty" xml:"maxVal,omitempty"`
  // {'en':'Required.
  // TRUE:Yes
  // FALSE:No', 'zh_CN':'必带参数。
  // TRUE：是
  // FALSE：否'}
  Required *string `json:"required,omitempty" xml:"required,omitempty"`
  // {'en':'Content (maximum 2000 characters), required when type = Enumeration, multiple separated by ; sign.', 'zh_CN':'内容（最多2000个字符），type = Enumeration 时必填，多个以 ; 号分隔。'}
  Content *string `json:"content,omitempty" xml:"content,omitempty"`
  // {'en':'Description.', 'zh_CN':'备注。'}
  Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
}

func (s CreateAPIDefinitionApiDefineParamLimitFormAddDTO) String() string {
  return tea.Prettify(s)
}

func (s CreateAPIDefinitionApiDefineParamLimitFormAddDTO) GoString() string {
  return s.String()
}

func (s *CreateAPIDefinitionApiDefineParamLimitFormAddDTO) SetName(v string) *CreateAPIDefinitionApiDefineParamLimitFormAddDTO {
  s.Name = &v
  return s
}

func (s *CreateAPIDefinitionApiDefineParamLimitFormAddDTO) SetType(v string) *CreateAPIDefinitionApiDefineParamLimitFormAddDTO {
  s.Type = &v
  return s
}

func (s *CreateAPIDefinitionApiDefineParamLimitFormAddDTO) SetMinVal(v string) *CreateAPIDefinitionApiDefineParamLimitFormAddDTO {
  s.MinVal = &v
  return s
}

func (s *CreateAPIDefinitionApiDefineParamLimitFormAddDTO) SetMaxVal(v string) *CreateAPIDefinitionApiDefineParamLimitFormAddDTO {
  s.MaxVal = &v
  return s
}

func (s *CreateAPIDefinitionApiDefineParamLimitFormAddDTO) SetRequired(v string) *CreateAPIDefinitionApiDefineParamLimitFormAddDTO {
  s.Required = &v
  return s
}

func (s *CreateAPIDefinitionApiDefineParamLimitFormAddDTO) SetContent(v string) *CreateAPIDefinitionApiDefineParamLimitFormAddDTO {
  s.Content = &v
  return s
}

func (s *CreateAPIDefinitionApiDefineParamLimitFormAddDTO) SetRemark(v string) *CreateAPIDefinitionApiDefineParamLimitFormAddDTO {
  s.Remark = &v
  return s
}

type CreateAPIDefinitionApiDefineParamLimitJsonAddDTO struct {
  // {'en':'Parameter Name.', 'zh_CN':'参数名称。'}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // {'en':'Parameter Type.
  // Array:Array
  // Json:JSON Object', 'zh_CN':'参数类型。
  // Array：数组
  // Json：JSON对象'}
  Type *string `json:"type,omitempty" xml:"type,omitempty"`
  // {'en':'type = Integer/Number:minimum value, type = String:minimum length, type = Boolean/Enumeration:leave blank.', 'zh_CN':'type = Integer/Number：最小值，type = String：最小长度，type = Boolean/Enumeration：置空。'}
  MinVal *string `json:"minVal,omitempty" xml:"minVal,omitempty"`
  // {'en':'type = Integer/Number:maximum value, type = String:maximum length, type = Boolean/Enumeration:leave blank.', 'zh_CN':'type = Integer/Number：最大值，type = String：最大长度，type = Boolean/Enumeration：置空。'}
  MaxVal *string `json:"maxVal,omitempty" xml:"maxVal,omitempty"`
  // {'en':'Required.
  // TRUE:Yes
  // FALSE:No', 'zh_CN':'必带参数。
  // TRUE：是
  // FALSE：否'}
  Required *string `json:"required,omitempty" xml:"required,omitempty"`
  // {'en':'Content (maximum 2000 characters), required when type = Enumeration, multiple separated by ; sign.', 'zh_CN':'内容（最多2000个字符），type = Enumeration 时必填，多个以 ; 号分隔。'}
  Content *string `json:"content,omitempty" xml:"content,omitempty"`
  // {'en':'Child node array, required when type= Array.', 'zh_CN':'子节点数组，type = Array 时必填。'}
  Children []*CreateAPIDefinitionApiDefineParamLimitJsonAddChildrenDTO `json:"children,omitempty" xml:"children,omitempty" type:"Repeated"`
}

func (s CreateAPIDefinitionApiDefineParamLimitJsonAddDTO) String() string {
  return tea.Prettify(s)
}

func (s CreateAPIDefinitionApiDefineParamLimitJsonAddDTO) GoString() string {
  return s.String()
}

func (s *CreateAPIDefinitionApiDefineParamLimitJsonAddDTO) SetName(v string) *CreateAPIDefinitionApiDefineParamLimitJsonAddDTO {
  s.Name = &v
  return s
}

func (s *CreateAPIDefinitionApiDefineParamLimitJsonAddDTO) SetType(v string) *CreateAPIDefinitionApiDefineParamLimitJsonAddDTO {
  s.Type = &v
  return s
}

func (s *CreateAPIDefinitionApiDefineParamLimitJsonAddDTO) SetMinVal(v string) *CreateAPIDefinitionApiDefineParamLimitJsonAddDTO {
  s.MinVal = &v
  return s
}

func (s *CreateAPIDefinitionApiDefineParamLimitJsonAddDTO) SetMaxVal(v string) *CreateAPIDefinitionApiDefineParamLimitJsonAddDTO {
  s.MaxVal = &v
  return s
}

func (s *CreateAPIDefinitionApiDefineParamLimitJsonAddDTO) SetRequired(v string) *CreateAPIDefinitionApiDefineParamLimitJsonAddDTO {
  s.Required = &v
  return s
}

func (s *CreateAPIDefinitionApiDefineParamLimitJsonAddDTO) SetContent(v string) *CreateAPIDefinitionApiDefineParamLimitJsonAddDTO {
  s.Content = &v
  return s
}

func (s *CreateAPIDefinitionApiDefineParamLimitJsonAddDTO) SetChildren(v []*CreateAPIDefinitionApiDefineParamLimitJsonAddChildrenDTO) *CreateAPIDefinitionApiDefineParamLimitJsonAddDTO {
  s.Children = v
  return s
}

type CreateAPIDefinitionApiDefineParamLimitJsonAddChildrenDTO struct {
  // {'en':'Parameter Name.', 'zh_CN':'参数名称。'}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // {'en':'Parameter Type.
  // Integer:Integer
  // Number:Number
  // String:String
  // Boolean:Boolean
  // Enumeration:Enumeration
  // Array:Array
  // Json:JSON Object', 'zh_CN':'参数类型。
  // Integer：整数
  // Number：数字
  // String：字符串
  // Boolean：布尔
  // Enumeration：枚举
  // Array：数组
  // Json：JSON对象'}
  Type *string `json:"type,omitempty" xml:"type,omitempty"`
  // {'en':'type = Integer/Number:minimum value, type = String:minimum length, type = Boolean/Enumeration:leave blank.', 'zh_CN':'type = Integer/Number：最小值，type = String：最小长度，type = Boolean/Enumeration：置空。'}
  MinVal *string `json:"minVal,omitempty" xml:"minVal,omitempty"`
  // {'en':'type = Integer/Number:maximum value, type = String:maximum length, type = Boolean/Enumeration:leave blank.', 'zh_CN':'type = Integer/Number：最大值，type = String：最大长度，type = Boolean/Enumeration：置空。'}
  MaxVal *string `json:"maxVal,omitempty" xml:"maxVal,omitempty"`
  // {'en':'Required.
  // TRUE:Yes
  // FALSE:No', 'zh_CN':'必带参数。
  // TRUE：是
  // FALSE：否'}
  Required *string `json:"required,omitempty" xml:"required,omitempty"`
  // {'en':'Content (maximum 2000 characters), required when type = Enumeration, multiple separated by ; sign.', 'zh_CN':'内容（最多2000个字符），type = Enumeration 时必填，多个以 ; 号分隔。'}
  Content *string `json:"content,omitempty" xml:"content,omitempty"`
  // {'en':'Child node array (JSON string), the structure is consistent with the parent node, required when type= Array.', 'zh_CN':'子节点数组（JSON字符串），结构与父结点一致，type = Array 时必填。'}
  Children []*string `json:"children,omitempty" xml:"children,omitempty" type:"Repeated"`
}

func (s CreateAPIDefinitionApiDefineParamLimitJsonAddChildrenDTO) String() string {
  return tea.Prettify(s)
}

func (s CreateAPIDefinitionApiDefineParamLimitJsonAddChildrenDTO) GoString() string {
  return s.String()
}

func (s *CreateAPIDefinitionApiDefineParamLimitJsonAddChildrenDTO) SetName(v string) *CreateAPIDefinitionApiDefineParamLimitJsonAddChildrenDTO {
  s.Name = &v
  return s
}

func (s *CreateAPIDefinitionApiDefineParamLimitJsonAddChildrenDTO) SetType(v string) *CreateAPIDefinitionApiDefineParamLimitJsonAddChildrenDTO {
  s.Type = &v
  return s
}

func (s *CreateAPIDefinitionApiDefineParamLimitJsonAddChildrenDTO) SetMinVal(v string) *CreateAPIDefinitionApiDefineParamLimitJsonAddChildrenDTO {
  s.MinVal = &v
  return s
}

func (s *CreateAPIDefinitionApiDefineParamLimitJsonAddChildrenDTO) SetMaxVal(v string) *CreateAPIDefinitionApiDefineParamLimitJsonAddChildrenDTO {
  s.MaxVal = &v
  return s
}

func (s *CreateAPIDefinitionApiDefineParamLimitJsonAddChildrenDTO) SetRequired(v string) *CreateAPIDefinitionApiDefineParamLimitJsonAddChildrenDTO {
  s.Required = &v
  return s
}

func (s *CreateAPIDefinitionApiDefineParamLimitJsonAddChildrenDTO) SetContent(v string) *CreateAPIDefinitionApiDefineParamLimitJsonAddChildrenDTO {
  s.Content = &v
  return s
}

func (s *CreateAPIDefinitionApiDefineParamLimitJsonAddChildrenDTO) SetChildren(v []*string) *CreateAPIDefinitionApiDefineParamLimitJsonAddChildrenDTO {
  s.Children = v
  return s
}

type CreateAPIDefinitionResponse struct {
  // {'en':'Please refer to the error code for exceptions.', 'zh_CN':'请参照错误码。','dictionary':'belong=WAAP-MS-Ext|dict=waap_retCodeEnum'}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {'en':'Description.', 'zh_CN':'描述信息。'}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {'en':'API definition ID.', 'zh_CN':'API定义ID。'}
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

type CreateAPIDefinitionRequestHeader struct {
  // {"zh_CN":"安全服务类型。有使用多个不同的安全服务时，需要填写具体的服务类型。","en":"Security service type. Please enter a specific service type, if you purchase multiple security services.","dictionary":"belong=WAAP-MS-Ext|dict=waap_serviceType"}
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

type CreateAPIDefinitionResponseHeader struct {
}

func (s CreateAPIDefinitionResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s CreateAPIDefinitionResponseHeader) GoString() string {
  return s.String()
}




type ListAPIDefinitionsRequest struct {
  // {'en':'Hostname array, by default all valid hostnames.', 'zh_CN':'域名数组，默认所有生效域名。'}
  DomainList []*string `json:"domainList,omitempty" xml:"domainList,omitempty" type:"Repeated"`
  // {'en':'API name array.', 'zh_CN':'API名称数组。'}
  NameList []*string `json:"nameList,omitempty" xml:"nameList,omitempty" type:"Repeated"`
  // {'en':'API grouped array.', 'zh_CN':'API分组数组。'}
  GroupNameList []*string `json:"groupNameList,omitempty" xml:"groupNameList,omitempty" type:"Repeated"`
  // {'en':'Sorting method, reverse order by default.
  // true: Modify time in reverse order
  // false: Modify time sequence', 'zh_CN':'排序方式，默认倒序。
  // true：修改时间倒序
  // false：修改时间正序'}
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

type ListAPIDefinitionsResponse struct {
  // {'en':'Please refer to the error code for exceptions.', 'zh_CN':'请参照错误码。','dictionary':'belong=WAAP-MS-Ext|dict=waap_retCodeEnum'}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {'en':'Description.', 'zh_CN':'描述信息。'}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {'en':'Data.', 'zh_CN':'出参数据。'}
  Data []*ListAPIDefinitionsApiDefineBasicVO `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
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

func (s *ListAPIDefinitionsResponse) SetData(v []*ListAPIDefinitionsApiDefineBasicVO) *ListAPIDefinitionsResponse {
  s.Data = v
  return s
}

type ListAPIDefinitionsApiDefineBasicVO struct {
  // {'en':'API definition ID.', 'zh_CN':'API定义ID。'}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {'en':'API name.', 'zh_CN':'API名称。'}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {'en':'API groups.', 'zh_CN':'API分组。'}
  GroupName *string `json:"groupName,omitempty" xml:"groupName,omitempty" require:"true"`
  // {'en':'Description.', 'zh_CN':'备注。'}
  Remark *string `json:"remark,omitempty" xml:"remark,omitempty" require:"true"`
  // {'en':'Attributed hostname.', 'zh_CN':'归属域名。'}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {'en':'API base path.', 'zh_CN':'端点路径。'}
  Path *string `json:"path,omitempty" xml:"path,omitempty" require:"true"`
  // {'en':'Creation time, format: yyyy-MM-dd HH:mm:ss.', 'zh_CN':'创建时间，格式：yyyy-MM-dd HH:mm:ss。'}
  CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty" require:"true"`
  // {'en':'Update time, format: yyyy-MM-dd HH:mm:ss.', 'zh_CN':'更新时间，格式：yyyy-MM-dd HH:mm:ss。'}
  UpdateTime *string `json:"updateTime,omitempty" xml:"updateTime,omitempty" require:"true"`
}

func (s ListAPIDefinitionsApiDefineBasicVO) String() string {
  return tea.Prettify(s)
}

func (s ListAPIDefinitionsApiDefineBasicVO) GoString() string {
  return s.String()
}

func (s *ListAPIDefinitionsApiDefineBasicVO) SetId(v string) *ListAPIDefinitionsApiDefineBasicVO {
  s.Id = &v
  return s
}

func (s *ListAPIDefinitionsApiDefineBasicVO) SetName(v string) *ListAPIDefinitionsApiDefineBasicVO {
  s.Name = &v
  return s
}

func (s *ListAPIDefinitionsApiDefineBasicVO) SetGroupName(v string) *ListAPIDefinitionsApiDefineBasicVO {
  s.GroupName = &v
  return s
}

func (s *ListAPIDefinitionsApiDefineBasicVO) SetRemark(v string) *ListAPIDefinitionsApiDefineBasicVO {
  s.Remark = &v
  return s
}

func (s *ListAPIDefinitionsApiDefineBasicVO) SetDomain(v string) *ListAPIDefinitionsApiDefineBasicVO {
  s.Domain = &v
  return s
}

func (s *ListAPIDefinitionsApiDefineBasicVO) SetPath(v string) *ListAPIDefinitionsApiDefineBasicVO {
  s.Path = &v
  return s
}

func (s *ListAPIDefinitionsApiDefineBasicVO) SetCreateTime(v string) *ListAPIDefinitionsApiDefineBasicVO {
  s.CreateTime = &v
  return s
}

func (s *ListAPIDefinitionsApiDefineBasicVO) SetUpdateTime(v string) *ListAPIDefinitionsApiDefineBasicVO {
  s.UpdateTime = &v
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

type ListAPIDefinitionsRequestHeader struct {
  // {"zh_CN":"安全服务类型。有使用多个不同的安全服务时，需要填写具体的服务类型。","en":"Security service type. Please enter a specific service type, if you purchase multiple security services.","dictionary":"belong=WAAP-MS-Ext|dict=waap_serviceType"}
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

type ListAPIDefinitionsResponseHeader struct {
}

func (s ListAPIDefinitionsResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ListAPIDefinitionsResponseHeader) GoString() string {
  return s.String()
}




