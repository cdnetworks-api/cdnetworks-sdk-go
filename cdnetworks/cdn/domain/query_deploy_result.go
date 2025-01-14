package domain

import (
	"github.com/alibabacloud-go/tea/tea"
)

type QueryDeployResultForTerraformRequest struct {
}

func (s QueryDeployResultForTerraformRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryDeployResultForTerraformRequest) GoString() string {
	return s.String()
}

type QueryDeployResultForTerraformResponse struct {
	// {"en":"The error code, when HTTPStatus is not 200, indicates the type of error the current request is calling.", "zh_CN":"错误代码，当HTTPStatus不为200时出现，表示当前请求调用的错误类型"}
	Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
	// {"en":"Response information, when success is successful", "zh_CN":"响应信息，成功时为success"}
	Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
	// {"en":"Response data", "zh_CN":"响应数据"}
	Data *QueryDeployResultForTerraformResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
}

func (s QueryDeployResultForTerraformResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryDeployResultForTerraformResponse) GoString() string {
	return s.String()
}

func (s *QueryDeployResultForTerraformResponse) SetCode(v string) *QueryDeployResultForTerraformResponse {
	s.Code = &v
	return s
}

func (s *QueryDeployResultForTerraformResponse) SetMessage(v string) *QueryDeployResultForTerraformResponse {
	s.Message = &v
	return s
}

func (s *QueryDeployResultForTerraformResponse) SetData(v *QueryDeployResultForTerraformResponseData) *QueryDeployResultForTerraformResponse {
	s.Data = v
	return s
}

type QueryDeployResultForTerraformResponseData struct {
	// {"en":"The cnc-request-id corresponding to the request record you want to query this time", "zh_CN":"本次想要查询的请求记录对应的cnc-request-id"}
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty" require:"true"`
	// {"en":"Task execution results for asynchronous requests, including the following four states: WAIT: Indicates that the request is pending execution INPROGRESS: indicates that the request is in execution SUCCESS: Indicates that the request has been successfully executed FAIL: Indicates that the request execution failed Note: 1. If the query result is a failure of execution, there will be a mechanism or manual intervention in the backend of the system until the deployment is successful.If the result is the asynchronous request task of the new domain name, the system will have a resubmission mechanism or human intervention. 2. Modify the asynchronous request task for the domain name configuration, as well as add the domain name, there will be the mechanism of reintroduction and human intervention.", "zh_CN":"异步请求的任务执行结果，包括以下4种状态： WAIT：表示该请求等待执行 INPROGRESS：表示该请求执行中 SUCCESS：表示该请求已经执行成功 FAIL：表示该请求执行失败 注意： 1、新建域名的异步请求任务，通过本接口如果查询到结果是执行失败，系统后台会会有重提机制或者人工干预，直到部署成功。 2、修改域名配置的异步请求任务，和新增域名一样，也会有重提机制和人工干预。"}
	DeployResult *string `json:"deployResult,omitempty" xml:"deployResult,omitempty" require:"true"`
}

func (s QueryDeployResultForTerraformResponseData) String() string {
	return tea.Prettify(s)
}

func (s QueryDeployResultForTerraformResponseData) GoString() string {
	return s.String()
}

func (s *QueryDeployResultForTerraformResponseData) SetRequestId(v string) *QueryDeployResultForTerraformResponseData {
	s.RequestId = &v
	return s
}

func (s *QueryDeployResultForTerraformResponseData) SetDeployResult(v string) *QueryDeployResultForTerraformResponseData {
	s.DeployResult = &v
	return s
}
