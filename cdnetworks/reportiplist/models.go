package reportiplist

// This file is auto-generated, don't edit it. Thanks.
import (
  "github.com/alibabacloud-go/tea/tea"
)

type QueryCdnIpListRequest struct {
  // {"en":"Domain list.
  // 1. All domains are queried if this field is not specified;
  // 2. Number of domains can be adjusted depending on different accounts. The default value is 20(this limit applies to the empty value);", "zh_CN":"域名列表
  // 域名个数限制根据账号可调，默认为20个（不传递时同样受此限制）；"}
  QueryCdnIpListDomainList *QueryCdnIpListDomainList `json:"domain-list,omitempty" xml:"domain-list,omitempty" require:"true"`
}

func (s QueryCdnIpListRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryCdnIpListRequest) GoString() string {
  return s.String()
}

func (s *QueryCdnIpListRequest) SetDomainList(v *QueryCdnIpListDomainList) *QueryCdnIpListRequest {
  s.QueryCdnIpListDomainList = v
  return s
}

type QueryCdnIpListDomainList struct {
  // {"en":"Domain", "zh_CN":"域名"}
  DomainName []*string `json:"domain-name,omitempty" xml:"domain-name,omitempty" require:"true" type:"Repeated"`
}

func (s QueryCdnIpListDomainList) String() string {
  return tea.Prettify(s)
}

func (s QueryCdnIpListDomainList) GoString() string {
  return s.String()
}

func (s *QueryCdnIpListDomainList) SetDomainName(v []*string) *QueryCdnIpListDomainList {
  s.DomainName = v
  return s
}

type QueryCdnIpListResponse struct {
  // {"en":"domainServerList", "zh_CN":"CDN服务IP数据"}
  Result []*QueryCdnIpListResponseResult `json:"domain-server-list,omitempty" xml:"domain-server-list,omitempty" require:"true" type:"Repeated"`
}

func (s QueryCdnIpListResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryCdnIpListResponse) GoString() string {
  return s.String()
}

func (s *QueryCdnIpListResponse) SetResult(v []*QueryCdnIpListResponseResult) *QueryCdnIpListResponse {
  s.Result = v
  return s
}

type QueryCdnIpListResponseResult struct     {
  // {"en":"Domain", "zh_CN":"域名"}
  DomainName *string `json:"domain-name,omitempty" xml:"domain-name,omitempty" require:"true"`
  // {"en":"serverList", "zh_CN":"服务数据"}
  ServerList []*QueryCdnIpListResponseResultServerList `json:"server-list,omitempty" xml:"server-list,omitempty" require:"true" type:"Repeated"`
}

func (s QueryCdnIpListResponseResult) String() string {
  return tea.Prettify(s)
}

func (s QueryCdnIpListResponseResult) GoString() string {
  return s.String()
}

func (s *QueryCdnIpListResponseResult) SetDomainName(v string) *QueryCdnIpListResponseResult {
  s.DomainName = &v
  return s
}

func (s *QueryCdnIpListResponseResult) SetServerList(v []*QueryCdnIpListResponseResultServerList) *QueryCdnIpListResponseResult {
  s.ServerList = v
  return s
}

type QueryCdnIpListResponseResultServerList struct     {
  // {"en":"Server node IP", "zh_CN":"覆盖节点IP"}
  Server *string `json:"server,omitempty" xml:"server,omitempty" require:"true"`
}

func (s QueryCdnIpListResponseResultServerList) String() string {
  return tea.Prettify(s)
}

func (s QueryCdnIpListResponseResultServerList) GoString() string {
  return s.String()
}

func (s *QueryCdnIpListResponseResultServerList) SetServer(v string) *QueryCdnIpListResponseResultServerList {
  s.Server = &v
  return s
}

type QueryCdnIpListPaths struct {
}

func (s QueryCdnIpListPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryCdnIpListPaths) GoString() string {
  return s.String()
}

type QueryCdnIpListParameters struct {
}

func (s QueryCdnIpListParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryCdnIpListParameters) GoString() string {
  return s.String()
}

type QueryCdnIpListRequestHeader struct {
}

func (s QueryCdnIpListRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryCdnIpListRequestHeader) GoString() string {
  return s.String()
}

type QueryCdnIpListResponseHeader struct {
}

func (s QueryCdnIpListResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryCdnIpListResponseHeader) GoString() string {
  return s.String()
}




type QueryServiceIPInfoDetailRequest struct {
  // {"en":"Domains: 1.Domain is not uploaded: Query all domains of the account (More than 20 domains will error,you can contact technical support for adjustment); 2.Domain is uploaded: Up to 20 domains are supported(you can contact technical support for adjustment).", "zh_CN":"域名: 1.未传递domain时:查询账号下所有全部域名(域名超过20个则报错,可联系技术支持调整)。"}
  Domain []*string `json:"domain,omitempty" xml:"domain,omitempty" type:"Repeated"`
}

func (s QueryServiceIPInfoDetailRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryServiceIPInfoDetailRequest) GoString() string {
  return s.String()
}

func (s *QueryServiceIPInfoDetailRequest) SetDomain(v []*string) *QueryServiceIPInfoDetailRequest {
  s.Domain = v
  return s
}

type QueryServiceIPInfoDetailResponse struct {
  // {"en":"Request result status code.", "zh_CN":"请求结果状态码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Request result information", "zh_CN":"请求结果信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Data", "zh_CN":"结果"}
  Data []*QueryServiceIPInfoDetailIpbInfoDetail `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s QueryServiceIPInfoDetailResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryServiceIPInfoDetailResponse) GoString() string {
  return s.String()
}

func (s *QueryServiceIPInfoDetailResponse) SetCode(v string) *QueryServiceIPInfoDetailResponse {
  s.Code = &v
  return s
}

func (s *QueryServiceIPInfoDetailResponse) SetMessage(v string) *QueryServiceIPInfoDetailResponse {
  s.Message = &v
  return s
}

func (s *QueryServiceIPInfoDetailResponse) SetData(v []*QueryServiceIPInfoDetailIpbInfoDetail) *QueryServiceIPInfoDetailResponse {
  s.Data = v
  return s
}

type QueryServiceIPInfoDetailIpbInfoDetail struct {
  // {"en":"The domain you added on the platform.", "zh_CN":"加速域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {"en":"The name of the view that serves the domain. A view is a special area division designed specifically for our CDN services, refining CDN coverage that may correspond to specific regions of a country, provinces, or even cities or ISPs.", "zh_CN":"域名覆盖的view名称。view是一种内部区域划分，旨在细化CDN的覆盖，可能对应国家的特定区域、省份，甚至具体到城市或ISP服务提供商。"}
  ViewName *string `json:"viewName,omitempty" xml:"viewName,omitempty" require:"true"`
  // {"en":"The country to which the view belongs, identified by standard country codes (e.g., CN for China, US for the United States).", "zh_CN":"view所属的国家，通过标准国家代码标识（例如：CN表示中国，US表示美国）。"}
  ViewCountry *string `json:"viewCountry,omitempty" xml:"viewCountry,omitempty" require:"true"`
  // {"en":"The IP address used to serve the domain (e.g., 192.0.2.1).", "zh_CN":"用于加速该域名的IP地址（例如：192.0.2.1）。"}
  Ip *string `json:"ip,omitempty" xml:"ip,omitempty" require:"true"`
  // {"en":"The country to which the IP address belongs, identified by standard country codes (e.g., CN for China, US for the United States).", "zh_CN":" IP地址所属的国家，通过标准国家代码标识（例如：CN表示中国，US表示美国）。"}
  IpCountry *string `json:"ipCountry,omitempty" xml:"ipCountry,omitempty" require:"true"`
  // {"en":"The Internet Service Provider (ISP) to which the IP address belongs.", "zh_CN":" IP地址所属的互联网服务提供商（ISP）。"}
  IpISP *string `json:"ipISP,omitempty" xml:"ipISP,omitempty" require:"true"`
}

func (s QueryServiceIPInfoDetailIpbInfoDetail) String() string {
  return tea.Prettify(s)
}

func (s QueryServiceIPInfoDetailIpbInfoDetail) GoString() string {
  return s.String()
}

func (s *QueryServiceIPInfoDetailIpbInfoDetail) SetDomain(v string) *QueryServiceIPInfoDetailIpbInfoDetail {
  s.Domain = &v
  return s
}

func (s *QueryServiceIPInfoDetailIpbInfoDetail) SetViewName(v string) *QueryServiceIPInfoDetailIpbInfoDetail {
  s.ViewName = &v
  return s
}

func (s *QueryServiceIPInfoDetailIpbInfoDetail) SetViewCountry(v string) *QueryServiceIPInfoDetailIpbInfoDetail {
  s.ViewCountry = &v
  return s
}

func (s *QueryServiceIPInfoDetailIpbInfoDetail) SetIp(v string) *QueryServiceIPInfoDetailIpbInfoDetail {
  s.Ip = &v
  return s
}

func (s *QueryServiceIPInfoDetailIpbInfoDetail) SetIpCountry(v string) *QueryServiceIPInfoDetailIpbInfoDetail {
  s.IpCountry = &v
  return s
}

func (s *QueryServiceIPInfoDetailIpbInfoDetail) SetIpISP(v string) *QueryServiceIPInfoDetailIpbInfoDetail {
  s.IpISP = &v
  return s
}

type QueryServiceIPInfoDetailPaths struct {
}

func (s QueryServiceIPInfoDetailPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryServiceIPInfoDetailPaths) GoString() string {
  return s.String()
}

type QueryServiceIPInfoDetailParameters struct {
}

func (s QueryServiceIPInfoDetailParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryServiceIPInfoDetailParameters) GoString() string {
  return s.String()
}

type QueryServiceIPInfoDetailRequestHeader struct {
}

func (s QueryServiceIPInfoDetailRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryServiceIPInfoDetailRequestHeader) GoString() string {
  return s.String()
}

type QueryServiceIPInfoDetailResponseHeader struct {
}

func (s QueryServiceIPInfoDetailResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryServiceIPInfoDetailResponseHeader) GoString() string {
  return s.String()
}




