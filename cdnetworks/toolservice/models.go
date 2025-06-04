package toolservice

// This file is auto-generated, don't edit it. Thanks.
import (
  "github.com/alibabacloud-go/tea/tea"
)

type IcpQueryServiceRequest struct {
}

func (s IcpQueryServiceRequest) String() string {
  return tea.Prettify(s)
}

func (s IcpQueryServiceRequest) GoString() string {
  return s.String()
}

type IcpQueryServiceResponse struct {
  // {'en':'domainIcpData', 'zh_CN':'域名备案信息'}
  DomainIcpDataList []*IcpQueryServiceResponseDomainIcpDataList `json:"domain-icp-data,omitempty" xml:"domain-icp-data,omitempty" require:"true" type:"Repeated"`
}

func (s IcpQueryServiceResponse) String() string {
  return tea.Prettify(s)
}

func (s IcpQueryServiceResponse) GoString() string {
  return s.String()
}

func (s *IcpQueryServiceResponse) SetDomainIcpDataList(v []*IcpQueryServiceResponseDomainIcpDataList) *IcpQueryServiceResponse {
  s.DomainIcpDataList = v
  return s
}

type IcpQueryServiceResponseDomainIcpDataList struct     {
  // {"en":"Domain", "zh_CN":"域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {"en":"Registration No.", "zh_CN":"备案号"}
  IcpNumber *string `json:"icp-number,omitempty" xml:"icp-number,omitempty" require:"true"`
}

func (s IcpQueryServiceResponseDomainIcpDataList) String() string {
  return tea.Prettify(s)
}

func (s IcpQueryServiceResponseDomainIcpDataList) GoString() string {
  return s.String()
}

func (s *IcpQueryServiceResponseDomainIcpDataList) SetDomain(v string) *IcpQueryServiceResponseDomainIcpDataList {
  s.Domain = &v
  return s
}

func (s *IcpQueryServiceResponseDomainIcpDataList) SetIcpNumber(v string) *IcpQueryServiceResponseDomainIcpDataList {
  s.IcpNumber = &v
  return s
}

type IcpQueryServicePaths struct {
}

func (s IcpQueryServicePaths) String() string {
  return tea.Prettify(s)
}

func (s IcpQueryServicePaths) GoString() string {
  return s.String()
}

type IcpQueryServiceParameters struct {
  // {"en":"Domain names, multiple domain names are separated by English semicolons. The maximum number of domain names is 20.", "zh_CN":"域名，多个以英文分号分隔。域名数上限为20个。"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
}

func (s IcpQueryServiceParameters) String() string {
  return tea.Prettify(s)
}

func (s IcpQueryServiceParameters) GoString() string {
  return s.String()
}

func (s *IcpQueryServiceParameters) SetDomain(v string) *IcpQueryServiceParameters {
  s.Domain = &v
  return s
}

type IcpQueryServiceRequestHeader struct {
}

func (s IcpQueryServiceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s IcpQueryServiceRequestHeader) GoString() string {
  return s.String()
}

type IcpQueryServiceResponseHeader struct {
}

func (s IcpQueryServiceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s IcpQueryServiceResponseHeader) GoString() string {
  return s.String()
}




type QueryForbiddingVisitorIPsByLabelCodeServiceRequest struct {
  // {"en":"List of forbidding Label Code", "zh_CN":"封禁标签列表"}
  LabelCodeList []*string `json:"labelCodeList,omitempty" xml:"labelCodeList,omitempty" require:"true" type:"Repeated"`
  // {"en":"List of forbidding IP, leave it empty to query all ", "zh_CN":"封禁IP列表,放空则查询全部"}
  IpList []*string `json:"ipList,omitempty" xml:"ipList,omitempty" type:"Repeated"`
  // {"en":"Current page number,the first page starts from 0,default 0 ", "zh_CN":"分页，当前页，第一页从0开始，默认0"}
  PageNo *int `json:"pageNo,omitempty" xml:"pageNo,omitempty"`
  // {"en":"Page size,must be greater than 0,default 100 ", "zh_CN":"每页大小，必须大于0，默认100"}
  PageSize *int `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s QueryForbiddingVisitorIPsByLabelCodeServiceRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryForbiddingVisitorIPsByLabelCodeServiceRequest) GoString() string {
  return s.String()
}

func (s *QueryForbiddingVisitorIPsByLabelCodeServiceRequest) SetLabelCodeList(v []*string) *QueryForbiddingVisitorIPsByLabelCodeServiceRequest {
  s.LabelCodeList = v
  return s
}

func (s *QueryForbiddingVisitorIPsByLabelCodeServiceRequest) SetIpList(v []*string) *QueryForbiddingVisitorIPsByLabelCodeServiceRequest {
  s.IpList = v
  return s
}

func (s *QueryForbiddingVisitorIPsByLabelCodeServiceRequest) SetPageNo(v int) *QueryForbiddingVisitorIPsByLabelCodeServiceRequest {
  s.PageNo = &v
  return s
}

func (s *QueryForbiddingVisitorIPsByLabelCodeServiceRequest) SetPageSize(v int) *QueryForbiddingVisitorIPsByLabelCodeServiceRequest {
  s.PageSize = &v
  return s
}

type QueryForbiddingVisitorIPsByLabelCodeServiceResponse struct {
  // {"en":"Result Code", "zh_CN":"响应码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Result Message", "zh_CN":"响应信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Result Data", "zh_CN":"响应数据"}
  Data *QueryForbiddingVisitorIPsByLabelCodeServiceLabelIpQueryResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s QueryForbiddingVisitorIPsByLabelCodeServiceResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryForbiddingVisitorIPsByLabelCodeServiceResponse) GoString() string {
  return s.String()
}

func (s *QueryForbiddingVisitorIPsByLabelCodeServiceResponse) SetCode(v string) *QueryForbiddingVisitorIPsByLabelCodeServiceResponse {
  s.Code = &v
  return s
}

func (s *QueryForbiddingVisitorIPsByLabelCodeServiceResponse) SetMessage(v string) *QueryForbiddingVisitorIPsByLabelCodeServiceResponse {
  s.Message = &v
  return s
}

func (s *QueryForbiddingVisitorIPsByLabelCodeServiceResponse) SetData(v *QueryForbiddingVisitorIPsByLabelCodeServiceLabelIpQueryResponseData) *QueryForbiddingVisitorIPsByLabelCodeServiceResponse {
  s.Data = v
  return s
}

type QueryForbiddingVisitorIPsByLabelCodeServiceLabelIpQueryResponseData struct {
  // {"en":"Total count ", "zh_CN":"总数据条数"}
  Total *int `json:"total,omitempty" xml:"total,omitempty" require:"true"`
  // {"en":"Current page number,the first page starts from 0,default 0 ", "zh_CN":"分页，当前页，第一页从0开始，默认0"}
  PageNo *int `json:"pageNo,omitempty" xml:"pageNo,omitempty" require:"true"`
  // {"en":"Page size,must be greater than 0,default 100 ", "zh_CN":"每页大小，必须大于0，默认100"}
  PageSize *int `json:"pageSize,omitempty" xml:"pageSize,omitempty" require:"true"`
  // {"en":"Query results", "zh_CN":"查询结果"}
  Result []*QueryForbiddingVisitorIPsByLabelCodeServiceLabelIpQueryDetailData `json:"result,omitempty" xml:"result,omitempty" require:"true" type:"Repeated"`
}

func (s QueryForbiddingVisitorIPsByLabelCodeServiceLabelIpQueryResponseData) String() string {
  return tea.Prettify(s)
}

func (s QueryForbiddingVisitorIPsByLabelCodeServiceLabelIpQueryResponseData) GoString() string {
  return s.String()
}

func (s *QueryForbiddingVisitorIPsByLabelCodeServiceLabelIpQueryResponseData) SetTotal(v int) *QueryForbiddingVisitorIPsByLabelCodeServiceLabelIpQueryResponseData {
  s.Total = &v
  return s
}

func (s *QueryForbiddingVisitorIPsByLabelCodeServiceLabelIpQueryResponseData) SetPageNo(v int) *QueryForbiddingVisitorIPsByLabelCodeServiceLabelIpQueryResponseData {
  s.PageNo = &v
  return s
}

func (s *QueryForbiddingVisitorIPsByLabelCodeServiceLabelIpQueryResponseData) SetPageSize(v int) *QueryForbiddingVisitorIPsByLabelCodeServiceLabelIpQueryResponseData {
  s.PageSize = &v
  return s
}

func (s *QueryForbiddingVisitorIPsByLabelCodeServiceLabelIpQueryResponseData) SetResult(v []*QueryForbiddingVisitorIPsByLabelCodeServiceLabelIpQueryDetailData) *QueryForbiddingVisitorIPsByLabelCodeServiceLabelIpQueryResponseData {
  s.Result = v
  return s
}

type QueryForbiddingVisitorIPsByLabelCodeServiceLabelIpQueryDetailData struct {
  // {"en":"Label Code", "zh_CN":"标签编码"}
  LabelCode *string `json:"labelCode,omitempty" xml:"labelCode,omitempty" require:"true"`
  // {"en":"Label Name", "zh_CN":"标签名称"}
  LabelName *string `json:"labelName,omitempty" xml:"labelName,omitempty" require:"true"`
  // {"en":"IP forbidding", "zh_CN":"封禁的IP"}
  Ip *string `json:"ip,omitempty" xml:"ip,omitempty" require:"true"`
  // {"en":"Start time of forbidden", "zh_CN":"封禁开始时间"}
  StartTime *int `json:"startTime,omitempty" xml:"startTime,omitempty" require:"true"`
  // {"en":"End time of forbidden", "zh_CN":"封禁结束时间"}
  EndTime *int `json:"endTime,omitempty" xml:"endTime,omitempty" require:"true"`
}

func (s QueryForbiddingVisitorIPsByLabelCodeServiceLabelIpQueryDetailData) String() string {
  return tea.Prettify(s)
}

func (s QueryForbiddingVisitorIPsByLabelCodeServiceLabelIpQueryDetailData) GoString() string {
  return s.String()
}

func (s *QueryForbiddingVisitorIPsByLabelCodeServiceLabelIpQueryDetailData) SetLabelCode(v string) *QueryForbiddingVisitorIPsByLabelCodeServiceLabelIpQueryDetailData {
  s.LabelCode = &v
  return s
}

func (s *QueryForbiddingVisitorIPsByLabelCodeServiceLabelIpQueryDetailData) SetLabelName(v string) *QueryForbiddingVisitorIPsByLabelCodeServiceLabelIpQueryDetailData {
  s.LabelName = &v
  return s
}

func (s *QueryForbiddingVisitorIPsByLabelCodeServiceLabelIpQueryDetailData) SetIp(v string) *QueryForbiddingVisitorIPsByLabelCodeServiceLabelIpQueryDetailData {
  s.Ip = &v
  return s
}

func (s *QueryForbiddingVisitorIPsByLabelCodeServiceLabelIpQueryDetailData) SetStartTime(v int) *QueryForbiddingVisitorIPsByLabelCodeServiceLabelIpQueryDetailData {
  s.StartTime = &v
  return s
}

func (s *QueryForbiddingVisitorIPsByLabelCodeServiceLabelIpQueryDetailData) SetEndTime(v int) *QueryForbiddingVisitorIPsByLabelCodeServiceLabelIpQueryDetailData {
  s.EndTime = &v
  return s
}

type QueryForbiddingVisitorIPsByLabelCodeServicePaths struct {
}

func (s QueryForbiddingVisitorIPsByLabelCodeServicePaths) String() string {
  return tea.Prettify(s)
}

func (s QueryForbiddingVisitorIPsByLabelCodeServicePaths) GoString() string {
  return s.String()
}

type QueryForbiddingVisitorIPsByLabelCodeServiceParameters struct {
}

func (s QueryForbiddingVisitorIPsByLabelCodeServiceParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryForbiddingVisitorIPsByLabelCodeServiceParameters) GoString() string {
  return s.String()
}

type QueryForbiddingVisitorIPsByLabelCodeServiceRequestHeader struct {
}

func (s QueryForbiddingVisitorIPsByLabelCodeServiceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryForbiddingVisitorIPsByLabelCodeServiceRequestHeader) GoString() string {
  return s.String()
}

type QueryForbiddingVisitorIPsByLabelCodeServiceResponseHeader struct {
}

func (s QueryForbiddingVisitorIPsByLabelCodeServiceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryForbiddingVisitorIPsByLabelCodeServiceResponseHeader) GoString() string {
  return s.String()
}




type IpDomainServiceRequest struct {
  // {"en":"IP", "zh_CN":"IP"}
  Ip []*string `json:"ip,omitempty" xml:"ip,omitempty" require:"true" type:"Repeated"`
}

func (s IpDomainServiceRequest) String() string {
  return tea.Prettify(s)
}

func (s IpDomainServiceRequest) GoString() string {
  return s.String()
}

func (s *IpDomainServiceRequest) SetIp(v []*string) *IpDomainServiceRequest {
  s.Ip = v
  return s
}

type IpDomainServiceResponse struct {
  // {"en":"request result status code", "zh_CN":"请求结果状态码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Request result information", "zh_CN":"请求结果信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Detailed data on the results of the request", "zh_CN":"请求结果的详细数据"}
  Data []*IpDomainServiceResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s IpDomainServiceResponse) String() string {
  return tea.Prettify(s)
}

func (s IpDomainServiceResponse) GoString() string {
  return s.String()
}

func (s *IpDomainServiceResponse) SetCode(v string) *IpDomainServiceResponse {
  s.Code = &v
  return s
}

func (s *IpDomainServiceResponse) SetMessage(v string) *IpDomainServiceResponse {
  s.Message = &v
  return s
}

func (s *IpDomainServiceResponse) SetData(v []*IpDomainServiceResponseData) *IpDomainServiceResponse {
  s.Data = v
  return s
}

type IpDomainServiceResponseData struct     {
  // {"en":"ip", "zh_CN":"IP名称"}
  Ip *string `json:"ip,omitempty" xml:"ip,omitempty" require:"true"`
  // {"en":"Whether to use:
  // 
  //   idle --IP not used yet;
  //   runing -- IP in use;
  //   out of range -- IP is not in a queryable range", "zh_CN":"是否使用:
  //   idle -- 暂未使用;
  //   runing -- 使用中;
  //   out of range -- 不在查询范围内的ip"}
  Status *string `json:"status,omitempty" xml:"status,omitempty" require:"true"`
  // {"en":"List of domain using this IP.The domain list of the IP that was idle or out of range was empty", "zh_CN":"用该IP的域名列表,未使用的ip/不在查询范围内的ip,域名列表为空"}
  DomainList []*string `json:"domainList,omitempty" xml:"domainList,omitempty" require:"true" type:"Repeated"`
}

func (s IpDomainServiceResponseData) String() string {
  return tea.Prettify(s)
}

func (s IpDomainServiceResponseData) GoString() string {
  return s.String()
}

func (s *IpDomainServiceResponseData) SetIp(v string) *IpDomainServiceResponseData {
  s.Ip = &v
  return s
}

func (s *IpDomainServiceResponseData) SetStatus(v string) *IpDomainServiceResponseData {
  s.Status = &v
  return s
}

func (s *IpDomainServiceResponseData) SetDomainList(v []*string) *IpDomainServiceResponseData {
  s.DomainList = v
  return s
}

type IpDomainServicePaths struct {
}

func (s IpDomainServicePaths) String() string {
  return tea.Prettify(s)
}

func (s IpDomainServicePaths) GoString() string {
  return s.String()
}

type IpDomainServiceParameters struct {
}

func (s IpDomainServiceParameters) String() string {
  return tea.Prettify(s)
}

func (s IpDomainServiceParameters) GoString() string {
  return s.String()
}

type IpDomainServiceRequestHeader struct {
}

func (s IpDomainServiceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s IpDomainServiceRequestHeader) GoString() string {
  return s.String()
}

type IpDomainServiceResponseHeader struct {
}

func (s IpDomainServiceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s IpDomainServiceResponseHeader) GoString() string {
  return s.String()
}




type QueryForbiddingVisitorIPsByDomainServiceRequest struct {
  // {"en":"List of forbidding domain", "zh_CN":"封禁域名列表"}
  DomainList []*string `json:"domainList,omitempty" xml:"domainList,omitempty" require:"true" type:"Repeated"`
  // {"en":"List of forbidding IP", "zh_CN":"封禁IP列表"}
  IpList []*string `json:"ipList,omitempty" xml:"ipList,omitempty" type:"Repeated"`
  // {"en":"Current page number,the first page starts from 0,default 0", "zh_CN":"分页，当前页，第一页从0开始，默认0"}
  PageNo *int `json:"pageNo,omitempty" xml:"pageNo,omitempty"`
  // {"en":"Page size,must be greater than 0,default 100, the maximum is 1000 ", "zh_CN":"每页大小，必须大于0，默认100，最大1000"}
  PageSize *int `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s QueryForbiddingVisitorIPsByDomainServiceRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryForbiddingVisitorIPsByDomainServiceRequest) GoString() string {
  return s.String()
}

func (s *QueryForbiddingVisitorIPsByDomainServiceRequest) SetDomainList(v []*string) *QueryForbiddingVisitorIPsByDomainServiceRequest {
  s.DomainList = v
  return s
}

func (s *QueryForbiddingVisitorIPsByDomainServiceRequest) SetIpList(v []*string) *QueryForbiddingVisitorIPsByDomainServiceRequest {
  s.IpList = v
  return s
}

func (s *QueryForbiddingVisitorIPsByDomainServiceRequest) SetPageNo(v int) *QueryForbiddingVisitorIPsByDomainServiceRequest {
  s.PageNo = &v
  return s
}

func (s *QueryForbiddingVisitorIPsByDomainServiceRequest) SetPageSize(v int) *QueryForbiddingVisitorIPsByDomainServiceRequest {
  s.PageSize = &v
  return s
}

type QueryForbiddingVisitorIPsByDomainServiceResponse struct {
  // {"en":"Result Code", "zh_CN":"响应码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Result Message", "zh_CN":"响应信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Result Data", "zh_CN":"响应数据"}
  Data map[string]interface{} `json:"data,omitempty" xml:"data,omitempty" require:"true"`
  // {"en":"Total count", "zh_CN":"总数据条数"}
  Total *int `json:"total,omitempty" xml:"total,omitempty" require:"true"`
  // {"en":"Current page number,the first page starts from 0,default 0", "zh_CN":"分页，当前页，第一页从0开始，默认0"}
  PageNo *int `json:"pageNo,omitempty" xml:"pageNo,omitempty" require:"true"`
  // {"en":"Page size,must be greater than 0,default 100, the maximum is 1000", "zh_CN":"每页大小，必须大于0，默认100，最大1000"}
  PageSize *int `json:"pageSize,omitempty" xml:"pageSize,omitempty" require:"true"`
  // {"en":"Query results", "zh_CN":"查询结果"}
  Result []*string `json:"result,omitempty" xml:"result,omitempty" require:"true" type:"Repeated"`
  // {"en":"Domain forbidding", "zh_CN":"封禁的域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {"en":"IP forbidding", "zh_CN":"封禁的IP"}
  Ip *string `json:"ip,omitempty" xml:"ip,omitempty" require:"true"`
  // {"en":"Start time of forbidden", "zh_CN":"封禁开始时间"}
  StartTime *int `json:"startTime,omitempty" xml:"startTime,omitempty" require:"true"`
  // {"en":"End time of forbidden", "zh_CN":"封禁结束时间"}
  EndTime *int `json:"endTime,omitempty" xml:"endTime,omitempty" require:"true"`
}

func (s QueryForbiddingVisitorIPsByDomainServiceResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryForbiddingVisitorIPsByDomainServiceResponse) GoString() string {
  return s.String()
}

func (s *QueryForbiddingVisitorIPsByDomainServiceResponse) SetCode(v string) *QueryForbiddingVisitorIPsByDomainServiceResponse {
  s.Code = &v
  return s
}

func (s *QueryForbiddingVisitorIPsByDomainServiceResponse) SetMessage(v string) *QueryForbiddingVisitorIPsByDomainServiceResponse {
  s.Message = &v
  return s
}

func (s *QueryForbiddingVisitorIPsByDomainServiceResponse) SetData(v map[string]interface{}) *QueryForbiddingVisitorIPsByDomainServiceResponse {
  s.Data = v
  return s
}

func (s *QueryForbiddingVisitorIPsByDomainServiceResponse) SetTotal(v int) *QueryForbiddingVisitorIPsByDomainServiceResponse {
  s.Total = &v
  return s
}

func (s *QueryForbiddingVisitorIPsByDomainServiceResponse) SetPageNo(v int) *QueryForbiddingVisitorIPsByDomainServiceResponse {
  s.PageNo = &v
  return s
}

func (s *QueryForbiddingVisitorIPsByDomainServiceResponse) SetPageSize(v int) *QueryForbiddingVisitorIPsByDomainServiceResponse {
  s.PageSize = &v
  return s
}

func (s *QueryForbiddingVisitorIPsByDomainServiceResponse) SetResult(v []*string) *QueryForbiddingVisitorIPsByDomainServiceResponse {
  s.Result = v
  return s
}

func (s *QueryForbiddingVisitorIPsByDomainServiceResponse) SetDomain(v string) *QueryForbiddingVisitorIPsByDomainServiceResponse {
  s.Domain = &v
  return s
}

func (s *QueryForbiddingVisitorIPsByDomainServiceResponse) SetIp(v string) *QueryForbiddingVisitorIPsByDomainServiceResponse {
  s.Ip = &v
  return s
}

func (s *QueryForbiddingVisitorIPsByDomainServiceResponse) SetStartTime(v int) *QueryForbiddingVisitorIPsByDomainServiceResponse {
  s.StartTime = &v
  return s
}

func (s *QueryForbiddingVisitorIPsByDomainServiceResponse) SetEndTime(v int) *QueryForbiddingVisitorIPsByDomainServiceResponse {
  s.EndTime = &v
  return s
}

type QueryForbiddingVisitorIPsByDomainServicePaths struct {
}

func (s QueryForbiddingVisitorIPsByDomainServicePaths) String() string {
  return tea.Prettify(s)
}

func (s QueryForbiddingVisitorIPsByDomainServicePaths) GoString() string {
  return s.String()
}

type QueryForbiddingVisitorIPsByDomainServiceParameters struct {
}

func (s QueryForbiddingVisitorIPsByDomainServiceParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryForbiddingVisitorIPsByDomainServiceParameters) GoString() string {
  return s.String()
}

type QueryForbiddingVisitorIPsByDomainServiceRequestHeader struct {
}

func (s QueryForbiddingVisitorIPsByDomainServiceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryForbiddingVisitorIPsByDomainServiceRequestHeader) GoString() string {
  return s.String()
}

type QueryForbiddingVisitorIPsByDomainServiceResponseHeader struct {
}

func (s QueryForbiddingVisitorIPsByDomainServiceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryForbiddingVisitorIPsByDomainServiceResponseHeader) GoString() string {
  return s.String()
}




