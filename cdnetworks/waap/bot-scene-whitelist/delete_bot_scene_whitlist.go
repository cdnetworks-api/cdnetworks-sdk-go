// This file is auto-generated, don't edit it. Thanks.
package bot_scene_whitelist

import (
	"github.com/alibabacloud-go/tea/tea"
)

type DeleteSpecificClientTrafficBypassRequest struct {
	// {"en":"Whitelist ID list.","zh_CN":"白名单ID列表。"}
	IdList []*string `json:"idList,omitempty" xml:"idList,omitempty" require:"true" type:"Repeated"`
}

func (s DeleteSpecificClientTrafficBypassRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteSpecificClientTrafficBypassRequest) GoString() string {
	return s.String()
}

func (s *DeleteSpecificClientTrafficBypassRequest) SetIdList(v []*string) *DeleteSpecificClientTrafficBypassRequest {
	s.IdList = v
	return s
}

type DeleteSpecificClientTrafficBypassResponse struct {
	// {"dictionary":"belong=WAAP-MS-Ext|dict=waap_retCodeEnum","en":"Please refer to the error code for exceptions.","zh_CN":"请参照错误码。"}
	Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
	// {"en":"Description.","zh_CN":"描述信息。"}
	Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
}

func (s DeleteSpecificClientTrafficBypassResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteSpecificClientTrafficBypassResponse) GoString() string {
	return s.String()
}

func (s *DeleteSpecificClientTrafficBypassResponse) SetCode(v string) *DeleteSpecificClientTrafficBypassResponse {
	s.Code = &v
	return s
}

func (s *DeleteSpecificClientTrafficBypassResponse) SetMsg(v string) *DeleteSpecificClientTrafficBypassResponse {
	s.Msg = &v
	return s
}
