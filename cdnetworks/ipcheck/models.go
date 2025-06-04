package ipcheck

// This file is auto-generated, don't edit it. Thanks.
import (
  "github.com/alibabacloud-go/tea/tea"
)

type QueryCDNServiceRealIPRequest struct {
}

func (s QueryCDNServiceRealIPRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryCDNServiceRealIPRequest) GoString() string {
  return s.String()
}

type QueryCDNServiceRealIPResponse struct {
  // {'en':'result', 'zh_CN':'结果'}
  Result *QueryCDNServiceRealIPResponseResult `json:"result,omitempty" xml:"result,omitempty" require:"true" type:"Struct"`
}

func (s QueryCDNServiceRealIPResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryCDNServiceRealIPResponse) GoString() string {
  return s.String()
}

func (s *QueryCDNServiceRealIPResponse) SetResult(v *QueryCDNServiceRealIPResponseResult) *QueryCDNServiceRealIPResponse {
  s.Result = v
  return s
}

type QueryCDNServiceRealIPResponseResult struct {
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {'en':'Real service IP list of domains', 'zh_CN':'域名对应的真实服务IP列表'}
  WhiteipList *QueryCDNServiceRealIPResponseResultWhiteipList `json:"whiteipList,omitempty" xml:"whiteipList,omitempty" require:"true" type:"Struct"`
}

func (s QueryCDNServiceRealIPResponseResult) String() string {
  return tea.Prettify(s)
}

func (s QueryCDNServiceRealIPResponseResult) GoString() string {
  return s.String()
}

func (s *QueryCDNServiceRealIPResponseResult) SetCode(v string) *QueryCDNServiceRealIPResponseResult {
  s.Code = &v
  return s
}

func (s *QueryCDNServiceRealIPResponseResult) SetWhiteipList(v *QueryCDNServiceRealIPResponseResultWhiteipList) *QueryCDNServiceRealIPResponseResult {
  s.WhiteipList = v
  return s
}

type QueryCDNServiceRealIPResponseResultWhiteipList struct {
  // {'en':'Ip List', 'zh_CN':'真实服务IP列表'}
  Whiteiplist []*string `json:"whiteiplist,omitempty" xml:"whiteiplist,omitempty" require:"true" type:"Repeated"`
  // {'en':'Domain List', 'zh_CN':'域名列表'}
  DomainName []*string `json:"domainName,omitempty" xml:"domainName,omitempty" require:"true" type:"Repeated"`
}

func (s QueryCDNServiceRealIPResponseResultWhiteipList) String() string {
  return tea.Prettify(s)
}

func (s QueryCDNServiceRealIPResponseResultWhiteipList) GoString() string {
  return s.String()
}

func (s *QueryCDNServiceRealIPResponseResultWhiteipList) SetWhiteiplist(v []*string) *QueryCDNServiceRealIPResponseResultWhiteipList {
  s.Whiteiplist = v
  return s
}

func (s *QueryCDNServiceRealIPResponseResultWhiteipList) SetDomainName(v []*string) *QueryCDNServiceRealIPResponseResultWhiteipList {
  s.DomainName = v
  return s
}

type QueryCDNServiceRealIPPaths struct {
}

func (s QueryCDNServiceRealIPPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryCDNServiceRealIPPaths) GoString() string {
  return s.String()
}

type QueryCDNServiceRealIPParameters struct {
  // {'en':'Domain names. Which are separated by semicolons, and it supports 20 domains at max.', 'zh_CN':'域名，以英文分号分隔，最多20个域名'}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {'en':'Node type. Default value is all. 
  //     Optional values:
  //     dyfu: dynamic relay;
  //     stfu: static relay; 
  //     fu: dynamic relays and static relays;
  //     edge: edge node;
  //     all: dynamic relays, static relaysand edge nodes.', 'zh_CN':'节点类型，不传默认all。
  // 	dyfu：动态父； stfu:静态父； fu：动态父+静态父；&nbsp; edge ：边缘机器；&nbsp; all：动静+边缘机器。'}
  Viewtype *string `json:"viewtype,omitempty" xml:"viewtype,omitempty"`
  // {'en':'IP form. Default value is ipseg. 
  //     Optional values:
  //     realip: real IP; 
  //     ipseg: IP segment.', 'zh_CN':'ip形式，不传默认 ipseg。	realip：真实IP ； ipseg：ip段。'}
  Iptype *string `json:"iptype,omitempty" xml:"iptype,omitempty"`
}

func (s QueryCDNServiceRealIPParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryCDNServiceRealIPParameters) GoString() string {
  return s.String()
}

func (s *QueryCDNServiceRealIPParameters) SetDomain(v string) *QueryCDNServiceRealIPParameters {
  s.Domain = &v
  return s
}

func (s *QueryCDNServiceRealIPParameters) SetViewtype(v string) *QueryCDNServiceRealIPParameters {
  s.Viewtype = &v
  return s
}

func (s *QueryCDNServiceRealIPParameters) SetIptype(v string) *QueryCDNServiceRealIPParameters {
  s.Iptype = &v
  return s
}

type QueryCDNServiceRealIPRequestHeader struct {
}

func (s QueryCDNServiceRealIPRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryCDNServiceRealIPRequestHeader) GoString() string {
  return s.String()
}

type QueryCDNServiceRealIPResponseHeader struct {
}

func (s QueryCDNServiceRealIPResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryCDNServiceRealIPResponseHeader) GoString() string {
  return s.String()
}




type QuerySpecificIPBelongRequest struct {
  // {'en':'IP address, use English comma to separate two items. Every IP address needs to following regular expression rule of   ((2[0-4]\\d|25[0-5]|1\\d\\d|0|[1-9]\\d?)\\.){3}(2[0-4]\\d|25[0-5]|1\\d\\d|0|[1-9]\\d?).   The default number of IPs cannot exceed 20 (you can contact technical support to adjust) .', 'zh_CN':'ip地址，以英文逗号分隔，每个ip都需要符合正则((2[0-4]\\d|25[0-5]|1\\d\\d|0|[1-9]\\d?)\\.){3}(2[0-4]\\d|25[0-5]|1\\d\\d|0|[1-9]\\d?)，ip个数默认不能超过20（可联系技术支持调整）'}
  Ip []*string `json:"ip,omitempty" xml:"ip,omitempty" require:"true" type:"Repeated"`
}

func (s QuerySpecificIPBelongRequest) String() string {
  return tea.Prettify(s)
}

func (s QuerySpecificIPBelongRequest) GoString() string {
  return s.String()
}

func (s *QuerySpecificIPBelongRequest) SetIp(v []*string) *QuerySpecificIPBelongRequest {
  s.Ip = v
  return s
}

type QuerySpecificIPBelongResponse struct {
  // {'en':'checkList', 'zh_CN':'结果数据'}
  CheckList []*QuerySpecificIPBelongResponseCheckList `json:"checkList,omitempty" xml:"checkList,omitempty" require:"true" type:"Repeated"`
}

func (s QuerySpecificIPBelongResponse) String() string {
  return tea.Prettify(s)
}

func (s QuerySpecificIPBelongResponse) GoString() string {
  return s.String()
}

func (s *QuerySpecificIPBelongResponse) SetCheckList(v []*QuerySpecificIPBelongResponseCheckList) *QuerySpecificIPBelongResponse {
  s.CheckList = v
  return s
}

type QuerySpecificIPBelongResponseCheckList struct     {
  // {'en':'yes: the IP belongs to Our system,
  //         no: the IP does not belong to Our system', 'zh_CN':'yes：ip属于我司，no：ip不属于我司'}
  Response *string `json:"response,omitempty" xml:"response,omitempty" require:"true"`
  // {'en':'IP addresses', 'zh_CN':'ip地址'}
  Ip *string `json:"ip,omitempty" xml:"ip,omitempty" require:"true"`
}

func (s QuerySpecificIPBelongResponseCheckList) String() string {
  return tea.Prettify(s)
}

func (s QuerySpecificIPBelongResponseCheckList) GoString() string {
  return s.String()
}

func (s *QuerySpecificIPBelongResponseCheckList) SetResponse(v string) *QuerySpecificIPBelongResponseCheckList {
  s.Response = &v
  return s
}

func (s *QuerySpecificIPBelongResponseCheckList) SetIp(v string) *QuerySpecificIPBelongResponseCheckList {
  s.Ip = &v
  return s
}

type QuerySpecificIPBelongPaths struct {
}

func (s QuerySpecificIPBelongPaths) String() string {
  return tea.Prettify(s)
}

func (s QuerySpecificIPBelongPaths) GoString() string {
  return s.String()
}

type QuerySpecificIPBelongParameters struct {
}

func (s QuerySpecificIPBelongParameters) String() string {
  return tea.Prettify(s)
}

func (s QuerySpecificIPBelongParameters) GoString() string {
  return s.String()
}

type QuerySpecificIPBelongRequestHeader struct {
}

func (s QuerySpecificIPBelongRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QuerySpecificIPBelongRequestHeader) GoString() string {
  return s.String()
}

type QuerySpecificIPBelongResponseHeader struct {
}

func (s QuerySpecificIPBelongResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QuerySpecificIPBelongResponseHeader) GoString() string {
  return s.String()
}




type IpInfoServiceRequest struct {
  // {'en':'The list of IP that needs to be querying is 20 times a single time.', 'zh_CN':'需要查询的IP列表，单次最大20个（联系技术支持可调上限）'}
  Ip []*string `json:"ip,omitempty" xml:"ip,omitempty" require:"true" type:"Repeated"`
}

func (s IpInfoServiceRequest) String() string {
  return tea.Prettify(s)
}

func (s IpInfoServiceRequest) GoString() string {
  return s.String()
}

func (s *IpInfoServiceRequest) SetIp(v []*string) *IpInfoServiceRequest {
  s.Ip = v
  return s
}

type IpInfoServiceResponse struct {
  // {'en':'result', 'zh_CN':'结果'}
  Result []*IpInfoServiceResponseResult `json:"result,omitempty" xml:"result,omitempty" require:"true" type:"Repeated"`
}

func (s IpInfoServiceResponse) String() string {
  return tea.Prettify(s)
}

func (s IpInfoServiceResponse) GoString() string {
  return s.String()
}

func (s *IpInfoServiceResponse) SetResult(v []*IpInfoServiceResponseResult) *IpInfoServiceResponse {
  s.Result = v
  return s
}

type IpInfoServiceResponseResult struct     {
  // {'en':'IP addresses', 'zh_CN':'IP地址'}
  Ip *string `json:"ip,omitempty" xml:"ip,omitempty" require:"true"`
  // {'en':'Whether to network the our IP
  // 
  //         1.true is the node IP of our CDN
  // 
  //         2.false is not the node IP of the CDN', 'zh_CN':'是否我司CDN的IP
  //         1.true 是我司CDN的节点IP
  //         2.false &nbsp; 不是我司CDN的节点IP'}
  IsCdnIp *bool `json:"isCdnIp,omitempty" xml:"isCdnIp,omitempty" require:"true"`
  // {'en':'If it is not a node of the CDN, it will not return; if it is not planned, it will return unknown.', 'zh_CN':'归属国家地区；不是我司CDN的节点，不返回；如未规划的则返回未知。'}
  Country *string `json:"country,omitempty" xml:"country,omitempty" require:"true"`
  // {'en':'If it is not a node of the CDN, it will not return; if it is not planned, it will return unknown.', 'zh_CN':'归属省份；不是我司CDN的节点，不返回；如未规划的则返回未知；'}
  Province *string `json:"province,omitempty" xml:"province,omitempty" require:"true"`
  // {'en':'If it is not a node of the CDN, it will not return; if it is not planned, it will return unknown.', 'zh_CN':'归属城市；不是我司CDN的节点，不返回；如未规划的则返回未知；'}
  City *string `json:"city,omitempty" xml:"city,omitempty" require:"true"`
  // {'en':'If it is not a node of the CDN, it will not return; if it is not planned, it will return unknown.', 'zh_CN':'归属运营商；不是我司CDN的节点，不返回；如未规划的则返回未知。'}
  Isp *string `json:"isp,omitempty" xml:"isp,omitempty" require:"true"`
}

func (s IpInfoServiceResponseResult) String() string {
  return tea.Prettify(s)
}

func (s IpInfoServiceResponseResult) GoString() string {
  return s.String()
}

func (s *IpInfoServiceResponseResult) SetIp(v string) *IpInfoServiceResponseResult {
  s.Ip = &v
  return s
}

func (s *IpInfoServiceResponseResult) SetIsCdnIp(v bool) *IpInfoServiceResponseResult {
  s.IsCdnIp = &v
  return s
}

func (s *IpInfoServiceResponseResult) SetCountry(v string) *IpInfoServiceResponseResult {
  s.Country = &v
  return s
}

func (s *IpInfoServiceResponseResult) SetProvince(v string) *IpInfoServiceResponseResult {
  s.Province = &v
  return s
}

func (s *IpInfoServiceResponseResult) SetCity(v string) *IpInfoServiceResponseResult {
  s.City = &v
  return s
}

func (s *IpInfoServiceResponseResult) SetIsp(v string) *IpInfoServiceResponseResult {
  s.Isp = &v
  return s
}

type IpInfoServicePaths struct {
}

func (s IpInfoServicePaths) String() string {
  return tea.Prettify(s)
}

func (s IpInfoServicePaths) GoString() string {
  return s.String()
}

type IpInfoServiceParameters struct {
}

func (s IpInfoServiceParameters) String() string {
  return tea.Prettify(s)
}

func (s IpInfoServiceParameters) GoString() string {
  return s.String()
}

type IpInfoServiceRequestHeader struct {
}

func (s IpInfoServiceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s IpInfoServiceRequestHeader) GoString() string {
  return s.String()
}

type IpInfoServiceResponseHeader struct {
}

func (s IpInfoServiceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s IpInfoServiceResponseHeader) GoString() string {
  return s.String()
}




