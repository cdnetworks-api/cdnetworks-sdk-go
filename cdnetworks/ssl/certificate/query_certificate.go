// This file is auto-generated, don't edit it. Thanks.
package certificate

import (
	"github.com/alibabacloud-go/tea/tea"
)

type QueryCertificateForTerraformRequest struct {
}

func (s QueryCertificateForTerraformRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryCertificateForTerraformRequest) GoString() string {
	return s.String()
}

type QueryCertificateForTerraformResponse struct {
	// {"en":"Response code, 0 means successful.", "zh_CN":"接口响应code，0代表成功。"}
	Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
	// {"en":"Response error message if failed.", "zh_CN":"接口响应信息，success代表成功，失败则提供失败信息。"}
	Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
	// {"en":"Response data array.", "zh_CN":"接口响应数据"}
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
	// {"en":"certificate Id", "zh_CN":"证书ID。"}
	CertificateId *int64 `json:"certificateId,omitempty" xml:"certificateId,omitempty" require:"true"`
	// {"en":"certificate name", "zh_CN":"证书名称。"}
	Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
	// {"en":"Certificate, PEM certificate.", "zh_CN":"证书内容，PEM格式。"}
	Certificate *string `json:"certificate,omitempty" xml:"certificate,omitempty" require:"true"`
	// {"en":"Private key of the certificate, PEM certificate. Not required when you specify a csrId.", "zh_CN":"证书密钥，PEM格式。"}
	PrivateKey *string `json:"privateKey,omitempty" xml:"privateKey,omitempty" require:"true"`
	// {"en":"Certificate related domains", "zh_CN":"在用该证书的加速域名列表"}
	Domains []*QueryCertificateForTerraformResponseDataDomains `json:"domains,omitempty" xml:"domains,omitempty" require:"true" type:"Repeated"`
}

func (s QueryCertificateForTerraformResponseData) String() string {
	return tea.Prettify(s)
}

func (s QueryCertificateForTerraformResponseData) GoString() string {
	return s.String()
}

func (s *QueryCertificateForTerraformResponseData) SetCertificateId(v int64) *QueryCertificateForTerraformResponseData {
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

type QueryCertificateForTerraformResponseDataDomains struct {
	// {"en":"Domain ID", "zh_CN":"域名id"}
	DomainId *int64 `json:"domainId,omitempty" xml:"domainId,omitempty" require:"true"`
	// {"en":"Domain name", "zh_CN":"域名名称"}
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
