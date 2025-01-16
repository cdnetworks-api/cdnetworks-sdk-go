// This file is auto-generated, don't edit it. Thanks.
package domain

import (
	"github.com/alibabacloud-go/tea/tea"
)

type RemoveProtectedHostnameRequest struct {
}

func (s RemoveProtectedHostnameRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveProtectedHostnameRequest) GoString() string {
	return s.String()
}

type RemoveProtectedHostnameResponse struct {
	// {'en':'Please refer to the error code for exceptions.', 'zh_CN':'请参照错误码。','dictionary':'belong=WAAP-MS-Ext|dict=waap_retCodeEnum'}
	Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
	// {"en":"Description.", "zh_CN":"描述信息。"}
	Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s RemoveProtectedHostnameResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveProtectedHostnameResponse) GoString() string {
	return s.String()
}

func (s *RemoveProtectedHostnameResponse) SetCode(v string) *RemoveProtectedHostnameResponse {
	s.Code = &v
	return s
}

func (s *RemoveProtectedHostnameResponse) SetMessage(v string) *RemoveProtectedHostnameResponse {
	s.Message = &v
	return s
}

type RemoveProtectedHostnameParameters struct {
	// {"en":"Hostname.", "zh_CN":"域名。"}
	Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
}

func (s RemoveProtectedHostnameParameters) String() string {
	return tea.Prettify(s)
}

func (s RemoveProtectedHostnameParameters) GoString() string {
	return s.String()
}

func (s *RemoveProtectedHostnameParameters) SetDomain(v string) *RemoveProtectedHostnameParameters {
	s.Domain = &v
	return s
}
