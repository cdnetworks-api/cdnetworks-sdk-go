package certificateapplication

// This file is auto-generated, don't edit it. Thanks.
import (
  "github.com/alibabacloud-go/tea/tea"
)

type GetDomainControlValidationContentRequest struct {
  // {"en":"Order ID", "zh_CN":"订单ID"}
  OrderId *string `json:"orderId,omitempty" xml:"orderId,omitempty"`
  // {"en":"Purchase Record ID", "zh_CN":"采购记录ID"}
  PurchaseRecordId *string `json:"purchaseRecordId,omitempty" xml:"purchaseRecordId,omitempty"`
}

func (s GetDomainControlValidationContentRequest) String() string {
  return tea.Prettify(s)
}

func (s GetDomainControlValidationContentRequest) GoString() string {
  return s.String()
}

func (s *GetDomainControlValidationContentRequest) SetOrderId(v string) *GetDomainControlValidationContentRequest {
  s.OrderId = &v
  return s
}

func (s *GetDomainControlValidationContentRequest) SetPurchaseRecordId(v string) *GetDomainControlValidationContentRequest {
  s.PurchaseRecordId = &v
  return s
}

type GetDomainControlValidationContentResponse struct {
  // {"en":"Result Code, success is 0 ", "zh_CN":"响应码，成功为0"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Result Message", "zh_CN":"响应信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Result Data", "zh_CN":"响应数据"}
  Data *GetDomainControlValidationContentResp `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s GetDomainControlValidationContentResponse) String() string {
  return tea.Prettify(s)
}

func (s GetDomainControlValidationContentResponse) GoString() string {
  return s.String()
}

func (s *GetDomainControlValidationContentResponse) SetCode(v string) *GetDomainControlValidationContentResponse {
  s.Code = &v
  return s
}

func (s *GetDomainControlValidationContentResponse) SetMessage(v string) *GetDomainControlValidationContentResponse {
  s.Message = &v
  return s
}

func (s *GetDomainControlValidationContentResponse) SetData(v *GetDomainControlValidationContentResp) *GetDomainControlValidationContentResponse {
  s.Data = v
  return s
}

type GetDomainControlValidationContentResp struct {
  // {"en":"Order ID", "zh_CN":"订单ID"}
  OrderId *string `json:"orderId,omitempty" xml:"orderId,omitempty" require:"true"`
  // {"en":"Validate Method", "zh_CN":"验证方式"}
  ValidateMethod *string `json:"validateMethod,omitempty" xml:"validateMethod,omitempty" require:"true"`
  // {"en":"Auto Validate", "zh_CN":"是否自动验证"}
  AutoValidate *string `json:"autoValidate,omitempty" xml:"autoValidate,omitempty" require:"true"`
  // {"en":"Domain Validate Information", "zh_CN":"域名验证内容列表"}
  DomainValidateInfos []*GetDomainControlValidationContentDomainValidateInfo `json:"domainValidateInfos,omitempty" xml:"domainValidateInfos,omitempty" require:"true" type:"Repeated"`
  // {"en":"Purchase Record ID", "zh_CN":"采购记录ID"}
  PurchaseRecordId *string `json:"purchaseRecordId,omitempty" xml:"purchaseRecordId,omitempty" require:"true"`
}

func (s GetDomainControlValidationContentResp) String() string {
  return tea.Prettify(s)
}

func (s GetDomainControlValidationContentResp) GoString() string {
  return s.String()
}

func (s *GetDomainControlValidationContentResp) SetOrderId(v string) *GetDomainControlValidationContentResp {
  s.OrderId = &v
  return s
}

func (s *GetDomainControlValidationContentResp) SetValidateMethod(v string) *GetDomainControlValidationContentResp {
  s.ValidateMethod = &v
  return s
}

func (s *GetDomainControlValidationContentResp) SetAutoValidate(v string) *GetDomainControlValidationContentResp {
  s.AutoValidate = &v
  return s
}

func (s *GetDomainControlValidationContentResp) SetDomainValidateInfos(v []*GetDomainControlValidationContentDomainValidateInfo) *GetDomainControlValidationContentResp {
  s.DomainValidateInfos = v
  return s
}

func (s *GetDomainControlValidationContentResp) SetPurchaseRecordId(v string) *GetDomainControlValidationContentResp {
  s.PurchaseRecordId = &v
  return s
}

type GetDomainControlValidationContentDomainValidateInfo struct {
  // {"en":"Domain ", "zh_CN":"域名 "}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {"en":"Validate File Name  ", "zh_CN":"验证文件名 "}
  ValidateFileName *string `json:"validateFileName,omitempty" xml:"validateFileName,omitempty" require:"true"`
  // {"en":"Validate value ", "zh_CN":"验证内容 "}
  ValidateContent *string `json:"validateContent,omitempty" xml:"validateContent,omitempty" require:"true"`
  // {"en":"Expire Time  ", "zh_CN":"过期时间 "}
  ExpireTime *string `json:"expireTime,omitempty" xml:"expireTime,omitempty" require:"true"`
  // {"en":"Validate Method", "zh_CN":"CA厂家实际的验证方式"}
  ValidateMethod *string `json:"validateMethod,omitempty" xml:"validateMethod,omitempty" require:"true"`
  // {"en":"Validate Status", "zh_CN":"验证状态,{INIT: 初始化, VALIDATE_WAIT: 待验证, VALIDATE_PROCESSING: 验证中, VALIDATE_SUCCESS: 验证成功, VALIDATE_FAILURE: 验证失败}"}
  ValidateStatus *string `json:"validateStatus,omitempty" xml:"validateStatus,omitempty" require:"true"`
  // {"en":"recordType", "zh_CN":"记录类型"}
  RecordType *string `json:"recordType,omitempty" xml:"recordType,omitempty" require:"true"`
}

func (s GetDomainControlValidationContentDomainValidateInfo) String() string {
  return tea.Prettify(s)
}

func (s GetDomainControlValidationContentDomainValidateInfo) GoString() string {
  return s.String()
}

func (s *GetDomainControlValidationContentDomainValidateInfo) SetDomain(v string) *GetDomainControlValidationContentDomainValidateInfo {
  s.Domain = &v
  return s
}

func (s *GetDomainControlValidationContentDomainValidateInfo) SetValidateFileName(v string) *GetDomainControlValidationContentDomainValidateInfo {
  s.ValidateFileName = &v
  return s
}

func (s *GetDomainControlValidationContentDomainValidateInfo) SetValidateContent(v string) *GetDomainControlValidationContentDomainValidateInfo {
  s.ValidateContent = &v
  return s
}

func (s *GetDomainControlValidationContentDomainValidateInfo) SetExpireTime(v string) *GetDomainControlValidationContentDomainValidateInfo {
  s.ExpireTime = &v
  return s
}

func (s *GetDomainControlValidationContentDomainValidateInfo) SetValidateMethod(v string) *GetDomainControlValidationContentDomainValidateInfo {
  s.ValidateMethod = &v
  return s
}

func (s *GetDomainControlValidationContentDomainValidateInfo) SetValidateStatus(v string) *GetDomainControlValidationContentDomainValidateInfo {
  s.ValidateStatus = &v
  return s
}

func (s *GetDomainControlValidationContentDomainValidateInfo) SetRecordType(v string) *GetDomainControlValidationContentDomainValidateInfo {
  s.RecordType = &v
  return s
}

type GetDomainControlValidationContentPaths struct {
}

func (s GetDomainControlValidationContentPaths) String() string {
  return tea.Prettify(s)
}

func (s GetDomainControlValidationContentPaths) GoString() string {
  return s.String()
}

type GetDomainControlValidationContentParameters struct {
}

func (s GetDomainControlValidationContentParameters) String() string {
  return tea.Prettify(s)
}

func (s GetDomainControlValidationContentParameters) GoString() string {
  return s.String()
}

type GetDomainControlValidationContentRequestHeader struct {
}

func (s GetDomainControlValidationContentRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetDomainControlValidationContentRequestHeader) GoString() string {
  return s.String()
}

type GetDomainControlValidationContentResponseHeader struct {
}

func (s GetDomainControlValidationContentResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetDomainControlValidationContentResponseHeader) GoString() string {
  return s.String()
}




type QueryCertificateSalesOrderDetailForWplusRequest struct {
  // {"en":"orderId", "zh_CN":"订单id"}
  OrderId *string `json:"orderId,omitempty" xml:"orderId,omitempty"`
  // {"en":"purchaseRecordId", "zh_CN":"采购记录ID"}
  PurchaseRecordId *string `json:"purchaseRecordId,omitempty" xml:"purchaseRecordId,omitempty"`
}

func (s QueryCertificateSalesOrderDetailForWplusRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryCertificateSalesOrderDetailForWplusRequest) GoString() string {
  return s.String()
}

func (s *QueryCertificateSalesOrderDetailForWplusRequest) SetOrderId(v string) *QueryCertificateSalesOrderDetailForWplusRequest {
  s.OrderId = &v
  return s
}

func (s *QueryCertificateSalesOrderDetailForWplusRequest) SetPurchaseRecordId(v string) *QueryCertificateSalesOrderDetailForWplusRequest {
  s.PurchaseRecordId = &v
  return s
}

type QueryCertificateSalesOrderDetailForWplusResponse struct {
  // {"en":"code", "zh_CN":"错误码，成功为0"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"error message", "zh_CN":"错误信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"data", "zh_CN":"订单详情"}
  Data *QueryCertificateSalesOrderDetailForWplusOrderDetail `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s QueryCertificateSalesOrderDetailForWplusResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryCertificateSalesOrderDetailForWplusResponse) GoString() string {
  return s.String()
}

func (s *QueryCertificateSalesOrderDetailForWplusResponse) SetCode(v string) *QueryCertificateSalesOrderDetailForWplusResponse {
  s.Code = &v
  return s
}

func (s *QueryCertificateSalesOrderDetailForWplusResponse) SetMessage(v string) *QueryCertificateSalesOrderDetailForWplusResponse {
  s.Message = &v
  return s
}

func (s *QueryCertificateSalesOrderDetailForWplusResponse) SetData(v *QueryCertificateSalesOrderDetailForWplusOrderDetail) *QueryCertificateSalesOrderDetailForWplusResponse {
  s.Data = v
  return s
}

type QueryCertificateSalesOrderDetailForWplusOrderDetail struct {
  // {"en":"orderId", "zh_CN":"订单id"}
  OrderId *string `json:"orderId,omitempty" xml:"orderId,omitempty" require:"true"`
  // {"en":"orderStatus
  // ACCEPT_SUCCESS: Received successfully
  // APPLYING: Under application
  // APPLy_FAILURE: Application preparation failed
  // VALIDATETWAIT: To be verified
  // VALIDATE_SUCCESS: Verification successful
  // VALIDATE_SAILURE: Validation failed
  // ISSUEAwaIT: To be issued
  // ISSUE_SUCCESS: Issued successfully
  // ISSUE_SFAILURE: Issuance failed
  // CANCELED: Cancel application
  // REVOKED: Revocation
  // DEPLO_SUCCESS: Deployment successful
  // DEPLOy_Failures: Deployment failed", "zh_CN":"订单状态
  // ACCEPT_SUCCESS: 接收成功
  // APPLYING: 申请中
  // APPLY_FAILURE：申请准备失败
  // VALIDATE_WAIT: 待验证
  // VALIDATE_SUCCESS: 验证成功
  // VALIDATE_FAILURE: 验证失败
  // ISSUE_WAIT: 待签发
  // ISSUE_SUCCESS: 签发成功
  // ISSUE_FAILURE: 签发失败
  // CANCELED: 取消申请
  // REVOKED: 吊销
  // DEPLOY_SUCCESS：部署成功
  // DEPLOY_FAILURE：部署失败"}
  OrderStatus *string `json:"orderStatus,omitempty" xml:"orderStatus,omitempty" require:"true"`
  // {"en":"certificateId", "zh_CN":"证书id"}
  CertificateId *int64 `json:"certificateId,omitempty" xml:"certificateId,omitempty" require:"true"`
  // {"en":"certificateName", "zh_CN":"证书名"}
  CertificateName *string `json:"certificateName,omitempty" xml:"certificateName,omitempty" require:"true"`
  // {"en":"certificateBrand", "zh_CN":"证书品牌"}
  CertificateBrand *string `json:"certificateBrand,omitempty" xml:"certificateBrand,omitempty" require:"true"`
  // {"en":"certificateType", "zh_CN":"证书类型"}
  CertificateType *string `json:"certificateType,omitempty" xml:"certificateType,omitempty" require:"true"`
  // {"en":"description", "zh_CN":"描述"}
  Description *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
  // {"en":"algorithm", "zh_CN":"算法"}
  Algorithm *string `json:"algorithm,omitempty" xml:"algorithm,omitempty" require:"true"`
  // {"en":"autoValidate", "zh_CN":"是否自动验证"}
  AutoValidate *string `json:"autoValidate,omitempty" xml:"autoValidate,omitempty" require:"true"`
  // {"en":"validateMethod", "zh_CN":"验证方式"}
  ValidateMethod *string `json:"validateMethod,omitempty" xml:"validateMethod,omitempty" require:"true"`
  // {"en":"autoRenew", "zh_CN":"是否自动续订"}
  AutoRenew *string `json:"autoRenew,omitempty" xml:"autoRenew,omitempty" require:"true"`
  // {"en":"validityDays", "zh_CN":"有效天数"}
  ValidityDays *int `json:"validityDays,omitempty" xml:"validityDays,omitempty" require:"true"`
  // {"en":"country", "zh_CN":"国家代码"}
  Country *string `json:"country,omitempty" xml:"country,omitempty" require:"true"`
  // {"en":"state", "zh_CN":"省"}
  State *string `json:"state,omitempty" xml:"state,omitempty" require:"true"`
  // {"en":"city", "zh_CN":"市"}
  City *string `json:"city,omitempty" xml:"city,omitempty" require:"true"`
  // {"en":"street", "zh_CN":"街道"}
  Street *string `json:"street,omitempty" xml:"street,omitempty" require:"true"`
  // {"en":"company", "zh_CN":"公司"}
  Company *string `json:"company,omitempty" xml:"company,omitempty" require:"true"`
  // {"en":"department", "zh_CN":"部门"}
  Department *string `json:"department,omitempty" xml:"department,omitempty" require:"true"`
  // {"en":"commonName", "zh_CN":"通用名"}
  CommonName *string `json:"commonName,omitempty" xml:"commonName,omitempty" require:"true"`
  // {"en":"email", "zh_CN":"邮箱"}
  Email *string `json:"email,omitempty" xml:"email,omitempty" require:"true"`
  // {"en":"subjectAlternativeNames", "zh_CN":"主体备用名称"}
  SubjectAlternativeNames *string `json:"subjectAlternativeNames,omitempty" xml:"subjectAlternativeNames,omitempty" require:"true"`
  // {"en":"createTime", "zh_CN":"创建时间"}
  CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty" require:"true"`
  // {"en":"errorMessage", "zh_CN":"错误信息"}
  ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty" require:"true"`
  // {"en":"certificateSpec", "zh_CN":"证书规格"}
  CertificateSpec *string `json:"certificateSpec,omitempty" xml:"certificateSpec,omitempty" require:"true"`
  // {"en":"domainType", "zh_CN":"域名类型"}
  DomainType *string `json:"domainType,omitempty" xml:"domainType,omitempty" require:"true"`
  // {"en":"autoDeploy", "zh_CN":"是否自动部署"}
  AutoDeploy *string `json:"autoDeploy,omitempty" xml:"autoDeploy,omitempty" require:"true"`
}

func (s QueryCertificateSalesOrderDetailForWplusOrderDetail) String() string {
  return tea.Prettify(s)
}

func (s QueryCertificateSalesOrderDetailForWplusOrderDetail) GoString() string {
  return s.String()
}

func (s *QueryCertificateSalesOrderDetailForWplusOrderDetail) SetOrderId(v string) *QueryCertificateSalesOrderDetailForWplusOrderDetail {
  s.OrderId = &v
  return s
}

func (s *QueryCertificateSalesOrderDetailForWplusOrderDetail) SetOrderStatus(v string) *QueryCertificateSalesOrderDetailForWplusOrderDetail {
  s.OrderStatus = &v
  return s
}

func (s *QueryCertificateSalesOrderDetailForWplusOrderDetail) SetCertificateId(v int64) *QueryCertificateSalesOrderDetailForWplusOrderDetail {
  s.CertificateId = &v
  return s
}

func (s *QueryCertificateSalesOrderDetailForWplusOrderDetail) SetCertificateName(v string) *QueryCertificateSalesOrderDetailForWplusOrderDetail {
  s.CertificateName = &v
  return s
}

func (s *QueryCertificateSalesOrderDetailForWplusOrderDetail) SetCertificateBrand(v string) *QueryCertificateSalesOrderDetailForWplusOrderDetail {
  s.CertificateBrand = &v
  return s
}

func (s *QueryCertificateSalesOrderDetailForWplusOrderDetail) SetCertificateType(v string) *QueryCertificateSalesOrderDetailForWplusOrderDetail {
  s.CertificateType = &v
  return s
}

func (s *QueryCertificateSalesOrderDetailForWplusOrderDetail) SetDescription(v string) *QueryCertificateSalesOrderDetailForWplusOrderDetail {
  s.Description = &v
  return s
}

func (s *QueryCertificateSalesOrderDetailForWplusOrderDetail) SetAlgorithm(v string) *QueryCertificateSalesOrderDetailForWplusOrderDetail {
  s.Algorithm = &v
  return s
}

func (s *QueryCertificateSalesOrderDetailForWplusOrderDetail) SetAutoValidate(v string) *QueryCertificateSalesOrderDetailForWplusOrderDetail {
  s.AutoValidate = &v
  return s
}

func (s *QueryCertificateSalesOrderDetailForWplusOrderDetail) SetValidateMethod(v string) *QueryCertificateSalesOrderDetailForWplusOrderDetail {
  s.ValidateMethod = &v
  return s
}

func (s *QueryCertificateSalesOrderDetailForWplusOrderDetail) SetAutoRenew(v string) *QueryCertificateSalesOrderDetailForWplusOrderDetail {
  s.AutoRenew = &v
  return s
}

func (s *QueryCertificateSalesOrderDetailForWplusOrderDetail) SetValidityDays(v int) *QueryCertificateSalesOrderDetailForWplusOrderDetail {
  s.ValidityDays = &v
  return s
}

func (s *QueryCertificateSalesOrderDetailForWplusOrderDetail) SetCountry(v string) *QueryCertificateSalesOrderDetailForWplusOrderDetail {
  s.Country = &v
  return s
}

func (s *QueryCertificateSalesOrderDetailForWplusOrderDetail) SetState(v string) *QueryCertificateSalesOrderDetailForWplusOrderDetail {
  s.State = &v
  return s
}

func (s *QueryCertificateSalesOrderDetailForWplusOrderDetail) SetCity(v string) *QueryCertificateSalesOrderDetailForWplusOrderDetail {
  s.City = &v
  return s
}

func (s *QueryCertificateSalesOrderDetailForWplusOrderDetail) SetStreet(v string) *QueryCertificateSalesOrderDetailForWplusOrderDetail {
  s.Street = &v
  return s
}

func (s *QueryCertificateSalesOrderDetailForWplusOrderDetail) SetCompany(v string) *QueryCertificateSalesOrderDetailForWplusOrderDetail {
  s.Company = &v
  return s
}

func (s *QueryCertificateSalesOrderDetailForWplusOrderDetail) SetDepartment(v string) *QueryCertificateSalesOrderDetailForWplusOrderDetail {
  s.Department = &v
  return s
}

func (s *QueryCertificateSalesOrderDetailForWplusOrderDetail) SetCommonName(v string) *QueryCertificateSalesOrderDetailForWplusOrderDetail {
  s.CommonName = &v
  return s
}

func (s *QueryCertificateSalesOrderDetailForWplusOrderDetail) SetEmail(v string) *QueryCertificateSalesOrderDetailForWplusOrderDetail {
  s.Email = &v
  return s
}

func (s *QueryCertificateSalesOrderDetailForWplusOrderDetail) SetSubjectAlternativeNames(v string) *QueryCertificateSalesOrderDetailForWplusOrderDetail {
  s.SubjectAlternativeNames = &v
  return s
}

func (s *QueryCertificateSalesOrderDetailForWplusOrderDetail) SetCreateTime(v string) *QueryCertificateSalesOrderDetailForWplusOrderDetail {
  s.CreateTime = &v
  return s
}

func (s *QueryCertificateSalesOrderDetailForWplusOrderDetail) SetErrorMessage(v string) *QueryCertificateSalesOrderDetailForWplusOrderDetail {
  s.ErrorMessage = &v
  return s
}

func (s *QueryCertificateSalesOrderDetailForWplusOrderDetail) SetCertificateSpec(v string) *QueryCertificateSalesOrderDetailForWplusOrderDetail {
  s.CertificateSpec = &v
  return s
}

func (s *QueryCertificateSalesOrderDetailForWplusOrderDetail) SetDomainType(v string) *QueryCertificateSalesOrderDetailForWplusOrderDetail {
  s.DomainType = &v
  return s
}

func (s *QueryCertificateSalesOrderDetailForWplusOrderDetail) SetAutoDeploy(v string) *QueryCertificateSalesOrderDetailForWplusOrderDetail {
  s.AutoDeploy = &v
  return s
}

type QueryCertificateSalesOrderDetailForWplusPaths struct {
}

func (s QueryCertificateSalesOrderDetailForWplusPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryCertificateSalesOrderDetailForWplusPaths) GoString() string {
  return s.String()
}

type QueryCertificateSalesOrderDetailForWplusParameters struct {
}

func (s QueryCertificateSalesOrderDetailForWplusParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryCertificateSalesOrderDetailForWplusParameters) GoString() string {
  return s.String()
}

type QueryCertificateSalesOrderDetailForWplusRequestHeader struct {
}

func (s QueryCertificateSalesOrderDetailForWplusRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryCertificateSalesOrderDetailForWplusRequestHeader) GoString() string {
  return s.String()
}

type QueryCertificateSalesOrderDetailForWplusResponseHeader struct {
}

func (s QueryCertificateSalesOrderDetailForWplusResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryCertificateSalesOrderDetailForWplusResponseHeader) GoString() string {
  return s.String()
}




type CreateCertificateApplyingOrderRequest struct {
  // {"en":"Name of the certificate. The certificate name cannot be the same as your existing certificate.
  // Range: <= 128 characters.", "zh_CN":"证书名称，最长不超过128个字符。证书名称不允许与已有的证书重名."}
  CertificateName *string `json:"certificateName,omitempty" xml:"certificateName,omitempty"`
  // {"en":"A description of the new certificate.
  // Range: <= 256 characters.", "zh_CN":"证书描述，最长不超过256个字符。"}
  Description *string `json:"description,omitempty" xml:"description,omitempty"`
  // {"en":"Automatically Renew your certificate. Selecting 'true' will reduce customization of other fields: 'certificateBrand' must be 'LE','validateMethod' must be 'HTTP', the certificate domain must configure in our acceleration platform; the domain  must CNAME to our acceleration platform.
  // Allowed values: true, false(default) .", "zh_CN":"是否自动续签，取值范围：true、false(默认) 。选择是否在证书到期前为您自动续订证书，开启自动续签，证书品牌必须选择'LE'；验证方式必须选择'HTTP'；申请证书包含的域名必须已在我司加速平台配置，且域名必须已CNAME到我司加速平台。"}
  AutoRenew *string `json:"autoRenew,omitempty" xml:"autoRenew,omitempty"`
  // {"en":"Certificate Brand. Selecting 'LE' means we apply your certificate through Let's Encrypt (https://letsencrypt.org/docs/challenge-types/) .
  // Allowed value: LE TrustAsia . Note: If the LE brand application fails, it will automatically switch to the ZeroSSL CA manufacturer for certificate application", "zh_CN":"证书品牌，取值范围：LE、TrustAsia 。 'LE '即Let's Encrypt。备注：如果LE品牌申请失败，会自动切换成ZeroSSL CA厂家进行证书申请"}
  CertificateBrand *string `json:"certificateBrand,omitempty" xml:"certificateBrand,omitempty" require:"true"`
  // {"en":"Certificate Specification,
  // Certificate brand GlobalSign, certificate specification GlobalSignOVMSSSL, certificate type OV, certificate validity days [360], billing;
  // Certificate brand GlobalSign, certificate specification GlobalSignOVSANsMSSL, certificate type OV, certificate validity days [360], billing;
  // Certificate brand TrustAsia, certificate specification TrustAsia TLSC1, certificate type DV, certificate validity days [360], billing;
  // Certificate brand TrustAsia, certificate specification TrustAsiaTLSSANsC1, certificate type DV, certificate validity days [360], billing;
  // Certificate brand TrustAsia, certificate specification TrustAsiaC1DVFree, certificate type DV, certificate validity days [90], free of charge;
  // Certificate brand TrustAsia, certificate specification TrustAsiaC1DVFreeSANS, certificate type DV, certificate validity days [90], free of charge;
  // Certificate brand LE, certificate specification LetsEncryptDVFree, certificate type DV, certificate validity days [90], free of charge;
  // Certificate Brand LE, Certificate Specification LetsEncryptDVFreeSANS, Certificate Type DV, Certificate Validity Days [90], free of charge", "zh_CN":"证书规格,
  // 证书品牌GlobalSign、证书规格GlobalSignOVMSSL、证书类型OV、证书有效天数[360]、计费;
  // 证书品牌GlobalSign、证书规格GlobalSignOVSANsMSSL、证书类型OV、证书有效天数[360]、计费;
  // 证书品牌TrustAsia、证书规格TrustAsiaTLSC1、证书类型DV、证书有效天数[360]、计费;
  // 证书品牌TrustAsia、证书规格TrustAsiaTLSSANsC1、证书类型DV、证书有效天数[360]、计费;
  // 证书品牌TrustAsia、证书规格TrustAsiaC1DVFree、证书类型DV、证书有效天数[90]、免费;
  // 证书品牌TrustAsia、证书规格TrustAsiaC1DVFreeSANS、证书类型DV、证书有效天数[90]、免费;
  // 证书品牌LE、证书规格LetsEncryptDVFree、证书类型DV、证书有效天数[90]、免费;
  // 证书品牌LE、证书规格LetsEncryptDVFreeSANS、证书类型DV、证书有效天数[90]、免费"}
  CertificateSpec *string `json:"certificateSpec,omitempty" xml:"certificateSpec,omitempty" require:"true"`
  // {"en":"Certificate Type. When 'certificateBrand' specifying 'LE',  'certificateType' only supports 'DV'.
  // TrustAsia supports DV and OV.", "zh_CN":"安全等级，取值范围: DV、OV 。 Let's Encrypt仅支持DV， TrustAsia支持DV和OV。"}
  CertificateType *string `json:"certificateType,omitempty" xml:"certificateType,omitempty" require:"true"`
  // {"en":"Certificate Algorithm. When 'certificateBrand' specifying 'LE',  'certificateType' only supports 'RSA2048' and 'ECDSA256'.
  // Allowed values: RSA2048, ECDSA256 .", "zh_CN":"证书算法，取值范围：RSA2048、ECDSA256。 Let's Encrypt仅支持RSA2048与ECDSA256  。"}
  Algorithm *string `json:"algorithm,omitempty" xml:"algorithm,omitempty" require:"true"`
  // {"en":"Automatically validate that you control the domains in the certificate you applied. Selecting 'true' will reduce customization of other fields: 'validateMethod' must be 'HTTP' , the certificate domain must configure in our acceleration platform; the domain must CNAME to our acceleration platform.
  // Allowed values: true, false .", "zh_CN":"是否自动验证，取值范围：true、false。选择是否帮您自动完成证书域名控制权验证。开启自动验证，证书品牌必须选择'LE'；验证方式必须选择'HTTP'；申请证书包含的域名必须已在我司加速平台配置，且域名必须已CNAME到我司加速平台。"}
  AutoValidate *string `json:"autoValidate,omitempty" xml:"autoValidate,omitempty" require:"true"`
  // {"en":"Validate Method selected for the order to validate that you control the domains in the certificate. Validate Method must be 'DNS' when 'subjectAlternativeNames'  is a wildcard hostname beginning with '*' such as '*.example.com' .
  // Allowed values: HTTP, DNS .  When applying for a certificate with an IP domain name, only HTTP authentication is supported.", "zh_CN":"验证方式，取值范围：HTTP、DNS 。选择证书域名控制权的验证方式。当申请证书域名'subjectAlternativeNames'为泛域名时，仅支持DNS验证。当申请证书域名为IP时，仅支持HTTP验证。"}
  ValidateMethod *string `json:"validateMethod,omitempty" xml:"validateMethod,omitempty" require:"true"`
  // {"en":"Whether to deploy automatically. Value range: true, false, true: automatic deployment, false: manual deployment.", "zh_CN":"是否自动部署，取值范围：true、false，true:自动部署,   false:手动部署"}
  AutoDeploy *string `json:"autoDeploy,omitempty" xml:"autoDeploy,omitempty" require:"true"`
  // {"en":"Validity Days selected for the certificate. Please refer to the description in the certificateSpec for details.", "zh_CN":"证书有效天数，详见certificateSpec描述"}
  ValidityDays *int `json:"validityDays,omitempty" xml:"validityDays,omitempty" require:"true"`
  // {"en":"Certificate Signing Request Informtion.", "zh_CN":"证书签名请求信息 。"}
  CreateCertificateApplyingOrderIdentificationInfo *CreateCertificateApplyingOrderIdentificationInfo `json:"identificationInfo,omitempty" xml:"identificationInfo,omitempty" require:"true"`
  // {"en":"Backup Certificate Brand.
  // The certificate brand LE for a single domain name is optional, and the backup CA certificate brand is TrustAsia ZeroSSL
  // Certificate brand: LE for multiple domain names, optional backup CA certificate brand: TrustAsia
  // Certificate brand ZeroSSL single domain name optional backup CA certificate brand TrustAsia LE
  // Certificate brand TrustAsia single domain name optional backup CA certificate brand is LE ZeroSSL
  // Certificate Brand TrustAsia Multi Domain Optional Backup CA Certificate Brand LE", "zh_CN":"备用CA证书品牌，
  // 证书品牌LE单域名可选备用CA证书品牌为TrustAsia、ZeroSSL
  // 证书品牌LE多域名可选备用CA证书品牌为TrustAsia
  // 证书品牌ZeroSSL单域名可选备用CA证书品牌为TrustAsia、LE
  // 证书品牌TrustAsia单域名可选备用CA证书品牌为LE、ZeroSSL
  // 证书品牌TrustAsia多域名可选备用CA证书品牌为LE"}
  BackupCertificateBrand *string `json:"backupCertificateBrand,omitempty" xml:"backupCertificateBrand,omitempty"`
  // {"en":"orgValidateMethod, default:  Enable pre audit, self_validate:  Self verify organizational information, none: Unorganized information, default value default.", "zh_CN":"组织验证方式, default: 开启预审核、self_validate: 自行验证组织信息、none: 无组织信息，默认值default。"}
  OrgValidateMethod *string `json:"orgValidateMethod,omitempty" xml:"orgValidateMethod,omitempty"`
  // {"en":"domainType, Value range: single multi", "zh_CN":"域名类型，取值范围：single、multi"}
  DomainType *string `json:"domainType,omitempty" xml:"domainType,omitempty"`
  // {"en":"batchApply, Value range: true false, Default false", "zh_CN":"是否批量申请，取值范围：true、false，默认false"}
  BatchApply *string `json:"batchApply,omitempty" xml:"batchApply,omitempty"`
  // {"en":"admin", "zh_CN":"订单管理联系人"}
  Admin *CreateCertificateApplyingOrderLinkman `json:"admin,omitempty" xml:"admin,omitempty"`
  // {"en":"tech", "zh_CN":"订单技术联系人"}
  Tech *CreateCertificateApplyingOrderLinkman `json:"tech,omitempty" xml:"tech,omitempty"`
  // {"en":"DNS Provider Infomation", "zh_CN":"DNS托管商信息"}
  DnsProviderInfos []*CreateCertificateApplyingOrderDnsProviderInfo `json:"dnsProviderInfos,omitempty" xml:"dnsProviderInfos,omitempty" type:"Repeated"`
}

func (s CreateCertificateApplyingOrderRequest) String() string {
  return tea.Prettify(s)
}

func (s CreateCertificateApplyingOrderRequest) GoString() string {
  return s.String()
}

func (s *CreateCertificateApplyingOrderRequest) SetCertificateName(v string) *CreateCertificateApplyingOrderRequest {
  s.CertificateName = &v
  return s
}

func (s *CreateCertificateApplyingOrderRequest) SetDescription(v string) *CreateCertificateApplyingOrderRequest {
  s.Description = &v
  return s
}

func (s *CreateCertificateApplyingOrderRequest) SetAutoRenew(v string) *CreateCertificateApplyingOrderRequest {
  s.AutoRenew = &v
  return s
}

func (s *CreateCertificateApplyingOrderRequest) SetCertificateBrand(v string) *CreateCertificateApplyingOrderRequest {
  s.CertificateBrand = &v
  return s
}

func (s *CreateCertificateApplyingOrderRequest) SetCertificateSpec(v string) *CreateCertificateApplyingOrderRequest {
  s.CertificateSpec = &v
  return s
}

func (s *CreateCertificateApplyingOrderRequest) SetCertificateType(v string) *CreateCertificateApplyingOrderRequest {
  s.CertificateType = &v
  return s
}

func (s *CreateCertificateApplyingOrderRequest) SetAlgorithm(v string) *CreateCertificateApplyingOrderRequest {
  s.Algorithm = &v
  return s
}

func (s *CreateCertificateApplyingOrderRequest) SetAutoValidate(v string) *CreateCertificateApplyingOrderRequest {
  s.AutoValidate = &v
  return s
}

func (s *CreateCertificateApplyingOrderRequest) SetValidateMethod(v string) *CreateCertificateApplyingOrderRequest {
  s.ValidateMethod = &v
  return s
}

func (s *CreateCertificateApplyingOrderRequest) SetAutoDeploy(v string) *CreateCertificateApplyingOrderRequest {
  s.AutoDeploy = &v
  return s
}

func (s *CreateCertificateApplyingOrderRequest) SetValidityDays(v int) *CreateCertificateApplyingOrderRequest {
  s.ValidityDays = &v
  return s
}

func (s *CreateCertificateApplyingOrderRequest) SetIdentificationInfo(v *CreateCertificateApplyingOrderIdentificationInfo) *CreateCertificateApplyingOrderRequest {
  s.CreateCertificateApplyingOrderIdentificationInfo = v
  return s
}

func (s *CreateCertificateApplyingOrderRequest) SetBackupCertificateBrand(v string) *CreateCertificateApplyingOrderRequest {
  s.BackupCertificateBrand = &v
  return s
}

func (s *CreateCertificateApplyingOrderRequest) SetOrgValidateMethod(v string) *CreateCertificateApplyingOrderRequest {
  s.OrgValidateMethod = &v
  return s
}

func (s *CreateCertificateApplyingOrderRequest) SetDomainType(v string) *CreateCertificateApplyingOrderRequest {
  s.DomainType = &v
  return s
}

func (s *CreateCertificateApplyingOrderRequest) SetBatchApply(v string) *CreateCertificateApplyingOrderRequest {
  s.BatchApply = &v
  return s
}

func (s *CreateCertificateApplyingOrderRequest) SetAdmin(v *CreateCertificateApplyingOrderLinkman) *CreateCertificateApplyingOrderRequest {
  s.Admin = v
  return s
}

func (s *CreateCertificateApplyingOrderRequest) SetTech(v *CreateCertificateApplyingOrderLinkman) *CreateCertificateApplyingOrderRequest {
  s.Tech = v
  return s
}

func (s *CreateCertificateApplyingOrderRequest) SetDnsProviderInfos(v []*CreateCertificateApplyingOrderDnsProviderInfo) *CreateCertificateApplyingOrderRequest {
  s.DnsProviderInfos = v
  return s
}

type CreateCertificateApplyingOrderIdentificationInfo struct {
  // {"en":"Country. An ISO-3166 country code.
  // Range: = 2 characters", "zh_CN":"国家，2个字符的ISO-3166国家代码"}
  Country *string `json:"country,omitempty" xml:"country,omitempty"`
  // {"en":"A state or province.", "zh_CN":"省"}
  State *string `json:"state,omitempty" xml:"state,omitempty"`
  // {"en":"A city", "zh_CN":"城市"}
  City *string `json:"city,omitempty" xml:"city,omitempty"`
  // {"en":"A company name", "zh_CN":"公司"}
  Company *string `json:"company,omitempty" xml:"company,omitempty"`
  // {"en":"The department associated with the certificate", "zh_CN":"部门"}
  Department *string `json:"department,omitempty" xml:"department,omitempty"`
  // {"en":"A common name of the certificate", "zh_CN":"证书的通用名称"}
  CommonName *string `json:"commonName,omitempty" xml:"commonName,omitempty" require:"true"`
  // {"en":"Email address", "zh_CN":"邮箱"}
  Email *string `json:"email,omitempty" xml:"email,omitempty"`
  // {"en":"The street where the company is located", "zh_CN":"街道"}
  Street *string `json:"street,omitempty" xml:"street,omitempty"`
  // {"en":"Hostnames that this certificate will serve. Each entry must be a valid hostname such as example.com or a wildcard hostname beginning with '*' such as.example.com.", "zh_CN":"此证书将提供服务的域名。 每个条目必须是有效的域名（例如 example.com）或以“*”开头的泛域名（例如 “*.example.com”）"}
  SubjectAlternativeNames []*string `json:"subjectAlternativeNames,omitempty" xml:"subjectAlternativeNames,omitempty" require:"true" type:"Repeated"`
  // {"en":"street1", "zh_CN":"街道"}
  Street1 *string `json:"street1,omitempty" xml:"street1,omitempty"`
  // {"en":"phone", "zh_CN":"联系电话"}
  Phone *string `json:"phone,omitempty" xml:"phone,omitempty"`
  // {"en":"postalCode", "zh_CN":"邮政编码"}
  PostalCode *string `json:"postalCode,omitempty" xml:"postalCode,omitempty"`
}

func (s CreateCertificateApplyingOrderIdentificationInfo) String() string {
  return tea.Prettify(s)
}

func (s CreateCertificateApplyingOrderIdentificationInfo) GoString() string {
  return s.String()
}

func (s *CreateCertificateApplyingOrderIdentificationInfo) SetCountry(v string) *CreateCertificateApplyingOrderIdentificationInfo {
  s.Country = &v
  return s
}

func (s *CreateCertificateApplyingOrderIdentificationInfo) SetState(v string) *CreateCertificateApplyingOrderIdentificationInfo {
  s.State = &v
  return s
}

func (s *CreateCertificateApplyingOrderIdentificationInfo) SetCity(v string) *CreateCertificateApplyingOrderIdentificationInfo {
  s.City = &v
  return s
}

func (s *CreateCertificateApplyingOrderIdentificationInfo) SetCompany(v string) *CreateCertificateApplyingOrderIdentificationInfo {
  s.Company = &v
  return s
}

func (s *CreateCertificateApplyingOrderIdentificationInfo) SetDepartment(v string) *CreateCertificateApplyingOrderIdentificationInfo {
  s.Department = &v
  return s
}

func (s *CreateCertificateApplyingOrderIdentificationInfo) SetCommonName(v string) *CreateCertificateApplyingOrderIdentificationInfo {
  s.CommonName = &v
  return s
}

func (s *CreateCertificateApplyingOrderIdentificationInfo) SetEmail(v string) *CreateCertificateApplyingOrderIdentificationInfo {
  s.Email = &v
  return s
}

func (s *CreateCertificateApplyingOrderIdentificationInfo) SetStreet(v string) *CreateCertificateApplyingOrderIdentificationInfo {
  s.Street = &v
  return s
}

func (s *CreateCertificateApplyingOrderIdentificationInfo) SetSubjectAlternativeNames(v []*string) *CreateCertificateApplyingOrderIdentificationInfo {
  s.SubjectAlternativeNames = v
  return s
}

func (s *CreateCertificateApplyingOrderIdentificationInfo) SetStreet1(v string) *CreateCertificateApplyingOrderIdentificationInfo {
  s.Street1 = &v
  return s
}

func (s *CreateCertificateApplyingOrderIdentificationInfo) SetPhone(v string) *CreateCertificateApplyingOrderIdentificationInfo {
  s.Phone = &v
  return s
}

func (s *CreateCertificateApplyingOrderIdentificationInfo) SetPostalCode(v string) *CreateCertificateApplyingOrderIdentificationInfo {
  s.PostalCode = &v
  return s
}

type CreateCertificateApplyingOrderLinkman struct {
  // {"en":"firstName, Self verify organizational information. This field is mandatory", "zh_CN":"名，自行验证组织信息此字段必填"}
  FirstName *string `json:"firstName,omitempty" xml:"firstName,omitempty"`
  // {"en":"lastName, Self verify organizational information. This field is mandatory", "zh_CN":"姓，自行验证组织信息此字段必填"}
  LastName *string `json:"lastName,omitempty" xml:"lastName,omitempty"`
  // {"en":"phone, Self verify organizational information. This field is mandatory", "zh_CN":"电话，自行验证组织信息此字段必填"}
  Phone *string `json:"phone,omitempty" xml:"phone,omitempty"`
  // {"en":"email, Self verify organizational information. This field is mandatory", "zh_CN":"邮箱，自行验证组织信息此字段必填"}
  Email *string `json:"email,omitempty" xml:"email,omitempty"`
  // {"en":"title, Self verify organizational information. This field is mandatory", "zh_CN":"职位，自行验证组织信息此字段必填"}
  Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s CreateCertificateApplyingOrderLinkman) String() string {
  return tea.Prettify(s)
}

func (s CreateCertificateApplyingOrderLinkman) GoString() string {
  return s.String()
}

func (s *CreateCertificateApplyingOrderLinkman) SetFirstName(v string) *CreateCertificateApplyingOrderLinkman {
  s.FirstName = &v
  return s
}

func (s *CreateCertificateApplyingOrderLinkman) SetLastName(v string) *CreateCertificateApplyingOrderLinkman {
  s.LastName = &v
  return s
}

func (s *CreateCertificateApplyingOrderLinkman) SetPhone(v string) *CreateCertificateApplyingOrderLinkman {
  s.Phone = &v
  return s
}

func (s *CreateCertificateApplyingOrderLinkman) SetEmail(v string) *CreateCertificateApplyingOrderLinkman {
  s.Email = &v
  return s
}

func (s *CreateCertificateApplyingOrderLinkman) SetTitle(v string) *CreateCertificateApplyingOrderLinkman {
  s.Title = &v
  return s
}

type CreateCertificateApplyingOrderDnsProviderInfo struct {
  // {"en":"domain", "zh_CN":"域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {"en":"DNS Provider Code, Support CloudDNS", "zh_CN":"DNS托管商编码，支持CloudDNS"}
  DnsProviderCode *string `json:"dnsProviderCode,omitempty" xml:"dnsProviderCode,omitempty" require:"true"`
  // {"en":"DNS Api Access, JSON format: The hosting provider is CloudDNS, JSON KEY is accessKey  and secretKey", "zh_CN":"DNS托管商API凭证，JSON格式，托管商为CloudDNS，JSON KEY为accessKey、secretKey"}
  DnsApiAccess *string `json:"dnsApiAccess,omitempty" xml:"dnsApiAccess,omitempty" require:"true"`
}

func (s CreateCertificateApplyingOrderDnsProviderInfo) String() string {
  return tea.Prettify(s)
}

func (s CreateCertificateApplyingOrderDnsProviderInfo) GoString() string {
  return s.String()
}

func (s *CreateCertificateApplyingOrderDnsProviderInfo) SetDomain(v string) *CreateCertificateApplyingOrderDnsProviderInfo {
  s.Domain = &v
  return s
}

func (s *CreateCertificateApplyingOrderDnsProviderInfo) SetDnsProviderCode(v string) *CreateCertificateApplyingOrderDnsProviderInfo {
  s.DnsProviderCode = &v
  return s
}

func (s *CreateCertificateApplyingOrderDnsProviderInfo) SetDnsApiAccess(v string) *CreateCertificateApplyingOrderDnsProviderInfo {
  s.DnsApiAccess = &v
  return s
}

type CreateCertificateApplyingOrderResponse struct {
  // {"en":"Result Code, success is 0", "zh_CN":"响应码，成功为0"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Result Message", "zh_CN":"响应信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Result Data", "zh_CN":"响应数据"}
  Data map[string]interface{} `json:"data,omitempty" xml:"data,omitempty" require:"true"`
  // {"en":"Order ID", "zh_CN":"订单ID"}
  OrderId *string `json:"orderId,omitempty" xml:"orderId,omitempty" require:"true"`
  // {"en":"Order ID List", "zh_CN":"订单ID列表"}
  OrderIds []*string `json:"orderIds,omitempty" xml:"orderIds,omitempty" require:"true" type:"Repeated"`
  // {"en":"Purchase Record ID List", "zh_CN":"采购记录ID列表"}
  PurchaseRecordIds []*string `json:"purchaseRecordIds,omitempty" xml:"purchaseRecordIds,omitempty" require:"true" type:"Repeated"`
}

func (s CreateCertificateApplyingOrderResponse) String() string {
  return tea.Prettify(s)
}

func (s CreateCertificateApplyingOrderResponse) GoString() string {
  return s.String()
}

func (s *CreateCertificateApplyingOrderResponse) SetCode(v string) *CreateCertificateApplyingOrderResponse {
  s.Code = &v
  return s
}

func (s *CreateCertificateApplyingOrderResponse) SetMessage(v string) *CreateCertificateApplyingOrderResponse {
  s.Message = &v
  return s
}

func (s *CreateCertificateApplyingOrderResponse) SetData(v map[string]interface{}) *CreateCertificateApplyingOrderResponse {
  s.Data = v
  return s
}

func (s *CreateCertificateApplyingOrderResponse) SetOrderId(v string) *CreateCertificateApplyingOrderResponse {
  s.OrderId = &v
  return s
}

func (s *CreateCertificateApplyingOrderResponse) SetOrderIds(v []*string) *CreateCertificateApplyingOrderResponse {
  s.OrderIds = v
  return s
}

func (s *CreateCertificateApplyingOrderResponse) SetPurchaseRecordIds(v []*string) *CreateCertificateApplyingOrderResponse {
  s.PurchaseRecordIds = v
  return s
}

type CreateCertificateApplyingOrderPaths struct {
}

func (s CreateCertificateApplyingOrderPaths) String() string {
  return tea.Prettify(s)
}

func (s CreateCertificateApplyingOrderPaths) GoString() string {
  return s.String()
}

type CreateCertificateApplyingOrderParameters struct {
}

func (s CreateCertificateApplyingOrderParameters) String() string {
  return tea.Prettify(s)
}

func (s CreateCertificateApplyingOrderParameters) GoString() string {
  return s.String()
}

type CreateCertificateApplyingOrderRequestHeader struct {
}

func (s CreateCertificateApplyingOrderRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s CreateCertificateApplyingOrderRequestHeader) GoString() string {
  return s.String()
}

type CreateCertificateApplyingOrderResponseHeader struct {
}

func (s CreateCertificateApplyingOrderResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s CreateCertificateApplyingOrderResponseHeader) GoString() string {
  return s.String()
}




type GetCertificateApplyingOrderListRequest struct {
  // {"en":"GetCertificateApplyingOrderListOrder ID", "zh_CN":"订单ID"}
  OrderId *string `json:"orderId,omitempty" xml:"orderId,omitempty"`
  // {"en":"Status of certificate apply order. Allowed values: ACCEPT_SUCCESS,APPLYING,ISSUE_SUCCESS,ISSUE_FAILURE,REVOKED,CANCELED. ", "zh_CN":"订单状态，取值范围：ACCEPT_SUCCESS,APPLYING,ISSUE_SUCCESS,ISSUE_FAILURE,REVOKED,CANCELED。
  //     ACCEPT_SUCCESS：订单提交成功，待后台发起证书申请；
  //     APPLYING：证书申请中，待域名验证；
  //     ISSUE_SUCCESS：证书签发成功，可进行证书配置与部署；
  //     ISSUE_FAILURE：证书签发失败；
  //     REVOKED：证书已吊销；
  //     CANCELED：已取消证书申请；
  //     "}
  OrderStatus []*string `json:"orderStatus,omitempty" xml:"orderStatus,omitempty" type:"Repeated"`
  // {"en":"Certificate Name ", "zh_CN":"证书名称"}
  CertificateName *string `json:"certificateName,omitempty" xml:"certificateName,omitempty"`
  // {"en":"Domain ", "zh_CN":"域名 "}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty"`
  // {"en":"The time when the certificate order was created. The format is yyyy-MM-dd HH:mm:ss ", "zh_CN":"订单创建时间，格式为：yyyy-MM-dd HH:mm:ss "}
  StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
  // {"en":"The time when the certificate order was finished. The format is yyyy-MM-dd HH:mm:ss ", "zh_CN":"订单结束时间，格式为：yyyy-MM-dd HH:mm:ss。证书签发成功或证书签发失败，则记录结束时间。 "}
  EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
  // {"en":"Page Parammeter.", "zh_CN":"分页参数"}
  GetCertificateApplyingOrderListPageParam *GetCertificateApplyingOrderListPageParam `json:"pageParam,omitempty" xml:"pageParam,omitempty"`
}

func (s GetCertificateApplyingOrderListRequest) String() string {
  return tea.Prettify(s)
}

func (s GetCertificateApplyingOrderListRequest) GoString() string {
  return s.String()
}

func (s *GetCertificateApplyingOrderListRequest) SetOrderId(v string) *GetCertificateApplyingOrderListRequest {
  s.OrderId = &v
  return s
}

func (s *GetCertificateApplyingOrderListRequest) SetOrderStatus(v []*string) *GetCertificateApplyingOrderListRequest {
  s.OrderStatus = v
  return s
}

func (s *GetCertificateApplyingOrderListRequest) SetCertificateName(v string) *GetCertificateApplyingOrderListRequest {
  s.CertificateName = &v
  return s
}

func (s *GetCertificateApplyingOrderListRequest) SetDomain(v string) *GetCertificateApplyingOrderListRequest {
  s.Domain = &v
  return s
}

func (s *GetCertificateApplyingOrderListRequest) SetStartTime(v string) *GetCertificateApplyingOrderListRequest {
  s.StartTime = &v
  return s
}

func (s *GetCertificateApplyingOrderListRequest) SetEndTime(v string) *GetCertificateApplyingOrderListRequest {
  s.EndTime = &v
  return s
}

func (s *GetCertificateApplyingOrderListRequest) SetPageParam(v *GetCertificateApplyingOrderListPageParam) *GetCertificateApplyingOrderListRequest {
  s.GetCertificateApplyingOrderListPageParam = v
  return s
}

type GetCertificateApplyingOrderListPageParam struct {
  // {"en":"Page Size,Range:1-100 (default:100)  ", "zh_CN":"每页大小，1-100，默认值100"}
  PageSize *int `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
  // {"en":"Page Number, the first page must be 0. 
  // Range:>=0 (default:0)  ", "zh_CN":"页码，当前页，第一页从0开始，必须大于等于0，默认值0"}
  PageNumber *int `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
}

func (s GetCertificateApplyingOrderListPageParam) String() string {
  return tea.Prettify(s)
}

func (s GetCertificateApplyingOrderListPageParam) GoString() string {
  return s.String()
}

func (s *GetCertificateApplyingOrderListPageParam) SetPageSize(v int) *GetCertificateApplyingOrderListPageParam {
  s.PageSize = &v
  return s
}

func (s *GetCertificateApplyingOrderListPageParam) SetPageNumber(v int) *GetCertificateApplyingOrderListPageParam {
  s.PageNumber = &v
  return s
}

type GetCertificateApplyingOrderListResponse struct {
  // {"en":"Result Code, success is 0 ", "zh_CN":"响应码，成功为0"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Result Message", "zh_CN":"响应信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Result Data", "zh_CN":"响应数据"}
  Data *GetCertificateApplyingOrderListResp `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s GetCertificateApplyingOrderListResponse) String() string {
  return tea.Prettify(s)
}

func (s GetCertificateApplyingOrderListResponse) GoString() string {
  return s.String()
}

func (s *GetCertificateApplyingOrderListResponse) SetCode(v string) *GetCertificateApplyingOrderListResponse {
  s.Code = &v
  return s
}

func (s *GetCertificateApplyingOrderListResponse) SetMessage(v string) *GetCertificateApplyingOrderListResponse {
  s.Message = &v
  return s
}

func (s *GetCertificateApplyingOrderListResponse) SetData(v *GetCertificateApplyingOrderListResp) *GetCertificateApplyingOrderListResponse {
  s.Data = v
  return s
}

type GetCertificateApplyingOrderListResp struct {
  // {"en":"GetCertificateApplyingOrderListOrder list", "zh_CN":"订单列表"}
  Orders []*GetCertificateApplyingOrderListOrder `json:"orders,omitempty" xml:"orders,omitempty" require:"true" type:"Repeated"`
  // {"en":"Page Info", "zh_CN":"分页数据"}
  GetCertificateApplyingOrderListPageInfo *GetCertificateApplyingOrderListPageInfo `json:"pageInfo,omitempty" xml:"pageInfo,omitempty" require:"true"`
}

func (s GetCertificateApplyingOrderListResp) String() string {
  return tea.Prettify(s)
}

func (s GetCertificateApplyingOrderListResp) GoString() string {
  return s.String()
}

func (s *GetCertificateApplyingOrderListResp) SetOrders(v []*GetCertificateApplyingOrderListOrder) *GetCertificateApplyingOrderListResp {
  s.Orders = v
  return s
}

func (s *GetCertificateApplyingOrderListResp) SetPageInfo(v *GetCertificateApplyingOrderListPageInfo) *GetCertificateApplyingOrderListResp {
  s.GetCertificateApplyingOrderListPageInfo = v
  return s
}

type GetCertificateApplyingOrderListOrder struct {
  // {"en":"GetCertificateApplyingOrderListOrder ID", "zh_CN":"订单ID"}
  OrderId *string `json:"orderId,omitempty" xml:"orderId,omitempty" require:"true"`
  // {"en":"Status of certificate apply order.
  // Allowed values: ACCEPT_SUCCESS,APPLYING,ISSUE_SUCCESS,ISSUE_FAILURE,REVOKED,CANCELED.  ", "zh_CN":"证书订单状态，取值范围：ACCEPT_SUCCESS,APPLYING,ISSUE_SUCCESS,ISSUE_FAILURE,REVOKED,CANCELED。"}
  OrderStatus *string `json:"orderStatus,omitempty" xml:"orderStatus,omitempty" require:"true"`
  // {"en":"Certificate ID ", "zh_CN":"证书ID "}
  CertificateId *string `json:"certificateId,omitempty" xml:"certificateId,omitempty" require:"true"`
  // {"en":"Certificate Name ", "zh_CN":"证书名称 "}
  CertificateName *string `json:"certificateName,omitempty" xml:"certificateName,omitempty" require:"true"`
  // {"en":"Certificate Brand ", "zh_CN":"证书品牌"}
  CertificateBrand *string `json:"certificateBrand,omitempty" xml:"certificateBrand,omitempty" require:"true"`
  // {"en":"Certificate Type ", "zh_CN":"证书类型 "}
  CertificateType *string `json:"certificateType,omitempty" xml:"certificateType,omitempty" require:"true"`
  // {"en":"Auto Renew ", "zh_CN":"是否自动续签 "}
  AutoRenew *string `json:"autoRenew,omitempty" xml:"autoRenew,omitempty" require:"true"`
  // {"en":"Description ", "zh_CN":"描述 "}
  Description *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
  // {"en":"Common Name of the certificate ", "zh_CN":"证书的通用名称 "}
  CommonName *string `json:"commonName,omitempty" xml:"commonName,omitempty" require:"true"`
  // {"en":"SAN ", "zh_CN":"主体备用名称 "}
  SubjectAlternativeNames *string `json:"subjectAlternativeNames,omitempty" xml:"subjectAlternativeNames,omitempty" require:"true"`
  // {"en":"Create Time ", "zh_CN":"创建时间 "}
  CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty" require:"true"`
}

func (s GetCertificateApplyingOrderListOrder) String() string {
  return tea.Prettify(s)
}

func (s GetCertificateApplyingOrderListOrder) GoString() string {
  return s.String()
}

func (s *GetCertificateApplyingOrderListOrder) SetOrderId(v string) *GetCertificateApplyingOrderListOrder {
  s.OrderId = &v
  return s
}

func (s *GetCertificateApplyingOrderListOrder) SetOrderStatus(v string) *GetCertificateApplyingOrderListOrder {
  s.OrderStatus = &v
  return s
}

func (s *GetCertificateApplyingOrderListOrder) SetCertificateId(v string) *GetCertificateApplyingOrderListOrder {
  s.CertificateId = &v
  return s
}

func (s *GetCertificateApplyingOrderListOrder) SetCertificateName(v string) *GetCertificateApplyingOrderListOrder {
  s.CertificateName = &v
  return s
}

func (s *GetCertificateApplyingOrderListOrder) SetCertificateBrand(v string) *GetCertificateApplyingOrderListOrder {
  s.CertificateBrand = &v
  return s
}

func (s *GetCertificateApplyingOrderListOrder) SetCertificateType(v string) *GetCertificateApplyingOrderListOrder {
  s.CertificateType = &v
  return s
}

func (s *GetCertificateApplyingOrderListOrder) SetAutoRenew(v string) *GetCertificateApplyingOrderListOrder {
  s.AutoRenew = &v
  return s
}

func (s *GetCertificateApplyingOrderListOrder) SetDescription(v string) *GetCertificateApplyingOrderListOrder {
  s.Description = &v
  return s
}

func (s *GetCertificateApplyingOrderListOrder) SetCommonName(v string) *GetCertificateApplyingOrderListOrder {
  s.CommonName = &v
  return s
}

func (s *GetCertificateApplyingOrderListOrder) SetSubjectAlternativeNames(v string) *GetCertificateApplyingOrderListOrder {
  s.SubjectAlternativeNames = &v
  return s
}

func (s *GetCertificateApplyingOrderListOrder) SetCreateTime(v string) *GetCertificateApplyingOrderListOrder {
  s.CreateTime = &v
  return s
}

type GetCertificateApplyingOrderListPageInfo struct {
  // {"en":"Total Number", "zh_CN":"总数"}
  TotalNumber *int `json:"totalNumber,omitempty" xml:"totalNumber,omitempty" require:"true"`
  // {"en":"Page Number", "zh_CN":"页码"}
  PageNumber *int `json:"pageNumber,omitempty" xml:"pageNumber,omitempty" require:"true"`
  // {"en":"Page Size", "zh_CN":"每页大小 "}
  PageSize *int `json:"pageSize,omitempty" xml:"pageSize,omitempty" require:"true"`
  // {"en":"Total Page Number", "zh_CN":"总页码数"}
  TotalPageNumber *int `json:"totalPageNumber,omitempty" xml:"totalPageNumber,omitempty" require:"true"`
}

func (s GetCertificateApplyingOrderListPageInfo) String() string {
  return tea.Prettify(s)
}

func (s GetCertificateApplyingOrderListPageInfo) GoString() string {
  return s.String()
}

func (s *GetCertificateApplyingOrderListPageInfo) SetTotalNumber(v int) *GetCertificateApplyingOrderListPageInfo {
  s.TotalNumber = &v
  return s
}

func (s *GetCertificateApplyingOrderListPageInfo) SetPageNumber(v int) *GetCertificateApplyingOrderListPageInfo {
  s.PageNumber = &v
  return s
}

func (s *GetCertificateApplyingOrderListPageInfo) SetPageSize(v int) *GetCertificateApplyingOrderListPageInfo {
  s.PageSize = &v
  return s
}

func (s *GetCertificateApplyingOrderListPageInfo) SetTotalPageNumber(v int) *GetCertificateApplyingOrderListPageInfo {
  s.TotalPageNumber = &v
  return s
}

type GetCertificateApplyingOrderListPaths struct {
}

func (s GetCertificateApplyingOrderListPaths) String() string {
  return tea.Prettify(s)
}

func (s GetCertificateApplyingOrderListPaths) GoString() string {
  return s.String()
}

type GetCertificateApplyingOrderListParameters struct {
}

func (s GetCertificateApplyingOrderListParameters) String() string {
  return tea.Prettify(s)
}

func (s GetCertificateApplyingOrderListParameters) GoString() string {
  return s.String()
}

type GetCertificateApplyingOrderListRequestHeader struct {
}

func (s GetCertificateApplyingOrderListRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetCertificateApplyingOrderListRequestHeader) GoString() string {
  return s.String()
}

type GetCertificateApplyingOrderListResponseHeader struct {
}

func (s GetCertificateApplyingOrderListResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetCertificateApplyingOrderListResponseHeader) GoString() string {
  return s.String()
}




type CancelCertificateApplyingOrderRequest struct {
  // {"en":"The orderId", "zh_CN":"订单ID"}
  OrderId *string `json:"orderId,omitempty" xml:"orderId,omitempty"`
  // {"en":"The purchaseRecordId", "zh_CN":"采购记录ID"}
  PurchaseRecordId *string `json:"purchaseRecordId,omitempty" xml:"purchaseRecordId,omitempty"`
}

func (s CancelCertificateApplyingOrderRequest) String() string {
  return tea.Prettify(s)
}

func (s CancelCertificateApplyingOrderRequest) GoString() string {
  return s.String()
}

func (s *CancelCertificateApplyingOrderRequest) SetOrderId(v string) *CancelCertificateApplyingOrderRequest {
  s.OrderId = &v
  return s
}

func (s *CancelCertificateApplyingOrderRequest) SetPurchaseRecordId(v string) *CancelCertificateApplyingOrderRequest {
  s.PurchaseRecordId = &v
  return s
}

type CancelCertificateApplyingOrderResponse struct {
  // {"en":"Result Code, 0 means successful ", "zh_CN":"响应码，成功为0"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Result Message", "zh_CN":"响应信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s CancelCertificateApplyingOrderResponse) String() string {
  return tea.Prettify(s)
}

func (s CancelCertificateApplyingOrderResponse) GoString() string {
  return s.String()
}

func (s *CancelCertificateApplyingOrderResponse) SetCode(v string) *CancelCertificateApplyingOrderResponse {
  s.Code = &v
  return s
}

func (s *CancelCertificateApplyingOrderResponse) SetMessage(v string) *CancelCertificateApplyingOrderResponse {
  s.Message = &v
  return s
}

type CancelCertificateApplyingOrderPaths struct {
}

func (s CancelCertificateApplyingOrderPaths) String() string {
  return tea.Prettify(s)
}

func (s CancelCertificateApplyingOrderPaths) GoString() string {
  return s.String()
}

type CancelCertificateApplyingOrderParameters struct {
}

func (s CancelCertificateApplyingOrderParameters) String() string {
  return tea.Prettify(s)
}

func (s CancelCertificateApplyingOrderParameters) GoString() string {
  return s.String()
}

type CancelCertificateApplyingOrderRequestHeader struct {
}

func (s CancelCertificateApplyingOrderRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s CancelCertificateApplyingOrderRequestHeader) GoString() string {
  return s.String()
}

type CancelCertificateApplyingOrderResponseHeader struct {
}

func (s CancelCertificateApplyingOrderResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s CancelCertificateApplyingOrderResponseHeader) GoString() string {
  return s.String()
}




