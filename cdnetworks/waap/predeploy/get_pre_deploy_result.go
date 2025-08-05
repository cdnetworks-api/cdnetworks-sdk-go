// This file is auto-generated, don't edit it. Thanks.
package predeploy

import (
	"github.com/alibabacloud-go/tea/tea"
)

type GetPreDeployResultRequest struct {
	// {"en":"Pre-deployment id.","zh_CN":"预部署id。"}
	PreId *string `json:"preId,omitempty" xml:"preId,omitempty" require:"true"`
}

func (s GetPreDeployResultRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPreDeployResultRequest) GoString() string {
	return s.String()
}

func (s *GetPreDeployResultRequest) SetPreId(v string) *GetPreDeployResultRequest {
	s.PreId = &v
	return s
}

type GetPreDeployResultResponse struct {
	// {"dictionary":"belong=WAAP-MS-Ext|dict=waap_retCodeEnum","en":"Please refer to the error code for exceptions.","zh_CN":"请参照错误码。"}
	Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
	// {"en":"Description.","zh_CN":"描述信息。"}
	Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
	// {"en":"Data.","zh_CN":"出参数据。"}
	Data *GetPreDeployResultResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
}

func (s GetPreDeployResultResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPreDeployResultResponse) GoString() string {
	return s.String()
}

func (s *GetPreDeployResultResponse) SetCode(v string) *GetPreDeployResultResponse {
	s.Code = &v
	return s
}

func (s *GetPreDeployResultResponse) SetMsg(v string) *GetPreDeployResultResponse {
	s.Msg = &v
	return s
}

func (s *GetPreDeployResultResponse) SetData(v *GetPreDeployResultResponseData) *GetPreDeployResultResponse {
	s.Data = v
	return s
}

type GetPreDeployResultResponseData struct {
	// {"en":"Deployment results.\nDEPLOYING: Deployment in progress.\nSUCCESS: Deployment successful.\nFAIL: Deployment failed.","zh_CN":"部署结果。\nDEPLOYING：部署中。\nSUCCESS：部署成功。\nFAIL：部署失败。","exampleValue":"DEPLOYING,SUCCESS,FAIL"}
	DeployStatus *string `json:"deployStatus,omitempty" xml:"deployStatus,omitempty" require:"true"`
	// {"en":"Host list.","zh_CN":"Host 列表。"}
	HostList []*HostList `json:"hostList,omitempty" xml:"hostList,omitempty" require:"true" type:"Repeated"`
}

func (s GetPreDeployResultResponseData) String() string {
	return tea.Prettify(s)
}

func (s GetPreDeployResultResponseData) GoString() string {
	return s.String()
}

func (s *GetPreDeployResultResponseData) SetDeployStatus(v string) *GetPreDeployResultResponseData {
	s.DeployStatus = &v
	return s
}

func (s *GetPreDeployResultResponseData) SetHostList(v []*HostList) *GetPreDeployResultResponseData {
	s.HostList = v
	return s
}

type HostList struct {
	// {"en":"Domain.","zh_CN":"域名。"}
	HostName *string `json:"hostName,omitempty" xml:"hostName,omitempty" require:"true"`
	// {"en":"IP address.","zh_CN":"IP地址。"}
	HostAddress *string `json:"hostAddress,omitempty" xml:"hostAddress,omitempty" require:"true"`
}

func (s HostList) String() string {
	return tea.Prettify(s)
}

func (s HostList) GoString() string {
	return s.String()
}

func (s *HostList) SetHostName(v string) *HostList {
	s.HostName = &v
	return s
}

func (s *HostList) SetHostAddress(v string) *HostList {
	s.HostAddress = &v
	return s
}
