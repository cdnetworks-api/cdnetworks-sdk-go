// This file is auto-generated, don't edit it. Thanks.
package whitelist

import (
	"github.com/alibabacloud-go/tea/tea"
)

type DeleteWhitelistRulesRequest struct {
	// {"en":"Rule ID List.", "zh_CN":"规则ID列表。"}
	IdList []*string `json:"idList,omitempty" xml:"idList,omitempty" require:"true" type:"Repeated"`
}

func (s DeleteWhitelistRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteWhitelistRulesRequest) GoString() string {
	return s.String()
}

func (s *DeleteWhitelistRulesRequest) SetIdList(v []*string) *DeleteWhitelistRulesRequest {
	s.IdList = v
	return s
}

type DeleteWhitelistRulesResponse struct {
	// {'en':'Please refer to the error code for exceptions.', 'zh_CN':'请参照错误码。','dictionary':'belong=WAAP-MS-Ext|dict=waap_retCodeEnum'}
	Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
	// {"en":"Description.", "zh_CN":"描述信息。"}
	Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s DeleteWhitelistRulesResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteWhitelistRulesResponse) GoString() string {
	return s.String()
}

func (s *DeleteWhitelistRulesResponse) SetCode(v string) *DeleteWhitelistRulesResponse {
	s.Code = &v
	return s
}

func (s *DeleteWhitelistRulesResponse) SetMessage(v string) *DeleteWhitelistRulesResponse {
	s.Message = &v
	return s
}
