package ipforbid

// This file is auto-generated, don't edit it. Thanks.
import (
  "github.com/alibabacloud-go/tea/tea"
)

type QueryForbiddingVisitorIPsByLabelCodeServiceRequest struct {
  // {"en":"List of forbidding Label Code","zh_CN":"封禁标签列表"}
  LabelCodeList []*string `json:"labelCodeList,omitempty" xml:"labelCodeList,omitempty" require:"true" type:"Repeated"`
  // {"en":"List of forbidding IP, leave it empty to query all","zh_CN":"封禁IP列表,放空则查询全部"}
  IpList []*string `json:"ipList,omitempty" xml:"ipList,omitempty" type:"Repeated"`
  // {"en":"Current page number,the first page starts from 0,default 0","zh_CN":"分页，当前页，第一页从0开始，默认0"}
  PageNo *QueryForbiddingVisitorIPsByLabelCodeServiceRequestPageNo `json:"pageNo,omitempty" xml:"pageNo,omitempty" type:"Struct"`
  // {"en":"Page size,must be greater than 0,default 100","zh_CN":"每页大小，必须大于0，默认100"}
  PageSize *QueryForbiddingVisitorIPsByLabelCodeServiceRequestPageSize `json:"pageSize,omitempty" xml:"pageSize,omitempty" type:"Struct"`
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

func (s *QueryForbiddingVisitorIPsByLabelCodeServiceRequest) SetPageNo(v *QueryForbiddingVisitorIPsByLabelCodeServiceRequestPageNo) *QueryForbiddingVisitorIPsByLabelCodeServiceRequest {
  s.PageNo = v
  return s
}

func (s *QueryForbiddingVisitorIPsByLabelCodeServiceRequest) SetPageSize(v *QueryForbiddingVisitorIPsByLabelCodeServiceRequestPageSize) *QueryForbiddingVisitorIPsByLabelCodeServiceRequest {
  s.PageSize = v
  return s
}

type QueryForbiddingVisitorIPsByLabelCodeServiceRequestPageNo struct {
}

func (s QueryForbiddingVisitorIPsByLabelCodeServiceRequestPageNo) String() string {
  return tea.Prettify(s)
}

func (s QueryForbiddingVisitorIPsByLabelCodeServiceRequestPageNo) GoString() string {
  return s.String()
}

type QueryForbiddingVisitorIPsByLabelCodeServiceRequestPageSize struct {
}

func (s QueryForbiddingVisitorIPsByLabelCodeServiceRequestPageSize) String() string {
  return tea.Prettify(s)
}

func (s QueryForbiddingVisitorIPsByLabelCodeServiceRequestPageSize) GoString() string {
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

type QueryForbiddingVisitorIPsByLabelCodeServiceResponse struct {
  // {"en":"Result Code","zh_CN":"响应码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Result Message","zh_CN":"响应信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Result Data","zh_CN":"响应数据"}
  Data *QueryForbiddingVisitorIPsByLabelCodeServiceResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
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

func (s *QueryForbiddingVisitorIPsByLabelCodeServiceResponse) SetData(v *QueryForbiddingVisitorIPsByLabelCodeServiceResponseData) *QueryForbiddingVisitorIPsByLabelCodeServiceResponse {
  s.Data = v
  return s
}

type QueryForbiddingVisitorIPsByLabelCodeServiceResponseData struct {
  // {"en":"Total count","zh_CN":"总数据条数"}
  Total *QueryForbiddingVisitorIPsByLabelCodeServiceResponseDataTotal `json:"total,omitempty" xml:"total,omitempty" require:"true" type:"Struct"`
  // {"en":"Current page number,the first page starts from 0,default 0","zh_CN":"分页，当前页，第一页从0开始，默认0"}
  PageNo *QueryForbiddingVisitorIPsByLabelCodeServiceResponseDataPageNo `json:"pageNo,omitempty" xml:"pageNo,omitempty" require:"true" type:"Struct"`
  // {"en":"Page size,must be greater than 0,default 100","zh_CN":"每页大小，必须大于0，默认100"}
  PageSize *QueryForbiddingVisitorIPsByLabelCodeServiceResponseDataPageSize `json:"pageSize,omitempty" xml:"pageSize,omitempty" require:"true" type:"Struct"`
  // {"en":"Query results","zh_CN":"查询结果"}
  Result []*QueryForbiddingVisitorIPsByLabelCodeServiceResponseDataResult `json:"result,omitempty" xml:"result,omitempty" require:"true" type:"Repeated"`
}

func (s QueryForbiddingVisitorIPsByLabelCodeServiceResponseData) String() string {
  return tea.Prettify(s)
}

func (s QueryForbiddingVisitorIPsByLabelCodeServiceResponseData) GoString() string {
  return s.String()
}

func (s *QueryForbiddingVisitorIPsByLabelCodeServiceResponseData) SetTotal(v *QueryForbiddingVisitorIPsByLabelCodeServiceResponseDataTotal) *QueryForbiddingVisitorIPsByLabelCodeServiceResponseData {
  s.Total = v
  return s
}

func (s *QueryForbiddingVisitorIPsByLabelCodeServiceResponseData) SetPageNo(v *QueryForbiddingVisitorIPsByLabelCodeServiceResponseDataPageNo) *QueryForbiddingVisitorIPsByLabelCodeServiceResponseData {
  s.PageNo = v
  return s
}

func (s *QueryForbiddingVisitorIPsByLabelCodeServiceResponseData) SetPageSize(v *QueryForbiddingVisitorIPsByLabelCodeServiceResponseDataPageSize) *QueryForbiddingVisitorIPsByLabelCodeServiceResponseData {
  s.PageSize = v
  return s
}

func (s *QueryForbiddingVisitorIPsByLabelCodeServiceResponseData) SetResult(v []*QueryForbiddingVisitorIPsByLabelCodeServiceResponseDataResult) *QueryForbiddingVisitorIPsByLabelCodeServiceResponseData {
  s.Result = v
  return s
}

type QueryForbiddingVisitorIPsByLabelCodeServiceResponseDataTotal struct {
}

func (s QueryForbiddingVisitorIPsByLabelCodeServiceResponseDataTotal) String() string {
  return tea.Prettify(s)
}

func (s QueryForbiddingVisitorIPsByLabelCodeServiceResponseDataTotal) GoString() string {
  return s.String()
}

type QueryForbiddingVisitorIPsByLabelCodeServiceResponseDataPageNo struct {
}

func (s QueryForbiddingVisitorIPsByLabelCodeServiceResponseDataPageNo) String() string {
  return tea.Prettify(s)
}

func (s QueryForbiddingVisitorIPsByLabelCodeServiceResponseDataPageNo) GoString() string {
  return s.String()
}

type QueryForbiddingVisitorIPsByLabelCodeServiceResponseDataPageSize struct {
}

func (s QueryForbiddingVisitorIPsByLabelCodeServiceResponseDataPageSize) String() string {
  return tea.Prettify(s)
}

func (s QueryForbiddingVisitorIPsByLabelCodeServiceResponseDataPageSize) GoString() string {
  return s.String()
}

type QueryForbiddingVisitorIPsByLabelCodeServiceResponseDataResult struct     {
  // {"en":"Label Code","zh_CN":"标签编码"}
  LabelCode *string `json:"labelCode,omitempty" xml:"labelCode,omitempty" require:"true"`
  // {"en":"Label Name","zh_CN":"标签名称"}
  LabelName *string `json:"labelName,omitempty" xml:"labelName,omitempty" require:"true"`
  // {"en":"IP forbidding","zh_CN":"封禁的IP"}
  Ip *string `json:"ip,omitempty" xml:"ip,omitempty" require:"true"`
  // {"en":"Start time of forbidden","zh_CN":"封禁开始时间"}
  StartTime *QueryForbiddingVisitorIPsByLabelCodeServiceResponseDataResultStartTime `json:"startTime,omitempty" xml:"startTime,omitempty" require:"true" type:"Struct"`
  // {"en":"End time of forbidden","zh_CN":"封禁结束时间"}
  EndTime *QueryForbiddingVisitorIPsByLabelCodeServiceResponseDataResultEndTime `json:"endTime,omitempty" xml:"endTime,omitempty" require:"true" type:"Struct"`
}

func (s QueryForbiddingVisitorIPsByLabelCodeServiceResponseDataResult) String() string {
  return tea.Prettify(s)
}

func (s QueryForbiddingVisitorIPsByLabelCodeServiceResponseDataResult) GoString() string {
  return s.String()
}

func (s *QueryForbiddingVisitorIPsByLabelCodeServiceResponseDataResult) SetLabelCode(v string) *QueryForbiddingVisitorIPsByLabelCodeServiceResponseDataResult {
  s.LabelCode = &v
  return s
}

func (s *QueryForbiddingVisitorIPsByLabelCodeServiceResponseDataResult) SetLabelName(v string) *QueryForbiddingVisitorIPsByLabelCodeServiceResponseDataResult {
  s.LabelName = &v
  return s
}

func (s *QueryForbiddingVisitorIPsByLabelCodeServiceResponseDataResult) SetIp(v string) *QueryForbiddingVisitorIPsByLabelCodeServiceResponseDataResult {
  s.Ip = &v
  return s
}

func (s *QueryForbiddingVisitorIPsByLabelCodeServiceResponseDataResult) SetStartTime(v *QueryForbiddingVisitorIPsByLabelCodeServiceResponseDataResultStartTime) *QueryForbiddingVisitorIPsByLabelCodeServiceResponseDataResult {
  s.StartTime = v
  return s
}

func (s *QueryForbiddingVisitorIPsByLabelCodeServiceResponseDataResult) SetEndTime(v *QueryForbiddingVisitorIPsByLabelCodeServiceResponseDataResultEndTime) *QueryForbiddingVisitorIPsByLabelCodeServiceResponseDataResult {
  s.EndTime = v
  return s
}

type QueryForbiddingVisitorIPsByLabelCodeServiceResponseDataResultStartTime struct {
}

func (s QueryForbiddingVisitorIPsByLabelCodeServiceResponseDataResultStartTime) String() string {
  return tea.Prettify(s)
}

func (s QueryForbiddingVisitorIPsByLabelCodeServiceResponseDataResultStartTime) GoString() string {
  return s.String()
}

type QueryForbiddingVisitorIPsByLabelCodeServiceResponseDataResultEndTime struct {
}

func (s QueryForbiddingVisitorIPsByLabelCodeServiceResponseDataResultEndTime) String() string {
  return tea.Prettify(s)
}

func (s QueryForbiddingVisitorIPsByLabelCodeServiceResponseDataResultEndTime) GoString() string {
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




type ForbidOrResumeVisitorIPsByLabelCodeServiceRequest struct {
  // {"en":"List of operating objects","zh_CN":"操作对象列表"}
  OperationObjectList []*ForbidOrResumeVisitorIPsByLabelCodeServiceRequestOperationObjectList `json:"operationObjectList,omitempty" xml:"operationObjectList,omitempty" require:"true" type:"Repeated"`
  // {"en":"Operation Type, 1: forbid; 2: resume","zh_CN":"操作类型， 1: 封禁； 2: 解禁"}
  OperationType *string `json:"operationType,omitempty" xml:"operationType,omitempty" require:"true"`
  // {"en":"Forbid Duration(minutes),The maximum value is 2628000 minutes(five years), and it will automatically be set to 2628000 if exceeded. Required for forbidding operation, non-required for resuming operation.","zh_CN":"封禁时长（分钟），最大值为2628000分钟（即五年），超过自动设置为2628000。封禁操作时，必填，解禁时非必填。"}
  ForbidTime *int `json:"forbidTime,omitempty" xml:"forbidTime,omitempty"`
}

func (s ForbidOrResumeVisitorIPsByLabelCodeServiceRequest) String() string {
  return tea.Prettify(s)
}

func (s ForbidOrResumeVisitorIPsByLabelCodeServiceRequest) GoString() string {
  return s.String()
}

func (s *ForbidOrResumeVisitorIPsByLabelCodeServiceRequest) SetOperationObjectList(v []*ForbidOrResumeVisitorIPsByLabelCodeServiceRequestOperationObjectList) *ForbidOrResumeVisitorIPsByLabelCodeServiceRequest {
  s.OperationObjectList = v
  return s
}

func (s *ForbidOrResumeVisitorIPsByLabelCodeServiceRequest) SetOperationType(v string) *ForbidOrResumeVisitorIPsByLabelCodeServiceRequest {
  s.OperationType = &v
  return s
}

func (s *ForbidOrResumeVisitorIPsByLabelCodeServiceRequest) SetForbidTime(v int) *ForbidOrResumeVisitorIPsByLabelCodeServiceRequest {
  s.ForbidTime = &v
  return s
}

type ForbidOrResumeVisitorIPsByLabelCodeServiceRequestOperationObjectList struct     {
  // {"en":"Label Code (Please contact technical support for assistance)","zh_CN":"标签编码（请联系专属技术支持获取）"}
  LabelCode *string `json:"labelCode,omitempty" xml:"labelCode,omitempty" require:"true"`
  // {"en":"List of IPs to forbid or resume. The maximum number of IPs is 10,000. The IP address can be v4 or v6,only supports IPV4 segment, does not support IPV6 segment.","zh_CN":"待封禁或解禁的IP列表。IP个数上限为10000个。支持IPV4、IPV6格式，仅支持IPV4段，不支持IPV6段"}
  IpList []*string `json:"ipList,omitempty" xml:"ipList,omitempty" require:"true" type:"Repeated"`
}

func (s ForbidOrResumeVisitorIPsByLabelCodeServiceRequestOperationObjectList) String() string {
  return tea.Prettify(s)
}

func (s ForbidOrResumeVisitorIPsByLabelCodeServiceRequestOperationObjectList) GoString() string {
  return s.String()
}

func (s *ForbidOrResumeVisitorIPsByLabelCodeServiceRequestOperationObjectList) SetLabelCode(v string) *ForbidOrResumeVisitorIPsByLabelCodeServiceRequestOperationObjectList {
  s.LabelCode = &v
  return s
}

func (s *ForbidOrResumeVisitorIPsByLabelCodeServiceRequestOperationObjectList) SetIpList(v []*string) *ForbidOrResumeVisitorIPsByLabelCodeServiceRequestOperationObjectList {
  s.IpList = v
  return s
}

type ForbidOrResumeVisitorIPsByLabelCodeServiceRequestHeader struct {
}

func (s ForbidOrResumeVisitorIPsByLabelCodeServiceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ForbidOrResumeVisitorIPsByLabelCodeServiceRequestHeader) GoString() string {
  return s.String()
}

type ForbidOrResumeVisitorIPsByLabelCodeServicePaths struct {
}

func (s ForbidOrResumeVisitorIPsByLabelCodeServicePaths) String() string {
  return tea.Prettify(s)
}

func (s ForbidOrResumeVisitorIPsByLabelCodeServicePaths) GoString() string {
  return s.String()
}

type ForbidOrResumeVisitorIPsByLabelCodeServiceParameters struct {
}

func (s ForbidOrResumeVisitorIPsByLabelCodeServiceParameters) String() string {
  return tea.Prettify(s)
}

func (s ForbidOrResumeVisitorIPsByLabelCodeServiceParameters) GoString() string {
  return s.String()
}

type ForbidOrResumeVisitorIPsByLabelCodeServiceResponse struct {
  // {"en":"Result Code.If it shows `PartialSuccess`, please pay attention to the details of partial failures in errCode","zh_CN":"响应码，如果为“PartialSuccess”，请关注errCode中部分失败的详情"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Result Message","zh_CN":"响应信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Result Data","zh_CN":"响应数据"}
  Data *ForbidOrResumeVisitorIPsByLabelCodeServiceResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
}

func (s ForbidOrResumeVisitorIPsByLabelCodeServiceResponse) String() string {
  return tea.Prettify(s)
}

func (s ForbidOrResumeVisitorIPsByLabelCodeServiceResponse) GoString() string {
  return s.String()
}

func (s *ForbidOrResumeVisitorIPsByLabelCodeServiceResponse) SetCode(v string) *ForbidOrResumeVisitorIPsByLabelCodeServiceResponse {
  s.Code = &v
  return s
}

func (s *ForbidOrResumeVisitorIPsByLabelCodeServiceResponse) SetMessage(v string) *ForbidOrResumeVisitorIPsByLabelCodeServiceResponse {
  s.Message = &v
  return s
}

func (s *ForbidOrResumeVisitorIPsByLabelCodeServiceResponse) SetData(v *ForbidOrResumeVisitorIPsByLabelCodeServiceResponseData) *ForbidOrResumeVisitorIPsByLabelCodeServiceResponse {
  s.Data = v
  return s
}

type ForbidOrResumeVisitorIPsByLabelCodeServiceResponseData struct {
  // {"en":"Error Code","zh_CN":"业务错误码"}
  ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty" require:"true"`
  // {"en":"Error Message","zh_CN":"业务错误信息"}
  ErrMessage *string `json:"errMessage,omitempty" xml:"errMessage,omitempty" require:"true"`
  // {"en":"Label Code","zh_CN":"标签编码"}
  LabelCode *string `json:"labelCode,omitempty" xml:"labelCode,omitempty" require:"true"`
  // {"en":"List of Failed IPs","zh_CN":"失败的IP列表"}
  FailedIpList []*string `json:"failedIpList,omitempty" xml:"failedIpList,omitempty" require:"true" type:"Repeated"`
}

func (s ForbidOrResumeVisitorIPsByLabelCodeServiceResponseData) String() string {
  return tea.Prettify(s)
}

func (s ForbidOrResumeVisitorIPsByLabelCodeServiceResponseData) GoString() string {
  return s.String()
}

func (s *ForbidOrResumeVisitorIPsByLabelCodeServiceResponseData) SetErrCode(v string) *ForbidOrResumeVisitorIPsByLabelCodeServiceResponseData {
  s.ErrCode = &v
  return s
}

func (s *ForbidOrResumeVisitorIPsByLabelCodeServiceResponseData) SetErrMessage(v string) *ForbidOrResumeVisitorIPsByLabelCodeServiceResponseData {
  s.ErrMessage = &v
  return s
}

func (s *ForbidOrResumeVisitorIPsByLabelCodeServiceResponseData) SetLabelCode(v string) *ForbidOrResumeVisitorIPsByLabelCodeServiceResponseData {
  s.LabelCode = &v
  return s
}

func (s *ForbidOrResumeVisitorIPsByLabelCodeServiceResponseData) SetFailedIpList(v []*string) *ForbidOrResumeVisitorIPsByLabelCodeServiceResponseData {
  s.FailedIpList = v
  return s
}

type ForbidOrResumeVisitorIPsByLabelCodeServiceResponseHeader struct {
}

func (s ForbidOrResumeVisitorIPsByLabelCodeServiceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ForbidOrResumeVisitorIPsByLabelCodeServiceResponseHeader) GoString() string {
  return s.String()
}




type ForbidOrResumeVisitorIPsByDomainServiceRequest struct {
  // {"en":"List of Domains","zh_CN":"域名列表"}
  DomainList []*string `json:"domainList,omitempty" xml:"domainList,omitempty" require:"true" type:"Repeated"`
  // {"en":"List of IPs to forbid or resume.The IP address can be v4 or v6,only supports IPV4 segment, does not support IPV6 segment.","zh_CN":"待封禁或解禁的IP列表。支持IPV4、IPV6格式，仅支持IPV4段，不支持IPV6段"}
  IpList []*string `json:"ipList,omitempty" xml:"ipList,omitempty" require:"true" type:"Repeated"`
  // {"en":"Operate Type(1: forbid; 2: resume;)","zh_CN":"操作类型(1: 封禁； 2: 解禁；)"}
  OperationType *int `json:"operationType,omitempty" xml:"operationType,omitempty" require:"true"`
  // {"en":"Forbid duration(minute), default 43200(means 30 days),max 2628000(means 5 years),it is recommended to specify a number,such as 30","zh_CN":"封禁时长（分钟），不传默认为43200（30天），最大支持2628000（5年），建议指定具体的数值，比如 30"}
  ForbidTime *int `json:"forbidTime,omitempty" xml:"forbidTime,omitempty"`
}

func (s ForbidOrResumeVisitorIPsByDomainServiceRequest) String() string {
  return tea.Prettify(s)
}

func (s ForbidOrResumeVisitorIPsByDomainServiceRequest) GoString() string {
  return s.String()
}

func (s *ForbidOrResumeVisitorIPsByDomainServiceRequest) SetDomainList(v []*string) *ForbidOrResumeVisitorIPsByDomainServiceRequest {
  s.DomainList = v
  return s
}

func (s *ForbidOrResumeVisitorIPsByDomainServiceRequest) SetIpList(v []*string) *ForbidOrResumeVisitorIPsByDomainServiceRequest {
  s.IpList = v
  return s
}

func (s *ForbidOrResumeVisitorIPsByDomainServiceRequest) SetOperationType(v int) *ForbidOrResumeVisitorIPsByDomainServiceRequest {
  s.OperationType = &v
  return s
}

func (s *ForbidOrResumeVisitorIPsByDomainServiceRequest) SetForbidTime(v int) *ForbidOrResumeVisitorIPsByDomainServiceRequest {
  s.ForbidTime = &v
  return s
}

type ForbidOrResumeVisitorIPsByDomainServiceRequestHeader struct {
}

func (s ForbidOrResumeVisitorIPsByDomainServiceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ForbidOrResumeVisitorIPsByDomainServiceRequestHeader) GoString() string {
  return s.String()
}

type ForbidOrResumeVisitorIPsByDomainServicePaths struct {
}

func (s ForbidOrResumeVisitorIPsByDomainServicePaths) String() string {
  return tea.Prettify(s)
}

func (s ForbidOrResumeVisitorIPsByDomainServicePaths) GoString() string {
  return s.String()
}

type ForbidOrResumeVisitorIPsByDomainServiceParameters struct {
}

func (s ForbidOrResumeVisitorIPsByDomainServiceParameters) String() string {
  return tea.Prettify(s)
}

func (s ForbidOrResumeVisitorIPsByDomainServiceParameters) GoString() string {
  return s.String()
}

type ForbidOrResumeVisitorIPsByDomainServiceResponse struct {
  // {"en":"Result Code.If it shows `36010032`,means `partial success`, please pay attention to the details of partial failures in errCode, adjust and try again","zh_CN":"响应码，如果为“36010032”，意味着部分成功，请关注errCode中部分失败的详情"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Result Message","zh_CN":"响应信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Result Data","zh_CN":"响应数据"}
  Data []*string `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
  // {"en":"Error Code","zh_CN":"业务错误码"}
  ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty" require:"true"`
  // {"en":"Error Message","zh_CN":"业务错误信息"}
  ErrMessage *string `json:"errMessage,omitempty" xml:"errMessage,omitempty" require:"true"`
  // {"en":"Failed Domain","zh_CN":"失败的域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {"en":"List of Failed IPs","zh_CN":"失败的IP列表"}
  FailedIpList []*string `json:"failedIpList,omitempty" xml:"failedIpList,omitempty" require:"true" type:"Repeated"`
}

func (s ForbidOrResumeVisitorIPsByDomainServiceResponse) String() string {
  return tea.Prettify(s)
}

func (s ForbidOrResumeVisitorIPsByDomainServiceResponse) GoString() string {
  return s.String()
}

func (s *ForbidOrResumeVisitorIPsByDomainServiceResponse) SetCode(v string) *ForbidOrResumeVisitorIPsByDomainServiceResponse {
  s.Code = &v
  return s
}

func (s *ForbidOrResumeVisitorIPsByDomainServiceResponse) SetMessage(v string) *ForbidOrResumeVisitorIPsByDomainServiceResponse {
  s.Message = &v
  return s
}

func (s *ForbidOrResumeVisitorIPsByDomainServiceResponse) SetData(v []*string) *ForbidOrResumeVisitorIPsByDomainServiceResponse {
  s.Data = v
  return s
}

func (s *ForbidOrResumeVisitorIPsByDomainServiceResponse) SetErrCode(v string) *ForbidOrResumeVisitorIPsByDomainServiceResponse {
  s.ErrCode = &v
  return s
}

func (s *ForbidOrResumeVisitorIPsByDomainServiceResponse) SetErrMessage(v string) *ForbidOrResumeVisitorIPsByDomainServiceResponse {
  s.ErrMessage = &v
  return s
}

func (s *ForbidOrResumeVisitorIPsByDomainServiceResponse) SetDomain(v string) *ForbidOrResumeVisitorIPsByDomainServiceResponse {
  s.Domain = &v
  return s
}

func (s *ForbidOrResumeVisitorIPsByDomainServiceResponse) SetFailedIpList(v []*string) *ForbidOrResumeVisitorIPsByDomainServiceResponse {
  s.FailedIpList = v
  return s
}

type ForbidOrResumeVisitorIPsByDomainServiceResponseHeader struct {
}

func (s ForbidOrResumeVisitorIPsByDomainServiceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ForbidOrResumeVisitorIPsByDomainServiceResponseHeader) GoString() string {
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




