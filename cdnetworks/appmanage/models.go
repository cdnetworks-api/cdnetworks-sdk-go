package appmanage

// This file is auto-generated, don't edit it. Thanks.
import (
  "github.com/alibabacloud-go/tea/tea"
)

type CreateAppRequest struct {
  // {"en":"Debug fingerprint, only applicable to Android. If it matches the production fingerprint, it can be left blank.","zh_CN":"调试指纹，仅对Android生效，如果与正式指纹一致可不填写。"}
  DebugFingerprintList []*CreateAppRequestDebugFingerprintList `json:"debugFingerprintList,omitempty" xml:"debugFingerprintList,omitempty" type:"Repeated"`
  // {"en":"Application name should not exceed 60 characters and does not support characters ',\",<,>,&,/.","zh_CN":"应用名称，长度不超过60，不支持字符',\",<,>,&,/。"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"Official fingerprint, mandatory for Android applications. Please ensure this is the fingerprint for the officially launched application, otherwise authentication will fail and the service cannot be activated. You can check with technical support for the method to obtain the fingerprint. The total length should not exceed 100 characters and must include colons, uppercase and lowercase letters from A to F, and numbers from 0 to 9. After removing the colons, the length should be between 32 and 64.","zh_CN":"正式指纹，Android应用必填，请确保该指纹为正式上线应用的指纹，否则将鉴权失败无法启用服务。获取指纹方法可找技术支持确认。总长度不超过100，仅包含冒号、A~F大小写字母和0~9的数字，去除冒号后长度介于32~64之间。"}
  Fingerprint *string `json:"fingerprint,omitempty" xml:"fingerprint,omitempty"`
  // {"en":"Please fill in the package name carefully; otherwise, it will not be usable. The length must be between 3 and 100 characters, consisting of letters, numbers, underscores, and hyphens, with sections of the package name separated by periods. Format example: com.maa.test.","zh_CN":"包名，请认真填写包名，否则将无法使用。长度必须在 3 到 100 个字符之间，并且由字母、数字、下划线和连字符组成，包名的各部分之间用点号分隔。格式如：com.maa.test。"}
  PackageName *string `json:"packageName,omitempty" xml:"packageName,omitempty" require:"true"`
  // {"dictionary":"belong=MAA-masp-portal-console|dict=AppType","en":"Application Type.","zh_CN":"应用类型。"}
  Type *int `json:"type,omitempty" xml:"type,omitempty" require:"true"`
  // {"en":"Application Platform, Android or iOS.","zh_CN":"应用平台，Android、iOS。"}
  Platform *string `json:"platform,omitempty" xml:"platform,omitempty" require:"true"`
}

func (s CreateAppRequest) String() string {
  return tea.Prettify(s)
}

func (s CreateAppRequest) GoString() string {
  return s.String()
}

func (s *CreateAppRequest) SetDebugFingerprintList(v []*CreateAppRequestDebugFingerprintList) *CreateAppRequest {
  s.DebugFingerprintList = v
  return s
}

func (s *CreateAppRequest) SetName(v string) *CreateAppRequest {
  s.Name = &v
  return s
}

func (s *CreateAppRequest) SetFingerprint(v string) *CreateAppRequest {
  s.Fingerprint = &v
  return s
}

func (s *CreateAppRequest) SetPackageName(v string) *CreateAppRequest {
  s.PackageName = &v
  return s
}

func (s *CreateAppRequest) SetType(v int) *CreateAppRequest {
  s.Type = &v
  return s
}

func (s *CreateAppRequest) SetPlatform(v string) *CreateAppRequest {
  s.Platform = &v
  return s
}

type CreateAppRequestDebugFingerprintList struct     {
  // {"en":"Fingerprint, with a total length not exceeding 100, containing only colons, letters A to F in both uppercase and lowercase, and digits 0 to 9. After removing colons, the length should be between 32 and 64.","zh_CN":"指纹，总长度不超过100，仅包含冒号、A~F大小写字母和0~9的数字，去除冒号后长度介于32~64之间。"}
  Fingerprint *string `json:"fingerprint,omitempty" xml:"fingerprint,omitempty"`
  // {"en":"Description, no more than 60 characters.","zh_CN":"描述，长度不超过60。"}
  Desc *string `json:"desc,omitempty" xml:"desc,omitempty"`
}

func (s CreateAppRequestDebugFingerprintList) String() string {
  return tea.Prettify(s)
}

func (s CreateAppRequestDebugFingerprintList) GoString() string {
  return s.String()
}

func (s *CreateAppRequestDebugFingerprintList) SetFingerprint(v string) *CreateAppRequestDebugFingerprintList {
  s.Fingerprint = &v
  return s
}

func (s *CreateAppRequestDebugFingerprintList) SetDesc(v string) *CreateAppRequestDebugFingerprintList {
  s.Desc = &v
  return s
}

type CreateAppRequestHeader struct {
}

func (s CreateAppRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s CreateAppRequestHeader) GoString() string {
  return s.String()
}

type CreateAppPaths struct {
}

func (s CreateAppPaths) String() string {
  return tea.Prettify(s)
}

func (s CreateAppPaths) GoString() string {
  return s.String()
}

type CreateAppParameters struct {
}

func (s CreateAppParameters) String() string {
  return tea.Prettify(s)
}

func (s CreateAppParameters) GoString() string {
  return s.String()
}

type CreateAppResponse struct {
  // {"dictionary":"belong=MAA-masp-portal-console|dict=wplus_code","en":"Response Code","zh_CN":"响应码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Return data, Application ID","zh_CN":"返回数据，应用ID"}
  Data *int `json:"data,omitempty" xml:"data,omitempty" require:"true"`
  // {"en":"Response Description","zh_CN":"响应描述"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s CreateAppResponse) String() string {
  return tea.Prettify(s)
}

func (s CreateAppResponse) GoString() string {
  return s.String()
}

func (s *CreateAppResponse) SetCode(v string) *CreateAppResponse {
  s.Code = &v
  return s
}

func (s *CreateAppResponse) SetData(v int) *CreateAppResponse {
  s.Data = &v
  return s
}

func (s *CreateAppResponse) SetMessage(v string) *CreateAppResponse {
  s.Message = &v
  return s
}

type CreateAppResponseHeader struct {
}

func (s CreateAppResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s CreateAppResponseHeader) GoString() string {
  return s.String()
}




type AddDebugFingerprintRequest struct {
  // {"en":"Application ID","zh_CN":"应用ID"}
  AppId *int `json:"appId,omitempty" xml:"appId,omitempty" require:"true"`
  // {"en":"Debug fingerprint, only applicable to Android.","zh_CN":"调试指纹，仅对Android生效。"}
  DebugFingerprintList []*AddDebugFingerprintRequestDebugFingerprintList `json:"debugFingerprintList,omitempty" xml:"debugFingerprintList,omitempty" require:"true" type:"Repeated"`
}

func (s AddDebugFingerprintRequest) String() string {
  return tea.Prettify(s)
}

func (s AddDebugFingerprintRequest) GoString() string {
  return s.String()
}

func (s *AddDebugFingerprintRequest) SetAppId(v int) *AddDebugFingerprintRequest {
  s.AppId = &v
  return s
}

func (s *AddDebugFingerprintRequest) SetDebugFingerprintList(v []*AddDebugFingerprintRequestDebugFingerprintList) *AddDebugFingerprintRequest {
  s.DebugFingerprintList = v
  return s
}

type AddDebugFingerprintRequestDebugFingerprintList struct     {
  // {"en":"Fingerprint, with a total length not exceeding 100, containing only colons, letters A to F in both uppercase and lowercase, and digits 0 to 9. After removing colons, the length should be between 32 and 64.","zh_CN":"指纹，总长度不超过100，仅包含冒号、A~F大小写字母和0~9的数字，去除冒号后长度介于32~64之间。"}
  Fingerprint *string `json:"fingerprint,omitempty" xml:"fingerprint,omitempty" require:"true"`
  // {"en":"Description, no more than 60 characters.","zh_CN":"描述，长度不超过60。"}
  Desc *string `json:"desc,omitempty" xml:"desc,omitempty"`
}

func (s AddDebugFingerprintRequestDebugFingerprintList) String() string {
  return tea.Prettify(s)
}

func (s AddDebugFingerprintRequestDebugFingerprintList) GoString() string {
  return s.String()
}

func (s *AddDebugFingerprintRequestDebugFingerprintList) SetFingerprint(v string) *AddDebugFingerprintRequestDebugFingerprintList {
  s.Fingerprint = &v
  return s
}

func (s *AddDebugFingerprintRequestDebugFingerprintList) SetDesc(v string) *AddDebugFingerprintRequestDebugFingerprintList {
  s.Desc = &v
  return s
}

type AddDebugFingerprintRequestHeader struct {
}

func (s AddDebugFingerprintRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s AddDebugFingerprintRequestHeader) GoString() string {
  return s.String()
}

type AddDebugFingerprintPaths struct {
}

func (s AddDebugFingerprintPaths) String() string {
  return tea.Prettify(s)
}

func (s AddDebugFingerprintPaths) GoString() string {
  return s.String()
}

type AddDebugFingerprintParameters struct {
}

func (s AddDebugFingerprintParameters) String() string {
  return tea.Prettify(s)
}

func (s AddDebugFingerprintParameters) GoString() string {
  return s.String()
}

type AddDebugFingerprintResponse struct {
  // {"dictionary":"belong=MAA-masp-portal-console|dict=wplus_code","en":"Response Code","zh_CN":"响应码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response Description","zh_CN":"响应描述"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s AddDebugFingerprintResponse) String() string {
  return tea.Prettify(s)
}

func (s AddDebugFingerprintResponse) GoString() string {
  return s.String()
}

func (s *AddDebugFingerprintResponse) SetCode(v string) *AddDebugFingerprintResponse {
  s.Code = &v
  return s
}

func (s *AddDebugFingerprintResponse) SetMessage(v string) *AddDebugFingerprintResponse {
  s.Message = &v
  return s
}

type AddDebugFingerprintResponseHeader struct {
}

func (s AddDebugFingerprintResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s AddDebugFingerprintResponseHeader) GoString() string {
  return s.String()
}




type DeleteAppRequest struct {
  // {"en":"Application ID","zh_CN":"应用ID"}
  AppId *int `json:"appId,omitempty" xml:"appId,omitempty" require:"true"`
}

func (s DeleteAppRequest) String() string {
  return tea.Prettify(s)
}

func (s DeleteAppRequest) GoString() string {
  return s.String()
}

func (s *DeleteAppRequest) SetAppId(v int) *DeleteAppRequest {
  s.AppId = &v
  return s
}

type DeleteAppRequestHeader struct {
}

func (s DeleteAppRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteAppRequestHeader) GoString() string {
  return s.String()
}

type DeleteAppPaths struct {
}

func (s DeleteAppPaths) String() string {
  return tea.Prettify(s)
}

func (s DeleteAppPaths) GoString() string {
  return s.String()
}

type DeleteAppParameters struct {
}

func (s DeleteAppParameters) String() string {
  return tea.Prettify(s)
}

func (s DeleteAppParameters) GoString() string {
  return s.String()
}

type DeleteAppResponse struct {
  // {"dictionary":"belong=MAA-masp-portal-console|dict=wplus_code","en":"Response Code","zh_CN":"响应码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response Description","zh_CN":"响应描述"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s DeleteAppResponse) String() string {
  return tea.Prettify(s)
}

func (s DeleteAppResponse) GoString() string {
  return s.String()
}

func (s *DeleteAppResponse) SetCode(v string) *DeleteAppResponse {
  s.Code = &v
  return s
}

func (s *DeleteAppResponse) SetMessage(v string) *DeleteAppResponse {
  s.Message = &v
  return s
}

type DeleteAppResponseHeader struct {
}

func (s DeleteAppResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteAppResponseHeader) GoString() string {
  return s.String()
}




