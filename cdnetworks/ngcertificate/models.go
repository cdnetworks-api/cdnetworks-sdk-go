package ngcertificate

// This file is auto-generated, don't edit it. Thanks.
import (
  "github.com/alibabacloud-go/tea/tea"
)

type GetACertificatePaths struct {
  // {"en":"ID of the certificate.","zh_CN": "证书ID。"}
  CertificateId *string `json:"certificateId,omitempty" xml:"certificateId,omitempty" require:"true"`
}

func (s GetACertificatePaths) String() string {
  return tea.Prettify(s)
}

func (s GetACertificatePaths) GoString() string {
  return s.String()
}

func (s *GetACertificatePaths) SetCertificateId(v string) *GetACertificatePaths {
  s.CertificateId = &v
  return s
}

type GetACertificateParameters struct {
}

func (s GetACertificateParameters) String() string {
  return tea.Prettify(s)
}

func (s GetACertificateParameters) GoString() string {
  return s.String()
}

type GetACertificateRequestHeader struct {
}

func (s GetACertificateRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetACertificateRequestHeader) GoString() string {
  return s.String()
}

type GetACertificateRequest struct {
}

func (s GetACertificateRequest) String() string {
  return tea.Prettify(s)
}

func (s GetACertificateRequest) GoString() string {
  return s.String()
}

type GetACertificateResponseHeader struct {
}

func (s GetACertificateResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetACertificateResponseHeader) GoString() string {
  return s.String()
}

type GetACertificateResponse struct {
  // {"en" : "ID of the certificate", "zh_CN": "证书ID。"}
  CertificateId *string `json:"certificateId,omitempty" xml:"certificateId,omitempty" require:"true"`
  // {"en" : "Name of the certificate.", "zh_CN": "证书名称。"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en" : "A description of the certificate.", "zh_CN": "证书描述。"}
  Description *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
  // {"en" : "Describes the versions of the certificate that have been created. You can obtain further details about each version by calling the Query a certificate version's details API.", "zh_CN": "证书版本列表。您可以通过调用'查询证书版本信息'接口来获取每个版本的更多信息。"}
  Versions []*GetACertificateResponseVersions `json:"versions,omitempty" xml:"versions,omitempty" require:"true" type:"Repeated"`
  // {"en" : "Enum: Off,LE 
  // A value of 'LE' indicates that auto renewal via Let's Encrypt (https://letsencrypt.org/docs/challenge-types/) is enabled.", "zh_CN": "取值范围: Off,LE 
  // 是否自动更新。'LE'值表示开启 Let's Encrypt 自动更新。"}
  AutoRenew *string `json:"autoRenew,omitempty" xml:"autoRenew,omitempty" require:"true"`
  // {"en" : "Indicates who is using the certificate in production.", "zh_CN": "证书在生产环境中的使用情况。"}
  UsageInProduction []*GetACertificateResponseUsageInProduction `json:"usageInProduction,omitempty" xml:"usageInProduction,omitempty" require:"true" type:"Repeated"`
  // {"en" : "Range: >= 1 
  // Indicates the version of the certificate deployed to staging.", "zh_CN": "取值范围: >= 1 
  // 表示部署到演练环境的证书版本。"}
  VersionInStaging *int `json:"versionInStaging,omitempty" xml:"versionInStaging,omitempty" require:"true"`
  // {"en" : "Indicates the customers using the certificate in staging.", "zh_CN": "证书在演练环境中的使用情况。"}
  UsageInStaging []*GetACertificateResponseUsageInStaging `json:"usageInStaging,omitempty" xml:"usageInStaging,omitempty" require:"true" type:"Repeated"`
  // {"en" : "Range: >= 1 
  // Indicates the version of the certificate deployed to production.", "zh_CN": "取值范围: >= 1 
  // 表示部署到生产环境的证书版本。"}
  VersionInProduction *int `json:"versionInProduction,omitempty" xml:"versionInProduction,omitempty" require:"true"`
  // {"en" : "A value of true requests the certificate to be auto-renewed as soon as possible instead of waiting for the certificate to expire in 15 days. The value will be set to false after a successful renewal.", "zh_CN": "是否强制更新。当值为true时表示要求尽快自动更新证书，而不是等待证书在 15 天后过期。 证书成功更新后，该值将设置为false。"}
  ForceRenew *bool `json:"forceRenew,omitempty" xml:"forceRenew,omitempty" require:"true"`
}

func (s GetACertificateResponse) String() string {
  return tea.Prettify(s)
}

func (s GetACertificateResponse) GoString() string {
  return s.String()
}

func (s *GetACertificateResponse) SetCertificateId(v string) *GetACertificateResponse {
  s.CertificateId = &v
  return s
}

func (s *GetACertificateResponse) SetName(v string) *GetACertificateResponse {
  s.Name = &v
  return s
}

func (s *GetACertificateResponse) SetDescription(v string) *GetACertificateResponse {
  s.Description = &v
  return s
}

func (s *GetACertificateResponse) SetVersions(v []*GetACertificateResponseVersions) *GetACertificateResponse {
  s.Versions = v
  return s
}

func (s *GetACertificateResponse) SetAutoRenew(v string) *GetACertificateResponse {
  s.AutoRenew = &v
  return s
}

func (s *GetACertificateResponse) SetUsageInProduction(v []*GetACertificateResponseUsageInProduction) *GetACertificateResponse {
  s.UsageInProduction = v
  return s
}

func (s *GetACertificateResponse) SetVersionInStaging(v int) *GetACertificateResponse {
  s.VersionInStaging = &v
  return s
}

func (s *GetACertificateResponse) SetUsageInStaging(v []*GetACertificateResponseUsageInStaging) *GetACertificateResponse {
  s.UsageInStaging = v
  return s
}

func (s *GetACertificateResponse) SetVersionInProduction(v int) *GetACertificateResponse {
  s.VersionInProduction = &v
  return s
}

func (s *GetACertificateResponse) SetForceRenew(v bool) *GetACertificateResponse {
  s.ForceRenew = &v
  return s
}

type GetACertificateResponseVersions struct     {
  // {"en" : "Range: >= 1 
  // Indicates the version number.", "zh_CN": "取值范围: >= 1 
  // 证书的版本号。"}
  Version *int `json:"version,omitempty" xml:"version,omitempty"`
  // {"en" : "Enum: uploaded,selfsigned 
  // Indicates the type of certificate, either 'uploaded' or 'self-signed'.", "zh_CN": "取值范围: uploaded,selfsigned 
  // 证书类型，包括'上传'和'自签名'2种类型。"}
  Type *string `json:"type,omitempty" xml:"type,omitempty"`
  // {"en" : "RFC 3339 format string indicating when the certificate expires.", "zh_CN": "RFC3339格式的日期，表示证书的过期时间。"}
  ExpirationTime *string `json:"expirationTime,omitempty" xml:"expirationTime,omitempty"`
  // {"en" : "RFC 3339 format string indicating when the certificate was created.", "zh_CN": "RFC3339格式的日期，表示证书的创建时间。"}
  CreationTime *string `json:"creationTime,omitempty" xml:"creationTime,omitempty"`
  // {"en" : "A unique fingerprint associated with the certificate.
  // ", "zh_CN": "与证书关联的唯一指纹。"}
  Fingerprint *string `json:"fingerprint,omitempty" xml:"fingerprint,omitempty"`
  // {"en" : "Comments about the certificate version.", "zh_CN": "证书版本的描述。"}
  Comments *string `json:"comments,omitempty" xml:"comments,omitempty"`
}

func (s GetACertificateResponseVersions) String() string {
  return tea.Prettify(s)
}

func (s GetACertificateResponseVersions) GoString() string {
  return s.String()
}

func (s *GetACertificateResponseVersions) SetVersion(v int) *GetACertificateResponseVersions {
  s.Version = &v
  return s
}

func (s *GetACertificateResponseVersions) SetType(v string) *GetACertificateResponseVersions {
  s.Type = &v
  return s
}

func (s *GetACertificateResponseVersions) SetExpirationTime(v string) *GetACertificateResponseVersions {
  s.ExpirationTime = &v
  return s
}

func (s *GetACertificateResponseVersions) SetCreationTime(v string) *GetACertificateResponseVersions {
  s.CreationTime = &v
  return s
}

func (s *GetACertificateResponseVersions) SetFingerprint(v string) *GetACertificateResponseVersions {
  s.Fingerprint = &v
  return s
}

func (s *GetACertificateResponseVersions) SetComments(v string) *GetACertificateResponseVersions {
  s.Comments = &v
  return s
}

type GetACertificateResponseUsageInProduction struct     {
  // {"en" : "ID of the property using the certificate", "zh_CN": "使用该证书的加速项目ID。"}
  PropertyId *string `json:"propertyId,omitempty" xml:"propertyId,omitempty"`
  // {"en" : "List of hostnames using the certificate.", "zh_CN": "使用该证书的加速域名列表。"}
  Hostnames []*string `json:"hostnames,omitempty" xml:"hostnames,omitempty" type:"Repeated"`
  // {"en" : "Names of the property's origins that use the certificate.", "zh_CN": "使用该证书的加速项目源站的名称。"}
  Origins []*string `json:"origins,omitempty" xml:"origins,omitempty" type:"Repeated"`
}

func (s GetACertificateResponseUsageInProduction) String() string {
  return tea.Prettify(s)
}

func (s GetACertificateResponseUsageInProduction) GoString() string {
  return s.String()
}

func (s *GetACertificateResponseUsageInProduction) SetPropertyId(v string) *GetACertificateResponseUsageInProduction {
  s.PropertyId = &v
  return s
}

func (s *GetACertificateResponseUsageInProduction) SetHostnames(v []*string) *GetACertificateResponseUsageInProduction {
  s.Hostnames = v
  return s
}

func (s *GetACertificateResponseUsageInProduction) SetOrigins(v []*string) *GetACertificateResponseUsageInProduction {
  s.Origins = v
  return s
}

type GetACertificateResponseUsageInStaging struct     {
  // {"en" : "ID of the property using the certificate.", "zh_CN": "使用该证书的加速项目的ID。"}
  PropertyId *string `json:"propertyId,omitempty" xml:"propertyId,omitempty"`
  // {"en" : "List of hostnames using the certificate.", "zh_CN": "使用该证书的加速域名列表。"}
  Hostnames []*string `json:"hostnames,omitempty" xml:"hostnames,omitempty" type:"Repeated"`
  // {"en" : "Names of the property's origins that use the certificate.", "zh_CN": "使用该证书的加速项目源站的名称。"}
  Origins []*string `json:"origins,omitempty" xml:"origins,omitempty" type:"Repeated"`
}

func (s GetACertificateResponseUsageInStaging) String() string {
  return tea.Prettify(s)
}

func (s GetACertificateResponseUsageInStaging) GoString() string {
  return s.String()
}

func (s *GetACertificateResponseUsageInStaging) SetPropertyId(v string) *GetACertificateResponseUsageInStaging {
  s.PropertyId = &v
  return s
}

func (s *GetACertificateResponseUsageInStaging) SetHostnames(v []*string) *GetACertificateResponseUsageInStaging {
  s.Hostnames = v
  return s
}

func (s *GetACertificateResponseUsageInStaging) SetOrigins(v []*string) *GetACertificateResponseUsageInStaging {
  s.Origins = v
  return s
}




type DownloadTheCsrPaths struct {
  // {"en" : "ID of the certificate.", "zh_CN": "证书的ID。"}
  CertificateId *string `json:"certificateId,omitempty" xml:"certificateId,omitempty" require:"true"`
}

func (s DownloadTheCsrPaths) String() string {
  return tea.Prettify(s)
}

func (s DownloadTheCsrPaths) GoString() string {
  return s.String()
}

func (s *DownloadTheCsrPaths) SetCertificateId(v string) *DownloadTheCsrPaths {
  s.CertificateId = &v
  return s
}

type DownloadTheCsrParameters struct {
  // {"en" : "Enum: sectigo 
  // This optional parameter can be used to request a domain control validation (DCV) file. Currently, we support the certificate authority, Sectigo, by specifying 'sectigo' as the value of the parameter.", "zh_CN": "取值范围: sectigo 
  // 可选参数，用于获取域控制验证(DCV)文件。目前，我们支持证书颁发机构Sectigo，您可将'sectigo'指定为参数值。"}
  Dcv *string `json:"dcv,omitempty" xml:"dcv,omitempty"`
}

func (s DownloadTheCsrParameters) String() string {
  return tea.Prettify(s)
}

func (s DownloadTheCsrParameters) GoString() string {
  return s.String()
}

func (s *DownloadTheCsrParameters) SetDcv(v string) *DownloadTheCsrParameters {
  s.Dcv = &v
  return s
}

type DownloadTheCsrRequestHeader struct {
}

func (s DownloadTheCsrRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s DownloadTheCsrRequestHeader) GoString() string {
  return s.String()
}

type DownloadTheCsrRequest struct {
}

func (s DownloadTheCsrRequest) String() string {
  return tea.Prettify(s)
}

func (s DownloadTheCsrRequest) GoString() string {
  return s.String()
}

type DownloadTheCsrResponseHeader struct {
}

func (s DownloadTheCsrResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s DownloadTheCsrResponseHeader) GoString() string {
  return s.String()
}

type DownloadTheCsrResponse struct {
  // {"en" : "ID of the certificate.", "zh_CN": "证书的ID。"}
  CertificateId *string `json:"certificateId,omitempty" xml:"certificateId,omitempty" require:"true"`
  // {"en" : "The CSR.", "zh_CN": "证书签名请求（CSR）文件。"}
  Csr *string `json:"csr,omitempty" xml:"csr,omitempty" require:"true"`
  // {"en" : "Range: >= 1 
  // Most recent version of the certifiate.", "zh_CN": "取值范围: >= 1 
  // 证书的最新版本。"}
  LatestVersion *int `json:"latestVersion,omitempty" xml:"latestVersion,omitempty" require:"true"`
  // {"en" : "This field is only returned if the dcv query parameter is specified.", "zh_CN": "仅当指定了dcv 查询参数时才会返回此字段。"}
  DcvFile *DownloadTheCsrResponseDcvFile `json:"dcvFile,omitempty" xml:"dcvFile,omitempty" require:"true" type:"Struct"`
}

func (s DownloadTheCsrResponse) String() string {
  return tea.Prettify(s)
}

func (s DownloadTheCsrResponse) GoString() string {
  return s.String()
}

func (s *DownloadTheCsrResponse) SetCertificateId(v string) *DownloadTheCsrResponse {
  s.CertificateId = &v
  return s
}

func (s *DownloadTheCsrResponse) SetCsr(v string) *DownloadTheCsrResponse {
  s.Csr = &v
  return s
}

func (s *DownloadTheCsrResponse) SetLatestVersion(v int) *DownloadTheCsrResponse {
  s.LatestVersion = &v
  return s
}

func (s *DownloadTheCsrResponse) SetDcvFile(v *DownloadTheCsrResponseDcvFile) *DownloadTheCsrResponse {
  s.DcvFile = v
  return s
}

type DownloadTheCsrResponseDcvFile struct {
  // {"en" : "A URI that is accessible on your hostnames using the certificate. The file's content will consist of a SHA-256 hash and the domain sectigo.com. Example:
  // 
  // <pre>
  // a6e96b23f9e2add3f79c2907ded1adc43961b05f6d702758a200a8ec8b4d7115
  // sectigo.com
  // </pre>", "zh_CN": "使用到该证书的加速域名可访问到的一个路径。该文件的内容由一个SHA-256哈希值和域名sectigo.com 组成。例如：
  // 
  // <pre>
  // a6e96b23f9e2add3f79c2907ded1adc43961b05f6d702758a200a8ec8b4d7115
  // sectigo.com
  // </pre>"}
  Uri *string `json:"uri,omitempty" xml:"uri,omitempty" require:"true"`
  // {"en" : "An RFC 3339 date indicating when the file expires.", "zh_CN": "RFC3339格式的日期，表示CSR文件的过期时间。"}
  ExpirationTime *string `json:"expirationTime,omitempty" xml:"expirationTime,omitempty" require:"true"`
}

func (s DownloadTheCsrResponseDcvFile) String() string {
  return tea.Prettify(s)
}

func (s DownloadTheCsrResponseDcvFile) GoString() string {
  return s.String()
}

func (s *DownloadTheCsrResponseDcvFile) SetUri(v string) *DownloadTheCsrResponseDcvFile {
  s.Uri = &v
  return s
}

func (s *DownloadTheCsrResponseDcvFile) SetExpirationTime(v string) *DownloadTheCsrResponseDcvFile {
  s.ExpirationTime = &v
  return s
}




type GetACertificateVersionsDetailsPaths struct {
  // {"en" : "ID of a certificate.", "zh_CN": "证书id。"}
  CertificateId *string `json:"certificateId,omitempty" xml:"certificateId,omitempty" require:"true"`
  // {"en" : "Version of the certificate. ", "zh_CN": "证书版本。"}
  Version *int `json:"version,omitempty" xml:"version,omitempty" require:"true"`
}

func (s GetACertificateVersionsDetailsPaths) String() string {
  return tea.Prettify(s)
}

func (s GetACertificateVersionsDetailsPaths) GoString() string {
  return s.String()
}

func (s *GetACertificateVersionsDetailsPaths) SetCertificateId(v string) *GetACertificateVersionsDetailsPaths {
  s.CertificateId = &v
  return s
}

func (s *GetACertificateVersionsDetailsPaths) SetVersion(v int) *GetACertificateVersionsDetailsPaths {
  s.Version = &v
  return s
}

type GetACertificateVersionsDetailsParameters struct {
}

func (s GetACertificateVersionsDetailsParameters) String() string {
  return tea.Prettify(s)
}

func (s GetACertificateVersionsDetailsParameters) GoString() string {
  return s.String()
}

type GetACertificateVersionsDetailsRequestHeader struct {
}

func (s GetACertificateVersionsDetailsRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetACertificateVersionsDetailsRequestHeader) GoString() string {
  return s.String()
}

type GetACertificateVersionsDetailsRequest struct {
}

func (s GetACertificateVersionsDetailsRequest) String() string {
  return tea.Prettify(s)
}

func (s GetACertificateVersionsDetailsRequest) GoString() string {
  return s.String()
}

type GetACertificateVersionsDetailsResponseHeader struct {
}

func (s GetACertificateVersionsDetailsResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetACertificateVersionsDetailsResponseHeader) GoString() string {
  return s.String()
}

type GetACertificateVersionsDetailsResponse struct {
  // {"en" : "Range: >= 1 
  // The certificate version.", "zh_CN": "取值范围: >= 1 
  // 证书版本。"}
  Version *int `json:"version,omitempty" xml:"version,omitempty" require:"true"`
  // {"en" : "Comments about the certificate version.
  // ", "zh_CN": "证书版本的描述。"}
  Comments *string `json:"comments,omitempty" xml:"comments,omitempty" require:"true"`
  // {"en" : "An RFC3339 format date indicating when the certificate version expires.", "zh_CN": "RFC3339格式的日期，表示证书版本的过期时间。"}
  ExpirationTime *string `json:"expirationTime,omitempty" xml:"expirationTime,omitempty" require:"true"`
  // {"en" : "An RFC 3339 format date indicating when the certificate version was created.", "zh_CN": "RFC3339格式的日期，用于表示证书版本创建的时间。"}
  CreationTime *string `json:"creationTime,omitempty" xml:"creationTime,omitempty" require:"true"`
  // {"en" : "The certificate subject.", "zh_CN": "证书主体。"}
  Subject *string `json:"subject,omitempty" xml:"subject,omitempty" require:"true"`
  // {"en" : "The signature algorithm used by the certificate.", "zh_CN": "证书使用的签名算法。"}
  SignatureAlgo *string `json:"signatureAlgo,omitempty" xml:"signatureAlgo,omitempty" require:"true"`
  // {"en" : "The serial number associated with the certificate.
  // ", "zh_CN": "与证书相关联的序列号。"}
  SerialNumber *string `json:"serialNumber,omitempty" xml:"serialNumber,omitempty" require:"true"`
  // {"en" : "Indicates whether the certificate version is currently deployed to production.", "zh_CN": "表示证书版本当前是否已部署到生产环境。"}
  InProduction *bool `json:"inProduction,omitempty" xml:"inProduction,omitempty" require:"true"`
  // {"en" : "Indicates whether the certificate version is currently deployed to staging.
  // ", "zh_CN": "表示证书版本当前是否已部署到演练环境。"}
  InStaging *bool `json:"inStaging,omitempty" xml:"inStaging,omitempty" require:"true"`
  // {"en" : "certificate fingerprint.", "zh_CN": "证书指纹。"}
  Fingerprint *string `json:"fingerprint,omitempty" xml:"fingerprint,omitempty" require:"true"`
  // {"en" : "The encryption algorithm.
  // ", "zh_CN": "加密算法。"}
  Algorithm *string `json:"algorithm,omitempty" xml:"algorithm,omitempty" require:"true"`
  // {"en" : "Number of bits used in encryption.
  // ", "zh_CN": "加密算法使用的位数。
  // "}
  KeyLength *int `json:"keyLength,omitempty" xml:"keyLength,omitempty" require:"true"`
  // {"en" : "List of hostnames served by the certificate. Wildcards are permitted, for example, *.domain.com.
  // ", "zh_CN": "证书所涵盖的域名列表（SAN）。允许使用通配符，例如，*.domain.com。
  // "}
  SubjectAlternativeNames []*string `json:"subjectAlternativeNames,omitempty" xml:"subjectAlternativeNames,omitempty" require:"true" type:"Repeated"`
  // {"en" : "Describes the certificate chain.", "zh_CN": "链证书。"}
  ChainCertificates []*GetACertificateVersionsDetailsResponseChainCertificates `json:"chainCertificates,omitempty" xml:"chainCertificates,omitempty" require:"true" type:"Repeated"`
  // {"en" : "Issuer of the certificate.", "zh_CN": "证书的颁发者。"}
  Issuer *string `json:"issuer,omitempty" xml:"issuer,omitempty" require:"true"`
}

func (s GetACertificateVersionsDetailsResponse) String() string {
  return tea.Prettify(s)
}

func (s GetACertificateVersionsDetailsResponse) GoString() string {
  return s.String()
}

func (s *GetACertificateVersionsDetailsResponse) SetVersion(v int) *GetACertificateVersionsDetailsResponse {
  s.Version = &v
  return s
}

func (s *GetACertificateVersionsDetailsResponse) SetComments(v string) *GetACertificateVersionsDetailsResponse {
  s.Comments = &v
  return s
}

func (s *GetACertificateVersionsDetailsResponse) SetExpirationTime(v string) *GetACertificateVersionsDetailsResponse {
  s.ExpirationTime = &v
  return s
}

func (s *GetACertificateVersionsDetailsResponse) SetCreationTime(v string) *GetACertificateVersionsDetailsResponse {
  s.CreationTime = &v
  return s
}

func (s *GetACertificateVersionsDetailsResponse) SetSubject(v string) *GetACertificateVersionsDetailsResponse {
  s.Subject = &v
  return s
}

func (s *GetACertificateVersionsDetailsResponse) SetSignatureAlgo(v string) *GetACertificateVersionsDetailsResponse {
  s.SignatureAlgo = &v
  return s
}

func (s *GetACertificateVersionsDetailsResponse) SetSerialNumber(v string) *GetACertificateVersionsDetailsResponse {
  s.SerialNumber = &v
  return s
}

func (s *GetACertificateVersionsDetailsResponse) SetInProduction(v bool) *GetACertificateVersionsDetailsResponse {
  s.InProduction = &v
  return s
}

func (s *GetACertificateVersionsDetailsResponse) SetInStaging(v bool) *GetACertificateVersionsDetailsResponse {
  s.InStaging = &v
  return s
}

func (s *GetACertificateVersionsDetailsResponse) SetFingerprint(v string) *GetACertificateVersionsDetailsResponse {
  s.Fingerprint = &v
  return s
}

func (s *GetACertificateVersionsDetailsResponse) SetAlgorithm(v string) *GetACertificateVersionsDetailsResponse {
  s.Algorithm = &v
  return s
}

func (s *GetACertificateVersionsDetailsResponse) SetKeyLength(v int) *GetACertificateVersionsDetailsResponse {
  s.KeyLength = &v
  return s
}

func (s *GetACertificateVersionsDetailsResponse) SetSubjectAlternativeNames(v []*string) *GetACertificateVersionsDetailsResponse {
  s.SubjectAlternativeNames = v
  return s
}

func (s *GetACertificateVersionsDetailsResponse) SetChainCertificates(v []*GetACertificateVersionsDetailsResponseChainCertificates) *GetACertificateVersionsDetailsResponse {
  s.ChainCertificates = v
  return s
}

func (s *GetACertificateVersionsDetailsResponse) SetIssuer(v string) *GetACertificateVersionsDetailsResponse {
  s.Issuer = &v
  return s
}

type GetACertificateVersionsDetailsResponseChainCertificates struct     {
  // {"en" : "Subject of the certificate.", "zh_CN": "证书主体。"}
  Subject *string `json:"subject,omitempty" xml:"subject,omitempty"`
  // {"en" : "RFC 3339 format date indicating when the certificate expires.", "zh_CN": "RFC3339格式的日期，表示证书的过期时间。"}
  ExpirationTime *string `json:"expirationTime,omitempty" xml:"expirationTime,omitempty"`
  // {"en" : "Algorithm of the certificate.", "zh_CN": "证书的算法。"}
  SignatureAlgo *string `json:"signatureAlgo,omitempty" xml:"signatureAlgo,omitempty"`
  // {"en" : "Issuer of the chain certificate.", "zh_CN": "链证书的颁发者。"}
  Issuer *string `json:"issuer,omitempty" xml:"issuer,omitempty"`
}

func (s GetACertificateVersionsDetailsResponseChainCertificates) String() string {
  return tea.Prettify(s)
}

func (s GetACertificateVersionsDetailsResponseChainCertificates) GoString() string {
  return s.String()
}

func (s *GetACertificateVersionsDetailsResponseChainCertificates) SetSubject(v string) *GetACertificateVersionsDetailsResponseChainCertificates {
  s.Subject = &v
  return s
}

func (s *GetACertificateVersionsDetailsResponseChainCertificates) SetExpirationTime(v string) *GetACertificateVersionsDetailsResponseChainCertificates {
  s.ExpirationTime = &v
  return s
}

func (s *GetACertificateVersionsDetailsResponseChainCertificates) SetSignatureAlgo(v string) *GetACertificateVersionsDetailsResponseChainCertificates {
  s.SignatureAlgo = &v
  return s
}

func (s *GetACertificateVersionsDetailsResponseChainCertificates) SetIssuer(v string) *GetACertificateVersionsDetailsResponseChainCertificates {
  s.Issuer = &v
  return s
}




type DeleteACertificatePaths struct {
  // {"en" : "ID of the certificate.", "zh_CN": "证书 id。"}
  CertificateId *string `json:"certificateId,omitempty" xml:"certificateId,omitempty" require:"true"`
}

func (s DeleteACertificatePaths) String() string {
  return tea.Prettify(s)
}

func (s DeleteACertificatePaths) GoString() string {
  return s.String()
}

func (s *DeleteACertificatePaths) SetCertificateId(v string) *DeleteACertificatePaths {
  s.CertificateId = &v
  return s
}

type DeleteACertificateParameters struct {
}

func (s DeleteACertificateParameters) String() string {
  return tea.Prettify(s)
}

func (s DeleteACertificateParameters) GoString() string {
  return s.String()
}

type DeleteACertificateRequestHeader struct {
}

func (s DeleteACertificateRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteACertificateRequestHeader) GoString() string {
  return s.String()
}

type DeleteACertificateRequest struct {
}

func (s DeleteACertificateRequest) String() string {
  return tea.Prettify(s)
}

func (s DeleteACertificateRequest) GoString() string {
  return s.String()
}

type DeleteACertificateResponseHeader struct {
}

func (s DeleteACertificateResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteACertificateResponseHeader) GoString() string {
  return s.String()
}

type DeleteACertificateResponse struct {
}

func (s DeleteACertificateResponse) String() string {
  return tea.Prettify(s)
}

func (s DeleteACertificateResponse) GoString() string {
  return s.String()
}




type CreateACertificatePaths struct {
}

func (s CreateACertificatePaths) String() string {
  return tea.Prettify(s)
}

func (s CreateACertificatePaths) GoString() string {
  return s.String()
}

type CreateACertificateParameters struct {
}

func (s CreateACertificateParameters) String() string {
  return tea.Prettify(s)
}

func (s CreateACertificateParameters) GoString() string {
  return s.String()
}

type CreateACertificateRequestHeader struct {
}

func (s CreateACertificateRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s CreateACertificateRequestHeader) GoString() string {
  return s.String()
}

type CreateACertificateRequest struct {
  // {"en" : "Name of the certificate.", "zh_CN": "证书名称。"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en" : "A description of the new certificate.
  // ", "zh_CN": "证书描述。"}
  Description *string `json:"description,omitempty" xml:"description,omitempty"`
  // {"en" : "Enum: LE,Off 
  // Specifying 'LE' requests that we automatically renew your certificate through Let's Encrypt (https://letsencrypt.org/docs/challenge-types/) when it is close to expiring. ", "zh_CN": "取值范围: LE,Off 
  // 是否开启证书自动更新。当值为'LE'时，我们将会在证书即将到期时通过Let's Encrypt自动更新您的证书。"}
  AutoRenew *string `json:"autoRenew,omitempty" xml:"autoRenew,omitempty"`
  // {"en" : "This object is used to specify the initial version of the certificate.", "zh_CN": "证书的第一个版本。"}
  NewVersion *CreateACertificateRequestNewVersion `json:"newVersion,omitempty" xml:"newVersion,omitempty" type:"Struct"`
  // {"en" : "Default: False 
  // A value of true requests the certificate to be auto-renewed as soon as possible instead of waiting for the certificate to expire in 15 days. The value will be set to false after a successful renewal.", "zh_CN": "默认值: False 
  // 是否强制更新。当值为true时表示要求尽快自动更新证书，而不是等待证书在 15 天后过期。 
  // 证书成功更新后，该值将设置为false。"}
  ForceRenew *bool `json:"forceRenew,omitempty" xml:"forceRenew,omitempty"`
}

func (s CreateACertificateRequest) String() string {
  return tea.Prettify(s)
}

func (s CreateACertificateRequest) GoString() string {
  return s.String()
}

func (s *CreateACertificateRequest) SetName(v string) *CreateACertificateRequest {
  s.Name = &v
  return s
}

func (s *CreateACertificateRequest) SetDescription(v string) *CreateACertificateRequest {
  s.Description = &v
  return s
}

func (s *CreateACertificateRequest) SetAutoRenew(v string) *CreateACertificateRequest {
  s.AutoRenew = &v
  return s
}

func (s *CreateACertificateRequest) SetNewVersion(v *CreateACertificateRequestNewVersion) *CreateACertificateRequest {
  s.NewVersion = v
  return s
}

func (s *CreateACertificateRequest) SetForceRenew(v bool) *CreateACertificateRequest {
  s.ForceRenew = &v
  return s
}

type CreateACertificateRequestNewVersion struct {
  // {"en" : "Comments about the certificate version.", "zh_CN": "证书版本的描述。"}
  Comments *string `json:"comments,omitempty" xml:"comments,omitempty"`
  // {"en" : "The value must be either the private key in PEM format and encrypted with the API key and timestamp OR the literal string 'RSA2048' or 'ECC256' if you opt to generate a self-signed certificate. The key must be encrypted with AES-128-CBC and base64-encoded. This helps protect your key when you upload it to CDN Pro.", "zh_CN": "用于指定证书私钥，必须是PEM格式的私钥。如果您选择生成自签名证书，则此处的值应为'RSA2048'或'ECC256'。请使用您API账号的密钥和时间戳对私钥进行加密再上传。请使用AES-128-CBC加密算法，并用base64编码。当您将私钥上传到CDN Pro时，这种加密方式可以保护您的私钥。"}
  PrivateKey *string `json:"privateKey,omitempty" xml:"privateKey,omitempty"`
  // {"en" : "Enter the certificate in PEM format or the literal string 'GENERATE' if you wish that we create a self-signed certificate valid for ten days for test purposes. If you enter 'GENERATE', then you must also specify either 'RSA2048' or 'ECC256' as the value of the privateKey field to choose the encryption for the self-signed certificate.", "zh_CN": "用于指定证书（公钥），必须是PEM格式。如果您希望系统自动生成自签名证书（有效期 10 天）用于测试，则此处的值应为'GENERATE'。 当您指定'GENERATE'时，必须同时指定'RSA2048'或'ECC256'作为 privateKey 字段的值，用于选择加密算法生成自签名证书。"}
  Certificate *string `json:"certificate,omitempty" xml:"certificate,omitempty"`
  // {"en" : "The chain certificate in PEM format.", "zh_CN": "用于指定链证书。必须是PEM格式"}
  ChainCert *string `json:"chainCert,omitempty" xml:"chainCert,omitempty"`
  // {"en" : "Information submitted when generating a self-signed certificate.", "zh_CN": "当您选择生成自签名证书时，需提交以下信息。"}
  IdentificationInfo *CreateACertificateRequestNewVersionIdentificationInfo `json:"identificationInfo,omitempty" xml:"identificationInfo,omitempty" type:"Struct"`
}

func (s CreateACertificateRequestNewVersion) String() string {
  return tea.Prettify(s)
}

func (s CreateACertificateRequestNewVersion) GoString() string {
  return s.String()
}

func (s *CreateACertificateRequestNewVersion) SetComments(v string) *CreateACertificateRequestNewVersion {
  s.Comments = &v
  return s
}

func (s *CreateACertificateRequestNewVersion) SetPrivateKey(v string) *CreateACertificateRequestNewVersion {
  s.PrivateKey = &v
  return s
}

func (s *CreateACertificateRequestNewVersion) SetCertificate(v string) *CreateACertificateRequestNewVersion {
  s.Certificate = &v
  return s
}

func (s *CreateACertificateRequestNewVersion) SetChainCert(v string) *CreateACertificateRequestNewVersion {
  s.ChainCert = &v
  return s
}

func (s *CreateACertificateRequestNewVersion) SetIdentificationInfo(v *CreateACertificateRequestNewVersionIdentificationInfo) *CreateACertificateRequestNewVersion {
  s.IdentificationInfo = v
  return s
}

type CreateACertificateRequestNewVersionIdentificationInfo struct {
  // {"en" : "Range: [ 2 .. 2 ] characters 
  // An ISO-3166 country code.", "zh_CN": "取值范围: [ 2 .. 2 ] 字符 
  // ISO-3166国家代码。"}
  Country *string `json:"country,omitempty" xml:"country,omitempty"`
  // {"en" : "A state or province.", "zh_CN": "州或省。"}
  State *string `json:"state,omitempty" xml:"state,omitempty"`
  // {"en" : "A city.", "zh_CN": "城市。"}
  City *string `json:"city,omitempty" xml:"city,omitempty"`
  // {"en" : "A company name.", "zh_CN": "公司名称。"}
  Company *string `json:"company,omitempty" xml:"company,omitempty"`
  // {"en" : "The department associated with the certificate.", "zh_CN": "与证书关联的部门。"}
  Department *string `json:"department,omitempty" xml:"department,omitempty"`
  // {"en" : "A common name for the certificate.", "zh_CN": "证书的通用名称。"}
  CommonName *string `json:"commonName,omitempty" xml:"commonName,omitempty" require:"true"`
  // {"en" : "An email address.", "zh_CN": "邮箱地址。"}
  Email *string `json:"email,omitempty" xml:"email,omitempty"`
  // {"en" : "Hostnames that this certificate will serve. Each entry must be a valid hostname such as domain.com or a wildcard hostname beginning with '*' such as *.domain.com.", "zh_CN": "此证书涵盖的加速域名（SAN）列表。每个条目必须是有效的加速域名，例如 domain.com
  // 或者是以'*'开头的泛域名，例如：*.domain.com。"}
  SubjectAlternativeNames []*string `json:"subjectAlternativeNames,omitempty" xml:"subjectAlternativeNames,omitempty" require:"true" type:"Repeated"`
}

func (s CreateACertificateRequestNewVersionIdentificationInfo) String() string {
  return tea.Prettify(s)
}

func (s CreateACertificateRequestNewVersionIdentificationInfo) GoString() string {
  return s.String()
}

func (s *CreateACertificateRequestNewVersionIdentificationInfo) SetCountry(v string) *CreateACertificateRequestNewVersionIdentificationInfo {
  s.Country = &v
  return s
}

func (s *CreateACertificateRequestNewVersionIdentificationInfo) SetState(v string) *CreateACertificateRequestNewVersionIdentificationInfo {
  s.State = &v
  return s
}

func (s *CreateACertificateRequestNewVersionIdentificationInfo) SetCity(v string) *CreateACertificateRequestNewVersionIdentificationInfo {
  s.City = &v
  return s
}

func (s *CreateACertificateRequestNewVersionIdentificationInfo) SetCompany(v string) *CreateACertificateRequestNewVersionIdentificationInfo {
  s.Company = &v
  return s
}

func (s *CreateACertificateRequestNewVersionIdentificationInfo) SetDepartment(v string) *CreateACertificateRequestNewVersionIdentificationInfo {
  s.Department = &v
  return s
}

func (s *CreateACertificateRequestNewVersionIdentificationInfo) SetCommonName(v string) *CreateACertificateRequestNewVersionIdentificationInfo {
  s.CommonName = &v
  return s
}

func (s *CreateACertificateRequestNewVersionIdentificationInfo) SetEmail(v string) *CreateACertificateRequestNewVersionIdentificationInfo {
  s.Email = &v
  return s
}

func (s *CreateACertificateRequestNewVersionIdentificationInfo) SetSubjectAlternativeNames(v []*string) *CreateACertificateRequestNewVersionIdentificationInfo {
  s.SubjectAlternativeNames = v
  return s
}

type CreateACertificateResponse struct {
}

func (s CreateACertificateResponse) String() string {
  return tea.Prettify(s)
}

func (s CreateACertificateResponse) GoString() string {
  return s.String()
}

type CreateACertificateResponseHeader struct {
  // {"en":"The Location header's value will be a URL allowing you to get details about the new certificate.  Example: <code>Location: http://ngapi.cdnetworks.com/cdn/certificates/d60b730d9a586425677940cc</code>", "zh_CN":"通过Location响应头返回新建的证书的URL。URL中包含证书的ID，可使用该ID调用'查询证书详情'接口来查看证书详细信息。URL示例：<code>Location: http://open.chinanetcenter.com/cdn/certificates/5dca2205f9e9cc0001df7b33"}
  Location *string `json:"Location,omitempty" xml:"Location,omitempty" require:"true"`
}

func (s CreateACertificateResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s CreateACertificateResponseHeader) GoString() string {
  return s.String()
}

func (s *CreateACertificateResponseHeader) SetLocation(v string) *CreateACertificateResponseHeader {
  s.Location = &v
  return s
}




type UpdateACertificatePaths struct {
  // {"en" : "ID of the certificate.", "zh_CN": "证书 id。"}
  CertificateId *string `json:"certificateId,omitempty" xml:"certificateId,omitempty" require:"true"`
}

func (s UpdateACertificatePaths) String() string {
  return tea.Prettify(s)
}

func (s UpdateACertificatePaths) GoString() string {
  return s.String()
}

func (s *UpdateACertificatePaths) SetCertificateId(v string) *UpdateACertificatePaths {
  s.CertificateId = &v
  return s
}

type UpdateACertificateParameters struct {
}

func (s UpdateACertificateParameters) String() string {
  return tea.Prettify(s)
}

func (s UpdateACertificateParameters) GoString() string {
  return s.String()
}

type UpdateACertificateRequestHeader struct {
}

func (s UpdateACertificateRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s UpdateACertificateRequestHeader) GoString() string {
  return s.String()
}

type UpdateACertificateRequest struct {
  // {"en" : "Name of the certificate.", "zh_CN": "证书名称。"}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // {"en" : "Description of the certificate.", "zh_CN": "证书描述。"}
  Description *string `json:"description,omitempty" xml:"description,omitempty"`
  // {"en" : "Enum: Off,LE 
  // Indicates whether the certificate will be renewed with Let's Encrypt.", "zh_CN": "取值范围: Off,LE 
  // 是否开启证书自动更新。当值为'LE'时，我们将会在证书即将到期时通过Let's Encrypt自动更新您的证书。"}
  AutoRenew *string `json:"autoRenew,omitempty" xml:"autoRenew,omitempty"`
  // {"en" : "If this field is present, a new version of the certificate will be created. If the identificationInfo field is not provided, then the information will be copied from the latest version of the certificate.", "zh_CN": "如果该字段存在，则将创建一个新的证书版本。如果没有提供identiationinfo字段，则相关信息将从证书的最新版本中复制。"}
  NewVersion *UpdateACertificateRequestNewVersion `json:"newVersion,omitempty" xml:"newVersion,omitempty" type:"Struct"`
  // {"en" : "Default: False 
  // A value of true requests the certificate to be auto-renewed as soon as possible instead of waiting for the certificate to expire in 15 days. The value will be set to false after a successful renewal.", "zh_CN": "默认值: False 
  // 是否强制更新。当值为true时表示要求尽快自动更新证书，而不是等待证书在 15 天后过期。 证书成功更新后，该值将设置为false。"}
  ForceRenew *bool `json:"forceRenew,omitempty" xml:"forceRenew,omitempty"`
}

func (s UpdateACertificateRequest) String() string {
  return tea.Prettify(s)
}

func (s UpdateACertificateRequest) GoString() string {
  return s.String()
}

func (s *UpdateACertificateRequest) SetName(v string) *UpdateACertificateRequest {
  s.Name = &v
  return s
}

func (s *UpdateACertificateRequest) SetDescription(v string) *UpdateACertificateRequest {
  s.Description = &v
  return s
}

func (s *UpdateACertificateRequest) SetAutoRenew(v string) *UpdateACertificateRequest {
  s.AutoRenew = &v
  return s
}

func (s *UpdateACertificateRequest) SetNewVersion(v *UpdateACertificateRequestNewVersion) *UpdateACertificateRequest {
  s.NewVersion = v
  return s
}

func (s *UpdateACertificateRequest) SetForceRenew(v bool) *UpdateACertificateRequest {
  s.ForceRenew = &v
  return s
}

type UpdateACertificateRequestNewVersion struct {
  // {"en" : "If not present, the value will be copied from the latest version of the certificate. Please refer to the description of the privateKey field in the Create a certificate API for details about the format.", "zh_CN": "如果未指定该字段，则将从证书的最新版本中复制。 具体格式请参考'创建证书'接口中privateKey字段的说明。"}
  PrivateKey *string `json:"privateKey,omitempty" xml:"privateKey,omitempty"`
  // {"en" : "If not present, the value will be copied from the latest version of the certificate. Please refer to the description of the certificate field in the Create a certificate API for details about the format.", "zh_CN": "如果未指定该字段，则将从证书的最新版本中复制。 具体格式请参考'创建证书'接口中certificate字段的说明。"}
  Certificate *string `json:"certificate,omitempty" xml:"certificate,omitempty"`
  // {"en" : "This field must be filled in if the privateKey and certificate fields are both omitted. In this case, only the chain certificate will be updated. The chain certificate must be in PEM format.", "zh_CN": "当privateKey和certificate字段都未指定时，该字段必须填写。在这种情况下，只有链证书将被更新。链证书的格式必须为PEM。"}
  ChainCert *string `json:"chainCert,omitempty" xml:"chainCert,omitempty"`
  // {"en" : "Information submitted when generating a self-signed certificate.", "zh_CN": "生成自签名证书时提交的信息。"}
  IdentificationInfo *UpdateACertificateRequestNewVersionIdentificationInfo `json:"identificationInfo,omitempty" xml:"identificationInfo,omitempty" type:"Struct"`
}

func (s UpdateACertificateRequestNewVersion) String() string {
  return tea.Prettify(s)
}

func (s UpdateACertificateRequestNewVersion) GoString() string {
  return s.String()
}

func (s *UpdateACertificateRequestNewVersion) SetPrivateKey(v string) *UpdateACertificateRequestNewVersion {
  s.PrivateKey = &v
  return s
}

func (s *UpdateACertificateRequestNewVersion) SetCertificate(v string) *UpdateACertificateRequestNewVersion {
  s.Certificate = &v
  return s
}

func (s *UpdateACertificateRequestNewVersion) SetChainCert(v string) *UpdateACertificateRequestNewVersion {
  s.ChainCert = &v
  return s
}

func (s *UpdateACertificateRequestNewVersion) SetIdentificationInfo(v *UpdateACertificateRequestNewVersionIdentificationInfo) *UpdateACertificateRequestNewVersion {
  s.IdentificationInfo = v
  return s
}

type UpdateACertificateRequestNewVersionIdentificationInfo struct {
  // {"en" : "Range: [ 2 .. 2 ] characters 
  // An ISO-3166 country code.", "zh_CN": "取值范围: [ 2 .. 2 ] 字符 
  // ISO-3166国家代码。"}
  Country *string `json:"country,omitempty" xml:"country,omitempty"`
  // {"en" : "A state or province.", "zh_CN": "州或省。"}
  State *string `json:"state,omitempty" xml:"state,omitempty"`
  // {"en" : "A city.", "zh_CN": "城市。"}
  City *string `json:"city,omitempty" xml:"city,omitempty"`
  // {"en" : "A company name.", "zh_CN": "公司名称。"}
  Company *string `json:"company,omitempty" xml:"company,omitempty"`
  // {"en" : "The department associated with the certificate.", "zh_CN": "与证书关联的部门。"}
  Department *string `json:"department,omitempty" xml:"department,omitempty"`
  // {"en" : "A common name for the certificate.", "zh_CN": "证书的通用名称。"}
  CommonName *string `json:"commonName,omitempty" xml:"commonName,omitempty" require:"true"`
  // {"en" : "An email address.", "zh_CN": "邮箱地址。"}
  Email *string `json:"email,omitempty" xml:"email,omitempty"`
  // {"en" : "Hostnames that this certificate will serve. Each entry must be a valid hostname such as domain.com or a wildcard hostname beginning with '*' such as *.domain.com.", "zh_CN": "此证书涵盖的加速域名。每个条目必须是有效的加速域名，例如 domain.com
  // 或者是以'*'开头的泛域名，例如：*.domain.com。"}
  SubjectAlternativeNames []*string `json:"subjectAlternativeNames,omitempty" xml:"subjectAlternativeNames,omitempty" require:"true" type:"Repeated"`
}

func (s UpdateACertificateRequestNewVersionIdentificationInfo) String() string {
  return tea.Prettify(s)
}

func (s UpdateACertificateRequestNewVersionIdentificationInfo) GoString() string {
  return s.String()
}

func (s *UpdateACertificateRequestNewVersionIdentificationInfo) SetCountry(v string) *UpdateACertificateRequestNewVersionIdentificationInfo {
  s.Country = &v
  return s
}

func (s *UpdateACertificateRequestNewVersionIdentificationInfo) SetState(v string) *UpdateACertificateRequestNewVersionIdentificationInfo {
  s.State = &v
  return s
}

func (s *UpdateACertificateRequestNewVersionIdentificationInfo) SetCity(v string) *UpdateACertificateRequestNewVersionIdentificationInfo {
  s.City = &v
  return s
}

func (s *UpdateACertificateRequestNewVersionIdentificationInfo) SetCompany(v string) *UpdateACertificateRequestNewVersionIdentificationInfo {
  s.Company = &v
  return s
}

func (s *UpdateACertificateRequestNewVersionIdentificationInfo) SetDepartment(v string) *UpdateACertificateRequestNewVersionIdentificationInfo {
  s.Department = &v
  return s
}

func (s *UpdateACertificateRequestNewVersionIdentificationInfo) SetCommonName(v string) *UpdateACertificateRequestNewVersionIdentificationInfo {
  s.CommonName = &v
  return s
}

func (s *UpdateACertificateRequestNewVersionIdentificationInfo) SetEmail(v string) *UpdateACertificateRequestNewVersionIdentificationInfo {
  s.Email = &v
  return s
}

func (s *UpdateACertificateRequestNewVersionIdentificationInfo) SetSubjectAlternativeNames(v []*string) *UpdateACertificateRequestNewVersionIdentificationInfo {
  s.SubjectAlternativeNames = v
  return s
}

type UpdateACertificateResponse struct {
}

func (s UpdateACertificateResponse) String() string {
  return tea.Prettify(s)
}

func (s UpdateACertificateResponse) GoString() string {
  return s.String()
}

type UpdateACertificateResponseHeader struct {
  // {"en":"The Location header returns a URL to the specific certificate version if a new one is created.  Example:  <code>Location: http://open.chinanetcenter.com/cdn/certificates/329f12c1fe6708c23c31e91f/versions/5</code>", "zh_CN":"通过Location响应头返回新创建的证书版本的URL。示例：<code>Location:http://open.chinanetcenter.com/cdn/certificates/329f12c1fe6708c23c31e91f/versions/5</code>"}
  Location *string `json:"Location,omitempty" xml:"Location,omitempty" require:"true"`
}

func (s UpdateACertificateResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s UpdateACertificateResponseHeader) GoString() string {
  return s.String()
}

func (s *UpdateACertificateResponseHeader) SetLocation(v string) *UpdateACertificateResponseHeader {
  s.Location = &v
  return s
}




type GetListOfCertificatesPaths struct {
}

func (s GetListOfCertificatesPaths) String() string {
  return tea.Prettify(s)
}

func (s GetListOfCertificatesPaths) GoString() string {
  return s.String()
}

type GetListOfCertificatesParameters struct {
  // {"en" : "This parameter specifies a keyword to be searched in the 'name' field of the certificate.", "zh_CN": "指定关键字，按证书名称查询证书。"}
  Search *string `json:"search,omitempty" xml:"search,omitempty"`
  // {"en" : "Default: 0 Range: >= 0 
  // The offset indicates the index of the first certificate to return. The default is 0.", "zh_CN": "默认值: 0 取值范围: >= 0 
  // 查询起始位置。"}
  Offset *int `json:"offset,omitempty" xml:"offset,omitempty"`
  // {"en" : "Default: 100 Range: <= 200 
  // The limit indicates the maximum number of certificates to return.", "zh_CN": "默认值: 100 取值范围: <= 200 
  // 每次查询的最大条数。"}
  Limit *int `json:"limit,omitempty" xml:"limit,omitempty"`
  // {"en" : "Enum: all,staging,production 
  // Filters the results based on where the certificate has been deployed. <table><tr><th>Value</th><th>Effect</th></tr><tr><td></td><td>An empty target parameter results in all certificates being returned whether or not they have been deployed.</td></tr><td>all</td><td>Return all certificates deployed to either staging or production environments.</td></tr><tr><td>staging</td><td>Only return certificates deployed to staging.</td></tr><tr><td>production</td><td>Only return certificates deployed to production.</td></tr></table>
  // ", "zh_CN": "取值范围: all,staging,production 
  // 根据证书的部署环境查询证书。 <table> <tr><th>取值</th><th>返回数据</th></tr> <tr><td></td><td>当该参数为空时，返回所有证书，无论证书是否已部署。</td></tr> <tr><td>all</td><td>返回所有已部署的证书，包括部署到演练环境和生产环境的证书。</td></tr> <tr><td>staging</td><td>只返回部署到演练环境的证书。</td></tr> <tr><td>production</td><td>只返回部署到生产环境的证书。</td></tr> </table>"}
  Target *string `json:"target,omitempty" xml:"target,omitempty"`
  // {"en" : "Enum: LE,Off 
  // Filter results based on auto renewal status. By default, all certificates are returned.", "zh_CN": "取值范围: LE,Off 
  // 根据是否设置自动更新来查询证书。LE指通过Let's Encrypt自动更新，Off指未开启自动更新。"}
  AutoRenew *string `json:"autoRenew,omitempty" xml:"autoRenew,omitempty"`
  // {"en" : "Enum: creationTime,lastUpdateTime,expirationTime,productionExpirationTime 
  // Default: lastUpdateTime 
  // Returns certificates sorted by this field.", "zh_CN": "取值范围: creationTime,lastUpdateTime,expirationTime,productionExpirationTime 
  // 默认值: lastUpdateTime 
  // 查询结果的排序依据。支持按证书的创建时间，最后一次更新时间，证书过期时间，证书在生产环境中的版本的过期时间来排序。"}
  SortBy *string `json:"sortBy,omitempty" xml:"sortBy,omitempty"`
  // {"en" : "Enum: asc,desc 
  // Default: desc 
  // Return results sorted in this order ('asc' for ascending order, 'desc' for descending order).
  // ", "zh_CN": "取值范围: asc,desc 
  // 默认值: desc 
  // 查询结果的排序顺序，即升序（asc）或者降序（desc）。"}
  SortOrder *string `json:"sortOrder,omitempty" xml:"sortOrder,omitempty"`
}

func (s GetListOfCertificatesParameters) String() string {
  return tea.Prettify(s)
}

func (s GetListOfCertificatesParameters) GoString() string {
  return s.String()
}

func (s *GetListOfCertificatesParameters) SetSearch(v string) *GetListOfCertificatesParameters {
  s.Search = &v
  return s
}

func (s *GetListOfCertificatesParameters) SetOffset(v int) *GetListOfCertificatesParameters {
  s.Offset = &v
  return s
}

func (s *GetListOfCertificatesParameters) SetLimit(v int) *GetListOfCertificatesParameters {
  s.Limit = &v
  return s
}

func (s *GetListOfCertificatesParameters) SetTarget(v string) *GetListOfCertificatesParameters {
  s.Target = &v
  return s
}

func (s *GetListOfCertificatesParameters) SetAutoRenew(v string) *GetListOfCertificatesParameters {
  s.AutoRenew = &v
  return s
}

func (s *GetListOfCertificatesParameters) SetSortBy(v string) *GetListOfCertificatesParameters {
  s.SortBy = &v
  return s
}

func (s *GetListOfCertificatesParameters) SetSortOrder(v string) *GetListOfCertificatesParameters {
  s.SortOrder = &v
  return s
}

type GetListOfCertificatesRequestHeader struct {
}

func (s GetListOfCertificatesRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetListOfCertificatesRequestHeader) GoString() string {
  return s.String()
}

type GetListOfCertificatesRequest struct {
}

func (s GetListOfCertificatesRequest) String() string {
  return tea.Prettify(s)
}

func (s GetListOfCertificatesRequest) GoString() string {
  return s.String()
}

type GetListOfCertificatesResponseHeader struct {
}

func (s GetListOfCertificatesResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetListOfCertificatesResponseHeader) GoString() string {
  return s.String()
}

type GetListOfCertificatesResponse struct {
  // {"en" : "Range: >= 0 
  // Total number of available certificates. The actual number returned depends on the query parameters.", "zh_CN": "取值范围: >= 0 
  // 证书的总数。返回的实际数量取决于查询参数。"}
  Count *int `json:"count,omitempty" xml:"count,omitempty" require:"true"`
  // {"en" : "List of certificates.", "zh_CN": "证书列表。"}
  Certificates []*GetListOfCertificatesResponseCertificates `json:"certificates,omitempty" xml:"certificates,omitempty" require:"true" type:"Repeated"`
}

func (s GetListOfCertificatesResponse) String() string {
  return tea.Prettify(s)
}

func (s GetListOfCertificatesResponse) GoString() string {
  return s.String()
}

func (s *GetListOfCertificatesResponse) SetCount(v int) *GetListOfCertificatesResponse {
  s.Count = &v
  return s
}

func (s *GetListOfCertificatesResponse) SetCertificates(v []*GetListOfCertificatesResponseCertificates) *GetListOfCertificatesResponse {
  s.Certificates = v
  return s
}

type GetListOfCertificatesResponseCertificates struct     {
  // {"en" : "An ID representing the certificate. You can call GET /cdn/certificates/{certificate ID} to get details about a certificate.", "zh_CN": "证书的ID。您可以通过调用'查询证书详情'接口来获取证书的详细信息。"}
  CertificateId *string `json:"certificateId,omitempty" xml:"certificateId,omitempty"`
  // {"en" : "Name of the certificate.", "zh_CN": "证书名称。"}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // {"en" : "Range: >= 1 
  // The latest version of the certificate.", "zh_CN": "取值范围: >= 1 
  // 证书的最新版本。"}
  LatestVersion *int `json:"latestVersion,omitempty" xml:"latestVersion,omitempty"`
  // {"en" : "Enum: LE,Off 
  // This is set to 'LE' if you have chosen to autorenew the certificate using Let's Encrypt.", "zh_CN": "取值范围: LE,Off 
  // 如果您选择使用Let's Encrypt自动续订证书，则将此项设置为'LE'。"}
  AutoRenew *string `json:"autoRenew,omitempty" xml:"autoRenew,omitempty"`
  // {"en" : "An RFC3339 format date indicates when the certificate was added to the system.", "zh_CN": "RFC3339格式的日期，表示证书在CDN Pro的创建时间。"}
  CreationTime *string `json:"creationTime,omitempty" xml:"creationTime,omitempty"`
  // {"en" : "An RFC3339 format date indicating when the latest version of the certificate expires.
  // ", "zh_CN": "RFC3339格式的日期，表示证书最新版本的到期时间。
  // "}
  ExpirationTime *string `json:"expirationTime,omitempty" xml:"expirationTime,omitempty"`
  // {"en" : "An RFC3339 format date indicates when the certificate was last updated.
  // ", "zh_CN": "RFC3339格式的日期，表示证书最近一次更新的时间。
  // "}
  LastUpdateTime *string `json:"lastUpdateTime,omitempty" xml:"lastUpdateTime,omitempty"`
  // {"en" : "Range: >= 1 
  // Indicates the version of the certificate deployed to production.", "zh_CN": "取值范围: >= 1 
  // 证书部署到生产环境的版本。"}
  VersionInProduction *int `json:"versionInProduction,omitempty" xml:"versionInProduction,omitempty"`
  // {"en" : "Range: >= 1 
  // Indicates the version of the certificate deployed to staging.", "zh_CN": "取值范围: >= 1 
  // 证书部署到演练环境的版本。"}
  VersionInStaging *int `json:"versionInStaging,omitempty" xml:"versionInStaging,omitempty"`
  // {"en" : "RFC 3339 date indicating when the version deployed to production will expire.", "zh_CN": "RFC3339格式的日期，表示部署到生产环境的版本的过期时间。"}
  ProductionExpirationTime *string `json:"productionExpirationTime,omitempty" xml:"productionExpirationTime,omitempty"`
  // {"en" : "RFC 3339 date indicating when the version deployed to staging will expire.", "zh_CN": "RFC3339格式的日期，表示部署到演练环境的版本的过期时间。"}
  StagingExpirationTime *string `json:"stagingExpirationTime,omitempty" xml:"stagingExpirationTime,omitempty"`
  // {"en" : "Setting the value to true requests the certificate to be auto-renewed as soon as possible instead of waiting for the certificate to expire in 15 days. The value will be set to false after a successful renewal.", "zh_CN": "是否强制更新。当值为true时表示要求尽快自动更新证书，而不是等待证书在 15 天后过期。 证书成功更新后，该值将设置为false。"}
  ForceRenew *bool `json:"forceRenew,omitempty" xml:"forceRenew,omitempty"`
}

func (s GetListOfCertificatesResponseCertificates) String() string {
  return tea.Prettify(s)
}

func (s GetListOfCertificatesResponseCertificates) GoString() string {
  return s.String()
}

func (s *GetListOfCertificatesResponseCertificates) SetCertificateId(v string) *GetListOfCertificatesResponseCertificates {
  s.CertificateId = &v
  return s
}

func (s *GetListOfCertificatesResponseCertificates) SetName(v string) *GetListOfCertificatesResponseCertificates {
  s.Name = &v
  return s
}

func (s *GetListOfCertificatesResponseCertificates) SetLatestVersion(v int) *GetListOfCertificatesResponseCertificates {
  s.LatestVersion = &v
  return s
}

func (s *GetListOfCertificatesResponseCertificates) SetAutoRenew(v string) *GetListOfCertificatesResponseCertificates {
  s.AutoRenew = &v
  return s
}

func (s *GetListOfCertificatesResponseCertificates) SetCreationTime(v string) *GetListOfCertificatesResponseCertificates {
  s.CreationTime = &v
  return s
}

func (s *GetListOfCertificatesResponseCertificates) SetExpirationTime(v string) *GetListOfCertificatesResponseCertificates {
  s.ExpirationTime = &v
  return s
}

func (s *GetListOfCertificatesResponseCertificates) SetLastUpdateTime(v string) *GetListOfCertificatesResponseCertificates {
  s.LastUpdateTime = &v
  return s
}

func (s *GetListOfCertificatesResponseCertificates) SetVersionInProduction(v int) *GetListOfCertificatesResponseCertificates {
  s.VersionInProduction = &v
  return s
}

func (s *GetListOfCertificatesResponseCertificates) SetVersionInStaging(v int) *GetListOfCertificatesResponseCertificates {
  s.VersionInStaging = &v
  return s
}

func (s *GetListOfCertificatesResponseCertificates) SetProductionExpirationTime(v string) *GetListOfCertificatesResponseCertificates {
  s.ProductionExpirationTime = &v
  return s
}

func (s *GetListOfCertificatesResponseCertificates) SetStagingExpirationTime(v string) *GetListOfCertificatesResponseCertificates {
  s.StagingExpirationTime = &v
  return s
}

func (s *GetListOfCertificatesResponseCertificates) SetForceRenew(v bool) *GetListOfCertificatesResponseCertificates {
  s.ForceRenew = &v
  return s
}




