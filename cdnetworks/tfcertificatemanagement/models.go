package tfcertificatemanagement

// This file is auto-generated, don't edit it. Thanks.
import (
  "github.com/alibabacloud-go/tea/tea"
)

type UpdateCertificateForTerraformRequest struct {
  // {"en":"Certificate name","zh_CN":"证书名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // {"en":"Certificate, PEM certificate.","zh_CN":"证书内容，PEM格式。例如：\n-----BEGIN CERTIFICATE-----\n……\n-----END CERTIFICATE-----"}
  Certificate *string `json:"certificate,omitempty" xml:"certificate,omitempty"`
  // {"en":"Private key of the certificate, PEM certificate. Not required when you specify a csrId.","zh_CN":"证书密钥，PEM格式。例如：\n-----BEGIN RSA PRIVATE KEY-----\n……\n-----BEGIN RSA PRIVATE KEY-----"}
  PrivateKey *string `json:"privateKey,omitempty" xml:"privateKey,omitempty"`
}

func (s UpdateCertificateForTerraformRequest) String() string {
  return tea.Prettify(s)
}

func (s UpdateCertificateForTerraformRequest) GoString() string {
  return s.String()
}

func (s *UpdateCertificateForTerraformRequest) SetName(v string) *UpdateCertificateForTerraformRequest {
  s.Name = &v
  return s
}

func (s *UpdateCertificateForTerraformRequest) SetCertificate(v string) *UpdateCertificateForTerraformRequest {
  s.Certificate = &v
  return s
}

func (s *UpdateCertificateForTerraformRequest) SetPrivateKey(v string) *UpdateCertificateForTerraformRequest {
  s.PrivateKey = &v
  return s
}

type UpdateCertificateForTerraformRequestHeader struct {
}

func (s UpdateCertificateForTerraformRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s UpdateCertificateForTerraformRequestHeader) GoString() string {
  return s.String()
}

type UpdateCertificateForTerraformPaths struct {
  // {"en":"The certificate ID you want to modify.","zh_CN":"需要修改的证书id"}
  CertificateId *int `json:"certificateId,omitempty" xml:"certificateId,omitempty" require:"true"`
}

func (s UpdateCertificateForTerraformPaths) String() string {
  return tea.Prettify(s)
}

func (s UpdateCertificateForTerraformPaths) GoString() string {
  return s.String()
}

func (s *UpdateCertificateForTerraformPaths) SetCertificateId(v int) *UpdateCertificateForTerraformPaths {
  s.CertificateId = &v
  return s
}

type UpdateCertificateForTerraformParameters struct {
}

func (s UpdateCertificateForTerraformParameters) String() string {
  return tea.Prettify(s)
}

func (s UpdateCertificateForTerraformParameters) GoString() string {
  return s.String()
}

type UpdateCertificateForTerraformResponse struct {
  // {"en":"Response code, 0 means successful.","zh_CN":"接口响应code，0代表成功。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response error message if failed.","zh_CN":"接口响应信息，success代表成功，失败则提供失败信息。"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Response data array.","zh_CN":"接口响应数据"}
  Data *string `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s UpdateCertificateForTerraformResponse) String() string {
  return tea.Prettify(s)
}

func (s UpdateCertificateForTerraformResponse) GoString() string {
  return s.String()
}

func (s *UpdateCertificateForTerraformResponse) SetCode(v string) *UpdateCertificateForTerraformResponse {
  s.Code = &v
  return s
}

func (s *UpdateCertificateForTerraformResponse) SetMessage(v string) *UpdateCertificateForTerraformResponse {
  s.Message = &v
  return s
}

func (s *UpdateCertificateForTerraformResponse) SetData(v string) *UpdateCertificateForTerraformResponse {
  s.Data = &v
  return s
}

type UpdateCertificateForTerraformResponseHeader struct {
}

func (s UpdateCertificateForTerraformResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s UpdateCertificateForTerraformResponseHeader) GoString() string {
  return s.String()
}




type AddCertificateForTerraformRequest struct {
  // {"en":"Certificate name","zh_CN":"证书名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"Certificate, PEM certificate, including CRT file and CA file.","zh_CN":"证书内容，PEM格式，包含CRT文件、CA文件。例如：\n-----BEGIN CERTIFICATE-----\n……\n-----END CERTIFICATE-----"}
  Certificate *string `json:"certificate,omitempty" xml:"certificate,omitempty" require:"true"`
  // {"en":"Private key of the certificate, PEM certificate. Not required when you specify a csrId.","zh_CN":"证书密钥，PEM格式。例如：\n-----BEGIN RSA PRIVATE KEY-----\n……\n-----BEGIN RSA PRIVATE KEY-----\n当指定csrId时，无需上传证书密钥。"}
  PrivateKey *string `json:"privateKey,omitempty" xml:"privateKey,omitempty" require:"true"`
}

func (s AddCertificateForTerraformRequest) String() string {
  return tea.Prettify(s)
}

func (s AddCertificateForTerraformRequest) GoString() string {
  return s.String()
}

func (s *AddCertificateForTerraformRequest) SetName(v string) *AddCertificateForTerraformRequest {
  s.Name = &v
  return s
}

func (s *AddCertificateForTerraformRequest) SetCertificate(v string) *AddCertificateForTerraformRequest {
  s.Certificate = &v
  return s
}

func (s *AddCertificateForTerraformRequest) SetPrivateKey(v string) *AddCertificateForTerraformRequest {
  s.PrivateKey = &v
  return s
}

type AddCertificateForTerraformRequestHeader struct {
}

func (s AddCertificateForTerraformRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s AddCertificateForTerraformRequestHeader) GoString() string {
  return s.String()
}

type AddCertificateForTerraformPaths struct {
}

func (s AddCertificateForTerraformPaths) String() string {
  return tea.Prettify(s)
}

func (s AddCertificateForTerraformPaths) GoString() string {
  return s.String()
}

type AddCertificateForTerraformParameters struct {
}

func (s AddCertificateForTerraformParameters) String() string {
  return tea.Prettify(s)
}

func (s AddCertificateForTerraformParameters) GoString() string {
  return s.String()
}

type AddCertificateForTerraformResponse struct {
  // {"en":"Status code","zh_CN":"状态码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response message","zh_CN":"响应信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Response data","zh_CN":"响应数据"}
  Data *AddCertificateForTerraformResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
}

func (s AddCertificateForTerraformResponse) String() string {
  return tea.Prettify(s)
}

func (s AddCertificateForTerraformResponse) GoString() string {
  return s.String()
}

func (s *AddCertificateForTerraformResponse) SetCode(v string) *AddCertificateForTerraformResponse {
  s.Code = &v
  return s
}

func (s *AddCertificateForTerraformResponse) SetMessage(v string) *AddCertificateForTerraformResponse {
  s.Message = &v
  return s
}

func (s *AddCertificateForTerraformResponse) SetData(v *AddCertificateForTerraformResponseData) *AddCertificateForTerraformResponse {
  s.Data = v
  return s
}

type AddCertificateForTerraformResponseData struct {
  // {"en":"Certificate ID","zh_CN":"证书ID"}
  CertificateId *int `json:"certificateId,omitempty" xml:"certificateId,omitempty" require:"true"`
}

func (s AddCertificateForTerraformResponseData) String() string {
  return tea.Prettify(s)
}

func (s AddCertificateForTerraformResponseData) GoString() string {
  return s.String()
}

func (s *AddCertificateForTerraformResponseData) SetCertificateId(v int) *AddCertificateForTerraformResponseData {
  s.CertificateId = &v
  return s
}

type AddCertificateForTerraformResponseHeader struct {
}

func (s AddCertificateForTerraformResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s AddCertificateForTerraformResponseHeader) GoString() string {
  return s.String()
}




type QueryCertificateForTerraformRequest struct {
}

func (s QueryCertificateForTerraformRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryCertificateForTerraformRequest) GoString() string {
  return s.String()
}

type QueryCertificateForTerraformRequestHeader struct {
}

func (s QueryCertificateForTerraformRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryCertificateForTerraformRequestHeader) GoString() string {
  return s.String()
}

type QueryCertificateForTerraformPaths struct {
  // {"en":"The certificate ID you want to query.","zh_CN":"需要查询的证书id"}
  CertificateId *int `json:"certificateId,omitempty" xml:"certificateId,omitempty" require:"true"`
}

func (s QueryCertificateForTerraformPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryCertificateForTerraformPaths) GoString() string {
  return s.String()
}

func (s *QueryCertificateForTerraformPaths) SetCertificateId(v int) *QueryCertificateForTerraformPaths {
  s.CertificateId = &v
  return s
}

type QueryCertificateForTerraformParameters struct {
}

func (s QueryCertificateForTerraformParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryCertificateForTerraformParameters) GoString() string {
  return s.String()
}

type QueryCertificateForTerraformResponse struct {
  // {"en":"Response code, 0 means successful.","zh_CN":"接口响应code，0代表成功。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response error message if failed.","zh_CN":"接口响应信息，success代表成功，失败则提供失败信息。"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Response data array.","zh_CN":"接口响应数据"}
  Data *QueryCertificateForTerraformResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
}

func (s QueryCertificateForTerraformResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryCertificateForTerraformResponse) GoString() string {
  return s.String()
}

func (s *QueryCertificateForTerraformResponse) SetCode(v string) *QueryCertificateForTerraformResponse {
  s.Code = &v
  return s
}

func (s *QueryCertificateForTerraformResponse) SetMessage(v string) *QueryCertificateForTerraformResponse {
  s.Message = &v
  return s
}

func (s *QueryCertificateForTerraformResponse) SetData(v *QueryCertificateForTerraformResponseData) *QueryCertificateForTerraformResponse {
  s.Data = v
  return s
}

type QueryCertificateForTerraformResponseData struct {
  // {"en":"certificate Id","zh_CN":"证书ID。"}
  CertificateId *int `json:"certificateId,omitempty" xml:"certificateId,omitempty" require:"true"`
  // {"en":"certificate name","zh_CN":"证书名称。"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"Certificate, PEM certificate.","zh_CN":"证书内容，PEM格式。"}
  Certificate *string `json:"certificate,omitempty" xml:"certificate,omitempty" require:"true"`
  // {"en":"Private key of the certificate, PEM certificate. Not required when you specify a csrId.","zh_CN":"证书密钥，PEM格式。"}
  PrivateKey *string `json:"privateKey,omitempty" xml:"privateKey,omitempty" require:"true"`
  // {"en":"Certificate related domains","zh_CN":"在用该证书的加速域名列表"}
  Domains []*QueryCertificateForTerraformResponseDataDomains `json:"domains,omitempty" xml:"domains,omitempty" require:"true" type:"Repeated"`
}

func (s QueryCertificateForTerraformResponseData) String() string {
  return tea.Prettify(s)
}

func (s QueryCertificateForTerraformResponseData) GoString() string {
  return s.String()
}

func (s *QueryCertificateForTerraformResponseData) SetCertificateId(v int) *QueryCertificateForTerraformResponseData {
  s.CertificateId = &v
  return s
}

func (s *QueryCertificateForTerraformResponseData) SetName(v string) *QueryCertificateForTerraformResponseData {
  s.Name = &v
  return s
}

func (s *QueryCertificateForTerraformResponseData) SetCertificate(v string) *QueryCertificateForTerraformResponseData {
  s.Certificate = &v
  return s
}

func (s *QueryCertificateForTerraformResponseData) SetPrivateKey(v string) *QueryCertificateForTerraformResponseData {
  s.PrivateKey = &v
  return s
}

func (s *QueryCertificateForTerraformResponseData) SetDomains(v []*QueryCertificateForTerraformResponseDataDomains) *QueryCertificateForTerraformResponseData {
  s.Domains = v
  return s
}

type QueryCertificateForTerraformResponseDataDomains struct     {
  // {"en":"Domain ID","zh_CN":"域名id"}
  DomainId *int64 `json:"domainId,omitempty" xml:"domainId,omitempty" require:"true"`
  // {"en":"Domain name","zh_CN":"域名名称"}
  DomainName *string `json:"domainName,omitempty" xml:"domainName,omitempty" require:"true"`
}

func (s QueryCertificateForTerraformResponseDataDomains) String() string {
  return tea.Prettify(s)
}

func (s QueryCertificateForTerraformResponseDataDomains) GoString() string {
  return s.String()
}

func (s *QueryCertificateForTerraformResponseDataDomains) SetDomainId(v int64) *QueryCertificateForTerraformResponseDataDomains {
  s.DomainId = &v
  return s
}

func (s *QueryCertificateForTerraformResponseDataDomains) SetDomainName(v string) *QueryCertificateForTerraformResponseDataDomains {
  s.DomainName = &v
  return s
}

type QueryCertificateForTerraformResponseHeader struct {
}

func (s QueryCertificateForTerraformResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryCertificateForTerraformResponseHeader) GoString() string {
  return s.String()
}




type QueryCertificateListForTerraformRequest struct {
  // {"en":"Response code, 0 means successful.","zh_CN":"接口响应code，0代表成功。"}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s QueryCertificateListForTerraformRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryCertificateListForTerraformRequest) GoString() string {
  return s.String()
}

func (s *QueryCertificateListForTerraformRequest) SetName(v string) *QueryCertificateListForTerraformRequest {
  s.Name = &v
  return s
}

type QueryCertificateListForTerraformRequestHeader struct {
}

func (s QueryCertificateListForTerraformRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryCertificateListForTerraformRequestHeader) GoString() string {
  return s.String()
}

type QueryCertificateListForTerraformPaths struct {
}

func (s QueryCertificateListForTerraformPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryCertificateListForTerraformPaths) GoString() string {
  return s.String()
}

type QueryCertificateListForTerraformParameters struct {
}

func (s QueryCertificateListForTerraformParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryCertificateListForTerraformParameters) GoString() string {
  return s.String()
}

type QueryCertificateListForTerraformResponse struct {
  // {"en":"Response code, 0 means successful.","zh_CN":"接口响应code，0代表成功。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response error message if failed.","zh_CN":"接口响应信息，success代表成功，失败则提供失败信息。"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Response data array.","zh_CN":"接口响应数据"}
  Data []*QueryCertificateListForTerraformResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s QueryCertificateListForTerraformResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryCertificateListForTerraformResponse) GoString() string {
  return s.String()
}

func (s *QueryCertificateListForTerraformResponse) SetCode(v string) *QueryCertificateListForTerraformResponse {
  s.Code = &v
  return s
}

func (s *QueryCertificateListForTerraformResponse) SetMessage(v string) *QueryCertificateListForTerraformResponse {
  s.Message = &v
  return s
}

func (s *QueryCertificateListForTerraformResponse) SetData(v []*QueryCertificateListForTerraformResponseData) *QueryCertificateListForTerraformResponse {
  s.Data = v
  return s
}

type QueryCertificateListForTerraformResponseData struct     {
  // {"en":"Certificate ID","zh_CN":"证书ID"}
  CertificateId *int64 `json:"certificateId,omitempty" xml:"certificateId,omitempty" require:"true"`
  // {"en":"Certificate name, unique to customer granularity","zh_CN":"证书名称，客户粒度下是唯一的"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"Remarks on cerfiticate file","zh_CN":"证书文件的备注"}
  Comment *string `json:"comment,omitempty" xml:"comment,omitempty" require:"true"`
  // {"en":"Shared, optional values are true and false, true represents shared certificates, false represents unshared certificates, default is false\nThis certificate allows cross-customer use when share-ssl is true. (The API does not support cross-customer use certificates. Contact customer service for manual configuration if required.)","zh_CN":"是否共享，true表示共享证书，false表示非共享证书"}
  ShareSsl *string `json:"shareSsl,omitempty" xml:"shareSsl,omitempty" require:"true"`
  // {"en":"Certificate effective start time (CST), such as 2016-08-01 07:00:00","zh_CN":"证书有效期的起始时间（CST时区），例如：2016-08-01 07:00:00"}
  CertificateValidityFrom *string `json:"certificateValidityFrom,omitempty" xml:"certificateValidityFrom,omitempty" require:"true"`
  // {"en":"Certificate effective end time (CST), such as 2018-08-01 19:00:00","zh_CN":"证书有效期的到期时间（CST时区），例如：2018-08-01 19:00:00"}
  CertificateValidityTo *string `json:"certificateValidityTo,omitempty" xml:"certificateValidityTo,omitempty" require:"true"`
  // {"en":"List of domain names using the current certificate","zh_CN":"使用当前证书的域名列表"}
  RelatedDomains []*QueryCertificateListForTerraformResponseDataRelatedDomains `json:"relatedDomains,omitempty" xml:"relatedDomains,omitempty" require:"true" type:"Repeated"`
  // {"en":"dns-names","zh_CN":"授权域名列表，证书使用者可选名称，父标签"}
  DnsNames []*string `json:"dnsNames,omitempty" xml:"dnsNames,omitempty" require:"true" type:"Repeated"`
  // {"en":"The CRT certificate serial number","zh_CN":"crt证书序列号"}
  CertificateSerial *string `json:"certificateSerial,omitempty" xml:"certificateSerial,omitempty" require:"true"`
  // {"en":"The CRT certificate issuer","zh_CN":"crt证书颁布者"}
  CertificateIssuer *string `json:"certificateIssuer,omitempty" xml:"certificateIssuer,omitempty" require:"true"`
}

func (s QueryCertificateListForTerraformResponseData) String() string {
  return tea.Prettify(s)
}

func (s QueryCertificateListForTerraformResponseData) GoString() string {
  return s.String()
}

func (s *QueryCertificateListForTerraformResponseData) SetCertificateId(v int64) *QueryCertificateListForTerraformResponseData {
  s.CertificateId = &v
  return s
}

func (s *QueryCertificateListForTerraformResponseData) SetName(v string) *QueryCertificateListForTerraformResponseData {
  s.Name = &v
  return s
}

func (s *QueryCertificateListForTerraformResponseData) SetComment(v string) *QueryCertificateListForTerraformResponseData {
  s.Comment = &v
  return s
}

func (s *QueryCertificateListForTerraformResponseData) SetShareSsl(v string) *QueryCertificateListForTerraformResponseData {
  s.ShareSsl = &v
  return s
}

func (s *QueryCertificateListForTerraformResponseData) SetCertificateValidityFrom(v string) *QueryCertificateListForTerraformResponseData {
  s.CertificateValidityFrom = &v
  return s
}

func (s *QueryCertificateListForTerraformResponseData) SetCertificateValidityTo(v string) *QueryCertificateListForTerraformResponseData {
  s.CertificateValidityTo = &v
  return s
}

func (s *QueryCertificateListForTerraformResponseData) SetRelatedDomains(v []*QueryCertificateListForTerraformResponseDataRelatedDomains) *QueryCertificateListForTerraformResponseData {
  s.RelatedDomains = v
  return s
}

func (s *QueryCertificateListForTerraformResponseData) SetDnsNames(v []*string) *QueryCertificateListForTerraformResponseData {
  s.DnsNames = v
  return s
}

func (s *QueryCertificateListForTerraformResponseData) SetCertificateSerial(v string) *QueryCertificateListForTerraformResponseData {
  s.CertificateSerial = &v
  return s
}

func (s *QueryCertificateListForTerraformResponseData) SetCertificateIssuer(v string) *QueryCertificateListForTerraformResponseData {
  s.CertificateIssuer = &v
  return s
}

type QueryCertificateListForTerraformResponseDataRelatedDomains struct     {
  // {"en":"Accelerated domain name ID","zh_CN":"加速域名ID"}
  DomainId *int64 `json:"domainId,omitempty" xml:"domainId,omitempty" require:"true"`
  // {"en":"Name of accelerated domain name","zh_CN":"加速域名的名称"}
  DomainName *string `json:"domainName,omitempty" xml:"domainName,omitempty" require:"true"`
}

func (s QueryCertificateListForTerraformResponseDataRelatedDomains) String() string {
  return tea.Prettify(s)
}

func (s QueryCertificateListForTerraformResponseDataRelatedDomains) GoString() string {
  return s.String()
}

func (s *QueryCertificateListForTerraformResponseDataRelatedDomains) SetDomainId(v int64) *QueryCertificateListForTerraformResponseDataRelatedDomains {
  s.DomainId = &v
  return s
}

func (s *QueryCertificateListForTerraformResponseDataRelatedDomains) SetDomainName(v string) *QueryCertificateListForTerraformResponseDataRelatedDomains {
  s.DomainName = &v
  return s
}

type QueryCertificateListForTerraformResponseHeader struct {
}

func (s QueryCertificateListForTerraformResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryCertificateListForTerraformResponseHeader) GoString() string {
  return s.String()
}




type DeleteCertificateForTerraformRequest struct {
}

func (s DeleteCertificateForTerraformRequest) String() string {
  return tea.Prettify(s)
}

func (s DeleteCertificateForTerraformRequest) GoString() string {
  return s.String()
}

type DeleteCertificateForTerraformRequestHeader struct {
}

func (s DeleteCertificateForTerraformRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteCertificateForTerraformRequestHeader) GoString() string {
  return s.String()
}

type DeleteCertificateForTerraformPaths struct {
  // {"en":"The certificate ID you want to delete.","zh_CN":"需要删除的证书id"}
  CertificateId *int `json:"certificateId,omitempty" xml:"certificateId,omitempty" require:"true"`
}

func (s DeleteCertificateForTerraformPaths) String() string {
  return tea.Prettify(s)
}

func (s DeleteCertificateForTerraformPaths) GoString() string {
  return s.String()
}

func (s *DeleteCertificateForTerraformPaths) SetCertificateId(v int) *DeleteCertificateForTerraformPaths {
  s.CertificateId = &v
  return s
}

type DeleteCertificateForTerraformParameters struct {
}

func (s DeleteCertificateForTerraformParameters) String() string {
  return tea.Prettify(s)
}

func (s DeleteCertificateForTerraformParameters) GoString() string {
  return s.String()
}

type DeleteCertificateForTerraformResponse struct {
  // {"en":"Response code, 0 means successful.","zh_CN":"接口响应code，0代表成功。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response error message if failed.","zh_CN":"接口响应信息，success代表成功，失败则提供失败信息。"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Response data array.","zh_CN":"接口响应数据"}
  Data *string `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s DeleteCertificateForTerraformResponse) String() string {
  return tea.Prettify(s)
}

func (s DeleteCertificateForTerraformResponse) GoString() string {
  return s.String()
}

func (s *DeleteCertificateForTerraformResponse) SetCode(v string) *DeleteCertificateForTerraformResponse {
  s.Code = &v
  return s
}

func (s *DeleteCertificateForTerraformResponse) SetMessage(v string) *DeleteCertificateForTerraformResponse {
  s.Message = &v
  return s
}

func (s *DeleteCertificateForTerraformResponse) SetData(v string) *DeleteCertificateForTerraformResponse {
  s.Data = &v
  return s
}

type DeleteCertificateForTerraformResponseHeader struct {
}

func (s DeleteCertificateForTerraformResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteCertificateForTerraformResponseHeader) GoString() string {
  return s.String()
}




