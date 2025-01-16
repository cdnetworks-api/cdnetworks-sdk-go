// This file is auto-generated, don't edit it. Thanks.
package ratelimit

import (
	"github.com/alibabacloud-go/tea/tea"
)

type DeleteRateLimitingRulesRequest struct {
	// {'en':'Rule ID list.', 'zh_CN':'规则ID列表。'}
	Ids []*string `json:"ids,omitempty" xml:"ids,omitempty" require:"true" type:"Repeated"`
}

func (s DeleteRateLimitingRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteRateLimitingRulesRequest) GoString() string {
	return s.String()
}

func (s *DeleteRateLimitingRulesRequest) SetIds(v []*string) *DeleteRateLimitingRulesRequest {
	s.Ids = v
	return s
}

type DeleteRateLimitingRulesResponse struct {
	// {'en':'Please refer to the error code for exceptions.', 'zh_CN':'请参照错误码。','dictionary':'belong=WAAP-MS-Ext|dict=waap_retCodeEnum'}
	Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
	// {'en':'Description.', 'zh_CN':'描述信息。'}
	Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s DeleteRateLimitingRulesResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteRateLimitingRulesResponse) GoString() string {
	return s.String()
}

func (s *DeleteRateLimitingRulesResponse) SetCode(v string) *DeleteRateLimitingRulesResponse {
	s.Code = &v
	return s
}

func (s *DeleteRateLimitingRulesResponse) SetMessage(v string) *DeleteRateLimitingRulesResponse {
	s.Message = &v
	return s
}
