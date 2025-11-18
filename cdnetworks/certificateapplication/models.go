package certificateapplication

// This file is auto-generated, don't edit it. Thanks.
import (
	"github.com/alibabacloud-go/tea/tea"
)

type ReissueCertificateForWplusRequest struct {
	// {"en":"Certificate id","zh_CN":"证书ID"}
	CertificateId *int64 `json:"certificateId,omitempty" xml:"certificateId,omitempty" require:"true"`
	// {"en":"Certificate description","zh_CN":"证书描述"}
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// {"en":"Certificate algorithm","zh_CN":"证书算法","exampleValue":"RSA2048,ECDSA256"}
	Algorithm *string `json:"algorithm,omitempty" xml:"algorithm,omitempty" require:"true"`
	// {"en":"Certificate validate method","zh_CN":"验证方式","exampleValue":"HTTP,DNS"}
	ValidateMethod *string `json:"validateMethod,omitempty" xml:"validateMethod,omitempty" require:"true"`
	// {"en":"Is auto validate","zh_CN":"是否自动验证","exampleValue":"true,false"}
	AutoValidate *string `json:"autoValidate,omitempty" xml:"autoValidate,omitempty" require:"true"`
	// {"en":"Is auto deploy","zh_CN":"是否自动部署","exampleValue":"true,false"}
	AutoDeploy *string `json:"autoDeploy,omitempty" xml:"autoDeploy,omitempty" require:"true"`
	// {"en":"Common name","zh_CN":"通用名称"}
	CommonName *string `json:"commonName,omitempty" xml:"commonName,omitempty" require:"true"`
	// {"en":"Subject alternative name","zh_CN":"主体备用名称"}
	SubjectAlternativeNames []*string `json:"subjectAlternativeNames,omitempty" xml:"subjectAlternativeNames,omitempty" require:"true" type:"Repeated"`
	// {"en":"DNS Provider Infomation","zh_CN":"DNS托管商信息"}
	DnsProviderInfos []*ReissueCertificateForWplusRequestDnsProviderInfos `json:"dnsProviderInfos,omitempty" xml:"dnsProviderInfos,omitempty" type:"Repeated"`
	// {"en":"Auto Renew","zh_CN":"是否自动续签, true 自动续签 false 不自动续签","exampleValue":"true,false"}
	AutoRenew *string `json:"autoRenew,omitempty" xml:"autoRenew,omitempty"`
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

func (s *ReissueCertificateForWplusRequest) SetDnsProviderInfos(v []*ReissueCertificateForWplusRequestDnsProviderInfos) *ReissueCertificateForWplusRequest {
	s.DnsProviderInfos = v
	return s
}

func (s *ReissueCertificateForWplusRequest) SetAutoRenew(v string) *ReissueCertificateForWplusRequest {
	s.AutoRenew = &v
	return s
}

type ReissueCertificateForWplusRequestDnsProviderInfos struct {
	// {"en":"domain","zh_CN":"域名"}
	Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
	// {"defaultValue":"false","en":"Whether to use alias verification","zh_CN":"是否使用别名方式","exampleValue":"true,false"}
	EnableDnsAliasMode *string `json:"enableDnsAliasMode,omitempty" xml:"enableDnsAliasMode,omitempty"`
	// {"en":"Validate alias","zh_CN":"验证别名"}
	ValidateAliasDomain *string `json:"validateAliasDomain,omitempty" xml:"validateAliasDomain,omitempty"`
	// {"en":"DNS Provider Code,\nSupport CloudDNS, The optional certificate brand is LE or TrustAsia or GlobalSign or ZeroSSSL.","zh_CN":"DNS托管商编码，\n支持CloudDNS，可选证书品牌为LE、TrustAsia、GlobalSign、ZeroSSL。","exampleValue":"CloudDNS, CloudFlare"}
	DnsProviderCode *string `json:"dnsProviderCode,omitempty" xml:"dnsProviderCode,omitempty"`
	// {"en":"DNS Api Access, JSON format:\nThe hosting provider is CloudDNS, JSON KEY is accessKey and secretKey.","zh_CN":"DNS托管商API凭证，JSON格式，\n托管商为CloudDNS，JSON KEY为accessKey、secretKey。"}
	DnsApiAccess *string `json:"dnsApiAccess,omitempty" xml:"dnsApiAccess,omitempty"`
}

func (s ReissueCertificateForWplusRequestDnsProviderInfos) String() string {
	return tea.Prettify(s)
}

func (s ReissueCertificateForWplusRequestDnsProviderInfos) GoString() string {
	return s.String()
}

func (s *ReissueCertificateForWplusRequestDnsProviderInfos) SetDomain(v string) *ReissueCertificateForWplusRequestDnsProviderInfos {
	s.Domain = &v
	return s
}

func (s *ReissueCertificateForWplusRequestDnsProviderInfos) SetEnableDnsAliasMode(v string) *ReissueCertificateForWplusRequestDnsProviderInfos {
	s.EnableDnsAliasMode = &v
	return s
}

func (s *ReissueCertificateForWplusRequestDnsProviderInfos) SetValidateAliasDomain(v string) *ReissueCertificateForWplusRequestDnsProviderInfos {
	s.ValidateAliasDomain = &v
	return s
}

func (s *ReissueCertificateForWplusRequestDnsProviderInfos) SetDnsProviderCode(v string) *ReissueCertificateForWplusRequestDnsProviderInfos {
	s.DnsProviderCode = &v
	return s
}

func (s *ReissueCertificateForWplusRequestDnsProviderInfos) SetDnsApiAccess(v string) *ReissueCertificateForWplusRequestDnsProviderInfos {
	s.DnsApiAccess = &v
	return s
}

type ReissueCertificateForWplusRequestHeader struct {
}

func (s ReissueCertificateForWplusRequestHeader) String() string {
	return tea.Prettify(s)
}

func (s ReissueCertificateForWplusRequestHeader) GoString() string {
	return s.String()
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

type ReissueCertificateForWplusResponse struct {
	// {"en":"code","zh_CN":"错误码，成功为0"}
	Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
	// {"en":"error message","zh_CN":"错误信息"}
	Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
	// {"en":"Response Data","zh_CN":"响应数据"}
	Data *ReissueCertificateForWplusResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
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

func (s *ReissueCertificateForWplusResponse) SetData(v *ReissueCertificateForWplusResponseData) *ReissueCertificateForWplusResponse {
	s.Data = v
	return s
}

type ReissueCertificateForWplusResponseData struct {
	// {"en":"Sale order id","zh_CN":"销售订单id"}
	SalesOrderId *string `json:"salesOrderId,omitempty" xml:"salesOrderId,omitempty" require:"true"`
	// {"en":"Validate alias List","zh_CN":"验证别名列表"}
	ValidateAliasDomains []*ReissueCertificateForWplusResponseDataValidateAliasDomains `json:"validateAliasDomains,omitempty" xml:"validateAliasDomains,omitempty" require:"true" type:"Repeated"`
}

func (s ReissueCertificateForWplusResponseData) String() string {
	return tea.Prettify(s)
}

func (s ReissueCertificateForWplusResponseData) GoString() string {
	return s.String()
}

func (s *ReissueCertificateForWplusResponseData) SetSalesOrderId(v string) *ReissueCertificateForWplusResponseData {
	s.SalesOrderId = &v
	return s
}

func (s *ReissueCertificateForWplusResponseData) SetValidateAliasDomains(v []*ReissueCertificateForWplusResponseDataValidateAliasDomains) *ReissueCertificateForWplusResponseData {
	s.ValidateAliasDomains = v
	return s
}

type ReissueCertificateForWplusResponseDataValidateAliasDomains struct {
	// {"en":"domain","zh_CN":"域名"}
	Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
	// {"en":"Validate alias","zh_CN":"验证别名"}
	ValidateAliasDomain *string `json:"validateAliasDomain,omitempty" xml:"validateAliasDomain,omitempty" require:"true"`
}

func (s ReissueCertificateForWplusResponseDataValidateAliasDomains) String() string {
	return tea.Prettify(s)
}

func (s ReissueCertificateForWplusResponseDataValidateAliasDomains) GoString() string {
	return s.String()
}

func (s *ReissueCertificateForWplusResponseDataValidateAliasDomains) SetDomain(v string) *ReissueCertificateForWplusResponseDataValidateAliasDomains {
	s.Domain = &v
	return s
}

func (s *ReissueCertificateForWplusResponseDataValidateAliasDomains) SetValidateAliasDomain(v string) *ReissueCertificateForWplusResponseDataValidateAliasDomains {
	s.ValidateAliasDomain = &v
	return s
}

type ReissueCertificateForWplusResponseHeader struct {
}

func (s ReissueCertificateForWplusResponseHeader) String() string {
	return tea.Prettify(s)
}

func (s ReissueCertificateForWplusResponseHeader) GoString() string {
	return s.String()
}

type GetDcvContentRequest struct {
	// {"en":"Order ID","zh_CN":"订单ID"}
	OrderId *string `json:"orderId,omitempty" xml:"orderId,omitempty"`
	// {"en":"Purchase Record ID","zh_CN":"采购记录ID"}
	PurchaseRecordId *string `json:"purchaseRecordId,omitempty" xml:"purchaseRecordId,omitempty"`
}

func (s GetDcvContentRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDcvContentRequest) GoString() string {
	return s.String()
}

func (s *GetDcvContentRequest) SetOrderId(v string) *GetDcvContentRequest {
	s.OrderId = &v
	return s
}

func (s *GetDcvContentRequest) SetPurchaseRecordId(v string) *GetDcvContentRequest {
	s.PurchaseRecordId = &v
	return s
}

type GetDcvContentRequestHeader struct {
}

func (s GetDcvContentRequestHeader) String() string {
	return tea.Prettify(s)
}

func (s GetDcvContentRequestHeader) GoString() string {
	return s.String()
}

type GetDcvContentPaths struct {
}

func (s GetDcvContentPaths) String() string {
	return tea.Prettify(s)
}

func (s GetDcvContentPaths) GoString() string {
	return s.String()
}

type GetDcvContentParameters struct {
}

func (s GetDcvContentParameters) String() string {
	return tea.Prettify(s)
}

func (s GetDcvContentParameters) GoString() string {
	return s.String()
}

type GetDcvContentResponse struct {
	// {"en":"Result Code, success is 0","zh_CN":"响应码，成功为0"}
	Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
	// {"en":"Result Message","zh_CN":"响应信息"}
	Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
	// {"en":"Result Data","zh_CN":"响应数据"}
	Data *GetDcvContentResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
}

func (s GetDcvContentResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDcvContentResponse) GoString() string {
	return s.String()
}

func (s *GetDcvContentResponse) SetCode(v string) *GetDcvContentResponse {
	s.Code = &v
	return s
}

func (s *GetDcvContentResponse) SetMessage(v string) *GetDcvContentResponse {
	s.Message = &v
	return s
}

func (s *GetDcvContentResponse) SetData(v *GetDcvContentResponseData) *GetDcvContentResponse {
	s.Data = v
	return s
}

type GetDcvContentResponseData struct {
	// {"en":"Order ID","zh_CN":"订单ID"}
	OrderId *string `json:"orderId,omitempty" xml:"orderId,omitempty" require:"true"`
	// {"en":"Validate Method","zh_CN":"验证方式"}
	ValidateMethod *string `json:"validateMethod,omitempty" xml:"validateMethod,omitempty" require:"true"`
	// {"en":"Auto Validate","zh_CN":"是否自动验证"}
	AutoValidate *string `json:"autoValidate,omitempty" xml:"autoValidate,omitempty" require:"true"`
	// {"en":"Domain Validate Information","zh_CN":"域名验证内容列表"}
	DomainValidateInfos []*GetDcvContentResponseDataDomainValidateInfos `json:"domainValidateInfos,omitempty" xml:"domainValidateInfos,omitempty" require:"true" type:"Repeated"`
	// {"en":"Purchase Record ID","zh_CN":"采购记录ID"}
	PurchaseRecordId *string `json:"purchaseRecordId,omitempty" xml:"purchaseRecordId,omitempty" require:"true"`
}

func (s GetDcvContentResponseData) String() string {
	return tea.Prettify(s)
}

func (s GetDcvContentResponseData) GoString() string {
	return s.String()
}

func (s *GetDcvContentResponseData) SetOrderId(v string) *GetDcvContentResponseData {
	s.OrderId = &v
	return s
}

func (s *GetDcvContentResponseData) SetValidateMethod(v string) *GetDcvContentResponseData {
	s.ValidateMethod = &v
	return s
}

func (s *GetDcvContentResponseData) SetAutoValidate(v string) *GetDcvContentResponseData {
	s.AutoValidate = &v
	return s
}

func (s *GetDcvContentResponseData) SetDomainValidateInfos(v []*GetDcvContentResponseDataDomainValidateInfos) *GetDcvContentResponseData {
	s.DomainValidateInfos = v
	return s
}

func (s *GetDcvContentResponseData) SetPurchaseRecordId(v string) *GetDcvContentResponseData {
	s.PurchaseRecordId = &v
	return s
}

type GetDcvContentResponseDataDomainValidateInfos struct {
	// {"en":"Domain","zh_CN":"域名"}
	Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
	// {"en":"Validate File Name","zh_CN":"验证文件名"}
	ValidateFileName *string `json:"validateFileName,omitempty" xml:"validateFileName,omitempty" require:"true"`
	// {"en":"Validate value","zh_CN":"验证内容"}
	ValidateContent *string `json:"validateContent,omitempty" xml:"validateContent,omitempty" require:"true"`
	// {"en":"Expire Time","zh_CN":"过期时间"}
	ExpireTime *string `json:"expireTime,omitempty" xml:"expireTime,omitempty" require:"true"`
	// {"en":"Validate Method","zh_CN":"CA厂家实际的验证方式"}
	ValidateMethod *string `json:"validateMethod,omitempty" xml:"validateMethod,omitempty" require:"true"`
	// {"en":"Validate Status","zh_CN":"验证状态,{INIT: 初始化, VALIDATE_WAIT: 待验证, VALIDATE_PROCESSING: 验证中, VALIDATE_SUCCESS: 验证成功, VALIDATE_FAILURE: 验证失败}","exampleValue":"INIT,VALIDATE_WAIT,VALIDATE_PROCESSING,VALIDATE_SUCCESS,VALIDATE_FAILURE"}
	ValidateStatus *string `json:"validateStatus,omitempty" xml:"validateStatus,omitempty" require:"true"`
	// {"en":"recordType","zh_CN":"记录类型"}
	RecordType *string `json:"recordType,omitempty" xml:"recordType,omitempty" require:"true"`
	// {"en":"Validate Alias Domain","zh_CN":"记录值,有验证别名时，CNAME对应的记录值"}
	RecordValue *string `json:"recordValue,omitempty" xml:"recordValue,omitempty" require:"true"`
}

func (s GetDcvContentResponseDataDomainValidateInfos) String() string {
	return tea.Prettify(s)
}

func (s GetDcvContentResponseDataDomainValidateInfos) GoString() string {
	return s.String()
}

func (s *GetDcvContentResponseDataDomainValidateInfos) SetDomain(v string) *GetDcvContentResponseDataDomainValidateInfos {
	s.Domain = &v
	return s
}

func (s *GetDcvContentResponseDataDomainValidateInfos) SetValidateFileName(v string) *GetDcvContentResponseDataDomainValidateInfos {
	s.ValidateFileName = &v
	return s
}

func (s *GetDcvContentResponseDataDomainValidateInfos) SetValidateContent(v string) *GetDcvContentResponseDataDomainValidateInfos {
	s.ValidateContent = &v
	return s
}

func (s *GetDcvContentResponseDataDomainValidateInfos) SetExpireTime(v string) *GetDcvContentResponseDataDomainValidateInfos {
	s.ExpireTime = &v
	return s
}

func (s *GetDcvContentResponseDataDomainValidateInfos) SetValidateMethod(v string) *GetDcvContentResponseDataDomainValidateInfos {
	s.ValidateMethod = &v
	return s
}

func (s *GetDcvContentResponseDataDomainValidateInfos) SetValidateStatus(v string) *GetDcvContentResponseDataDomainValidateInfos {
	s.ValidateStatus = &v
	return s
}

func (s *GetDcvContentResponseDataDomainValidateInfos) SetRecordType(v string) *GetDcvContentResponseDataDomainValidateInfos {
	s.RecordType = &v
	return s
}

func (s *GetDcvContentResponseDataDomainValidateInfos) SetRecordValue(v string) *GetDcvContentResponseDataDomainValidateInfos {
	s.RecordValue = &v
	return s
}

type GetDcvContentResponseHeader struct {
}

func (s GetDcvContentResponseHeader) String() string {
	return tea.Prettify(s)
}

func (s GetDcvContentResponseHeader) GoString() string {
	return s.String()
}

type GetCertificateApplicationOrderRequest struct {
	// {"en":"orderId","zh_CN":"订单id"}
	OrderId *string `json:"orderId,omitempty" xml:"orderId,omitempty"`
	// {"en":"purchaseRecordId","zh_CN":"采购记录ID"}
	PurchaseRecordId *string `json:"purchaseRecordId,omitempty" xml:"purchaseRecordId,omitempty"`
}

func (s GetCertificateApplicationOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s GetCertificateApplicationOrderRequest) GoString() string {
	return s.String()
}

func (s *GetCertificateApplicationOrderRequest) SetOrderId(v string) *GetCertificateApplicationOrderRequest {
	s.OrderId = &v
	return s
}

func (s *GetCertificateApplicationOrderRequest) SetPurchaseRecordId(v string) *GetCertificateApplicationOrderRequest {
	s.PurchaseRecordId = &v
	return s
}

type GetCertificateApplicationOrderRequestHeader struct {
}

func (s GetCertificateApplicationOrderRequestHeader) String() string {
	return tea.Prettify(s)
}

func (s GetCertificateApplicationOrderRequestHeader) GoString() string {
	return s.String()
}

type GetCertificateApplicationOrderPaths struct {
}

func (s GetCertificateApplicationOrderPaths) String() string {
	return tea.Prettify(s)
}

func (s GetCertificateApplicationOrderPaths) GoString() string {
	return s.String()
}

type GetCertificateApplicationOrderParameters struct {
}

func (s GetCertificateApplicationOrderParameters) String() string {
	return tea.Prettify(s)
}

func (s GetCertificateApplicationOrderParameters) GoString() string {
	return s.String()
}

type GetCertificateApplicationOrderResponse struct {
	// {"en":"code","zh_CN":"错误码，成功为0"}
	Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
	// {"en":"error message","zh_CN":"错误信息"}
	Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
	// {"en":"data","zh_CN":"订单详情"}
	Data *GetCertificateApplicationOrderResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
}

func (s GetCertificateApplicationOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCertificateApplicationOrderResponse) GoString() string {
	return s.String()
}

func (s *GetCertificateApplicationOrderResponse) SetCode(v string) *GetCertificateApplicationOrderResponse {
	s.Code = &v
	return s
}

func (s *GetCertificateApplicationOrderResponse) SetMessage(v string) *GetCertificateApplicationOrderResponse {
	s.Message = &v
	return s
}

func (s *GetCertificateApplicationOrderResponse) SetData(v *GetCertificateApplicationOrderResponseData) *GetCertificateApplicationOrderResponse {
	s.Data = v
	return s
}

type GetCertificateApplicationOrderResponseData struct {
	// {"en":"orderId","zh_CN":"订单id"}
	OrderId *string `json:"orderId,omitempty" xml:"orderId,omitempty" require:"true"`
	// {"en":"orderStatus\nACCEPT_SUCCESS: Received successfully\nAPPLYING: Under application\nAPPLy_FAILURE: Application preparation failed\nVALIDATETWAIT: To be verified\nVALIDATE_PROCESSING: Verification processing\nVALIDATE_SUCCESS: Verification successful\nVALIDATE_SAILURE: Validation failed\nISSUEAwaIT: To be issued\nISSUE_SUCCESS: Issued successfully\nISSUE_SFAILURE: Issuance failed\nCANCELED: Cancel application\nREVOKED: Revocation\nDEPLO_SUCCESS: Deployment successful\nDEPLOy_Failures: Deployment failed","zh_CN":"订单状态\nACCEPT_SUCCESS: 接收成功\nAPPLYING: 申请中\nAPPLY_FAILURE：申请准备失败\nVALIDATE_WAIT: 待验证\nVALIDATE_PROCESSING: 验证中\nVALIDATE_SUCCESS: 验证成功\nVALIDATE_FAILURE: 验证失败\nISSUE_WAIT: 待签发\nISSUE_SUCCESS: 签发成功\nISSUE_FAILURE: 签发失败\nCANCELED: 取消申请\nREVOKED: 吊销\nDEPLOY_SUCCESS：部署成功\nDEPLOY_FAILURE：部署失败"}
	OrderStatus *string `json:"orderStatus,omitempty" xml:"orderStatus,omitempty" require:"true"`
	// {"en":"certificateId","zh_CN":"证书id"}
	CertificateId *int64 `json:"certificateId,omitempty" xml:"certificateId,omitempty" require:"true"`
	// {"en":"certificateName","zh_CN":"证书名"}
	CertificateName *string `json:"certificateName,omitempty" xml:"certificateName,omitempty" require:"true"`
	// {"en":"certificateBrand","zh_CN":"证书品牌"}
	CertificateBrand *string `json:"certificateBrand,omitempty" xml:"certificateBrand,omitempty" require:"true"`
	// {"en":"certificateType","zh_CN":"证书类型"}
	CertificateType *string `json:"certificateType,omitempty" xml:"certificateType,omitempty" require:"true"`
	// {"en":"description","zh_CN":"描述"}
	Description *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	// {"en":"algorithm","zh_CN":"算法"}
	Algorithm *string `json:"algorithm,omitempty" xml:"algorithm,omitempty" require:"true"`
	// {"en":"autoValidate","zh_CN":"是否自动验证"}
	AutoValidate *string `json:"autoValidate,omitempty" xml:"autoValidate,omitempty" require:"true"`
	// {"en":"validateMethod","zh_CN":"验证方式"}
	ValidateMethod *string `json:"validateMethod,omitempty" xml:"validateMethod,omitempty" require:"true"`
	// {"en":"autoRenew","zh_CN":"是否自动续订"}
	AutoRenew *string `json:"autoRenew,omitempty" xml:"autoRenew,omitempty" require:"true"`
	// {"en":"validityDays","zh_CN":"有效天数"}
	ValidityDays *int `json:"validityDays,omitempty" xml:"validityDays,omitempty" require:"true"`
	// {"en":"country","zh_CN":"国家代码"}
	Country *string `json:"country,omitempty" xml:"country,omitempty" require:"true"`
	// {"en":"state","zh_CN":"省"}
	State *string `json:"state,omitempty" xml:"state,omitempty" require:"true"`
	// {"en":"city","zh_CN":"市"}
	City *string `json:"city,omitempty" xml:"city,omitempty" require:"true"`
	// {"en":"street","zh_CN":"街道"}
	Street *string `json:"street,omitempty" xml:"street,omitempty" require:"true"`
	// {"en":"company","zh_CN":"公司"}
	Company *string `json:"company,omitempty" xml:"company,omitempty" require:"true"`
	// {"en":"department","zh_CN":"部门"}
	Department *string `json:"department,omitempty" xml:"department,omitempty" require:"true"`
	// {"en":"commonName","zh_CN":"通用名"}
	CommonName *string `json:"commonName,omitempty" xml:"commonName,omitempty" require:"true"`
	// {"en":"email","zh_CN":"邮箱"}
	Email *string `json:"email,omitempty" xml:"email,omitempty" require:"true"`
	// {"en":"subjectAlternativeNames","zh_CN":"主体备用名称"}
	SubjectAlternativeNames []*string `json:"subjectAlternativeNames,omitempty" xml:"subjectAlternativeNames,omitempty" require:"true" type:"Repeated"`
	// {"en":"createTime","zh_CN":"创建时间"}
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty" require:"true"`
	// {"en":"errorMessage","zh_CN":"错误信息"}
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty" require:"true"`
	// {"en":"certificateSpec","zh_CN":"证书规格"}
	CertificateSpec *string `json:"certificateSpec,omitempty" xml:"certificateSpec,omitempty" require:"true"`
	// {"en":"domainType","zh_CN":"域名类型"}
	DomainType *string `json:"domainType,omitempty" xml:"domainType,omitempty" require:"true"`
	// {"en":"autoDeploy","zh_CN":"是否自动部署"}
	AutoDeploy *string `json:"autoDeploy,omitempty" xml:"autoDeploy,omitempty" require:"true"`
	// {"en":"remainValidityDays","zh_CN":"剩余有效期"}
	RemainValidityDays *int `json:"remainValidityDays,omitempty" xml:"remainValidityDays,omitempty" require:"true"`
	// {"en":"dnsProviderInfos","zh_CN":"DNS托管商信息"}
	DnsProviderInfos []*GetCertificateApplicationOrderResponseDataDnsProviderInfos `json:"dnsProviderInfos,omitempty" xml:"dnsProviderInfos,omitempty" require:"true" type:"Repeated"`
	// {"en":"Primary Certificate Brand","zh_CN":"主证书品牌"}
	PrimaryCertificateBrand *string `json:"primaryCertificateBrand,omitempty" xml:"primaryCertificateBrand,omitempty" require:"true"`
	// {"en":"Backup Certificate Brand","zh_CN":"备用CA证书品牌"}
	BackupCertificateBrand *string `json:"backupCertificateBrand,omitempty" xml:"backupCertificateBrand,omitempty" require:"true"`
	// {"en":"Org Validate Method","zh_CN":"组织验证方式"}
	OrgValidateMethod *string `json:"orgValidateMethod,omitempty" xml:"orgValidateMethod,omitempty" require:"true"`
	// {"en":"Phone","zh_CN":"组织联系电话"}
	Phone *string `json:"phone,omitempty" xml:"phone,omitempty" require:"true"`
	// {"en":"Street","zh_CN":"街道2"}
	Street1 *string `json:"street1,omitempty" xml:"street1,omitempty" require:"true"`
	// {"en":"Postal Code","zh_CN":"邮编"}
	PostalCode *string `json:"postalCode,omitempty" xml:"postalCode,omitempty" require:"true"`
	// {"en":"admin","zh_CN":"订单管理联系人"}
	Admin *GetCertificateApplicationOrderResponseDataAdmin `json:"admin,omitempty" xml:"admin,omitempty" require:"true" type:"Struct"`
	// {"en":"Technical contact person for the order.","zh_CN":"订单技术联系人"}
	Tech *GetCertificateApplicationOrderResponseDataTech `json:"tech,omitempty" xml:"tech,omitempty" require:"true" type:"Struct"`
}

func (s GetCertificateApplicationOrderResponseData) String() string {
	return tea.Prettify(s)
}

func (s GetCertificateApplicationOrderResponseData) GoString() string {
	return s.String()
}

func (s *GetCertificateApplicationOrderResponseData) SetOrderId(v string) *GetCertificateApplicationOrderResponseData {
	s.OrderId = &v
	return s
}

func (s *GetCertificateApplicationOrderResponseData) SetOrderStatus(v string) *GetCertificateApplicationOrderResponseData {
	s.OrderStatus = &v
	return s
}

func (s *GetCertificateApplicationOrderResponseData) SetCertificateId(v int64) *GetCertificateApplicationOrderResponseData {
	s.CertificateId = &v
	return s
}

func (s *GetCertificateApplicationOrderResponseData) SetCertificateName(v string) *GetCertificateApplicationOrderResponseData {
	s.CertificateName = &v
	return s
}

func (s *GetCertificateApplicationOrderResponseData) SetCertificateBrand(v string) *GetCertificateApplicationOrderResponseData {
	s.CertificateBrand = &v
	return s
}

func (s *GetCertificateApplicationOrderResponseData) SetCertificateType(v string) *GetCertificateApplicationOrderResponseData {
	s.CertificateType = &v
	return s
}

func (s *GetCertificateApplicationOrderResponseData) SetDescription(v string) *GetCertificateApplicationOrderResponseData {
	s.Description = &v
	return s
}

func (s *GetCertificateApplicationOrderResponseData) SetAlgorithm(v string) *GetCertificateApplicationOrderResponseData {
	s.Algorithm = &v
	return s
}

func (s *GetCertificateApplicationOrderResponseData) SetAutoValidate(v string) *GetCertificateApplicationOrderResponseData {
	s.AutoValidate = &v
	return s
}

func (s *GetCertificateApplicationOrderResponseData) SetValidateMethod(v string) *GetCertificateApplicationOrderResponseData {
	s.ValidateMethod = &v
	return s
}

func (s *GetCertificateApplicationOrderResponseData) SetAutoRenew(v string) *GetCertificateApplicationOrderResponseData {
	s.AutoRenew = &v
	return s
}

func (s *GetCertificateApplicationOrderResponseData) SetValidityDays(v int) *GetCertificateApplicationOrderResponseData {
	s.ValidityDays = &v
	return s
}

func (s *GetCertificateApplicationOrderResponseData) SetCountry(v string) *GetCertificateApplicationOrderResponseData {
	s.Country = &v
	return s
}

func (s *GetCertificateApplicationOrderResponseData) SetState(v string) *GetCertificateApplicationOrderResponseData {
	s.State = &v
	return s
}

func (s *GetCertificateApplicationOrderResponseData) SetCity(v string) *GetCertificateApplicationOrderResponseData {
	s.City = &v
	return s
}

func (s *GetCertificateApplicationOrderResponseData) SetStreet(v string) *GetCertificateApplicationOrderResponseData {
	s.Street = &v
	return s
}

func (s *GetCertificateApplicationOrderResponseData) SetCompany(v string) *GetCertificateApplicationOrderResponseData {
	s.Company = &v
	return s
}

func (s *GetCertificateApplicationOrderResponseData) SetDepartment(v string) *GetCertificateApplicationOrderResponseData {
	s.Department = &v
	return s
}

func (s *GetCertificateApplicationOrderResponseData) SetCommonName(v string) *GetCertificateApplicationOrderResponseData {
	s.CommonName = &v
	return s
}

func (s *GetCertificateApplicationOrderResponseData) SetEmail(v string) *GetCertificateApplicationOrderResponseData {
	s.Email = &v
	return s
}

func (s *GetCertificateApplicationOrderResponseData) SetSubjectAlternativeNames(v []*string) *GetCertificateApplicationOrderResponseData {
	s.SubjectAlternativeNames = v
	return s
}

func (s *GetCertificateApplicationOrderResponseData) SetCreateTime(v string) *GetCertificateApplicationOrderResponseData {
	s.CreateTime = &v
	return s
}

func (s *GetCertificateApplicationOrderResponseData) SetErrorMessage(v string) *GetCertificateApplicationOrderResponseData {
	s.ErrorMessage = &v
	return s
}

func (s *GetCertificateApplicationOrderResponseData) SetCertificateSpec(v string) *GetCertificateApplicationOrderResponseData {
	s.CertificateSpec = &v
	return s
}

func (s *GetCertificateApplicationOrderResponseData) SetDomainType(v string) *GetCertificateApplicationOrderResponseData {
	s.DomainType = &v
	return s
}

func (s *GetCertificateApplicationOrderResponseData) SetAutoDeploy(v string) *GetCertificateApplicationOrderResponseData {
	s.AutoDeploy = &v
	return s
}

func (s *GetCertificateApplicationOrderResponseData) SetRemainValidityDays(v int) *GetCertificateApplicationOrderResponseData {
	s.RemainValidityDays = &v
	return s
}

func (s *GetCertificateApplicationOrderResponseData) SetDnsProviderInfos(v []*GetCertificateApplicationOrderResponseDataDnsProviderInfos) *GetCertificateApplicationOrderResponseData {
	s.DnsProviderInfos = v
	return s
}

func (s *GetCertificateApplicationOrderResponseData) SetPrimaryCertificateBrand(v string) *GetCertificateApplicationOrderResponseData {
	s.PrimaryCertificateBrand = &v
	return s
}

func (s *GetCertificateApplicationOrderResponseData) SetBackupCertificateBrand(v string) *GetCertificateApplicationOrderResponseData {
	s.BackupCertificateBrand = &v
	return s
}

func (s *GetCertificateApplicationOrderResponseData) SetOrgValidateMethod(v string) *GetCertificateApplicationOrderResponseData {
	s.OrgValidateMethod = &v
	return s
}

func (s *GetCertificateApplicationOrderResponseData) SetPhone(v string) *GetCertificateApplicationOrderResponseData {
	s.Phone = &v
	return s
}

func (s *GetCertificateApplicationOrderResponseData) SetStreet1(v string) *GetCertificateApplicationOrderResponseData {
	s.Street1 = &v
	return s
}

func (s *GetCertificateApplicationOrderResponseData) SetPostalCode(v string) *GetCertificateApplicationOrderResponseData {
	s.PostalCode = &v
	return s
}

func (s *GetCertificateApplicationOrderResponseData) SetAdmin(v *GetCertificateApplicationOrderResponseDataAdmin) *GetCertificateApplicationOrderResponseData {
	s.Admin = v
	return s
}

func (s *GetCertificateApplicationOrderResponseData) SetTech(v *GetCertificateApplicationOrderResponseDataTech) *GetCertificateApplicationOrderResponseData {
	s.Tech = v
	return s
}

type GetCertificateApplicationOrderResponseDataDnsProviderInfos struct {
	// {"en":"domain","zh_CN":"域名"}
	Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
	// {"en":"Support CloudDNS, The optional certificate brand is LE or TrustAsia or GlobalSign or ZeroSSL;Support CloudFlare, The optional certificate brand is LE or ZeroSSL;Support DNSPod.cn, The optional certificate brand is LE or ZeroSSL;Support DNSPod.com, The optional certificate brand is LE or ZeroSSL;Support TencentDNSPod, The optional certificate brand is LE or ZeroSSL;Support Aliyun, The optional certificate brand is LE or ZeroSSL;Support GoDaddy, The optional certificate brand is LE or ZeroSSL;Support AmazonRoute53, The optional certificate brand is LE or ZeroSSL;Support GoogleDomains, The optional certificate brand is LE or ZeroSSL;Support AkamaiEdgeDNS, The optional certificate brand is LE or ZeroSSL.","zh_CN":"DNS托管商编码，支持CloudDNS，可选证书品牌为LE、TrustAsia、GlobalSign、ZeroSSL；支持CloudFlare，可选证书品牌为LE、ZeroSSL；支持DNSPod.cn，可选证书品牌为LE、ZeroSSL；支持DNSPod.com，可选证书品牌为LE、ZeroSSL；支持TencentDNSPod，可选证书品牌为LE、ZeroSSL；支持Aliyun，可选证书品牌为LE、ZeroSSL；支持GoDaddy，可选证书品牌为LE、ZeroSSL；支持AmazonRoute53，可选证书品牌为LE、ZeroSSL；支持GoogleDomains，可选证书品牌为LE、ZeroSSL；支持AkamaiEdgeDNS，可选证书品牌为LE、ZeroSSL。","exampleValue":"CloudDNS,CloudFlare"}
	DnsProviderCode *string `json:"dnsProviderCode,omitempty" xml:"dnsProviderCode,omitempty" require:"true"`
	// {"en":"DNS Api Access, JSON format:The hosting provider is CloudDNS, JSON KEY is accessKey and secretKey;The hosting provider is CloudFlare, The global API key JSON KEY is CF_Key CF_Email, The restricted API token JSON KEY is CF_Token, CF_ZoneID or CF_Sccount-ID.","zh_CN":"DNS托管商API凭证，JSON格式，托管商为CloudDNS，JSON KEY为accessKey、secretKey；托管商为CloudFlare，全局API密钥JSON KEY为CF_Key、CF_Email，限制性API令牌JSON KEY为CF_Token、CF_Zone_ID或CF_Account_ID；托管商为DNSPod.cn，JSON KEY为DP_Id、DP_Key；托管商为DNSPod.com，JSON KEY为DPI_Id、DPI_Key；托管商为TencentDNSPod，JSON KEY为Tencent_SecretId、Tencent_SecretKey；托管商为Aliyun，JSON KEY为Ali_Key、Ali_Secret；托管商为GoDaddy，JSON KEY为GD_Key、GD_Secret；托管商为AmazonRoute53，JSON KEY为AWS_ACCESS_KEY_ID、AWS_SECRET_ACCESS_KEY；托管商为GoogleDomains，JSON KEY为GOOGLEDOMAINS_ACCESS_TOKEN、GOOGLEDOMAINS_ZONE；托管商为AkamaiEdgeDNS，JSON KEY为AKAMAI_CLIENT_TOKEN、AKAMAI_ACCESS_TOKEN、AKAMAI_CLIENT_SECRET、AKAMAI_HOST；"}
	DnsApiAccess *string `json:"dnsApiAccess,omitempty" xml:"dnsApiAccess,omitempty" require:"true"`
	// {"defaultValue":"false","en":"Whether to use alias verification","zh_CN":"是否使用别名方式","exampleValue":"true,false"}
	EnableDnsAliasMode *string `json:"enableDnsAliasMode,omitempty" xml:"enableDnsAliasMode,omitempty" require:"true"`
	// {"en":"Validate alias","zh_CN":"验证别名"}
	ValidateAliasDomain *string `json:"validateAliasDomain,omitempty" xml:"validateAliasDomain,omitempty" require:"true"`
}

func (s GetCertificateApplicationOrderResponseDataDnsProviderInfos) String() string {
	return tea.Prettify(s)
}

func (s GetCertificateApplicationOrderResponseDataDnsProviderInfos) GoString() string {
	return s.String()
}

func (s *GetCertificateApplicationOrderResponseDataDnsProviderInfos) SetDomain(v string) *GetCertificateApplicationOrderResponseDataDnsProviderInfos {
	s.Domain = &v
	return s
}

func (s *GetCertificateApplicationOrderResponseDataDnsProviderInfos) SetDnsProviderCode(v string) *GetCertificateApplicationOrderResponseDataDnsProviderInfos {
	s.DnsProviderCode = &v
	return s
}

func (s *GetCertificateApplicationOrderResponseDataDnsProviderInfos) SetDnsApiAccess(v string) *GetCertificateApplicationOrderResponseDataDnsProviderInfos {
	s.DnsApiAccess = &v
	return s
}

func (s *GetCertificateApplicationOrderResponseDataDnsProviderInfos) SetEnableDnsAliasMode(v string) *GetCertificateApplicationOrderResponseDataDnsProviderInfos {
	s.EnableDnsAliasMode = &v
	return s
}

func (s *GetCertificateApplicationOrderResponseDataDnsProviderInfos) SetValidateAliasDomain(v string) *GetCertificateApplicationOrderResponseDataDnsProviderInfos {
	s.ValidateAliasDomain = &v
	return s
}

type GetCertificateApplicationOrderResponseDataAdmin struct {
	// {"en":"firstName, Self verify organizational information. This field is mandatory","zh_CN":"名，自行验证组织信息此字段必填"}
	FirstName *string `json:"firstName,omitempty" xml:"firstName,omitempty" require:"true"`
	// {"en":"lastName, Self verify organizational information. This field is mandatory","zh_CN":"姓，自行验证组织信息此字段必填"}
	LastName *string `json:"lastName,omitempty" xml:"lastName,omitempty" require:"true"`
	// {"en":"phone, Self verify organizational information. This field is mandatory","zh_CN":"电话，自行验证组织信息此字段必填"}
	Phone *string `json:"phone,omitempty" xml:"phone,omitempty" require:"true"`
	// {"en":"email, Self verify organizational information. This field is mandatory","zh_CN":"邮箱，自行验证组织信息此字段必填"}
	Email *string `json:"email,omitempty" xml:"email,omitempty" require:"true"`
	// {"en":"title, Self verify organizational information. This field is mandatory","zh_CN":"职位，自行验证组织信息此字段必填"}
	Title *string `json:"title,omitempty" xml:"title,omitempty" require:"true"`
}

func (s GetCertificateApplicationOrderResponseDataAdmin) String() string {
	return tea.Prettify(s)
}

func (s GetCertificateApplicationOrderResponseDataAdmin) GoString() string {
	return s.String()
}

func (s *GetCertificateApplicationOrderResponseDataAdmin) SetFirstName(v string) *GetCertificateApplicationOrderResponseDataAdmin {
	s.FirstName = &v
	return s
}

func (s *GetCertificateApplicationOrderResponseDataAdmin) SetLastName(v string) *GetCertificateApplicationOrderResponseDataAdmin {
	s.LastName = &v
	return s
}

func (s *GetCertificateApplicationOrderResponseDataAdmin) SetPhone(v string) *GetCertificateApplicationOrderResponseDataAdmin {
	s.Phone = &v
	return s
}

func (s *GetCertificateApplicationOrderResponseDataAdmin) SetEmail(v string) *GetCertificateApplicationOrderResponseDataAdmin {
	s.Email = &v
	return s
}

func (s *GetCertificateApplicationOrderResponseDataAdmin) SetTitle(v string) *GetCertificateApplicationOrderResponseDataAdmin {
	s.Title = &v
	return s
}

type GetCertificateApplicationOrderResponseDataTech struct {
	// {"en":"firstName, Self verify organizational information. This field is mandatory","zh_CN":"名，自行验证组织信息此字段必填"}
	FirstName *string `json:"firstName,omitempty" xml:"firstName,omitempty" require:"true"`
	// {"en":"lastName, Self verify organizational information. This field is mandatory","zh_CN":"姓，自行验证组织信息此字段必填"}
	LastName *string `json:"lastName,omitempty" xml:"lastName,omitempty" require:"true"`
	// {"en":"phone, Self verify organizational information. This field is mandatory","zh_CN":"电话，自行验证组织信息此字段必填"}
	Phone *string `json:"phone,omitempty" xml:"phone,omitempty" require:"true"`
	// {"en":"email, Self verify organizational information. This field is mandatory","zh_CN":"邮箱，自行验证组织信息此字段必填"}
	Email *string `json:"email,omitempty" xml:"email,omitempty" require:"true"`
	// {"en":"title, Self verify organizational information. This field is mandatory","zh_CN":"职位，自行验证组织信息此字段必填"}
	Title *string `json:"title,omitempty" xml:"title,omitempty" require:"true"`
}

func (s GetCertificateApplicationOrderResponseDataTech) String() string {
	return tea.Prettify(s)
}

func (s GetCertificateApplicationOrderResponseDataTech) GoString() string {
	return s.String()
}

func (s *GetCertificateApplicationOrderResponseDataTech) SetFirstName(v string) *GetCertificateApplicationOrderResponseDataTech {
	s.FirstName = &v
	return s
}

func (s *GetCertificateApplicationOrderResponseDataTech) SetLastName(v string) *GetCertificateApplicationOrderResponseDataTech {
	s.LastName = &v
	return s
}

func (s *GetCertificateApplicationOrderResponseDataTech) SetPhone(v string) *GetCertificateApplicationOrderResponseDataTech {
	s.Phone = &v
	return s
}

func (s *GetCertificateApplicationOrderResponseDataTech) SetEmail(v string) *GetCertificateApplicationOrderResponseDataTech {
	s.Email = &v
	return s
}

func (s *GetCertificateApplicationOrderResponseDataTech) SetTitle(v string) *GetCertificateApplicationOrderResponseDataTech {
	s.Title = &v
	return s
}

type GetCertificateApplicationOrderResponseHeader struct {
}

func (s GetCertificateApplicationOrderResponseHeader) String() string {
	return tea.Prettify(s)
}

func (s GetCertificateApplicationOrderResponseHeader) GoString() string {
	return s.String()
}

type CreateCertificateApplicationOrderRequest struct {
	// {"en":"Name of the certificate. The certificate name cannot be the same as your existing certificate.\nRange: <= 128 characters.","zh_CN":"证书名称，最长不超过128个字符。证书名称不允许与已有的证书重名."}
	CertificateName *string `json:"certificateName,omitempty" xml:"certificateName,omitempty"`
	// {"en":"A description of the new certificate.\nRange: <= 256 characters.","zh_CN":"证书描述，最长不超过256个字符。"}
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// {"defaultValue":"false","en":"Automatically Renew your certificate. Selecting 'true' will reduce customization of other fields: 'validateMethod' must be 'HTTP', the certificate domain must configure in our acceleration platform; the domain  must CNAME to our acceleration platform.","zh_CN":"是否自动续签。选择是否在证书到期前为您自动续订证书，开启自动续签，验证方式必须选择'HTTP'；申请证书包含的域名必须已在我司加速平台配置，且域名必须已CNAME到我司加速平台。","exampleValue":"true,false"}
	AutoRenew *string `json:"autoRenew,omitempty" xml:"autoRenew,omitempty"`
	// {"en":"Certificate Brand. Selecting 'LE' means we apply your certificate through Let's Encrypt (https://letsencrypt.org/docs/challenge-types/) .","zh_CN":"证书品牌。LE即Let's Encrypt。","exampleValue":"LE,TrustAsia,GlobalSign"}
	CertificateBrand *string `json:"certificateBrand,omitempty" xml:"certificateBrand,omitempty" require:"true"`
	// {"en":"Certificate Specification,\nCertificate brand GlobalSign, certificate specification GlobalSignOVMSSSL, certificate type OV, certificate validity days [360], billing;\nCertificate brand GlobalSign, certificate specification GlobalSignOVSANsMSSL, certificate type OV, certificate validity days [360], billing;\nCertificate brand TrustAsia, certificate specification TrustAsia TLSC1, certificate type DV, certificate validity days [360], billing;\nCertificate brand TrustAsia, certificate specification TrustAsiaTLSSANsC1, certificate type DV, certificate validity days [360], billing;\nCertificate brand TrustAsia, certificate specification TrustAsiaC1DVFree, certificate type DV, certificate validity days [90], free of charge;\nCertificate brand TrustAsia, certificate specification TrustAsiaC1DVFreeSANS, certificate type DV, certificate validity days [90], free of charge;\nCertificate brand LE, certificate specification LetsEncryptDVFree, certificate type DV, certificate validity days [90], free of charge;\nCertificate Brand LE, Certificate Specification LetsEncryptDVFreeSANS, Certificate Type DV, Certificate Validity Days [90], free of charge","zh_CN":"证书规格,\n证书品牌GlobalSign、证书规格GlobalSignOVMSSL、证书类型OV、证书有效天数[360]、计费;\n证书品牌GlobalSign、证书规格GlobalSignOVSANsMSSL、证书类型OV、证书有效天数[360]、计费;\n证书品牌TrustAsia、证书规格TrustAsiaTLSC1、证书类型DV、证书有效天数[360]、计费;\n证书品牌TrustAsia、证书规格TrustAsiaTLSSANsC1、证书类型DV、证书有效天数[360]、计费;\n证书品牌TrustAsia、证书规格TrustAsiaC1DVFree、证书类型DV、证书有效天数[90]、免费;\n证书品牌TrustAsia、证书规格TrustAsiaC1DVFreeSANS、证书类型DV、证书有效天数[90]、免费;\n证书品牌LE、证书规格LetsEncryptDVFree、证书类型DV、证书有效天数[90]、免费;\n证书品牌LE、证书规格LetsEncryptDVFreeSANS、证书类型DV、证书有效天数[90]、免费","exampleValue":"GlobalSignOVMSSL,GlobalSignOVSANsMSSL,TrustAsiaTLSC1,TrustAsiaTLSSANsC1,TrustAsiaC1DVFree,TrustAsiaC1DVFreeSANS,LetsEncryptDVFree,LetsEncryptDVFreeSANS"}
	CertificateSpec *string `json:"certificateSpec,omitempty" xml:"certificateSpec,omitempty" require:"true"`
	// {"en":"Certificate Type. When 'certificateBrand' specifying 'LE',  'certificateType' only supports 'DV'.","zh_CN":"安全等级。 Let's Encrypt仅支持DV， TrustAsia支持DV和OV。","exampleValue":"DV,OV"}
	CertificateType *string `json:"certificateType,omitempty" xml:"certificateType,omitempty" require:"true"`
	// {"en":"Certificate Algorithm. When 'certificateBrand' specifying 'LE',  'certificateType' only supports 'RSA2048' and 'ECDSA256'.","zh_CN":"证书算法。Let's Encrypt仅支持RSA2048与ECDSA256  。","exampleValue":"RSA2048,ECDSA256"}
	Algorithm *string `json:"algorithm,omitempty" xml:"algorithm,omitempty" require:"true"`
	// {"en":"Automatically validate that you control the domains in the certificate you applied. Selecting 'true' will reduce customization of other fields: 'validateMethod' must be 'HTTP' , the certificate domain must configure in our acceleration platform; the domain must CNAME to our acceleration platform.","zh_CN":"是否自动验证。选择是否帮您自动完成证书域名控制权验证。开启自动验证，验证方式必须选择'HTTP'；申请证书包含的域名必须已在我司加速平台配置，且域名必须已CNAME到我司加速平台。","exampleValue":"true,false"}
	AutoValidate *string `json:"autoValidate,omitempty" xml:"autoValidate,omitempty" require:"true"`
	// {"en":"Validate Method selected for the order to validate that you control the domains in the certificate. Validate Method must be 'DNS' when 'subjectAlternativeNames'  is a wildcard hostname beginning with '*' such as '*.example.com' .\nWhen applying for a certificate with an IP domain name, only HTTP authentication is supported.","zh_CN":"验证方式。选择证书域名控制权的验证方式。当申请证书域名'subjectAlternativeNames'为泛域名时，仅支持DNS验证。当申请证书域名为IP时，仅支持HTTP验证。","exampleValue":"HTTP,DNS"}
	ValidateMethod *string `json:"validateMethod,omitempty" xml:"validateMethod,omitempty" require:"true"`
	// {"en":"Whether to deploy automatically. true: automatic deployment, false: manual deployment.","zh_CN":"是否自动部署，true:自动部署，false:手动部署.","exampleValue":"true,false"}
	AutoDeploy *string `json:"autoDeploy,omitempty" xml:"autoDeploy,omitempty" require:"true"`
	// {"en":"Validity Days selected for the certificate. Please refer to the description in the certificateSpec for details.","zh_CN":"证书有效天数，详见certificateSpec描述"}
	ValidityDays *int `json:"validityDays,omitempty" xml:"validityDays,omitempty" require:"true"`
	// {"en":"Certificate Signing Request Informtion.","zh_CN":"证书签名请求信息 。"}
	IdentificationInfo *CreateCertificateApplicationOrderRequestIdentificationInfo `json:"identificationInfo,omitempty" xml:"identificationInfo,omitempty" require:"true" type:"Struct"`
	// {"en":"Backup Certificate Brand.\n\nSingle Domain:\n- LE (Backup: TrustAsia, ZeroSSL)\n- ZeroSSL (Backup: TrustAsia, LE)\n- TrustAsia (Backup: LE, ZeroSSL)\n\nMulti-Domain:\n- LE (Backup: TrustAsia)\n- TrustAsia (Backup: LE)","zh_CN":"备用CA证书品牌。\n\n单域名:\n- LE (备用: TrustAsia, ZeroSSL)\n- ZeroSSL (备用: TrustAsia, LE)\n- TrustAsia (备用: LE, ZeroSSL)\n\n多域名:\n- LE (备用: TrustAsia)\n- TrustAsia (备用: LE)","exampleValue":"LE,TrustAsia,ZeroSSL"}
	BackupCertificateBrand *string `json:"backupCertificateBrand,omitempty" xml:"backupCertificateBrand,omitempty"`
	// {"en":"orgValidateMethod, default:  Enable pre audit, self_validate:  Self verify organizational information, none: Unorganized information, default value default.","zh_CN":"组织验证方式, default: 开启预审核、self_validate: 自行验证组织信息、none: 无组织信息，默认值default。"}
	OrgValidateMethod *string `json:"orgValidateMethod,omitempty" xml:"orgValidateMethod,omitempty"`
	// {"en":"domainType","zh_CN":"域名类型","exampleValue":"single,multi"}
	DomainType *string `json:"domainType,omitempty" xml:"domainType,omitempty"`
	// {"defaultValue":"false","en":"batchApply","zh_CN":"是否批量申请","exampleValue":"true,false"}
	BatchApply *string `json:"batchApply,omitempty" xml:"batchApply,omitempty"`
	// {"en":"admin","zh_CN":"订单管理联系人"}
	Admin *CreateCertificateApplicationOrderRequestAdmin `json:"admin,omitempty" xml:"admin,omitempty" type:"Struct"`
	// {"en":"tech","zh_CN":"订单技术联系人"}
	Tech *CreateCertificateApplicationOrderRequestTech `json:"tech,omitempty" xml:"tech,omitempty" type:"Struct"`
	// {"en":"DNS Provider Infomation","zh_CN":"DNS托管商信息"}
	DnsProviderInfos []*CreateCertificateApplicationOrderRequestDnsProviderInfos `json:"dnsProviderInfos,omitempty" xml:"dnsProviderInfos,omitempty" type:"Repeated"`
}

func (s CreateCertificateApplicationOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateCertificateApplicationOrderRequest) GoString() string {
	return s.String()
}

func (s *CreateCertificateApplicationOrderRequest) SetCertificateName(v string) *CreateCertificateApplicationOrderRequest {
	s.CertificateName = &v
	return s
}

func (s *CreateCertificateApplicationOrderRequest) SetDescription(v string) *CreateCertificateApplicationOrderRequest {
	s.Description = &v
	return s
}

func (s *CreateCertificateApplicationOrderRequest) SetAutoRenew(v string) *CreateCertificateApplicationOrderRequest {
	s.AutoRenew = &v
	return s
}

func (s *CreateCertificateApplicationOrderRequest) SetCertificateBrand(v string) *CreateCertificateApplicationOrderRequest {
	s.CertificateBrand = &v
	return s
}

func (s *CreateCertificateApplicationOrderRequest) SetCertificateSpec(v string) *CreateCertificateApplicationOrderRequest {
	s.CertificateSpec = &v
	return s
}

func (s *CreateCertificateApplicationOrderRequest) SetCertificateType(v string) *CreateCertificateApplicationOrderRequest {
	s.CertificateType = &v
	return s
}

func (s *CreateCertificateApplicationOrderRequest) SetAlgorithm(v string) *CreateCertificateApplicationOrderRequest {
	s.Algorithm = &v
	return s
}

func (s *CreateCertificateApplicationOrderRequest) SetAutoValidate(v string) *CreateCertificateApplicationOrderRequest {
	s.AutoValidate = &v
	return s
}

func (s *CreateCertificateApplicationOrderRequest) SetValidateMethod(v string) *CreateCertificateApplicationOrderRequest {
	s.ValidateMethod = &v
	return s
}

func (s *CreateCertificateApplicationOrderRequest) SetAutoDeploy(v string) *CreateCertificateApplicationOrderRequest {
	s.AutoDeploy = &v
	return s
}

func (s *CreateCertificateApplicationOrderRequest) SetValidityDays(v int) *CreateCertificateApplicationOrderRequest {
	s.ValidityDays = &v
	return s
}

func (s *CreateCertificateApplicationOrderRequest) SetIdentificationInfo(v *CreateCertificateApplicationOrderRequestIdentificationInfo) *CreateCertificateApplicationOrderRequest {
	s.IdentificationInfo = v
	return s
}

func (s *CreateCertificateApplicationOrderRequest) SetBackupCertificateBrand(v string) *CreateCertificateApplicationOrderRequest {
	s.BackupCertificateBrand = &v
	return s
}

func (s *CreateCertificateApplicationOrderRequest) SetOrgValidateMethod(v string) *CreateCertificateApplicationOrderRequest {
	s.OrgValidateMethod = &v
	return s
}

func (s *CreateCertificateApplicationOrderRequest) SetDomainType(v string) *CreateCertificateApplicationOrderRequest {
	s.DomainType = &v
	return s
}

func (s *CreateCertificateApplicationOrderRequest) SetBatchApply(v string) *CreateCertificateApplicationOrderRequest {
	s.BatchApply = &v
	return s
}

func (s *CreateCertificateApplicationOrderRequest) SetAdmin(v *CreateCertificateApplicationOrderRequestAdmin) *CreateCertificateApplicationOrderRequest {
	s.Admin = v
	return s
}

func (s *CreateCertificateApplicationOrderRequest) SetTech(v *CreateCertificateApplicationOrderRequestTech) *CreateCertificateApplicationOrderRequest {
	s.Tech = v
	return s
}

func (s *CreateCertificateApplicationOrderRequest) SetDnsProviderInfos(v []*CreateCertificateApplicationOrderRequestDnsProviderInfos) *CreateCertificateApplicationOrderRequest {
	s.DnsProviderInfos = v
	return s
}

type CreateCertificateApplicationOrderRequestIdentificationInfo struct {
	// {"en":"Country. An ISO-3166 country code.\nRange: = 2 characters","zh_CN":"国家，2个字符的ISO-3166国家代码"}
	Country *string `json:"country,omitempty" xml:"country,omitempty"`
	// {"en":"A state or province.","zh_CN":"省"}
	State *string `json:"state,omitempty" xml:"state,omitempty"`
	// {"en":"A city","zh_CN":"城市"}
	City *string `json:"city,omitempty" xml:"city,omitempty"`
	// {"en":"A company name","zh_CN":"公司"}
	Company *string `json:"company,omitempty" xml:"company,omitempty"`
	// {"en":"The department associated with the certificate","zh_CN":"部门"}
	Department *string `json:"department,omitempty" xml:"department,omitempty"`
	// {"en":"A common name of the certificate","zh_CN":"证书的通用名称"}
	CommonName *string `json:"commonName,omitempty" xml:"commonName,omitempty" require:"true"`
	// {"en":"Email address","zh_CN":"邮箱"}
	Email *string `json:"email,omitempty" xml:"email,omitempty"`
	// {"en":"The street where the company is located","zh_CN":"街道"}
	Street *string `json:"street,omitempty" xml:"street,omitempty"`
	// {"en":"Hostnames that this certificate will serve. Each entry must be a valid hostname such as example.com or a wildcard hostname beginning with '*' such as.example.com.","zh_CN":"此证书将提供服务的域名。 每个条目必须是有效的域名（例如 example.com）或以“*”开头的泛域名（例如 “*.example.com”）"}
	SubjectAlternativeNames []*string `json:"subjectAlternativeNames,omitempty" xml:"subjectAlternativeNames,omitempty" require:"true" type:"Repeated"`
	// {"en":"street1","zh_CN":"街道"}
	Street1 *string `json:"street1,omitempty" xml:"street1,omitempty"`
	// {"en":"phone","zh_CN":"联系电话"}
	Phone *string `json:"phone,omitempty" xml:"phone,omitempty"`
	// {"en":"postalCode","zh_CN":"邮政编码"}
	PostalCode *string `json:"postalCode,omitempty" xml:"postalCode,omitempty"`
}

func (s CreateCertificateApplicationOrderRequestIdentificationInfo) String() string {
	return tea.Prettify(s)
}

func (s CreateCertificateApplicationOrderRequestIdentificationInfo) GoString() string {
	return s.String()
}

func (s *CreateCertificateApplicationOrderRequestIdentificationInfo) SetCountry(v string) *CreateCertificateApplicationOrderRequestIdentificationInfo {
	s.Country = &v
	return s
}

func (s *CreateCertificateApplicationOrderRequestIdentificationInfo) SetState(v string) *CreateCertificateApplicationOrderRequestIdentificationInfo {
	s.State = &v
	return s
}

func (s *CreateCertificateApplicationOrderRequestIdentificationInfo) SetCity(v string) *CreateCertificateApplicationOrderRequestIdentificationInfo {
	s.City = &v
	return s
}

func (s *CreateCertificateApplicationOrderRequestIdentificationInfo) SetCompany(v string) *CreateCertificateApplicationOrderRequestIdentificationInfo {
	s.Company = &v
	return s
}

func (s *CreateCertificateApplicationOrderRequestIdentificationInfo) SetDepartment(v string) *CreateCertificateApplicationOrderRequestIdentificationInfo {
	s.Department = &v
	return s
}

func (s *CreateCertificateApplicationOrderRequestIdentificationInfo) SetCommonName(v string) *CreateCertificateApplicationOrderRequestIdentificationInfo {
	s.CommonName = &v
	return s
}

func (s *CreateCertificateApplicationOrderRequestIdentificationInfo) SetEmail(v string) *CreateCertificateApplicationOrderRequestIdentificationInfo {
	s.Email = &v
	return s
}

func (s *CreateCertificateApplicationOrderRequestIdentificationInfo) SetStreet(v string) *CreateCertificateApplicationOrderRequestIdentificationInfo {
	s.Street = &v
	return s
}

func (s *CreateCertificateApplicationOrderRequestIdentificationInfo) SetSubjectAlternativeNames(v []*string) *CreateCertificateApplicationOrderRequestIdentificationInfo {
	s.SubjectAlternativeNames = v
	return s
}

func (s *CreateCertificateApplicationOrderRequestIdentificationInfo) SetStreet1(v string) *CreateCertificateApplicationOrderRequestIdentificationInfo {
	s.Street1 = &v
	return s
}

func (s *CreateCertificateApplicationOrderRequestIdentificationInfo) SetPhone(v string) *CreateCertificateApplicationOrderRequestIdentificationInfo {
	s.Phone = &v
	return s
}

func (s *CreateCertificateApplicationOrderRequestIdentificationInfo) SetPostalCode(v string) *CreateCertificateApplicationOrderRequestIdentificationInfo {
	s.PostalCode = &v
	return s
}

type CreateCertificateApplicationOrderRequestAdmin struct {
	// {"en":"firstName, Self verify organizational information. This field is mandatory","zh_CN":"名，自行验证组织信息此字段必填"}
	FirstName *string `json:"firstName,omitempty" xml:"firstName,omitempty"`
	// {"en":"lastName, Self verify organizational information. This field is mandatory","zh_CN":"姓，自行验证组织信息此字段必填"}
	LastName *string `json:"lastName,omitempty" xml:"lastName,omitempty"`
	// {"en":"phone, Self verify organizational information. This field is mandatory","zh_CN":"电话，自行验证组织信息此字段必填"}
	Phone *string `json:"phone,omitempty" xml:"phone,omitempty"`
	// {"en":"email, Self verify organizational information. This field is mandatory","zh_CN":"邮箱，自行验证组织信息此字段必填"}
	Email *string `json:"email,omitempty" xml:"email,omitempty"`
	// {"en":"title, Self verify organizational information. This field is mandatory","zh_CN":"职位，自行验证组织信息此字段必填"}
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s CreateCertificateApplicationOrderRequestAdmin) String() string {
	return tea.Prettify(s)
}

func (s CreateCertificateApplicationOrderRequestAdmin) GoString() string {
	return s.String()
}

func (s *CreateCertificateApplicationOrderRequestAdmin) SetFirstName(v string) *CreateCertificateApplicationOrderRequestAdmin {
	s.FirstName = &v
	return s
}

func (s *CreateCertificateApplicationOrderRequestAdmin) SetLastName(v string) *CreateCertificateApplicationOrderRequestAdmin {
	s.LastName = &v
	return s
}

func (s *CreateCertificateApplicationOrderRequestAdmin) SetPhone(v string) *CreateCertificateApplicationOrderRequestAdmin {
	s.Phone = &v
	return s
}

func (s *CreateCertificateApplicationOrderRequestAdmin) SetEmail(v string) *CreateCertificateApplicationOrderRequestAdmin {
	s.Email = &v
	return s
}

func (s *CreateCertificateApplicationOrderRequestAdmin) SetTitle(v string) *CreateCertificateApplicationOrderRequestAdmin {
	s.Title = &v
	return s
}

type CreateCertificateApplicationOrderRequestTech struct {
	// {"en":"firstName, Self verify organizational information. This field is mandatory","zh_CN":"名，自行验证组织信息此字段必填"}
	FirstName *string `json:"firstName,omitempty" xml:"firstName,omitempty"`
	// {"en":"lastName, Self verify organizational information. This field is mandatory","zh_CN":"姓，自行验证组织信息此字段必填"}
	LastName *string `json:"lastName,omitempty" xml:"lastName,omitempty"`
	// {"en":"phone, Self verify organizational information. This field is mandatory","zh_CN":"电话，自行验证组织信息此字段必填"}
	Phone *string `json:"phone,omitempty" xml:"phone,omitempty"`
	// {"en":"email, Self verify organizational information. This field is mandatory","zh_CN":"邮箱，自行验证组织信息此字段必填"}
	Email *string `json:"email,omitempty" xml:"email,omitempty"`
	// {"en":"title, Self verify organizational information. This field is mandatory","zh_CN":"职位，自行验证组织信息此字段必填"}
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s CreateCertificateApplicationOrderRequestTech) String() string {
	return tea.Prettify(s)
}

func (s CreateCertificateApplicationOrderRequestTech) GoString() string {
	return s.String()
}

func (s *CreateCertificateApplicationOrderRequestTech) SetFirstName(v string) *CreateCertificateApplicationOrderRequestTech {
	s.FirstName = &v
	return s
}

func (s *CreateCertificateApplicationOrderRequestTech) SetLastName(v string) *CreateCertificateApplicationOrderRequestTech {
	s.LastName = &v
	return s
}

func (s *CreateCertificateApplicationOrderRequestTech) SetPhone(v string) *CreateCertificateApplicationOrderRequestTech {
	s.Phone = &v
	return s
}

func (s *CreateCertificateApplicationOrderRequestTech) SetEmail(v string) *CreateCertificateApplicationOrderRequestTech {
	s.Email = &v
	return s
}

func (s *CreateCertificateApplicationOrderRequestTech) SetTitle(v string) *CreateCertificateApplicationOrderRequestTech {
	s.Title = &v
	return s
}

type CreateCertificateApplicationOrderRequestDnsProviderInfos struct {
	// {"en":"domain","zh_CN":"域名"}
	Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
	// {"en":"Support CloudDNS, The optional certificate brand is LE or TrustAsia or GlobalSign or ZeroSSL;\nSupport CloudFlare, The optional certificate brand is LE or ZeroSSL;\nSupport DNSPod.cn, The optional certificate brand is LE or ZeroSSL;\nSupport DNSPod.com, The optional certificate brand is LE or ZeroSSL;\nSupport TencentDNSPod, The optional certificate brand is LE or ZeroSSL;\nSupport Aliyun, The optional certificate brand is LE or ZeroSSL;\nSupport GoDaddy, The optional certificate brand is LE or ZeroSSL;\nSupport AmazonRoute53, The optional certificate brand is LE or ZeroSSL;\nSupport GoogleDomains, The optional certificate brand is LE or ZeroSSL;\nSupport AkamaiEdgeDNS, The optional certificate brand is LE or ZeroSSL.","zh_CN":"DNS托管商编码，\n支持CloudDNS，可选证书品牌为LE、TrustAsia、GlobalSign、ZeroSSL；\n支持CloudFlare，可选证书品牌为LE、ZeroSSL；\n支持DNSPod.cn，可选证书品牌为LE、ZeroSSL；\n支持DNSPod.com，可选证书品牌为LE、ZeroSSL；\n支持TencentDNSPod，可选证书品牌为LE、ZeroSSL；\n支持Aliyun，可选证书品牌为LE、ZeroSSL；\n支持GoDaddy，可选证书品牌为LE、ZeroSSL；\n支持AmazonRoute53，可选证书品牌为LE、ZeroSSL；\n支持GoogleDomains，可选证书品牌为LE、ZeroSSL；\n支持AkamaiEdgeDNS，可选证书品牌为LE、ZeroSSL。","exampleValue":"CloudDNS,CloudFlare"}
	DnsProviderCode *string `json:"dnsProviderCode,omitempty" xml:"dnsProviderCode,omitempty"`
	// {"en":"DNS Api Access, JSON format:\nThe hosting provider is CloudDNS, JSON KEY is accessKey and secretKey;\nThe hosting provider is CloudFlare, The global API key JSON KEY is CF_Key CF_Email, The restricted API token JSON KEY is CF_Token, CF_ZoneID or CF_Sccount-ID.","zh_CN":"DNS托管商API凭证，JSON格式，\n托管商为CloudDNS，JSON KEY为accessKey、secretKey；\n托管商为CloudFlare，全局API密钥JSON KEY为CF_Key、CF_Email，限制性API令牌JSON KEY为CF_Token、CF_Zone_ID或CF_Account_ID；\n托管商为DNSPod.cn，JSON KEY为DP_Id、DP_Key；\n托管商为DNSPod.com，JSON KEY为DPI_Id、DPI_Key；\n托管商为TencentDNSPod，JSON KEY为Tencent_SecretId、Tencent_SecretKey；\n托管商为Aliyun，JSON KEY为Ali_Key、Ali_Secret；\n托管商为GoDaddy，JSON KEY为GD_Key、GD_Secret；\n托管商为AmazonRoute53，JSON KEY为AWS_ACCESS_KEY_ID、AWS_SECRET_ACCESS_KEY；\n托管商为GoogleDomains，JSON KEY为GOOGLEDOMAINS_ACCESS_TOKEN、GOOGLEDOMAINS_ZONE；\n托管商为AkamaiEdgeDNS，JSON KEY为AKAMAI_CLIENT_TOKEN、AKAMAI_ACCESS_TOKEN、AKAMAI_CLIENT_SECRET、AKAMAI_HOST；"}
	DnsApiAccess *string `json:"dnsApiAccess,omitempty" xml:"dnsApiAccess,omitempty"`
	// {"defaultValue":"false","en":"Whether to use alias verification","zh_CN":"是否使用别名方式","exampleValue":"true,false"}
	EnableDnsAliasMode *string `json:"enableDnsAliasMode,omitempty" xml:"enableDnsAliasMode,omitempty"`
	// {"en":"Validate alias","zh_CN":"验证别名"}
	ValidateAliasDomain *string `json:"validateAliasDomain,omitempty" xml:"validateAliasDomain,omitempty"`
}

func (s CreateCertificateApplicationOrderRequestDnsProviderInfos) String() string {
	return tea.Prettify(s)
}

func (s CreateCertificateApplicationOrderRequestDnsProviderInfos) GoString() string {
	return s.String()
}

func (s *CreateCertificateApplicationOrderRequestDnsProviderInfos) SetDomain(v string) *CreateCertificateApplicationOrderRequestDnsProviderInfos {
	s.Domain = &v
	return s
}

func (s *CreateCertificateApplicationOrderRequestDnsProviderInfos) SetDnsProviderCode(v string) *CreateCertificateApplicationOrderRequestDnsProviderInfos {
	s.DnsProviderCode = &v
	return s
}

func (s *CreateCertificateApplicationOrderRequestDnsProviderInfos) SetDnsApiAccess(v string) *CreateCertificateApplicationOrderRequestDnsProviderInfos {
	s.DnsApiAccess = &v
	return s
}

func (s *CreateCertificateApplicationOrderRequestDnsProviderInfos) SetEnableDnsAliasMode(v string) *CreateCertificateApplicationOrderRequestDnsProviderInfos {
	s.EnableDnsAliasMode = &v
	return s
}

func (s *CreateCertificateApplicationOrderRequestDnsProviderInfos) SetValidateAliasDomain(v string) *CreateCertificateApplicationOrderRequestDnsProviderInfos {
	s.ValidateAliasDomain = &v
	return s
}

type CreateCertificateApplicationOrderRequestHeader struct {
}

func (s CreateCertificateApplicationOrderRequestHeader) String() string {
	return tea.Prettify(s)
}

func (s CreateCertificateApplicationOrderRequestHeader) GoString() string {
	return s.String()
}

type CreateCertificateApplicationOrderPaths struct {
}

func (s CreateCertificateApplicationOrderPaths) String() string {
	return tea.Prettify(s)
}

func (s CreateCertificateApplicationOrderPaths) GoString() string {
	return s.String()
}

type CreateCertificateApplicationOrderParameters struct {
}

func (s CreateCertificateApplicationOrderParameters) String() string {
	return tea.Prettify(s)
}

func (s CreateCertificateApplicationOrderParameters) GoString() string {
	return s.String()
}

type CreateCertificateApplicationOrderResponse struct {
	// {"en":"Result Code, success is 0","zh_CN":"响应码，成功为0"}
	Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
	// {"en":"Result Message","zh_CN":"响应信息"}
	Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
	// {"en":"Result Data","zh_CN":"响应数据"}
	Data *CreateCertificateApplicationOrderResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
}

func (s CreateCertificateApplicationOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateCertificateApplicationOrderResponse) GoString() string {
	return s.String()
}

func (s *CreateCertificateApplicationOrderResponse) SetCode(v string) *CreateCertificateApplicationOrderResponse {
	s.Code = &v
	return s
}

func (s *CreateCertificateApplicationOrderResponse) SetMessage(v string) *CreateCertificateApplicationOrderResponse {
	s.Message = &v
	return s
}

func (s *CreateCertificateApplicationOrderResponse) SetData(v *CreateCertificateApplicationOrderResponseData) *CreateCertificateApplicationOrderResponse {
	s.Data = v
	return s
}

type CreateCertificateApplicationOrderResponseData struct {
	// {"en":"Order ID","zh_CN":"订单ID"}
	OrderId *string `json:"orderId,omitempty" xml:"orderId,omitempty" require:"true"`
	// {"en":"Order ID List","zh_CN":"订单ID列表"}
	OrderIds []*string `json:"orderIds,omitempty" xml:"orderIds,omitempty" require:"true" type:"Repeated"`
	// {"en":"Purchase Record ID List","zh_CN":"采购记录ID列表"}
	PurchaseRecordIds []*string `json:"purchaseRecordIds,omitempty" xml:"purchaseRecordIds,omitempty" require:"true" type:"Repeated"`
	// {"en":"Validate alias List","zh_CN":"验证别名列表"}
	ValidateAliasDomains []*CreateCertificateApplicationOrderResponseDataValidateAliasDomains `json:"validateAliasDomains,omitempty" xml:"validateAliasDomains,omitempty" require:"true" type:"Repeated"`
	// {"en":"Domain and order relation info list","zh_CN":"域名与订单关联关系"}
	DomainOrderRelationInfos []*CreateCertificateApplicationOrderResponseDataDomainOrderRelationInfos `json:"domainOrderRelationInfos,omitempty" xml:"domainOrderRelationInfos,omitempty" require:"true" type:"Repeated"`
}

func (s CreateCertificateApplicationOrderResponseData) String() string {
	return tea.Prettify(s)
}

func (s CreateCertificateApplicationOrderResponseData) GoString() string {
	return s.String()
}

func (s *CreateCertificateApplicationOrderResponseData) SetOrderId(v string) *CreateCertificateApplicationOrderResponseData {
	s.OrderId = &v
	return s
}

func (s *CreateCertificateApplicationOrderResponseData) SetOrderIds(v []*string) *CreateCertificateApplicationOrderResponseData {
	s.OrderIds = v
	return s
}

func (s *CreateCertificateApplicationOrderResponseData) SetPurchaseRecordIds(v []*string) *CreateCertificateApplicationOrderResponseData {
	s.PurchaseRecordIds = v
	return s
}

func (s *CreateCertificateApplicationOrderResponseData) SetValidateAliasDomains(v []*CreateCertificateApplicationOrderResponseDataValidateAliasDomains) *CreateCertificateApplicationOrderResponseData {
	s.ValidateAliasDomains = v
	return s
}

func (s *CreateCertificateApplicationOrderResponseData) SetDomainOrderRelationInfos(v []*CreateCertificateApplicationOrderResponseDataDomainOrderRelationInfos) *CreateCertificateApplicationOrderResponseData {
	s.DomainOrderRelationInfos = v
	return s
}

type CreateCertificateApplicationOrderResponseDataValidateAliasDomains struct {
	// {"en":"domain","zh_CN":"域名"}
	Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
	// {"en":"Validate alias","zh_CN":"验证别名"}
	ValidateAliasDomain *string `json:"validateAliasDomain,omitempty" xml:"validateAliasDomain,omitempty" require:"true"`
}

func (s CreateCertificateApplicationOrderResponseDataValidateAliasDomains) String() string {
	return tea.Prettify(s)
}

func (s CreateCertificateApplicationOrderResponseDataValidateAliasDomains) GoString() string {
	return s.String()
}

func (s *CreateCertificateApplicationOrderResponseDataValidateAliasDomains) SetDomain(v string) *CreateCertificateApplicationOrderResponseDataValidateAliasDomains {
	s.Domain = &v
	return s
}

func (s *CreateCertificateApplicationOrderResponseDataValidateAliasDomains) SetValidateAliasDomain(v string) *CreateCertificateApplicationOrderResponseDataValidateAliasDomains {
	s.ValidateAliasDomain = &v
	return s
}

type CreateCertificateApplicationOrderResponseDataDomainOrderRelationInfos struct {
	// {"en":"Domain name","zh_CN":"域名名称"}
	DomainName *string `json:"domainName,omitempty" xml:"domainName,omitempty" require:"true"`
	// {"en":"Order ID","zh_CN":"订单ID"}
	SalesOrderId *string `json:"salesOrderId,omitempty" xml:"salesOrderId,omitempty" require:"true"`
	// {"en":"Purchase Record ID","zh_CN":"采购记录ID"}
	PurchaseRecordId *string `json:"purchaseRecordId,omitempty" xml:"purchaseRecordId,omitempty" require:"true"`
}

func (s CreateCertificateApplicationOrderResponseDataDomainOrderRelationInfos) String() string {
	return tea.Prettify(s)
}

func (s CreateCertificateApplicationOrderResponseDataDomainOrderRelationInfos) GoString() string {
	return s.String()
}

func (s *CreateCertificateApplicationOrderResponseDataDomainOrderRelationInfos) SetDomainName(v string) *CreateCertificateApplicationOrderResponseDataDomainOrderRelationInfos {
	s.DomainName = &v
	return s
}

func (s *CreateCertificateApplicationOrderResponseDataDomainOrderRelationInfos) SetSalesOrderId(v string) *CreateCertificateApplicationOrderResponseDataDomainOrderRelationInfos {
	s.SalesOrderId = &v
	return s
}

func (s *CreateCertificateApplicationOrderResponseDataDomainOrderRelationInfos) SetPurchaseRecordId(v string) *CreateCertificateApplicationOrderResponseDataDomainOrderRelationInfos {
	s.PurchaseRecordId = &v
	return s
}

type CreateCertificateApplicationOrderResponseHeader struct {
}

func (s CreateCertificateApplicationOrderResponseHeader) String() string {
	return tea.Prettify(s)
}

func (s CreateCertificateApplicationOrderResponseHeader) GoString() string {
	return s.String()
}

type ListCertificateApplicationOrdersRequest struct {
	// {"en":"Order ID","zh_CN":"订单ID"}
	OrderId *string `json:"orderId,omitempty" xml:"orderId,omitempty"`
	// {"en":"Status of certificate apply order.","zh_CN":"订单状态:\nACCEPT_SUCCESS：订单提交成功，待后台发起证书申请；\nAPPLYING：证书申请中，待域名验证；\nISSUE_SUCCESS：证书签发成功，可进行证书配置与部署；\nISSUE_FAILURE：证书签发失败；\nREVOKED：证书已吊销；\nCANCELED：已取消证书申请；","exampleValue":"ACCEPT_SUCCESS,APPLYING,ISSUE_SUCCESS,ISSUE_FAILURE,REVOKED,CANCELED"}
	OrderStatus []*string `json:"orderStatus,omitempty" xml:"orderStatus,omitempty" type:"Repeated"`
	// {"en":"Certificate Name","zh_CN":"证书名称"}
	CertificateName *string `json:"certificateName,omitempty" xml:"certificateName,omitempty"`
	// {"en":"Domain","zh_CN":"域名"}
	Domain *string `json:"domain,omitempty" xml:"domain,omitempty"`
	// {"en":"The time when the certificate order was created. The format is yyyy-MM-dd HH:mm:ss","zh_CN":"订单创建时间，格式为：yyyy-MM-dd HH:mm:ss"}
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// {"en":"The time when the certificate order was finished. The format is yyyy-MM-dd HH:mm:ss","zh_CN":"订单结束时间，格式为：yyyy-MM-dd HH:mm:ss。证书签发成功或证书签发失败，则记录结束时间。"}
	EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// {"en":"Page Parammeter.","zh_CN":"分页参数"}
	PageParam *ListCertificateApplicationOrdersRequestPageParam `json:"pageParam,omitempty" xml:"pageParam,omitempty" type:"Struct"`
}

func (s ListCertificateApplicationOrdersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListCertificateApplicationOrdersRequest) GoString() string {
	return s.String()
}

func (s *ListCertificateApplicationOrdersRequest) SetOrderId(v string) *ListCertificateApplicationOrdersRequest {
	s.OrderId = &v
	return s
}

func (s *ListCertificateApplicationOrdersRequest) SetOrderStatus(v []*string) *ListCertificateApplicationOrdersRequest {
	s.OrderStatus = v
	return s
}

func (s *ListCertificateApplicationOrdersRequest) SetCertificateName(v string) *ListCertificateApplicationOrdersRequest {
	s.CertificateName = &v
	return s
}

func (s *ListCertificateApplicationOrdersRequest) SetDomain(v string) *ListCertificateApplicationOrdersRequest {
	s.Domain = &v
	return s
}

func (s *ListCertificateApplicationOrdersRequest) SetStartTime(v string) *ListCertificateApplicationOrdersRequest {
	s.StartTime = &v
	return s
}

func (s *ListCertificateApplicationOrdersRequest) SetEndTime(v string) *ListCertificateApplicationOrdersRequest {
	s.EndTime = &v
	return s
}

func (s *ListCertificateApplicationOrdersRequest) SetPageParam(v *ListCertificateApplicationOrdersRequestPageParam) *ListCertificateApplicationOrdersRequest {
	s.PageParam = v
	return s
}

type ListCertificateApplicationOrdersRequestPageParam struct {
	// {"en":"Page Size,Range:1-100","zh_CN":"每页大小，1-100","exampleValue":"00"}
	PageSize *int `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// {"defaultValue":"0","en":"Page Number, the first page must be 0.Range:>=0","zh_CN":"页码，当前页，第一页从0开始，必须大于等于0"}
	PageNumber *int `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
}

func (s ListCertificateApplicationOrdersRequestPageParam) String() string {
	return tea.Prettify(s)
}

func (s ListCertificateApplicationOrdersRequestPageParam) GoString() string {
	return s.String()
}

func (s *ListCertificateApplicationOrdersRequestPageParam) SetPageSize(v int) *ListCertificateApplicationOrdersRequestPageParam {
	s.PageSize = &v
	return s
}

func (s *ListCertificateApplicationOrdersRequestPageParam) SetPageNumber(v int) *ListCertificateApplicationOrdersRequestPageParam {
	s.PageNumber = &v
	return s
}

type ListCertificateApplicationOrdersRequestHeader struct {
}

func (s ListCertificateApplicationOrdersRequestHeader) String() string {
	return tea.Prettify(s)
}

func (s ListCertificateApplicationOrdersRequestHeader) GoString() string {
	return s.String()
}

type ListCertificateApplicationOrdersPaths struct {
}

func (s ListCertificateApplicationOrdersPaths) String() string {
	return tea.Prettify(s)
}

func (s ListCertificateApplicationOrdersPaths) GoString() string {
	return s.String()
}

type ListCertificateApplicationOrdersParameters struct {
}

func (s ListCertificateApplicationOrdersParameters) String() string {
	return tea.Prettify(s)
}

func (s ListCertificateApplicationOrdersParameters) GoString() string {
	return s.String()
}

type ListCertificateApplicationOrdersResponse struct {
	// {"en":"Result Code, success is 0","zh_CN":"响应码，成功为0"}
	Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
	// {"en":"Result Message","zh_CN":"响应信息"}
	Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
	// {"en":"Result Data","zh_CN":"响应数据"}
	Data *ListCertificateApplicationOrdersResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
}

func (s ListCertificateApplicationOrdersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListCertificateApplicationOrdersResponse) GoString() string {
	return s.String()
}

func (s *ListCertificateApplicationOrdersResponse) SetCode(v string) *ListCertificateApplicationOrdersResponse {
	s.Code = &v
	return s
}

func (s *ListCertificateApplicationOrdersResponse) SetMessage(v string) *ListCertificateApplicationOrdersResponse {
	s.Message = &v
	return s
}

func (s *ListCertificateApplicationOrdersResponse) SetData(v *ListCertificateApplicationOrdersResponseData) *ListCertificateApplicationOrdersResponse {
	s.Data = v
	return s
}

type ListCertificateApplicationOrdersResponseData struct {
	// {"en":"Order list","zh_CN":"订单列表"}
	Orders []*ListCertificateApplicationOrdersResponseDataOrders `json:"orders,omitempty" xml:"orders,omitempty" require:"true" type:"Repeated"`
	// {"en":"Page Info","zh_CN":"分页数据"}
	PageInfo *ListCertificateApplicationOrdersResponseDataPageInfo `json:"pageInfo,omitempty" xml:"pageInfo,omitempty" require:"true" type:"Struct"`
}

func (s ListCertificateApplicationOrdersResponseData) String() string {
	return tea.Prettify(s)
}

func (s ListCertificateApplicationOrdersResponseData) GoString() string {
	return s.String()
}

func (s *ListCertificateApplicationOrdersResponseData) SetOrders(v []*ListCertificateApplicationOrdersResponseDataOrders) *ListCertificateApplicationOrdersResponseData {
	s.Orders = v
	return s
}

func (s *ListCertificateApplicationOrdersResponseData) SetPageInfo(v *ListCertificateApplicationOrdersResponseDataPageInfo) *ListCertificateApplicationOrdersResponseData {
	s.PageInfo = v
	return s
}

type ListCertificateApplicationOrdersResponseDataOrders struct {
	// {"en":"Order ID","zh_CN":"订单ID"}
	OrderId *string `json:"orderId,omitempty" xml:"orderId,omitempty" require:"true"`
	// {"en":"Status of certificate apply order.","zh_CN":"证书订单状态","exampleValue":"ACCEPT_SUCCESS,APPLYING,ISSUE_SUCCESS,ISSUE_FAILURE,REVOKED,CANCELED"}
	OrderStatus *string `json:"orderStatus,omitempty" xml:"orderStatus,omitempty" require:"true"`
	// {"en":"Certificate ID","zh_CN":"证书ID"}
	CertificateId *int `json:"certificateId,omitempty" xml:"certificateId,omitempty" require:"true"`
	// {"en":"Certificate Name","zh_CN":"证书名称"}
	CertificateName *string `json:"certificateName,omitempty" xml:"certificateName,omitempty" require:"true"`
	// {"en":"Certificate Brand","zh_CN":"证书品牌"}
	CertificateBrand *string `json:"certificateBrand,omitempty" xml:"certificateBrand,omitempty" require:"true"`
	// {"en":"Certificate Type","zh_CN":"证书类型"}
	CertificateType *string `json:"certificateType,omitempty" xml:"certificateType,omitempty" require:"true"`
	// {"en":"Auto Renew","zh_CN":"是否自动续签"}
	AutoRenew *string `json:"autoRenew,omitempty" xml:"autoRenew,omitempty" require:"true"`
	// {"en":"Description","zh_CN":"描述"}
	Description *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	// {"en":"Common Name of the certificate","zh_CN":"证书的通用名称"}
	CommonName *string `json:"commonName,omitempty" xml:"commonName,omitempty" require:"true"`
	// {"en":"SAN","zh_CN":"主体备用名称"}
	SubjectAlternativeNames []*string `json:"subjectAlternativeNames,omitempty" xml:"subjectAlternativeNames,omitempty" require:"true" type:"Repeated"`
	// {"en":"Create Time","zh_CN":"创建时间"}
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty" require:"true"`
}

func (s ListCertificateApplicationOrdersResponseDataOrders) String() string {
	return tea.Prettify(s)
}

func (s ListCertificateApplicationOrdersResponseDataOrders) GoString() string {
	return s.String()
}

func (s *ListCertificateApplicationOrdersResponseDataOrders) SetOrderId(v string) *ListCertificateApplicationOrdersResponseDataOrders {
	s.OrderId = &v
	return s
}

func (s *ListCertificateApplicationOrdersResponseDataOrders) SetOrderStatus(v string) *ListCertificateApplicationOrdersResponseDataOrders {
	s.OrderStatus = &v
	return s
}

func (s *ListCertificateApplicationOrdersResponseDataOrders) SetCertificateId(v int) *ListCertificateApplicationOrdersResponseDataOrders {
	s.CertificateId = &v
	return s
}

func (s *ListCertificateApplicationOrdersResponseDataOrders) SetCertificateName(v string) *ListCertificateApplicationOrdersResponseDataOrders {
	s.CertificateName = &v
	return s
}

func (s *ListCertificateApplicationOrdersResponseDataOrders) SetCertificateBrand(v string) *ListCertificateApplicationOrdersResponseDataOrders {
	s.CertificateBrand = &v
	return s
}

func (s *ListCertificateApplicationOrdersResponseDataOrders) SetCertificateType(v string) *ListCertificateApplicationOrdersResponseDataOrders {
	s.CertificateType = &v
	return s
}

func (s *ListCertificateApplicationOrdersResponseDataOrders) SetAutoRenew(v string) *ListCertificateApplicationOrdersResponseDataOrders {
	s.AutoRenew = &v
	return s
}

func (s *ListCertificateApplicationOrdersResponseDataOrders) SetDescription(v string) *ListCertificateApplicationOrdersResponseDataOrders {
	s.Description = &v
	return s
}

func (s *ListCertificateApplicationOrdersResponseDataOrders) SetCommonName(v string) *ListCertificateApplicationOrdersResponseDataOrders {
	s.CommonName = &v
	return s
}

func (s *ListCertificateApplicationOrdersResponseDataOrders) SetSubjectAlternativeNames(v []*string) *ListCertificateApplicationOrdersResponseDataOrders {
	s.SubjectAlternativeNames = v
	return s
}

func (s *ListCertificateApplicationOrdersResponseDataOrders) SetCreateTime(v string) *ListCertificateApplicationOrdersResponseDataOrders {
	s.CreateTime = &v
	return s
}

type ListCertificateApplicationOrdersResponseDataPageInfo struct {
	// {"en":"Total Number","zh_CN":"总数"}
	TotalNumber *int `json:"totalNumber,omitempty" xml:"totalNumber,omitempty" require:"true"`
	// {"en":"Page Number","zh_CN":"页码"}
	PageNumber *int `json:"pageNumber,omitempty" xml:"pageNumber,omitempty" require:"true"`
	// {"en":"Page Size","zh_CN":"每页大小"}
	PageSize *int `json:"pageSize,omitempty" xml:"pageSize,omitempty" require:"true"`
	// {"en":"Total Page Number","zh_CN":"总页码数"}
	TotalPageNumber *int `json:"totalPageNumber,omitempty" xml:"totalPageNumber,omitempty" require:"true"`
}

func (s ListCertificateApplicationOrdersResponseDataPageInfo) String() string {
	return tea.Prettify(s)
}

func (s ListCertificateApplicationOrdersResponseDataPageInfo) GoString() string {
	return s.String()
}

func (s *ListCertificateApplicationOrdersResponseDataPageInfo) SetTotalNumber(v int) *ListCertificateApplicationOrdersResponseDataPageInfo {
	s.TotalNumber = &v
	return s
}

func (s *ListCertificateApplicationOrdersResponseDataPageInfo) SetPageNumber(v int) *ListCertificateApplicationOrdersResponseDataPageInfo {
	s.PageNumber = &v
	return s
}

func (s *ListCertificateApplicationOrdersResponseDataPageInfo) SetPageSize(v int) *ListCertificateApplicationOrdersResponseDataPageInfo {
	s.PageSize = &v
	return s
}

func (s *ListCertificateApplicationOrdersResponseDataPageInfo) SetTotalPageNumber(v int) *ListCertificateApplicationOrdersResponseDataPageInfo {
	s.TotalPageNumber = &v
	return s
}

type ListCertificateApplicationOrdersResponseHeader struct {
}

func (s ListCertificateApplicationOrdersResponseHeader) String() string {
	return tea.Prettify(s)
}

func (s ListCertificateApplicationOrdersResponseHeader) GoString() string {
	return s.String()
}

type CancelCertificateApplicationOrderRequest struct {
	// {"en":"The orderId","zh_CN":"订单ID"}
	OrderId *string `json:"orderId,omitempty" xml:"orderId,omitempty"`
	// {"en":"The purchaseRecordId","zh_CN":"采购记录ID"}
	PurchaseRecordId *string `json:"purchaseRecordId,omitempty" xml:"purchaseRecordId,omitempty"`
}

func (s CancelCertificateApplicationOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s CancelCertificateApplicationOrderRequest) GoString() string {
	return s.String()
}

func (s *CancelCertificateApplicationOrderRequest) SetOrderId(v string) *CancelCertificateApplicationOrderRequest {
	s.OrderId = &v
	return s
}

func (s *CancelCertificateApplicationOrderRequest) SetPurchaseRecordId(v string) *CancelCertificateApplicationOrderRequest {
	s.PurchaseRecordId = &v
	return s
}

type CancelCertificateApplicationOrderRequestHeader struct {
}

func (s CancelCertificateApplicationOrderRequestHeader) String() string {
	return tea.Prettify(s)
}

func (s CancelCertificateApplicationOrderRequestHeader) GoString() string {
	return s.String()
}

type CancelCertificateApplicationOrderPaths struct {
}

func (s CancelCertificateApplicationOrderPaths) String() string {
	return tea.Prettify(s)
}

func (s CancelCertificateApplicationOrderPaths) GoString() string {
	return s.String()
}

type CancelCertificateApplicationOrderParameters struct {
}

func (s CancelCertificateApplicationOrderParameters) String() string {
	return tea.Prettify(s)
}

func (s CancelCertificateApplicationOrderParameters) GoString() string {
	return s.String()
}

type CancelCertificateApplicationOrderResponse struct {
	// {"en":"Result Code, 0 means successful","zh_CN":"响应码，成功为0"}
	Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
	// {"en":"Result Message","zh_CN":"响应信息"}
	Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s CancelCertificateApplicationOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s CancelCertificateApplicationOrderResponse) GoString() string {
	return s.String()
}

func (s *CancelCertificateApplicationOrderResponse) SetCode(v string) *CancelCertificateApplicationOrderResponse {
	s.Code = &v
	return s
}

func (s *CancelCertificateApplicationOrderResponse) SetMessage(v string) *CancelCertificateApplicationOrderResponse {
	s.Message = &v
	return s
}

type CancelCertificateApplicationOrderResponseHeader struct {
}

func (s CancelCertificateApplicationOrderResponseHeader) String() string {
	return tea.Prettify(s)
}

func (s CancelCertificateApplicationOrderResponseHeader) GoString() string {
	return s.String()
}

type GetCertificateApplicationOrderForTerraformRequest struct {
	// {"en":"The order ID.","zh_CN":"订单id"}
	OrderId *string `json:"orderId,omitempty" xml:"orderId,omitempty"`
	// {"en":"The purchase record ID.","zh_CN":"采购记录ID"}
	PurchaseRecordId *string `json:"purchaseRecordId,omitempty" xml:"purchaseRecordId,omitempty"`
}

func (s GetCertificateApplicationOrderForTerraformRequest) String() string {
	return tea.Prettify(s)
}

func (s GetCertificateApplicationOrderForTerraformRequest) GoString() string {
	return s.String()
}

func (s *GetCertificateApplicationOrderForTerraformRequest) SetOrderId(v string) *GetCertificateApplicationOrderForTerraformRequest {
	s.OrderId = &v
	return s
}

func (s *GetCertificateApplicationOrderForTerraformRequest) SetPurchaseRecordId(v string) *GetCertificateApplicationOrderForTerraformRequest {
	s.PurchaseRecordId = &v
	return s
}

type GetCertificateApplicationOrderForTerraformRequestHeader struct {
}

func (s GetCertificateApplicationOrderForTerraformRequestHeader) String() string {
	return tea.Prettify(s)
}

func (s GetCertificateApplicationOrderForTerraformRequestHeader) GoString() string {
	return s.String()
}

type GetCertificateApplicationOrderForTerraformPaths struct {
}

func (s GetCertificateApplicationOrderForTerraformPaths) String() string {
	return tea.Prettify(s)
}

func (s GetCertificateApplicationOrderForTerraformPaths) GoString() string {
	return s.String()
}

type GetCertificateApplicationOrderForTerraformParameters struct {
}

func (s GetCertificateApplicationOrderForTerraformParameters) String() string {
	return tea.Prettify(s)
}

func (s GetCertificateApplicationOrderForTerraformParameters) GoString() string {
	return s.String()
}

type GetCertificateApplicationOrderForTerraformResponse struct {
	// {"en":"The error code. 0 indicates success.","zh_CN":"错误码，成功为0"}
	Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
	// {"en":"error message","zh_CN":"错误信息"}
	Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
	// {"en":"The order details.","zh_CN":"订单详情"}
	Data *GetCertificateApplicationOrderForTerraformResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
}

func (s GetCertificateApplicationOrderForTerraformResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCertificateApplicationOrderForTerraformResponse) GoString() string {
	return s.String()
}

func (s *GetCertificateApplicationOrderForTerraformResponse) SetCode(v string) *GetCertificateApplicationOrderForTerraformResponse {
	s.Code = &v
	return s
}

func (s *GetCertificateApplicationOrderForTerraformResponse) SetMessage(v string) *GetCertificateApplicationOrderForTerraformResponse {
	s.Message = &v
	return s
}

func (s *GetCertificateApplicationOrderForTerraformResponse) SetData(v *GetCertificateApplicationOrderForTerraformResponseData) *GetCertificateApplicationOrderForTerraformResponse {
	s.Data = v
	return s
}

type GetCertificateApplicationOrderForTerraformResponseData struct {
	// {"en":"orderId","zh_CN":"订单id"}
	OrderId *string `json:"orderId,omitempty" xml:"orderId,omitempty" require:"true"`
	// {"en":"orderStatus\nACCEPT_SUCCESS: Received successfully\nAPPLYING: Under application\nAPPLy_FAILURE: Application preparation failed\nVALIDATETWAIT: To be verified\nVALIDATE_PROCESSING: Verification processing\nVALIDATE_SUCCESS: Verification successful\nVALIDATE_SAILURE: Validation failed\nISSUEAwaIT: To be issued\nISSUE_SUCCESS: Issued successfully\nISSUE_SFAILURE: Issuance failed\nCANCELED: Cancel application\nREVOKED: Revocation\nDEPLO_SUCCESS: Deployment successful\nDEPLOy_Failures: Deployment failed","zh_CN":"订单状态\nACCEPT_SUCCESS: 接收成功\nAPPLYING: 申请中\nAPPLY_FAILURE：申请准备失败\nVALIDATE_WAIT: 待验证\nVALIDATE_PROCESSING: 验证中\nVALIDATE_SUCCESS: 验证成功\nVALIDATE_FAILURE: 验证失败\nISSUE_WAIT: 待签发\nISSUE_SUCCESS: 签发成功\nISSUE_FAILURE: 签发失败\nCANCELED: 取消申请\nREVOKED: 吊销\nDEPLOY_SUCCESS：部署成功\nDEPLOY_FAILURE：部署失败"}
	OrderStatus *string `json:"orderStatus,omitempty" xml:"orderStatus,omitempty" require:"true"`
	// {"en":"certificateId","zh_CN":"证书id"}
	CertificateId *int64 `json:"certificateId,omitempty" xml:"certificateId,omitempty" require:"true"`
	// {"en":"certificateName","zh_CN":"证书名"}
	CertificateName *string `json:"certificateName,omitempty" xml:"certificateName,omitempty" require:"true"`
	// {"en":"certificateBrand","zh_CN":"证书品牌"}
	CertificateBrand *string `json:"certificateBrand,omitempty" xml:"certificateBrand,omitempty" require:"true"`
	// {"en":"certificateType","zh_CN":"证书类型"}
	CertificateType *string `json:"certificateType,omitempty" xml:"certificateType,omitempty" require:"true"`
	// {"en":"description","zh_CN":"描述"}
	Description *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	// {"en":"algorithm","zh_CN":"算法"}
	Algorithm *string `json:"algorithm,omitempty" xml:"algorithm,omitempty" require:"true"`
	// {"en":"autoValidate","zh_CN":"是否自动验证"}
	AutoValidate *string `json:"autoValidate,omitempty" xml:"autoValidate,omitempty" require:"true"`
	// {"en":"validateMethod","zh_CN":"验证方式"}
	ValidateMethod *string `json:"validateMethod,omitempty" xml:"validateMethod,omitempty" require:"true"`
	// {"en":"autoRenew","zh_CN":"是否自动续订"}
	AutoRenew *string `json:"autoRenew,omitempty" xml:"autoRenew,omitempty" require:"true"`
	// {"en":"validityDays","zh_CN":"有效天数"}
	ValidityDays *int `json:"validityDays,omitempty" xml:"validityDays,omitempty" require:"true"`
	// {"en":"country","zh_CN":"国家代码"}
	Country *string `json:"country,omitempty" xml:"country,omitempty" require:"true"`
	// {"en":"state","zh_CN":"省"}
	State *string `json:"state,omitempty" xml:"state,omitempty" require:"true"`
	// {"en":"city","zh_CN":"市"}
	City *string `json:"city,omitempty" xml:"city,omitempty" require:"true"`
	// {"en":"street","zh_CN":"街道"}
	Street *string `json:"street,omitempty" xml:"street,omitempty" require:"true"`
	// {"en":"company","zh_CN":"公司"}
	Company *string `json:"company,omitempty" xml:"company,omitempty" require:"true"`
	// {"en":"department","zh_CN":"部门"}
	Department *string `json:"department,omitempty" xml:"department,omitempty" require:"true"`
	// {"en":"commonName","zh_CN":"通用名"}
	CommonName *string `json:"commonName,omitempty" xml:"commonName,omitempty" require:"true"`
	// {"en":"email","zh_CN":"邮箱"}
	Email *string `json:"email,omitempty" xml:"email,omitempty" require:"true"`
	// {"en":"subjectAlternativeNames","zh_CN":"主体备用名称"}
	SubjectAlternativeNames []*string `json:"subjectAlternativeNames,omitempty" xml:"subjectAlternativeNames,omitempty" require:"true" type:"Repeated"`
	// {"en":"createTime","zh_CN":"创建时间"}
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty" require:"true"`
	// {"en":"errorMessage","zh_CN":"错误信息"}
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty" require:"true"`
	// {"en":"certificateSpec","zh_CN":"证书规格"}
	CertificateSpec *string `json:"certificateSpec,omitempty" xml:"certificateSpec,omitempty" require:"true"`
	// {"en":"domainType","zh_CN":"域名类型"}
	DomainType *string `json:"domainType,omitempty" xml:"domainType,omitempty" require:"true"`
	// {"en":"autoDeploy","zh_CN":"是否自动部署"}
	AutoDeploy *string `json:"autoDeploy,omitempty" xml:"autoDeploy,omitempty" require:"true"`
	// {"en":"remainValidityDays","zh_CN":"剩余有效期"}
	RemainValidityDays *int `json:"remainValidityDays,omitempty" xml:"remainValidityDays,omitempty" require:"true"`
	// {"en":"dnsProviderInfos","zh_CN":"DNS托管商信息"}
	DnsProviderInfos []*GetCertificateApplicationOrderForTerraformResponseDataDnsProviderInfos `json:"dnsProviderInfos,omitempty" xml:"dnsProviderInfos,omitempty" require:"true" type:"Repeated"`
	// {"en":"Primary Certificate Brand","zh_CN":"主证书品牌"}
	PrimaryCertificateBrand *string `json:"primaryCertificateBrand,omitempty" xml:"primaryCertificateBrand,omitempty" require:"true"`
	// {"en":"Backup Certificate Brand","zh_CN":"备用CA证书品牌"}
	BackupCertificateBrand *string `json:"backupCertificateBrand,omitempty" xml:"backupCertificateBrand,omitempty" require:"true"`
	// {"en":"Org Validate Method","zh_CN":"组织验证方式"}
	OrgValidateMethod *string `json:"orgValidateMethod,omitempty" xml:"orgValidateMethod,omitempty" require:"true"`
	// {"en":"Phone","zh_CN":"组织联系电话"}
	Phone *string `json:"phone,omitempty" xml:"phone,omitempty" require:"true"`
	// {"en":"Street","zh_CN":"街道2"}
	Street1 *string `json:"street1,omitempty" xml:"street1,omitempty" require:"true"`
	// {"en":"Postal Code","zh_CN":"邮编"}
	PostalCode *string `json:"postalCode,omitempty" xml:"postalCode,omitempty" require:"true"`
	// {"en":"admin","zh_CN":"订单管理联系人"}
	Admin *GetCertificateApplicationOrderForTerraformResponseDataAdmin `json:"admin,omitempty" xml:"admin,omitempty" require:"true" type:"Struct"`
	// {"en":"Technical contact person for the order.","zh_CN":"订单技术联系人"}
	Tech *GetCertificateApplicationOrderForTerraformResponseDataTech `json:"tech,omitempty" xml:"tech,omitempty" require:"true" type:"Struct"`
}

func (s GetCertificateApplicationOrderForTerraformResponseData) String() string {
	return tea.Prettify(s)
}

func (s GetCertificateApplicationOrderForTerraformResponseData) GoString() string {
	return s.String()
}

func (s *GetCertificateApplicationOrderForTerraformResponseData) SetOrderId(v string) *GetCertificateApplicationOrderForTerraformResponseData {
	s.OrderId = &v
	return s
}

func (s *GetCertificateApplicationOrderForTerraformResponseData) SetOrderStatus(v string) *GetCertificateApplicationOrderForTerraformResponseData {
	s.OrderStatus = &v
	return s
}

func (s *GetCertificateApplicationOrderForTerraformResponseData) SetCertificateId(v int64) *GetCertificateApplicationOrderForTerraformResponseData {
	s.CertificateId = &v
	return s
}

func (s *GetCertificateApplicationOrderForTerraformResponseData) SetCertificateName(v string) *GetCertificateApplicationOrderForTerraformResponseData {
	s.CertificateName = &v
	return s
}

func (s *GetCertificateApplicationOrderForTerraformResponseData) SetCertificateBrand(v string) *GetCertificateApplicationOrderForTerraformResponseData {
	s.CertificateBrand = &v
	return s
}

func (s *GetCertificateApplicationOrderForTerraformResponseData) SetCertificateType(v string) *GetCertificateApplicationOrderForTerraformResponseData {
	s.CertificateType = &v
	return s
}

func (s *GetCertificateApplicationOrderForTerraformResponseData) SetDescription(v string) *GetCertificateApplicationOrderForTerraformResponseData {
	s.Description = &v
	return s
}

func (s *GetCertificateApplicationOrderForTerraformResponseData) SetAlgorithm(v string) *GetCertificateApplicationOrderForTerraformResponseData {
	s.Algorithm = &v
	return s
}

func (s *GetCertificateApplicationOrderForTerraformResponseData) SetAutoValidate(v string) *GetCertificateApplicationOrderForTerraformResponseData {
	s.AutoValidate = &v
	return s
}

func (s *GetCertificateApplicationOrderForTerraformResponseData) SetValidateMethod(v string) *GetCertificateApplicationOrderForTerraformResponseData {
	s.ValidateMethod = &v
	return s
}

func (s *GetCertificateApplicationOrderForTerraformResponseData) SetAutoRenew(v string) *GetCertificateApplicationOrderForTerraformResponseData {
	s.AutoRenew = &v
	return s
}

func (s *GetCertificateApplicationOrderForTerraformResponseData) SetValidityDays(v int) *GetCertificateApplicationOrderForTerraformResponseData {
	s.ValidityDays = &v
	return s
}

func (s *GetCertificateApplicationOrderForTerraformResponseData) SetCountry(v string) *GetCertificateApplicationOrderForTerraformResponseData {
	s.Country = &v
	return s
}

func (s *GetCertificateApplicationOrderForTerraformResponseData) SetState(v string) *GetCertificateApplicationOrderForTerraformResponseData {
	s.State = &v
	return s
}

func (s *GetCertificateApplicationOrderForTerraformResponseData) SetCity(v string) *GetCertificateApplicationOrderForTerraformResponseData {
	s.City = &v
	return s
}

func (s *GetCertificateApplicationOrderForTerraformResponseData) SetStreet(v string) *GetCertificateApplicationOrderForTerraformResponseData {
	s.Street = &v
	return s
}

func (s *GetCertificateApplicationOrderForTerraformResponseData) SetCompany(v string) *GetCertificateApplicationOrderForTerraformResponseData {
	s.Company = &v
	return s
}

func (s *GetCertificateApplicationOrderForTerraformResponseData) SetDepartment(v string) *GetCertificateApplicationOrderForTerraformResponseData {
	s.Department = &v
	return s
}

func (s *GetCertificateApplicationOrderForTerraformResponseData) SetCommonName(v string) *GetCertificateApplicationOrderForTerraformResponseData {
	s.CommonName = &v
	return s
}

func (s *GetCertificateApplicationOrderForTerraformResponseData) SetEmail(v string) *GetCertificateApplicationOrderForTerraformResponseData {
	s.Email = &v
	return s
}

func (s *GetCertificateApplicationOrderForTerraformResponseData) SetSubjectAlternativeNames(v []*string) *GetCertificateApplicationOrderForTerraformResponseData {
	s.SubjectAlternativeNames = v
	return s
}

func (s *GetCertificateApplicationOrderForTerraformResponseData) SetCreateTime(v string) *GetCertificateApplicationOrderForTerraformResponseData {
	s.CreateTime = &v
	return s
}

func (s *GetCertificateApplicationOrderForTerraformResponseData) SetErrorMessage(v string) *GetCertificateApplicationOrderForTerraformResponseData {
	s.ErrorMessage = &v
	return s
}

func (s *GetCertificateApplicationOrderForTerraformResponseData) SetCertificateSpec(v string) *GetCertificateApplicationOrderForTerraformResponseData {
	s.CertificateSpec = &v
	return s
}

func (s *GetCertificateApplicationOrderForTerraformResponseData) SetDomainType(v string) *GetCertificateApplicationOrderForTerraformResponseData {
	s.DomainType = &v
	return s
}

func (s *GetCertificateApplicationOrderForTerraformResponseData) SetAutoDeploy(v string) *GetCertificateApplicationOrderForTerraformResponseData {
	s.AutoDeploy = &v
	return s
}

func (s *GetCertificateApplicationOrderForTerraformResponseData) SetRemainValidityDays(v int) *GetCertificateApplicationOrderForTerraformResponseData {
	s.RemainValidityDays = &v
	return s
}

func (s *GetCertificateApplicationOrderForTerraformResponseData) SetDnsProviderInfos(v []*GetCertificateApplicationOrderForTerraformResponseDataDnsProviderInfos) *GetCertificateApplicationOrderForTerraformResponseData {
	s.DnsProviderInfos = v
	return s
}

func (s *GetCertificateApplicationOrderForTerraformResponseData) SetPrimaryCertificateBrand(v string) *GetCertificateApplicationOrderForTerraformResponseData {
	s.PrimaryCertificateBrand = &v
	return s
}

func (s *GetCertificateApplicationOrderForTerraformResponseData) SetBackupCertificateBrand(v string) *GetCertificateApplicationOrderForTerraformResponseData {
	s.BackupCertificateBrand = &v
	return s
}

func (s *GetCertificateApplicationOrderForTerraformResponseData) SetOrgValidateMethod(v string) *GetCertificateApplicationOrderForTerraformResponseData {
	s.OrgValidateMethod = &v
	return s
}

func (s *GetCertificateApplicationOrderForTerraformResponseData) SetPhone(v string) *GetCertificateApplicationOrderForTerraformResponseData {
	s.Phone = &v
	return s
}

func (s *GetCertificateApplicationOrderForTerraformResponseData) SetStreet1(v string) *GetCertificateApplicationOrderForTerraformResponseData {
	s.Street1 = &v
	return s
}

func (s *GetCertificateApplicationOrderForTerraformResponseData) SetPostalCode(v string) *GetCertificateApplicationOrderForTerraformResponseData {
	s.PostalCode = &v
	return s
}

func (s *GetCertificateApplicationOrderForTerraformResponseData) SetAdmin(v *GetCertificateApplicationOrderForTerraformResponseDataAdmin) *GetCertificateApplicationOrderForTerraformResponseData {
	s.Admin = v
	return s
}

func (s *GetCertificateApplicationOrderForTerraformResponseData) SetTech(v *GetCertificateApplicationOrderForTerraformResponseDataTech) *GetCertificateApplicationOrderForTerraformResponseData {
	s.Tech = v
	return s
}

type GetCertificateApplicationOrderForTerraformResponseDataDnsProviderInfos struct {
	// {"en":"domain","zh_CN":"域名"}
	Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
	// {"en":"Support CloudDNS, The optional certificate brand is LE or TrustAsia or GlobalSign or ZeroSSL;Support CloudFlare, The optional certificate brand is LE or ZeroSSL;Support DNSPod.cn, The optional certificate brand is LE or ZeroSSL;Support DNSPod.com, The optional certificate brand is LE or ZeroSSL;Support TencentDNSPod, The optional certificate brand is LE or ZeroSSL;Support Aliyun, The optional certificate brand is LE or ZeroSSL;Support GoDaddy, The optional certificate brand is LE or ZeroSSL;Support AmazonRoute53, The optional certificate brand is LE or ZeroSSL;Support GoogleDomains, The optional certificate brand is LE or ZeroSSL;Support AkamaiEdgeDNS, The optional certificate brand is LE or ZeroSSL.","zh_CN":"DNS托管商编码，支持CloudDNS，可选证书品牌为LE、TrustAsia、GlobalSign、ZeroSSL；支持CloudFlare，可选证书品牌为LE、ZeroSSL；支持DNSPod.cn，可选证书品牌为LE、ZeroSSL；支持DNSPod.com，可选证书品牌为LE、ZeroSSL；支持TencentDNSPod，可选证书品牌为LE、ZeroSSL；支持Aliyun，可选证书品牌为LE、ZeroSSL；支持GoDaddy，可选证书品牌为LE、ZeroSSL；支持AmazonRoute53，可选证书品牌为LE、ZeroSSL；支持GoogleDomains，可选证书品牌为LE、ZeroSSL；支持AkamaiEdgeDNS，可选证书品牌为LE、ZeroSSL。","exampleValue":"CloudDNS,CloudFlare"}
	DnsProviderCode *string `json:"dnsProviderCode,omitempty" xml:"dnsProviderCode,omitempty" require:"true"`
	// {"en":"DNS Api Access, JSON format:The hosting provider is CloudDNS, JSON KEY is accessKey and secretKey;The hosting provider is CloudFlare, The global API key JSON KEY is CF_Key CF_Email, The restricted API token JSON KEY is CF_Token, CF_ZoneID or CF_Sccount-ID.","zh_CN":"DNS托管商API凭证，JSON格式，托管商为CloudDNS，JSON KEY为accessKey、secretKey；托管商为CloudFlare，全局API密钥JSON KEY为CF_Key、CF_Email，限制性API令牌JSON KEY为CF_Token、CF_Zone_ID或CF_Account_ID；托管商为DNSPod.cn，JSON KEY为DP_Id、DP_Key；托管商为DNSPod.com，JSON KEY为DPI_Id、DPI_Key；托管商为TencentDNSPod，JSON KEY为Tencent_SecretId、Tencent_SecretKey；托管商为Aliyun，JSON KEY为Ali_Key、Ali_Secret；托管商为GoDaddy，JSON KEY为GD_Key、GD_Secret；托管商为AmazonRoute53，JSON KEY为AWS_ACCESS_KEY_ID、AWS_SECRET_ACCESS_KEY；托管商为GoogleDomains，JSON KEY为GOOGLEDOMAINS_ACCESS_TOKEN、GOOGLEDOMAINS_ZONE；托管商为AkamaiEdgeDNS，JSON KEY为AKAMAI_CLIENT_TOKEN、AKAMAI_ACCESS_TOKEN、AKAMAI_CLIENT_SECRET、AKAMAI_HOST；"}
	DnsApiAccess *string `json:"dnsApiAccess,omitempty" xml:"dnsApiAccess,omitempty" require:"true"`
	// {"defaultValue":"false","en":"Whether to use alias verification","zh_CN":"是否使用别名方式","exampleValue":"true,false"}
	EnableDnsAliasMode *string `json:"enableDnsAliasMode,omitempty" xml:"enableDnsAliasMode,omitempty" require:"true"`
	// {"en":"Validate alias","zh_CN":"验证别名"}
	ValidateAliasDomain *string `json:"validateAliasDomain,omitempty" xml:"validateAliasDomain,omitempty" require:"true"`
}

func (s GetCertificateApplicationOrderForTerraformResponseDataDnsProviderInfos) String() string {
	return tea.Prettify(s)
}

func (s GetCertificateApplicationOrderForTerraformResponseDataDnsProviderInfos) GoString() string {
	return s.String()
}

func (s *GetCertificateApplicationOrderForTerraformResponseDataDnsProviderInfos) SetDomain(v string) *GetCertificateApplicationOrderForTerraformResponseDataDnsProviderInfos {
	s.Domain = &v
	return s
}

func (s *GetCertificateApplicationOrderForTerraformResponseDataDnsProviderInfos) SetDnsProviderCode(v string) *GetCertificateApplicationOrderForTerraformResponseDataDnsProviderInfos {
	s.DnsProviderCode = &v
	return s
}

func (s *GetCertificateApplicationOrderForTerraformResponseDataDnsProviderInfos) SetDnsApiAccess(v string) *GetCertificateApplicationOrderForTerraformResponseDataDnsProviderInfos {
	s.DnsApiAccess = &v
	return s
}

func (s *GetCertificateApplicationOrderForTerraformResponseDataDnsProviderInfos) SetEnableDnsAliasMode(v string) *GetCertificateApplicationOrderForTerraformResponseDataDnsProviderInfos {
	s.EnableDnsAliasMode = &v
	return s
}

func (s *GetCertificateApplicationOrderForTerraformResponseDataDnsProviderInfos) SetValidateAliasDomain(v string) *GetCertificateApplicationOrderForTerraformResponseDataDnsProviderInfos {
	s.ValidateAliasDomain = &v
	return s
}

type GetCertificateApplicationOrderForTerraformResponseDataAdmin struct {
	// {"en":"firstName, Self verify organizational information. This field is mandatory","zh_CN":"名，自行验证组织信息此字段必填"}
	FirstName *string `json:"firstName,omitempty" xml:"firstName,omitempty" require:"true"`
	// {"en":"lastName, Self verify organizational information. This field is mandatory","zh_CN":"姓，自行验证组织信息此字段必填"}
	LastName *string `json:"lastName,omitempty" xml:"lastName,omitempty" require:"true"`
	// {"en":"phone, Self verify organizational information. This field is mandatory","zh_CN":"电话，自行验证组织信息此字段必填"}
	Phone *string `json:"phone,omitempty" xml:"phone,omitempty" require:"true"`
	// {"en":"email, Self verify organizational information. This field is mandatory","zh_CN":"邮箱，自行验证组织信息此字段必填"}
	Email *string `json:"email,omitempty" xml:"email,omitempty" require:"true"`
	// {"en":"title, Self verify organizational information. This field is mandatory","zh_CN":"职位，自行验证组织信息此字段必填"}
	Title *string `json:"title,omitempty" xml:"title,omitempty" require:"true"`
}

func (s GetCertificateApplicationOrderForTerraformResponseDataAdmin) String() string {
	return tea.Prettify(s)
}

func (s GetCertificateApplicationOrderForTerraformResponseDataAdmin) GoString() string {
	return s.String()
}

func (s *GetCertificateApplicationOrderForTerraformResponseDataAdmin) SetFirstName(v string) *GetCertificateApplicationOrderForTerraformResponseDataAdmin {
	s.FirstName = &v
	return s
}

func (s *GetCertificateApplicationOrderForTerraformResponseDataAdmin) SetLastName(v string) *GetCertificateApplicationOrderForTerraformResponseDataAdmin {
	s.LastName = &v
	return s
}

func (s *GetCertificateApplicationOrderForTerraformResponseDataAdmin) SetPhone(v string) *GetCertificateApplicationOrderForTerraformResponseDataAdmin {
	s.Phone = &v
	return s
}

func (s *GetCertificateApplicationOrderForTerraformResponseDataAdmin) SetEmail(v string) *GetCertificateApplicationOrderForTerraformResponseDataAdmin {
	s.Email = &v
	return s
}

func (s *GetCertificateApplicationOrderForTerraformResponseDataAdmin) SetTitle(v string) *GetCertificateApplicationOrderForTerraformResponseDataAdmin {
	s.Title = &v
	return s
}

type GetCertificateApplicationOrderForTerraformResponseDataTech struct {
	// {"en":"firstName, Self verify organizational information. This field is mandatory","zh_CN":"名，自行验证组织信息此字段必填"}
	FirstName *string `json:"firstName,omitempty" xml:"firstName,omitempty" require:"true"`
	// {"en":"lastName, Self verify organizational information. This field is mandatory","zh_CN":"姓，自行验证组织信息此字段必填"}
	LastName *string `json:"lastName,omitempty" xml:"lastName,omitempty" require:"true"`
	// {"en":"phone, Self verify organizational information. This field is mandatory","zh_CN":"电话，自行验证组织信息此字段必填"}
	Phone *string `json:"phone,omitempty" xml:"phone,omitempty" require:"true"`
	// {"en":"email, Self verify organizational information. This field is mandatory","zh_CN":"邮箱，自行验证组织信息此字段必填"}
	Email *string `json:"email,omitempty" xml:"email,omitempty" require:"true"`
	// {"en":"title, Self verify organizational information. This field is mandatory","zh_CN":"职位，自行验证组织信息此字段必填"}
	Title *string `json:"title,omitempty" xml:"title,omitempty" require:"true"`
}

func (s GetCertificateApplicationOrderForTerraformResponseDataTech) String() string {
	return tea.Prettify(s)
}

func (s GetCertificateApplicationOrderForTerraformResponseDataTech) GoString() string {
	return s.String()
}

func (s *GetCertificateApplicationOrderForTerraformResponseDataTech) SetFirstName(v string) *GetCertificateApplicationOrderForTerraformResponseDataTech {
	s.FirstName = &v
	return s
}

func (s *GetCertificateApplicationOrderForTerraformResponseDataTech) SetLastName(v string) *GetCertificateApplicationOrderForTerraformResponseDataTech {
	s.LastName = &v
	return s
}

func (s *GetCertificateApplicationOrderForTerraformResponseDataTech) SetPhone(v string) *GetCertificateApplicationOrderForTerraformResponseDataTech {
	s.Phone = &v
	return s
}

func (s *GetCertificateApplicationOrderForTerraformResponseDataTech) SetEmail(v string) *GetCertificateApplicationOrderForTerraformResponseDataTech {
	s.Email = &v
	return s
}

func (s *GetCertificateApplicationOrderForTerraformResponseDataTech) SetTitle(v string) *GetCertificateApplicationOrderForTerraformResponseDataTech {
	s.Title = &v
	return s
}

type GetCertificateApplicationOrderForTerraformResponseHeader struct {
}

func (s GetCertificateApplicationOrderForTerraformResponseHeader) String() string {
	return tea.Prettify(s)
}

func (s GetCertificateApplicationOrderForTerraformResponseHeader) GoString() string {
	return s.String()
}

type ListCertificateApplicationOrdersForTerraformRequest struct {
	// {"en":"Order ID","zh_CN":"订单ID"}
	OrderId *string `json:"orderId,omitempty" xml:"orderId,omitempty"`
	// {"en":"Status of certificate apply order.","zh_CN":"订单状态:\nACCEPT_SUCCESS：订单提交成功，待后台发起证书申请；\nAPPLYING：证书申请中，待域名验证；\nISSUE_SUCCESS：证书签发成功，可进行证书配置与部署；\nISSUE_FAILURE：证书签发失败；\nREVOKED：证书已吊销；\nCANCELED：已取消证书申请；","exampleValue":"ACCEPT_SUCCESS,APPLYING,ISSUE_SUCCESS,ISSUE_FAILURE,REVOKED,CANCELED"}
	OrderStatus []*string `json:"orderStatus,omitempty" xml:"orderStatus,omitempty" type:"Repeated"`
	// {"en":"Certificate Name","zh_CN":"证书名称"}
	CertificateName *string `json:"certificateName,omitempty" xml:"certificateName,omitempty"`
	// {"en":"Domain","zh_CN":"域名"}
	Domain *string `json:"domain,omitempty" xml:"domain,omitempty"`
	// {"en":"The time when the certificate order was created. The format is yyyy-MM-dd HH:mm:ss","zh_CN":"订单创建时间，格式为：yyyy-MM-dd HH:mm:ss"}
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// {"en":"The time when the certificate order was finished. The format is yyyy-MM-dd HH:mm:ss","zh_CN":"订单结束时间，格式为：yyyy-MM-dd HH:mm:ss。证书签发成功或证书签发失败，则记录结束时间。"}
	EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// {"en":"Page Parammeter.","zh_CN":"分页参数"}
	PageParam *ListCertificateApplicationOrdersForTerraformRequestPageParam `json:"pageParam,omitempty" xml:"pageParam,omitempty" type:"Struct"`
}

func (s ListCertificateApplicationOrdersForTerraformRequest) String() string {
	return tea.Prettify(s)
}

func (s ListCertificateApplicationOrdersForTerraformRequest) GoString() string {
	return s.String()
}

func (s *ListCertificateApplicationOrdersForTerraformRequest) SetOrderId(v string) *ListCertificateApplicationOrdersForTerraformRequest {
	s.OrderId = &v
	return s
}

func (s *ListCertificateApplicationOrdersForTerraformRequest) SetOrderStatus(v []*string) *ListCertificateApplicationOrdersForTerraformRequest {
	s.OrderStatus = v
	return s
}

func (s *ListCertificateApplicationOrdersForTerraformRequest) SetCertificateName(v string) *ListCertificateApplicationOrdersForTerraformRequest {
	s.CertificateName = &v
	return s
}

func (s *ListCertificateApplicationOrdersForTerraformRequest) SetDomain(v string) *ListCertificateApplicationOrdersForTerraformRequest {
	s.Domain = &v
	return s
}

func (s *ListCertificateApplicationOrdersForTerraformRequest) SetStartTime(v string) *ListCertificateApplicationOrdersForTerraformRequest {
	s.StartTime = &v
	return s
}

func (s *ListCertificateApplicationOrdersForTerraformRequest) SetEndTime(v string) *ListCertificateApplicationOrdersForTerraformRequest {
	s.EndTime = &v
	return s
}

func (s *ListCertificateApplicationOrdersForTerraformRequest) SetPageParam(v *ListCertificateApplicationOrdersForTerraformRequestPageParam) *ListCertificateApplicationOrdersForTerraformRequest {
	s.PageParam = v
	return s
}

type ListCertificateApplicationOrdersForTerraformRequestPageParam struct {
	// {"en":"Page Size,Range:1-100","zh_CN":"每页大小，1-100","exampleValue":"00"}
	PageSize *int `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// {"defaultValue":"0","en":"Page Number, the first page must be 0.Range:>=0","zh_CN":"页码，当前页，第一页从0开始，必须大于等于0"}
	PageNumber *int `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
}

func (s ListCertificateApplicationOrdersForTerraformRequestPageParam) String() string {
	return tea.Prettify(s)
}

func (s ListCertificateApplicationOrdersForTerraformRequestPageParam) GoString() string {
	return s.String()
}

func (s *ListCertificateApplicationOrdersForTerraformRequestPageParam) SetPageSize(v int) *ListCertificateApplicationOrdersForTerraformRequestPageParam {
	s.PageSize = &v
	return s
}

func (s *ListCertificateApplicationOrdersForTerraformRequestPageParam) SetPageNumber(v int) *ListCertificateApplicationOrdersForTerraformRequestPageParam {
	s.PageNumber = &v
	return s
}

type ListCertificateApplicationOrdersForTerraformRequestHeader struct {
}

func (s ListCertificateApplicationOrdersForTerraformRequestHeader) String() string {
	return tea.Prettify(s)
}

func (s ListCertificateApplicationOrdersForTerraformRequestHeader) GoString() string {
	return s.String()
}

type ListCertificateApplicationOrdersForTerraformPaths struct {
}

func (s ListCertificateApplicationOrdersForTerraformPaths) String() string {
	return tea.Prettify(s)
}

func (s ListCertificateApplicationOrdersForTerraformPaths) GoString() string {
	return s.String()
}

type ListCertificateApplicationOrdersForTerraformParameters struct {
}

func (s ListCertificateApplicationOrdersForTerraformParameters) String() string {
	return tea.Prettify(s)
}

func (s ListCertificateApplicationOrdersForTerraformParameters) GoString() string {
	return s.String()
}

type ListCertificateApplicationOrdersForTerraformResponse struct {
	// {"en":"Result Code, success is 0","zh_CN":"响应码，成功为0"}
	Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
	// {"en":"Result Message","zh_CN":"响应信息"}
	Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
	// {"en":"Result Data","zh_CN":"响应数据"}
	Data *ListCertificateApplicationOrdersForTerraformResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
}

func (s ListCertificateApplicationOrdersForTerraformResponse) String() string {
	return tea.Prettify(s)
}

func (s ListCertificateApplicationOrdersForTerraformResponse) GoString() string {
	return s.String()
}

func (s *ListCertificateApplicationOrdersForTerraformResponse) SetCode(v string) *ListCertificateApplicationOrdersForTerraformResponse {
	s.Code = &v
	return s
}

func (s *ListCertificateApplicationOrdersForTerraformResponse) SetMessage(v string) *ListCertificateApplicationOrdersForTerraformResponse {
	s.Message = &v
	return s
}

func (s *ListCertificateApplicationOrdersForTerraformResponse) SetData(v *ListCertificateApplicationOrdersForTerraformResponseData) *ListCertificateApplicationOrdersForTerraformResponse {
	s.Data = v
	return s
}

type ListCertificateApplicationOrdersForTerraformResponseData struct {
	// {"en":"Order list","zh_CN":"订单列表"}
	Orders []*ListCertificateApplicationOrdersForTerraformResponseDataOrders `json:"orders,omitempty" xml:"orders,omitempty" require:"true" type:"Repeated"`
	// {"en":"Page Info","zh_CN":"分页数据"}
	PageInfo *ListCertificateApplicationOrdersForTerraformResponseDataPageInfo `json:"pageInfo,omitempty" xml:"pageInfo,omitempty" require:"true" type:"Struct"`
}

func (s ListCertificateApplicationOrdersForTerraformResponseData) String() string {
	return tea.Prettify(s)
}

func (s ListCertificateApplicationOrdersForTerraformResponseData) GoString() string {
	return s.String()
}

func (s *ListCertificateApplicationOrdersForTerraformResponseData) SetOrders(v []*ListCertificateApplicationOrdersForTerraformResponseDataOrders) *ListCertificateApplicationOrdersForTerraformResponseData {
	s.Orders = v
	return s
}

func (s *ListCertificateApplicationOrdersForTerraformResponseData) SetPageInfo(v *ListCertificateApplicationOrdersForTerraformResponseDataPageInfo) *ListCertificateApplicationOrdersForTerraformResponseData {
	s.PageInfo = v
	return s
}

type ListCertificateApplicationOrdersForTerraformResponseDataOrders struct {
	// {"en":"Order ID","zh_CN":"订单ID"}
	OrderId *string `json:"orderId,omitempty" xml:"orderId,omitempty" require:"true"`
	// {"en":"Status of certificate apply order.","zh_CN":"证书订单状态","exampleValue":"ACCEPT_SUCCESS,APPLYING,ISSUE_SUCCESS,ISSUE_FAILURE,REVOKED,CANCELED"}
	OrderStatus *string `json:"orderStatus,omitempty" xml:"orderStatus,omitempty" require:"true"`
	// {"en":"Certificate ID","zh_CN":"证书ID"}
	CertificateId *int `json:"certificateId,omitempty" xml:"certificateId,omitempty" require:"true"`
	// {"en":"Certificate Name","zh_CN":"证书名称"}
	CertificateName *string `json:"certificateName,omitempty" xml:"certificateName,omitempty" require:"true"`
	// {"en":"Certificate Brand","zh_CN":"证书品牌"}
	CertificateBrand *string `json:"certificateBrand,omitempty" xml:"certificateBrand,omitempty" require:"true"`
	// {"en":"Certificate Type","zh_CN":"证书类型"}
	CertificateType *string `json:"certificateType,omitempty" xml:"certificateType,omitempty" require:"true"`
	// {"en":"Auto Renew","zh_CN":"是否自动续签"}
	AutoRenew *string `json:"autoRenew,omitempty" xml:"autoRenew,omitempty" require:"true"`
	// {"en":"Description","zh_CN":"描述"}
	Description *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	// {"en":"Common Name of the certificate","zh_CN":"证书的通用名称"}
	CommonName *string `json:"commonName,omitempty" xml:"commonName,omitempty" require:"true"`
	// {"en":"SAN","zh_CN":"主体备用名称"}
	SubjectAlternativeNames []*string `json:"subjectAlternativeNames,omitempty" xml:"subjectAlternativeNames,omitempty" require:"true" type:"Repeated"`
	// {"en":"Create Time","zh_CN":"创建时间"}
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty" require:"true"`
}

func (s ListCertificateApplicationOrdersForTerraformResponseDataOrders) String() string {
	return tea.Prettify(s)
}

func (s ListCertificateApplicationOrdersForTerraformResponseDataOrders) GoString() string {
	return s.String()
}

func (s *ListCertificateApplicationOrdersForTerraformResponseDataOrders) SetOrderId(v string) *ListCertificateApplicationOrdersForTerraformResponseDataOrders {
	s.OrderId = &v
	return s
}

func (s *ListCertificateApplicationOrdersForTerraformResponseDataOrders) SetOrderStatus(v string) *ListCertificateApplicationOrdersForTerraformResponseDataOrders {
	s.OrderStatus = &v
	return s
}

func (s *ListCertificateApplicationOrdersForTerraformResponseDataOrders) SetCertificateId(v int) *ListCertificateApplicationOrdersForTerraformResponseDataOrders {
	s.CertificateId = &v
	return s
}

func (s *ListCertificateApplicationOrdersForTerraformResponseDataOrders) SetCertificateName(v string) *ListCertificateApplicationOrdersForTerraformResponseDataOrders {
	s.CertificateName = &v
	return s
}

func (s *ListCertificateApplicationOrdersForTerraformResponseDataOrders) SetCertificateBrand(v string) *ListCertificateApplicationOrdersForTerraformResponseDataOrders {
	s.CertificateBrand = &v
	return s
}

func (s *ListCertificateApplicationOrdersForTerraformResponseDataOrders) SetCertificateType(v string) *ListCertificateApplicationOrdersForTerraformResponseDataOrders {
	s.CertificateType = &v
	return s
}

func (s *ListCertificateApplicationOrdersForTerraformResponseDataOrders) SetAutoRenew(v string) *ListCertificateApplicationOrdersForTerraformResponseDataOrders {
	s.AutoRenew = &v
	return s
}

func (s *ListCertificateApplicationOrdersForTerraformResponseDataOrders) SetDescription(v string) *ListCertificateApplicationOrdersForTerraformResponseDataOrders {
	s.Description = &v
	return s
}

func (s *ListCertificateApplicationOrdersForTerraformResponseDataOrders) SetCommonName(v string) *ListCertificateApplicationOrdersForTerraformResponseDataOrders {
	s.CommonName = &v
	return s
}

func (s *ListCertificateApplicationOrdersForTerraformResponseDataOrders) SetSubjectAlternativeNames(v []*string) *ListCertificateApplicationOrdersForTerraformResponseDataOrders {
	s.SubjectAlternativeNames = v
	return s
}

func (s *ListCertificateApplicationOrdersForTerraformResponseDataOrders) SetCreateTime(v string) *ListCertificateApplicationOrdersForTerraformResponseDataOrders {
	s.CreateTime = &v
	return s
}

type ListCertificateApplicationOrdersForTerraformResponseDataPageInfo struct {
	// {"en":"Total Number","zh_CN":"总数"}
	TotalNumber *int `json:"totalNumber,omitempty" xml:"totalNumber,omitempty" require:"true"`
	// {"en":"Page Number","zh_CN":"页码"}
	PageNumber *int `json:"pageNumber,omitempty" xml:"pageNumber,omitempty" require:"true"`
	// {"en":"Page Size","zh_CN":"每页大小"}
	PageSize *int `json:"pageSize,omitempty" xml:"pageSize,omitempty" require:"true"`
	// {"en":"Total Page Number","zh_CN":"总页码数"}
	TotalPageNumber *int `json:"totalPageNumber,omitempty" xml:"totalPageNumber,omitempty" require:"true"`
}

func (s ListCertificateApplicationOrdersForTerraformResponseDataPageInfo) String() string {
	return tea.Prettify(s)
}

func (s ListCertificateApplicationOrdersForTerraformResponseDataPageInfo) GoString() string {
	return s.String()
}

func (s *ListCertificateApplicationOrdersForTerraformResponseDataPageInfo) SetTotalNumber(v int) *ListCertificateApplicationOrdersForTerraformResponseDataPageInfo {
	s.TotalNumber = &v
	return s
}

func (s *ListCertificateApplicationOrdersForTerraformResponseDataPageInfo) SetPageNumber(v int) *ListCertificateApplicationOrdersForTerraformResponseDataPageInfo {
	s.PageNumber = &v
	return s
}

func (s *ListCertificateApplicationOrdersForTerraformResponseDataPageInfo) SetPageSize(v int) *ListCertificateApplicationOrdersForTerraformResponseDataPageInfo {
	s.PageSize = &v
	return s
}

func (s *ListCertificateApplicationOrdersForTerraformResponseDataPageInfo) SetTotalPageNumber(v int) *ListCertificateApplicationOrdersForTerraformResponseDataPageInfo {
	s.TotalPageNumber = &v
	return s
}

type ListCertificateApplicationOrdersForTerraformResponseHeader struct {
}

func (s ListCertificateApplicationOrdersForTerraformResponseHeader) String() string {
	return tea.Prettify(s)
}

func (s ListCertificateApplicationOrdersForTerraformResponseHeader) GoString() string {
	return s.String()
}

type CreateCertificateApplicationOrderForTerraformRequest struct {
	// {"en":"Name of the certificate. The certificate name cannot be the same as your existing certificate.\nRange: <= 128 characters.","zh_CN":"证书名称，最长不超过128个字符。证书名称不允许与已有的证书重名."}
	CertificateName *string `json:"certificateName,omitempty" xml:"certificateName,omitempty"`
	// {"en":"A description of the new certificate.\nRange: <= 256 characters.","zh_CN":"证书描述，最长不超过256个字符。"}
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// {"defaultValue":"false","en":"Automatically Renew your certificate. Selecting 'true' will reduce customization of other fields: 'validateMethod' must be 'HTTP', the certificate domain must configure in our acceleration platform; the domain  must CNAME to our acceleration platform.","zh_CN":"是否自动续签。选择是否在证书到期前为您自动续订证书，开启自动续签，验证方式必须选择'HTTP'；申请证书包含的域名必须已在我司加速平台配置，且域名必须已CNAME到我司加速平台。","exampleValue":"true,false"}
	AutoRenew *string `json:"autoRenew,omitempty" xml:"autoRenew,omitempty"`
	// {"en":"Certificate Brand. Selecting 'LE' means we apply your certificate through Let's Encrypt (https://letsencrypt.org/docs/challenge-types/) .","zh_CN":"证书品牌。LE即Let's Encrypt。","exampleValue":"LE,TrustAsia,GlobalSign"}
	CertificateBrand *string `json:"certificateBrand,omitempty" xml:"certificateBrand,omitempty" require:"true"`
	// {"en":"Certificate Specification,\nCertificate brand GlobalSign, certificate specification GlobalSignOVMSSSL, certificate type OV, certificate validity days [360], billing;\nCertificate brand GlobalSign, certificate specification GlobalSignOVSANsMSSL, certificate type OV, certificate validity days [360], billing;\nCertificate brand TrustAsia, certificate specification TrustAsia TLSC1, certificate type DV, certificate validity days [360], billing;\nCertificate brand TrustAsia, certificate specification TrustAsiaTLSSANsC1, certificate type DV, certificate validity days [360], billing;\nCertificate brand TrustAsia, certificate specification TrustAsiaC1DVFree, certificate type DV, certificate validity days [90], free of charge;\nCertificate brand TrustAsia, certificate specification TrustAsiaC1DVFreeSANS, certificate type DV, certificate validity days [90], free of charge;\nCertificate brand LE, certificate specification LetsEncryptDVFree, certificate type DV, certificate validity days [90], free of charge;\nCertificate Brand LE, Certificate Specification LetsEncryptDVFreeSANS, Certificate Type DV, Certificate Validity Days [90], free of charge","zh_CN":"证书规格,\n证书品牌GlobalSign、证书规格GlobalSignOVMSSL、证书类型OV、证书有效天数[360]、计费;\n证书品牌GlobalSign、证书规格GlobalSignOVSANsMSSL、证书类型OV、证书有效天数[360]、计费;\n证书品牌TrustAsia、证书规格TrustAsiaTLSC1、证书类型DV、证书有效天数[360]、计费;\n证书品牌TrustAsia、证书规格TrustAsiaTLSSANsC1、证书类型DV、证书有效天数[360]、计费;\n证书品牌TrustAsia、证书规格TrustAsiaC1DVFree、证书类型DV、证书有效天数[90]、免费;\n证书品牌TrustAsia、证书规格TrustAsiaC1DVFreeSANS、证书类型DV、证书有效天数[90]、免费;\n证书品牌LE、证书规格LetsEncryptDVFree、证书类型DV、证书有效天数[90]、免费;\n证书品牌LE、证书规格LetsEncryptDVFreeSANS、证书类型DV、证书有效天数[90]、免费","exampleValue":"GlobalSignOVMSSL,GlobalSignOVSANsMSSL,TrustAsiaTLSC1,TrustAsiaTLSSANsC1,TrustAsiaC1DVFree,TrustAsiaC1DVFreeSANS,LetsEncryptDVFree,LetsEncryptDVFreeSANS"}
	CertificateSpec *string `json:"certificateSpec,omitempty" xml:"certificateSpec,omitempty" require:"true"`
	// {"en":"Certificate Type. When 'certificateBrand' specifying 'LE',  'certificateType' only supports 'DV'.","zh_CN":"安全等级。 Let's Encrypt仅支持DV， TrustAsia支持DV和OV。","exampleValue":"DV,OV"}
	CertificateType *string `json:"certificateType,omitempty" xml:"certificateType,omitempty" require:"true"`
	// {"en":"Certificate Algorithm. When 'certificateBrand' specifying 'LE',  'certificateType' only supports 'RSA2048' and 'ECDSA256'.","zh_CN":"证书算法。Let's Encrypt仅支持RSA2048与ECDSA256  。","exampleValue":"RSA2048,ECDSA256"}
	Algorithm *string `json:"algorithm,omitempty" xml:"algorithm,omitempty" require:"true"`
	// {"en":"Automatically validate that you control the domains in the certificate you applied. Selecting 'true' will reduce customization of other fields: 'validateMethod' must be 'HTTP' , the certificate domain must configure in our acceleration platform; the domain must CNAME to our acceleration platform.","zh_CN":"是否自动验证。选择是否帮您自动完成证书域名控制权验证。开启自动验证，验证方式必须选择'HTTP'；申请证书包含的域名必须已在我司加速平台配置，且域名必须已CNAME到我司加速平台。","exampleValue":"true,false"}
	AutoValidate *string `json:"autoValidate,omitempty" xml:"autoValidate,omitempty" require:"true"`
	// {"en":"Validate Method selected for the order to validate that you control the domains in the certificate. Validate Method must be 'DNS' when 'subjectAlternativeNames'  is a wildcard hostname beginning with '*' such as '*.example.com' .\nWhen applying for a certificate with an IP domain name, only HTTP authentication is supported.","zh_CN":"验证方式。选择证书域名控制权的验证方式。当申请证书域名'subjectAlternativeNames'为泛域名时，仅支持DNS验证。当申请证书域名为IP时，仅支持HTTP验证。","exampleValue":"HTTP,DNS"}
	ValidateMethod *string `json:"validateMethod,omitempty" xml:"validateMethod,omitempty" require:"true"`
	// {"en":"Whether to deploy automatically. true: automatic deployment, false: manual deployment.","zh_CN":"是否自动部署，true:自动部署，false:手动部署.","exampleValue":"true,false"}
	AutoDeploy *string `json:"autoDeploy,omitempty" xml:"autoDeploy,omitempty" require:"true"`
	// {"en":"Validity Days selected for the certificate. Please refer to the description in the certificateSpec for details.","zh_CN":"证书有效天数，详见certificateSpec描述"}
	ValidityDays *int `json:"validityDays,omitempty" xml:"validityDays,omitempty" require:"true"`
	// {"en":"Certificate Signing Request Informtion.","zh_CN":"证书签名请求信息 。"}
	IdentificationInfo *CreateCertificateApplicationOrderForTerraformRequestIdentificationInfo `json:"identificationInfo,omitempty" xml:"identificationInfo,omitempty" require:"true" type:"Struct"`
	// {"en":"Backup Certificate Brand.\n\nSingle Domain:\n- LE (Backup: TrustAsia, ZeroSSL)\n- ZeroSSL (Backup: TrustAsia, LE)\n- TrustAsia (Backup: LE, ZeroSSL)\n\nMulti-Domain:\n- LE (Backup: TrustAsia)\n- TrustAsia (Backup: LE)","zh_CN":"备用CA证书品牌。\n\n单域名:\n- LE (备用: TrustAsia, ZeroSSL)\n- ZeroSSL (备用: TrustAsia, LE)\n- TrustAsia (备用: LE, ZeroSSL)\n\n多域名:\n- LE (备用: TrustAsia)\n- TrustAsia (备用: LE)","exampleValue":"LE,TrustAsia,ZeroSSL"}
	BackupCertificateBrand *string `json:"backupCertificateBrand,omitempty" xml:"backupCertificateBrand,omitempty"`
	// {"en":"orgValidateMethod, default:  Enable pre audit, self_validate:  Self verify organizational information, none: Unorganized information, default value default.","zh_CN":"组织验证方式, default: 开启预审核、self_validate: 自行验证组织信息、none: 无组织信息，默认值default。"}
	OrgValidateMethod *string `json:"orgValidateMethod,omitempty" xml:"orgValidateMethod,omitempty"`
	// {"en":"domainType","zh_CN":"域名类型","exampleValue":"single,multi"}
	DomainType *string `json:"domainType,omitempty" xml:"domainType,omitempty"`
	// {"defaultValue":"false","en":"batchApply","zh_CN":"是否批量申请","exampleValue":"true,false"}
	BatchApply *string `json:"batchApply,omitempty" xml:"batchApply,omitempty"`
	// {"en":"admin","zh_CN":"订单管理联系人"}
	Admin *CreateCertificateApplicationOrderForTerraformRequestAdmin `json:"admin,omitempty" xml:"admin,omitempty" type:"Struct"`
	// {"en":"Technical contact person for the order.","zh_CN":"订单技术联系人"}
	Tech *CreateCertificateApplicationOrderForTerraformRequestTech `json:"tech,omitempty" xml:"tech,omitempty" type:"Struct"`
	// {"en":"DNS Provider Infomation","zh_CN":"DNS托管商信息"}
	DnsProviderInfos []*CreateCertificateApplicationOrderForTerraformRequestDnsProviderInfos `json:"dnsProviderInfos,omitempty" xml:"dnsProviderInfos,omitempty" type:"Repeated"`
}

func (s CreateCertificateApplicationOrderForTerraformRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateCertificateApplicationOrderForTerraformRequest) GoString() string {
	return s.String()
}

func (s *CreateCertificateApplicationOrderForTerraformRequest) SetCertificateName(v string) *CreateCertificateApplicationOrderForTerraformRequest {
	s.CertificateName = &v
	return s
}

func (s *CreateCertificateApplicationOrderForTerraformRequest) SetDescription(v string) *CreateCertificateApplicationOrderForTerraformRequest {
	s.Description = &v
	return s
}

func (s *CreateCertificateApplicationOrderForTerraformRequest) SetAutoRenew(v string) *CreateCertificateApplicationOrderForTerraformRequest {
	s.AutoRenew = &v
	return s
}

func (s *CreateCertificateApplicationOrderForTerraformRequest) SetCertificateBrand(v string) *CreateCertificateApplicationOrderForTerraformRequest {
	s.CertificateBrand = &v
	return s
}

func (s *CreateCertificateApplicationOrderForTerraformRequest) SetCertificateSpec(v string) *CreateCertificateApplicationOrderForTerraformRequest {
	s.CertificateSpec = &v
	return s
}

func (s *CreateCertificateApplicationOrderForTerraformRequest) SetCertificateType(v string) *CreateCertificateApplicationOrderForTerraformRequest {
	s.CertificateType = &v
	return s
}

func (s *CreateCertificateApplicationOrderForTerraformRequest) SetAlgorithm(v string) *CreateCertificateApplicationOrderForTerraformRequest {
	s.Algorithm = &v
	return s
}

func (s *CreateCertificateApplicationOrderForTerraformRequest) SetAutoValidate(v string) *CreateCertificateApplicationOrderForTerraformRequest {
	s.AutoValidate = &v
	return s
}

func (s *CreateCertificateApplicationOrderForTerraformRequest) SetValidateMethod(v string) *CreateCertificateApplicationOrderForTerraformRequest {
	s.ValidateMethod = &v
	return s
}

func (s *CreateCertificateApplicationOrderForTerraformRequest) SetAutoDeploy(v string) *CreateCertificateApplicationOrderForTerraformRequest {
	s.AutoDeploy = &v
	return s
}

func (s *CreateCertificateApplicationOrderForTerraformRequest) SetValidityDays(v int) *CreateCertificateApplicationOrderForTerraformRequest {
	s.ValidityDays = &v
	return s
}

func (s *CreateCertificateApplicationOrderForTerraformRequest) SetIdentificationInfo(v *CreateCertificateApplicationOrderForTerraformRequestIdentificationInfo) *CreateCertificateApplicationOrderForTerraformRequest {
	s.IdentificationInfo = v
	return s
}

func (s *CreateCertificateApplicationOrderForTerraformRequest) SetBackupCertificateBrand(v string) *CreateCertificateApplicationOrderForTerraformRequest {
	s.BackupCertificateBrand = &v
	return s
}

func (s *CreateCertificateApplicationOrderForTerraformRequest) SetOrgValidateMethod(v string) *CreateCertificateApplicationOrderForTerraformRequest {
	s.OrgValidateMethod = &v
	return s
}

func (s *CreateCertificateApplicationOrderForTerraformRequest) SetDomainType(v string) *CreateCertificateApplicationOrderForTerraformRequest {
	s.DomainType = &v
	return s
}

func (s *CreateCertificateApplicationOrderForTerraformRequest) SetBatchApply(v string) *CreateCertificateApplicationOrderForTerraformRequest {
	s.BatchApply = &v
	return s
}

func (s *CreateCertificateApplicationOrderForTerraformRequest) SetAdmin(v *CreateCertificateApplicationOrderForTerraformRequestAdmin) *CreateCertificateApplicationOrderForTerraformRequest {
	s.Admin = v
	return s
}

func (s *CreateCertificateApplicationOrderForTerraformRequest) SetTech(v *CreateCertificateApplicationOrderForTerraformRequestTech) *CreateCertificateApplicationOrderForTerraformRequest {
	s.Tech = v
	return s
}

func (s *CreateCertificateApplicationOrderForTerraformRequest) SetDnsProviderInfos(v []*CreateCertificateApplicationOrderForTerraformRequestDnsProviderInfos) *CreateCertificateApplicationOrderForTerraformRequest {
	s.DnsProviderInfos = v
	return s
}

type CreateCertificateApplicationOrderForTerraformRequestIdentificationInfo struct {
	// {"en":"Country. An ISO-3166 country code.\nRange: = 2 characters","zh_CN":"国家，2个字符的ISO-3166国家代码"}
	Country *string `json:"country,omitempty" xml:"country,omitempty"`
	// {"en":"A state or province.","zh_CN":"省"}
	State *string `json:"state,omitempty" xml:"state,omitempty"`
	// {"en":"A city","zh_CN":"城市"}
	City *string `json:"city,omitempty" xml:"city,omitempty"`
	// {"en":"A company name","zh_CN":"公司"}
	Company *string `json:"company,omitempty" xml:"company,omitempty"`
	// {"en":"The department associated with the certificate","zh_CN":"部门"}
	Department *string `json:"department,omitempty" xml:"department,omitempty"`
	// {"en":"A common name of the certificate","zh_CN":"证书的通用名称"}
	CommonName *string `json:"commonName,omitempty" xml:"commonName,omitempty" require:"true"`
	// {"en":"Email address","zh_CN":"邮箱"}
	Email *string `json:"email,omitempty" xml:"email,omitempty"`
	// {"en":"The street where the company is located","zh_CN":"街道"}
	Street *string `json:"street,omitempty" xml:"street,omitempty"`
	// {"en":"Hostnames that this certificate will serve. Each entry must be a valid hostname such as example.com or a wildcard hostname beginning with '*' such as.example.com.","zh_CN":"此证书将提供服务的域名。 每个条目必须是有效的域名（例如 example.com）或以“*”开头的泛域名（例如 “*.example.com”）"}
	SubjectAlternativeNames []*string `json:"subjectAlternativeNames,omitempty" xml:"subjectAlternativeNames,omitempty" require:"true" type:"Repeated"`
	// {"en":"street1","zh_CN":"街道"}
	Street1 *string `json:"street1,omitempty" xml:"street1,omitempty"`
	// {"en":"phone","zh_CN":"联系电话"}
	Phone *string `json:"phone,omitempty" xml:"phone,omitempty"`
	// {"en":"postalCode","zh_CN":"邮政编码"}
	PostalCode *string `json:"postalCode,omitempty" xml:"postalCode,omitempty"`
}

func (s CreateCertificateApplicationOrderForTerraformRequestIdentificationInfo) String() string {
	return tea.Prettify(s)
}

func (s CreateCertificateApplicationOrderForTerraformRequestIdentificationInfo) GoString() string {
	return s.String()
}

func (s *CreateCertificateApplicationOrderForTerraformRequestIdentificationInfo) SetCountry(v string) *CreateCertificateApplicationOrderForTerraformRequestIdentificationInfo {
	s.Country = &v
	return s
}

func (s *CreateCertificateApplicationOrderForTerraformRequestIdentificationInfo) SetState(v string) *CreateCertificateApplicationOrderForTerraformRequestIdentificationInfo {
	s.State = &v
	return s
}

func (s *CreateCertificateApplicationOrderForTerraformRequestIdentificationInfo) SetCity(v string) *CreateCertificateApplicationOrderForTerraformRequestIdentificationInfo {
	s.City = &v
	return s
}

func (s *CreateCertificateApplicationOrderForTerraformRequestIdentificationInfo) SetCompany(v string) *CreateCertificateApplicationOrderForTerraformRequestIdentificationInfo {
	s.Company = &v
	return s
}

func (s *CreateCertificateApplicationOrderForTerraformRequestIdentificationInfo) SetDepartment(v string) *CreateCertificateApplicationOrderForTerraformRequestIdentificationInfo {
	s.Department = &v
	return s
}

func (s *CreateCertificateApplicationOrderForTerraformRequestIdentificationInfo) SetCommonName(v string) *CreateCertificateApplicationOrderForTerraformRequestIdentificationInfo {
	s.CommonName = &v
	return s
}

func (s *CreateCertificateApplicationOrderForTerraformRequestIdentificationInfo) SetEmail(v string) *CreateCertificateApplicationOrderForTerraformRequestIdentificationInfo {
	s.Email = &v
	return s
}

func (s *CreateCertificateApplicationOrderForTerraformRequestIdentificationInfo) SetStreet(v string) *CreateCertificateApplicationOrderForTerraformRequestIdentificationInfo {
	s.Street = &v
	return s
}

func (s *CreateCertificateApplicationOrderForTerraformRequestIdentificationInfo) SetSubjectAlternativeNames(v []*string) *CreateCertificateApplicationOrderForTerraformRequestIdentificationInfo {
	s.SubjectAlternativeNames = v
	return s
}

func (s *CreateCertificateApplicationOrderForTerraformRequestIdentificationInfo) SetStreet1(v string) *CreateCertificateApplicationOrderForTerraformRequestIdentificationInfo {
	s.Street1 = &v
	return s
}

func (s *CreateCertificateApplicationOrderForTerraformRequestIdentificationInfo) SetPhone(v string) *CreateCertificateApplicationOrderForTerraformRequestIdentificationInfo {
	s.Phone = &v
	return s
}

func (s *CreateCertificateApplicationOrderForTerraformRequestIdentificationInfo) SetPostalCode(v string) *CreateCertificateApplicationOrderForTerraformRequestIdentificationInfo {
	s.PostalCode = &v
	return s
}

type CreateCertificateApplicationOrderForTerraformRequestAdmin struct {
	// {"en":"firstName, Self verify organizational information. This field is mandatory","zh_CN":"名，自行验证组织信息此字段必填"}
	FirstName *string `json:"firstName,omitempty" xml:"firstName,omitempty"`
	// {"en":"lastName, Self verify organizational information. This field is mandatory","zh_CN":"姓，自行验证组织信息此字段必填"}
	LastName *string `json:"lastName,omitempty" xml:"lastName,omitempty"`
	// {"en":"phone, Self verify organizational information. This field is mandatory","zh_CN":"电话，自行验证组织信息此字段必填"}
	Phone *string `json:"phone,omitempty" xml:"phone,omitempty"`
	// {"en":"email, Self verify organizational information. This field is mandatory","zh_CN":"邮箱，自行验证组织信息此字段必填"}
	Email *string `json:"email,omitempty" xml:"email,omitempty"`
	// {"en":"title, Self verify organizational information. This field is mandatory","zh_CN":"职位，自行验证组织信息此字段必填"}
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s CreateCertificateApplicationOrderForTerraformRequestAdmin) String() string {
	return tea.Prettify(s)
}

func (s CreateCertificateApplicationOrderForTerraformRequestAdmin) GoString() string {
	return s.String()
}

func (s *CreateCertificateApplicationOrderForTerraformRequestAdmin) SetFirstName(v string) *CreateCertificateApplicationOrderForTerraformRequestAdmin {
	s.FirstName = &v
	return s
}

func (s *CreateCertificateApplicationOrderForTerraformRequestAdmin) SetLastName(v string) *CreateCertificateApplicationOrderForTerraformRequestAdmin {
	s.LastName = &v
	return s
}

func (s *CreateCertificateApplicationOrderForTerraformRequestAdmin) SetPhone(v string) *CreateCertificateApplicationOrderForTerraformRequestAdmin {
	s.Phone = &v
	return s
}

func (s *CreateCertificateApplicationOrderForTerraformRequestAdmin) SetEmail(v string) *CreateCertificateApplicationOrderForTerraformRequestAdmin {
	s.Email = &v
	return s
}

func (s *CreateCertificateApplicationOrderForTerraformRequestAdmin) SetTitle(v string) *CreateCertificateApplicationOrderForTerraformRequestAdmin {
	s.Title = &v
	return s
}

type CreateCertificateApplicationOrderForTerraformRequestTech struct {
	// {"en":"firstName, Self verify organizational information. This field is mandatory","zh_CN":"名，自行验证组织信息此字段必填"}
	FirstName *string `json:"firstName,omitempty" xml:"firstName,omitempty"`
	// {"en":"lastName, Self verify organizational information. This field is mandatory","zh_CN":"姓，自行验证组织信息此字段必填"}
	LastName *string `json:"lastName,omitempty" xml:"lastName,omitempty"`
	// {"en":"phone, Self verify organizational information. This field is mandatory","zh_CN":"电话，自行验证组织信息此字段必填"}
	Phone *string `json:"phone,omitempty" xml:"phone,omitempty"`
	// {"en":"email, Self verify organizational information. This field is mandatory","zh_CN":"邮箱，自行验证组织信息此字段必填"}
	Email *string `json:"email,omitempty" xml:"email,omitempty"`
	// {"en":"title, Self verify organizational information. This field is mandatory","zh_CN":"职位，自行验证组织信息此字段必填"}
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s CreateCertificateApplicationOrderForTerraformRequestTech) String() string {
	return tea.Prettify(s)
}

func (s CreateCertificateApplicationOrderForTerraformRequestTech) GoString() string {
	return s.String()
}

func (s *CreateCertificateApplicationOrderForTerraformRequestTech) SetFirstName(v string) *CreateCertificateApplicationOrderForTerraformRequestTech {
	s.FirstName = &v
	return s
}

func (s *CreateCertificateApplicationOrderForTerraformRequestTech) SetLastName(v string) *CreateCertificateApplicationOrderForTerraformRequestTech {
	s.LastName = &v
	return s
}

func (s *CreateCertificateApplicationOrderForTerraformRequestTech) SetPhone(v string) *CreateCertificateApplicationOrderForTerraformRequestTech {
	s.Phone = &v
	return s
}

func (s *CreateCertificateApplicationOrderForTerraformRequestTech) SetEmail(v string) *CreateCertificateApplicationOrderForTerraformRequestTech {
	s.Email = &v
	return s
}

func (s *CreateCertificateApplicationOrderForTerraformRequestTech) SetTitle(v string) *CreateCertificateApplicationOrderForTerraformRequestTech {
	s.Title = &v
	return s
}

type CreateCertificateApplicationOrderForTerraformRequestDnsProviderInfos struct {
	// {"en":"domain","zh_CN":"域名"}
	Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
	// {"en":"Support CloudDNS, The optional certificate brand is LE or TrustAsia or GlobalSign or ZeroSSL;Support CloudFlare, The optional certificate brand is LE or ZeroSSL;Support DNSPod.cn, The optional certificate brand is LE or ZeroSSL;Support DNSPod.com, The optional certificate brand is LE or ZeroSSL;Support TencentDNSPod, The optional certificate brand is LE or ZeroSSL;Support Aliyun, The optional certificate brand is LE or ZeroSSL;Support GoDaddy, The optional certificate brand is LE or ZeroSSL;Support AmazonRoute53, The optional certificate brand is LE or ZeroSSL;Support GoogleDomains, The optional certificate brand is LE or ZeroSSL;Support AkamaiEdgeDNS, The optional certificate brand is LE or ZeroSSL.","zh_CN":"DNS托管商编码，支持CloudDNS，可选证书品牌为LE、TrustAsia、GlobalSign、ZeroSSL；支持CloudFlare，可选证书品牌为LE、ZeroSSL；支持DNSPod.cn，可选证书品牌为LE、ZeroSSL；支持DNSPod.com，可选证书品牌为LE、ZeroSSL；支持TencentDNSPod，可选证书品牌为LE、ZeroSSL；支持Aliyun，可选证书品牌为LE、ZeroSSL；支持GoDaddy，可选证书品牌为LE、ZeroSSL；支持AmazonRoute53，可选证书品牌为LE、ZeroSSL；支持GoogleDomains，可选证书品牌为LE、ZeroSSL；支持AkamaiEdgeDNS，可选证书品牌为LE、ZeroSSL。","exampleValue":"CloudDNS,CloudFlare"}
	DnsProviderCode *string `json:"dnsProviderCode,omitempty" xml:"dnsProviderCode,omitempty"`
	// {"en":"DNS Api Access, JSON format:The hosting provider is CloudDNS, JSON KEY is accessKey and secretKey;The hosting provider is CloudFlare, The global API key JSON KEY is CF_Key CF_Email, The restricted API token JSON KEY is CF_Token, CF_ZoneID or CF_Sccount-ID.","zh_CN":"DNS托管商API凭证，JSON格式，托管商为CloudDNS，JSON KEY为accessKey、secretKey；托管商为CloudFlare，全局API密钥JSON KEY为CF_Key、CF_Email，限制性API令牌JSON KEY为CF_Token、CF_Zone_ID或CF_Account_ID；托管商为DNSPod.cn，JSON KEY为DP_Id、DP_Key；托管商为DNSPod.com，JSON KEY为DPI_Id、DPI_Key；托管商为TencentDNSPod，JSON KEY为Tencent_SecretId、Tencent_SecretKey；托管商为Aliyun，JSON KEY为Ali_Key、Ali_Secret；托管商为GoDaddy，JSON KEY为GD_Key、GD_Secret；托管商为AmazonRoute53，JSON KEY为AWS_ACCESS_KEY_ID、AWS_SECRET_ACCESS_KEY；托管商为GoogleDomains，JSON KEY为GOOGLEDOMAINS_ACCESS_TOKEN、GOOGLEDOMAINS_ZONE；托管商为AkamaiEdgeDNS，JSON KEY为AKAMAI_CLIENT_TOKEN、AKAMAI_ACCESS_TOKEN、AKAMAI_CLIENT_SECRET、AKAMAI_HOST；"}
	DnsApiAccess *string `json:"dnsApiAccess,omitempty" xml:"dnsApiAccess,omitempty"`
	// {"defaultValue":"false","en":"Whether to use alias verification","zh_CN":"是否使用别名方式","exampleValue":"true,false"}
	EnableDnsAliasMode *string `json:"enableDnsAliasMode,omitempty" xml:"enableDnsAliasMode,omitempty"`
	// {"en":"Validate alias","zh_CN":"验证别名"}
	ValidateAliasDomain *string `json:"validateAliasDomain,omitempty" xml:"validateAliasDomain,omitempty"`
}

func (s CreateCertificateApplicationOrderForTerraformRequestDnsProviderInfos) String() string {
	return tea.Prettify(s)
}

func (s CreateCertificateApplicationOrderForTerraformRequestDnsProviderInfos) GoString() string {
	return s.String()
}

func (s *CreateCertificateApplicationOrderForTerraformRequestDnsProviderInfos) SetDomain(v string) *CreateCertificateApplicationOrderForTerraformRequestDnsProviderInfos {
	s.Domain = &v
	return s
}

func (s *CreateCertificateApplicationOrderForTerraformRequestDnsProviderInfos) SetDnsProviderCode(v string) *CreateCertificateApplicationOrderForTerraformRequestDnsProviderInfos {
	s.DnsProviderCode = &v
	return s
}

func (s *CreateCertificateApplicationOrderForTerraformRequestDnsProviderInfos) SetDnsApiAccess(v string) *CreateCertificateApplicationOrderForTerraformRequestDnsProviderInfos {
	s.DnsApiAccess = &v
	return s
}

func (s *CreateCertificateApplicationOrderForTerraformRequestDnsProviderInfos) SetEnableDnsAliasMode(v string) *CreateCertificateApplicationOrderForTerraformRequestDnsProviderInfos {
	s.EnableDnsAliasMode = &v
	return s
}

func (s *CreateCertificateApplicationOrderForTerraformRequestDnsProviderInfos) SetValidateAliasDomain(v string) *CreateCertificateApplicationOrderForTerraformRequestDnsProviderInfos {
	s.ValidateAliasDomain = &v
	return s
}

type CreateCertificateApplicationOrderForTerraformRequestHeader struct {
}

func (s CreateCertificateApplicationOrderForTerraformRequestHeader) String() string {
	return tea.Prettify(s)
}

func (s CreateCertificateApplicationOrderForTerraformRequestHeader) GoString() string {
	return s.String()
}

type CreateCertificateApplicationOrderForTerraformPaths struct {
}

func (s CreateCertificateApplicationOrderForTerraformPaths) String() string {
	return tea.Prettify(s)
}

func (s CreateCertificateApplicationOrderForTerraformPaths) GoString() string {
	return s.String()
}

type CreateCertificateApplicationOrderForTerraformParameters struct {
}

func (s CreateCertificateApplicationOrderForTerraformParameters) String() string {
	return tea.Prettify(s)
}

func (s CreateCertificateApplicationOrderForTerraformParameters) GoString() string {
	return s.String()
}

type CreateCertificateApplicationOrderForTerraformResponse struct {
	// {"en":"Result Code, success is 0","zh_CN":"响应码，成功为0"}
	Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
	// {"en":"Result Message","zh_CN":"响应信息"}
	Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
	// {"en":"Result Data","zh_CN":"响应数据"}
	Data *CreateCertificateApplicationOrderForTerraformResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
}

func (s CreateCertificateApplicationOrderForTerraformResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateCertificateApplicationOrderForTerraformResponse) GoString() string {
	return s.String()
}

func (s *CreateCertificateApplicationOrderForTerraformResponse) SetCode(v string) *CreateCertificateApplicationOrderForTerraformResponse {
	s.Code = &v
	return s
}

func (s *CreateCertificateApplicationOrderForTerraformResponse) SetMessage(v string) *CreateCertificateApplicationOrderForTerraformResponse {
	s.Message = &v
	return s
}

func (s *CreateCertificateApplicationOrderForTerraformResponse) SetData(v *CreateCertificateApplicationOrderForTerraformResponseData) *CreateCertificateApplicationOrderForTerraformResponse {
	s.Data = v
	return s
}

type CreateCertificateApplicationOrderForTerraformResponseData struct {
	// {"en":"Order ID","zh_CN":"订单ID"}
	OrderId *string `json:"orderId,omitempty" xml:"orderId,omitempty" require:"true"`
	// {"en":"Order ID List","zh_CN":"订单ID列表"}
	OrderIds []*string `json:"orderIds,omitempty" xml:"orderIds,omitempty" require:"true" type:"Repeated"`
	// {"en":"Purchase Record ID List","zh_CN":"采购记录ID列表"}
	PurchaseRecordIds []*string `json:"purchaseRecordIds,omitempty" xml:"purchaseRecordIds,omitempty" require:"true" type:"Repeated"`
	// {"en":"Validate alias List","zh_CN":"验证别名列表"}
	ValidateAliasDomains []*CreateCertificateApplicationOrderForTerraformResponseDataValidateAliasDomains `json:"validateAliasDomains,omitempty" xml:"validateAliasDomains,omitempty" require:"true" type:"Repeated"`
	// {"en":"Domain and order relation info list","zh_CN":"域名与订单关联关系"}
	DomainOrderRelationInfos []*CreateCertificateApplicationOrderForTerraformResponseDataDomainOrderRelationInfos `json:"domainOrderRelationInfos,omitempty" xml:"domainOrderRelationInfos,omitempty" require:"true" type:"Repeated"`
}

func (s CreateCertificateApplicationOrderForTerraformResponseData) String() string {
	return tea.Prettify(s)
}

func (s CreateCertificateApplicationOrderForTerraformResponseData) GoString() string {
	return s.String()
}

func (s *CreateCertificateApplicationOrderForTerraformResponseData) SetOrderId(v string) *CreateCertificateApplicationOrderForTerraformResponseData {
	s.OrderId = &v
	return s
}

func (s *CreateCertificateApplicationOrderForTerraformResponseData) SetOrderIds(v []*string) *CreateCertificateApplicationOrderForTerraformResponseData {
	s.OrderIds = v
	return s
}

func (s *CreateCertificateApplicationOrderForTerraformResponseData) SetPurchaseRecordIds(v []*string) *CreateCertificateApplicationOrderForTerraformResponseData {
	s.PurchaseRecordIds = v
	return s
}

func (s *CreateCertificateApplicationOrderForTerraformResponseData) SetValidateAliasDomains(v []*CreateCertificateApplicationOrderForTerraformResponseDataValidateAliasDomains) *CreateCertificateApplicationOrderForTerraformResponseData {
	s.ValidateAliasDomains = v
	return s
}

func (s *CreateCertificateApplicationOrderForTerraformResponseData) SetDomainOrderRelationInfos(v []*CreateCertificateApplicationOrderForTerraformResponseDataDomainOrderRelationInfos) *CreateCertificateApplicationOrderForTerraformResponseData {
	s.DomainOrderRelationInfos = v
	return s
}

type CreateCertificateApplicationOrderForTerraformResponseDataValidateAliasDomains struct {
	// {"en":"domain","zh_CN":"域名"}
	Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
	// {"en":"Validate alias","zh_CN":"验证别名"}
	ValidateAliasDomain *string `json:"validateAliasDomain,omitempty" xml:"validateAliasDomain,omitempty" require:"true"`
}

func (s CreateCertificateApplicationOrderForTerraformResponseDataValidateAliasDomains) String() string {
	return tea.Prettify(s)
}

func (s CreateCertificateApplicationOrderForTerraformResponseDataValidateAliasDomains) GoString() string {
	return s.String()
}

func (s *CreateCertificateApplicationOrderForTerraformResponseDataValidateAliasDomains) SetDomain(v string) *CreateCertificateApplicationOrderForTerraformResponseDataValidateAliasDomains {
	s.Domain = &v
	return s
}

func (s *CreateCertificateApplicationOrderForTerraformResponseDataValidateAliasDomains) SetValidateAliasDomain(v string) *CreateCertificateApplicationOrderForTerraformResponseDataValidateAliasDomains {
	s.ValidateAliasDomain = &v
	return s
}

type CreateCertificateApplicationOrderForTerraformResponseDataDomainOrderRelationInfos struct {
	// {"en":"Domain name","zh_CN":"域名名称"}
	DomainName *string `json:"domainName,omitempty" xml:"domainName,omitempty" require:"true"`
	// {"en":"Order ID","zh_CN":"订单ID"}
	SalesOrderId *string `json:"salesOrderId,omitempty" xml:"salesOrderId,omitempty" require:"true"`
	// {"en":"Purchase Record ID","zh_CN":"采购记录ID"}
	PurchaseRecordId *string `json:"purchaseRecordId,omitempty" xml:"purchaseRecordId,omitempty" require:"true"`
}

func (s CreateCertificateApplicationOrderForTerraformResponseDataDomainOrderRelationInfos) String() string {
	return tea.Prettify(s)
}

func (s CreateCertificateApplicationOrderForTerraformResponseDataDomainOrderRelationInfos) GoString() string {
	return s.String()
}

func (s *CreateCertificateApplicationOrderForTerraformResponseDataDomainOrderRelationInfos) SetDomainName(v string) *CreateCertificateApplicationOrderForTerraformResponseDataDomainOrderRelationInfos {
	s.DomainName = &v
	return s
}

func (s *CreateCertificateApplicationOrderForTerraformResponseDataDomainOrderRelationInfos) SetSalesOrderId(v string) *CreateCertificateApplicationOrderForTerraformResponseDataDomainOrderRelationInfos {
	s.SalesOrderId = &v
	return s
}

func (s *CreateCertificateApplicationOrderForTerraformResponseDataDomainOrderRelationInfos) SetPurchaseRecordId(v string) *CreateCertificateApplicationOrderForTerraformResponseDataDomainOrderRelationInfos {
	s.PurchaseRecordId = &v
	return s
}

type CreateCertificateApplicationOrderForTerraformResponseHeader struct {
}

func (s CreateCertificateApplicationOrderForTerraformResponseHeader) String() string {
	return tea.Prettify(s)
}

func (s CreateCertificateApplicationOrderForTerraformResponseHeader) GoString() string {
	return s.String()
}

type CancelCertificateApplicationOrderForTerraformRequest struct {
	// {"en":"The orderId","zh_CN":"订单ID"}
	OrderId *string `json:"orderId,omitempty" xml:"orderId,omitempty"`
	// {"en":"The purchaseRecordId","zh_CN":"采购记录ID"}
	PurchaseRecordId *string `json:"purchaseRecordId,omitempty" xml:"purchaseRecordId,omitempty"`
}

func (s CancelCertificateApplicationOrderForTerraformRequest) String() string {
	return tea.Prettify(s)
}

func (s CancelCertificateApplicationOrderForTerraformRequest) GoString() string {
	return s.String()
}

func (s *CancelCertificateApplicationOrderForTerraformRequest) SetOrderId(v string) *CancelCertificateApplicationOrderForTerraformRequest {
	s.OrderId = &v
	return s
}

func (s *CancelCertificateApplicationOrderForTerraformRequest) SetPurchaseRecordId(v string) *CancelCertificateApplicationOrderForTerraformRequest {
	s.PurchaseRecordId = &v
	return s
}

type CancelCertificateApplicationOrderForTerraformRequestHeader struct {
}

func (s CancelCertificateApplicationOrderForTerraformRequestHeader) String() string {
	return tea.Prettify(s)
}

func (s CancelCertificateApplicationOrderForTerraformRequestHeader) GoString() string {
	return s.String()
}

type CancelCertificateApplicationOrderForTerraformPaths struct {
}

func (s CancelCertificateApplicationOrderForTerraformPaths) String() string {
	return tea.Prettify(s)
}

func (s CancelCertificateApplicationOrderForTerraformPaths) GoString() string {
	return s.String()
}

type CancelCertificateApplicationOrderForTerraformParameters struct {
}

func (s CancelCertificateApplicationOrderForTerraformParameters) String() string {
	return tea.Prettify(s)
}

func (s CancelCertificateApplicationOrderForTerraformParameters) GoString() string {
	return s.String()
}

type CancelCertificateApplicationOrderForTerraformResponse struct {
	// {"en":"Result Code, 0 means successful","zh_CN":"响应码，成功为0"}
	Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
	// {"en":"Result Message","zh_CN":"响应信息"}
	Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s CancelCertificateApplicationOrderForTerraformResponse) String() string {
	return tea.Prettify(s)
}

func (s CancelCertificateApplicationOrderForTerraformResponse) GoString() string {
	return s.String()
}

func (s *CancelCertificateApplicationOrderForTerraformResponse) SetCode(v string) *CancelCertificateApplicationOrderForTerraformResponse {
	s.Code = &v
	return s
}

func (s *CancelCertificateApplicationOrderForTerraformResponse) SetMessage(v string) *CancelCertificateApplicationOrderForTerraformResponse {
	s.Message = &v
	return s
}

type CancelCertificateApplicationOrderForTerraformResponseHeader struct {
}

func (s CancelCertificateApplicationOrderForTerraformResponseHeader) String() string {
	return tea.Prettify(s)
}

func (s CancelCertificateApplicationOrderForTerraformResponseHeader) GoString() string {
	return s.String()
}
