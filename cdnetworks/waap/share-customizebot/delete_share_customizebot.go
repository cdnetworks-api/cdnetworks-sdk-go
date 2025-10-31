// This file is auto-generated, don't edit it. Thanks.
package share_customizebot

import (
	"github.com/alibabacloud-go/tea/tea"
)

type DeleteShareCustomizeBotsRequest struct {
	// {"en":"Rule ID list.","zh_CN":"规则ID列表。"}
	IdList []*string `json:"idList,omitempty" xml:"idList,omitempty" require:"true" type:"Repeated"`
}

func (s DeleteShareCustomizeBotsRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteShareCustomizeBotsRequest) GoString() string {
	return s.String()
}

func (s *DeleteShareCustomizeBotsRequest) SetIdList(v []*string) *DeleteShareCustomizeBotsRequest {
	s.IdList = v
	return s
}

type DeleteShareCustomizeBotsResponse struct {
	// {"dictionary":"belong=WAAP-MS-Ext|dict=waap_retCodeEnum","en":"Please refer to the error code for exceptions.","zh_CN":"请参照错误码。"}
	Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
	// {"en":"Description.","zh_CN":"描述信息。"}
	Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
	// {"en":"Data.","zh_CN":"出参数据。"}
	Data *bool `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s DeleteShareCustomizeBotsResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteShareCustomizeBotsResponse) GoString() string {
	return s.String()
}

func (s *DeleteShareCustomizeBotsResponse) SetCode(v string) *DeleteShareCustomizeBotsResponse {
	s.Code = &v
	return s
}

func (s *DeleteShareCustomizeBotsResponse) SetMsg(v string) *DeleteShareCustomizeBotsResponse {
	s.Msg = &v
	return s
}

func (s *DeleteShareCustomizeBotsResponse) SetData(v bool) *DeleteShareCustomizeBotsResponse {
	s.Data = &v
	return s
}
