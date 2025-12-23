package certificationmanage

// This file is auto-generated, don't edit it. Thanks.
import (
  "github.com/alibabacloud-go/tea/tea"
)

type QueryCertificateRequest struct {
}

func (s QueryCertificateRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryCertificateRequest) GoString() string {
  return s.String()
}

type QueryCertificateRequestHeader struct {
}

func (s QueryCertificateRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryCertificateRequestHeader) GoString() string {
  return s.String()
}

type QueryCertificatePaths struct {
  // {"en":"Certificate IDNote: 1. Refer to the url in the request example, 100166 for certificate-id2.After the new certificate is successfully submitted, the certificate-id corresponding to the certificate can be queried in the location access url in the return parameter. You can also query certificate-id through the QueryCertificateList interface","zh_CN":"证书ID注意：1、参看请求示例中的url，100166对应的就是certificate-id2、新增证书成功提交后，返回参数中的location访问url中，能够查询到证书对应的certificate-id；也可以通过【查看证书列表】接口查询到certificate-id"}
  CertificateId *int `json:"certificate-id,omitempty" xml:"certificate-id,omitempty" require:"true"`
}

func (s QueryCertificatePaths) String() string {
  return tea.Prettify(s)
}

func (s QueryCertificatePaths) GoString() string {
  return s.String()
}

func (s *QueryCertificatePaths) SetCertificateId(v int) *QueryCertificatePaths {
  s.CertificateId = &v
  return s
}

type QueryCertificateParameters struct {
}

func (s QueryCertificateParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryCertificateParameters) GoString() string {
  return s.String()
}

type QueryCertificateResponse struct {
  // {"en":"List of domain names using the current certificate","zh_CN":"使用当前证书的域名列表"}
  RelatedDomains *int64 `json:"related-domains,omitempty" xml:"related-domains,omitempty" require:"true"`
  // {"en":"The time of the latest update certificate (CST time zone), for example: 2018-08-01 19:00:01","zh_CN":"最新更新证书的时间（CST时区），例如：2018-08-01 19:00:01"}
  CertificateUpdateTime *string `json:"certificate-update-time,omitempty" xml:"certificate-update-time,omitempty" require:"true"`
  // {"en":"Accelerated domain name ID","zh_CN":"加速域名ID"}
  DomainId *int `json:"domain-id,omitempty" xml:"domain-id,omitempty" require:"true"`
  // {"en":"Name of accelerated domain name","zh_CN":"加速域名的名称"}
  DomainName *string `json:"domain-name,omitempty" xml:"domain-name,omitempty" require:"true"`
  // {"en":"Shared, true means shared certificate, false means non-shared certificate","zh_CN":"是否共享，true表示共享证书，false表示非共享证书"}
  ShareSsl *string `json:"share-ssl,omitempty" xml:"share-ssl,omitempty" require:"true"`
  // {"en":"Certificate ca file content md5","zh_CN":"证书ca文件内容md5"}
  CaMd5 *string `json:"ca-md5,omitempty" xml:"ca-md5,omitempty" require:"true"`
  // {"en":"Certificate crt file content md5","zh_CN":"证书crt文件内容md5"}
  CrtMd5 *string `json:"crt-md5,omitempty" xml:"crt-md5,omitempty" require:"true"`
  // {"en":"Certificate key file content md5","zh_CN":"证书key文件内容md5"}
  KeyMd5 *string `json:"key-md5,omitempty" xml:"key-md5,omitempty" require:"true"`
  // {"en":"The CRT certificate serial number","zh_CN":"crt证书序列号"}
  CertificateSerial *string `json:"certificate-serial,omitempty" xml:"certificate-serial,omitempty" require:"true"`
  // {"en":"The CRT certificate issuer","zh_CN":"crt证书颁布者"}
  CertificateIssuer *string `json:"certificate-issuer,omitempty" xml:"certificate-issuer,omitempty" require:"true"`
  // {"en":"Certificate name, unique to customer granularity","zh_CN":"证书名称，客户粒度下是唯一的"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"Remarks on certificate file","zh_CN":"证书文件的备注"}
  Comment *string `json:"comment,omitempty" xml:"comment,omitempty" require:"true"`
  // {"en":"The expiration time of the certificate validity period (CST time zone), for example: 2018-08-01 19:00:00","zh_CN":"证书有效期的到期时间（CST时区），例如：2018-08-01 19:00:00"}
  CertificateValidityTo *string `json:"certificate-validity-to,omitempty" xml:"certificate-validity-to,omitempty" require:"true"`
  // {"en":"The start time of the certificate validity period (CST time zone), for example: 2016-08-01 07:00:00","zh_CN":"证书有效期的起始时间（CST时区），例如：2016-08-01 07:00:00"}
  CertificateValidityFrom *string `json:"certificate-validity-from,omitempty" xml:"certificate-validity-from,omitempty" require:"true"`
  // {"en":"dns-names","zh_CN":"授权域名列表，证书使用者可选名称，父标签"}
  DnsNames *int64 `json:"dns-names,omitempty" xml:"dns-names,omitempty" require:"true"`
  // {"en":"dns-name list, such as\n<dns-names>\n<dns-name>*.example.com</dns-name>\n<dns-name>test.example.com</dns-name>\n</dns-names>","zh_CN":"授权域名，格式如：\n<dns-names>\n<dns-name>*.example.com</dns-name>\n<dns-name>test.example.com</dns-name>\n</dns-names>"}
  DnsName *string `json:"dns-name,omitempty" xml:"dns-name,omitempty" require:"true"`
}

func (s QueryCertificateResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryCertificateResponse) GoString() string {
  return s.String()
}

func (s *QueryCertificateResponse) SetRelatedDomains(v int64) *QueryCertificateResponse {
  s.RelatedDomains = &v
  return s
}

func (s *QueryCertificateResponse) SetCertificateUpdateTime(v string) *QueryCertificateResponse {
  s.CertificateUpdateTime = &v
  return s
}

func (s *QueryCertificateResponse) SetDomainId(v int) *QueryCertificateResponse {
  s.DomainId = &v
  return s
}

func (s *QueryCertificateResponse) SetDomainName(v string) *QueryCertificateResponse {
  s.DomainName = &v
  return s
}

func (s *QueryCertificateResponse) SetShareSsl(v string) *QueryCertificateResponse {
  s.ShareSsl = &v
  return s
}

func (s *QueryCertificateResponse) SetCaMd5(v string) *QueryCertificateResponse {
  s.CaMd5 = &v
  return s
}

func (s *QueryCertificateResponse) SetCrtMd5(v string) *QueryCertificateResponse {
  s.CrtMd5 = &v
  return s
}

func (s *QueryCertificateResponse) SetKeyMd5(v string) *QueryCertificateResponse {
  s.KeyMd5 = &v
  return s
}

func (s *QueryCertificateResponse) SetCertificateSerial(v string) *QueryCertificateResponse {
  s.CertificateSerial = &v
  return s
}

func (s *QueryCertificateResponse) SetCertificateIssuer(v string) *QueryCertificateResponse {
  s.CertificateIssuer = &v
  return s
}

func (s *QueryCertificateResponse) SetName(v string) *QueryCertificateResponse {
  s.Name = &v
  return s
}

func (s *QueryCertificateResponse) SetComment(v string) *QueryCertificateResponse {
  s.Comment = &v
  return s
}

func (s *QueryCertificateResponse) SetCertificateValidityTo(v string) *QueryCertificateResponse {
  s.CertificateValidityTo = &v
  return s
}

func (s *QueryCertificateResponse) SetCertificateValidityFrom(v string) *QueryCertificateResponse {
  s.CertificateValidityFrom = &v
  return s
}

func (s *QueryCertificateResponse) SetDnsNames(v int64) *QueryCertificateResponse {
  s.DnsNames = &v
  return s
}

func (s *QueryCertificateResponse) SetDnsName(v string) *QueryCertificateResponse {
  s.DnsName = &v
  return s
}

type QueryCertificateResponseHeader struct {
}

func (s QueryCertificateResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryCertificateResponseHeader) GoString() string {
  return s.String()
}




type QueryCertificatePaginationRequest struct {
  // {"en":"Page Number","zh_CN":"当前页数"}
  PageNumber *int `json:"page-number,omitempty" xml:"page-number,omitempty"`
  // {"en":"Page Size","zh_CN":"每页记录数"}
  PageSize *int `json:"page-size,omitempty" xml:"page-size,omitempty"`
}

func (s QueryCertificatePaginationRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryCertificatePaginationRequest) GoString() string {
  return s.String()
}

func (s *QueryCertificatePaginationRequest) SetPageNumber(v int) *QueryCertificatePaginationRequest {
  s.PageNumber = &v
  return s
}

func (s *QueryCertificatePaginationRequest) SetPageSize(v int) *QueryCertificatePaginationRequest {
  s.PageSize = &v
  return s
}

type QueryCertificatePaginationRequestHeader struct {
}

func (s QueryCertificatePaginationRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryCertificatePaginationRequestHeader) GoString() string {
  return s.String()
}

type QueryCertificatePaginationPaths struct {
}

func (s QueryCertificatePaginationPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryCertificatePaginationPaths) GoString() string {
  return s.String()
}

type QueryCertificatePaginationParameters struct {
}

func (s QueryCertificatePaginationParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryCertificatePaginationParameters) GoString() string {
  return s.String()
}

type QueryCertificatePaginationResponse struct {
  // {"en":"Certificate list information","zh_CN":"证书列表信息"}
  SslCertificates []*QueryCertificatePaginationResponseSslCertificates `json:"ssl-certificates,omitempty" xml:"ssl-certificates,omitempty" require:"true" type:"Repeated"`
  // {"en":"Page Info","zh_CN":"分页信息"}
  PageInfo *QueryCertificatePaginationResponsePageInfo `json:"page-info,omitempty" xml:"page-info,omitempty" require:"true" type:"Struct"`
}

func (s QueryCertificatePaginationResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryCertificatePaginationResponse) GoString() string {
  return s.String()
}

func (s *QueryCertificatePaginationResponse) SetSslCertificates(v []*QueryCertificatePaginationResponseSslCertificates) *QueryCertificatePaginationResponse {
  s.SslCertificates = v
  return s
}

func (s *QueryCertificatePaginationResponse) SetPageInfo(v *QueryCertificatePaginationResponsePageInfo) *QueryCertificatePaginationResponse {
  s.PageInfo = v
  return s
}

type QueryCertificatePaginationResponseSslCertificates struct     {
  // {"en":"Certificate ID","zh_CN":"证书ID"}
  CertificateId *string `json:"certificate-id,omitempty" xml:"certificate-id,omitempty" require:"true"`
  // {"en":"Certificate name, unique to customer granularity","zh_CN":"证书名称，客户粒度下是唯一的"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"Remarks on certificate file","zh_CN":"证书文件的备注"}
  Comment *string `json:"comment,omitempty" xml:"comment,omitempty" require:"true"`
  // {"en":"Indicates whether the certificate is shared. `true` means shared, `false` means not shared.","zh_CN":"是否共享，true表示共享证书，false表示非共享证书"}
  ShareSsl *string `json:"share-ssl,omitempty" xml:"share-ssl,omitempty" require:"true"`
  // {"en":"Certificate validity start time (CST), such as 2016-08-01 07:00:00","zh_CN":"证书有效期的起始时间（CST时区），例如：2016-08-01 07:00:00"}
  CertificateValidityFrom *string `json:"certificate-validity-from,omitempty" xml:"certificate-validity-from,omitempty" require:"true"`
  // {"en":"Certificate validity end time (CST), such as 2018-08-01 19:00:00","zh_CN":"证书有效期的到期时间（CST时区），例如：2018-08-01 19:00:00"}
  CertificateValidityTo *string `json:"certificate-validity-to,omitempty" xml:"certificate-validity-to,omitempty" require:"true"`
  // {"en":"List of domain names using the current certificate","zh_CN":"使用当前证书的域名列表"}
  RelatedDomains []*QueryCertificatePaginationResponseSslCertificatesRelatedDomains `json:"related-domains,omitempty" xml:"related-domains,omitempty" require:"true" type:"Repeated"`
  // {"en":"List of authorized domain names, also known as Subject Alternative Names (SANs).","zh_CN":"授权域名列表，证书使用者可选名称"}
  DnsNames []*string `json:"dns-names,omitempty" xml:"dns-names,omitempty" require:"true" type:"Repeated"`
  // {"en":"The CRT certificate serial number","zh_CN":"crt证书序列号"}
  CertificateSerial *string `json:"certificate-serial,omitempty" xml:"certificate-serial,omitempty" require:"true"`
  // {"en":"The MD5 value of CRT file.","zh_CN":"CRT文件内容的md5值。"}
  CrtMd5 *string `json:"crt-md5,omitempty" xml:"crt-md5,omitempty" require:"true"`
  // {"en":"The MD5 value of KEY file.","zh_CN":"KEY文件内容的md5值。"}
  KeyMd5 *string `json:"key-md5,omitempty" xml:"key-md5,omitempty" require:"true"`
  // {"en":"The MD5 value of CA file.","zh_CN":"CA文件内容的md5值。"}
  CaMd5 *string `json:"ca-md5,omitempty" xml:"ca-md5,omitempty" require:"true"`
  // {"en":"The CRT certificate issuer","zh_CN":"crt证书颁布者"}
  CertificateIssuer *string `json:"certificate-issuer,omitempty" xml:"certificate-issuer,omitempty" require:"true"`
}

func (s QueryCertificatePaginationResponseSslCertificates) String() string {
  return tea.Prettify(s)
}

func (s QueryCertificatePaginationResponseSslCertificates) GoString() string {
  return s.String()
}

func (s *QueryCertificatePaginationResponseSslCertificates) SetCertificateId(v string) *QueryCertificatePaginationResponseSslCertificates {
  s.CertificateId = &v
  return s
}

func (s *QueryCertificatePaginationResponseSslCertificates) SetName(v string) *QueryCertificatePaginationResponseSslCertificates {
  s.Name = &v
  return s
}

func (s *QueryCertificatePaginationResponseSslCertificates) SetComment(v string) *QueryCertificatePaginationResponseSslCertificates {
  s.Comment = &v
  return s
}

func (s *QueryCertificatePaginationResponseSslCertificates) SetShareSsl(v string) *QueryCertificatePaginationResponseSslCertificates {
  s.ShareSsl = &v
  return s
}

func (s *QueryCertificatePaginationResponseSslCertificates) SetCertificateValidityFrom(v string) *QueryCertificatePaginationResponseSslCertificates {
  s.CertificateValidityFrom = &v
  return s
}

func (s *QueryCertificatePaginationResponseSslCertificates) SetCertificateValidityTo(v string) *QueryCertificatePaginationResponseSslCertificates {
  s.CertificateValidityTo = &v
  return s
}

func (s *QueryCertificatePaginationResponseSslCertificates) SetRelatedDomains(v []*QueryCertificatePaginationResponseSslCertificatesRelatedDomains) *QueryCertificatePaginationResponseSslCertificates {
  s.RelatedDomains = v
  return s
}

func (s *QueryCertificatePaginationResponseSslCertificates) SetDnsNames(v []*string) *QueryCertificatePaginationResponseSslCertificates {
  s.DnsNames = v
  return s
}

func (s *QueryCertificatePaginationResponseSslCertificates) SetCertificateSerial(v string) *QueryCertificatePaginationResponseSslCertificates {
  s.CertificateSerial = &v
  return s
}

func (s *QueryCertificatePaginationResponseSslCertificates) SetCrtMd5(v string) *QueryCertificatePaginationResponseSslCertificates {
  s.CrtMd5 = &v
  return s
}

func (s *QueryCertificatePaginationResponseSslCertificates) SetKeyMd5(v string) *QueryCertificatePaginationResponseSslCertificates {
  s.KeyMd5 = &v
  return s
}

func (s *QueryCertificatePaginationResponseSslCertificates) SetCaMd5(v string) *QueryCertificatePaginationResponseSslCertificates {
  s.CaMd5 = &v
  return s
}

func (s *QueryCertificatePaginationResponseSslCertificates) SetCertificateIssuer(v string) *QueryCertificatePaginationResponseSslCertificates {
  s.CertificateIssuer = &v
  return s
}

type QueryCertificatePaginationResponseSslCertificatesRelatedDomains struct     {
  // {"en":"Accelerated domain name ID","zh_CN":"加速域名ID"}
  DomainId *string `json:"domain-id,omitempty" xml:"domain-id,omitempty" require:"true"`
  // {"en":"Name of accelerated domain name","zh_CN":"加速域名的名称"}
  DomainName *string `json:"domain-name,omitempty" xml:"domain-name,omitempty" require:"true"`
}

func (s QueryCertificatePaginationResponseSslCertificatesRelatedDomains) String() string {
  return tea.Prettify(s)
}

func (s QueryCertificatePaginationResponseSslCertificatesRelatedDomains) GoString() string {
  return s.String()
}

func (s *QueryCertificatePaginationResponseSslCertificatesRelatedDomains) SetDomainId(v string) *QueryCertificatePaginationResponseSslCertificatesRelatedDomains {
  s.DomainId = &v
  return s
}

func (s *QueryCertificatePaginationResponseSslCertificatesRelatedDomains) SetDomainName(v string) *QueryCertificatePaginationResponseSslCertificatesRelatedDomains {
  s.DomainName = &v
  return s
}

type QueryCertificatePaginationResponsePageInfo struct {
  // {"en":"Total Number","zh_CN":"总记录数"}
  TotalNumber *int `json:"total-number,omitempty" xml:"total-number,omitempty" require:"true"`
  // {"en":"Total Page Number","zh_CN":"总页数"}
  TotalPageNumber *int `json:"total-page-number,omitempty" xml:"total-page-number,omitempty" require:"true"`
  // {"en":"Page Number","zh_CN":"当前页数"}
  PageNumber *int `json:"page-number,omitempty" xml:"page-number,omitempty" require:"true"`
  // {"en":"Page Size","zh_CN":"每页记录数"}
  PageSize *int `json:"page-size,omitempty" xml:"page-size,omitempty" require:"true"`
}

func (s QueryCertificatePaginationResponsePageInfo) String() string {
  return tea.Prettify(s)
}

func (s QueryCertificatePaginationResponsePageInfo) GoString() string {
  return s.String()
}

func (s *QueryCertificatePaginationResponsePageInfo) SetTotalNumber(v int) *QueryCertificatePaginationResponsePageInfo {
  s.TotalNumber = &v
  return s
}

func (s *QueryCertificatePaginationResponsePageInfo) SetTotalPageNumber(v int) *QueryCertificatePaginationResponsePageInfo {
  s.TotalPageNumber = &v
  return s
}

func (s *QueryCertificatePaginationResponsePageInfo) SetPageNumber(v int) *QueryCertificatePaginationResponsePageInfo {
  s.PageNumber = &v
  return s
}

func (s *QueryCertificatePaginationResponsePageInfo) SetPageSize(v int) *QueryCertificatePaginationResponsePageInfo {
  s.PageSize = &v
  return s
}

type QueryCertificatePaginationResponseHeader struct {
}

func (s QueryCertificatePaginationResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryCertificatePaginationResponseHeader) GoString() string {
  return s.String()
}




type GetCertificateContentInfosRequest struct {
  // {"en":"Certificate ID list","zh_CN":"证书ID列表"}
  CertificateIdList []*string `json:"certificateIdList,omitempty" xml:"certificateIdList,omitempty" require:"true" type:"Repeated"`
}

func (s GetCertificateContentInfosRequest) String() string {
  return tea.Prettify(s)
}

func (s GetCertificateContentInfosRequest) GoString() string {
  return s.String()
}

func (s *GetCertificateContentInfosRequest) SetCertificateIdList(v []*string) *GetCertificateContentInfosRequest {
  s.CertificateIdList = v
  return s
}

type GetCertificateContentInfosRequestHeader struct {
  // {"en":"Date.  Formatting to comply with RFC 1123 specifications as for Thu, 17 May 2012 19:37:58 GMT","zh_CN":"时间戳，格式需符合RFC 1123规范，如Thu, 17 May 2012 19:37:58 GMT"}
  Date *string `json:"date,omitempty" xml:"date,omitempty" require:"true"`
}

func (s GetCertificateContentInfosRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetCertificateContentInfosRequestHeader) GoString() string {
  return s.String()
}

func (s *GetCertificateContentInfosRequestHeader) SetDate(v string) *GetCertificateContentInfosRequestHeader {
  s.Date = &v
  return s
}

type GetCertificateContentInfosPaths struct {
}

func (s GetCertificateContentInfosPaths) String() string {
  return tea.Prettify(s)
}

func (s GetCertificateContentInfosPaths) GoString() string {
  return s.String()
}

type GetCertificateContentInfosParameters struct {
}

func (s GetCertificateContentInfosParameters) String() string {
  return tea.Prettify(s)
}

func (s GetCertificateContentInfosParameters) GoString() string {
  return s.String()
}

type GetCertificateContentInfosResponse struct {
  // {"en":"Certificate id","zh_CN":"证书ID"}
  CertificateId *int64 `json:"certificateId,omitempty" xml:"certificateId,omitempty" require:"true"`
  // {"en":"The certificate name is unique under a customer.\nNote:\n1. Support single domain, multi domain and wildcard domain name.\n2. For example, the authorized domain name of certificate A is.example.com , a.example2.com, b.example2.com test.example.com , a.example2.com, can be associated with certificate A, while the domain name c.example2.com cannot use certificate A.","zh_CN":"证书名称，客户粒度下是唯一的\n\n注意：\n1、支持单域名、多域名和泛域名证书，即证书的授权域名允许为多个域名，或者泛域名\n2、例如：证书A的授权域名为.example.com，a.example2.com，b.example2.com，则域名test.example.com，a.example2.com，都可以关联使用证书A，而域名c.example2.com不能使用证书A"}
  CertificateName *string `json:"certificateName,omitempty" xml:"certificateName,omitempty" require:"true"`
  // {"en":"Certificate content. The encryption algorithm is the same as the creating a certificate.\nEncryption algorithm: First get md5 value of the request header 'Date'. Then take the left 8 bits of MD5 value as the key and the right 8 bits as the IV. Finally, encode the encrypted binary content with Base64. For example:\ndate=`env LANG=\"en_US.UTF-8\" date -u \"+%a, %d %b %Y %H:%M:%S GMT\"`\nmd5str=`echo -n $date | openssl md5`\nkey=`echo -n ${md5str:$((-32)):$((8))}|hexdump -e '8/1 \"%02X\"'`\niv=`echo -n ${md5str:$((-8))}|hexdump -e '8/1 \"%02X\"'`\nNote:\n1.Decryption command: file is the content of the certificate file (CRT, key, CA) responded by the interface\ncat $file |base64 -d|openssl enc -des -K $key -iv $iv -nosalt -d\n2. The date time of the HTTP request header must be the same as that of the above certificate encryption.","zh_CN":"crt文件内容，已加密。加密算法同新增证书的加密算法。\n加密算法说明：将http头中的Date值做md5运算，将md5值的左8位作为key，右8位作为iv，然后对文件内容作des加密，将加密后的二进制内容作base64编码，以下为示例参考：\ndate=`env LANG=\"en_US.UTF-8\" date -u \"+%a, %d %b %Y %H:%M:%S GMT\"`\nmd5str=`echo -n $date | openssl md5`\nkey=`echo -n ${md5str:$((-32)):$((8))}|hexdump -e '8/1 \"%02X\"'`\niv=`echo -n ${md5str:$((-8))}|hexdump -e '8/1 \"%02X\"'`\n注意：\n1、解密命令如下，file就是接口响应的证书文件内容（crt，key，ca）\ncat $file |base64 -d|openssl enc -des -K $key -iv $iv -nosalt -d\n2、http请求头的Date时间必须和上述证书加解密的时间一致"}
  Certificate *string `json:"certificate,omitempty" xml:"certificate,omitempty" require:"true"`
  // {"en":"Certificate key. The encryption algorithm is the same as creating a  new certificate","zh_CN":"key文件内容，已加密。加密算法同新增证书的加密算法。"}
  PrivateKey *string `json:"privateKey,omitempty" xml:"privateKey,omitempty" require:"true"`
  // {"en":"Certificate expiration time","zh_CN":"证书过期时间"}
  ExpirationTime *string `json:"expirationTime,omitempty" xml:"expirationTime,omitempty" require:"true"`
}

func (s GetCertificateContentInfosResponse) String() string {
  return tea.Prettify(s)
}

func (s GetCertificateContentInfosResponse) GoString() string {
  return s.String()
}

func (s *GetCertificateContentInfosResponse) SetCertificateId(v int64) *GetCertificateContentInfosResponse {
  s.CertificateId = &v
  return s
}

func (s *GetCertificateContentInfosResponse) SetCertificateName(v string) *GetCertificateContentInfosResponse {
  s.CertificateName = &v
  return s
}

func (s *GetCertificateContentInfosResponse) SetCertificate(v string) *GetCertificateContentInfosResponse {
  s.Certificate = &v
  return s
}

func (s *GetCertificateContentInfosResponse) SetPrivateKey(v string) *GetCertificateContentInfosResponse {
  s.PrivateKey = &v
  return s
}

func (s *GetCertificateContentInfosResponse) SetExpirationTime(v string) *GetCertificateContentInfosResponse {
  s.ExpirationTime = &v
  return s
}

type GetCertificateContentInfosResponseHeader struct {
}

func (s GetCertificateContentInfosResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetCertificateContentInfosResponseHeader) GoString() string {
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

type GetCertificateContentRequestHeader struct {
  // {"en":"Date.  Formatting to comply with RFC 1123 specifications as for Thu, 17 May 2012 19:37:58 GMT","zh_CN":"时间戳，格式需符合RFC 1123规范，如Thu, 17 May 2012 19:37:58 GMT"}
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

type GetCertificateContentPaths struct {
  // {"en":"Certificate ID, corresponding to the* in the interface address","zh_CN":"证书ID，对应接口url中的*\n注意：参看请求示例中的url，100166对应的就是certificate-id"}
  CertificateId *int `json:"certificate-id,omitempty" xml:"certificate-id,omitempty" require:"true"`
}

func (s GetCertificateContentPaths) String() string {
  return tea.Prettify(s)
}

func (s GetCertificateContentPaths) GoString() string {
  return s.String()
}

func (s *GetCertificateContentPaths) SetCertificateId(v int) *GetCertificateContentPaths {
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

type GetCertificateContentResponse struct {
  // {"en":"certificate id","zh_CN":"证书ID"}
  CertificateId *int64 `json:"certificate-id,omitempty" xml:"certificate-id,omitempty" require:"true"`
  // {"en":"The certificate name is unique under a customer.\nNote:\n1. Support single domain, multi domain and wildcard domain name.\n2. For example, the authorized domain name of certificate A is.example.com , a.example2.com, b.example2.com test.example.com , a.example2.com, can be associated with certificate A, while the domain name c.example2.com cannot use certificate A.","zh_CN":"证书名称，客户粒度下是唯一的\n\n注意：\n1、支持单域名、多域名和泛域名证书，即证书的授权域名允许为多个域名，或者泛域名\n2、例如：证书A的授权域名为.example.com，a.example2.com，b.example2.com，则域名test.example.com，a.example2.com，都可以关联使用证书A，而域名c.example2.com不能使用证书A"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"Comment for a certificate.","zh_CN":"证书文件的备注"}
  Comment *string `json:"comment,omitempty" xml:"comment,omitempty" require:"true"`
  // {"en":"Specified a shared certificate. True means shared.","zh_CN":"是否共享，true表示共享证书，false表示非共享证书"}
  ShareSsl *bool `json:"share-ssl,omitempty" xml:"share-ssl,omitempty" require:"true"`
  // {"en":"Certificate content. The encryption algorithm is the same as the creating a certificate.\nEncryption algorithm: First get md5 value of the request header 'Date'. Then take the left 8 bits of MD5 value as the key and the right 8 bits as the IV. Finally, encode the encrypted binary content with Base64. For example:\ndate=`env LANG=\"en_US.UTF-8\" date -u \"+%a, %d %b %Y %H:%M:%S GMT\"`\nmd5str=`echo -n $date | openssl md5`\nkey=`echo -n ${md5str:$((-32)):$((8))}|hexdump -e '8/1 \"%02X\"'`\niv=`echo -n ${md5str:$((-8))}|hexdump -e '8/1 \"%02X\"'`\nNote:\n1.Decryption command: file is the content of the certificate file (CRT, key, CA) responded by the interface\ncat $file |base64 -d|openssl enc -des -K $key -iv $iv -nosalt -d\n2. The date time of the HTTP request header must be the same as that of the above certificate encryption.","zh_CN":"crt文件内容，已加密。加密算法同新增证书的加密算法。\n加密算法说明：将http头中的Date值做md5运算，将md5值的左8位作为key，右8位作为iv，然后对文件内容作des加密，将加密后的二进制内容作base64编码，以下为示例参考：\ndate=`env LANG=\"en_US.UTF-8\" date -u \"+%a, %d %b %Y %H:%M:%S GMT\"`\nmd5str=`echo -n $date | openssl md5`\nkey=`echo -n ${md5str:$((-32)):$((8))}|hexdump -e '8/1 \"%02X\"'`\niv=`echo -n ${md5str:$((-8))}|hexdump -e '8/1 \"%02X\"'`\n注意：\n1、解密命令如下，file就是接口响应的证书文件内容（crt，key，ca）\ncat $file |base64 -d|openssl enc -des -K $key -iv $iv -nosalt -d\n2、http请求头的Date时间必须和上述证书加解密的时间一致"}
  SslCertificate *string `json:"ssl-certificate,omitempty" xml:"ssl-certificate,omitempty" require:"true"`
  // {"en":"Certificate key. The encryption algorithm is the same as creating a  new certificate","zh_CN":"key文件内容，已加密。加密算法同新增证书的加密算法。"}
  SslKey *string `json:"ssl-key,omitempty" xml:"ssl-key,omitempty" require:"true"`
  // {"en":"Certificate chain. The encryption algorithm is the same as the new certificate.","zh_CN":"ca文件内容，已加密。加密方式同新增证书的加密方式。"}
  SslCertificateChain *string `json:"ssl-certificate-chain,omitempty" xml:"ssl-certificate-chain,omitempty" require:"true"`
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

type QueryDomainMultiCertConfigRequestHeader struct {
}

func (s QueryDomainMultiCertConfigRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryDomainMultiCertConfigRequestHeader) GoString() string {
  return s.String()
}

type QueryDomainMultiCertConfigPaths struct {
  // {"en":"The domain or domainId which needs a query config.","zh_CN":"需要查询配置的域名或域名id"}
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

type QueryDomainMultiCertConfigResponse struct {
  // {"en":"code","zh_CN":"错误码，成功为0"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"message","zh_CN":"结果信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"data","zh_CN":"返回数据"}
  Data *QueryDomainMultiCertConfigResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
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

func (s *QueryDomainMultiCertConfigResponse) SetData(v *QueryDomainMultiCertConfigResponseData) *QueryDomainMultiCertConfigResponse {
  s.Data = v
  return s
}

type QueryDomainMultiCertConfigResponseData struct {
  // {"en":"domainId","zh_CN":"域名ID"}
  DomainId *string `json:"domainId,omitempty" xml:"domainId,omitempty" require:"true"`
  // {"en":"domainName","zh_CN":"域名名称"}
  DomainName *string `json:"domainName,omitempty" xml:"domainName,omitempty" require:"true"`
  // {"en":"TLSVersion","zh_CN":"TLS版本"}
  TLSVersion *string `json:"TLSVersion,omitempty" xml:"TLSVersion,omitempty" require:"true"`
  // {"en":"enableOCSP","zh_CN":"启用OCSP，默认不启用，可选值true、false。没配置返回"}
  EnableOCSP *string `json:"enableOCSP,omitempty" xml:"enableOCSP,omitempty" require:"true"`
  // {"en":"cipherSuites","zh_CN":"设置cipher加密套件"}
  CipherSuites *string `json:"cipherSuites,omitempty" xml:"cipherSuites,omitempty" require:"true"`
  // {"en":"domainCertConfigs","zh_CN":"证书配置"}
  DomainCertConfigs []*QueryDomainMultiCertConfigResponseDataDomainCertConfigs `json:"domainCertConfigs,omitempty" xml:"domainCertConfigs,omitempty" require:"true" type:"Repeated"`
}

func (s QueryDomainMultiCertConfigResponseData) String() string {
  return tea.Prettify(s)
}

func (s QueryDomainMultiCertConfigResponseData) GoString() string {
  return s.String()
}

func (s *QueryDomainMultiCertConfigResponseData) SetDomainId(v string) *QueryDomainMultiCertConfigResponseData {
  s.DomainId = &v
  return s
}

func (s *QueryDomainMultiCertConfigResponseData) SetDomainName(v string) *QueryDomainMultiCertConfigResponseData {
  s.DomainName = &v
  return s
}

func (s *QueryDomainMultiCertConfigResponseData) SetTLSVersion(v string) *QueryDomainMultiCertConfigResponseData {
  s.TLSVersion = &v
  return s
}

func (s *QueryDomainMultiCertConfigResponseData) SetEnableOCSP(v string) *QueryDomainMultiCertConfigResponseData {
  s.EnableOCSP = &v
  return s
}

func (s *QueryDomainMultiCertConfigResponseData) SetCipherSuites(v string) *QueryDomainMultiCertConfigResponseData {
  s.CipherSuites = &v
  return s
}

func (s *QueryDomainMultiCertConfigResponseData) SetDomainCertConfigs(v []*QueryDomainMultiCertConfigResponseDataDomainCertConfigs) *QueryDomainMultiCertConfigResponseData {
  s.DomainCertConfigs = v
  return s
}

type QueryDomainMultiCertConfigResponseDataDomainCertConfigs struct     {
  // {"en":"certUsage","zh_CN":"证书用途"}
  CertUsage *string `json:"certUsage,omitempty" xml:"certUsage,omitempty" require:"true"`
  // {"en":"certInfos","zh_CN":"证书信息"}
  CertInfos []*QueryDomainMultiCertConfigResponseDataDomainCertConfigsCertInfos `json:"certInfos,omitempty" xml:"certInfos,omitempty" require:"true" type:"Repeated"`
}

func (s QueryDomainMultiCertConfigResponseDataDomainCertConfigs) String() string {
  return tea.Prettify(s)
}

func (s QueryDomainMultiCertConfigResponseDataDomainCertConfigs) GoString() string {
  return s.String()
}

func (s *QueryDomainMultiCertConfigResponseDataDomainCertConfigs) SetCertUsage(v string) *QueryDomainMultiCertConfigResponseDataDomainCertConfigs {
  s.CertUsage = &v
  return s
}

func (s *QueryDomainMultiCertConfigResponseDataDomainCertConfigs) SetCertInfos(v []*QueryDomainMultiCertConfigResponseDataDomainCertConfigsCertInfos) *QueryDomainMultiCertConfigResponseDataDomainCertConfigs {
  s.CertInfos = v
  return s
}

type QueryDomainMultiCertConfigResponseDataDomainCertConfigsCertInfos struct     {
  // {"en":"certificateId","zh_CN":"证书ID"}
  CertificateId *int64 `json:"certificateId,omitempty" xml:"certificateId,omitempty" require:"true"`
  // {"en":"certificateName","zh_CN":"证书名称"}
  CertificateName *string `json:"certificateName,omitempty" xml:"certificateName,omitempty" require:"true"`
  // {"en":"comment","zh_CN":"备注"}
  Comment *string `json:"comment,omitempty" xml:"comment,omitempty" require:"true"`
  // {"en":"serial","zh_CN":"序列号"}
  Serial *string `json:"serial,omitempty" xml:"serial,omitempty" require:"true"`
}

func (s QueryDomainMultiCertConfigResponseDataDomainCertConfigsCertInfos) String() string {
  return tea.Prettify(s)
}

func (s QueryDomainMultiCertConfigResponseDataDomainCertConfigsCertInfos) GoString() string {
  return s.String()
}

func (s *QueryDomainMultiCertConfigResponseDataDomainCertConfigsCertInfos) SetCertificateId(v int64) *QueryDomainMultiCertConfigResponseDataDomainCertConfigsCertInfos {
  s.CertificateId = &v
  return s
}

func (s *QueryDomainMultiCertConfigResponseDataDomainCertConfigsCertInfos) SetCertificateName(v string) *QueryDomainMultiCertConfigResponseDataDomainCertConfigsCertInfos {
  s.CertificateName = &v
  return s
}

func (s *QueryDomainMultiCertConfigResponseDataDomainCertConfigsCertInfos) SetComment(v string) *QueryDomainMultiCertConfigResponseDataDomainCertConfigsCertInfos {
  s.Comment = &v
  return s
}

func (s *QueryDomainMultiCertConfigResponseDataDomainCertConfigsCertInfos) SetSerial(v string) *QueryDomainMultiCertConfigResponseDataDomainCertConfigsCertInfos {
  s.Serial = &v
  return s
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




type GetCertificateAssociatedHostnamesRequest struct {
}

func (s GetCertificateAssociatedHostnamesRequest) String() string {
  return tea.Prettify(s)
}

func (s GetCertificateAssociatedHostnamesRequest) GoString() string {
  return s.String()
}

type GetCertificateAssociatedHostnamesResponse struct {
  // {"en":"Response code, 0 means successful.", "zh_CN":"接口响应code，0代表成功。"}
  Code *int `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response error message if failed.", "zh_CN":"接口响应信息，success代表成功，失败则提供失败信息。"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Response data array.", "zh_CN":"接口响应数据"}
  Data *GetCertificateAssociatedHostnamesResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
}

func (s GetCertificateAssociatedHostnamesResponse) String() string {
  return tea.Prettify(s)
}

func (s GetCertificateAssociatedHostnamesResponse) GoString() string {
  return s.String()
}

func (s *GetCertificateAssociatedHostnamesResponse) SetCode(v int) *GetCertificateAssociatedHostnamesResponse {
  s.Code = &v
  return s
}

func (s *GetCertificateAssociatedHostnamesResponse) SetMessage(v string) *GetCertificateAssociatedHostnamesResponse {
  s.Message = &v
  return s
}

func (s *GetCertificateAssociatedHostnamesResponse) SetData(v *GetCertificateAssociatedHostnamesResponseData) *GetCertificateAssociatedHostnamesResponse {
  s.Data = v
  return s
}

type GetCertificateAssociatedHostnamesResponseData struct {
  // {"en":"Certificate ID", "zh_CN":"证书ID"}
  CertificateId *int `json:"certificateId,omitempty" xml:"certificateId,omitempty" require:"true"`
  // {"en":"Certificate related domains", "zh_CN":"在用该证书的加速域名列表"}
  Domains []*GetCertificateAssociatedHostnamesResponseDataDomains `json:"domains,omitempty" xml:"domains,omitempty" require:"true" type:"Repeated"`
}

func (s GetCertificateAssociatedHostnamesResponseData) String() string {
  return tea.Prettify(s)
}

func (s GetCertificateAssociatedHostnamesResponseData) GoString() string {
  return s.String()
}

func (s *GetCertificateAssociatedHostnamesResponseData) SetCertificateId(v int) *GetCertificateAssociatedHostnamesResponseData {
  s.CertificateId = &v
  return s
}

func (s *GetCertificateAssociatedHostnamesResponseData) SetDomains(v []*GetCertificateAssociatedHostnamesResponseDataDomains) *GetCertificateAssociatedHostnamesResponseData {
  s.Domains = v
  return s
}

type GetCertificateAssociatedHostnamesResponseDataDomains struct     {
  // {"en":"Domain ID", "zh_CN":"域名id"}
  DomainId *int `json:"domainId,omitempty" xml:"domainId,omitempty" require:"true"`
  // {"en":"Domain name", "zh_CN":"域名名称"}
  DomainName *string `json:"domainName,omitempty" xml:"domainName,omitempty" require:"true"`
}

func (s GetCertificateAssociatedHostnamesResponseDataDomains) String() string {
  return tea.Prettify(s)
}

func (s GetCertificateAssociatedHostnamesResponseDataDomains) GoString() string {
  return s.String()
}

func (s *GetCertificateAssociatedHostnamesResponseDataDomains) SetDomainId(v int) *GetCertificateAssociatedHostnamesResponseDataDomains {
  s.DomainId = &v
  return s
}

func (s *GetCertificateAssociatedHostnamesResponseDataDomains) SetDomainName(v string) *GetCertificateAssociatedHostnamesResponseDataDomains {
  s.DomainName = &v
  return s
}

type GetCertificateAssociatedHostnamesPaths struct {
  // {"en":"The certificate ID you want to query.", "zh_CN":"需要查询的证书id"}
  CertificateId *int `json:"certificateId,omitempty" xml:"certificateId,omitempty" require:"true"`
}

func (s GetCertificateAssociatedHostnamesPaths) String() string {
  return tea.Prettify(s)
}

func (s GetCertificateAssociatedHostnamesPaths) GoString() string {
  return s.String()
}

func (s *GetCertificateAssociatedHostnamesPaths) SetCertificateId(v int) *GetCertificateAssociatedHostnamesPaths {
  s.CertificateId = &v
  return s
}

type GetCertificateAssociatedHostnamesParameters struct {
}

func (s GetCertificateAssociatedHostnamesParameters) String() string {
  return tea.Prettify(s)
}

func (s GetCertificateAssociatedHostnamesParameters) GoString() string {
  return s.String()
}

type GetCertificateAssociatedHostnamesRequestHeader struct {
}

func (s GetCertificateAssociatedHostnamesRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetCertificateAssociatedHostnamesRequestHeader) GoString() string {
  return s.String()
}

type GetCertificateAssociatedHostnamesResponseHeader struct {
}

func (s GetCertificateAssociatedHostnamesResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetCertificateAssociatedHostnamesResponseHeader) GoString() string {
  return s.String()
}




type GetCertificateV2Request struct {
}

func (s GetCertificateV2Request) String() string {
  return tea.Prettify(s)
}

func (s GetCertificateV2Request) GoString() string {
  return s.String()
}

type GetCertificateV2RequestHeader struct {
}

func (s GetCertificateV2RequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetCertificateV2RequestHeader) GoString() string {
  return s.String()
}

type GetCertificateV2Paths struct {
  // {"en":"The certificate ID you want to query.","zh_CN":"需要查询的证书id"}
  CertificateId *int `json:"certificateId,omitempty" xml:"certificateId,omitempty" require:"true"`
}

func (s GetCertificateV2Paths) String() string {
  return tea.Prettify(s)
}

func (s GetCertificateV2Paths) GoString() string {
  return s.String()
}

func (s *GetCertificateV2Paths) SetCertificateId(v int) *GetCertificateV2Paths {
  s.CertificateId = &v
  return s
}

type GetCertificateV2Parameters struct {
}

func (s GetCertificateV2Parameters) String() string {
  return tea.Prettify(s)
}

func (s GetCertificateV2Parameters) GoString() string {
  return s.String()
}

type GetCertificateV2Response struct {
  // {"en":"Response code, 0 means successful.","zh_CN":"接口响应code，0代表成功。"}
  Code *int `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response error message if failed.","zh_CN":"接口响应信息，success代表成功，失败则提供失败信息。"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Response data array.","zh_CN":"接口响应数据"}
  Data *GetCertificateV2ResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
}

func (s GetCertificateV2Response) String() string {
  return tea.Prettify(s)
}

func (s GetCertificateV2Response) GoString() string {
  return s.String()
}

func (s *GetCertificateV2Response) SetCode(v int) *GetCertificateV2Response {
  s.Code = &v
  return s
}

func (s *GetCertificateV2Response) SetMessage(v string) *GetCertificateV2Response {
  s.Message = &v
  return s
}

func (s *GetCertificateV2Response) SetData(v *GetCertificateV2ResponseData) *GetCertificateV2Response {
  s.Data = v
  return s
}

type GetCertificateV2ResponseData struct {
  // {"en":"Certificate ID","zh_CN":"证书ID。"}
  CertificateId *int `json:"certificateId,omitempty" xml:"certificateId,omitempty" require:"true"`
  // {"en":"Certificate name","zh_CN":"证书名称。"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"Comment","zh_CN":"备注信息。"}
  Comment *string `json:"comment,omitempty" xml:"comment,omitempty" require:"true"`
  // {"en":"Certificate serial","zh_CN":"证书序列号。"}
  Serial *string `json:"serial,omitempty" xml:"serial,omitempty" require:"true"`
  // {"en":"Issue Date","zh_CN":"签发时间。"}
  NotBefore *string `json:"notBefore,omitempty" xml:"notBefore,omitempty" require:"true"`
  // {"en":"Expiration Date","zh_CN":"到期时间。"}
  NotAfter *string `json:"notAfter,omitempty" xml:"notAfter,omitempty" require:"true"`
  // {"en":"Common name","zh_CN":"证书绑定的主域名。"}
  CommonName *string `json:"commonName,omitempty" xml:"commonName,omitempty" require:"true"`
  // {"en":"Subject Alternative Names","zh_CN":"证书绑定的附加域名列表。"}
  SubjectAlternativeNames []*GetCertificateV2ResponseDataSubjectAlternativeNames `json:"subjectAlternativeNames,omitempty" xml:"subjectAlternativeNames,omitempty" require:"true" type:"Repeated"`
}

func (s GetCertificateV2ResponseData) String() string {
  return tea.Prettify(s)
}

func (s GetCertificateV2ResponseData) GoString() string {
  return s.String()
}

func (s *GetCertificateV2ResponseData) SetCertificateId(v int) *GetCertificateV2ResponseData {
  s.CertificateId = &v
  return s
}

func (s *GetCertificateV2ResponseData) SetName(v string) *GetCertificateV2ResponseData {
  s.Name = &v
  return s
}

func (s *GetCertificateV2ResponseData) SetComment(v string) *GetCertificateV2ResponseData {
  s.Comment = &v
  return s
}

func (s *GetCertificateV2ResponseData) SetSerial(v string) *GetCertificateV2ResponseData {
  s.Serial = &v
  return s
}

func (s *GetCertificateV2ResponseData) SetNotBefore(v string) *GetCertificateV2ResponseData {
  s.NotBefore = &v
  return s
}

func (s *GetCertificateV2ResponseData) SetNotAfter(v string) *GetCertificateV2ResponseData {
  s.NotAfter = &v
  return s
}

func (s *GetCertificateV2ResponseData) SetCommonName(v string) *GetCertificateV2ResponseData {
  s.CommonName = &v
  return s
}

func (s *GetCertificateV2ResponseData) SetSubjectAlternativeNames(v []*GetCertificateV2ResponseDataSubjectAlternativeNames) *GetCertificateV2ResponseData {
  s.SubjectAlternativeNames = v
  return s
}

type GetCertificateV2ResponseDataSubjectAlternativeNames struct     {
  // {"en":"Subject Alternative Name","zh_CN":"证书绑定的附加域名"}
  SubjectAlternativeName *string `json:"subjectAlternativeName,omitempty" xml:"subjectAlternativeName,omitempty" require:"true"`
}

func (s GetCertificateV2ResponseDataSubjectAlternativeNames) String() string {
  return tea.Prettify(s)
}

func (s GetCertificateV2ResponseDataSubjectAlternativeNames) GoString() string {
  return s.String()
}

func (s *GetCertificateV2ResponseDataSubjectAlternativeNames) SetSubjectAlternativeName(v string) *GetCertificateV2ResponseDataSubjectAlternativeNames {
  s.SubjectAlternativeName = &v
  return s
}

type GetCertificateV2ResponseHeader struct {
}

func (s GetCertificateV2ResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetCertificateV2ResponseHeader) GoString() string {
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

type QueryCertificateListRequestHeader struct {
}

func (s QueryCertificateListRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryCertificateListRequestHeader) GoString() string {
  return s.String()
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

type QueryCertificateListResponse struct {
  // {"en":"Certificate list information","zh_CN":"证书列表信息"}
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
  // {"en":"Certificate ID","zh_CN":"证书ID"}
  CertificateId *string `json:"certificate-id,omitempty" xml:"certificate-id,omitempty" require:"true"`
  // {"en":"Certificate name, unique to customer granularity","zh_CN":"证书名称，客户粒度下是唯一的"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"Remarks on cerfiticate file","zh_CN":"证书文件的备注"}
  Comment *string `json:"comment,omitempty" xml:"comment,omitempty" require:"true"`
  // {"en":"Shared, optional values are true and false, true represents shared certificates, false represents unshared certificates, default is false\nThis certificate allows cross-customer use when share-ssl is true. (The API does not support cross-customer use certificates. Contact customer service for manual configuration if required.)","zh_CN":"是否共享，true表示共享证书，false表示非共享证书"}
  ShareSsl *string `json:"share-ssl,omitempty" xml:"share-ssl,omitempty" require:"true"`
  // {"en":"Certificate validity start time (CST), such as 2016-08-01 07:00:00","zh_CN":"证书有效期的起始时间（CST时区），例如：2016-08-01 07:00:00"}
  CertificateValidityFrom *string `json:"certificate-validity-from,omitempty" xml:"certificate-validity-from,omitempty" require:"true"`
  // {"en":"Certificate validity end time (CST), such as 2018-08-01 19:00:00","zh_CN":"证书有效期的到期时间（CST时区），例如：2018-08-01 19:00:00"}
  CertificateValidityTo *string `json:"certificate-validity-to,omitempty" xml:"certificate-validity-to,omitempty" require:"true"`
  // {"en":"List of domain names using the current certificate","zh_CN":"使用当前证书的域名列表"}
  RelatedDomains []*QueryCertificateListResponseSslCertificatesRelatedDomains `json:"related-domains,omitempty" xml:"related-domains,omitempty" require:"true" type:"Repeated"`
  // {"en":"dns-names","zh_CN":"授权域名列表，证书使用者可选名称，父标签"}
  DnsNames []*string `json:"dns-names,omitempty" xml:"dns-names,omitempty" require:"true" type:"Repeated"`
  // {"en":"The CRT certificate serial number","zh_CN":"crt证书序列号"}
  CertificateSerial *string `json:"certificate-serial,omitempty" xml:"certificate-serial,omitempty" require:"true"`
  // {"en":"The MD5 value of CRT file.","zh_CN":"CRT文件内容的md5值。"}
  CrtMd5 *string `json:"crt-md5,omitempty" xml:"crt-md5,omitempty" require:"true"`
  // {"en":"The MD5 value of KEY file.","zh_CN":"KEY文件内容的md5值。"}
  KeyMd5 *string `json:"key-md5,omitempty" xml:"key-md5,omitempty" require:"true"`
  // {"en":"The MD5 value of CA file.","zh_CN":"CA。"}
  CaMd5 *string `json:"ca-md5,omitempty" xml:"ca-md5,omitempty" require:"true"`
  // {"en":"The CRT certificate issuer","zh_CN":"crt证书颁布者"}
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
  // {"en":"Accelerated domain name ID","zh_CN":"加速域名ID"}
  DomainId *string `json:"domain-id,omitempty" xml:"domain-id,omitempty" require:"true"`
  // {"en":"Name of accelerated domain name","zh_CN":"加速域名的名称"}
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

type QueryCertificateListResponseHeader struct {
}

func (s QueryCertificateListResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryCertificateListResponseHeader) GoString() string {
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

type GetCertificateContentRequestHeader struct {
}

func (s GetCertificateContentRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetCertificateContentRequestHeader) GoString() string {
  return s.String()
}

type GetCertificateContentPaths struct {
  // {"en":"The certificate ID you want to query.","zh_CN":"需要查询的证书id"}
  CertificateId *int `json:"certificateId,omitempty" xml:"certificateId,omitempty" require:"true"`
}

func (s GetCertificateContentPaths) String() string {
  return tea.Prettify(s)
}

func (s GetCertificateContentPaths) GoString() string {
  return s.String()
}

func (s *GetCertificateContentPaths) SetCertificateId(v int) *GetCertificateContentPaths {
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

type GetCertificateContentResponse struct {
  // {"en":"Response code, 0 means successful.","zh_CN":"接口响应code，0代表成功。"}
  Code *int `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response error message if failed.","zh_CN":"接口响应信息，success代表成功，失败则提供失败信息。"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Response data array.","zh_CN":"接口响应数据"}
  Data *GetCertificateContentResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
}

func (s GetCertificateContentResponse) String() string {
  return tea.Prettify(s)
}

func (s GetCertificateContentResponse) GoString() string {
  return s.String()
}

func (s *GetCertificateContentResponse) SetCode(v int) *GetCertificateContentResponse {
  s.Code = &v
  return s
}

func (s *GetCertificateContentResponse) SetMessage(v string) *GetCertificateContentResponse {
  s.Message = &v
  return s
}

func (s *GetCertificateContentResponse) SetData(v *GetCertificateContentResponseData) *GetCertificateContentResponse {
  s.Data = v
  return s
}

type GetCertificateContentResponseData struct {
  // {"en":"Certificate ID","zh_CN":"证书ID"}
  CertificateId *int `json:"certificateId,omitempty" xml:"certificateId,omitempty" require:"true"`
  // {"en":"Certificate content","zh_CN":"证书内容，PEM格式"}
  Certificate *string `json:"certificate,omitempty" xml:"certificate,omitempty" require:"true"`
}

func (s GetCertificateContentResponseData) String() string {
  return tea.Prettify(s)
}

func (s GetCertificateContentResponseData) GoString() string {
  return s.String()
}

func (s *GetCertificateContentResponseData) SetCertificateId(v int) *GetCertificateContentResponseData {
  s.CertificateId = &v
  return s
}

func (s *GetCertificateContentResponseData) SetCertificate(v string) *GetCertificateContentResponseData {
  s.Certificate = &v
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




type UpdateCertificateV2Request struct {
  // {"en":"Certificate name","zh_CN":"证书名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"Certificate, PEM certificate.","zh_CN":"证书内容，PEM格式。例如：\n-----BEGIN CERTIFICATE-----\n……\n-----END CERTIFICATE-----"}
  Certificate *string `json:"certificate,omitempty" xml:"certificate,omitempty"`
  // {"en":"Private key of the certificate, PEM certificate. Not required when you specify a csrId.","zh_CN":"证书密钥，PEM格式。例如：\n-----BEGIN RSA PRIVATE KEY-----\n……\n-----BEGIN RSA PRIVATE KEY-----"}
  PrivateKey *string `json:"privateKey,omitempty" xml:"privateKey,omitempty"`
  // {"en":"Automatically renew","zh_CN":"是否自动续签","exampleValue":"true,false"}
  AutoRenew *string `json:"autoRenew,omitempty" xml:"autoRenew,omitempty"`
  // {"en":"Do you want to send an expiration alarm","zh_CN":"是否需要证书过期告警","exampleValue":"true,false"}
  IsNeedAlarm *string `json:"isNeedAlarm,omitempty" xml:"isNeedAlarm,omitempty"`
  // {"en":"Csr ID, the id returned by the interface \"Create CSR","zh_CN":"csrId，证书申请文件的id"}
  CsrId *int `json:"csrId,omitempty" xml:"csrId,omitempty"`
  // {"en":"Comment","zh_CN":"备注信息"}
  Comment *string `json:"comment,omitempty" xml:"comment,omitempty"`
}

func (s UpdateCertificateV2Request) String() string {
  return tea.Prettify(s)
}

func (s UpdateCertificateV2Request) GoString() string {
  return s.String()
}

func (s *UpdateCertificateV2Request) SetName(v string) *UpdateCertificateV2Request {
  s.Name = &v
  return s
}

func (s *UpdateCertificateV2Request) SetCertificate(v string) *UpdateCertificateV2Request {
  s.Certificate = &v
  return s
}

func (s *UpdateCertificateV2Request) SetPrivateKey(v string) *UpdateCertificateV2Request {
  s.PrivateKey = &v
  return s
}

func (s *UpdateCertificateV2Request) SetAutoRenew(v string) *UpdateCertificateV2Request {
  s.AutoRenew = &v
  return s
}

func (s *UpdateCertificateV2Request) SetIsNeedAlarm(v string) *UpdateCertificateV2Request {
  s.IsNeedAlarm = &v
  return s
}

func (s *UpdateCertificateV2Request) SetCsrId(v int) *UpdateCertificateV2Request {
  s.CsrId = &v
  return s
}

func (s *UpdateCertificateV2Request) SetComment(v string) *UpdateCertificateV2Request {
  s.Comment = &v
  return s
}

type UpdateCertificateV2RequestHeader struct {
}

func (s UpdateCertificateV2RequestHeader) String() string {
  return tea.Prettify(s)
}

func (s UpdateCertificateV2RequestHeader) GoString() string {
  return s.String()
}

type UpdateCertificateV2Paths struct {
  // {"en":"The certificate ID you want to modify.","zh_CN":"需要修改的证书id"}
  CertificateId *int `json:"certificateId,omitempty" xml:"certificateId,omitempty" require:"true"`
}

func (s UpdateCertificateV2Paths) String() string {
  return tea.Prettify(s)
}

func (s UpdateCertificateV2Paths) GoString() string {
  return s.String()
}

func (s *UpdateCertificateV2Paths) SetCertificateId(v int) *UpdateCertificateV2Paths {
  s.CertificateId = &v
  return s
}

type UpdateCertificateV2Parameters struct {
}

func (s UpdateCertificateV2Parameters) String() string {
  return tea.Prettify(s)
}

func (s UpdateCertificateV2Parameters) GoString() string {
  return s.String()
}

type UpdateCertificateV2Response struct {
  // {"en":"Response code, 0 means successful.","zh_CN":"接口响应code，0代表成功。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response error message if failed.","zh_CN":"接口响应信息，success代表成功，失败则提供失败信息。"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Response data array.","zh_CN":"接口响应数据"}
  Data *string `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s UpdateCertificateV2Response) String() string {
  return tea.Prettify(s)
}

func (s UpdateCertificateV2Response) GoString() string {
  return s.String()
}

func (s *UpdateCertificateV2Response) SetCode(v string) *UpdateCertificateV2Response {
  s.Code = &v
  return s
}

func (s *UpdateCertificateV2Response) SetMessage(v string) *UpdateCertificateV2Response {
  s.Message = &v
  return s
}

func (s *UpdateCertificateV2Response) SetData(v string) *UpdateCertificateV2Response {
  s.Data = &v
  return s
}

type UpdateCertificateV2ResponseHeader struct {
}

func (s UpdateCertificateV2ResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s UpdateCertificateV2ResponseHeader) GoString() string {
  return s.String()
}




type CreateCertificateV2Request struct {
  // {"en":"Certificate name","zh_CN":"证书名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"Certificate, PEM certificate, including CRT file and CA file.","zh_CN":"证书内容，PEM格式，包含CRT文件、CA文件。例如：\n-----BEGIN CERTIFICATE-----\n……\n-----END CERTIFICATE-----"}
  Certificate *string `json:"certificate,omitempty" xml:"certificate,omitempty" require:"true"`
  // {"en":"Private key of the certificate, PEM certificate. Not required when you specify a csrId.","zh_CN":"证书密钥，PEM格式。例如：\n-----BEGIN RSA PRIVATE KEY-----\n……\n-----BEGIN RSA PRIVATE KEY-----\n当指定csrId时，无需上传证书密钥。"}
  PrivateKey *string `json:"privateKey,omitempty" xml:"privateKey,omitempty"`
  // {"en":"Csr ID, the id returned by the interface \"Create CSR\".","zh_CN":"csrId，证书申请文件的id。"}
  CsrId *int `json:"csrId,omitempty" xml:"csrId,omitempty"`
  // {"en":"comment","zh_CN":"备注"}
  Comment *string `json:"comment,omitempty" xml:"comment,omitempty"`
}

func (s CreateCertificateV2Request) String() string {
  return tea.Prettify(s)
}

func (s CreateCertificateV2Request) GoString() string {
  return s.String()
}

func (s *CreateCertificateV2Request) SetName(v string) *CreateCertificateV2Request {
  s.Name = &v
  return s
}

func (s *CreateCertificateV2Request) SetCertificate(v string) *CreateCertificateV2Request {
  s.Certificate = &v
  return s
}

func (s *CreateCertificateV2Request) SetPrivateKey(v string) *CreateCertificateV2Request {
  s.PrivateKey = &v
  return s
}

func (s *CreateCertificateV2Request) SetCsrId(v int) *CreateCertificateV2Request {
  s.CsrId = &v
  return s
}

func (s *CreateCertificateV2Request) SetComment(v string) *CreateCertificateV2Request {
  s.Comment = &v
  return s
}

type CreateCertificateV2RequestHeader struct {
}

func (s CreateCertificateV2RequestHeader) String() string {
  return tea.Prettify(s)
}

func (s CreateCertificateV2RequestHeader) GoString() string {
  return s.String()
}

type CreateCertificateV2Paths struct {
}

func (s CreateCertificateV2Paths) String() string {
  return tea.Prettify(s)
}

func (s CreateCertificateV2Paths) GoString() string {
  return s.String()
}

type CreateCertificateV2Parameters struct {
}

func (s CreateCertificateV2Parameters) String() string {
  return tea.Prettify(s)
}

func (s CreateCertificateV2Parameters) GoString() string {
  return s.String()
}

type CreateCertificateV2Response struct {
  // {"en":"Status code","zh_CN":"状态码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response message","zh_CN":"响应信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s CreateCertificateV2Response) String() string {
  return tea.Prettify(s)
}

func (s CreateCertificateV2Response) GoString() string {
  return s.String()
}

func (s *CreateCertificateV2Response) SetCode(v string) *CreateCertificateV2Response {
  s.Code = &v
  return s
}

func (s *CreateCertificateV2Response) SetMessage(v string) *CreateCertificateV2Response {
  s.Message = &v
  return s
}

type CreateCertificateV2ResponseHeader struct {
  // {"en":"The URL used to access the certificate file where certificate-id is the unique token that the system generates for the certificate and whose value is a string","zh_CN":"用于访问该证书文件的URL，其中certificate-id为系统为该证书生成的唯一标示，其值为字符串"}
  Location *string `json:"location,omitempty" xml:"location,omitempty" require:"true"`
}

func (s CreateCertificateV2ResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s CreateCertificateV2ResponseHeader) GoString() string {
  return s.String()
}

func (s *CreateCertificateV2ResponseHeader) SetLocation(v string) *CreateCertificateV2ResponseHeader {
  s.Location = &v
  return s
}




