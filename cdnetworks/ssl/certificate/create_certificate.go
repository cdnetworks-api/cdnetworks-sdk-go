// This file is auto-generated, don't edit it. Thanks.
package certificate

import (
	"github.com/alibabacloud-go/tea/tea"
)

type AddCertificateForTerraformRequest struct {
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

type AddCertificateForTerraformResponse struct {
	// {"en":"Status code", "zh_CN":"状态码"}
	Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
	// {"en":"Response message", "zh_CN":"响应信息"}
	Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
	// {"en":"Response data", "zh_CN":"响应数据"}
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
	// {"en":"Certificate ID", "zh_CN":"证书ID"}
	CertificateId *int64 `json:"certificateId,omitempty" xml:"certificateId,omitempty" require:"true"`
}

func (s AddCertificateForTerraformResponseData) String() string {
	return tea.Prettify(s)
}

func (s AddCertificateForTerraformResponseData) GoString() string {
	return s.String()
}

func (s *AddCertificateForTerraformResponseData) SetCertificateId(v int64) *AddCertificateForTerraformResponseData {
	s.CertificateId = &v
	return s
}
