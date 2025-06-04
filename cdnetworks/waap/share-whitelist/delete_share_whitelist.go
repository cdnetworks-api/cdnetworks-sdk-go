// This file is auto-generated, don't edit it. Thanks.
package share_whitelist

import (
	"github.com/alibabacloud-go/tea/tea"
)

type DeleteShareWhitelistRuleRequest struct {
	// {"en":"Rule ID List.", "zh_CN":"规则ID列表。"}
	IdList []*string `json:"idList,omitempty" xml:"idList,omitempty" require:"true" type:"Repeated"`
}

func (s DeleteShareWhitelistRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteShareWhitelistRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteShareWhitelistRuleRequest) SetIdList(v []*string) *DeleteShareWhitelistRuleRequest {
	s.IdList = v
	return s
}

type DeleteShareWhitelistRuleResponse struct {
	// {'en':'Please refer to the error code for exceptions.', 'zh_CN':'请参照错误码。','dictionary':'belong=WAAP-MS-Ext|dict=waap_retCodeEnum'}
	Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
	// {"en":"Description.", "zh_CN":"描述信息。"}
	Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
}

func (s DeleteShareWhitelistRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteShareWhitelistRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteShareWhitelistRuleResponse) SetCode(v string) *DeleteShareWhitelistRuleResponse {
	s.Code = &v
	return s
}

func (s *DeleteShareWhitelistRuleResponse) SetMsg(v string) *DeleteShareWhitelistRuleResponse {
	s.Msg = &v
	return s
}
