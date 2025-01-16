// This file is auto-generated, don't edit it. Thanks.
package certificate

import (
	"github.com/alibabacloud-go/tea/tea"
)

type DeleteCertificateForTerraformRequest struct {
}

func (s DeleteCertificateForTerraformRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteCertificateForTerraformRequest) GoString() string {
	return s.String()
}

type DeleteCertificateForTerraformResponse struct {
	// {"en":"Response code, 0 means successful.", "zh_CN":"接口响应code，0代表成功。"}
	Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
	// {"en":"Response error message if failed.", "zh_CN":"接口响应信息，success代表成功，失败则提供失败信息。"}
	Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
	// {"en":"Response data array.", "zh_CN":"接口响应数据"}
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
