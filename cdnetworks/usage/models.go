package usage

// This file is auto-generated, don't edit it. Thanks.
import (
  "github.com/alibabacloud-go/tea/tea"
)

type GetPeakUsageRequest struct {
  // {"en":"Start time:
  //   1.The format is yyyy-MM-dd; 
  //   2.And smaller than the current time and 'dateTo'; 
  //   3.Only the data of the last 2 years can be queried.
  //   4.Period between dataFrom and dateTo should not be longer than 31 days.", "zh_CN":"开始时间:
  //   1.格式为yyyy-MM-dd; 
  //   2.并且小于当前时间和dateTo; 
  //   3.只能查询最近2年内数据;
  //   4.dateFrom和dateTo相差不能超过31天。"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty" require:"true"`
  // {"en":"End time:
  //   1.The format is yyyy-MM-dd; 
  //   2.Must be greater than 'dateFrom'; 
  //   3.If it's greater than the current time, then the current time is assigned as the value", "zh_CN":"结束时间:
  //   1.格式为yyyy-MM-dd; 
  //   2.必须大于dateFrom; 
  //   3.如果大于当前时间,则重新赋值为当前时间。"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty" require:"true"`
  // {"en":"Contract, only support send one contract at a time.", "zh_CN":"合同,仅支持传单个"}
  Contract *string `json:"contract,omitempty" xml:"contract,omitempty" require:"true"`
}

func (s GetPeakUsageRequest) String() string {
  return tea.Prettify(s)
}

func (s GetPeakUsageRequest) GoString() string {
  return s.String()
}

func (s *GetPeakUsageRequest) SetDateFrom(v string) *GetPeakUsageRequest {
  s.DateFrom = &v
  return s
}

func (s *GetPeakUsageRequest) SetDateTo(v string) *GetPeakUsageRequest {
  s.DateTo = &v
  return s
}

func (s *GetPeakUsageRequest) SetContract(v string) *GetPeakUsageRequest {
  s.Contract = &v
  return s
}

type GetPeakUsageResponse struct {
  // {"en":"contract", "zh_CN":"合同"}
  Contract *string `json:"contract,omitempty" xml:"contract,omitempty" require:"true"`
  // {"en":"Date time, the format is yyyy-MM-dd HH:mm", "zh_CN":"时间,格式是 yyyy-MM-dd HH:mm"}
  Peakdate *string `json:"peakdate,omitempty" xml:"peakdate,omitempty" require:"true"`
  // {"en":"peak usage, the unit is byte", "zh_CN":"峰值用量,单位为byte"}
  PeakusageByte *string `json:"peakusageByte,omitempty" xml:"peakusageByte,omitempty" require:"true"`
  // {"en":"peak usage, the unit is GB", "zh_CN":"峰值用量,单位为GB"}
  PeakusageGb *string `json:"peakusageGb,omitempty" xml:"peakusageGb,omitempty" require:"true"`
  DomainData []*GetPeakUsageResponseDomainData `json:"domainData,omitempty" xml:"domainData,omitempty" require:"true" type:"Repeated"`
}

func (s GetPeakUsageResponse) String() string {
  return tea.Prettify(s)
}

func (s GetPeakUsageResponse) GoString() string {
  return s.String()
}

func (s *GetPeakUsageResponse) SetContract(v string) *GetPeakUsageResponse {
  s.Contract = &v
  return s
}

func (s *GetPeakUsageResponse) SetPeakdate(v string) *GetPeakUsageResponse {
  s.Peakdate = &v
  return s
}

func (s *GetPeakUsageResponse) SetPeakusageByte(v string) *GetPeakUsageResponse {
  s.PeakusageByte = &v
  return s
}

func (s *GetPeakUsageResponse) SetPeakusageGb(v string) *GetPeakUsageResponse {
  s.PeakusageGb = &v
  return s
}

func (s *GetPeakUsageResponse) SetDomainData(v []*GetPeakUsageResponseDomainData) *GetPeakUsageResponse {
  s.DomainData = v
  return s
}

type GetPeakUsageResponseDomainData struct     {
  // {"en":"domain under the contract", "zh_CN":"合同下的域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {"en":"domain peak usage, the unit is byte", "zh_CN":"域名峰值用量,单位为byte"}
  UsageByte *string `json:"usageByte,omitempty" xml:"usageByte,omitempty" require:"true"`
  // {"en":"domain peak usage, the unit is GB", "zh_CN":"域名峰值用量,单位为GB"}
  UsageGb *string `json:"usageGb,omitempty" xml:"usageGb,omitempty" require:"true"`
}

func (s GetPeakUsageResponseDomainData) String() string {
  return tea.Prettify(s)
}

func (s GetPeakUsageResponseDomainData) GoString() string {
  return s.String()
}

func (s *GetPeakUsageResponseDomainData) SetDomain(v string) *GetPeakUsageResponseDomainData {
  s.Domain = &v
  return s
}

func (s *GetPeakUsageResponseDomainData) SetUsageByte(v string) *GetPeakUsageResponseDomainData {
  s.UsageByte = &v
  return s
}

func (s *GetPeakUsageResponseDomainData) SetUsageGb(v string) *GetPeakUsageResponseDomainData {
  s.UsageGb = &v
  return s
}

type GetPeakUsagePaths struct {
}

func (s GetPeakUsagePaths) String() string {
  return tea.Prettify(s)
}

func (s GetPeakUsagePaths) GoString() string {
  return s.String()
}

type GetPeakUsageParameters struct {
}

func (s GetPeakUsageParameters) String() string {
  return tea.Prettify(s)
}

func (s GetPeakUsageParameters) GoString() string {
  return s.String()
}

type GetPeakUsageRequestHeader struct {
}

func (s GetPeakUsageRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetPeakUsageRequestHeader) GoString() string {
  return s.String()
}

type GetPeakUsageResponseHeader struct {
}

func (s GetPeakUsageResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetPeakUsageResponseHeader) GoString() string {
  return s.String()
}




type GetContractUsageRequest struct {
  // {"en":"Start time:
  //   1.The format is yyyy-MM-dd; 
  //   2.And smaller than the current time and 'dateTo'; 
  //   3.Only the data of the last 2 years can be queried.
  //   4.Period between dataFrom and dateTo should not be longer than 31 days", "zh_CN":"开始时间：
  //   1.格式为yyyy-MM-dd; 
  //   2.并且小于当前时间和dateTo; 
  //   3.只能查询最近2年内数据;
  //   4.dateFrom和dateTo相差不能超过31天"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty" require:"true"`
  // {"en":"End time:
  //   1.The format is yyyy-MM-dd; 
  //   2.Must be greater than 'dateFrom'; 
  //   3.If it's greater than the current time, then the current time is assigned as the value.", "zh_CN":"结束时间：
  //   1.格式为yyyy-MM-dd; 
  //   2.必须大于dateFrom; 
  //   3.如果大于当前时间,则重新赋值为当前时间;"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty" require:"true"`
  // {"en":"contract, only support send one contract at a time", "zh_CN":"合同,仅支持传一个合同"}
  Contract *string `json:"contract,omitempty" xml:"contract,omitempty" require:"true"`
  // {"en":"Data granularity:
  // h: hourly
  // d: dayly
  // If not specified, daily is set as the default value", "zh_CN":"数据粒度：
  // h: 小时粒度
  // d: 天粒度
  // 不传默认为天粒度"}
  Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetContractUsageRequest) String() string {
  return tea.Prettify(s)
}

func (s GetContractUsageRequest) GoString() string {
  return s.String()
}

func (s *GetContractUsageRequest) SetDateFrom(v string) *GetContractUsageRequest {
  s.DateFrom = &v
  return s
}

func (s *GetContractUsageRequest) SetDateTo(v string) *GetContractUsageRequest {
  s.DateTo = &v
  return s
}

func (s *GetContractUsageRequest) SetContract(v string) *GetContractUsageRequest {
  s.Contract = &v
  return s
}

func (s *GetContractUsageRequest) SetType(v string) *GetContractUsageRequest {
  s.Type = &v
  return s
}

type GetContractUsageResponse struct {
  // {"en":"contract", "zh_CN":"合同"}
  Contract *string `json:"contract,omitempty" xml:"contract,omitempty" require:"true"`
  Usage []*GetContractUsageResponseUsage `json:"usage,omitempty" xml:"usage,omitempty" require:"true" type:"Repeated"`
}

func (s GetContractUsageResponse) String() string {
  return tea.Prettify(s)
}

func (s GetContractUsageResponse) GoString() string {
  return s.String()
}

func (s *GetContractUsageResponse) SetContract(v string) *GetContractUsageResponse {
  s.Contract = &v
  return s
}

func (s *GetContractUsageResponse) SetUsage(v []*GetContractUsageResponseUsage) *GetContractUsageResponse {
  s.Usage = v
  return s
}

type GetContractUsageResponseUsage struct     {
  // {"en":"DateTime, the format is yyyy-MM-dd HH:mm", "zh_CN":"时间,格式是 yyyy-MM-dd HH:mm"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {"en":"Usage, the unit is byte", "zh_CN":"用量,单位为byte"}
  Usage *string `json:"usage,omitempty" xml:"usage,omitempty" require:"true"`
}

func (s GetContractUsageResponseUsage) String() string {
  return tea.Prettify(s)
}

func (s GetContractUsageResponseUsage) GoString() string {
  return s.String()
}

func (s *GetContractUsageResponseUsage) SetTimestamp(v string) *GetContractUsageResponseUsage {
  s.Timestamp = &v
  return s
}

func (s *GetContractUsageResponseUsage) SetUsage(v string) *GetContractUsageResponseUsage {
  s.Usage = &v
  return s
}

type GetContractUsagePaths struct {
}

func (s GetContractUsagePaths) String() string {
  return tea.Prettify(s)
}

func (s GetContractUsagePaths) GoString() string {
  return s.String()
}

type GetContractUsageParameters struct {
}

func (s GetContractUsageParameters) String() string {
  return tea.Prettify(s)
}

func (s GetContractUsageParameters) GoString() string {
  return s.String()
}

type GetContractUsageRequestHeader struct {
}

func (s GetContractUsageRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetContractUsageRequestHeader) GoString() string {
  return s.String()
}

type GetContractUsageResponseHeader struct {
}

func (s GetContractUsageResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetContractUsageResponseHeader) GoString() string {
  return s.String()
}




type GetDomainUsageRequest struct {
  // {"en":"Start time 1.The format is yyyy-MM-dd; 2.And smaller than the current time and 'dateTo'; 3.Only the data of the last 2 years can be queried.4.Period between dataFrom and dateTo should not be longer than 31 days", "zh_CN":"开始时间 1.格式为yyyy-MM-dd； 2.并且小于当前时间和dateTo； 3.只能查询最近2年内数据；4.dateFrom和dateTo相差不能超过31天；"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty" require:"true"`
  // {"en":"End time 1.The format is yyyy-MM-dd; 2.Must be greater than 'dateFrom'; 3.If it's greater than the current time, then the current time is assigned as the value.", "zh_CN":"结束时间 1.格式为yyyy-MM-dd； 2.必须大于dateFrom； 3.如果大于当前时间，则重新赋值为当前时间；"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty" require:"true"`
  // {"en":"Domain. Only support send one domain at a time", "zh_CN":"域名，仅支持传单个域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  // {"en":"contract that the domain belong, only support send one contract at a time", "zh_CN":"域名所属合同，仅支持传一个合同"}
  Contract *string `json:"contract,omitempty" xml:"contract,omitempty" require:"true"`
  // {"en":"Data granularity
  // 
  // h: hourly
  // 
  // d: daily
  // 
  // If not specified, daily is set as the default value", "zh_CN":"数据粒度：
  // 
  // h：小时粒度
  // 
  // d：天粒度
  // 
  // 不传默认为天粒度"}
  Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetDomainUsageRequest) String() string {
  return tea.Prettify(s)
}

func (s GetDomainUsageRequest) GoString() string {
  return s.String()
}

func (s *GetDomainUsageRequest) SetDateFrom(v string) *GetDomainUsageRequest {
  s.DateFrom = &v
  return s
}

func (s *GetDomainUsageRequest) SetDateTo(v string) *GetDomainUsageRequest {
  s.DateTo = &v
  return s
}

func (s *GetDomainUsageRequest) SetDomain(v string) *GetDomainUsageRequest {
  s.Domain = &v
  return s
}

func (s *GetDomainUsageRequest) SetContract(v string) *GetDomainUsageRequest {
  s.Contract = &v
  return s
}

func (s *GetDomainUsageRequest) SetType(v string) *GetDomainUsageRequest {
  s.Type = &v
  return s
}

type GetDomainUsageResponse struct {
  // {"en":"Domain", "zh_CN":"域名"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
  Usage []*GetDomainUsageResponseUsage `json:"usage,omitempty" xml:"usage,omitempty" require:"true" type:"Repeated"`
}

func (s GetDomainUsageResponse) String() string {
  return tea.Prettify(s)
}

func (s GetDomainUsageResponse) GoString() string {
  return s.String()
}

func (s *GetDomainUsageResponse) SetDomain(v string) *GetDomainUsageResponse {
  s.Domain = &v
  return s
}

func (s *GetDomainUsageResponse) SetUsage(v []*GetDomainUsageResponseUsage) *GetDomainUsageResponse {
  s.Usage = v
  return s
}

type GetDomainUsageResponseUsage struct     {
  // {"en":"DateTime, the format is yyyy-MM-dd HH:mm", "zh_CN":"时间，格式是 yyyy-MM-dd HH:mm"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
  // {"en":"Usage,the unit is byte", "zh_CN":"用量，单位为byte"}
  Usage *string `json:"usage,omitempty" xml:"usage,omitempty" require:"true"`
}

func (s GetDomainUsageResponseUsage) String() string {
  return tea.Prettify(s)
}

func (s GetDomainUsageResponseUsage) GoString() string {
  return s.String()
}

func (s *GetDomainUsageResponseUsage) SetTimestamp(v string) *GetDomainUsageResponseUsage {
  s.Timestamp = &v
  return s
}

func (s *GetDomainUsageResponseUsage) SetUsage(v string) *GetDomainUsageResponseUsage {
  s.Usage = &v
  return s
}

type GetDomainUsagePaths struct {
}

func (s GetDomainUsagePaths) String() string {
  return tea.Prettify(s)
}

func (s GetDomainUsagePaths) GoString() string {
  return s.String()
}

type GetDomainUsageParameters struct {
}

func (s GetDomainUsageParameters) String() string {
  return tea.Prettify(s)
}

func (s GetDomainUsageParameters) GoString() string {
  return s.String()
}

type GetDomainUsageRequestHeader struct {
}

func (s GetDomainUsageRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetDomainUsageRequestHeader) GoString() string {
  return s.String()
}

type GetDomainUsageResponseHeader struct {
}

func (s GetDomainUsageResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetDomainUsageResponseHeader) GoString() string {
  return s.String()
}




