// This file is auto-generated, don't edit it. Thanks.
package certificate

import (
	"github.com/alibabacloud-go/tea/tea"
)

type QueryCertificateListForTerraformRequest struct {
	// {"en":"Response code, 0 means successful.", "zh_CN":"接口响应code，0代表成功。"}
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

type QueryCertificateListForTerraformResponse struct {
	// {"en":"Response code, 0 means successful.", "zh_CN":"接口响应code，0代表成功。"}
	Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
	// {"en":"Response error message if failed.", "zh_CN":"接口响应信息，success代表成功，失败则提供失败信息。"}
	Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
	// {"en":"Response data array.", "zh_CN":"接口响应数据"}
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

type QueryCertificateListForTerraformResponseData struct {
	// {"en":"Certificate ID", "zh_CN":"证书ID"}
	CertificateId *int64 `json:"certificateId,omitempty" xml:"certificateId,omitempty" require:"true"`
	// {"en":"Certificate name, unique to customer granularity", "zh_CN":"证书名称，客户粒度下是唯一的"}
	Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
	// {"en":"Remarks on cerfiticate file", "zh_CN":"证书文件的备注"}
	Comment *string `json:"comment,omitempty" xml:"comment,omitempty" require:"true"`
	// {"en":"Shared, optional values are true and false, true represents shared certificates, false represents unshared certificates, default is false
	// This certificate allows cross-customer use when share-ssl is true. (The API does not support cross-customer use certificates. Contact customer service for manual configuration if required.)", "zh_CN":"是否共享，true表示共享证书，false表示非共享证书"}
	ShareSsl *string `json:"shareSsl,omitempty" xml:"shareSsl,omitempty" require:"true"`
	// {"en":"Certificate effective start time (CST), such as 2016-08-01 07:00:00", "zh_CN":"证书有效期的起始时间（CST时区），例如：2016-08-01 07:00:00"}
	CertificateValidityFrom *string `json:"certificateValidityFrom,omitempty" xml:"certificateValidityFrom,omitempty" require:"true"`
	// {"en":"Certificate effective end time (CST), such as 2018-08-01 19:00:00", "zh_CN":"证书有效期的到期时间（CST时区），例如：2018-08-01 19:00:00"}
	CertificateValidityTo *string `json:"certificateValidityTo,omitempty" xml:"certificateValidityTo,omitempty" require:"true"`
	// {"en":"List of domain names using the current certificate", "zh_CN":"使用当前证书的域名列表"}
	RelatedDomains []*QueryCertificateListForTerraformResponseDataRelatedDomains `json:"relatedDomains,omitempty" xml:"relatedDomains,omitempty" require:"true" type:"Repeated"`
	// {"en":"dns-names", "zh_CN":"授权域名列表，证书使用者可选名称，父标签"}
	DnsNames []*string `json:"dnsNames,omitempty" xml:"dnsNames,omitempty" require:"true" type:"Repeated"`
	// {"en":"The CRT certificate serial number", "zh_CN":"crt证书序列号"}
	CertificateSerial *string `json:"certificateSerial,omitempty" xml:"certificateSerial,omitempty" require:"true"`
	// {"en":"The CRT certificate issuer", "zh_CN":"crt证书颁布者"}
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

type QueryCertificateListForTerraformResponseDataRelatedDomains struct {
	// {"en":"Accelerated domain name ID", "zh_CN":"加速域名ID"}
	DomainId *int64 `json:"domainId,omitempty" xml:"domainId,omitempty" require:"true"`
	// {"en":"Name of accelerated domain name", "zh_CN":"加速域名的名称"}
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
