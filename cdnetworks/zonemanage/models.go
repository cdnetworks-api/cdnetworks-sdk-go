package zonemanage

// This file is auto-generated, don't edit it. Thanks.
import (
  "github.com/alibabacloud-go/tea/tea"
)

type UpdateRecordByIDRequest struct {
  // {"en":"host name", "zh_CN":"域名头"}
  HostName *string `json:"hostName,omitempty" xml:"hostName,omitempty" require:"true"`
  // {"en":"(SOA,A,AAAA,CNAME,TXT,SRV,MX,Resolve_CNAME,RP,SPF,PTR,NS)", "zh_CN":"record类型值（SOA,A,AAAA,CNAME,TXT,SRV,MX,Resolve_CNAME,RP,SPF,PTR,NS）"}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
  // {"en":"Record value (SRV value: '${priority} ${weight} ${port} ${target}'; RP value: '${mailboxName} ${domainName}' )", "zh_CN":"记录值(SRV值为：'${priority} ${weight} ${port} ${target}'; RP值为：'${mailboxName} ${domainName}')"}
  Value *string `json:"value,omitempty" xml:"value,omitempty" require:"true"`
  // {"en":"ttl", "zh_CN":"ttl"}
  Ttl *int `json:"ttl,omitempty" xml:"ttl,omitempty" require:"true"`
  // {"en":"MX preference (Only available on MX type, default is 10)", "zh_CN":"MX类型的优先级(只有在MX类型才有, 默认为10)"}
  Data *int `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s UpdateRecordByIDRequest) String() string {
  return tea.Prettify(s)
}

func (s UpdateRecordByIDRequest) GoString() string {
  return s.String()
}

func (s *UpdateRecordByIDRequest) SetHostName(v string) *UpdateRecordByIDRequest {
  s.HostName = &v
  return s
}

func (s *UpdateRecordByIDRequest) SetType(v string) *UpdateRecordByIDRequest {
  s.Type = &v
  return s
}

func (s *UpdateRecordByIDRequest) SetValue(v string) *UpdateRecordByIDRequest {
  s.Value = &v
  return s
}

func (s *UpdateRecordByIDRequest) SetTtl(v int) *UpdateRecordByIDRequest {
  s.Ttl = &v
  return s
}

func (s *UpdateRecordByIDRequest) SetData(v int) *UpdateRecordByIDRequest {
  s.Data = &v
  return s
}

type UpdateRecordByIDResponse struct {
  // {"en":"return data", "zh_CN":"返回值"}
  Data *UpdateRecordByIDResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
  // {"en":"code,success is 0", "zh_CN":"状态码,成功为0"}
  Code *int `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"error message or Success", "zh_CN":"错误信息或Success"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s UpdateRecordByIDResponse) String() string {
  return tea.Prettify(s)
}

func (s UpdateRecordByIDResponse) GoString() string {
  return s.String()
}

func (s *UpdateRecordByIDResponse) SetData(v *UpdateRecordByIDResponseData) *UpdateRecordByIDResponse {
  s.Data = v
  return s
}

func (s *UpdateRecordByIDResponse) SetCode(v int) *UpdateRecordByIDResponse {
  s.Code = &v
  return s
}

func (s *UpdateRecordByIDResponse) SetMessage(v string) *UpdateRecordByIDResponse {
  s.Message = &v
  return s
}

type UpdateRecordByIDResponseData struct {
  // {"en":"zone id", "zh_CN":"zone的ID"}
  ZoneId *int `json:"zoneId,omitempty" xml:"zoneId,omitempty" require:"true"`
  // {"en":"record id", "zh_CN":"record的ID"}
  RecordId *int `json:"recordId,omitempty" xml:"recordId,omitempty" require:"true"`
  // {"en":"host name", "zh_CN":"域名头"}
  HostName *string `json:"hostName,omitempty" xml:"hostName,omitempty" require:"true"`
  // {"en":"(SOA,A,AAAA,CNAME,TXT,SRV,MX,Resolve_CNAME,RP,SPF,PTR,NS)", "zh_CN":"record类型值（SOA,A,AAAA,CNAME,TXT,SRV,MX,Resolve_CNAME,RP,SPF,PTR,NS）"}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
  // {"en":"Record value (RP and SRV do not have this attribute)", "zh_CN":"记录值(RP与SRV没有这个属性)"}
  Value *string `json:"value,omitempty" xml:"value,omitempty" require:"true"`
  // {"en":"ttl", "zh_CN":"ttl"}
  Ttl *int `json:"ttl,omitempty" xml:"ttl,omitempty" require:"true"`
  // {"en":"priority(Only available on SRV type)", "zh_CN":"优先(只有在SRV类型才有)"}
  Priority *int `json:"priority,omitempty" xml:"priority,omitempty" require:"true"`
  // {"en":"weight(Only available on SRV type)", "zh_CN":"权重(只有在SRV类型才有)"}
  Weight *int `json:"weight,omitempty" xml:"weight,omitempty" require:"true"`
  // {"en":"port(Only available on SRV type)", "zh_CN":"端口(只有在SRV类型才有)"}
  Port *int `json:"port,omitempty" xml:"port,omitempty" require:"true"`
  // {"en":"target domain(Only available on SRV type)", "zh_CN":"目标域名(只有在SRV类型才有)"}
  Target *string `json:"target,omitempty" xml:"target,omitempty" require:"true"`
  // {"en":"mail box name(Only available on RP type)", "zh_CN":"电子邮件域名(只有在RP类型才有)"}
  MailboxName *string `json:"mailboxName,omitempty" xml:"mailboxName,omitempty" require:"true"`
  // {"en":"domain name(Only available on RP type)", "zh_CN":"域名(只有在RP类型才有)"}
  DomainName *string `json:"domainName,omitempty" xml:"domainName,omitempty" require:"true"`
  // {"en":"preference(Only available on MX type)", "zh_CN":"优先级(只有在MX类型才有)"}
  Preference *int `json:"preference,omitempty" xml:"preference,omitempty" require:"true"`
  // {'en':'Unicode name of hostName','zh_CN':'域名头的unicode名称'}
  HostUnicode *string `json:"hostUnicode,omitempty" xml:"hostUnicode,omitempty" require:"true"`
}

func (s UpdateRecordByIDResponseData) String() string {
  return tea.Prettify(s)
}

func (s UpdateRecordByIDResponseData) GoString() string {
  return s.String()
}

func (s *UpdateRecordByIDResponseData) SetZoneId(v int) *UpdateRecordByIDResponseData {
  s.ZoneId = &v
  return s
}

func (s *UpdateRecordByIDResponseData) SetRecordId(v int) *UpdateRecordByIDResponseData {
  s.RecordId = &v
  return s
}

func (s *UpdateRecordByIDResponseData) SetHostName(v string) *UpdateRecordByIDResponseData {
  s.HostName = &v
  return s
}

func (s *UpdateRecordByIDResponseData) SetType(v string) *UpdateRecordByIDResponseData {
  s.Type = &v
  return s
}

func (s *UpdateRecordByIDResponseData) SetValue(v string) *UpdateRecordByIDResponseData {
  s.Value = &v
  return s
}

func (s *UpdateRecordByIDResponseData) SetTtl(v int) *UpdateRecordByIDResponseData {
  s.Ttl = &v
  return s
}

func (s *UpdateRecordByIDResponseData) SetPriority(v int) *UpdateRecordByIDResponseData {
  s.Priority = &v
  return s
}

func (s *UpdateRecordByIDResponseData) SetWeight(v int) *UpdateRecordByIDResponseData {
  s.Weight = &v
  return s
}

func (s *UpdateRecordByIDResponseData) SetPort(v int) *UpdateRecordByIDResponseData {
  s.Port = &v
  return s
}

func (s *UpdateRecordByIDResponseData) SetTarget(v string) *UpdateRecordByIDResponseData {
  s.Target = &v
  return s
}

func (s *UpdateRecordByIDResponseData) SetMailboxName(v string) *UpdateRecordByIDResponseData {
  s.MailboxName = &v
  return s
}

func (s *UpdateRecordByIDResponseData) SetDomainName(v string) *UpdateRecordByIDResponseData {
  s.DomainName = &v
  return s
}

func (s *UpdateRecordByIDResponseData) SetPreference(v int) *UpdateRecordByIDResponseData {
  s.Preference = &v
  return s
}

func (s *UpdateRecordByIDResponseData) SetHostUnicode(v string) *UpdateRecordByIDResponseData {
  s.HostUnicode = &v
  return s
}

type UpdateRecordByIDPaths struct {
  // {"en":"zone's ID", "zh_CN":"zone的ID"}
  ZoneId *int `json:"zoneId,omitempty" xml:"zoneId,omitempty" require:"true"`
  // {"en":"record's ID", "zh_CN":"record的ID"}
  RecordId *int `json:"recordId,omitempty" xml:"recordId,omitempty" require:"true"`
}

func (s UpdateRecordByIDPaths) String() string {
  return tea.Prettify(s)
}

func (s UpdateRecordByIDPaths) GoString() string {
  return s.String()
}

func (s *UpdateRecordByIDPaths) SetZoneId(v int) *UpdateRecordByIDPaths {
  s.ZoneId = &v
  return s
}

func (s *UpdateRecordByIDPaths) SetRecordId(v int) *UpdateRecordByIDPaths {
  s.RecordId = &v
  return s
}

type UpdateRecordByIDParameters struct {
}

func (s UpdateRecordByIDParameters) String() string {
  return tea.Prettify(s)
}

func (s UpdateRecordByIDParameters) GoString() string {
  return s.String()
}

type UpdateRecordByIDRequestHeader struct {
}

func (s UpdateRecordByIDRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s UpdateRecordByIDRequestHeader) GoString() string {
  return s.String()
}

type UpdateRecordByIDResponseHeader struct {
}

func (s UpdateRecordByIDResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s UpdateRecordByIDResponseHeader) GoString() string {
  return s.String()
}




type DeleteZTSBulkRequest struct {
  // {"en":"List of zoneIds for which ZTS configurations need to be deleted","zh_CN":"需要删除ZTS配置的zoneId列表"}
  ZoneIds []*int `json:"zoneIds,omitempty" xml:"zoneIds,omitempty" require:"true" type:"Repeated"`
}

func (s DeleteZTSBulkRequest) String() string {
  return tea.Prettify(s)
}

func (s DeleteZTSBulkRequest) GoString() string {
  return s.String()
}

func (s *DeleteZTSBulkRequest) SetZoneIds(v []*int) *DeleteZTSBulkRequest {
  s.ZoneIds = v
  return s
}

type DeleteZTSBulkRequestHeader struct {
}

func (s DeleteZTSBulkRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteZTSBulkRequestHeader) GoString() string {
  return s.String()
}

type DeleteZTSBulkPaths struct {
}

func (s DeleteZTSBulkPaths) String() string {
  return tea.Prettify(s)
}

func (s DeleteZTSBulkPaths) GoString() string {
  return s.String()
}

type DeleteZTSBulkParameters struct {
}

func (s DeleteZTSBulkParameters) String() string {
  return tea.Prettify(s)
}

func (s DeleteZTSBulkParameters) GoString() string {
  return s.String()
}

type DeleteZTSBulkResponse struct {
  // {"en":"Status code, 0 indicates success","zh_CN":"状态码,成功为0"}
  Code *int `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Error message or Success","zh_CN":"错误信息或Success"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s DeleteZTSBulkResponse) String() string {
  return tea.Prettify(s)
}

func (s DeleteZTSBulkResponse) GoString() string {
  return s.String()
}

func (s *DeleteZTSBulkResponse) SetCode(v int) *DeleteZTSBulkResponse {
  s.Code = &v
  return s
}

func (s *DeleteZTSBulkResponse) SetMessage(v string) *DeleteZTSBulkResponse {
  s.Message = &v
  return s
}

type DeleteZTSBulkResponseHeader struct {
}

func (s DeleteZTSBulkResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteZTSBulkResponseHeader) GoString() string {
  return s.String()
}




type BulkUpdateRecordsByZoneIDRequest struct {
  // {"en":"request data", "zh_CN":"请求数据"}
  Data []*BulkUpdateRecordsByZoneIDRequestData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s BulkUpdateRecordsByZoneIDRequest) String() string {
  return tea.Prettify(s)
}

func (s BulkUpdateRecordsByZoneIDRequest) GoString() string {
  return s.String()
}

func (s *BulkUpdateRecordsByZoneIDRequest) SetData(v []*BulkUpdateRecordsByZoneIDRequestData) *BulkUpdateRecordsByZoneIDRequest {
  s.Data = v
  return s
}

type BulkUpdateRecordsByZoneIDRequestData struct     {
  // {"en":"record id", "zh_CN":"record的ID"}
  RecordId *int `json:"recordId,omitempty" xml:"recordId,omitempty" require:"true"`
  // {"en":"host name", "zh_CN":"域名头"}
  HostName *string `json:"hostName,omitempty" xml:"hostName,omitempty" require:"true"`
  // {"en":"(SOA,A,AAAA,CNAME,TXT,SRV,MX,ALIAS,RP,SPF,PTR,NS)", "zh_CN":"record类型值（SOA,A,AAAA,CNAME,TXT,SRV,MX,ALIAS,RP,SPF,PTR,NS）"}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
  // {"en":"Record value (SRV value: '${priority} ${weight} ${port} ${target}'; RP value: '${mailboxName} ${domainName}' )", "zh_CN":"记录值(SRV值为：'${priority} ${weight} ${port} ${target}'; RP值为：'${mailboxName} ${domainName}')"}
  Value *string `json:"value,omitempty" xml:"value,omitempty" require:"true"`
  // {"en":"ttl", "zh_CN":"ttl"}
  Ttl *int `json:"ttl,omitempty" xml:"ttl,omitempty" require:"true"`
  // {"en":"MX preference (Only available on MX type, default is 10)", "zh_CN":"MX类型的优先级(只有在MX类型才有, 默认为10)"}
  Data *int `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s BulkUpdateRecordsByZoneIDRequestData) String() string {
  return tea.Prettify(s)
}

func (s BulkUpdateRecordsByZoneIDRequestData) GoString() string {
  return s.String()
}

func (s *BulkUpdateRecordsByZoneIDRequestData) SetRecordId(v int) *BulkUpdateRecordsByZoneIDRequestData {
  s.RecordId = &v
  return s
}

func (s *BulkUpdateRecordsByZoneIDRequestData) SetHostName(v string) *BulkUpdateRecordsByZoneIDRequestData {
  s.HostName = &v
  return s
}

func (s *BulkUpdateRecordsByZoneIDRequestData) SetType(v string) *BulkUpdateRecordsByZoneIDRequestData {
  s.Type = &v
  return s
}

func (s *BulkUpdateRecordsByZoneIDRequestData) SetValue(v string) *BulkUpdateRecordsByZoneIDRequestData {
  s.Value = &v
  return s
}

func (s *BulkUpdateRecordsByZoneIDRequestData) SetTtl(v int) *BulkUpdateRecordsByZoneIDRequestData {
  s.Ttl = &v
  return s
}

func (s *BulkUpdateRecordsByZoneIDRequestData) SetData(v int) *BulkUpdateRecordsByZoneIDRequestData {
  s.Data = &v
  return s
}

type BulkUpdateRecordsByZoneIDResponse struct {
  // {"en":"return data", "zh_CN":"返回值"}
  Data *BulkUpdateRecordsByZoneIDResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
  // {"en":"code,success is 0", "zh_CN":"状态码,成功为0"}
  Code *int `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"error message or Success", "zh_CN":"错误信息或Success"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s BulkUpdateRecordsByZoneIDResponse) String() string {
  return tea.Prettify(s)
}

func (s BulkUpdateRecordsByZoneIDResponse) GoString() string {
  return s.String()
}

func (s *BulkUpdateRecordsByZoneIDResponse) SetData(v *BulkUpdateRecordsByZoneIDResponseData) *BulkUpdateRecordsByZoneIDResponse {
  s.Data = v
  return s
}

func (s *BulkUpdateRecordsByZoneIDResponse) SetCode(v int) *BulkUpdateRecordsByZoneIDResponse {
  s.Code = &v
  return s
}

func (s *BulkUpdateRecordsByZoneIDResponse) SetMessage(v string) *BulkUpdateRecordsByZoneIDResponse {
  s.Message = &v
  return s
}

type BulkUpdateRecordsByZoneIDResponseData struct {
  // {"en":"record type(SOA,A,AAAA,CNAME,TXT,SRV,MX,ALIAS,RP,SPF,PTR,NS)", "zh_CN":"数据类型（SOA,A,AAAA,CNAME,TXT,SRV,MX,ALIAS,RP,SPF,PTR,NS）"}
  A []*BulkUpdateRecordsByZoneIDResponseDataA `json:"A,omitempty" xml:"A,omitempty" require:"true" type:"Repeated"`
}

func (s BulkUpdateRecordsByZoneIDResponseData) String() string {
  return tea.Prettify(s)
}

func (s BulkUpdateRecordsByZoneIDResponseData) GoString() string {
  return s.String()
}

func (s *BulkUpdateRecordsByZoneIDResponseData) SetA(v []*BulkUpdateRecordsByZoneIDResponseDataA) *BulkUpdateRecordsByZoneIDResponseData {
  s.A = v
  return s
}

type BulkUpdateRecordsByZoneIDResponseDataA struct     {
  // {"en":"zone id", "zh_CN":"zone的ID"}
  ZoneId *int `json:"zoneId,omitempty" xml:"zoneId,omitempty" require:"true"`
  // {"en":"record id", "zh_CN":"record的ID"}
  RecordId *int `json:"recordId,omitempty" xml:"recordId,omitempty" require:"true"`
  // {"en":"host name", "zh_CN":"域名头"}
  HostName *string `json:"hostName,omitempty" xml:"hostName,omitempty" require:"true"`
  // {"en":"record type", "zh_CN":"record类型值"}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
  // {"en":"Record value (RP and SRV do not have this attribute)", "zh_CN":"记录值(RP与SRV没有这个属性)"}
  Value *string `json:"value,omitempty" xml:"value,omitempty" require:"true"`
  // {"en":"ttl", "zh_CN":"ttl"}
  Ttl *int `json:"ttl,omitempty" xml:"ttl,omitempty" require:"true"`
  // {"en":"priority(Only available on SRV type)", "zh_CN":"优先(只有在SRV类型才有)"}
  Priority *int `json:"priority,omitempty" xml:"priority,omitempty" require:"true"`
  // {"en":"weight(Only available on SRV type)", "zh_CN":"权重(只有在SRV类型才有)"}
  Weight *int `json:"weight,omitempty" xml:"weight,omitempty" require:"true"`
  // {"en":"port(Only available on SRV type)", "zh_CN":"端口(只有在SRV类型才有)"}
  Port *int `json:"port,omitempty" xml:"port,omitempty" require:"true"`
  // {"en":"target domain(Only available on SRV type)", "zh_CN":"目标域名(只有在SRV类型才有)"}
  Target *string `json:"target,omitempty" xml:"target,omitempty" require:"true"`
  // {"en":"mail box name(Only available on RP type)", "zh_CN":"电子邮件域名(只有在RP类型才有)"}
  MailboxName *string `json:"mailboxName,omitempty" xml:"mailboxName,omitempty" require:"true"`
  // {"en":"domain name(Only available on RP type)", "zh_CN":"域名(只有在RP类型才有)"}
  DomainName *string `json:"domainName,omitempty" xml:"domainName,omitempty" require:"true"`
  // {"en":"preference(Only available on MX type)", "zh_CN":"优先级(只有在MX类型才有)"}
  Preference *int `json:"preference,omitempty" xml:"preference,omitempty" require:"true"`
  // {'en':'Unicode name of hostName','zh_CN':'域名头的unicode名称'}
  HostUnicode *string `json:"hostUnicode,omitempty" xml:"hostUnicode,omitempty" require:"true"`
}

func (s BulkUpdateRecordsByZoneIDResponseDataA) String() string {
  return tea.Prettify(s)
}

func (s BulkUpdateRecordsByZoneIDResponseDataA) GoString() string {
  return s.String()
}

func (s *BulkUpdateRecordsByZoneIDResponseDataA) SetZoneId(v int) *BulkUpdateRecordsByZoneIDResponseDataA {
  s.ZoneId = &v
  return s
}

func (s *BulkUpdateRecordsByZoneIDResponseDataA) SetRecordId(v int) *BulkUpdateRecordsByZoneIDResponseDataA {
  s.RecordId = &v
  return s
}

func (s *BulkUpdateRecordsByZoneIDResponseDataA) SetHostName(v string) *BulkUpdateRecordsByZoneIDResponseDataA {
  s.HostName = &v
  return s
}

func (s *BulkUpdateRecordsByZoneIDResponseDataA) SetType(v string) *BulkUpdateRecordsByZoneIDResponseDataA {
  s.Type = &v
  return s
}

func (s *BulkUpdateRecordsByZoneIDResponseDataA) SetValue(v string) *BulkUpdateRecordsByZoneIDResponseDataA {
  s.Value = &v
  return s
}

func (s *BulkUpdateRecordsByZoneIDResponseDataA) SetTtl(v int) *BulkUpdateRecordsByZoneIDResponseDataA {
  s.Ttl = &v
  return s
}

func (s *BulkUpdateRecordsByZoneIDResponseDataA) SetPriority(v int) *BulkUpdateRecordsByZoneIDResponseDataA {
  s.Priority = &v
  return s
}

func (s *BulkUpdateRecordsByZoneIDResponseDataA) SetWeight(v int) *BulkUpdateRecordsByZoneIDResponseDataA {
  s.Weight = &v
  return s
}

func (s *BulkUpdateRecordsByZoneIDResponseDataA) SetPort(v int) *BulkUpdateRecordsByZoneIDResponseDataA {
  s.Port = &v
  return s
}

func (s *BulkUpdateRecordsByZoneIDResponseDataA) SetTarget(v string) *BulkUpdateRecordsByZoneIDResponseDataA {
  s.Target = &v
  return s
}

func (s *BulkUpdateRecordsByZoneIDResponseDataA) SetMailboxName(v string) *BulkUpdateRecordsByZoneIDResponseDataA {
  s.MailboxName = &v
  return s
}

func (s *BulkUpdateRecordsByZoneIDResponseDataA) SetDomainName(v string) *BulkUpdateRecordsByZoneIDResponseDataA {
  s.DomainName = &v
  return s
}

func (s *BulkUpdateRecordsByZoneIDResponseDataA) SetPreference(v int) *BulkUpdateRecordsByZoneIDResponseDataA {
  s.Preference = &v
  return s
}

func (s *BulkUpdateRecordsByZoneIDResponseDataA) SetHostUnicode(v string) *BulkUpdateRecordsByZoneIDResponseDataA {
  s.HostUnicode = &v
  return s
}

type BulkUpdateRecordsByZoneIDPaths struct {
  // {"en":"zone's ID", "zh_CN":"zone的ID"}
  ZoneId *int `json:"zoneId,omitempty" xml:"zoneId,omitempty" require:"true"`
}

func (s BulkUpdateRecordsByZoneIDPaths) String() string {
  return tea.Prettify(s)
}

func (s BulkUpdateRecordsByZoneIDPaths) GoString() string {
  return s.String()
}

func (s *BulkUpdateRecordsByZoneIDPaths) SetZoneId(v int) *BulkUpdateRecordsByZoneIDPaths {
  s.ZoneId = &v
  return s
}

type BulkUpdateRecordsByZoneIDParameters struct {
}

func (s BulkUpdateRecordsByZoneIDParameters) String() string {
  return tea.Prettify(s)
}

func (s BulkUpdateRecordsByZoneIDParameters) GoString() string {
  return s.String()
}

type BulkUpdateRecordsByZoneIDRequestHeader struct {
}

func (s BulkUpdateRecordsByZoneIDRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s BulkUpdateRecordsByZoneIDRequestHeader) GoString() string {
  return s.String()
}

type BulkUpdateRecordsByZoneIDResponseHeader struct {
}

func (s BulkUpdateRecordsByZoneIDResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s BulkUpdateRecordsByZoneIDResponseHeader) GoString() string {
  return s.String()
}




type GetRecordListByZoneRequest struct {
}

func (s GetRecordListByZoneRequest) String() string {
  return tea.Prettify(s)
}

func (s GetRecordListByZoneRequest) GoString() string {
  return s.String()
}

type GetRecordListByZoneResponse struct {
  // {"en":"return data", "zh_CN":"返回值"}
  Data *GetRecordListByZoneResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
  // {"en":"code,success is 0", "zh_CN":"状态码,成功为0"}
  Code *int `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"error message or Success", "zh_CN":"错误信息或Success"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s GetRecordListByZoneResponse) String() string {
  return tea.Prettify(s)
}

func (s GetRecordListByZoneResponse) GoString() string {
  return s.String()
}

func (s *GetRecordListByZoneResponse) SetData(v *GetRecordListByZoneResponseData) *GetRecordListByZoneResponse {
  s.Data = v
  return s
}

func (s *GetRecordListByZoneResponse) SetCode(v int) *GetRecordListByZoneResponse {
  s.Code = &v
  return s
}

func (s *GetRecordListByZoneResponse) SetMessage(v string) *GetRecordListByZoneResponse {
  s.Message = &v
  return s
}

type GetRecordListByZoneResponseData struct {
  // {"en":"record type(SOA,A,AAAA,CNAME,TXT,SRV,MX,ALIAS,RP,SPF,PTR,NS)", "zh_CN":"数据类型（SOA,A,AAAA,CNAME,TXT,SRV,MX,ALIAS,RP,SPF,PTR,NS）"}
  A []*GetRecordListByZoneResponseDataA `json:"A,omitempty" xml:"A,omitempty" require:"true" type:"Repeated"`
}

func (s GetRecordListByZoneResponseData) String() string {
  return tea.Prettify(s)
}

func (s GetRecordListByZoneResponseData) GoString() string {
  return s.String()
}

func (s *GetRecordListByZoneResponseData) SetA(v []*GetRecordListByZoneResponseDataA) *GetRecordListByZoneResponseData {
  s.A = v
  return s
}

type GetRecordListByZoneResponseDataA struct     {
  // {"en":"zone id", "zh_CN":"zone的ID"}
  ZoneId *int `json:"zoneId,omitempty" xml:"zoneId,omitempty" require:"true"`
  // {"en":"record id", "zh_CN":"record的ID"}
  RecordId *int `json:"recordId,omitempty" xml:"recordId,omitempty" require:"true"`
  // {"en":"host name", "zh_CN":"域名头"}
  HostName *string `json:"hostName,omitempty" xml:"hostName,omitempty" require:"true"`
  // {"en":"record type", "zh_CN":"record类型值"}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
  // {"en":"Record value (RP and SRV do not have this attribute)", "zh_CN":"记录值(RP与SRV没有这个属性)"}
  Value *string `json:"value,omitempty" xml:"value,omitempty" require:"true"`
  // {"en":"ttl", "zh_CN":"ttl"}
  Ttl *int `json:"ttl,omitempty" xml:"ttl,omitempty" require:"true"`
  // {"en":"priority(Only available on SRV type)", "zh_CN":"优先(只有在SRV类型才有)"}
  Priority *int `json:"priority,omitempty" xml:"priority,omitempty" require:"true"`
  // {"en":"weight(Only available on SRV type)", "zh_CN":"权重(只有在SRV类型才有)"}
  Weight *int `json:"weight,omitempty" xml:"weight,omitempty" require:"true"`
  // {"en":"port(Only available on SRV type)", "zh_CN":"端口(只有在SRV类型才有)"}
  Port *int `json:"port,omitempty" xml:"port,omitempty" require:"true"`
  // {"en":"target domain(Only available on SRV type)", "zh_CN":"目标域名(只有在SRV类型才有)"}
  Target *string `json:"target,omitempty" xml:"target,omitempty" require:"true"`
  // {"en":"mail box name(Only available on RP type)", "zh_CN":"电子邮件域名(只有在RP类型才有)"}
  MailboxName *string `json:"mailboxName,omitempty" xml:"mailboxName,omitempty" require:"true"`
  // {"en":"domain name(Only available on RP type)", "zh_CN":"域名(只有在RP类型才有)"}
  DomainName *string `json:"domainName,omitempty" xml:"domainName,omitempty" require:"true"`
  // {"en":"preference(Only available on MX type)", "zh_CN":"优先级(只有在MX类型才有)"}
  Preference *int `json:"preference,omitempty" xml:"preference,omitempty" require:"true"`
  // {'en':'Unicode name of hostName','zh_CN':'域名头的unicode名称'}
  HostUnicode *string `json:"hostUnicode,omitempty" xml:"hostUnicode,omitempty" require:"true"`
}

func (s GetRecordListByZoneResponseDataA) String() string {
  return tea.Prettify(s)
}

func (s GetRecordListByZoneResponseDataA) GoString() string {
  return s.String()
}

func (s *GetRecordListByZoneResponseDataA) SetZoneId(v int) *GetRecordListByZoneResponseDataA {
  s.ZoneId = &v
  return s
}

func (s *GetRecordListByZoneResponseDataA) SetRecordId(v int) *GetRecordListByZoneResponseDataA {
  s.RecordId = &v
  return s
}

func (s *GetRecordListByZoneResponseDataA) SetHostName(v string) *GetRecordListByZoneResponseDataA {
  s.HostName = &v
  return s
}

func (s *GetRecordListByZoneResponseDataA) SetType(v string) *GetRecordListByZoneResponseDataA {
  s.Type = &v
  return s
}

func (s *GetRecordListByZoneResponseDataA) SetValue(v string) *GetRecordListByZoneResponseDataA {
  s.Value = &v
  return s
}

func (s *GetRecordListByZoneResponseDataA) SetTtl(v int) *GetRecordListByZoneResponseDataA {
  s.Ttl = &v
  return s
}

func (s *GetRecordListByZoneResponseDataA) SetPriority(v int) *GetRecordListByZoneResponseDataA {
  s.Priority = &v
  return s
}

func (s *GetRecordListByZoneResponseDataA) SetWeight(v int) *GetRecordListByZoneResponseDataA {
  s.Weight = &v
  return s
}

func (s *GetRecordListByZoneResponseDataA) SetPort(v int) *GetRecordListByZoneResponseDataA {
  s.Port = &v
  return s
}

func (s *GetRecordListByZoneResponseDataA) SetTarget(v string) *GetRecordListByZoneResponseDataA {
  s.Target = &v
  return s
}

func (s *GetRecordListByZoneResponseDataA) SetMailboxName(v string) *GetRecordListByZoneResponseDataA {
  s.MailboxName = &v
  return s
}

func (s *GetRecordListByZoneResponseDataA) SetDomainName(v string) *GetRecordListByZoneResponseDataA {
  s.DomainName = &v
  return s
}

func (s *GetRecordListByZoneResponseDataA) SetPreference(v int) *GetRecordListByZoneResponseDataA {
  s.Preference = &v
  return s
}

func (s *GetRecordListByZoneResponseDataA) SetHostUnicode(v string) *GetRecordListByZoneResponseDataA {
  s.HostUnicode = &v
  return s
}

type GetRecordListByZonePaths struct {
  // {"en":"zone's ID", "zh_CN":"zone的ID"}
  ZoneId *int `json:"zoneId,omitempty" xml:"zoneId,omitempty" require:"true"`
}

func (s GetRecordListByZonePaths) String() string {
  return tea.Prettify(s)
}

func (s GetRecordListByZonePaths) GoString() string {
  return s.String()
}

func (s *GetRecordListByZonePaths) SetZoneId(v int) *GetRecordListByZonePaths {
  s.ZoneId = &v
  return s
}

type GetRecordListByZoneParameters struct {
  // {"en":"Keywords of hostname you want to filtering", "zh_CN":"要过滤的hostname的关键字"}
  HostName *string `json:"hostName,omitempty" xml:"hostName,omitempty"`
  // {"en":"The record types you want to filter, case insensitive, separated by comma, for example: types=SOA,A,AAAA,CNAME,TXT,SRV,MX,ALIAS,RP,SPF,PTR,NS", "zh_CN":"要过滤的解析记录类型，大小写不敏感，多个类型用逗号分隔，比如：types=SOA,A,AAAA,CNAME,TXT,SRV,MX,ALIAS,RP,SPF,PTR,NS"}
  Types *string `json:"types,omitempty" xml:"types,omitempty"`
}

func (s GetRecordListByZoneParameters) String() string {
  return tea.Prettify(s)
}

func (s GetRecordListByZoneParameters) GoString() string {
  return s.String()
}

func (s *GetRecordListByZoneParameters) SetHostName(v string) *GetRecordListByZoneParameters {
  s.HostName = &v
  return s
}

func (s *GetRecordListByZoneParameters) SetTypes(v string) *GetRecordListByZoneParameters {
  s.Types = &v
  return s
}

type GetRecordListByZoneRequestHeader struct {
}

func (s GetRecordListByZoneRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetRecordListByZoneRequestHeader) GoString() string {
  return s.String()
}

type GetRecordListByZoneResponseHeader struct {
}

func (s GetRecordListByZoneResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetRecordListByZoneResponseHeader) GoString() string {
  return s.String()
}




type CreateZoneRequest struct {
  // {"en":"zone name", "zh_CN":"zone名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"contract item(eg:100000000-10)", "zh_CN":"合同编号(例:100000000-10)"}
  ContractItem *string `json:"contractItem,omitempty" xml:"contractItem,omitempty" require:"true"`
  // {"en":"default ttl", "zh_CN":"默认ttl"}
  Ttl *int `json:"ttl,omitempty" xml:"ttl,omitempty"`
}

func (s CreateZoneRequest) String() string {
  return tea.Prettify(s)
}

func (s CreateZoneRequest) GoString() string {
  return s.String()
}

func (s *CreateZoneRequest) SetName(v string) *CreateZoneRequest {
  s.Name = &v
  return s
}

func (s *CreateZoneRequest) SetContractItem(v string) *CreateZoneRequest {
  s.ContractItem = &v
  return s
}

func (s *CreateZoneRequest) SetTtl(v int) *CreateZoneRequest {
  s.Ttl = &v
  return s
}

type CreateZoneResponse struct {
  // {"en":"return data", "zh_CN":"返回值"}
  Data *CreateZoneResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
  // {"en":"code,success is 0", "zh_CN":"状态码,成功为0"}
  Code *int `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"error message or Success", "zh_CN":"错误信息或Success"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s CreateZoneResponse) String() string {
  return tea.Prettify(s)
}

func (s CreateZoneResponse) GoString() string {
  return s.String()
}

func (s *CreateZoneResponse) SetData(v *CreateZoneResponseData) *CreateZoneResponse {
  s.Data = v
  return s
}

func (s *CreateZoneResponse) SetCode(v int) *CreateZoneResponse {
  s.Code = &v
  return s
}

func (s *CreateZoneResponse) SetMessage(v string) *CreateZoneResponse {
  s.Message = &v
  return s
}

type CreateZoneResponseData struct {
  // {"en":"zone id", "zh_CN":"zone的ID"}
  ZoneId *int `json:"zoneId,omitempty" xml:"zoneId,omitempty" require:"true"`
  // {"en":"zone name", "zh_CN":"域名"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"Serial Number", "zh_CN":"序列号"}
  Serial *int `json:"serial,omitempty" xml:"serial,omitempty" require:"true"`
  // {"en":"TTL", "zh_CN":"TTL"}
  Ttl *int `json:"ttl,omitempty" xml:"ttl,omitempty" require:"true"`
  // {"en":"STATUS CODE ( 0: NEW OR MODIFY ; 1:PENDING TO STAGE; 2: IN STAGE; 3: PENDING TO PRODUCTION;4: IN PRODUCTION; -2: ERROR TO STAGE; -4: ERROR TO PRODUCTION )", "zh_CN":"状态码（0: NEW; 1:PENDING TO STAGE; 2: IN STAGE; 3: PENDING TO PRODUCTION;4: IN PRODUCTION; -2: ERROR TO STAGE; -4: ERROR TO PRODUCTION）"}
  StatusCode *int `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
  // {"en":"status code description", "zh_CN":"状态码说明"}
  StatusText *string `json:"statusText,omitempty" xml:"statusText,omitempty" require:"true"`
  // {"en":"Creation time, UTC time(eg: 2012-11-12T02:08:34Z )", "zh_CN":"创建时间，UTC时间(如: 2012-11-12T02:08:34Z )"}
  DateCreated *string `json:"dateCreated,omitempty" xml:"dateCreated,omitempty" require:"true"`
  // {"en":"Modified time, UTC time(eg: 2012-11-12T02:08:34Z )", "zh_CN":"修改时间，UTC时间(如: 2012-11-12T02:08:34Z )"}
  DateModified *string `json:"dateModified,omitempty" xml:"dateModified,omitempty" require:"true"`
  // {"en":"Published time, UTC time(eg: 2012-11-12T02:08:34Z )", "zh_CN":"发布时间，UTC时间(如: 2012-11-12T02:08:34Z )"}
  DatePublished *string `json:"datePublished,omitempty" xml:"datePublished,omitempty" require:"true"`
  // {'en':'product','zh_CN':'产品'}
  Product *string `json:"product,omitempty" xml:"product,omitempty" require:"true"`
  // {'en':'The unicode name of the zone','zh_CN':'zone的unicode名称'}
  Unicode *string `json:"unicode,omitempty" xml:"unicode,omitempty" require:"true"`
  // {'en':'customer info','zh_CN':'客户信息'}
  Customer *CreateZoneResponseDataCustomer `json:"customer,omitempty" xml:"customer,omitempty" require:"true" type:"Struct"`
}

func (s CreateZoneResponseData) String() string {
  return tea.Prettify(s)
}

func (s CreateZoneResponseData) GoString() string {
  return s.String()
}

func (s *CreateZoneResponseData) SetZoneId(v int) *CreateZoneResponseData {
  s.ZoneId = &v
  return s
}

func (s *CreateZoneResponseData) SetName(v string) *CreateZoneResponseData {
  s.Name = &v
  return s
}

func (s *CreateZoneResponseData) SetSerial(v int) *CreateZoneResponseData {
  s.Serial = &v
  return s
}

func (s *CreateZoneResponseData) SetTtl(v int) *CreateZoneResponseData {
  s.Ttl = &v
  return s
}

func (s *CreateZoneResponseData) SetStatusCode(v int) *CreateZoneResponseData {
  s.StatusCode = &v
  return s
}

func (s *CreateZoneResponseData) SetStatusText(v string) *CreateZoneResponseData {
  s.StatusText = &v
  return s
}

func (s *CreateZoneResponseData) SetDateCreated(v string) *CreateZoneResponseData {
  s.DateCreated = &v
  return s
}

func (s *CreateZoneResponseData) SetDateModified(v string) *CreateZoneResponseData {
  s.DateModified = &v
  return s
}

func (s *CreateZoneResponseData) SetDatePublished(v string) *CreateZoneResponseData {
  s.DatePublished = &v
  return s
}

func (s *CreateZoneResponseData) SetProduct(v string) *CreateZoneResponseData {
  s.Product = &v
  return s
}

func (s *CreateZoneResponseData) SetUnicode(v string) *CreateZoneResponseData {
  s.Unicode = &v
  return s
}

func (s *CreateZoneResponseData) SetCustomer(v *CreateZoneResponseDataCustomer) *CreateZoneResponseData {
  s.Customer = v
  return s
}

type CreateZoneResponseDataCustomer struct {
  // {'en':'customer name','zh_CN':'客户名称'}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
}

func (s CreateZoneResponseDataCustomer) String() string {
  return tea.Prettify(s)
}

func (s CreateZoneResponseDataCustomer) GoString() string {
  return s.String()
}

func (s *CreateZoneResponseDataCustomer) SetName(v string) *CreateZoneResponseDataCustomer {
  s.Name = &v
  return s
}

type CreateZonePaths struct {
}

func (s CreateZonePaths) String() string {
  return tea.Prettify(s)
}

func (s CreateZonePaths) GoString() string {
  return s.String()
}

type CreateZoneParameters struct {
  // {"en":"Whether to restore deleted zone data, the default value is -1. When there is deleted zone data before the newly created zone, if the value is -1, an error 409 will be reported; if the value is 0, the old zone will be restored; if the value is 1, a new zone will be created and the deleted zone in the database will be completely deleted.", "zh_CN":"是否恢复已删除zone数据判断，默认值为-1。当新建的zone之前有已删除的zone数据时，值为-1时,则报错409; 值为0时，则恢复旧zone; 值为1时，新建新zone，彻底删除数据库中已删除zone。"}
  RestoreFlag *int `json:"restoreFlag,omitempty" xml:"restoreFlag,omitempty"`
}

func (s CreateZoneParameters) String() string {
  return tea.Prettify(s)
}

func (s CreateZoneParameters) GoString() string {
  return s.String()
}

func (s *CreateZoneParameters) SetRestoreFlag(v int) *CreateZoneParameters {
  s.RestoreFlag = &v
  return s
}

type CreateZoneRequestHeader struct {
}

func (s CreateZoneRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s CreateZoneRequestHeader) GoString() string {
  return s.String()
}

type CreateZoneResponseHeader struct {
}

func (s CreateZoneResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s CreateZoneResponseHeader) GoString() string {
  return s.String()
}




type DeployZoneConfigRequest struct {
}

func (s DeployZoneConfigRequest) String() string {
  return tea.Prettify(s)
}

func (s DeployZoneConfigRequest) GoString() string {
  return s.String()
}

type DeployZoneConfigResponse struct {
  // {"en":"return data", "zh_CN":"返回值"}
  Data *DeployZoneConfigResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
  // {"en":"code,success is 0", "zh_CN":"状态码,成功为0"}
  Code *int `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"error message or Success", "zh_CN":"错误信息或Success"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s DeployZoneConfigResponse) String() string {
  return tea.Prettify(s)
}

func (s DeployZoneConfigResponse) GoString() string {
  return s.String()
}

func (s *DeployZoneConfigResponse) SetData(v *DeployZoneConfigResponseData) *DeployZoneConfigResponse {
  s.Data = v
  return s
}

func (s *DeployZoneConfigResponse) SetCode(v int) *DeployZoneConfigResponse {
  s.Code = &v
  return s
}

func (s *DeployZoneConfigResponse) SetMessage(v string) *DeployZoneConfigResponse {
  s.Message = &v
  return s
}

type DeployZoneConfigResponseData struct {
  // {"en":"zone's ID", "zh_CN":"zone的ID"}
  ZoneId *int `json:"zoneId,omitempty" xml:"zoneId,omitempty" require:"true"`
  // {"en":"status description", "zh_CN":"状态说明"}
  Phase *string `json:"phase,omitempty" xml:"phase,omitempty" require:"true"`
}

func (s DeployZoneConfigResponseData) String() string {
  return tea.Prettify(s)
}

func (s DeployZoneConfigResponseData) GoString() string {
  return s.String()
}

func (s *DeployZoneConfigResponseData) SetZoneId(v int) *DeployZoneConfigResponseData {
  s.ZoneId = &v
  return s
}

func (s *DeployZoneConfigResponseData) SetPhase(v string) *DeployZoneConfigResponseData {
  s.Phase = &v
  return s
}

type DeployZoneConfigPaths struct {
  // {"en":"zone's ID", "zh_CN":"zone的ID"}
  ZoneId *int `json:"zoneId,omitempty" xml:"zoneId,omitempty" require:"true"`
}

func (s DeployZoneConfigPaths) String() string {
  return tea.Prettify(s)
}

func (s DeployZoneConfigPaths) GoString() string {
  return s.String()
}

func (s *DeployZoneConfigPaths) SetZoneId(v int) *DeployZoneConfigPaths {
  s.ZoneId = &v
  return s
}

type DeployZoneConfigParameters struct {
  // {"en":"The status that needs to be deployed, the value can be staging or production", "zh_CN":"需要部署的状态，值可以为staging或者production"}
  DeployType *string `json:"deployType,omitempty" xml:"deployType,omitempty"`
}

func (s DeployZoneConfigParameters) String() string {
  return tea.Prettify(s)
}

func (s DeployZoneConfigParameters) GoString() string {
  return s.String()
}

func (s *DeployZoneConfigParameters) SetDeployType(v string) *DeployZoneConfigParameters {
  s.DeployType = &v
  return s
}

type DeployZoneConfigRequestHeader struct {
}

func (s DeployZoneConfigRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s DeployZoneConfigRequestHeader) GoString() string {
  return s.String()
}

type DeployZoneConfigResponseHeader struct {
}

func (s DeployZoneConfigResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s DeployZoneConfigResponseHeader) GoString() string {
  return s.String()
}




type GetZoneByIDRequest struct {
}

func (s GetZoneByIDRequest) String() string {
  return tea.Prettify(s)
}

func (s GetZoneByIDRequest) GoString() string {
  return s.String()
}

type GetZoneByIDResponse struct {
  // {"en":"result data", "zh_CN":"返回数据"}
  Data *GetZoneByIDResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
  // {"en":"code,success is 0", "zh_CN":"状态码,成功为0"}
  Code *int `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"error message or Success", "zh_CN":"错误信息或Success"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s GetZoneByIDResponse) String() string {
  return tea.Prettify(s)
}

func (s GetZoneByIDResponse) GoString() string {
  return s.String()
}

func (s *GetZoneByIDResponse) SetData(v *GetZoneByIDResponseData) *GetZoneByIDResponse {
  s.Data = v
  return s
}

func (s *GetZoneByIDResponse) SetCode(v int) *GetZoneByIDResponse {
  s.Code = &v
  return s
}

func (s *GetZoneByIDResponse) SetMessage(v string) *GetZoneByIDResponse {
  s.Message = &v
  return s
}

type GetZoneByIDResponseData struct {
  // {"en":"zone id", "zh_CN":"zone的ID"}
  ZoneId *int `json:"zoneId,omitempty" xml:"zoneId,omitempty" require:"true"`
  // {"en":"zone name", "zh_CN":"域名"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"Serial Number", "zh_CN":"序列号"}
  Serial *int `json:"serial,omitempty" xml:"serial,omitempty" require:"true"`
  // {"en":"TTL", "zh_CN":"TTL"}
  Ttl *int `json:"ttl,omitempty" xml:"ttl,omitempty" require:"true"`
  // {"en":"STATUS CODE ( 0: NEW OR MODIFY ; 1:PENDING TO STAGE; 2: IN STAGE; 3: PENDING TO PRODUCTION;4: IN PRODUCTION; -2: ERROR TO STAGE; -4: ERROR TO PRODUCTION )", "zh_CN":"状态码（0: NEW; 1:PENDING TO STAGE; 2: IN STAGE; 3: PENDING TO PRODUCTION;4: IN PRODUCTION; -2: ERROR TO STAGE; -4: ERROR TO PRODUCTION）"}
  StatusCode *int `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
  // {"en":"status code description", "zh_CN":"状态码说明"}
  StatusText *string `json:"statusText,omitempty" xml:"statusText,omitempty" require:"true"`
  // {"en":"Creation time, UTC time(eg: 2012-11-12T02:08:34Z )", "zh_CN":"创建时间，UTC时间(如: 2012-11-12T02:08:34Z )"}
  DateCreated *string `json:"dateCreated,omitempty" xml:"dateCreated,omitempty" require:"true"`
  // {"en":"Modified time, UTC time(eg: 2012-11-12T02:08:34Z )", "zh_CN":"修改时间，UTC时间(如: 2012-11-12T02:08:34Z )"}
  DateModified *string `json:"dateModified,omitempty" xml:"dateModified,omitempty" require:"true"`
  // {"en":"Published time, UTC time(eg: 2012-11-12T02:08:34Z )", "zh_CN":"发布时间，UTC时间(如: 2012-11-12T02:08:34Z )"}
  DatePublished *string `json:"datePublished,omitempty" xml:"datePublished,omitempty" require:"true"`
  // {'en':'product','zh_CN':'产品'}
  Product *string `json:"product,omitempty" xml:"product,omitempty" require:"true"`
  // {'en':'The unicode name of the zone','zh_CN':'zone的unicode名称'}
  Unicode *string `json:"unicode,omitempty" xml:"unicode,omitempty" require:"true"`
  // {'en':'customer info','zh_CN':'客户信息'}
  Customer *GetZoneByIDResponseDataCustomer `json:"customer,omitempty" xml:"customer,omitempty" require:"true" type:"Struct"`
}

func (s GetZoneByIDResponseData) String() string {
  return tea.Prettify(s)
}

func (s GetZoneByIDResponseData) GoString() string {
  return s.String()
}

func (s *GetZoneByIDResponseData) SetZoneId(v int) *GetZoneByIDResponseData {
  s.ZoneId = &v
  return s
}

func (s *GetZoneByIDResponseData) SetName(v string) *GetZoneByIDResponseData {
  s.Name = &v
  return s
}

func (s *GetZoneByIDResponseData) SetSerial(v int) *GetZoneByIDResponseData {
  s.Serial = &v
  return s
}

func (s *GetZoneByIDResponseData) SetTtl(v int) *GetZoneByIDResponseData {
  s.Ttl = &v
  return s
}

func (s *GetZoneByIDResponseData) SetStatusCode(v int) *GetZoneByIDResponseData {
  s.StatusCode = &v
  return s
}

func (s *GetZoneByIDResponseData) SetStatusText(v string) *GetZoneByIDResponseData {
  s.StatusText = &v
  return s
}

func (s *GetZoneByIDResponseData) SetDateCreated(v string) *GetZoneByIDResponseData {
  s.DateCreated = &v
  return s
}

func (s *GetZoneByIDResponseData) SetDateModified(v string) *GetZoneByIDResponseData {
  s.DateModified = &v
  return s
}

func (s *GetZoneByIDResponseData) SetDatePublished(v string) *GetZoneByIDResponseData {
  s.DatePublished = &v
  return s
}

func (s *GetZoneByIDResponseData) SetProduct(v string) *GetZoneByIDResponseData {
  s.Product = &v
  return s
}

func (s *GetZoneByIDResponseData) SetUnicode(v string) *GetZoneByIDResponseData {
  s.Unicode = &v
  return s
}

func (s *GetZoneByIDResponseData) SetCustomer(v *GetZoneByIDResponseDataCustomer) *GetZoneByIDResponseData {
  s.Customer = v
  return s
}

type GetZoneByIDResponseDataCustomer struct {
  // {'en':'customer name','zh_CN':'客户名称'}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
}

func (s GetZoneByIDResponseDataCustomer) String() string {
  return tea.Prettify(s)
}

func (s GetZoneByIDResponseDataCustomer) GoString() string {
  return s.String()
}

func (s *GetZoneByIDResponseDataCustomer) SetName(v string) *GetZoneByIDResponseDataCustomer {
  s.Name = &v
  return s
}

type GetZoneByIDPaths struct {
  // {"en":"zone's ID", "zh_CN":"zone的ID"}
  ZoneId *int `json:"zoneId,omitempty" xml:"zoneId,omitempty" require:"true"`
}

func (s GetZoneByIDPaths) String() string {
  return tea.Prettify(s)
}

func (s GetZoneByIDPaths) GoString() string {
  return s.String()
}

func (s *GetZoneByIDPaths) SetZoneId(v int) *GetZoneByIDPaths {
  s.ZoneId = &v
  return s
}

type GetZoneByIDParameters struct {
}

func (s GetZoneByIDParameters) String() string {
  return tea.Prettify(s)
}

func (s GetZoneByIDParameters) GoString() string {
  return s.String()
}

type GetZoneByIDRequestHeader struct {
}

func (s GetZoneByIDRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetZoneByIDRequestHeader) GoString() string {
  return s.String()
}

type GetZoneByIDResponseHeader struct {
}

func (s GetZoneByIDResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetZoneByIDResponseHeader) GoString() string {
  return s.String()
}




type BatchCreateZoneRequest struct {
  // {"en":"zone name, comma split, up to 20 can be created at the same time", "zh_CN":"zone名称，逗号分割，最多可支持同时创建20个"}
  Names *string `json:"names,omitempty" xml:"names,omitempty" require:"true"`
  // {"en":"contract item(eg:100000000-10)", "zh_CN":"合同编号(例:100000000-10)"}
  ContractItem *string `json:"contractItem,omitempty" xml:"contractItem,omitempty" require:"true"`
  // {"en":"default ttl", "zh_CN":"默认ttl"}
  Ttl *int `json:"ttl,omitempty" xml:"ttl,omitempty"`
}

func (s BatchCreateZoneRequest) String() string {
  return tea.Prettify(s)
}

func (s BatchCreateZoneRequest) GoString() string {
  return s.String()
}

func (s *BatchCreateZoneRequest) SetNames(v string) *BatchCreateZoneRequest {
  s.Names = &v
  return s
}

func (s *BatchCreateZoneRequest) SetContractItem(v string) *BatchCreateZoneRequest {
  s.ContractItem = &v
  return s
}

func (s *BatchCreateZoneRequest) SetTtl(v int) *BatchCreateZoneRequest {
  s.Ttl = &v
  return s
}

type BatchCreateZoneResponse struct {
  // {"en":"return data", "zh_CN":"返回值"}
  Data *BatchCreateZoneResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
  // {"en":"code,success is 0", "zh_CN":"状态码,成功为0"}
  Code *int `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"error message or Success", "zh_CN":"错误信息或Success"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s BatchCreateZoneResponse) String() string {
  return tea.Prettify(s)
}

func (s BatchCreateZoneResponse) GoString() string {
  return s.String()
}

func (s *BatchCreateZoneResponse) SetData(v *BatchCreateZoneResponseData) *BatchCreateZoneResponse {
  s.Data = v
  return s
}

func (s *BatchCreateZoneResponse) SetCode(v int) *BatchCreateZoneResponse {
  s.Code = &v
  return s
}

func (s *BatchCreateZoneResponse) SetMessage(v string) *BatchCreateZoneResponse {
  s.Message = &v
  return s
}

type BatchCreateZoneResponseData struct {
  // {"en":"Object or string type. If it fails, a string will be returned. The zone information is returned if succeeds", "zh_CN": "对象或者字符串类型，如果失败将会返回字符串。成功时返回zone信息"}
  K *BatchCreateZoneResponseDataK `json:"K,omitempty" xml:"K,omitempty" require:"true" type:"Struct"`
}

func (s BatchCreateZoneResponseData) String() string {
  return tea.Prettify(s)
}

func (s BatchCreateZoneResponseData) GoString() string {
  return s.String()
}

func (s *BatchCreateZoneResponseData) SetK(v *BatchCreateZoneResponseDataK) *BatchCreateZoneResponseData {
  s.K = v
  return s
}

type BatchCreateZoneResponseDataK struct {
  // {"en":"zone id", "zh_CN":"zone的ID"}
  ZoneId *int `json:"zoneId,omitempty" xml:"zoneId,omitempty" require:"true"`
  // {"en":"zone name", "zh_CN":"域名"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"Serial Number", "zh_CN":"序列号"}
  Serial *int `json:"serial,omitempty" xml:"serial,omitempty" require:"true"`
  // {"en":"TTL", "zh_CN":"TTL"}
  Ttl *int `json:"ttl,omitempty" xml:"ttl,omitempty" require:"true"`
  // {"en":"STATUS CODE ( 0: NEW OR MODIFY ; 1:PENDING TO STAGE; 2: IN STAGE; 3: PENDING TO PRODUCTION;4: IN PRODUCTION; -2: ERROR TO STAGE; -4: ERROR TO PRODUCTION )", "zh_CN":"状态码（0: NEW; 1:PENDING TO STAGE; 2: IN STAGE; 3: PENDING TO PRODUCTION;4: IN PRODUCTION; -2: ERROR TO STAGE; -4: ERROR TO PRODUCTION）"}
  StatusCode *int `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
  // {"en":"status code description", "zh_CN":"状态码说明"}
  StatusText *string `json:"statusText,omitempty" xml:"statusText,omitempty" require:"true"`
  // {"en":"Creation time, UTC time(eg: 2012-11-12T02:08:34Z )", "zh_CN":"创建时间，UTC时间(如: 2012-11-12T02:08:34Z )"}
  DateCreated *string `json:"dateCreated,omitempty" xml:"dateCreated,omitempty" require:"true"`
  // {"en":"Modified time, UTC time(eg: 2012-11-12T02:08:34Z )", "zh_CN":"修改时间，UTC时间(如: 2012-11-12T02:08:34Z )"}
  DateModified *string `json:"dateModified,omitempty" xml:"dateModified,omitempty" require:"true"`
  // {"en":"Published time, UTC time(eg: 2012-11-12T02:08:34Z )", "zh_CN":"发布时间，UTC时间(如: 2012-11-12T02:08:34Z )"}
  DatePublished *string `json:"datePublished,omitempty" xml:"datePublished,omitempty" require:"true"`
  // {'en':'product','zh_CN':'产品'}
  Product *string `json:"product,omitempty" xml:"product,omitempty" require:"true"`
  // {'en':'The unicode name of the zone','zh_CN':'zone的unicode名称'}
  Unicode *string `json:"unicode,omitempty" xml:"unicode,omitempty" require:"true"`
  // {'en':'customer info','zh_CN':'客户信息'}
  Customer *BatchCreateZoneResponseDataKCustomer `json:"customer,omitempty" xml:"customer,omitempty" require:"true" type:"Struct"`
}

func (s BatchCreateZoneResponseDataK) String() string {
  return tea.Prettify(s)
}

func (s BatchCreateZoneResponseDataK) GoString() string {
  return s.String()
}

func (s *BatchCreateZoneResponseDataK) SetZoneId(v int) *BatchCreateZoneResponseDataK {
  s.ZoneId = &v
  return s
}

func (s *BatchCreateZoneResponseDataK) SetName(v string) *BatchCreateZoneResponseDataK {
  s.Name = &v
  return s
}

func (s *BatchCreateZoneResponseDataK) SetSerial(v int) *BatchCreateZoneResponseDataK {
  s.Serial = &v
  return s
}

func (s *BatchCreateZoneResponseDataK) SetTtl(v int) *BatchCreateZoneResponseDataK {
  s.Ttl = &v
  return s
}

func (s *BatchCreateZoneResponseDataK) SetStatusCode(v int) *BatchCreateZoneResponseDataK {
  s.StatusCode = &v
  return s
}

func (s *BatchCreateZoneResponseDataK) SetStatusText(v string) *BatchCreateZoneResponseDataK {
  s.StatusText = &v
  return s
}

func (s *BatchCreateZoneResponseDataK) SetDateCreated(v string) *BatchCreateZoneResponseDataK {
  s.DateCreated = &v
  return s
}

func (s *BatchCreateZoneResponseDataK) SetDateModified(v string) *BatchCreateZoneResponseDataK {
  s.DateModified = &v
  return s
}

func (s *BatchCreateZoneResponseDataK) SetDatePublished(v string) *BatchCreateZoneResponseDataK {
  s.DatePublished = &v
  return s
}

func (s *BatchCreateZoneResponseDataK) SetProduct(v string) *BatchCreateZoneResponseDataK {
  s.Product = &v
  return s
}

func (s *BatchCreateZoneResponseDataK) SetUnicode(v string) *BatchCreateZoneResponseDataK {
  s.Unicode = &v
  return s
}

func (s *BatchCreateZoneResponseDataK) SetCustomer(v *BatchCreateZoneResponseDataKCustomer) *BatchCreateZoneResponseDataK {
  s.Customer = v
  return s
}

type BatchCreateZoneResponseDataKCustomer struct {
  // {'en':'customer name','zh_CN':'客户名称'}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
}

func (s BatchCreateZoneResponseDataKCustomer) String() string {
  return tea.Prettify(s)
}

func (s BatchCreateZoneResponseDataKCustomer) GoString() string {
  return s.String()
}

func (s *BatchCreateZoneResponseDataKCustomer) SetName(v string) *BatchCreateZoneResponseDataKCustomer {
  s.Name = &v
  return s
}

type BatchCreateZonePaths struct {
}

func (s BatchCreateZonePaths) String() string {
  return tea.Prettify(s)
}

func (s BatchCreateZonePaths) GoString() string {
  return s.String()
}

type BatchCreateZoneParameters struct {
  // {"en":"Whether to restore deleted zone data, the default value is -1. When there is deleted zone data before the newly created zone, if the value is -1, an error 409 will be reported; if the value is 0, the old zone will be restored; if the value is 1, a new zone will be created and the deleted zone in the database will be completely deleted.", "zh_CN":"是否恢复已删除zone数据判断，默认值为-1。当新建的zone之前有已删除的zone数据时，值为-1时,则报错409; 值为0时，则恢复旧zone; 值为1时，新建新zone，彻底删除数据库中已删除zone。"}
  RestoreFlag *int `json:"restoreFlag,omitempty" xml:"restoreFlag,omitempty"`
}

func (s BatchCreateZoneParameters) String() string {
  return tea.Prettify(s)
}

func (s BatchCreateZoneParameters) GoString() string {
  return s.String()
}

func (s *BatchCreateZoneParameters) SetRestoreFlag(v int) *BatchCreateZoneParameters {
  s.RestoreFlag = &v
  return s
}

type BatchCreateZoneRequestHeader struct {
}

func (s BatchCreateZoneRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s BatchCreateZoneRequestHeader) GoString() string {
  return s.String()
}

type BatchCreateZoneResponseHeader struct {
}

func (s BatchCreateZoneResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s BatchCreateZoneResponseHeader) GoString() string {
  return s.String()
}




type GetZoneListRequest struct {
}

func (s GetZoneListRequest) String() string {
  return tea.Prettify(s)
}

func (s GetZoneListRequest) GoString() string {
  return s.String()
}

type GetZoneListResponse struct {
  // {"en":"result data", "zh_CN":"返回数据"}
  Data *GetZoneListResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
  // {"en":"code,success is 0", "zh_CN":"状态码,成功为0"}
  Code *int `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"error message or Success", "zh_CN":"错误信息或Success"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s GetZoneListResponse) String() string {
  return tea.Prettify(s)
}

func (s GetZoneListResponse) GoString() string {
  return s.String()
}

func (s *GetZoneListResponse) SetData(v *GetZoneListResponseData) *GetZoneListResponse {
  s.Data = v
  return s
}

func (s *GetZoneListResponse) SetCode(v int) *GetZoneListResponse {
  s.Code = &v
  return s
}

func (s *GetZoneListResponse) SetMessage(v string) *GetZoneListResponse {
  s.Message = &v
  return s
}

type GetZoneListResponseData struct {
  // {"en":"Pagination info", "zh_CN":"分页信息"}
  PageInfo *GetZoneListResponseDataPageInfo `json:"pageInfo,omitempty" xml:"pageInfo,omitempty" require:"true" type:"Struct"`
  // {"en":"return data", "zh_CN":"返回值"}
  Results []*GetZoneListResponseDataResults `json:"results,omitempty" xml:"results,omitempty" require:"true" type:"Repeated"`
}

func (s GetZoneListResponseData) String() string {
  return tea.Prettify(s)
}

func (s GetZoneListResponseData) GoString() string {
  return s.String()
}

func (s *GetZoneListResponseData) SetPageInfo(v *GetZoneListResponseDataPageInfo) *GetZoneListResponseData {
  s.PageInfo = v
  return s
}

func (s *GetZoneListResponseData) SetResults(v []*GetZoneListResponseDataResults) *GetZoneListResponseData {
  s.Results = v
  return s
}

type GetZoneListResponseDataPageInfo struct {
  // {"en":"total pages", "zh_CN":"总页数"}
  Pages *int `json:"pages,omitempty" xml:"pages,omitempty" require:"true"`
  // {"en":"total count", "zh_CN":"总数量"}
  Count *int `json:"count,omitempty" xml:"count,omitempty" require:"true"`
  // {"en":"num of per page", "zh_CN":"每页数量"}
  PageSize *int `json:"pageSize,omitempty" xml:"pageSize,omitempty" require:"true"`
  // {"en":"current page", "zh_CN":"当前页"}
  Page *int `json:"page,omitempty" xml:"page,omitempty" require:"true"`
}

func (s GetZoneListResponseDataPageInfo) String() string {
  return tea.Prettify(s)
}

func (s GetZoneListResponseDataPageInfo) GoString() string {
  return s.String()
}

func (s *GetZoneListResponseDataPageInfo) SetPages(v int) *GetZoneListResponseDataPageInfo {
  s.Pages = &v
  return s
}

func (s *GetZoneListResponseDataPageInfo) SetCount(v int) *GetZoneListResponseDataPageInfo {
  s.Count = &v
  return s
}

func (s *GetZoneListResponseDataPageInfo) SetPageSize(v int) *GetZoneListResponseDataPageInfo {
  s.PageSize = &v
  return s
}

func (s *GetZoneListResponseDataPageInfo) SetPage(v int) *GetZoneListResponseDataPageInfo {
  s.Page = &v
  return s
}

type GetZoneListResponseDataResults struct     {
  // {"en":"zone id", "zh_CN":"zone的ID"}
  ZoneId *int `json:"zoneId,omitempty" xml:"zoneId,omitempty" require:"true"`
  // {"en":"zone name", "zh_CN":"域名"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"Serial Number", "zh_CN":"序列号"}
  Serial *int `json:"serial,omitempty" xml:"serial,omitempty" require:"true"`
  // {"en":"TTL", "zh_CN":"TTL"}
  Ttl *int `json:"ttl,omitempty" xml:"ttl,omitempty" require:"true"`
  // {"en":"STATUS CODE ( 0: NEW OR MODIFY ; 1:PENDING TO STAGE; 2: IN STAGE; 3: PENDING TO PRODUCTION;4: IN PRODUCTION; -2: ERROR TO STAGE; -4: ERROR TO PRODUCTION )", "zh_CN":"状态码（0: NEW; 1:PENDING TO STAGE; 2: IN STAGE; 3: PENDING TO PRODUCTION;4: IN PRODUCTION; -2: ERROR TO STAGE; -4: ERROR TO PRODUCTION）"}
  StatusCode *int `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
  // {"en":"status code description", "zh_CN":"状态码说明"}
  StatusText *string `json:"statusText,omitempty" xml:"statusText,omitempty" require:"true"`
  // {"en":"Creation time, UTC time(eg: 2012-11-12T02:08:34Z )", "zh_CN":"创建时间，UTC时间(如: 2012-11-12T02:08:34Z )"}
  DateCreated *string `json:"dateCreated,omitempty" xml:"dateCreated,omitempty" require:"true"`
  // {"en":"Modified time, UTC time(eg: 2012-11-12T02:08:34Z )", "zh_CN":"修改时间，UTC时间(如: 2012-11-12T02:08:34Z )"}
  DateModified *string `json:"dateModified,omitempty" xml:"dateModified,omitempty" require:"true"`
  // {"en":"Published time, UTC time(eg: 2012-11-12T02:08:34Z )", "zh_CN":"发布时间，UTC时间(如: 2012-11-12T02:08:34Z )"}
  DatePublished *string `json:"datePublished,omitempty" xml:"datePublished,omitempty" require:"true"`
  // {'en':'product','zh_CN':'产品'}
  Product *string `json:"product,omitempty" xml:"product,omitempty" require:"true"`
  // {"en":"The unicode name of the zone","zh_CN":"zone的unicode名称"}
  Unicode *string `json:"unicode,omitempty" xml:"unicode,omitempty" require:"true"`
  // {'en':'customer info','zh_CN':'客户信息'}
  Customer *GetZoneListResponseDataResultsCustomer `json:"customer,omitempty" xml:"customer,omitempty" require:"true" type:"Struct"`
}

func (s GetZoneListResponseDataResults) String() string {
  return tea.Prettify(s)
}

func (s GetZoneListResponseDataResults) GoString() string {
  return s.String()
}

func (s *GetZoneListResponseDataResults) SetZoneId(v int) *GetZoneListResponseDataResults {
  s.ZoneId = &v
  return s
}

func (s *GetZoneListResponseDataResults) SetName(v string) *GetZoneListResponseDataResults {
  s.Name = &v
  return s
}

func (s *GetZoneListResponseDataResults) SetSerial(v int) *GetZoneListResponseDataResults {
  s.Serial = &v
  return s
}

func (s *GetZoneListResponseDataResults) SetTtl(v int) *GetZoneListResponseDataResults {
  s.Ttl = &v
  return s
}

func (s *GetZoneListResponseDataResults) SetStatusCode(v int) *GetZoneListResponseDataResults {
  s.StatusCode = &v
  return s
}

func (s *GetZoneListResponseDataResults) SetStatusText(v string) *GetZoneListResponseDataResults {
  s.StatusText = &v
  return s
}

func (s *GetZoneListResponseDataResults) SetDateCreated(v string) *GetZoneListResponseDataResults {
  s.DateCreated = &v
  return s
}

func (s *GetZoneListResponseDataResults) SetDateModified(v string) *GetZoneListResponseDataResults {
  s.DateModified = &v
  return s
}

func (s *GetZoneListResponseDataResults) SetDatePublished(v string) *GetZoneListResponseDataResults {
  s.DatePublished = &v
  return s
}

func (s *GetZoneListResponseDataResults) SetProduct(v string) *GetZoneListResponseDataResults {
  s.Product = &v
  return s
}

func (s *GetZoneListResponseDataResults) SetUnicode(v string) *GetZoneListResponseDataResults {
  s.Unicode = &v
  return s
}

func (s *GetZoneListResponseDataResults) SetCustomer(v *GetZoneListResponseDataResultsCustomer) *GetZoneListResponseDataResults {
  s.Customer = v
  return s
}

type GetZoneListResponseDataResultsCustomer struct {
  // {'en':'customer name','zh_CN':'客户名称'}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
}

func (s GetZoneListResponseDataResultsCustomer) String() string {
  return tea.Prettify(s)
}

func (s GetZoneListResponseDataResultsCustomer) GoString() string {
  return s.String()
}

func (s *GetZoneListResponseDataResultsCustomer) SetName(v string) *GetZoneListResponseDataResultsCustomer {
  s.Name = &v
  return s
}

type GetZoneListPaths struct {
}

func (s GetZoneListPaths) String() string {
  return tea.Prettify(s)
}

func (s GetZoneListPaths) GoString() string {
  return s.String()
}

type GetZoneListParameters struct {
  // {"en":"zone name's keywords, use to search zones", "zh_CN":"zone名称的关键字，用于搜索"}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // {"en":"Page number, use to get the specific zone list", "zh_CN":"页码，用于获取指定范围的zone列表"}
  Page *int `json:"page,omitempty" xml:"page,omitempty"`
  // {"en":"Number of zones been returned each time", "zh_CN":"每次返回zone的数量"}
  PageSize *int `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s GetZoneListParameters) String() string {
  return tea.Prettify(s)
}

func (s GetZoneListParameters) GoString() string {
  return s.String()
}

func (s *GetZoneListParameters) SetName(v string) *GetZoneListParameters {
  s.Name = &v
  return s
}

func (s *GetZoneListParameters) SetPage(v int) *GetZoneListParameters {
  s.Page = &v
  return s
}

func (s *GetZoneListParameters) SetPageSize(v int) *GetZoneListParameters {
  s.PageSize = &v
  return s
}

type GetZoneListRequestHeader struct {
}

func (s GetZoneListRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetZoneListRequestHeader) GoString() string {
  return s.String()
}

type GetZoneListResponseHeader struct {
}

func (s GetZoneListResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetZoneListResponseHeader) GoString() string {
  return s.String()
}




type BatchDeleteZoneRequest struct {
}

func (s BatchDeleteZoneRequest) String() string {
  return tea.Prettify(s)
}

func (s BatchDeleteZoneRequest) GoString() string {
  return s.String()
}

type BatchDeleteZoneResponse struct {
  // {"en":"If it fails, the name of the failed zone is returned, separated by commas", "zh_CN":"如果失败，返回的是失败zone的名称，逗号分割"}
  Data *string `json:"data,omitempty" xml:"data,omitempty" require:"true"`
  // {"en":"code,success is 0", "zh_CN":"状态码,成功为0"}
  Code *int `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"error message or Success", "zh_CN":"错误信息或Success"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s BatchDeleteZoneResponse) String() string {
  return tea.Prettify(s)
}

func (s BatchDeleteZoneResponse) GoString() string {
  return s.String()
}

func (s *BatchDeleteZoneResponse) SetData(v string) *BatchDeleteZoneResponse {
  s.Data = &v
  return s
}

func (s *BatchDeleteZoneResponse) SetCode(v int) *BatchDeleteZoneResponse {
  s.Code = &v
  return s
}

func (s *BatchDeleteZoneResponse) SetMessage(v string) *BatchDeleteZoneResponse {
  s.Message = &v
  return s
}

type BatchDeleteZonePaths struct {
  // {"en":"zone's ID list, split by comma", "zh_CN":"zone的ID列表，逗号分割"}
  ZoneIds *string `json:"zoneIds,omitempty" xml:"zoneIds,omitempty" require:"true"`
}

func (s BatchDeleteZonePaths) String() string {
  return tea.Prettify(s)
}

func (s BatchDeleteZonePaths) GoString() string {
  return s.String()
}

func (s *BatchDeleteZonePaths) SetZoneIds(v string) *BatchDeleteZonePaths {
  s.ZoneIds = &v
  return s
}

type BatchDeleteZoneParameters struct {
}

func (s BatchDeleteZoneParameters) String() string {
  return tea.Prettify(s)
}

func (s BatchDeleteZoneParameters) GoString() string {
  return s.String()
}

type BatchDeleteZoneRequestHeader struct {
}

func (s BatchDeleteZoneRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s BatchDeleteZoneRequestHeader) GoString() string {
  return s.String()
}

type BatchDeleteZoneResponseHeader struct {
}

func (s BatchDeleteZoneResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s BatchDeleteZoneResponseHeader) GoString() string {
  return s.String()
}




type DeleteRecordByIDRequest struct {
}

func (s DeleteRecordByIDRequest) String() string {
  return tea.Prettify(s)
}

func (s DeleteRecordByIDRequest) GoString() string {
  return s.String()
}

type DeleteRecordByIDResponse struct {
  // {"en":"return data", "zh_CN":"返回值"}
  Data *DeleteRecordByIDResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
  // {"en":"code,success is 0", "zh_CN":"状态码,成功为0"}
  Code *int `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"error message or Success", "zh_CN":"错误信息或Success"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s DeleteRecordByIDResponse) String() string {
  return tea.Prettify(s)
}

func (s DeleteRecordByIDResponse) GoString() string {
  return s.String()
}

func (s *DeleteRecordByIDResponse) SetData(v *DeleteRecordByIDResponseData) *DeleteRecordByIDResponse {
  s.Data = v
  return s
}

func (s *DeleteRecordByIDResponse) SetCode(v int) *DeleteRecordByIDResponse {
  s.Code = &v
  return s
}

func (s *DeleteRecordByIDResponse) SetMessage(v string) *DeleteRecordByIDResponse {
  s.Message = &v
  return s
}

type DeleteRecordByIDResponseData struct {
}

func (s DeleteRecordByIDResponseData) String() string {
  return tea.Prettify(s)
}

func (s DeleteRecordByIDResponseData) GoString() string {
  return s.String()
}

type DeleteRecordByIDPaths struct {
  // {"en":"zone's ID", "zh_CN":"zone的ID"}
  ZoneId *int `json:"zoneId,omitempty" xml:"zoneId,omitempty" require:"true"`
  // {"en":"record's ID", "zh_CN":"record的ID"}
  RecordId *int `json:"recordId,omitempty" xml:"recordId,omitempty" require:"true"`
}

func (s DeleteRecordByIDPaths) String() string {
  return tea.Prettify(s)
}

func (s DeleteRecordByIDPaths) GoString() string {
  return s.String()
}

func (s *DeleteRecordByIDPaths) SetZoneId(v int) *DeleteRecordByIDPaths {
  s.ZoneId = &v
  return s
}

func (s *DeleteRecordByIDPaths) SetRecordId(v int) *DeleteRecordByIDPaths {
  s.RecordId = &v
  return s
}

type DeleteRecordByIDParameters struct {
}

func (s DeleteRecordByIDParameters) String() string {
  return tea.Prettify(s)
}

func (s DeleteRecordByIDParameters) GoString() string {
  return s.String()
}

type DeleteRecordByIDRequestHeader struct {
}

func (s DeleteRecordByIDRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteRecordByIDRequestHeader) GoString() string {
  return s.String()
}

type DeleteRecordByIDResponseHeader struct {
}

func (s DeleteRecordByIDResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteRecordByIDResponseHeader) GoString() string {
  return s.String()
}




type DeleteZoneRequest struct {
}

func (s DeleteZoneRequest) String() string {
  return tea.Prettify(s)
}

func (s DeleteZoneRequest) GoString() string {
  return s.String()
}

type DeleteZoneResponse struct {
  // {"en":"return data", "zh_CN":"返回值"}
  Data *string `json:"data,omitempty" xml:"data,omitempty" require:"true"`
  // {"en":"code,success is 0", "zh_CN":"状态码,成功为0"}
  Code *int `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"error message or Success", "zh_CN":"错误信息或Success"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s DeleteZoneResponse) String() string {
  return tea.Prettify(s)
}

func (s DeleteZoneResponse) GoString() string {
  return s.String()
}

func (s *DeleteZoneResponse) SetData(v string) *DeleteZoneResponse {
  s.Data = &v
  return s
}

func (s *DeleteZoneResponse) SetCode(v int) *DeleteZoneResponse {
  s.Code = &v
  return s
}

func (s *DeleteZoneResponse) SetMessage(v string) *DeleteZoneResponse {
  s.Message = &v
  return s
}

type DeleteZonePaths struct {
  // {"en":"zone's ID", "zh_CN":"zone的ID"}
  ZoneId *int `json:"zoneId,omitempty" xml:"zoneId,omitempty" require:"true"`
}

func (s DeleteZonePaths) String() string {
  return tea.Prettify(s)
}

func (s DeleteZonePaths) GoString() string {
  return s.String()
}

func (s *DeleteZonePaths) SetZoneId(v int) *DeleteZonePaths {
  s.ZoneId = &v
  return s
}

type DeleteZoneParameters struct {
}

func (s DeleteZoneParameters) String() string {
  return tea.Prettify(s)
}

func (s DeleteZoneParameters) GoString() string {
  return s.String()
}

type DeleteZoneRequestHeader struct {
}

func (s DeleteZoneRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteZoneRequestHeader) GoString() string {
  return s.String()
}

type DeleteZoneResponseHeader struct {
}

func (s DeleteZoneResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteZoneResponseHeader) GoString() string {
  return s.String()
}




type BulkCreateRecordsByZoneRequest struct {
  // {"en":"request data", "zh_CN":"请求数据"}
  Data []*BulkCreateRecordsByZoneRequestData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s BulkCreateRecordsByZoneRequest) String() string {
  return tea.Prettify(s)
}

func (s BulkCreateRecordsByZoneRequest) GoString() string {
  return s.String()
}

func (s *BulkCreateRecordsByZoneRequest) SetData(v []*BulkCreateRecordsByZoneRequestData) *BulkCreateRecordsByZoneRequest {
  s.Data = v
  return s
}

type BulkCreateRecordsByZoneRequestData struct     {
  // {"en":"host name", "zh_CN":"域名头"}
  HostName *string `json:"hostName,omitempty" xml:"hostName,omitempty" require:"true"`
  // {"en":"(SOA,A,AAAA,CNAME,TXT,SRV,MX,ALIAS,RP,SPF,PTR,NS)", "zh_CN":"record类型值（SOA,A,AAAA,CNAME,TXT,SRV,MX,ALIAS,RP,SPF,PTR,NS）"}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
  // {"en":"Record value (SRV value: '${priority} ${weight} ${port} ${target}'; RP value: '${mailboxName} ${domainName}' )", "zh_CN":"记录值(SRV值为：'${priority} ${weight} ${port} ${target}'; RP值为：'${mailboxName} ${domainName}')"}
  Value *string `json:"value,omitempty" xml:"value,omitempty" require:"true"`
  // {"en":"ttl", "zh_CN":"ttl"}
  Ttl *int `json:"ttl,omitempty" xml:"ttl,omitempty" require:"true"`
  // {"en":"MX preference (Only available on MX type, default is 10)", "zh_CN":"MX类型的优先级(只有在MX类型才有, 默认为10)"}
  Data *int `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s BulkCreateRecordsByZoneRequestData) String() string {
  return tea.Prettify(s)
}

func (s BulkCreateRecordsByZoneRequestData) GoString() string {
  return s.String()
}

func (s *BulkCreateRecordsByZoneRequestData) SetHostName(v string) *BulkCreateRecordsByZoneRequestData {
  s.HostName = &v
  return s
}

func (s *BulkCreateRecordsByZoneRequestData) SetType(v string) *BulkCreateRecordsByZoneRequestData {
  s.Type = &v
  return s
}

func (s *BulkCreateRecordsByZoneRequestData) SetValue(v string) *BulkCreateRecordsByZoneRequestData {
  s.Value = &v
  return s
}

func (s *BulkCreateRecordsByZoneRequestData) SetTtl(v int) *BulkCreateRecordsByZoneRequestData {
  s.Ttl = &v
  return s
}

func (s *BulkCreateRecordsByZoneRequestData) SetData(v int) *BulkCreateRecordsByZoneRequestData {
  s.Data = &v
  return s
}

type BulkCreateRecordsByZoneResponse struct {
  // {"en":"return data", "zh_CN":"返回值"}
  Data *BulkCreateRecordsByZoneResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
  // {"en":"code,success is 0", "zh_CN":"状态码,成功为0"}
  Code *int `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"error message or Success", "zh_CN":"错误信息或Success"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s BulkCreateRecordsByZoneResponse) String() string {
  return tea.Prettify(s)
}

func (s BulkCreateRecordsByZoneResponse) GoString() string {
  return s.String()
}

func (s *BulkCreateRecordsByZoneResponse) SetData(v *BulkCreateRecordsByZoneResponseData) *BulkCreateRecordsByZoneResponse {
  s.Data = v
  return s
}

func (s *BulkCreateRecordsByZoneResponse) SetCode(v int) *BulkCreateRecordsByZoneResponse {
  s.Code = &v
  return s
}

func (s *BulkCreateRecordsByZoneResponse) SetMessage(v string) *BulkCreateRecordsByZoneResponse {
  s.Message = &v
  return s
}

type BulkCreateRecordsByZoneResponseData struct {
  // {"en":"record type(SOA,A,AAAA,CNAME,TXT,SRV,MX,ALIAS,RP,SPF,PTR,NS)", "zh_CN":"数据类型（SOA,A,AAAA,CNAME,TXT,SRV,MX,ALIAS,RP,SPF,PTR,NS）"}
  A []*BulkCreateRecordsByZoneResponseDataA `json:"A,omitempty" xml:"A,omitempty" require:"true" type:"Repeated"`
}

func (s BulkCreateRecordsByZoneResponseData) String() string {
  return tea.Prettify(s)
}

func (s BulkCreateRecordsByZoneResponseData) GoString() string {
  return s.String()
}

func (s *BulkCreateRecordsByZoneResponseData) SetA(v []*BulkCreateRecordsByZoneResponseDataA) *BulkCreateRecordsByZoneResponseData {
  s.A = v
  return s
}

type BulkCreateRecordsByZoneResponseDataA struct     {
  // {"en":"zone id", "zh_CN":"zone的ID"}
  ZoneId *int `json:"zoneId,omitempty" xml:"zoneId,omitempty" require:"true"`
  // {"en":"record id", "zh_CN":"record的ID"}
  RecordId *int `json:"recordId,omitempty" xml:"recordId,omitempty" require:"true"`
  // {"en":"host name", "zh_CN":"域名头"}
  HostName *string `json:"hostName,omitempty" xml:"hostName,omitempty" require:"true"`
  // {"en":"record type", "zh_CN":"record类型值"}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
  // {"en":"Record value (RP and SRV do not have this attribute)", "zh_CN":"记录值(RP与SRV没有这个属性)"}
  Value *string `json:"value,omitempty" xml:"value,omitempty" require:"true"`
  // {"en":"ttl", "zh_CN":"ttl"}
  Ttl *int `json:"ttl,omitempty" xml:"ttl,omitempty" require:"true"`
  // {"en":"priority(Only available on SRV type)", "zh_CN":"优先(只有在SRV类型才有)"}
  Priority *int `json:"priority,omitempty" xml:"priority,omitempty" require:"true"`
  // {"en":"weight(Only available on SRV type)", "zh_CN":"权重(只有在SRV类型才有)"}
  Weight *int `json:"weight,omitempty" xml:"weight,omitempty" require:"true"`
  // {"en":"port(Only available on SRV type)", "zh_CN":"端口(只有在SRV类型才有)"}
  Port *int `json:"port,omitempty" xml:"port,omitempty" require:"true"`
  // {"en":"target domain(Only available on SRV type)", "zh_CN":"目标域名(只有在SRV类型才有)"}
  Target *string `json:"target,omitempty" xml:"target,omitempty" require:"true"`
  // {"en":"mail box name(Only available on RP type)", "zh_CN":"电子邮件域名(只有在RP类型才有)"}
  MailboxName *string `json:"mailboxName,omitempty" xml:"mailboxName,omitempty" require:"true"`
  // {"en":"domain name(Only available on RP type)", "zh_CN":"域名(只有在RP类型才有)"}
  DomainName *string `json:"domainName,omitempty" xml:"domainName,omitempty" require:"true"`
  // {"en":"preference(Only available on MX type)", "zh_CN":"优先级(只有在MX类型才有)"}
  Preference *int `json:"preference,omitempty" xml:"preference,omitempty" require:"true"`
  // {'en':'Unicode name of hostName','zh_CN':'域名头的unicode名称'}
  HostUnicode *string `json:"hostUnicode,omitempty" xml:"hostUnicode,omitempty" require:"true"`
}

func (s BulkCreateRecordsByZoneResponseDataA) String() string {
  return tea.Prettify(s)
}

func (s BulkCreateRecordsByZoneResponseDataA) GoString() string {
  return s.String()
}

func (s *BulkCreateRecordsByZoneResponseDataA) SetZoneId(v int) *BulkCreateRecordsByZoneResponseDataA {
  s.ZoneId = &v
  return s
}

func (s *BulkCreateRecordsByZoneResponseDataA) SetRecordId(v int) *BulkCreateRecordsByZoneResponseDataA {
  s.RecordId = &v
  return s
}

func (s *BulkCreateRecordsByZoneResponseDataA) SetHostName(v string) *BulkCreateRecordsByZoneResponseDataA {
  s.HostName = &v
  return s
}

func (s *BulkCreateRecordsByZoneResponseDataA) SetType(v string) *BulkCreateRecordsByZoneResponseDataA {
  s.Type = &v
  return s
}

func (s *BulkCreateRecordsByZoneResponseDataA) SetValue(v string) *BulkCreateRecordsByZoneResponseDataA {
  s.Value = &v
  return s
}

func (s *BulkCreateRecordsByZoneResponseDataA) SetTtl(v int) *BulkCreateRecordsByZoneResponseDataA {
  s.Ttl = &v
  return s
}

func (s *BulkCreateRecordsByZoneResponseDataA) SetPriority(v int) *BulkCreateRecordsByZoneResponseDataA {
  s.Priority = &v
  return s
}

func (s *BulkCreateRecordsByZoneResponseDataA) SetWeight(v int) *BulkCreateRecordsByZoneResponseDataA {
  s.Weight = &v
  return s
}

func (s *BulkCreateRecordsByZoneResponseDataA) SetPort(v int) *BulkCreateRecordsByZoneResponseDataA {
  s.Port = &v
  return s
}

func (s *BulkCreateRecordsByZoneResponseDataA) SetTarget(v string) *BulkCreateRecordsByZoneResponseDataA {
  s.Target = &v
  return s
}

func (s *BulkCreateRecordsByZoneResponseDataA) SetMailboxName(v string) *BulkCreateRecordsByZoneResponseDataA {
  s.MailboxName = &v
  return s
}

func (s *BulkCreateRecordsByZoneResponseDataA) SetDomainName(v string) *BulkCreateRecordsByZoneResponseDataA {
  s.DomainName = &v
  return s
}

func (s *BulkCreateRecordsByZoneResponseDataA) SetPreference(v int) *BulkCreateRecordsByZoneResponseDataA {
  s.Preference = &v
  return s
}

func (s *BulkCreateRecordsByZoneResponseDataA) SetHostUnicode(v string) *BulkCreateRecordsByZoneResponseDataA {
  s.HostUnicode = &v
  return s
}

type BulkCreateRecordsByZonePaths struct {
  // {"en":"zone's ID", "zh_CN":"zone的ID"}
  ZoneId *int `json:"zoneId,omitempty" xml:"zoneId,omitempty" require:"true"`
}

func (s BulkCreateRecordsByZonePaths) String() string {
  return tea.Prettify(s)
}

func (s BulkCreateRecordsByZonePaths) GoString() string {
  return s.String()
}

func (s *BulkCreateRecordsByZonePaths) SetZoneId(v int) *BulkCreateRecordsByZonePaths {
  s.ZoneId = &v
  return s
}

type BulkCreateRecordsByZoneParameters struct {
}

func (s BulkCreateRecordsByZoneParameters) String() string {
  return tea.Prettify(s)
}

func (s BulkCreateRecordsByZoneParameters) GoString() string {
  return s.String()
}

type BulkCreateRecordsByZoneRequestHeader struct {
}

func (s BulkCreateRecordsByZoneRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s BulkCreateRecordsByZoneRequestHeader) GoString() string {
  return s.String()
}

type BulkCreateRecordsByZoneResponseHeader struct {
}

func (s BulkCreateRecordsByZoneResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s BulkCreateRecordsByZoneResponseHeader) GoString() string {
  return s.String()
}




type GetRecordbyIDRequest struct {
}

func (s GetRecordbyIDRequest) String() string {
  return tea.Prettify(s)
}

func (s GetRecordbyIDRequest) GoString() string {
  return s.String()
}

type GetRecordbyIDResponse struct {
  // {"en":"return data", "zh_CN":"返回值"}
  Data *GetRecordbyIDResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
  // {"en":"code,success is 0", "zh_CN":"状态码,成功为0"}
  Code *int `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"error message or Success", "zh_CN":"错误信息或Success"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s GetRecordbyIDResponse) String() string {
  return tea.Prettify(s)
}

func (s GetRecordbyIDResponse) GoString() string {
  return s.String()
}

func (s *GetRecordbyIDResponse) SetData(v *GetRecordbyIDResponseData) *GetRecordbyIDResponse {
  s.Data = v
  return s
}

func (s *GetRecordbyIDResponse) SetCode(v int) *GetRecordbyIDResponse {
  s.Code = &v
  return s
}

func (s *GetRecordbyIDResponse) SetMessage(v string) *GetRecordbyIDResponse {
  s.Message = &v
  return s
}

type GetRecordbyIDResponseData struct {
  // {"en":"zone id", "zh_CN":"zone的ID"}
  ZoneId *int `json:"zoneId,omitempty" xml:"zoneId,omitempty" require:"true"`
  // {"en":"record id", "zh_CN":"record的ID"}
  RecordId *int `json:"recordId,omitempty" xml:"recordId,omitempty" require:"true"`
  // {"en":"host name", "zh_CN":"域名头"}
  HostName *string `json:"hostName,omitempty" xml:"hostName,omitempty" require:"true"`
  // {"en":"(SOA,A,AAAA,CNAME,TXT,SRV,MX,ALIAS,RP,SPF,PTR,NS)", "zh_CN":"record类型值（SOA,A,AAAA,CNAME,TXT,SRV,MX,ALIAS,RP,SPF,PTR,NS）"}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
  // {"en":"Record value (RP and SRV do not have this attribute)", "zh_CN":"记录值(RP与SRV没有这个属性)"}
  Value *string `json:"value,omitempty" xml:"value,omitempty" require:"true"`
  // {"en":"ttl", "zh_CN":"ttl"}
  Ttl *int `json:"ttl,omitempty" xml:"ttl,omitempty" require:"true"`
  // {"en":"priority(Only available on SRV type)", "zh_CN":"优先(只有在SRV类型才有)"}
  Priority *int `json:"priority,omitempty" xml:"priority,omitempty" require:"true"`
  // {"en":"weight(Only available on SRV type)", "zh_CN":"权重(只有在SRV类型才有)"}
  Weight *int `json:"weight,omitempty" xml:"weight,omitempty" require:"true"`
  // {"en":"port(Only available on SRV type)", "zh_CN":"端口(只有在SRV类型才有)"}
  Port *int `json:"port,omitempty" xml:"port,omitempty" require:"true"`
  // {"en":"target domain(Only available on SRV type)", "zh_CN":"目标域名(只有在SRV类型才有)"}
  Target *string `json:"target,omitempty" xml:"target,omitempty" require:"true"`
  // {"en":"mail box name(Only available on RP type)", "zh_CN":"电子邮件域名(只有在RP类型才有)"}
  MailboxName *string `json:"mailboxName,omitempty" xml:"mailboxName,omitempty" require:"true"`
  // {"en":"domain name(Only available on RP type)", "zh_CN":"域名(只有在RP类型才有)"}
  DomainName *string `json:"domainName,omitempty" xml:"domainName,omitempty" require:"true"`
  // {"en":"preference(Only available on MX type)", "zh_CN":"优先级(只有在MX类型才有)"}
  Preference *int `json:"preference,omitempty" xml:"preference,omitempty" require:"true"`
  // {'en':'Unicode name of hostName','zh_CN':'域名头的unicode名称'}
  HostUnicode *string `json:"hostUnicode,omitempty" xml:"hostUnicode,omitempty" require:"true"`
}

func (s GetRecordbyIDResponseData) String() string {
  return tea.Prettify(s)
}

func (s GetRecordbyIDResponseData) GoString() string {
  return s.String()
}

func (s *GetRecordbyIDResponseData) SetZoneId(v int) *GetRecordbyIDResponseData {
  s.ZoneId = &v
  return s
}

func (s *GetRecordbyIDResponseData) SetRecordId(v int) *GetRecordbyIDResponseData {
  s.RecordId = &v
  return s
}

func (s *GetRecordbyIDResponseData) SetHostName(v string) *GetRecordbyIDResponseData {
  s.HostName = &v
  return s
}

func (s *GetRecordbyIDResponseData) SetType(v string) *GetRecordbyIDResponseData {
  s.Type = &v
  return s
}

func (s *GetRecordbyIDResponseData) SetValue(v string) *GetRecordbyIDResponseData {
  s.Value = &v
  return s
}

func (s *GetRecordbyIDResponseData) SetTtl(v int) *GetRecordbyIDResponseData {
  s.Ttl = &v
  return s
}

func (s *GetRecordbyIDResponseData) SetPriority(v int) *GetRecordbyIDResponseData {
  s.Priority = &v
  return s
}

func (s *GetRecordbyIDResponseData) SetWeight(v int) *GetRecordbyIDResponseData {
  s.Weight = &v
  return s
}

func (s *GetRecordbyIDResponseData) SetPort(v int) *GetRecordbyIDResponseData {
  s.Port = &v
  return s
}

func (s *GetRecordbyIDResponseData) SetTarget(v string) *GetRecordbyIDResponseData {
  s.Target = &v
  return s
}

func (s *GetRecordbyIDResponseData) SetMailboxName(v string) *GetRecordbyIDResponseData {
  s.MailboxName = &v
  return s
}

func (s *GetRecordbyIDResponseData) SetDomainName(v string) *GetRecordbyIDResponseData {
  s.DomainName = &v
  return s
}

func (s *GetRecordbyIDResponseData) SetPreference(v int) *GetRecordbyIDResponseData {
  s.Preference = &v
  return s
}

func (s *GetRecordbyIDResponseData) SetHostUnicode(v string) *GetRecordbyIDResponseData {
  s.HostUnicode = &v
  return s
}

type GetRecordbyIDPaths struct {
  // {"en":"zone's ID", "zh_CN":"zone的ID"}
  ZoneId *int `json:"zoneId,omitempty" xml:"zoneId,omitempty" require:"true"`
  // {"en":"record's ID", "zh_CN":"Record的ID"}
  RecordId *int `json:"recordId,omitempty" xml:"recordId,omitempty" require:"true"`
}

func (s GetRecordbyIDPaths) String() string {
  return tea.Prettify(s)
}

func (s GetRecordbyIDPaths) GoString() string {
  return s.String()
}

func (s *GetRecordbyIDPaths) SetZoneId(v int) *GetRecordbyIDPaths {
  s.ZoneId = &v
  return s
}

func (s *GetRecordbyIDPaths) SetRecordId(v int) *GetRecordbyIDPaths {
  s.RecordId = &v
  return s
}

type GetRecordbyIDParameters struct {
}

func (s GetRecordbyIDParameters) String() string {
  return tea.Prettify(s)
}

func (s GetRecordbyIDParameters) GoString() string {
  return s.String()
}

type GetRecordbyIDRequestHeader struct {
}

func (s GetRecordbyIDRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetRecordbyIDRequestHeader) GoString() string {
  return s.String()
}

type GetRecordbyIDResponseHeader struct {
}

func (s GetRecordbyIDResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetRecordbyIDResponseHeader) GoString() string {
  return s.String()
}




type UpdateZoneRequest struct {
  // {"en":"Default TTL setting of this zone", "zh_CN":"zone的默认TTL配置"}
  Ttl *int `json:"ttl,omitempty" xml:"ttl,omitempty" require:"true"`
}

func (s UpdateZoneRequest) String() string {
  return tea.Prettify(s)
}

func (s UpdateZoneRequest) GoString() string {
  return s.String()
}

func (s *UpdateZoneRequest) SetTtl(v int) *UpdateZoneRequest {
  s.Ttl = &v
  return s
}

type UpdateZoneResponse struct {
}

func (s UpdateZoneResponse) String() string {
  return tea.Prettify(s)
}

func (s UpdateZoneResponse) GoString() string {
  return s.String()
}

type UpdateZonePaths struct {
  // {"en":"zone's ID", "zh_CN":"zone的ID"}
  ZoneId *int `json:"zoneId,omitempty" xml:"zoneId,omitempty" require:"true"`
}

func (s UpdateZonePaths) String() string {
  return tea.Prettify(s)
}

func (s UpdateZonePaths) GoString() string {
  return s.String()
}

func (s *UpdateZonePaths) SetZoneId(v int) *UpdateZonePaths {
  s.ZoneId = &v
  return s
}

type UpdateZoneParameters struct {
}

func (s UpdateZoneParameters) String() string {
  return tea.Prettify(s)
}

func (s UpdateZoneParameters) GoString() string {
  return s.String()
}

type UpdateZoneRequestHeader struct {
}

func (s UpdateZoneRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s UpdateZoneRequestHeader) GoString() string {
  return s.String()
}

type UpdateZoneResponseHeader struct {
}

func (s UpdateZoneResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s UpdateZoneResponseHeader) GoString() string {
  return s.String()
}




type UpdateZTSBulkRequest struct {
  // {"en":"ZTS configuration object","zh_CN":"ZTS配置对象"}
  Data []*UpdateZTSBulkRequestData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s UpdateZTSBulkRequest) String() string {
  return tea.Prettify(s)
}

func (s UpdateZTSBulkRequest) GoString() string {
  return s.String()
}

func (s *UpdateZTSBulkRequest) SetData(v []*UpdateZTSBulkRequestData) *UpdateZTSBulkRequest {
  s.Data = v
  return s
}

type UpdateZTSBulkRequestData struct     {
  // {"en":"zoneId","zh_CN":"zoneId"}
  ZoneId *int `json:"zoneId,omitempty" xml:"zoneId,omitempty" require:"true"`
  // {"en":"Status: 0 - Enabled, 1 - Disabled","zh_CN":"状态，0-启用，1-停用"}
  St *int `json:"st,omitempty" xml:"st,omitempty" require:"true"`
  // {"en":"","zh_CN":""}
  Axfr *UpdateZTSBulkRequestDataAxfr `json:"axfr,omitempty" xml:"axfr,omitempty" require:"true" type:"Struct"`
}

func (s UpdateZTSBulkRequestData) String() string {
  return tea.Prettify(s)
}

func (s UpdateZTSBulkRequestData) GoString() string {
  return s.String()
}

func (s *UpdateZTSBulkRequestData) SetZoneId(v int) *UpdateZTSBulkRequestData {
  s.ZoneId = &v
  return s
}

func (s *UpdateZTSBulkRequestData) SetSt(v int) *UpdateZTSBulkRequestData {
  s.St = &v
  return s
}

func (s *UpdateZTSBulkRequestData) SetAxfr(v *UpdateZTSBulkRequestDataAxfr) *UpdateZTSBulkRequestData {
  s.Axfr = v
  return s
}

type UpdateZTSBulkRequestDataAxfr struct {
  // {"en":"Source DNS server address","zh_CN":"源DNS服务器地址"}
  MasterAddr *string `json:"masterAddr,omitempty" xml:"masterAddr,omitempty" require:"true"`
  // {"en":"Optional: Source DNS server port, default is 53","zh_CN":"选填，源DNS服务器端口，默认53"}
  Port *int `json:"port,omitempty" xml:"port,omitempty"`
  // {"en":"Optional: Enable NS override, use CloudDNS's NS to override the NS configured on the original DNS, default is disabled.","zh_CN":"选填，启用NS覆盖，用CloudDNS的NS覆盖原DNS上配的NS，默认不启用"}
  EnableSoaNsRewriting *bool `json:"enableSoaNsRewriting,omitempty" xml:"enableSoaNsRewriting,omitempty"`
  // {"en":"Optional: Leave blank to disable encryption","zh_CN":"选填，不填不加密"}
  Key *UpdateZTSBulkRequestDataAxfrKey `json:"key,omitempty" xml:"key,omitempty" type:"Struct"`
  // {"en":"","zh_CN":""}
  SoaPoll *UpdateZTSBulkRequestDataAxfrSoaPoll `json:"soaPoll,omitempty" xml:"soaPoll,omitempty" type:"Struct"`
  // {"en":"Optional: Whitelist configuration, specify which ZTS servers are allowed to access the source DNS server","zh_CN":"选填，白名单配置，允许哪些ZTS服务器接入源DNS服务器"}
  Whitelist []*string `json:"whitelist,omitempty" xml:"whitelist,omitempty" type:"Repeated"`
}

func (s UpdateZTSBulkRequestDataAxfr) String() string {
  return tea.Prettify(s)
}

func (s UpdateZTSBulkRequestDataAxfr) GoString() string {
  return s.String()
}

func (s *UpdateZTSBulkRequestDataAxfr) SetMasterAddr(v string) *UpdateZTSBulkRequestDataAxfr {
  s.MasterAddr = &v
  return s
}

func (s *UpdateZTSBulkRequestDataAxfr) SetPort(v int) *UpdateZTSBulkRequestDataAxfr {
  s.Port = &v
  return s
}

func (s *UpdateZTSBulkRequestDataAxfr) SetEnableSoaNsRewriting(v bool) *UpdateZTSBulkRequestDataAxfr {
  s.EnableSoaNsRewriting = &v
  return s
}

func (s *UpdateZTSBulkRequestDataAxfr) SetKey(v *UpdateZTSBulkRequestDataAxfrKey) *UpdateZTSBulkRequestDataAxfr {
  s.Key = v
  return s
}

func (s *UpdateZTSBulkRequestDataAxfr) SetSoaPoll(v *UpdateZTSBulkRequestDataAxfrSoaPoll) *UpdateZTSBulkRequestDataAxfr {
  s.SoaPoll = v
  return s
}

func (s *UpdateZTSBulkRequestDataAxfr) SetWhitelist(v []*string) *UpdateZTSBulkRequestDataAxfr {
  s.Whitelist = v
  return s
}

type UpdateZTSBulkRequestDataAxfrKey struct {
  // {"en":"Encryption algorithms currently supported: hmac-md5/hmac-sha1/hmac-sha224/hmac-sha256/hmac-sha384/hmac-sha512","zh_CN":"加密算法，目前支持：hmac-md5/hmac-sha1/hmac-sha224/hmac-sha256/hmac-sha384/hmac-sha512"}
  Algorithm *string `json:"algorithm,omitempty" xml:"algorithm,omitempty" require:"true"`
  // {"en":"Encryption username.","zh_CN":"加密用户名"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"Encryption key","zh_CN":"加密密钥"}
  Secret *string `json:"secret,omitempty" xml:"secret,omitempty" require:"true"`
}

func (s UpdateZTSBulkRequestDataAxfrKey) String() string {
  return tea.Prettify(s)
}

func (s UpdateZTSBulkRequestDataAxfrKey) GoString() string {
  return s.String()
}

func (s *UpdateZTSBulkRequestDataAxfrKey) SetAlgorithm(v string) *UpdateZTSBulkRequestDataAxfrKey {
  s.Algorithm = &v
  return s
}

func (s *UpdateZTSBulkRequestDataAxfrKey) SetName(v string) *UpdateZTSBulkRequestDataAxfrKey {
  s.Name = &v
  return s
}

func (s *UpdateZTSBulkRequestDataAxfrKey) SetSecret(v string) *UpdateZTSBulkRequestDataAxfrKey {
  s.Secret = &v
  return s
}

type UpdateZTSBulkRequestDataAxfrSoaPoll struct {
  // {"en":"Synchronization frequency, in seconds, default is 300","zh_CN":"同步频率，单位秒，默认300"}
  PsInterval *int `json:"psInterval,omitempty" xml:"psInterval,omitempty"`
  // {"en":"Timeout, in seconds, default is 10","zh_CN":"超时时间，单位秒，默认10"}
  PsTimeout *int `json:"psTimeout,omitempty" xml:"psTimeout,omitempty"`
}

func (s UpdateZTSBulkRequestDataAxfrSoaPoll) String() string {
  return tea.Prettify(s)
}

func (s UpdateZTSBulkRequestDataAxfrSoaPoll) GoString() string {
  return s.String()
}

func (s *UpdateZTSBulkRequestDataAxfrSoaPoll) SetPsInterval(v int) *UpdateZTSBulkRequestDataAxfrSoaPoll {
  s.PsInterval = &v
  return s
}

func (s *UpdateZTSBulkRequestDataAxfrSoaPoll) SetPsTimeout(v int) *UpdateZTSBulkRequestDataAxfrSoaPoll {
  s.PsTimeout = &v
  return s
}

type UpdateZTSBulkRequestHeader struct {
}

func (s UpdateZTSBulkRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s UpdateZTSBulkRequestHeader) GoString() string {
  return s.String()
}

type UpdateZTSBulkPaths struct {
}

func (s UpdateZTSBulkPaths) String() string {
  return tea.Prettify(s)
}

func (s UpdateZTSBulkPaths) GoString() string {
  return s.String()
}

type UpdateZTSBulkParameters struct {
}

func (s UpdateZTSBulkParameters) String() string {
  return tea.Prettify(s)
}

func (s UpdateZTSBulkParameters) GoString() string {
  return s.String()
}

type UpdateZTSBulkResponse struct {
  // {"en":"Status code, 0 indicates success.","zh_CN":"状态码,成功为0"}
  Code *int `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Error message or Success","zh_CN":"错误信息或Success"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s UpdateZTSBulkResponse) String() string {
  return tea.Prettify(s)
}

func (s UpdateZTSBulkResponse) GoString() string {
  return s.String()
}

func (s *UpdateZTSBulkResponse) SetCode(v int) *UpdateZTSBulkResponse {
  s.Code = &v
  return s
}

func (s *UpdateZTSBulkResponse) SetMessage(v string) *UpdateZTSBulkResponse {
  s.Message = &v
  return s
}

type UpdateZTSBulkResponseHeader struct {
}

func (s UpdateZTSBulkResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s UpdateZTSBulkResponseHeader) GoString() string {
  return s.String()
}




