// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/tea"
)

type QueryDispatchDomainsRequest struct {
  // {"en":"Number per page", "zh_CN":"分页查询条数"}
  Limit *int `json:"limit,omitempty" xml:"limit,omitempty" require:"true"`
  // {"en":"Start number", "zh_CN":"分页查询起始记录"}
  Start *int `json:"start,omitempty" xml:"start,omitempty" require:"true"`
  // {"en":"Leave empty to return Chinese result(default), en: Return English result", "zh_CN":"为空返回中文结果（默认），en：返回英文结果"}
  Language *string `json:"language,omitempty" xml:"language,omitempty" require:"true"`
  // {"en":"Domain name(Fuzzy search), leave empty to return all", "zh_CN":"域名模糊搜索参数  域名模糊搜索，如果没有填写域名，则返回该用户的所有域名及域名相应信息。"}
  DomainName *string `json:"domainName,omitempty" xml:"domainName,omitempty" require:"true"`
}

func (s QueryDispatchDomainsRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryDispatchDomainsRequest) GoString() string {
  return s.String()
}

func (s *QueryDispatchDomainsRequest) SetLimit(v int) *QueryDispatchDomainsRequest {
  s.Limit = &v
  return s
}

func (s *QueryDispatchDomainsRequest) SetStart(v int) *QueryDispatchDomainsRequest {
  s.Start = &v
  return s
}

func (s *QueryDispatchDomainsRequest) SetLanguage(v string) *QueryDispatchDomainsRequest {
  s.Language = &v
  return s
}

func (s *QueryDispatchDomainsRequest) SetDomainName(v string) *QueryDispatchDomainsRequest {
  s.DomainName = &v
  return s
}

type QueryDispatchDomainsResponse struct {
  // {"en":"Status code message", "zh_CN":"状态码信息"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  Content *QueryDispatchDomainsResponseContent `json:"content,omitempty" xml:"content,omitempty" require:"true" type:"Struct"`
  // {"en":"Status code", "zh_CN":"状态码"}
  ResCode *string `json:"resCode,omitempty" xml:"resCode,omitempty" require:"true"`
}

func (s QueryDispatchDomainsResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryDispatchDomainsResponse) GoString() string {
  return s.String()
}

func (s *QueryDispatchDomainsResponse) SetMsg(v string) *QueryDispatchDomainsResponse {
  s.Msg = &v
  return s
}

func (s *QueryDispatchDomainsResponse) SetContent(v *QueryDispatchDomainsResponseContent) *QueryDispatchDomainsResponse {
  s.Content = v
  return s
}

func (s *QueryDispatchDomainsResponse) SetResCode(v string) *QueryDispatchDomainsResponse {
  s.ResCode = &v
  return s
}

type QueryDispatchDomainsResponseContent struct {
  Rows []*QueryDispatchDomainsResponseContentRows `json:"rows,omitempty" xml:"rows,omitempty" require:"true" type:"Repeated"`
  // {"en":"Total count", "zh_CN":"用户域名总数量"}
  Count *int `json:"count,omitempty" xml:"count,omitempty" require:"true"`
}

func (s QueryDispatchDomainsResponseContent) String() string {
  return tea.Prettify(s)
}

func (s QueryDispatchDomainsResponseContent) GoString() string {
  return s.String()
}

func (s *QueryDispatchDomainsResponseContent) SetRows(v []*QueryDispatchDomainsResponseContentRows) *QueryDispatchDomainsResponseContent {
  s.Rows = v
  return s
}

func (s *QueryDispatchDomainsResponseContent) SetCount(v int) *QueryDispatchDomainsResponseContent {
  s.Count = &v
  return s
}

type QueryDispatchDomainsResponseContentRows struct     {
  // {"en":"Domain ID", "zh_CN":"域名ID标识"}
  DomainId *int `json:"domainId,omitempty" xml:"domainId,omitempty" require:"true"`
  // {"en":"Domain name", "zh_CN":"域名"}
  DomainName *string `json:"domainName,omitempty" xml:"domainName,omitempty" require:"true"`
  // {"en":"Policy count", "zh_CN":"域名策略数量"}
  PolicyCount *int `json:"policyCount,omitempty" xml:"policyCount,omitempty" require:"true"`
  // {"en":"Dispatch domain", "zh_CN":"调度CNAME"}
  DispatchCname *string `json:"dispatchCname,omitempty" xml:"dispatchCname,omitempty" require:"true"`
}

func (s QueryDispatchDomainsResponseContentRows) String() string {
  return tea.Prettify(s)
}

func (s QueryDispatchDomainsResponseContentRows) GoString() string {
  return s.String()
}

func (s *QueryDispatchDomainsResponseContentRows) SetDomainId(v int) *QueryDispatchDomainsResponseContentRows {
  s.DomainId = &v
  return s
}

func (s *QueryDispatchDomainsResponseContentRows) SetDomainName(v string) *QueryDispatchDomainsResponseContentRows {
  s.DomainName = &v
  return s
}

func (s *QueryDispatchDomainsResponseContentRows) SetPolicyCount(v int) *QueryDispatchDomainsResponseContentRows {
  s.PolicyCount = &v
  return s
}

func (s *QueryDispatchDomainsResponseContentRows) SetDispatchCname(v string) *QueryDispatchDomainsResponseContentRows {
  s.DispatchCname = &v
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

type RequestHeader struct {
}

func (s RequestHeader) String() string {
  return tea.Prettify(s)
}

func (s RequestHeader) GoString() string {
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


