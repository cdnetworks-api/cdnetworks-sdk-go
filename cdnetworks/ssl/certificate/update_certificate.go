// This file is auto-generated, don't edit it. Thanks.
package certificate

import (
	"github.com/alibabacloud-go/tea/tea"
)

type UpdateCertificateForTerraformRequest struct {
	// {"en":"Certificate name", "zh_CN":"证书名称"}
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
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

type UpdateCertificateForTerraformResponse struct {
	// {"en":"Response code, 0 means successful.", "zh_CN":"接口响应code，0代表成功。"}
	Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
	// {"en":"Response error message if failed.", "zh_CN":"接口响应信息，success代表成功，失败则提供失败信息。"}
	Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
	// {"en":"Response data array.", "zh_CN":"接口响应数据"}
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
