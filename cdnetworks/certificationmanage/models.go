package certificationmanage

// This file is auto-generated, don't edit it. Thanks.
import (
  "github.com/alibabacloud-go/tea/tea"
)

type ReissueCertificateForWplusRequest struct {
  // {"en":"Certificate id", "zh_CN":"证书ID"}
  CertificateId *int64 `json:"certificateId,omitempty" xml:"certificateId,omitempty" require:"true"`
  // {"en":"Certificate description", "zh_CN":"证书描述"}
  Description *string `json:"description,omitempty" xml:"description,omitempty"`
  // {"en":"Certificate algorithm, RSA2048 or ECDSA256", "zh_CN":"证书算法，取值范围：RSA2048、ECDSA256"}
  Algorithm *string `json:"algorithm,omitempty" xml:"algorithm,omitempty" require:"true"`
  // {"en":"Certificate validate method, HTTP or DNS", "zh_CN":"验证方式，取值范围：HTTP、DNS"}
  ValidateMethod *string `json:"validateMethod,omitempty" xml:"validateMethod,omitempty" require:"true"`
  // {"en":"Is auto validate, true or false", "zh_CN":"是否自动验证，取值范围：true、false"}
  AutoValidate *string `json:"autoValidate,omitempty" xml:"autoValidate,omitempty" require:"true"`
  // {"en":"Is auto deploy, true or false", "zh_CN":"是否自动部署，取值范围：true、false"}
  AutoDeploy *string `json:"autoDeploy,omitempty" xml:"autoDeploy,omitempty" require:"true"`
  // {"en":"Common name", "zh_CN":"通用名称"}
  CommonName *string `json:"commonName,omitempty" xml:"commonName,omitempty" require:"true"`
  // {"en":"Subject alternative name", "zh_CN":"主体备用名称"}
  SubjectAlternativeNames []*string `json:"subjectAlternativeNames,omitempty" xml:"subjectAlternativeNames,omitempty" require:"true" type:"Repeated"`
}

func (s ReissueCertificateForWplusRequest) String() string {
  return tea.Prettify(s)
}

func (s ReissueCertificateForWplusRequest) GoString() string {
  return s.String()
}

func (s *ReissueCertificateForWplusRequest) SetCertificateId(v int64) *ReissueCertificateForWplusRequest {
  s.CertificateId = &v
  return s
}

func (s *ReissueCertificateForWplusRequest) SetDescription(v string) *ReissueCertificateForWplusRequest {
  s.Description = &v
  return s
}

func (s *ReissueCertificateForWplusRequest) SetAlgorithm(v string) *ReissueCertificateForWplusRequest {
  s.Algorithm = &v
  return s
}

func (s *ReissueCertificateForWplusRequest) SetValidateMethod(v string) *ReissueCertificateForWplusRequest {
  s.ValidateMethod = &v
  return s
}

func (s *ReissueCertificateForWplusRequest) SetAutoValidate(v string) *ReissueCertificateForWplusRequest {
  s.AutoValidate = &v
  return s
}

func (s *ReissueCertificateForWplusRequest) SetAutoDeploy(v string) *ReissueCertificateForWplusRequest {
  s.AutoDeploy = &v
  return s
}

func (s *ReissueCertificateForWplusRequest) SetCommonName(v string) *ReissueCertificateForWplusRequest {
  s.CommonName = &v
  return s
}

func (s *ReissueCertificateForWplusRequest) SetSubjectAlternativeNames(v []*string) *ReissueCertificateForWplusRequest {
  s.SubjectAlternativeNames = v
  return s
}

type ReissueCertificateForWplusResponse struct {
  // {"en":"code", "zh_CN":"错误码，成功为0"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"error message", "zh_CN":"错误信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Sale order id", "zh_CN":"销售订单id"}
  Data *string `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s ReissueCertificateForWplusResponse) String() string {
  return tea.Prettify(s)
}

func (s ReissueCertificateForWplusResponse) GoString() string {
  return s.String()
}

func (s *ReissueCertificateForWplusResponse) SetCode(v string) *ReissueCertificateForWplusResponse {
  s.Code = &v
  return s
}

func (s *ReissueCertificateForWplusResponse) SetMessage(v string) *ReissueCertificateForWplusResponse {
  s.Message = &v
  return s
}

func (s *ReissueCertificateForWplusResponse) SetData(v string) *ReissueCertificateForWplusResponse {
  s.Data = &v
  return s
}

type ReissueCertificateForWplusPaths struct {
}

func (s ReissueCertificateForWplusPaths) String() string {
  return tea.Prettify(s)
}

func (s ReissueCertificateForWplusPaths) GoString() string {
  return s.String()
}

type ReissueCertificateForWplusParameters struct {
}

func (s ReissueCertificateForWplusParameters) String() string {
  return tea.Prettify(s)
}

func (s ReissueCertificateForWplusParameters) GoString() string {
  return s.String()
}

type ReissueCertificateForWplusRequestHeader struct {
}

func (s ReissueCertificateForWplusRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ReissueCertificateForWplusRequestHeader) GoString() string {
  return s.String()
}

type ReissueCertificateForWplusResponseHeader struct {
}

func (s ReissueCertificateForWplusResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ReissueCertificateForWplusResponseHeader) GoString() string {
  return s.String()
}




type GetCertificateContentRequest struct {
}

func (s GetCertificateContentRequest) String() string {
  return tea.Prettify(s)
}

func (s GetCertificateContentRequest) GoString() string {
  return s.String()
}

type GetCertificateContentResponse struct {
  // {"en":"certificate id", "zh_CN":"证书ID"}
  CertificateId *int64 `json:"certificate-id,omitempty" xml:"certificate-id,omitempty" require:"true"`
  // {"en":"The certificate name is unique under a customer.
  // Note:
  // 1. Support single domain, multi domain and wildcard domain name.
  // 2. For example, the authorized domain name of certificate A is *.example.com , a.example2.com, b.example2.com test.example.com , a.example2.com, can be associated with certificate A, while the domain name c.example2.com cannot use certificate A.", "zh_CN":"证书名称，客户粒度下是唯一的
  // 
  // 注意：
  // 1、支持单域名、多域名和泛域名证书，即证书的授权域名允许为多个域名，或者泛域名
  // 2、例如：证书A的授权域名为 *.example.com，a.example2.com，b.example2.com，则域名test.example.com，a.example2.com，都可以关联使用证书A，而域名c.example2.com不能使用证书A"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"Comment for a certificate.", "zh_CN":"证书文件的备注"}
  Comment *string `json:"comment,omitempty" xml:"comment,omitempty" require:"true"`
  // {"en":"Specified a shared certificate. True means shared.", "zh_CN":"是否共享，true表示共享证书，false表示非共享证书"}
  ShareSsl *bool `json:"share-ssl:,omitempty" xml:"share-ssl:,omitempty" require:"true"`
  // {"en":"Certificate content. The encryption algorithm is the same as the creating a certificate.
  // Encryption algorithm: First get md5 value of the request header 'Date'. Then take the left 8 bits of MD5 value as the key and the right 8 bits as the IV. At lase, encode the encrypted binary content with Base64. For example:
  // date=`env LANG=\"en_US.UTF-8\" date -u \"+%a, %d %b %Y %H:%M:%S GMT\"`
  // md5str=`echo -n $date | openssl md5`
  // key=`echo -n ${md5str:$((-32)):$((8))}|hexdump -e '8/1 \"%02X\"'`
  // iv=`echo -n ${md5str:$((-8))}|hexdump -e '8/1 \"%02X\"'`
  // Note:
  // 1.Decryption command: file is the content of the certificate file (CRT, key, CA) responded by the interface
  // cat $file |base64 -d|openssl enc -des -K $key -iv $iv -nosalt -d
  // 2. The date time of the HTTP request header must be the same as that of the above certificate encryption.", "zh_CN":"crt文件内容，已加密。加密算法同新增证书的加密算法。
  // 加密算法说明：将http头中的Date值做md5运算，将md5值的左8位作为key，右8位作为iv，然后对文件内容作des加密，将加密后的二进制内容作base64编码，以下为示例参考：
  // date=`env LANG=\"en_US.UTF-8\" date -u \"+%a, %d %b %Y %H:%M:%S GMT\"`
  // md5str=`echo -n $date | openssl md5`
  // key=`echo -n ${md5str:$((-32)):$((8))}|hexdump -e '8/1 \"%02X\"'`
  // iv=`echo -n ${md5str:$((-8))}|hexdump -e '8/1 \"%02X\"'`
  // 注意：
  // 1、解密命令如下，file就是接口响应的证书文件内容（crt，key，ca）
  // cat $file |base64 -d|openssl enc -des -K $key -iv $iv -nosalt -d
  // 2、http请求头的Date时间必须和上述证书加解密的时间一致"}
  SslCertificate *string `json:"ssl-certificate,omitempty" xml:"ssl-certificate,omitempty"`
  // {"en":"Certificate key. The encryption algorithm is the same as creating a  new certificate", "zh_CN":"key文件内容，已加密。加密算法同新增证书的加密算法。"}
  SslKey *string `json:"ssl-key,omitempty" xml:"ssl-key,omitempty"`
  // {"en":"Certificate chain. The encryption algorithm is the same as the new certificate.", "zh_CN":"ca文件内容，已加密。加密方式同新增证书的加密方式。"}
  SslCertificateChain *string `json:"ssl-certificate-chain,omitempty" xml:"ssl-certificate-chain,omitempty"`
}

func (s GetCertificateContentResponse) String() string {
  return tea.Prettify(s)
}

func (s GetCertificateContentResponse) GoString() string {
  return s.String()
}

func (s *GetCertificateContentResponse) SetCertificateId(v int64) *GetCertificateContentResponse {
  s.CertificateId = &v
  return s
}

func (s *GetCertificateContentResponse) SetName(v string) *GetCertificateContentResponse {
  s.Name = &v
  return s
}

func (s *GetCertificateContentResponse) SetComment(v string) *GetCertificateContentResponse {
  s.Comment = &v
  return s
}

func (s *GetCertificateContentResponse) SetShareSsl(v bool) *GetCertificateContentResponse {
  s.ShareSsl = &v
  return s
}

func (s *GetCertificateContentResponse) SetSslCertificate(v string) *GetCertificateContentResponse {
  s.SslCertificate = &v
  return s
}

func (s *GetCertificateContentResponse) SetSslKey(v string) *GetCertificateContentResponse {
  s.SslKey = &v
  return s
}

func (s *GetCertificateContentResponse) SetSslCertificateChain(v string) *GetCertificateContentResponse {
  s.SslCertificateChain = &v
  return s
}

type GetCertificateContentPaths struct {
  // {"en":"Certificate ID, corresponding to the* in the interface address", "zh_CN":"证书ID，对应接口url中的*
  // 注意：参看请求示例中的url，100166对应的就是certificate-id"}
  CertificateId *int64 `json:"certificate-id,omitempty" xml:"certificate-id,omitempty" require:"true"`
}

func (s GetCertificateContentPaths) String() string {
  return tea.Prettify(s)
}

func (s GetCertificateContentPaths) GoString() string {
  return s.String()
}

func (s *GetCertificateContentPaths) SetCertificateId(v int64) *GetCertificateContentPaths {
  s.CertificateId = &v
  return s
}

type GetCertificateContentParameters struct {
}

func (s GetCertificateContentParameters) String() string {
  return tea.Prettify(s)
}

func (s GetCertificateContentParameters) GoString() string {
  return s.String()
}

type GetCertificateContentRequestHeader struct {
  // {"en":"Date.  Formatting to comply with RFC 1123 specifications as for Thu, 17 May 2012 19:37:58 GMT ", "zh_CN":"时间戳，格式需符合RFC 1123规范，如Thu, 17 May 2012 19:37:58 GMT "}
  Date *string `json:"date,omitempty" xml:"date,omitempty" require:"true"`
}

func (s GetCertificateContentRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetCertificateContentRequestHeader) GoString() string {
  return s.String()
}

func (s *GetCertificateContentRequestHeader) SetDate(v string) *GetCertificateContentRequestHeader {
  s.Date = &v
  return s
}

type GetCertificateContentResponseHeader struct {
}

func (s GetCertificateContentResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetCertificateContentResponseHeader) GoString() string {
  return s.String()
}




type QueryDomainMultiCertConfigRequest struct {
}

func (s QueryDomainMultiCertConfigRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryDomainMultiCertConfigRequest) GoString() string {
  return s.String()
}

type QueryDomainMultiCertConfigCertInfo struct {
  // {"en":"certificateId", "zh_CN":"证书ID"}
  CertificateId *int64 `json:"certificateId,omitempty" xml:"certificateId,omitempty" require:"true"`
  // {"en":"certificateName", "zh_CN":"证书名称"}
  CertificateName *string `json:"certificateName,omitempty" xml:"certificateName,omitempty" require:"true"`
  // {"en":"comment", "zh_CN":"备注"}
  Comment *string `json:"comment,omitempty" xml:"comment,omitempty" require:"true"`
  // {"en":"serial", "zh_CN":"序列号"}
  Serial *string `json:"serial,omitempty" xml:"serial,omitempty" require:"true"`
}

func (s QueryDomainMultiCertConfigCertInfo) String() string {
  return tea.Prettify(s)
}

func (s QueryDomainMultiCertConfigCertInfo) GoString() string {
  return s.String()
}

func (s *QueryDomainMultiCertConfigCertInfo) SetCertificateId(v int64) *QueryDomainMultiCertConfigCertInfo {
  s.CertificateId = &v
  return s
}

func (s *QueryDomainMultiCertConfigCertInfo) SetCertificateName(v string) *QueryDomainMultiCertConfigCertInfo {
  s.CertificateName = &v
  return s
}

func (s *QueryDomainMultiCertConfigCertInfo) SetComment(v string) *QueryDomainMultiCertConfigCertInfo {
  s.Comment = &v
  return s
}

func (s *QueryDomainMultiCertConfigCertInfo) SetSerial(v string) *QueryDomainMultiCertConfigCertInfo {
  s.Serial = &v
  return s
}

type QueryDomainMultiCertConfigDomainCertConfig struct {
  // {"en":"certUsage", "zh_CN":"证书用途 "}
  CertUsage *string `json:"certUsage,omitempty" xml:"certUsage,omitempty" require:"true"`
  // {"en":"certInfos", "zh_CN":"证书信息"}
  CertInfos []*QueryDomainMultiCertConfigCertInfo `json:"certInfos,omitempty" xml:"certInfos,omitempty" require:"true" type:"Repeated"`
}

func (s QueryDomainMultiCertConfigDomainCertConfig) String() string {
  return tea.Prettify(s)
}

func (s QueryDomainMultiCertConfigDomainCertConfig) GoString() string {
  return s.String()
}

func (s *QueryDomainMultiCertConfigDomainCertConfig) SetCertUsage(v string) *QueryDomainMultiCertConfigDomainCertConfig {
  s.CertUsage = &v
  return s
}

func (s *QueryDomainMultiCertConfigDomainCertConfig) SetCertInfos(v []*QueryDomainMultiCertConfigCertInfo) *QueryDomainMultiCertConfigDomainCertConfig {
  s.CertInfos = v
  return s
}

type QueryDomainMultiCertConfigData struct {
  // {"en":"domainId", "zh_CN":"域名ID"}
  DomainId *string `json:"domainId,omitempty" xml:"domainId,omitempty" require:"true"`
  // {"en":"domainName", "zh_CN":"域名名称"}
  DomainName *string `json:"domainName,omitempty" xml:"domainName,omitempty" require:"true"`
  // {"en":"TLSVersion", "zh_CN":"TLS版本"}
  TLSVersion *string `json:"TLSVersion,omitempty" xml:"TLSVersion,omitempty" require:"true"`
  // {"en":"enableOCSP", "zh_CN":"启用OCSP，默认不启用，可选值true、false。没配置返回"}
  EnableOCSP *string `json:"enableOCSP,omitempty" xml:"enableOCSP,omitempty" require:"true"`
  // {"en":"cipherSuites", "zh_CN":"设置cipher加密套件"}
  CipherSuites *string `json:"cipherSuites,omitempty" xml:"cipherSuites,omitempty" require:"true"`
  // {"en":"domainCertConfigs", "zh_CN":"证书配置"}
  DomainCertConfigs []*QueryDomainMultiCertConfigDomainCertConfig `json:"domainCertConfigs,omitempty" xml:"domainCertConfigs,omitempty" require:"true" type:"Repeated"`
}

func (s QueryDomainMultiCertConfigData) String() string {
  return tea.Prettify(s)
}

func (s QueryDomainMultiCertConfigData) GoString() string {
  return s.String()
}

func (s *QueryDomainMultiCertConfigData) SetDomainId(v string) *QueryDomainMultiCertConfigData {
  s.DomainId = &v
  return s
}

func (s *QueryDomainMultiCertConfigData) SetDomainName(v string) *QueryDomainMultiCertConfigData {
  s.DomainName = &v
  return s
}

func (s *QueryDomainMultiCertConfigData) SetTLSVersion(v string) *QueryDomainMultiCertConfigData {
  s.TLSVersion = &v
  return s
}

func (s *QueryDomainMultiCertConfigData) SetEnableOCSP(v string) *QueryDomainMultiCertConfigData {
  s.EnableOCSP = &v
  return s
}

func (s *QueryDomainMultiCertConfigData) SetCipherSuites(v string) *QueryDomainMultiCertConfigData {
  s.CipherSuites = &v
  return s
}

func (s *QueryDomainMultiCertConfigData) SetDomainCertConfigs(v []*QueryDomainMultiCertConfigDomainCertConfig) *QueryDomainMultiCertConfigData {
  s.DomainCertConfigs = v
  return s
}

type QueryDomainMultiCertConfigResponse struct {
  // {"en":"code", "zh_CN":"错误码，成功为0"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"message", "zh_CN":"结果信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"data", "zh_CN":"返回数据"}
  QueryDomainMultiCertConfigData *QueryDomainMultiCertConfigData `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s QueryDomainMultiCertConfigResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryDomainMultiCertConfigResponse) GoString() string {
  return s.String()
}

func (s *QueryDomainMultiCertConfigResponse) SetCode(v string) *QueryDomainMultiCertConfigResponse {
  s.Code = &v
  return s
}

func (s *QueryDomainMultiCertConfigResponse) SetMessage(v string) *QueryDomainMultiCertConfigResponse {
  s.Message = &v
  return s
}

func (s *QueryDomainMultiCertConfigResponse) SetData(v *QueryDomainMultiCertConfigData) *QueryDomainMultiCertConfigResponse {
  s.QueryDomainMultiCertConfigData = v
  return s
}

type QueryDomainMultiCertConfigPaths struct {
  // {"en":"The domain whoes need query config.", "zh_CN":"需要查询配置的域名或域名id"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
}

func (s QueryDomainMultiCertConfigPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryDomainMultiCertConfigPaths) GoString() string {
  return s.String()
}

func (s *QueryDomainMultiCertConfigPaths) SetDomain(v string) *QueryDomainMultiCertConfigPaths {
  s.Domain = &v
  return s
}

type QueryDomainMultiCertConfigParameters struct {
}

func (s QueryDomainMultiCertConfigParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryDomainMultiCertConfigParameters) GoString() string {
  return s.String()
}

type QueryDomainMultiCertConfigRequestHeader struct {
}

func (s QueryDomainMultiCertConfigRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryDomainMultiCertConfigRequestHeader) GoString() string {
  return s.String()
}

type QueryDomainMultiCertConfigResponseHeader struct {
}

func (s QueryDomainMultiCertConfigResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryDomainMultiCertConfigResponseHeader) GoString() string {
  return s.String()
}




type DeleteCertificateRequest struct {
}

func (s DeleteCertificateRequest) String() string {
  return tea.Prettify(s)
}

func (s DeleteCertificateRequest) GoString() string {
  return s.String()
}

type DeleteCertificateResponse struct {
  // {"en":"Response code, 0 means successful.", "zh_CN":"接口响应code，0代表成功。"}
  Code *int `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response error message if failed.", "zh_CN":"接口响应信息，success代表成功，失败则提供失败信息。"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Response data array.", "zh_CN":"接口响应数据"}
  Data []*string `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s DeleteCertificateResponse) String() string {
  return tea.Prettify(s)
}

func (s DeleteCertificateResponse) GoString() string {
  return s.String()
}

func (s *DeleteCertificateResponse) SetCode(v int) *DeleteCertificateResponse {
  s.Code = &v
  return s
}

func (s *DeleteCertificateResponse) SetMessage(v string) *DeleteCertificateResponse {
  s.Message = &v
  return s
}

func (s *DeleteCertificateResponse) SetData(v []*string) *DeleteCertificateResponse {
  s.Data = v
  return s
}

type DeleteCertificatePaths struct {
  // {"en":"The certificate ID you want to delete.", "zh_CN":"需要删除的证书id"}
  CertificateId *int `json:"certificateId,omitempty" xml:"certificateId,omitempty" require:"true"`
}

func (s DeleteCertificatePaths) String() string {
  return tea.Prettify(s)
}

func (s DeleteCertificatePaths) GoString() string {
  return s.String()
}

func (s *DeleteCertificatePaths) SetCertificateId(v int) *DeleteCertificatePaths {
  s.CertificateId = &v
  return s
}

type DeleteCertificateParameters struct {
}

func (s DeleteCertificateParameters) String() string {
  return tea.Prettify(s)
}

func (s DeleteCertificateParameters) GoString() string {
  return s.String()
}

type DeleteCertificateRequestHeader struct {
}

func (s DeleteCertificateRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteCertificateRequestHeader) GoString() string {
  return s.String()
}

type DeleteCertificateResponseHeader struct {
}

func (s DeleteCertificateResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteCertificateResponseHeader) GoString() string {
  return s.String()
}




type QueryCertificateRelatedDomainsRequest struct {
}

func (s QueryCertificateRelatedDomainsRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryCertificateRelatedDomainsRequest) GoString() string {
  return s.String()
}

type QueryCertificateRelatedDomainsResponse struct {
  // {"en":"Response code, 0 means successful.", "zh_CN":"接口响应code，0代表成功。"}
  Code *int `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response error message if failed.", "zh_CN":"接口响应信息，success代表成功，失败则提供失败信息。"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Response data array.", "zh_CN":"接口响应数据"}
  Data *QueryCertificateRelatedDomainsResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
}

func (s QueryCertificateRelatedDomainsResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryCertificateRelatedDomainsResponse) GoString() string {
  return s.String()
}

func (s *QueryCertificateRelatedDomainsResponse) SetCode(v int) *QueryCertificateRelatedDomainsResponse {
  s.Code = &v
  return s
}

func (s *QueryCertificateRelatedDomainsResponse) SetMessage(v string) *QueryCertificateRelatedDomainsResponse {
  s.Message = &v
  return s
}

func (s *QueryCertificateRelatedDomainsResponse) SetData(v *QueryCertificateRelatedDomainsResponseData) *QueryCertificateRelatedDomainsResponse {
  s.Data = v
  return s
}

type QueryCertificateRelatedDomainsResponseData struct {
  // {"en":"Certificate ID", "zh_CN":"证书ID"}
  CertificateId *int `json:"certificateId,omitempty" xml:"certificateId,omitempty" require:"true"`
  // {"en":"Certificate related domains", "zh_CN":"在用该证书的加速域名列表"}
  Domains []*QueryCertificateRelatedDomainsResponseDataDomains `json:"domains,omitempty" xml:"domains,omitempty" require:"true" type:"Repeated"`
}

func (s QueryCertificateRelatedDomainsResponseData) String() string {
  return tea.Prettify(s)
}

func (s QueryCertificateRelatedDomainsResponseData) GoString() string {
  return s.String()
}

func (s *QueryCertificateRelatedDomainsResponseData) SetCertificateId(v int) *QueryCertificateRelatedDomainsResponseData {
  s.CertificateId = &v
  return s
}

func (s *QueryCertificateRelatedDomainsResponseData) SetDomains(v []*QueryCertificateRelatedDomainsResponseDataDomains) *QueryCertificateRelatedDomainsResponseData {
  s.Domains = v
  return s
}

type QueryCertificateRelatedDomainsResponseDataDomains struct     {
  // {"en":"Domain ID", "zh_CN":"域名id"}
  DomainId *int `json:"domainId,omitempty" xml:"domainId,omitempty" require:"true"`
  // {"en":"Domain name", "zh_CN":"域名名称"}
  DomainName *string `json:"domainName,omitempty" xml:"domainName,omitempty" require:"true"`
}

func (s QueryCertificateRelatedDomainsResponseDataDomains) String() string {
  return tea.Prettify(s)
}

func (s QueryCertificateRelatedDomainsResponseDataDomains) GoString() string {
  return s.String()
}

func (s *QueryCertificateRelatedDomainsResponseDataDomains) SetDomainId(v int) *QueryCertificateRelatedDomainsResponseDataDomains {
  s.DomainId = &v
  return s
}

func (s *QueryCertificateRelatedDomainsResponseDataDomains) SetDomainName(v string) *QueryCertificateRelatedDomainsResponseDataDomains {
  s.DomainName = &v
  return s
}

type QueryCertificateRelatedDomainsPaths struct {
  // {"en":"The certificate ID you want to query.", "zh_CN":"需要查询的证书id"}
  CertificateId *int `json:"certificateId,omitempty" xml:"certificateId,omitempty" require:"true"`
}

func (s QueryCertificateRelatedDomainsPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryCertificateRelatedDomainsPaths) GoString() string {
  return s.String()
}

func (s *QueryCertificateRelatedDomainsPaths) SetCertificateId(v int) *QueryCertificateRelatedDomainsPaths {
  s.CertificateId = &v
  return s
}

type QueryCertificateRelatedDomainsParameters struct {
}

func (s QueryCertificateRelatedDomainsParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryCertificateRelatedDomainsParameters) GoString() string {
  return s.String()
}

type QueryCertificateRelatedDomainsRequestHeader struct {
}

func (s QueryCertificateRelatedDomainsRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryCertificateRelatedDomainsRequestHeader) GoString() string {
  return s.String()
}

type QueryCertificateRelatedDomainsResponseHeader struct {
}

func (s QueryCertificateRelatedDomainsResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryCertificateRelatedDomainsResponseHeader) GoString() string {
  return s.String()
}




type QueryCertificateInfoRequest struct {
  // {"en":"The certificate ID you want to query.", "zh_CN":"指定要查询的证书ID，在URI上"}
  CertificateId *int `json:"certificateId,omitempty" xml:"certificateId,omitempty" require:"true"`
}

func (s QueryCertificateInfoRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryCertificateInfoRequest) GoString() string {
  return s.String()
}

func (s *QueryCertificateInfoRequest) SetCertificateId(v int) *QueryCertificateInfoRequest {
  s.CertificateId = &v
  return s
}

type QueryCertificateInfoResponse struct {
  // {"en":"Response code, 0 means successful.", "zh_CN":"接口响应code，0代表成功。"}
  Code *int `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response error message if failed.", "zh_CN":"接口响应信息，success代表成功，失败则提供失败信息。"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Response data array.", "zh_CN":"接口响应数据"}
  Data *QueryCertificateInfoResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
}

func (s QueryCertificateInfoResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryCertificateInfoResponse) GoString() string {
  return s.String()
}

func (s *QueryCertificateInfoResponse) SetCode(v int) *QueryCertificateInfoResponse {
  s.Code = &v
  return s
}

func (s *QueryCertificateInfoResponse) SetMessage(v string) *QueryCertificateInfoResponse {
  s.Message = &v
  return s
}

func (s *QueryCertificateInfoResponse) SetData(v *QueryCertificateInfoResponseData) *QueryCertificateInfoResponse {
  s.Data = v
  return s
}

type QueryCertificateInfoResponseData struct {
  // {"en":"certificate Id", "zh_CN":"证书ID。"}
  CertificateId *int `json:"certificateId,omitempty" xml:"certificateId,omitempty" require:"true"`
  // {"en":"certificate name", "zh_CN":"证书名称。"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"comment", "zh_CN":"备注信息。"}
  Comment *string `json:"comment,omitempty" xml:"comment,omitempty" require:"true"`
  // {"en":"certificate serial", "zh_CN":"证书序列号。"}
  Serial *string `json:"serial,omitempty" xml:"serial,omitempty" require:"true"`
  // {"en":"not befoe", "zh_CN":"签发时间。"}
  NotBefore *string `json:"notBefore,omitempty" xml:"notBefore,omitempty" require:"true"`
  // {"en":"not after", "zh_CN":"到期时间。"}
  NotAfter *string `json:"notAfter,omitempty" xml:"notAfter,omitempty" require:"true"`
  // {"en":"common name", "zh_CN":"证书绑定的主域名。"}
  CommonName *string `json:"commonName,omitempty" xml:"commonName,omitempty" require:"true"`
  // {"en":"Subject Alternative Names", "zh_CN":"证书绑定的附加域名列表。"}
  SubjectAlternativeNames []*QueryCertificateInfoResponseDataSubjectAlternativeNames `json:"subjectAlternativeNames,omitempty" xml:"subjectAlternativeNames,omitempty" require:"true" type:"Repeated"`
}

func (s QueryCertificateInfoResponseData) String() string {
  return tea.Prettify(s)
}

func (s QueryCertificateInfoResponseData) GoString() string {
  return s.String()
}

func (s *QueryCertificateInfoResponseData) SetCertificateId(v int) *QueryCertificateInfoResponseData {
  s.CertificateId = &v
  return s
}

func (s *QueryCertificateInfoResponseData) SetName(v string) *QueryCertificateInfoResponseData {
  s.Name = &v
  return s
}

func (s *QueryCertificateInfoResponseData) SetComment(v string) *QueryCertificateInfoResponseData {
  s.Comment = &v
  return s
}

func (s *QueryCertificateInfoResponseData) SetSerial(v string) *QueryCertificateInfoResponseData {
  s.Serial = &v
  return s
}

func (s *QueryCertificateInfoResponseData) SetNotBefore(v string) *QueryCertificateInfoResponseData {
  s.NotBefore = &v
  return s
}

func (s *QueryCertificateInfoResponseData) SetNotAfter(v string) *QueryCertificateInfoResponseData {
  s.NotAfter = &v
  return s
}

func (s *QueryCertificateInfoResponseData) SetCommonName(v string) *QueryCertificateInfoResponseData {
  s.CommonName = &v
  return s
}

func (s *QueryCertificateInfoResponseData) SetSubjectAlternativeNames(v []*QueryCertificateInfoResponseDataSubjectAlternativeNames) *QueryCertificateInfoResponseData {
  s.SubjectAlternativeNames = v
  return s
}

type QueryCertificateInfoResponseDataSubjectAlternativeNames struct     {
  // {"en":"Subject Alternative Name", "zh_CN":"证书绑定的附加域名"}
  SubjectAlternativeName *string `json:"subjectAlternativeName,omitempty" xml:"subjectAlternativeName,omitempty" require:"true"`
}

func (s QueryCertificateInfoResponseDataSubjectAlternativeNames) String() string {
  return tea.Prettify(s)
}

func (s QueryCertificateInfoResponseDataSubjectAlternativeNames) GoString() string {
  return s.String()
}

func (s *QueryCertificateInfoResponseDataSubjectAlternativeNames) SetSubjectAlternativeName(v string) *QueryCertificateInfoResponseDataSubjectAlternativeNames {
  s.SubjectAlternativeName = &v
  return s
}

type QueryCertificateInfoPaths struct {
  // {"en":"The certificate ID you want to query.", "zh_CN":"需要查询的证书id"}
  CertificateId *int `json:"certificateId,omitempty" xml:"certificateId,omitempty" require:"true"`
}

func (s QueryCertificateInfoPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryCertificateInfoPaths) GoString() string {
  return s.String()
}

func (s *QueryCertificateInfoPaths) SetCertificateId(v int) *QueryCertificateInfoPaths {
  s.CertificateId = &v
  return s
}

type QueryCertificateInfoParameters struct {
}

func (s QueryCertificateInfoParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryCertificateInfoParameters) GoString() string {
  return s.String()
}

type QueryCertificateInfoRequestHeader struct {
}

func (s QueryCertificateInfoRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryCertificateInfoRequestHeader) GoString() string {
  return s.String()
}

type QueryCertificateInfoResponseHeader struct {
}

func (s QueryCertificateInfoResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryCertificateInfoResponseHeader) GoString() string {
  return s.String()
}




type QueryCertificateListRequest struct {
}

func (s QueryCertificateListRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryCertificateListRequest) GoString() string {
  return s.String()
}

type QueryCertificateListResponse struct {
  // {"en":"Certificate list information", "zh_CN":"证书列表信息"}
  SslCertificates []*QueryCertificateListResponseSslCertificates `json:"ssl-certificates,omitempty" xml:"ssl-certificates,omitempty" require:"true" type:"Repeated"`
}

func (s QueryCertificateListResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryCertificateListResponse) GoString() string {
  return s.String()
}

func (s *QueryCertificateListResponse) SetSslCertificates(v []*QueryCertificateListResponseSslCertificates) *QueryCertificateListResponse {
  s.SslCertificates = v
  return s
}

type QueryCertificateListResponseSslCertificates struct     {
  // {"en":"Certificate ID", "zh_CN":"证书ID"}
  CertificateId *string `json:"certificate-id,omitempty" xml:"certificate-id,omitempty" require:"true"`
  // {"en":"Certificate name, unique to customer granularity", "zh_CN":"证书名称，客户粒度下是唯一的"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"Remarks on cerfiticate file", "zh_CN":"证书文件的备注"}
  Comment *string `json:"comment,omitempty" xml:"comment,omitempty" require:"true"`
  // {"en":"Shared, optional values are true and false, true represents shared certificates, false represents unshared certificates, default is false
  // This certificate allows cross-customer use when share-ssl is true. (The API does not support cross-customer use certificates. Contact customer service for manual configuration if required.)", "zh_CN":"是否共享，true表示共享证书，false表示非共享证书"}
  ShareSsl *string `json:"share-ssl,omitempty" xml:"share-ssl,omitempty" require:"true"`
  // {"en":"Certificate effective start time (CST), such as 2016-08-01 07:00:00", "zh_CN":"证书有效期的起始时间（CST时区），例如：2016-08-01 07:00:00"}
  CertificateValidityFrom *string `json:"certificate-validity-from,omitempty" xml:"certificate-validity-from,omitempty" require:"true"`
  // {"en":"Certificate effective end time (CST), such as 2018-08-01 19:00:00", "zh_CN":"证书有效期的到期时间（CST时区），例如：2018-08-01 19:00:00"}
  CertificateValidityTo *string `json:"certificate-validity-to,omitempty" xml:"certificate-validity-to,omitempty" require:"true"`
  // {"en":"List of domain names using the current certificate", "zh_CN":"使用当前证书的域名列表"}
  RelatedDomains []*QueryCertificateListResponseSslCertificatesRelatedDomains `json:"related-domains,omitempty" xml:"related-domains,omitempty" require:"true" type:"Repeated"`
  // {"en":"dns-names", "zh_CN":"授权域名列表，证书使用者可选名称，父标签"}
  DnsNames []*string `json:"dns-names,omitempty" xml:"dns-names,omitempty" require:"true" type:"Repeated"`
  // {"en":"The CRT certificate serial number", "zh_CN":"crt证书序列号"}
  CertificateSerial *string `json:"certificate-serial,omitempty" xml:"certificate-serial,omitempty" require:"true"`
  // {"en":"The MD5 value of CRT file.", "zh_CN":"CRT文件内容的md5值。"}
  CrtMd5 *string `json:"crt-md5,omitempty" xml:"crt-md5,omitempty" require:"true"`
  // {"en":"The MD5 value of KEY file.", "zh_CN":"KEY文件内容的md5值。"}
  KeyMd5 *string `json:"key-md5,omitempty" xml:"key-md5,omitempty" require:"true"`
  // {"en":"The MD5 value of CA file.", "zh_CN":"CA。"}
  CaMd5 *string `json:"ca-md5,omitempty" xml:"ca-md5,omitempty" require:"true"`
  // {"en":"The CRT certificate issuer", "zh_CN":"crt证书颁布者"}
  CertificateIssuer *string `json:"certificate-issuer,omitempty" xml:"certificate-issuer,omitempty" require:"true"`
}

func (s QueryCertificateListResponseSslCertificates) String() string {
  return tea.Prettify(s)
}

func (s QueryCertificateListResponseSslCertificates) GoString() string {
  return s.String()
}

func (s *QueryCertificateListResponseSslCertificates) SetCertificateId(v string) *QueryCertificateListResponseSslCertificates {
  s.CertificateId = &v
  return s
}

func (s *QueryCertificateListResponseSslCertificates) SetName(v string) *QueryCertificateListResponseSslCertificates {
  s.Name = &v
  return s
}

func (s *QueryCertificateListResponseSslCertificates) SetComment(v string) *QueryCertificateListResponseSslCertificates {
  s.Comment = &v
  return s
}

func (s *QueryCertificateListResponseSslCertificates) SetShareSsl(v string) *QueryCertificateListResponseSslCertificates {
  s.ShareSsl = &v
  return s
}

func (s *QueryCertificateListResponseSslCertificates) SetCertificateValidityFrom(v string) *QueryCertificateListResponseSslCertificates {
  s.CertificateValidityFrom = &v
  return s
}

func (s *QueryCertificateListResponseSslCertificates) SetCertificateValidityTo(v string) *QueryCertificateListResponseSslCertificates {
  s.CertificateValidityTo = &v
  return s
}

func (s *QueryCertificateListResponseSslCertificates) SetRelatedDomains(v []*QueryCertificateListResponseSslCertificatesRelatedDomains) *QueryCertificateListResponseSslCertificates {
  s.RelatedDomains = v
  return s
}

func (s *QueryCertificateListResponseSslCertificates) SetDnsNames(v []*string) *QueryCertificateListResponseSslCertificates {
  s.DnsNames = v
  return s
}

func (s *QueryCertificateListResponseSslCertificates) SetCertificateSerial(v string) *QueryCertificateListResponseSslCertificates {
  s.CertificateSerial = &v
  return s
}

func (s *QueryCertificateListResponseSslCertificates) SetCrtMd5(v string) *QueryCertificateListResponseSslCertificates {
  s.CrtMd5 = &v
  return s
}

func (s *QueryCertificateListResponseSslCertificates) SetKeyMd5(v string) *QueryCertificateListResponseSslCertificates {
  s.KeyMd5 = &v
  return s
}

func (s *QueryCertificateListResponseSslCertificates) SetCaMd5(v string) *QueryCertificateListResponseSslCertificates {
  s.CaMd5 = &v
  return s
}

func (s *QueryCertificateListResponseSslCertificates) SetCertificateIssuer(v string) *QueryCertificateListResponseSslCertificates {
  s.CertificateIssuer = &v
  return s
}

type QueryCertificateListResponseSslCertificatesRelatedDomains struct     {
  // {"en":"Accelerated domain name ID", "zh_CN":"加速域名ID"}
  DomainId *string `json:"domain-id,omitempty" xml:"domain-id,omitempty" require:"true"`
  // {"en":"Name of accelerated domain name", "zh_CN":"加速域名的名称"}
  DomainName *string `json:"domain-name,omitempty" xml:"domain-name,omitempty" require:"true"`
}

func (s QueryCertificateListResponseSslCertificatesRelatedDomains) String() string {
  return tea.Prettify(s)
}

func (s QueryCertificateListResponseSslCertificatesRelatedDomains) GoString() string {
  return s.String()
}

func (s *QueryCertificateListResponseSslCertificatesRelatedDomains) SetDomainId(v string) *QueryCertificateListResponseSslCertificatesRelatedDomains {
  s.DomainId = &v
  return s
}

func (s *QueryCertificateListResponseSslCertificatesRelatedDomains) SetDomainName(v string) *QueryCertificateListResponseSslCertificatesRelatedDomains {
  s.DomainName = &v
  return s
}

type QueryCertificateListPaths struct {
}

func (s QueryCertificateListPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryCertificateListPaths) GoString() string {
  return s.String()
}

type QueryCertificateListParameters struct {
}

func (s QueryCertificateListParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryCertificateListParameters) GoString() string {
  return s.String()
}

type QueryCertificateListRequestHeader struct {
}

func (s QueryCertificateListRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryCertificateListRequestHeader) GoString() string {
  return s.String()
}

type QueryCertificateListResponseHeader struct {
}

func (s QueryCertificateListResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryCertificateListResponseHeader) GoString() string {
  return s.String()
}




type QueryCertificateContentRequest struct {
}

func (s QueryCertificateContentRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryCertificateContentRequest) GoString() string {
  return s.String()
}

type QueryCertificateContentResponse struct {
  // {"en":"Response code, 0 means successful.", "zh_CN":"接口响应code，0代表成功。"}
  Code *int `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response error message if failed.", "zh_CN":"接口响应信息，success代表成功，失败则提供失败信息。"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Response data array.", "zh_CN":"接口响应数据"}
  Data *QueryCertificateContentResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
}

func (s QueryCertificateContentResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryCertificateContentResponse) GoString() string {
  return s.String()
}

func (s *QueryCertificateContentResponse) SetCode(v int) *QueryCertificateContentResponse {
  s.Code = &v
  return s
}

func (s *QueryCertificateContentResponse) SetMessage(v string) *QueryCertificateContentResponse {
  s.Message = &v
  return s
}

func (s *QueryCertificateContentResponse) SetData(v *QueryCertificateContentResponseData) *QueryCertificateContentResponse {
  s.Data = v
  return s
}

type QueryCertificateContentResponseData struct {
  // {"en":"Certificate ID", "zh_CN":"证书ID"}
  CertificateId *int `json:"certificateId,omitempty" xml:"certificateId,omitempty" require:"true"`
  // {"en":"Certificate content", "zh_CN":"证书内容，PEM格式"}
  Certificate *string `json:"certificate,omitempty" xml:"certificate,omitempty" require:"true"`
}

func (s QueryCertificateContentResponseData) String() string {
  return tea.Prettify(s)
}

func (s QueryCertificateContentResponseData) GoString() string {
  return s.String()
}

func (s *QueryCertificateContentResponseData) SetCertificateId(v int) *QueryCertificateContentResponseData {
  s.CertificateId = &v
  return s
}

func (s *QueryCertificateContentResponseData) SetCertificate(v string) *QueryCertificateContentResponseData {
  s.Certificate = &v
  return s
}

type QueryCertificateContentPaths struct {
  // {"en":"The certificate ID you want to query.", "zh_CN":"需要查询的证书id"}
  CertificateId *int `json:"certificateId,omitempty" xml:"certificateId,omitempty" require:"true"`
}

func (s QueryCertificateContentPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryCertificateContentPaths) GoString() string {
  return s.String()
}

func (s *QueryCertificateContentPaths) SetCertificateId(v int) *QueryCertificateContentPaths {
  s.CertificateId = &v
  return s
}

type QueryCertificateContentParameters struct {
}

func (s QueryCertificateContentParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryCertificateContentParameters) GoString() string {
  return s.String()
}

type QueryCertificateContentRequestHeader struct {
}

func (s QueryCertificateContentRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryCertificateContentRequestHeader) GoString() string {
  return s.String()
}

type QueryCertificateContentResponseHeader struct {
}

func (s QueryCertificateContentResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryCertificateContentResponseHeader) GoString() string {
  return s.String()
}




type EditCertificateV2Request struct {
  // {"en":"Certificate name", "zh_CN":"证书名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"Certificate, PEM certificate.", "zh_CN":"证书内容，PEM格式。例如：
  // -----BEGIN CERTIFICATE-----
  // ……
  // -----END CERTIFICATE-----"}
  Certificate *string `json:"certificate,omitempty" xml:"certificate,omitempty"`
  // {"en":"Private key of the certificate, PEM certificate. Not required when you specify a csrId.", "zh_CN":"证书密钥，PEM格式。例如：
  // -----BEGIN RSA PRIVATE KEY-----
  // ……
  // -----BEGIN RSA PRIVATE KEY-----"}
  PrivateKey *string `json:"privateKey,omitempty" xml:"privateKey,omitempty"`
  // {"en":"Automatically renew", "zh_CN":"是否自动续签，取值范围false或者true"}
  AutoRenew *string `json:"autoRenew,omitempty" xml:"autoRenew,omitempty"`
  // {"en":"Do you want to send an expiration alarm", "zh_CN":"是否需要证书过期告警，取值范围false或者true"}
  IsNeedAlarm *string `json:"isNeedAlarm,omitempty" xml:"isNeedAlarm,omitempty"`
  // {"en":"Csr ID, the id returned by the interface "Create CSR".", "zh_CN":"csrId，证书申请文件的id"}
  CsrId *int `json:"csrId,omitempty" xml:"csrId,omitempty"`
  // {"en":"Comment", "zh_CN":"备注信息"}
  Comment *string `json:"comment,omitempty" xml:"comment,omitempty"`
}

func (s EditCertificateV2Request) String() string {
  return tea.Prettify(s)
}

func (s EditCertificateV2Request) GoString() string {
  return s.String()
}

func (s *EditCertificateV2Request) SetName(v string) *EditCertificateV2Request {
  s.Name = &v
  return s
}

func (s *EditCertificateV2Request) SetCertificate(v string) *EditCertificateV2Request {
  s.Certificate = &v
  return s
}

func (s *EditCertificateV2Request) SetPrivateKey(v string) *EditCertificateV2Request {
  s.PrivateKey = &v
  return s
}

func (s *EditCertificateV2Request) SetAutoRenew(v string) *EditCertificateV2Request {
  s.AutoRenew = &v
  return s
}

func (s *EditCertificateV2Request) SetIsNeedAlarm(v string) *EditCertificateV2Request {
  s.IsNeedAlarm = &v
  return s
}

func (s *EditCertificateV2Request) SetCsrId(v int) *EditCertificateV2Request {
  s.CsrId = &v
  return s
}

func (s *EditCertificateV2Request) SetComment(v string) *EditCertificateV2Request {
  s.Comment = &v
  return s
}

type EditCertificateV2Response struct {
  // {"en":"Response code, 0 means successful.", "zh_CN":"接口响应code，0代表成功。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response error message if failed.", "zh_CN":"接口响应信息，success代表成功，失败则提供失败信息。"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Response data array.", "zh_CN":"接口响应数据"}
  Data *string `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s EditCertificateV2Response) String() string {
  return tea.Prettify(s)
}

func (s EditCertificateV2Response) GoString() string {
  return s.String()
}

func (s *EditCertificateV2Response) SetCode(v string) *EditCertificateV2Response {
  s.Code = &v
  return s
}

func (s *EditCertificateV2Response) SetMessage(v string) *EditCertificateV2Response {
  s.Message = &v
  return s
}

func (s *EditCertificateV2Response) SetData(v string) *EditCertificateV2Response {
  s.Data = &v
  return s
}

type EditCertificateV2Paths struct {
  // {"en":"The certificate ID you want to modify.", "zh_CN":"需要修改的证书id"}
  CertificateId *int `json:"certificateId,omitempty" xml:"certificateId,omitempty" require:"true"`
}

func (s EditCertificateV2Paths) String() string {
  return tea.Prettify(s)
}

func (s EditCertificateV2Paths) GoString() string {
  return s.String()
}

func (s *EditCertificateV2Paths) SetCertificateId(v int) *EditCertificateV2Paths {
  s.CertificateId = &v
  return s
}

type EditCertificateV2Parameters struct {
}

func (s EditCertificateV2Parameters) String() string {
  return tea.Prettify(s)
}

func (s EditCertificateV2Parameters) GoString() string {
  return s.String()
}

type EditCertificateV2RequestHeader struct {
}

func (s EditCertificateV2RequestHeader) String() string {
  return tea.Prettify(s)
}

func (s EditCertificateV2RequestHeader) GoString() string {
  return s.String()
}

type EditCertificateV2ResponseHeader struct {
}

func (s EditCertificateV2ResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s EditCertificateV2ResponseHeader) GoString() string {
  return s.String()
}




type AddCertificateServiceV2Request struct {
  // {"en":"Certificate name", "zh_CN":"证书名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"Certificate, PEM certificate, including CRT file and CA file.", "zh_CN":"证书内容，PEM格式，包含CRT文件、CA文件。例如：
  // -----BEGIN CERTIFICATE-----
  // ……
  // -----END CERTIFICATE-----"}
  Certificate *string `json:"certificate,omitempty" xml:"certificate,omitempty" require:"true"`
  // {"en":"Private key of the certificate, PEM certificate. Not required when you specify a csrId.", "zh_CN":"证书密钥，PEM格式。例如：
  // -----BEGIN RSA PRIVATE KEY-----
  // ……
  // -----BEGIN RSA PRIVATE KEY-----
  // 当指定csrId时，无需上传证书密钥。"}
  PrivateKey *string `json:"privateKey,omitempty" xml:"privateKey,omitempty"`
  // {"en":"Csr ID, the id returned by the interface \"Create CSR\".", "zh_CN":"csrId，证书申请文件的id。"}
  CsrId *int `json:"csrId,omitempty" xml:"csrId,omitempty"`
  // {"en":"comment", "zh_CN":"备注"}
  Comment *string `json:"comment,omitempty" xml:"comment,omitempty"`
}

func (s AddCertificateServiceV2Request) String() string {
  return tea.Prettify(s)
}

func (s AddCertificateServiceV2Request) GoString() string {
  return s.String()
}

func (s *AddCertificateServiceV2Request) SetName(v string) *AddCertificateServiceV2Request {
  s.Name = &v
  return s
}

func (s *AddCertificateServiceV2Request) SetCertificate(v string) *AddCertificateServiceV2Request {
  s.Certificate = &v
  return s
}

func (s *AddCertificateServiceV2Request) SetPrivateKey(v string) *AddCertificateServiceV2Request {
  s.PrivateKey = &v
  return s
}

func (s *AddCertificateServiceV2Request) SetCsrId(v int) *AddCertificateServiceV2Request {
  s.CsrId = &v
  return s
}

func (s *AddCertificateServiceV2Request) SetComment(v string) *AddCertificateServiceV2Request {
  s.Comment = &v
  return s
}

type AddCertificateServiceV2Response struct {
  // {"en":"Status code", "zh_CN":"状态码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response message", "zh_CN":"响应信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s AddCertificateServiceV2Response) String() string {
  return tea.Prettify(s)
}

func (s AddCertificateServiceV2Response) GoString() string {
  return s.String()
}

func (s *AddCertificateServiceV2Response) SetCode(v string) *AddCertificateServiceV2Response {
  s.Code = &v
  return s
}

func (s *AddCertificateServiceV2Response) SetMessage(v string) *AddCertificateServiceV2Response {
  s.Message = &v
  return s
}

type AddCertificateServiceV2Paths struct {
}

func (s AddCertificateServiceV2Paths) String() string {
  return tea.Prettify(s)
}

func (s AddCertificateServiceV2Paths) GoString() string {
  return s.String()
}

type AddCertificateServiceV2Parameters struct {
}

func (s AddCertificateServiceV2Parameters) String() string {
  return tea.Prettify(s)
}

func (s AddCertificateServiceV2Parameters) GoString() string {
  return s.String()
}

type AddCertificateServiceV2RequestHeader struct {
}

func (s AddCertificateServiceV2RequestHeader) String() string {
  return tea.Prettify(s)
}

func (s AddCertificateServiceV2RequestHeader) GoString() string {
  return s.String()
}

type AddCertificateServiceV2ResponseHeader struct {
  // {"en":"The URL used to access the certificate file where certificate-id is the unique token that the system generates for the certificate and whose value is a string", "zh_CN":"用于访问该证书文件的URL，其中certificate-id为系统为该证书生成的唯一标示，其值为字符串"}
  Location *string `json:"location,omitempty" xml:"location,omitempty" require:"true"`
}

func (s AddCertificateServiceV2ResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s AddCertificateServiceV2ResponseHeader) GoString() string {
  return s.String()
}

func (s *AddCertificateServiceV2ResponseHeader) SetLocation(v string) *AddCertificateServiceV2ResponseHeader {
  s.Location = &v
  return s
}




