package domain

import (
	"github.com/alibabacloud-go/tea/tea"
)

type DeleteDomainForTerraformRequest struct {
}

func (s DeleteDomainForTerraformRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDomainForTerraformRequest) GoString() string {
	return s.String()
}

type DeleteDomainForTerraformResponse struct {
	// {"en":"Response code, 0 means successful.", "zh_CN":"接口响应code，0代表成功。"}
	Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
	// {"en":"Response error message if failed.", "zh_CN":"接口响应信息，success代表成功，失败则提供失败信息。"}
	Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
	// {"en":"Response data.", "zh_CN":"接口响应数据"}
	Data *string `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s DeleteDomainForTerraformResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDomainForTerraformResponse) GoString() string {
	return s.String()
}

func (s *DeleteDomainForTerraformResponse) SetCode(v string) *DeleteDomainForTerraformResponse {
	s.Code = &v
	return s
}

func (s *DeleteDomainForTerraformResponse) SetMessage(v string) *DeleteDomainForTerraformResponse {
	s.Message = &v
	return s
}

func (s *DeleteDomainForTerraformResponse) SetData(v string) *DeleteDomainForTerraformResponse {
	s.Data = &v
	return s
}
