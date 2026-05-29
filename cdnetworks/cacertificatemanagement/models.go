package cacertificatemanagement

// This file is auto-generated, don't edit it. Thanks.
import (
  "github.com/alibabacloud-go/tea/tea"
)

type DeleteACaCertificateRequest struct {
}

func (s DeleteACaCertificateRequest) String() string {
  return tea.Prettify(s)
}

func (s DeleteACaCertificateRequest) GoString() string {
  return s.String()
}

type DeleteACaCertificateResponse struct {
  // {"en":"Response code, 0 means successful.", "zh_CN":"接口响应code，0代表成功。"}
  Code *int `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response error message if failed.", "zh_CN":"接口响应信息，success代表成功，失败则提供失败信息。"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s DeleteACaCertificateResponse) String() string {
  return tea.Prettify(s)
}

func (s DeleteACaCertificateResponse) GoString() string {
  return s.String()
}

func (s *DeleteACaCertificateResponse) SetCode(v int) *DeleteACaCertificateResponse {
  s.Code = &v
  return s
}

func (s *DeleteACaCertificateResponse) SetMessage(v string) *DeleteACaCertificateResponse {
  s.Message = &v
  return s
}

type DeleteACaCertificatePaths struct {
  // {"en":"The certificate ID you want to delete.", "zh_CN":"需要删除的证书id"}
  CertificateId *int `json:"certificateId,omitempty" xml:"certificateId,omitempty" require:"true"`
}

func (s DeleteACaCertificatePaths) String() string {
  return tea.Prettify(s)
}

func (s DeleteACaCertificatePaths) GoString() string {
  return s.String()
}

func (s *DeleteACaCertificatePaths) SetCertificateId(v int) *DeleteACaCertificatePaths {
  s.CertificateId = &v
  return s
}

type DeleteACaCertificateParameters struct {
}

func (s DeleteACaCertificateParameters) String() string {
  return tea.Prettify(s)
}

func (s DeleteACaCertificateParameters) GoString() string {
  return s.String()
}

type DeleteACaCertificateRequestHeader struct {
}

func (s DeleteACaCertificateRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteACaCertificateRequestHeader) GoString() string {
  return s.String()
}

type DeleteACaCertificateResponseHeader struct {
}

func (s DeleteACaCertificateResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteACaCertificateResponseHeader) GoString() string {
  return s.String()
}




type UpdateACaCertificateRequest struct {
  // {"en":"Certificate name","zh_CN":"证书名称"}
  CertificateName *string `json:"certificateName,omitempty" xml:"certificateName,omitempty" require:"true"`
  // {"en":"Certificate content, in PEM format.","zh_CN":"证书内容，PEM格式。例如：\n-----BEGIN CERTIFICATE-----\n……\n-----END CERTIFICATE-----"}
  Content *string `json:"content,omitempty" xml:"content,omitempty" require:"true"`
  // {"en":"comment.","zh_CN":"备注"}
  Comment *string `json:"comment,omitempty" xml:"comment,omitempty"`
  // {"en":"Certificate type","zh_CN":"证书的类型。ROOT表示根证书，NODE表示子证书","exampleValue":"ROOT,NODE"}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
  // {"en":"parent certificate Id,if the certificate type is NODE, it cannot be empty.","zh_CN":"父证书id，如果证书类型为NODE，则必填"}
  ParentCertificateId *string `json:"parentCertificateId,omitempty" xml:"parentCertificateId,omitempty"`
  // {"en":"Skip the constraints that have expired.","zh_CN":"是否跳过校验证书已过期的约束","exampleValue":"true,false"}
  SkipExpired *bool `json:"skipExpired,omitempty" xml:"skipExpired,omitempty" require:"true"`
}

func (s UpdateACaCertificateRequest) String() string {
  return tea.Prettify(s)
}

func (s UpdateACaCertificateRequest) GoString() string {
  return s.String()
}

func (s *UpdateACaCertificateRequest) SetCertificateName(v string) *UpdateACaCertificateRequest {
  s.CertificateName = &v
  return s
}

func (s *UpdateACaCertificateRequest) SetContent(v string) *UpdateACaCertificateRequest {
  s.Content = &v
  return s
}

func (s *UpdateACaCertificateRequest) SetComment(v string) *UpdateACaCertificateRequest {
  s.Comment = &v
  return s
}

func (s *UpdateACaCertificateRequest) SetType(v string) *UpdateACaCertificateRequest {
  s.Type = &v
  return s
}

func (s *UpdateACaCertificateRequest) SetParentCertificateId(v string) *UpdateACaCertificateRequest {
  s.ParentCertificateId = &v
  return s
}

func (s *UpdateACaCertificateRequest) SetSkipExpired(v bool) *UpdateACaCertificateRequest {
  s.SkipExpired = &v
  return s
}

type UpdateACaCertificateRequestHeader struct {
}

func (s UpdateACaCertificateRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s UpdateACaCertificateRequestHeader) GoString() string {
  return s.String()
}

type UpdateACaCertificatePaths struct {
  // {"en":"The certificate ID you want to modify.","zh_CN":"需要修改的证书id"}
  CertificateId *string `json:"certificateId,omitempty" xml:"certificateId,omitempty" require:"true"`
}

func (s UpdateACaCertificatePaths) String() string {
  return tea.Prettify(s)
}

func (s UpdateACaCertificatePaths) GoString() string {
  return s.String()
}

func (s *UpdateACaCertificatePaths) SetCertificateId(v string) *UpdateACaCertificatePaths {
  s.CertificateId = &v
  return s
}

type UpdateACaCertificateParameters struct {
}

func (s UpdateACaCertificateParameters) String() string {
  return tea.Prettify(s)
}

func (s UpdateACaCertificateParameters) GoString() string {
  return s.String()
}

type UpdateACaCertificateResponse struct {
  // {"en":"Status code","zh_CN":"状态码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response message","zh_CN":"响应信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s UpdateACaCertificateResponse) String() string {
  return tea.Prettify(s)
}

func (s UpdateACaCertificateResponse) GoString() string {
  return s.String()
}

func (s *UpdateACaCertificateResponse) SetCode(v string) *UpdateACaCertificateResponse {
  s.Code = &v
  return s
}

func (s *UpdateACaCertificateResponse) SetMessage(v string) *UpdateACaCertificateResponse {
  s.Message = &v
  return s
}

type UpdateACaCertificateResponseHeader struct {
}

func (s UpdateACaCertificateResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s UpdateACaCertificateResponseHeader) GoString() string {
  return s.String()
}




type GetCaCertificateRequest struct {
}

func (s GetCaCertificateRequest) String() string {
  return tea.Prettify(s)
}

func (s GetCaCertificateRequest) GoString() string {
  return s.String()
}

type GetCaCertificateRequestHeader struct {
}

func (s GetCaCertificateRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetCaCertificateRequestHeader) GoString() string {
  return s.String()
}

type GetCaCertificatePaths struct {
  // {"en":"CA certificate ID","zh_CN":"CA证书ID"}
  CertificateId *string `json:"certificate-id,omitempty" xml:"certificate-id,omitempty" require:"true"`
}

func (s GetCaCertificatePaths) String() string {
  return tea.Prettify(s)
}

func (s GetCaCertificatePaths) GoString() string {
  return s.String()
}

func (s *GetCaCertificatePaths) SetCertificateId(v string) *GetCaCertificatePaths {
  s.CertificateId = &v
  return s
}

type GetCaCertificateParameters struct {
}

func (s GetCaCertificateParameters) String() string {
  return tea.Prettify(s)
}

func (s GetCaCertificateParameters) GoString() string {
  return s.String()
}

type GetCaCertificateResponse struct {
  // {"en":"Status code","zh_CN":"状态码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response message","zh_CN":"响应信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"detail","zh_CN":"详情"}
  Data *GetCaCertificateResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
  // {"en":"Certificate ID.","zh_CN":"证书ID。"}
  CertificateId *string `json:"certificateId,omitempty" xml:"certificateId,omitempty" require:"true"`
  // {"en":"Certificate name.","zh_CN":"证书名称。"}
  CertificateName *string `json:"certificateName,omitempty" xml:"certificateName,omitempty" require:"true"`
  // {"en":"Certificate type","zh_CN":"证书类型","exampleValue":"ROOT,NODE"}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
  // {"en":"Parent certificate Id.","zh_CN":"父证书id。"}
  ParentCertificateId *string `json:"parentCertificateId,omitempty" xml:"parentCertificateId,omitempty" require:"true"`
  // {"en":"Parent certificate name.","zh_CN":"父证书名称。"}
  ParentCertificateName *string `json:"parentCertificateName,omitempty" xml:"parentCertificateName,omitempty" require:"true"`
  // {"en":"Certificate validity from.","zh_CN":"证书有效期开始时间。"}
  ValidityFrom *string `json:"validityFrom,omitempty" xml:"validityFrom,omitempty" require:"true"`
  // {"en":"Certificate validity to.","zh_CN":"证书有效期结束时间。"}
  ValidityTo *string `json:"validityTo,omitempty" xml:"validityTo,omitempty" require:"true"`
  // {"en":"Certificate create time.","zh_CN":"证书创建时间。"}
  CreationTime *string `json:"creationTime,omitempty" xml:"creationTime,omitempty" require:"true"`
  // {"en":"algorithm","zh_CN":"私钥算法。"}
  Algorithm *string `json:"algorithm,omitempty" xml:"algorithm,omitempty" require:"true"`
  // {"en":"version","zh_CN":"version"}
  Version *string `json:"version,omitempty" xml:"version,omitempty" require:"true"`
  // {"en":"issuer","zh_CN":"颁发机构"}
  Issuer *string `json:"issuer,omitempty" xml:"issuer,omitempty" require:"true"`
  // {"en":"serial number","zh_CN":"序列号"}
  SerialNumber *string `json:"serialNumber,omitempty" xml:"serialNumber,omitempty" require:"true"`
  // {"en":"certificate content","zh_CN":"证书内容"}
  Content *string `json:"content,omitempty" xml:"content,omitempty" require:"true"`
  // {"en":"Certificate expire status, 0-normal, 1-nearly expired, 2-already expired.","zh_CN":"证书过期状态，0-正常，1-临近过期，2-已过期。","exampleValue":"0,1,2"}
  ExpireStatus *string `json:"expireStatus,omitempty" xml:"expireStatus,omitempty" require:"true"`
  // {"en":"comment.","zh_CN":"备注。"}
  Comment *string `json:"comment,omitempty" xml:"comment,omitempty" require:"true"`
}

func (s GetCaCertificateResponse) String() string {
  return tea.Prettify(s)
}

func (s GetCaCertificateResponse) GoString() string {
  return s.String()
}

func (s *GetCaCertificateResponse) SetCode(v string) *GetCaCertificateResponse {
  s.Code = &v
  return s
}

func (s *GetCaCertificateResponse) SetMessage(v string) *GetCaCertificateResponse {
  s.Message = &v
  return s
}

func (s *GetCaCertificateResponse) SetData(v *GetCaCertificateResponseData) *GetCaCertificateResponse {
  s.Data = v
  return s
}

func (s *GetCaCertificateResponse) SetCertificateId(v string) *GetCaCertificateResponse {
  s.CertificateId = &v
  return s
}

func (s *GetCaCertificateResponse) SetCertificateName(v string) *GetCaCertificateResponse {
  s.CertificateName = &v
  return s
}

func (s *GetCaCertificateResponse) SetType(v string) *GetCaCertificateResponse {
  s.Type = &v
  return s
}

func (s *GetCaCertificateResponse) SetParentCertificateId(v string) *GetCaCertificateResponse {
  s.ParentCertificateId = &v
  return s
}

func (s *GetCaCertificateResponse) SetParentCertificateName(v string) *GetCaCertificateResponse {
  s.ParentCertificateName = &v
  return s
}

func (s *GetCaCertificateResponse) SetValidityFrom(v string) *GetCaCertificateResponse {
  s.ValidityFrom = &v
  return s
}

func (s *GetCaCertificateResponse) SetValidityTo(v string) *GetCaCertificateResponse {
  s.ValidityTo = &v
  return s
}

func (s *GetCaCertificateResponse) SetCreationTime(v string) *GetCaCertificateResponse {
  s.CreationTime = &v
  return s
}

func (s *GetCaCertificateResponse) SetAlgorithm(v string) *GetCaCertificateResponse {
  s.Algorithm = &v
  return s
}

func (s *GetCaCertificateResponse) SetVersion(v string) *GetCaCertificateResponse {
  s.Version = &v
  return s
}

func (s *GetCaCertificateResponse) SetIssuer(v string) *GetCaCertificateResponse {
  s.Issuer = &v
  return s
}

func (s *GetCaCertificateResponse) SetSerialNumber(v string) *GetCaCertificateResponse {
  s.SerialNumber = &v
  return s
}

func (s *GetCaCertificateResponse) SetContent(v string) *GetCaCertificateResponse {
  s.Content = &v
  return s
}

func (s *GetCaCertificateResponse) SetExpireStatus(v string) *GetCaCertificateResponse {
  s.ExpireStatus = &v
  return s
}

func (s *GetCaCertificateResponse) SetComment(v string) *GetCaCertificateResponse {
  s.Comment = &v
  return s
}

type GetCaCertificateResponseData struct {
}

func (s GetCaCertificateResponseData) String() string {
  return tea.Prettify(s)
}

func (s GetCaCertificateResponseData) GoString() string {
  return s.String()
}

type GetCaCertificateResponseHeader struct {
}

func (s GetCaCertificateResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetCaCertificateResponseHeader) GoString() string {
  return s.String()
}




type AssociateDomainWithCaCertificateRequest struct {
  // {"en":"Certificate ID","zh_CN":"证书ID"}
  CertificateId *string `json:"certificateId,omitempty" xml:"certificateId,omitempty" require:"true"`
  // {"en":"Certificate usage, Allowed values: CLIENT_MTLS, ORIGIN_MTLS,VERIFY_ORIGIN","zh_CN":"证书用途，取值范围：CLIENT_MTLS、ORIGIN_MTLS、VERIFY_ORIGIN"}
  Usage *string `json:"usage,omitempty" xml:"usage,omitempty" require:"true"`
  // {"en":"Associated domains.","zh_CN":"关联域名列表"}
  Domains []*string `json:"domains,omitempty" xml:"domains,omitempty" require:"true" type:"Repeated"`
}

func (s AssociateDomainWithCaCertificateRequest) String() string {
  return tea.Prettify(s)
}

func (s AssociateDomainWithCaCertificateRequest) GoString() string {
  return s.String()
}

func (s *AssociateDomainWithCaCertificateRequest) SetCertificateId(v string) *AssociateDomainWithCaCertificateRequest {
  s.CertificateId = &v
  return s
}

func (s *AssociateDomainWithCaCertificateRequest) SetUsage(v string) *AssociateDomainWithCaCertificateRequest {
  s.Usage = &v
  return s
}

func (s *AssociateDomainWithCaCertificateRequest) SetDomains(v []*string) *AssociateDomainWithCaCertificateRequest {
  s.Domains = v
  return s
}

type AssociateDomainWithCaCertificateRequestHeader struct {
}

func (s AssociateDomainWithCaCertificateRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s AssociateDomainWithCaCertificateRequestHeader) GoString() string {
  return s.String()
}

type AssociateDomainWithCaCertificatePaths struct {
}

func (s AssociateDomainWithCaCertificatePaths) String() string {
  return tea.Prettify(s)
}

func (s AssociateDomainWithCaCertificatePaths) GoString() string {
  return s.String()
}

type AssociateDomainWithCaCertificateParameters struct {
}

func (s AssociateDomainWithCaCertificateParameters) String() string {
  return tea.Prettify(s)
}

func (s AssociateDomainWithCaCertificateParameters) GoString() string {
  return s.String()
}

type AssociateDomainWithCaCertificateResponse struct {
  // {"en":"Status code","zh_CN":"状态码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response message","zh_CN":"响应信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s AssociateDomainWithCaCertificateResponse) String() string {
  return tea.Prettify(s)
}

func (s AssociateDomainWithCaCertificateResponse) GoString() string {
  return s.String()
}

func (s *AssociateDomainWithCaCertificateResponse) SetCode(v string) *AssociateDomainWithCaCertificateResponse {
  s.Code = &v
  return s
}

func (s *AssociateDomainWithCaCertificateResponse) SetMessage(v string) *AssociateDomainWithCaCertificateResponse {
  s.Message = &v
  return s
}

type AssociateDomainWithCaCertificateResponseHeader struct {
}

func (s AssociateDomainWithCaCertificateResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s AssociateDomainWithCaCertificateResponseHeader) GoString() string {
  return s.String()
}




type CreateACaCertificateRequest struct {
  // {"en":"Certificate name","zh_CN":"证书名称"}
  CertificateName *string `json:"certificateName,omitempty" xml:"certificateName,omitempty" require:"true"`
  // {"en":"Certificate content, in PEM format.","zh_CN":"证书内容，PEM格式。例如：\n-----BEGIN CERTIFICATE-----\n……\n-----END CERTIFICATE-----"}
  Content *string `json:"content,omitempty" xml:"content,omitempty" require:"true"`
  // {"en":"comment.","zh_CN":"备注"}
  Comment *string `json:"comment,omitempty" xml:"comment,omitempty"`
  // {"en":"Certificate type","zh_CN":"证书的类型。ROOT表示根证书，NODE表示子证书","exampleValue":"ROOT,NODE"}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
  // {"en":"parent certificate Id,if the certificate type is NODE, it cannot be empty.","zh_CN":"父证书ID，如果证书类型为NODE，则必填"}
  ParentCertificateId *string `json:"parentCertificateId,omitempty" xml:"parentCertificateId,omitempty"`
  // {"en":"Skip the constraints that have expired.","zh_CN":"是否跳过校验证书已过期的约束"}
  SkipExpired *bool `json:"skipExpired,omitempty" xml:"skipExpired,omitempty" require:"true"`
}

func (s CreateACaCertificateRequest) String() string {
  return tea.Prettify(s)
}

func (s CreateACaCertificateRequest) GoString() string {
  return s.String()
}

func (s *CreateACaCertificateRequest) SetCertificateName(v string) *CreateACaCertificateRequest {
  s.CertificateName = &v
  return s
}

func (s *CreateACaCertificateRequest) SetContent(v string) *CreateACaCertificateRequest {
  s.Content = &v
  return s
}

func (s *CreateACaCertificateRequest) SetComment(v string) *CreateACaCertificateRequest {
  s.Comment = &v
  return s
}

func (s *CreateACaCertificateRequest) SetType(v string) *CreateACaCertificateRequest {
  s.Type = &v
  return s
}

func (s *CreateACaCertificateRequest) SetParentCertificateId(v string) *CreateACaCertificateRequest {
  s.ParentCertificateId = &v
  return s
}

func (s *CreateACaCertificateRequest) SetSkipExpired(v bool) *CreateACaCertificateRequest {
  s.SkipExpired = &v
  return s
}

type CreateACaCertificateRequestHeader struct {
}

func (s CreateACaCertificateRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s CreateACaCertificateRequestHeader) GoString() string {
  return s.String()
}

type CreateACaCertificatePaths struct {
}

func (s CreateACaCertificatePaths) String() string {
  return tea.Prettify(s)
}

func (s CreateACaCertificatePaths) GoString() string {
  return s.String()
}

type CreateACaCertificateParameters struct {
}

func (s CreateACaCertificateParameters) String() string {
  return tea.Prettify(s)
}

func (s CreateACaCertificateParameters) GoString() string {
  return s.String()
}

type CreateACaCertificateResponse struct {
  // {"en":"The URL used to access the certificate file where certificate-id is the unique token that the system generates for the certificate and whose value is a string","zh_CN":"用于访问该证书文件的URL，其中certificate-id为系统为该证书生成的唯一标示，其值为字符串"}
  Location *string `json:"location,omitempty" xml:"location,omitempty" require:"true"`
  // {"en":"Status code","zh_CN":"状态码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response message","zh_CN":"响应信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s CreateACaCertificateResponse) String() string {
  return tea.Prettify(s)
}

func (s CreateACaCertificateResponse) GoString() string {
  return s.String()
}

func (s *CreateACaCertificateResponse) SetLocation(v string) *CreateACaCertificateResponse {
  s.Location = &v
  return s
}

func (s *CreateACaCertificateResponse) SetCode(v string) *CreateACaCertificateResponse {
  s.Code = &v
  return s
}

func (s *CreateACaCertificateResponse) SetMessage(v string) *CreateACaCertificateResponse {
  s.Message = &v
  return s
}

type CreateACaCertificateResponseHeader struct {
}

func (s CreateACaCertificateResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s CreateACaCertificateResponseHeader) GoString() string {
  return s.String()
}




type ListCaCertificatesRequest struct {
  // {"en":"Certificate name","zh_CN":"证书名称"}
  CertificateName *string `json:"certificateName,omitempty" xml:"certificateName,omitempty"`
  // {"en":"Page number, must great than 0.","zh_CN":"当前页数,大于0"}
  PageNumber *int `json:"pageNumber,omitempty" xml:"pageNumber,omitempty" require:"true"`
  // {"en":"Page size,must great than 0 and not allowed to exceed 500.","zh_CN":"每页记录数,大于0小于500"}
  PageSize *int `json:"pageSize,omitempty" xml:"pageSize,omitempty" require:"true"`
  // {"en":"related domain names.","zh_CN":"关联的域名"}
  RelateDomains []*string `json:"relateDomains,omitempty" xml:"relateDomains,omitempty" type:"Repeated"`
  // {"en":"Certificate expire status, 0-normal, 1-nearly expired, 2-already expired.","zh_CN":"证书过期状态，0-正常，1-临近过期，2-已过期","exampleValue":"0,1,2"}
  ExpireStatus []*string `json:"expireStatus,omitempty" xml:"expireStatus,omitempty" type:"Repeated"`
  // {"en":"algorithm.","zh_CN":"私钥算法"}
  Algorithm *string `json:"algorithm,omitempty" xml:"algorithm,omitempty"`
  // {"en":"Match mode, effects on parameter certificateName and relateDomains.","zh_CN":"匹配方式,作用于证书名称以及关联域名"}
  IsLike *string `json:"isLike,omitempty" xml:"isLike,omitempty" require:"true"`
}

func (s ListCaCertificatesRequest) String() string {
  return tea.Prettify(s)
}

func (s ListCaCertificatesRequest) GoString() string {
  return s.String()
}

func (s *ListCaCertificatesRequest) SetCertificateName(v string) *ListCaCertificatesRequest {
  s.CertificateName = &v
  return s
}

func (s *ListCaCertificatesRequest) SetPageNumber(v int) *ListCaCertificatesRequest {
  s.PageNumber = &v
  return s
}

func (s *ListCaCertificatesRequest) SetPageSize(v int) *ListCaCertificatesRequest {
  s.PageSize = &v
  return s
}

func (s *ListCaCertificatesRequest) SetRelateDomains(v []*string) *ListCaCertificatesRequest {
  s.RelateDomains = v
  return s
}

func (s *ListCaCertificatesRequest) SetExpireStatus(v []*string) *ListCaCertificatesRequest {
  s.ExpireStatus = v
  return s
}

func (s *ListCaCertificatesRequest) SetAlgorithm(v string) *ListCaCertificatesRequest {
  s.Algorithm = &v
  return s
}

func (s *ListCaCertificatesRequest) SetIsLike(v string) *ListCaCertificatesRequest {
  s.IsLike = &v
  return s
}

type ListCaCertificatesRequestHeader struct {
}

func (s ListCaCertificatesRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ListCaCertificatesRequestHeader) GoString() string {
  return s.String()
}

type ListCaCertificatesPaths struct {
}

func (s ListCaCertificatesPaths) String() string {
  return tea.Prettify(s)
}

func (s ListCaCertificatesPaths) GoString() string {
  return s.String()
}

type ListCaCertificatesParameters struct {
}

func (s ListCaCertificatesParameters) String() string {
  return tea.Prettify(s)
}

func (s ListCaCertificatesParameters) GoString() string {
  return s.String()
}

type ListCaCertificatesResponse struct {
  // {"en":"Result Code, success is 0","zh_CN":"响应码，成功为0"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Result Message","zh_CN":"响应信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Certificate list","zh_CN":"ca证书列表"}
  Certificates []*ListCaCertificatesResponseCertificates `json:"certificates,omitempty" xml:"certificates,omitempty" require:"true" type:"Repeated"`
  // {"en":"Page Info","zh_CN":"分页数据"}
  PageInfo *ListCaCertificatesResponsePageInfo `json:"pageInfo,omitempty" xml:"pageInfo,omitempty" require:"true" type:"Struct"`
}

func (s ListCaCertificatesResponse) String() string {
  return tea.Prettify(s)
}

func (s ListCaCertificatesResponse) GoString() string {
  return s.String()
}

func (s *ListCaCertificatesResponse) SetCode(v string) *ListCaCertificatesResponse {
  s.Code = &v
  return s
}

func (s *ListCaCertificatesResponse) SetMessage(v string) *ListCaCertificatesResponse {
  s.Message = &v
  return s
}

func (s *ListCaCertificatesResponse) SetCertificates(v []*ListCaCertificatesResponseCertificates) *ListCaCertificatesResponse {
  s.Certificates = v
  return s
}

func (s *ListCaCertificatesResponse) SetPageInfo(v *ListCaCertificatesResponsePageInfo) *ListCaCertificatesResponse {
  s.PageInfo = v
  return s
}

type ListCaCertificatesResponseCertificates struct     {
  // {"en":"Certificate ID.","zh_CN":"证书ID。"}
  CertificateId *string `json:"certificateId,omitempty" xml:"certificateId,omitempty" require:"true"`
  // {"en":"Certificate name.","zh_CN":"证书名称。"}
  CertificateName *string `json:"certificateName,omitempty" xml:"certificateName,omitempty" require:"true"`
  // {"en":"Certificate type, ROOT or NODE.","zh_CN":"证书类型,取值ROOT或NODE。"}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
  // {"en":"Parent certificate Id.","zh_CN":"父证书id。"}
  ParentCertificateId *string `json:"parentCertificateId,omitempty" xml:"parentCertificateId,omitempty" require:"true"`
  // {"en":"Certificate validity from.","zh_CN":"证书有效期开始时间。"}
  ValidityFrom *string `json:"validityFrom,omitempty" xml:"validityFrom,omitempty" require:"true"`
  // {"en":"Certificate validity to.","zh_CN":"证书有效期结束时间。"}
  ValidityTo *string `json:"validityTo,omitempty" xml:"validityTo,omitempty" require:"true"`
  // {"en":"Related domain list.","zh_CN":"关联的域名。"}
  RelatedDomainList []*ListCaCertificatesResponseCertificatesRelatedDomainList `json:"relatedDomainList,omitempty" xml:"relatedDomainList,omitempty" require:"true" type:"Repeated"`
  // {"en":"algorithm","zh_CN":"私钥算法。"}
  Algorithm *string `json:"algorithm,omitempty" xml:"algorithm,omitempty" require:"true"`
  // {"en":"Certificate expire status, 0-normal, 1-nearly expired, 2-already expired.","zh_CN":"证书过期状态，0-正常，1-临近过期，2-已过期。"}
  ExpireStatus *string `json:"expireStatus,omitempty" xml:"expireStatus,omitempty" require:"true"`
  // {"en":"comment.","zh_CN":"备注。"}
  Comment *string `json:"comment,omitempty" xml:"comment,omitempty" require:"true"`
}

func (s ListCaCertificatesResponseCertificates) String() string {
  return tea.Prettify(s)
}

func (s ListCaCertificatesResponseCertificates) GoString() string {
  return s.String()
}

func (s *ListCaCertificatesResponseCertificates) SetCertificateId(v string) *ListCaCertificatesResponseCertificates {
  s.CertificateId = &v
  return s
}

func (s *ListCaCertificatesResponseCertificates) SetCertificateName(v string) *ListCaCertificatesResponseCertificates {
  s.CertificateName = &v
  return s
}

func (s *ListCaCertificatesResponseCertificates) SetType(v string) *ListCaCertificatesResponseCertificates {
  s.Type = &v
  return s
}

func (s *ListCaCertificatesResponseCertificates) SetParentCertificateId(v string) *ListCaCertificatesResponseCertificates {
  s.ParentCertificateId = &v
  return s
}

func (s *ListCaCertificatesResponseCertificates) SetValidityFrom(v string) *ListCaCertificatesResponseCertificates {
  s.ValidityFrom = &v
  return s
}

func (s *ListCaCertificatesResponseCertificates) SetValidityTo(v string) *ListCaCertificatesResponseCertificates {
  s.ValidityTo = &v
  return s
}

func (s *ListCaCertificatesResponseCertificates) SetRelatedDomainList(v []*ListCaCertificatesResponseCertificatesRelatedDomainList) *ListCaCertificatesResponseCertificates {
  s.RelatedDomainList = v
  return s
}

func (s *ListCaCertificatesResponseCertificates) SetAlgorithm(v string) *ListCaCertificatesResponseCertificates {
  s.Algorithm = &v
  return s
}

func (s *ListCaCertificatesResponseCertificates) SetExpireStatus(v string) *ListCaCertificatesResponseCertificates {
  s.ExpireStatus = &v
  return s
}

func (s *ListCaCertificatesResponseCertificates) SetComment(v string) *ListCaCertificatesResponseCertificates {
  s.Comment = &v
  return s
}

type ListCaCertificatesResponseCertificatesRelatedDomainList struct     {
  // {"en":"Certificate usage, CLIENT_MTLS or ORIGIN_MTLS.","zh_CN":"证书用途，取值CLIENT_MTLS或ORIGIN_MTLS"}
  Usage *string `json:"usage,omitempty" xml:"usage,omitempty" require:"true"`
  // {"en":"Domain list","zh_CN":"域名列表"}
  Domains []*string `json:"domains,omitempty" xml:"domains,omitempty" require:"true" type:"Repeated"`
}

func (s ListCaCertificatesResponseCertificatesRelatedDomainList) String() string {
  return tea.Prettify(s)
}

func (s ListCaCertificatesResponseCertificatesRelatedDomainList) GoString() string {
  return s.String()
}

func (s *ListCaCertificatesResponseCertificatesRelatedDomainList) SetUsage(v string) *ListCaCertificatesResponseCertificatesRelatedDomainList {
  s.Usage = &v
  return s
}

func (s *ListCaCertificatesResponseCertificatesRelatedDomainList) SetDomains(v []*string) *ListCaCertificatesResponseCertificatesRelatedDomainList {
  s.Domains = v
  return s
}

type ListCaCertificatesResponsePageInfo struct {
  // {"en":"Total Number","zh_CN":"总数"}
  TotalNumber *int `json:"totalNumber,omitempty" xml:"totalNumber,omitempty" require:"true"`
  // {"en":"Page Number","zh_CN":"页码"}
  PageNumber *int `json:"pageNumber,omitempty" xml:"pageNumber,omitempty" require:"true"`
  // {"en":"Page Size","zh_CN":"每页大小"}
  PageSize *int `json:"pageSize,omitempty" xml:"pageSize,omitempty" require:"true"`
  // {"en":"Total Page Number","zh_CN":"总页码数"}
  TotalPageNumber *int `json:"totalPageNumber,omitempty" xml:"totalPageNumber,omitempty" require:"true"`
}

func (s ListCaCertificatesResponsePageInfo) String() string {
  return tea.Prettify(s)
}

func (s ListCaCertificatesResponsePageInfo) GoString() string {
  return s.String()
}

func (s *ListCaCertificatesResponsePageInfo) SetTotalNumber(v int) *ListCaCertificatesResponsePageInfo {
  s.TotalNumber = &v
  return s
}

func (s *ListCaCertificatesResponsePageInfo) SetPageNumber(v int) *ListCaCertificatesResponsePageInfo {
  s.PageNumber = &v
  return s
}

func (s *ListCaCertificatesResponsePageInfo) SetPageSize(v int) *ListCaCertificatesResponsePageInfo {
  s.PageSize = &v
  return s
}

func (s *ListCaCertificatesResponsePageInfo) SetTotalPageNumber(v int) *ListCaCertificatesResponsePageInfo {
  s.TotalPageNumber = &v
  return s
}

type ListCaCertificatesResponseHeader struct {
}

func (s ListCaCertificatesResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ListCaCertificatesResponseHeader) GoString() string {
  return s.String()
}




type DisassociateDomainWithCaCertificateRequest struct {
  // {"en":"Certificate ID","zh_CN":"证书ID"}
  CertificateId *string `json:"certificateId,omitempty" xml:"certificateId,omitempty" require:"true"`
  // {"en":"Certificate usage, Allowed value: CLIENT_MTLS,ORIGIN_MTLS,VERIFY_ORIGIN .","zh_CN":"证书用途，取值范围：CLIENT_MTLS、ORIGIN_MTLS、 VERIFY_ORIGIN"}
  Usage *string `json:"usage,omitempty" xml:"usage,omitempty" require:"true"`
  // {"en":"Associated domains.","zh_CN":"关联域名列表"}
  Domains []*string `json:"domains,omitempty" xml:"domains,omitempty" require:"true" type:"Repeated"`
}

func (s DisassociateDomainWithCaCertificateRequest) String() string {
  return tea.Prettify(s)
}

func (s DisassociateDomainWithCaCertificateRequest) GoString() string {
  return s.String()
}

func (s *DisassociateDomainWithCaCertificateRequest) SetCertificateId(v string) *DisassociateDomainWithCaCertificateRequest {
  s.CertificateId = &v
  return s
}

func (s *DisassociateDomainWithCaCertificateRequest) SetUsage(v string) *DisassociateDomainWithCaCertificateRequest {
  s.Usage = &v
  return s
}

func (s *DisassociateDomainWithCaCertificateRequest) SetDomains(v []*string) *DisassociateDomainWithCaCertificateRequest {
  s.Domains = v
  return s
}

type DisassociateDomainWithCaCertificateRequestHeader struct {
}

func (s DisassociateDomainWithCaCertificateRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s DisassociateDomainWithCaCertificateRequestHeader) GoString() string {
  return s.String()
}

type DisassociateDomainWithCaCertificatePaths struct {
}

func (s DisassociateDomainWithCaCertificatePaths) String() string {
  return tea.Prettify(s)
}

func (s DisassociateDomainWithCaCertificatePaths) GoString() string {
  return s.String()
}

type DisassociateDomainWithCaCertificateParameters struct {
}

func (s DisassociateDomainWithCaCertificateParameters) String() string {
  return tea.Prettify(s)
}

func (s DisassociateDomainWithCaCertificateParameters) GoString() string {
  return s.String()
}

type DisassociateDomainWithCaCertificateResponse struct {
  // {"en":"Status code","zh_CN":"响应码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response message","zh_CN":"响应信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s DisassociateDomainWithCaCertificateResponse) String() string {
  return tea.Prettify(s)
}

func (s DisassociateDomainWithCaCertificateResponse) GoString() string {
  return s.String()
}

func (s *DisassociateDomainWithCaCertificateResponse) SetCode(v string) *DisassociateDomainWithCaCertificateResponse {
  s.Code = &v
  return s
}

func (s *DisassociateDomainWithCaCertificateResponse) SetMessage(v string) *DisassociateDomainWithCaCertificateResponse {
  s.Message = &v
  return s
}

type DisassociateDomainWithCaCertificateResponseHeader struct {
}

func (s DisassociateDomainWithCaCertificateResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s DisassociateDomainWithCaCertificateResponseHeader) GoString() string {
  return s.String()
}




