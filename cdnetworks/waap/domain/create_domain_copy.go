// This file is auto-generated, don't edit it. Thanks.
package domain

import (
	"github.com/alibabacloud-go/tea/tea"
)

type UsingExistingHostnameToAddNewHostnameRequest struct {
	// {"en":"The reference hostname.", "zh_CN":"指定域名。"}
	SourceDomain *string `json:"sourceDomain,omitempty" xml:"sourceDomain,omitempty" require:"true"`
	// {"en":"Hostname to be accessed.", "zh_CN":"目标域名。"}
	TargetDomains []*string `json:"targetDomains,omitempty" xml:"targetDomains,omitempty" require:"true" type:"Repeated"`
}

func (s UsingExistingHostnameToAddNewHostnameRequest) String() string {
	return tea.Prettify(s)
}

func (s UsingExistingHostnameToAddNewHostnameRequest) GoString() string {
	return s.String()
}

func (s *UsingExistingHostnameToAddNewHostnameRequest) SetSourceDomain(v string) *UsingExistingHostnameToAddNewHostnameRequest {
	s.SourceDomain = &v
	return s
}

func (s *UsingExistingHostnameToAddNewHostnameRequest) SetTargetDomains(v []*string) *UsingExistingHostnameToAddNewHostnameRequest {
	s.TargetDomains = v
	return s
}

type UsingExistingHostnameToAddNewHostnameResponse struct {
	// {'en':'Please refer to the error code for exceptions.', 'zh_CN':'请参照错误码。','dictionary':'belong=WAAP-MS-Ext|dict=waap_retCodeEnum'}
	Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
	// {"en":"Description.", "zh_CN":"描述信息。"}
	Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s UsingExistingHostnameToAddNewHostnameResponse) String() string {
	return tea.Prettify(s)
}

func (s UsingExistingHostnameToAddNewHostnameResponse) GoString() string {
	return s.String()
}

func (s *UsingExistingHostnameToAddNewHostnameResponse) SetCode(v string) *UsingExistingHostnameToAddNewHostnameResponse {
	s.Code = &v
	return s
}

func (s *UsingExistingHostnameToAddNewHostnameResponse) SetMessage(v string) *UsingExistingHostnameToAddNewHostnameResponse {
	s.Message = &v
	return s
}
