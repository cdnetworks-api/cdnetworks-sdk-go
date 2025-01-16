// This file is auto-generated, don't edit it. Thanks.
package customizerule

import (
	"github.com/alibabacloud-go/tea/tea"
)

type DeleteCustomRuleRequest struct {
	// {'en':'Rule ID List.', 'zh_CN':'规则ID列表。'}
	IdList []*string `json:"idList,omitempty" xml:"idList,omitempty" require:"true" type:"Repeated"`
}

func (s DeleteCustomRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteCustomRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteCustomRuleRequest) SetIdList(v []*string) *DeleteCustomRuleRequest {
	s.IdList = v
	return s
}

type DeleteCustomRuleResponse struct {
	// {'en':'Please refer to the error code for exceptions.', 'zh_CN':'请参照错误码。','dictionary':'belong=WAAP-MS-Ext|dict=waap_retCodeEnum'}
	Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
	// {'en':'Description.', 'zh_CN':'描述信息。'}
	Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s DeleteCustomRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteCustomRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteCustomRuleResponse) SetCode(v string) *DeleteCustomRuleResponse {
	s.Code = &v
	return s
}

func (s *DeleteCustomRuleResponse) SetMessage(v string) *DeleteCustomRuleResponse {
	s.Message = &v
	return s
}
