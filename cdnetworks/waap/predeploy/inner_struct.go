package predeploy

import "github.com/alibabacloud-go/tea/tea"

type RequestHeader struct {
	// {"dictionary":"belong=WAAP-MS-Ext|dict=waap_serviceType","en":"Service type of the Acceleration domain:\n\n1. If not specified, it is considered as 'no restriction on service type.'\n\n2. For multiple Application server types, please separate them with an English semicolon \";\"","zh_CN":"加速域名的服务类型：\n\n1.未传递视为不限服务类型\n\n2.多个服务类型请使用英文分号\";\"分隔"}
	ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty"`
}

func (s RequestHeader) String() string {
	return tea.Prettify(s)
}

func (s RequestHeader) GoString() string {
	return s.String()
}

func (s *RequestHeader) SetServiceType(v string) *RequestHeader {
	s.ServiceType = &v
	return s
}

type Paths struct {
}

func (s Paths) String() string {
	return tea.Prettify(s)
}

func (s Paths) GoString() string {
	return s.String()
}

type Parameters struct {
}

func (s Parameters) String() string {
	return tea.Prettify(s)
}

func (s Parameters) GoString() string {
	return s.String()
}

type ResponseHeader struct {
}

func (s ResponseHeader) String() string {
	return tea.Prettify(s)
}

func (s ResponseHeader) GoString() string {
	return s.String()
}

type PreDeployResponseData struct {
	// {"en":"Pre-deployment id.","zh_CN":"预部署id。"}
	PreId *string `json:"preId,omitempty" xml:"preId,omitempty" require:"true"`
}

func (s PreDeployResponseData) String() string {
	return tea.Prettify(s)
}

func (s PreDeployResponseData) GoString() string {
	return s.String()
}

func (s *PreDeployResponseData) SetPreId(v string) *PreDeployResponseData {
	s.PreId = &v
	return s
}
