// This file is auto-generated, don't edit it. Thanks.
package share_customizerule

import (
	"github.com/alibabacloud-go/tea/tea"
)

type DeleteSharedCustomRulesRequest struct {
	// {'en':'Rule ID list.', 'zh_CN':'规则ID列表。'}
	IdList []*string `json:"idList,omitempty" xml:"idList,omitempty" require:"true" type:"Repeated"`
}

func (s DeleteSharedCustomRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteSharedCustomRulesRequest) GoString() string {
	return s.String()
}

func (s *DeleteSharedCustomRulesRequest) SetIdList(v []*string) *DeleteSharedCustomRulesRequest {
	s.IdList = v
	return s
}

type DeleteSharedCustomRulesResponse struct {
	// {'en':'Please refer to the error code for exceptions.', 'zh_CN':'请参照错误码。','dictionary':'belong=WAAP-MS-Ext|dict=waap_retCodeEnum'}
	Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
	// {'en':'Description.', 'zh_CN':'描述信息。'}
	Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
	// {'en':'Data.', 'zh_CN':'出参数据。'}
	Data *bool `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s DeleteSharedCustomRulesResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteSharedCustomRulesResponse) GoString() string {
	return s.String()
}

func (s *DeleteSharedCustomRulesResponse) SetCode(v string) *DeleteSharedCustomRulesResponse {
	s.Code = &v
	return s
}

func (s *DeleteSharedCustomRulesResponse) SetMsg(v string) *DeleteSharedCustomRulesResponse {
	s.Msg = &v
	return s
}

func (s *DeleteSharedCustomRulesResponse) SetData(v bool) *DeleteSharedCustomRulesResponse {
	s.Data = &v
	return s
}
