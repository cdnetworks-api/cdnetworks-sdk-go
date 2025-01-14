// This file is auto-generated, don't edit it. Thanks.
package domain

import (
	"github.com/alibabacloud-go/tea/tea"
)

type QueryPagingDomainListForTerraformRequest struct {
	// {"en":"Page number must be a positive integer greater than 0.If not passed, then no paging. If it is passed, pageSize is required.", "zh_CN":"分页的页码，必须为大于0的正整数。不传默认不分页，若传参则pageSize必填.。"}
	PageNumber *int `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// {"en":"Number of domain name data items for paging, must be a positive integer greater than 0.If not passed, then no paging. If it is passed, pageSize is required.", "zh_CN":"分页的域名数据条数，必须大于0的正整数。不传默认不分页，若传参则pageSize必填.。"}
	PageSize *int `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// {"en":"Specify the service type to be queried. Multiple services are allowed. Data will be returned if any one service is satisfied. If not passed, all services will be checked by default. For example: [wsa,waf], returns all domains whose services include wsa or include waf.", "zh_CN":"指定查询的服务，允许多个服务，任意一个服务满足就返回数据，不传默认查全部服务。如：[wsa,waf], 则返回服务包含wsa或包含waf的所有域名。"}
	ServiceTypes []*string `json:"serviceTypes,omitempty" xml:"serviceTypes,omitempty" type:"Repeated"`
	// {"en":"Specify the accelerated domain name for the query. Multiple domain names are allowed. If not specified, all domain names will be searched by default.", "zh_CN":"指定查询的加速域名，允许多个域名，不传默认查全部域名。"}
	DomainNames []*string `json:"domainNames,omitempty" xml:"domainNames,omitempty" type:"Repeated"`
	// {"en":"RFC3339 formatted date indicating the starting date. Example: 2024-01-01T22:30:00+08:00", "zh_CN":"	查询开始时间，支持时间格式如：2024-01-01T22:30:00+08:00"}
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// {"en":"RFC3339 formatted date indicating the ending date. Example: 2024-01-01T22:30:00+08:00", "zh_CN":"查询结束时间，支持时间格式如：2024-01-01T22:30:00+08:00"}
	EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// {"en":" Status of the accelerated domain. Optional value: enabled, disabled, deploying, checking, disabling, deployFailed, disableFailed.", "zh_CN":"加速域名的状态：enabled表示已启用；disabled表示已禁用；deploying表示配置部署中；checking表示审核中；disabling表示禁用中；deployFailed表示配置部署失败；disableFailed表示禁用失败。不传默认查全部"}
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s QueryPagingDomainListForTerraformRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryPagingDomainListForTerraformRequest) GoString() string {
	return s.String()
}

func (s *QueryPagingDomainListForTerraformRequest) SetPageNumber(v int) *QueryPagingDomainListForTerraformRequest {
	s.PageNumber = &v
	return s
}

func (s *QueryPagingDomainListForTerraformRequest) SetPageSize(v int) *QueryPagingDomainListForTerraformRequest {
	s.PageSize = &v
	return s
}

func (s *QueryPagingDomainListForTerraformRequest) SetServiceTypes(v []*string) *QueryPagingDomainListForTerraformRequest {
	s.ServiceTypes = v
	return s
}

func (s *QueryPagingDomainListForTerraformRequest) SetDomainNames(v []*string) *QueryPagingDomainListForTerraformRequest {
	s.DomainNames = v
	return s
}

func (s *QueryPagingDomainListForTerraformRequest) SetStartTime(v string) *QueryPagingDomainListForTerraformRequest {
	s.StartTime = &v
	return s
}

func (s *QueryPagingDomainListForTerraformRequest) SetEndTime(v string) *QueryPagingDomainListForTerraformRequest {
	s.EndTime = &v
	return s
}

func (s *QueryPagingDomainListForTerraformRequest) SetStatus(v string) *QueryPagingDomainListForTerraformRequest {
	s.Status = &v
	return s
}

type QueryPagingDomainListForTerraformResponse struct {
	// {"en":"Response code, 0 means successful.", "zh_CN":"接口响应code，0代表成功。"}
	Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
	// {"en":"Response error message if failed.", "zh_CN":"接口响应信息，success代表成功，失败则提供失败信息。"}
	Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
	// {"en":"Response data.", "zh_CN":"接口响应数据"}
	Data *QueryPagingDomainListForTerraformResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
}

func (s QueryPagingDomainListForTerraformResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryPagingDomainListForTerraformResponse) GoString() string {
	return s.String()
}

func (s *QueryPagingDomainListForTerraformResponse) SetCode(v string) *QueryPagingDomainListForTerraformResponse {
	s.Code = &v
	return s
}

func (s *QueryPagingDomainListForTerraformResponse) SetMessage(v string) *QueryPagingDomainListForTerraformResponse {
	s.Message = &v
	return s
}

func (s *QueryPagingDomainListForTerraformResponse) SetData(v *QueryPagingDomainListForTerraformResponseData) *QueryPagingDomainListForTerraformResponse {
	s.Data = v
	return s
}

type QueryPagingDomainListForTerraformResponseData struct {
	// {"en":"Responses the page number of the data", "zh_CN":"所有满足条件的数据总条数"}
	TotalCount *int `json:"totalCount,omitempty" xml:"totalCount,omitempty" require:"true"`
	// {"en":"total pages", "zh_CN":"总页数"}
	TotalPageNumber *int `json:"totalPageNumber,omitempty" xml:"totalPageNumber,omitempty" require:"true"`
	// {"en":"Responses the page number of the data", "zh_CN":"返回数据的页码"}
	PageNumber *int `json:"pageNumber,omitempty" xml:"pageNumber,omitempty" require:"true"`
	// {"en":"Number of data page", "zh_CN":"每个页面的数据条数"}
	PageSize *int `json:"pageSize,omitempty" xml:"pageSize,omitempty" require:"true"`
	// {"en":"Domain list.", "zh_CN":"域名列表"}
	ResultList []*QueryPagingDomainListForTerraformResponseDataResultList `json:"resultList,omitempty" xml:"resultList,omitempty" require:"true" type:"Repeated"`
}

func (s QueryPagingDomainListForTerraformResponseData) String() string {
	return tea.Prettify(s)
}

func (s QueryPagingDomainListForTerraformResponseData) GoString() string {
	return s.String()
}

func (s *QueryPagingDomainListForTerraformResponseData) SetTotalCount(v int) *QueryPagingDomainListForTerraformResponseData {
	s.TotalCount = &v
	return s
}

func (s *QueryPagingDomainListForTerraformResponseData) SetTotalPageNumber(v int) *QueryPagingDomainListForTerraformResponseData {
	s.TotalPageNumber = &v
	return s
}

func (s *QueryPagingDomainListForTerraformResponseData) SetPageNumber(v int) *QueryPagingDomainListForTerraformResponseData {
	s.PageNumber = &v
	return s
}

func (s *QueryPagingDomainListForTerraformResponseData) SetPageSize(v int) *QueryPagingDomainListForTerraformResponseData {
	s.PageSize = &v
	return s
}

func (s *QueryPagingDomainListForTerraformResponseData) SetResultList(v []*QueryPagingDomainListForTerraformResponseDataResultList) *QueryPagingDomainListForTerraformResponseData {
	s.ResultList = v
	return s
}

type QueryPagingDomainListForTerraformResponseDataResultList struct {
	// {"en":"Cname of the accelerated domain", "zh_CN":"加速域名cname，如：a1.example.com.wscdns.com"}
	Cname *string `json:"cname,omitempty" xml:"cname,omitempty" require:"true"`
	// {"en":"Create time of the accelerated domain. Example: 2024-01-01T22:30:00+08:00", "zh_CN":"	域名创建时间，时间格式如：2024-01-01T22:30:00+08:00"}
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty" require:"true"`
	// {"en":"Corresponding domain ID", "zh_CN":"对应的域名ID"}
	DomainId *string `json:"domainId,omitempty" xml:"domainId,omitempty" require:"true"`
	// {"en":"Accelerated domain name", "zh_CN":"加速域名名称"}
	DomainName *string `json:"domainName,omitempty" xml:"domainName,omitempty" require:"true"`
	// {"en":"Service type for accelerated domain name", "zh_CN":"加速域名的服务，如[wsa,waf]。"}
	ServiceTypes []*string `json:"serviceTypes,omitempty" xml:"serviceTypes,omitempty" require:"true" type:"Repeated"`
	// {"en":"Status of the accelerated domain. Optional value: enabled, disabled, deploying, checking, disabling.", "zh_CN":"	加速域名的状态：enabled表示已启用；disabled表示已禁用；deploying表示部署中；checking表示审核中；disabling表示禁用中；deployFailed表示配置部署失败；disableFailed表示禁用失败"}
	Status *string `json:"status,omitempty" xml:"status,omitempty" require:"true"`
	// {"en":"Accelerated domain enabling status: true indicates that it is enabled, false indicates that it is disabled.", "zh_CN":"加速域名启用状态：true为启用，false为禁用。"}
	Enabled *string `json:"enabled,omitempty" xml:"enabled,omitempty" require:"true"`
}

func (s QueryPagingDomainListForTerraformResponseDataResultList) String() string {
	return tea.Prettify(s)
}

func (s QueryPagingDomainListForTerraformResponseDataResultList) GoString() string {
	return s.String()
}

func (s *QueryPagingDomainListForTerraformResponseDataResultList) SetCname(v string) *QueryPagingDomainListForTerraformResponseDataResultList {
	s.Cname = &v
	return s
}

func (s *QueryPagingDomainListForTerraformResponseDataResultList) SetCreateTime(v string) *QueryPagingDomainListForTerraformResponseDataResultList {
	s.CreateTime = &v
	return s
}

func (s *QueryPagingDomainListForTerraformResponseDataResultList) SetDomainId(v string) *QueryPagingDomainListForTerraformResponseDataResultList {
	s.DomainId = &v
	return s
}

func (s *QueryPagingDomainListForTerraformResponseDataResultList) SetDomainName(v string) *QueryPagingDomainListForTerraformResponseDataResultList {
	s.DomainName = &v
	return s
}

func (s *QueryPagingDomainListForTerraformResponseDataResultList) SetServiceTypes(v []*string) *QueryPagingDomainListForTerraformResponseDataResultList {
	s.ServiceTypes = v
	return s
}

func (s *QueryPagingDomainListForTerraformResponseDataResultList) SetStatus(v string) *QueryPagingDomainListForTerraformResponseDataResultList {
	s.Status = &v
	return s
}

func (s *QueryPagingDomainListForTerraformResponseDataResultList) SetEnabled(v string) *QueryPagingDomainListForTerraformResponseDataResultList {
	s.Enabled = &v
	return s
}
